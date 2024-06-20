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
    def fm0(p0, p1, globalState):
        return (True) == ((not(False)) and (not(True)))

    @staticmethod
    def fm1(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxninyjx"))

    @staticmethod
    def fm2(globalState):
        return _dafny.CodePoint('s')

    @staticmethod
    def fm3(globalState):
        return 332

    @staticmethod
    def fm4(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.MultiSet([False, False])), D1_DC3(_dafny.MultiSet([True]))])) + (_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(False))]))), D1_DC3(_dafny.MultiSet([True, False])), D1_DC3(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))]))

    @staticmethod
    def fm16(p0, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, not(not(True))]), _dafny.SeqWithoutIsStrInference([True])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, False, True]), _dafny.SeqWithoutIsStrInference([False])]))

    @staticmethod
    def fm17(globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lheve")))) * (len(_dafny.SeqWithoutIsStrInference([False]))) for d_0_i0_ in range(default__.abs(960))])

    @staticmethod
    def fm18(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psg"))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wufuwmsuw")) for d_1_i0_ in range(default__.abs(918))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfu"))])))

    @staticmethod
    def fm19(p0, globalState):
        return _dafny.MultiSet([(146) * ((_dafny.MultiSet([False, not(False)])).cardinality)])

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(not(not(False))), True])) + (_dafny.SeqWithoutIsStrInference([False, False]))

    @staticmethod
    def fm21(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([241])).Elements:
                d_3_v0_: int = compr_0_
                if (d_3_v0_) in (_dafny.SeqWithoutIsStrInference([241])):
                    coll0_ = coll0_.union(_dafny.Set([(d_3_v0_) + (-99)]))
            return _dafny.Set(coll0_)
        return D4_DC12(True, (_dafny.Map({-894: _dafny.SeqWithoutIsStrInference([283])})) != (_dafny.Map({328: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2_i0_ in range(default__.abs(831))]))])})), ((0) - ((_dafny.MultiSet([878, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du"))), len(_dafny.Map({True: True})), len(iife0_()
)])).cardinality)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnlbg")))))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return D8_DC20((_dafny.SeqWithoutIsStrInference([not(not(False))])) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm23(p0, p1, globalState):
        return ((_dafny.Map({(D11_DC31(False, 948, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_4_i0_ in range(default__.abs(-910))])), True)).cf56: True})) | (_dafny.Map({False: False}))) | ((_dafny.Map({False: False})) | (_dafny.Map({False: not(not(not(True)))})))

    @staticmethod
    def fm24(p0, p1, globalState):
        return D8_DC23()

    @staticmethod
    def fm25(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: D8
            for compr_1_ in (_dafny.MultiSet((D14_DC38(_dafny.SeqWithoutIsStrInference([D8_DC23()]))).cf70)).Elements:
                d_5_v0_: D8 = compr_1_
                if (d_5_v0_) in (_dafny.MultiSet((D14_DC38(_dafny.SeqWithoutIsStrInference([D8_DC23()]))).cf70)):
                    coll1_ = coll1_.union(_dafny.Set([d_5_v0_]))
            return _dafny.Set(coll1_)
        return _dafny.MultiSet([(iife1_()
        ) | (_dafny.Set({D8_DC23()})), _dafny.Set({D8_DC23()}), (_dafny.Set({D8_DC23(), D8_DC23(), D8_DC23()})) | (_dafny.Set({D8_DC23(), D8_DC23(), D8_DC23()}))])

    @staticmethod
    def fm26(globalState):
        if (_dafny.SeqWithoutIsStrInference([not(True)])) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])})):
            return D8_DC21(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([324]))})), True, not(True), True, 194)
        elif True:
            return D8_DC21(len(_dafny.Set({-729, (0) - (-109)})), False, True, True, -579)

    @staticmethod
    def fm27(globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: str
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n')])).Elements:
                d_6_v0_: str = compr_2_
                if (d_6_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n')])):
                    coll2_ = coll2_.union(_dafny.Set([d_6_v0_]))
            return _dafny.Set(coll2_)
        return (iife2_()
        ) | ((_dafny.Set({_dafny.CodePoint('v')})) - (_dafny.Set({_dafny.CodePoint('n')})))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in (_dafny.Set({-250})).Elements:
                d_7_v0_: int = compr_3_
                if (d_7_v0_) in (_dafny.Set({-250})):
                    coll3_ = coll3_.union(_dafny.Set([(d_7_v0_) + (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')]))})))]))
            return _dafny.Set(coll3_)
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(223, 890):
                d_8_v1_: int = compr_4_
                if ((223) <= (d_8_v1_)) and ((d_8_v1_) < (890)):
                    coll4_ = coll4_.union(_dafny.Set([default__.safeDivisionInt(d_8_v1_, 291)]))
            return _dafny.Set(coll4_)
        def iife5_():
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, not(False)])): True})).keys.Elements:
                    d_10_v2_: int = compr_7_
                    if (d_10_v2_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, not(False)])): True})):
                        coll7_[(d_10_v2_) * (15)] = True
                return _dafny.Map(coll7_)
            coll5_ = _dafny.Set()
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, not(False)])): True})).keys.Elements:
                    d_10_v2_: int = compr_6_
                    if (d_10_v2_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, not(False)])): True})):
                        coll6_[(d_10_v2_) * (15)] = True
                return _dafny.Map(coll6_)
            compr_5_: _dafny.Set
            for compr_5_ in (_dafny.MultiSet([_dafny.Set({-278, len(_dafny.SeqWithoutIsStrInference([436]))}), _dafny.Set({len(iife6_()
            )})])).Elements:
                d_11_v3_: _dafny.Set = compr_5_
                if (d_11_v3_) in (_dafny.MultiSet([_dafny.Set({-278, len(_dafny.SeqWithoutIsStrInference([436]))}), _dafny.Set({len(iife7_()
                )})])):
                    coll5_ = coll5_.union(_dafny.Set([d_11_v3_]))
            return _dafny.Set(coll5_)
        return (_dafny.Set({_dafny.Set({708, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnk"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True), True])), len(_dafny.Map({802: False}))])), 475}), iife3_()
        })) | ((_dafny.Set({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(0) - (len(iife4_()
)) for d_9_i0_ in range(default__.abs(769))]))})})) | (iife5_()
        ))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        source0_ = D9_DC27(False, _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('v')]))}))
        if source0_.is_DC26:
            d_12___mcc_h0_ = source0_.cf45
            d_13___mcc_h1_ = source0_.cf46
            d_14_cf46_ = d_13___mcc_h1_
            d_15_cf45_ = d_12___mcc_h0_
            def iife8_():
                coll8_ = _dafny.Set()
                compr_8_: int
                for compr_8_ in _dafny.IntegerRange(809, 957):
                    d_16_v0_: int = compr_8_
                    if ((809) <= (d_16_v0_)) and ((d_16_v0_) < (957)):
                        coll8_ = coll8_.union(_dafny.Set([default__.safeModuloInt(d_16_v0_, (d_14_cf46_).f26)]))
                return _dafny.Set(coll8_)
            return _dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efkgpwsh")))}), iife8_()
            ])
        elif source0_.is_DC27:
            d_17___mcc_h2_ = source0_.cf47
            d_18___mcc_h3_ = source0_.cf48
            d_19_cf48_ = d_18___mcc_h3_
            d_20_cf47_ = d_17___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([d_19_cf48_])) + (_dafny.SeqWithoutIsStrInference([d_19_cf48_]))
        elif True:
            d_21___mcc_h4_ = source0_.cf44
            d_22_cf44_ = d_21___mcc_h4_
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_23_i0_ in range(default__.abs(697))]))])).Elements:
                    d_24_v1_: int = compr_9_
                    if (d_24_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_23_i0_ in range(default__.abs(697))]))])):
                        coll9_ = coll9_.union(_dafny.Set([(d_24_v1_) - (-706)]))
                return _dafny.Set(coll9_)
            def iife10_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(836, -360):
                    d_25_v2_: int = compr_10_
                    if ((836) <= (d_25_v2_)) and ((d_25_v2_) < (-360)):
                        coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_25_v2_, 13)]))
                return _dafny.Set(coll10_)
            return (_dafny.SeqWithoutIsStrInference([iife9_()
            ])) + (_dafny.SeqWithoutIsStrInference([iife10_()
            ]))

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D1_DC5(320, True) for d_26_i0_ in range(default__.abs(-865))]) if True else _dafny.SeqWithoutIsStrInference([D1_DC5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwmmy"))), False), D1_DC5(301, True)])))) - (_dafny.MultiSet([D1_DC5(-624, False), D1_DC5(len(_dafny.SeqWithoutIsStrInference([True])), False)]))

    @staticmethod
    def fm31(p0, globalState):
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in (_dafny.Map({530: False})).keys.Elements:
                d_27_v0_: int = compr_11_
                if (d_27_v0_) in (_dafny.Map({530: False})):
                    def iife12_():
                        coll12_ = _dafny.Set()
                        compr_12_: str
                        for compr_12_ in (_dafny.Map({_dafny.CodePoint('j'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnkgmf"))})).keys.Elements:
                            d_28_v1_: str = compr_12_
                            if (d_28_v1_) in (_dafny.Map({_dafny.CodePoint('j'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnkgmf"))})):
                                coll12_ = coll12_.union(_dafny.Set([d_28_v1_]))
                        return _dafny.Set(coll12_)
                    coll11_ = coll11_.union(_dafny.Set([default__.safeDivisionInt(d_27_v0_, len(iife12_()
))]))
            return _dafny.Set(coll11_)
        return ((_dafny.Set({154})) - (iife11_()
        )) | (_dafny.Set({881}))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        if not(not(not (False) or (True))):
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehocjrafq")))])])
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(_dafny.MultiSet([True, (D8_DC21((_dafny.MultiSet([387])).cardinality, True, True, not(True), 990)).cf36, True])).cardinality]) for d_29_i0_ in range(default__.abs(790))])

    @staticmethod
    def fm33(p0, globalState):
        return D1_DC4((-221) != ((D11_DC30(False, -315)).cf52), (-347) + ((_dafny.MultiSet([True])).cardinality), _dafny.Map({False: not(True)}))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return D1_DC5((_dafny.MultiSet([False])).cardinality, not(False))

    @staticmethod
    def fm35(p0, p1, globalState):
        def iife13_():
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(-241, 396):
                    d_30_v0_: int = compr_16_
                    if ((-241) <= (d_30_v0_)) and ((d_30_v0_) < (396)):
                        coll16_[default__.safeDivisionInt(d_30_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwfabehw"))))] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_31_i0_ in range(default__.abs(248))])))
                return _dafny.Map(coll16_)
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(-481, 89):
                    d_32_v1_: int = compr_17_
                    if ((-481) <= (d_32_v1_)) and ((d_32_v1_) < (89)):
                        coll17_[(d_32_v1_) * (len(_dafny.Set({741})))] = -324
                return _dafny.Map(coll17_)
            coll13_ = _dafny.Set()
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(-241, 396):
                    d_30_v0_: int = compr_14_
                    if ((-241) <= (d_30_v0_)) and ((d_30_v0_) < (396)):
                        coll14_[default__.safeDivisionInt(d_30_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwfabehw"))))] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_31_i0_ in range(default__.abs(248))])))
                return _dafny.Map(coll14_)
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-481, 89):
                    d_32_v1_: int = compr_15_
                    if ((-481) <= (d_32_v1_)) and ((d_32_v1_) < (89)):
                        coll15_[(d_32_v1_) * (len(_dafny.Set({741})))] = -324
                return _dafny.Map(coll15_)
            compr_13_: _dafny.Map
            for compr_13_ in ((_dafny.SeqWithoutIsStrInference([iife14_()
            ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): -954}), iife15_()
            ]))).Elements:
                d_33_v2_: _dafny.Map = compr_13_
                if (d_33_v2_) in ((_dafny.SeqWithoutIsStrInference([iife16_()
                ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): -954}), iife17_()
                ]))):
                    coll13_ = coll13_.union(_dafny.Set([d_33_v2_]))
            return _dafny.Set(coll13_)
        return iife13_()
        

    @staticmethod
    def fm36(p0, p1, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))) | (_dafny.MultiSet([True, not(False)]))) - (_dafny.MultiSet([not(False)]))

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        return D0_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_34_i0_ in range(default__.abs(893))]))

    @staticmethod
    def fm38(p0, p1, globalState):
        if (_dafny.Set({False})) in (_dafny.Map({_dafny.Set({False}): True})):
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(918, -354):
                    d_36_v0_: int = compr_18_
                    if ((918) <= (d_36_v0_)) and ((d_36_v0_) < (-354)):
                        coll18_[default__.safeModuloInt(d_36_v0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_37_i2_ in range(default__.abs(232))])) for d_38_i1_ in range(default__.abs(571))])))] = not(True)
                return _dafny.Map(coll18_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([412 for d_35_i0_ in range(default__.abs(318))])): not(True)})) | (iife18_()
            )
        elif True:
            return (D15_DC40(_dafny.Map({(D4_DC12(True, False, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-207])), 640])).cardinality)).cf23: True}))).cf73

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: int
            for compr_19_ in (_dafny.Map({937: -278})).keys.Elements:
                d_39_v0_: int = compr_19_
                if (d_39_v0_) in (_dafny.Map({937: -278})):
                    coll19_ = coll19_.union(_dafny.Set([(d_39_v0_) - (313)]))
            return _dafny.Set(coll19_)
        return _dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usfsg")))]), (D16_DC44(_dafny.MultiSet([len(_dafny.Map({False: False})), len(_dafny.Set({-290}))]))).cf82, (_dafny.MultiSet([(_dafny.MultiSet([352, (0) - ((_dafny.MultiSet([4])).cardinality), (0) - (len(_dafny.SeqWithoutIsStrInference([884, 762])))])).cardinality, (0) - (-430)])).intersection(_dafny.MultiSet([-57, (0) - (-119)])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({-953: 886})): D15_DC41(len(_dafny.Map({True: len(_dafny.Map({True: 483}))})), _dafny.SeqWithoutIsStrInference([296]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gp")))})), len(_dafny.Set({_dafny.Set({True}), _dafny.Set({False, (D6_DC16(True, True)).cf30}), _dafny.Set({False, True}), _dafny.Set({False})})), 580, 128, (_dafny.MultiSet([816, len(_dafny.SeqWithoutIsStrInference([(0) - (-119)])), 131])).cardinality]))) | (_dafny.MultiSet([len(iife19_()
        ), len(_dafny.Map({204: False})), 707, (0) - (len(_dafny.SeqWithoutIsStrInference([not(not(True))]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))]))])

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.CodePoint('l'), _dafny.CodePoint('g')])) | ((_dafny.MultiSet([_dafny.CodePoint('l'), _dafny.CodePoint('r'), _dafny.CodePoint('v'), _dafny.CodePoint('f'), _dafny.CodePoint('t')])) | (_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('o')])))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_40_v0_: bool
        d_40_v0_ = False
        if d_40_v0_:
            d_41_v1_: int
            d_41_v1_ = 9
            d_42_v2_: _dafny.Map
            d_42_v2_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([default__.fm3(globalState)])) != (_dafny.SeqWithoutIsStrInference([d_41_v1_, p2])): d_40_v0_})
            d_41_v1_ = len(d_42_v2_)
            d_43_v3_: _dafny.Map
            d_43_v3_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_41_v1_, 469]): not(d_40_v0_)})
            d_44_v4_: _dafny.Seq
            d_44_v4_ = _dafny.SeqWithoutIsStrInference([len(d_43_v3_), d_41_v1_, (0) - (d_41_v1_)])
            d_45_v5_: _dafny.Set
            d_45_v5_ = _dafny.Set({d_44_v4_, _dafny.SeqWithoutIsStrInference([d_41_v1_]), d_44_v4_})
            d_45_v5_ = d_45_v5_
            d_46_v6_: _dafny.MultiSet
            d_46_v6_ = _dafny.MultiSet([p1, p1])
            d_47_v7_: str
            d_47_v7_ = _dafny.CodePoint('m')
            d_46_v6_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_48_i0_ in range(default__.abs(281))]), (p1).set(default__.safeIndex(len(p1), len(p1)), d_47_v7_)])
            d_49_v8_: _dafny.Array
            def lambda0_(d_50_v0_):
                def lambda1_(d_51_i1_):
                    return d_50_v0_

                return lambda1_

            init0_ = lambda0_(d_40_v0_)
            nw0_ = _dafny.Array(None, 24)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_49_v8_ = nw0_
            index0_ = default__.safeIndex(613, (d_49_v8_).length(0))
            (d_49_v8_)[index0_] = not(default__.fm0(d_40_v0_, d_40_v0_, globalState))
            index1_ = default__.safeIndex(613, (d_49_v8_).length(0))
            rhs0_ = d_40_v0_
            rhs1_ = True
            lhs0_ = globalState
            lhs1_ = d_49_v8_
            lhs2_ = default__.safeIndex(613, (d_49_v8_).length(0))
            lhs0_.f2 = rhs0_
            lhs1_[lhs2_] = rhs1_
            index2_ = default__.safeIndex(613, (d_49_v8_).length(0))
            (d_49_v8_)[index2_] = default__.fm0((p1) <= (p1), True, globalState)
        elif True:
            d_52_v9_: _dafny.Map
            d_52_v9_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_53_i2_ in range(default__.abs(-2))]): p2})
            if (d_52_v9_) != (d_52_v9_):
                d_54_v10_: int
                d_54_v10_ = 844
                d_55_v11_: _dafny.Map
                d_55_v11_ = _dafny.Map({d_40_v0_: d_40_v0_})
                d_56_v12_: D1
                d_56_v12_ = D1_DC4(d_40_v0_, p2, (d_55_v11_).set(d_40_v0_, d_40_v0_))
                d_57_v13_: _dafny.MultiSet
                d_57_v13_ = _dafny.MultiSet([(d_56_v12_).cf7, d_40_v0_])
                d_54_v10_ = (default__.safeDivisionInt(p0, 323)) - (default__.safeModuloInt((d_57_v13_).cardinality, p0))
                d_58_v14_: _dafny.Set
                d_58_v14_ = _dafny.Set({d_40_v0_})
                d_59_v15_: _dafny.Seq
                d_59_v15_ = _dafny.SeqWithoutIsStrInference([len(p1)])
                d_60_v16_: C8
                nw1_ = C8()
                nw1_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_61_i3_ in range(default__.abs(909))]), d_58_v14_, d_40_v0_, ((d_59_v15_) + (d_59_v15_)).set(default__.safeIndex(p2, len((d_59_v15_) + (d_59_v15_))), p3))
                d_60_v16_ = nw1_
                d_62_v17_: _dafny.Array
                def lambda2_(d_63_v0_):
                    def lambda3_(d_64_i4_):
                        return d_63_v0_

                    return lambda3_

                init1_ = lambda2_(d_40_v0_)
                nw2_ = _dafny.Array(None, 18)
                for i0_1_ in range(nw2_.length(0)):
                    nw2_[i0_1_] = init1_(i0_1_)
                d_62_v17_ = nw2_
                index3_ = default__.safeIndex(392, (d_62_v17_).length(0))
                (d_62_v17_)[index3_] = d_40_v0_
                index4_ = default__.safeIndex(392, (d_62_v17_).length(0))
                (d_62_v17_)[index4_] = d_40_v0_
                d_65_v18_: _dafny.Array
                nw3_ = _dafny.Array(int(0), 17)
                d_65_v18_ = nw3_
                index5_ = default__.safeIndex(816, (d_65_v18_).length(0))
                (d_65_v18_)[index5_] = p0
                index6_ = default__.safeIndex(816, (d_65_v18_).length(0))
                (d_65_v18_)[index6_] = (default__.fm3(globalState)) - (p3)
                d_66_v19_: C2
                nw4_ = C2()
                nw4_.ctor__(len(d_59_v15_))
                d_66_v19_ = nw4_
            elif True:
                d_67_v20_: _dafny.Array
                nw5_ = _dafny.Array(D2.default()(), 18)
                d_67_v20_ = nw5_
                d_68_v21_: _dafny.Array
                d_68_v21_ = d_67_v20_
                d_68_v21_ = (d_67_v20_ if d_40_v0_ else d_67_v20_)
                d_69_v22_: _dafny.Seq
                d_69_v22_ = _dafny.SeqWithoutIsStrInference([672])
                d_70_v23_: _dafny.Array
                def lambda4_(d_71_p1_):
                    def lambda5_(d_72_i6_):
                        return default__.safeDivisionInt(d_72_i6_, len(d_71_p1_))

                    return lambda5_

                init2_ = lambda4_(p1)
                nw6_ = _dafny.Array(None, 28)
                for i0_2_ in range(nw6_.length(0)):
                    nw6_[i0_2_] = init2_(i0_2_)
                d_70_v23_ = nw6_
                d_73_v24_: _dafny.Seq
                d_73_v24_ = _dafny.SeqWithoutIsStrInference([d_70_v23_])
                d_74_v25_: _dafny.Array
                nw7_ = _dafny.Array(None, 13)
                nw7_[int(0)] = d_69_v22_
                nw7_[int(1)] = d_69_v22_
                nw7_[int(2)] = d_69_v22_
                nw7_[int(3)] = _dafny.SeqWithoutIsStrInference([p0 for d_75_i5_ in range(default__.abs(102))])
                nw7_[int(4)] = d_69_v22_
                nw7_[int(5)] = d_69_v22_
                nw7_[int(6)] = _dafny.SeqWithoutIsStrInference([p2, len(d_73_v24_), p2, p2, p0])
                nw7_[int(7)] = (d_69_v22_) + (default__.fm17(globalState))
                nw7_[int(8)] = default__.fm17(globalState)
                nw7_[int(9)] = d_69_v22_
                nw7_[int(10)] = d_69_v22_
                nw7_[int(11)] = d_69_v22_
                nw7_[int(12)] = d_69_v22_
                d_74_v25_ = nw7_
                index7_ = default__.safeIndex(829, (d_74_v25_).length(0))
                (d_74_v25_)[index7_] = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_40_v0_}))])
                d_76_v26_: _dafny.Map
                d_76_v26_ = _dafny.Map({not(d_40_v0_): d_69_v22_})
                d_77_v27_: _dafny.Seq
                d_77_v27_ = _dafny.SeqWithoutIsStrInference([((d_76_v26_)[d_40_v0_] if (d_40_v0_) in (d_76_v26_) else _dafny.SeqWithoutIsStrInference([p2, p0, p2]))])
                index8_ = default__.safeIndex(829, (d_74_v25_).length(0))
                (d_74_v25_)[index8_] = (default__.fm17(globalState)) + ((d_77_v27_)[default__.safeIndex(p3, len(d_77_v27_))])
                d_78_v28_: int
                d_78_v28_ = 907
                d_78_v28_ = p2
                d_79_v29_: _dafny.MultiSet
                d_79_v29_ = _dafny.MultiSet([d_78_v28_])
                (globalState).f2 = ((d_79_v29_) - (_dafny.MultiSet([p3, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p2: d_40_v0_})) for d_80_i7_ in range(default__.abs(327))]))]))).issubset(_dafny.MultiSet((d_74_v25_)[default__.safeIndex(829, (d_74_v25_).length(0))]))
                d_78_v28_ = p0
            d_81_v30_: _dafny.Array
            def lambda6_(d_82_v0_):
                def lambda7_(d_83_i8_):
                    return not (False) or (d_82_v0_)

                return lambda7_

            init3_ = lambda6_(d_40_v0_)
            nw8_ = _dafny.Array(None, 28)
            for i0_3_ in range(nw8_.length(0)):
                nw8_[i0_3_] = init3_(i0_3_)
            d_81_v30_ = nw8_
            index9_ = default__.safeIndex(37, (d_81_v30_).length(0))
            (d_81_v30_)[index9_] = d_40_v0_
            index10_ = default__.safeIndex(37, (d_81_v30_).length(0))
            (d_81_v30_)[index10_] = d_40_v0_
            d_84_v31_: T0
            nw9_ = C5()
            nw9_.ctor__(d_40_v0_, _dafny.SeqWithoutIsStrInference([p0 for d_85_i9_ in range(default__.abs(168))]))
            d_84_v31_ = nw9_
            d_86_v32_: str
            d_86_v32_ = _dafny.CodePoint('m')
            d_87_v33_: _dafny.Map
            d_87_v33_ = _dafny.Map({d_84_v31_: d_86_v32_})
            (globalState).f2 = (d_84_v31_) not in (d_87_v33_)
            d_88_v34_: C5
            nw10_ = C5()
            nw10_.ctor__(not ((d_81_v30_)[default__.safeIndex(37, (d_81_v30_).length(0))]) or (d_84_v31_.f16), _dafny.SeqWithoutIsStrInference([p0 for d_89_i10_ in range(default__.abs(172))]))
            d_88_v34_ = nw10_
            d_86_v32_ = d_86_v32_
        d_90_v35_: int
        d_90_v35_ = -579
        d_91_v36_: D4
        d_91_v36_ = D4_DC12(True, False, p2)
        pat_let_tv0_ = p3
        pat_let_tv1_ = p3
        pat_let_tv2_ = d_40_v0_
        pat_let_tv3_ = d_40_v0_
        pat_let_tv4_ = d_40_v0_
        pat_let_tv5_ = d_40_v0_
        pat_let_tv6_ = d_40_v0_
        def lambda8_(source1_):
            if source1_.is_DC12:
                d_92___mcc_h0_ = source1_.cf21
                d_93___mcc_h1_ = source1_.cf22
                d_94___mcc_h2_ = source1_.cf23
                d_95_cf23_ = d_94___mcc_h2_
                d_96_cf22_ = d_93___mcc_h1_
                d_97_cf21_ = d_92___mcc_h0_
                return pat_let_tv0_
            elif source1_.is_DC13:
                d_98___mcc_h3_ = source1_.cf24
                d_99___mcc_h4_ = source1_.cf25
                d_100___mcc_h5_ = source1_.cf26
                d_101_cf26_ = d_100___mcc_h5_
                d_102_cf25_ = d_99___mcc_h4_
                d_103_cf24_ = d_98___mcc_h3_
                return (0) - (pat_let_tv1_)
            elif True:
                d_104___mcc_h6_ = source1_.cf20
                d_105_cf20_ = d_104___mcc_h6_
                return (0) - (len((_dafny.SeqWithoutIsStrInference([pat_let_tv2_, not(pat_let_tv3_), pat_let_tv4_])) + (_dafny.SeqWithoutIsStrInference([pat_let_tv5_, pat_let_tv6_]))))

        d_90_v35_ = lambda8_(d_91_v36_)
        d_106_v37_: str
        d_106_v37_ = _dafny.CodePoint('j')
        d_107_v38_: _dafny.Seq
        d_107_v38_ = _dafny.SeqWithoutIsStrInference([d_40_v0_])
        d_108_v39_: _dafny.Map
        d_108_v39_ = _dafny.Map({d_106_v37_: d_107_v38_})
        d_109_v40_: _dafny.Seq
        d_109_v40_ = _dafny.SeqWithoutIsStrInference([d_107_v38_])
        d_108_v39_ = (d_108_v39_).set(d_106_v37_, (d_107_v38_) + ((d_109_v40_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsxwg"))), len(d_109_v40_))]))
        d_110_v41_: _dafny.Map
        d_110_v41_ = _dafny.Map({p0: p3})
        d_111_v42_: _dafny.Seq
        d_111_v42_ = _dafny.SeqWithoutIsStrInference([d_90_v35_, p3])
        d_112_v43_: C3
        nw11_ = C3()
        nw11_.ctor__(p1, d_40_v0_, d_111_v42_)
        d_112_v43_ = nw11_
        d_113_v44_: D12
        d_113_v44_ = D12_DC35(d_40_v0_, d_40_v0_, d_110_v41_, True, d_112_v43_)
        d_114_v45_: _dafny.Array
        nw12_ = _dafny.Array(None, 12)
        nw12_[int(0)] = d_40_v0_
        nw12_[int(1)] = (d_113_v44_).cf64
        nw12_[int(2)] = d_40_v0_
        nw12_[int(3)] = d_40_v0_
        nw12_[int(4)] = d_40_v0_
        nw12_[int(5)] = d_40_v0_
        nw12_[int(6)] = d_40_v0_
        nw12_[int(7)] = d_40_v0_
        nw12_[int(8)] = d_40_v0_
        nw12_[int(9)] = d_40_v0_
        nw12_[int(10)] = d_40_v0_
        nw12_[int(11)] = False
        d_114_v45_ = nw12_
        d_115_v46_: D0
        d_115_v46_ = D0_DC2((_dafny.MultiSet(d_107_v38_)).cardinality, p0, p3, d_114_v45_)
        d_116_v47_: C6
        nw13_ = C6()
        def iife20_(_pat_let0_0):
            def iife21_(d_117_dt__update__tmp_h0_):
                def iife22_(_pat_let1_0):
                    def iife23_(d_118_dt__update_hcf3_h0_):
                        return D0_DC2((d_117_dt__update__tmp_h0_).cf2, d_118_dt__update_hcf3_h0_, (d_117_dt__update__tmp_h0_).cf4, (d_117_dt__update__tmp_h0_).cf5)
                    return iife23_(_pat_let1_0)
                return iife22_(-516)
            return iife21_(_pat_let0_0)
        nw13_.ctor__(iife20_(d_115_v46_))
        d_116_v47_ = nw13_
        d_119_v48_: _dafny.Set
        d_119_v48_ = _dafny.Set({d_116_v47_})
        hi0_ = len(d_119_v48_)
        for d_120_i11_ in range(d_90_v35_, hi0_):
            d_121_v49_: _dafny.Array
            nw14_ = _dafny.Array(None, 5)
            nw14_[int(0)] = d_114_v45_
            nw14_[int(1)] = d_114_v45_
            nw14_[int(2)] = d_114_v45_
            nw14_[int(3)] = (d_114_v45_ if d_40_v0_ else d_114_v45_)
            nw14_[int(4)] = d_114_v45_
            d_121_v49_ = nw14_
            index11_ = default__.safeIndex(987, (d_121_v49_).length(0))
            (d_121_v49_)[index11_] = d_114_v45_
            d_122_v50_: _dafny.Map
            d_122_v50_ = _dafny.Map({d_40_v0_: 260})
            d_123_v51_: _dafny.Array
            nw15_ = _dafny.Array(None, 3)
            nw15_[int(0)] = d_111_v42_
            nw15_[int(1)] = _dafny.SeqWithoutIsStrInference([d_120_i11_ for d_124_i12_ in range(default__.abs(793))])
            nw15_[int(2)] = d_111_v42_
            d_123_v51_ = nw15_
            d_125_v52_: D2
            d_125_v52_ = D2_DC8(d_122_v50_, d_123_v51_, d_40_v0_, d_40_v0_, d_120_i11_)
            d_126_v53_: D6
            d_126_v53_ = D6_DC17(d_125_v52_)
            index12_ = default__.safeIndex(987, (d_121_v49_).length(0))
            rhs2_ = d_114_v45_
            rhs3_ = d_126_v53_
            lhs3_ = d_121_v49_
            lhs4_ = default__.safeIndex(987, (d_121_v49_).length(0))
            lhs3_[lhs4_] = rhs2_
            d_126_v53_ = rhs3_
            d_106_v37_ = _dafny.CodePoint('v')
            d_127_v54_: _dafny.Array
            def lambda9_(d_128_v37_, d_129_p1_):
                def lambda10_(d_130_i13_):
                    return (_dafny.SeqWithoutIsStrInference([d_128_v37_ for d_131_i14_ in range(default__.abs(229))])) + (d_129_p1_)

                return lambda10_

            init4_ = lambda9_(d_106_v37_, p1)
            nw16_ = _dafny.Array(None, 28)
            for i0_4_ in range(nw16_.length(0)):
                nw16_[i0_4_] = init4_(i0_4_)
            d_127_v54_ = nw16_
            index13_ = default__.safeIndex(35, (d_127_v54_).length(0))
            (d_127_v54_)[index13_] = p1
            index14_ = default__.safeIndex(35, (d_127_v54_).length(0))
            (d_127_v54_)[index14_] = (d_112_v43_).f24
            rhs4_ = not(d_40_v0_)
            rhs5_ = (d_107_v38_)[default__.safeIndex(len(d_111_v42_), len(d_107_v38_))]
            lhs5_ = globalState
            lhs6_ = globalState
            lhs5_.f2 = rhs4_
            lhs6_.f2 = rhs5_
        d_132_v55_: _dafny.Set
        d_132_v55_ = _dafny.Set({default__.fm0(d_40_v0_, d_40_v0_, globalState), d_40_v0_})
        d_133_v56_: _dafny.Seq
        d_133_v56_ = _dafny.SeqWithoutIsStrInference([d_132_v55_, d_132_v55_, d_132_v55_])
        if not (((d_133_v56_)[default__.safeIndex((d_111_v42_)[default__.safeIndex(d_90_v35_, len(d_111_v42_))], len(d_133_v56_))]).issubset(d_132_v55_)) or (False):
            d_134_v57_: _dafny.Map
            d_134_v57_ = _dafny.Map({d_40_v0_: p3})
            d_135_v58_: C2
            nw17_ = C2()
            nw17_.ctor__((_dafny.MultiSet([d_40_v0_])).cardinality)
            d_135_v58_ = nw17_
            d_136_v59_: _dafny.Map
            d_136_v59_ = _dafny.Map({p2: d_135_v58_})
            d_137_v60_: _dafny.Seq
            d_137_v60_ = _dafny.SeqWithoutIsStrInference([d_136_v59_, d_136_v59_])
            d_138_v61_: T0
            nw18_ = C5()
            nw18_.ctor__(d_40_v0_, (d_116_v47_).fm6(d_90_v35_, d_40_v0_, d_106_v37_, p1, globalState))
            d_138_v61_ = nw18_
            d_139_v62_: _dafny.Map
            d_139_v62_ = _dafny.Map({p0: d_40_v0_})
            d_140_v63_: _dafny.Map
            d_140_v63_ = _dafny.Map({d_138_v61_: _dafny.Map({d_40_v0_: len(d_139_v62_)})})
            d_141_v64_: _dafny.Array
            nw19_ = _dafny.Array(None, 23)
            nw19_[int(0)] = d_134_v57_
            nw19_[int(1)] = (d_134_v57_).set(d_40_v0_, 194)
            nw19_[int(2)] = d_134_v57_
            nw19_[int(3)] = (d_134_v57_) | (d_134_v57_)
            nw19_[int(4)] = _dafny.Map({d_40_v0_: p2})
            nw19_[int(5)] = d_134_v57_
            nw19_[int(6)] = (d_134_v57_) | (d_134_v57_)
            nw19_[int(7)] = d_134_v57_
            nw19_[int(8)] = d_134_v57_
            nw19_[int(9)] = d_134_v57_
            nw19_[int(10)] = d_134_v57_
            nw19_[int(11)] = d_134_v57_
            nw19_[int(12)] = (d_134_v57_) | (d_134_v57_)
            nw19_[int(13)] = d_134_v57_
            nw19_[int(14)] = _dafny.Map({d_40_v0_: len(d_111_v42_)})
            nw19_[int(15)] = (d_134_v57_) | (d_134_v57_)
            nw19_[int(16)] = d_134_v57_
            nw19_[int(17)] = d_134_v57_
            nw19_[int(18)] = d_134_v57_
            nw19_[int(19)] = (d_134_v57_) | ((d_134_v57_).set(d_40_v0_, len(d_137_v60_)))
            nw19_[int(20)] = ((d_140_v63_)[d_138_v61_] if (d_138_v61_) in (d_140_v63_) else d_134_v57_)
            nw19_[int(21)] = d_134_v57_
            nw19_[int(22)] = _dafny.Map({d_138_v61_.f16: p3})
            d_141_v64_ = nw19_
            d_142_v65_: _dafny.Map
            d_142_v65_ = _dafny.Map({d_40_v0_: (d_134_v57_).set(d_40_v0_, (d_135_v58_).f25)})
            index15_ = default__.safeIndex(396, (d_141_v64_).length(0))
            (d_141_v64_)[index15_] = ((d_142_v65_)[d_40_v0_] if (d_40_v0_) in (d_142_v65_) else d_134_v57_)
            d_143_v66_: C0
            nw20_ = C0()
            nw20_.ctor__(p3, d_114_v45_)
            d_143_v66_ = nw20_
            d_144_v67_: T1
            nw21_ = C1()
            nw21_.ctor__(((d_135_v58_).f25) - (903), len(_dafny.Set({(d_107_v38_).set(default__.safeIndex(p0, len(d_107_v38_)), False)})), (d_107_v38_)[default__.safeIndex((D9_DC26((d_135_v58_).f25, d_143_v66_)).cf45, len(d_107_v38_))], d_111_v42_, d_90_v35_)
            d_144_v67_ = nw21_
            index16_ = default__.safeIndex(396, (d_141_v64_).length(0))
            rhs6_ = d_134_v57_
            rhs7_ = (d_135_v58_).f25
            rhs8_ = d_144_v67_
            lhs7_ = d_141_v64_
            lhs8_ = default__.safeIndex(396, (d_141_v64_).length(0))
            lhs7_[lhs8_] = rhs6_
            d_90_v35_ = rhs7_
            d_144_v67_ = rhs8_
            d_145_v68_: _dafny.Array
            nw22_ = _dafny.Array(None, 9)
            nw22_[int(0)] = (d_144_v67_).fm13(d_144_v67_.f16, globalState)
            nw22_[int(1)] = p3
            nw22_[int(2)] = (d_144_v67_).f28
            nw22_[int(3)] = (d_135_v58_).f25
            nw22_[int(4)] = p0
            nw22_[int(5)] = p2
            nw22_[int(6)] = p2
            nw22_[int(7)] = len(p1)
            nw22_[int(8)] = (d_135_v58_).f25
            d_145_v68_ = nw22_
            d_146_v69_: _dafny.Array
            nw23_ = _dafny.Array(None, 13)
            nw23_[int(0)] = (d_145_v68_ if d_144_v67_.f16 else d_145_v68_)
            nw23_[int(1)] = d_145_v68_
            nw23_[int(2)] = d_145_v68_
            nw23_[int(3)] = d_145_v68_
            nw23_[int(4)] = d_145_v68_
            nw23_[int(5)] = d_145_v68_
            nw23_[int(6)] = d_145_v68_
            nw23_[int(7)] = d_145_v68_
            nw23_[int(8)] = d_145_v68_
            nw23_[int(9)] = d_145_v68_
            nw23_[int(10)] = d_145_v68_
            nw23_[int(11)] = d_145_v68_
            nw23_[int(12)] = d_145_v68_
            d_146_v69_ = nw23_
            index17_ = default__.safeIndex(127, (d_146_v69_).length(0))
            (d_146_v69_)[index17_] = d_145_v68_
            index18_ = default__.safeIndex(127, (d_146_v69_).length(0))
            (d_146_v69_)[index18_] = d_145_v68_
            if d_138_v61_.f16:
                d_90_v35_ = d_90_v35_
                (globalState).f5 = d_40_v0_
                d_147_v70_: _dafny.Array
                nw24_ = _dafny.Array(None, 27)
                nw24_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_143_v66_).f26])
                nw24_[int(1)] = (d_144_v67_).f17
                nw24_[int(2)] = (d_138_v61_).f17
                nw24_[int(3)] = (d_144_v67_).f17
                nw24_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_144_v67_).f28 for d_148_i15_ in range(default__.abs(-381))])
                nw24_[int(5)] = (d_138_v61_).f17
                nw24_[int(6)] = (d_144_v67_).f17
                nw24_[int(7)] = _dafny.SeqWithoutIsStrInference([len((d_138_v61_).f17) for d_149_i16_ in range(default__.abs(230))])
                nw24_[int(8)] = (d_138_v61_).f17
                nw24_[int(9)] = (d_144_v67_).f17
                nw24_[int(10)] = (d_138_v61_).f17
                nw24_[int(11)] = (d_144_v67_).f17
                nw24_[int(12)] = (d_138_v61_).f17
                nw24_[int(13)] = d_111_v42_
                nw24_[int(14)] = (d_138_v61_).f17
                nw24_[int(15)] = (d_138_v61_).f17
                nw24_[int(16)] = (_dafny.SeqWithoutIsStrInference([(d_144_v67_).f28 for d_150_i17_ in range(default__.abs(503))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([(d_144_v67_).f28 for d_151_i17_ in range(default__.abs(503))]))), -749)
                nw24_[int(17)] = (d_144_v67_).f17
                nw24_[int(18)] = _dafny.SeqWithoutIsStrInference([p2])
                nw24_[int(19)] = (d_138_v61_).f17
                nw24_[int(20)] = (d_138_v61_).f17
                nw24_[int(21)] = d_111_v42_
                nw24_[int(22)] = (d_144_v67_).f17
                nw24_[int(23)] = (d_144_v67_).f17
                nw24_[int(24)] = (d_138_v61_).f17
                nw24_[int(25)] = d_111_v42_
                nw24_[int(26)] = (d_144_v67_).f17
                d_147_v70_ = nw24_
                d_152_v71_: D2
                d_152_v71_ = D2_DC8(_dafny.Map({d_144_v67_.f16: d_90_v35_}), d_147_v70_, d_40_v0_, d_138_v61_.f16, d_90_v35_)
                d_153_v72_: D6
                d_153_v72_ = D6_DC17(d_152_v71_)
                pat_let_tv7_ = d_152_v71_
                d_154_v73_: _dafny.MultiSet
                def iife24_(_pat_let2_0):
                    def iife25_(d_155_dt__update__tmp_h1_):
                        def iife26_(_pat_let3_0):
                            def iife27_(d_156_dt__update_hcf31_h0_):
                                return D6_DC17(d_156_dt__update_hcf31_h0_)
                            return iife27_(_pat_let3_0)
                        return iife26_(pat_let_tv7_)
                    return iife25_(_pat_let2_0)
                d_154_v73_ = _dafny.MultiSet([iife24_(d_153_v72_), d_153_v72_, d_153_v72_])
                d_157_v74_: _dafny.Seq
                d_157_v74_ = _dafny.SeqWithoutIsStrInference([d_154_v73_])
                d_158_v75_: _dafny.Seq
                d_158_v75_ = d_157_v74_
                d_159_v76_: _dafny.Set
                d_159_v76_ = _dafny.Set({d_158_v75_})
                d_160_v77_: _dafny.MultiSet
                d_160_v77_ = _dafny.MultiSet([(d_159_v76_).issubset(d_159_v76_)])
                index19_ = default__.safeIndex(910, (d_114_v45_).length(0))
                (d_114_v45_)[index19_] = d_144_v67_.f16
                d_161_v78_: _dafny.Array
                nw25_ = _dafny.Array(_dafny.Seq({}), 14)
                d_161_v78_ = nw25_
                index20_ = default__.safeIndex(207, (d_161_v78_).length(0))
                (d_161_v78_)[index20_] = (d_107_v38_ if d_144_v67_.f16 else d_107_v38_)
                d_162_v79_: _dafny.Map
                d_162_v79_ = _dafny.Map({_dafny.CodePoint('y'): d_144_v67_})
                d_163_v80_: _dafny.Map
                d_163_v80_ = _dafny.Map({len(d_162_v79_): d_143_v66_})
                index21_ = default__.safeIndex(910, (d_114_v45_).length(0))
                index22_ = default__.safeIndex(207, (d_161_v78_).length(0))
                rhs9_ = d_138_v61_.f16
                rhs10_ = d_160_v77_
                rhs11_ = True
                rhs12_ = (d_109_v40_)[default__.safeIndex(d_90_v35_, len(d_109_v40_))]
                rhs13_ = (0) - (default__.safeModuloInt(default__.fm3(globalState), len(d_163_v80_)))
                lhs9_ = d_138_v61_
                lhs10_ = d_114_v45_
                lhs11_ = default__.safeIndex(910, (d_114_v45_).length(0))
                lhs12_ = d_161_v78_
                lhs13_ = default__.safeIndex(207, (d_161_v78_).length(0))
                lhs9_.f16 = rhs9_
                d_160_v77_ = rhs10_
                lhs10_[lhs11_] = rhs11_
                lhs12_[lhs13_] = rhs12_
                d_90_v35_ = rhs13_
                (d_143_v66_).f27 = d_143_v66_.f27
                (globalState).f5 = (d_40_v0_) or ((d_114_v45_)[default__.safeIndex(910, (d_114_v45_).length(0))])
            elif True:
                (d_138_v61_).f16 = not(d_138_v61_.f16)
                d_164_v81_: D12
                d_164_v81_ = D12_DC34(d_114_v45_, p3, d_145_v68_, d_106_v37_)
                d_165_v82_: _dafny.Seq
                d_165_v82_ = _dafny.SeqWithoutIsStrInference([d_143_v66_.f27])
                d_166_v83_: _dafny.Array
                nw26_ = _dafny.Array(None, 10)
                nw26_[int(0)] = d_143_v66_.f27
                nw26_[int(1)] = d_143_v66_.f27
                nw26_[int(2)] = (d_164_v81_).cf59
                nw26_[int(3)] = d_143_v66_.f27
                nw26_[int(4)] = d_143_v66_.f27
                nw26_[int(5)] = d_143_v66_.f27
                nw26_[int(6)] = d_143_v66_.f27
                nw26_[int(7)] = d_143_v66_.f27
                nw26_[int(8)] = (d_165_v82_)[default__.safeIndex((d_135_v58_).f25, len(d_165_v82_))]
                nw26_[int(9)] = d_114_v45_
                d_166_v83_ = nw26_
                d_167_v84_: _dafny.Map
                d_167_v84_ = _dafny.Map({d_106_v37_: d_166_v83_})
                d_167_v84_ = (d_167_v84_).set(d_106_v37_, d_166_v83_)
                d_168_v85_: _dafny.MultiSet
                d_168_v85_ = _dafny.MultiSet([d_138_v61_.f16])
                (globalState).f2 = (d_168_v85_).isdisjoint(d_168_v85_)
                d_169_v86_: C7
                nw27_ = C7()
                nw27_.ctor__(d_145_v68_, d_144_v67_.f16, (d_138_v61_).f17)
                d_169_v86_ = nw27_
                d_169_v86_ = d_169_v86_
                d_170_v87_: C5
                nw28_ = C5()
                nw28_.ctor__(d_144_v67_.f16, (d_138_v61_).f17)
                d_170_v87_ = nw28_
                d_171_v88_: _dafny.Map
                d_171_v88_ = _dafny.Map({d_170_v87_: d_40_v0_})
                d_172_v89_: _dafny.Map
                d_172_v89_ = _dafny.Map({d_138_v61_.f16: d_138_v61_.f16})
                d_173_v90_: D1
                d_173_v90_ = D1_DC4(d_138_v61_.f16, len(d_171_v88_), (d_172_v89_ if d_138_v61_.f16 else d_172_v89_))
                rhs14_ = ((d_107_v38_) == (_dafny.SeqWithoutIsStrInference([d_144_v67_.f16, d_138_v61_.f16])) if d_138_v61_.f16 else d_40_v0_)
                rhs15_ = (d_173_v90_ if (((d_116_v47_).f21).cf2) <= (p3) else d_173_v90_)
                lhs14_ = d_144_v67_
                lhs14_.f16 = rhs14_
                d_173_v90_ = rhs15_
            index23_ = default__.safeIndex(640, (d_114_v45_).length(0))
            (d_114_v45_)[index23_] = d_144_v67_.f16
            index24_ = default__.safeIndex(640, (d_114_v45_).length(0))
            (d_114_v45_)[index24_] = d_40_v0_
            (globalState).f2 = not(d_40_v0_)
        elif True:
            d_174_v91_: _dafny.MultiSet
            d_174_v91_ = _dafny.MultiSet([d_40_v0_, d_40_v0_, d_40_v0_])
            d_175_v92_: _dafny.Map
            d_175_v92_ = _dafny.Map({d_40_v0_: len(_dafny.Map({len(d_111_v42_): d_40_v0_}))})
            d_176_v93_: _dafny.MultiSet
            d_176_v93_ = _dafny.MultiSet([d_174_v91_, default__.fm36(len(d_175_v92_), d_40_v0_, globalState), _dafny.MultiSet([d_40_v0_])])
            (globalState).f2 = ((d_174_v91_).set(d_40_v0_, default__.abs(d_90_v35_))) not in ((d_176_v93_).intersection(_dafny.MultiSet([_dafny.MultiSet([d_40_v0_])])))
            d_107_v38_ = default__.fm20(d_40_v0_, d_40_v0_, not(d_40_v0_), default__.safeModuloInt(577, p2), globalState)
            (globalState).f5 = (d_40_v0_) == (d_40_v0_)
            d_177_v94_: _dafny.Array
            nw29_ = _dafny.Array(_dafny.CodePoint('D'), 17)
            d_177_v94_ = nw29_
            index25_ = default__.safeIndex(336, (d_177_v94_).length(0))
            (d_177_v94_)[index25_] = d_106_v37_
            index26_ = default__.safeIndex(336, (d_177_v94_).length(0))
            (d_177_v94_)[index26_] = d_106_v37_
            d_40_v0_ = (d_107_v38_)[default__.safeIndex((((d_110_v41_)[len(p1)] if (len(p1)) in (d_110_v41_) else 487) if d_40_v0_ else p0), len(d_107_v38_))]
        if ((d_40_v0_) in (d_107_v38_) if (not(d_40_v0_)) and (d_40_v0_) else not((p1) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykfd"))))):
            (globalState).f5 = (p2) != (p0)
            d_90_v35_ = d_90_v35_
            d_178_v95_: _dafny.Seq
            d_178_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwqu"))
            d_178_v95_ = d_178_v95_
            d_179_v96_: C1
            nw30_ = C1()
            nw30_.ctor__(527, (p3 if d_40_v0_ else p0), False, d_111_v42_, p0)
            d_179_v96_ = nw30_
            d_180_v97_: _dafny.MultiSet
            d_180_v97_ = _dafny.MultiSet([d_106_v37_, d_106_v37_])
            rhs16_ = default__.safeModuloInt(((d_180_v97_ if d_40_v0_ else (default__.fm40(p3, d_40_v0_, d_132_v55_, 582, globalState)).set(d_106_v37_, default__.abs(p0)))).cardinality, default__.safeModuloInt(p2, p2))
            rhs17_ = d_179_v96_
            d_90_v35_ = rhs16_
            d_179_v96_ = rhs17_
            d_181_v98_: _dafny.MultiSet
            d_181_v98_ = _dafny.MultiSet([d_107_v38_])
            index27_ = default__.safeIndex(94, (d_114_v45_).length(0))
            (d_114_v45_)[index27_] = (_dafny.MultiSet([d_107_v38_])).ispropersubset(d_181_v98_)
            d_182_v99_: _dafny.Set
            d_182_v99_ = _dafny.Set({(d_112_v43_).fm11(d_107_v38_, d_40_v0_, (0) - (-325), d_40_v0_, globalState), (d_112_v43_).f24, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgwwcrgf"))})
            d_183_v100_: _dafny.MultiSet
            d_183_v100_ = _dafny.MultiSet([p0, p3])
            index28_ = default__.safeIndex(208, (d_114_v45_).length(0))
            (d_114_v45_)[index28_] = (d_183_v100_).issubset(d_183_v100_)
            d_184_v101_: C8
            nw31_ = C8()
            nw31_.ctor__(d_178_v95_, d_132_v55_, d_40_v0_, (d_111_v42_).set(default__.safeIndex(p2, len(d_111_v42_)), (d_179_v96_).f30))
            d_184_v101_ = nw31_
            d_185_v102_: _dafny.Set
            d_185_v102_ = _dafny.Set({d_184_v101_, d_184_v101_})
            d_186_v103_: _dafny.Map
            d_186_v103_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_106_v37_ for d_187_i18_ in range(default__.abs(-662))]): len(_dafny.SeqWithoutIsStrInference([not(d_40_v0_), False, d_40_v0_]))})
            index29_ = default__.safeIndex(94, (d_114_v45_).length(0))
            index30_ = default__.safeIndex(208, (d_114_v45_).length(0))
            def iife28_():
                coll20_ = _dafny.Set()
                compr_20_: _dafny.Seq
                for compr_20_ in (d_186_v103_).keys.Elements:
                    d_188_v104_: _dafny.Seq = compr_20_
                    if (d_188_v104_) in (d_186_v103_):
                        coll20_ = coll20_.union(_dafny.Set([d_188_v104_]))
                return _dafny.Set(coll20_)
            rhs18_ = (d_185_v102_).isdisjoint(d_185_v102_)
            rhs19_ = iife28_()

            rhs20_ = not(d_40_v0_)
            rhs21_ = default__.safeDivisionInt(p0, (0) - (p2))
            lhs15_ = d_114_v45_
            lhs16_ = default__.safeIndex(94, (d_114_v45_).length(0))
            lhs17_ = d_114_v45_
            lhs18_ = default__.safeIndex(208, (d_114_v45_).length(0))
            lhs15_[lhs16_] = rhs18_
            d_182_v99_ = rhs19_
            lhs17_[lhs18_] = rhs20_
            d_90_v35_ = rhs21_
        elif True:
            d_189_v105_: T1
            nw32_ = C1()
            nw32_.ctor__(d_90_v35_, p3, d_40_v0_, d_111_v42_, p2)
            d_189_v105_ = nw32_
            d_190_v106_: _dafny.Map
            d_190_v106_ = _dafny.Map({d_189_v105_: d_114_v45_})
            d_191_v107_: _dafny.Map
            d_191_v107_ = _dafny.Map({(d_114_v45_ if d_40_v0_ else ((d_190_v106_)[d_189_v105_] if (d_189_v105_) in (d_190_v106_) else d_114_v45_)): d_189_v105_.f16})
            d_191_v107_ = (d_191_v107_) | (d_191_v107_)
            d_192_v108_: _dafny.Array
            def lambda11_(d_193_v105_):
                def lambda12_(d_194_i19_):
                    return D1_DC5((d_193_v105_).f28, d_193_v105_.f16)

                return lambda12_

            init5_ = lambda11_(d_189_v105_)
            nw33_ = _dafny.Array(None, 13)
            for i0_5_ in range(nw33_.length(0)):
                nw33_[i0_5_] = init5_(i0_5_)
            d_192_v108_ = nw33_
            index31_ = default__.safeIndex(256, (d_192_v108_).length(0))
            (d_192_v108_)[index31_] = default__.fm34(13, 795, d_132_v55_, len(d_107_v38_), globalState)
            d_195_v109_: D1
            d_195_v109_ = D1_DC5((d_189_v105_).f28, True)
            pat_let_tv8_ = d_40_v0_
            pat_let_tv9_ = d_40_v0_
            index32_ = default__.safeIndex(256, (d_192_v108_).length(0))
            def iife29_(_pat_let4_0):
                def iife30_(d_196_dt__update__tmp_h2_):
                    def iife31_(_pat_let5_0):
                        def iife32_(d_197_dt__update_hcf11_h0_):
                            return D1_DC5((d_196_dt__update__tmp_h2_).cf10, d_197_dt__update_hcf11_h0_)
                        return iife32_(_pat_let5_0)
                    return iife31_((True if pat_let_tv8_ else pat_let_tv9_))
                return iife30_(_pat_let4_0)
            rhs22_ = iife29_(d_195_v109_)
            rhs23_ = (p2) > (p0)
            lhs19_ = d_192_v108_
            lhs20_ = default__.safeIndex(256, (d_192_v108_).length(0))
            lhs19_[lhs20_] = rhs22_
            d_40_v0_ = rhs23_
            d_198_v110_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_198_v110_ = nw34_
            index33_ = default__.safeIndex(669, (d_198_v110_).length(0))
            (d_198_v110_)[index33_] = d_106_v37_
            index34_ = default__.safeIndex(669, (d_198_v110_).length(0))
            (d_198_v110_)[index34_] = d_106_v37_
            d_199_v111_: D2
            d_199_v111_ = D2_DC7(d_132_v55_)
            index35_ = default__.safeIndex(57, (d_114_v45_).length(0))
            (d_114_v45_)[index35_] = d_189_v105_.f16
            d_200_v113_: _dafny.MultiSet
            d_200_v113_ = _dafny.MultiSet([(d_198_v110_)[default__.safeIndex(669, (d_198_v110_).length(0))], (d_198_v110_)[default__.safeIndex(669, (d_198_v110_).length(0))], d_106_v37_])
            d_201_v114_: D9
            def iife33_():
                coll21_ = _dafny.Map()
                compr_21_: str
                for compr_21_ in (d_200_v113_).Elements:
                    d_202_v112_: str = compr_21_
                    if (d_202_v112_) in (d_200_v113_):
                        coll21_[d_202_v112_] = True
                return _dafny.Map(coll21_)
            d_201_v114_ = D9_DC25(iife33_()
)
            index36_ = default__.safeIndex(57, (d_114_v45_).length(0))
            rhs24_ = d_199_v111_
            rhs25_ = ((d_112_v43_).f24)[default__.safeIndex(d_90_v35_, len((d_112_v43_).f24))]
            rhs26_ = (True if d_189_v105_.f16 else d_189_v105_.f16)
            rhs27_ = not(d_40_v0_)
            rhs28_ = default__.fm38(d_201_v114_, (d_189_v105_).f28, globalState)
            lhs21_ = globalState
            lhs22_ = d_114_v45_
            lhs23_ = default__.safeIndex(57, (d_114_v45_).length(0))
            lhs24_ = d_189_v105_
            lhs25_ = globalState
            d_199_v111_ = rhs24_
            lhs21_.f15 = rhs25_
            lhs22_[lhs23_] = rhs26_
            lhs24_.f16 = rhs27_
            lhs25_.f8 = rhs28_
            (d_189_v105_).f16 = not (d_40_v0_) or (not(True))

    @staticmethod
    def Main(noArgsParameter__):
        d_203_v0_: _dafny.Seq
        d_203_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dg"))
        d_204_v1_: int
        d_204_v1_ = 879
        d_205_v2_: str
        d_205_v2_ = _dafny.CodePoint('t')
        d_206_v3_: _dafny.Array
        nw35_ = _dafny.Array(None, 12)
        nw35_[int(0)] = d_203_v0_
        nw35_[int(1)] = d_203_v0_
        nw35_[int(2)] = d_203_v0_
        nw35_[int(3)] = d_203_v0_
        nw35_[int(4)] = d_203_v0_
        nw35_[int(5)] = d_203_v0_
        nw35_[int(6)] = d_203_v0_
        nw35_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
        nw35_[int(8)] = (d_203_v0_).set(default__.safeIndex(d_204_v1_, len(d_203_v0_)), d_205_v2_)
        nw35_[int(9)] = d_203_v0_
        nw35_[int(10)] = (d_203_v0_).set(default__.safeIndex(d_204_v1_, len(d_203_v0_)), d_205_v2_)
        nw35_[int(11)] = d_203_v0_
        d_206_v3_ = nw35_
        d_207_v4_: _dafny.Array
        nw36_ = _dafny.Array(_dafny.MultiSet({}), 5)
        d_207_v4_ = nw36_
        d_208_v5_: _dafny.Array
        def lambda13_(d_209_i0_):
            return True

        init6_ = lambda13_
        nw37_ = _dafny.Array(None, 26)
        for i0_6_ in range(nw37_.length(0)):
            nw37_[i0_6_] = init6_(i0_6_)
        d_208_v5_ = nw37_
        d_210_v6_: _dafny.Seq
        d_210_v6_ = _dafny.SeqWithoutIsStrInference([d_208_v5_])
        d_211_v7_: bool
        d_211_v7_ = True
        d_212_v8_: _dafny.Map
        d_212_v8_ = _dafny.Map({d_211_v7_: d_211_v7_})
        d_213_v9_: _dafny.Map
        d_213_v9_ = _dafny.Map({d_204_v1_: ((d_212_v8_)[False] if (False) in (d_212_v8_) else d_211_v7_)})
        d_214_v10_: _dafny.Seq
        d_214_v10_ = _dafny.SeqWithoutIsStrInference([d_211_v7_])
        d_215_v11_: _dafny.Map
        d_215_v11_ = _dafny.Map({d_214_v10_: d_211_v7_})
        d_216_v12_: _dafny.Seq
        d_216_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_204_v1_, len(d_203_v0_), d_204_v1_, d_204_v1_, d_204_v1_])])
        d_217_v13_: _dafny.Seq
        d_217_v13_ = _dafny.SeqWithoutIsStrInference([d_204_v1_, d_204_v1_])
        d_218_globalState_: GlobalState
        nw38_ = GlobalState()
        nw38_.ctor__(d_206_v3_, d_207_v4_, False, 526, (d_210_v6_)[default__.safeIndex(126, len(d_210_v6_))], True, True, 72, d_213_v9_, (d_215_v11_) | (d_215_v11_), d_216_v12_, False, d_217_v13_, _dafny.SeqWithoutIsStrInference([d_205_v2_ for d_219_i1_ in range(default__.abs(156))]), True, _dafny.CodePoint('w'))
        d_218_globalState_ = nw38_
        d_220_v14_: _dafny.Map
        d_220_v14_ = _dafny.Map({not(default__.fm0(d_211_v7_, d_211_v7_, d_218_globalState_)): d_208_v5_})
        d_204_v1_ = len(((d_220_v14_) | (d_220_v14_)).set(d_211_v7_, d_208_v5_))
        d_221_v15_: _dafny.Array
        def lambda14_(d_222_v1_):
            def lambda15_(d_223_i2_):
                return (_dafny.Map({d_222_v1_: -637})) | (_dafny.Map({d_222_v1_: d_222_v1_}))

            return lambda15_

        init7_ = lambda14_(d_204_v1_)
        nw39_ = _dafny.Array(None, 22)
        for i0_7_ in range(nw39_.length(0)):
            nw39_[i0_7_] = init7_(i0_7_)
        d_221_v15_ = nw39_
        d_224_v16_: _dafny.Map
        d_224_v16_ = _dafny.Map({d_204_v1_: d_204_v1_})
        d_225_v17_: _dafny.Map
        d_225_v17_ = _dafny.Map({_dafny.CodePoint('f'): len(d_203_v0_)})
        d_226_v18_: _dafny.Map
        d_226_v18_ = _dafny.Map({((d_224_v16_)[d_204_v1_] if (d_204_v1_) in (d_224_v16_) else len(default__.fm1(d_218_globalState_))): len(d_225_v17_)})
        index37_ = default__.safeIndex(251, (d_221_v15_).length(0))
        (d_221_v15_)[index37_] = d_226_v18_
        index38_ = default__.safeIndex(251, (d_221_v15_).length(0))
        (d_221_v15_)[index38_] = d_226_v18_
        d_204_v1_ = (d_204_v1_) * (d_204_v1_)
        d_227_v19_: _dafny.Array
        nw40_ = _dafny.Array(int(0), 15)
        d_227_v19_ = nw40_
        d_228_v20_: _dafny.Map
        d_228_v20_ = _dafny.Map({d_227_v19_: d_226_v18_})
        d_229_v21_: _dafny.Array
        nw41_ = _dafny.Array(None, 13)
        nw41_[int(0)] = len((d_214_v10_) + (d_214_v10_))
        nw41_[int(1)] = (d_204_v1_) - (948)
        nw41_[int(2)] = (0) - (d_204_v1_)
        nw41_[int(3)] = d_204_v1_
        nw41_[int(4)] = d_204_v1_
        nw41_[int(5)] = (0) - (len((d_228_v20_) | (d_228_v20_)))
        nw41_[int(6)] = len(_dafny.Set({d_211_v7_, d_211_v7_}))
        nw41_[int(7)] = d_204_v1_
        nw41_[int(8)] = d_204_v1_
        nw41_[int(9)] = d_204_v1_
        nw41_[int(10)] = d_204_v1_
        nw41_[int(11)] = d_204_v1_
        nw41_[int(12)] = d_204_v1_
        d_229_v21_ = nw41_
        d_227_v19_ = d_227_v19_
        d_230_i3_: int
        d_230_i3_ = 0
        with _dafny.label("0"):
            while (d_211_v7_) and (d_211_v7_):
                with _dafny.c_label("0"):
                    if (d_230_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_230_i3_ = (d_230_i3_) + (1)
                    d_204_v1_ = d_204_v1_
                    default__.m0(d_204_v1_, (d_203_v0_).set(default__.safeIndex(d_204_v1_, len(d_203_v0_)), d_205_v2_), len(_dafny.SeqWithoutIsStrInference([d_204_v1_, (0) - ((0) - (d_204_v1_))])), (d_204_v1_) - (132), d_218_globalState_)
                    d_231_v22_: _dafny.Map
                    d_231_v22_ = _dafny.Map({d_210_v6_: d_211_v7_})
                    d_231_v22_ = (d_231_v22_).set((d_210_v6_) + (d_210_v6_), (d_214_v10_)[default__.safeIndex(d_204_v1_, len(d_214_v10_))])
                    default__.m0(d_204_v1_, d_203_v0_, (d_204_v1_) + (d_204_v1_), d_204_v1_, d_218_globalState_)
                    pass
            pass
        def iife34_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(327, 229):
                d_232_v23_: int = compr_22_
                if ((327) <= (d_232_v23_)) and ((d_232_v23_) < (229)):
                    coll22_[default__.safeModuloInt(d_232_v23_, d_204_v1_)] = d_211_v7_
            return _dafny.Map(coll22_)
        default__.m0(d_204_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "casqliqs")), default__.safeDivisionInt(d_204_v1_, len(d_203_v0_)), len(iife34_()
        ), d_218_globalState_)
        d_204_v1_ = d_204_v1_
        d_233_i4_: int
        d_233_i4_ = 0
        with _dafny.label("1"):
            while d_211_v7_:
                with _dafny.c_label("1"):
                    if (d_233_i4_) >= (100):
                        raise _dafny.Break("1")
                    d_233_i4_ = (d_233_i4_) + (1)
                    d_234_v24_: _dafny.Map
                    d_234_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqs")): d_204_v1_})
                    d_235_v26_: _dafny.Map
                    d_235_v26_ = _dafny.Map({d_203_v0_: _dafny.Set({d_211_v7_})})
                    d_236_v28_: _dafny.Seq
                    d_236_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_205_v2_ for d_237_i5_ in range(default__.abs(-396))])])
                    d_238_v29_: _dafny.Seq
                    def iife35_():
                        coll23_ = _dafny.Map()
                        compr_23_: _dafny.Seq
                        for compr_23_ in (d_235_v26_).keys.Elements:
                            d_239_v25_: _dafny.Seq = compr_23_
                            if (d_239_v25_) in (d_235_v26_):
                                coll23_[d_239_v25_] = d_204_v1_
                        return _dafny.Map(coll23_)
                    def iife36_():
                        coll24_ = _dafny.Map()
                        compr_24_: _dafny.Seq
                        for compr_24_ in ((d_236_v28_).set(default__.safeIndex((0) - (d_204_v1_), len(d_236_v28_)), d_203_v0_)).Elements:
                            d_240_v27_: _dafny.Seq = compr_24_
                            if (d_240_v27_) in ((d_236_v28_).set(default__.safeIndex((0) - (d_204_v1_), len(d_236_v28_)), d_203_v0_)):
                                coll24_[d_240_v27_] = d_204_v1_
                        return _dafny.Map(coll24_)
                    d_238_v29_ = _dafny.SeqWithoutIsStrInference([iife35_()
                    , _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsb")): d_204_v1_}), (iife36_()
                    ).set(d_203_v0_, len(d_203_v0_))])
                    d_234_v24_ = ((d_238_v29_)[default__.safeIndex(d_204_v1_, len(d_238_v29_))]) | (d_234_v24_)
                    nw42_ = _dafny.Array(False, 22)
                    d_208_v5_ = nw42_
                    d_204_v1_ = (0) - ((0) - (d_204_v1_))
                    (d_218_globalState_).f2 = (len(d_226_v18_)) <= (d_204_v1_)
                    pass
            pass
        index39_ = default__.safeIndex(595, (d_206_v3_).length(0))
        (d_206_v3_)[index39_] = d_203_v0_
        d_241_v30_: _dafny.Map
        d_241_v30_ = _dafny.Map({d_211_v7_: d_205_v2_})
        d_242_v31_: _dafny.Map
        d_242_v31_ = _dafny.Map({not(d_211_v7_): d_241_v30_})
        index40_ = default__.safeIndex(405, (d_227_v19_).length(0))
        (d_227_v19_)[index40_] = (len(d_242_v31_) if ((d_212_v8_)[d_211_v7_] if (d_211_v7_) in (d_212_v8_) else d_211_v7_) else d_204_v1_)
        d_243_v32_: _dafny.MultiSet
        d_243_v32_ = _dafny.MultiSet([d_205_v2_, d_205_v2_, default__.fm2(d_218_globalState_)])
        index41_ = default__.safeIndex(595, (d_206_v3_).length(0))
        index42_ = default__.safeIndex(405, (d_227_v19_).length(0))
        rhs29_ = (D0_DC0(default__.fm1(d_218_globalState_))).cf0
        rhs30_ = (d_243_v32_).cardinality
        rhs31_ = d_208_v5_
        lhs26_ = d_206_v3_
        lhs27_ = default__.safeIndex(595, (d_206_v3_).length(0))
        lhs28_ = d_227_v19_
        lhs29_ = default__.safeIndex(405, (d_227_v19_).length(0))
        lhs26_[lhs27_] = rhs29_
        lhs28_[lhs29_] = rhs30_
        d_208_v5_ = rhs31_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_208_v5_).length(0)):
            d_244_i6_: int = guard_loop_0_
            if (True) and (((0) <= (d_244_i6_)) and ((d_244_i6_) < ((d_208_v5_).length(0)))):
                (d_208_v5_)[(d_244_i6_)] = (_dafny.MultiSet([d_211_v7_])).isdisjoint((_dafny.MultiSet([d_211_v7_])).intersection(_dafny.MultiSet([d_211_v7_])))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_208_v5_).length(0)):
            d_245_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_245_i7_)) and ((d_245_i7_) < ((d_208_v5_).length(0)))):
                (d_208_v5_)[(d_245_i7_)] = not(d_211_v7_)
        if d_211_v7_:
            d_246_v34_: _dafny.Map
            def iife37_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(333, 468):
                    d_247_v33_: int = compr_25_
                    if ((333) <= (d_247_v33_)) and ((d_247_v33_) < (468)):
                        coll25_[default__.safeDivisionInt(d_247_v33_, d_204_v1_)] = d_211_v7_
                return _dafny.Map(coll25_)
            d_246_v34_ = _dafny.Map({d_211_v7_: iife37_()
            })
            d_204_v1_ = (len((d_246_v34_) | (d_246_v34_)) if d_211_v7_ else default__.fm3(d_218_globalState_))
            d_248_v35_: D0
            d_248_v35_ = D0_DC1(default__.fm1(d_218_globalState_))
            source2_ = d_248_v35_
            if source2_.is_DC1:
                d_249___mcc_h0_ = source2_.cf1
                d_250_cf1_ = d_249___mcc_h0_
                d_251_v36_: _dafny.MultiSet
                d_251_v36_ = _dafny.MultiSet([d_211_v7_, d_211_v7_, d_211_v7_])
                index43_ = default__.safeIndex(310, (d_207_v4_).length(0))
                (d_207_v4_)[index43_] = d_251_v36_
                index44_ = default__.safeIndex(635, (d_208_v5_).length(0))
                (d_208_v5_)[index44_] = d_211_v7_
                d_252_v37_: _dafny.Seq
                d_252_v37_ = _dafny.SeqWithoutIsStrInference([d_251_v36_])
                d_253_v38_: D1
                d_253_v38_ = D1_DC3((d_252_v37_)[default__.safeIndex(d_204_v1_, len(d_252_v37_))])
                d_254_v39_: D1
                d_254_v39_ = D1_DC4(d_211_v7_, d_204_v1_, d_212_v8_)
                pat_let_tv10_ = d_251_v36_
                index45_ = default__.safeIndex(310, (d_207_v4_).length(0))
                index46_ = default__.safeIndex(635, (d_208_v5_).length(0))
                def iife38_(_pat_let6_0):
                    def iife39_(d_255_dt__update__tmp_h0_):
                        def iife40_(_pat_let7_0):
                            def iife41_(d_256_dt__update_hcf6_h0_):
                                return D1_DC3(d_256_dt__update_hcf6_h0_)
                            return iife41_(_pat_let7_0)
                        return iife40_(pat_let_tv10_)
                    return iife39_(_pat_let6_0)
                rhs32_ = (iife38_(d_253_v38_)).cf6
                rhs33_ = (d_254_v39_).cf7
                rhs34_ = not(d_211_v7_)
                rhs35_ = not(not(not(d_211_v7_)))
                lhs30_ = d_207_v4_
                lhs31_ = default__.safeIndex(310, (d_207_v4_).length(0))
                lhs32_ = d_218_globalState_
                lhs33_ = d_208_v5_
                lhs34_ = default__.safeIndex(635, (d_208_v5_).length(0))
                lhs35_ = d_218_globalState_
                lhs30_[lhs31_] = rhs32_
                lhs32_.f2 = rhs33_
                lhs33_[lhs34_] = rhs34_
                lhs35_.f5 = rhs35_
                d_204_v1_ = (len(d_203_v0_)) - (d_204_v1_)
                d_257_v41_: _dafny.Seq
                def iife42_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(-26, 365):
                        d_258_v40_: int = compr_26_
                        if ((-26) <= (d_258_v40_)) and ((d_258_v40_) < (365)):
                            coll26_[(d_258_v40_) * (d_204_v1_)] = False
                    return _dafny.Map(coll26_)
                d_257_v41_ = _dafny.SeqWithoutIsStrInference([iife42_()
                , (d_213_v9_) | (d_213_v9_)])
                d_257_v41_ = ((d_257_v41_).set(default__.safeIndex(d_204_v1_, len(d_257_v41_)), d_213_v9_) if (d_208_v5_)[default__.safeIndex(635, (d_208_v5_).length(0))] else (d_257_v41_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_204_v1_: (d_208_v5_)[default__.safeIndex(635, (d_208_v5_).length(0))]}), d_213_v9_, d_213_v9_])))
                d_213_v9_ = (d_213_v9_).set((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], (d_250_cf1_) != (d_203_v0_))
            elif source2_.is_DC2:
                d_259___mcc_h1_ = source2_.cf2
                d_260___mcc_h2_ = source2_.cf3
                d_261___mcc_h3_ = source2_.cf4
                d_262___mcc_h4_ = source2_.cf5
                d_263_cf5_ = d_262___mcc_h4_
                d_264_cf4_ = d_261___mcc_h3_
                d_265_cf3_ = d_260___mcc_h2_
                d_266_cf2_ = d_259___mcc_h1_
                d_263_cf5_ = d_208_v5_
                default__.m0(d_264_cf4_, (d_203_v0_) + ((d_206_v3_)[default__.safeIndex(595, (d_206_v3_).length(0))]), d_266_cf2_, (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], d_218_globalState_)
                d_267_v42_: _dafny.MultiSet
                d_267_v42_ = _dafny.MultiSet([d_211_v7_])
                d_268_v43_: _dafny.Seq
                d_268_v43_ = _dafny.SeqWithoutIsStrInference([D1_DC3(d_267_v42_)])
                index47_ = default__.safeIndex(717, (d_208_v5_).length(0))
                (d_208_v5_)[index47_] = (default__.fm4(d_211_v7_, -367, d_218_globalState_)) != (d_268_v43_)
                index48_ = default__.safeIndex(405, (d_227_v19_).length(0))
                index49_ = default__.safeIndex(717, (d_208_v5_).length(0))
                index50_ = default__.safeIndex(405, (d_227_v19_).length(0))
                rhs36_ = d_264_cf4_
                rhs37_ = d_211_v7_
                rhs38_ = d_211_v7_
                rhs39_ = d_204_v1_
                rhs40_ = d_265_cf3_
                lhs36_ = d_227_v19_
                lhs37_ = default__.safeIndex(405, (d_227_v19_).length(0))
                lhs38_ = d_208_v5_
                lhs39_ = default__.safeIndex(717, (d_208_v5_).length(0))
                lhs40_ = d_218_globalState_
                lhs41_ = d_227_v19_
                lhs42_ = default__.safeIndex(405, (d_227_v19_).length(0))
                lhs36_[lhs37_] = rhs36_
                lhs38_[lhs39_] = rhs37_
                lhs40_.f5 = rhs38_
                d_265_cf3_ = rhs39_
                lhs41_[lhs42_] = rhs40_
                (d_218_globalState_).f2 = (d_208_v5_)[default__.safeIndex(717, (d_208_v5_).length(0))]
            elif True:
                d_269___mcc_h5_ = source2_.cf0
                d_270_cf0_ = d_269___mcc_h5_
                d_271_v44_: _dafny.MultiSet
                d_271_v44_ = _dafny.MultiSet([d_211_v7_, False])
                index51_ = default__.safeIndex(245, (d_207_v4_).length(0))
                (d_207_v4_)[index51_] = (d_271_v44_).set(d_211_v7_, default__.abs(d_204_v1_))
                index52_ = default__.safeIndex(245, (d_207_v4_).length(0))
                rhs41_ = d_271_v44_
                rhs42_ = d_211_v7_
                lhs43_ = d_207_v4_
                lhs44_ = default__.safeIndex(245, (d_207_v4_).length(0))
                lhs45_ = d_218_globalState_
                lhs43_[lhs44_] = rhs41_
                lhs45_.f5 = rhs42_
                default__.m0((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxynt")), (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], default__.safeModuloInt((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], 255), d_218_globalState_)
                d_272_v45_: _dafny.MultiSet
                d_272_v45_ = _dafny.MultiSet([((d_207_v4_)[default__.safeIndex(245, (d_207_v4_).length(0))]).cardinality])
                default__.m0(d_204_v1_, (d_206_v3_)[default__.safeIndex(595, (d_206_v3_).length(0))], (d_272_v45_).cardinality, (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], d_218_globalState_)
                index53_ = default__.safeIndex(765, (d_208_v5_).length(0))
                (d_208_v5_)[index53_] = d_211_v7_
                index54_ = default__.safeIndex(765, (d_208_v5_).length(0))
                (d_208_v5_)[index54_] = (d_270_cf0_) == (_dafny.SeqWithoutIsStrInference([d_205_v2_ for d_273_i8_ in range(default__.abs(-422))]))
            index55_ = default__.safeIndex(405, (d_227_v19_).length(0))
            (d_227_v19_)[index55_] = d_204_v1_
            d_274_v46_: _dafny.Array
            nw43_ = _dafny.Array(_dafny.Seq({}), 22)
            d_274_v46_ = nw43_
            d_275_v47_: _dafny.Seq
            d_275_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_211_v7_, d_211_v7_, d_211_v7_, d_211_v7_, d_211_v7_]), d_214_v10_, d_214_v10_, (d_214_v10_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_211_v7_])), len(d_214_v10_)), d_211_v7_), _dafny.SeqWithoutIsStrInference([False])])
            index56_ = default__.safeIndex(908, (d_274_v46_).length(0))
            (d_274_v46_)[index56_] = d_275_v47_
            index57_ = default__.safeIndex(908, (d_274_v46_).length(0))
            index58_ = default__.safeIndex(405, (d_227_v19_).length(0))
            rhs43_ = d_275_v47_
            rhs44_ = len((d_217_v13_) + (d_217_v13_))
            rhs45_ = d_211_v7_
            lhs46_ = d_274_v46_
            lhs47_ = default__.safeIndex(908, (d_274_v46_).length(0))
            lhs48_ = d_227_v19_
            lhs49_ = default__.safeIndex(405, (d_227_v19_).length(0))
            lhs50_ = d_218_globalState_
            lhs46_[lhs47_] = rhs43_
            lhs48_[lhs49_] = rhs44_
            lhs50_.f5 = rhs45_
            d_276_v48_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_276_v48_ = nw44_
            index59_ = default__.safeIndex(23, (d_276_v48_).length(0))
            (d_276_v48_)[index59_] = d_208_v5_
            index60_ = default__.safeIndex(23, (d_276_v48_).length(0))
            (d_276_v48_)[index60_] = d_208_v5_
        elif True:
            d_277_v49_: C7
            nw45_ = C7()
            nw45_.ctor__(d_229_v21_, (not(d_211_v7_)) == (d_211_v7_), (_dafny.SeqWithoutIsStrInference([(d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], d_204_v1_])).set(default__.safeIndex(len(d_203_v0_), len(_dafny.SeqWithoutIsStrInference([(d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], d_204_v1_]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xse")))))
            d_277_v49_ = nw45_
            d_278_v50_: C5
            nw46_ = C5()
            nw46_.ctor__(d_211_v7_, d_217_v13_)
            d_278_v50_ = nw46_
            d_204_v1_ = -693
            index61_ = default__.safeIndex(405, (d_227_v19_).length(0))
            rhs46_ = not(d_211_v7_)
            rhs47_ = (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]
            lhs51_ = d_218_globalState_
            lhs52_ = d_227_v19_
            lhs53_ = default__.safeIndex(405, (d_227_v19_).length(0))
            lhs51_.f5 = rhs46_
            lhs52_[lhs53_] = rhs47_
            d_279_v51_: _dafny.Array
            nw47_ = _dafny.Array(_dafny.Map({}), 12)
            d_279_v51_ = nw47_
            index62_ = default__.safeIndex(405, (d_227_v19_).length(0))
            rhs48_ = (d_279_v51_ if default__.fm0(d_211_v7_, d_211_v7_, d_218_globalState_) else d_279_v51_)
            rhs49_ = ((_dafny.SeqWithoutIsStrInference([d_211_v7_, d_211_v7_, d_211_v7_, d_211_v7_, d_211_v7_])) + (_dafny.SeqWithoutIsStrInference([d_211_v7_]))) != (d_214_v10_)
            rhs50_ = default__.fm3(d_218_globalState_)
            lhs54_ = d_227_v19_
            lhs55_ = default__.safeIndex(405, (d_227_v19_).length(0))
            d_279_v51_ = rhs48_
            d_211_v7_ = rhs49_
            lhs54_[lhs55_] = rhs50_
        index63_ = default__.safeIndex(405, (d_227_v19_).length(0))
        (d_227_v19_)[index63_] = (0) - (len(d_214_v10_))
        d_280_v52_: _dafny.MultiSet
        d_280_v52_ = _dafny.MultiSet([d_211_v7_])
        d_281_v53_: C7
        nw48_ = C7()
        nw48_.ctor__(d_229_v21_, ((d_280_v52_).set(d_211_v7_, default__.abs(d_204_v1_))) == (d_280_v52_), _dafny.SeqWithoutIsStrInference([(d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))] for d_282_i9_ in range(default__.abs(384))]))
        d_281_v53_ = nw48_
        d_283_v54_: D3
        d_283_v54_ = D3_DC10(d_211_v7_)
        d_283_v54_ = d_283_v54_
        d_284_v55_: _dafny.Set
        d_284_v55_ = _dafny.Set({d_211_v7_, not(d_211_v7_)})
        hi1_ = ((d_280_v52_)[d_211_v7_] if (d_211_v7_) in (d_280_v52_) else len(d_284_v55_))
        for d_285_i10_ in range(d_204_v1_, hi1_):
            if d_211_v7_:
                d_286_v56_: _dafny.Array
                nw49_ = _dafny.Array(D12.default()(), 13)
                d_286_v56_ = nw49_
                d_287_v57_: C8
                nw50_ = C8()
                nw50_.ctor__(d_203_v0_, d_284_v55_, d_211_v7_, _dafny.SeqWithoutIsStrInference([529, d_204_v1_]))
                d_287_v57_ = nw50_
                d_288_v58_: _dafny.Map
                d_288_v58_ = _dafny.Map({d_287_v57_: d_211_v7_})
                d_289_v59_: _dafny.Map
                d_289_v59_ = _dafny.Map({d_288_v58_: d_286_v56_})
                d_290_v60_: _dafny.Array
                nw51_ = _dafny.Array(None, 8)
                nw51_[int(0)] = d_286_v56_
                nw51_[int(1)] = d_286_v56_
                nw51_[int(2)] = d_286_v56_
                nw51_[int(3)] = ((d_289_v59_)[_dafny.Map({d_287_v57_: d_211_v7_})] if (_dafny.Map({d_287_v57_: d_211_v7_})) in (d_289_v59_) else d_286_v56_)
                nw51_[int(4)] = d_286_v56_
                nw51_[int(5)] = d_286_v56_
                nw51_[int(6)] = d_286_v56_
                nw51_[int(7)] = d_286_v56_
                d_290_v60_ = nw51_
                index64_ = default__.safeIndex(448, (d_290_v60_).length(0))
                (d_290_v60_)[index64_] = d_286_v56_
                d_291_v61_: _dafny.Map
                d_291_v61_ = _dafny.Map({d_211_v7_: (0) - ((0) - ((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]))})
                d_292_v62_: _dafny.Map
                d_292_v62_ = _dafny.Map({(d_204_v1_) * (d_204_v1_): (d_214_v10_) + (d_214_v10_)})
                d_293_v63_: _dafny.MultiSet
                d_293_v63_ = _dafny.MultiSet([_dafny.MultiSet([True]), d_280_v52_, d_280_v52_])
                d_294_v64_: T1
                nw52_ = C1()
                nw52_.ctor__((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], d_204_v1_, default__.fm0(d_211_v7_, d_211_v7_, d_218_globalState_), d_217_v13_, d_204_v1_)
                d_294_v64_ = nw52_
                d_295_v65_: T0
                nw53_ = C1()
                nw53_.ctor__(len(d_291_v61_), (_dafny.MultiSet([d_294_v64_, d_294_v64_])).cardinality, d_211_v7_, (d_294_v64_).f17, (d_294_v64_).f28)
                d_295_v65_ = nw53_
                d_296_v66_: _dafny.Map
                d_296_v66_ = _dafny.Map({(d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]: d_295_v65_})
                d_297_v67_: C3
                nw54_ = C3()
                nw54_.ctor__(d_203_v0_, (D1_DC5(len(d_287_v57_.f19), d_295_v65_.f16)).cf11, (d_295_v65_).f17)
                d_297_v67_ = nw54_
                d_298_v68_: _dafny.Map
                d_298_v68_ = _dafny.Map({d_297_v67_: len(default__.fm31(d_294_v64_.f16, d_218_globalState_))})
                index65_ = default__.safeIndex(448, (d_290_v60_).length(0))
                index66_ = default__.safeIndex(405, (d_227_v19_).length(0))
                index67_ = default__.safeIndex(405, (d_227_v19_).length(0))
                rhs51_ = d_286_v56_
                rhs52_ = _dafny.Map({(d_280_v52_) in (d_293_v63_): (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]})
                rhs53_ = (d_285_i10_) - (((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))] if d_211_v7_ else len(d_296_v66_)))
                rhs54_ = d_292_v62_
                rhs55_ = (default__.safeDivisionInt(len(d_217_v13_), ((d_280_v52_).set(False, default__.abs(((d_298_v68_)[d_297_v67_] if (d_297_v67_) in (d_298_v68_) else (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))])))).cardinality)) * (default__.safeDivisionInt(d_285_i10_, d_285_i10_))
                lhs56_ = d_290_v60_
                lhs57_ = default__.safeIndex(448, (d_290_v60_).length(0))
                lhs58_ = d_227_v19_
                lhs59_ = default__.safeIndex(405, (d_227_v19_).length(0))
                lhs60_ = d_227_v19_
                lhs61_ = default__.safeIndex(405, (d_227_v19_).length(0))
                lhs56_[lhs57_] = rhs51_
                d_291_v61_ = rhs52_
                lhs58_[lhs59_] = rhs53_
                d_292_v62_ = rhs54_
                lhs60_[lhs61_] = rhs55_
                d_299_v69_: _dafny.Map
                d_299_v69_ = _dafny.Map({_dafny.Map({d_205_v2_: d_281_v53_}): (0) - ((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))])})
                d_300_v70_: _dafny.Map
                d_300_v70_ = _dafny.Map({d_285_i10_: d_281_v53_})
                rhs56_ = d_299_v69_
                rhs57_ = ((d_300_v70_) | (d_300_v70_)) != ((d_300_v70_) | (d_300_v70_))
                d_299_v69_ = rhs56_
                d_211_v7_ = rhs57_
                d_301_v71_: _dafny.Seq
                d_301_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_212_v8_) for d_302_i11_ in range(default__.abs(601))]), (d_295_v65_).f17, (d_295_v65_).f17, (d_295_v65_).f17, (d_294_v64_).f17])
                d_303_v72_: C1
                nw55_ = C1()
                nw55_.ctor__(d_204_v1_, d_204_v1_, False, ((d_301_v71_)[default__.safeIndex((d_294_v64_).f28, len(d_301_v71_))]) + (_dafny.SeqWithoutIsStrInference([d_204_v1_, 427])), d_285_i10_)
                d_303_v72_ = nw55_
                d_304_v75_: _dafny.Map
                d_304_v75_ = _dafny.Map({d_284_v55_: (default__.fm26(d_218_globalState_)).cf35})
                d_305_v76_: _dafny.Set
                def iife43_():
                    def iife45_():
                        coll29_ = _dafny.Map()
                        compr_29_: _dafny.Set
                        for compr_29_ in (d_304_v75_).keys.Elements:
                            d_306_v74_: _dafny.Set = compr_29_
                            if (d_306_v74_) in (d_304_v75_):
                                coll29_[d_306_v74_] = d_205_v2_
                        return _dafny.Map(coll29_)
                    coll27_ = _dafny.Map()
                    def iife44_():
                        coll28_ = _dafny.Map()
                        compr_28_: _dafny.Set
                        for compr_28_ in (d_304_v75_).keys.Elements:
                            d_306_v74_: _dafny.Set = compr_28_
                            if (d_306_v74_) in (d_304_v75_):
                                coll28_[d_306_v74_] = d_205_v2_
                        return _dafny.Map(coll28_)
                    compr_27_: _dafny.Set
                    for compr_27_ in (iife44_()
                    ).keys.Elements:
                        d_307_v73_: _dafny.Set = compr_27_
                        if (d_307_v73_) in (iife45_()
                        ):
                            coll27_[d_307_v73_] = d_294_v64_.f16
                    return _dafny.Map(coll27_)
                d_305_v76_ = _dafny.Set({len(iife43_()
                ), -809})
                d_308_v77_: _dafny.Set
                d_308_v77_ = _dafny.Set({d_285_i10_, len(d_305_v76_), (d_303_v72_).f30, default__.safeModuloInt((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))))})
                def iife46_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(-708, 256):
                        d_309_v78_: int = compr_30_
                        if ((-708) <= (d_309_v78_)) and ((d_309_v78_) < (256)):
                            coll30_ = coll30_.union(_dafny.Set([(d_309_v78_) - ((d_303_v72_).f30)]))
                    return _dafny.Set(coll30_)
                d_305_v76_ = iife46_()
                
                d_310_v79_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Seq({}), 18)
                d_310_v79_ = nw56_
                d_311_v81_: _dafny.Seq
                d_311_v81_ = _dafny.SeqWithoutIsStrInference([d_214_v10_])
                index68_ = default__.safeIndex(405, (d_227_v19_).length(0))
                index69_ = default__.safeIndex(405, (d_227_v19_).length(0))
                def iife47_():
                    coll31_ = _dafny.Set()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(76, -151):
                        d_312_v80_: int = compr_31_
                        if ((76) <= (d_312_v80_)) and ((d_312_v80_) < (-151)):
                            coll31_ = coll31_.union(_dafny.Set([(d_312_v80_) * ((d_303_v72_).f30)]))
                    return _dafny.Set(coll31_)
                rhs58_ = (d_308_v77_).intersection(iife47_()
                )
                rhs59_ = ((d_303_v72_).f30) + ((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))])
                rhs60_ = ((d_311_v81_)[default__.safeIndex(len(d_214_v10_), len(d_311_v81_))]) + (d_214_v10_)
                rhs61_ = d_310_v79_
                rhs62_ = d_285_i10_
                lhs62_ = d_227_v19_
                lhs63_ = default__.safeIndex(405, (d_227_v19_).length(0))
                lhs64_ = d_227_v19_
                lhs65_ = default__.safeIndex(405, (d_227_v19_).length(0))
                d_305_v76_ = rhs58_
                lhs62_[lhs63_] = rhs59_
                d_214_v10_ = rhs60_
                d_310_v79_ = rhs61_
                lhs64_[lhs65_] = rhs62_
            elif True:
                d_206_v3_ = d_206_v3_
                d_313_v82_: C3
                nw57_ = C3()
                nw57_.ctor__(((d_206_v3_)[default__.safeIndex(595, (d_206_v3_).length(0))]) + (default__.fm1(d_218_globalState_)), d_211_v7_, d_217_v13_)
                d_313_v82_ = nw57_
                (d_218_globalState_).f5 = d_211_v7_
                d_314_v83_: C0
                nw58_ = C0()
                nw58_.ctor__((0) - ((0) - (d_204_v1_)), d_208_v5_)
                d_314_v83_ = nw58_
                d_315_v84_: C3
                nw59_ = C3()
                nw59_.ctor__((d_313_v82_).f24, d_211_v7_, ((d_217_v13_).set(default__.safeIndex((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], len(d_217_v13_)), d_204_v1_)) + (d_217_v13_))
                d_315_v84_ = nw59_
                d_315_v84_ = d_313_v82_
            d_211_v7_ = True
            nw60_ = _dafny.Array(None, 2)
            nw60_[int(0)] = True
            nw60_[int(1)] = d_211_v7_
            d_208_v5_ = nw60_
            if d_211_v7_:
                d_316_v85_: T0
                nw61_ = C5()
                nw61_.ctor__(d_211_v7_, d_217_v13_)
                d_316_v85_ = nw61_
                d_317_v86_: _dafny.Map
                d_317_v86_ = _dafny.Map({d_316_v85_: d_285_i10_})
                d_318_v87_: _dafny.MultiSet
                d_318_v87_ = _dafny.MultiSet([d_317_v86_])
                d_319_v88_: _dafny.Map
                d_319_v88_ = _dafny.Map({d_211_v7_: ((d_318_v87_)[_dafny.Map({d_316_v85_: d_204_v1_})] if (_dafny.Map({d_316_v85_: d_204_v1_})) in (d_318_v87_) else d_285_i10_)})
                d_319_v88_ = (d_319_v88_).set(False, d_204_v1_)
                d_320_v89_: T1
                nw62_ = C1()
                nw62_.ctor__(55, len((d_206_v3_)[default__.safeIndex(595, (d_206_v3_).length(0))]), d_211_v7_, (d_316_v85_).f17, (0) - ((d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]))
                d_320_v89_ = nw62_
                d_321_v90_: _dafny.Map
                d_321_v90_ = _dafny.Map({default__.fm2(d_218_globalState_): d_320_v89_})
                rhs63_ = (default__.fm0(d_211_v7_, d_316_v85_.f16, d_218_globalState_) if (not(d_320_v89_.f16) if d_316_v85_.f16 else d_211_v7_) else d_316_v85_.f16)
                rhs64_ = ((d_321_v90_)[_dafny.CodePoint('v')] if (_dafny.CodePoint('v')) in (d_321_v90_) else d_320_v89_)
                rhs65_ = d_280_v52_
                lhs66_ = d_218_globalState_
                lhs66_.f2 = rhs63_
                d_320_v89_ = rhs64_
                d_280_v52_ = rhs65_
                d_322_v91_: D11
                d_322_v91_ = D11_DC29(d_205_v2_)
                d_322_v91_ = d_322_v91_
                d_323_v92_: _dafny.Set
                d_323_v92_ = _dafny.Set({(d_320_v89_).f28, (d_320_v89_).f28})
                index70_ = default__.safeIndex(405, (d_227_v19_).length(0))
                (d_227_v19_)[index70_] = len((d_323_v92_).intersection(d_323_v92_))
                index71_ = default__.safeIndex(405, (d_227_v19_).length(0))
                (d_227_v19_)[index71_] = default__.fm3(d_218_globalState_)
            elif True:
                index72_ = default__.safeIndex(595, (d_206_v3_).length(0))
                (d_206_v3_)[index72_] = (d_206_v3_)[default__.safeIndex(595, (d_206_v3_).length(0))]
                d_212_v8_ = (d_212_v8_).set(d_211_v7_, not(d_211_v7_))
                index73_ = default__.safeIndex(405, (d_227_v19_).length(0))
                rhs66_ = (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]
                rhs67_ = d_211_v7_
                rhs68_ = (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]
                rhs69_ = d_285_i10_
                lhs67_ = d_218_globalState_
                lhs68_ = d_227_v19_
                lhs69_ = default__.safeIndex(405, (d_227_v19_).length(0))
                d_204_v1_ = rhs66_
                lhs67_.f5 = rhs67_
                d_204_v1_ = rhs68_
                lhs68_[lhs69_] = rhs69_
                d_324_v93_: C4
                nw63_ = C4()
                nw63_.ctor__(578, default__.fm0(((d_212_v8_)[d_211_v7_] if (d_211_v7_) in (d_212_v8_) else d_211_v7_), d_211_v7_, d_218_globalState_))
                d_324_v93_ = nw63_
                d_325_v94_: _dafny.Seq
                d_325_v94_ = _dafny.SeqWithoutIsStrInference([d_214_v10_])
                rhs70_ = (d_214_v10_) == ((d_214_v10_) + ((d_325_v94_)[default__.safeIndex(779, len(d_325_v94_))]))
                rhs71_ = default__.fm2(d_218_globalState_)
                rhs72_ = d_324_v93_
                d_211_v7_ = rhs70_
                d_205_v2_ = rhs71_
                d_324_v93_ = rhs72_
                d_326_v95_: C3
                nw64_ = C3()
                nw64_.ctor__((d_206_v3_)[default__.safeIndex(595, (d_206_v3_).length(0))], d_211_v7_, _dafny.SeqWithoutIsStrInference([-573, (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))]]))
                d_326_v95_ = nw64_
                d_327_v96_: _dafny.Array
                def lambda16_(d_328_v2_):
                    def lambda17_(d_329_i12_):
                        return d_328_v2_

                    return lambda17_

                init8_ = lambda16_(d_205_v2_)
                nw65_ = _dafny.Array(None, 8)
                for i0_8_ in range(nw65_.length(0)):
                    nw65_[i0_8_] = init8_(i0_8_)
                d_327_v96_ = nw65_
                index74_ = default__.safeIndex(284, (d_327_v96_).length(0))
                (d_327_v96_)[index74_] = d_205_v2_
                d_330_v97_: C8
                nw66_ = C8()
                nw66_.ctor__(d_203_v0_, d_284_v55_, d_211_v7_, _dafny.SeqWithoutIsStrInference([(d_217_v13_)[default__.safeIndex(814, len(d_217_v13_))], len(default__.fm17(d_218_globalState_)), (d_227_v19_)[default__.safeIndex(405, (d_227_v19_).length(0))], (d_217_v13_)[default__.safeIndex(d_285_i10_, len(d_217_v13_))]]))
                d_330_v97_ = nw66_
                d_331_v98_: _dafny.Map
                d_331_v98_ = _dafny.Map({d_330_v97_: d_211_v7_})
                index75_ = default__.safeIndex(405, (d_227_v19_).length(0))
                index76_ = default__.safeIndex(284, (d_327_v96_).length(0))
                rhs73_ = ((d_331_v98_)[d_330_v97_] if (d_330_v97_) in (d_331_v98_) else d_211_v7_)
                rhs74_ = (d_324_v93_).f22
                rhs75_ = d_326_v95_
                rhs76_ = _dafny.CodePoint('i')
                lhs70_ = d_324_v93_
                lhs71_ = d_227_v19_
                lhs72_ = default__.safeIndex(405, (d_227_v19_).length(0))
                lhs73_ = d_327_v96_
                lhs74_ = default__.safeIndex(284, (d_327_v96_).length(0))
                lhs70_.f23 = rhs73_
                lhs71_[lhs72_] = rhs74_
                d_326_v95_ = rhs75_
                lhs73_[lhs74_] = rhs76_
        _dafny.print((d_203_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_204_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_205_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_206_v3_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v4_)[0]) == (_dafny.MultiSet([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_210_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_211_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v8_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v9_) == (_dafny.Map({879: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v10_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_215_v11_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v12_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([879, 879, 879, 879, 2])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v13_) == (_dafny.SeqWithoutIsStrInference([879, 879]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_218_globalState_).f0)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_218_globalState_).f1)[0]) == (_dafny.MultiSet([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_.f8) == (_dafny.Map({318: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f9) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f10) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([879, 879, 879, 879, 2])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f12) == (_dafny.SeqWithoutIsStrInference([879, 879]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f13) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_220_v14_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[0]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[1]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[2]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[3]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[4]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[5]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[6]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[7]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[8]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[9]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[10]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[11]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[12]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[13]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[14]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[15]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[16]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[17]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[18]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[19]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[20]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v15_)[21]) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v16_) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v17_) == (_dafny.Map({_dafny.CodePoint('f'): 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v18_) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v19_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_228_v20_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v21_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_230_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_233_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v30_) == (_dafny.Map({True: _dafny.CodePoint('t')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v31_) == (_dafny.Map({False: _dafny.Map({True: _dafny.CodePoint('t')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v32_) == (_dafny.MultiSet([_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v52_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_281_v53_).f20)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v54_).cf19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v55_) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
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
        return f'D0.DC1({self.cf1.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(False, int(0), _dafny.Map({}))
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

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(_dafny.Map({}), _dafny.Array(None, 0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
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
        return lambda: D3_DC10(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D4_DC12(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC14(D5, NamedTuple('DC14', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf27 == __o.cf27
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
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC19(D7, NamedTuple('DC19', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(int(0), False, False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC21(D8, NamedTuple('DC21', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26(int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)

class D9_DC26(D9, NamedTuple('DC26', [('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC28(D10, NamedTuple('DC28', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)

class D11_DC30(D11, NamedTuple('DC30', [('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC34(_dafny.Array(None, 0), int(0), _dafny.Array(None, 0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)

class D12_DC34(D12, NamedTuple('DC34', [('cf59', Any), ('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf63', Any), ('cf64', Any), ('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)

class D13_DC37(D13, NamedTuple('DC37', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC39(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)

class D14_DC39(D14, NamedTuple('DC39', [('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC38(D14, NamedTuple('DC38', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC41(int(0), _dafny.Seq({}), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D15_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)

class D15_DC41(D15, NamedTuple('DC41', [('cf74', Any), ('cf75', Any), ('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC41({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {self.cf77.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC41) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf78', Any), ('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC43(D15, NamedTuple('DC43', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC45(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D16_DC47)

class D16_DC45(D16, NamedTuple('DC45', [('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC46(D16, NamedTuple('DC46', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC47(D16, NamedTuple('DC47', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC47({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC47) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    def m1(self, p0, p1, p2, p3, globalState):
        pass


class T1:
    pass
    def m11(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: bool = False
        self.f5: bool = False
        self.f8: _dafny.Map = _dafny.Map({})
        self.f15: str = _dafny.CodePoint('D')
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f3: int = int(0)
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f6: bool = False
        self._f7: int = int(0)
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: _dafny.Seq = _dafny.Seq({})
        self._f11: bool = False
        self._f12: _dafny.Seq = _dafny.Seq({})
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15

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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f9(self):
        return self._f9
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
    @property
    def f14(self):
        return self._f14

class C0:
    def  __init__(self):
        self.f27: _dafny.Array = _dafny.Array(None, 0)
        self._f26: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f26, f27):
        (self)._f26 = f26
        (self).f27 = f27

    @property
    def f26(self):
        return self._f26

class C1(T0, T1):
    def  __init__(self):
        self._f16: bool = False
        self._f17: _dafny.Seq = _dafny.Seq({})
        self._f28: int = int(0)
        self._f29: int = int(0)
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f29, f30, f16, f17, f28):
        (self)._f29 = f29
        (self)._f30 = f30
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f28 = f28

    def fm13(self, p0, globalState):
        return 519

    def fm14(self, p0, p1, p2, globalState):
        def iife48_():
            coll32_ = _dafny.Set()
            compr_32_: int
            for compr_32_ in _dafny.IntegerRange(847, 900):
                d_332_v0_: int = compr_32_
                if ((847) <= (d_332_v0_)) and ((d_332_v0_) < (900)):
                    coll32_ = coll32_.union(_dafny.Set([(d_332_v0_) + ((self).f29)]))
            return _dafny.Set(coll32_)
        return default__.safeModuloInt(((self).f29) - ((self).f29), len(iife48_()
        ))

    def fm15(self, p0, p1, p2, p3, globalState):
        return (D4_DC12(self.f16, self.f16, (self).f28)).cf23

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        d_333_v0_: int
        d_333_v0_ = 138
        rhs77_ = ((p1 if self.f16 else self.f16)) or (p2)
        rhs78_ = default__.fm3(globalState)
        lhs75_ = self
        lhs75_.f16 = rhs77_
        d_333_v0_ = rhs78_
        d_334_v1_: _dafny.Map
        d_334_v1_ = _dafny.Map({((self).f29) - (253): p1})
        d_334_v1_ = (d_334_v1_).set((self).f30, self.f16)
        d_335_v2_: _dafny.Array
        nw67_ = _dafny.Array(False, 9)
        d_335_v2_ = nw67_
        index77_ = default__.safeIndex(135, (d_335_v2_).length(0))
        (d_335_v2_)[index77_] = self.f16
        d_336_v3_: str
        d_336_v3_ = _dafny.CodePoint('r')
        d_337_v4_: _dafny.Seq
        d_337_v4_ = _dafny.SeqWithoutIsStrInference([((0) - ((self).fm14((self).f28, len(p0), d_336_v3_, globalState))) - (d_333_v0_)])
        index78_ = default__.safeIndex(135, (d_335_v2_).length(0))
        rhs79_ = not(p1)
        rhs80_ = self.f16
        rhs81_ = (d_333_v0_) != ((self).f29)
        rhs82_ = (self).f17
        lhs76_ = globalState
        lhs77_ = globalState
        lhs78_ = d_335_v2_
        lhs79_ = default__.safeIndex(135, (d_335_v2_).length(0))
        lhs76_.f2 = rhs79_
        lhs77_.f5 = rhs80_
        lhs78_[lhs79_] = rhs81_
        d_337_v4_ = rhs82_
        d_338_v5_: C0
        nw68_ = C0()
        nw68_.ctor__((default__.fm16((self).f30, globalState)).cardinality, d_335_v2_)
        d_338_v5_ = nw68_
        d_339_v6_: D0
        d_339_v6_ = D0_DC0(p0)
        pat_let_tv11_ = globalState
        def iife49_(_pat_let8_0):
            def iife50_(d_340_dt__update__tmp_h0_):
                def iife51_(_pat_let9_0):
                    def iife52_(d_341_dt__update_hcf0_h0_):
                        return D0_DC0(d_341_dt__update_hcf0_h0_)
                    return iife52_(_pat_let9_0)
                return iife51_(default__.fm1(pat_let_tv11_))
            return iife50_(_pat_let8_0)
        source3_ = iife49_(d_339_v6_)
        if source3_.is_DC1:
            d_342___mcc_h0_ = source3_.cf1
            d_343_cf1_ = d_342___mcc_h0_
            d_344_v7_: _dafny.Seq
            d_345_v8_: bool
            d_346_v9_: bool
            out0_: _dafny.Seq
            out1_: bool
            out2_: bool
            out0_, out1_, out2_ = (self).m12(globalState)
            d_344_v7_ = out0_
            d_345_v8_ = out1_
            d_346_v9_ = out2_
            d_347_v10_: _dafny.Set
            d_347_v10_ = _dafny.Set({len(d_343_cf1_)})
            d_348_v11_: C0
            nw69_ = C0()
            nw69_.ctor__(default__.safeDivisionInt(d_333_v0_, len(d_347_v10_)), d_335_v2_)
            d_348_v11_ = nw69_
            (globalState).f2 = d_346_v9_
            d_333_v0_ = default__.safeModuloInt(((d_338_v5_).f26) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtmijpy")))), p3)
        elif source3_.is_DC2:
            d_349___mcc_h1_ = source3_.cf2
            d_350___mcc_h2_ = source3_.cf3
            d_351___mcc_h3_ = source3_.cf4
            d_352___mcc_h4_ = source3_.cf5
            d_353_cf5_ = d_352___mcc_h4_
            d_354_cf4_ = d_351___mcc_h3_
            d_355_cf3_ = d_350___mcc_h2_
            d_356_cf2_ = d_349___mcc_h1_
            d_357_v12_: _dafny.Array
            nw70_ = _dafny.Array(int(0), 20)
            d_357_v12_ = nw70_
            index79_ = default__.safeIndex(871, (d_357_v12_).length(0))
            (d_357_v12_)[index79_] = d_333_v0_
            index80_ = default__.safeIndex(871, (d_357_v12_).length(0))
            (d_357_v12_)[index80_] = (0) - ((self).f28)
            d_358_v13_: _dafny.Set
            d_358_v13_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_336_v3_ for d_359_i0_ in range(default__.abs(691))]), p0})
            (globalState).f5 = default__.fm0((d_358_v13_) == (_dafny.Set({p0, p0})), p2, globalState)
            d_360_v14_: C0
            nw71_ = C0()
            nw71_.ctor__((d_338_v5_).f26, (d_353_cf5_ if p2 else d_338_v5_.f27))
            d_360_v14_ = nw71_
            d_354_cf4_ = (d_360_v14_).f26
        elif True:
            d_361___mcc_h5_ = source3_.cf0
            d_362_cf0_ = d_361___mcc_h5_
            d_363_v15_: _dafny.Map
            d_363_v15_ = _dafny.Map({(d_338_v5_).f26: d_335_v2_})
            d_364_v16_: _dafny.Map
            d_364_v16_ = _dafny.Map({d_333_v0_: ((d_363_v15_)[(self).f30] if ((self).f30) in (d_363_v15_) else d_338_v5_.f27)})
            d_364_v16_ = (d_364_v16_).set(297, d_338_v5_.f27)
            d_365_v17_: _dafny.Map
            d_365_v17_ = _dafny.Map({p2: (d_335_v2_)[default__.safeIndex(135, (d_335_v2_).length(0))]})
            d_366_v18_: D1
            d_366_v18_ = D1_DC4(p2, (d_338_v5_).f26, (d_365_v17_) | (d_365_v17_))
            d_367_v19_: _dafny.Map
            d_367_v19_ = _dafny.Map({d_336_v3_: len((self).f17)})
            d_368_v20_: _dafny.Map
            d_368_v20_ = _dafny.Map({p2: d_367_v19_})
            d_366_v18_ = D1_DC4(False, len(d_368_v20_), d_365_v17_)
            if (False) and (((d_338_v5_).f26) == ((d_338_v5_).f26)):
                def iife53_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in ((self).f17).Elements:
                        d_369_v21_: int = compr_33_
                        if (d_369_v21_) in ((self).f17):
                            coll33_[default__.safeModuloInt(d_369_v21_, d_333_v0_)] = (self).f28
                    return _dafny.Map(coll33_)
                d_333_v0_ = default__.safeDivisionInt((self).f28, len(iife53_()
                ))
                d_370_v22_: C0
                nw72_ = C0()
                nw72_.ctor__(d_333_v0_, d_335_v2_)
                d_370_v22_ = nw72_
                d_371_v23_: _dafny.Map
                d_371_v23_ = _dafny.Map({(d_335_v2_)[default__.safeIndex(135, (d_335_v2_).length(0))]: 443})
                (globalState).f5 = (p2) not in (d_371_v23_)
                d_372_v24_: _dafny.Array
                nw73_ = _dafny.Array(D2.default()(), 9)
                d_372_v24_ = nw73_
                d_373_v25_: _dafny.Array
                def lambda18_(d_374_cf0_):
                    def lambda19_(d_375_i1_):
                        return d_374_cf0_

                    return lambda19_

                init9_ = lambda18_(d_362_cf0_)
                nw74_ = _dafny.Array(None, 18)
                for i0_9_ in range(nw74_.length(0)):
                    nw74_[i0_9_] = init9_(i0_9_)
                d_373_v25_ = nw74_
                index81_ = default__.safeIndex(713, (d_373_v25_).length(0))
                (d_373_v25_)[index81_] = (d_362_cf0_) + (p0)
                d_376_v26_: _dafny.Array
                d_376_v26_ = d_372_v24_
                index82_ = default__.safeIndex(713, (d_373_v25_).length(0))
                rhs83_ = (p1 if self.f16 else self.f16)
                rhs84_ = (d_376_v26_)
                rhs85_ = d_336_v3_
                rhs86_ = (p0).set(default__.safeIndex(len((default__.fm17(globalState)) + (d_337_v4_)), len(p0)), d_336_v3_)
                rhs87_ = p1
                lhs80_ = globalState
                lhs81_ = globalState
                lhs82_ = d_373_v25_
                lhs83_ = default__.safeIndex(713, (d_373_v25_).length(0))
                lhs84_ = self
                lhs80_.f5 = rhs83_
                d_372_v24_ = rhs84_
                lhs81_.f15 = rhs85_
                lhs82_[lhs83_] = rhs86_
                lhs84_.f16 = rhs87_
                d_377_v27_: _dafny.Seq
                d_377_v27_ = _dafny.SeqWithoutIsStrInference([d_334_v1_, d_334_v1_, d_334_v1_, d_334_v1_])
                d_378_v28_: _dafny.Array
                nw75_ = _dafny.Array(None, 1)
                nw75_[int(0)] = (d_377_v27_)[default__.safeIndex((d_370_v22_).f26, len(d_377_v27_))]
                d_378_v28_ = nw75_
                def iife54_():
                    coll34_ = _dafny.Map()
                    compr_34_: int
                    for compr_34_ in ((self).f17).Elements:
                        d_379_v29_: int = compr_34_
                        if (d_379_v29_) in ((self).f17):
                            coll34_[(d_379_v29_) + ((d_338_v5_).f26)] = p1
                    return _dafny.Map(coll34_)
                rhs88_ = (default__.safeModuloInt(p3, 151)) < ((self).f30)
                rhs89_ = ((self).fm14(p3, (d_338_v5_).f26, d_336_v3_, globalState)) > (len((_dafny.Map({len(d_365_v17_): (d_335_v2_)[default__.safeIndex(135, (d_335_v2_).length(0))]})) | (iife54_()
                )))
                rhs90_ = d_378_v28_
                lhs85_ = globalState
                lhs86_ = globalState
                lhs85_.f2 = rhs88_
                lhs86_.f5 = rhs89_
                d_378_v28_ = rhs90_
            elif True:
                (globalState).f2 = (((d_338_v5_).f26) * ((self).f29)) > (((self).f17)[default__.safeIndex(d_333_v0_, len((self).f17))])
                (globalState).f15 = d_336_v3_
                d_380_v30_: C0
                nw76_ = C0()
                nw76_.ctor__((self).fm13(p1, globalState), d_338_v5_.f27)
                d_380_v30_ = nw76_
                d_381_v31_: _dafny.MultiSet
                d_381_v31_ = _dafny.MultiSet([(d_380_v30_).f26])
                d_381_v31_ = (_dafny.MultiSet((self).f17)).intersection(d_381_v31_)
                (globalState).f5 = ((self).f17) <= (d_337_v4_)
            d_382_v32_: _dafny.Array
            nw77_ = _dafny.Array(int(0), 1)
            d_382_v32_ = nw77_
            d_383_v33_: _dafny.MultiSet
            d_383_v33_ = _dafny.MultiSet([self.f16, default__.fm0(not(p1), default__.fm0(p2, self.f16, globalState), globalState), True])
            d_384_v34_: _dafny.Map
            d_384_v34_ = _dafny.Map({d_382_v32_: (0) - ((d_383_v33_).cardinality)})
            d_384_v34_ = d_384_v34_
        d_385_v35_: _dafny.Array
        nw78_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_385_v35_ = nw78_
        index83_ = default__.safeIndex(851, (d_385_v35_).length(0))
        (d_385_v35_)[index83_] = d_335_v2_
        d_386_v36_: _dafny.Seq
        d_386_v36_ = _dafny.SeqWithoutIsStrInference([self.f16])
        pat_let_tv12_ = d_338_v5_
        pat_let_tv13_ = d_338_v5_
        pat_let_tv14_ = d_386_v36_
        pat_let_tv15_ = p2
        index84_ = default__.safeIndex(851, (d_385_v35_).length(0))
        def lambda20_(source4_):
            if source4_.is_DC12:
                d_387___mcc_h6_ = source4_.cf21
                d_388___mcc_h7_ = source4_.cf22
                d_389___mcc_h8_ = source4_.cf23
                d_390_cf23_ = d_389___mcc_h8_
                d_391_cf22_ = d_388___mcc_h7_
                d_392_cf21_ = d_387___mcc_h6_
                return (pat_let_tv12_).f26
            elif source4_.is_DC13:
                d_393___mcc_h9_ = source4_.cf24
                d_394___mcc_h10_ = source4_.cf25
                d_395___mcc_h11_ = source4_.cf26
                d_396_cf26_ = d_395___mcc_h11_
                d_397_cf25_ = d_394___mcc_h10_
                d_398_cf24_ = d_393___mcc_h9_
                return default__.safeModuloInt(len(_dafny.Map({(pat_let_tv13_).f26: self.f16})), len(pat_let_tv14_))
            elif True:
                d_399___mcc_h12_ = source4_.cf20
                d_400_cf20_ = d_399___mcc_h12_
                return (0) - ((_dafny.MultiSet([_dafny.Set({pat_let_tv15_})])).cardinality)

        rhs91_ = ((d_386_v36_) + (d_386_v36_)) + (d_386_v36_)
        rhs92_ = lambda20_(D4_DC13((self).f28, (self).f30, (self).f30))
        rhs93_ = (self).f29
        rhs94_ = p3
        rhs95_ = d_338_v5_.f27
        lhs87_ = d_385_v35_
        lhs88_ = default__.safeIndex(851, (d_385_v35_).length(0))
        r0 = rhs91_
        d_333_v0_ = rhs92_
        d_333_v0_ = rhs93_
        d_333_v0_ = rhs94_
        lhs87_[lhs88_] = rhs95_
        r0 = _dafny.SeqWithoutIsStrInference([p2])
        return r0

    def m11(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: bool = False
        r2: int = int(0)
        hi2_ = ((self).f29) + ((self).f30)
        for d_401_i0_ in range((self).f28, hi2_):
            r2 = default__.fm3(globalState)
            r0 = default__.fm2(globalState)
            d_402_v0_: _dafny.Array
            nw79_ = _dafny.Array(False, 2)
            d_402_v0_ = nw79_
            d_403_v1_: C0
            nw80_ = C0()
            nw80_.ctor__(d_401_i0_, d_402_v0_)
            d_403_v1_ = nw80_
            d_404_v2_: _dafny.Seq
            d_404_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anddf"))
            d_405_v3_: _dafny.MultiSet
            d_405_v3_ = _dafny.MultiSet([d_402_v0_])
            d_406_v4_: _dafny.Map
            d_406_v4_ = _dafny.Map({d_404_v2_: (d_405_v3_) - (d_405_v3_)})
            d_406_v4_ = (d_406_v4_).set((d_404_v2_ if self.f16 else d_404_v2_), d_405_v3_)
        hi3_ = len((self).f17)
        for d_407_i1_ in range((self).f28, hi3_):
            d_408_v5_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 6)
            d_408_v5_ = nw81_
            d_409_v6_: _dafny.Map
            d_409_v6_ = _dafny.Map({(self).f17: (self).f30})
            index85_ = default__.safeIndex(666, (d_408_v5_).length(0))
            (d_408_v5_)[index85_] = ((d_409_v6_)[(self).f17] if ((self).f17) in (d_409_v6_) else d_407_i1_)
            index86_ = default__.safeIndex(864, (d_408_v5_).length(0))
            (d_408_v5_)[index86_] = 759
            d_410_v7_: _dafny.Map
            d_410_v7_ = _dafny.Map({self.f16: 13})
            d_411_v8_: _dafny.Seq
            d_411_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lw"))
            d_412_v9_: _dafny.MultiSet
            d_412_v9_ = _dafny.MultiSet([d_410_v7_, d_410_v7_, (_dafny.Map({self.f16: default__.fm3(globalState)})) | (_dafny.Map({self.f16: len(d_411_v8_)}))])
            d_413_v10_: _dafny.Set
            d_413_v10_ = _dafny.Set({self.f16})
            d_414_v11_: _dafny.MultiSet
            d_414_v11_ = _dafny.MultiSet([(self).f30, 393])
            index87_ = default__.safeIndex(666, (d_408_v5_).length(0))
            index88_ = default__.safeIndex(864, (d_408_v5_).length(0))
            rhs96_ = (d_412_v9_).cardinality
            rhs97_ = len(d_411_v8_)
            rhs98_ = (not((d_413_v10_).issubset(d_413_v10_))) and (not(((self).f17) == ((self).f17)))
            rhs99_ = self.f16
            rhs100_ = (not (self.f16) or (default__.fm0(self.f16, self.f16, globalState))) in (_dafny.Map({self.f16: d_414_v11_}))
            lhs89_ = d_408_v5_
            lhs90_ = default__.safeIndex(666, (d_408_v5_).length(0))
            lhs91_ = d_408_v5_
            lhs92_ = default__.safeIndex(864, (d_408_v5_).length(0))
            lhs93_ = globalState
            lhs94_ = globalState
            lhs95_ = globalState
            lhs89_[lhs90_] = rhs96_
            lhs91_[lhs92_] = rhs97_
            lhs93_.f2 = rhs98_
            lhs94_.f5 = rhs99_
            lhs95_.f5 = rhs100_
            d_415_v12_: D1
            d_415_v12_ = D1_DC4(self.f16, (43) + ((self).f29), _dafny.Map({self.f16: self.f16}))
            source5_ = d_415_v12_
            if source5_.is_DC4:
                d_416___mcc_h0_ = source5_.cf7
                d_417___mcc_h1_ = source5_.cf8
                d_418___mcc_h2_ = source5_.cf9
                d_419_cf9_ = d_418___mcc_h2_
                d_420_cf8_ = d_417___mcc_h1_
                d_421_cf7_ = d_416___mcc_h0_
                d_422_v13_: _dafny.Seq
                d_422_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hndbi"))])
                index89_ = default__.safeIndex(666, (d_408_v5_).length(0))
                (d_408_v5_)[index89_] = (d_407_i1_) + (default__.safeDivisionInt(len(d_422_v13_), (self).f30))
                d_423_v14_: _dafny.Seq
                d_423_v14_ = _dafny.SeqWithoutIsStrInference([False])
                d_420_cf8_ = (len(d_423_v14_)) + ((self).f29)
                d_424_v15_: _dafny.Seq
                d_424_v15_ = _dafny.SeqWithoutIsStrInference([d_408_v5_])
                index90_ = default__.safeIndex(666, (d_408_v5_).length(0))
                (d_408_v5_)[index90_] = len((d_424_v15_) + (d_424_v15_))
                d_425_v16_: _dafny.Array
                def lambda21_(d_426_i2_):
                    return False

                init10_ = lambda21_
                nw82_ = _dafny.Array(None, 18)
                for i0_10_ in range(nw82_.length(0)):
                    nw82_[i0_10_] = init10_(i0_10_)
                d_425_v16_ = nw82_
                d_427_v17_: C0
                nw83_ = C0()
                nw83_.ctor__((d_408_v5_)[default__.safeIndex(666, (d_408_v5_).length(0))], d_425_v16_)
                d_427_v17_ = nw83_
            elif source5_.is_DC5:
                d_428___mcc_h3_ = source5_.cf10
                d_429___mcc_h4_ = source5_.cf11
                d_430_cf11_ = d_429___mcc_h4_
                d_431_cf10_ = d_428___mcc_h3_
                d_432_v18_: _dafny.MultiSet
                d_432_v18_ = _dafny.MultiSet([d_430_cf11_, not(self.f16)])
                d_433_v19_: D4
                d_433_v19_ = D4_DC12(self.f16, (self.f16 if d_430_cf11_ else d_430_cf11_), (d_432_v18_).cardinality)
                d_433_v19_ = d_433_v19_
                d_434_v20_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_434_v20_ = nw84_
                d_434_v20_ = d_434_v20_
                (self).f16 = self.f16
                d_435_v21_: _dafny.Array
                def lambda22_(d_436_cf11_):
                    def lambda23_(d_437_i3_):
                        return d_436_cf11_

                    return lambda23_

                init11_ = lambda22_(d_430_cf11_)
                nw85_ = _dafny.Array(None, 23)
                for i0_11_ in range(nw85_.length(0)):
                    nw85_[i0_11_] = init11_(i0_11_)
                d_435_v21_ = nw85_
                index91_ = default__.safeIndex(478, (d_435_v21_).length(0))
                (d_435_v21_)[index91_] = True
                index92_ = default__.safeIndex(478, (d_435_v21_).length(0))
                (d_435_v21_)[index92_] = self.f16
            elif source5_.is_DC6:
                d_438_v22_: _dafny.Array
                nw86_ = _dafny.Array(False, 19)
                d_438_v22_ = nw86_
                d_439_v23_: C0
                nw87_ = C0()
                nw87_.ctor__((self).f29, d_438_v22_)
                d_439_v23_ = nw87_
                d_440_v24_: _dafny.MultiSet
                d_440_v24_ = _dafny.MultiSet([self.f16, (_dafny.SeqWithoutIsStrInference([(self).f30])) == (((self).f17).set(default__.safeIndex((self).f29, len((self).f17)), d_407_i1_)), self.f16, self.f16, True])
                d_441_v25_: _dafny.Set
                d_441_v25_ = _dafny.Set({(self).f29, len((d_410_v7_) | (d_410_v7_))})
                index93_ = default__.safeIndex(666, (d_408_v5_).length(0))
                rhs101_ = len(_dafny.Set({d_411_v8_, (d_411_v8_) + (default__.fm1(globalState)), d_411_v8_, (d_411_v8_) + (d_411_v8_)}))
                rhs102_ = (d_440_v24_) - (_dafny.MultiSet([self.f16]))
                rhs103_ = (d_441_v25_).intersection(d_441_v25_)
                lhs96_ = d_408_v5_
                lhs97_ = default__.safeIndex(666, (d_408_v5_).length(0))
                lhs96_[lhs97_] = rhs101_
                d_440_v24_ = rhs102_
                d_441_v25_ = rhs103_
                d_442_v26_: _dafny.Map
                d_442_v26_ = _dafny.Map({(d_408_v5_)[default__.safeIndex(666, (d_408_v5_).length(0))]: self.f16})
                d_443_v27_: C0
                nw88_ = C0()
                nw88_.ctor__(default__.safeModuloInt((0) - ((0) - (len(d_442_v26_))), (self).f29), d_439_v23_.f27)
                d_443_v27_ = nw88_
                d_444_v28_: str
                d_444_v28_ = _dafny.CodePoint('y')
                (globalState).f15 = d_444_v28_
            elif True:
                d_445___mcc_h5_ = source5_.cf6
                d_446_cf6_ = d_445___mcc_h5_
                index94_ = default__.safeIndex(666, (d_408_v5_).length(0))
                (d_408_v5_)[index94_] = len(d_411_v8_)
                (globalState).f5 = (len(_dafny.Map({(self).f28: d_408_v5_}))) < ((self).f30)
                d_447_v29_: _dafny.Seq
                d_448_v30_: bool
                d_449_v31_: bool
                out3_: _dafny.Seq
                out4_: bool
                out5_: bool
                out3_, out4_, out5_ = (self).m12(globalState)
                d_447_v29_ = out3_
                d_448_v30_ = out4_
                d_449_v31_ = out5_
                d_450_v32_: _dafny.Map
                d_450_v32_ = _dafny.Map({D1_DC3(d_446_cf6_): (d_447_v29_) + (d_447_v29_)})
                d_451_v33_: D1
                d_451_v33_ = D1_DC3(d_446_cf6_)
                d_452_v34_: str
                d_452_v34_ = _dafny.CodePoint('q')
                rhs104_ = ((((d_450_v32_)[d_451_v33_] if (d_451_v33_) in (d_450_v32_) else d_411_v8_)).set(default__.safeIndex((self).f30, len(((d_450_v32_)[d_451_v33_] if (d_451_v33_) in (d_450_v32_) else d_411_v8_))), d_452_v34_)).set(default__.safeIndex((self).f28, len((((d_450_v32_)[d_451_v33_] if (d_451_v33_) in (d_450_v32_) else d_411_v8_)).set(default__.safeIndex((self).f30, len(((d_450_v32_)[d_451_v33_] if (d_451_v33_) in (d_450_v32_) else d_411_v8_))), d_452_v34_))), _dafny.CodePoint('i'))
                rhs105_ = self.f16
                rhs106_ = default__.fm0(d_449_v31_, d_448_v30_, globalState)
                rhs107_ = d_448_v30_
                rhs108_ = d_449_v31_
                lhs98_ = globalState
                lhs99_ = globalState
                d_447_v29_ = rhs104_
                lhs98_.f2 = rhs105_
                r1 = rhs106_
                r1 = rhs107_
                lhs99_.f5 = rhs108_
            d_453_v36_: _dafny.Map
            d_453_v36_ = _dafny.Map({(self).f30: d_407_i1_})
            d_454_v38_: _dafny.Seq
            d_454_v38_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_455_v39_: _dafny.Seq
            d_455_v39_ = _dafny.SeqWithoutIsStrInference([d_454_v38_])
            d_456_v40_: _dafny.Seq
            d_456_v40_ = _dafny.SeqWithoutIsStrInference([(d_453_v36_).set(len((d_455_v39_)[default__.safeIndex(len(d_454_v38_), len(d_455_v39_))]), (self).f30)])
            d_457_v42_: _dafny.Array
            nw89_ = _dafny.Array(None, 12)
            def iife55_():
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in (((self).f17).set(default__.safeIndex((self).f29, len((self).f17)), (self).f29)).Elements:
                    d_458_v35_: int = compr_35_
                    if (d_458_v35_) in (((self).f17).set(default__.safeIndex((self).f29, len((self).f17)), (self).f29)):
                        coll35_[default__.safeModuloInt(d_458_v35_, d_407_i1_)] = (d_408_v5_)[default__.safeIndex(666, (d_408_v5_).length(0))]
                return _dafny.Map(coll35_)
            nw89_[int(0)] = iife55_()
            
            nw89_[int(1)] = (d_453_v36_) | (d_453_v36_)
            nw89_[int(2)] = (d_453_v36_) | (d_453_v36_)
            def iife56_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in ((self).f17).Elements:
                    d_459_v37_: int = compr_36_
                    if (d_459_v37_) in ((self).f17):
                        coll36_[default__.safeModuloInt(d_459_v37_, (d_408_v5_)[default__.safeIndex(666, (d_408_v5_).length(0))])] = (0) - ((self).f30)
                return _dafny.Map(coll36_)
            nw89_[int(3)] = iife56_()
            
            nw89_[int(4)] = d_453_v36_
            nw89_[int(5)] = _dafny.Map({d_407_i1_: (self).f30})
            nw89_[int(6)] = d_453_v36_
            nw89_[int(7)] = (d_456_v40_)[default__.safeIndex(((d_410_v7_)[self.f16] if (self.f16) in (d_410_v7_) else d_407_i1_), len(d_456_v40_))]
            def iife57_():
                coll37_ = _dafny.Map()
                compr_37_: int
                for compr_37_ in (d_414_v11_).Elements:
                    d_460_v41_: int = compr_37_
                    if (d_460_v41_) in (d_414_v11_):
                        coll37_[(d_460_v41_) + ((self).f29)] = len(d_454_v38_)
                return _dafny.Map(coll37_)
            nw89_[int(8)] = iife57_()
            
            nw89_[int(9)] = d_453_v36_
            nw89_[int(10)] = d_453_v36_
            nw89_[int(11)] = (d_453_v36_).set((self).f29, (self).f29)
            d_457_v42_ = nw89_
            d_457_v42_ = d_457_v42_
            r2 = default__.safeModuloInt(60, (d_407_i1_) * (d_407_i1_))
        d_461_v43_: _dafny.Array
        nw90_ = _dafny.Array(int(0), 24)
        d_461_v43_ = nw90_
        index95_ = default__.safeIndex(495, (d_461_v43_).length(0))
        (d_461_v43_)[index95_] = (self).f29
        index96_ = default__.safeIndex(495, (d_461_v43_).length(0))
        rhs109_ = (self).f30
        rhs110_ = (self).f28
        lhs100_ = d_461_v43_
        lhs101_ = default__.safeIndex(495, (d_461_v43_).length(0))
        lhs100_[lhs101_] = rhs109_
        r2 = rhs110_
        d_462_v44_: _dafny.Array
        nw91_ = _dafny.Array(False, 27)
        d_462_v44_ = nw91_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_462_v44_).length(0)):
            d_463_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_463_i4_)) and ((d_463_i4_) < ((d_462_v44_).length(0)))):
                (d_462_v44_)[(d_463_i4_)] = self.f16
        d_464_v45_: D4
        d_464_v45_ = D4_DC13((d_461_v43_)[default__.safeIndex(495, (d_461_v43_).length(0))], default__.safeModuloInt(len(_dafny.Map({not(self.f16): d_462_v44_})), (self).f28), -224)
        source6_ = d_464_v45_
        if source6_.is_DC12:
            d_465___mcc_h6_ = source6_.cf21
            d_466___mcc_h7_ = source6_.cf22
            d_467___mcc_h8_ = source6_.cf23
            d_468_cf23_ = d_467___mcc_h8_
            d_469_cf22_ = d_466___mcc_h7_
            d_470_cf21_ = d_465___mcc_h6_
            d_471_v46_: _dafny.Seq
            d_471_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wivplsm"))
            d_472_v47_: _dafny.Seq
            d_472_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wugmj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kamk")), d_471_v46_, d_471_v46_])
            d_473_v48_: _dafny.MultiSet
            d_473_v48_ = _dafny.MultiSet([self.f16, False])
            d_474_v49_: _dafny.Map
            d_474_v49_ = _dafny.Map({d_472_v47_: (_dafny.MultiSet([not(False), d_470_cf21_, self.f16]) if self.f16 else d_473_v48_)})
            d_474_v49_ = (d_474_v49_).set((default__.fm18((0) - (d_468_cf23_), globalState)) + (d_472_v47_), d_473_v48_)
            (self).f16 = self.f16
            if d_470_cf21_:
                d_475_v50_: _dafny.Set
                d_475_v50_ = _dafny.Set({((self).f28) + (748), (self).f28, 811})
                d_476_v51_: _dafny.Array
                nw92_ = _dafny.Array(_dafny.Seq({}), 11)
                d_476_v51_ = nw92_
                index97_ = default__.safeIndex(581, (d_476_v51_).length(0))
                (d_476_v51_)[index97_] = _dafny.SeqWithoutIsStrInference([(self).fm13(d_470_cf21_, globalState)])
                index98_ = default__.safeIndex(581, (d_476_v51_).length(0))
                rhs111_ = d_475_v50_
                rhs112_ = ((self).f17) + ((self).f17)
                lhs102_ = d_476_v51_
                lhs103_ = default__.safeIndex(581, (d_476_v51_).length(0))
                d_475_v50_ = rhs111_
                lhs102_[lhs103_] = rhs112_
                index99_ = default__.safeIndex(58, (d_462_v44_).length(0))
                (d_462_v44_)[index99_] = self.f16
                d_477_v52_: _dafny.Array
                def lambda24_(d_478_cf22_):
                    def lambda25_(d_479_i5_):
                        return (_dafny.CodePoint('f') if d_478_cf22_ else _dafny.CodePoint('d'))

                    return lambda25_

                init12_ = lambda24_(d_469_cf22_)
                nw93_ = _dafny.Array(None, 20)
                for i0_12_ in range(nw93_.length(0)):
                    nw93_[i0_12_] = init12_(i0_12_)
                d_477_v52_ = nw93_
                index100_ = default__.safeIndex(610, (d_477_v52_).length(0))
                (d_477_v52_)[index100_] = _dafny.CodePoint('h')
                d_480_v53_: str
                d_480_v53_ = _dafny.CodePoint('q')
                index101_ = default__.safeIndex(58, (d_462_v44_).length(0))
                index102_ = default__.safeIndex(610, (d_477_v52_).length(0))
                rhs113_ = ((self).f29) != ((self).f29)
                rhs114_ = not(((self).f29) < (default__.safeDivisionInt((self).f29, (self).f29)))
                rhs115_ = d_480_v53_
                lhs104_ = d_462_v44_
                lhs105_ = default__.safeIndex(58, (d_462_v44_).length(0))
                lhs106_ = d_477_v52_
                lhs107_ = default__.safeIndex(610, (d_477_v52_).length(0))
                lhs104_[lhs105_] = rhs113_
                r1 = rhs114_
                lhs106_[lhs107_] = rhs115_
                d_480_v53_ = _dafny.CodePoint('i')
                d_475_v50_ = d_475_v50_
                r1 = ((0) - (d_468_cf23_)) != (((d_476_v51_)[default__.safeIndex(581, (d_476_v51_).length(0))])[default__.safeIndex(786, len((d_476_v51_)[default__.safeIndex(581, (d_476_v51_).length(0))]))])
            elif True:
                index103_ = default__.safeIndex(495, (d_461_v43_).length(0))
                (d_461_v43_)[index103_] = (self).f29
                d_481_v54_: _dafny.Map
                d_481_v54_ = _dafny.Map({self.f16: (self).f29})
                d_482_v55_: _dafny.Array
                def lambda26_(d_483_i6_):
                    return (self).f17

                init13_ = lambda26_
                nw94_ = _dafny.Array(None, 4)
                for i0_13_ in range(nw94_.length(0)):
                    nw94_[i0_13_] = init13_(i0_13_)
                d_482_v55_ = nw94_
                d_484_v56_: _dafny.Seq
                d_484_v56_ = _dafny.SeqWithoutIsStrInference([True])
                d_485_v57_: _dafny.Map
                d_485_v57_ = _dafny.Map({d_470_cf21_: D2_DC8(d_481_v54_, d_482_v55_, d_469_cf22_, d_469_cf22_, len(_dafny.Map({len(d_484_v56_): True})))})
                d_486_v58_: D2
                d_486_v58_ = D2_DC8(d_481_v54_, d_482_v55_, d_470_cf21_, default__.fm0(self.f16, self.f16, globalState), d_468_cf23_)
                d_485_v57_ = (d_485_v57_).set(default__.fm0(self.f16, d_469_cf22_, globalState), d_486_v58_)
                d_471_v46_ = default__.fm1(globalState)
                r2 = (self).fm13(d_469_cf22_, globalState)
                (globalState).f5 = d_469_cf22_
            if d_470_cf21_:
                d_487_v59_: _dafny.Array
                nw95_ = _dafny.Array(_dafny.Set({}), 22)
                d_487_v59_ = nw95_
                index104_ = default__.safeIndex(76, (d_487_v59_).length(0))
                (d_487_v59_)[index104_] = _dafny.Set({d_471_v46_})
                index105_ = default__.safeIndex(76, (d_487_v59_).length(0))
                (d_487_v59_)[index105_] = _dafny.Set({d_471_v46_})
                d_469_cf22_ = d_470_cf21_
                d_488_v60_: str
                d_488_v60_ = _dafny.CodePoint('p')
                d_489_v61_: _dafny.MultiSet
                d_489_v61_ = _dafny.MultiSet([d_488_v60_, _dafny.CodePoint('s')])
                d_490_v62_: _dafny.Seq
                d_490_v62_ = _dafny.SeqWithoutIsStrInference([d_489_v61_, (d_489_v61_).set(d_488_v60_, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sey"))))), _dafny.MultiSet([d_488_v60_, d_488_v60_, d_488_v60_, d_488_v60_, d_488_v60_]), d_489_v61_])
                d_491_v63_: _dafny.Map
                d_491_v63_ = _dafny.Map({len(d_490_v62_): d_468_cf23_})
                rhs116_ = d_491_v63_
                rhs117_ = d_469_cf22_
                lhs108_ = self
                d_491_v63_ = rhs116_
                lhs108_.f16 = rhs117_
                d_492_v64_: D0
                d_492_v64_ = D0_DC0(d_471_v46_)
                d_492_v64_ = d_492_v64_
                d_491_v63_ = (d_491_v63_).set((self).f28, (self).f29)
            elif True:
                d_493_v65_: D1
                d_493_v65_ = D1_DC3((d_473_v48_) - (d_473_v48_))
                d_493_v65_ = d_493_v65_
                d_494_v66_: C0
                nw96_ = C0()
                nw96_.ctor__((self).f29, (D0_DC2(len(d_471_v46_), (self).f28, (d_461_v43_)[default__.safeIndex(495, (d_461_v43_).length(0))], d_462_v44_)).cf5)
                d_494_v66_ = nw96_
                d_495_v67_: _dafny.Set
                d_495_v67_ = _dafny.Set({d_470_cf21_, d_469_cf22_})
                d_496_v68_: D2
                d_496_v68_ = D2_DC7(d_495_v67_)
                d_497_v69_: _dafny.Array
                nw97_ = _dafny.Array(None, 14)
                nw97_[int(0)] = d_496_v68_
                nw97_[int(1)] = d_496_v68_
                nw97_[int(2)] = d_496_v68_
                nw97_[int(3)] = D2_DC7(d_495_v67_)
                nw97_[int(4)] = d_496_v68_
                nw97_[int(5)] = d_496_v68_
                nw97_[int(6)] = d_496_v68_
                nw97_[int(7)] = d_496_v68_
                nw97_[int(8)] = d_496_v68_
                nw97_[int(9)] = d_496_v68_
                nw97_[int(10)] = d_496_v68_
                nw97_[int(11)] = D2_DC7(d_495_v67_)
                nw97_[int(12)] = d_496_v68_
                nw97_[int(13)] = d_496_v68_
                d_497_v69_ = nw97_
                d_498_v70_: str
                d_498_v70_ = _dafny.CodePoint('a')
                d_499_v71_: _dafny.Map
                d_499_v71_ = _dafny.Map({d_497_v69_: default__.safeDivisionInt((d_494_v66_).f26, (self).fm14((d_494_v66_).f26, (d_461_v43_)[default__.safeIndex(495, (d_461_v43_).length(0))], d_498_v70_, globalState))})
                d_499_v71_ = (d_499_v71_).set(d_497_v69_, len((self).f17))
                d_470_cf21_ = d_469_cf22_
                d_500_v72_: _dafny.Array
                nw98_ = _dafny.Array(_dafny.Map({}), 27)
                d_500_v72_ = nw98_
                d_501_v73_: _dafny.Map
                d_501_v73_ = _dafny.Map({False: d_470_cf21_})
                index106_ = default__.safeIndex(997, (d_500_v72_).length(0))
                (d_500_v72_)[index106_] = d_501_v73_
                index107_ = default__.safeIndex(997, (d_500_v72_).length(0))
                (d_500_v72_)[index107_] = d_501_v73_
        elif source6_.is_DC13:
            d_502___mcc_h9_ = source6_.cf24
            d_503___mcc_h10_ = source6_.cf25
            d_504___mcc_h11_ = source6_.cf26
            d_505_cf26_ = d_504___mcc_h11_
            d_506_cf25_ = d_503___mcc_h10_
            d_507_cf24_ = d_502___mcc_h9_
            index108_ = default__.safeIndex(495, (d_461_v43_).length(0))
            (d_461_v43_)[index108_] = (self).f28
            d_508_v74_: _dafny.Seq
            d_508_v74_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_508_v74_ = (d_508_v74_) + (d_508_v74_)
            d_509_v75_: _dafny.Set
            d_509_v75_ = _dafny.Set({d_461_v43_})
            d_509_v75_ = ((d_509_v75_).intersection(d_509_v75_)) - (d_509_v75_)
            d_510_v76_: _dafny.Seq
            d_510_v76_ = _dafny.SeqWithoutIsStrInference([(d_461_v43_)[default__.safeIndex(495, (d_461_v43_).length(0))], len(d_509_v75_), (d_461_v43_)[default__.safeIndex(495, (d_461_v43_).length(0))]])
            d_511_v77_: _dafny.Seq
            d_511_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "js"))
            d_512_v78_: _dafny.MultiSet
            d_512_v78_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixeqyo"))) <= (d_511_v77_), self.f16, (self.f16) and (self.f16), self.f16])
            index109_ = default__.safeIndex(495, (d_461_v43_).length(0))
            rhs118_ = d_508_v74_
            rhs119_ = d_510_v76_
            rhs120_ = (_dafny.MultiSet(d_508_v74_)) | (d_512_v78_)
            rhs121_ = (self).f30
            lhs109_ = d_461_v43_
            lhs110_ = default__.safeIndex(495, (d_461_v43_).length(0))
            d_508_v74_ = rhs118_
            d_510_v76_ = rhs119_
            d_512_v78_ = rhs120_
            lhs109_[lhs110_] = rhs121_
        elif True:
            d_513___mcc_h12_ = source6_.cf20
            d_514_cf20_ = d_513___mcc_h12_
            d_515_v79_: _dafny.Array
            nw99_ = _dafny.Array(_dafny.Seq({}), 26)
            d_515_v79_ = nw99_
            index110_ = default__.safeIndex(239, (d_515_v79_).length(0))
            (d_515_v79_)[index110_] = (self).f17
            index111_ = default__.safeIndex(239, (d_515_v79_).length(0))
            (d_515_v79_)[index111_] = ((self).f17) + ((self).f17)
            d_516_v80_: _dafny.Seq
            d_516_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lju"))
            d_516_v80_ = d_516_v80_
            d_517_v81_: D0
            d_517_v81_ = D0_DC0((d_516_v80_) + (d_516_v80_))
            d_517_v81_ = d_517_v81_
            d_518_v82_: _dafny.Set
            d_518_v82_ = _dafny.Set({False, default__.fm0(not(False), self.f16, globalState), True, self.f16, ((self).f28) != ((self).f28)})
            d_518_v82_ = d_518_v82_
        r1 = (self.f16 if self.f16 else self.f16)
        d_519_v83_: str
        d_519_v83_ = _dafny.CodePoint('c')
        r0 = d_519_v83_
        r1 = self.f16
        r2 = ((d_461_v43_)[default__.safeIndex(495, (d_461_v43_).length(0))]) - ((0) - (-271))
        return r0, r1, r2

    def m12(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: bool = False
        d_520_v0_: _dafny.Seq
        d_520_v0_ = _dafny.SeqWithoutIsStrInference([self.f16, self.f16])
        d_521_v1_: _dafny.Array
        nw100_ = _dafny.Array(None, 5)
        nw100_[int(0)] = False
        nw100_[int(1)] = (_dafny.MultiSet([self.f16])).ispropersubset(_dafny.MultiSet(d_520_v0_))
        nw100_[int(2)] = (self.f16 if self.f16 else not(True))
        nw100_[int(3)] = ((self).f29) <= ((self).f29)
        nw100_[int(4)] = not((len(_dafny.SeqWithoutIsStrInference([(self).f30]))) >= (len((self).f17)))
        d_521_v1_ = nw100_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_521_v1_).length(0)):
            d_522_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_522_i0_)) and ((d_522_i0_) < ((d_521_v1_).length(0)))):
                (d_521_v1_)[(d_522_i0_)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucfop"))).set(default__.safeIndex(len(_dafny.Map({_dafny.Map({(self).f29: self.f16}): _dafny.Set({(_dafny.MultiSet([self.f16])).cardinality})})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucfop")))), _dafny.CodePoint('b'))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpiib")))
        hi4_ = (self).f30
        for d_523_i1_ in range(((self).f28) * (566), hi4_):
            if default__.fm0((d_520_v0_) < (d_520_v0_), self.f16, globalState):
                d_524_v2_: int
                d_524_v2_ = 460
                d_524_v2_ = (self).f30
                d_525_v3_: _dafny.Seq
                d_525_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
                d_526_v4_: _dafny.Map
                d_526_v4_ = _dafny.Map({len(d_525_v3_): False})
                d_527_v5_: str
                d_527_v5_ = _dafny.CodePoint('p')
                d_528_v6_: _dafny.Map
                d_528_v6_ = _dafny.Map({d_526_v4_: d_527_v5_})
                d_528_v6_ = (d_528_v6_).set(d_526_v4_, d_527_v5_)
                (globalState).f5 = False
                d_524_v2_ = (0) - (default__.safeModuloInt((self).f28, (self).f28))
                d_529_v7_: D1
                d_529_v7_ = D1_DC4(self.f16, d_524_v2_, _dafny.Map({self.f16: self.f16}))
                rhs122_ = (d_529_v7_).cf8
                rhs123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdpvkliok"))
                rhs124_ = d_524_v2_
                d_524_v2_ = rhs122_
                d_525_v3_ = rhs123_
                d_524_v2_ = rhs124_
            elif True:
                d_530_v8_: _dafny.Map
                d_530_v8_ = _dafny.Map({self.f16: (0) - ((0) - ((self).f28))})
                d_531_v9_: _dafny.Array
                def lambda27_(d_532_i2_):
                    return (self).f17

                init14_ = lambda27_
                nw101_ = _dafny.Array(None, 23)
                for i0_14_ in range(nw101_.length(0)):
                    nw101_[i0_14_] = init14_(i0_14_)
                d_531_v9_ = nw101_
                pat_let_tv16_ = d_531_v9_
                def iife58_(_pat_let10_0):
                    def iife59_(d_533_dt__update__tmp_h0_):
                        def iife60_(_pat_let11_0):
                            def iife61_(d_534_dt__update_hcf14_h0_):
                                def iife62_(_pat_let12_0):
                                    def iife63_(d_535_dt__update_hcf16_h0_):
                                        return D2_DC8((d_533_dt__update__tmp_h0_).cf13, d_534_dt__update_hcf14_h0_, (d_533_dt__update__tmp_h0_).cf15, d_535_dt__update_hcf16_h0_, (d_533_dt__update__tmp_h0_).cf17)
                                    return iife63_(_pat_let12_0)
                                return iife62_(True)
                            return iife61_(_pat_let11_0)
                        return iife60_(pat_let_tv16_)
                    return iife59_(_pat_let10_0)
                d_530_v8_ = ((d_530_v8_) | (d_530_v8_)) | ((iife58_(D2_DC8(d_530_v8_, d_531_v9_, self.f16, self.f16, (self).f29))).cf13)
                d_536_v10_: int
                d_536_v10_ = 704
                d_537_v11_: str
                d_537_v11_ = _dafny.CodePoint('k')
                rhs125_ = d_523_i1_
                rhs126_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufg"))).set(default__.safeIndex(default__.safeDivisionInt(d_523_i1_, (self).f30), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufg")))), d_537_v11_)
                rhs127_ = (self.f16 if default__.fm0(self.f16, self.f16, globalState) else True)
                lhs111_ = globalState
                d_536_v10_ = rhs125_
                r0 = rhs126_
                lhs111_.f2 = rhs127_
                d_536_v10_ = ((self).f28) - ((self).f28)
                d_538_v12_: _dafny.Array
                def lambda28_(d_539_v11_):
                    def lambda29_(d_540_i3_):
                        return d_539_v11_

                    return lambda29_

                init15_ = lambda28_(d_537_v11_)
                nw102_ = _dafny.Array(None, 19)
                for i0_15_ in range(nw102_.length(0)):
                    nw102_[i0_15_] = init15_(i0_15_)
                d_538_v12_ = nw102_
                d_541_v13_: _dafny.Map
                d_541_v13_ = _dafny.Map({True: d_538_v12_})
                d_542_v14_: _dafny.Seq
                d_542_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fj"))
                d_543_v15_: _dafny.MultiSet
                d_543_v15_ = _dafny.MultiSet([d_536_v10_, (self).f30, len(d_542_v14_)])
                d_544_v16_: _dafny.Array
                nw103_ = _dafny.Array(None, 17)
                nw103_[int(0)] = default__.safeModuloInt((self).fm14((self).f29, (self).f29, _dafny.CodePoint('h'), globalState), -34)
                nw103_[int(1)] = (self).f30
                nw103_[int(2)] = default__.fm3(globalState)
                nw103_[int(3)] = d_523_i1_
                nw103_[int(4)] = d_536_v10_
                nw103_[int(5)] = ((self).f29) * (len(d_530_v8_))
                nw103_[int(6)] = (self).f28
                nw103_[int(7)] = 84
                nw103_[int(8)] = (self).f30
                nw103_[int(9)] = (self).f30
                nw103_[int(10)] = (self).f29
                nw103_[int(11)] = (self).f28
                nw103_[int(12)] = len(d_541_v13_)
                nw103_[int(13)] = ((d_543_v15_) - (d_543_v15_)).cardinality
                nw103_[int(14)] = default__.safeDivisionInt((0) - (d_536_v10_), 210)
                nw103_[int(15)] = len(d_542_v14_)
                nw103_[int(16)] = (self).f30
                d_544_v16_ = nw103_
                rhs128_ = d_544_v16_
                rhs129_ = self.f16
                rhs130_ = self.f16
                lhs112_ = globalState
                lhs113_ = globalState
                d_544_v16_ = rhs128_
                lhs112_.f2 = rhs129_
                lhs113_.f5 = rhs130_
                index112_ = default__.safeIndex(695, (d_521_v1_).length(0))
                (d_521_v1_)[index112_] = default__.fm0(self.f16, self.f16, globalState)
                index113_ = default__.safeIndex(695, (d_521_v1_).length(0))
                rhs131_ = not(self.f16)
                rhs132_ = self.f16
                rhs133_ = default__.fm0(((self).f28) != ((self).f30), (self.f16 if self.f16 else self.f16), globalState)
                rhs134_ = (d_520_v0_)[default__.safeIndex((default__.fm3(globalState) if self.f16 else (self).f28), len(d_520_v0_))]
                lhs114_ = d_521_v1_
                lhs115_ = default__.safeIndex(695, (d_521_v1_).length(0))
                lhs116_ = globalState
                r2 = rhs131_
                lhs114_[lhs115_] = rhs132_
                lhs116_.f5 = rhs133_
                r2 = rhs134_
            (globalState).f2 = self.f16
            d_545_v17_: C0
            nw104_ = C0()
            nw104_.ctor__((len(_dafny.Map({(self).f30: self.f16}))) * ((self).f28), d_521_v1_)
            d_545_v17_ = nw104_
            d_546_v18_: _dafny.Map
            d_546_v18_ = _dafny.Map({d_521_v1_: d_521_v1_})
            d_547_v19_: _dafny.Array
            nw105_ = _dafny.Array(None, 12)
            nw105_[int(0)] = d_545_v17_.f27
            nw105_[int(1)] = d_521_v1_
            nw105_[int(2)] = d_545_v17_.f27
            nw105_[int(3)] = d_521_v1_
            nw105_[int(4)] = d_521_v1_
            nw105_[int(5)] = d_545_v17_.f27
            nw105_[int(6)] = (d_521_v1_ if self.f16 else d_545_v17_.f27)
            nw105_[int(7)] = d_521_v1_
            nw105_[int(8)] = d_545_v17_.f27
            nw105_[int(9)] = ((d_546_v18_)[d_545_v17_.f27] if (d_545_v17_.f27) in (d_546_v18_) else d_545_v17_.f27)
            nw105_[int(10)] = d_521_v1_
            nw105_[int(11)] = d_521_v1_
            d_547_v19_ = nw105_
            nw106_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_547_v19_ = nw106_
        d_548_v20_: int
        d_548_v20_ = 354
        d_548_v20_ = (self).f28
        d_549_v21_: _dafny.Set
        d_549_v21_ = _dafny.Set({d_521_v1_})
        d_550_v22_: _dafny.Set
        d_550_v22_ = _dafny.Set({self.f16, self.f16})
        d_551_v23_: _dafny.Map
        d_551_v23_ = _dafny.Map({self.f16: d_550_v22_})
        r2 = (self.f16 if (d_549_v21_).ispropersubset(d_549_v21_) else (d_551_v23_) == (d_551_v23_))
        d_548_v20_ = 497
        d_552_v24_: _dafny.Seq
        d_552_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldcs"))
        d_553_v25_: _dafny.Map
        d_553_v25_ = _dafny.Map({len(d_552_v24_): self.f16})
        (globalState).f8 = (d_553_v25_) | (_dafny.Map({(self).f30: not(self.f16)}))
        r0 = default__.fm1(globalState)
        d_554_v26_: _dafny.MultiSet
        d_554_v26_ = _dafny.MultiSet([(self).f28])
        r1 = ((d_554_v26_) - (default__.fm19((self).f28, globalState))).isdisjoint(d_554_v26_)
        r2 = self.f16
        return r0, r1, r2

    def m13(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_555_v0_: _dafny.Array
        def lambda30_(d_556_p2_):
            def lambda31_(d_557_i0_):
                return d_556_p2_

            return lambda31_

        init16_ = lambda30_(p2)
        nw107_ = _dafny.Array(None, 25)
        for i0_16_ in range(nw107_.length(0)):
            nw107_[i0_16_] = init16_(i0_16_)
        d_555_v0_ = nw107_
        d_558_v1_: C0
        nw108_ = C0()
        nw108_.ctor__(876, d_555_v0_)
        d_558_v1_ = nw108_
        d_559_v2_: _dafny.Array
        nw109_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_559_v2_ = nw109_
        d_559_v2_ = d_559_v2_
        d_560_i1_: int
        d_560_i1_ = 0
        with _dafny.label("2"):
            while (not(p2) if self.f16 else ((self).f28) <= ((self).f30)):
                with _dafny.c_label("2"):
                    if (d_560_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_560_i1_ = (d_560_i1_) + (1)
                    r0 = (self).f28
                    r0 = 965
                    (globalState).f2 = (p2 if self.f16 else self.f16)
                    d_561_v3_: _dafny.Set
                    d_561_v3_ = _dafny.Set({p1, self.f16, False, p2, self.f16})
                    d_562_v4_: D2
                    d_562_v4_ = D2_DC7((_dafny.Set({self.f16})) | (d_561_v3_))
                    d_562_v4_ = d_562_v4_
                    pass
            pass
        d_563_v5_: _dafny.Set
        d_563_v5_ = _dafny.Set({p1, p2})
        d_564_v6_: _dafny.Seq
        d_564_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "curdxgmyi"))
        d_565_v7_: _dafny.Map
        d_565_v7_ = _dafny.Map({len(d_563_v5_): len(d_564_v6_)})
        d_566_v8_: _dafny.Seq
        d_566_v8_ = _dafny.SeqWithoutIsStrInference([d_565_v7_, d_565_v7_])
        d_567_v10_: str
        d_567_v10_ = _dafny.CodePoint('p')
        d_568_v11_: _dafny.Map
        d_568_v11_ = _dafny.Map({d_567_v10_: p1})
        d_569_v13_: _dafny.MultiSet
        def iife64_():
            coll38_ = _dafny.Map()
            compr_38_: int
            for compr_38_ in _dafny.IntegerRange(357, 697):
                d_570_v9_: int = compr_38_
                if ((357) <= (d_570_v9_)) and ((d_570_v9_) < (697)):
                    coll38_[(d_570_v9_) + (196)] = (d_558_v1_).f26
            return _dafny.Map(coll38_)
        def iife65_():
            coll39_ = _dafny.Map()
            compr_39_: int
            for compr_39_ in _dafny.IntegerRange(983, 580):
                d_571_v12_: int = compr_39_
                if ((983) <= (d_571_v12_)) and ((d_571_v12_) < (580)):
                    coll39_[(d_571_v12_) * (323)] = (d_558_v1_).f26
            return _dafny.Map(coll39_)
        d_569_v13_ = _dafny.MultiSet([(d_566_v8_)[default__.safeIndex((self).f28, len(d_566_v8_))], (iife64_()
        ).set((self).f30, len(d_568_v11_)), _dafny.Map({(self).f28: (d_558_v1_).f26}), iife65_()
        , d_565_v7_])
        (globalState).f5 = (_dafny.MultiSet((d_566_v8_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f30: -324}), d_565_v7_, d_565_v7_])))).isdisjoint(d_569_v13_)
        d_572_v15_: _dafny.Array
        nw110_ = _dafny.Array(None, 17)
        nw110_[int(0)] = (self).f28
        nw110_[int(1)] = (self).f30
        nw110_[int(2)] = (self).f30
        nw110_[int(3)] = default__.safeModuloInt((d_558_v1_).f26, (self).f28)
        nw110_[int(4)] = (d_558_v1_).f26
        nw110_[int(5)] = (d_558_v1_).f26
        nw110_[int(6)] = ((self).f28) + ((0) - ((d_558_v1_).f26))
        nw110_[int(7)] = ((d_558_v1_).f26 if p1 else 700)
        nw110_[int(8)] = (self).f30
        nw110_[int(9)] = -614
        nw110_[int(10)] = (self).f28
        nw110_[int(11)] = (self).fm13(self.f16, globalState)
        nw110_[int(12)] = (self).f30
        nw110_[int(13)] = (self).f28
        def iife66_():
            coll40_ = _dafny.Set()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(965, -1):
                d_573_v14_: int = compr_40_
                if ((965) <= (d_573_v14_)) and ((d_573_v14_) < (-1)):
                    coll40_ = coll40_.union(_dafny.Set([(d_573_v14_) - ((d_558_v1_).f26)]))
            return _dafny.Set(coll40_)
        nw110_[int(14)] = default__.safeModuloInt(len(iife66_()
        ), 40)
        nw110_[int(15)] = (self).f28
        nw110_[int(16)] = (self).f28
        d_572_v15_ = nw110_
        index114_ = default__.safeIndex(288, (d_572_v15_).length(0))
        (d_572_v15_)[index114_] = ((0) - (len(d_564_v6_))) + ((self).f29)
        d_574_v16_: _dafny.Array
        def lambda32_(d_575_v1_):
            def lambda33_(d_576_i2_):
                return _dafny.MultiSet([_dafny.MultiSet([(d_575_v1_).f26])])

            return lambda33_

        init17_ = lambda32_(d_558_v1_)
        nw111_ = _dafny.Array(None, 2)
        for i0_17_ in range(nw111_.length(0)):
            nw111_[i0_17_] = init17_(i0_17_)
        d_574_v16_ = nw111_
        d_577_v17_: _dafny.MultiSet
        d_577_v17_ = _dafny.MultiSet([(self).f29])
        d_578_v18_: _dafny.MultiSet
        d_578_v18_ = _dafny.MultiSet([d_577_v17_, d_577_v17_])
        index115_ = default__.safeIndex(969, (d_574_v16_).length(0))
        (d_574_v16_)[index115_] = (d_578_v18_) | (d_578_v18_)
        d_579_v19_: _dafny.Array
        nw112_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_579_v19_ = nw112_
        d_580_v20_: _dafny.Array
        nw113_ = _dafny.Array(None, 1)
        nw113_[int(0)] = d_567_v10_
        d_580_v20_ = nw113_
        index116_ = default__.safeIndex(262, (d_579_v19_).length(0))
        (d_579_v19_)[index116_] = d_580_v20_
        d_581_v21_: _dafny.Seq
        d_581_v21_ = _dafny.SeqWithoutIsStrInference([d_577_v17_, d_577_v17_])
        d_582_v22_: _dafny.Seq
        d_582_v22_ = _dafny.SeqWithoutIsStrInference([p2])
        index117_ = default__.safeIndex(288, (d_572_v15_).length(0))
        index118_ = default__.safeIndex(969, (d_574_v16_).length(0))
        index119_ = default__.safeIndex(262, (d_579_v19_).length(0))
        rhs135_ = -533
        rhs136_ = (self).f29
        rhs137_ = (_dafny.MultiSet((d_581_v21_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([463, (0) - ((d_558_v1_).f26)])])))) - (d_578_v18_)
        rhs138_ = len((d_582_v22_) + (default__.fm20(default__.fm0(p1, p1, globalState), self.f16, p1, default__.fm3(globalState), globalState)))
        rhs139_ = (d_580_v20_ if p2 else d_580_v20_)
        lhs117_ = d_572_v15_
        lhs118_ = default__.safeIndex(288, (d_572_v15_).length(0))
        lhs119_ = d_574_v16_
        lhs120_ = default__.safeIndex(969, (d_574_v16_).length(0))
        lhs121_ = d_579_v19_
        lhs122_ = default__.safeIndex(262, (d_579_v19_).length(0))
        lhs117_[lhs118_] = rhs135_
        r1 = rhs136_
        lhs119_[lhs120_] = rhs137_
        r0 = rhs138_
        lhs121_[lhs122_] = rhs139_
        d_583_v23_: _dafny.Map
        d_583_v23_ = _dafny.Map({default__.fm3(globalState): p1})
        d_584_v24_: C0
        nw114_ = C0()
        nw114_.ctor__((len(d_583_v23_) if p1 else (d_558_v1_).f26), d_558_v1_.f27)
        d_584_v24_ = nw114_
        r0 = (((0) - ((self).f28)) + ((self).f28) if not(p1) else len((self).f17))
        d_585_v25_: _dafny.MultiSet
        d_585_v25_ = _dafny.MultiSet([d_582_v22_])
        d_586_v26_: _dafny.MultiSet
        d_586_v26_ = _dafny.MultiSet([(self).f17])
        d_587_v27_: _dafny.Seq
        d_587_v27_ = _dafny.SeqWithoutIsStrInference([(self).f17])
        r1 = ((d_585_v25_).set(d_582_v22_, default__.abs(((d_586_v26_)[(d_587_v27_)[default__.safeIndex(((self).f17)[default__.safeIndex((d_558_v1_).f26, len((self).f17))], len(d_587_v27_))]] if ((d_587_v27_)[default__.safeIndex(((self).f17)[default__.safeIndex((d_558_v1_).f26, len((self).f17))], len(d_587_v27_))]) in (d_586_v26_) else (self).fm15((self).f17, len(d_582_v22_), len(d_564_v6_), _dafny.CodePoint('i'), globalState))))).cardinality
        return r0, r1

    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30

class C2:
    def  __init__(self):
        self._f25: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f25):
        (self)._f25 = f25

    def fm12(self, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([True, False])})) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([not(False), False]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(True), True, False, True, False]), _dafny.SeqWithoutIsStrInference([False, True])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([True, False])})))

    def m9(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_588_v0_: _dafny.Set
        d_588_v0_ = _dafny.Set({(self).f25})
        d_589_i0_: int
        d_589_i0_ = 0
        with _dafny.label("3"):
            while not((default__.safeDivisionInt((self).f25, 873)) >= (((self).f25) * (len(d_588_v0_)))):
                with _dafny.c_label("3"):
                    if (d_589_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_589_i0_ = (d_589_i0_) + (1)
                    d_590_v1_: _dafny.MultiSet
                    d_590_v1_ = _dafny.MultiSet([(389) * ((self).f25), ((self).f25) + ((self).f25)])
                    d_591_v2_: str
                    d_591_v2_ = _dafny.CodePoint('r')
                    d_592_v3_: _dafny.MultiSet
                    d_592_v3_ = _dafny.MultiSet([d_591_v2_, d_591_v2_, d_591_v2_])
                    d_593_v4_: bool
                    d_593_v4_ = False
                    d_594_v5_: _dafny.Set
                    d_594_v5_ = _dafny.Set({d_593_v4_})
                    d_590_v1_ = (_dafny.MultiSet([(self).f25, ((d_592_v3_).set(_dafny.CodePoint('c'), default__.abs(len(d_594_v5_)))).cardinality, (self).f25])).intersection((d_590_v1_) - (p2))
                    d_595_v6_: _dafny.Map
                    d_595_v6_ = _dafny.Map({d_593_v4_: (self).f25})
                    d_596_v7_: _dafny.MultiSet
                    d_596_v7_ = _dafny.MultiSet([(d_595_v6_).set(d_593_v4_, (self).f25)])
                    d_597_v8_: D4
                    d_597_v8_ = D4_DC11((d_596_v7_) - (d_596_v7_))
                    source7_ = d_597_v8_
                    if source7_.is_DC12:
                        d_598___mcc_h0_ = source7_.cf21
                        d_599___mcc_h1_ = source7_.cf22
                        d_600___mcc_h2_ = source7_.cf23
                        d_601_cf23_ = d_600___mcc_h2_
                        d_602_cf22_ = d_599___mcc_h1_
                        d_603_cf21_ = d_598___mcc_h0_
                        d_604_v9_: _dafny.Array
                        nw115_ = _dafny.Array(False, 5)
                        d_604_v9_ = nw115_
                        d_605_v10_: C0
                        nw116_ = C0()
                        nw116_.ctor__(default__.safeModuloInt((self).f25, (self).f25), d_604_v9_)
                        d_605_v10_ = nw116_
                        (globalState).f5 = d_593_v4_
                        d_606_v11_: _dafny.Seq
                        d_606_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlvehw"))
                        d_607_v12_: _dafny.Seq
                        d_607_v12_ = _dafny.SeqWithoutIsStrInference([False, d_593_v4_, d_603_cf21_, d_602_cf22_])
                        rhs140_ = (d_607_v12_)[default__.safeIndex(d_601_cf23_, len(d_607_v12_))]
                        rhs141_ = (d_606_v11_) + (default__.fm1(globalState))
                        lhs123_ = globalState
                        lhs123_.f5 = rhs140_
                        d_606_v11_ = rhs141_
                        d_608_v13_: _dafny.Array
                        def lambda34_(d_609_i1_):
                            return (d_609_i1_) + ((0) - ((self).f25))

                        init18_ = lambda34_
                        nw117_ = _dafny.Array(None, 6)
                        for i0_18_ in range(nw117_.length(0)):
                            nw117_[i0_18_] = init18_(i0_18_)
                        d_608_v13_ = nw117_
                        d_610_v14_: _dafny.MultiSet
                        d_610_v14_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_602_cf22_]), d_607_v12_])
                        d_611_v15_: _dafny.Map
                        d_611_v15_ = _dafny.Map({((d_605_v10_).f26) + ((d_610_v14_).cardinality): d_608_v13_})
                        d_612_v16_: _dafny.Map
                        d_612_v16_ = _dafny.Map({d_593_v4_: p0})
                        d_608_v13_ = ((d_611_v15_)[(0) - (((self).f25) - (d_601_cf23_))] if ((0) - (((self).f25) - (d_601_cf23_))) in (d_611_v15_) else ((d_612_v16_)[False] if (False) in (d_612_v16_) else p0))
                    elif source7_.is_DC13:
                        d_613___mcc_h3_ = source7_.cf24
                        d_614___mcc_h4_ = source7_.cf25
                        d_615___mcc_h5_ = source7_.cf26
                        d_616_cf26_ = d_615___mcc_h5_
                        d_617_cf25_ = d_614___mcc_h4_
                        d_618_cf24_ = d_613___mcc_h3_
                        d_619_v17_: _dafny.Array
                        nw118_ = _dafny.Array(None, 2)
                        nw118_[int(0)] = d_593_v4_
                        nw118_[int(1)] = True
                        d_619_v17_ = nw118_
                        d_620_v18_: D0
                        d_620_v18_ = D0_DC2((0) - ((self).f25), (self).f25, d_616_cf26_, d_619_v17_)
                        pat_let_tv17_ = d_617_cf25_
                        pat_let_tv18_ = d_618_cf24_
                        def iife67_(_pat_let13_0):
                            def iife68_(d_621_dt__update__tmp_h0_):
                                def iife69_(_pat_let14_0):
                                    def iife70_(d_622_dt__update_hcf4_h0_):
                                        def iife71_(_pat_let15_0):
                                            def iife72_(d_623_dt__update_hcf2_h0_):
                                                return D0_DC2(d_623_dt__update_hcf2_h0_, (d_621_dt__update__tmp_h0_).cf3, d_622_dt__update_hcf4_h0_, (d_621_dt__update__tmp_h0_).cf5)
                                            return iife72_(_pat_let15_0)
                                        return iife71_(pat_let_tv18_)
                                    return iife70_(_pat_let14_0)
                                return iife69_(pat_let_tv17_)
                            return iife68_(_pat_let13_0)
                        d_617_cf25_ = (iife67_(d_620_v18_)).cf4
                        (globalState).f2 = d_593_v4_
                        d_624_v19_: _dafny.Map
                        d_624_v19_ = _dafny.Map({_dafny.MultiSet([d_616_cf26_]): True})
                        (globalState).f5 = ((d_624_v19_)[p2] if (p2) in (d_624_v19_) else (d_617_cf25_) != (543))
                        d_625_v20_: _dafny.Seq
                        d_625_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsh"))
                        d_626_v21_: D4
                        d_626_v21_ = D4_DC12(d_593_v4_, True, len(d_625_v20_))
                        rhs142_ = d_591_v2_
                        rhs143_ = d_626_v21_
                        lhs124_ = globalState
                        lhs124_.f15 = rhs142_
                        d_626_v21_ = rhs143_
                    elif True:
                        d_627___mcc_h6_ = source7_.cf20
                        d_628_cf20_ = d_627___mcc_h6_
                        r0 = (self).f25
                        index120_ = default__.safeIndex(372, (p0).length(0))
                        (p0)[index120_] = (self).f25
                        d_629_v22_: _dafny.Seq
                        d_629_v22_ = _dafny.SeqWithoutIsStrInference([d_593_v4_])
                        index121_ = default__.safeIndex(372, (p0).length(0))
                        (p0)[index121_] = (0) - ((0) - (((self).f25) - (((d_590_v1_)[(self).f25] if ((self).f25) in (d_590_v1_) else len(d_629_v22_)))))
                        d_630_v23_: _dafny.Array
                        nw119_ = _dafny.Array(D0.default()(), 23)
                        d_630_v23_ = nw119_
                        d_631_v24_: _dafny.Map
                        d_631_v24_ = _dafny.Map({(0) - ((p0)[default__.safeIndex(372, (p0).length(0))]): (self).f25})
                        d_632_v25_: _dafny.Map
                        d_632_v25_ = _dafny.Map({d_631_v24_: d_593_v4_})
                        d_633_v26_: _dafny.Map
                        d_633_v26_ = _dafny.Map({d_630_v23_: ((d_632_v25_)[d_631_v24_] if (d_631_v24_) in (d_632_v25_) else d_593_v4_)})
                        d_633_v26_ = (d_633_v26_).set(d_630_v23_, d_593_v4_)
                        d_634_v27_: _dafny.Array
                        nw120_ = _dafny.Array(_dafny.Array(None, 0), 24)
                        d_634_v27_ = nw120_
                        d_635_v28_: _dafny.Map
                        d_635_v28_ = _dafny.Map({d_593_v4_: d_593_v4_})
                        d_636_v29_: _dafny.Array
                        nw121_ = _dafny.Array(None, 28)
                        nw121_[int(0)] = ((d_635_v28_)[d_593_v4_] if (d_593_v4_) in (d_635_v28_) else d_593_v4_)
                        nw121_[int(1)] = d_593_v4_
                        nw121_[int(2)] = False
                        nw121_[int(3)] = d_593_v4_
                        nw121_[int(4)] = d_593_v4_
                        nw121_[int(5)] = d_593_v4_
                        nw121_[int(6)] = d_593_v4_
                        nw121_[int(7)] = d_593_v4_
                        nw121_[int(8)] = d_593_v4_
                        nw121_[int(9)] = d_593_v4_
                        nw121_[int(10)] = d_593_v4_
                        nw121_[int(11)] = d_593_v4_
                        nw121_[int(12)] = d_593_v4_
                        nw121_[int(13)] = d_593_v4_
                        nw121_[int(14)] = d_593_v4_
                        nw121_[int(15)] = d_593_v4_
                        nw121_[int(16)] = d_593_v4_
                        nw121_[int(17)] = d_593_v4_
                        nw121_[int(18)] = d_593_v4_
                        nw121_[int(19)] = not(not(d_593_v4_))
                        nw121_[int(20)] = d_593_v4_
                        nw121_[int(21)] = d_593_v4_
                        nw121_[int(22)] = d_593_v4_
                        nw121_[int(23)] = d_593_v4_
                        nw121_[int(24)] = d_593_v4_
                        nw121_[int(25)] = d_593_v4_
                        nw121_[int(26)] = True
                        nw121_[int(27)] = d_593_v4_
                        d_636_v29_ = nw121_
                        index122_ = default__.safeIndex(316, (d_634_v27_).length(0))
                        (d_634_v27_)[index122_] = d_636_v29_
                        index123_ = default__.safeIndex(316, (d_634_v27_).length(0))
                        (d_634_v27_)[index123_] = d_636_v29_
                    d_637_v30_: _dafny.Array
                    nw122_ = _dafny.Array(False, 4)
                    d_637_v30_ = nw122_
                    index124_ = default__.safeIndex(77, (d_637_v30_).length(0))
                    (d_637_v30_)[index124_] = d_593_v4_
                    index125_ = default__.safeIndex(77, (d_637_v30_).length(0))
                    (d_637_v30_)[index125_] = d_593_v4_
                    if not (not(d_593_v4_)) or ((default__.fm3(globalState)) < ((self).f25)):
                        d_638_v31_: D3
                        d_638_v31_ = D3_DC10((d_637_v30_)[default__.safeIndex(77, (d_637_v30_).length(0))])
                        (globalState).f5 = (d_638_v31_).cf19
                        (globalState).f2 = False
                        d_639_v32_: _dafny.Map
                        d_639_v32_ = _dafny.Map({d_593_v4_: d_594_v5_})
                        d_594_v5_ = ((((d_639_v32_)[(d_637_v30_)[default__.safeIndex(77, (d_637_v30_).length(0))]] if ((d_637_v30_)[default__.safeIndex(77, (d_637_v30_).length(0))]) in (d_639_v32_) else d_594_v5_)) | (d_594_v5_)).intersection(_dafny.Set({False}))
                        (globalState).f5 = (d_637_v30_)[default__.safeIndex(77, (d_637_v30_).length(0))]
                        index126_ = default__.safeIndex(77, (d_637_v30_).length(0))
                        (d_637_v30_)[index126_] = ((self).f25) == (((0) - ((self).f25)) + ((0) - ((p2).cardinality)))
                    elif True:
                        d_640_v34_: _dafny.Map
                        d_640_v34_ = _dafny.Map({d_595_v6_: (self).f25})
                        d_641_v35_: C0
                        nw123_ = C0()
                        def iife73_():
                            coll41_ = _dafny.Map()
                            compr_41_: _dafny.Map
                            for compr_41_ in (d_640_v34_).keys.Elements:
                                d_642_v33_: _dafny.Map = compr_41_
                                if (d_642_v33_) in (d_640_v34_):
                                    coll41_[d_642_v33_] = (self).f25
                            return _dafny.Map(coll41_)
                        nw123_.ctor__(len(iife73_()
                        ), (d_637_v30_ if d_593_v4_ else d_637_v30_))
                        d_641_v35_ = nw123_
                        index127_ = default__.safeIndex(77, (d_637_v30_).length(0))
                        rhs144_ = (d_641_v35_).f26
                        rhs145_ = (d_637_v30_)[default__.safeIndex(77, (d_637_v30_).length(0))]
                        lhs125_ = d_637_v30_
                        lhs126_ = default__.safeIndex(77, (d_637_v30_).length(0))
                        r0 = rhs144_
                        lhs125_[lhs126_] = rhs145_
                        r0 = ((d_641_v35_).f26) * ((d_641_v35_).f26)
                        d_643_v36_: _dafny.Map
                        d_643_v36_ = _dafny.Map({d_591_v2_: D1_DC6()})
                        d_644_v37_: D1
                        d_644_v37_ = D1_DC6()
                        d_643_v36_ = (d_643_v36_).set(d_591_v2_, d_644_v37_)
                        d_637_v30_ = d_641_v35_.f27
                    pass
            pass
        d_645_v38_: bool
        d_645_v38_ = False
        d_646_i2_: int
        d_646_i2_ = 0
        with _dafny.label("4"):
            while d_645_v38_:
                with _dafny.c_label("4"):
                    if (d_646_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_646_i2_ = (d_646_i2_) + (1)
                    d_647_v39_: _dafny.Array
                    nw124_ = _dafny.Array(D0.default()(), 22)
                    d_647_v39_ = nw124_
                    d_648_v40_: _dafny.Array
                    nw125_ = _dafny.Array(False, 16)
                    d_648_v40_ = nw125_
                    index128_ = default__.safeIndex(67, (d_647_v39_).length(0))
                    (d_647_v39_)[index128_] = D0_DC2((self).f25, (self).f25, (self).f25, d_648_v40_)
                    d_649_v41_: D0
                    d_649_v41_ = D0_DC2((self).f25, ((p2)[(self).f25] if ((self).f25) in (p2) else (self).f25), ((self).f25) + (730), d_648_v40_)
                    index129_ = default__.safeIndex(67, (d_647_v39_).length(0))
                    rhs146_ = d_649_v41_
                    rhs147_ = -984
                    lhs127_ = d_647_v39_
                    lhs128_ = default__.safeIndex(67, (d_647_v39_).length(0))
                    lhs127_[lhs128_] = rhs146_
                    r0 = rhs147_
                    index130_ = default__.safeIndex(446, (p0).length(0))
                    (p0)[index130_] = default__.fm3(globalState)
                    index131_ = default__.safeIndex(446, (p0).length(0))
                    (p0)[index131_] = -683
                    d_650_v42_: _dafny.MultiSet
                    d_650_v42_ = _dafny.MultiSet([d_645_v38_])
                    d_651_v43_: D1
                    d_651_v43_ = D1_DC3(d_650_v42_)
                    d_651_v43_ = d_651_v43_
                    if d_645_v38_:
                        d_652_v44_: _dafny.Array
                        def lambda35_(d_653_p0_):
                            def lambda36_(d_654_i3_):
                                return (d_654_i3_) + ((d_653_p0_)[default__.safeIndex(446, (d_653_p0_).length(0))])

                            return lambda36_

                        init19_ = lambda35_(p0)
                        nw126_ = _dafny.Array(None, 23)
                        for i0_19_ in range(nw126_.length(0)):
                            nw126_[i0_19_] = init19_(i0_19_)
                        d_652_v44_ = nw126_
                        d_652_v44_ = d_652_v44_
                        index132_ = default__.safeIndex(446, (p0).length(0))
                        (p0)[index132_] = (((p2).cardinality) * (257)) + ((p0)[default__.safeIndex(446, (p0).length(0))])
                        d_655_v45_: _dafny.Array
                        nw127_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                        d_655_v45_ = nw127_
                        index133_ = default__.safeIndex(467, (d_655_v45_).length(0))
                        (d_655_v45_)[index133_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_656_i4_ in range(default__.abs(372))])
                        d_657_v46_: _dafny.Seq
                        d_657_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmr"))
                        index134_ = default__.safeIndex(467, (d_655_v45_).length(0))
                        (d_655_v45_)[index134_] = d_657_v46_
                        d_658_v47_: _dafny.Seq
                        d_658_v47_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                        d_659_v48_: T0
                        nw128_ = C1()
                        nw128_.ctor__((self).f25, (p0)[default__.safeIndex(446, (p0).length(0))], d_645_v38_, d_658_v47_, (p0)[default__.safeIndex(446, (p0).length(0))])
                        d_659_v48_ = nw128_
                        d_660_v49_: _dafny.Map
                        d_660_v49_ = _dafny.Map({not(d_645_v38_): d_659_v48_})
                        rhs148_ = d_645_v38_
                        rhs149_ = d_660_v49_
                        lhs129_ = globalState
                        lhs129_.f2 = rhs148_
                        d_660_v49_ = rhs149_
                        r0 = (((d_650_v42_)[d_659_v48_.f16] if (d_659_v48_.f16) in (d_650_v42_) else default__.fm3(globalState)) if d_659_v48_.f16 else ((p0)[default__.safeIndex(446, (p0).length(0))]) * ((p0)[default__.safeIndex(446, (p0).length(0))]))
                    elif True:
                        d_645_v38_ = not(d_645_v38_)
                        d_661_v50_: _dafny.Seq
                        d_661_v50_ = _dafny.SeqWithoutIsStrInference([not(d_645_v38_)])
                        index135_ = default__.safeIndex(446, (p0).length(0))
                        (p0)[index135_] = default__.safeDivisionInt((self).f25, len(d_661_v50_))
                        d_662_v51_: _dafny.Map
                        d_662_v51_ = _dafny.Map({d_645_v38_: d_645_v38_})
                        d_663_v52_: D1
                        d_663_v52_ = D1_DC4(not(d_645_v38_), (self).f25, d_662_v51_)
                        d_664_v53_: _dafny.Map
                        d_664_v53_ = _dafny.Map({d_663_v52_: d_588_v0_})
                        d_664_v53_ = (d_664_v53_).set(d_663_v52_, _dafny.Set({(self).f25}))
                        d_665_v54_: _dafny.Array
                        nw129_ = _dafny.Array(_dafny.Array(None, 0), 24)
                        d_665_v54_ = nw129_
                        d_666_v55_: D6
                        d_666_v55_ = D6_DC15(d_665_v54_)
                        d_665_v54_ = (d_666_v55_).cf28
                        index136_ = default__.safeIndex(446, (p0).length(0))
                        (p0)[index136_] = (self).f25
                    pass
            pass
        r0 = ((self).f25 if d_645_v38_ else default__.safeDivisionInt((self).f25, (self).f25))
        d_667_v56_: _dafny.Array
        nw130_ = _dafny.Array(False, 13)
        d_667_v56_ = nw130_
        d_668_v57_: C0
        nw131_ = C0()
        nw131_.ctor__((self).f25, d_667_v56_)
        d_668_v57_ = nw131_
        index137_ = default__.safeIndex(641, (d_667_v56_).length(0))
        (d_667_v56_)[index137_] = d_645_v38_
        index138_ = default__.safeIndex(641, (d_667_v56_).length(0))
        (d_667_v56_)[index138_] = (((self).f25) * (-768)) <= ((self).f25)
        d_669_v58_: str
        d_669_v58_ = _dafny.CodePoint('a')
        (globalState).f15 = d_669_v58_
        d_670_v59_: _dafny.Array
        nw132_ = _dafny.Array(D2.default()(), 20)
        d_670_v59_ = nw132_
        d_671_v60_: _dafny.Map
        d_671_v60_ = _dafny.Map({d_670_v59_: (d_667_v56_)[default__.safeIndex(641, (d_667_v56_).length(0))]})
        d_672_v61_: _dafny.Array
        d_672_v61_ = d_670_v59_
        r0 = len(_dafny.Set({(self).f25, len((d_671_v60_) | ((d_671_v60_).set(d_672_v61_, (d_667_v56_)[default__.safeIndex(641, (d_667_v56_).length(0))])))}))
        return r0

    def m10(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_673_v0_: str
        d_673_v0_ = _dafny.CodePoint('k')
        (globalState).f15 = d_673_v0_
        r0 = default__.fm3(globalState)
        d_674_v1_: D6
        d_674_v1_ = D6_DC16(not(not(p0)), p0)
        def lambda37_(source8_):
            if source8_.is_DC16:
                d_675___mcc_h0_ = source8_.cf29
                d_676___mcc_h1_ = source8_.cf30
                d_677_cf30_ = d_676___mcc_h1_
                d_678_cf29_ = d_675___mcc_h0_
                return (self).f25
            elif source8_.is_DC17:
                d_679___mcc_h2_ = source8_.cf31
                d_680_cf31_ = d_679___mcc_h2_
                return (self).f25
            elif source8_.is_DC18:
                d_681___mcc_h3_ = source8_.cf32
                d_682_cf32_ = d_681___mcc_h3_
                return ((self).f25) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_683_i0_ in range(default__.abs(566))])))
            elif True:
                d_684___mcc_h4_ = source8_.cf28
                d_685_cf28_ = d_684___mcc_h4_
                return default__.safeModuloInt(584, (0) - ((0) - ((self).f25)))

        r0 = lambda37_((d_674_v1_ if p0 else d_674_v1_))
        hi5_ = (self).f25
        for d_686_i1_ in range((self).f25, hi5_):
            d_687_v2_: _dafny.Seq
            d_687_v2_ = _dafny.SeqWithoutIsStrInference([459])
            d_688_v3_: C1
            nw133_ = C1()
            nw133_.ctor__((self).f25, (self).f25, p0, d_687_v2_, d_686_i1_)
            d_688_v3_ = nw133_
            (globalState).f2 = False
            d_689_v4_: _dafny.Seq
            d_689_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbabos"))
            d_690_v5_: _dafny.Map
            d_690_v5_ = _dafny.Map({p0: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0])): d_689_v4_}))})
            d_691_v6_: _dafny.Array
            nw134_ = _dafny.Array(_dafny.Seq({}), 20)
            d_691_v6_ = nw134_
            d_692_v7_: _dafny.Set
            d_692_v7_ = _dafny.Set({(d_688_v3_).f30, (self).f25})
            d_693_v8_: _dafny.Map
            d_693_v8_ = _dafny.Map({d_688_v3_: p0})
            d_694_v9_: D1
            d_694_v9_ = D1_DC5(len(_dafny.Map({len(d_692_v7_): 99})), ((d_693_v8_)[d_688_v3_] if (d_688_v3_) in (d_693_v8_) else p0))
            d_695_v10_: D2
            d_695_v10_ = D2_DC8(d_690_v5_, d_691_v6_, p0, (d_694_v9_).cf11, (0) - ((self).f25))
            d_696_v11_: _dafny.MultiSet
            d_696_v11_ = _dafny.MultiSet([((d_688_v3_).f30) + ((d_688_v3_).f29), (d_688_v3_).fm14((self).f25, (d_688_v3_).f30, d_673_v0_, globalState), (d_695_v10_).cf17])
            d_696_v11_ = d_696_v11_
            d_697_v12_: D1
            d_697_v12_ = D1_DC6()
            d_697_v12_ = d_697_v12_
        d_698_v13_: _dafny.Seq
        d_698_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lskhib"))
        d_698_v13_ = (d_698_v13_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otcfiuf")))
        r0 = (len(d_698_v13_)) * ((self).f25)
        r0 = (self).f25
        d_699_v14_: _dafny.Array
        nw135_ = _dafny.Array(int(0), 25)
        d_699_v14_ = nw135_
        r1 = (d_699_v14_ if (p0) and (not(p0)) else d_699_v14_)
        return r0, r1

    @property
    def f25(self):
        return self._f25

class C3(T0):
    def  __init__(self):
        self._f16: bool = False
        self._f17: _dafny.Seq = _dafny.Seq({})
        self._f24: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f24, f16, f17):
        (self)._f24 = f24
        (self).f16 = f16
        (self)._f17 = f17

    def fm11(self, p0, p1, p2, p3, globalState):
        return (self).f24

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        d_700_v1_: _dafny.Array
        def lambda38_(d_701_i0_):
            def iife74_():
                coll42_ = _dafny.Set()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(757, 81):
                    d_702_v0_: int = compr_42_
                    if ((757) <= (d_702_v0_)) and ((d_702_v0_) < (81)):
                        coll42_ = coll42_.union(_dafny.Set([(d_702_v0_) + ((_dafny.MultiSet([self.f16])).cardinality)]))
                return _dafny.Set(coll42_)
            return iife74_()
            

        init20_ = lambda38_
        nw136_ = _dafny.Array(None, 15)
        for i0_20_ in range(nw136_.length(0)):
            nw136_[i0_20_] = init20_(i0_20_)
        d_700_v1_ = nw136_
        d_703_v2_: _dafny.Set
        d_703_v2_ = _dafny.Set({len(_dafny.Map({p2: self.f16}))})
        index139_ = default__.safeIndex(267, (d_700_v1_).length(0))
        (d_700_v1_)[index139_] = d_703_v2_
        index140_ = default__.safeIndex(267, (d_700_v1_).length(0))
        rhs150_ = d_703_v2_
        rhs151_ = self.f16
        rhs152_ = ((p3) < (p3) if (d_703_v2_).ispropersubset(d_703_v2_) else self.f16)
        lhs130_ = d_700_v1_
        lhs131_ = default__.safeIndex(267, (d_700_v1_).length(0))
        lhs132_ = globalState
        lhs133_ = globalState
        lhs130_[lhs131_] = rhs150_
        lhs132_.f5 = rhs151_
        lhs133_.f2 = rhs152_
        (self).f16 = p1
        d_704_v3_: _dafny.Array
        nw137_ = _dafny.Array(False, 1)
        d_704_v3_ = nw137_
        d_705_v4_: _dafny.Set
        d_705_v4_ = _dafny.Set({d_704_v3_, d_704_v3_})
        d_706_i1_: int
        d_706_i1_ = 0
        with _dafny.label("5"):
            while (d_704_v3_) in ((d_705_v4_) - (d_705_v4_)):
                with _dafny.c_label("5"):
                    if (d_706_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_706_i1_ = (d_706_i1_) + (1)
                    d_707_v5_: D3
                    d_707_v5_ = D3_DC9((self).f17)
                    d_708_v6_: _dafny.Map
                    d_708_v6_ = _dafny.Map({d_704_v3_: p3})
                    d_709_v7_: _dafny.Map
                    d_709_v7_ = _dafny.Map({d_707_v5_: len((d_708_v6_) | (d_708_v6_))})
                    d_709_v7_ = (d_709_v7_).set(d_707_v5_, p3)
                    d_710_v8_: _dafny.Array
                    nw138_ = _dafny.Array(int(0), 12)
                    d_710_v8_ = nw138_
                    d_711_v9_: _dafny.Seq
                    d_711_v9_ = _dafny.SeqWithoutIsStrInference([self.f16, p1])
                    index141_ = default__.safeIndex(940, (d_710_v8_).length(0))
                    (d_710_v8_)[index141_] = len(d_711_v9_)
                    index142_ = default__.safeIndex(940, (d_710_v8_).length(0))
                    (d_710_v8_)[index142_] = p3
                    d_712_v10_: _dafny.Map
                    d_712_v10_ = _dafny.Map({p0: not(False)})
                    d_712_v10_ = (d_712_v10_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_713_i2_ in range(default__.abs(-629))]), self.f16)
                    index143_ = default__.safeIndex(940, (d_710_v8_).length(0))
                    (d_710_v8_)[index143_] = (d_710_v8_)[default__.safeIndex(940, (d_710_v8_).length(0))]
                    pass
            pass
        d_714_v11_: _dafny.Map
        d_714_v11_ = _dafny.Map({p3: (self).f24})
        d_715_v12_: _dafny.Map
        d_715_v12_ = _dafny.Map({self.f16: self.f16})
        d_716_v13_: _dafny.Seq
        d_716_v13_ = _dafny.SeqWithoutIsStrInference([d_715_v12_])
        d_717_v14_: _dafny.Map
        d_717_v14_ = _dafny.Map({d_714_v11_: (d_716_v13_ if self.f16 else _dafny.SeqWithoutIsStrInference([d_715_v12_ for d_718_i3_ in range(default__.abs(607))]))})
        d_717_v14_ = (d_717_v14_).set(d_714_v11_, d_716_v13_)
        d_719_v15_: int
        d_719_v15_ = -688
        d_719_v15_ = (0) - ((d_719_v15_) * (d_719_v15_))
        index144_ = default__.safeIndex(159, (d_704_v3_).length(0))
        (d_704_v3_)[index144_] = self.f16
        index145_ = default__.safeIndex(159, (d_704_v3_).length(0))
        rhs153_ = self.f16
        rhs154_ = not(default__.fm0(p1, self.f16, globalState))
        lhs134_ = globalState
        lhs135_ = d_704_v3_
        lhs136_ = default__.safeIndex(159, (d_704_v3_).length(0))
        lhs134_.f5 = rhs153_
        lhs135_[lhs136_] = rhs154_
        d_720_v16_: _dafny.Seq
        d_720_v16_ = _dafny.SeqWithoutIsStrInference([p1])
        r0 = (d_720_v16_) + (d_720_v16_)
        return r0

    def m8(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        d_721_v0_: int
        d_721_v0_ = 69
        d_722_v1_: _dafny.Seq
        d_722_v1_ = _dafny.SeqWithoutIsStrInference([self.f16, not(self.f16)])
        d_723_v2_: _dafny.Map
        d_723_v2_ = _dafny.Map({-279: default__.safeDivisionInt(d_721_v0_, len(d_722_v1_))})
        d_724_v3_: _dafny.Map
        d_724_v3_ = _dafny.Map({False: d_721_v0_})
        d_723_v2_ = (d_723_v2_).set(((self).f17)[default__.safeIndex(d_721_v0_, len((self).f17))], (d_721_v0_) * (len(d_724_v3_)))
        d_725_i0_: int
        d_725_i0_ = 0
        with _dafny.label("6"):
            while default__.fm0(self.f16, self.f16, globalState):
                with _dafny.c_label("6"):
                    if (d_725_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_725_i0_ = (d_725_i0_) + (1)
                    if (default__.safeDivisionInt(len((self).f24), (0) - (d_721_v0_))) == (878):
                        (globalState).f2 = self.f16
                        d_726_v4_: _dafny.Array
                        nw139_ = _dafny.Array(False, 18)
                        d_726_v4_ = nw139_
                        index146_ = default__.safeIndex(547, (d_726_v4_).length(0))
                        (d_726_v4_)[index146_] = not(self.f16)
                        index147_ = default__.safeIndex(547, (d_726_v4_).length(0))
                        rhs155_ = (d_723_v2_ if not(self.f16) else d_723_v2_)
                        rhs156_ = default__.safeDivisionInt((0) - (len((self).f24)), d_721_v0_)
                        rhs157_ = self.f16
                        lhs137_ = d_726_v4_
                        lhs138_ = default__.safeIndex(547, (d_726_v4_).length(0))
                        d_723_v2_ = rhs155_
                        d_721_v0_ = rhs156_
                        lhs137_[lhs138_] = rhs157_
                        d_727_v5_: _dafny.Array
                        nw140_ = _dafny.Array(_dafny.Map({}), 12)
                        d_727_v5_ = nw140_
                        d_728_v6_: str
                        d_728_v6_ = _dafny.CodePoint('w')
                        index148_ = default__.safeIndex(830, (d_727_v5_).length(0))
                        (d_727_v5_)[index148_] = _dafny.Map({d_721_v0_: d_728_v6_})
                        d_729_v7_: D3
                        d_729_v7_ = D3_DC10((d_726_v4_)[default__.safeIndex(547, (d_726_v4_).length(0))])
                        d_730_v8_: _dafny.Set
                        d_730_v8_ = _dafny.Set({self.f16, self.f16})
                        d_731_v9_: _dafny.Map
                        d_731_v9_ = _dafny.Map({d_721_v0_: d_728_v6_})
                        index149_ = default__.safeIndex(830, (d_727_v5_).length(0))
                        rhs158_ = (((d_726_v4_)[default__.safeIndex(547, (d_726_v4_).length(0))] if (d_726_v4_)[default__.safeIndex(547, (d_726_v4_).length(0))] else not((d_726_v4_)[default__.safeIndex(547, (d_726_v4_).length(0))]))) == (((d_726_v4_)[default__.safeIndex(547, (d_726_v4_).length(0))]) in (d_730_v8_))
                        rhs159_ = (d_731_v9_ if default__.fm0(not(self.f16), self.f16, globalState) else (d_731_v9_) | (d_731_v9_))
                        rhs160_ = d_729_v7_
                        lhs139_ = globalState
                        lhs140_ = d_727_v5_
                        lhs141_ = default__.safeIndex(830, (d_727_v5_).length(0))
                        lhs139_.f5 = rhs158_
                        lhs140_[lhs141_] = rhs159_
                        d_729_v7_ = rhs160_
                        r1 = (0) - ((0) - ((d_721_v0_) * (d_721_v0_)))
                        d_732_v10_: _dafny.Map
                        d_732_v10_ = _dafny.Map({True: (d_726_v4_)[default__.safeIndex(547, (d_726_v4_).length(0))]})
                        d_733_v11_: D1
                        d_733_v11_ = D1_DC4(False, (0) - (d_721_v0_), d_732_v10_)
                        d_734_v12_: _dafny.MultiSet
                        d_734_v12_ = _dafny.MultiSet([_dafny.Map({not(False): 832}), (d_724_v3_) | (d_724_v3_), d_724_v3_])
                        d_735_v13_: _dafny.Seq
                        d_735_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])
                        d_736_v14_: D4
                        d_736_v14_ = D4_DC11(d_734_v12_)
                        rhs161_ = d_733_v11_
                        rhs162_ = (_dafny.MultiSet([d_724_v3_, d_724_v3_])).intersection((d_736_v14_).cf20)
                        rhs163_ = (d_721_v0_) - (d_721_v0_)
                        rhs164_ = d_721_v0_
                        rhs165_ = (_dafny.SeqWithoutIsStrInference([((self).f24)[default__.safeIndex(d_721_v0_, len((self).f24))], _dafny.CodePoint('a')])).set(default__.safeIndex((d_721_v0_) * (d_721_v0_), len(_dafny.SeqWithoutIsStrInference([((self).f24)[default__.safeIndex(d_721_v0_, len((self).f24))], _dafny.CodePoint('a')]))), (d_728_v6_ if self.f16 else ((d_731_v9_)[d_721_v0_] if (d_721_v0_) in (d_731_v9_) else d_728_v6_)))
                        d_733_v11_ = rhs161_
                        d_734_v12_ = rhs162_
                        d_721_v0_ = rhs163_
                        d_721_v0_ = rhs164_
                        d_735_v13_ = rhs165_
                    elif True:
                        rhs166_ = d_721_v0_
                        rhs167_ = -554
                        d_721_v0_ = rhs166_
                        r1 = rhs167_
                        (globalState).f2 = self.f16
                        d_737_v15_: _dafny.Map
                        d_737_v15_ = _dafny.Map({self.f16: self.f16})
                        d_722_v1_ = (d_722_v1_ if ((d_737_v15_)[self.f16] if (self.f16) in (d_737_v15_) else self.f16) else d_722_v1_)
                        d_738_v16_: _dafny.Set
                        d_738_v16_ = _dafny.Set({(0) - (((d_724_v3_)[self.f16] if (self.f16) in (d_724_v3_) else d_721_v0_)), d_721_v0_})
                        d_739_v17_: _dafny.MultiSet
                        d_739_v17_ = _dafny.MultiSet([d_738_v16_])
                        d_739_v17_ = d_739_v17_
                        d_740_v18_: _dafny.Array
                        nw141_ = _dafny.Array(False, 28)
                        d_740_v18_ = nw141_
                        d_741_v19_: C0
                        nw142_ = C0()
                        nw142_.ctor__(d_721_v0_, d_740_v18_)
                        d_741_v19_ = nw142_
                    d_742_v20_: D4
                    d_742_v20_ = D4_DC12(self.f16, default__.fm0(self.f16, self.f16, globalState), d_721_v0_)
                    pat_let_tv19_ = d_721_v0_
                    pat_let_tv20_ = d_721_v0_
                    pat_let_tv21_ = d_742_v20_
                    pat_let_tv22_ = globalState
                    d_743_v21_: _dafny.Array
                    nw143_ = _dafny.Array(None, 17)
                    nw143_[int(0)] = d_742_v20_
                    nw143_[int(1)] = D4_DC12(self.f16, self.f16, default__.fm3(globalState))
                    nw143_[int(2)] = default__.fm21(d_721_v0_, d_721_v0_, globalState)
                    nw143_[int(3)] = d_742_v20_
                    nw143_[int(4)] = d_742_v20_
                    nw143_[int(5)] = d_742_v20_
                    def iife75_(_pat_let16_0):
                        def iife76_(d_744_dt__update__tmp_h0_):
                            def iife77_(_pat_let17_0):
                                def iife78_(d_745_dt__update_hcf23_h0_):
                                    return D4_DC12((d_744_dt__update__tmp_h0_).cf21, (d_744_dt__update__tmp_h0_).cf22, d_745_dt__update_hcf23_h0_)
                                return iife78_(_pat_let17_0)
                            return iife77_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv19_])))
                        return iife76_(_pat_let16_0)
                    nw143_[int(6)] = iife75_(d_742_v20_)
                    nw143_[int(7)] = D4_DC12(self.f16, self.f16, d_721_v0_)
                    nw143_[int(8)] = d_742_v20_
                    def iife79_(_pat_let18_0):
                        def iife80_(d_746_dt__update__tmp_h1_):
                            def iife81_(_pat_let19_0):
                                def iife82_(d_747_dt__update_hcf23_h1_):
                                    return D4_DC12((d_746_dt__update__tmp_h1_).cf21, (d_746_dt__update__tmp_h1_).cf22, d_747_dt__update_hcf23_h1_)
                                return iife82_(_pat_let19_0)
                            return iife81_(pat_let_tv20_)
                        return iife80_(_pat_let18_0)
                    nw143_[int(9)] = iife79_(default__.fm21(d_721_v0_, d_721_v0_, globalState))
                    nw143_[int(10)] = D4_DC12(self.f16, self.f16, (default__.fm21(d_721_v0_, d_721_v0_, globalState)).cf23)
                    def iife83_(_pat_let20_0):
                        def iife84_(d_748_dt__update__tmp_h2_):
                            def iife85_(_pat_let21_0):
                                def iife86_(d_749_dt__update_hcf22_h0_):
                                    def iife88_(_pat_let23_0):
                                        def iife89_(d_750_dt__update__tmp_h3_):
                                            def iife90_(_pat_let24_0):
                                                def iife91_(d_751_dt__update_hcf21_h1_):
                                                    return D4_DC12(d_751_dt__update_hcf21_h1_, (d_750_dt__update__tmp_h3_).cf22, (d_750_dt__update__tmp_h3_).cf23)
                                                return iife91_(_pat_let24_0)
                                            return iife90_(not(self.f16))
                                        return iife89_(_pat_let23_0)
                                    def iife87_(_pat_let22_0):
                                        def iife92_(d_752_dt__update_hcf21_h0_):
                                            return D4_DC12(d_752_dt__update_hcf21_h0_, d_749_dt__update_hcf22_h0_, (d_748_dt__update__tmp_h2_).cf23)
                                        return iife92_(_pat_let22_0)
                                    return iife87_(default__.fm0(self.f16, (iife88_(pat_let_tv21_)).cf21, pat_let_tv22_))
                                return iife86_(_pat_let21_0)
                            return iife85_(not(self.f16))
                        return iife84_(_pat_let20_0)
                    nw143_[int(11)] = iife83_(d_742_v20_)
                    nw143_[int(12)] = d_742_v20_
                    nw143_[int(13)] = d_742_v20_
                    nw143_[int(14)] = d_742_v20_
                    nw143_[int(15)] = d_742_v20_
                    nw143_[int(16)] = d_742_v20_
                    d_743_v21_ = nw143_
                    index150_ = default__.safeIndex(671, (d_743_v21_).length(0))
                    (d_743_v21_)[index150_] = d_742_v20_
                    index151_ = default__.safeIndex(671, (d_743_v21_).length(0))
                    (d_743_v21_)[index151_] = D4_DC12(False, default__.fm0(self.f16, self.f16, globalState), default__.safeModuloInt(d_721_v0_, d_721_v0_))
                    d_753_v22_: _dafny.Map
                    d_753_v22_ = _dafny.Map({default__.safeModuloInt((0) - (d_721_v0_), d_721_v0_): not(self.f16)})
                    d_753_v22_ = (d_753_v22_).set(d_721_v0_, self.f16)
                    d_754_v23_: _dafny.Map
                    d_754_v23_ = _dafny.Map({(self.f16 if self.f16 else self.f16): (d_721_v0_) <= (((d_724_v3_)[True] if (True) in (d_724_v3_) else d_721_v0_))})
                    d_755_v24_: _dafny.Seq
                    d_755_v24_ = _dafny.SeqWithoutIsStrInference([(d_754_v23_).set(self.f16, self.f16)])
                    d_756_v25_: D4
                    d_756_v25_ = D4_DC13(len(d_723_v2_), d_721_v0_, d_721_v0_)
                    d_754_v23_ = (d_754_v23_) | ((d_755_v24_)[default__.safeIndex((d_756_v25_).cf24, len(d_755_v24_))])
                    pass
            pass
        (self).f16 = self.f16
        d_757_v26_: _dafny.Array
        nw144_ = _dafny.Array(False, 20)
        d_757_v26_ = nw144_
        index152_ = default__.safeIndex(870, (d_757_v26_).length(0))
        (d_757_v26_)[index152_] = self.f16
        d_758_v27_: D0
        d_758_v27_ = D0_DC2(d_721_v0_, d_721_v0_, d_721_v0_, d_757_v26_)
        d_759_v28_: _dafny.MultiSet
        d_759_v28_ = _dafny.MultiSet([D0_DC2(d_721_v0_, d_721_v0_, d_721_v0_, d_757_v26_)])
        index153_ = default__.safeIndex(870, (d_757_v26_).length(0))
        (d_757_v26_)[index153_] = (d_758_v27_) not in (d_759_v28_)
        (globalState).f5 = not((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))])
        if not ((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]) or (self.f16):
            d_721_v0_ = d_721_v0_
            d_760_v29_: _dafny.MultiSet
            d_760_v29_ = _dafny.MultiSet([(d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]])
            if ((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]) not in (d_760_v29_):
                d_761_v30_: str
                d_761_v30_ = _dafny.CodePoint('w')
                (globalState).f15 = (_dafny.CodePoint('m') if self.f16 else d_761_v30_)
                d_762_v31_: _dafny.Seq
                d_762_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dita"))
                d_762_v31_ = ((self).f24).set(default__.safeIndex((0) - (d_721_v0_), len((self).f24)), d_761_v30_)
                d_763_v32_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.Seq({}), 16)
                d_763_v32_ = nw145_
                d_764_v33_: D2
                d_764_v33_ = D2_DC8(d_724_v3_, d_763_v32_, False, (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))], d_721_v0_)
                d_765_v34_: D2
                d_765_v34_ = D2_DC8(d_724_v3_, (d_764_v33_).cf14, self.f16, (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))], len(d_724_v3_))
                d_766_v35_: D6
                d_766_v35_ = D6_DC17(d_765_v34_)
                d_767_v36_: _dafny.MultiSet
                d_767_v36_ = _dafny.MultiSet([d_766_v35_])
                d_768_v37_: _dafny.Seq
                d_768_v37_ = _dafny.SeqWithoutIsStrInference([d_767_v36_])
                d_769_v38_: _dafny.Seq
                d_769_v38_ = d_768_v37_
                rhs168_ = default__.fm3(globalState)
                rhs169_ = (d_768_v37_)
                r0 = rhs168_
                d_768_v37_ = rhs169_
                (globalState).f2 = self.f16
                d_770_v39_: _dafny.Array
                nw146_ = _dafny.Array(None, 24)
                nw146_[int(0)] = d_761_v30_
                nw146_[int(1)] = d_761_v30_
                nw146_[int(2)] = d_761_v30_
                nw146_[int(3)] = d_761_v30_
                nw146_[int(4)] = d_761_v30_
                nw146_[int(5)] = d_761_v30_
                nw146_[int(6)] = d_761_v30_
                nw146_[int(7)] = ((d_762_v31_)[default__.safeIndex(d_721_v0_, len(d_762_v31_))] if (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))] else d_761_v30_)
                nw146_[int(8)] = d_761_v30_
                nw146_[int(9)] = d_761_v30_
                nw146_[int(10)] = d_761_v30_
                nw146_[int(11)] = d_761_v30_
                nw146_[int(12)] = d_761_v30_
                nw146_[int(13)] = _dafny.CodePoint('g')
                nw146_[int(14)] = d_761_v30_
                nw146_[int(15)] = d_761_v30_
                nw146_[int(16)] = d_761_v30_
                nw146_[int(17)] = d_761_v30_
                nw146_[int(18)] = default__.fm2(globalState)
                nw146_[int(19)] = d_761_v30_
                nw146_[int(20)] = d_761_v30_
                nw146_[int(21)] = d_761_v30_
                nw146_[int(22)] = d_761_v30_
                nw146_[int(23)] = d_761_v30_
                d_770_v39_ = nw146_
                d_770_v39_ = d_770_v39_
            elif True:
                d_771_v40_: _dafny.Array
                nw147_ = _dafny.Array(D6.default()(), 20)
                d_771_v40_ = nw147_
                d_772_v41_: D6
                d_772_v41_ = D6_DC16(True, self.f16)
                index154_ = default__.safeIndex(393, (d_771_v40_).length(0))
                (d_771_v40_)[index154_] = d_772_v41_
                index155_ = default__.safeIndex(393, (d_771_v40_).length(0))
                (d_771_v40_)[index155_] = D6_DC16((d_721_v0_) < (len((self).f24)), self.f16)
                d_773_v42_: _dafny.Array
                def lambda39_(d_774_i1_):
                    return ((self).f17) + ((self).f17)

                init21_ = lambda39_
                nw148_ = _dafny.Array(None, 25)
                for i0_21_ in range(nw148_.length(0)):
                    nw148_[i0_21_] = init21_(i0_21_)
                d_773_v42_ = nw148_
                d_775_v43_: _dafny.Map
                d_775_v43_ = _dafny.Map({d_722_v1_: d_721_v0_})
                index156_ = default__.safeIndex(544, (d_773_v42_).length(0))
                (d_773_v42_)[index156_] = _dafny.SeqWithoutIsStrInference([d_721_v0_, len(d_775_v43_), d_721_v0_])
                index157_ = default__.safeIndex(544, (d_773_v42_).length(0))
                (d_773_v42_)[index157_] = default__.fm17(globalState)
                d_776_v44_: _dafny.Set
                d_776_v44_ = _dafny.Set({(0) - (d_721_v0_), d_721_v0_, d_721_v0_, (0) - (d_721_v0_), d_721_v0_})
                (self).f16 = not(not((d_776_v44_).ispropersubset(d_776_v44_)))
                d_777_v45_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_777_v45_ = nw149_
                d_778_v46_: _dafny.Set
                d_778_v46_ = _dafny.Set({d_722_v1_, d_722_v1_})
                d_779_v47_: _dafny.Map
                d_779_v47_ = _dafny.Map({d_777_v45_: d_778_v46_})
                d_779_v47_ = (d_779_v47_).set(d_777_v45_, _dafny.Set({d_722_v1_, (default__.fm22(self.f16, _dafny.MultiSet((d_773_v42_)[default__.safeIndex(544, (d_773_v42_).length(0))]), (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))], globalState)).cf34}))
                d_780_v48_: _dafny.Array
                nw150_ = _dafny.Array(int(0), 15)
                d_780_v48_ = nw150_
                index158_ = default__.safeIndex(903, (d_780_v48_).length(0))
                (d_780_v48_)[index158_] = d_721_v0_
                index159_ = default__.safeIndex(903, (d_780_v48_).length(0))
                rhs170_ = d_721_v0_
                rhs171_ = d_760_v29_
                rhs172_ = -58
                rhs173_ = d_758_v27_
                lhs142_ = d_780_v48_
                lhs143_ = default__.safeIndex(903, (d_780_v48_).length(0))
                r1 = rhs170_
                d_760_v29_ = rhs171_
                lhs142_[lhs143_] = rhs172_
                d_758_v27_ = rhs173_
            if not((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]):
                d_781_v49_: _dafny.Array
                nw151_ = _dafny.Array(int(0), 8)
                d_781_v49_ = nw151_
                rhs174_ = default__.fm3(globalState)
                rhs175_ = d_781_v49_
                rhs176_ = d_721_v0_
                rhs177_ = (((0) - (d_721_v0_) if self.f16 else d_721_v0_)) + (d_721_v0_)
                d_721_v0_ = rhs174_
                d_781_v49_ = rhs175_
                r0 = rhs176_
                d_721_v0_ = rhs177_
                r1 = len(((((self).f24) + ((self).f24)).set(default__.safeIndex(d_721_v0_, len(((self).f24) + ((self).f24))), _dafny.CodePoint('x'))) + ((self).f24))
                r0 = 960
                d_782_v50_: _dafny.Array
                nw152_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_782_v50_ = nw152_
                index160_ = default__.safeIndex(834, (d_782_v50_).length(0))
                (d_782_v50_)[index160_] = d_781_v49_
                index161_ = default__.safeIndex(834, (d_782_v50_).length(0))
                (d_782_v50_)[index161_] = d_781_v49_
                d_783_v51_: _dafny.Map
                d_783_v51_ = _dafny.Map({(d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]: (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]})
                d_784_v52_: _dafny.Map
                d_784_v52_ = _dafny.Map({d_783_v51_: (d_781_v49_ if self.f16 else (d_782_v50_)[default__.safeIndex(834, (d_782_v50_).length(0))])})
                d_785_v53_: _dafny.Array
                def lambda40_(d_786_v51_):
                    def lambda41_(d_787_i2_):
                        return (d_786_v51_).set(self.f16, self.f16)

                    return lambda41_

                init22_ = lambda40_(d_783_v51_)
                nw153_ = _dafny.Array(None, 24)
                for i0_22_ in range(nw153_.length(0)):
                    nw153_[i0_22_] = init22_(i0_22_)
                d_785_v53_ = nw153_
                index162_ = default__.safeIndex(631, (d_785_v53_).length(0))
                (d_785_v53_)[index162_] = d_783_v51_
                index163_ = default__.safeIndex(631, (d_785_v53_).length(0))
                rhs178_ = d_784_v52_
                rhs179_ = ((d_783_v51_) | (d_783_v51_)) | ((default__.fm23(_dafny.MultiSet([(0) - (d_721_v0_), d_721_v0_, d_721_v0_]), not(not((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))])), globalState)).set(self.f16, (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]))
                lhs144_ = d_785_v53_
                lhs145_ = default__.safeIndex(631, (d_785_v53_).length(0))
                d_784_v52_ = rhs178_
                lhs144_[lhs145_] = rhs179_
            elif True:
                def iife93_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(147, 327):
                        d_788_v54_: int = compr_43_
                        if ((147) <= (d_788_v54_)) and ((d_788_v54_) < (327)):
                            coll43_[default__.safeModuloInt(d_788_v54_, (d_760_v29_).cardinality)] = len((self).f24)
                    return _dafny.Map(coll43_)
                rhs180_ = iife93_()

                rhs181_ = d_721_v0_
                d_723_v2_ = rhs180_
                r1 = rhs181_
                d_789_v55_: _dafny.Seq
                d_789_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "he"))
                d_789_v55_ = d_789_v55_
                d_790_v56_: _dafny.Array
                def lambda42_(d_791_v0_):
                    def lambda43_(d_792_i3_):
                        return (d_792_i3_) + (d_791_v0_)

                    return lambda43_

                init23_ = lambda42_(d_721_v0_)
                nw154_ = _dafny.Array(None, 22)
                for i0_23_ in range(nw154_.length(0)):
                    nw154_[i0_23_] = init23_(i0_23_)
                d_790_v56_ = nw154_
                def lambda44_(d_793_v0_):
                    def lambda45_(d_794_i4_):
                        return default__.safeDivisionInt(d_794_i4_, d_793_v0_)

                    return lambda45_

                init24_ = lambda44_(d_721_v0_)
                nw155_ = _dafny.Array(None, 25)
                for i0_24_ in range(nw155_.length(0)):
                    nw155_[i0_24_] = init24_(i0_24_)
                d_790_v56_ = nw155_
                r1 = d_721_v0_
                d_724_v3_ = (d_724_v3_) | (_dafny.Map({not((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]): d_721_v0_}))
            d_795_v57_: D8
            d_795_v57_ = D8_DC23()
            d_796_v58_: _dafny.Set
            d_796_v58_ = _dafny.Set({d_795_v57_, d_795_v57_, d_795_v57_})
            d_797_v59_: str
            d_797_v59_ = _dafny.CodePoint('s')
            d_798_v60_: _dafny.Map
            d_798_v60_ = _dafny.Map({d_797_v59_: len((self).f24)})
            d_799_v61_: _dafny.MultiSet
            d_799_v61_ = _dafny.MultiSet([(d_796_v58_) | (d_796_v58_), _dafny.Set({d_795_v57_, default__.fm24(((self).f17)[default__.safeIndex(len(d_798_v60_), len((self).f17))], d_797_v59_, globalState)}), d_796_v58_, d_796_v58_, _dafny.Set({d_795_v57_})])
            d_800_v62_: _dafny.Seq
            d_800_v62_ = _dafny.SeqWithoutIsStrInference([(len((self).f17)) - (d_721_v0_)])
            index164_ = default__.safeIndex(870, (d_757_v26_).length(0))
            rhs182_ = default__.fm25(True, d_721_v0_, globalState)
            rhs183_ = 675
            rhs184_ = (not(self.f16) if (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))] else self.f16)
            rhs185_ = d_800_v62_
            rhs186_ = ((_dafny.SeqWithoutIsStrInference([self.f16, self.f16])) + (d_722_v1_)) + (d_722_v1_)
            lhs146_ = d_757_v26_
            lhs147_ = default__.safeIndex(870, (d_757_v26_).length(0))
            d_799_v61_ = rhs182_
            d_721_v0_ = rhs183_
            lhs146_[lhs147_] = rhs184_
            d_800_v62_ = rhs185_
            d_722_v1_ = rhs186_
            d_801_v63_: _dafny.Array
            def lambda46_(d_802_v0_):
                def lambda47_(d_803_i5_):
                    return (d_803_i5_) * (d_802_v0_)

                return lambda47_

            init25_ = lambda46_(d_721_v0_)
            nw156_ = _dafny.Array(None, 9)
            for i0_25_ in range(nw156_.length(0)):
                nw156_[i0_25_] = init25_(i0_25_)
            d_801_v63_ = nw156_
            index165_ = default__.safeIndex(829, (d_801_v63_).length(0))
            (d_801_v63_)[index165_] = d_721_v0_
            index166_ = default__.safeIndex(829, (d_801_v63_).length(0))
            (d_801_v63_)[index166_] = (0) - (d_721_v0_)
        elif True:
            d_804_v64_: _dafny.Set
            d_804_v64_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([default__.fm3(globalState)]), _dafny.SeqWithoutIsStrInference([(0) - (d_721_v0_)]), _dafny.SeqWithoutIsStrInference([d_721_v0_])})
            if (d_804_v64_).ispropersubset(d_804_v64_):
                d_805_v65_: _dafny.Array
                nw157_ = _dafny.Array(int(0), 15)
                d_805_v65_ = nw157_
                index167_ = default__.safeIndex(750, (d_805_v65_).length(0))
                (d_805_v65_)[index167_] = (0) - (d_721_v0_)
                index168_ = default__.safeIndex(750, (d_805_v65_).length(0))
                (d_805_v65_)[index168_] = ((d_724_v3_)[(d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]] if ((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]) in (d_724_v3_) else (d_721_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isodoen")))))
                (globalState).f2 = (((d_805_v65_)[default__.safeIndex(750, (d_805_v65_).length(0))]) > (d_721_v0_)) or (False)
                index169_ = default__.safeIndex(870, (d_757_v26_).length(0))
                (d_757_v26_)[index169_] = not((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))])
                d_724_v3_ = (d_724_v3_).set(self.f16, (d_805_v65_)[default__.safeIndex(750, (d_805_v65_).length(0))])
                def iife94_(_pat_let25_0):
                    def iife95_(d_806_dt__update__tmp_h5_):
                        def iife96_(_pat_let26_0):
                            def iife97_(d_807_dt__update_hcf39_h0_):
                                return D8_DC21((d_806_dt__update__tmp_h5_).cf35, (d_806_dt__update__tmp_h5_).cf36, (d_806_dt__update__tmp_h5_).cf37, (d_806_dt__update__tmp_h5_).cf38, d_807_dt__update_hcf39_h0_)
                            return iife97_(_pat_let26_0)
                        return iife96_(856)
                    return iife95_(_pat_let25_0)
                default__.m0((d_805_v65_)[default__.safeIndex(750, (d_805_v65_).length(0))], (self).f24, (iife94_(default__.fm26(globalState))).cf35, d_721_v0_, globalState)
            elif True:
                d_808_v66_: _dafny.Map
                d_808_v66_ = _dafny.Map({self.f16: self.f16})
                d_809_v67_: C1
                nw158_ = C1()
                nw158_.ctor__((d_721_v0_) * (d_721_v0_), d_721_v0_, default__.fm0(((d_808_v66_)[(d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]] if ((d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]) in (d_808_v66_) else self.f16), False, globalState), (self).f17, d_721_v0_)
                d_809_v67_ = nw158_
                d_721_v0_ = 853
                r0 = (d_809_v67_).f30
                d_810_v68_: _dafny.Array
                def lambda48_(d_811_v67_):
                    def lambda49_(d_812_i6_):
                        return default__.safeDivisionInt(d_812_i6_, (d_811_v67_).f30)

                    return lambda49_

                init26_ = lambda48_(d_809_v67_)
                nw159_ = _dafny.Array(None, 7)
                for i0_26_ in range(nw159_.length(0)):
                    nw159_[i0_26_] = init26_(i0_26_)
                d_810_v68_ = nw159_
                d_813_v69_: D6
                d_813_v69_ = D6_DC18(d_810_v68_)
                index170_ = default__.safeIndex(870, (d_757_v26_).length(0))
                index171_ = default__.safeIndex(870, (d_757_v26_).length(0))
                rhs187_ = (d_813_v69_).cf32
                rhs188_ = not(not((len(d_722_v1_)) <= (d_721_v0_)))
                rhs189_ = (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]
                rhs190_ = (d_722_v1_)[default__.safeIndex(d_721_v0_, len(d_722_v1_))]
                rhs191_ = d_810_v68_
                lhs148_ = d_757_v26_
                lhs149_ = default__.safeIndex(870, (d_757_v26_).length(0))
                lhs150_ = globalState
                lhs151_ = d_757_v26_
                lhs152_ = default__.safeIndex(870, (d_757_v26_).length(0))
                d_810_v68_ = rhs187_
                lhs148_[lhs149_] = rhs188_
                lhs150_.f5 = rhs189_
                lhs151_[lhs152_] = rhs190_
                d_810_v68_ = rhs191_
                d_814_v70_: C2
                nw160_ = C2()
                nw160_.ctor__(d_721_v0_)
                d_814_v70_ = nw160_
                d_814_v70_ = (d_814_v70_ if self.f16 else d_814_v70_)
            d_815_v71_: _dafny.Seq
            d_815_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmq"))
            d_816_v72_: str
            d_816_v72_ = _dafny.CodePoint('i')
            d_815_v71_ = (d_815_v71_).set(default__.safeIndex(d_721_v0_, len(d_815_v71_)), d_816_v72_)
            d_721_v0_ = (d_721_v0_) - (d_721_v0_)
            d_817_v73_: _dafny.Map
            d_817_v73_ = _dafny.Map({self.f16: (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]})
            d_818_v74_: _dafny.Map
            d_818_v74_ = _dafny.Map({(d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]: d_757_v26_})
            d_819_v75_: _dafny.MultiSet
            d_819_v75_ = _dafny.MultiSet([len(d_818_v74_), d_721_v0_])
            (globalState).f5 = (_dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference([d_721_v0_, len(d_722_v1_)])).set(default__.safeIndex((0) - (d_721_v0_), len(_dafny.SeqWithoutIsStrInference([d_721_v0_, len(d_722_v1_)]))), len(d_817_v73_))), d_721_v0_, (0) - (d_721_v0_), d_721_v0_])).ispropersubset((d_819_v75_ if (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))] else _dafny.MultiSet([len(d_815_v71_)])))
            (globalState).f2 = (d_757_v26_)[default__.safeIndex(870, (d_757_v26_).length(0))]
        r0 = d_721_v0_
        r1 = d_721_v0_
        d_820_v76_: _dafny.Seq
        d_820_v76_ = _dafny.SeqWithoutIsStrInference([default__.fm17(globalState), _dafny.SeqWithoutIsStrInference([d_721_v0_, d_721_v0_]), (self).f17, (self).f17])
        d_821_v77_: _dafny.Map
        d_821_v77_ = _dafny.Map({len((d_820_v76_)[default__.safeIndex(d_721_v0_, len(d_820_v76_))]): False})
        r2 = d_821_v77_
        return r0, r1, r2

    @property
    def f24(self):
        return self._f24

class C4:
    def  __init__(self):
        self.f23: bool = False
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f22, f23):
        (self)._f22 = f22
        (self).f23 = f23

    def fm9(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([533, 795]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f22, (self).f22])).cardinality])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f22, (self).f22, (self).f22, (self).f22, (self).f22]), _dafny.SeqWithoutIsStrInference([(self).f22 for d_822_i0_ in range(default__.abs(726))])]))) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f22]) for d_823_i1_ in range(default__.abs(-128))]))

    def fm10(self, p0, p1, globalState):
        if (D1_DC5(len(_dafny.SeqWithoutIsStrInference([self.f23, False])), False)).cf11:
            return _dafny.Map({(self).f22: _dafny.MultiSet([(self).f22, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([D2_DC7(_dafny.Set({self.f23}))])): (self).f22}))])})
        elif self.f23:
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, self.f23, self.f23, self.f23, self.f23])): _dafny.MultiSet([len(_dafny.Map({(self).f22: self.f23})), 702, (0) - ((self).f22), (self).f22, 62])})
        elif True:
            return _dafny.Map({508: _dafny.MultiSet([(self).f22])})

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        d_824_v0_: str
        d_824_v0_ = _dafny.CodePoint('j')
        d_825_v1_: _dafny.Seq
        d_825_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvsewr"))
        (globalState).f2 = (d_824_v0_) not in (d_825_v1_)
        if self.f23:
            d_826_v2_: _dafny.MultiSet
            d_826_v2_ = _dafny.MultiSet([p1])
            d_827_v3_: D1
            d_827_v3_ = D1_DC3(d_826_v2_)
            source9_ = d_827_v3_
            if source9_.is_DC4:
                d_828___mcc_h0_ = source9_.cf7
                d_829___mcc_h1_ = source9_.cf8
                d_830___mcc_h2_ = source9_.cf9
                d_831_cf9_ = d_830___mcc_h2_
                d_832_cf8_ = d_829___mcc_h1_
                d_833_cf7_ = d_828___mcc_h0_
                d_834_v4_: _dafny.Array
                def lambda50_(d_835_p2_):
                    def lambda51_(d_836_i0_):
                        return (d_836_i0_) + (d_835_p2_)

                    return lambda51_

                init27_ = lambda50_(p2)
                nw161_ = _dafny.Array(None, 7)
                for i0_27_ in range(nw161_.length(0)):
                    nw161_[i0_27_] = init27_(i0_27_)
                d_834_v4_ = nw161_
                index172_ = default__.safeIndex(375, (d_834_v4_).length(0))
                (d_834_v4_)[index172_] = (368 if (self).fm9(p2, globalState) else default__.fm3(globalState))
                index173_ = default__.safeIndex(375, (d_834_v4_).length(0))
                rhs192_ = d_832_cf8_
                rhs193_ = p2
                rhs194_ = p2
                rhs195_ = d_832_cf8_
                rhs196_ = (-974) * ((self).f22)
                lhs153_ = d_834_v4_
                lhs154_ = default__.safeIndex(375, (d_834_v4_).length(0))
                lhs153_[lhs154_] = rhs192_
                r2 = rhs193_
                d_832_cf8_ = rhs194_
                r2 = rhs195_
                r2 = rhs196_
                d_831_cf9_ = (d_831_cf9_).set(d_833_cf7_, self.f23)
                d_837_v5_: _dafny.Map
                d_837_v5_ = _dafny.Map({772: (self).f22})
                d_838_v6_: _dafny.Seq
                d_838_v6_ = _dafny.SeqWithoutIsStrInference([d_833_cf7_])
                d_837_v5_ = (d_837_v5_).set((0) - ((0) - (len(_dafny.Map({len(d_838_v6_): p0})))), d_832_cf8_)
                d_839_v7_: _dafny.Map
                d_839_v7_ = _dafny.Map({d_824_v0_: self.f23})
                d_840_v8_: D9
                d_840_v8_ = D9_DC25(d_839_v7_)
                d_841_v9_: _dafny.Map
                d_841_v9_ = _dafny.Map({len((d_840_v8_).cf44): False})
                d_842_v10_: _dafny.Array
                nw162_ = _dafny.Array(None, 16)
                nw162_[int(0)] = ((d_841_v9_)[(d_834_v4_)[default__.safeIndex(375, (d_834_v4_).length(0))]] if ((d_834_v4_)[default__.safeIndex(375, (d_834_v4_).length(0))]) in (d_841_v9_) else p1)
                nw162_[int(1)] = not(self.f23)
                nw162_[int(2)] = d_833_cf7_
                nw162_[int(3)] = p0
                nw162_[int(4)] = p0
                nw162_[int(5)] = self.f23
                nw162_[int(6)] = d_833_cf7_
                nw162_[int(7)] = self.f23
                nw162_[int(8)] = p0
                nw162_[int(9)] = p1
                nw162_[int(10)] = p1
                nw162_[int(11)] = False
                nw162_[int(12)] = False
                nw162_[int(13)] = p1
                nw162_[int(14)] = self.f23
                nw162_[int(15)] = p1
                d_842_v10_ = nw162_
                d_843_v11_: C0
                nw163_ = C0()
                nw163_.ctor__((d_834_v4_)[default__.safeIndex(375, (d_834_v4_).length(0))], d_842_v10_)
                d_843_v11_ = nw163_
            elif source9_.is_DC5:
                d_844___mcc_h3_ = source9_.cf10
                d_845___mcc_h4_ = source9_.cf11
                d_846_cf11_ = d_845___mcc_h4_
                d_847_cf10_ = d_844___mcc_h3_
                d_848_v12_: _dafny.Seq
                d_848_v12_ = _dafny.SeqWithoutIsStrInference([False, self.f23])
                d_849_v13_: _dafny.Map
                d_849_v13_ = _dafny.Map({False: p1})
                d_850_v14_: _dafny.Set
                d_850_v14_ = _dafny.Set({(self).f22})
                d_851_v15_: _dafny.Map
                d_851_v15_ = _dafny.Map({d_850_v14_: ((d_849_v13_)[p1] if (p1) in (d_849_v13_) else not(False))})
                d_852_v16_: _dafny.Seq
                d_852_v16_ = _dafny.SeqWithoutIsStrInference([p2, len(d_851_v15_)])
                d_853_v17_: C1
                nw164_ = C1()
                nw164_.ctor__(len(d_848_v12_), len(d_849_v13_), not(p1), (_dafny.SeqWithoutIsStrInference([default__.fm3(globalState)])) + (d_852_v16_), (self).f22)
                d_853_v17_ = nw164_
                (globalState).f2 = not(d_846_cf11_)
                d_854_v18_: _dafny.Map
                d_854_v18_ = _dafny.Map({(d_853_v17_).f30: p0})
                d_855_v19_: _dafny.Array
                nw165_ = _dafny.Array(None, 8)
                nw165_[int(0)] = p1
                nw165_[int(1)] = d_846_cf11_
                nw165_[int(2)] = d_846_cf11_
                nw165_[int(3)] = p0
                nw165_[int(4)] = False
                nw165_[int(5)] = d_846_cf11_
                nw165_[int(6)] = (d_848_v12_)[default__.safeIndex((d_853_v17_).f30, len(d_848_v12_))]
                nw165_[int(7)] = ((d_854_v18_)[len(d_848_v12_)] if (len(d_848_v12_)) in (d_854_v18_) else default__.fm0(p0, p1, globalState))
                d_855_v19_ = nw165_
                d_855_v19_ = d_855_v19_
                d_856_v20_: C0
                nw166_ = C0()
                nw166_.ctor__((d_853_v17_).fm13(self.f23, globalState), d_855_v19_)
                d_856_v20_ = nw166_
            elif source9_.is_DC6:
                r2 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_857_i1_ in range(default__.abs(447))]))
                d_858_v21_: _dafny.Array
                nw167_ = _dafny.Array(False, 21)
                d_858_v21_ = nw167_
                index174_ = default__.safeIndex(330, (d_858_v21_).length(0))
                (d_858_v21_)[index174_] = (p0) == (p0)
                index175_ = default__.safeIndex(330, (d_858_v21_).length(0))
                (d_858_v21_)[index175_] = not(True)
                (globalState).f2 = p0
                d_859_v22_: _dafny.Array
                def lambda52_(d_860_p2_):
                    def lambda53_(d_861_i2_):
                        return _dafny.Map({False: d_860_p2_})

                    return lambda53_

                init28_ = lambda52_(p2)
                nw168_ = _dafny.Array(None, 6)
                for i0_28_ in range(nw168_.length(0)):
                    nw168_[i0_28_] = init28_(i0_28_)
                d_859_v22_ = nw168_
                d_862_v23_: _dafny.Map
                d_862_v23_ = _dafny.Map({d_824_v0_: d_859_v22_})
                d_862_v23_ = (d_862_v23_).set(default__.fm2(globalState), d_859_v22_)
            elif True:
                d_863___mcc_h5_ = source9_.cf6
                d_864_cf6_ = d_863___mcc_h5_
                d_865_v24_: _dafny.MultiSet
                d_865_v24_ = _dafny.MultiSet([d_825_v1_, d_825_v1_])
                d_866_v25_: _dafny.Seq
                d_866_v25_ = _dafny.SeqWithoutIsStrInference([(d_865_v24_).cardinality, len(_dafny.SeqWithoutIsStrInference([d_824_v0_ for d_867_i3_ in range(default__.abs(234))]))])
                d_868_v26_: D3
                d_868_v26_ = D3_DC10(p1)
                d_869_v27_: _dafny.Set
                d_869_v27_ = _dafny.Set({False, not((d_868_v26_).cf19)})
                d_870_v28_: C1
                nw169_ = C1()
                nw169_.ctor__((0) - ((self).f22), p2, p0, d_866_v25_, len(d_869_v27_))
                d_870_v28_ = nw169_
                d_871_v29_: _dafny.MultiSet
                d_871_v29_ = _dafny.MultiSet([d_870_v28_])
                r2 = default__.safeDivisionInt(((d_871_v29_)[d_870_v28_] if (d_870_v28_) in (d_871_v29_) else (d_870_v28_).fm14(p2, (d_870_v28_).f30, d_824_v0_, globalState)), default__.safeModuloInt((d_870_v28_).f29, len((_dafny.Map({(d_870_v28_).f29: p2})).set(p2, (d_870_v28_).f29))))
                d_872_v30_: D8
                d_872_v30_ = D8_DC23()
                d_873_v31_: D8
                d_873_v31_ = D8_DC24((d_872_v30_ if p0 else d_872_v30_))
                d_873_v31_ = d_873_v31_
                d_874_v32_: _dafny.Map
                d_874_v32_ = _dafny.Map({d_825_v1_: (self).fm9(p2, globalState)})
                d_874_v32_ = (d_874_v32_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weptditry")), self.f23)
                d_875_v33_: _dafny.Map
                d_875_v33_ = _dafny.Map({not(False): p0})
                d_876_v34_: _dafny.Map
                d_876_v34_ = _dafny.Map({len((d_875_v33_).set(p1, p1)): p2})
                (globalState).f5 = ((d_870_v28_).fm15((d_866_v25_).set(default__.safeIndex((self).f22, len(d_866_v25_)), (d_870_v28_).f30), len(d_869_v27_), (d_870_v28_).f30, d_824_v0_, globalState)) > (((d_876_v34_)[p2] if (p2) in (d_876_v34_) else p2))
            d_877_v35_: _dafny.Array
            nw170_ = _dafny.Array(False, 22)
            d_877_v35_ = nw170_
            index176_ = default__.safeIndex(738, (d_877_v35_).length(0))
            (d_877_v35_)[index176_] = p0
            d_878_v36_: _dafny.Array
            nw171_ = _dafny.Array(None, 24)
            nw171_[int(0)] = d_877_v35_
            nw171_[int(1)] = d_877_v35_
            nw171_[int(2)] = d_877_v35_
            nw171_[int(3)] = d_877_v35_
            nw171_[int(4)] = d_877_v35_
            nw171_[int(5)] = d_877_v35_
            nw171_[int(6)] = d_877_v35_
            nw171_[int(7)] = d_877_v35_
            nw171_[int(8)] = d_877_v35_
            nw171_[int(9)] = d_877_v35_
            nw171_[int(10)] = d_877_v35_
            nw171_[int(11)] = d_877_v35_
            nw171_[int(12)] = d_877_v35_
            nw171_[int(13)] = d_877_v35_
            nw171_[int(14)] = d_877_v35_
            nw171_[int(15)] = (d_877_v35_ if p0 else d_877_v35_)
            nw171_[int(16)] = d_877_v35_
            nw171_[int(17)] = d_877_v35_
            nw171_[int(18)] = d_877_v35_
            nw171_[int(19)] = d_877_v35_
            nw171_[int(20)] = d_877_v35_
            nw171_[int(21)] = d_877_v35_
            nw171_[int(22)] = (d_877_v35_ if False else d_877_v35_)
            nw171_[int(23)] = d_877_v35_
            d_878_v36_ = nw171_
            d_879_v37_: _dafny.Seq
            d_879_v37_ = _dafny.SeqWithoutIsStrInference([p2])
            d_880_v38_: _dafny.Map
            d_880_v38_ = _dafny.Map({p2: 569})
            d_881_v39_: _dafny.Map
            d_881_v39_ = _dafny.Map({p1: (d_879_v37_)[default__.safeIndex((self).f22, len(d_879_v37_))]})
            index177_ = default__.safeIndex(738, (d_877_v35_).length(0))
            rhs197_ = (((0) - ((self).f22)) > (p2)) or (p1)
            rhs198_ = p2
            rhs199_ = len((d_879_v37_).set(default__.safeIndex(default__.safeModuloInt((self).f22, len(_dafny.Set({996, len(d_825_v1_)}))), len(d_879_v37_)), ((d_880_v38_)[p2] if (p2) in (d_880_v38_) else default__.fm3(globalState))))
            rhs200_ = (default__.fm3(globalState)) < (len(d_881_v39_))
            rhs201_ = d_878_v36_
            lhs155_ = globalState
            lhs156_ = d_877_v35_
            lhs157_ = default__.safeIndex(738, (d_877_v35_).length(0))
            lhs155_.f2 = rhs197_
            r2 = rhs198_
            r2 = rhs199_
            lhs156_[lhs157_] = rhs200_
            d_878_v36_ = rhs201_
            if ((d_825_v1_) + (default__.fm1(globalState))) == (d_825_v1_):
                d_882_v40_: C3
                nw172_ = C3()
                nw172_.ctor__(d_825_v1_, self.f23, default__.fm17(globalState))
                d_882_v40_ = nw172_
                d_883_v41_: _dafny.Map
                d_883_v41_ = _dafny.Map({d_824_v0_: self.f23})
                d_884_v42_: _dafny.Map
                d_884_v42_ = _dafny.Map({D9_DC25(d_883_v41_): (p2) - ((self).f22)})
                d_885_v44_: _dafny.Map
                d_885_v44_ = _dafny.Map({d_824_v0_: p2})
                d_886_v45_: D9
                def iife98_():
                    coll44_ = _dafny.Map()
                    compr_44_: str
                    for compr_44_ in (d_885_v44_).keys.Elements:
                        d_887_v43_: str = compr_44_
                        if (d_887_v43_) in (d_885_v44_):
                            coll44_[d_887_v43_] = p0
                    return _dafny.Map(coll44_)
                d_886_v45_ = D9_DC25(iife98_()
)
                index178_ = default__.safeIndex(738, (d_877_v35_).length(0))
                index179_ = default__.safeIndex(738, (d_877_v35_).length(0))
                rhs202_ = p0
                rhs203_ = p0
                rhs204_ = (d_884_v42_).set(d_886_v45_, 539)
                rhs205_ = (self).f22
                lhs158_ = d_877_v35_
                lhs159_ = default__.safeIndex(738, (d_877_v35_).length(0))
                lhs160_ = d_877_v35_
                lhs161_ = default__.safeIndex(738, (d_877_v35_).length(0))
                lhs158_[lhs159_] = rhs202_
                lhs160_[lhs161_] = rhs203_
                d_884_v42_ = rhs204_
                r2 = rhs205_
                d_888_v46_: _dafny.Array
                nw173_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_888_v46_ = nw173_
                d_889_v47_: _dafny.Array
                nw174_ = _dafny.Array(int(0), 16)
                d_889_v47_ = nw174_
                index180_ = default__.safeIndex(235, (d_888_v46_).length(0))
                (d_888_v46_)[index180_] = d_889_v47_
                index181_ = default__.safeIndex(235, (d_888_v46_).length(0))
                (d_888_v46_)[index181_] = d_889_v47_
                d_890_v48_: _dafny.Seq
                d_890_v48_ = _dafny.SeqWithoutIsStrInference([(d_877_v35_)[default__.safeIndex(738, (d_877_v35_).length(0))], p0, False, ((self).f22) in (_dafny.Set({(self).f22}))])
                d_891_v51_: _dafny.Array
                def lambda54_(d_892_v37_):
                    def lambda55_(d_893_i4_):
                        def iife99_():
                            coll45_ = _dafny.Set()
                            compr_45_: int
                            for compr_45_ in (d_892_v37_).Elements:
                                d_894_v49_: int = compr_45_
                                if (d_894_v49_) in (d_892_v37_):
                                    coll45_ = coll45_.union(_dafny.Set([(d_894_v49_) + ((0) - (-821))]))
                            return _dafny.Set(coll45_)
                        def iife100_():
                            coll46_ = _dafny.Set()
                            compr_46_: int
                            for compr_46_ in _dafny.IntegerRange(375, 298):
                                d_895_v50_: int = compr_46_
                                if ((375) <= (d_895_v50_)) and ((d_895_v50_) < (298)):
                                    coll46_ = coll46_.union(_dafny.Set([(d_895_v50_) - ((self).f22)]))
                            return _dafny.Set(coll46_)
                        return (_dafny.Set({iife99_()
                        , iife100_()
                        })) - (_dafny.Set({_dafny.Set({(self).f22})}))

                    return lambda55_

                init29_ = lambda54_(d_879_v37_)
                nw175_ = _dafny.Array(None, 10)
                for i0_29_ in range(nw175_.length(0)):
                    nw175_[i0_29_] = init29_(i0_29_)
                d_891_v51_ = nw175_
                d_896_v52_: _dafny.Set
                d_896_v52_ = _dafny.Set({(0) - (p2)})
                d_897_v53_: _dafny.Set
                d_897_v53_ = _dafny.Set({d_896_v52_, d_896_v52_})
                index182_ = default__.safeIndex(405, (d_891_v51_).length(0))
                (d_891_v51_)[index182_] = d_897_v53_
                d_898_v54_: _dafny.Map
                d_898_v54_ = _dafny.Map({d_825_v1_: _dafny.CodePoint('g')})
                d_899_v56_: _dafny.Seq
                d_899_v56_ = _dafny.SeqWithoutIsStrInference([d_825_v1_])
                index183_ = default__.safeIndex(235, (d_888_v46_).length(0))
                index184_ = default__.safeIndex(405, (d_891_v51_).length(0))
                def iife101_():
                    coll47_ = _dafny.Map()
                    compr_47_: _dafny.Seq
                    for compr_47_ in (d_899_v56_).Elements:
                        d_900_v55_: _dafny.Seq = compr_47_
                        if (d_900_v55_) in (d_899_v56_):
                            coll47_[d_900_v55_] = _dafny.CodePoint('e')
                    return _dafny.Map(coll47_)
                rhs206_ = d_889_v47_
                rhs207_ = _dafny.SeqWithoutIsStrInference([False])
                rhs208_ = d_897_v53_
                rhs209_ = -22
                rhs210_ = (iife101_()
                ) | (d_898_v54_)
                lhs162_ = d_888_v46_
                lhs163_ = default__.safeIndex(235, (d_888_v46_).length(0))
                lhs164_ = d_891_v51_
                lhs165_ = default__.safeIndex(405, (d_891_v51_).length(0))
                lhs162_[lhs163_] = rhs206_
                d_890_v48_ = rhs207_
                lhs164_[lhs165_] = rhs208_
                r2 = rhs209_
                d_898_v54_ = rhs210_
                d_901_v57_: _dafny.Seq
                d_901_v57_ = _dafny.SeqWithoutIsStrInference([d_889_v47_])
                d_902_v58_: _dafny.Array
                nw176_ = _dafny.Array(None, 16)
                nw176_[int(0)] = (d_901_v57_) + (d_901_v57_)
                nw176_[int(1)] = (d_901_v57_) + (_dafny.SeqWithoutIsStrInference([d_889_v47_]))
                nw176_[int(2)] = (d_901_v57_) + (d_901_v57_)
                nw176_[int(3)] = _dafny.SeqWithoutIsStrInference([d_889_v47_])
                nw176_[int(4)] = d_901_v57_
                nw176_[int(5)] = (d_901_v57_ if p1 else d_901_v57_)
                nw176_[int(6)] = d_901_v57_
                nw176_[int(7)] = d_901_v57_
                nw176_[int(8)] = d_901_v57_
                nw176_[int(9)] = (d_901_v57_) + (d_901_v57_)
                nw176_[int(10)] = d_901_v57_
                nw176_[int(11)] = d_901_v57_
                nw176_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_889_v47_])) + (d_901_v57_)
                nw176_[int(13)] = _dafny.SeqWithoutIsStrInference([d_889_v47_])
                nw176_[int(14)] = _dafny.SeqWithoutIsStrInference([(d_888_v46_)[default__.safeIndex(235, (d_888_v46_).length(0))], d_889_v47_])
                nw176_[int(15)] = (d_901_v57_) + (d_901_v57_)
                d_902_v58_ = nw176_
                index185_ = default__.safeIndex(642, (d_902_v58_).length(0))
                (d_902_v58_)[index185_] = (d_901_v57_) + (d_901_v57_)
                index186_ = default__.safeIndex(642, (d_902_v58_).length(0))
                (d_902_v58_)[index186_] = d_901_v57_
            elif True:
                index187_ = default__.safeIndex(738, (d_877_v35_).length(0))
                rhs211_ = (p2) - (p2)
                rhs212_ = not((((self).f22) - ((self).f22)) not in (d_879_v37_))
                lhs166_ = d_877_v35_
                lhs167_ = default__.safeIndex(738, (d_877_v35_).length(0))
                r2 = rhs211_
                lhs166_[lhs167_] = rhs212_
                (globalState).f2 = ((d_877_v35_)[default__.safeIndex(738, (d_877_v35_).length(0))]) and ((_dafny.MultiSet([p0])).issubset(d_826_v2_))
                d_903_v59_: _dafny.Map
                d_903_v59_ = _dafny.Map({self.f23: self.f23})
                d_904_v60_: _dafny.Map
                d_904_v60_ = _dafny.Map({((d_877_v35_)[default__.safeIndex(738, (d_877_v35_).length(0))]) and (self.f23): D1_DC4(self.f23, p2, d_903_v59_)})
                d_904_v60_ = (d_904_v60_) | (d_904_v60_)
                d_877_v35_ = d_877_v35_
                d_905_v61_: _dafny.Set
                d_905_v61_ = _dafny.Set({d_824_v0_})
                d_906_v62_: _dafny.Map
                d_906_v62_ = _dafny.Map({d_824_v0_: False})
                d_907_v64_: _dafny.Map
                def iife102_():
                    coll48_ = _dafny.Set()
                    compr_48_: str
                    for compr_48_ in (d_906_v62_).keys.Elements:
                        d_908_v63_: str = compr_48_
                        if (d_908_v63_) in (d_906_v62_):
                            coll48_ = coll48_.union(_dafny.Set([d_908_v63_]))
                    return _dafny.Set(coll48_)
                d_907_v64_ = _dafny.Map({(d_905_v61_) - (iife102_()
                ): d_877_v35_})
                rhs213_ = (p2) - (p2)
                rhs214_ = p2
                rhs215_ = (self).f22
                rhs216_ = ((d_907_v64_)[default__.fm27(globalState)] if (default__.fm27(globalState)) in (d_907_v64_) else d_877_v35_)
                rhs217_ = (self).f22
                r2 = rhs213_
                r2 = rhs214_
                r2 = rhs215_
                d_877_v35_ = rhs216_
                r2 = rhs217_
            r2 = default__.fm3(globalState)
            d_909_v65_: _dafny.Seq
            d_909_v65_ = _dafny.SeqWithoutIsStrInference([False])
            d_909_v65_ = (d_909_v65_) + (d_909_v65_)
        elif True:
            d_910_v66_: _dafny.Set
            d_910_v66_ = _dafny.Set({p2})
            d_911_v67_: _dafny.Set
            d_911_v67_ = _dafny.Set({len(d_910_v66_), p2})
            d_912_v68_: _dafny.Set
            d_912_v68_ = _dafny.Set({d_911_v67_, d_911_v67_, d_910_v66_})
            d_913_v69_: _dafny.Seq
            d_913_v69_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_910_v66_})])
            d_914_v70_: _dafny.Seq
            d_914_v70_ = _dafny.SeqWithoutIsStrInference([-236, p2])
            d_915_v71_: C1
            nw177_ = C1()
            nw177_.ctor__(p2, p2, self.f23, d_914_v70_, len(d_825_v1_))
            d_915_v71_ = nw177_
            d_916_v72_: _dafny.Map
            d_916_v72_ = _dafny.Map({d_915_v71_: (self).f22})
            d_917_v73_: _dafny.Map
            d_917_v73_ = _dafny.Map({d_916_v72_: _dafny.Set({_dafny.Set({(d_915_v71_).f30, (d_915_v71_).f30, (d_915_v71_).f30})})})
            d_918_v76_: _dafny.Seq
            d_918_v76_ = _dafny.SeqWithoutIsStrInference([d_910_v66_])
            d_919_v79_: _dafny.Array
            nw178_ = _dafny.Array(None, 24)
            nw178_[int(0)] = d_912_v68_
            nw178_[int(1)] = _dafny.Set({d_911_v67_})
            nw178_[int(2)] = (d_913_v69_)[default__.safeIndex(p2, len(d_913_v69_))]
            nw178_[int(3)] = (d_912_v68_) | (d_912_v68_)
            nw178_[int(4)] = (d_912_v68_) | (d_912_v68_)
            nw178_[int(5)] = default__.fm28(p1, (self).f22, p0, d_824_v0_, globalState)
            nw178_[int(6)] = d_912_v68_
            nw178_[int(7)] = d_912_v68_
            def iife103_():
                coll49_ = _dafny.Set()
                compr_49_: _dafny.Set
                for compr_49_ in (default__.fm29(p1, p2, p0, globalState)).Elements:
                    d_920_v74_: _dafny.Set = compr_49_
                    if (d_920_v74_) in (default__.fm29(p1, p2, p0, globalState)):
                        coll49_ = coll49_.union(_dafny.Set([d_920_v74_]))
                return _dafny.Set(coll49_)
            nw178_[int(8)] = ((d_917_v73_)[d_916_v72_] if (d_916_v72_) in (d_917_v73_) else iife103_()
            )
            def iife104_():
                coll50_ = _dafny.Set()
                compr_50_: _dafny.Set
                for compr_50_ in (_dafny.Map({d_911_v67_: (d_915_v71_).f30})).keys.Elements:
                    d_921_v75_: _dafny.Set = compr_50_
                    if (d_921_v75_) in (_dafny.Map({d_911_v67_: (d_915_v71_).f30})):
                        coll50_ = coll50_.union(_dafny.Set([d_921_v75_]))
                return _dafny.Set(coll50_)
            nw178_[int(9)] = iife104_()
            
            nw178_[int(10)] = d_912_v68_
            nw178_[int(11)] = (d_912_v68_ if p0 else d_912_v68_)
            nw178_[int(12)] = d_912_v68_
            nw178_[int(13)] = (d_912_v68_) - (_dafny.Set({d_911_v67_, _dafny.Set({p2}), d_911_v67_, d_910_v66_, d_910_v66_}))
            nw178_[int(14)] = d_912_v68_
            nw178_[int(15)] = d_912_v68_
            nw178_[int(16)] = d_912_v68_
            def iife105_():
                coll51_ = _dafny.Set()
                compr_51_: _dafny.Set
                for compr_51_ in (d_918_v76_).Elements:
                    d_922_v77_: _dafny.Set = compr_51_
                    if (d_922_v77_) in (d_918_v76_):
                        coll51_ = coll51_.union(_dafny.Set([d_922_v77_]))
                return _dafny.Set(coll51_)
            nw178_[int(17)] = iife105_()
            
            nw178_[int(18)] = d_912_v68_
            nw178_[int(19)] = d_912_v68_
            nw178_[int(20)] = d_912_v68_
            nw178_[int(21)] = d_912_v68_
            def iife106_():
                coll52_ = _dafny.Set()
                compr_52_: _dafny.Set
                for compr_52_ in (_dafny.Map({d_911_v67_: p0})).keys.Elements:
                    d_923_v78_: _dafny.Set = compr_52_
                    if (d_923_v78_) in (_dafny.Map({d_911_v67_: p0})):
                        coll52_ = coll52_.union(_dafny.Set([d_923_v78_]))
                return _dafny.Set(coll52_)
            nw178_[int(22)] = iife106_()
            
            nw178_[int(23)] = (_dafny.Set({d_910_v66_})) - (d_912_v68_)
            d_919_v79_ = nw178_
            index188_ = default__.safeIndex(440, (d_919_v79_).length(0))
            (d_919_v79_)[index188_] = _dafny.Set({d_911_v67_})
            index189_ = default__.safeIndex(440, (d_919_v79_).length(0))
            (d_919_v79_)[index189_] = default__.fm28(self.f23, (d_915_v71_).f29, not (p0) or (p1), d_824_v0_, globalState)
            d_924_v80_: _dafny.Map
            d_924_v80_ = _dafny.Map({p1: (self).f22})
            d_925_v81_: _dafny.Array
            nw179_ = _dafny.Array(_dafny.Seq({}), 4)
            d_925_v81_ = nw179_
            d_926_v82_: D2
            d_926_v82_ = D2_DC8(d_924_v80_, d_925_v81_, self.f23, p1, (d_915_v71_).f30)
            d_927_v83_: D6
            d_927_v83_ = D6_DC17(d_926_v82_)
            d_928_v84_: _dafny.MultiSet
            d_928_v84_ = _dafny.MultiSet([d_927_v83_, D6_DC17(d_926_v82_)])
            d_929_v85_: _dafny.Seq
            d_929_v85_ = _dafny.SeqWithoutIsStrInference([d_928_v84_])
            d_930_v86_: _dafny.Seq
            d_930_v86_ = d_929_v85_
            d_931_v87_: _dafny.Map
            d_931_v87_ = _dafny.Map({d_930_v86_: (d_824_v0_ if p0 else d_824_v0_)})
            rhs218_ = (d_914_v70_)[default__.safeIndex((self).f22, len(d_914_v70_))]
            rhs219_ = ((d_931_v87_)[d_930_v86_] if (d_930_v86_) in (d_931_v87_) else d_824_v0_)
            lhs168_ = globalState
            r2 = rhs218_
            lhs168_.f15 = rhs219_
            d_824_v0_ = _dafny.CodePoint('t')
            d_824_v0_ = d_824_v0_
            d_932_v88_: C1
            nw180_ = C1()
            nw180_.ctor__((d_915_v71_).f29, p2, p1, d_914_v70_, (self).f22)
            d_932_v88_ = nw180_
        (globalState).f2 = not(self.f23)
        d_933_v89_: _dafny.Map
        d_933_v89_ = _dafny.Map({p1: d_824_v0_})
        d_934_v90_: _dafny.Set
        d_934_v90_ = _dafny.Set({p0, not(default__.fm0(False, False, globalState))})
        d_935_v91_: _dafny.Array
        nw181_ = _dafny.Array(None, 3)
        nw181_[int(0)] = default__.safeDivisionInt(len(d_933_v89_), (self).f22)
        nw181_[int(1)] = default__.safeModuloInt(p2, p2)
        nw181_[int(2)] = len(d_934_v90_)
        d_935_v91_ = nw181_
        d_936_v92_: _dafny.Array
        nw182_ = _dafny.Array(False, 11)
        d_936_v92_ = nw182_
        d_937_v93_: _dafny.Map
        d_937_v93_ = _dafny.Map({d_936_v92_: default__.fm3(globalState)})
        d_938_v94_: _dafny.Seq
        d_938_v94_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_939_v95_: _dafny.Map
        d_939_v95_ = _dafny.Map({647: _dafny.SeqWithoutIsStrInference([(d_938_v94_)[default__.safeIndex(p2, len(d_938_v94_))], p1])})
        d_940_v96_: _dafny.Map
        d_940_v96_ = _dafny.Map({(0) - (((d_937_v93_)[d_936_v92_] if (d_936_v92_) in (d_937_v93_) else len(d_939_v95_))): (self).f22})
        index190_ = default__.safeIndex(578, (d_935_v91_).length(0))
        (d_935_v91_)[index190_] = len(d_940_v96_)
        index191_ = default__.safeIndex(578, (d_935_v91_).length(0))
        (d_935_v91_)[index191_] = p2
        d_941_i5_: int
        d_941_i5_ = 0
        with _dafny.label("7"):
            while p1:
                with _dafny.c_label("7"):
                    if (d_941_i5_) >= (100):
                        raise _dafny.Break("7")
                    d_941_i5_ = (d_941_i5_) + (1)
                    d_936_v92_ = d_936_v92_
                    d_942_v97_: _dafny.MultiSet
                    d_942_v97_ = _dafny.MultiSet([self.f23])
                    d_942_v97_ = (d_942_v97_) | (d_942_v97_)
                    d_943_v98_: _dafny.MultiSet
                    d_943_v98_ = _dafny.MultiSet([(d_935_v91_)[default__.safeIndex(578, (d_935_v91_).length(0))], len(d_938_v94_)])
                    d_944_v99_: _dafny.Map
                    d_944_v99_ = _dafny.Map({(self).f22: p1})
                    index192_ = default__.safeIndex(506, (d_936_v92_).length(0))
                    (d_936_v92_)[index192_] = ((d_943_v98_).set(len(d_944_v99_), default__.abs((0) - ((d_935_v91_)[default__.safeIndex(578, (d_935_v91_).length(0))])))).isdisjoint(d_943_v98_)
                    index193_ = default__.safeIndex(506, (d_936_v92_).length(0))
                    (d_936_v92_)[index193_] = ((391) - ((d_935_v91_)[default__.safeIndex(578, (d_935_v91_).length(0))])) < (default__.safeModuloInt((self).f22, (self).f22))
                    d_945_v100_: D4
                    d_945_v100_ = D4_DC13(p2, (0) - ((d_935_v91_)[default__.safeIndex(578, (d_935_v91_).length(0))]), (self).f22)
                    index194_ = default__.safeIndex(578, (d_935_v91_).length(0))
                    (d_935_v91_)[index194_] = (d_945_v100_).cf25
                    pass
            pass
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_936_v92_).length(0)):
            d_946_i6_: int = guard_loop_4_
            if (True) and (((0) <= (d_946_i6_)) and ((d_946_i6_) < ((d_936_v92_).length(0)))):
                (d_936_v92_)[(d_946_i6_)] = p0
        d_947_v101_: _dafny.MultiSet
        d_947_v101_ = _dafny.MultiSet([(self).f22, 628, 437, (0) - (p2), len(d_825_v1_)])
        d_948_v102_: _dafny.Set
        d_948_v102_ = _dafny.Set({d_947_v101_})
        def iife107_():
            coll53_ = _dafny.Set()
            compr_53_: _dafny.MultiSet
            for compr_53_ in (d_948_v102_).Elements:
                d_949_v103_: _dafny.MultiSet = compr_53_
                if (d_949_v103_) in (d_948_v102_):
                    coll53_ = coll53_.union(_dafny.Set([d_949_v103_]))
            return _dafny.Set(coll53_)
        r0 = (iife107_()
        ) | ((d_948_v102_ if self.f23 else d_948_v102_))
        r1 = p1
        r2 = (self).f22
        return r0, r1, r2

    def m7(self, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: bool = False
        r3: int = int(0)
        d_950_v0_: _dafny.Map
        d_950_v0_ = _dafny.Map({self.f23: self.f23})
        d_951_v1_: D1
        d_951_v1_ = D1_DC4(self.f23, (self).f22, d_950_v0_)
        d_952_v2_: _dafny.Seq
        d_952_v2_ = _dafny.SeqWithoutIsStrInference([d_951_v1_, d_951_v1_, d_951_v1_])
        d_952_v2_ = (_dafny.SeqWithoutIsStrInference([d_951_v1_])) + ((d_952_v2_) + (d_952_v2_))
        d_953_v3_: D8
        d_953_v3_ = D8_DC21((0) - ((0) - ((self).f22)), self.f23, self.f23, not(self.f23), (self).f22)
        d_954_i0_: int
        d_954_i0_ = 0
        with _dafny.label("8"):
            def lambda56_(source10_):
                if source10_.is_DC21:
                    d_959___mcc_h0_ = source10_.cf35
                    d_960___mcc_h1_ = source10_.cf36
                    d_961___mcc_h2_ = source10_.cf37
                    d_962___mcc_h3_ = source10_.cf38
                    d_963___mcc_h4_ = source10_.cf39
                    d_964_cf39_ = d_963___mcc_h4_
                    d_965_cf38_ = d_962___mcc_h3_
                    d_966_cf37_ = d_961___mcc_h2_
                    d_967_cf36_ = d_960___mcc_h1_
                    d_968_cf35_ = d_959___mcc_h0_
                    return d_966_cf37_
                elif source10_.is_DC22:
                    d_969___mcc_h5_ = source10_.cf40
                    d_970___mcc_h6_ = source10_.cf41
                    d_971___mcc_h7_ = source10_.cf42
                    d_972_cf42_ = d_971___mcc_h7_
                    d_973_cf41_ = d_970___mcc_h6_
                    d_974_cf40_ = d_969___mcc_h5_
                    return not(d_973_cf41_.f16)
                elif source10_.is_DC23:
                    return self.f23
                elif source10_.is_DC20:
                    d_975___mcc_h8_ = source10_.cf34
                    d_976_cf34_ = d_975___mcc_h8_
                    return self.f23
                elif True:
                    d_977___mcc_h9_ = source10_.cf43
                    d_978_cf43_ = d_977___mcc_h9_
                    return ((D3_DC9(_dafny.SeqWithoutIsStrInference([(self).f22]))).cf18) < (_dafny.SeqWithoutIsStrInference([(self).f22, (0) - ((self).f22)]))

            while lambda56_(d_953_v3_):
                with _dafny.c_label("8"):
                    if (d_954_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_954_i0_ = (d_954_i0_) + (1)
                    (self).f23 = ((self).f22) >= ((self).f22)
                    d_955_v4_: _dafny.Array
                    nw183_ = _dafny.Array(int(0), 25)
                    d_955_v4_ = nw183_
                    d_956_v5_: _dafny.Map
                    d_956_v5_ = _dafny.Map({(self).f22: (self).f22})
                    index195_ = default__.safeIndex(482, (d_955_v4_).length(0))
                    (d_955_v4_)[index195_] = (len(d_956_v5_)) - ((self).f22)
                    d_957_v6_: _dafny.Array
                    nw184_ = _dafny.Array(_dafny.Array(None, 0), 20)
                    d_957_v6_ = nw184_
                    d_958_v7_: _dafny.Array
                    nw185_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_958_v7_ = nw185_
                    index196_ = default__.safeIndex(246, (d_957_v6_).length(0))
                    (d_957_v6_)[index196_] = d_958_v7_
                    index197_ = default__.safeIndex(482, (d_955_v4_).length(0))
                    index198_ = default__.safeIndex(246, (d_957_v6_).length(0))
                    rhs220_ = (0) - ((self).f22)
                    rhs221_ = self.f23
                    rhs222_ = self.f23
                    rhs223_ = ((self).f22) == ((self).f22)
                    rhs224_ = d_958_v7_
                    lhs169_ = d_955_v4_
                    lhs170_ = default__.safeIndex(482, (d_955_v4_).length(0))
                    lhs171_ = globalState
                    lhs172_ = globalState
                    lhs173_ = globalState
                    lhs174_ = d_957_v6_
                    lhs175_ = default__.safeIndex(246, (d_957_v6_).length(0))
                    lhs169_[lhs170_] = rhs220_
                    lhs171_.f5 = rhs221_
                    lhs172_.f2 = rhs222_
                    lhs173_.f2 = rhs223_
                    lhs174_[lhs175_] = rhs224_
                    (globalState).f2 = (default__.safeDivisionInt((self).f22, (self).f22)) >= ((d_955_v4_)[default__.safeIndex(482, (d_955_v4_).length(0))])
                    (globalState).f15 = _dafny.CodePoint('f')
                    pass
            pass
        r0 = default__.safeModuloInt(len(_dafny.Map({self.f23: False})), ((self).f22) * (858))
        hi6_ = (self).f22
        for d_979_i1_ in range((self).f22, hi6_):
            r2 = not (self.f23) or (self.f23)
            r2 = self.f23
            d_980_v8_: _dafny.Seq
            d_980_v8_ = _dafny.SeqWithoutIsStrInference([self.f23])
            d_980_v8_ = _dafny.SeqWithoutIsStrInference([(d_980_v8_) <= ((d_980_v8_).set(default__.safeIndex(default__.fm3(globalState), len(d_980_v8_)), self.f23))])
            d_981_v9_: _dafny.Array
            nw186_ = _dafny.Array(None, 7)
            nw186_[int(0)] = d_953_v3_
            nw186_[int(1)] = d_953_v3_
            nw186_[int(2)] = D8_DC21((self).f22, self.f23, self.f23, True, d_979_i1_)
            nw186_[int(3)] = d_953_v3_
            nw186_[int(4)] = d_953_v3_
            nw186_[int(5)] = d_953_v3_
            nw186_[int(6)] = d_953_v3_
            d_981_v9_ = nw186_
            d_982_v10_: C2
            nw187_ = C2()
            nw187_.ctor__((self).f22)
            d_982_v10_ = nw187_
            d_983_v11_: _dafny.Map
            d_983_v11_ = _dafny.Map({(self).f22: d_982_v10_})
            pat_let_tv23_ = d_983_v11_
            index199_ = default__.safeIndex(383, (d_981_v9_).length(0))
            def iife108_(_pat_let27_0):
                def iife109_(d_984_dt__update__tmp_h0_):
                    def iife110_(_pat_let28_0):
                        def iife111_(d_985_dt__update_hcf35_h0_):
                            return D8_DC21(d_985_dt__update_hcf35_h0_, (d_984_dt__update__tmp_h0_).cf36, (d_984_dt__update__tmp_h0_).cf37, (d_984_dt__update__tmp_h0_).cf38, (d_984_dt__update__tmp_h0_).cf39)
                        return iife111_(_pat_let28_0)
                    return iife110_(len(pat_let_tv23_))
                return iife109_(_pat_let27_0)
            (d_981_v9_)[index199_] = (iife108_(d_953_v3_) if True else default__.fm26(globalState))
            d_986_v12_: _dafny.Array
            nw188_ = _dafny.Array(int(0), 24)
            d_986_v12_ = nw188_
            index200_ = default__.safeIndex(152, (d_986_v12_).length(0))
            (d_986_v12_)[index200_] = ((d_982_v10_).f25 if self.f23 else (0) - ((d_982_v10_).f25))
            d_987_v13_: _dafny.MultiSet
            d_987_v13_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([self.f23])) + (d_980_v8_)])
            d_988_v14_: _dafny.Map
            d_988_v14_ = _dafny.Map({self.f23: 491})
            d_989_v15_: _dafny.Map
            d_989_v15_ = _dafny.Map({d_979_i1_: self.f23})
            d_990_v16_: _dafny.Array
            nw189_ = _dafny.Array(_dafny.Seq({}), 23)
            d_990_v16_ = nw189_
            d_991_v17_: D2
            d_991_v17_ = D2_DC8(d_988_v14_, d_990_v16_, self.f23, self.f23, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjqdbdskh")))))
            d_992_v18_: _dafny.Array
            nw190_ = _dafny.Array(None, 25)
            nw190_[int(0)] = self.f23
            nw190_[int(1)] = self.f23
            nw190_[int(2)] = self.f23
            nw190_[int(3)] = self.f23
            nw190_[int(4)] = self.f23
            nw190_[int(5)] = not(self.f23)
            nw190_[int(6)] = (d_988_v14_) != (d_988_v14_)
            nw190_[int(7)] = False
            nw190_[int(8)] = (self.f23 if (d_980_v8_)[default__.safeIndex((d_982_v10_).f25, len(d_980_v8_))] else ((d_989_v15_)[d_979_i1_] if (d_979_i1_) in (d_989_v15_) else (self).fm9(d_979_i1_, globalState)))
            nw190_[int(9)] = self.f23
            nw190_[int(10)] = False
            nw190_[int(11)] = self.f23
            nw190_[int(12)] = (499) != (d_979_i1_)
            nw190_[int(13)] = self.f23
            nw190_[int(14)] = self.f23
            nw190_[int(15)] = not (self.f23) or (self.f23)
            nw190_[int(16)] = self.f23
            nw190_[int(17)] = True
            nw190_[int(18)] = not((d_979_i1_) > (59))
            nw190_[int(19)] = self.f23
            nw190_[int(20)] = ((d_991_v17_).cf15 if self.f23 else self.f23)
            nw190_[int(21)] = self.f23
            nw190_[int(22)] = not(not(self.f23))
            nw190_[int(23)] = self.f23
            nw190_[int(24)] = self.f23
            d_992_v18_ = nw190_
            d_993_v19_: _dafny.MultiSet
            d_993_v19_ = _dafny.MultiSet([d_979_i1_])
            d_994_v20_: _dafny.Map
            d_994_v20_ = _dafny.Map({d_992_v18_: d_992_v18_})
            index201_ = default__.safeIndex(383, (d_981_v9_).length(0))
            index202_ = default__.safeIndex(152, (d_986_v12_).length(0))
            rhs225_ = D8_DC21((0) - (((0) - ((self).f22) if True else (0) - (d_979_i1_))), (d_993_v19_) != (d_993_v19_), self.f23, self.f23, d_979_i1_)
            rhs226_ = (self).f22
            rhs227_ = d_979_i1_
            rhs228_ = d_987_v13_
            rhs229_ = ((d_994_v20_)[d_992_v18_] if (d_992_v18_) in (d_994_v20_) else d_992_v18_)
            lhs176_ = d_981_v9_
            lhs177_ = default__.safeIndex(383, (d_981_v9_).length(0))
            lhs178_ = d_986_v12_
            lhs179_ = default__.safeIndex(152, (d_986_v12_).length(0))
            lhs176_[lhs177_] = rhs225_
            r0 = rhs226_
            lhs178_[lhs179_] = rhs227_
            d_987_v13_ = rhs228_
            d_992_v18_ = rhs229_
        d_995_v21_: D4
        d_995_v21_ = D4_DC12(self.f23, self.f23, (self).f22)
        source11_ = d_995_v21_
        if source11_.is_DC12:
            d_996___mcc_h10_ = source11_.cf21
            d_997___mcc_h11_ = source11_.cf22
            d_998___mcc_h12_ = source11_.cf23
            d_999_cf23_ = d_998___mcc_h12_
            d_1000_cf22_ = d_997___mcc_h11_
            d_1001_cf21_ = d_996___mcc_h10_
            d_1002_v22_: _dafny.Seq
            d_1002_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eh"))
            r1 = d_1002_v22_
            d_1003_v23_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Seq({}), 1)
            d_1003_v23_ = nw191_
            d_1004_v24_: _dafny.Seq
            d_1004_v24_ = _dafny.SeqWithoutIsStrInference([not(d_1001_cf21_)])
            index203_ = default__.safeIndex(118, (d_1003_v23_).length(0))
            (d_1003_v23_)[index203_] = d_1004_v24_
            index204_ = default__.safeIndex(118, (d_1003_v23_).length(0))
            (d_1003_v23_)[index204_] = d_1004_v24_
            r2 = (d_1002_v22_) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) + (d_1002_v22_))
            d_1005_v25_: _dafny.Map
            d_1005_v25_ = _dafny.Map({False: (self).f22})
            d_1005_v25_ = (d_1005_v25_).set((d_1001_cf21_) and (True), (self).f22)
        elif source11_.is_DC13:
            d_1006___mcc_h13_ = source11_.cf24
            d_1007___mcc_h14_ = source11_.cf25
            d_1008___mcc_h15_ = source11_.cf26
            d_1009_cf26_ = d_1008___mcc_h15_
            d_1010_cf25_ = d_1007___mcc_h14_
            d_1011_cf24_ = d_1006___mcc_h13_
            d_1012_v26_: _dafny.Map
            d_1012_v26_ = _dafny.Map({d_1010_cf25_: False})
            d_1013_v27_: _dafny.MultiSet
            d_1013_v27_ = _dafny.MultiSet([d_1011_cf24_, len(d_1012_v26_)])
            d_1014_v28_: _dafny.Seq
            d_1014_v28_ = _dafny.SeqWithoutIsStrInference([len(default__.fm23(d_1013_v27_, self.f23, globalState)), len(d_1012_v26_)])
            d_1015_v29_: _dafny.Array
            nw192_ = _dafny.Array(False, 28)
            d_1015_v29_ = nw192_
            d_1016_v30_: C0
            nw193_ = C0()
            nw193_.ctor__(len((d_1014_v28_) + (_dafny.SeqWithoutIsStrInference([d_1010_cf25_ for d_1017_i2_ in range(default__.abs(567))]))), d_1015_v29_)
            d_1016_v30_ = nw193_
            (globalState).f2 = self.f23
            (globalState).f2 = self.f23
            if self.f23:
                d_1018_v31_: _dafny.Map
                d_1018_v31_ = _dafny.Map({self.f23: d_1011_cf24_})
                d_1019_v32_: D3
                d_1019_v32_ = D3_DC9(_dafny.SeqWithoutIsStrInference([len(d_1018_v31_), (d_1016_v30_).f26, (d_1016_v30_).f26]))
                pat_let_tv24_ = d_1014_v28_
                pat_let_tv25_ = d_1014_v28_
                pat_let_tv26_ = d_1014_v28_
                d_1020_v33_: _dafny.Array
                nw194_ = _dafny.Array(None, 15)
                nw194_[int(0)] = d_1019_v32_
                nw194_[int(1)] = d_1019_v32_
                nw194_[int(2)] = d_1019_v32_
                def iife112_(_pat_let29_0):
                    def iife113_(d_1021_dt__update__tmp_h1_):
                        def iife114_(_pat_let30_0):
                            def iife115_(d_1022_dt__update_hcf18_h0_):
                                return D3_DC9(d_1022_dt__update_hcf18_h0_)
                            return iife115_(_pat_let30_0)
                        return iife114_(pat_let_tv24_)
                    return iife113_(_pat_let29_0)
                nw194_[int(3)] = iife112_(d_1019_v32_)
                nw194_[int(4)] = d_1019_v32_
                def iife116_(_pat_let31_0):
                    def iife117_(d_1023_dt__update__tmp_h2_):
                        def iife118_(_pat_let32_0):
                            def iife119_(d_1024_dt__update_hcf18_h1_):
                                return D3_DC9(d_1024_dt__update_hcf18_h1_)
                            return iife119_(_pat_let32_0)
                        return iife118_(pat_let_tv25_)
                    return iife117_(_pat_let31_0)
                nw194_[int(5)] = iife116_(d_1019_v32_)
                def iife120_(_pat_let33_0):
                    def iife121_(d_1025_dt__update__tmp_h3_):
                        def iife122_(_pat_let34_0):
                            def iife123_(d_1026_dt__update_hcf18_h2_):
                                return D3_DC9(d_1026_dt__update_hcf18_h2_)
                            return iife123_(_pat_let34_0)
                        return iife122_(pat_let_tv26_)
                    return iife121_(_pat_let33_0)
                nw194_[int(6)] = iife120_(d_1019_v32_)
                nw194_[int(7)] = (d_1019_v32_ if self.f23 else d_1019_v32_)
                nw194_[int(8)] = (d_1019_v32_ if self.f23 else d_1019_v32_)
                nw194_[int(9)] = d_1019_v32_
                nw194_[int(10)] = d_1019_v32_
                nw194_[int(11)] = d_1019_v32_
                nw194_[int(12)] = d_1019_v32_
                nw194_[int(13)] = d_1019_v32_
                nw194_[int(14)] = d_1019_v32_
                d_1020_v33_ = nw194_
                index205_ = default__.safeIndex(975, (d_1020_v33_).length(0))
                (d_1020_v33_)[index205_] = d_1019_v32_
                index206_ = default__.safeIndex(975, (d_1020_v33_).length(0))
                (d_1020_v33_)[index206_] = d_1019_v32_
                (d_1016_v30_).f27 = (d_1015_v29_ if (default__.fm17(globalState)) == (d_1014_v28_) else d_1015_v29_)
                d_1027_v34_: _dafny.Seq
                d_1027_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bglc"))
                d_1028_v35_: C3
                nw195_ = C3()
                nw195_.ctor__(d_1027_v34_, self.f23, (d_1014_v28_) + (d_1014_v28_))
                d_1028_v35_ = nw195_
                d_1028_v35_ = d_1028_v35_
                (globalState).f2 = self.f23
                d_1029_v36_: C0
                nw196_ = C0()
                nw196_.ctor__(d_1011_cf24_, d_1015_v29_)
                d_1029_v36_ = nw196_
            elif True:
                r3 = (d_1010_cf25_) - ((d_1016_v30_).f26)
                r1 = default__.fm1(globalState)
                d_1030_v37_: C2
                nw197_ = C2()
                nw197_.ctor__((d_1016_v30_).f26)
                d_1030_v37_ = nw197_
                d_1011_cf24_ = (d_1016_v30_).f26
                d_1031_v38_: _dafny.Array
                nw198_ = _dafny.Array(_dafny.Seq({}), 2)
                d_1031_v38_ = nw198_
                index207_ = default__.safeIndex(608, (d_1031_v38_).length(0))
                (d_1031_v38_)[index207_] = d_1014_v28_
                index208_ = default__.safeIndex(608, (d_1031_v38_).length(0))
                (d_1031_v38_)[index208_] = d_1014_v28_
        elif True:
            d_1032___mcc_h16_ = source11_.cf20
            d_1033_cf20_ = d_1032___mcc_h16_
            (globalState).f2 = self.f23
            d_1034_v39_: _dafny.Set
            d_1034_v39_ = _dafny.Set({self.f23})
            if ((_dafny.Set({self.f23, self.f23})) | (d_1034_v39_)).issubset(d_1034_v39_):
                d_1035_v40_: _dafny.Array
                nw199_ = _dafny.Array(int(0), 12)
                d_1035_v40_ = nw199_
                nw200_ = _dafny.Array(int(0), 10)
                d_1035_v40_ = nw200_
                d_1036_v41_: _dafny.Map
                d_1036_v41_ = _dafny.Map({False: (self).f22})
                d_1037_v42_: _dafny.Seq
                d_1037_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap"))
                d_1036_v41_ = (d_1036_v41_).set((self).fm9((self).f22, globalState), len(d_1037_v42_))
                r0 = (self).f22
                rhs230_ = True
                rhs231_ = d_1035_v40_
                rhs232_ = (self).f22
                rhs233_ = self.f23
                lhs180_ = globalState
                lhs181_ = globalState
                lhs180_.f5 = rhs230_
                d_1035_v40_ = rhs231_
                r0 = rhs232_
                lhs181_.f2 = rhs233_
                r2 = self.f23
            elif True:
                r2 = self.f23
                d_1038_v43_: D0
                d_1038_v43_ = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
                d_1039_v44_: _dafny.Map
                d_1039_v44_ = _dafny.Map({d_1038_v43_: self.f23})
                d_1040_v46_: _dafny.Set
                d_1040_v46_ = _dafny.Set({d_1038_v43_})
                d_1041_v47_: _dafny.Set
                def iife124_():
                    coll54_ = _dafny.Map()
                    compr_54_: D0
                    for compr_54_ in (d_1040_v46_).Elements:
                        d_1042_v45_: D0 = compr_54_
                        if (d_1042_v45_) in (d_1040_v46_):
                            coll54_[d_1042_v45_] = self.f23
                    return _dafny.Map(coll54_)
                d_1041_v47_ = _dafny.Set({iife124_()
                })
                (self).f23 = (d_1039_v44_) not in (d_1041_v47_)
                d_1043_v48_: D1
                d_1043_v48_ = D1_DC5((35) - (174), self.f23)
                d_1044_v49_: _dafny.MultiSet
                d_1044_v49_ = _dafny.MultiSet([(self.f23 if True else self.f23)])
                d_1045_v50_: _dafny.Map
                d_1045_v50_ = _dafny.Map({222: self.f23})
                d_1046_v51_: _dafny.Map
                d_1046_v51_ = _dafny.Map({self.f23: (self).f22})
                rhs234_ = default__.fm0(False, (d_1045_v50_) != ((d_1045_v50_).set((self).f22, self.f23)), globalState)
                rhs235_ = d_1043_v48_
                rhs236_ = (d_1044_v49_).set(self.f23, default__.abs(((d_1046_v51_)[self.f23] if (self.f23) in (d_1046_v51_) else (self).f22)))
                lhs182_ = globalState
                lhs182_.f2 = rhs234_
                d_1043_v48_ = rhs235_
                d_1044_v49_ = rhs236_
                d_1047_v52_: _dafny.Array
                nw201_ = _dafny.Array(None, 3)
                nw201_[int(0)] = self.f23
                nw201_[int(1)] = self.f23
                nw201_[int(2)] = (self.f23) and (self.f23)
                d_1047_v52_ = nw201_
                d_1047_v52_ = (d_1047_v52_ if self.f23 else d_1047_v52_)
                d_1048_v53_: _dafny.Array
                nw202_ = _dafny.Array(int(0), 19)
                d_1048_v53_ = nw202_
                index209_ = default__.safeIndex(328, (d_1048_v53_).length(0))
                (d_1048_v53_)[index209_] = len(_dafny.Map({True: self.f23}))
                d_1049_v54_: C2
                nw203_ = C2()
                nw203_.ctor__(len(d_950_v0_))
                d_1049_v54_ = nw203_
                d_1050_v55_: _dafny.Array
                nw204_ = _dafny.Array(_dafny.MultiSet({}), 5)
                d_1050_v55_ = nw204_
                d_1051_v56_: _dafny.MultiSet
                d_1051_v56_ = _dafny.MultiSet([D1_DC5((d_1049_v54_).f25, self.f23), d_1043_v48_])
                index210_ = default__.safeIndex(862, (d_1050_v55_).length(0))
                (d_1050_v55_)[index210_] = d_1051_v56_
                d_1052_v57_: _dafny.Seq
                d_1052_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fs"))
                index211_ = default__.safeIndex(328, (d_1048_v53_).length(0))
                index212_ = default__.safeIndex(862, (d_1050_v55_).length(0))
                rhs237_ = (self).f22
                rhs238_ = d_1049_v54_
                rhs239_ = (default__.fm30(self.f23, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ut")), True, self.f23, globalState)) - (d_1051_v56_)
                rhs240_ = self.f23
                rhs241_ = default__.safeDivisionInt(len(d_1052_v57_), ((0) - ((d_1049_v54_).f25)) - (len(_dafny.SeqWithoutIsStrInference([self.f23, not(self.f23)]))))
                lhs183_ = d_1048_v53_
                lhs184_ = default__.safeIndex(328, (d_1048_v53_).length(0))
                lhs185_ = d_1050_v55_
                lhs186_ = default__.safeIndex(862, (d_1050_v55_).length(0))
                lhs187_ = self
                lhs183_[lhs184_] = rhs237_
                d_1049_v54_ = rhs238_
                lhs185_[lhs186_] = rhs239_
                lhs187_.f23 = rhs240_
                r0 = rhs241_
            d_1053_v58_: _dafny.Seq
            d_1053_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gao"))
            d_1054_v59_: str
            d_1054_v59_ = _dafny.CodePoint('x')
            default__.m0(len(((d_1053_v58_ if self.f23 else d_1053_v58_)).set(default__.safeIndex((0) - ((self).f22), len((d_1053_v58_ if self.f23 else d_1053_v58_))), d_1054_v59_)), d_1053_v58_, (self).f22, default__.safeModuloInt(len(_dafny.Set({self.f23})), 992), globalState)
            d_1055_v60_: _dafny.Seq
            d_1055_v60_ = _dafny.SeqWithoutIsStrInference([(self).f22])
            rhs242_ = (self).f22
            rhs243_ = self.f23
            rhs244_ = _dafny.CodePoint('w')
            rhs245_ = ((self).f22 if self.f23 else (self).f22)
            rhs246_ = (d_1055_v60_) + ((d_1055_v60_) + (d_1055_v60_))
            lhs188_ = globalState
            lhs189_ = globalState
            r0 = rhs242_
            lhs188_.f2 = rhs243_
            lhs189_.f15 = rhs244_
            r0 = rhs245_
            d_1055_v60_ = rhs246_
        d_1056_v61_: _dafny.Array
        nw205_ = _dafny.Array(D2.default()(), 29)
        d_1056_v61_ = nw205_
        d_1057_v62_: _dafny.Map
        d_1057_v62_ = _dafny.Map({self.f23: (self).f22})
        d_1058_v63_: _dafny.Array
        nw206_ = _dafny.Array(_dafny.Seq({}), 3)
        d_1058_v63_ = nw206_
        d_1059_v64_: D2
        d_1059_v64_ = D2_DC8(d_1057_v62_, d_1058_v63_, self.f23, self.f23, (self).f22)
        index213_ = default__.safeIndex(594, (d_1056_v61_).length(0))
        (d_1056_v61_)[index213_] = d_1059_v64_
        index214_ = default__.safeIndex(594, (d_1056_v61_).length(0))
        (d_1056_v61_)[index214_] = d_1059_v64_
        r0 = -398
        d_1060_v65_: _dafny.Seq
        d_1060_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
        r1 = d_1060_v65_
        r2 = self.f23
        d_1061_v66_: _dafny.Seq
        d_1061_v66_ = _dafny.SeqWithoutIsStrInference([self.f23, self.f23, self.f23, (self.f23) == (self.f23), ((d_950_v0_)[False] if (False) in (d_950_v0_) else self.f23)])
        r3 = len(d_1061_v66_)
        return r0, r1, r2, r3

    @property
    def f22(self):
        return self._f22

class C5(T0):
    def  __init__(self):
        self._f16: bool = False
        self._f17: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f16, f17):
        (self).f16 = f16
        (self)._f17 = f17

    def fm7(self, p0, globalState):
        if self.f16:
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])
        elif True:
            return (D0_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1062_i0_ in range(default__.abs(117))]))).cf1

    def fm8(self, p0, globalState):
        return ((self).f17) <= ((D3_DC9((self).f17)).cf18)

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        (self).f16 = not (p2) or (default__.fm0(p2, p1, globalState))
        d_1063_v0_: _dafny.Array
        nw207_ = _dafny.Array(int(0), 29)
        d_1063_v0_ = nw207_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1063_v0_).length(0)):
            d_1064_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_1064_i0_)) and ((d_1064_i0_) < ((d_1063_v0_).length(0)))):
                (d_1063_v0_)[(d_1064_i0_)] = (d_1064_i0_) + (8)
        d_1065_v1_: C4
        nw208_ = C4()
        nw208_.ctor__(p3, True)
        d_1065_v1_ = nw208_
        d_1066_v2_: _dafny.Array
        nw209_ = _dafny.Array(_dafny.Seq({}), 15)
        d_1066_v2_ = nw209_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1066_v2_).length(0)):
            d_1067_i1_: int = guard_loop_6_
            if (True) and (((0) <= (d_1067_i1_)) and ((d_1067_i1_) < ((d_1066_v2_).length(0)))):
                (d_1066_v2_)[(d_1067_i1_)] = _dafny.SeqWithoutIsStrInference([59 for d_1068_i2_ in range(default__.abs(-728))])
        d_1069_v3_: _dafny.Array
        nw210_ = _dafny.Array(D0.default()(), 11)
        d_1069_v3_ = nw210_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1069_v3_).length(0)):
            d_1070_i3_: int = guard_loop_7_
            if (True) and (((0) <= (d_1070_i3_)) and ((d_1070_i3_) < ((d_1069_v3_).length(0)))):
                (d_1069_v3_)[(d_1070_i3_)] = D0_DC0(p0)
        hi7_ = (d_1065_v1_).f22
        for d_1071_i4_ in range(p3, hi7_):
            d_1072_v4_: int
            d_1072_v4_ = 577
            d_1072_v4_ = p3
            d_1073_v5_: int
            d_1074_v6_: _dafny.Seq
            d_1075_v7_: bool
            d_1076_v8_: int
            out6_: int
            out7_: _dafny.Seq
            out8_: bool
            out9_: int
            out6_, out7_, out8_, out9_ = (d_1065_v1_).m7(globalState)
            d_1073_v5_ = out6_
            d_1074_v6_ = out7_
            d_1075_v7_ = out8_
            d_1076_v8_ = out9_
            d_1077_v9_: _dafny.Array
            nw211_ = _dafny.Array(False, 14)
            d_1077_v9_ = nw211_
            index215_ = default__.safeIndex(633, (d_1077_v9_).length(0))
            (d_1077_v9_)[index215_] = p2
            index216_ = default__.safeIndex(818, (d_1063_v0_).length(0))
            (d_1063_v0_)[index216_] = (d_1065_v1_).f22
            d_1078_v10_: _dafny.Set
            d_1078_v10_ = _dafny.Set({d_1071_i4_})
            d_1079_v11_: _dafny.Array
            nw212_ = _dafny.Array(None, 6)
            nw212_[int(0)] = d_1078_v10_
            nw212_[int(1)] = d_1078_v10_
            nw212_[int(2)] = default__.fm31(self.f16, globalState)
            nw212_[int(3)] = default__.fm31(p2, globalState)
            nw212_[int(4)] = default__.fm31(d_1065_v1_.f23, globalState)
            nw212_[int(5)] = (d_1078_v10_).intersection(d_1078_v10_)
            d_1079_v11_ = nw212_
            index217_ = default__.safeIndex(458, (d_1079_v11_).length(0))
            (d_1079_v11_)[index217_] = d_1078_v10_
            d_1080_v12_: str
            d_1080_v12_ = _dafny.CodePoint('h')
            index218_ = default__.safeIndex(633, (d_1077_v9_).length(0))
            index219_ = default__.safeIndex(818, (d_1063_v0_).length(0))
            index220_ = default__.safeIndex(458, (d_1079_v11_).length(0))
            rhs247_ = self.f16
            rhs248_ = d_1080_v12_
            rhs249_ = 716
            rhs250_ = d_1078_v10_
            lhs190_ = d_1077_v9_
            lhs191_ = default__.safeIndex(633, (d_1077_v9_).length(0))
            lhs192_ = globalState
            lhs193_ = d_1063_v0_
            lhs194_ = default__.safeIndex(818, (d_1063_v0_).length(0))
            lhs195_ = d_1079_v11_
            lhs196_ = default__.safeIndex(458, (d_1079_v11_).length(0))
            lhs190_[lhs191_] = rhs247_
            lhs192_.f15 = rhs248_
            lhs193_[lhs194_] = rhs249_
            lhs195_[lhs196_] = rhs250_
            d_1081_v13_: _dafny.Map
            d_1081_v13_ = _dafny.Map({(d_1065_v1_).f22: d_1073_v5_})
            d_1082_v14_: _dafny.Array
            nw213_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_1082_v14_ = nw213_
            d_1083_v15_: _dafny.Seq
            d_1083_v15_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_1084_v16_: _dafny.Map
            d_1084_v16_ = _dafny.Map({not(self.f16): d_1082_v14_})
            d_1085_v17_: _dafny.Map
            d_1085_v17_ = _dafny.Map({p3: ((d_1084_v16_)[p1] if (p1) in (d_1084_v16_) else d_1082_v14_)})
            index221_ = default__.safeIndex(633, (d_1077_v9_).length(0))
            rhs251_ = ((d_1083_v15_)[default__.safeIndex(d_1076_v8_, len(d_1083_v15_))]) and ((self).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjxld")), globalState))
            rhs252_ = _dafny.Map({d_1076_v8_: (0) - ((0) - (default__.safeDivisionInt(d_1073_v5_, d_1072_v4_)))})
            rhs253_ = ((d_1085_v17_)[d_1071_i4_] if (d_1071_i4_) in (d_1085_v17_) else (d_1082_v14_ if p2 else d_1082_v14_))
            lhs197_ = d_1077_v9_
            lhs198_ = default__.safeIndex(633, (d_1077_v9_).length(0))
            lhs197_[lhs198_] = rhs251_
            d_1081_v13_ = rhs252_
            d_1082_v14_ = rhs253_
        d_1086_v18_: _dafny.Seq
        d_1086_v18_ = _dafny.SeqWithoutIsStrInference([self.f16])
        r0 = d_1086_v18_
        return r0

    def m5(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        if not (self.f16) or (self.f16):
            if self.f16:
                d_1087_v0_: _dafny.Seq
                d_1087_v0_ = _dafny.SeqWithoutIsStrInference([default__.fm19(125, globalState)])
                d_1088_v1_: int
                d_1088_v1_ = 42
                d_1089_v2_: _dafny.Seq
                d_1089_v2_ = _dafny.SeqWithoutIsStrInference([self.f16, not(self.f16), self.f16, self.f16])
                d_1090_v3_: _dafny.MultiSet
                d_1090_v3_ = _dafny.MultiSet([self.f16, self.f16])
                d_1091_v4_: _dafny.Map
                d_1091_v4_ = _dafny.Map({(d_1090_v3_).cardinality: not(self.f16)})
                d_1092_v5_: str
                d_1092_v5_ = _dafny.CodePoint('y')
                d_1093_v6_: _dafny.MultiSet
                d_1093_v6_ = _dafny.MultiSet([d_1092_v5_])
                d_1094_v7_: _dafny.Map
                d_1094_v7_ = _dafny.Map({True: d_1088_v1_})
                d_1095_v8_: _dafny.Array
                nw214_ = _dafny.Array(None, 28)
                nw214_[int(0)] = d_1088_v1_
                nw214_[int(1)] = 203
                nw214_[int(2)] = default__.safeModuloInt(d_1088_v1_, d_1088_v1_)
                nw214_[int(3)] = (d_1088_v1_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([d_1088_v1_]))))
                nw214_[int(4)] = d_1088_v1_
                nw214_[int(5)] = d_1088_v1_
                nw214_[int(6)] = d_1088_v1_
                nw214_[int(7)] = d_1088_v1_
                nw214_[int(8)] = (d_1088_v1_) - (len(_dafny.SeqWithoutIsStrInference([self.f16, self.f16])))
                nw214_[int(9)] = d_1088_v1_
                nw214_[int(10)] = len(d_1089_v2_)
                nw214_[int(11)] = (187) + (d_1088_v1_)
                nw214_[int(12)] = d_1088_v1_
                nw214_[int(13)] = d_1088_v1_
                nw214_[int(14)] = d_1088_v1_
                nw214_[int(15)] = (len(d_1091_v4_)) * (-267)
                nw214_[int(16)] = d_1088_v1_
                nw214_[int(17)] = d_1088_v1_
                nw214_[int(18)] = (d_1090_v3_).cardinality
                nw214_[int(19)] = (d_1088_v1_) - (d_1088_v1_)
                nw214_[int(20)] = (d_1088_v1_) - (((self).f17)[default__.safeIndex(d_1088_v1_, len((self).f17))])
                nw214_[int(21)] = -718
                nw214_[int(22)] = d_1088_v1_
                nw214_[int(23)] = d_1088_v1_
                nw214_[int(24)] = d_1088_v1_
                nw214_[int(25)] = default__.safeDivisionInt(((d_1093_v6_)[d_1092_v5_] if (d_1092_v5_) in (d_1093_v6_) else d_1088_v1_), len(d_1094_v7_))
                nw214_[int(26)] = (d_1088_v1_) * (len((self).f17))
                nw214_[int(27)] = d_1088_v1_
                d_1095_v8_ = nw214_
                index222_ = default__.safeIndex(522, (d_1095_v8_).length(0))
                (d_1095_v8_)[index222_] = d_1088_v1_
                d_1096_v9_: _dafny.Map
                d_1096_v9_ = _dafny.Map({self.f16: d_1089_v2_})
                index223_ = default__.safeIndex(522, (d_1095_v8_).length(0))
                rhs254_ = ((d_1087_v0_ if self.f16 else d_1087_v0_)) + (default__.fm32(102, False, self.f16, d_1092_v5_, globalState))
                rhs255_ = ((len(d_1096_v9_)) * (d_1088_v1_)) - (default__.fm3(globalState))
                lhs199_ = d_1095_v8_
                lhs200_ = default__.safeIndex(522, (d_1095_v8_).length(0))
                d_1087_v0_ = rhs254_
                lhs199_[lhs200_] = rhs255_
                index224_ = default__.safeIndex(522, (d_1095_v8_).length(0))
                (d_1095_v8_)[index224_] = (d_1095_v8_)[default__.safeIndex(522, (d_1095_v8_).length(0))]
                d_1097_v10_: _dafny.Seq
                d_1097_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klmwlqbvh"))
                default__.m0(d_1088_v1_, d_1097_v10_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvw"))) + (d_1097_v10_)), default__.safeDivisionInt((d_1095_v8_)[default__.safeIndex(522, (d_1095_v8_).length(0))], (d_1095_v8_)[default__.safeIndex(522, (d_1095_v8_).length(0))]), globalState)
                index225_ = default__.safeIndex(522, (d_1095_v8_).length(0))
                (d_1095_v8_)[index225_] = ((d_1095_v8_)[default__.safeIndex(522, (d_1095_v8_).length(0))]) * ((d_1088_v1_) * ((d_1095_v8_)[default__.safeIndex(522, (d_1095_v8_).length(0))]))
                index226_ = default__.safeIndex(522, (d_1095_v8_).length(0))
                (d_1095_v8_)[index226_] = (0) - (d_1088_v1_)
            elif True:
                d_1098_v11_: int
                d_1098_v11_ = 0
                d_1099_v12_: _dafny.Seq
                d_1099_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sf"))
                d_1098_v11_ = len(((d_1099_v12_) + (default__.fm1(globalState))) + (d_1099_v12_))
                d_1098_v11_ = d_1098_v11_
                d_1098_v11_ = d_1098_v11_
                d_1100_v13_: D0
                d_1100_v13_ = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvkvrqxb")))
                d_1101_v14_: str
                d_1101_v14_ = _dafny.CodePoint('s')
                d_1102_v15_: _dafny.MultiSet
                d_1102_v15_ = _dafny.MultiSet([len((d_1099_v12_ if self.f16 else ((d_1100_v13_).cf1).set(default__.safeIndex(819, len((d_1100_v13_).cf1)), d_1101_v14_))), len(_dafny.SeqWithoutIsStrInference([d_1098_v11_ for d_1103_i0_ in range(default__.abs(185))]))])
                d_1102_v15_ = _dafny.MultiSet([(d_1098_v11_) * (d_1098_v11_), d_1098_v11_, d_1098_v11_])
                d_1104_v16_: _dafny.Array
                nw215_ = _dafny.Array(int(0), 10)
                d_1104_v16_ = nw215_
                d_1104_v16_ = (d_1104_v16_ if self.f16 else d_1104_v16_)
            source12_ = default__.fm33((self).f17, globalState)
            if source12_.is_DC4:
                d_1105___mcc_h0_ = source12_.cf7
                d_1106___mcc_h1_ = source12_.cf8
                d_1107___mcc_h2_ = source12_.cf9
                d_1108_cf9_ = d_1107___mcc_h2_
                d_1109_cf8_ = d_1106___mcc_h1_
                d_1110_cf7_ = d_1105___mcc_h0_
                d_1111_v17_: _dafny.Map
                d_1111_v17_ = _dafny.Map({d_1110_cf7_: -752})
                d_1111_v17_ = d_1111_v17_
                d_1112_v18_: _dafny.Seq
                d_1112_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                d_1109_cf8_ = len(((d_1112_v18_) + (d_1112_v18_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlh"))))
                d_1113_v19_: _dafny.MultiSet
                d_1113_v19_ = _dafny.MultiSet([d_1109_cf8_])
                d_1114_v20_: C1
                nw216_ = C1()
                nw216_.ctor__(d_1109_cf8_, d_1109_cf8_, (_dafny.MultiSet([(0) - (d_1109_cf8_)])).isdisjoint(d_1113_v19_), ((self).f17) + ((self).f17), default__.safeDivisionInt(d_1109_cf8_, d_1109_cf8_))
                d_1114_v20_ = nw216_
                d_1115_v21_: _dafny.Array
                def lambda57_(d_1116_v18_):
                    def lambda58_(d_1117_i1_):
                        return d_1116_v18_

                    return lambda58_

                init30_ = lambda57_(d_1112_v18_)
                nw217_ = _dafny.Array(None, 8)
                for i0_30_ in range(nw217_.length(0)):
                    nw217_[i0_30_] = init30_(i0_30_)
                d_1115_v21_ = nw217_
                index227_ = default__.safeIndex(902, (d_1115_v21_).length(0))
                (d_1115_v21_)[index227_] = (d_1112_v18_) + (d_1112_v18_)
                index228_ = default__.safeIndex(902, (d_1115_v21_).length(0))
                rhs256_ = d_1114_v20_
                rhs257_ = d_1112_v18_
                lhs201_ = d_1115_v21_
                lhs202_ = default__.safeIndex(902, (d_1115_v21_).length(0))
                d_1114_v20_ = rhs256_
                lhs201_[lhs202_] = rhs257_
                d_1118_v22_: D1
                d_1118_v22_ = D1_DC5(((d_1111_v17_)[d_1110_cf7_] if (d_1110_cf7_) in (d_1111_v17_) else len((d_1115_v21_)[default__.safeIndex(902, (d_1115_v21_).length(0))])), d_1110_cf7_)
                default__.m0(d_1109_cf8_, d_1112_v18_, default__.safeDivisionInt((d_1118_v22_).cf10, (d_1114_v20_).f30), d_1109_cf8_, globalState)
            elif source12_.is_DC5:
                d_1119___mcc_h3_ = source12_.cf10
                d_1120___mcc_h4_ = source12_.cf11
                d_1121_cf11_ = d_1120___mcc_h4_
                d_1122_cf10_ = d_1119___mcc_h3_
                d_1123_v23_: _dafny.Seq
                d_1123_v23_ = _dafny.SeqWithoutIsStrInference([d_1121_cf11_])
                d_1124_v24_: _dafny.Seq
                d_1124_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnjvqpd"))
                d_1125_v25_: C3
                nw218_ = C3()
                nw218_.ctor__(d_1124_v24_, self.f16, (self).f17)
                d_1125_v25_ = nw218_
                d_1126_v26_: _dafny.MultiSet
                d_1126_v26_ = _dafny.MultiSet([d_1122_cf10_, len(_dafny.Set({d_1125_v25_}))])
                d_1127_v27_: _dafny.Set
                d_1127_v27_ = _dafny.Set({default__.fm3(globalState)})
                d_1128_v28_: _dafny.Map
                d_1128_v28_ = _dafny.Map({d_1121_cf11_: d_1122_cf10_})
                d_1129_v29_: _dafny.Array
                nw219_ = _dafny.Array(None, 21)
                nw219_[int(0)] = d_1122_cf10_
                nw219_[int(1)] = (0) - (d_1122_cf10_)
                nw219_[int(2)] = len(d_1123_v23_)
                nw219_[int(3)] = len(d_1124_v24_)
                nw219_[int(4)] = 575
                nw219_[int(5)] = (_dafny.MultiSet([d_1122_cf10_, len((self).f17)])).cardinality
                nw219_[int(6)] = d_1122_cf10_
                nw219_[int(7)] = d_1122_cf10_
                nw219_[int(8)] = d_1122_cf10_
                nw219_[int(9)] = d_1122_cf10_
                nw219_[int(10)] = d_1122_cf10_
                nw219_[int(11)] = d_1122_cf10_
                nw219_[int(12)] = len(d_1123_v23_)
                nw219_[int(13)] = d_1122_cf10_
                nw219_[int(14)] = d_1122_cf10_
                nw219_[int(15)] = d_1122_cf10_
                nw219_[int(16)] = d_1122_cf10_
                nw219_[int(17)] = ((d_1126_v26_)[396] if (396) in (d_1126_v26_) else len(d_1127_v27_))
                nw219_[int(18)] = len(d_1124_v24_)
                nw219_[int(19)] = d_1122_cf10_
                nw219_[int(20)] = ((d_1128_v28_)[self.f16] if (self.f16) in (d_1128_v28_) else d_1122_cf10_)
                d_1129_v29_ = nw219_
                d_1130_v30_: _dafny.MultiSet
                d_1130_v30_ = _dafny.MultiSet([(d_1129_v29_ if d_1121_cf11_ else d_1129_v29_)])
                d_1131_v31_: _dafny.Seq
                d_1131_v31_ = _dafny.SeqWithoutIsStrInference([d_1130_v30_])
                d_1132_v32_: _dafny.Seq
                d_1132_v32_ = _dafny.SeqWithoutIsStrInference([d_1124_v24_])
                d_1130_v30_ = (d_1131_v31_)[default__.safeIndex(len((d_1132_v32_) + (default__.fm18(d_1122_cf10_, globalState))), len(d_1131_v31_))]
                d_1133_v33_: _dafny.Array
                nw220_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                d_1133_v33_ = nw220_
                index229_ = default__.safeIndex(643, (d_1133_v33_).length(0))
                (d_1133_v33_)[index229_] = d_1124_v24_
                index230_ = default__.safeIndex(643, (d_1133_v33_).length(0))
                rhs258_ = d_1124_v24_
                rhs259_ = (d_1124_v24_) + ((d_1125_v25_).f24)
                rhs260_ = d_1121_cf11_
                lhs203_ = d_1133_v33_
                lhs204_ = default__.safeIndex(643, (d_1133_v33_).length(0))
                lhs205_ = globalState
                d_1124_v24_ = rhs258_
                lhs203_[lhs204_] = rhs259_
                lhs205_.f5 = rhs260_
                d_1122_cf10_ = d_1122_cf10_
                (globalState).f2 = not(self.f16)
            elif source12_.is_DC6:
                d_1134_v34_: _dafny.Map
                d_1134_v34_ = _dafny.Map({self.f16: len(_dafny.Map({self.f16: self.f16}))})
                d_1135_v35_: _dafny.Seq
                d_1135_v35_ = _dafny.SeqWithoutIsStrInference([self.f16])
                d_1134_v34_ = (d_1134_v34_).set(not(self.f16), len(d_1135_v35_))
                d_1136_v36_: int
                d_1136_v36_ = 39
                d_1137_v37_: D8
                d_1137_v37_ = D8_DC21(d_1136_v36_, self.f16, self.f16, False, d_1136_v36_)
                d_1138_v38_: D4
                d_1138_v38_ = D4_DC12(self.f16, self.f16, (d_1137_v37_).cf39)
                d_1139_v39_: _dafny.Array
                def lambda59_(d_1140_i2_):
                    return self.f16

                init31_ = lambda59_
                nw221_ = _dafny.Array(None, 23)
                for i0_31_ in range(nw221_.length(0)):
                    nw221_[i0_31_] = init31_(i0_31_)
                d_1139_v39_ = nw221_
                d_1141_v40_: C0
                nw222_ = C0()
                nw222_.ctor__((d_1136_v36_) * ((d_1138_v38_).cf23), d_1139_v39_)
                d_1141_v40_ = nw222_
                d_1142_v41_: _dafny.Map
                d_1142_v41_ = _dafny.Map({(0) - ((d_1141_v40_).f26): self.f16})
                d_1143_v42_: str
                d_1143_v42_ = _dafny.CodePoint('x')
                (self).f16 = not(((self.f16) or (((d_1142_v41_)[len(_dafny.Map({self.f16: d_1143_v42_}))] if (len(_dafny.Map({self.f16: d_1143_v42_}))) in (d_1142_v41_) else self.f16))) == (self.f16))
                rhs261_ = not((d_1136_v36_) <= ((d_1141_v40_).f26))
                rhs262_ = (0) - (d_1136_v36_)
                lhs206_ = globalState
                lhs206_.f2 = rhs261_
                d_1136_v36_ = rhs262_
            elif True:
                d_1144___mcc_h5_ = source12_.cf6
                d_1145_cf6_ = d_1144___mcc_h5_
                d_1146_v43_: _dafny.Array
                nw223_ = _dafny.Array(False, 9)
                d_1146_v43_ = nw223_
                def lambda60_(d_1147_i3_):
                    return self.f16

                init32_ = lambda60_
                nw224_ = _dafny.Array(None, 15)
                for i0_32_ in range(nw224_.length(0)):
                    nw224_[i0_32_] = init32_(i0_32_)
                d_1146_v43_ = nw224_
                d_1148_v44_: int
                d_1148_v44_ = -578
                d_1149_v45_: _dafny.Seq
                d_1149_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbnnl"))
                default__.m0(d_1148_v44_, d_1149_v45_, d_1148_v44_, (d_1148_v44_ if False else d_1148_v44_), globalState)
                d_1150_v46_: C4
                nw225_ = C4()
                nw225_.ctor__(d_1148_v44_, self.f16)
                d_1150_v46_ = nw225_
                d_1151_v47_: _dafny.MultiSet
                d_1151_v47_ = _dafny.MultiSet([d_1150_v46_, d_1150_v46_, d_1150_v46_])
                d_1152_v49_: _dafny.Seq
                d_1152_v49_ = _dafny.SeqWithoutIsStrInference([default__.fm17(globalState)])
                d_1153_v50_: _dafny.Map
                d_1153_v50_ = _dafny.Map({(d_1145_cf6_).cardinality: d_1152_v49_})
                d_1154_v51_: _dafny.Map
                def iife125_():
                    coll55_ = _dafny.Map()
                    compr_55_: _dafny.Seq
                    for compr_55_ in (((d_1153_v50_)[(d_1150_v46_).f22] if ((d_1150_v46_).f22) in (d_1153_v50_) else d_1152_v49_)).Elements:
                        d_1155_v48_: _dafny.Seq = compr_55_
                        if (d_1155_v48_) in (((d_1153_v50_)[(d_1150_v46_).f22] if ((d_1150_v46_).f22) in (d_1153_v50_) else d_1152_v49_)):
                            coll55_[d_1155_v48_] = d_1150_v46_.f23
                    return _dafny.Map(coll55_)
                d_1154_v51_ = _dafny.Map({d_1151_v47_: (0) - (len(iife125_()
                ))})
                d_1156_v52_: _dafny.Seq
                d_1156_v52_ = _dafny.SeqWithoutIsStrInference([d_1150_v46_])
                d_1154_v51_ = (d_1154_v51_).set(_dafny.MultiSet((d_1156_v52_) + (d_1156_v52_)), default__.fm3(globalState))
                d_1157_v53_: D1
                d_1157_v53_ = D1_DC5((d_1150_v46_).f22, d_1150_v46_.f23)
                (self).f16 = ((d_1157_v53_).cf10) <= (((d_1150_v46_).f22) - ((d_1150_v46_).f22))
            d_1158_v54_: int
            d_1158_v54_ = 649
            d_1158_v54_ = default__.safeModuloInt((0) - (default__.fm3(globalState)), (0) - (d_1158_v54_))
            d_1159_v55_: C2
            nw226_ = C2()
            nw226_.ctor__(d_1158_v54_)
            d_1159_v55_ = nw226_
            (globalState).f2 = not(not((self.f16 if self.f16 else self.f16)))
        elif True:
            d_1160_v56_: int
            d_1160_v56_ = 119
            d_1160_v56_ = default__.safeDivisionInt(d_1160_v56_, d_1160_v56_)
            d_1161_v57_: _dafny.Set
            d_1161_v57_ = _dafny.Set({default__.fm0(self.f16, self.f16, globalState)})
            d_1162_v58_: _dafny.Seq
            d_1162_v58_ = _dafny.SeqWithoutIsStrInference([True, (d_1161_v57_).isdisjoint(d_1161_v57_), self.f16, (d_1160_v56_) > ((0) - (d_1160_v56_))])
            (self).f16 = (d_1162_v58_)[default__.safeIndex(d_1160_v56_, len(d_1162_v58_))]
            d_1163_v59_: _dafny.MultiSet
            d_1163_v59_ = _dafny.MultiSet([d_1160_v56_])
            d_1164_v60_: _dafny.Map
            d_1164_v60_ = _dafny.Map({d_1163_v59_: self.f16})
            d_1164_v60_ = (d_1164_v60_).set(((_dafny.MultiSet([d_1160_v56_])).set(d_1160_v56_, default__.abs(d_1160_v56_))).set(d_1160_v56_, default__.abs(len(((self).f17).set(default__.safeIndex((0) - ((0) - (d_1160_v56_)), len((self).f17)), d_1160_v56_)))), not(not(self.f16)))
            d_1165_v61_: _dafny.Seq
            d_1165_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbufmbxty"))
            source13_ = default__.fm34(d_1160_v56_, d_1160_v56_, d_1161_v57_, len(d_1165_v61_), globalState)
            if source13_.is_DC4:
                d_1166___mcc_h6_ = source13_.cf7
                d_1167___mcc_h7_ = source13_.cf8
                d_1168___mcc_h8_ = source13_.cf9
                d_1169_cf9_ = d_1168___mcc_h8_
                d_1170_cf8_ = d_1167___mcc_h7_
                d_1171_cf7_ = d_1166___mcc_h6_
                d_1172_v62_: str
                d_1172_v62_ = _dafny.CodePoint('s')
                d_1173_v63_: _dafny.Map
                d_1173_v63_ = _dafny.Map({default__.safeModuloInt((0) - (d_1160_v56_), d_1160_v56_): d_1172_v62_})
                d_1173_v63_ = (d_1173_v63_).set(d_1160_v56_, d_1172_v62_)
                d_1174_v64_: C2
                nw227_ = C2()
                nw227_.ctor__(default__.fm3(globalState))
                d_1174_v64_ = nw227_
                d_1165_v61_ = (d_1165_v61_) + (d_1165_v61_)
                d_1175_v65_: _dafny.Array
                nw228_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_1175_v65_ = nw228_
                d_1176_v66_: _dafny.Array
                def lambda61_(d_1177_v64_):
                    def lambda62_(d_1178_i4_):
                        return (d_1178_i4_) * ((d_1177_v64_).f25)

                    return lambda62_

                init33_ = lambda61_(d_1174_v64_)
                nw229_ = _dafny.Array(None, 16)
                for i0_33_ in range(nw229_.length(0)):
                    nw229_[i0_33_] = init33_(i0_33_)
                d_1176_v66_ = nw229_
                index231_ = default__.safeIndex(15, (d_1175_v65_).length(0))
                (d_1175_v65_)[index231_] = d_1176_v66_
                d_1179_v67_: _dafny.MultiSet
                d_1179_v67_ = _dafny.MultiSet([d_1171_cf7_, d_1171_cf7_])
                index232_ = default__.safeIndex(15, (d_1175_v65_).length(0))
                rhs263_ = (d_1160_v56_) + ((0) - (((0) - (d_1160_v56_) if self.f16 else (d_1174_v64_).f25)))
                rhs264_ = d_1176_v66_
                rhs265_ = (self).fm8(_dafny.SeqWithoutIsStrInference([d_1172_v62_ for d_1180_i5_ in range(default__.abs(-637))]), globalState)
                rhs266_ = _dafny.SeqWithoutIsStrInference([(self.f16) and (self.f16), d_1171_cf7_, ((D8_DC21(62, not(self.f16), self.f16, d_1171_cf7_, (d_1179_v67_).cardinality)).cf35) == (d_1170_cf8_), True])
                rhs267_ = (d_1179_v67_).issubset((d_1179_v67_) | (_dafny.MultiSet(d_1162_v58_)))
                lhs207_ = d_1175_v65_
                lhs208_ = default__.safeIndex(15, (d_1175_v65_).length(0))
                lhs209_ = self
                lhs210_ = globalState
                d_1170_cf8_ = rhs263_
                lhs207_[lhs208_] = rhs264_
                lhs209_.f16 = rhs265_
                d_1162_v58_ = rhs266_
                lhs210_.f2 = rhs267_
            elif source13_.is_DC5:
                d_1181___mcc_h9_ = source13_.cf10
                d_1182___mcc_h10_ = source13_.cf11
                d_1183_cf11_ = d_1182___mcc_h10_
                d_1184_cf10_ = d_1181___mcc_h9_
                d_1185_v68_: _dafny.Map
                d_1185_v68_ = _dafny.Map({d_1184_cf10_: (self).f17})
                d_1186_v69_: str
                d_1186_v69_ = _dafny.CodePoint('d')
                d_1187_v70_: _dafny.Map
                d_1187_v70_ = _dafny.Map({d_1186_v69_: 483})
                def iife126_():
                    coll56_ = _dafny.Set()
                    compr_56_: str
                    for compr_56_ in (d_1187_v70_).keys.Elements:
                        d_1188_v71_: str = compr_56_
                        if (d_1188_v71_) in (d_1187_v70_):
                            coll56_ = coll56_.union(_dafny.Set([d_1188_v71_]))
                    return _dafny.Set(coll56_)
                d_1185_v68_ = (d_1185_v68_).set(len(iife126_()
                ), (self).f17)
                d_1183_cf11_ = self.f16
                d_1189_v72_: _dafny.Array
                def lambda63_(d_1190_cf10_):
                    def lambda64_(d_1191_i6_):
                        return default__.safeModuloInt(d_1191_i6_, d_1190_cf10_)

                    return lambda64_

                init34_ = lambda63_(d_1184_cf10_)
                nw230_ = _dafny.Array(None, 11)
                for i0_34_ in range(nw230_.length(0)):
                    nw230_[i0_34_] = init34_(i0_34_)
                d_1189_v72_ = nw230_
                index233_ = default__.safeIndex(367, (d_1189_v72_).length(0))
                (d_1189_v72_)[index233_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfgkuwyr")))
                index234_ = default__.safeIndex(367, (d_1189_v72_).length(0))
                rhs268_ = ((d_1160_v56_) - (d_1160_v56_)) - (-640)
                rhs269_ = ((self).f17)[default__.safeIndex(d_1160_v56_, len((self).f17))]
                rhs270_ = d_1165_v61_
                rhs271_ = (d_1160_v56_ if self.f16 else d_1160_v56_)
                lhs211_ = d_1189_v72_
                lhs212_ = default__.safeIndex(367, (d_1189_v72_).length(0))
                lhs211_[lhs212_] = rhs268_
                d_1184_cf10_ = rhs269_
                d_1165_v61_ = rhs270_
                d_1184_cf10_ = rhs271_
                d_1160_v56_ = (len((d_1165_v61_) + (d_1165_v61_)) if self.f16 else (len(d_1165_v61_)) * (-682))
            elif source13_.is_DC6:
                d_1192_v73_: str
                d_1192_v73_ = _dafny.CodePoint('w')
                (globalState).f15 = d_1192_v73_
                d_1160_v56_ = d_1160_v56_
                d_1193_v74_: _dafny.Map
                d_1193_v74_ = _dafny.Map({self.f16: True})
                d_1194_v75_: _dafny.Map
                d_1194_v75_ = _dafny.Map({d_1160_v56_: False})
                d_1193_v74_ = (d_1193_v74_).set(((d_1194_v75_)[d_1160_v56_] if (d_1160_v56_) in (d_1194_v75_) else self.f16), default__.fm0(self.f16, self.f16, globalState))
                (globalState).f2 = default__.fm0(not(self.f16), self.f16, globalState)
            elif True:
                d_1195___mcc_h11_ = source13_.cf6
                d_1196_cf6_ = d_1195___mcc_h11_
                rhs272_ = d_1160_v56_
                rhs273_ = d_1165_v61_
                d_1160_v56_ = rhs272_
                d_1165_v61_ = rhs273_
                d_1197_v76_: _dafny.Set
                d_1197_v76_ = _dafny.Set({d_1160_v56_})
                d_1198_v77_: _dafny.Map
                d_1198_v77_ = _dafny.Map({(d_1196_cf6_).cardinality: len(d_1197_v76_)})
                d_1160_v56_ = (d_1160_v56_) + ((len(d_1198_v77_)) + (d_1160_v56_))
                d_1160_v56_ = (0) - (default__.safeModuloInt(d_1160_v56_, d_1160_v56_))
                d_1160_v56_ = (((d_1196_cf6_).intersection(d_1196_cf6_)) - (d_1196_cf6_)).cardinality
            d_1199_v78_: _dafny.Map
            d_1199_v78_ = _dafny.Map({(self.f16) or (self.f16): self.f16})
            d_1199_v78_ = (d_1199_v78_).set(self.f16, self.f16)
        d_1200_v79_: _dafny.Array
        def lambda65_(d_1201_i8_):
            return self.f16

        init35_ = lambda65_
        nw231_ = _dafny.Array(None, 14)
        for i0_35_ in range(nw231_.length(0)):
            nw231_[i0_35_] = init35_(i0_35_)
        d_1200_v79_ = nw231_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1200_v79_).length(0)):
            d_1202_i7_: int = guard_loop_8_
            if (True) and (((0) <= (d_1202_i7_)) and ((d_1202_i7_) < ((d_1200_v79_).length(0)))):
                (d_1200_v79_)[(d_1202_i7_)] = ((len(_dafny.Map({_dafny.CodePoint('r'): -381}))) - (836)) == ((319 if True else -786))
        d_1200_v79_ = d_1200_v79_
        d_1203_v80_: str
        d_1203_v80_ = _dafny.CodePoint('n')
        d_1204_v81_: _dafny.Seq
        d_1204_v81_ = _dafny.SeqWithoutIsStrInference([(d_1203_v80_ if self.f16 else d_1203_v80_), d_1203_v80_, d_1203_v80_])
        d_1205_v82_: int
        d_1205_v82_ = 516
        (globalState).f15 = (d_1204_v81_)[default__.safeIndex(default__.safeDivisionInt(d_1205_v82_, d_1205_v82_), len(d_1204_v81_))]
        d_1206_v83_: _dafny.Seq
        d_1206_v83_ = _dafny.SeqWithoutIsStrInference([self.f16])
        d_1207_v84_: D0
        d_1207_v84_ = D0_DC1(d_1204_v81_)
        d_1208_v85_: _dafny.Set
        d_1208_v85_ = _dafny.Set({(d_1206_v83_)[default__.safeIndex(len((d_1207_v84_).cf1), len(d_1206_v83_))], not(not(True)), not(self.f16)})
        if (d_1208_v85_).ispropersubset(d_1208_v85_):
            d_1209_v86_: _dafny.Array
            def lambda66_(d_1210_v82_):
                def lambda67_(d_1211_i9_):
                    return (d_1211_i9_) + (d_1210_v82_)

                return lambda67_

            init36_ = lambda66_(d_1205_v82_)
            nw232_ = _dafny.Array(None, 18)
            for i0_36_ in range(nw232_.length(0)):
                nw232_[i0_36_] = init36_(i0_36_)
            d_1209_v86_ = nw232_
            index235_ = default__.safeIndex(658, (d_1209_v86_).length(0))
            (d_1209_v86_)[index235_] = d_1205_v82_
            d_1212_v87_: _dafny.MultiSet
            d_1212_v87_ = _dafny.MultiSet([d_1203_v80_, (d_1204_v81_)[default__.safeIndex(d_1205_v82_, len(d_1204_v81_))], _dafny.CodePoint('j')])
            index236_ = default__.safeIndex(658, (d_1209_v86_).length(0))
            (d_1209_v86_)[index236_] = (d_1212_v87_).cardinality
            d_1213_v88_: _dafny.Set
            d_1213_v88_ = _dafny.Set({d_1205_v82_, (d_1209_v86_)[default__.safeIndex(658, (d_1209_v86_).length(0))]})
            d_1214_v90_: D9
            d_1214_v90_ = D9_DC27(self.f16, d_1213_v88_)
            d_1215_v91_: _dafny.Seq
            d_1215_v91_ = _dafny.SeqWithoutIsStrInference([(d_1214_v90_).cf48, _dafny.Set({d_1205_v82_})])
            d_1216_v92_: _dafny.MultiSet
            def iife127_():
                coll57_ = _dafny.Set()
                compr_57_: int
                for compr_57_ in ((self).f17).Elements:
                    d_1217_v89_: int = compr_57_
                    if (d_1217_v89_) in ((self).f17):
                        coll57_ = coll57_.union(_dafny.Set([default__.safeModuloInt(d_1217_v89_, 564)]))
                return _dafny.Set(coll57_)
            d_1216_v92_ = _dafny.MultiSet([d_1213_v88_, iife127_()
            , (d_1213_v88_) - (d_1213_v88_), (d_1215_v91_)[default__.safeIndex(d_1205_v82_, len(d_1215_v91_))], d_1213_v88_])
            d_1216_v92_ = d_1216_v92_
            d_1218_v93_: _dafny.Array
            nw233_ = _dafny.Array(None, 10)
            nw233_[int(0)] = d_1203_v80_
            nw233_[int(1)] = d_1203_v80_
            nw233_[int(2)] = d_1203_v80_
            nw233_[int(3)] = d_1203_v80_
            nw233_[int(4)] = d_1203_v80_
            nw233_[int(5)] = _dafny.CodePoint('k')
            nw233_[int(6)] = d_1203_v80_
            nw233_[int(7)] = d_1203_v80_
            nw233_[int(8)] = d_1203_v80_
            nw233_[int(9)] = d_1203_v80_
            d_1218_v93_ = nw233_
            d_1218_v93_ = d_1218_v93_
            index237_ = default__.safeIndex(658, (d_1209_v86_).length(0))
            rhs274_ = self.f16
            rhs275_ = (d_1209_v86_)[default__.safeIndex(658, (d_1209_v86_).length(0))]
            rhs276_ = (default__.fm3(globalState)) <= ((d_1209_v86_)[default__.safeIndex(658, (d_1209_v86_).length(0))])
            lhs213_ = globalState
            lhs214_ = d_1209_v86_
            lhs215_ = default__.safeIndex(658, (d_1209_v86_).length(0))
            lhs216_ = globalState
            lhs213_.f5 = rhs274_
            lhs214_[lhs215_] = rhs275_
            lhs216_.f2 = rhs276_
            index238_ = default__.safeIndex(658, (d_1209_v86_).length(0))
            (d_1209_v86_)[index238_] = (d_1209_v86_)[default__.safeIndex(658, (d_1209_v86_).length(0))]
        elif True:
            d_1219_v94_: D0
            d_1219_v94_ = D0_DC0((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlth"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilunfi"))))
            d_1219_v94_ = d_1219_v94_
            d_1220_v95_: _dafny.Array
            nw234_ = _dafny.Array(int(0), 29)
            d_1220_v95_ = nw234_
            index239_ = default__.safeIndex(954, (d_1200_v79_).length(0))
            (d_1200_v79_)[index239_] = (self).fm8(d_1204_v81_, globalState)
            index240_ = default__.safeIndex(954, (d_1200_v79_).length(0))
            rhs277_ = d_1203_v80_
            rhs278_ = d_1220_v95_
            rhs279_ = self.f16
            lhs217_ = globalState
            lhs218_ = d_1200_v79_
            lhs219_ = default__.safeIndex(954, (d_1200_v79_).length(0))
            lhs217_.f15 = rhs277_
            d_1220_v95_ = rhs278_
            lhs218_[lhs219_] = rhs279_
            if (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))]:
                d_1221_v96_: _dafny.Array
                nw235_ = _dafny.Array(False, 27)
                d_1221_v96_ = nw235_
                d_1222_v97_: D0
                d_1222_v97_ = D0_DC2(d_1205_v82_, d_1205_v82_, len(_dafny.Set({d_1205_v82_})), d_1221_v96_)
                d_1223_v98_: C0
                nw236_ = C0()
                nw236_.ctor__(d_1205_v82_, (d_1222_v97_).cf5)
                d_1223_v98_ = nw236_
                d_1224_v99_: _dafny.Seq
                d_1224_v99_ = _dafny.SeqWithoutIsStrInference([d_1221_v96_, d_1221_v96_])
                d_1225_v100_: _dafny.Map
                d_1225_v100_ = _dafny.Map({(d_1206_v83_)[default__.safeIndex((0) - ((d_1223_v98_).f26), len(d_1206_v83_))]: (d_1223_v98_).f26})
                d_1226_v101_: _dafny.Map
                d_1226_v101_ = _dafny.Map({self.f16: d_1221_v96_})
                d_1227_v102_: _dafny.Array
                nw237_ = _dafny.Array(None, 18)
                nw237_[int(0)] = ((d_1224_v99_)[default__.safeIndex(len(d_1225_v100_), len(d_1224_v99_))] if (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))] else d_1221_v96_)
                nw237_[int(1)] = d_1223_v98_.f27
                nw237_[int(2)] = d_1223_v98_.f27
                nw237_[int(3)] = ((d_1226_v101_)[False] if (False) in (d_1226_v101_) else d_1223_v98_.f27)
                nw237_[int(4)] = d_1221_v96_
                nw237_[int(5)] = d_1223_v98_.f27
                nw237_[int(6)] = d_1223_v98_.f27
                nw237_[int(7)] = d_1221_v96_
                nw237_[int(8)] = d_1221_v96_
                nw237_[int(9)] = d_1221_v96_
                nw237_[int(10)] = d_1223_v98_.f27
                nw237_[int(11)] = d_1221_v96_
                nw237_[int(12)] = d_1223_v98_.f27
                nw237_[int(13)] = d_1223_v98_.f27
                nw237_[int(14)] = d_1221_v96_
                nw237_[int(15)] = d_1223_v98_.f27
                nw237_[int(16)] = d_1223_v98_.f27
                nw237_[int(17)] = d_1223_v98_.f27
                d_1227_v102_ = nw237_
                index241_ = default__.safeIndex(166, (d_1227_v102_).length(0))
                (d_1227_v102_)[index241_] = d_1221_v96_
                index242_ = default__.safeIndex(166, (d_1227_v102_).length(0))
                rhs280_ = (len(_dafny.SeqWithoutIsStrInference([self.f16]))) + ((0) - (d_1205_v82_))
                rhs281_ = d_1203_v80_
                rhs282_ = d_1223_v98_.f27
                lhs220_ = globalState
                lhs221_ = d_1227_v102_
                lhs222_ = default__.safeIndex(166, (d_1227_v102_).length(0))
                d_1205_v82_ = rhs280_
                lhs220_.f15 = rhs281_
                lhs221_[lhs222_] = rhs282_
                d_1228_v104_: _dafny.Map
                d_1228_v104_ = _dafny.Map({d_1205_v82_: (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))]})
                d_1229_v106_: _dafny.Map
                def iife128_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in (d_1228_v104_).keys.Elements:
                        d_1230_v103_: int = compr_58_
                        if (d_1230_v103_) in (d_1228_v104_):
                            coll58_[(d_1230_v103_) * (586)] = (D6_DC16(self.f16, (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))])).cf29
                    return _dafny.Map(coll58_)
                def iife129_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(709, 931):
                        d_1231_v105_: int = compr_59_
                        if ((709) <= (d_1231_v105_)) and ((d_1231_v105_) < (931)):
                            coll59_[default__.safeDivisionInt(d_1231_v105_, (d_1223_v98_).f26)] = (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))]
                    return _dafny.Map(coll59_)
                d_1229_v106_ = _dafny.Map({(iife128_()
                ) | (iife129_()
                ): default__.safeDivisionInt(748, (d_1223_v98_).f26)})
                d_1229_v106_ = (d_1229_v106_).set((d_1228_v104_).set((d_1223_v98_).f26, False), (d_1223_v98_).f26)
                d_1232_v107_: _dafny.Set
                d_1232_v107_ = _dafny.Set({d_1208_v85_, d_1208_v85_, d_1208_v85_, d_1208_v85_, d_1208_v85_})
                d_1233_v108_: _dafny.Array
                def lambda68_(d_1234_v80_):
                    def lambda69_(d_1235_i10_):
                        return d_1234_v80_

                    return lambda69_

                init37_ = lambda68_(d_1203_v80_)
                nw238_ = _dafny.Array(None, 17)
                for i0_37_ in range(nw238_.length(0)):
                    nw238_[i0_37_] = init37_(i0_37_)
                d_1233_v108_ = nw238_
                d_1236_v109_: _dafny.Array
                d_1236_v109_ = d_1233_v108_
                index243_ = default__.safeIndex(954, (d_1200_v79_).length(0))
                rhs283_ = default__.fm0(self.f16, self.f16, globalState)
                rhs284_ = (d_1232_v107_) != (d_1232_v107_)
                rhs285_ = ((len(_dafny.SeqWithoutIsStrInference([d_1233_v108_, d_1233_v108_, d_1233_v108_, (d_1236_v109_), d_1233_v108_])) if not((d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))]) else d_1205_v82_)) + (d_1205_v82_)
                rhs286_ = self.f16
                rhs287_ = ((d_1223_v98_).f26) != (d_1205_v82_)
                lhs223_ = d_1200_v79_
                lhs224_ = default__.safeIndex(954, (d_1200_v79_).length(0))
                lhs225_ = globalState
                lhs226_ = self
                lhs227_ = globalState
                lhs223_[lhs224_] = rhs283_
                lhs225_.f5 = rhs284_
                d_1205_v82_ = rhs285_
                lhs226_.f16 = rhs286_
                lhs227_.f2 = rhs287_
                index244_ = default__.safeIndex(75, (d_1220_v95_).length(0))
                (d_1220_v95_)[index244_] = 724
                d_1237_v110_: _dafny.Array
                nw239_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_1237_v110_ = nw239_
                index245_ = default__.safeIndex(835, (d_1237_v110_).length(0))
                (d_1237_v110_)[index245_] = (_dafny.SeqWithoutIsStrInference([d_1203_v80_ for d_1238_i11_ in range(default__.abs(775))])) + (d_1204_v81_)
                d_1239_v111_: _dafny.Set
                d_1239_v111_ = _dafny.Set({default__.fm3(globalState), (d_1223_v98_).f26})
                d_1240_v112_: D8
                d_1240_v112_ = D8_DC21(d_1205_v82_, self.f16, (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], (d_1223_v98_).f26)
                index246_ = default__.safeIndex(542, (d_1220_v95_).length(0))
                (d_1220_v95_)[index246_] = (len(d_1239_v111_) if (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))] else (d_1240_v112_).cf35)
                index247_ = default__.safeIndex(75, (d_1220_v95_).length(0))
                index248_ = default__.safeIndex(835, (d_1237_v110_).length(0))
                index249_ = default__.safeIndex(542, (d_1220_v95_).length(0))
                rhs288_ = (d_1223_v98_).f26
                rhs289_ = d_1204_v81_
                rhs290_ = d_1225_v100_
                rhs291_ = (d_1205_v82_) * ((d_1223_v98_).f26)
                lhs228_ = d_1220_v95_
                lhs229_ = default__.safeIndex(75, (d_1220_v95_).length(0))
                lhs230_ = d_1237_v110_
                lhs231_ = default__.safeIndex(835, (d_1237_v110_).length(0))
                lhs232_ = d_1220_v95_
                lhs233_ = default__.safeIndex(542, (d_1220_v95_).length(0))
                lhs228_[lhs229_] = rhs288_
                lhs230_[lhs231_] = rhs289_
                d_1225_v100_ = rhs290_
                lhs232_[lhs233_] = rhs291_
            elif True:
                d_1241_v113_: _dafny.Array
                nw240_ = _dafny.Array(None, 20)
                nw240_[int(0)] = d_1204_v81_
                nw240_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1203_v80_ for d_1242_i12_ in range(default__.abs(632))])
                nw240_[int(2)] = d_1204_v81_
                nw240_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgjcs"))
                nw240_[int(4)] = d_1204_v81_
                nw240_[int(5)] = d_1204_v81_
                nw240_[int(6)] = (d_1204_v81_) + (d_1204_v81_)
                nw240_[int(7)] = d_1204_v81_
                nw240_[int(8)] = d_1204_v81_
                nw240_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyruluxc"))
                nw240_[int(10)] = d_1204_v81_
                nw240_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1203_v80_ for d_1243_i13_ in range(default__.abs(317))])
                nw240_[int(12)] = (d_1204_v81_) + (d_1204_v81_)
                nw240_[int(13)] = d_1204_v81_
                nw240_[int(14)] = _dafny.SeqWithoutIsStrInference([d_1203_v80_ for d_1244_i14_ in range(default__.abs(229))])
                nw240_[int(15)] = d_1204_v81_
                nw240_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfvxlqw"))
                nw240_[int(17)] = d_1204_v81_
                nw240_[int(18)] = d_1204_v81_
                nw240_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evndkdyhw"))
                d_1241_v113_ = nw240_
                index250_ = default__.safeIndex(823, (d_1241_v113_).length(0))
                (d_1241_v113_)[index250_] = (d_1204_v81_) + (d_1204_v81_)
                index251_ = default__.safeIndex(823, (d_1241_v113_).length(0))
                rhs292_ = default__.fm3(globalState)
                rhs293_ = d_1206_v83_
                rhs294_ = (d_1204_v81_) + (_dafny.SeqWithoutIsStrInference([d_1203_v80_ for d_1245_i15_ in range(default__.abs(205))]))
                lhs234_ = d_1241_v113_
                lhs235_ = default__.safeIndex(823, (d_1241_v113_).length(0))
                d_1205_v82_ = rhs292_
                r0 = rhs293_
                lhs234_[lhs235_] = rhs294_
                d_1246_v114_: _dafny.Map
                d_1246_v114_ = _dafny.Map({(944) + (d_1205_v82_): self.f16})
                d_1247_v115_: _dafny.Array
                nw241_ = _dafny.Array(False, 1)
                d_1247_v115_ = nw241_
                d_1248_v116_: C0
                nw242_ = C0()
                nw242_.ctor__(len((d_1241_v113_)[default__.safeIndex(823, (d_1241_v113_).length(0))]), d_1247_v115_)
                d_1248_v116_ = nw242_
                d_1249_v117_: D9
                d_1249_v117_ = D9_DC26(d_1205_v82_, d_1248_v116_)
                d_1246_v114_ = (d_1246_v114_).set((d_1249_v117_).cf45, (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))])
                d_1250_v118_: _dafny.Array
                nw243_ = _dafny.Array(_dafny.Seq({}), 14)
                d_1250_v118_ = nw243_
                d_1251_v119_: _dafny.Map
                d_1251_v119_ = _dafny.Map({((d_1248_v116_).f26) + ((d_1248_v116_).f26): d_1250_v118_})
                d_1251_v119_ = (d_1251_v119_).set(d_1205_v82_, d_1250_v118_)
                d_1205_v82_ = (d_1205_v82_) + ((d_1248_v116_).f26)
                d_1252_v120_: D8
                d_1252_v120_ = D8_DC21((d_1248_v116_).f26, (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], self.f16, (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], (d_1248_v116_).f26)
                d_1253_v121_: _dafny.MultiSet
                d_1253_v121_ = _dafny.MultiSet([(d_1252_v120_).cf35, 141])
                d_1254_v122_: _dafny.MultiSet
                d_1254_v122_ = _dafny.MultiSet([d_1203_v80_, d_1203_v80_, d_1203_v80_])
                d_1205_v82_ = ((d_1253_v121_).cardinality) * ((d_1254_v122_).cardinality)
            if (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))]:
                d_1255_v123_: C3
                nw244_ = C3()
                nw244_.ctor__(d_1204_v81_, self.f16, (_dafny.SeqWithoutIsStrInference([d_1205_v82_])) + (_dafny.SeqWithoutIsStrInference([len(d_1206_v83_) for d_1256_i16_ in range(default__.abs(655))])))
                d_1255_v123_ = nw244_
                d_1257_v124_: _dafny.Map
                d_1257_v124_ = _dafny.Map({d_1205_v82_: d_1205_v82_})
                d_1258_v125_: _dafny.Set
                d_1258_v125_ = _dafny.Set({d_1257_v124_})
                d_1259_v126_: _dafny.Array
                nw245_ = _dafny.Array(False, 23)
                d_1259_v126_ = nw245_
                d_1260_v127_: _dafny.Map
                d_1260_v127_ = _dafny.Map({(d_1258_v125_).ispropersubset(default__.fm35(d_1205_v82_, _dafny.CodePoint('o'), globalState)): d_1259_v126_})
                d_1260_v127_ = (d_1260_v127_).set((d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], d_1259_v126_)
                (self).f16 = (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))]
                d_1205_v82_ = d_1205_v82_
                d_1205_v82_ = d_1205_v82_
            elif True:
                d_1261_v128_: _dafny.Map
                d_1261_v128_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikbah")): default__.fm0(self.f16, (d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], globalState)})
                d_1262_v129_: _dafny.Map
                d_1262_v129_ = _dafny.Map({(self.f16) or (self.f16): (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhohi")): self.f16})) | (d_1261_v128_)})
                d_1262_v129_ = (d_1262_v129_).set((d_1200_v79_)[default__.safeIndex(954, (d_1200_v79_).length(0))], d_1261_v128_)
                d_1263_v130_: _dafny.MultiSet
                d_1263_v130_ = _dafny.MultiSet([d_1204_v81_, d_1204_v81_, d_1204_v81_])
                d_1264_v131_: _dafny.Seq
                d_1264_v131_ = _dafny.SeqWithoutIsStrInference([d_1204_v81_, d_1204_v81_])
                rhs295_ = self.f16
                rhs296_ = d_1205_v82_
                rhs297_ = (_dafny.MultiSet(d_1264_v131_)).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljveyjw")), default__.abs(default__.safeModuloInt(d_1205_v82_, d_1205_v82_)))
                rhs298_ = d_1205_v82_
                lhs236_ = globalState
                lhs236_.f5 = rhs295_
                d_1205_v82_ = rhs296_
                d_1263_v130_ = rhs297_
                d_1205_v82_ = rhs298_
                (self).f16 = False
                d_1205_v82_ = (d_1205_v82_) * ((0) - (d_1205_v82_))
                d_1265_v132_: _dafny.MultiSet
                d_1265_v132_ = _dafny.MultiSet([d_1205_v82_])
                (globalState).f2 = (d_1205_v82_) in (d_1265_v132_)
            rhs299_ = d_1205_v82_
            rhs300_ = ((0) - (d_1205_v82_)) != (d_1205_v82_)
            lhs237_ = self
            d_1205_v82_ = rhs299_
            lhs237_.f16 = rhs300_
        d_1266_v133_: D8
        d_1266_v133_ = D8_DC24(D8_DC23())
        d_1266_v133_ = d_1266_v133_
        r0 = _dafny.SeqWithoutIsStrInference([self.f16])
        return r0


class C6:
    def  __init__(self):
        self._f21: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f21):
        (self)._f21 = f21

    def fm6(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lc"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(not(True)), False, not(False), False})) for d_1267_i0_ in range(default__.abs(655))])))])

    def m3(self, p0, globalState):
        d_1268_v0_: _dafny.Seq
        d_1268_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpaxlupea"))
        d_1269_v1_: int
        d_1269_v1_ = 76
        hi8_ = (d_1269_v1_) * (d_1269_v1_)
        for d_1270_i0_ in range(len(d_1268_v0_), hi8_):
            d_1271_v2_: _dafny.Array
            nw246_ = _dafny.Array(int(0), 3)
            d_1271_v2_ = nw246_
            index252_ = default__.safeIndex(553, (d_1271_v2_).length(0))
            (d_1271_v2_)[index252_] = (d_1270_i0_) + (d_1270_i0_)
            index253_ = default__.safeIndex(553, (d_1271_v2_).length(0))
            (d_1271_v2_)[index253_] = d_1269_v1_
            d_1272_v3_: _dafny.Set
            d_1272_v3_ = _dafny.Set({p0})
            d_1273_v4_: D2
            d_1273_v4_ = D2_DC7(d_1272_v3_)
            d_1274_v5_: _dafny.MultiSet
            d_1274_v5_ = _dafny.MultiSet([len(((d_1273_v4_).cf12) | (d_1272_v3_)), 441])
            d_1275_v6_: str
            d_1275_v6_ = _dafny.CodePoint('g')
            d_1276_v7_: _dafny.Set
            d_1276_v7_ = _dafny.Set({d_1275_v6_, d_1275_v6_, default__.fm2(globalState)})
            d_1269_v1_ = ((d_1274_v5_)[len(d_1276_v7_)] if (len(d_1276_v7_)) in (d_1274_v5_) else (d_1269_v1_) * (d_1270_i0_))
            index254_ = default__.safeIndex(553, (d_1271_v2_).length(0))
            (d_1271_v2_)[index254_] = (0) - (d_1270_i0_)
            d_1277_v8_: D1
            d_1277_v8_ = D1_DC6()
            d_1277_v8_ = d_1277_v8_
        hi9_ = 169
        for d_1278_i1_ in range(d_1269_v1_, hi9_):
            d_1279_v9_: _dafny.MultiSet
            d_1279_v9_ = _dafny.MultiSet([p0])
            d_1280_v10_: D1
            d_1280_v10_ = D1_DC3(d_1279_v9_)
            pat_let_tv27_ = d_1279_v9_
            def iife130_(_pat_let35_0):
                def iife131_(d_1281_dt__update__tmp_h0_):
                    def iife132_(_pat_let36_0):
                        def iife133_(d_1282_dt__update_hcf6_h0_):
                            return D1_DC3(d_1282_dt__update_hcf6_h0_)
                        return iife133_(_pat_let36_0)
                    return iife132_(pat_let_tv27_)
                return iife131_(_pat_let35_0)
            source14_ = iife130_(d_1280_v10_)
            if source14_.is_DC4:
                d_1283___mcc_h0_ = source14_.cf7
                d_1284___mcc_h1_ = source14_.cf8
                d_1285___mcc_h2_ = source14_.cf9
                d_1286_cf9_ = d_1285___mcc_h2_
                d_1287_cf8_ = d_1284___mcc_h1_
                d_1288_cf7_ = d_1283___mcc_h0_
                d_1289_v11_: _dafny.Set
                d_1289_v11_ = _dafny.Set({d_1269_v1_})
                d_1290_v12_: _dafny.Set
                d_1290_v12_ = _dafny.Set({d_1289_v11_, d_1289_v11_})
                d_1291_v13_: _dafny.Map
                d_1291_v13_ = _dafny.Map({175: d_1278_i1_})
                def iife134_():
                    coll60_ = _dafny.Set()
                    compr_60_: int
                    for compr_60_ in (d_1291_v13_).keys.Elements:
                        d_1292_v14_: int = compr_60_
                        if (d_1292_v14_) in (d_1291_v13_):
                            coll60_ = coll60_.union(_dafny.Set([default__.safeDivisionInt(d_1292_v14_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgmhqd"))))]))
                    return _dafny.Set(coll60_)
                (globalState).f2 = ((_dafny.Set({d_1289_v11_})).intersection(_dafny.Set({iife134_()
                }))).ispropersubset(d_1290_v12_)
                d_1293_v15_: bool
                d_1294_v16_: bool
                d_1295_v17_: bool
                d_1296_v18_: int
                out10_: bool
                out11_: bool
                out12_: bool
                out13_: int
                out10_, out11_, out12_, out13_ = (self).m4(globalState)
                d_1293_v15_ = out10_
                d_1294_v16_ = out11_
                d_1295_v17_ = out12_
                d_1296_v18_ = out13_
                d_1297_v19_: _dafny.Seq
                d_1297_v19_ = _dafny.SeqWithoutIsStrInference([d_1268_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sluvhgu"))])
                d_1298_v20_: str
                d_1298_v20_ = _dafny.CodePoint('w')
                d_1299_v21_: _dafny.MultiSet
                d_1299_v21_ = _dafny.MultiSet([((d_1297_v19_)[default__.safeIndex(d_1287_cf8_, len(d_1297_v19_))]).set(default__.safeIndex(len(d_1268_v0_), len((d_1297_v19_)[default__.safeIndex(d_1287_cf8_, len(d_1297_v19_))])), d_1298_v20_), d_1268_v0_, d_1268_v0_])
                d_1299_v21_ = d_1299_v21_
                d_1300_v22_: _dafny.Seq
                d_1300_v22_ = _dafny.SeqWithoutIsStrInference([-967])
                d_1301_v23_: T0
                nw247_ = C1()
                nw247_.ctor__(919, len(d_1268_v0_), d_1288_cf7_, d_1300_v22_, d_1287_cf8_)
                d_1301_v23_ = nw247_
                d_1302_v24_: _dafny.Map
                d_1302_v24_ = _dafny.Map({d_1294_v16_: d_1301_v23_})
                d_1303_v25_: D8
                d_1303_v25_ = D8_DC22(d_1293_v15_, d_1301_v23_, d_1295_v17_)
                d_1304_v26_: _dafny.Seq
                d_1304_v26_ = _dafny.SeqWithoutIsStrInference([d_1301_v23_, (d_1303_v25_).cf41])
                d_1305_v27_: _dafny.Map
                d_1305_v27_ = _dafny.Map({d_1293_v15_: len(d_1300_v22_)})
                d_1306_v28_: _dafny.Map
                d_1306_v28_ = _dafny.Map({((d_1302_v24_)[d_1288_cf7_] if (d_1288_cf7_) in (d_1302_v24_) else (d_1304_v26_)[default__.safeIndex(len(d_1305_v27_), len(d_1304_v26_))]): d_1301_v23_.f16})
                d_1306_v28_ = (d_1306_v28_).set(d_1301_v23_, d_1294_v16_)
            elif source14_.is_DC5:
                d_1307___mcc_h3_ = source14_.cf10
                d_1308___mcc_h4_ = source14_.cf11
                d_1309_cf11_ = d_1308___mcc_h4_
                d_1310_cf10_ = d_1307___mcc_h3_
                d_1311_v29_: _dafny.Array
                nw248_ = _dafny.Array(int(0), 20)
                d_1311_v29_ = nw248_
                index255_ = default__.safeIndex(733, (d_1311_v29_).length(0))
                (d_1311_v29_)[index255_] = d_1278_i1_
                index256_ = default__.safeIndex(733, (d_1311_v29_).length(0))
                (d_1311_v29_)[index256_] = default__.safeModuloInt(d_1310_cf10_, d_1310_cf10_)
                (globalState).f2 = (p0 if p0 else not(p0))
                d_1312_v30_: _dafny.Seq
                d_1312_v30_ = _dafny.SeqWithoutIsStrInference([d_1309_cf11_])
                d_1313_v31_: _dafny.Set
                d_1313_v31_ = _dafny.Set({p0, False})
                d_1312_v30_ = (d_1312_v30_) + ((d_1312_v30_ if d_1309_cf11_ else (d_1312_v30_).set(default__.safeIndex(len(d_1313_v31_), len(d_1312_v30_)), d_1309_cf11_)))
                d_1314_v32_: _dafny.Set
                d_1314_v32_ = _dafny.Set({d_1269_v1_})
                (globalState).f2 = default__.fm0((d_1314_v32_).ispropersubset(d_1314_v32_), default__.fm0(default__.fm0(p0, p0, globalState), p0, globalState), globalState)
            elif source14_.is_DC6:
                d_1315_v33_: str
                d_1315_v33_ = _dafny.CodePoint('a')
                pat_let_tv28_ = d_1315_v33_
                def iife135_(_pat_let37_0):
                    def iife136_(d_1316_dt__update__tmp_h1_):
                        def iife137_(_pat_let38_0):
                            def iife138_(d_1317_dt__update_hcf50_h0_):
                                return D11_DC29(d_1317_dt__update_hcf50_h0_)
                            return iife138_(_pat_let38_0)
                        return iife137_(pat_let_tv28_)
                    return iife136_(_pat_let37_0)
                (globalState).f5 = ((d_1268_v0_).set(default__.safeIndex(d_1269_v1_, len(d_1268_v0_)), (iife135_(D11_DC29(d_1315_v33_))).cf50)) <= (d_1268_v0_)
                d_1269_v1_ = d_1278_i1_
                d_1318_v34_: _dafny.Set
                d_1318_v34_ = _dafny.Set({d_1315_v33_, _dafny.CodePoint('u')})
                d_1318_v34_ = d_1318_v34_
                d_1319_v35_: _dafny.Array
                nw249_ = _dafny.Array(False, 18)
                d_1319_v35_ = nw249_
                d_1320_v36_: _dafny.Seq
                d_1320_v36_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1321_v37_: _dafny.Array
                nw250_ = _dafny.Array(None, 2)
                nw250_[int(0)] = len(d_1320_v36_)
                nw250_[int(1)] = (0) - (d_1278_i1_)
                d_1321_v37_ = nw250_
                d_1322_v38_: _dafny.Map
                d_1322_v38_ = _dafny.Map({(d_1319_v35_ if p0 else d_1319_v35_): d_1321_v37_})
                d_1322_v38_ = (d_1322_v38_).set(d_1319_v35_, d_1321_v37_)
            elif True:
                d_1323___mcc_h5_ = source14_.cf6
                d_1324_cf6_ = d_1323___mcc_h5_
                d_1325_v39_: _dafny.Map
                d_1325_v39_ = _dafny.Map({p0: d_1278_i1_})
                d_1326_v40_: str
                d_1326_v40_ = _dafny.CodePoint('t')
                d_1327_v41_: _dafny.Map
                d_1327_v41_ = _dafny.Map({d_1269_v1_: p0})
                d_1328_v42_: _dafny.Seq
                d_1328_v42_ = _dafny.SeqWithoutIsStrInference([len(d_1327_v41_)])
                d_1329_v43_: _dafny.Map
                d_1329_v43_ = _dafny.Map({d_1326_v40_: d_1328_v42_})
                d_1330_v44_: _dafny.Seq
                d_1330_v44_ = _dafny.SeqWithoutIsStrInference([d_1328_v42_])
                d_1331_v45_: _dafny.Array
                nw251_ = _dafny.Array(None, 18)
                nw251_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1269_v1_])
                nw251_[int(1)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1332_i2_ in range(default__.abs(888))]))])
                nw251_[int(2)] = _dafny.SeqWithoutIsStrInference([d_1269_v1_])
                nw251_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1269_v1_ for d_1333_i3_ in range(default__.abs(925))])
                nw251_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1278_i1_])
                nw251_[int(5)] = ((d_1329_v43_)[d_1326_v40_] if (d_1326_v40_) in (d_1329_v43_) else d_1328_v42_)
                nw251_[int(6)] = d_1328_v42_
                nw251_[int(7)] = _dafny.SeqWithoutIsStrInference([668])
                nw251_[int(8)] = d_1328_v42_
                nw251_[int(9)] = d_1328_v42_
                nw251_[int(10)] = d_1328_v42_
                nw251_[int(11)] = d_1328_v42_
                nw251_[int(12)] = d_1328_v42_
                nw251_[int(13)] = (d_1330_v44_)[default__.safeIndex(d_1269_v1_, len(d_1330_v44_))]
                nw251_[int(14)] = d_1328_v42_
                nw251_[int(15)] = d_1328_v42_
                nw251_[int(16)] = d_1328_v42_
                nw251_[int(17)] = (D3_DC9(d_1328_v42_)).cf18
                d_1331_v45_ = nw251_
                d_1334_v46_: D2
                d_1334_v46_ = D2_DC8(d_1325_v39_, d_1331_v45_, p0, False, d_1278_i1_)
                d_1335_v47_: _dafny.Map
                d_1335_v47_ = _dafny.Map({(d_1334_v46_).cf15: d_1268_v0_})
                d_1269_v1_ = (d_1278_i1_) - (len(d_1335_v47_))
                d_1336_v48_: _dafny.Array
                def lambda70_(d_1337_p0_):
                    def lambda71_(d_1338_i4_):
                        return _dafny.SeqWithoutIsStrInference([d_1337_p0_, d_1337_p0_])

                    return lambda71_

                init38_ = lambda70_(p0)
                nw252_ = _dafny.Array(None, 21)
                for i0_38_ in range(nw252_.length(0)):
                    nw252_[i0_38_] = init38_(i0_38_)
                d_1336_v48_ = nw252_
                d_1339_v49_: _dafny.Seq
                d_1339_v49_ = _dafny.SeqWithoutIsStrInference([False, p0])
                index257_ = default__.safeIndex(491, (d_1336_v48_).length(0))
                (d_1336_v48_)[index257_] = d_1339_v49_
                index258_ = default__.safeIndex(491, (d_1336_v48_).length(0))
                (d_1336_v48_)[index258_] = d_1339_v49_
                d_1340_v50_: _dafny.Array
                nw253_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_1340_v50_ = nw253_
                d_1341_v51_: _dafny.Array
                nw254_ = _dafny.Array(False, 29)
                d_1341_v51_ = nw254_
                index259_ = default__.safeIndex(937, (d_1340_v50_).length(0))
                (d_1340_v50_)[index259_] = d_1341_v51_
                index260_ = default__.safeIndex(937, (d_1340_v50_).length(0))
                (d_1340_v50_)[index260_] = d_1341_v51_
                index261_ = default__.safeIndex(491, (d_1336_v48_).length(0))
                (d_1336_v48_)[index261_] = d_1339_v49_
            d_1342_v52_: _dafny.MultiSet
            d_1342_v52_ = _dafny.MultiSet([d_1269_v1_, (d_1278_i1_ if p0 else d_1269_v1_), (-937) + (d_1278_i1_), d_1269_v1_])
            d_1343_v53_: _dafny.Seq
            d_1343_v53_ = _dafny.SeqWithoutIsStrInference([-710, d_1269_v1_])
            d_1344_v54_: _dafny.Map
            d_1344_v54_ = _dafny.Map({(_dafny.MultiSet(d_1343_v53_)).cardinality: p0})
            d_1345_v55_: _dafny.Map
            d_1345_v55_ = _dafny.Map({d_1344_v54_: d_1278_i1_})
            d_1346_v56_: _dafny.Seq
            d_1346_v56_ = _dafny.SeqWithoutIsStrInference([d_1268_v0_])
            d_1342_v52_ = (d_1342_v52_).set(len(d_1345_v55_), default__.abs((0) - (len(((d_1346_v56_)[default__.safeIndex(d_1269_v1_, len(d_1346_v56_))]) + (d_1268_v0_)))))
            d_1347_v57_: _dafny.Array
            nw255_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_1347_v57_ = nw255_
            d_1348_v58_: _dafny.Map
            d_1348_v58_ = _dafny.Map({d_1347_v57_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1349_i5_ in range(default__.abs(567))])})
            d_1350_v59_: _dafny.Map
            d_1350_v59_ = _dafny.Map({d_1269_v1_: d_1347_v57_})
            d_1351_v60_: _dafny.Seq
            d_1351_v60_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1352_v61_: D12
            d_1352_v61_ = D12_DC33(((d_1350_v59_)[len(d_1351_v60_)] if (len(d_1351_v60_)) in (d_1350_v59_) else d_1347_v57_))
            d_1268_v0_ = ((d_1348_v58_)[(d_1352_v61_).cf58] if ((d_1352_v61_).cf58) in (d_1348_v58_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utfdflun")))
            d_1353_v62_: _dafny.Set
            d_1353_v62_ = _dafny.Set({d_1278_i1_})
            d_1354_v63_: _dafny.Array
            nw256_ = _dafny.Array(None, 20)
            nw256_[int(0)] = p0
            nw256_[int(1)] = p0
            nw256_[int(2)] = (p0) == (p0)
            nw256_[int(3)] = (d_1279_v9_).isdisjoint(d_1279_v9_)
            nw256_[int(4)] = p0
            nw256_[int(5)] = p0
            nw256_[int(6)] = not(p0)
            nw256_[int(7)] = p0
            nw256_[int(8)] = p0
            nw256_[int(9)] = p0
            nw256_[int(10)] = not(not((p0 if p0 else p0)))
            nw256_[int(11)] = False
            nw256_[int(12)] = (d_1269_v1_) < (d_1269_v1_)
            nw256_[int(13)] = False
            nw256_[int(14)] = p0
            nw256_[int(15)] = (p0) == (p0)
            nw256_[int(16)] = p0
            nw256_[int(17)] = (d_1353_v62_).ispropersubset(default__.fm31(p0, globalState))
            nw256_[int(18)] = p0
            nw256_[int(19)] = ((d_1343_v53_)[default__.safeIndex(d_1278_i1_, len(d_1343_v53_))]) >= (len(d_1351_v60_))
            d_1354_v63_ = nw256_
            index262_ = default__.safeIndex(563, (d_1354_v63_).length(0))
            (d_1354_v63_)[index262_] = p0
            index263_ = default__.safeIndex(563, (d_1354_v63_).length(0))
            (d_1354_v63_)[index263_] = (-628) in (d_1342_v52_)
        hi10_ = d_1269_v1_
        for d_1355_i6_ in range(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfmtbfldy"))), hi10_):
            (globalState).f5 = (p0 if True else False)
            d_1356_v64_: _dafny.Array
            def lambda72_(d_1357_v1_):
                def lambda73_(d_1358_i7_):
                    return (d_1358_i7_) + (d_1357_v1_)

                return lambda73_

            init39_ = lambda72_(d_1269_v1_)
            nw257_ = _dafny.Array(None, 18)
            for i0_39_ in range(nw257_.length(0)):
                nw257_[i0_39_] = init39_(i0_39_)
            d_1356_v64_ = nw257_
            d_1359_v65_: _dafny.Seq
            d_1359_v65_ = _dafny.SeqWithoutIsStrInference([d_1356_v64_, d_1356_v64_])
            d_1359_v65_ = d_1359_v65_
            d_1360_v66_: _dafny.MultiSet
            d_1360_v66_ = _dafny.MultiSet([p0, p0, p0])
            d_1361_v67_: D11
            d_1361_v67_ = D11_DC30(p0, d_1355_i6_)
            d_1269_v1_ = ((d_1360_v66_)[(d_1361_v67_).cf51] if ((d_1361_v67_).cf51) in (d_1360_v66_) else d_1269_v1_)
            d_1269_v1_ = d_1355_i6_
        hi11_ = default__.safeDivisionInt(d_1269_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkxb"))))
        for d_1362_i8_ in range(d_1269_v1_, hi11_):
            d_1363_v68_: str
            d_1363_v68_ = _dafny.CodePoint('o')
            (globalState).f2 = (d_1363_v68_) not in (d_1268_v0_)
            d_1364_v69_: _dafny.Seq
            d_1364_v69_ = _dafny.SeqWithoutIsStrInference([d_1269_v1_, d_1362_i8_])
            d_1365_v70_: D3
            d_1365_v70_ = D3_DC9(d_1364_v69_)
            d_1366_v71_: _dafny.Map
            d_1366_v71_ = _dafny.Map({p0: (d_1365_v70_).cf18})
            d_1269_v1_ = (_dafny.MultiSet([len(((d_1366_v71_)[p0] if (p0) in (d_1366_v71_) else (self).fm6(default__.fm3(globalState), p0, d_1363_v68_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1367_i9_ in range(default__.abs(96))]), globalState)))])).cardinality
            d_1368_v72_: _dafny.Map
            d_1368_v72_ = _dafny.Map({p0: d_1362_i8_})
            d_1369_v73_: _dafny.Array
            nw258_ = _dafny.Array(_dafny.Seq({}), 8)
            d_1369_v73_ = nw258_
            d_1370_v74_: D2
            d_1370_v74_ = D2_DC8(d_1368_v72_, d_1369_v73_, p0, p0, d_1269_v1_)
            pat_let_tv29_ = p0
            pat_let_tv30_ = d_1369_v73_
            d_1371_v75_: D6
            def iife139_(_pat_let39_0):
                def iife140_(d_1372_dt__update__tmp_h2_):
                    def iife141_(_pat_let40_0):
                        def iife142_(d_1373_dt__update_hcf15_h0_):
                            def iife143_(_pat_let41_0):
                                def iife144_(d_1374_dt__update_hcf14_h0_):
                                    return D2_DC8((d_1372_dt__update__tmp_h2_).cf13, d_1374_dt__update_hcf14_h0_, d_1373_dt__update_hcf15_h0_, (d_1372_dt__update__tmp_h2_).cf16, (d_1372_dt__update__tmp_h2_).cf17)
                                return iife144_(_pat_let41_0)
                            return iife143_(pat_let_tv30_)
                        return iife142_(_pat_let40_0)
                    return iife141_(pat_let_tv29_)
                return iife140_(_pat_let39_0)
            d_1371_v75_ = D6_DC17(iife139_(d_1370_v74_))
            d_1375_v76_: _dafny.MultiSet
            d_1375_v76_ = _dafny.MultiSet([d_1371_v75_, d_1371_v75_])
            d_1376_v77_: _dafny.Seq
            d_1376_v77_ = _dafny.SeqWithoutIsStrInference([(d_1375_v76_).set(d_1371_v75_, default__.abs(d_1269_v1_))])
            d_1377_v78_: _dafny.Seq
            d_1377_v78_ = d_1376_v77_
            d_1377_v78_ = d_1377_v78_
            d_1378_v79_: C2
            nw259_ = C2()
            nw259_.ctor__(d_1269_v1_)
            d_1378_v79_ = nw259_
            d_1379_v80_: _dafny.Set
            d_1379_v80_ = _dafny.Set({d_1378_v79_, d_1378_v79_, d_1378_v79_, d_1378_v79_})
            d_1379_v80_ = d_1379_v80_
        hi12_ = d_1269_v1_
        for d_1380_i10_ in range(d_1269_v1_, hi12_):
            (globalState).f5 = not(False)
            if (d_1269_v1_) <= ((d_1269_v1_) + (d_1380_i10_)):
                d_1381_v81_: _dafny.Set
                d_1381_v81_ = _dafny.Set({default__.fm0(p0, p0, globalState), False})
                d_1382_v82_: _dafny.MultiSet
                d_1382_v82_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfw"))), len(_dafny.SeqWithoutIsStrInference([len(d_1381_v81_)])), len(default__.fm20(p0, not(p0), p0, d_1380_i10_, globalState)), -685, d_1380_i10_])
                d_1383_v83_: D1
                d_1383_v83_ = D1_DC5(d_1269_v1_, False)
                d_1384_v84_: C1
                nw260_ = C1()
                nw260_.ctor__(d_1380_i10_, ((d_1382_v82_)[(d_1383_v83_).cf10] if ((d_1383_v83_).cf10) in (d_1382_v82_) else d_1380_i10_), (p0 if not(p0) else p0), (_dafny.SeqWithoutIsStrInference([d_1380_i10_])) + (_dafny.SeqWithoutIsStrInference([d_1269_v1_, d_1269_v1_, (_dafny.MultiSet([787])).cardinality, d_1380_i10_])), d_1380_i10_)
                d_1384_v84_ = nw260_
                d_1385_v85_: _dafny.Array
                def lambda74_(d_1386_i11_):
                    return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1387_i12_ in range(default__.abs(483))])

                init40_ = lambda74_
                nw261_ = _dafny.Array(None, 14)
                for i0_40_ in range(nw261_.length(0)):
                    nw261_[i0_40_] = init40_(i0_40_)
                d_1385_v85_ = nw261_
                index264_ = default__.safeIndex(722, (d_1385_v85_).length(0))
                (d_1385_v85_)[index264_] = (d_1268_v0_) + (d_1268_v0_)
                index265_ = default__.safeIndex(722, (d_1385_v85_).length(0))
                rhs301_ = d_1269_v1_
                rhs302_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbdbgrpij"))) + (d_1268_v0_)
                lhs238_ = d_1385_v85_
                lhs239_ = default__.safeIndex(722, (d_1385_v85_).length(0))
                d_1269_v1_ = rhs301_
                lhs238_[lhs239_] = rhs302_
                d_1269_v1_ = d_1380_i10_
                d_1388_v86_: _dafny.Seq
                d_1388_v86_ = _dafny.SeqWithoutIsStrInference([(d_1384_v84_).f29])
                d_1269_v1_ = ((d_1382_v82_)[(d_1384_v84_).f30] if ((d_1384_v84_).f30) in (d_1382_v82_) else default__.safeDivisionInt(len(_dafny.Map({d_1388_v86_: p0})), -373))
                d_1389_v87_: _dafny.Array
                nw262_ = _dafny.Array(False, 12)
                d_1389_v87_ = nw262_
                d_1390_v88_: C0
                nw263_ = C0()
                nw263_.ctor__((d_1384_v84_).f30, d_1389_v87_)
                d_1390_v88_ = nw263_
            elif True:
                (globalState).f15 = default__.fm2(globalState)
                d_1391_v89_: _dafny.Array
                nw264_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_1391_v89_ = nw264_
                index266_ = default__.safeIndex(903, (d_1391_v89_).length(0))
                (d_1391_v89_)[index266_] = d_1268_v0_
                d_1392_v90_: str
                d_1392_v90_ = _dafny.CodePoint('a')
                index267_ = default__.safeIndex(903, (d_1391_v89_).length(0))
                (d_1391_v89_)[index267_] = (d_1268_v0_).set(default__.safeIndex(-864, len(d_1268_v0_)), d_1392_v90_)
                d_1393_v91_: _dafny.Set
                d_1393_v91_ = _dafny.Set({len(_dafny.Set({d_1269_v1_, default__.fm3(globalState)}))})
                d_1394_v92_: _dafny.Seq
                d_1394_v92_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1395_v93_: _dafny.Map
                d_1395_v93_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prrne")): d_1380_i10_})
                d_1396_v94_: _dafny.Map
                d_1396_v94_ = _dafny.Map({(d_1394_v92_)[default__.safeIndex(d_1269_v1_, len(d_1394_v92_))]: len((d_1395_v93_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pn")), d_1380_i10_))})
                d_1397_v95_: _dafny.Array
                def lambda75_(d_1398_i10_):
                    def lambda76_(d_1399_i13_):
                        return _dafny.SeqWithoutIsStrInference([d_1398_i10_])

                    return lambda76_

                init41_ = lambda75_(d_1380_i10_)
                nw265_ = _dafny.Array(None, 25)
                for i0_41_ in range(nw265_.length(0)):
                    nw265_[i0_41_] = init41_(i0_41_)
                d_1397_v95_ = nw265_
                d_1400_v96_: D2
                d_1400_v96_ = D2_DC8(d_1396_v94_, d_1397_v95_, p0, p0, d_1380_i10_)
                d_1401_v97_: _dafny.Set
                d_1401_v97_ = _dafny.Set({d_1400_v96_, d_1400_v96_})
                d_1402_v98_: D9
                d_1402_v98_ = D9_DC27(p0, (d_1393_v91_ if p0 else _dafny.Set({len(d_1401_v97_)})))
                d_1403_v99_: _dafny.MultiSet
                d_1403_v99_ = _dafny.MultiSet([p0])
                rhs303_ = ((d_1403_v99_).cardinality) >= (214)
                rhs304_ = d_1402_v98_
                rhs305_ = d_1380_i10_
                lhs240_ = globalState
                lhs240_.f2 = rhs303_
                d_1402_v98_ = rhs304_
                d_1269_v1_ = rhs305_
                (globalState).f5 = p0
                d_1392_v90_ = d_1392_v90_
            d_1404_v100_: _dafny.Map
            d_1404_v100_ = _dafny.Map({not(p0): (0) - (d_1269_v1_)})
            d_1269_v1_ = (0) - (default__.safeDivisionInt(default__.safeDivisionInt(len(d_1268_v0_), len(d_1404_v100_)), default__.safeModuloInt(d_1269_v1_, d_1380_i10_)))
            d_1405_v101_: _dafny.Array
            nw266_ = _dafny.Array(None, 8)
            nw266_[int(0)] = p0
            nw266_[int(1)] = p0
            nw266_[int(2)] = p0
            nw266_[int(3)] = p0
            nw266_[int(4)] = p0
            nw266_[int(5)] = p0
            nw266_[int(6)] = False
            nw266_[int(7)] = p0
            d_1405_v101_ = nw266_
            d_1406_v102_: _dafny.Array
            nw267_ = _dafny.Array(None, 6)
            nw267_[int(0)] = d_1405_v101_
            nw267_[int(1)] = ((self).f21).cf5
            nw267_[int(2)] = d_1405_v101_
            nw267_[int(3)] = d_1405_v101_
            nw267_[int(4)] = d_1405_v101_
            nw267_[int(5)] = d_1405_v101_
            d_1406_v102_ = nw267_
            d_1407_v103_: D6
            d_1407_v103_ = D6_DC15(d_1406_v102_)
            pat_let_tv31_ = d_1406_v102_
            d_1408_v104_: _dafny.Map
            def iife145_(_pat_let42_0):
                def iife146_(d_1409_dt__update__tmp_h3_):
                    def iife147_(_pat_let43_0):
                        def iife148_(d_1410_dt__update_hcf28_h0_):
                            return D6_DC15(d_1410_dt__update_hcf28_h0_)
                        return iife148_(_pat_let43_0)
                    return iife147_(pat_let_tv31_)
                return iife146_(_pat_let42_0)
            d_1408_v104_ = _dafny.Map({705: iife145_(d_1407_v103_)})
            d_1408_v104_ = (d_1408_v104_).set((d_1380_i10_) - (d_1380_i10_), d_1407_v103_)
        d_1411_v105_: _dafny.MultiSet
        d_1411_v105_ = _dafny.MultiSet([d_1269_v1_])
        d_1412_v106_: D4
        d_1412_v106_ = D4_DC12(p0, not(p0), d_1269_v1_)
        d_1413_v107_: _dafny.MultiSet
        d_1413_v107_ = _dafny.MultiSet([p0])
        d_1414_v108_: _dafny.Seq
        d_1414_v108_ = _dafny.SeqWithoutIsStrInference([d_1269_v1_])
        d_1415_v109_: C3
        nw268_ = C3()
        nw268_.ctor__(d_1268_v0_, p0, d_1414_v108_)
        d_1415_v109_ = nw268_
        d_1416_v110_: _dafny.Map
        d_1416_v110_ = _dafny.Map({d_1415_v109_: p0})
        d_1417_v111_: D8
        d_1417_v111_ = D8_DC21(d_1269_v1_, p0, True, p0, d_1269_v1_)
        d_1418_v112_: _dafny.Set
        d_1418_v112_ = _dafny.Set({d_1269_v1_})
        d_1419_v113_: _dafny.Map
        d_1419_v113_ = _dafny.Map({p0: d_1269_v1_})
        d_1420_v114_: D11
        d_1420_v114_ = D11_DC31(p0, d_1269_v1_, d_1269_v1_, p0)
        d_1421_v115_: _dafny.Seq
        d_1421_v115_ = _dafny.SeqWithoutIsStrInference([(d_1420_v114_).cf53])
        d_1422_v116_: _dafny.Map
        d_1422_v116_ = _dafny.Map({505: 617})
        d_1423_v117_: _dafny.Array
        nw269_ = _dafny.Array(None, 23)
        nw269_[int(0)] = (d_1269_v1_) + (len(d_1268_v0_))
        nw269_[int(1)] = d_1269_v1_
        nw269_[int(2)] = d_1269_v1_
        nw269_[int(3)] = d_1269_v1_
        nw269_[int(4)] = d_1269_v1_
        nw269_[int(5)] = ((d_1411_v105_)[d_1269_v1_] if (d_1269_v1_) in (d_1411_v105_) else (d_1412_v106_).cf23)
        nw269_[int(6)] = (d_1269_v1_) + ((d_1413_v107_).cardinality)
        nw269_[int(7)] = d_1269_v1_
        nw269_[int(8)] = d_1269_v1_
        nw269_[int(9)] = (0) - (d_1269_v1_)
        nw269_[int(10)] = d_1269_v1_
        nw269_[int(11)] = (0) - (d_1269_v1_)
        nw269_[int(12)] = len((d_1416_v110_) | ((d_1416_v110_).set(d_1415_v109_, p0)))
        nw269_[int(13)] = (0) - (len((_dafny.Map({(d_1417_v111_).cf37: len(d_1418_v112_)})) | (d_1419_v113_)))
        nw269_[int(14)] = (0) - (d_1269_v1_)
        nw269_[int(15)] = d_1269_v1_
        nw269_[int(16)] = ((default__.fm36(d_1269_v1_, p0, globalState)).intersection(d_1413_v107_)).cardinality
        nw269_[int(17)] = (d_1269_v1_ if p0 else d_1269_v1_)
        nw269_[int(18)] = (d_1269_v1_) + (len(d_1421_v115_))
        nw269_[int(19)] = d_1269_v1_
        nw269_[int(20)] = len(d_1422_v116_)
        nw269_[int(21)] = default__.safeDivisionInt(d_1269_v1_, 153)
        nw269_[int(22)] = d_1269_v1_
        d_1423_v117_ = nw269_
        index268_ = default__.safeIndex(346, (d_1423_v117_).length(0))
        (d_1423_v117_)[index268_] = d_1269_v1_
        index269_ = default__.safeIndex(346, (d_1423_v117_).length(0))
        (d_1423_v117_)[index269_] = (0) - (default__.fm3(globalState))

    def m4(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_1424_v0_: bool
        d_1424_v0_ = False
        if d_1424_v0_:
            source15_ = (self).f21
            if source15_.is_DC1:
                d_1425___mcc_h0_ = source15_.cf1
                d_1426_cf1_ = d_1425___mcc_h0_
                r3 = default__.fm3(globalState)
                d_1427_v1_: int
                d_1427_v1_ = 408
                r3 = (d_1427_v1_) * (d_1427_v1_)
                d_1428_v2_: _dafny.Map
                d_1428_v2_ = _dafny.Map({18: d_1427_v1_})
                d_1429_v3_: _dafny.Map
                d_1429_v3_ = _dafny.Map({d_1424_v0_: len(d_1428_v2_)})
                d_1430_v4_: _dafny.MultiSet
                d_1430_v4_ = _dafny.MultiSet([d_1429_v3_])
                d_1431_v5_: _dafny.Seq
                d_1431_v5_ = _dafny.SeqWithoutIsStrInference([d_1430_v4_, _dafny.MultiSet([d_1429_v3_])])
                d_1432_v6_: D4
                d_1432_v6_ = D4_DC11((d_1431_v5_)[default__.safeIndex(667, len(d_1431_v5_))])
                d_1432_v6_ = d_1432_v6_
                d_1433_v7_: _dafny.Seq
                d_1433_v7_ = _dafny.SeqWithoutIsStrInference([d_1427_v1_])
                d_1434_v8_: str
                d_1434_v8_ = _dafny.CodePoint('i')
                d_1435_v9_: D0
                d_1435_v9_ = D0_DC1((d_1426_cf1_).set(default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmnwsjn"))).set(default__.safeIndex(d_1427_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmnwsjn")))), d_1434_v8_)), len(d_1426_cf1_)), d_1434_v8_))
                d_1436_v10_: _dafny.Map
                d_1436_v10_ = _dafny.Map({d_1427_v1_: d_1435_v9_})
                d_1437_v11_: C1
                nw270_ = C1()
                nw270_.ctor__(931, d_1427_v1_, d_1424_v0_, d_1433_v7_, len(d_1436_v10_))
                d_1437_v11_ = nw270_
                d_1438_v12_: _dafny.Map
                d_1438_v12_ = _dafny.Map({d_1427_v1_: d_1437_v11_})
                d_1438_v12_ = (d_1438_v12_).set(len(d_1428_v2_), d_1437_v11_)
            elif source15_.is_DC2:
                d_1439___mcc_h1_ = source15_.cf2
                d_1440___mcc_h2_ = source15_.cf3
                d_1441___mcc_h3_ = source15_.cf4
                d_1442___mcc_h4_ = source15_.cf5
                d_1443_cf5_ = d_1442___mcc_h4_
                d_1444_cf4_ = d_1441___mcc_h3_
                d_1445_cf3_ = d_1440___mcc_h2_
                d_1446_cf2_ = d_1439___mcc_h1_
                d_1447_v13_: _dafny.Map
                d_1447_v13_ = _dafny.Map({d_1445_cf3_: d_1424_v0_})
                d_1448_v14_: _dafny.Seq
                d_1448_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxaubieu"))
                r2 = ((d_1447_v13_)[(0) - ((0) - ((d_1446_cf2_) + (len(d_1448_v14_))))] if ((0) - ((0) - ((d_1446_cf2_) + (len(d_1448_v14_))))) in (d_1447_v13_) else (_dafny.SeqWithoutIsStrInference([d_1444_cf4_ for d_1449_i0_ in range(default__.abs(-21))])) != (_dafny.SeqWithoutIsStrInference([d_1446_cf2_])))
                d_1450_v15_: _dafny.Map
                d_1450_v15_ = _dafny.Map({d_1424_v0_: d_1448_v14_})
                d_1451_v16_: _dafny.Map
                d_1451_v16_ = _dafny.Map({len(d_1450_v15_): d_1445_cf3_})
                d_1452_v17_: _dafny.Set
                d_1452_v17_ = _dafny.Set({d_1451_v16_})
                d_1452_v17_ = (d_1452_v17_) | (d_1452_v17_)
                d_1447_v13_ = d_1447_v13_
                d_1453_v18_: _dafny.Seq
                d_1453_v18_ = _dafny.SeqWithoutIsStrInference([d_1424_v0_, not(d_1424_v0_)])
                index270_ = default__.safeIndex(41, (d_1443_cf5_).length(0))
                (d_1443_cf5_)[index270_] = (d_1453_v18_)[default__.safeIndex(len(d_1453_v18_), len(d_1453_v18_))]
                d_1454_v19_: str
                d_1454_v19_ = _dafny.CodePoint('m')
                index271_ = default__.safeIndex(41, (d_1443_cf5_).length(0))
                rhs306_ = d_1454_v19_
                rhs307_ = d_1424_v0_
                rhs308_ = ((0) - (d_1446_cf2_)) * (default__.safeModuloInt(d_1444_cf4_, len(d_1453_v18_)))
                rhs309_ = d_1454_v19_
                lhs241_ = globalState
                lhs242_ = d_1443_cf5_
                lhs243_ = default__.safeIndex(41, (d_1443_cf5_).length(0))
                lhs244_ = globalState
                lhs241_.f15 = rhs306_
                lhs242_[lhs243_] = rhs307_
                d_1445_cf3_ = rhs308_
                lhs244_.f15 = rhs309_
            elif True:
                d_1455___mcc_h5_ = source15_.cf0
                d_1456_cf0_ = d_1455___mcc_h5_
                (globalState).f5 = d_1424_v0_
                d_1456_cf0_ = d_1456_cf0_
                r3 = len(d_1456_cf0_)
                d_1457_v20_: _dafny.Array
                nw271_ = _dafny.Array(_dafny.MultiSet({}), 14)
                d_1457_v20_ = nw271_
                d_1458_v21_: int
                d_1458_v21_ = -675
                d_1459_v22_: _dafny.MultiSet
                d_1459_v22_ = _dafny.MultiSet([d_1458_v21_, d_1458_v21_])
                d_1460_v23_: _dafny.Seq
                d_1460_v23_ = _dafny.SeqWithoutIsStrInference([d_1458_v21_, d_1458_v21_, d_1458_v21_])
                index272_ = default__.safeIndex(436, (d_1457_v20_).length(0))
                (d_1457_v20_)[index272_] = (_dafny.MultiSet([(d_1459_v22_).cardinality])).intersection(_dafny.MultiSet(d_1460_v23_))
                index273_ = default__.safeIndex(436, (d_1457_v20_).length(0))
                (d_1457_v20_)[index273_] = d_1459_v22_
            d_1461_v24_: _dafny.Array
            nw272_ = _dafny.Array(None, 24)
            d_1461_v24_ = nw272_
            d_1462_v25_: int
            d_1462_v25_ = 645
            d_1463_v26_: C4
            nw273_ = C4()
            nw273_.ctor__(52, d_1424_v0_)
            d_1463_v26_ = nw273_
            d_1464_v27_: str
            d_1464_v27_ = _dafny.CodePoint('u')
            d_1465_v28_: _dafny.Map
            d_1465_v28_ = _dafny.Map({d_1463_v26_: _dafny.CodePoint('m')})
            d_1466_v29_: _dafny.Seq
            d_1466_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1463_v26_: d_1464_v27_}), d_1465_v28_, d_1465_v28_])
            d_1467_v30_: _dafny.Seq
            d_1467_v30_ = _dafny.SeqWithoutIsStrInference([len((d_1466_v29_)[default__.safeIndex((d_1463_v26_).f22, len(d_1466_v29_))])])
            d_1468_v31_: C1
            nw274_ = C1()
            nw274_.ctor__(d_1462_v25_, d_1462_v25_, True, d_1467_v30_, len(d_1467_v30_))
            d_1468_v31_ = nw274_
            index274_ = default__.safeIndex(804, (d_1461_v24_).length(0))
            (d_1461_v24_)[index274_] = d_1468_v31_
            index275_ = default__.safeIndex(804, (d_1461_v24_).length(0))
            (d_1461_v24_)[index275_] = d_1468_v31_
            d_1469_v32_: _dafny.Array
            nw275_ = _dafny.Array(False, 12)
            d_1469_v32_ = nw275_
            d_1470_v33_: _dafny.MultiSet
            d_1470_v33_ = _dafny.MultiSet([d_1469_v32_, d_1469_v32_])
            d_1471_v34_: _dafny.Array
            nw276_ = _dafny.Array(int(0), 5)
            d_1471_v34_ = nw276_
            d_1472_v35_: _dafny.Set
            d_1472_v35_ = _dafny.Set({d_1471_v34_})
            d_1473_v36_: D11
            d_1473_v36_ = D11_DC30(d_1463_v26_.f23, (d_1463_v26_).f22)
            d_1474_v37_: _dafny.MultiSet
            d_1474_v37_ = _dafny.MultiSet([not((d_1473_v36_).cf51), d_1463_v26_.f23, d_1463_v26_.f23])
            d_1475_v38_: _dafny.Set
            d_1475_v38_ = _dafny.Set({d_1463_v26_.f23, not(d_1424_v0_)})
            d_1476_v39_: _dafny.Map
            d_1476_v39_ = _dafny.Map({d_1463_v26_.f23: (d_1463_v26_).f22})
            d_1477_v40_: _dafny.Array
            nw277_ = _dafny.Array(None, 16)
            nw277_[int(0)] = d_1462_v25_
            nw277_[int(1)] = (d_1468_v31_).f30
            nw277_[int(2)] = (d_1470_v33_).cardinality
            nw277_[int(3)] = (d_1468_v31_).fm15(_dafny.SeqWithoutIsStrInference([(d_1468_v31_).f30 for d_1478_i1_ in range(default__.abs(324))]), len(d_1472_v35_), d_1462_v25_, d_1464_v27_, globalState)
            nw277_[int(4)] = d_1462_v25_
            nw277_[int(5)] = (d_1468_v31_).f30
            nw277_[int(6)] = d_1462_v25_
            nw277_[int(7)] = (d_1474_v37_).cardinality
            nw277_[int(8)] = (0) - ((d_1462_v25_ if d_1463_v26_.f23 else len(d_1475_v38_)))
            nw277_[int(9)] = len(_dafny.Map({d_1469_v32_: d_1424_v0_}))
            nw277_[int(10)] = 32
            nw277_[int(11)] = (d_1468_v31_).f30
            nw277_[int(12)] = -898
            nw277_[int(13)] = (d_1468_v31_).f29
            nw277_[int(14)] = (d_1463_v26_).f22
            nw277_[int(15)] = ((d_1476_v39_)[d_1424_v0_] if (d_1424_v0_) in (d_1476_v39_) else 116)
            d_1477_v40_ = nw277_
            d_1471_v34_ = d_1471_v34_
            d_1479_v41_: _dafny.MultiSet
            d_1479_v41_ = _dafny.MultiSet([(d_1468_v31_).f30, (d_1468_v31_).f29])
            index276_ = default__.safeIndex(112, (d_1471_v34_).length(0))
            (d_1471_v34_)[index276_] = ((d_1479_v41_).intersection(d_1479_v41_)).cardinality
            index277_ = default__.safeIndex(112, (d_1471_v34_).length(0))
            (d_1471_v34_)[index277_] = (d_1468_v31_).f29
            d_1480_v42_: _dafny.Seq
            d_1480_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhpqpdgw"))
            d_1462_v25_ = (0) - (len(d_1480_v42_))
        elif True:
            d_1481_v43_: _dafny.Seq
            d_1481_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aebefuri"))
            d_1482_v44_: D0
            d_1482_v44_ = D0_DC0(d_1481_v43_)
            d_1482_v44_ = d_1482_v44_
            d_1483_v45_: int
            d_1483_v45_ = 700
            d_1484_v46_: _dafny.Seq
            d_1484_v46_ = _dafny.SeqWithoutIsStrInference([d_1483_v45_])
            d_1485_v47_: C4
            nw278_ = C4()
            nw278_.ctor__((_dafny.MultiSet([len(d_1484_v46_)])).cardinality, (d_1424_v0_ if d_1424_v0_ else default__.fm0(d_1424_v0_, d_1424_v0_, globalState)))
            d_1485_v47_ = nw278_
            d_1486_v48_: C2
            nw279_ = C2()
            nw279_.ctor__(len(d_1484_v46_))
            d_1486_v48_ = nw279_
            d_1487_v49_: _dafny.Map
            d_1487_v49_ = _dafny.Map({d_1424_v0_: d_1485_v47_.f23})
            d_1488_v50_: _dafny.Seq
            d_1488_v50_ = _dafny.SeqWithoutIsStrInference([((d_1487_v49_)[d_1424_v0_] if (d_1424_v0_) in (d_1487_v49_) else d_1424_v0_), d_1424_v0_])
            d_1489_v51_: _dafny.Map
            d_1489_v51_ = _dafny.Map({len(d_1488_v50_): d_1486_v48_})
            d_1490_v52_: _dafny.Map
            d_1490_v52_ = _dafny.Map({d_1424_v0_: len(d_1489_v51_)})
            d_1491_v53_: str
            d_1491_v53_ = _dafny.CodePoint('l')
            d_1492_v54_: _dafny.Map
            d_1492_v54_ = _dafny.Map({d_1491_v53_: d_1424_v0_})
            d_1493_v55_: C1
            nw280_ = C1()
            nw280_.ctor__(len(d_1490_v52_), (d_1486_v48_).f25, ((d_1492_v54_)[d_1491_v53_] if (d_1491_v53_) in (d_1492_v54_) else d_1485_v47_.f23), d_1484_v46_, (d_1486_v48_).f25)
            d_1493_v55_ = nw280_
            d_1494_v56_: C3
            nw281_ = C3()
            nw281_.ctor__(d_1481_v43_, d_1485_v47_.f23, (d_1484_v46_ if d_1424_v0_ else (self).fm6(822, not(not(d_1485_v47_.f23)), d_1491_v53_, d_1481_v43_, globalState)))
            d_1494_v56_ = nw281_
        d_1495_v57_: str
        d_1495_v57_ = _dafny.CodePoint('h')
        d_1496_v58_: int
        d_1496_v58_ = 461
        d_1497_v59_: _dafny.Seq
        d_1497_v59_ = _dafny.SeqWithoutIsStrInference([d_1424_v0_])
        d_1498_v60_: T0
        nw282_ = C1()
        nw282_.ctor__(d_1496_v58_, d_1496_v58_, d_1424_v0_, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_1497_v59_)).cardinality, d_1496_v58_, -126]), d_1496_v58_)
        d_1498_v60_ = nw282_
        d_1499_v61_: D8
        d_1499_v61_ = D8_DC22(d_1424_v0_, d_1498_v60_, d_1424_v0_)
        d_1500_v62_: _dafny.Map
        d_1500_v62_ = _dafny.Map({d_1495_v57_: (d_1499_v61_).cf40})
        d_1501_v63_: D9
        d_1501_v63_ = D9_DC25(d_1500_v62_)
        d_1502_v64_: _dafny.Set
        d_1502_v64_ = _dafny.Set({d_1501_v63_, d_1501_v63_, d_1501_v63_})
        d_1503_v65_: _dafny.Map
        d_1503_v65_ = _dafny.Map({d_1501_v63_: default__.fm3(globalState)})
        d_1504_v67_: C5
        nw283_ = C5()
        def iife149_():
            coll61_ = _dafny.Set()
            compr_61_: D9
            for compr_61_ in (d_1503_v65_).keys.Elements:
                d_1505_v66_: D9 = compr_61_
                if (d_1505_v66_) in (d_1503_v65_):
                    coll61_ = coll61_.union(_dafny.Set([d_1505_v66_]))
            return _dafny.Set(coll61_)
        nw283_.ctor__((d_1502_v64_) != (iife149_()
        ), _dafny.SeqWithoutIsStrInference([d_1496_v58_, d_1496_v58_, d_1496_v58_, d_1496_v58_]))
        d_1504_v67_ = nw283_
        d_1496_v58_ = default__.safeDivisionInt(d_1496_v58_, d_1496_v58_)
        d_1506_v68_: _dafny.Seq
        d_1506_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kp"))
        d_1507_v69_: D0
        d_1507_v69_ = D0_DC1(d_1506_v68_)
        d_1508_v70_: _dafny.MultiSet
        d_1508_v70_ = _dafny.MultiSet([987])
        d_1507_v69_ = default__.fm37(d_1424_v0_, (d_1508_v70_).cardinality, d_1424_v0_, globalState)
        hi13_ = d_1496_v58_
        for d_1509_i2_ in range(default__.fm3(globalState), hi13_):
            d_1510_v71_: D3
            d_1510_v71_ = D3_DC10((d_1497_v59_)[default__.safeIndex(d_1509_i2_, len(d_1497_v59_))])
            source16_ = d_1510_v71_
            if source16_.is_DC10:
                d_1511___mcc_h6_ = source16_.cf19
                d_1512_cf19_ = d_1511___mcc_h6_
                d_1513_v72_: _dafny.Map
                d_1513_v72_ = _dafny.Map({d_1496_v58_: default__.safeDivisionInt(d_1496_v58_, (0) - (d_1509_i2_))})
                d_1513_v72_ = (d_1513_v72_).set(d_1509_i2_, ((d_1498_v60_).f17)[default__.safeIndex(len(d_1506_v68_), len((d_1498_v60_).f17))])
                r3 = d_1496_v58_
                r3 = d_1496_v58_
                d_1514_v73_: _dafny.Set
                d_1514_v73_ = _dafny.Set({d_1496_v58_})
                d_1515_v74_: _dafny.Set
                d_1515_v74_ = _dafny.Set({d_1514_v73_})
                d_1516_v75_: _dafny.Map
                d_1516_v75_ = _dafny.Map({d_1495_v57_: d_1495_v57_})
                d_1517_v76_: _dafny.Array
                def lambda77_(d_1518_v60_):
                    def lambda78_(d_1519_i3_):
                        return not(d_1518_v60_.f16)

                    return lambda78_

                init42_ = lambda77_(d_1498_v60_)
                nw284_ = _dafny.Array(None, 4)
                for i0_42_ in range(nw284_.length(0)):
                    nw284_[i0_42_] = init42_(i0_42_)
                d_1517_v76_ = nw284_
                d_1520_v77_: _dafny.Map
                d_1520_v77_ = _dafny.Map({d_1509_i2_: d_1517_v76_})
                d_1521_v78_: _dafny.Array
                def lambda79_(d_1522_i2_):
                    def lambda80_(d_1523_i4_):
                        return (d_1523_i4_) * (d_1522_i2_)

                    return lambda80_

                init43_ = lambda79_(d_1509_i2_)
                nw285_ = _dafny.Array(None, 11)
                for i0_43_ in range(nw285_.length(0)):
                    nw285_[i0_43_] = init43_(i0_43_)
                d_1521_v78_ = nw285_
                d_1524_v79_: D12
                d_1524_v79_ = D12_DC34(((d_1520_v77_)[63] if (63) in (d_1520_v77_) else d_1517_v76_), d_1509_i2_, d_1521_v78_, d_1495_v57_)
                d_1525_v80_: _dafny.Array
                nw286_ = _dafny.Array(None, 26)
                nw286_[int(0)] = d_1498_v60_.f16
                nw286_[int(1)] = d_1512_cf19_
                nw286_[int(2)] = d_1498_v60_.f16
                nw286_[int(3)] = not(d_1498_v60_.f16)
                nw286_[int(4)] = (d_1509_i2_) > (default__.fm3(globalState))
                nw286_[int(5)] = (d_1515_v74_) == (d_1515_v74_)
                nw286_[int(6)] = default__.fm0(False, True, globalState)
                nw286_[int(7)] = default__.fm0(d_1424_v0_, False, globalState)
                nw286_[int(8)] = (d_1496_v58_) >= (d_1509_i2_)
                nw286_[int(9)] = False
                nw286_[int(10)] = d_1512_cf19_
                nw286_[int(11)] = d_1512_cf19_
                nw286_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_1495_v57_, ((d_1516_v75_)[(d_1524_v79_).cf62] if ((d_1524_v79_).cf62) in (d_1516_v75_) else default__.fm2(globalState)), d_1495_v57_, d_1495_v57_, d_1495_v57_])) != (d_1506_v68_)
                nw286_[int(13)] = not (d_1424_v0_) or (not(d_1424_v0_))
                nw286_[int(14)] = (d_1498_v60_.f16 if d_1498_v60_.f16 else d_1498_v60_.f16)
                nw286_[int(15)] = d_1498_v60_.f16
                nw286_[int(16)] = d_1498_v60_.f16
                nw286_[int(17)] = d_1498_v60_.f16
                nw286_[int(18)] = d_1512_cf19_
                nw286_[int(19)] = not (d_1498_v60_.f16) or (d_1498_v60_.f16)
                nw286_[int(20)] = d_1424_v0_
                nw286_[int(21)] = (d_1504_v67_).fm8(d_1506_v68_, globalState)
                nw286_[int(22)] = False
                nw286_[int(23)] = False
                nw286_[int(24)] = False
                nw286_[int(25)] = (d_1498_v60_.f16 if not(not(d_1424_v0_)) else d_1424_v0_)
                d_1525_v80_ = nw286_
                index278_ = default__.safeIndex(327, (d_1517_v76_).length(0))
                (d_1517_v76_)[index278_] = d_1498_v60_.f16
                index279_ = default__.safeIndex(327, (d_1517_v76_).length(0))
                (d_1517_v76_)[index279_] = d_1512_cf19_
            elif True:
                d_1526___mcc_h7_ = source16_.cf18
                d_1527_cf18_ = d_1526___mcc_h7_
                (globalState).f15 = d_1495_v57_
                d_1508_v70_ = d_1508_v70_
                d_1528_v81_: _dafny.MultiSet
                d_1528_v81_ = _dafny.MultiSet([d_1424_v0_, d_1424_v0_, False])
                d_1529_v82_: D1
                d_1529_v82_ = D1_DC3((_dafny.MultiSet(d_1497_v59_)).intersection(d_1528_v81_))
                d_1529_v82_ = (d_1529_v82_ if d_1498_v60_.f16 else d_1529_v82_)
                d_1530_v83_: _dafny.Map
                d_1530_v83_ = _dafny.Map({d_1496_v58_: d_1424_v0_})
                d_1531_v84_: _dafny.Map
                d_1531_v84_ = _dafny.Map({d_1495_v57_: (d_1530_v83_) | (d_1530_v83_)})
                d_1531_v84_ = (d_1531_v84_).set(d_1495_v57_, (default__.fm38(d_1501_v63_, d_1509_i2_, globalState)) | (d_1530_v83_))
            d_1532_v85_: C1
            nw287_ = C1()
            nw287_.ctor__(d_1509_i2_, d_1496_v58_, d_1498_v60_.f16, (d_1498_v60_).f17, d_1496_v58_)
            d_1532_v85_ = nw287_
            d_1533_v86_: D1
            d_1533_v86_ = D1_DC5(d_1496_v58_, d_1424_v0_)
            d_1424_v0_ = ((d_1533_v86_).cf10) > (default__.safeDivisionInt((d_1532_v85_).f30, len(_dafny.Set({(d_1504_v67_).fm8(d_1506_v68_, globalState), True}))))
            d_1534_v87_: _dafny.Set
            d_1534_v87_ = _dafny.Set({d_1424_v0_, d_1498_v60_.f16, d_1498_v60_.f16})
            d_1535_v88_: D11
            d_1535_v88_ = D11_DC31(d_1424_v0_, (d_1532_v85_).f29, (0) - ((d_1532_v85_).fm14(d_1509_i2_, d_1496_v58_, d_1495_v57_, globalState)), d_1424_v0_)
            d_1536_v89_: _dafny.Set
            d_1536_v89_ = _dafny.Set({d_1535_v88_})
            d_1537_v90_: _dafny.Seq
            d_1537_v90_ = _dafny.SeqWithoutIsStrInference([d_1536_v89_])
            d_1538_v91_: _dafny.Map
            d_1538_v91_ = _dafny.Map({(d_1532_v85_).f29: d_1498_v60_.f16})
            d_1539_v92_: _dafny.Array
            def lambda81_(d_1540_v60_):
                def lambda82_(d_1541_i5_):
                    return d_1540_v60_.f16

                return lambda82_

            init44_ = lambda81_(d_1498_v60_)
            nw288_ = _dafny.Array(None, 22)
            for i0_44_ in range(nw288_.length(0)):
                nw288_[i0_44_] = init44_(i0_44_)
            d_1539_v92_ = nw288_
            d_1542_v93_: C0
            nw289_ = C0()
            nw289_.ctor__(d_1496_v58_, d_1539_v92_)
            d_1542_v93_ = nw289_
            d_1543_v94_: _dafny.Array
            nw290_ = _dafny.Array(None, 9)
            nw290_[int(0)] = len(d_1534_v87_)
            nw290_[int(1)] = d_1509_i2_
            nw290_[int(2)] = d_1509_i2_
            nw290_[int(3)] = len((d_1537_v90_)[default__.safeIndex((d_1532_v85_).f29, len(d_1537_v90_))])
            nw290_[int(4)] = d_1496_v58_
            nw290_[int(5)] = d_1496_v58_
            nw290_[int(6)] = d_1496_v58_
            nw290_[int(7)] = (d_1532_v85_).f29
            nw290_[int(8)] = (D9_DC26(len(d_1538_v91_), d_1542_v93_)).cf45
            d_1543_v94_ = nw290_
            d_1544_v95_: D6
            d_1544_v95_ = D6_DC18(d_1543_v94_)
            source17_ = d_1544_v95_
            if source17_.is_DC16:
                d_1545___mcc_h8_ = source17_.cf29
                d_1546___mcc_h9_ = source17_.cf30
                d_1547_cf30_ = d_1546___mcc_h9_
                d_1548_cf29_ = d_1545___mcc_h8_
                d_1549_v96_: C4
                nw291_ = C4()
                nw291_.ctor__((d_1532_v85_).f29, d_1547_cf30_)
                d_1549_v96_ = nw291_
                r2 = d_1547_cf30_
                d_1550_v97_: _dafny.Map
                d_1550_v97_ = _dafny.Map({_dafny.CodePoint('r'): (0) - (d_1509_i2_)})
                d_1551_v98_: D11
                d_1551_v98_ = D11_DC29(d_1495_v57_)
                d_1550_v97_ = (d_1550_v97_).set((d_1551_v98_).cf50, (0) - ((d_1532_v85_).f29))
                d_1496_v58_ = (D1_DC5((d_1532_v85_).f29, False)).cf10
            elif source17_.is_DC17:
                d_1552___mcc_h10_ = source17_.cf31
                d_1553_cf31_ = d_1552___mcc_h10_
                d_1496_v58_ = (d_1532_v85_).f30
                d_1554_v99_: _dafny.Array
                nw292_ = _dafny.Array(_dafny.Set({}), 28)
                d_1554_v99_ = nw292_
                index280_ = default__.safeIndex(280, (d_1554_v99_).length(0))
                (d_1554_v99_)[index280_] = d_1534_v87_
                index281_ = default__.safeIndex(280, (d_1554_v99_).length(0))
                (d_1554_v99_)[index281_] = d_1534_v87_
                d_1496_v58_ = default__.safeDivisionInt((d_1532_v85_).f29, (d_1532_v85_).f30)
                def iife150_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(119, 403):
                        d_1555_v100_: int = compr_62_
                        if ((119) <= (d_1555_v100_)) and ((d_1555_v100_) < (403)):
                            coll62_[(d_1555_v100_) * ((d_1532_v85_).f29)] = d_1498_v60_.f16
                    return _dafny.Map(coll62_)
                r3 = (len((iife150_()
                ) | ((d_1538_v91_).set((d_1542_v93_).f26, d_1424_v0_)))) + ((d_1532_v85_).f29)
            elif source17_.is_DC18:
                d_1556___mcc_h11_ = source17_.cf32
                d_1557_cf32_ = d_1556___mcc_h11_
                d_1558_v101_: _dafny.Map
                d_1558_v101_ = _dafny.Map({(d_1532_v85_).f30: _dafny.CodePoint('m')})
                d_1559_v102_: _dafny.Array
                nw293_ = _dafny.Array(None, 16)
                nw293_[int(0)] = ((d_1558_v101_)[(d_1532_v85_).f29] if ((d_1532_v85_).f29) in (d_1558_v101_) else d_1495_v57_)
                nw293_[int(1)] = d_1495_v57_
                nw293_[int(2)] = d_1495_v57_
                nw293_[int(3)] = d_1495_v57_
                nw293_[int(4)] = d_1495_v57_
                nw293_[int(5)] = d_1495_v57_
                nw293_[int(6)] = d_1495_v57_
                nw293_[int(7)] = d_1495_v57_
                nw293_[int(8)] = d_1495_v57_
                nw293_[int(9)] = d_1495_v57_
                nw293_[int(10)] = d_1495_v57_
                nw293_[int(11)] = d_1495_v57_
                nw293_[int(12)] = _dafny.CodePoint('h')
                nw293_[int(13)] = d_1495_v57_
                nw293_[int(14)] = d_1495_v57_
                nw293_[int(15)] = d_1495_v57_
                d_1559_v102_ = nw293_
                index282_ = default__.safeIndex(631, (d_1559_v102_).length(0))
                (d_1559_v102_)[index282_] = d_1495_v57_
                index283_ = default__.safeIndex(631, (d_1559_v102_).length(0))
                (d_1559_v102_)[index283_] = d_1495_v57_
                index284_ = default__.safeIndex(531, (d_1543_v94_).length(0))
                (d_1543_v94_)[index284_] = ((d_1542_v93_).f26) * (d_1509_i2_)
                d_1560_v103_: _dafny.Array
                nw294_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_1560_v103_ = nw294_
                d_1561_v104_: _dafny.Map
                d_1561_v104_ = _dafny.Map({(d_1532_v85_).f30: (d_1498_v60_).f17})
                index285_ = default__.safeIndex(531, (d_1543_v94_).length(0))
                rhs310_ = (0) - ((d_1532_v85_).fm15(((d_1561_v104_)[(d_1508_v70_).cardinality] if ((d_1508_v70_).cardinality) in (d_1561_v104_) else (d_1498_v60_).f17), (d_1496_v58_) - ((d_1542_v93_).f26), (d_1542_v93_).f26, d_1495_v57_, globalState))
                rhs311_ = d_1560_v103_
                rhs312_ = (d_1532_v85_).f30
                lhs245_ = d_1543_v94_
                lhs246_ = default__.safeIndex(531, (d_1543_v94_).length(0))
                lhs245_[lhs246_] = rhs310_
                d_1560_v103_ = rhs311_
                d_1496_v58_ = rhs312_
                d_1562_v105_: C4
                nw295_ = C4()
                nw295_.ctor__(len((_dafny.SeqWithoutIsStrInference([d_1495_v57_ for d_1563_i6_ in range(default__.abs(241))])) + ((d_1506_v68_).set(default__.safeIndex((d_1532_v85_).f29, len(d_1506_v68_)), d_1495_v57_))), (False if d_1498_v60_.f16 else d_1424_v0_))
                d_1562_v105_ = nw295_
                d_1496_v58_ = default__.safeModuloInt(default__.safeModuloInt((d_1543_v94_)[default__.safeIndex(531, (d_1543_v94_).length(0))], 100), (d_1532_v85_).f29)
            elif True:
                d_1564___mcc_h12_ = source17_.cf28
                d_1565_cf28_ = d_1564___mcc_h12_
                d_1566_v106_: _dafny.MultiSet
                d_1566_v106_ = _dafny.MultiSet([d_1498_v60_.f16])
                d_1567_v107_: D0
                d_1567_v107_ = D0_DC2((d_1532_v85_).f29, (len(default__.fm20(not(d_1424_v0_), d_1498_v60_.f16, not(d_1498_v60_.f16), (D4_DC12(d_1424_v0_, True, (d_1566_v106_).cardinality)).cf23, globalState))) * ((d_1532_v85_).f30), len(d_1497_v59_), d_1542_v93_.f27)
                d_1567_v107_ = d_1567_v107_
                d_1568_v108_: _dafny.Array
                def lambda83_(d_1569_v57_):
                    def lambda84_(d_1570_i7_):
                        return d_1569_v57_

                    return lambda84_

                init45_ = lambda83_(d_1495_v57_)
                nw296_ = _dafny.Array(None, 19)
                for i0_45_ in range(nw296_.length(0)):
                    nw296_[i0_45_] = init45_(i0_45_)
                d_1568_v108_ = nw296_
                d_1571_v109_: T1
                nw297_ = C1()
                nw297_.ctor__((d_1542_v93_).f26, (d_1542_v93_).f26, d_1498_v60_.f16, ((d_1498_v60_).f17).set(default__.safeIndex((_dafny.MultiSet([d_1568_v108_, d_1568_v108_])).cardinality, len((d_1498_v60_).f17)), 875), (d_1532_v85_).f30)
                d_1571_v109_ = nw297_
                d_1572_v110_: _dafny.Map
                d_1572_v110_ = _dafny.Map({d_1571_v109_: not((d_1571_v109_.f16 if d_1498_v60_.f16 else d_1571_v109_.f16))})
                d_1573_v111_: _dafny.Seq
                d_1573_v111_ = _dafny.SeqWithoutIsStrInference([d_1571_v109_])
                d_1572_v110_ = (d_1572_v110_).set((d_1573_v111_)[default__.safeIndex((d_1532_v85_).f29, len(d_1573_v111_))], d_1571_v109_.f16)
                (d_1498_v60_).f16 = d_1498_v60_.f16
                d_1574_v112_: _dafny.Seq
                d_1574_v112_ = _dafny.SeqWithoutIsStrInference([d_1506_v68_])
                d_1497_v59_ = (d_1497_v59_).set(default__.safeIndex((d_1532_v85_).f30, len(d_1497_v59_)), ((d_1532_v85_).f29) != (len(d_1574_v112_)))
        d_1575_v113_: _dafny.Map
        d_1575_v113_ = _dafny.Map({d_1496_v58_: d_1496_v58_})
        hi14_ = d_1496_v58_
        for d_1576_i8_ in range((len(d_1575_v113_)) * (-870), hi14_):
            d_1577_v114_: D8
            d_1577_v114_ = D8_DC21(d_1496_v58_, (d_1576_i8_) <= (d_1576_i8_), (d_1496_v58_) != (d_1496_v58_), (d_1424_v0_) in (d_1497_v59_), (0) - (len(d_1506_v68_)))
            source18_ = d_1577_v114_
            if source18_.is_DC21:
                d_1578___mcc_h13_ = source18_.cf35
                d_1579___mcc_h14_ = source18_.cf36
                d_1580___mcc_h15_ = source18_.cf37
                d_1581___mcc_h16_ = source18_.cf38
                d_1582___mcc_h17_ = source18_.cf39
                d_1583_cf39_ = d_1582___mcc_h17_
                d_1584_cf38_ = d_1581___mcc_h16_
                d_1585_cf37_ = d_1580___mcc_h15_
                d_1586_cf36_ = d_1579___mcc_h14_
                d_1587_cf35_ = d_1578___mcc_h13_
                r3 = d_1583_cf39_
                d_1583_cf39_ = d_1583_cf39_
                d_1588_v115_: _dafny.Map
                d_1588_v115_ = _dafny.Map({d_1498_v60_.f16: d_1587_cf35_})
                d_1589_v116_: _dafny.Seq
                d_1589_v116_ = _dafny.SeqWithoutIsStrInference([d_1588_v115_])
                d_1589_v116_ = d_1589_v116_
                (d_1498_v60_).f16 = d_1498_v60_.f16
            elif source18_.is_DC22:
                d_1590___mcc_h18_ = source18_.cf40
                d_1591___mcc_h19_ = source18_.cf41
                d_1592___mcc_h20_ = source18_.cf42
                d_1593_cf42_ = d_1592___mcc_h20_
                d_1594_cf41_ = d_1591___mcc_h19_
                d_1595_cf40_ = d_1590___mcc_h18_
                (globalState).f5 = d_1424_v0_
                d_1596_v117_: _dafny.Map
                d_1596_v117_ = _dafny.Map({default__.fm0(d_1595_cf40_, d_1595_cf40_, globalState): d_1506_v68_})
                d_1597_v118_: _dafny.Seq
                d_1597_v118_ = _dafny.SeqWithoutIsStrInference([len(d_1596_v117_), (d_1496_v58_) * (d_1496_v58_), d_1496_v58_])
                d_1597_v118_ = (((d_1594_cf41_).f17).set(default__.safeIndex((0) - (len((d_1594_cf41_).f17)), len((d_1594_cf41_).f17)), d_1576_i8_) if False else default__.fm17(globalState))
                d_1598_v119_: C3
                nw298_ = C3()
                nw298_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gj")), (d_1504_v67_).fm8(_dafny.SeqWithoutIsStrInference([d_1495_v57_ for d_1599_i9_ in range(default__.abs(248))]), globalState), (d_1498_v60_).f17)
                d_1598_v119_ = nw298_
                d_1600_v120_: _dafny.Map
                d_1600_v120_ = _dafny.Map({d_1496_v58_: d_1498_v60_.f16})
                d_1601_v121_: C1
                nw299_ = C1()
                nw299_.ctor__(d_1576_i8_, len((d_1598_v119_).f24), d_1498_v60_.f16, d_1597_v118_, default__.safeModuloInt(len(d_1600_v120_), (d_1508_v70_).cardinality))
                d_1601_v121_ = nw299_
            elif source18_.is_DC23:
                d_1602_v122_: _dafny.Map
                d_1602_v122_ = _dafny.Map({d_1576_i8_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvrpwreug"))})
                d_1603_v123_: _dafny.Map
                d_1603_v123_ = _dafny.Map({-794: d_1498_v60_.f16})
                d_1604_v125_: _dafny.Seq
                def iife151_():
                    coll63_ = _dafny.Map()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(918, 620):
                        d_1605_v124_: int = compr_63_
                        if ((918) <= (d_1605_v124_)) and ((d_1605_v124_) < (620)):
                            coll63_[(d_1605_v124_) * (d_1576_i8_)] = d_1576_i8_
                    return _dafny.Map(coll63_)
                d_1604_v125_ = _dafny.SeqWithoutIsStrInference([d_1575_v113_, iife151_()
                ])
                d_1606_v126_: C3
                nw300_ = C3()
                nw300_.ctor__(d_1506_v68_, d_1498_v60_.f16, (d_1498_v60_).f17)
                d_1606_v126_ = nw300_
                d_1607_v127_: D12
                d_1607_v127_ = D12_DC35(False, d_1498_v60_.f16, (d_1604_v125_)[default__.safeIndex(d_1496_v58_, len(d_1604_v125_))], d_1498_v60_.f16, d_1606_v126_)
                d_1602_v122_ = (d_1602_v122_).set((0) - (len((d_1603_v123_) | (_dafny.Map({d_1496_v58_: (d_1607_v127_).cf63})))), (d_1606_v126_).f24)
                d_1608_v128_: _dafny.Seq
                out14_: _dafny.Seq
                out14_ = (d_1504_v67_).m5(globalState)
                d_1608_v128_ = out14_
                d_1506_v68_ = default__.fm1(globalState)
                d_1609_v129_: _dafny.Array
                nw301_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_1609_v129_ = nw301_
                d_1610_v130_: D6
                d_1610_v130_ = D6_DC15(d_1609_v129_)
                d_1611_v131_: _dafny.Array
                def lambda85_(d_1612_v60_):
                    def lambda86_(d_1613_i10_):
                        return (d_1612_v60_.f16) or (d_1612_v60_.f16)

                    return lambda86_

                init46_ = lambda85_(d_1498_v60_)
                nw302_ = _dafny.Array(None, 24)
                for i0_46_ in range(nw302_.length(0)):
                    nw302_[i0_46_] = init46_(i0_46_)
                d_1611_v131_ = nw302_
                index286_ = default__.safeIndex(317, (d_1611_v131_).length(0))
                (d_1611_v131_)[index286_] = ((d_1606_v126_).f24) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uncqkk")))
                index287_ = default__.safeIndex(742, (d_1611_v131_).length(0))
                (d_1611_v131_)[index287_] = default__.fm0(d_1498_v60_.f16, True, globalState)
                index288_ = default__.safeIndex(317, (d_1611_v131_).length(0))
                index289_ = default__.safeIndex(742, (d_1611_v131_).length(0))
                rhs313_ = d_1610_v130_
                rhs314_ = d_1498_v60_.f16
                rhs315_ = not((((0) - (d_1496_v58_)) == (len(d_1608_v128_))) and (d_1424_v0_))
                lhs247_ = d_1611_v131_
                lhs248_ = default__.safeIndex(317, (d_1611_v131_).length(0))
                lhs249_ = d_1611_v131_
                lhs250_ = default__.safeIndex(742, (d_1611_v131_).length(0))
                d_1610_v130_ = rhs313_
                lhs247_[lhs248_] = rhs314_
                lhs249_[lhs250_] = rhs315_
            elif source18_.is_DC20:
                d_1614___mcc_h21_ = source18_.cf34
                d_1615_cf34_ = d_1614___mcc_h21_
                d_1616_v132_: _dafny.Map
                d_1616_v132_ = _dafny.Map({(0) - ((d_1576_i8_) + (d_1576_i8_)): not(not((d_1498_v60_.f16 if d_1498_v60_.f16 else (d_1504_v67_).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwy")), globalState))))})
                d_1616_v132_ = (d_1616_v132_).set(d_1576_i8_, True)
                r3 = (794) + (d_1576_i8_)
                d_1617_v133_: _dafny.Array
                def lambda87_(d_1618_v60_, d_1619_v68_):
                    def lambda88_(d_1620_i11_):
                        return (d_1619_v68_ if d_1618_v60_.f16 else (D0_DC0(d_1619_v68_)).cf0)

                    return lambda88_

                init47_ = lambda87_(d_1498_v60_, d_1506_v68_)
                nw303_ = _dafny.Array(None, 19)
                for i0_47_ in range(nw303_.length(0)):
                    nw303_[i0_47_] = init47_(i0_47_)
                d_1617_v133_ = nw303_
                index290_ = default__.safeIndex(363, (d_1617_v133_).length(0))
                (d_1617_v133_)[index290_] = (d_1506_v68_) + (d_1506_v68_)
                d_1621_v134_: _dafny.Set
                d_1621_v134_ = _dafny.Set({((d_1616_v132_)[836] if (836) in (d_1616_v132_) else d_1424_v0_), d_1498_v60_.f16, d_1498_v60_.f16})
                index291_ = default__.safeIndex(363, (d_1617_v133_).length(0))
                rhs316_ = d_1498_v60_.f16
                rhs317_ = d_1506_v68_
                rhs318_ = d_1621_v134_
                rhs319_ = d_1496_v58_
                lhs251_ = globalState
                lhs252_ = d_1617_v133_
                lhs253_ = default__.safeIndex(363, (d_1617_v133_).length(0))
                lhs251_.f5 = rhs316_
                lhs252_[lhs253_] = rhs317_
                d_1621_v134_ = rhs318_
                r3 = rhs319_
                d_1622_v135_: _dafny.Array
                def lambda89_(d_1623_i8_):
                    def lambda90_(d_1624_i12_):
                        return (d_1624_i12_) * (d_1623_i8_)

                    return lambda90_

                init48_ = lambda89_(d_1576_i8_)
                nw304_ = _dafny.Array(None, 13)
                for i0_48_ in range(nw304_.length(0)):
                    nw304_[i0_48_] = init48_(i0_48_)
                d_1622_v135_ = nw304_
                d_1625_v136_: _dafny.Map
                d_1625_v136_ = _dafny.Map({True: d_1576_i8_})
                d_1626_v137_: D3
                d_1626_v137_ = D3_DC9(default__.fm17(globalState))
                d_1627_v138_: _dafny.Map
                d_1627_v138_ = _dafny.Map({d_1576_i8_: d_1495_v57_})
                d_1628_v139_: _dafny.Map
                d_1628_v139_ = _dafny.Map({d_1626_v137_: len(d_1627_v138_)})
                d_1629_v141_: _dafny.Array
                nw305_ = _dafny.Array(None, 25)
                nw305_[int(0)] = (d_1498_v60_).f17
                nw305_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1576_i8_ for d_1630_i13_ in range(default__.abs(727))])
                nw305_[int(2)] = _dafny.SeqWithoutIsStrInference([d_1576_i8_])
                nw305_[int(3)] = _dafny.SeqWithoutIsStrInference([len((d_1617_v133_)[default__.safeIndex(363, (d_1617_v133_).length(0))])])
                nw305_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1576_i8_ for d_1631_i14_ in range(default__.abs(906))])
                nw305_[int(5)] = (d_1498_v60_).f17
                nw305_[int(6)] = (d_1498_v60_).f17
                nw305_[int(7)] = (d_1498_v60_).f17
                nw305_[int(8)] = (d_1498_v60_).f17
                nw305_[int(9)] = (d_1498_v60_).f17
                nw305_[int(10)] = (d_1498_v60_).f17
                nw305_[int(11)] = (d_1498_v60_).f17
                nw305_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1496_v58_, d_1496_v58_, d_1576_i8_])
                nw305_[int(13)] = (d_1498_v60_).f17
                nw305_[int(14)] = (d_1498_v60_).f17
                nw305_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1576_i8_, d_1496_v58_, len(_dafny.SeqWithoutIsStrInference([d_1495_v57_ for d_1632_i15_ in range(default__.abs(320))]))])
                nw305_[int(16)] = (d_1498_v60_).f17
                nw305_[int(17)] = ((d_1498_v60_).f17).set(default__.safeIndex(d_1576_i8_, len((d_1498_v60_).f17)), len(d_1628_v139_))
                nw305_[int(18)] = (d_1498_v60_).f17
                def iife152_():
                    coll64_ = _dafny.Set()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(263, 597):
                        d_1633_v140_: int = compr_64_
                        if ((263) <= (d_1633_v140_)) and ((d_1633_v140_) < (597)):
                            coll64_ = coll64_.union(_dafny.Set([(d_1633_v140_) * (d_1576_i8_)]))
                    return _dafny.Set(coll64_)
                nw305_[int(19)] = _dafny.SeqWithoutIsStrInference([len(iife152_()
                )])
                nw305_[int(20)] = _dafny.SeqWithoutIsStrInference([d_1576_i8_])
                nw305_[int(21)] = (d_1498_v60_).f17
                nw305_[int(22)] = (d_1498_v60_).f17
                nw305_[int(23)] = (d_1498_v60_).f17
                nw305_[int(24)] = (d_1498_v60_).f17
                d_1629_v141_ = nw305_
                d_1634_v142_: _dafny.Seq
                d_1634_v142_ = _dafny.SeqWithoutIsStrInference([d_1629_v141_, d_1629_v141_, d_1629_v141_])
                d_1635_v143_: D2
                d_1635_v143_ = D2_DC8(d_1625_v136_, (d_1634_v142_)[default__.safeIndex(len((d_1506_v68_).set(default__.safeIndex(d_1496_v58_, len(d_1506_v68_)), d_1495_v57_)), len(d_1634_v142_))], d_1498_v60_.f16, not(d_1424_v0_), d_1576_i8_)
                d_1636_v144_: _dafny.Map
                d_1636_v144_ = _dafny.Map({d_1498_v60_.f16: d_1498_v60_.f16})
                d_1637_v145_: _dafny.Map
                d_1637_v145_ = _dafny.Map({(d_1635_v143_).cf16: (d_1636_v144_) | (d_1636_v144_)})
                d_1638_v146_: _dafny.Map
                d_1638_v146_ = _dafny.Map({d_1506_v68_: d_1622_v135_})
                index292_ = default__.safeIndex(363, (d_1617_v133_).length(0))
                rhs320_ = (((d_1617_v133_)[default__.safeIndex(363, (d_1617_v133_).length(0))]) + (d_1506_v68_)).set(default__.safeIndex(d_1576_i8_, len(((d_1617_v133_)[default__.safeIndex(363, (d_1617_v133_).length(0))]) + (d_1506_v68_))), d_1495_v57_)
                rhs321_ = ((d_1638_v146_)[d_1506_v68_] if (d_1506_v68_) in (d_1638_v146_) else d_1622_v135_)
                rhs322_ = (((d_1637_v145_).set(not(d_1498_v60_.f16), _dafny.Map({d_1498_v60_.f16: d_1498_v60_.f16}))) | (d_1637_v145_)) | (d_1637_v145_)
                lhs254_ = d_1617_v133_
                lhs255_ = default__.safeIndex(363, (d_1617_v133_).length(0))
                lhs254_[lhs255_] = rhs320_
                d_1622_v135_ = rhs321_
                d_1637_v145_ = rhs322_
            elif True:
                d_1639___mcc_h22_ = source18_.cf43
                d_1640_cf43_ = d_1639___mcc_h22_
                d_1641_v147_: _dafny.Array
                nw306_ = _dafny.Array(_dafny.Map({}), 10)
                d_1641_v147_ = nw306_
                d_1642_v148_: _dafny.Map
                d_1642_v148_ = _dafny.Map({d_1424_v0_: d_1498_v60_.f16})
                index293_ = default__.safeIndex(15, (d_1641_v147_).length(0))
                (d_1641_v147_)[index293_] = _dafny.Map({not(((d_1642_v148_)[d_1498_v60_.f16] if (d_1498_v60_.f16) in (d_1642_v148_) else d_1498_v60_.f16)): d_1496_v58_})
                d_1643_v149_: _dafny.Map
                d_1643_v149_ = _dafny.Map({d_1498_v60_.f16: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phnuuwd")))})
                index294_ = default__.safeIndex(15, (d_1641_v147_).length(0))
                (d_1641_v147_)[index294_] = (d_1643_v149_) | (d_1643_v149_)
                d_1644_v150_: _dafny.Map
                d_1644_v150_ = _dafny.Map({d_1504_v67_: d_1498_v60_.f16})
                d_1644_v150_ = (d_1644_v150_).set(d_1504_v67_, d_1498_v60_.f16)
                d_1645_v151_: _dafny.Array
                def lambda91_(d_1646_v68_):
                    def lambda92_(d_1647_i16_):
                        return (d_1647_i16_) * (len(d_1646_v68_))

                    return lambda92_

                init49_ = lambda91_(d_1506_v68_)
                nw307_ = _dafny.Array(None, 6)
                for i0_49_ in range(nw307_.length(0)):
                    nw307_[i0_49_] = init49_(i0_49_)
                d_1645_v151_ = nw307_
                index295_ = default__.safeIndex(934, (d_1645_v151_).length(0))
                (d_1645_v151_)[index295_] = d_1496_v58_
                index296_ = default__.safeIndex(934, (d_1645_v151_).length(0))
                (d_1645_v151_)[index296_] = d_1576_i8_
                (d_1498_v60_).f16 = True
            d_1648_v152_: _dafny.Set
            d_1648_v152_ = _dafny.Set({d_1498_v60_.f16, (d_1498_v60_.f16 if d_1424_v0_ else d_1498_v60_.f16)})
            d_1648_v152_ = (d_1648_v152_) | (d_1648_v152_)
            d_1496_v58_ = len(((d_1498_v60_).f17 if (d_1424_v0_) or (not(d_1424_v0_)) else (d_1498_v60_).f17))
            d_1496_v58_ = (d_1576_i8_) + (d_1496_v58_)
        d_1649_v153_: _dafny.Map
        d_1649_v153_ = _dafny.Map({d_1498_v60_.f16: d_1496_v58_})
        d_1650_v154_: _dafny.Map
        d_1650_v154_ = _dafny.Map({len(d_1649_v153_): d_1424_v0_})
        d_1651_v155_: _dafny.Seq
        d_1651_v155_ = _dafny.SeqWithoutIsStrInference([d_1650_v154_, d_1650_v154_])
        r0 = (d_1651_v155_) < ((d_1651_v155_) + (d_1651_v155_))
        d_1652_v156_: _dafny.MultiSet
        d_1652_v156_ = _dafny.MultiSet([d_1649_v153_, d_1649_v153_, _dafny.Map({d_1498_v60_.f16: d_1496_v58_})])
        pat_let_tv32_ = d_1508_v70_
        pat_let_tv33_ = d_1508_v70_
        pat_let_tv34_ = d_1498_v60_
        pat_let_tv35_ = d_1498_v60_
        def lambda93_(source19_):
            if source19_.is_DC12:
                d_1653___mcc_h23_ = source19_.cf21
                d_1654___mcc_h24_ = source19_.cf22
                d_1655___mcc_h25_ = source19_.cf23
                d_1656_cf23_ = d_1655___mcc_h25_
                d_1657_cf22_ = d_1654___mcc_h24_
                d_1658_cf21_ = d_1653___mcc_h23_
                return (pat_let_tv32_).issubset(pat_let_tv33_)
            elif source19_.is_DC13:
                d_1659___mcc_h26_ = source19_.cf24
                d_1660___mcc_h27_ = source19_.cf25
                d_1661___mcc_h28_ = source19_.cf26
                d_1662_cf26_ = d_1661___mcc_h28_
                d_1663_cf25_ = d_1660___mcc_h27_
                d_1664_cf24_ = d_1659___mcc_h26_
                return pat_let_tv34_.f16
            elif True:
                d_1665___mcc_h29_ = source19_.cf20
                d_1666_cf20_ = d_1665___mcc_h29_
                return pat_let_tv35_.f16

        r1 = lambda93_(D4_DC11(d_1652_v156_))
        r2 = not(d_1498_v60_.f16)
        r3 = default__.fm3(globalState)
        return r0, r1, r2, r3

    @property
    def f21(self):
        return self._f21

class C7(T0):
    def  __init__(self):
        self._f16: bool = False
        self._f17: _dafny.Seq = _dafny.Seq({})
        self._f20: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f20, f16, f17):
        (self)._f20 = f20
        (self).f16 = f16
        (self)._f17 = f17

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        d_1667_v0_: int
        d_1667_v0_ = 33
        d_1667_v0_ = default__.fm3(globalState)
        (globalState).f5 = not (self.f16) or ((p3) == ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1667_v0_]))).cardinality))
        d_1668_v1_: _dafny.Array
        nw308_ = _dafny.Array(False, 11)
        d_1668_v1_ = nw308_
        d_1669_v2_: D0
        d_1669_v2_ = D0_DC2(d_1667_v0_, d_1667_v0_, d_1667_v0_, d_1668_v1_)
        hi15_ = default__.safeModuloInt(929, (d_1669_v2_).cf3)
        for d_1670_i0_ in range((p3) + (d_1667_v0_), hi15_):
            d_1671_v3_: _dafny.Seq
            d_1671_v3_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1672_v4_: _dafny.Array
            nw309_ = _dafny.Array(None, 24)
            nw309_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            nw309_[int(1)] = p0
            nw309_[int(2)] = p0
            nw309_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1673_i1_ in range(default__.abs(802))])
            nw309_[int(4)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1674_i2_ in range(default__.abs(-237))])
            nw309_[int(5)] = p0
            nw309_[int(6)] = p0
            nw309_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "magnsasim"))
            nw309_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edeabhm"))
            nw309_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiq"))
            nw309_[int(10)] = p0
            nw309_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mul"))
            nw309_[int(12)] = p0
            nw309_[int(13)] = (p0) + (p0)
            nw309_[int(14)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1675_i3_ in range(default__.abs(619))])
            nw309_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbhpqphi")) if False else p0)
            nw309_[int(16)] = p0
            nw309_[int(17)] = p0
            nw309_[int(18)] = (p0) + (default__.fm1(globalState))
            nw309_[int(19)] = (p0) + (p0)
            nw309_[int(20)] = p0
            nw309_[int(21)] = (d_1671_v3_)[default__.safeIndex(d_1667_v0_, len(d_1671_v3_))]
            nw309_[int(22)] = p0
            nw309_[int(23)] = p0
            d_1672_v4_ = nw309_
            d_1672_v4_ = d_1672_v4_
            d_1676_v5_: _dafny.Map
            d_1676_v5_ = _dafny.Map({p0: self.f16})
            d_1676_v5_ = (d_1676_v5_).set((p0) + (p0), self.f16)
            d_1677_v6_: C3
            nw310_ = C3()
            nw310_.ctor__(p0, p1, _dafny.SeqWithoutIsStrInference([d_1670_i0_ for d_1678_i4_ in range(default__.abs(680))]))
            d_1677_v6_ = nw310_
            d_1679_v7_: D11
            d_1679_v7_ = D11_DC30(p2, (0) - (d_1667_v0_))
            d_1667_v0_ = ((d_1679_v7_).cf52) - (p3)
        if p1:
            d_1680_v8_: str
            d_1680_v8_ = _dafny.CodePoint('o')
            d_1681_v9_: _dafny.Map
            d_1681_v9_ = _dafny.Map({d_1680_v8_: p3})
            d_1667_v0_ = len(d_1681_v9_)
            d_1668_v1_ = d_1668_v1_
            d_1682_v10_: _dafny.Seq
            d_1682_v10_ = _dafny.SeqWithoutIsStrInference([d_1667_v0_])
            index297_ = default__.safeIndex(109, ((self).f20).length(0))
            ((self).f20)[index297_] = 252
            index298_ = default__.safeIndex(109, ((self).f20).length(0))
            rhs323_ = (((self).f17) + ((self).f17)) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([self.f16]))): _dafny.SeqWithoutIsStrInference([len(p0)])})) for d_1683_i5_ in range(default__.abs(271))]))
            rhs324_ = default__.fm3(globalState)
            rhs325_ = p3
            lhs256_ = (self).f20
            lhs257_ = default__.safeIndex(109, ((self).f20).length(0))
            d_1682_v10_ = rhs323_
            d_1667_v0_ = rhs324_
            lhs256_[lhs257_] = rhs325_
            d_1684_v11_: _dafny.Set
            d_1684_v11_ = _dafny.Set({default__.fm0(p2, p1, globalState)})
            d_1685_v12_: _dafny.Set
            d_1685_v12_ = _dafny.Set({((self).f20)[default__.safeIndex(109, ((self).f20).length(0))], p3, p3, default__.fm3(globalState)})
            d_1686_v13_: _dafny.Map
            d_1686_v13_ = _dafny.Map({d_1684_v11_: d_1685_v12_})
            d_1686_v13_ = (d_1686_v13_).set(d_1684_v11_, d_1685_v12_)
            d_1687_v14_: _dafny.Array
            def lambda94_(d_1688_i6_):
                return (d_1688_i6_) * (((self).f20)[default__.safeIndex(109, ((self).f20).length(0))])

            init50_ = lambda94_
            nw311_ = _dafny.Array(None, 26)
            for i0_50_ in range(nw311_.length(0)):
                nw311_[i0_50_] = init50_(i0_50_)
            d_1687_v14_ = nw311_
            d_1687_v14_ = d_1687_v14_
        elif True:
            d_1689_v15_: _dafny.Array
            nw312_ = _dafny.Array(int(0), 16)
            d_1689_v15_ = nw312_
            d_1690_v16_: str
            d_1690_v16_ = _dafny.CodePoint('q')
            d_1691_v17_: _dafny.Map
            d_1691_v17_ = _dafny.Map({d_1690_v16_: p3})
            d_1692_v18_: _dafny.Set
            d_1692_v18_ = _dafny.Set({p3})
            nw313_ = _dafny.Array(None, 12)
            nw313_[int(0)] = d_1667_v0_
            nw313_[int(1)] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eio")) if p2 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1693_i7_ in range(default__.abs(601))]))).set(default__.safeIndex(d_1667_v0_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eio")) if p2 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1694_i7_ in range(default__.abs(601))])))), d_1690_v16_))
            nw313_[int(2)] = ((d_1691_v17_)[d_1690_v16_] if (d_1690_v16_) in (d_1691_v17_) else d_1667_v0_)
            nw313_[int(3)] = d_1667_v0_
            nw313_[int(4)] = default__.safeModuloInt(p3, len(p0))
            nw313_[int(5)] = (d_1667_v0_ if p2 else p3)
            nw313_[int(6)] = (0) - ((0) - (p3))
            nw313_[int(7)] = default__.safeDivisionInt(d_1667_v0_, d_1667_v0_)
            nw313_[int(8)] = p3
            nw313_[int(9)] = len(d_1692_v18_)
            nw313_[int(10)] = p3
            nw313_[int(11)] = (956 if p2 else len(d_1692_v18_))
            d_1689_v15_ = nw313_
            d_1695_v19_: _dafny.Map
            d_1695_v19_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwadxtdi")): p1})
            d_1696_v20_: _dafny.Seq
            d_1696_v20_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_1697_v21_: D4
            d_1697_v21_ = D4_DC12(p2, default__.fm0(p1, p2, globalState), len(d_1696_v20_))
            d_1695_v19_ = (d_1695_v19_).set((p0) + (p0), (d_1697_v21_).cf21)
            d_1698_v22_: _dafny.Seq
            d_1698_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmrqcudhj"))
            d_1698_v22_ = d_1698_v22_
            d_1699_v23_: C1
            nw314_ = C1()
            nw314_.ctor__(p3, default__.safeDivisionInt(p3, len(d_1692_v18_)), p2, (self).f17, d_1667_v0_)
            d_1699_v23_ = nw314_
            d_1700_v24_: _dafny.Seq
            d_1700_v24_ = _dafny.SeqWithoutIsStrInference([d_1667_v0_])
            d_1701_v25_: _dafny.Map
            d_1701_v25_ = _dafny.Map({(d_1699_v23_).f29: not(self.f16)})
            d_1702_v26_: C0
            nw315_ = C0()
            nw315_.ctor__((d_1699_v23_).f29, d_1668_v1_)
            d_1702_v26_ = nw315_
            d_1703_v27_: _dafny.MultiSet
            d_1703_v27_ = _dafny.MultiSet([p1])
            rhs326_ = (default__.fm17(globalState)).set(default__.safeIndex((D9_DC26(len(d_1701_v25_), d_1702_v26_)).cf45, len(default__.fm17(globalState))), (d_1667_v0_) - ((d_1699_v23_).f30))
            rhs327_ = (0) - (((d_1702_v26_).f26) - (d_1667_v0_))
            rhs328_ = default__.fm0((not(self.f16) if p2 else p1), (d_1692_v18_).ispropersubset(d_1692_v18_), globalState)
            rhs329_ = (d_1703_v27_).issubset(d_1703_v27_)
            rhs330_ = d_1695_v19_
            lhs258_ = globalState
            lhs259_ = globalState
            d_1700_v24_ = rhs326_
            d_1667_v0_ = rhs327_
            lhs258_.f5 = rhs328_
            lhs259_.f2 = rhs329_
            d_1695_v19_ = rhs330_
        d_1704_v28_: D0
        d_1704_v28_ = D0_DC0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egdqwsb")))
        if ((d_1704_v28_).cf0) < (p0):
            d_1705_v29_: _dafny.Map
            d_1705_v29_ = _dafny.Map({d_1668_v1_: self.f16})
            d_1706_v30_: _dafny.Map
            d_1706_v30_ = d_1705_v29_
            d_1707_v31_: _dafny.Array
            nw316_ = _dafny.Array(None, 21)
            nw316_[int(0)] = d_1705_v29_
            nw316_[int(1)] = d_1705_v29_
            nw316_[int(2)] = _dafny.Map({d_1668_v1_: p1})
            nw316_[int(3)] = d_1705_v29_
            nw316_[int(4)] = d_1705_v29_
            nw316_[int(5)] = d_1705_v29_
            nw316_[int(6)] = d_1705_v29_
            nw316_[int(7)] = (d_1705_v29_) | (d_1705_v29_)
            nw316_[int(8)] = _dafny.Map({d_1668_v1_: p1})
            nw316_[int(9)] = (d_1705_v29_) | (d_1705_v29_)
            nw316_[int(10)] = d_1705_v29_
            nw316_[int(11)] = (d_1706_v30_)
            nw316_[int(12)] = d_1705_v29_
            nw316_[int(13)] = d_1705_v29_
            nw316_[int(14)] = d_1705_v29_
            nw316_[int(15)] = d_1705_v29_
            nw316_[int(16)] = (d_1705_v29_).set(d_1668_v1_, p2)
            nw316_[int(17)] = d_1705_v29_
            nw316_[int(18)] = _dafny.Map({d_1668_v1_: not(self.f16)})
            nw316_[int(19)] = (d_1705_v29_).set(d_1668_v1_, p2)
            nw316_[int(20)] = d_1705_v29_
            d_1707_v31_ = nw316_
            index299_ = default__.safeIndex(632, (d_1707_v31_).length(0))
            (d_1707_v31_)[index299_] = (d_1705_v29_) | (d_1705_v29_)
            index300_ = default__.safeIndex(632, (d_1707_v31_).length(0))
            (d_1707_v31_)[index300_] = d_1705_v29_
            d_1708_v32_: _dafny.Set
            d_1708_v32_ = _dafny.Set({(self).f20, (self).f20})
            d_1709_v33_: C5
            nw317_ = C5()
            nw317_.ctor__(self.f16, ((self).f17).set(default__.safeIndex(len(d_1708_v32_), len((self).f17)), p3))
            d_1709_v33_ = nw317_
            index301_ = default__.safeIndex(489, (d_1668_v1_).length(0))
            (d_1668_v1_)[index301_] = p2
            index302_ = default__.safeIndex(489, (d_1668_v1_).length(0))
            (d_1668_v1_)[index302_] = self.f16
            d_1710_v34_: _dafny.Map
            d_1710_v34_ = _dafny.Map({True: d_1667_v0_})
            d_1711_v35_: _dafny.Array
            def lambda95_(d_1712_i8_):
                return (self).f17

            init51_ = lambda95_
            nw318_ = _dafny.Array(None, 9)
            for i0_51_ in range(nw318_.length(0)):
                nw318_[i0_51_] = init51_(i0_51_)
            d_1711_v35_ = nw318_
            d_1713_v36_: D2
            d_1713_v36_ = D2_DC8(d_1710_v34_, d_1711_v35_, not(self.f16), p1, p3)
            source20_ = D6_DC17(d_1713_v36_)
            if source20_.is_DC16:
                d_1714___mcc_h0_ = source20_.cf29
                d_1715___mcc_h1_ = source20_.cf30
                d_1716_cf30_ = d_1715___mcc_h1_
                d_1717_cf29_ = d_1714___mcc_h0_
                d_1718_v37_: _dafny.Seq
                d_1718_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lutqh"))
                index303_ = default__.safeIndex(489, (d_1668_v1_).length(0))
                rhs331_ = d_1716_cf30_
                rhs332_ = (d_1667_v0_) + (d_1667_v0_)
                rhs333_ = (len((p0) + (d_1718_v37_))) <= (d_1667_v0_)
                rhs334_ = ((d_1709_v33_).fm7(_dafny.SeqWithoutIsStrInference([d_1717_cf29_, d_1717_cf29_]), globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "soa")))
                lhs260_ = globalState
                lhs261_ = d_1668_v1_
                lhs262_ = default__.safeIndex(489, (d_1668_v1_).length(0))
                lhs260_.f2 = rhs331_
                d_1667_v0_ = rhs332_
                lhs261_[lhs262_] = rhs333_
                d_1718_v37_ = rhs334_
                index304_ = default__.safeIndex(489, (d_1668_v1_).length(0))
                (d_1668_v1_)[index304_] = p1
                d_1719_v38_: _dafny.Map
                d_1719_v38_ = _dafny.Map({default__.fm3(globalState): _dafny.SeqWithoutIsStrInference([(self).f20])})
                d_1720_v39_: _dafny.Seq
                d_1720_v39_ = _dafny.SeqWithoutIsStrInference([(self).f20])
                d_1719_v38_ = (d_1719_v38_).set(p3, (d_1720_v39_) + (d_1720_v39_))
                d_1716_cf30_ = p1
            elif source20_.is_DC17:
                d_1721___mcc_h2_ = source20_.cf31
                d_1722_cf31_ = d_1721___mcc_h2_
                d_1667_v0_ = ((677) - ((0) - (p3))) - ((d_1667_v0_) + (p3))
                d_1723_v40_: _dafny.Map
                d_1723_v40_ = _dafny.Map({self.f16: (self).f20})
                d_1724_v41_: _dafny.Seq
                d_1724_v41_ = _dafny.SeqWithoutIsStrInference([(self).f20, (self).f20, (self).f20, (self).f20, (self).f20])
                d_1725_v42_: _dafny.MultiSet
                d_1725_v42_ = _dafny.MultiSet([(self).f20, (self).f20, (self).f20, ((d_1723_v40_)[(d_1668_v1_)[default__.safeIndex(489, (d_1668_v1_).length(0))]] if ((d_1668_v1_)[default__.safeIndex(489, (d_1668_v1_).length(0))]) in (d_1723_v40_) else (self).f20), (d_1724_v41_)[default__.safeIndex(p3, len(d_1724_v41_))]])
                d_1725_v42_ = ((d_1725_v42_) | (d_1725_v42_)) | (d_1725_v42_)
                d_1726_v43_: _dafny.Map
                d_1726_v43_ = _dafny.Map({p0: p1})
                d_1723_v40_ = (d_1723_v40_).set(not(((d_1726_v43_)[p0] if (p0) in (d_1726_v43_) else p2)), (self).f20)
                d_1727_v44_: _dafny.Array
                nw319_ = _dafny.Array(False, 11)
                d_1727_v44_ = nw319_
                d_1728_v45_: C0
                nw320_ = C0()
                nw320_.ctor__(p3, d_1727_v44_)
                d_1728_v45_ = nw320_
                index305_ = default__.safeIndex(489, (d_1668_v1_).length(0))
                rhs335_ = (True) and (p2)
                rhs336_ = d_1728_v45_
                lhs263_ = d_1668_v1_
                lhs264_ = default__.safeIndex(489, (d_1668_v1_).length(0))
                lhs263_[lhs264_] = rhs335_
                d_1728_v45_ = rhs336_
            elif source20_.is_DC18:
                d_1729___mcc_h3_ = source20_.cf32
                d_1730_cf32_ = d_1729___mcc_h3_
                index306_ = default__.safeIndex(629, (d_1730_cf32_).length(0))
                (d_1730_cf32_)[index306_] = p3
                index307_ = default__.safeIndex(629, (d_1730_cf32_).length(0))
                (d_1730_cf32_)[index307_] = p3
                (globalState).f2 = p1
                d_1731_v46_: C4
                nw321_ = C4()
                nw321_.ctor__(235, self.f16)
                d_1731_v46_ = nw321_
                (globalState).f5 = (d_1668_v1_)[default__.safeIndex(489, (d_1668_v1_).length(0))]
            elif True:
                d_1732___mcc_h4_ = source20_.cf28
                d_1733_cf28_ = d_1732___mcc_h4_
                index308_ = default__.safeIndex(489, (d_1668_v1_).length(0))
                (d_1668_v1_)[index308_] = (d_1668_v1_)[default__.safeIndex(489, (d_1668_v1_).length(0))]
                d_1734_v47_: D0
                d_1734_v47_ = D0_DC1(p0)
                (globalState).f5 = (p0) <= (((d_1734_v47_).cf1 if p2 else p0))
                d_1735_v48_: _dafny.MultiSet
                d_1735_v48_ = _dafny.MultiSet([d_1667_v0_, p3])
                d_1736_v49_: _dafny.MultiSet
                d_1736_v49_ = _dafny.MultiSet([(d_1735_v48_).intersection(d_1735_v48_), default__.fm19(d_1667_v0_, globalState), _dafny.MultiSet(default__.fm17(globalState))])
                d_1736_v49_ = default__.fm39((d_1668_v1_)[default__.safeIndex(489, (d_1668_v1_).length(0))], _dafny.SeqWithoutIsStrInference([p3 for d_1737_i9_ in range(default__.abs(498))]), (p0)[default__.safeIndex(p3, len(p0))], globalState)
                (self).m2(globalState)
            d_1738_v50_: _dafny.Map
            d_1738_v50_ = _dafny.Map({p3: p3})
            rhs337_ = not(not(not(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuu"))) + (p0)) < (p0))))
            rhs338_ = ((d_1738_v50_)[d_1667_v0_] if (d_1667_v0_) in (d_1738_v50_) else (d_1667_v0_ if (d_1668_v1_)[default__.safeIndex(489, (d_1668_v1_).length(0))] else d_1667_v0_))
            rhs339_ = p3
            lhs265_ = globalState
            lhs265_.f5 = rhs337_
            d_1667_v0_ = rhs338_
            d_1667_v0_ = rhs339_
        elif True:
            d_1739_v51_: str
            d_1739_v51_ = _dafny.CodePoint('q')
            if ((d_1667_v0_) * ((D12_DC34(d_1668_v1_, p3, (self).f20, d_1739_v51_)).cf60)) >= (p3):
                d_1740_v52_: D11
                d_1740_v52_ = D11_DC31(p1, len((_dafny.SeqWithoutIsStrInference([d_1739_v51_ for d_1741_i10_ in range(default__.abs(721))])) + (p0)), d_1667_v0_, p2)
                d_1740_v52_ = d_1740_v52_
                d_1742_v53_: _dafny.MultiSet
                d_1742_v53_ = _dafny.MultiSet([d_1667_v0_])
                default__.m0(((d_1742_v53_)[p3] if (p3) in (d_1742_v53_) else (0) - (default__.fm3(globalState))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "re")), ((self).f17)[default__.safeIndex(default__.fm3(globalState), len((self).f17))], d_1667_v0_, globalState)
                (globalState).f2 = self.f16
                d_1743_v54_: _dafny.Seq
                d_1743_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykvixfv"))
                d_1743_v54_ = (d_1743_v54_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ir")))
                d_1744_v55_: _dafny.Set
                d_1744_v55_ = _dafny.Set({p2})
                index309_ = default__.safeIndex(455, ((self).f20).length(0))
                ((self).f20)[index309_] = len((d_1744_v55_).intersection(d_1744_v55_))
                d_1745_v56_: _dafny.Seq
                d_1745_v56_ = _dafny.SeqWithoutIsStrInference([(self).f17, default__.fm17(globalState), (self).f17, _dafny.SeqWithoutIsStrInference([(0) - (d_1667_v0_) for d_1746_i11_ in range(default__.abs(-271))]), (self).f17])
                index310_ = default__.safeIndex(455, ((self).f20).length(0))
                rhs340_ = len(p0)
                rhs341_ = d_1745_v56_
                rhs342_ = d_1743_v54_
                lhs266_ = (self).f20
                lhs267_ = default__.safeIndex(455, ((self).f20).length(0))
                lhs266_[lhs267_] = rhs340_
                d_1745_v56_ = rhs341_
                d_1743_v54_ = rhs342_
            elif True:
                (globalState).f5 = not(p1)
                (globalState).f5 = ((p0) + (_dafny.SeqWithoutIsStrInference([d_1739_v51_ for d_1747_i12_ in range(default__.abs(914))]))) == (p0)
                d_1667_v0_ = p3
                d_1748_v57_: _dafny.Array
                nw322_ = _dafny.Array(None, 3)
                nw322_[int(0)] = d_1668_v1_
                nw322_[int(1)] = (d_1668_v1_ if p1 else d_1668_v1_)
                nw322_[int(2)] = d_1668_v1_
                d_1748_v57_ = nw322_
                index311_ = default__.safeIndex(297, (d_1748_v57_).length(0))
                (d_1748_v57_)[index311_] = d_1668_v1_
                index312_ = default__.safeIndex(297, (d_1748_v57_).length(0))
                (d_1748_v57_)[index312_] = d_1668_v1_
                d_1667_v0_ = d_1667_v0_
            if (d_1667_v0_) > (d_1667_v0_):
                d_1749_v58_: D2
                d_1749_v58_ = D2_DC7(_dafny.Set({self.f16, self.f16, p1}))
                d_1750_v59_: _dafny.Seq
                d_1750_v59_ = _dafny.SeqWithoutIsStrInference([d_1749_v58_])
                d_1751_v60_: _dafny.Seq
                d_1751_v60_ = _dafny.SeqWithoutIsStrInference([d_1750_v59_, d_1750_v59_])
                d_1752_v61_: _dafny.Map
                d_1752_v61_ = _dafny.Map({_dafny.MultiSet((d_1751_v60_)[default__.safeIndex(len(p0), len(d_1751_v60_))]): p3})
                d_1753_v62_: _dafny.Set
                d_1753_v62_ = _dafny.Set({p2, self.f16})
                d_1754_v63_: _dafny.MultiSet
                d_1754_v63_ = _dafny.MultiSet([D2_DC7(d_1753_v62_)])
                d_1752_v61_ = (d_1752_v61_).set(d_1754_v63_, d_1667_v0_)
                d_1755_v64_: _dafny.Seq
                d_1755_v64_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                d_1756_v65_: T0
                nw323_ = C1()
                nw323_.ctor__((0) - (p3), d_1667_v0_, p1, (d_1755_v64_)[default__.safeIndex(len(p0), len(d_1755_v64_))], 154)
                d_1756_v65_ = nw323_
                d_1756_v65_ = d_1756_v65_
                (self).m2(globalState)
                (globalState).f2 = d_1756_v65_.f16
                d_1757_v66_: _dafny.Seq
                d_1757_v66_ = _dafny.SeqWithoutIsStrInference([(d_1667_v0_) + (p3)])
                d_1758_v68_: _dafny.MultiSet
                d_1758_v68_ = _dafny.MultiSet([d_1739_v51_])
                def iife153_():
                    coll65_ = _dafny.Map()
                    compr_65_: str
                    for compr_65_ in (d_1758_v68_).Elements:
                        d_1759_v67_: str = compr_65_
                        if (d_1759_v67_) in (d_1758_v68_):
                            coll65_[d_1759_v67_] = len(d_1757_v66_)
                    return _dafny.Map(coll65_)
                rhs343_ = ((d_1757_v66_) + (d_1757_v66_)) + ((d_1756_v65_).f17)
                rhs344_ = (len(iife153_()
                )) + (default__.safeModuloInt(d_1667_v0_, len(p0)))
                rhs345_ = ((False) or (True)) == (self.f16)
                rhs346_ = p3
                rhs347_ = True
                lhs268_ = globalState
                lhs269_ = globalState
                d_1757_v66_ = rhs343_
                d_1667_v0_ = rhs344_
                lhs268_.f5 = rhs345_
                d_1667_v0_ = rhs346_
                lhs269_.f5 = rhs347_
            elif True:
                d_1760_v69_: _dafny.Map
                d_1760_v69_ = _dafny.Map({p1: p1})
                d_1761_v70_: D1
                d_1761_v70_ = D1_DC4(not(self.f16), d_1667_v0_, d_1760_v69_)
                d_1762_v71_: _dafny.Seq
                d_1762_v71_ = _dafny.SeqWithoutIsStrInference([(d_1761_v70_).cf9, d_1760_v69_, d_1760_v69_])
                index313_ = default__.safeIndex(605, (d_1668_v1_).length(0))
                (d_1668_v1_)[index313_] = (d_1762_v71_) != (d_1762_v71_)
                index314_ = default__.safeIndex(605, (d_1668_v1_).length(0))
                (d_1668_v1_)[index314_] = p2
                d_1763_v72_: _dafny.Map
                d_1763_v72_ = _dafny.Map({d_1667_v0_: p3})
                d_1764_v73_: C3
                nw324_ = C3()
                nw324_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")), p1, _dafny.SeqWithoutIsStrInference([len((self).f17)]))
                d_1764_v73_ = nw324_
                d_1765_v74_: D12
                d_1765_v74_ = D12_DC35(p1, (p3) < ((0) - (d_1667_v0_)), d_1763_v72_, (d_1668_v1_)[default__.safeIndex(605, (d_1668_v1_).length(0))], d_1764_v73_)
                index315_ = default__.safeIndex(605, (d_1668_v1_).length(0))
                index316_ = default__.safeIndex(605, (d_1668_v1_).length(0))
                rhs348_ = p1
                rhs349_ = (d_1668_v1_)[default__.safeIndex(605, (d_1668_v1_).length(0))]
                rhs350_ = self.f16
                rhs351_ = d_1765_v74_
                lhs270_ = globalState
                lhs271_ = d_1668_v1_
                lhs272_ = default__.safeIndex(605, (d_1668_v1_).length(0))
                lhs273_ = d_1668_v1_
                lhs274_ = default__.safeIndex(605, (d_1668_v1_).length(0))
                lhs270_.f5 = rhs348_
                lhs271_[lhs272_] = rhs349_
                lhs273_[lhs274_] = rhs350_
                d_1765_v74_ = rhs351_
                d_1766_v75_: _dafny.Set
                d_1766_v75_ = _dafny.Set({(0) - (d_1667_v0_), 406})
                rhs352_ = p2
                rhs353_ = (0) - (default__.safeDivisionInt((p3) * (len(d_1766_v75_)), (d_1667_v0_) * (p3)))
                lhs275_ = self
                lhs275_.f16 = rhs352_
                d_1667_v0_ = rhs353_
                d_1767_v76_: int
                d_1768_v77_: int
                d_1769_v78_: _dafny.Map
                out15_: int
                out16_: int
                out17_: _dafny.Map
                out15_, out16_, out17_ = (d_1764_v73_).m8(globalState)
                d_1767_v76_ = out15_
                d_1768_v77_ = out16_
                d_1769_v78_ = out17_
                d_1770_v79_: C2
                nw325_ = C2()
                nw325_.ctor__(d_1667_v0_)
                d_1770_v79_ = nw325_
            d_1771_v80_: _dafny.Map
            d_1771_v80_ = _dafny.Map({d_1668_v1_: default__.fm0(p2, self.f16, globalState)})
            d_1772_v81_: _dafny.Seq
            d_1772_v81_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1771_v80_: False})])
            d_1773_v82_: C1
            nw326_ = C1()
            nw326_.ctor__(len(d_1772_v81_), d_1667_v0_, p2, (self).f17, d_1667_v0_)
            d_1773_v82_ = nw326_
            d_1774_v83_: _dafny.Seq
            d_1774_v83_ = _dafny.SeqWithoutIsStrInference([d_1773_v82_, d_1773_v82_])
            d_1775_v84_: _dafny.Seq
            d_1775_v84_ = _dafny.SeqWithoutIsStrInference([(d_1774_v83_)[default__.safeIndex((d_1773_v82_).f30, len(d_1774_v83_))], d_1773_v82_, d_1773_v82_])
            d_1776_v85_: _dafny.MultiSet
            d_1776_v85_ = _dafny.MultiSet([(d_1773_v82_).f30, 128])
            d_1777_v86_: _dafny.Array
            nw327_ = _dafny.Array(None, 14)
            nw327_[int(0)] = d_1773_v82_
            nw327_[int(1)] = d_1773_v82_
            nw327_[int(2)] = d_1773_v82_
            nw327_[int(3)] = d_1773_v82_
            nw327_[int(4)] = (d_1775_v84_)[default__.safeIndex((d_1776_v85_).cardinality, len(d_1775_v84_))]
            nw327_[int(5)] = d_1773_v82_
            nw327_[int(6)] = d_1773_v82_
            nw327_[int(7)] = d_1773_v82_
            nw327_[int(8)] = d_1773_v82_
            nw327_[int(9)] = d_1773_v82_
            nw327_[int(10)] = d_1773_v82_
            nw327_[int(11)] = d_1773_v82_
            nw327_[int(12)] = d_1773_v82_
            nw327_[int(13)] = (d_1775_v84_)[default__.safeIndex(943, len(d_1775_v84_))]
            d_1777_v86_ = nw327_
            index317_ = default__.safeIndex(248, (d_1777_v86_).length(0))
            (d_1777_v86_)[index317_] = d_1773_v82_
            index318_ = default__.safeIndex(248, (d_1777_v86_).length(0))
            (d_1777_v86_)[index318_] = d_1773_v82_
            d_1778_v87_: T0
            nw328_ = C1()
            nw328_.ctor__(default__.safeModuloInt(len((self).f17), -478), (0) - (((self).f17)[default__.safeIndex(d_1667_v0_, len((self).f17))]), False, ((self).f17).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jstytmew"))), len((self).f17)), (d_1773_v82_).f30), (d_1773_v82_).f29)
            d_1778_v87_ = nw328_
            d_1778_v87_ = d_1778_v87_
            d_1779_v88_: _dafny.Array
            nw329_ = _dafny.Array(_dafny.CodePoint('D'), 26)
            d_1779_v88_ = nw329_
            index319_ = default__.safeIndex(508, (d_1779_v88_).length(0))
            (d_1779_v88_)[index319_] = d_1739_v51_
            index320_ = default__.safeIndex(508, (d_1779_v88_).length(0))
            (d_1779_v88_)[index320_] = d_1739_v51_
        hi16_ = len((self).f17)
        for d_1780_i13_ in range(d_1667_v0_, hi16_):
            d_1781_v89_: _dafny.Array
            def lambda96_(d_1782_v0_):
                def lambda97_(d_1783_i14_):
                    return _dafny.Map({d_1782_v0_: (0) - (-813)})

                return lambda97_

            init52_ = lambda96_(d_1667_v0_)
            nw330_ = _dafny.Array(None, 6)
            for i0_52_ in range(nw330_.length(0)):
                nw330_[i0_52_] = init52_(i0_52_)
            d_1781_v89_ = nw330_
            d_1784_v90_: _dafny.Set
            d_1784_v90_ = _dafny.Set({-759})
            d_1785_v91_: _dafny.Set
            d_1785_v91_ = _dafny.Set({p3, len(d_1784_v90_), p3})
            d_1786_v92_: _dafny.MultiSet
            d_1786_v92_ = _dafny.MultiSet([d_1667_v0_])
            d_1787_v93_: _dafny.Map
            d_1787_v93_ = _dafny.Map({d_1667_v0_: d_1786_v92_})
            rhs354_ = (len(p0)) not in (d_1784_v90_)
            rhs355_ = default__.safeDivisionInt(len(d_1787_v93_), d_1780_i13_)
            rhs356_ = d_1781_v89_
            rhs357_ = p1
            lhs276_ = self
            lhs277_ = globalState
            lhs276_.f16 = rhs354_
            d_1667_v0_ = rhs355_
            d_1781_v89_ = rhs356_
            lhs277_.f2 = rhs357_
            (globalState).f5 = self.f16
            d_1788_v94_: _dafny.Map
            d_1788_v94_ = _dafny.Map({self.f16: p1})
            d_1789_v95_: _dafny.Map
            d_1789_v95_ = _dafny.Map({len(p0): len(d_1788_v94_)})
            d_1667_v0_ = ((d_1789_v95_)[default__.fm3(globalState)] if (default__.fm3(globalState)) in (d_1789_v95_) else p3)
            if False:
                d_1790_v96_: D1
                d_1790_v96_ = D1_DC3(_dafny.MultiSet([p1]))
                d_1791_v97_: _dafny.Map
                d_1791_v97_ = _dafny.Map({(d_1667_v0_) - (len(p0)): d_1790_v96_})
                d_1791_v97_ = (d_1791_v97_).set(d_1780_i13_, d_1790_v96_)
                d_1792_v98_: _dafny.MultiSet
                d_1792_v98_ = _dafny.MultiSet([d_1668_v1_, d_1668_v1_, d_1668_v1_, d_1668_v1_, d_1668_v1_])
                rhs358_ = not(((d_1788_v94_)[(self.f16 if p2 else p2)] if ((self.f16 if p2 else p2)) in (d_1788_v94_) else self.f16))
                rhs359_ = d_1792_v98_
                lhs278_ = globalState
                lhs278_.f5 = rhs358_
                d_1792_v98_ = rhs359_
                index321_ = default__.safeIndex(903, ((self).f20).length(0))
                ((self).f20)[index321_] = default__.safeModuloInt(d_1667_v0_, d_1667_v0_)
                d_1793_v99_: str
                d_1793_v99_ = _dafny.CodePoint('g')
                index322_ = default__.safeIndex(903, ((self).f20).length(0))
                ((self).f20)[index322_] = (((self).f17)[default__.safeIndex((_dafny.MultiSet([d_1793_v99_, d_1793_v99_, d_1793_v99_])).cardinality, len((self).f17))]) - (p3)
                index323_ = default__.safeIndex(903, ((self).f20).length(0))
                ((self).f20)[index323_] = d_1780_i13_
                index324_ = default__.safeIndex(903, ((self).f20).length(0))
                ((self).f20)[index324_] = ((self).f20)[default__.safeIndex(903, ((self).f20).length(0))]
            elif True:
                (globalState).f5 = p2
                d_1794_v100_: _dafny.Map
                d_1794_v100_ = _dafny.Map({p2: d_1667_v0_})
                d_1794_v100_ = d_1794_v100_
                pat_let_tv36_ = d_1667_v0_
                pat_let_tv37_ = d_1668_v1_
                pat_let_tv38_ = d_1668_v1_
                d_1795_v101_: C6
                nw331_ = C6()
                def iife155_(_pat_let45_0):
                    def iife156_(d_1796_dt__update__tmp_h0_):
                        def iife157_(_pat_let46_0):
                            def iife158_(d_1797_dt__update_hcf3_h0_):
                                def iife159_(_pat_let47_0):
                                    def iife160_(d_1798_dt__update_hcf2_h0_):
                                        def iife161_(_pat_let48_0):
                                            def iife162_(d_1799_dt__update_hcf5_h0_):
                                                return D0_DC2(d_1798_dt__update_hcf2_h0_, d_1797_dt__update_hcf3_h0_, (d_1796_dt__update__tmp_h0_).cf4, d_1799_dt__update_hcf5_h0_)
                                            return iife162_(_pat_let48_0)
                                        return iife161_(pat_let_tv37_)
                                    return iife160_(_pat_let47_0)
                                return iife159_((0) - (-269))
                            return iife158_(_pat_let46_0)
                        return iife157_(pat_let_tv36_)
                    return iife156_(_pat_let45_0)
                def iife154_(_pat_let44_0):
                    def iife163_(d_1800_dt__update__tmp_h1_):
                        def iife164_(_pat_let49_0):
                            def iife165_(d_1801_dt__update_hcf5_h1_):
                                def iife166_(_pat_let50_0):
                                    def iife167_(d_1802_dt__update_hcf3_h1_):
                                        return D0_DC2((d_1800_dt__update__tmp_h1_).cf2, d_1802_dt__update_hcf3_h1_, (d_1800_dt__update__tmp_h1_).cf4, d_1801_dt__update_hcf5_h1_)
                                    return iife167_(_pat_let50_0)
                                return iife166_(d_1780_i13_)
                            return iife165_(_pat_let49_0)
                        return iife164_(pat_let_tv38_)
                    return iife163_(_pat_let44_0)
                nw331_.ctor__(iife154_(iife155_(d_1669_v2_)))
                d_1795_v101_ = nw331_
                d_1803_v102_: _dafny.Map
                d_1803_v102_ = _dafny.Map({(-131 if False else p3): (d_1795_v101_ if False else d_1795_v101_)})
                d_1804_v103_: _dafny.Seq
                d_1804_v103_ = _dafny.SeqWithoutIsStrInference([d_1795_v101_])
                d_1795_v101_ = ((d_1803_v102_)[153] if (153) in (d_1803_v102_) else (d_1804_v103_)[default__.safeIndex(((d_1789_v95_)[p3] if (p3) in (d_1789_v95_) else (0) - (len(d_1794_v100_))), len(d_1804_v103_))])
                d_1805_v104_: str
                d_1805_v104_ = _dafny.CodePoint('p')
                d_1806_v105_: _dafny.MultiSet
                d_1806_v105_ = _dafny.MultiSet([not((d_1805_v104_) not in (p0))])
                d_1806_v105_ = (d_1806_v105_).intersection((d_1806_v105_).set(self.f16, default__.abs(default__.fm3(globalState))))
                d_1807_v106_: _dafny.Seq
                d_1807_v106_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1808_v107_: _dafny.Map
                d_1808_v107_ = _dafny.Map({self.f16: (d_1807_v106_) + (d_1807_v106_)})
                d_1808_v107_ = (d_1808_v107_).set((p0) == (_dafny.SeqWithoutIsStrInference([d_1805_v104_ for d_1809_i15_ in range(default__.abs(77))])), d_1807_v106_)
        d_1810_v108_: _dafny.Seq
        d_1810_v108_ = _dafny.SeqWithoutIsStrInference([False])
        r0 = d_1810_v108_
        return r0

    def m2(self, globalState):
        d_1811_v0_: int
        d_1811_v0_ = 411
        d_1812_v1_: _dafny.Seq
        d_1812_v1_ = _dafny.SeqWithoutIsStrInference([self.f16])
        d_1813_v2_: _dafny.MultiSet
        d_1813_v2_ = _dafny.MultiSet([(0) - (len(d_1812_v1_)), d_1811_v0_, d_1811_v0_, d_1811_v0_, d_1811_v0_])
        d_1814_v3_: _dafny.Array
        nw332_ = _dafny.Array(None, 25)
        nw332_[int(0)] = self.f16
        nw332_[int(1)] = self.f16
        nw332_[int(2)] = self.f16
        nw332_[int(3)] = self.f16
        nw332_[int(4)] = self.f16
        nw332_[int(5)] = self.f16
        nw332_[int(6)] = False
        nw332_[int(7)] = (d_1813_v2_).issubset(d_1813_v2_)
        nw332_[int(8)] = (d_1811_v0_) > (d_1811_v0_)
        nw332_[int(9)] = self.f16
        nw332_[int(10)] = self.f16
        nw332_[int(11)] = self.f16
        nw332_[int(12)] = self.f16
        nw332_[int(13)] = self.f16
        nw332_[int(14)] = default__.fm0(not(self.f16), self.f16, globalState)
        nw332_[int(15)] = (d_1813_v2_).issubset(d_1813_v2_)
        nw332_[int(16)] = self.f16
        nw332_[int(17)] = (self.f16) == (self.f16)
        nw332_[int(18)] = (d_1812_v1_)[default__.safeIndex(d_1811_v0_, len(d_1812_v1_))]
        nw332_[int(19)] = self.f16
        nw332_[int(20)] = False
        nw332_[int(21)] = self.f16
        nw332_[int(22)] = self.f16
        nw332_[int(23)] = self.f16
        nw332_[int(24)] = self.f16
        d_1814_v3_ = nw332_
        d_1815_v4_: D1
        d_1815_v4_ = D1_DC5(520, False)
        index325_ = default__.safeIndex(847, (d_1814_v3_).length(0))
        (d_1814_v3_)[index325_] = (self.f16 if self.f16 else (d_1815_v4_).cf11)
        d_1816_v5_: D0
        d_1816_v5_ = D0_DC2(184, (d_1813_v2_).cardinality, d_1811_v0_, d_1814_v3_)
        d_1817_v6_: _dafny.Seq
        d_1817_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv"))
        index326_ = default__.safeIndex(847, (d_1814_v3_).length(0))
        rhs360_ = (d_1816_v5_).cf2
        rhs361_ = ((d_1811_v0_) - (d_1811_v0_)) == (len(d_1817_v6_))
        lhs279_ = d_1814_v3_
        lhs280_ = default__.safeIndex(847, (d_1814_v3_).length(0))
        d_1811_v0_ = rhs360_
        lhs279_[lhs280_] = rhs361_
        d_1818_i0_: int
        d_1818_i0_ = 0
        with _dafny.label("9"):
            while self.f16:
                with _dafny.c_label("9"):
                    if (d_1818_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1818_i0_ = (d_1818_i0_) + (1)
                    d_1819_v7_: str
                    d_1819_v7_ = _dafny.CodePoint('r')
                    d_1820_v8_: _dafny.Map
                    d_1820_v8_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdxlfemr")): d_1819_v7_})
                    d_1820_v8_ = (d_1820_v8_).set(d_1817_v6_, d_1819_v7_)
                    d_1814_v3_ = d_1814_v3_
                    if (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]:
                        index327_ = default__.safeIndex(931, ((self).f20).length(0))
                        ((self).f20)[index327_] = len(_dafny.SeqWithoutIsStrInference([d_1811_v0_ for d_1821_i1_ in range(default__.abs(533))]))
                        d_1822_v9_: _dafny.Array
                        def lambda98_(d_1823_v6_):
                            def lambda99_(d_1824_i2_):
                                return d_1823_v6_

                            return lambda99_

                        init53_ = lambda98_(d_1817_v6_)
                        nw333_ = _dafny.Array(None, 12)
                        for i0_53_ in range(nw333_.length(0)):
                            nw333_[i0_53_] = init53_(i0_53_)
                        d_1822_v9_ = nw333_
                        index328_ = default__.safeIndex(321, (d_1822_v9_).length(0))
                        (d_1822_v9_)[index328_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmofc"))
                        index329_ = default__.safeIndex(832, ((self).f20).length(0))
                        ((self).f20)[index329_] = d_1811_v0_
                        d_1825_v10_: _dafny.Seq
                        d_1825_v10_ = _dafny.SeqWithoutIsStrInference([d_1817_v6_, d_1817_v6_])
                        d_1826_v11_: _dafny.Set
                        d_1826_v11_ = _dafny.Set({self.f16, (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))], self.f16})
                        index330_ = default__.safeIndex(931, ((self).f20).length(0))
                        index331_ = default__.safeIndex(321, (d_1822_v9_).length(0))
                        index332_ = default__.safeIndex(832, ((self).f20).length(0))
                        rhs362_ = d_1811_v0_
                        rhs363_ = ((d_1825_v10_)[default__.safeIndex(d_1811_v0_, len(d_1825_v10_))]) + (d_1817_v6_)
                        rhs364_ = len(d_1826_v11_)
                        rhs365_ = (not ((d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]) or ((d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))])) and ((d_1811_v0_) >= (d_1811_v0_))
                        rhs366_ = (d_1811_v0_) - (((0) - (d_1811_v0_)) * (d_1811_v0_))
                        lhs281_ = (self).f20
                        lhs282_ = default__.safeIndex(931, ((self).f20).length(0))
                        lhs283_ = d_1822_v9_
                        lhs284_ = default__.safeIndex(321, (d_1822_v9_).length(0))
                        lhs285_ = globalState
                        lhs286_ = (self).f20
                        lhs287_ = default__.safeIndex(832, ((self).f20).length(0))
                        lhs281_[lhs282_] = rhs362_
                        lhs283_[lhs284_] = rhs363_
                        d_1811_v0_ = rhs364_
                        lhs285_.f2 = rhs365_
                        lhs286_[lhs287_] = rhs366_
                        d_1811_v0_ = (((d_1813_v2_)[((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]] if (((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]) in (d_1813_v2_) else ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]) if (True if self.f16 else (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]) else d_1811_v0_)
                        d_1827_v12_: C0
                        nw334_ = C0()
                        nw334_.ctor__(742, d_1814_v3_)
                        d_1827_v12_ = nw334_
                        d_1828_v13_: _dafny.Map
                        d_1828_v13_ = _dafny.Map({d_1827_v12_: ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]})
                        d_1829_v14_: _dafny.Array
                        nw335_ = _dafny.Array(None, 21)
                        nw335_[int(0)] = ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]
                        nw335_[int(1)] = ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]
                        nw335_[int(2)] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1830_i3_ in range(default__.abs(-983))])))
                        nw335_[int(3)] = (0) - (((self).f20)[default__.safeIndex(931, ((self).f20).length(0))])
                        nw335_[int(4)] = len(_dafny.SeqWithoutIsStrInference([(d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]]))
                        nw335_[int(5)] = (_dafny.MultiSet([d_1811_v0_])).cardinality
                        nw335_[int(6)] = ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]
                        nw335_[int(7)] = d_1811_v0_
                        nw335_[int(8)] = ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]
                        nw335_[int(9)] = 420
                        nw335_[int(10)] = len(d_1828_v13_)
                        nw335_[int(11)] = d_1811_v0_
                        nw335_[int(12)] = d_1811_v0_
                        nw335_[int(13)] = ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]
                        nw335_[int(14)] = (d_1827_v12_).f26
                        nw335_[int(15)] = d_1811_v0_
                        nw335_[int(16)] = -696
                        nw335_[int(17)] = default__.fm3(globalState)
                        nw335_[int(18)] = d_1811_v0_
                        nw335_[int(19)] = (d_1827_v12_).f26
                        nw335_[int(20)] = ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))]
                        d_1829_v14_ = nw335_
                        d_1831_v15_: _dafny.Map
                        d_1831_v15_ = _dafny.Map({True: d_1829_v14_})
                        d_1831_v15_ = (d_1831_v15_).set(self.f16, d_1829_v14_)
                        (globalState).f5 = (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]
                        d_1832_v17_: _dafny.Set
                        def iife168_():
                            coll66_ = _dafny.Set()
                            compr_66_: int
                            for compr_66_ in _dafny.IntegerRange(-397, -384):
                                d_1833_v16_: int = compr_66_
                                if ((-397) <= (d_1833_v16_)) and ((d_1833_v16_) < (-384)):
                                    coll66_ = coll66_.union(_dafny.Set([default__.safeModuloInt(d_1833_v16_, (d_1827_v12_).f26)]))
                            return _dafny.Set(coll66_)
                        d_1832_v17_ = _dafny.Set({(d_1827_v12_).f26, (d_1827_v12_).f26, len(iife168_()
                        )})
                        d_1834_v18_: D9
                        d_1834_v18_ = D9_DC26(len((d_1825_v10_)[default__.safeIndex((d_1827_v12_).f26, len(d_1825_v10_))]), d_1827_v12_)
                        index333_ = default__.safeIndex(931, ((self).f20).length(0))
                        index334_ = default__.safeIndex(847, (d_1814_v3_).length(0))
                        rhs367_ = (d_1827_v12_ if (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))] else (d_1834_v18_).cf46)
                        rhs368_ = (d_1827_v12_).f26
                        rhs369_ = (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]
                        rhs370_ = (d_1832_v17_) - (_dafny.Set({((self).f20)[default__.safeIndex(931, ((self).f20).length(0))], ((self).f20)[default__.safeIndex(931, ((self).f20).length(0))], d_1811_v0_, (d_1827_v12_).f26, d_1811_v0_}))
                        lhs288_ = (self).f20
                        lhs289_ = default__.safeIndex(931, ((self).f20).length(0))
                        lhs290_ = d_1814_v3_
                        lhs291_ = default__.safeIndex(847, (d_1814_v3_).length(0))
                        d_1827_v12_ = rhs367_
                        lhs288_[lhs289_] = rhs368_
                        lhs290_[lhs291_] = rhs369_
                        d_1832_v17_ = rhs370_
                    elif True:
                        d_1817_v6_ = (_dafny.SeqWithoutIsStrInference([default__.fm2(globalState), d_1819_v7_])) + ((d_1817_v6_) + (d_1817_v6_))
                        d_1835_v19_: _dafny.Seq
                        d_1835_v19_ = _dafny.SeqWithoutIsStrInference([d_1817_v6_])
                        d_1836_v20_: _dafny.Map
                        d_1836_v20_ = _dafny.Map({d_1835_v19_: default__.safeModuloInt(d_1811_v0_, d_1811_v0_)})
                        d_1836_v20_ = (d_1836_v20_).set(d_1835_v19_, d_1811_v0_)
                        d_1817_v6_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yerr")) if True else d_1817_v6_)
                        d_1837_v21_: _dafny.Map
                        d_1837_v21_ = _dafny.Map({((d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))] if (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))] else self.f16): (len(_dafny.Set({(d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))], (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]}))) - (37)})
                        d_1837_v21_ = (d_1837_v21_).set(not(self.f16), default__.fm3(globalState))
                        (globalState).f5 = (d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]
                    d_1838_v22_: _dafny.Set
                    d_1838_v22_ = _dafny.Set({d_1811_v0_, d_1811_v0_, (0) - (d_1811_v0_)})
                    d_1839_v23_: D9
                    d_1839_v23_ = D9_DC27((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxw"))) <= (d_1817_v6_), d_1838_v22_)
                    d_1839_v23_ = d_1839_v23_
                    pass
            pass
        d_1840_v24_: C2
        nw336_ = C2()
        nw336_.ctor__((0) - (default__.safeDivisionInt(d_1811_v0_, d_1811_v0_)))
        d_1840_v24_ = nw336_
        d_1814_v3_ = d_1814_v3_
        hi17_ = d_1811_v0_
        for d_1841_i4_ in range(676, hi17_):
            d_1811_v0_ = default__.safeModuloInt(d_1841_i4_, len(_dafny.Map({(d_1840_v24_).f25: d_1841_i4_})))
            d_1842_v25_: _dafny.MultiSet
            d_1842_v25_ = _dafny.MultiSet([self.f16])
            d_1843_v26_: _dafny.Seq
            d_1843_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_1814_v3_)[default__.safeIndex(847, (d_1814_v3_).length(0))]]), d_1842_v25_])
            d_1842_v25_ = (d_1843_v26_)[default__.safeIndex((d_1840_v24_).f25, len(d_1843_v26_))]
            d_1844_v27_: str
            d_1844_v27_ = _dafny.CodePoint('l')
            d_1845_v28_: _dafny.Set
            d_1845_v28_ = _dafny.Set({d_1844_v27_})
            d_1845_v28_ = default__.fm27(globalState)
            d_1846_v29_: C6
            nw337_ = C6()
            nw337_.ctor__(d_1816_v5_)
            d_1846_v29_ = nw337_
        d_1811_v0_ = len(((self).f17) + ((_dafny.SeqWithoutIsStrInference([109 for d_1847_i5_ in range(default__.abs(-893))])) + ((self).f17)))

    @property
    def f20(self):
        return self._f20

class C8(T0):
    def  __init__(self):
        self._f16: bool = False
        self._f17: _dafny.Seq = _dafny.Seq({})
        self.f19: _dafny.Set = _dafny.Set({})
        self._f18: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f18, f19, f16, f17):
        (self)._f18 = f18
        (self).f19 = f19
        (self).f16 = f16
        (self)._f17 = f17

    def fm5(self, globalState):
        return self.f16

    def m1(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        hi18_ = p3
        for d_1848_i0_ in range(p3, hi18_):
            d_1849_v0_: _dafny.Map
            d_1849_v0_ = _dafny.Map({p3: p2})
            d_1849_v0_ = (d_1849_v0_).set(d_1848_i0_, not(p1))
            (globalState).f2 = p2
            d_1850_v1_: int
            d_1850_v1_ = 960
            d_1850_v1_ = d_1848_i0_
            d_1851_v2_: _dafny.Array
            nw338_ = _dafny.Array(int(0), 24)
            d_1851_v2_ = nw338_
            d_1852_v3_: _dafny.Map
            d_1852_v3_ = _dafny.Map({self.f16: d_1851_v2_})
            d_1852_v3_ = (d_1852_v3_).set(self.f16, d_1851_v2_)
        hi19_ = p3
        for d_1853_i1_ in range(p3, hi19_):
            d_1854_v4_: int
            d_1854_v4_ = 284
            d_1855_v5_: _dafny.Array
            nw339_ = _dafny.Array(int(0), 26)
            d_1855_v5_ = nw339_
            d_1856_v6_: _dafny.Array
            def lambda100_(d_1857_p2_):
                def lambda101_(d_1858_i2_):
                    return d_1857_p2_

                return lambda101_

            init54_ = lambda100_(p2)
            nw340_ = _dafny.Array(None, 20)
            for i0_54_ in range(nw340_.length(0)):
                nw340_[i0_54_] = init54_(i0_54_)
            d_1856_v6_ = nw340_
            index335_ = default__.safeIndex(850, (d_1856_v6_).length(0))
            (d_1856_v6_)[index335_] = (_dafny.Set({p1, p2})).isdisjoint(self.f19)
            d_1859_v7_: D0
            d_1859_v7_ = D0_DC0(default__.fm1(globalState))
            d_1860_v8_: _dafny.Map
            d_1860_v8_ = _dafny.Map({p2: d_1859_v7_})
            index336_ = default__.safeIndex(850, (d_1856_v6_).length(0))
            rhs371_ = len(d_1860_v8_)
            rhs372_ = d_1855_v5_
            rhs373_ = not(p2)
            lhs292_ = d_1856_v6_
            lhs293_ = default__.safeIndex(850, (d_1856_v6_).length(0))
            d_1854_v4_ = rhs371_
            d_1855_v5_ = rhs372_
            lhs292_[lhs293_] = rhs373_
            d_1854_v4_ = -278
            d_1861_v9_: C6
            nw341_ = C6()
            nw341_.ctor__(D0_DC2(((self).f17)[default__.safeIndex(p3, len((self).f17))], p3, d_1854_v4_, d_1856_v6_))
            d_1861_v9_ = nw341_
            d_1862_v10_: _dafny.Seq
            d_1862_v10_ = _dafny.SeqWithoutIsStrInference([d_1856_v6_])
            (globalState).f5 = (d_1856_v6_) in (d_1862_v10_)
        d_1863_v11_: _dafny.Array
        nw342_ = _dafny.Array(int(0), 2)
        d_1863_v11_ = nw342_
        d_1864_v12_: _dafny.Map
        d_1864_v12_ = _dafny.Map({p1: d_1863_v11_})
        hi20_ = (0) - (len((d_1864_v12_).set(self.f16, d_1863_v11_)))
        for d_1865_i3_ in range(p3, hi20_):
            d_1866_v13_: _dafny.Seq
            d_1866_v13_ = _dafny.SeqWithoutIsStrInference([p3])
            d_1866_v13_ = (self).f17
            (globalState).f5 = True
            d_1867_v14_: T1
            nw343_ = C1()
            nw343_.ctor__(len(self.f19), p3, p1, (self).f17, p3)
            d_1867_v14_ = nw343_
            d_1868_v15_: _dafny.Map
            d_1868_v15_ = _dafny.Map({len(self.f19): d_1867_v14_})
            d_1868_v15_ = (d_1868_v15_).set(p3, d_1867_v14_)
            if self.f16:
                d_1869_v16_: D3
                d_1869_v16_ = D3_DC10(True)
                d_1870_v17_: _dafny.Seq
                d_1870_v17_ = _dafny.SeqWithoutIsStrInference([d_1867_v14_.f16, d_1867_v14_.f16, (d_1869_v16_).cf19])
                r0 = (d_1870_v17_).set(default__.safeIndex((p3) - (835), len(d_1870_v17_)), p2)
                d_1871_v18_: _dafny.Array
                nw344_ = _dafny.Array(False, 20)
                d_1871_v18_ = nw344_
                d_1872_v19_: str
                d_1872_v19_ = _dafny.CodePoint('q')
                d_1873_v20_: _dafny.MultiSet
                d_1873_v20_ = _dafny.MultiSet([d_1872_v19_])
                index337_ = default__.safeIndex(176, (d_1871_v18_).length(0))
                (d_1871_v18_)[index337_] = (_dafny.MultiSet([d_1872_v19_])) == (d_1873_v20_)
                d_1874_v21_: _dafny.Seq
                d_1874_v21_ = _dafny.SeqWithoutIsStrInference([(p0).set(default__.safeIndex((d_1867_v14_).f28, len(p0)), d_1872_v19_), (self).f18, p0])
                d_1875_v22_: D0
                d_1875_v22_ = D0_DC1(p0)
                d_1876_v23_: _dafny.Set
                d_1876_v23_ = _dafny.Set({(self).f18, p0, (d_1874_v21_)[default__.safeIndex((0) - (len(_dafny.Map({self.f16: len(_dafny.Set({(d_1867_v14_).f28}))}))), len(d_1874_v21_))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apnbo")), (((self).f18) + ((d_1875_v22_).cf1)).set(default__.safeIndex((d_1867_v14_).f28, len(((self).f18) + ((d_1875_v22_).cf1))), d_1872_v19_)})
                d_1877_v24_: int
                d_1877_v24_ = 320
                index338_ = default__.safeIndex(176, (d_1871_v18_).length(0))
                rhs374_ = d_1867_v14_.f16
                rhs375_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktslxunay"))) < ((self).f18)
                rhs376_ = d_1876_v23_
                rhs377_ = (d_1870_v17_) + (d_1870_v17_)
                rhs378_ = d_1865_i3_
                lhs294_ = d_1871_v18_
                lhs295_ = default__.safeIndex(176, (d_1871_v18_).length(0))
                lhs296_ = globalState
                lhs294_[lhs295_] = rhs374_
                lhs296_.f2 = rhs375_
                d_1876_v23_ = rhs376_
                d_1870_v17_ = rhs377_
                d_1877_v24_ = rhs378_
                d_1866_v13_ = (d_1867_v14_).f17
                d_1878_v25_: _dafny.Seq
                d_1878_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jco"))
                d_1879_v26_: _dafny.Set
                d_1879_v26_ = _dafny.Set({d_1863_v11_})
                d_1880_v27_: _dafny.MultiSet
                d_1880_v27_ = _dafny.MultiSet([p2])
                d_1881_v28_: _dafny.MultiSet
                d_1881_v28_ = _dafny.MultiSet([len((d_1867_v14_).f17), (d_1880_v27_).cardinality, d_1865_i3_])
                rhs379_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfcilcoa"))
                rhs380_ = d_1879_v26_
                rhs381_ = ((_dafny.MultiSet([d_1877_v24_])).set(-137, default__.abs((0) - (p3)))) == (d_1881_v28_)
                lhs297_ = globalState
                d_1878_v25_ = rhs379_
                d_1879_v26_ = rhs380_
                lhs297_.f2 = rhs381_
                d_1882_v29_: _dafny.Map
                d_1882_v29_ = _dafny.Map({p2: (self).fm5(globalState)})
                d_1882_v29_ = (d_1882_v29_).set(p2, True)
            elif True:
                d_1883_v30_: str
                d_1883_v30_ = _dafny.CodePoint('b')
                (globalState).f15 = d_1883_v30_
                d_1884_v31_: int
                d_1884_v31_ = 418
                rhs382_ = (d_1867_v14_).f28
                rhs383_ = default__.fm3(globalState)
                d_1884_v31_ = rhs382_
                d_1884_v31_ = rhs383_
                d_1885_v32_: _dafny.Array
                nw345_ = _dafny.Array(False, 13)
                d_1885_v32_ = nw345_
                index339_ = default__.safeIndex(252, (d_1885_v32_).length(0))
                (d_1885_v32_)[index339_] = ((d_1867_v14_).f28) < ((d_1867_v14_).f28)
                index340_ = default__.safeIndex(252, (d_1885_v32_).length(0))
                (d_1885_v32_)[index340_] = p2
                rhs384_ = (len((self).f18)) - ((p3) * (d_1865_i3_))
                rhs385_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvfjtwdb"))), ((d_1867_v14_).f28) + (d_1865_i3_))
                rhs386_ = (not(d_1867_v14_.f16)) == ((p1) and (not((d_1885_v32_)[default__.safeIndex(252, (d_1885_v32_).length(0))])))
                rhs387_ = (0) - (d_1884_v31_)
                lhs298_ = globalState
                d_1884_v31_ = rhs384_
                d_1884_v31_ = rhs385_
                lhs298_.f2 = rhs386_
                d_1884_v31_ = rhs387_
                d_1886_v33_: C3
                nw346_ = C3()
                nw346_.ctor__((self).f18, d_1867_v14_.f16, (d_1867_v14_).f17)
                d_1886_v33_ = nw346_
        if self.f16:
            d_1887_v34_: _dafny.Map
            d_1887_v34_ = _dafny.Map({d_1863_v11_: p0})
            d_1888_v35_: _dafny.Map
            d_1888_v35_ = _dafny.Map({p2: p0})
            d_1889_v36_: _dafny.Map
            d_1889_v36_ = _dafny.Map({self.f16: ((d_1888_v35_)[p1] if (p1) in (d_1888_v35_) else p0)})
            d_1887_v34_ = (d_1887_v34_).set(d_1863_v11_, ((d_1888_v35_)[p2] if (p2) in (d_1888_v35_) else p0))
            d_1890_v37_: _dafny.Seq
            d_1890_v37_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_1891_v38_: _dafny.Set
            d_1891_v38_ = _dafny.Set({len(d_1890_v37_), (0) - (p3), -403})
            d_1891_v38_ = (d_1891_v38_) | (default__.fm31(p2, globalState))
            d_1892_v39_: _dafny.Array
            def lambda102_(d_1893_i4_):
                return _dafny.CodePoint('q')

            init55_ = lambda102_
            nw347_ = _dafny.Array(None, 20)
            for i0_55_ in range(nw347_.length(0)):
                nw347_[i0_55_] = init55_(i0_55_)
            d_1892_v39_ = nw347_
            d_1894_v40_: _dafny.Seq
            d_1894_v40_ = _dafny.SeqWithoutIsStrInference([d_1892_v39_, d_1892_v39_, d_1892_v39_, d_1892_v39_])
            d_1894_v40_ = (d_1894_v40_) + (d_1894_v40_)
            d_1895_v41_: _dafny.Array
            nw348_ = _dafny.Array(False, 5)
            d_1895_v41_ = nw348_
            d_1896_v42_: D11
            d_1896_v42_ = D11_DC31(p1, 919, p3, p1)
            d_1897_v43_: _dafny.Map
            d_1897_v43_ = _dafny.Map({D11_DC32(d_1896_v42_): d_1895_v41_})
            d_1898_v44_: D11
            d_1898_v44_ = D11_DC32(d_1896_v42_)
            d_1899_v45_: _dafny.MultiSet
            d_1899_v45_ = _dafny.MultiSet([762])
            d_1900_v46_: D0
            d_1900_v46_ = D0_DC2((d_1899_v45_).cardinality, p3, p3, d_1895_v41_)
            rhs388_ = ((d_1897_v43_)[d_1898_v44_] if (d_1898_v44_) in (d_1897_v43_) else (d_1900_v46_).cf5)
            rhs389_ = d_1863_v11_
            d_1895_v41_ = rhs388_
            d_1863_v11_ = rhs389_
            d_1901_v47_: int
            d_1901_v47_ = -83
            d_1902_v48_: D11
            d_1902_v48_ = D11_DC29(default__.fm2(globalState))
            pat_let_tv39_ = d_1896_v42_
            def iife169_(_pat_let51_0):
                def iife170_(d_1903_dt__update__tmp_h0_):
                    def iife171_(_pat_let52_0):
                        def iife172_(d_1904_dt__update_hcf57_h0_):
                            return D11_DC32(d_1904_dt__update_hcf57_h0_)
                        return iife172_(_pat_let52_0)
                    return iife171_(pat_let_tv39_)
                return iife170_(_pat_let51_0)
            rhs390_ = d_1901_v47_
            rhs391_ = d_1902_v48_
            rhs392_ = d_1901_v47_
            rhs393_ = iife169_(d_1898_v44_)
            rhs394_ = True
            lhs299_ = self
            d_1901_v47_ = rhs390_
            d_1902_v48_ = rhs391_
            d_1901_v47_ = rhs392_
            d_1898_v44_ = rhs393_
            lhs299_.f16 = rhs394_
        elif True:
            if p2:
                d_1905_v49_: _dafny.Seq
                d_1905_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebxvgp"))
                d_1906_v50_: _dafny.Map
                d_1906_v50_ = _dafny.Map({p3: d_1905_v49_})
                d_1907_v51_: _dafny.Seq
                d_1907_v51_ = _dafny.SeqWithoutIsStrInference([d_1905_v49_])
                d_1908_v52_: _dafny.MultiSet
                d_1908_v52_ = _dafny.MultiSet([p1])
                d_1905_v49_ = ((d_1906_v50_)[default__.safeModuloInt(p3, default__.fm3(globalState))] if (default__.safeModuloInt(p3, default__.fm3(globalState))) in (d_1906_v50_) else ((d_1907_v51_)[default__.safeIndex((d_1908_v52_).cardinality, len(d_1907_v51_))] if self.f16 else (d_1907_v51_)[default__.safeIndex(p3, len(d_1907_v51_))]))
                d_1909_v53_: _dafny.Array
                nw349_ = _dafny.Array(False, 11)
                d_1909_v53_ = nw349_
                index341_ = default__.safeIndex(819, (d_1909_v53_).length(0))
                (d_1909_v53_)[index341_] = (False) and (p2)
                d_1910_v54_: _dafny.Map
                d_1910_v54_ = _dafny.Map({self.f16: p1})
                d_1911_v55_: D1
                d_1911_v55_ = D1_DC5(len((d_1910_v54_).set(p2, p1)), p2)
                d_1912_v56_: _dafny.Seq
                d_1912_v56_ = _dafny.SeqWithoutIsStrInference([self.f16, default__.fm0(self.f16, self.f16, globalState), self.f16, p2])
                d_1913_v57_: C5
                nw350_ = C5()
                nw350_.ctor__(self.f16, (self).f17)
                d_1913_v57_ = nw350_
                d_1914_v58_: _dafny.Map
                d_1914_v58_ = _dafny.Map({d_1913_v57_: p3})
                index342_ = default__.safeIndex(819, (d_1909_v53_).length(0))
                (d_1909_v53_)[index342_] = ((d_1912_v56_)[default__.safeIndex(((d_1914_v58_)[d_1913_v57_] if (d_1913_v57_) in (d_1914_v58_) else p3), len(d_1912_v56_))] if (d_1911_v55_).cf11 else p1)
                d_1915_v59_: _dafny.Map
                d_1915_v59_ = _dafny.Map({p1: len((self).f18)})
                d_1916_v60_: _dafny.Array
                nw351_ = _dafny.Array(None, 20)
                nw351_[int(0)] = _dafny.SeqWithoutIsStrInference([p3, (0) - (p3), p3, p3, p3])
                nw351_[int(1)] = _dafny.SeqWithoutIsStrInference([-740])
                nw351_[int(2)] = (self).f17
                nw351_[int(3)] = (self).f17
                nw351_[int(4)] = (self).f17
                nw351_[int(5)] = (self).f17
                nw351_[int(6)] = (self).f17
                nw351_[int(7)] = (self).f17
                nw351_[int(8)] = (self).f17
                nw351_[int(9)] = _dafny.SeqWithoutIsStrInference([p3, p3])
                nw351_[int(10)] = (self).f17
                nw351_[int(11)] = (self).f17
                nw351_[int(12)] = (self).f17
                nw351_[int(13)] = _dafny.SeqWithoutIsStrInference([p3])
                nw351_[int(14)] = (self).f17
                nw351_[int(15)] = ((self).f17).set(default__.safeIndex(474, len((self).f17)), p3)
                nw351_[int(16)] = (self).f17
                nw351_[int(17)] = (self).f17
                nw351_[int(18)] = (self).f17
                nw351_[int(19)] = (self).f17
                d_1916_v60_ = nw351_
                d_1917_v61_: D2
                d_1917_v61_ = D2_DC8(d_1915_v59_, d_1916_v60_, p1, self.f16, p3)
                d_1918_v62_: D6
                d_1918_v62_ = D6_DC17(d_1917_v61_)
                d_1918_v62_ = d_1918_v62_
                d_1919_v63_: _dafny.Set
                d_1919_v63_ = _dafny.Set({p3})
                d_1920_v64_: _dafny.Map
                d_1920_v64_ = _dafny.Map({d_1919_v63_: d_1863_v11_})
                def iife173_():
                    coll67_ = _dafny.Set()
                    compr_67_: int
                    for compr_67_ in (d_1919_v63_).Elements:
                        d_1921_v65_: int = compr_67_
                        if (d_1921_v65_) in (d_1919_v63_):
                            coll67_ = coll67_.union(_dafny.Set([default__.safeDivisionInt(d_1921_v65_, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([_dafny.Map({439: not(True)})])).cardinality: False})])))]))
                    return _dafny.Set(coll67_)
                d_1920_v64_ = (d_1920_v64_).set(iife173_()
                , d_1863_v11_)
                d_1922_v66_: _dafny.Array
                nw352_ = _dafny.Array(None, 1)
                nw352_[int(0)] = self.f19
                d_1922_v66_ = nw352_
                d_1922_v66_ = d_1922_v66_
            elif True:
                (globalState).f2 = (p0) < ((self).f18)
                d_1923_v67_: D11
                d_1923_v67_ = D11_DC31(self.f16, p3, p3, True)
                d_1923_v67_ = D11_DC31(self.f16, p3, p3, self.f16)
                d_1924_v68_: int
                d_1924_v68_ = 248
                d_1924_v68_ = d_1924_v68_
                d_1924_v68_ = (default__.fm3(globalState) if not(self.f16) else 6)
                (globalState).f15 = ((self).f18)[default__.safeIndex(d_1924_v68_, len((self).f18))]
            d_1925_v69_: _dafny.Array
            def lambda103_(d_1926_i5_):
                return False

            init56_ = lambda103_
            nw353_ = _dafny.Array(None, 15)
            for i0_56_ in range(nw353_.length(0)):
                nw353_[i0_56_] = init56_(i0_56_)
            d_1925_v69_ = nw353_
            index343_ = default__.safeIndex(152, (d_1925_v69_).length(0))
            (d_1925_v69_)[index343_] = default__.fm0(p2, p2, globalState)
            index344_ = default__.safeIndex(152, (d_1925_v69_).length(0))
            (d_1925_v69_)[index344_] = p2
            d_1927_v70_: int
            d_1927_v70_ = -337
            rhs395_ = ((0) - (p3)) * (d_1927_v70_)
            rhs396_ = default__.safeDivisionInt(15, p3)
            d_1927_v70_ = rhs395_
            d_1927_v70_ = rhs396_
            d_1928_v71_: _dafny.Seq
            d_1928_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlvn"))
            d_1928_v71_ = ((self).f18) + ((self).f18)
            d_1927_v70_ = (len((self).f17)) + (p3)
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_1863_v11_).length(0)):
            d_1929_i6_: int = guard_loop_9_
            if (True) and (((0) <= (d_1929_i6_)) and ((d_1929_i6_) < ((d_1863_v11_).length(0)))):
                (d_1863_v11_)[(d_1929_i6_)] = (d_1929_i6_) * (p3)
        d_1930_v72_: int
        d_1930_v72_ = -430
        d_1931_v73_: _dafny.Seq
        d_1931_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmsoc"))
        rhs397_ = (d_1863_v11_ if self.f16 else d_1863_v11_)
        rhs398_ = len((default__.fm20(self.f16, not(p2), True, d_1930_v72_, globalState)) + (_dafny.SeqWithoutIsStrInference([p1, self.f16])))
        rhs399_ = p2
        rhs400_ = not(self.f16)
        rhs401_ = p0
        lhs300_ = globalState
        lhs301_ = globalState
        d_1863_v11_ = rhs397_
        d_1930_v72_ = rhs398_
        lhs300_.f5 = rhs399_
        lhs301_.f5 = rhs400_
        d_1931_v73_ = rhs401_
        d_1932_v74_: _dafny.Seq
        d_1932_v74_ = _dafny.SeqWithoutIsStrInference([p2, (p0) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))), p1])
        r0 = d_1932_v74_
        return r0

    @property
    def f18(self):
        return self._f18
