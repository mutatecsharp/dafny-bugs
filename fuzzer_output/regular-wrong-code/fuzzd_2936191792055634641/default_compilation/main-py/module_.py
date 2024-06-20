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
        return _dafny.CodePoint('n')

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return (not((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True, True])])).isdisjoint(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True])])))) == (True)

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        source0_ = D27_DC66(_dafny.SeqWithoutIsStrInference([False]), len(_dafny.Map({True: False})))
        if source0_.is_DC65:
            d_0___mcc_h0_ = source0_.cf107
            d_1___mcc_h1_ = source0_.cf108
            d_2___mcc_h2_ = source0_.cf109
            d_3___mcc_h3_ = source0_.cf110
            d_4_cf110_ = d_3___mcc_h3_
            d_5_cf109_ = d_2___mcc_h2_
            d_6_cf108_ = d_1___mcc_h1_
            d_7_cf107_ = d_0___mcc_h0_
            return _dafny.Map({d_7_cf107_: d_4_cf110_})
        elif source0_.is_DC66:
            d_8___mcc_h4_ = source0_.cf111
            d_9___mcc_h5_ = source0_.cf112
            d_10_cf112_ = d_9___mcc_h5_
            d_11_cf111_ = d_8___mcc_h4_
            if True:
                return _dafny.Map({True: False})
            elif True:
                return _dafny.Map({True: True})
        elif source0_.is_DC67:
            d_12___mcc_h6_ = source0_.cf113
            d_13___mcc_h7_ = source0_.cf114
            d_14___mcc_h8_ = source0_.cf115
            d_15_cf115_ = d_14___mcc_h8_
            d_16_cf114_ = d_13___mcc_h7_
            d_17_cf113_ = d_12___mcc_h6_
            return (_dafny.Map({True: not(False)})) | (_dafny.Map({True: True}))
        elif True:
            d_18___mcc_h9_ = source0_.cf106
            d_19_cf106_ = d_18___mcc_h9_
            return (_dafny.Map({not(False): True})) | (_dafny.Map({not(True): False}))

    @staticmethod
    def fm11(globalState):
        source1_ = D5_DC12()
        if source1_.is_DC12:
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: D8
                for compr_0_ in (_dafny.Set({D8_DC18(_dafny.MultiSet([-307, len(_dafny.SeqWithoutIsStrInference([False]))]))})).Elements:
                    d_20_v0_: D8 = compr_0_
                    if (d_20_v0_) in (_dafny.Set({D8_DC18(_dafny.MultiSet([-307, len(_dafny.SeqWithoutIsStrInference([False]))]))})):
                        coll0_[d_20_v0_] = (0) - (len(_dafny.Map({not(not(True)): len(_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({True: 558}))])).cardinality: 156}))})))
                return _dafny.Map(coll0_)
            return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('a')]))])) | (_dafny.MultiSet([len(iife0_()
            )]))
        elif source1_.is_DC13:
            d_21___mcc_h0_ = source1_.cf16
            d_22___mcc_h1_ = source1_.cf17
            d_23___mcc_h2_ = source1_.cf18
            d_24_cf18_ = d_23___mcc_h2_
            d_25_cf17_ = d_22___mcc_h1_
            d_26_cf16_ = d_21___mcc_h0_
            return _dafny.MultiSet([646, d_25_cf17_])
        elif True:
            d_27___mcc_h3_ = source1_.cf15
            d_28_cf15_ = d_27___mcc_h3_
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([688]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), (0) - (len((D5_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbfo")), 514, False)).cf16)), 832])))

    @staticmethod
    def fm12(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bt"))

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return 369

    @staticmethod
    def fm14(globalState):
        return (((_dafny.MultiSet([True, False, True])) | (_dafny.MultiSet([True]))).cardinality) + (len(_dafny.Set({164, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.MultiSet([891]): not(False)}) for d_29_i0_ in range(default__.abs(572))]))})))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        if (True if True else False):
            return (_dafny.MultiSet([True])) - (_dafny.MultiSet([True]))
        elif True:
            return _dafny.MultiSet([True])

    @staticmethod
    def fm16(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxrlr"))

    @staticmethod
    def fm19(p0, p1, p2, globalState):
        return (_dafny.Map({not(False): not(True)})) | (_dafny.Map({True: False}))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False, True, True, False, not(True)])))])

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.MultiSet
            for compr_1_ in (_dafny.Set({_dafny.MultiSet([(_dafny.MultiSet([267])).cardinality, 946, -175])})).Elements:
                d_30_v0_: _dafny.MultiSet = compr_1_
                if (d_30_v0_) in (_dafny.Set({_dafny.MultiSet([(_dafny.MultiSet([267])).cardinality, 946, -175])})):
                    coll1_[d_30_v0_] = True
            return _dafny.Map(coll1_)
        source2_ = iife1_()

        d_31___mcc_h0_ = source2_
        d_32_cf94_ = d_31___mcc_h0_
        return (_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        if not(not (not(True)) or (not(False))):
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(-471, 969):
                    d_33_v0_: int = compr_2_
                    if ((-471) <= (d_33_v0_)) and ((d_33_v0_) < (969)):
                        coll2_ = coll2_.union(_dafny.Set([default__.safeDivisionInt(d_33_v0_, len(_dafny.SeqWithoutIsStrInference([917 for d_34_i0_ in range(default__.abs(445))])))]))
                return _dafny.Set(coll2_)
            return (iife2_()
            ) | (_dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_35_i1_ in range(default__.abs(900))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_36_i2_ in range(default__.abs(466))])), len(_dafny.Map({True: False})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([576]))}))}))}))
        elif True:
            return _dafny.Set({168, -164})

    @staticmethod
    def fm23(p0, globalState):
        return _dafny.MultiSet([(_dafny.CodePoint('x') if True else _dafny.CodePoint('o'))])

    @staticmethod
    def fm24(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtft"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_37_i0_ in range(default__.abs(-262))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpf"))))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(611, 962):
                d_38_v0_: int = compr_3_
                if ((611) <= (d_38_v0_)) and ((d_38_v0_) < (962)):
                    coll3_[(d_38_v0_) * (283)] = 331
            return _dafny.Map(coll3_)
        return (_dafny.Map({-921: _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlrlyjs")))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odsdhie")))})})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxobxhhst"))): iife3_()
        })) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhycailwn"))): _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nimexgr"))): len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hosqaqqwm"))): True})): 964}))})})))

    @staticmethod
    def fm26(p0, p1, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({True, not(True)}))) | ((_dafny.Set({True, False, not(not(not(True)))}) if False else _dafny.Set({False})))

    @staticmethod
    def fm27(globalState):
        return D13_DC29((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iusforw"))), False)

    @staticmethod
    def fm28(globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: D13
            for compr_4_ in (_dafny.Map({D13_DC29(False, True): True})).keys.Elements:
                d_39_v0_: D13 = compr_4_
                if (d_39_v0_) in (_dafny.Map({D13_DC29(False, True): True})):
                    coll4_[d_39_v0_] = True
            return _dafny.Map(coll4_)
        return ((_dafny.MultiSet([iife4_()
        ])) - (_dafny.MultiSet([_dafny.Map({D13_DC29(False, False): False}), _dafny.Map({D13_DC29(True, False): True}), _dafny.Map({D13_DC29(False, True): True})]))) | (_dafny.MultiSet([_dafny.Map({D13_DC29(True, False): False})]))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return (_dafny.Map({not(not(not(not(True)))): 464})) | ((_dafny.Map({False: (_dafny.MultiSet([False])).cardinality})) | (_dafny.Map({False: (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({221: _dafny.SeqWithoutIsStrInference([True, False])})) for d_40_i0_ in range(default__.abs(856))]))).cardinality)})))

    @staticmethod
    def fm30(p0, p1, globalState):
        return ((_dafny.MultiSet([228])) - (_dafny.MultiSet([len(_dafny.Map({650: _dafny.Map({True: 169})})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qshnukah")))]))) - (_dafny.MultiSet([len(_dafny.Map({not((D23_DC56(False, True, True, False)).cf92): False})), (0) - ((_dafny.MultiSet([116])).cardinality), len(_dafny.Map({True: -65})), -43, 249]))

    @staticmethod
    def fm31(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([-879, 774, -161, -747])) + (_dafny.SeqWithoutIsStrInference([-625]))

    @staticmethod
    def fm32(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([917, 878, 489, 542])).Elements:
                d_41_v0_: int = compr_5_
                if (d_41_v0_) in (_dafny.SeqWithoutIsStrInference([917, 878, 489, 542])):
                    coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_41_v0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))])))]))
            return _dafny.Set(coll5_)
        return D7_DC17(not(not (True) or (True)), (iife5_()
).isdisjoint(_dafny.Set({len(_dafny.Map({False: (_dafny.MultiSet([839])).cardinality}))})), True, True)

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return ((D31_DC74(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n'), _dafny.CodePoint('x')]): _dafny.CodePoint('g')}))).cf126) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')]): _dafny.CodePoint('h')}))

    @staticmethod
    def fm34(p0, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkhgoyrc"))): True})) | (_dafny.Map({len(_dafny.Map({True: 669})): True}))) | ((_dafny.Map({102: True})) | (_dafny.Map({725: True})))

    @staticmethod
    def fm35(globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: _dafny.Map
            for compr_6_ in (_dafny.MultiSet([_dafny.Map({not(False): False}), _dafny.Map({False: True})])).Elements:
                d_42_v0_: _dafny.Map = compr_6_
                if (d_42_v0_) in (_dafny.MultiSet([_dafny.Map({not(False): False}), _dafny.Map({False: True})])):
                    coll6_[d_42_v0_] = -213
            return _dafny.Map(coll6_)
        return ((_dafny.Map({_dafny.Map({True: True}): 18})) | (_dafny.Map({_dafny.Map({False: False}): len(_dafny.SeqWithoutIsStrInference([True, False]))}))) | ((_dafny.Map({_dafny.Map({False: False}): 41})) | (iife6_()
        ))

    @staticmethod
    def fm36(p0, p1, globalState):
        return D10_DC21((_dafny.Set({False})) | (_dafny.Set({not(False)})))

    @staticmethod
    def fm37(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({False}), _dafny.Set({True, False, False, False, not(False)}), _dafny.Set({False, True}), _dafny.Set({not(not(True))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(True), False, False}), _dafny.Set({True})]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({False, not(False), True}) for d_43_i0_ in range(default__.abs(-284))]))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        if True:
            return D5_DC11(_dafny.Map({not(True): _dafny.CodePoint('c')}))
        elif True:
            return D5_DC11(_dafny.Map({not(True): _dafny.CodePoint('v')}))

    @staticmethod
    def fm39(p0, globalState):
        return _dafny.Set({_dafny.Set({False})})

    @staticmethod
    def fm40(p0, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): -975})) | ((_dafny.Map({(D20_DC48(False, _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(True), False]), len(_dafny.Map({False: not(True)})))).cf75: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_44_i1_ in range(default__.abs(302))]))) for d_45_i0_ in range(default__.abs(131))])): _dafny.CodePoint('v')}))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfpqhopax")))})))

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        return D16_DC38((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbmyxmjv"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_46_i0_ in range(default__.abs(853))])), True, (len(_dafny.Set({438, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qra")))})) if not(False) else len(_dafny.SeqWithoutIsStrInference([False]))), True)

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        source3_ = D26_DC63(False, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_47_i0_ in range(default__.abs(-795))]))), len(_dafny.Map({True: 863})), not(False))
        if source3_.is_DC62:
            d_48___mcc_h0_ = source3_.cf100
            d_49___mcc_h1_ = source3_.cf101
            d_50_cf101_ = d_49___mcc_h1_
            d_51_cf100_ = d_48___mcc_h0_
            return True
        elif source3_.is_DC63:
            d_52___mcc_h2_ = source3_.cf102
            d_53___mcc_h3_ = source3_.cf103
            d_54___mcc_h4_ = source3_.cf104
            d_55___mcc_h5_ = source3_.cf105
            d_56_cf105_ = d_55___mcc_h5_
            d_57_cf104_ = d_54___mcc_h4_
            d_58_cf103_ = d_53___mcc_h3_
            d_59_cf102_ = d_52___mcc_h2_
            return False
        elif True:
            d_60___mcc_h6_ = source3_.cf99
            d_61_cf99_ = d_60___mcc_h6_
            return True

    @staticmethod
    def fm43(p0, p1, globalState):
        return _dafny.Set({_dafny.MultiSet([True, True]), _dafny.MultiSet([True, True])})

    @staticmethod
    def fm44(globalState):
        return _dafny.MultiSet([D16_DC39(D16_DC37(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})))])

    @staticmethod
    def fm45(p0, globalState):
        return D25_DC60(True, not(False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdydk")))

    @staticmethod
    def fm46(globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhr"))])).intersection(((D32_DC77(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiba")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orlal"))]))).cf129) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgoxp"))])))

    @staticmethod
    def fm47(p0, p1, globalState):
        return _dafny.Map({_dafny.Map({not(False): False}): (len(_dafny.SeqWithoutIsStrInference([-990 for d_62_i0_ in range(default__.abs(88))]))) + ((_dafny.MultiSet([824])).cardinality)})

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: _dafny.CodePoint('v')})) | (_dafny.Map({True: _dafny.CodePoint('e')}))

    @staticmethod
    def fm49(globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(23, -251):
                d_63_v0_: int = compr_7_
                if ((23) <= (d_63_v0_)) and ((d_63_v0_) < (-251)):
                    coll7_ = coll7_.union(_dafny.Set([default__.safeDivisionInt(d_63_v0_, -694)]))
            return _dafny.Set(coll7_)
        return D20_DC48((4) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djfgxi")))), (_dafny.SeqWithoutIsStrInference([not(False), not(True)])) + (_dafny.SeqWithoutIsStrInference([not(False), False])), _dafny.SeqWithoutIsStrInference([False]), len(iife7_()
))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference([True])})

    @staticmethod
    def fm51(p0, globalState):
        return (_dafny.Set({_dafny.Map({False: not(True)})})).intersection((_dafny.Set({_dafny.Map({True: not(not(True))})})).intersection(_dafny.Set({_dafny.Map({False: not(True)}), _dafny.Map({False: not(False)})})))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        def iife8_():
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(-986, 207):
                    d_65_v1_: int = compr_10_
                    if ((-986) <= (d_65_v1_)) and ((d_65_v1_) < (207)):
                        coll10_[(d_65_v1_) + (183)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qchgrep"))
                return _dafny.Map(coll10_)
            coll8_ = _dafny.Map()
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(-986, 207):
                    d_65_v1_: int = compr_9_
                    if ((-986) <= (d_65_v1_)) and ((d_65_v1_) < (207)):
                        coll9_[(d_65_v1_) + (183)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qchgrep"))
                return _dafny.Map(coll9_)
            compr_8_: D8
            for compr_8_ in (_dafny.Map({D8_DC19(len(iife9_()
), True): (0) - (-893)})).keys.Elements:
                d_66_v0_: D8 = compr_8_
                if (d_66_v0_) in (_dafny.Map({D8_DC19(len(iife10_()
), True): (0) - (-893)})):
                    coll8_[d_66_v0_] = -258
            return _dafny.Map(coll8_)
        return ((_dafny.Map({D8_DC19((_dafny.MultiSet([-113, 655])).cardinality, True): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyprqgfxf"))): len(_dafny.Set({True}))}))})) | (_dafny.Map({D8_DC19(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_64_i0_ in range(default__.abs(-318))])), False): 411}))) | ((iife8_()
        ) | (_dafny.Map({D8_DC19(993, False): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtdlx")))]))})))

    @staticmethod
    def fm53(globalState):
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: _dafny.Seq
            for compr_11_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])})).Elements:
                d_67_v0_: _dafny.Seq = compr_11_
                if (d_67_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([True])})):
                    coll11_ = coll11_.union(_dafny.Set([d_67_v0_]))
            return _dafny.Set(coll11_)
        return (_dafny.Map({len(_dafny.Map({False: -166})): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})), len(_dafny.SeqWithoutIsStrInference([733, len(_dafny.Map({True: 930})), len(iife11_()
        ), len(_dafny.Map({len(_dafny.Set({True})): False})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')]))]))]))}))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arlqe")) for d_68_i0_ in range(default__.abs(875))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yi")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgyhqbpjr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyuepp")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_69_i1_ in range(default__.abs(391))])]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fljjpgts"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfbd"))])))

    @staticmethod
    def m0(p0, p1, globalState):
        hi0_ = (p0) + (p0)
        for d_70_i0_ in range((0) - (p0), hi0_):
            d_71_v0_: _dafny.Seq
            d_71_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngu"))
            d_72_v1_: D16
            d_72_v1_ = D16_DC38(d_71_v0_, not(not(p1)), p0, p1)
            d_73_v2_: _dafny.Seq
            d_73_v2_ = _dafny.SeqWithoutIsStrInference([(d_72_v1_).cf57, d_71_v0_, d_71_v0_])
            d_74_v3_: _dafny.MultiSet
            d_74_v3_ = _dafny.MultiSet([p1])
            d_75_v4_: D1
            d_75_v4_ = D1_DC1(d_74_v3_)
            d_73_v2_ = default__.fm54(p0, d_75_v4_, default__.fm14(globalState), globalState)
            d_76_v5_: _dafny.Array
            def lambda0_(d_77_p1_):
                def lambda1_(d_78_i1_):
                    return _dafny.Set({d_77_p1_})

                return lambda1_

            init0_ = lambda0_(p1)
            nw0_ = _dafny.Array(None, 10)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_76_v5_ = nw0_
            d_79_v6_: str
            d_79_v6_ = _dafny.CodePoint('w')
            d_80_v7_: C9
            nw1_ = C9()
            nw1_.ctor__((d_76_v5_ if p1 else d_76_v5_), d_79_v6_)
            d_80_v7_ = nw1_
            d_81_v8_: _dafny.Map
            d_81_v8_ = _dafny.Map({p0: d_79_v6_})
            d_82_v9_: _dafny.Map
            d_82_v9_ = _dafny.Map({len(d_81_v8_): p1})
            source4_ = D14_DC32(((0) - (default__.fm14(globalState))) - (p0), ((d_82_v9_)[p0] if (p0) in (d_82_v9_) else True), p1)
            if source4_.is_DC31:
                d_83___mcc_h0_ = source4_.cf48
                d_84_cf48_ = d_83___mcc_h0_
                d_85_v10_: bool
                d_85_v10_ = True
                d_85_v10_ = p1
                d_85_v10_ = p1
                d_86_v11_: _dafny.Seq
                d_86_v11_ = _dafny.SeqWithoutIsStrInference([d_70_i0_, p0])
                d_87_v12_: _dafny.Seq
                d_87_v12_ = _dafny.SeqWithoutIsStrInference([(d_86_v11_)[default__.safeIndex(d_84_cf48_, len(d_86_v11_))], d_70_i0_, d_84_cf48_])
                d_88_v13_: _dafny.Map
                d_88_v13_ = _dafny.Map({d_87_v12_: p1})
                d_89_v14_: _dafny.Seq
                d_89_v14_ = _dafny.SeqWithoutIsStrInference([d_85_v10_, d_85_v10_, p1])
                d_88_v13_ = (d_88_v13_).set(d_87_v12_, (d_89_v14_)[default__.safeIndex(d_70_i0_, len(d_89_v14_))])
                d_90_v15_: C0
                nw2_ = C0()
                nw2_.ctor__()
                d_90_v15_ = nw2_
                d_91_v16_: C0
                d_91_v16_ = d_90_v15_
                d_91_v16_ = d_91_v16_
            elif source4_.is_DC32:
                d_92___mcc_h1_ = source4_.cf49
                d_93___mcc_h2_ = source4_.cf50
                d_94___mcc_h3_ = source4_.cf51
                d_95_cf51_ = d_94___mcc_h3_
                d_96_cf50_ = d_93___mcc_h2_
                d_97_cf49_ = d_92___mcc_h1_
                d_98_v17_: _dafny.Array
                def lambda2_(d_99_cf51_):
                    def lambda3_(d_100_i2_):
                        return d_99_cf51_

                    return lambda3_

                init1_ = lambda2_(d_95_cf51_)
                nw3_ = _dafny.Array(None, 1)
                for i0_1_ in range(nw3_.length(0)):
                    nw3_[i0_1_] = init1_(i0_1_)
                d_98_v17_ = nw3_
                index0_ = default__.safeIndex(782, (d_98_v17_).length(0))
                (d_98_v17_)[index0_] = p1
                d_101_v18_: T0
                nw4_ = C7()
                nw4_.ctor__()
                d_101_v18_ = nw4_
                index1_ = default__.safeIndex(782, (d_98_v17_).length(0))
                rhs0_ = d_95_cf51_
                rhs1_ = d_101_v18_
                lhs0_ = d_98_v17_
                lhs1_ = default__.safeIndex(782, (d_98_v17_).length(0))
                lhs0_[lhs1_] = rhs0_
                d_101_v18_ = rhs1_
                d_102_v19_: _dafny.Seq
                d_102_v19_ = _dafny.SeqWithoutIsStrInference([not(d_96_cf50_), True, d_95_cf51_, p1, (d_98_v17_)[default__.safeIndex(782, (d_98_v17_).length(0))]])
                d_103_v20_: _dafny.Seq
                d_103_v20_ = _dafny.SeqWithoutIsStrInference([d_97_cf49_])
                d_104_v21_: _dafny.Map
                d_104_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_103_v20_), (0) - (d_70_i0_), p0]): d_96_cf50_})
                d_105_v22_: _dafny.Seq
                d_105_v22_ = _dafny.SeqWithoutIsStrInference([(d_102_v19_)[default__.safeIndex(d_70_i0_, len(d_102_v19_))], d_95_cf51_, not (((d_104_v21_)[default__.fm31(-181, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otxi"))), globalState)] if (default__.fm31(-181, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otxi"))), globalState)) in (d_104_v21_) else d_96_cf50_)) or ((d_98_v17_)[default__.safeIndex(782, (d_98_v17_).length(0))])])
                d_105_v22_ = ((d_102_v19_) + (d_102_v19_)).set(default__.safeIndex(d_97_cf49_, len((d_102_v19_) + (d_102_v19_))), (d_70_i0_) != (d_70_i0_))
                d_106_v23_: _dafny.Map
                d_106_v23_ = _dafny.Map({(p0) >= (840): d_70_i0_})
                d_106_v23_ = (d_106_v23_).set(not((d_96_cf50_ if d_96_cf50_ else d_95_cf51_)), d_70_i0_)
                d_107_v24_: _dafny.Map
                d_107_v24_ = _dafny.Map({(d_98_v17_)[default__.safeIndex(782, (d_98_v17_).length(0))]: d_79_v6_})
                d_108_v25_: _dafny.Map
                d_108_v25_ = _dafny.Map({d_103_v20_: (d_107_v24_ if (d_105_v22_)[default__.safeIndex(d_97_cf49_, len(d_105_v22_))] else (d_107_v24_).set(not(p1), d_79_v6_))})
                d_108_v25_ = (d_108_v25_) | (d_108_v25_)
            elif source4_.is_DC30:
                d_109___mcc_h4_ = source4_.cf47
                d_110_cf47_ = d_109___mcc_h4_
                d_111_v26_: _dafny.Array
                nw5_ = _dafny.Array(None, 4)
                nw5_[int(0)] = p1
                nw5_[int(1)] = not(p1)
                nw5_[int(2)] = ((d_80_v7_).fm3(-674, d_70_i0_, globalState) if p1 else p1)
                nw5_[int(3)] = True
                d_111_v26_ = nw5_
                index2_ = default__.safeIndex(24, (d_111_v26_).length(0))
                (d_111_v26_)[index2_] = False
                d_112_v27_: int
                d_112_v27_ = 681
                d_113_v28_: _dafny.Map
                d_113_v28_ = _dafny.Map({(_dafny.MultiSet([d_70_i0_, d_112_v27_])).cardinality: d_70_i0_})
                d_114_v29_: C6
                nw6_ = C6()
                nw6_.ctor__(d_113_v28_)
                d_114_v29_ = nw6_
                d_115_v30_: _dafny.Map
                d_115_v30_ = _dafny.Map({p1: d_114_v29_})
                d_116_v31_: _dafny.Array
                nw7_ = _dafny.Array(None, 28)
                nw7_[int(0)] = d_114_v29_
                nw7_[int(1)] = d_114_v29_
                nw7_[int(2)] = d_114_v29_
                nw7_[int(3)] = d_114_v29_
                nw7_[int(4)] = d_114_v29_
                nw7_[int(5)] = d_114_v29_
                nw7_[int(6)] = ((d_115_v30_)[p1] if (p1) in (d_115_v30_) else d_114_v29_)
                nw7_[int(7)] = d_114_v29_
                nw7_[int(8)] = d_114_v29_
                nw7_[int(9)] = d_114_v29_
                nw7_[int(10)] = d_114_v29_
                nw7_[int(11)] = d_114_v29_
                nw7_[int(12)] = d_114_v29_
                nw7_[int(13)] = d_114_v29_
                nw7_[int(14)] = d_114_v29_
                nw7_[int(15)] = d_114_v29_
                nw7_[int(16)] = d_114_v29_
                nw7_[int(17)] = d_114_v29_
                nw7_[int(18)] = d_114_v29_
                nw7_[int(19)] = d_114_v29_
                nw7_[int(20)] = d_114_v29_
                nw7_[int(21)] = d_114_v29_
                nw7_[int(22)] = d_114_v29_
                nw7_[int(23)] = d_114_v29_
                nw7_[int(24)] = d_114_v29_
                nw7_[int(25)] = d_114_v29_
                nw7_[int(26)] = d_114_v29_
                nw7_[int(27)] = d_114_v29_
                d_116_v31_ = nw7_
                d_117_v32_: D28
                d_117_v32_ = D28_DC69(p0, d_116_v31_, d_79_v6_)
                index3_ = default__.safeIndex(24, (d_111_v26_).length(0))
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: int
                    for compr_12_ in _dafny.IntegerRange(-295, 165):
                        d_118_v33_: int = compr_12_
                        if ((-295) <= (d_118_v33_)) and ((d_118_v33_) < (165)):
                            coll12_[(d_118_v33_) - (d_112_v27_)] = _dafny.Map({d_70_i0_: p1})
                    return _dafny.Map(coll12_)
                rhs2_ = p1
                rhs3_ = (len(iife12_()
                )) * (p0)
                rhs4_ = 334
                rhs5_ = d_117_v32_
                rhs6_ = _dafny.SeqWithoutIsStrInference([d_79_v6_ for d_119_i3_ in range(default__.abs(40))])
                lhs2_ = d_111_v26_
                lhs3_ = default__.safeIndex(24, (d_111_v26_).length(0))
                lhs2_[lhs3_] = rhs2_
                d_112_v27_ = rhs3_
                d_112_v27_ = rhs4_
                d_117_v32_ = rhs5_
                d_71_v0_ = rhs6_
                d_120_v34_: C3
                nw8_ = C3()
                nw8_.ctor__(d_76_v5_, default__.fm0(p0, (0) - (d_112_v27_), globalState))
                d_120_v34_ = nw8_
                d_121_v35_: D20
                d_121_v35_ = D20_DC47(d_76_v5_)
                d_122_v36_: C10
                nw9_ = C10()
                nw9_.ctor__(p0, p1, (d_121_v35_).cf73, d_79_v6_)
                d_122_v36_ = nw9_
                d_123_v37_: bool
                d_123_v37_ = (d_122_v36_).f9
                d_124_v38_: _dafny.Array
                nw10_ = _dafny.Array(int(0), 23)
                d_124_v38_ = nw10_
                d_125_v39_: _dafny.Map
                d_125_v39_ = _dafny.Map({(d_122_v36_).f9: d_124_v38_})
                d_126_v40_: _dafny.Array
                nw11_ = _dafny.Array(None, 21)
                nw11_[int(0)] = d_124_v38_
                nw11_[int(1)] = (d_124_v38_ if False else ((d_125_v39_)[(d_122_v36_).f9] if ((d_122_v36_).f9) in (d_125_v39_) else d_124_v38_))
                nw11_[int(2)] = d_124_v38_
                nw11_[int(3)] = d_124_v38_
                nw11_[int(4)] = d_124_v38_
                nw11_[int(5)] = d_124_v38_
                nw11_[int(6)] = d_124_v38_
                nw11_[int(7)] = d_124_v38_
                nw11_[int(8)] = d_124_v38_
                nw11_[int(9)] = d_124_v38_
                nw11_[int(10)] = d_124_v38_
                nw11_[int(11)] = d_124_v38_
                nw11_[int(12)] = d_124_v38_
                nw11_[int(13)] = d_124_v38_
                nw11_[int(14)] = d_124_v38_
                nw11_[int(15)] = ((d_125_v39_)[p1] if (p1) in (d_125_v39_) else d_124_v38_)
                nw11_[int(16)] = d_124_v38_
                nw11_[int(17)] = d_124_v38_
                nw11_[int(18)] = d_124_v38_
                nw11_[int(19)] = d_124_v38_
                nw11_[int(20)] = d_124_v38_
                d_126_v40_ = nw11_
                index4_ = default__.safeIndex(409, (d_126_v40_).length(0))
                (d_126_v40_)[index4_] = d_124_v38_
                index5_ = default__.safeIndex(24, (d_111_v26_).length(0))
                index6_ = default__.safeIndex(409, (d_126_v40_).length(0))
                rhs7_ = d_122_v36_
                rhs8_ = d_123_v37_
                rhs9_ = ((0) - ((0) - ((d_122_v36_).f8))) + (d_112_v27_)
                rhs10_ = not (not ((d_111_v26_)[default__.safeIndex(24, (d_111_v26_).length(0))]) or (False)) or ((len(d_71_v0_)) >= (-330))
                rhs11_ = d_124_v38_
                lhs4_ = d_111_v26_
                lhs5_ = default__.safeIndex(24, (d_111_v26_).length(0))
                lhs6_ = d_126_v40_
                lhs7_ = default__.safeIndex(409, (d_126_v40_).length(0))
                d_122_v36_ = rhs7_
                d_123_v37_ = rhs8_
                d_112_v27_ = rhs9_
                lhs4_[lhs5_] = rhs10_
                lhs6_[lhs7_] = rhs11_
                d_71_v0_ = d_71_v0_
            elif True:
                d_127___mcc_h5_ = source4_.cf52
                d_128_cf52_ = d_127___mcc_h5_
                d_129_v41_: bool
                d_129_v41_ = True
                d_129_v41_ = d_129_v41_
                d_130_v42_: _dafny.Seq
                d_131_v43_: int
                out0_: _dafny.Seq
                out1_: int
                out0_, out1_ = (d_80_v7_).m1(globalState)
                d_130_v42_ = out0_
                d_131_v43_ = out1_
                d_132_v44_: C0
                nw12_ = C0()
                nw12_.ctor__()
                d_132_v44_ = nw12_
                d_133_v45_: _dafny.Map
                d_133_v45_ = _dafny.Map({d_132_v44_: d_71_v0_})
                d_133_v45_ = (d_133_v45_).set(d_132_v44_, default__.fm12(globalState))
                d_134_v46_: _dafny.Array
                nw13_ = _dafny.Array(None, 7)
                d_134_v46_ = nw13_
                index7_ = default__.safeIndex(373, (d_134_v46_).length(0))
                (d_134_v46_)[index7_] = d_80_v7_
                index8_ = default__.safeIndex(373, (d_134_v46_).length(0))
                (d_134_v46_)[index8_] = d_80_v7_
            d_135_v47_: bool
            d_135_v47_ = False
            d_135_v47_ = p1
        d_136_i4_: int
        d_136_i4_ = 0
        with _dafny.label("0"):
            while (not(p1) if (p1) or (p1) else p1):
                with _dafny.c_label("0"):
                    if (d_136_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_136_i4_ = (d_136_i4_) + (1)
                    d_137_v48_: _dafny.Array
                    nw14_ = _dafny.Array(None, 16)
                    d_137_v48_ = nw14_
                    d_138_v49_: C5
                    nw15_ = C5()
                    nw15_.ctor__()
                    d_138_v49_ = nw15_
                    index9_ = default__.safeIndex(779, (d_137_v48_).length(0))
                    (d_137_v48_)[index9_] = d_138_v49_
                    index10_ = default__.safeIndex(779, (d_137_v48_).length(0))
                    (d_137_v48_)[index10_] = d_138_v49_
                    d_139_v50_: bool
                    d_139_v50_ = True
                    d_139_v50_ = p1
                    d_140_v51_: _dafny.MultiSet
                    d_140_v51_ = _dafny.MultiSet([686])
                    d_141_v52_: D8
                    d_141_v52_ = D8_DC18(d_140_v51_)
                    source5_ = d_141_v52_
                    if source5_.is_DC19:
                        d_142___mcc_h6_ = source5_.cf29
                        d_143___mcc_h7_ = source5_.cf30
                        d_144_cf30_ = d_143___mcc_h7_
                        d_145_cf29_ = d_142___mcc_h6_
                        d_145_cf29_ = default__.fm13(p1, (d_144_cf30_) == (d_139_v50_), (p1 if p1 else p1), globalState)
                        d_146_v53_: _dafny.Array
                        nw16_ = _dafny.Array(_dafny.Set({}), 27)
                        d_146_v53_ = nw16_
                        d_147_v54_: str
                        d_147_v54_ = _dafny.CodePoint('x')
                        d_148_v55_: C9
                        nw17_ = C9()
                        nw17_.ctor__(d_146_v53_, d_147_v54_)
                        d_148_v55_ = nw17_
                        d_149_v56_: _dafny.Map
                        d_149_v56_ = _dafny.Map({d_145_cf29_: p0})
                        d_150_v57_: _dafny.Map
                        d_150_v57_ = _dafny.Map({d_148_v55_: d_149_v56_})
                        d_151_v59_: C6
                        nw18_ = C6()
                        def iife13_():
                            coll13_ = _dafny.Map()
                            compr_13_: int
                            for compr_13_ in _dafny.IntegerRange(146, 160):
                                d_152_v58_: int = compr_13_
                                if ((146) <= (d_152_v58_)) and ((d_152_v58_) < (160)):
                                    coll13_[default__.safeDivisionInt(d_152_v58_, p0)] = d_145_cf29_
                            return _dafny.Map(coll13_)
                        nw18_.ctor__(((d_150_v57_)[d_148_v55_] if (d_148_v55_) in (d_150_v57_) else iife13_()
                        ))
                        d_151_v59_ = nw18_
                        d_153_v60_: _dafny.Array
                        nw19_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_153_v60_ = nw19_
                        d_154_v61_: _dafny.Array
                        nw20_ = _dafny.Array(_dafny.Seq({}), 7)
                        d_154_v61_ = nw20_
                        index11_ = default__.safeIndex(428, (d_153_v60_).length(0))
                        (d_153_v60_)[index11_] = d_154_v61_
                        d_155_v62_: C1
                        nw21_ = C1()
                        nw21_.ctor__(d_146_v53_, d_147_v54_)
                        d_155_v62_ = nw21_
                        d_156_v63_: _dafny.Seq
                        d_156_v63_ = _dafny.SeqWithoutIsStrInference([d_155_v62_])
                        d_157_v64_: _dafny.Map
                        d_157_v64_ = _dafny.Map({d_145_cf29_: d_156_v63_})
                        index12_ = default__.safeIndex(428, (d_153_v60_).length(0))
                        rhs12_ = d_154_v61_
                        rhs13_ = default__.safeModuloInt(p0, d_145_cf29_)
                        rhs14_ = len(((d_157_v64_).set(p0, d_156_v63_) if p1 else d_157_v64_))
                        lhs8_ = d_153_v60_
                        lhs9_ = default__.safeIndex(428, (d_153_v60_).length(0))
                        lhs8_[lhs9_] = rhs12_
                        d_145_cf29_ = rhs13_
                        d_145_cf29_ = rhs14_
                        d_158_v65_: C2
                        nw22_ = C2()
                        nw22_.ctor__(d_146_v53_, d_147_v54_)
                        d_158_v65_ = nw22_
                        nw23_ = C2()
                        nw23_.ctor__(d_146_v53_, d_147_v54_)
                        d_158_v65_ = nw23_
                    elif True:
                        d_159___mcc_h8_ = source5_.cf28
                        d_160_cf28_ = d_159___mcc_h8_
                        d_139_v50_ = False
                        d_161_v66_: _dafny.Array
                        def lambda4_(d_162_v50_):
                            def lambda5_(d_163_i5_):
                                return d_162_v50_

                            return lambda5_

                        init2_ = lambda4_(d_139_v50_)
                        nw24_ = _dafny.Array(None, 24)
                        for i0_2_ in range(nw24_.length(0)):
                            nw24_[i0_2_] = init2_(i0_2_)
                        d_161_v66_ = nw24_
                        d_164_v67_: _dafny.Map
                        d_164_v67_ = _dafny.Map({default__.fm1((0) - (p0), p0, d_139_v50_, False, globalState): d_161_v66_})
                        d_165_v68_: _dafny.Set
                        d_165_v68_ = _dafny.Set({(D30_DC73(p0, d_139_v50_, d_161_v66_)).cf125, d_161_v66_, d_161_v66_, ((d_164_v67_)[d_139_v50_] if (d_139_v50_) in (d_164_v67_) else d_161_v66_)})
                        d_165_v68_ = (d_165_v68_) | (d_165_v68_)
                        d_166_v69_: _dafny.Array
                        nw25_ = _dafny.Array(int(0), 2)
                        d_166_v69_ = nw25_
                        index13_ = default__.safeIndex(654, (d_166_v69_).length(0))
                        (d_166_v69_)[index13_] = (p0) * (276)
                        d_167_v70_: int
                        d_167_v70_ = -223
                        d_168_v71_: D1
                        d_168_v71_ = D1_DC1(_dafny.MultiSet([d_139_v50_]))
                        d_169_v72_: _dafny.MultiSet
                        d_169_v72_ = _dafny.MultiSet([p1])
                        index14_ = default__.safeIndex(654, (d_166_v69_).length(0))
                        rhs15_ = (p0) * (p0)
                        rhs16_ = default__.safeDivisionInt(d_167_v70_, p0)
                        rhs17_ = d_166_v69_
                        rhs18_ = (d_169_v72_).ispropersubset((_dafny.MultiSet([p1, d_139_v50_])) | ((d_168_v71_).cf1))
                        lhs10_ = d_166_v69_
                        lhs11_ = default__.safeIndex(654, (d_166_v69_).length(0))
                        lhs10_[lhs11_] = rhs15_
                        d_167_v70_ = rhs16_
                        d_166_v69_ = rhs17_
                        d_139_v50_ = rhs18_
                        d_170_v73_: _dafny.Seq
                        d_170_v73_ = _dafny.SeqWithoutIsStrInference([d_139_v50_, d_139_v50_, p1, p1, p1])
                        d_171_v74_: _dafny.Set
                        d_171_v74_ = _dafny.Set({d_139_v50_})
                        d_172_v75_: D20
                        d_172_v75_ = D20_DC48(True, (d_170_v73_) + (default__.fm21((d_169_v72_).cardinality, d_167_v70_, d_171_v74_, globalState)), d_170_v73_, (d_166_v69_)[default__.safeIndex(654, (d_166_v69_).length(0))])
                        pat_let_tv0_ = d_166_v69_
                        pat_let_tv1_ = d_166_v69_
                        pat_let_tv2_ = p1
                        def iife14_(_pat_let0_0):
                            def iife15_(d_173_dt__update__tmp_h0_):
                                def iife16_(_pat_let1_0):
                                    def iife17_(d_174_dt__update_hcf77_h0_):
                                        def iife18_(_pat_let2_0):
                                            def iife19_(d_175_dt__update_hcf74_h0_):
                                                return D20_DC48(d_175_dt__update_hcf74_h0_, (d_173_dt__update__tmp_h0_).cf75, (d_173_dt__update__tmp_h0_).cf76, d_174_dt__update_hcf77_h0_)
                                            return iife19_(_pat_let2_0)
                                        return iife18_(pat_let_tv2_)
                                    return iife17_(_pat_let1_0)
                                return iife16_((pat_let_tv1_)[default__.safeIndex(654, (pat_let_tv0_).length(0))])
                            return iife15_(_pat_let0_0)
                        d_172_v75_ = iife14_(default__.fm49(globalState))
                    d_176_v76_: _dafny.MultiSet
                    d_176_v76_ = _dafny.MultiSet([_dafny.MultiSet([True, d_139_v50_])])
                    d_177_v77_: _dafny.MultiSet
                    d_177_v77_ = _dafny.MultiSet([d_139_v50_, p1])
                    d_178_v78_: _dafny.Seq
                    d_178_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dasn"))
                    d_179_v79_: _dafny.Map
                    d_179_v79_ = _dafny.Map({p0: len(d_178_v78_)})
                    d_180_v80_: _dafny.Map
                    d_180_v80_ = _dafny.Map({len(d_179_v79_): p1})
                    d_181_v81_: _dafny.Seq
                    d_181_v81_ = _dafny.SeqWithoutIsStrInference([d_180_v80_])
                    d_182_v82_: _dafny.Seq
                    d_182_v82_ = _dafny.SeqWithoutIsStrInference([d_179_v79_])
                    d_183_v83_: _dafny.Map
                    d_183_v83_ = _dafny.Map({False: p0})
                    d_184_v84_: _dafny.Array
                    nw26_ = _dafny.Array(None, 20)
                    nw26_[int(0)] = 248
                    nw26_[int(1)] = default__.safeModuloInt(p0, ((d_176_v76_)[d_177_v77_] if (d_177_v77_) in (d_176_v76_) else p0))
                    nw26_[int(2)] = p0
                    nw26_[int(3)] = p0
                    nw26_[int(4)] = p0
                    nw26_[int(5)] = len(d_178_v78_)
                    nw26_[int(6)] = 238
                    nw26_[int(7)] = p0
                    nw26_[int(8)] = (0) - (default__.fm14(globalState))
                    nw26_[int(9)] = p0
                    nw26_[int(10)] = (641) * (665)
                    nw26_[int(11)] = default__.fm13(False, d_139_v50_, p1, globalState)
                    nw26_[int(12)] = p0
                    nw26_[int(13)] = p0
                    nw26_[int(14)] = (0) - (len(d_181_v81_))
                    nw26_[int(15)] = p0
                    nw26_[int(16)] = (p0) + ((0) - (p0))
                    nw26_[int(17)] = p0
                    nw26_[int(18)] = (default__.fm14(globalState)) + ((0) - (len(d_182_v82_)))
                    nw26_[int(19)] = len(d_183_v83_)
                    d_184_v84_ = nw26_
                    d_185_v85_: _dafny.Seq
                    d_185_v85_ = _dafny.SeqWithoutIsStrInference([d_139_v50_, d_139_v50_, False, d_139_v50_])
                    index15_ = default__.safeIndex(494, (d_184_v84_).length(0))
                    (d_184_v84_)[index15_] = (default__.fm30(d_139_v50_, d_185_v85_, globalState)).cardinality
                    index16_ = default__.safeIndex(494, (d_184_v84_).length(0))
                    (d_184_v84_)[index16_] = default__.safeDivisionInt((0) - (default__.fm13(d_139_v50_, p1, False, globalState)), p0)
                    pass
            pass
        d_186_v86_: _dafny.Seq
        d_186_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnttcvusf"))
        d_187_v87_: _dafny.Seq
        d_187_v87_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_188_i7_ in range(default__.abs(849))]), d_186_v86_, d_186_v86_, d_186_v86_])
        hi1_ = (_dafny.MultiSet([default__.fm1(len(d_187_v87_), p0, p1, p1, globalState)])).cardinality
        for d_189_i6_ in range(p0, hi1_):
            d_190_v88_: int
            d_190_v88_ = 204
            d_190_v88_ = p0
            d_191_v89_: _dafny.Set
            d_191_v89_ = _dafny.Set({False, p1})
            d_192_v90_: _dafny.Seq
            d_192_v90_ = _dafny.SeqWithoutIsStrInference([p1])
            d_193_v91_: _dafny.Array
            nw27_ = _dafny.Array(None, 19)
            nw27_[int(0)] = d_191_v89_
            nw27_[int(1)] = _dafny.Set({default__.fm1(p0, len(d_192_v90_), p1, True, globalState)})
            nw27_[int(2)] = d_191_v89_
            nw27_[int(3)] = d_191_v89_
            nw27_[int(4)] = d_191_v89_
            nw27_[int(5)] = d_191_v89_
            nw27_[int(6)] = d_191_v89_
            nw27_[int(7)] = d_191_v89_
            nw27_[int(8)] = d_191_v89_
            nw27_[int(9)] = d_191_v89_
            nw27_[int(10)] = _dafny.Set({False})
            nw27_[int(11)] = d_191_v89_
            nw27_[int(12)] = d_191_v89_
            nw27_[int(13)] = d_191_v89_
            nw27_[int(14)] = _dafny.Set({p1})
            nw27_[int(15)] = d_191_v89_
            nw27_[int(16)] = d_191_v89_
            nw27_[int(17)] = d_191_v89_
            nw27_[int(18)] = d_191_v89_
            d_193_v91_ = nw27_
            d_194_v92_: _dafny.Seq
            d_194_v92_ = _dafny.SeqWithoutIsStrInference([d_189_i6_])
            d_195_v93_: C2
            nw28_ = C2()
            nw28_.ctor__(d_193_v91_, default__.fm0((_dafny.MultiSet(d_194_v92_)).cardinality, d_190_v88_, globalState))
            d_195_v93_ = nw28_
            d_196_v94_: _dafny.Map
            d_196_v94_ = _dafny.Map({p1: d_195_v93_})
            d_197_v95_: _dafny.Map
            d_197_v95_ = _dafny.Map({len(d_196_v94_): d_190_v88_})
            d_198_v96_: C6
            nw29_ = C6()
            nw29_.ctor__(d_197_v95_)
            d_198_v96_ = nw29_
            d_199_v97_: _dafny.Map
            d_199_v97_ = _dafny.Map({d_198_v96_: not(p1)})
            d_199_v97_ = (d_199_v97_).set(d_198_v96_, p1)
            d_200_v98_: _dafny.Array
            nw30_ = _dafny.Array(False, 24)
            d_200_v98_ = nw30_
            index17_ = default__.safeIndex(256, (d_200_v98_).length(0))
            (d_200_v98_)[index17_] = p1
            index18_ = default__.safeIndex(389, (d_200_v98_).length(0))
            (d_200_v98_)[index18_] = p1
            d_201_v99_: T1
            nw31_ = C2()
            nw31_.ctor__(d_193_v91_, _dafny.CodePoint('u'))
            d_201_v99_ = nw31_
            d_202_v100_: _dafny.Set
            d_202_v100_ = _dafny.Set({d_201_v99_, d_201_v99_, d_201_v99_})
            d_203_v101_: _dafny.Set
            d_203_v101_ = _dafny.Set({d_202_v100_})
            index19_ = default__.safeIndex(256, (d_200_v98_).length(0))
            index20_ = default__.safeIndex(389, (d_200_v98_).length(0))
            rhs19_ = ((len(d_191_v89_)) + (d_190_v88_)) == (len(default__.fm24(p0, globalState)))
            rhs20_ = d_186_v86_
            rhs21_ = (p1) == (False)
            rhs22_ = _dafny.Set({(d_202_v100_).intersection(d_202_v100_)})
            lhs12_ = d_200_v98_
            lhs13_ = default__.safeIndex(256, (d_200_v98_).length(0))
            lhs14_ = d_200_v98_
            lhs15_ = default__.safeIndex(389, (d_200_v98_).length(0))
            lhs12_[lhs13_] = rhs19_
            d_186_v86_ = rhs20_
            lhs14_[lhs15_] = rhs21_
            d_203_v101_ = rhs22_
            d_204_v102_: _dafny.Map
            d_204_v102_ = _dafny.Map({p0: (d_200_v98_)[default__.safeIndex(256, (d_200_v98_).length(0))]})
            index21_ = default__.safeIndex(256, (d_200_v98_).length(0))
            (d_200_v98_)[index21_] = not ((p1) and ((d_201_v99_).fm3(default__.fm14(globalState), len(d_204_v102_), globalState))) or ((p1 if p1 else p1))
        d_205_v103_: _dafny.Set
        d_205_v103_ = _dafny.Set({p1, p1, default__.fm1(p0, len(d_186_v86_), p1, False, globalState), p1, p1})
        d_206_v104_: _dafny.Array
        nw32_ = _dafny.Array(None, 5)
        nw32_[int(0)] = d_205_v103_
        nw32_[int(1)] = d_205_v103_
        nw32_[int(2)] = d_205_v103_
        nw32_[int(3)] = d_205_v103_
        nw32_[int(4)] = d_205_v103_
        d_206_v104_ = nw32_
        d_207_v105_: C2
        nw33_ = C2()
        nw33_.ctor__(d_206_v104_, _dafny.CodePoint('i'))
        d_207_v105_ = nw33_
        hi2_ = p0
        for d_208_i8_ in range(p0, hi2_):
            d_209_v106_: bool
            d_209_v106_ = False
            d_210_v107_: _dafny.Seq
            d_210_v107_ = _dafny.SeqWithoutIsStrInference([p0])
            d_209_v106_ = not(default__.fm1(len(d_210_v107_), d_208_i8_, d_209_v106_, not(p1), globalState))
            if d_209_v106_:
                d_209_v106_ = p1
                d_211_v108_: str
                d_211_v108_ = _dafny.CodePoint('a')
                d_212_v109_: C3
                nw34_ = C3()
                nw34_.ctor__(d_206_v104_, d_211_v108_)
                d_212_v109_ = nw34_
                d_212_v109_ = d_212_v109_
                d_209_v106_ = (d_212_v109_).fm5(p0, p1, p0, p0, globalState)
                d_213_v110_: C9
                nw35_ = C9()
                nw35_.ctor__((d_206_v104_ if d_209_v106_ else d_206_v104_), d_211_v108_)
                d_213_v110_ = nw35_
                d_214_v111_: _dafny.Array
                nw36_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                d_214_v111_ = nw36_
                d_215_v112_: _dafny.Map
                d_215_v112_ = _dafny.Map({d_211_v108_: d_211_v108_})
                index22_ = default__.safeIndex(776, (d_214_v111_).length(0))
                (d_214_v111_)[index22_] = ((d_215_v112_)[d_211_v108_] if (d_211_v108_) in (d_215_v112_) else d_211_v108_)
                d_216_v113_: _dafny.Map
                d_216_v113_ = _dafny.Map({491: p1})
                d_217_v114_: _dafny.Seq
                d_217_v114_ = _dafny.SeqWithoutIsStrInference([not((d_207_v105_).fm5(d_208_i8_, ((d_216_v113_)[d_208_i8_] if (d_208_i8_) in (d_216_v113_) else False), p0, 840, globalState)), p1])
                d_218_v115_: _dafny.Map
                d_218_v115_ = _dafny.Map({default__.safeModuloInt(len(d_217_v114_), len(d_216_v113_)): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))) < (d_186_v86_)})
                index23_ = default__.safeIndex(776, (d_214_v111_).length(0))
                rhs23_ = (d_186_v86_) + (_dafny.SeqWithoutIsStrInference([d_211_v108_ for d_219_i9_ in range(default__.abs(-524))]))
                rhs24_ = d_211_v108_
                rhs25_ = ((d_218_v115_)[len((d_186_v86_ if p1 else d_186_v86_))] if (len((d_186_v86_ if p1 else d_186_v86_))) in (d_218_v115_) else p1)
                lhs16_ = d_214_v111_
                lhs17_ = default__.safeIndex(776, (d_214_v111_).length(0))
                d_186_v86_ = rhs23_
                lhs16_[lhs17_] = rhs24_
                d_209_v106_ = rhs25_
            elif True:
                d_220_v116_: int
                d_220_v116_ = 401
                d_221_v117_: _dafny.MultiSet
                d_221_v117_ = _dafny.MultiSet([878])
                d_222_v118_: str
                d_222_v118_ = _dafny.CodePoint('s')
                d_223_v119_: _dafny.Array
                nw37_ = _dafny.Array(None, 3)
                nw37_[int(0)] = d_222_v118_
                nw37_[int(1)] = d_222_v118_
                nw37_[int(2)] = d_222_v118_
                d_223_v119_ = nw37_
                d_224_v120_: _dafny.Set
                d_224_v120_ = _dafny.Set({d_223_v119_})
                d_225_v121_: _dafny.MultiSet
                d_225_v121_ = _dafny.MultiSet([d_209_v106_, d_209_v106_])
                rhs26_ = ((-464 if d_209_v106_ else d_208_i8_)) + (default__.safeModuloInt(d_220_v116_, p0))
                rhs27_ = (0) - (p0)
                rhs28_ = (d_208_i8_) < (((d_221_v117_)[len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_226_i10_ in range(default__.abs(277))]))] if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_227_i10_ in range(default__.abs(277))]))) in (d_221_v117_) else len(d_224_v120_)))
                rhs29_ = (p1 if False else not((d_225_v121_).issubset(_dafny.MultiSet([p1]))))
                d_220_v116_ = rhs26_
                d_220_v116_ = rhs27_
                d_209_v106_ = rhs28_
                d_209_v106_ = rhs29_
                def iife20_():
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(674, 620):
                        d_228_v122_: int = compr_14_
                        if ((674) <= (d_228_v122_)) and ((d_228_v122_) < (620)):
                            coll14_[(d_228_v122_) - (p0)] = p0
                    return _dafny.Map(coll14_)
                d_220_v116_ = len(iife20_()
                )
                d_229_v123_: C9
                nw38_ = C9()
                nw38_.ctor__((d_206_v104_ if d_209_v106_ else d_206_v104_), _dafny.CodePoint('u'))
                d_229_v123_ = nw38_
                d_229_v123_ = d_229_v123_
                d_230_v124_: C9
                nw39_ = C9()
                nw39_.ctor__(d_206_v104_, (d_222_v118_ if True else d_222_v118_))
                d_230_v124_ = nw39_
                d_220_v116_ = d_220_v116_
            d_231_v125_: int
            d_231_v125_ = 894
            d_231_v125_ = (d_208_i8_) + (len(d_186_v86_))
            d_231_v125_ = p0
        d_232_v126_: D1
        d_232_v126_ = D1_DC2(p0, p0)
        d_232_v126_ = d_232_v126_

    @staticmethod
    def Main(noArgsParameter__):
        d_233_v0_: bool
        d_233_v0_ = True
        d_234_v1_: _dafny.Set
        d_234_v1_ = _dafny.Set({d_233_v0_})
        d_235_v2_: _dafny.Array
        nw40_ = _dafny.Array(False, 25)
        d_235_v2_ = nw40_
        d_236_globalState_: GlobalState
        nw41_ = GlobalState()
        nw41_.ctor__((d_234_v1_) - (d_234_v1_), 214, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([836, 40])): d_233_v0_})) for d_237_i0_ in range(default__.abs(358))]), (d_235_v2_ if d_233_v0_ else d_235_v2_))
        d_236_globalState_ = nw41_
        index24_ = default__.safeIndex(398, (d_235_v2_).length(0))
        (d_235_v2_)[index24_] = d_233_v0_
        d_238_v3_: int
        d_238_v3_ = 263
        index25_ = default__.safeIndex(398, (d_235_v2_).length(0))
        (d_235_v2_)[index25_] = (len(_dafny.Set({d_233_v0_, False, d_233_v0_}))) >= (d_238_v3_)
        d_239_v4_: str
        d_239_v4_ = _dafny.CodePoint('r')
        d_240_v5_: _dafny.Array
        nw42_ = _dafny.Array(None, 12)
        nw42_[int(0)] = d_239_v4_
        nw42_[int(1)] = _dafny.CodePoint('d')
        nw42_[int(2)] = d_239_v4_
        nw42_[int(3)] = d_239_v4_
        nw42_[int(4)] = d_239_v4_
        nw42_[int(5)] = _dafny.CodePoint('h')
        nw42_[int(6)] = d_239_v4_
        nw42_[int(7)] = (d_239_v4_ if d_233_v0_ else d_239_v4_)
        nw42_[int(8)] = (default__.fm0(d_238_v3_, d_238_v3_, d_236_globalState_) if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))] else default__.fm0(d_238_v3_, d_238_v3_, d_236_globalState_))
        nw42_[int(9)] = _dafny.CodePoint('l')
        nw42_[int(10)] = d_239_v4_
        nw42_[int(11)] = d_239_v4_
        d_240_v5_ = nw42_
        index26_ = default__.safeIndex(218, (d_240_v5_).length(0))
        (d_240_v5_)[index26_] = d_239_v4_
        d_241_v6_: _dafny.Seq
        d_241_v6_ = _dafny.SeqWithoutIsStrInference([(0) - (d_238_v3_)])
        index27_ = default__.safeIndex(218, (d_240_v5_).length(0))
        rhs30_ = _dafny.CodePoint('m')
        rhs31_ = (d_241_v6_)[default__.safeIndex(330, len(d_241_v6_))]
        rhs32_ = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
        lhs18_ = d_240_v5_
        lhs19_ = default__.safeIndex(218, (d_240_v5_).length(0))
        lhs18_[lhs19_] = rhs30_
        d_238_v3_ = rhs31_
        d_233_v0_ = rhs32_
        if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]:
            d_242_v7_: _dafny.MultiSet
            d_242_v7_ = _dafny.MultiSet([(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])])
            d_243_v8_: _dafny.Map
            d_243_v8_ = _dafny.Map({d_238_v3_: d_242_v7_})
            d_244_v9_: _dafny.Map
            d_244_v9_ = _dafny.Map({d_238_v3_: (d_242_v7_).issubset(((d_243_v8_)[389] if (389) in (d_243_v8_) else _dafny.MultiSet([(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]])))})
            d_245_v10_: D1
            d_245_v10_ = D1_DC1(d_242_v7_)
            d_246_v11_: D2
            d_246_v11_ = D2_DC3(d_235_v2_)
            index28_ = default__.safeIndex(398, (d_235_v2_).length(0))
            rhs33_ = ((d_244_v9_)[default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference([((d_245_v10_).cf1).cardinality]))), d_238_v3_)] if (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference([((d_245_v10_).cf1).cardinality]))), d_238_v3_)) in (d_244_v9_) else d_233_v0_)
            rhs34_ = False
            rhs35_ = (d_246_v11_).cf4
            lhs20_ = d_235_v2_
            lhs21_ = default__.safeIndex(398, (d_235_v2_).length(0))
            d_233_v0_ = rhs33_
            lhs20_[lhs21_] = rhs34_
            d_235_v2_ = rhs35_
            d_247_v12_: _dafny.Set
            d_247_v12_ = _dafny.Set({d_235_v2_})
            d_248_v13_: _dafny.Set
            d_248_v13_ = d_247_v12_
            index29_ = default__.safeIndex(398, (d_235_v2_).length(0))
            (d_235_v2_)[index29_] = ((d_248_v13_)).issubset((d_247_v12_) - (d_247_v12_))
            d_249_v14_: _dafny.Map
            d_249_v14_ = _dafny.Map({-845: d_235_v2_})
            d_250_v15_: _dafny.Seq
            d_250_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
            d_251_v16_: _dafny.Set
            d_251_v16_ = _dafny.Set({d_238_v3_, d_238_v3_})
            d_252_v17_: _dafny.Array
            nw43_ = _dafny.Array(None, 13)
            nw43_[int(0)] = d_235_v2_
            nw43_[int(1)] = ((d_249_v14_)[len(d_250_v15_)] if (len(d_250_v15_)) in (d_249_v14_) else d_235_v2_)
            nw43_[int(2)] = d_235_v2_
            nw43_[int(3)] = d_235_v2_
            nw43_[int(4)] = d_235_v2_
            nw43_[int(5)] = (d_235_v2_ if default__.fm1(d_238_v3_, len(d_251_v16_), True, d_233_v0_, d_236_globalState_) else d_235_v2_)
            nw43_[int(6)] = d_235_v2_
            nw43_[int(7)] = d_235_v2_
            nw43_[int(8)] = d_235_v2_
            nw43_[int(9)] = (d_246_v11_).cf4
            nw43_[int(10)] = d_235_v2_
            nw43_[int(11)] = d_235_v2_
            nw43_[int(12)] = d_235_v2_
            d_252_v17_ = nw43_
            index30_ = default__.safeIndex(320, (d_252_v17_).length(0))
            (d_252_v17_)[index30_] = d_235_v2_
            index31_ = default__.safeIndex(320, (d_252_v17_).length(0))
            (d_252_v17_)[index31_] = (d_246_v11_).cf4
            d_253_v18_: _dafny.Array
            def lambda6_(d_254_v16_):
                def lambda7_(d_255_i1_):
                    return d_254_v16_

                return lambda7_

            init3_ = lambda6_(d_251_v16_)
            nw44_ = _dafny.Array(None, 18)
            for i0_3_ in range(nw44_.length(0)):
                nw44_[i0_3_] = init3_(i0_3_)
            d_253_v18_ = nw44_
            d_256_v19_: _dafny.Map
            d_256_v19_ = _dafny.Map({d_238_v3_: d_253_v18_})
            d_256_v19_ = (d_256_v19_).set(d_238_v3_, d_253_v18_)
            d_257_v20_: _dafny.Set
            d_257_v20_ = _dafny.Set({default__.fm2(d_238_v3_, d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], len(d_250_v15_), d_236_globalState_)})
            d_251_v16_ = _dafny.Set({default__.safeModuloInt(len(d_257_v20_), d_238_v3_)})
        elif True:
            d_258_v21_: _dafny.Map
            d_258_v21_ = _dafny.Map({d_235_v2_: (-994) > ((0) - (len(d_241_v6_)))})
            d_238_v3_ = len(d_258_v21_)
            d_259_v22_: _dafny.Seq
            d_259_v22_ = _dafny.SeqWithoutIsStrInference([d_233_v0_])
            if (((_dafny.MultiSet(d_259_v22_)) - (_dafny.MultiSet([True]))).cardinality) == (463):
                d_260_v23_: _dafny.Map
                d_260_v23_ = _dafny.Map({(d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))]: not(d_233_v0_)})
                d_260_v23_ = (d_260_v23_).set((d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))], (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
                d_261_v24_: _dafny.Seq
                d_261_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp"))
                d_262_v25_: _dafny.MultiSet
                d_262_v25_ = _dafny.MultiSet([d_238_v3_, d_238_v3_])
                d_263_v26_: _dafny.Seq
                d_263_v26_ = _dafny.SeqWithoutIsStrInference([d_262_v25_])
                rhs36_ = len((d_261_v24_).set(default__.safeIndex((len(d_261_v24_)) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tslisom"))).set(default__.safeIndex(len(d_261_v24_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tslisom")))), (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))]))), len(d_261_v24_)), (d_261_v24_)[default__.safeIndex(d_238_v3_, len(d_261_v24_))]))
                rhs37_ = (-232) * (d_238_v3_)
                rhs38_ = (((_dafny.MultiSet([d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_])).cardinality) + (d_238_v3_) if (d_259_v22_)[default__.safeIndex(((d_263_v26_)[default__.safeIndex(d_238_v3_, len(d_263_v26_))]).cardinality, len(d_259_v22_))] else (0) - (d_238_v3_))
                rhs39_ = d_238_v3_
                rhs40_ = d_233_v0_
                d_238_v3_ = rhs36_
                d_238_v3_ = rhs37_
                d_238_v3_ = rhs38_
                d_238_v3_ = rhs39_
                d_233_v0_ = rhs40_
                d_264_v27_: _dafny.Seq
                d_264_v27_ = _dafny.SeqWithoutIsStrInference([d_234_v1_])
                default__.m0(len((d_234_v1_).intersection((d_264_v27_)[default__.safeIndex((0) - (d_238_v3_), len(d_264_v27_))])), d_233_v0_, d_236_globalState_)
                d_265_v28_: D2
                d_265_v28_ = D2_DC5(d_261_v24_)
                d_265_v28_ = d_265_v28_
                d_266_v29_: _dafny.Array
                nw45_ = _dafny.Array(None, 25)
                nw45_[int(0)] = d_234_v1_
                nw45_[int(1)] = default__.fm26(d_238_v3_, d_238_v3_, d_236_globalState_)
                nw45_[int(2)] = d_234_v1_
                nw45_[int(3)] = d_234_v1_
                nw45_[int(4)] = d_234_v1_
                nw45_[int(5)] = d_234_v1_
                nw45_[int(6)] = d_234_v1_
                nw45_[int(7)] = d_234_v1_
                nw45_[int(8)] = d_234_v1_
                nw45_[int(9)] = d_234_v1_
                nw45_[int(10)] = _dafny.Set({d_233_v0_})
                nw45_[int(11)] = d_234_v1_
                nw45_[int(12)] = d_234_v1_
                nw45_[int(13)] = default__.fm26(d_238_v3_, len(d_259_v22_), d_236_globalState_)
                nw45_[int(14)] = _dafny.Set({d_233_v0_})
                nw45_[int(15)] = d_234_v1_
                nw45_[int(16)] = d_234_v1_
                nw45_[int(17)] = d_234_v1_
                nw45_[int(18)] = d_234_v1_
                nw45_[int(19)] = d_234_v1_
                nw45_[int(20)] = d_234_v1_
                nw45_[int(21)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], False})
                nw45_[int(22)] = d_234_v1_
                nw45_[int(23)] = d_234_v1_
                nw45_[int(24)] = default__.fm26(len(d_261_v24_), len(d_234_v1_), d_236_globalState_)
                d_266_v29_ = nw45_
                d_267_v30_: C4
                nw46_ = C4()
                nw46_.ctor__(d_238_v3_, 116, d_266_v29_, d_239_v4_)
                d_267_v30_ = nw46_
            elif True:
                d_268_v31_: _dafny.Map
                d_268_v31_ = _dafny.Map({d_234_v1_: (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
                d_269_v32_: _dafny.MultiSet
                d_269_v32_ = _dafny.MultiSet([(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (d_268_v31_) == (d_268_v31_)])
                d_238_v3_ = (0) - ((d_269_v32_).cardinality)
                default__.m0(d_238_v3_, ((d_241_v6_)[default__.safeIndex(681, len(d_241_v6_))]) >= (d_238_v3_), d_236_globalState_)
                d_270_v33_: _dafny.Seq
                d_270_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfq"))
                d_270_v33_ = d_270_v33_
                d_271_v35_: _dafny.MultiSet
                d_271_v35_ = _dafny.MultiSet([_dafny.Set({d_233_v0_, d_233_v0_})])
                d_272_v36_: D28
                def iife21_():
                    coll15_ = _dafny.Map()
                    compr_15_: _dafny.Set
                    for compr_15_ in (d_271_v35_).Elements:
                        d_273_v34_: _dafny.Set = compr_15_
                        if (d_273_v34_) in (d_271_v35_):
                            coll15_[d_273_v34_] = d_238_v3_
                    return _dafny.Map(coll15_)
                d_272_v36_ = D28_DC68(iife21_()
)
                default__.m0(d_238_v3_, (d_259_v22_)[default__.safeIndex(len((d_272_v36_).cf116), len(d_259_v22_))], d_236_globalState_)
                d_274_v37_: _dafny.Array
                def lambda8_(d_275_v1_):
                    def lambda9_(d_276_i2_):
                        return d_275_v1_

                    return lambda9_

                init4_ = lambda8_(d_234_v1_)
                nw47_ = _dafny.Array(None, 9)
                for i0_4_ in range(nw47_.length(0)):
                    nw47_[i0_4_] = init4_(i0_4_)
                d_274_v37_ = nw47_
                d_277_v38_: C3
                nw48_ = C3()
                nw48_.ctor__(d_274_v37_, d_239_v4_)
                d_277_v38_ = nw48_
            d_238_v3_ = (d_238_v3_) + (d_238_v3_)
            if d_233_v0_:
                d_278_v39_: _dafny.Seq
                d_278_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvwankh"))
                d_278_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qf"))
                d_279_v40_: _dafny.Array
                nw49_ = _dafny.Array(None, 14)
                nw49_[int(0)] = d_234_v1_
                nw49_[int(1)] = d_234_v1_
                nw49_[int(2)] = d_234_v1_
                nw49_[int(3)] = d_234_v1_
                nw49_[int(4)] = d_234_v1_
                nw49_[int(5)] = d_234_v1_
                nw49_[int(6)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
                nw49_[int(7)] = d_234_v1_
                nw49_[int(8)] = d_234_v1_
                nw49_[int(9)] = d_234_v1_
                nw49_[int(10)] = d_234_v1_
                nw49_[int(11)] = d_234_v1_
                nw49_[int(12)] = default__.fm26(d_238_v3_, 944, d_236_globalState_)
                nw49_[int(13)] = d_234_v1_
                d_279_v40_ = nw49_
                d_280_v41_: C2
                nw50_ = C2()
                nw50_.ctor__(d_279_v40_, (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])
                d_280_v41_ = nw50_
                d_278_v39_ = d_278_v39_
                index32_ = default__.safeIndex(398, (d_235_v2_).length(0))
                (d_235_v2_)[index32_] = d_233_v0_
                d_281_v42_: T0
                nw51_ = C9()
                nw51_.ctor__(d_279_v40_, d_239_v4_)
                d_281_v42_ = nw51_
                nw52_ = C7()
                nw52_.ctor__()
                d_281_v42_ = nw52_
            elif True:
                d_282_v43_: _dafny.Map
                d_282_v43_ = _dafny.Map({d_233_v0_: True})
                d_283_v44_: _dafny.Set
                d_283_v44_ = _dafny.Set({d_282_v43_})
                d_284_v45_: _dafny.Seq
                d_284_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                d_285_v46_: _dafny.Map
                d_285_v46_ = _dafny.Map({(d_283_v44_ if not(d_233_v0_) else _dafny.Set({d_282_v43_})): d_284_v45_})
                d_285_v46_ = (d_285_v46_).set(default__.fm51(d_233_v0_, d_236_globalState_), (d_284_v45_) + (d_284_v45_))
                default__.m0(default__.safeModuloInt(d_238_v3_, 832), d_233_v0_, d_236_globalState_)
                d_286_v47_: _dafny.Map
                d_286_v47_ = _dafny.Map({d_238_v3_: 912})
                d_286_v47_ = d_286_v47_
                d_287_v48_: D5
                d_287_v48_ = D5_DC13(d_284_v45_, d_238_v3_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
                index33_ = default__.safeIndex(398, (d_235_v2_).length(0))
                rhs41_ = default__.fm14(d_236_globalState_)
                rhs42_ = (d_287_v48_).cf18
                rhs43_ = ((d_282_v43_)[d_233_v0_] if (d_233_v0_) in (d_282_v43_) else not((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]))
                rhs44_ = d_238_v3_
                rhs45_ = not((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
                lhs22_ = d_235_v2_
                lhs23_ = default__.safeIndex(398, (d_235_v2_).length(0))
                d_238_v3_ = rhs41_
                d_233_v0_ = rhs42_
                d_233_v0_ = rhs43_
                d_238_v3_ = rhs44_
                lhs22_[lhs23_] = rhs45_
                d_288_v49_: D27
                d_288_v49_ = D27_DC65(default__.fm1(d_238_v3_, d_238_v3_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, d_236_globalState_), (_dafny.SeqWithoutIsStrInference([(d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))] for d_289_i3_ in range(default__.abs(412))])) + (d_284_v45_), (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
                d_288_v49_ = D27_DC65((default__.fm0(d_238_v3_, 890, d_236_globalState_)) not in (d_284_v45_), (d_284_v45_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uia"))), (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], not(True))
            d_290_v50_: _dafny.Map
            d_290_v50_ = _dafny.Map({len(default__.fm34((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)): d_238_v3_})
            d_291_v51_: _dafny.Map
            d_291_v51_ = _dafny.Map({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]: d_290_v50_})
            d_292_v52_: C5
            nw53_ = C5()
            nw53_.ctor__()
            d_292_v52_ = nw53_
            d_293_v53_: _dafny.Seq
            d_293_v53_ = _dafny.SeqWithoutIsStrInference([d_292_v52_])
            d_294_v54_: D27
            d_294_v54_ = D27_DC64(d_293_v53_)
            rhs46_ = d_291_v51_
            rhs47_ = d_238_v3_
            rhs48_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nacylgo")))) <= (d_238_v3_)
            rhs49_ = d_294_v54_
            d_291_v51_ = rhs46_
            d_238_v3_ = rhs47_
            d_233_v0_ = rhs48_
            d_294_v54_ = rhs49_
        default__.m0(d_238_v3_, d_233_v0_, d_236_globalState_)
        d_295_v55_: _dafny.Seq
        d_295_v55_ = _dafny.SeqWithoutIsStrInference([d_233_v0_])
        hi3_ = len((d_295_v55_) + (d_295_v55_))
        for d_296_i4_ in range(d_238_v3_, hi3_):
            index34_ = default__.safeIndex(398, (d_235_v2_).length(0))
            (d_235_v2_)[index34_] = ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))] if d_233_v0_ else d_233_v0_)
            d_297_v56_: _dafny.Map
            d_297_v56_ = _dafny.Map({(d_235_v2_ if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))] else d_235_v2_): (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            d_297_v56_ = (d_297_v56_).set(d_235_v2_, True)
            index35_ = default__.safeIndex(218, (d_240_v5_).length(0))
            (d_240_v5_)[index35_] = d_239_v4_
            d_298_v57_: _dafny.Map
            d_298_v57_ = _dafny.Map({d_238_v3_: (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            d_299_v58_: _dafny.Array
            nw54_ = _dafny.Array(None, 9)
            nw54_[int(0)] = d_234_v1_
            nw54_[int(1)] = _dafny.Set({(D5_DC13(_dafny.SeqWithoutIsStrInference([d_239_v4_ for d_300_i5_ in range(default__.abs(582))]), len(d_298_v57_), d_233_v0_)).cf18, d_233_v0_})
            nw54_[int(2)] = d_234_v1_
            nw54_[int(3)] = d_234_v1_
            nw54_[int(4)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, d_233_v0_})
            nw54_[int(5)] = _dafny.Set({not(d_233_v0_)})
            nw54_[int(6)] = d_234_v1_
            nw54_[int(7)] = d_234_v1_
            nw54_[int(8)] = d_234_v1_
            d_299_v58_ = nw54_
            d_301_v59_: C3
            nw55_ = C3()
            nw55_.ctor__(d_299_v58_, d_239_v4_)
            d_301_v59_ = nw55_
        d_302_v60_: _dafny.Map
        d_302_v60_ = _dafny.Map({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]: d_241_v6_})
        default__.m0(len(((d_302_v60_).set((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], _dafny.SeqWithoutIsStrInference([313 for d_303_i6_ in range(default__.abs(629))]))) | ((d_302_v60_).set(d_233_v0_, d_241_v6_))), (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)
        d_304_v61_: _dafny.Map
        d_304_v61_ = _dafny.Map({-494: d_238_v3_})
        d_305_v62_: C6
        nw56_ = C6()
        nw56_.ctor__(d_304_v61_)
        d_305_v62_ = nw56_
        d_306_v63_: _dafny.Seq
        d_306_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "waupnh"))
        d_307_v64_: _dafny.Map
        d_307_v64_ = _dafny.Map({not((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]): d_233_v0_})
        hi4_ = default__.safeModuloInt(len(d_306_v63_), len(d_307_v64_))
        for d_308_i7_ in range((148) + (d_238_v3_), hi4_):
            d_309_v65_: C5
            nw57_ = C5()
            nw57_.ctor__()
            d_309_v65_ = nw57_
            d_310_v66_: _dafny.Seq
            d_310_v66_ = _dafny.SeqWithoutIsStrInference([d_309_v65_])
            d_311_v67_: D27
            d_311_v67_ = D27_DC64(d_310_v66_)
            d_311_v67_ = d_311_v67_
            if d_233_v0_:
                index36_ = default__.safeIndex(398, (d_235_v2_).length(0))
                (d_235_v2_)[index36_] = not ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]) or (d_233_v0_)
                d_312_v68_: _dafny.Map
                d_312_v68_ = _dafny.Map({default__.safeModuloInt(d_238_v3_, d_238_v3_): False})
                def iife22_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(424, 861):
                        d_313_v69_: int = compr_16_
                        if ((424) <= (d_313_v69_)) and ((d_313_v69_) < (861)):
                            coll16_[(d_313_v69_) + (d_238_v3_)] = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
                    return _dafny.Map(coll16_)
                d_312_v68_ = ((d_312_v68_).set(d_308_i7_, d_233_v0_)) | (iife22_()
                )
                d_314_v70_: _dafny.Seq
                d_314_v70_ = _dafny.SeqWithoutIsStrInference([d_307_v64_, d_307_v64_])
                d_315_v71_: _dafny.Map
                d_316_v72_: bool
                d_317_v73_: bool
                d_318_v74_: _dafny.Array
                out2_: _dafny.Map
                out3_: bool
                out4_: bool
                out5_: _dafny.Array
                out2_, out3_, out4_, out5_ = (d_305_v62_).m9(d_314_v70_, False, d_236_globalState_)
                d_315_v71_ = out2_
                d_316_v72_ = out3_
                d_317_v73_ = out4_
                d_318_v74_ = out5_
                d_319_v75_: _dafny.Map
                d_319_v75_ = _dafny.Map({(d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))]: d_308_i7_})
                d_319_v75_ = (d_319_v75_).set((d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))], (0) - (d_308_i7_))
                d_306_v63_ = d_306_v63_
            elif True:
                d_320_v76_: _dafny.Array
                def lambda10_(d_321_v63_):
                    def lambda11_(d_322_i8_):
                        return d_321_v63_

                    return lambda11_

                init5_ = lambda10_(d_306_v63_)
                nw58_ = _dafny.Array(None, 25)
                for i0_5_ in range(nw58_.length(0)):
                    nw58_[i0_5_] = init5_(i0_5_)
                d_320_v76_ = nw58_
                index37_ = default__.safeIndex(31, (d_320_v76_).length(0))
                (d_320_v76_)[index37_] = d_306_v63_
                d_323_v77_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Set({}), 26)
                d_323_v77_ = nw59_
                d_324_v78_: T1
                nw60_ = C1()
                nw60_.ctor__((d_323_v77_ if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))] else d_323_v77_), (d_306_v63_)[default__.safeIndex(d_308_i7_, len(d_306_v63_))])
                d_324_v78_ = nw60_
                d_325_v79_: C0
                nw61_ = C0()
                nw61_.ctor__()
                d_325_v79_ = nw61_
                d_326_v80_: _dafny.MultiSet
                d_326_v80_ = _dafny.MultiSet([136])
                d_327_v81_: C0
                d_327_v81_ = d_325_v79_
                index38_ = default__.safeIndex(31, (d_320_v76_).length(0))
                index39_ = default__.safeIndex(398, (d_235_v2_).length(0))
                rhs50_ = (d_306_v63_).set(default__.safeIndex((0) - (((d_326_v80_).cardinality) + (default__.fm13((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, d_233_v0_, d_236_globalState_))), len(d_306_v63_)), (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])
                rhs51_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rinkuq"))) < (d_306_v63_)
                rhs52_ = d_324_v78_
                rhs53_ = ((d_327_v81_ if d_233_v0_ else d_327_v81_))
                lhs24_ = d_320_v76_
                lhs25_ = default__.safeIndex(31, (d_320_v76_).length(0))
                lhs26_ = d_235_v2_
                lhs27_ = default__.safeIndex(398, (d_235_v2_).length(0))
                lhs24_[lhs25_] = rhs50_
                lhs26_[lhs27_] = rhs51_
                d_324_v78_ = rhs52_
                d_325_v79_ = rhs53_
                d_328_v82_: _dafny.Array
                nw62_ = _dafny.Array(_dafny.Map({}), 12)
                d_328_v82_ = nw62_
                d_329_v83_: _dafny.Array
                nw63_ = _dafny.Array(int(0), 5)
                d_329_v83_ = nw63_
                index40_ = default__.safeIndex(377, (d_328_v82_).length(0))
                (d_328_v82_)[index40_] = (_dafny.Map({d_329_v83_: len((d_320_v76_)[default__.safeIndex(31, (d_320_v76_).length(0))])})).set(d_329_v83_, d_308_i7_)
                d_330_v84_: T0
                nw64_ = C7()
                nw64_.ctor__()
                d_330_v84_ = nw64_
                d_331_v85_: _dafny.Map
                d_331_v85_ = _dafny.Map({d_330_v84_: d_238_v3_})
                index41_ = default__.safeIndex(377, (d_328_v82_).length(0))
                (d_328_v82_)[index41_] = _dafny.Map({d_329_v83_: default__.safeDivisionInt((0) - (len(d_331_v85_)), d_308_i7_)})
                d_332_v86_: _dafny.Map
                d_332_v86_ = _dafny.Map({d_238_v3_: (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
                d_332_v86_ = (d_332_v86_).set(d_238_v3_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
                d_333_v87_: C7
                nw65_ = C7()
                nw65_.ctor__()
                d_333_v87_ = nw65_
                d_233_v0_ = (d_233_v0_) or ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
            d_238_v3_ = 868
            d_334_v90_: _dafny.Array
            nw66_ = _dafny.Array(None, 23)
            nw66_[int(0)] = d_234_v1_
            nw66_[int(1)] = d_234_v1_
            nw66_[int(2)] = d_234_v1_
            nw66_[int(3)] = d_234_v1_
            nw66_[int(4)] = d_234_v1_
            nw66_[int(5)] = d_234_v1_
            nw66_[int(6)] = _dafny.Set({d_233_v0_})
            nw66_[int(7)] = d_234_v1_
            nw66_[int(8)] = d_234_v1_
            nw66_[int(9)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw66_[int(10)] = d_234_v1_
            nw66_[int(11)] = _dafny.Set({d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw66_[int(12)] = d_234_v1_
            nw66_[int(13)] = d_234_v1_
            nw66_[int(14)] = _dafny.Set({d_233_v0_})
            nw66_[int(15)] = d_234_v1_
            nw66_[int(16)] = d_234_v1_
            nw66_[int(17)] = d_234_v1_
            nw66_[int(18)] = _dafny.Set({d_233_v0_})
            nw66_[int(19)] = d_234_v1_
            nw66_[int(20)] = d_234_v1_
            nw66_[int(21)] = _dafny.Set({d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw66_[int(22)] = d_234_v1_
            d_334_v90_ = nw66_
            d_335_v91_: _dafny.Array
            nw67_ = _dafny.Array(None, 24)
            d_335_v91_ = nw67_
            d_336_v92_: D28
            d_336_v92_ = D28_DC69(d_238_v3_, d_335_v91_, d_239_v4_)
            d_337_v93_: T1
            nw68_ = C11()
            def iife23_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in ((d_305_v62_).f10).keys.Elements:
                    d_338_v88_: int = compr_17_
                    if (d_338_v88_) in ((d_305_v62_).f10):
                        def iife24_():
                            coll18_ = _dafny.Map()
                            compr_18_: int
                            for compr_18_ in (_dafny.SeqWithoutIsStrInference([d_238_v3_])).Elements:
                                d_339_v89_: int = compr_18_
                                if (d_339_v89_) in (_dafny.SeqWithoutIsStrInference([d_238_v3_])):
                                    coll18_[default__.safeModuloInt(d_339_v89_, d_238_v3_)] = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
                            return _dafny.Map(coll18_)
                        coll17_[(d_338_v88_) + (d_238_v3_)] = iife24_()

                return _dafny.Map(coll17_)
            nw68_.ctor__(iife23_()
            , d_308_i7_, d_334_v90_, (d_336_v92_).cf119)
            d_337_v93_ = nw68_
            d_340_v94_: _dafny.Map
            d_340_v94_ = _dafny.Map({d_337_v93_: (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            d_341_v95_: _dafny.Array
            nw69_ = _dafny.Array(None, 22)
            nw69_[int(0)] = _dafny.Set({False, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw69_[int(1)] = _dafny.Set({d_233_v0_})
            nw69_[int(2)] = d_234_v1_
            nw69_[int(3)] = d_234_v1_
            nw69_[int(4)] = d_234_v1_
            nw69_[int(5)] = d_234_v1_
            nw69_[int(6)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw69_[int(7)] = d_234_v1_
            nw69_[int(8)] = _dafny.Set({not(d_233_v0_), False})
            nw69_[int(9)] = d_234_v1_
            nw69_[int(10)] = default__.fm26(d_238_v3_, d_308_i7_, d_236_globalState_)
            nw69_[int(11)] = _dafny.Set({((d_340_v94_)[d_337_v93_] if (d_337_v93_) in (d_340_v94_) else d_233_v0_)})
            nw69_[int(12)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw69_[int(13)] = d_234_v1_
            nw69_[int(14)] = d_234_v1_
            nw69_[int(15)] = d_234_v1_
            nw69_[int(16)] = d_234_v1_
            nw69_[int(17)] = d_234_v1_
            nw69_[int(18)] = _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]})
            nw69_[int(19)] = _dafny.Set({d_233_v0_})
            nw69_[int(20)] = d_234_v1_
            nw69_[int(21)] = d_234_v1_
            d_341_v95_ = nw69_
            d_342_v96_: T1
            nw70_ = C4()
            nw70_.ctor__(d_308_i7_, d_238_v3_, (d_337_v93_).f4, (d_337_v93_).f5)
            d_342_v96_ = nw70_
            d_343_v97_: _dafny.Set
            d_343_v97_ = _dafny.Set({d_309_v65_, d_309_v65_})
            nw71_ = C10()
            nw71_.ctor__((d_308_i7_) * (len(d_241_v6_)), (_dafny.Set({d_309_v65_})) != (d_343_v97_), (d_337_v93_).f4, _dafny.CodePoint('l'))
            d_337_v93_ = nw71_
        d_344_v98_: _dafny.Seq
        d_344_v98_ = _dafny.SeqWithoutIsStrInference([d_235_v2_, d_235_v2_])
        d_345_v99_: D18
        d_345_v99_ = D18_DC45((d_344_v98_)[default__.safeIndex(d_238_v3_, len(d_344_v98_))], True, d_238_v3_, -410)
        d_238_v3_ = (0) - ((d_345_v99_).cf70)
        d_238_v3_ = default__.safeModuloInt(d_238_v3_, (0) - (default__.safeDivisionInt(d_238_v3_, len(d_306_v63_))))
        d_346_v100_: _dafny.Map
        d_346_v100_ = _dafny.Map({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]: 199})
        default__.m0(len((d_346_v100_).set((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_238_v3_)), (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)
        if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]:
            d_347_v101_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.Set({}), 15)
            d_347_v101_ = nw72_
            d_348_v102_: C9
            nw73_ = C9()
            nw73_.ctor__(d_347_v101_, d_239_v4_)
            d_348_v102_ = nw73_
            d_349_v103_: _dafny.Map
            d_349_v103_ = _dafny.Map({d_348_v102_: d_238_v3_})
            d_350_v104_: _dafny.Array
            nw74_ = _dafny.Array(None, 23)
            nw74_[int(0)] = d_234_v1_
            nw74_[int(1)] = d_234_v1_
            nw74_[int(2)] = d_234_v1_
            nw74_[int(3)] = d_234_v1_
            nw74_[int(4)] = d_234_v1_
            nw74_[int(5)] = d_234_v1_
            nw74_[int(6)] = d_234_v1_
            nw74_[int(7)] = d_234_v1_
            nw74_[int(8)] = d_234_v1_
            nw74_[int(9)] = d_234_v1_
            nw74_[int(10)] = d_234_v1_
            nw74_[int(11)] = d_234_v1_
            nw74_[int(12)] = d_234_v1_
            nw74_[int(13)] = d_234_v1_
            nw74_[int(14)] = d_234_v1_
            nw74_[int(15)] = default__.fm26(d_238_v3_, d_238_v3_, d_236_globalState_)
            nw74_[int(16)] = default__.fm26(((d_349_v103_)[d_348_v102_] if (d_348_v102_) in (d_349_v103_) else d_238_v3_), d_238_v3_, d_236_globalState_)
            nw74_[int(17)] = d_234_v1_
            nw74_[int(18)] = d_234_v1_
            nw74_[int(19)] = d_234_v1_
            nw74_[int(20)] = d_234_v1_
            nw74_[int(21)] = d_234_v1_
            nw74_[int(22)] = _dafny.Set({not((d_348_v102_).fm5(len(d_306_v63_), d_233_v0_, d_238_v3_, d_238_v3_, d_236_globalState_))})
            d_350_v104_ = nw74_
            d_351_v105_: T1
            nw75_ = C3()
            nw75_.ctor__(d_350_v104_, (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])
            d_351_v105_ = nw75_
            d_352_v106_: _dafny.Map
            d_352_v106_ = _dafny.Map({d_239_v4_: d_351_v105_})
            d_353_v107_: _dafny.Map
            d_353_v107_ = _dafny.Map({(((d_305_v62_).f10)[d_238_v3_] if (d_238_v3_) in ((d_305_v62_).f10) else len(d_352_v106_)): d_235_v2_})
            d_354_v108_: _dafny.Map
            d_354_v108_ = _dafny.Map({d_306_v63_: d_235_v2_})
            d_355_v109_: D2
            d_355_v109_ = D2_DC3(d_235_v2_)
            d_356_v110_: _dafny.Set
            d_356_v110_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([D7_DC17((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, False, d_233_v0_) for d_357_i9_ in range(default__.abs(-303))])})
            d_358_v111_: _dafny.Array
            def lambda12_(d_359_v3_):
                def lambda13_(d_360_i10_):
                    return (d_360_i10_) + (d_359_v3_)

                return lambda13_

            init6_ = lambda12_(d_238_v3_)
            nw76_ = _dafny.Array(None, 26)
            for i0_6_ in range(nw76_.length(0)):
                nw76_[i0_6_] = init6_(i0_6_)
            d_358_v111_ = nw76_
            d_361_v112_: _dafny.Map
            d_361_v112_ = _dafny.Map({d_358_v111_: d_238_v3_})
            d_362_v113_: _dafny.Array
            nw77_ = _dafny.Array(None, 22)
            nw77_[int(0)] = d_305_v62_
            nw77_[int(1)] = d_305_v62_
            nw77_[int(2)] = d_305_v62_
            nw77_[int(3)] = d_305_v62_
            nw77_[int(4)] = d_305_v62_
            nw77_[int(5)] = d_305_v62_
            nw77_[int(6)] = d_305_v62_
            nw77_[int(7)] = d_305_v62_
            nw77_[int(8)] = d_305_v62_
            nw77_[int(9)] = d_305_v62_
            nw77_[int(10)] = d_305_v62_
            nw77_[int(11)] = d_305_v62_
            nw77_[int(12)] = d_305_v62_
            nw77_[int(13)] = d_305_v62_
            nw77_[int(14)] = d_305_v62_
            nw77_[int(15)] = d_305_v62_
            nw77_[int(16)] = d_305_v62_
            nw77_[int(17)] = d_305_v62_
            nw77_[int(18)] = d_305_v62_
            nw77_[int(19)] = d_305_v62_
            nw77_[int(20)] = d_305_v62_
            nw77_[int(21)] = d_305_v62_
            d_362_v113_ = nw77_
            d_363_v114_: C5
            nw78_ = C5()
            nw78_.ctor__()
            d_363_v114_ = nw78_
            d_364_v115_: _dafny.Seq
            d_364_v115_ = _dafny.SeqWithoutIsStrInference([d_363_v114_])
            index42_ = default__.safeIndex(398, (d_235_v2_).length(0))
            rhs54_ = _dafny.Map({d_238_v3_: ((d_354_v108_)[d_306_v63_] if (d_306_v63_) in (d_354_v108_) else (d_355_v109_).cf4)})
            rhs55_ = (not(((0) - (d_238_v3_)) == (d_238_v3_)) if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))] else (d_238_v3_) != (d_238_v3_))
            rhs56_ = (((d_305_v62_).f10)[(len(d_356_v110_)) - (((d_361_v112_)[d_358_v111_] if (d_358_v111_) in (d_361_v112_) else d_238_v3_))] if ((len(d_356_v110_)) - (((d_361_v112_)[d_358_v111_] if (d_358_v111_) in (d_361_v112_) else d_238_v3_))) in ((d_305_v62_).f10) else (D28_DC69(d_238_v3_, d_362_v113_, (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])).cf117)
            rhs57_ = default__.safeModuloInt(-421, default__.safeModuloInt((0) - (d_238_v3_), len(d_364_v115_)))
            lhs28_ = d_235_v2_
            lhs29_ = default__.safeIndex(398, (d_235_v2_).length(0))
            d_353_v107_ = rhs54_
            lhs28_[lhs29_] = rhs55_
            d_238_v3_ = rhs56_
            d_238_v3_ = rhs57_
            d_365_v116_: _dafny.Array
            nw79_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_365_v116_ = nw79_
            index43_ = default__.safeIndex(89, (d_365_v116_).length(0))
            (d_365_v116_)[index43_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_366_i11_ in range(default__.abs(729))])) + (d_306_v63_)
            index44_ = default__.safeIndex(89, (d_365_v116_).length(0))
            (d_365_v116_)[index44_] = (d_306_v63_) + (d_306_v63_)
            d_367_v117_: C7
            nw80_ = C7()
            nw80_.ctor__()
            d_367_v117_ = nw80_
            index45_ = default__.safeIndex(398, (d_235_v2_).length(0))
            (d_235_v2_)[index45_] = d_233_v0_
            d_368_v118_: _dafny.Seq
            d_368_v118_ = _dafny.SeqWithoutIsStrInference([d_358_v111_])
            d_238_v3_ = (D11_DC25(not(not((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])), 558, (d_368_v118_)[default__.safeIndex(len(d_306_v63_), len(d_368_v118_))], d_238_v3_)).cf39
        elif True:
            index46_ = default__.safeIndex(398, (d_235_v2_).length(0))
            (d_235_v2_)[index46_] = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
            d_369_v119_: _dafny.Array
            nw81_ = _dafny.Array(_dafny.Set({}), 14)
            d_369_v119_ = nw81_
            d_370_v120_: C9
            nw82_ = C9()
            nw82_.ctor__(d_369_v119_, (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])
            d_370_v120_ = nw82_
            d_371_v121_: _dafny.Seq
            d_371_v121_ = _dafny.SeqWithoutIsStrInference([d_370_v120_, d_370_v120_])
            d_372_v122_: _dafny.Array
            nw83_ = _dafny.Array(None, 6)
            nw83_[int(0)] = d_370_v120_
            nw83_[int(1)] = d_370_v120_
            nw83_[int(2)] = d_370_v120_
            nw83_[int(3)] = (d_371_v121_)[default__.safeIndex(d_238_v3_, len(d_371_v121_))]
            nw83_[int(4)] = d_370_v120_
            nw83_[int(5)] = d_370_v120_
            d_372_v122_ = nw83_
            index47_ = default__.safeIndex(800, (d_372_v122_).length(0))
            (d_372_v122_)[index47_] = d_370_v120_
            index48_ = default__.safeIndex(800, (d_372_v122_).length(0))
            (d_372_v122_)[index48_] = d_370_v120_
            d_238_v3_ = -721
            d_234_v1_ = d_234_v1_
            d_373_v123_: _dafny.MultiSet
            d_373_v123_ = _dafny.MultiSet([(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_, d_233_v0_, d_233_v0_])
            index49_ = default__.safeIndex(398, (d_235_v2_).length(0))
            (d_235_v2_)[index49_] = ((0) - (d_238_v3_)) >= ((d_373_v123_).cardinality)
        d_374_v124_: _dafny.Map
        d_374_v124_ = _dafny.Map({(d_241_v6_) + (_dafny.SeqWithoutIsStrInference([d_238_v3_, d_238_v3_, d_238_v3_, d_238_v3_])): d_238_v3_})
        d_374_v124_ = (d_374_v124_).set(d_241_v6_, d_238_v3_)
        d_346_v100_ = d_346_v100_
        source6_ = D1_DC2(d_238_v3_, ((d_241_v6_)[default__.safeIndex(-276, len(d_241_v6_))]) + (d_238_v3_))
        if source6_.is_DC2:
            d_375___mcc_h0_ = source6_.cf2
            d_376___mcc_h1_ = source6_.cf3
            d_377_cf3_ = d_376___mcc_h1_
            d_378_cf2_ = d_375___mcc_h0_
            if (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]:
                d_379_v125_: _dafny.MultiSet
                d_379_v125_ = _dafny.MultiSet([(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (len(d_241_v6_)) != (340), (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], not((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])])
                d_379_v125_ = d_379_v125_
                index50_ = default__.safeIndex(218, (d_240_v5_).length(0))
                (d_240_v5_)[index50_] = d_239_v4_
                default__.m0(962, d_233_v0_, d_236_globalState_)
                d_380_v126_: _dafny.Array
                def lambda14_(d_381_cf3_):
                    def lambda15_(d_382_i12_):
                        return default__.safeModuloInt(d_382_i12_, d_381_cf3_)

                    return lambda15_

                init7_ = lambda14_(d_377_cf3_)
                nw84_ = _dafny.Array(None, 1)
                for i0_7_ in range(nw84_.length(0)):
                    nw84_[i0_7_] = init7_(i0_7_)
                d_380_v126_ = nw84_
                d_383_v127_: _dafny.Set
                d_383_v127_ = _dafny.Set({len(_dafny.Map({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]: (d_379_v125_).cardinality}))})
                index51_ = default__.safeIndex(934, (d_380_v126_).length(0))
                (d_380_v126_)[index51_] = len(d_383_v127_)
                d_384_v128_: _dafny.MultiSet
                d_384_v128_ = _dafny.MultiSet([d_380_v126_])
                index52_ = default__.safeIndex(934, (d_380_v126_).length(0))
                rhs58_ = (((d_384_v128_) | (d_384_v128_)) | (d_384_v128_)).cardinality
                rhs59_ = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
                rhs60_ = (d_378_cf2_) - ((-812) + (d_378_cf2_))
                lhs30_ = d_380_v126_
                lhs31_ = default__.safeIndex(934, (d_380_v126_).length(0))
                d_378_cf2_ = rhs58_
                d_233_v0_ = rhs59_
                lhs30_[lhs31_] = rhs60_
                d_385_v129_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.MultiSet({}), 24)
                d_385_v129_ = nw85_
                index53_ = default__.safeIndex(734, (d_385_v129_).length(0))
                (d_385_v129_)[index53_] = d_379_v125_
                index54_ = default__.safeIndex(734, (d_385_v129_).length(0))
                (d_385_v129_)[index54_] = d_379_v125_
            elif True:
                d_386_v130_: C0
                nw86_ = C0()
                nw86_.ctor__()
                d_386_v130_ = nw86_
                d_387_v131_: C0
                d_387_v131_ = d_386_v130_
                d_387_v131_ = d_386_v130_
                d_388_v132_: _dafny.Array
                def lambda16_(d_389_v1_):
                    def lambda17_(d_390_i13_):
                        return d_389_v1_

                    return lambda17_

                init8_ = lambda16_(d_234_v1_)
                nw87_ = _dafny.Array(None, 4)
                for i0_8_ in range(nw87_.length(0)):
                    nw87_[i0_8_] = init8_(i0_8_)
                d_388_v132_ = nw87_
                d_391_v133_: C10
                nw88_ = C10()
                nw88_.ctor__(d_377_cf3_, d_233_v0_, d_388_v132_, (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])
                d_391_v133_ = nw88_
                d_392_v134_: D30
                d_392_v134_ = D30_DC72(d_391_v133_)
                d_393_v135_: _dafny.Array
                nw89_ = _dafny.Array(None, 27)
                nw89_[int(0)] = (d_392_v134_).cf122
                nw89_[int(1)] = d_391_v133_
                nw89_[int(2)] = d_391_v133_
                nw89_[int(3)] = d_391_v133_
                nw89_[int(4)] = d_391_v133_
                nw89_[int(5)] = d_391_v133_
                nw89_[int(6)] = d_391_v133_
                nw89_[int(7)] = d_391_v133_
                nw89_[int(8)] = d_391_v133_
                nw89_[int(9)] = d_391_v133_
                nw89_[int(10)] = d_391_v133_
                nw89_[int(11)] = d_391_v133_
                nw89_[int(12)] = d_391_v133_
                nw89_[int(13)] = d_391_v133_
                nw89_[int(14)] = d_391_v133_
                nw89_[int(15)] = d_391_v133_
                nw89_[int(16)] = d_391_v133_
                nw89_[int(17)] = d_391_v133_
                nw89_[int(18)] = d_391_v133_
                nw89_[int(19)] = d_391_v133_
                nw89_[int(20)] = d_391_v133_
                nw89_[int(21)] = d_391_v133_
                nw89_[int(22)] = d_391_v133_
                nw89_[int(23)] = d_391_v133_
                nw89_[int(24)] = d_391_v133_
                nw89_[int(25)] = d_391_v133_
                nw89_[int(26)] = d_391_v133_
                d_393_v135_ = nw89_
                index55_ = default__.safeIndex(571, (d_393_v135_).length(0))
                (d_393_v135_)[index55_] = d_391_v133_
                index56_ = default__.safeIndex(571, (d_393_v135_).length(0))
                rhs61_ = d_391_v133_
                rhs62_ = d_378_cf2_
                rhs63_ = d_374_v124_
                lhs32_ = d_393_v135_
                lhs33_ = default__.safeIndex(571, (d_393_v135_).length(0))
                lhs32_[lhs33_] = rhs61_
                d_378_cf2_ = rhs62_
                d_374_v124_ = rhs63_
                d_394_v136_: T0
                nw90_ = C9()
                nw90_.ctor__(d_388_v132_, d_239_v4_)
                d_394_v136_ = nw90_
                d_394_v136_ = d_394_v136_
                d_395_v137_: D8
                d_395_v137_ = D8_DC19(d_377_cf3_, d_233_v0_)
                d_396_v138_: _dafny.Map
                d_396_v138_ = _dafny.Map({(d_395_v137_ if (d_391_v133_).f9 else d_395_v137_): (d_391_v133_).f8})
                d_397_v139_: D6
                d_397_v139_ = D6_DC14(d_241_v6_)
                index57_ = default__.safeIndex(398, (d_235_v2_).length(0))
                index58_ = default__.safeIndex(218, (d_240_v5_).length(0))
                rhs64_ = (d_295_v55_) + ((d_295_v55_).set(default__.safeIndex(d_377_cf3_, len(d_295_v55_)), False))
                rhs65_ = ((d_391_v133_).f8) in ((_dafny.SeqWithoutIsStrInference([716, d_378_cf2_])) + ((d_397_v139_).cf19))
                rhs66_ = (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))]
                rhs67_ = d_238_v3_
                rhs68_ = (default__.fm52((d_391_v133_).f8, 973, d_238_v3_, (0) - ((d_391_v133_).f8), d_236_globalState_)) | (d_396_v138_)
                lhs34_ = d_235_v2_
                lhs35_ = default__.safeIndex(398, (d_235_v2_).length(0))
                lhs36_ = d_240_v5_
                lhs37_ = default__.safeIndex(218, (d_240_v5_).length(0))
                d_295_v55_ = rhs64_
                lhs34_[lhs35_] = rhs65_
                lhs36_[lhs37_] = rhs66_
                d_378_cf2_ = rhs67_
                d_396_v138_ = rhs68_
                d_398_v140_: _dafny.Set
                d_398_v140_ = _dafny.Set({default__.fm26((0) - (-20), (d_391_v133_).f8, d_236_globalState_), d_234_v1_, d_234_v1_, d_234_v1_})
                d_399_v141_: D18
                d_399_v141_ = D18_DC44(d_398_v140_)
                d_399_v141_ = D18_DC44((d_398_v140_).intersection(d_398_v140_))
            d_306_v63_ = d_306_v63_
            d_378_cf2_ = (0) - ((514) * (len(d_241_v6_)))
            if ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]) or (d_233_v0_):
                d_377_cf3_ = default__.fm13(d_233_v0_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)
                default__.m0(d_238_v3_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)
                d_400_v142_: C7
                nw91_ = C7()
                nw91_.ctor__()
                d_400_v142_ = nw91_
                index59_ = default__.safeIndex(398, (d_235_v2_).length(0))
                rhs69_ = (0) - ((d_378_cf2_) - (len(d_295_v55_)))
                rhs70_ = (d_239_v4_ if d_233_v0_ else default__.fm0(d_238_v3_, len(_dafny.Map({d_233_v0_: not(False)})), d_236_globalState_))
                rhs71_ = d_233_v0_
                lhs38_ = d_235_v2_
                lhs39_ = default__.safeIndex(398, (d_235_v2_).length(0))
                d_377_cf3_ = rhs69_
                d_239_v4_ = rhs70_
                lhs38_[lhs39_] = rhs71_
                d_401_v143_: _dafny.Array
                def lambda18_(d_402_i14_):
                    return default__.safeDivisionInt(d_402_i14_, 724)

                init9_ = lambda18_
                nw92_ = _dafny.Array(None, 21)
                for i0_9_ in range(nw92_.length(0)):
                    nw92_[i0_9_] = init9_(i0_9_)
                d_401_v143_ = nw92_
                index60_ = default__.safeIndex(373, (d_401_v143_).length(0))
                (d_401_v143_)[index60_] = d_378_cf2_
                d_403_v144_: _dafny.MultiSet
                d_403_v144_ = _dafny.MultiSet([d_378_cf2_])
                index61_ = default__.safeIndex(373, (d_401_v143_).length(0))
                (d_401_v143_)[index61_] = (((d_403_v144_) | (d_403_v144_)) - ((d_403_v144_) - (d_403_v144_))).cardinality
            elif True:
                d_306_v63_ = d_306_v63_
                d_404_v145_: C8
                nw93_ = C8()
                nw93_.ctor__()
                d_404_v145_ = nw93_
                d_404_v145_ = d_404_v145_
                d_405_v146_: T0
                nw94_ = C7()
                nw94_.ctor__()
                d_405_v146_ = nw94_
                d_406_v147_: C7
                nw95_ = C7()
                nw95_.ctor__()
                d_406_v147_ = nw95_
                d_407_v148_: _dafny.Array
                nw96_ = _dafny.Array(_dafny.Seq({}), 3)
                d_407_v148_ = nw96_
                d_408_v149_: _dafny.Set
                d_408_v149_ = _dafny.Set({d_235_v2_, d_235_v2_})
                d_409_v150_: _dafny.Set
                d_409_v150_ = d_408_v149_
                d_410_v151_: _dafny.Map
                d_410_v151_ = _dafny.Map({d_378_cf2_: d_233_v0_})
                d_411_v154_: _dafny.Array
                nw97_ = _dafny.Array(None, 22)
                nw97_[int(0)] = d_410_v151_
                nw97_[int(1)] = d_410_v151_
                nw97_[int(2)] = d_410_v151_
                nw97_[int(3)] = d_410_v151_
                nw97_[int(4)] = d_410_v151_
                nw97_[int(5)] = d_410_v151_
                nw97_[int(6)] = d_410_v151_
                nw97_[int(7)] = (d_410_v151_).set(702, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])
                nw97_[int(8)] = (d_410_v151_).set((0) - (d_238_v3_), not(d_233_v0_))
                nw97_[int(9)] = d_410_v151_
                nw97_[int(10)] = d_410_v151_
                nw97_[int(11)] = d_410_v151_
                nw97_[int(12)] = d_410_v151_
                nw97_[int(13)] = _dafny.Map({(default__.fm11(d_236_globalState_)).cardinality: d_233_v0_})
                nw97_[int(14)] = d_410_v151_
                nw97_[int(15)] = _dafny.Map({d_238_v3_: d_233_v0_})
                def iife25_():
                    coll19_ = _dafny.Map()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(976, 470):
                        d_412_v152_: int = compr_19_
                        if ((976) <= (d_412_v152_)) and ((d_412_v152_) < (470)):
                            coll19_[(d_412_v152_) + (d_377_cf3_)] = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
                    return _dafny.Map(coll19_)
                nw97_[int(16)] = iife25_()
                
                nw97_[int(17)] = d_410_v151_
                nw97_[int(18)] = d_410_v151_
                nw97_[int(19)] = d_410_v151_
                def iife26_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(410, 813):
                        d_413_v153_: int = compr_20_
                        if ((410) <= (d_413_v153_)) and ((d_413_v153_) < (813)):
                            coll20_[(d_413_v153_) + (d_377_cf3_)] = d_233_v0_
                    return _dafny.Map(coll20_)
                nw97_[int(20)] = iife26_()
                
                nw97_[int(21)] = d_410_v151_
                d_411_v154_ = nw97_
                d_414_v155_: D4
                d_414_v155_ = D4_DC10(d_409_v150_, d_377_cf3_, d_306_v63_, d_411_v154_)
                rhs72_ = (d_414_v155_).cf13
                rhs73_ = d_377_cf3_
                rhs74_ = d_405_v146_
                rhs75_ = d_406_v147_
                rhs76_ = d_407_v148_
                d_306_v63_ = rhs72_
                d_377_cf3_ = rhs73_
                d_405_v146_ = rhs74_
                d_406_v147_ = rhs75_
                d_407_v148_ = rhs76_
                d_415_v156_: C0
                nw98_ = C0()
                nw98_.ctor__()
                d_415_v156_ = nw98_
                d_416_v157_: _dafny.Map
                d_416_v157_ = _dafny.Map({((d_346_v100_)[d_233_v0_] if (d_233_v0_) in (d_346_v100_) else d_377_cf3_): _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqvlgxso"))): not(d_233_v0_)})})
                d_417_v158_: _dafny.Array
                nw99_ = _dafny.Array(_dafny.Set({}), 6)
                d_417_v158_ = nw99_
                d_418_v159_: D20
                d_418_v159_ = D20_DC47(d_417_v158_)
                d_419_v160_: C11
                nw100_ = C11()
                nw100_.ctor__(d_416_v157_, d_377_cf3_, (d_418_v159_).cf73, (d_240_v5_)[default__.safeIndex(218, (d_240_v5_).length(0))])
                d_419_v160_ = nw100_
        elif True:
            d_420___mcc_h2_ = source6_.cf1
            d_421_cf1_ = d_420___mcc_h2_
            d_233_v0_ = not (not ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]) or ((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))])) or ((_dafny.SeqWithoutIsStrInference([len((d_306_v63_).set(default__.safeIndex(d_238_v3_, len(d_306_v63_)), d_239_v4_))])) != (_dafny.SeqWithoutIsStrInference([d_238_v3_ for d_422_i15_ in range(default__.abs(-668))])))
            d_423_v161_: _dafny.MultiSet
            d_423_v161_ = _dafny.MultiSet([d_238_v3_, d_238_v3_])
            d_424_v162_: D6
            d_424_v162_ = D6_DC14(d_241_v6_)
            d_423_v161_ = ((d_423_v161_) | (_dafny.MultiSet((d_424_v162_).cf19))).set(len(d_295_v55_), default__.abs(default__.fm13(d_233_v0_, not(default__.fm1(d_238_v3_, d_238_v3_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)), default__.fm1(d_238_v3_, d_238_v3_, (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_), d_236_globalState_)))
            d_425_v163_: _dafny.Map
            d_425_v163_ = _dafny.Map({d_239_v4_: d_238_v3_})
            default__.m0(len(d_425_v163_), (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_236_globalState_)
            d_426_v164_: _dafny.Map
            d_426_v164_ = _dafny.Map({d_233_v0_: default__.fm53(d_236_globalState_)})
            d_427_v165_: _dafny.Array
            nw101_ = _dafny.Array(None, 11)
            nw101_[int(0)] = d_238_v3_
            nw101_[int(1)] = d_238_v3_
            nw101_[int(2)] = d_238_v3_
            nw101_[int(3)] = len(d_426_v164_)
            nw101_[int(4)] = d_238_v3_
            nw101_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gy")))
            nw101_[int(6)] = d_238_v3_
            nw101_[int(7)] = d_238_v3_
            nw101_[int(8)] = (_dafny.MultiSet([(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]])).cardinality
            nw101_[int(9)] = 670
            nw101_[int(10)] = d_238_v3_
            d_427_v165_ = nw101_
            d_428_v166_: _dafny.Map
            d_428_v166_ = _dafny.Map({d_295_v55_: d_427_v165_})
            d_428_v166_ = (d_428_v166_).set(d_295_v55_, d_427_v165_)
        d_429_v167_: D14
        d_429_v167_ = D14_DC31(d_238_v3_)
        d_430_v168_: D14
        d_430_v168_ = D14_DC33(d_429_v167_)
        pat_let_tv3_ = d_238_v3_
        pat_let_tv4_ = d_307_v64_
        pat_let_tv5_ = d_235_v2_
        pat_let_tv6_ = d_235_v2_
        pat_let_tv7_ = d_235_v2_
        pat_let_tv8_ = d_235_v2_
        pat_let_tv9_ = d_307_v64_
        pat_let_tv10_ = d_238_v3_
        pat_let_tv11_ = d_307_v64_
        pat_let_tv12_ = d_235_v2_
        pat_let_tv13_ = d_235_v2_
        pat_let_tv14_ = d_235_v2_
        pat_let_tv15_ = d_235_v2_
        pat_let_tv16_ = d_307_v64_
        pat_let_tv17_ = d_233_v0_
        pat_let_tv18_ = d_238_v3_
        def lambda19_(source8_):
            if source8_.is_DC31:
                d_431___mcc_h6_ = source8_.cf48
                d_432_cf48_ = d_431___mcc_h6_
                return D15_DC35(d_432_cf48_)
            elif source8_.is_DC32:
                d_433___mcc_h7_ = source8_.cf49
                d_434___mcc_h8_ = source8_.cf50
                d_435___mcc_h9_ = source8_.cf51
                d_436_cf51_ = d_435___mcc_h9_
                d_437_cf50_ = d_434___mcc_h8_
                d_438_cf49_ = d_433___mcc_h7_
                return D15_DC35(pat_let_tv3_)
            elif source8_.is_DC30:
                d_439___mcc_h10_ = source8_.cf47
                d_440_cf47_ = d_439___mcc_h10_
                return D15_DC35(len((_dafny.SeqWithoutIsStrInference([((pat_let_tv4_)[(pat_let_tv6_)[default__.safeIndex(398, (pat_let_tv5_).length(0))]] if ((pat_let_tv8_)[default__.safeIndex(398, (pat_let_tv7_).length(0))]) in (pat_let_tv9_) else False)])).set(default__.safeIndex(pat_let_tv10_, len(_dafny.SeqWithoutIsStrInference([((pat_let_tv11_)[(pat_let_tv13_)[default__.safeIndex(398, (pat_let_tv12_).length(0))]] if ((pat_let_tv15_)[default__.safeIndex(398, (pat_let_tv14_).length(0))]) in (pat_let_tv16_) else False)]))), pat_let_tv17_)))
            elif True:
                d_441___mcc_h11_ = source8_.cf52
                d_442_cf52_ = d_441___mcc_h11_
                return D15_DC35(pat_let_tv18_)

        source7_ = lambda19_(d_430_v168_)
        if source7_.is_DC35:
            d_443___mcc_h3_ = source7_.cf54
            d_444_cf54_ = d_443___mcc_h3_
            d_445_v169_: _dafny.MultiSet
            d_445_v169_ = _dafny.MultiSet([(_dafny.MultiSet(d_241_v6_)).cardinality, (((d_305_v62_).f10)[d_444_cf54_] if (d_444_cf54_) in ((d_305_v62_).f10) else len(d_306_v63_))])
            d_446_v170_: _dafny.MultiSet
            d_446_v170_ = _dafny.MultiSet([d_444_cf54_, d_238_v3_, ((d_445_v169_)[d_238_v3_] if (d_238_v3_) in (d_445_v169_) else d_238_v3_)])
            d_447_v171_: _dafny.Seq
            d_447_v171_ = _dafny.SeqWithoutIsStrInference([d_446_v170_, d_446_v170_])
            d_448_v172_: _dafny.Array
            nw102_ = _dafny.Array(None, 4)
            nw102_[int(0)] = ((d_447_v171_)[default__.safeIndex(d_238_v3_, len(d_447_v171_))]).cardinality
            nw102_[int(1)] = d_444_cf54_
            nw102_[int(2)] = default__.safeModuloInt(len(d_306_v63_), d_444_cf54_)
            nw102_[int(3)] = d_238_v3_
            d_448_v172_ = nw102_
            index62_ = default__.safeIndex(206, (d_448_v172_).length(0))
            (d_448_v172_)[index62_] = d_238_v3_
            index63_ = default__.safeIndex(206, (d_448_v172_).length(0))
            (d_448_v172_)[index63_] = (d_444_cf54_) * ((0) - (default__.safeDivisionInt((0) - (((d_346_v100_)[d_233_v0_] if (d_233_v0_) in (d_346_v100_) else (d_241_v6_)[default__.safeIndex(d_444_cf54_, len(d_241_v6_))])), d_444_cf54_)))
            index64_ = default__.safeIndex(206, (d_448_v172_).length(0))
            rhs77_ = 718
            rhs78_ = d_295_v55_
            lhs40_ = d_448_v172_
            lhs41_ = default__.safeIndex(206, (d_448_v172_).length(0))
            lhs40_[lhs41_] = rhs77_
            d_295_v55_ = rhs78_
            d_449_v173_: _dafny.Seq
            d_449_v173_ = _dafny.SeqWithoutIsStrInference([d_234_v1_, _dafny.Set({(d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], d_233_v0_}), d_234_v1_])
            d_450_v174_: _dafny.Seq
            d_450_v174_ = _dafny.SeqWithoutIsStrInference([d_241_v6_, d_241_v6_])
            d_451_v175_: _dafny.MultiSet
            d_451_v175_ = _dafny.MultiSet([(d_450_v174_) != (d_450_v174_), not((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]), d_233_v0_])
            d_452_v176_: D2
            d_452_v176_ = D2_DC5(d_306_v63_)
            d_453_v177_: _dafny.MultiSet
            d_453_v177_ = _dafny.MultiSet([D2_DC5(d_306_v63_), d_452_v176_, d_452_v176_, d_452_v176_])
            index65_ = default__.safeIndex(398, (d_235_v2_).length(0))
            rhs79_ = d_444_cf54_
            rhs80_ = d_449_v173_
            rhs81_ = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
            rhs82_ = ((0) - (d_238_v3_)) >= ((d_448_v172_)[default__.safeIndex(206, (d_448_v172_).length(0))])
            rhs83_ = ((d_451_v175_).set((d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))], default__.abs((d_453_v177_).cardinality))) | (d_451_v175_)
            lhs42_ = d_235_v2_
            lhs43_ = default__.safeIndex(398, (d_235_v2_).length(0))
            d_444_cf54_ = rhs79_
            d_449_v173_ = rhs80_
            lhs42_[lhs43_] = rhs81_
            d_233_v0_ = rhs82_
            d_451_v175_ = rhs83_
            index66_ = default__.safeIndex(398, (d_235_v2_).length(0))
            (d_235_v2_)[index66_] = (d_235_v2_)[default__.safeIndex(398, (d_235_v2_).length(0))]
        elif source7_.is_DC34:
            d_454___mcc_h4_ = source7_.cf53
            d_455_cf53_ = d_454___mcc_h4_
            d_238_v3_ = 748
            (d_236_globalState_).f3 = d_235_v2_
            index67_ = default__.safeIndex(398, (d_235_v2_).length(0))
            rhs84_ = d_305_v62_
            rhs85_ = ((d_306_v63_) <= (default__.fm12(d_236_globalState_))) == ((_dafny.CodePoint('d')) in (d_306_v63_))
            rhs86_ = d_238_v3_
            rhs87_ = d_306_v63_
            lhs44_ = d_235_v2_
            lhs45_ = default__.safeIndex(398, (d_235_v2_).length(0))
            d_305_v62_ = rhs84_
            lhs44_[lhs45_] = rhs85_
            d_238_v3_ = rhs86_
            d_306_v63_ = rhs87_
            d_456_v178_: C7
            nw103_ = C7()
            nw103_.ctor__()
            d_456_v178_ = nw103_
        elif True:
            d_457___mcc_h5_ = source7_.cf55
            d_458_cf55_ = d_457___mcc_h5_
            d_238_v3_ = (d_238_v3_) + ((((_dafny.MultiSet([d_233_v0_])).set(d_233_v0_, default__.abs(d_238_v3_))).cardinality) - (d_238_v3_))
            d_459_v179_: _dafny.Seq
            d_459_v179_ = _dafny.SeqWithoutIsStrInference([d_307_v64_])
            d_460_v180_: _dafny.Map
            d_461_v181_: bool
            d_462_v182_: bool
            d_463_v183_: _dafny.Array
            out6_: _dafny.Map
            out7_: bool
            out8_: bool
            out9_: _dafny.Array
            out6_, out7_, out8_, out9_ = (d_305_v62_).m9((d_459_v179_) + (d_459_v179_), not((d_295_v55_)[default__.safeIndex(d_238_v3_, len(d_295_v55_))]), d_236_globalState_)
            d_460_v180_ = out6_
            d_461_v181_ = out7_
            d_462_v182_ = out8_
            d_463_v183_ = out9_
            d_464_v184_: _dafny.Array
            nw104_ = _dafny.Array(int(0), 7)
            d_464_v184_ = nw104_
            d_465_v185_: _dafny.Seq
            d_465_v185_ = _dafny.SeqWithoutIsStrInference([d_464_v184_])
            d_465_v185_ = d_465_v185_
            d_462_v182_ = d_233_v0_
        _dafny.print(_dafny.string_of(d_233_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v1_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_235_v2_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_236_globalState_).f0) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_globalState_.f2) == (_dafny.SeqWithoutIsStrInference([1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_globalState_.f3)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_238_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_239_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v6_) == (_dafny.SeqWithoutIsStrInference([-263]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v55_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v60_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([-263])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v61_) == (_dafny.Map({-494: -4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_305_v62_).f10) == (_dafny.Map({-494: -4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_306_v63_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_307_v64_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_344_v98_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_345_v99_).cf68)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v99_).cf69))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v99_).cf70))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v99_).cf71))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_346_v100_) == (_dafny.Map({False: 199}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v124_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([-263, -721, -721, -721, -721]): -721, _dafny.SeqWithoutIsStrInference([-263]): -721}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_429_v167_).cf48))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_430_v168_).cf52).cf48))
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
        return lambda: D1_DC2(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({self.cf6.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC8(D3, NamedTuple('DC8', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(_dafny.Set({}), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {self.cf13.VerbatimString(True)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC12(D5, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({self.cf16.VerbatimString(True)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC15(D6, NamedTuple('DC15', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC17(False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC19(D8, NamedTuple('DC19', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC20(D9, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC22(D10, NamedTuple('DC22', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({self.cf33.VerbatimString(True)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25(False, int(0), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)

class D11_DC25(D11, NamedTuple('DC25', [('cf38', Any), ('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)

class D12_DC27(D12, NamedTuple('DC27', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC29(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)

class D13_DC29(D13, NamedTuple('DC29', [('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC28(D13, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC31(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D14_DC33)

class D14_DC31(D14, NamedTuple('DC31', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC32(D14, NamedTuple('DC32', [('cf49', Any), ('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC30(D14, NamedTuple('DC30', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC33(D14, NamedTuple('DC33', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC33({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC33) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC35(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D15_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)

class D15_DC35(D15, NamedTuple('DC35', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC36(D15, NamedTuple('DC36', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D16_DC37)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)

class D16_DC38(D16, NamedTuple('DC38', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({self.cf57.VerbatimString(True)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC37(D16, NamedTuple('DC37', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC41(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)

class D17_DC41(D17, NamedTuple('DC41', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC42(D17, NamedTuple('DC42', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC43(D17, NamedTuple('DC43', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC45(_dafny.Array(None, 0), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D18_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)

class D18_DC45(D18, NamedTuple('DC45', [('cf68', Any), ('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC45({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC45) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)

class D19_DC46(D19, NamedTuple('DC46', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC48(False, _dafny.Seq({}), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D20_DC47)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)

class D20_DC48(D20, NamedTuple('DC48', [('cf74', Any), ('cf75', Any), ('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC47(D20, NamedTuple('DC47', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC47({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC47) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC49(D20, NamedTuple('DC49', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC51(False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D21_DC50)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)

class D21_DC51(D21, NamedTuple('DC51', [('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC50(D21, NamedTuple('DC50', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC50({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC50) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC52(D21, NamedTuple('DC52', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)

class D22_DC53(D22, NamedTuple('DC53', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC55(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D23_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D23_DC54)

class D23_DC55(D23, NamedTuple('DC55', [('cf85', Any), ('cf86', Any), ('cf87', Any), ('cf88', Any), ('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({self.cf85.VerbatimString(True)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC56(D23, NamedTuple('DC56', [('cf90', Any), ('cf91', Any), ('cf92', Any), ('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC56({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC56) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC57(D23, NamedTuple('DC57', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC54(D23, NamedTuple('DC54', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC54({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC54) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D24_DC58)

class D24_DC58(D24, NamedTuple('DC58', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC58({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC58) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC60(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D25_DC60)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D25_DC59)

class D25_DC60(D25, NamedTuple('DC60', [('cf96', Any), ('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC60({_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)}, {self.cf98.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC60) and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC59(D25, NamedTuple('DC59', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC59({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC59) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC62(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D26_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D26_DC63)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D26_DC61)

class D26_DC62(D26, NamedTuple('DC62', [('cf100', Any), ('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC62({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC62) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC63(D26, NamedTuple('DC63', [('cf102', Any), ('cf103', Any), ('cf104', Any), ('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC63({_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC63) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104 and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC61(D26, NamedTuple('DC61', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC61({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC61) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC65(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D27_DC65)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D27_DC66)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D27_DC67)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D27_DC64)

class D27_DC65(D27, NamedTuple('DC65', [('cf107', Any), ('cf108', Any), ('cf109', Any), ('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC65({_dafny.string_of(self.cf107)}, {self.cf108.VerbatimString(True)}, {_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC65) and self.cf107 == __o.cf107 and self.cf108 == __o.cf108 and self.cf109 == __o.cf109 and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC66(D27, NamedTuple('DC66', [('cf111', Any), ('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC66({_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC66) and self.cf111 == __o.cf111 and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC67(D27, NamedTuple('DC67', [('cf113', Any), ('cf114', Any), ('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC67({_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)}, {_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC67) and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC64(D27, NamedTuple('DC64', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC64({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC64) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC69(int(0), _dafny.Array(None, 0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D28_DC69)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D28_DC68)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D28_DC70)

class D28_DC69(D28, NamedTuple('DC69', [('cf117', Any), ('cf118', Any), ('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC69({_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC69) and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC68(D28, NamedTuple('DC68', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC68({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC68) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC70(D28, NamedTuple('DC70', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC70({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC70) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D29_DC71)

class D29_DC71(D29, NamedTuple('DC71', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC71({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC71) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC73(int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D30_DC73)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D30_DC72)

class D30_DC73(D30, NamedTuple('DC73', [('cf123', Any), ('cf124', Any), ('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC73({_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)}, {_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC73) and self.cf123 == __o.cf123 and self.cf124 == __o.cf124 and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC72(D30, NamedTuple('DC72', [('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC72({_dafny.string_of(self.cf122)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC72) and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: D31_DC75(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D31_DC75)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D31_DC74)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D31_DC76)

class D31_DC75(D31, NamedTuple('DC75', [('cf127', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC75({_dafny.string_of(self.cf127)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC75) and self.cf127 == __o.cf127
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC74(D31, NamedTuple('DC74', [('cf126', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC74({_dafny.string_of(self.cf126)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC74) and self.cf126 == __o.cf126
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC76(D31, NamedTuple('DC76', [('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC76({_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC76) and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()


class D32:
    @classmethod
    def default(cls, ):
        return lambda: D32_DC78(False, int(0), False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D32_DC78)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D32_DC79)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D32_DC77)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D32_DC80)

class D32_DC78(D32, NamedTuple('DC78', [('cf130', Any), ('cf131', Any), ('cf132', Any), ('cf133', Any), ('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC78({_dafny.string_of(self.cf130)}, {_dafny.string_of(self.cf131)}, {_dafny.string_of(self.cf132)}, {_dafny.string_of(self.cf133)}, {_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC78) and self.cf130 == __o.cf130 and self.cf131 == __o.cf131 and self.cf132 == __o.cf132 and self.cf133 == __o.cf133 and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC79(D32, NamedTuple('DC79', [('cf135', Any), ('cf136', Any), ('cf137', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC79({_dafny.string_of(self.cf135)}, {_dafny.string_of(self.cf136)}, {_dafny.string_of(self.cf137)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC79) and self.cf135 == __o.cf135 and self.cf136 == __o.cf136 and self.cf137 == __o.cf137
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC77(D32, NamedTuple('DC77', [('cf129', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC77({_dafny.string_of(self.cf129)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC77) and self.cf129 == __o.cf129
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC80(D32, NamedTuple('DC80', [('cf138', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC80({_dafny.string_of(self.cf138)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC80) and self.cf138 == __o.cf138
    def __hash__(self) -> int:
        return super().__hash__()


class D33:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D33_DC81)

class D33_DC81(D33, NamedTuple('DC81', [('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC81({_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC81) and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, globalState):
        pass

    def m2(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: _dafny.Seq = _dafny.Seq({})
        self.f3: _dafny.Array = _dafny.Array(None, 0)
        self._f0: _dafny.Set = _dafny.Set({})
        self._f1: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1

class C0(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, globalState):
        return ((_dafny.MultiSet([True, not(not(False)), True, True, False]) if False else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(True), not(False)])))).isdisjoint((_dafny.MultiSet([not(False), True, False])).intersection(_dafny.MultiSet([True])))

    def fm17(self, p0, p1, p2, globalState):
        return D8_DC18((_dafny.MultiSet([len(_dafny.Map({224: not(True)})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phmadbtnq")))])) - ((D8_DC18(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlf"))), len(_dafny.Map({True: len(_dafny.Map({True: False}))}))]))).cf28))

    def fm18(self, p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaymeth")))) for d_466_i0_ in range(default__.abs(639))])) + (_dafny.SeqWithoutIsStrInference([-14, (0) - (len(_dafny.Map({203: -39})))])))) <= (default__.safeModuloInt(686, 737))


class C1(T1, T0):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f4, f5):
        (self)._f4 = f4
        (self)._f5 = f5

    def fm4(self, p0, globalState):
        def iife27_():
            coll21_ = _dafny.Map()
            compr_21_: _dafny.Seq
            for compr_21_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpqxlhwy")): 356})).keys.Elements:
                d_467_v0_: _dafny.Seq = compr_21_
                if (d_467_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpqxlhwy")): 356})):
                    coll21_[d_467_v0_] = not(False)
            return _dafny.Map(coll21_)
        return default__.safeDivisionInt((0) - ((0) - (default__.safeModuloInt(-466, len(iife27_()
        )))), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwts"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))))

    def fm5(self, p0, p1, p2, p3, globalState):
        return True

    def fm3(self, p0, p1, globalState):
        return False

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        d_468_v0_: _dafny.Array
        nw105_ = _dafny.Array(None, 20)
        d_468_v0_ = nw105_
        d_469_v1_: _dafny.Array
        d_469_v1_ = d_468_v0_
        d_470_v2_: bool
        d_470_v2_ = True
        d_471_v3_: _dafny.Array
        nw106_ = _dafny.Array(None, 25)
        nw106_[int(0)] = d_468_v0_
        nw106_[int(1)] = d_468_v0_
        nw106_[int(2)] = d_468_v0_
        nw106_[int(3)] = d_468_v0_
        nw106_[int(4)] = d_468_v0_
        nw106_[int(5)] = d_468_v0_
        nw106_[int(6)] = d_468_v0_
        nw106_[int(7)] = d_468_v0_
        nw106_[int(8)] = (d_469_v1_)
        nw106_[int(9)] = (d_468_v0_ if d_470_v2_ else d_468_v0_)
        nw106_[int(10)] = d_468_v0_
        nw106_[int(11)] = d_468_v0_
        nw106_[int(12)] = d_468_v0_
        nw106_[int(13)] = d_468_v0_
        nw106_[int(14)] = d_468_v0_
        nw106_[int(15)] = d_468_v0_
        nw106_[int(16)] = d_468_v0_
        nw106_[int(17)] = d_468_v0_
        nw106_[int(18)] = d_468_v0_
        nw106_[int(19)] = d_468_v0_
        nw106_[int(20)] = d_468_v0_
        nw106_[int(21)] = d_468_v0_
        nw106_[int(22)] = d_468_v0_
        nw106_[int(23)] = d_468_v0_
        nw106_[int(24)] = d_468_v0_
        d_471_v3_ = nw106_
        index68_ = default__.safeIndex(396, (d_471_v3_).length(0))
        (d_471_v3_)[index68_] = d_468_v0_
        d_472_v4_: C0
        nw107_ = C0()
        nw107_.ctor__()
        d_472_v4_ = nw107_
        index69_ = default__.safeIndex(396, (d_471_v3_).length(0))
        nw108_ = _dafny.Array(None, 7)
        nw108_[int(0)] = d_472_v4_
        nw108_[int(1)] = d_472_v4_
        nw108_[int(2)] = d_472_v4_
        nw108_[int(3)] = d_472_v4_
        nw108_[int(4)] = d_472_v4_
        nw108_[int(5)] = d_472_v4_
        nw108_[int(6)] = d_472_v4_
        (d_471_v3_)[index69_] = nw108_
        d_473_i0_: int
        d_473_i0_ = 0
        with _dafny.label("1"):
            while d_470_v2_:
                with _dafny.c_label("1"):
                    if (d_473_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_473_i0_ = (d_473_i0_) + (1)
                    d_474_v5_: _dafny.Array
                    nw109_ = _dafny.Array(D11.default()(), 29)
                    d_474_v5_ = nw109_
                    d_475_v6_: _dafny.Array
                    def lambda20_(d_476_i1_):
                        return _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewp")))})

                    init10_ = lambda20_
                    nw110_ = _dafny.Array(None, 22)
                    for i0_10_ in range(nw110_.length(0)):
                        nw110_[i0_10_] = init10_(i0_10_)
                    d_475_v6_ = nw110_
                    index70_ = default__.safeIndex(298, (d_474_v5_).length(0))
                    (d_474_v5_)[index70_] = D11_DC24(d_475_v6_)
                    d_477_v7_: D11
                    d_477_v7_ = D11_DC24(d_475_v6_)
                    index71_ = default__.safeIndex(298, (d_474_v5_).length(0))
                    (d_474_v5_)[index71_] = d_477_v7_
                    d_478_v8_: int
                    d_478_v8_ = 6
                    r1 = (d_478_v8_) - (d_478_v8_)
                    d_479_v9_: _dafny.Array
                    nw111_ = _dafny.Array(_dafny.Seq({}), 13)
                    d_479_v9_ = nw111_
                    d_480_v10_: D4
                    d_480_v10_ = D4_DC9(d_479_v9_)
                    d_481_v11_: _dafny.Seq
                    d_481_v11_ = _dafny.SeqWithoutIsStrInference([d_470_v2_])
                    d_480_v10_ = (d_480_v10_ if (d_478_v8_) < (len(d_481_v11_)) else d_480_v10_)
                    d_470_v2_ = False
                    pass
            pass
        d_482_v12_: _dafny.Array
        def lambda21_(d_483_i2_):
            return (d_483_i2_) - (-463)

        init11_ = lambda21_
        nw112_ = _dafny.Array(None, 13)
        for i0_11_ in range(nw112_.length(0)):
            nw112_[i0_11_] = init11_(i0_11_)
        d_482_v12_ = nw112_
        d_484_v13_: _dafny.Map
        d_484_v13_ = _dafny.Map({d_470_v2_: d_482_v12_})
        d_484_v13_ = (d_484_v13_ if d_470_v2_ else d_484_v13_)
        d_485_v14_: _dafny.Array
        def lambda22_(d_486_v2_):
            def lambda23_(d_487_i3_):
                return (_dafny.MultiSet([d_486_v2_])).ispropersubset(_dafny.MultiSet([d_486_v2_, d_486_v2_, d_486_v2_]))

            return lambda23_

        init12_ = lambda22_(d_470_v2_)
        nw113_ = _dafny.Array(None, 6)
        for i0_12_ in range(nw113_.length(0)):
            nw113_[i0_12_] = init12_(i0_12_)
        d_485_v14_ = nw113_
        (globalState).f3 = d_485_v14_
        d_488_v15_: int
        d_488_v15_ = 282
        r1 = d_488_v15_
        d_489_v16_: _dafny.Map
        d_489_v16_ = _dafny.Map({d_470_v2_: d_470_v2_})
        d_490_v17_: _dafny.Map
        d_490_v17_ = _dafny.Map({len(d_489_v16_): not(d_470_v2_)})
        d_490_v17_ = (d_490_v17_).set((0) - (d_488_v15_), d_470_v2_)
        d_491_v18_: _dafny.Map
        d_491_v18_ = _dafny.Map({d_482_v12_: d_488_v15_})
        d_492_v19_: _dafny.Seq
        d_492_v19_ = _dafny.SeqWithoutIsStrInference([len(d_491_v18_)])
        r0 = (d_492_v19_) + (d_492_v19_)
        r1 = d_488_v15_
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_493_v0_: bool
        d_493_v0_ = True
        d_494_v1_: _dafny.Seq
        d_494_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwx"))
        d_495_v2_: int
        d_495_v2_ = 425
        d_496_v3_: _dafny.Map
        d_496_v3_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_494_v1_])): d_495_v2_})
        d_497_v4_: _dafny.Map
        d_497_v4_ = _dafny.Map({d_493_v0_: d_496_v3_})
        d_498_i0_: int
        d_498_i0_ = 0
        with _dafny.label("2"):
            while (d_493_v0_) not in (d_497_v4_):
                with _dafny.c_label("2"):
                    if (d_498_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_498_i0_ = (d_498_i0_) + (1)
                    d_499_v5_: C0
                    nw114_ = C0()
                    nw114_.ctor__()
                    d_499_v5_ = nw114_
                    d_495_v2_ = (285) * (d_495_v2_)
                    d_493_v0_ = d_493_v0_
                    d_500_v6_: _dafny.Seq
                    d_500_v6_ = _dafny.SeqWithoutIsStrInference([d_495_v2_, d_495_v2_])
                    d_501_v7_: _dafny.Map
                    d_501_v7_ = _dafny.Map({(self).fm4(d_500_v6_, globalState): d_493_v0_})
                    d_502_v8_: _dafny.MultiSet
                    d_502_v8_ = _dafny.MultiSet([d_501_v7_])
                    if (d_502_v8_).ispropersubset(d_502_v8_):
                        d_503_v9_: _dafny.MultiSet
                        d_503_v9_ = _dafny.MultiSet([False])
                        d_504_v10_: _dafny.Set
                        d_504_v10_ = _dafny.Set({(self).fm4(d_500_v6_, globalState), d_495_v2_, d_495_v2_, ((d_503_v9_)[d_493_v0_] if (d_493_v0_) in (d_503_v9_) else (0) - (len(d_494_v1_))), d_495_v2_})
                        d_505_v11_: _dafny.Seq
                        d_505_v11_ = _dafny.SeqWithoutIsStrInference([d_493_v0_])
                        rhs88_ = (default__.fm22((self).f5, d_493_v0_, d_495_v2_, globalState)).intersection(d_504_v10_)
                        rhs89_ = not((d_505_v11_)[default__.safeIndex(d_495_v2_, len(d_505_v11_))])
                        rhs90_ = (d_495_v2_ if d_493_v0_ else d_495_v2_)
                        d_504_v10_ = rhs88_
                        d_493_v0_ = rhs89_
                        d_495_v2_ = rhs90_
                        d_506_v12_: _dafny.Seq
                        d_506_v12_ = _dafny.SeqWithoutIsStrInference([(d_505_v11_).set(default__.safeIndex(len(d_500_v6_), len(d_505_v11_)), d_493_v0_), d_505_v11_, _dafny.SeqWithoutIsStrInference([d_493_v0_, d_493_v0_, d_493_v0_, d_493_v0_, d_493_v0_])])
                        d_507_v13_: _dafny.Map
                        d_507_v13_ = _dafny.Map({d_493_v0_: d_493_v0_})
                        d_508_v14_: _dafny.Map
                        d_508_v14_ = _dafny.Map({not(((self).f5) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcajlchw")))): len(d_507_v13_)})
                        rhs91_ = (d_505_v11_)[default__.safeIndex(d_495_v2_, len(d_505_v11_))]
                        rhs92_ = default__.safeDivisionInt((self).fm4(_dafny.SeqWithoutIsStrInference([d_495_v2_ for d_509_i1_ in range(default__.abs(36))]), globalState), d_495_v2_)
                        rhs93_ = _dafny.SeqWithoutIsStrInference([d_505_v11_, d_505_v11_])
                        rhs94_ = d_508_v14_
                        d_493_v0_ = rhs91_
                        d_495_v2_ = rhs92_
                        d_506_v12_ = rhs93_
                        d_508_v14_ = rhs94_
                        d_510_v15_: _dafny.Array
                        def lambda24_(d_511_i2_):
                            return (self).f5

                        init13_ = lambda24_
                        nw115_ = _dafny.Array(None, 12)
                        for i0_13_ in range(nw115_.length(0)):
                            nw115_[i0_13_] = init13_(i0_13_)
                        d_510_v15_ = nw115_
                        d_512_v16_: D10
                        d_512_v16_ = D10_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fenhui")), d_510_v15_, d_495_v2_)
                        d_513_v17_: _dafny.Array
                        def lambda25_(d_514_v2_):
                            def lambda26_(d_515_i3_):
                                return (d_515_i3_) * (d_514_v2_)

                            return lambda26_

                        init14_ = lambda25_(d_495_v2_)
                        nw116_ = _dafny.Array(None, 6)
                        for i0_14_ in range(nw116_.length(0)):
                            nw116_[i0_14_] = init14_(i0_14_)
                        d_513_v17_ = nw116_
                        d_516_v18_: _dafny.Map
                        d_516_v18_ = _dafny.Map({d_513_v17_: d_495_v2_})
                        d_517_v19_: _dafny.Map
                        d_517_v19_ = _dafny.Map({d_495_v2_: d_516_v18_})
                        d_493_v0_ = (default__.safeDivisionInt(d_495_v2_, (d_512_v16_).cf35)) < ((d_495_v2_) - (len(((d_517_v19_)[d_495_v2_] if (d_495_v2_) in (d_517_v19_) else (d_516_v18_).set(d_513_v17_, (0) - (len(d_505_v11_)))))))
                        d_518_v20_: _dafny.Map
                        d_518_v20_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(0) - (len(d_494_v1_))]) if d_493_v0_ else d_500_v6_): d_493_v0_})
                        d_518_v20_ = d_518_v20_
                        d_519_v21_: _dafny.Array
                        nw117_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_519_v21_ = nw117_
                        d_520_v22_: _dafny.Array
                        def lambda27_(d_521_v0_):
                            def lambda28_(d_522_i4_):
                                return d_521_v0_

                            return lambda28_

                        init15_ = lambda27_(d_493_v0_)
                        nw118_ = _dafny.Array(None, 25)
                        for i0_15_ in range(nw118_.length(0)):
                            nw118_[i0_15_] = init15_(i0_15_)
                        d_520_v22_ = nw118_
                        index72_ = default__.safeIndex(796, (d_519_v21_).length(0))
                        (d_519_v21_)[index72_] = d_520_v22_
                        index73_ = default__.safeIndex(796, (d_519_v21_).length(0))
                        (d_519_v21_)[index73_] = d_520_v22_
                    elif True:
                        d_523_v23_: C0
                        nw119_ = C0()
                        nw119_.ctor__()
                        d_523_v23_ = nw119_
                        d_494_v1_ = d_494_v1_
                        d_524_v24_: _dafny.Array
                        nw120_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                        d_524_v24_ = nw120_
                        d_524_v24_ = d_524_v24_
                        d_525_v25_: _dafny.MultiSet
                        d_525_v25_ = _dafny.MultiSet([(self).f5, (self).f5])
                        d_493_v0_ = (d_525_v25_) == (default__.fm23(d_493_v0_, globalState))
                        d_526_v26_: _dafny.Map
                        d_526_v26_ = _dafny.Map({d_493_v0_: d_495_v2_})
                        d_526_v26_ = (d_526_v26_).set(True, d_495_v2_)
                    pass
            pass
        d_527_v27_: _dafny.Array
        def lambda29_(d_528_i5_):
            return (self).f5

        init16_ = lambda29_
        nw121_ = _dafny.Array(None, 25)
        for i0_16_ in range(nw121_.length(0)):
            nw121_[i0_16_] = init16_(i0_16_)
        d_527_v27_ = nw121_
        index74_ = default__.safeIndex(494, (d_527_v27_).length(0))
        (d_527_v27_)[index74_] = (self).f5
        d_529_v28_: _dafny.Seq
        d_529_v28_ = _dafny.SeqWithoutIsStrInference([not(d_493_v0_), True])
        index75_ = default__.safeIndex(494, (d_527_v27_).length(0))
        (d_527_v27_)[index75_] = ((self).f5 if (d_493_v0_) in (d_529_v28_) else (self).f5)
        d_493_v0_ = d_493_v0_
        d_495_v2_ = default__.safeModuloInt(len(_dafny.Set({d_495_v2_})), d_495_v2_)
        hi5_ = d_495_v2_
        for d_530_i6_ in range(len(d_494_v1_), hi5_):
            d_531_v29_: C0
            nw122_ = C0()
            nw122_.ctor__()
            d_531_v29_ = nw122_
            d_495_v2_ = d_495_v2_
            d_532_v30_: _dafny.Array
            nw123_ = _dafny.Array(False, 23)
            d_532_v30_ = nw123_
            d_533_v31_: _dafny.Map
            d_533_v31_ = _dafny.Map({d_493_v0_: d_532_v30_})
            d_534_v32_: _dafny.Array
            nw124_ = _dafny.Array(None, 14)
            nw124_[int(0)] = d_532_v30_
            nw124_[int(1)] = d_532_v30_
            nw124_[int(2)] = d_532_v30_
            nw124_[int(3)] = (((d_533_v31_)[d_493_v0_] if (d_493_v0_) in (d_533_v31_) else d_532_v30_) if d_493_v0_ else d_532_v30_)
            nw124_[int(4)] = d_532_v30_
            nw124_[int(5)] = d_532_v30_
            nw124_[int(6)] = d_532_v30_
            nw124_[int(7)] = d_532_v30_
            nw124_[int(8)] = d_532_v30_
            nw124_[int(9)] = d_532_v30_
            nw124_[int(10)] = d_532_v30_
            nw124_[int(11)] = d_532_v30_
            nw124_[int(12)] = d_532_v30_
            nw124_[int(13)] = d_532_v30_
            d_534_v32_ = nw124_
            rhs95_ = d_534_v32_
            rhs96_ = d_532_v30_
            d_534_v32_ = rhs95_
            d_532_v30_ = rhs96_
            index76_ = default__.safeIndex(950, (d_532_v30_).length(0))
            (d_532_v30_)[index76_] = d_493_v0_
            d_535_v33_: _dafny.Array
            nw125_ = _dafny.Array(_dafny.Seq({}), 11)
            d_535_v33_ = nw125_
            d_536_v34_: D13
            d_536_v34_ = D13_DC28(d_535_v33_)
            d_537_v35_: _dafny.Array
            nw126_ = _dafny.Array(None, 6)
            nw126_[int(0)] = (d_536_v34_).cf44
            nw126_[int(1)] = d_535_v33_
            nw126_[int(2)] = d_535_v33_
            nw126_[int(3)] = d_535_v33_
            nw126_[int(4)] = (d_535_v33_ if d_493_v0_ else d_535_v33_)
            nw126_[int(5)] = d_535_v33_
            d_537_v35_ = nw126_
            index77_ = default__.safeIndex(956, (d_537_v35_).length(0))
            (d_537_v35_)[index77_] = d_535_v33_
            d_538_v36_: _dafny.Seq
            d_538_v36_ = _dafny.SeqWithoutIsStrInference([d_535_v33_, d_535_v33_])
            index78_ = default__.safeIndex(950, (d_532_v30_).length(0))
            index79_ = default__.safeIndex(956, (d_537_v35_).length(0))
            rhs97_ = d_495_v2_
            rhs98_ = d_493_v0_
            rhs99_ = not(not(not(d_493_v0_)))
            rhs100_ = (d_535_v33_ if d_493_v0_ else (d_538_v36_)[default__.safeIndex(994, len(d_538_v36_))])
            lhs46_ = d_532_v30_
            lhs47_ = default__.safeIndex(950, (d_532_v30_).length(0))
            lhs48_ = d_537_v35_
            lhs49_ = default__.safeIndex(956, (d_537_v35_).length(0))
            d_495_v2_ = rhs97_
            lhs46_[lhs47_] = rhs98_
            d_493_v0_ = rhs99_
            lhs48_[lhs49_] = rhs100_
        d_539_v37_: _dafny.Map
        d_539_v37_ = _dafny.Map({d_493_v0_: (d_527_v27_)[default__.safeIndex(494, (d_527_v27_).length(0))]})
        d_540_v38_: D5
        d_540_v38_ = D5_DC11(d_539_v37_)
        d_541_v39_: _dafny.Map
        d_541_v39_ = _dafny.Map({d_495_v2_: d_540_v38_})
        d_542_v40_: _dafny.Seq
        d_542_v40_ = _dafny.SeqWithoutIsStrInference([d_541_v39_])
        d_542_v40_ = (d_542_v40_) + (d_542_v40_)
        d_543_v41_: _dafny.Array
        def lambda30_(d_544_v0_):
            def lambda31_(d_545_i7_):
                return (d_545_i7_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_544_v0_, d_544_v0_]))).cardinality)

            return lambda31_

        init17_ = lambda30_(d_493_v0_)
        nw127_ = _dafny.Array(None, 3)
        for i0_17_ in range(nw127_.length(0)):
            nw127_[i0_17_] = init17_(i0_17_)
        d_543_v41_ = nw127_
        r0 = (d_543_v41_ if False else d_543_v41_)
        return r0

    def m14(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_546_v0_: bool
        d_546_v0_ = False
        d_547_v1_: _dafny.Seq
        d_547_v1_ = _dafny.SeqWithoutIsStrInference([p1])
        d_548_v2_: _dafny.MultiSet
        d_548_v2_ = _dafny.MultiSet([((_dafny.MultiSet([d_546_v0_, d_546_v0_])).cardinality) not in (d_547_v1_), d_546_v0_])
        d_549_v3_: _dafny.Array
        def lambda32_(d_550_v0_):
            def lambda33_(d_551_i0_):
                return (d_550_v0_) == (d_550_v0_)

            return lambda33_

        init18_ = lambda32_(d_546_v0_)
        nw128_ = _dafny.Array(None, 12)
        for i0_18_ in range(nw128_.length(0)):
            nw128_[i0_18_] = init18_(i0_18_)
        d_549_v3_ = nw128_
        index80_ = default__.safeIndex(910, (d_549_v3_).length(0))
        (d_549_v3_)[index80_] = d_546_v0_
        d_552_v4_: _dafny.Seq
        d_552_v4_ = _dafny.SeqWithoutIsStrInference([(d_548_v2_ if d_546_v0_ else d_548_v2_)])
        d_553_v5_: _dafny.Set
        d_553_v5_ = _dafny.Set({d_546_v0_, not(d_546_v0_), d_546_v0_})
        d_554_v6_: _dafny.Map
        d_554_v6_ = _dafny.Map({True: (0) - ((0) - (len(d_553_v5_)))})
        index81_ = default__.safeIndex(910, (d_549_v3_).length(0))
        rhs101_ = (d_552_v4_)[default__.safeIndex(((d_554_v6_)[not(d_546_v0_)] if (not(d_546_v0_)) in (d_554_v6_) else (self).fm4(_dafny.SeqWithoutIsStrInference([p1 for d_555_i1_ in range(default__.abs(154))]), globalState)), len(d_552_v4_))]
        rhs102_ = d_546_v0_
        lhs50_ = d_549_v3_
        lhs51_ = default__.safeIndex(910, (d_549_v3_).length(0))
        d_548_v2_ = rhs101_
        lhs50_[lhs51_] = rhs102_
        if (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]:
            d_556_v7_: _dafny.Set
            d_556_v7_ = _dafny.Set({(self).f5})
            d_557_v8_: _dafny.Seq
            d_557_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyekk"))
            d_558_v9_: _dafny.Array
            nw129_ = _dafny.Array(_dafny.CodePoint('D'), 7)
            d_558_v9_ = nw129_
            d_559_v10_: _dafny.Map
            d_559_v10_ = _dafny.Map({d_557_v8_: d_558_v9_})
            d_560_v11_: _dafny.Array
            nw130_ = _dafny.Array(None, 11)
            nw130_[int(0)] = p0
            nw130_[int(1)] = (self).fm4(d_547_v1_, globalState)
            nw130_[int(2)] = p0
            nw130_[int(3)] = (p0) - (p0)
            nw130_[int(4)] = (p0) - (p0)
            nw130_[int(5)] = p0
            nw130_[int(6)] = len((d_554_v6_) | (d_554_v6_))
            nw130_[int(7)] = (0) - (p0)
            nw130_[int(8)] = len(d_556_v7_)
            nw130_[int(9)] = len((d_559_v10_ if (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))] else (d_559_v10_).set(d_557_v8_, d_558_v9_)))
            nw130_[int(10)] = (0) - (p1)
            d_560_v11_ = nw130_
            d_560_v11_ = (d_560_v11_ if False else d_560_v11_)
            if True:
                index82_ = default__.safeIndex(910, (d_549_v3_).length(0))
                (d_549_v3_)[index82_] = (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]
                r2 = p0
                d_546_v0_ = (not((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]) if ((self).f5) not in (d_557_v8_) else True)
                index83_ = default__.safeIndex(910, (d_549_v3_).length(0))
                (d_549_v3_)[index83_] = not (d_546_v0_) or (not((self).fm5(p1, (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], p1, p1, globalState)))
                d_561_v12_: _dafny.Seq
                d_561_v12_ = _dafny.SeqWithoutIsStrInference([d_547_v1_, d_547_v1_])
                d_562_v13_: _dafny.Set
                d_562_v13_ = _dafny.Set({d_547_v1_, d_547_v1_, (d_547_v1_) + (d_547_v1_), (d_547_v1_) + ((d_561_v12_)[default__.safeIndex(p0, len(d_561_v12_))]), d_547_v1_})
                d_562_v13_ = d_562_v13_
            elif True:
                d_563_v14_: _dafny.Set
                d_563_v14_ = _dafny.Set({d_557_v8_, d_557_v8_, d_557_v8_})
                def iife28_():
                    coll22_ = _dafny.Set()
                    compr_22_: _dafny.Seq
                    for compr_22_ in (d_563_v14_).Elements:
                        d_564_v15_: _dafny.Seq = compr_22_
                        if (d_564_v15_) in (d_563_v14_):
                            coll22_ = coll22_.union(_dafny.Set([d_564_v15_]))
                    return _dafny.Set(coll22_)
                d_563_v14_ = ((d_563_v14_ if (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))] else iife28_()
                )) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kw"))}))
                d_565_v16_: C0
                nw131_ = C0()
                nw131_.ctor__()
                d_565_v16_ = nw131_
                r1 = p0
                index84_ = default__.safeIndex(856, (d_560_v11_).length(0))
                (d_560_v11_)[index84_] = p0
                index85_ = default__.safeIndex(856, (d_560_v11_).length(0))
                (d_560_v11_)[index85_] = (self).fm4((d_547_v1_) + (d_547_v1_), globalState)
                d_566_v17_: _dafny.Map
                d_566_v17_ = _dafny.Map({(d_560_v11_)[default__.safeIndex(856, (d_560_v11_).length(0))]: ((d_565_v16_).fm17(_dafny.SeqWithoutIsStrInference([608 for d_567_i2_ in range(default__.abs(-198))]), d_546_v0_, d_557_v8_, globalState)).cf28})
                d_554_v6_ = (d_554_v6_).set(d_546_v0_, (0) - (len(d_566_v17_)))
            d_568_v19_: _dafny.Map
            def iife29_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in (_dafny.SeqWithoutIsStrInference([p1 for d_569_i3_ in range(default__.abs(709))])).Elements:
                    d_570_v18_: int = compr_23_
                    if (d_570_v18_) in (_dafny.SeqWithoutIsStrInference([p1 for d_569_i3_ in range(default__.abs(709))])):
                        coll23_ = coll23_.union(_dafny.Set([(d_570_v18_) - (492)]))
                return _dafny.Set(coll23_)
            d_568_v19_ = _dafny.Map({iife29_()
            : (0) - (len(d_553_v5_))})
            d_546_v0_ = not(default__.fm1(91, p1, (len(d_568_v19_)) != ((0) - (p1)), d_546_v0_, globalState))
            d_571_v20_: _dafny.Array
            def lambda34_(d_572_v1_):
                def lambda35_(d_573_i4_):
                    return d_572_v1_

                return lambda35_

            init19_ = lambda34_(d_547_v1_)
            nw132_ = _dafny.Array(None, 17)
            for i0_19_ in range(nw132_.length(0)):
                nw132_[i0_19_] = init19_(i0_19_)
            d_571_v20_ = nw132_
            source9_ = D13_DC28(d_571_v20_)
            if source9_.is_DC29:
                d_574___mcc_h0_ = source9_.cf45
                d_575___mcc_h1_ = source9_.cf46
                d_576_cf46_ = d_575___mcc_h1_
                d_577_cf45_ = d_574___mcc_h0_
                d_578_v21_: C0
                nw133_ = C0()
                nw133_.ctor__()
                d_578_v21_ = nw133_
                d_579_v22_: _dafny.Array
                nw134_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_579_v22_ = nw134_
                d_579_v22_ = d_579_v22_
                r0 = (0) - (p0)
                d_580_v23_: D13
                d_580_v23_ = D13_DC28(d_571_v20_)
                d_581_v24_: _dafny.Set
                d_581_v24_ = _dafny.Set({(0) - (p0)})
                d_582_v26_: _dafny.Seq
                def iife30_():
                    coll24_ = _dafny.Set()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(945, 999):
                        d_583_v25_: int = compr_24_
                        if ((945) <= (d_583_v25_)) and ((d_583_v25_) < (999)):
                            coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_583_v25_, p1)]))
                    return _dafny.Set(coll24_)
                d_582_v26_ = _dafny.SeqWithoutIsStrInference([(iife30_()
                ).issubset(d_581_v24_), d_577_cf45_, (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]])
                rhs103_ = d_580_v23_
                rhs104_ = (d_582_v26_)[default__.safeIndex(p1, len(d_582_v26_))]
                d_580_v23_ = rhs103_
                d_577_cf45_ = rhs104_
            elif True:
                d_584___mcc_h2_ = source9_.cf44
                d_585_cf44_ = d_584___mcc_h2_
                rhs105_ = not((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))])
                rhs106_ = (self).fm3(((d_554_v6_)[d_546_v0_] if (d_546_v0_) in (d_554_v6_) else (0) - (p1)), (0) - (p0), globalState)
                rhs107_ = (((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))] if d_546_v0_ else d_546_v0_)) and ((self).fm3(p1, p1, globalState))
                d_546_v0_ = rhs105_
                r3 = rhs106_
                r3 = rhs107_
                d_586_v27_: _dafny.Seq
                d_586_v27_ = _dafny.SeqWithoutIsStrInference([d_560_v11_, d_560_v11_, d_560_v11_])
                d_560_v11_ = (d_586_v27_)[default__.safeIndex(p0, len(d_586_v27_))]
                d_587_v28_: _dafny.Map
                d_587_v28_ = _dafny.Map({(self).f5: (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]})
                d_587_v28_ = (d_587_v28_).set((self).f5, (p1) == (p0))
                r0 = len(((d_557_v8_).set(default__.safeIndex(p0, len(d_557_v8_)), _dafny.CodePoint('d'))) + (d_557_v8_))
            rhs108_ = default__.fm24((self).fm4(d_547_v1_, globalState), globalState)
            rhs109_ = d_546_v0_
            d_557_v8_ = rhs108_
            d_546_v0_ = rhs109_
        elif True:
            if d_546_v0_:
                index86_ = default__.safeIndex(910, (d_549_v3_).length(0))
                (d_549_v3_)[index86_] = d_546_v0_
                d_588_v29_: _dafny.Map
                d_588_v29_ = _dafny.Map({d_546_v0_: False})
                d_589_v30_: _dafny.Map
                d_589_v30_ = _dafny.Map({(d_588_v29_) | (d_588_v29_): d_548_v2_})
                d_589_v30_ = d_589_v30_
                d_590_v31_: _dafny.Seq
                d_590_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laj"))
                d_590_v31_ = d_590_v31_
                r2 = default__.safeDivisionInt(p0, p0)
                d_591_v32_: D1
                d_591_v32_ = D1_DC1(_dafny.MultiSet([d_546_v0_]))
                d_592_v33_: _dafny.MultiSet
                d_592_v33_ = _dafny.MultiSet([d_591_v32_])
                d_593_v34_: _dafny.Seq
                d_593_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm1(p1, p1, (self).fm5(p1, (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], 787, ((_dafny.MultiSet([(d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]])).set((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], default__.abs(p1))).cardinality, globalState), d_546_v0_, globalState)])
                d_594_v35_: _dafny.Set
                d_594_v35_ = _dafny.Set({(self).f5})
                d_595_v36_: _dafny.Array
                nw135_ = _dafny.Array(None, 20)
                nw135_[int(0)] = 65
                nw135_[int(1)] = p1
                nw135_[int(2)] = (self).fm4(d_547_v1_, globalState)
                nw135_[int(3)] = p0
                nw135_[int(4)] = p1
                nw135_[int(5)] = p1
                nw135_[int(6)] = p0
                nw135_[int(7)] = (d_592_v33_).cardinality
                nw135_[int(8)] = p0
                nw135_[int(9)] = (d_548_v2_).cardinality
                nw135_[int(10)] = p0
                nw135_[int(11)] = p0
                nw135_[int(12)] = p1
                nw135_[int(13)] = len(d_590_v31_)
                nw135_[int(14)] = (0) - (p1)
                nw135_[int(15)] = len(d_593_v34_)
                nw135_[int(16)] = p1
                nw135_[int(17)] = 72
                nw135_[int(18)] = len(d_594_v35_)
                nw135_[int(19)] = p1
                d_595_v36_ = nw135_
                d_596_v37_: _dafny.Map
                d_596_v37_ = _dafny.Map({d_595_v36_: (d_593_v34_) < (d_593_v34_)})
                d_596_v37_ = d_596_v37_
            elif True:
                r1 = (0) - ((d_548_v2_).cardinality)
                d_597_v38_: _dafny.Array
                nw136_ = _dafny.Array(int(0), 12)
                d_597_v38_ = nw136_
                d_598_v39_: _dafny.Map
                d_598_v39_ = _dafny.Map({d_597_v38_: (d_553_v5_).ispropersubset(d_553_v5_)})
                d_598_v39_ = (d_598_v39_).set(d_597_v38_, d_546_v0_)
                d_599_v40_: _dafny.Array
                nw137_ = _dafny.Array(_dafny.Set({}), 14)
                d_599_v40_ = nw137_
                d_600_v41_: D13
                d_600_v41_ = D13_DC29(False, (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))])
                d_601_v42_: _dafny.Seq
                d_601_v42_ = _dafny.SeqWithoutIsStrInference([(d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]])
                pat_let_tv19_ = d_546_v0_
                d_602_v43_: _dafny.Array
                nw138_ = _dafny.Array(None, 18)
                nw138_[int(0)] = d_600_v41_
                nw138_[int(1)] = d_600_v41_
                nw138_[int(2)] = d_600_v41_
                nw138_[int(3)] = d_600_v41_
                nw138_[int(4)] = d_600_v41_
                nw138_[int(5)] = d_600_v41_
                nw138_[int(6)] = d_600_v41_
                nw138_[int(7)] = d_600_v41_
                nw138_[int(8)] = d_600_v41_
                nw138_[int(9)] = D13_DC29(d_546_v0_, (d_601_v42_)[default__.safeIndex(p1, len(d_601_v42_))])
                nw138_[int(10)] = D13_DC29(not(d_546_v0_), True)
                nw138_[int(11)] = D13_DC29((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], d_546_v0_)
                nw138_[int(12)] = d_600_v41_
                def iife31_(_pat_let3_0):
                    def iife32_(d_603_dt__update__tmp_h0_):
                        def iife33_(_pat_let4_0):
                            def iife34_(d_604_dt__update_hcf46_h0_):
                                return D13_DC29((d_603_dt__update__tmp_h0_).cf45, d_604_dt__update_hcf46_h0_)
                            return iife34_(_pat_let4_0)
                        return iife33_(pat_let_tv19_)
                    return iife32_(_pat_let3_0)
                nw138_[int(13)] = iife31_(d_600_v41_)
                nw138_[int(14)] = d_600_v41_
                nw138_[int(15)] = d_600_v41_
                nw138_[int(16)] = d_600_v41_
                nw138_[int(17)] = d_600_v41_
                d_602_v43_ = nw138_
                d_605_v44_: _dafny.Set
                d_605_v44_ = _dafny.Set({d_602_v43_})
                index87_ = default__.safeIndex(112, (d_599_v40_).length(0))
                (d_599_v40_)[index87_] = (d_605_v44_) | (d_605_v44_)
                d_606_v45_: _dafny.Seq
                d_606_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xo"))
                index88_ = default__.safeIndex(910, (d_549_v3_).length(0))
                index89_ = default__.safeIndex(112, (d_599_v40_).length(0))
                rhs110_ = (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]
                rhs111_ = ((d_605_v44_ if d_546_v0_ else _dafny.Set({d_602_v43_, d_602_v43_}))) | (d_605_v44_)
                rhs112_ = d_606_v45_
                rhs113_ = (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]
                rhs114_ = (d_598_v39_) | (d_598_v39_)
                lhs52_ = d_549_v3_
                lhs53_ = default__.safeIndex(910, (d_549_v3_).length(0))
                lhs54_ = d_599_v40_
                lhs55_ = default__.safeIndex(112, (d_599_v40_).length(0))
                lhs52_[lhs53_] = rhs110_
                lhs54_[lhs55_] = rhs111_
                d_606_v45_ = rhs112_
                d_546_v0_ = rhs113_
                d_598_v39_ = rhs114_
                d_607_v46_: bool
                d_607_v46_ = not(default__.fm1(p0, (0) - (p1), False, d_546_v0_, globalState))
                r3 = (d_607_v46_)
                r0 = (0) - (p0)
            d_546_v0_ = not (False) or (d_546_v0_)
            d_608_v47_: _dafny.Array
            nw139_ = _dafny.Array(int(0), 16)
            d_608_v47_ = nw139_
            d_609_v48_: D11
            d_609_v48_ = D11_DC25(False, p0, d_608_v47_, p1)
            if (d_609_v48_).cf38:
                d_610_v49_: str
                d_610_v49_ = _dafny.CodePoint('y')
                d_610_v49_ = d_610_v49_
                d_611_v50_: _dafny.Seq
                d_611_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyoke"))
                d_546_v0_ = (self).fm3(len(d_611_v50_), (self).fm4(d_547_v1_, globalState), globalState)
                d_612_v51_: _dafny.Map
                d_612_v51_ = _dafny.Map({p1: d_549_v3_})
                d_613_v52_: _dafny.Set
                d_613_v52_ = _dafny.Set({d_549_v3_, ((d_612_v51_)[710] if (710) in (d_612_v51_) else d_549_v3_), d_549_v3_})
                index90_ = default__.safeIndex(910, (d_549_v3_).length(0))
                (d_549_v3_)[index90_] = (d_613_v52_).ispropersubset(d_613_v52_)
                d_614_v53_: D2
                d_614_v53_ = D2_DC3(d_549_v3_)
                d_614_v53_ = d_614_v53_
                d_611_v50_ = (d_611_v50_).set(default__.safeIndex(p1, len(d_611_v50_)), d_610_v49_)
            elif True:
                d_615_v54_: D8
                d_615_v54_ = D8_DC19(p0, (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))])
                d_616_v55_: _dafny.Map
                d_616_v55_ = _dafny.Map({605: d_546_v0_})
                pat_let_tv20_ = d_608_v47_
                d_617_v56_: _dafny.Array
                nw140_ = _dafny.Array(None, 24)
                nw140_[int(0)] = d_609_v48_
                nw140_[int(1)] = D11_DC25(d_546_v0_, p1, d_608_v47_, p0)
                nw140_[int(2)] = d_609_v48_
                nw140_[int(3)] = D11_DC25(d_546_v0_, p1, d_608_v47_, 826)
                nw140_[int(4)] = D11_DC25(False, p0, d_608_v47_, p1)
                nw140_[int(5)] = d_609_v48_
                nw140_[int(6)] = d_609_v48_
                def iife35_(_pat_let5_0):
                    def iife36_(d_618_dt__update__tmp_h1_):
                        def iife37_(_pat_let6_0):
                            def iife38_(d_619_dt__update_hcf40_h0_):
                                return D11_DC25((d_618_dt__update__tmp_h1_).cf38, (d_618_dt__update__tmp_h1_).cf39, d_619_dt__update_hcf40_h0_, (d_618_dt__update__tmp_h1_).cf41)
                            return iife38_(_pat_let6_0)
                        return iife37_(pat_let_tv20_)
                    return iife36_(_pat_let5_0)
                nw140_[int(7)] = iife35_(D11_DC25((D8_DC19(len(d_547_v1_), (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))])).cf30, p1, d_608_v47_, p0))
                nw140_[int(8)] = (d_609_v48_ if (self).fm3(p0, len(_dafny.Map({(d_615_v54_).cf29: d_546_v0_})), globalState) else d_609_v48_)
                nw140_[int(9)] = D11_DC25(not(((d_616_v55_)[p1] if (p1) in (d_616_v55_) else d_546_v0_)), 181, d_608_v47_, p0)
                nw140_[int(10)] = d_609_v48_
                nw140_[int(11)] = d_609_v48_
                nw140_[int(12)] = d_609_v48_
                nw140_[int(13)] = d_609_v48_
                nw140_[int(14)] = D11_DC25((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))], p1, d_608_v47_, len(_dafny.SeqWithoutIsStrInference([p0])))
                nw140_[int(15)] = d_609_v48_
                nw140_[int(16)] = D11_DC25(d_546_v0_, p0, d_608_v47_, p0)
                nw140_[int(17)] = D11_DC25((self).fm5(len(_dafny.Set({(self).f5})), not(True), p0, p0, globalState), p1, d_608_v47_, (self).fm4(d_547_v1_, globalState))
                nw140_[int(18)] = d_609_v48_
                nw140_[int(19)] = d_609_v48_
                nw140_[int(20)] = d_609_v48_
                nw140_[int(21)] = d_609_v48_
                nw140_[int(22)] = d_609_v48_
                nw140_[int(23)] = d_609_v48_
                d_617_v56_ = nw140_
                index91_ = default__.safeIndex(596, (d_617_v56_).length(0))
                (d_617_v56_)[index91_] = d_609_v48_
                d_620_v57_: str
                d_620_v57_ = _dafny.CodePoint('x')
                d_621_v58_: _dafny.Set
                d_621_v58_ = _dafny.Set({p0, p1, (0) - (p0)})
                d_622_v59_: D2
                d_622_v59_ = D2_DC6(p0)
                d_623_v60_: _dafny.Set
                d_623_v60_ = _dafny.Set({d_622_v59_, d_622_v59_})
                index92_ = default__.safeIndex(596, (d_617_v56_).length(0))
                def iife39_():
                    coll25_ = _dafny.Set()
                    compr_25_: D2
                    for compr_25_ in (d_623_v60_).Elements:
                        d_624_v61_: D2 = compr_25_
                        if (d_624_v61_) in (d_623_v60_):
                            coll25_ = coll25_.union(_dafny.Set([d_624_v61_]))
                    return _dafny.Set(coll25_)
                rhs115_ = True
                rhs116_ = d_609_v48_
                rhs117_ = (d_621_v58_).isdisjoint((d_621_v58_) - (d_621_v58_))
                rhs118_ = len(iife39_()
                )
                rhs119_ = ((self).f5 if d_546_v0_ else _dafny.CodePoint('p'))
                lhs56_ = d_617_v56_
                lhs57_ = default__.safeIndex(596, (d_617_v56_).length(0))
                r3 = rhs115_
                lhs56_[lhs57_] = rhs116_
                r3 = rhs117_
                r1 = rhs118_
                d_620_v57_ = rhs119_
                d_620_v57_ = d_620_v57_
                d_625_v62_: _dafny.Set
                d_625_v62_ = _dafny.Set({d_608_v47_, d_608_v47_, d_608_v47_})
                d_626_v63_: C0
                nw141_ = C0()
                nw141_.ctor__()
                d_626_v63_ = nw141_
                d_627_v64_: _dafny.Array
                nw142_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                d_627_v64_ = nw142_
                index93_ = default__.safeIndex(27, (d_627_v64_).length(0))
                (d_627_v64_)[index93_] = (self).f5
                pat_let_tv21_ = d_546_v0_
                index94_ = default__.safeIndex(27, (d_627_v64_).length(0))
                def iife40_(_pat_let7_0):
                    def iife41_(d_628_dt__update__tmp_h2_):
                        def iife42_(_pat_let8_0):
                            def iife43_(d_629_dt__update_hcf30_h0_):
                                return D8_DC19((d_628_dt__update__tmp_h2_).cf29, d_629_dt__update_hcf30_h0_)
                            return iife43_(_pat_let8_0)
                        return iife42_(pat_let_tv21_)
                    return iife41_(_pat_let7_0)
                rhs120_ = d_625_v62_
                rhs121_ = (d_626_v63_ if (iife40_(d_615_v54_)).cf30 else d_626_v63_)
                rhs122_ = (self).f5
                rhs123_ = (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]
                lhs58_ = d_627_v64_
                lhs59_ = default__.safeIndex(27, (d_627_v64_).length(0))
                d_625_v62_ = rhs120_
                d_626_v63_ = rhs121_
                lhs58_[lhs59_] = rhs122_
                d_546_v0_ = rhs123_
                d_620_v57_ = d_620_v57_
                d_630_v65_: _dafny.Array
                def lambda36_(d_631_i5_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apibiv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkanpg")))

                init20_ = lambda36_
                nw143_ = _dafny.Array(None, 26)
                for i0_20_ in range(nw143_.length(0)):
                    nw143_[i0_20_] = init20_(i0_20_)
                d_630_v65_ = nw143_
                d_632_v66_: _dafny.Seq
                d_632_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnotyjd"))
                index95_ = default__.safeIndex(680, (d_630_v65_).length(0))
                (d_630_v65_)[index95_] = d_632_v66_
                index96_ = default__.safeIndex(910, (d_549_v3_).length(0))
                index97_ = default__.safeIndex(680, (d_630_v65_).length(0))
                rhs124_ = ((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]) == (d_546_v0_)
                rhs125_ = d_632_v66_
                rhs126_ = (p0) + (p0)
                lhs60_ = d_549_v3_
                lhs61_ = default__.safeIndex(910, (d_549_v3_).length(0))
                lhs62_ = d_630_v65_
                lhs63_ = default__.safeIndex(680, (d_630_v65_).length(0))
                lhs60_[lhs61_] = rhs124_
                lhs62_[lhs63_] = rhs125_
                r1 = rhs126_
            index98_ = default__.safeIndex(910, (d_549_v3_).length(0))
            (d_549_v3_)[index98_] = d_546_v0_
            if d_546_v0_:
                d_633_v67_: _dafny.Array
                nw144_ = _dafny.Array(_dafny.Set({}), 5)
                d_633_v67_ = nw144_
                d_634_v68_: _dafny.Map
                d_634_v68_ = _dafny.Map({p1: p0})
                d_635_v69_: D8
                d_635_v69_ = D8_DC19(p0, d_546_v0_)
                index99_ = default__.safeIndex(567, (d_633_v67_).length(0))
                (d_633_v67_)[index99_] = _dafny.Set({((d_634_v68_)[p1] if (p1) in (d_634_v68_) else (d_635_v69_).cf29)})
                d_636_v70_: _dafny.Set
                d_636_v70_ = _dafny.Set({p1, 669, (0) - (p0)})
                index100_ = default__.safeIndex(567, (d_633_v67_).length(0))
                (d_633_v67_)[index100_] = (d_636_v70_) | (d_636_v70_)
                d_637_v71_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_637_v71_ = nw145_
                d_637_v71_ = d_637_v71_
                index101_ = default__.safeIndex(487, (d_608_v47_).length(0))
                (d_608_v47_)[index101_] = p1
                index102_ = default__.safeIndex(487, (d_608_v47_).length(0))
                (d_608_v47_)[index102_] = p1
                d_638_v72_: _dafny.Seq
                d_638_v72_ = _dafny.SeqWithoutIsStrInference([True, (self).fm3(p0, (d_608_v47_)[default__.safeIndex(487, (d_608_v47_).length(0))], globalState)])
                index103_ = default__.safeIndex(910, (d_549_v3_).length(0))
                (d_549_v3_)[index103_] = (default__.fm24(len(d_638_v72_), globalState)) <= (default__.fm24(857, globalState))
                default__.m0(p1, ((self).f5) in (default__.fm24(p0, globalState)), globalState)
            elif True:
                d_639_v73_: _dafny.Array
                nw146_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_639_v73_ = nw146_
                d_640_v74_: _dafny.Seq
                d_640_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpwliykl"))
                index104_ = default__.safeIndex(693, (d_639_v73_).length(0))
                (d_639_v73_)[index104_] = (d_640_v74_).set(default__.safeIndex(p1, len(d_640_v74_)), (self).f5)
                d_641_v75_: D2
                d_641_v75_ = D2_DC4((d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))])
                d_642_v76_: _dafny.Map
                d_642_v76_ = _dafny.Map({d_641_v75_: _dafny.SeqWithoutIsStrInference([(self).f5 for d_643_i6_ in range(default__.abs(410))])})
                d_644_v77_: _dafny.Map
                d_644_v77_ = _dafny.Map({((d_642_v76_)[D2_DC4(d_546_v0_)] if (D2_DC4(d_546_v0_)) in (d_642_v76_) else _dafny.SeqWithoutIsStrInference([(self).f5 for d_645_i7_ in range(default__.abs(973))])): d_640_v74_})
                index105_ = default__.safeIndex(693, (d_639_v73_).length(0))
                (d_639_v73_)[index105_] = (((d_644_v77_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr"))) in (d_644_v77_) else d_640_v74_)) + ((d_640_v74_) + (d_640_v74_))
                d_646_v78_: _dafny.Array
                nw147_ = _dafny.Array(None, 19)
                nw147_[int(0)] = (self).f5
                nw147_[int(1)] = (self).f5
                nw147_[int(2)] = (self).f5
                nw147_[int(3)] = (self).f5
                nw147_[int(4)] = (self).f5
                nw147_[int(5)] = (self).f5
                nw147_[int(6)] = _dafny.CodePoint('s')
                nw147_[int(7)] = (self).f5
                nw147_[int(8)] = (self).f5
                nw147_[int(9)] = default__.fm0(p0, p0, globalState)
                nw147_[int(10)] = (self).f5
                nw147_[int(11)] = (self).f5
                nw147_[int(12)] = _dafny.CodePoint('k')
                nw147_[int(13)] = (self).f5
                nw147_[int(14)] = default__.fm0(p1, (0) - (p1), globalState)
                nw147_[int(15)] = _dafny.CodePoint('g')
                nw147_[int(16)] = default__.fm0(p0, (0) - (p1), globalState)
                nw147_[int(17)] = ((d_639_v73_)[default__.safeIndex(693, (d_639_v73_).length(0))])[default__.safeIndex(len(d_547_v1_), len((d_639_v73_)[default__.safeIndex(693, (d_639_v73_).length(0))]))]
                nw147_[int(18)] = (self).f5
                d_646_v78_ = nw147_
                index106_ = default__.safeIndex(720, (d_646_v78_).length(0))
                (d_646_v78_)[index106_] = (d_640_v74_)[default__.safeIndex(p1, len(d_640_v74_))]
                index107_ = default__.safeIndex(720, (d_646_v78_).length(0))
                (d_646_v78_)[index107_] = (self).f5
                index108_ = default__.safeIndex(910, (d_549_v3_).length(0))
                (d_549_v3_)[index108_] = d_546_v0_
                index109_ = default__.safeIndex(259, (d_608_v47_).length(0))
                (d_608_v47_)[index109_] = len((d_640_v74_) + (d_640_v74_))
                index110_ = default__.safeIndex(259, (d_608_v47_).length(0))
                (d_608_v47_)[index110_] = (p1) * (len((d_639_v73_)[default__.safeIndex(693, (d_639_v73_).length(0))]))
                r3 = (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]
        hi6_ = (self).fm4(d_547_v1_, globalState)
        for d_647_i8_ in range(p0, hi6_):
            d_648_v79_: _dafny.Seq
            d_648_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
            d_649_v80_: D1
            d_649_v80_ = D1_DC2(len(d_648_v79_), p0)
            d_650_v81_: _dafny.Seq
            d_650_v81_ = _dafny.SeqWithoutIsStrInference([(d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))]])
            d_651_v82_: D7
            d_651_v82_ = D7_DC16((self).f5)
            d_652_v83_: _dafny.Array
            nw148_ = _dafny.Array(None, 24)
            nw148_[int(0)] = (self).f5
            nw148_[int(1)] = _dafny.CodePoint('t')
            nw148_[int(2)] = (self).f5
            nw148_[int(3)] = (self).f5
            nw148_[int(4)] = _dafny.CodePoint('g')
            nw148_[int(5)] = (self).f5
            nw148_[int(6)] = (d_651_v82_).cf23
            nw148_[int(7)] = (self).f5
            nw148_[int(8)] = (self).f5
            nw148_[int(9)] = (self).f5
            nw148_[int(10)] = (self).f5
            nw148_[int(11)] = (self).f5
            nw148_[int(12)] = (self).f5
            nw148_[int(13)] = (self).f5
            nw148_[int(14)] = _dafny.CodePoint('u')
            nw148_[int(15)] = (self).f5
            nw148_[int(16)] = (self).f5
            nw148_[int(17)] = (self).f5
            nw148_[int(18)] = (self).f5
            nw148_[int(19)] = _dafny.CodePoint('n')
            nw148_[int(20)] = (self).f5
            nw148_[int(21)] = (self).f5
            nw148_[int(22)] = (self).f5
            nw148_[int(23)] = (self).f5
            d_652_v83_ = nw148_
            d_653_v84_: _dafny.Seq
            d_653_v84_ = _dafny.SeqWithoutIsStrInference([(D10_DC22((D5_DC13(d_648_v79_, d_647_i8_, (d_549_v3_)[default__.safeIndex(910, (d_549_v3_).length(0))])).cf16, d_652_v83_, 239)).cf33, d_648_v79_, _dafny.SeqWithoutIsStrInference([(self).f5])])
            d_654_v85_: _dafny.MultiSet
            d_654_v85_ = _dafny.MultiSet([len(d_553_v5_), p1, len(d_547_v1_), p0, d_647_i8_])
            d_655_v86_: _dafny.Array
            nw149_ = _dafny.Array(None, 10)
            nw149_[int(0)] = d_647_i8_
            nw149_[int(1)] = p1
            nw149_[int(2)] = p1
            nw149_[int(3)] = (d_649_v80_).cf2
            nw149_[int(4)] = default__.safeDivisionInt(415, (d_547_v1_)[default__.safeIndex(len(d_650_v81_), len(d_547_v1_))])
            nw149_[int(5)] = len(d_653_v84_)
            nw149_[int(6)] = ((d_654_v85_).cardinality) + (len(d_547_v1_))
            nw149_[int(7)] = p1
            nw149_[int(8)] = (self).fm4(d_547_v1_, globalState)
            nw149_[int(9)] = (0) - (len((d_648_v79_) + ((d_653_v84_)[default__.safeIndex((0) - (d_647_i8_), len(d_653_v84_))])))
            d_655_v86_ = nw149_
            d_655_v86_ = d_655_v86_
            r1 = d_647_i8_
            index111_ = default__.safeIndex(910, (d_549_v3_).length(0))
            (d_549_v3_)[index111_] = (361) == (p1)
            index112_ = default__.safeIndex(910, (d_549_v3_).length(0))
            (d_549_v3_)[index112_] = (self).fm3(p1, p1, globalState)
        d_656_v87_: _dafny.Seq
        d_656_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwenu"))
        d_657_v88_: _dafny.Set
        d_657_v88_ = _dafny.Set({(0) - (len(d_656_v87_))})
        d_658_v89_: _dafny.Map
        d_658_v89_ = _dafny.Map({d_657_v88_: p0})
        r3 = (_dafny.Map({d_657_v88_: p1})) != (d_658_v89_)
        d_659_v90_: D1
        d_659_v90_ = D1_DC1(d_548_v2_)
        d_659_v90_ = D1_DC1(d_548_v2_)
        d_660_v91_: D8
        d_660_v91_ = D8_DC19(p1, d_546_v0_)
        pat_let_tv22_ = d_546_v0_
        index113_ = default__.safeIndex(910, (d_549_v3_).length(0))
        def lambda37_(source10_):
            if source10_.is_DC19:
                d_661___mcc_h3_ = source10_.cf29
                d_662___mcc_h4_ = source10_.cf30
                d_663_cf30_ = d_662___mcc_h4_
                d_664_cf29_ = d_661___mcc_h3_
                return d_663_cf30_
            elif True:
                d_665___mcc_h5_ = source10_.cf28
                d_666_cf28_ = d_665___mcc_h5_
                return pat_let_tv22_

        (d_549_v3_)[index113_] = lambda37_(d_660_v91_)
        d_667_v92_: _dafny.MultiSet
        d_667_v92_ = _dafny.MultiSet([d_547_v1_])
        d_668_v93_: _dafny.Seq
        d_668_v93_ = _dafny.SeqWithoutIsStrInference([d_667_v92_, d_667_v92_])
        r0 = ((d_668_v93_)[default__.safeIndex(p1, len(d_668_v93_))]).cardinality
        r1 = p1
        r2 = (0) - ((p1) - ((0) - (p1)))
        r3 = d_546_v0_
        return r0, r1, r2, r3


class C2(T1, T0):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f4, f5):
        (self)._f4 = f4
        (self)._f5 = f5

    def fm4(self, p0, globalState):
        return (default__.safeModuloInt(916, len(_dafny.Map({False: (_dafny.MultiSet([(self).f5, (self).f5, (self).f5])).cardinality})))) + (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), (self).f5])) + (_dafny.SeqWithoutIsStrInference([(self).f5]))))

    def fm5(self, p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([True, True, True, False, False])).intersection(_dafny.MultiSet([True]))).isdisjoint((_dafny.MultiSet([False, True, False])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))))

    def fm3(self, p0, p1, globalState):
        return (803) < (default__.safeModuloInt(320, 224))

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        d_669_v0_: bool
        d_669_v0_ = False
        d_670_v1_: _dafny.Array
        nw150_ = _dafny.Array(False, 11)
        d_670_v1_ = nw150_
        index114_ = default__.safeIndex(122, (d_670_v1_).length(0))
        (d_670_v1_)[index114_] = d_669_v0_
        d_671_v2_: int
        d_671_v2_ = 559
        d_672_v3_: _dafny.Seq
        d_672_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrpoqg"))
        d_673_v4_: _dafny.MultiSet
        d_673_v4_ = _dafny.MultiSet([d_671_v2_, (d_671_v2_ if not(default__.fm1(d_671_v2_, d_671_v2_, d_669_v0_, d_669_v0_, globalState)) else (self).fm4(_dafny.SeqWithoutIsStrInference([(0) - (len(d_672_v3_)) for d_674_i0_ in range(default__.abs(127))]), globalState)), (d_671_v2_) - (len(d_672_v3_))])
        d_675_v5_: _dafny.Set
        d_675_v5_ = _dafny.Set({d_669_v0_})
        index115_ = default__.safeIndex(122, (d_670_v1_).length(0))
        rhs127_ = (default__.safeDivisionInt(d_671_v2_, len(d_672_v3_))) < (((default__.fm15(d_669_v0_, d_669_v0_, d_671_v2_, globalState)).cardinality) - (-447))
        rhs128_ = ((d_673_v4_)[len(d_675_v5_)] if (len(d_675_v5_)) in (d_673_v4_) else (d_671_v2_) * (d_671_v2_))
        rhs129_ = d_671_v2_
        rhs130_ = ((self).f5) not in (d_672_v3_)
        lhs64_ = d_670_v1_
        lhs65_ = default__.safeIndex(122, (d_670_v1_).length(0))
        d_669_v0_ = rhs127_
        r1 = rhs128_
        r1 = rhs129_
        lhs64_[lhs65_] = rhs130_
        d_676_v6_: _dafny.MultiSet
        d_676_v6_ = _dafny.MultiSet([d_669_v0_])
        if (d_669_v0_) and (((0) - ((d_676_v6_).cardinality)) != (d_671_v2_)):
            d_671_v2_ = (0) - (d_671_v2_)
            d_677_v7_: _dafny.Array
            def lambda38_(d_678_i1_):
                return default__.safeDivisionInt(d_678_i1_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f5 for d_679_i2_ in range(default__.abs(909))]))).cardinality)

            init21_ = lambda38_
            nw151_ = _dafny.Array(None, 24)
            for i0_21_ in range(nw151_.length(0)):
                nw151_[i0_21_] = init21_(i0_21_)
            d_677_v7_ = nw151_
            index116_ = default__.safeIndex(641, (d_677_v7_).length(0))
            (d_677_v7_)[index116_] = d_671_v2_
            d_680_v8_: _dafny.Set
            d_680_v8_ = _dafny.Set({d_670_v1_, d_670_v1_})
            d_681_v9_: _dafny.Set
            d_681_v9_ = d_680_v8_
            d_682_v10_: _dafny.Map
            d_682_v10_ = _dafny.Map({(0) - (d_671_v2_): (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]})
            d_683_v11_: _dafny.Array
            nw152_ = _dafny.Array(None, 7)
            nw152_[int(0)] = d_682_v10_
            nw152_[int(1)] = d_682_v10_
            nw152_[int(2)] = d_682_v10_
            nw152_[int(3)] = d_682_v10_
            nw152_[int(4)] = d_682_v10_
            nw152_[int(5)] = _dafny.Map({len(d_672_v3_): d_669_v0_})
            nw152_[int(6)] = d_682_v10_
            d_683_v11_ = nw152_
            d_684_v12_: D4
            d_684_v12_ = D4_DC10(d_681_v9_, d_671_v2_, d_672_v3_, d_683_v11_)
            index117_ = default__.safeIndex(641, (d_677_v7_).length(0))
            (d_677_v7_)[index117_] = (0) - ((d_684_v12_).cf12)
            d_685_v13_: _dafny.Set
            d_685_v13_ = _dafny.Set({d_671_v2_, -787, -779, 99})
            d_671_v2_ = (0) - (default__.safeDivisionInt(len(d_685_v13_), d_671_v2_))
            index118_ = default__.safeIndex(122, (d_670_v1_).length(0))
            (d_670_v1_)[index118_] = d_669_v0_
            if (self).fm3(d_671_v2_, (d_677_v7_)[default__.safeIndex(641, (d_677_v7_).length(0))], globalState):
                (globalState).f2 = _dafny.SeqWithoutIsStrInference([d_671_v2_ for d_686_i3_ in range(default__.abs(724))])
                d_687_v14_: _dafny.Seq
                d_687_v14_ = _dafny.SeqWithoutIsStrInference([False, (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]])
                index119_ = default__.safeIndex(122, (d_670_v1_).length(0))
                index120_ = default__.safeIndex(122, (d_670_v1_).length(0))
                rhs131_ = (d_687_v14_) != (_dafny.SeqWithoutIsStrInference([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]]))
                rhs132_ = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
                rhs133_ = d_677_v7_
                rhs134_ = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
                rhs135_ = ((0) - (d_671_v2_) if d_669_v0_ else d_671_v2_)
                lhs66_ = d_670_v1_
                lhs67_ = default__.safeIndex(122, (d_670_v1_).length(0))
                lhs68_ = d_670_v1_
                lhs69_ = default__.safeIndex(122, (d_670_v1_).length(0))
                lhs66_[lhs67_] = rhs131_
                d_669_v0_ = rhs132_
                d_677_v7_ = rhs133_
                lhs68_[lhs69_] = rhs134_
                r1 = rhs135_
                d_688_v15_: _dafny.Set
                d_688_v15_ = _dafny.Set({d_672_v3_, default__.fm16((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_671_v2_, globalState)})
                d_688_v15_ = _dafny.Set({(d_672_v3_).set(default__.safeIndex((d_677_v7_)[default__.safeIndex(641, (d_677_v7_).length(0))], len(d_672_v3_)), (self).f5)})
                d_689_v16_: _dafny.Array
                nw153_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_689_v16_ = nw153_
                index121_ = default__.safeIndex(66, (d_689_v16_).length(0))
                (d_689_v16_)[index121_] = d_670_v1_
                index122_ = default__.safeIndex(66, (d_689_v16_).length(0))
                (d_689_v16_)[index122_] = d_670_v1_
                d_690_v17_: C0
                nw154_ = C0()
                nw154_.ctor__()
                d_690_v17_ = nw154_
            elif True:
                index123_ = default__.safeIndex(641, (d_677_v7_).length(0))
                (d_677_v7_)[index123_] = ((0) - ((d_677_v7_)[default__.safeIndex(641, (d_677_v7_).length(0))])) * (default__.safeModuloInt(d_671_v2_, d_671_v2_))
                d_691_v18_: D5
                d_691_v18_ = D5_DC13(d_672_v3_, d_671_v2_, (d_669_v0_) or ((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]))
                d_692_v19_: _dafny.Seq
                d_692_v19_ = _dafny.SeqWithoutIsStrInference([d_670_v1_, d_670_v1_])
                pat_let_tv23_ = d_672_v3_
                index124_ = default__.safeIndex(122, (d_670_v1_).length(0))
                def iife45_(_pat_let10_0):
                    def iife46_(d_693_dt__update__tmp_h0_):
                        def iife47_(_pat_let11_0):
                            def iife48_(d_694_dt__update_hcf16_h0_):
                                return D5_DC13(d_694_dt__update_hcf16_h0_, (d_693_dt__update__tmp_h0_).cf17, (d_693_dt__update__tmp_h0_).cf18)
                            return iife48_(_pat_let11_0)
                        return iife47_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyfu")))
                    return iife46_(_pat_let10_0)
                def iife44_(_pat_let9_0):
                    def iife49_(d_695_dt__update__tmp_h1_):
                        def iife50_(_pat_let12_0):
                            def iife51_(d_696_dt__update_hcf16_h1_):
                                return D5_DC13(d_696_dt__update_hcf16_h1_, (d_695_dt__update__tmp_h1_).cf17, (d_695_dt__update__tmp_h1_).cf18)
                            return iife51_(_pat_let12_0)
                        return iife50_(pat_let_tv23_)
                    return iife49_(_pat_let9_0)
                rhs136_ = False
                rhs137_ = iife44_(iife45_(d_691_v18_))
                rhs138_ = _dafny.SeqWithoutIsStrInference([d_670_v1_, d_670_v1_])
                rhs139_ = (D6_DC15((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_669_v0_, False)).cf20
                lhs70_ = d_670_v1_
                lhs71_ = default__.safeIndex(122, (d_670_v1_).length(0))
                d_669_v0_ = rhs136_
                d_691_v18_ = rhs137_
                d_692_v19_ = rhs138_
                lhs70_[lhs71_] = rhs139_
                d_697_v20_: _dafny.Array
                nw155_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                d_697_v20_ = nw155_
                index125_ = default__.safeIndex(566, (d_697_v20_).length(0))
                (d_697_v20_)[index125_] = (self).f5
                d_698_v21_: _dafny.Seq
                d_698_v21_ = _dafny.SeqWithoutIsStrInference([d_669_v0_, d_669_v0_, (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]])
                d_699_v22_: _dafny.Map
                d_699_v22_ = _dafny.Map({(d_698_v21_) <= (d_698_v21_): d_669_v0_})
                d_700_v23_: _dafny.Seq
                d_700_v23_ = _dafny.SeqWithoutIsStrInference([-282])
                index126_ = default__.safeIndex(566, (d_697_v20_).length(0))
                index127_ = default__.safeIndex(122, (d_670_v1_).length(0))
                def iife52_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(81, 72):
                        d_702_v24_: int = compr_26_
                        if ((81) <= (d_702_v24_)) and ((d_702_v24_) < (72)):
                            coll26_[default__.safeModuloInt(d_702_v24_, d_671_v2_)] = True
                    return _dafny.Map(coll26_)
                rhs140_ = (self).f5
                rhs141_ = (len(d_700_v23_)) == (d_671_v2_)
                rhs142_ = (d_672_v3_) != (_dafny.SeqWithoutIsStrInference([(self).f5 for d_701_i4_ in range(default__.abs(444))]))
                rhs143_ = ((d_699_v22_) | ((default__.fm19(d_669_v0_, (d_673_v4_).cardinality, ((d_699_v22_)[d_669_v0_] if (d_669_v0_) in (d_699_v22_) else d_669_v0_), globalState)))).set((len(iife52_()
                )) <= (d_671_v2_), ((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]) or ((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]))
                rhs144_ = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
                lhs72_ = d_697_v20_
                lhs73_ = default__.safeIndex(566, (d_697_v20_).length(0))
                lhs74_ = d_670_v1_
                lhs75_ = default__.safeIndex(122, (d_670_v1_).length(0))
                lhs72_[lhs73_] = rhs140_
                d_669_v0_ = rhs141_
                lhs74_[lhs75_] = rhs142_
                d_699_v22_ = rhs143_
                d_669_v0_ = rhs144_
                index128_ = default__.safeIndex(641, (d_677_v7_).length(0))
                (d_677_v7_)[index128_] = len((d_672_v3_) + (d_672_v3_))
                index129_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index129_] = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
        elif True:
            d_703_v25_: C0
            nw156_ = C0()
            nw156_.ctor__()
            d_703_v25_ = nw156_
            d_669_v0_ = d_669_v0_
            d_672_v3_ = _dafny.SeqWithoutIsStrInference([(self).f5 for d_704_i5_ in range(default__.abs(661))])
            d_671_v2_ = d_671_v2_
            d_705_v26_: _dafny.Array
            def lambda39_(d_706_v2_):
                def lambda40_(d_707_i6_):
                    return default__.safeDivisionInt(d_707_i6_, d_706_v2_)

                return lambda40_

            init22_ = lambda39_(d_671_v2_)
            nw157_ = _dafny.Array(None, 3)
            for i0_22_ in range(nw157_.length(0)):
                nw157_[i0_22_] = init22_(i0_22_)
            d_705_v26_ = nw157_
            d_708_v27_: _dafny.Map
            d_708_v27_ = _dafny.Map({D8_DC18(d_673_v4_): d_705_v26_})
            d_709_v28_: _dafny.Map
            d_709_v28_ = _dafny.Map({(default__.fm16(d_669_v0_, d_671_v2_, globalState)).set(default__.safeIndex(len(d_708_v27_), len(default__.fm16(d_669_v0_, d_671_v2_, globalState))), (self).f5): d_669_v0_})
            d_710_v29_: _dafny.Map
            d_710_v29_ = _dafny.Map({(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]: (self).fm3(625, d_671_v2_, globalState)})
            d_711_v30_: _dafny.Map
            d_711_v30_ = _dafny.Map({d_709_v28_: not(not(not (((d_710_v29_)[(self).fm3(d_671_v2_, 1, globalState)] if ((self).fm3(d_671_v2_, 1, globalState)) in (d_710_v29_) else d_669_v0_)) or (d_669_v0_)))})
            d_711_v30_ = (d_711_v30_).set((d_709_v28_) | (d_709_v28_), (d_671_v2_) == ((d_676_v6_).cardinality))
        index130_ = default__.safeIndex(122, (d_670_v1_).length(0))
        (d_670_v1_)[index130_] = not(True)
        d_712_v31_: _dafny.Seq
        d_712_v31_ = _dafny.SeqWithoutIsStrInference([((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] else (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]), (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], False, default__.fm1(d_671_v2_, d_671_v2_, (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_669_v0_, globalState)])
        index131_ = default__.safeIndex(122, (d_670_v1_).length(0))
        (d_670_v1_)[index131_] = not((d_712_v31_)[default__.safeIndex(d_671_v2_, len(d_712_v31_))])
        d_713_v32_: D2
        d_713_v32_ = D2_DC3(d_670_v1_)
        pat_let_tv24_ = d_670_v1_
        def iife53_(_pat_let13_0):
            def iife54_(d_714_dt__update__tmp_h2_):
                def iife55_(_pat_let14_0):
                    def iife56_(d_715_dt__update_hcf4_h0_):
                        return D2_DC3(d_715_dt__update_hcf4_h0_)
                    return iife56_(_pat_let14_0)
                return iife55_(pat_let_tv24_)
            return iife54_(_pat_let13_0)
        source11_ = iife53_(d_713_v32_)
        if source11_.is_DC4:
            d_716___mcc_h0_ = source11_.cf5
            d_717_cf5_ = d_716___mcc_h0_
            d_718_v33_: _dafny.Set
            d_718_v33_ = _dafny.Set({d_671_v2_, d_671_v2_, d_671_v2_})
            d_719_v34_: _dafny.MultiSet
            d_719_v34_ = _dafny.MultiSet([_dafny.Set({d_671_v2_, d_671_v2_})])
            d_720_v35_: _dafny.Map
            d_720_v35_ = _dafny.Map({((_dafny.MultiSet([d_718_v33_])) - (d_719_v34_)).cardinality: 139})
            d_720_v35_ = (d_720_v35_).set(d_671_v2_, (0) - (d_671_v2_))
            d_721_v36_: _dafny.Map
            d_721_v36_ = _dafny.Map({d_672_v3_: d_669_v0_})
            d_722_v37_: _dafny.Seq
            d_722_v37_ = _dafny.SeqWithoutIsStrInference([d_675_v5_])
            index132_ = default__.safeIndex(122, (d_670_v1_).length(0))
            rhs145_ = not(not(((d_722_v37_)[default__.safeIndex(len(d_718_v33_), len(d_722_v37_))]).issubset(_dafny.Set({d_717_cf5_}))))
            rhs146_ = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
            rhs147_ = (_dafny.Map({d_672_v3_: d_717_cf5_})) | (d_721_v36_)
            rhs148_ = d_672_v3_
            lhs76_ = d_670_v1_
            lhs77_ = default__.safeIndex(122, (d_670_v1_).length(0))
            d_717_cf5_ = rhs145_
            lhs76_[lhs77_] = rhs146_
            d_721_v36_ = rhs147_
            d_672_v3_ = rhs148_
            d_723_v38_: _dafny.Array
            nw158_ = _dafny.Array(int(0), 11)
            d_723_v38_ = nw158_
            index133_ = default__.safeIndex(434, (d_723_v38_).length(0))
            (d_723_v38_)[index133_] = (d_671_v2_) + (983)
            d_724_v39_: _dafny.Map
            d_724_v39_ = _dafny.Map({(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]: -502})
            index134_ = default__.safeIndex(434, (d_723_v38_).length(0))
            (d_723_v38_)[index134_] = (len((d_724_v39_) | (d_724_v39_))) + (len(d_672_v3_))
            d_725_v40_: _dafny.Seq
            d_725_v40_ = _dafny.SeqWithoutIsStrInference([d_676_v6_, d_676_v6_, d_676_v6_, d_676_v6_])
            d_726_v41_: _dafny.Seq
            d_726_v41_ = _dafny.SeqWithoutIsStrInference([(d_723_v38_)[default__.safeIndex(434, (d_723_v38_).length(0))], (d_723_v38_)[default__.safeIndex(434, (d_723_v38_).length(0))], 361])
            d_727_v42_: _dafny.Seq
            d_727_v42_ = _dafny.SeqWithoutIsStrInference([d_726_v41_, d_726_v41_, _dafny.SeqWithoutIsStrInference([547])])
            index135_ = default__.safeIndex(434, (d_723_v38_).length(0))
            rhs149_ = (277) >= ((d_723_v38_)[default__.safeIndex(434, (d_723_v38_).length(0))])
            rhs150_ = d_671_v2_
            rhs151_ = ((d_723_v38_)[default__.safeIndex(434, (d_723_v38_).length(0))]) + (len((d_725_v40_) + (default__.fm20(d_671_v2_, (d_727_v42_)[default__.safeIndex(d_671_v2_, len(d_727_v42_))], (self).f5, globalState))))
            lhs78_ = d_723_v38_
            lhs79_ = default__.safeIndex(434, (d_723_v38_).length(0))
            d_669_v0_ = rhs149_
            r1 = rhs150_
            lhs78_[lhs79_] = rhs151_
        elif source11_.is_DC5:
            d_728___mcc_h1_ = source11_.cf6
            d_729_cf6_ = d_728___mcc_h1_
            d_730_v43_: _dafny.Seq
            d_730_v43_ = _dafny.SeqWithoutIsStrInference([d_671_v2_])
            r1 = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccxgxm")) if (self).fm5(len(d_730_v43_), (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_671_v2_, -1, globalState) else (d_729_cf6_).set(default__.safeIndex((self).fm4(d_730_v43_, globalState), len(d_729_cf6_)), (self).f5)))
            d_731_v44_: _dafny.Map
            d_731_v44_ = _dafny.Map({d_712_v31_: d_730_v43_})
            index136_ = default__.safeIndex(122, (d_670_v1_).length(0))
            (d_670_v1_)[index136_] = default__.fm1(d_671_v2_, d_671_v2_, True, (d_671_v2_) not in (((d_731_v44_)[default__.fm21(d_671_v2_, (d_730_v43_)[default__.safeIndex(d_671_v2_, len(d_730_v43_))], d_675_v5_, globalState)] if (default__.fm21(d_671_v2_, (d_730_v43_)[default__.safeIndex(d_671_v2_, len(d_730_v43_))], d_675_v5_, globalState)) in (d_731_v44_) else _dafny.SeqWithoutIsStrInference([d_671_v2_]))), globalState)
            d_672_v3_ = d_729_cf6_
            if False:
                d_732_v45_: _dafny.Array
                nw159_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_732_v45_ = nw159_
                d_733_v46_: _dafny.Map
                d_733_v46_ = _dafny.Map({d_732_v45_: d_669_v0_})
                d_671_v2_ = len((d_733_v46_) | (d_733_v46_))
                d_734_v47_: _dafny.Array
                def lambda41_(d_735_v2_):
                    def lambda42_(d_736_i7_):
                        return (d_736_i7_) - (d_735_v2_)

                    return lambda42_

                init23_ = lambda41_(d_671_v2_)
                nw160_ = _dafny.Array(None, 14)
                for i0_23_ in range(nw160_.length(0)):
                    nw160_[i0_23_] = init23_(i0_23_)
                d_734_v47_ = nw160_
                index137_ = default__.safeIndex(917, (d_734_v47_).length(0))
                (d_734_v47_)[index137_] = (d_671_v2_) * (d_671_v2_)
                index138_ = default__.safeIndex(917, (d_734_v47_).length(0))
                (d_734_v47_)[index138_] = default__.safeModuloInt(d_671_v2_, (self).fm4(d_730_v43_, globalState))
                index139_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index139_] = (d_673_v4_).isdisjoint(_dafny.MultiSet([d_671_v2_]))
                index140_ = default__.safeIndex(917, (d_734_v47_).length(0))
                (d_734_v47_)[index140_] = (d_671_v2_) * ((d_671_v2_ if not(d_669_v0_) else 216))
                d_737_v48_: _dafny.Map
                d_737_v48_ = _dafny.Map({(d_734_v47_)[default__.safeIndex(917, (d_734_v47_).length(0))]: (d_734_v47_)[default__.safeIndex(917, (d_734_v47_).length(0))]})
                d_737_v48_ = (d_737_v48_).set((len(d_672_v3_)) + ((d_734_v47_)[default__.safeIndex(917, (d_734_v47_).length(0))]), len(d_730_v43_))
            elif True:
                d_738_v49_: _dafny.Array
                def lambda43_(d_739_v2_):
                    def lambda44_(d_740_i8_):
                        return (d_740_i8_) * (d_739_v2_)

                    return lambda44_

                init24_ = lambda43_(d_671_v2_)
                nw161_ = _dafny.Array(None, 22)
                for i0_24_ in range(nw161_.length(0)):
                    nw161_[i0_24_] = init24_(i0_24_)
                d_738_v49_ = nw161_
                d_741_v50_: _dafny.Map
                d_741_v50_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))): d_738_v49_})
                d_738_v49_ = (((d_741_v50_)[355] if (355) in (d_741_v50_) else d_738_v49_) if d_669_v0_ else d_738_v49_)
                (globalState).f2 = d_730_v43_
                index141_ = default__.safeIndex(951, (d_738_v49_).length(0))
                (d_738_v49_)[index141_] = d_671_v2_
                d_742_v51_: _dafny.Seq
                d_742_v51_ = _dafny.SeqWithoutIsStrInference([d_738_v49_, d_738_v49_])
                index142_ = default__.safeIndex(951, (d_738_v49_).length(0))
                index143_ = default__.safeIndex(122, (d_670_v1_).length(0))
                index144_ = default__.safeIndex(122, (d_670_v1_).length(0))
                rhs152_ = (d_671_v2_ if d_669_v0_ else 763)
                rhs153_ = d_671_v2_
                rhs154_ = d_671_v2_
                rhs155_ = True
                rhs156_ = (d_742_v51_) == (d_742_v51_)
                lhs80_ = d_738_v49_
                lhs81_ = default__.safeIndex(951, (d_738_v49_).length(0))
                lhs82_ = d_670_v1_
                lhs83_ = default__.safeIndex(122, (d_670_v1_).length(0))
                lhs84_ = d_670_v1_
                lhs85_ = default__.safeIndex(122, (d_670_v1_).length(0))
                lhs80_[lhs81_] = rhs152_
                r1 = rhs153_
                r1 = rhs154_
                lhs82_[lhs83_] = rhs155_
                lhs84_[lhs85_] = rhs156_
                d_743_v52_: _dafny.Map
                d_743_v52_ = _dafny.Map({(d_738_v49_)[default__.safeIndex(951, (d_738_v49_).length(0))]: _dafny.CodePoint('h')})
                d_743_v52_ = (d_743_v52_).set(len((d_712_v31_) + (d_712_v31_)), (self).f5)
                d_744_v53_: C0
                nw162_ = C0()
                nw162_.ctor__()
                d_744_v53_ = nw162_
                d_745_v54_: _dafny.Map
                d_745_v54_ = _dafny.Map({d_669_v0_: d_744_v53_})
                d_746_v55_: bool
                d_746_v55_ = d_669_v0_
                d_745_v54_ = (d_745_v54_).set((d_746_v55_), d_744_v53_)
        elif source11_.is_DC6:
            d_747___mcc_h2_ = source11_.cf7
            d_748_cf7_ = d_747___mcc_h2_
            d_749_v56_: _dafny.Seq
            d_749_v56_ = _dafny.SeqWithoutIsStrInference([d_671_v2_])
            d_750_v57_: D6
            d_750_v57_ = D6_DC14((d_749_v56_ if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] else d_749_v56_))
            d_750_v57_ = d_750_v57_
            d_751_v58_: C0
            nw163_ = C0()
            nw163_.ctor__()
            d_751_v58_ = nw163_
            if d_669_v0_:
                index145_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index145_] = d_669_v0_
                index146_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index146_] = d_669_v0_
                index147_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index147_] = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
                r1 = d_671_v2_
                d_752_v59_: _dafny.Map
                d_752_v59_ = _dafny.Map({(self).f5: ((self).f5 if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] else (self).f5)})
                d_753_v60_: D7
                d_753_v60_ = D7_DC16((self).f5)
                d_752_v59_ = (d_752_v59_).set(((d_753_v60_).cf23 if False else _dafny.CodePoint('b')), default__.fm0(d_748_cf7_, d_671_v2_, globalState))
            elif True:
                d_754_v61_: _dafny.Map
                d_754_v61_ = _dafny.Map({d_671_v2_: (0) - (d_748_cf7_)})
                d_755_v62_: _dafny.Map
                d_755_v62_ = _dafny.Map({d_676_v6_: d_754_v61_})
                d_756_v63_: _dafny.Array
                def lambda45_(d_757_v2_):
                    def lambda46_(d_758_i9_):
                        return (d_758_i9_) - ((0) - (d_757_v2_))

                    return lambda46_

                init25_ = lambda45_(d_671_v2_)
                nw164_ = _dafny.Array(None, 14)
                for i0_25_ in range(nw164_.length(0)):
                    nw164_[i0_25_] = init25_(i0_25_)
                d_756_v63_ = nw164_
                d_759_v64_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_759_v64_ = nw165_
                d_760_v65_: _dafny.Array
                def lambda47_(d_761_v3_):
                    def lambda48_(d_762_i10_):
                        return d_761_v3_

                    return lambda48_

                init26_ = lambda47_(d_672_v3_)
                nw166_ = _dafny.Array(None, 10)
                for i0_26_ in range(nw166_.length(0)):
                    nw166_[i0_26_] = init26_(i0_26_)
                d_760_v65_ = nw166_
                index148_ = default__.safeIndex(394, (d_759_v64_).length(0))
                (d_759_v64_)[index148_] = d_760_v65_
                index149_ = default__.safeIndex(394, (d_759_v64_).length(0))
                rhs157_ = d_748_cf7_
                rhs158_ = d_755_v62_
                rhs159_ = (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]
                rhs160_ = d_756_v63_
                rhs161_ = d_760_v65_
                lhs86_ = d_759_v64_
                lhs87_ = default__.safeIndex(394, (d_759_v64_).length(0))
                r1 = rhs157_
                d_755_v62_ = rhs158_
                d_669_v0_ = rhs159_
                d_756_v63_ = rhs160_
                lhs86_[lhs87_] = rhs161_
                index150_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index150_] = (((d_676_v6_) | (d_676_v6_)).cardinality) < ((d_671_v2_) * (d_671_v2_))
                d_671_v2_ = d_671_v2_
                d_763_v66_: _dafny.Map
                d_763_v66_ = _dafny.Map({(d_672_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fusnvwurp"))): (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]})
                d_763_v66_ = d_763_v66_
                index151_ = default__.safeIndex(117, (d_756_v63_).length(0))
                (d_756_v63_)[index151_] = len(d_672_v3_)
                index152_ = default__.safeIndex(117, (d_756_v63_).length(0))
                (d_756_v63_)[index152_] = (len(d_712_v31_)) + (d_671_v2_)
            d_764_v67_: D8
            d_764_v67_ = D8_DC18(d_673_v4_)
            d_765_v68_: _dafny.Seq
            d_765_v68_ = _dafny.SeqWithoutIsStrInference([d_673_v4_, d_673_v4_])
            d_766_v69_: _dafny.Map
            d_766_v69_ = _dafny.Map({d_676_v6_: (0) - (d_748_cf7_)})
            d_767_v70_: _dafny.Array
            nw167_ = _dafny.Array(None, 10)
            nw167_[int(0)] = (d_673_v4_) - ((d_764_v67_).cf28)
            nw167_[int(1)] = d_673_v4_
            nw167_[int(2)] = _dafny.MultiSet([(d_673_v4_).cardinality, (0) - ((d_749_v56_)[default__.safeIndex(d_671_v2_, len(d_749_v56_))]), d_748_cf7_])
            nw167_[int(3)] = d_673_v4_
            nw167_[int(4)] = (d_673_v4_).set(404, default__.abs(d_671_v2_))
            nw167_[int(5)] = d_673_v4_
            nw167_[int(6)] = d_673_v4_
            nw167_[int(7)] = (d_765_v68_)[default__.safeIndex(d_671_v2_, len(d_765_v68_))]
            nw167_[int(8)] = d_673_v4_
            nw167_[int(9)] = (d_673_v4_).set(d_748_cf7_, default__.abs(((d_766_v69_)[_dafny.MultiSet([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_669_v0_])] if (_dafny.MultiSet([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_669_v0_])) in (d_766_v69_) else d_748_cf7_)))
            d_767_v70_ = nw167_
            index153_ = default__.safeIndex(710, (d_767_v70_).length(0))
            (d_767_v70_)[index153_] = d_673_v4_
            index154_ = default__.safeIndex(710, (d_767_v70_).length(0))
            (d_767_v70_)[index154_] = _dafny.MultiSet([((d_676_v6_)[d_669_v0_] if (d_669_v0_) in (d_676_v6_) else d_671_v2_), ((d_676_v6_).cardinality if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] else d_671_v2_)])
        elif source11_.is_DC3:
            d_768___mcc_h3_ = source11_.cf4
            d_769_cf4_ = d_768___mcc_h3_
            d_770_v71_: _dafny.Map
            d_770_v71_ = _dafny.Map({((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]) and ((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]): d_669_v0_})
            d_770_v71_ = (d_770_v71_).set(default__.fm1(d_671_v2_, d_671_v2_, (self).fm3(d_671_v2_, d_671_v2_, globalState), d_669_v0_, globalState), d_669_v0_)
            d_771_v72_: _dafny.Seq
            d_771_v72_ = _dafny.SeqWithoutIsStrInference([d_672_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwrcaq"))])
            d_772_v73_: D5
            d_772_v73_ = D5_DC13((d_771_v72_)[default__.safeIndex(len(d_675_v5_), len(d_771_v72_))], d_671_v2_, (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))])
            d_773_v74_: _dafny.Map
            d_773_v74_ = _dafny.Map({d_671_v2_: (d_772_v73_).cf17})
            d_773_v74_ = (d_773_v74_).set(default__.safeDivisionInt(d_671_v2_, len(_dafny.SeqWithoutIsStrInference([len(d_675_v5_)]))), default__.safeModuloInt(802, d_671_v2_))
            if d_669_v0_:
                d_774_v75_: D1
                d_774_v75_ = D1_DC1(d_676_v6_)
                d_774_v75_ = (D1_DC1(d_676_v6_) if not((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]) else d_774_v75_)
                d_775_v76_: _dafny.Array
                def lambda49_(d_776_v2_):
                    def lambda50_(d_777_i11_):
                        return (d_777_i11_) - ((_dafny.MultiSet([d_776_v2_])).cardinality)

                    return lambda50_

                init27_ = lambda49_(d_671_v2_)
                nw168_ = _dafny.Array(None, 20)
                for i0_27_ in range(nw168_.length(0)):
                    nw168_[i0_27_] = init27_(i0_27_)
                d_775_v76_ = nw168_
                index155_ = default__.safeIndex(648, (d_775_v76_).length(0))
                (d_775_v76_)[index155_] = d_671_v2_
                index156_ = default__.safeIndex(648, (d_775_v76_).length(0))
                (d_775_v76_)[index156_] = default__.safeDivisionInt(d_671_v2_, default__.safeModuloInt(d_671_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_778_i12_ in range(default__.abs(679))]))))
                d_669_v0_ = d_669_v0_
                d_773_v74_ = (d_773_v74_).set((d_775_v76_)[default__.safeIndex(648, (d_775_v76_).length(0))], (d_775_v76_)[default__.safeIndex(648, (d_775_v76_).length(0))])
                index157_ = default__.safeIndex(648, (d_775_v76_).length(0))
                (d_775_v76_)[index157_] = ((d_775_v76_)[default__.safeIndex(648, (d_775_v76_).length(0))] if not((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]) else len((d_712_v31_) + (d_712_v31_)))
            elif True:
                d_779_v77_: D10
                d_779_v77_ = D10_DC21(_dafny.Set({d_669_v0_, d_669_v0_}))
                index158_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index158_] = (default__.safeModuloInt(d_671_v2_, len((d_779_v77_).cf32))) == ((667) + (len(d_773_v74_)))
                d_780_v79_: _dafny.Array
                def lambda51_(d_781_v2_):
                    def lambda52_(d_782_i13_):
                        def iife57_():
                            coll27_ = _dafny.Set()
                            compr_27_: int
                            for compr_27_ in (_dafny.SeqWithoutIsStrInference([d_781_v2_ for d_783_i14_ in range(default__.abs(11))])).Elements:
                                d_784_v78_: int = compr_27_
                                if (d_784_v78_) in (_dafny.SeqWithoutIsStrInference([d_781_v2_ for d_783_i14_ in range(default__.abs(11))])):
                                    coll27_ = coll27_.union(_dafny.Set([(d_784_v78_) * (len(_dafny.Map({len(_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykjtoq")))})): True})): 792})))]))
                            return _dafny.Set(coll27_)
                        return iife57_()
                        

                    return lambda52_

                init28_ = lambda51_(d_671_v2_)
                nw169_ = _dafny.Array(None, 23)
                for i0_28_ in range(nw169_.length(0)):
                    nw169_[i0_28_] = init28_(i0_28_)
                d_780_v79_ = nw169_
                d_785_v80_: D11
                d_785_v80_ = D11_DC24(d_780_v79_)
                d_786_v81_: _dafny.Seq
                d_786_v81_ = _dafny.SeqWithoutIsStrInference([d_780_v79_, d_780_v79_, d_780_v79_, d_780_v79_, d_780_v79_])
                d_787_v82_: _dafny.Array
                nw170_ = _dafny.Array(None, 25)
                nw170_[int(0)] = d_780_v79_
                nw170_[int(1)] = d_780_v79_
                nw170_[int(2)] = d_780_v79_
                nw170_[int(3)] = d_780_v79_
                nw170_[int(4)] = d_780_v79_
                nw170_[int(5)] = d_780_v79_
                nw170_[int(6)] = d_780_v79_
                nw170_[int(7)] = d_780_v79_
                nw170_[int(8)] = d_780_v79_
                nw170_[int(9)] = (d_785_v80_).cf37
                nw170_[int(10)] = (d_780_v79_ if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] else d_780_v79_)
                nw170_[int(11)] = d_780_v79_
                nw170_[int(12)] = d_780_v79_
                nw170_[int(13)] = d_780_v79_
                nw170_[int(14)] = d_780_v79_
                nw170_[int(15)] = d_780_v79_
                nw170_[int(16)] = (d_786_v81_)[default__.safeIndex((0) - (d_671_v2_), len(d_786_v81_))]
                nw170_[int(17)] = d_780_v79_
                nw170_[int(18)] = d_780_v79_
                nw170_[int(19)] = d_780_v79_
                nw170_[int(20)] = d_780_v79_
                nw170_[int(21)] = (d_785_v80_).cf37
                nw170_[int(22)] = d_780_v79_
                nw170_[int(23)] = d_780_v79_
                nw170_[int(24)] = d_780_v79_
                d_787_v82_ = nw170_
                index159_ = default__.safeIndex(215, (d_787_v82_).length(0))
                (d_787_v82_)[index159_] = d_780_v79_
                d_788_v83_: _dafny.MultiSet
                d_788_v83_ = _dafny.MultiSet([d_672_v3_, (d_672_v3_) + (d_672_v3_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmaps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))])
                index160_ = default__.safeIndex(215, (d_787_v82_).length(0))
                rhs162_ = d_780_v79_
                rhs163_ = (d_675_v5_).ispropersubset(_dafny.Set({(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]}))
                rhs164_ = d_788_v83_
                rhs165_ = (d_672_v3_) + (((d_672_v3_).set(default__.safeIndex(276, len(d_672_v3_)), (self).f5)) + (d_672_v3_))
                lhs88_ = d_787_v82_
                lhs89_ = default__.safeIndex(215, (d_787_v82_).length(0))
                lhs88_[lhs89_] = rhs162_
                d_669_v0_ = rhs163_
                d_788_v83_ = rhs164_
                d_672_v3_ = rhs165_
                d_789_v84_: str
                d_789_v84_ = _dafny.CodePoint('p')
                d_790_v85_: _dafny.Array
                nw171_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_790_v85_ = nw171_
                index161_ = default__.safeIndex(185, (d_790_v85_).length(0))
                (d_790_v85_)[index161_] = (self).f5
                index162_ = default__.safeIndex(185, (d_790_v85_).length(0))
                rhs166_ = (self).f5
                rhs167_ = default__.fm0(d_671_v2_, d_671_v2_, globalState)
                lhs90_ = d_790_v85_
                lhs91_ = default__.safeIndex(185, (d_790_v85_).length(0))
                d_789_v84_ = rhs166_
                lhs90_[lhs91_] = rhs167_
                d_791_v86_: C0
                nw172_ = C0()
                nw172_.ctor__()
                d_791_v86_ = nw172_
                d_792_v87_: _dafny.Map
                d_792_v87_ = _dafny.Map({d_671_v2_: _dafny.SeqWithoutIsStrInference([d_791_v86_, d_791_v86_])})
                d_793_v88_: _dafny.Map
                d_793_v88_ = _dafny.Map({d_671_v2_: d_792_v87_})
                d_793_v88_ = (d_793_v88_).set(d_671_v2_, d_792_v87_)
                d_671_v2_ = 740
            d_794_v89_: _dafny.Array
            def lambda53_(d_795_i15_):
                return default__.safeModuloInt(d_795_i15_, 445)

            init29_ = lambda53_
            nw173_ = _dafny.Array(None, 16)
            for i0_29_ in range(nw173_.length(0)):
                nw173_[i0_29_] = init29_(i0_29_)
            d_794_v89_ = nw173_
            index163_ = default__.safeIndex(961, (d_794_v89_).length(0))
            (d_794_v89_)[index163_] = 930
            d_796_v90_: D11
            d_796_v90_ = D11_DC25((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_671_v2_, d_794_v89_, d_671_v2_)
            d_797_v91_: _dafny.Array
            nw174_ = _dafny.Array(None, 5)
            nw174_[int(0)] = d_796_v90_
            nw174_[int(1)] = D11_DC25((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))], d_671_v2_, d_794_v89_, d_671_v2_)
            nw174_[int(2)] = d_796_v90_
            nw174_[int(3)] = d_796_v90_
            nw174_[int(4)] = d_796_v90_
            d_797_v91_ = nw174_
            d_798_v92_: _dafny.Map
            d_798_v92_ = _dafny.Map({-171: d_797_v91_})
            d_799_v93_: _dafny.MultiSet
            d_799_v93_ = _dafny.MultiSet([((d_798_v92_)[708] if (708) in (d_798_v92_) else d_797_v91_), d_797_v91_])
            index164_ = default__.safeIndex(961, (d_794_v89_).length(0))
            (d_794_v89_)[index164_] = (((d_799_v93_) | (d_799_v93_)) | (_dafny.MultiSet([d_797_v91_]))).cardinality
        elif True:
            d_800___mcc_h4_ = source11_.cf8
            d_801_cf8_ = d_800___mcc_h4_
            if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]:
                d_802_v94_: C0
                nw175_ = C0()
                nw175_.ctor__()
                d_802_v94_ = nw175_
                nw176_ = C0()
                nw176_.ctor__()
                d_802_v94_ = nw176_
                d_803_v95_: T1
                nw177_ = C1()
                nw177_.ctor__((self).f4, (self).f5)
                d_803_v95_ = nw177_
                d_804_v96_: _dafny.Array
                def lambda54_(d_805_v2_):
                    def lambda55_(d_806_i16_):
                        return default__.safeDivisionInt(d_806_i16_, d_805_v2_)

                    return lambda55_

                init30_ = lambda54_(d_671_v2_)
                nw178_ = _dafny.Array(None, 12)
                for i0_30_ in range(nw178_.length(0)):
                    nw178_[i0_30_] = init30_(i0_30_)
                d_804_v96_ = nw178_
                index165_ = default__.safeIndex(507, (d_804_v96_).length(0))
                (d_804_v96_)[index165_] = 612
                d_807_v97_: _dafny.Seq
                d_807_v97_ = _dafny.SeqWithoutIsStrInference([d_671_v2_])
                index166_ = default__.safeIndex(507, (d_804_v96_).length(0))
                rhs168_ = (0) - ((d_671_v2_ if (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))] else len((d_807_v97_) + (_dafny.SeqWithoutIsStrInference([d_671_v2_])))))
                rhs169_ = (d_803_v95_ if d_669_v0_ else d_803_v95_)
                rhs170_ = (329) + (d_671_v2_)
                rhs171_ = d_671_v2_
                lhs92_ = d_804_v96_
                lhs93_ = default__.safeIndex(507, (d_804_v96_).length(0))
                d_671_v2_ = rhs168_
                d_803_v95_ = rhs169_
                r1 = rhs170_
                lhs92_[lhs93_] = rhs171_
                d_808_v98_: _dafny.Array
                nw179_ = _dafny.Array(None, 16)
                d_808_v98_ = nw179_
                d_809_v99_: _dafny.Array
                d_809_v99_ = d_808_v98_
                d_810_v100_: _dafny.Set
                d_810_v100_ = _dafny.Set({d_808_v98_, d_809_v99_, d_809_v99_})
                index167_ = default__.safeIndex(122, (d_670_v1_).length(0))
                (d_670_v1_)[index167_] = not(((d_810_v100_) - (d_810_v100_)).ispropersubset(_dafny.Set({d_809_v99_})))
                r1 = d_671_v2_
                d_811_v101_: _dafny.Map
                d_811_v101_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f5 for d_812_i17_ in range(default__.abs(-331))])) + (d_672_v3_): (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]})
                index168_ = default__.safeIndex(122, (d_670_v1_).length(0))
                rhs172_ = ((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]) and (d_669_v0_)
                rhs173_ = d_811_v101_
                rhs174_ = (d_676_v6_) - ((d_676_v6_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]]))))
                lhs94_ = d_670_v1_
                lhs95_ = default__.safeIndex(122, (d_670_v1_).length(0))
                lhs94_[lhs95_] = rhs172_
                d_811_v101_ = rhs173_
                d_676_v6_ = rhs174_
            elif True:
                d_813_v102_: C1
                nw180_ = C1()
                nw180_.ctor__(((self).f4 if d_669_v0_ else (self).f4), (d_672_v3_)[default__.safeIndex(d_671_v2_, len(d_672_v3_))])
                d_813_v102_ = nw180_
                d_814_v103_: _dafny.Map
                d_814_v103_ = _dafny.Map({d_671_v2_: d_671_v2_})
                d_815_v104_: _dafny.Map
                d_815_v104_ = _dafny.Map({len(d_675_v5_): d_814_v103_})
                d_816_v107_: _dafny.Map
                d_816_v107_ = _dafny.Map({d_671_v2_: d_815_v104_})
                d_817_v109_: D2
                d_817_v109_ = D2_DC4((d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))])
                d_818_v110_: _dafny.Seq
                d_818_v110_ = _dafny.SeqWithoutIsStrInference([d_671_v2_])
                d_819_v113_: D14
                def iife58_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(-267, 748):
                        d_820_v111_: int = compr_28_
                        if ((-267) <= (d_820_v111_)) and ((d_820_v111_) < (748)):
                            def iife59_():
                                coll29_ = _dafny.Map()
                                compr_29_: int
                                for compr_29_ in _dafny.IntegerRange(570, 540):
                                    d_821_v112_: int = compr_29_
                                    if ((570) <= (d_821_v112_)) and ((d_821_v112_) < (540)):
                                        coll29_[(d_821_v112_) - (294)] = d_671_v2_
                                return _dafny.Map(coll29_)
                            coll28_[default__.safeModuloInt(d_820_v111_, d_671_v2_)] = iife59_()

                    return _dafny.Map(coll28_)
                d_819_v113_ = D14_DC30(iife58_()
)
                d_822_v114_: _dafny.Array
                nw181_ = _dafny.Array(None, 29)
                nw181_[int(0)] = (d_815_v104_) | (d_815_v104_)
                nw181_[int(1)] = (d_815_v104_) | (d_815_v104_)
                nw181_[int(2)] = _dafny.Map({d_671_v2_: d_814_v103_})
                def iife60_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(16, 730):
                        d_823_v105_: int = compr_30_
                        if ((16) <= (d_823_v105_)) and ((d_823_v105_) < (730)):
                            coll30_[(d_823_v105_) - (d_671_v2_)] = d_814_v103_
                    return _dafny.Map(coll30_)
                nw181_[int(3)] = iife60_()
                
                nw181_[int(4)] = d_815_v104_
                nw181_[int(5)] = d_815_v104_
                nw181_[int(6)] = (d_815_v104_) | (d_815_v104_)
                nw181_[int(7)] = d_815_v104_
                nw181_[int(8)] = (d_815_v104_) | (_dafny.Map({len(d_672_v3_): d_814_v103_}))
                nw181_[int(9)] = d_815_v104_
                nw181_[int(10)] = d_815_v104_
                nw181_[int(11)] = _dafny.Map({d_671_v2_: _dafny.Map({d_671_v2_: d_671_v2_})})
                def iife61_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(504, 891):
                        d_824_v106_: int = compr_31_
                        if ((504) <= (d_824_v106_)) and ((d_824_v106_) < (891)):
                            coll31_[default__.safeDivisionInt(d_824_v106_, d_671_v2_)] = d_814_v103_
                    return _dafny.Map(coll31_)
                nw181_[int(12)] = iife61_()
                
                def iife62_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in _dafny.IntegerRange(26, 851):
                        d_825_v108_: int = compr_32_
                        if ((26) <= (d_825_v108_)) and ((d_825_v108_) < (851)):
                            coll32_[(d_825_v108_) * (d_671_v2_)] = d_814_v103_
                    return _dafny.Map(coll32_)
                nw181_[int(13)] = ((d_816_v107_)[d_671_v2_] if (d_671_v2_) in (d_816_v107_) else iife62_()
                )
                nw181_[int(14)] = d_815_v104_
                nw181_[int(15)] = d_815_v104_
                nw181_[int(16)] = d_815_v104_
                nw181_[int(17)] = (d_815_v104_) | (d_815_v104_)
                nw181_[int(18)] = ((d_815_v104_).set(d_671_v2_, d_814_v103_)) | (_dafny.Map({d_671_v2_: d_814_v103_}))
                nw181_[int(19)] = d_815_v104_
                nw181_[int(20)] = (d_815_v104_) | (d_815_v104_)
                nw181_[int(21)] = d_815_v104_
                nw181_[int(22)] = default__.fm25(d_669_v0_, (d_817_v109_).cf5, d_818_v110_, globalState)
                nw181_[int(23)] = (d_815_v104_).set(-582, d_814_v103_)
                nw181_[int(24)] = d_815_v104_
                nw181_[int(25)] = d_815_v104_
                nw181_[int(26)] = (d_819_v113_).cf47
                nw181_[int(27)] = d_815_v104_
                nw181_[int(28)] = (d_815_v104_) | (d_815_v104_)
                d_822_v114_ = nw181_
                index169_ = default__.safeIndex(495, (d_822_v114_).length(0))
                (d_822_v114_)[index169_] = (d_815_v104_) | (d_815_v104_)
                index170_ = default__.safeIndex(495, (d_822_v114_).length(0))
                (d_822_v114_)[index170_] = d_815_v104_
                d_672_v3_ = d_672_v3_
                d_826_v115_: D15
                d_826_v115_ = D15_DC34(d_712_v31_)
                d_827_v116_: _dafny.Array
                nw182_ = _dafny.Array(None, 13)
                nw182_[int(0)] = default__.fm21(len(d_672_v3_), d_671_v2_, d_675_v5_, globalState)
                nw182_[int(1)] = ((d_826_v115_).cf53).set(default__.safeIndex(d_671_v2_, len((d_826_v115_).cf53)), d_669_v0_)
                nw182_[int(2)] = d_712_v31_
                nw182_[int(3)] = ((_dafny.SeqWithoutIsStrInference([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]])).set(default__.safeIndex(d_671_v2_, len(_dafny.SeqWithoutIsStrInference([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]]))), d_669_v0_)) + (d_712_v31_)
                nw182_[int(4)] = (d_712_v31_) + (d_712_v31_)
                nw182_[int(5)] = d_712_v31_
                nw182_[int(6)] = (d_712_v31_) + (d_712_v31_)
                nw182_[int(7)] = d_712_v31_
                nw182_[int(8)] = (d_712_v31_) + (d_712_v31_)
                nw182_[int(9)] = (_dafny.SeqWithoutIsStrInference([d_669_v0_])) + (default__.fm21(d_671_v2_, d_671_v2_, _dafny.Set({d_669_v0_}), globalState))
                nw182_[int(10)] = _dafny.SeqWithoutIsStrInference([(d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]])
                nw182_[int(11)] = (d_712_v31_) + ((d_712_v31_).set(default__.safeIndex(d_671_v2_, len(d_712_v31_)), (d_670_v1_)[default__.safeIndex(122, (d_670_v1_).length(0))]))
                nw182_[int(12)] = d_712_v31_
                d_827_v116_ = nw182_
                index171_ = default__.safeIndex(200, (d_827_v116_).length(0))
                (d_827_v116_)[index171_] = default__.fm21(d_671_v2_, d_671_v2_, d_675_v5_, globalState)
                index172_ = default__.safeIndex(200, (d_827_v116_).length(0))
                (d_827_v116_)[index172_] = d_712_v31_
                r1 = d_671_v2_
            d_828_v117_: _dafny.Array
            def lambda56_(d_829_i18_):
                return (d_829_i18_) + (887)

            init31_ = lambda56_
            nw183_ = _dafny.Array(None, 13)
            for i0_31_ in range(nw183_.length(0)):
                nw183_[i0_31_] = init31_(i0_31_)
            d_828_v117_ = nw183_
            index173_ = default__.safeIndex(449, (d_828_v117_).length(0))
            (d_828_v117_)[index173_] = d_671_v2_
            index174_ = default__.safeIndex(449, (d_828_v117_).length(0))
            (d_828_v117_)[index174_] = d_671_v2_
            d_672_v3_ = d_672_v3_
            r1 = -575
        d_830_v118_: _dafny.Seq
        d_830_v118_ = _dafny.SeqWithoutIsStrInference([(d_712_v31_) + (d_712_v31_)])
        d_830_v118_ = d_830_v118_
        r0 = _dafny.SeqWithoutIsStrInference([((d_673_v4_).intersection(d_673_v4_)).cardinality, default__.safeModuloInt(d_671_v2_, len(_dafny.SeqWithoutIsStrInference([d_671_v2_]))), 270, len(_dafny.Map({d_669_v0_: d_671_v2_}))])
        r1 = d_671_v2_
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_831_v0_: bool
        d_831_v0_ = True
        rhs175_ = False
        rhs176_ = d_831_v0_
        rhs177_ = d_831_v0_
        d_831_v0_ = rhs175_
        d_831_v0_ = rhs176_
        d_831_v0_ = rhs177_
        d_832_v1_: _dafny.Map
        d_832_v1_ = _dafny.Map({d_831_v0_: d_831_v0_})
        d_833_v2_: int
        d_833_v2_ = -346
        d_834_v3_: _dafny.Set
        d_834_v3_ = _dafny.Set({default__.fm0(len(d_832_v1_), d_833_v2_, globalState)})
        d_835_v4_: _dafny.Seq
        d_835_v4_ = _dafny.SeqWithoutIsStrInference([(self).f5])
        def iife63_():
            coll33_ = _dafny.Set()
            compr_33_: str
            for compr_33_ in (d_835_v4_).Elements:
                d_836_v5_: str = compr_33_
                if (d_836_v5_) in (d_835_v4_):
                    coll33_ = coll33_.union(_dafny.Set([d_836_v5_]))
            return _dafny.Set(coll33_)
        if not(((d_834_v3_) - (d_834_v3_)).isdisjoint((iife63_()
        ).intersection(d_834_v3_))):
            d_837_v6_: _dafny.Array
            def lambda57_(d_838_v0_):
                def lambda58_(d_839_i0_):
                    return not (d_838_v0_) or (True)

                return lambda58_

            init32_ = lambda57_(d_831_v0_)
            nw184_ = _dafny.Array(None, 26)
            for i0_32_ in range(nw184_.length(0)):
                nw184_[i0_32_] = init32_(i0_32_)
            d_837_v6_ = nw184_
            index175_ = default__.safeIndex(675, (d_837_v6_).length(0))
            (d_837_v6_)[index175_] = d_831_v0_
            index176_ = default__.safeIndex(675, (d_837_v6_).length(0))
            (d_837_v6_)[index176_] = not(d_831_v0_)
            d_840_v7_: _dafny.Set
            d_840_v7_ = _dafny.Set({d_833_v2_, 804})
            d_841_v8_: _dafny.Seq
            d_841_v8_ = _dafny.SeqWithoutIsStrInference([d_833_v2_])
            d_842_v9_: _dafny.Map
            d_842_v9_ = _dafny.Map({d_840_v7_: d_841_v8_})
            d_842_v9_ = (d_842_v9_) | (d_842_v9_)
            d_833_v2_ = d_833_v2_
            d_843_v10_: _dafny.Array
            def lambda59_(d_844_v2_):
                def lambda60_(d_845_i1_):
                    return (d_845_i1_) * (d_844_v2_)

                return lambda60_

            init33_ = lambda59_(d_833_v2_)
            nw185_ = _dafny.Array(None, 21)
            for i0_33_ in range(nw185_.length(0)):
                nw185_[i0_33_] = init33_(i0_33_)
            d_843_v10_ = nw185_
            index177_ = default__.safeIndex(401, (d_843_v10_).length(0))
            (d_843_v10_)[index177_] = d_833_v2_
            d_846_v11_: _dafny.Set
            d_846_v11_ = _dafny.Set({not(not(d_831_v0_))})
            d_847_v12_: str
            d_847_v12_ = _dafny.CodePoint('c')
            d_848_v13_: D6
            d_848_v13_ = D6_DC15((d_837_v6_)[default__.safeIndex(675, (d_837_v6_).length(0))], d_831_v0_, True)
            d_849_v14_: _dafny.Seq
            d_849_v14_ = _dafny.SeqWithoutIsStrInference([d_831_v0_])
            index178_ = default__.safeIndex(401, (d_843_v10_).length(0))
            index179_ = default__.safeIndex(675, (d_837_v6_).length(0))
            rhs178_ = (d_833_v2_) * (792)
            rhs179_ = (_dafny.Set({(d_848_v13_).cf22, (d_837_v6_)[default__.safeIndex(675, (d_837_v6_).length(0))]})) | (default__.fm26(d_833_v2_, d_833_v2_, globalState))
            rhs180_ = (len(d_849_v14_)) < (d_833_v2_)
            rhs181_ = d_847_v12_
            rhs182_ = d_833_v2_
            lhs96_ = d_843_v10_
            lhs97_ = default__.safeIndex(401, (d_843_v10_).length(0))
            lhs98_ = d_837_v6_
            lhs99_ = default__.safeIndex(675, (d_837_v6_).length(0))
            lhs96_[lhs97_] = rhs178_
            d_846_v11_ = rhs179_
            lhs98_[lhs99_] = rhs180_
            d_847_v12_ = rhs181_
            d_833_v2_ = rhs182_
            d_850_v15_: _dafny.MultiSet
            d_850_v15_ = _dafny.MultiSet([(d_843_v10_)[default__.safeIndex(401, (d_843_v10_).length(0))]])
            d_851_v16_: _dafny.Map
            d_851_v16_ = _dafny.Map({d_833_v2_: (_dafny.MultiSet([d_833_v2_, d_833_v2_])) | (d_850_v15_)})
            d_851_v16_ = (d_851_v16_).set((d_843_v10_)[default__.safeIndex(401, (d_843_v10_).length(0))], d_850_v15_)
        elif True:
            if not(d_831_v0_):
                d_852_v17_: _dafny.Array
                nw186_ = _dafny.Array(None, 3)
                nw186_[int(0)] = d_832_v1_
                nw186_[int(1)] = d_832_v1_
                nw186_[int(2)] = d_832_v1_
                d_852_v17_ = nw186_
                d_853_v18_: _dafny.Map
                d_853_v18_ = _dafny.Map({d_833_v2_: d_852_v17_})
                d_854_v19_: _dafny.Array
                nw187_ = _dafny.Array(None, 22)
                nw187_[int(0)] = d_852_v17_
                nw187_[int(1)] = d_852_v17_
                nw187_[int(2)] = d_852_v17_
                nw187_[int(3)] = d_852_v17_
                nw187_[int(4)] = d_852_v17_
                nw187_[int(5)] = d_852_v17_
                nw187_[int(6)] = d_852_v17_
                nw187_[int(7)] = d_852_v17_
                nw187_[int(8)] = d_852_v17_
                nw187_[int(9)] = d_852_v17_
                nw187_[int(10)] = d_852_v17_
                nw187_[int(11)] = d_852_v17_
                nw187_[int(12)] = d_852_v17_
                nw187_[int(13)] = d_852_v17_
                nw187_[int(14)] = d_852_v17_
                nw187_[int(15)] = d_852_v17_
                nw187_[int(16)] = d_852_v17_
                nw187_[int(17)] = d_852_v17_
                nw187_[int(18)] = ((d_853_v18_)[-100] if (-100) in (d_853_v18_) else d_852_v17_)
                nw187_[int(19)] = d_852_v17_
                nw187_[int(20)] = d_852_v17_
                nw187_[int(21)] = d_852_v17_
                d_854_v19_ = nw187_
                index180_ = default__.safeIndex(638, (d_854_v19_).length(0))
                (d_854_v19_)[index180_] = d_852_v17_
                index181_ = default__.safeIndex(638, (d_854_v19_).length(0))
                (d_854_v19_)[index181_] = d_852_v17_
                d_831_v0_ = ((D7_DC16((self).f5)).cf23) in (_dafny.SeqWithoutIsStrInference([(self).f5 for d_855_i2_ in range(default__.abs(308))]))
                d_856_v20_: _dafny.Seq
                d_856_v20_ = _dafny.SeqWithoutIsStrInference([d_833_v2_])
                d_857_v21_: _dafny.MultiSet
                d_857_v21_ = _dafny.MultiSet([d_856_v20_, d_856_v20_])
                d_858_v22_: _dafny.Set
                d_858_v22_ = _dafny.Set({59})
                d_859_v23_: _dafny.Array
                nw188_ = _dafny.Array(None, 21)
                nw188_[int(0)] = (False if d_831_v0_ else d_831_v0_)
                nw188_[int(1)] = d_831_v0_
                nw188_[int(2)] = default__.fm1(d_833_v2_, (d_857_v21_).cardinality, d_831_v0_, d_831_v0_, globalState)
                nw188_[int(3)] = d_831_v0_
                nw188_[int(4)] = not(not(not(not(d_831_v0_))))
                nw188_[int(5)] = d_831_v0_
                nw188_[int(6)] = d_831_v0_
                nw188_[int(7)] = d_831_v0_
                nw188_[int(8)] = d_831_v0_
                nw188_[int(9)] = d_831_v0_
                nw188_[int(10)] = d_831_v0_
                nw188_[int(11)] = d_831_v0_
                nw188_[int(12)] = not(d_831_v0_)
                nw188_[int(13)] = d_831_v0_
                nw188_[int(14)] = not(d_831_v0_)
                nw188_[int(15)] = d_831_v0_
                nw188_[int(16)] = (d_833_v2_) != (d_833_v2_)
                nw188_[int(17)] = not(d_831_v0_)
                nw188_[int(18)] = d_831_v0_
                nw188_[int(19)] = d_831_v0_
                nw188_[int(20)] = (d_858_v22_) != (_dafny.Set({161}))
                d_859_v23_ = nw188_
                index182_ = default__.safeIndex(138, (d_859_v23_).length(0))
                (d_859_v23_)[index182_] = d_831_v0_
                d_860_v24_: _dafny.MultiSet
                d_860_v24_ = _dafny.MultiSet([(self).f5])
                d_861_v25_: _dafny.Map
                d_861_v25_ = _dafny.Map({d_831_v0_: d_833_v2_})
                index183_ = default__.safeIndex(138, (d_859_v23_).length(0))
                rhs183_ = ((d_860_v24_)[(self).f5] if ((self).f5) in (d_860_v24_) else ((d_861_v25_)[not(d_831_v0_)] if (not(d_831_v0_)) in (d_861_v25_) else (0) - ((0) - (d_833_v2_))))
                rhs184_ = (d_860_v24_).ispropersubset((d_860_v24_).intersection(d_860_v24_))
                lhs100_ = d_859_v23_
                lhs101_ = default__.safeIndex(138, (d_859_v23_).length(0))
                d_833_v2_ = rhs183_
                lhs100_[lhs101_] = rhs184_
                d_833_v2_ = d_833_v2_
                index184_ = default__.safeIndex(138, (d_859_v23_).length(0))
                (d_859_v23_)[index184_] = (d_831_v0_) or ((d_859_v23_)[default__.safeIndex(138, (d_859_v23_).length(0))])
            elif True:
                d_862_v26_: _dafny.Array
                nw189_ = _dafny.Array(int(0), 7)
                d_862_v26_ = nw189_
                d_863_v27_: _dafny.Set
                d_863_v27_ = _dafny.Set({d_862_v26_, d_862_v26_})
                d_863_v27_ = d_863_v27_
                d_831_v0_ = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_864_i3_ in range(default__.abs(369))])) != ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_865_i4_ in range(default__.abs(669))])) + (d_835_v4_))
                d_866_v28_: _dafny.Array
                nw190_ = _dafny.Array(False, 24)
                d_866_v28_ = nw190_
                index185_ = default__.safeIndex(617, (d_866_v28_).length(0))
                (d_866_v28_)[index185_] = not(not (d_831_v0_) or (d_831_v0_))
                d_867_v29_: D2
                d_867_v29_ = D2_DC4(d_831_v0_)
                index186_ = default__.safeIndex(617, (d_866_v28_).length(0))
                (d_866_v28_)[index186_] = (d_831_v0_) and ((d_867_v29_).cf5)
                d_868_v30_: C0
                nw191_ = C0()
                nw191_.ctor__()
                d_868_v30_ = nw191_
                d_869_v31_: D7
                d_869_v31_ = D7_DC16(default__.fm0(d_833_v2_, d_833_v2_, globalState))
                d_870_v32_: _dafny.Seq
                d_870_v32_ = _dafny.SeqWithoutIsStrInference([d_869_v31_, d_869_v31_])
                d_871_v33_: _dafny.MultiSet
                d_871_v33_ = _dafny.MultiSet([d_870_v32_])
                d_871_v33_ = d_871_v33_
            d_872_v34_: _dafny.Array
            nw192_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_872_v34_ = nw192_
            d_873_v35_: _dafny.Array
            nw193_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_873_v35_ = nw193_
            index187_ = default__.safeIndex(216, (d_872_v34_).length(0))
            (d_872_v34_)[index187_] = d_873_v35_
            index188_ = default__.safeIndex(216, (d_872_v34_).length(0))
            (d_872_v34_)[index188_] = (d_873_v35_ if d_831_v0_ else d_873_v35_)
            d_874_v36_: _dafny.Map
            d_874_v36_ = _dafny.Map({d_833_v2_: d_831_v0_})
            d_875_v37_: _dafny.MultiSet
            d_875_v37_ = _dafny.MultiSet([d_833_v2_, d_833_v2_])
            d_876_v38_: _dafny.Array
            nw194_ = _dafny.Array(None, 8)
            nw194_[int(0)] = ((d_874_v36_)[(0) - (d_833_v2_)] if ((0) - (d_833_v2_)) in (d_874_v36_) else d_831_v0_)
            nw194_[int(1)] = (False if d_831_v0_ else not(d_831_v0_))
            nw194_[int(2)] = (d_833_v2_) < (d_833_v2_)
            nw194_[int(3)] = (d_875_v37_).issubset(_dafny.MultiSet([d_833_v2_]))
            nw194_[int(4)] = d_831_v0_
            nw194_[int(5)] = d_831_v0_
            nw194_[int(6)] = not(False)
            nw194_[int(7)] = d_831_v0_
            d_876_v38_ = nw194_
            index189_ = default__.safeIndex(664, (d_876_v38_).length(0))
            (d_876_v38_)[index189_] = d_831_v0_
            index190_ = default__.safeIndex(664, (d_876_v38_).length(0))
            (d_876_v38_)[index190_] = d_831_v0_
            if d_831_v0_:
                d_833_v2_ = d_833_v2_
                index191_ = default__.safeIndex(664, (d_876_v38_).length(0))
                (d_876_v38_)[index191_] = (d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]
                d_877_v39_: D13
                d_877_v39_ = D13_DC29((d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))], not ((d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]) or ((d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]))
                d_877_v39_ = default__.fm27(globalState)
                d_878_v40_: _dafny.Array
                nw195_ = _dafny.Array(_dafny.Set({}), 1)
                d_878_v40_ = nw195_
                index192_ = default__.safeIndex(298, (d_878_v40_).length(0))
                def iife64_():
                    coll34_ = _dafny.Set()
                    compr_34_: str
                    for compr_34_ in (d_834_v3_).Elements:
                        d_879_v41_: str = compr_34_
                        if (d_879_v41_) in (d_834_v3_):
                            coll34_ = coll34_.union(_dafny.Set([d_879_v41_]))
                    return _dafny.Set(coll34_)
                (d_878_v40_)[index192_] = iife64_()
                
                index193_ = default__.safeIndex(298, (d_878_v40_).length(0))
                (d_878_v40_)[index193_] = _dafny.Set({(self).f5})
                d_880_v42_: _dafny.MultiSet
                d_880_v42_ = _dafny.MultiSet([(d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]])
                d_881_v43_: _dafny.Map
                d_881_v43_ = _dafny.Map({(d_880_v42_).cardinality: d_833_v2_})
                d_831_v0_ = (-377) == (len(d_881_v43_))
            elif True:
                d_882_v44_: _dafny.Array
                nw196_ = _dafny.Array(_dafny.Map({}), 14)
                d_882_v44_ = nw196_
                index194_ = default__.safeIndex(458, (d_882_v44_).length(0))
                (d_882_v44_)[index194_] = _dafny.Map({d_835_v4_: d_833_v2_})
                d_883_v46_: _dafny.MultiSet
                d_883_v46_ = _dafny.MultiSet([d_835_v4_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')]), d_835_v4_])
                d_884_v47_: _dafny.Set
                d_884_v47_ = _dafny.Set({(d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]})
                d_885_v48_: _dafny.Map
                d_885_v48_ = _dafny.Map({d_835_v4_: len(d_884_v47_)})
                index195_ = default__.safeIndex(458, (d_882_v44_).length(0))
                def iife65_():
                    coll35_ = _dafny.Map()
                    compr_35_: _dafny.Seq
                    for compr_35_ in (d_883_v46_).Elements:
                        d_886_v45_: _dafny.Seq = compr_35_
                        if (d_886_v45_) in (d_883_v46_):
                            coll35_[d_886_v45_] = len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([(d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]])}))
                    return _dafny.Map(coll35_)
                (d_882_v44_)[index195_] = ((iife65_()
                ) | (d_885_v48_)) | ((d_885_v48_) | (d_885_v48_))
                d_887_v49_: _dafny.Map
                d_887_v49_ = _dafny.Map({d_833_v2_: 623})
                d_888_v50_: _dafny.MultiSet
                d_888_v50_ = _dafny.MultiSet([d_831_v0_])
                d_887_v49_ = (d_887_v49_).set(((d_888_v50_)[(d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]] if ((d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]) in (d_888_v50_) else d_833_v2_), d_833_v2_)
                d_833_v2_ = d_833_v2_
                d_889_v51_: T0
                nw197_ = C0()
                nw197_.ctor__()
                d_889_v51_ = nw197_
                d_890_v52_: _dafny.Map
                d_890_v52_ = _dafny.Map({(len(d_835_v4_) if d_831_v0_ else -133): _dafny.MultiSet([d_833_v2_])})
                rhs185_ = d_889_v51_
                rhs186_ = (113) > (d_833_v2_)
                rhs187_ = d_890_v52_
                rhs188_ = (d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]
                d_889_v51_ = rhs185_
                d_831_v0_ = rhs186_
                d_890_v52_ = rhs187_
                d_831_v0_ = rhs188_
                d_891_v53_: _dafny.Seq
                d_891_v53_ = _dafny.SeqWithoutIsStrInference([d_876_v38_])
                d_891_v53_ = d_891_v53_
            if d_831_v0_:
                d_892_v54_: _dafny.Seq
                d_892_v54_ = _dafny.SeqWithoutIsStrInference([d_833_v2_])
                d_893_v55_: _dafny.Map
                d_893_v55_ = _dafny.Map({(d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]: d_892_v54_})
                d_831_v0_ = ((d_893_v55_) | (d_893_v55_)) == (d_893_v55_)
                d_894_v56_: C1
                nw198_ = C1()
                nw198_.ctor__((self).f4, (self).f5)
                d_894_v56_ = nw198_
                index196_ = default__.safeIndex(664, (d_876_v38_).length(0))
                (d_876_v38_)[index196_] = not((d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))])
                d_895_v57_: _dafny.Array
                def lambda61_(d_896_v4_):
                    def lambda62_(d_897_i5_):
                        return (d_897_i5_) - (len(d_896_v4_))

                    return lambda62_

                init34_ = lambda61_(d_835_v4_)
                nw199_ = _dafny.Array(None, 21)
                for i0_34_ in range(nw199_.length(0)):
                    nw199_[i0_34_] = init34_(i0_34_)
                d_895_v57_ = nw199_
                index197_ = default__.safeIndex(703, (d_895_v57_).length(0))
                (d_895_v57_)[index197_] = (d_833_v2_ if (d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))] else d_833_v2_)
                index198_ = default__.safeIndex(703, (d_895_v57_).length(0))
                (d_895_v57_)[index198_] = default__.safeDivisionInt((d_894_v56_).fm4(_dafny.SeqWithoutIsStrInference([d_833_v2_, d_833_v2_]), globalState), d_833_v2_)
                index199_ = default__.safeIndex(703, (d_895_v57_).length(0))
                (d_895_v57_)[index199_] = (0) - (d_833_v2_)
            elif True:
                index200_ = default__.safeIndex(664, (d_876_v38_).length(0))
                (d_876_v38_)[index200_] = ((d_832_v1_)[d_831_v0_] if (d_831_v0_) in (d_832_v1_) else (d_833_v2_) == (d_833_v2_))
                d_898_v58_: _dafny.Seq
                d_898_v58_ = _dafny.SeqWithoutIsStrInference([d_833_v2_])
                d_899_v59_: _dafny.Seq
                d_899_v59_ = _dafny.SeqWithoutIsStrInference([d_898_v58_, d_898_v58_])
                (globalState).f2 = (d_898_v58_) + (((d_899_v59_)[default__.safeIndex(d_833_v2_, len(d_899_v59_))]) + (d_898_v58_))
                d_835_v4_ = d_835_v4_
                d_900_v60_: _dafny.Array
                nw200_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_900_v60_ = nw200_
                d_901_v61_: _dafny.Array
                nw201_ = _dafny.Array(None, 5)
                d_901_v61_ = nw201_
                index201_ = default__.safeIndex(959, (d_900_v60_).length(0))
                (d_900_v60_)[index201_] = d_901_v61_
                d_902_v62_: _dafny.Array
                d_902_v62_ = d_901_v61_
                index202_ = default__.safeIndex(959, (d_900_v60_).length(0))
                (d_900_v60_)[index202_] = d_902_v62_
                d_903_v63_: _dafny.Map
                d_903_v63_ = _dafny.Map({d_898_v58_: d_831_v0_})
                d_903_v63_ = (d_903_v63_).set(d_898_v58_, not ((d_876_v38_)[default__.safeIndex(664, (d_876_v38_).length(0))]) or (d_831_v0_))
        d_904_v64_: _dafny.Seq
        d_904_v64_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_905_i6_ in range(default__.abs(629))])), d_833_v2_])
        d_833_v2_ = (d_833_v2_) - (default__.safeModuloInt(len(d_904_v64_), d_833_v2_))
        rhs189_ = ((d_831_v0_ if d_831_v0_ else (self).fm3(d_833_v2_, d_833_v2_, globalState))) == (d_831_v0_)
        rhs190_ = d_833_v2_
        d_831_v0_ = rhs189_
        d_833_v2_ = rhs190_
        if not((D14_DC32(d_833_v2_, d_831_v0_, False)).cf51):
            d_833_v2_ = d_833_v2_
            rhs191_ = (d_835_v4_) <= (d_835_v4_)
            rhs192_ = d_833_v2_
            d_831_v0_ = rhs191_
            d_833_v2_ = rhs192_
            d_831_v0_ = d_831_v0_
            d_833_v2_ = d_833_v2_
            d_906_v65_: C1
            nw202_ = C1()
            nw202_.ctor__((self).f4, _dafny.CodePoint('a'))
            d_906_v65_ = nw202_
        elif True:
            d_907_v66_: _dafny.MultiSet
            d_907_v66_ = _dafny.MultiSet([True])
            d_908_v67_: _dafny.Map
            d_908_v67_ = _dafny.Map({d_907_v66_: d_831_v0_})
            d_908_v67_ = (d_908_v67_).set(d_907_v66_, (_dafny.MultiSet([d_831_v0_])).ispropersubset(d_907_v66_))
            d_831_v0_ = (len(d_835_v4_)) != (d_833_v2_)
            d_909_v69_: _dafny.Map
            def iife66_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in (_dafny.Map({d_833_v2_: d_833_v2_})).keys.Elements:
                    d_911_v68_: int = compr_36_
                    if (d_911_v68_) in (_dafny.Map({d_833_v2_: d_833_v2_})):
                        coll36_[(d_911_v68_) - ((0) - (len(d_835_v4_)))] = d_831_v0_
                return _dafny.Map(coll36_)
            d_909_v69_ = _dafny.Map({((d_904_v64_).set(default__.safeIndex(992, len(d_904_v64_)), d_833_v2_)) + (d_904_v64_): (d_833_v2_) + (len((_dafny.SeqWithoutIsStrInference([(self).f5 for d_910_i7_ in range(default__.abs(-714))])).set(default__.safeIndex(len((iife66_()
            ).set(d_833_v2_, d_831_v0_)), len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_912_i7_ in range(default__.abs(-714))]))), (self).f5)))})
            d_913_v70_: _dafny.Map
            d_913_v70_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_831_v0_, d_831_v0_, d_831_v0_])): d_833_v2_})
            d_909_v69_ = (d_909_v69_).set((d_904_v64_ if d_831_v0_ else d_904_v64_), ((self).fm4(_dafny.SeqWithoutIsStrInference([len(d_913_v70_), d_833_v2_]), globalState)) - (d_833_v2_))
            d_831_v0_ = not (d_831_v0_) or (True)
            d_909_v69_ = d_909_v69_
        d_914_v71_: _dafny.Map
        d_914_v71_ = _dafny.Map({(self).f5: d_831_v0_})
        d_915_i8_: int
        d_915_i8_ = 0
        with _dafny.label("3"):
            while ((d_914_v71_)[(d_835_v4_)[default__.safeIndex(d_833_v2_, len(d_835_v4_))]] if ((d_835_v4_)[default__.safeIndex(d_833_v2_, len(d_835_v4_))]) in (d_914_v71_) else d_831_v0_):
                with _dafny.c_label("3"):
                    if (d_915_i8_) >= (100):
                        raise _dafny.Break("3")
                    d_915_i8_ = (d_915_i8_) + (1)
                    d_916_v72_: _dafny.Map
                    d_916_v72_ = _dafny.Map({not(d_831_v0_): d_831_v0_})
                    d_917_v73_: _dafny.Array
                    nw203_ = _dafny.Array(None, 25)
                    nw203_[int(0)] = d_832_v1_
                    nw203_[int(1)] = d_832_v1_
                    nw203_[int(2)] = _dafny.Map({d_831_v0_: d_831_v0_})
                    nw203_[int(3)] = d_832_v1_
                    nw203_[int(4)] = d_832_v1_
                    nw203_[int(5)] = d_832_v1_
                    nw203_[int(6)] = d_832_v1_
                    nw203_[int(7)] = d_832_v1_
                    nw203_[int(8)] = ((d_916_v72_)).set(d_831_v0_, d_831_v0_)
                    nw203_[int(9)] = d_832_v1_
                    nw203_[int(10)] = d_832_v1_
                    nw203_[int(11)] = d_832_v1_
                    nw203_[int(12)] = _dafny.Map({False: d_831_v0_})
                    nw203_[int(13)] = d_832_v1_
                    nw203_[int(14)] = _dafny.Map({d_831_v0_: d_831_v0_})
                    nw203_[int(15)] = d_832_v1_
                    nw203_[int(16)] = (d_832_v1_).set(d_831_v0_, default__.fm1(d_833_v2_, (d_904_v64_)[default__.safeIndex(d_833_v2_, len(d_904_v64_))], False, d_831_v0_, globalState))
                    nw203_[int(17)] = d_832_v1_
                    nw203_[int(18)] = d_832_v1_
                    nw203_[int(19)] = d_832_v1_
                    nw203_[int(20)] = d_832_v1_
                    nw203_[int(21)] = _dafny.Map({True: d_831_v0_})
                    nw203_[int(22)] = (d_832_v1_).set(d_831_v0_, d_831_v0_)
                    nw203_[int(23)] = d_832_v1_
                    nw203_[int(24)] = d_832_v1_
                    d_917_v73_ = nw203_
                    d_918_v74_: int
                    d_919_v75_: _dafny.Array
                    d_920_v76_: _dafny.Map
                    out10_: int
                    out11_: _dafny.Array
                    out12_: _dafny.Map
                    out10_, out11_, out12_ = (self).m13(d_917_v73_, globalState)
                    d_918_v74_ = out10_
                    d_919_v75_ = out11_
                    d_920_v76_ = out12_
                    d_921_v77_: D8
                    d_921_v77_ = D8_DC19(d_833_v2_, d_831_v0_)
                    d_922_v78_: _dafny.Array
                    nw204_ = _dafny.Array(None, 7)
                    nw204_[int(0)] = default__.fm1(648, d_918_v74_, d_831_v0_, d_831_v0_, globalState)
                    nw204_[int(1)] = False
                    nw204_[int(2)] = True
                    nw204_[int(3)] = d_831_v0_
                    nw204_[int(4)] = d_831_v0_
                    nw204_[int(5)] = (d_918_v74_) > ((d_921_v77_).cf29)
                    nw204_[int(6)] = d_831_v0_
                    d_922_v78_ = nw204_
                    index203_ = default__.safeIndex(422, (d_922_v78_).length(0))
                    (d_922_v78_)[index203_] = d_831_v0_
                    index204_ = default__.safeIndex(422, (d_922_v78_).length(0))
                    (d_922_v78_)[index204_] = False
                    d_833_v2_ = ((-428) + ((self).fm4(d_904_v64_, globalState))) * (d_833_v2_)
                    d_923_v79_: _dafny.MultiSet
                    d_923_v79_ = _dafny.MultiSet([(d_922_v78_)[default__.safeIndex(422, (d_922_v78_).length(0))]])
                    d_924_v80_: _dafny.Map
                    d_924_v80_ = _dafny.Map({d_918_v74_: (d_923_v79_).cardinality})
                    d_925_v81_: _dafny.Set
                    d_925_v81_ = _dafny.Set({d_924_v80_, d_924_v80_, d_924_v80_})
                    d_926_v82_: _dafny.Map
                    d_926_v82_ = _dafny.Map({(self).fm5(len(d_925_v81_), d_831_v0_, d_833_v2_, 508, globalState): len(d_904_v64_)})
                    d_927_v83_: _dafny.Seq
                    d_927_v83_ = _dafny.SeqWithoutIsStrInference([d_926_v82_, _dafny.Map({False: d_833_v2_}), d_926_v82_])
                    d_928_v84_: _dafny.Map
                    d_928_v84_ = _dafny.Map({d_918_v74_: d_926_v82_})
                    d_929_v85_: _dafny.Map
                    d_929_v85_ = _dafny.Map({(d_927_v83_ if (d_922_v78_)[default__.safeIndex(422, (d_922_v78_).length(0))] else d_927_v83_): ((d_928_v84_)[(d_923_v79_).cardinality] if ((d_923_v79_).cardinality) in (d_928_v84_) else d_926_v82_)})
                    d_930_v86_: _dafny.Set
                    d_930_v86_ = _dafny.Set({False, d_831_v0_, default__.fm1(d_918_v74_, (0) - (d_833_v2_), True, True, globalState), d_831_v0_, (d_922_v78_)[default__.safeIndex(422, (d_922_v78_).length(0))]})
                    d_929_v85_ = (d_929_v85_).set((d_927_v83_) + ((d_927_v83_).set(default__.safeIndex(d_833_v2_, len(d_927_v83_)), (d_926_v82_).set((d_922_v78_)[default__.safeIndex(422, (d_922_v78_).length(0))], d_833_v2_))), _dafny.Map({d_831_v0_: len(d_930_v86_)}))
                    pass
            pass
        d_931_v87_: _dafny.Array
        nw205_ = _dafny.Array(int(0), 4)
        d_931_v87_ = nw205_
        r0 = d_931_v87_
        return r0

    def m13(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        d_932_v0_: bool
        d_932_v0_ = False
        d_933_v1_: _dafny.Array
        nw206_ = _dafny.Array(None, 22)
        nw206_[int(0)] = d_932_v0_
        nw206_[int(1)] = d_932_v0_
        nw206_[int(2)] = d_932_v0_
        nw206_[int(3)] = d_932_v0_
        nw206_[int(4)] = d_932_v0_
        nw206_[int(5)] = d_932_v0_
        nw206_[int(6)] = d_932_v0_
        nw206_[int(7)] = d_932_v0_
        nw206_[int(8)] = d_932_v0_
        nw206_[int(9)] = d_932_v0_
        nw206_[int(10)] = d_932_v0_
        nw206_[int(11)] = d_932_v0_
        nw206_[int(12)] = d_932_v0_
        nw206_[int(13)] = True
        nw206_[int(14)] = d_932_v0_
        nw206_[int(15)] = d_932_v0_
        nw206_[int(16)] = d_932_v0_
        nw206_[int(17)] = d_932_v0_
        nw206_[int(18)] = d_932_v0_
        nw206_[int(19)] = d_932_v0_
        nw206_[int(20)] = d_932_v0_
        nw206_[int(21)] = d_932_v0_
        d_933_v1_ = nw206_
        (globalState).f3 = (d_933_v1_ if d_932_v0_ else (d_933_v1_ if not(d_932_v0_) else d_933_v1_))
        d_934_v2_: D2
        d_934_v2_ = D2_DC5(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_935_i0_ in range(default__.abs(746))]))
        d_936_v3_: _dafny.Seq
        d_936_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dt"))
        d_937_v4_: int
        d_937_v4_ = 897
        d_938_v5_: D5
        d_938_v5_ = D5_DC13(d_936_v3_, d_937_v4_, d_932_v0_)
        d_934_v2_ = D2_DC5((d_938_v5_).cf16)
        if d_932_v0_:
            d_939_v6_: _dafny.Seq
            d_939_v6_ = _dafny.SeqWithoutIsStrInference([d_933_v1_, d_933_v1_, d_933_v1_, d_933_v1_, d_933_v1_])
            d_932_v0_ = (d_939_v6_) < (d_939_v6_)
            d_940_v7_: _dafny.Array
            nw207_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_940_v7_ = nw207_
            d_941_v8_: _dafny.Array
            def lambda63_(d_942_v4_):
                def lambda64_(d_943_i1_):
                    return default__.safeDivisionInt(d_943_i1_, d_942_v4_)

                return lambda64_

            init35_ = lambda63_(d_937_v4_)
            nw208_ = _dafny.Array(None, 7)
            for i0_35_ in range(nw208_.length(0)):
                nw208_[i0_35_] = init35_(i0_35_)
            d_941_v8_ = nw208_
            index205_ = default__.safeIndex(475, (d_940_v7_).length(0))
            (d_940_v7_)[index205_] = d_941_v8_
            index206_ = default__.safeIndex(475, (d_940_v7_).length(0))
            (d_940_v7_)[index206_] = d_941_v8_
            r0 = d_937_v4_
            d_944_v9_: _dafny.Seq
            d_944_v9_ = _dafny.SeqWithoutIsStrInference([(d_940_v7_)[default__.safeIndex(475, (d_940_v7_).length(0))], (d_940_v7_)[default__.safeIndex(475, (d_940_v7_).length(0))], (d_940_v7_)[default__.safeIndex(475, (d_940_v7_).length(0))]])
            d_941_v8_ = ((d_940_v7_)[default__.safeIndex(475, (d_940_v7_).length(0))] if d_932_v0_ else (d_944_v9_)[default__.safeIndex(d_937_v4_, len(d_944_v9_))])
            r0 = d_937_v4_
        elif True:
            d_945_v10_: D2
            d_945_v10_ = D2_DC4(d_932_v0_)
            index207_ = default__.safeIndex(455, (d_933_v1_).length(0))
            (d_933_v1_)[index207_] = not((d_945_v10_).cf5)
            index208_ = default__.safeIndex(455, (d_933_v1_).length(0))
            (d_933_v1_)[index208_] = default__.fm1(d_937_v4_, (0) - (d_937_v4_), False, False, globalState)
            d_946_v11_: D2
            d_946_v11_ = D2_DC6(d_937_v4_)
            source12_ = D2_DC7(d_946_v11_)
            if source12_.is_DC4:
                d_947___mcc_h0_ = source12_.cf5
                d_948_cf5_ = d_947___mcc_h0_
                d_949_v12_: _dafny.Map
                d_949_v12_ = _dafny.Map({False: (d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]})
                d_949_v12_ = (d_949_v12_).set(d_948_cf5_, d_932_v0_)
                d_950_v13_: _dafny.Seq
                d_950_v13_ = _dafny.SeqWithoutIsStrInference([d_936_v3_, d_936_v3_])
                d_937_v4_ = len((d_950_v13_) + ((d_950_v13_) + (d_950_v13_)))
                d_951_v14_: str
                d_951_v14_ = _dafny.CodePoint('l')
                d_952_v15_: _dafny.Array
                nw209_ = _dafny.Array(D13.default()(), 26)
                d_952_v15_ = nw209_
                d_953_v16_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Seq({}), 4)
                d_953_v16_ = nw210_
                d_954_v17_: D13
                d_954_v17_ = D13_DC28(d_953_v16_)
                d_955_v18_: _dafny.Map
                d_955_v18_ = _dafny.Map({d_937_v4_: d_954_v17_})
                d_956_v19_: D13
                d_956_v19_ = D13_DC29(d_932_v0_, default__.fm1(len(d_955_v18_), (0) - (d_937_v4_), (d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))], False, globalState))
                d_957_v20_: _dafny.Map
                d_957_v20_ = _dafny.Map({d_956_v19_: False})
                d_958_v22_: _dafny.MultiSet
                d_958_v22_ = _dafny.MultiSet([d_957_v20_, d_957_v20_])
                d_959_v23_: _dafny.Array
                nw211_ = _dafny.Array(None, 8)
                def iife67_():
                    coll37_ = _dafny.Map()
                    compr_37_: D13
                    for compr_37_ in (d_957_v20_).keys.Elements:
                        d_960_v21_: D13 = compr_37_
                        if (d_960_v21_) in (d_957_v20_):
                            coll37_[d_960_v21_] = d_948_cf5_
                    return _dafny.Map(coll37_)
                nw211_[int(0)] = _dafny.MultiSet([d_957_v20_, d_957_v20_, iife67_()
                ])
                nw211_[int(1)] = d_958_v22_
                nw211_[int(2)] = (d_958_v22_).intersection(d_958_v22_)
                nw211_[int(3)] = d_958_v22_
                nw211_[int(4)] = d_958_v22_
                nw211_[int(5)] = (d_958_v22_) | (d_958_v22_)
                nw211_[int(6)] = default__.fm28(globalState)
                nw211_[int(7)] = d_958_v22_
                d_959_v23_ = nw211_
                index209_ = default__.safeIndex(828, (d_959_v23_).length(0))
                (d_959_v23_)[index209_] = (default__.fm28(globalState)) - (d_958_v22_)
                d_961_v25_: _dafny.Set
                d_961_v25_ = _dafny.Set({d_956_v19_})
                index210_ = default__.safeIndex(828, (d_959_v23_).length(0))
                def iife68_():
                    coll38_ = _dafny.Map()
                    compr_38_: D13
                    for compr_38_ in (d_961_v25_).Elements:
                        d_962_v24_: D13 = compr_38_
                        if (d_962_v24_) in (d_961_v25_):
                            coll38_[d_962_v24_] = d_932_v0_
                    return _dafny.Map(coll38_)
                rhs193_ = d_951_v14_
                rhs194_ = d_952_v15_
                rhs195_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_957_v20_) | (iife68_()
                )]))
                lhs102_ = d_959_v23_
                lhs103_ = default__.safeIndex(828, (d_959_v23_).length(0))
                d_951_v14_ = rhs193_
                d_952_v15_ = rhs194_
                lhs102_[lhs103_] = rhs195_
                d_948_cf5_ = not(False)
            elif source12_.is_DC5:
                d_963___mcc_h1_ = source12_.cf6
                d_964_cf6_ = d_963___mcc_h1_
                d_965_v26_: _dafny.Seq
                d_965_v26_ = _dafny.SeqWithoutIsStrInference([d_932_v0_])
                d_966_v28_: _dafny.Seq
                d_966_v28_ = _dafny.SeqWithoutIsStrInference([-267])
                index211_ = default__.safeIndex(455, (d_933_v1_).length(0))
                def iife69_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in (d_966_v28_).Elements:
                        d_967_v27_: int = compr_39_
                        if (d_967_v27_) in (d_966_v28_):
                            coll39_[default__.safeModuloInt(d_967_v27_, d_937_v4_)] = 75
                    return _dafny.Map(coll39_)
                (d_933_v1_)[index211_] = (not((d_965_v26_)[default__.safeIndex(d_937_v4_, len(d_965_v26_))])) and ((_dafny.Map({len(d_964_cf6_): 243})) != (iife69_()
                ))
                d_968_v29_: C0
                nw212_ = C0()
                nw212_.ctor__()
                d_968_v29_ = nw212_
                d_966_v28_ = d_966_v28_
                d_969_v30_: _dafny.Set
                d_969_v30_ = _dafny.Set({d_937_v4_})
                d_970_v31_: D8
                d_970_v31_ = D8_DC19(d_937_v4_, (d_937_v4_) < (len(d_969_v30_)))
                d_970_v31_ = d_970_v31_
            elif source12_.is_DC6:
                d_971___mcc_h2_ = source12_.cf7
                d_972_cf7_ = d_971___mcc_h2_
                d_973_v32_: D14
                d_973_v32_ = D14_DC33(D14_DC31(d_972_cf7_))
                d_974_v33_: D14
                d_974_v33_ = D14_DC33(d_973_v32_)
                d_975_v34_: D14
                d_975_v34_ = D14_DC33(d_974_v33_)
                pat_let_tv25_ = d_974_v33_
                d_976_v35_: _dafny.Seq
                def iife70_(_pat_let15_0):
                    def iife71_(d_977_dt__update__tmp_h0_):
                        def iife72_(_pat_let16_0):
                            def iife73_(d_978_dt__update_hcf52_h0_):
                                return D14_DC33(d_978_dt__update_hcf52_h0_)
                            return iife73_(_pat_let16_0)
                        return iife72_(pat_let_tv25_)
                    return iife71_(_pat_let15_0)
                d_976_v35_ = _dafny.SeqWithoutIsStrInference([d_975_v34_, d_975_v34_, iife70_(d_975_v34_), d_975_v34_, d_975_v34_])
                d_979_v36_: _dafny.Seq
                d_979_v36_ = _dafny.SeqWithoutIsStrInference([d_976_v35_, d_976_v35_, d_976_v35_, d_976_v35_, d_976_v35_])
                d_980_v37_: _dafny.Map
                d_980_v37_ = _dafny.Map({(self).f5: d_979_v36_})
                d_980_v37_ = (d_980_v37_).set((self).f5, d_979_v36_)
                d_981_v38_: _dafny.Map
                d_981_v38_ = _dafny.Map({d_932_v0_: d_972_cf7_})
                d_982_v39_: _dafny.Seq
                d_982_v39_ = _dafny.SeqWithoutIsStrInference([len(d_981_v38_), (0) - (d_937_v4_)])
                d_983_v40_: _dafny.Seq
                d_983_v40_ = _dafny.SeqWithoutIsStrInference([(self).fm4(d_982_v39_, globalState)])
                (globalState).f2 = d_983_v40_
                (globalState).f2 = (d_982_v39_) + (_dafny.SeqWithoutIsStrInference([d_937_v4_ for d_984_i2_ in range(default__.abs(286))]))
                index212_ = default__.safeIndex(455, (d_933_v1_).length(0))
                (d_933_v1_)[index212_] = not((d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))])
            elif source12_.is_DC3:
                d_985___mcc_h3_ = source12_.cf4
                d_986_cf4_ = d_985___mcc_h3_
                d_932_v0_ = False
                d_932_v0_ = (d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]
                d_932_v0_ = (d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]
                d_987_v41_: D14
                d_987_v41_ = D14_DC31(d_937_v4_)
                r0 = (d_987_v41_).cf48
            elif True:
                d_988___mcc_h4_ = source12_.cf8
                d_989_cf8_ = d_988___mcc_h4_
                index213_ = default__.safeIndex(455, (d_933_v1_).length(0))
                (d_933_v1_)[index213_] = ((d_937_v4_) <= (d_937_v4_)) == (d_932_v0_)
                d_932_v0_ = default__.fm1(d_937_v4_, d_937_v4_, d_932_v0_, d_932_v0_, globalState)
                d_937_v4_ = d_937_v4_
                r0 = d_937_v4_
            d_990_v42_: _dafny.Map
            d_990_v42_ = _dafny.Map({d_932_v0_: (d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]})
            index214_ = default__.safeIndex(455, (d_933_v1_).length(0))
            (d_933_v1_)[index214_] = ((d_990_v42_)[(d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]] if ((d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]) in (d_990_v42_) else (self).fm3(d_937_v4_, d_937_v4_, globalState))
            d_991_v43_: _dafny.MultiSet
            d_991_v43_ = _dafny.MultiSet([d_937_v4_, d_937_v4_])
            d_992_v44_: _dafny.Map
            d_992_v44_ = _dafny.Map({(d_991_v43_ if d_932_v0_ else d_991_v43_): len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faqhwwpjf"))) + (d_936_v3_))})
            d_992_v44_ = _dafny.Map({d_991_v43_: d_937_v4_})
            d_993_v45_: _dafny.Map
            d_993_v45_ = _dafny.Map({_dafny.CodePoint('n'): (d_933_v1_)[default__.safeIndex(455, (d_933_v1_).length(0))]})
            d_993_v45_ = (d_993_v45_).set((self).f5, d_932_v0_)
        d_994_v46_: _dafny.Array
        nw213_ = _dafny.Array(D1.default()(), 28)
        d_994_v46_ = nw213_
        d_995_v47_: _dafny.MultiSet
        d_995_v47_ = _dafny.MultiSet([d_932_v0_, d_932_v0_])
        d_996_v48_: D1
        d_996_v48_ = D1_DC1(d_995_v47_)
        pat_let_tv26_ = d_995_v47_
        index215_ = default__.safeIndex(565, (d_994_v46_).length(0))
        def iife74_(_pat_let17_0):
            def iife75_(d_997_dt__update__tmp_h1_):
                def iife76_(_pat_let18_0):
                    def iife77_(d_998_dt__update_hcf1_h0_):
                        return D1_DC1(d_998_dt__update_hcf1_h0_)
                    return iife77_(_pat_let18_0)
                return iife76_(pat_let_tv26_)
            return iife75_(_pat_let17_0)
        (d_994_v46_)[index215_] = iife74_(d_996_v48_)
        pat_let_tv27_ = d_995_v47_
        pat_let_tv28_ = d_932_v0_
        pat_let_tv29_ = d_995_v47_
        index216_ = default__.safeIndex(565, (d_994_v46_).length(0))
        def iife79_(_pat_let20_0):
            def iife80_(d_999_dt__update__tmp_h2_):
                def iife81_(_pat_let21_0):
                    def iife82_(d_1000_dt__update_hcf1_h1_):
                        return D1_DC1(d_1000_dt__update_hcf1_h1_)
                    return iife82_(_pat_let21_0)
                return iife81_(pat_let_tv27_)
            return iife80_(_pat_let20_0)
        def iife78_(_pat_let19_0):
            def iife83_(d_1001_dt__update__tmp_h3_):
                def iife84_(_pat_let22_0):
                    def iife85_(d_1002_dt__update_hcf1_h2_):
                        return D1_DC1(d_1002_dt__update_hcf1_h2_)
                    return iife85_(_pat_let22_0)
                return iife84_((_dafny.MultiSet([pat_let_tv28_])) - (pat_let_tv29_))
            return iife83_(_pat_let19_0)
        (d_994_v46_)[index216_] = iife78_(iife79_(d_996_v48_))
        r0 = (366) - (d_937_v4_)
        d_1003_v49_: C0
        nw214_ = C0()
        nw214_.ctor__()
        d_1003_v49_ = nw214_
        d_1004_v50_: _dafny.Set
        d_1004_v50_ = _dafny.Set({(self).f5, (self).f5})
        d_1005_v51_: _dafny.MultiSet
        d_1005_v51_ = _dafny.MultiSet([len(d_1004_v50_)])
        d_1006_v52_: _dafny.Seq
        d_1006_v52_ = _dafny.SeqWithoutIsStrInference([d_937_v4_])
        d_1007_v53_: _dafny.Map
        d_1007_v53_ = _dafny.Map({d_932_v0_: (self).fm4(d_1006_v52_, globalState)})
        r0 = len((d_936_v3_ if default__.fm1((d_1005_v51_).cardinality, len(d_1007_v53_), d_932_v0_, False, globalState) else d_936_v3_))
        d_1008_v54_: _dafny.Array
        nw215_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_1008_v54_ = nw215_
        r1 = d_1008_v54_
        d_1009_v55_: _dafny.Map
        d_1009_v55_ = _dafny.Map({d_932_v0_: d_932_v0_})
        d_1010_v56_: _dafny.Map
        d_1010_v56_ = _dafny.Map({len(d_1009_v55_): (_dafny.SeqWithoutIsStrInference([not(True)])).set(default__.safeIndex(d_937_v4_, len(_dafny.SeqWithoutIsStrInference([not(True)]))), d_932_v0_)})
        r2 = (d_1010_v56_) | (d_1010_v56_)
        return r0, r1, r2


class C3(T1, T0):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f4, f5):
        (self)._f4 = f4
        (self)._f5 = f5

    def fm4(self, p0, globalState):
        return ((0) - ((D1_DC2(325, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asvmv"))))).cf3)) - (default__.safeModuloInt(878, 778))

    def fm5(self, p0, p1, p2, p3, globalState):
        return ((not(True)) in (_dafny.Map({False: _dafny.Map({-780: True})}))) and ((_dafny.MultiSet([999, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bimvtv"))), (0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "js")))}))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tg")))])).ispropersubset((D8_DC18(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))]))).cf28))

    def fm3(self, p0, p1, globalState):
        return not(((self).f5) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcrfnbjfd"))))

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        d_1011_v0_: bool
        d_1011_v0_ = True
        d_1012_v1_: _dafny.Array
        def lambda65_(d_1013_v0_):
            def lambda66_(d_1014_i0_):
                return d_1013_v0_

            return lambda66_

        init36_ = lambda65_(d_1011_v0_)
        nw216_ = _dafny.Array(None, 3)
        for i0_36_ in range(nw216_.length(0)):
            nw216_[i0_36_] = init36_(i0_36_)
        d_1012_v1_ = nw216_
        rhs196_ = not(d_1011_v0_)
        rhs197_ = d_1012_v1_
        lhs104_ = globalState
        d_1011_v0_ = rhs196_
        lhs104_.f3 = rhs197_
        index217_ = default__.safeIndex(117, (d_1012_v1_).length(0))
        (d_1012_v1_)[index217_] = d_1011_v0_
        d_1015_v2_: _dafny.Seq
        d_1015_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjluy"))
        index218_ = default__.safeIndex(117, (d_1012_v1_).length(0))
        (d_1012_v1_)[index218_] = ((self).f5) in (d_1015_v2_)
        d_1016_v3_: int
        d_1016_v3_ = 733
        d_1017_v4_: _dafny.Map
        d_1017_v4_ = _dafny.Map({d_1016_v3_: (self).f4})
        d_1018_v5_: C2
        nw217_ = C2()
        nw217_.ctor__(((d_1017_v4_)[d_1016_v3_] if (d_1016_v3_) in (d_1017_v4_) else (self).f4), (self).f5)
        d_1018_v5_ = nw217_
        d_1019_v6_: _dafny.Array
        nw218_ = _dafny.Array(None, 7)
        nw218_[int(0)] = (self).f5
        nw218_[int(1)] = (self).f5
        nw218_[int(2)] = (self).f5
        nw218_[int(3)] = _dafny.CodePoint('a')
        nw218_[int(4)] = _dafny.CodePoint('j')
        nw218_[int(5)] = (self).f5
        nw218_[int(6)] = _dafny.CodePoint('b')
        d_1019_v6_ = nw218_
        index219_ = default__.safeIndex(956, (d_1019_v6_).length(0))
        (d_1019_v6_)[index219_] = (self).f5
        index220_ = default__.safeIndex(956, (d_1019_v6_).length(0))
        (d_1019_v6_)[index220_] = (self).f5
        (globalState).f3 = d_1012_v1_
        d_1020_v7_: C1
        nw219_ = C1()
        nw219_.ctor__((self).f4, (self).f5)
        d_1020_v7_ = nw219_
        d_1021_v8_: _dafny.Seq
        d_1021_v8_ = _dafny.SeqWithoutIsStrInference([607])
        d_1022_v9_: _dafny.Map
        d_1022_v9_ = _dafny.Map({(d_1012_v1_)[default__.safeIndex(117, (d_1012_v1_).length(0))]: (d_1012_v1_)[default__.safeIndex(117, (d_1012_v1_).length(0))]})
        r0 = (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-712: _dafny.Set({True})})) for d_1023_i1_ in range(default__.abs(79))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1016_v3_, d_1016_v3_, len(d_1021_v8_), d_1016_v3_, len(d_1022_v9_)]))]))
        r1 = d_1016_v3_
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1024_v0_: bool
        d_1024_v0_ = False
        d_1025_i0_: int
        d_1025_i0_ = 0
        with _dafny.label("4"):
            while d_1024_v0_:
                with _dafny.c_label("4"):
                    if (d_1025_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_1025_i0_ = (d_1025_i0_) + (1)
                    d_1026_v1_: int
                    d_1026_v1_ = -202
                    d_1027_v2_: _dafny.Seq
                    d_1027_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adlfp"))
                    d_1028_v3_: _dafny.Seq
                    d_1028_v3_ = _dafny.SeqWithoutIsStrInference([d_1026_v1_, d_1026_v1_])
                    d_1029_v4_: _dafny.Seq
                    d_1029_v4_ = _dafny.SeqWithoutIsStrInference([d_1026_v1_, (0) - (len(d_1027_v2_)), (0) - ((d_1028_v3_)[default__.safeIndex(559, len(d_1028_v3_))])])
                    d_1026_v1_ = (self).fm4(d_1029_v4_, globalState)
                    d_1030_v5_: _dafny.Seq
                    d_1030_v5_ = _dafny.SeqWithoutIsStrInference([d_1024_v0_])
                    d_1031_v6_: _dafny.Set
                    d_1031_v6_ = _dafny.Set({(self).fm4(d_1029_v4_, globalState), len(d_1027_v2_), d_1026_v1_, len(d_1030_v5_)})
                    d_1032_v7_: _dafny.Seq
                    d_1032_v7_ = _dafny.SeqWithoutIsStrInference([d_1031_v6_, (d_1031_v6_).intersection(d_1031_v6_)])
                    d_1032_v7_ = d_1032_v7_
                    d_1033_v8_: _dafny.Array
                    def lambda67_(d_1034_v1_):
                        def lambda68_(d_1035_i1_):
                            return default__.safeDivisionInt(d_1035_i1_, d_1034_v1_)

                        return lambda68_

                    init37_ = lambda67_(d_1026_v1_)
                    nw220_ = _dafny.Array(None, 22)
                    for i0_37_ in range(nw220_.length(0)):
                        nw220_[i0_37_] = init37_(i0_37_)
                    d_1033_v8_ = nw220_
                    d_1036_v9_: _dafny.Map
                    d_1036_v9_ = _dafny.Map({(d_1026_v1_) != ((0) - (d_1026_v1_)): d_1033_v8_})
                    d_1036_v9_ = d_1036_v9_
                    d_1026_v1_ = d_1026_v1_
                    pass
            pass
        d_1037_v10_: int
        d_1037_v10_ = 929
        d_1038_v11_: _dafny.Map
        d_1038_v11_ = _dafny.Map({d_1037_v10_: d_1037_v10_})
        d_1039_v12_: _dafny.Seq
        d_1039_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjdpd"))
        d_1040_v13_: _dafny.Seq
        d_1040_v13_ = _dafny.SeqWithoutIsStrInference([d_1039_v12_])
        if (len(d_1038_v11_)) >= ((0) - (len(d_1040_v13_))):
            d_1041_v14_: C1
            nw221_ = C1()
            nw221_.ctor__((self).f4, default__.fm0((0) - (d_1037_v10_), 957, globalState))
            d_1041_v14_ = nw221_
            d_1042_v15_: _dafny.Array
            def lambda69_(d_1043_v0_):
                def lambda70_(d_1044_i2_):
                    return d_1043_v0_

                return lambda70_

            init38_ = lambda69_(d_1024_v0_)
            nw222_ = _dafny.Array(None, 13)
            for i0_38_ in range(nw222_.length(0)):
                nw222_[i0_38_] = init38_(i0_38_)
            d_1042_v15_ = nw222_
            index221_ = default__.safeIndex(35, (d_1042_v15_).length(0))
            (d_1042_v15_)[index221_] = True
            index222_ = default__.safeIndex(35, (d_1042_v15_).length(0))
            (d_1042_v15_)[index222_] = (d_1024_v0_) and (d_1024_v0_)
            d_1037_v10_ = d_1037_v10_
            d_1045_v16_: _dafny.Seq
            d_1045_v16_ = _dafny.SeqWithoutIsStrInference([d_1037_v10_, d_1037_v10_])
            d_1046_v17_: _dafny.MultiSet
            d_1046_v17_ = _dafny.MultiSet([_dafny.Map({d_1037_v10_: len(d_1045_v16_)})])
            d_1047_v18_: _dafny.Set
            d_1047_v18_ = _dafny.Set({d_1024_v0_})
            d_1048_v19_: _dafny.MultiSet
            d_1048_v19_ = _dafny.MultiSet([(d_1042_v15_)[default__.safeIndex(35, (d_1042_v15_).length(0))], (d_1042_v15_)[default__.safeIndex(35, (d_1042_v15_).length(0))]])
            d_1049_v20_: _dafny.Array
            nw223_ = _dafny.Array(None, 21)
            nw223_[int(0)] = d_1037_v10_
            nw223_[int(1)] = (d_1037_v10_) - (d_1037_v10_)
            nw223_[int(2)] = (d_1037_v10_) * (((d_1046_v17_)[d_1038_v11_] if (d_1038_v11_) in (d_1046_v17_) else d_1037_v10_))
            nw223_[int(3)] = d_1037_v10_
            nw223_[int(4)] = d_1037_v10_
            nw223_[int(5)] = default__.safeModuloInt(d_1037_v10_, d_1037_v10_)
            nw223_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypjujowm")))
            nw223_[int(7)] = d_1037_v10_
            nw223_[int(8)] = default__.safeDivisionInt(d_1037_v10_, len(d_1047_v18_))
            nw223_[int(9)] = d_1037_v10_
            nw223_[int(10)] = -406
            nw223_[int(11)] = ((d_1048_v19_)[d_1024_v0_] if (d_1024_v0_) in (d_1048_v19_) else d_1037_v10_)
            nw223_[int(12)] = d_1037_v10_
            nw223_[int(13)] = d_1037_v10_
            nw223_[int(14)] = (0) - (d_1037_v10_)
            nw223_[int(15)] = len(d_1039_v12_)
            nw223_[int(16)] = d_1037_v10_
            nw223_[int(17)] = d_1037_v10_
            nw223_[int(18)] = d_1037_v10_
            nw223_[int(19)] = (336) - (d_1037_v10_)
            nw223_[int(20)] = d_1037_v10_
            d_1049_v20_ = nw223_
            index223_ = default__.safeIndex(643, (d_1049_v20_).length(0))
            (d_1049_v20_)[index223_] = 468
            index224_ = default__.safeIndex(643, (d_1049_v20_).length(0))
            (d_1049_v20_)[index224_] = default__.safeModuloInt(d_1037_v10_, d_1037_v10_)
            index225_ = default__.safeIndex(35, (d_1042_v15_).length(0))
            rhs198_ = (self).fm5(d_1037_v10_, (d_1037_v10_) <= (((d_1038_v11_)[(d_1049_v20_)[default__.safeIndex(643, (d_1049_v20_).length(0))]] if ((d_1049_v20_)[default__.safeIndex(643, (d_1049_v20_).length(0))]) in (d_1038_v11_) else len(d_1039_v12_))), len(_dafny.Map({(self).f5: (d_1042_v15_)[default__.safeIndex(35, (d_1042_v15_).length(0))]})), (d_1049_v20_)[default__.safeIndex(643, (d_1049_v20_).length(0))], globalState)
            rhs199_ = d_1049_v20_
            rhs200_ = d_1024_v0_
            lhs105_ = d_1042_v15_
            lhs106_ = default__.safeIndex(35, (d_1042_v15_).length(0))
            d_1024_v0_ = rhs198_
            d_1049_v20_ = rhs199_
            lhs105_[lhs106_] = rhs200_
        elif True:
            d_1039_v12_ = d_1039_v12_
            d_1024_v0_ = True
            d_1050_v21_: _dafny.Map
            d_1050_v21_ = _dafny.Map({d_1024_v0_: _dafny.SeqWithoutIsStrInference([(self).f5 for d_1051_i3_ in range(default__.abs(-749))])})
            d_1052_v22_: _dafny.Set
            d_1052_v22_ = _dafny.Set({d_1024_v0_})
            d_1024_v0_ = default__.fm1((0) - (len(((d_1039_v12_).set(default__.safeIndex(d_1037_v10_, len(d_1039_v12_)), (self).f5)) + (d_1039_v12_))), ((0) - ((0) - ((0) - (len(((d_1050_v21_)[d_1024_v0_] if (d_1024_v0_) in (d_1050_v21_) else d_1039_v12_))))) if True else d_1037_v10_), (_dafny.Set({d_1024_v0_, d_1024_v0_, d_1024_v0_})).ispropersubset(d_1052_v22_), False, globalState)
            d_1053_v23_: D14
            d_1053_v23_ = D14_DC32(d_1037_v10_, d_1024_v0_, d_1024_v0_)
            d_1054_v24_: _dafny.Seq
            d_1054_v24_ = _dafny.SeqWithoutIsStrInference([(d_1053_v23_).cf49])
            d_1055_v25_: _dafny.Array
            nw224_ = _dafny.Array(False, 23)
            d_1055_v25_ = nw224_
            d_1056_v26_: _dafny.Seq
            d_1056_v26_ = _dafny.SeqWithoutIsStrInference([d_1055_v25_])
            d_1057_v27_: _dafny.Array
            nw225_ = _dafny.Array(None, 3)
            nw225_[int(0)] = len(d_1054_v24_)
            nw225_[int(1)] = len(_dafny.Map({d_1037_v10_: len(d_1056_v26_)}))
            nw225_[int(2)] = (0) - (d_1037_v10_)
            d_1057_v27_ = nw225_
            index226_ = default__.safeIndex(979, (d_1057_v27_).length(0))
            (d_1057_v27_)[index226_] = d_1037_v10_
            index227_ = default__.safeIndex(979, (d_1057_v27_).length(0))
            (d_1057_v27_)[index227_] = d_1037_v10_
            d_1037_v10_ = (0) - (len(d_1056_v26_))
        d_1058_v29_: _dafny.Seq
        def iife86_():
            coll40_ = _dafny.Set()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(678, -935):
                d_1059_v28_: int = compr_40_
                if ((678) <= (d_1059_v28_)) and ((d_1059_v28_) < (-935)):
                    coll40_ = coll40_.union(_dafny.Set([default__.safeDivisionInt(d_1059_v28_, d_1037_v10_)]))
            return _dafny.Set(coll40_)
        d_1058_v29_ = _dafny.SeqWithoutIsStrInference([d_1037_v10_, len(iife86_()
        )])
        d_1037_v10_ = (d_1037_v10_ if d_1024_v0_ else len((d_1058_v29_) + (_dafny.SeqWithoutIsStrInference([d_1037_v10_ for d_1060_i4_ in range(default__.abs(-255))]))))
        d_1037_v10_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjknrytek"))) + ((d_1039_v12_) + (d_1039_v12_)))
        d_1061_v30_: _dafny.MultiSet
        d_1061_v30_ = _dafny.MultiSet([d_1037_v10_])
        d_1037_v10_ = (d_1061_v30_).cardinality
        d_1037_v10_ = ((d_1038_v11_)[d_1037_v10_] if (d_1037_v10_) in (d_1038_v11_) else d_1037_v10_)
        d_1062_v31_: _dafny.Array
        nw226_ = _dafny.Array(int(0), 6)
        d_1062_v31_ = nw226_
        r0 = d_1062_v31_
        return r0

    def m11(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1063_v0_: _dafny.Array
        def lambda71_(d_1064_i0_):
            return _dafny.SeqWithoutIsStrInference([not(False), False, not(False)])

        init39_ = lambda71_
        nw227_ = _dafny.Array(None, 23)
        for i0_39_ in range(nw227_.length(0)):
            nw227_[i0_39_] = init39_(i0_39_)
        d_1063_v0_ = nw227_
        d_1065_v1_: bool
        d_1065_v1_ = False
        d_1066_v2_: _dafny.Seq
        d_1066_v2_ = _dafny.SeqWithoutIsStrInference([d_1065_v1_, d_1065_v1_, d_1065_v1_, False])
        index228_ = default__.safeIndex(25, (d_1063_v0_).length(0))
        (d_1063_v0_)[index228_] = d_1066_v2_
        d_1067_v3_: int
        d_1067_v3_ = 794
        d_1068_v4_: _dafny.Set
        d_1068_v4_ = _dafny.Set({d_1065_v1_})
        index229_ = default__.safeIndex(25, (d_1063_v0_).length(0))
        (d_1063_v0_)[index229_] = ((d_1066_v2_) + (d_1066_v2_)).set(default__.safeIndex(d_1067_v3_, len((d_1066_v2_) + (d_1066_v2_))), (d_1068_v4_).isdisjoint(d_1068_v4_))
        d_1065_v1_ = d_1065_v1_
        d_1069_i1_: int
        d_1069_i1_ = 0
        with _dafny.label("5"):
            while False:
                with _dafny.c_label("5"):
                    if (d_1069_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_1069_i1_ = (d_1069_i1_) + (1)
                    d_1070_v5_: _dafny.Map
                    d_1070_v5_ = _dafny.Map({(self).f5: (self).f5})
                    d_1071_v6_: _dafny.Map
                    d_1071_v6_ = _dafny.Map({len(d_1070_v5_): d_1065_v1_})
                    if (d_1065_v1_) or (not (True) or (((d_1071_v6_)[d_1067_v3_] if (d_1067_v3_) in (d_1071_v6_) else d_1065_v1_))):
                        d_1072_v7_: _dafny.Seq
                        d_1072_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "al"))
                        d_1073_v8_: str
                        d_1073_v8_ = _dafny.CodePoint('d')
                        d_1074_v9_: _dafny.Seq
                        d_1074_v9_ = _dafny.SeqWithoutIsStrInference([d_1067_v3_, d_1067_v3_])
                        rhs201_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbrvy"))).set(default__.safeIndex((d_1074_v9_)[default__.safeIndex(d_1067_v3_, len(d_1074_v9_))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbrvy")))), d_1073_v8_)
                        rhs202_ = not((d_1065_v1_) or (d_1065_v1_))
                        rhs203_ = d_1073_v8_
                        d_1072_v7_ = rhs201_
                        d_1065_v1_ = rhs202_
                        d_1073_v8_ = rhs203_
                        d_1065_v1_ = d_1065_v1_
                        d_1065_v1_ = d_1065_v1_
                        d_1075_v10_: _dafny.Set
                        d_1075_v10_ = _dafny.Set({d_1067_v3_})
                        d_1076_v11_: _dafny.Map
                        d_1076_v11_ = _dafny.Map({((_dafny.SeqWithoutIsStrInference([not(d_1065_v1_), d_1065_v1_, d_1065_v1_])) + (((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))]).set(default__.safeIndex(len(d_1075_v10_), len((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))])), d_1065_v1_))).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1065_v1_: (0) - (len(d_1074_v9_))}) for d_1077_i2_ in range(default__.abs(174))]))), len((_dafny.SeqWithoutIsStrInference([not(d_1065_v1_), d_1065_v1_, d_1065_v1_])) + (((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))]).set(default__.safeIndex(len(d_1075_v10_), len((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))])), d_1065_v1_)))), d_1065_v1_): (d_1065_v1_ if d_1065_v1_ else not(d_1065_v1_))})
                        d_1078_v12_: D6
                        d_1078_v12_ = D6_DC15(d_1065_v1_, d_1065_v1_, d_1065_v1_)
                        d_1065_v1_ = not(not(((d_1076_v11_)[d_1066_v2_] if (d_1066_v2_) in (d_1076_v11_) else (d_1078_v12_).cf21)))
                        d_1067_v3_ = (0) - (d_1067_v3_)
                    elif True:
                        d_1079_v13_: C1
                        nw228_ = C1()
                        nw228_.ctor__((self).f4, (self).f5)
                        d_1079_v13_ = nw228_
                        d_1080_v14_: _dafny.Array
                        nw229_ = _dafny.Array(int(0), 1)
                        d_1080_v14_ = nw229_
                        d_1081_v15_: D11
                        d_1081_v15_ = D11_DC25(d_1065_v1_, 157, d_1080_v14_, (0) - (d_1067_v3_))
                        d_1082_v16_: _dafny.Set
                        d_1082_v16_ = _dafny.Set({89, 918, d_1067_v3_})
                        d_1083_v17_: _dafny.Seq
                        d_1083_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                        d_1084_v18_: _dafny.MultiSet
                        d_1084_v18_ = _dafny.MultiSet([(self).fm5(d_1067_v3_, d_1065_v1_, -836, d_1067_v3_, globalState)])
                        d_1085_v19_: _dafny.MultiSet
                        d_1085_v19_ = _dafny.MultiSet([_dafny.CodePoint('m'), (self).f5, (self).f5, (self).f5])
                        d_1086_v20_: _dafny.Seq
                        d_1086_v20_ = _dafny.SeqWithoutIsStrInference([len(d_1083_v17_), (d_1085_v19_).cardinality])
                        d_1087_v21_: _dafny.Map
                        d_1087_v21_ = _dafny.Map({d_1065_v1_: d_1065_v1_})
                        d_1088_v22_: _dafny.Array
                        nw230_ = _dafny.Array(None, 26)
                        nw230_[int(0)] = d_1065_v1_
                        nw230_[int(1)] = True
                        nw230_[int(2)] = False
                        nw230_[int(3)] = False
                        nw230_[int(4)] = default__.fm1((0) - (d_1067_v3_), d_1067_v3_, d_1065_v1_, d_1065_v1_, globalState)
                        nw230_[int(5)] = (d_1081_v15_).cf38
                        nw230_[int(6)] = ((0) - (d_1067_v3_)) >= (d_1067_v3_)
                        nw230_[int(7)] = not(d_1065_v1_)
                        nw230_[int(8)] = (_dafny.MultiSet(d_1066_v2_)).ispropersubset(_dafny.MultiSet([d_1065_v1_]))
                        nw230_[int(9)] = d_1065_v1_
                        nw230_[int(10)] = d_1065_v1_
                        nw230_[int(11)] = d_1065_v1_
                        nw230_[int(12)] = d_1065_v1_
                        nw230_[int(13)] = d_1065_v1_
                        nw230_[int(14)] = (self).fm5(len(d_1082_v16_), d_1065_v1_, len(d_1083_v17_), d_1067_v3_, globalState)
                        nw230_[int(15)] = d_1065_v1_
                        nw230_[int(16)] = (d_1068_v4_) == (d_1068_v4_)
                        nw230_[int(17)] = d_1065_v1_
                        nw230_[int(18)] = d_1065_v1_
                        nw230_[int(19)] = d_1065_v1_
                        nw230_[int(20)] = not(d_1065_v1_)
                        nw230_[int(21)] = (d_1067_v3_) <= (len(_dafny.Map({d_1065_v1_: d_1084_v18_})))
                        nw230_[int(22)] = (d_1079_v13_).fm5(d_1067_v3_, d_1065_v1_, (d_1086_v20_)[default__.safeIndex(len((d_1087_v21_).set(d_1065_v1_, d_1065_v1_)), len(d_1086_v20_))], d_1067_v3_, globalState)
                        nw230_[int(23)] = d_1065_v1_
                        nw230_[int(24)] = d_1065_v1_
                        nw230_[int(25)] = False
                        d_1088_v22_ = nw230_
                        index230_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        (d_1088_v22_)[index230_] = d_1065_v1_
                        index231_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        (d_1088_v22_)[index231_] = d_1065_v1_
                        d_1089_v23_: D1
                        d_1089_v23_ = D1_DC2(d_1067_v3_, len(_dafny.SeqWithoutIsStrInference([True, (d_1088_v22_)[default__.safeIndex(956, (d_1088_v22_).length(0))]])))
                        d_1090_v24_: _dafny.Map
                        d_1090_v24_ = _dafny.Map({True: d_1067_v3_})
                        d_1091_v25_: _dafny.Map
                        d_1091_v25_ = _dafny.Map({d_1089_v23_: ((d_1090_v24_)[True] if (True) in (d_1090_v24_) else d_1067_v3_)})
                        d_1092_v26_: _dafny.MultiSet
                        d_1092_v26_ = _dafny.MultiSet([len(d_1086_v20_)])
                        d_1093_v27_: _dafny.Seq
                        d_1093_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yofklgq"))])
                        index232_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        index233_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        rhs204_ = (807) + (((0) - (d_1067_v3_)) + (-678))
                        rhs205_ = ((d_1079_v13_).fm4(_dafny.SeqWithoutIsStrInference([(d_1092_v26_).cardinality]), globalState)) * ((d_1067_v3_) + (d_1067_v3_))
                        rhs206_ = ((d_1091_v25_) | (d_1091_v25_)).set(d_1089_v23_, d_1067_v3_)
                        rhs207_ = default__.fm1(d_1067_v3_, len(default__.fm29(len(d_1093_v27_), d_1065_v1_, d_1067_v3_, globalState)), (d_1088_v22_)[default__.safeIndex(956, (d_1088_v22_).length(0))], (d_1088_v22_)[default__.safeIndex(956, (d_1088_v22_).length(0))], globalState)
                        rhs208_ = d_1065_v1_
                        lhs107_ = d_1088_v22_
                        lhs108_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        lhs109_ = d_1088_v22_
                        lhs110_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        d_1067_v3_ = rhs204_
                        d_1067_v3_ = rhs205_
                        d_1091_v25_ = rhs206_
                        lhs107_[lhs108_] = rhs207_
                        lhs109_[lhs110_] = rhs208_
                        d_1094_v28_: _dafny.Map
                        d_1094_v28_ = _dafny.Map({d_1065_v1_: (self).f5})
                        d_1095_v29_: D5
                        d_1095_v29_ = D5_DC11(d_1094_v28_)
                        d_1096_v30_: D2
                        d_1096_v30_ = D2_DC3(d_1088_v22_)
                        d_1097_v31_: _dafny.Array
                        nw231_ = _dafny.Array(None, 4)
                        nw231_[int(0)] = d_1088_v22_
                        nw231_[int(1)] = d_1088_v22_
                        nw231_[int(2)] = (d_1096_v30_).cf4
                        nw231_[int(3)] = d_1088_v22_
                        d_1097_v31_ = nw231_
                        d_1098_v32_: _dafny.Map
                        d_1098_v32_ = _dafny.Map({d_1097_v31_: d_1065_v1_})
                        rhs209_ = ((d_1098_v32_)[d_1097_v31_] if (d_1097_v31_) in (d_1098_v32_) else (not((d_1088_v22_)[default__.safeIndex(956, (d_1088_v22_).length(0))])) == (d_1065_v1_))
                        rhs210_ = d_1067_v3_
                        rhs211_ = default__.safeModuloInt((0) - (len((d_1086_v20_) + (d_1086_v20_))), d_1067_v3_)
                        rhs212_ = d_1080_v14_
                        rhs213_ = d_1095_v29_
                        d_1065_v1_ = rhs209_
                        d_1067_v3_ = rhs210_
                        d_1067_v3_ = rhs211_
                        d_1080_v14_ = rhs212_
                        d_1095_v29_ = rhs213_
                        index234_ = default__.safeIndex(956, (d_1088_v22_).length(0))
                        (d_1088_v22_)[index234_] = (d_1088_v22_)[default__.safeIndex(956, (d_1088_v22_).length(0))]
                    d_1099_v33_: _dafny.Seq
                    d_1099_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvdnequn"))
                    d_1100_v34_: D2
                    d_1100_v34_ = D2_DC6(len(d_1099_v33_))
                    d_1101_v35_: D2
                    d_1101_v35_ = D2_DC7(d_1100_v34_)
                    d_1102_v36_: D2
                    d_1102_v36_ = D2_DC7(d_1101_v35_)
                    d_1103_v37_: _dafny.Map
                    d_1103_v37_ = _dafny.Map({d_1065_v1_: d_1102_v36_})
                    d_1104_v38_: _dafny.Set
                    d_1104_v38_ = _dafny.Set({d_1103_v37_})
                    d_1105_v39_: D7
                    d_1105_v39_ = D7_DC17(d_1065_v1_, d_1065_v1_, d_1065_v1_, d_1065_v1_)
                    d_1106_v40_: _dafny.Map
                    d_1106_v40_ = _dafny.Map({-176: d_1104_v38_})
                    d_1107_v41_: _dafny.Seq
                    d_1107_v41_ = _dafny.SeqWithoutIsStrInference([d_1104_v38_, d_1104_v38_, d_1104_v38_, d_1104_v38_])
                    d_1108_v42_: _dafny.Array
                    nw232_ = _dafny.Array(None, 27)
                    nw232_[int(0)] = d_1104_v38_
                    nw232_[int(1)] = (d_1104_v38_) - (_dafny.Set({d_1103_v37_, _dafny.Map({(d_1105_v39_).cf24: d_1102_v36_}), d_1103_v37_}))
                    nw232_[int(2)] = _dafny.Set({d_1103_v37_})
                    nw232_[int(3)] = d_1104_v38_
                    nw232_[int(4)] = (d_1104_v38_).intersection(d_1104_v38_)
                    nw232_[int(5)] = (_dafny.Set({d_1103_v37_})) | (d_1104_v38_)
                    nw232_[int(6)] = ((d_1106_v40_)[d_1067_v3_] if (d_1067_v3_) in (d_1106_v40_) else d_1104_v38_)
                    nw232_[int(7)] = d_1104_v38_
                    nw232_[int(8)] = d_1104_v38_
                    nw232_[int(9)] = (d_1104_v38_) | (d_1104_v38_)
                    nw232_[int(10)] = d_1104_v38_
                    nw232_[int(11)] = d_1104_v38_
                    nw232_[int(12)] = d_1104_v38_
                    nw232_[int(13)] = _dafny.Set({d_1103_v37_})
                    nw232_[int(14)] = (d_1107_v41_)[default__.safeIndex(len(d_1099_v33_), len(d_1107_v41_))]
                    nw232_[int(15)] = d_1104_v38_
                    nw232_[int(16)] = _dafny.Set({d_1103_v37_, _dafny.Map({True: d_1102_v36_}), _dafny.Map({d_1065_v1_: d_1102_v36_})})
                    nw232_[int(17)] = d_1104_v38_
                    nw232_[int(18)] = d_1104_v38_
                    nw232_[int(19)] = (d_1104_v38_) | (_dafny.Set({_dafny.Map({d_1065_v1_: d_1102_v36_})}))
                    nw232_[int(20)] = (d_1104_v38_).intersection(_dafny.Set({d_1103_v37_, d_1103_v37_}))
                    nw232_[int(21)] = (d_1104_v38_) | (d_1104_v38_)
                    nw232_[int(22)] = (d_1104_v38_) | (d_1104_v38_)
                    nw232_[int(23)] = d_1104_v38_
                    nw232_[int(24)] = d_1104_v38_
                    nw232_[int(25)] = d_1104_v38_
                    nw232_[int(26)] = d_1104_v38_
                    d_1108_v42_ = nw232_
                    index235_ = default__.safeIndex(297, (d_1108_v42_).length(0))
                    (d_1108_v42_)[index235_] = d_1104_v38_
                    d_1109_v43_: _dafny.Seq
                    d_1109_v43_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1067_v3_: d_1065_v1_})])
                    d_1110_v44_: _dafny.MultiSet
                    d_1110_v44_ = _dafny.MultiSet([d_1071_v6_])
                    d_1111_v45_: _dafny.Seq
                    d_1111_v45_ = _dafny.SeqWithoutIsStrInference([d_1067_v3_, d_1067_v3_])
                    d_1112_v46_: _dafny.Set
                    d_1112_v46_ = _dafny.Set({len(_dafny.Set({d_1111_v45_, d_1111_v45_, _dafny.SeqWithoutIsStrInference([d_1067_v3_ for d_1113_i3_ in range(default__.abs(-754))])}))})
                    index236_ = default__.safeIndex(297, (d_1108_v42_).length(0))
                    rhs214_ = d_1104_v38_
                    rhs215_ = d_1099_v33_
                    rhs216_ = (0) - ((((_dafny.MultiSet(d_1109_v43_)).intersection(d_1110_v44_)).cardinality if d_1065_v1_ else d_1067_v3_))
                    rhs217_ = len(d_1112_v46_)
                    rhs218_ = default__.safeDivisionInt(len((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))]), d_1067_v3_)
                    lhs111_ = d_1108_v42_
                    lhs112_ = default__.safeIndex(297, (d_1108_v42_).length(0))
                    lhs111_[lhs112_] = rhs214_
                    d_1099_v33_ = rhs215_
                    d_1067_v3_ = rhs216_
                    d_1067_v3_ = rhs217_
                    d_1067_v3_ = rhs218_
                    d_1114_v47_: _dafny.Array
                    nw233_ = _dafny.Array(False, 7)
                    d_1114_v47_ = nw233_
                    index237_ = default__.safeIndex(714, (d_1114_v47_).length(0))
                    (d_1114_v47_)[index237_] = (len(d_1111_v45_)) == (d_1067_v3_)
                    index238_ = default__.safeIndex(537, (d_1114_v47_).length(0))
                    (d_1114_v47_)[index238_] = d_1065_v1_
                    d_1115_v48_: _dafny.Array
                    def lambda72_(d_1116_i4_):
                        return default__.safeDivisionInt(d_1116_i4_, 793)

                    init40_ = lambda72_
                    nw234_ = _dafny.Array(None, 23)
                    for i0_40_ in range(nw234_.length(0)):
                        nw234_[i0_40_] = init40_(i0_40_)
                    d_1115_v48_ = nw234_
                    index239_ = default__.safeIndex(998, (d_1115_v48_).length(0))
                    (d_1115_v48_)[index239_] = d_1067_v3_
                    d_1117_v49_: _dafny.Map
                    d_1117_v49_ = _dafny.Map({d_1065_v1_: d_1067_v3_})
                    index240_ = default__.safeIndex(714, (d_1114_v47_).length(0))
                    index241_ = default__.safeIndex(537, (d_1114_v47_).length(0))
                    index242_ = default__.safeIndex(998, (d_1115_v48_).length(0))
                    rhs219_ = ((d_1065_v1_ if d_1065_v1_ else d_1065_v1_) if (d_1117_v49_) == (d_1117_v49_) else True)
                    rhs220_ = d_1065_v1_
                    rhs221_ = (d_1067_v3_ if d_1065_v1_ else (d_1067_v3_) * (d_1067_v3_))
                    lhs113_ = d_1114_v47_
                    lhs114_ = default__.safeIndex(714, (d_1114_v47_).length(0))
                    lhs115_ = d_1114_v47_
                    lhs116_ = default__.safeIndex(537, (d_1114_v47_).length(0))
                    lhs117_ = d_1115_v48_
                    lhs118_ = default__.safeIndex(998, (d_1115_v48_).length(0))
                    lhs113_[lhs114_] = rhs219_
                    lhs115_[lhs116_] = rhs220_
                    lhs117_[lhs118_] = rhs221_
                    d_1118_v50_: _dafny.MultiSet
                    d_1118_v50_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shaofskxb"))])
                    index243_ = default__.safeIndex(714, (d_1114_v47_).length(0))
                    rhs222_ = not((d_1118_v50_).ispropersubset((d_1118_v50_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmd")), default__.abs((d_1115_v48_)[default__.safeIndex(998, (d_1115_v48_).length(0))]))))
                    rhs223_ = default__.safeModuloInt((d_1115_v48_)[default__.safeIndex(998, (d_1115_v48_).length(0))], (d_1115_v48_)[default__.safeIndex(998, (d_1115_v48_).length(0))])
                    lhs119_ = d_1114_v47_
                    lhs120_ = default__.safeIndex(714, (d_1114_v47_).length(0))
                    lhs119_[lhs120_] = rhs222_
                    d_1067_v3_ = rhs223_
                    pass
            pass
        d_1119_v51_: _dafny.Map
        d_1119_v51_ = _dafny.Map({d_1067_v3_: d_1065_v1_})
        d_1120_v52_: _dafny.MultiSet
        d_1120_v52_ = _dafny.MultiSet([len(_dafny.Set({d_1119_v51_, d_1119_v51_})), d_1067_v3_, d_1067_v3_, d_1067_v3_])
        d_1121_v53_: _dafny.Array
        def lambda73_(d_1122_v3_):
            def lambda74_(d_1123_i5_):
                return default__.safeDivisionInt(d_1123_i5_, d_1122_v3_)

            return lambda74_

        init41_ = lambda73_(d_1067_v3_)
        nw235_ = _dafny.Array(None, 17)
        for i0_41_ in range(nw235_.length(0)):
            nw235_[i0_41_] = init41_(i0_41_)
        d_1121_v53_ = nw235_
        d_1124_v54_: _dafny.MultiSet
        d_1124_v54_ = _dafny.MultiSet([(self).f5])
        d_1125_v55_: _dafny.Seq
        d_1125_v55_ = _dafny.SeqWithoutIsStrInference([d_1067_v3_])
        rhs224_ = ((d_1120_v52_).intersection(d_1120_v52_) if not((d_1124_v54_).isdisjoint(d_1124_v54_)) else (d_1120_v52_) | (d_1120_v52_))
        rhs225_ = d_1121_v53_
        rhs226_ = (0) - (((0) - (default__.safeModuloInt(d_1067_v3_, d_1067_v3_))) * (default__.safeDivisionInt((d_1125_v55_)[default__.safeIndex(d_1067_v3_, len(d_1125_v55_))], d_1067_v3_)))
        d_1120_v52_ = rhs224_
        d_1121_v53_ = rhs225_
        d_1067_v3_ = rhs226_
        hi7_ = d_1067_v3_
        for d_1126_i6_ in range((len((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))])) + (d_1067_v3_), hi7_):
            d_1127_v56_: D14
            d_1127_v56_ = D14_DC31(d_1126_i6_)
            d_1128_v57_: _dafny.MultiSet
            d_1128_v57_ = _dafny.MultiSet([d_1127_v56_])
            d_1129_v58_: _dafny.Seq
            d_1129_v58_ = _dafny.SeqWithoutIsStrInference([d_1127_v56_, d_1127_v56_, d_1127_v56_, d_1127_v56_])
            d_1065_v1_ = (d_1128_v57_).ispropersubset(_dafny.MultiSet(d_1129_v58_))
            d_1130_v59_: C1
            nw236_ = C1()
            nw236_.ctor__((self).f4, (self).f5)
            d_1130_v59_ = nw236_
            d_1131_v60_: _dafny.Seq
            d_1131_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cglegqjc"))
            d_1131_v60_ = d_1131_v60_
            d_1132_v61_: D6
            d_1132_v61_ = D6_DC15(d_1065_v1_, not (d_1065_v1_) or (d_1065_v1_), d_1065_v1_)
            d_1132_v61_ = d_1132_v61_
        d_1065_v1_ = (d_1067_v3_) >= (len((d_1063_v0_)[default__.safeIndex(25, (d_1063_v0_).length(0))]))
        r0 = (d_1119_v51_).set(d_1067_v3_, not(d_1065_v1_))
        return r0

    def m12(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: int = int(0)
        d_1133_v0_: bool
        d_1133_v0_ = False
        d_1134_v1_: _dafny.Seq
        d_1134_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ms"))
        d_1135_v2_: _dafny.Map
        d_1135_v2_ = _dafny.Map({d_1133_v0_: d_1134_v1_})
        d_1136_v3_: _dafny.MultiSet
        d_1136_v3_ = _dafny.MultiSet([p0, p0, len(d_1135_v2_), p0])
        d_1137_v4_: _dafny.MultiSet
        d_1137_v4_ = _dafny.MultiSet([False, True])
        d_1138_v5_: _dafny.Seq
        d_1138_v5_ = _dafny.SeqWithoutIsStrInference([True])
        d_1139_v6_: _dafny.Seq
        d_1139_v6_ = _dafny.SeqWithoutIsStrInference([(d_1138_v5_)[default__.safeIndex((self).fm4(_dafny.SeqWithoutIsStrInference([p0]), globalState), len(d_1138_v5_))]])
        d_1140_v7_: _dafny.Array
        nw237_ = _dafny.Array(None, 10)
        nw237_[int(0)] = p0
        nw237_[int(1)] = default__.safeDivisionInt(p0, p0)
        nw237_[int(2)] = p0
        nw237_[int(3)] = default__.safeModuloInt((d_1136_v3_).cardinality, p0)
        nw237_[int(4)] = p0
        nw237_[int(5)] = p0
        nw237_[int(6)] = p0
        nw237_[int(7)] = ((d_1137_v4_)[d_1133_v0_] if (d_1133_v0_) in (d_1137_v4_) else p0)
        nw237_[int(8)] = ((default__.fm30(d_1133_v0_, d_1139_v6_, globalState)).set(p0, default__.abs(p0))).cardinality
        nw237_[int(9)] = p0
        d_1140_v7_ = nw237_
        index244_ = default__.safeIndex(933, (d_1140_v7_).length(0))
        (d_1140_v7_)[index244_] = len(d_1134_v1_)
        index245_ = default__.safeIndex(933, (d_1140_v7_).length(0))
        (d_1140_v7_)[index245_] = p0
        d_1141_v8_: _dafny.Array
        nw238_ = _dafny.Array(False, 23)
        d_1141_v8_ = nw238_
        (globalState).f3 = d_1141_v8_
        d_1142_v9_: C0
        nw239_ = C0()
        nw239_.ctor__()
        d_1142_v9_ = nw239_
        d_1142_v9_ = d_1142_v9_
        hi8_ = (d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))]
        for d_1143_i0_ in range(p0, hi8_):
            d_1144_v10_: _dafny.Map
            d_1144_v10_ = _dafny.Map({p0: _dafny.Map({p0: (d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))]})})
            d_1145_v11_: D14
            d_1145_v11_ = D14_DC30(d_1144_v10_)
            pat_let_tv30_ = d_1144_v10_
            d_1146_v12_: _dafny.Seq
            def iife87_(_pat_let23_0):
                def iife88_(d_1147_dt__update__tmp_h0_):
                    def iife89_(_pat_let24_0):
                        def iife90_(d_1148_dt__update_hcf47_h0_):
                            return D14_DC30(d_1148_dt__update_hcf47_h0_)
                        return iife90_(_pat_let24_0)
                    return iife89_(pat_let_tv30_)
                return iife88_(_pat_let23_0)
            d_1146_v12_ = _dafny.SeqWithoutIsStrInference([iife87_(d_1145_v11_), d_1145_v11_, d_1145_v11_, d_1145_v11_, d_1145_v11_])
            d_1149_v13_: _dafny.Seq
            d_1149_v13_ = _dafny.SeqWithoutIsStrInference([d_1146_v12_, d_1146_v12_, d_1146_v12_, d_1146_v12_])
            d_1146_v12_ = (d_1149_v13_)[default__.safeIndex(d_1143_i0_, len(d_1149_v13_))]
            d_1150_v14_: _dafny.Set
            d_1150_v14_ = _dafny.Set({d_1133_v0_})
            index246_ = default__.safeIndex(933, (d_1140_v7_).length(0))
            rhs227_ = (_dafny.SeqWithoutIsStrInference([(0) - (default__.safeDivisionInt(p0, d_1143_i0_)), default__.safeDivisionInt((d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))], len(d_1150_v14_)), (d_1136_v3_).cardinality])).set(default__.safeIndex((d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))], len(_dafny.SeqWithoutIsStrInference([(0) - (default__.safeDivisionInt(p0, d_1143_i0_)), default__.safeDivisionInt((d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))], len(d_1150_v14_)), (d_1136_v3_).cardinality]))), len((d_1134_v1_ if d_1133_v0_ else d_1134_v1_)))
            rhs228_ = (d_1143_i0_) + (d_1143_i0_)
            lhs121_ = globalState
            lhs122_ = d_1140_v7_
            lhs123_ = default__.safeIndex(933, (d_1140_v7_).length(0))
            lhs121_.f2 = rhs227_
            lhs122_[lhs123_] = rhs228_
            d_1151_v15_: C2
            nw240_ = C2()
            nw240_.ctor__((self).f4, default__.fm0(684, p0, globalState))
            d_1151_v15_ = nw240_
            d_1152_v16_: _dafny.Map
            d_1152_v16_ = _dafny.Map({d_1133_v0_: d_1133_v0_})
            d_1153_v17_: C1
            nw241_ = C1()
            nw241_.ctor__((self).f4, (self).f5)
            d_1153_v17_ = nw241_
            d_1154_v18_: _dafny.Map
            d_1154_v18_ = _dafny.Map({d_1133_v0_: d_1153_v17_})
            d_1155_v19_: _dafny.Map
            d_1155_v19_ = _dafny.Map({(d_1143_i0_ if d_1133_v0_ else len(d_1154_v18_)): (0) - ((d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))])})
            d_1156_v20_: _dafny.MultiSet
            d_1156_v20_ = _dafny.MultiSet([d_1141_v8_, d_1141_v8_, d_1141_v8_])
            index247_ = default__.safeIndex(933, (d_1140_v7_).length(0))
            def iife91_():
                coll41_ = _dafny.Map()
                compr_41_: int
                for compr_41_ in _dafny.IntegerRange(28, 15):
                    d_1157_v21_: int = compr_41_
                    if ((28) <= (d_1157_v21_)) and ((d_1157_v21_) < (15)):
                        coll41_[default__.safeDivisionInt(d_1157_v21_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nukgrcx"))))] = p0
                return _dafny.Map(coll41_)
            rhs229_ = (0) - (default__.safeModuloInt(d_1143_i0_, ((d_1156_v20_)[d_1141_v8_] if (d_1141_v8_) in (d_1156_v20_) else d_1143_i0_)))
            rhs230_ = d_1151_v15_
            rhs231_ = d_1152_v16_
            rhs232_ = not(not(((d_1152_v16_)[((0) - ((d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))])) >= ((0) - (d_1143_i0_))] if (((0) - ((d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))])) >= ((0) - (d_1143_i0_))) in (d_1152_v16_) else (False) and (d_1133_v0_))))
            rhs233_ = ((d_1155_v19_) | (_dafny.Map({(d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))]: p0}))) | ((d_1155_v19_) | (iife91_()
            ))
            lhs124_ = d_1140_v7_
            lhs125_ = default__.safeIndex(933, (d_1140_v7_).length(0))
            lhs124_[lhs125_] = rhs229_
            d_1151_v15_ = rhs230_
            d_1152_v16_ = rhs231_
            d_1133_v0_ = rhs232_
            d_1155_v19_ = rhs233_
            d_1158_v22_: str
            d_1158_v22_ = _dafny.CodePoint('e')
            d_1158_v22_ = d_1158_v22_
        d_1133_v0_ = d_1133_v0_
        r1 = d_1133_v0_
        r0 = default__.fm16(False, p0, globalState)
        r1 = d_1133_v0_
        d_1159_v23_: _dafny.Map
        d_1159_v23_ = _dafny.Map({d_1133_v0_: default__.safeDivisionInt(-68, (d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))])})
        r2 = ((d_1159_v23_)[False] if (False) in (d_1159_v23_) else default__.safeDivisionInt(p0, (d_1140_v7_)[default__.safeIndex(933, (d_1140_v7_).length(0))]))
        return r0, r1, r2


class C4(T1, T0):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        self.f11: int = int(0)
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f11, f12, f4, f5):
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f4 = f4
        (self)._f5 = f5

    def fm4(self, p0, globalState):
        return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([not(False), False])), self.f11)

    def fm5(self, p0, p1, p2, p3, globalState):
        if (False if not(True) else not(True)):
            return True
        elif True:
            return (_dafny.SeqWithoutIsStrInference([True])) < (_dafny.SeqWithoutIsStrInference([True]))

    def fm3(self, p0, p1, globalState):
        return not(((361) + ((self).f12)) > (default__.safeModuloInt(self.f11, len(_dafny.Map({False: True})))))

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        (self).f11 = self.f11
        d_1160_v0_: bool
        d_1160_v0_ = True
        d_1160_v0_ = not((default__.safeModuloInt(self.f11, self.f11)) <= ((self).f12))
        d_1161_v1_: _dafny.Seq
        d_1161_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])
        d_1162_v2_: _dafny.Seq
        d_1162_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhg"))
        rhs234_ = d_1161_v1_
        rhs235_ = (d_1162_v2_) != (d_1162_v2_)
        d_1161_v1_ = rhs234_
        d_1160_v0_ = rhs235_
        d_1163_v3_: _dafny.Map
        d_1163_v3_ = _dafny.Map({(self).f12: d_1160_v0_})
        d_1164_v4_: _dafny.MultiSet
        d_1164_v4_ = _dafny.MultiSet([self.f11, (0) - (self.f11), 691, (self).f12, len(d_1163_v3_)])
        d_1164_v4_ = d_1164_v4_
        hi9_ = ((0) - ((self).f12) if d_1160_v0_ else (0) - (self.f11))
        for d_1165_i0_ in range((self).f12, hi9_):
            d_1166_v5_: _dafny.Array
            nw242_ = _dafny.Array(False, 29)
            d_1166_v5_ = nw242_
            index248_ = default__.safeIndex(385, (d_1166_v5_).length(0))
            (d_1166_v5_)[index248_] = d_1160_v0_
            d_1167_v6_: bool
            d_1167_v6_ = not(d_1160_v0_)
            d_1168_v7_: _dafny.Seq
            d_1168_v7_ = _dafny.SeqWithoutIsStrInference([d_1160_v0_, (d_1167_v6_), d_1160_v0_])
            index249_ = default__.safeIndex(385, (d_1166_v5_).length(0))
            rhs236_ = d_1160_v0_
            rhs237_ = d_1168_v7_
            lhs126_ = d_1166_v5_
            lhs127_ = default__.safeIndex(385, (d_1166_v5_).length(0))
            lhs126_[lhs127_] = rhs236_
            d_1168_v7_ = rhs237_
            d_1169_v8_: _dafny.Set
            d_1169_v8_ = _dafny.Set({d_1166_v5_, d_1166_v5_})
            d_1170_v9_: _dafny.Set
            d_1170_v9_ = d_1169_v8_
            d_1171_v13_: _dafny.Array
            nw243_ = _dafny.Array(None, 17)
            def iife92_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(919, 303):
                    d_1172_v10_: int = compr_42_
                    if ((919) <= (d_1172_v10_)) and ((d_1172_v10_) < (303)):
                        coll42_[default__.safeDivisionInt(d_1172_v10_, d_1165_i0_)] = True
                return _dafny.Map(coll42_)
            nw243_[int(0)] = iife92_()
            
            nw243_[int(1)] = d_1163_v3_
            nw243_[int(2)] = d_1163_v3_
            nw243_[int(3)] = d_1163_v3_
            nw243_[int(4)] = d_1163_v3_
            nw243_[int(5)] = d_1163_v3_
            nw243_[int(6)] = _dafny.Map({d_1165_i0_: d_1160_v0_})
            nw243_[int(7)] = (d_1163_v3_).set(d_1165_i0_, ((d_1163_v3_)[self.f11] if (self.f11) in (d_1163_v3_) else d_1160_v0_))
            nw243_[int(8)] = d_1163_v3_
            nw243_[int(9)] = d_1163_v3_
            def iife93_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(737, -102):
                    d_1173_v11_: int = compr_43_
                    if ((737) <= (d_1173_v11_)) and ((d_1173_v11_) < (-102)):
                        coll43_[default__.safeModuloInt(d_1173_v11_, self.f11)] = d_1160_v0_
                return _dafny.Map(coll43_)
            nw243_[int(10)] = iife93_()
            
            nw243_[int(11)] = d_1163_v3_
            nw243_[int(12)] = d_1163_v3_
            def iife94_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in _dafny.IntegerRange(-117, -575):
                    d_1174_v12_: int = compr_44_
                    if ((-117) <= (d_1174_v12_)) and ((d_1174_v12_) < (-575)):
                        coll44_[(d_1174_v12_) - (d_1165_i0_)] = (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]
                return _dafny.Map(coll44_)
            nw243_[int(13)] = iife94_()
            
            nw243_[int(14)] = d_1163_v3_
            nw243_[int(15)] = d_1163_v3_
            nw243_[int(16)] = d_1163_v3_
            d_1171_v13_ = nw243_
            d_1175_v14_: D4
            d_1175_v14_ = D4_DC10(d_1170_v9_, len(_dafny.SeqWithoutIsStrInference([(self).f12 for d_1176_i1_ in range(default__.abs(251))])), d_1162_v2_, d_1171_v13_)
            d_1177_v15_: _dafny.Map
            d_1177_v15_ = _dafny.Map({d_1166_v5_: (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]})
            if (default__.safeDivisionInt(d_1165_i0_, (d_1175_v14_).cf12)) <= (len(d_1177_v15_)):
                d_1178_v16_: _dafny.Array
                def lambda75_(d_1179_v0_):
                    def lambda76_(d_1180_i2_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1179_v0_])])

                    return lambda76_

                init42_ = lambda75_(d_1160_v0_)
                nw244_ = _dafny.Array(None, 13)
                for i0_42_ in range(nw244_.length(0)):
                    nw244_[i0_42_] = init42_(i0_42_)
                d_1178_v16_ = nw244_
                d_1181_v17_: _dafny.MultiSet
                d_1181_v17_ = _dafny.MultiSet([(d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))], (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]])
                d_1182_v18_: _dafny.Seq
                d_1182_v18_ = _dafny.SeqWithoutIsStrInference([d_1181_v17_])
                index250_ = default__.safeIndex(865, (d_1178_v16_).length(0))
                (d_1178_v16_)[index250_] = d_1182_v18_
                d_1183_v19_: D1
                d_1183_v19_ = D1_DC1(d_1181_v17_)
                index251_ = default__.safeIndex(865, (d_1178_v16_).length(0))
                (d_1178_v16_)[index251_] = (_dafny.SeqWithoutIsStrInference([(d_1183_v19_).cf1, d_1181_v17_, d_1181_v17_])) + (d_1182_v18_)
                index252_ = default__.safeIndex(385, (d_1166_v5_).length(0))
                rhs238_ = True
                rhs239_ = default__.fm1((d_1165_i0_) * (d_1165_i0_), self.f11, (d_1160_v0_ if d_1160_v0_ else not(d_1160_v0_)), (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))], globalState)
                lhs128_ = d_1166_v5_
                lhs129_ = default__.safeIndex(385, (d_1166_v5_).length(0))
                lhs128_[lhs129_] = rhs238_
                d_1160_v0_ = rhs239_
                d_1184_v20_: _dafny.Array
                nw245_ = _dafny.Array(None, 3)
                nw245_[int(0)] = d_1162_v2_
                nw245_[int(1)] = d_1162_v2_
                nw245_[int(2)] = d_1162_v2_
                d_1184_v20_ = nw245_
                d_1184_v20_ = d_1184_v20_
                d_1185_v21_: _dafny.Seq
                d_1185_v21_ = _dafny.SeqWithoutIsStrInference([(self).f12])
                d_1186_v22_: _dafny.Seq
                d_1186_v22_ = _dafny.SeqWithoutIsStrInference([self.f11, len(d_1185_v21_)])
                d_1187_v23_: _dafny.Set
                d_1187_v23_ = _dafny.Set({(D6_DC14(d_1185_v21_)).cf19})
                d_1187_v23_ = d_1187_v23_
                d_1160_v0_ = (self.f11) == (d_1165_i0_)
            elif True:
                index253_ = default__.safeIndex(385, (d_1166_v5_).length(0))
                (d_1166_v5_)[index253_] = not(not(not(not((d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]))))
                d_1188_v24_: _dafny.Array
                nw246_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_1188_v24_ = nw246_
                d_1189_v25_: _dafny.Array
                nw247_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_1189_v25_ = nw247_
                d_1190_v26_: _dafny.Seq
                d_1190_v26_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1189_v25_, d_1189_v25_, d_1189_v25_])).cardinality])
                d_1191_v27_: _dafny.Map
                d_1191_v27_ = _dafny.Map({d_1160_v0_: d_1162_v2_})
                d_1192_v28_: _dafny.Map
                d_1192_v28_ = _dafny.Map({d_1188_v24_: (self).fm5(len((d_1190_v26_).set(default__.safeIndex(self.f11, len(d_1190_v26_)), 755)), not(d_1160_v0_), len(_dafny.Set({not((d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]), d_1160_v0_})), len(d_1191_v27_), globalState)})
                d_1192_v28_ = (d_1192_v28_).set((d_1188_v24_ if (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))] else d_1188_v24_), (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))])
                d_1193_v29_: _dafny.Map
                d_1193_v29_ = _dafny.Map({(self).f5: (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]})
                d_1194_v30_: _dafny.Map
                d_1194_v30_ = _dafny.Map({(self).f5: len(d_1168_v7_)})
                d_1195_v31_: _dafny.Array
                nw248_ = _dafny.Array(None, 4)
                nw248_[int(0)] = self.f11
                nw248_[int(1)] = len(d_1194_v30_)
                nw248_[int(2)] = (self).f12
                nw248_[int(3)] = default__.safeModuloInt(d_1165_i0_, (self).fm4(d_1190_v26_, globalState))
                d_1195_v31_ = nw248_
                rhs240_ = d_1193_v29_
                rhs241_ = d_1195_v31_
                rhs242_ = (154) < (437)
                rhs243_ = (-847) - ((len(d_1162_v2_)) - ((self).f12))
                lhs130_ = self
                d_1193_v29_ = rhs240_
                d_1195_v31_ = rhs241_
                d_1160_v0_ = rhs242_
                lhs130_.f11 = rhs243_
                index254_ = default__.safeIndex(208, (d_1195_v31_).length(0))
                (d_1195_v31_)[index254_] = default__.safeModuloInt((self).f12, self.f11)
                index255_ = default__.safeIndex(208, (d_1195_v31_).length(0))
                rhs244_ = d_1189_v25_
                rhs245_ = (d_1195_v31_ if (self.f11) <= ((self).f12) else d_1195_v31_)
                rhs246_ = (869) - (d_1165_i0_)
                rhs247_ = (d_1166_v5_)[default__.safeIndex(385, (d_1166_v5_).length(0))]
                lhs131_ = d_1195_v31_
                lhs132_ = default__.safeIndex(208, (d_1195_v31_).length(0))
                d_1189_v25_ = rhs244_
                d_1195_v31_ = rhs245_
                lhs131_[lhs132_] = rhs246_
                d_1160_v0_ = rhs247_
                (self).f11 = self.f11
            d_1196_v32_: _dafny.Array
            def lambda77_(d_1197_i3_):
                return (d_1197_i3_) * ((self).f12)

            init43_ = lambda77_
            nw249_ = _dafny.Array(None, 20)
            for i0_43_ in range(nw249_.length(0)):
                nw249_[i0_43_] = init43_(i0_43_)
            d_1196_v32_ = nw249_
            d_1198_v33_: _dafny.Map
            d_1198_v33_ = _dafny.Map({d_1162_v2_: d_1196_v32_})
            d_1199_v34_: _dafny.Seq
            d_1199_v34_ = _dafny.SeqWithoutIsStrInference([len(d_1198_v33_), len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1200_i4_ in range(default__.abs(56))]))])
            d_1201_v35_: _dafny.Map
            d_1201_v35_ = _dafny.Map({(self).f12: default__.safeDivisionInt((0) - ((d_1199_v34_)[default__.safeIndex(len(d_1168_v7_), len(d_1199_v34_))]), 537)})
            d_1201_v35_ = (d_1201_v35_).set(self.f11, 916)
            d_1202_v36_: T1
            nw250_ = C1()
            nw250_.ctor__((self).f4, (self).f5)
            d_1202_v36_ = nw250_
            d_1202_v36_ = d_1202_v36_
        d_1203_v37_: _dafny.Array
        nw251_ = _dafny.Array(None, 27)
        nw251_[int(0)] = (self).f5
        nw251_[int(1)] = _dafny.CodePoint('e')
        nw251_[int(2)] = (self).f5
        nw251_[int(3)] = (self).f5
        nw251_[int(4)] = (self).f5
        nw251_[int(5)] = (self).f5
        nw251_[int(6)] = _dafny.CodePoint('n')
        nw251_[int(7)] = (self).f5
        nw251_[int(8)] = (self).f5
        nw251_[int(9)] = (self).f5
        nw251_[int(10)] = (self).f5
        nw251_[int(11)] = (self).f5
        nw251_[int(12)] = (self).f5
        nw251_[int(13)] = (self).f5
        nw251_[int(14)] = (self).f5
        nw251_[int(15)] = (self).f5
        nw251_[int(16)] = default__.fm0(self.f11, (self).f12, globalState)
        nw251_[int(17)] = (self).f5
        nw251_[int(18)] = (self).f5
        nw251_[int(19)] = (self).f5
        nw251_[int(20)] = (self).f5
        nw251_[int(21)] = (self).f5
        nw251_[int(22)] = _dafny.CodePoint('a')
        nw251_[int(23)] = (self).f5
        nw251_[int(24)] = (self).f5
        nw251_[int(25)] = default__.fm0(294, len(_dafny.Map({(self).f12: len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1204_i5_ in range(default__.abs(497))]))})), globalState)
        nw251_[int(26)] = (self).f5
        d_1203_v37_ = nw251_
        index256_ = default__.safeIndex(499, (d_1203_v37_).length(0))
        (d_1203_v37_)[index256_] = (self).f5
        index257_ = default__.safeIndex(499, (d_1203_v37_).length(0))
        (d_1203_v37_)[index257_] = (self).f5
        d_1205_v38_: _dafny.Seq
        d_1205_v38_ = _dafny.SeqWithoutIsStrInference([(self).f12])
        d_1206_v39_: D6
        d_1206_v39_ = D6_DC14(d_1205_v38_)
        r0 = (default__.fm31((self).f12, (self).f12, globalState)) + ((d_1206_v39_).cf19)
        r1 = (self).fm4(_dafny.SeqWithoutIsStrInference([self.f11]), globalState)
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1207_v0_: _dafny.Array
        def lambda78_(d_1208_i1_):
            return False

        init44_ = lambda78_
        nw252_ = _dafny.Array(None, 23)
        for i0_44_ in range(nw252_.length(0)):
            nw252_[i0_44_] = init44_(i0_44_)
        d_1207_v0_ = nw252_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_1207_v0_).length(0)):
            d_1209_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_1209_i0_)) and ((d_1209_i0_) < ((d_1207_v0_).length(0)))):
                (d_1207_v0_)[(d_1209_i0_)] = False
        d_1210_v1_: _dafny.Array
        def lambda79_(d_1211_i3_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))

        init45_ = lambda79_
        nw253_ = _dafny.Array(None, 4)
        for i0_45_ in range(nw253_.length(0)):
            nw253_[i0_45_] = init45_(i0_45_)
        d_1210_v1_ = nw253_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1210_v1_).length(0)):
            d_1212_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_1212_i2_)) and ((d_1212_i2_) < ((d_1210_v1_).length(0)))):
                (d_1210_v1_)[(d_1212_i2_)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baooybn"))
        d_1213_v2_: C0
        nw254_ = C0()
        nw254_.ctor__()
        d_1213_v2_ = nw254_
        d_1214_v3_: _dafny.Seq
        d_1214_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljvlju"))
        d_1214_v3_ = d_1214_v3_
        d_1207_v0_ = d_1207_v0_
        d_1215_v4_: _dafny.Array
        def lambda80_(d_1216_i4_):
            return (_dafny.Map({self.f11: (0) - ((self).f12)})) | (_dafny.Map({(self).f12: self.f11}))

        init46_ = lambda80_
        nw255_ = _dafny.Array(None, 16)
        for i0_46_ in range(nw255_.length(0)):
            nw255_[i0_46_] = init46_(i0_46_)
        d_1215_v4_ = nw255_
        index258_ = default__.safeIndex(625, (d_1215_v4_).length(0))
        (d_1215_v4_)[index258_] = _dafny.Map({self.f11: -171})
        d_1217_v5_: _dafny.Seq
        d_1217_v5_ = _dafny.SeqWithoutIsStrInference([self.f11, 55])
        d_1218_v6_: bool
        d_1218_v6_ = True
        d_1219_v7_: _dafny.Set
        d_1219_v7_ = _dafny.Set({d_1218_v6_, d_1218_v6_})
        d_1220_v8_: _dafny.MultiSet
        d_1220_v8_ = _dafny.MultiSet([len(d_1219_v7_), (self).f12, self.f11])
        d_1221_v9_: _dafny.MultiSet
        d_1221_v9_ = _dafny.MultiSet([d_1218_v6_])
        d_1222_v10_: _dafny.MultiSet
        d_1222_v10_ = _dafny.MultiSet([(self).fm4(_dafny.SeqWithoutIsStrInference([(self).f12]), globalState), (d_1217_v5_)[default__.safeIndex(679, len(d_1217_v5_))], ((d_1220_v8_)[(d_1221_v9_).cardinality] if ((d_1221_v9_).cardinality) in (d_1220_v8_) else len(d_1219_v7_))])
        d_1223_v11_: _dafny.Map
        d_1223_v11_ = _dafny.Map({((d_1220_v8_)[(0) - (self.f11)] if ((0) - (self.f11)) in (d_1220_v8_) else self.f11): 369})
        index259_ = default__.safeIndex(625, (d_1215_v4_).length(0))
        rhs248_ = len(d_1214_v3_)
        rhs249_ = ((_dafny.Map({self.f11: (self).f12})) | (d_1223_v11_)).set((self).f12, 267)
        lhs133_ = self
        lhs134_ = d_1215_v4_
        lhs135_ = default__.safeIndex(625, (d_1215_v4_).length(0))
        lhs133_.f11 = rhs248_
        lhs134_[lhs135_] = rhs249_
        d_1224_v12_: _dafny.Array
        nw256_ = _dafny.Array(int(0), 7)
        d_1224_v12_ = nw256_
        d_1225_v13_: D11
        d_1225_v13_ = D11_DC25(d_1218_v6_, self.f11, d_1224_v12_, (self).f12)
        r0 = (d_1225_v13_).cf40
        return r0

    @property
    def f12(self):
        return self._f12

class C5:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def m10(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        r3: int = int(0)
        d_1226_v0_: _dafny.Array
        def lambda81_(d_1227_p2_):
            def lambda82_(d_1228_i0_):
                return d_1227_p2_

            return lambda82_

        init47_ = lambda81_(p2)
        nw257_ = _dafny.Array(None, 11)
        for i0_47_ in range(nw257_.length(0)):
            nw257_[i0_47_] = init47_(i0_47_)
        d_1226_v0_ = nw257_
        d_1229_v1_: _dafny.Map
        d_1229_v1_ = _dafny.Map({260: d_1226_v0_})
        d_1229_v1_ = (d_1229_v1_ if p2 else d_1229_v1_)
        d_1230_v2_: _dafny.MultiSet
        d_1230_v2_ = _dafny.MultiSet([p0])
        d_1231_v3_: _dafny.Seq
        d_1231_v3_ = _dafny.SeqWithoutIsStrInference([d_1230_v2_, d_1230_v2_])
        r1 = (((d_1231_v3_)[default__.safeIndex(761, len(d_1231_v3_))]).cardinality) * (p0)
        d_1232_v4_: _dafny.Seq
        d_1232_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcyi"))
        r0 = d_1232_v4_
        d_1233_v5_: _dafny.Array
        nw258_ = _dafny.Array(int(0), 25)
        d_1233_v5_ = nw258_
        index260_ = default__.safeIndex(979, (d_1233_v5_).length(0))
        (d_1233_v5_)[index260_] = p0
        index261_ = default__.safeIndex(979, (d_1233_v5_).length(0))
        rhs250_ = (d_1232_v4_) + ((d_1232_v4_ if p2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbpbcw"))))
        rhs251_ = p0
        rhs252_ = (p0) + (p0)
        lhs136_ = d_1233_v5_
        lhs137_ = default__.safeIndex(979, (d_1233_v5_).length(0))
        d_1232_v4_ = rhs250_
        lhs136_[lhs137_] = rhs251_
        r1 = rhs252_
        d_1234_v6_: _dafny.Map
        d_1234_v6_ = _dafny.Map({p2: (0) - (p0)})
        hi10_ = ((d_1234_v6_)[not(p2)] if (not(p2)) in (d_1234_v6_) else p0)
        for d_1235_i1_ in range(p0, hi10_):
            d_1236_v7_: bool
            d_1236_v7_ = False
            d_1236_v7_ = not (d_1236_v7_) or (p2)
            d_1237_v8_: _dafny.Array
            def lambda83_(d_1238_i1_, d_1239_p2_, d_1240_v5_, d_1241_p0_):
                def lambda84_(d_1242_i2_):
                    return (_dafny.SeqWithoutIsStrInference([d_1238_i1_, len(_dafny.SeqWithoutIsStrInference([d_1239_p2_])), -80, (d_1240_v5_)[default__.safeIndex(979, (d_1240_v5_).length(0))], d_1241_p0_])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1240_v5_)[default__.safeIndex(979, (d_1240_v5_).length(0))]: (d_1240_v5_)[default__.safeIndex(979, (d_1240_v5_).length(0))]}))]))

                return lambda84_

            init48_ = lambda83_(d_1235_i1_, p2, d_1233_v5_, p0)
            nw259_ = _dafny.Array(None, 29)
            for i0_48_ in range(nw259_.length(0)):
                nw259_[i0_48_] = init48_(i0_48_)
            d_1237_v8_ = nw259_
            d_1243_v9_: _dafny.Seq
            d_1243_v9_ = _dafny.SeqWithoutIsStrInference([d_1235_i1_])
            index262_ = default__.safeIndex(886, (d_1237_v8_).length(0))
            (d_1237_v8_)[index262_] = (d_1243_v9_) + (d_1243_v9_)
            index263_ = default__.safeIndex(886, (d_1237_v8_).length(0))
            (d_1237_v8_)[index263_] = d_1243_v9_
            d_1244_v10_: _dafny.Map
            d_1244_v10_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1235_i1_, d_1235_i1_]): (0) - (d_1235_i1_)})
            d_1245_v11_: D6
            d_1245_v11_ = D6_DC14(_dafny.SeqWithoutIsStrInference([(d_1233_v5_)[default__.safeIndex(979, (d_1233_v5_).length(0))]]))
            r3 = len(((d_1244_v10_).set((d_1245_v11_).cf19, p0)) | (d_1244_v10_))
            d_1246_v12_: _dafny.Set
            d_1246_v12_ = _dafny.Set({846, default__.fm14(globalState)})
            if ((d_1246_v12_).isdisjoint(d_1246_v12_)) or (not(p2)):
                index264_ = default__.safeIndex(979, (d_1233_v5_).length(0))
                rhs253_ = d_1232_v4_
                rhs254_ = d_1235_i1_
                lhs138_ = d_1233_v5_
                lhs139_ = default__.safeIndex(979, (d_1233_v5_).length(0))
                d_1232_v4_ = rhs253_
                lhs138_[lhs139_] = rhs254_
                d_1247_v13_: _dafny.Array
                nw260_ = _dafny.Array(D5.default()(), 29)
                d_1247_v13_ = nw260_
                d_1248_v14_: D5
                d_1248_v14_ = D5_DC12()
                index265_ = default__.safeIndex(36, (d_1247_v13_).length(0))
                (d_1247_v13_)[index265_] = d_1248_v14_
                index266_ = default__.safeIndex(36, (d_1247_v13_).length(0))
                (d_1247_v13_)[index266_] = D5_DC12()
                d_1249_v15_: _dafny.Seq
                d_1249_v15_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1250_v16_: D2
                d_1250_v16_ = D2_DC3(d_1226_v0_)
                d_1251_v17_: _dafny.Map
                d_1251_v17_ = _dafny.Map({d_1250_v16_: p0})
                d_1252_v18_: _dafny.Set
                d_1252_v18_ = _dafny.Set({d_1234_v6_, d_1234_v6_})
                index267_ = default__.safeIndex(979, (d_1233_v5_).length(0))
                def iife95_():
                    coll45_ = _dafny.Set()
                    compr_45_: _dafny.Map
                    for compr_45_ in (d_1252_v18_).Elements:
                        d_1253_v19_: _dafny.Map = compr_45_
                        if (d_1253_v19_) in (d_1252_v18_):
                            coll45_ = coll45_.union(_dafny.Set([d_1253_v19_]))
                    return _dafny.Set(coll45_)
                rhs255_ = not((d_1235_i1_) <= (len(d_1249_v15_)))
                rhs256_ = (0) - (p0)
                rhs257_ = len(d_1251_v17_)
                rhs258_ = (p0) * (default__.fm14(globalState))
                rhs259_ = ((d_1252_v18_) != (iife95_()
                ) if ((d_1233_v5_)[default__.safeIndex(979, (d_1233_v5_).length(0))]) > ((d_1233_v5_)[default__.safeIndex(979, (d_1233_v5_).length(0))]) else False)
                lhs140_ = d_1233_v5_
                lhs141_ = default__.safeIndex(979, (d_1233_v5_).length(0))
                d_1236_v7_ = rhs255_
                lhs140_[lhs141_] = rhs256_
                r1 = rhs257_
                r1 = rhs258_
                d_1236_v7_ = rhs259_
                index268_ = default__.safeIndex(135, (d_1226_v0_).length(0))
                (d_1226_v0_)[index268_] = d_1236_v7_
                index269_ = default__.safeIndex(135, (d_1226_v0_).length(0))
                (d_1226_v0_)[index269_] = d_1236_v7_
                d_1254_v20_: _dafny.Map
                d_1254_v20_ = _dafny.Map({p0: d_1236_v7_})
                d_1255_v21_: _dafny.Set
                d_1255_v21_ = _dafny.Set({((d_1254_v20_)[p0] if (p0) in (d_1254_v20_) else (d_1226_v0_)[default__.safeIndex(135, (d_1226_v0_).length(0))]), d_1236_v7_})
                d_1256_v22_: _dafny.Array
                nw261_ = _dafny.Array(None, 11)
                nw261_[int(0)] = default__.fm26((d_1233_v5_)[default__.safeIndex(979, (d_1233_v5_).length(0))], p0, globalState)
                nw261_[int(1)] = d_1255_v21_
                nw261_[int(2)] = d_1255_v21_
                nw261_[int(3)] = d_1255_v21_
                nw261_[int(4)] = d_1255_v21_
                nw261_[int(5)] = d_1255_v21_
                nw261_[int(6)] = d_1255_v21_
                nw261_[int(7)] = d_1255_v21_
                nw261_[int(8)] = d_1255_v21_
                nw261_[int(9)] = d_1255_v21_
                nw261_[int(10)] = d_1255_v21_
                d_1256_v22_ = nw261_
                d_1257_v23_: str
                d_1257_v23_ = _dafny.CodePoint('r')
                d_1258_v24_: C2
                nw262_ = C2()
                nw262_.ctor__(d_1256_v22_, d_1257_v23_)
                d_1258_v24_ = nw262_
            elif True:
                d_1259_v25_: _dafny.Map
                d_1259_v25_ = _dafny.Map({d_1235_i1_: len(_dafny.SeqWithoutIsStrInference([d_1236_v7_, d_1236_v7_, True, not(d_1236_v7_)]))})
                r3 = default__.safeModuloInt(len(d_1232_v4_), ((d_1259_v25_)[d_1235_i1_] if (d_1235_i1_) in (d_1259_v25_) else (d_1233_v5_)[default__.safeIndex(979, (d_1233_v5_).length(0))]))
                d_1260_v26_: _dafny.Set
                d_1260_v26_ = _dafny.Set({default__.fm16(p2, d_1235_i1_, globalState), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1261_i3_ in range(default__.abs(396))]), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okftu")) if False else d_1232_v4_)})
                d_1260_v26_ = d_1260_v26_
                def iife96_():
                    coll46_ = _dafny.Set()
                    compr_46_: int
                    for compr_46_ in (_dafny.MultiSet([(0) - (-187)])).Elements:
                        d_1262_v27_: int = compr_46_
                        if (d_1262_v27_) in (_dafny.MultiSet([(0) - (-187)])):
                            def iife97_():
                                coll47_ = _dafny.Set()
                                compr_47_: int
                                for compr_47_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(_dafny.MultiSet([940])).cardinality})) for d_1263_i4_ in range(default__.abs(549))])).Elements:
                                    d_1264_v28_: int = compr_47_
                                    if (d_1264_v28_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(_dafny.MultiSet([940])).cardinality})) for d_1263_i4_ in range(default__.abs(549))])):
                                        coll47_ = coll47_.union(_dafny.Set([(d_1264_v28_) - (-160)]))
                                return _dafny.Set(coll47_)
                            coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1262_v27_, len((D16_DC37(iife97_()
)).cf56))]))
                    return _dafny.Set(coll46_)
                d_1246_v12_ = (iife96_()
                ) - (d_1246_v12_)
                d_1265_v29_: D6
                d_1265_v29_ = D6_DC15(p2, d_1236_v7_, d_1236_v7_)
                d_1266_v30_: _dafny.Set
                d_1266_v30_ = _dafny.Set({_dafny.Map({d_1265_v29_: len(d_1260_v26_)})})
                d_1267_v31_: D8
                d_1267_v31_ = D8_DC19(len((d_1266_v30_).intersection(d_1266_v30_)), p2)
                d_1268_v32_: _dafny.Map
                d_1268_v32_ = _dafny.Map({d_1233_v5_: d_1232_v4_})
                d_1269_v33_: _dafny.Map
                d_1269_v33_ = _dafny.Map({d_1236_v7_: d_1232_v4_})
                rhs260_ = ((d_1230_v2_) | (d_1230_v2_)) | (d_1230_v2_)
                rhs261_ = d_1267_v31_
                rhs262_ = (p0) * (len(((d_1268_v32_)[d_1233_v5_] if (d_1233_v5_) in (d_1268_v32_) else ((d_1269_v33_)[p2] if (p2) in (d_1269_v33_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evqfa"))))))
                d_1230_v2_ = rhs260_
                d_1267_v31_ = rhs261_
                r3 = rhs262_
                d_1270_v34_: _dafny.Array
                nw263_ = _dafny.Array(_dafny.Map({}), 8)
                d_1270_v34_ = nw263_
                d_1271_v35_: _dafny.Map
                d_1271_v35_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itrexs")))): _dafny.CodePoint('u')})
                index270_ = default__.safeIndex(910, (d_1270_v34_).length(0))
                (d_1270_v34_)[index270_] = d_1271_v35_
                index271_ = default__.safeIndex(910, (d_1270_v34_).length(0))
                (d_1270_v34_)[index271_] = d_1271_v35_
        index272_ = default__.safeIndex(979, (d_1233_v5_).length(0))
        (d_1233_v5_)[index272_] = (d_1233_v5_)[default__.safeIndex(979, (d_1233_v5_).length(0))]
        d_1272_v36_: _dafny.MultiSet
        d_1272_v36_ = _dafny.MultiSet([not(p2)])
        r0 = default__.fm16(p2, (d_1272_v36_).cardinality, globalState)
        d_1273_v37_: _dafny.Seq
        d_1273_v37_ = _dafny.SeqWithoutIsStrInference([p2, p2])
        r1 = ((0) - (default__.safeModuloInt(len(d_1273_v37_), 945))) - (442)
        d_1274_v38_: str
        d_1274_v38_ = _dafny.CodePoint('s')
        r2 = (d_1274_v38_ if (d_1232_v4_) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1275_i5_ in range(default__.abs(598))])) else d_1274_v38_)
        r3 = default__.fm14(globalState)
        return r0, r1, r2, r3


class C6:
    def  __init__(self):
        self._f10: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f10):
        (self)._f10 = f10

    def m9(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_1276_v0_: int
        d_1276_v0_ = 266
        hi11_ = d_1276_v0_
        for d_1277_i0_ in range(d_1276_v0_, hi11_):
            d_1278_v1_: _dafny.Seq
            d_1278_v1_ = _dafny.SeqWithoutIsStrInference([p1])
            r2 = (d_1278_v1_)[default__.safeIndex(default__.fm13(p1, True, p1, globalState), len(d_1278_v1_))]
            d_1279_v3_: _dafny.Array
            def lambda85_(d_1280_i0_):
                def lambda86_(d_1281_i1_):
                    return (d_1281_i1_) * (d_1280_i0_)

                return lambda86_

            init49_ = lambda85_(d_1277_i0_)
            nw264_ = _dafny.Array(None, 17)
            for i0_49_ in range(nw264_.length(0)):
                nw264_[i0_49_] = init49_(i0_49_)
            d_1279_v3_ = nw264_
            d_1282_v4_: _dafny.Map
            d_1282_v4_ = _dafny.Map({_dafny.Set({(0) - (d_1276_v0_), d_1277_i0_, (0) - (d_1277_i0_)}): d_1279_v3_})
            def iife98_():
                coll48_ = _dafny.Set()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(10, -345):
                    d_1283_v2_: int = compr_48_
                    if ((10) <= (d_1283_v2_)) and ((d_1283_v2_) < (-345)):
                        coll48_ = coll48_.union(_dafny.Set([(d_1283_v2_) + (d_1277_i0_)]))
                return _dafny.Set(coll48_)
            d_1276_v0_ = len((_dafny.Map({iife98_()
            : d_1279_v3_})) | (d_1282_v4_))
            d_1276_v0_ = d_1277_i0_
            r1 = (d_1277_i0_) <= (990)
        d_1284_v5_: _dafny.Seq
        d_1284_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcoyirdm"))
        d_1285_v6_: str
        d_1285_v6_ = _dafny.CodePoint('a')
        d_1286_v7_: _dafny.MultiSet
        d_1286_v7_ = _dafny.MultiSet([(d_1284_v5_).set(default__.safeIndex(d_1276_v0_, len(d_1284_v5_)), d_1285_v6_), d_1284_v5_])
        d_1287_v8_: D5
        d_1287_v8_ = D5_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbin")), d_1276_v0_, p1)
        hi12_ = d_1276_v0_
        for d_1288_i2_ in range(((d_1286_v7_)[(d_1287_v8_).cf16] if ((d_1287_v8_).cf16) in (d_1286_v7_) else d_1276_v0_), hi12_):
            d_1289_v9_: _dafny.Array
            nw265_ = _dafny.Array(False, 26)
            d_1289_v9_ = nw265_
            index273_ = default__.safeIndex(57, (d_1289_v9_).length(0))
            (d_1289_v9_)[index273_] = p1
            index274_ = default__.safeIndex(57, (d_1289_v9_).length(0))
            (d_1289_v9_)[index274_] = p1
            r1 = not((D7_DC17((d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))], p1, p1, (d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))])).cf27)
            d_1290_v10_: _dafny.Array
            def lambda87_(d_1291_v0_):
                def lambda88_(d_1292_i3_):
                    return default__.safeDivisionInt(d_1292_i3_, d_1291_v0_)

                return lambda88_

            init50_ = lambda87_(d_1276_v0_)
            nw266_ = _dafny.Array(None, 16)
            for i0_50_ in range(nw266_.length(0)):
                nw266_[i0_50_] = init50_(i0_50_)
            d_1290_v10_ = nw266_
            index275_ = default__.safeIndex(737, (d_1290_v10_).length(0))
            (d_1290_v10_)[index275_] = d_1276_v0_
            d_1293_v11_: _dafny.Map
            d_1293_v11_ = _dafny.Map({d_1284_v5_: d_1289_v9_})
            d_1294_v12_: _dafny.Seq
            d_1294_v12_ = _dafny.SeqWithoutIsStrInference([d_1289_v9_, ((d_1293_v11_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "woupfmq"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "woupfmq"))) in (d_1293_v11_) else d_1289_v9_), d_1289_v9_])
            d_1295_v13_: _dafny.Array
            nw267_ = _dafny.Array(None, 16)
            nw267_[int(0)] = d_1289_v9_
            nw267_[int(1)] = d_1289_v9_
            nw267_[int(2)] = ((d_1293_v11_)[d_1284_v5_] if (d_1284_v5_) in (d_1293_v11_) else d_1289_v9_)
            nw267_[int(3)] = d_1289_v9_
            nw267_[int(4)] = (d_1294_v12_)[default__.safeIndex(175, len(d_1294_v12_))]
            nw267_[int(5)] = d_1289_v9_
            nw267_[int(6)] = d_1289_v9_
            nw267_[int(7)] = d_1289_v9_
            nw267_[int(8)] = d_1289_v9_
            nw267_[int(9)] = (d_1289_v9_ if (d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))] else d_1289_v9_)
            nw267_[int(10)] = d_1289_v9_
            nw267_[int(11)] = d_1289_v9_
            nw267_[int(12)] = d_1289_v9_
            nw267_[int(13)] = (d_1294_v12_)[default__.safeIndex(d_1288_i2_, len(d_1294_v12_))]
            nw267_[int(14)] = d_1289_v9_
            nw267_[int(15)] = d_1289_v9_
            d_1295_v13_ = nw267_
            d_1296_v14_: _dafny.MultiSet
            d_1296_v14_ = _dafny.MultiSet([(d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))]])
            d_1297_v15_: _dafny.Map
            d_1297_v15_ = _dafny.Map({d_1296_v14_: d_1276_v0_})
            d_1298_v16_: _dafny.MultiSet
            d_1298_v16_ = _dafny.MultiSet([default__.fm13(p1, (d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))], p1, globalState), len(d_1297_v15_)])
            d_1299_v17_: _dafny.Seq
            d_1299_v17_ = _dafny.SeqWithoutIsStrInference([(((self).f10)[(d_1298_v16_).cardinality] if ((d_1298_v16_).cardinality) in ((self).f10) else d_1288_i2_), len(_dafny.SeqWithoutIsStrInference([497]))])
            d_1300_v18_: _dafny.Seq
            d_1300_v18_ = _dafny.SeqWithoutIsStrInference([d_1299_v17_, d_1299_v17_])
            d_1301_v20_: D2
            d_1301_v20_ = D2_DC5(d_1284_v5_)
            d_1302_v21_: _dafny.Map
            d_1302_v21_ = _dafny.Map({len(d_1299_v17_): d_1301_v20_})
            index276_ = default__.safeIndex(737, (d_1290_v10_).length(0))
            def iife99_():
                coll49_ = _dafny.Map()
                compr_49_: int
                for compr_49_ in (d_1298_v16_).Elements:
                    d_1305_v19_: int = compr_49_
                    if (d_1305_v19_) in (d_1298_v16_):
                        coll49_[default__.safeModuloInt(d_1305_v19_, d_1276_v0_)] = D2_DC5(d_1284_v5_)
                return _dafny.Map(coll49_)
            rhs263_ = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1288_i2_ for d_1303_i5_ in range(default__.abs(510))]) for d_1304_i4_ in range(default__.abs(-173))])) == ((d_1300_v18_ if not((d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))]) else d_1300_v18_))
            rhs264_ = d_1276_v0_
            rhs265_ = p1
            rhs266_ = default__.fm1(len((iife99_()
            ) | (d_1302_v21_)), len(_dafny.Map({(d_1289_v9_)[default__.safeIndex(57, (d_1289_v9_).length(0))]: d_1276_v0_})), p1, p1, globalState)
            rhs267_ = d_1295_v13_
            lhs142_ = d_1290_v10_
            lhs143_ = default__.safeIndex(737, (d_1290_v10_).length(0))
            r1 = rhs263_
            lhs142_[lhs143_] = rhs264_
            r2 = rhs265_
            r1 = rhs266_
            d_1295_v13_ = rhs267_
            d_1306_v22_: _dafny.Set
            d_1306_v22_ = _dafny.Set({d_1288_i2_})
            d_1306_v22_ = d_1306_v22_
        d_1307_v23_: _dafny.Map
        d_1307_v23_ = _dafny.Map({d_1284_v5_: p1})
        d_1307_v23_ = (d_1307_v23_).set(d_1284_v5_, p1)
        d_1308_v24_: _dafny.Set
        d_1308_v24_ = _dafny.Set({_dafny.CodePoint('s')})
        d_1309_v25_: _dafny.Map
        d_1309_v25_ = _dafny.Map({(d_1308_v24_) - (d_1308_v24_): d_1276_v0_})
        rhs268_ = default__.safeModuloInt(d_1276_v0_, d_1276_v0_)
        rhs269_ = p1
        rhs270_ = p1
        rhs271_ = (d_1276_v0_) - (d_1276_v0_)
        rhs272_ = d_1309_v25_
        d_1276_v0_ = rhs268_
        r2 = rhs269_
        r2 = rhs270_
        d_1276_v0_ = rhs271_
        d_1309_v25_ = rhs272_
        d_1310_v26_: _dafny.Array
        def lambda89_(d_1311_p1_):
            def lambda90_(d_1312_i6_):
                return _dafny.Set({d_1311_p1_, True})

            return lambda90_

        init51_ = lambda89_(p1)
        nw268_ = _dafny.Array(None, 19)
        for i0_51_ in range(nw268_.length(0)):
            nw268_[i0_51_] = init51_(i0_51_)
        d_1310_v26_ = nw268_
        d_1313_v27_: C1
        nw269_ = C1()
        nw269_.ctor__(d_1310_v26_, default__.fm0(d_1276_v0_, d_1276_v0_, globalState))
        d_1313_v27_ = nw269_
        d_1314_v28_: _dafny.MultiSet
        d_1314_v28_ = _dafny.MultiSet([p1])
        d_1315_i7_: int
        d_1315_i7_ = 0
        with _dafny.label("6"):
            while (d_1314_v28_).isdisjoint(d_1314_v28_):
                with _dafny.c_label("6"):
                    if (d_1315_i7_) >= (100):
                        raise _dafny.Break("6")
                    d_1315_i7_ = (d_1315_i7_) + (1)
                    d_1316_v29_: int
                    d_1317_v30_: int
                    d_1318_v31_: int
                    d_1319_v32_: bool
                    out13_: int
                    out14_: int
                    out15_: int
                    out16_: bool
                    out13_, out14_, out15_, out16_ = (d_1313_v27_).m14(default__.safeDivisionInt(d_1276_v0_, d_1276_v0_), d_1276_v0_, globalState)
                    d_1316_v29_ = out13_
                    d_1317_v30_ = out14_
                    d_1318_v31_ = out15_
                    d_1319_v32_ = out16_
                    d_1320_v33_: _dafny.Map
                    d_1320_v33_ = _dafny.Map({d_1319_v32_: d_1319_v32_})
                    d_1321_v34_: _dafny.Seq
                    d_1321_v34_ = _dafny.SeqWithoutIsStrInference([(d_1320_v33_) | (d_1320_v33_), (d_1320_v33_) | (d_1320_v33_), default__.fm2(796, d_1319_v32_, d_1319_v32_, d_1276_v0_, globalState)])
                    d_1321_v34_ = ((p0) + (p0)) + (_dafny.SeqWithoutIsStrInference([(d_1320_v33_).set(True, d_1319_v32_), d_1320_v33_]))
                    d_1322_v35_: _dafny.Array
                    nw270_ = _dafny.Array(int(0), 26)
                    d_1322_v35_ = nw270_
                    index277_ = default__.safeIndex(931, (d_1322_v35_).length(0))
                    (d_1322_v35_)[index277_] = (d_1316_v29_ if d_1319_v32_ else 5)
                    index278_ = default__.safeIndex(931, (d_1322_v35_).length(0))
                    (d_1322_v35_)[index278_] = default__.safeModuloInt((len(d_1284_v5_)) - (d_1316_v29_), d_1316_v29_)
                    d_1323_v36_: _dafny.Seq
                    d_1323_v36_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_1324_v37_: D11
                    d_1324_v37_ = D11_DC25(p1, d_1316_v29_, d_1322_v35_, len(d_1323_v36_))
                    source13_ = d_1324_v37_
                    if source13_.is_DC25:
                        d_1325___mcc_h0_ = source13_.cf38
                        d_1326___mcc_h1_ = source13_.cf39
                        d_1327___mcc_h2_ = source13_.cf40
                        d_1328___mcc_h3_ = source13_.cf41
                        d_1329_cf41_ = d_1328___mcc_h3_
                        d_1330_cf40_ = d_1327___mcc_h2_
                        d_1331_cf39_ = d_1326___mcc_h1_
                        d_1332_cf38_ = d_1325___mcc_h0_
                        d_1333_v38_: _dafny.Seq
                        d_1333_v38_ = _dafny.SeqWithoutIsStrInference([d_1276_v0_, d_1276_v0_])
                        d_1334_v39_: _dafny.Map
                        d_1334_v39_ = _dafny.Map({d_1320_v33_: (d_1313_v27_).fm5((d_1313_v27_).fm4(d_1333_v38_, globalState), d_1332_cf38_, d_1329_cf41_, d_1276_v0_, globalState)})
                        d_1334_v39_ = d_1334_v39_
                        index279_ = default__.safeIndex(931, (d_1322_v35_).length(0))
                        (d_1322_v35_)[index279_] = ((d_1322_v35_)[default__.safeIndex(931, (d_1322_v35_).length(0))]) * (d_1331_cf39_)
                        d_1335_v40_: _dafny.Array
                        nw271_ = _dafny.Array(_dafny.Map({}), 15)
                        d_1335_v40_ = nw271_
                        index280_ = default__.safeIndex(852, (d_1335_v40_).length(0))
                        def iife100_():
                            coll50_ = _dafny.Map()
                            compr_50_: int
                            for compr_50_ in (d_1333_v38_).Elements:
                                d_1336_v41_: int = compr_50_
                                if (d_1336_v41_) in (d_1333_v38_):
                                    coll50_[(d_1336_v41_) * (d_1276_v0_)] = len((d_1284_v5_).set(default__.safeIndex(d_1329_cf41_, len(d_1284_v5_)), d_1285_v6_))
                            return _dafny.Map(coll50_)
                        (d_1335_v40_)[index280_] = iife100_()
                        
                        d_1337_v42_: _dafny.Array
                        nw272_ = _dafny.Array(False, 8)
                        d_1337_v42_ = nw272_
                        index281_ = default__.safeIndex(852, (d_1335_v40_).length(0))
                        rhs273_ = d_1337_v42_
                        rhs274_ = (self).f10
                        rhs275_ = d_1332_cf38_
                        lhs144_ = globalState
                        lhs145_ = d_1335_v40_
                        lhs146_ = default__.safeIndex(852, (d_1335_v40_).length(0))
                        lhs144_.f3 = rhs273_
                        lhs145_[lhs146_] = rhs274_
                        d_1319_v32_ = rhs275_
                        d_1276_v0_ = d_1316_v29_
                    elif source13_.is_DC24:
                        d_1338___mcc_h4_ = source13_.cf37
                        d_1339_cf37_ = d_1338___mcc_h4_
                        r2 = False
                        r2 = (d_1319_v32_ if p1 else (d_1319_v32_) == (not(d_1319_v32_)))
                        d_1340_v43_: _dafny.Set
                        d_1340_v43_ = _dafny.Set({750})
                        d_1320_v33_ = (d_1320_v33_).set(d_1319_v32_, (d_1340_v43_).ispropersubset(d_1340_v43_))
                        d_1341_v44_: _dafny.Seq
                        d_1341_v44_ = _dafny.SeqWithoutIsStrInference([(default__.fm31(d_1276_v0_, d_1276_v0_, globalState)) + (default__.fm31(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axjc"))), d_1317_v30_, globalState))])
                        d_1342_v45_: _dafny.Seq
                        d_1342_v45_ = _dafny.SeqWithoutIsStrInference([(d_1322_v35_)[default__.safeIndex(931, (d_1322_v35_).length(0))], d_1317_v30_])
                        d_1316_v29_ = len((d_1341_v44_).set(default__.safeIndex((d_1322_v35_)[default__.safeIndex(931, (d_1322_v35_).length(0))], len(d_1341_v44_)), d_1342_v45_))
                    elif True:
                        d_1343___mcc_h5_ = source13_.cf42
                        d_1344_cf42_ = d_1343___mcc_h5_
                        index282_ = default__.safeIndex(931, (d_1322_v35_).length(0))
                        (d_1322_v35_)[index282_] = d_1316_v29_
                        d_1345_v46_: _dafny.Map
                        d_1345_v46_ = _dafny.Map({(d_1285_v6_) in (d_1284_v5_): d_1320_v33_})
                        d_1346_v47_: _dafny.Set
                        d_1346_v47_ = _dafny.Set({False})
                        d_1347_v48_: _dafny.Seq
                        d_1347_v48_ = _dafny.SeqWithoutIsStrInference([d_1322_v35_, d_1322_v35_])
                        d_1345_v46_ = (d_1345_v46_).set((d_1346_v47_).ispropersubset(d_1346_v47_), (default__.fm2(d_1318_v31_, p1, d_1319_v32_, (0) - (len(d_1347_v48_)), globalState)).set(d_1319_v32_, p1))
                        d_1319_v32_ = ((d_1323_v36_)[default__.safeIndex((0) - (len(_dafny.Set({len(d_1284_v5_)}))), len(d_1323_v36_))]) or (d_1319_v32_)
                        d_1348_v49_: _dafny.Map
                        d_1348_v49_ = _dafny.Map({_dafny.Map({p1: d_1318_v31_}): False})
                        d_1319_v32_ = (d_1317_v30_) < (len(d_1348_v49_))
                    pass
            pass
        d_1349_v51_: _dafny.Map
        d_1349_v51_ = _dafny.Map({p1: d_1276_v0_})
        d_1350_v52_: D2
        d_1350_v52_ = D2_DC6(len(d_1349_v51_))
        d_1351_v53_: _dafny.Seq
        d_1351_v53_ = _dafny.SeqWithoutIsStrInference([(d_1350_v52_).cf7])
        d_1352_v54_: _dafny.Map
        d_1352_v54_ = _dafny.Map({d_1351_v53_: len(_dafny.SeqWithoutIsStrInference([d_1276_v0_]))})
        def iife101_():
            coll51_ = _dafny.Map()
            compr_51_: _dafny.Seq
            for compr_51_ in (((d_1352_v54_).set(_dafny.SeqWithoutIsStrInference([d_1276_v0_]), d_1276_v0_) if p1 else d_1352_v54_)).keys.Elements:
                d_1353_v50_: _dafny.Seq = compr_51_
                if (d_1353_v50_) in (((d_1352_v54_).set(_dafny.SeqWithoutIsStrInference([d_1276_v0_]), d_1276_v0_) if p1 else d_1352_v54_)):
                    coll51_[d_1353_v50_] = not(p1)
            return _dafny.Map(coll51_)
        r0 = iife101_()
        
        d_1354_v55_: _dafny.Array
        def lambda91_(d_1355_p1_):
            def lambda92_(d_1356_i8_):
                return d_1355_p1_

            return lambda92_

        init52_ = lambda91_(p1)
        nw273_ = _dafny.Array(None, 20)
        for i0_52_ in range(nw273_.length(0)):
            nw273_[i0_52_] = init52_(i0_52_)
        d_1354_v55_ = nw273_
        d_1357_v56_: _dafny.Set
        d_1357_v56_ = _dafny.Set({d_1354_v55_})
        r1 = (d_1357_v56_).issubset(d_1357_v56_)
        r2 = p1
        d_1358_v57_: _dafny.Array
        def lambda93_(d_1359_v28_):
            def lambda94_(d_1360_i9_):
                return d_1359_v28_

            return lambda94_

        init53_ = lambda93_(d_1314_v28_)
        nw274_ = _dafny.Array(None, 17)
        for i0_53_ in range(nw274_.length(0)):
            nw274_[i0_53_] = init53_(i0_53_)
        d_1358_v57_ = nw274_
        r3 = d_1358_v57_
        return r0, r1, r2, r3

    @property
    def f10(self):
        return self._f10

class C7(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm3(self, p0, p1, globalState):
        return not (((D2_DC4(False)).cf5) == (True)) or (True)

    def fm10(self, p0, p1, p2, globalState):
        return 913

    def m8(self, p0, p1, p2, globalState):
        r0: bool = False
        hi13_ = p1
        for d_1361_i0_ in range(p1, hi13_):
            d_1362_v0_: str
            d_1362_v0_ = _dafny.CodePoint('c')
            d_1363_v1_: _dafny.Map
            d_1363_v1_ = _dafny.Map({d_1362_v0_: p2})
            d_1364_v2_: _dafny.Array
            def lambda95_(d_1365_p2_):
                def lambda96_(d_1366_i1_):
                    return d_1365_p2_

                return lambda96_

            init54_ = lambda95_(p2)
            nw275_ = _dafny.Array(None, 9)
            for i0_54_ in range(nw275_.length(0)):
                nw275_[i0_54_] = init54_(i0_54_)
            d_1364_v2_ = nw275_
            d_1367_v3_: _dafny.Map
            d_1367_v3_ = _dafny.Map({d_1364_v2_: p2})
            d_1363_v1_ = (d_1363_v1_).set(d_1362_v0_, default__.fm1(d_1361_i0_, d_1361_i0_, ((d_1367_v3_)[d_1364_v2_] if (d_1364_v2_) in (d_1367_v3_) else not(True)), p2, globalState))
            index283_ = default__.safeIndex(968, (d_1364_v2_).length(0))
            (d_1364_v2_)[index283_] = (self).fm3(d_1361_i0_, d_1361_i0_, globalState)
            d_1368_v4_: D1
            d_1368_v4_ = D1_DC2(p1, d_1361_i0_)
            index284_ = default__.safeIndex(968, (d_1364_v2_).length(0))
            (d_1364_v2_)[index284_] = (((0) - (d_1361_i0_)) != (299) if False else ((d_1368_v4_).cf2) <= (p1))
            d_1369_v5_: int
            d_1369_v5_ = 671
            d_1370_v6_: _dafny.Seq
            d_1370_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ph"))
            d_1371_v7_: _dafny.Map
            d_1371_v7_ = _dafny.Map({d_1370_v6_: default__.safeModuloInt(134, p1)})
            d_1369_v5_ = ((d_1371_v7_)[d_1370_v6_] if (d_1370_v6_) in (d_1371_v7_) else p1)
            d_1372_v8_: _dafny.Map
            d_1372_v8_ = _dafny.Map({d_1362_v0_: d_1362_v0_})
            d_1372_v8_ = (d_1372_v8_).set(d_1362_v0_, d_1362_v0_)
        if p2:
            r0 = p2
            d_1373_v9_: _dafny.Array
            nw276_ = _dafny.Array(D2.default()(), 29)
            d_1373_v9_ = nw276_
            d_1374_v10_: D2
            d_1374_v10_ = D2_DC4(p2)
            d_1375_v11_: D2
            d_1375_v11_ = D2_DC7(d_1374_v10_)
            d_1376_v12_: D2
            d_1376_v12_ = D2_DC7(d_1375_v11_)
            index285_ = default__.safeIndex(403, (d_1373_v9_).length(0))
            (d_1373_v9_)[index285_] = d_1376_v12_
            index286_ = default__.safeIndex(403, (d_1373_v9_).length(0))
            (d_1373_v9_)[index286_] = D2_DC7(d_1374_v10_)
            if default__.fm1(p1, p1, p2, not((p2) or (p2)), globalState):
                d_1377_v13_: str
                d_1377_v13_ = _dafny.CodePoint('i')
                d_1378_v14_: _dafny.Seq
                d_1378_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcfkl"))
                d_1379_v15_: _dafny.Map
                d_1379_v15_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1377_v13_, d_1377_v13_]): (d_1378_v14_) < (d_1378_v14_)})
                d_1379_v15_ = d_1379_v15_
                d_1380_v16_: _dafny.Array
                def lambda97_(d_1381_v14_, d_1382_p2_, d_1383_p1_):
                    def lambda98_(d_1384_i2_):
                        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-893, len(d_1381_v14_)]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1382_p2_: d_1383_p1_})), d_1383_p1_])))

                    return lambda98_

                init55_ = lambda97_(d_1378_v14_, p2, p1)
                nw277_ = _dafny.Array(None, 6)
                for i0_55_ in range(nw277_.length(0)):
                    nw277_[i0_55_] = init55_(i0_55_)
                d_1380_v16_ = nw277_
                d_1385_v17_: _dafny.MultiSet
                d_1385_v17_ = _dafny.MultiSet([p1])
                index287_ = default__.safeIndex(856, (d_1380_v16_).length(0))
                (d_1380_v16_)[index287_] = (d_1385_v17_) - (d_1385_v17_)
                index288_ = default__.safeIndex(856, (d_1380_v16_).length(0))
                (d_1380_v16_)[index288_] = default__.fm11(globalState)
                d_1386_v18_: _dafny.Seq
                d_1386_v18_ = _dafny.SeqWithoutIsStrInference([p1])
                (globalState).f2 = d_1386_v18_
                r0 = not((D5_DC13(d_1378_v14_, 531, p2)).cf18)
                r0 = p2
            elif True:
                d_1387_v19_: int
                d_1387_v19_ = 520
                d_1387_v19_ = 280
                d_1388_v20_: _dafny.MultiSet
                d_1388_v20_ = _dafny.MultiSet([p2])
                d_1389_v21_: D1
                d_1389_v21_ = D1_DC1(d_1388_v20_)
                d_1390_v22_: _dafny.Map
                d_1390_v22_ = _dafny.Map({p1: p2})
                d_1391_v23_: _dafny.Map
                d_1391_v23_ = _dafny.Map({(self).fm10(d_1389_v21_, p2, d_1390_v22_, globalState): d_1388_v20_})
                d_1392_v24_: _dafny.Seq
                d_1392_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxjqhtn"))
                d_1393_v25_: _dafny.MultiSet
                d_1393_v25_ = _dafny.MultiSet([d_1392_v24_])
                d_1394_v26_: D1
                d_1394_v26_ = D1_DC1(((d_1391_v23_)[(d_1393_v25_).cardinality] if ((d_1393_v25_).cardinality) in (d_1391_v23_) else d_1388_v20_))
                d_1394_v26_ = d_1389_v21_
                d_1387_v19_ = (p1) * (p1)
                d_1395_v27_: str
                d_1395_v27_ = _dafny.CodePoint('a')
                d_1396_v28_: _dafny.Seq
                d_1396_v28_ = _dafny.SeqWithoutIsStrInference([p1, p1, d_1387_v19_])
                d_1397_v29_: _dafny.Array
                nw278_ = _dafny.Array(None, 1)
                nw278_[int(0)] = (d_1396_v28_) + (d_1396_v28_)
                d_1397_v29_ = nw278_
                index289_ = default__.safeIndex(551, (d_1397_v29_).length(0))
                (d_1397_v29_)[index289_] = (d_1396_v28_).set(default__.safeIndex(d_1387_v19_, len(d_1396_v28_)), d_1387_v19_)
                index290_ = default__.safeIndex(551, (d_1397_v29_).length(0))
                rhs276_ = d_1395_v27_
                rhs277_ = d_1387_v19_
                rhs278_ = (d_1396_v28_) + ((D6_DC14(_dafny.SeqWithoutIsStrInference([d_1387_v19_, d_1387_v19_]))).cf19)
                lhs147_ = d_1397_v29_
                lhs148_ = default__.safeIndex(551, (d_1397_v29_).length(0))
                d_1395_v27_ = rhs276_
                d_1387_v19_ = rhs277_
                lhs147_[lhs148_] = rhs278_
                d_1398_v30_: _dafny.Map
                d_1398_v30_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([(D7_DC16(_dafny.CodePoint('j'))).cf23 for d_1399_i3_ in range(default__.abs(567))]))) != (d_1387_v19_): p1})
                d_1400_v31_: _dafny.Array
                def lambda99_(d_1401_p2_):
                    def lambda100_(d_1402_i4_):
                        return d_1401_p2_

                    return lambda100_

                init56_ = lambda99_(p2)
                nw279_ = _dafny.Array(None, 5)
                for i0_56_ in range(nw279_.length(0)):
                    nw279_[i0_56_] = init56_(i0_56_)
                d_1400_v31_ = nw279_
                index291_ = default__.safeIndex(568, (d_1400_v31_).length(0))
                (d_1400_v31_)[index291_] = p2
                d_1403_v32_: _dafny.MultiSet
                d_1403_v32_ = _dafny.MultiSet([(0) - (d_1387_v19_), (0) - (p1), d_1387_v19_])
                d_1404_v33_: _dafny.Map
                d_1404_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rblvmkb")): (0) - (((d_1403_v32_)[d_1387_v19_] if (d_1387_v19_) in (d_1403_v32_) else p1))})
                index292_ = default__.safeIndex(568, (d_1400_v31_).length(0))
                rhs279_ = d_1398_v30_
                rhs280_ = not(not((d_1403_v32_).ispropersubset(d_1403_v32_)))
                rhs281_ = (d_1395_v27_ if (D2_DC4(p2)).cf5 else (d_1392_v24_)[default__.safeIndex(len(_dafny.Map({d_1387_v19_: d_1387_v19_})), len(d_1392_v24_))])
                rhs282_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkagjtkuc")): p1})
                lhs149_ = d_1400_v31_
                lhs150_ = default__.safeIndex(568, (d_1400_v31_).length(0))
                d_1398_v30_ = rhs279_
                lhs149_[lhs150_] = rhs280_
                d_1395_v27_ = rhs281_
                d_1404_v33_ = rhs282_
            d_1405_v34_: int
            d_1405_v34_ = 879
            d_1405_v34_ = d_1405_v34_
            d_1406_v35_: _dafny.Seq
            d_1406_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vx"))
            d_1407_v36_: _dafny.Array
            nw280_ = _dafny.Array(None, 7)
            nw280_[int(0)] = d_1406_v35_
            nw280_[int(1)] = (d_1406_v35_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1408_i5_ in range(default__.abs(375))]))
            nw280_[int(2)] = d_1406_v35_
            nw280_[int(3)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1409_i6_ in range(default__.abs(182))]) if False else d_1406_v35_)
            nw280_[int(4)] = d_1406_v35_
            nw280_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpley"))
            nw280_[int(6)] = default__.fm12(globalState)
            d_1407_v36_ = nw280_
            d_1407_v36_ = d_1407_v36_
        elif True:
            d_1410_v37_: _dafny.Set
            d_1410_v37_ = _dafny.Set({p2})
            d_1411_v38_: _dafny.MultiSet
            d_1411_v38_ = _dafny.MultiSet([(d_1410_v37_).ispropersubset(d_1410_v37_)])
            d_1411_v38_ = d_1411_v38_
            d_1412_v39_: _dafny.Array
            nw281_ = _dafny.Array(int(0), 8)
            d_1412_v39_ = nw281_
            d_1413_v40_: _dafny.Map
            d_1413_v40_ = _dafny.Map({p1: p1})
            index293_ = default__.safeIndex(585, (d_1412_v39_).length(0))
            (d_1412_v39_)[index293_] = ((d_1413_v40_)[p1] if (p1) in (d_1413_v40_) else p1)
            index294_ = default__.safeIndex(585, (d_1412_v39_).length(0))
            (d_1412_v39_)[index294_] = (p1) * (p1)
            d_1414_v41_: _dafny.Array
            nw282_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1414_v41_ = nw282_
            d_1415_v42_: _dafny.Seq
            d_1415_v42_ = _dafny.SeqWithoutIsStrInference([d_1414_v41_])
            d_1416_v43_: _dafny.Seq
            d_1416_v43_ = _dafny.SeqWithoutIsStrInference([d_1414_v41_, (d_1415_v42_)[default__.safeIndex(p1, len(d_1415_v42_))]])
            d_1417_v44_: _dafny.Array
            nw283_ = _dafny.Array(None, 12)
            nw283_[int(0)] = d_1414_v41_
            nw283_[int(1)] = d_1414_v41_
            nw283_[int(2)] = d_1414_v41_
            nw283_[int(3)] = d_1414_v41_
            nw283_[int(4)] = d_1414_v41_
            nw283_[int(5)] = (d_1414_v41_ if p2 else d_1414_v41_)
            nw283_[int(6)] = d_1414_v41_
            nw283_[int(7)] = d_1414_v41_
            nw283_[int(8)] = d_1414_v41_
            nw283_[int(9)] = (d_1416_v43_)[default__.safeIndex((d_1412_v39_)[default__.safeIndex(585, (d_1412_v39_).length(0))], len(d_1416_v43_))]
            nw283_[int(10)] = d_1414_v41_
            nw283_[int(11)] = d_1414_v41_
            d_1417_v44_ = nw283_
            index295_ = default__.safeIndex(220, (d_1417_v44_).length(0))
            (d_1417_v44_)[index295_] = d_1414_v41_
            index296_ = default__.safeIndex(220, (d_1417_v44_).length(0))
            (d_1417_v44_)[index296_] = d_1414_v41_
            r0 = p2
            d_1418_v45_: _dafny.Seq
            d_1418_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsco"))
            d_1418_v45_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1419_i7_ in range(default__.abs(940))])
        d_1420_v46_: C5
        nw284_ = C5()
        nw284_.ctor__()
        d_1420_v46_ = nw284_
        r0 = p2
        d_1421_v47_: str
        d_1421_v47_ = _dafny.CodePoint('u')
        d_1422_v48_: _dafny.Seq
        d_1422_v48_ = _dafny.SeqWithoutIsStrInference([d_1421_v47_, d_1421_v47_])
        d_1423_v49_: _dafny.Seq
        d_1423_v49_ = _dafny.SeqWithoutIsStrInference([len(d_1422_v48_)])
        d_1424_v50_: D6
        d_1424_v50_ = D6_DC14(d_1423_v49_)
        source14_ = d_1424_v50_
        if source14_.is_DC15:
            d_1425___mcc_h0_ = source14_.cf20
            d_1426___mcc_h1_ = source14_.cf21
            d_1427___mcc_h2_ = source14_.cf22
            d_1428_cf22_ = d_1427___mcc_h2_
            d_1429_cf21_ = d_1426___mcc_h1_
            d_1430_cf20_ = d_1425___mcc_h0_
            d_1431_v51_: int
            d_1431_v51_ = 171
            rhs283_ = p1
            rhs284_ = d_1431_v51_
            rhs285_ = (0) - ((len(d_1423_v49_) if (d_1431_v51_) <= (len(d_1422_v48_)) else 915))
            d_1431_v51_ = rhs283_
            d_1431_v51_ = rhs284_
            d_1431_v51_ = rhs285_
            if d_1430_cf20_:
                d_1432_v52_: _dafny.Array
                nw285_ = _dafny.Array(_dafny.Set({}), 9)
                d_1432_v52_ = nw285_
                d_1433_v53_: C4
                nw286_ = C4()
                nw286_.ctor__(p1, (0) - (p1), d_1432_v52_, d_1421_v47_)
                d_1433_v53_ = nw286_
                d_1434_v54_: C3
                nw287_ = C3()
                nw287_.ctor__(d_1432_v52_, d_1421_v47_)
                d_1434_v54_ = nw287_
                d_1434_v54_ = d_1434_v54_
                r0 = (p0).isdisjoint((p0).intersection(p0))
                d_1435_v55_: _dafny.Map
                d_1435_v55_ = _dafny.Map({d_1433_v53_: _dafny.Set({(_dafny.SeqWithoutIsStrInference([d_1421_v47_ for d_1436_i8_ in range(default__.abs(-737))])).set(default__.safeIndex(d_1433_v53_.f11, len(_dafny.SeqWithoutIsStrInference([d_1421_v47_ for d_1437_i8_ in range(default__.abs(-737))]))), d_1421_v47_)})})
                d_1438_v56_: _dafny.Set
                d_1438_v56_ = _dafny.Set({d_1422_v48_})
                d_1435_v55_ = (d_1435_v55_).set(d_1433_v53_, d_1438_v56_)
                d_1439_v57_: _dafny.Array
                def lambda101_(d_1440_cf20_):
                    def lambda102_(d_1441_i9_):
                        return d_1440_cf20_

                    return lambda102_

                init57_ = lambda101_(d_1430_cf20_)
                nw288_ = _dafny.Array(None, 29)
                for i0_57_ in range(nw288_.length(0)):
                    nw288_[i0_57_] = init57_(i0_57_)
                d_1439_v57_ = nw288_
                d_1442_v58_: _dafny.Map
                d_1442_v58_ = _dafny.Map({d_1439_v57_: d_1428_cf22_})
                d_1443_v59_: _dafny.Map
                d_1443_v59_ = _dafny.Map({d_1442_v58_: d_1431_v51_})
                d_1443_v59_ = (d_1443_v59_).set(d_1442_v58_, default__.safeModuloInt((d_1433_v53_).f12, (d_1433_v53_).f12))
            elif True:
                d_1444_v60_: D7
                d_1444_v60_ = D7_DC17(True, d_1429_cf21_, d_1428_cf22_, d_1430_cf20_)
                d_1445_v61_: _dafny.Seq
                d_1445_v61_ = _dafny.SeqWithoutIsStrInference([d_1444_v60_, default__.fm32(d_1430_cf20_, d_1429_cf21_, globalState)])
                d_1445_v61_ = (d_1445_v61_ if d_1430_cf20_ else (d_1445_v61_) + ((d_1445_v61_).set(default__.safeIndex(p1, len(d_1445_v61_)), d_1444_v60_)))
                d_1446_v62_: _dafny.Array
                def lambda103_(d_1447_cf20_, d_1448_cf22_):
                    def lambda104_(d_1449_i10_):
                        return _dafny.Set({d_1447_cf20_, d_1448_cf22_})

                    return lambda104_

                init58_ = lambda103_(d_1430_cf20_, d_1428_cf22_)
                nw289_ = _dafny.Array(None, 16)
                for i0_58_ in range(nw289_.length(0)):
                    nw289_[i0_58_] = init58_(i0_58_)
                d_1446_v62_ = nw289_
                d_1450_v63_: C1
                nw290_ = C1()
                nw290_.ctor__(d_1446_v62_, _dafny.CodePoint('s'))
                d_1450_v63_ = nw290_
                d_1451_v64_: _dafny.Array
                nw291_ = _dafny.Array(_dafny.Seq({}), 8)
                d_1451_v64_ = nw291_
                d_1452_v65_: _dafny.Seq
                d_1452_v65_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1453_v66_: _dafny.Map
                d_1453_v66_ = _dafny.Map({-110: d_1429_cf21_})
                d_1454_v67_: _dafny.Map
                d_1454_v67_ = _dafny.Map({d_1428_cf22_: p1})
                d_1455_v68_: _dafny.MultiSet
                d_1455_v68_ = _dafny.MultiSet([d_1454_v67_, d_1454_v67_])
                d_1456_v69_: _dafny.Map
                d_1456_v69_ = _dafny.Map({(d_1455_v68_).cardinality: p1})
                d_1457_v70_: _dafny.Set
                d_1457_v70_ = _dafny.Set({p2})
                d_1458_v71_: _dafny.MultiSet
                d_1458_v71_ = _dafny.MultiSet([p2, p2])
                d_1459_v72_: _dafny.Array
                nw292_ = _dafny.Array(None, 24)
                nw292_[int(0)] = len(p0)
                nw292_[int(1)] = 755
                nw292_[int(2)] = len(d_1452_v65_)
                nw292_[int(3)] = len(_dafny.Map({True: d_1431_v51_}))
                nw292_[int(4)] = p1
                nw292_[int(5)] = d_1431_v51_
                nw292_[int(6)] = d_1431_v51_
                nw292_[int(7)] = d_1431_v51_
                nw292_[int(8)] = len(d_1422_v48_)
                nw292_[int(9)] = (0) - (d_1431_v51_)
                nw292_[int(10)] = p1
                nw292_[int(11)] = (0) - (d_1431_v51_)
                nw292_[int(12)] = len(d_1453_v66_)
                nw292_[int(13)] = 554
                nw292_[int(14)] = d_1431_v51_
                nw292_[int(15)] = (_dafny.MultiSet([p1])).cardinality
                nw292_[int(16)] = p1
                nw292_[int(17)] = p1
                nw292_[int(18)] = p1
                nw292_[int(19)] = d_1431_v51_
                nw292_[int(20)] = len(d_1456_v69_)
                nw292_[int(21)] = len(d_1457_v70_)
                nw292_[int(22)] = d_1431_v51_
                nw292_[int(23)] = (d_1458_v71_).cardinality
                d_1459_v72_ = nw292_
                d_1460_v73_: _dafny.Seq
                d_1460_v73_ = _dafny.SeqWithoutIsStrInference([d_1459_v72_])
                index297_ = default__.safeIndex(314, (d_1451_v64_).length(0))
                (d_1451_v64_)[index297_] = (_dafny.SeqWithoutIsStrInference([d_1459_v72_])) + (d_1460_v73_)
                d_1461_v74_: D1
                d_1461_v74_ = D1_DC1(_dafny.MultiSet([d_1428_cf22_]))
                index298_ = default__.safeIndex(314, (d_1451_v64_).length(0))
                rhs286_ = d_1460_v73_
                rhs287_ = _dafny.MultiSet([d_1429_cf21_])
                rhs288_ = p1
                rhs289_ = not((d_1450_v63_).fm3((d_1431_v51_ if d_1429_cf21_ else 498), default__.safeDivisionInt((0) - (p1), (self).fm10(d_1461_v74_, d_1428_cf22_, d_1453_v66_, globalState)), globalState))
                lhs151_ = d_1451_v64_
                lhs152_ = default__.safeIndex(314, (d_1451_v64_).length(0))
                lhs151_[lhs152_] = rhs286_
                d_1458_v71_ = rhs287_
                d_1431_v51_ = rhs288_
                d_1428_cf22_ = rhs289_
                d_1430_cf20_ = d_1429_cf21_
                d_1462_v75_: _dafny.Map
                d_1462_v75_ = _dafny.Map({d_1459_v72_: (d_1429_cf21_) and (d_1429_cf21_)})
                d_1462_v75_ = (d_1462_v75_).set(d_1459_v72_, d_1430_cf20_)
            if False:
                d_1463_v76_: _dafny.Map
                d_1463_v76_ = _dafny.Map({p1: d_1430_cf20_})
                d_1464_v77_: _dafny.Array
                nw293_ = _dafny.Array(None, 6)
                nw293_[int(0)] = len(d_1463_v76_)
                nw293_[int(1)] = d_1431_v51_
                nw293_[int(2)] = p1
                nw293_[int(3)] = 491
                nw293_[int(4)] = p1
                nw293_[int(5)] = d_1431_v51_
                d_1464_v77_ = nw293_
                d_1465_v78_: _dafny.Seq
                d_1465_v78_ = _dafny.SeqWithoutIsStrInference([d_1464_v77_])
                d_1465_v78_ = ((d_1465_v78_) + (d_1465_v78_)) + (d_1465_v78_)
                d_1431_v51_ = default__.safeModuloInt(len((d_1422_v48_) + (d_1422_v48_)), p1)
                d_1466_v79_: _dafny.Array
                nw294_ = _dafny.Array(None, 8)
                nw294_[int(0)] = d_1430_cf20_
                nw294_[int(1)] = False
                nw294_[int(2)] = p2
                nw294_[int(3)] = p2
                nw294_[int(4)] = d_1430_cf20_
                nw294_[int(5)] = True
                nw294_[int(6)] = d_1428_cf22_
                nw294_[int(7)] = p2
                d_1466_v79_ = nw294_
                d_1467_v80_: _dafny.Array
                def lambda105_(d_1468_v76_):
                    def lambda106_(d_1469_i11_):
                        return d_1468_v76_

                    return lambda106_

                init59_ = lambda105_(d_1463_v76_)
                nw295_ = _dafny.Array(None, 14)
                for i0_59_ in range(nw295_.length(0)):
                    nw295_[i0_59_] = init59_(i0_59_)
                d_1467_v80_ = nw295_
                d_1470_v81_: D4
                d_1470_v81_ = D4_DC10(_dafny.Set({d_1466_v79_, d_1466_v79_}), len(d_1463_v76_), d_1422_v48_, d_1467_v80_)
                d_1430_cf20_ = (self).fm3(d_1431_v51_, (d_1470_v81_).cf12, globalState)
                d_1471_v82_: _dafny.Map
                d_1471_v82_ = _dafny.Map({d_1428_cf22_: d_1422_v48_})
                d_1472_v83_: _dafny.Map
                d_1472_v83_ = _dafny.Map({len(d_1423_v49_): _dafny.SeqWithoutIsStrInference([len(d_1471_v82_)])})
                d_1472_v83_ = (d_1472_v83_).set((p1) - (d_1431_v51_), d_1423_v49_)
                nw296_ = _dafny.Array(None, 5)
                nw296_[int(0)] = (0) - (default__.fm14(globalState))
                nw296_[int(1)] = len((d_1422_v48_) + (d_1422_v48_))
                nw296_[int(2)] = p1
                nw296_[int(3)] = p1
                nw296_[int(4)] = d_1431_v51_
                d_1464_v77_ = nw296_
            elif True:
                d_1473_v84_: C5
                nw297_ = C5()
                nw297_.ctor__()
                d_1473_v84_ = nw297_
                d_1474_v85_: _dafny.Array
                nw298_ = _dafny.Array(_dafny.Set({}), 22)
                d_1474_v85_ = nw298_
                d_1475_v86_: _dafny.Map
                d_1475_v86_ = _dafny.Map({p2: d_1429_cf21_})
                d_1476_v87_: T1
                nw299_ = C1()
                nw299_.ctor__(d_1474_v85_, default__.fm0(len(p0), len(d_1475_v86_), globalState))
                d_1476_v87_ = nw299_
                d_1476_v87_ = d_1476_v87_
                d_1421_v47_ = _dafny.CodePoint('v')
                d_1477_v88_: _dafny.MultiSet
                d_1477_v88_ = _dafny.MultiSet([p2])
                d_1478_v89_: _dafny.Array
                nw300_ = _dafny.Array(None, 7)
                nw300_[int(0)] = d_1428_cf22_
                nw300_[int(1)] = d_1428_cf22_
                nw300_[int(2)] = d_1429_cf21_
                nw300_[int(3)] = d_1428_cf22_
                nw300_[int(4)] = d_1429_cf21_
                nw300_[int(5)] = d_1430_cf20_
                nw300_[int(6)] = not (default__.fm1(-183, (d_1477_v88_).cardinality, not(d_1430_cf20_), d_1430_cf20_, globalState)) or (d_1430_cf20_)
                d_1478_v89_ = nw300_
                index299_ = default__.safeIndex(845, (d_1478_v89_).length(0))
                (d_1478_v89_)[index299_] = p2
                d_1479_v90_: _dafny.Map
                d_1479_v90_ = _dafny.Map({d_1478_v89_: (p0 if d_1428_cf22_ else p0)})
                d_1480_v91_: _dafny.Set
                d_1480_v91_ = _dafny.Set({d_1429_cf21_})
                index300_ = default__.safeIndex(845, (d_1478_v89_).length(0))
                rhs290_ = ((d_1475_v86_)[d_1428_cf22_] if (d_1428_cf22_) in (d_1475_v86_) else d_1430_cf20_)
                rhs291_ = (d_1422_v48_) + (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1481_i12_ in range(default__.abs(595))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1482_i12_ in range(default__.abs(595))]))), _dafny.CodePoint('g'))) + (d_1422_v48_))
                rhs292_ = len(d_1479_v90_)
                rhs293_ = (d_1480_v91_).issubset(_dafny.Set({d_1428_cf22_, True, d_1430_cf20_}))
                rhs294_ = d_1428_cf22_
                lhs153_ = d_1478_v89_
                lhs154_ = default__.safeIndex(845, (d_1478_v89_).length(0))
                d_1428_cf22_ = rhs290_
                d_1422_v48_ = rhs291_
                d_1431_v51_ = rhs292_
                lhs153_[lhs154_] = rhs293_
                d_1430_cf20_ = rhs294_
                d_1483_v92_: _dafny.Map
                d_1483_v92_ = _dafny.Map({(d_1478_v89_)[default__.safeIndex(845, (d_1478_v89_).length(0))]: p1})
                d_1483_v92_ = (d_1483_v92_).set((p1) != ((0) - (d_1431_v51_)), d_1431_v51_)
            d_1484_v93_: _dafny.Seq
            d_1484_v93_ = _dafny.SeqWithoutIsStrInference([d_1428_cf22_, d_1430_cf20_])
            d_1485_v94_: _dafny.Map
            d_1485_v94_ = _dafny.Map({p2: d_1484_v93_})
            d_1486_v95_: _dafny.Map
            d_1486_v95_ = _dafny.Map({d_1428_cf22_: len(d_1422_v48_)})
            d_1487_v96_: _dafny.MultiSet
            d_1487_v96_ = _dafny.MultiSet([d_1421_v47_, d_1421_v47_, _dafny.CodePoint('r')])
            d_1488_v97_: _dafny.MultiSet
            d_1488_v97_ = _dafny.MultiSet([False, p2, d_1430_cf20_, d_1429_cf21_])
            d_1489_v98_: _dafny.MultiSet
            d_1489_v98_ = _dafny.MultiSet([d_1431_v51_, p1, len((d_1422_v48_).set(default__.safeIndex(d_1431_v51_, len(d_1422_v48_)), d_1421_v47_)), d_1431_v51_])
            d_1490_v99_: _dafny.Array
            nw301_ = _dafny.Array(None, 26)
            nw301_[int(0)] = d_1431_v51_
            nw301_[int(1)] = p1
            nw301_[int(2)] = d_1431_v51_
            nw301_[int(3)] = len(((d_1485_v94_)[d_1430_cf20_] if (d_1430_cf20_) in (d_1485_v94_) else d_1484_v93_))
            nw301_[int(4)] = p1
            nw301_[int(5)] = p1
            nw301_[int(6)] = (d_1423_v49_)[default__.safeIndex(p1, len(d_1423_v49_))]
            nw301_[int(7)] = len(d_1486_v95_)
            nw301_[int(8)] = 741
            nw301_[int(9)] = 52
            nw301_[int(10)] = ((d_1487_v96_)[d_1421_v47_] if (d_1421_v47_) in (d_1487_v96_) else p1)
            nw301_[int(11)] = 813
            nw301_[int(12)] = d_1431_v51_
            nw301_[int(13)] = (d_1488_v97_).cardinality
            nw301_[int(14)] = len(d_1484_v93_)
            nw301_[int(15)] = (d_1489_v98_).cardinality
            nw301_[int(16)] = len(default__.fm33(d_1430_cf20_, not(d_1429_cf21_), (0) - (d_1431_v51_), globalState))
            nw301_[int(17)] = d_1431_v51_
            nw301_[int(18)] = len((d_1484_v93_).set(default__.safeIndex(d_1431_v51_, len(d_1484_v93_)), True))
            nw301_[int(19)] = p1
            nw301_[int(20)] = d_1431_v51_
            nw301_[int(21)] = p1
            nw301_[int(22)] = default__.fm14(globalState)
            nw301_[int(23)] = 823
            nw301_[int(24)] = len(p0)
            nw301_[int(25)] = len(d_1422_v48_)
            d_1490_v99_ = nw301_
            d_1491_v100_: _dafny.Set
            d_1491_v100_ = _dafny.Set({d_1490_v99_, d_1490_v99_})
            d_1429_cf21_ = not((d_1491_v100_).ispropersubset((d_1491_v100_) | (d_1491_v100_)))
        elif True:
            d_1492___mcc_h3_ = source14_.cf19
            d_1493_cf19_ = d_1492___mcc_h3_
            r0 = p2
            d_1494_v101_: _dafny.Array
            nw302_ = _dafny.Array(_dafny.Set({}), 28)
            d_1494_v101_ = nw302_
            index301_ = default__.safeIndex(176, (d_1494_v101_).length(0))
            (d_1494_v101_)[index301_] = _dafny.Set({p2})
            d_1495_v102_: _dafny.Seq
            d_1495_v102_ = _dafny.SeqWithoutIsStrInference([True])
            d_1496_v103_: _dafny.Set
            d_1496_v103_ = _dafny.Set({p2, ((self).fm3(-959, 610, globalState)) in (d_1495_v102_), p2})
            index302_ = default__.safeIndex(176, (d_1494_v101_).length(0))
            (d_1494_v101_)[index302_] = d_1496_v103_
            d_1497_v104_: D10
            d_1497_v104_ = D10_DC21(_dafny.Set({p2}))
            source15_ = d_1497_v104_
            if source15_.is_DC22:
                d_1498___mcc_h4_ = source15_.cf33
                d_1499___mcc_h5_ = source15_.cf34
                d_1500___mcc_h6_ = source15_.cf35
                d_1501_cf35_ = d_1500___mcc_h6_
                d_1502_cf34_ = d_1499___mcc_h5_
                d_1503_cf33_ = d_1498___mcc_h4_
                d_1504_v105_: D5
                d_1504_v105_ = D5_DC13(d_1503_cf33_, p1, p2)
                d_1505_v106_: _dafny.MultiSet
                d_1505_v106_ = _dafny.MultiSet([d_1501_cf35_])
                d_1506_v107_: _dafny.Set
                d_1506_v107_ = _dafny.Set({d_1505_v106_})
                pat_let_tv31_ = p2
                pat_let_tv32_ = d_1501_cf35_
                d_1507_v108_: _dafny.Array
                nw303_ = _dafny.Array(None, 23)
                nw303_[int(0)] = p2
                nw303_[int(1)] = p2
                nw303_[int(2)] = p2
                nw303_[int(3)] = (d_1501_cf35_) > (d_1501_cf35_)
                nw303_[int(4)] = not(p2)
                nw303_[int(5)] = p2
                nw303_[int(6)] = p2
                nw303_[int(7)] = False
                nw303_[int(8)] = p2
                nw303_[int(9)] = p2
                nw303_[int(10)] = p2
                nw303_[int(11)] = p2
                nw303_[int(12)] = p2
                nw303_[int(13)] = p2
                nw303_[int(14)] = not(not(default__.fm1(d_1501_cf35_, d_1501_cf35_, p2, not(p2), globalState)))
                def iife102_(_pat_let25_0):
                    def iife103_(d_1508_dt__update__tmp_h0_):
                        def iife104_(_pat_let26_0):
                            def iife105_(d_1509_dt__update_hcf18_h0_):
                                def iife106_(_pat_let27_0):
                                    def iife107_(d_1510_dt__update_hcf17_h0_):
                                        return D5_DC13((d_1508_dt__update__tmp_h0_).cf16, d_1510_dt__update_hcf17_h0_, d_1509_dt__update_hcf18_h0_)
                                    return iife107_(_pat_let27_0)
                                return iife106_((0) - (pat_let_tv32_))
                            return iife105_(_pat_let26_0)
                        return iife104_(pat_let_tv31_)
                    return iife103_(_pat_let25_0)
                nw303_[int(15)] = (p2) and ((iife102_(d_1504_v105_)).cf18)
                nw303_[int(16)] = p2
                nw303_[int(17)] = (d_1506_v107_) == (d_1506_v107_)
                nw303_[int(18)] = True
                nw303_[int(19)] = (d_1495_v102_) < (d_1495_v102_)
                nw303_[int(20)] = (p1) != (p1)
                nw303_[int(21)] = p2
                nw303_[int(22)] = p2
                d_1507_v108_ = nw303_
                index303_ = default__.safeIndex(915, (d_1507_v108_).length(0))
                (d_1507_v108_)[index303_] = p2
                index304_ = default__.safeIndex(915, (d_1507_v108_).length(0))
                (d_1507_v108_)[index304_] = p2
                d_1511_v109_: _dafny.Map
                d_1511_v109_ = _dafny.Map({not(p2): (d_1507_v108_)[default__.safeIndex(915, (d_1507_v108_).length(0))]})
                d_1512_v110_: _dafny.Map
                d_1512_v110_ = _dafny.Map({d_1421_v47_: d_1511_v109_})
                d_1513_v111_: _dafny.Seq
                d_1513_v111_ = _dafny.SeqWithoutIsStrInference([d_1512_v110_])
                d_1512_v110_ = ((d_1513_v111_)[default__.safeIndex(d_1501_cf35_, len(d_1513_v111_))]).set(_dafny.CodePoint('s'), (d_1511_v109_) | (d_1511_v109_))
                d_1514_v112_: C3
                nw304_ = C3()
                nw304_.ctor__(d_1494_v101_, _dafny.CodePoint('u'))
                d_1514_v112_ = nw304_
                d_1514_v112_ = d_1514_v112_
                d_1501_cf35_ = d_1501_cf35_
            elif source15_.is_DC21:
                d_1515___mcc_h7_ = source15_.cf32
                d_1516_cf32_ = d_1515___mcc_h7_
                d_1517_v113_: _dafny.MultiSet
                d_1517_v113_ = _dafny.MultiSet([_dafny.CodePoint('b'), d_1421_v47_])
                d_1518_v114_: _dafny.Map
                d_1518_v114_ = _dafny.Map({(_dafny.MultiSet([d_1421_v47_])).intersection(d_1517_v113_): p2})
                d_1518_v114_ = (d_1518_v114_).set(d_1517_v113_, p2)
                d_1519_v115_: C1
                nw305_ = C1()
                nw305_.ctor__(d_1494_v101_, d_1421_v47_)
                d_1519_v115_ = nw305_
                d_1520_v116_: _dafny.Map
                d_1520_v116_ = _dafny.Map({_dafny.CodePoint('g'): p2})
                r0 = not(((d_1519_v115_).fm5((0) - (p1), p2, p1, len(d_1520_v116_), globalState)) == (p2))
                d_1521_v117_: int
                d_1521_v117_ = 55
                d_1521_v117_ = p1
            elif True:
                d_1522___mcc_h8_ = source15_.cf36
                d_1523_cf36_ = d_1522___mcc_h8_
                d_1524_v118_: _dafny.Map
                d_1524_v118_ = _dafny.Map({p2: p1})
                d_1525_v119_: _dafny.Map
                d_1525_v119_ = _dafny.Map({d_1524_v118_: ((default__.fm19(p2, p1, (d_1495_v102_)[default__.safeIndex(p1, len(d_1495_v102_))], globalState))) | (_dafny.Map({p2: p2}))})
                d_1526_v120_: D17
                d_1526_v120_ = D17_DC40(d_1525_v119_)
                d_1525_v119_ = (d_1526_v120_).cf62
                d_1527_v121_: int
                d_1527_v121_ = 302
                d_1527_v121_ = d_1527_v121_
                d_1528_v122_: D1
                d_1528_v122_ = D1_DC1(_dafny.MultiSet([not(False)]))
                d_1529_v123_: _dafny.MultiSet
                d_1529_v123_ = _dafny.MultiSet([p2])
                d_1530_v124_: _dafny.Map
                d_1530_v124_ = _dafny.Map({(d_1529_v123_).cardinality: not(p2)})
                d_1531_v125_: _dafny.Map
                d_1531_v125_ = _dafny.Map({(0) - ((self).fm10(d_1528_v122_, ((d_1530_v124_)[p1] if (p1) in (d_1530_v124_) else p2), d_1530_v124_, globalState)): d_1527_v121_})
                d_1531_v125_ = (d_1531_v125_).set(d_1527_v121_, (self).fm10(d_1528_v122_, p2, default__.fm34(p2, globalState), globalState))
                d_1527_v121_ = d_1527_v121_
            d_1532_v126_: _dafny.Map
            d_1532_v126_ = _dafny.Map({_dafny.Set({p2, p2}): d_1496_v103_})
            d_1532_v126_ = d_1532_v126_
        d_1533_v127_: _dafny.Array
        def lambda107_(d_1534_i13_):
            return (d_1534_i13_) * (-729)

        init60_ = lambda107_
        nw306_ = _dafny.Array(None, 23)
        for i0_60_ in range(nw306_.length(0)):
            nw306_[i0_60_] = init60_(i0_60_)
        d_1533_v127_ = nw306_
        d_1535_v128_: _dafny.Seq
        d_1535_v128_ = _dafny.SeqWithoutIsStrInference([True])
        d_1536_v129_: D15
        d_1536_v129_ = D15_DC34(d_1535_v128_)
        d_1537_v130_: _dafny.Map
        d_1537_v130_ = _dafny.Map({p2: d_1533_v127_})
        d_1538_v131_: _dafny.Map
        d_1538_v131_ = _dafny.Map({d_1536_v129_: ((d_1537_v130_)[default__.fm1(-171, len(d_1535_v128_), p2, p2, globalState)] if (default__.fm1(-171, len(d_1535_v128_), p2, p2, globalState)) in (d_1537_v130_) else d_1533_v127_)})
        d_1539_v132_: _dafny.Array
        nw307_ = _dafny.Array(None, 3)
        nw307_[int(0)] = d_1533_v127_
        nw307_[int(1)] = ((d_1538_v131_)[d_1536_v129_] if (d_1536_v129_) in (d_1538_v131_) else d_1533_v127_)
        nw307_[int(2)] = d_1533_v127_
        d_1539_v132_ = nw307_
        index305_ = default__.safeIndex(851, (d_1539_v132_).length(0))
        (d_1539_v132_)[index305_] = d_1533_v127_
        index306_ = default__.safeIndex(851, (d_1539_v132_).length(0))
        (d_1539_v132_)[index306_] = d_1533_v127_
        r0 = ((d_1535_v128_ if p2 else d_1535_v128_)) < (d_1535_v128_)
        return r0


class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def m6(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: bool = False
        d_1540_v0_: _dafny.Array
        nw308_ = _dafny.Array(False, 25)
        d_1540_v0_ = nw308_
        d_1541_v1_: bool
        d_1541_v1_ = True
        index307_ = default__.safeIndex(220, (d_1540_v0_).length(0))
        (d_1540_v0_)[index307_] = d_1541_v1_
        d_1542_v2_: str
        d_1542_v2_ = _dafny.CodePoint('f')
        d_1543_v3_: _dafny.Seq
        d_1543_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vteybig"))
        index308_ = default__.safeIndex(220, (d_1540_v0_).length(0))
        (d_1540_v0_)[index308_] = (d_1542_v2_) not in (d_1543_v3_)
        d_1544_v4_: T0
        nw309_ = C7()
        nw309_.ctor__()
        d_1544_v4_ = nw309_
        d_1544_v4_ = d_1544_v4_
        if (p2) != ((187) * (p0)):
            d_1545_v5_: _dafny.Set
            d_1545_v5_ = _dafny.Set({False})
            d_1546_v6_: _dafny.Set
            d_1546_v6_ = _dafny.Set({d_1545_v5_, d_1545_v5_, (_dafny.Set({d_1541_v1_})).intersection(d_1545_v5_)})
            d_1547_v7_: D18
            d_1547_v7_ = D18_DC44(d_1546_v6_)
            d_1546_v6_ = (d_1547_v7_).cf67
            r2 = d_1541_v1_
            d_1548_v8_: int
            d_1548_v8_ = 970
            d_1549_v9_: _dafny.Set
            d_1549_v9_ = _dafny.Set({d_1543_v3_})
            d_1548_v8_ = len((d_1549_v9_) | (d_1549_v9_))
            d_1548_v8_ = (0) - ((0) - (p2))
            d_1550_v10_: _dafny.Seq
            d_1550_v10_ = _dafny.SeqWithoutIsStrInference([d_1548_v8_])
            d_1551_v11_: _dafny.Map
            d_1551_v11_ = _dafny.Map({p2: (d_1550_v10_) < (d_1550_v10_)})
            d_1551_v11_ = (d_1551_v11_).set(p0, True)
        elif True:
            d_1552_v12_: int
            d_1552_v12_ = 980
            d_1553_v13_: _dafny.Map
            d_1553_v13_ = _dafny.Map({d_1542_v2_: d_1552_v12_})
            d_1552_v12_ = ((len(d_1553_v13_)) - (p2)) - ((p3) * (p3))
            d_1554_v14_: _dafny.Map
            d_1554_v14_ = _dafny.Map({p3: p2})
            d_1555_v15_: C6
            nw310_ = C6()
            nw310_.ctor__((d_1554_v14_) | (d_1554_v14_))
            d_1555_v15_ = nw310_
            d_1556_v16_: C5
            nw311_ = C5()
            nw311_.ctor__()
            d_1556_v16_ = nw311_
            d_1557_v17_: _dafny.Set
            d_1557_v17_ = _dafny.Set({True})
            d_1558_v18_: _dafny.MultiSet
            d_1558_v18_ = _dafny.MultiSet([len(d_1557_v17_), p0, (0) - (p2), p0])
            d_1559_v19_: _dafny.Array
            nw312_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_1559_v19_ = nw312_
            d_1560_v20_: _dafny.Array
            def lambda108_(d_1561_v12_):
                def lambda109_(d_1562_i0_):
                    return default__.safeDivisionInt(d_1562_i0_, d_1561_v12_)

                return lambda109_

            init61_ = lambda108_(d_1552_v12_)
            nw313_ = _dafny.Array(None, 15)
            for i0_61_ in range(nw313_.length(0)):
                nw313_[i0_61_] = init61_(i0_61_)
            d_1560_v20_ = nw313_
            index309_ = default__.safeIndex(114, (d_1559_v19_).length(0))
            (d_1559_v19_)[index309_] = d_1560_v20_
            d_1563_v21_: _dafny.Array
            def lambda110_(d_1564_v3_):
                def lambda111_(d_1565_i1_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nut"))) + (d_1564_v3_)

                return lambda111_

            init62_ = lambda110_(d_1543_v3_)
            nw314_ = _dafny.Array(None, 9)
            for i0_62_ in range(nw314_.length(0)):
                nw314_[i0_62_] = init62_(i0_62_)
            d_1563_v21_ = nw314_
            index310_ = default__.safeIndex(869, (d_1563_v21_).length(0))
            (d_1563_v21_)[index310_] = d_1543_v3_
            d_1566_v22_: _dafny.Seq
            d_1566_v22_ = _dafny.SeqWithoutIsStrInference([(0) - (p3)])
            index311_ = default__.safeIndex(114, (d_1559_v19_).length(0))
            index312_ = default__.safeIndex(869, (d_1563_v21_).length(0))
            rhs295_ = p0
            rhs296_ = ((d_1558_v18_).set(p3, default__.abs(d_1552_v12_))) | (_dafny.MultiSet((d_1566_v22_ if d_1541_v1_ else d_1566_v22_)))
            rhs297_ = d_1560_v20_
            rhs298_ = (p2) * (p2)
            rhs299_ = d_1543_v3_
            lhs155_ = d_1559_v19_
            lhs156_ = default__.safeIndex(114, (d_1559_v19_).length(0))
            lhs157_ = d_1563_v21_
            lhs158_ = default__.safeIndex(869, (d_1563_v21_).length(0))
            d_1552_v12_ = rhs295_
            d_1558_v18_ = rhs296_
            lhs155_[lhs156_] = rhs297_
            d_1552_v12_ = rhs298_
            lhs157_[lhs158_] = rhs299_
            if (d_1543_v3_) != (d_1543_v3_):
                d_1567_v23_: C6
                nw315_ = C6()
                nw315_.ctor__((d_1555_v15_).f10)
                d_1567_v23_ = nw315_
                d_1568_v24_: _dafny.Array
                nw316_ = _dafny.Array(_dafny.Seq({}), 9)
                d_1568_v24_ = nw316_
                d_1568_v24_ = (d_1568_v24_ if (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))] else (d_1568_v24_ if (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))] else d_1568_v24_))
                d_1569_v25_: _dafny.Map
                d_1569_v25_ = _dafny.Map({(d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]: p3})
                d_1570_v26_: _dafny.Seq
                d_1570_v26_ = _dafny.SeqWithoutIsStrInference([d_1569_v25_])
                d_1571_v27_: _dafny.Seq
                d_1571_v27_ = _dafny.SeqWithoutIsStrInference([True])
                d_1554_v14_ = (d_1554_v14_).set((0) - (len(d_1570_v26_)), ((0) - (len(d_1571_v27_))) + (p0))
                index313_ = default__.safeIndex(220, (d_1540_v0_).length(0))
                (d_1540_v0_)[index313_] = (d_1544_v4_).fm3(p2, p3, globalState)
                d_1572_v28_: _dafny.Set
                d_1572_v28_ = _dafny.Set({d_1552_v12_})
                d_1572_v28_ = d_1572_v28_
            elif True:
                d_1573_v29_: _dafny.Array
                nw317_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_1573_v29_ = nw317_
                d_1574_v30_: D10
                d_1574_v30_ = D10_DC22(d_1543_v3_, d_1573_v29_, ((default__.fm23((d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], globalState)).set(d_1542_v2_, default__.abs(p0))).cardinality)
                d_1575_v31_: _dafny.Seq
                d_1575_v31_ = _dafny.SeqWithoutIsStrInference([d_1541_v1_, (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]])
                d_1576_v32_: _dafny.Seq
                d_1576_v32_ = _dafny.SeqWithoutIsStrInference([d_1563_v21_, (d_1563_v21_ if default__.fm1(686, p2, False, (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], globalState) else d_1563_v21_), d_1563_v21_, d_1563_v21_, d_1563_v21_])
                d_1577_v33_: _dafny.Map
                d_1577_v33_ = _dafny.Map({d_1541_v1_: d_1575_v31_})
                d_1578_v34_: _dafny.Map
                d_1578_v34_ = _dafny.Map({(((d_1555_v15_).f10)[len((_dafny.SeqWithoutIsStrInference([d_1542_v2_ for d_1579_i2_ in range(default__.abs(216))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_1542_v2_ for d_1580_i2_ in range(default__.abs(216))]))), _dafny.CodePoint('u')))] if (len((_dafny.SeqWithoutIsStrInference([d_1542_v2_ for d_1581_i2_ in range(default__.abs(216))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_1542_v2_ for d_1582_i2_ in range(default__.abs(216))]))), _dafny.CodePoint('u')))) in ((d_1555_v15_).f10) else len(d_1577_v33_)): d_1556_v16_})
                index314_ = default__.safeIndex(220, (d_1540_v0_).length(0))
                index315_ = default__.safeIndex(220, (d_1540_v0_).length(0))
                rhs300_ = (d_1575_v31_)[default__.safeIndex(p3, len(d_1575_v31_))]
                rhs301_ = (d_1576_v32_)[default__.safeIndex(p2, len(d_1576_v32_))]
                rhs302_ = (p2) >= (p2)
                rhs303_ = ((d_1578_v34_)[(p0) * (default__.fm14(globalState))] if ((p0) * (default__.fm14(globalState))) in (d_1578_v34_) else d_1556_v16_)
                rhs304_ = d_1574_v30_
                lhs159_ = d_1540_v0_
                lhs160_ = default__.safeIndex(220, (d_1540_v0_).length(0))
                lhs161_ = d_1540_v0_
                lhs162_ = default__.safeIndex(220, (d_1540_v0_).length(0))
                lhs159_[lhs160_] = rhs300_
                d_1563_v21_ = rhs301_
                lhs161_[lhs162_] = rhs302_
                d_1556_v16_ = rhs303_
                d_1574_v30_ = rhs304_
                d_1552_v12_ = len(default__.fm35(globalState))
                r1 = not((d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))])
                r3 = True
                d_1541_v1_ = d_1541_v1_
        index316_ = default__.safeIndex(220, (d_1540_v0_).length(0))
        (d_1540_v0_)[index316_] = not(d_1541_v1_)
        d_1583_v35_: _dafny.Set
        d_1583_v35_ = _dafny.Set({(d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]})
        d_1584_v36_: D10
        d_1584_v36_ = D10_DC21(d_1583_v35_)
        d_1585_v37_: _dafny.Seq
        d_1585_v37_ = _dafny.SeqWithoutIsStrInference([d_1583_v35_, d_1583_v35_])
        d_1586_v38_: D17
        d_1586_v38_ = D17_DC41(p3)
        pat_let_tv33_ = d_1583_v35_
        pat_let_tv34_ = d_1540_v0_
        pat_let_tv35_ = d_1540_v0_
        d_1587_v39_: _dafny.Array
        nw318_ = _dafny.Array(None, 24)
        nw318_[int(0)] = d_1584_v36_
        nw318_[int(1)] = d_1584_v36_
        nw318_[int(2)] = d_1584_v36_
        def iife108_(_pat_let28_0):
            def iife109_(d_1588_dt__update__tmp_h0_):
                def iife110_(_pat_let29_0):
                    def iife111_(d_1589_dt__update_hcf32_h0_):
                        return D10_DC21(d_1589_dt__update_hcf32_h0_)
                    return iife111_(_pat_let29_0)
                return iife110_(pat_let_tv33_)
            return iife109_(_pat_let28_0)
        nw318_[int(3)] = iife108_(d_1584_v36_)
        nw318_[int(4)] = D10_DC21((d_1585_v37_)[default__.safeIndex(p0, len(d_1585_v37_))])
        def iife112_(_pat_let30_0):
            def iife113_(d_1590_dt__update__tmp_h1_):
                def iife114_(_pat_let31_0):
                    def iife115_(d_1591_dt__update_hcf32_h1_):
                        return D10_DC21(d_1591_dt__update_hcf32_h1_)
                    return iife115_(_pat_let31_0)
                return iife114_(_dafny.Set({(pat_let_tv35_)[default__.safeIndex(220, (pat_let_tv34_).length(0))]}))
            return iife113_(_pat_let30_0)
        nw318_[int(5)] = iife112_(d_1584_v36_)
        nw318_[int(6)] = D10_DC21(d_1583_v35_)
        nw318_[int(7)] = d_1584_v36_
        nw318_[int(8)] = d_1584_v36_
        nw318_[int(9)] = (d_1584_v36_ if (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))] else d_1584_v36_)
        nw318_[int(10)] = d_1584_v36_
        nw318_[int(11)] = d_1584_v36_
        nw318_[int(12)] = d_1584_v36_
        nw318_[int(13)] = default__.fm36(d_1586_v38_, False, globalState)
        nw318_[int(14)] = d_1584_v36_
        nw318_[int(15)] = d_1584_v36_
        nw318_[int(16)] = d_1584_v36_
        nw318_[int(17)] = d_1584_v36_
        nw318_[int(18)] = d_1584_v36_
        nw318_[int(19)] = d_1584_v36_
        nw318_[int(20)] = d_1584_v36_
        nw318_[int(21)] = D10_DC21(d_1583_v35_)
        nw318_[int(22)] = d_1584_v36_
        nw318_[int(23)] = d_1584_v36_
        d_1587_v39_ = nw318_
        pat_let_tv36_ = d_1583_v35_
        pat_let_tv37_ = p3
        pat_let_tv38_ = p3
        pat_let_tv39_ = globalState
        index317_ = default__.safeIndex(286, (d_1587_v39_).length(0))
        def iife117_(_pat_let33_0):
            def iife118_(d_1592_dt__update__tmp_h2_):
                def iife119_(_pat_let34_0):
                    def iife120_(d_1593_dt__update_hcf32_h2_):
                        return D10_DC21(d_1593_dt__update_hcf32_h2_)
                    return iife120_(_pat_let34_0)
                return iife119_(pat_let_tv36_)
            return iife118_(_pat_let33_0)
        def iife116_(_pat_let32_0):
            def iife121_(d_1594_dt__update__tmp_h3_):
                def iife122_(_pat_let35_0):
                    def iife123_(d_1595_dt__update_hcf32_h3_):
                        return D10_DC21(d_1595_dt__update_hcf32_h3_)
                    return iife123_(_pat_let35_0)
                return iife122_(default__.fm26(pat_let_tv37_, pat_let_tv38_, pat_let_tv39_))
            return iife121_(_pat_let32_0)
        (d_1587_v39_)[index317_] = iife116_(iife117_(D10_DC21(_dafny.Set({d_1541_v1_, False}))))
        d_1596_v40_: _dafny.Array
        nw319_ = _dafny.Array(int(0), 16)
        d_1596_v40_ = nw319_
        d_1597_v41_: _dafny.Map
        d_1597_v41_ = _dafny.Map({d_1596_v40_: (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]})
        d_1598_v42_: _dafny.Seq
        d_1598_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojrqkow")), d_1543_v3_])
        d_1599_v43_: _dafny.Seq
        d_1599_v43_ = _dafny.SeqWithoutIsStrInference([len((d_1597_v41_).set(d_1596_v40_, (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))])), len(d_1598_v42_), -350, p2, 159])
        d_1600_v44_: _dafny.Array
        nw320_ = _dafny.Array(None, 1)
        nw320_[int(0)] = (d_1599_v43_) + (d_1599_v43_)
        d_1600_v44_ = nw320_
        index318_ = default__.safeIndex(6, (d_1600_v44_).length(0))
        (d_1600_v44_)[index318_] = d_1599_v43_
        index319_ = default__.safeIndex(286, (d_1587_v39_).length(0))
        index320_ = default__.safeIndex(6, (d_1600_v44_).length(0))
        rhs305_ = d_1584_v36_
        rhs306_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]: p0})) for d_1601_i3_ in range(default__.abs(-906))])
        lhs163_ = d_1587_v39_
        lhs164_ = default__.safeIndex(286, (d_1587_v39_).length(0))
        lhs165_ = d_1600_v44_
        lhs166_ = default__.safeIndex(6, (d_1600_v44_).length(0))
        lhs163_[lhs164_] = rhs305_
        lhs165_[lhs166_] = rhs306_
        if (d_1542_v2_) in (d_1543_v3_):
            index321_ = default__.safeIndex(875, (d_1596_v40_).length(0))
            (d_1596_v40_)[index321_] = p2
            index322_ = default__.safeIndex(875, (d_1596_v40_).length(0))
            (d_1596_v40_)[index322_] = 858
            r3 = not(d_1541_v1_)
            d_1602_v45_: _dafny.Array
            nw321_ = _dafny.Array(D6.default()(), 7)
            d_1602_v45_ = nw321_
            d_1603_v46_: _dafny.Seq
            d_1603_v46_ = _dafny.SeqWithoutIsStrInference([d_1541_v1_])
            d_1604_v47_: D6
            d_1604_v47_ = D6_DC15((d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], (d_1603_v46_)[default__.safeIndex(default__.fm13(True, d_1541_v1_, (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], globalState), len(d_1603_v46_))])
            index323_ = default__.safeIndex(805, (d_1602_v45_).length(0))
            (d_1602_v45_)[index323_] = d_1604_v47_
            index324_ = default__.safeIndex(805, (d_1602_v45_).length(0))
            (d_1602_v45_)[index324_] = d_1604_v47_
            index325_ = default__.safeIndex(875, (d_1596_v40_).length(0))
            (d_1596_v40_)[index325_] = (0) - (p2)
            if (d_1544_v4_).fm3((d_1596_v40_)[default__.safeIndex(875, (d_1596_v40_).length(0))], 124, globalState):
                d_1605_v48_: _dafny.MultiSet
                d_1605_v48_ = _dafny.MultiSet([False])
                index326_ = default__.safeIndex(875, (d_1596_v40_).length(0))
                (d_1596_v40_)[index326_] = (((0) - ((0) - (len(d_1543_v3_)))) + (p0)) - (((d_1605_v48_).cardinality) + ((d_1605_v48_).cardinality))
                d_1606_v49_: C5
                nw322_ = C5()
                nw322_.ctor__()
                d_1606_v49_ = nw322_
                d_1607_v50_: D16
                d_1607_v50_ = D16_DC38(d_1543_v3_, d_1541_v1_, p3, (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))])
                d_1608_v51_: _dafny.Array
                nw323_ = _dafny.Array(None, 17)
                nw323_[int(0)] = d_1543_v3_
                nw323_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilyxr"))
                nw323_[int(2)] = d_1543_v3_
                nw323_[int(3)] = d_1543_v3_
                nw323_[int(4)] = d_1543_v3_
                nw323_[int(5)] = ((d_1598_v42_)[default__.safeIndex(p0, len(d_1598_v42_))]) + (d_1543_v3_)
                nw323_[int(6)] = d_1543_v3_
                nw323_[int(7)] = d_1543_v3_
                nw323_[int(8)] = d_1543_v3_
                nw323_[int(9)] = d_1543_v3_
                nw323_[int(10)] = d_1543_v3_
                nw323_[int(11)] = d_1543_v3_
                nw323_[int(12)] = (d_1543_v3_).set(default__.safeIndex((_dafny.MultiSet([(d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], d_1541_v1_])).cardinality, len(d_1543_v3_)), d_1542_v2_)
                nw323_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkarnsgh"))
                nw323_[int(14)] = ((d_1607_v50_).cf57 if d_1541_v1_ else (d_1543_v3_).set(default__.safeIndex((0) - (p2), len(d_1543_v3_)), _dafny.CodePoint('y')))
                nw323_[int(15)] = (d_1543_v3_) + (d_1543_v3_)
                nw323_[int(16)] = (d_1543_v3_ if default__.fm1((d_1596_v40_)[default__.safeIndex(875, (d_1596_v40_).length(0))], p0, d_1541_v1_, d_1541_v1_, globalState) else d_1543_v3_)
                d_1608_v51_ = nw323_
                index327_ = default__.safeIndex(836, (d_1608_v51_).length(0))
                (d_1608_v51_)[index327_] = d_1543_v3_
                index328_ = default__.safeIndex(836, (d_1608_v51_).length(0))
                (d_1608_v51_)[index328_] = d_1543_v3_
                index329_ = default__.safeIndex(220, (d_1540_v0_).length(0))
                (d_1540_v0_)[index329_] = d_1541_v1_
                d_1609_v52_: _dafny.Map
                d_1609_v52_ = _dafny.Map({default__.safeModuloInt(p0, len((d_1608_v51_)[default__.safeIndex(836, (d_1608_v51_).length(0))])): d_1541_v1_})
                d_1609_v52_ = (d_1609_v52_).set(433, not (not(d_1541_v1_)) or (d_1541_v1_))
            elif True:
                r3 = d_1541_v1_
                index330_ = default__.safeIndex(875, (d_1596_v40_).length(0))
                (d_1596_v40_)[index330_] = len((_dafny.SeqWithoutIsStrInference([p0 for d_1610_i4_ in range(default__.abs(938))])).set(default__.safeIndex((d_1596_v40_)[default__.safeIndex(875, (d_1596_v40_).length(0))], len(_dafny.SeqWithoutIsStrInference([p0 for d_1611_i4_ in range(default__.abs(938))]))), p2))
                d_1612_v54_: _dafny.Map
                d_1612_v54_ = _dafny.Map({d_1603_v46_: d_1541_v1_})
                d_1613_v55_: _dafny.Array
                nw324_ = _dafny.Array(None, 28)
                nw324_[int(0)] = d_1585_v37_
                nw324_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1583_v35_])
                nw324_[int(2)] = d_1585_v37_
                nw324_[int(3)] = d_1585_v37_
                nw324_[int(4)] = d_1585_v37_
                nw324_[int(5)] = default__.fm37(d_1598_v42_, globalState)
                nw324_[int(6)] = d_1585_v37_
                nw324_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1583_v35_ for d_1614_i5_ in range(default__.abs(320))])
                nw324_[int(8)] = (d_1585_v37_) + (d_1585_v37_)
                nw324_[int(9)] = (d_1585_v37_) + (d_1585_v37_)
                nw324_[int(10)] = (d_1585_v37_ if not(False) else d_1585_v37_)
                nw324_[int(11)] = (d_1585_v37_) + (_dafny.SeqWithoutIsStrInference([d_1583_v35_]))
                nw324_[int(12)] = default__.fm37(d_1598_v42_, globalState)
                nw324_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1583_v35_])
                nw324_[int(14)] = (d_1585_v37_ if (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))] else d_1585_v37_)
                def iife124_():
                    coll52_ = _dafny.Map()
                    compr_52_: _dafny.Seq
                    for compr_52_ in (d_1612_v54_).keys.Elements:
                        d_1617_v53_: _dafny.Seq = compr_52_
                        if (d_1617_v53_) in (d_1612_v54_):
                            coll52_[d_1617_v53_] = d_1541_v1_
                    return _dafny.Map(coll52_)
                nw324_[int(15)] = (_dafny.SeqWithoutIsStrInference([d_1583_v35_ for d_1615_i6_ in range(default__.abs(-724))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_1583_v35_ for d_1616_i6_ in range(default__.abs(-724))]))), default__.fm26((0) - (p2), len(iife124_()
                ), globalState))
                nw324_[int(16)] = d_1585_v37_
                nw324_[int(17)] = d_1585_v37_
                nw324_[int(18)] = _dafny.SeqWithoutIsStrInference([d_1583_v35_])
                nw324_[int(19)] = (d_1585_v37_) + (d_1585_v37_)
                nw324_[int(20)] = d_1585_v37_
                nw324_[int(21)] = (default__.fm37(_dafny.SeqWithoutIsStrInference([d_1543_v3_ for d_1618_i7_ in range(default__.abs(197))]), globalState)).set(default__.safeIndex(len(d_1543_v3_), len(default__.fm37(_dafny.SeqWithoutIsStrInference([d_1543_v3_ for d_1619_i7_ in range(default__.abs(197))]), globalState))), d_1583_v35_)
                nw324_[int(22)] = d_1585_v37_
                nw324_[int(23)] = (d_1585_v37_) + (d_1585_v37_)
                nw324_[int(24)] = (_dafny.SeqWithoutIsStrInference([_dafny.Set({d_1541_v1_, (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], not(default__.fm1(p2, (d_1596_v40_)[default__.safeIndex(875, (d_1596_v40_).length(0))], (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))], d_1541_v1_, globalState))})])) + (d_1585_v37_)
                nw324_[int(25)] = d_1585_v37_
                nw324_[int(26)] = (d_1585_v37_) + (d_1585_v37_)
                nw324_[int(27)] = (_dafny.SeqWithoutIsStrInference([d_1583_v35_ for d_1620_i8_ in range(default__.abs(458))])) + (d_1585_v37_)
                d_1613_v55_ = nw324_
                index331_ = default__.safeIndex(84, (d_1613_v55_).length(0))
                (d_1613_v55_)[index331_] = d_1585_v37_
                d_1621_v56_: _dafny.Map
                d_1621_v56_ = _dafny.Map({d_1603_v46_: _dafny.SeqWithoutIsStrInference([d_1583_v35_ for d_1622_i9_ in range(default__.abs(-262))])})
                index332_ = default__.safeIndex(84, (d_1613_v55_).length(0))
                (d_1613_v55_)[index332_] = ((d_1621_v56_)[d_1603_v46_] if (d_1603_v46_) in (d_1621_v56_) else d_1585_v37_)
                d_1623_v57_: _dafny.Map
                d_1623_v57_ = _dafny.Map({p2: d_1541_v1_})
                d_1624_v58_: _dafny.MultiSet
                d_1624_v58_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktvdbuymx"))) + (_dafny.SeqWithoutIsStrInference([d_1542_v2_ for d_1625_i10_ in range(default__.abs(890))])), (d_1543_v3_) + (default__.fm16(d_1541_v1_, (_dafny.MultiSet([d_1541_v1_])).cardinality, globalState)), d_1543_v3_, (d_1543_v3_) + (default__.fm16(d_1541_v1_, len(d_1623_v57_), globalState))])
                d_1624_v58_ = (_dafny.MultiSet([d_1543_v3_, d_1543_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")), d_1543_v3_, d_1543_v3_])).intersection(d_1624_v58_)
                d_1626_v59_: _dafny.Map
                d_1626_v59_ = _dafny.Map({(d_1543_v3_) + (d_1543_v3_): d_1541_v1_})
                d_1626_v59_ = (d_1626_v59_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yukmtwjs")), d_1541_v1_)
        elif True:
            d_1543_v3_ = d_1543_v3_
            d_1627_v60_: D15
            d_1627_v60_ = D15_DC35(p3)
            d_1628_v61_: D15
            d_1628_v61_ = D15_DC36(d_1627_v60_)
            d_1629_v62_: _dafny.Map
            d_1629_v62_ = _dafny.Map({(d_1628_v61_ if d_1541_v1_ else d_1628_v61_): (d_1543_v3_) + (d_1543_v3_)})
            d_1629_v62_ = (d_1629_v62_).set(d_1628_v61_, d_1543_v3_)
            d_1630_v63_: int
            d_1630_v63_ = -609
            d_1630_v63_ = p3
            d_1630_v63_ = ((p3) * (p3)) + (d_1630_v63_)
            d_1631_v64_: _dafny.Map
            d_1631_v64_ = _dafny.Map({p2: d_1541_v1_})
            index333_ = default__.safeIndex(220, (d_1540_v0_).length(0))
            (d_1540_v0_)[index333_] = (612) in (d_1631_v64_)
        d_1632_v65_: _dafny.Map
        d_1632_v65_ = _dafny.Map({True: d_1542_v2_})
        d_1633_v66_: _dafny.Map
        d_1633_v66_ = _dafny.Map({d_1541_v1_: D5_DC11(d_1632_v65_)})
        r0 = (d_1633_v66_ if ((d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))] if (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))] else (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]) else d_1633_v66_)
        r1 = d_1541_v1_
        r2 = (d_1540_v0_)[default__.safeIndex(220, (d_1540_v0_).length(0))]
        d_1634_v67_: D2
        d_1634_v67_ = D2_DC4(d_1541_v1_)
        r3 = (d_1634_v67_).cf5
        return r0, r1, r2, r3

    def m7(self, globalState):
        r0: D2 = D2.default()()
        d_1635_v0_: bool
        d_1635_v0_ = True
        d_1636_v1_: _dafny.Map
        d_1636_v1_ = _dafny.Map({False: not(d_1635_v0_)})
        d_1637_v2_: int
        d_1637_v2_ = 830
        hi14_ = d_1637_v2_
        for d_1638_i0_ in range(default__.fm13(d_1635_v0_, d_1635_v0_, ((d_1636_v1_)[d_1635_v0_] if (d_1635_v0_) in (d_1636_v1_) else False), globalState), hi14_):
            d_1639_v3_: _dafny.Seq
            d_1639_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uri"))
            d_1640_v4_: _dafny.Array
            nw325_ = _dafny.Array(None, 6)
            nw325_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1641_i1_ in range(default__.abs(-647))])
            nw325_[int(1)] = d_1639_v3_
            nw325_[int(2)] = (d_1639_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcvosk")))
            nw325_[int(3)] = d_1639_v3_
            nw325_[int(4)] = d_1639_v3_
            nw325_[int(5)] = (d_1639_v3_) + (d_1639_v3_)
            d_1640_v4_ = nw325_
            d_1640_v4_ = d_1640_v4_
            d_1642_v5_: D15
            d_1642_v5_ = D15_DC34(_dafny.SeqWithoutIsStrInference([not(d_1635_v0_), d_1635_v0_, d_1635_v0_, default__.fm1(d_1637_v2_, d_1637_v2_, d_1635_v0_, d_1635_v0_, globalState)]))
            d_1643_v6_: _dafny.Seq
            d_1643_v6_ = _dafny.SeqWithoutIsStrInference([True, d_1635_v0_, d_1635_v0_, d_1635_v0_, not(d_1635_v0_)])
            pat_let_tv40_ = d_1643_v6_
            def iife125_(_pat_let36_0):
                def iife126_(d_1644_dt__update__tmp_h0_):
                    def iife127_(_pat_let37_0):
                        def iife128_(d_1645_dt__update_hcf53_h0_):
                            return D15_DC34(d_1645_dt__update_hcf53_h0_)
                        return iife128_(_pat_let37_0)
                    return iife127_(pat_let_tv40_)
                return iife126_(_pat_let36_0)
            source16_ = iife125_(d_1642_v5_)
            if source16_.is_DC35:
                d_1646___mcc_h0_ = source16_.cf54
                d_1647_cf54_ = d_1646___mcc_h0_
                d_1648_v7_: str
                d_1648_v7_ = _dafny.CodePoint('f')
                d_1649_v8_: _dafny.Set
                d_1649_v8_ = _dafny.Set({(_dafny.CodePoint('m')) in (d_1639_v3_), d_1635_v0_, d_1635_v0_})
                rhs307_ = default__.fm24((0) - (((0) - (d_1638_i0_)) * (d_1638_i0_)), globalState)
                rhs308_ = d_1648_v7_
                rhs309_ = d_1649_v8_
                d_1639_v3_ = rhs307_
                d_1648_v7_ = rhs308_
                d_1649_v8_ = rhs309_
                d_1650_v9_: _dafny.Seq
                d_1650_v9_ = _dafny.SeqWithoutIsStrInference([default__.fm13(d_1635_v0_, d_1635_v0_, not(d_1635_v0_), globalState)])
                def iife129_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in (d_1650_v9_).Elements:
                        d_1651_v10_: int = compr_53_
                        if (d_1651_v10_) in (d_1650_v9_):
                            coll53_ = coll53_.union(_dafny.Set([(d_1651_v10_) - ((_dafny.MultiSet([668])).cardinality)]))
                    return _dafny.Set(coll53_)
                def iife130_():
                    coll54_ = _dafny.Set()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(281, -758):
                        d_1652_v11_: int = compr_54_
                        if ((281) <= (d_1652_v11_)) and ((d_1652_v11_) < (-758)):
                            coll54_ = coll54_.union(_dafny.Set([(d_1652_v11_) - (-561)]))
                    return _dafny.Set(coll54_)
                d_1635_v0_ = (d_1635_v0_ if (iife129_()
                ) == (iife130_()
                ) else d_1635_v0_)
                d_1653_v12_: _dafny.Array
                nw326_ = _dafny.Array(_dafny.Set({}), 3)
                d_1653_v12_ = nw326_
                d_1654_v13_: C3
                nw327_ = C3()
                nw327_.ctor__(d_1653_v12_, (d_1639_v3_)[default__.safeIndex(d_1638_i0_, len(d_1639_v3_))])
                d_1654_v13_ = nw327_
                d_1635_v0_ = ((d_1636_v1_)[True] if (True) in (d_1636_v1_) else d_1635_v0_)
            elif source16_.is_DC34:
                d_1655___mcc_h1_ = source16_.cf53
                d_1656_cf53_ = d_1655___mcc_h1_
                d_1657_v14_: _dafny.Array
                def lambda112_(d_1658_v0_):
                    def lambda113_(d_1659_i2_):
                        return d_1658_v0_

                    return lambda113_

                init63_ = lambda112_(d_1635_v0_)
                nw328_ = _dafny.Array(None, 9)
                for i0_63_ in range(nw328_.length(0)):
                    nw328_[i0_63_] = init63_(i0_63_)
                d_1657_v14_ = nw328_
                d_1660_v15_: _dafny.MultiSet
                d_1660_v15_ = _dafny.MultiSet([d_1657_v14_, d_1657_v14_])
                d_1661_v16_: D2
                d_1661_v16_ = D2_DC3(d_1657_v14_)
                pat_let_tv41_ = d_1657_v14_
                d_1662_v17_: D2
                def iife131_(_pat_let38_0):
                    def iife132_(d_1663_dt__update__tmp_h1_):
                        def iife133_(_pat_let39_0):
                            def iife134_(d_1664_dt__update_hcf4_h0_):
                                return D2_DC3(d_1664_dt__update_hcf4_h0_)
                            return iife134_(_pat_let39_0)
                        return iife133_(pat_let_tv41_)
                    return iife132_(_pat_let38_0)
                d_1662_v17_ = D2_DC3((iife131_(d_1661_v16_)).cf4)
                d_1665_v18_: _dafny.Seq
                d_1665_v18_ = _dafny.SeqWithoutIsStrInference([d_1661_v16_])
                d_1666_v19_: _dafny.Set
                d_1666_v19_ = _dafny.Set({d_1635_v0_})
                d_1667_v20_: _dafny.Map
                d_1667_v20_ = _dafny.Map({default__.fm31(d_1637_v2_, len(d_1666_v19_), globalState): _dafny.CodePoint('c')})
                pat_let_tv42_ = d_1657_v14_
                def iife135_(_pat_let40_0):
                    def iife136_(d_1668_dt__update__tmp_h2_):
                        def iife137_(_pat_let41_0):
                            def iife138_(d_1669_dt__update_hcf4_h1_):
                                return D2_DC3(d_1669_dt__update_hcf4_h1_)
                            return iife138_(_pat_let41_0)
                        return iife137_(pat_let_tv42_)
                    return iife136_(_pat_let40_0)
                rhs310_ = (d_1638_i0_) == (d_1638_i0_)
                rhs311_ = ((d_1636_v1_)[d_1635_v0_] if (d_1635_v0_) in (d_1636_v1_) else (d_1660_v15_).isdisjoint(d_1660_v15_))
                rhs312_ = default__.fm1(len(((d_1665_v18_).set(default__.safeIndex(d_1637_v2_, len(d_1665_v18_)), d_1662_v17_)) + (_dafny.SeqWithoutIsStrInference([d_1661_v16_, iife135_(d_1662_v17_), d_1662_v17_]))), len(d_1667_v20_), d_1635_v0_, not (d_1635_v0_) or (d_1635_v0_), globalState)
                d_1635_v0_ = rhs310_
                d_1635_v0_ = rhs311_
                d_1635_v0_ = rhs312_
                index334_ = default__.safeIndex(936, (d_1657_v14_).length(0))
                (d_1657_v14_)[index334_] = False
                index335_ = default__.safeIndex(936, (d_1657_v14_).length(0))
                (d_1657_v14_)[index335_] = d_1635_v0_
                d_1670_v21_: C0
                nw329_ = C0()
                nw329_.ctor__()
                d_1670_v21_ = nw329_
                d_1671_v22_: _dafny.Map
                d_1671_v22_ = _dafny.Map({(d_1657_v14_)[default__.safeIndex(936, (d_1657_v14_).length(0))]: d_1637_v2_})
                d_1672_v24_: _dafny.Seq
                d_1672_v24_ = _dafny.SeqWithoutIsStrInference([d_1637_v2_, d_1638_i0_])
                d_1673_v25_: _dafny.Map
                def iife139_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in _dafny.IntegerRange(-212, 277):
                        d_1674_v23_: int = compr_55_
                        if ((-212) <= (d_1674_v23_)) and ((d_1674_v23_) < (277)):
                            coll55_[default__.safeModuloInt(d_1674_v23_, d_1638_i0_)] = d_1638_i0_
                    return _dafny.Map(coll55_)
                d_1673_v25_ = _dafny.Map({iife139_()
                : (d_1672_v24_)[default__.safeIndex(d_1637_v2_, len(d_1672_v24_))]})
                d_1675_v26_: _dafny.Map
                d_1675_v26_ = _dafny.Map({d_1638_i0_: 76})
                d_1676_v27_: D10
                d_1676_v27_ = D10_DC21(default__.fm26(((d_1671_v22_)[True] if (True) in (d_1671_v22_) else d_1638_i0_), ((d_1673_v25_)[d_1675_v26_] if (d_1675_v26_) in (d_1673_v25_) else d_1638_i0_), globalState))
                d_1666_v19_ = ((d_1666_v19_).intersection(d_1666_v19_)) - (((d_1676_v27_).cf32) | (d_1666_v19_))
            elif True:
                d_1677___mcc_h2_ = source16_.cf55
                d_1678_cf55_ = d_1677___mcc_h2_
                d_1679_v28_: _dafny.Array
                nw330_ = _dafny.Array(False, 17)
                d_1679_v28_ = nw330_
                index336_ = default__.safeIndex(266, (d_1679_v28_).length(0))
                (d_1679_v28_)[index336_] = d_1635_v0_
                index337_ = default__.safeIndex(266, (d_1679_v28_).length(0))
                (d_1679_v28_)[index337_] = d_1635_v0_
                index338_ = default__.safeIndex(266, (d_1679_v28_).length(0))
                (d_1679_v28_)[index338_] = d_1635_v0_
                d_1680_v29_: _dafny.Map
                d_1680_v29_ = _dafny.Map({default__.fm1(d_1637_v2_, d_1637_v2_, d_1635_v0_, d_1635_v0_, globalState): d_1638_i0_})
                d_1680_v29_ = (d_1680_v29_).set(d_1635_v0_, (d_1638_i0_) + (len(d_1643_v6_)))
                d_1681_v30_: str
                d_1681_v30_ = _dafny.CodePoint('l')
                d_1681_v30_ = d_1681_v30_
            d_1636_v1_ = (d_1636_v1_) | ((d_1636_v1_).set(True, d_1635_v0_))
            d_1682_v31_: _dafny.MultiSet
            d_1682_v31_ = _dafny.MultiSet([d_1637_v2_])
            d_1683_v32_: _dafny.Map
            d_1683_v32_ = _dafny.Map({861: (0) - (d_1637_v2_)})
            d_1684_v33_: _dafny.Seq
            d_1684_v33_ = _dafny.SeqWithoutIsStrInference([len(d_1643_v6_), d_1637_v2_, (0) - (len(d_1683_v32_))])
            d_1685_v34_: _dafny.Seq
            d_1685_v34_ = _dafny.SeqWithoutIsStrInference([(d_1684_v33_)[default__.safeIndex(d_1637_v2_, len(d_1684_v33_))], (0) - (d_1638_i0_), d_1638_i0_])
            d_1686_v35_: _dafny.Seq
            d_1686_v35_ = _dafny.SeqWithoutIsStrInference([d_1637_v2_, len(d_1636_v1_), (d_1682_v31_).cardinality, len(d_1684_v33_), -724])
            d_1687_v36_: _dafny.Array
            nw331_ = _dafny.Array(None, 28)
            nw331_[int(0)] = d_1635_v0_
            nw331_[int(1)] = d_1635_v0_
            nw331_[int(2)] = d_1635_v0_
            nw331_[int(3)] = d_1635_v0_
            nw331_[int(4)] = d_1635_v0_
            nw331_[int(5)] = d_1635_v0_
            nw331_[int(6)] = d_1635_v0_
            nw331_[int(7)] = not(d_1635_v0_)
            nw331_[int(8)] = d_1635_v0_
            nw331_[int(9)] = d_1635_v0_
            nw331_[int(10)] = d_1635_v0_
            nw331_[int(11)] = d_1635_v0_
            nw331_[int(12)] = not(d_1635_v0_)
            nw331_[int(13)] = d_1635_v0_
            nw331_[int(14)] = d_1635_v0_
            nw331_[int(15)] = d_1635_v0_
            nw331_[int(16)] = d_1635_v0_
            nw331_[int(17)] = d_1635_v0_
            nw331_[int(18)] = False
            nw331_[int(19)] = False
            nw331_[int(20)] = d_1635_v0_
            nw331_[int(21)] = d_1635_v0_
            nw331_[int(22)] = d_1635_v0_
            nw331_[int(23)] = d_1635_v0_
            nw331_[int(24)] = d_1635_v0_
            nw331_[int(25)] = d_1635_v0_
            nw331_[int(26)] = d_1635_v0_
            nw331_[int(27)] = d_1635_v0_
            d_1687_v36_ = nw331_
            d_1688_v37_: _dafny.MultiSet
            d_1688_v37_ = _dafny.MultiSet([d_1687_v36_, d_1687_v36_, d_1687_v36_])
            d_1689_v38_: _dafny.Array
            nw332_ = _dafny.Array(None, 19)
            nw332_[int(0)] = (d_1686_v35_).set(default__.safeIndex(d_1638_i0_, len(d_1686_v35_)), d_1638_i0_)
            nw332_[int(1)] = d_1684_v33_
            nw332_[int(2)] = (d_1685_v34_) + (_dafny.SeqWithoutIsStrInference([d_1637_v2_, (0) - (d_1638_i0_)]))
            nw332_[int(3)] = d_1686_v35_
            nw332_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1638_i0_ for d_1690_i3_ in range(default__.abs(762))])
            nw332_[int(5)] = d_1684_v33_
            nw332_[int(6)] = d_1686_v35_
            nw332_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1638_i0_])
            nw332_[int(8)] = d_1685_v34_
            nw332_[int(9)] = (d_1684_v33_) + (d_1684_v33_)
            nw332_[int(10)] = d_1685_v34_
            nw332_[int(11)] = d_1684_v33_
            nw332_[int(12)] = d_1685_v34_
            nw332_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1638_i0_, d_1638_i0_])
            nw332_[int(14)] = d_1686_v35_
            nw332_[int(15)] = ((d_1684_v33_) + (d_1685_v34_)).set(default__.safeIndex(d_1637_v2_, len((d_1684_v33_) + (d_1685_v34_))), (d_1688_v37_).cardinality)
            nw332_[int(16)] = _dafny.SeqWithoutIsStrInference([len(d_1643_v6_), 188, d_1638_i0_])
            nw332_[int(17)] = d_1685_v34_
            nw332_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_1638_i0_ for d_1691_i4_ in range(default__.abs(-432))])) + (d_1686_v35_)
            d_1689_v38_ = nw332_
            index339_ = default__.safeIndex(81, (d_1689_v38_).length(0))
            (d_1689_v38_)[index339_] = _dafny.SeqWithoutIsStrInference([d_1638_i0_, d_1637_v2_, d_1637_v2_])
            d_1692_v39_: _dafny.Seq
            d_1692_v39_ = _dafny.SeqWithoutIsStrInference([d_1684_v33_, d_1686_v35_, d_1686_v35_])
            d_1693_v40_: _dafny.Array
            def lambda114_(d_1694_i5_):
                return _dafny.MultiSet([_dafny.CodePoint('d')])

            init64_ = lambda114_
            nw333_ = _dafny.Array(None, 4)
            for i0_64_ in range(nw333_.length(0)):
                nw333_[i0_64_] = init64_(i0_64_)
            d_1693_v40_ = nw333_
            d_1695_v41_: _dafny.MultiSet
            d_1695_v41_ = _dafny.MultiSet([d_1693_v40_])
            index340_ = default__.safeIndex(81, (d_1689_v38_).length(0))
            (d_1689_v38_)[index340_] = (d_1692_v39_)[default__.safeIndex(((d_1695_v41_)[d_1693_v40_] if (d_1693_v40_) in (d_1695_v41_) else d_1637_v2_), len(d_1692_v39_))]
        d_1696_v42_: str
        d_1696_v42_ = _dafny.CodePoint('v')
        d_1697_v43_: _dafny.MultiSet
        d_1697_v43_ = _dafny.MultiSet([d_1696_v42_])
        hi15_ = (d_1697_v43_).cardinality
        for d_1698_i6_ in range(d_1637_v2_, hi15_):
            d_1699_v44_: _dafny.Seq
            d_1699_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxncedkul"))
            d_1635_v0_ = (d_1699_v44_) <= ((d_1699_v44_) + (d_1699_v44_))
            d_1700_v45_: C7
            nw334_ = C7()
            nw334_.ctor__()
            d_1700_v45_ = nw334_
            d_1701_v46_: _dafny.Array
            nw335_ = _dafny.Array(_dafny.Seq({}), 14)
            d_1701_v46_ = nw335_
            d_1702_v47_: _dafny.Seq
            d_1702_v47_ = _dafny.SeqWithoutIsStrInference([d_1635_v0_])
            d_1703_v48_: _dafny.MultiSet
            d_1703_v48_ = _dafny.MultiSet([d_1702_v47_, _dafny.SeqWithoutIsStrInference([d_1635_v0_])])
            d_1704_v49_: _dafny.Seq
            d_1704_v49_ = _dafny.SeqWithoutIsStrInference([(d_1703_v48_).cardinality, d_1698_i6_])
            index341_ = default__.safeIndex(498, (d_1701_v46_).length(0))
            (d_1701_v46_)[index341_] = d_1704_v49_
            index342_ = default__.safeIndex(498, (d_1701_v46_).length(0))
            (d_1701_v46_)[index342_] = _dafny.SeqWithoutIsStrInference([default__.fm13(d_1635_v0_, d_1635_v0_, d_1635_v0_, globalState)])
            d_1705_v50_: _dafny.Array
            nw336_ = _dafny.Array(_dafny.Set({}), 22)
            d_1705_v50_ = nw336_
            d_1706_v51_: C2
            nw337_ = C2()
            nw337_.ctor__(d_1705_v50_, _dafny.CodePoint('o'))
            d_1706_v51_ = nw337_
            d_1707_v52_: _dafny.Seq
            d_1707_v52_ = _dafny.SeqWithoutIsStrInference([d_1706_v51_, d_1706_v51_])
            d_1635_v0_ = (d_1707_v52_) < (d_1707_v52_)
        d_1708_v53_: _dafny.Set
        d_1708_v53_ = _dafny.Set({d_1637_v2_, d_1637_v2_, d_1637_v2_})
        d_1709_v54_: _dafny.Map
        d_1709_v54_ = _dafny.Map({d_1635_v0_: len(d_1708_v53_)})
        d_1710_v55_: _dafny.Seq
        d_1710_v55_ = _dafny.SeqWithoutIsStrInference([len(d_1709_v54_), len(d_1636_v1_), d_1637_v2_])
        d_1711_v56_: C0
        nw338_ = C0()
        nw338_.ctor__()
        d_1711_v56_ = nw338_
        d_1712_v57_: _dafny.Map
        d_1712_v57_ = _dafny.Map({-64: d_1635_v0_})
        d_1713_v58_: _dafny.MultiSet
        d_1713_v58_ = _dafny.MultiSet([(d_1710_v55_)[default__.safeIndex(len(d_1712_v57_), len(d_1710_v55_))], len((d_1710_v55_).set(default__.safeIndex((0) - (d_1637_v2_), len(d_1710_v55_)), d_1637_v2_)), len(d_1710_v55_), d_1637_v2_])
        d_1714_v59_: _dafny.Array
        nw339_ = _dafny.Array(None, 23)
        nw339_[int(0)] = d_1635_v0_
        nw339_[int(1)] = default__.fm1(len(d_1710_v55_), len(_dafny.Set({d_1711_v56_, d_1711_v56_})), d_1635_v0_, d_1635_v0_, globalState)
        nw339_[int(2)] = not(d_1635_v0_)
        nw339_[int(3)] = d_1635_v0_
        nw339_[int(4)] = (d_1711_v56_).fm18(d_1635_v0_, d_1637_v2_, d_1635_v0_, (0) - ((d_1713_v58_).cardinality), globalState)
        nw339_[int(5)] = not(d_1635_v0_)
        nw339_[int(6)] = d_1635_v0_
        nw339_[int(7)] = True
        nw339_[int(8)] = d_1635_v0_
        nw339_[int(9)] = d_1635_v0_
        nw339_[int(10)] = not(d_1635_v0_)
        nw339_[int(11)] = d_1635_v0_
        nw339_[int(12)] = d_1635_v0_
        nw339_[int(13)] = d_1635_v0_
        nw339_[int(14)] = False
        nw339_[int(15)] = not((d_1711_v56_).fm18(d_1635_v0_, d_1637_v2_, d_1635_v0_, d_1637_v2_, globalState))
        nw339_[int(16)] = d_1635_v0_
        nw339_[int(17)] = d_1635_v0_
        nw339_[int(18)] = not(d_1635_v0_)
        nw339_[int(19)] = d_1635_v0_
        nw339_[int(20)] = d_1635_v0_
        nw339_[int(21)] = d_1635_v0_
        nw339_[int(22)] = d_1635_v0_
        d_1714_v59_ = nw339_
        d_1715_v60_: D2
        d_1715_v60_ = D2_DC3(d_1714_v59_)
        source17_ = d_1715_v60_
        if source17_.is_DC4:
            d_1716___mcc_h3_ = source17_.cf5
            d_1717_cf5_ = d_1716___mcc_h3_
            d_1718_v61_: _dafny.Seq
            d_1718_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smxtmhf"))
            d_1719_v62_: _dafny.Seq
            d_1719_v62_ = _dafny.SeqWithoutIsStrInference([d_1718_v61_, d_1718_v61_])
            d_1720_v63_: _dafny.Array
            nw340_ = _dafny.Array(None, 7)
            nw340_[int(0)] = default__.fm13(d_1717_cf5_, d_1717_cf5_, True, globalState)
            nw340_[int(1)] = d_1637_v2_
            nw340_[int(2)] = d_1637_v2_
            nw340_[int(3)] = d_1637_v2_
            nw340_[int(4)] = d_1637_v2_
            nw340_[int(5)] = (0) - (len(d_1719_v62_))
            nw340_[int(6)] = d_1637_v2_
            d_1720_v63_ = nw340_
            d_1721_v64_: _dafny.Seq
            d_1721_v64_ = _dafny.SeqWithoutIsStrInference([d_1720_v63_, d_1720_v63_, d_1720_v63_])
            d_1722_v65_: _dafny.Map
            d_1722_v65_ = _dafny.Map({d_1696_v42_: (d_1721_v64_)[default__.safeIndex(d_1637_v2_, len(d_1721_v64_))]})
            d_1723_v66_: _dafny.Array
            nw341_ = _dafny.Array(None, 28)
            nw341_[int(0)] = d_1720_v63_
            nw341_[int(1)] = (d_1721_v64_)[default__.safeIndex(len(d_1718_v61_), len(d_1721_v64_))]
            nw341_[int(2)] = d_1720_v63_
            nw341_[int(3)] = d_1720_v63_
            nw341_[int(4)] = d_1720_v63_
            nw341_[int(5)] = (d_1720_v63_ if d_1635_v0_ else d_1720_v63_)
            nw341_[int(6)] = d_1720_v63_
            nw341_[int(7)] = d_1720_v63_
            nw341_[int(8)] = d_1720_v63_
            nw341_[int(9)] = d_1720_v63_
            nw341_[int(10)] = d_1720_v63_
            nw341_[int(11)] = d_1720_v63_
            nw341_[int(12)] = (d_1720_v63_ if d_1717_cf5_ else d_1720_v63_)
            nw341_[int(13)] = d_1720_v63_
            nw341_[int(14)] = d_1720_v63_
            nw341_[int(15)] = d_1720_v63_
            nw341_[int(16)] = d_1720_v63_
            nw341_[int(17)] = d_1720_v63_
            nw341_[int(18)] = d_1720_v63_
            nw341_[int(19)] = d_1720_v63_
            nw341_[int(20)] = d_1720_v63_
            nw341_[int(21)] = d_1720_v63_
            nw341_[int(22)] = d_1720_v63_
            nw341_[int(23)] = d_1720_v63_
            nw341_[int(24)] = d_1720_v63_
            nw341_[int(25)] = d_1720_v63_
            nw341_[int(26)] = (((d_1722_v65_)[d_1696_v42_] if (d_1696_v42_) in (d_1722_v65_) else d_1720_v63_) if d_1635_v0_ else d_1720_v63_)
            nw341_[int(27)] = d_1720_v63_
            d_1723_v66_ = nw341_
            index343_ = default__.safeIndex(210, (d_1723_v66_).length(0))
            (d_1723_v66_)[index343_] = d_1720_v63_
            d_1724_v67_: D7
            d_1724_v67_ = D7_DC17(d_1717_cf5_, d_1717_cf5_, d_1717_cf5_, d_1635_v0_)
            index344_ = default__.safeIndex(210, (d_1723_v66_).length(0))
            rhs313_ = d_1637_v2_
            rhs314_ = ((D7_DC17(d_1635_v0_, d_1635_v0_, d_1717_cf5_, d_1717_cf5_) if d_1717_cf5_ else d_1724_v67_)).cf25
            rhs315_ = d_1720_v63_
            rhs316_ = default__.fm14(globalState)
            lhs167_ = d_1723_v66_
            lhs168_ = default__.safeIndex(210, (d_1723_v66_).length(0))
            d_1637_v2_ = rhs313_
            d_1635_v0_ = rhs314_
            lhs167_[lhs168_] = rhs315_
            d_1637_v2_ = rhs316_
            d_1725_v68_: _dafny.Map
            d_1725_v68_ = _dafny.Map({d_1635_v0_: d_1718_v61_})
            d_1718_v61_ = ((d_1725_v68_)[d_1635_v0_] if (d_1635_v0_) in (d_1725_v68_) else d_1718_v61_)
            d_1637_v2_ = d_1637_v2_
            d_1637_v2_ = d_1637_v2_
        elif source17_.is_DC5:
            d_1726___mcc_h4_ = source17_.cf6
            d_1727_cf6_ = d_1726___mcc_h4_
            d_1635_v0_ = not(not((D7_DC17(d_1635_v0_, d_1635_v0_, d_1635_v0_, not(not(d_1635_v0_)))).cf27))
            index345_ = default__.safeIndex(172, (d_1714_v59_).length(0))
            (d_1714_v59_)[index345_] = d_1635_v0_
            index346_ = default__.safeIndex(172, (d_1714_v59_).length(0))
            (d_1714_v59_)[index346_] = d_1635_v0_
            d_1728_v69_: C0
            nw342_ = C0()
            nw342_.ctor__()
            d_1728_v69_ = nw342_
            index347_ = default__.safeIndex(172, (d_1714_v59_).length(0))
            (d_1714_v59_)[index347_] = ((d_1712_v57_)[d_1637_v2_] if (d_1637_v2_) in (d_1712_v57_) else (d_1714_v59_)[default__.safeIndex(172, (d_1714_v59_).length(0))])
        elif source17_.is_DC6:
            d_1729___mcc_h5_ = source17_.cf7
            d_1730_cf7_ = d_1729___mcc_h5_
            d_1731_v70_: _dafny.Seq
            d_1731_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvnt"))
            if not((d_1731_v70_) <= (d_1731_v70_)):
                d_1731_v70_ = (d_1731_v70_) + (d_1731_v70_)
                d_1712_v57_ = (d_1712_v57_).set(default__.safeDivisionInt(d_1637_v2_, -564), d_1635_v0_)
                d_1637_v2_ = default__.safeDivisionInt(d_1730_cf7_, len((d_1731_v70_) + (d_1731_v70_)))
                d_1635_v0_ = d_1635_v0_
                d_1635_v0_ = d_1635_v0_
            elif True:
                d_1732_v71_: _dafny.Map
                d_1732_v71_ = _dafny.Map({d_1731_v70_: d_1637_v2_})
                d_1730_cf7_ = (0) - (((d_1732_v71_)[((_dafny.SeqWithoutIsStrInference([d_1696_v42_ for d_1733_i7_ in range(default__.abs(985))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvdekoxc")))).set(default__.safeIndex(d_1637_v2_, len((_dafny.SeqWithoutIsStrInference([d_1696_v42_ for d_1734_i7_ in range(default__.abs(985))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvdekoxc"))))), d_1696_v42_)] if (((_dafny.SeqWithoutIsStrInference([d_1696_v42_ for d_1735_i7_ in range(default__.abs(985))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvdekoxc")))).set(default__.safeIndex(d_1637_v2_, len((_dafny.SeqWithoutIsStrInference([d_1696_v42_ for d_1736_i7_ in range(default__.abs(985))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvdekoxc"))))), d_1696_v42_)) in (d_1732_v71_) else default__.fm13(d_1635_v0_, d_1635_v0_, not(d_1635_v0_), globalState)))
                d_1635_v0_ = d_1635_v0_
                d_1737_v72_: _dafny.Array
                def lambda115_(d_1738_cf7_):
                    def lambda116_(d_1739_i8_):
                        return (d_1739_i8_) + (d_1738_cf7_)

                    return lambda116_

                init65_ = lambda115_(d_1730_cf7_)
                nw343_ = _dafny.Array(None, 20)
                for i0_65_ in range(nw343_.length(0)):
                    nw343_[i0_65_] = init65_(i0_65_)
                d_1737_v72_ = nw343_
                d_1740_v73_: _dafny.Map
                d_1740_v73_ = _dafny.Map({d_1737_v72_: d_1730_cf7_})
                d_1740_v73_ = (d_1740_v73_).set(d_1737_v72_, d_1730_cf7_)
                d_1635_v0_ = d_1635_v0_
                d_1635_v0_ = (d_1635_v0_) or ((d_1637_v2_) == (d_1730_cf7_))
            d_1741_v74_: _dafny.Seq
            d_1741_v74_ = _dafny.SeqWithoutIsStrInference([(d_1711_v56_).fm3(-252, len(d_1636_v1_), globalState)])
            d_1741_v74_ = d_1741_v74_
            d_1742_v75_: D5
            d_1742_v75_ = D5_DC11(_dafny.Map({d_1635_v0_: d_1696_v42_}))
            d_1743_v76_: _dafny.Map
            d_1743_v76_ = _dafny.Map({d_1635_v0_: d_1696_v42_})
            d_1744_v77_: _dafny.Map
            d_1744_v77_ = _dafny.Map({default__.fm0(d_1637_v2_, d_1730_cf7_, globalState): _dafny.Map({d_1635_v0_: d_1696_v42_})})
            pat_let_tv43_ = d_1744_v77_
            pat_let_tv44_ = d_1696_v42_
            pat_let_tv45_ = d_1696_v42_
            pat_let_tv46_ = d_1744_v77_
            pat_let_tv47_ = d_1743_v76_
            pat_let_tv48_ = d_1742_v75_
            d_1745_v78_: _dafny.Array
            nw344_ = _dafny.Array(None, 26)
            nw344_[int(0)] = d_1742_v75_
            nw344_[int(1)] = d_1742_v75_
            nw344_[int(2)] = d_1742_v75_
            nw344_[int(3)] = D5_DC11(d_1743_v76_)
            nw344_[int(4)] = d_1742_v75_
            nw344_[int(5)] = d_1742_v75_
            nw344_[int(6)] = d_1742_v75_
            nw344_[int(7)] = d_1742_v75_
            nw344_[int(8)] = d_1742_v75_
            nw344_[int(9)] = D5_DC11((d_1743_v76_).set(d_1635_v0_, d_1696_v42_))
            nw344_[int(10)] = default__.fm38(d_1742_v75_, d_1730_cf7_, len(_dafny.SeqWithoutIsStrInference([d_1730_cf7_])), d_1637_v2_, globalState)
            nw344_[int(11)] = d_1742_v75_
            nw344_[int(12)] = (D5_DC11((d_1743_v76_).set(d_1635_v0_, (d_1731_v70_)[default__.safeIndex(d_1730_cf7_, len(d_1731_v70_))])) if d_1635_v0_ else d_1742_v75_)
            def iife140_(_pat_let42_0):
                def iife141_(d_1746_dt__update__tmp_h3_):
                    def iife142_(_pat_let43_0):
                        def iife143_(d_1747_dt__update_hcf15_h0_):
                            return D5_DC11(d_1747_dt__update_hcf15_h0_)
                        return iife143_(_pat_let43_0)
                    return iife142_(((pat_let_tv43_)[pat_let_tv44_] if (pat_let_tv45_) in (pat_let_tv46_) else pat_let_tv47_))
                return iife141_(_pat_let42_0)
            nw344_[int(13)] = iife140_(d_1742_v75_)
            nw344_[int(14)] = D5_DC11(d_1743_v76_)
            nw344_[int(15)] = d_1742_v75_
            nw344_[int(16)] = d_1742_v75_
            nw344_[int(17)] = d_1742_v75_
            nw344_[int(18)] = d_1742_v75_
            nw344_[int(19)] = D5_DC11(d_1743_v76_)
            nw344_[int(20)] = d_1742_v75_
            nw344_[int(21)] = d_1742_v75_
            nw344_[int(22)] = d_1742_v75_
            nw344_[int(23)] = D5_DC11(d_1743_v76_)
            nw344_[int(24)] = d_1742_v75_
            def iife144_(_pat_let44_0):
                def iife145_(d_1748_dt__update__tmp_h4_):
                    def iife146_(_pat_let45_0):
                        def iife147_(d_1749_dt__update_hcf15_h1_):
                            return D5_DC11(d_1749_dt__update_hcf15_h1_)
                        return iife147_(_pat_let45_0)
                    return iife146_((pat_let_tv48_).cf15)
                return iife145_(_pat_let44_0)
            nw344_[int(25)] = iife144_(d_1742_v75_)
            d_1745_v78_ = nw344_
            index348_ = default__.safeIndex(471, (d_1745_v78_).length(0))
            (d_1745_v78_)[index348_] = d_1742_v75_
            index349_ = default__.safeIndex(471, (d_1745_v78_).length(0))
            rhs317_ = d_1742_v75_
            rhs318_ = d_1696_v42_
            lhs169_ = d_1745_v78_
            lhs170_ = default__.safeIndex(471, (d_1745_v78_).length(0))
            lhs169_[lhs170_] = rhs317_
            d_1696_v42_ = rhs318_
            index350_ = default__.safeIndex(691, (d_1714_v59_).length(0))
            (d_1714_v59_)[index350_] = True
            index351_ = default__.safeIndex(691, (d_1714_v59_).length(0))
            (d_1714_v59_)[index351_] = ((d_1636_v1_)[d_1635_v0_] if (d_1635_v0_) in (d_1636_v1_) else (d_1731_v70_) < (d_1731_v70_))
        elif source17_.is_DC3:
            d_1750___mcc_h6_ = source17_.cf4
            d_1751_cf4_ = d_1750___mcc_h6_
            d_1752_v79_: C7
            nw345_ = C7()
            nw345_.ctor__()
            d_1752_v79_ = nw345_
            d_1753_v80_: _dafny.Map
            d_1753_v80_ = _dafny.Map({d_1637_v2_: _dafny.Map({d_1635_v0_: 481})})
            d_1753_v80_ = d_1753_v80_
            d_1754_v81_: _dafny.Seq
            d_1754_v81_ = _dafny.SeqWithoutIsStrInference([d_1635_v0_, d_1635_v0_, (d_1711_v56_).fm18(d_1635_v0_, d_1637_v2_, d_1635_v0_, 54, globalState), d_1635_v0_, d_1635_v0_])
            rhs319_ = d_1637_v2_
            rhs320_ = ((d_1754_v81_)[default__.safeIndex(d_1637_v2_, len(d_1754_v81_))]) or (d_1635_v0_)
            rhs321_ = d_1635_v0_
            rhs322_ = d_1696_v42_
            d_1637_v2_ = rhs319_
            d_1635_v0_ = rhs320_
            d_1635_v0_ = rhs321_
            d_1696_v42_ = rhs322_
            d_1754_v81_ = (d_1754_v81_).set(default__.safeIndex(d_1637_v2_, len(d_1754_v81_)), d_1635_v0_)
        elif True:
            d_1755___mcc_h7_ = source17_.cf8
            d_1756_cf8_ = d_1755___mcc_h7_
            d_1757_v82_: _dafny.MultiSet
            d_1757_v82_ = _dafny.MultiSet([d_1710_v55_, _dafny.SeqWithoutIsStrInference([(0) - (d_1637_v2_) for d_1758_i9_ in range(default__.abs(-841))]), d_1710_v55_])
            d_1635_v0_ = (d_1757_v82_).issubset(d_1757_v82_)
            d_1635_v0_ = ((d_1712_v57_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uadctpfh")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uadctpfh")))) in (d_1712_v57_) else (d_1711_v56_).fm3(d_1637_v2_, d_1637_v2_, globalState))
            d_1759_v83_: _dafny.Seq
            d_1759_v83_ = _dafny.SeqWithoutIsStrInference([d_1635_v0_])
            rhs323_ = (d_1711_v56_).fm18(d_1635_v0_, (default__.fm13(not(d_1635_v0_), d_1635_v0_, False, globalState)) * (d_1637_v2_), d_1635_v0_, d_1637_v2_, globalState)
            rhs324_ = (d_1637_v2_) - (default__.safeModuloInt(d_1637_v2_, d_1637_v2_))
            rhs325_ = (d_1637_v2_) != ((0) - (default__.fm13((d_1759_v83_)[default__.safeIndex(d_1637_v2_, len(d_1759_v83_))], d_1635_v0_, d_1635_v0_, globalState)))
            rhs326_ = (d_1713_v58_).ispropersubset(d_1713_v58_)
            d_1635_v0_ = rhs323_
            d_1637_v2_ = rhs324_
            d_1635_v0_ = rhs325_
            d_1635_v0_ = rhs326_
            d_1637_v2_ = d_1637_v2_
        d_1760_v84_: _dafny.Array
        def lambda117_(d_1761_v42_, d_1762_v55_, d_1763_v0_):
            def lambda118_(d_1764_i10_):
                return _dafny.Set({True, (D5_DC13(_dafny.SeqWithoutIsStrInference([d_1761_v42_ for d_1765_i11_ in range(default__.abs(625))]), len(d_1762_v55_), d_1763_v0_)).cf18, False})

            return lambda118_

        init66_ = lambda117_(d_1696_v42_, d_1710_v55_, d_1635_v0_)
        nw346_ = _dafny.Array(None, 14)
        for i0_66_ in range(nw346_.length(0)):
            nw346_[i0_66_] = init66_(i0_66_)
        d_1760_v84_ = nw346_
        d_1766_v85_: C1
        nw347_ = C1()
        nw347_.ctor__(d_1760_v84_, d_1696_v42_)
        d_1766_v85_ = nw347_
        d_1767_i12_: int
        d_1767_i12_ = 0
        with _dafny.label("7"):
            while d_1635_v0_:
                with _dafny.c_label("7"):
                    if (d_1767_i12_) >= (100):
                        raise _dafny.Break("7")
                    d_1767_i12_ = (d_1767_i12_) + (1)
                    d_1768_v86_: _dafny.Array
                    nw348_ = _dafny.Array(int(0), 25)
                    d_1768_v86_ = nw348_
                    index352_ = default__.safeIndex(345, (d_1768_v86_).length(0))
                    (d_1768_v86_)[index352_] = d_1637_v2_
                    index353_ = default__.safeIndex(345, (d_1768_v86_).length(0))
                    (d_1768_v86_)[index353_] = ((d_1709_v54_)[d_1635_v0_] if (d_1635_v0_) in (d_1709_v54_) else d_1637_v2_)
                    d_1635_v0_ = d_1635_v0_
                    d_1769_v87_: _dafny.Map
                    d_1769_v87_ = _dafny.Map({not(d_1635_v0_): d_1696_v42_})
                    d_1770_v88_: D5
                    d_1770_v88_ = D5_DC11(d_1769_v87_)
                    d_1771_v89_: _dafny.Map
                    d_1771_v89_ = _dafny.Map({d_1770_v88_: d_1635_v0_})
                    d_1772_v90_: _dafny.Seq
                    d_1772_v90_ = _dafny.SeqWithoutIsStrInference([d_1771_v89_])
                    d_1772_v90_ = d_1772_v90_
                    d_1773_v91_: _dafny.Map
                    d_1773_v91_ = _dafny.Map({(d_1768_v86_)[default__.safeIndex(345, (d_1768_v86_).length(0))]: (0) - (-668)})
                    d_1773_v91_ = (d_1773_v91_).set(279, ((d_1709_v54_)[d_1635_v0_] if (d_1635_v0_) in (d_1709_v54_) else (d_1768_v86_)[default__.safeIndex(345, (d_1768_v86_).length(0))]))
                    pass
            pass
        d_1774_v92_: _dafny.MultiSet
        d_1774_v92_ = _dafny.MultiSet([d_1635_v0_, d_1635_v0_, d_1635_v0_])
        if (d_1774_v92_).ispropersubset(d_1774_v92_):
            index354_ = default__.safeIndex(259, (d_1714_v59_).length(0))
            (d_1714_v59_)[index354_] = d_1635_v0_
            index355_ = default__.safeIndex(259, (d_1714_v59_).length(0))
            (d_1714_v59_)[index355_] = d_1635_v0_
            d_1635_v0_ = default__.fm1(d_1637_v2_, d_1637_v2_, (d_1714_v59_)[default__.safeIndex(259, (d_1714_v59_).length(0))], d_1635_v0_, globalState)
            d_1775_v93_: _dafny.Set
            d_1775_v93_ = _dafny.Set({not(d_1635_v0_)})
            rhs327_ = d_1775_v93_
            rhs328_ = d_1696_v42_
            d_1775_v93_ = rhs327_
            d_1696_v42_ = rhs328_
            d_1776_v94_: _dafny.Array
            nw349_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
            d_1776_v94_ = nw349_
            d_1776_v94_ = (d_1776_v94_ if d_1635_v0_ else d_1776_v94_)
            d_1777_v95_: _dafny.Seq
            d_1777_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smj"))
            d_1635_v0_ = (d_1777_v95_) <= (d_1777_v95_)
        elif True:
            d_1778_v96_: _dafny.MultiSet
            d_1778_v96_ = _dafny.MultiSet([d_1714_v59_, d_1714_v59_, d_1714_v59_])
            d_1778_v96_ = _dafny.MultiSet([d_1714_v59_, d_1714_v59_])
            d_1636_v1_ = (d_1636_v1_).set(((0) - (d_1637_v2_)) >= ((0) - (d_1637_v2_)), not(d_1635_v0_))
            d_1779_v97_: _dafny.Array
            nw350_ = _dafny.Array(int(0), 7)
            d_1779_v97_ = nw350_
            d_1779_v97_ = d_1779_v97_
            d_1780_v98_: _dafny.Seq
            d_1780_v98_ = _dafny.SeqWithoutIsStrInference([True, d_1635_v0_])
            d_1781_v99_: _dafny.Map
            d_1781_v99_ = _dafny.Map({d_1635_v0_: (d_1780_v98_) + (d_1780_v98_)})
            d_1782_v100_: _dafny.Set
            d_1782_v100_ = _dafny.Set({d_1635_v0_})
            d_1781_v99_ = (d_1781_v99_).set((d_1766_v85_).fm5(d_1637_v2_, d_1635_v0_, len(d_1782_v100_), d_1637_v2_, globalState), ((d_1780_v98_).set(default__.safeIndex(d_1637_v2_, len(d_1780_v98_)), not((d_1766_v85_).fm5(d_1637_v2_, d_1635_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))), d_1637_v2_, globalState)))) + (d_1780_v98_))
            d_1783_v101_: D7
            d_1783_v101_ = D7_DC17(d_1635_v0_, True, not(d_1635_v0_), d_1635_v0_)
            d_1784_v102_: _dafny.Seq
            d_1784_v102_ = _dafny.SeqWithoutIsStrInference([d_1783_v101_])
            d_1784_v102_ = _dafny.SeqWithoutIsStrInference([d_1783_v101_])
        d_1785_v103_: D2
        d_1785_v103_ = D2_DC6(374)
        r0 = d_1785_v103_
        return r0


class C9(T0, T1):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f4, f5):
        (self)._f4 = f4
        (self)._f5 = f5

    def fm3(self, p0, p1, globalState):
        return ((True) or (True)) not in (_dafny.MultiSet([False]))

    def fm4(self, p0, globalState):
        return default__.safeDivisionInt(len(_dafny.Set({False, False})), (len(_dafny.SeqWithoutIsStrInference([(D5_DC13(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1786_i0_ in range(default__.abs(-635))]), 556, False)).cf17]))) - (len(_dafny.SeqWithoutIsStrInference([True, True]))))

    def fm5(self, p0, p1, p2, p3, globalState):
        return (len(_dafny.Map({not(False): True}))) >= (703)

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        d_1787_v0_: bool
        d_1787_v0_ = True
        d_1788_v1_: int
        d_1788_v1_ = -526
        source18_ = D2_DC4((d_1787_v0_) or (default__.fm1(d_1788_v1_, d_1788_v1_, d_1787_v0_, d_1787_v0_, globalState)))
        if source18_.is_DC4:
            d_1789___mcc_h0_ = source18_.cf5
            d_1790_cf5_ = d_1789___mcc_h0_
            d_1791_v2_: C5
            nw351_ = C5()
            nw351_.ctor__()
            d_1791_v2_ = nw351_
            d_1790_cf5_ = d_1787_v0_
            if d_1790_cf5_:
                d_1792_v3_: C5
                nw352_ = C5()
                nw352_.ctor__()
                d_1792_v3_ = nw352_
                d_1793_v4_: _dafny.Seq
                d_1793_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                d_1794_v5_: _dafny.Array
                nw353_ = _dafny.Array(None, 8)
                nw353_[int(0)] = d_1787_v0_
                nw353_[int(1)] = d_1787_v0_
                nw353_[int(2)] = d_1790_cf5_
                nw353_[int(3)] = True
                nw353_[int(4)] = d_1790_cf5_
                nw353_[int(5)] = d_1790_cf5_
                nw353_[int(6)] = d_1790_cf5_
                nw353_[int(7)] = d_1787_v0_
                d_1794_v5_ = nw353_
                d_1795_v6_: _dafny.Set
                d_1795_v6_ = _dafny.Set({d_1794_v5_})
                d_1796_v7_: _dafny.Set
                d_1796_v7_ = d_1795_v6_
                d_1797_v8_: _dafny.Map
                d_1797_v8_ = _dafny.Map({d_1790_cf5_: d_1796_v7_})
                d_1798_v9_: _dafny.Map
                d_1799_v10_: int
                d_1800_v11_: bool
                d_1801_v12_: _dafny.Map
                out17_: _dafny.Map
                out18_: int
                out19_: bool
                out20_: _dafny.Map
                out17_, out18_, out19_, out20_ = (self).m5(d_1793_v4_, (d_1797_v8_).set(not(d_1787_v0_), d_1796_v7_), d_1788_v1_, globalState)
                d_1798_v9_ = out17_
                d_1799_v10_ = out18_
                d_1800_v11_ = out19_
                d_1801_v12_ = out20_
                d_1800_v11_ = d_1800_v11_
                d_1802_v13_: _dafny.Set
                d_1802_v13_ = _dafny.Set({d_1800_v11_, True, d_1800_v11_})
                d_1803_v14_: _dafny.Set
                d_1803_v14_ = _dafny.Set({d_1802_v13_, _dafny.Set({d_1800_v11_, True, d_1790_cf5_}), (d_1802_v13_) | (_dafny.Set({d_1800_v11_}))})
                d_1803_v14_ = default__.fm39((d_1793_v4_) + (d_1793_v4_), globalState)
                d_1804_v15_: _dafny.Array
                def lambda119_(d_1805_v1_):
                    def lambda120_(d_1806_i0_):
                        return (d_1806_i0_) + (d_1805_v1_)

                    return lambda120_

                init67_ = lambda119_(d_1788_v1_)
                nw354_ = _dafny.Array(None, 17)
                for i0_67_ in range(nw354_.length(0)):
                    nw354_[i0_67_] = init67_(i0_67_)
                d_1804_v15_ = nw354_
                d_1807_v16_: _dafny.Map
                d_1807_v16_ = _dafny.Map({d_1788_v1_: d_1804_v15_})
                d_1808_v17_: _dafny.MultiSet
                d_1808_v17_ = _dafny.MultiSet([d_1787_v0_])
                d_1807_v16_ = (d_1807_v16_).set(default__.safeDivisionInt(d_1799_v10_, ((d_1808_v17_)[True] if (True) in (d_1808_v17_) else d_1799_v10_)), d_1804_v15_)
            elif True:
                d_1809_v18_: _dafny.Seq
                d_1809_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcer"))
                d_1788_v1_ = len(d_1809_v18_)
                d_1810_v19_: C2
                nw355_ = C2()
                nw355_.ctor__((self).f4, ((self).f5 if d_1790_cf5_ else _dafny.CodePoint('p')))
                d_1810_v19_ = nw355_
                d_1811_v20_: str
                d_1811_v20_ = _dafny.CodePoint('t')
                rhs329_ = (not (not(d_1787_v0_)) or (False) if True else not (d_1787_v0_) or (True))
                rhs330_ = d_1810_v19_
                rhs331_ = d_1811_v20_
                d_1787_v0_ = rhs329_
                d_1810_v19_ = rhs330_
                d_1811_v20_ = rhs331_
                d_1790_cf5_ = d_1790_cf5_
                d_1812_v21_: _dafny.Seq
                d_1812_v21_ = _dafny.SeqWithoutIsStrInference([d_1788_v1_])
                d_1813_v22_: _dafny.Map
                d_1813_v22_ = _dafny.Map({d_1812_v21_: d_1788_v1_})
                d_1787_v0_ = (d_1812_v21_) not in (((d_1813_v22_)).set(d_1812_v21_, len(_dafny.Map({631: not(d_1790_cf5_)}))))
                r1 = default__.safeDivisionInt(d_1788_v1_, d_1788_v1_)
            d_1814_v23_: _dafny.Set
            d_1814_v23_ = _dafny.Set({d_1788_v1_, d_1788_v1_, d_1788_v1_})
            d_1815_v24_: _dafny.Seq
            d_1815_v24_ = _dafny.SeqWithoutIsStrInference([d_1790_cf5_, (635) != (d_1788_v1_), (d_1814_v23_).issubset(d_1814_v23_)])
            d_1816_v25_: _dafny.Array
            nw356_ = _dafny.Array(int(0), 8)
            d_1816_v25_ = nw356_
            index356_ = default__.safeIndex(407, (d_1816_v25_).length(0))
            (d_1816_v25_)[index356_] = d_1788_v1_
            d_1817_v26_: _dafny.Seq
            d_1817_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
            d_1818_v27_: str
            d_1818_v27_ = _dafny.CodePoint('n')
            index357_ = default__.safeIndex(407, (d_1816_v25_).length(0))
            rhs332_ = d_1815_v24_
            rhs333_ = d_1788_v1_
            rhs334_ = d_1788_v1_
            rhs335_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nisbhej"))).set(default__.safeIndex(d_1788_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nisbhej")))), (self).f5)
            rhs336_ = _dafny.CodePoint('b')
            lhs171_ = d_1816_v25_
            lhs172_ = default__.safeIndex(407, (d_1816_v25_).length(0))
            d_1815_v24_ = rhs332_
            lhs171_[lhs172_] = rhs333_
            r1 = rhs334_
            d_1817_v26_ = rhs335_
            d_1818_v27_ = rhs336_
        elif source18_.is_DC5:
            d_1819___mcc_h1_ = source18_.cf6
            d_1820_cf6_ = d_1819___mcc_h1_
            d_1787_v0_ = True
            d_1787_v0_ = d_1787_v0_
            d_1821_v28_: D20
            d_1821_v28_ = D20_DC47((self).f4)
            d_1822_v29_: C3
            nw357_ = C3()
            nw357_.ctor__((d_1821_v28_).cf73, (self).f5)
            d_1822_v29_ = nw357_
            if d_1787_v0_:
                r1 = d_1788_v1_
                d_1823_v30_: _dafny.MultiSet
                d_1823_v30_ = _dafny.MultiSet([d_1787_v0_])
                d_1824_v31_: _dafny.Map
                d_1824_v31_ = _dafny.Map({d_1788_v1_: (default__.fm15(d_1787_v0_, d_1787_v0_, d_1788_v1_, globalState)) == (d_1823_v30_)})
                d_1824_v31_ = (d_1824_v31_).set(d_1788_v1_, d_1787_v0_)
                d_1825_v32_: _dafny.Seq
                d_1825_v32_ = _dafny.SeqWithoutIsStrInference([d_1787_v0_, d_1787_v0_, not(True), d_1787_v0_])
                d_1826_v33_: _dafny.Seq
                d_1826_v33_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1825_v32_), (d_1823_v30_) | (d_1823_v30_)])
                d_1826_v33_ = (_dafny.SeqWithoutIsStrInference([d_1823_v30_, d_1823_v30_])) + (d_1826_v33_)
                d_1825_v32_ = (((d_1825_v32_) + (d_1825_v32_)) + (d_1825_v32_)).set(default__.safeIndex((0) - ((d_1788_v1_) * (d_1788_v1_)), len(((d_1825_v32_) + (d_1825_v32_)) + (d_1825_v32_))), (d_1787_v0_) == (d_1787_v0_))
                d_1827_v34_: _dafny.Array
                nw358_ = _dafny.Array(False, 12)
                d_1827_v34_ = nw358_
                d_1828_v35_: _dafny.Set
                d_1828_v35_ = _dafny.Set({d_1827_v34_, d_1827_v34_, d_1827_v34_})
                d_1829_v36_: _dafny.Seq
                d_1829_v36_ = _dafny.SeqWithoutIsStrInference([d_1788_v1_])
                index358_ = default__.safeIndex(772, (d_1827_v34_).length(0))
                (d_1827_v34_)[index358_] = (d_1788_v1_) == ((self).fm4(d_1829_v36_, globalState))
                d_1830_v37_: _dafny.Array
                def lambda121_(d_1831_v1_):
                    def lambda122_(d_1832_i1_):
                        return (d_1832_i1_) * (d_1831_v1_)

                    return lambda122_

                init68_ = lambda121_(d_1788_v1_)
                nw359_ = _dafny.Array(None, 10)
                for i0_68_ in range(nw359_.length(0)):
                    nw359_[i0_68_] = init68_(i0_68_)
                d_1830_v37_ = nw359_
                d_1833_v38_: _dafny.Map
                d_1833_v38_ = _dafny.Map({d_1787_v0_: d_1828_v35_})
                index359_ = default__.safeIndex(772, (d_1827_v34_).length(0))
                rhs337_ = (((d_1833_v38_)[True] if (True) in (d_1833_v38_) else d_1828_v35_)) | (d_1828_v35_)
                rhs338_ = d_1787_v0_
                rhs339_ = d_1820_cf6_
                rhs340_ = d_1830_v37_
                rhs341_ = d_1787_v0_
                lhs173_ = d_1827_v34_
                lhs174_ = default__.safeIndex(772, (d_1827_v34_).length(0))
                d_1828_v35_ = rhs337_
                lhs173_[lhs174_] = rhs338_
                d_1820_cf6_ = rhs339_
                d_1830_v37_ = rhs340_
                d_1787_v0_ = rhs341_
            elif True:
                r1 = default__.safeDivisionInt(d_1788_v1_, d_1788_v1_)
                d_1834_v39_: _dafny.Set
                d_1834_v39_ = _dafny.Set({(0) - (d_1788_v1_)})
                d_1835_v40_: _dafny.Set
                d_1835_v40_ = _dafny.Set({d_1834_v39_, d_1834_v39_})
                d_1836_v41_: _dafny.Seq
                d_1836_v41_ = _dafny.SeqWithoutIsStrInference([d_1788_v1_, len(d_1835_v40_), d_1788_v1_, default__.fm13(False, d_1787_v0_, d_1787_v0_, globalState)])
                d_1837_v42_: _dafny.Set
                d_1837_v42_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(D14_DC31(d_1788_v1_)).cf48 for d_1838_i2_ in range(default__.abs(960))]), d_1836_v41_, (d_1836_v41_) + (d_1836_v41_)})
                d_1837_v42_ = (d_1837_v42_) - (d_1837_v42_)
                d_1839_v43_: _dafny.MultiSet
                d_1839_v43_ = _dafny.MultiSet([d_1787_v0_])
                d_1787_v0_ = not((d_1839_v43_).issubset((d_1839_v43_) - (d_1839_v43_)))
                d_1840_v44_: _dafny.Map
                d_1840_v44_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1787_v0_, d_1787_v0_, d_1787_v0_]): d_1788_v1_})
                d_1841_v45_: D14
                d_1841_v45_ = D14_DC32(d_1788_v1_, d_1787_v0_, d_1787_v0_)
                d_1842_v46_: _dafny.Map
                d_1842_v46_ = _dafny.Map({d_1841_v45_: _dafny.SeqWithoutIsStrInference([(self).f5 for d_1843_i3_ in range(default__.abs(957))])})
                d_1844_v48_: _dafny.Map
                d_1844_v48_ = _dafny.Map({d_1841_v45_: _dafny.Map({d_1787_v0_: d_1787_v0_})})
                d_1845_v49_: D21
                d_1845_v49_ = D21_DC50(d_1844_v48_)
                def iife148_():
                    coll56_ = _dafny.Map()
                    compr_56_: D14
                    for compr_56_ in ((d_1845_v49_).cf79).keys.Elements:
                        d_1846_v47_: D14 = compr_56_
                        if (d_1846_v47_) in ((d_1845_v49_).cf79):
                            coll56_[d_1846_v47_] = d_1820_cf6_
                    return _dafny.Map(coll56_)
                rhs342_ = default__.fm40(d_1788_v1_, globalState)
                rhs343_ = (iife148_()
                 if (d_1788_v1_) != (d_1788_v1_) else _dafny.Map({d_1841_v45_: d_1820_cf6_}))
                rhs344_ = (d_1820_cf6_) + ((((d_1820_cf6_).set(default__.safeIndex(d_1788_v1_, len(d_1820_cf6_)), (self).f5)) + (d_1820_cf6_)).set(default__.safeIndex((0) - ((d_1839_v43_).cardinality), len(((d_1820_cf6_).set(default__.safeIndex(d_1788_v1_, len(d_1820_cf6_)), (self).f5)) + (d_1820_cf6_))), _dafny.CodePoint('u')))
                rhs345_ = d_1788_v1_
                rhs346_ = _dafny.SeqWithoutIsStrInference([(self).f5 for d_1847_i4_ in range(default__.abs(478))])
                d_1840_v44_ = rhs342_
                d_1842_v46_ = rhs343_
                d_1820_cf6_ = rhs344_
                r1 = rhs345_
                d_1820_cf6_ = rhs346_
                r1 = d_1788_v1_
        elif source18_.is_DC6:
            d_1848___mcc_h2_ = source18_.cf7
            d_1849_cf7_ = d_1848___mcc_h2_
            d_1850_v50_: _dafny.MultiSet
            d_1850_v50_ = _dafny.MultiSet([d_1787_v0_])
            d_1851_v51_: _dafny.Set
            d_1851_v51_ = _dafny.Set({d_1788_v1_, (d_1850_v50_).cardinality})
            d_1852_v52_: D16
            d_1852_v52_ = D16_DC37(d_1851_v51_)
            d_1853_v53_: D16
            d_1853_v53_ = D16_DC39(D16_DC39(d_1852_v52_))
            source19_ = d_1853_v53_
            if source19_.is_DC38:
                d_1854___mcc_h5_ = source19_.cf57
                d_1855___mcc_h6_ = source19_.cf58
                d_1856___mcc_h7_ = source19_.cf59
                d_1857___mcc_h8_ = source19_.cf60
                d_1858_cf60_ = d_1857___mcc_h8_
                d_1859_cf59_ = d_1856___mcc_h7_
                d_1860_cf58_ = d_1855___mcc_h6_
                d_1861_cf57_ = d_1854___mcc_h5_
                d_1862_v54_: _dafny.Seq
                d_1862_v54_ = _dafny.SeqWithoutIsStrInference([d_1788_v1_, d_1859_cf59_, d_1788_v1_])
                nw360_ = _dafny.Array(None, 4)
                nw360_[int(0)] = d_1860_cf58_
                nw360_[int(1)] = (d_1849_cf7_) in (d_1862_v54_)
                nw360_[int(2)] = d_1787_v0_
                nw360_[int(3)] = (default__.fm1(d_1849_cf7_, d_1849_cf7_, True, d_1860_cf58_, globalState) if d_1858_cf60_ else True)
                (globalState).f3 = nw360_
                d_1863_v55_: _dafny.MultiSet
                d_1863_v55_ = _dafny.MultiSet([(d_1861_cf57_ if True else d_1861_cf57_), (d_1861_cf57_) + (d_1861_cf57_)])
                r1 = ((d_1863_v55_)[_dafny.SeqWithoutIsStrInference([(self).f5 for d_1864_i5_ in range(default__.abs(184))])] if (_dafny.SeqWithoutIsStrInference([(self).f5 for d_1865_i5_ in range(default__.abs(184))])) in (d_1863_v55_) else (-840) - (d_1849_cf7_))
                d_1866_v56_: T1
                nw361_ = C2()
                nw361_.ctor__((self).f4, (self).f5)
                d_1866_v56_ = nw361_
                d_1866_v56_ = d_1866_v56_
                d_1867_v57_: _dafny.Array
                def lambda123_(d_1868_v0_):
                    def lambda124_(d_1869_i6_):
                        return d_1868_v0_

                    return lambda124_

                init69_ = lambda123_(d_1787_v0_)
                nw362_ = _dafny.Array(None, 1)
                for i0_69_ in range(nw362_.length(0)):
                    nw362_[i0_69_] = init69_(i0_69_)
                d_1867_v57_ = nw362_
                index360_ = default__.safeIndex(907, (d_1867_v57_).length(0))
                (d_1867_v57_)[index360_] = d_1858_cf60_
                index361_ = default__.safeIndex(907, (d_1867_v57_).length(0))
                (d_1867_v57_)[index361_] = d_1858_cf60_
            elif source19_.is_DC37:
                d_1870___mcc_h9_ = source19_.cf56
                d_1871_cf56_ = d_1870___mcc_h9_
                d_1872_v58_: T0
                nw363_ = C0()
                nw363_.ctor__()
                d_1872_v58_ = nw363_
                d_1873_v59_: _dafny.Seq
                d_1873_v59_ = _dafny.SeqWithoutIsStrInference([d_1872_v58_, d_1872_v58_, d_1872_v58_, d_1872_v58_])
                d_1873_v59_ = _dafny.SeqWithoutIsStrInference([d_1872_v58_])
                d_1874_v60_: str
                d_1874_v60_ = _dafny.CodePoint('n')
                d_1874_v60_ = (self).f5
                d_1787_v0_ = not ((d_1849_cf7_) == (d_1788_v1_)) or (not (d_1787_v0_) or (d_1787_v0_))
                d_1875_v61_: _dafny.Map
                d_1875_v61_ = _dafny.Map({(0) - (d_1788_v1_): d_1874_v60_})
                d_1876_v62_: _dafny.Map
                d_1876_v62_ = _dafny.Map({default__.fm0(d_1849_cf7_, d_1788_v1_, globalState): d_1875_v61_})
                d_1877_v63_: _dafny.Array
                nw364_ = _dafny.Array(None, 4)
                nw364_[int(0)] = d_1875_v61_
                nw364_[int(1)] = d_1875_v61_
                nw364_[int(2)] = (d_1875_v61_) | (d_1875_v61_)
                nw364_[int(3)] = (d_1875_v61_) | (((d_1876_v62_)[d_1874_v60_] if (d_1874_v60_) in (d_1876_v62_) else d_1875_v61_))
                d_1877_v63_ = nw364_
                d_1877_v63_ = d_1877_v63_
            elif True:
                d_1878___mcc_h10_ = source19_.cf61
                d_1879_cf61_ = d_1878___mcc_h10_
                d_1849_cf7_ = (default__.fm41(d_1787_v0_, d_1788_v1_, d_1849_cf7_, d_1788_v1_, globalState)).cf59
                d_1880_v64_: _dafny.Array
                nw365_ = _dafny.Array(int(0), 12)
                d_1880_v64_ = nw365_
                d_1881_v65_: _dafny.MultiSet
                d_1881_v65_ = _dafny.MultiSet([d_1849_cf7_, d_1788_v1_])
                d_1882_v66_: _dafny.Seq
                d_1882_v66_ = _dafny.SeqWithoutIsStrInference([d_1788_v1_, (d_1881_v65_).cardinality])
                index362_ = default__.safeIndex(912, (d_1880_v64_).length(0))
                (d_1880_v64_)[index362_] = (self).fm4(d_1882_v66_, globalState)
                index363_ = default__.safeIndex(912, (d_1880_v64_).length(0))
                (d_1880_v64_)[index363_] = d_1849_cf7_
                d_1883_v67_: _dafny.Seq
                d_1883_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "milgoae"))
                d_1883_v67_ = (d_1883_v67_) + (d_1883_v67_)
                d_1884_v68_: bool
                d_1884_v68_ = d_1787_v0_
                d_1885_v69_: _dafny.Map
                d_1885_v69_ = _dafny.Map({d_1787_v0_: (d_1880_v64_)[default__.safeIndex(912, (d_1880_v64_).length(0))]})
                d_1886_v70_: _dafny.Array
                nw366_ = _dafny.Array(False, 11)
                d_1886_v70_ = nw366_
                d_1887_v71_: D18
                d_1887_v71_ = D18_DC45(d_1886_v70_, d_1787_v0_, d_1788_v1_, 474)
                d_1884_v68_ = default__.fm42((d_1880_v64_)[default__.safeIndex(912, (d_1880_v64_).length(0))], ((d_1885_v69_)[d_1787_v0_] if (d_1787_v0_) in (d_1885_v69_) else d_1849_cf7_), d_1787_v0_, (d_1887_v71_).cf70, globalState)
            d_1888_v72_: _dafny.Seq
            d_1888_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f5: -467})])
            d_1889_v73_: C4
            nw367_ = C4()
            nw367_.ctor__(len(_dafny.SeqWithoutIsStrInference([(self).f5])), len(d_1888_v72_), (self).f4, (self).f5)
            d_1889_v73_ = nw367_
            (d_1889_v73_).f11 = -506
            d_1787_v0_ = not (d_1787_v0_) or (d_1787_v0_)
        elif source18_.is_DC3:
            d_1890___mcc_h3_ = source18_.cf4
            d_1891_cf4_ = d_1890___mcc_h3_
            d_1892_v74_: str
            d_1892_v74_ = _dafny.CodePoint('g')
            d_1892_v74_ = d_1892_v74_
            d_1893_v75_: _dafny.Set
            d_1893_v75_ = _dafny.Set({d_1787_v0_, d_1787_v0_, d_1787_v0_})
            d_1788_v1_ = default__.fm13(d_1787_v0_, d_1787_v0_, (d_1788_v1_) == (len(d_1893_v75_)), globalState)
            d_1894_v76_: _dafny.Seq
            d_1894_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hh"))
            d_1895_v77_: D7
            d_1895_v77_ = D7_DC16(d_1892_v74_)
            d_1896_v78_: _dafny.Array
            def lambda125_(d_1897_v0_):
                def lambda126_(d_1898_i7_):
                    return _dafny.SeqWithoutIsStrInference([d_1897_v0_])

                return lambda126_

            init70_ = lambda125_(d_1787_v0_)
            nw368_ = _dafny.Array(None, 2)
            for i0_70_ in range(nw368_.length(0)):
                nw368_[i0_70_] = init70_(i0_70_)
            d_1896_v78_ = nw368_
            d_1899_v79_: D4
            d_1899_v79_ = D4_DC9(d_1896_v78_)
            d_1900_v80_: _dafny.Map
            d_1900_v80_ = _dafny.Map({d_1895_v77_: d_1899_v79_})
            d_1901_v81_: _dafny.Map
            d_1901_v81_ = _dafny.Map({(d_1900_v80_).set(d_1895_v77_, d_1899_v79_): (0) - (d_1788_v1_)})
            rhs347_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyyph"))) + (d_1894_v76_)
            rhs348_ = _dafny.CodePoint('p')
            rhs349_ = _dafny.SeqWithoutIsStrInference([default__.fm13(d_1787_v0_, d_1787_v0_, d_1787_v0_, globalState), ((d_1901_v81_)[d_1900_v80_] if (d_1900_v80_) in (d_1901_v81_) else d_1788_v1_), d_1788_v1_])
            d_1894_v76_ = rhs347_
            d_1892_v74_ = rhs348_
            r0 = rhs349_
            d_1902_v82_: _dafny.Array
            nw369_ = _dafny.Array(int(0), 28)
            d_1902_v82_ = nw369_
            index364_ = default__.safeIndex(143, (d_1902_v82_).length(0))
            (d_1902_v82_)[index364_] = d_1788_v1_
            index365_ = default__.safeIndex(143, (d_1902_v82_).length(0))
            (d_1902_v82_)[index365_] = default__.safeDivisionInt(d_1788_v1_, default__.fm14(globalState))
        elif True:
            d_1903___mcc_h4_ = source18_.cf8
            d_1904_cf8_ = d_1903___mcc_h4_
            d_1787_v0_ = True
            d_1905_v83_: C3
            nw370_ = C3()
            nw370_.ctor__((D20_DC47((self).f4)).cf73, _dafny.CodePoint('s'))
            d_1905_v83_ = nw370_
            d_1906_v84_: _dafny.Seq
            d_1906_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enprlnjp"))
            d_1907_v85_: _dafny.Seq
            d_1907_v85_ = _dafny.SeqWithoutIsStrInference([d_1788_v1_])
            d_1908_v86_: D14
            d_1908_v86_ = D14_DC32(d_1788_v1_, d_1787_v0_, d_1787_v0_)
            d_1909_v87_: _dafny.Seq
            d_1909_v87_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_1907_v85_)[default__.safeIndex((0) - ((d_1908_v86_).cf49), len(d_1907_v85_))])])
            d_1906_v84_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))).set(default__.safeIndex((d_1909_v87_)[default__.safeIndex(d_1788_v1_, len(d_1909_v87_))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))), (self).f5)
            d_1909_v87_ = (d_1907_v85_).set(default__.safeIndex(d_1788_v1_, len(d_1907_v85_)), d_1788_v1_)
        hi16_ = (d_1788_v1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrob"))))
        for d_1910_i8_ in range((0) - (d_1788_v1_), hi16_):
            d_1911_v88_: _dafny.Map
            d_1911_v88_ = _dafny.Map({(0) - (d_1910_i8_): _dafny.CodePoint('b')})
            d_1911_v88_ = (d_1911_v88_).set(d_1910_i8_, (self).f5)
            d_1912_v89_: _dafny.Seq
            d_1912_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eh"))
            rhs350_ = d_1912_v89_
            rhs351_ = (d_1910_i8_) != (d_1910_i8_)
            d_1912_v89_ = rhs350_
            d_1787_v0_ = rhs351_
            d_1913_v90_: _dafny.Set
            d_1913_v90_ = _dafny.Set({d_1787_v0_, d_1787_v0_})
            d_1914_v91_: _dafny.Seq
            d_1914_v91_ = _dafny.SeqWithoutIsStrInference([d_1787_v0_, d_1787_v0_, d_1787_v0_, d_1787_v0_, not((_dafny.Set({d_1787_v0_})).issubset(d_1913_v90_))])
            d_1915_v92_: _dafny.Seq
            d_1915_v92_ = _dafny.SeqWithoutIsStrInference([d_1910_i8_])
            d_1787_v0_ = (d_1914_v91_)[default__.safeIndex((_dafny.MultiSet(d_1915_v92_)).cardinality, len(d_1914_v91_))]
            d_1916_v93_: _dafny.MultiSet
            d_1916_v93_ = _dafny.MultiSet([(self).f5])
            d_1917_v94_: _dafny.MultiSet
            d_1917_v94_ = _dafny.MultiSet([d_1787_v0_, d_1787_v0_, d_1787_v0_, (d_1916_v93_).issubset(d_1916_v93_), d_1787_v0_])
            d_1917_v94_ = d_1917_v94_
        d_1918_v95_: _dafny.Seq
        d_1918_v95_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_1787_v0_})), d_1788_v1_])
        d_1919_v96_: _dafny.Seq
        d_1919_v96_ = _dafny.SeqWithoutIsStrInference([d_1918_v95_])
        d_1920_v97_: _dafny.Map
        d_1920_v97_ = _dafny.Map({d_1788_v1_: d_1918_v95_})
        d_1921_v98_: _dafny.Seq
        d_1921_v98_ = _dafny.SeqWithoutIsStrInference([(len(d_1919_v96_)) != (len(d_1920_v97_))])
        d_1787_v0_ = (d_1921_v98_)[default__.safeIndex(default__.safeModuloInt(d_1788_v1_, (0) - (d_1788_v1_)), len(d_1921_v98_))]
        if d_1787_v0_:
            if not (d_1787_v0_) or (d_1787_v0_):
                d_1922_v99_: _dafny.MultiSet
                d_1922_v99_ = _dafny.MultiSet([d_1788_v1_])
                d_1923_v100_: _dafny.Seq
                d_1923_v100_ = _dafny.SeqWithoutIsStrInference([d_1922_v99_])
                d_1924_v101_: _dafny.Map
                d_1924_v101_ = _dafny.Map({(self).fm5(d_1788_v1_, d_1787_v0_, (0) - (d_1788_v1_), d_1788_v1_, globalState): ((d_1923_v100_)[default__.safeIndex(d_1788_v1_, len(d_1923_v100_))]).cardinality})
                d_1925_v102_: _dafny.Array
                nw371_ = _dafny.Array(False, 1)
                d_1925_v102_ = nw371_
                index366_ = default__.safeIndex(219, (d_1925_v102_).length(0))
                (d_1925_v102_)[index366_] = d_1787_v0_
                d_1926_v103_: _dafny.Map
                d_1926_v103_ = _dafny.Map({d_1788_v1_: 540})
                d_1927_v104_: _dafny.Map
                d_1927_v104_ = _dafny.Map({d_1787_v0_: d_1787_v0_})
                d_1928_v105_: _dafny.Seq
                d_1928_v105_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), (self).f5, (self).f5])
                d_1929_v106_: T1
                nw372_ = C2()
                nw372_.ctor__((self).f4, (d_1928_v105_)[default__.safeIndex(d_1788_v1_, len(d_1928_v105_))])
                d_1929_v106_ = nw372_
                d_1930_v107_: _dafny.Seq
                d_1930_v107_ = _dafny.SeqWithoutIsStrInference([d_1929_v106_])
                d_1931_v108_: _dafny.Set
                d_1931_v108_ = _dafny.Set({d_1787_v0_})
                d_1932_v109_: _dafny.Array
                nw373_ = _dafny.Array(None, 24)
                nw373_[int(0)] = 613
                nw373_[int(1)] = d_1788_v1_
                nw373_[int(2)] = d_1788_v1_
                nw373_[int(3)] = d_1788_v1_
                nw373_[int(4)] = d_1788_v1_
                nw373_[int(5)] = len(d_1926_v103_)
                nw373_[int(6)] = len(d_1927_v104_)
                nw373_[int(7)] = d_1788_v1_
                nw373_[int(8)] = d_1788_v1_
                nw373_[int(9)] = len(d_1930_v107_)
                nw373_[int(10)] = d_1788_v1_
                nw373_[int(11)] = d_1788_v1_
                nw373_[int(12)] = d_1788_v1_
                nw373_[int(13)] = d_1788_v1_
                nw373_[int(14)] = len(d_1931_v108_)
                nw373_[int(15)] = d_1788_v1_
                nw373_[int(16)] = d_1788_v1_
                nw373_[int(17)] = d_1788_v1_
                nw373_[int(18)] = d_1788_v1_
                nw373_[int(19)] = len(d_1928_v105_)
                nw373_[int(20)] = d_1788_v1_
                nw373_[int(21)] = d_1788_v1_
                nw373_[int(22)] = d_1788_v1_
                nw373_[int(23)] = d_1788_v1_
                d_1932_v109_ = nw373_
                d_1933_v110_: _dafny.Map
                d_1933_v110_ = _dafny.Map({d_1932_v109_: d_1788_v1_})
                index367_ = default__.safeIndex(219, (d_1925_v102_).length(0))
                rhs352_ = _dafny.Map({d_1787_v0_: d_1788_v1_})
                rhs353_ = (d_1788_v1_) >= (d_1788_v1_)
                rhs354_ = d_1787_v0_
                rhs355_ = (d_1788_v1_) < (len((_dafny.Map({d_1932_v109_: d_1788_v1_})) | (d_1933_v110_)))
                lhs175_ = d_1925_v102_
                lhs176_ = default__.safeIndex(219, (d_1925_v102_).length(0))
                d_1924_v101_ = rhs352_
                lhs175_[lhs176_] = rhs353_
                d_1787_v0_ = rhs354_
                d_1787_v0_ = rhs355_
                d_1788_v1_ = (self).fm4(_dafny.SeqWithoutIsStrInference([(D1_DC2(-871, d_1788_v1_)).cf2, (0) - (d_1788_v1_)]), globalState)
                d_1788_v1_ = (d_1788_v1_) - ((d_1788_v1_) - (len(default__.fm34(False, globalState))))
                d_1934_v111_: C1
                nw374_ = C1()
                nw374_.ctor__((self).f4, (d_1929_v106_).f5)
                d_1934_v111_ = nw374_
                d_1787_v0_ = (d_1921_v98_)[default__.safeIndex((default__.fm14(globalState)) * (d_1788_v1_), len(d_1921_v98_))]
            elif True:
                d_1935_v112_: C7
                nw375_ = C7()
                nw375_.ctor__()
                d_1935_v112_ = nw375_
                rhs356_ = d_1935_v112_
                rhs357_ = d_1788_v1_
                rhs358_ = d_1787_v0_
                d_1935_v112_ = rhs356_
                r1 = rhs357_
                d_1787_v0_ = rhs358_
                d_1936_v113_: C2
                nw376_ = C2()
                nw376_.ctor__((self).f4, (self).f5)
                d_1936_v113_ = nw376_
                d_1937_v114_: _dafny.Set
                d_1937_v114_ = _dafny.Set({d_1787_v0_, d_1787_v0_})
                d_1938_v115_: D10
                d_1938_v115_ = D10_DC21(d_1937_v114_)
                d_1939_v116_: D10
                d_1939_v116_ = D10_DC23(d_1938_v115_)
                d_1940_v117_: _dafny.Array
                nw377_ = _dafny.Array(False, 10)
                d_1940_v117_ = nw377_
                d_1941_v118_: _dafny.Map
                d_1941_v118_ = _dafny.Map({d_1940_v117_: d_1787_v0_})
                d_1942_v119_: _dafny.Map
                d_1942_v119_ = _dafny.Map({d_1939_v116_: (d_1940_v117_) not in (d_1941_v118_)})
                d_1942_v119_ = (d_1942_v119_).set(d_1939_v116_, d_1787_v0_)
                d_1943_v120_: _dafny.Array
                nw378_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1943_v120_ = nw378_
                d_1944_v121_: _dafny.Array
                def lambda127_(d_1945_i9_):
                    return default__.safeModuloInt(d_1945_i9_, -59)

                init71_ = lambda127_
                nw379_ = _dafny.Array(None, 2)
                for i0_71_ in range(nw379_.length(0)):
                    nw379_[i0_71_] = init71_(i0_71_)
                d_1944_v121_ = nw379_
                index368_ = default__.safeIndex(57, (d_1943_v120_).length(0))
                (d_1943_v120_)[index368_] = d_1944_v121_
                d_1946_v122_: _dafny.Seq
                d_1946_v122_ = _dafny.SeqWithoutIsStrInference([d_1944_v121_, d_1944_v121_])
                index369_ = default__.safeIndex(57, (d_1943_v120_).length(0))
                rhs359_ = (d_1946_v122_)[default__.safeIndex(261, len(d_1946_v122_))]
                rhs360_ = d_1788_v1_
                rhs361_ = (d_1788_v1_ if d_1787_v0_ else (d_1788_v1_) + (d_1788_v1_))
                lhs177_ = d_1943_v120_
                lhs178_ = default__.safeIndex(57, (d_1943_v120_).length(0))
                lhs177_[lhs178_] = rhs359_
                r1 = rhs360_
                d_1788_v1_ = rhs361_
                d_1947_v123_: _dafny.Seq
                d_1947_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evcvcc"))
                d_1948_v124_: D16
                d_1948_v124_ = D16_DC38((d_1947_v123_ if d_1787_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbw"))), d_1787_v0_, 797, d_1787_v0_)
                d_1949_v125_: str
                d_1949_v125_ = _dafny.CodePoint('d')
                rhs362_ = d_1948_v124_
                rhs363_ = (self).f5
                d_1948_v124_ = rhs362_
                d_1949_v125_ = rhs363_
            d_1950_v126_: _dafny.Seq
            d_1950_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyqox"))
            d_1951_v127_: D16
            d_1951_v127_ = D16_DC38(d_1950_v126_, d_1787_v0_, d_1788_v1_, d_1787_v0_)
            d_1952_v128_: D7
            d_1952_v128_ = D7_DC17((d_1951_v127_).cf58, d_1787_v0_, True, d_1787_v0_)
            pat_let_tv49_ = d_1787_v0_
            d_1953_v129_: D7
            def iife149_(_pat_let46_0):
                def iife150_(d_1954_dt__update__tmp_h0_):
                    def iife151_(_pat_let47_0):
                        def iife152_(d_1955_dt__update_hcf26_h0_):
                            def iife153_(_pat_let48_0):
                                def iife154_(d_1956_dt__update_hcf25_h0_):
                                    return D7_DC17((d_1954_dt__update__tmp_h0_).cf24, d_1956_dt__update_hcf25_h0_, d_1955_dt__update_hcf26_h0_, (d_1954_dt__update__tmp_h0_).cf27)
                                return iife154_(_pat_let48_0)
                            return iife153_(True)
                        return iife152_(_pat_let47_0)
                    return iife151_(pat_let_tv49_)
                return iife150_(_pat_let46_0)
            d_1953_v129_ = D7_DC17((self).fm3(len(d_1950_v126_), 261, globalState), (iife149_(d_1952_v128_)).cf24, d_1787_v0_, d_1787_v0_)
            rhs364_ = True
            rhs365_ = D7_DC17(d_1787_v0_, d_1787_v0_, d_1787_v0_, default__.fm1(d_1788_v1_, len(d_1950_v126_), d_1787_v0_, d_1787_v0_, globalState))
            d_1787_v0_ = rhs364_
            d_1953_v129_ = rhs365_
            d_1957_v130_: _dafny.Array
            nw380_ = _dafny.Array(int(0), 25)
            d_1957_v130_ = nw380_
            index370_ = default__.safeIndex(13, (d_1957_v130_).length(0))
            (d_1957_v130_)[index370_] = d_1788_v1_
            index371_ = default__.safeIndex(13, (d_1957_v130_).length(0))
            (d_1957_v130_)[index371_] = default__.safeDivisionInt((d_1788_v1_) * (d_1788_v1_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgucd")))) + (d_1788_v1_))
            d_1958_v131_: _dafny.Map
            d_1958_v131_ = _dafny.Map({(d_1788_v1_) - ((0) - (d_1788_v1_)): (d_1788_v1_) != ((d_1957_v130_)[default__.safeIndex(13, (d_1957_v130_).length(0))])})
            d_1787_v0_ = ((d_1958_v131_)[d_1788_v1_] if (d_1788_v1_) in (d_1958_v131_) else d_1787_v0_)
            d_1959_v132_: C2
            nw381_ = C2()
            nw381_.ctor__((self).f4, (self).f5)
            d_1959_v132_ = nw381_
        elif True:
            r1 = d_1788_v1_
            d_1960_v133_: _dafny.Seq
            d_1960_v133_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmep"))
            d_1961_v134_: _dafny.Array
            nw382_ = _dafny.Array(int(0), 15)
            d_1961_v134_ = nw382_
            d_1962_v135_: D11
            d_1962_v135_ = D11_DC25(d_1787_v0_, len(d_1960_v133_), d_1961_v134_, len(_dafny.Set({d_1788_v1_, -789})))
            d_1963_v136_: D11
            d_1963_v136_ = D11_DC26(d_1962_v135_)
            pat_let_tv50_ = d_1962_v135_
            def iife155_(_pat_let49_0):
                def iife156_(d_1964_dt__update__tmp_h1_):
                    def iife157_(_pat_let50_0):
                        def iife158_(d_1965_dt__update_hcf42_h0_):
                            return D11_DC26(d_1965_dt__update_hcf42_h0_)
                        return iife158_(_pat_let50_0)
                    return iife157_(D11_DC26(pat_let_tv50_))
                return iife156_(_pat_let49_0)
            source20_ = (iife155_(d_1963_v136_) if (d_1788_v1_) >= (798) else d_1963_v136_)
            if source20_.is_DC25:
                d_1966___mcc_h11_ = source20_.cf38
                d_1967___mcc_h12_ = source20_.cf39
                d_1968___mcc_h13_ = source20_.cf40
                d_1969___mcc_h14_ = source20_.cf41
                d_1970_cf41_ = d_1969___mcc_h14_
                d_1971_cf40_ = d_1968___mcc_h13_
                d_1972_cf39_ = d_1967___mcc_h12_
                d_1973_cf38_ = d_1966___mcc_h11_
                d_1974_v137_: _dafny.Set
                d_1974_v137_ = _dafny.Set({d_1973_cf38_, d_1973_cf38_, d_1787_v0_, d_1787_v0_, True})
                d_1921_v98_ = (default__.fm21(d_1972_cf39_, d_1970_cf41_, d_1974_v137_, globalState)) + (d_1921_v98_)
                d_1975_v138_: C1
                nw383_ = C1()
                nw383_.ctor__((self).f4, (self).f5)
                d_1975_v138_ = nw383_
                d_1975_v138_ = d_1975_v138_
                d_1787_v0_ = d_1787_v0_
                d_1973_cf38_ = (d_1973_cf38_) == (d_1787_v0_)
            elif source20_.is_DC24:
                d_1976___mcc_h15_ = source20_.cf37
                d_1977_cf37_ = d_1976___mcc_h15_
                d_1978_v139_: D16
                d_1978_v139_ = D16_DC38(d_1960_v133_, d_1787_v0_, (0) - (d_1788_v1_), d_1787_v0_)
                d_1979_v140_: _dafny.MultiSet
                d_1979_v140_ = _dafny.MultiSet([d_1788_v1_])
                pat_let_tv51_ = d_1921_v98_
                pat_let_tv52_ = d_1979_v140_
                pat_let_tv53_ = d_1921_v98_
                def iife159_(_pat_let51_0):
                    def iife160_(d_1980_dt__update__tmp_h2_):
                        def iife161_(_pat_let52_0):
                            def iife162_(d_1981_dt__update_hcf58_h0_):
                                return D16_DC38((d_1980_dt__update__tmp_h2_).cf57, d_1981_dt__update_hcf58_h0_, (d_1980_dt__update__tmp_h2_).cf59, (d_1980_dt__update__tmp_h2_).cf60)
                            return iife162_(_pat_let52_0)
                        return iife161_((pat_let_tv51_)[default__.safeIndex((pat_let_tv52_).cardinality, len(pat_let_tv53_))])
                    return iife160_(_pat_let51_0)
                d_1787_v0_ = (iife159_(d_1978_v139_)).cf58
                d_1982_v141_: _dafny.Array
                nw384_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_1982_v141_ = nw384_
                d_1982_v141_ = d_1982_v141_
                d_1983_v142_: str
                d_1983_v142_ = _dafny.CodePoint('l')
                d_1983_v142_ = d_1983_v142_
                d_1984_v143_: _dafny.Map
                d_1984_v143_ = _dafny.Map({default__.fm14(globalState): 666})
                d_1787_v0_ = ((len(d_1984_v143_)) * (default__.fm13(d_1787_v0_, True, d_1787_v0_, globalState))) == (d_1788_v1_)
            elif True:
                d_1985___mcc_h16_ = source20_.cf42
                d_1986_cf42_ = d_1985___mcc_h16_
                index372_ = default__.safeIndex(942, (d_1961_v134_).length(0))
                (d_1961_v134_)[index372_] = d_1788_v1_
                d_1987_v144_: _dafny.Map
                d_1987_v144_ = _dafny.Map({(self).f5: d_1788_v1_})
                index373_ = default__.safeIndex(942, (d_1961_v134_).length(0))
                (d_1961_v134_)[index373_] = ((d_1987_v144_)[(self).f5] if ((self).f5) in (d_1987_v144_) else default__.safeModuloInt(-855, len(d_1960_v133_)))
                d_1988_v145_: _dafny.Array
                def lambda128_(d_1989_v134_):
                    def lambda129_(d_1990_i10_):
                        return (_dafny.MultiSet([(d_1989_v134_)[default__.safeIndex(942, (d_1989_v134_).length(0))]])) | (_dafny.MultiSet([696, (d_1989_v134_)[default__.safeIndex(942, (d_1989_v134_).length(0))]]))

                    return lambda129_

                init72_ = lambda128_(d_1961_v134_)
                nw385_ = _dafny.Array(None, 23)
                for i0_72_ in range(nw385_.length(0)):
                    nw385_[i0_72_] = init72_(i0_72_)
                d_1988_v145_ = nw385_
                d_1988_v145_ = d_1988_v145_
                d_1991_v146_: _dafny.Set
                d_1991_v146_ = _dafny.Set({d_1788_v1_, len(_dafny.SeqWithoutIsStrInference([(d_1961_v134_)[default__.safeIndex(942, (d_1961_v134_).length(0))] for d_1992_i11_ in range(default__.abs(-678))])), default__.safeDivisionInt(d_1788_v1_, d_1788_v1_), (d_1961_v134_)[default__.safeIndex(942, (d_1961_v134_).length(0))]})
                d_1991_v146_ = d_1991_v146_
                d_1993_v147_: _dafny.Map
                d_1993_v147_ = _dafny.Map({d_1787_v0_: d_1787_v0_})
                d_1787_v0_ = ((d_1993_v147_)[(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1994_i12_ in range(default__.abs(724))])) == (d_1960_v133_)] if ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_1995_i12_ in range(default__.abs(724))])) == (d_1960_v133_)) in (d_1993_v147_) else d_1787_v0_)
            d_1996_v148_: _dafny.Array
            def lambda130_(d_1997_v1_, d_1998_v0_):
                def lambda131_(d_1999_i13_):
                    return (D14_DC32(d_1997_v1_, d_1998_v0_, d_1998_v0_)).cf50

                return lambda131_

            init73_ = lambda130_(d_1788_v1_, d_1787_v0_)
            nw386_ = _dafny.Array(None, 16)
            for i0_73_ in range(nw386_.length(0)):
                nw386_[i0_73_] = init73_(i0_73_)
            d_1996_v148_ = nw386_
            index374_ = default__.safeIndex(417, (d_1996_v148_).length(0))
            (d_1996_v148_)[index374_] = False
            index375_ = default__.safeIndex(417, (d_1996_v148_).length(0))
            (d_1996_v148_)[index375_] = (default__.fm14(globalState)) != (d_1788_v1_)
            d_2000_v149_: _dafny.Map
            d_2000_v149_ = _dafny.Map({d_1960_v133_: (d_1996_v148_)[default__.safeIndex(417, (d_1996_v148_).length(0))]})
            d_2000_v149_ = (d_2000_v149_).set(d_1960_v133_, (d_1996_v148_)[default__.safeIndex(417, (d_1996_v148_).length(0))])
            d_2001_v150_: _dafny.MultiSet
            d_2001_v150_ = _dafny.MultiSet([(d_1996_v148_)[default__.safeIndex(417, (d_1996_v148_).length(0))]])
            d_2002_v151_: _dafny.Map
            d_2002_v151_ = _dafny.Map({(d_2001_v150_).cardinality: d_1788_v1_})
            d_2003_v153_: C6
            nw387_ = C6()
            def iife163_():
                coll57_ = _dafny.Map()
                compr_57_: int
                for compr_57_ in _dafny.IntegerRange(221, 62):
                    d_2004_v152_: int = compr_57_
                    if ((221) <= (d_2004_v152_)) and ((d_2004_v152_) < (62)):
                        coll57_[(d_2004_v152_) + (len(d_1921_v98_))] = 6
                return _dafny.Map(coll57_)
            nw387_.ctor__((d_2002_v151_) | (iife163_()
            ))
            d_2003_v153_ = nw387_
        d_2005_v154_: _dafny.Array
        nw388_ = _dafny.Array(None, 4)
        nw388_[int(0)] = d_1787_v0_
        nw388_[int(1)] = d_1787_v0_
        nw388_[int(2)] = not(False)
        nw388_[int(3)] = d_1787_v0_
        d_2005_v154_ = nw388_
        d_2006_v155_: _dafny.Map
        d_2006_v155_ = _dafny.Map({d_1787_v0_: _dafny.SeqWithoutIsStrInference([d_1788_v1_ for d_2007_i14_ in range(default__.abs(327))])})
        d_2008_v156_: _dafny.Seq
        d_2008_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtcajlcn"))
        d_2009_v157_: _dafny.Map
        d_2009_v157_ = _dafny.Map({(self).fm3(d_1788_v1_, len(((d_2006_v155_)[d_1787_v0_] if (d_1787_v0_) in (d_2006_v155_) else d_1918_v95_)), globalState): len(d_2008_v156_)})
        d_2010_v158_: _dafny.Seq
        d_2010_v158_ = _dafny.SeqWithoutIsStrInference([d_2009_v157_, d_2009_v157_, d_2009_v157_, d_2009_v157_])
        d_2011_v159_: D17
        d_2011_v159_ = D17_DC42(d_2005_v154_, len(d_2010_v158_))
        source21_ = d_2011_v159_
        if source21_.is_DC41:
            d_2012___mcc_h17_ = source21_.cf63
            d_2013_cf63_ = d_2012___mcc_h17_
            d_2014_v160_: _dafny.Map
            d_2014_v160_ = _dafny.Map({(self).f5: d_2013_cf63_})
            d_2015_v161_: _dafny.MultiSet
            d_2015_v161_ = _dafny.MultiSet([len(_dafny.Set({d_2013_cf63_})), d_2013_cf63_])
            d_2016_v162_: _dafny.Array
            nw389_ = _dafny.Array(None, 19)
            nw389_[int(0)] = d_1788_v1_
            nw389_[int(1)] = d_1788_v1_
            nw389_[int(2)] = ((d_2014_v160_)[(self).f5] if ((self).f5) in (d_2014_v160_) else len(default__.fm16(False, (d_2015_v161_).cardinality, globalState)))
            nw389_[int(3)] = d_2013_cf63_
            nw389_[int(4)] = d_1788_v1_
            nw389_[int(5)] = d_2013_cf63_
            nw389_[int(6)] = d_2013_cf63_
            nw389_[int(7)] = ((d_2014_v160_)[_dafny.CodePoint('o')] if (_dafny.CodePoint('o')) in (d_2014_v160_) else d_2013_cf63_)
            nw389_[int(8)] = 712
            nw389_[int(9)] = d_2013_cf63_
            nw389_[int(10)] = d_1788_v1_
            nw389_[int(11)] = d_2013_cf63_
            nw389_[int(12)] = d_1788_v1_
            nw389_[int(13)] = -138
            nw389_[int(14)] = d_1788_v1_
            nw389_[int(15)] = (self).fm4(d_1918_v95_, globalState)
            nw389_[int(16)] = d_2013_cf63_
            nw389_[int(17)] = d_2013_cf63_
            nw389_[int(18)] = (0) - (d_1788_v1_)
            d_2016_v162_ = nw389_
            d_2017_v163_: _dafny.MultiSet
            d_2017_v163_ = _dafny.MultiSet([d_2016_v162_, d_2016_v162_, d_2016_v162_, d_2016_v162_, d_2016_v162_])
            d_2017_v163_ = d_2017_v163_
            d_2018_v165_: _dafny.Map
            def iife164_():
                coll58_ = _dafny.Set()
                compr_58_: int
                for compr_58_ in _dafny.IntegerRange(777, 78):
                    d_2019_v164_: int = compr_58_
                    if ((777) <= (d_2019_v164_)) and ((d_2019_v164_) < (78)):
                        coll58_ = coll58_.union(_dafny.Set([(d_2019_v164_) + (d_1788_v1_)]))
                return _dafny.Set(coll58_)
            d_2018_v165_ = _dafny.Map({d_1788_v1_: iife164_()
            })
            d_2020_v166_: _dafny.Set
            d_2020_v166_ = _dafny.Set({d_2013_cf63_})
            d_2021_v167_: _dafny.Map
            d_2021_v167_ = _dafny.Map({len(((d_2018_v165_)[d_2013_cf63_] if (d_2013_cf63_) in (d_2018_v165_) else d_2020_v166_)): d_2013_cf63_})
            d_2022_v168_: _dafny.Map
            d_2022_v168_ = _dafny.Map({d_2008_v156_: (d_2013_cf63_) in (d_2021_v167_)})
            index376_ = default__.safeIndex(192, (d_2005_v154_).length(0))
            (d_2005_v154_)[index376_] = not((d_1788_v1_) != (d_1788_v1_))
            d_2023_v169_: bool
            d_2023_v169_ = d_1787_v0_
            index377_ = default__.safeIndex(192, (d_2005_v154_).length(0))
            rhs366_ = d_2022_v168_
            rhs367_ = (d_2023_v169_)
            lhs179_ = d_2005_v154_
            lhs180_ = default__.safeIndex(192, (d_2005_v154_).length(0))
            d_2022_v168_ = rhs366_
            lhs179_[lhs180_] = rhs367_
            if ((self).f5) in ((d_2008_v156_) + (d_2008_v156_)):
                d_1788_v1_ = (d_2013_cf63_) + (-47)
                r1 = d_2013_cf63_
                d_1788_v1_ = (d_2013_cf63_) - ((d_1918_v95_)[default__.safeIndex(len(d_1921_v98_), len(d_1918_v95_))])
                d_2024_v170_: _dafny.Seq
                d_2024_v170_ = _dafny.SeqWithoutIsStrInference([d_2008_v156_])
                d_2024_v170_ = (d_2024_v170_) + (d_2024_v170_)
                index378_ = default__.safeIndex(192, (d_2005_v154_).length(0))
                (d_2005_v154_)[index378_] = ((d_2015_v161_) | (d_2015_v161_)).isdisjoint(default__.fm30((d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))], d_1921_v98_, globalState))
            elif True:
                d_1788_v1_ = d_2013_cf63_
                d_2025_v171_: _dafny.Map
                d_2025_v171_ = _dafny.Map({d_1788_v1_: d_1787_v0_})
                index379_ = default__.safeIndex(192, (d_2005_v154_).length(0))
                (d_2005_v154_)[index379_] = (len((d_2025_v171_ if d_1787_v0_ else d_2025_v171_))) != (d_2013_cf63_)
                r1 = default__.fm13(d_1787_v0_, (d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))], (d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))], globalState)
                d_2026_v173_: _dafny.Map
                def iife165_():
                    coll59_ = _dafny.Map()
                    compr_59_: str
                    for compr_59_ in (_dafny.Map({(self).f5: (_dafny.Map({(d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))]: d_2013_cf63_})).set(False, d_2013_cf63_)})).keys.Elements:
                        d_2027_v172_: str = compr_59_
                        if (d_2027_v172_) in (_dafny.Map({(self).f5: (_dafny.Map({(d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))]: d_2013_cf63_})).set(False, d_2013_cf63_)})):
                            coll59_[d_2027_v172_] = D8_DC19(d_1788_v1_, (d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))])
                    return _dafny.Map(coll59_)
                d_2026_v173_ = _dafny.Map({(d_2008_v156_) + (d_2008_v156_): iife165_()
                })
                d_2028_v174_: D8
                d_2028_v174_ = D8_DC19(d_2013_cf63_, False)
                d_2029_v175_: _dafny.Map
                d_2029_v175_ = _dafny.Map({(self).f5: d_2028_v174_})
                d_2026_v173_ = (d_2026_v173_).set((d_2008_v156_) + (d_2008_v156_), d_2029_v175_)
                d_2030_v176_: _dafny.Array
                nw390_ = _dafny.Array(_dafny.Set({}), 21)
                d_2030_v176_ = nw390_
                d_2030_v176_ = (self).f4
            index380_ = default__.safeIndex(944, (d_2016_v162_).length(0))
            (d_2016_v162_)[index380_] = d_2013_cf63_
            d_2031_v177_: _dafny.Array
            def lambda132_(d_2032_i15_):
                return D7_DC16((self).f5)

            init74_ = lambda132_
            nw391_ = _dafny.Array(None, 18)
            for i0_74_ in range(nw391_.length(0)):
                nw391_[i0_74_] = init74_(i0_74_)
            d_2031_v177_ = nw391_
            d_2033_v178_: D7
            d_2033_v178_ = D7_DC16((self).f5)
            index381_ = default__.safeIndex(409, (d_2031_v177_).length(0))
            (d_2031_v177_)[index381_] = d_2033_v178_
            index382_ = default__.safeIndex(192, (d_2005_v154_).length(0))
            index383_ = default__.safeIndex(944, (d_2016_v162_).length(0))
            index384_ = default__.safeIndex(409, (d_2031_v177_).length(0))
            rhs368_ = d_1787_v0_
            rhs369_ = d_2013_cf63_
            rhs370_ = d_2013_cf63_
            rhs371_ = (D7_DC16((self).f5) if (d_2005_v154_)[default__.safeIndex(192, (d_2005_v154_).length(0))] else d_2033_v178_)
            lhs181_ = d_2005_v154_
            lhs182_ = default__.safeIndex(192, (d_2005_v154_).length(0))
            lhs183_ = d_2016_v162_
            lhs184_ = default__.safeIndex(944, (d_2016_v162_).length(0))
            lhs185_ = d_2031_v177_
            lhs186_ = default__.safeIndex(409, (d_2031_v177_).length(0))
            lhs181_[lhs182_] = rhs368_
            r1 = rhs369_
            lhs183_[lhs184_] = rhs370_
            lhs185_[lhs186_] = rhs371_
        elif source21_.is_DC42:
            d_2034___mcc_h18_ = source21_.cf64
            d_2035___mcc_h19_ = source21_.cf65
            d_2036_cf65_ = d_2035___mcc_h19_
            d_2037_cf64_ = d_2034___mcc_h18_
            d_2038_v179_: _dafny.Array
            nw392_ = _dafny.Array(int(0), 22)
            d_2038_v179_ = nw392_
            index385_ = default__.safeIndex(68, (d_2038_v179_).length(0))
            (d_2038_v179_)[index385_] = d_1788_v1_
            index386_ = default__.safeIndex(68, (d_2038_v179_).length(0))
            (d_2038_v179_)[index386_] = d_1788_v1_
            d_2039_v180_: C0
            nw393_ = C0()
            nw393_.ctor__()
            d_2039_v180_ = nw393_
            d_2008_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srfmniqdo"))
            d_2040_v181_: _dafny.Array
            def lambda133_(d_2041_v95_):
                def lambda134_(d_2042_i16_):
                    return d_2041_v95_

                return lambda134_

            init75_ = lambda133_(d_1918_v95_)
            nw394_ = _dafny.Array(None, 26)
            for i0_75_ in range(nw394_.length(0)):
                nw394_[i0_75_] = init75_(i0_75_)
            d_2040_v181_ = nw394_
            d_2043_v182_: _dafny.Map
            d_2043_v182_ = _dafny.Map({d_2009_v157_: d_2040_v181_})
            d_2044_v183_: _dafny.Map
            d_2044_v183_ = _dafny.Map({d_1787_v0_: d_1787_v0_})
            d_2045_v184_: _dafny.Array
            nw395_ = _dafny.Array(None, 29)
            nw395_[int(0)] = d_2040_v181_
            nw395_[int(1)] = d_2040_v181_
            nw395_[int(2)] = d_2040_v181_
            nw395_[int(3)] = ((d_2043_v182_)[d_2009_v157_] if (d_2009_v157_) in (d_2043_v182_) else d_2040_v181_)
            nw395_[int(4)] = d_2040_v181_
            nw395_[int(5)] = d_2040_v181_
            nw395_[int(6)] = d_2040_v181_
            nw395_[int(7)] = d_2040_v181_
            nw395_[int(8)] = d_2040_v181_
            nw395_[int(9)] = d_2040_v181_
            nw395_[int(10)] = d_2040_v181_
            nw395_[int(11)] = d_2040_v181_
            nw395_[int(12)] = d_2040_v181_
            nw395_[int(13)] = d_2040_v181_
            nw395_[int(14)] = d_2040_v181_
            nw395_[int(15)] = d_2040_v181_
            nw395_[int(16)] = (d_2040_v181_ if d_1787_v0_ else d_2040_v181_)
            nw395_[int(17)] = (d_2040_v181_ if d_1787_v0_ else d_2040_v181_)
            nw395_[int(18)] = d_2040_v181_
            nw395_[int(19)] = d_2040_v181_
            nw395_[int(20)] = d_2040_v181_
            nw395_[int(21)] = d_2040_v181_
            nw395_[int(22)] = (d_2040_v181_ if d_1787_v0_ else d_2040_v181_)
            nw395_[int(23)] = d_2040_v181_
            nw395_[int(24)] = d_2040_v181_
            nw395_[int(25)] = d_2040_v181_
            nw395_[int(26)] = d_2040_v181_
            nw395_[int(27)] = (d_2040_v181_ if ((d_2044_v183_)[not(not(default__.fm1(866, d_1788_v1_, d_1787_v0_, d_1787_v0_, globalState)))] if (not(not(default__.fm1(866, d_1788_v1_, d_1787_v0_, d_1787_v0_, globalState)))) in (d_2044_v183_) else d_1787_v0_) else d_2040_v181_)
            nw395_[int(28)] = d_2040_v181_
            d_2045_v184_ = nw395_
            index387_ = default__.safeIndex(869, (d_2045_v184_).length(0))
            (d_2045_v184_)[index387_] = d_2040_v181_
            d_2046_v185_: _dafny.MultiSet
            d_2046_v185_ = _dafny.MultiSet([d_1787_v0_])
            index388_ = default__.safeIndex(68, (d_2038_v179_).length(0))
            index389_ = default__.safeIndex(869, (d_2045_v184_).length(0))
            rhs372_ = not (d_1787_v0_) or (d_1787_v0_)
            rhs373_ = d_2036_cf65_
            rhs374_ = (default__.safeModuloInt((d_2038_v179_)[default__.safeIndex(68, (d_2038_v179_).length(0))], d_2036_cf65_)) * ((d_2046_v185_).cardinality)
            rhs375_ = d_2040_v181_
            rhs376_ = d_1787_v0_
            lhs187_ = d_2038_v179_
            lhs188_ = default__.safeIndex(68, (d_2038_v179_).length(0))
            lhs189_ = d_2045_v184_
            lhs190_ = default__.safeIndex(869, (d_2045_v184_).length(0))
            d_1787_v0_ = rhs372_
            lhs187_[lhs188_] = rhs373_
            r1 = rhs374_
            lhs189_[lhs190_] = rhs375_
            d_1787_v0_ = rhs376_
        elif source21_.is_DC40:
            d_2047___mcc_h20_ = source21_.cf62
            d_2048_cf62_ = d_2047___mcc_h20_
            d_2049_v186_: C0
            nw396_ = C0()
            nw396_.ctor__()
            d_2049_v186_ = nw396_
            d_1787_v0_ = ((0) - (d_1788_v1_)) > ((0) - (d_1788_v1_))
            if d_1787_v0_:
                d_2008_v156_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tulyymy"))) + ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_2050_i17_ in range(default__.abs(-503))])) + (d_2008_v156_))
                d_2051_v187_: _dafny.Map
                d_2051_v187_ = _dafny.Map({(self).f5: len(d_1918_v95_)})
                d_2051_v187_ = (d_2051_v187_).set((self).f5, (0) - (d_1788_v1_))
                d_2052_v188_: C5
                nw397_ = C5()
                nw397_.ctor__()
                d_2052_v188_ = nw397_
                d_2053_v189_: C6
                nw398_ = C6()
                nw398_.ctor__(_dafny.Map({d_1788_v1_: d_1788_v1_}))
                d_2053_v189_ = nw398_
                d_2054_v190_: C7
                nw399_ = C7()
                nw399_.ctor__()
                d_2054_v190_ = nw399_
                d_2055_v191_: _dafny.Map
                d_2055_v191_ = _dafny.Map({d_1787_v0_: d_2054_v190_})
                d_2055_v191_ = (d_2055_v191_).set(d_1787_v0_, d_2054_v190_)
            elif True:
                d_1788_v1_ = d_1788_v1_
                d_2056_v192_: _dafny.Array
                nw400_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_2056_v192_ = nw400_
                nw401_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_2056_v192_ = nw401_
                d_2057_v193_: _dafny.Map
                d_2057_v193_ = _dafny.Map({d_1787_v0_: d_2008_v156_})
                d_2057_v193_ = (d_2057_v193_).set(d_1787_v0_, d_2008_v156_)
                r1 = default__.safeModuloInt(194, (d_1788_v1_) * (d_1788_v1_))
                d_1788_v1_ = default__.safeModuloInt(d_1788_v1_, d_1788_v1_)
            d_1787_v0_ = False
        elif True:
            d_2058___mcc_h21_ = source21_.cf66
            d_2059_cf66_ = d_2058___mcc_h21_
            d_1921_v98_ = (d_1921_v98_) + (d_1921_v98_)
            (globalState).f3 = d_2005_v154_
            d_2060_v194_: _dafny.Array
            nw402_ = _dafny.Array(None, 3)
            d_2060_v194_ = nw402_
            d_2061_v195_: _dafny.Seq
            d_2061_v195_ = _dafny.SeqWithoutIsStrInference([d_2060_v194_])
            d_2062_v196_: _dafny.Array
            nw403_ = _dafny.Array(None, 5)
            nw403_[int(0)] = (d_2061_v195_) + (d_2061_v195_)
            nw403_[int(1)] = (d_2061_v195_) + (d_2061_v195_)
            nw403_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_2060_v194_, d_2060_v194_])) + (d_2061_v195_)
            nw403_[int(3)] = d_2061_v195_
            nw403_[int(4)] = d_2061_v195_
            d_2062_v196_ = nw403_
            index390_ = default__.safeIndex(827, (d_2062_v196_).length(0))
            (d_2062_v196_)[index390_] = (d_2061_v195_ if d_1787_v0_ else d_2061_v195_)
            index391_ = default__.safeIndex(827, (d_2062_v196_).length(0))
            rhs377_ = (d_1918_v95_) + (d_1918_v95_)
            rhs378_ = d_2061_v195_
            lhs191_ = globalState
            lhs192_ = d_2062_v196_
            lhs193_ = default__.safeIndex(827, (d_2062_v196_).length(0))
            lhs191_.f2 = rhs377_
            lhs192_[lhs193_] = rhs378_
            rhs379_ = not(not (d_1787_v0_) or (not(not(d_1787_v0_))))
            rhs380_ = d_1788_v1_
            d_1787_v0_ = rhs379_
            d_1788_v1_ = rhs380_
        if (self).fm5(d_1788_v1_, d_1787_v0_, d_1788_v1_, d_1788_v1_, globalState):
            d_2063_v197_: _dafny.Array
            nw404_ = _dafny.Array(None, 4)
            nw404_[int(0)] = d_2005_v154_
            nw404_[int(1)] = d_2005_v154_
            nw404_[int(2)] = d_2005_v154_
            nw404_[int(3)] = d_2005_v154_
            d_2063_v197_ = nw404_
            rhs381_ = len(d_2008_v156_)
            rhs382_ = d_2063_v197_
            rhs383_ = (not(not(True))) or (((self).f5) in (d_2008_v156_))
            rhs384_ = not(d_1787_v0_)
            d_1788_v1_ = rhs381_
            d_2063_v197_ = rhs382_
            d_1787_v0_ = rhs383_
            d_1787_v0_ = rhs384_
            d_1787_v0_ = (self).fm3(len(d_2008_v156_), len((d_1921_v98_).set(default__.safeIndex((0) - (d_1788_v1_), len(d_1921_v98_)), d_1787_v0_)), globalState)
            d_2064_v198_: _dafny.Set
            d_2064_v198_ = _dafny.Set({d_1787_v0_, d_1787_v0_})
            d_2065_v199_: _dafny.Seq
            d_2065_v199_ = _dafny.SeqWithoutIsStrInference([d_2064_v198_])
            d_1788_v1_ = default__.safeModuloInt(len((d_2065_v199_)[default__.safeIndex(d_1788_v1_, len(d_2065_v199_))]), d_1788_v1_)
            d_2066_v200_: _dafny.Array
            nw405_ = _dafny.Array(_dafny.CodePoint('D'), 14)
            d_2066_v200_ = nw405_
            d_2067_v201_: _dafny.Map
            d_2067_v201_ = _dafny.Map({True: d_2066_v200_})
            d_2068_v202_: T1
            nw406_ = C3()
            nw406_.ctor__((self).f4, (self).f5)
            d_2068_v202_ = nw406_
            rhs385_ = d_1787_v0_
            rhs386_ = (d_2067_v201_) | (d_2067_v201_)
            rhs387_ = d_2068_v202_
            d_1787_v0_ = rhs385_
            d_2067_v201_ = rhs386_
            d_2068_v202_ = rhs387_
            d_1788_v1_ = len(d_1921_v98_)
        elif True:
            d_2069_v203_: _dafny.Array
            nw407_ = _dafny.Array(int(0), 26)
            d_2069_v203_ = nw407_
            index392_ = default__.safeIndex(919, (d_2069_v203_).length(0))
            (d_2069_v203_)[index392_] = d_1788_v1_
            d_2070_v204_: _dafny.Set
            d_2070_v204_ = _dafny.Set({d_1788_v1_, d_1788_v1_, d_1788_v1_})
            d_2071_v205_: C8
            nw408_ = C8()
            nw408_.ctor__()
            d_2071_v205_ = nw408_
            d_2072_v206_: _dafny.Map
            d_2072_v206_ = _dafny.Map({(d_2070_v204_) - (d_2070_v204_): d_2071_v205_})
            d_2073_v207_: _dafny.Map
            d_2073_v207_ = _dafny.Map({d_1788_v1_: d_1787_v0_})
            d_2074_v208_: _dafny.Seq
            d_2074_v208_ = _dafny.SeqWithoutIsStrInference([(d_2073_v207_) | (d_2073_v207_), (d_2073_v207_) | (d_2073_v207_)])
            d_2075_v209_: D7
            d_2075_v209_ = D7_DC17(d_1787_v0_, d_1787_v0_, not(d_1787_v0_), d_1787_v0_)
            d_2076_v210_: _dafny.Map
            d_2076_v210_ = _dafny.Map({default__.fm13((d_2075_v209_).cf25, d_1787_v0_, not(d_1787_v0_), globalState): len(d_2008_v156_)})
            index393_ = default__.safeIndex(919, (d_2069_v203_).length(0))
            rhs388_ = len(d_2074_v208_)
            rhs389_ = d_1787_v0_
            rhs390_ = d_2072_v206_
            rhs391_ = default__.safeDivisionInt(default__.safeModuloInt(d_1788_v1_, ((d_2076_v210_)[d_1788_v1_] if (d_1788_v1_) in (d_2076_v210_) else d_1788_v1_)), d_1788_v1_)
            rhs392_ = (len(d_1918_v95_)) - (default__.safeModuloInt(len(d_2073_v207_), d_1788_v1_))
            lhs194_ = d_2069_v203_
            lhs195_ = default__.safeIndex(919, (d_2069_v203_).length(0))
            lhs194_[lhs195_] = rhs388_
            d_1787_v0_ = rhs389_
            d_2072_v206_ = rhs390_
            d_1788_v1_ = rhs391_
            r1 = rhs392_
            rhs393_ = d_1787_v0_
            rhs394_ = (d_1787_v0_ if d_1787_v0_ else True)
            d_1787_v0_ = rhs393_
            d_1787_v0_ = rhs394_
            index394_ = default__.safeIndex(801, (d_2005_v154_).length(0))
            (d_2005_v154_)[index394_] = d_1787_v0_
            d_2077_v211_: _dafny.MultiSet
            d_2077_v211_ = _dafny.MultiSet([d_1787_v0_])
            index395_ = default__.safeIndex(801, (d_2005_v154_).length(0))
            (d_2005_v154_)[index395_] = (d_2077_v211_) == (d_2077_v211_)
            d_2078_v212_: T0
            nw409_ = C7()
            nw409_.ctor__()
            d_2078_v212_ = nw409_
            d_2078_v212_ = d_2078_v212_
            index396_ = default__.safeIndex(801, (d_2005_v154_).length(0))
            (d_2005_v154_)[index396_] = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2079_i18_ in range(default__.abs(-532))])) == ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_2080_i19_ in range(default__.abs(289))])) + (d_2008_v156_))
        r0 = d_1918_v95_
        r1 = d_1788_v1_
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_2081_v0_: bool
        d_2081_v0_ = True
        d_2081_v0_ = d_2081_v0_
        d_2082_v1_: int
        d_2082_v1_ = 376
        d_2083_v2_: _dafny.Map
        d_2083_v2_ = _dafny.Map({d_2081_v0_: (self).f5})
        d_2084_v3_: _dafny.Map
        d_2084_v3_ = _dafny.Map({(self).f5: (0) - (d_2082_v1_)})
        d_2085_v4_: _dafny.Seq
        d_2085_v4_ = _dafny.SeqWithoutIsStrInference([(d_2082_v1_ if d_2081_v0_ else d_2082_v1_), (0) - (d_2082_v1_), (d_2082_v1_) + (len((d_2083_v2_).set(not(d_2081_v0_), (self).f5))), default__.safeDivisionInt(((d_2084_v3_)[(self).f5] if ((self).f5) in (d_2084_v3_) else 966), d_2082_v1_)])
        d_2082_v1_ = (d_2085_v4_)[default__.safeIndex(default__.safeModuloInt(default__.fm13(d_2081_v0_, d_2081_v0_, d_2081_v0_, globalState), 857), len(d_2085_v4_))]
        d_2086_v5_: _dafny.Seq
        d_2086_v5_ = _dafny.SeqWithoutIsStrInference([d_2081_v0_])
        hi17_ = len(d_2086_v5_)
        for d_2087_i0_ in range(((0) - (d_2082_v1_)) - (d_2082_v1_), hi17_):
            d_2088_v6_: _dafny.Seq
            d_2088_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hycyid"))
            d_2082_v1_ = (0) - (len(((d_2088_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snftantu")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfkcaacp")))))
            if d_2081_v0_:
                d_2089_v7_: _dafny.Array
                nw410_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_2089_v7_ = nw410_
                index397_ = default__.safeIndex(760, (d_2089_v7_).length(0))
                (d_2089_v7_)[index397_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udrapuaxq"))
                index398_ = default__.safeIndex(760, (d_2089_v7_).length(0))
                (d_2089_v7_)[index398_] = d_2088_v6_
                d_2081_v0_ = d_2081_v0_
                d_2081_v0_ = d_2081_v0_
                d_2090_v8_: _dafny.Map
                d_2090_v8_ = _dafny.Map({d_2081_v0_: d_2082_v1_})
                d_2091_v9_: _dafny.Set
                d_2091_v9_ = _dafny.Set({d_2081_v0_, not(d_2081_v0_), True})
                d_2092_v10_: _dafny.MultiSet
                d_2092_v10_ = _dafny.MultiSet([d_2081_v0_])
                d_2093_v11_: _dafny.Array
                nw411_ = _dafny.Array(None, 13)
                nw411_[int(0)] = d_2087_i0_
                nw411_[int(1)] = d_2087_i0_
                nw411_[int(2)] = d_2087_i0_
                nw411_[int(3)] = d_2087_i0_
                nw411_[int(4)] = len((d_2089_v7_)[default__.safeIndex(760, (d_2089_v7_).length(0))])
                nw411_[int(5)] = d_2087_i0_
                nw411_[int(6)] = d_2082_v1_
                nw411_[int(7)] = d_2082_v1_
                nw411_[int(8)] = d_2087_i0_
                nw411_[int(9)] = d_2082_v1_
                nw411_[int(10)] = d_2087_i0_
                nw411_[int(11)] = d_2087_i0_
                nw411_[int(12)] = len(d_2088_v6_)
                d_2093_v11_ = nw411_
                d_2094_v12_: _dafny.Set
                d_2094_v12_ = _dafny.Set({len(d_2085_v4_)})
                d_2095_v13_: _dafny.Array
                nw412_ = _dafny.Array(None, 16)
                nw412_[int(0)] = d_2087_i0_
                nw412_[int(1)] = len((d_2089_v7_)[default__.safeIndex(760, (d_2089_v7_).length(0))])
                nw412_[int(2)] = 800
                nw412_[int(3)] = d_2082_v1_
                nw412_[int(4)] = d_2082_v1_
                nw412_[int(5)] = len(d_2090_v8_)
                nw412_[int(6)] = d_2087_i0_
                nw412_[int(7)] = len(d_2091_v9_)
                nw412_[int(8)] = (d_2082_v1_) * (d_2082_v1_)
                nw412_[int(9)] = 848
                nw412_[int(10)] = d_2082_v1_
                nw412_[int(11)] = d_2087_i0_
                nw412_[int(12)] = d_2082_v1_
                nw412_[int(13)] = default__.safeModuloInt((d_2092_v10_).cardinality, (D11_DC25(d_2081_v0_, d_2082_v1_, d_2093_v11_, d_2082_v1_)).cf41)
                nw412_[int(14)] = (0) - (len(d_2094_v12_))
                nw412_[int(15)] = (0) - ((d_2082_v1_) * (len(d_2086_v5_)))
                d_2095_v13_ = nw412_
                index399_ = default__.safeIndex(922, (d_2093_v11_).length(0))
                (d_2093_v11_)[index399_] = len(_dafny.Map({_dafny.Set({d_2087_i0_}): d_2081_v0_}))
                index400_ = default__.safeIndex(922, (d_2093_v11_).length(0))
                (d_2093_v11_)[index400_] = d_2082_v1_
                default__.m0(d_2082_v1_, (d_2082_v1_) > (d_2082_v1_), globalState)
            elif True:
                d_2082_v1_ = d_2087_i0_
                d_2096_v14_: D8
                d_2096_v14_ = D8_DC19(d_2087_i0_, True)
                d_2097_v15_: _dafny.Map
                d_2097_v15_ = _dafny.Map({d_2087_i0_: d_2082_v1_})
                d_2082_v1_ = default__.safeModuloInt((d_2096_v14_).cf29, len(d_2097_v15_))
                d_2098_v16_: _dafny.Map
                d_2098_v16_ = _dafny.Map({d_2082_v1_: d_2081_v0_})
                d_2099_v17_: D15
                d_2099_v17_ = D15_DC35(d_2082_v1_)
                rhs395_ = (len(d_2085_v4_)) + (d_2087_i0_)
                rhs396_ = ((d_2098_v16_)[(d_2099_v17_).cf54] if ((d_2099_v17_).cf54) in (d_2098_v16_) else d_2081_v0_)
                d_2082_v1_ = rhs395_
                d_2081_v0_ = rhs396_
                d_2081_v0_ = d_2081_v0_
                d_2082_v1_ = len(d_2088_v6_)
            d_2100_v18_: _dafny.Map
            d_2100_v18_ = _dafny.Map({d_2081_v0_: default__.fm14(globalState)})
            d_2100_v18_ = (d_2100_v18_).set(d_2081_v0_, d_2087_i0_)
            d_2082_v1_ = (d_2087_i0_ if (d_2087_i0_) >= (d_2087_i0_) else (0) - ((default__.fm14(globalState)) + (408)))
        if (-362) == (d_2082_v1_):
            d_2082_v1_ = ((d_2085_v4_)[default__.safeIndex(d_2082_v1_, len(d_2085_v4_))]) - (d_2082_v1_)
            d_2082_v1_ = d_2082_v1_
            d_2082_v1_ = (d_2082_v1_ if d_2081_v0_ else default__.safeDivisionInt(d_2082_v1_, d_2082_v1_))
            d_2101_v19_: C0
            nw413_ = C0()
            nw413_.ctor__()
            d_2101_v19_ = nw413_
            d_2101_v19_ = d_2101_v19_
            d_2081_v0_ = d_2081_v0_
        elif True:
            d_2102_v20_: C0
            nw414_ = C0()
            nw414_.ctor__()
            d_2102_v20_ = nw414_
            d_2103_v21_: C2
            nw415_ = C2()
            nw415_.ctor__((self).f4, (self).f5)
            d_2103_v21_ = nw415_
            d_2103_v21_ = d_2103_v21_
            d_2104_v22_: _dafny.Map
            d_2104_v22_ = _dafny.Map({d_2082_v1_: d_2102_v20_})
            d_2105_v23_: _dafny.Map
            d_2105_v23_ = _dafny.Map({((d_2104_v22_)[d_2082_v1_] if (d_2082_v1_) in (d_2104_v22_) else d_2102_v20_): 630})
            d_2082_v1_ = default__.safeModuloInt(d_2082_v1_, len(d_2105_v23_))
            d_2106_v24_: C7
            nw416_ = C7()
            nw416_.ctor__()
            d_2106_v24_ = nw416_
            if (d_2081_v0_) and (d_2081_v0_):
                d_2107_v25_: _dafny.Array
                def lambda135_(d_2108_i1_):
                    return (self).f5

                init76_ = lambda135_
                nw417_ = _dafny.Array(None, 20)
                for i0_76_ in range(nw417_.length(0)):
                    nw417_[i0_76_] = init76_(i0_76_)
                d_2107_v25_ = nw417_
                index401_ = default__.safeIndex(947, (d_2107_v25_).length(0))
                (d_2107_v25_)[index401_] = (self).f5
                d_2109_v26_: _dafny.Seq
                d_2109_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehiaxp"))
                index402_ = default__.safeIndex(947, (d_2107_v25_).length(0))
                rhs397_ = (self).f5
                rhs398_ = (not(d_2081_v0_) if (len(d_2109_v26_)) == (d_2082_v1_) else not((self).fm3(d_2082_v1_, (0) - (d_2082_v1_), globalState)))
                rhs399_ = (d_2106_v24_).fm3((d_2082_v1_ if True else d_2082_v1_), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnxs")))) - (d_2082_v1_), globalState)
                lhs196_ = d_2107_v25_
                lhs197_ = default__.safeIndex(947, (d_2107_v25_).length(0))
                lhs196_[lhs197_] = rhs397_
                d_2081_v0_ = rhs398_
                d_2081_v0_ = rhs399_
                d_2110_v27_: _dafny.Array
                nw418_ = _dafny.Array(False, 12)
                d_2110_v27_ = nw418_
                index403_ = default__.safeIndex(248, (d_2110_v27_).length(0))
                (d_2110_v27_)[index403_] = (d_2082_v1_) <= (d_2082_v1_)
                index404_ = default__.safeIndex(248, (d_2110_v27_).length(0))
                (d_2110_v27_)[index404_] = (d_2106_v24_).fm3(len(d_2109_v26_), (0) - (d_2082_v1_), globalState)
                d_2111_v28_: _dafny.Array
                nw419_ = _dafny.Array(int(0), 26)
                d_2111_v28_ = nw419_
                index405_ = default__.safeIndex(377, (d_2111_v28_).length(0))
                (d_2111_v28_)[index405_] = d_2082_v1_
                d_2112_v29_: _dafny.Map
                d_2112_v29_ = _dafny.Map({d_2110_v27_: d_2109_v26_})
                d_2113_v30_: _dafny.Map
                d_2113_v30_ = _dafny.Map({(self).fm3(d_2082_v1_, 304, globalState): _dafny.SeqWithoutIsStrInference([(self).f5 for d_2114_i2_ in range(default__.abs(762))])})
                d_2115_v31_: _dafny.Seq
                d_2115_v31_ = _dafny.SeqWithoutIsStrInference([(((d_2112_v29_)[d_2110_v27_] if (d_2110_v27_) in (d_2112_v29_) else ((d_2113_v30_)[(d_2110_v27_)[default__.safeIndex(248, (d_2110_v27_).length(0))]] if ((d_2110_v27_)[default__.safeIndex(248, (d_2110_v27_).length(0))]) in (d_2113_v30_) else d_2109_v26_))) + (d_2109_v26_), d_2109_v26_])
                d_2116_v32_: _dafny.MultiSet
                d_2116_v32_ = _dafny.MultiSet([(d_2110_v27_)[default__.safeIndex(248, (d_2110_v27_).length(0))]])
                index406_ = default__.safeIndex(377, (d_2111_v28_).length(0))
                index407_ = default__.safeIndex(248, (d_2110_v27_).length(0))
                rhs400_ = ((default__.fm16(False, d_2082_v1_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsnh")))) + (d_2109_v26_)
                rhs401_ = len((d_2115_v31_)[default__.safeIndex(((d_2116_v32_).cardinality) - (d_2082_v1_), len(d_2115_v31_))])
                rhs402_ = d_2082_v1_
                rhs403_ = (d_2086_v5_) == ((d_2086_v5_) + (d_2086_v5_))
                rhs404_ = True
                lhs198_ = d_2111_v28_
                lhs199_ = default__.safeIndex(377, (d_2111_v28_).length(0))
                lhs200_ = d_2110_v27_
                lhs201_ = default__.safeIndex(248, (d_2110_v27_).length(0))
                d_2109_v26_ = rhs400_
                d_2082_v1_ = rhs401_
                lhs198_[lhs199_] = rhs402_
                d_2081_v0_ = rhs403_
                lhs200_[lhs201_] = rhs404_
                d_2109_v26_ = (d_2109_v26_) + (((d_2109_v26_).set(default__.safeIndex(d_2082_v1_, len(d_2109_v26_)), (d_2107_v25_)[default__.safeIndex(947, (d_2107_v25_).length(0))]) if (d_2110_v27_)[default__.safeIndex(248, (d_2110_v27_).length(0))] else d_2109_v26_))
                d_2081_v0_ = (d_2110_v27_)[default__.safeIndex(248, (d_2110_v27_).length(0))]
            elif True:
                d_2081_v0_ = d_2081_v0_
                d_2117_v33_: _dafny.MultiSet
                d_2117_v33_ = _dafny.MultiSet([(self).f5, _dafny.CodePoint('x')])
                d_2118_v34_: _dafny.Set
                d_2118_v34_ = _dafny.Set({d_2082_v1_, ((d_2117_v33_).intersection(d_2117_v33_)).cardinality})
                d_2118_v34_ = d_2118_v34_
                d_2119_v35_: _dafny.Array
                def lambda136_(d_2120_v0_):
                    def lambda137_(d_2121_i3_):
                        return _dafny.Map({d_2120_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhbmkh"))})

                    return lambda137_

                init77_ = lambda136_(d_2081_v0_)
                nw420_ = _dafny.Array(None, 16)
                for i0_77_ in range(nw420_.length(0)):
                    nw420_[i0_77_] = init77_(i0_77_)
                d_2119_v35_ = nw420_
                d_2122_v36_: _dafny.Seq
                d_2122_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
                d_2123_v37_: _dafny.Map
                d_2123_v37_ = _dafny.Map({d_2081_v0_: d_2122_v36_})
                index408_ = default__.safeIndex(855, (d_2119_v35_).length(0))
                (d_2119_v35_)[index408_] = d_2123_v37_
                index409_ = default__.safeIndex(855, (d_2119_v35_).length(0))
                (d_2119_v35_)[index409_] = _dafny.Map({d_2081_v0_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap"))).set(default__.safeIndex(d_2082_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap")))), (self).f5)})
                d_2124_v38_: str
                d_2124_v38_ = _dafny.CodePoint('l')
                d_2124_v38_ = (self).f5
                rhs405_ = d_2122_v36_
                rhs406_ = d_2081_v0_
                d_2122_v36_ = rhs405_
                d_2081_v0_ = rhs406_
        d_2125_v39_: _dafny.MultiSet
        d_2125_v39_ = _dafny.MultiSet([d_2081_v0_, d_2081_v0_])
        hi18_ = d_2082_v1_
        for d_2126_i4_ in range((0) - ((((d_2125_v39_).set(d_2081_v0_, default__.abs(d_2082_v1_))).cardinality) - (d_2082_v1_)), hi18_):
            d_2081_v0_ = (d_2126_i4_) == (d_2082_v1_)
            d_2127_v40_: _dafny.Map
            d_2127_v40_ = _dafny.Map({default__.fm43(d_2081_v0_, _dafny.MultiSet([d_2082_v1_, (0) - (d_2082_v1_), d_2082_v1_, default__.fm13(d_2081_v0_, d_2081_v0_, d_2081_v0_, globalState)]), globalState): d_2081_v0_})
            d_2128_v41_: _dafny.Array
            nw421_ = _dafny.Array(int(0), 3)
            d_2128_v41_ = nw421_
            d_2129_v42_: D11
            d_2129_v42_ = D11_DC25(d_2081_v0_, d_2126_i4_, d_2128_v41_, len(_dafny.SeqWithoutIsStrInference([d_2082_v1_])))
            d_2130_v43_: _dafny.Set
            d_2130_v43_ = _dafny.Set({d_2125_v39_, (d_2125_v39_).set(d_2081_v0_, default__.abs((d_2129_v42_).cf41)), d_2125_v39_})
            d_2131_v44_: _dafny.Seq
            d_2131_v44_ = _dafny.SeqWithoutIsStrInference([d_2130_v43_])
            d_2132_v45_: _dafny.Seq
            d_2132_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auc"))
            d_2081_v0_ = ((d_2127_v40_)[((d_2131_v44_)[default__.safeIndex(618, len(d_2131_v44_))]) | (d_2130_v43_)] if (((d_2131_v44_)[default__.safeIndex(618, len(d_2131_v44_))]) | (d_2130_v43_)) in (d_2127_v40_) else (len(d_2132_v45_)) != (d_2082_v1_))
            if d_2081_v0_:
                d_2133_v46_: _dafny.Map
                d_2133_v46_ = _dafny.Map({d_2128_v41_: False})
                d_2133_v46_ = (d_2133_v46_).set(d_2128_v41_, (self).fm3(885, d_2126_i4_, globalState))
                d_2134_v47_: _dafny.Array
                nw422_ = _dafny.Array(None, 4)
                nw422_[int(0)] = d_2081_v0_
                nw422_[int(1)] = d_2081_v0_
                nw422_[int(2)] = d_2081_v0_
                nw422_[int(3)] = d_2081_v0_
                d_2134_v47_ = nw422_
                d_2135_v48_: _dafny.Map
                d_2135_v48_ = _dafny.Map({d_2134_v47_: True})
                d_2135_v48_ = (d_2135_v48_).set(d_2134_v47_, d_2081_v0_)
                d_2136_v49_: _dafny.Array
                def lambda138_(d_2137_i4_, d_2138_v0_):
                    def lambda139_(d_2139_i5_):
                        return D21_DC50(_dafny.Map({D14_DC32((_dafny.MultiSet([d_2137_i4_])).cardinality, d_2138_v0_, d_2138_v0_): _dafny.Map({d_2138_v0_: d_2138_v0_})}))

                    return lambda139_

                init78_ = lambda138_(d_2126_i4_, d_2081_v0_)
                nw423_ = _dafny.Array(None, 25)
                for i0_78_ in range(nw423_.length(0)):
                    nw423_[i0_78_] = init78_(i0_78_)
                d_2136_v49_ = nw423_
                d_2136_v49_ = d_2136_v49_
                index410_ = default__.safeIndex(93, (d_2134_v47_).length(0))
                (d_2134_v47_)[index410_] = d_2081_v0_
                index411_ = default__.safeIndex(604, (d_2134_v47_).length(0))
                (d_2134_v47_)[index411_] = d_2081_v0_
                index412_ = default__.safeIndex(335, (d_2134_v47_).length(0))
                (d_2134_v47_)[index412_] = (d_2085_v4_) != (d_2085_v4_)
                index413_ = default__.safeIndex(93, (d_2134_v47_).length(0))
                index414_ = default__.safeIndex(604, (d_2134_v47_).length(0))
                index415_ = default__.safeIndex(335, (d_2134_v47_).length(0))
                rhs407_ = d_2081_v0_
                rhs408_ = d_2125_v39_
                rhs409_ = (d_2086_v5_)[default__.safeIndex(d_2126_i4_, len(d_2086_v5_))]
                rhs410_ = not ((self).fm3(default__.fm13(d_2081_v0_, d_2081_v0_, d_2081_v0_, globalState), default__.fm14(globalState), globalState)) or (d_2081_v0_)
                rhs411_ = d_2081_v0_
                lhs202_ = d_2134_v47_
                lhs203_ = default__.safeIndex(93, (d_2134_v47_).length(0))
                lhs204_ = d_2134_v47_
                lhs205_ = default__.safeIndex(604, (d_2134_v47_).length(0))
                lhs206_ = d_2134_v47_
                lhs207_ = default__.safeIndex(335, (d_2134_v47_).length(0))
                d_2081_v0_ = rhs407_
                d_2125_v39_ = rhs408_
                lhs202_[lhs203_] = rhs409_
                lhs204_[lhs205_] = rhs410_
                lhs206_[lhs207_] = rhs411_
                d_2140_v50_: D18
                d_2140_v50_ = D18_DC45(d_2134_v47_, d_2081_v0_, d_2082_v1_, 128)
                d_2141_v51_: _dafny.Map
                d_2141_v51_ = _dafny.Map({D5_DC11(d_2083_v2_): (d_2140_v50_).cf69})
                d_2082_v1_ = (0) - (len((d_2132_v45_).set(default__.safeIndex(len(default__.fm31(180, len(d_2141_v51_), globalState)), len(d_2132_v45_)), (self).f5)))
            elif True:
                d_2081_v0_ = d_2081_v0_
                d_2132_v45_ = (default__.fm12(globalState) if d_2081_v0_ else (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2142_i6_ in range(default__.abs(-992))]) if d_2081_v0_ else d_2132_v45_))
                d_2143_v52_: _dafny.Map
                d_2143_v52_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx"))) + (d_2132_v45_): d_2132_v45_})
                d_2143_v52_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f5 for d_2144_i7_ in range(default__.abs(311))])) + (d_2132_v45_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmmasuhrp"))})
                d_2145_v53_: _dafny.Array
                def lambda140_(d_2146_i8_):
                    return False

                init79_ = lambda140_
                nw424_ = _dafny.Array(None, 22)
                for i0_79_ in range(nw424_.length(0)):
                    nw424_[i0_79_] = init79_(i0_79_)
                d_2145_v53_ = nw424_
                d_2147_v54_: _dafny.Set
                d_2147_v54_ = _dafny.Set({d_2145_v53_})
                d_2148_v55_: _dafny.Set
                d_2148_v55_ = d_2147_v54_
                d_2149_v56_: _dafny.Map
                d_2149_v56_ = _dafny.Map({d_2081_v0_: d_2148_v55_})
                d_2150_v57_: _dafny.Map
                d_2151_v58_: int
                d_2152_v59_: bool
                d_2153_v60_: _dafny.Map
                out21_: _dafny.Map
                out22_: int
                out23_: bool
                out24_: _dafny.Map
                out21_, out22_, out23_, out24_ = (self).m5(d_2132_v45_, d_2149_v56_, d_2126_i4_, globalState)
                d_2150_v57_ = out21_
                d_2151_v58_ = out22_
                d_2152_v59_ = out23_
                d_2153_v60_ = out24_
                d_2154_v61_: _dafny.Map
                d_2154_v61_ = _dafny.Map({not((d_2082_v1_) > (d_2151_v58_)): (not(d_2152_v59_) if d_2081_v0_ else d_2152_v59_)})
                d_2154_v61_ = (d_2154_v61_).set(d_2152_v59_, d_2152_v59_)
            d_2082_v1_ = (d_2082_v1_) - (d_2126_i4_)
        d_2155_v62_: _dafny.Array
        def lambda141_(d_2156_v1_):
            def lambda142_(d_2157_i9_):
                return (d_2157_i9_) * (d_2156_v1_)

            return lambda142_

        init80_ = lambda141_(d_2082_v1_)
        nw425_ = _dafny.Array(None, 3)
        for i0_80_ in range(nw425_.length(0)):
            nw425_[i0_80_] = init80_(i0_80_)
        d_2155_v62_ = nw425_
        d_2158_v63_: _dafny.Map
        d_2158_v63_ = _dafny.Map({d_2155_v62_: d_2082_v1_})
        d_2158_v63_ = d_2158_v63_
        r0 = d_2155_v62_
        return r0

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: str = _dafny.CodePoint('D')
        d_2159_v0_: bool
        d_2159_v0_ = True
        d_2160_i0_: int
        d_2160_i0_ = 0
        with _dafny.label("8"):
            while d_2159_v0_:
                with _dafny.c_label("8"):
                    if (d_2160_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_2160_i0_ = (d_2160_i0_) + (1)
                    d_2159_v0_ = not(d_2159_v0_)
                    r3 = (self).f5
                    d_2161_v1_: int
                    d_2161_v1_ = 444
                    d_2162_v2_: _dafny.Map
                    d_2162_v2_ = _dafny.Map({(d_2161_v1_) * (p1): len(_dafny.SeqWithoutIsStrInference([d_2159_v0_, d_2159_v0_]))})
                    d_2163_v3_: _dafny.Array
                    nw426_ = _dafny.Array(int(0), 15)
                    d_2163_v3_ = nw426_
                    d_2164_v4_: _dafny.MultiSet
                    d_2164_v4_ = _dafny.MultiSet([d_2163_v3_, d_2163_v3_])
                    d_2165_v5_: _dafny.Seq
                    d_2165_v5_ = _dafny.SeqWithoutIsStrInference([((d_2164_v4_)[d_2163_v3_] if (d_2163_v3_) in (d_2164_v4_) else p0), d_2161_v1_, p1])
                    rhs412_ = ((self).fm5(d_2161_v1_, d_2159_v0_, p0, p1, globalState) if (d_2159_v0_) == (d_2159_v0_) else d_2159_v0_)
                    rhs413_ = ((d_2162_v2_)[(d_2165_v5_)[default__.safeIndex(d_2161_v1_, len(d_2165_v5_))]] if ((d_2165_v5_)[default__.safeIndex(d_2161_v1_, len(d_2165_v5_))]) in (d_2162_v2_) else d_2161_v1_)
                    r0 = rhs412_
                    d_2161_v1_ = rhs413_
                    d_2166_v6_: _dafny.Array
                    def lambda143_(d_2167_p1_):
                        def lambda144_(d_2168_i1_):
                            return (_dafny.Set({d_2167_p1_})).ispropersubset(_dafny.Set({d_2167_p1_}))

                        return lambda144_

                    init81_ = lambda143_(p1)
                    nw427_ = _dafny.Array(None, 5)
                    for i0_81_ in range(nw427_.length(0)):
                        nw427_[i0_81_] = init81_(i0_81_)
                    d_2166_v6_ = nw427_
                    index416_ = default__.safeIndex(398, (d_2166_v6_).length(0))
                    (d_2166_v6_)[index416_] = (p0) > (d_2161_v1_)
                    index417_ = default__.safeIndex(398, (d_2166_v6_).length(0))
                    (d_2166_v6_)[index417_] = not (default__.fm1(d_2161_v1_, p1, True, d_2159_v0_, globalState)) or (d_2159_v0_)
                    pass
            pass
        d_2169_v7_: _dafny.Seq
        d_2169_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bveyecq"))
        d_2170_v8_: _dafny.Array
        nw428_ = _dafny.Array(None, 19)
        nw428_[int(0)] = default__.fm16(d_2159_v0_, p0, globalState)
        nw428_[int(1)] = d_2169_v7_
        nw428_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f5 for d_2171_i2_ in range(default__.abs(284))])
        nw428_[int(3)] = default__.fm12(globalState)
        nw428_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhgvbpyh"))) + (d_2169_v7_)
        nw428_[int(5)] = d_2169_v7_
        nw428_[int(6)] = d_2169_v7_
        nw428_[int(7)] = d_2169_v7_
        nw428_[int(8)] = d_2169_v7_
        nw428_[int(9)] = (d_2169_v7_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvhh")))
        nw428_[int(10)] = default__.fm24(-430, globalState)
        nw428_[int(11)] = d_2169_v7_
        nw428_[int(12)] = d_2169_v7_
        nw428_[int(13)] = d_2169_v7_
        nw428_[int(14)] = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2172_i3_ in range(default__.abs(95))])) + (d_2169_v7_)
        nw428_[int(15)] = d_2169_v7_
        nw428_[int(16)] = d_2169_v7_
        nw428_[int(17)] = (d_2169_v7_) + (d_2169_v7_)
        nw428_[int(18)] = (d_2169_v7_).set(default__.safeIndex(p0, len(d_2169_v7_)), _dafny.CodePoint('l'))
        d_2170_v8_ = nw428_
        index418_ = default__.safeIndex(27, (d_2170_v8_).length(0))
        (d_2170_v8_)[index418_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2173_i4_ in range(default__.abs(409))])
        d_2174_v9_: int
        d_2174_v9_ = 334
        d_2175_v10_: _dafny.Map
        d_2175_v10_ = _dafny.Map({p1: d_2169_v7_})
        d_2176_v11_: _dafny.Seq
        d_2176_v11_ = _dafny.SeqWithoutIsStrInference([d_2175_v10_])
        d_2177_v12_: _dafny.Seq
        d_2177_v12_ = _dafny.SeqWithoutIsStrInference([d_2174_v9_])
        d_2178_v13_: D2
        d_2178_v13_ = D2_DC6(p0)
        index419_ = default__.safeIndex(27, (d_2170_v8_).length(0))
        rhs414_ = d_2169_v7_
        rhs415_ = (((0) - (d_2174_v9_)) - (d_2174_v9_)) - ((0) - (d_2174_v9_))
        rhs416_ = (d_2176_v11_)[default__.safeIndex(d_2174_v9_, len(d_2176_v11_))]
        rhs417_ = ((self).fm4(d_2177_v12_, globalState)) + (d_2174_v9_)
        rhs418_ = ((d_2178_v13_).cf7) <= (default__.fm14(globalState))
        lhs208_ = d_2170_v8_
        lhs209_ = default__.safeIndex(27, (d_2170_v8_).length(0))
        lhs208_[lhs209_] = rhs414_
        d_2174_v9_ = rhs415_
        d_2175_v10_ = rhs416_
        d_2174_v9_ = rhs417_
        d_2159_v0_ = rhs418_
        d_2179_v14_: D5
        d_2179_v14_ = D5_DC13(d_2169_v7_, len(d_2169_v7_), d_2159_v0_)
        if (d_2179_v14_).cf18:
            d_2180_v15_: _dafny.Array
            nw429_ = _dafny.Array(False, 5)
            d_2180_v15_ = nw429_
            d_2181_v16_: _dafny.Map
            d_2181_v16_ = _dafny.Map({d_2180_v15_: d_2159_v0_})
            d_2182_v17_: _dafny.Seq
            d_2182_v17_ = _dafny.SeqWithoutIsStrInference([d_2159_v0_])
            d_2183_v18_: _dafny.MultiSet
            d_2183_v18_ = _dafny.MultiSet([d_2159_v0_])
            d_2184_v19_: _dafny.Array
            nw430_ = _dafny.Array(None, 25)
            nw430_[int(0)] = (_dafny.MultiSet(d_2177_v12_)).cardinality
            nw430_[int(1)] = len(((d_2181_v16_).set(d_2180_v15_, False)) | (d_2181_v16_))
            nw430_[int(2)] = (d_2174_v9_) * ((0) - ((self).fm4(d_2177_v12_, globalState)))
            nw430_[int(3)] = d_2174_v9_
            nw430_[int(4)] = p1
            nw430_[int(5)] = (p0) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))))
            nw430_[int(6)] = (p0) * (p1)
            nw430_[int(7)] = len((d_2177_v12_) + (d_2177_v12_))
            nw430_[int(8)] = 917
            nw430_[int(9)] = len(d_2182_v17_)
            nw430_[int(10)] = d_2174_v9_
            nw430_[int(11)] = (((d_2183_v18_)[default__.fm1(-304, p1, d_2159_v0_, d_2159_v0_, globalState)] if (default__.fm1(-304, p1, d_2159_v0_, d_2159_v0_, globalState)) in (d_2183_v18_) else p1)) + (len(_dafny.SeqWithoutIsStrInference([d_2159_v0_, d_2159_v0_, d_2159_v0_])))
            nw430_[int(12)] = 58
            nw430_[int(13)] = p1
            nw430_[int(14)] = (p0) + ((0) - (p1))
            nw430_[int(15)] = d_2174_v9_
            nw430_[int(16)] = p1
            nw430_[int(17)] = (p1) - (d_2174_v9_)
            nw430_[int(18)] = (d_2174_v9_) - (p1)
            nw430_[int(19)] = d_2174_v9_
            nw430_[int(20)] = (self).fm4(d_2177_v12_, globalState)
            nw430_[int(21)] = p0
            nw430_[int(22)] = d_2174_v9_
            nw430_[int(23)] = default__.safeDivisionInt(p0, 970)
            nw430_[int(24)] = (p0) + (p1)
            d_2184_v19_ = nw430_
            index420_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            (d_2184_v19_)[index420_] = p1
            d_2185_v20_: _dafny.Set
            d_2185_v20_ = _dafny.Set({d_2159_v0_})
            d_2186_v21_: D16
            d_2186_v21_ = D16_DC38((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))], not(d_2159_v0_), len((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]), default__.fm1(p1, p0, d_2159_v0_, d_2159_v0_, globalState))
            d_2187_v22_: _dafny.Set
            d_2187_v22_ = _dafny.Set({d_2185_v20_, _dafny.Set({(d_2186_v21_).cf58})})
            d_2188_v23_: D18
            d_2188_v23_ = D18_DC44(d_2187_v22_)
            d_2189_v24_: _dafny.Map
            d_2189_v24_ = _dafny.Map({d_2159_v0_: d_2159_v0_})
            d_2190_v25_: _dafny.Array
            nw431_ = _dafny.Array(None, 20)
            nw431_[int(0)] = (self).f5
            nw431_[int(1)] = default__.fm0(966, (d_2177_v12_)[default__.safeIndex(p1, len(d_2177_v12_))], globalState)
            nw431_[int(2)] = (self).f5
            nw431_[int(3)] = (self).f5
            nw431_[int(4)] = (self).f5
            nw431_[int(5)] = default__.fm0(d_2174_v9_, d_2174_v9_, globalState)
            nw431_[int(6)] = (self).f5
            nw431_[int(7)] = (self).f5
            nw431_[int(8)] = default__.fm0(len(_dafny.Map({(self).f5: d_2189_v24_})), d_2174_v9_, globalState)
            nw431_[int(9)] = (self).f5
            nw431_[int(10)] = (self).f5
            nw431_[int(11)] = (self).f5
            nw431_[int(12)] = (self).f5
            nw431_[int(13)] = (self).f5
            nw431_[int(14)] = (self).f5
            nw431_[int(15)] = (self).f5
            nw431_[int(16)] = (self).f5
            nw431_[int(17)] = (self).f5
            nw431_[int(18)] = ((self).f5 if not(d_2159_v0_) else (self).f5)
            nw431_[int(19)] = (self).f5
            d_2190_v25_ = nw431_
            index421_ = default__.safeIndex(747, (d_2190_v25_).length(0))
            (d_2190_v25_)[index421_] = (self).f5
            d_2191_v26_: _dafny.Map
            d_2191_v26_ = _dafny.Map({_dafny.CodePoint('g'): d_2159_v0_})
            d_2192_v27_: _dafny.Map
            d_2192_v27_ = _dafny.Map({(d_2191_v26_) | (_dafny.Map({(self).f5: d_2159_v0_})): p1})
            index422_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            index423_ = default__.safeIndex(747, (d_2190_v25_).length(0))
            rhs419_ = len((d_2177_v12_) + (_dafny.SeqWithoutIsStrInference([d_2174_v9_])))
            rhs420_ = d_2188_v23_
            rhs421_ = (self).f5
            rhs422_ = d_2192_v27_
            lhs210_ = d_2184_v19_
            lhs211_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            lhs212_ = d_2190_v25_
            lhs213_ = default__.safeIndex(747, (d_2190_v25_).length(0))
            lhs210_[lhs211_] = rhs419_
            d_2188_v23_ = rhs420_
            lhs212_[lhs213_] = rhs421_
            d_2192_v27_ = rhs422_
            index424_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            (d_2184_v19_)[index424_] = p1
            index425_ = default__.safeIndex(747, (d_2190_v25_).length(0))
            (d_2190_v25_)[index425_] = _dafny.CodePoint('g')
            d_2193_v28_: _dafny.Map
            d_2193_v28_ = _dafny.Map({d_2159_v0_: d_2180_v15_})
            d_2194_v29_: D17
            d_2194_v29_ = D17_DC42(d_2180_v15_, d_2174_v9_)
            d_2195_v30_: _dafny.Map
            d_2195_v30_ = _dafny.Map({921: d_2180_v15_})
            d_2196_v31_: _dafny.Array
            nw432_ = _dafny.Array(None, 26)
            nw432_[int(0)] = d_2180_v15_
            nw432_[int(1)] = d_2180_v15_
            nw432_[int(2)] = ((d_2193_v28_)[d_2159_v0_] if (d_2159_v0_) in (d_2193_v28_) else d_2180_v15_)
            nw432_[int(3)] = (d_2194_v29_).cf64
            nw432_[int(4)] = d_2180_v15_
            nw432_[int(5)] = (d_2180_v15_ if not(not(True)) else d_2180_v15_)
            nw432_[int(6)] = d_2180_v15_
            nw432_[int(7)] = d_2180_v15_
            nw432_[int(8)] = d_2180_v15_
            nw432_[int(9)] = d_2180_v15_
            nw432_[int(10)] = ((d_2195_v30_)[p1] if (p1) in (d_2195_v30_) else d_2180_v15_)
            nw432_[int(11)] = d_2180_v15_
            nw432_[int(12)] = (d_2180_v15_ if d_2159_v0_ else d_2180_v15_)
            nw432_[int(13)] = d_2180_v15_
            nw432_[int(14)] = d_2180_v15_
            nw432_[int(15)] = d_2180_v15_
            nw432_[int(16)] = d_2180_v15_
            nw432_[int(17)] = d_2180_v15_
            nw432_[int(18)] = d_2180_v15_
            nw432_[int(19)] = d_2180_v15_
            nw432_[int(20)] = d_2180_v15_
            nw432_[int(21)] = d_2180_v15_
            nw432_[int(22)] = d_2180_v15_
            nw432_[int(23)] = d_2180_v15_
            nw432_[int(24)] = d_2180_v15_
            nw432_[int(25)] = d_2180_v15_
            d_2196_v31_ = nw432_
            index426_ = default__.safeIndex(36, (d_2196_v31_).length(0))
            (d_2196_v31_)[index426_] = d_2180_v15_
            d_2197_v32_: D18
            d_2197_v32_ = D18_DC45(d_2180_v15_, d_2159_v0_, d_2174_v9_, (d_2184_v19_)[default__.safeIndex(217, (d_2184_v19_).length(0))])
            index427_ = default__.safeIndex(36, (d_2196_v31_).length(0))
            index428_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            index429_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            rhs423_ = default__.safeModuloInt(p1, (0) - ((d_2184_v19_)[default__.safeIndex(217, (d_2184_v19_).length(0))]))
            rhs424_ = (d_2197_v32_).cf68
            rhs425_ = default__.fm14(globalState)
            rhs426_ = (d_2184_v19_)[default__.safeIndex(217, (d_2184_v19_).length(0))]
            rhs427_ = (d_2184_v19_)[default__.safeIndex(217, (d_2184_v19_).length(0))]
            lhs214_ = d_2196_v31_
            lhs215_ = default__.safeIndex(36, (d_2196_v31_).length(0))
            lhs216_ = d_2184_v19_
            lhs217_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            lhs218_ = d_2184_v19_
            lhs219_ = default__.safeIndex(217, (d_2184_v19_).length(0))
            d_2174_v9_ = rhs423_
            lhs214_[lhs215_] = rhs424_
            lhs216_[lhs217_] = rhs425_
            d_2174_v9_ = rhs426_
            lhs218_[lhs219_] = rhs427_
            d_2198_v33_: _dafny.Array
            nw433_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_2198_v33_ = nw433_
            index430_ = default__.safeIndex(617, (d_2198_v33_).length(0))
            (d_2198_v33_)[index430_] = d_2184_v19_
            index431_ = default__.safeIndex(617, (d_2198_v33_).length(0))
            def lambda145_(d_2199_v0_):
                def lambda146_(d_2200_i5_):
                    return (d_2200_i5_) - (len(_dafny.Map({not(d_2199_v0_): _dafny.Map({d_2199_v0_: 730})})))

                return lambda146_

            init82_ = lambda145_(d_2159_v0_)
            nw434_ = _dafny.Array(None, 10)
            for i0_82_ in range(nw434_.length(0)):
                nw434_[i0_82_] = init82_(i0_82_)
            (d_2198_v33_)[index431_] = nw434_
        elif True:
            d_2201_v34_: D16
            d_2201_v34_ = D16_DC38((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))], d_2159_v0_, p1, d_2159_v0_)
            d_2202_v35_: D16
            d_2202_v35_ = D16_DC39(d_2201_v34_)
            d_2203_v36_: _dafny.MultiSet
            d_2203_v36_ = _dafny.MultiSet([d_2202_v35_, d_2202_v35_, d_2202_v35_])
            d_2204_v37_: _dafny.Seq
            d_2204_v37_ = _dafny.SeqWithoutIsStrInference([d_2203_v36_, d_2203_v36_, d_2203_v36_, _dafny.MultiSet([d_2202_v35_, d_2202_v35_, d_2202_v35_, d_2202_v35_])])
            d_2205_v38_: _dafny.Map
            d_2205_v38_ = _dafny.Map({(self).f5: d_2159_v0_})
            d_2206_v39_: _dafny.Seq
            d_2206_v39_ = _dafny.SeqWithoutIsStrInference([d_2159_v0_, d_2159_v0_])
            d_2207_v40_: _dafny.Seq
            d_2207_v40_ = _dafny.SeqWithoutIsStrInference([d_2203_v36_, _dafny.MultiSet([d_2202_v35_, d_2202_v35_, d_2202_v35_]), (d_2203_v36_) | ((d_2204_v37_)[default__.safeIndex(len(d_2205_v38_), len(d_2204_v37_))]), (default__.fm44(globalState)), ((d_2203_v36_).set(d_2202_v35_, default__.abs(len(d_2206_v39_))) if d_2159_v0_ else d_2203_v36_)])
            d_2207_v40_ = (d_2207_v40_) + (d_2204_v37_)
            d_2208_v41_: C7
            nw435_ = C7()
            nw435_.ctor__()
            d_2208_v41_ = nw435_
            d_2209_v42_: _dafny.Map
            d_2209_v42_ = _dafny.Map({(self).f5: d_2208_v41_})
            d_2210_v43_: _dafny.Map
            d_2210_v43_ = _dafny.Map({(d_2209_v42_) | (d_2209_v42_): d_2159_v0_})
            d_2210_v43_ = d_2210_v43_
            d_2211_v44_: _dafny.Array
            def lambda147_(d_2212_i6_):
                return True

            init83_ = lambda147_
            nw436_ = _dafny.Array(None, 15)
            for i0_83_ in range(nw436_.length(0)):
                nw436_[i0_83_] = init83_(i0_83_)
            d_2211_v44_ = nw436_
            rhs428_ = d_2211_v44_
            rhs429_ = d_2159_v0_
            lhs220_ = globalState
            lhs220_.f3 = rhs428_
            r0 = rhs429_
            d_2174_v9_ = len(d_2206_v39_)
            d_2213_v45_: _dafny.MultiSet
            d_2213_v45_ = _dafny.MultiSet([(d_2211_v44_ if d_2159_v0_ else d_2211_v44_), d_2211_v44_, d_2211_v44_, d_2211_v44_, d_2211_v44_])
            d_2213_v45_ = d_2213_v45_
        if d_2159_v0_:
            d_2174_v9_ = (d_2174_v9_ if d_2159_v0_ else default__.fm14(globalState))
            d_2174_v9_ = p1
            d_2214_v46_: _dafny.Seq
            d_2214_v46_ = _dafny.SeqWithoutIsStrInference([True])
            d_2215_v47_: _dafny.Seq
            d_2215_v47_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_2174_v9_, d_2174_v9_, d_2159_v0_, (d_2214_v46_)[default__.safeIndex(d_2174_v9_, len(d_2214_v46_))], globalState)])
            r3 = ((self).f5 if (d_2159_v0_) == ((d_2215_v47_)[default__.safeIndex(d_2174_v9_, len(d_2215_v47_))]) else (self).f5)
            r0 = d_2159_v0_
            d_2174_v9_ = p1
        elif True:
            d_2216_v48_: _dafny.Array
            def lambda148_(d_2217_p0_):
                def lambda149_(d_2218_i7_):
                    return (d_2218_i7_) * (d_2217_p0_)

                return lambda149_

            init84_ = lambda148_(p0)
            nw437_ = _dafny.Array(None, 10)
            for i0_84_ in range(nw437_.length(0)):
                nw437_[i0_84_] = init84_(i0_84_)
            d_2216_v48_ = nw437_
            index432_ = default__.safeIndex(542, (d_2216_v48_).length(0))
            (d_2216_v48_)[index432_] = p0
            index433_ = default__.safeIndex(542, (d_2216_v48_).length(0))
            (d_2216_v48_)[index433_] = 878
            d_2159_v0_ = (self).fm5(p0, d_2159_v0_, p0, (d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))], globalState)
            if False:
                d_2219_v49_: _dafny.MultiSet
                d_2219_v49_ = _dafny.MultiSet([d_2159_v0_, d_2159_v0_, d_2159_v0_, d_2159_v0_])
                d_2219_v49_ = d_2219_v49_
                d_2220_v50_: D7
                d_2220_v50_ = D7_DC16((self).f5)
                d_2221_v51_: _dafny.MultiSet
                d_2221_v51_ = _dafny.MultiSet([_dafny.CodePoint('b'), (d_2220_v50_).cf23])
                d_2222_v52_: _dafny.Map
                d_2222_v52_ = _dafny.Map({490: (d_2221_v51_).cardinality})
                d_2223_v53_: _dafny.Seq
                d_2223_v53_ = _dafny.SeqWithoutIsStrInference([d_2219_v49_, _dafny.MultiSet([d_2159_v0_, d_2159_v0_])])
                d_2224_v54_: _dafny.Seq
                d_2224_v54_ = _dafny.SeqWithoutIsStrInference([d_2159_v0_, d_2159_v0_, d_2159_v0_])
                d_2225_v55_: D11
                d_2225_v55_ = D11_DC25(d_2159_v0_, -927, d_2216_v48_, p1)
                d_2226_v56_: _dafny.Array
                nw438_ = _dafny.Array(None, 28)
                nw438_[int(0)] = d_2159_v0_
                nw438_[int(1)] = (d_2174_v9_) > ((0) - (p0))
                nw438_[int(2)] = (d_2159_v0_ if not(d_2159_v0_) else (self).fm5((0) - (d_2174_v9_), d_2159_v0_, p0, p1, globalState))
                nw438_[int(3)] = d_2159_v0_
                nw438_[int(4)] = (self).fm3((d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))], ((d_2222_v52_)[len(d_2169_v7_)] if (len(d_2169_v7_)) in (d_2222_v52_) else len(_dafny.SeqWithoutIsStrInference([p1]))), globalState)
                nw438_[int(5)] = d_2159_v0_
                nw438_[int(6)] = d_2159_v0_
                nw438_[int(7)] = d_2159_v0_
                nw438_[int(8)] = False
                nw438_[int(9)] = (d_2174_v9_) == (p0)
                nw438_[int(10)] = (d_2221_v51_).issubset(_dafny.MultiSet([(self).f5]))
                nw438_[int(11)] = False
                nw438_[int(12)] = d_2159_v0_
                nw438_[int(13)] = d_2159_v0_
                nw438_[int(14)] = d_2159_v0_
                nw438_[int(15)] = d_2159_v0_
                nw438_[int(16)] = d_2159_v0_
                nw438_[int(17)] = (False if (self).fm5(d_2174_v9_, d_2159_v0_, (d_2219_v49_).cardinality, len(d_2169_v7_), globalState) else d_2159_v0_)
                nw438_[int(18)] = d_2159_v0_
                nw438_[int(19)] = (d_2159_v0_ if not(d_2159_v0_) else d_2159_v0_)
                nw438_[int(20)] = d_2159_v0_
                nw438_[int(21)] = not(d_2159_v0_)
                nw438_[int(22)] = (_dafny.MultiSet(d_2224_v54_)).ispropersubset((d_2223_v53_)[default__.safeIndex(p0, len(d_2223_v53_))])
                nw438_[int(23)] = ((d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))]) < (p1)
                nw438_[int(24)] = d_2159_v0_
                nw438_[int(25)] = d_2159_v0_
                nw438_[int(26)] = (d_2225_v55_).cf38
                nw438_[int(27)] = (d_2224_v54_) <= (d_2224_v54_)
                d_2226_v56_ = nw438_
                index434_ = default__.safeIndex(710, (d_2226_v56_).length(0))
                (d_2226_v56_)[index434_] = (d_2219_v49_).isdisjoint(d_2219_v49_)
                d_2227_v57_: _dafny.Set
                d_2227_v57_ = _dafny.Set({d_2224_v54_})
                d_2228_v58_: _dafny.Map
                d_2228_v58_ = _dafny.Map({(d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))]: d_2227_v57_})
                d_2229_v59_: D23
                d_2229_v59_ = D23_DC54(((d_2228_v58_)[(self).fm4(_dafny.SeqWithoutIsStrInference([(d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))] for d_2230_i8_ in range(default__.abs(224))]), globalState)] if ((self).fm4(_dafny.SeqWithoutIsStrInference([(d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))] for d_2231_i8_ in range(default__.abs(224))]), globalState)) in (d_2228_v58_) else d_2227_v57_))
                index435_ = default__.safeIndex(710, (d_2226_v56_).length(0))
                rhs430_ = d_2216_v48_
                rhs431_ = (d_2227_v57_).ispropersubset((d_2229_v59_).cf84)
                lhs221_ = d_2226_v56_
                lhs222_ = default__.safeIndex(710, (d_2226_v56_).length(0))
                d_2216_v48_ = rhs430_
                lhs221_[lhs222_] = rhs431_
                index436_ = default__.safeIndex(542, (d_2216_v48_).length(0))
                (d_2216_v48_)[index436_] = ((d_2222_v52_)[(d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))]] if ((d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))]) in (d_2222_v52_) else default__.safeModuloInt(len(d_2169_v7_), p0))
                d_2232_v60_: C5
                nw439_ = C5()
                nw439_.ctor__()
                d_2232_v60_ = nw439_
                d_2233_v61_: _dafny.Map
                d_2233_v61_ = _dafny.Map({(0) - ((0) - ((d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))])): (d_2226_v56_)[default__.safeIndex(710, (d_2226_v56_).length(0))]})
                d_2234_v62_: _dafny.Map
                d_2234_v62_ = _dafny.Map({d_2233_v61_: (self).f4})
                d_2235_v63_: C3
                nw440_ = C3()
                nw440_.ctor__(((d_2234_v62_)[d_2233_v61_] if (d_2233_v61_) in (d_2234_v62_) else (self).f4), default__.fm0(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pufjmt"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pufjmt")))), (self).f5)), (d_2216_v48_)[default__.safeIndex(542, (d_2216_v48_).length(0))], globalState))
                d_2235_v63_ = nw440_
            elif True:
                d_2236_v64_: _dafny.Seq
                d_2236_v64_ = _dafny.SeqWithoutIsStrInference([d_2216_v48_, d_2216_v48_, d_2216_v48_])
                d_2237_v65_: _dafny.Seq
                d_2237_v65_ = _dafny.SeqWithoutIsStrInference([False, d_2159_v0_])
                d_2236_v64_ = (((d_2236_v64_) + (d_2236_v64_)).set(default__.safeIndex(len(d_2237_v65_), len((d_2236_v64_) + (d_2236_v64_))), d_2216_v48_)) + (d_2236_v64_)
                d_2238_v66_: _dafny.Map
                d_2238_v66_ = _dafny.Map({_dafny.MultiSet(d_2177_v12_): (d_2237_v65_)[default__.safeIndex(len(d_2169_v7_), len(d_2237_v65_))]})
                d_2239_v67_: _dafny.Map
                d_2239_v67_ = d_2238_v66_
                index437_ = default__.safeIndex(542, (d_2216_v48_).length(0))
                rhs432_ = len(((d_2238_v66_) | ((d_2239_v67_))) | (d_2238_v66_))
                rhs433_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2240_i9_ in range(default__.abs(670))])
                lhs223_ = d_2216_v48_
                lhs224_ = default__.safeIndex(542, (d_2216_v48_).length(0))
                lhs223_[lhs224_] = rhs432_
                d_2169_v7_ = rhs433_
                index438_ = default__.safeIndex(542, (d_2216_v48_).length(0))
                (d_2216_v48_)[index438_] = p0
                d_2241_v68_: _dafny.Array
                nw441_ = _dafny.Array(_dafny.Map({}), 17)
                d_2241_v68_ = nw441_
                index439_ = default__.safeIndex(695, (d_2241_v68_).length(0))
                (d_2241_v68_)[index439_] = _dafny.Map({d_2159_v0_: False})
                d_2242_v69_: _dafny.Map
                d_2242_v69_ = _dafny.Map({d_2159_v0_: d_2159_v0_})
                index440_ = default__.safeIndex(695, (d_2241_v68_).length(0))
                (d_2241_v68_)[index440_] = (d_2242_v69_ if d_2159_v0_ else (_dafny.Map({d_2159_v0_: False})) | (d_2242_v69_))
                d_2243_v70_: _dafny.Array
                nw442_ = _dafny.Array(False, 18)
                d_2243_v70_ = nw442_
                d_2244_v71_: _dafny.MultiSet
                d_2244_v71_ = _dafny.MultiSet([d_2243_v70_, d_2243_v70_])
                d_2245_v72_: D25
                d_2245_v72_ = D25_DC59(d_2244_v71_)
                d_2244_v71_ = ((d_2245_v72_).cf95) - (d_2244_v71_)
            d_2246_v73_: _dafny.Array
            def lambda150_(d_2247_v0_):
                def lambda151_(d_2248_i10_):
                    return d_2247_v0_

                return lambda151_

            init85_ = lambda150_(d_2159_v0_)
            nw443_ = _dafny.Array(None, 27)
            for i0_85_ in range(nw443_.length(0)):
                nw443_[i0_85_] = init85_(i0_85_)
            d_2246_v73_ = nw443_
            index441_ = default__.safeIndex(685, (d_2246_v73_).length(0))
            (d_2246_v73_)[index441_] = d_2159_v0_
            d_2249_v74_: _dafny.MultiSet
            d_2249_v74_ = _dafny.MultiSet([(self).f5, (self).f5])
            index442_ = default__.safeIndex(685, (d_2246_v73_).length(0))
            (d_2246_v73_)[index442_] = (d_2249_v74_) == ((d_2249_v74_ if d_2159_v0_ else d_2249_v74_))
            r0 = (d_2246_v73_)[default__.safeIndex(685, (d_2246_v73_).length(0))]
        if (default__.fm45(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtaqiowis"))), globalState)).cf97:
            r3 = (self).f5
            d_2250_v75_: _dafny.Seq
            d_2250_v75_ = _dafny.SeqWithoutIsStrInference([d_2159_v0_, d_2159_v0_, d_2159_v0_, False, d_2159_v0_])
            d_2251_v76_: _dafny.Seq
            d_2251_v76_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_2174_v9_, p1, (d_2250_v75_)[default__.safeIndex(p0, len(d_2250_v75_))], d_2159_v0_, globalState), d_2159_v0_, True])
            d_2252_v77_: D23
            d_2252_v77_ = D23_DC55((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))], d_2159_v0_, d_2159_v0_, (d_2159_v0_) or (d_2159_v0_), d_2250_v75_)
            source22_ = d_2252_v77_
            if source22_.is_DC55:
                d_2253___mcc_h0_ = source22_.cf85
                d_2254___mcc_h1_ = source22_.cf86
                d_2255___mcc_h2_ = source22_.cf87
                d_2256___mcc_h3_ = source22_.cf88
                d_2257___mcc_h4_ = source22_.cf89
                d_2258_cf89_ = d_2257___mcc_h4_
                d_2259_cf88_ = d_2256___mcc_h3_
                d_2260_cf87_ = d_2255___mcc_h2_
                d_2261_cf86_ = d_2254___mcc_h1_
                d_2262_cf85_ = d_2253___mcc_h0_
                d_2259_cf88_ = (d_2260_cf87_) == ((d_2159_v0_) and (d_2159_v0_))
                r2 = default__.fm11(globalState)
                d_2263_v78_: D16
                d_2263_v78_ = D16_DC38(d_2262_cf85_, d_2259_cf88_, p0, d_2159_v0_)
                d_2264_v79_: _dafny.Array
                nw444_ = _dafny.Array(None, 20)
                nw444_[int(0)] = d_2260_cf87_
                nw444_[int(1)] = d_2159_v0_
                nw444_[int(2)] = True
                nw444_[int(3)] = d_2159_v0_
                nw444_[int(4)] = d_2261_cf86_
                nw444_[int(5)] = d_2159_v0_
                nw444_[int(6)] = d_2260_cf87_
                nw444_[int(7)] = d_2261_cf86_
                nw444_[int(8)] = d_2261_cf86_
                nw444_[int(9)] = (d_2263_v78_).cf58
                nw444_[int(10)] = d_2260_cf87_
                nw444_[int(11)] = False
                nw444_[int(12)] = d_2259_cf88_
                nw444_[int(13)] = not(d_2260_cf87_)
                nw444_[int(14)] = d_2261_cf86_
                nw444_[int(15)] = d_2260_cf87_
                nw444_[int(16)] = d_2261_cf86_
                nw444_[int(17)] = d_2159_v0_
                nw444_[int(18)] = d_2260_cf87_
                nw444_[int(19)] = d_2159_v0_
                d_2264_v79_ = nw444_
                d_2265_v80_: D17
                d_2265_v80_ = D17_DC42(d_2264_v79_, p1)
                d_2265_v80_ = d_2265_v80_
                d_2266_v81_: _dafny.Array
                nw445_ = _dafny.Array(D11.default()(), 24)
                d_2266_v81_ = nw445_
                d_2267_v82_: _dafny.Array
                nw446_ = _dafny.Array(int(0), 9)
                d_2267_v82_ = nw446_
                d_2268_v83_: D18
                d_2268_v83_ = D18_DC45(d_2264_v79_, d_2159_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjvmy"))), p1)
                d_2269_v84_: D11
                d_2269_v84_ = D11_DC25(d_2261_cf86_, p1, d_2267_v82_, (d_2268_v83_).cf70)
                index443_ = default__.safeIndex(40, (d_2266_v81_).length(0))
                (d_2266_v81_)[index443_] = d_2269_v84_
                index444_ = default__.safeIndex(363, (d_2264_v79_).length(0))
                (d_2264_v79_)[index444_] = (False) and (not(d_2159_v0_))
                index445_ = default__.safeIndex(40, (d_2266_v81_).length(0))
                index446_ = default__.safeIndex(363, (d_2264_v79_).length(0))
                rhs434_ = d_2269_v84_
                rhs435_ = d_2259_cf88_
                lhs225_ = d_2266_v81_
                lhs226_ = default__.safeIndex(40, (d_2266_v81_).length(0))
                lhs227_ = d_2264_v79_
                lhs228_ = default__.safeIndex(363, (d_2264_v79_).length(0))
                lhs225_[lhs226_] = rhs434_
                lhs227_[lhs228_] = rhs435_
            elif source22_.is_DC56:
                d_2270___mcc_h5_ = source22_.cf90
                d_2271___mcc_h6_ = source22_.cf91
                d_2272___mcc_h7_ = source22_.cf92
                d_2273___mcc_h8_ = source22_.cf93
                d_2274_cf93_ = d_2273___mcc_h8_
                d_2275_cf92_ = d_2272___mcc_h7_
                d_2276_cf91_ = d_2271___mcc_h6_
                d_2277_cf90_ = d_2270___mcc_h5_
                d_2278_v85_: _dafny.Array
                def lambda152_(d_2279_i11_):
                    return (_dafny.MultiSet([D7_DC16(_dafny.CodePoint('x')), D7_DC16((self).f5)])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D7_DC16((self).f5)])))

                init86_ = lambda152_
                nw447_ = _dafny.Array(None, 18)
                for i0_86_ in range(nw447_.length(0)):
                    nw447_[i0_86_] = init86_(i0_86_)
                d_2278_v85_ = nw447_
                index447_ = default__.safeIndex(524, (d_2278_v85_).length(0))
                (d_2278_v85_)[index447_] = d_2275_cf92_
                index448_ = default__.safeIndex(524, (d_2278_v85_).length(0))
                (d_2278_v85_)[index448_] = not (d_2277_cf90_) or (not(d_2159_v0_))
                d_2280_v86_: C0
                nw448_ = C0()
                nw448_.ctor__()
                d_2280_v86_ = nw448_
                d_2281_v87_: _dafny.Map
                d_2281_v87_ = _dafny.Map({d_2274_cf93_: d_2280_v86_})
                d_2282_v88_: _dafny.Seq
                d_2282_v88_ = _dafny.SeqWithoutIsStrInference([d_2280_v86_, d_2280_v86_, ((d_2281_v87_)[d_2159_v0_] if (d_2159_v0_) in (d_2281_v87_) else d_2280_v86_), ((d_2281_v87_)[d_2159_v0_] if (d_2159_v0_) in (d_2281_v87_) else d_2280_v86_)])
                d_2283_v89_: _dafny.Array
                nw449_ = _dafny.Array(None, 24)
                nw449_[int(0)] = d_2280_v86_
                nw449_[int(1)] = d_2280_v86_
                nw449_[int(2)] = d_2280_v86_
                nw449_[int(3)] = d_2280_v86_
                nw449_[int(4)] = d_2280_v86_
                nw449_[int(5)] = d_2280_v86_
                nw449_[int(6)] = d_2280_v86_
                nw449_[int(7)] = (d_2282_v88_)[default__.safeIndex(p1, len(d_2282_v88_))]
                nw449_[int(8)] = d_2280_v86_
                nw449_[int(9)] = d_2280_v86_
                nw449_[int(10)] = d_2280_v86_
                nw449_[int(11)] = d_2280_v86_
                nw449_[int(12)] = d_2280_v86_
                nw449_[int(13)] = d_2280_v86_
                nw449_[int(14)] = d_2280_v86_
                nw449_[int(15)] = d_2280_v86_
                nw449_[int(16)] = d_2280_v86_
                nw449_[int(17)] = d_2280_v86_
                nw449_[int(18)] = d_2280_v86_
                nw449_[int(19)] = d_2280_v86_
                nw449_[int(20)] = d_2280_v86_
                nw449_[int(21)] = d_2280_v86_
                nw449_[int(22)] = d_2280_v86_
                nw449_[int(23)] = d_2280_v86_
                d_2283_v89_ = nw449_
                d_2284_v90_: _dafny.Array
                d_2284_v90_ = d_2283_v89_
                d_2285_v91_: _dafny.MultiSet
                d_2285_v91_ = _dafny.MultiSet([d_2283_v89_, d_2284_v90_, d_2284_v90_, d_2284_v90_])
                d_2285_v91_ = d_2285_v91_
                d_2174_v9_ = d_2174_v9_
                d_2286_v92_: _dafny.Map
                d_2286_v92_ = _dafny.Map({d_2159_v0_: d_2174_v9_})
                d_2287_v93_: _dafny.Map
                d_2287_v93_ = _dafny.Map({d_2275_cf92_: d_2274_cf93_})
                d_2288_v94_: D17
                d_2288_v94_ = D17_DC40(_dafny.Map({d_2286_v92_: d_2287_v93_}))
                d_2289_v95_: D17
                d_2289_v95_ = D17_DC43(d_2288_v94_)
                d_2290_v96_: _dafny.Array
                nw450_ = _dafny.Array(None, 3)
                nw450_[int(0)] = (self).f5
                nw450_[int(1)] = (self).f5
                nw450_[int(2)] = (self).f5
                d_2290_v96_ = nw450_
                d_2291_v97_: D7
                d_2291_v97_ = D7_DC16((self).f5)
                index449_ = default__.safeIndex(827, (d_2290_v96_).length(0))
                (d_2290_v96_)[index449_] = (d_2291_v97_).cf23
                index450_ = default__.safeIndex(524, (d_2278_v85_).length(0))
                index451_ = default__.safeIndex(827, (d_2290_v96_).length(0))
                rhs436_ = d_2289_v95_
                rhs437_ = (d_2276_cf91_ if not (True) or (d_2277_cf90_) else ((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwsytbx"))))
                rhs438_ = d_2159_v0_
                rhs439_ = (self).f5
                rhs440_ = (0) - (d_2174_v9_)
                lhs229_ = d_2278_v85_
                lhs230_ = default__.safeIndex(524, (d_2278_v85_).length(0))
                lhs231_ = d_2290_v96_
                lhs232_ = default__.safeIndex(827, (d_2290_v96_).length(0))
                d_2289_v95_ = rhs436_
                d_2276_cf91_ = rhs437_
                lhs229_[lhs230_] = rhs438_
                lhs231_[lhs232_] = rhs439_
                d_2174_v9_ = rhs440_
            elif source22_.is_DC57:
                d_2292_v98_: _dafny.Set
                d_2292_v98_ = _dafny.Set({p1, (d_2177_v12_)[default__.safeIndex(d_2174_v9_, len(d_2177_v12_))], p1})
                d_2293_v99_: _dafny.MultiSet
                d_2293_v99_ = _dafny.MultiSet([d_2159_v0_, d_2159_v0_])
                d_2294_v100_: _dafny.Map
                d_2294_v100_ = _dafny.Map({d_2292_v98_: d_2293_v99_})
                d_2174_v9_ = len(((_dafny.Map({d_2292_v98_: d_2293_v99_})) | (d_2294_v100_)) | (d_2294_v100_))
                d_2295_v101_: _dafny.Array
                nw451_ = _dafny.Array(None, 22)
                nw451_[int(0)] = d_2159_v0_
                nw451_[int(1)] = d_2159_v0_
                nw451_[int(2)] = d_2159_v0_
                nw451_[int(3)] = d_2159_v0_
                nw451_[int(4)] = True
                nw451_[int(5)] = d_2159_v0_
                nw451_[int(6)] = True
                nw451_[int(7)] = not(d_2159_v0_)
                nw451_[int(8)] = d_2159_v0_
                nw451_[int(9)] = not(d_2159_v0_)
                nw451_[int(10)] = d_2159_v0_
                nw451_[int(11)] = d_2159_v0_
                nw451_[int(12)] = d_2159_v0_
                nw451_[int(13)] = not(not(True))
                nw451_[int(14)] = d_2159_v0_
                nw451_[int(15)] = d_2159_v0_
                nw451_[int(16)] = d_2159_v0_
                nw451_[int(17)] = d_2159_v0_
                nw451_[int(18)] = d_2159_v0_
                nw451_[int(19)] = False
                nw451_[int(20)] = d_2159_v0_
                nw451_[int(21)] = d_2159_v0_
                d_2295_v101_ = nw451_
                d_2296_v102_: _dafny.Map
                d_2296_v102_ = _dafny.Map({p0: d_2295_v101_})
                d_2297_v103_: _dafny.Map
                d_2297_v103_ = _dafny.Map({d_2159_v0_: p1})
                d_2298_v104_: _dafny.Seq
                d_2298_v104_ = _dafny.SeqWithoutIsStrInference([d_2297_v103_])
                d_2296_v102_ = (d_2296_v102_).set(len(((d_2298_v104_)[default__.safeIndex(83, len(d_2298_v104_))]) | (d_2297_v103_)), d_2295_v101_)
                d_2299_v105_: _dafny.Set
                d_2299_v105_ = _dafny.Set({d_2159_v0_})
                d_2300_v106_: _dafny.Seq
                d_2300_v106_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2159_v0_}), (_dafny.Set({d_2159_v0_})) - (d_2299_v105_)])
                rhs441_ = not (not(d_2159_v0_)) or (d_2159_v0_)
                rhs442_ = len(d_2300_v106_)
                rhs443_ = default__.safeModuloInt(p0, p1)
                d_2159_v0_ = rhs441_
                d_2174_v9_ = rhs442_
                d_2174_v9_ = rhs443_
                d_2301_v107_: _dafny.Array
                nw452_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_2301_v107_ = nw452_
                d_2302_v108_: _dafny.Array
                nw453_ = _dafny.Array(int(0), 21)
                d_2302_v108_ = nw453_
                index452_ = default__.safeIndex(891, (d_2301_v107_).length(0))
                (d_2301_v107_)[index452_] = d_2302_v108_
                index453_ = default__.safeIndex(891, (d_2301_v107_).length(0))
                nw454_ = _dafny.Array(int(0), 8)
                (d_2301_v107_)[index453_] = nw454_
            elif True:
                d_2303___mcc_h9_ = source22_.cf84
                d_2304_cf84_ = d_2303___mcc_h9_
                d_2305_v109_: _dafny.Array
                def lambda153_(d_2306_v0_):
                    def lambda154_(d_2307_i12_):
                        return d_2306_v0_

                    return lambda154_

                init87_ = lambda153_(d_2159_v0_)
                nw455_ = _dafny.Array(None, 6)
                for i0_87_ in range(nw455_.length(0)):
                    nw455_[i0_87_] = init87_(i0_87_)
                d_2305_v109_ = nw455_
                d_2308_v110_: _dafny.Map
                d_2308_v110_ = _dafny.Map({d_2159_v0_: D17_DC42(d_2305_v109_, p0)})
                d_2309_v111_: _dafny.Map
                d_2309_v111_ = _dafny.Map({True: not(d_2159_v0_)})
                d_2310_v112_: D17
                d_2310_v112_ = D17_DC42(d_2305_v109_, p1)
                pat_let_tv54_ = d_2177_v12_
                def iife166_(_pat_let53_0):
                    def iife167_(d_2311_dt__update__tmp_h1_):
                        def iife168_(_pat_let54_0):
                            def iife169_(d_2312_dt__update_hcf65_h0_):
                                return D17_DC42((d_2311_dt__update__tmp_h1_).cf64, d_2312_dt__update_hcf65_h0_)
                            return iife169_(_pat_let54_0)
                        return iife168_(len(pat_let_tv54_))
                    return iife167_(_pat_let53_0)
                d_2308_v110_ = (d_2308_v110_).set((self).fm5(d_2174_v9_, d_2159_v0_, p1, len(d_2309_v111_), globalState), iife166_(d_2310_v112_))
                d_2313_v113_: _dafny.MultiSet
                d_2313_v113_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f5 for d_2314_i13_ in range(default__.abs(774))]), d_2169_v7_])
                r0 = (default__.fm46(globalState)).issubset((d_2313_v113_).set((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))], default__.abs(903)))
                d_2315_v114_: _dafny.Seq
                d_2315_v114_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2159_v0_: d_2159_v0_}), d_2309_v111_])
                d_2316_v115_: _dafny.Map
                d_2316_v115_ = _dafny.Map({d_2159_v0_: 99})
                d_2317_v116_: _dafny.Map
                d_2317_v116_ = _dafny.Map({d_2159_v0_: ((d_2315_v114_)[default__.safeIndex(len(d_2316_v115_), len(d_2315_v114_))]) | (d_2309_v111_)})
                rhs444_ = 595
                rhs445_ = d_2317_v116_
                d_2174_v9_ = rhs444_
                d_2317_v116_ = rhs445_
                d_2174_v9_ = d_2174_v9_
            d_2318_v118_: C6
            nw456_ = C6()
            def iife170_():
                coll60_ = _dafny.Map()
                compr_60_: int
                for compr_60_ in _dafny.IntegerRange(-726, 473):
                    d_2319_v117_: int = compr_60_
                    if ((-726) <= (d_2319_v117_)) and ((d_2319_v117_) < (473)):
                        coll60_[(d_2319_v117_) + (d_2174_v9_)] = (0) - (p0)
                return _dafny.Map(coll60_)
            nw456_.ctor__((iife170_()
            ) | (_dafny.Map({p1: len(d_2177_v12_)})))
            d_2318_v118_ = nw456_
            r0 = d_2159_v0_
            d_2320_v119_: _dafny.Array
            nw457_ = _dafny.Array(False, 15)
            d_2320_v119_ = nw457_
            d_2321_v120_: _dafny.Map
            d_2321_v120_ = _dafny.Map({-552: d_2320_v119_})
            d_2322_v121_: _dafny.MultiSet
            d_2322_v121_ = _dafny.MultiSet([((d_2321_v120_)[p0] if (p0) in (d_2321_v120_) else d_2320_v119_)])
            d_2323_v122_: D25
            d_2323_v122_ = D25_DC59(d_2322_v121_)
            source23_ = d_2323_v122_
            if source23_.is_DC60:
                d_2324___mcc_h10_ = source23_.cf96
                d_2325___mcc_h11_ = source23_.cf97
                d_2326___mcc_h12_ = source23_.cf98
                d_2327_cf98_ = d_2326___mcc_h12_
                d_2328_cf97_ = d_2325___mcc_h11_
                d_2329_cf96_ = d_2324___mcc_h10_
                d_2330_v123_: _dafny.Map
                d_2330_v123_ = _dafny.Map({d_2329_cf96_: d_2174_v9_})
                d_2331_v124_: _dafny.Seq
                d_2331_v124_ = _dafny.SeqWithoutIsStrInference([(d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]])
                rhs446_ = d_2169_v7_
                rhs447_ = (d_2331_v124_)[default__.safeIndex(p0, len(d_2331_v124_))]
                rhs448_ = (d_2330_v123_).set(d_2159_v0_, d_2174_v9_)
                rhs449_ = d_2159_v0_
                r1 = rhs446_
                d_2169_v7_ = rhs447_
                d_2330_v123_ = rhs448_
                d_2328_cf97_ = rhs449_
                def iife171_():
                    coll61_ = _dafny.Map()
                    compr_61_: int
                    for compr_61_ in ((d_2318_v118_).f10).keys.Elements:
                        d_2332_v125_: int = compr_61_
                        if (d_2332_v125_) in ((d_2318_v118_).f10):
                            coll61_[(d_2332_v125_) - (p1)] = ((d_2327_cf98_ if d_2329_cf96_ else d_2169_v7_)).set(default__.safeIndex(p0, len((d_2327_cf98_ if d_2329_cf96_ else d_2169_v7_))), (self).f5)
                    return _dafny.Map(coll61_)
                d_2175_v10_ = iife171_()
                
                d_2333_v126_: _dafny.Array
                def lambda155_(d_2334_v9_):
                    def lambda156_(d_2335_i14_):
                        return (d_2335_i14_) + (d_2334_v9_)

                    return lambda156_

                init88_ = lambda155_(d_2174_v9_)
                nw458_ = _dafny.Array(None, 6)
                for i0_88_ in range(nw458_.length(0)):
                    nw458_[i0_88_] = init88_(i0_88_)
                d_2333_v126_ = nw458_
                index454_ = default__.safeIndex(999, (d_2333_v126_).length(0))
                (d_2333_v126_)[index454_] = p1
                index455_ = default__.safeIndex(999, (d_2333_v126_).length(0))
                (d_2333_v126_)[index455_] = ((self).fm4(d_2177_v12_, globalState)) * (p1)
                index456_ = default__.safeIndex(999, (d_2333_v126_).length(0))
                rhs450_ = default__.fm14(globalState)
                rhs451_ = (len(d_2327_cf98_)) * ((p0) * (p1))
                lhs233_ = d_2333_v126_
                lhs234_ = default__.safeIndex(999, (d_2333_v126_).length(0))
                lhs233_[lhs234_] = rhs450_
                d_2174_v9_ = rhs451_
            elif True:
                d_2336___mcc_h13_ = source23_.cf95
                d_2337_cf95_ = d_2336___mcc_h13_
                d_2338_v127_: _dafny.Array
                nw459_ = _dafny.Array(_dafny.Seq({}), 25)
                d_2338_v127_ = nw459_
                d_2339_v128_: _dafny.Map
                d_2339_v128_ = _dafny.Map({not(d_2159_v0_): (d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]})
                d_2340_v129_: _dafny.Seq
                d_2340_v129_ = _dafny.SeqWithoutIsStrInference([d_2169_v7_, ((d_2339_v128_)[d_2159_v0_] if (d_2159_v0_) in (d_2339_v128_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qna")))])
                index457_ = default__.safeIndex(140, (d_2338_v127_).length(0))
                (d_2338_v127_)[index457_] = d_2340_v129_
                index458_ = default__.safeIndex(140, (d_2338_v127_).length(0))
                (d_2338_v127_)[index458_] = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkkyq"))).set(default__.safeIndex((_dafny.MultiSet([True, d_2159_v0_, d_2159_v0_])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkkyq")))), (self).f5), (d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))], d_2169_v7_])
                d_2341_v130_: _dafny.MultiSet
                d_2341_v130_ = _dafny.MultiSet([False, False, d_2159_v0_])
                d_2342_v131_: _dafny.Map
                d_2342_v131_ = _dafny.Map({d_2341_v130_: (d_2177_v12_)[default__.safeIndex(d_2174_v9_, len(d_2177_v12_))]})
                d_2343_v132_: _dafny.Array
                nw460_ = _dafny.Array(None, 11)
                nw460_[int(0)] = p0
                nw460_[int(1)] = p0
                nw460_[int(2)] = -17
                nw460_[int(3)] = p1
                nw460_[int(4)] = d_2174_v9_
                nw460_[int(5)] = p0
                nw460_[int(6)] = p1
                nw460_[int(7)] = len(d_2342_v131_)
                nw460_[int(8)] = d_2174_v9_
                nw460_[int(9)] = p1
                nw460_[int(10)] = p0
                d_2343_v132_ = nw460_
                d_2344_v133_: _dafny.MultiSet
                d_2344_v133_ = _dafny.MultiSet([d_2174_v9_, p0, len(_dafny.SeqWithoutIsStrInference([d_2343_v132_, d_2343_v132_])), p0])
                r2 = (d_2344_v133_) | (d_2344_v133_)
                d_2174_v9_ = len((d_2318_v118_).f10)
                d_2345_v134_: D15
                d_2345_v134_ = D15_DC35(d_2174_v9_)
                d_2346_v135_: D15
                d_2346_v135_ = D15_DC36(d_2345_v134_)
                d_2347_v136_: _dafny.Set
                d_2347_v136_ = _dafny.Set({d_2320_v119_})
                d_2348_v137_: _dafny.Set
                d_2348_v137_ = d_2347_v136_
                d_2349_v138_: _dafny.Array
                nw461_ = _dafny.Array(_dafny.Map({}), 23)
                d_2349_v138_ = nw461_
                d_2350_v139_: D4
                d_2350_v139_ = D4_DC10(d_2348_v137_, p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2351_i15_ in range(default__.abs(834))]), d_2349_v138_)
                d_2352_v140_: T1
                nw462_ = C4()
                nw462_.ctor__(d_2174_v9_, d_2174_v9_, (self).f4, _dafny.CodePoint('w'))
                d_2352_v140_ = nw462_
                d_2353_v141_: _dafny.Map
                d_2353_v141_ = _dafny.Map({(d_2350_v139_).cf12: d_2352_v140_})
                d_2354_v142_: _dafny.Map
                d_2354_v142_ = _dafny.Map({d_2346_v135_: d_2353_v141_})
                d_2354_v142_ = (d_2354_v142_).set(d_2346_v135_, d_2353_v141_)
        elif True:
            d_2355_v143_: _dafny.Set
            d_2355_v143_ = _dafny.Set({p1, p0})
            d_2356_v145_: D14
            d_2356_v145_ = D14_DC32(len(d_2169_v7_), d_2159_v0_, d_2159_v0_)
            d_2357_v146_: _dafny.Array
            nw463_ = _dafny.Array(int(0), 18)
            d_2357_v146_ = nw463_
            d_2358_v147_: _dafny.Map
            d_2358_v147_ = _dafny.Map({d_2159_v0_: d_2357_v146_})
            d_2359_v150_: _dafny.Array
            nw464_ = _dafny.Array(None, 17)
            nw464_[int(0)] = (_dafny.Set({d_2174_v9_})) | (d_2355_v143_)
            nw464_[int(1)] = (_dafny.Set({p0}) if d_2159_v0_ else d_2355_v143_)
            def iife172_():
                coll62_ = _dafny.Set()
                compr_62_: int
                for compr_62_ in _dafny.IntegerRange(-228, 726):
                    d_2360_v144_: int = compr_62_
                    if ((-228) <= (d_2360_v144_)) and ((d_2360_v144_) < (726)):
                        coll62_ = coll62_.union(_dafny.Set([(d_2360_v144_) - (p0)]))
                return _dafny.Set(coll62_)
            nw464_[int(2)] = iife172_()
            
            nw464_[int(3)] = _dafny.Set({d_2174_v9_, p0})
            nw464_[int(4)] = _dafny.Set({p0})
            nw464_[int(5)] = d_2355_v143_
            nw464_[int(6)] = d_2355_v143_
            nw464_[int(7)] = d_2355_v143_
            nw464_[int(8)] = d_2355_v143_
            nw464_[int(9)] = _dafny.Set({(d_2356_v145_).cf49, d_2174_v9_})
            nw464_[int(10)] = d_2355_v143_
            def iife173_():
                coll63_ = _dafny.Set()
                compr_63_: int
                for compr_63_ in (_dafny.MultiSet([len(d_2358_v147_)])).Elements:
                    d_2361_v148_: int = compr_63_
                    if (d_2361_v148_) in (_dafny.MultiSet([len(d_2358_v147_)])):
                        coll63_ = coll63_.union(_dafny.Set([(d_2361_v148_) * (len(_dafny.Map({False: _dafny.Set({True})})))]))
                return _dafny.Set(coll63_)
            nw464_[int(11)] = (d_2355_v143_).intersection(iife173_()
            )
            nw464_[int(12)] = d_2355_v143_
            def iife174_():
                coll64_ = _dafny.Set()
                compr_64_: int
                for compr_64_ in (_dafny.SeqWithoutIsStrInference([d_2174_v9_])).Elements:
                    d_2362_v149_: int = compr_64_
                    if (d_2362_v149_) in (_dafny.SeqWithoutIsStrInference([d_2174_v9_])):
                        coll64_ = coll64_.union(_dafny.Set([(d_2362_v149_) + ((D20_DC48(False, _dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([not(True), not(True), True]), (_dafny.MultiSet([True, False])).cardinality)).cf77)]))
                return _dafny.Set(coll64_)
            nw464_[int(13)] = iife174_()
            
            nw464_[int(14)] = d_2355_v143_
            nw464_[int(15)] = d_2355_v143_
            nw464_[int(16)] = d_2355_v143_
            d_2359_v150_ = nw464_
            nw465_ = _dafny.Array(_dafny.Set({}), 11)
            d_2359_v150_ = nw465_
            d_2363_v151_: _dafny.Array
            def lambda157_(d_2364_v0_):
                def lambda158_(d_2365_i16_):
                    return d_2364_v0_

                return lambda158_

            init89_ = lambda157_(d_2159_v0_)
            nw466_ = _dafny.Array(None, 13)
            for i0_89_ in range(nw466_.length(0)):
                nw466_[i0_89_] = init89_(i0_89_)
            d_2363_v151_ = nw466_
            index459_ = default__.safeIndex(307, (d_2363_v151_).length(0))
            (d_2363_v151_)[index459_] = not(d_2159_v0_)
            d_2366_v152_: _dafny.Array
            nw467_ = _dafny.Array(D25.default()(), 13)
            d_2366_v152_ = nw467_
            d_2367_v153_: _dafny.Seq
            d_2367_v153_ = _dafny.SeqWithoutIsStrInference([d_2366_v152_])
            index460_ = default__.safeIndex(307, (d_2363_v151_).length(0))
            (d_2363_v151_)[index460_] = (d_2366_v152_) not in (d_2367_v153_)
            d_2368_v154_: D16
            d_2368_v154_ = D16_DC38((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))], (d_2363_v151_)[default__.safeIndex(307, (d_2363_v151_).length(0))], p1, (d_2363_v151_)[default__.safeIndex(307, (d_2363_v151_).length(0))])
            d_2369_v155_: _dafny.MultiSet
            d_2369_v155_ = _dafny.MultiSet([default__.fm41(d_2159_v0_, (_dafny.MultiSet([d_2174_v9_, d_2174_v9_, p0])).cardinality, 974, len((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]), globalState), d_2368_v154_, d_2368_v154_])
            d_2159_v0_ = (d_2369_v155_) != (d_2369_v155_)
            r0 = d_2159_v0_
            d_2370_v156_: _dafny.Set
            d_2370_v156_ = _dafny.Set({d_2363_v151_})
            d_2371_v157_: _dafny.Map
            d_2371_v157_ = _dafny.Map({(d_2363_v151_)[default__.safeIndex(307, (d_2363_v151_).length(0))]: d_2370_v156_})
            d_2372_v158_: _dafny.Map
            d_2373_v159_: int
            d_2374_v160_: bool
            d_2375_v161_: _dafny.Map
            out25_: _dafny.Map
            out26_: int
            out27_: bool
            out28_: _dafny.Map
            out25_, out26_, out27_, out28_ = (self).m5(d_2169_v7_, d_2371_v157_, (p1) - (d_2174_v9_), globalState)
            d_2372_v158_ = out25_
            d_2373_v159_ = out26_
            d_2374_v160_ = out27_
            d_2375_v161_ = out28_
        d_2376_v162_: _dafny.Map
        d_2376_v162_ = _dafny.Map({d_2159_v0_: d_2169_v7_})
        d_2377_v163_: _dafny.MultiSet
        d_2377_v163_ = _dafny.MultiSet([len((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]), d_2174_v9_, d_2174_v9_, d_2174_v9_, len(d_2376_v162_)])
        d_2378_v164_: D8
        d_2378_v164_ = D8_DC18(d_2377_v163_)
        source24_ = d_2378_v164_
        if source24_.is_DC19:
            d_2379___mcc_h14_ = source24_.cf29
            d_2380___mcc_h15_ = source24_.cf30
            d_2381_cf30_ = d_2380___mcc_h15_
            d_2382_cf29_ = d_2379___mcc_h14_
            default__.m0(len((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]), d_2381_cf30_, globalState)
            d_2159_v0_ = d_2381_cf30_
            d_2174_v9_ = default__.safeDivisionInt(p1, d_2382_cf29_)
            d_2383_v165_: _dafny.Map
            d_2383_v165_ = _dafny.Map({d_2177_v12_: False})
            d_2383_v165_ = (d_2383_v165_).set((_dafny.SeqWithoutIsStrInference([p1 for d_2384_i17_ in range(default__.abs(-998))])) + (d_2177_v12_), (d_2177_v12_) != (d_2177_v12_))
        elif True:
            d_2385___mcc_h16_ = source24_.cf28
            d_2386_cf28_ = d_2385___mcc_h16_
            d_2387_v166_: _dafny.Array
            nw468_ = _dafny.Array(int(0), 1)
            d_2387_v166_ = nw468_
            d_2388_v167_: _dafny.Seq
            d_2388_v167_ = _dafny.SeqWithoutIsStrInference([d_2159_v0_])
            index461_ = default__.safeIndex(460, (d_2387_v166_).length(0))
            (d_2387_v166_)[index461_] = (d_2174_v9_) * (len(d_2388_v167_))
            index462_ = default__.safeIndex(727, (d_2387_v166_).length(0))
            (d_2387_v166_)[index462_] = p1
            d_2389_v168_: D7
            d_2389_v168_ = D7_DC17(d_2159_v0_, d_2159_v0_, d_2159_v0_, d_2159_v0_)
            index463_ = default__.safeIndex(460, (d_2387_v166_).length(0))
            index464_ = default__.safeIndex(727, (d_2387_v166_).length(0))
            rhs452_ = not(d_2159_v0_)
            rhs453_ = (d_2159_v0_ if (d_2389_v168_).cf24 else d_2159_v0_)
            rhs454_ = -611
            rhs455_ = default__.safeModuloInt(d_2174_v9_, p1)
            lhs235_ = d_2387_v166_
            lhs236_ = default__.safeIndex(460, (d_2387_v166_).length(0))
            lhs237_ = d_2387_v166_
            lhs238_ = default__.safeIndex(727, (d_2387_v166_).length(0))
            d_2159_v0_ = rhs452_
            d_2159_v0_ = rhs453_
            lhs235_[lhs236_] = rhs454_
            lhs237_[lhs238_] = rhs455_
            if d_2159_v0_:
                index465_ = default__.safeIndex(460, (d_2387_v166_).length(0))
                rhs456_ = d_2387_v166_
                rhs457_ = default__.fm14(globalState)
                rhs458_ = False
                rhs459_ = (d_2388_v167_)[default__.safeIndex(d_2174_v9_, len(d_2388_v167_))]
                lhs239_ = d_2387_v166_
                lhs240_ = default__.safeIndex(460, (d_2387_v166_).length(0))
                d_2387_v166_ = rhs456_
                lhs239_[lhs240_] = rhs457_
                r0 = rhs458_
                r0 = rhs459_
                d_2390_v169_: _dafny.Array
                def lambda159_(d_2391_v7_, d_2392_v0_, d_2393_v167_):
                    def lambda160_(d_2394_i18_):
                        return _dafny.SeqWithoutIsStrInference([D5_DC11(_dafny.Map({(D23_DC55(d_2391_v7_, d_2392_v0_, True, d_2392_v0_, d_2393_v167_)).cf88: (self).f5})), D5_DC11(_dafny.Map({d_2392_v0_: (self).f5})), D5_DC11(_dafny.Map({False: (self).f5})), D5_DC11(_dafny.Map({True: _dafny.CodePoint('u')}))])

                    return lambda160_

                init90_ = lambda159_(d_2169_v7_, d_2159_v0_, d_2388_v167_)
                nw469_ = _dafny.Array(None, 8)
                for i0_90_ in range(nw469_.length(0)):
                    nw469_[i0_90_] = init90_(i0_90_)
                d_2390_v169_ = nw469_
                d_2395_v170_: _dafny.Map
                d_2395_v170_ = _dafny.Map({d_2159_v0_: (self).f5})
                d_2396_v171_: D5
                d_2396_v171_ = D5_DC11(d_2395_v170_)
                d_2397_v172_: _dafny.Seq
                d_2397_v172_ = _dafny.SeqWithoutIsStrInference([d_2396_v171_, d_2396_v171_, D5_DC11(d_2395_v170_)])
                index466_ = default__.safeIndex(151, (d_2390_v169_).length(0))
                (d_2390_v169_)[index466_] = d_2397_v172_
                pat_let_tv55_ = d_2395_v170_
                pat_let_tv56_ = d_2159_v0_
                index467_ = default__.safeIndex(151, (d_2390_v169_).length(0))
                def iife175_(_pat_let55_0):
                    def iife176_(d_2398_dt__update__tmp_h2_):
                        def iife177_(_pat_let56_0):
                            def iife178_(d_2399_dt__update_hcf15_h0_):
                                return D5_DC11(d_2399_dt__update_hcf15_h0_)
                            return iife178_(_pat_let56_0)
                        return iife177_((pat_let_tv55_).set(pat_let_tv56_, (self).f5))
                    return iife176_(_pat_let55_0)
                (d_2390_v169_)[index467_] = (_dafny.SeqWithoutIsStrInference([d_2396_v171_])) + (_dafny.SeqWithoutIsStrInference([d_2396_v171_, d_2396_v171_, d_2396_v171_, iife175_(d_2396_v171_)]))
                d_2159_v0_ = not(d_2159_v0_)
                d_2400_v173_: _dafny.Map
                d_2400_v173_ = _dafny.Map({d_2159_v0_: not(d_2159_v0_)})
                d_2400_v173_ = (d_2400_v173_).set(d_2159_v0_, d_2159_v0_)
                d_2401_v174_: _dafny.Array
                nw470_ = _dafny.Array(False, 6)
                d_2401_v174_ = nw470_
                d_2402_v175_: _dafny.Set
                d_2402_v175_ = _dafny.Set({d_2401_v174_, d_2401_v174_})
                d_2403_v176_: _dafny.Set
                d_2403_v176_ = d_2402_v175_
                d_2404_v177_: _dafny.Map
                d_2404_v177_ = _dafny.Map({d_2159_v0_: d_2403_v176_})
                d_2405_v178_: _dafny.Seq
                d_2405_v178_ = _dafny.SeqWithoutIsStrInference([d_2404_v177_])
                d_2406_v179_: _dafny.Map
                d_2407_v180_: int
                d_2408_v181_: bool
                d_2409_v182_: _dafny.Map
                out29_: _dafny.Map
                out30_: int
                out31_: bool
                out32_: _dafny.Map
                out29_, out30_, out31_, out32_ = (self).m5(default__.fm12(globalState), (d_2405_v178_)[default__.safeIndex(p1, len(d_2405_v178_))], p0, globalState)
                d_2406_v179_ = out29_
                d_2407_v180_ = out30_
                d_2408_v181_ = out31_
                d_2409_v182_ = out32_
            elif True:
                index468_ = default__.safeIndex(460, (d_2387_v166_).length(0))
                (d_2387_v166_)[index468_] = (self).fm4(default__.fm31(623, p0, globalState), globalState)
                d_2410_v183_: _dafny.Array
                nw471_ = _dafny.Array(False, 14)
                d_2410_v183_ = nw471_
                index469_ = default__.safeIndex(779, (d_2410_v183_).length(0))
                (d_2410_v183_)[index469_] = d_2159_v0_
                d_2411_v184_: _dafny.Map
                d_2411_v184_ = _dafny.Map({len((d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]): p0})
                index470_ = default__.safeIndex(779, (d_2410_v183_).length(0))
                (d_2410_v183_)[index470_] = (default__.safeModuloInt((d_2387_v166_)[default__.safeIndex(460, (d_2387_v166_).length(0))], (d_2387_v166_)[default__.safeIndex(460, (d_2387_v166_).length(0))])) != (((d_2411_v184_)[(d_2387_v166_)[default__.safeIndex(460, (d_2387_v166_).length(0))]] if ((d_2387_v166_)[default__.safeIndex(460, (d_2387_v166_).length(0))]) in (d_2411_v184_) else p1))
                index471_ = default__.safeIndex(779, (d_2410_v183_).length(0))
                (d_2410_v183_)[index471_] = (d_2159_v0_) or (d_2159_v0_)
                index472_ = default__.safeIndex(460, (d_2387_v166_).length(0))
                (d_2387_v166_)[index472_] = (d_2177_v12_)[default__.safeIndex((0) - ((p1) * (d_2174_v9_)), len(d_2177_v12_))]
                d_2387_v166_ = d_2387_v166_
            d_2412_v185_: _dafny.Array
            nw472_ = _dafny.Array(_dafny.MultiSet({}), 1)
            d_2412_v185_ = nw472_
            rhs460_ = d_2159_v0_
            rhs461_ = d_2412_v185_
            r0 = rhs460_
            d_2412_v185_ = rhs461_
            (globalState).f2 = ((d_2177_v12_) + (d_2177_v12_) if d_2159_v0_ else d_2177_v12_)
        d_2413_v186_: _dafny.Array
        nw473_ = _dafny.Array(False, 8)
        d_2413_v186_ = nw473_
        d_2414_v187_: _dafny.Seq
        d_2414_v187_ = _dafny.SeqWithoutIsStrInference([d_2413_v186_])
        r0 = not((d_2413_v186_) not in ((d_2414_v187_) + ((d_2414_v187_).set(default__.safeIndex(p1, len(d_2414_v187_)), d_2413_v186_))))
        d_2415_v188_: _dafny.Map
        d_2415_v188_ = _dafny.Map({d_2413_v186_: (d_2170_v8_)[default__.safeIndex(27, (d_2170_v8_).length(0))]})
        r1 = ((d_2415_v188_)[d_2413_v186_] if (d_2413_v186_) in (d_2415_v188_) else default__.fm12(globalState))
        r2 = d_2377_v163_
        r3 = default__.fm0((0) - ((p0) + (d_2174_v9_)), p0, globalState)
        return r0, r1, r2, r3

    def m5(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: bool = False
        r3: _dafny.Map = _dafny.Map({})
        d_2416_v0_: _dafny.Array
        def lambda161_(d_2417_i0_):
            return default__.safeDivisionInt(d_2417_i0_, 277)

        init91_ = lambda161_
        nw474_ = _dafny.Array(None, 13)
        for i0_91_ in range(nw474_.length(0)):
            nw474_[i0_91_] = init91_(i0_91_)
        d_2416_v0_ = nw474_
        d_2416_v0_ = d_2416_v0_
        r1 = (p2) + (len(p0))
        d_2418_v1_: str
        d_2418_v1_ = _dafny.CodePoint('k')
        d_2418_v1_ = d_2418_v1_
        d_2419_v2_: _dafny.Seq
        d_2419_v2_ = _dafny.SeqWithoutIsStrInference([p2])
        if (((d_2419_v2_) + (d_2419_v2_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uf"))), len((d_2419_v2_) + (d_2419_v2_))), p2)) != (d_2419_v2_):
            d_2420_v3_: bool
            d_2420_v3_ = True
            d_2421_v4_: _dafny.Seq
            d_2421_v4_ = _dafny.SeqWithoutIsStrInference([False, d_2420_v3_, d_2420_v3_, d_2420_v3_])
            d_2422_v5_: _dafny.Map
            d_2422_v5_ = _dafny.Map({d_2420_v3_: d_2421_v4_})
            d_2423_v6_: _dafny.MultiSet
            d_2423_v6_ = _dafny.MultiSet([((d_2422_v5_)[d_2420_v3_] if (d_2420_v3_) in (d_2422_v5_) else d_2421_v4_), d_2421_v4_])
            index473_ = default__.safeIndex(310, (d_2416_v0_).length(0))
            (d_2416_v0_)[index473_] = (d_2423_v6_).cardinality
            index474_ = default__.safeIndex(310, (d_2416_v0_).length(0))
            (d_2416_v0_)[index474_] = (((0) - (p2)) + ((0) - (p2))) * (-393)
            d_2424_v7_: _dafny.Array
            nw475_ = _dafny.Array(None, 3)
            nw475_[int(0)] = (d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))]
            nw475_[int(1)] = (0) - ((d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))])
            nw475_[int(2)] = (d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))]
            d_2424_v7_ = nw475_
            d_2425_v8_: _dafny.Seq
            d_2425_v8_ = _dafny.SeqWithoutIsStrInference([d_2424_v7_, d_2424_v7_, d_2424_v7_, d_2424_v7_, d_2424_v7_])
            d_2425_v8_ = _dafny.SeqWithoutIsStrInference([d_2424_v7_])
            r1 = (d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))]
            d_2426_v9_: _dafny.Array
            nw476_ = _dafny.Array(False, 26)
            d_2426_v9_ = nw476_
            index475_ = default__.safeIndex(827, (d_2426_v9_).length(0))
            (d_2426_v9_)[index475_] = (p2) != ((d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))])
            index476_ = default__.safeIndex(827, (d_2426_v9_).length(0))
            (d_2426_v9_)[index476_] = (p2) > (default__.safeModuloInt((d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))], (d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))]))
            index477_ = default__.safeIndex(310, (d_2416_v0_).length(0))
            (d_2416_v0_)[index477_] = (d_2416_v0_)[default__.safeIndex(310, (d_2416_v0_).length(0))]
        elif True:
            r1 = default__.safeDivisionInt(p2, 408)
            d_2427_v10_: bool
            d_2427_v10_ = True
            r2 = d_2427_v10_
            d_2428_v11_: _dafny.MultiSet
            d_2428_v11_ = _dafny.MultiSet([p2])
            d_2429_v12_: _dafny.Set
            d_2429_v12_ = _dafny.Set({d_2428_v11_, d_2428_v11_, _dafny.MultiSet(d_2419_v2_)})
            d_2430_v13_: _dafny.Seq
            d_2430_v13_ = _dafny.SeqWithoutIsStrInference([d_2429_v12_])
            d_2431_v14_: _dafny.Seq
            d_2431_v14_ = _dafny.SeqWithoutIsStrInference([default__.fm22(d_2418_v1_, d_2427_v10_, len((d_2430_v13_)[default__.safeIndex(p2, len(d_2430_v13_))]), globalState)])
            d_2432_v15_: _dafny.Seq
            d_2432_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gq"))
            rhs462_ = d_2427_v10_
            rhs463_ = d_2431_v14_
            rhs464_ = d_2432_v15_
            d_2427_v10_ = rhs462_
            d_2431_v14_ = rhs463_
            d_2432_v15_ = rhs464_
            d_2433_v16_: _dafny.Seq
            d_2433_v16_ = _dafny.SeqWithoutIsStrInference([d_2427_v10_, d_2427_v10_])
            d_2434_v17_: _dafny.Set
            d_2434_v17_ = _dafny.Set({not(d_2427_v10_)})
            d_2435_v18_: _dafny.Array
            nw477_ = _dafny.Array(None, 7)
            nw477_[int(0)] = d_2433_v16_
            nw477_[int(1)] = (d_2433_v16_) + (d_2433_v16_)
            nw477_[int(2)] = (d_2433_v16_) + (d_2433_v16_)
            nw477_[int(3)] = (d_2433_v16_) + (d_2433_v16_)
            nw477_[int(4)] = (default__.fm21(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emr"))), len(_dafny.Map({d_2427_v10_: d_2427_v10_})), d_2434_v17_, globalState)).set(default__.safeIndex((0) - (p2), len(default__.fm21(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emr"))), len(_dafny.Map({d_2427_v10_: d_2427_v10_})), d_2434_v17_, globalState))), d_2427_v10_)
            nw477_[int(5)] = default__.fm21(p2, p2, d_2434_v17_, globalState)
            nw477_[int(6)] = d_2433_v16_
            d_2435_v18_ = nw477_
            index478_ = default__.safeIndex(929, (d_2435_v18_).length(0))
            (d_2435_v18_)[index478_] = d_2433_v16_
            index479_ = default__.safeIndex(929, (d_2435_v18_).length(0))
            (d_2435_v18_)[index479_] = ((d_2433_v16_) + (d_2433_v16_)) + (d_2433_v16_)
            d_2436_v19_: _dafny.Array
            nw478_ = _dafny.Array(False, 6)
            d_2436_v19_ = nw478_
            (globalState).f3 = d_2436_v19_
        d_2437_v20_: C0
        nw479_ = C0()
        nw479_.ctor__()
        d_2437_v20_ = nw479_
        d_2438_v21_: _dafny.Set
        d_2438_v21_ = _dafny.Set({d_2437_v20_, d_2437_v20_})
        d_2439_v22_: _dafny.Map
        d_2439_v22_ = _dafny.Map({p2: len(d_2438_v21_)})
        r3 = d_2439_v22_
        hi19_ = len(_dafny.SeqWithoutIsStrInference([(self).fm4(d_2419_v2_, globalState), p2, 600]))
        for d_2440_i1_ in range(p2, hi19_):
            d_2441_v23_: bool
            d_2441_v23_ = False
            if not ((_dafny.SeqWithoutIsStrInference([p2])) == (d_2419_v2_)) or ((D7_DC17((d_2437_v20_).fm3(846, p2, globalState), d_2441_v23_, d_2441_v23_, d_2441_v23_)).cf27):
                d_2442_v24_: _dafny.Set
                d_2442_v24_ = _dafny.Set({d_2440_i1_})
                d_2443_v25_: _dafny.Map
                d_2443_v25_ = _dafny.Map({default__.fm0(277, d_2440_i1_, globalState): d_2442_v24_})
                d_2444_v26_: _dafny.Map
                d_2444_v26_ = _dafny.Map({len(d_2443_v25_): _dafny.SeqWithoutIsStrInference([(0) - (p2)])})
                d_2444_v26_ = (d_2444_v26_).set(d_2440_i1_, d_2419_v2_)
                d_2445_v27_: _dafny.MultiSet
                d_2445_v27_ = _dafny.MultiSet([d_2441_v23_, d_2441_v23_])
                d_2446_v28_: _dafny.Map
                d_2446_v28_ = _dafny.Map({((d_2439_v22_)[932] if (932) in (d_2439_v22_) else d_2440_i1_): d_2445_v27_})
                d_2447_v29_: _dafny.Map
                d_2447_v29_ = _dafny.Map({len(d_2446_v28_): d_2441_v23_})
                d_2448_v30_: _dafny.MultiSet
                d_2448_v30_ = _dafny.MultiSet([False, ((d_2447_v29_)[p2] if (p2) in (d_2447_v29_) else not(d_2441_v23_)), d_2441_v23_, not((len(d_2419_v2_)) > (len((p0).set(default__.safeIndex(len(_dafny.Map({d_2441_v23_: d_2440_i1_})), len(p0)), _dafny.CodePoint('c'))))), d_2441_v23_])
                d_2445_v27_ = ((d_2448_v30_) | (d_2445_v27_)) | (_dafny.MultiSet([d_2441_v23_]))
                r2 = d_2441_v23_
                d_2449_v31_: _dafny.Array
                nw480_ = _dafny.Array(False, 2)
                d_2449_v31_ = nw480_
                index480_ = default__.safeIndex(531, (d_2449_v31_).length(0))
                (d_2449_v31_)[index480_] = True
                index481_ = default__.safeIndex(531, (d_2449_v31_).length(0))
                (d_2449_v31_)[index481_] = d_2441_v23_
                r1 = p2
            elif True:
                d_2441_v23_ = (len(d_2419_v2_)) == ((len(p0) if not(d_2441_v23_) else -897))
                r1 = p2
                d_2441_v23_ = d_2441_v23_
                r1 = p2
                r1 = (d_2440_i1_) + (p2)
            d_2441_v23_ = d_2441_v23_
            if d_2441_v23_:
                rhs465_ = d_2441_v23_
                rhs466_ = (_dafny.Map({d_2440_i1_: (self).fm4(d_2419_v2_, globalState)})) | (d_2439_v22_)
                rhs467_ = d_2418_v1_
                rhs468_ = (d_2440_i1_) != ((0) - ((d_2440_i1_) * (p2)))
                rhs469_ = d_2441_v23_
                r2 = rhs465_
                d_2439_v22_ = rhs466_
                d_2418_v1_ = rhs467_
                d_2441_v23_ = rhs468_
                d_2441_v23_ = rhs469_
                d_2450_v32_: C7
                nw481_ = C7()
                nw481_.ctor__()
                d_2450_v32_ = nw481_
                d_2451_v33_: _dafny.Array
                nw482_ = _dafny.Array(None, 17)
                nw482_[int(0)] = d_2450_v32_
                nw482_[int(1)] = d_2450_v32_
                nw482_[int(2)] = d_2450_v32_
                nw482_[int(3)] = d_2450_v32_
                nw482_[int(4)] = d_2450_v32_
                nw482_[int(5)] = d_2450_v32_
                nw482_[int(6)] = d_2450_v32_
                nw482_[int(7)] = d_2450_v32_
                nw482_[int(8)] = d_2450_v32_
                nw482_[int(9)] = d_2450_v32_
                nw482_[int(10)] = d_2450_v32_
                nw482_[int(11)] = d_2450_v32_
                nw482_[int(12)] = d_2450_v32_
                nw482_[int(13)] = d_2450_v32_
                nw482_[int(14)] = d_2450_v32_
                nw482_[int(15)] = d_2450_v32_
                nw482_[int(16)] = d_2450_v32_
                d_2451_v33_ = nw482_
                d_2452_v34_: _dafny.Map
                d_2452_v34_ = _dafny.Map({p2: d_2451_v33_})
                d_2453_v35_: _dafny.Array
                nw483_ = _dafny.Array(None, 29)
                nw483_[int(0)] = d_2451_v33_
                nw483_[int(1)] = ((d_2452_v34_)[(0) - (d_2440_i1_)] if ((0) - (d_2440_i1_)) in (d_2452_v34_) else d_2451_v33_)
                nw483_[int(2)] = d_2451_v33_
                nw483_[int(3)] = d_2451_v33_
                nw483_[int(4)] = d_2451_v33_
                nw483_[int(5)] = d_2451_v33_
                nw483_[int(6)] = d_2451_v33_
                nw483_[int(7)] = d_2451_v33_
                nw483_[int(8)] = d_2451_v33_
                nw483_[int(9)] = d_2451_v33_
                nw483_[int(10)] = d_2451_v33_
                nw483_[int(11)] = ((d_2452_v34_)[d_2440_i1_] if (d_2440_i1_) in (d_2452_v34_) else d_2451_v33_)
                nw483_[int(12)] = d_2451_v33_
                nw483_[int(13)] = d_2451_v33_
                nw483_[int(14)] = d_2451_v33_
                nw483_[int(15)] = (d_2451_v33_ if True else d_2451_v33_)
                nw483_[int(16)] = d_2451_v33_
                nw483_[int(17)] = d_2451_v33_
                nw483_[int(18)] = d_2451_v33_
                nw483_[int(19)] = d_2451_v33_
                nw483_[int(20)] = d_2451_v33_
                nw483_[int(21)] = d_2451_v33_
                nw483_[int(22)] = d_2451_v33_
                nw483_[int(23)] = d_2451_v33_
                nw483_[int(24)] = d_2451_v33_
                nw483_[int(25)] = d_2451_v33_
                nw483_[int(26)] = d_2451_v33_
                nw483_[int(27)] = d_2451_v33_
                nw483_[int(28)] = d_2451_v33_
                d_2453_v35_ = nw483_
                d_2453_v35_ = (d_2453_v35_ if d_2441_v23_ else d_2453_v35_)
                d_2454_v36_: _dafny.Seq
                d_2454_v36_ = _dafny.SeqWithoutIsStrInference([d_2441_v23_, d_2441_v23_, d_2441_v23_, d_2441_v23_, d_2441_v23_])
                d_2455_v37_: D15
                d_2455_v37_ = D15_DC36(D15_DC34(d_2454_v36_))
                d_2456_v38_: D15
                d_2456_v38_ = D15_DC34(d_2454_v36_)
                rhs470_ = d_2441_v23_
                rhs471_ = (D15_DC36(d_2456_v38_) if d_2441_v23_ else D15_DC36(d_2456_v38_))
                rhs472_ = d_2441_v23_
                rhs473_ = (d_2418_v1_ if d_2441_v23_ else _dafny.CodePoint('y'))
                r2 = rhs470_
                d_2455_v37_ = rhs471_
                d_2441_v23_ = rhs472_
                d_2418_v1_ = rhs473_
                index482_ = default__.safeIndex(389, (d_2416_v0_).length(0))
                (d_2416_v0_)[index482_] = (len(p0) if d_2441_v23_ else d_2440_i1_)
                index483_ = default__.safeIndex(389, (d_2416_v0_).length(0))
                (d_2416_v0_)[index483_] = (0) - ((self).fm4(d_2419_v2_, globalState))
                r2 = d_2441_v23_
            elif True:
                r1 = default__.safeModuloInt(d_2440_i1_, (len(d_2439_v22_)) * (238))
                d_2457_v39_: _dafny.Array
                nw484_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_2457_v39_ = nw484_
                index484_ = default__.safeIndex(240, (d_2457_v39_).length(0))
                (d_2457_v39_)[index484_] = (p0) + (p0)
                index485_ = default__.safeIndex(240, (d_2457_v39_).length(0))
                (d_2457_v39_)[index485_] = default__.fm12(globalState)
                d_2416_v0_ = d_2416_v0_
                d_2458_v40_: _dafny.Map
                d_2458_v40_ = _dafny.Map({(0) - (p2): d_2441_v23_})
                d_2458_v40_ = (d_2458_v40_).set(p2, d_2441_v23_)
                r2 = d_2441_v23_
            r1 = default__.safeModuloInt(default__.fm13(d_2441_v23_, d_2441_v23_, default__.fm1(d_2440_i1_, p2, False, d_2441_v23_, globalState), globalState), (len(p0)) - (d_2440_i1_))
        d_2459_v41_: bool
        d_2459_v41_ = True
        d_2460_v42_: _dafny.Map
        d_2460_v42_ = _dafny.Map({d_2459_v41_: p2})
        d_2461_v43_: _dafny.Map
        d_2461_v43_ = _dafny.Map({d_2459_v41_: default__.safeDivisionInt(((d_2460_v42_)[True] if (True) in (d_2460_v42_) else p2), p2)})
        r0 = d_2461_v43_
        r1 = p2
        r2 = (len((_dafny.SeqWithoutIsStrInference([(self).f5 for d_2462_i2_ in range(default__.abs(682))])).set(default__.safeIndex(len(p0), len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_2463_i2_ in range(default__.abs(682))]))), d_2418_v1_))) >= (p2)
        r3 = (d_2439_v22_) | (d_2439_v22_)
        return r0, r1, r2, r3


class C10(T1, T0):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f8, f9, f4, f5):
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f4 = f4
        (self)._f5 = f5

    def fm4(self, p0, globalState):
        return default__.safeDivisionInt((0) - ((4) * ((self).f8)), (self).f8)

    def fm5(self, p0, p1, p2, p3, globalState):
        return not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chgqrdurm"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr"))))

    def fm3(self, p0, p1, globalState):
        return False

    def fm8(self, p0, p1, p2, globalState):
        return (self).f8

    def fm9(self, p0, p1, globalState):
        return ((_dafny.Map({(self).f8: (self).f8})) | (_dafny.Map({(self).f8: (0) - ((self).f8)}))) | ((_dafny.Map({(self).f8: (self).f8})) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f9]))).cardinality: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kckfhdycy"))): (self).f9}))})))

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        if (self).f9:
            r1 = (-907) * ((self).f8)
            d_2464_v0_: _dafny.Map
            d_2464_v0_ = _dafny.Map({(self).f8: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxaegu"))})
            d_2465_v2_: D2
            d_2465_v2_ = D2_DC4((self).f9)
            d_2466_v3_: _dafny.Map
            def iife179_():
                coll65_ = _dafny.Map()
                compr_65_: int
                for compr_65_ in _dafny.IntegerRange(443, 341):
                    d_2467_v1_: int = compr_65_
                    if ((443) <= (d_2467_v1_)) and ((d_2467_v1_) < (341)):
                        coll65_[default__.safeDivisionInt(d_2467_v1_, 550)] = _dafny.SeqWithoutIsStrInference([(self).f5 for d_2468_i0_ in range(default__.abs(16))])
                return _dafny.Map(coll65_)
            def iife180_(_pat_let57_0):
                def iife181_(d_2469_dt__update__tmp_h0_):
                    def iife182_(_pat_let58_0):
                        def iife183_(d_2470_dt__update_hcf5_h0_):
                            return D2_DC4(d_2470_dt__update_hcf5_h0_)
                        return iife183_(_pat_let58_0)
                    return iife182_((self).f9)
                return iife181_(_pat_let57_0)
            d_2466_v3_ = _dafny.Map({(0) - (len((d_2464_v0_) | (iife179_()
            ))): iife180_(d_2465_v2_)})
            d_2466_v3_ = (d_2466_v3_) | (d_2466_v3_)
            d_2471_v4_: _dafny.Seq
            d_2471_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpeh"))
            d_2471_v4_ = d_2471_v4_
            r1 = (self).f8
            r1 = (self).f8
        elif True:
            d_2472_v5_: _dafny.Array
            nw485_ = _dafny.Array(False, 26)
            d_2472_v5_ = nw485_
            d_2473_v6_: _dafny.Set
            d_2473_v6_ = _dafny.Set({d_2472_v5_})
            d_2474_v7_: _dafny.Set
            d_2474_v7_ = d_2473_v6_
            d_2475_v8_: _dafny.Map
            d_2475_v8_ = _dafny.Map({d_2474_v7_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2476_i1_ in range(default__.abs(27))]))})
            d_2477_v9_: _dafny.Seq
            d_2477_v9_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5])
            d_2478_v10_: _dafny.Seq
            d_2478_v10_ = _dafny.SeqWithoutIsStrInference([(d_2475_v8_).set(d_2474_v7_, len(d_2477_v9_))])
            index486_ = default__.safeIndex(860, (d_2472_v5_).length(0))
            (d_2472_v5_)[index486_] = False
            d_2479_v11_: bool
            d_2479_v11_ = False
            d_2480_v12_: _dafny.Map
            d_2480_v12_ = _dafny.Map({d_2479_v11_: (self).f5})
            d_2481_v13_: D5
            d_2481_v13_ = D5_DC11(d_2480_v12_)
            d_2482_v15_: _dafny.Map
            d_2482_v15_ = _dafny.Map({(D5_DC13(d_2477_v9_, (self).f8, (self).f9)).cf17: (self).f8})
            d_2483_v16_: _dafny.MultiSet
            d_2483_v16_ = _dafny.MultiSet([(self).f8, len(d_2482_v15_)])
            d_2484_v17_: D2
            d_2484_v17_ = D2_DC4(not(d_2479_v11_))
            index487_ = default__.safeIndex(860, (d_2472_v5_).length(0))
            def iife184_():
                coll66_ = _dafny.Map()
                compr_66_: int
                for compr_66_ in (d_2483_v16_).Elements:
                    d_2485_v14_: int = compr_66_
                    if (d_2485_v14_) in (d_2483_v16_):
                        coll66_[default__.safeModuloInt(d_2485_v14_, (0) - ((self).f8))] = (_dafny.MultiSet([(self).f5, (self).f5, (self).f5, (self).f5])).cardinality
                return _dafny.Map(coll66_)
            rhs474_ = (_dafny.SeqWithoutIsStrInference([d_2475_v8_])).set(default__.safeIndex((self).fm8((d_2481_v13_).cf15, iife184_()
            , _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltvuc")), globalState), len(_dafny.SeqWithoutIsStrInference([d_2475_v8_]))), d_2475_v8_)
            rhs475_ = (d_2484_v17_).cf5
            rhs476_ = (self).f9
            rhs477_ = d_2479_v11_
            lhs241_ = d_2472_v5_
            lhs242_ = default__.safeIndex(860, (d_2472_v5_).length(0))
            d_2478_v10_ = rhs474_
            lhs241_[lhs242_] = rhs475_
            d_2479_v11_ = rhs476_
            d_2479_v11_ = rhs477_
            d_2486_v18_: C0
            nw486_ = C0()
            nw486_.ctor__()
            d_2486_v18_ = nw486_
            d_2482_v15_ = (d_2482_v15_).set((self).f8, (self).f8)
            d_2479_v11_ = (d_2486_v18_).fm3((self).f8, -990, globalState)
            d_2479_v11_ = not((d_2477_v9_) == (d_2477_v9_))
        r1 = (self).f8
        d_2487_v19_: bool
        d_2487_v19_ = True
        d_2488_v20_: _dafny.Array
        nw487_ = _dafny.Array(int(0), 25)
        d_2488_v20_ = nw487_
        index488_ = default__.safeIndex(59, (d_2488_v20_).length(0))
        (d_2488_v20_)[index488_] = -895
        d_2489_v21_: _dafny.Map
        d_2489_v21_ = _dafny.Map({(self).f8: d_2487_v19_})
        d_2490_v22_: _dafny.Map
        d_2490_v22_ = _dafny.Map({(self).f8: len(d_2489_v21_)})
        d_2491_v23_: _dafny.Map
        d_2491_v23_ = _dafny.Map({(self).f8: d_2490_v22_})
        d_2492_v24_: D14
        d_2492_v24_ = D14_DC30((d_2491_v23_).set((self).f8, d_2490_v22_))
        d_2493_v25_: _dafny.Seq
        d_2493_v25_ = _dafny.SeqWithoutIsStrInference([d_2492_v24_])
        d_2494_v26_: _dafny.Map
        d_2494_v26_ = _dafny.Map({(self).f8: (self).f5})
        d_2495_v27_: _dafny.Seq
        d_2495_v27_ = _dafny.SeqWithoutIsStrInference([d_2493_v25_])
        index489_ = default__.safeIndex(59, (d_2488_v20_).length(0))
        rhs478_ = (self).fm3(default__.safeDivisionInt((self).f8, default__.fm14(globalState)), (len(d_2494_v26_)) + ((self).f8), globalState)
        rhs479_ = (self).f8
        rhs480_ = (d_2495_v27_)[default__.safeIndex((self).f8, len(d_2495_v27_))]
        rhs481_ = (default__.safeModuloInt((self).f8, 229)) - ((self).f8)
        lhs243_ = d_2488_v20_
        lhs244_ = default__.safeIndex(59, (d_2488_v20_).length(0))
        d_2487_v19_ = rhs478_
        lhs243_[lhs244_] = rhs479_
        d_2493_v25_ = rhs480_
        r1 = rhs481_
        d_2496_v28_: _dafny.Array
        nw488_ = _dafny.Array(_dafny.Seq({}), 27)
        d_2496_v28_ = nw488_
        d_2497_v29_: _dafny.Seq
        d_2497_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rt"))
        d_2498_v30_: _dafny.Seq
        d_2498_v30_ = _dafny.SeqWithoutIsStrInference([(d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))], ((d_2490_v22_)[len(d_2497_v29_)] if (len(d_2497_v29_)) in (d_2490_v22_) else (d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]), (d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]])
        index490_ = default__.safeIndex(650, (d_2496_v28_).length(0))
        (d_2496_v28_)[index490_] = d_2498_v30_
        d_2499_v31_: _dafny.Map
        d_2499_v31_ = _dafny.Map({(d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]: d_2497_v29_})
        index491_ = default__.safeIndex(650, (d_2496_v28_).length(0))
        (d_2496_v28_)[index491_] = default__.fm31(len(((d_2499_v31_)[(self).f8] if ((self).f8) in (d_2499_v31_) else d_2497_v29_)), 690, globalState)
        d_2500_v32_: _dafny.Seq
        d_2500_v32_ = _dafny.SeqWithoutIsStrInference([d_2487_v19_])
        d_2501_v33_: _dafny.Map
        d_2501_v33_ = _dafny.Map({d_2500_v32_: (d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]})
        d_2501_v33_ = ((_dafny.Map({d_2500_v32_: 891})) | (d_2501_v33_)) | ((d_2501_v33_).set(d_2500_v32_, (self).f8))
        d_2502_v34_: _dafny.Array
        nw489_ = _dafny.Array(False, 11)
        d_2502_v34_ = nw489_
        d_2503_v35_: D18
        d_2503_v35_ = D18_DC45(d_2502_v34_, (self).f9, (self).f8, -585)
        d_2504_v36_: _dafny.Seq
        d_2504_v36_ = _dafny.SeqWithoutIsStrInference([d_2502_v34_, (d_2503_v35_).cf68])
        d_2505_i2_: int
        d_2505_i2_ = 0
        with _dafny.label("9"):
            while ((d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]) != (len(d_2504_v36_)):
                with _dafny.c_label("9"):
                    if (d_2505_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_2505_i2_ = (d_2505_i2_) + (1)
                    d_2506_v37_: _dafny.Array
                    def lambda162_(d_2507_v29_, d_2508_v20_):
                        def lambda163_(d_2509_i3_):
                            return (d_2507_v29_)[default__.safeIndex((d_2508_v20_)[default__.safeIndex(59, (d_2508_v20_).length(0))], len(d_2507_v29_))]

                        return lambda163_

                    init92_ = lambda162_(d_2497_v29_, d_2488_v20_)
                    nw490_ = _dafny.Array(None, 20)
                    for i0_92_ in range(nw490_.length(0)):
                        nw490_[i0_92_] = init92_(i0_92_)
                    d_2506_v37_ = nw490_
                    d_2506_v37_ = d_2506_v37_
                    d_2510_v38_: _dafny.MultiSet
                    d_2510_v38_ = _dafny.MultiSet([(d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))], (self).f8])
                    index492_ = default__.safeIndex(59, (d_2488_v20_).length(0))
                    rhs482_ = d_2488_v20_
                    rhs483_ = (d_2510_v38_).ispropersubset(d_2510_v38_)
                    rhs484_ = ((d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]) * ((d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))])
                    lhs245_ = d_2488_v20_
                    lhs246_ = default__.safeIndex(59, (d_2488_v20_).length(0))
                    d_2488_v20_ = rhs482_
                    d_2487_v19_ = rhs483_
                    lhs245_[lhs246_] = rhs484_
                    d_2511_v39_: _dafny.Seq
                    d_2511_v39_ = _dafny.SeqWithoutIsStrInference([(d_2498_v30_ if d_2487_v19_ else d_2498_v30_)])
                    d_2511_v39_ = _dafny.SeqWithoutIsStrInference([d_2498_v30_, (d_2498_v30_).set(default__.safeIndex((0) - ((d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]), len(d_2498_v30_)), (self).f8), d_2498_v30_])
                    d_2512_v40_: str
                    d_2512_v40_ = _dafny.CodePoint('o')
                    d_2512_v40_ = (self).f5
                    pass
            pass
        r0 = (((d_2496_v28_)[default__.safeIndex(650, (d_2496_v28_).length(0))]) + (_dafny.SeqWithoutIsStrInference([(self).f8 for d_2513_i4_ in range(default__.abs(332))]))).set(default__.safeIndex((d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))], len(((d_2496_v28_)[default__.safeIndex(650, (d_2496_v28_).length(0))]) + (_dafny.SeqWithoutIsStrInference([(self).f8 for d_2514_i4_ in range(default__.abs(332))])))), 330)
        d_2515_v41_: D11
        d_2515_v41_ = D11_DC25((self).f9, (d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))], d_2488_v20_, (self).f8)
        d_2516_v42_: D11
        d_2516_v42_ = D11_DC26(d_2515_v41_)
        d_2517_v43_: D11
        d_2517_v43_ = D11_DC26(d_2516_v42_)
        d_2518_v44_: D11
        d_2518_v44_ = D11_DC26(d_2516_v42_)
        d_2519_v45_: D11
        d_2519_v45_ = D11_DC26(d_2517_v43_)
        d_2520_v46_: _dafny.MultiSet
        d_2520_v46_ = _dafny.MultiSet([(d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))], (d_2488_v20_)[default__.safeIndex(59, (d_2488_v20_).length(0))]])
        d_2521_v47_: _dafny.Map
        d_2521_v47_ = _dafny.Map({d_2519_v45_: d_2520_v46_})
        r1 = ((0) - ((0) - (len((d_2521_v47_) | (d_2521_v47_))))) + ((self).f8)
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        if (self).f9:
            if (self).f9:
                d_2522_v0_: _dafny.Array
                def lambda164_(d_2523_i0_):
                    return False

                init93_ = lambda164_
                nw491_ = _dafny.Array(None, 10)
                for i0_93_ in range(nw491_.length(0)):
                    nw491_[i0_93_] = init93_(i0_93_)
                d_2522_v0_ = nw491_
                d_2524_v1_: _dafny.Map
                d_2524_v1_ = _dafny.Map({d_2522_v0_: (self).f9})
                d_2524_v1_ = (d_2524_v1_).set(d_2522_v0_, (self).f9)
                d_2525_v2_: _dafny.Map
                d_2525_v2_ = _dafny.Map({(self).f9: (self).f9})
                d_2526_v3_: _dafny.Seq
                d_2526_v3_ = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f8])
                d_2527_v5_: _dafny.Seq
                d_2527_v5_ = _dafny.SeqWithoutIsStrInference([d_2525_v2_])
                d_2528_v6_: _dafny.Array
                nw492_ = _dafny.Array(None, 8)
                nw492_[int(0)] = len(_dafny.Set({(self).f8, (self).f8, (self).f8, (0) - ((self).f8), default__.fm14(globalState)}))
                nw492_[int(1)] = (self).f8
                nw492_[int(2)] = len(d_2525_v2_)
                nw492_[int(3)] = (self).f8
                nw492_[int(4)] = default__.safeDivisionInt((self).f8, (self).f8)
                nw492_[int(5)] = (d_2526_v3_)[default__.safeIndex((self).f8, len(d_2526_v3_))]
                def iife185_():
                    coll67_ = _dafny.Map()
                    compr_67_: _dafny.Map
                    for compr_67_ in (d_2527_v5_).Elements:
                        d_2529_v4_: _dafny.Map = compr_67_
                        if (d_2529_v4_) in (d_2527_v5_):
                            coll67_[d_2529_v4_] = (self).f8
                    return _dafny.Map(coll67_)
                nw492_[int(6)] = len((iife185_()
                ) | (default__.fm47(True, (self).f8, globalState)))
                nw492_[int(7)] = (self).f8
                d_2528_v6_ = nw492_
                index493_ = default__.safeIndex(62, (d_2528_v6_).length(0))
                def iife186_():
                    coll68_ = _dafny.Set()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(172, 986):
                        d_2530_v7_: int = compr_68_
                        if ((172) <= (d_2530_v7_)) and ((d_2530_v7_) < (986)):
                            coll68_ = coll68_.union(_dafny.Set([default__.safeModuloInt(d_2530_v7_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f5 for d_2531_i1_ in range(default__.abs(248))]))).cardinality)]))
                    return _dafny.Set(coll68_)
                (d_2528_v6_)[index493_] = (0) - (len(iife186_()
                ))
                index494_ = default__.safeIndex(62, (d_2528_v6_).length(0))
                (d_2528_v6_)[index494_] = ((0) - ((self).f8)) * ((self).f8)
                index495_ = default__.safeIndex(62, (d_2528_v6_).length(0))
                (d_2528_v6_)[index495_] = (0) - ((d_2528_v6_)[default__.safeIndex(62, (d_2528_v6_).length(0))])
                d_2532_v8_: _dafny.Seq
                d_2532_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([340, (self).f8])])
                d_2533_v9_: _dafny.Seq
                d_2533_v9_ = _dafny.SeqWithoutIsStrInference([(self).f9, (len(d_2532_v8_)) >= (len(default__.fm29((d_2528_v6_)[default__.safeIndex(62, (d_2528_v6_).length(0))], (self).f9, (self).f8, globalState))), (self).f9])
                d_2533_v9_ = d_2533_v9_
                index496_ = default__.safeIndex(62, (d_2528_v6_).length(0))
                (d_2528_v6_)[index496_] = (-975) * ((d_2528_v6_)[default__.safeIndex(62, (d_2528_v6_).length(0))])
            elif True:
                d_2534_v10_: _dafny.Seq
                d_2534_v10_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])
                d_2535_v11_: _dafny.Set
                d_2535_v11_ = _dafny.Set({len(_dafny.Map({(self).f8: (self).f9}))})
                d_2536_v12_: _dafny.Map
                d_2536_v12_ = _dafny.Map({(self).f9: (self).f9})
                d_2537_v13_: _dafny.Seq
                d_2537_v13_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5])
                d_2538_v14_: _dafny.Map
                d_2538_v14_ = _dafny.Map({(self).f8: len(d_2537_v13_)})
                d_2539_v15_: C6
                nw493_ = C6()
                nw493_.ctor__(d_2538_v14_)
                d_2539_v15_ = nw493_
                d_2540_v16_: _dafny.Set
                d_2540_v16_ = _dafny.Set({d_2539_v15_, d_2539_v15_})
                d_2541_v17_: _dafny.Map
                d_2541_v17_ = _dafny.Map({(self).f9: (self).f8})
                d_2542_v18_: _dafny.MultiSet
                d_2542_v18_ = _dafny.MultiSet([len(d_2541_v17_)])
                d_2543_v19_: _dafny.Seq
                d_2543_v19_ = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f8])
                d_2544_v20_: _dafny.Array
                def lambda165_(d_2545_i2_):
                    return (self).f9

                init94_ = lambda165_
                nw494_ = _dafny.Array(None, 24)
                for i0_94_ in range(nw494_.length(0)):
                    nw494_[i0_94_] = init94_(i0_94_)
                d_2544_v20_ = nw494_
                d_2546_v21_: _dafny.Seq
                d_2546_v21_ = _dafny.SeqWithoutIsStrInference([d_2544_v20_])
                d_2547_v22_: D2
                d_2547_v22_ = D2_DC4((self).f9)
                d_2548_v23_: _dafny.Array
                nw495_ = _dafny.Array(None, 14)
                nw495_[int(0)] = not(not((self).f9))
                nw495_[int(1)] = (len(d_2534_v10_)) < ((self).f8)
                nw495_[int(2)] = ((self).f8) > ((self).f8)
                nw495_[int(3)] = (d_2535_v11_).issubset(d_2535_v11_)
                nw495_[int(4)] = ((d_2536_v12_)[(self).f9] if ((self).f9) in (d_2536_v12_) else (self).f9)
                nw495_[int(5)] = (d_2539_v15_) not in (d_2540_v16_)
                nw495_[int(6)] = ((self).f9) and ((self).f9)
                nw495_[int(7)] = ((d_2542_v18_).cardinality) <= ((0) - ((self).f8))
                nw495_[int(8)] = not(not(((self).f8) <= ((d_2543_v19_)[default__.safeIndex((self).f8, len(d_2543_v19_))])))
                nw495_[int(9)] = (d_2544_v20_) not in (d_2546_v21_)
                nw495_[int(10)] = (self).f9
                nw495_[int(11)] = (self).f9
                nw495_[int(12)] = (d_2547_v22_).cf5
                nw495_[int(13)] = False
                d_2548_v23_ = nw495_
                d_2549_v24_: _dafny.Map
                d_2549_v24_ = _dafny.Map({(self).f9: d_2543_v19_})
                index497_ = default__.safeIndex(964, (d_2548_v23_).length(0))
                (d_2548_v23_)[index497_] = (_dafny.Map({(self).f9: d_2543_v19_})) == (d_2549_v24_)
                index498_ = default__.safeIndex(964, (d_2548_v23_).length(0))
                (d_2548_v23_)[index498_] = (self).f9
                d_2550_v25_: _dafny.MultiSet
                d_2550_v25_ = _dafny.MultiSet([((d_2548_v23_)[default__.safeIndex(964, (d_2548_v23_).length(0))]) == ((self).f9), (self).f9, (self).f9])
                d_2550_v25_ = (d_2550_v25_) - (d_2550_v25_)
                d_2551_v26_: _dafny.Map
                d_2551_v26_ = _dafny.Map({(d_2538_v14_) | (d_2538_v14_): (self).f9})
                d_2551_v26_ = (d_2551_v26_).set(_dafny.Map({(self).f8: len(d_2537_v13_)}), (d_2548_v23_)[default__.safeIndex(964, (d_2548_v23_).length(0))])
                d_2552_v27_: int
                d_2552_v27_ = 569
                d_2553_v28_: C0
                nw496_ = C0()
                nw496_.ctor__()
                d_2553_v28_ = nw496_
                d_2554_v29_: _dafny.Array
                def lambda166_(d_2555_v13_):
                    def lambda167_(d_2556_i3_):
                        return (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2557_i4_ in range(default__.abs(573))])) + (d_2555_v13_)

                    return lambda167_

                init95_ = lambda166_(d_2537_v13_)
                nw497_ = _dafny.Array(None, 9)
                for i0_95_ in range(nw497_.length(0)):
                    nw497_[i0_95_] = init95_(i0_95_)
                d_2554_v29_ = nw497_
                index499_ = default__.safeIndex(189, (d_2554_v29_).length(0))
                (d_2554_v29_)[index499_] = d_2537_v13_
                index500_ = default__.safeIndex(189, (d_2554_v29_).length(0))
                rhs485_ = len(((d_2543_v19_).set(default__.safeIndex(len(d_2543_v19_), len(d_2543_v19_)), (0) - (d_2552_v27_))) + (default__.fm31(d_2552_v27_, (self).f8, globalState)))
                rhs486_ = d_2553_v28_
                rhs487_ = ((d_2537_v13_ if default__.fm1(d_2552_v27_, (self).f8, (self).f9, True, globalState) else _dafny.SeqWithoutIsStrInference([(self).f5 for d_2558_i5_ in range(default__.abs(20))])) if (self).f9 else d_2537_v13_)
                lhs247_ = d_2554_v29_
                lhs248_ = default__.safeIndex(189, (d_2554_v29_).length(0))
                d_2552_v27_ = rhs485_
                d_2553_v28_ = rhs486_
                lhs247_[lhs248_] = rhs487_
                index501_ = default__.safeIndex(964, (d_2548_v23_).length(0))
                rhs488_ = (((self).f9 if (self).f9 else (self).f9) if (self).f9 else True)
                rhs489_ = (d_2543_v19_) + (_dafny.SeqWithoutIsStrInference([(0) - (len(d_2535_v11_)), d_2552_v27_, (0) - (d_2552_v27_)]))
                rhs490_ = d_2552_v27_
                lhs249_ = d_2548_v23_
                lhs250_ = default__.safeIndex(964, (d_2548_v23_).length(0))
                lhs251_ = globalState
                lhs249_[lhs250_] = rhs488_
                lhs251_.f2 = rhs489_
                d_2552_v27_ = rhs490_
            if not (default__.fm1((self).f8, (self).f8, (self).f9, (self).f9, globalState)) or ((self).f9):
                d_2559_v30_: int
                d_2559_v30_ = 729
                d_2560_v31_: _dafny.Seq
                d_2560_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gc"))
                d_2561_v32_: _dafny.Array
                nw498_ = _dafny.Array(int(0), 12)
                d_2561_v32_ = nw498_
                d_2562_v33_: _dafny.Map
                d_2562_v33_ = _dafny.Map({d_2559_v30_: d_2561_v32_})
                d_2563_v34_: _dafny.Seq
                d_2563_v34_ = _dafny.SeqWithoutIsStrInference([d_2561_v32_])
                d_2564_v35_: _dafny.Array
                nw499_ = _dafny.Array(None, 23)
                nw499_[int(0)] = ((d_2562_v33_)[d_2559_v30_] if (d_2559_v30_) in (d_2562_v33_) else d_2561_v32_)
                nw499_[int(1)] = d_2561_v32_
                nw499_[int(2)] = d_2561_v32_
                nw499_[int(3)] = d_2561_v32_
                nw499_[int(4)] = d_2561_v32_
                nw499_[int(5)] = d_2561_v32_
                nw499_[int(6)] = d_2561_v32_
                nw499_[int(7)] = d_2561_v32_
                nw499_[int(8)] = d_2561_v32_
                nw499_[int(9)] = d_2561_v32_
                nw499_[int(10)] = (d_2563_v34_)[default__.safeIndex(d_2559_v30_, len(d_2563_v34_))]
                nw499_[int(11)] = d_2561_v32_
                nw499_[int(12)] = d_2561_v32_
                nw499_[int(13)] = (d_2561_v32_ if (self).f9 else d_2561_v32_)
                nw499_[int(14)] = d_2561_v32_
                nw499_[int(15)] = d_2561_v32_
                nw499_[int(16)] = d_2561_v32_
                nw499_[int(17)] = d_2561_v32_
                nw499_[int(18)] = d_2561_v32_
                nw499_[int(19)] = d_2561_v32_
                nw499_[int(20)] = d_2561_v32_
                nw499_[int(21)] = d_2561_v32_
                nw499_[int(22)] = d_2561_v32_
                d_2564_v35_ = nw499_
                index502_ = default__.safeIndex(576, (d_2564_v35_).length(0))
                (d_2564_v35_)[index502_] = d_2561_v32_
                index503_ = default__.safeIndex(32, (d_2561_v32_).length(0))
                (d_2561_v32_)[index503_] = (0) - ((self).f8)
                d_2565_v36_: str
                d_2565_v36_ = _dafny.CodePoint('t')
                index504_ = default__.safeIndex(576, (d_2564_v35_).length(0))
                index505_ = default__.safeIndex(32, (d_2561_v32_).length(0))
                rhs491_ = (self).f8
                rhs492_ = d_2560_v31_
                rhs493_ = d_2561_v32_
                rhs494_ = (self).f8
                rhs495_ = (self).f5
                lhs252_ = d_2564_v35_
                lhs253_ = default__.safeIndex(576, (d_2564_v35_).length(0))
                lhs254_ = d_2561_v32_
                lhs255_ = default__.safeIndex(32, (d_2561_v32_).length(0))
                d_2559_v30_ = rhs491_
                d_2560_v31_ = rhs492_
                lhs252_[lhs253_] = rhs493_
                lhs254_[lhs255_] = rhs494_
                d_2565_v36_ = rhs495_
                d_2566_v37_: bool
                d_2566_v37_ = False
                d_2566_v37_ = (((d_2561_v32_)[default__.safeIndex(32, (d_2561_v32_).length(0))]) != ((self).f8)) == (d_2566_v37_)
                d_2566_v37_ = ((d_2561_v32_)[default__.safeIndex(32, (d_2561_v32_).length(0))]) > ((d_2561_v32_)[default__.safeIndex(32, (d_2561_v32_).length(0))])
                d_2567_v38_: C9
                nw500_ = C9()
                nw500_.ctor__((self).f4, (self).f5)
                d_2567_v38_ = nw500_
                d_2568_v39_: _dafny.Map
                d_2568_v39_ = _dafny.Map({(self).f5: (0) - ((self).f8)})
                d_2568_v39_ = d_2568_v39_
            elif True:
                d_2569_v40_: _dafny.Seq
                d_2569_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fodfeqy"))
                d_2570_v41_: bool
                d_2570_v41_ = True
                d_2571_v42_: T0
                nw501_ = C7()
                nw501_.ctor__()
                d_2571_v42_ = nw501_
                d_2572_v43_: _dafny.Map
                d_2572_v43_ = _dafny.Map({(self).f8: d_2571_v42_})
                d_2573_v44_: _dafny.Seq
                d_2573_v44_ = _dafny.SeqWithoutIsStrInference([d_2571_v42_])
                d_2574_v45_: D16
                d_2574_v45_ = D16_DC38(d_2569_v40_, (self).f9, 92, (self).f9)
                d_2575_v46_: _dafny.Map
                d_2575_v46_ = _dafny.Map({((d_2572_v43_)[default__.fm13(d_2570_v41_, d_2570_v41_, False, globalState)] if (default__.fm13(d_2570_v41_, d_2570_v41_, False, globalState)) in (d_2572_v43_) else (d_2573_v44_)[default__.safeIndex(942, len(d_2573_v44_))]): (d_2574_v45_).cf57})
                d_2576_v47_: _dafny.Seq
                d_2576_v47_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_2577_v48_: _dafny.Set
                d_2577_v48_ = _dafny.Set({(self).f8})
                d_2578_v49_: _dafny.Map
                d_2578_v49_ = _dafny.Map({d_2576_v47_: d_2577_v48_})
                d_2579_v50_: _dafny.Map
                d_2579_v50_ = _dafny.Map({(self).f9: (self).f8})
                d_2580_v51_: _dafny.Array
                def lambda168_(d_2581_i7_):
                    return (d_2581_i7_) - ((self).f8)

                init96_ = lambda168_
                nw502_ = _dafny.Array(None, 16)
                for i0_96_ in range(nw502_.length(0)):
                    nw502_[i0_96_] = init96_(i0_96_)
                d_2580_v51_ = nw502_
                rhs496_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgbxxirqm"))
                rhs497_ = not ((d_2578_v49_) != (d_2578_v49_)) or ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emovxn"))) < (((d_2575_v46_)[d_2571_v42_] if (d_2571_v42_) in (d_2575_v46_) else d_2569_v40_)))
                rhs498_ = (d_2576_v47_)[default__.safeIndex(len(d_2579_v50_), len(d_2576_v47_))]
                rhs499_ = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2582_i6_ in range(default__.abs(-134))])) == (d_2569_v40_)
                rhs500_ = d_2580_v51_
                d_2569_v40_ = rhs496_
                d_2570_v41_ = rhs497_
                d_2570_v41_ = rhs498_
                d_2570_v41_ = rhs499_
                r0 = rhs500_
                d_2583_v52_: _dafny.Map
                d_2583_v52_ = _dafny.Map({d_2570_v41_: ((self).f8) > ((self).f8)})
                d_2570_v41_ = ((d_2583_v52_)[(self).f9] if ((self).f9) in (d_2583_v52_) else (self).f9)
                d_2584_v53_: str
                d_2584_v53_ = _dafny.CodePoint('u')
                d_2585_v54_: _dafny.Map
                d_2585_v54_ = _dafny.Map({(self).f9: d_2584_v53_})
                d_2586_v55_: _dafny.Map
                d_2586_v55_ = _dafny.Map({(self).f8: 209})
                index506_ = default__.safeIndex(426, (d_2580_v51_).length(0))
                (d_2580_v51_)[index506_] = ((self).f8) * ((self).fm8(d_2585_v54_, d_2586_v55_, default__.fm12(globalState), globalState))
                index507_ = default__.safeIndex(426, (d_2580_v51_).length(0))
                rhs501_ = d_2580_v51_
                rhs502_ = d_2584_v53_
                rhs503_ = (self).f8
                lhs256_ = d_2580_v51_
                lhs257_ = default__.safeIndex(426, (d_2580_v51_).length(0))
                d_2580_v51_ = rhs501_
                d_2584_v53_ = rhs502_
                lhs256_[lhs257_] = rhs503_
                d_2570_v41_ = d_2570_v41_
                index508_ = default__.safeIndex(426, (d_2580_v51_).length(0))
                (d_2580_v51_)[index508_] = default__.safeDivisionInt((self).f8, (0) - ((len(d_2577_v48_)) * (146)))
            d_2587_v56_: _dafny.Map
            d_2587_v56_ = _dafny.Map({(self).f9: default__.fm14(globalState)})
            default__.m0((self).f8, ((self).f8) < (len(d_2587_v56_)), globalState)
            d_2588_v57_: D1
            d_2588_v57_ = D1_DC1(default__.fm15((self).f9, (self).f9, (self).f8, globalState))
            source25_ = (d_2588_v57_ if (self).f9 else d_2588_v57_)
            if source25_.is_DC2:
                d_2589___mcc_h0_ = source25_.cf2
                d_2590___mcc_h1_ = source25_.cf3
                d_2591_cf3_ = d_2590___mcc_h1_
                d_2592_cf2_ = d_2589___mcc_h0_
                d_2593_v58_: _dafny.Array
                def lambda169_(d_2594_cf2_):
                    def lambda170_(d_2595_i8_):
                        return default__.safeModuloInt(d_2595_i8_, d_2594_cf2_)

                    return lambda170_

                init97_ = lambda169_(d_2592_cf2_)
                nw503_ = _dafny.Array(None, 4)
                for i0_97_ in range(nw503_.length(0)):
                    nw503_[i0_97_] = init97_(i0_97_)
                d_2593_v58_ = nw503_
                r0 = d_2593_v58_
                default__.m0(d_2592_cf2_, (self).f9, globalState)
                d_2596_v59_: C8
                nw504_ = C8()
                nw504_.ctor__()
                d_2596_v59_ = nw504_
                d_2597_v60_: _dafny.Map
                d_2597_v60_ = _dafny.Map({not ((self).f9) or ((self).f9): d_2596_v59_})
                d_2598_v61_: _dafny.Seq
                d_2598_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                d_2599_v62_: _dafny.Seq
                d_2599_v62_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                d_2600_v63_: D23
                d_2600_v63_ = D23_DC55(d_2598_v61_, (self).f9, (self).f9, (self).f9, d_2599_v62_)
                d_2596_v59_ = ((d_2597_v60_)[((self).f9 if not((d_2600_v63_).cf88) else (self).f9)] if (((self).f9 if not((d_2600_v63_).cf88) else (self).f9)) in (d_2597_v60_) else d_2596_v59_)
                d_2601_v64_: str
                d_2601_v64_ = _dafny.CodePoint('x')
                d_2601_v64_ = d_2601_v64_
            elif True:
                d_2602___mcc_h2_ = source25_.cf1
                d_2603_cf1_ = d_2602___mcc_h2_
                d_2604_v65_: _dafny.Map
                d_2604_v65_ = _dafny.Map({(self).f8: (self).f8})
                d_2604_v65_ = (d_2604_v65_) | ((d_2604_v65_ if (self).f9 else d_2604_v65_))
                d_2605_v66_: _dafny.Array
                def lambda171_(d_2606_i9_):
                    return not((self).f9)

                init98_ = lambda171_
                nw505_ = _dafny.Array(None, 25)
                for i0_98_ in range(nw505_.length(0)):
                    nw505_[i0_98_] = init98_(i0_98_)
                d_2605_v66_ = nw505_
                d_2607_v67_: _dafny.Seq
                d_2607_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idq"))
                index509_ = default__.safeIndex(630, (d_2605_v66_).length(0))
                (d_2605_v66_)[index509_] = (default__.fm0(74, (self).f8, globalState)) not in (d_2607_v67_)
                d_2608_v68_: _dafny.Set
                d_2608_v68_ = _dafny.Set({d_2607_v67_})
                d_2609_v69_: _dafny.Map
                d_2609_v69_ = _dafny.Map({d_2605_v66_: False})
                d_2610_v70_: _dafny.Array
                nw506_ = _dafny.Array(None, 24)
                nw506_[int(0)] = ((self).f8) * ((self).f8)
                nw506_[int(1)] = (self).f8
                nw506_[int(2)] = (self).f8
                nw506_[int(3)] = default__.safeModuloInt((self).f8, (self).f8)
                nw506_[int(4)] = (self).f8
                nw506_[int(5)] = 467
                nw506_[int(6)] = (self).f8
                nw506_[int(7)] = (0) - ((0) - ((self).f8))
                nw506_[int(8)] = (0) - (len((d_2607_v67_) + (d_2607_v67_)))
                nw506_[int(9)] = (self).f8
                nw506_[int(10)] = default__.safeModuloInt(len(d_2607_v67_), (0) - (len(d_2608_v68_)))
                nw506_[int(11)] = (self).f8
                nw506_[int(12)] = (self).f8
                nw506_[int(13)] = 963
                nw506_[int(14)] = (self).f8
                nw506_[int(15)] = (self).f8
                nw506_[int(16)] = (self).f8
                nw506_[int(17)] = (self).f8
                nw506_[int(18)] = (0) - (len(d_2609_v69_))
                nw506_[int(19)] = (self).f8
                nw506_[int(20)] = (self).f8
                nw506_[int(21)] = (0) - ((self).f8)
                nw506_[int(22)] = ((self).f8) - ((0) - ((self).f8))
                nw506_[int(23)] = (0) - ((self).f8)
                d_2610_v70_ = nw506_
                d_2611_v71_: _dafny.Seq
                d_2611_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wu"))])
                index510_ = default__.safeIndex(430, (d_2610_v70_).length(0))
                (d_2610_v70_)[index510_] = len(d_2611_v71_)
                d_2612_v72_: int
                d_2612_v72_ = 695
                d_2613_v73_: _dafny.Map
                d_2613_v73_ = _dafny.Map({d_2612_v72_: d_2605_v66_})
                index511_ = default__.safeIndex(630, (d_2605_v66_).length(0))
                index512_ = default__.safeIndex(430, (d_2610_v70_).length(0))
                rhs504_ = (self).f9
                rhs505_ = (d_2607_v67_).set(default__.safeIndex((self).f8, len(d_2607_v67_)), (self).f5)
                rhs506_ = len((d_2613_v73_) | (d_2613_v73_))
                rhs507_ = ((self).f8) - (d_2612_v72_)
                lhs258_ = d_2605_v66_
                lhs259_ = default__.safeIndex(630, (d_2605_v66_).length(0))
                lhs260_ = d_2610_v70_
                lhs261_ = default__.safeIndex(430, (d_2610_v70_).length(0))
                lhs258_[lhs259_] = rhs504_
                d_2607_v67_ = rhs505_
                lhs260_[lhs261_] = rhs506_
                d_2612_v72_ = rhs507_
                default__.m0((self).f8, (d_2605_v66_)[default__.safeIndex(630, (d_2605_v66_).length(0))], globalState)
                d_2614_v74_: D17
                d_2614_v74_ = D17_DC42(d_2605_v66_, (self).f8)
                d_2604_v65_ = (d_2604_v65_).set(len(_dafny.Set({d_2614_v74_, d_2614_v74_})), (self).f8)
            d_2615_v75_: bool
            d_2615_v75_ = True
            d_2615_v75_ = not(((self).f9 if (self).f9 else d_2615_v75_))
        elif True:
            d_2616_v76_: int
            d_2616_v76_ = 303
            d_2616_v76_ = (0) - (((0) * (d_2616_v76_)) * (d_2616_v76_))
            d_2617_v77_: _dafny.Array
            nw507_ = _dafny.Array(_dafny.Seq({}), 14)
            d_2617_v77_ = nw507_
            d_2618_v78_: _dafny.Seq
            d_2618_v78_ = _dafny.SeqWithoutIsStrInference([d_2616_v76_, (self).f8, d_2616_v76_])
            index513_ = default__.safeIndex(339, (d_2617_v77_).length(0))
            (d_2617_v77_)[index513_] = d_2618_v78_
            index514_ = default__.safeIndex(339, (d_2617_v77_).length(0))
            (d_2617_v77_)[index514_] = (d_2618_v78_) + (d_2618_v78_)
            d_2619_v79_: _dafny.Seq
            d_2619_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            d_2620_v80_: _dafny.Seq
            d_2620_v80_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_2616_v76_ = default__.safeDivisionInt((self).f8, (len(d_2619_v79_)) * ((_dafny.MultiSet(d_2620_v80_)).cardinality))
            d_2616_v76_ = default__.safeDivisionInt(d_2616_v76_, 861)
            d_2621_v81_: _dafny.Set
            d_2621_v81_ = _dafny.Set({(self).f9, (self).f9})
            d_2622_v82_: _dafny.Map
            d_2622_v82_ = _dafny.Map({(self).f5: d_2621_v81_})
            d_2623_v83_: _dafny.Map
            d_2623_v83_ = _dafny.Map({(self).f8: (self).f9})
            d_2624_v84_: C0
            nw508_ = C0()
            nw508_.ctor__()
            d_2624_v84_ = nw508_
            d_2625_v85_: _dafny.Map
            d_2625_v85_ = _dafny.Map({(self).f9: d_2624_v84_})
            rhs508_ = default__.safeDivisionInt((self).f8, d_2616_v76_)
            rhs509_ = (d_2621_v81_) - ((((d_2622_v82_)[(self).f5] if ((self).f5) in (d_2622_v82_) else d_2621_v81_)) | (_dafny.Set({not(default__.fm1((self).f8, d_2616_v76_, (self).f9, False, globalState)), ((d_2623_v83_)[len(d_2625_v85_)] if (len(d_2625_v85_)) in (d_2623_v83_) else (d_2624_v84_).fm18((self).f9, d_2616_v76_, not((self).f9), (self).f8, globalState))})))
            d_2616_v76_ = rhs508_
            d_2621_v81_ = rhs509_
        d_2626_v86_: _dafny.Seq
        d_2626_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgwfnpj"))
        d_2627_v87_: _dafny.Seq
        d_2627_v87_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehknsls"))])
        d_2628_v88_: _dafny.MultiSet
        d_2628_v88_ = _dafny.MultiSet([d_2626_v86_, d_2626_v86_, (d_2627_v87_)[default__.safeIndex(-467, len(d_2627_v87_))], d_2626_v86_, (d_2627_v87_)[default__.safeIndex(716, len(d_2627_v87_))]])
        if ((self).f9) and ((d_2628_v88_).isdisjoint(d_2628_v88_)):
            d_2629_v89_: _dafny.Array
            nw509_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_2629_v89_ = nw509_
            d_2629_v89_ = d_2629_v89_
            d_2630_v90_: _dafny.Array
            nw510_ = _dafny.Array(int(0), 11)
            d_2630_v90_ = nw510_
            d_2631_v91_: _dafny.Map
            d_2631_v91_ = _dafny.Map({d_2630_v90_: (d_2626_v86_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnuurg")))})
            d_2631_v91_ = d_2631_v91_
            d_2632_v92_: _dafny.Map
            d_2632_v92_ = _dafny.Map({default__.fm12(globalState): ((self).f9) == ((self).f9)})
            d_2632_v92_ = d_2632_v92_
            d_2633_v93_: int
            d_2633_v93_ = 978
            d_2634_v94_: _dafny.MultiSet
            d_2634_v94_ = _dafny.MultiSet([d_2630_v90_])
            d_2635_v95_: _dafny.Array
            def lambda172_(d_2636_i10_):
                return (self).f9

            init99_ = lambda172_
            nw511_ = _dafny.Array(None, 28)
            for i0_99_ in range(nw511_.length(0)):
                nw511_[i0_99_] = init99_(i0_99_)
            d_2635_v95_ = nw511_
            d_2637_v96_: _dafny.Map
            d_2637_v96_ = _dafny.Map({d_2635_v95_: (self).f9})
            d_2638_v97_: _dafny.Map
            d_2638_v97_ = _dafny.Map({d_2633_v93_: (self).f9})
            d_2639_v98_: _dafny.Seq
            d_2639_v98_ = _dafny.SeqWithoutIsStrInference([False])
            d_2640_v99_: _dafny.MultiSet
            d_2640_v99_ = _dafny.MultiSet([(self).f9, (d_2639_v98_)[default__.safeIndex(d_2633_v93_, len(d_2639_v98_))]])
            d_2641_v100_: _dafny.MultiSet
            d_2641_v100_ = _dafny.MultiSet([(d_2637_v96_) == (d_2637_v96_), (self).f9, (self).f9, ((0) - (len(d_2638_v97_))) <= (((d_2640_v99_)[(self).f9] if ((self).f9) in (d_2640_v99_) else d_2633_v93_))])
            rhs510_ = d_2628_v88_
            rhs511_ = ((d_2634_v94_)[d_2630_v90_] if (d_2630_v90_) in (d_2634_v94_) else d_2633_v93_)
            rhs512_ = (d_2640_v99_).cardinality
            d_2628_v88_ = rhs510_
            d_2633_v93_ = rhs511_
            d_2633_v93_ = rhs512_
            d_2633_v93_ = d_2633_v93_
        elif True:
            d_2642_v101_: C3
            nw512_ = C3()
            nw512_.ctor__((self).f4, (self).f5)
            d_2642_v101_ = nw512_
            d_2643_v102_: _dafny.Map
            d_2643_v102_ = _dafny.Map({(self).f9: 683})
            d_2644_v103_: _dafny.Array
            nw513_ = _dafny.Array(False, 19)
            d_2644_v103_ = nw513_
            d_2645_v104_: D17
            d_2645_v104_ = D17_DC42(d_2644_v103_, -284)
            d_2646_v105_: _dafny.Seq
            d_2646_v105_ = _dafny.SeqWithoutIsStrInference([(d_2645_v104_).cf65])
            d_2647_v106_: _dafny.Set
            d_2647_v106_ = _dafny.Set({(d_2642_v101_).fm5(len(d_2643_v102_), (self).f9, (self).f8, len(d_2646_v105_), globalState)})
            d_2647_v106_ = d_2647_v106_
            d_2648_v107_: D10
            d_2648_v107_ = D10_DC21(_dafny.Set({(self).f9}))
            d_2649_v108_: _dafny.Map
            d_2649_v108_ = _dafny.Map({default__.fm21((self).f8, 562, (d_2648_v107_).cf32, globalState): (self).f8})
            d_2650_v109_: _dafny.Seq
            d_2650_v109_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_2649_v108_ = (d_2649_v108_).set(((d_2650_v109_) + ((d_2650_v109_).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chfoomfv")))), len(d_2650_v109_)), (self).f9))).set(default__.safeIndex((self).f8, len((d_2650_v109_) + ((d_2650_v109_).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chfoomfv")))), len(d_2650_v109_)), (self).f9)))), (self).f9), (self).f8)
            d_2651_v110_: _dafny.Set
            d_2651_v110_ = _dafny.Set({_dafny.Set({(self).f9})})
            d_2652_v111_: _dafny.Seq
            d_2652_v111_ = _dafny.SeqWithoutIsStrInference([default__.fm26((self).f8, len(d_2626_v86_), globalState), _dafny.Set({(self).f9})])
            def iife187_():
                coll69_ = _dafny.Set()
                compr_69_: _dafny.Set
                for compr_69_ in (d_2652_v111_).Elements:
                    d_2653_v112_: _dafny.Set = compr_69_
                    if (d_2653_v112_) in (d_2652_v111_):
                        coll69_ = coll69_.union(_dafny.Set([d_2653_v112_]))
                return _dafny.Set(coll69_)
            d_2651_v110_ = iife187_()
            
            d_2654_v113_: _dafny.Map
            d_2654_v113_ = _dafny.Map({(self).f5: ((0) - ((self).f8) if (self).f9 else (self).f8)})
            d_2654_v113_ = (d_2654_v113_).set((self).f5, (self).f8)
        if (self).f9:
            d_2655_v114_: _dafny.Array
            def lambda173_(d_2656_i11_):
                return default__.safeModuloInt(d_2656_i11_, (self).f8)

            init100_ = lambda173_
            nw514_ = _dafny.Array(None, 21)
            for i0_100_ in range(nw514_.length(0)):
                nw514_[i0_100_] = init100_(i0_100_)
            d_2655_v114_ = nw514_
            index515_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            (d_2655_v114_)[index515_] = (self).f8
            d_2657_v115_: _dafny.Map
            d_2657_v115_ = _dafny.Map({True: (self).f8})
            index516_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            (d_2655_v114_)[index516_] = len(_dafny.SeqWithoutIsStrInference([((self).f8) > ((0) - (((d_2657_v115_)[True] if (True) in (d_2657_v115_) else (self).f8))), ((self).f8) > ((self).f8)]))
            d_2626_v86_ = d_2626_v86_
            d_2658_v116_: D16
            d_2658_v116_ = D16_DC38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbqpxnm")), (self).f9, (d_2655_v114_)[default__.safeIndex(424, (d_2655_v114_).length(0))], (self).f9)
            d_2659_v117_: _dafny.Map
            d_2659_v117_ = _dafny.Map({(self).f9: D16_DC39(d_2658_v116_)})
            d_2660_v118_: _dafny.Seq
            d_2660_v118_ = _dafny.SeqWithoutIsStrInference([d_2659_v117_])
            d_2661_v119_: _dafny.Seq
            d_2661_v119_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f8, (self).f8})])
            d_2662_v120_: D26
            d_2662_v120_ = D26_DC61(d_2661_v119_)
            index517_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            (d_2655_v114_)[index517_] = ((self).f8) + ((len((d_2660_v118_)[default__.safeIndex(-739, len(d_2660_v118_))]) if (self).f9 else len((d_2662_v120_).cf99)))
            d_2663_v121_: _dafny.Seq
            d_2663_v121_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_2664_v122_: bool
            d_2664_v122_ = False
            d_2665_v123_: _dafny.Array
            nw515_ = _dafny.Array(None, 3)
            nw515_[int(0)] = (self).f9
            nw515_[int(1)] = (self).f9
            nw515_[int(2)] = ((0) - ((self).f8)) >= ((d_2655_v114_)[default__.safeIndex(424, (d_2655_v114_).length(0))])
            d_2665_v123_ = nw515_
            index518_ = default__.safeIndex(980, (d_2665_v123_).length(0))
            (d_2665_v123_)[index518_] = d_2664_v122_
            d_2666_v124_: _dafny.Set
            d_2666_v124_ = _dafny.Set({(self).f9, (self).f9})
            index519_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            index520_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            index521_ = default__.safeIndex(980, (d_2665_v123_).length(0))
            rhs513_ = default__.fm21((self).f8, (self).f8, d_2666_v124_, globalState)
            rhs514_ = (self).f9
            rhs515_ = (d_2655_v114_)[default__.safeIndex(424, (d_2655_v114_).length(0))]
            rhs516_ = (d_2655_v114_)[default__.safeIndex(424, (d_2655_v114_).length(0))]
            rhs517_ = default__.fm1(default__.safeDivisionInt((self).f8, (self).f8), 344, (d_2626_v86_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2667_i12_ in range(default__.abs(-310))])), (self).f9, globalState)
            lhs262_ = d_2655_v114_
            lhs263_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            lhs264_ = d_2655_v114_
            lhs265_ = default__.safeIndex(424, (d_2655_v114_).length(0))
            lhs266_ = d_2665_v123_
            lhs267_ = default__.safeIndex(980, (d_2665_v123_).length(0))
            d_2663_v121_ = rhs513_
            d_2664_v122_ = rhs514_
            lhs262_[lhs263_] = rhs515_
            lhs264_[lhs265_] = rhs516_
            lhs266_[lhs267_] = rhs517_
            d_2668_v125_: _dafny.Seq
            d_2668_v125_ = _dafny.SeqWithoutIsStrInference([(self).f8, (0) - ((self).f8)])
            d_2669_v126_: _dafny.Map
            d_2669_v126_ = _dafny.Map({d_2668_v125_: (self).f8})
            d_2670_v127_: _dafny.Map
            d_2670_v127_ = d_2669_v126_
            d_2671_v128_: _dafny.Map
            d_2671_v128_ = _dafny.Map({(d_2670_v127_ if not((self).f9) else d_2670_v127_): d_2664_v122_})
            d_2672_v129_: _dafny.Map
            d_2672_v129_ = _dafny.Map({(self).f8: d_2669_v126_})
            d_2673_v130_: _dafny.Map
            d_2673_v130_ = _dafny.Map({(d_2665_v123_)[default__.safeIndex(980, (d_2665_v123_).length(0))]: (self).f5})
            d_2674_v131_: _dafny.Map
            d_2674_v131_ = _dafny.Map({(self).f8: 311})
            d_2671_v128_ = (d_2671_v128_).set(((d_2672_v129_)[(self).fm8(d_2673_v130_, d_2674_v131_, d_2626_v86_, globalState)] if ((self).fm8(d_2673_v130_, d_2674_v131_, d_2626_v86_, globalState)) in (d_2672_v129_) else d_2669_v126_), (d_2665_v123_)[default__.safeIndex(980, (d_2665_v123_).length(0))])
        elif True:
            d_2675_v132_: int
            d_2675_v132_ = -444
            d_2676_v133_: _dafny.Seq
            d_2676_v133_ = _dafny.SeqWithoutIsStrInference([(self).f8])
            d_2677_v135_: C8
            nw516_ = C8()
            nw516_.ctor__()
            d_2677_v135_ = nw516_
            d_2678_v136_: _dafny.Map
            d_2678_v136_ = _dafny.Map({d_2677_v135_: (self).f9})
            d_2679_v137_: _dafny.Map
            d_2679_v137_ = _dafny.Map({(self).f9: (self).f8})
            d_2680_v138_: _dafny.Set
            d_2680_v138_ = _dafny.Set({-145, len(d_2679_v137_)})
            d_2681_v139_: _dafny.Array
            nw517_ = _dafny.Array(None, 7)
            nw517_[int(0)] = ((0) - (default__.fm14(globalState))) < ((d_2676_v133_)[default__.safeIndex(d_2675_v132_, len(d_2676_v133_))])
            nw517_[int(1)] = (self).f9
            nw517_[int(2)] = (self).f9
            def iife188_():
                coll70_ = _dafny.Map()
                compr_70_: str
                for compr_70_ in (d_2626_v86_).Elements:
                    d_2682_v134_: str = compr_70_
                    if (d_2682_v134_) in (d_2626_v86_):
                        coll70_[d_2682_v134_] = d_2675_v132_
                return _dafny.Map(coll70_)
            nw517_[int(3)] = ((self).f8) not in (default__.fm31(len(iife188_()
            ), len(d_2626_v86_), globalState))
            nw517_[int(4)] = ((d_2678_v136_)[d_2677_v135_] if (d_2677_v135_) in (d_2678_v136_) else (self).f9)
            nw517_[int(5)] = (self).f9
            nw517_[int(6)] = (_dafny.Set({(self).f8, len(d_2680_v138_)})).isdisjoint(_dafny.Set({(self).f8}))
            d_2681_v139_ = nw517_
            index522_ = default__.safeIndex(969, (d_2681_v139_).length(0))
            (d_2681_v139_)[index522_] = (self).f9
            d_2683_v140_: _dafny.Array
            nw518_ = _dafny.Array(None, 1)
            nw518_[int(0)] = 351
            d_2683_v140_ = nw518_
            d_2684_v141_: _dafny.MultiSet
            d_2684_v141_ = _dafny.MultiSet([(self).f9])
            d_2685_v142_: D1
            d_2685_v142_ = D1_DC1(d_2684_v141_)
            index523_ = default__.safeIndex(553, (d_2683_v140_).length(0))
            (d_2683_v140_)[index523_] = ((_dafny.MultiSet([(self).f9])).intersection((d_2685_v142_).cf1)).cardinality
            d_2686_v143_: _dafny.Set
            d_2686_v143_ = _dafny.Set({(self).f9})
            d_2687_v144_: _dafny.MultiSet
            d_2687_v144_ = _dafny.MultiSet([(self).f8, (self).f8, -770, d_2675_v132_, len(d_2686_v143_)])
            index524_ = default__.safeIndex(969, (d_2681_v139_).length(0))
            index525_ = default__.safeIndex(553, (d_2683_v140_).length(0))
            rhs518_ = 35
            rhs519_ = (self).f8
            rhs520_ = (d_2687_v144_).ispropersubset((d_2687_v144_).intersection(d_2687_v144_))
            rhs521_ = len(d_2676_v133_)
            rhs522_ = default__.safeDivisionInt((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxesrvvq")))) + (d_2675_v132_), d_2675_v132_)
            lhs268_ = d_2681_v139_
            lhs269_ = default__.safeIndex(969, (d_2681_v139_).length(0))
            lhs270_ = d_2683_v140_
            lhs271_ = default__.safeIndex(553, (d_2683_v140_).length(0))
            d_2675_v132_ = rhs518_
            d_2675_v132_ = rhs519_
            lhs268_[lhs269_] = rhs520_
            lhs270_[lhs271_] = rhs521_
            d_2675_v132_ = rhs522_
            d_2626_v86_ = d_2626_v86_
            index526_ = default__.safeIndex(553, (d_2683_v140_).length(0))
            (d_2683_v140_)[index526_] = default__.safeDivisionInt(d_2675_v132_, (self).f8)
            d_2688_v145_: _dafny.Map
            d_2688_v145_ = _dafny.Map({(d_2681_v139_)[default__.safeIndex(969, (d_2681_v139_).length(0))]: (self).f5})
            d_2689_v146_: D1
            d_2689_v146_ = D1_DC2(len(default__.fm21(d_2675_v132_, d_2675_v132_, d_2686_v143_, globalState)), d_2675_v132_)
            d_2690_v147_: _dafny.Map
            d_2690_v147_ = _dafny.Map({d_2684_v141_: (d_2683_v140_)[default__.safeIndex(553, (d_2683_v140_).length(0))]})
            index527_ = default__.safeIndex(969, (d_2681_v139_).length(0))
            (d_2681_v139_)[index527_] = ((self).fm8(d_2688_v145_, (self).fm9(d_2689_v146_, d_2626_v86_, globalState), d_2626_v86_, globalState)) == (len(d_2690_v147_))
            d_2691_v148_: D6
            d_2691_v148_ = D6_DC15(True, True, False)
            d_2692_v149_: _dafny.Map
            d_2692_v149_ = _dafny.Map({(d_2691_v148_).cf20: not((self).f9)})
            d_2693_v150_: _dafny.Map
            d_2693_v150_ = d_2692_v149_
            d_2694_v151_: _dafny.Map
            d_2694_v151_ = _dafny.Map({d_2626_v86_: d_2692_v149_})
            d_2695_v152_: _dafny.Seq
            d_2695_v152_ = _dafny.SeqWithoutIsStrInference([d_2683_v140_, d_2683_v140_])
            d_2696_v153_: _dafny.Map
            d_2696_v153_ = _dafny.Map({d_2695_v152_: d_2694_v151_})
            index528_ = default__.safeIndex(969, (d_2681_v139_).length(0))
            (d_2681_v139_)[index528_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skiyru"))) + (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2697_i13_ in range(default__.abs(-117))]))) in ((d_2694_v151_) | (((d_2696_v153_)[d_2695_v152_] if (d_2695_v152_) in (d_2696_v153_) else (d_2694_v151_).set(d_2626_v86_, d_2693_v150_))))
        hi20_ = -926
        for d_2698_i14_ in range(len((d_2626_v86_) + (_dafny.SeqWithoutIsStrInference([(self).f5, (self).f5]))), hi20_):
            d_2699_v154_: _dafny.Array
            def lambda174_(d_2700_i14_):
                def lambda175_(d_2701_i15_):
                    return (d_2701_i15_) * (d_2700_i14_)

                return lambda175_

            init101_ = lambda174_(d_2698_i14_)
            nw519_ = _dafny.Array(None, 2)
            for i0_101_ in range(nw519_.length(0)):
                nw519_[i0_101_] = init101_(i0_101_)
            d_2699_v154_ = nw519_
            index529_ = default__.safeIndex(237, (d_2699_v154_).length(0))
            (d_2699_v154_)[index529_] = (self).f8
            d_2702_v155_: _dafny.Array
            nw520_ = _dafny.Array(None, 2)
            nw520_[int(0)] = d_2699_v154_
            nw520_[int(1)] = d_2699_v154_
            d_2702_v155_ = nw520_
            d_2703_v156_: _dafny.Seq
            d_2703_v156_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_2704_v157_: _dafny.Seq
            d_2704_v157_ = _dafny.SeqWithoutIsStrInference([(self).f9, True, (self).f9, (d_2703_v156_) != (d_2703_v156_)])
            d_2705_v158_: bool
            d_2705_v158_ = False
            d_2706_v159_: _dafny.Map
            d_2706_v159_ = _dafny.Map({d_2705_v158_: (self).f8})
            d_2707_v160_: _dafny.Seq
            d_2707_v160_ = _dafny.SeqWithoutIsStrInference([len(d_2706_v159_), (d_2698_i14_) - (default__.fm14(globalState))])
            d_2708_v161_: _dafny.Map
            d_2708_v161_ = _dafny.Map({(self).f5: 335})
            index530_ = default__.safeIndex(237, (d_2699_v154_).length(0))
            rhs523_ = d_2707_v160_
            rhs524_ = default__.fm13(not(d_2705_v158_), d_2705_v158_, ((self).f8) <= (len(d_2708_v161_)), globalState)
            rhs525_ = d_2702_v155_
            rhs526_ = _dafny.SeqWithoutIsStrInference([d_2705_v158_])
            rhs527_ = default__.fm1((0) - ((self).f8), d_2698_i14_, (self).f9, (d_2705_v158_ if d_2705_v158_ else True), globalState)
            lhs272_ = globalState
            lhs273_ = d_2699_v154_
            lhs274_ = default__.safeIndex(237, (d_2699_v154_).length(0))
            lhs272_.f2 = rhs523_
            lhs273_[lhs274_] = rhs524_
            d_2702_v155_ = rhs525_
            d_2704_v157_ = rhs526_
            d_2705_v158_ = rhs527_
            if (self).f9:
                d_2709_v162_: _dafny.MultiSet
                d_2709_v162_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2698_i14_])])
                index531_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                index532_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                index533_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                rhs528_ = (self).fm4(_dafny.SeqWithoutIsStrInference([(self).f8 for d_2710_i16_ in range(default__.abs(270))]), globalState)
                rhs529_ = d_2698_i14_
                rhs530_ = d_2699_v154_
                rhs531_ = (d_2709_v162_).cardinality
                rhs532_ = d_2626_v86_
                lhs275_ = d_2699_v154_
                lhs276_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                lhs277_ = d_2699_v154_
                lhs278_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                lhs279_ = d_2699_v154_
                lhs280_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                lhs275_[lhs276_] = rhs528_
                lhs277_[lhs278_] = rhs529_
                d_2699_v154_ = rhs530_
                lhs279_[lhs280_] = rhs531_
                d_2626_v86_ = rhs532_
                d_2711_v163_: _dafny.Array
                nw521_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                d_2711_v163_ = nw521_
                index534_ = default__.safeIndex(736, (d_2711_v163_).length(0))
                (d_2711_v163_)[index534_] = _dafny.CodePoint('l')
                index535_ = default__.safeIndex(736, (d_2711_v163_).length(0))
                (d_2711_v163_)[index535_] = (self).f5
                d_2712_v164_: D23
                d_2712_v164_ = D23_DC55(d_2626_v86_, (self).f9, (self).f9, (self).f9, d_2704_v157_)
                d_2713_v165_: _dafny.Map
                d_2713_v165_ = _dafny.Map({(self).f8: len((d_2712_v164_).cf85)})
                d_2714_v166_: _dafny.Seq
                d_2714_v166_ = _dafny.SeqWithoutIsStrInference([d_2713_v165_])
                d_2715_v167_: _dafny.Set
                d_2715_v167_ = _dafny.Set({465})
                d_2716_v168_: D1
                d_2716_v168_ = D1_DC2((0) - (d_2698_i14_), len(d_2715_v167_))
                d_2717_v169_: _dafny.Map
                d_2717_v169_ = _dafny.Map({d_2705_v158_: (d_2711_v163_)[default__.safeIndex(736, (d_2711_v163_).length(0))]})
                d_2714_v166_ = (((_dafny.SeqWithoutIsStrInference([_dafny.Map({d_2698_i14_: len(d_2626_v86_)}), d_2713_v165_, d_2713_v165_])) + (d_2714_v166_)) + (d_2714_v166_)).set(default__.safeIndex((self).f8, len(((_dafny.SeqWithoutIsStrInference([_dafny.Map({d_2698_i14_: len(d_2626_v86_)}), d_2713_v165_, d_2713_v165_])) + (d_2714_v166_)) + (d_2714_v166_))), ((self).fm9(d_2716_v168_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpjskhwm")), globalState)).set(d_2698_i14_, (self).fm8(d_2717_v169_, d_2713_v165_, d_2626_v86_, globalState)))
                d_2718_v170_: _dafny.Map
                d_2718_v170_ = _dafny.Map({d_2698_i14_: d_2699_v154_})
                d_2719_v171_: _dafny.Map
                d_2719_v171_ = _dafny.Map({d_2705_v158_: d_2718_v170_})
                d_2720_v172_: _dafny.Array
                nw522_ = _dafny.Array(None, 17)
                nw522_[int(0)] = d_2718_v170_
                nw522_[int(1)] = d_2718_v170_
                nw522_[int(2)] = (d_2718_v170_) | (d_2718_v170_)
                nw522_[int(3)] = _dafny.Map({(d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]: d_2699_v154_})
                nw522_[int(4)] = d_2718_v170_
                nw522_[int(5)] = (d_2718_v170_) | (d_2718_v170_)
                nw522_[int(6)] = (_dafny.Map({(d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]: d_2699_v154_})) | (_dafny.Map({len(_dafny.Map({(self).f9: (self).f9})): d_2699_v154_}))
                nw522_[int(7)] = (d_2718_v170_) | (d_2718_v170_)
                nw522_[int(8)] = d_2718_v170_
                nw522_[int(9)] = d_2718_v170_
                nw522_[int(10)] = ((d_2719_v171_)[d_2705_v158_] if (d_2705_v158_) in (d_2719_v171_) else d_2718_v170_)
                nw522_[int(11)] = d_2718_v170_
                nw522_[int(12)] = d_2718_v170_
                nw522_[int(13)] = (d_2718_v170_ if (self).f9 else d_2718_v170_)
                nw522_[int(14)] = (_dafny.Map({(d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]: d_2699_v154_})).set((self).f8, d_2699_v154_)
                nw522_[int(15)] = (d_2718_v170_) | (d_2718_v170_)
                nw522_[int(16)] = d_2718_v170_
                d_2720_v172_ = nw522_
                index536_ = default__.safeIndex(255, (d_2720_v172_).length(0))
                (d_2720_v172_)[index536_] = d_2718_v170_
                d_2721_v173_: _dafny.Seq
                d_2721_v173_ = _dafny.SeqWithoutIsStrInference([d_2718_v170_])
                index537_ = default__.safeIndex(255, (d_2720_v172_).length(0))
                (d_2720_v172_)[index537_] = (((d_2721_v173_)[default__.safeIndex(d_2698_i14_, len(d_2721_v173_))]) | (d_2718_v170_)) | (d_2718_v170_)
                d_2722_v174_: _dafny.Map
                d_2722_v174_ = _dafny.Map({(self).f9: not((self).f9)})
                d_2722_v174_ = ((d_2722_v174_) | (d_2722_v174_)) | (d_2722_v174_)
            elif True:
                d_2705_v158_ = ((d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]) != (d_2698_i14_)
                d_2723_v175_: C4
                nw523_ = C4()
                nw523_.ctor__(((self).f8) * ((d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]), ((self).f8) - ((d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]), (self).f4, (self).f5)
                d_2723_v175_ = nw523_
                d_2724_v176_: _dafny.Array
                nw524_ = _dafny.Array(_dafny.Map({}), 4)
                d_2724_v176_ = nw524_
                index538_ = default__.safeIndex(540, (d_2724_v176_).length(0))
                def iife189_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(170, -457):
                        d_2725_v177_: int = compr_71_
                        if ((170) <= (d_2725_v177_)) and ((d_2725_v177_) < (-457)):
                            coll71_[default__.safeModuloInt(d_2725_v177_, (_dafny.MultiSet([False])).cardinality)] = d_2705_v158_
                    return _dafny.Map(coll71_)
                (d_2724_v176_)[index538_] = iife189_()
                
                d_2726_v178_: _dafny.Map
                d_2726_v178_ = _dafny.Map({(d_2698_i14_) * ((self).f8): d_2705_v158_})
                index539_ = default__.safeIndex(540, (d_2724_v176_).length(0))
                (d_2724_v176_)[index539_] = d_2726_v178_
                d_2727_v179_: _dafny.Map
                d_2727_v179_ = _dafny.Map({False: (self).f5})
                d_2728_v180_: D5
                d_2728_v180_ = D5_DC11((d_2727_v179_).set(d_2705_v158_, _dafny.CodePoint('y')))
                d_2729_v181_: _dafny.Map
                d_2729_v181_ = _dafny.Map({len(d_2626_v86_): (d_2723_v175_).f12})
                rhs533_ = d_2728_v180_
                rhs534_ = (self).f9
                rhs535_ = (self).fm8(default__.fm48(d_2705_v158_, False, default__.fm1(d_2723_v175_.f11, d_2698_i14_, (self).f9, d_2705_v158_, globalState), (self).f8, globalState), d_2729_v181_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwjmmom"))) + (d_2626_v86_), globalState)
                lhs281_ = d_2723_v175_
                d_2728_v180_ = rhs533_
                d_2705_v158_ = rhs534_
                lhs281_.f11 = rhs535_
                index540_ = default__.safeIndex(237, (d_2699_v154_).length(0))
                (d_2699_v154_)[index540_] = (d_2723_v175_).f12
            index541_ = default__.safeIndex(237, (d_2699_v154_).length(0))
            (d_2699_v154_)[index541_] = (d_2699_v154_)[default__.safeIndex(237, (d_2699_v154_).length(0))]
            d_2730_v182_: _dafny.Array
            def lambda176_(d_2731_i17_):
                return (self).f5

            init102_ = lambda176_
            nw525_ = _dafny.Array(None, 13)
            for i0_102_ in range(nw525_.length(0)):
                nw525_[i0_102_] = init102_(i0_102_)
            d_2730_v182_ = nw525_
            d_2732_v183_: _dafny.Seq
            d_2732_v183_ = _dafny.SeqWithoutIsStrInference([d_2730_v182_])
            rhs536_ = d_2705_v158_
            rhs537_ = (self).f9
            rhs538_ = d_2732_v183_
            d_2705_v158_ = rhs536_
            d_2705_v158_ = rhs537_
            d_2732_v183_ = rhs538_
        d_2733_v184_: C2
        nw526_ = C2()
        nw526_.ctor__((self).f4, (d_2626_v86_)[default__.safeIndex((self).f8, len(d_2626_v86_))])
        d_2733_v184_ = nw526_
        if False:
            d_2734_v185_: int
            d_2734_v185_ = 92
            d_2735_v186_: C5
            nw527_ = C5()
            nw527_.ctor__()
            d_2735_v186_ = nw527_
            d_2736_v187_: D21
            d_2736_v187_ = D21_DC51((self).f9, d_2735_v186_)
            d_2737_v188_: _dafny.Map
            d_2737_v188_ = _dafny.Map({(self).f9: (self).f9})
            d_2738_v189_: _dafny.Seq
            d_2738_v189_ = _dafny.SeqWithoutIsStrInference([((d_2737_v188_)[(self).f9] if ((self).f9) in (d_2737_v188_) else (self).f9)])
            d_2739_v190_: D23
            d_2739_v190_ = D23_DC55(d_2626_v86_, (self).f9, ((d_2736_v187_).cf80 if (self).f9 else (self).f9), (d_2733_v184_).fm5((self).f8, (self).f9, 262, len(d_2626_v86_), globalState), d_2738_v189_)
            d_2740_v191_: _dafny.Set
            d_2740_v191_ = _dafny.Set({d_2737_v188_})
            d_2741_v192_: _dafny.Seq
            d_2741_v192_ = _dafny.SeqWithoutIsStrInference([(self).f8, d_2734_v185_])
            rhs539_ = default__.safeDivisionInt((self).f8, len((d_2740_v191_ if (self).f9 else d_2740_v191_)))
            rhs540_ = d_2739_v190_
            rhs541_ = ((d_2741_v192_)[default__.safeIndex(d_2734_v185_, len(d_2741_v192_))]) + ((self).f8)
            rhs542_ = (0) - (d_2734_v185_)
            d_2734_v185_ = rhs539_
            d_2739_v190_ = rhs540_
            d_2734_v185_ = rhs541_
            d_2734_v185_ = rhs542_
            d_2742_v193_: _dafny.Array
            nw528_ = _dafny.Array(False, 28)
            d_2742_v193_ = nw528_
            index542_ = default__.safeIndex(216, (d_2742_v193_).length(0))
            (d_2742_v193_)[index542_] = (self).f9
            d_2743_v194_: bool
            d_2743_v194_ = False
            index543_ = default__.safeIndex(705, (d_2742_v193_).length(0))
            (d_2742_v193_)[index543_] = (d_2738_v189_) <= (d_2738_v189_)
            d_2744_v195_: _dafny.MultiSet
            d_2744_v195_ = _dafny.MultiSet([d_2734_v185_])
            d_2745_v196_: _dafny.Map
            d_2745_v196_ = _dafny.Map({True: (self).f5})
            d_2746_v197_: _dafny.Map
            d_2746_v197_ = _dafny.Map({d_2738_v189_: d_2743_v194_})
            d_2747_v198_: _dafny.Map
            d_2747_v198_ = _dafny.Map({d_2734_v185_: len(d_2746_v197_)})
            d_2748_v199_: _dafny.Set
            d_2748_v199_ = _dafny.Set({d_2743_v194_, (self).f9})
            d_2749_v200_: _dafny.Array
            nw529_ = _dafny.Array(None, 28)
            nw529_[int(0)] = d_2734_v185_
            nw529_[int(1)] = (self).f8
            nw529_[int(2)] = 601
            nw529_[int(3)] = (0) - (len(d_2626_v86_))
            nw529_[int(4)] = d_2734_v185_
            nw529_[int(5)] = ((d_2744_v195_)[len(_dafny.Map({d_2734_v185_: (self).f9}))] if (len(_dafny.Map({d_2734_v185_: (self).f9}))) in (d_2744_v195_) else d_2734_v185_)
            nw529_[int(6)] = (0) - ((self).f8)
            nw529_[int(7)] = (self).f8
            nw529_[int(8)] = (self).fm8(d_2745_v196_, d_2747_v198_, d_2626_v86_, globalState)
            nw529_[int(9)] = d_2734_v185_
            nw529_[int(10)] = 795
            nw529_[int(11)] = (self).f8
            nw529_[int(12)] = d_2734_v185_
            nw529_[int(13)] = len(d_2748_v199_)
            nw529_[int(14)] = d_2734_v185_
            nw529_[int(15)] = (self).f8
            nw529_[int(16)] = len(d_2741_v192_)
            nw529_[int(17)] = (self).f8
            nw529_[int(18)] = len(_dafny.Map({d_2743_v194_: (d_2741_v192_)[default__.safeIndex(len(d_2737_v188_), len(d_2741_v192_))]}))
            nw529_[int(19)] = (self).f8
            nw529_[int(20)] = 69
            nw529_[int(21)] = -586
            nw529_[int(22)] = 941
            nw529_[int(23)] = (self).f8
            nw529_[int(24)] = (self).f8
            nw529_[int(25)] = (self).f8
            nw529_[int(26)] = (0) - ((self).f8)
            nw529_[int(27)] = len(d_2741_v192_)
            d_2749_v200_ = nw529_
            d_2750_v201_: D11
            d_2750_v201_ = D11_DC25((self).f9, d_2734_v185_, d_2749_v200_, d_2734_v185_)
            d_2751_v202_: _dafny.Map
            d_2751_v202_ = _dafny.Map({d_2738_v189_: (self).f8})
            index544_ = default__.safeIndex(216, (d_2742_v193_).length(0))
            index545_ = default__.safeIndex(705, (d_2742_v193_).length(0))
            rhs543_ = len((d_2626_v86_) + (d_2626_v86_))
            rhs544_ = not ((self).f9) or (d_2743_v194_)
            rhs545_ = (d_2750_v201_).cf40
            rhs546_ = not (not((((d_2751_v202_)[_dafny.SeqWithoutIsStrInference([d_2743_v194_, d_2743_v194_])] if (_dafny.SeqWithoutIsStrInference([d_2743_v194_, d_2743_v194_])) in (d_2751_v202_) else -173)) >= (d_2734_v185_))) or ((d_2734_v185_) >= (d_2734_v185_))
            rhs547_ = not((d_2739_v190_).cf87)
            lhs282_ = d_2742_v193_
            lhs283_ = default__.safeIndex(216, (d_2742_v193_).length(0))
            lhs284_ = d_2742_v193_
            lhs285_ = default__.safeIndex(705, (d_2742_v193_).length(0))
            d_2734_v185_ = rhs543_
            lhs282_[lhs283_] = rhs544_
            r0 = rhs545_
            d_2743_v194_ = rhs546_
            lhs284_[lhs285_] = rhs547_
            d_2752_v203_: D15
            d_2752_v203_ = D15_DC36(D15_DC35(d_2734_v185_))
            d_2753_v204_: _dafny.Map
            d_2753_v204_ = _dafny.Map({((self).f9) or ((d_2742_v193_)[default__.safeIndex(216, (d_2742_v193_).length(0))]): d_2752_v203_})
            d_2753_v204_ = (d_2753_v204_).set(d_2743_v194_, d_2752_v203_)
            d_2754_v205_: _dafny.MultiSet
            d_2754_v205_ = _dafny.MultiSet([(self).f5])
            d_2755_v206_: _dafny.Map
            d_2755_v206_ = _dafny.Map({d_2747_v198_: (self).f9})
            index546_ = default__.safeIndex(216, (d_2742_v193_).length(0))
            rhs548_ = (d_2742_v193_)[default__.safeIndex(216, (d_2742_v193_).length(0))]
            rhs549_ = (0) - ((self).f8)
            rhs550_ = default__.fm13((d_2754_v205_).issubset(d_2754_v205_), ((d_2755_v206_)[d_2747_v198_] if (d_2747_v198_) in (d_2755_v206_) else (d_2742_v193_)[default__.safeIndex(216, (d_2742_v193_).length(0))]), (d_2742_v193_)[default__.safeIndex(216, (d_2742_v193_).length(0))], globalState)
            lhs286_ = d_2742_v193_
            lhs287_ = default__.safeIndex(216, (d_2742_v193_).length(0))
            lhs286_[lhs287_] = rhs548_
            d_2734_v185_ = rhs549_
            d_2734_v185_ = rhs550_
            d_2756_v207_: _dafny.Seq
            d_2756_v207_ = _dafny.SeqWithoutIsStrInference([d_2749_v200_])
            d_2757_v208_: _dafny.Array
            nw530_ = _dafny.Array(None, 8)
            nw530_[int(0)] = (self).f5
            nw530_[int(1)] = (self).f5
            nw530_[int(2)] = (self).f5
            nw530_[int(3)] = (self).f5
            nw530_[int(4)] = (self).f5
            nw530_[int(5)] = _dafny.CodePoint('w')
            nw530_[int(6)] = (self).f5
            nw530_[int(7)] = (self).f5
            d_2757_v208_ = nw530_
            d_2758_v209_: _dafny.Array
            nw531_ = _dafny.Array(_dafny.Map({}), 8)
            d_2758_v209_ = nw531_
            d_2759_v210_: _dafny.Map
            d_2759_v210_ = _dafny.Map({d_2744_v195_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvi")))})
            index547_ = default__.safeIndex(892, (d_2758_v209_).length(0))
            (d_2758_v209_)[index547_] = (d_2759_v210_) | (d_2759_v210_)
            d_2760_v211_: _dafny.Set
            d_2760_v211_ = _dafny.Set({d_2749_v200_, d_2749_v200_, d_2749_v200_})
            index548_ = default__.safeIndex(892, (d_2758_v209_).length(0))
            rhs551_ = d_2756_v207_
            rhs552_ = d_2757_v208_
            rhs553_ = (d_2760_v211_).ispropersubset(_dafny.Set({d_2749_v200_, d_2749_v200_}))
            rhs554_ = d_2759_v210_
            lhs288_ = d_2758_v209_
            lhs289_ = default__.safeIndex(892, (d_2758_v209_).length(0))
            d_2756_v207_ = rhs551_
            d_2757_v208_ = rhs552_
            d_2743_v194_ = rhs553_
            lhs288_[lhs289_] = rhs554_
        elif True:
            d_2761_v212_: _dafny.Array
            def lambda177_(d_2762_i18_):
                return (d_2762_i18_) - ((self).f8)

            init103_ = lambda177_
            nw532_ = _dafny.Array(None, 13)
            for i0_103_ in range(nw532_.length(0)):
                nw532_[i0_103_] = init103_(i0_103_)
            d_2761_v212_ = nw532_
            index549_ = default__.safeIndex(395, (d_2761_v212_).length(0))
            (d_2761_v212_)[index549_] = (self).f8
            index550_ = default__.safeIndex(395, (d_2761_v212_).length(0))
            (d_2761_v212_)[index550_] = (self).f8
            d_2763_v213_: _dafny.Map
            d_2763_v213_ = _dafny.Map({(self).f5: d_2626_v86_})
            d_2764_v214_: _dafny.Array
            nw533_ = _dafny.Array(None, 6)
            nw533_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxlj"))
            nw533_[int(1)] = (d_2626_v86_) + (d_2626_v86_)
            nw533_[int(2)] = d_2626_v86_
            nw533_[int(3)] = (d_2626_v86_) + (d_2626_v86_)
            nw533_[int(4)] = d_2626_v86_
            nw533_[int(5)] = ((d_2763_v213_)[(self).f5] if ((self).f5) in (d_2763_v213_) else d_2626_v86_)
            d_2764_v214_ = nw533_
            index551_ = default__.safeIndex(803, (d_2764_v214_).length(0))
            (d_2764_v214_)[index551_] = d_2626_v86_
            index552_ = default__.safeIndex(803, (d_2764_v214_).length(0))
            (d_2764_v214_)[index552_] = (d_2626_v86_) + (d_2626_v86_)
            d_2765_v215_: C7
            nw534_ = C7()
            nw534_.ctor__()
            d_2765_v215_ = nw534_
            d_2766_v216_: _dafny.Map
            d_2766_v216_ = _dafny.Map({(d_2761_v212_)[default__.safeIndex(395, (d_2761_v212_).length(0))]: (self).f9})
            d_2767_v217_: _dafny.Seq
            d_2767_v217_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])
            d_2766_v216_ = (d_2766_v216_).set((self).f8, (d_2767_v217_) != (d_2767_v217_))
            d_2768_v218_: bool
            d_2768_v218_ = True
            d_2768_v218_ = not((d_2626_v86_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hufpuyy"))))
        d_2769_v219_: _dafny.Array
        nw535_ = _dafny.Array(int(0), 16)
        d_2769_v219_ = nw535_
        r0 = d_2769_v219_
        return r0

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9

class C11(T1, T0):
    def  __init__(self):
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: str = _dafny.CodePoint('D')
        self._f7: int = int(0)
        self._f6: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f6, f7, f4, f5):
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f4 = f4
        (self)._f5 = f5

    def fm4(self, p0, globalState):
        return ((self).f7) * ((self).f7)

    def fm5(self, p0, p1, p2, p3, globalState):
        return not((True if False else ((self).f7) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfwsjafps"))))))

    def fm3(self, p0, p1, globalState):
        if False:
            def iife190_():
                coll72_ = _dafny.Map()
                compr_72_: int
                for compr_72_ in (_dafny.Map({(self).f7: (self).f5})).keys.Elements:
                    d_2770_v0_: int = compr_72_
                    if (d_2770_v0_) in (_dafny.Map({(self).f7: (self).f5})):
                        coll72_[(d_2770_v0_) - (-860)] = D2_DC4(True)
                return _dafny.Map(coll72_)
            return (_dafny.Map({iife190_()
            : True})) != (_dafny.Map({_dafny.Map({(self).f7: D2_DC4(False)}): True}))
        elif True:
            return True

    def fm6(self, p0, p1, globalState):
        return len(((_dafny.Map({(_dafny.MultiSet([(self).f7])).cardinality: _dafny.SeqWithoutIsStrInference([not((False))])})) | (_dafny.Map({(0) - ((self).f7): _dafny.SeqWithoutIsStrInference([True, False])}))) | ((_dafny.Map({(self).f7: _dafny.SeqWithoutIsStrInference([True, False])})) | (_dafny.Map({(self).f7: _dafny.SeqWithoutIsStrInference([False, False, False])}))))

    def fm7(self, p0, p1, p2, p3, globalState):
        return (self).f7

    def m1(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        if False:
            d_2771_v0_: _dafny.Seq
            d_2771_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnsxyet"))
            d_2772_v1_: _dafny.Array
            nw536_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_2772_v1_ = nw536_
            d_2773_v2_: _dafny.Map
            d_2773_v2_ = _dafny.Map({d_2772_v1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gspoppq"))})
            d_2771_v0_ = (((d_2773_v2_)[d_2772_v1_] if (d_2772_v1_) in (d_2773_v2_) else d_2771_v0_)).set(default__.safeIndex(len(d_2771_v0_), len(((d_2773_v2_)[d_2772_v1_] if (d_2772_v1_) in (d_2773_v2_) else d_2771_v0_))), (self).f5)
            d_2774_v3_: _dafny.Array
            nw537_ = _dafny.Array(_dafny.Seq({}), 23)
            d_2774_v3_ = nw537_
            d_2775_v4_: D4
            d_2775_v4_ = D4_DC9(d_2774_v3_)
            d_2774_v3_ = (d_2775_v4_).cf10
            d_2776_v5_: _dafny.MultiSet
            d_2776_v5_ = _dafny.MultiSet([(0) - ((self).f7)])
            d_2777_v6_: D1
            d_2777_v6_ = D1_DC2((self).f7, ((d_2776_v5_)[(self).f7] if ((self).f7) in (d_2776_v5_) else (self).f7))
            d_2778_v7_: bool
            d_2778_v7_ = True
            d_2779_v8_: _dafny.Map
            d_2779_v8_ = _dafny.Map({d_2775_v4_: (self).f7})
            def iife191_(_pat_let59_0):
                def iife192_(d_2780_dt__update__tmp_h0_):
                    def iife193_(_pat_let60_0):
                        def iife194_(d_2781_dt__update_hcf2_h0_):
                            return D1_DC2(d_2781_dt__update_hcf2_h0_, (d_2780_dt__update__tmp_h0_).cf3)
                        return iife194_(_pat_let60_0)
                    return iife193_((self).f7)
                return iife192_(_pat_let59_0)
            r1 = (self).fm7(iife191_(d_2777_v6_), d_2778_v7_, (779) + ((self).f7), not(default__.fm1((self).f7, len((d_2779_v8_).set(d_2775_v4_, (self).f7)), d_2778_v7_, d_2778_v7_, globalState)), globalState)
            r1 = (self).f7
            if (d_2778_v7_) or (False):
                d_2778_v7_ = d_2778_v7_
                d_2782_v9_: _dafny.Array
                def lambda178_(d_2783_v0_):
                    def lambda179_(d_2784_i0_):
                        return d_2783_v0_

                    return lambda179_

                init104_ = lambda178_(d_2771_v0_)
                nw538_ = _dafny.Array(None, 14)
                for i0_104_ in range(nw538_.length(0)):
                    nw538_[i0_104_] = init104_(i0_104_)
                d_2782_v9_ = nw538_
                d_2782_v9_ = (d_2782_v9_ if True else d_2782_v9_)
                d_2785_v10_: T1
                nw539_ = C1()
                nw539_.ctor__((self).f4, (self).f5)
                d_2785_v10_ = nw539_
                d_2786_v11_: _dafny.Seq
                d_2786_v11_ = _dafny.SeqWithoutIsStrInference([d_2785_v10_, d_2785_v10_])
                d_2787_v12_: _dafny.Seq
                d_2787_v12_ = _dafny.SeqWithoutIsStrInference([d_2778_v7_, (self).fm3((0) - (len(d_2786_v11_)), len(_dafny.SeqWithoutIsStrInference([(d_2771_v0_)[default__.safeIndex((self).f7, len(d_2771_v0_))] for d_2788_i1_ in range(default__.abs(161))])), globalState)])
                d_2787_v12_ = ((d_2787_v12_) + (_dafny.SeqWithoutIsStrInference([False, d_2778_v7_]))) + (d_2787_v12_)
                d_2778_v7_ = (d_2787_v12_)[default__.safeIndex(len(d_2787_v12_), len(d_2787_v12_))]
                d_2789_v13_: _dafny.Map
                d_2789_v13_ = _dafny.Map({d_2778_v7_: default__.safeDivisionInt(964, 292)})
                d_2790_v14_: _dafny.Set
                d_2790_v14_ = _dafny.Set({(self).f7})
                d_2791_v15_: _dafny.Map
                d_2791_v15_ = _dafny.Map({d_2771_v0_: d_2778_v7_})
                d_2792_v16_: _dafny.Set
                d_2792_v16_ = _dafny.Set({d_2778_v7_})
                d_2793_v17_: _dafny.Seq
                d_2793_v17_ = _dafny.SeqWithoutIsStrInference([(self).f7, len(default__.fm21(len(d_2791_v15_), (self).f7, d_2792_v16_, globalState))])
                d_2789_v13_ = (d_2789_v13_).set((d_2778_v7_) and ((self).fm3(len(d_2790_v14_), (self).f7, globalState)), (self).fm4(d_2793_v17_, globalState))
            elif True:
                index553_ = default__.safeIndex(309, (d_2772_v1_).length(0))
                (d_2772_v1_)[index553_] = (self).f5
                d_2794_v18_: C9
                nw540_ = C9()
                nw540_.ctor__((self).f4, (self).f5)
                d_2794_v18_ = nw540_
                index554_ = default__.safeIndex(309, (d_2772_v1_).length(0))
                rhs555_ = (((self).f5 if d_2778_v7_ else (self).f5) if d_2778_v7_ else (self).f5)
                rhs556_ = d_2794_v18_
                lhs290_ = d_2772_v1_
                lhs291_ = default__.safeIndex(309, (d_2772_v1_).length(0))
                lhs290_[lhs291_] = rhs555_
                d_2794_v18_ = rhs556_
                d_2795_v19_: _dafny.Seq
                d_2795_v19_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_2796_v20_: _dafny.Array
                nw541_ = _dafny.Array(None, 6)
                nw541_[int(0)] = (self).f7
                nw541_[int(1)] = (self).f7
                nw541_[int(2)] = (self).f7
                nw541_[int(3)] = (self).f7
                nw541_[int(4)] = (self).f7
                nw541_[int(5)] = len(d_2795_v19_)
                d_2796_v20_ = nw541_
                d_2797_v21_: _dafny.Map
                d_2797_v21_ = _dafny.Map({_dafny.CodePoint('o'): d_2796_v20_})
                d_2797_v21_ = (d_2797_v21_).set((self).f5, d_2796_v20_)
                index555_ = default__.safeIndex(309, (d_2772_v1_).length(0))
                (d_2772_v1_)[index555_] = (self).f5
                d_2798_v22_: _dafny.Map
                d_2798_v22_ = _dafny.Map({(self).f7: False})
                d_2799_v23_: bool
                d_2799_v23_ = ((d_2798_v22_)[264] if (264) in (d_2798_v22_) else d_2778_v7_)
                d_2799_v23_ = d_2799_v23_
                d_2800_v24_: C1
                nw542_ = C1()
                nw542_.ctor__((self).f4, (self).f5)
                d_2800_v24_ = nw542_
        elif True:
            d_2801_v25_: _dafny.Array
            def lambda180_(d_2802_i2_):
                return (d_2802_i2_) * ((D14_DC31((self).f7)).cf48)

            init105_ = lambda180_
            nw543_ = _dafny.Array(None, 18)
            for i0_105_ in range(nw543_.length(0)):
                nw543_[i0_105_] = init105_(i0_105_)
            d_2801_v25_ = nw543_
            index556_ = default__.safeIndex(924, (d_2801_v25_).length(0))
            (d_2801_v25_)[index556_] = (self).f7
            d_2803_v26_: bool
            d_2803_v26_ = True
            d_2804_v27_: D13
            d_2804_v27_ = D13_DC29(d_2803_v26_, d_2803_v26_)
            d_2805_v28_: _dafny.MultiSet
            d_2805_v28_ = _dafny.MultiSet([(d_2804_v27_).cf46])
            index557_ = default__.safeIndex(924, (d_2801_v25_).length(0))
            (d_2801_v25_)[index557_] = (d_2805_v28_).cardinality
            d_2806_v29_: C2
            nw544_ = C2()
            nw544_.ctor__((self).f4, (self).f5)
            d_2806_v29_ = nw544_
            r1 = (d_2801_v25_)[default__.safeIndex(924, (d_2801_v25_).length(0))]
            d_2807_v30_: _dafny.Array
            def lambda181_(d_2808_i3_):
                return _dafny.Map({False: True})

            init106_ = lambda181_
            nw545_ = _dafny.Array(None, 16)
            for i0_106_ in range(nw545_.length(0)):
                nw545_[i0_106_] = init106_(i0_106_)
            d_2807_v30_ = nw545_
            d_2807_v30_ = d_2807_v30_
            d_2809_v31_: _dafny.Set
            d_2809_v31_ = _dafny.Set({d_2803_v26_, d_2803_v26_})
            d_2810_v32_: _dafny.Seq
            d_2810_v32_ = _dafny.SeqWithoutIsStrInference([((self).f7) == (len(d_2809_v31_))])
            d_2811_v33_: C5
            nw546_ = C5()
            nw546_.ctor__()
            d_2811_v33_ = nw546_
            d_2812_v34_: _dafny.Array
            nw547_ = _dafny.Array(None, 7)
            nw547_[int(0)] = d_2811_v33_
            nw547_[int(1)] = d_2811_v33_
            nw547_[int(2)] = d_2811_v33_
            nw547_[int(3)] = d_2811_v33_
            nw547_[int(4)] = d_2811_v33_
            nw547_[int(5)] = d_2811_v33_
            nw547_[int(6)] = d_2811_v33_
            d_2812_v34_ = nw547_
            d_2813_v35_: D21
            d_2813_v35_ = D21_DC51(d_2803_v26_, d_2811_v33_)
            d_2814_v36_: _dafny.Seq
            d_2814_v36_ = _dafny.SeqWithoutIsStrInference([d_2811_v33_, d_2811_v33_, (d_2813_v35_).cf81])
            index558_ = default__.safeIndex(55, (d_2812_v34_).length(0))
            (d_2812_v34_)[index558_] = (d_2814_v36_)[default__.safeIndex(default__.fm14(globalState), len(d_2814_v36_))]
            d_2815_v37_: _dafny.Map
            d_2815_v37_ = _dafny.Map({_dafny.CodePoint('j'): (self).fm5((self).f7, not(d_2803_v26_), (self).f7, (self).f7, globalState)})
            index559_ = default__.safeIndex(55, (d_2812_v34_).length(0))
            rhs557_ = ((d_2810_v32_).set(default__.safeIndex((self).f7, len(d_2810_v32_)), d_2803_v26_)).set(default__.safeIndex((d_2801_v25_)[default__.safeIndex(924, (d_2801_v25_).length(0))], len((d_2810_v32_).set(default__.safeIndex((self).f7, len(d_2810_v32_)), d_2803_v26_))), d_2803_v26_)
            rhs558_ = (d_2803_v26_ if d_2803_v26_ else (d_2806_v29_).fm5((d_2801_v25_)[default__.safeIndex(924, (d_2801_v25_).length(0))], ((d_2815_v37_)[(self).f5] if ((self).f5) in (d_2815_v37_) else d_2803_v26_), len(default__.fm2((self).f7, d_2803_v26_, d_2803_v26_, 457, globalState)), (self).f7, globalState))
            rhs559_ = d_2811_v33_
            lhs292_ = d_2812_v34_
            lhs293_ = default__.safeIndex(55, (d_2812_v34_).length(0))
            d_2810_v32_ = rhs557_
            d_2803_v26_ = rhs558_
            lhs292_[lhs293_] = rhs559_
        r1 = (0) - ((0) - (default__.safeDivisionInt((self).f7, (self).f7)))
        d_2816_v38_: bool
        d_2816_v38_ = False
        d_2817_v39_: _dafny.Seq
        d_2817_v39_ = _dafny.SeqWithoutIsStrInference([d_2816_v38_, False])
        d_2818_v40_: _dafny.Set
        d_2818_v40_ = _dafny.Set({(self).f7})
        d_2819_v41_: _dafny.Seq
        d_2819_v41_ = _dafny.SeqWithoutIsStrInference([d_2818_v40_, d_2818_v40_])
        d_2820_v42_: _dafny.Seq
        d_2820_v42_ = _dafny.SeqWithoutIsStrInference([d_2817_v39_])
        d_2821_v43_: _dafny.Map
        d_2821_v43_ = _dafny.Map({(self).f7: ((_dafny.SeqWithoutIsStrInference([d_2817_v39_])).set(default__.safeIndex(len(d_2819_v41_), len(_dafny.SeqWithoutIsStrInference([d_2817_v39_]))), _dafny.SeqWithoutIsStrInference([d_2816_v38_, d_2816_v38_, d_2816_v38_]))) + (d_2820_v42_)})
        d_2821_v43_ = (d_2821_v43_).set((self).f7, _dafny.SeqWithoutIsStrInference([d_2817_v39_, d_2817_v39_]))
        d_2822_v44_: _dafny.Map
        d_2822_v44_ = _dafny.Map({False: d_2816_v38_})
        d_2823_v45_: _dafny.Map
        d_2823_v45_ = d_2822_v44_
        d_2824_v46_: _dafny.Array
        nw548_ = _dafny.Array(None, 28)
        nw548_[int(0)] = d_2822_v44_
        nw548_[int(1)] = (d_2822_v44_) | (d_2822_v44_)
        nw548_[int(2)] = (_dafny.Map({d_2816_v38_: d_2816_v38_})) | (d_2822_v44_)
        nw548_[int(3)] = d_2822_v44_
        nw548_[int(4)] = d_2822_v44_
        nw548_[int(5)] = ((_dafny.Map({d_2816_v38_: d_2816_v38_})).set((d_2817_v39_)[default__.safeIndex((self).f7, len(d_2817_v39_))], d_2816_v38_)) | (d_2822_v44_)
        nw548_[int(6)] = (d_2822_v44_).set(d_2816_v38_, d_2816_v38_)
        nw548_[int(7)] = _dafny.Map({d_2816_v38_: d_2816_v38_})
        nw548_[int(8)] = (d_2822_v44_ if d_2816_v38_ else d_2822_v44_)
        nw548_[int(9)] = d_2822_v44_
        nw548_[int(10)] = d_2822_v44_
        nw548_[int(11)] = d_2822_v44_
        nw548_[int(12)] = d_2822_v44_
        nw548_[int(13)] = _dafny.Map({d_2816_v38_: d_2816_v38_})
        nw548_[int(14)] = _dafny.Map({d_2816_v38_: d_2816_v38_})
        nw548_[int(15)] = d_2822_v44_
        nw548_[int(16)] = (_dafny.Map({d_2816_v38_: d_2816_v38_})) | (_dafny.Map({d_2816_v38_: d_2816_v38_}))
        nw548_[int(17)] = _dafny.Map({d_2816_v38_: d_2816_v38_})
        nw548_[int(18)] = (d_2822_v44_) | (_dafny.Map({d_2816_v38_: d_2816_v38_}))
        nw548_[int(19)] = d_2822_v44_
        nw548_[int(20)] = d_2822_v44_
        nw548_[int(21)] = d_2822_v44_
        nw548_[int(22)] = (d_2822_v44_) | (d_2822_v44_)
        nw548_[int(23)] = (d_2822_v44_ if d_2816_v38_ else _dafny.Map({d_2816_v38_: True}))
        nw548_[int(24)] = _dafny.Map({True: not(True)})
        nw548_[int(25)] = (d_2823_v45_)
        nw548_[int(26)] = d_2822_v44_
        nw548_[int(27)] = d_2822_v44_
        d_2824_v46_ = nw548_
        index560_ = default__.safeIndex(109, (d_2824_v46_).length(0))
        (d_2824_v46_)[index560_] = d_2822_v44_
        d_2825_v47_: _dafny.Map
        d_2825_v47_ = _dafny.Map({(False if d_2816_v38_ else not(d_2816_v38_)): d_2822_v44_})
        index561_ = default__.safeIndex(109, (d_2824_v46_).length(0))
        (d_2824_v46_)[index561_] = ((d_2825_v47_)[d_2816_v38_] if (d_2816_v38_) in (d_2825_v47_) else d_2822_v44_)
        d_2826_v48_: _dafny.Array
        nw549_ = _dafny.Array(_dafny.Seq({}), 15)
        d_2826_v48_ = nw549_
        d_2827_v49_: D13
        d_2827_v49_ = D13_DC28(d_2826_v48_)
        source26_ = d_2827_v49_
        if source26_.is_DC29:
            d_2828___mcc_h0_ = source26_.cf45
            d_2829___mcc_h1_ = source26_.cf46
            d_2830_cf46_ = d_2829___mcc_h1_
            d_2831_cf45_ = d_2828___mcc_h0_
            d_2832_v50_: C0
            nw550_ = C0()
            nw550_.ctor__()
            d_2832_v50_ = nw550_
            d_2833_v51_: _dafny.Array
            nw551_ = _dafny.Array(D26.default()(), 20)
            d_2833_v51_ = nw551_
            d_2833_v51_ = d_2833_v51_
            if d_2830_cf46_:
                d_2830_cf46_ = d_2830_cf46_
                d_2834_v52_: _dafny.Seq
                d_2834_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa"))
                d_2835_v53_: _dafny.Seq
                d_2835_v53_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_2836_v54_: _dafny.Map
                d_2836_v54_ = _dafny.Map({d_2816_v38_: d_2835_v53_})
                d_2837_v55_: D1
                d_2837_v55_ = D1_DC2((self).f7, (self).f7)
                d_2838_v56_: _dafny.Set
                d_2838_v56_ = _dafny.Set({d_2834_v52_})
                d_2839_v57_: C5
                nw552_ = C5()
                nw552_.ctor__()
                d_2839_v57_ = nw552_
                d_2840_v58_: _dafny.Seq
                d_2840_v58_ = _dafny.SeqWithoutIsStrInference([d_2839_v57_])
                d_2841_v59_: D27
                d_2841_v59_ = D27_DC64(d_2840_v58_)
                d_2842_v60_: _dafny.Array
                nw553_ = _dafny.Array(None, 22)
                nw553_[int(0)] = (self).f7
                nw553_[int(1)] = len(d_2822_v44_)
                nw553_[int(2)] = 144
                nw553_[int(3)] = (self).f7
                nw553_[int(4)] = len(d_2834_v52_)
                nw553_[int(5)] = (self).f7
                nw553_[int(6)] = ((self).f7) - (len(((d_2836_v54_)[not(d_2830_cf46_)] if (not(d_2830_cf46_)) in (d_2836_v54_) else d_2835_v53_)))
                nw553_[int(7)] = ((self).f7) + (len(d_2835_v53_))
                nw553_[int(8)] = (self).f7
                nw553_[int(9)] = (self).f7
                nw553_[int(10)] = (self).f7
                nw553_[int(11)] = (self).f7
                nw553_[int(12)] = len(d_2834_v52_)
                nw553_[int(13)] = default__.fm14(globalState)
                nw553_[int(14)] = (493) * ((self).f7)
                nw553_[int(15)] = (self).fm7(d_2837_v55_, d_2831_cf45_, (self).f7, d_2816_v38_, globalState)
                nw553_[int(16)] = (self).f7
                nw553_[int(17)] = (self).f7
                nw553_[int(18)] = default__.safeModuloInt(len(d_2838_v56_), (self).f7)
                nw553_[int(19)] = ((self).f7) - ((self).f7)
                nw553_[int(20)] = len((d_2841_v59_).cf106)
                nw553_[int(21)] = (self).f7
                d_2842_v60_ = nw553_
                d_2842_v60_ = d_2842_v60_
                d_2830_cf46_ = d_2816_v38_
                d_2843_v61_: T1
                nw554_ = C10()
                nw554_.ctor__((0) - ((self).f7), d_2831_cf45_, (self).f4, (self).f5)
                d_2843_v61_ = nw554_
                r1 = (default__.safeDivisionInt(len(_dafny.Map({(self).f7: d_2843_v61_})), (self).f7)) * (895)
                d_2844_v62_: _dafny.Array
                nw555_ = _dafny.Array(False, 1)
                d_2844_v62_ = nw555_
                index562_ = default__.safeIndex(892, (d_2844_v62_).length(0))
                (d_2844_v62_)[index562_] = d_2816_v38_
                index563_ = default__.safeIndex(892, (d_2844_v62_).length(0))
                (d_2844_v62_)[index563_] = d_2816_v38_
            elif True:
                d_2845_v63_: str
                d_2845_v63_ = _dafny.CodePoint('g')
                d_2845_v63_ = d_2845_v63_
                r1 = (75) - ((self).f7)
                d_2846_v64_: _dafny.Seq
                d_2846_v64_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                d_2847_v65_: C2
                nw556_ = C2()
                nw556_.ctor__((self).f4, (d_2846_v64_)[default__.safeIndex((self).f7, len(d_2846_v64_))])
                d_2847_v65_ = nw556_
                d_2830_cf46_ = d_2831_cf45_
                d_2848_v66_: D26
                d_2848_v66_ = D26_DC63(d_2830_cf46_, (self).fm6((self).f7, (self).f7, globalState), 219, d_2831_cf45_)
                rhs560_ = not (d_2830_cf46_) or (d_2830_cf46_)
                rhs561_ = (d_2848_v66_).cf104
                d_2831_cf45_ = rhs560_
                r1 = rhs561_
            d_2849_v67_: C8
            nw557_ = C8()
            nw557_.ctor__()
            d_2849_v67_ = nw557_
        elif True:
            d_2850___mcc_h2_ = source26_.cf44
            d_2851_cf44_ = d_2850___mcc_h2_
            d_2852_v68_: _dafny.Array
            nw558_ = _dafny.Array(None, 2)
            nw558_[int(0)] = d_2816_v38_
            nw558_[int(1)] = d_2816_v38_
            d_2852_v68_ = nw558_
            index564_ = default__.safeIndex(573, (d_2852_v68_).length(0))
            (d_2852_v68_)[index564_] = (self).fm5((self).f7, d_2816_v38_, (self).f7, (self).f7, globalState)
            d_2853_v69_: _dafny.Seq
            d_2853_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxiotcre"))
            index565_ = default__.safeIndex(573, (d_2852_v68_).length(0))
            (d_2852_v68_)[index565_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhybnu"))) == (d_2853_v69_)
            d_2854_v70_: _dafny.MultiSet
            d_2854_v70_ = _dafny.MultiSet([(self).f7])
            d_2855_v71_: D2
            d_2855_v71_ = D2_DC4(not((self).fm5((self).f7, (d_2852_v68_)[default__.safeIndex(573, (d_2852_v68_).length(0))], ((d_2854_v70_)[(self).f7] if ((self).f7) in (d_2854_v70_) else 663), (self).f7, globalState)))
            d_2855_v71_ = d_2855_v71_
            d_2856_v72_: D23
            d_2856_v72_ = D23_DC56(not(d_2816_v38_), (d_2852_v68_)[default__.safeIndex(573, (d_2852_v68_).length(0))], default__.fm1((self).f7, (self).f7, (d_2852_v68_)[default__.safeIndex(573, (d_2852_v68_).length(0))], d_2816_v38_, globalState), d_2816_v38_)
            d_2816_v38_ = ((d_2822_v44_)[(d_2852_v68_)[default__.safeIndex(573, (d_2852_v68_).length(0))]] if ((d_2852_v68_)[default__.safeIndex(573, (d_2852_v68_).length(0))]) in (d_2822_v44_) else (d_2856_v72_).cf92)
            d_2857_v73_: C7
            nw559_ = C7()
            nw559_.ctor__()
            d_2857_v73_ = nw559_
            nw560_ = C7()
            nw560_.ctor__()
            d_2857_v73_ = nw560_
        d_2858_v74_: _dafny.Array
        def lambda182_(d_2859_i4_):
            return _dafny.CodePoint('d')

        init107_ = lambda182_
        nw561_ = _dafny.Array(None, 18)
        for i0_107_ in range(nw561_.length(0)):
            nw561_[i0_107_] = init107_(i0_107_)
        d_2858_v74_ = nw561_
        d_2860_v75_: _dafny.Seq
        d_2860_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cs"))
        index566_ = default__.safeIndex(838, (d_2858_v74_).length(0))
        (d_2858_v74_)[index566_] = (d_2860_v75_)[default__.safeIndex((self).f7, len(d_2860_v75_))]
        index567_ = default__.safeIndex(838, (d_2858_v74_).length(0))
        rhs562_ = ((self).f5 if d_2816_v38_ else (self).f5)
        rhs563_ = default__.safeDivisionInt((self).f7, (self).f7)
        rhs564_ = (not(d_2816_v38_) if ((self).f7) == ((self).f7) else d_2816_v38_)
        lhs294_ = d_2858_v74_
        lhs295_ = default__.safeIndex(838, (d_2858_v74_).length(0))
        lhs294_[lhs295_] = rhs562_
        r1 = rhs563_
        d_2816_v38_ = rhs564_
        d_2861_v76_: _dafny.Map
        d_2861_v76_ = _dafny.Map({d_2816_v38_: (self).f7})
        d_2862_v77_: _dafny.Seq
        d_2862_v77_ = _dafny.SeqWithoutIsStrInference([(self).f7, (0) - (default__.safeDivisionInt((self).f7, len(d_2861_v76_)))])
        r0 = d_2862_v77_
        r1 = (self).f7
        return r0, r1

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_2863_v0_: C0
        nw562_ = C0()
        nw562_.ctor__()
        d_2863_v0_ = nw562_
        d_2864_v1_: bool
        d_2864_v1_ = False
        if d_2864_v1_:
            d_2865_v2_: T1
            nw563_ = C10()
            nw563_.ctor__(default__.safeDivisionInt((self).f7, (self).f7), d_2864_v1_, (self).f4, _dafny.CodePoint('y'))
            d_2865_v2_ = nw563_
            d_2865_v2_ = d_2865_v2_
            d_2866_v3_: _dafny.Array
            nw564_ = _dafny.Array(int(0), 24)
            d_2866_v3_ = nw564_
            d_2867_v4_: _dafny.Map
            d_2867_v4_ = _dafny.Map({d_2866_v3_: d_2864_v1_})
            d_2868_v5_: _dafny.Seq
            d_2868_v5_ = _dafny.SeqWithoutIsStrInference([(self).f7, len(_dafny.SeqWithoutIsStrInference([d_2864_v1_]))])
            d_2869_v6_: _dafny.Set
            d_2869_v6_ = _dafny.Set({d_2864_v1_, not((d_2866_v3_) not in (d_2867_v4_)), ((self).f7) != ((d_2868_v5_)[default__.safeIndex(611, len(d_2868_v5_))])})
            d_2869_v6_ = (d_2869_v6_) - ((d_2869_v6_) - (d_2869_v6_))
            index568_ = default__.safeIndex(638, (d_2866_v3_).length(0))
            (d_2866_v3_)[index568_] = (self).f7
            index569_ = default__.safeIndex(638, (d_2866_v3_).length(0))
            (d_2866_v3_)[index569_] = (self).f7
            index570_ = default__.safeIndex(638, (d_2866_v3_).length(0))
            rhs565_ = d_2864_v1_
            rhs566_ = (d_2866_v3_)[default__.safeIndex(638, (d_2866_v3_).length(0))]
            lhs296_ = d_2866_v3_
            lhs297_ = default__.safeIndex(638, (d_2866_v3_).length(0))
            d_2864_v1_ = rhs565_
            lhs296_[lhs297_] = rhs566_
            index571_ = default__.safeIndex(7, ((d_2865_v2_).f4).length(0))
            ((d_2865_v2_).f4)[index571_] = d_2869_v6_
            d_2870_v7_: D15
            d_2870_v7_ = D15_DC35((d_2865_v2_).fm4(_dafny.SeqWithoutIsStrInference([(self).f7, (d_2866_v3_)[default__.safeIndex(638, (d_2866_v3_).length(0))]]), globalState))
            index572_ = default__.safeIndex(7, ((d_2865_v2_).f4).length(0))
            index573_ = default__.safeIndex(638, (d_2866_v3_).length(0))
            rhs567_ = d_2864_v1_
            rhs568_ = not (d_2864_v1_) or (d_2864_v1_)
            rhs569_ = (d_2869_v6_) | (_dafny.Set({True}))
            rhs570_ = (d_2870_v7_).cf54
            lhs298_ = (d_2865_v2_).f4
            lhs299_ = default__.safeIndex(7, ((d_2865_v2_).f4).length(0))
            lhs300_ = d_2866_v3_
            lhs301_ = default__.safeIndex(638, (d_2866_v3_).length(0))
            d_2864_v1_ = rhs567_
            d_2864_v1_ = rhs568_
            lhs298_[lhs299_] = rhs569_
            lhs300_[lhs301_] = rhs570_
        elif True:
            d_2871_v8_: _dafny.Array
            nw565_ = _dafny.Array(int(0), 23)
            d_2871_v8_ = nw565_
            d_2872_v9_: _dafny.Map
            d_2872_v9_ = _dafny.Map({d_2864_v1_: len(_dafny.Set({(self).f7}))})
            index574_ = default__.safeIndex(103, (d_2871_v8_).length(0))
            (d_2871_v8_)[index574_] = len((d_2872_v9_) | ((d_2872_v9_).set(d_2864_v1_, (0) - ((self).f7))))
            index575_ = default__.safeIndex(103, (d_2871_v8_).length(0))
            (d_2871_v8_)[index575_] = (self).f7
            d_2873_v10_: D11
            d_2873_v10_ = D11_DC25(d_2864_v1_, (self).f7, d_2871_v8_, (self).f7)
            d_2874_v11_: _dafny.Seq
            d_2874_v11_ = _dafny.SeqWithoutIsStrInference([d_2871_v8_, d_2871_v8_, d_2871_v8_])
            d_2875_v12_: _dafny.Array
            nw566_ = _dafny.Array(None, 21)
            nw566_[int(0)] = d_2871_v8_
            nw566_[int(1)] = (d_2873_v10_).cf40
            nw566_[int(2)] = d_2871_v8_
            nw566_[int(3)] = d_2871_v8_
            nw566_[int(4)] = (d_2873_v10_).cf40
            nw566_[int(5)] = d_2871_v8_
            nw566_[int(6)] = d_2871_v8_
            nw566_[int(7)] = d_2871_v8_
            nw566_[int(8)] = (d_2873_v10_).cf40
            nw566_[int(9)] = d_2871_v8_
            nw566_[int(10)] = d_2871_v8_
            nw566_[int(11)] = d_2871_v8_
            nw566_[int(12)] = d_2871_v8_
            nw566_[int(13)] = (d_2874_v11_)[default__.safeIndex((d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], len(d_2874_v11_))]
            nw566_[int(14)] = d_2871_v8_
            nw566_[int(15)] = d_2871_v8_
            nw566_[int(16)] = d_2871_v8_
            nw566_[int(17)] = d_2871_v8_
            nw566_[int(18)] = d_2871_v8_
            nw566_[int(19)] = d_2871_v8_
            nw566_[int(20)] = d_2871_v8_
            d_2875_v12_ = nw566_
            d_2876_v13_: _dafny.Array
            nw567_ = _dafny.Array(False, 15)
            d_2876_v13_ = nw567_
            d_2877_v14_: _dafny.Map
            d_2877_v14_ = _dafny.Map({d_2875_v12_: d_2876_v13_})
            d_2878_v15_: _dafny.Set
            d_2878_v15_ = _dafny.Set({(self).f7})
            index576_ = default__.safeIndex(103, (d_2871_v8_).length(0))
            index577_ = default__.safeIndex(103, (d_2871_v8_).length(0))
            rhs571_ = (-119) - ((d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))])
            rhs572_ = ((d_2877_v14_)[d_2875_v12_] if (d_2875_v12_) in (d_2877_v14_) else d_2876_v13_)
            rhs573_ = default__.safeDivisionInt(len(d_2878_v15_), (0) - ((d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))]))
            lhs302_ = d_2871_v8_
            lhs303_ = default__.safeIndex(103, (d_2871_v8_).length(0))
            lhs304_ = globalState
            lhs305_ = d_2871_v8_
            lhs306_ = default__.safeIndex(103, (d_2871_v8_).length(0))
            lhs302_[lhs303_] = rhs571_
            lhs304_.f3 = rhs572_
            lhs305_[lhs306_] = rhs573_
            d_2864_v1_ = not(d_2864_v1_)
            d_2879_v16_: _dafny.Set
            d_2879_v16_ = _dafny.Set({d_2864_v1_})
            if (d_2879_v16_).ispropersubset((d_2879_v16_).intersection(d_2879_v16_)):
                d_2880_v17_: C3
                nw568_ = C3()
                nw568_.ctor__((self).f4, (self).f5)
                d_2880_v17_ = nw568_
                d_2881_v18_: D26
                d_2881_v18_ = D26_DC63(d_2864_v1_, (self).f7, (d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], d_2864_v1_)
                d_2864_v1_ = (d_2881_v18_).cf102
                d_2882_v19_: _dafny.Map
                d_2882_v19_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))] for d_2883_i0_ in range(default__.abs(-874))]): (self).f7})
                d_2884_v20_: _dafny.Seq
                d_2884_v20_ = _dafny.SeqWithoutIsStrInference([(d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], (d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))]])
                d_2885_v21_: D1
                d_2885_v21_ = D1_DC2((d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], (d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))])
                d_2882_v19_ = (d_2882_v19_).set(d_2884_v20_, (self).fm7(d_2885_v21_, default__.fm1((self).f7, (d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], d_2864_v1_, d_2864_v1_, globalState), (d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], d_2864_v1_, globalState))
                index578_ = default__.safeIndex(103, (d_2871_v8_).length(0))
                (d_2871_v8_)[index578_] = (self).f7
                index579_ = default__.safeIndex(103, (d_2871_v8_).length(0))
                (d_2871_v8_)[index579_] = (self).f7
            elif True:
                d_2886_v22_: _dafny.MultiSet
                d_2886_v22_ = _dafny.MultiSet([(self).f7])
                d_2864_v1_ = not(not (((d_2886_v22_).set((d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], default__.abs(50))).issubset(d_2886_v22_)) or ((d_2864_v1_) or (d_2864_v1_)))
                d_2887_v23_: _dafny.Seq
                d_2887_v23_ = _dafny.SeqWithoutIsStrInference([d_2864_v1_, not(d_2864_v1_), d_2864_v1_, False, d_2864_v1_])
                d_2887_v23_ = (d_2887_v23_) + (_dafny.SeqWithoutIsStrInference([not(d_2864_v1_), d_2864_v1_, d_2864_v1_, d_2864_v1_]))
                index580_ = default__.safeIndex(103, (d_2871_v8_).length(0))
                (d_2871_v8_)[index580_] = (self).f7
                d_2888_v24_: _dafny.Seq
                d_2888_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjves"))
                d_2889_v25_: _dafny.Map
                d_2889_v25_ = _dafny.Map({d_2888_v24_: d_2864_v1_})
                d_2890_v26_: C5
                nw569_ = C5()
                nw569_.ctor__()
                d_2890_v26_ = nw569_
                d_2891_v27_: _dafny.Seq
                d_2891_v27_ = _dafny.SeqWithoutIsStrInference([d_2890_v26_])
                d_2892_v28_: D21
                d_2892_v28_ = D21_DC51((d_2887_v23_)[default__.safeIndex((self).f7, len(d_2887_v23_))], (d_2891_v27_)[default__.safeIndex((self).f7, len(d_2891_v27_))])
                d_2893_v29_: D21
                d_2893_v29_ = D21_DC51(not(d_2864_v1_), (d_2892_v28_).cf81)
                rhs574_ = ((d_2864_v1_) and ((d_2893_v29_).cf80)) or ((d_2864_v1_ if (d_2887_v23_)[default__.safeIndex((d_2871_v8_)[default__.safeIndex(103, (d_2871_v8_).length(0))], len(d_2887_v23_))] else d_2864_v1_))
                rhs575_ = (d_2889_v25_) | (d_2889_v25_)
                d_2864_v1_ = rhs574_
                d_2889_v25_ = rhs575_
                d_2894_v30_: _dafny.MultiSet
                d_2894_v30_ = _dafny.MultiSet([d_2864_v1_])
                d_2864_v1_ = not((d_2894_v30_).issubset((d_2894_v30_).set(d_2864_v1_, default__.abs((self).f7))))
            d_2895_v32_: C6
            nw570_ = C6()
            def iife195_():
                coll73_ = _dafny.Map()
                compr_73_: int
                for compr_73_ in _dafny.IntegerRange(72, 862):
                    d_2896_v31_: int = compr_73_
                    if ((72) <= (d_2896_v31_)) and ((d_2896_v31_) < (862)):
                        coll73_[(d_2896_v31_) * (465)] = (0) - ((self).f7)
                return _dafny.Map(coll73_)
            nw570_.ctor__(iife195_()
            )
            d_2895_v32_ = nw570_
            d_2895_v32_ = d_2895_v32_
        d_2897_v33_: D7
        d_2897_v33_ = D7_DC17(d_2864_v1_, d_2864_v1_, d_2864_v1_, d_2864_v1_)
        source27_ = d_2897_v33_
        if source27_.is_DC17:
            d_2898___mcc_h0_ = source27_.cf24
            d_2899___mcc_h1_ = source27_.cf25
            d_2900___mcc_h2_ = source27_.cf26
            d_2901___mcc_h3_ = source27_.cf27
            d_2902_cf27_ = d_2901___mcc_h3_
            d_2903_cf26_ = d_2900___mcc_h2_
            d_2904_cf25_ = d_2899___mcc_h1_
            d_2905_cf24_ = d_2898___mcc_h0_
            d_2904_cf25_ = d_2864_v1_
            d_2906_v34_: _dafny.Seq
            d_2906_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qybbhcq"))
            d_2907_v35_: T1
            nw571_ = C10()
            nw571_.ctor__(len(d_2906_v34_), d_2903_cf26_, (self).f4, (self).f5)
            d_2907_v35_ = nw571_
            d_2908_v36_: _dafny.Array
            nw572_ = _dafny.Array(False, 1)
            d_2908_v36_ = nw572_
            d_2909_v37_: _dafny.Map
            d_2909_v37_ = _dafny.Map({d_2907_v35_: d_2908_v36_})
            d_2909_v37_ = _dafny.Map({d_2907_v35_: d_2908_v36_})
            d_2904_cf25_ = d_2903_cf26_
            d_2910_v38_: int
            d_2910_v38_ = 567
            d_2911_v39_: _dafny.Set
            d_2911_v39_ = _dafny.Set({d_2905_cf24_, d_2903_cf26_})
            index581_ = default__.safeIndex(138, ((d_2907_v35_).f4).length(0))
            ((d_2907_v35_).f4)[index581_] = d_2911_v39_
            index582_ = default__.safeIndex(138, ((d_2907_v35_).f4).length(0))
            rhs576_ = d_2903_cf26_
            rhs577_ = len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_2912_i1_ in range(default__.abs(303))]))
            rhs578_ = (d_2911_v39_) - (d_2911_v39_)
            lhs307_ = (d_2907_v35_).f4
            lhs308_ = default__.safeIndex(138, ((d_2907_v35_).f4).length(0))
            d_2905_cf24_ = rhs576_
            d_2910_v38_ = rhs577_
            lhs307_[lhs308_] = rhs578_
        elif True:
            d_2913___mcc_h4_ = source27_.cf23
            d_2914_cf23_ = d_2913___mcc_h4_
            d_2915_v40_: _dafny.Array
            def lambda183_(d_2916_i2_):
                return default__.safeModuloInt(d_2916_i2_, (self).f7)

            init108_ = lambda183_
            nw573_ = _dafny.Array(None, 29)
            for i0_108_ in range(nw573_.length(0)):
                nw573_[i0_108_] = init108_(i0_108_)
            d_2915_v40_ = nw573_
            d_2917_v41_: bool
            d_2918_v42_: bool
            d_2919_v43_: _dafny.Seq
            out33_: bool
            out34_: bool
            out35_: _dafny.Seq
            out33_, out34_, out35_ = (self).m3((self).f7, d_2915_v40_, globalState)
            d_2917_v41_ = out33_
            d_2918_v42_ = out34_
            d_2919_v43_ = out35_
            d_2920_v44_: C5
            nw574_ = C5()
            nw574_.ctor__()
            d_2920_v44_ = nw574_
            d_2921_v45_: _dafny.MultiSet
            d_2921_v45_ = _dafny.MultiSet([(self).f7, (0) - ((self).f7), (self).f7, (self).f7])
            d_2922_v46_: _dafny.MultiSet
            d_2922_v46_ = _dafny.MultiSet([(d_2921_v45_).cardinality, (self).f7, (0) - ((self).f7), (self).f7, (self).f7])
            d_2923_v47_: _dafny.Map
            d_2923_v47_ = _dafny.Map({(self).f7: ((d_2921_v45_)[(self).f7] if ((self).f7) in (d_2921_v45_) else -441)})
            d_2924_v48_: _dafny.Seq
            d_2924_v48_ = _dafny.SeqWithoutIsStrInference([(self).f7])
            d_2925_v49_: _dafny.Array
            nw575_ = _dafny.Array(None, 14)
            nw575_[int(0)] = (d_2864_v1_) == (d_2864_v1_)
            nw575_[int(1)] = d_2864_v1_
            nw575_[int(2)] = d_2917_v41_
            nw575_[int(3)] = False
            nw575_[int(4)] = default__.fm1(((d_2923_v47_)[(self).f7] if ((self).f7) in (d_2923_v47_) else (self).f7), (_dafny.MultiSet(d_2924_v48_)).cardinality, d_2917_v41_, d_2864_v1_, globalState)
            nw575_[int(5)] = d_2917_v41_
            nw575_[int(6)] = d_2918_v42_
            nw575_[int(7)] = not(d_2917_v41_)
            nw575_[int(8)] = d_2864_v1_
            nw575_[int(9)] = d_2864_v1_
            nw575_[int(10)] = d_2918_v42_
            nw575_[int(11)] = d_2864_v1_
            nw575_[int(12)] = d_2864_v1_
            nw575_[int(13)] = d_2864_v1_
            d_2925_v49_ = nw575_
            (globalState).f3 = d_2925_v49_
            d_2926_v50_: _dafny.Seq
            d_2926_v50_ = _dafny.SeqWithoutIsStrInference([(default__.fm49(globalState)).cf75])
            d_2927_v52_: _dafny.Map
            d_2927_v52_ = _dafny.Map({(self).f5: d_2917_v41_})
            d_2928_v53_: _dafny.Set
            d_2928_v53_ = _dafny.Set({d_2918_v42_, d_2864_v1_, (self).fm3((self).f7, len(d_2927_v52_), globalState)})
            d_2929_v54_: _dafny.Set
            d_2929_v54_ = _dafny.Set({default__.fm21((self).f7, len(d_2924_v48_), d_2928_v53_, globalState), _dafny.SeqWithoutIsStrInference([d_2917_v41_, d_2917_v41_])})
            d_2930_v55_: _dafny.Map
            d_2930_v55_ = _dafny.Map({d_2917_v41_: d_2929_v54_})
            def iife196_():
                coll74_ = _dafny.Set()
                compr_74_: _dafny.Seq
                for compr_74_ in (d_2926_v50_).Elements:
                    d_2931_v51_: _dafny.Seq = compr_74_
                    if (d_2931_v51_) in (d_2926_v50_):
                        coll74_ = coll74_.union(_dafny.Set([d_2931_v51_]))
                return _dafny.Set(coll74_)
            if (iife196_()
            ).isdisjoint(((d_2930_v55_)[d_2864_v1_] if (d_2864_v1_) in (d_2930_v55_) else default__.fm50(d_2918_v42_, d_2864_v1_, d_2918_v42_, globalState))):
                d_2918_v42_ = d_2917_v41_
                d_2922_v46_ = _dafny.MultiSet(d_2924_v48_)
                d_2932_v56_: C6
                nw576_ = C6()
                nw576_.ctor__(d_2923_v47_)
                d_2932_v56_ = nw576_
                d_2933_v57_: _dafny.MultiSet
                d_2933_v57_ = _dafny.MultiSet([_dafny.Map({d_2932_v56_: (self).fm5((self).f7, not(not(d_2864_v1_)), (self).f7, (self).f7, globalState)})])
                d_2934_v58_: _dafny.Seq
                d_2934_v58_ = _dafny.SeqWithoutIsStrInference([False])
                d_2935_v59_: _dafny.Map
                d_2935_v59_ = _dafny.Map({True: (d_2934_v58_)[default__.safeIndex((self).f7, len(d_2934_v58_))]})
                d_2936_v60_: _dafny.Seq
                d_2937_v61_: int
                d_2938_v62_: str
                d_2939_v63_: int
                out36_: _dafny.Seq
                out37_: int
                out38_: str
                out39_: int
                out36_, out37_, out38_, out39_ = (d_2920_v44_).m10((d_2933_v57_).cardinality, d_2935_v59_, d_2918_v42_, globalState)
                d_2936_v60_ = out36_
                d_2937_v61_ = out37_
                d_2938_v62_ = out38_
                d_2939_v63_ = out39_
                index583_ = default__.safeIndex(322, (d_2925_v49_).length(0))
                (d_2925_v49_)[index583_] = d_2864_v1_
                index584_ = default__.safeIndex(322, (d_2925_v49_).length(0))
                (d_2925_v49_)[index584_] = d_2864_v1_
                d_2940_v64_: C2
                nw577_ = C2()
                nw577_.ctor__((self).f4, (self).f5)
                d_2940_v64_ = nw577_
            elif True:
                d_2918_v42_ = not(d_2918_v42_)
                r0 = d_2915_v40_
                d_2941_v65_: _dafny.Array
                nw578_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_2941_v65_ = nw578_
                index585_ = default__.safeIndex(640, (d_2941_v65_).length(0))
                (d_2941_v65_)[index585_] = d_2919_v43_
                d_2942_v66_: _dafny.Seq
                d_2942_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "insl")), _dafny.SeqWithoutIsStrInference([(self).f5 for d_2943_i3_ in range(default__.abs(652))]), d_2919_v43_, d_2919_v43_])
                d_2944_v67_: _dafny.Array
                nw579_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_2944_v67_ = nw579_
                d_2945_v68_: _dafny.MultiSet
                d_2945_v68_ = _dafny.MultiSet([d_2944_v67_])
                d_2946_v69_: _dafny.Map
                d_2946_v69_ = _dafny.Map({d_2917_v41_: (self).f7})
                index586_ = default__.safeIndex(640, (d_2941_v65_).length(0))
                (d_2941_v65_)[index586_] = (d_2942_v66_)[default__.safeIndex(default__.safeModuloInt(((d_2945_v68_)[d_2944_v67_] if (d_2944_v67_) in (d_2945_v68_) else ((d_2946_v69_)[d_2918_v42_] if (d_2918_v42_) in (d_2946_v69_) else (0) - ((self).f7))), (self).f7), len(d_2942_v66_))]
                d_2947_v70_: D7
                d_2947_v70_ = D7_DC16(d_2914_cf23_)
                d_2948_v71_: C9
                nw580_ = C9()
                nw580_.ctor__((self).f4, (d_2947_v70_).cf23)
                d_2948_v71_ = nw580_
                d_2949_v72_: int
                d_2949_v72_ = 889
                d_2950_v73_: _dafny.Seq
                d_2950_v73_ = _dafny.SeqWithoutIsStrInference([d_2920_v44_])
                d_2951_v74_: _dafny.Seq
                d_2951_v74_ = _dafny.SeqWithoutIsStrInference([(d_2950_v73_)[default__.safeIndex(444, len(d_2950_v73_))], d_2920_v44_])
                d_2952_v75_: D18
                d_2952_v75_ = D18_DC44(_dafny.Set({d_2928_v53_}))
                d_2953_v76_: _dafny.Seq
                d_2953_v76_ = _dafny.SeqWithoutIsStrInference([d_2952_v75_])
                d_2954_v77_: _dafny.Seq
                d_2954_v77_ = _dafny.SeqWithoutIsStrInference([not(d_2917_v41_)])
                d_2949_v72_ = ((len(d_2953_v76_)) * (len((d_2941_v65_)[default__.safeIndex(640, (d_2941_v65_).length(0))])) if (d_2951_v74_) == (d_2950_v73_) else len((d_2954_v77_ if d_2864_v1_ else d_2954_v77_)))
        d_2955_v78_: int
        d_2955_v78_ = 824
        d_2956_v79_: _dafny.Seq
        d_2956_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
        rhs579_ = 116
        rhs580_ = ((d_2956_v79_ if d_2864_v1_ else d_2956_v79_)) + (d_2956_v79_)
        rhs581_ = d_2864_v1_
        rhs582_ = d_2955_v78_
        d_2955_v78_ = rhs579_
        d_2956_v79_ = rhs580_
        d_2864_v1_ = rhs581_
        d_2955_v78_ = rhs582_
        d_2957_v80_: _dafny.Seq
        d_2957_v80_ = _dafny.SeqWithoutIsStrInference([d_2864_v1_])
        d_2958_v81_: D20
        d_2958_v81_ = D20_DC48(d_2864_v1_, d_2957_v80_, d_2957_v80_, d_2955_v78_)
        d_2959_v82_: D20
        d_2959_v82_ = D20_DC49(d_2958_v81_)
        source28_ = D20_DC49(d_2958_v81_)
        if source28_.is_DC48:
            d_2960___mcc_h5_ = source28_.cf74
            d_2961___mcc_h6_ = source28_.cf75
            d_2962___mcc_h7_ = source28_.cf76
            d_2963___mcc_h8_ = source28_.cf77
            d_2964_cf77_ = d_2963___mcc_h8_
            d_2965_cf76_ = d_2962___mcc_h7_
            d_2966_cf75_ = d_2961___mcc_h6_
            d_2967_cf74_ = d_2960___mcc_h5_
            d_2967_cf74_ = d_2864_v1_
            d_2968_v83_: C5
            nw581_ = C5()
            nw581_.ctor__()
            d_2968_v83_ = nw581_
            d_2969_v84_: _dafny.Seq
            d_2969_v84_ = _dafny.SeqWithoutIsStrInference([d_2968_v83_, d_2968_v83_])
            d_2970_v85_: D27
            d_2970_v85_ = D27_DC64(d_2969_v84_)
            pat_let_tv57_ = d_2969_v84_
            pat_let_tv58_ = d_2964_cf77_
            pat_let_tv59_ = d_2969_v84_
            pat_let_tv60_ = d_2968_v83_
            def iife197_(_pat_let61_0):
                def iife198_(d_2971_dt__update__tmp_h0_):
                    def iife199_(_pat_let62_0):
                        def iife200_(d_2972_dt__update_hcf106_h0_):
                            return D27_DC64(d_2972_dt__update_hcf106_h0_)
                        return iife200_(_pat_let62_0)
                    return iife199_((pat_let_tv57_).set(default__.safeIndex(pat_let_tv58_, len(pat_let_tv59_)), pat_let_tv60_))
                return iife198_(_pat_let61_0)
            d_2970_v85_ = iife197_(d_2970_v85_)
            d_2973_v86_: D13
            d_2973_v86_ = D13_DC29(d_2864_v1_, d_2864_v1_)
            d_2974_v87_: _dafny.Set
            d_2974_v87_ = _dafny.Set({(self).f7})
            d_2975_v88_: D16
            d_2975_v88_ = D16_DC39(D16_DC37(d_2974_v87_))
            d_2976_v89_: D16
            d_2976_v89_ = D16_DC37(d_2974_v87_)
            d_2977_v90_: _dafny.MultiSet
            d_2977_v90_ = _dafny.MultiSet([d_2975_v88_, D16_DC39(d_2976_v89_), d_2975_v88_])
            d_2978_v91_: _dafny.MultiSet
            d_2978_v91_ = d_2977_v90_
            d_2979_v92_: _dafny.MultiSet
            d_2979_v92_ = (d_2978_v91_)
            d_2980_v93_: _dafny.Map
            d_2980_v93_ = _dafny.Map({(self).f5: d_2979_v92_})
            def iife201_():
                coll75_ = _dafny.Set()
                compr_75_: str
                for compr_75_ in (d_2980_v93_).keys.Elements:
                    d_2981_v94_: str = compr_75_
                    if (d_2981_v94_) in (d_2980_v93_):
                        coll75_ = coll75_.union(_dafny.Set([d_2981_v94_]))
                return _dafny.Set(coll75_)
            rhs583_ = default__.safeDivisionInt((d_2964_cf77_) - ((self).f7), default__.safeModuloInt((self).f7, d_2955_v78_))
            rhs584_ = default__.safeDivisionInt(d_2955_v78_, d_2964_cf77_)
            rhs585_ = d_2973_v86_
            rhs586_ = d_2955_v78_
            rhs587_ = (len(iife201_()
            )) > (-443)
            d_2955_v78_ = rhs583_
            d_2955_v78_ = rhs584_
            d_2973_v86_ = rhs585_
            d_2955_v78_ = rhs586_
            d_2864_v1_ = rhs587_
            d_2982_v95_: _dafny.Array
            nw582_ = _dafny.Array(int(0), 19)
            d_2982_v95_ = nw582_
            index587_ = default__.safeIndex(82, (d_2982_v95_).length(0))
            (d_2982_v95_)[index587_] = (self).f7
            d_2983_v96_: _dafny.Array
            nw583_ = _dafny.Array(None, 14)
            nw583_[int(0)] = d_2967_cf74_
            nw583_[int(1)] = d_2864_v1_
            nw583_[int(2)] = d_2967_cf74_
            nw583_[int(3)] = d_2967_cf74_
            nw583_[int(4)] = True
            nw583_[int(5)] = d_2967_cf74_
            nw583_[int(6)] = d_2967_cf74_
            nw583_[int(7)] = d_2864_v1_
            nw583_[int(8)] = d_2967_cf74_
            nw583_[int(9)] = d_2967_cf74_
            nw583_[int(10)] = d_2864_v1_
            nw583_[int(11)] = d_2864_v1_
            nw583_[int(12)] = d_2967_cf74_
            nw583_[int(13)] = d_2864_v1_
            d_2983_v96_ = nw583_
            d_2984_v97_: _dafny.Set
            d_2984_v97_ = _dafny.Set({d_2983_v96_})
            d_2985_v98_: _dafny.Set
            d_2985_v98_ = d_2984_v97_
            d_2986_v99_: _dafny.Array
            def lambda184_(d_2987_cf77_, d_2988_v1_):
                def lambda185_(d_2989_i4_):
                    return _dafny.Map({d_2987_cf77_: d_2988_v1_})

                return lambda185_

            init109_ = lambda184_(d_2964_cf77_, d_2864_v1_)
            nw584_ = _dafny.Array(None, 18)
            for i0_109_ in range(nw584_.length(0)):
                nw584_[i0_109_] = init109_(i0_109_)
            d_2986_v99_ = nw584_
            d_2990_v100_: D4
            d_2990_v100_ = D4_DC10((d_2985_v98_ if False else d_2985_v98_), len(d_2974_v87_), default__.fm16(False, d_2955_v78_, globalState), d_2986_v99_)
            d_2991_v101_: _dafny.Seq
            d_2991_v101_ = _dafny.SeqWithoutIsStrInference([d_2955_v78_])
            index588_ = default__.safeIndex(82, (d_2982_v95_).length(0))
            rhs588_ = len((d_2991_v101_ if d_2967_cf74_ else d_2991_v101_))
            rhs589_ = d_2955_v78_
            rhs590_ = d_2990_v100_
            rhs591_ = (d_2957_v80_).set(default__.safeIndex((d_2955_v78_ if d_2967_cf74_ else d_2964_cf77_), len(d_2957_v80_)), not(d_2864_v1_))
            rhs592_ = d_2955_v78_
            lhs309_ = d_2982_v95_
            lhs310_ = default__.safeIndex(82, (d_2982_v95_).length(0))
            lhs309_[lhs310_] = rhs588_
            d_2955_v78_ = rhs589_
            d_2990_v100_ = rhs590_
            d_2965_cf76_ = rhs591_
            d_2955_v78_ = rhs592_
        elif source28_.is_DC47:
            d_2992___mcc_h9_ = source28_.cf73
            d_2993_cf73_ = d_2992___mcc_h9_
            d_2864_v1_ = not(d_2864_v1_)
            d_2994_v102_: C8
            nw585_ = C8()
            nw585_.ctor__()
            d_2994_v102_ = nw585_
            d_2995_v103_: D8
            d_2995_v103_ = D8_DC19(len(d_2956_v79_), True)
            d_2996_v104_: D8
            d_2996_v104_ = D8_DC19(len(d_2957_v80_), (d_2995_v103_).cf30)
            d_2997_v105_: _dafny.Array
            nw586_ = _dafny.Array(None, 9)
            nw586_[int(0)] = (self).f7
            nw586_[int(1)] = (self).f7
            nw586_[int(2)] = d_2955_v78_
            nw586_[int(3)] = d_2955_v78_
            nw586_[int(4)] = (self).f7
            nw586_[int(5)] = (self).f7
            nw586_[int(6)] = (self).f7
            nw586_[int(7)] = (d_2996_v104_).cf29
            nw586_[int(8)] = (self).f7
            d_2997_v105_ = nw586_
            d_2998_v106_: _dafny.MultiSet
            d_2998_v106_ = _dafny.MultiSet([d_2997_v105_])
            d_2999_v107_: _dafny.Map
            d_2999_v107_ = _dafny.Map({d_2864_v1_: _dafny.MultiSet([d_2997_v105_])})
            d_2998_v106_ = (d_2998_v106_ if d_2864_v1_ else (((d_2999_v107_)[d_2864_v1_] if (d_2864_v1_) in (d_2999_v107_) else d_2998_v106_)) - (d_2998_v106_))
            d_2955_v78_ = d_2955_v78_
        elif True:
            d_3000___mcc_h10_ = source28_.cf78
            d_3001_cf78_ = d_3000___mcc_h10_
            d_3002_v108_: _dafny.Seq
            d_3002_v108_ = _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([(self).f5 for d_3003_i5_ in range(default__.abs(-718))])).set(default__.safeIndex((self).f7, len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_3004_i5_ in range(default__.abs(-718))]))), (self).f5)), (_dafny.MultiSet([d_2864_v1_])).cardinality, d_2955_v78_])
            d_3005_v109_: _dafny.MultiSet
            d_3005_v109_ = _dafny.MultiSet([d_2864_v1_, d_2864_v1_])
            d_3006_v110_: _dafny.Map
            d_3006_v110_ = _dafny.Map({(self).f7: d_3002_v108_})
            d_3007_v111_: _dafny.Array
            nw587_ = _dafny.Array(None, 29)
            nw587_[int(0)] = d_3002_v108_
            nw587_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2955_v78_])
            nw587_[int(2)] = d_3002_v108_
            nw587_[int(3)] = _dafny.SeqWithoutIsStrInference([d_2955_v78_ for d_3008_i6_ in range(default__.abs(652))])
            nw587_[int(4)] = _dafny.SeqWithoutIsStrInference([d_2955_v78_ for d_3009_i7_ in range(default__.abs(324))])
            nw587_[int(5)] = d_3002_v108_
            nw587_[int(6)] = ((d_3002_v108_).set(default__.safeIndex(d_2955_v78_, len(d_3002_v108_)), d_2955_v78_)) + (default__.fm31(d_2955_v78_, 763, globalState))
            nw587_[int(7)] = d_3002_v108_
            nw587_[int(8)] = d_3002_v108_
            nw587_[int(9)] = (d_3002_v108_) + (d_3002_v108_)
            nw587_[int(10)] = d_3002_v108_
            nw587_[int(11)] = d_3002_v108_
            nw587_[int(12)] = _dafny.SeqWithoutIsStrInference([len(d_2956_v79_) for d_3010_i8_ in range(default__.abs(47))])
            nw587_[int(13)] = (_dafny.SeqWithoutIsStrInference([((d_3005_v109_)[d_2864_v1_] if (d_2864_v1_) in (d_3005_v109_) else d_2955_v78_)])) + (d_3002_v108_)
            nw587_[int(14)] = (d_3002_v108_) + (d_3002_v108_)
            nw587_[int(15)] = d_3002_v108_
            nw587_[int(16)] = d_3002_v108_
            nw587_[int(17)] = d_3002_v108_
            nw587_[int(18)] = d_3002_v108_
            nw587_[int(19)] = ((d_3002_v108_).set(default__.safeIndex(len(((d_3006_v110_)[d_2955_v78_] if (d_2955_v78_) in (d_3006_v110_) else d_3002_v108_)), len(d_3002_v108_)), (self).f7)) + (d_3002_v108_)
            nw587_[int(20)] = d_3002_v108_
            nw587_[int(21)] = d_3002_v108_
            nw587_[int(22)] = d_3002_v108_
            nw587_[int(23)] = d_3002_v108_
            nw587_[int(24)] = _dafny.SeqWithoutIsStrInference([(self).f7])
            nw587_[int(25)] = (d_3002_v108_) + (d_3002_v108_)
            nw587_[int(26)] = d_3002_v108_
            nw587_[int(27)] = _dafny.SeqWithoutIsStrInference([(0) - (d_2955_v78_) for d_3011_i9_ in range(default__.abs(900))])
            nw587_[int(28)] = _dafny.SeqWithoutIsStrInference([(self).fm6(len(d_2957_v80_), 148, globalState)])
            d_3007_v111_ = nw587_
            index589_ = default__.safeIndex(225, (d_3007_v111_).length(0))
            (d_3007_v111_)[index589_] = _dafny.SeqWithoutIsStrInference([d_2955_v78_ for d_3012_i10_ in range(default__.abs(904))])
            index590_ = default__.safeIndex(225, (d_3007_v111_).length(0))
            (d_3007_v111_)[index590_] = d_3002_v108_
            d_2955_v78_ = (self).f7
            d_3013_v112_: _dafny.Array
            nw588_ = _dafny.Array(int(0), 27)
            d_3013_v112_ = nw588_
            r0 = (d_3013_v112_ if d_2864_v1_ else d_3013_v112_)
            d_2864_v1_ = d_2864_v1_
        d_2955_v78_ = (self).f7
        d_3014_v113_: _dafny.Array
        nw589_ = _dafny.Array(int(0), 20)
        d_3014_v113_ = nw589_
        r0 = d_3014_v113_
        return r0

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_3015_v0_: bool
        d_3015_v0_ = False
        d_3016_v1_: _dafny.Set
        d_3016_v1_ = _dafny.Set({d_3015_v0_, d_3015_v0_, not(d_3015_v0_), d_3015_v0_})
        d_3017_v2_: _dafny.Seq
        d_3017_v2_ = _dafny.SeqWithoutIsStrInference([d_3015_v0_])
        d_3018_v3_: _dafny.Array
        nw590_ = _dafny.Array(None, 13)
        nw590_[int(0)] = _dafny.Set({d_3015_v0_, d_3015_v0_, d_3015_v0_})
        nw590_[int(1)] = d_3016_v1_
        nw590_[int(2)] = d_3016_v1_
        nw590_[int(3)] = d_3016_v1_
        nw590_[int(4)] = d_3016_v1_
        nw590_[int(5)] = (d_3016_v1_).intersection(d_3016_v1_)
        nw590_[int(6)] = d_3016_v1_
        nw590_[int(7)] = (d_3016_v1_) - (d_3016_v1_)
        nw590_[int(8)] = (_dafny.Set({d_3015_v0_, d_3015_v0_, d_3015_v0_})).intersection(d_3016_v1_)
        nw590_[int(9)] = _dafny.Set({d_3015_v0_, d_3015_v0_, (d_3017_v2_)[default__.safeIndex((self).f7, len(d_3017_v2_))]})
        nw590_[int(10)] = (d_3016_v1_) | (d_3016_v1_)
        nw590_[int(11)] = (d_3016_v1_) - (d_3016_v1_)
        nw590_[int(12)] = _dafny.Set({default__.fm1((self).f7, (self).f7, d_3015_v0_, default__.fm1(len(d_3016_v1_), 483, d_3015_v0_, d_3015_v0_, globalState), globalState), d_3015_v0_, d_3015_v0_, False, d_3015_v0_})
        d_3018_v3_ = nw590_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_3018_v3_).length(0)):
            d_3019_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_3019_i0_)) and ((d_3019_i0_) < ((d_3018_v3_).length(0)))):
                (d_3018_v3_)[(d_3019_i0_)] = d_3016_v1_
        hi21_ = p0
        for d_3020_i1_ in range((self).f7, hi21_):
            if d_3015_v0_:
                d_3021_v4_: int
                d_3021_v4_ = 693
                d_3021_v4_ = 967
                d_3015_v0_ = d_3015_v0_
                d_3022_v5_: _dafny.Seq
                d_3022_v5_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_3023_v6_: _dafny.MultiSet
                d_3023_v6_ = _dafny.MultiSet([(self).f7])
                d_3021_v4_ = (((_dafny.MultiSet([(d_3022_v5_)[default__.safeIndex(len(_dafny.Set({True})), len(d_3022_v5_))], d_3021_v4_, ((d_3023_v6_)[d_3021_v4_] if (d_3021_v4_) in (d_3023_v6_) else d_3020_i1_), 745])).intersection(d_3023_v6_)).set((self).f7, default__.abs(979))).cardinality
                d_3024_v7_: D2
                d_3024_v7_ = D2_DC4(d_3015_v0_)
                d_3024_v7_ = d_3024_v7_
                d_3016_v1_ = (d_3016_v1_) - (_dafny.Set({False, d_3015_v0_}))
            elif True:
                d_3025_v8_: D15
                d_3025_v8_ = D15_DC35(p0)
                d_3026_v9_: int
                d_3026_v9_ = 405
                d_3027_v10_: _dafny.MultiSet
                d_3027_v10_ = _dafny.MultiSet([default__.fm15(d_3015_v0_, d_3015_v0_, d_3020_i1_, globalState)])
                d_3028_v11_: _dafny.Seq
                d_3028_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ceearghr"))
                rhs593_ = d_3025_v8_
                rhs594_ = ((d_3027_v10_)[_dafny.MultiSet([d_3015_v0_, d_3015_v0_, False])] if (_dafny.MultiSet([d_3015_v0_, d_3015_v0_, False])) in (d_3027_v10_) else len(d_3028_v11_))
                d_3025_v8_ = rhs593_
                d_3026_v9_ = rhs594_
                d_3029_v12_: C7
                nw591_ = C7()
                nw591_.ctor__()
                d_3029_v12_ = nw591_
                d_3026_v9_ = p0
                d_3026_v9_ = (0) - (d_3026_v9_)
                d_3026_v9_ = (p0) - ((p0) - ((self).f7))
            r0 = d_3015_v0_
            d_3030_v13_: int
            d_3030_v13_ = 607
            d_3030_v13_ = (self).f7
            d_3030_v13_ = ((54) + (-703)) * (293)
        d_3015_v0_ = not((d_3015_v0_) == (d_3015_v0_))
        d_3031_v14_: _dafny.Array
        nw592_ = _dafny.Array(None, 7)
        nw592_[int(0)] = d_3015_v0_
        nw592_[int(1)] = d_3015_v0_
        nw592_[int(2)] = d_3015_v0_
        nw592_[int(3)] = d_3015_v0_
        nw592_[int(4)] = d_3015_v0_
        nw592_[int(5)] = d_3015_v0_
        nw592_[int(6)] = (d_3017_v2_)[default__.safeIndex(p0, len(d_3017_v2_))]
        d_3031_v14_ = nw592_
        d_3032_v15_: _dafny.Seq
        d_3032_v15_ = _dafny.SeqWithoutIsStrInference([d_3031_v14_, d_3031_v14_, d_3031_v14_, d_3031_v14_, d_3031_v14_])
        d_3033_v16_: _dafny.Map
        d_3033_v16_ = _dafny.Map({(d_3032_v15_)[default__.safeIndex(p0, len(d_3032_v15_))]: d_3015_v0_})
        d_3034_v17_: _dafny.Map
        d_3034_v17_ = _dafny.Map({p0: d_3015_v0_})
        d_3015_v0_ = ((d_3033_v16_)[d_3031_v14_] if (d_3031_v14_) in (d_3033_v16_) else ((d_3034_v17_)[default__.fm14(globalState)] if (default__.fm14(globalState)) in (d_3034_v17_) else d_3015_v0_))
        d_3035_v18_: _dafny.Array
        nw593_ = _dafny.Array(_dafny.Seq({}), 14)
        d_3035_v18_ = nw593_
        d_3036_v19_: _dafny.Seq
        d_3036_v19_ = _dafny.SeqWithoutIsStrInference([-904])
        index591_ = default__.safeIndex(959, (d_3035_v18_).length(0))
        (d_3035_v18_)[index591_] = (d_3036_v19_) + (d_3036_v19_)
        index592_ = default__.safeIndex(959, (d_3035_v18_).length(0))
        (d_3035_v18_)[index592_] = (_dafny.SeqWithoutIsStrInference([(self).f7 for d_3037_i2_ in range(default__.abs(230))])) + ((d_3036_v19_) + (d_3036_v19_))
        d_3038_v20_: _dafny.Seq
        d_3038_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcjvago"))
        r2 = d_3038_v20_
        r0 = d_3015_v0_
        r1 = d_3015_v0_
        r2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkqgfndn"))
        return r0, r1, r2

    @property
    def f7(self):
        return self._f7
    @property
    def f6(self):
        return self._f6
