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
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm1(p0, p1, globalState):
        return not(not(not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_0_i0_ in range(default__.abs(750))])) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhvt"))))))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False])])), 242])) + (_dafny.SeqWithoutIsStrInference([(0) - (-242)]))) + ((_dafny.SeqWithoutIsStrInference([-647, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), 578, 227]) if True else _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1_i0_ in range(default__.abs(-573))])])).cardinality, (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([395, -527])) for d_2_i1_ in range(default__.abs(571))]))))])))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        if False:
            return (_dafny.Set({_dafny.Set({not(False)})})).intersection(_dafny.Set({_dafny.Set({(D10_DC29(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})), -791])), False)).cf57}), _dafny.Set({False}), _dafny.Set({True, True}), _dafny.Set({False}), _dafny.Set({True})}))
        elif True:
            return _dafny.Set({_dafny.Set({False, False}), _dafny.Set({False})})

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({True, True, True}) if True else _dafny.Set({False}))) - ((_dafny.Set({True, True, True, False})).intersection(_dafny.Set({False})))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return D2_DC6((0) - ((0) - (default__.safeDivisionInt(207, -76))), (_dafny.MultiSet([_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Map({678: _dafny.SeqWithoutIsStrInference([True])}))})), 590]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tx")))]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-348 for d_3_i0_ in range(default__.abs(25))]))])).issubset(_dafny.MultiSet([_dafny.MultiSet([732])])), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})) for d_4_i1_ in range(default__.abs(472))])), False, (0) - (len(_dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference([False, False]))}))))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        if (not(False)) or (not(not(False))):
            if False:
                return D3_DC8(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), _dafny.CodePoint('h')])), not(True), (0) - (len(_dafny.Map({-603: False}))))
            elif True:
                return D3_DC8(_dafny.MultiSet([_dafny.CodePoint('o')]), False, 251)
        elif True:
            return D3_DC8(_dafny.MultiSet([_dafny.CodePoint('v')]), not(False), (_dafny.MultiSet([False])).cardinality)

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('u')

    @staticmethod
    def fm15(p0, p1, globalState):
        source0_ = D4_DC10(_dafny.MultiSet([(0) - (-492)]))
        if source0_.is_DC11:
            d_5___mcc_h0_ = source0_.cf23
            d_6___mcc_h1_ = source0_.cf24
            d_7___mcc_h2_ = source0_.cf25
            d_8_cf25_ = d_7___mcc_h2_
            d_9_cf24_ = d_6___mcc_h1_
            d_10_cf23_ = d_5___mcc_h0_
            return D0_DC1(_dafny.MultiSet([d_9_cf24_]))
        elif True:
            d_11___mcc_h3_ = source0_.cf22
            d_12_cf22_ = d_11___mcc_h3_
            return D0_DC1(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm18(p0, p1, globalState):
        source1_ = D21_DC52(_dafny.CodePoint('k'))
        if source1_.is_DC52:
            d_13___mcc_h0_ = source1_.cf89
            d_14_cf89_ = d_13___mcc_h0_
            return D0_DC0(True)
        elif source1_.is_DC51:
            d_15___mcc_h1_ = source1_.cf88
            d_16_cf88_ = d_15___mcc_h1_
            return D0_DC0(False)
        elif True:
            d_17___mcc_h2_ = source1_.cf90
            d_18_cf90_ = d_17___mcc_h2_
            if not(False):
                return D0_DC0(not(False))
            elif True:
                return D0_DC0(True)

    @staticmethod
    def fm21(p0, p1, globalState):
        if (1) != ((0) - (-669)):
            return _dafny.CodePoint('h')
        elif True:
            return _dafny.CodePoint('x')

    @staticmethod
    def fm23(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D4_DC10(_dafny.MultiSet([656])), D4_DC10(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_19_i0_ in range(default__.abs(-781))])), 413]))]))])

    @staticmethod
    def fm24(p0, p1, globalState):
        source2_ = D20_DC49(False, _dafny.CodePoint('w'))
        if source2_.is_DC48:
            d_20___mcc_h0_ = source2_.cf82
            d_21___mcc_h1_ = source2_.cf83
            d_22___mcc_h2_ = source2_.cf84
            d_23_cf84_ = d_22___mcc_h2_
            d_24_cf83_ = d_21___mcc_h1_
            d_25_cf82_ = d_20___mcc_h0_
            return (470) - (613)
        elif source2_.is_DC49:
            d_26___mcc_h3_ = source2_.cf85
            d_27___mcc_h4_ = source2_.cf86
            d_28_cf86_ = d_27___mcc_h4_
            d_29_cf85_ = d_26___mcc_h3_
            return default__.safeDivisionInt(len(_dafny.Set({(_dafny.MultiSet([d_29_cf85_])).cardinality, -440})), 609)
        elif source2_.is_DC47:
            d_30___mcc_h5_ = source2_.cf81
            d_31_cf81_ = d_30___mcc_h5_
            return 132
        elif True:
            d_32___mcc_h6_ = source2_.cf87
            d_33_cf87_ = d_32___mcc_h6_
            return 626

    @staticmethod
    def fm25(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(550, -304):
                d_34_v0_: int = compr_0_
                if ((550) <= (d_34_v0_)) and ((d_34_v0_) < (-304)):
                    coll0_[default__.safeModuloInt(d_34_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_35_i0_ in range(default__.abs(-980))])))] = _dafny.Map({False: False})
            return _dafny.Map(coll0_)
        return (_dafny.MultiSet([len(_dafny.Map({(0) - (len(_dafny.Set({_dafny.CodePoint('u'), _dafny.CodePoint('u')}))): len(_dafny.Map({False: -198}))})), 496])) | ((_dafny.MultiSet([len(_dafny.Set({len(iife0_()
        ), len(_dafny.SeqWithoutIsStrInference([165, 557]))})), len(_dafny.Map({True: not(True)}))])) - (_dafny.MultiSet([297, (0) - (len(_dafny.Set({len(_dafny.Map({68: True}))})))])))

    @staticmethod
    def fm26(p0, globalState):
        return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([True])})) | ((D26_DC62(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False, True, False, not(True)])}))).cf100)

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Map
            for compr_1_ in (_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): not(True)}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): False})})).Elements:
                d_36_v0_: _dafny.Map = compr_1_
                if (d_36_v0_) in (_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): not(True)}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): False})})):
                    coll1_[d_36_v0_] = _dafny.MultiSet([_dafny.CodePoint('p')])
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Map
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({727: False}) for d_37_i0_ in range(default__.abs(484))])).Elements:
                d_38_v1_: _dafny.Map = compr_2_
                if (d_38_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({727: False}) for d_37_i0_ in range(default__.abs(484))])):
                    coll2_[d_38_v1_] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('o')]))
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Map
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({107: False}), _dafny.Map({-709: (D10_DC29(433, True)).cf57})])).Elements:
                d_39_v2_: _dafny.Map = compr_3_
                if (d_39_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({107: False}), _dafny.Map({-709: (D10_DC29(433, True)).cf57})])):
                    coll3_[d_39_v2_] = _dafny.MultiSet([_dafny.CodePoint('a'), _dafny.CodePoint('r')])
            return _dafny.Map(coll3_)
        return (iife1_()
        ) | ((iife2_()
        ) | (iife3_()
        ))

    @staticmethod
    def fm32(p0, globalState):
        return (0) - (default__.safeDivisionInt(323, default__.safeDivisionInt((_dafny.MultiSet([826])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvfvdpjmg"))))))

    @staticmethod
    def fm33(p0, globalState):
        return (_dafny.MultiSet([True, False, True, True])) - (_dafny.MultiSet([False]))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        source3_ = D8_DC22((0) - (-494), len(_dafny.SeqWithoutIsStrInference([772])))
        if source3_.is_DC22:
            d_40___mcc_h0_ = source3_.cf41
            d_41___mcc_h1_ = source3_.cf42
            d_42_cf42_ = d_41___mcc_h1_
            d_43_cf41_ = d_40___mcc_h0_
            return _dafny.CodePoint('w')
        elif source3_.is_DC23:
            d_44___mcc_h2_ = source3_.cf43
            d_45___mcc_h3_ = source3_.cf44
            d_46___mcc_h4_ = source3_.cf45
            d_47_cf45_ = d_46___mcc_h4_
            d_48_cf44_ = d_45___mcc_h3_
            d_49_cf43_ = d_44___mcc_h2_
            if True:
                return _dafny.CodePoint('n')
            elif True:
                return _dafny.CodePoint('j')
        elif source3_.is_DC24:
            d_50___mcc_h5_ = source3_.cf46
            d_51___mcc_h6_ = source3_.cf47
            d_52___mcc_h7_ = source3_.cf48
            d_53___mcc_h8_ = source3_.cf49
            d_54___mcc_h9_ = source3_.cf50
            d_55_cf50_ = d_54___mcc_h9_
            d_56_cf49_ = d_53___mcc_h8_
            d_57_cf48_ = d_52___mcc_h7_
            d_58_cf47_ = d_51___mcc_h6_
            d_59_cf46_ = d_50___mcc_h5_
            return _dafny.CodePoint('r')
        elif source3_.is_DC21:
            d_60___mcc_h10_ = source3_.cf40
            d_61_cf40_ = d_60___mcc_h10_
            if True:
                return _dafny.CodePoint('p')
            elif True:
                return _dafny.CodePoint('a')
        elif True:
            d_62___mcc_h11_ = source3_.cf51
            d_63_cf51_ = d_62___mcc_h11_
            return _dafny.CodePoint('l')

    @staticmethod
    def fm35(p0, globalState):
        source4_ = D8_DC24(True, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_64_i0_ in range(default__.abs(-792))])), _dafny.SeqWithoutIsStrInference([True]), False, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({True: len(_dafny.Map({False: True}))}): 905}))]))
        if source4_.is_DC22:
            d_65___mcc_h0_ = source4_.cf41
            d_66___mcc_h1_ = source4_.cf42
            d_67_cf42_ = d_66___mcc_h1_
            d_68_cf41_ = d_65___mcc_h0_
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: D0
                for compr_4_ in (_dafny.SeqWithoutIsStrInference([D0_DC1(_dafny.MultiSet([True])) for d_69_i1_ in range(default__.abs(682))])).Elements:
                    d_70_v0_: D0 = compr_4_
                    if (d_70_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC1(_dafny.MultiSet([True])) for d_69_i1_ in range(default__.abs(682))])):
                        coll4_[d_70_v0_] = False
                return _dafny.Map(coll4_)
            return ((D8_DC21(iife4_()
)).cf40) | (_dafny.Map({D0_DC1(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))): False}))
        elif source4_.is_DC23:
            d_71___mcc_h2_ = source4_.cf43
            d_72___mcc_h3_ = source4_.cf44
            d_73___mcc_h4_ = source4_.cf45
            d_74_cf45_ = d_73___mcc_h4_
            d_75_cf44_ = d_72___mcc_h3_
            d_76_cf43_ = d_71___mcc_h2_
            return _dafny.Map({D0_DC1(_dafny.MultiSet([False])): False})
        elif source4_.is_DC24:
            d_77___mcc_h5_ = source4_.cf46
            d_78___mcc_h6_ = source4_.cf47
            d_79___mcc_h7_ = source4_.cf48
            d_80___mcc_h8_ = source4_.cf49
            d_81___mcc_h9_ = source4_.cf50
            d_82_cf50_ = d_81___mcc_h9_
            d_83_cf49_ = d_80___mcc_h8_
            d_84_cf48_ = d_79___mcc_h7_
            d_85_cf47_ = d_78___mcc_h6_
            d_86_cf46_ = d_77___mcc_h5_
            if d_86_cf46_:
                return _dafny.Map({D0_DC1(_dafny.MultiSet([d_83_cf49_, d_83_cf49_])): d_86_cf46_})
            elif True:
                return _dafny.Map({D0_DC1(_dafny.MultiSet(d_84_cf48_)): d_83_cf49_})
        elif source4_.is_DC21:
            d_87___mcc_h10_ = source4_.cf40
            d_88_cf40_ = d_87___mcc_h10_
            return d_88_cf40_
        elif True:
            d_89___mcc_h11_ = source4_.cf51
            d_90_cf51_ = d_89___mcc_h11_
            return (_dafny.Map({D0_DC1(_dafny.MultiSet([True])): False})) | (_dafny.Map({D0_DC1(_dafny.MultiSet([False])): True}))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return D0_DC1(_dafny.MultiSet([False]))

    @staticmethod
    def fm37(p0, p1, globalState):
        return (_dafny.Map({False: False}))

    @staticmethod
    def fm38(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsk")))])).Elements:
                d_91_v0_: int = compr_5_
                if (d_91_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsk")))])):
                    coll5_[default__.safeDivisionInt(d_91_v0_, -244)] = (0) - ((0) - (-729))
            return _dafny.Map(coll5_)
        return (iife5_()
        ) | (_dafny.Map({312: 669}))

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        return len(_dafny.Map({D3_DC7(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: -880}), _dafny.Map({True: -879})])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('u'), _dafny.CodePoint('l')])}))

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        return _dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(True)), False, True, False]))) | (_dafny.MultiSet([False]))})

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('d')

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({988: len(_dafny.SeqWithoutIsStrInference([202 for d_92_i0_ in range(default__.abs(59))]))})))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, not(False), True}))])) + (_dafny.SeqWithoutIsStrInference([(0) - (-690) for d_93_i1_ in range(default__.abs(5))])))

    @staticmethod
    def fm43(globalState):
        source5_ = D8_DC25(D8_DC24(True, len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbqytxrv")): len(_dafny.Map({not(False): True}))})), _dafny.SeqWithoutIsStrInference([True]), not(True), _dafny.SeqWithoutIsStrInference([-972 for d_94_i0_ in range(default__.abs(549))])))
        if source5_.is_DC22:
            d_95___mcc_h0_ = source5_.cf41
            d_96___mcc_h1_ = source5_.cf42
            d_97_cf42_ = d_96___mcc_h1_
            d_98_cf41_ = d_95___mcc_h0_
            return _dafny.MultiSet([-766, 342, d_98_cf41_, d_97_cf42_])
        elif source5_.is_DC23:
            d_99___mcc_h2_ = source5_.cf43
            d_100___mcc_h3_ = source5_.cf44
            d_101___mcc_h4_ = source5_.cf45
            d_102_cf45_ = d_101___mcc_h4_
            d_103_cf44_ = d_100___mcc_h3_
            d_104_cf43_ = d_99___mcc_h2_
            return (_dafny.MultiSet([(0) - ((0) - ((0) - (-246)))])) | (_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_105_i1_ in range(default__.abs(427))])))]))
        elif source5_.is_DC24:
            d_106___mcc_h5_ = source5_.cf46
            d_107___mcc_h6_ = source5_.cf47
            d_108___mcc_h7_ = source5_.cf48
            d_109___mcc_h8_ = source5_.cf49
            d_110___mcc_h9_ = source5_.cf50
            d_111_cf50_ = d_110___mcc_h9_
            d_112_cf49_ = d_109___mcc_h8_
            d_113_cf48_ = d_108___mcc_h7_
            d_114_cf47_ = d_107___mcc_h6_
            d_115_cf46_ = d_106___mcc_h5_
            return _dafny.MultiSet([d_114_cf47_, d_114_cf47_, d_114_cf47_, d_114_cf47_])
        elif source5_.is_DC21:
            d_116___mcc_h10_ = source5_.cf40
            d_117_cf40_ = d_116___mcc_h10_
            return (_dafny.MultiSet([832, len(_dafny.Set({not(False)})), 789, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuonpos"))): False}))])) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_118_i2_ in range(default__.abs(351))]))]))
        elif True:
            d_119___mcc_h11_ = source5_.cf51
            d_120_cf51_ = d_119___mcc_h11_
            return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "schrvre"))), 887])) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrovpuj"))), (0) - (len(_dafny.SeqWithoutIsStrInference([True, False])))]))

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return _dafny.Set({(len(_dafny.SeqWithoutIsStrInference([True]))) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))))})

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        return len((((D16_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llk")), 894)).cf73) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hf")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_121_i0_ in range(default__.abs(401))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lacfjew")))))

    @staticmethod
    def fm47(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(-128, 591):
                d_123_v0_: int = compr_6_
                if ((-128) <= (d_123_v0_)) and ((d_123_v0_) < (591)):
                    coll6_[(d_123_v0_) * (-966)] = True
            return _dafny.Map(coll6_)
        return (_dafny.Map({len(_dafny.Map({668: False})): _dafny.SeqWithoutIsStrInference([(D21_DC52(_dafny.CodePoint('j'))).cf89 for d_122_i0_ in range(default__.abs(-920))])})) | (_dafny.Map({len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([not((D10_DC31(False, not(not(True)), len(_dafny.Map({_dafny.CodePoint('j'): _dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbpvswxo"))), 340])), len(_dafny.Set({True, True})), 481, len(iife6_()
), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnyamwon"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kd")))}))})})))).cf61), False])})): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_124_i1_ in range(default__.abs(396))])}))

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([5 for d_125_i0_ in range(default__.abs(937))]), _dafny.SeqWithoutIsStrInference([753]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_126_i2_ in range(default__.abs(-478))]))) for d_127_i1_ in range(default__.abs(371))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([111]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([916, 995]))])), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([908 for d_128_i3_ in range(default__.abs(230))]))]), _dafny.MultiSet([989]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([886]))]))])])) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([382]), _dafny.SeqWithoutIsStrInference([-56, 776]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality, 36])])):
            return _dafny.CodePoint('d')
        elif True:
            return _dafny.CodePoint('j')

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(not(True))]))]) if False else _dafny.SeqWithoutIsStrInference([220, len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxr")))): True}))])))) | ((_dafny.MultiSet([18, -111])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))), 478, 708, -3])))

    @staticmethod
    def fm50(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(112, 890):
                d_129_v0_: int = compr_7_
                if ((112) <= (d_129_v0_)) and ((d_129_v0_) < (890)):
                    coll7_[(d_129_v0_) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxkxsgkqe"))): len(_dafny.Set({False}))})))] = True
            return _dafny.Map(coll7_)
        return _dafny.Map({True: (_dafny.Map({len(iife7_()
        ): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpyifw")))}))})

    @staticmethod
    def fm51(p0, p1, globalState):
        if not(((0) - (-184)) != (len(_dafny.Map({284: False})))):
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, True, not((D26_DC63(-626, 673, True)).cf103)])])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]) for d_130_i0_ in range(default__.abs(397))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([False])]))

    @staticmethod
    def fm52(p0, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: _dafny.Seq
            for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True])])).Elements:
                d_131_v0_: _dafny.Seq = compr_8_
                if (d_131_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True])])):
                    coll8_ = coll8_.union(_dafny.Set([d_131_v0_]))
            return _dafny.Set(coll8_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False])})).intersection(iife8_()
        )) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([not(False)])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(False), False]), _dafny.SeqWithoutIsStrInference([True, False])})))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D10_DC31(not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvutgcy"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))), not (False) or (True), ((_dafny.MultiSet([False])).cardinality) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_132_i0_ in range(default__.abs(415))]))))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: _dafny.Seq
            for compr_9_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_134_i1_ in range(default__.abs(276))])), len(_dafny.Map({True: len(_dafny.Map({272: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_135_i2_ in range(default__.abs(639))]))}))}))])])).Elements:
                d_136_v0_: _dafny.Seq = compr_9_
                if (d_136_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_134_i1_ in range(default__.abs(276))])), len(_dafny.Map({True: len(_dafny.Map({272: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_135_i2_ in range(default__.abs(639))]))}))}))])])):
                    coll9_ = coll9_.union(_dafny.Set([d_136_v0_]))
            return _dafny.Set(coll9_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqw"))) for d_133_i0_ in range(default__.abs(746))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqg")))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])})) - (iife9_()
        )

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        source6_ = D10_DC29(100, True)
        if source6_.is_DC29:
            d_137___mcc_h0_ = source6_.cf56
            d_138___mcc_h1_ = source6_.cf57
            d_139_cf57_ = d_138___mcc_h1_
            d_140_cf56_ = d_137___mcc_h0_
            return _dafny.Set({False})
        elif source6_.is_DC30:
            d_141___mcc_h2_ = source6_.cf58
            d_142___mcc_h3_ = source6_.cf59
            d_143_cf59_ = d_142___mcc_h3_
            d_144_cf58_ = d_141___mcc_h2_
            return _dafny.Set({False})
        elif source6_.is_DC31:
            d_145___mcc_h4_ = source6_.cf60
            d_146___mcc_h5_ = source6_.cf61
            d_147___mcc_h6_ = source6_.cf62
            d_148_cf62_ = d_147___mcc_h6_
            d_149_cf61_ = d_146___mcc_h5_
            d_150_cf60_ = d_145___mcc_h4_
            if not(d_150_cf60_):
                return _dafny.Set({d_149_cf61_})
            elif True:
                return _dafny.Set({d_149_cf61_})
        elif source6_.is_DC28:
            d_151___mcc_h7_ = source6_.cf55
            d_152_cf55_ = d_151___mcc_h7_
            return _dafny.Set({(d_152_cf55_).f43, (d_152_cf55_).f43, (d_152_cf55_).f43, not((d_152_cf55_).f43), (d_152_cf55_).f43})
        elif True:
            d_153___mcc_h8_ = source6_.cf63
            d_154_cf63_ = d_153___mcc_h8_
            return _dafny.Set({False, not(True)})

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        return D2_DC6(default__.safeDivisionInt(875, -348), (True if not(True) else not(not(not(not(True))))), -364, not((55) > (len(_dafny.Map({232: len(_dafny.Map({True: len(_dafny.Map({False: False}))}))})))), 667)

    @staticmethod
    def fm57(p0, p1, globalState):
        return (_dafny.Map({833: False})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcbb"))): False}))

    @staticmethod
    def fm58(p0, p1, globalState):
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: str
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])).Elements:
                d_155_v0_: str = compr_10_
                if (d_155_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])):
                    coll10_ = coll10_.union(_dafny.Set([d_155_v0_]))
            return _dafny.Set(coll10_)
        return D14_DC37(default__.safeModuloInt((D26_DC63(len(_dafny.Map({not(True): len(iife10_()
)})), len(_dafny.Set({70})), True)).cf101, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_156_i0_ in range(default__.abs(323))]))), True, default__.safeDivisionInt(723, len(_dafny.Map({-464: len(_dafny.Map({_dafny.CodePoint('s'): (D8_DC22(len(_dafny.SeqWithoutIsStrInference([-737])), len(_dafny.Map({False: False})))).cf41}))}))))

    @staticmethod
    def fm59(p0, p1, p2, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_157_i0_ in range(default__.abs(991))])).Elements:
                d_158_v0_: int = compr_11_
                if (d_158_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_157_i0_ in range(default__.abs(991))])):
                    coll11_[default__.safeModuloInt(d_158_v0_, len(_dafny.Set({False})))] = _dafny.Set({326})
            return _dafny.Map(coll11_)
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekgojohii"))), 374, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False])), 185, 958]))])).cardinality, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_159_i1_ in range(default__.abs(976))]))})), 248, 565])).Elements:
                d_160_v1_: int = compr_12_
                if (d_160_v1_) in (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekgojohii"))), 374, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False])), 185, 958]))])).cardinality, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_159_i1_ in range(default__.abs(976))]))})), 248, 565])):
                    def iife13_():
                        coll13_ = _dafny.Set()
                        compr_13_: int
                        for compr_13_ in _dafny.IntegerRange(901, 58):
                            d_161_v2_: int = compr_13_
                            if ((901) <= (d_161_v2_)) and ((d_161_v2_) < (58)):
                                coll13_ = coll13_.union(_dafny.Set([(d_161_v2_) + ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_162_i2_ in range(default__.abs(789))])), len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))).cardinality: -470}))])).cardinality)]))
                        return _dafny.Set(coll13_)
                    coll12_[(d_160_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrr"))))] = iife13_()

            return _dafny.Map(coll12_)
        return ((iife11_()
        ) | (_dafny.Map({-574: _dafny.Set({88})}))) | (iife12_()
        )

    @staticmethod
    def fm60(globalState):
        return ((_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.Set({len(_dafny.Map({True: 110}))}))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcr"))), len(_dafny.Set({(D22_DC55(len(_dafny.Set({_dafny.CodePoint('p'), _dafny.CodePoint('q'), _dafny.CodePoint('l'), _dafny.CodePoint('l')})), 372, True)).cf93})), -613])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([543 for d_163_i0_ in range(default__.abs(342))])}))) | ((_dafny.Map({False: _dafny.SeqWithoutIsStrInference([(0) - (-801) for d_164_i1_ in range(default__.abs(746))])})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgo"))) for d_165_i2_ in range(default__.abs(123))])})))

    @staticmethod
    def fm61(p0, p1, p2, globalState):
        source7_ = D16_DC41(D16_DC39(_dafny.Set({_dafny.SeqWithoutIsStrInference([707 for d_166_i0_ in range(default__.abs(522))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Set({False})): False})), 395, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_167_i1_ in range(default__.abs(-512))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkgia"))), 806]), _dafny.SeqWithoutIsStrInference([124])})))
        if source7_.is_DC40:
            d_168___mcc_h0_ = source7_.cf73
            d_169___mcc_h1_ = source7_.cf74
            d_170_cf74_ = d_169___mcc_h1_
            d_171_cf73_ = d_168___mcc_h0_
            return D8_DC21(_dafny.Map({D0_DC1(_dafny.MultiSet([False])): False}))
        elif source7_.is_DC39:
            d_172___mcc_h2_ = source7_.cf72
            d_173_cf72_ = d_172___mcc_h2_
            return D8_DC21(_dafny.Map({D0_DC1(_dafny.MultiSet([not(False), not(False)])): not(True)}))
        elif True:
            d_174___mcc_h3_ = source7_.cf75
            d_175_cf75_ = d_174___mcc_h3_
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: D0
                for compr_14_ in (_dafny.Map({D0_DC1(_dafny.MultiSet([not(False)])): 618})).keys.Elements:
                    d_176_v0_: D0 = compr_14_
                    if (d_176_v0_) in (_dafny.Map({D0_DC1(_dafny.MultiSet([not(False)])): 618})):
                        coll14_[d_176_v0_] = True
                return _dafny.Map(coll14_)
            return D8_DC21(iife14_()
)

    @staticmethod
    def fm62(p0, p1, p2, p3, globalState):
        return D21_DC51(_dafny.Map({285: _dafny.Set({len(_dafny.SeqWithoutIsStrInference([902]))})}))

    @staticmethod
    def fm63(p0, p1, p2, globalState):
        return _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_177_i0_ in range(default__.abs(687))]))), (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eix")) for d_178_i1_ in range(default__.abs(385))]))))})

    @staticmethod
    def fm64(p0, p1, p2, p3, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in (_dafny.Map({279: (D16_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stc")), len(_dafny.Map({False: True})))).cf73})).keys.Elements:
                d_179_v0_: int = compr_15_
                if (d_179_v0_) in (_dafny.Map({279: (D16_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stc")), len(_dafny.Map({False: True})))).cf73})):
                    coll15_[(d_179_v0_) + (len(_dafny.Map({True: False})))] = D0_DC1(_dafny.MultiSet([False, True]))
            return _dafny.Map(coll15_)
        return iife15_()
        

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        return ((_dafny.Map({True: D8_DC22(-188, 0)}) if not(True) else _dafny.Map({not(not(False)): D8_DC22(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))), 398)}))) | ((D29_DC67(_dafny.Map({True: D8_DC22((0) - (len(_dafny.Map({False: True}))), 792)}))).cf107)

    @staticmethod
    def fm66(globalState):
        return D6_DC15((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fflctru")))])) + (_dafny.SeqWithoutIsStrInference([137])))

    @staticmethod
    def fm67(globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}), _dafny.Map({True: True})])) <= (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}), _dafny.Map({not(not(not(True))): False})])): _dafny.Map({True: not(False)})})

    @staticmethod
    def fm68(p0, p1, p2, globalState):
        return D20_DC49((_dafny.SeqWithoutIsStrInference([not(False)])) == (_dafny.SeqWithoutIsStrInference([False, False, True, False])), _dafny.CodePoint('q'))

    @staticmethod
    def fm69(p0, p1, p2, p3, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(-784, 969):
                d_180_v0_: int = compr_16_
                if ((-784) <= (d_180_v0_)) and ((d_180_v0_) < (969)):
                    coll16_[(d_180_v0_) - (len(_dafny.Set({689, 252})))] = False
            return _dafny.Map(coll16_)
        return (_dafny.Map({D16_DC39(_dafny.Set({_dafny.SeqWithoutIsStrInference([814]), _dafny.SeqWithoutIsStrInference([920])})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrtq"))})) | (_dafny.Map({D16_DC39(_dafny.Set({(D8_DC24(False, len(iife16_()
), _dafny.SeqWithoutIsStrInference([False]), not(not(True)), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))]))).cf50, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neweldmjf"))): 783})) for d_181_i0_ in range(default__.abs(141))]))])})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbhwx"))}))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_182_v0_: _dafny.Map
        d_182_v0_ = _dafny.Map({555: default__.safeModuloInt(p2, p2)})
        d_182_v0_ = d_182_v0_
        d_183_v1_: _dafny.Map
        d_183_v1_ = _dafny.Map({p2: p3})
        d_184_v2_: str
        d_184_v2_ = _dafny.CodePoint('y')
        d_185_v3_: D20
        d_185_v3_ = D20_DC49(p3, d_184_v2_)
        d_186_v4_: _dafny.MultiSet
        d_186_v4_ = _dafny.MultiSet([d_185_v3_])
        d_187_v5_: _dafny.MultiSet
        d_187_v5_ = _dafny.MultiSet([p1])
        d_188_v6_: _dafny.Set
        d_188_v6_ = _dafny.Set({((d_187_v5_)[p3] if (p3) in (d_187_v5_) else p2)})
        hi0_ = 312
        for d_189_i0_ in range(((0) - (len(d_183_v1_))) * (default__.fm46(((d_186_v4_)[default__.fm68(p3, p1, p1, globalState)] if (default__.fm68(p3, p1, p1, globalState)) in (d_186_v4_) else p2), len(d_188_v6_), p1, globalState)), hi0_):
            d_190_v7_: _dafny.Array
            def lambda0_(d_191_p2_, d_192_p3_, d_193_p1_):
                def lambda1_(d_194_i1_):
                    return D2_DC6((0) - (d_191_p2_), d_192_p3_, 597, d_193_p1_, d_191_p2_)

                return lambda1_

            init0_ = lambda0_(p2, p3, p1)
            nw0_ = _dafny.Array(None, 15)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_190_v7_ = nw0_
            d_195_v8_: _dafny.MultiSet
            d_195_v8_ = _dafny.MultiSet([d_190_v7_, d_190_v7_])
            d_196_v9_: _dafny.Seq
            d_196_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuxvyifu"))
            (globalState).f26 = (_dafny.MultiSet([d_190_v7_])).isdisjoint((d_195_v8_).set(d_190_v7_, default__.abs(len(d_196_v9_))))
            d_197_v10_: _dafny.MultiSet
            d_197_v10_ = _dafny.MultiSet([d_189_i0_])
            d_198_v11_: _dafny.Array
            nw1_ = _dafny.Array(None, 16)
            nw1_[int(0)] = 537
            nw1_[int(1)] = d_189_i0_
            nw1_[int(2)] = len(_dafny.Set({p2}))
            nw1_[int(3)] = (p2) + ((0) - (len(d_196_v9_)))
            nw1_[int(4)] = (d_189_i0_ if p3 else (d_197_v10_).cardinality)
            nw1_[int(5)] = p2
            nw1_[int(6)] = p2
            nw1_[int(7)] = p2
            nw1_[int(8)] = default__.fm39(p1, d_197_v10_, p3, globalState)
            nw1_[int(9)] = len(d_196_v9_)
            nw1_[int(10)] = p2
            nw1_[int(11)] = d_189_i0_
            nw1_[int(12)] = p2
            nw1_[int(13)] = (p2) * (p2)
            nw1_[int(14)] = len((d_182_v0_) | (d_182_v0_))
            nw1_[int(15)] = ((p0)[default__.safeIndex(d_189_i0_, len(p0))]) * (p2)
            d_198_v11_ = nw1_
            index0_ = default__.safeIndex(656, (d_198_v11_).length(0))
            (d_198_v11_)[index0_] = d_189_i0_
            index1_ = default__.safeIndex(656, (d_198_v11_).length(0))
            (d_198_v11_)[index1_] = d_189_i0_
            if p1:
                (globalState).f13 = p2
                d_199_v12_: _dafny.Map
                d_199_v12_ = _dafny.Map({p3: d_184_v2_})
                d_200_v13_: _dafny.Map
                d_200_v13_ = _dafny.Map({p1: False})
                d_201_v14_: _dafny.Map
                d_201_v14_ = _dafny.Map({len(d_199_v12_): d_200_v13_})
                d_202_v15_: _dafny.Seq
                d_202_v15_ = _dafny.SeqWithoutIsStrInference([p1])
                d_203_v16_: _dafny.Seq
                d_203_v16_ = _dafny.SeqWithoutIsStrInference([d_202_v15_])
                d_204_v17_: C9
                nw2_ = C9()
                nw2_.ctor__((d_187_v5_).cardinality, d_201_v14_, d_184_v2_, (len(d_196_v9_)) * (len(d_200_v13_)), p1, ((_dafny.SeqWithoutIsStrInference([p3])) + (d_202_v15_)).set(default__.safeIndex(708, len((_dafny.SeqWithoutIsStrInference([p3])) + (d_202_v15_))), p3), (d_203_v16_) + (d_203_v16_), d_189_i0_, ((d_198_v11_)[default__.safeIndex(656, (d_198_v11_).length(0))]) > (p2), p3, p3)
                d_204_v17_ = nw2_
                (globalState).f13 = default__.safeDivisionInt(len(d_202_v15_), ((d_204_v17_).f38) * (d_189_i0_))
                (globalState).f26 = p3
                d_205_v18_: _dafny.Array
                def lambda2_(d_206_p1_):
                    def lambda3_(d_207_i2_):
                        return d_206_p1_

                    return lambda3_

                init1_ = lambda2_(p1)
                nw3_ = _dafny.Array(None, 17)
                for i0_1_ in range(nw3_.length(0)):
                    nw3_[i0_1_] = init1_(i0_1_)
                d_205_v18_ = nw3_
                d_208_v19_: _dafny.Map
                d_208_v19_ = _dafny.Map({d_196_v9_: d_205_v18_})
                d_208_v19_ = (d_208_v19_).set((d_196_v9_) + ((d_196_v9_).set(default__.safeIndex(d_189_i0_, len(d_196_v9_)), d_184_v2_)), d_205_v18_)
            elif True:
                d_209_v20_: _dafny.Seq
                d_209_v20_ = _dafny.SeqWithoutIsStrInference([p3, p3])
                d_210_v21_: _dafny.Seq
                d_210_v21_ = _dafny.SeqWithoutIsStrInference([d_209_v20_])
                d_211_v22_: _dafny.Array
                nw4_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                d_211_v22_ = nw4_
                d_212_v23_: C0
                nw5_ = C0()
                nw5_.ctor__(d_211_v22_)
                d_212_v23_ = nw5_
                d_213_v24_: _dafny.Seq
                d_213_v24_ = _dafny.SeqWithoutIsStrInference([d_212_v23_])
                d_214_v25_: C6
                nw6_ = C6()
                nw6_.ctor__(d_184_v2_, ((d_182_v0_)[220] if (220) in (d_182_v0_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))), p3, _dafny.SeqWithoutIsStrInference([p1]), d_210_v21_, d_189_i0_, p1, (d_213_v24_) != (d_213_v24_), p3)
                d_214_v25_ = nw6_
                d_215_v26_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Map({}), 3)
                d_215_v26_ = nw7_
                d_216_v27_: _dafny.Array
                nw8_ = _dafny.Array(False, 16)
                d_216_v27_ = nw8_
                d_217_v28_: _dafny.Map
                d_217_v28_ = _dafny.Map({d_216_v27_: _dafny.MultiSet([p3, ((d_183_v1_)[d_189_i0_] if (d_189_i0_) in (d_183_v1_) else False)])})
                index2_ = default__.safeIndex(897, (d_215_v26_).length(0))
                (d_215_v26_)[index2_] = d_217_v28_
                index3_ = default__.safeIndex(897, (d_215_v26_).length(0))
                rhs0_ = d_217_v28_
                rhs1_ = d_189_i0_
                lhs0_ = d_215_v26_
                lhs1_ = default__.safeIndex(897, (d_215_v26_).length(0))
                lhs2_ = globalState
                lhs0_[lhs1_] = rhs0_
                lhs2_.f13 = rhs1_
                d_197_v10_ = (d_197_v10_).intersection((D4_DC10(d_197_v10_)).cf22)
                d_218_v29_: _dafny.Set
                d_218_v29_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_189_i0_]), p0, _dafny.SeqWithoutIsStrInference([(d_198_v11_)[default__.safeIndex(656, (d_198_v11_).length(0))], p2]), p0})
                d_219_v30_: D16
                d_219_v30_ = D16_DC39(d_218_v29_)
                d_220_v31_: _dafny.Map
                d_220_v31_ = _dafny.Map({d_219_v30_: d_196_v9_})
                d_221_v32_: D0
                d_221_v32_ = D0_DC0(p3)
                index4_ = default__.safeIndex(656, (d_198_v11_).length(0))
                rhs2_ = default__.fm69(p1, p2, len(_dafny.SeqWithoutIsStrInference([d_184_v2_ for d_222_i3_ in range(default__.abs(921))])), not((d_221_v32_).cf0), globalState)
                rhs3_ = (0) - (p2)
                lhs3_ = d_198_v11_
                lhs4_ = default__.safeIndex(656, (d_198_v11_).length(0))
                d_220_v31_ = rhs2_
                lhs3_[lhs4_] = rhs3_
                d_223_v33_: T2
                nw9_ = C8()
                nw9_.ctor__(p3, d_184_v2_, 689, p1, d_209_v20_, _dafny.SeqWithoutIsStrInference([d_209_v20_ for d_224_i4_ in range(default__.abs(418))]), p2, p3, True, p3)
                d_223_v33_ = nw9_
                d_225_v34_: _dafny.Seq
                d_225_v34_ = _dafny.SeqWithoutIsStrInference([d_223_v33_])
                d_226_v35_: _dafny.Seq
                d_226_v35_ = _dafny.SeqWithoutIsStrInference([p0])
                d_227_v36_: T3
                nw10_ = C5()
                nw10_.ctor__(p1, len(d_225_v34_), len(d_226_v35_), (d_214_v25_).fm6(globalState), d_209_v20_, d_210_v21_, (0) - (len(d_196_v9_)), d_223_v33_.f28, True, True)
                d_227_v36_ = nw10_
                d_228_v37_: _dafny.Array
                nw11_ = _dafny.Array(None, 3)
                nw11_[int(0)] = d_227_v36_
                nw11_[int(1)] = d_227_v36_
                nw11_[int(2)] = d_227_v36_
                d_228_v37_ = nw11_
                d_229_v38_: _dafny.MultiSet
                d_229_v38_ = _dafny.MultiSet([(d_198_v11_ if d_223_v33_.f28 else d_198_v11_), d_198_v11_, d_198_v11_, d_198_v11_, (d_198_v11_ if p3 else d_198_v11_)])
                rhs4_ = (d_229_v38_).cardinality
                rhs5_ = d_189_i0_
                rhs6_ = d_228_v37_
                rhs7_ = (d_227_v36_).fm6(globalState)
                rhs8_ = not (((d_198_v11_)[default__.safeIndex(656, (d_198_v11_).length(0))]) <= (d_189_i0_)) or ((d_209_v20_)[default__.safeIndex(len(_dafny.Set({d_223_v33_.f29})), len(d_209_v20_))])
                lhs5_ = globalState
                lhs6_ = globalState
                r0 = rhs4_
                r0 = rhs5_
                d_228_v37_ = rhs6_
                lhs5_.f26 = rhs7_
                lhs6_.f12 = rhs8_
            d_230_v39_: _dafny.Map
            d_230_v39_ = _dafny.Map({p1: p2})
            d_231_v40_: _dafny.Seq
            d_231_v40_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({default__.fm1(p1, len(d_196_v9_), globalState): p2})).set(p1, p2), (d_230_v39_) | (_dafny.Map({p1: d_189_i0_})), (d_230_v39_) | (d_230_v39_), d_230_v39_])
            d_231_v40_ = d_231_v40_
        d_232_v41_: _dafny.Set
        d_232_v41_ = _dafny.Set({_dafny.CodePoint('j')})
        d_233_v42_: _dafny.Seq
        d_233_v42_ = _dafny.SeqWithoutIsStrInference([p3, p1, True])
        d_234_v43_: _dafny.MultiSet
        d_234_v43_ = _dafny.MultiSet([d_232_v41_])
        d_235_v44_: _dafny.Seq
        d_235_v44_ = _dafny.SeqWithoutIsStrInference([d_233_v42_])
        d_236_v45_: C10
        nw12_ = C10()
        nw12_.ctor__(d_184_v2_, p2, p3, d_233_v42_, d_235_v44_, p2, (d_233_v42_)[default__.safeIndex(p2, len(d_233_v42_))], p1, p3)
        d_236_v45_ = nw12_
        d_237_v46_: T3
        nw13_ = C4()
        nw13_.ctor__(p2, False, _dafny.SeqWithoutIsStrInference([not(p3), p3]), d_235_v44_, p2, p1, p1, p3)
        d_237_v46_ = nw13_
        d_238_v47_: _dafny.Map
        d_238_v47_ = _dafny.Map({d_236_v45_: d_237_v46_})
        d_239_v48_: _dafny.Seq
        d_239_v48_ = _dafny.SeqWithoutIsStrInference([d_238_v47_, d_238_v47_, d_238_v47_, d_238_v47_, d_238_v47_])
        d_240_v49_: _dafny.MultiSet
        d_240_v49_ = _dafny.MultiSet([(d_239_v48_)[default__.safeIndex(len(p0), len(d_239_v48_))], d_238_v47_])
        d_241_v50_: _dafny.Array
        nw14_ = _dafny.Array(None, 23)
        nw14_[int(0)] = p1
        nw14_[int(1)] = (True if p1 else p3)
        nw14_[int(2)] = not (p1) or (p3)
        nw14_[int(3)] = (default__.fm39(not(p1), _dafny.MultiSet([p2, p2, p2, (0) - (p2), len(_dafny.SeqWithoutIsStrInference([d_184_v2_ for d_242_i5_ in range(default__.abs(574))]))]), p1, globalState)) in (d_183_v1_)
        nw14_[int(4)] = p1
        nw14_[int(5)] = p3
        nw14_[int(6)] = p3
        nw14_[int(7)] = p3
        nw14_[int(8)] = p3
        nw14_[int(9)] = (p3) and (p1)
        nw14_[int(10)] = (p3) == (p3)
        nw14_[int(11)] = p3
        nw14_[int(12)] = p3
        nw14_[int(13)] = True
        nw14_[int(14)] = p1
        nw14_[int(15)] = (d_184_v2_) in (d_232_v41_)
        nw14_[int(16)] = p1
        nw14_[int(17)] = ((d_183_v1_)[(0) - (len(d_233_v42_))] if ((0) - (len(d_233_v42_))) in (d_183_v1_) else p1)
        nw14_[int(18)] = (p2) < (p2)
        nw14_[int(19)] = not((d_234_v43_).isdisjoint(d_234_v43_))
        nw14_[int(20)] = p3
        nw14_[int(21)] = p1
        nw14_[int(22)] = (d_240_v49_) == (d_240_v49_)
        d_241_v50_ = nw14_
        index5_ = default__.safeIndex(311, (d_241_v50_).length(0))
        (d_241_v50_)[index5_] = d_237_v46_.f29
        d_243_v51_: _dafny.Seq
        d_243_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujnosv"))
        d_244_v52_: _dafny.MultiSet
        d_244_v52_ = _dafny.MultiSet([d_243_v51_])
        d_245_v53_: _dafny.Seq
        d_245_v53_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_243_v51_ for d_246_i6_ in range(default__.abs(518))])])
        d_247_v54_: _dafny.MultiSet
        d_247_v54_ = _dafny.MultiSet([p2])
        d_248_v55_: T2
        nw15_ = C6()
        nw15_.ctor__(d_184_v2_, (d_237_v46_).f30, p3, d_233_v42_, (d_237_v46_).f33, p2, d_237_v46_.f35, d_237_v46_.f35, p3)
        d_248_v55_ = nw15_
        d_249_v56_: _dafny.Map
        d_249_v56_ = _dafny.Map({len(d_243_v51_): d_248_v55_})
        d_250_v57_: _dafny.Seq
        d_250_v57_ = _dafny.SeqWithoutIsStrInference([((d_249_v56_)[330] if (330) in (d_249_v56_) else d_248_v55_)])
        d_251_v58_: _dafny.Map
        d_251_v58_ = _dafny.Map({d_184_v2_: len(d_250_v57_)})
        index6_ = default__.safeIndex(311, (d_241_v50_).length(0))
        rhs9_ = (d_237_v46_).fm6(globalState)
        rhs10_ = ((d_244_v52_).intersection(d_244_v52_)) | (_dafny.MultiSet((d_245_v53_)[default__.safeIndex(((d_247_v54_)[(0) - ((0) - (((d_251_v58_)[d_184_v2_] if (d_184_v2_) in (d_251_v58_) else (d_237_v46_).f30)))] if ((0) - ((0) - (((d_251_v58_)[d_184_v2_] if (d_184_v2_) in (d_251_v58_) else (d_237_v46_).f30)))) in (d_247_v54_) else len(d_243_v51_)), len(d_245_v53_))]))
        lhs7_ = d_241_v50_
        lhs8_ = default__.safeIndex(311, (d_241_v50_).length(0))
        lhs7_[lhs8_] = rhs9_
        d_244_v52_ = rhs10_
        r0 = (d_237_v46_).f30
        d_252_v59_: _dafny.Array
        nw16_ = _dafny.Array(int(0), 27)
        d_252_v59_ = nw16_
        d_252_v59_ = d_252_v59_
        hi1_ = (d_237_v46_).f30
        for d_253_i7_ in range((d_248_v55_).f30, hi1_):
            (d_237_v46_).f31 = False
            (globalState).f13 = ((d_237_v46_).f30) - ((d_237_v46_).f34)
            d_254_v60_: D0
            d_254_v60_ = D0_DC1(d_187_v5_)
            d_255_v61_: C10
            nw17_ = C10()
            nw17_.ctor__(default__.fm14(386, ((d_248_v55_).f32)[default__.safeIndex((d_237_v46_).f34, len((d_248_v55_).f32))], (d_248_v55_).f30, d_254_v60_, globalState), -339, p1, ((d_237_v46_).f32) + (d_233_v42_), ((d_237_v46_).f33) + (d_235_v44_), (d_237_v46_).f30, (d_241_v50_)[default__.safeIndex(311, (d_241_v50_).length(0))], d_237_v46_.f31, ((d_236_v45_).fm7(d_247_v54_, (0) - (-158), globalState) if True else d_248_v55_.f29))
            d_255_v61_ = nw17_
            d_256_v62_: T1
            nw18_ = C2()
            nw18_.ctor__((d_237_v46_).f34, (d_248_v55_).f30, d_248_v55_.f29, d_248_v55_.f29, False)
            d_256_v62_ = nw18_
            d_257_v63_: _dafny.Seq
            d_257_v63_ = _dafny.SeqWithoutIsStrInference([d_256_v62_, d_256_v62_])
            d_258_v64_: _dafny.Seq
            d_258_v64_ = (_dafny.SeqWithoutIsStrInference([d_256_v62_])) + (d_257_v63_)
            d_258_v64_ = d_258_v64_
        r0 = (0) - (default__.safeModuloInt(p2, len((d_243_v51_ if (d_241_v50_)[default__.safeIndex(311, (d_241_v50_).length(0))] else d_243_v51_))))
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_259_v0_: _dafny.Array
        nw19_ = _dafny.Array(int(0), 17)
        d_259_v0_ = nw19_
        d_260_v1_: int
        d_260_v1_ = 20
        d_261_v2_: _dafny.Seq
        d_261_v2_ = _dafny.SeqWithoutIsStrInference([d_260_v1_])
        d_262_v3_: _dafny.Map
        d_262_v3_ = _dafny.Map({d_259_v0_: d_261_v2_})
        d_263_v4_: bool
        d_263_v4_ = True
        d_264_v5_: _dafny.Map
        d_264_v5_ = _dafny.Map({((d_262_v3_)[d_259_v0_] if (d_259_v0_) in (d_262_v3_) else d_261_v2_): d_263_v4_})
        d_265_v6_: _dafny.Seq
        d_265_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhhbc"))
        d_266_v7_: _dafny.Seq
        d_266_v7_ = _dafny.SeqWithoutIsStrInference([d_265_v6_])
        d_267_v8_: _dafny.Array
        nw20_ = _dafny.Array(_dafny.Array(None, 0), 21)
        d_267_v8_ = nw20_
        d_268_v9_: _dafny.Array
        def lambda4_(d_269_v4_):
            def lambda5_(d_270_i0_):
                return d_269_v4_

            return lambda5_

        init2_ = lambda4_(d_263_v4_)
        nw21_ = _dafny.Array(None, 1)
        for i0_2_ in range(nw21_.length(0)):
            nw21_[i0_2_] = init2_(i0_2_)
        d_268_v9_ = nw21_
        d_271_globalState_: GlobalState
        nw22_ = GlobalState()
        nw22_.ctor__(324, False, True, False, 862, (d_264_v5_).set(d_261_v2_, d_263_v4_), 38, d_265_v6_, (d_266_v7_)[default__.safeIndex(d_260_v1_, len(d_266_v7_))], True, 167, -961, True, 457, True, d_267_v8_, 745, _dafny.CodePoint('d'), d_268_v9_, True, True, d_268_v9_, d_259_v0_, True, _dafny.CodePoint('y'), d_261_v2_, False, 426)
        d_271_globalState_ = nw22_
        d_260_v1_ = 569
        if d_263_v4_:
            d_272_v10_: _dafny.MultiSet
            d_272_v10_ = _dafny.MultiSet([d_260_v1_])
            d_273_v11_: _dafny.Map
            d_273_v11_ = _dafny.Map({d_263_v4_: d_260_v1_})
            d_274_v12_: _dafny.Seq
            d_274_v12_ = _dafny.SeqWithoutIsStrInference([d_263_v4_])
            d_275_v13_: _dafny.Map
            d_275_v13_ = _dafny.Map({(0) - (d_260_v1_): len(d_274_v12_)})
            d_260_v1_ = default__.safeModuloInt(((d_272_v10_)[((d_273_v11_)[d_263_v4_] if (d_263_v4_) in (d_273_v11_) else d_260_v1_)] if (((d_273_v11_)[d_263_v4_] if (d_263_v4_) in (d_273_v11_) else d_260_v1_)) in (d_272_v10_) else d_260_v1_), ((d_275_v13_)[d_260_v1_] if (d_260_v1_) in (d_275_v13_) else d_260_v1_))
            d_276_v14_: int
            out0_: int
            out0_ = default__.m0(d_261_v2_, d_263_v4_, d_260_v1_, (d_263_v4_) not in ((_dafny.Map({d_263_v4_: d_259_v0_})).set(True, d_259_v0_)), d_271_globalState_)
            d_276_v14_ = out0_
            (d_271_globalState_).f1 = (d_261_v2_) == (_dafny.SeqWithoutIsStrInference([(0) - (d_276_v14_) for d_277_i1_ in range(default__.abs(940))]))
            d_278_v15_: _dafny.Set
            d_278_v15_ = _dafny.Set({d_263_v4_, d_263_v4_})
            d_279_v16_: _dafny.Map
            d_279_v16_ = _dafny.Map({d_268_v9_: (d_263_v4_) in (d_278_v15_)})
            d_279_v16_ = d_279_v16_
            d_280_v17_: _dafny.MultiSet
            d_280_v17_ = _dafny.MultiSet([d_263_v4_, d_263_v4_])
            d_281_v18_: str
            d_281_v18_ = _dafny.CodePoint('b')
            d_282_v19_: _dafny.Map
            d_282_v19_ = _dafny.Map({d_263_v4_: d_281_v18_})
            d_283_v20_: _dafny.Array
            nw23_ = _dafny.Array(None, 11)
            nw23_[int(0)] = d_274_v12_
            nw23_[int(1)] = d_274_v12_
            nw23_[int(2)] = (d_274_v12_) + (default__.fm0(d_276_v14_, d_263_v4_, d_263_v4_, (d_280_v17_).cardinality, d_271_globalState_))
            nw23_[int(3)] = _dafny.SeqWithoutIsStrInference([d_263_v4_])
            nw23_[int(4)] = (d_274_v12_ if d_263_v4_ else d_274_v12_)
            nw23_[int(5)] = (d_274_v12_).set(default__.safeIndex((0) - (len(d_282_v19_)), len(d_274_v12_)), d_263_v4_)
            nw23_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_263_v4_, d_263_v4_])) + (d_274_v12_)
            nw23_[int(7)] = d_274_v12_
            nw23_[int(8)] = default__.fm0(d_260_v1_, d_263_v4_, not(d_263_v4_), (d_280_v17_).cardinality, d_271_globalState_)
            nw23_[int(9)] = d_274_v12_
            nw23_[int(10)] = (d_274_v12_) + (d_274_v12_)
            d_283_v20_ = nw23_
            index7_ = default__.safeIndex(78, (d_283_v20_).length(0))
            (d_283_v20_)[index7_] = d_274_v12_
            index8_ = default__.safeIndex(78, (d_283_v20_).length(0))
            (d_283_v20_)[index8_] = d_274_v12_
        elif True:
            index9_ = default__.safeIndex(648, (d_268_v9_).length(0))
            (d_268_v9_)[index9_] = d_263_v4_
            d_284_v21_: _dafny.MultiSet
            d_284_v21_ = _dafny.MultiSet([not(d_263_v4_)])
            d_285_v22_: _dafny.Map
            d_285_v22_ = _dafny.Map({d_263_v4_: (d_263_v4_) not in (d_284_v21_)})
            index10_ = default__.safeIndex(648, (d_268_v9_).length(0))
            rhs11_ = ((d_285_v22_)[not (d_263_v4_) or (d_263_v4_)] if (not (d_263_v4_) or (d_263_v4_)) in (d_285_v22_) else d_263_v4_)
            rhs12_ = not(default__.fm1((False if d_263_v4_ else d_263_v4_), d_260_v1_, d_271_globalState_))
            lhs9_ = d_268_v9_
            lhs10_ = default__.safeIndex(648, (d_268_v9_).length(0))
            lhs11_ = d_271_globalState_
            lhs9_[lhs10_] = rhs11_
            lhs11_.f26 = rhs12_
            d_286_v23_: _dafny.MultiSet
            d_286_v23_ = _dafny.MultiSet([d_260_v1_])
            d_287_v24_: _dafny.Array
            def lambda6_(d_288_i2_):
                return False

            init3_ = lambda6_
            nw24_ = _dafny.Array(None, 28)
            for i0_3_ in range(nw24_.length(0)):
                nw24_[i0_3_] = init3_(i0_3_)
            d_287_v24_ = nw24_
            d_289_v25_: _dafny.Set
            d_289_v25_ = _dafny.Set({d_287_v24_})
            d_290_v26_: _dafny.Map
            d_290_v26_ = _dafny.Map({not((d_268_v9_)[default__.safeIndex(648, (d_268_v9_).length(0))]): d_287_v24_})
            d_291_v27_: _dafny.Map
            d_291_v27_ = _dafny.Map({d_286_v23_: (d_289_v25_).isdisjoint(_dafny.Set({d_287_v24_, ((d_290_v26_)[d_263_v4_] if (d_263_v4_) in (d_290_v26_) else d_287_v24_)}))})
            (d_271_globalState_).f1 = ((d_291_v27_)[d_286_v23_] if (d_286_v23_) in (d_291_v27_) else d_263_v4_)
            d_292_v28_: _dafny.Seq
            d_292_v28_ = _dafny.SeqWithoutIsStrInference([(d_268_v9_)[default__.safeIndex(648, (d_268_v9_).length(0))]])
            d_293_v29_: int
            out1_: int
            out1_ = default__.m0(d_261_v2_, (d_292_v28_)[default__.safeIndex(780, len(d_292_v28_))], d_260_v1_, not(d_263_v4_), d_271_globalState_)
            d_293_v29_ = out1_
            d_292_v28_ = d_292_v28_
            d_294_v30_: _dafny.Set
            d_294_v30_ = _dafny.Set({len(d_292_v28_)})
            d_295_v31_: _dafny.Set
            d_295_v31_ = _dafny.Set({d_292_v28_, d_292_v28_, d_292_v28_})
            d_296_v33_: _dafny.MultiSet
            def iife17_():
                coll17_ = _dafny.Set()
                compr_17_: _dafny.Seq
                for compr_17_ in (d_295_v31_).Elements:
                    d_297_v32_: _dafny.Seq = compr_17_
                    if (d_297_v32_) in (d_295_v31_):
                        coll17_ = coll17_.union(_dafny.Set([d_297_v32_]))
                return _dafny.Set(coll17_)
            d_296_v33_ = _dafny.MultiSet([iife17_()
            , d_295_v31_])
            (d_271_globalState_).f26 = not(((d_294_v30_).isdisjoint(d_294_v30_)) == ((d_293_v29_) > (((d_296_v33_)[d_295_v31_] if (d_295_v31_) in (d_296_v33_) else d_293_v29_))))
        hi2_ = d_260_v1_
        for d_298_i3_ in range(d_260_v1_, hi2_):
            (d_271_globalState_).f1 = d_263_v4_
            d_299_v34_: _dafny.MultiSet
            d_299_v34_ = _dafny.MultiSet([d_298_i3_])
            d_300_v35_: C7
            nw25_ = C7()
            nw25_.ctor__(d_298_i3_, d_299_v34_, (d_260_v1_) * (568), not (d_263_v4_) or (d_263_v4_), d_263_v4_, d_263_v4_)
            d_300_v35_ = nw25_
            (d_271_globalState_).f12 = (d_298_i3_) >= ((d_300_v35_).f41)
            (d_271_globalState_).f13 = d_298_i3_
        rhs13_ = d_260_v1_
        rhs14_ = 417
        lhs12_ = d_271_globalState_
        lhs12_.f13 = rhs13_
        d_260_v1_ = rhs14_
        d_301_v36_: C7
        nw26_ = C7()
        nw26_.ctor__(d_260_v1_, _dafny.MultiSet(d_261_v2_), 780, d_263_v4_, False, d_263_v4_)
        d_301_v36_ = nw26_
        d_302_v37_: _dafny.Map
        d_302_v37_ = _dafny.Map({d_268_v9_: d_301_v36_})
        d_303_v38_: _dafny.Map
        d_303_v38_ = _dafny.Map({(d_301_v36_).f41: (d_302_v37_).set(d_268_v9_, d_301_v36_)})
        d_302_v37_ = (d_302_v37_) | (((d_303_v38_)[default__.fm46(811, d_260_v1_, d_263_v4_, d_271_globalState_)] if (default__.fm46(811, d_260_v1_, d_263_v4_, d_271_globalState_)) in (d_303_v38_) else d_302_v37_))
        d_304_v39_: _dafny.Map
        d_304_v39_ = _dafny.Map({(d_261_v2_)[default__.safeIndex(d_260_v1_, len(d_261_v2_))]: ((d_301_v36_).f41) - (d_260_v1_)})
        d_305_v40_: str
        d_305_v40_ = _dafny.CodePoint('m')
        d_306_v41_: _dafny.Set
        d_306_v41_ = _dafny.Set({d_305_v40_})
        d_304_v39_ = (d_304_v39_).set(default__.safeModuloInt((0) - (d_260_v1_), d_260_v1_), len((d_306_v41_) | (d_306_v41_)))
        d_307_v42_: _dafny.Array
        nw27_ = _dafny.Array(_dafny.Seq({}), 6)
        d_307_v42_ = nw27_
        d_308_v43_: _dafny.Array
        d_308_v43_ = d_307_v42_
        d_307_v42_ = (d_308_v43_)
        hi3_ = (d_301_v36_).f41
        for d_309_i4_ in range(d_260_v1_, hi3_):
            if d_263_v4_:
                d_310_v44_: D22
                d_310_v44_ = D22_DC55(263, d_260_v1_, d_263_v4_)
                d_311_v45_: D3
                d_311_v45_ = D3_DC9((d_310_v44_).cf94, d_263_v4_, d_268_v9_, d_263_v4_)
                d_312_v46_: _dafny.Array
                d_313_v47_: int
                out2_: _dafny.Array
                out3_: int
                out2_, out3_ = (d_301_v36_).m8((d_311_v45_).cf20, d_271_globalState_)
                d_312_v46_ = out2_
                d_313_v47_ = out3_
                d_260_v1_ = d_309_i4_
                (d_271_globalState_).f12 = d_263_v4_
                d_314_v48_: _dafny.Array
                d_315_v49_: int
                out4_: _dafny.Array
                out5_: int
                out4_, out5_ = (d_301_v36_).m8(d_268_v9_, d_271_globalState_)
                d_314_v48_ = out4_
                d_315_v49_ = out5_
                (d_271_globalState_).f26 = d_263_v4_
            elif True:
                (d_271_globalState_).f1 = (934) <= (d_309_i4_)
                d_316_v50_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.Set({}), 13)
                d_316_v50_ = nw28_
                d_317_v51_: _dafny.Set
                d_317_v51_ = _dafny.Set({d_263_v4_})
                d_318_v52_: _dafny.Seq
                d_318_v52_ = _dafny.SeqWithoutIsStrInference([d_263_v4_])
                d_319_v53_: _dafny.Seq
                d_319_v53_ = _dafny.SeqWithoutIsStrInference([d_318_v52_])
                d_320_v54_: D14
                d_320_v54_ = D14_DC37((d_301_v36_).f41, d_263_v4_, (d_301_v36_).f41)
                d_321_v55_: _dafny.MultiSet
                d_321_v55_ = _dafny.MultiSet([True])
                d_322_v56_: _dafny.Map
                d_322_v56_ = _dafny.Map({d_263_v4_: not(False)})
                nw29_ = _dafny.Array(None, 28)
                nw29_[int(0)] = d_317_v51_
                nw29_[int(1)] = d_317_v51_
                nw29_[int(2)] = d_317_v51_
                nw29_[int(3)] = (d_317_v51_) - (_dafny.Set({d_263_v4_, d_263_v4_}))
                nw29_[int(4)] = (d_317_v51_) | (d_317_v51_)
                nw29_[int(5)] = (d_317_v51_) - (default__.fm11(d_319_v53_, not(d_263_v4_), d_263_v4_, (d_301_v36_).f41, d_271_globalState_))
                nw29_[int(6)] = d_317_v51_
                nw29_[int(7)] = _dafny.Set({False})
                nw29_[int(8)] = d_317_v51_
                nw29_[int(9)] = d_317_v51_
                nw29_[int(10)] = (d_317_v51_) - (d_317_v51_)
                nw29_[int(11)] = _dafny.Set({d_263_v4_, d_263_v4_, d_263_v4_})
                nw29_[int(12)] = (d_317_v51_) - (d_317_v51_)
                nw29_[int(13)] = _dafny.Set({d_263_v4_, d_263_v4_, d_263_v4_})
                nw29_[int(14)] = d_317_v51_
                nw29_[int(15)] = d_317_v51_
                nw29_[int(16)] = d_317_v51_
                nw29_[int(17)] = (d_317_v51_)
                nw29_[int(18)] = default__.fm45((d_320_v54_).cf70, d_321_v55_, ((d_322_v56_)[(d_301_v36_).fm19(d_271_globalState_)] if ((d_301_v36_).fm19(d_271_globalState_)) in (d_322_v56_) else True), d_309_i4_, d_271_globalState_)
                nw29_[int(19)] = d_317_v51_
                nw29_[int(20)] = d_317_v51_
                nw29_[int(21)] = (d_317_v51_).intersection(d_317_v51_)
                nw29_[int(22)] = d_317_v51_
                nw29_[int(23)] = _dafny.Set({(d_318_v52_)[default__.safeIndex((d_301_v36_).f41, len(d_318_v52_))], d_263_v4_})
                nw29_[int(24)] = (d_317_v51_).intersection(d_317_v51_)
                nw29_[int(25)] = d_317_v51_
                nw29_[int(26)] = d_317_v51_
                nw29_[int(27)] = (default__.fm45(len(d_265_v6_), d_321_v55_, (d_301_v36_).fm19(d_271_globalState_), 492, d_271_globalState_)) | (d_317_v51_)
                d_316_v50_ = nw29_
                d_323_v57_: D1
                d_323_v57_ = D1_DC2(d_268_v9_)
                pat_let_tv0_ = d_268_v9_
                def iife18_(_pat_let0_0):
                    def iife19_(d_324_dt__update__tmp_h0_):
                        def iife20_(_pat_let1_0):
                            def iife21_(d_325_dt__update_hcf2_h0_):
                                return D1_DC2(d_325_dt__update_hcf2_h0_)
                            return iife21_(_pat_let1_0)
                        return iife20_(pat_let_tv0_)
                    return iife19_(_pat_let0_0)
                d_323_v57_ = iife18_(D1_DC2(d_268_v9_))
                d_326_v58_: _dafny.Array
                d_327_v59_: int
                out6_: _dafny.Array
                out7_: int
                out6_, out7_ = (d_301_v36_).m8(d_268_v9_, d_271_globalState_)
                d_326_v58_ = out6_
                d_327_v59_ = out7_
                d_328_v60_: _dafny.Array
                def lambda7_(d_329_v40_):
                    def lambda8_(d_330_i5_):
                        return d_329_v40_

                    return lambda8_

                init4_ = lambda7_(d_305_v40_)
                nw30_ = _dafny.Array(None, 11)
                for i0_4_ in range(nw30_.length(0)):
                    nw30_[i0_4_] = init4_(i0_4_)
                d_328_v60_ = nw30_
                d_331_v61_: C0
                nw31_ = C0()
                nw31_.ctor__(d_328_v60_)
                d_331_v61_ = nw31_
                d_331_v61_ = d_331_v61_
            d_332_v62_: C2
            nw32_ = C2()
            nw32_.ctor__((len(_dafny.SeqWithoutIsStrInference([d_265_v6_]))) * (len(d_265_v6_)), d_260_v1_, d_263_v4_, d_263_v4_, d_263_v4_)
            d_332_v62_ = nw32_
            d_333_v63_: _dafny.Array
            def lambda9_(d_334_v4_):
                def lambda10_(d_335_i6_):
                    return _dafny.MultiSet([d_334_v4_])

                return lambda10_

            init5_ = lambda9_(d_263_v4_)
            nw33_ = _dafny.Array(None, 23)
            for i0_5_ in range(nw33_.length(0)):
                nw33_[i0_5_] = init5_(i0_5_)
            d_333_v63_ = nw33_
            d_336_v64_: _dafny.MultiSet
            d_336_v64_ = _dafny.MultiSet([d_263_v4_])
            index11_ = default__.safeIndex(405, (d_333_v63_).length(0))
            (d_333_v63_)[index11_] = (d_336_v64_) | (d_336_v64_)
            d_337_v65_: _dafny.Seq
            d_337_v65_ = _dafny.SeqWithoutIsStrInference([d_263_v4_, d_263_v4_])
            index12_ = default__.safeIndex(405, (d_333_v63_).length(0))
            (d_333_v63_)[index12_] = ((_dafny.MultiSet([d_263_v4_, d_263_v4_]) if (d_337_v65_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_263_v4_: d_263_v4_}) for d_338_i7_ in range(default__.abs(292))])), len(d_337_v65_))] else _dafny.MultiSet(((d_337_v65_).set(default__.safeIndex(len(d_265_v6_), len(d_337_v65_)), d_263_v4_)).set(default__.safeIndex((d_301_v36_).f41, len((d_337_v65_).set(default__.safeIndex(len(d_265_v6_), len(d_337_v65_)), d_263_v4_))), d_263_v4_)))).intersection(d_336_v64_)
            (d_271_globalState_).f13 = (0) - (len(d_265_v6_))
        d_339_v66_: _dafny.Seq
        d_339_v66_ = _dafny.SeqWithoutIsStrInference([d_263_v4_])
        index13_ = default__.safeIndex(179, (d_259_v0_).length(0))
        (d_259_v0_)[index13_] = (len(d_339_v66_)) * ((d_301_v36_).f41)
        d_340_v67_: D8
        d_340_v67_ = D8_DC24(False, len(d_265_v6_), d_339_v66_, True, d_261_v2_)
        d_341_v68_: _dafny.MultiSet
        d_341_v68_ = _dafny.MultiSet([d_305_v40_])
        d_342_v69_: D3
        d_342_v69_ = D3_DC8(d_341_v68_, d_263_v4_, (d_301_v36_).f41)
        d_343_v70_: _dafny.Set
        d_343_v70_ = _dafny.Set({False, d_263_v4_, d_263_v4_, (d_340_v67_).cf49, (d_342_v69_).cf16})
        index14_ = default__.safeIndex(179, (d_259_v0_).length(0))
        (d_259_v0_)[index14_] = default__.safeDivisionInt(default__.safeDivisionInt((d_301_v36_).f41, len(d_343_v70_)), d_260_v1_)
        (d_271_globalState_).f1 = (not(False) if d_263_v4_ else d_263_v4_)
        d_344_i8_: int
        d_344_i8_ = 0
        with _dafny.label("0"):
            while d_263_v4_:
                with _dafny.c_label("0"):
                    if (d_344_i8_) >= (100):
                        raise _dafny.Break("0")
                    d_344_i8_ = (d_344_i8_) + (1)
                    d_345_v71_: str
                    d_346_v72_: _dafny.Seq
                    d_347_v73_: int
                    d_348_v74_: _dafny.Seq
                    out8_: str
                    out9_: _dafny.Seq
                    out10_: int
                    out11_: _dafny.Seq
                    out8_, out9_, out10_, out11_ = (d_301_v36_).m7(False, d_268_v9_, (d_343_v70_).issubset(_dafny.Set({(d_301_v36_).fm19(d_271_globalState_)})), d_271_globalState_)
                    d_345_v71_ = out8_
                    d_346_v72_ = out9_
                    d_347_v73_ = out10_
                    d_348_v74_ = out11_
                    (d_271_globalState_).f1 = default__.fm1((d_339_v66_)[default__.safeIndex(-315, len(d_339_v66_))], (d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], d_271_globalState_)
                    d_347_v73_ = (d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))]
                    d_349_v75_: _dafny.Array
                    nw34_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                    d_349_v75_ = nw34_
                    d_350_v76_: D20
                    d_350_v76_ = D20_DC50(D20_DC47(d_349_v75_))
                    d_351_v77_: _dafny.Map
                    d_351_v77_ = _dafny.Map({d_263_v4_: d_350_v76_})
                    d_352_v78_: D20
                    d_352_v78_ = D20_DC50(((d_351_v77_)[d_263_v4_] if (d_263_v4_) in (d_351_v77_) else d_350_v76_))
                    source8_ = d_352_v78_
                    if source8_.is_DC48:
                        d_353___mcc_h0_ = source8_.cf82
                        d_354___mcc_h1_ = source8_.cf83
                        d_355___mcc_h2_ = source8_.cf84
                        d_356_cf84_ = d_355___mcc_h2_
                        d_357_cf83_ = d_354___mcc_h1_
                        d_358_cf82_ = d_353___mcc_h0_
                        d_359_v79_: _dafny.Seq
                        d_359_v79_ = _dafny.SeqWithoutIsStrInference([(D1_DC2(d_268_v9_)).cf2, d_268_v9_, d_268_v9_, d_268_v9_])
                        (d_271_globalState_).f21 = (d_359_v79_)[default__.safeIndex(d_358_cf82_, len(d_359_v79_))]
                        d_266_v7_ = (((d_266_v7_).set(default__.safeIndex(507, len(d_266_v7_)), (d_301_v36_).fm3(d_263_v4_, False, d_271_globalState_))).set(default__.safeIndex(default__.safeDivisionInt(default__.fm39(d_263_v4_, (d_301_v36_).f42, d_263_v4_, d_271_globalState_), -975), len((d_266_v7_).set(default__.safeIndex(507, len(d_266_v7_)), (d_301_v36_).fm3(d_263_v4_, False, d_271_globalState_)))), d_348_v74_)).set(default__.safeIndex((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], len(((d_266_v7_).set(default__.safeIndex(507, len(d_266_v7_)), (d_301_v36_).fm3(d_263_v4_, False, d_271_globalState_))).set(default__.safeIndex(default__.safeDivisionInt(default__.fm39(d_263_v4_, (d_301_v36_).f42, d_263_v4_, d_271_globalState_), -975), len((d_266_v7_).set(default__.safeIndex(507, len(d_266_v7_)), (d_301_v36_).fm3(d_263_v4_, False, d_271_globalState_)))), d_348_v74_))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydtxhh")))
                        d_356_cf84_ = d_356_cf84_
                        d_304_v39_ = (d_304_v39_).set(((d_301_v36_).f41) + ((d_301_v36_).f41), (d_301_v36_).f41)
                    elif source8_.is_DC49:
                        d_360___mcc_h3_ = source8_.cf85
                        d_361___mcc_h4_ = source8_.cf86
                        d_362_cf86_ = d_361___mcc_h4_
                        d_363_cf85_ = d_360___mcc_h3_
                        index15_ = default__.safeIndex(259, (d_268_v9_).length(0))
                        (d_268_v9_)[index15_] = (d_301_v36_).fm19(d_271_globalState_)
                        d_364_v81_: _dafny.Map
                        d_364_v81_ = _dafny.Map({d_362_cf86_: d_345_v71_})
                        d_365_v82_: C2
                        nw35_ = C2()
                        def iife22_():
                            coll18_ = _dafny.Map()
                            compr_18_: str
                            for compr_18_ in (d_364_v81_).keys.Elements:
                                d_366_v80_: str = compr_18_
                                if (d_366_v80_) in (d_364_v81_):
                                    coll18_[d_366_v80_] = d_263_v4_
                            return _dafny.Map(coll18_)
                        def iife23_():
                            coll19_ = _dafny.Map()
                            compr_19_: str
                            for compr_19_ in (d_364_v81_).keys.Elements:
                                d_367_v80_: str = compr_19_
                                if (d_367_v80_) in (d_364_v81_):
                                    coll19_[d_367_v80_] = d_263_v4_
                            return _dafny.Map(coll19_)
                        nw35_.ctor__((((d_301_v36_).f42)[default__.fm46((0) - (len(iife22_()
                        )), (d_301_v36_).f41, d_363_cf85_, d_271_globalState_)] if (default__.fm46((0) - (len(iife23_()
                        )), (d_301_v36_).f41, d_363_cf85_, d_271_globalState_)) in ((d_301_v36_).f42) else (((d_301_v36_).f42)[(d_301_v36_).f41] if ((d_301_v36_).f41) in ((d_301_v36_).f42) else (d_301_v36_).f41)), len(d_348_v74_), d_263_v4_, not((d_339_v66_)[default__.safeIndex(393, len(d_339_v66_))]), d_263_v4_)
                        d_365_v82_ = nw35_
                        d_368_v83_: _dafny.Map
                        d_368_v83_ = _dafny.Map({len(d_261_v2_): d_365_v82_})
                        d_369_v84_: _dafny.Array
                        nw36_ = _dafny.Array(None, 23)
                        nw36_[int(0)] = d_365_v82_
                        nw36_[int(1)] = d_365_v82_
                        nw36_[int(2)] = ((d_368_v83_)[(d_301_v36_).f41] if ((d_301_v36_).f41) in (d_368_v83_) else d_365_v82_)
                        nw36_[int(3)] = d_365_v82_
                        nw36_[int(4)] = d_365_v82_
                        nw36_[int(5)] = d_365_v82_
                        nw36_[int(6)] = d_365_v82_
                        nw36_[int(7)] = d_365_v82_
                        nw36_[int(8)] = d_365_v82_
                        nw36_[int(9)] = d_365_v82_
                        nw36_[int(10)] = d_365_v82_
                        nw36_[int(11)] = d_365_v82_
                        nw36_[int(12)] = ((d_368_v83_)[22] if (22) in (d_368_v83_) else d_365_v82_)
                        nw36_[int(13)] = d_365_v82_
                        nw36_[int(14)] = d_365_v82_
                        nw36_[int(15)] = d_365_v82_
                        nw36_[int(16)] = d_365_v82_
                        nw36_[int(17)] = d_365_v82_
                        nw36_[int(18)] = d_365_v82_
                        nw36_[int(19)] = d_365_v82_
                        nw36_[int(20)] = d_365_v82_
                        nw36_[int(21)] = d_365_v82_
                        nw36_[int(22)] = d_365_v82_
                        d_369_v84_ = nw36_
                        index16_ = default__.safeIndex(383, (d_369_v84_).length(0))
                        (d_369_v84_)[index16_] = d_365_v82_
                        d_370_v85_: D20
                        d_370_v85_ = D20_DC49(d_263_v4_, d_305_v40_)
                        d_371_v86_: D5
                        d_371_v86_ = D5_DC13(d_363_cf85_, _dafny.Map({not(d_363_cf85_): (d_365_v82_).f46}), d_305_v40_)
                        d_372_v87_: _dafny.Array
                        nw37_ = _dafny.Array(None, 29)
                        nw37_[int(0)] = d_345_v71_
                        nw37_[int(1)] = d_305_v40_
                        nw37_[int(2)] = default__.fm41(d_363_cf85_, d_362_cf86_, d_263_v4_, (d_301_v36_).f41, d_271_globalState_)
                        nw37_[int(3)] = d_305_v40_
                        nw37_[int(4)] = d_345_v71_
                        nw37_[int(5)] = d_345_v71_
                        nw37_[int(6)] = d_345_v71_
                        nw37_[int(7)] = _dafny.CodePoint('t')
                        nw37_[int(8)] = d_305_v40_
                        nw37_[int(9)] = d_345_v71_
                        nw37_[int(10)] = d_362_cf86_
                        nw37_[int(11)] = (d_370_v85_).cf86
                        nw37_[int(12)] = d_345_v71_
                        nw37_[int(13)] = d_305_v40_
                        nw37_[int(14)] = _dafny.CodePoint('x')
                        nw37_[int(15)] = _dafny.CodePoint('t')
                        nw37_[int(16)] = (d_371_v86_).cf29
                        nw37_[int(17)] = d_345_v71_
                        nw37_[int(18)] = d_345_v71_
                        nw37_[int(19)] = d_345_v71_
                        nw37_[int(20)] = _dafny.CodePoint('a')
                        nw37_[int(21)] = d_345_v71_
                        nw37_[int(22)] = d_305_v40_
                        nw37_[int(23)] = d_362_cf86_
                        nw37_[int(24)] = d_305_v40_
                        nw37_[int(25)] = d_305_v40_
                        nw37_[int(26)] = d_362_cf86_
                        nw37_[int(27)] = d_362_cf86_
                        nw37_[int(28)] = default__.fm48(d_363_cf85_, d_363_cf85_, d_260_v1_, d_271_globalState_)
                        d_372_v87_ = nw37_
                        index17_ = default__.safeIndex(880, (d_372_v87_).length(0))
                        (d_372_v87_)[index17_] = _dafny.CodePoint('m')
                        index18_ = default__.safeIndex(259, (d_268_v9_).length(0))
                        index19_ = default__.safeIndex(383, (d_369_v84_).length(0))
                        index20_ = default__.safeIndex(880, (d_372_v87_).length(0))
                        rhs15_ = (d_365_v82_).fm30(d_363_cf85_, d_271_globalState_)
                        rhs16_ = d_365_v82_
                        rhs17_ = default__.safeModuloInt((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], 908)
                        rhs18_ = not(d_263_v4_)
                        rhs19_ = d_362_cf86_
                        lhs13_ = d_268_v9_
                        lhs14_ = default__.safeIndex(259, (d_268_v9_).length(0))
                        lhs15_ = d_369_v84_
                        lhs16_ = default__.safeIndex(383, (d_369_v84_).length(0))
                        lhs17_ = d_271_globalState_
                        lhs18_ = d_372_v87_
                        lhs19_ = default__.safeIndex(880, (d_372_v87_).length(0))
                        lhs13_[lhs14_] = rhs15_
                        lhs15_[lhs16_] = rhs16_
                        d_347_v73_ = rhs17_
                        lhs17_.f1 = rhs18_
                        lhs18_[lhs19_] = rhs19_
                        (d_271_globalState_).f26 = (d_268_v9_)[default__.safeIndex(259, (d_268_v9_).length(0))]
                        d_301_v36_ = d_301_v36_
                        d_362_cf86_ = (d_372_v87_)[default__.safeIndex(880, (d_372_v87_).length(0))]
                    elif source8_.is_DC47:
                        d_373___mcc_h5_ = source8_.cf81
                        d_374_cf81_ = d_373___mcc_h5_
                        d_375_v88_: _dafny.Array
                        d_376_v89_: int
                        out12_: _dafny.Array
                        out13_: int
                        out12_, out13_ = (d_301_v36_).m8(d_268_v9_, d_271_globalState_)
                        d_375_v88_ = out12_
                        d_376_v89_ = out13_
                        d_377_v90_: _dafny.Map
                        d_377_v90_ = _dafny.Map({d_263_v4_: d_263_v4_})
                        d_378_v91_: _dafny.Map
                        d_378_v91_ = _dafny.Map({(d_301_v36_).f41: d_377_v90_})
                        d_379_v92_: _dafny.Set
                        d_379_v92_ = _dafny.Set({d_259_v0_})
                        d_380_v93_: C9
                        nw38_ = C9()
                        nw38_.ctor__(d_376_v89_, d_378_v91_, d_305_v40_, d_376_v89_, not(d_263_v4_), (d_339_v66_).set(default__.safeIndex(74, len(d_339_v66_)), (d_301_v36_).fm19(d_271_globalState_)), (_dafny.SeqWithoutIsStrInference([d_339_v66_, d_339_v66_])).set(default__.safeIndex((d_301_v36_).f41, len(_dafny.SeqWithoutIsStrInference([d_339_v66_, d_339_v66_]))), d_339_v66_), (d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], d_263_v4_, (d_379_v92_).ispropersubset(d_379_v92_), d_263_v4_)
                        d_380_v93_ = nw38_
                        d_380_v93_ = d_380_v93_
                        (d_271_globalState_).f1 = not(not(((_dafny.SeqWithoutIsStrInference([d_305_v40_ for d_381_i9_ in range(default__.abs(851))])).set(default__.safeIndex((d_301_v36_).f41, len(_dafny.SeqWithoutIsStrInference([d_305_v40_ for d_382_i9_ in range(default__.abs(851))]))), d_345_v71_)) < ((d_348_v74_) + (d_348_v74_))))
                        index21_ = default__.safeIndex(179, (d_259_v0_).length(0))
                        (d_259_v0_)[index21_] = ((len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D21_DC53(D21_DC53(D21_DC52(d_345_v71_))), D21_DC53(D21_DC51(_dafny.Map({(d_301_v36_).f41: _dafny.Set({(d_301_v36_).f41, d_376_v89_})})))])) for d_383_i10_ in range(default__.abs(265))]))) - (len(d_265_v6_))) * ((d_301_v36_).f41)
                    elif True:
                        d_384___mcc_h6_ = source8_.cf87
                        d_385_cf87_ = d_384___mcc_h6_
                        d_386_v94_: _dafny.Array
                        nw39_ = _dafny.Array(_dafny.Array(None, 0), 23)
                        d_386_v94_ = nw39_
                        index22_ = default__.safeIndex(486, (d_386_v94_).length(0))
                        (d_386_v94_)[index22_] = d_268_v9_
                        index23_ = default__.safeIndex(486, (d_386_v94_).length(0))
                        (d_386_v94_)[index23_] = d_268_v9_
                        d_387_v95_: _dafny.Set
                        d_387_v95_ = _dafny.Set({d_301_v36_, d_301_v36_, d_301_v36_})
                        d_388_v96_: _dafny.Set
                        d_388_v96_ = _dafny.Set({d_387_v95_})
                        d_389_v97_: _dafny.Seq
                        d_389_v97_ = _dafny.SeqWithoutIsStrInference([d_388_v96_])
                        d_390_v98_: D24
                        d_390_v98_ = D24_DC57(_dafny.Set({d_387_v95_, d_387_v95_}))
                        d_391_v99_: _dafny.Array
                        nw40_ = _dafny.Array(None, 15)
                        nw40_[int(0)] = d_388_v96_
                        nw40_[int(1)] = _dafny.Set({d_387_v95_, d_387_v95_})
                        nw40_[int(2)] = d_388_v96_
                        nw40_[int(3)] = (d_388_v96_) | (_dafny.Set({d_387_v95_}))
                        nw40_[int(4)] = d_388_v96_
                        nw40_[int(5)] = d_388_v96_
                        nw40_[int(6)] = (d_388_v96_) - ((d_389_v97_)[default__.safeIndex(len(d_348_v74_), len(d_389_v97_))])
                        nw40_[int(7)] = d_388_v96_
                        nw40_[int(8)] = d_388_v96_
                        nw40_[int(9)] = d_388_v96_
                        nw40_[int(10)] = _dafny.Set({d_387_v95_, d_387_v95_})
                        nw40_[int(11)] = d_388_v96_
                        nw40_[int(12)] = d_388_v96_
                        nw40_[int(13)] = (d_390_v98_).cf96
                        nw40_[int(14)] = d_388_v96_
                        d_391_v99_ = nw40_
                        index24_ = default__.safeIndex(53, (d_391_v99_).length(0))
                        (d_391_v99_)[index24_] = d_388_v96_
                        index25_ = default__.safeIndex(53, (d_391_v99_).length(0))
                        (d_391_v99_)[index25_] = d_388_v96_
                        d_392_v100_: D3
                        d_392_v100_ = D3_DC9(d_263_v4_, d_263_v4_, (d_386_v94_)[default__.safeIndex(486, (d_386_v94_).length(0))], d_263_v4_)
                        pat_let_tv1_ = d_263_v4_
                        pat_let_tv2_ = d_263_v4_
                        d_393_v101_: _dafny.Map
                        def iife24_(_pat_let2_0):
                            def iife25_(d_394_dt__update__tmp_h1_):
                                def iife26_(_pat_let3_0):
                                    def iife27_(d_395_dt__update_hcf21_h0_):
                                        def iife28_(_pat_let4_0):
                                            def iife29_(d_396_dt__update_hcf18_h0_):
                                                return D3_DC9(d_396_dt__update_hcf18_h0_, (d_394_dt__update__tmp_h1_).cf19, (d_394_dt__update__tmp_h1_).cf20, d_395_dt__update_hcf21_h0_)
                                            return iife29_(_pat_let4_0)
                                        return iife28_(pat_let_tv2_)
                                    return iife27_(_pat_let3_0)
                                return iife26_(pat_let_tv1_)
                            return iife25_(_pat_let2_0)
                        d_393_v101_ = _dafny.Map({iife24_(d_392_v100_): (d_301_v36_).f41})
                        d_397_v102_: _dafny.Map
                        d_397_v102_ = _dafny.Map({(d_301_v36_).f41: d_305_v40_})
                        d_398_v103_: D2
                        d_398_v103_ = D2_DC6((d_301_v36_).f41, d_263_v4_, default__.fm46((d_301_v36_).f41, len(d_393_v101_), True, d_271_globalState_), d_263_v4_, len(d_397_v102_))
                        d_260_v1_ = ((d_398_v103_).cf13) - ((len(d_348_v74_)) * (d_260_v1_))
                        d_399_v104_: str
                        d_400_v105_: _dafny.Seq
                        d_401_v106_: int
                        d_402_v107_: _dafny.Seq
                        out14_: str
                        out15_: _dafny.Seq
                        out16_: int
                        out17_: _dafny.Seq
                        out14_, out15_, out16_, out17_ = (d_301_v36_).m7((d_263_v4_) in (d_339_v66_), (d_386_v94_)[default__.safeIndex(486, (d_386_v94_).length(0))], default__.fm1(True, (default__.fm33(d_263_v4_, d_271_globalState_)).cardinality, d_271_globalState_), d_271_globalState_)
                        d_399_v104_ = out14_
                        d_400_v105_ = out15_
                        d_401_v106_ = out16_
                        d_402_v107_ = out17_
                    pass
            pass
        d_403_v108_: _dafny.Seq
        d_403_v108_ = _dafny.SeqWithoutIsStrInference([d_339_v66_])
        d_404_v109_: C4
        nw41_ = C4()
        nw41_.ctor__((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], d_263_v4_, d_339_v66_, d_403_v108_, ((d_301_v36_).f41) + (len(d_343_v70_)), True, d_263_v4_, d_263_v4_)
        d_404_v109_ = nw41_
        d_266_v7_ = ((d_266_v7_).set(default__.safeIndex(d_260_v1_, len(d_266_v7_)), _dafny.SeqWithoutIsStrInference([d_305_v40_ for d_405_i11_ in range(default__.abs(904))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssisbjr"))]))
        d_406_v110_: _dafny.MultiSet
        d_406_v110_ = _dafny.MultiSet([not(True)])
        if (d_406_v110_).ispropersubset(d_406_v110_):
            d_407_v111_: _dafny.Array
            nw42_ = _dafny.Array(None, 2)
            d_407_v111_ = nw42_
            d_408_v112_: T0
            nw43_ = C1()
            nw43_.ctor__(False, d_263_v4_, d_263_v4_)
            d_408_v112_ = nw43_
            index26_ = default__.safeIndex(55, (d_407_v111_).length(0))
            (d_407_v111_)[index26_] = d_408_v112_
            d_409_v114_: _dafny.Array
            def lambda11_(d_410_v1_, d_411_v36_):
                def lambda12_(d_412_i12_):
                    def iife30_():
                        coll20_ = _dafny.Set()
                        compr_20_: int
                        for compr_20_ in _dafny.IntegerRange(718, 593):
                            d_413_v113_: int = compr_20_
                            if ((718) <= (d_413_v113_)) and ((d_413_v113_) < (593)):
                                coll20_ = coll20_.union(_dafny.Set([(d_413_v113_) + ((d_411_v36_).f41)]))
                        return _dafny.Set(coll20_)
                    return (_dafny.Set({646, (0) - (d_410_v1_)})) | (iife30_()
                    )

                return lambda12_

            init6_ = lambda11_(d_260_v1_, d_301_v36_)
            nw44_ = _dafny.Array(None, 20)
            for i0_6_ in range(nw44_.length(0)):
                nw44_[i0_6_] = init6_(i0_6_)
            d_409_v114_ = nw44_
            index27_ = default__.safeIndex(55, (d_407_v111_).length(0))
            rhs20_ = d_408_v112_
            rhs21_ = d_409_v114_
            rhs22_ = d_406_v110_
            lhs20_ = d_407_v111_
            lhs21_ = default__.safeIndex(55, (d_407_v111_).length(0))
            lhs20_[lhs21_] = rhs20_
            d_409_v114_ = rhs21_
            d_406_v110_ = rhs22_
            d_414_v115_: str
            d_415_v116_: _dafny.Seq
            d_416_v117_: int
            d_417_v118_: _dafny.Seq
            out18_: str
            out19_: _dafny.Seq
            out20_: int
            out21_: _dafny.Seq
            out18_, out19_, out20_, out21_ = (d_301_v36_).m7(False, d_268_v9_, d_408_v112_.f29, d_271_globalState_)
            d_414_v115_ = out18_
            d_415_v116_ = out19_
            d_416_v117_ = out20_
            d_417_v118_ = out21_
            d_418_v119_: _dafny.Map
            d_418_v119_ = _dafny.Map({d_263_v4_: (d_301_v36_).f41})
            d_419_v120_: T3
            nw45_ = C4()
            nw45_.ctor__(len(d_417_v118_), d_408_v112_.f28, d_339_v66_, d_403_v108_, ((d_418_v119_)[d_263_v4_] if (d_263_v4_) in (d_418_v119_) else (d_301_v36_).f41), d_263_v4_, (d_404_v109_).fm6(d_271_globalState_), d_408_v112_.f28)
            d_419_v120_ = nw45_
            d_420_v121_: _dafny.MultiSet
            d_420_v121_ = _dafny.MultiSet([d_419_v120_])
            d_421_v122_: D10
            d_421_v122_ = D10_DC30(d_305_v40_, d_420_v121_)
            source9_ = (d_421_v122_ if d_263_v4_ else d_421_v122_)
            if source9_.is_DC29:
                d_422___mcc_h7_ = source9_.cf56
                d_423___mcc_h8_ = source9_.cf57
                d_424_cf57_ = d_423___mcc_h8_
                d_425_cf56_ = d_422___mcc_h7_
                d_426_v123_: str
                d_427_v124_: _dafny.Seq
                d_428_v125_: int
                d_429_v126_: _dafny.Seq
                out22_: str
                out23_: _dafny.Seq
                out24_: int
                out25_: _dafny.Seq
                out22_, out23_, out24_, out25_ = (d_301_v36_).m7(d_263_v4_, d_268_v9_, d_263_v4_, d_271_globalState_)
                d_426_v123_ = out22_
                d_427_v124_ = out23_
                d_428_v125_ = out24_
                d_429_v126_ = out25_
                index28_ = default__.safeIndex(179, (d_259_v0_).length(0))
                (d_259_v0_)[index28_] = (0) - (len(_dafny.SeqWithoutIsStrInference([(d_419_v120_).f30])))
                d_430_v127_: C7
                nw46_ = C7()
                nw46_.ctor__((d_261_v2_)[default__.safeIndex(d_428_v125_, len(d_261_v2_))], ((d_301_v36_).f42) | ((d_301_v36_).f42), (-421) - (len(d_265_v6_)), d_419_v120_.f35, d_424_cf57_, d_419_v120_.f31)
                d_430_v127_ = nw46_
                rhs23_ = (d_301_v36_).f41
                rhs24_ = (d_419_v120_).fm3((_dafny.SeqWithoutIsStrInference([d_419_v120_.f28])) < (d_339_v66_), d_408_v112_.f28, d_271_globalState_)
                rhs25_ = not (d_408_v112_.f28) or (not (d_408_v112_.f29) or (d_263_v4_))
                rhs26_ = ((d_301_v36_).f42).ispropersubset((d_430_v127_).f42)
                lhs22_ = d_271_globalState_
                lhs23_ = d_419_v120_
                d_260_v1_ = rhs23_
                d_429_v126_ = rhs24_
                lhs22_.f12 = rhs25_
                lhs23_.f35 = rhs26_
            elif source9_.is_DC30:
                d_431___mcc_h9_ = source9_.cf58
                d_432___mcc_h10_ = source9_.cf59
                d_433_cf59_ = d_432___mcc_h10_
                d_434_cf58_ = d_431___mcc_h9_
                rhs27_ = d_417_v118_
                rhs28_ = d_268_v9_
                lhs24_ = d_271_globalState_
                d_417_v118_ = rhs27_
                lhs24_.f21 = rhs28_
                d_435_v128_: _dafny.Set
                d_435_v128_ = _dafny.Set({(d_301_v36_).f41})
                rhs29_ = d_419_v120_.f31
                rhs30_ = (d_435_v128_).issubset(d_435_v128_)
                lhs25_ = d_271_globalState_
                lhs26_ = d_271_globalState_
                lhs25_.f26 = rhs29_
                lhs26_.f26 = rhs30_
                d_436_v129_: _dafny.Seq
                d_436_v129_ = _dafny.SeqWithoutIsStrInference([d_261_v2_])
                d_418_v119_ = (d_418_v119_).set(d_419_v120_.f35, len(d_436_v129_))
                d_437_v130_: _dafny.Seq
                d_437_v130_ = _dafny.SeqWithoutIsStrInference([default__.fm66(d_271_globalState_)])
                d_438_v131_: T2
                nw47_ = C5()
                nw47_.ctor__(d_419_v120_.f31, len(_dafny.SeqWithoutIsStrInference([d_305_v40_ for d_439_i13_ in range(default__.abs(535))])), (0) - ((d_301_v36_).f41), d_419_v120_.f35, (d_419_v120_).f32, _dafny.SeqWithoutIsStrInference([(d_419_v120_).f32, (d_419_v120_).f32, (d_419_v120_).f32, _dafny.SeqWithoutIsStrInference([d_419_v120_.f29, d_419_v120_.f29, False, d_263_v4_, d_419_v120_.f31]), (d_419_v120_).f32]), len(d_437_v130_), True, d_408_v112_.f29, d_419_v120_.f35)
                d_438_v131_ = nw47_
                d_440_v132_: _dafny.MultiSet
                d_440_v132_ = _dafny.MultiSet([d_438_v131_])
                d_440_v132_ = (_dafny.MultiSet([d_438_v131_])) - (d_440_v132_)
            elif source9_.is_DC31:
                d_441___mcc_h11_ = source9_.cf60
                d_442___mcc_h12_ = source9_.cf61
                d_443___mcc_h13_ = source9_.cf62
                d_444_cf62_ = d_443___mcc_h13_
                d_445_cf61_ = d_442___mcc_h12_
                d_446_cf60_ = d_441___mcc_h11_
                d_447_v133_: C2
                nw48_ = C2()
                nw48_.ctor__(d_260_v1_, len((d_419_v120_).f32), d_419_v120_.f28, d_419_v120_.f29, d_419_v120_.f29)
                d_447_v133_ = nw48_
                d_448_v134_: D9
                d_448_v134_ = D9_DC26(d_447_v133_)
                d_449_v135_: _dafny.Array
                nw49_ = _dafny.Array(None, 22)
                d_449_v135_ = nw49_
                d_450_v136_: _dafny.Array
                nw50_ = _dafny.Array(_dafny.Seq({}), 21)
                d_450_v136_ = nw50_
                d_451_v137_: C3
                nw51_ = C3()
                nw51_.ctor__(-956, d_450_v136_, d_305_v40_, -341, d_419_v120_.f31, d_339_v66_, _dafny.SeqWithoutIsStrInference([d_339_v66_ for d_452_i14_ in range(default__.abs(860))]), 775, d_419_v120_.f28, d_446_cf60_, d_446_cf60_)
                d_451_v137_ = nw51_
                index29_ = default__.safeIndex(105, (d_449_v135_).length(0))
                (d_449_v135_)[index29_] = d_451_v137_
                d_453_v138_: _dafny.Array
                nw52_ = _dafny.Array(None, 1)
                nw52_[int(0)] = d_417_v118_
                d_453_v138_ = nw52_
                index30_ = default__.safeIndex(31, (d_453_v138_).length(0))
                (d_453_v138_)[index30_] = d_417_v118_
                index31_ = default__.safeIndex(105, (d_449_v135_).length(0))
                index32_ = default__.safeIndex(31, (d_453_v138_).length(0))
                rhs31_ = (d_261_v2_)[default__.safeIndex((d_419_v120_).f34, len(d_261_v2_))]
                rhs32_ = D9_DC26(d_447_v133_)
                rhs33_ = d_451_v137_
                rhs34_ = d_417_v118_
                rhs35_ = (d_419_v120_).f34
                lhs27_ = d_271_globalState_
                lhs28_ = d_449_v135_
                lhs29_ = default__.safeIndex(105, (d_449_v135_).length(0))
                lhs30_ = d_453_v138_
                lhs31_ = default__.safeIndex(31, (d_453_v138_).length(0))
                lhs27_.f13 = rhs31_
                d_448_v134_ = rhs32_
                lhs28_[lhs29_] = rhs33_
                lhs30_[lhs31_] = rhs34_
                d_260_v1_ = rhs35_
                d_454_v139_: bool
                d_455_v140_: int
                d_456_v141_: _dafny.Seq
                out26_: bool
                out27_: int
                out28_: _dafny.Seq
                out26_, out27_, out28_ = (d_419_v120_).m1(d_271_globalState_)
                d_454_v139_ = out26_
                d_455_v140_ = out27_
                d_456_v141_ = out28_
                d_455_v140_ = d_260_v1_
                d_457_v142_: D8
                d_457_v142_ = D8_DC24(d_408_v112_.f28, d_455_v140_, _dafny.SeqWithoutIsStrInference([d_446_cf60_]), d_263_v4_, d_261_v2_)
                d_458_v143_: D8
                d_458_v143_ = D8_DC25(d_457_v142_)
                d_459_v144_: D8
                d_459_v144_ = D8_DC25(d_458_v143_)
                d_460_v145_: _dafny.Map
                d_460_v145_ = _dafny.Map({True: d_414_v115_})
                rhs36_ = d_459_v144_
                rhs37_ = (d_419_v120_.f28) in (d_460_v145_)
                lhs32_ = d_419_v120_
                d_459_v144_ = rhs36_
                lhs32_.f29 = rhs37_
            elif source9_.is_DC28:
                d_461___mcc_h14_ = source9_.cf55
                d_462_cf55_ = d_461___mcc_h14_
                index33_ = default__.safeIndex(549, (d_268_v9_).length(0))
                (d_268_v9_)[index33_] = (d_416_v117_) >= ((d_406_v110_).cardinality)
                index34_ = default__.safeIndex(549, (d_268_v9_).length(0))
                (d_268_v9_)[index34_] = ((d_419_v120_).f32) != ((d_419_v120_).f32)
                d_463_v146_: _dafny.Array
                nw53_ = _dafny.Array(_dafny.CodePoint('D'), 4)
                d_463_v146_ = nw53_
                d_464_v147_: C0
                nw54_ = C0()
                nw54_.ctor__(d_463_v146_)
                d_464_v147_ = nw54_
                d_465_v148_: _dafny.Map
                d_465_v148_ = _dafny.Map({((d_301_v36_).f41) + ((d_301_v36_).f41): (d_404_v109_).fm6(d_271_globalState_)})
                d_466_v149_: _dafny.Seq
                d_466_v149_ = _dafny.SeqWithoutIsStrInference([(d_301_v36_).f42])
                d_467_v150_: D5
                d_467_v150_ = D5_DC13(((d_301_v36_).f42).ispropersubset((d_466_v149_)[default__.safeIndex((d_419_v120_).f34, len(d_466_v149_))]), d_418_v119_, _dafny.CodePoint('e'))
                d_468_v151_: _dafny.Map
                d_468_v151_ = _dafny.Map({default__.safeModuloInt((d_419_v120_).f30, (d_419_v120_).f30): d_464_v147_})
                rhs38_ = ((d_468_v151_)[((d_301_v36_).f41 if d_419_v120_.f31 else (d_419_v120_).f30)] if (((d_301_v36_).f41 if d_419_v120_.f31 else (d_419_v120_).f30)) in (d_468_v151_) else d_464_v147_)
                rhs39_ = (d_465_v148_).set(d_416_v117_, (d_408_v112_.f28) and (d_419_v120_.f29))
                rhs40_ = not((_dafny.Set({d_419_v120_.f35, (d_462_cf55_).f43, d_419_v120_.f29, d_408_v112_.f28})).ispropersubset((d_343_v70_) - (default__.fm45((d_301_v36_).f41, d_406_v110_, d_408_v112_.f29, d_260_v1_, d_271_globalState_))))
                rhs41_ = (((d_301_v36_).f41) - (len(d_417_v118_)) if ((d_301_v36_).fm20((d_419_v120_).f30, (d_268_v9_)[default__.safeIndex(549, (d_268_v9_).length(0))], d_271_globalState_)) == ((d_301_v36_).f41) else d_260_v1_)
                rhs42_ = d_467_v150_
                lhs33_ = d_271_globalState_
                lhs34_ = d_271_globalState_
                d_464_v147_ = rhs38_
                d_465_v148_ = rhs39_
                lhs33_.f26 = rhs40_
                lhs34_.f13 = rhs41_
                d_467_v150_ = rhs42_
                d_469_v152_: _dafny.Array
                nw55_ = _dafny.Array(False, 11)
                d_469_v152_ = nw55_
                d_470_v153_: _dafny.Map
                d_470_v153_ = _dafny.Map({d_469_v152_: (d_268_v9_)[default__.safeIndex(549, (d_268_v9_).length(0))]})
                index35_ = default__.safeIndex(549, (d_268_v9_).length(0))
                (d_268_v9_)[index35_] = ((0) - (default__.safeModuloInt((d_301_v36_).f41, d_260_v1_))) >= (default__.safeModuloInt(433, len(d_470_v153_)))
                index36_ = default__.safeIndex(251, (d_267_v8_).length(0))
                (d_267_v8_)[index36_] = d_259_v0_
                d_471_v154_: _dafny.Seq
                d_471_v154_ = _dafny.SeqWithoutIsStrInference([d_259_v0_, d_259_v0_, d_259_v0_])
                index37_ = default__.safeIndex(251, (d_267_v8_).length(0))
                (d_267_v8_)[index37_] = (d_471_v154_)[default__.safeIndex((d_301_v36_).f41, len(d_471_v154_))]
            elif True:
                d_472___mcc_h15_ = source9_.cf63
                d_473_cf63_ = d_472___mcc_h15_
                d_406_v110_ = _dafny.MultiSet([(_dafny.Set({d_408_v112_.f29})).issubset(d_343_v70_), True])
                (d_271_globalState_).f13 = 481
                d_474_v155_: int
                out29_: int
                out29_ = default__.m0(_dafny.SeqWithoutIsStrInference([(d_419_v120_).f34 for d_475_i15_ in range(default__.abs(555))]), not(d_408_v112_.f28), (_dafny.MultiSet([d_408_v112_.f29, d_419_v120_.f35])).cardinality, d_419_v120_.f29, d_271_globalState_)
                d_474_v155_ = out29_
                d_476_v156_: _dafny.Set
                d_476_v156_ = _dafny.Set({(d_419_v120_).f30, d_416_v117_, (d_419_v120_).f30, d_260_v1_, len(_dafny.Map({d_408_v112_.f28: d_404_v109_}))})
                d_477_v157_: T2
                nw56_ = C6()
                nw56_.ctor__(d_414_v115_, len((d_265_v6_) + (d_417_v118_)), (_dafny.Set({(d_301_v36_).f41, 313})).issubset(d_476_v156_), d_339_v66_, (d_419_v120_).f33, d_416_v117_, d_419_v120_.f35, ((d_419_v120_).f32)[default__.safeIndex(d_416_v117_, len((d_419_v120_).f32))], d_408_v112_.f28)
                d_477_v157_ = nw56_
                index38_ = default__.safeIndex(179, (d_259_v0_).length(0))
                rhs43_ = d_474_v155_
                rhs44_ = False
                rhs45_ = d_477_v157_
                rhs46_ = d_477_v157_.f29
                lhs35_ = d_259_v0_
                lhs36_ = default__.safeIndex(179, (d_259_v0_).length(0))
                lhs37_ = d_271_globalState_
                lhs38_ = d_477_v157_
                lhs35_[lhs36_] = rhs43_
                lhs37_.f12 = rhs44_
                d_477_v157_ = rhs45_
                lhs38_.f31 = rhs46_
            if False:
                d_478_v158_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_478_v158_ = nw57_
                d_479_v159_: C0
                nw58_ = C0()
                nw58_.ctor__(d_478_v158_)
                d_479_v159_ = nw58_
                d_480_v160_: _dafny.Seq
                d_480_v160_ = _dafny.SeqWithoutIsStrInference([d_479_v159_, d_479_v159_])
                rhs47_ = (0) - (((d_301_v36_).fm20(len(d_480_v160_), d_419_v120_.f28, d_271_globalState_)) - ((d_301_v36_).f41))
                rhs48_ = (0) - ((d_419_v120_).f34)
                lhs39_ = d_271_globalState_
                lhs39_.f13 = rhs47_
                d_260_v1_ = rhs48_
                d_481_v161_: D1
                d_481_v161_ = D1_DC3(d_419_v120_.f31, d_416_v117_, d_479_v159_, default__.fm32(d_408_v112_.f29, d_271_globalState_))
                (d_271_globalState_).f13 = (d_481_v161_).cf4
                d_482_v162_: _dafny.Set
                d_482_v162_ = _dafny.Set({d_261_v2_, _dafny.SeqWithoutIsStrInference([d_260_v1_, (d_419_v120_).f34])})
                d_483_v163_: D16
                d_483_v163_ = D16_DC39((d_482_v162_) - (d_482_v162_))
                d_483_v163_ = D16_DC39(_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_419_v120_).f34 for d_484_i16_ in range(default__.abs(405))])}))
                (d_419_v120_).f31 = (d_419_v120_).fm6(d_271_globalState_)
                d_485_v164_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Seq({}), 3)
                d_485_v164_ = nw59_
                d_486_v165_: C3
                nw60_ = C3()
                nw60_.ctor__(len(d_261_v2_), d_485_v164_, _dafny.CodePoint('d'), (0) - ((d_301_v36_).f41), d_419_v120_.f28, (d_419_v120_).f32, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_419_v120_.f28, d_419_v120_.f31]) for d_487_i17_ in range(default__.abs(18))]), (len((d_419_v120_).f32)) + ((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))]), d_408_v112_.f28, not(d_419_v120_.f29), (d_260_v1_) != ((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))]))
                d_486_v165_ = nw60_
            elif True:
                d_488_v166_: C10
                nw61_ = C10()
                nw61_.ctor__((d_305_v40_ if (d_301_v36_).fm19(d_271_globalState_) else (d_417_v118_)[default__.safeIndex((d_419_v120_).f34, len(d_417_v118_))]), (d_419_v120_).f30, d_408_v112_.f28, (d_419_v120_).f32, (d_419_v120_).f33, (d_419_v120_).f34, d_408_v112_.f29, d_408_v112_.f29, ((d_301_v36_).f42) == (default__.fm25(d_271_globalState_)))
                d_488_v166_ = nw61_
                rhs49_ = (((d_301_v36_).f42).intersection((d_301_v36_).f42)).ispropersubset((d_301_v36_).f42)
                rhs50_ = (d_419_v120_.f31 if d_419_v120_.f29 else default__.fm1(d_263_v4_, 419, d_271_globalState_))
                rhs51_ = ((d_301_v36_).f41) >= ((d_301_v36_).f41)
                rhs52_ = d_488_v166_
                lhs40_ = d_271_globalState_
                lhs41_ = d_271_globalState_
                lhs42_ = d_408_v112_
                lhs40_.f1 = rhs49_
                lhs41_.f12 = rhs50_
                lhs42_.f28 = rhs51_
                d_488_v166_ = rhs52_
                d_489_v167_: C8
                nw62_ = C8()
                nw62_.ctor__(d_419_v120_.f29, default__.fm48(d_419_v120_.f28, d_419_v120_.f29, (d_419_v120_).f30, d_271_globalState_), ((d_301_v36_).f41) + ((d_301_v36_).f41), d_408_v112_.f28, d_339_v66_, (d_403_v108_) + ((d_419_v120_).f33), (d_301_v36_).f41, d_408_v112_.f28, (d_306_v41_).ispropersubset(d_306_v41_), False)
                d_489_v167_ = nw62_
                d_490_v168_: T4
                nw63_ = C8()
                nw63_.ctor__(d_408_v112_.f28, d_414_v115_, len(_dafny.Map({d_419_v120_.f29: d_419_v120_.f31})), d_419_v120_.f29, (d_419_v120_).f32, (d_403_v108_) + ((d_419_v120_).f33), (d_301_v36_).f41, d_419_v120_.f28, ((d_419_v120_).f30) >= ((d_301_v36_).f41), d_419_v120_.f31)
                d_490_v168_ = nw63_
                d_490_v168_ = d_490_v168_
                d_491_v169_: _dafny.Map
                d_491_v169_ = _dafny.Map({d_419_v120_.f35: (d_490_v168_).fm7((d_301_v36_).f42, (d_419_v120_).f34, d_271_globalState_)})
                d_492_v170_: _dafny.Map
                d_492_v170_ = _dafny.Map({d_419_v120_.f28: (d_491_v169_) | (d_491_v169_)})
                d_493_v171_: _dafny.Array
                nw64_ = _dafny.Array(_dafny.Seq({}), 12)
                d_493_v171_ = nw64_
                d_494_v172_: _dafny.Array
                def lambda13_(d_495_v118_):
                    def lambda14_(d_496_i18_):
                        return d_495_v118_

                    return lambda14_

                init7_ = lambda13_(d_417_v118_)
                nw65_ = _dafny.Array(None, 4)
                for i0_7_ in range(nw65_.length(0)):
                    nw65_[i0_7_] = init7_(i0_7_)
                d_494_v172_ = nw65_
                index39_ = default__.safeIndex(985, (d_494_v172_).length(0))
                (d_494_v172_)[index39_] = (d_419_v120_).fm3(d_419_v120_.f35, not(d_419_v120_.f31), d_271_globalState_)
                d_497_v173_: _dafny.Map
                d_497_v173_ = _dafny.Map({d_304_v39_: True})
                d_498_v175_: _dafny.Array
                nw66_ = _dafny.Array(None, 15)
                nw66_[int(0)] = d_305_v40_
                nw66_[int(1)] = d_414_v115_
                nw66_[int(2)] = d_414_v115_
                nw66_[int(3)] = d_305_v40_
                nw66_[int(4)] = d_490_v168_.f36
                nw66_[int(5)] = d_305_v40_
                nw66_[int(6)] = d_414_v115_
                nw66_[int(7)] = d_414_v115_
                nw66_[int(8)] = d_490_v168_.f36
                nw66_[int(9)] = d_490_v168_.f36
                nw66_[int(10)] = d_490_v168_.f36
                nw66_[int(11)] = d_490_v168_.f36
                nw66_[int(12)] = d_414_v115_
                nw66_[int(13)] = d_305_v40_
                nw66_[int(14)] = d_305_v40_
                d_498_v175_ = nw66_
                d_499_v176_: C0
                nw67_ = C0()
                nw67_.ctor__(d_498_v175_)
                d_499_v176_ = nw67_
                d_500_v177_: _dafny.Map
                d_500_v177_ = _dafny.Map({default__.fm39(d_419_v120_.f35, (d_301_v36_).f42, d_419_v120_.f28, d_271_globalState_): d_499_v176_})
                d_501_v178_: _dafny.Map
                d_501_v178_ = _dafny.Map({len(d_500_v177_): d_419_v120_.f29})
                index40_ = default__.safeIndex(985, (d_494_v172_).length(0))
                def iife31_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(783, 265):
                        d_502_v174_: int = compr_21_
                        if ((783) <= (d_502_v174_)) and ((d_502_v174_) < (265)):
                            coll21_[default__.safeDivisionInt(d_502_v174_, (d_490_v168_).f34)] = (d_490_v168_).f30
                    return _dafny.Map(coll21_)
                def iife32_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in _dafny.IntegerRange(783, 265):
                        d_503_v174_: int = compr_22_
                        if ((783) <= (d_503_v174_)) and ((d_503_v174_) < (265)):
                            coll22_[default__.safeDivisionInt(d_503_v174_, (d_490_v168_).f34)] = (d_490_v168_).f30
                    return _dafny.Map(coll22_)
                rhs53_ = default__.fm67(d_271_globalState_)
                rhs54_ = not((((d_497_v173_)[iife31_()
                ] if (iife32_()
                ) in (d_497_v173_) else d_408_v112_.f29) if d_419_v120_.f31 else ((d_501_v178_)[(d_490_v168_).f30] if ((d_490_v168_).f30) in (d_501_v178_) else d_408_v112_.f29)))
                rhs55_ = d_493_v171_
                rhs56_ = d_417_v118_
                rhs57_ = ((d_417_v118_ if not(d_408_v112_.f28) else d_417_v118_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqm")) if d_490_v168_.f29 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xb"))))
                lhs43_ = d_419_v120_
                lhs44_ = d_494_v172_
                lhs45_ = default__.safeIndex(985, (d_494_v172_).length(0))
                d_492_v170_ = rhs53_
                lhs43_.f29 = rhs54_
                d_493_v171_ = rhs55_
                d_417_v118_ = rhs56_
                lhs44_[lhs45_] = rhs57_
                index41_ = default__.safeIndex(122, (d_268_v9_).length(0))
                (d_268_v9_)[index41_] = False
                index42_ = default__.safeIndex(122, (d_268_v9_).length(0))
                (d_268_v9_)[index42_] = ((d_419_v120_).f32)[default__.safeIndex((d_301_v36_).f41, len((d_419_v120_).f32))]
            d_504_v179_: C1
            nw68_ = C1()
            nw68_.ctor__(d_419_v120_.f35, d_419_v120_.f29, d_408_v112_.f29)
            d_504_v179_ = nw68_
            d_505_v180_: D10
            d_505_v180_ = D10_DC28(d_504_v179_)
            pat_let_tv3_ = d_504_v179_
            d_506_v181_: _dafny.Array
            nw69_ = _dafny.Array(None, 26)
            nw69_[int(0)] = d_505_v180_
            nw69_[int(1)] = D10_DC28(d_504_v179_)
            nw69_[int(2)] = d_505_v180_
            nw69_[int(3)] = D10_DC28(d_504_v179_)
            nw69_[int(4)] = D10_DC28(d_504_v179_)
            nw69_[int(5)] = d_505_v180_
            nw69_[int(6)] = d_505_v180_
            nw69_[int(7)] = d_505_v180_
            nw69_[int(8)] = d_505_v180_
            nw69_[int(9)] = d_505_v180_
            def iife33_(_pat_let5_0):
                def iife34_(d_507_dt__update__tmp_h2_):
                    def iife35_(_pat_let6_0):
                        def iife36_(d_508_dt__update_hcf55_h0_):
                            return D10_DC28(d_508_dt__update_hcf55_h0_)
                        return iife36_(_pat_let6_0)
                    return iife35_(pat_let_tv3_)
                return iife34_(_pat_let5_0)
            nw69_[int(10)] = iife33_(d_505_v180_)
            nw69_[int(11)] = d_505_v180_
            nw69_[int(12)] = d_505_v180_
            nw69_[int(13)] = d_505_v180_
            nw69_[int(14)] = d_505_v180_
            nw69_[int(15)] = D10_DC28((d_505_v180_).cf55)
            nw69_[int(16)] = d_505_v180_
            nw69_[int(17)] = d_505_v180_
            nw69_[int(18)] = d_505_v180_
            nw69_[int(19)] = d_505_v180_
            nw69_[int(20)] = d_505_v180_
            nw69_[int(21)] = d_505_v180_
            nw69_[int(22)] = d_505_v180_
            nw69_[int(23)] = d_505_v180_
            nw69_[int(24)] = D10_DC28(d_504_v179_)
            nw69_[int(25)] = d_505_v180_
            d_506_v181_ = nw69_
            d_509_v182_: D20
            d_509_v182_ = D20_DC48((d_419_v120_).f30, d_506_v181_, d_419_v120_.f29)
            d_509_v182_ = d_509_v182_
        elif True:
            d_510_v183_: D3
            d_510_v183_ = D3_DC9(False, not(d_263_v4_), d_268_v9_, d_263_v4_)
            d_511_v184_: _dafny.Array
            d_512_v185_: int
            out30_: _dafny.Array
            out31_: int
            out30_, out31_ = (d_301_v36_).m8((d_510_v183_).cf20, d_271_globalState_)
            d_511_v184_ = out30_
            d_512_v185_ = out31_
            d_513_v186_: _dafny.Array
            nw70_ = _dafny.Array(None, 27)
            d_513_v186_ = nw70_
            d_514_v187_: T3
            nw71_ = C5()
            nw71_.ctor__(not(True), d_260_v1_, d_512_v185_, d_263_v4_, d_339_v66_, d_403_v108_, (d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], d_263_v4_, d_263_v4_, d_263_v4_)
            d_514_v187_ = nw71_
            index43_ = default__.safeIndex(899, (d_513_v186_).length(0))
            (d_513_v186_)[index43_] = d_514_v187_
            index44_ = default__.safeIndex(899, (d_513_v186_).length(0))
            (d_513_v186_)[index44_] = d_514_v187_
            d_515_v188_: _dafny.Seq
            d_515_v188_ = _dafny.SeqWithoutIsStrInference([d_404_v109_, d_404_v109_, d_404_v109_, d_404_v109_, d_404_v109_])
            d_516_v189_: D17
            d_516_v189_ = D17_DC42((d_515_v188_) + (d_515_v188_))
            d_517_v190_: _dafny.Map
            d_517_v190_ = _dafny.Map({d_514_v187_.f31: d_515_v188_})
            d_516_v189_ = D17_DC42(((d_517_v190_)[(d_514_v187_).fm6(d_271_globalState_)] if ((d_514_v187_).fm6(d_271_globalState_)) in (d_517_v190_) else d_515_v188_))
            d_518_v191_: _dafny.Array
            d_519_v192_: int
            out32_: _dafny.Array
            out33_: int
            out32_, out33_ = (d_301_v36_).m8(d_268_v9_, d_271_globalState_)
            d_518_v191_ = out32_
            d_519_v192_ = out33_
            d_520_v193_: _dafny.Array
            def lambda15_(d_521_v40_):
                def lambda16_(d_522_i19_):
                    return d_521_v40_

                return lambda16_

            init8_ = lambda15_(d_305_v40_)
            nw72_ = _dafny.Array(None, 8)
            for i0_8_ in range(nw72_.length(0)):
                nw72_[i0_8_] = init8_(i0_8_)
            d_520_v193_ = nw72_
            d_523_v194_: C0
            nw73_ = C0()
            nw73_.ctor__(d_520_v193_)
            d_523_v194_ = nw73_
            d_524_v195_: D1
            d_524_v195_ = D1_DC3(d_514_v187_.f28, (0) - (d_512_v185_), d_523_v194_, (d_261_v2_)[default__.safeIndex((d_301_v36_).f41, len(d_261_v2_))])
            d_525_v196_: C1
            nw74_ = C1()
            nw74_.ctor__(d_514_v187_.f28, d_514_v187_.f28, not((d_524_v195_).cf3))
            d_525_v196_ = nw74_
        d_526_v197_: _dafny.Seq
        d_526_v197_ = _dafny.SeqWithoutIsStrInference([d_342_v69_])
        if (_dafny.MultiSet([d_526_v197_])) == (_dafny.MultiSet([d_526_v197_, d_526_v197_, d_526_v197_])):
            d_304_v39_ = (d_304_v39_).set((d_301_v36_).f41, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fy"))))
            d_527_v198_: T1
            nw75_ = C7()
            nw75_.ctor__((((d_301_v36_).f42)[-183] if (-183) in ((d_301_v36_).f42) else d_260_v1_), (d_301_v36_).f42, len(d_343_v70_), d_263_v4_, d_263_v4_, d_263_v4_)
            d_527_v198_ = nw75_
            d_528_v199_: _dafny.Seq
            d_528_v199_ = _dafny.SeqWithoutIsStrInference([d_527_v198_, d_527_v198_, d_527_v198_, d_527_v198_])
            d_529_v200_: _dafny.Map
            d_529_v200_ = _dafny.Map({d_404_v109_: d_528_v199_})
            d_530_v201_: _dafny.Set
            d_530_v201_ = _dafny.Set({len(d_265_v6_)})
            d_531_v202_: _dafny.Seq
            d_531_v202_ = _dafny.SeqWithoutIsStrInference([d_528_v199_, d_528_v199_])
            d_532_v203_: _dafny.Seq
            d_532_v203_ = d_528_v199_
            d_533_v204_: _dafny.Array
            nw76_ = _dafny.Array(None, 27)
            nw76_[int(0)] = (d_528_v199_ if d_263_v4_ else d_528_v199_)
            nw76_[int(1)] = d_528_v199_
            nw76_[int(2)] = d_528_v199_
            nw76_[int(3)] = (d_528_v199_) + (d_528_v199_)
            nw76_[int(4)] = (d_528_v199_) + (d_528_v199_)
            nw76_[int(5)] = (d_528_v199_) + (d_528_v199_)
            nw76_[int(6)] = (d_528_v199_) + (_dafny.SeqWithoutIsStrInference([d_527_v198_]))
            nw76_[int(7)] = d_528_v199_
            nw76_[int(8)] = (d_528_v199_) + (d_528_v199_)
            nw76_[int(9)] = ((d_528_v199_).set(default__.safeIndex(473, len(d_528_v199_)), d_527_v198_)) + (d_528_v199_)
            nw76_[int(10)] = d_528_v199_
            nw76_[int(11)] = d_528_v199_
            nw76_[int(12)] = _dafny.SeqWithoutIsStrInference([d_527_v198_])
            nw76_[int(13)] = (d_528_v199_).set(default__.safeIndex((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))], len(d_528_v199_)), d_527_v198_)
            nw76_[int(14)] = (d_528_v199_) + (d_528_v199_)
            nw76_[int(15)] = d_528_v199_
            nw76_[int(16)] = d_528_v199_
            nw76_[int(17)] = d_528_v199_
            nw76_[int(18)] = (d_528_v199_) + (((d_529_v200_)[d_404_v109_] if (d_404_v109_) in (d_529_v200_) else d_528_v199_))
            nw76_[int(19)] = _dafny.SeqWithoutIsStrInference([d_527_v198_, d_527_v198_, d_527_v198_, d_527_v198_, d_527_v198_])
            nw76_[int(20)] = (d_528_v199_) + (((d_528_v199_).set(default__.safeIndex(len(d_530_v201_), len(d_528_v199_)), d_527_v198_)).set(default__.safeIndex((d_301_v36_).f41, len((d_528_v199_).set(default__.safeIndex(len(d_530_v201_), len(d_528_v199_)), d_527_v198_))), d_527_v198_))
            nw76_[int(21)] = (d_531_v202_)[default__.safeIndex((((d_301_v36_).f42)[576] if (576) in ((d_301_v36_).f42) else -151), len(d_531_v202_))]
            nw76_[int(22)] = (d_532_v203_)
            nw76_[int(23)] = _dafny.SeqWithoutIsStrInference([d_527_v198_])
            nw76_[int(24)] = d_528_v199_
            nw76_[int(25)] = (d_528_v199_) + (d_528_v199_)
            nw76_[int(26)] = (d_528_v199_) + (d_528_v199_)
            d_533_v204_ = nw76_
            index45_ = default__.safeIndex(187, (d_533_v204_).length(0))
            (d_533_v204_)[index45_] = (d_528_v199_) + (_dafny.SeqWithoutIsStrInference([d_527_v198_]))
            index46_ = default__.safeIndex(187, (d_533_v204_).length(0))
            (d_533_v204_)[index46_] = ((d_528_v199_ if d_527_v198_.f28 else d_528_v199_)) + (d_528_v199_)
            d_534_v205_: str
            d_535_v206_: _dafny.Seq
            d_536_v207_: int
            d_537_v208_: _dafny.Seq
            out34_: str
            out35_: _dafny.Seq
            out36_: int
            out37_: _dafny.Seq
            out34_, out35_, out36_, out37_ = (d_301_v36_).m7(d_527_v198_.f31, d_268_v9_, d_527_v198_.f29, d_271_globalState_)
            d_534_v205_ = out34_
            d_535_v206_ = out35_
            d_536_v207_ = out36_
            d_537_v208_ = out37_
            d_538_v209_: C1
            nw77_ = C1()
            nw77_.ctor__(not((d_301_v36_).fm19(d_271_globalState_)), False, d_527_v198_.f28)
            d_538_v209_ = nw77_
            d_539_v210_: D10
            d_539_v210_ = D10_DC28(d_538_v209_)
            pat_let_tv4_ = d_538_v209_
            d_540_v211_: _dafny.Array
            nw78_ = _dafny.Array(None, 17)
            nw78_[int(0)] = d_539_v210_
            nw78_[int(1)] = d_539_v210_
            nw78_[int(2)] = d_539_v210_
            nw78_[int(3)] = d_539_v210_
            nw78_[int(4)] = d_539_v210_
            nw78_[int(5)] = d_539_v210_
            nw78_[int(6)] = d_539_v210_
            nw78_[int(7)] = d_539_v210_
            def iife37_(_pat_let7_0):
                def iife38_(d_541_dt__update__tmp_h3_):
                    def iife39_(_pat_let8_0):
                        def iife40_(d_542_dt__update_hcf55_h1_):
                            return D10_DC28(d_542_dt__update_hcf55_h1_)
                        return iife40_(_pat_let8_0)
                    return iife39_(pat_let_tv4_)
                return iife38_(_pat_let7_0)
            nw78_[int(8)] = iife37_(d_539_v210_)
            nw78_[int(9)] = d_539_v210_
            nw78_[int(10)] = d_539_v210_
            nw78_[int(11)] = d_539_v210_
            nw78_[int(12)] = d_539_v210_
            nw78_[int(13)] = d_539_v210_
            nw78_[int(14)] = d_539_v210_
            nw78_[int(15)] = d_539_v210_
            nw78_[int(16)] = d_539_v210_
            d_540_v211_ = nw78_
            d_543_v212_: D20
            d_543_v212_ = D20_DC50(D20_DC48((d_301_v36_).f41, d_540_v211_, (d_538_v209_).f43))
            source10_ = d_543_v212_
            if source10_.is_DC48:
                d_544___mcc_h16_ = source10_.cf82
                d_545___mcc_h17_ = source10_.cf83
                d_546___mcc_h18_ = source10_.cf84
                d_547_cf84_ = d_546___mcc_h18_
                d_548_cf83_ = d_545___mcc_h17_
                d_549_cf82_ = d_544___mcc_h16_
                d_550_v213_: bool
                d_551_v214_: int
                d_552_v215_: _dafny.Seq
                out38_: bool
                out39_: int
                out40_: _dafny.Seq
                out38_, out39_, out40_ = (d_404_v109_).m1(d_271_globalState_)
                d_550_v213_ = out38_
                d_551_v214_ = out39_
                d_552_v215_ = out40_
                d_551_v214_ = (default__.safeModuloInt((d_301_v36_).f41, 50)) - (len(d_339_v66_))
                d_553_v216_: _dafny.Seq
                d_553_v216_ = _dafny.SeqWithoutIsStrInference([d_343_v70_])
                d_343_v70_ = (d_553_v216_)[default__.safeIndex(d_536_v207_, len(d_553_v216_))]
                index47_ = default__.safeIndex(179, (d_259_v0_).length(0))
                (d_259_v0_)[index47_] = (d_301_v36_).f41
            elif source10_.is_DC49:
                d_554___mcc_h19_ = source10_.cf85
                d_555___mcc_h20_ = source10_.cf86
                d_556_cf86_ = d_555___mcc_h20_
                d_557_cf85_ = d_554___mcc_h19_
                (d_527_v198_).f31 = (d_538_v209_).f43
                d_558_v217_: bool
                d_559_v218_: int
                d_560_v219_: _dafny.Seq
                out41_: bool
                out42_: int
                out43_: _dafny.Seq
                out41_, out42_, out43_ = (d_404_v109_).m1(d_271_globalState_)
                d_558_v217_ = out41_
                d_559_v218_ = out42_
                d_560_v219_ = out43_
                (d_527_v198_).f29 = (d_265_v6_) == ((d_560_v219_ if d_263_v4_ else d_265_v6_))
                d_561_v220_: D4
                d_561_v220_ = D4_DC11((d_527_v198_).f30, d_557_cf85_, 854)
                d_562_v221_: _dafny.Map
                d_562_v221_ = _dafny.Map({d_534_v205_: (d_538_v209_).f43})
                d_563_v222_: C10
                nw79_ = C10()
                nw79_.ctor__(d_534_v205_, (d_561_v220_).cf25, (d_527_v198_.f29 if (d_538_v209_).fm22(d_559_v218_, default__.fm24(False, (d_301_v36_).f41, d_271_globalState_), d_527_v198_.f29, d_271_globalState_) else (d_538_v209_).f43), d_339_v66_, d_403_v108_, d_260_v1_, d_527_v198_.f31, d_527_v198_.f31, ((d_562_v221_)[d_556_cf86_] if (d_556_cf86_) in (d_562_v221_) else d_527_v198_.f28))
                d_563_v222_ = nw79_
            elif source10_.is_DC47:
                d_564___mcc_h21_ = source10_.cf81
                d_565_cf81_ = d_564___mcc_h21_
                d_566_v223_: T3
                nw80_ = C4()
                nw80_.ctor__(d_260_v1_, d_527_v198_.f31, d_339_v66_, d_403_v108_, (d_301_v36_).f41, (d_538_v209_).f43, d_527_v198_.f31, d_527_v198_.f29)
                d_566_v223_ = nw80_
                d_567_v224_: _dafny.Seq
                d_567_v224_ = _dafny.SeqWithoutIsStrInference([d_566_v223_])
                d_568_v225_: C10
                nw81_ = C10()
                nw81_.ctor__(d_305_v40_, len(d_567_v224_), d_566_v223_.f31, (d_566_v223_).f32, d_403_v108_, d_536_v207_, d_566_v223_.f28, d_263_v4_, d_566_v223_.f29)
                d_568_v225_ = nw81_
                d_569_v226_: _dafny.Array
                nw82_ = _dafny.Array(None, 20)
                nw82_[int(0)] = d_568_v225_
                nw82_[int(1)] = d_568_v225_
                nw82_[int(2)] = d_568_v225_
                nw82_[int(3)] = d_568_v225_
                nw82_[int(4)] = d_568_v225_
                nw82_[int(5)] = d_568_v225_
                nw82_[int(6)] = d_568_v225_
                nw82_[int(7)] = d_568_v225_
                nw82_[int(8)] = d_568_v225_
                nw82_[int(9)] = d_568_v225_
                nw82_[int(10)] = d_568_v225_
                nw82_[int(11)] = d_568_v225_
                nw82_[int(12)] = d_568_v225_
                nw82_[int(13)] = d_568_v225_
                nw82_[int(14)] = d_568_v225_
                nw82_[int(15)] = d_568_v225_
                nw82_[int(16)] = d_568_v225_
                nw82_[int(17)] = d_568_v225_
                nw82_[int(18)] = d_568_v225_
                nw82_[int(19)] = d_568_v225_
                d_569_v226_ = nw82_
                d_569_v226_ = d_569_v226_
                (d_271_globalState_).f13 = (d_566_v223_).f30
                d_570_v227_: _dafny.Map
                d_570_v227_ = _dafny.Map({d_566_v223_.f28: len(_dafny.SeqWithoutIsStrInference([d_305_v40_ for d_571_i20_ in range(default__.abs(834))]))})
                d_572_v228_: _dafny.Set
                d_572_v228_ = _dafny.Set({d_527_v198_})
                d_573_v229_: _dafny.Map
                d_573_v229_ = _dafny.Map({-925: default__.fm41(d_566_v223_.f31, d_305_v40_, d_527_v198_.f28, len(d_572_v228_), d_271_globalState_)})
                d_574_v230_: T2
                nw83_ = C6()
                nw83_.ctor__(d_305_v40_, 925, d_566_v223_.f35, _dafny.SeqWithoutIsStrInference([d_566_v223_.f31]), (d_566_v223_).f33, ((d_570_v227_)[False] if (False) in (d_570_v227_) else len(d_573_v229_)), d_566_v223_.f31, False, d_527_v198_.f31)
                d_574_v230_ = nw83_
                d_575_v231_: _dafny.Map
                d_575_v231_ = _dafny.Map({d_574_v230_: (d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))]})
                d_576_v232_: _dafny.Seq
                d_576_v232_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_574_v230_: ((d_301_v36_).f42).cardinality}), d_575_v231_])
                d_577_v233_: C6
                nw84_ = C6()
                nw84_.ctor__(d_305_v40_, default__.safeModuloInt(723, len(d_570_v227_)), (d_530_v201_).issubset(d_530_v201_), d_339_v66_, default__.fm51(d_527_v198_.f31, d_305_v40_, d_271_globalState_), d_260_v1_, d_527_v198_.f28, d_527_v198_.f29, (d_575_v231_) in (d_576_v232_))
                d_577_v233_ = nw84_
                index48_ = default__.safeIndex(365, (d_268_v9_).length(0))
                (d_268_v9_)[index48_] = not(d_566_v223_.f31)
                index49_ = default__.safeIndex(365, (d_268_v9_).length(0))
                (d_268_v9_)[index49_] = d_574_v230_.f31
            elif True:
                d_578___mcc_h22_ = source10_.cf87
                d_579_cf87_ = d_578___mcc_h22_
                d_580_v234_: _dafny.Array
                def lambda17_(d_581_i21_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtf"))

                init9_ = lambda17_
                nw85_ = _dafny.Array(None, 15)
                for i0_9_ in range(nw85_.length(0)):
                    nw85_[i0_9_] = init9_(i0_9_)
                d_580_v234_ = nw85_
                d_580_v234_ = d_580_v234_
                (d_271_globalState_).f13 = (d_536_v207_) + (((0) - ((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))]) if d_527_v198_.f31 else d_536_v207_))
                d_536_v207_ = (0) - (default__.safeDivisionInt((0) - ((d_527_v198_).f30), (d_301_v36_).f41))
                d_582_v235_: _dafny.Map
                d_582_v235_ = _dafny.Map({d_530_v201_: d_265_v6_})
                (d_271_globalState_).f13 = default__.fm32(not((d_582_v235_) == (d_582_v235_)), d_271_globalState_)
            d_583_v236_: _dafny.Array
            def lambda18_(d_584_v198_):
                def lambda19_(d_585_i22_):
                    return (_dafny.Map({d_584_v198_.f31: not(d_584_v198_.f29)})) | (_dafny.Map({d_584_v198_.f28: not(False)}))

                return lambda19_

            init10_ = lambda18_(d_527_v198_)
            nw86_ = _dafny.Array(None, 22)
            for i0_10_ in range(nw86_.length(0)):
                nw86_[i0_10_] = init10_(i0_10_)
            d_583_v236_ = nw86_
            d_586_v237_: D22
            d_586_v237_ = D22_DC55((d_527_v198_).f30, d_260_v1_, (d_538_v209_).f43)
            d_587_v238_: _dafny.Map
            d_587_v238_ = _dafny.Map({d_527_v198_.f29: (d_586_v237_).cf94})
            index50_ = default__.safeIndex(147, (d_583_v236_).length(0))
            (d_583_v236_)[index50_] = d_587_v238_
            d_588_v239_: T4
            nw87_ = C8()
            nw87_.ctor__(False, d_534_v205_, 694, not((d_538_v209_).f43), (d_339_v66_).set(default__.safeIndex(d_260_v1_, len(d_339_v66_)), (d_538_v209_).f43), _dafny.SeqWithoutIsStrInference([d_339_v66_]), (d_301_v36_).f41, (d_538_v209_).f43, d_527_v198_.f31, False)
            d_588_v239_ = nw87_
            d_589_v240_: _dafny.Map
            d_589_v240_ = _dafny.Map({d_527_v198_.f31: d_588_v239_})
            d_590_v241_: _dafny.Array
            nw88_ = _dafny.Array(None, 6)
            nw88_[int(0)] = (d_589_v240_).set(d_588_v239_.f31, d_588_v239_)
            nw88_[int(1)] = (d_589_v240_) | (d_589_v240_)
            nw88_[int(2)] = (d_589_v240_ if d_527_v198_.f28 else d_589_v240_)
            nw88_[int(3)] = d_589_v240_
            nw88_[int(4)] = _dafny.Map({d_588_v239_.f31: d_588_v239_})
            nw88_[int(5)] = d_589_v240_
            d_590_v241_ = nw88_
            d_591_v242_: _dafny.Array
            nw89_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_591_v242_ = nw89_
            d_592_v243_: _dafny.Map
            d_592_v243_ = _dafny.Map({d_591_v242_: not(not (d_588_v239_.f28) or (d_588_v239_.f35))})
            index51_ = default__.safeIndex(147, (d_583_v236_).length(0))
            index52_ = default__.safeIndex(179, (d_259_v0_).length(0))
            rhs58_ = d_587_v238_
            rhs59_ = d_590_v241_
            rhs60_ = ((d_588_v239_).f34) > ((len(d_530_v201_) if d_263_v4_ else 797))
            rhs61_ = ((d_301_v36_).f41) + ((d_259_v0_)[default__.safeIndex(179, (d_259_v0_).length(0))])
            rhs62_ = ((d_592_v243_)[d_591_v242_] if (d_591_v242_) in (d_592_v243_) else d_588_v239_.f29)
            lhs46_ = d_583_v236_
            lhs47_ = default__.safeIndex(147, (d_583_v236_).length(0))
            lhs48_ = d_271_globalState_
            lhs49_ = d_259_v0_
            lhs50_ = default__.safeIndex(179, (d_259_v0_).length(0))
            lhs51_ = d_271_globalState_
            lhs46_[lhs47_] = rhs58_
            d_590_v241_ = rhs59_
            lhs48_.f1 = rhs60_
            lhs49_[lhs50_] = rhs61_
            lhs51_.f12 = rhs62_
        elif True:
            rhs63_ = d_259_v0_
            rhs64_ = (d_301_v36_).fm19(d_271_globalState_)
            lhs52_ = d_271_globalState_
            d_259_v0_ = rhs63_
            lhs52_.f1 = rhs64_
            d_593_v244_: _dafny.Set
            d_593_v244_ = _dafny.Set({len(_dafny.Map({_dafny.Set({False}): (d_301_v36_).f41})), len(d_265_v6_)})
            (d_271_globalState_).f1 = not((len(d_593_v244_)) < ((d_301_v36_).f41))
            index53_ = default__.safeIndex(179, (d_259_v0_).length(0))
            (d_259_v0_)[index53_] = ((d_301_v36_).f41) + ((d_301_v36_).f41)
            index54_ = default__.safeIndex(159, (d_268_v9_).length(0))
            (d_268_v9_)[index54_] = d_263_v4_
            index55_ = default__.safeIndex(159, (d_268_v9_).length(0))
            (d_268_v9_)[index55_] = not((not (d_263_v4_) or (d_263_v4_) if d_263_v4_ else d_263_v4_))
            if (d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))]:
                (d_271_globalState_).f1 = (d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))]
                index56_ = default__.safeIndex(179, (d_259_v0_).length(0))
                (d_259_v0_)[index56_] = 501
                d_594_v245_: C1
                nw90_ = C1()
                nw90_.ctor__((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], (d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], d_263_v4_)
                d_594_v245_ = nw90_
                d_260_v1_ = (d_301_v36_).f41
                d_595_v246_: D7
                d_595_v246_ = D7_DC20(D7_DC19((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], (d_301_v36_).f41, -249))
                d_596_v247_: _dafny.Map
                d_596_v247_ = _dafny.Map({d_595_v246_: d_339_v66_})
                d_597_v248_: D7
                d_597_v248_ = D7_DC19((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], (d_301_v36_).f41, (d_301_v36_).f41)
                d_596_v247_ = (d_596_v247_).set(D7_DC20(d_597_v248_), _dafny.SeqWithoutIsStrInference([d_263_v4_]))
            elif True:
                d_598_v249_: _dafny.Array
                def lambda20_(d_599_v9_):
                    def lambda21_(d_600_i23_):
                        return (d_599_v9_)[default__.safeIndex(159, (d_599_v9_).length(0))]

                    return lambda21_

                init11_ = lambda20_(d_268_v9_)
                nw91_ = _dafny.Array(None, 16)
                for i0_11_ in range(nw91_.length(0)):
                    nw91_[i0_11_] = init11_(i0_11_)
                d_598_v249_ = nw91_
                d_601_v250_: str
                d_602_v251_: _dafny.Seq
                d_603_v252_: int
                d_604_v253_: _dafny.Seq
                out44_: str
                out45_: _dafny.Seq
                out46_: int
                out47_: _dafny.Seq
                out44_, out45_, out46_, out47_ = (d_301_v36_).m7(d_263_v4_, d_598_v249_, (d_301_v36_).fm19(d_271_globalState_), d_271_globalState_)
                d_601_v250_ = out44_
                d_602_v251_ = out45_
                d_603_v252_ = out46_
                d_604_v253_ = out47_
                index57_ = default__.safeIndex(179, (d_259_v0_).length(0))
                rhs65_ = (d_301_v36_).f41
                rhs66_ = (d_265_v6_) + (d_604_v253_)
                rhs67_ = d_265_v6_
                rhs68_ = default__.safeDivisionInt((d_301_v36_).f41, default__.fm32((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], d_271_globalState_))
                lhs53_ = d_259_v0_
                lhs54_ = default__.safeIndex(179, (d_259_v0_).length(0))
                lhs53_[lhs54_] = rhs65_
                d_265_v6_ = rhs66_
                d_604_v253_ = rhs67_
                d_603_v252_ = rhs68_
                d_605_v254_: C5
                nw92_ = C5()
                nw92_.ctor__((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], default__.safeModuloInt(default__.fm24(False, (d_301_v36_).f41, d_271_globalState_), d_603_v252_), d_260_v1_, (d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], d_339_v66_, (d_403_v108_) + (d_403_v108_), (0) - (len(_dafny.Map({True: (d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))]}))), not(not((d_263_v4_) or (d_263_v4_))), (d_260_v1_) == ((d_301_v36_).f41), not ((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))]) or (False))
                d_605_v254_ = nw92_
                d_606_v255_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_606_v255_ = nw93_
                index58_ = default__.safeIndex(700, (d_606_v255_).length(0))
                (d_606_v255_)[index58_] = (d_305_v40_ if d_263_v4_ else d_305_v40_)
                d_607_v256_: _dafny.Map
                d_607_v256_ = _dafny.Map({d_404_v109_: (d_301_v36_).f41})
                index59_ = default__.safeIndex(700, (d_606_v255_).length(0))
                (d_606_v255_)[index59_] = (d_601_v250_ if (99) > ((d_605_v254_).f48) else default__.fm48((d_268_v9_)[default__.safeIndex(159, (d_268_v9_).length(0))], True, len(d_607_v256_), d_271_globalState_))
                (d_271_globalState_).f13 = -574
        d_608_v257_: _dafny.MultiSet
        d_608_v257_ = _dafny.MultiSet([d_268_v9_, d_268_v9_])
        index60_ = default__.safeIndex(179, (d_259_v0_).length(0))
        (d_259_v0_)[index60_] = (default__.safeModuloInt((d_608_v257_).cardinality, (d_301_v36_).f41)) + (329)
        _dafny.print(_dafny.string_of((d_259_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_260_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v2_) == (_dafny.SeqWithoutIsStrInference([20]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_262_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_263_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v5_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([20]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_265_v6_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m')]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssisbjr"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_271_globalState_).f5) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([20]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_271_globalState_).f7).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_271_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_271_globalState_).f18)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_.f21)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_271_globalState_).f22)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_271_globalState_).f25) == (_dafny.SeqWithoutIsStrInference([20]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_globalState_.f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_301_v36_).f41))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v36_).f42) == (_dafny.MultiSet([20]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_302_v37_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_303_v38_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v39_) == (_dafny.Map({20: 0, 0: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_305_v40_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v41_) == (_dafny.Set({_dafny.CodePoint('m')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v66_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_340_v67_).cf46))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_340_v67_).cf47))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_340_v67_).cf48) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_340_v67_).cf49))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_340_v67_).cf50) == (_dafny.SeqWithoutIsStrInference([20]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v68_) == (_dafny.MultiSet([_dafny.CodePoint('m')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v69_).cf15) == (_dafny.MultiSet([_dafny.CodePoint('m')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v69_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v69_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v70_) == (_dafny.Set({False, True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_344_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_403_v108_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_406_v110_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_526_v197_) == (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.MultiSet([_dafny.CodePoint('m')]), True, 417)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_608_v257_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
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
        return lambda: D1_DC3(False, int(0), None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0), False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(_dafny.MultiSet({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(False, _dafny.Map({}), _dafny.CodePoint('D'))
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

class D5_DC13(D5, NamedTuple('DC13', [('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({self.cf26.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC16(D6, NamedTuple('DC16', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC19(D7, NamedTuple('DC19', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC22(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC22(D8, NamedTuple('DC22', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27(None, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC27(D9, NamedTuple('DC27', [('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf53)}, {self.cf54.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D10_DC32)

class D10_DC29(D10, NamedTuple('DC29', [('cf56', Any), ('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC31(D10, NamedTuple('DC31', [('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC32(D10, NamedTuple('DC32', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC32({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC32) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)

class D11_DC33(D11, NamedTuple('DC33', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC34(D12, NamedTuple('DC34', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC35(D13, NamedTuple('DC35', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC37(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC37(D14, NamedTuple('DC37', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)

class D15_DC38(D15, NamedTuple('DC38', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)

class D16_DC40(D16, NamedTuple('DC40', [('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({self.cf73.VerbatimString(True)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC41(D16, NamedTuple('DC41', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC43()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)

class D17_DC43(D17, NamedTuple('DC43', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC44(D17, NamedTuple('DC44', [('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC42(D17, NamedTuple('DC42', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D18_DC45)

class D18_DC45(D18, NamedTuple('DC45', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC45({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC45) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)

class D19_DC46(D19, NamedTuple('DC46', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC48(int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D20_DC47)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)

class D20_DC48(D20, NamedTuple('DC48', [('cf82', Any), ('cf83', Any), ('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC49(D20, NamedTuple('DC49', [('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC47(D20, NamedTuple('DC47', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC47({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC47) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC50(D20, NamedTuple('DC50', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC52(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)

class D21_DC52(D21, NamedTuple('DC52', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC51(D21, NamedTuple('DC51', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC53(D21, NamedTuple('DC53', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC55(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D22_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D22_DC54)

class D22_DC55(D22, NamedTuple('DC55', [('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC55({_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC55) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC54(D22, NamedTuple('DC54', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC54({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC54) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D23_DC56)

class D23_DC56(D23, NamedTuple('DC56', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC56({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC56) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC58()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D24_DC58)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D24_DC59)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D24_DC57)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)

class D24_DC58(D24, NamedTuple('DC58', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC58'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC58)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC59(D24, NamedTuple('DC59', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC59({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC59) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC57(D24, NamedTuple('DC57', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC57({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC57) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC60(D24, NamedTuple('DC60', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D25_DC61)

class D25_DC61(D25, NamedTuple('DC61', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC61({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC61) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC63(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D26_DC63)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D26_DC62)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D26_DC64)

class D26_DC63(D26, NamedTuple('DC63', [('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC63({_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC63) and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC62(D26, NamedTuple('DC62', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC62({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC62) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC64(D26, NamedTuple('DC64', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC64({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC64) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D27_DC65)

class D27_DC65(D27, NamedTuple('DC65', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC65({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC65) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D28_DC66)

class D28_DC66(D28, NamedTuple('DC66', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC66({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC66) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC68(_dafny.Seq({}), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D29_DC68)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D29_DC67)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D29_DC69)

class D29_DC68(D29, NamedTuple('DC68', [('cf108', Any), ('cf109', Any), ('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC68({_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC68) and self.cf108 == __o.cf108 and self.cf109 == __o.cf109 and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC67(D29, NamedTuple('DC67', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC67({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC67) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC69(D29, NamedTuple('DC69', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC69({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC69) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value

class T1:
    pass
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value

class T2:
    pass
    def m1(self, globalState):
        pass


class T3:
    pass
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value

class T4:
    pass
    @property
    def f36(self):
        return self._f36
    @f36.setter
    def f36(self, value):
        self._f36 = value

class GlobalState:
    def  __init__(self):
        self.f1: bool = False
        self.f12: bool = False
        self.f13: int = int(0)
        self.f21: _dafny.Array = _dafny.Array(None, 0)
        self.f26: bool = False
        self._f0: int = int(0)
        self._f2: bool = False
        self._f3: bool = False
        self._f4: int = int(0)
        self._f5: _dafny.Map = _dafny.Map({})
        self._f6: int = int(0)
        self._f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: bool = False
        self._f10: int = int(0)
        self._f11: int = int(0)
        self._f14: bool = False
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        self._f16: int = int(0)
        self._f17: str = _dafny.CodePoint('D')
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f19: bool = False
        self._f20: bool = False
        self._f22: _dafny.Array = _dafny.Array(None, 0)
        self._f23: bool = False
        self._f24: str = _dafny.CodePoint('D')
        self._f25: _dafny.Seq = _dafny.Seq({})
        self._f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self)._f25 = f25
        (self).f26 = f26
        (self)._f27 = f27

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
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
    def f11(self):
        return self._f11
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    @property
    def f27(self):
        return self._f27

class C0:
    def  __init__(self):
        self.f37: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f37):
        (self).f37 = f37


class C1(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f43: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    def ctor__(self, f43, f28, f29):
        (self)._f43 = f43
        (self).f28 = f28
        (self).f29 = f29

    def fm2(self, globalState):
        source11_ = D3_DC8(_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('e'), _dafny.CodePoint('y'), _dafny.CodePoint('y'), _dafny.CodePoint('a')]), False, 611)
        if source11_.is_DC8:
            d_609___mcc_h0_ = source11_.cf15
            d_610___mcc_h1_ = source11_.cf16
            d_611___mcc_h2_ = source11_.cf17
            d_612_cf17_ = d_611___mcc_h2_
            d_613_cf16_ = d_610___mcc_h1_
            d_614_cf15_ = d_609___mcc_h0_
            return (_dafny.Map({(self).f43: d_612_cf17_})) | (_dafny.Map({(self).f43: (0) - (d_612_cf17_)}))
        elif source11_.is_DC9:
            d_615___mcc_h3_ = source11_.cf18
            d_616___mcc_h4_ = source11_.cf19
            d_617___mcc_h5_ = source11_.cf20
            d_618___mcc_h6_ = source11_.cf21
            d_619_cf21_ = d_618___mcc_h6_
            d_620_cf20_ = d_617___mcc_h5_
            d_621_cf19_ = d_616___mcc_h4_
            d_622_cf18_ = d_615___mcc_h3_
            return (_dafny.Map({(self).f43: 513})) | (_dafny.Map({False: 615}))
        elif True:
            d_623___mcc_h7_ = source11_.cf14
            d_624_cf14_ = d_623___mcc_h7_
            return (_dafny.Map({(D2_DC6(643, self.f28, 167, (self).f43, len(_dafny.SeqWithoutIsStrInference([-627])))).cf12: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('m')]))})) | (_dafny.Map({(self).f43: 342}))

    def fm22(self, p0, p1, p2, globalState):
        return True

    def m9(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        r3: bool = False
        d_625_v0_: int
        d_625_v0_ = 690
        (self).f28 = not(default__.fm1(self.f28, d_625_v0_, globalState))
        d_626_v1_: _dafny.Map
        d_626_v1_ = _dafny.Map({d_625_v0_: d_625_v0_})
        d_627_v2_: _dafny.Map
        d_627_v2_ = _dafny.Map({self.f29: ((d_626_v1_)[193] if (193) in (d_626_v1_) else len(_dafny.Set({d_625_v0_})))})
        d_628_v3_: _dafny.MultiSet
        d_628_v3_ = _dafny.MultiSet([len(d_627_v2_)])
        r3 = (d_628_v3_) == ((d_628_v3_) | (d_628_v3_))
        d_626_v1_ = (d_626_v1_) | (_dafny.Map({d_625_v0_: 854}))
        d_629_v4_: _dafny.Map
        d_629_v4_ = _dafny.Map({410: self.f28})
        d_630_v5_: _dafny.Seq
        d_630_v5_ = _dafny.SeqWithoutIsStrInference([d_625_v0_])
        d_625_v0_ = ((0) - (len((d_629_v4_).set(d_625_v0_, self.f28)))) * ((545) - (len(d_630_v5_)))
        d_631_v6_: _dafny.Seq
        d_631_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avexbc"))
        d_632_v7_: _dafny.MultiSet
        d_632_v7_ = _dafny.MultiSet([(self).f43, self.f29, ((self).f43) == ((self).f43), default__.fm1((self).f43, d_625_v0_, globalState)])
        d_633_v8_: _dafny.Set
        d_633_v8_ = _dafny.Set({False})
        d_634_v9_: _dafny.Seq
        d_634_v9_ = _dafny.SeqWithoutIsStrInference([d_627_v2_, d_627_v2_])
        d_635_v10_: D3
        d_635_v10_ = D3_DC7(d_634_v9_)
        d_636_v11_: _dafny.Seq
        d_636_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_625_v0_: d_625_v0_})])
        d_637_v12_: _dafny.Set
        d_637_v12_ = _dafny.Set({297})
        pat_let_tv5_ = d_632_v7_
        pat_let_tv6_ = d_632_v7_
        pat_let_tv7_ = d_632_v7_
        pat_let_tv8_ = d_632_v7_
        def lambda22_(source12_):
            if source12_.is_DC8:
                d_638___mcc_h0_ = source12_.cf15
                d_639___mcc_h1_ = source12_.cf16
                d_640___mcc_h2_ = source12_.cf17
                d_641_cf17_ = d_640___mcc_h2_
                d_642_cf16_ = d_639___mcc_h1_
                d_643_cf15_ = d_638___mcc_h0_
                return (pat_let_tv5_).intersection((D0_DC1(pat_let_tv6_)).cf1)
            elif source12_.is_DC9:
                d_644___mcc_h3_ = source12_.cf18
                d_645___mcc_h4_ = source12_.cf19
                d_646___mcc_h5_ = source12_.cf20
                d_647___mcc_h6_ = source12_.cf21
                d_648_cf21_ = d_647___mcc_h6_
                d_649_cf20_ = d_646___mcc_h5_
                d_650_cf19_ = d_645___mcc_h4_
                d_651_cf18_ = d_644___mcc_h3_
                return pat_let_tv7_
            elif True:
                d_652___mcc_h7_ = source12_.cf14
                d_653_cf14_ = d_652___mcc_h7_
                return pat_let_tv8_

        rhs69_ = (d_631_v6_) + (d_631_v6_)
        rhs70_ = (d_631_v6_)[default__.safeIndex(len(d_633_v8_), len(d_631_v6_))]
        rhs71_ = lambda22_(d_635_v10_)
        rhs72_ = ((d_626_v1_) | (d_626_v1_)) | (((d_636_v11_)[default__.safeIndex(len(d_631_v6_), len(d_636_v11_))]) | ((d_636_v11_)[default__.safeIndex(d_625_v0_, len(d_636_v11_))]))
        rhs73_ = (self).fm22(len(d_631_v6_), d_625_v0_, (d_637_v12_).issubset(_dafny.Set({len(d_630_v5_)})), globalState)
        lhs55_ = globalState
        d_631_v6_ = rhs69_
        r2 = rhs70_
        d_632_v7_ = rhs71_
        d_626_v1_ = rhs72_
        lhs55_.f12 = rhs73_
        if not(default__.fm1((self).f43, (d_625_v0_) * (d_625_v0_), globalState)):
            d_654_v13_: str
            d_654_v13_ = _dafny.CodePoint('t')
            d_655_v14_: _dafny.Map
            d_655_v14_ = _dafny.Map({d_654_v13_: d_630_v5_})
            d_656_v15_: _dafny.Array
            nw94_ = _dafny.Array(None, 24)
            nw94_[int(0)] = d_630_v5_
            nw94_[int(1)] = d_630_v5_
            nw94_[int(2)] = default__.fm23(self.f28, default__.fm24(self.f29, 489, globalState), globalState)
            nw94_[int(3)] = d_630_v5_
            nw94_[int(4)] = (d_630_v5_).set(default__.safeIndex(d_625_v0_, len(d_630_v5_)), default__.fm24(False, default__.fm24(False, d_625_v0_, globalState), globalState))
            nw94_[int(5)] = d_630_v5_
            nw94_[int(6)] = d_630_v5_
            nw94_[int(7)] = d_630_v5_
            nw94_[int(8)] = d_630_v5_
            nw94_[int(9)] = _dafny.SeqWithoutIsStrInference([d_625_v0_])
            nw94_[int(10)] = d_630_v5_
            nw94_[int(11)] = d_630_v5_
            nw94_[int(12)] = _dafny.SeqWithoutIsStrInference([len(d_631_v6_) for d_657_i0_ in range(default__.abs(376))])
            nw94_[int(13)] = _dafny.SeqWithoutIsStrInference([d_625_v0_ for d_658_i1_ in range(default__.abs(955))])
            nw94_[int(14)] = default__.fm23((self).f43, len(d_631_v6_), globalState)
            nw94_[int(15)] = d_630_v5_
            nw94_[int(16)] = d_630_v5_
            nw94_[int(17)] = ((d_655_v14_)[d_654_v13_] if (d_654_v13_) in (d_655_v14_) else d_630_v5_)
            nw94_[int(18)] = default__.fm23(self.f29, d_625_v0_, globalState)
            nw94_[int(19)] = d_630_v5_
            nw94_[int(20)] = d_630_v5_
            nw94_[int(21)] = (_dafny.SeqWithoutIsStrInference([d_625_v0_, 698, 327])).set(default__.safeIndex((d_630_v5_)[default__.safeIndex(99, len(d_630_v5_))], len(_dafny.SeqWithoutIsStrInference([d_625_v0_, 698, 327]))), (d_628_v3_).cardinality)
            nw94_[int(22)] = (d_630_v5_) + (d_630_v5_)
            nw94_[int(23)] = (d_630_v5_ if self.f28 else d_630_v5_)
            d_656_v15_ = nw94_
            index61_ = default__.safeIndex(956, (d_656_v15_).length(0))
            (d_656_v15_)[index61_] = d_630_v5_
            d_659_v16_: D6
            d_659_v16_ = D6_DC15(_dafny.SeqWithoutIsStrInference([d_625_v0_ for d_660_i2_ in range(default__.abs(-544))]))
            index62_ = default__.safeIndex(956, (d_656_v15_).length(0))
            (d_656_v15_)[index62_] = (d_659_v16_).cf32
            d_661_v17_: D4
            d_661_v17_ = D4_DC10((d_628_v3_) | (d_628_v3_))
            source13_ = d_661_v17_
            if source13_.is_DC11:
                d_662___mcc_h8_ = source13_.cf23
                d_663___mcc_h9_ = source13_.cf24
                d_664___mcc_h10_ = source13_.cf25
                d_665_cf25_ = d_664___mcc_h10_
                d_666_cf24_ = d_663___mcc_h9_
                d_667_cf23_ = d_662___mcc_h8_
                d_668_v18_: _dafny.MultiSet
                d_668_v18_ = _dafny.MultiSet([d_654_v13_])
                d_669_v19_: _dafny.Seq
                d_669_v19_ = _dafny.SeqWithoutIsStrInference([(self).f43, self.f28])
                (globalState).f12 = (d_628_v3_).isdisjoint((d_628_v3_).set((D3_DC8(d_668_v18_, (d_669_v19_)[default__.safeIndex(d_665_cf25_, len(d_669_v19_))], d_667_cf23_)).cf17, default__.abs(len(d_631_v6_))))
                d_670_v20_: _dafny.Array
                nw95_ = _dafny.Array(int(0), 22)
                d_670_v20_ = nw95_
                index63_ = default__.safeIndex(251, (d_670_v20_).length(0))
                (d_670_v20_)[index63_] = default__.fm24(d_666_cf24_, 818, globalState)
                d_671_v21_: _dafny.Seq
                d_671_v21_ = _dafny.SeqWithoutIsStrInference([d_631_v6_, ((d_631_v6_) + (d_631_v6_)).set(default__.safeIndex(len(_dafny.Map({d_625_v0_: d_669_v19_})), len((d_631_v6_) + (d_631_v6_))), d_654_v13_)])
                index64_ = default__.safeIndex(251, (d_670_v20_).length(0))
                rhs74_ = len((d_671_v21_)[default__.safeIndex((d_665_cf25_) - (default__.fm24((self).f43, default__.fm24(self.f28, d_667_cf23_, globalState), globalState)), len(d_671_v21_))])
                rhs75_ = d_669_v19_
                lhs56_ = d_670_v20_
                lhs57_ = default__.safeIndex(251, (d_670_v20_).length(0))
                lhs56_[lhs57_] = rhs74_
                d_669_v19_ = rhs75_
                (globalState).f12 = self.f29
                d_665_cf25_ = ((d_670_v20_)[default__.safeIndex(251, (d_670_v20_).length(0))]) - ((d_625_v0_) + (d_665_cf25_))
            elif True:
                d_672___mcc_h11_ = source13_.cf22
                d_673_cf22_ = d_672___mcc_h11_
                d_674_v22_: _dafny.Array
                nw96_ = _dafny.Array(False, 19)
                d_674_v22_ = nw96_
                index65_ = default__.safeIndex(439, (d_674_v22_).length(0))
                (d_674_v22_)[index65_] = self.f29
                index66_ = default__.safeIndex(439, (d_674_v22_).length(0))
                (d_674_v22_)[index66_] = (self).f43
                (globalState).f12 = self.f29
                r1 = (d_625_v0_) + ((d_625_v0_) - (len(_dafny.Set({((d_626_v1_)[d_625_v0_] if (d_625_v0_) in (d_626_v1_) else d_625_v0_)}))))
                (globalState).f13 = len(d_631_v6_)
            d_631_v6_ = d_631_v6_
            (globalState).f1 = ((self).f43) == (True)
            (globalState).f13 = (d_625_v0_) - (d_625_v0_)
        elif True:
            if self.f29:
                (globalState).f26 = (self).f43
                d_675_v23_: _dafny.Array
                nw97_ = _dafny.Array(False, 8)
                d_675_v23_ = nw97_
                index67_ = default__.safeIndex(434, (d_675_v23_).length(0))
                (d_675_v23_)[index67_] = default__.fm1(self.f28, d_625_v0_, globalState)
                index68_ = default__.safeIndex(434, (d_675_v23_).length(0))
                (d_675_v23_)[index68_] = (self.f29 if self.f28 else (self).f43)
                d_676_v24_: _dafny.MultiSet
                d_676_v24_ = _dafny.MultiSet([d_630_v5_])
                (globalState).f1 = (d_676_v24_) == (d_676_v24_)
                r1 = d_625_v0_
                r1 = d_625_v0_
            elif True:
                d_677_v25_: _dafny.Array
                nw98_ = _dafny.Array(int(0), 28)
                d_677_v25_ = nw98_
                index69_ = default__.safeIndex(104, (d_677_v25_).length(0))
                (d_677_v25_)[index69_] = (d_625_v0_) - (d_625_v0_)
                index70_ = default__.safeIndex(104, (d_677_v25_).length(0))
                (d_677_v25_)[index70_] = default__.safeModuloInt(d_625_v0_, 15)
                r1 = d_625_v0_
                index71_ = default__.safeIndex(104, (d_677_v25_).length(0))
                (d_677_v25_)[index71_] = ((d_677_v25_)[default__.safeIndex(104, (d_677_v25_).length(0))] if self.f29 else ((d_677_v25_)[default__.safeIndex(104, (d_677_v25_).length(0))]) * ((d_677_v25_)[default__.safeIndex(104, (d_677_v25_).length(0))]))
                d_678_v26_: _dafny.Map
                d_678_v26_ = _dafny.Map({(0) - (default__.safeDivisionInt((d_677_v25_)[default__.safeIndex(104, (d_677_v25_).length(0))], (d_677_v25_)[default__.safeIndex(104, (d_677_v25_).length(0))])): d_631_v6_})
                d_678_v26_ = (d_678_v26_).set((d_677_v25_)[default__.safeIndex(104, (d_677_v25_).length(0))], d_631_v6_)
                d_679_v27_: _dafny.Seq
                d_679_v27_ = _dafny.SeqWithoutIsStrInference([(self).f43])
                d_680_v28_: _dafny.Map
                d_680_v28_ = _dafny.Map({(d_625_v0_) > (len(d_631_v6_)): (d_679_v27_)[default__.safeIndex((d_632_v7_).cardinality, len(d_679_v27_))]})
                d_680_v28_ = (d_680_v28_).set(self.f29, self.f28)
            r2 = _dafny.CodePoint('h')
            d_681_v29_: _dafny.Array
            nw99_ = _dafny.Array(False, 12)
            d_681_v29_ = nw99_
            index72_ = default__.safeIndex(209, (d_681_v29_).length(0))
            (d_681_v29_)[index72_] = self.f28
            index73_ = default__.safeIndex(209, (d_681_v29_).length(0))
            (d_681_v29_)[index73_] = not((d_625_v0_) >= (d_625_v0_))
            r3 = (not((self).fm22(d_625_v0_, d_625_v0_, self.f28, globalState)) if not (self.f28) or ((self).f43) else self.f29)
            d_682_v30_: _dafny.Array
            def lambda23_(d_683_v0_):
                def lambda24_(d_684_i3_):
                    return (d_684_i3_) * (d_683_v0_)

                return lambda24_

            init12_ = lambda23_(d_625_v0_)
            nw100_ = _dafny.Array(None, 20)
            for i0_12_ in range(nw100_.length(0)):
                nw100_[i0_12_] = init12_(i0_12_)
            d_682_v30_ = nw100_
            index74_ = default__.safeIndex(445, (d_682_v30_).length(0))
            (d_682_v30_)[index74_] = (d_625_v0_) * (d_625_v0_)
            d_685_v31_: _dafny.Map
            d_685_v31_ = _dafny.Map({self.f28: self.f28})
            d_686_v32_: _dafny.Set
            d_686_v32_ = _dafny.Set({d_685_v31_, d_685_v31_})
            index75_ = default__.safeIndex(445, (d_682_v30_).length(0))
            (d_682_v30_)[index75_] = len(d_686_v32_)
        d_687_v33_: str
        d_687_v33_ = _dafny.CodePoint('r')
        r0 = d_687_v33_
        r1 = (d_625_v0_) - (d_625_v0_)
        r2 = d_687_v33_
        def lambda25_(source14_):
            if source14_.is_DC16:
                d_688___mcc_h12_ = source14_.cf33
                d_689_cf33_ = d_688___mcc_h12_
                return (self).f43
            elif source14_.is_DC15:
                d_690___mcc_h13_ = source14_.cf32
                d_691_cf32_ = d_690___mcc_h13_
                return not(self.f28)
            elif True:
                d_692___mcc_h14_ = source14_.cf34
                d_693_cf34_ = d_692___mcc_h14_
                return self.f29

        r3 = lambda25_(D6_DC15(d_630_v5_))
        return r0, r1, r2, r3

    def m10(self, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_694_v0_: _dafny.Seq
        d_694_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjeniah"))
        d_695_v1_: int
        d_695_v1_ = 229
        (globalState).f13 = (0) - (((0) - (len(d_694_v0_))) * (d_695_v1_))
        d_696_v2_: _dafny.Array
        def lambda26_(d_697_v1_):
            def lambda27_(d_698_i0_):
                return default__.safeModuloInt(d_698_i0_, d_697_v1_)

            return lambda27_

        init13_ = lambda26_(d_695_v1_)
        nw101_ = _dafny.Array(None, 20)
        for i0_13_ in range(nw101_.length(0)):
            nw101_[i0_13_] = init13_(i0_13_)
        d_696_v2_ = nw101_
        d_699_v3_: _dafny.Set
        d_699_v3_ = _dafny.Set({d_696_v2_})
        d_700_v4_: _dafny.Array
        nw102_ = _dafny.Array(None, 26)
        d_700_v4_ = nw102_
        d_701_v5_: D2
        d_701_v5_ = D2_DC5(d_700_v4_)
        d_702_v6_: _dafny.Map
        d_702_v6_ = _dafny.Map({(_dafny.Set({d_696_v2_})).issubset(d_699_v3_): (d_701_v5_ if self.f29 else D2_DC5(d_700_v4_))})
        rhs76_ = (0) - ((d_695_v1_) - (default__.safeModuloInt(d_695_v1_, 390)))
        rhs77_ = False
        rhs78_ = d_702_v6_
        lhs58_ = globalState
        d_695_v1_ = rhs76_
        lhs58_.f12 = rhs77_
        d_702_v6_ = rhs78_
        d_703_i1_: int
        d_703_i1_ = 0
        with _dafny.label("1"):
            while not(self.f28):
                with _dafny.c_label("1"):
                    if (d_703_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_703_i1_ = (d_703_i1_) + (1)
                    d_704_v7_: _dafny.Array
                    nw103_ = _dafny.Array(False, 16)
                    d_704_v7_ = nw103_
                    (globalState).f21 = d_704_v7_
                    d_705_v8_: _dafny.Array
                    def lambda28_(d_706_i2_):
                        return _dafny.CodePoint('i')

                    init14_ = lambda28_
                    nw104_ = _dafny.Array(None, 9)
                    for i0_14_ in range(nw104_.length(0)):
                        nw104_[i0_14_] = init14_(i0_14_)
                    d_705_v8_ = nw104_
                    d_707_v9_: C0
                    nw105_ = C0()
                    nw105_.ctor__(d_705_v8_)
                    d_707_v9_ = nw105_
                    d_708_v10_: _dafny.Seq
                    d_708_v10_ = _dafny.SeqWithoutIsStrInference([111, d_695_v1_, d_695_v1_])
                    d_709_v11_: _dafny.Seq
                    d_709_v11_ = _dafny.SeqWithoutIsStrInference([self.f29])
                    d_710_v12_: int
                    out48_: int
                    out48_ = default__.m0(d_708_v10_, (d_709_v11_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([758 for d_711_i3_ in range(default__.abs(368))])), len(d_709_v11_))], (0) - (d_695_v1_), self.f29, globalState)
                    d_710_v12_ = out48_
                    d_712_v13_: _dafny.MultiSet
                    d_712_v13_ = _dafny.MultiSet([d_710_v12_, d_710_v12_])
                    d_713_v14_: _dafny.Set
                    d_713_v14_ = _dafny.Set({d_712_v13_})
                    d_714_v15_: _dafny.Map
                    d_714_v15_ = _dafny.Map({(d_713_v14_ if self.f28 else d_713_v14_): d_710_v12_})
                    d_714_v15_ = (d_714_v15_).set((_dafny.Set({default__.fm25(globalState)})) - (d_713_v14_), len(d_708_v10_))
                    pass
            pass
        d_715_v16_: _dafny.Array
        nw106_ = _dafny.Array(False, 1)
        d_715_v16_ = nw106_
        index76_ = default__.safeIndex(690, (d_715_v16_).length(0))
        (d_715_v16_)[index76_] = (self).f43
        index77_ = default__.safeIndex(690, (d_715_v16_).length(0))
        (d_715_v16_)[index77_] = self.f29
        d_716_v17_: _dafny.Array
        nw107_ = _dafny.Array(_dafny.CodePoint('D'), 18)
        d_716_v17_ = nw107_
        d_717_v18_: C0
        nw108_ = C0()
        nw108_.ctor__(d_716_v17_)
        d_717_v18_ = nw108_
        d_718_v19_: _dafny.Set
        d_718_v19_ = _dafny.Set({d_695_v1_, d_695_v1_})
        d_695_v1_ = default__.safeDivisionInt(442, len(_dafny.Map({d_718_v19_: d_695_v1_})))
        d_719_v20_: _dafny.Map
        d_719_v20_ = _dafny.Map({not((d_715_v16_)[default__.safeIndex(690, (d_715_v16_).length(0))]): _dafny.SeqWithoutIsStrInference([self.f28, not(not(not(self.f29))), self.f29, self.f28, self.f29])})
        r0 = (d_719_v20_) | ((d_719_v20_) | (default__.fm26(self.f29, globalState)))
        return r0

    @property
    def f43(self):
        return self._f43

class C2(T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f30: int = int(0)
        self._f46: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f46, f30, f31, f28, f29):
        (self)._f46 = f46
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm3(self, p0, p1, globalState):
        return (D5_DC12(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlpme")))).cf26

    def fm2(self, globalState):
        return ((_dafny.Map({self.f28: (self).f30}) if self.f28 else _dafny.Map({self.f28: (self).f30}))) | (_dafny.Map({self.f29: -74}))

    def fm29(self, globalState):
        source15_ = (D5_DC13(self.f29, _dafny.Map({self.f31: (self).f30}), _dafny.CodePoint('k')) if self.f29 else D5_DC13(self.f31, _dafny.Map({self.f29: (self).f46}), _dafny.CodePoint('m')))
        if source15_.is_DC13:
            d_720___mcc_h0_ = source15_.cf27
            d_721___mcc_h1_ = source15_.cf28
            d_722___mcc_h2_ = source15_.cf29
            d_723_cf29_ = d_722___mcc_h2_
            d_724_cf28_ = d_721___mcc_h1_
            d_725_cf27_ = d_720___mcc_h0_
            return (self).f46
        elif source15_.is_DC14:
            d_726___mcc_h3_ = source15_.cf30
            d_727___mcc_h4_ = source15_.cf31
            d_728_cf31_ = d_727___mcc_h4_
            d_729_cf30_ = d_726___mcc_h3_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_728_cf31_]))).cardinality
        elif True:
            d_730___mcc_h5_ = source15_.cf26
            d_731_cf26_ = d_730___mcc_h5_
            return (self).f46

    def fm30(self, p0, globalState):
        return self.f28

    def m13(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_732_v0_: _dafny.Seq
        d_732_v0_ = _dafny.SeqWithoutIsStrInference([(0) - (p3)])
        d_733_i0_: int
        d_733_i0_ = 0
        with _dafny.label("2"):
            while (d_732_v0_) < (d_732_v0_):
                with _dafny.c_label("2"):
                    if (d_733_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_733_i0_ = (d_733_i0_) + (1)
                    d_734_v1_: _dafny.Array
                    nw109_ = _dafny.Array(False, 11)
                    d_734_v1_ = nw109_
                    d_735_v2_: _dafny.Map
                    d_735_v2_ = _dafny.Map({(self).f30: (0) - ((_dafny.MultiSet([d_734_v1_])).cardinality)})
                    (globalState).f13 = (p3 if p2 else ((d_735_v2_)[(self).f46] if ((self).f46) in (d_735_v2_) else (self).f30))
                    (globalState).f13 = (self).f30
                    if self.f31:
                        (globalState).f12 = p2
                        d_736_v3_: _dafny.Seq
                        d_736_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgyuevq"))
                        d_736_v3_ = (self).fm3(self.f28, (p3) <= (p3), globalState)
                        (globalState).f13 = ((p1)[self.f28] if (self.f28) in (p1) else ((self).f46 if not(True) else len(d_736_v3_)))
                        d_732_v0_ = _dafny.SeqWithoutIsStrInference([(-583) - ((self).f46)])
                        (globalState).f13 = (self).f46
                    elif True:
                        (globalState).f13 = (self).f46
                        (globalState).f13 = (self).f46
                        index78_ = default__.safeIndex(333, (d_734_v1_).length(0))
                        (d_734_v1_)[index78_] = self.f31
                        d_737_v4_: _dafny.Seq
                        d_737_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                        index79_ = default__.safeIndex(333, (d_734_v1_).length(0))
                        (d_734_v1_)[index79_] = (len(d_737_v4_)) < ((self).f46)
                        index80_ = default__.safeIndex(333, (d_734_v1_).length(0))
                        (d_734_v1_)[index80_] = (d_734_v1_)[default__.safeIndex(333, (d_734_v1_).length(0))]
                        d_738_v5_: _dafny.Array
                        nw110_ = _dafny.Array(False, 4)
                        d_738_v5_ = nw110_
                        d_739_v6_: _dafny.Map
                        d_739_v6_ = _dafny.Map({d_738_v5_: (d_737_v4_) + (d_737_v4_)})
                        d_739_v6_ = (d_739_v6_).set(d_738_v5_, d_737_v4_)
                    d_740_v7_: _dafny.Array
                    nw111_ = _dafny.Array(_dafny.Set({}), 27)
                    d_740_v7_ = nw111_
                    d_741_v8_: _dafny.Set
                    d_741_v8_ = _dafny.Set({self.f31, self.f28})
                    index81_ = default__.safeIndex(542, (d_740_v7_).length(0))
                    (d_740_v7_)[index81_] = d_741_v8_
                    d_742_v9_: _dafny.Map
                    d_742_v9_ = _dafny.Map({self.f28: self.f29})
                    index82_ = default__.safeIndex(542, (d_740_v7_).length(0))
                    rhs79_ = (_dafny.Set({self.f28, self.f29})) | (d_741_v8_)
                    rhs80_ = ((self).f46 if (True) == (self.f31) else (self).f46)
                    rhs81_ = ((self).f46) + ((len(d_742_v9_)) * ((self).f30))
                    lhs59_ = d_740_v7_
                    lhs60_ = default__.safeIndex(542, (d_740_v7_).length(0))
                    lhs61_ = globalState
                    lhs62_ = globalState
                    lhs59_[lhs60_] = rhs79_
                    lhs61_.f13 = rhs80_
                    lhs62_.f13 = rhs81_
                    pass
            pass
        d_743_v10_: _dafny.Array
        def lambda29_(d_744_p3_):
            def lambda30_(d_745_i2_):
                return (_dafny.Set({(self).f46, -550})) | (_dafny.Set({d_744_p3_}))

            return lambda30_

        init15_ = lambda29_(p3)
        nw112_ = _dafny.Array(None, 27)
        for i0_15_ in range(nw112_.length(0)):
            nw112_[i0_15_] = init15_(i0_15_)
        d_743_v10_ = nw112_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_743_v10_).length(0)):
            d_746_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_746_i1_)) and ((d_746_i1_) < ((d_743_v10_).length(0)))):
                (d_743_v10_)[(d_746_i1_)] = ((_dafny.Set({(self).f30, (self).f30})) - (_dafny.Set({(d_732_v0_)[default__.safeIndex((self).f46, len(d_732_v0_))]}))).intersection((_dafny.Set({p3})) | (_dafny.Set({950})))
        d_747_v11_: C1
        nw113_ = C1()
        nw113_.ctor__(self.f29, self.f29, not (True) or (not(self.f28)))
        d_747_v11_ = nw113_
        d_748_v12_: _dafny.Array
        nw114_ = _dafny.Array(_dafny.CodePoint('D'), 27)
        d_748_v12_ = nw114_
        d_749_v13_: C0
        nw115_ = C0()
        nw115_.ctor__(d_748_v12_)
        d_749_v13_ = nw115_
        d_750_v14_: _dafny.Seq
        d_750_v14_ = _dafny.SeqWithoutIsStrInference([True])
        d_751_v15_: _dafny.Seq
        d_751_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuygtgj"))
        d_750_v14_ = (default__.fm0(default__.safeModuloInt(len(d_751_v15_), (self).f30), self.f29, self.f29, 808, globalState)).set(default__.safeIndex((self).f30, len(default__.fm0(default__.safeModuloInt(len(d_751_v15_), (self).f30), self.f29, self.f29, 808, globalState))), True)
        d_752_v16_: _dafny.Map
        d_752_v16_ = _dafny.Map({default__.fm1(p2, len(p1), globalState): d_749_v13_.f37})
        d_753_v17_: C0
        nw116_ = C0()
        nw116_.ctor__(((d_752_v16_)[p2] if (p2) in (d_752_v16_) else d_749_v13_.f37))
        d_753_v17_ = nw116_
        r0 = (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_754_i3_ in range(default__.abs(61))])) + (d_751_v15_))) <= ((d_732_v0_)[default__.safeIndex((self).f30, len(d_732_v0_))])
        return r0

    @property
    def f46(self):
        return self._f46

class C3(T4, T3, T2, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f36: str = _dafny.CodePoint('D')
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        self._f44: int = int(0)
        self._f45: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f36(self):
        return self._f36
    @f36.setter
    def f36(self, value):
        self._f36 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f44, f45, f36, f34, f35, f32, f33, f30, f31, f28, f29):
        (self)._f44 = f44
        (self)._f45 = f45
        (self).f36 = f36
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm7(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsrndyx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsqkph")))) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfxuo"))) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_755_i0_ in range(default__.abs(673))])))

    def fm6(self, globalState):
        return self.f29

    def fm4(self, p0, p1, p2, p3, globalState):
        return _dafny.Set({((self).f34 if self.f35 else (0) - ((self).f34)), (self).f34, ((self).f34) - ((self).f30), (self).f30, default__.safeDivisionInt((self).f30, (self).f30)})

    def fm5(self, p0, p1, p2, p3, globalState):
        return (((self).f32) + ((self).f32)) + ((self).f32)

    def fm3(self, p0, p1, globalState):
        source16_ = (D5_DC14(self.f35, len(_dafny.Map({(self).f44: (self).f44}))) if self.f28 else D5_DC14(self.f29, (self).f44))
        if source16_.is_DC13:
            d_756___mcc_h0_ = source16_.cf27
            d_757___mcc_h1_ = source16_.cf28
            d_758___mcc_h2_ = source16_.cf29
            d_759_cf29_ = d_758___mcc_h2_
            d_760_cf28_ = d_757___mcc_h1_
            d_761_cf27_ = d_756___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cca")))
        elif source16_.is_DC14:
            d_762___mcc_h3_ = source16_.cf30
            d_763___mcc_h4_ = source16_.cf31
            d_764_cf31_ = d_763___mcc_h4_
            d_765_cf30_ = d_762___mcc_h3_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auph"))
        elif True:
            d_766___mcc_h5_ = source16_.cf26
            d_767_cf26_ = d_766___mcc_h5_
            return d_767_cf26_

    def fm2(self, globalState):
        return ((_dafny.Map({self.f29: 230})) | (_dafny.Map({self.f28: (self).f44}))) | (_dafny.Map({self.f31: (self).f34}))

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (globalState).f26 = ((self).f44) >= (66)
        d_768_v0_: _dafny.Array
        nw117_ = _dafny.Array(None, 4)
        nw117_[int(0)] = self.f36
        nw117_[int(1)] = self.f36
        nw117_[int(2)] = self.f36
        nw117_[int(3)] = _dafny.CodePoint('o')
        d_768_v0_ = nw117_
        d_769_v1_: _dafny.Set
        d_769_v1_ = _dafny.Set({d_768_v0_, d_768_v0_})
        if ((d_769_v1_) - (d_769_v1_)).ispropersubset(d_769_v1_):
            d_770_v2_: _dafny.Seq
            d_770_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmi"))
            d_771_v3_: _dafny.Map
            d_771_v3_ = _dafny.Map({(self).f30: len((d_770_v2_).set(default__.safeIndex((self).f30, len(d_770_v2_)), (d_770_v2_)[default__.safeIndex((self).f30, len(d_770_v2_))]))})
            d_772_v4_: _dafny.MultiSet
            d_772_v4_ = _dafny.MultiSet([(self).f44])
            d_773_v5_: T1
            nw118_ = C2()
            def iife41_(_pat_let9_0):
                def iife42_(d_774_dt__update__tmp_h0_):
                    def iife43_(_pat_let10_0):
                        def iife44_(d_775_dt__update_hcf31_h0_):
                            return D5_DC14((d_774_dt__update__tmp_h0_).cf30, d_775_dt__update_hcf31_h0_)
                        return iife44_(_pat_let10_0)
                    return iife43_(-18)
                return iife42_(_pat_let9_0)
            nw118_.ctor__(len(default__.fm31(default__.fm32(not(self.f28), globalState), d_771_v3_, (iife41_(D5_DC14(self.f35, (self).f44))).cf30, ((self).f32)[default__.safeIndex((d_772_v4_).cardinality, len((self).f32))], globalState)), 454, ((self).f34) == ((self).f30), self.f35, ((self).f32)[default__.safeIndex(900, len((self).f32))])
            d_773_v5_ = nw118_
            d_773_v5_ = (d_773_v5_ if self.f31 else d_773_v5_)
            d_776_v6_: _dafny.Map
            d_776_v6_ = _dafny.Map({d_773_v5_.f31: (self).f44})
            d_777_v7_: D5
            d_777_v7_ = D5_DC13(d_773_v5_.f31, d_776_v6_, self.f36)
            source17_ = d_777_v7_
            if source17_.is_DC13:
                d_778___mcc_h0_ = source17_.cf27
                d_779___mcc_h1_ = source17_.cf28
                d_780___mcc_h2_ = source17_.cf29
                d_781_cf29_ = d_780___mcc_h2_
                d_782_cf28_ = d_779___mcc_h1_
                d_783_cf27_ = d_778___mcc_h0_
                (globalState).f13 = default__.safeModuloInt((_dafny.MultiSet([(self).f34])).cardinality, (self).f34)
                d_784_v8_: _dafny.Seq
                d_784_v8_ = _dafny.SeqWithoutIsStrInference([(self).f34, (d_773_v5_).f30])
                d_785_v9_: _dafny.Array
                nw119_ = _dafny.Array(None, 10)
                nw119_[int(0)] = default__.fm32(self.f28, globalState)
                nw119_[int(1)] = (d_784_v8_)[default__.safeIndex((d_773_v5_).f30, len(d_784_v8_))]
                nw119_[int(2)] = ((d_773_v5_).f30) + (len(_dafny.SeqWithoutIsStrInference([d_781_cf29_ for d_786_i0_ in range(default__.abs(238))])))
                nw119_[int(3)] = len(d_784_v8_)
                nw119_[int(4)] = (0) - ((self).f30)
                nw119_[int(5)] = (d_773_v5_).f30
                nw119_[int(6)] = (self).f34
                nw119_[int(7)] = -91
                nw119_[int(8)] = default__.safeDivisionInt((d_773_v5_).f30, (d_773_v5_).f30)
                nw119_[int(9)] = (self).f30
                d_785_v9_ = nw119_
                nw120_ = _dafny.Array(int(0), 1)
                d_785_v9_ = nw120_
                d_787_v10_: C0
                nw121_ = C0()
                nw121_.ctor__(d_768_v0_)
                d_787_v10_ = nw121_
                d_788_v11_: D1
                d_788_v11_ = D1_DC3(d_773_v5_.f29, (self).f34, d_787_v10_, (self).f44)
                d_789_v12_: D1
                d_789_v12_ = D1_DC4(d_788_v11_)
                d_790_v13_: _dafny.MultiSet
                d_790_v13_ = _dafny.MultiSet([d_789_v12_, d_789_v12_])
                d_791_v14_: _dafny.Seq
                d_791_v14_ = _dafny.SeqWithoutIsStrInference([d_789_v12_])
                pat_let_tv9_ = d_788_v11_
                d_792_v15_: _dafny.Array
                nw122_ = _dafny.Array(None, 14)
                def iife45_(_pat_let11_0):
                    def iife46_(d_793_dt__update__tmp_h1_):
                        def iife47_(_pat_let12_0):
                            def iife48_(d_794_dt__update_hcf7_h0_):
                                return D1_DC4(d_794_dt__update_hcf7_h0_)
                            return iife48_(_pat_let12_0)
                        return iife47_(pat_let_tv9_)
                    return iife46_(_pat_let11_0)
                nw122_[int(0)] = _dafny.MultiSet([d_789_v12_, iife45_(d_789_v12_), d_789_v12_, D1_DC4(d_788_v11_), d_789_v12_])
                nw122_[int(1)] = (d_790_v13_) | (d_790_v13_)
                nw122_[int(2)] = d_790_v13_
                nw122_[int(3)] = d_790_v13_
                nw122_[int(4)] = d_790_v13_
                nw122_[int(5)] = (d_790_v13_).set(D1_DC4(d_788_v11_), default__.abs(len(d_770_v2_)))
                nw122_[int(6)] = d_790_v13_
                nw122_[int(7)] = d_790_v13_
                nw122_[int(8)] = (d_790_v13_).intersection(d_790_v13_)
                nw122_[int(9)] = _dafny.MultiSet([D1_DC4(d_788_v11_)])
                nw122_[int(10)] = d_790_v13_
                nw122_[int(11)] = d_790_v13_
                nw122_[int(12)] = (_dafny.MultiSet(d_791_v14_) if d_773_v5_.f29 else d_790_v13_)
                nw122_[int(13)] = d_790_v13_
                d_792_v15_ = nw122_
                index83_ = default__.safeIndex(980, (d_792_v15_).length(0))
                (d_792_v15_)[index83_] = _dafny.MultiSet([d_789_v12_])
                d_795_v16_: _dafny.Seq
                d_795_v16_ = _dafny.SeqWithoutIsStrInference([d_791_v14_, d_791_v14_])
                index84_ = default__.safeIndex(980, (d_792_v15_).length(0))
                (d_792_v15_)[index84_] = _dafny.MultiSet((d_791_v14_ if d_773_v5_.f31 else (d_795_v16_)[default__.safeIndex((d_773_v5_).f30, len(d_795_v16_))]))
                (globalState).f12 = d_773_v5_.f29
            elif source17_.is_DC14:
                d_796___mcc_h3_ = source17_.cf30
                d_797___mcc_h4_ = source17_.cf31
                d_798_cf31_ = d_797___mcc_h4_
                d_799_cf30_ = d_796___mcc_h3_
                (d_773_v5_).f31 = not(d_773_v5_.f31)
                (globalState).f1 = not(self.f28)
                d_800_v17_: _dafny.Array
                nw123_ = _dafny.Array(False, 28)
                d_800_v17_ = nw123_
                index85_ = default__.safeIndex(42, (d_800_v17_).length(0))
                (d_800_v17_)[index85_] = d_773_v5_.f31
                d_801_v18_: _dafny.Map
                d_801_v18_ = _dafny.Map({242: False})
                d_802_v19_: D2
                d_802_v19_ = D2_DC6((self).f34, d_773_v5_.f28, len(d_801_v18_), self.f31, (d_773_v5_).f30)
                d_803_v20_: _dafny.Seq
                d_803_v20_ = _dafny.SeqWithoutIsStrInference([d_802_v19_, d_802_v19_])
                d_804_v21_: D2
                d_804_v21_ = D2_DC6(default__.safeModuloInt(default__.fm32(self.f31, globalState), (self).f30), (d_802_v19_) not in (d_803_v20_), (d_773_v5_).f30, d_773_v5_.f29, (self).f44)
                d_805_v22_: D5
                d_805_v22_ = D5_DC14(d_773_v5_.f31, (self).f30)
                index86_ = default__.safeIndex(42, (d_800_v17_).length(0))
                rhs82_ = self.f31
                rhs83_ = (d_804_v21_ if True else d_802_v19_)
                rhs84_ = (227) * ((self).f34)
                rhs85_ = default__.fm1((d_805_v22_).cf30, ((self).f30) + ((self).f30), globalState)
                rhs86_ = (d_773_v5_.f31 if (self.f28 if self.f35 else d_773_v5_.f31) else (_dafny.SeqWithoutIsStrInference([d_799_cf30_])) == ((self).f32))
                lhs63_ = d_800_v17_
                lhs64_ = default__.safeIndex(42, (d_800_v17_).length(0))
                lhs65_ = globalState
                lhs66_ = d_773_v5_
                lhs63_[lhs64_] = rhs82_
                d_802_v19_ = rhs83_
                d_798_cf31_ = rhs84_
                lhs65_.f12 = rhs85_
                lhs66_.f28 = rhs86_
                (globalState).f13 = default__.safeDivisionInt((-11 if ((self).f32)[default__.safeIndex(d_798_cf31_, len((self).f32))] else (d_773_v5_).f30), (self).f30)
            elif True:
                d_806___mcc_h5_ = source17_.cf26
                d_807_cf26_ = d_806___mcc_h5_
                d_808_v23_: _dafny.Map
                d_808_v23_ = _dafny.Map({d_771_v3_: d_773_v5_.f28})
                d_809_v24_: _dafny.Set
                d_810_v25_: bool
                out49_: _dafny.Set
                out50_: bool
                out49_, out50_ = (self).m12((self).f30, d_808_v23_, globalState)
                d_809_v24_ = out49_
                d_810_v25_ = out50_
                (globalState).f13 = ((self).f34) - ((d_773_v5_).f30)
                d_811_v26_: _dafny.Array
                nw124_ = _dafny.Array(int(0), 15)
                d_811_v26_ = nw124_
                index87_ = default__.safeIndex(914, (d_811_v26_).length(0))
                (d_811_v26_)[index87_] = len(d_776_v6_)
                index88_ = default__.safeIndex(914, (d_811_v26_).length(0))
                rhs87_ = default__.fm32(d_810_v25_, globalState)
                rhs88_ = d_773_v5_.f29
                lhs67_ = d_811_v26_
                lhs68_ = default__.safeIndex(914, (d_811_v26_).length(0))
                lhs69_ = d_773_v5_
                lhs67_[lhs68_] = rhs87_
                lhs69_.f31 = rhs88_
                d_812_v27_: C1
                nw125_ = C1()
                nw125_.ctor__(self.f31, d_773_v5_.f29, d_773_v5_.f28)
                d_812_v27_ = nw125_
                d_813_v28_: _dafny.MultiSet
                d_813_v28_ = _dafny.MultiSet([d_812_v27_])
                (globalState).f1 = ((d_813_v28_).set(d_812_v27_, default__.abs(895))).isdisjoint((d_813_v28_).set(d_812_v27_, default__.abs((self).f34)))
            (globalState).f1 = d_773_v5_.f31
            d_814_v29_: C0
            nw126_ = C0()
            nw126_.ctor__(d_768_v0_)
            d_814_v29_ = nw126_
            (d_773_v5_).f28 = True
        elif True:
            d_815_v30_: _dafny.Map
            d_815_v30_ = _dafny.Map({self.f35: self.f28})
            d_816_v31_: D0
            d_816_v31_ = D0_DC1(_dafny.MultiSet([not(((d_815_v30_)[False] if (False) in (d_815_v30_) else not(self.f28)))]))
            d_816_v31_ = d_816_v31_
            (globalState).f13 = (self).f30
            (self).f31 = ((self).f30) >= ((self).f34)
            d_817_v32_: _dafny.Set
            d_817_v32_ = _dafny.Set({self.f28, self.f35})
            d_818_v33_: _dafny.Seq
            d_818_v33_ = _dafny.SeqWithoutIsStrInference([len(d_817_v32_)])
            d_819_v34_: _dafny.Map
            d_819_v34_ = _dafny.Map({-187: len(d_818_v33_)})
            d_820_v35_: _dafny.Seq
            d_820_v35_ = _dafny.SeqWithoutIsStrInference([(d_819_v34_).set((self).f44, (self).f34)])
            rhs89_ = d_820_v35_
            rhs90_ = self.f28
            rhs91_ = self.f29
            lhs70_ = globalState
            lhs71_ = globalState
            d_820_v35_ = rhs89_
            lhs70_.f26 = rhs90_
            lhs71_.f1 = rhs91_
            d_821_v36_: C0
            nw127_ = C0()
            nw127_.ctor__(d_768_v0_)
            d_821_v36_ = nw127_
        d_822_v37_: _dafny.MultiSet
        d_822_v37_ = _dafny.MultiSet([self.f31, self.f31])
        hi4_ = ((self).f44) - (((d_822_v37_)[self.f31] if (self.f31) in (d_822_v37_) else (self).f44))
        for d_823_i1_ in range((self).f30, hi4_):
            (globalState).f12 = self.f28
            d_824_v38_: D2
            d_824_v38_ = D2_DC6(51, self.f28, (self).f30, self.f35, 290)
            d_825_v39_: _dafny.MultiSet
            d_825_v39_ = _dafny.MultiSet([(self).f34])
            d_826_v40_: _dafny.Seq
            d_826_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksmswsjps"))
            d_827_v41_: _dafny.Set
            d_827_v41_ = _dafny.Set({d_826_v40_, d_826_v40_, d_826_v40_, d_826_v40_, d_826_v40_})
            d_828_v42_: _dafny.Seq
            d_828_v42_ = _dafny.SeqWithoutIsStrInference([(self).f34, (self).f44])
            d_829_v43_: _dafny.Seq
            d_829_v43_ = _dafny.SeqWithoutIsStrInference([len(d_828_v42_)])
            d_830_v44_: _dafny.Set
            d_830_v44_ = _dafny.Set({len(d_828_v42_)})
            d_831_v45_: _dafny.Set
            d_831_v45_ = _dafny.Set({self.f28})
            d_832_v46_: _dafny.Map
            d_832_v46_ = _dafny.Map({(self).f44: self.f28})
            d_833_v47_: _dafny.Array
            nw128_ = _dafny.Array(None, 25)
            nw128_[int(0)] = not((d_824_v38_).cf10)
            nw128_[int(1)] = (self).fm7(d_825_v39_, 125, globalState)
            nw128_[int(2)] = (d_827_v41_).issubset(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aa")), d_826_v40_}))
            nw128_[int(3)] = self.f29
            nw128_[int(4)] = self.f31
            nw128_[int(5)] = self.f29
            nw128_[int(6)] = ((self).f34) < (d_823_i1_)
            nw128_[int(7)] = self.f31
            nw128_[int(8)] = True
            nw128_[int(9)] = False
            nw128_[int(10)] = self.f35
            nw128_[int(11)] = self.f35
            nw128_[int(12)] = self.f35
            nw128_[int(13)] = self.f35
            nw128_[int(14)] = not(self.f29)
            nw128_[int(15)] = self.f28
            nw128_[int(16)] = (False if self.f31 else self.f31)
            nw128_[int(17)] = (d_830_v44_) == (d_830_v44_)
            nw128_[int(18)] = not(self.f31)
            nw128_[int(19)] = self.f31
            nw128_[int(20)] = (d_831_v45_).ispropersubset(d_831_v45_)
            nw128_[int(21)] = (135) in (d_832_v46_)
            nw128_[int(22)] = self.f29
            nw128_[int(23)] = (len(d_828_v42_)) not in (_dafny.MultiSet([len(d_826_v40_), (self).f34]))
            nw128_[int(24)] = False
            d_833_v47_ = nw128_
            index89_ = default__.safeIndex(277, (d_833_v47_).length(0))
            (d_833_v47_)[index89_] = (default__.fm32(self.f29, globalState)) == (d_823_i1_)
            index90_ = default__.safeIndex(277, (d_833_v47_).length(0))
            (d_833_v47_)[index90_] = (self).fm7(d_825_v39_, default__.safeDivisionInt(len(d_831_v45_), (self).f44), globalState)
            d_834_v48_: _dafny.MultiSet
            d_834_v48_ = _dafny.MultiSet([self.f36, _dafny.CodePoint('o')])
            d_835_v49_: int
            out51_: int
            out51_ = default__.m0(d_829_v43_, not((self.f36) in (d_826_v40_)), (d_834_v48_).cardinality, not(self.f28), globalState)
            d_835_v49_ = out51_
            d_836_v50_: _dafny.Map
            d_836_v50_ = _dafny.Map({_dafny.MultiSet([False]): d_768_v0_})
            d_837_v51_: D0
            d_837_v51_ = D0_DC1(default__.fm33(self.f29, globalState))
            d_836_v50_ = (d_836_v50_).set((d_837_v51_).cf1, d_768_v0_)
        d_838_v52_: _dafny.Seq
        d_838_v52_ = _dafny.SeqWithoutIsStrInference([d_822_v37_, d_822_v37_, d_822_v37_, d_822_v37_])
        d_839_v53_: C0
        nw129_ = C0()
        nw129_.ctor__(d_768_v0_)
        d_839_v53_ = nw129_
        d_840_v54_: _dafny.Seq
        d_840_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrneeb"))
        d_841_v55_: D1
        d_841_v55_ = D1_DC3(self.f35, (self).f34, d_839_v53_, len(d_840_v54_))
        d_842_v56_: _dafny.MultiSet
        d_842_v56_ = _dafny.MultiSet([(d_822_v37_).cardinality, (self).f34])
        d_843_v57_: _dafny.Array
        nw130_ = _dafny.Array(None, 26)
        nw130_[int(0)] = (self).f34
        nw130_[int(1)] = (self).f30
        nw130_[int(2)] = ((d_838_v52_)[default__.safeIndex((d_841_v55_).cf4, len(d_838_v52_))]).cardinality
        nw130_[int(3)] = (self).f34
        nw130_[int(4)] = ((self).f44) + (len(d_840_v54_))
        nw130_[int(5)] = ((self).f44) * ((self).f34)
        nw130_[int(6)] = (self).f30
        nw130_[int(7)] = -578
        nw130_[int(8)] = (self).f44
        nw130_[int(9)] = (self).f30
        nw130_[int(10)] = (self).f34
        nw130_[int(11)] = (self).f30
        nw130_[int(12)] = (self).f44
        nw130_[int(13)] = (self).f34
        nw130_[int(14)] = (len(_dafny.Map({(self).f30: 337}))) * ((self).f30)
        nw130_[int(15)] = ((self).f30) + ((self).f44)
        nw130_[int(16)] = (0) - ((self).f34)
        nw130_[int(17)] = (self).f34
        nw130_[int(18)] = ((self).f34) - ((self).f30)
        nw130_[int(19)] = (self).f34
        nw130_[int(20)] = ((0) - ((self).f34) if self.f28 else (self).f44)
        nw130_[int(21)] = ((d_842_v56_)[(0) - ((self).f44)] if ((0) - ((self).f44)) in (d_842_v56_) else len(_dafny.SeqWithoutIsStrInference([self.f36])))
        nw130_[int(22)] = default__.safeDivisionInt((self).f44, len(d_840_v54_))
        nw130_[int(23)] = (self).f44
        nw130_[int(24)] = (self).f44
        nw130_[int(25)] = ((d_822_v37_)[not(self.f35)] if (not(self.f35)) in (d_822_v37_) else (self).f44)
        d_843_v57_ = nw130_
        d_843_v57_ = d_843_v57_
        if (self).fm6(globalState):
            if self.f35:
                d_844_v58_: _dafny.Array
                nw131_ = _dafny.Array(False, 23)
                d_844_v58_ = nw131_
                rhs92_ = self.f31
                rhs93_ = d_844_v58_
                rhs94_ = d_843_v57_
                rhs95_ = self.f28
                rhs96_ = False
                lhs72_ = globalState
                lhs73_ = globalState
                lhs74_ = self
                r0 = rhs92_
                lhs72_.f21 = rhs93_
                d_843_v57_ = rhs94_
                lhs73_.f12 = rhs95_
                lhs74_.f28 = rhs96_
                d_845_v59_: _dafny.Seq
                d_845_v59_ = _dafny.SeqWithoutIsStrInference([d_843_v57_, d_843_v57_])
                rhs97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                rhs98_ = d_844_v58_
                rhs99_ = ((self).fm3(not((d_843_v57_) not in (d_845_v59_)), self.f29, globalState)).set(default__.safeIndex(default__.safeDivisionInt((d_842_v56_).cardinality, (self).f30), len((self).fm3(not((d_843_v57_) not in (d_845_v59_)), self.f29, globalState))), (d_840_v54_)[default__.safeIndex((0) - (len(d_838_v52_)), len(d_840_v54_))])
                lhs75_ = globalState
                d_840_v54_ = rhs97_
                lhs75_.f21 = rhs98_
                d_840_v54_ = rhs99_
                index91_ = default__.safeIndex(616, (d_843_v57_).length(0))
                (d_843_v57_)[index91_] = (self).f30
                index92_ = default__.safeIndex(616, (d_843_v57_).length(0))
                (d_843_v57_)[index92_] = ((self).f34) - ((self).f30)
                d_846_v60_: _dafny.Seq
                d_846_v60_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_843_v57_)[default__.safeIndex(616, (d_843_v57_).length(0))])])
                (globalState).f12 = ((0) - ((len(d_846_v60_)) * ((d_843_v57_)[default__.safeIndex(616, (d_843_v57_).length(0))]))) > ((self).f34)
                index93_ = default__.safeIndex(836, (d_844_v58_).length(0))
                (d_844_v58_)[index93_] = not(False)
                index94_ = default__.safeIndex(836, (d_844_v58_).length(0))
                (d_844_v58_)[index94_] = ((self).f44) < ((self).f30)
            elif True:
                r2 = d_840_v54_
                d_847_v61_: _dafny.Map
                d_847_v61_ = _dafny.Map({self.f31: self.f36})
                d_848_v62_: _dafny.Map
                d_848_v62_ = _dafny.Map({((d_847_v61_)[self.f35] if (self.f35) in (d_847_v61_) else self.f36): len(d_840_v54_)})
                index95_ = default__.safeIndex(169, (d_843_v57_).length(0))
                (d_843_v57_)[index95_] = default__.safeDivisionInt(len(d_848_v62_), len(_dafny.Map({_dafny.CodePoint('g'): (self).f34})))
                d_849_v63_: D2
                d_849_v63_ = D2_DC6((self).f34, self.f31, (self).f34, self.f29, (self).f30)
                d_850_v64_: _dafny.MultiSet
                d_850_v64_ = _dafny.MultiSet([self.f36, self.f36, self.f36, self.f36])
                d_851_v65_: D3
                d_851_v65_ = D3_DC8(d_850_v64_, self.f28, (self).f44)
                d_852_v66_: _dafny.Array
                nw132_ = _dafny.Array(None, 11)
                nw132_[int(0)] = (d_849_v63_).cf10
                nw132_[int(1)] = self.f31
                nw132_[int(2)] = (d_822_v37_) != (_dafny.MultiSet([self.f31]))
                nw132_[int(3)] = False
                nw132_[int(4)] = ((self).f34) != (((d_851_v65_).cf15).cardinality)
                nw132_[int(5)] = self.f35
                nw132_[int(6)] = self.f29
                nw132_[int(7)] = self.f35
                nw132_[int(8)] = self.f29
                nw132_[int(9)] = (len(d_840_v54_)) <= ((self).f30)
                nw132_[int(10)] = self.f28
                d_852_v66_ = nw132_
                index96_ = default__.safeIndex(467, (d_852_v66_).length(0))
                (d_852_v66_)[index96_] = (self.f35) or (self.f29)
                d_853_v67_: _dafny.Set
                d_853_v67_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([self.f36 for d_854_i2_ in range(default__.abs(895))]))})
                index97_ = default__.safeIndex(169, (d_843_v57_).length(0))
                index98_ = default__.safeIndex(467, (d_852_v66_).length(0))
                def iife49_():
                    coll23_ = _dafny.Set()
                    compr_23_: int
                    for compr_23_ in (d_853_v67_).Elements:
                        d_855_v68_: int = compr_23_
                        if (d_855_v68_) in (d_853_v67_):
                            coll23_ = coll23_.union(_dafny.Set([default__.safeModuloInt(d_855_v68_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aohlb"))))]))
                    return _dafny.Set(coll23_)
                rhs100_ = (self).f44
                rhs101_ = (iife49_()
                ).issubset(_dafny.Set({(self).f30}))
                lhs76_ = d_843_v57_
                lhs77_ = default__.safeIndex(169, (d_843_v57_).length(0))
                lhs78_ = d_852_v66_
                lhs79_ = default__.safeIndex(467, (d_852_v66_).length(0))
                lhs76_[lhs77_] = rhs100_
                lhs78_[lhs79_] = rhs101_
                (globalState).f13 = ((d_849_v63_).cf11) * ((d_843_v57_)[default__.safeIndex(169, (d_843_v57_).length(0))])
                d_856_v69_: _dafny.Map
                d_856_v69_ = _dafny.Map({(d_852_v66_)[default__.safeIndex(467, (d_852_v66_).length(0))]: (d_843_v57_)[default__.safeIndex(169, (d_843_v57_).length(0))]})
                d_857_v70_: _dafny.MultiSet
                d_857_v70_ = _dafny.MultiSet([d_856_v69_, d_856_v69_, (d_856_v69_ if self.f35 else d_856_v69_)])
                d_857_v70_ = d_857_v70_
                (globalState).f12 = (d_852_v66_)[default__.safeIndex(467, (d_852_v66_).length(0))]
            d_858_v71_: _dafny.Array
            def lambda31_(d_859_i3_):
                return (self.f28 if self.f29 else self.f35)

            init16_ = lambda31_
            nw133_ = _dafny.Array(None, 1)
            for i0_16_ in range(nw133_.length(0)):
                nw133_[i0_16_] = init16_(i0_16_)
            d_858_v71_ = nw133_
            index99_ = default__.safeIndex(727, (d_858_v71_).length(0))
            (d_858_v71_)[index99_] = self.f29
            index100_ = default__.safeIndex(727, (d_858_v71_).length(0))
            (d_858_v71_)[index100_] = (not (self.f28) or (self.f28)) not in ((self).f32)
            (globalState).f13 = default__.fm32(True, globalState)
            d_860_v72_: _dafny.Map
            d_860_v72_ = _dafny.Map({(self).f30: self.f31})
            (self).f36 = default__.fm34(len((d_860_v72_).set((self).f44, self.f28)), (self).f34, (self).f44, self.f31, globalState)
            index101_ = default__.safeIndex(146, (d_843_v57_).length(0))
            (d_843_v57_)[index101_] = (self).f30
            index102_ = default__.safeIndex(146, (d_843_v57_).length(0))
            (d_843_v57_)[index102_] = default__.safeModuloInt(default__.fm32((d_858_v71_)[default__.safeIndex(727, (d_858_v71_).length(0))], globalState), (self).f30)
        elif True:
            (globalState).f13 = (default__.fm32(self.f28, globalState)) * ((self).f30)
            r0 = self.f28
            d_861_v73_: _dafny.Seq
            d_861_v73_ = _dafny.SeqWithoutIsStrInference([(self).f34, 307, default__.safeDivisionInt((self).f34, 381), (self).f44, (self).f30])
            d_862_v74_: _dafny.MultiSet
            d_862_v74_ = _dafny.MultiSet([d_861_v73_])
            rhs102_ = _dafny.SeqWithoutIsStrInference([((d_862_v74_).cardinality if self.f31 else (self).f44), (0) - ((self).f30), (self).f44, (self).f34, default__.safeModuloInt((self).f34, -445)])
            rhs103_ = ((self).f30) not in (_dafny.MultiSet((d_861_v73_) + (d_861_v73_)))
            rhs104_ = not(((self).f44) < (((self).f30) + ((self).f34)))
            lhs80_ = self
            lhs81_ = globalState
            d_861_v73_ = rhs102_
            lhs80_.f29 = rhs103_
            lhs81_.f26 = rhs104_
            d_863_v75_: _dafny.Array
            nw134_ = _dafny.Array(None, 10)
            nw134_[int(0)] = True
            nw134_[int(1)] = self.f28
            nw134_[int(2)] = (self).fm7(_dafny.MultiSet(d_861_v73_), (self).f30, globalState)
            nw134_[int(3)] = not(self.f29)
            nw134_[int(4)] = True
            nw134_[int(5)] = not(self.f29)
            nw134_[int(6)] = self.f28
            nw134_[int(7)] = self.f31
            nw134_[int(8)] = self.f28
            nw134_[int(9)] = self.f35
            d_863_v75_ = nw134_
            d_864_v76_: _dafny.Map
            d_864_v76_ = _dafny.Map({(self).f44: d_863_v75_})
            d_865_v77_: _dafny.Map
            d_865_v77_ = _dafny.Map({(self).f44: d_840_v54_})
            d_864_v76_ = (d_864_v76_).set(((self).f34) - (len((((self).f32).set(default__.safeIndex(len(d_865_v77_), len((self).f32)), self.f29)).set(default__.safeIndex((self).f30, len(((self).f32).set(default__.safeIndex(len(d_865_v77_), len((self).f32)), self.f29))), self.f31))), d_863_v75_)
            (globalState).f13 = (self).f30
        r2 = d_840_v54_
        d_866_v78_: _dafny.Map
        d_866_v78_ = _dafny.Map({(d_842_v56_) != (d_842_v56_): self.f28})
        d_867_v79_: _dafny.Map
        d_867_v79_ = _dafny.Map({(self).f30: (self).f30})
        r0 = ((d_866_v78_)[((self).f30) in (d_867_v79_)] if (((self).f30) in (d_867_v79_)) in (d_866_v78_) else self.f28)
        d_868_v80_: _dafny.Map
        d_868_v80_ = _dafny.Map({self.f36: not(self.f31)})
        d_869_v81_: _dafny.Seq
        d_869_v81_ = _dafny.SeqWithoutIsStrInference([len(d_868_v80_)])
        d_870_v82_: D6
        d_870_v82_ = D6_DC15(d_869_v81_)
        d_871_v83_: _dafny.Map
        d_871_v83_ = _dafny.Map({d_870_v82_: (self).f34})
        d_872_v84_: _dafny.Seq
        d_872_v84_ = _dafny.SeqWithoutIsStrInference([(self).f44, ((d_871_v83_)[d_870_v82_] if (d_870_v82_) in (d_871_v83_) else 601)])
        r1 = (d_872_v84_)[default__.safeIndex((self).f34, len(d_872_v84_))]
        r2 = ((_dafny.SeqWithoutIsStrInference([self.f36 for d_873_i4_ in range(default__.abs(926))])).set(default__.safeIndex((self).f44, len(_dafny.SeqWithoutIsStrInference([self.f36 for d_874_i4_ in range(default__.abs(926))]))), self.f36)) + (((d_840_v54_) + (d_840_v54_)).set(default__.safeIndex((self).f44, len((d_840_v54_) + (d_840_v54_))), self.f36))
        return r0, r1, r2

    def m12(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        d_875_v0_: _dafny.Array
        nw135_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
        d_875_v0_ = nw135_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_875_v0_).length(0)):
            d_876_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_876_i0_)) and ((d_876_i0_) < ((d_875_v0_).length(0)))):
                (d_875_v0_)[(d_876_i0_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pftnlhv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mekvfibbt")))
        d_877_v1_: _dafny.Seq
        d_877_v1_ = _dafny.SeqWithoutIsStrInference([(self).f34])
        d_878_v2_: int
        out52_: int
        out52_ = default__.m0((_dafny.SeqWithoutIsStrInference([(0) - ((self).f30), (self).f44])) + (d_877_v1_), self.f35, (self).f44, self.f31, globalState)
        d_878_v2_ = out52_
        d_879_v3_: D3
        d_879_v3_ = D3_DC8(_dafny.MultiSet([self.f36]), self.f28, p0)
        if (self.f29) or ((d_879_v3_).cf16):
            d_880_v4_: _dafny.Array
            def lambda32_(d_881_i1_):
                return self.f36

            init17_ = lambda32_
            nw136_ = _dafny.Array(None, 16)
            for i0_17_ in range(nw136_.length(0)):
                nw136_[i0_17_] = init17_(i0_17_)
            d_880_v4_ = nw136_
            d_882_v5_: C0
            nw137_ = C0()
            nw137_.ctor__(d_880_v4_)
            d_882_v5_ = nw137_
            d_883_v6_: C0
            nw138_ = C0()
            nw138_.ctor__(d_882_v5_.f37)
            d_883_v6_ = nw138_
            d_884_v7_: _dafny.Seq
            d_884_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaganir"))
            (globalState).f13 = ((0) - (len(d_884_v7_))) + ((self).f34)
            d_885_v8_: _dafny.Set
            d_885_v8_ = _dafny.Set({d_880_v4_})
            d_886_v9_: int
            out53_: int
            out53_ = default__.m0(_dafny.SeqWithoutIsStrInference([(self).f44]), not((d_885_v8_).ispropersubset(d_885_v8_)), p0, self.f28, globalState)
            d_886_v9_ = out53_
            d_887_v10_: _dafny.Map
            d_887_v10_ = _dafny.Map({p0: self.f35})
            d_888_v11_: C2
            nw139_ = C2()
            nw139_.ctor__((self).f30, (0) - (len(d_887_v10_)), self.f35, (self.f31) or (self.f29), self.f28)
            d_888_v11_ = nw139_
        elif True:
            d_889_v12_: _dafny.Array
            nw140_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_889_v12_ = nw140_
            d_890_v13_: _dafny.Array
            def lambda33_(d_891_i2_):
                return self.f29

            init18_ = lambda33_
            nw141_ = _dafny.Array(None, 17)
            for i0_18_ in range(nw141_.length(0)):
                nw141_[i0_18_] = init18_(i0_18_)
            d_890_v13_ = nw141_
            nw142_ = _dafny.Array(None, 13)
            nw142_[int(0)] = d_890_v13_
            nw142_[int(1)] = d_890_v13_
            nw142_[int(2)] = d_890_v13_
            nw142_[int(3)] = d_890_v13_
            nw142_[int(4)] = d_890_v13_
            nw142_[int(5)] = d_890_v13_
            nw142_[int(6)] = d_890_v13_
            nw142_[int(7)] = d_890_v13_
            nw142_[int(8)] = d_890_v13_
            nw142_[int(9)] = d_890_v13_
            nw142_[int(10)] = d_890_v13_
            nw142_[int(11)] = d_890_v13_
            nw142_[int(12)] = d_890_v13_
            d_889_v12_ = nw142_
            d_892_v14_: _dafny.Seq
            d_892_v14_ = _dafny.SeqWithoutIsStrInference([self.f28, self.f35])
            d_893_v15_: _dafny.MultiSet
            d_893_v15_ = _dafny.MultiSet([(self).f34])
            d_894_v16_: _dafny.Array
            nw143_ = _dafny.Array(int(0), 22)
            d_894_v16_ = nw143_
            d_895_v17_: _dafny.Map
            d_895_v17_ = _dafny.Map({d_894_v16_: d_890_v13_})
            d_896_v18_: _dafny.Seq
            d_896_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
            d_897_v19_: _dafny.MultiSet
            d_897_v19_ = _dafny.MultiSet([d_896_v18_])
            rhs105_ = ((_dafny.MultiSet([d_896_v18_])).intersection((d_897_v19_).set(d_896_v18_, default__.abs((self).f34)))).cardinality
            rhs106_ = self.f28
            rhs107_ = (self).f32
            rhs108_ = (d_893_v15_).set(d_878_v2_, default__.abs(len(_dafny.SeqWithoutIsStrInference([p0, (self).f30, (d_877_v1_)[default__.safeIndex(p0, len(d_877_v1_))]]))))
            rhs109_ = d_895_v17_
            lhs82_ = globalState
            d_878_v2_ = rhs105_
            lhs82_.f1 = rhs106_
            d_892_v14_ = rhs107_
            d_893_v15_ = rhs108_
            d_895_v17_ = rhs109_
            index103_ = default__.safeIndex(604, (d_890_v13_).length(0))
            (d_890_v13_)[index103_] = not((((self).f32)[default__.safeIndex((self).f44, len((self).f32))]) == (self.f29))
            index104_ = default__.safeIndex(604, (d_890_v13_).length(0))
            rhs110_ = self.f35
            rhs111_ = (d_892_v14_) + (((self).f32) + ((self).f32))
            rhs112_ = d_878_v2_
            rhs113_ = (d_896_v18_)[default__.safeIndex(d_878_v2_, len(d_896_v18_))]
            rhs114_ = self.f31
            lhs83_ = globalState
            lhs84_ = self
            lhs85_ = d_890_v13_
            lhs86_ = default__.safeIndex(604, (d_890_v13_).length(0))
            lhs83_.f26 = rhs110_
            d_892_v14_ = rhs111_
            d_878_v2_ = rhs112_
            lhs84_.f36 = rhs113_
            lhs85_[lhs86_] = rhs114_
            d_898_v20_: C1
            nw144_ = C1()
            nw144_.ctor__(self.f35, ((self).f32)[default__.safeIndex(729, len((self).f32))], self.f31)
            d_898_v20_ = nw144_
            (globalState).f1 = self.f28
        d_899_i3_: int
        d_899_i3_ = 0
        with _dafny.label("3"):
            while self.f29:
                with _dafny.c_label("3"):
                    if (d_899_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_899_i3_ = (d_899_i3_) + (1)
                    d_900_v21_: _dafny.Seq
                    d_900_v21_ = _dafny.SeqWithoutIsStrInference([self.f28])
                    rhs115_ = ((self).f32 if self.f28 else d_900_v21_)
                    rhs116_ = d_878_v2_
                    rhs117_ = not (self.f31) or (self.f28)
                    lhs87_ = globalState
                    d_900_v21_ = rhs115_
                    d_878_v2_ = rhs116_
                    lhs87_.f12 = rhs117_
                    d_901_v22_: _dafny.Array
                    def lambda34_(d_902_p0_):
                        def lambda35_(d_903_i4_):
                            return (_dafny.MultiSet([(_dafny.MultiSet([_dafny.Map({self.f35: self.f28})])).cardinality])).issubset(_dafny.MultiSet([d_902_p0_]))

                        return lambda35_

                    init19_ = lambda34_(p0)
                    nw145_ = _dafny.Array(None, 18)
                    for i0_19_ in range(nw145_.length(0)):
                        nw145_[i0_19_] = init19_(i0_19_)
                    d_901_v22_ = nw145_
                    index105_ = default__.safeIndex(595, (d_901_v22_).length(0))
                    (d_901_v22_)[index105_] = self.f29
                    index106_ = default__.safeIndex(595, (d_901_v22_).length(0))
                    (d_901_v22_)[index106_] = self.f28
                    d_904_v23_: _dafny.MultiSet
                    d_904_v23_ = _dafny.MultiSet([(self).f34, (self).f34, default__.fm32(not((d_901_v22_)[default__.safeIndex(595, (d_901_v22_).length(0))]), globalState), (self).f44])
                    r1 = (self).fm7(d_904_v23_, (self).f44, globalState)
                    d_905_v24_: _dafny.Array
                    nw146_ = _dafny.Array(_dafny.CodePoint('D'), 4)
                    d_905_v24_ = nw146_
                    index107_ = default__.safeIndex(246, (d_905_v24_).length(0))
                    (d_905_v24_)[index107_] = self.f36
                    d_906_v25_: D5
                    d_906_v25_ = D5_DC13(self.f29, _dafny.Map({(d_901_v22_)[default__.safeIndex(595, (d_901_v22_).length(0))]: (self).f30}), _dafny.CodePoint('y'))
                    index108_ = default__.safeIndex(246, (d_905_v24_).length(0))
                    (d_905_v24_)[index108_] = (d_906_v25_).cf29
                    pass
            pass
        d_907_v26_: C2
        nw147_ = C2()
        nw147_.ctor__((self).f34, (0) - (d_878_v2_), not(self.f35), not(self.f31), self.f31)
        d_907_v26_ = nw147_
        d_908_v27_: _dafny.Seq
        d_908_v27_ = _dafny.SeqWithoutIsStrInference([self.f36, self.f36])
        d_909_v28_: _dafny.Map
        d_909_v28_ = _dafny.Map({self.f36: (d_908_v27_) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_910_i5_ in range(default__.abs(377))]))})
        d_911_v29_: _dafny.Array
        nw148_ = _dafny.Array(_dafny.MultiSet({}), 9)
        d_911_v29_ = nw148_
        d_912_v30_: _dafny.MultiSet
        d_912_v30_ = _dafny.MultiSet([self.f36])
        d_913_v31_: _dafny.Seq
        d_913_v31_ = _dafny.SeqWithoutIsStrInference([d_911_v29_])
        rhs118_ = ((d_912_v30_)[self.f36] if (self.f36) in (d_912_v30_) else (len(_dafny.Map({self.f35: self.f35}))) * (p0))
        rhs119_ = d_909_v28_
        rhs120_ = (d_913_v31_)[default__.safeIndex(len((self).f32), len(d_913_v31_))]
        lhs88_ = globalState
        lhs88_.f13 = rhs118_
        d_909_v28_ = rhs119_
        d_911_v29_ = rhs120_
        d_914_v32_: _dafny.Set
        d_914_v32_ = _dafny.Set({d_878_v2_})
        r0 = (d_914_v32_) | ((d_914_v32_) | (d_914_v32_))
        r1 = self.f28
        return r0, r1

    @property
    def f44(self):
        return self._f44
    @property
    def f45(self):
        return self._f45

class C4(T3, T2, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f34, f35, f32, f33, f30, f31, f28, f29):
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm6(self, globalState):
        return not(self.f31)

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife50_():
            coll24_ = _dafny.Set()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(460, 658):
                d_915_v0_: int = compr_24_
                if ((460) <= (d_915_v0_)) and ((d_915_v0_) < (658)):
                    coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_915_v0_, (self).f34)]))
            return _dafny.Set(coll24_)
        return ((iife50_()
        ) - (_dafny.Set({972}))) | ((_dafny.Set({(self).f34})) | (_dafny.Set({len(_dafny.Set({self.f28}))})))

    def fm5(self, p0, p1, p2, p3, globalState):
        return (self).f32

    def fm3(self, p0, p1, globalState):
        if (_dafny.MultiSet([(self).f30])).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f30 for d_916_i0_ in range(default__.abs(845))]))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe"))
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjcbfbtrq"))

    def fm2(self, globalState):
        return _dafny.Map({(_dafny.Set({not(not(((self).f32)[default__.safeIndex((self).f34, len((self).f32))]))})).issubset(_dafny.Set({self.f31})): (self).f34})

    def fm44(self, p0, globalState):
        return (default__.safeModuloInt((self).f30, (0) - ((self).f30))) > ((self).f30)

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (self).f35 = self.f31
        d_917_i0_: int
        d_917_i0_ = 0
        with _dafny.label("4"):
            while (self).fm6(globalState):
                with _dafny.c_label("4"):
                    if (d_917_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_917_i0_ = (d_917_i0_) + (1)
                    d_918_v0_: _dafny.MultiSet
                    d_918_v0_ = _dafny.MultiSet([self.f28])
                    d_919_v1_: _dafny.Set
                    d_919_v1_ = _dafny.Set({self.f31})
                    d_920_v2_: _dafny.Array
                    nw149_ = _dafny.Array(None, 20)
                    nw149_[int(0)] = self.f31
                    nw149_[int(1)] = self.f28
                    nw149_[int(2)] = self.f28
                    nw149_[int(3)] = self.f29
                    nw149_[int(4)] = self.f28
                    nw149_[int(5)] = self.f29
                    nw149_[int(6)] = self.f28
                    nw149_[int(7)] = self.f28
                    nw149_[int(8)] = self.f28
                    nw149_[int(9)] = False
                    nw149_[int(10)] = self.f29
                    nw149_[int(11)] = self.f29
                    nw149_[int(12)] = self.f35
                    nw149_[int(13)] = self.f29
                    nw149_[int(14)] = self.f29
                    nw149_[int(15)] = True
                    nw149_[int(16)] = True
                    nw149_[int(17)] = (self).fm6(globalState)
                    nw149_[int(18)] = self.f35
                    nw149_[int(19)] = (self).fm6(globalState)
                    d_920_v2_ = nw149_
                    d_921_v3_: _dafny.Map
                    d_921_v3_ = _dafny.Map({d_920_v2_: d_919_v1_})
                    d_922_v4_: _dafny.Seq
                    d_922_v4_ = _dafny.SeqWithoutIsStrInference([d_919_v1_])
                    d_923_v5_: _dafny.Array
                    nw150_ = _dafny.Array(None, 23)
                    nw150_[int(0)] = default__.fm45((self).f30, d_918_v0_, self.f35, len((self).f32), globalState)
                    nw150_[int(1)] = ((d_919_v1_)) - (_dafny.Set({self.f29, False, self.f29, self.f35, self.f31}))
                    nw150_[int(2)] = d_919_v1_
                    nw150_[int(3)] = ((d_921_v3_)[d_920_v2_] if (d_920_v2_) in (d_921_v3_) else d_919_v1_)
                    nw150_[int(4)] = (d_919_v1_) - (d_919_v1_)
                    nw150_[int(5)] = d_919_v1_
                    nw150_[int(6)] = d_919_v1_
                    nw150_[int(7)] = d_919_v1_
                    nw150_[int(8)] = d_919_v1_
                    nw150_[int(9)] = (_dafny.Set({self.f35})) - (default__.fm45((0) - ((self).f30), d_918_v0_, self.f35, (self).f34, globalState))
                    nw150_[int(10)] = (d_919_v1_).intersection(d_919_v1_)
                    nw150_[int(11)] = (d_922_v4_)[default__.safeIndex(133, len(d_922_v4_))]
                    nw150_[int(12)] = d_919_v1_
                    nw150_[int(13)] = d_919_v1_
                    nw150_[int(14)] = d_919_v1_
                    nw150_[int(15)] = d_919_v1_
                    nw150_[int(16)] = d_919_v1_
                    nw150_[int(17)] = d_919_v1_
                    nw150_[int(18)] = default__.fm45((self).f30, d_918_v0_, self.f35, len(_dafny.Map({(self).fm6(globalState): (self).f34})), globalState)
                    nw150_[int(19)] = _dafny.Set({True})
                    nw150_[int(20)] = d_919_v1_
                    nw150_[int(21)] = d_919_v1_
                    nw150_[int(22)] = d_919_v1_
                    d_923_v5_ = nw150_
                    d_923_v5_ = d_923_v5_
                    d_924_v6_: _dafny.MultiSet
                    d_924_v6_ = _dafny.MultiSet([(self).f30, 355])
                    if ((self).f32)[default__.safeIndex(default__.fm46((self).f34, (d_924_v6_).cardinality, self.f28, globalState), len((self).f32))]:
                        d_925_v7_: _dafny.Set
                        d_925_v7_ = _dafny.Set({(self).f34, (self).f34})
                        r1 = default__.safeDivisionInt(len((d_925_v7_) - (d_925_v7_)), (d_924_v6_).cardinality)
                        d_926_v8_: _dafny.Map
                        d_926_v8_ = _dafny.Map({(self).f34: self.f29})
                        d_927_v9_: _dafny.Array
                        nw151_ = _dafny.Array(None, 12)
                        nw151_[int(0)] = len((self).f32)
                        nw151_[int(1)] = (self).f30
                        nw151_[int(2)] = (self).f30
                        nw151_[int(3)] = ((self).f34) * ((self).f30)
                        nw151_[int(4)] = len(d_926_v8_)
                        nw151_[int(5)] = (self).f30
                        nw151_[int(6)] = (self).f30
                        nw151_[int(7)] = default__.safeDivisionInt((self).f30, (self).f30)
                        nw151_[int(8)] = (self).f34
                        nw151_[int(9)] = ((self).f34 if self.f29 else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khil"))))
                        nw151_[int(10)] = (0) - (default__.safeModuloInt((self).f34, (self).f34))
                        nw151_[int(11)] = (self).f34
                        d_927_v9_ = nw151_
                        index109_ = default__.safeIndex(655, (d_927_v9_).length(0))
                        (d_927_v9_)[index109_] = ((self).f30) * ((0) - ((self).f30))
                        index110_ = default__.safeIndex(655, (d_927_v9_).length(0))
                        (d_927_v9_)[index110_] = 802
                        (globalState).f13 = len((d_925_v7_) | (d_925_v7_))
                        index111_ = default__.safeIndex(655, (d_927_v9_).length(0))
                        (d_927_v9_)[index111_] = (0) - ((0) - ((d_927_v9_)[default__.safeIndex(655, (d_927_v9_).length(0))]))
                        (globalState).f1 = ((self).f30) < ((d_927_v9_)[default__.safeIndex(655, (d_927_v9_).length(0))])
                    elif True:
                        index112_ = default__.safeIndex(443, (d_920_v2_).length(0))
                        (d_920_v2_)[index112_] = ((self).f30) == ((self).f30)
                        index113_ = default__.safeIndex(443, (d_920_v2_).length(0))
                        (d_920_v2_)[index113_] = self.f29
                        (globalState).f13 = ((d_924_v6_) - (d_924_v6_)).cardinality
                        (globalState).f13 = (self).f30
                        r1 = len((self).f32)
                        index114_ = default__.safeIndex(443, (d_920_v2_).length(0))
                        (d_920_v2_)[index114_] = self.f28
                    d_928_v10_: str
                    d_928_v10_ = _dafny.CodePoint('p')
                    d_928_v10_ = d_928_v10_
                    d_929_v11_: _dafny.Array
                    nw152_ = _dafny.Array(None, 18)
                    nw152_[int(0)] = d_920_v2_
                    nw152_[int(1)] = d_920_v2_
                    nw152_[int(2)] = d_920_v2_
                    nw152_[int(3)] = d_920_v2_
                    nw152_[int(4)] = d_920_v2_
                    nw152_[int(5)] = d_920_v2_
                    nw152_[int(6)] = d_920_v2_
                    nw152_[int(7)] = d_920_v2_
                    nw152_[int(8)] = d_920_v2_
                    nw152_[int(9)] = d_920_v2_
                    nw152_[int(10)] = d_920_v2_
                    nw152_[int(11)] = d_920_v2_
                    nw152_[int(12)] = d_920_v2_
                    nw152_[int(13)] = d_920_v2_
                    nw152_[int(14)] = d_920_v2_
                    nw152_[int(15)] = d_920_v2_
                    nw152_[int(16)] = d_920_v2_
                    nw152_[int(17)] = d_920_v2_
                    d_929_v11_ = nw152_
                    index115_ = default__.safeIndex(658, (d_929_v11_).length(0))
                    (d_929_v11_)[index115_] = d_920_v2_
                    d_930_v12_: _dafny.MultiSet
                    d_930_v12_ = _dafny.MultiSet([((self).f32) + ((self).f32), ((self).f32) + ((self).f32), (self).f32, ((self).f32) + ((self).fm5((self).f30, (self).f34, not(self.f35), (self).f30, globalState))])
                    d_931_v13_: _dafny.Set
                    d_931_v13_ = _dafny.Set({(self).f34})
                    d_932_v14_: _dafny.Set
                    d_932_v14_ = d_931_v13_
                    d_933_v15_: _dafny.Seq
                    d_933_v15_ = _dafny.SeqWithoutIsStrInference([(d_932_v14_), d_931_v13_, d_931_v13_])
                    d_934_v16_: _dafny.Map
                    d_934_v16_ = _dafny.Map({self.f35: (d_931_v13_).issubset((d_933_v15_)[default__.safeIndex(459, len(d_933_v15_))])})
                    index116_ = default__.safeIndex(658, (d_929_v11_).length(0))
                    rhs121_ = d_920_v2_
                    rhs122_ = (d_930_v12_) - ((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])).set(default__.safeIndex((self).f30, len(_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32]))), (self).f32))) | (d_930_v12_))
                    rhs123_ = d_931_v13_
                    rhs124_ = (d_934_v16_) | (d_934_v16_)
                    lhs89_ = d_929_v11_
                    lhs90_ = default__.safeIndex(658, (d_929_v11_).length(0))
                    lhs89_[lhs90_] = rhs121_
                    d_930_v12_ = rhs122_
                    d_931_v13_ = rhs123_
                    d_934_v16_ = rhs124_
                    pass
            pass
        r1 = (self).f30
        if self.f35:
            d_935_v17_: _dafny.Array
            nw153_ = _dafny.Array(_dafny.Map({}), 12)
            d_935_v17_ = nw153_
            d_936_v18_: str
            d_936_v18_ = _dafny.CodePoint('p')
            d_937_v19_: _dafny.Map
            d_937_v19_ = _dafny.Map({768: d_936_v18_})
            d_938_v20_: _dafny.Map
            d_938_v20_ = _dafny.Map({(self).f30: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qefsgg"))})
            index117_ = default__.safeIndex(445, (d_935_v17_).length(0))
            (d_935_v17_)[index117_] = (default__.fm47(self.f29, d_937_v19_, globalState)) | (d_938_v20_)
            d_939_v21_: _dafny.Array
            nw154_ = _dafny.Array(False, 26)
            d_939_v21_ = nw154_
            d_940_v22_: _dafny.Map
            d_940_v22_ = _dafny.Map({d_939_v21_: self.f28})
            d_941_v23_: _dafny.Seq
            d_941_v23_ = _dafny.SeqWithoutIsStrInference([(self).f34])
            d_942_v24_: D8
            d_942_v24_ = D8_DC24(self.f28, (self).f34, (self).f32, self.f31, d_941_v23_)
            d_943_v25_: _dafny.Array
            nw155_ = _dafny.Array(None, 9)
            nw155_[int(0)] = (len(d_940_v22_)) > ((self).f34)
            nw155_[int(1)] = self.f31
            nw155_[int(2)] = (True if self.f31 else default__.fm1(self.f31, (self).f30, globalState))
            nw155_[int(3)] = (711) == (238)
            nw155_[int(4)] = ((d_941_v23_)[default__.safeIndex((self).f34, len(d_941_v23_))]) not in ((d_942_v24_).cf50)
            nw155_[int(5)] = True
            nw155_[int(6)] = self.f29
            nw155_[int(7)] = self.f31
            nw155_[int(8)] = ((self).f32) != (_dafny.SeqWithoutIsStrInference([self.f31, self.f31]))
            d_943_v25_ = nw155_
            index118_ = default__.safeIndex(445, (d_935_v17_).length(0))
            rhs125_ = self.f31
            rhs126_ = ((self).f30) + (845)
            rhs127_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yowl")))) + ((self).f34): _dafny.SeqWithoutIsStrInference([d_936_v18_ for d_944_i1_ in range(default__.abs(526))])})
            rhs128_ = d_943_v25_
            rhs129_ = True
            lhs91_ = globalState
            lhs92_ = d_935_v17_
            lhs93_ = default__.safeIndex(445, (d_935_v17_).length(0))
            lhs94_ = globalState
            lhs95_ = globalState
            lhs91_.f26 = rhs125_
            r1 = rhs126_
            lhs92_[lhs93_] = rhs127_
            lhs94_.f21 = rhs128_
            lhs95_.f12 = rhs129_
            (globalState).f13 = ((self).f34) + ((self).f34)
            d_945_v26_: _dafny.Seq
            d_945_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uiiih"))
            (self).f31 = ((self).f32)[default__.safeIndex((0) - ((0) - (len((d_945_v26_) + (d_945_v26_)))), len((self).f32))]
            d_946_v27_: _dafny.Map
            d_946_v27_ = _dafny.Map({(self).f30: self.f29})
            d_947_v28_: _dafny.Array
            nw156_ = _dafny.Array(None, 5)
            nw156_[int(0)] = len(_dafny.Set({931}))
            nw156_[int(1)] = (0) - ((self).f34)
            nw156_[int(2)] = len((d_946_v27_) | (_dafny.Map({-557: self.f35})))
            nw156_[int(3)] = default__.safeModuloInt((self).f30, (self).f34)
            nw156_[int(4)] = (self).f30
            d_947_v28_ = nw156_
            index119_ = default__.safeIndex(748, (d_947_v28_).length(0))
            (d_947_v28_)[index119_] = len(_dafny.Map({not(self.f28): (self).f34}))
            index120_ = default__.safeIndex(748, (d_947_v28_).length(0))
            (d_947_v28_)[index120_] = (self).f30
            d_948_v29_: _dafny.Array
            nw157_ = _dafny.Array(_dafny.Seq({}), 24)
            d_948_v29_ = nw157_
            d_949_v30_: _dafny.MultiSet
            d_949_v30_ = _dafny.MultiSet([len(d_941_v23_)])
            d_950_v31_: _dafny.Set
            d_950_v31_ = _dafny.Set({(self).f34, (d_949_v30_).cardinality})
            d_951_v32_: D1
            d_951_v32_ = D1_DC2(d_939_v21_)
            d_952_v33_: D1
            d_952_v33_ = D1_DC4(d_951_v32_)
            d_953_v34_: D1
            d_953_v34_ = D1_DC4(d_951_v32_)
            d_954_v35_: _dafny.Map
            d_954_v35_ = _dafny.Map({(self).f34: d_953_v34_})
            d_955_v36_: _dafny.Map
            d_955_v36_ = _dafny.Map({len(d_950_v31_): d_954_v35_})
            d_956_v37_: _dafny.Map
            d_956_v37_ = _dafny.Map({self.f35: (0) - (len(d_955_v36_))})
            d_957_v38_: _dafny.Array
            nw158_ = _dafny.Array(None, 27)
            nw158_[int(0)] = _dafny.CodePoint('t')
            nw158_[int(1)] = d_936_v18_
            nw158_[int(2)] = d_936_v18_
            nw158_[int(3)] = d_936_v18_
            nw158_[int(4)] = d_936_v18_
            nw158_[int(5)] = d_936_v18_
            nw158_[int(6)] = d_936_v18_
            nw158_[int(7)] = d_936_v18_
            nw158_[int(8)] = ((d_937_v19_)[125] if (125) in (d_937_v19_) else d_936_v18_)
            nw158_[int(9)] = d_936_v18_
            nw158_[int(10)] = _dafny.CodePoint('o')
            nw158_[int(11)] = _dafny.CodePoint('b')
            nw158_[int(12)] = d_936_v18_
            nw158_[int(13)] = d_936_v18_
            nw158_[int(14)] = d_936_v18_
            nw158_[int(15)] = d_936_v18_
            nw158_[int(16)] = d_936_v18_
            nw158_[int(17)] = d_936_v18_
            nw158_[int(18)] = d_936_v18_
            nw158_[int(19)] = d_936_v18_
            nw158_[int(20)] = default__.fm48(True, default__.fm1(self.f28, (self).f34, globalState), len((self).f32), globalState)
            nw158_[int(21)] = d_936_v18_
            nw158_[int(22)] = d_936_v18_
            nw158_[int(23)] = d_936_v18_
            nw158_[int(24)] = _dafny.CodePoint('l')
            nw158_[int(25)] = d_936_v18_
            nw158_[int(26)] = d_936_v18_
            d_957_v38_ = nw158_
            d_958_v39_: C0
            nw159_ = C0()
            nw159_.ctor__(d_957_v38_)
            d_958_v39_ = nw159_
            d_959_v40_: D1
            d_959_v40_ = D1_DC3(self.f35, (self).f30, d_958_v39_, len((self).f32))
            d_960_v41_: C3
            nw160_ = C3()
            nw160_.ctor__((self).f34, d_948_v29_, d_936_v18_, (self).f34, not(self.f35), ((self).f32).set(default__.safeIndex((self).f34, len((self).f32)), self.f29), (self).f33, ((d_956_v37_)[self.f29] if (self.f29) in (d_956_v37_) else (self).f30), (d_959_v40_).cf3, not(self.f29), self.f31)
            d_960_v41_ = nw160_
        elif True:
            d_961_v42_: _dafny.Set
            d_961_v42_ = _dafny.Set({self.f31})
            d_962_v43_: _dafny.Set
            d_962_v43_ = d_961_v42_
            d_963_v44_: str
            d_963_v44_ = _dafny.CodePoint('k')
            d_964_v45_: _dafny.Seq
            d_964_v45_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f34)])
            d_965_v46_: D4
            d_965_v46_ = D4_DC10(default__.fm49(_dafny.MultiSet([977, len((d_962_v43_))]), d_963_v44_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - ((self).f34)]), d_964_v45_, d_964_v45_]), globalState))
            d_965_v46_ = d_965_v46_
            d_966_v47_: _dafny.Array
            nw161_ = _dafny.Array(False, 18)
            d_966_v47_ = nw161_
            (globalState).f21 = d_966_v47_
            (globalState).f1 = self.f35
            d_967_v48_: _dafny.Seq
            d_967_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hngyxcapn"))
            d_968_v49_: _dafny.Array
            nw162_ = _dafny.Array(None, 3)
            nw162_[int(0)] = d_963_v44_
            nw162_[int(1)] = d_963_v44_
            nw162_[int(2)] = (d_967_v48_)[default__.safeIndex((self).f30, len(d_967_v48_))]
            d_968_v49_ = nw162_
            index121_ = default__.safeIndex(724, (d_968_v49_).length(0))
            (d_968_v49_)[index121_] = d_963_v44_
            index122_ = default__.safeIndex(724, (d_968_v49_).length(0))
            (d_968_v49_)[index122_] = (d_963_v44_ if (self).fm6(globalState) else d_963_v44_)
            d_969_v50_: _dafny.Map
            d_969_v50_ = _dafny.Map({d_966_v47_: not(True)})
            (self).f28 = (len(d_969_v50_)) != ((0) - ((self).f34))
        d_970_v51_: _dafny.Set
        d_970_v51_ = _dafny.Set({(self).f30})
        d_971_v52_: _dafny.MultiSet
        d_971_v52_ = _dafny.MultiSet([(self).f34, (self).f30])
        d_972_v53_: _dafny.Map
        d_972_v53_ = _dafny.Map({self.f28: len((self).f32)})
        rhs130_ = default__.safeDivisionInt((self).f30, (self).f30)
        rhs131_ = default__.safeModuloInt(((d_971_v52_).intersection(d_971_v52_)).cardinality, (len(d_972_v53_) if self.f29 else (self).f30))
        rhs132_ = (d_970_v51_).intersection(d_970_v51_)
        lhs96_ = globalState
        r1 = rhs130_
        lhs96_.f13 = rhs131_
        d_970_v51_ = rhs132_
        d_973_v54_: _dafny.Seq
        d_973_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmjcaq"))
        d_974_v55_: _dafny.Set
        d_974_v55_ = _dafny.Set({d_973_v54_})
        d_975_v56_: str
        d_975_v56_ = _dafny.CodePoint('h')
        (globalState).f13 = len(((d_974_v55_).intersection(_dafny.Set({d_973_v54_, d_973_v54_, (d_973_v54_).set(default__.safeIndex((0) - ((self).f34), len(d_973_v54_)), d_975_v56_)}))) - ((d_974_v55_) | (d_974_v55_)))
        d_976_v57_: _dafny.Map
        d_976_v57_ = _dafny.Map({(self).f34: self.f31})
        r0 = not(((d_976_v57_)[((self).f30) + ((self).f34)] if (((self).f30) + ((self).f34)) in (d_976_v57_) else self.f35))
        r1 = (self).f30
        d_977_v58_: D5
        d_977_v58_ = D5_DC12(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpyunkfoa")))
        r2 = (d_973_v54_) + ((d_977_v58_).cf26)
        return r0, r1, r2


class C5(T3, T2, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        self._f47: bool = False
        self._f48: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f47, f48, f34, f35, f32, f33, f30, f31, f28, f29):
        (self)._f47 = f47
        (self)._f48 = f48
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm6(self, globalState):
        return ((self).f30) > ((self).f34)

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife51_():
            coll25_ = _dafny.Set()
            compr_25_: int
            for compr_25_ in (_dafny.SeqWithoutIsStrInference([-735])).Elements:
                d_978_v0_: int = compr_25_
                if (d_978_v0_) in (_dafny.SeqWithoutIsStrInference([-735])):
                    coll25_ = coll25_.union(_dafny.Set([(d_978_v0_) - (599)]))
            return _dafny.Set(coll25_)
        return (iife51_()
        ) | ((_dafny.Set({(self).f48})) - (_dafny.Set({255})))

    def fm5(self, p0, p1, p2, p3, globalState):
        return (self).f32

    def fm3(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_979_i0_ in range(default__.abs(741))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfcfykdh")))

    def fm2(self, globalState):
        return ((_dafny.Map({self.f29: (self).f48})) | (_dafny.Map({self.f28: (self).f48}))) | (_dafny.Map({self.f28: (self).f30}))

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1 = (self).f48
        d_980_v0_: str
        d_980_v0_ = _dafny.CodePoint('a')
        d_981_v1_: _dafny.MultiSet
        d_981_v1_ = _dafny.MultiSet([d_980_v0_, d_980_v0_])
        d_982_v2_: _dafny.MultiSet
        d_982_v2_ = _dafny.MultiSet([(self).f48])
        d_983_v3_: _dafny.Map
        d_983_v3_ = _dafny.Map({d_981_v1_: default__.fm39((self).f47, d_982_v2_, self.f31, globalState)})
        d_984_v4_: _dafny.Map
        d_984_v4_ = _dafny.Map({d_983_v3_: (self).f47})
        d_984_v4_ = (d_984_v4_).set(_dafny.Map({d_981_v1_: (self).f34}), not((self).fm6(globalState)))
        d_985_v5_: _dafny.Seq
        d_985_v5_ = _dafny.SeqWithoutIsStrInference([(self).f48])
        d_986_v6_: _dafny.Map
        d_986_v6_ = _dafny.Map({not(self.f28): (d_985_v5_) <= (_dafny.SeqWithoutIsStrInference([(self).f48 for d_987_i0_ in range(default__.abs(641))]))})
        d_988_v7_: _dafny.Seq
        d_988_v7_ = _dafny.SeqWithoutIsStrInference([((self).f32)[default__.safeIndex((self).f34, len((self).f32))]])
        d_989_v8_: _dafny.Array
        nw163_ = _dafny.Array(D4.default()(), 13)
        d_989_v8_ = nw163_
        d_990_v9_: D4
        d_990_v9_ = D4_DC11(len(_dafny.SeqWithoutIsStrInference([d_980_v0_ for d_991_i1_ in range(default__.abs(248))])), True, (self).f30)
        index123_ = default__.safeIndex(423, (d_989_v8_).length(0))
        (d_989_v8_)[index123_] = d_990_v9_
        index124_ = default__.safeIndex(423, (d_989_v8_).length(0))
        rhs133_ = d_986_v6_
        rhs134_ = d_988_v7_
        rhs135_ = D4_DC11((self).f30, ((self).f34) >= ((self).f48), -352)
        rhs136_ = not (self.f35) or (((self).f34) == ((self).f30))
        lhs97_ = d_989_v8_
        lhs98_ = default__.safeIndex(423, (d_989_v8_).length(0))
        lhs99_ = self
        d_986_v6_ = rhs133_
        d_988_v7_ = rhs134_
        lhs97_[lhs98_] = rhs135_
        lhs99_.f35 = rhs136_
        d_992_v10_: _dafny.Map
        d_992_v10_ = _dafny.Map({False: 941})
        d_993_v11_: _dafny.MultiSet
        d_993_v11_ = _dafny.MultiSet([self.f29])
        d_994_v12_: _dafny.Map
        d_994_v12_ = _dafny.Map({True: d_980_v0_})
        d_995_v13_: _dafny.Array
        nw164_ = _dafny.Array(None, 29)
        nw164_[int(0)] = (self).f30
        nw164_[int(1)] = (self).f30
        nw164_[int(2)] = (self).f48
        nw164_[int(3)] = (d_985_v5_)[default__.safeIndex(((d_982_v2_)[(self).f34] if ((self).f34) in (d_982_v2_) else (self).f30), len(d_985_v5_))]
        nw164_[int(4)] = (self).f30
        nw164_[int(5)] = (self).f30
        nw164_[int(6)] = (self).f34
        nw164_[int(7)] = (self).f34
        nw164_[int(8)] = 381
        nw164_[int(9)] = (self).f30
        nw164_[int(10)] = (0) - ((self).f34)
        nw164_[int(11)] = (self).f48
        nw164_[int(12)] = (self).f48
        nw164_[int(13)] = (self).f48
        nw164_[int(14)] = len(d_985_v5_)
        nw164_[int(15)] = 803
        nw164_[int(16)] = (self).f34
        nw164_[int(17)] = (self).f34
        nw164_[int(18)] = (self).f48
        nw164_[int(19)] = (_dafny.MultiSet([self.f29, False, default__.fm1(self.f28, ((d_992_v10_)[self.f28] if (self.f28) in (d_992_v10_) else (self).f48), globalState), self.f35, True])).cardinality
        nw164_[int(20)] = (self).f48
        nw164_[int(21)] = (d_993_v11_).cardinality
        nw164_[int(22)] = (self).f34
        nw164_[int(23)] = (self).f48
        nw164_[int(24)] = (self).f48
        nw164_[int(25)] = (self).f34
        nw164_[int(26)] = len(d_994_v12_)
        nw164_[int(27)] = (self).f34
        nw164_[int(28)] = len(_dafny.SeqWithoutIsStrInference([d_980_v0_ for d_996_i2_ in range(default__.abs(656))]))
        d_995_v13_ = nw164_
        d_997_v14_: _dafny.Set
        d_997_v14_ = _dafny.Set({d_995_v13_, d_995_v13_, d_995_v13_, d_995_v13_, d_995_v13_})
        d_997_v14_ = d_997_v14_
        d_998_v15_: D8
        d_998_v15_ = D8_DC22(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djojhst"))), default__.safeModuloInt(948, (self).f48))
        source18_ = d_998_v15_
        if source18_.is_DC22:
            d_999___mcc_h0_ = source18_.cf41
            d_1000___mcc_h1_ = source18_.cf42
            d_1001_cf42_ = d_1000___mcc_h1_
            d_1002_cf41_ = d_999___mcc_h0_
            rhs137_ = self.f31
            rhs138_ = (self).fm3(self.f28, self.f29, globalState)
            lhs100_ = globalState
            lhs100_.f1 = rhs137_
            r2 = rhs138_
            (globalState).f13 = d_1001_cf42_
            if not(self.f28):
                d_1003_v16_: _dafny.Set
                d_1003_v16_ = _dafny.Set({d_993_v11_, d_993_v11_})
                d_1004_v17_: _dafny.Map
                d_1004_v17_ = _dafny.Map({d_993_v11_: (self).f48})
                d_1005_v19_: D8
                d_1005_v19_ = D8_DC24(True, len(d_986_v6_), (self).f32, self.f35, d_985_v5_)
                d_1006_v20_: D8
                d_1006_v20_ = D8_DC25(d_1005_v19_)
                d_1007_v21_: _dafny.Array
                nw165_ = _dafny.Array(None, 8)
                nw165_[int(0)] = d_1003_v16_
                def iife52_():
                    coll26_ = _dafny.Set()
                    compr_26_: _dafny.MultiSet
                    for compr_26_ in (d_1004_v17_).keys.Elements:
                        d_1008_v18_: _dafny.MultiSet = compr_26_
                        if (d_1008_v18_) in (d_1004_v17_):
                            coll26_ = coll26_.union(_dafny.Set([d_1008_v18_]))
                    return _dafny.Set(coll26_)
                nw165_[int(1)] = iife52_()
                
                nw165_[int(2)] = (d_1003_v16_).intersection(default__.fm40(d_1002_cf41_, self.f29, d_1006_v20_, globalState))
                nw165_[int(3)] = d_1003_v16_
                nw165_[int(4)] = d_1003_v16_
                nw165_[int(5)] = d_1003_v16_
                nw165_[int(6)] = (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f31, self.f35, self.f35]))})).intersection(d_1003_v16_)
                nw165_[int(7)] = _dafny.Set({d_993_v11_, _dafny.MultiSet([False, self.f28]), d_993_v11_, d_993_v11_})
                d_1007_v21_ = nw165_
                index125_ = default__.safeIndex(263, (d_1007_v21_).length(0))
                (d_1007_v21_)[index125_] = d_1003_v16_
                index126_ = default__.safeIndex(263, (d_1007_v21_).length(0))
                (d_1007_v21_)[index126_] = d_1003_v16_
                (globalState).f13 = default__.fm39(self.f35, d_982_v2_, (not(self.f35)) or (self.f29), globalState)
                d_1009_v22_: _dafny.Array
                nw166_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1009_v22_ = nw166_
                d_1010_v23_: _dafny.Array
                def lambda36_(d_1011_i3_):
                    return self.f31

                init20_ = lambda36_
                nw167_ = _dafny.Array(None, 27)
                for i0_20_ in range(nw167_.length(0)):
                    nw167_[i0_20_] = init20_(i0_20_)
                d_1010_v23_ = nw167_
                index127_ = default__.safeIndex(580, (d_1009_v22_).length(0))
                (d_1009_v22_)[index127_] = d_1010_v23_
                index128_ = default__.safeIndex(529, (d_995_v13_).length(0))
                (d_995_v13_)[index128_] = d_1002_cf41_
                d_1012_v24_: _dafny.Array
                nw168_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_1012_v24_ = nw168_
                index129_ = default__.safeIndex(580, (d_1009_v22_).length(0))
                index130_ = default__.safeIndex(529, (d_995_v13_).length(0))
                rhs139_ = ((_dafny.Map({self.f29: (self).f47})).set(self.f35, self.f35)).set(self.f35, (self).f47)
                rhs140_ = d_1010_v23_
                rhs141_ = d_1002_cf41_
                rhs142_ = d_1012_v24_
                rhs143_ = ((self).f32) + ((self).f32)
                lhs101_ = d_1009_v22_
                lhs102_ = default__.safeIndex(580, (d_1009_v22_).length(0))
                lhs103_ = d_995_v13_
                lhs104_ = default__.safeIndex(529, (d_995_v13_).length(0))
                d_986_v6_ = rhs139_
                lhs101_[lhs102_] = rhs140_
                lhs103_[lhs104_] = rhs141_
                d_1012_v24_ = rhs142_
                d_988_v7_ = rhs143_
                d_1013_v25_: _dafny.Array
                nw169_ = _dafny.Array(int(0), 26)
                d_1013_v25_ = nw169_
                d_1014_v26_: _dafny.Map
                d_1014_v26_ = _dafny.Map({183: self.f35})
                d_1015_v27_: _dafny.Seq
                d_1015_v27_ = _dafny.SeqWithoutIsStrInference([d_980_v0_])
                d_1016_v28_: C1
                nw170_ = C1()
                nw170_.ctor__(self.f29, self.f31, self.f35)
                d_1016_v28_ = nw170_
                d_1017_v29_: _dafny.Set
                d_1017_v29_ = _dafny.Set({(self).f48, 696, d_1002_cf41_, len(_dafny.Map({d_1016_v28_: default__.fm41(self.f31, d_980_v0_, self.f28, (self).f48, globalState)}))})
                d_1018_v30_: _dafny.Array
                def lambda37_(d_1019_v0_):
                    def lambda38_(d_1020_i4_):
                        return d_1019_v0_

                    return lambda38_

                init21_ = lambda37_(d_980_v0_)
                nw171_ = _dafny.Array(None, 2)
                for i0_21_ in range(nw171_.length(0)):
                    nw171_[i0_21_] = init21_(i0_21_)
                d_1018_v30_ = nw171_
                d_1021_v31_: C0
                nw172_ = C0()
                nw172_.ctor__(d_1018_v30_)
                d_1021_v31_ = nw172_
                d_1022_v32_: _dafny.Map
                d_1022_v32_ = _dafny.Map({D9_DC27(d_1021_v31_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxwqwtkoe"))): d_1010_v23_})
                nw173_ = _dafny.Array(None, 22)
                nw173_[int(0)] = d_1002_cf41_
                nw173_[int(1)] = d_1001_cf42_
                nw173_[int(2)] = ((self).f30) + (d_1001_cf42_)
                nw173_[int(3)] = default__.fm39(((d_1014_v26_)[len(d_1014_v26_)] if (len(d_1014_v26_)) in (d_1014_v26_) else self.f28), d_982_v2_, self.f31, globalState)
                nw173_[int(4)] = default__.safeModuloInt(d_1001_cf42_, len(d_1015_v27_))
                nw173_[int(5)] = len((d_985_v5_) + (_dafny.SeqWithoutIsStrInference([(d_995_v13_)[default__.safeIndex(529, (d_995_v13_).length(0))]])))
                nw173_[int(6)] = (self).f30
                nw173_[int(7)] = (default__.fm39(False, d_982_v2_, self.f29, globalState)) * ((self).f48)
                nw173_[int(8)] = (self).f30
                nw173_[int(9)] = -634
                nw173_[int(10)] = d_1001_cf42_
                nw173_[int(11)] = d_1002_cf41_
                nw173_[int(12)] = (self).f34
                nw173_[int(13)] = default__.safeDivisionInt((self).f30, len((self).fm4(default__.fm1(self.f35, (d_995_v13_)[default__.safeIndex(529, (d_995_v13_).length(0))], globalState), d_1017_v29_, d_1001_cf42_, self.f28, globalState)))
                nw173_[int(14)] = default__.fm39(self.f31, d_982_v2_, not((d_1016_v28_).f43), globalState)
                nw173_[int(15)] = (self).f30
                nw173_[int(16)] = len((d_1022_v32_) | (d_1022_v32_))
                nw173_[int(17)] = (self).f48
                nw173_[int(18)] = (self).f30
                nw173_[int(19)] = (self).f34
                nw173_[int(20)] = (d_995_v13_)[default__.safeIndex(529, (d_995_v13_).length(0))]
                nw173_[int(21)] = (self).f34
                d_1013_v25_ = nw173_
                d_1001_cf42_ = default__.fm39(self.f29, _dafny.MultiSet(default__.fm42(d_1015_v27_, (self).f48, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkp")), globalState)), (d_1016_v28_).f43, globalState)
            elif True:
                (self).f35 = ((self).f30) == (default__.fm39((self).fm6(globalState), d_982_v2_, (self).f47, globalState))
                d_1001_cf42_ = 441
                r1 = ((0) - ((((d_993_v11_)[not(self.f29)] if (not(self.f29)) in (d_993_v11_) else len(d_985_v5_))) + (545)) if not(self.f28) else (self).f30)
                d_1023_v33_: _dafny.Array
                nw174_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_1023_v33_ = nw174_
                d_1024_v34_: _dafny.Set
                d_1024_v34_ = _dafny.Set({(self).f48})
                d_1025_v35_: _dafny.Map
                d_1025_v35_ = _dafny.Map({default__.fm41(self.f35, d_980_v0_, self.f29, (self).f34, globalState): (d_982_v2_).set((0) - ((self).f30), default__.abs(len(d_1024_v34_)))})
                index131_ = default__.safeIndex(825, (d_1023_v33_).length(0))
                (d_1023_v33_)[index131_] = ((d_1025_v35_)[default__.fm41(False, d_980_v0_, self.f28, default__.fm39(False, d_982_v2_, self.f35, globalState), globalState)] if (default__.fm41(False, d_980_v0_, self.f28, default__.fm39(False, d_982_v2_, self.f35, globalState), globalState)) in (d_1025_v35_) else d_982_v2_)
                index132_ = default__.safeIndex(825, (d_1023_v33_).length(0))
                (d_1023_v33_)[index132_] = default__.fm43(globalState)
                d_1026_v36_: _dafny.Array
                nw175_ = _dafny.Array(None, 1)
                nw175_[int(0)] = self.f35
                d_1026_v36_ = nw175_
                index133_ = default__.safeIndex(643, (d_1026_v36_).length(0))
                (d_1026_v36_)[index133_] = (self).f47
                index134_ = default__.safeIndex(643, (d_1026_v36_).length(0))
                (d_1026_v36_)[index134_] = ((self).f30) == ((self).f48)
            d_1027_v37_: _dafny.Map
            d_1027_v37_ = _dafny.Map({(self).f34: (self).f30})
            d_1027_v37_ = (d_1027_v37_).set(d_1002_cf41_, (self).f34)
        elif source18_.is_DC23:
            d_1028___mcc_h2_ = source18_.cf43
            d_1029___mcc_h3_ = source18_.cf44
            d_1030___mcc_h4_ = source18_.cf45
            d_1031_cf45_ = d_1030___mcc_h4_
            d_1032_cf44_ = d_1029___mcc_h3_
            d_1033_cf43_ = d_1028___mcc_h2_
            d_980_v0_ = d_980_v0_
            d_1034_v38_: _dafny.Seq
            d_1034_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nujp"))
            (globalState).f1 = ((d_1034_v38_).set(default__.safeIndex((self).f48, len(d_1034_v38_)), d_980_v0_)) < (d_1034_v38_)
            (self).f35 = False
            (globalState).f13 = (self).f48
        elif source18_.is_DC24:
            d_1035___mcc_h5_ = source18_.cf46
            d_1036___mcc_h6_ = source18_.cf47
            d_1037___mcc_h7_ = source18_.cf48
            d_1038___mcc_h8_ = source18_.cf49
            d_1039___mcc_h9_ = source18_.cf50
            d_1040_cf50_ = d_1039___mcc_h9_
            d_1041_cf49_ = d_1038___mcc_h8_
            d_1042_cf48_ = d_1037___mcc_h7_
            d_1043_cf47_ = d_1036___mcc_h6_
            d_1044_cf46_ = d_1035___mcc_h5_
            d_982_v2_ = (d_982_v2_) - (d_982_v2_)
            d_980_v0_ = _dafny.CodePoint('i')
            d_1045_v39_: _dafny.Seq
            d_1045_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvj"))
            d_1046_v40_: _dafny.Array
            nw176_ = _dafny.Array(None, 5)
            nw176_[int(0)] = d_1045_v39_
            nw176_[int(1)] = d_1045_v39_
            nw176_[int(2)] = d_1045_v39_
            nw176_[int(3)] = (d_1045_v39_) + (d_1045_v39_)
            nw176_[int(4)] = (d_1045_v39_).set(default__.safeIndex((0) - ((self).f34), len(d_1045_v39_)), d_980_v0_)
            d_1046_v40_ = nw176_
            index135_ = default__.safeIndex(528, (d_1046_v40_).length(0))
            (d_1046_v40_)[index135_] = d_1045_v39_
            d_1047_v41_: _dafny.Seq
            d_1047_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_980_v0_ for d_1048_i5_ in range(default__.abs(849))])])
            index136_ = default__.safeIndex(528, (d_1046_v40_).length(0))
            (d_1046_v40_)[index136_] = (d_1047_v41_)[default__.safeIndex((self).f30, len(d_1047_v41_))]
            d_1049_v42_: _dafny.Seq
            d_1049_v42_ = _dafny.SeqWithoutIsStrInference([d_985_v5_, d_1040_cf50_, d_1040_cf50_])
            d_1050_v43_: _dafny.Map
            d_1050_v43_ = _dafny.Map({(d_1049_v42_)[default__.safeIndex((0) - (d_1043_cf47_), len(d_1049_v42_))]: d_1041_cf49_})
            d_1051_v44_: D6
            d_1051_v44_ = D6_DC15(d_985_v5_)
            d_1050_v43_ = ((d_1050_v43_) | (d_1050_v43_)).set((d_1051_v44_).cf32, d_1041_cf49_)
        elif source18_.is_DC21:
            d_1052___mcc_h10_ = source18_.cf40
            d_1053_cf40_ = d_1052___mcc_h10_
            d_1054_v45_: C2
            nw177_ = C2()
            nw177_.ctor__((self).f48, (self).f30, self.f31, self.f28, self.f35)
            d_1054_v45_ = nw177_
            d_1055_v46_: _dafny.Array
            nw178_ = _dafny.Array(None, 8)
            nw178_[int(0)] = (self).f47
            nw178_[int(1)] = self.f35
            nw178_[int(2)] = self.f29
            nw178_[int(3)] = True
            nw178_[int(4)] = self.f35
            nw178_[int(5)] = self.f35
            nw178_[int(6)] = self.f28
            nw178_[int(7)] = self.f35
            d_1055_v46_ = nw178_
            d_1056_v47_: D3
            d_1056_v47_ = D3_DC9((d_1054_v45_).fm30(self.f35, globalState), (self).fm6(globalState), d_1055_v46_, (d_988_v7_)[default__.safeIndex((self).f48, len(d_988_v7_))])
            d_1057_v48_: D6
            d_1057_v48_ = D6_DC15(d_985_v5_)
            d_1058_v49_: _dafny.Seq
            d_1058_v49_ = _dafny.SeqWithoutIsStrInference([d_985_v5_])
            d_1059_v50_: D8
            d_1059_v50_ = D8_DC24(self.f35, (self).f34, (self).f32, self.f28, (d_1058_v49_)[default__.safeIndex((self).f34, len(d_1058_v49_))])
            pat_let_tv10_ = d_1059_v50_
            d_1060_v51_: bool
            d_1061_v52_: str
            out54_: bool
            out55_: str
            def iife53_(_pat_let13_0):
                def iife54_(d_1062_dt__update__tmp_h0_):
                    def iife55_(_pat_let14_0):
                        def iife56_(d_1063_dt__update_hcf32_h0_):
                            return D6_DC15(d_1063_dt__update_hcf32_h0_)
                        return iife56_(_pat_let14_0)
                    return iife55_((pat_let_tv10_).cf50)
                return iife54_(_pat_let13_0)
            out54_, out55_ = (self).m14(d_980_v0_, (d_1056_v47_).cf20, iife53_(d_1057_v48_), globalState)
            d_1060_v51_ = out54_
            d_1061_v52_ = out55_
            (self).f29 = self.f29
            (globalState).f13 = (self).f34
        elif True:
            d_1064___mcc_h11_ = source18_.cf51
            d_1065_cf51_ = d_1064___mcc_h11_
            d_1066_v53_: _dafny.Map
            d_1066_v53_ = _dafny.Map({(self).f30: (self).f34})
            d_1067_v54_: _dafny.Seq
            d_1067_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
            d_1068_v55_: _dafny.Seq
            d_1068_v55_ = _dafny.SeqWithoutIsStrInference([d_1066_v53_])
            d_1066_v53_ = (_dafny.Map({len(d_1067_v54_): len(d_1067_v54_)})) | ((d_1068_v55_)[default__.safeIndex((self).f48, len(d_1068_v55_))])
            (globalState).f13 = 342
            source19_ = d_998_v15_
            if source19_.is_DC22:
                d_1069___mcc_h12_ = source19_.cf41
                d_1070___mcc_h13_ = source19_.cf42
                d_1071_cf42_ = d_1070___mcc_h13_
                d_1072_cf41_ = d_1069___mcc_h12_
                d_1073_v56_: _dafny.Map
                d_1073_v56_ = _dafny.Map({((self).f30) - (d_1071_cf42_): self.f28})
                d_1073_v56_ = (d_1073_v56_).set((991) * ((self).f34), ((self).f34) == (d_1071_cf42_))
                d_1074_v58_: _dafny.Map
                d_1074_v58_ = _dafny.Map({self.f29: d_1073_v56_})
                d_1075_v59_: _dafny.Array
                nw179_ = _dafny.Array(None, 10)
                def iife57_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in (d_985_v5_).Elements:
                        d_1076_v57_: int = compr_27_
                        if (d_1076_v57_) in (d_985_v5_):
                            coll27_[default__.safeDivisionInt(d_1076_v57_, len(_dafny.Map({self.f35: len(_dafny.Set({d_985_v5_}))})))] = self.f31
                    return _dafny.Map(coll27_)
                nw179_[int(0)] = iife57_()
                
                nw179_[int(1)] = d_1073_v56_
                nw179_[int(2)] = d_1073_v56_
                nw179_[int(3)] = d_1073_v56_
                nw179_[int(4)] = ((d_1074_v58_)[(self).f47] if ((self).f47) in (d_1074_v58_) else d_1073_v56_)
                nw179_[int(5)] = _dafny.Map({(self).f30: self.f35})
                nw179_[int(6)] = d_1073_v56_
                nw179_[int(7)] = d_1073_v56_
                nw179_[int(8)] = d_1073_v56_
                nw179_[int(9)] = d_1073_v56_
                d_1075_v59_ = nw179_
                d_1077_v60_: _dafny.Array
                nw180_ = _dafny.Array(None, 9)
                nw180_[int(0)] = d_1075_v59_
                nw180_[int(1)] = d_1075_v59_
                nw180_[int(2)] = d_1075_v59_
                nw180_[int(3)] = d_1075_v59_
                nw180_[int(4)] = d_1075_v59_
                nw180_[int(5)] = d_1075_v59_
                nw180_[int(6)] = d_1075_v59_
                nw180_[int(7)] = (d_1075_v59_ if self.f35 else d_1075_v59_)
                nw180_[int(8)] = d_1075_v59_
                d_1077_v60_ = nw180_
                index137_ = default__.safeIndex(381, (d_1077_v60_).length(0))
                (d_1077_v60_)[index137_] = d_1075_v59_
                d_1078_v61_: _dafny.Set
                d_1078_v61_ = _dafny.Set({len(_dafny.Map({self.f29: -133})), len(d_986_v6_), (self).f48, d_1072_cf41_})
                d_1079_v63_: _dafny.Array
                nw181_ = _dafny.Array(None, 17)
                nw181_[int(0)] = self.f28
                def iife58_():
                    coll28_ = _dafny.Set()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(57, 761):
                        d_1080_v62_: int = compr_28_
                        if ((57) <= (d_1080_v62_)) and ((d_1080_v62_) < (761)):
                            coll28_ = coll28_.union(_dafny.Set([(d_1080_v62_) * ((self).f30)]))
                    return _dafny.Set(coll28_)
                nw181_[int(1)] = (iife58_()
                ).ispropersubset(d_1078_v61_)
                nw181_[int(2)] = (not((self).f47)) or ((self).f47)
                nw181_[int(3)] = (default__.fm1(self.f35, d_1072_cf41_, globalState)) or (self.f28)
                nw181_[int(4)] = False
                nw181_[int(5)] = self.f28
                nw181_[int(6)] = self.f31
                nw181_[int(7)] = (self).f47
                nw181_[int(8)] = False
                nw181_[int(9)] = (self).f47
                nw181_[int(10)] = self.f35
                nw181_[int(11)] = (self).fm6(globalState)
                nw181_[int(12)] = self.f28
                nw181_[int(13)] = (self.f29) not in ((_dafny.MultiSet(d_988_v7_)).set(not(self.f31), default__.abs(d_1071_cf42_)))
                nw181_[int(14)] = self.f29
                nw181_[int(15)] = self.f35
                nw181_[int(16)] = self.f31
                d_1079_v63_ = nw181_
                index138_ = default__.safeIndex(135, (d_1079_v63_).length(0))
                (d_1079_v63_)[index138_] = self.f29
                d_1081_v64_: _dafny.Seq
                d_1081_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_980_v0_ for d_1082_i6_ in range(default__.abs(889))])])
                index139_ = default__.safeIndex(381, (d_1077_v60_).length(0))
                index140_ = default__.safeIndex(135, (d_1079_v63_).length(0))
                rhs144_ = 788
                rhs145_ = d_1075_v59_
                rhs146_ = (len(((d_1081_v64_)[default__.safeIndex(d_1072_cf41_, len(d_1081_v64_))]) + (d_1067_v54_))) <= ((d_993_v11_).cardinality)
                lhs105_ = globalState
                lhs106_ = d_1077_v60_
                lhs107_ = default__.safeIndex(381, (d_1077_v60_).length(0))
                lhs108_ = d_1079_v63_
                lhs109_ = default__.safeIndex(135, (d_1079_v63_).length(0))
                lhs105_.f13 = rhs144_
                lhs106_[lhs107_] = rhs145_
                lhs108_[lhs109_] = rhs146_
                (globalState).f13 = (self).f34
                d_1083_v65_: _dafny.Array
                def lambda39_(d_1084_v11_):
                    def lambda40_(d_1085_i7_):
                        return d_1084_v11_

                    return lambda40_

                init22_ = lambda39_(d_993_v11_)
                nw182_ = _dafny.Array(None, 17)
                for i0_22_ in range(nw182_.length(0)):
                    nw182_[i0_22_] = init22_(i0_22_)
                d_1083_v65_ = nw182_
                d_1086_v66_: _dafny.Map
                d_1086_v66_ = _dafny.Map({d_1083_v65_: d_1072_cf41_})
                d_1086_v66_ = (d_1086_v66_).set(d_1083_v65_, (self).f34)
            elif source19_.is_DC23:
                d_1087___mcc_h14_ = source19_.cf43
                d_1088___mcc_h15_ = source19_.cf44
                d_1089___mcc_h16_ = source19_.cf45
                d_1090_cf45_ = d_1089___mcc_h16_
                d_1091_cf44_ = d_1088___mcc_h15_
                d_1092_cf43_ = d_1087___mcc_h14_
                (globalState).f13 = (self).f34
                d_1093_v67_: _dafny.Seq
                d_1093_v67_ = _dafny.SeqWithoutIsStrInference([d_1067_v54_, (d_1067_v54_) + (d_1067_v54_)])
                d_1093_v67_ = (d_1093_v67_) + (d_1093_v67_)
                d_1094_v68_: _dafny.Set
                d_1094_v68_ = _dafny.Set({d_1090_cf45_, (self).f47})
                d_1095_v69_: C1
                nw183_ = C1()
                nw183_.ctor__((len(d_1094_v68_)) == ((self).f30), self.f28, self.f31)
                d_1095_v69_ = nw183_
                d_995_v13_ = d_995_v13_
            elif source19_.is_DC24:
                d_1096___mcc_h17_ = source19_.cf46
                d_1097___mcc_h18_ = source19_.cf47
                d_1098___mcc_h19_ = source19_.cf48
                d_1099___mcc_h20_ = source19_.cf49
                d_1100___mcc_h21_ = source19_.cf50
                d_1101_cf50_ = d_1100___mcc_h21_
                d_1102_cf49_ = d_1099___mcc_h20_
                d_1103_cf48_ = d_1098___mcc_h19_
                d_1104_cf47_ = d_1097___mcc_h18_
                d_1105_cf46_ = d_1096___mcc_h17_
                (globalState).f13 = (default__.safeModuloInt((self).f48, (self).f34)) - ((self).f34)
                d_1106_v70_: D3
                d_1106_v70_ = D3_DC8(_dafny.MultiSet([d_980_v0_]), (self).f47, 735)
                pat_let_tv11_ = d_1102_cf49_
                def iife59_(_pat_let15_0):
                    def iife60_(d_1107_dt__update__tmp_h1_):
                        def iife61_(_pat_let16_0):
                            def iife62_(d_1108_dt__update_hcf16_h0_):
                                return D3_DC8((d_1107_dt__update__tmp_h1_).cf15, d_1108_dt__update_hcf16_h0_, (d_1107_dt__update__tmp_h1_).cf17)
                            return iife62_(_pat_let16_0)
                        return iife61_(pat_let_tv11_)
                    return iife60_(_pat_let15_0)
                d_1106_v70_ = iife59_(d_1106_v70_)
                d_1109_v71_: _dafny.Array
                def lambda41_(d_1110_i8_):
                    return _dafny.Map({(self).f34: self.f31})

                init23_ = lambda41_
                nw184_ = _dafny.Array(None, 1)
                for i0_23_ in range(nw184_.length(0)):
                    nw184_[i0_23_] = init23_(i0_23_)
                d_1109_v71_ = nw184_
                d_1111_v72_: _dafny.Map
                d_1111_v72_ = _dafny.Map({(self).f34: d_1102_cf49_})
                index141_ = default__.safeIndex(224, (d_1109_v71_).length(0))
                (d_1109_v71_)[index141_] = (d_1111_v72_) | (d_1111_v72_)
                index142_ = default__.safeIndex(224, (d_1109_v71_).length(0))
                (d_1109_v71_)[index142_] = (d_1111_v72_ if (True if self.f29 else ((self).f32)[default__.safeIndex((self).f48, len((self).f32))]) else _dafny.Map({(self).f30: (self).f47}))
                d_1112_v74_: _dafny.Set
                def iife63_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in ((d_1109_v71_)[default__.safeIndex(224, (d_1109_v71_).length(0))]).keys.Elements:
                        d_1113_v73_: int = compr_29_
                        if (d_1113_v73_) in ((d_1109_v71_)[default__.safeIndex(224, (d_1109_v71_).length(0))]):
                            coll29_ = coll29_.union(_dafny.Set([(d_1113_v73_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmcvtul"))))]))
                    return _dafny.Set(coll29_)
                d_1112_v74_ = _dafny.Set({(0) - (len(iife63_()
                )), (self).f30})
                d_1114_v76_: _dafny.Seq
                d_1114_v76_ = _dafny.SeqWithoutIsStrInference([d_1112_v74_])
                d_1115_v77_: _dafny.Array
                nw185_ = _dafny.Array(None, 13)
                nw185_[int(0)] = d_1112_v74_
                nw185_[int(1)] = (d_1112_v74_) - (_dafny.Set({(self).f34, (self).f34}))
                nw185_[int(2)] = _dafny.Set({(0) - ((self).f30), (self).f30})
                def iife64_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(718, -585):
                        d_1116_v75_: int = compr_30_
                        if ((718) <= (d_1116_v75_)) and ((d_1116_v75_) < (-585)):
                            coll30_ = coll30_.union(_dafny.Set([default__.safeModuloInt(d_1116_v75_, 470)]))
                    return _dafny.Set(coll30_)
                nw185_[int(3)] = iife64_()
                
                nw185_[int(4)] = d_1112_v74_
                nw185_[int(5)] = (d_1112_v74_) - (d_1112_v74_)
                nw185_[int(6)] = (d_1112_v74_) - (d_1112_v74_)
                nw185_[int(7)] = (d_1114_v76_)[default__.safeIndex(len(d_1101_cf50_), len(d_1114_v76_))]
                nw185_[int(8)] = d_1112_v74_
                nw185_[int(9)] = d_1112_v74_
                nw185_[int(10)] = d_1112_v74_
                nw185_[int(11)] = d_1112_v74_
                nw185_[int(12)] = d_1112_v74_
                d_1115_v77_ = nw185_
                index143_ = default__.safeIndex(248, (d_1115_v77_).length(0))
                (d_1115_v77_)[index143_] = _dafny.Set({d_1104_cf47_, (self).f30})
                index144_ = default__.safeIndex(248, (d_1115_v77_).length(0))
                def iife65_():
                    coll31_ = _dafny.Set()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(-346, 278):
                        d_1117_v78_: int = compr_31_
                        if ((-346) <= (d_1117_v78_)) and ((d_1117_v78_) < (278)):
                            coll31_ = coll31_.union(_dafny.Set([default__.safeDivisionInt(d_1117_v78_, d_1104_cf47_)]))
                    return _dafny.Set(coll31_)
                (d_1115_v77_)[index144_] = (self).fm4(d_1102_cf49_, iife65_()
                , (self).f34, d_1105_cf46_, globalState)
            elif source19_.is_DC21:
                d_1118___mcc_h22_ = source19_.cf40
                d_1119_cf40_ = d_1118___mcc_h22_
                d_1120_v79_: _dafny.Map
                d_1120_v79_ = _dafny.Map({self.f29: (d_982_v2_) | (_dafny.MultiSet([(self).f30]))})
                rhs147_ = ((d_1120_v79_)[(self.f28) == (self.f29)] if ((self.f28) == (self.f29)) in (d_1120_v79_) else (d_982_v2_).set((self).f48, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gonoudi"))))))
                rhs148_ = (self).f47
                rhs149_ = ((self).f48) - ((self).f30)
                lhs110_ = globalState
                lhs111_ = globalState
                d_982_v2_ = rhs147_
                lhs110_.f12 = rhs148_
                lhs111_.f13 = rhs149_
                d_1121_v80_: _dafny.Array
                nw186_ = _dafny.Array(None, 6)
                nw186_[int(0)] = False
                nw186_[int(1)] = ((self).f47) == (self.f29)
                nw186_[int(2)] = ((self).f30) > ((self).f48)
                nw186_[int(3)] = not(self.f28)
                nw186_[int(4)] = True
                nw186_[int(5)] = self.f28
                d_1121_v80_ = nw186_
                index145_ = default__.safeIndex(690, (d_1121_v80_).length(0))
                (d_1121_v80_)[index145_] = (self.f31 if self.f31 else self.f29)
                index146_ = default__.safeIndex(690, (d_1121_v80_).length(0))
                (d_1121_v80_)[index146_] = False
                d_1122_v81_: _dafny.MultiSet
                d_1122_v81_ = _dafny.MultiSet([(d_989_v8_)[default__.safeIndex(423, (d_989_v8_).length(0))]])
                index147_ = default__.safeIndex(690, (d_1121_v80_).length(0))
                def iife66_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in (d_985_v5_).Elements:
                        d_1123_v82_: int = compr_32_
                        if (d_1123_v82_) in (d_985_v5_):
                            coll32_[default__.safeModuloInt(d_1123_v82_, (self).f30)] = d_1067_v54_
                    return _dafny.Map(coll32_)
                rhs150_ = d_1121_v80_
                rhs151_ = not(((len(iife66_()
                )) - (len(d_1067_v54_))) >= (len(d_988_v7_)))
                rhs152_ = _dafny.MultiSet([D4_DC11((self).f30, (d_1121_v80_)[default__.safeIndex(690, (d_1121_v80_).length(0))], (self).f30)])
                lhs112_ = d_1121_v80_
                lhs113_ = default__.safeIndex(690, (d_1121_v80_).length(0))
                d_1121_v80_ = rhs150_
                lhs112_[lhs113_] = rhs151_
                d_1122_v81_ = rhs152_
                (self).f29 = False
            elif True:
                d_1124___mcc_h23_ = source19_.cf51
                d_1125_cf51_ = d_1124___mcc_h23_
                (self).f35 = ((d_988_v7_) + ((self).f32)) < (((self).fm5((self).f48, (self).f48, (self).f47, (self).f48, globalState)) + ((self).f32))
                d_1126_v83_: C1
                nw187_ = C1()
                nw187_.ctor__((d_1067_v54_) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1127_i9_ in range(default__.abs(314))])), ((self).f34) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsm")))), self.f31)
                d_1126_v83_ = nw187_
                (globalState).f26 = (d_1126_v83_).f43
                d_1128_v84_: _dafny.Array
                nw188_ = _dafny.Array(False, 29)
                d_1128_v84_ = nw188_
                (globalState).f21 = d_1128_v84_
            rhs153_ = default__.fm39(self.f28, d_982_v2_, not(self.f31), globalState)
            rhs154_ = ((self).f30) < ((self).f48)
            rhs155_ = (358) == (((d_993_v11_)[self.f28] if (self.f28) in (d_993_v11_) else len((d_985_v5_).set(default__.safeIndex((self).f34, len(d_985_v5_)), (self).f48))))
            lhs114_ = globalState
            lhs115_ = self
            lhs116_ = globalState
            lhs114_.f13 = rhs153_
            lhs115_.f29 = rhs154_
            lhs116_.f12 = rhs155_
        hi5_ = (0) - (-182)
        for d_1129_i10_ in range((len(_dafny.Map({(self).f34: d_982_v2_}))) - ((0) - ((self).f48)), hi5_):
            (globalState).f12 = default__.fm1(self.f35, (d_982_v2_).cardinality, globalState)
            (globalState).f12 = ((169) != (d_1129_i10_)) and (self.f35)
            (globalState).f13 = ((self).f34) * ((self).f48)
            d_1130_v85_: _dafny.Map
            d_1130_v85_ = _dafny.Map({self.f31: d_993_v11_})
            d_1130_v85_ = (d_1130_v85_).set(self.f28, _dafny.MultiSet([self.f31]))
        d_1131_v86_: _dafny.Seq
        d_1131_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjtojxk"))
        r0 = (len(d_1131_v86_)) > ((self).f34)
        r1 = (((self).f34) + ((self).f30) if not(self.f28) else (0) - ((self).f34))
        r2 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))) + (d_1131_v86_)).set(default__.safeIndex(213, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))) + (d_1131_v86_))), d_980_v0_)
        return r0, r1, r2

    def m14(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: str = _dafny.CodePoint('D')
        hi6_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxgmww")))
        for d_1132_i0_ in range((self).f30, hi6_):
            (globalState).f12 = self.f28
            (globalState).f13 = default__.safeDivisionInt(d_1132_i0_, (d_1132_i0_) - (d_1132_i0_))
            d_1133_v0_: _dafny.Map
            d_1133_v0_ = _dafny.Map({73: len(_dafny.SeqWithoutIsStrInference([p0 for d_1134_i1_ in range(default__.abs(137))]))})
            d_1135_v1_: T3
            nw189_ = C4()
            nw189_.ctor__((self).f30, self.f29, (self).f32, (self).f33, len(d_1133_v0_), self.f35, default__.fm1(False, (self).f48, globalState), (self).f47)
            d_1135_v1_ = nw189_
            d_1136_v2_: _dafny.Map
            d_1136_v2_ = _dafny.Map({d_1132_i0_: d_1135_v1_})
            d_1137_v3_: D10
            d_1137_v3_ = D10_DC30(default__.fm41(self.f31, p0, self.f31, (self).f34, globalState), _dafny.MultiSet([d_1135_v1_, d_1135_v1_, ((d_1136_v2_)[764] if (764) in (d_1136_v2_) else d_1135_v1_), d_1135_v1_]))
            d_1138_v4_: _dafny.MultiSet
            d_1138_v4_ = _dafny.MultiSet([d_1135_v1_, d_1135_v1_, d_1135_v1_, d_1135_v1_, d_1135_v1_])
            pat_let_tv12_ = d_1138_v4_
            def iife67_(_pat_let17_0):
                def iife68_(d_1139_dt__update__tmp_h0_):
                    def iife69_(_pat_let18_0):
                        def iife70_(d_1140_dt__update_hcf59_h0_):
                            return D10_DC30((d_1139_dt__update__tmp_h0_).cf58, d_1140_dt__update_hcf59_h0_)
                        return iife70_(_pat_let18_0)
                    return iife69_(pat_let_tv12_)
                return iife68_(_pat_let17_0)
            source20_ = iife67_(d_1137_v3_)
            if source20_.is_DC29:
                d_1141___mcc_h0_ = source20_.cf56
                d_1142___mcc_h1_ = source20_.cf57
                d_1143_cf57_ = d_1142___mcc_h1_
                d_1144_cf56_ = d_1141___mcc_h0_
                (globalState).f13 = (d_1144_cf56_) - (((self).f34) * ((self).f34))
                d_1145_v5_: _dafny.Set
                d_1145_v5_ = _dafny.Set({p0, p0, p0})
                index148_ = default__.safeIndex(637, (p1).length(0))
                (p1)[index148_] = (self.f31) not in (_dafny.Map({d_1135_v1_.f31: d_1145_v5_}))
                index149_ = default__.safeIndex(637, (p1).length(0))
                (p1)[index149_] = d_1135_v1_.f29
                d_1146_v6_: _dafny.Seq
                d_1146_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xji"))
                d_1147_v7_: _dafny.Map
                d_1147_v7_ = _dafny.Map({d_1146_v6_: (self).f30})
                d_1147_v7_ = d_1147_v7_
                d_1148_v8_: _dafny.Array
                def lambda42_(d_1149_p0_):
                    def lambda43_(d_1150_i2_):
                        return _dafny.SeqWithoutIsStrInference([d_1149_p0_ for d_1151_i3_ in range(default__.abs(315))])

                    return lambda43_

                init24_ = lambda42_(p0)
                nw190_ = _dafny.Array(None, 29)
                for i0_24_ in range(nw190_.length(0)):
                    nw190_[i0_24_] = init24_(i0_24_)
                d_1148_v8_ = nw190_
                index150_ = default__.safeIndex(472, (d_1148_v8_).length(0))
                (d_1148_v8_)[index150_] = d_1146_v6_
                index151_ = default__.safeIndex(472, (d_1148_v8_).length(0))
                (d_1148_v8_)[index151_] = _dafny.SeqWithoutIsStrInference([p0 for d_1152_i4_ in range(default__.abs(822))])
            elif source20_.is_DC30:
                d_1153___mcc_h2_ = source20_.cf58
                d_1154___mcc_h3_ = source20_.cf59
                d_1155_cf59_ = d_1154___mcc_h3_
                d_1156_cf58_ = d_1153___mcc_h2_
                (globalState).f13 = (self).f48
                d_1157_v9_: _dafny.Array
                nw191_ = _dafny.Array(None, 4)
                nw191_[int(0)] = (d_1135_v1_).f34
                nw191_[int(1)] = (d_1135_v1_).f34
                nw191_[int(2)] = (self).f34
                nw191_[int(3)] = (d_1135_v1_).f34
                d_1157_v9_ = nw191_
                d_1158_v10_: _dafny.Map
                d_1158_v10_ = _dafny.Map({d_1157_v9_: d_1135_v1_.f28})
                d_1159_v11_: D7
                d_1159_v11_ = D7_DC19(d_1135_v1_.f31, (0) - ((d_1135_v1_).f34), len(d_1158_v10_))
                (globalState).f13 = (d_1159_v11_).cf38
                d_1160_v12_: _dafny.Seq
                d_1160_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dplmqrvf"))
                d_1161_v13_: _dafny.Map
                d_1161_v13_ = _dafny.Map({(self).f48: d_1160_v12_})
                d_1162_v14_: _dafny.Set
                d_1162_v14_ = _dafny.Set({self.f31})
                d_1163_v16_: _dafny.Map
                d_1163_v16_ = _dafny.Map({self.f28: self.f28})
                d_1164_v17_: _dafny.Seq
                d_1164_v17_ = _dafny.SeqWithoutIsStrInference([len((d_1135_v1_).f32), len(d_1163_v16_)])
                d_1165_v18_: _dafny.Array
                nw192_ = _dafny.Array(None, 26)
                nw192_[int(0)] = (d_1135_v1_).f32
                nw192_[int(1)] = (d_1135_v1_).f32
                nw192_[int(2)] = (d_1135_v1_).f32
                nw192_[int(3)] = (d_1135_v1_).f32
                nw192_[int(4)] = (d_1135_v1_).f32
                nw192_[int(5)] = (d_1135_v1_).fm5(len(((d_1161_v13_)[(self).f48] if ((self).f48) in (d_1161_v13_) else d_1160_v12_)), (self).f34, (self).f47, (d_1135_v1_).f30, globalState)
                nw192_[int(6)] = (d_1135_v1_).f32
                nw192_[int(7)] = (d_1135_v1_).f32
                nw192_[int(8)] = (self).f32
                nw192_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1135_v1_.f28, d_1135_v1_.f29])
                nw192_[int(10)] = (d_1135_v1_).f32
                nw192_[int(11)] = (d_1135_v1_).f32
                nw192_[int(12)] = (d_1135_v1_).f32
                nw192_[int(13)] = (d_1135_v1_).f32
                nw192_[int(14)] = (d_1135_v1_).f32
                def iife71_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in (d_1164_v17_).Elements:
                        d_1166_v15_: int = compr_33_
                        if (d_1166_v15_) in (d_1164_v17_):
                            coll33_[default__.safeModuloInt(d_1166_v15_, len(_dafny.Map({self.f29: (self).f30})))] = d_1135_v1_.f28
                    return _dafny.Map(coll33_)
                nw192_[int(15)] = (d_1135_v1_).fm5(len(d_1162_v14_), (self).f30, d_1135_v1_.f29, len(iife71_()
                ), globalState)
                nw192_[int(16)] = _dafny.SeqWithoutIsStrInference([self.f29])
                nw192_[int(17)] = (self).f32
                nw192_[int(18)] = _dafny.SeqWithoutIsStrInference([d_1135_v1_.f35])
                nw192_[int(19)] = (d_1135_v1_).f32
                nw192_[int(20)] = (d_1135_v1_).f32
                nw192_[int(21)] = (self).f32
                nw192_[int(22)] = (d_1135_v1_).f32
                nw192_[int(23)] = (d_1135_v1_).f32
                nw192_[int(24)] = _dafny.SeqWithoutIsStrInference([(self).f47, d_1135_v1_.f35, self.f28])
                nw192_[int(25)] = (d_1135_v1_).f32
                d_1165_v18_ = nw192_
                d_1167_v19_: T4
                nw193_ = C3()
                nw193_.ctor__(893, d_1165_v18_, p0, (default__.fm46(len(d_1160_v12_), (self).f34, self.f31, globalState)) * ((d_1135_v1_).f34), not(not(((self).f32)[default__.safeIndex((d_1135_v1_).f30, len((self).f32))])), (_dafny.SeqWithoutIsStrInference([False])) + ((d_1135_v1_).f32), ((d_1135_v1_).f33) + ((self).f33), (d_1135_v1_).f30, not((d_1135_v1_.f31 if d_1135_v1_.f35 else self.f35)), (True if d_1135_v1_.f31 else False), True)
                d_1167_v19_ = nw193_
                d_1167_v19_ = d_1167_v19_
                d_1168_v20_: _dafny.Map
                d_1168_v20_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0 for d_1169_i5_ in range(default__.abs(632))]): d_1162_v14_})
                d_1168_v20_ = (d_1168_v20_).set((d_1160_v12_) + (((self).fm3((self).f47, d_1135_v1_.f28, globalState)).set(default__.safeIndex(((d_1133_v0_)[len(d_1133_v0_)] if (len(d_1133_v0_)) in (d_1133_v0_) else len((self).f32)), len((self).fm3((self).f47, d_1135_v1_.f28, globalState))), d_1167_v19_.f36)), d_1162_v14_)
            elif source20_.is_DC31:
                d_1170___mcc_h4_ = source20_.cf60
                d_1171___mcc_h5_ = source20_.cf61
                d_1172___mcc_h6_ = source20_.cf62
                d_1173_cf62_ = d_1172___mcc_h6_
                d_1174_cf61_ = d_1171___mcc_h5_
                d_1175_cf60_ = d_1170___mcc_h4_
                index152_ = default__.safeIndex(628, (p1).length(0))
                (p1)[index152_] = ((d_1135_v1_).f32)[default__.safeIndex(959, len((d_1135_v1_).f32))]
                d_1176_v21_: D7
                d_1176_v21_ = D7_DC19((self).f47, (d_1135_v1_).f34, (d_1135_v1_).f34)
                index153_ = default__.safeIndex(628, (p1).length(0))
                (p1)[index153_] = (d_1176_v21_).cf36
                d_1177_v22_: _dafny.Array
                nw194_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_1177_v22_ = nw194_
                d_1177_v22_ = d_1177_v22_
                d_1178_v23_: C1
                nw195_ = C1()
                nw195_.ctor__(False, (p1)[default__.safeIndex(628, (p1).length(0))], False)
                d_1178_v23_ = nw195_
                d_1179_v24_: D10
                d_1179_v24_ = D10_DC28(d_1178_v23_)
                d_1180_v25_: _dafny.MultiSet
                d_1180_v25_ = _dafny.MultiSet([d_1179_v24_])
                d_1181_v26_: _dafny.Set
                d_1181_v26_ = _dafny.Set({(d_1180_v25_) - (d_1180_v25_)})
                d_1181_v26_ = d_1181_v26_
                (globalState).f13 = d_1132_i0_
            elif source20_.is_DC28:
                d_1182___mcc_h7_ = source20_.cf55
                d_1183_cf55_ = d_1182___mcc_h7_
                r1 = p0
                d_1184_v27_: _dafny.Array
                nw196_ = _dafny.Array(int(0), 12)
                d_1184_v27_ = nw196_
                d_1185_v28_: _dafny.Map
                d_1185_v28_ = _dafny.Map({d_1184_v27_: (self.f28) or (self.f31)})
                d_1185_v28_ = (d_1185_v28_)
                nw197_ = _dafny.Array(int(0), 25)
                d_1184_v27_ = nw197_
                (globalState).f12 = not((d_1183_cf55_).f43)
            elif True:
                d_1186___mcc_h8_ = source20_.cf63
                d_1187_cf63_ = d_1186___mcc_h8_
                d_1188_v29_: _dafny.Map
                d_1188_v29_ = _dafny.Map({self.f29: self.f28})
                d_1189_v30_: _dafny.MultiSet
                d_1189_v30_ = _dafny.MultiSet([((d_1135_v1_).f34) >= ((self).f48), ((d_1188_v29_)[d_1135_v1_.f28] if (d_1135_v1_.f28) in (d_1188_v29_) else d_1135_v1_.f28), (self.f28 if d_1135_v1_.f35 else d_1135_v1_.f28), d_1135_v1_.f28])
                (globalState).f13 = (d_1189_v30_).cardinality
                (d_1135_v1_).f35 = d_1135_v1_.f35
                r1 = p0
                (globalState).f13 = default__.fm46((d_1135_v1_).f30, ((self).f48 if self.f31 else (d_1135_v1_).f30), d_1135_v1_.f28, globalState)
            d_1190_v31_: _dafny.MultiSet
            d_1190_v31_ = _dafny.MultiSet([False])
            if (((self).f32) == ((self).f32) if (d_1190_v31_).issubset(_dafny.MultiSet([self.f35, (self).f47])) else True):
                (globalState).f13 = (self).f30
                d_1191_v32_: _dafny.Seq
                d_1191_v32_ = _dafny.SeqWithoutIsStrInference([(d_1135_v1_).f34, d_1132_i0_, (d_1135_v1_).f34])
                d_1192_v33_: D8
                d_1192_v33_ = D8_DC24(d_1135_v1_.f35, (d_1135_v1_).f30, (self).f32, d_1135_v1_.f31, (d_1191_v32_).set(default__.safeIndex((self).f48, len(d_1191_v32_)), (self).f48))
                d_1193_v34_: D0
                d_1193_v34_ = D0_DC0(d_1135_v1_.f31)
                d_1194_v35_: C2
                nw198_ = C2()
                nw198_.ctor__(len(d_1133_v0_), (d_1192_v33_).cf47, (d_1193_v34_).cf0, not(d_1135_v1_.f29), self.f29)
                d_1194_v35_ = nw198_
                d_1195_v36_: _dafny.Map
                d_1195_v36_ = _dafny.Map({not((d_1135_v1_.f29) == (self.f35)): d_1135_v1_.f35})
                d_1195_v36_ = (d_1195_v36_).set((self).f47, (self).f47)
                d_1196_v37_: _dafny.Array
                nw199_ = _dafny.Array(_dafny.Map({}), 7)
                d_1196_v37_ = nw199_
                d_1197_v38_: _dafny.Map
                d_1197_v38_ = _dafny.Map({self.f29: (self).f34})
                d_1198_v39_: _dafny.Set
                d_1198_v39_ = _dafny.Set({(d_1135_v1_).f30, (d_1135_v1_).f34, len(d_1197_v38_), (self).f48, (d_1194_v35_).f46})
                d_1199_v41_: _dafny.Map
                def iife72_():
                    coll34_ = _dafny.Set()
                    compr_34_: int
                    for compr_34_ in _dafny.IntegerRange(127, 328):
                        d_1200_v40_: int = compr_34_
                        if ((127) <= (d_1200_v40_)) and ((d_1200_v40_) < (328)):
                            coll34_ = coll34_.union(_dafny.Set([default__.safeModuloInt(d_1200_v40_, (d_1135_v1_).f34)]))
                    return _dafny.Set(coll34_)
                d_1199_v41_ = _dafny.Map({iife72_()
                : p1})
                index154_ = default__.safeIndex(739, (d_1196_v37_).length(0))
                (d_1196_v37_)[index154_] = (_dafny.Map({d_1198_v39_: p1})) | (d_1199_v41_)
                index155_ = default__.safeIndex(739, (d_1196_v37_).length(0))
                (d_1196_v37_)[index155_] = d_1199_v41_
                d_1201_v42_: C1
                nw200_ = C1()
                nw200_.ctor__(True, ((d_1195_v36_)[self.f35] if (self.f35) in (d_1195_v36_) else self.f29), d_1135_v1_.f35)
                d_1201_v42_ = nw200_
                d_1202_v43_: _dafny.Map
                d_1202_v43_ = _dafny.Map({(d_1135_v1_).f34: D10_DC28(d_1201_v42_)})
                d_1203_v44_: _dafny.Seq
                d_1203_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mo"))
                (d_1135_v1_).f29 = (len(d_1202_v43_)) <= (len(d_1203_v44_))
            elif True:
                (globalState).f26 = not(d_1135_v1_.f29)
                d_1204_v45_: _dafny.Seq
                d_1204_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqrfmscn"))
                index156_ = default__.safeIndex(824, (p1).length(0))
                (p1)[index156_] = (d_1204_v45_) != (d_1204_v45_)
                index157_ = default__.safeIndex(824, (p1).length(0))
                (p1)[index157_] = d_1135_v1_.f31
                d_1205_v46_: _dafny.Map
                d_1205_v46_ = _dafny.Map({d_1190_v31_: p0})
                d_1205_v46_ = (d_1205_v46_).set((d_1190_v31_).set(not(d_1135_v1_.f31), default__.abs((d_1135_v1_).f30)), (d_1204_v45_)[default__.safeIndex((d_1135_v1_).f34, len(d_1204_v45_))])
                (globalState).f26 = (default__.fm1(d_1135_v1_.f28, (d_1135_v1_).f34, globalState)) in ((self).f32)
                (self).f35 = d_1135_v1_.f28
        if not(((self).f32)[default__.safeIndex(-172, len((self).f32))]):
            r1 = p0
            (globalState).f13 = (560) * ((self).f34)
            d_1206_v47_: _dafny.Map
            d_1206_v47_ = _dafny.Map({self.f29: (self).f32})
            d_1207_v48_: D3
            d_1207_v48_ = D3_DC9(self.f35, (self).f47, p1, self.f28)
            d_1208_v49_: _dafny.Array
            nw201_ = _dafny.Array(None, 21)
            nw201_[int(0)] = (self).f32
            nw201_[int(1)] = (self).f32
            nw201_[int(2)] = (self).f32
            nw201_[int(3)] = (self).f32
            nw201_[int(4)] = (self).f32
            nw201_[int(5)] = ((d_1206_v47_)[(d_1207_v48_).cf21] if ((d_1207_v48_).cf21) in (d_1206_v47_) else (self).f32)
            nw201_[int(6)] = (self).f32
            nw201_[int(7)] = (self).f32
            nw201_[int(8)] = _dafny.SeqWithoutIsStrInference([False, (self).f47, self.f35, not(True)])
            nw201_[int(9)] = (self).f32
            nw201_[int(10)] = (self).f32
            nw201_[int(11)] = (self).f32
            nw201_[int(12)] = (self).f32
            nw201_[int(13)] = _dafny.SeqWithoutIsStrInference([True, self.f31])
            nw201_[int(14)] = (self).f32
            nw201_[int(15)] = (self).f32
            nw201_[int(16)] = (self).f32
            nw201_[int(17)] = (self).f32
            nw201_[int(18)] = _dafny.SeqWithoutIsStrInference([True, self.f28, self.f28])
            nw201_[int(19)] = (self).f32
            nw201_[int(20)] = _dafny.SeqWithoutIsStrInference([(self).f47])
            d_1208_v49_ = nw201_
            d_1209_v50_: C3
            nw202_ = C3()
            nw202_.ctor__(default__.safeDivisionInt((self).f30, (self).f30), d_1208_v49_, p0, ((self).f48) * ((self).f30), ((self).f47) == (self.f28), (self).f32, (self).f33, (self).f30, False, (self).f47, self.f35)
            d_1209_v50_ = nw202_
            (globalState).f12 = (_dafny.SeqWithoutIsStrInference([(self).f30 for d_1210_i6_ in range(default__.abs(816))])) < (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaqupmrc"))) for d_1211_i7_ in range(default__.abs(685))]))
            d_1212_v51_: _dafny.Seq
            d_1212_v51_ = _dafny.SeqWithoutIsStrInference([(d_1209_v50_).f44])
            d_1213_v52_: _dafny.MultiSet
            d_1213_v52_ = _dafny.MultiSet([d_1212_v51_, _dafny.SeqWithoutIsStrInference([(d_1209_v50_).f44 for d_1214_i8_ in range(default__.abs(618))]), (p2).cf32])
            d_1215_v53_: _dafny.Map
            d_1215_v53_ = _dafny.Map({self.f35: d_1213_v52_})
            d_1216_v54_: _dafny.Set
            d_1216_v54_ = _dafny.Set({(self).f34, len(d_1212_v51_), (self).f34})
            d_1215_v53_ = (d_1215_v53_).set((_dafny.Set({(self).f34})).isdisjoint(d_1216_v54_), d_1213_v52_)
        elif True:
            d_1217_v55_: _dafny.Map
            d_1217_v55_ = _dafny.Map({self.f28: ((0) - ((self).f48)) - ((self).f30)})
            d_1218_v56_: _dafny.Seq
            d_1218_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axjmgmig"))
            d_1217_v55_ = (d_1217_v55_).set((d_1218_v56_) <= (d_1218_v56_), (self).f30)
            d_1219_v57_: _dafny.Map
            d_1219_v57_ = _dafny.Map({(self).f48: p0})
            d_1219_v57_ = (d_1219_v57_).set((self).f30, p0)
            d_1220_v58_: _dafny.Seq
            d_1220_v58_ = _dafny.SeqWithoutIsStrInference([(self).f48, (self).f48])
            d_1221_v59_: _dafny.Map
            d_1221_v59_ = _dafny.Map({(len(_dafny.Set({self.f35}))) - ((self).f34): d_1220_v58_})
            d_1222_v60_: _dafny.Map
            d_1222_v60_ = _dafny.Map({(self).f34: self.f35})
            d_1221_v59_ = (d_1221_v59_).set(len(d_1222_v60_), d_1220_v58_)
            (globalState).f1 = not(self.f35)
            d_1223_v61_: _dafny.Seq
            d_1223_v61_ = _dafny.SeqWithoutIsStrInference([d_1218_v56_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yumismnte"))])
            d_1224_v63_: _dafny.Map
            d_1224_v63_ = _dafny.Map({(self).f30: 692})
            d_1225_v64_: _dafny.MultiSet
            d_1225_v64_ = _dafny.MultiSet([(self).f47])
            d_1226_v65_: _dafny.Array
            def lambda44_(d_1227_i9_):
                return (d_1227_i9_) + ((D7_DC19(self.f35, len(_dafny.SeqWithoutIsStrInference([(self).f30 for d_1228_i10_ in range(default__.abs(146))])), (self).f34)).cf37)

            init25_ = lambda44_
            nw203_ = _dafny.Array(None, 9)
            for i0_25_ in range(nw203_.length(0)):
                nw203_[i0_25_] = init25_(i0_25_)
            d_1226_v65_ = nw203_
            d_1229_v66_: D0
            d_1229_v66_ = D0_DC0(self.f35)
            d_1230_v67_: _dafny.Seq
            d_1230_v67_ = _dafny.SeqWithoutIsStrInference([d_1229_v66_])
            d_1231_v68_: _dafny.Set
            d_1231_v68_ = _dafny.Set({(self).f48})
            d_1232_v69_: D2
            d_1232_v69_ = D2_DC6((0) - ((self).f30), self.f35, (self).f34, (self).f47, (self).f34)
            d_1233_v70_: _dafny.MultiSet
            d_1233_v70_ = _dafny.MultiSet([(self).f48, (self).f34])
            d_1234_v71_: _dafny.Array
            nw204_ = _dafny.Array(None, 29)
            def iife73_():
                coll35_ = _dafny.Set()
                compr_35_: _dafny.Seq
                for compr_35_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa"))})).Elements:
                    d_1235_v62_: _dafny.Seq = compr_35_
                    if (d_1235_v62_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa"))})):
                        coll35_ = coll35_.union(_dafny.Set([d_1235_v62_]))
                return _dafny.Set(coll35_)
            nw204_[int(0)] = (len((self).fm4(True, _dafny.Set({(self).f48, (self).f48}), len(((d_1223_v61_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuntapvl"))), len(d_1223_v61_))]).set(default__.safeIndex((self).f48, len((d_1223_v61_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuntapvl"))), len(d_1223_v61_))])), _dafny.CodePoint('j'))), self.f28, globalState))) + (len(iife73_()
            ))
            nw204_[int(1)] = ((d_1224_v63_)[(self).f48] if ((self).f48) in (d_1224_v63_) else (self).f30)
            nw204_[int(2)] = ((self).f48) - ((self).f30)
            nw204_[int(3)] = ((0) - (len(d_1220_v58_))) * (((d_1225_v64_)[(D10_DC31(self.f29, (self).f47, (self).f48)).cf61] if ((D10_DC31(self.f29, (self).f47, (self).f48)).cf61) in (d_1225_v64_) else (self).f34))
            nw204_[int(4)] = (0) - (default__.safeModuloInt((self).f34, (self).f30))
            nw204_[int(5)] = (self).f30
            nw204_[int(6)] = default__.safeDivisionInt((self).f30, (self).f34)
            nw204_[int(7)] = len(_dafny.SeqWithoutIsStrInference([d_1226_v65_, d_1226_v65_, d_1226_v65_, d_1226_v65_]))
            nw204_[int(8)] = len(d_1230_v67_)
            nw204_[int(9)] = default__.safeModuloInt((self).f30, (self).f48)
            nw204_[int(10)] = (self).f48
            nw204_[int(11)] = ((self).f34) - ((self).f34)
            nw204_[int(12)] = len(d_1231_v68_)
            nw204_[int(13)] = default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({self.f29: (self).f48})]))), (self).f30)
            nw204_[int(14)] = ((d_1232_v69_).cf11) * ((0) - ((self).f30))
            nw204_[int(15)] = (self).f34
            nw204_[int(16)] = (self).f34
            nw204_[int(17)] = (self).f48
            nw204_[int(18)] = (self).f30
            nw204_[int(19)] = (self).f30
            nw204_[int(20)] = (self).f30
            nw204_[int(21)] = (self).f30
            nw204_[int(22)] = (0) - ((self).f34)
            nw204_[int(23)] = (self).f48
            nw204_[int(24)] = len(d_1231_v68_)
            nw204_[int(25)] = -628
            nw204_[int(26)] = ((d_1233_v70_)[(self).f30] if ((self).f30) in (d_1233_v70_) else (self).f34)
            nw204_[int(27)] = ((self).f48) * (default__.fm39((self).f47, d_1233_v70_, self.f31, globalState))
            nw204_[int(28)] = len(d_1217_v55_)
            d_1234_v71_ = nw204_
            index158_ = default__.safeIndex(384, (d_1234_v71_).length(0))
            (d_1234_v71_)[index158_] = (self).f48
            d_1236_v72_: _dafny.Map
            d_1236_v72_ = _dafny.Map({not(True): d_1226_v65_})
            index159_ = default__.safeIndex(384, (d_1234_v71_).length(0))
            rhs156_ = 808
            rhs157_ = (len(d_1236_v72_)) * (default__.fm39((self).f47, d_1233_v70_, self.f31, globalState))
            rhs158_ = (0) - ((self).f34)
            lhs117_ = d_1234_v71_
            lhs118_ = default__.safeIndex(384, (d_1234_v71_).length(0))
            lhs119_ = globalState
            lhs120_ = globalState
            lhs117_[lhs118_] = rhs156_
            lhs119_.f13 = rhs157_
            lhs120_.f13 = rhs158_
        d_1237_v73_: D6
        d_1237_v73_ = D6_DC16(p1)
        d_1238_v74_: D6
        d_1238_v74_ = D6_DC17(d_1237_v73_)
        d_1238_v74_ = d_1238_v74_
        def lambda45_(source21_):
            if source21_.is_DC16:
                d_1239___mcc_h9_ = source21_.cf33
                d_1240_cf33_ = d_1239___mcc_h9_
                return (self).f47
            elif source21_.is_DC15:
                d_1241___mcc_h10_ = source21_.cf32
                d_1242_cf32_ = d_1241___mcc_h10_
                return False
            elif True:
                d_1243___mcc_h11_ = source21_.cf34
                d_1244_cf34_ = d_1243___mcc_h11_
                return (self).f47

        (self).f28 = lambda45_(p2)
        d_1245_v75_: _dafny.Map
        d_1245_v75_ = _dafny.Map({(self).f48: self.f35})
        d_1245_v75_ = d_1245_v75_
        d_1246_v76_: _dafny.Set
        d_1246_v76_ = _dafny.Set({p1})
        d_1247_v77_: _dafny.Map
        d_1247_v77_ = _dafny.Map({(self).f34: (self).f34})
        d_1248_v78_: _dafny.Map
        d_1248_v78_ = _dafny.Map({(p1) not in (d_1246_v76_): d_1247_v77_})
        d_1248_v78_ = (default__.fm50(d_1245_v75_, globalState)) | (d_1248_v78_)
        d_1249_v79_: _dafny.Set
        d_1249_v79_ = _dafny.Set({(self).f47, True})
        d_1250_v80_: _dafny.Map
        d_1250_v80_ = _dafny.Map({default__.fm1(self.f31, (self).f30, globalState): (self).f30})
        d_1251_v81_: _dafny.MultiSet
        d_1251_v81_ = _dafny.MultiSet([((d_1250_v80_)[(self).f47] if ((self).f47) in (d_1250_v80_) else (self).f30), (self).f30, (self).f34, (0) - ((self).f30), (self).f30])
        r0 = (len(d_1249_v79_)) == ((d_1251_v81_).cardinality)
        d_1252_v82_: _dafny.Seq
        d_1252_v82_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        r1 = (d_1252_v82_)[default__.safeIndex((self).f48, len(d_1252_v82_))]
        return r0, r1

    @property
    def f47(self):
        return self._f47
    @property
    def f48(self):
        return self._f48

class C6(T4, T2, T3, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f36: str = _dafny.CodePoint('D')
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f36(self):
        return self._f36
    @f36.setter
    def f36(self, value):
        self._f36 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f36, f34, f35, f32, f33, f30, f31, f28, f29):
        (self).f36 = f36
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm7(self, p0, p1, globalState):
        return self.f31

    def fm6(self, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')]))

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife74_():
            coll36_ = _dafny.Set()
            compr_36_: int
            for compr_36_ in _dafny.IntegerRange(924, 550):
                d_1253_v0_: int = compr_36_
                if ((924) <= (d_1253_v0_)) and ((d_1253_v0_) < (550)):
                    coll36_ = coll36_.union(_dafny.Set([(d_1253_v0_) * (len((self).f32))]))
            return _dafny.Set(coll36_)
        return ((iife74_()
        ) - (_dafny.Set({-281}))).intersection(_dafny.Set({543, len(_dafny.SeqWithoutIsStrInference([(self).f34, (self).f34, (self).f30])), (self).f30, (self).f30}))

    def fm5(self, p0, p1, p2, p3, globalState):
        return ((self).f33)[default__.safeIndex(((self).f30 if False else (self).f30), len((self).f33))]

    def fm3(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crmrunqsx"))

    def fm2(self, globalState):
        return _dafny.Map({self.f29: (self).f30})

    def fm27(self, p0, p1, p2, globalState):
        if (False if self.f29 else self.f29):
            return (0) - ((self).f30)
        elif True:
            return (self).f34

    def fm28(self, p0, p1, globalState):
        return -675

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        hi7_ = (self).f34
        for d_1254_i0_ in range((self).f30, hi7_):
            d_1255_v0_: _dafny.Array
            nw205_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_1255_v0_ = nw205_
            d_1256_v1_: _dafny.Seq
            d_1256_v1_ = _dafny.SeqWithoutIsStrInference([d_1255_v0_])
            d_1257_v2_: _dafny.Map
            d_1257_v2_ = _dafny.Map({(d_1254_i0_) * (d_1254_i0_): (d_1256_v1_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([380])), len(d_1256_v1_))]})
            d_1258_v3_: _dafny.Map
            d_1258_v3_ = _dafny.Map({self.f35: d_1255_v0_})
            d_1257_v2_ = (d_1257_v2_).set(default__.safeDivisionInt(d_1254_i0_, -558), ((d_1258_v3_)[self.f28] if (self.f28) in (d_1258_v3_) else d_1255_v0_))
            d_1259_v4_: _dafny.Seq
            d_1259_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xku"))
            d_1260_v5_: _dafny.Map
            d_1260_v5_ = _dafny.Map({-301: d_1259_v4_})
            if (d_1254_i0_) != (len(d_1260_v5_)):
                d_1261_v6_: _dafny.Seq
                d_1261_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({476, d_1254_i0_})])
                d_1262_v7_: _dafny.Map
                d_1262_v7_ = _dafny.Map({(d_1261_v6_)[default__.safeIndex((self).f30, len(d_1261_v6_))]: _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f32: (self).f34}))])})
                d_1263_v8_: _dafny.Set
                d_1263_v8_ = _dafny.Set({12, d_1254_i0_, d_1254_i0_})
                d_1264_v9_: _dafny.Array
                nw206_ = _dafny.Array(_dafny.Seq({}), 1)
                d_1264_v9_ = nw206_
                d_1265_v10_: _dafny.Map
                d_1265_v10_ = _dafny.Map({self.f31: self.f28})
                d_1266_v11_: T4
                nw207_ = C3()
                nw207_.ctor__(d_1254_i0_, d_1264_v9_, (d_1259_v4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f34])), len(d_1259_v4_))], (self).f34, self.f29, (self).f32, _dafny.SeqWithoutIsStrInference([(self).f32 for d_1267_i1_ in range(default__.abs(658))]), len(_dafny.SeqWithoutIsStrInference([len(d_1265_v10_), 387])), self.f29, self.f31, self.f35)
                d_1266_v11_ = nw207_
                d_1268_v12_: _dafny.Seq
                d_1268_v12_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yor"))), (self).f30, len(_dafny.Map({(self).f34: d_1266_v11_}))])
                d_1262_v7_ = (d_1262_v7_).set((d_1263_v8_ if self.f29 else d_1263_v8_), (d_1268_v12_) + (d_1268_v12_))
                d_1269_v13_: _dafny.Seq
                d_1269_v13_ = _dafny.SeqWithoutIsStrInference([d_1264_v9_, d_1264_v9_])
                d_1270_v14_: _dafny.Map
                d_1270_v14_ = _dafny.Map({self.f36: _dafny.CodePoint('i')})
                d_1271_v15_: C3
                nw208_ = C3()
                nw208_.ctor__(d_1254_i0_, (d_1269_v13_)[default__.safeIndex((self).f34, len(d_1269_v13_))], ((d_1270_v14_)[d_1266_v11_.f36] if (d_1266_v11_.f36) in (d_1270_v14_) else _dafny.CodePoint('u')), 35, self.f28, _dafny.SeqWithoutIsStrInference([d_1266_v11_.f29, d_1266_v11_.f28]), (self).f33, d_1254_i0_, (not(d_1266_v11_.f31)) or (self.f28), True, self.f31)
                d_1271_v15_ = nw208_
                d_1272_v16_: _dafny.Array
                nw209_ = _dafny.Array(None, 10)
                d_1272_v16_ = nw209_
                d_1273_v17_: _dafny.MultiSet
                d_1273_v17_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twqhca"))])
                d_1274_v18_: _dafny.Map
                d_1274_v18_ = _dafny.Map({not((d_1259_v4_) in (d_1273_v17_)): d_1272_v16_})
                rhs159_ = ((self).f30) * (((d_1266_v11_).f34 if d_1266_v11_.f29 else (d_1271_v15_).f44))
                rhs160_ = ((d_1274_v18_)[not(d_1266_v11_.f28)] if (not(d_1266_v11_.f28)) in (d_1274_v18_) else d_1272_v16_)
                rhs161_ = not(not(not((d_1266_v11_.f36) in (d_1259_v4_))))
                lhs121_ = globalState
                r1 = rhs159_
                d_1272_v16_ = rhs160_
                lhs121_.f26 = rhs161_
                d_1275_v19_: C2
                nw210_ = C2()
                nw210_.ctor__((self).f30, (self).f30, d_1266_v11_.f29, d_1266_v11_.f31, d_1266_v11_.f35)
                d_1275_v19_ = nw210_
                d_1276_v20_: _dafny.Map
                d_1276_v20_ = _dafny.Map({d_1266_v11_.f31: d_1259_v4_})
                d_1277_v21_: _dafny.Map
                d_1277_v21_ = _dafny.Map({(len(_dafny.Map({d_1275_v19_: len(d_1259_v4_)}))) + ((d_1275_v19_).f46): (d_1276_v20_ if self.f31 else _dafny.Map({self.f29: d_1259_v4_}))})
                d_1278_v22_: D5
                d_1278_v22_ = D5_DC12(_dafny.SeqWithoutIsStrInference([d_1266_v11_.f36 for d_1279_i2_ in range(default__.abs(-852))]))
                d_1277_v21_ = (d_1277_v21_).set((d_1271_v15_).f44, (_dafny.Map({d_1266_v11_.f28: (d_1278_v22_).cf26}) if not(self.f31) else d_1276_v20_))
                rhs162_ = self.f31
                rhs163_ = d_1254_i0_
                rhs164_ = (308) * ((self).f34)
                lhs122_ = self
                lhs123_ = globalState
                lhs124_ = globalState
                lhs122_.f35 = rhs162_
                lhs123_.f13 = rhs163_
                lhs124_.f13 = rhs164_
            elif True:
                d_1280_v23_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Seq({}), 10)
                d_1280_v23_ = nw211_
                d_1281_v24_: C3
                nw212_ = C3()
                nw212_.ctor__(-200, d_1280_v23_, _dafny.CodePoint('s'), (self).fm28((self).f30, True, globalState), self.f35, (self).f32, (self).f33, d_1254_i0_, ((self).f32) <= ((self).f32), self.f29, ((self).f34) == (len(d_1259_v4_)))
                d_1281_v24_ = nw212_
                d_1259_v4_ = _dafny.SeqWithoutIsStrInference([self.f36 for d_1282_i3_ in range(default__.abs(495))])
                d_1283_v25_: _dafny.Map
                d_1283_v25_ = _dafny.Map({d_1281_v24_: len((self).f32)})
                d_1284_v26_: _dafny.Seq
                d_1284_v26_ = _dafny.SeqWithoutIsStrInference([d_1283_v25_, d_1283_v25_, d_1283_v25_, d_1283_v25_])
                d_1285_v27_: _dafny.Map
                d_1285_v27_ = _dafny.Map({self.f29: (self).f30})
                r1 = (0) - ((len(d_1285_v27_) if (d_1284_v26_) <= (d_1284_v26_) else d_1254_i0_))
                d_1286_v28_: D5
                d_1286_v28_ = D5_DC14(self.f35, (self).f30)
                d_1287_v29_: _dafny.Array
                def lambda46_(d_1288_i4_):
                    return self.f36

                init26_ = lambda46_
                nw213_ = _dafny.Array(None, 19)
                for i0_26_ in range(nw213_.length(0)):
                    nw213_[i0_26_] = init26_(i0_26_)
                d_1287_v29_ = nw213_
                d_1289_v30_: C0
                nw214_ = C0()
                nw214_.ctor__((d_1287_v29_ if (d_1286_v28_).cf30 else d_1287_v29_))
                d_1289_v30_ = nw214_
                (globalState).f13 = (d_1281_v24_).f44
            d_1290_v31_: _dafny.Map
            d_1290_v31_ = _dafny.Map({_dafny.CodePoint('a'): self.f29})
            d_1291_v32_: _dafny.MultiSet
            d_1291_v32_ = _dafny.MultiSet([605, (self).f34, (self).f30, len(d_1290_v31_)])
            d_1292_v33_: _dafny.Seq
            d_1292_v33_ = _dafny.SeqWithoutIsStrInference([True, (d_1291_v32_).isdisjoint(d_1291_v32_)])
            d_1292_v33_ = ((self).f32) + (_dafny.SeqWithoutIsStrInference([False, not(self.f35)]))
            d_1293_v34_: _dafny.Set
            d_1293_v34_ = _dafny.Set({(self).f30})
            def iife75_():
                coll37_ = _dafny.Set()
                compr_37_: int
                for compr_37_ in (d_1293_v34_).Elements:
                    d_1294_v35_: int = compr_37_
                    if (d_1294_v35_) in (d_1293_v34_):
                        coll37_ = coll37_.union(_dafny.Set([default__.safeModuloInt(d_1294_v35_, len(_dafny.Map({_dafny.CodePoint('v'): -246})))]))
                return _dafny.Set(coll37_)
            d_1293_v34_ = ((d_1293_v34_) - (iife75_()
            )) - (d_1293_v34_)
        d_1295_v36_: _dafny.Array
        nw215_ = _dafny.Array(int(0), 1)
        d_1295_v36_ = nw215_
        d_1296_v37_: _dafny.Map
        d_1296_v37_ = _dafny.Map({d_1295_v36_: (self).f30})
        d_1297_v38_: D7
        d_1297_v38_ = D7_DC18((d_1296_v37_).set(d_1295_v36_, (self).f30))
        pat_let_tv13_ = d_1296_v37_
        index160_ = default__.safeIndex(301, (d_1295_v36_).length(0))
        def iife76_(_pat_let19_0):
            def iife77_(d_1298_dt__update__tmp_h0_):
                def iife78_(_pat_let20_0):
                    def iife79_(d_1299_dt__update_hcf35_h0_):
                        return D7_DC18(d_1299_dt__update_hcf35_h0_)
                    return iife79_(_pat_let20_0)
                return iife78_(pat_let_tv13_)
            return iife77_(_pat_let19_0)
        (d_1295_v36_)[index160_] = len((iife76_(d_1297_v38_)).cf35)
        d_1300_v39_: _dafny.MultiSet
        d_1300_v39_ = _dafny.MultiSet([True])
        d_1301_v40_: D0
        d_1301_v40_ = D0_DC1(d_1300_v39_)
        d_1302_v41_: _dafny.Map
        d_1302_v41_ = _dafny.Map({d_1301_v40_: self.f29})
        d_1303_v42_: D8
        d_1303_v42_ = D8_DC21(_dafny.Map({d_1301_v40_: self.f35}))
        d_1304_v44_: _dafny.Set
        d_1304_v44_ = _dafny.Set({d_1301_v40_})
        d_1305_v46_: _dafny.Seq
        d_1305_v46_ = _dafny.SeqWithoutIsStrInference([(D8_DC21(d_1302_v41_)).cf40])
        pat_let_tv14_ = d_1300_v39_
        d_1306_v47_: _dafny.Array
        nw216_ = _dafny.Array(None, 25)
        nw216_[int(0)] = d_1302_v41_
        nw216_[int(1)] = (_dafny.Map({d_1301_v40_: self.f28}) if self.f28 else d_1302_v41_)
        nw216_[int(2)] = (_dafny.Map({d_1301_v40_: not(not(self.f35))})) | (d_1302_v41_)
        nw216_[int(3)] = d_1302_v41_
        nw216_[int(4)] = ((d_1303_v42_).cf40) | (d_1302_v41_)
        nw216_[int(5)] = d_1302_v41_
        nw216_[int(6)] = d_1302_v41_
        nw216_[int(7)] = d_1302_v41_
        nw216_[int(8)] = d_1302_v41_
        nw216_[int(9)] = _dafny.Map({D0_DC1(d_1300_v39_): default__.fm1(self.f31, (self).f34, globalState)})
        nw216_[int(10)] = d_1302_v41_
        nw216_[int(11)] = default__.fm35(self.f35, globalState)
        nw216_[int(12)] = d_1302_v41_
        nw216_[int(13)] = (_dafny.Map({default__.fm36(893, (self).f34, self.f31, globalState): not(self.f35)})) | (d_1302_v41_)
        def iife80_():
            coll38_ = _dafny.Map()
            compr_38_: D0
            for compr_38_ in (d_1304_v44_).Elements:
                d_1307_v43_: D0 = compr_38_
                if (d_1307_v43_) in (d_1304_v44_):
                    coll38_[d_1307_v43_] = self.f35
            return _dafny.Map(coll38_)
        nw216_[int(14)] = (d_1302_v41_ if self.f28 else iife80_()
        )
        def iife81_():
            coll39_ = _dafny.Map()
            compr_39_: D0
            for compr_39_ in (d_1302_v41_).keys.Elements:
                d_1308_v45_: D0 = compr_39_
                if (d_1308_v45_) in (d_1302_v41_):
                    coll39_[d_1308_v45_] = True
            return _dafny.Map(coll39_)
        nw216_[int(15)] = (d_1302_v41_) | (iife81_()
        )
        nw216_[int(16)] = (d_1305_v46_)[default__.safeIndex((0) - ((self).f34), len(d_1305_v46_))]
        def iife82_(_pat_let21_0):
            def iife83_(d_1309_dt__update__tmp_h1_):
                def iife84_(_pat_let22_0):
                    def iife85_(d_1310_dt__update_hcf1_h0_):
                        return D0_DC1(d_1310_dt__update_hcf1_h0_)
                    return iife85_(_pat_let22_0)
                return iife84_(pat_let_tv14_)
            return iife83_(_pat_let21_0)
        nw216_[int(17)] = _dafny.Map({iife82_(D0_DC1(d_1300_v39_)): self.f28})
        nw216_[int(18)] = d_1302_v41_
        nw216_[int(19)] = d_1302_v41_
        nw216_[int(20)] = d_1302_v41_
        nw216_[int(21)] = d_1302_v41_
        nw216_[int(22)] = default__.fm35(self.f35, globalState)
        nw216_[int(23)] = d_1302_v41_
        nw216_[int(24)] = _dafny.Map({d_1301_v40_: self.f28})
        d_1306_v47_ = nw216_
        d_1311_v48_: _dafny.Array
        nw217_ = _dafny.Array(False, 13)
        d_1311_v48_ = nw217_
        index161_ = default__.safeIndex(592, (d_1311_v48_).length(0))
        (d_1311_v48_)[index161_] = self.f35
        d_1312_v49_: _dafny.Array
        nw218_ = _dafny.Array(_dafny.CodePoint('D'), 29)
        d_1312_v49_ = nw218_
        d_1313_v50_: _dafny.Seq
        d_1313_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehackpxcv"))
        d_1314_v51_: _dafny.Map
        d_1314_v51_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([self.f36 for d_1315_i5_ in range(default__.abs(781))]): self.f29})
        index162_ = default__.safeIndex(984, (d_1312_v49_).length(0))
        (d_1312_v49_)[index162_] = default__.fm34(len(d_1313_v50_), ((d_1300_v39_)[self.f35] if (self.f35) in (d_1300_v39_) else len(default__.fm37(self.f36, self.f31, globalState))), (self).f30, ((d_1314_v51_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exfveuas"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exfveuas"))) in (d_1314_v51_) else True), globalState)
        index163_ = default__.safeIndex(301, (d_1295_v36_).length(0))
        index164_ = default__.safeIndex(592, (d_1311_v48_).length(0))
        index165_ = default__.safeIndex(984, (d_1312_v49_).length(0))
        rhs165_ = ((self).fm28((self).f34, self.f35, globalState)) + ((self).fm28((self).f30, self.f29, globalState))
        rhs166_ = d_1306_v47_
        rhs167_ = not(True)
        rhs168_ = (self.f29) or (self.f28)
        rhs169_ = _dafny.CodePoint('u')
        lhs125_ = d_1295_v36_
        lhs126_ = default__.safeIndex(301, (d_1295_v36_).length(0))
        lhs127_ = globalState
        lhs128_ = d_1311_v48_
        lhs129_ = default__.safeIndex(592, (d_1311_v48_).length(0))
        lhs130_ = d_1312_v49_
        lhs131_ = default__.safeIndex(984, (d_1312_v49_).length(0))
        lhs125_[lhs126_] = rhs165_
        d_1306_v47_ = rhs166_
        lhs127_.f26 = rhs167_
        lhs128_[lhs129_] = rhs168_
        lhs130_[lhs131_] = rhs169_
        d_1316_v52_: _dafny.MultiSet
        d_1316_v52_ = _dafny.MultiSet([(self).fm27((d_1311_v48_)[default__.safeIndex(592, (d_1311_v48_).length(0))], (d_1295_v36_)[default__.safeIndex(301, (d_1295_v36_).length(0))], (self).f30, globalState)])
        d_1317_v53_: _dafny.Seq
        d_1317_v53_ = _dafny.SeqWithoutIsStrInference([(d_1295_v36_)[default__.safeIndex(301, (d_1295_v36_).length(0))], (d_1295_v36_)[default__.safeIndex(301, (d_1295_v36_).length(0))], (self).f30, len(_dafny.SeqWithoutIsStrInference([-627, (self).f34])), (0) - ((self).f30)])
        d_1318_v54_: C2
        nw219_ = C2()
        nw219_.ctor__((d_1295_v36_)[default__.safeIndex(301, (d_1295_v36_).length(0))], ((d_1316_v52_)[(d_1317_v53_)[default__.safeIndex((d_1295_v36_)[default__.safeIndex(301, (d_1295_v36_).length(0))], len(d_1317_v53_))]] if ((d_1317_v53_)[default__.safeIndex((d_1295_v36_)[default__.safeIndex(301, (d_1295_v36_).length(0))], len(d_1317_v53_))]) in (d_1316_v52_) else (self).f34), self.f28, self.f35, not(self.f29))
        d_1318_v54_ = nw219_
        d_1319_v55_: _dafny.Map
        d_1319_v55_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f30 for d_1320_i6_ in range(default__.abs(799))]): len(d_1317_v53_)})
        rhs170_ = len(((d_1313_v50_ if (len((self).f32)) >= ((self).f30) else d_1313_v50_)).set(default__.safeIndex(((d_1318_v54_).f46) - ((0) - (len(d_1319_v55_))), len((d_1313_v50_ if (len((self).f32)) >= ((self).f30) else d_1313_v50_))), (d_1312_v49_)[default__.safeIndex(984, (d_1312_v49_).length(0))]))
        rhs171_ = ((d_1313_v50_) + (d_1313_v50_)) + (_dafny.SeqWithoutIsStrInference([self.f36]))
        rhs172_ = (len((d_1313_v50_) + ((_dafny.SeqWithoutIsStrInference([(d_1312_v49_)[default__.safeIndex(984, (d_1312_v49_).length(0))] for d_1321_i7_ in range(default__.abs(178))])).set(default__.safeIndex((self).f34, len(_dafny.SeqWithoutIsStrInference([(d_1312_v49_)[default__.safeIndex(984, (d_1312_v49_).length(0))] for d_1322_i7_ in range(default__.abs(178))]))), (d_1312_v49_)[default__.safeIndex(984, (d_1312_v49_).length(0))]))) if self.f35 else (self).f34)
        rhs173_ = not (self.f35) or (self.f28)
        lhs132_ = globalState
        lhs133_ = globalState
        r1 = rhs170_
        d_1313_v50_ = rhs171_
        lhs132_.f13 = rhs172_
        lhs133_.f12 = rhs173_
        index166_ = default__.safeIndex(592, (d_1311_v48_).length(0))
        (d_1311_v48_)[index166_] = self.f29
        index167_ = default__.safeIndex(592, (d_1311_v48_).length(0))
        (d_1311_v48_)[index167_] = (d_1318_v54_).fm30(False, globalState)
        r0 = self.f31
        r1 = (self).f34
        r2 = d_1313_v50_
        return r0, r1, r2

    def m11(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_1323_v0_: D8
        d_1323_v0_ = D8_DC22((self).f34, (self).f34)
        d_1324_v1_: D8
        d_1324_v1_ = D8_DC25(D8_DC25(d_1323_v0_))
        source22_ = d_1324_v1_
        if source22_.is_DC22:
            d_1325___mcc_h0_ = source22_.cf41
            d_1326___mcc_h1_ = source22_.cf42
            d_1327_cf42_ = d_1326___mcc_h1_
            d_1328_cf41_ = d_1325___mcc_h0_
            d_1329_v2_: _dafny.Array
            nw220_ = _dafny.Array(_dafny.Map({}), 24)
            d_1329_v2_ = nw220_
            d_1330_v3_: _dafny.Array
            nw221_ = _dafny.Array(int(0), 1)
            d_1330_v3_ = nw221_
            index168_ = default__.safeIndex(745, (d_1329_v2_).length(0))
            (d_1329_v2_)[index168_] = _dafny.Map({not(p0): d_1330_v3_})
            d_1331_v4_: _dafny.Map
            d_1331_v4_ = _dafny.Map({self.f28: d_1330_v3_})
            index169_ = default__.safeIndex(745, (d_1329_v2_).length(0))
            (d_1329_v2_)[index169_] = (d_1331_v4_) | (d_1331_v4_)
            d_1332_v5_: _dafny.Array
            def lambda47_(d_1333_i0_):
                return False

            init27_ = lambda47_
            nw222_ = _dafny.Array(None, 3)
            for i0_27_ in range(nw222_.length(0)):
                nw222_[i0_27_] = init27_(i0_27_)
            d_1332_v5_ = nw222_
            d_1334_v6_: D1
            d_1334_v6_ = D1_DC2(d_1332_v5_)
            d_1335_v7_: D1
            d_1335_v7_ = D1_DC4(d_1334_v6_)
            source23_ = d_1335_v7_
            if source23_.is_DC3:
                d_1336___mcc_h12_ = source23_.cf3
                d_1337___mcc_h13_ = source23_.cf4
                d_1338___mcc_h14_ = source23_.cf5
                d_1339___mcc_h15_ = source23_.cf6
                d_1340_cf6_ = d_1339___mcc_h15_
                d_1341_cf5_ = d_1338___mcc_h14_
                d_1342_cf4_ = d_1337___mcc_h13_
                d_1343_cf3_ = d_1336___mcc_h12_
                d_1344_v8_: _dafny.Seq
                d_1344_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyollv"))
                d_1345_v9_: _dafny.Seq
                d_1345_v9_ = _dafny.SeqWithoutIsStrInference([len(d_1344_v8_)])
                d_1346_v10_: int
                out56_: int
                out56_ = default__.m0(d_1345_v9_, not(self.f29), d_1328_cf41_, p0, globalState)
                d_1346_v10_ = out56_
                d_1347_v11_: _dafny.Array
                nw223_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_1347_v11_ = nw223_
                index170_ = default__.safeIndex(346, (d_1347_v11_).length(0))
                (d_1347_v11_)[index170_] = (d_1332_v5_ if self.f31 else d_1332_v5_)
                index171_ = default__.safeIndex(577, (d_1330_v3_).length(0))
                (d_1330_v3_)[index171_] = d_1342_cf4_
                d_1348_v12_: _dafny.MultiSet
                d_1348_v12_ = _dafny.MultiSet([not(self.f35)])
                index172_ = default__.safeIndex(346, (d_1347_v11_).length(0))
                index173_ = default__.safeIndex(577, (d_1330_v3_).length(0))
                rhs174_ = d_1332_v5_
                rhs175_ = ((d_1348_v12_)[self.f31] if (self.f31) in (d_1348_v12_) else (self).f30)
                rhs176_ = d_1340_cf6_
                rhs177_ = d_1324_v1_
                lhs134_ = d_1347_v11_
                lhs135_ = default__.safeIndex(346, (d_1347_v11_).length(0))
                lhs136_ = d_1330_v3_
                lhs137_ = default__.safeIndex(577, (d_1330_v3_).length(0))
                lhs134_[lhs135_] = rhs174_
                r2 = rhs175_
                lhs136_[lhs137_] = rhs176_
                d_1324_v1_ = rhs177_
                d_1344_v8_ = d_1344_v8_
                r2 = d_1328_cf41_
            elif source23_.is_DC2:
                d_1349___mcc_h16_ = source23_.cf2
                d_1350_cf2_ = d_1349___mcc_h16_
                d_1351_v13_: _dafny.Array
                nw224_ = _dafny.Array(None, 6)
                d_1351_v13_ = nw224_
                d_1352_v14_: _dafny.MultiSet
                d_1352_v14_ = _dafny.MultiSet([d_1327_cf42_])
                d_1353_v15_: _dafny.Set
                d_1353_v15_ = _dafny.Set({d_1352_v14_})
                d_1354_v16_: C2
                nw225_ = C2()
                nw225_.ctor__(643, len(d_1353_v15_), self.f28, p0, not(self.f28))
                d_1354_v16_ = nw225_
                index174_ = default__.safeIndex(680, (d_1351_v13_).length(0))
                (d_1351_v13_)[index174_] = d_1354_v16_
                d_1355_v17_: D9
                d_1355_v17_ = D9_DC26(d_1354_v16_)
                index175_ = default__.safeIndex(680, (d_1351_v13_).length(0))
                (d_1351_v13_)[index175_] = (d_1355_v17_).cf52
                d_1356_v18_: _dafny.MultiSet
                d_1356_v18_ = _dafny.MultiSet([(self).f32, (self).f32])
                (globalState).f13 = ((self).f30) - ((d_1356_v18_).cardinality)
                d_1357_v19_: _dafny.Set
                d_1357_v19_ = _dafny.Set({(self).f30})
                index176_ = default__.safeIndex(529, (d_1330_v3_).length(0))
                (d_1330_v3_)[index176_] = len(d_1357_v19_)
                d_1358_v20_: _dafny.Map
                d_1358_v20_ = _dafny.Map({self.f31: (d_1354_v16_).f46})
                d_1359_v21_: _dafny.Map
                d_1359_v21_ = _dafny.Map({d_1330_v3_: len(d_1358_v20_)})
                index177_ = default__.safeIndex(529, (d_1330_v3_).length(0))
                (d_1330_v3_)[index177_] = ((d_1359_v21_)[d_1330_v3_] if (d_1330_v3_) in (d_1359_v21_) else (self).fm27(self.f29, d_1328_cf41_, 736, globalState))
                d_1360_v22_: _dafny.Map
                d_1360_v22_ = _dafny.Map({self.f35: self.f31})
                d_1361_v23_: _dafny.Seq
                d_1361_v23_ = _dafny.SeqWithoutIsStrInference([d_1327_cf42_])
                d_1360_v22_ = (d_1360_v22_).set(not((self.f31 if False else self.f31)), (d_1361_v23_) == (d_1361_v23_))
            elif True:
                d_1362___mcc_h17_ = source23_.cf7
                d_1363_cf7_ = d_1362___mcc_h17_
                d_1364_v24_: _dafny.Map
                d_1364_v24_ = _dafny.Map({d_1332_v5_: self.f36})
                d_1365_v25_: _dafny.Map
                d_1365_v25_ = _dafny.Map({p0: d_1364_v24_})
                d_1366_v26_: _dafny.Array
                nw226_ = _dafny.Array(None, 15)
                nw226_[int(0)] = d_1364_v24_
                nw226_[int(1)] = d_1364_v24_
                nw226_[int(2)] = d_1364_v24_
                nw226_[int(3)] = d_1364_v24_
                nw226_[int(4)] = (d_1364_v24_) | (d_1364_v24_)
                nw226_[int(5)] = (d_1364_v24_) | (d_1364_v24_)
                nw226_[int(6)] = d_1364_v24_
                nw226_[int(7)] = d_1364_v24_
                nw226_[int(8)] = ((d_1364_v24_).set(d_1332_v5_, self.f36)) | (_dafny.Map({d_1332_v5_: self.f36}))
                nw226_[int(9)] = d_1364_v24_
                nw226_[int(10)] = _dafny.Map({d_1332_v5_: self.f36})
                nw226_[int(11)] = ((d_1365_v25_)[self.f31] if (self.f31) in (d_1365_v25_) else d_1364_v24_)
                nw226_[int(12)] = (d_1364_v24_) | (d_1364_v24_)
                nw226_[int(13)] = d_1364_v24_
                nw226_[int(14)] = d_1364_v24_
                d_1366_v26_ = nw226_
                index178_ = default__.safeIndex(525, (d_1366_v26_).length(0))
                (d_1366_v26_)[index178_] = d_1364_v24_
                d_1367_v27_: D1
                d_1367_v27_ = D1_DC2(d_1332_v5_)
                d_1368_v28_: C1
                nw227_ = C1()
                nw227_.ctor__(self.f35, self.f29, False)
                d_1368_v28_ = nw227_
                d_1369_v29_: _dafny.Map
                d_1369_v29_ = _dafny.Map({d_1368_v28_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auuk")))})
                index179_ = default__.safeIndex(525, (d_1366_v26_).length(0))
                rhs178_ = d_1364_v24_
                rhs179_ = False
                rhs180_ = (0) - ((0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(D5_DC13(self.f28, _dafny.Map({True: d_1328_cf41_}), self.f36)).cf29 for d_1370_i1_ in range(default__.abs(781))])), ((self).f34) + (((d_1369_v29_)[d_1368_v28_] if (d_1368_v28_) in (d_1369_v29_) else (self).fm28((_dafny.MultiSet([self.f35])).cardinality, self.f31, globalState))))))
                rhs181_ = d_1327_cf42_
                rhs182_ = d_1367_v27_
                lhs138_ = d_1366_v26_
                lhs139_ = default__.safeIndex(525, (d_1366_v26_).length(0))
                lhs140_ = self
                lhs141_ = globalState
                lhs142_ = globalState
                lhs138_[lhs139_] = rhs178_
                lhs140_.f29 = rhs179_
                lhs141_.f13 = rhs180_
                lhs142_.f13 = rhs181_
                d_1367_v27_ = rhs182_
                d_1371_v30_: _dafny.Array
                def lambda48_(d_1372_i2_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhgvtevm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chiwr")))

                init28_ = lambda48_
                nw228_ = _dafny.Array(None, 15)
                for i0_28_ in range(nw228_.length(0)):
                    nw228_[i0_28_] = init28_(i0_28_)
                d_1371_v30_ = nw228_
                d_1373_v31_: _dafny.Seq
                d_1373_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yguw"))
                index180_ = default__.safeIndex(746, (d_1371_v30_).length(0))
                (d_1371_v30_)[index180_] = (d_1373_v31_ if p0 else d_1373_v31_)
                index181_ = default__.safeIndex(746, (d_1371_v30_).length(0))
                (d_1371_v30_)[index181_] = (d_1373_v31_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhyw")))
                d_1328_cf41_ = d_1328_cf41_
                d_1328_cf41_ = len((_dafny.SeqWithoutIsStrInference([self.f36 for d_1374_i3_ in range(default__.abs(554))])) + (d_1373_v31_))
            d_1375_v33_: _dafny.MultiSet
            d_1375_v33_ = _dafny.MultiSet([d_1328_cf41_])
            d_1376_v34_: D4
            d_1376_v34_ = D4_DC10(d_1375_v33_)
            d_1377_v35_: _dafny.Map
            d_1377_v35_ = _dafny.Map({(d_1376_v34_).cf22: True})
            d_1378_v36_: _dafny.Seq
            def iife86_():
                coll40_ = _dafny.Map()
                compr_40_: _dafny.MultiSet
                for compr_40_ in (d_1377_v35_).keys.Elements:
                    d_1379_v32_: _dafny.MultiSet = compr_40_
                    if (d_1379_v32_) in (d_1377_v35_):
                        coll40_[d_1379_v32_] = self.f29
                return _dafny.Map(coll40_)
            d_1378_v36_ = _dafny.SeqWithoutIsStrInference([len(iife86_()
            ), d_1328_cf41_])
            d_1380_v37_: _dafny.Set
            d_1380_v37_ = _dafny.Set({(self).f34})
            (globalState).f13 = len((d_1378_v36_).set(default__.safeIndex(len(d_1380_v37_), len(d_1378_v36_)), ((self).f30) + (d_1328_cf41_)))
            (globalState).f13 = (0) - ((self).f30)
        elif source22_.is_DC23:
            d_1381___mcc_h2_ = source22_.cf43
            d_1382___mcc_h3_ = source22_.cf44
            d_1383___mcc_h4_ = source22_.cf45
            d_1384_cf45_ = d_1383___mcc_h4_
            d_1385_cf44_ = d_1382___mcc_h3_
            d_1386_cf43_ = d_1381___mcc_h2_
            d_1387_v38_: _dafny.Array
            nw229_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_1387_v38_ = nw229_
            d_1388_v39_: C0
            nw230_ = C0()
            nw230_.ctor__(d_1387_v38_)
            d_1388_v39_ = nw230_
            d_1389_v40_: _dafny.Map
            d_1389_v40_ = _dafny.Map({(self).f34: d_1388_v39_})
            d_1390_v41_: _dafny.Seq
            d_1390_v41_ = _dafny.SeqWithoutIsStrInference([835])
            d_1391_v42_: _dafny.Seq
            d_1391_v42_ = _dafny.SeqWithoutIsStrInference([161, len(d_1390_v41_)])
            d_1392_v43_: _dafny.Set
            d_1392_v43_ = _dafny.Set({d_1391_v42_})
            d_1393_v44_: _dafny.Map
            d_1393_v44_ = _dafny.Map({len(d_1392_v43_): d_1385_cf44_})
            d_1394_v45_: _dafny.Set
            d_1394_v45_ = _dafny.Set({True})
            d_1395_v46_: _dafny.Map
            d_1395_v46_ = _dafny.Map({d_1393_v44_: d_1394_v45_})
            d_1396_v47_: _dafny.Seq
            d_1396_v47_ = _dafny.SeqWithoutIsStrInference([((d_1395_v46_)[d_1393_v44_] if (d_1393_v44_) in (d_1395_v46_) else d_1394_v45_), _dafny.Set({d_1385_cf44_, d_1384_cf45_}), d_1394_v45_])
            d_1397_v48_: _dafny.Array
            nw231_ = _dafny.Array(None, 28)
            nw231_[int(0)] = d_1388_v39_
            nw231_[int(1)] = d_1388_v39_
            nw231_[int(2)] = d_1388_v39_
            nw231_[int(3)] = d_1388_v39_
            nw231_[int(4)] = d_1388_v39_
            nw231_[int(5)] = d_1388_v39_
            nw231_[int(6)] = d_1388_v39_
            nw231_[int(7)] = d_1388_v39_
            nw231_[int(8)] = d_1388_v39_
            nw231_[int(9)] = d_1388_v39_
            nw231_[int(10)] = d_1388_v39_
            nw231_[int(11)] = d_1388_v39_
            nw231_[int(12)] = d_1388_v39_
            nw231_[int(13)] = d_1388_v39_
            nw231_[int(14)] = d_1388_v39_
            nw231_[int(15)] = d_1388_v39_
            nw231_[int(16)] = d_1388_v39_
            nw231_[int(17)] = d_1388_v39_
            nw231_[int(18)] = d_1388_v39_
            nw231_[int(19)] = ((d_1389_v40_)[len(d_1396_v47_)] if (len(d_1396_v47_)) in (d_1389_v40_) else d_1388_v39_)
            nw231_[int(20)] = d_1388_v39_
            nw231_[int(21)] = d_1388_v39_
            nw231_[int(22)] = d_1388_v39_
            nw231_[int(23)] = d_1388_v39_
            nw231_[int(24)] = d_1388_v39_
            nw231_[int(25)] = d_1388_v39_
            nw231_[int(26)] = d_1388_v39_
            nw231_[int(27)] = d_1388_v39_
            d_1397_v48_ = nw231_
            d_1397_v48_ = d_1397_v48_
            r2 = (len((self).f32)) - ((self).f30)
            (globalState).f13 = 262
            (globalState).f12 = d_1386_cf43_
        elif source22_.is_DC24:
            d_1398___mcc_h5_ = source22_.cf46
            d_1399___mcc_h6_ = source22_.cf47
            d_1400___mcc_h7_ = source22_.cf48
            d_1401___mcc_h8_ = source22_.cf49
            d_1402___mcc_h9_ = source22_.cf50
            d_1403_cf50_ = d_1402___mcc_h9_
            d_1404_cf49_ = d_1401___mcc_h8_
            d_1405_cf48_ = d_1400___mcc_h7_
            d_1406_cf47_ = d_1399___mcc_h6_
            d_1407_cf46_ = d_1398___mcc_h5_
            d_1408_v49_: _dafny.Seq
            d_1408_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_1409_v50_: _dafny.Map
            d_1409_v50_ = _dafny.Map({(0) - ((self).f34): self.f36})
            d_1410_v51_: _dafny.Array
            nw232_ = _dafny.Array(None, 16)
            nw232_[int(0)] = self.f36
            nw232_[int(1)] = self.f36
            nw232_[int(2)] = self.f36
            nw232_[int(3)] = self.f36
            nw232_[int(4)] = _dafny.CodePoint('l')
            nw232_[int(5)] = self.f36
            nw232_[int(6)] = self.f36
            nw232_[int(7)] = ((d_1409_v50_)[len(d_1408_v49_)] if (len(d_1408_v49_)) in (d_1409_v50_) else self.f36)
            nw232_[int(8)] = _dafny.CodePoint('c')
            nw232_[int(9)] = self.f36
            nw232_[int(10)] = self.f36
            nw232_[int(11)] = self.f36
            nw232_[int(12)] = self.f36
            nw232_[int(13)] = self.f36
            nw232_[int(14)] = self.f36
            nw232_[int(15)] = _dafny.CodePoint('f')
            d_1410_v51_ = nw232_
            d_1411_v52_: C0
            nw233_ = C0()
            nw233_.ctor__(d_1410_v51_)
            d_1411_v52_ = nw233_
            source24_ = D1_DC3(self.f35, (d_1403_cf50_)[default__.safeIndex(len(d_1408_v49_), len(d_1403_cf50_))], d_1411_v52_, default__.safeModuloInt(d_1406_cf47_, d_1406_cf47_))
            if source24_.is_DC3:
                d_1412___mcc_h18_ = source24_.cf3
                d_1413___mcc_h19_ = source24_.cf4
                d_1414___mcc_h20_ = source24_.cf5
                d_1415___mcc_h21_ = source24_.cf6
                d_1416_cf6_ = d_1415___mcc_h21_
                d_1417_cf5_ = d_1414___mcc_h20_
                d_1418_cf4_ = d_1413___mcc_h19_
                d_1419_cf3_ = d_1412___mcc_h18_
                d_1420_v53_: C0
                nw234_ = C0()
                nw234_.ctor__(d_1411_v52_.f37)
                d_1420_v53_ = nw234_
                d_1421_v54_: T0
                nw235_ = C1()
                nw235_.ctor__(d_1404_cf49_, False, True)
                d_1421_v54_ = nw235_
                d_1422_v55_: _dafny.Map
                d_1422_v55_ = _dafny.Map({d_1421_v54_: self.f36})
                d_1418_cf4_ = (default__.safeModuloInt(len(d_1422_v55_), d_1416_cf6_)) - (d_1418_cf4_)
                d_1423_v56_: _dafny.Array
                def lambda49_(d_1424_i4_):
                    return (d_1424_i4_) - ((self).f34)

                init29_ = lambda49_
                nw236_ = _dafny.Array(None, 29)
                for i0_29_ in range(nw236_.length(0)):
                    nw236_[i0_29_] = init29_(i0_29_)
                d_1423_v56_ = nw236_
                d_1425_v57_: _dafny.Map
                d_1425_v57_ = _dafny.Map({self.f29: self.f36})
                d_1426_v58_: _dafny.Map
                d_1426_v58_ = _dafny.Map({d_1423_v56_: len(d_1425_v57_)})
                d_1427_v59_: _dafny.MultiSet
                d_1427_v59_ = _dafny.MultiSet([d_1421_v54_.f28])
                d_1428_v60_: C2
                nw237_ = C2()
                nw237_.ctor__((self).f30, len(d_1426_v58_), (d_1427_v59_).isdisjoint(_dafny.MultiSet([d_1421_v54_.f28, self.f28])), not(not(d_1421_v54_.f28)), not(d_1421_v54_.f28))
                d_1428_v60_ = nw237_
                r3 = (self).fm27(((self).f32)[default__.safeIndex(d_1416_cf6_, len((self).f32))], (len(d_1403_cf50_)) * (d_1418_cf4_), d_1418_cf4_, globalState)
            elif source24_.is_DC2:
                d_1429___mcc_h22_ = source24_.cf2
                d_1430_cf2_ = d_1429___mcc_h22_
                (self).f36 = self.f36
                d_1431_v61_: C0
                nw238_ = C0()
                nw238_.ctor__(d_1411_v52_.f37)
                d_1431_v61_ = nw238_
                r1 = not((d_1405_cf48_)[default__.safeIndex((self).f30, len(d_1405_cf48_))])
                (globalState).f13 = (self).f30
            elif True:
                d_1432___mcc_h23_ = source24_.cf7
                d_1433_cf7_ = d_1432___mcc_h23_
                d_1434_v62_: C0
                nw239_ = C0()
                nw239_.ctor__(d_1411_v52_.f37)
                d_1434_v62_ = nw239_
                r2 = default__.safeModuloInt((self).f34, d_1406_cf47_)
                d_1435_v63_: _dafny.Map
                d_1435_v63_ = _dafny.Map({(self).f34: not(d_1407_cf46_)})
                d_1436_v64_: _dafny.Map
                d_1436_v64_ = _dafny.Map({(self).f34: len(d_1435_v63_)})
                d_1437_v65_: _dafny.Array
                nw240_ = _dafny.Array(None, 23)
                nw240_[int(0)] = d_1405_cf48_
                nw240_[int(1)] = (self).f32
                nw240_[int(2)] = d_1405_cf48_
                nw240_[int(3)] = (self).f32
                nw240_[int(4)] = d_1405_cf48_
                nw240_[int(5)] = (self).f32
                nw240_[int(6)] = d_1405_cf48_
                nw240_[int(7)] = default__.fm0(d_1406_cf47_, self.f35, p0, len(d_1436_v64_), globalState)
                nw240_[int(8)] = (self).f32
                nw240_[int(9)] = (self).f32
                nw240_[int(10)] = d_1405_cf48_
                nw240_[int(11)] = (self).f32
                nw240_[int(12)] = _dafny.SeqWithoutIsStrInference([self.f35, d_1404_cf49_, d_1407_cf46_, self.f31, self.f29])
                nw240_[int(13)] = (((self).f32).set(default__.safeIndex(d_1406_cf47_, len((self).f32)), p0)).set(default__.safeIndex((self).f34, len(((self).f32).set(default__.safeIndex(d_1406_cf47_, len((self).f32)), p0))), self.f35)
                nw240_[int(14)] = d_1405_cf48_
                nw240_[int(15)] = ((self).f32).set(default__.safeIndex(d_1406_cf47_, len((self).f32)), not(d_1407_cf46_))
                nw240_[int(16)] = (self).f32
                nw240_[int(17)] = _dafny.SeqWithoutIsStrInference([p0])
                nw240_[int(18)] = (self).f32
                nw240_[int(19)] = (self).f32
                nw240_[int(20)] = _dafny.SeqWithoutIsStrInference([self.f35])
                nw240_[int(21)] = d_1405_cf48_
                nw240_[int(22)] = _dafny.SeqWithoutIsStrInference([p0, d_1407_cf46_, True, self.f28, False])
                d_1437_v65_ = nw240_
                d_1438_v66_: _dafny.Set
                d_1438_v66_ = _dafny.Set({self.f31, d_1404_cf49_})
                d_1439_v67_: D2
                d_1439_v67_ = D2_DC6(d_1406_cf47_, False, (self).f30, d_1407_cf46_, d_1406_cf47_)
                d_1440_v68_: C3
                nw241_ = C3()
                nw241_.ctor__((self).f34, d_1437_v65_, self.f36, (self).f34, (default__.fm1(self.f35, len(default__.fm38(False, len(d_1438_v66_), globalState)), globalState)) and (d_1404_cf49_), ((((self).f33)[default__.safeIndex(d_1406_cf47_, len((self).f33))]).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f36])), len(((self).f33)[default__.safeIndex(d_1406_cf47_, len((self).f33))])), p0)).set(default__.safeIndex(540, len((((self).f33)[default__.safeIndex(d_1406_cf47_, len((self).f33))]).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f36])), len(((self).f33)[default__.safeIndex(d_1406_cf47_, len((self).f33))])), p0))), self.f35), (self).f33, (self).f34, self.f35, (d_1439_v67_).cf10, p0)
                d_1440_v68_ = nw241_
                r3 = (d_1406_cf47_) * ((d_1440_v68_).f44)
            if d_1407_cf46_:
                d_1441_v69_: _dafny.Map
                d_1441_v69_ = _dafny.Map({(self).f30: not(d_1404_cf49_)})
                d_1442_v70_: _dafny.Set
                d_1442_v70_ = _dafny.Set({d_1441_v69_})
                d_1443_v71_: D7
                d_1443_v71_ = D7_DC19(self.f29, (0) - (d_1406_cf47_), (self).f30)
                d_1442_v70_ = ((d_1442_v70_) - (d_1442_v70_) if (d_1443_v71_).cf36 else d_1442_v70_)
                d_1444_v72_: _dafny.Map
                d_1444_v72_ = _dafny.Map({d_1406_cf47_: _dafny.Set({(self).f34})})
                d_1445_v73_: _dafny.Set
                d_1445_v73_ = _dafny.Set({self.f35, self.f31})
                d_1446_v74_: _dafny.Array
                nw242_ = _dafny.Array(None, 20)
                nw242_[int(0)] = (self).f34
                nw242_[int(1)] = (self).f34
                nw242_[int(2)] = (0) - (d_1406_cf47_)
                nw242_[int(3)] = (self).f30
                nw242_[int(4)] = len(d_1444_v72_)
                nw242_[int(5)] = (self).f30
                nw242_[int(6)] = (self).f34
                nw242_[int(7)] = d_1406_cf47_
                nw242_[int(8)] = len(d_1445_v73_)
                nw242_[int(9)] = d_1406_cf47_
                nw242_[int(10)] = (self).f34
                nw242_[int(11)] = (0) - ((self).f30)
                nw242_[int(12)] = (self).f30
                nw242_[int(13)] = (self).f30
                nw242_[int(14)] = 300
                nw242_[int(15)] = d_1406_cf47_
                nw242_[int(16)] = (self).f30
                nw242_[int(17)] = (self).f34
                nw242_[int(18)] = (self).f34
                nw242_[int(19)] = 162
                d_1446_v74_ = nw242_
                d_1447_v75_: _dafny.Seq
                d_1447_v75_ = _dafny.SeqWithoutIsStrInference([d_1446_v74_])
                d_1448_v76_: _dafny.Map
                d_1448_v76_ = _dafny.Map({default__.safeDivisionInt((self).f34, (self).f30): (self).fm27(self.f31, len(d_1447_v75_), (self).fm28((self).f34, (self).fm7(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f34])), d_1406_cf47_, globalState), globalState), globalState)})
                d_1448_v76_ = (d_1448_v76_).set((0) - (d_1406_cf47_), default__.safeModuloInt(d_1406_cf47_, d_1406_cf47_))
                r3 = (default__.safeModuloInt(len(d_1408_v49_), (self).f34) if self.f35 else (636) - ((self).f30))
                d_1449_v77_: _dafny.Array
                nw243_ = _dafny.Array(None, 11)
                d_1449_v77_ = nw243_
                d_1450_v78_: _dafny.MultiSet
                d_1450_v78_ = _dafny.MultiSet([(0) - (d_1406_cf47_)])
                d_1451_v79_: C1
                nw244_ = C1()
                nw244_.ctor__(self.f31, self.f35, (self).fm7(d_1450_v78_, (self).f34, globalState))
                d_1451_v79_ = nw244_
                index182_ = default__.safeIndex(942, (d_1449_v77_).length(0))
                (d_1449_v77_)[index182_] = d_1451_v79_
                d_1452_v80_: _dafny.Array
                nw245_ = _dafny.Array(_dafny.Map({}), 2)
                d_1452_v80_ = nw245_
                d_1453_v81_: _dafny.Array
                def lambda50_(d_1454_i5_):
                    return self.f29

                init30_ = lambda50_
                nw246_ = _dafny.Array(None, 26)
                for i0_30_ in range(nw246_.length(0)):
                    nw246_[i0_30_] = init30_(i0_30_)
                d_1453_v81_ = nw246_
                d_1455_v82_: _dafny.Map
                d_1455_v82_ = _dafny.Map({((d_1448_v76_)[(self).f30] if ((self).f30) in (d_1448_v76_) else (self).f30): d_1453_v81_})
                index183_ = default__.safeIndex(257, (d_1452_v80_).length(0))
                (d_1452_v80_)[index183_] = d_1455_v82_
                d_1456_v83_: _dafny.Map
                d_1456_v83_ = _dafny.Map({False: d_1451_v79_})
                d_1457_v84_: _dafny.Map
                d_1457_v84_ = _dafny.Map({self.f29: d_1455_v82_})
                index184_ = default__.safeIndex(942, (d_1449_v77_).length(0))
                index185_ = default__.safeIndex(257, (d_1452_v80_).length(0))
                rhs183_ = default__.safeDivisionInt((self).f34, d_1406_cf47_)
                rhs184_ = ((d_1456_v83_)[self.f29] if (self.f29) in (d_1456_v83_) else (D10_DC28(d_1451_v79_)).cf55)
                rhs185_ = ((((d_1457_v84_)[(d_1451_v79_).f43] if ((d_1451_v79_).f43) in (d_1457_v84_) else d_1455_v82_)) | (d_1455_v82_)) | (d_1455_v82_)
                lhs143_ = globalState
                lhs144_ = d_1449_v77_
                lhs145_ = default__.safeIndex(942, (d_1449_v77_).length(0))
                lhs146_ = d_1452_v80_
                lhs147_ = default__.safeIndex(257, (d_1452_v80_).length(0))
                lhs143_.f13 = rhs183_
                lhs144_[lhs145_] = rhs184_
                lhs146_[lhs147_] = rhs185_
                d_1458_v85_: _dafny.Array
                nw247_ = _dafny.Array(None, 11)
                nw247_[int(0)] = d_1453_v81_
                nw247_[int(1)] = d_1453_v81_
                nw247_[int(2)] = d_1453_v81_
                nw247_[int(3)] = d_1453_v81_
                nw247_[int(4)] = d_1453_v81_
                nw247_[int(5)] = d_1453_v81_
                nw247_[int(6)] = d_1453_v81_
                nw247_[int(7)] = d_1453_v81_
                nw247_[int(8)] = d_1453_v81_
                nw247_[int(9)] = (d_1453_v81_ if self.f31 else d_1453_v81_)
                nw247_[int(10)] = d_1453_v81_
                d_1458_v85_ = nw247_
                d_1459_v86_: _dafny.Seq
                d_1459_v86_ = _dafny.SeqWithoutIsStrInference([d_1453_v81_, d_1453_v81_])
                index186_ = default__.safeIndex(233, (d_1458_v85_).length(0))
                (d_1458_v85_)[index186_] = (d_1459_v86_)[default__.safeIndex((0) - (len(d_1408_v49_)), len(d_1459_v86_))]
                index187_ = default__.safeIndex(233, (d_1458_v85_).length(0))
                (d_1458_v85_)[index187_] = d_1453_v81_
            elif True:
                (globalState).f1 = self.f28
                (globalState).f1 = default__.fm1(False, (self).fm28((self).f30, d_1404_cf49_, globalState), globalState)
                d_1460_v87_: _dafny.Map
                d_1460_v87_ = _dafny.Map({d_1407_cf46_: not(self.f31)})
                d_1460_v87_ = (d_1460_v87_).set(self.f29, ((self).f34) >= ((self).f34))
                d_1461_v88_: C0
                nw248_ = C0()
                nw248_.ctor__(d_1411_v52_.f37)
                d_1461_v88_ = nw248_
                (globalState).f26 = (self.f31) and (self.f31)
            d_1462_v89_: D10
            d_1462_v89_ = D10_DC29(d_1406_cf47_, default__.fm1(self.f29, d_1406_cf47_, globalState))
            (self).f28 = (d_1462_v89_).cf57
            (globalState).f1 = not (self.f35) or (d_1404_cf49_)
        elif source22_.is_DC21:
            d_1463___mcc_h10_ = source22_.cf40
            d_1464_cf40_ = d_1463___mcc_h10_
            d_1465_v90_: _dafny.Seq
            d_1465_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggytjvvm"))
            d_1466_v91_: _dafny.Array
            nw249_ = _dafny.Array(None, 29)
            nw249_[int(0)] = d_1465_v90_
            nw249_[int(1)] = d_1465_v90_
            nw249_[int(2)] = (d_1465_v90_ if False else d_1465_v90_)
            nw249_[int(3)] = d_1465_v90_
            nw249_[int(4)] = d_1465_v90_
            nw249_[int(5)] = d_1465_v90_
            nw249_[int(6)] = d_1465_v90_
            nw249_[int(7)] = (d_1465_v90_) + (d_1465_v90_)
            nw249_[int(8)] = d_1465_v90_
            nw249_[int(9)] = d_1465_v90_
            nw249_[int(10)] = (d_1465_v90_).set(default__.safeIndex((self).f30, len(d_1465_v90_)), self.f36)
            nw249_[int(11)] = d_1465_v90_
            nw249_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqoj"))
            nw249_[int(13)] = d_1465_v90_
            nw249_[int(14)] = d_1465_v90_
            nw249_[int(15)] = d_1465_v90_
            nw249_[int(16)] = d_1465_v90_
            nw249_[int(17)] = d_1465_v90_
            nw249_[int(18)] = d_1465_v90_
            nw249_[int(19)] = d_1465_v90_
            nw249_[int(20)] = (d_1465_v90_) + (d_1465_v90_)
            nw249_[int(21)] = _dafny.SeqWithoutIsStrInference([self.f36 for d_1467_i6_ in range(default__.abs(-167))])
            nw249_[int(22)] = d_1465_v90_
            nw249_[int(23)] = (self).fm3(self.f29, p0, globalState)
            nw249_[int(24)] = d_1465_v90_
            nw249_[int(25)] = d_1465_v90_
            nw249_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lskefdq"))
            nw249_[int(27)] = d_1465_v90_
            nw249_[int(28)] = ((d_1465_v90_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))).set(default__.safeIndex((0) - ((self).f34), len((d_1465_v90_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))))), self.f36)
            d_1466_v91_ = nw249_
            index188_ = default__.safeIndex(600, (d_1466_v91_).length(0))
            (d_1466_v91_)[index188_] = (d_1465_v90_) + (d_1465_v90_)
            index189_ = default__.safeIndex(600, (d_1466_v91_).length(0))
            rhs186_ = d_1465_v90_
            rhs187_ = 852
            lhs148_ = d_1466_v91_
            lhs149_ = default__.safeIndex(600, (d_1466_v91_).length(0))
            lhs150_ = globalState
            lhs148_[lhs149_] = rhs186_
            lhs150_.f13 = rhs187_
            if self.f35:
                (self).f36 = self.f36
                d_1468_v92_: _dafny.Array
                nw250_ = _dafny.Array(False, 19)
                d_1468_v92_ = nw250_
                d_1469_v93_: _dafny.Seq
                d_1469_v93_ = _dafny.SeqWithoutIsStrInference([(self).f30, (0) - ((self).f34)])
                d_1470_v94_: _dafny.Map
                d_1470_v94_ = _dafny.Map({d_1469_v93_: True})
                d_1471_v95_: _dafny.Array
                def lambda51_(d_1472_i7_):
                    return self.f36

                init31_ = lambda51_
                nw251_ = _dafny.Array(None, 23)
                for i0_31_ in range(nw251_.length(0)):
                    nw251_[i0_31_] = init31_(i0_31_)
                d_1471_v95_ = nw251_
                d_1473_v96_: C0
                nw252_ = C0()
                nw252_.ctor__(d_1471_v95_)
                d_1473_v96_ = nw252_
                d_1474_v97_: _dafny.Seq
                d_1474_v97_ = _dafny.SeqWithoutIsStrInference([d_1473_v96_])
                d_1475_v98_: D9
                d_1475_v98_ = D9_DC27((d_1474_v97_)[default__.safeIndex((self).f34, len(d_1474_v97_))], d_1465_v90_)
                d_1476_v99_: _dafny.MultiSet
                d_1476_v99_ = _dafny.MultiSet([d_1475_v98_])
                index190_ = default__.safeIndex(260, (d_1468_v92_).length(0))
                (d_1468_v92_)[index190_] = (len(d_1470_v94_)) > ((d_1476_v99_).cardinality)
                index191_ = default__.safeIndex(260, (d_1468_v92_).length(0))
                (d_1468_v92_)[index191_] = (d_1465_v90_) <= ((d_1466_v91_)[default__.safeIndex(600, (d_1466_v91_).length(0))])
                d_1477_v100_: _dafny.Array
                def lambda52_(d_1478_i8_):
                    return D8_DC21(_dafny.Map({D0_DC1(_dafny.MultiSet((self).f32)): self.f35}))

                init32_ = lambda52_
                nw253_ = _dafny.Array(None, 27)
                for i0_32_ in range(nw253_.length(0)):
                    nw253_[i0_32_] = init32_(i0_32_)
                d_1477_v100_ = nw253_
                d_1479_v101_: _dafny.Map
                d_1479_v101_ = _dafny.Map({(self).f30: self.f31})
                d_1480_v102_: _dafny.Map
                d_1480_v102_ = _dafny.Map({p0: (d_1468_v92_)[default__.safeIndex(260, (d_1468_v92_).length(0))]})
                d_1481_v103_: _dafny.Map
                d_1481_v103_ = _dafny.Map({len(d_1479_v101_): d_1480_v102_})
                d_1482_v104_: _dafny.Map
                d_1482_v104_ = _dafny.Map({d_1477_v100_: len((d_1481_v103_) | (d_1481_v103_))})
                d_1482_v104_ = d_1482_v104_
                d_1483_v105_: _dafny.MultiSet
                d_1483_v105_ = _dafny.MultiSet([(self).f30])
                d_1484_v106_: T2
                nw254_ = C5()
                nw254_.ctor__(self.f29, (self).f34, (self).f34, self.f35, default__.fm0((d_1483_v105_).cardinality, p0, self.f29, (self).f34, globalState), (D14_DC36(default__.fm51(p0, self.f36, globalState))).cf67, len(d_1479_v101_), self.f28, self.f29, p0)
                d_1484_v106_ = nw254_
                d_1485_v107_: _dafny.Map
                d_1485_v107_ = _dafny.Map({(self).f34: d_1484_v106_})
                d_1486_v108_: _dafny.Seq
                d_1486_v108_ = _dafny.SeqWithoutIsStrInference([d_1484_v106_.f29, True, d_1484_v106_.f28, self.f28, (d_1468_v92_)[default__.safeIndex(260, (d_1468_v92_).length(0))]])
                d_1487_v109_: _dafny.MultiSet
                d_1487_v109_ = _dafny.MultiSet([(d_1484_v106_).f32, d_1486_v108_])
                def iife87_():
                    coll41_ = _dafny.Set()
                    compr_41_: _dafny.Seq
                    for compr_41_ in (d_1487_v109_).Elements:
                        d_1488_v110_: _dafny.Seq = compr_41_
                        if (d_1488_v110_) in (d_1487_v109_):
                            coll41_ = coll41_.union(_dafny.Set([d_1488_v110_]))
                    return _dafny.Set(coll41_)
                rhs188_ = d_1485_v107_
                rhs189_ = ((d_1484_v106_).f32).set(default__.safeIndex((self).f34, len((d_1484_v106_).f32)), (self.f36) not in ((d_1466_v91_)[default__.safeIndex(600, (d_1466_v91_).length(0))]))
                rhs190_ = (default__.fm52(D8_DC21((d_1464_cf40_).set(default__.fm36((d_1484_v106_).f30, (d_1483_v105_).cardinality, self.f29, globalState), d_1484_v106_.f29)), globalState)).issubset((_dafny.Set({d_1486_v108_, (self).f32, (d_1484_v106_).f32})) - (iife87_()
                ))
                lhs151_ = d_1484_v106_
                d_1485_v107_ = rhs188_
                d_1486_v108_ = rhs189_
                lhs151_.f29 = rhs190_
                d_1480_v102_ = (d_1480_v102_).set(not(((d_1479_v101_)[(self).f34] if ((self).f34) in (d_1479_v101_) else not(self.f28))), p0)
            elif True:
                d_1489_v111_: _dafny.Array
                nw255_ = _dafny.Array(int(0), 29)
                d_1489_v111_ = nw255_
                index192_ = default__.safeIndex(286, (d_1489_v111_).length(0))
                (d_1489_v111_)[index192_] = (self).f30
                index193_ = default__.safeIndex(286, (d_1489_v111_).length(0))
                (d_1489_v111_)[index193_] = ((self).f34) - (default__.safeDivisionInt(len((self).f32), (self).f34))
                d_1490_v112_: _dafny.Set
                d_1490_v112_ = _dafny.Set({(self.f29 if p0 else self.f28)})
                d_1491_v113_: _dafny.MultiSet
                d_1491_v113_ = _dafny.MultiSet([self.f28])
                d_1492_v114_: _dafny.Seq
                d_1492_v114_ = _dafny.SeqWithoutIsStrInference([((d_1491_v113_)[self.f35] if (self.f35) in (d_1491_v113_) else (self).f34), (d_1489_v111_)[default__.safeIndex(286, (d_1489_v111_).length(0))]])
                d_1490_v112_ = _dafny.Set({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pk")))) >= (len(d_1492_v114_)), self.f28, False})
                d_1493_v115_: int
                out57_: int
                out57_ = default__.m0(d_1492_v114_, self.f29, default__.safeDivisionInt((0) - (-105), (self).f34), self.f35, globalState)
                d_1493_v115_ = out57_
                d_1494_v116_: _dafny.Map
                d_1494_v116_ = _dafny.Map({66: self.f28})
                d_1495_v117_: _dafny.MultiSet
                d_1495_v117_ = _dafny.MultiSet([(d_1489_v111_)[default__.safeIndex(286, (d_1489_v111_).length(0))]])
                d_1496_v118_: _dafny.Map
                d_1496_v118_ = _dafny.Map({(d_1489_v111_)[default__.safeIndex(286, (d_1489_v111_).length(0))]: (self).f30})
                d_1497_v119_: C4
                nw256_ = C4()
                nw256_.ctor__(len((d_1466_v91_)[default__.safeIndex(600, (d_1466_v91_).length(0))]), (True) and (((d_1494_v116_)[len((self).f32)] if (len((self).f32)) in (d_1494_v116_) else (self).fm7(d_1495_v117_, d_1493_v115_, globalState))), (self).f32, _dafny.SeqWithoutIsStrInference([(self).f32 for d_1498_i9_ in range(default__.abs(846))]), (d_1491_v113_).cardinality, self.f35, self.f29, (len((d_1466_v91_)[default__.safeIndex(600, (d_1466_v91_).length(0))])) >= (len(d_1496_v118_)))
                d_1497_v119_ = nw256_
                d_1499_v120_: _dafny.Array
                nw257_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                d_1499_v120_ = nw257_
                d_1500_v121_: C0
                nw258_ = C0()
                nw258_.ctor__(d_1499_v120_)
                d_1500_v121_ = nw258_
                index194_ = default__.safeIndex(600, (d_1466_v91_).length(0))
                rhs191_ = (d_1465_v90_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gobhvtjm")))
                rhs192_ = d_1497_v119_
                rhs193_ = (D1_DC3(((self).f32)[default__.safeIndex((d_1489_v111_)[default__.safeIndex(286, (d_1489_v111_).length(0))], len((self).f32))], 302, d_1500_v121_, len((self).f32))).cf6
                lhs152_ = d_1466_v91_
                lhs153_ = default__.safeIndex(600, (d_1466_v91_).length(0))
                lhs152_[lhs153_] = rhs191_
                d_1497_v119_ = rhs192_
                r3 = rhs193_
                index195_ = default__.safeIndex(286, (d_1489_v111_).length(0))
                (d_1489_v111_)[index195_] = len((d_1490_v112_).intersection((d_1490_v112_) - (d_1490_v112_)))
            d_1501_v122_: _dafny.Array
            nw259_ = _dafny.Array(int(0), 19)
            d_1501_v122_ = nw259_
            index196_ = default__.safeIndex(255, (d_1501_v122_).length(0))
            (d_1501_v122_)[index196_] = (self).f34
            index197_ = default__.safeIndex(255, (d_1501_v122_).length(0))
            rhs194_ = ((-486) * (len(d_1465_v90_))) - (default__.safeDivisionInt((self).f34, (self).f34))
            rhs195_ = self.f29
            rhs196_ = ((d_1465_v90_) + ((d_1466_v91_)[default__.safeIndex(600, (d_1466_v91_).length(0))])) + (d_1465_v90_)
            rhs197_ = default__.safeModuloInt((0) - ((self).f30), (0) - ((self).f30))
            lhs154_ = self
            lhs155_ = d_1501_v122_
            lhs156_ = default__.safeIndex(255, (d_1501_v122_).length(0))
            r2 = rhs194_
            lhs154_.f35 = rhs195_
            d_1465_v90_ = rhs196_
            lhs155_[lhs156_] = rhs197_
            (self).f36 = _dafny.CodePoint('a')
        elif True:
            d_1502___mcc_h11_ = source22_.cf51
            d_1503_cf51_ = d_1502___mcc_h11_
            (self).f35 = not (self.f31) or (p0)
            r1 = not (True) or (not(p0))
            (self).f28 = not(p0)
            (globalState).f26 = p0
        if ((self).f30) == ((self).f34):
            d_1504_v123_: _dafny.Seq
            d_1504_v123_ = _dafny.SeqWithoutIsStrInference([(self).f34, (self).f30, (self).f30])
            d_1505_v124_: int
            out58_: int
            out58_ = default__.m0(d_1504_v123_, not(p0), (self).f30, self.f35, globalState)
            d_1505_v124_ = out58_
            d_1506_v125_: _dafny.Seq
            d_1506_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsw"))
            d_1506_v125_ = d_1506_v125_
            d_1507_v126_: _dafny.Map
            d_1507_v126_ = _dafny.Map({(self).f34: (self).f34})
            d_1508_v127_: _dafny.Seq
            d_1508_v127_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f34 for d_1509_i10_ in range(default__.abs(274))])])
            d_1510_v128_: D14
            d_1510_v128_ = D14_DC36((self).f33)
            d_1511_v129_: _dafny.MultiSet
            d_1511_v129_ = _dafny.MultiSet([d_1505_v124_])
            d_1512_v130_: _dafny.Seq
            d_1512_v130_ = _dafny.SeqWithoutIsStrInference([d_1511_v129_, _dafny.MultiSet([(self).f30, (self).f34])])
            d_1513_v131_: C5
            nw260_ = C5()
            nw260_.ctor__(self.f31, ((d_1507_v126_)[len((self).f32)] if (len((self).f32)) in (d_1507_v126_) else default__.fm32(True, globalState)), (self).f30, (d_1504_v123_) < ((d_1508_v127_)[default__.safeIndex((self).f30, len(d_1508_v127_))]), (self).f32, (d_1510_v128_).cf67, default__.safeDivisionInt((self).f34, default__.fm46(((d_1512_v130_)[default__.safeIndex((self).f34, len(d_1512_v130_))]).cardinality, (self).f34, False, globalState)), default__.fm1(True, (d_1504_v123_)[default__.safeIndex(len(d_1504_v123_), len(d_1504_v123_))], globalState), ((self).f32)[default__.safeIndex(d_1505_v124_, len((self).f32))], (not(self.f28)) or (p0))
            d_1513_v131_ = nw260_
            d_1514_v132_: D4
            d_1514_v132_ = D4_DC11(len(d_1506_v125_), self.f29, (d_1513_v131_).f48)
            source25_ = d_1514_v132_
            if source25_.is_DC11:
                d_1515___mcc_h24_ = source25_.cf23
                d_1516___mcc_h25_ = source25_.cf24
                d_1517___mcc_h26_ = source25_.cf25
                d_1518_cf25_ = d_1517___mcc_h26_
                d_1519_cf24_ = d_1516___mcc_h25_
                d_1520_cf23_ = d_1515___mcc_h24_
                d_1521_v133_: _dafny.Array
                nw261_ = _dafny.Array(None, 10)
                nw261_[int(0)] = self.f36
                nw261_[int(1)] = self.f36
                nw261_[int(2)] = self.f36
                nw261_[int(3)] = self.f36
                nw261_[int(4)] = _dafny.CodePoint('q')
                nw261_[int(5)] = self.f36
                nw261_[int(6)] = self.f36
                nw261_[int(7)] = self.f36
                nw261_[int(8)] = self.f36
                nw261_[int(9)] = _dafny.CodePoint('q')
                d_1521_v133_ = nw261_
                d_1522_v134_: C0
                nw262_ = C0()
                nw262_.ctor__(d_1521_v133_)
                d_1522_v134_ = nw262_
                d_1523_v135_: _dafny.Map
                d_1523_v135_ = _dafny.Map({d_1506_v125_: d_1522_v134_})
                d_1523_v135_ = _dafny.Map({d_1506_v125_: (d_1522_v134_ if (d_1513_v131_).f47 else (D1_DC3(p0, (0) - ((self).f34), d_1522_v134_, d_1505_v124_)).cf5)})
                d_1524_v136_: _dafny.Map
                d_1524_v136_ = _dafny.Map({(d_1513_v131_).f47: True})
                d_1524_v136_ = (d_1524_v136_).set(self.f28, (d_1505_v124_) < (d_1520_cf23_))
                d_1506_v125_ = (d_1506_v125_) + (d_1506_v125_)
                (globalState).f1 = ((self).f32)[default__.safeIndex(d_1518_cf25_, len((self).f32))]
            elif True:
                d_1525___mcc_h27_ = source25_.cf22
                d_1526_cf22_ = d_1525___mcc_h27_
                d_1527_v137_: _dafny.Map
                d_1527_v137_ = _dafny.Map({self.f31: self.f31})
                d_1528_v139_: C4
                nw263_ = C4()
                def iife88_():
                    coll42_ = _dafny.Map()
                    compr_42_: int
                    for compr_42_ in _dafny.IntegerRange(910, 734):
                        d_1529_v138_: int = compr_42_
                        if ((910) <= (d_1529_v138_)) and ((d_1529_v138_) < (734)):
                            coll42_[default__.safeModuloInt(d_1529_v138_, (self).f34)] = 341
                    return _dafny.Map(coll42_)
                nw263_.ctor__(d_1505_v124_, True, _dafny.SeqWithoutIsStrInference([not(self.f29)]), (((self).f33).set(default__.safeIndex(d_1505_v124_, len((self).f33)), (self).f32)) + (_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32, (self).fm5((d_1513_v131_).f48, (0) - (len(d_1527_v137_)), self.f28, d_1505_v124_, globalState)])), (self).f34, ((self).f30) not in (iife88_()
                ), (d_1513_v131_).f47, self.f29)
                d_1528_v139_ = nw263_
                r1 = ((d_1513_v131_).f48) >= ((self).f34)
                d_1506_v125_ = ((d_1506_v125_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhye")))) + (d_1506_v125_)
                d_1530_v140_: _dafny.Array
                def lambda53_(d_1531_p0_):
                    def lambda54_(d_1532_i11_):
                        return d_1531_p0_

                    return lambda54_

                init33_ = lambda53_(p0)
                nw264_ = _dafny.Array(None, 24)
                for i0_33_ in range(nw264_.length(0)):
                    nw264_[i0_33_] = init33_(i0_33_)
                d_1530_v140_ = nw264_
                d_1533_v141_: _dafny.Map
                d_1533_v141_ = _dafny.Map({(self).f30: (d_1513_v131_).f47})
                d_1534_v142_: _dafny.MultiSet
                d_1534_v142_ = _dafny.MultiSet([self.f29])
                index198_ = default__.safeIndex(721, (d_1530_v140_).length(0))
                (d_1530_v140_)[index198_] = (len(d_1533_v141_)) <= ((d_1534_v142_).cardinality)
                index199_ = default__.safeIndex(721, (d_1530_v140_).length(0))
                (d_1530_v140_)[index199_] = (((self).f30) == (166)) == ((d_1513_v131_).f47)
            d_1535_v143_: _dafny.Array
            def lambda55_(d_1536_v125_):
                def lambda56_(d_1537_i12_):
                    return d_1536_v125_

                return lambda56_

            init34_ = lambda55_(d_1506_v125_)
            nw265_ = _dafny.Array(None, 28)
            for i0_34_ in range(nw265_.length(0)):
                nw265_[i0_34_] = init34_(i0_34_)
            d_1535_v143_ = nw265_
            rhs198_ = (self).f30
            rhs199_ = d_1505_v124_
            rhs200_ = d_1535_v143_
            rhs201_ = (self).f34
            rhs202_ = ((d_1513_v131_).fm3(self.f29, (d_1513_v131_).f47, globalState)) + (d_1506_v125_)
            lhs157_ = globalState
            r3 = rhs198_
            lhs157_.f13 = rhs199_
            d_1535_v143_ = rhs200_
            r2 = rhs201_
            d_1506_v125_ = rhs202_
        elif True:
            d_1538_v144_: _dafny.Seq
            d_1538_v144_ = _dafny.SeqWithoutIsStrInference([(self).f34])
            d_1539_v145_: int
            out59_: int
            out59_ = default__.m0((d_1538_v144_) + (_dafny.SeqWithoutIsStrInference([(self).f30])), self.f29, (self).f34, p0, globalState)
            d_1539_v145_ = out59_
            r2 = d_1539_v145_
            d_1540_v146_: _dafny.MultiSet
            d_1540_v146_ = _dafny.MultiSet([(self).f30, d_1539_v145_, (self).f30, d_1539_v145_])
            d_1541_v147_: D4
            d_1541_v147_ = D4_DC10(d_1540_v146_)
            d_1542_v148_: _dafny.Seq
            d_1542_v148_ = _dafny.SeqWithoutIsStrInference([d_1540_v146_, (d_1541_v147_).cf22, (d_1540_v146_) - (d_1540_v146_)])
            d_1543_v149_: _dafny.Array
            def lambda57_(d_1544_i13_):
                return not(self.f31)

            init35_ = lambda57_
            nw266_ = _dafny.Array(None, 22)
            for i0_35_ in range(nw266_.length(0)):
                nw266_[i0_35_] = init35_(i0_35_)
            d_1543_v149_ = nw266_
            rhs203_ = d_1543_v149_
            rhs204_ = 354
            rhs205_ = default__.fm32(not (p0) or (p0), globalState)
            rhs206_ = ((d_1542_v148_) + (d_1542_v148_)) + (d_1542_v148_)
            lhs158_ = globalState
            lhs159_ = globalState
            lhs158_.f21 = rhs203_
            lhs159_.f13 = rhs204_
            d_1539_v145_ = rhs205_
            d_1542_v148_ = rhs206_
            d_1539_v145_ = d_1539_v145_
            r3 = default__.safeModuloInt((0) - ((_dafny.MultiSet([self.f29])).cardinality), ((self).f34) * ((self).f34))
        d_1545_v150_: _dafny.Array
        def lambda58_(d_1546_i14_):
            return self.f36

        init36_ = lambda58_
        nw267_ = _dafny.Array(None, 2)
        for i0_36_ in range(nw267_.length(0)):
            nw267_[i0_36_] = init36_(i0_36_)
        d_1545_v150_ = nw267_
        d_1547_v151_: C0
        nw268_ = C0()
        nw268_.ctor__(d_1545_v150_)
        d_1547_v151_ = nw268_
        d_1548_v152_: _dafny.Map
        d_1548_v152_ = _dafny.Map({p0: self.f35})
        d_1549_v153_: _dafny.Seq
        d_1549_v153_ = _dafny.SeqWithoutIsStrInference([d_1548_v152_, d_1548_v152_, d_1548_v152_, d_1548_v152_, d_1548_v152_])
        d_1550_v154_: D1
        d_1550_v154_ = D1_DC3(self.f29, (self).f30, d_1547_v151_, ((self).f30) + (len(d_1549_v153_)))
        source26_ = d_1550_v154_
        if source26_.is_DC3:
            d_1551___mcc_h28_ = source26_.cf3
            d_1552___mcc_h29_ = source26_.cf4
            d_1553___mcc_h30_ = source26_.cf5
            d_1554___mcc_h31_ = source26_.cf6
            d_1555_cf6_ = d_1554___mcc_h31_
            d_1556_cf5_ = d_1553___mcc_h30_
            d_1557_cf4_ = d_1552___mcc_h29_
            d_1558_cf3_ = d_1551___mcc_h28_
            d_1559_v155_: _dafny.Array
            nw269_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_1559_v155_ = nw269_
            d_1560_v156_: _dafny.Seq
            d_1560_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiirlaq"))
            index200_ = default__.safeIndex(860, (d_1559_v155_).length(0))
            (d_1559_v155_)[index200_] = (d_1560_v156_) + (d_1560_v156_)
            d_1561_v157_: _dafny.Seq
            d_1561_v157_ = _dafny.SeqWithoutIsStrInference([(self).f30])
            index201_ = default__.safeIndex(860, (d_1559_v155_).length(0))
            rhs207_ = ((self).f32)[default__.safeIndex((d_1561_v157_)[default__.safeIndex((self).f34, len(d_1561_v157_))], len((self).f32))]
            rhs208_ = _dafny.SeqWithoutIsStrInference([self.f36 for d_1562_i15_ in range(default__.abs(910))])
            rhs209_ = self.f28
            lhs160_ = d_1559_v155_
            lhs161_ = default__.safeIndex(860, (d_1559_v155_).length(0))
            lhs162_ = self
            r0 = rhs207_
            lhs160_[lhs161_] = rhs208_
            lhs162_.f29 = rhs209_
            d_1563_v158_: _dafny.Set
            d_1563_v158_ = _dafny.Set({((self).f30) == (d_1557_cf4_), True})
            d_1564_v159_: _dafny.Set
            d_1564_v159_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((self).f30))])})
            rhs210_ = ((0) - (d_1557_cf4_)) * (len((_dafny.Set({d_1561_v157_})).intersection(d_1564_v159_)))
            rhs211_ = not(((self).f30) != (((self).f34) + (d_1555_cf6_)))
            rhs212_ = self.f36
            rhs213_ = (d_1563_v158_) - (d_1563_v158_)
            rhs214_ = (len(d_1563_v158_)) >= ((self).f34)
            lhs163_ = globalState
            lhs164_ = globalState
            lhs165_ = self
            lhs166_ = self
            lhs163_.f13 = rhs210_
            lhs164_.f26 = rhs211_
            lhs165_.f36 = rhs212_
            d_1563_v158_ = rhs213_
            lhs166_.f29 = rhs214_
            d_1565_v160_: C2
            nw270_ = C2()
            nw270_.ctor__(d_1555_cf6_, default__.fm32(not(p0), globalState), not(self.f35), ((self).f32)[default__.safeIndex(d_1557_cf4_, len((self).f32))], self.f31)
            d_1565_v160_ = nw270_
            d_1566_v161_: D9
            d_1566_v161_ = D9_DC26(d_1565_v160_)
            d_1567_v162_: D9
            d_1567_v162_ = D9_DC26((d_1566_v161_).cf52)
            source27_ = d_1566_v161_
            if source27_.is_DC27:
                d_1568___mcc_h34_ = source27_.cf53
                d_1569___mcc_h35_ = source27_.cf54
                d_1570_cf54_ = d_1569___mcc_h35_
                d_1571_cf53_ = d_1568___mcc_h34_
                d_1572_v163_: T3
                nw271_ = C4()
                nw271_.ctor__((d_1565_v160_).f46, d_1558_cf3_, (self).f32, (self).f33, len(d_1548_v152_), self.f35, self.f35, self.f29)
                d_1572_v163_ = nw271_
                d_1573_v164_: _dafny.MultiSet
                d_1573_v164_ = _dafny.MultiSet([d_1572_v163_, d_1572_v163_, d_1572_v163_, d_1572_v163_, d_1572_v163_])
                d_1574_v165_: D10
                d_1574_v165_ = D10_DC30(self.f36, (d_1573_v164_).set(d_1572_v163_, default__.abs((d_1565_v160_).f46)))
                d_1574_v165_ = d_1574_v165_
                d_1575_v166_: _dafny.Array
                nw272_ = _dafny.Array(False, 8)
                d_1575_v166_ = nw272_
                index202_ = default__.safeIndex(323, (d_1575_v166_).length(0))
                (d_1575_v166_)[index202_] = d_1572_v163_.f31
                index203_ = default__.safeIndex(323, (d_1575_v166_).length(0))
                (d_1575_v166_)[index203_] = p0
                d_1576_v167_: _dafny.Map
                d_1576_v167_ = _dafny.Map({556: d_1572_v163_.f28})
                d_1576_v167_ = d_1576_v167_
                d_1577_v168_: _dafny.Map
                d_1577_v168_ = _dafny.Map({((0) - (-700)) <= ((d_1561_v157_)[default__.safeIndex(len((d_1570_cf54_).set(default__.safeIndex((d_1565_v160_).f46, len(d_1570_cf54_)), self.f36)), len(d_1561_v157_))]): (_dafny.MultiSet([not(d_1572_v163_.f35), False, (d_1575_v166_)[default__.safeIndex(323, (d_1575_v166_).length(0))]])).intersection(_dafny.MultiSet([d_1572_v163_.f31]))})
                d_1578_v169_: _dafny.MultiSet
                d_1578_v169_ = _dafny.MultiSet([False])
                d_1579_v170_: _dafny.Map
                d_1579_v170_ = _dafny.Map({(d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))]: d_1555_cf6_})
                d_1580_v171_: D9
                d_1580_v171_ = D9_DC27(d_1547_v151_, (d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))])
                d_1581_v172_: _dafny.Seq
                d_1581_v172_ = _dafny.SeqWithoutIsStrInference([d_1580_v171_])
                rhs215_ = ((d_1565_v160_).f46) * (((d_1565_v160_).fm29(globalState)) - (564))
                rhs216_ = ((d_1578_v169_).set(d_1572_v163_.f29, default__.abs(len(d_1579_v170_)))).ispropersubset(_dafny.MultiSet([d_1572_v163_.f28, self.f28, (d_1575_v166_)[default__.safeIndex(323, (d_1575_v166_).length(0))]]))
                rhs217_ = d_1558_cf3_
                rhs218_ = ((d_1577_v168_) | (d_1577_v168_)).set(not (d_1572_v163_.f28) or (d_1572_v163_.f29), _dafny.MultiSet((self).f32))
                rhs219_ = len(((self).fm3((d_1581_v172_) <= (d_1581_v172_), d_1558_cf3_, globalState)).set(default__.safeIndex((d_1565_v160_).f46, len((self).fm3((d_1581_v172_) <= (d_1581_v172_), d_1558_cf3_, globalState))), self.f36))
                lhs167_ = globalState
                lhs168_ = globalState
                d_1555_cf6_ = rhs215_
                lhs167_.f26 = rhs216_
                lhs168_.f1 = rhs217_
                d_1577_v168_ = rhs218_
                d_1555_cf6_ = rhs219_
            elif True:
                d_1582___mcc_h36_ = source27_.cf52
                d_1583_cf52_ = d_1582___mcc_h36_
                d_1560_v156_ = (d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))]
                (globalState).f12 = True
                d_1584_v173_: _dafny.MultiSet
                d_1584_v173_ = _dafny.MultiSet([d_1560_v156_])
                d_1585_v174_: _dafny.Map
                d_1585_v174_ = _dafny.Map({(d_1584_v173_) | (_dafny.MultiSet([(d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))], (d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))], (d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))], (d_1559_v155_)[default__.safeIndex(860, (d_1559_v155_).length(0))]])): (d_1565_v160_).f46})
                d_1585_v174_ = (d_1585_v174_).set(d_1584_v173_, d_1555_cf6_)
                d_1586_v175_: _dafny.MultiSet
                d_1586_v175_ = _dafny.MultiSet([self.f28, self.f35])
                d_1586_v175_ = (d_1586_v175_) - (d_1586_v175_)
            d_1587_v176_: _dafny.Seq
            d_1587_v176_ = _dafny.SeqWithoutIsStrInference([d_1560_v156_])
            d_1587_v176_ = (_dafny.SeqWithoutIsStrInference([d_1560_v156_ for d_1588_i16_ in range(default__.abs(-677))])) + (_dafny.SeqWithoutIsStrInference([d_1560_v156_]))
        elif source26_.is_DC2:
            d_1589___mcc_h32_ = source26_.cf2
            d_1590_cf2_ = d_1589___mcc_h32_
            d_1591_v177_: _dafny.Set
            d_1591_v177_ = _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([self.f36 for d_1592_i17_ in range(default__.abs(108))]))), (self).f34, (self).f30, default__.fm32(not(self.f35), globalState)})
            d_1591_v177_ = (d_1591_v177_).intersection(d_1591_v177_)
            (globalState).f1 = self.f35
            (globalState).f13 = 804
            d_1593_v178_: _dafny.Map
            d_1593_v178_ = _dafny.Map({(self).f34: (self).f30})
            d_1593_v178_ = d_1593_v178_
        elif True:
            d_1594___mcc_h33_ = source26_.cf7
            d_1595_cf7_ = d_1594___mcc_h33_
            (self).f31 = self.f29
            d_1596_v179_: _dafny.MultiSet
            d_1596_v179_ = _dafny.MultiSet([True])
            d_1597_v180_: D0
            d_1597_v180_ = D0_DC1(d_1596_v179_)
            d_1598_v181_: _dafny.Map
            d_1598_v181_ = _dafny.Map({d_1597_v180_: self.f29})
            d_1599_v182_: D8
            d_1599_v182_ = D8_DC21((d_1598_v181_) | (d_1598_v181_))
            source28_ = d_1599_v182_
            if source28_.is_DC22:
                d_1600___mcc_h37_ = source28_.cf41
                d_1601___mcc_h38_ = source28_.cf42
                d_1602_cf42_ = d_1601___mcc_h38_
                d_1603_cf41_ = d_1600___mcc_h37_
                d_1604_v183_: _dafny.Array
                def lambda59_(d_1605_cf41_):
                    def lambda60_(d_1606_i18_):
                        return default__.safeModuloInt(d_1606_i18_, d_1605_cf41_)

                    return lambda60_

                init37_ = lambda59_(d_1603_cf41_)
                nw273_ = _dafny.Array(None, 8)
                for i0_37_ in range(nw273_.length(0)):
                    nw273_[i0_37_] = init37_(i0_37_)
                d_1604_v183_ = nw273_
                index204_ = default__.safeIndex(571, (d_1604_v183_).length(0))
                (d_1604_v183_)[index204_] = 593
                index205_ = default__.safeIndex(571, (d_1604_v183_).length(0))
                (d_1604_v183_)[index205_] = (d_1602_cf42_) + ((d_1602_cf42_) * ((self).f34))
                d_1607_v184_: _dafny.Array
                nw274_ = _dafny.Array(False, 25)
                d_1607_v184_ = nw274_
                index206_ = default__.safeIndex(803, (d_1607_v184_).length(0))
                (d_1607_v184_)[index206_] = p0
                d_1608_v185_: _dafny.Set
                d_1608_v185_ = _dafny.Set({self.f28})
                d_1609_v186_: D7
                d_1609_v186_ = D7_DC19(self.f31, 530, 368)
                index207_ = default__.safeIndex(803, (d_1607_v184_).length(0))
                index208_ = default__.safeIndex(571, (d_1604_v183_).length(0))
                rhs220_ = (d_1608_v185_).ispropersubset((d_1608_v185_ if p0 else d_1608_v185_))
                rhs221_ = self.f31
                rhs222_ = (0) - ((self).f34)
                rhs223_ = not((d_1609_v186_).cf36)
                lhs169_ = d_1607_v184_
                lhs170_ = default__.safeIndex(803, (d_1607_v184_).length(0))
                lhs171_ = self
                lhs172_ = d_1604_v183_
                lhs173_ = default__.safeIndex(571, (d_1604_v183_).length(0))
                lhs174_ = globalState
                lhs169_[lhs170_] = rhs220_
                lhs171_.f28 = rhs221_
                lhs172_[lhs173_] = rhs222_
                lhs174_.f26 = rhs223_
                d_1610_v187_: _dafny.Seq
                d_1610_v187_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f28: (self).f34})])
                d_1611_v188_: _dafny.Map
                d_1611_v188_ = _dafny.Map({(d_1607_v184_)[default__.safeIndex(803, (d_1607_v184_).length(0))]: d_1603_cf41_})
                (globalState).f13 = len(((d_1610_v187_)[default__.safeIndex(-962, len(d_1610_v187_))]) | ((d_1611_v188_) | (d_1611_v188_)))
                d_1612_v189_: _dafny.Seq
                d_1612_v189_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwd"))
                rhs224_ = 742
                rhs225_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uy"))) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_1613_i19_ in range(default__.abs(863))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvqi")))
                rhs226_ = (d_1607_v184_)[default__.safeIndex(803, (d_1607_v184_).length(0))]
                lhs175_ = globalState
                r3 = rhs224_
                d_1612_v189_ = rhs225_
                lhs175_.f1 = rhs226_
            elif source28_.is_DC23:
                d_1614___mcc_h39_ = source28_.cf43
                d_1615___mcc_h40_ = source28_.cf44
                d_1616___mcc_h41_ = source28_.cf45
                d_1617_cf45_ = d_1616___mcc_h41_
                d_1618_cf44_ = d_1615___mcc_h40_
                d_1619_cf43_ = d_1614___mcc_h39_
                d_1620_v190_: _dafny.Seq
                d_1620_v190_ = _dafny.SeqWithoutIsStrInference([(self).f34])
                d_1621_v191_: C4
                nw275_ = C4()
                nw275_.ctor__((d_1596_v179_).cardinality, self.f28, ((self).f32) + (((self).f33)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f36 for d_1622_i20_ in range(default__.abs(229))])), len((self).f33))]), (self).f33, (self).f30, False, d_1618_cf44_, (d_1620_v190_) < ((d_1620_v190_).set(default__.safeIndex(526, len(d_1620_v190_)), (self).f30)))
                d_1621_v191_ = nw275_
                d_1623_v192_: C1
                nw276_ = C1()
                nw276_.ctor__(False, self.f31, d_1619_cf43_)
                d_1623_v192_ = nw276_
                d_1623_v192_ = d_1623_v192_
                d_1624_v193_: _dafny.Map
                d_1624_v193_ = _dafny.Map({self.f36: d_1618_cf44_})
                d_1625_v194_: _dafny.Map
                d_1625_v194_ = _dafny.Map({d_1624_v193_: d_1619_cf43_})
                d_1626_v195_: _dafny.Map
                d_1626_v195_ = _dafny.Map({d_1625_v194_: self.f28})
                d_1627_v196_: _dafny.Seq
                d_1627_v196_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvyflppiw"))
                d_1628_v197_: _dafny.MultiSet
                d_1628_v197_ = _dafny.MultiSet([(self).f30])
                d_1629_v198_: _dafny.Map
                d_1629_v198_ = _dafny.Map({(self).f34: self.f35})
                d_1630_v199_: _dafny.Array
                nw277_ = _dafny.Array(None, 28)
                nw277_[int(0)] = self.f31
                nw277_[int(1)] = ((d_1626_v195_)[d_1625_v194_] if (d_1625_v194_) in (d_1626_v195_) else self.f29)
                nw277_[int(2)] = self.f29
                nw277_[int(3)] = (d_1627_v196_) <= (d_1627_v196_)
                nw277_[int(4)] = (d_1623_v192_).f43
                nw277_[int(5)] = d_1617_cf45_
                nw277_[int(6)] = not ((d_1623_v192_).f43) or (p0)
                nw277_[int(7)] = d_1618_cf44_
                nw277_[int(8)] = d_1619_cf43_
                nw277_[int(9)] = self.f35
                nw277_[int(10)] = self.f28
                nw277_[int(11)] = (d_1623_v192_).f43
                nw277_[int(12)] = self.f31
                nw277_[int(13)] = (d_1623_v192_).f43
                nw277_[int(14)] = (d_1617_cf45_) in ((self).f32)
                nw277_[int(15)] = (d_1619_cf43_) or (((d_1548_v152_)[self.f28] if (self.f28) in (d_1548_v152_) else (self).fm7(d_1628_v197_, (self).f34, globalState)))
                nw277_[int(16)] = (d_1623_v192_).f43
                nw277_[int(17)] = d_1619_cf43_
                nw277_[int(18)] = (d_1623_v192_).f43
                nw277_[int(19)] = True
                nw277_[int(20)] = (d_1623_v192_).f43
                nw277_[int(21)] = self.f29
                nw277_[int(22)] = d_1619_cf43_
                nw277_[int(23)] = not((d_1623_v192_).f43)
                nw277_[int(24)] = self.f29
                nw277_[int(25)] = ((d_1629_v198_)[(self).f34] if ((self).f34) in (d_1629_v198_) else self.f28)
                nw277_[int(26)] = self.f35
                nw277_[int(27)] = True
                d_1630_v199_ = nw277_
                index209_ = default__.safeIndex(577, (d_1630_v199_).length(0))
                (d_1630_v199_)[index209_] = self.f28
                index210_ = default__.safeIndex(577, (d_1630_v199_).length(0))
                rhs227_ = ((d_1628_v197_)[(self).f34] if ((self).f34) in (d_1628_v197_) else (0) - ((self).f34))
                rhs228_ = d_1619_cf43_
                lhs176_ = globalState
                lhs177_ = d_1630_v199_
                lhs178_ = default__.safeIndex(577, (d_1630_v199_).length(0))
                lhs176_.f13 = rhs227_
                lhs177_[lhs178_] = rhs228_
                index211_ = default__.safeIndex(577, (d_1630_v199_).length(0))
                (d_1630_v199_)[index211_] = (d_1630_v199_)[default__.safeIndex(577, (d_1630_v199_).length(0))]
            elif source28_.is_DC24:
                d_1631___mcc_h42_ = source28_.cf46
                d_1632___mcc_h43_ = source28_.cf47
                d_1633___mcc_h44_ = source28_.cf48
                d_1634___mcc_h45_ = source28_.cf49
                d_1635___mcc_h46_ = source28_.cf50
                d_1636_cf50_ = d_1635___mcc_h46_
                d_1637_cf49_ = d_1634___mcc_h45_
                d_1638_cf48_ = d_1633___mcc_h44_
                d_1639_cf47_ = d_1632___mcc_h43_
                d_1640_cf46_ = d_1631___mcc_h42_
                d_1641_v200_: _dafny.Array
                nw278_ = _dafny.Array(None, 2)
                nw278_[int(0)] = d_1637_cf49_
                nw278_[int(1)] = ((self).f30) <= (d_1639_cf47_)
                d_1641_v200_ = nw278_
                index212_ = default__.safeIndex(703, (d_1641_v200_).length(0))
                (d_1641_v200_)[index212_] = self.f28
                index213_ = default__.safeIndex(703, (d_1641_v200_).length(0))
                (d_1641_v200_)[index213_] = not(self.f35)
                d_1642_v201_: _dafny.Array
                def lambda61_(d_1643_i21_):
                    return default__.safeModuloInt(d_1643_i21_, (self).f30)

                init38_ = lambda61_
                nw279_ = _dafny.Array(None, 1)
                for i0_38_ in range(nw279_.length(0)):
                    nw279_[i0_38_] = init38_(i0_38_)
                d_1642_v201_ = nw279_
                index214_ = default__.safeIndex(703, (d_1641_v200_).length(0))
                rhs229_ = d_1637_cf49_
                rhs230_ = (default__.safeModuloInt(d_1639_cf47_, 22)) - (((d_1596_v179_).cardinality) + (d_1639_cf47_))
                rhs231_ = d_1642_v201_
                rhs232_ = (645) <= (d_1639_cf47_)
                lhs179_ = d_1641_v200_
                lhs180_ = default__.safeIndex(703, (d_1641_v200_).length(0))
                d_1640_cf46_ = rhs229_
                r3 = rhs230_
                d_1642_v201_ = rhs231_
                lhs179_[lhs180_] = rhs232_
                d_1644_v202_: T3
                nw280_ = C4()
                nw280_.ctor__((self).f30, self.f35, _dafny.SeqWithoutIsStrInference([self.f31, p0, self.f35, d_1637_cf49_, d_1640_cf46_]), (self).f33, (self).f34, self.f35, self.f31, (d_1641_v200_)[default__.safeIndex(703, (d_1641_v200_).length(0))])
                d_1644_v202_ = nw280_
                d_1645_v203_: _dafny.MultiSet
                d_1645_v203_ = _dafny.MultiSet([d_1644_v202_, d_1644_v202_])
                d_1646_v204_: _dafny.Map
                d_1646_v204_ = _dafny.Map({D10_DC30(self.f36, d_1645_v203_): ((d_1644_v202_).f34) - ((0) - (d_1639_cf47_))})
                d_1646_v204_ = (d_1646_v204_).set(D10_DC30(self.f36, _dafny.MultiSet([d_1644_v202_])), d_1639_cf47_)
                d_1647_v205_: _dafny.Seq
                d_1647_v205_ = _dafny.SeqWithoutIsStrInference([d_1547_v151_.f37, d_1547_v151_.f37, d_1547_v151_.f37])
                d_1636_cf50_ = _dafny.SeqWithoutIsStrInference([len(d_1647_v205_)])
            elif source28_.is_DC21:
                d_1648___mcc_h47_ = source28_.cf40
                d_1649_cf40_ = d_1648___mcc_h47_
                d_1650_v206_: _dafny.Seq
                d_1650_v206_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkeeiw"))
                (globalState).f12 = (self.f36) in (d_1650_v206_)
                d_1651_v207_: _dafny.MultiSet
                d_1651_v207_ = _dafny.MultiSet([(self).f30])
                d_1652_v208_: _dafny.Set
                d_1652_v208_ = _dafny.Set({len(d_1548_v152_), (self).f34, (self).f30})
                d_1653_v209_: _dafny.Seq
                d_1653_v209_ = _dafny.SeqWithoutIsStrInference([(self).f34, (d_1596_v179_).cardinality, (self).f30, (d_1596_v179_).cardinality, (self).f30])
                d_1654_v210_: _dafny.Set
                d_1654_v210_ = _dafny.Set({not(not(self.f28)), self.f28, False})
                d_1655_v211_: _dafny.Array
                nw281_ = _dafny.Array(None, 8)
                nw281_[int(0)] = ((self).f34) in (d_1651_v207_)
                nw281_[int(1)] = (default__.fm42(d_1650_v206_, ((d_1651_v207_)[len(d_1652_v208_)] if (len(d_1652_v208_)) in (d_1651_v207_) else (self).f30), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvyovs")), globalState)) == (d_1653_v209_)
                nw281_[int(2)] = self.f31
                nw281_[int(3)] = (d_1654_v210_).isdisjoint(d_1654_v210_)
                nw281_[int(4)] = self.f31
                nw281_[int(5)] = self.f29
                nw281_[int(6)] = p0
                nw281_[int(7)] = self.f28
                d_1655_v211_ = nw281_
                index215_ = default__.safeIndex(917, (d_1655_v211_).length(0))
                (d_1655_v211_)[index215_] = ((self).f30) > ((default__.fm53(_dafny.CodePoint('v'), (self).f30, globalState)).cf62)
                index216_ = default__.safeIndex(917, (d_1655_v211_).length(0))
                (d_1655_v211_)[index216_] = default__.fm1(p0, (self).f30, globalState)
                d_1656_v212_: _dafny.Array
                nw282_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1656_v212_ = nw282_
                d_1657_v213_: _dafny.Array
                nw283_ = _dafny.Array(int(0), 3)
                d_1657_v213_ = nw283_
                index217_ = default__.safeIndex(170, (d_1656_v212_).length(0))
                (d_1656_v212_)[index217_] = d_1657_v213_
                index218_ = default__.safeIndex(170, (d_1656_v212_).length(0))
                rhs233_ = d_1657_v213_
                rhs234_ = (self).f30
                rhs235_ = (self).f34
                rhs236_ = (self).fm6(globalState)
                lhs181_ = d_1656_v212_
                lhs182_ = default__.safeIndex(170, (d_1656_v212_).length(0))
                lhs183_ = globalState
                lhs184_ = globalState
                lhs185_ = self
                lhs181_[lhs182_] = rhs233_
                lhs183_.f13 = rhs234_
                lhs184_.f13 = rhs235_
                lhs185_.f35 = rhs236_
                d_1658_v214_: _dafny.Map
                d_1658_v214_ = _dafny.Map({p0: (self).f30})
                d_1658_v214_ = (d_1658_v214_).set(self.f35, (self).f34)
            elif True:
                d_1659___mcc_h48_ = source28_.cf51
                d_1660_cf51_ = d_1659___mcc_h48_
                (globalState).f13 = (self).f34
                (globalState).f1 = p0
                (globalState).f13 = default__.safeModuloInt(((self).f34) - (669), (0) - ((0) - (-218)))
                d_1661_v215_: _dafny.Map
                d_1661_v215_ = _dafny.Map({(self).f34: p0})
                rhs237_ = not(p0)
                rhs238_ = (0) - (default__.fm46((0) - ((self).f34), (self).f34, ((d_1661_v215_)[(self).f30] if ((self).f30) in (d_1661_v215_) else p0), globalState))
                lhs186_ = self
                lhs186_.f31 = rhs237_
                r3 = rhs238_
            d_1662_v216_: _dafny.Seq
            d_1662_v216_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylj"))
            d_1662_v216_ = d_1662_v216_
            rhs239_ = (self).f34
            rhs240_ = (self).f30
            lhs187_ = globalState
            r2 = rhs239_
            lhs187_.f13 = rhs240_
        d_1663_v217_: _dafny.Array
        nw284_ = _dafny.Array(int(0), 24)
        d_1663_v217_ = nw284_
        d_1664_v218_: _dafny.Array
        nw285_ = _dafny.Array(False, 24)
        d_1664_v218_ = nw285_
        d_1665_v219_: _dafny.Seq
        d_1665_v219_ = _dafny.SeqWithoutIsStrInference([d_1664_v218_, d_1664_v218_, d_1664_v218_, d_1664_v218_])
        index219_ = default__.safeIndex(484, (d_1663_v217_).length(0))
        (d_1663_v217_)[index219_] = len(d_1665_v219_)
        index220_ = default__.safeIndex(484, (d_1663_v217_).length(0))
        (d_1663_v217_)[index220_] = (self).f30
        r2 = (d_1663_v217_)[default__.safeIndex(484, (d_1663_v217_).length(0))]
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1663_v217_).length(0)):
            d_1666_i22_: int = guard_loop_2_
            if (True) and (((0) <= (d_1666_i22_)) and ((d_1666_i22_) < ((d_1663_v217_).length(0)))):
                (d_1663_v217_)[(d_1666_i22_)] = (d_1666_i22_) * ((self).f30)
        r0 = True
        r1 = self.f31
        d_1667_v220_: _dafny.Set
        d_1667_v220_ = _dafny.Set({(d_1663_v217_)[default__.safeIndex(484, (d_1663_v217_).length(0))]})
        d_1668_v221_: _dafny.Map
        d_1668_v221_ = _dafny.Map({(self).f30: d_1667_v220_})
        d_1669_v222_: _dafny.Map
        d_1669_v222_ = _dafny.Map({p0: (d_1668_v221_) | (d_1668_v221_)})
        r2 = (0) - (len(((d_1669_v222_)[self.f31] if (self.f31) in (d_1669_v222_) else (d_1668_v221_) | (d_1668_v221_))))
        d_1670_v223_: _dafny.Seq
        d_1670_v223_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urgiph"))
        d_1671_v224_: _dafny.Seq
        d_1671_v224_ = _dafny.SeqWithoutIsStrInference([(self).f34, len(d_1670_v223_), ((d_1663_v217_)[default__.safeIndex(484, (d_1663_v217_).length(0))]) * ((d_1663_v217_)[default__.safeIndex(484, (d_1663_v217_).length(0))])])
        r3 = (d_1671_v224_)[default__.safeIndex((d_1663_v217_)[default__.safeIndex(484, (d_1663_v217_).length(0))], len(d_1671_v224_))]
        return r0, r1, r2, r3


class C7(T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f30: int = int(0)
        self._f41: int = int(0)
        self._f42: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f41, f42, f30, f31, f28, f29):
        (self)._f41 = f41
        (self)._f42 = f42
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm3(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1672_i0_ in range(default__.abs(557))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lytqd")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gig")) if self.f28 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbkwbpt"))))

    def fm2(self, globalState):
        return _dafny.Map({self.f28: (self).f41})

    def fm19(self, globalState):
        return self.f29

    def fm20(self, p0, p1, globalState):
        return (0) - ((0) - ((self).f30))

    def m7(self, p0, p1, p2, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        r2: int = int(0)
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (self).f29 = not (self.f31) or (p2)
        (self).f31 = p2
        d_1673_v0_: _dafny.Map
        d_1673_v0_ = _dafny.Map({_dafny.Map({p0: False}): (self).f41})
        hi8_ = (self).f41
        for d_1674_i0_ in range(len(d_1673_v0_), hi8_):
            d_1675_v1_: _dafny.Seq
            d_1675_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fiw"))
            r3 = ((self).fm3(p0, self.f31, globalState)) + ((d_1675_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phvd"))))
            (globalState).f13 = ((self).f30) - (d_1674_i0_)
            r2 = 445
            d_1676_v2_: _dafny.MultiSet
            d_1676_v2_ = _dafny.MultiSet([(self).fm20((self).f41, False, globalState), (self).f30, d_1674_i0_])
            d_1676_v2_ = (self).f42
        d_1677_v3_: _dafny.Seq
        d_1677_v3_ = _dafny.SeqWithoutIsStrInference([(self).f41])
        d_1678_v4_: _dafny.Map
        d_1678_v4_ = _dafny.Map({not(p2): -371})
        (globalState).f13 = (d_1677_v3_)[default__.safeIndex(((d_1678_v4_)[self.f29] if (self.f29) in (d_1678_v4_) else (self).f41), len(d_1677_v3_))]
        d_1679_v5_: str
        d_1679_v5_ = _dafny.CodePoint('r')
        d_1680_v6_: _dafny.Seq
        d_1680_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdeidmrnl"))
        d_1681_v7_: _dafny.Array
        nw286_ = _dafny.Array(None, 28)
        nw286_[int(0)] = d_1679_v5_
        nw286_[int(1)] = _dafny.CodePoint('q')
        nw286_[int(2)] = d_1679_v5_
        nw286_[int(3)] = d_1679_v5_
        nw286_[int(4)] = d_1679_v5_
        nw286_[int(5)] = d_1679_v5_
        nw286_[int(6)] = d_1679_v5_
        nw286_[int(7)] = d_1679_v5_
        nw286_[int(8)] = d_1679_v5_
        nw286_[int(9)] = d_1679_v5_
        nw286_[int(10)] = _dafny.CodePoint('e')
        nw286_[int(11)] = default__.fm21(d_1679_v5_, len(_dafny.Map({(self).f41: False})), globalState)
        nw286_[int(12)] = d_1679_v5_
        nw286_[int(13)] = d_1679_v5_
        nw286_[int(14)] = d_1679_v5_
        nw286_[int(15)] = d_1679_v5_
        nw286_[int(16)] = d_1679_v5_
        nw286_[int(17)] = d_1679_v5_
        nw286_[int(18)] = d_1679_v5_
        nw286_[int(19)] = d_1679_v5_
        nw286_[int(20)] = default__.fm21(default__.fm21(d_1679_v5_, (self).f30, globalState), (self).f30, globalState)
        nw286_[int(21)] = (d_1680_v6_)[default__.safeIndex((self).f30, len(d_1680_v6_))]
        nw286_[int(22)] = _dafny.CodePoint('e')
        nw286_[int(23)] = d_1679_v5_
        nw286_[int(24)] = d_1679_v5_
        nw286_[int(25)] = d_1679_v5_
        nw286_[int(26)] = d_1679_v5_
        nw286_[int(27)] = d_1679_v5_
        d_1681_v7_ = nw286_
        d_1682_v8_: C0
        nw287_ = C0()
        nw287_.ctor__(d_1681_v7_)
        d_1682_v8_ = nw287_
        d_1683_v9_: _dafny.Map
        d_1683_v9_ = _dafny.Map({(self).f41: (self).f30})
        d_1684_v10_: _dafny.Array
        nw288_ = _dafny.Array(None, 20)
        nw288_[int(0)] = (self).f41
        nw288_[int(1)] = ((0) - ((self).f41) if p0 else 277)
        nw288_[int(2)] = (self).f41
        nw288_[int(3)] = (0) - ((self).f41)
        nw288_[int(4)] = (0) - ((self).f41)
        nw288_[int(5)] = (self).fm20((self).f41, not(True), globalState)
        nw288_[int(6)] = (self).f30
        nw288_[int(7)] = (self).f30
        nw288_[int(8)] = (self).f30
        nw288_[int(9)] = (self).fm20((0) - ((0) - ((self).f30)), p0, globalState)
        nw288_[int(10)] = ((0) - (len(d_1683_v9_))) * ((self).f41)
        nw288_[int(11)] = (self).fm20((self).f41, self.f31, globalState)
        nw288_[int(12)] = (self).f30
        nw288_[int(13)] = ((self).f41) + ((self).fm20((self).f41, p0, globalState))
        nw288_[int(14)] = (0) - (((d_1678_v4_)[self.f28] if (self.f28) in (d_1678_v4_) else (self).f30))
        nw288_[int(15)] = (self).f30
        nw288_[int(16)] = len(_dafny.Set({(self).f41}))
        nw288_[int(17)] = (0) - ((self).f41)
        nw288_[int(18)] = (self).f30
        nw288_[int(19)] = default__.safeModuloInt((self).fm20((self).f30, (self).fm19(globalState), globalState), (self).f41)
        d_1684_v10_ = nw288_
        index221_ = default__.safeIndex(943, (d_1684_v10_).length(0))
        (d_1684_v10_)[index221_] = default__.safeDivisionInt((self).f30, (0) - ((self).f30))
        index222_ = default__.safeIndex(398, (p1).length(0))
        (p1)[index222_] = p0
        index223_ = default__.safeIndex(513, (d_1684_v10_).length(0))
        (d_1684_v10_)[index223_] = (self).f41
        d_1685_v11_: _dafny.Map
        d_1685_v11_ = _dafny.Map({(self).f41: p2})
        d_1686_v12_: _dafny.MultiSet
        d_1686_v12_ = _dafny.MultiSet([self.f28])
        index224_ = default__.safeIndex(943, (d_1684_v10_).length(0))
        index225_ = default__.safeIndex(398, (p1).length(0))
        index226_ = default__.safeIndex(513, (d_1684_v10_).length(0))
        rhs241_ = (self).f30
        rhs242_ = (self).f41
        rhs243_ = not(((self).f30) > (((self).f30) - (len(d_1685_v11_))))
        rhs244_ = (_dafny.MultiSet([p2])).ispropersubset(d_1686_v12_)
        rhs245_ = ((self).f30) - ((self).f30)
        lhs188_ = globalState
        lhs189_ = d_1684_v10_
        lhs190_ = default__.safeIndex(943, (d_1684_v10_).length(0))
        lhs191_ = p1
        lhs192_ = default__.safeIndex(398, (p1).length(0))
        lhs193_ = self
        lhs194_ = d_1684_v10_
        lhs195_ = default__.safeIndex(513, (d_1684_v10_).length(0))
        lhs188_.f13 = rhs241_
        lhs189_[lhs190_] = rhs242_
        lhs191_[lhs192_] = rhs243_
        lhs193_.f31 = rhs244_
        lhs194_[lhs195_] = rhs245_
        d_1687_v13_: D5
        d_1687_v13_ = D5_DC13(p0, d_1678_v4_, d_1679_v5_)
        d_1688_v14_: D5
        d_1688_v14_ = D5_DC13(self.f28, (d_1687_v13_).cf28, d_1679_v5_)
        pat_let_tv15_ = d_1679_v5_
        pat_let_tv16_ = d_1679_v5_
        def lambda62_(source29_):
            if source29_.is_DC13:
                d_1689___mcc_h0_ = source29_.cf27
                d_1690___mcc_h1_ = source29_.cf28
                d_1691___mcc_h2_ = source29_.cf29
                d_1692_cf29_ = d_1691___mcc_h2_
                d_1693_cf28_ = d_1690___mcc_h1_
                d_1694_cf27_ = d_1689___mcc_h0_
                return _dafny.CodePoint('e')
            elif source29_.is_DC14:
                d_1695___mcc_h3_ = source29_.cf30
                d_1696___mcc_h4_ = source29_.cf31
                d_1697_cf31_ = d_1696___mcc_h4_
                d_1698_cf30_ = d_1695___mcc_h3_
                return pat_let_tv15_
            elif True:
                d_1699___mcc_h5_ = source29_.cf26
                d_1700_cf26_ = d_1699___mcc_h5_
                return pat_let_tv16_

        r0 = lambda62_(d_1687_v13_)
        d_1701_v15_: _dafny.Seq
        d_1701_v15_ = _dafny.SeqWithoutIsStrInference([d_1683_v9_, d_1683_v9_, d_1683_v9_, _dafny.Map({(self).fm20((self).f41, p0, globalState): len(d_1685_v11_)})])
        r1 = (d_1701_v15_ if p0 else _dafny.SeqWithoutIsStrInference([d_1683_v9_ for d_1702_i1_ in range(default__.abs(561))]))
        r2 = default__.safeModuloInt(len(d_1680_v6_), (d_1684_v10_)[default__.safeIndex(943, (d_1684_v10_).length(0))])
        r3 = d_1680_v6_
        return r0, r1, r2, r3

    def m8(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_1703_v0_: T0
        nw289_ = C1()
        nw289_.ctor__(self.f29, self.f29, True)
        d_1703_v0_ = nw289_
        d_1704_v1_: _dafny.Map
        d_1704_v1_ = _dafny.Map({d_1703_v0_: 161})
        d_1705_v2_: str
        d_1705_v2_ = _dafny.CodePoint('w')
        d_1706_v3_: _dafny.Map
        d_1706_v3_ = _dafny.Map({d_1703_v0_.f28: (self).fm19(globalState)})
        d_1707_v4_: T2
        nw290_ = C6()
        nw290_.ctor__(d_1705_v2_, (self).f30, self.f29, _dafny.SeqWithoutIsStrInference([self.f29]), default__.fm51(d_1703_v0_.f29, d_1705_v2_, globalState), len(d_1706_v3_), self.f29, self.f29, self.f29)
        d_1707_v4_ = nw290_
        d_1708_v5_: _dafny.Map
        d_1708_v5_ = _dafny.Map({(((d_1704_v1_)[d_1703_v0_] if (d_1703_v0_) in (d_1704_v1_) else (self).f30)) <= ((self).f41): (len(_dafny.SeqWithoutIsStrInference([d_1707_v4_, d_1707_v4_]))) + ((self).f30)})
        d_1708_v5_ = (d_1708_v5_).set(self.f31, (d_1707_v4_).f30)
        d_1709_v6_: D5
        d_1709_v6_ = D5_DC13(d_1707_v4_.f31, d_1708_v5_, d_1705_v2_)
        d_1710_v7_: _dafny.Map
        d_1710_v7_ = _dafny.Map({default__.fm24(d_1703_v0_.f29, 487, globalState): (d_1707_v4_).f30})
        d_1711_v8_: _dafny.Seq
        d_1711_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbffckyc"))
        d_1712_v9_: C5
        nw291_ = C5()
        nw291_.ctor__((d_1709_v6_).cf27, 593, (self).f41, self.f31, (d_1707_v4_).f32, (d_1707_v4_).f33, len(_dafny.Map({d_1710_v7_: (0) - ((d_1707_v4_).f30)})), True, ((d_1707_v4_).f32)[default__.safeIndex(((self).f42).cardinality, len((d_1707_v4_).f32))], (_dafny.SeqWithoutIsStrInference([d_1705_v2_ for d_1713_i0_ in range(default__.abs(418))])) != (d_1711_v8_))
        d_1712_v9_ = nw291_
        (d_1703_v0_).f29 = ((d_1707_v4_).f32)[default__.safeIndex((self).f30, len((d_1707_v4_).f32))]
        d_1714_v10_: _dafny.Seq
        d_1714_v10_ = _dafny.SeqWithoutIsStrInference([len(d_1708_v5_)])
        hi9_ = (d_1712_v9_).f48
        for d_1715_i1_ in range((len(d_1714_v10_)) - ((self).f30), hi9_):
            d_1716_v11_: D8
            d_1716_v11_ = D8_DC24((d_1712_v9_).f47, (self).f41, (d_1707_v4_).f32, (d_1712_v9_).f47, d_1714_v10_)
            d_1717_v12_: C4
            nw292_ = C4()
            nw292_.ctor__(len(d_1711_v8_), (d_1712_v9_).f47, (d_1716_v11_).cf48, ((d_1707_v4_).f33) + ((d_1707_v4_).f33), (d_1707_v4_).f30, ((d_1707_v4_).f30) < ((d_1712_v9_).f48), (d_1712_v9_).f47, d_1703_v0_.f29)
            d_1717_v12_ = nw292_
            index227_ = default__.safeIndex(581, (p0).length(0))
            (p0)[index227_] = self.f29
            index228_ = default__.safeIndex(581, (p0).length(0))
            (p0)[index228_] = (d_1717_v12_).fm44((0) - (d_1715_i1_), globalState)
            d_1718_v13_: _dafny.Map
            d_1718_v13_ = _dafny.Map({d_1715_i1_: d_1703_v0_.f28})
            d_1718_v13_ = (d_1718_v13_).set(len(((d_1707_v4_).f32).set(default__.safeIndex(-685, len((d_1707_v4_).f32)), d_1703_v0_.f29)), d_1707_v4_.f29)
            d_1719_v14_: _dafny.Array
            def lambda63_(d_1720_i2_):
                return (d_1720_i2_) + (962)

            init39_ = lambda63_
            nw293_ = _dafny.Array(None, 13)
            for i0_39_ in range(nw293_.length(0)):
                nw293_[i0_39_] = init39_(i0_39_)
            d_1719_v14_ = nw293_
            d_1721_v15_: _dafny.Map
            d_1721_v15_ = _dafny.Map({d_1719_v14_: len((d_1707_v4_).f32)})
            d_1722_v16_: D7
            d_1722_v16_ = D7_DC20(D7_DC20(D7_DC18(d_1721_v15_)))
            d_1723_v17_: D7
            d_1723_v17_ = D7_DC20(d_1722_v16_)
            pat_let_tv17_ = d_1722_v16_
            d_1724_v18_: _dafny.Array
            nw294_ = _dafny.Array(None, 19)
            nw294_[int(0)] = d_1723_v17_
            nw294_[int(1)] = d_1723_v17_
            nw294_[int(2)] = d_1723_v17_
            nw294_[int(3)] = d_1723_v17_
            nw294_[int(4)] = d_1723_v17_
            nw294_[int(5)] = d_1723_v17_
            nw294_[int(6)] = d_1723_v17_
            nw294_[int(7)] = D7_DC20(d_1722_v16_)
            nw294_[int(8)] = d_1723_v17_
            nw294_[int(9)] = d_1723_v17_
            nw294_[int(10)] = D7_DC20(d_1722_v16_)
            nw294_[int(11)] = d_1723_v17_
            nw294_[int(12)] = d_1723_v17_
            nw294_[int(13)] = d_1723_v17_
            nw294_[int(14)] = d_1723_v17_
            nw294_[int(15)] = d_1723_v17_
            def iife89_(_pat_let23_0):
                def iife90_(d_1725_dt__update__tmp_h0_):
                    def iife91_(_pat_let24_0):
                        def iife92_(d_1726_dt__update_hcf39_h0_):
                            return D7_DC20(d_1726_dt__update_hcf39_h0_)
                        return iife92_(_pat_let24_0)
                    return iife91_(pat_let_tv17_)
                return iife90_(_pat_let23_0)
            nw294_[int(16)] = iife89_(d_1723_v17_)
            nw294_[int(17)] = d_1723_v17_
            nw294_[int(18)] = d_1723_v17_
            d_1724_v18_ = nw294_
            d_1727_v19_: _dafny.Array
            d_1727_v19_ = d_1724_v18_
            d_1728_v20_: _dafny.Array
            nw295_ = _dafny.Array(None, 17)
            nw295_[int(0)] = (d_1724_v18_ if (p0)[default__.safeIndex(581, (p0).length(0))] else d_1724_v18_)
            nw295_[int(1)] = d_1724_v18_
            nw295_[int(2)] = d_1724_v18_
            nw295_[int(3)] = d_1724_v18_
            nw295_[int(4)] = d_1724_v18_
            nw295_[int(5)] = d_1724_v18_
            nw295_[int(6)] = (d_1727_v19_)
            nw295_[int(7)] = d_1724_v18_
            nw295_[int(8)] = d_1724_v18_
            nw295_[int(9)] = d_1724_v18_
            nw295_[int(10)] = d_1724_v18_
            nw295_[int(11)] = d_1724_v18_
            nw295_[int(12)] = d_1724_v18_
            nw295_[int(13)] = d_1724_v18_
            nw295_[int(14)] = d_1724_v18_
            nw295_[int(15)] = d_1724_v18_
            nw295_[int(16)] = d_1724_v18_
            d_1728_v20_ = nw295_
            d_1729_v21_: _dafny.Array
            nw296_ = _dafny.Array(_dafny.CodePoint('D'), 8)
            d_1729_v21_ = nw296_
            index229_ = default__.safeIndex(294, (d_1729_v21_).length(0))
            (d_1729_v21_)[index229_] = d_1705_v2_
            index230_ = default__.safeIndex(294, (d_1729_v21_).length(0))
            rhs246_ = d_1728_v20_
            rhs247_ = default__.safeDivisionInt((d_1712_v9_).f48, default__.safeModuloInt((self).f41, d_1715_i1_))
            rhs248_ = _dafny.CodePoint('x')
            rhs249_ = d_1707_v4_.f29
            rhs250_ = d_1707_v4_.f31
            lhs196_ = d_1729_v21_
            lhs197_ = default__.safeIndex(294, (d_1729_v21_).length(0))
            lhs198_ = globalState
            lhs199_ = d_1703_v0_
            d_1728_v20_ = rhs246_
            r1 = rhs247_
            lhs196_[lhs197_] = rhs248_
            lhs198_.f12 = rhs249_
            lhs199_.f29 = rhs250_
        if self.f29:
            if (d_1712_v9_).f47:
                d_1730_v22_: _dafny.Array
                def lambda64_(d_1731_i3_):
                    return _dafny.SeqWithoutIsStrInference([False])

                init40_ = lambda64_
                nw297_ = _dafny.Array(None, 27)
                for i0_40_ in range(nw297_.length(0)):
                    nw297_[i0_40_] = init40_(i0_40_)
                d_1730_v22_ = nw297_
                rhs251_ = d_1730_v22_
                rhs252_ = (d_1707_v4_).f30
                rhs253_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgo"))) + (d_1711_v8_)
                lhs200_ = globalState
                d_1730_v22_ = rhs251_
                lhs200_.f13 = rhs252_
                d_1711_v8_ = rhs253_
                d_1732_v23_: _dafny.Map
                d_1732_v23_ = _dafny.Map({(d_1707_v4_).f30: d_1708_v5_})
                rhs254_ = default__.fm46(len(((d_1732_v23_)[len((D8_DC24((d_1712_v9_).f47, (d_1712_v9_).f48, (d_1707_v4_).f32, d_1703_v0_.f29, d_1714_v10_)).cf48)] if (len((D8_DC24((d_1712_v9_).f47, (d_1712_v9_).f48, (d_1707_v4_).f32, d_1703_v0_.f29, d_1714_v10_)).cf48)) in (d_1732_v23_) else d_1708_v5_)), (d_1712_v9_).f48, d_1703_v0_.f28, globalState)
                rhs255_ = d_1705_v2_
                lhs201_ = globalState
                lhs201_.f13 = rhs254_
                d_1705_v2_ = rhs255_
                (globalState).f26 = False
                d_1733_v24_: C1
                nw298_ = C1()
                nw298_.ctor__((d_1712_v9_).fm6(globalState), d_1703_v0_.f28, d_1707_v4_.f31)
                d_1733_v24_ = nw298_
                d_1734_v25_: _dafny.Map
                d_1734_v25_ = _dafny.Map({d_1733_v24_: len(d_1711_v8_)})
                d_1734_v25_ = (d_1734_v25_).set((d_1733_v24_ if True else d_1733_v24_), (d_1707_v4_).f30)
                (globalState).f26 = (d_1712_v9_).f47
            elif True:
                d_1735_v26_: _dafny.Array
                nw299_ = _dafny.Array(_dafny.Map({}), 23)
                d_1735_v26_ = nw299_
                d_1735_v26_ = d_1735_v26_
                rhs256_ = d_1703_v0_.f28
                rhs257_ = d_1705_v2_
                rhs258_ = d_1705_v2_
                lhs202_ = d_1703_v0_
                lhs202_.f28 = rhs256_
                d_1705_v2_ = rhs257_
                d_1705_v2_ = rhs258_
                index231_ = default__.safeIndex(198, (p0).length(0))
                (p0)[index231_] = (d_1712_v9_).f47
                index232_ = default__.safeIndex(198, (p0).length(0))
                (p0)[index232_] = (d_1712_v9_).f47
                d_1736_v27_: C1
                nw300_ = C1()
                nw300_.ctor__(d_1703_v0_.f29, False, ((d_1707_v4_).f30) >= ((d_1712_v9_).f48))
                d_1736_v27_ = nw300_
                d_1737_v28_: _dafny.Array
                nw301_ = _dafny.Array(int(0), 6)
                d_1737_v28_ = nw301_
                index233_ = default__.safeIndex(862, (d_1737_v28_).length(0))
                (d_1737_v28_)[index233_] = (self).f41
                index234_ = default__.safeIndex(862, (d_1737_v28_).length(0))
                (d_1737_v28_)[index234_] = (0) - ((0) - ((d_1707_v4_).f30))
            d_1738_v29_: _dafny.Map
            d_1738_v29_ = _dafny.Map({(self).f30: (d_1711_v8_) + (d_1711_v8_)})
            d_1738_v29_ = ((d_1738_v29_) | (d_1738_v29_)) | (d_1738_v29_)
            d_1739_v30_: _dafny.Array
            def lambda65_(d_1740_v2_):
                def lambda66_(d_1741_i4_):
                    return d_1740_v2_

                return lambda66_

            init41_ = lambda65_(d_1705_v2_)
            nw302_ = _dafny.Array(None, 5)
            for i0_41_ in range(nw302_.length(0)):
                nw302_[i0_41_] = init41_(i0_41_)
            d_1739_v30_ = nw302_
            index235_ = default__.safeIndex(109, (d_1739_v30_).length(0))
            (d_1739_v30_)[index235_] = _dafny.CodePoint('a')
            index236_ = default__.safeIndex(109, (d_1739_v30_).length(0))
            (d_1739_v30_)[index236_] = d_1705_v2_
            d_1742_v31_: _dafny.MultiSet
            d_1742_v31_ = _dafny.MultiSet([self.f28])
            d_1743_v32_: _dafny.Map
            d_1743_v32_ = _dafny.Map({(d_1707_v4_).f30: False})
            r1 = (((d_1742_v31_).set(d_1703_v0_.f29, default__.abs(len((d_1707_v4_).f32)))).cardinality) - (default__.safeDivisionInt(((d_1708_v5_)[not((D2_DC6((d_1707_v4_).f30, True, (d_1707_v4_).f30, d_1707_v4_.f29, (d_1707_v4_).f30)).cf12)] if (not((D2_DC6((d_1707_v4_).f30, True, (d_1707_v4_).f30, d_1707_v4_.f29, (d_1707_v4_).f30)).cf12)) in (d_1708_v5_) else len(d_1743_v32_)), (d_1707_v4_).f30))
            index237_ = default__.safeIndex(549, (p0).length(0))
            (p0)[index237_] = d_1707_v4_.f29
            index238_ = default__.safeIndex(549, (p0).length(0))
            (p0)[index238_] = d_1703_v0_.f29
        elif True:
            d_1744_v33_: _dafny.Map
            d_1744_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f30 for d_1745_i5_ in range(default__.abs(665))]): (d_1707_v4_).f30})
            (globalState).f13 = len((d_1744_v33_) | (d_1744_v33_))
            (d_1707_v4_).f28 = d_1707_v4_.f28
            d_1746_v34_: D4
            d_1746_v34_ = D4_DC11((d_1707_v4_).f30, d_1703_v0_.f28, (self).f41)
            d_1747_v35_: _dafny.Map
            d_1747_v35_ = _dafny.Map({d_1746_v34_: (self).f41})
            d_1748_v36_: _dafny.Seq
            d_1748_v36_ = _dafny.SeqWithoutIsStrInference([d_1747_v35_])
            rhs259_ = default__.safeModuloInt(96, (d_1712_v9_).f48)
            rhs260_ = default__.safeModuloInt((0) - (((d_1707_v4_).f30) - (default__.fm46(len(d_1711_v8_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtbdkgdg"))), d_1703_v0_.f28, globalState))), (d_1707_v4_).f30)
            rhs261_ = d_1711_v8_
            rhs262_ = default__.fm1((d_1748_v36_) < ((d_1748_v36_).set(default__.safeIndex((self).f30, len(d_1748_v36_)), d_1747_v35_)), (d_1712_v9_).f48, globalState)
            lhs203_ = globalState
            lhs204_ = globalState
            lhs205_ = d_1707_v4_
            lhs203_.f13 = rhs259_
            lhs204_.f13 = rhs260_
            d_1711_v8_ = rhs261_
            lhs205_.f31 = rhs262_
            d_1749_v37_: str
            d_1750_v38_: _dafny.Seq
            d_1751_v39_: int
            d_1752_v40_: _dafny.Seq
            out60_: str
            out61_: _dafny.Seq
            out62_: int
            out63_: _dafny.Seq
            out60_, out61_, out62_, out63_ = (self).m7(not(d_1703_v0_.f28), p0, default__.fm1((d_1712_v9_).f47, (self).f30, globalState), globalState)
            d_1749_v37_ = out60_
            d_1750_v38_ = out61_
            d_1751_v39_ = out62_
            d_1752_v40_ = out63_
            (d_1703_v0_).f29 = not(((d_1712_v9_).f47 if d_1703_v0_.f28 else (d_1707_v4_.f29) or (d_1707_v4_.f31)))
        d_1753_v41_: _dafny.Array
        nw303_ = _dafny.Array(int(0), 21)
        d_1753_v41_ = nw303_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1753_v41_).length(0)):
            d_1754_i6_: int = guard_loop_3_
            if (True) and (((0) <= (d_1754_i6_)) and ((d_1754_i6_) < ((d_1753_v41_).length(0)))):
                (d_1753_v41_)[(d_1754_i6_)] = default__.safeDivisionInt(d_1754_i6_, default__.safeDivisionInt((d_1707_v4_).f30, len(d_1711_v8_)))
        r0 = p0
        d_1755_v42_: _dafny.Set
        d_1755_v42_ = _dafny.Set({len(d_1711_v8_)})
        d_1756_v43_: _dafny.Map
        d_1756_v43_ = _dafny.Map({d_1707_v4_.f29: d_1755_v42_})
        r1 = (0) - (len(((d_1755_v42_).intersection(d_1755_v42_)).intersection(((d_1756_v43_)[d_1703_v0_.f29] if (d_1703_v0_.f29) in (d_1756_v43_) else _dafny.Set({(self).f30})))))
        return r0, r1

    @property
    def f41(self):
        return self._f41
    @property
    def f42(self):
        return self._f42

class C8(T4, T2, T3, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f36: str = _dafny.CodePoint('D')
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        self.f40: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f36(self):
        return self._f36
    @f36.setter
    def f36(self, value):
        self._f36 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f40, f36, f34, f35, f32, f33, f30, f31, f28, f29):
        (self).f40 = f40
        (self).f36 = f36
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm7(self, p0, p1, globalState):
        return self.f35

    def fm6(self, globalState):
        return self.f35

    def fm4(self, p0, p1, p2, p3, globalState):
        return _dafny.Set({(self).f34})

    def fm5(self, p0, p1, p2, p3, globalState):
        if self.f40:
            return (self).f32
        elif True:
            return ((self).f32) + (_dafny.SeqWithoutIsStrInference([self.f29]))

    def fm3(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlucvm")) if self.f28 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ef")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgolrgbnu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbm"))))

    def fm2(self, globalState):
        source30_ = D4_DC10(_dafny.MultiSet([747, 611]))
        if source30_.is_DC11:
            d_1757___mcc_h0_ = source30_.cf23
            d_1758___mcc_h1_ = source30_.cf24
            d_1759___mcc_h2_ = source30_.cf25
            d_1760_cf25_ = d_1759___mcc_h2_
            d_1761_cf24_ = d_1758___mcc_h1_
            d_1762_cf23_ = d_1757___mcc_h0_
            return _dafny.Map({False: (0) - ((self).f30)})
        elif True:
            d_1763___mcc_h3_ = source30_.cf22
            d_1764_cf22_ = d_1763___mcc_h3_
            return (_dafny.Map({self.f35: (0) - ((self).f30)})) | (_dafny.Map({self.f40: len(_dafny.Map({self.f29: self.f31}))}))

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (self).f36 = _dafny.CodePoint('a')
        if (((self).f30) < ((self).f30)) and (self.f28):
            d_1765_v0_: _dafny.Seq
            d_1765_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suoxd"))
            d_1766_v1_: _dafny.MultiSet
            d_1766_v1_ = _dafny.MultiSet([(self).f34, len(d_1765_v0_)])
            d_1767_v2_: _dafny.Seq
            d_1767_v2_ = _dafny.SeqWithoutIsStrInference([(self).f30])
            d_1768_v3_: _dafny.Seq
            d_1768_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f34, (d_1766_v1_).cardinality, (self).f30]), d_1767_v2_, d_1767_v2_, d_1767_v2_])
            d_1769_v4_: _dafny.Array
            def lambda67_(d_1770_i0_):
                return (d_1770_i0_) * ((self).f34)

            init42_ = lambda67_
            nw304_ = _dafny.Array(None, 28)
            for i0_42_ in range(nw304_.length(0)):
                nw304_[i0_42_] = init42_(i0_42_)
            d_1769_v4_ = nw304_
            d_1771_v5_: _dafny.Map
            d_1771_v5_ = _dafny.Map({d_1769_v4_: d_1765_v0_})
            d_1772_v6_: _dafny.MultiSet
            d_1772_v6_ = _dafny.MultiSet([self.f31])
            d_1773_v7_: T0
            nw305_ = C7()
            nw305_.ctor__(default__.safeModuloInt(-494, (self).f30), (default__.fm49(d_1766_v1_, self.f36, d_1768_v3_, globalState)) - (d_1766_v1_), len(((_dafny.Map({d_1769_v4_: d_1765_v0_})).set(d_1769_v4_, d_1765_v0_)) | (d_1771_v5_)), self.f40, (default__.fm34((self).f34, (self).f30, (d_1772_v6_).cardinality, not(self.f35), globalState)) in (d_1765_v0_), self.f29)
            d_1773_v7_ = nw305_
            d_1773_v7_ = d_1773_v7_
            r1 = (0) - (default__.fm39(True, d_1766_v1_, self.f40, globalState))
            index239_ = default__.safeIndex(605, (d_1769_v4_).length(0))
            (d_1769_v4_)[index239_] = len(((self).f32) + ((self).f32))
            index240_ = default__.safeIndex(605, (d_1769_v4_).length(0))
            (d_1769_v4_)[index240_] = (self).f34
            d_1774_v9_: _dafny.Map
            d_1774_v9_ = _dafny.Map({(self).f30: d_1769_v4_})
            d_1775_v10_: _dafny.Map
            d_1775_v10_ = _dafny.Map({(self).f34: True})
            d_1776_v11_: _dafny.Map
            d_1776_v11_ = _dafny.Map({len(d_1774_v9_): d_1775_v10_})
            d_1777_v13_: _dafny.Set
            def iife93_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in ((d_1776_v11_).set((self).f34, d_1775_v10_)).keys.Elements:
                    d_1778_v8_: int = compr_43_
                    if (d_1778_v8_) in ((d_1776_v11_).set((self).f34, d_1775_v10_)):
                        coll43_[(d_1778_v8_) + (24)] = self.f29
                return _dafny.Map(coll43_)
            def iife94_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in _dafny.IntegerRange(985, 854):
                    d_1779_v12_: int = compr_44_
                    if ((985) <= (d_1779_v12_)) and ((d_1779_v12_) < (854)):
                        coll44_[default__.safeDivisionInt(d_1779_v12_, (d_1769_v4_)[default__.safeIndex(605, (d_1769_v4_).length(0))])] = False
                return _dafny.Map(coll44_)
            d_1777_v13_ = _dafny.Set({iife93_()
            , iife94_()
            , d_1775_v10_})
            def iife95_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(-637, 477):
                    d_1780_v14_: int = compr_45_
                    if ((-637) <= (d_1780_v14_)) and ((d_1780_v14_) < (477)):
                        coll45_[(d_1780_v14_) * ((self).f30)] = d_1773_v7_.f28
                return _dafny.Map(coll45_)
            d_1777_v13_ = _dafny.Set({d_1775_v10_, iife95_()
            })
            d_1772_v6_ = d_1772_v6_
        elif True:
            d_1781_v15_: _dafny.Map
            d_1781_v15_ = _dafny.Map({True: self.f35})
            d_1782_v16_: _dafny.Seq
            d_1782_v16_ = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f34])
            d_1783_v17_: _dafny.Map
            d_1783_v17_ = _dafny.Map({d_1781_v15_: d_1782_v16_})
            d_1784_v18_: int
            out64_: int
            out64_ = default__.m0(((d_1783_v17_)[d_1781_v15_] if (d_1781_v15_) in (d_1783_v17_) else d_1782_v16_), self.f28, ((self).f30) - (-749), self.f29, globalState)
            d_1784_v18_ = out64_
            (globalState).f13 = (self).f34
            d_1785_v19_: _dafny.Array
            def lambda68_(d_1786_i1_):
                return (d_1786_i1_) * (len(_dafny.SeqWithoutIsStrInference([self.f28, self.f29])))

            init43_ = lambda68_
            nw306_ = _dafny.Array(None, 6)
            for i0_43_ in range(nw306_.length(0)):
                nw306_[i0_43_] = init43_(i0_43_)
            d_1785_v19_ = nw306_
            index241_ = default__.safeIndex(495, (d_1785_v19_).length(0))
            (d_1785_v19_)[index241_] = (self).f30
            index242_ = default__.safeIndex(495, (d_1785_v19_).length(0))
            (d_1785_v19_)[index242_] = (self).f34
            (globalState).f1 = self.f28
            d_1781_v15_ = (d_1781_v15_).set(self.f40, not((self.f40 if self.f28 else self.f31)))
        d_1787_i2_: int
        d_1787_i2_ = 0
        with _dafny.label("5"):
            while (not((self.f40) == (True)) if self.f29 else True):
                with _dafny.c_label("5"):
                    if (d_1787_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_1787_i2_ = (d_1787_i2_) + (1)
                    d_1788_v20_: _dafny.Seq
                    d_1788_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idpxt"))
                    d_1789_v21_: _dafny.Seq
                    d_1789_v21_ = _dafny.SeqWithoutIsStrInference([self.f35, self.f35, ((self).f30) >= (len(d_1788_v20_))])
                    d_1789_v21_ = d_1789_v21_
                    d_1790_v22_: _dafny.Array
                    def lambda69_(d_1791_i3_):
                        return _dafny.CodePoint('u')

                    init44_ = lambda69_
                    nw307_ = _dafny.Array(None, 13)
                    for i0_44_ in range(nw307_.length(0)):
                        nw307_[i0_44_] = init44_(i0_44_)
                    d_1790_v22_ = nw307_
                    d_1792_v23_: C0
                    nw308_ = C0()
                    nw308_.ctor__(d_1790_v22_)
                    d_1792_v23_ = nw308_
                    source31_ = D9_DC27(d_1792_v23_, d_1788_v20_)
                    if source31_.is_DC27:
                        d_1793___mcc_h0_ = source31_.cf53
                        d_1794___mcc_h1_ = source31_.cf54
                        d_1795_cf54_ = d_1794___mcc_h1_
                        d_1796_cf53_ = d_1793___mcc_h0_
                        d_1797_v24_: D9
                        d_1797_v24_ = D9_DC27(d_1796_cf53_, (d_1788_v20_ if ((self).f32)[default__.safeIndex((self).f34, len((self).f32))] else d_1795_cf54_))
                        rhs263_ = d_1797_v24_
                        rhs264_ = (self).fm3(self.f31, self.f40, globalState)
                        rhs265_ = self.f29
                        d_1797_v24_ = rhs263_
                        r2 = rhs264_
                        r0 = rhs265_
                        (globalState).f12 = self.f35
                        (globalState).f13 = ((self).f34 if self.f28 else (0) - ((self).f30))
                        d_1798_v25_: _dafny.Set
                        d_1798_v25_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f30 for d_1799_i4_ in range(default__.abs(707))])})
                        d_1800_v26_: _dafny.Map
                        d_1800_v26_ = _dafny.Map({(self).f30: d_1798_v25_})
                        d_1801_v27_: _dafny.Seq
                        d_1801_v27_ = _dafny.SeqWithoutIsStrInference([(self).f34])
                        d_1802_v29_: _dafny.Map
                        d_1802_v29_ = _dafny.Map({len(d_1795_cf54_): self.f36})
                        d_1803_v31_: D16
                        d_1803_v31_ = D16_DC39(d_1798_v25_)
                        d_1804_v32_: _dafny.Seq
                        d_1804_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f30, (self).f34]), d_1801_v27_, _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emyatjs")))), (self).f34])])
                        d_1805_v34_: _dafny.Array
                        nw309_ = _dafny.Array(None, 23)
                        nw309_[int(0)] = d_1798_v25_
                        nw309_[int(1)] = (d_1798_v25_).intersection(d_1798_v25_)
                        nw309_[int(2)] = ((d_1800_v26_)[(self).f34] if ((self).f34) in (d_1800_v26_) else d_1798_v25_)
                        nw309_[int(3)] = d_1798_v25_
                        nw309_[int(4)] = d_1798_v25_
                        nw309_[int(5)] = d_1798_v25_
                        nw309_[int(6)] = (_dafny.Set({d_1801_v27_, d_1801_v27_}) if self.f40 else d_1798_v25_)
                        nw309_[int(7)] = d_1798_v25_
                        nw309_[int(8)] = d_1798_v25_
                        nw309_[int(9)] = (_dafny.Set({(_dafny.SeqWithoutIsStrInference([(self).f34])).set(default__.safeIndex(-949, len(_dafny.SeqWithoutIsStrInference([(self).f34]))), (self).f34), d_1801_v27_, d_1801_v27_})).intersection(d_1798_v25_)
                        nw309_[int(10)] = d_1798_v25_
                        nw309_[int(11)] = d_1798_v25_
                        def iife96_():
                            coll46_ = _dafny.Set()
                            compr_46_: _dafny.Seq
                            for compr_46_ in (d_1798_v25_).Elements:
                                d_1806_v28_: _dafny.Seq = compr_46_
                                if (d_1806_v28_) in (d_1798_v25_):
                                    coll46_ = coll46_.union(_dafny.Set([d_1806_v28_]))
                            return _dafny.Set(coll46_)
                        nw309_[int(12)] = iife96_()
                        
                        def iife97_():
                            coll47_ = _dafny.Set()
                            compr_47_: int
                            for compr_47_ in (d_1802_v29_).keys.Elements:
                                d_1807_v30_: int = compr_47_
                                if (d_1807_v30_) in (d_1802_v29_):
                                    coll47_ = coll47_.union(_dafny.Set([(d_1807_v30_) * (-831)]))
                            return _dafny.Set(coll47_)
                        nw309_[int(13)] = (_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f34, len(_dafny.Map({self.f35: False})), len(iife97_()
                        ), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wt")))])})) - ((d_1803_v31_).cf72)
                        nw309_[int(14)] = d_1798_v25_
                        nw309_[int(15)] = d_1798_v25_
                        nw309_[int(16)] = d_1798_v25_
                        nw309_[int(17)] = d_1798_v25_
                        nw309_[int(18)] = d_1798_v25_
                        nw309_[int(19)] = (default__.fm54(536, d_1801_v27_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1808_i5_ in range(default__.abs(892))])), globalState)) - (d_1798_v25_)
                        nw309_[int(20)] = d_1798_v25_
                        nw309_[int(21)] = _dafny.Set({d_1801_v27_})
                        def iife98_():
                            coll48_ = _dafny.Set()
                            compr_48_: _dafny.Seq
                            for compr_48_ in (d_1804_v32_).Elements:
                                d_1809_v33_: _dafny.Seq = compr_48_
                                if (d_1809_v33_) in (d_1804_v32_):
                                    coll48_ = coll48_.union(_dafny.Set([d_1809_v33_]))
                            return _dafny.Set(coll48_)
                        nw309_[int(22)] = iife98_()
                        
                        d_1805_v34_ = nw309_
                        index243_ = default__.safeIndex(926, (d_1805_v34_).length(0))
                        (d_1805_v34_)[index243_] = d_1798_v25_
                        index244_ = default__.safeIndex(926, (d_1805_v34_).length(0))
                        def iife99_():
                            coll49_ = _dafny.Map()
                            compr_49_: int
                            for compr_49_ in _dafny.IntegerRange(642, 927):
                                d_1810_v35_: int = compr_49_
                                if ((642) <= (d_1810_v35_)) and ((d_1810_v35_) < (927)):
                                    coll49_[(d_1810_v35_) * ((self).f34)] = (self).f30
                            return _dafny.Map(coll49_)
                        rhs266_ = d_1798_v25_
                        rhs267_ = ((self).f30) - ((self).f30)
                        rhs268_ = default__.safeModuloInt(((self).f30) - (len(iife99_()
                        )), (default__.fm32(self.f40, globalState)) - (default__.fm32(self.f35, globalState)))
                        rhs269_ = (self).f34
                        lhs206_ = d_1805_v34_
                        lhs207_ = default__.safeIndex(926, (d_1805_v34_).length(0))
                        lhs208_ = globalState
                        lhs209_ = globalState
                        lhs206_[lhs207_] = rhs266_
                        lhs208_.f13 = rhs267_
                        r1 = rhs268_
                        lhs209_.f13 = rhs269_
                    elif True:
                        d_1811___mcc_h2_ = source31_.cf52
                        d_1812_cf52_ = d_1811___mcc_h2_
                        d_1813_v36_: _dafny.Array
                        nw310_ = _dafny.Array(False, 22)
                        d_1813_v36_ = nw310_
                        (globalState).f21 = (d_1813_v36_ if self.f40 else d_1813_v36_)
                        d_1814_v37_: _dafny.Array
                        def lambda70_(d_1815_v20_):
                            def lambda71_(d_1816_i6_):
                                return (d_1816_i6_) - (len(_dafny.SeqWithoutIsStrInference([len(d_1815_v20_), (self).f30])))

                            return lambda71_

                        init45_ = lambda70_(d_1788_v20_)
                        nw311_ = _dafny.Array(None, 7)
                        for i0_45_ in range(nw311_.length(0)):
                            nw311_[i0_45_] = init45_(i0_45_)
                        d_1814_v37_ = nw311_
                        index245_ = default__.safeIndex(20, (d_1814_v37_).length(0))
                        (d_1814_v37_)[index245_] = ((self).f34) - ((self).f30)
                        index246_ = default__.safeIndex(20, (d_1814_v37_).length(0))
                        rhs270_ = ((self).f30) > (default__.safeModuloInt((d_1812_cf52_).f46, (d_1812_cf52_).f46))
                        rhs271_ = (self).f30
                        lhs210_ = globalState
                        lhs211_ = d_1814_v37_
                        lhs212_ = default__.safeIndex(20, (d_1814_v37_).length(0))
                        lhs210_.f1 = rhs270_
                        lhs211_[lhs212_] = rhs271_
                        d_1817_v38_: _dafny.Map
                        d_1817_v38_ = _dafny.Map({(self).f34: self.f40})
                        d_1818_v39_: C2
                        nw312_ = C2()
                        nw312_.ctor__(default__.safeDivisionInt(len(d_1817_v38_), (d_1814_v37_)[default__.safeIndex(20, (d_1814_v37_).length(0))]), default__.fm24(self.f31, len(d_1788_v20_), globalState), self.f29, self.f35, True)
                        d_1818_v39_ = nw312_
                        (globalState).f13 = ((d_1818_v39_).f46) - (514)
                    (globalState).f12 = self.f40
                    d_1819_v40_: _dafny.Array
                    nw313_ = _dafny.Array(_dafny.Array(None, 0), 13)
                    d_1819_v40_ = nw313_
                    d_1820_v41_: _dafny.Array
                    nw314_ = _dafny.Array(False, 4)
                    d_1820_v41_ = nw314_
                    nw315_ = _dafny.Array(None, 24)
                    nw315_[int(0)] = (d_1820_v41_ if self.f31 else d_1820_v41_)
                    nw315_[int(1)] = d_1820_v41_
                    nw315_[int(2)] = d_1820_v41_
                    nw315_[int(3)] = d_1820_v41_
                    nw315_[int(4)] = d_1820_v41_
                    nw315_[int(5)] = d_1820_v41_
                    nw315_[int(6)] = d_1820_v41_
                    nw315_[int(7)] = d_1820_v41_
                    nw315_[int(8)] = d_1820_v41_
                    nw315_[int(9)] = d_1820_v41_
                    nw315_[int(10)] = d_1820_v41_
                    nw315_[int(11)] = d_1820_v41_
                    nw315_[int(12)] = d_1820_v41_
                    nw315_[int(13)] = d_1820_v41_
                    nw315_[int(14)] = d_1820_v41_
                    nw315_[int(15)] = d_1820_v41_
                    nw315_[int(16)] = d_1820_v41_
                    nw315_[int(17)] = d_1820_v41_
                    nw315_[int(18)] = d_1820_v41_
                    nw315_[int(19)] = d_1820_v41_
                    nw315_[int(20)] = d_1820_v41_
                    nw315_[int(21)] = d_1820_v41_
                    nw315_[int(22)] = d_1820_v41_
                    nw315_[int(23)] = d_1820_v41_
                    d_1819_v40_ = nw315_
                    pass
            pass
        (self).f40 = self.f28
        (self).f35 = ((self).f30) != ((self).f30)
        d_1821_v42_: _dafny.MultiSet
        d_1821_v42_ = _dafny.MultiSet([default__.fm1(self.f29, (self).f30, globalState)])
        d_1822_v43_: D0
        d_1822_v43_ = D0_DC1(d_1821_v42_)
        d_1823_v44_: _dafny.Map
        d_1823_v44_ = _dafny.Map({d_1822_v43_: self.f40})
        d_1824_v45_: D8
        d_1824_v45_ = D8_DC21(d_1823_v44_)
        d_1825_v46_: _dafny.Seq
        d_1825_v46_ = _dafny.SeqWithoutIsStrInference([d_1824_v45_])
        d_1826_v47_: _dafny.Seq
        d_1826_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D8_DC21(_dafny.Map({D0_DC1(_dafny.MultiSet([self.f31])): self.f29})) for d_1827_i7_ in range(default__.abs(76))]), d_1825_v46_, _dafny.SeqWithoutIsStrInference([d_1824_v45_ for d_1828_i8_ in range(default__.abs(47))]), _dafny.SeqWithoutIsStrInference([d_1824_v45_, d_1824_v45_, d_1824_v45_, d_1824_v45_])])
        if (d_1826_v47_) < (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1824_v45_])])):
            d_1829_v48_: _dafny.Array
            nw316_ = _dafny.Array(_dafny.Seq({}), 21)
            d_1829_v48_ = nw316_
            d_1830_v49_: C3
            nw317_ = C3()
            nw317_.ctor__((self).f34, d_1829_v48_, _dafny.CodePoint('i'), 101, self.f35, (_dafny.SeqWithoutIsStrInference([True])) + ((self).f32), (self).f33, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))), self.f28, self.f29, (False) and (False))
            d_1830_v49_ = nw317_
            d_1831_v50_: _dafny.Set
            d_1831_v50_ = _dafny.Set({self.f35})
            d_1832_v51_: _dafny.Array
            def lambda72_(d_1833_i9_):
                return self.f36

            init46_ = lambda72_
            nw318_ = _dafny.Array(None, 18)
            for i0_46_ in range(nw318_.length(0)):
                nw318_[i0_46_] = init46_(i0_46_)
            d_1832_v51_ = nw318_
            d_1834_v52_: C0
            nw319_ = C0()
            nw319_.ctor__(d_1832_v51_)
            d_1834_v52_ = nw319_
            d_1835_v53_: _dafny.Map
            d_1836_v54_: bool
            d_1837_v55_: bool
            d_1838_v56_: int
            out65_: _dafny.Map
            out66_: bool
            out67_: bool
            out68_: int
            out65_, out66_, out67_, out68_ = (self).m6(default__.safeDivisionInt((self).f34, (d_1830_v49_).f44), default__.fm39(True, _dafny.MultiSet([len(d_1831_v50_)]), self.f31, globalState), self.f36, d_1834_v52_, globalState)
            d_1835_v53_ = out65_
            d_1836_v54_ = out66_
            d_1837_v55_ = out67_
            d_1838_v56_ = out68_
            d_1839_v57_: _dafny.Seq
            d_1839_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fu"))
            r2 = (d_1839_v57_) + (d_1839_v57_)
            (self).f36 = self.f36
            d_1840_v58_: _dafny.Map
            d_1840_v58_ = _dafny.Map({self.f35: (d_1821_v42_).cardinality})
            d_1840_v58_ = (d_1840_v58_).set((self).fm6(globalState), (self).f34)
        elif True:
            r1 = ((self).f30) * ((self).f30)
            r1 = (len((self).fm3(self.f31, self.f35, globalState))) - ((self).f30)
            (globalState).f1 = self.f35
            d_1841_v59_: _dafny.Array
            nw320_ = _dafny.Array(_dafny.Seq({}), 19)
            d_1841_v59_ = nw320_
            d_1842_v60_: C4
            nw321_ = C4()
            nw321_.ctor__(657, False, (self).f32, (self).f33, (self).f34, False, self.f31, self.f28)
            d_1842_v60_ = nw321_
            d_1843_v61_: _dafny.Seq
            d_1843_v61_ = _dafny.SeqWithoutIsStrInference([d_1842_v60_, d_1842_v60_, d_1842_v60_, d_1842_v60_, d_1842_v60_])
            index247_ = default__.safeIndex(542, (d_1841_v59_).length(0))
            (d_1841_v59_)[index247_] = (d_1843_v61_) + (d_1843_v61_)
            d_1844_v62_: D17
            d_1844_v62_ = D17_DC42(d_1843_v61_)
            index248_ = default__.safeIndex(542, (d_1841_v59_).length(0))
            rhs272_ = (d_1844_v62_).cf76
            rhs273_ = ((self).f30) >= ((self).f34)
            lhs213_ = d_1841_v59_
            lhs214_ = default__.safeIndex(542, (d_1841_v59_).length(0))
            lhs215_ = globalState
            lhs213_[lhs214_] = rhs272_
            lhs215_.f1 = rhs273_
            d_1845_v63_: _dafny.Map
            d_1845_v63_ = _dafny.Map({self.f28: (self).f30})
            d_1845_v63_ = (d_1845_v63_).set(self.f28, (self).f34)
        r0 = self.f35
        d_1846_v64_: _dafny.Map
        d_1846_v64_ = _dafny.Map({self.f31: self.f36})
        r1 = (len(d_1846_v64_)) * ((self).f34)
        d_1847_v65_: _dafny.Seq
        d_1847_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
        r2 = d_1847_v65_
        return r0, r1, r2

    def m6(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_1848_v0_: _dafny.Set
        d_1848_v0_ = _dafny.Set({self.f35, self.f29, self.f35})
        hi10_ = default__.safeDivisionInt(len(d_1848_v0_), (self).f30)
        for d_1849_i0_ in range((p1 if self.f31 else (self).f30), hi10_):
            d_1850_v1_: _dafny.Seq
            d_1850_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vq"))
            d_1851_v2_: _dafny.Seq
            d_1851_v2_ = _dafny.SeqWithoutIsStrInference([p1, d_1849_i0_, 43, d_1849_i0_, p0])
            d_1852_v4_: _dafny.Map
            d_1852_v4_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f34])): p1})
            d_1853_v5_: _dafny.Array
            def lambda73_(d_1854_i1_):
                return (d_1854_i1_) + ((0) - ((self).f34))

            init47_ = lambda73_
            nw322_ = _dafny.Array(None, 7)
            for i0_47_ in range(nw322_.length(0)):
                nw322_[i0_47_] = init47_(i0_47_)
            d_1853_v5_ = nw322_
            d_1855_v6_: _dafny.Map
            d_1855_v6_ = _dafny.Map({d_1853_v5_: self.f29})
            d_1856_v7_: T2
            nw323_ = C5()
            def iife100_():
                coll50_ = _dafny.Map()
                compr_50_: int
                for compr_50_ in (d_1852_v4_).keys.Elements:
                    d_1857_v3_: int = compr_50_
                    if (d_1857_v3_) in (d_1852_v4_):
                        coll50_[(d_1857_v3_) - ((self).f30)] = self.f28
                return _dafny.Map(coll50_)
            nw323_.ctor__(self.f35, len(iife100_()
            ), (0) - ((0) - (p0)), self.f28, default__.fm0((self).f34, False, self.f40, len(d_1850_v1_), globalState), ((self).f33).set(default__.safeIndex((self).f30, len((self).f33)), (self).f32), len(d_1855_v6_), False, self.f31, True)
            d_1856_v7_ = nw323_
            d_1858_v8_: _dafny.Map
            d_1858_v8_ = _dafny.Map({d_1851_v2_: d_1856_v7_})
            (globalState).f13 = len((d_1850_v1_).set(default__.safeIndex(len(d_1858_v8_), len(d_1850_v1_)), self.f36))
            (self).f35 = False
            (self).f36 = self.f36
            d_1859_v9_: D16
            d_1859_v9_ = D16_DC40(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcj")), (len(_dafny.Set({(_dafny.MultiSet(d_1851_v2_)).cardinality, (self).f34}))) + ((d_1856_v7_).f30))
            source32_ = d_1859_v9_
            if source32_.is_DC40:
                d_1860___mcc_h0_ = source32_.cf73
                d_1861___mcc_h1_ = source32_.cf74
                d_1862_cf74_ = d_1861___mcc_h1_
                d_1863_cf73_ = d_1860___mcc_h0_
                d_1864_v10_: _dafny.Array
                nw324_ = _dafny.Array(False, 19)
                d_1864_v10_ = nw324_
                d_1865_v11_: _dafny.Seq
                d_1865_v11_ = _dafny.SeqWithoutIsStrInference([d_1864_v10_, d_1864_v10_])
                d_1866_v12_: _dafny.Seq
                d_1866_v12_ = d_1865_v11_
                d_1865_v11_ = (d_1866_v12_)
                (globalState).f26 = self.f35
                d_1867_v13_: _dafny.Set
                d_1867_v13_ = _dafny.Set({default__.fm55((0) - (p0), d_1856_v7_.f29, len(_dafny.Set({True, d_1856_v7_.f31})), globalState), d_1848_v0_})
                d_1867_v13_ = d_1867_v13_
                d_1868_v14_: _dafny.Seq
                d_1868_v14_ = _dafny.SeqWithoutIsStrInference([d_1848_v0_])
                d_1869_v15_: D2
                d_1869_v15_ = D2_DC6(len(_dafny.SeqWithoutIsStrInference([(self).f34])), d_1856_v7_.f29, 423, self.f29, len((d_1868_v14_)[default__.safeIndex((d_1856_v7_).f30, len(d_1868_v14_))]))
                d_1869_v15_ = d_1869_v15_
            elif source32_.is_DC39:
                d_1870___mcc_h2_ = source32_.cf72
                d_1871_cf72_ = d_1870___mcc_h2_
                d_1872_v16_: _dafny.Array
                nw325_ = _dafny.Array(_dafny.Seq({}), 17)
                d_1872_v16_ = nw325_
                d_1873_v17_: C4
                nw326_ = C4()
                nw326_.ctor__(d_1849_i0_, default__.fm1(d_1856_v7_.f31, p0, globalState), _dafny.SeqWithoutIsStrInference([d_1856_v7_.f31]), (d_1856_v7_).f33, len(default__.fm0(553, self.f35, self.f29, p1, globalState)), d_1856_v7_.f31, False, d_1856_v7_.f28)
                d_1873_v17_ = nw326_
                d_1874_v18_: C3
                nw327_ = C3()
                nw327_.ctor__((self).f34, (d_1872_v16_ if d_1856_v7_.f31 else d_1872_v16_), self.f36, len(_dafny.SeqWithoutIsStrInference([len((d_1856_v7_).f32) for d_1875_i2_ in range(default__.abs(-686))])), self.f31, ((d_1856_v7_).f32) + (_dafny.SeqWithoutIsStrInference([d_1856_v7_.f31, self.f31, not(d_1856_v7_.f29)])), (d_1856_v7_).f33, (d_1856_v7_).f30, default__.fm1(self.f35, len((_dafny.Map({p1: d_1873_v17_})).set((0) - (len(d_1850_v1_)), d_1873_v17_)), globalState), self.f40, d_1856_v7_.f29)
                d_1874_v18_ = nw327_
                index249_ = default__.safeIndex(691, (d_1853_v5_).length(0))
                (d_1853_v5_)[index249_] = (self).f30
                index250_ = default__.safeIndex(691, (d_1853_v5_).length(0))
                (d_1853_v5_)[index250_] = default__.safeDivisionInt((d_1856_v7_).f30, (d_1856_v7_).f30)
                d_1876_v19_: _dafny.Array
                def lambda74_(d_1877_v7_):
                    def lambda75_(d_1878_i3_):
                        return (_dafny.MultiSet([_dafny.MultiSet([self.f29, self.f40])])).ispropersubset((_dafny.MultiSet([_dafny.MultiSet([d_1877_v7_.f28])])))

                    return lambda75_

                init48_ = lambda74_(d_1856_v7_)
                nw328_ = _dafny.Array(None, 19)
                for i0_48_ in range(nw328_.length(0)):
                    nw328_[i0_48_] = init48_(i0_48_)
                d_1876_v19_ = nw328_
                index251_ = default__.safeIndex(223, (d_1876_v19_).length(0))
                (d_1876_v19_)[index251_] = self.f28
                index252_ = default__.safeIndex(223, (d_1876_v19_).length(0))
                (d_1876_v19_)[index252_] = (p1) < (p0)
                d_1876_v19_ = d_1876_v19_
            elif True:
                d_1879___mcc_h3_ = source32_.cf75
                d_1880_cf75_ = d_1879___mcc_h3_
                (globalState).f1 = self.f40
                d_1881_v20_: _dafny.Array
                nw329_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_1881_v20_ = nw329_
                d_1882_v21_: _dafny.Array
                def lambda76_(d_1883_v7_):
                    def lambda77_(d_1884_i4_):
                        return d_1883_v7_.f28

                    return lambda77_

                init49_ = lambda76_(d_1856_v7_)
                nw330_ = _dafny.Array(None, 13)
                for i0_49_ in range(nw330_.length(0)):
                    nw330_[i0_49_] = init49_(i0_49_)
                d_1882_v21_ = nw330_
                index253_ = default__.safeIndex(720, (d_1881_v20_).length(0))
                (d_1881_v20_)[index253_] = d_1882_v21_
                index254_ = default__.safeIndex(720, (d_1881_v20_).length(0))
                rhs274_ = 522
                rhs275_ = d_1882_v21_
                lhs216_ = d_1881_v20_
                lhs217_ = default__.safeIndex(720, (d_1881_v20_).length(0))
                r3 = rhs274_
                lhs216_[lhs217_] = rhs275_
                d_1885_v22_: _dafny.Map
                d_1885_v22_ = _dafny.Map({d_1882_v21_: len(d_1850_v1_)})
                d_1886_v23_: C4
                nw331_ = C4()
                nw331_.ctor__((d_1856_v7_).f30, d_1856_v7_.f28, (d_1856_v7_).f32, (((d_1856_v7_).f33).set(default__.safeIndex(p1, len((d_1856_v7_).f33)), (self).f32)) + ((self).f33), default__.safeModuloInt((d_1856_v7_).f30, 143), ((d_1856_v7_).f30) == (len(d_1885_v22_)), self.f31, d_1856_v7_.f28)
                d_1886_v23_ = nw331_
                d_1887_v24_: _dafny.Set
                d_1887_v24_ = _dafny.Set({self.f36, default__.fm21(p2, p0, globalState)})
                d_1888_v25_: _dafny.Map
                d_1888_v25_ = _dafny.Map({self.f29: (self).f30})
                d_1889_v26_: D2
                d_1889_v26_ = D2_DC6(len(d_1887_v24_), d_1856_v7_.f29, 230, d_1856_v7_.f31, ((d_1888_v25_)[self.f28] if (self.f28) in (d_1888_v25_) else len(d_1851_v2_)))
                d_1890_v27_: _dafny.Array
                nw332_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                d_1890_v27_ = nw332_
                d_1891_v28_: _dafny.Seq
                d_1891_v28_ = _dafny.SeqWithoutIsStrInference([d_1850_v1_])
                rhs276_ = default__.fm56((self).f30, (d_1850_v1_) <= (d_1850_v1_), ((0) - ((d_1856_v7_).f30)) * (p0), globalState)
                rhs277_ = default__.safeDivisionInt(p1, (self).f30)
                rhs278_ = p1
                rhs279_ = d_1890_v27_
                rhs280_ = ((d_1891_v28_)[default__.safeIndex(len(d_1850_v1_), len(d_1891_v28_))]) != ((_dafny.SeqWithoutIsStrInference([p2 for d_1892_i5_ in range(default__.abs(443))])) + (d_1850_v1_))
                lhs218_ = globalState
                lhs219_ = globalState
                lhs220_ = d_1856_v7_
                d_1889_v26_ = rhs276_
                lhs218_.f13 = rhs277_
                lhs219_.f13 = rhs278_
                d_1890_v27_ = rhs279_
                lhs220_.f28 = rhs280_
        d_1893_v29_: _dafny.Array
        nw333_ = _dafny.Array(False, 14)
        d_1893_v29_ = nw333_
        d_1894_v30_: _dafny.Map
        d_1894_v30_ = _dafny.Map({self.f28: self.f29})
        d_1895_v31_: _dafny.Map
        d_1895_v31_ = _dafny.Map({self.f40: len(d_1894_v30_)})
        d_1896_v32_: _dafny.Map
        d_1896_v32_ = _dafny.Map({not(self.f40): default__.fm57(p2, len(d_1895_v31_), globalState)})
        index255_ = default__.safeIndex(82, (d_1893_v29_).length(0))
        (d_1893_v29_)[index255_] = default__.fm1(self.f29, default__.fm24(False, len(d_1896_v32_), globalState), globalState)
        index256_ = default__.safeIndex(82, (d_1893_v29_).length(0))
        (d_1893_v29_)[index256_] = ((self).f32)[default__.safeIndex((0) - ((0) - ((self).f34)), len((self).f32))]
        d_1897_i6_: int
        d_1897_i6_ = 0
        with _dafny.label("6"):
            while not(self.f28):
                with _dafny.c_label("6"):
                    if (d_1897_i6_) >= (100):
                        raise _dafny.Break("6")
                    d_1897_i6_ = (d_1897_i6_) + (1)
                    d_1898_v33_: _dafny.MultiSet
                    d_1898_v33_ = _dafny.MultiSet([self.f40])
                    d_1899_v34_: _dafny.MultiSet
                    d_1899_v34_ = _dafny.MultiSet([(self).f34, p1])
                    d_1900_v35_: _dafny.Seq
                    d_1900_v35_ = _dafny.SeqWithoutIsStrInference([(self).f34])
                    d_1901_v36_: D4
                    d_1901_v36_ = D4_DC11(p1, self.f40, ((d_1898_v33_)[self.f28] if (self.f28) in (d_1898_v33_) else p0))
                    d_1902_v37_: _dafny.MultiSet
                    d_1902_v37_ = _dafny.MultiSet([d_1901_v36_])
                    d_1903_v38_: _dafny.Seq
                    d_1903_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ff"))
                    d_1904_v39_: T0
                    nw334_ = C7()
                    nw334_.ctor__(((d_1898_v33_)[(d_1893_v29_)[default__.safeIndex(82, (d_1893_v29_).length(0))]] if ((d_1893_v29_)[default__.safeIndex(82, (d_1893_v29_).length(0))]) in (d_1898_v33_) else p0), (d_1899_v34_) - (_dafny.MultiSet([len(d_1900_v35_)])), p0, (d_1902_v37_).ispropersubset(d_1902_v37_), ((self).f32)[default__.safeIndex(len(d_1903_v38_), len((self).f32))], self.f29)
                    d_1904_v39_ = nw334_
                    d_1905_v40_: _dafny.Array
                    nw335_ = _dafny.Array(None, 9)
                    d_1905_v40_ = nw335_
                    d_1906_v41_: _dafny.Seq
                    d_1906_v41_ = _dafny.SeqWithoutIsStrInference([default__.fm49(d_1899_v34_, self.f36, _dafny.SeqWithoutIsStrInference([d_1900_v35_, _dafny.SeqWithoutIsStrInference([len(d_1903_v38_), len(_dafny.Map({d_1905_v40_: self.f31}))])]), globalState), d_1899_v34_])
                    rhs281_ = d_1904_v39_
                    rhs282_ = ((d_1906_v41_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1900_v35_) for d_1907_i7_ in range(default__.abs(963))]))) < (((d_1906_v41_).set(default__.safeIndex((self).f34, len(d_1906_v41_)), d_1899_v34_)) + ((d_1906_v41_).set(default__.safeIndex(492, len(d_1906_v41_)), _dafny.MultiSet([(self).f34]))))
                    d_1904_v39_ = rhs281_
                    r1 = rhs282_
                    index257_ = default__.safeIndex(82, (d_1893_v29_).length(0))
                    rhs283_ = (self).fm7(d_1899_v34_, (0) - ((self).f30), globalState)
                    rhs284_ = ((self).f30) * ((0) - ((self).f34))
                    lhs221_ = d_1893_v29_
                    lhs222_ = default__.safeIndex(82, (d_1893_v29_).length(0))
                    lhs223_ = globalState
                    lhs221_[lhs222_] = rhs283_
                    lhs223_.f13 = rhs284_
                    (globalState).f21 = d_1893_v29_
                    d_1908_v42_: _dafny.Array
                    nw336_ = _dafny.Array(_dafny.Seq({}), 26)
                    d_1908_v42_ = nw336_
                    index258_ = default__.safeIndex(529, (d_1908_v42_).length(0))
                    (d_1908_v42_)[index258_] = (self).f32
                    index259_ = default__.safeIndex(529, (d_1908_v42_).length(0))
                    (d_1908_v42_)[index259_] = (self).f32
                    pass
            pass
        d_1909_v43_: _dafny.Seq
        d_1909_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opebgfkic"))
        d_1910_v44_: _dafny.Array
        nw337_ = _dafny.Array(None, 13)
        nw337_[int(0)] = p1
        nw337_[int(1)] = (0) - ((0) - ((self).f30))
        nw337_[int(2)] = p0
        nw337_[int(3)] = len(d_1909_v43_)
        nw337_[int(4)] = p0
        nw337_[int(5)] = (self).f34
        nw337_[int(6)] = (self).f30
        nw337_[int(7)] = (self).f34
        nw337_[int(8)] = (self).f30
        nw337_[int(9)] = (0) - (len(d_1909_v43_))
        nw337_[int(10)] = p1
        nw337_[int(11)] = (self).f30
        nw337_[int(12)] = p0
        d_1910_v44_ = nw337_
        d_1911_v45_: _dafny.Seq
        d_1911_v45_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f28])), (self).f30])
        d_1912_v46_: _dafny.Map
        d_1912_v46_ = _dafny.Map({d_1910_v44_: (default__.fm23((d_1893_v29_)[default__.safeIndex(82, (d_1893_v29_).length(0))], p0, globalState)) + (d_1911_v45_)})
        d_1913_v47_: _dafny.Set
        d_1913_v47_ = _dafny.Set({len(d_1911_v45_)})
        d_1914_v48_: _dafny.Map
        d_1914_v48_ = _dafny.Map({d_1913_v47_: d_1912_v46_})
        d_1915_v49_: _dafny.Set
        d_1915_v49_ = d_1848_v0_
        d_1916_v50_: _dafny.Map
        d_1916_v50_ = _dafny.Map({not(self.f28): d_1912_v46_})
        d_1917_v51_: _dafny.Map
        d_1917_v51_ = _dafny.Map({len((d_1915_v49_)): ((d_1916_v50_)[self.f40] if (self.f40) in (d_1916_v50_) else d_1912_v46_)})
        d_1912_v46_ = ((d_1914_v48_)[d_1913_v47_] if (d_1913_v47_) in (d_1914_v48_) else ((d_1917_v51_)[p0] if (p0) in (d_1917_v51_) else d_1912_v46_))
        index260_ = default__.safeIndex(551, (d_1910_v44_).length(0))
        (d_1910_v44_)[index260_] = (self).f30
        index261_ = default__.safeIndex(551, (d_1910_v44_).length(0))
        (d_1910_v44_)[index261_] = (0) - ((((self).f30 if self.f29 else p0)) * (-661))
        (self).f29 = self.f29
        d_1918_v52_: _dafny.Map
        d_1918_v52_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([d_1895_v31_, d_1895_v31_])})
        d_1919_v53_: _dafny.Map
        d_1919_v53_ = _dafny.Map({26: self.f31})
        d_1920_v54_: _dafny.Seq
        d_1920_v54_ = _dafny.SeqWithoutIsStrInference([d_1895_v31_])
        d_1921_v55_: D3
        d_1921_v55_ = D3_DC7(((d_1918_v52_)[(0) - (len(d_1919_v53_))] if ((0) - (len(d_1919_v53_))) in (d_1918_v52_) else d_1920_v54_))
        d_1922_v56_: _dafny.Map
        d_1922_v56_ = _dafny.Map({d_1921_v55_: (self).f30})
        d_1923_v57_: _dafny.Map
        d_1923_v57_ = _dafny.Map({d_1922_v56_: d_1910_v44_})
        r0 = d_1923_v57_
        r1 = (self).fm7(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC6(438, self.f35, 871, self.f28, len(d_1911_v45_))).cf11 for d_1924_i8_ in range(default__.abs(868))])), default__.safeDivisionInt((_dafny.MultiSet([(d_1893_v29_)[default__.safeIndex(82, (d_1893_v29_).length(0))], self.f40, self.f35])).cardinality, len(d_1909_v43_)), globalState)
        r2 = self.f28
        r3 = p1
        return r0, r1, r2, r3


class C9(T4, T3, T2, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f36: str = _dafny.CodePoint('D')
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        self._f38: int = int(0)
        self._f39: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f36(self):
        return self._f36
    @f36.setter
    def f36(self, value):
        self._f36 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f38, f39, f36, f34, f35, f32, f33, f30, f31, f28, f29):
        (self)._f38 = f38
        (self)._f39 = f39
        (self).f36 = f36
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm7(self, p0, p1, globalState):
        if True:
            return False
        elif True:
            return (D0_DC0(self.f35)).cf0

    def fm6(self, globalState):
        return ((_dafny.Map({not(False): self.f35}) if self.f29 else _dafny.Map({self.f29: (D0_DC0(self.f35)).cf0}))) != ((_dafny.Map({not(((self).f32)[default__.safeIndex((self).f30, len((self).f32))]): self.f31})) | (_dafny.Map({self.f31: self.f28})))

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife101_():
            coll51_ = _dafny.Map()
            compr_51_: int
            for compr_51_ in (_dafny.Set({(self).f38})).Elements:
                d_1925_v0_: int = compr_51_
                if (d_1925_v0_) in (_dafny.Set({(self).f38})):
                    coll51_[default__.safeDivisionInt(d_1925_v0_, (self).f38)] = self.f31
            return _dafny.Map(coll51_)
        return _dafny.Set({(((D4_DC10(_dafny.MultiSet([(self).f34]))).cf22) - (_dafny.MultiSet([571, ((D0_DC1(_dafny.MultiSet([self.f28]))).cf1).cardinality, (self).f34, (self).f30, len(iife101_()
        )]))).cardinality})

    def fm5(self, p0, p1, p2, p3, globalState):
        return (self).f32

    def fm3(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxjjkbq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "buh")))

    def fm2(self, globalState):
        return (_dafny.Map({self.f31: (self).f38})) | ((_dafny.Map({self.f29: (self).f30})) | (_dafny.Map({not(self.f35): (self).f34})))

    def fm16(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({(0) - ((self).f30)})).intersection(_dafny.Set({(self).f38, (self).f30}))).ispropersubset(_dafny.Set({715, (self).f30}))

    def fm17(self, p0, p1, globalState):
        return (len((_dafny.Map({_dafny.Set({self.f35}): _dafny.MultiSet([self.f28, not(self.f31), self.f31])})) | (_dafny.Map({_dafny.Set({True, self.f31}): _dafny.MultiSet([self.f28])})))) + ((self).f34)

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1926_v0_: _dafny.Set
        d_1926_v0_ = _dafny.Set({self.f35, self.f35})
        if (d_1926_v0_).isdisjoint(_dafny.Set({self.f29, self.f35})):
            d_1927_v1_: _dafny.Map
            d_1927_v1_ = _dafny.Map({self.f35: (self).f30})
            if (((d_1927_v1_)[self.f31] if (self.f31) in (d_1927_v1_) else (self).f34)) <= ((self).f30):
                d_1928_v2_: _dafny.Array
                def lambda78_(d_1929_i0_):
                    return self.f36

                init50_ = lambda78_
                nw338_ = _dafny.Array(None, 1)
                for i0_50_ in range(nw338_.length(0)):
                    nw338_[i0_50_] = init50_(i0_50_)
                d_1928_v2_ = nw338_
                d_1930_v3_: C0
                nw339_ = C0()
                nw339_.ctor__(d_1928_v2_)
                d_1930_v3_ = nw339_
                d_1931_v4_: D1
                d_1931_v4_ = D1_DC3(self.f35, 777, d_1930_v3_, (0) - ((self).f34))
                d_1931_v4_ = d_1931_v4_
                d_1932_v5_: _dafny.Seq
                d_1932_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyotbppo"))
                d_1933_v6_: D5
                d_1933_v6_ = D5_DC12(d_1932_v5_)
                d_1934_v7_: D0
                d_1934_v7_ = D0_DC0(self.f35)
                d_1935_v8_: D2
                d_1935_v8_ = D2_DC6((self).f30, self.f35, len(d_1932_v5_), not((d_1934_v7_).cf0), -57)
                d_1936_v9_: _dafny.MultiSet
                d_1936_v9_ = _dafny.MultiSet([self.f28, self.f35, self.f28])
                d_1937_v10_: _dafny.Array
                nw340_ = _dafny.Array(None, 10)
                nw340_[int(0)] = (self).f30
                nw340_[int(1)] = (self).f30
                nw340_[int(2)] = ((self).f30) * ((self).f34)
                nw340_[int(3)] = (len((d_1933_v6_).cf26)) + ((self).f38)
                nw340_[int(4)] = (self).f34
                nw340_[int(5)] = (d_1935_v8_).cf11
                nw340_[int(6)] = (self).f34
                nw340_[int(7)] = (d_1936_v9_).cardinality
                nw340_[int(8)] = (self).f34
                nw340_[int(9)] = (self).f38
                d_1937_v10_ = nw340_
                d_1938_v11_: _dafny.Set
                d_1938_v11_ = _dafny.Set({(self).f34, (self).f34})
                d_1939_v12_: _dafny.MultiSet
                d_1939_v12_ = _dafny.MultiSet([(self).f30])
                d_1940_v13_: _dafny.Array
                nw341_ = _dafny.Array(None, 16)
                nw341_[int(0)] = True
                nw341_[int(1)] = self.f31
                nw341_[int(2)] = self.f28
                nw341_[int(3)] = (self).fm16(self.f31, d_1934_v7_, d_1938_v11_, (self).f38, globalState)
                nw341_[int(4)] = self.f35
                nw341_[int(5)] = self.f31
                nw341_[int(6)] = self.f29
                nw341_[int(7)] = self.f35
                nw341_[int(8)] = self.f31
                nw341_[int(9)] = not(True)
                nw341_[int(10)] = self.f28
                nw341_[int(11)] = self.f28
                nw341_[int(12)] = (self).fm7((d_1939_v12_).set((self).f34, default__.abs((self).f38)), (self).f30, globalState)
                nw341_[int(13)] = self.f35
                nw341_[int(14)] = self.f29
                nw341_[int(15)] = self.f28
                d_1940_v13_ = nw341_
                d_1941_v14_: D3
                d_1941_v14_ = D3_DC9(self.f28, False, d_1940_v13_, self.f31)
                index262_ = default__.safeIndex(176, (d_1937_v10_).length(0))
                (d_1937_v10_)[index262_] = (0) - (((self).f30 if (d_1941_v14_).cf21 else (self).f30))
                index263_ = default__.safeIndex(176, (d_1937_v10_).length(0))
                (d_1937_v10_)[index263_] = (self).f38
                d_1942_v15_: _dafny.Map
                d_1942_v15_ = _dafny.Map({self.f36: self.f35})
                d_1942_v15_ = (d_1942_v15_).set(self.f36, self.f35)
                d_1943_v16_: _dafny.Seq
                d_1943_v16_ = _dafny.SeqWithoutIsStrInference([((self).f34) - ((d_1937_v10_)[default__.safeIndex(176, (d_1937_v10_).length(0))])])
                (globalState).f13 = (d_1943_v16_)[default__.safeIndex((self).f34, len(d_1943_v16_))]
                d_1944_v18_: _dafny.Seq
                def iife102_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(144, 351):
                        d_1945_v17_: int = compr_52_
                        if ((144) <= (d_1945_v17_)) and ((d_1945_v17_) < (351)):
                            coll52_ = coll52_.union(_dafny.Set([(d_1945_v17_) - ((self).f38)]))
                    return _dafny.Set(coll52_)
                d_1944_v18_ = _dafny.SeqWithoutIsStrInference([iife102_()
                ])
                d_1946_v21_: _dafny.Map
                def iife103_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in (_dafny.Map({len(d_1932_v5_): self.f28})).keys.Elements:
                        d_1947_v20_: int = compr_53_
                        if (d_1947_v20_) in (_dafny.Map({len(d_1932_v5_): self.f28})):
                            coll53_ = coll53_.union(_dafny.Set([(d_1947_v20_) + (len(_dafny.SeqWithoutIsStrInference([454, 198])))]))
                    return _dafny.Set(coll53_)
                d_1946_v21_ = _dafny.Map({(self).f30: iife103_()
                })
                d_1948_v22_: _dafny.Set
                d_1948_v22_ = _dafny.Set({_dafny.Set({(self).f38}), ((d_1946_v21_)[(self).f38] if ((self).f38) in (d_1946_v21_) else (d_1944_v18_)[default__.safeIndex((self).f38, len(d_1944_v18_))]), d_1938_v11_, d_1938_v11_, _dafny.Set({-489, (self).f38})})
                index264_ = default__.safeIndex(176, (d_1937_v10_).length(0))
                def iife104_():
                    coll54_ = _dafny.Set()
                    compr_54_: _dafny.Set
                    for compr_54_ in (d_1944_v18_).Elements:
                        d_1949_v19_: _dafny.Set = compr_54_
                        if (d_1949_v19_) in (d_1944_v18_):
                            coll54_ = coll54_.union(_dafny.Set([d_1949_v19_]))
                    return _dafny.Set(coll54_)
                (d_1937_v10_)[index264_] = len(_dafny.SeqWithoutIsStrInference([(d_1948_v22_).issubset(iife104_()
                ), (self.f31 if self.f28 else False), self.f31]))
            elif True:
                d_1950_v23_: _dafny.Set
                d_1950_v23_ = _dafny.Set({(self).f30})
                (globalState).f1 = (self).fm16(((self).f38) >= ((self).f34), (default__.fm18(self.f31, self.f29, globalState) if self.f29 else D0_DC0(self.f35)), d_1950_v23_, (self).f38, globalState)
                d_1951_v24_: _dafny.Seq
                d_1951_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srednb"))
                d_1952_v25_: _dafny.Map
                d_1952_v25_ = _dafny.Map({(d_1951_v24_) == (d_1951_v24_): not((self.f35 if self.f31 else self.f28))})
                d_1952_v25_ = (d_1952_v25_).set(self.f31, self.f29)
                (self).f36 = self.f36
                d_1953_v26_: _dafny.Map
                d_1953_v26_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([self.f36 for d_1954_i1_ in range(default__.abs(-92))])) if True else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "im")))): not(self.f35)})
                r0 = not(((d_1953_v26_)[(self).f38] if ((self).f38) in (d_1953_v26_) else self.f31))
                d_1955_v27_: _dafny.Array
                nw342_ = _dafny.Array(False, 12)
                d_1955_v27_ = nw342_
                index265_ = default__.safeIndex(556, (d_1955_v27_).length(0))
                (d_1955_v27_)[index265_] = self.f29
                index266_ = default__.safeIndex(556, (d_1955_v27_).length(0))
                (d_1955_v27_)[index266_] = self.f28
            d_1956_v28_: _dafny.Seq
            d_1956_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bd"))
            d_1957_v29_: _dafny.Array
            nw343_ = _dafny.Array(None, 2)
            nw343_[int(0)] = d_1956_v28_
            nw343_[int(1)] = (d_1956_v28_) + (d_1956_v28_)
            d_1957_v29_ = nw343_
            index267_ = default__.safeIndex(798, (d_1957_v29_).length(0))
            (d_1957_v29_)[index267_] = d_1956_v28_
            d_1958_v30_: _dafny.Array
            nw344_ = _dafny.Array(None, 16)
            d_1958_v30_ = nw344_
            d_1959_v31_: _dafny.Array
            def lambda79_(d_1960_i2_):
                return self.f36

            init51_ = lambda79_
            nw345_ = _dafny.Array(None, 22)
            for i0_51_ in range(nw345_.length(0)):
                nw345_[i0_51_] = init51_(i0_51_)
            d_1959_v31_ = nw345_
            d_1961_v32_: C0
            nw346_ = C0()
            nw346_.ctor__(d_1959_v31_)
            d_1961_v32_ = nw346_
            index268_ = default__.safeIndex(900, (d_1958_v30_).length(0))
            (d_1958_v30_)[index268_] = d_1961_v32_
            d_1962_v33_: _dafny.Seq
            d_1962_v33_ = _dafny.SeqWithoutIsStrInference([(self).f38, (self).f30])
            d_1963_v34_: T2
            nw347_ = C5()
            nw347_.ctor__(self.f28, (self).f34, (self).f30, self.f31, (self).f32, (self).f33, (d_1962_v33_)[default__.safeIndex((self).f30, len(d_1962_v33_))], self.f28, self.f35, not(self.f35))
            d_1963_v34_ = nw347_
            d_1964_v35_: _dafny.Map
            d_1964_v35_ = _dafny.Map({True: d_1963_v34_})
            d_1965_v36_: _dafny.Map
            d_1965_v36_ = _dafny.Map({len(d_1956_v28_): default__.fm46((self).f34, (d_1963_v34_).f30, self.f35, globalState)})
            index269_ = default__.safeIndex(798, (d_1957_v29_).length(0))
            index270_ = default__.safeIndex(900, (d_1958_v30_).length(0))
            rhs285_ = (self).fm3(self.f28, self.f28, globalState)
            rhs286_ = (len(d_1964_v35_)) * (((self).f38) * (((d_1965_v36_)[(d_1962_v33_)[default__.safeIndex((d_1963_v34_).f30, len(d_1962_v33_))]] if ((d_1962_v33_)[default__.safeIndex((d_1963_v34_).f30, len(d_1962_v33_))]) in (d_1965_v36_) else (d_1963_v34_).f30)))
            rhs287_ = (d_1963_v34_).f30
            rhs288_ = d_1961_v32_
            lhs224_ = d_1957_v29_
            lhs225_ = default__.safeIndex(798, (d_1957_v29_).length(0))
            lhs226_ = globalState
            lhs227_ = globalState
            lhs228_ = d_1958_v30_
            lhs229_ = default__.safeIndex(900, (d_1958_v30_).length(0))
            lhs224_[lhs225_] = rhs285_
            lhs226_.f13 = rhs286_
            lhs227_.f13 = rhs287_
            lhs228_[lhs229_] = rhs288_
            r1 = ((620) - ((self).f38)) * ((d_1963_v34_).f30)
            d_1966_v37_: _dafny.Array
            nw348_ = _dafny.Array(False, 19)
            d_1966_v37_ = nw348_
            (globalState).f21 = d_1966_v37_
            d_1967_v38_: _dafny.Map
            d_1967_v38_ = _dafny.Map({D0_DC1(_dafny.MultiSet([self.f31, self.f28, d_1963_v34_.f29])): self.f29})
            d_1968_v39_: D8
            d_1968_v39_ = D8_DC21(d_1967_v38_)
            d_1969_v40_: D8
            d_1969_v40_ = D8_DC25(d_1968_v39_)
            rhs289_ = d_1966_v37_
            rhs290_ = d_1966_v37_
            rhs291_ = d_1969_v40_
            rhs292_ = d_1963_v34_.f28
            lhs230_ = self
            d_1966_v37_ = rhs289_
            d_1966_v37_ = rhs290_
            d_1969_v40_ = rhs291_
            lhs230_.f28 = rhs292_
        elif True:
            d_1970_v41_: _dafny.MultiSet
            d_1970_v41_ = _dafny.MultiSet([(self).f38, (self).f38])
            d_1971_v42_: _dafny.Array
            nw349_ = _dafny.Array(None, 4)
            nw349_[int(0)] = (d_1970_v41_).cardinality
            nw349_[int(1)] = (self).f34
            nw349_[int(2)] = (self).f34
            nw349_[int(3)] = (self).f34
            d_1971_v42_ = nw349_
            index271_ = default__.safeIndex(436, (d_1971_v42_).length(0))
            (d_1971_v42_)[index271_] = (self).f30
            index272_ = default__.safeIndex(436, (d_1971_v42_).length(0))
            (d_1971_v42_)[index272_] = (self).f34
            (globalState).f1 = False
            (globalState).f26 = self.f35
            d_1972_v43_: C8
            nw350_ = C8()
            nw350_.ctor__(self.f35, self.f36, (self).f34, self.f28, (self).f32, (self).f33, (self).f30, self.f31, self.f29, self.f31)
            d_1972_v43_ = nw350_
            d_1973_v44_: _dafny.Map
            d_1973_v44_ = _dafny.Map({d_1972_v43_: (self).f30})
            d_1974_v45_: _dafny.Set
            d_1974_v45_ = _dafny.Set({(0) - ((self).f34), (self).f30, default__.fm24(self.f35, (d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))], globalState), (self).f38, len(d_1973_v44_)})
            if ((0) - (((self).f34) + ((self).f34))) >= (len(d_1974_v45_)):
                index273_ = default__.safeIndex(436, (d_1971_v42_).length(0))
                rhs293_ = ((self).f38) == ((self).f34)
                rhs294_ = (self).f38
                lhs231_ = globalState
                lhs232_ = d_1971_v42_
                lhs233_ = default__.safeIndex(436, (d_1971_v42_).length(0))
                lhs231_.f1 = rhs293_
                lhs232_[lhs233_] = rhs294_
                index274_ = default__.safeIndex(436, (d_1971_v42_).length(0))
                (d_1971_v42_)[index274_] = (self).f34
                d_1975_v46_: _dafny.Array
                def lambda80_(d_1976_i3_):
                    return self.f28

                init52_ = lambda80_
                nw351_ = _dafny.Array(None, 18)
                for i0_52_ in range(nw351_.length(0)):
                    nw351_[i0_52_] = init52_(i0_52_)
                d_1975_v46_ = nw351_
                index275_ = default__.safeIndex(426, (d_1975_v46_).length(0))
                (d_1975_v46_)[index275_] = self.f35
                d_1977_v47_: _dafny.Set
                d_1977_v47_ = _dafny.Set({_dafny.Set({(self).f30})})
                d_1978_v48_: T0
                nw352_ = C7()
                nw352_.ctor__((self).f34, d_1970_v41_, (self).f38, self.f35, self.f35, d_1972_v43_.f40)
                d_1978_v48_ = nw352_
                d_1979_v49_: _dafny.Map
                d_1979_v49_ = _dafny.Map({len(((self).f32 if d_1978_v48_.f29 else ((self).f32).set(default__.safeIndex((self).f34, len((self).f32)), d_1972_v43_.f40))): d_1978_v48_})
                index276_ = default__.safeIndex(426, (d_1975_v46_).length(0))
                rhs295_ = self.f29
                rhs296_ = (d_1977_v47_).intersection(d_1977_v47_)
                rhs297_ = ((d_1979_v49_)[default__.safeDivisionInt((d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okb"))))] if (default__.safeDivisionInt((d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okb"))))) in (d_1979_v49_) else d_1978_v48_)
                lhs234_ = d_1975_v46_
                lhs235_ = default__.safeIndex(426, (d_1975_v46_).length(0))
                lhs234_[lhs235_] = rhs295_
                d_1977_v47_ = rhs296_
                d_1978_v48_ = rhs297_
                d_1980_v50_: C1
                nw353_ = C1()
                nw353_.ctor__(self.f31, True, self.f29)
                d_1980_v50_ = nw353_
                d_1981_v51_: _dafny.Map
                d_1981_v51_ = _dafny.Map({self.f36: _dafny.Map({d_1980_v50_: default__.fm39(d_1972_v43_.f40, d_1970_v41_, ((self).f32)[default__.safeIndex((self).f38, len((self).f32))], globalState)})})
                d_1982_v52_: _dafny.Map
                d_1982_v52_ = _dafny.Map({d_1980_v50_: (self).f38})
                d_1981_v51_ = (d_1981_v51_).set(self.f36, d_1982_v52_)
                d_1983_v53_: _dafny.Seq
                d_1983_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhctyq"))
                d_1984_v54_: _dafny.Array
                nw354_ = _dafny.Array(None, 15)
                nw354_[int(0)] = self.f36
                nw354_[int(1)] = self.f36
                nw354_[int(2)] = self.f36
                nw354_[int(3)] = _dafny.CodePoint('h')
                nw354_[int(4)] = self.f36
                nw354_[int(5)] = default__.fm34((self).f38, (self).f38, len(d_1983_v53_), self.f28, globalState)
                nw354_[int(6)] = self.f36
                nw354_[int(7)] = default__.fm41(self.f29, self.f36, self.f29, (self).f38, globalState)
                nw354_[int(8)] = (d_1983_v53_)[default__.safeIndex((self).f30, len(d_1983_v53_))]
                nw354_[int(9)] = self.f36
                nw354_[int(10)] = _dafny.CodePoint('v')
                nw354_[int(11)] = _dafny.CodePoint('a')
                nw354_[int(12)] = (d_1983_v53_)[default__.safeIndex((d_1970_v41_).cardinality, len(d_1983_v53_))]
                nw354_[int(13)] = self.f36
                nw354_[int(14)] = self.f36
                d_1984_v54_ = nw354_
                index277_ = default__.safeIndex(164, (d_1984_v54_).length(0))
                (d_1984_v54_)[index277_] = self.f36
                index278_ = default__.safeIndex(164, (d_1984_v54_).length(0))
                rhs298_ = self.f36
                rhs299_ = self.f36
                rhs300_ = not (self.f28) or ((d_1975_v46_)[default__.safeIndex(426, (d_1975_v46_).length(0))])
                lhs236_ = self
                lhs237_ = d_1984_v54_
                lhs238_ = default__.safeIndex(164, (d_1984_v54_).length(0))
                lhs239_ = globalState
                lhs236_.f36 = rhs298_
                lhs237_[lhs238_] = rhs299_
                lhs239_.f26 = rhs300_
            elif True:
                d_1985_v55_: _dafny.Map
                d_1985_v55_ = _dafny.Map({True: self.f35})
                d_1986_v56_: _dafny.Set
                d_1986_v56_ = _dafny.Set({self.f36, default__.fm41(self.f31, self.f36, self.f35, (self).f30, globalState)})
                d_1987_v57_: D0
                d_1987_v57_ = D0_DC0(self.f28)
                d_1988_v58_: _dafny.Map
                d_1988_v58_ = _dafny.Map({(d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))]: self.f31})
                d_1989_v59_: _dafny.Map
                d_1989_v59_ = _dafny.Map({(d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))]: default__.fm1(d_1972_v43_.f40, len(d_1988_v58_), globalState)})
                d_1990_v60_: _dafny.Array
                nw355_ = _dafny.Array(None, 26)
                nw355_[int(0)] = d_1972_v43_.f40
                nw355_[int(1)] = self.f28
                nw355_[int(2)] = not(self.f35)
                nw355_[int(3)] = self.f35
                nw355_[int(4)] = self.f35
                nw355_[int(5)] = self.f31
                nw355_[int(6)] = ((d_1985_v55_)[self.f35] if (self.f35) in (d_1985_v55_) else True)
                nw355_[int(7)] = d_1972_v43_.f40
                nw355_[int(8)] = not(d_1972_v43_.f40)
                nw355_[int(9)] = ((d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))]) in (d_1970_v41_)
                nw355_[int(10)] = (_dafny.CodePoint('s')) in (d_1986_v56_)
                nw355_[int(11)] = d_1972_v43_.f40
                nw355_[int(12)] = self.f35
                nw355_[int(13)] = self.f31
                nw355_[int(14)] = self.f29
                nw355_[int(15)] = (self).fm16(not(self.f28), d_1987_v57_, (self).fm4(self.f31, _dafny.Set({(self).f30}), (self).f30, self.f29, globalState), len((self).f32), globalState)
                nw355_[int(16)] = d_1972_v43_.f40
                nw355_[int(17)] = self.f35
                nw355_[int(18)] = self.f28
                nw355_[int(19)] = ((self).f34) == (848)
                nw355_[int(20)] = (_dafny.SeqWithoutIsStrInference([self.f36 for d_1991_i4_ in range(default__.abs(357))])) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdyxsb")))
                nw355_[int(21)] = self.f29
                nw355_[int(22)] = (self.f31) or (self.f35)
                nw355_[int(23)] = ((self).f30) in (d_1988_v58_)
                nw355_[int(24)] = (False) or (not(self.f35))
                nw355_[int(25)] = not (self.f31) or (self.f29)
                d_1990_v60_ = nw355_
                index279_ = default__.safeIndex(225, (d_1990_v60_).length(0))
                (d_1990_v60_)[index279_] = (self.f35 if True else True)
                d_1992_v61_: _dafny.MultiSet
                d_1992_v61_ = _dafny.MultiSet([self.f31])
                index280_ = default__.safeIndex(225, (d_1990_v60_).length(0))
                (d_1990_v60_)[index280_] = (self.f31 if (d_1992_v61_).ispropersubset(d_1992_v61_) else self.f29)
                index281_ = default__.safeIndex(436, (d_1971_v42_).length(0))
                (d_1971_v42_)[index281_] = (self).f34
                (self).f36 = self.f36
                (globalState).f12 = ((d_1985_v55_) | (_dafny.Map({self.f31: d_1972_v43_.f40}))) == ((d_1985_v55_) | (_dafny.Map({True: d_1972_v43_.f40})))
                (self).f36 = self.f36
            index282_ = default__.safeIndex(436, (d_1971_v42_).length(0))
            rhs301_ = not((not(self.f35)) or (self.f28))
            rhs302_ = (d_1971_v42_)[default__.safeIndex(436, (d_1971_v42_).length(0))]
            rhs303_ = d_1972_v43_.f40
            lhs240_ = d_1972_v43_
            lhs241_ = d_1971_v42_
            lhs242_ = default__.safeIndex(436, (d_1971_v42_).length(0))
            lhs243_ = globalState
            lhs240_.f40 = rhs301_
            lhs241_[lhs242_] = rhs302_
            lhs243_.f12 = rhs303_
        d_1993_v62_: _dafny.Array
        def lambda81_(d_1994_i5_):
            return (False if self.f35 else self.f28)

        init53_ = lambda81_
        nw356_ = _dafny.Array(None, 23)
        for i0_53_ in range(nw356_.length(0)):
            nw356_[i0_53_] = init53_(i0_53_)
        d_1993_v62_ = nw356_
        index283_ = default__.safeIndex(22, (d_1993_v62_).length(0))
        (d_1993_v62_)[index283_] = self.f28
        d_1995_v63_: _dafny.Seq
        d_1995_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxqix"))
        d_1996_v64_: _dafny.Seq
        d_1996_v64_ = _dafny.SeqWithoutIsStrInference([d_1995_v63_, d_1995_v63_])
        index284_ = default__.safeIndex(22, (d_1993_v62_).length(0))
        (d_1993_v62_)[index284_] = (d_1996_v64_) == (d_1996_v64_)
        d_1997_v65_: _dafny.Array
        nw357_ = _dafny.Array(int(0), 10)
        d_1997_v65_ = nw357_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1997_v65_).length(0)):
            d_1998_i6_: int = guard_loop_4_
            if (True) and (((0) <= (d_1998_i6_)) and ((d_1998_i6_) < ((d_1997_v65_).length(0)))):
                (d_1997_v65_)[(d_1998_i6_)] = (d_1998_i6_) - ((self).f30)
        r1 = len(d_1926_v0_)
        r1 = -371
        d_1999_v66_: _dafny.Map
        d_1999_v66_ = _dafny.Map({(self).f34: len((self).f32)})
        if (d_1999_v66_) == ((d_1999_v66_) | (d_1999_v66_)):
            r0 = self.f28
            d_2000_v67_: _dafny.Array
            def lambda82_(d_2001_i7_):
                return self.f36

            init54_ = lambda82_
            nw358_ = _dafny.Array(None, 17)
            for i0_54_ in range(nw358_.length(0)):
                nw358_[i0_54_] = init54_(i0_54_)
            d_2000_v67_ = nw358_
            d_2002_v68_: C0
            nw359_ = C0()
            nw359_.ctor__(d_2000_v67_)
            d_2002_v68_ = nw359_
            d_2003_v69_: D1
            d_2003_v69_ = D1_DC3(self.f29, len(d_1995_v63_), d_2002_v68_, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xevgdv")))))
            d_2004_v70_: C4
            nw360_ = C4()
            nw360_.ctor__((self).f34, (self).fm6(globalState), (_dafny.SeqWithoutIsStrInference([self.f35])) + ((self).f32), (_dafny.SeqWithoutIsStrInference([(self).f32])) + (_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])), default__.safeModuloInt((self).f34, (self).f34), (d_2003_v69_).cf3, self.f35, self.f31)
            d_2004_v70_ = nw360_
            rhs304_ = d_2004_v70_
            rhs305_ = self.f31
            rhs306_ = d_1997_v65_
            lhs244_ = self
            d_2004_v70_ = rhs304_
            lhs244_.f31 = rhs305_
            d_1997_v65_ = rhs306_
            d_2005_v71_: _dafny.Map
            d_2005_v71_ = _dafny.Map({(self).f38: d_1997_v65_})
            d_1997_v65_ = ((d_2005_v71_)[(self).f38] if ((self).f38) in (d_2005_v71_) else d_1997_v65_)
            (self).f28 = (d_1926_v0_).ispropersubset(d_1926_v0_)
            d_2006_v72_: _dafny.Array
            nw361_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_2006_v72_ = nw361_
            index285_ = default__.safeIndex(750, (d_2006_v72_).length(0))
            (d_2006_v72_)[index285_] = d_1993_v62_
            index286_ = default__.safeIndex(750, (d_2006_v72_).length(0))
            (d_2006_v72_)[index286_] = d_1993_v62_
        elif True:
            d_2007_v73_: _dafny.Map
            d_2007_v73_ = _dafny.Map({(d_1993_v62_)[default__.safeIndex(22, (d_1993_v62_).length(0))]: (0) - ((self).f34)})
            d_2008_v74_: _dafny.Seq
            d_2008_v74_ = _dafny.SeqWithoutIsStrInference([len(d_2007_v73_), default__.safeDivisionInt((self).f34, (_dafny.MultiSet([d_1997_v65_])).cardinality), (self).f38])
            d_2009_v75_: _dafny.Seq
            d_2009_v75_ = _dafny.SeqWithoutIsStrInference([d_2008_v74_, d_2008_v74_])
            d_2008_v74_ = ((d_2008_v74_) + (_dafny.SeqWithoutIsStrInference([(self).f38 for d_2010_i8_ in range(default__.abs(832))]))) + ((d_2008_v74_) + ((d_2009_v75_)[default__.safeIndex((self).f30, len(d_2009_v75_))]))
            index287_ = default__.safeIndex(781, (d_1997_v65_).length(0))
            (d_1997_v65_)[index287_] = len(d_1926_v0_)
            index288_ = default__.safeIndex(781, (d_1997_v65_).length(0))
            (d_1997_v65_)[index288_] = default__.safeDivisionInt(-825, (self).f34)
            d_2011_v76_: _dafny.Map
            d_2011_v76_ = _dafny.Map({(d_2008_v74_)[default__.safeIndex((self).f30, len(d_2008_v74_))]: self.f29})
            d_2012_v77_: _dafny.Map
            d_2012_v77_ = _dafny.Map({d_1993_v62_: len(d_2011_v76_)})
            d_2013_v78_: _dafny.Map
            d_2013_v78_ = _dafny.Map({(d_1997_v65_)[default__.safeIndex(781, (d_1997_v65_).length(0))]: _dafny.SeqWithoutIsStrInference([self.f36 for d_2014_i9_ in range(default__.abs(836))])})
            d_2015_v79_: _dafny.Map
            d_2015_v79_ = _dafny.Map({self.f35: _dafny.CodePoint('h')})
            d_2016_v80_: _dafny.Array
            nw362_ = _dafny.Array(int(0), 23)
            d_2016_v80_ = nw362_
            d_2017_v81_: C5
            nw363_ = C5()
            nw363_.ctor__(((self).f34) == (793), len(d_2012_v77_), len(((d_2013_v78_)[(self).f38] if ((self).f38) in (d_2013_v78_) else d_1995_v63_)), (d_1993_v62_)[default__.safeIndex(22, (d_1993_v62_).length(0))], (self).f32, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f31, self.f29]), ((self).f32).set(default__.safeIndex(len(d_2015_v79_), len((self).f32)), self.f28), (self).f32]), len((d_2008_v74_).set(default__.safeIndex(len(_dafny.Set({d_2016_v80_})), len(d_2008_v74_)), (d_2008_v74_)[default__.safeIndex((self).f34, len(d_2008_v74_))])), (len((self).f32)) < ((self).f30), self.f28, ((self).f32)[default__.safeIndex((self).f34, len((self).f32))])
            d_2017_v81_ = nw363_
            d_2018_v82_: _dafny.MultiSet
            d_2018_v82_ = _dafny.MultiSet([(self).f30, len(d_1995_v63_)])
            d_2019_v83_: _dafny.MultiSet
            d_2019_v83_ = _dafny.MultiSet([290, (d_2018_v82_).cardinality])
            (globalState).f12 = (self).fm7(d_2018_v82_, (0) - (len((d_2008_v74_) + (d_2008_v74_))), globalState)
            d_2020_v84_: D8
            d_2020_v84_ = D8_DC22((self).f34, (self).f30)
            d_2021_v85_: D2
            d_2021_v85_ = D2_DC6((d_2020_v84_).cf41, self.f28, (d_2017_v81_).f48, (d_2017_v81_).f47, -48)
            pat_let_tv18_ = d_2011_v76_
            pat_let_tv19_ = d_2011_v76_
            def iife105_(_pat_let25_0):
                def iife106_(d_2022_dt__update__tmp_h0_):
                    def iife107_(_pat_let26_0):
                        def iife108_(d_2023_dt__update_hcf9_h0_):
                            def iife109_(_pat_let27_0):
                                def iife110_(d_2024_dt__update_hcf12_h0_):
                                    def iife111_(_pat_let28_0):
                                        def iife112_(d_2025_dt__update_hcf13_h0_):
                                            return D2_DC6(d_2023_dt__update_hcf9_h0_, (d_2022_dt__update__tmp_h0_).cf10, (d_2022_dt__update__tmp_h0_).cf11, d_2024_dt__update_hcf12_h0_, d_2025_dt__update_hcf13_h0_)
                                        return iife112_(_pat_let28_0)
                                    return iife111_(len((pat_let_tv18_) | (pat_let_tv19_)))
                                return iife110_(_pat_let27_0)
                            return iife109_(not(self.f29))
                        return iife108_(_pat_let26_0)
                    return iife107_((self).f38)
                return iife106_(_pat_let25_0)
            d_2021_v85_ = iife105_(d_2021_v85_)
        r0 = self.f31
        r1 = (self).f30
        d_2026_v86_: _dafny.MultiSet
        d_2026_v86_ = _dafny.MultiSet([(self).f34, (0) - ((self).f38)])
        d_2027_v87_: D14
        d_2027_v87_ = D14_DC37(default__.fm46((self).f34, (self).f38, self.f29, globalState), (self).fm7(d_2026_v86_, (self).f30, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfybmsjun"))))
        pat_let_tv20_ = d_1995_v63_
        pat_let_tv21_ = d_1995_v63_
        def lambda83_(source33_):
            if source33_.is_DC37:
                d_2028___mcc_h0_ = source33_.cf68
                d_2029___mcc_h1_ = source33_.cf69
                d_2030___mcc_h2_ = source33_.cf70
                d_2031_cf70_ = d_2030___mcc_h2_
                d_2032_cf69_ = d_2029___mcc_h1_
                d_2033_cf68_ = d_2028___mcc_h0_
                return pat_let_tv20_
            elif True:
                d_2034___mcc_h3_ = source33_.cf67
                d_2035_cf67_ = d_2034___mcc_h3_
                return pat_let_tv21_

        r2 = lambda83_(d_2027_v87_)
        return r0, r1, r2

    def m4(self, globalState):
        r0: _dafny.Map = _dafny.Map({})
        (self).f36 = self.f36
        (globalState).f13 = (0) - ((self).f38)
        (self).f28 = self.f29
        d_2036_v0_: D0
        d_2036_v0_ = D0_DC0(self.f35)
        if (d_2036_v0_).cf0:
            d_2037_v1_: _dafny.Seq
            d_2037_v1_ = _dafny.SeqWithoutIsStrInference([(self).f34, (self).f34])
            d_2038_v2_: _dafny.Set
            d_2038_v2_ = _dafny.Set({d_2037_v1_})
            d_2039_v3_: D16
            d_2039_v3_ = D16_DC39(d_2038_v2_)
            d_2040_v4_: D16
            d_2040_v4_ = D16_DC41(d_2039_v3_)
            d_2041_v5_: D16
            d_2041_v5_ = D16_DC41(D16_DC41(d_2040_v4_))
            d_2042_v6_: _dafny.Array
            nw364_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
            d_2042_v6_ = nw364_
            d_2043_v7_: D20
            d_2043_v7_ = D20_DC47(d_2042_v6_)
            pat_let_tv22_ = d_2042_v6_
            def iife113_(_pat_let29_0):
                def iife114_(d_2044_dt__update__tmp_h0_):
                    def iife115_(_pat_let30_0):
                        def iife116_(d_2045_dt__update_hcf81_h0_):
                            return D20_DC47(d_2045_dt__update_hcf81_h0_)
                        return iife116_(_pat_let30_0)
                    return iife115_(pat_let_tv22_)
                return iife114_(_pat_let29_0)
            rhs307_ = d_2041_v5_
            rhs308_ = (self).f38
            rhs309_ = ((self).f32)[default__.safeIndex((self).f34, len((self).f32))]
            rhs310_ = (iife113_(d_2043_v7_)).cf81
            rhs311_ = (self).f30
            lhs245_ = globalState
            lhs246_ = globalState
            lhs247_ = globalState
            d_2041_v5_ = rhs307_
            lhs245_.f13 = rhs308_
            lhs246_.f1 = rhs309_
            d_2042_v6_ = rhs310_
            lhs247_.f13 = rhs311_
            d_2046_v8_: _dafny.MultiSet
            d_2046_v8_ = _dafny.MultiSet([False, self.f28])
            if (_dafny.MultiSet([self.f31, self.f35, self.f35])).ispropersubset(d_2046_v8_):
                (globalState).f13 = (self).f34
                d_2047_v9_: _dafny.Map
                d_2047_v9_ = _dafny.Map({self.f36: len(((self).f33)[default__.safeIndex((0) - ((self).f30), len((self).f33))])})
                d_2048_v10_: _dafny.Seq
                d_2048_v10_ = _dafny.SeqWithoutIsStrInference([d_2047_v9_])
                (self).f28 = ((self).f38) >= (len(d_2048_v10_))
                d_2049_v11_: D20
                d_2049_v11_ = D20_DC49(self.f29, self.f36)
                d_2050_v12_: C1
                nw365_ = C1()
                nw365_.ctor__(self.f29, (self.f28) in ((self).f32), not((d_2049_v11_).cf85))
                d_2050_v12_ = nw365_
                (globalState).f1 = self.f35
                d_2051_v13_: _dafny.Array
                nw366_ = _dafny.Array(False, 10)
                d_2051_v13_ = nw366_
                index289_ = default__.safeIndex(702, (d_2051_v13_).length(0))
                (d_2051_v13_)[index289_] = True
                index290_ = default__.safeIndex(312, (d_2051_v13_).length(0))
                (d_2051_v13_)[index290_] = self.f35
                d_2052_v14_: _dafny.Seq
                d_2052_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wc"))
                index291_ = default__.safeIndex(702, (d_2051_v13_).length(0))
                index292_ = default__.safeIndex(312, (d_2051_v13_).length(0))
                rhs312_ = ((self).f34) <= ((0) - (((0) - (len(d_2037_v1_))) - ((0) - ((self).f38))))
                rhs313_ = self.f31
                rhs314_ = (default__.safeDivisionInt((self).f38, len(d_2037_v1_))) * ((self).f38)
                rhs315_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                rhs316_ = (self).f30
                lhs248_ = d_2051_v13_
                lhs249_ = default__.safeIndex(702, (d_2051_v13_).length(0))
                lhs250_ = d_2051_v13_
                lhs251_ = default__.safeIndex(312, (d_2051_v13_).length(0))
                lhs252_ = globalState
                lhs253_ = globalState
                lhs248_[lhs249_] = rhs312_
                lhs250_[lhs251_] = rhs313_
                lhs252_.f13 = rhs314_
                d_2052_v14_ = rhs315_
                lhs253_.f13 = rhs316_
            elif True:
                d_2053_v15_: _dafny.Array
                nw367_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_2053_v15_ = nw367_
                index293_ = default__.safeIndex(3, (d_2053_v15_).length(0))
                (d_2053_v15_)[index293_] = d_2046_v8_
                index294_ = default__.safeIndex(3, (d_2053_v15_).length(0))
                (d_2053_v15_)[index294_] = d_2046_v8_
                (self).f28 = self.f35
                d_2054_v16_: D8
                d_2054_v16_ = D8_DC24(self.f31, (self).f34, (self).f32, self.f35, d_2037_v1_)
                d_2055_v17_: D8
                d_2055_v17_ = D8_DC25(d_2054_v16_)
                d_2056_v18_: _dafny.Seq
                d_2056_v18_ = _dafny.SeqWithoutIsStrInference([d_2055_v17_, d_2055_v17_])
                d_2057_v19_: _dafny.MultiSet
                d_2057_v19_ = _dafny.MultiSet([d_2056_v18_, d_2056_v18_])
                (globalState).f26 = ((d_2057_v19_) - (d_2057_v19_)).ispropersubset(d_2057_v19_)
                d_2058_v20_: _dafny.Array
                nw368_ = _dafny.Array(None, 29)
                d_2058_v20_ = nw368_
                d_2059_v21_: _dafny.MultiSet
                d_2059_v21_ = _dafny.MultiSet([d_2058_v20_])
                (globalState).f26 = not((d_2059_v21_).ispropersubset(d_2059_v21_))
                d_2060_v22_: _dafny.MultiSet
                d_2060_v22_ = _dafny.MultiSet([(self).f34, (self).f30])
                d_2061_v23_: _dafny.Map
                d_2061_v23_ = _dafny.Map({(self).f38: default__.safeModuloInt((0) - (((d_2060_v22_).set(-797, default__.abs((self).f38))).cardinality), (self).f38)})
                d_2062_v24_: _dafny.Set
                d_2062_v24_ = _dafny.Set({self.f31})
                (globalState).f13 = ((d_2061_v23_)[len((d_2062_v24_).intersection(_dafny.Set({self.f35})))] if (len((d_2062_v24_).intersection(_dafny.Set({self.f35})))) in (d_2061_v23_) else ((self).f34) * ((self).f38))
            d_2063_v25_: _dafny.Array
            def lambda84_(d_2064_v1_):
                def lambda85_(d_2065_i0_):
                    return (_dafny.Map({(self).f34: 238})) | (_dafny.Map({(d_2064_v1_)[default__.safeIndex(-294, len(d_2064_v1_))]: (self).f30}))

                return lambda85_

            init55_ = lambda84_(d_2037_v1_)
            nw369_ = _dafny.Array(None, 13)
            for i0_55_ in range(nw369_.length(0)):
                nw369_[i0_55_] = init55_(i0_55_)
            d_2063_v25_ = nw369_
            d_2066_v26_: _dafny.Map
            d_2066_v26_ = _dafny.Map({713: (self).f34})
            index295_ = default__.safeIndex(646, (d_2063_v25_).length(0))
            (d_2063_v25_)[index295_] = (d_2066_v26_) | (d_2066_v26_)
            index296_ = default__.safeIndex(646, (d_2063_v25_).length(0))
            (d_2063_v25_)[index296_] = _dafny.Map({len((self).f32): (self).f38})
            d_2067_v27_: _dafny.MultiSet
            d_2067_v27_ = _dafny.MultiSet([(0) - ((self).f30)])
            d_2068_v28_: D14
            d_2068_v28_ = D14_DC37(((d_2067_v27_)[(self).f38] if ((self).f38) in (d_2067_v27_) else (self).f34), self.f35, default__.safeModuloInt(625, (self).f38))
            d_2068_v28_ = default__.fm58(self.f28, ((self).f33)[default__.safeIndex((self).f38, len((self).f33))], globalState)
            d_2069_v29_: _dafny.Set
            d_2069_v29_ = _dafny.Set({(self).f30})
            d_2070_v30_: int
            out69_: int
            out69_ = default__.m0(d_2037_v1_, self.f28, (self).f34, ((self).f38) in (d_2069_v29_), globalState)
            d_2070_v30_ = out69_
        elif True:
            nw370_ = _dafny.Array(False, 27)
            (globalState).f21 = nw370_
            if self.f35:
                d_2071_v31_: _dafny.Seq
                d_2071_v31_ = _dafny.SeqWithoutIsStrInference([(self).f34])
                d_2072_v32_: int
                out70_: int
                out70_ = default__.m0((d_2071_v31_) + (d_2071_v31_), self.f31, (self).f38, self.f35, globalState)
                d_2072_v32_ = out70_
                (globalState).f26 = False
                (self).f35 = not(self.f31)
                d_2073_v33_: _dafny.Map
                d_2073_v33_ = _dafny.Map({(self).f30: (self).f38})
                d_2074_v35_: _dafny.Map
                def iife117_():
                    coll55_ = _dafny.Set()
                    compr_55_: int
                    for compr_55_ in (d_2073_v33_).keys.Elements:
                        d_2075_v34_: int = compr_55_
                        if (d_2075_v34_) in (d_2073_v33_):
                            coll55_ = coll55_.union(_dafny.Set([default__.safeDivisionInt(d_2075_v34_, len(_dafny.SeqWithoutIsStrInference([True, not(True)])))]))
                    return _dafny.Set(coll55_)
                d_2074_v35_ = _dafny.Map({self.f28: iife117_()
                })
                d_2076_v36_: D2
                d_2076_v36_ = D2_DC6((0) - (len((d_2074_v35_ if self.f31 else d_2074_v35_))), ((self).f32)[default__.safeIndex((self).f34, len((self).f32))], (self).f34, self.f31, default__.safeDivisionInt(d_2072_v32_, 423))
                d_2077_v37_: _dafny.Set
                d_2077_v37_ = _dafny.Set({self.f31})
                d_2078_v38_: _dafny.Seq
                d_2078_v38_ = _dafny.SeqWithoutIsStrInference([d_2077_v37_, _dafny.Set({self.f28}), d_2077_v37_, default__.fm45((self).f38, _dafny.MultiSet((self).f32), self.f28, (0) - (d_2072_v32_), globalState), d_2077_v37_])
                d_2079_v39_: _dafny.Set
                d_2079_v39_ = _dafny.Set({(d_2078_v38_)[default__.safeIndex(len((self).f32), len(d_2078_v38_))]})
                d_2080_v40_: _dafny.Array
                def lambda86_(d_2081_i1_):
                    return self.f35

                init56_ = lambda86_
                nw371_ = _dafny.Array(None, 3)
                for i0_56_ in range(nw371_.length(0)):
                    nw371_[i0_56_] = init56_(i0_56_)
                d_2080_v40_ = nw371_
                d_2082_v41_: _dafny.Seq
                d_2082_v41_ = _dafny.SeqWithoutIsStrInference([d_2080_v40_])
                d_2083_v42_: _dafny.Seq
                d_2083_v42_ = _dafny.SeqWithoutIsStrInference([(d_2082_v41_)[default__.safeIndex(700, len(d_2082_v41_))], d_2080_v40_])
                d_2076_v36_ = D2_DC6(len(d_2079_v39_), self.f28, (len((self).f32) if not(self.f29) else 968), self.f29, len(d_2083_v42_))
                d_2084_v43_: _dafny.Map
                d_2084_v43_ = _dafny.Map({d_2071_v31_: True})
                d_2084_v43_ = ((d_2084_v43_).set(d_2071_v31_, self.f28)).set((_dafny.SeqWithoutIsStrInference([(self).f30 for d_2085_i2_ in range(default__.abs(991))])) + (d_2071_v31_), self.f35)
            elif True:
                d_2086_v44_: _dafny.Set
                d_2086_v44_ = _dafny.Set({self.f35})
                d_2087_v45_: _dafny.Map
                d_2087_v45_ = _dafny.Map({self.f36: (self).f30})
                d_2088_v46_: _dafny.Seq
                d_2088_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk"))
                d_2089_v47_: _dafny.Array
                nw372_ = _dafny.Array(None, 11)
                nw372_[int(0)] = (self).f38
                nw372_[int(1)] = (self).f34
                nw372_[int(2)] = (self).f38
                nw372_[int(3)] = len(d_2086_v44_)
                nw372_[int(4)] = len(d_2087_v45_)
                nw372_[int(5)] = (self).f38
                nw372_[int(6)] = len(d_2088_v46_)
                nw372_[int(7)] = (self).f38
                nw372_[int(8)] = len(d_2088_v46_)
                nw372_[int(9)] = (self).f30
                nw372_[int(10)] = (self).f38
                d_2089_v47_ = nw372_
                d_2090_v48_: _dafny.MultiSet
                d_2090_v48_ = _dafny.MultiSet([d_2089_v47_])
                d_2091_v50_: _dafny.Array
                nw373_ = _dafny.Array(None, 16)
                nw373_[int(0)] = (self).f34
                nw373_[int(1)] = ((d_2090_v48_).cardinality) - ((self).f38)
                nw373_[int(2)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyvg")))) + (len((self).f33))
                nw373_[int(3)] = 140
                nw373_[int(4)] = (self).f38
                nw373_[int(5)] = ((self).f30) * ((self).f30)
                nw373_[int(6)] = (self).f30
                nw373_[int(7)] = len((_dafny.SeqWithoutIsStrInference([self.f36 for d_2092_i3_ in range(default__.abs(816))])) + (d_2088_v46_))
                nw373_[int(8)] = (self).f34
                nw373_[int(9)] = (self).f30
                nw373_[int(10)] = (self).f38
                nw373_[int(11)] = len(d_2088_v46_)
                nw373_[int(12)] = (self).f34
                nw373_[int(13)] = (self).f30
                nw373_[int(14)] = default__.safeModuloInt((self).f30, (self).f30)
                def iife118_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in _dafny.IntegerRange(-644, 310):
                        d_2093_v49_: int = compr_56_
                        if ((-644) <= (d_2093_v49_)) and ((d_2093_v49_) < (310)):
                            coll56_[(d_2093_v49_) - ((self).f30)] = self.f31
                    return _dafny.Map(coll56_)
                nw373_[int(15)] = len(_dafny.Map({self.f29: len(iife118_()
                )}))
                d_2091_v50_ = nw373_
                index297_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                (d_2091_v50_)[index297_] = (self).f38
                d_2094_v51_: _dafny.Set
                d_2094_v51_ = _dafny.Set({(self).f30})
                d_2095_v52_: D21
                d_2095_v52_ = D21_DC51(default__.fm59(self.f35, self.f35, self.f28, globalState))
                d_2096_v53_: _dafny.Map
                d_2096_v53_ = _dafny.Map({(self).f30: d_2094_v51_})
                index298_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                def iife119_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in _dafny.IntegerRange(-564, 901):
                        d_2097_v54_: int = compr_57_
                        if ((-564) <= (d_2097_v54_)) and ((d_2097_v54_) < (901)):
                            coll57_ = coll57_.union(_dafny.Set([default__.safeModuloInt(d_2097_v54_, len(_dafny.Map({self.f29: self.f35})))]))
                    return _dafny.Set(coll57_)
                rhs317_ = ((self).f34) not in (d_2094_v51_)
                rhs318_ = False
                rhs319_ = default__.fm48(((d_2095_v52_).cf88) != ((d_2096_v53_).set((self).f38, iife119_()
                )), self.f35, (self).f38, globalState)
                rhs320_ = len(_dafny.SeqWithoutIsStrInference([(self).f38 for d_2098_i4_ in range(default__.abs(756))]))
                lhs254_ = globalState
                lhs255_ = globalState
                lhs256_ = self
                lhs257_ = d_2091_v50_
                lhs258_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                lhs254_.f12 = rhs317_
                lhs255_.f12 = rhs318_
                lhs256_.f36 = rhs319_
                lhs257_[lhs258_] = rhs320_
                (self).f36 = self.f36
                d_2099_v55_: _dafny.Array
                def lambda87_(d_2100_v46_):
                    def lambda88_(d_2101_i5_):
                        return d_2100_v46_

                    return lambda88_

                init57_ = lambda87_(d_2088_v46_)
                nw374_ = _dafny.Array(None, 3)
                for i0_57_ in range(nw374_.length(0)):
                    nw374_[i0_57_] = init57_(i0_57_)
                d_2099_v55_ = nw374_
                index299_ = default__.safeIndex(958, (d_2099_v55_).length(0))
                (d_2099_v55_)[index299_] = d_2088_v46_
                index300_ = default__.safeIndex(958, (d_2099_v55_).length(0))
                (d_2099_v55_)[index300_] = (self).fm3(not(not(self.f35)), False, globalState)
                d_2102_v56_: _dafny.Array
                def lambda89_(d_2103_i6_):
                    return _dafny.MultiSet([self.f29])

                init58_ = lambda89_
                nw375_ = _dafny.Array(None, 6)
                for i0_58_ in range(nw375_.length(0)):
                    nw375_[i0_58_] = init58_(i0_58_)
                d_2102_v56_ = nw375_
                d_2104_v57_: _dafny.MultiSet
                d_2104_v57_ = _dafny.MultiSet([self.f28, self.f31, self.f31])
                index301_ = default__.safeIndex(962, (d_2102_v56_).length(0))
                (d_2102_v56_)[index301_] = d_2104_v57_
                d_2105_v58_: _dafny.Array
                nw376_ = _dafny.Array(False, 25)
                d_2105_v58_ = nw376_
                index302_ = default__.safeIndex(323, (d_2105_v58_).length(0))
                (d_2105_v58_)[index302_] = self.f28
                index303_ = default__.safeIndex(962, (d_2102_v56_).length(0))
                index304_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                index305_ = default__.safeIndex(323, (d_2105_v58_).length(0))
                rhs321_ = (d_2099_v55_)[default__.safeIndex(958, (d_2099_v55_).length(0))]
                rhs322_ = (d_2104_v57_).set(self.f31, default__.abs((self).f38))
                rhs323_ = (939) < ((0) - ((self).f38))
                rhs324_ = (d_2091_v50_)[default__.safeIndex(832, (d_2091_v50_).length(0))]
                rhs325_ = not(not(self.f28))
                lhs259_ = d_2102_v56_
                lhs260_ = default__.safeIndex(962, (d_2102_v56_).length(0))
                lhs261_ = globalState
                lhs262_ = d_2091_v50_
                lhs263_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                lhs264_ = d_2105_v58_
                lhs265_ = default__.safeIndex(323, (d_2105_v58_).length(0))
                d_2088_v46_ = rhs321_
                lhs259_[lhs260_] = rhs322_
                lhs261_.f1 = rhs323_
                lhs262_[lhs263_] = rhs324_
                lhs264_[lhs265_] = rhs325_
                d_2106_v59_: _dafny.MultiSet
                d_2106_v59_ = _dafny.MultiSet([(self).f34])
                d_2107_v60_: D17
                d_2107_v60_ = D17_DC44(d_2106_v59_, (self).f34)
                d_2108_v61_: _dafny.Map
                d_2108_v61_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([self.f28])) < ((self).f32): self.f29})
                index306_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                def iife120_(_pat_let31_0):
                    def iife121_(d_2109_dt__update__tmp_h1_):
                        def iife122_(_pat_let32_0):
                            def iife123_(d_2110_dt__update_hcf78_h0_):
                                return D17_DC44((d_2109_dt__update__tmp_h1_).cf77, d_2110_dt__update_hcf78_h0_)
                            return iife123_(_pat_let32_0)
                        return iife122_(695)
                    return iife121_(_pat_let31_0)
                rhs326_ = iife120_(d_2107_v60_)
                rhs327_ = len(d_2108_v61_)
                lhs266_ = d_2091_v50_
                lhs267_ = default__.safeIndex(832, (d_2091_v50_).length(0))
                d_2107_v60_ = rhs326_
                lhs266_[lhs267_] = rhs327_
            d_2111_v62_: _dafny.Seq
            d_2111_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvpm"))
            d_2111_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnhqon"))
            d_2112_v63_: _dafny.Array
            def lambda90_(d_2113_v62_):
                def lambda91_(d_2114_i7_):
                    return (_dafny.MultiSet([self.f36])).ispropersubset(_dafny.MultiSet([(d_2113_v62_)[default__.safeIndex((self).f38, len(d_2113_v62_))]]))

                return lambda91_

            init59_ = lambda90_(d_2111_v62_)
            nw377_ = _dafny.Array(None, 26)
            for i0_59_ in range(nw377_.length(0)):
                nw377_[i0_59_] = init59_(i0_59_)
            d_2112_v63_ = nw377_
            index307_ = default__.safeIndex(361, (d_2112_v63_).length(0))
            (d_2112_v63_)[index307_] = self.f28
            index308_ = default__.safeIndex(96, (d_2112_v63_).length(0))
            (d_2112_v63_)[index308_] = self.f31
            d_2115_v64_: _dafny.MultiSet
            d_2115_v64_ = _dafny.MultiSet([self.f29, self.f29])
            d_2116_v65_: _dafny.Seq
            d_2116_v65_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f34), ((d_2115_v64_)[self.f31] if (self.f31) in (d_2115_v64_) else len(_dafny.Set({False}))), 431])
            index309_ = default__.safeIndex(361, (d_2112_v63_).length(0))
            index310_ = default__.safeIndex(96, (d_2112_v63_).length(0))
            rhs328_ = default__.safeDivisionInt(len(d_2116_v65_), (d_2116_v65_)[default__.safeIndex(984, len(d_2116_v65_))])
            rhs329_ = ((0) - ((self).f38)) >= ((self).f34)
            rhs330_ = self.f28
            lhs268_ = globalState
            lhs269_ = d_2112_v63_
            lhs270_ = default__.safeIndex(361, (d_2112_v63_).length(0))
            lhs271_ = d_2112_v63_
            lhs272_ = default__.safeIndex(96, (d_2112_v63_).length(0))
            lhs268_.f13 = rhs328_
            lhs269_[lhs270_] = rhs329_
            lhs271_[lhs272_] = rhs330_
            d_2117_v66_: _dafny.MultiSet
            d_2117_v66_ = _dafny.MultiSet([(self).f30])
            rhs331_ = (_dafny.MultiSet([(self).f38])) - (d_2117_v66_)
            rhs332_ = 224
            lhs273_ = globalState
            d_2117_v66_ = rhs331_
            lhs273_.f13 = rhs332_
        (globalState).f26 = self.f28
        d_2118_v67_: _dafny.Seq
        d_2118_v67_ = _dafny.SeqWithoutIsStrInference([self.f29])
        d_2119_v68_: _dafny.Array
        def lambda92_(d_2120_i8_):
            return (d_2120_i8_) + ((self).f30)

        init60_ = lambda92_
        nw378_ = _dafny.Array(None, 8)
        for i0_60_ in range(nw378_.length(0)):
            nw378_[i0_60_] = init60_(i0_60_)
        d_2119_v68_ = nw378_
        d_2121_v69_: _dafny.MultiSet
        d_2121_v69_ = _dafny.MultiSet([(0) - ((self).f34), (self).f38])
        index311_ = default__.safeIndex(129, (d_2119_v68_).length(0))
        (d_2119_v68_)[index311_] = (d_2121_v69_).cardinality
        index312_ = default__.safeIndex(129, (d_2119_v68_).length(0))
        rhs333_ = d_2118_v67_
        rhs334_ = (self).f30
        rhs335_ = (0) - ((self).f34)
        rhs336_ = (self).f38
        rhs337_ = (self).f30
        lhs274_ = globalState
        lhs275_ = globalState
        lhs276_ = globalState
        lhs277_ = d_2119_v68_
        lhs278_ = default__.safeIndex(129, (d_2119_v68_).length(0))
        d_2118_v67_ = rhs333_
        lhs274_.f13 = rhs334_
        lhs275_.f13 = rhs335_
        lhs276_.f13 = rhs336_
        lhs277_[lhs278_] = rhs337_
        d_2122_v70_: _dafny.Map
        d_2122_v70_ = _dafny.Map({(self).f38: self.f35})
        d_2123_v71_: _dafny.Map
        d_2123_v71_ = _dafny.Map({d_2122_v70_: (d_2119_v68_)[default__.safeIndex(129, (d_2119_v68_).length(0))]})
        d_2124_v72_: _dafny.Seq
        d_2124_v72_ = _dafny.SeqWithoutIsStrInference([d_2123_v71_, (d_2123_v71_) | (d_2123_v71_), d_2123_v71_, d_2123_v71_])
        r0 = (d_2124_v72_)[default__.safeIndex(338, len(d_2124_v72_))]
        return r0

    def m5(self, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_2125_v0_: _dafny.Array
        nw379_ = _dafny.Array(int(0), 24)
        d_2125_v0_ = nw379_
        d_2126_v1_: _dafny.Map
        d_2126_v1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([self.f36 for d_2127_i0_ in range(default__.abs(778))]): (d_2125_v0_ if False else d_2125_v0_)})
        d_2126_v1_ = (d_2126_v1_).set((self).fm3(self.f28, self.f31, globalState), d_2125_v0_)
        d_2128_v2_: _dafny.Array
        def lambda93_(d_2129_i1_):
            return ((0) - ((0) - ((self).f34))) >= (len(_dafny.Map({self.f28: self.f31})))

        init61_ = lambda93_
        nw380_ = _dafny.Array(None, 4)
        for i0_61_ in range(nw380_.length(0)):
            nw380_[i0_61_] = init61_(i0_61_)
        d_2128_v2_ = nw380_
        index313_ = default__.safeIndex(811, (d_2128_v2_).length(0))
        (d_2128_v2_)[index313_] = not((self.f31) and (True))
        d_2130_v3_: _dafny.MultiSet
        d_2130_v3_ = _dafny.MultiSet([self.f28])
        d_2131_v4_: _dafny.Map
        d_2131_v4_ = _dafny.Map({(self).f34: d_2130_v3_})
        index314_ = default__.safeIndex(195, (d_2125_v0_).length(0))
        (d_2125_v0_)[index314_] = len((self).fm3(self.f35, ((self).f32)[default__.safeIndex((self).f30, len((self).f32))], globalState))
        d_2132_v5_: _dafny.MultiSet
        d_2132_v5_ = _dafny.MultiSet([(self).f38, 325])
        index315_ = default__.safeIndex(654, (d_2125_v0_).length(0))
        (d_2125_v0_)[index315_] = default__.fm39(self.f31, d_2132_v5_, self.f28, globalState)
        index316_ = default__.safeIndex(811, (d_2128_v2_).length(0))
        index317_ = default__.safeIndex(195, (d_2125_v0_).length(0))
        index318_ = default__.safeIndex(654, (d_2125_v0_).length(0))
        rhs338_ = self.f36
        rhs339_ = self.f29
        rhs340_ = ((_dafny.Map({-867: d_2130_v3_})) | (d_2131_v4_)) | ((d_2131_v4_) | (d_2131_v4_))
        rhs341_ = (self).f30
        rhs342_ = (self).f38
        lhs279_ = self
        lhs280_ = d_2128_v2_
        lhs281_ = default__.safeIndex(811, (d_2128_v2_).length(0))
        lhs282_ = d_2125_v0_
        lhs283_ = default__.safeIndex(195, (d_2125_v0_).length(0))
        lhs284_ = d_2125_v0_
        lhs285_ = default__.safeIndex(654, (d_2125_v0_).length(0))
        lhs279_.f36 = rhs338_
        lhs280_[lhs281_] = rhs339_
        d_2131_v4_ = rhs340_
        lhs282_[lhs283_] = rhs341_
        lhs284_[lhs285_] = rhs342_
        if ((self).f32) < ((self).f32):
            (globalState).f13 = (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]
            d_2125_v0_ = d_2125_v0_
            d_2133_v6_: _dafny.Seq
            d_2133_v6_ = _dafny.SeqWithoutIsStrInference([self.f36, self.f36])
            d_2134_v7_: _dafny.Seq
            d_2134_v7_ = _dafny.SeqWithoutIsStrInference([True, False, (self.f29 if not(self.f35) else self.f29), self.f28])
            d_2135_v8_: _dafny.Set
            d_2135_v8_ = _dafny.Set({(self).f34})
            rhs343_ = (_dafny.SeqWithoutIsStrInference([self.f36])) + (d_2133_v6_)
            rhs344_ = ((self).f34) < ((0) - ((self).f38))
            rhs345_ = ((self).f32) + (d_2134_v7_)
            rhs346_ = (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]
            rhs347_ = d_2135_v8_
            lhs286_ = globalState
            lhs287_ = globalState
            d_2133_v6_ = rhs343_
            lhs286_.f26 = rhs344_
            d_2134_v7_ = rhs345_
            lhs287_.f13 = rhs346_
            d_2135_v8_ = rhs347_
            d_2136_v9_: _dafny.Map
            d_2136_v9_ = _dafny.Map({d_2125_v0_: (self).fm3(self.f31, (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))], globalState)})
            d_2136_v9_ = (d_2136_v9_).set(d_2125_v0_, d_2133_v6_)
            d_2137_v10_: _dafny.Map
            d_2137_v10_ = _dafny.Map({995: not((d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))])})
            d_2137_v10_ = (d_2137_v10_).set((d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))], self.f35)
        elif True:
            d_2138_v11_: _dafny.Seq
            d_2138_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aki"))
            d_2139_v12_: _dafny.Seq
            d_2139_v12_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ox")))])
            d_2140_v13_: _dafny.Seq
            d_2140_v13_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - ((d_2139_v12_)[default__.safeIndex((0) - ((self).f30), len(d_2139_v12_))]))])
            rhs348_ = (_dafny.SeqWithoutIsStrInference([self.f36 for d_2141_i2_ in range(default__.abs(272))])) + (d_2138_v11_)
            rhs349_ = not ((self.f28) or ((d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))])) or ((default__.fm60(globalState)) != ((_dafny.Map({self.f28: d_2139_v12_})).set((d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))], d_2139_v12_)))
            rhs350_ = (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]
            lhs288_ = self
            d_2138_v11_ = rhs348_
            lhs288_.f35 = rhs349_
            r0 = rhs350_
            index319_ = default__.safeIndex(195, (d_2125_v0_).length(0))
            (d_2125_v0_)[index319_] = (0) - (((self).f38) + (len((self).f32)))
            d_2142_v14_: _dafny.Map
            out71_: _dafny.Map
            out71_ = (self).m4(globalState)
            d_2142_v14_ = out71_
            (globalState).f13 = (self).f30
            d_2143_v15_: D0
            d_2143_v15_ = D0_DC1(d_2130_v3_)
            pat_let_tv23_ = d_2130_v3_
            d_2144_v16_: _dafny.Map
            def iife124_(_pat_let33_0):
                def iife125_(d_2145_dt__update__tmp_h0_):
                    def iife126_(_pat_let34_0):
                        def iife127_(d_2146_dt__update_hcf1_h0_):
                            return D0_DC1(d_2146_dt__update_hcf1_h0_)
                        return iife127_(_pat_let34_0)
                    return iife126_(pat_let_tv23_)
                return iife125_(_pat_let33_0)
            d_2144_v16_ = _dafny.Map({iife124_(d_2143_v15_): True})
            d_2147_v17_: D8
            d_2147_v17_ = D8_DC21(d_2144_v16_)
            d_2148_v18_: _dafny.Map
            d_2148_v18_ = _dafny.Map({(self).f38: d_2144_v16_})
            d_2149_v19_: _dafny.Set
            d_2149_v19_ = _dafny.Set({(d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))], self.f29})
            d_2150_v21_: _dafny.MultiSet
            d_2150_v21_ = _dafny.MultiSet([d_2143_v15_])
            pat_let_tv24_ = d_2150_v21_
            pat_let_tv25_ = d_2150_v21_
            d_2151_v22_: _dafny.Array
            nw381_ = _dafny.Array(None, 18)
            nw381_[int(0)] = d_2147_v17_
            def iife128_(_pat_let35_0):
                def iife129_(d_2152_dt__update__tmp_h1_):
                    def iife131_():
                        coll58_ = _dafny.Map()
                        compr_58_: D0
                        for compr_58_ in (pat_let_tv24_).Elements:
                            d_2153_v20_: D0 = compr_58_
                            if (d_2153_v20_) in (pat_let_tv25_):
                                coll58_[d_2153_v20_] = self.f29
                        return _dafny.Map(coll58_)
                    def iife130_(_pat_let36_0):
                        def iife132_(d_2154_dt__update_hcf40_h0_):
                            return D8_DC21(d_2154_dt__update_hcf40_h0_)
                        return iife132_(_pat_let36_0)
                    return iife130_(iife131_()
                    )
                return iife129_(_pat_let35_0)
            nw381_[int(1)] = iife128_(D8_DC21(((d_2148_v18_)[len(d_2149_v19_)] if (len(d_2149_v19_)) in (d_2148_v18_) else _dafny.Map({d_2143_v15_: (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))]}))))
            nw381_[int(2)] = d_2147_v17_
            nw381_[int(3)] = d_2147_v17_
            nw381_[int(4)] = d_2147_v17_
            nw381_[int(5)] = default__.fm61(d_2138_v11_, (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))], (self).f30, globalState)
            nw381_[int(6)] = D8_DC21(d_2144_v16_)
            nw381_[int(7)] = d_2147_v17_
            nw381_[int(8)] = D8_DC21(d_2144_v16_)
            nw381_[int(9)] = d_2147_v17_
            nw381_[int(10)] = d_2147_v17_
            nw381_[int(11)] = D8_DC21(_dafny.Map({d_2143_v15_: self.f35}))
            nw381_[int(12)] = d_2147_v17_
            nw381_[int(13)] = d_2147_v17_
            nw381_[int(14)] = D8_DC21(_dafny.Map({D0_DC1(_dafny.MultiSet([self.f31, True])): self.f29}))
            nw381_[int(15)] = d_2147_v17_
            nw381_[int(16)] = d_2147_v17_
            nw381_[int(17)] = d_2147_v17_
            d_2151_v22_ = nw381_
            index320_ = default__.safeIndex(661, (d_2151_v22_).length(0))
            (d_2151_v22_)[index320_] = d_2147_v17_
            index321_ = default__.safeIndex(661, (d_2151_v22_).length(0))
            rhs351_ = d_2147_v17_
            rhs352_ = ((d_2132_v5_) | (_dafny.MultiSet(d_2140_v13_))) - ((d_2132_v5_) - (d_2132_v5_))
            lhs289_ = d_2151_v22_
            lhs290_ = default__.safeIndex(661, (d_2151_v22_).length(0))
            lhs289_[lhs290_] = rhs351_
            d_2132_v5_ = rhs352_
        (globalState).f13 = (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]
        d_2155_i3_: int
        d_2155_i3_ = 0
        with _dafny.label("7"):
            while ((len(_dafny.Map({d_2125_v0_: (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))]}))) * ((self).f34)) > (((d_2130_v3_).set(((self).f32)[default__.safeIndex(((d_2130_v3_)[(d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))]] if ((d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))]) in (d_2130_v3_) else (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]), len((self).f32))], default__.abs((self).f34))).cardinality):
                with _dafny.c_label("7"):
                    if (d_2155_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_2155_i3_ = (d_2155_i3_) + (1)
                    d_2156_v23_: _dafny.Seq
                    d_2156_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wejhofhx"))
                    d_2157_v24_: _dafny.Map
                    d_2157_v24_ = _dafny.Map({d_2156_v23_: (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))]})
                    d_2158_v25_: _dafny.Array
                    nw382_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                    d_2158_v25_ = nw382_
                    d_2159_v26_: C0
                    nw383_ = C0()
                    nw383_.ctor__(d_2158_v25_)
                    d_2159_v26_ = nw383_
                    d_2160_v27_: D9
                    d_2160_v27_ = D9_DC27(d_2159_v26_, d_2156_v23_)
                    d_2161_v28_: _dafny.Map
                    d_2161_v28_ = _dafny.Map({self.f28: d_2160_v27_})
                    d_2162_v29_: _dafny.Map
                    d_2162_v29_ = _dafny.Map({(d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]: d_2161_v28_})
                    rhs353_ = d_2157_v24_
                    rhs354_ = default__.safeModuloInt(len((d_2162_v29_) | (d_2162_v29_)), (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))])
                    rhs355_ = d_2125_v0_
                    lhs291_ = globalState
                    d_2157_v24_ = rhs353_
                    lhs291_.f13 = rhs354_
                    d_2125_v0_ = rhs355_
                    d_2163_v30_: _dafny.Seq
                    d_2163_v30_ = _dafny.SeqWithoutIsStrInference([(d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))], (self).f38])
                    d_2164_v31_: _dafny.Array
                    nw384_ = _dafny.Array(None, 16)
                    nw384_[int(0)] = d_2163_v30_
                    nw384_[int(1)] = d_2163_v30_
                    nw384_[int(2)] = (d_2163_v30_).set(default__.safeIndex((d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))], len(d_2163_v30_)), (self).f38)
                    nw384_[int(3)] = d_2163_v30_
                    nw384_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_2132_v5_).cardinality for d_2165_i4_ in range(default__.abs(920))])
                    nw384_[int(5)] = d_2163_v30_
                    nw384_[int(6)] = (d_2163_v30_) + (d_2163_v30_)
                    nw384_[int(7)] = d_2163_v30_
                    nw384_[int(8)] = (d_2163_v30_) + (_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30]))
                    nw384_[int(9)] = default__.fm23(self.f31, (d_2130_v3_).cardinality, globalState)
                    nw384_[int(10)] = d_2163_v30_
                    nw384_[int(11)] = ((d_2163_v30_) + (d_2163_v30_)).set(default__.safeIndex((self).f38, len((d_2163_v30_) + (d_2163_v30_))), (self).f34)
                    nw384_[int(12)] = (default__.fm23(not(self.f35), (self).f38, globalState) if self.f29 else _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idwc"))) for d_2166_i5_ in range(default__.abs(107))]))
                    nw384_[int(13)] = (d_2163_v30_) + (d_2163_v30_)
                    nw384_[int(14)] = d_2163_v30_
                    nw384_[int(15)] = d_2163_v30_
                    d_2164_v31_ = nw384_
                    index322_ = default__.safeIndex(195, (d_2125_v0_).length(0))
                    rhs356_ = d_2164_v31_
                    rhs357_ = d_2125_v0_
                    rhs358_ = self.f31
                    rhs359_ = (len(d_2156_v23_) if (self).fm6(globalState) else (0) - (((self).f38) * ((self).f34)))
                    lhs292_ = globalState
                    lhs293_ = d_2125_v0_
                    lhs294_ = default__.safeIndex(195, (d_2125_v0_).length(0))
                    d_2164_v31_ = rhs356_
                    d_2125_v0_ = rhs357_
                    lhs292_.f1 = rhs358_
                    lhs293_[lhs294_] = rhs359_
                    d_2167_v32_: _dafny.Map
                    d_2167_v32_ = _dafny.Map({(0) - (len(d_2156_v23_)): ((self).f34) + ((self).f38)})
                    d_2168_v33_: _dafny.Map
                    d_2168_v33_ = _dafny.Map({d_2167_v32_: (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]})
                    d_2167_v32_ = (d_2167_v32_).set(len((d_2168_v33_) | (d_2168_v33_)), (0) - (default__.safeModuloInt(-265, (self).f38)))
                    d_2169_v34_: _dafny.Map
                    d_2169_v34_ = _dafny.Map({not (self.f31) or (self.f31): default__.fm62(self.f36, (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))], self.f29, self.f28, globalState)})
                    d_2170_v35_: _dafny.Map
                    d_2170_v35_ = _dafny.Map({self.f28: (self).f34})
                    d_2171_v36_: _dafny.Set
                    d_2171_v36_ = _dafny.Set({(self).f38, (self).f34})
                    d_2172_v37_: D21
                    d_2172_v37_ = D21_DC51(_dafny.Map({len(d_2170_v35_): d_2171_v36_}))
                    d_2169_v34_ = (d_2169_v34_).set(self.f35, d_2172_v37_)
                    pass
            pass
        hi11_ = (self).f38
        for d_2173_i6_ in range((self).f34, hi11_):
            d_2174_v38_: _dafny.Seq
            d_2174_v38_ = _dafny.SeqWithoutIsStrInference([d_2128_v2_])
            d_2175_v39_: _dafny.Seq
            d_2175_v39_ = d_2174_v38_
            d_2175_v39_ = d_2175_v39_
            index323_ = default__.safeIndex(195, (d_2125_v0_).length(0))
            (d_2125_v0_)[index323_] = (0) - ((88) + (d_2173_i6_))
            d_2176_v40_: _dafny.Map
            d_2176_v40_ = _dafny.Map({self.f36: d_2128_v2_})
            d_2176_v40_ = (d_2176_v40_).set(self.f36, d_2128_v2_)
            if self.f35:
                index324_ = default__.safeIndex(195, (d_2125_v0_).length(0))
                (d_2125_v0_)[index324_] = 55
                (globalState).f21 = d_2128_v2_
                d_2177_v41_: _dafny.Seq
                d_2177_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvq"))
                d_2178_v42_: C4
                nw385_ = C4()
                nw385_.ctor__(718, True, (_dafny.SeqWithoutIsStrInference([not(self.f31)])) + ((self).f32), (self).f33, len(d_2177_v41_), (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))], self.f35, self.f29)
                d_2178_v42_ = nw385_
                d_2179_v43_: _dafny.Map
                d_2179_v43_ = _dafny.Map({False: self.f35})
                rhs360_ = (d_2128_v2_)[default__.safeIndex(811, (d_2128_v2_).length(0))]
                rhs361_ = d_2177_v41_
                rhs362_ = d_2125_v0_
                rhs363_ = not((d_2130_v3_).issubset(_dafny.MultiSet([((d_2179_v43_)[self.f28] if (self.f28) in (d_2179_v43_) else self.f29), self.f31, False, self.f29])))
                rhs364_ = d_2178_v42_
                lhs295_ = self
                lhs296_ = self
                lhs295_.f31 = rhs360_
                d_2177_v41_ = rhs361_
                d_2125_v0_ = rhs362_
                lhs296_.f29 = rhs363_
                d_2178_v42_ = rhs364_
                d_2180_v44_: _dafny.Seq
                d_2180_v44_ = _dafny.SeqWithoutIsStrInference([d_2130_v3_])
                index325_ = default__.safeIndex(811, (d_2128_v2_).length(0))
                (d_2128_v2_)[index325_] = (self.f35 if (((d_2180_v44_)[default__.safeIndex((self).f38, len(d_2180_v44_))]).set(self.f31, default__.abs((d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]))).isdisjoint(d_2130_v3_) else True)
                d_2181_v45_: _dafny.Array
                def lambda94_(d_2182_v41_):
                    def lambda95_(d_2183_i7_):
                        return d_2182_v41_

                    return lambda95_

                init62_ = lambda94_(d_2177_v41_)
                nw386_ = _dafny.Array(None, 8)
                for i0_62_ in range(nw386_.length(0)):
                    nw386_[i0_62_] = init62_(i0_62_)
                d_2181_v45_ = nw386_
                index326_ = default__.safeIndex(145, (d_2181_v45_).length(0))
                (d_2181_v45_)[index326_] = d_2177_v41_
                index327_ = default__.safeIndex(145, (d_2181_v45_).length(0))
                (d_2181_v45_)[index327_] = ((d_2177_v41_) + (d_2177_v41_)) + ((self).fm3(self.f28, self.f31, globalState))
            elif True:
                d_2184_v46_: _dafny.Map
                d_2184_v46_ = _dafny.Map({True: self.f31})
                (globalState).f13 = (self).fm17(len((_dafny.Map({self.f28: self.f28})) | ((d_2184_v46_).set(self.f28, True))), d_2173_i6_, globalState)
                d_2185_v47_: _dafny.Seq
                d_2185_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmac"))
                d_2185_v47_ = d_2185_v47_
                d_2186_v48_: _dafny.Array
                nw387_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_2186_v48_ = nw387_
                d_2187_v49_: C0
                nw388_ = C0()
                nw388_.ctor__(d_2186_v48_)
                d_2187_v49_ = nw388_
                d_2187_v49_ = d_2187_v49_
                (globalState).f13 = d_2173_i6_
                index328_ = default__.safeIndex(195, (d_2125_v0_).length(0))
                (d_2125_v0_)[index328_] = (d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]
        d_2188_v50_: _dafny.MultiSet
        d_2188_v50_ = _dafny.MultiSet([d_2130_v3_, d_2130_v3_, default__.fm33(self.f29, globalState)])
        r0 = ((d_2188_v50_) | (((d_2188_v50_).set(d_2130_v3_, default__.abs((d_2125_v0_)[default__.safeIndex(195, (d_2125_v0_).length(0))]))).set(d_2130_v3_, default__.abs(444)))).cardinality
        nw389_ = _dafny.Array(_dafny.Set({}), 19)
        r1 = nw389_
        r2 = d_2128_v2_
        return r0, r1, r2

    @property
    def f38(self):
        return self._f38
    @property
    def f39(self):
        return self._f39

class C10(T4, T3, T2, T1, T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f31: bool = False
        self._f35: bool = False
        self._f36: str = _dafny.CodePoint('D')
        self._f34: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    @f29.setter
    def f29(self, value):
        self._f29 = value
    @property
    def f31(self):
        return self._f31
    @f31.setter
    def f31(self, value):
        self._f31 = value
    @property
    def f35(self):
        return self._f35
    @f35.setter
    def f35(self, value):
        self._f35 = value
    @property
    def f36(self):
        return self._f36
    @f36.setter
    def f36(self, value):
        self._f36 = value
    @property
    def f34(self):
        return self._f34
    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33
    @property
    def f30(self):
        return self._f30
    def ctor__(self, f36, f34, f35, f32, f33, f30, f31, f28, f29):
        (self).f36 = f36
        (self)._f34 = f34
        (self).f35 = f35
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f30 = f30
        (self).f31 = f31
        (self).f28 = f28
        (self).f29 = f29

    def fm7(self, p0, p1, globalState):
        return not(not((not(not(self.f29))) or (self.f35)))

    def fm6(self, globalState):
        return ((D0_DC0(self.f28) if self.f35 else D0_DC0(self.f31))).cf0

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife133_():
            coll59_ = _dafny.Set()
            compr_59_: int
            for compr_59_ in _dafny.IntegerRange(507, 18):
                d_2189_v0_: int = compr_59_
                if ((507) <= (d_2189_v0_)) and ((d_2189_v0_) < (18)):
                    coll59_ = coll59_.union(_dafny.Set([(d_2189_v0_) - ((self).f34)]))
            return _dafny.Set(coll59_)
        return (_dafny.Set({(self).f34, (self).f34})) - (iife133_()
        )

    def fm5(self, p0, p1, p2, p3, globalState):
        source34_ = D0_DC1(_dafny.MultiSet((self).f32))
        if source34_.is_DC1:
            d_2190___mcc_h0_ = source34_.cf1
            d_2191_cf1_ = d_2190___mcc_h0_
            return ((self).f32) + (_dafny.SeqWithoutIsStrInference([self.f28]))
        elif True:
            d_2192___mcc_h1_ = source34_.cf0
            d_2193_cf0_ = d_2192___mcc_h1_
            return (((self).f33)[default__.safeIndex((self).f34, len((self).f33))]) + ((self).f32)

    def fm3(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgpghnx"))) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_2194_i0_ in range(default__.abs(991))]))

    def fm2(self, globalState):
        if self.f31:
            return (_dafny.Map({self.f28: (0) - ((self).f30)})) | (_dafny.Map({self.f31: (self).f30}))
        elif True:
            return (_dafny.Map({self.f29: (0) - ((self).f34)})) | (_dafny.Map({self.f35: 936}))

    def fm8(self, p0, p1, p2, globalState):
        return (361) * ((369) * ((self).f30))

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (self).f31 = self.f28
        d_2195_v0_: _dafny.MultiSet
        d_2195_v0_ = _dafny.MultiSet([self.f29, self.f31, self.f29])
        d_2196_v1_: D0
        d_2196_v1_ = D0_DC1(d_2195_v0_)
        source35_ = d_2196_v1_
        if source35_.is_DC1:
            d_2197___mcc_h0_ = source35_.cf1
            d_2198_cf1_ = d_2197___mcc_h0_
            d_2199_v2_: _dafny.Array
            nw390_ = _dafny.Array(_dafny.CodePoint('D'), 7)
            d_2199_v2_ = nw390_
            d_2200_v3_: C0
            nw391_ = C0()
            nw391_.ctor__(d_2199_v2_)
            d_2200_v3_ = nw391_
            d_2201_v4_: _dafny.Seq
            d_2201_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdvneponb"))
            d_2202_v5_: _dafny.Seq
            d_2202_v5_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f34)])
            d_2203_v6_: _dafny.MultiSet
            d_2203_v6_ = _dafny.MultiSet([(self).f30])
            d_2204_v7_: _dafny.Array
            def lambda96_(d_2205_i0_):
                return self.f28

            init63_ = lambda96_
            nw392_ = _dafny.Array(None, 19)
            for i0_63_ in range(nw392_.length(0)):
                nw392_[i0_63_] = init63_(i0_63_)
            d_2204_v7_ = nw392_
            d_2206_v8_: _dafny.Map
            d_2206_v8_ = _dafny.Map({d_2204_v7_: d_2203_v6_})
            d_2207_v9_: _dafny.Array
            nw393_ = _dafny.Array(None, 12)
            nw393_[int(0)] = 15
            nw393_[int(1)] = len((d_2201_v4_) + (d_2201_v4_))
            nw393_[int(2)] = len(d_2202_v5_)
            nw393_[int(3)] = (d_2203_v6_).cardinality
            nw393_[int(4)] = (self).f34
            nw393_[int(5)] = (self).f34
            nw393_[int(6)] = (self).f30
            nw393_[int(7)] = (self).f30
            nw393_[int(8)] = (self).f30
            nw393_[int(9)] = len(d_2206_v8_)
            nw393_[int(10)] = (self).f30
            nw393_[int(11)] = default__.safeDivisionInt((self).f34, 590)
            d_2207_v9_ = nw393_
            index329_ = default__.safeIndex(469, (d_2207_v9_).length(0))
            (d_2207_v9_)[index329_] = (self).f30
            index330_ = default__.safeIndex(469, (d_2207_v9_).length(0))
            (d_2207_v9_)[index330_] = (0) - (((self).fm8(False, self.f36, self.f29, globalState)) + ((d_2202_v5_)[default__.safeIndex((self).f34, len(d_2202_v5_))]))
            index331_ = default__.safeIndex(469, (d_2207_v9_).length(0))
            (d_2207_v9_)[index331_] = (self).f30
            d_2208_v10_: _dafny.Map
            d_2208_v10_ = _dafny.Map({self.f28: (D1_DC2(d_2204_v7_)).cf2})
            d_2209_v11_: _dafny.MultiSet
            d_2209_v11_ = _dafny.MultiSet([((d_2208_v10_)[self.f28] if (self.f28) in (d_2208_v10_) else d_2204_v7_)])
            d_2210_v12_: _dafny.Set
            d_2210_v12_ = _dafny.Set({(d_2207_v9_)[default__.safeIndex(469, (d_2207_v9_).length(0))]})
            d_2211_v14_: int
            out72_: int
            def iife134_():
                coll60_ = _dafny.Set()
                compr_60_: int
                for compr_60_ in _dafny.IntegerRange(580, 286):
                    d_2212_v13_: int = compr_60_
                    if ((580) <= (d_2212_v13_)) and ((d_2212_v13_) < (286)):
                        coll60_ = coll60_.union(_dafny.Set([(d_2212_v13_) * (-415)]))
                return _dafny.Set(coll60_)
            out72_ = default__.m0((_dafny.SeqWithoutIsStrInference([((d_2209_v11_)[d_2204_v7_] if (d_2204_v7_) in (d_2209_v11_) else (d_2207_v9_)[default__.safeIndex(469, (d_2207_v9_).length(0))])])) + (d_2202_v5_), self.f29, ((d_2203_v6_)[len(_dafny.SeqWithoutIsStrInference([self.f28]))] if (len(_dafny.SeqWithoutIsStrInference([self.f28]))) in (d_2203_v6_) else len(d_2202_v5_)), (iife134_()
            ).ispropersubset(d_2210_v12_), globalState)
            d_2211_v14_ = out72_
        elif True:
            d_2213___mcc_h1_ = source35_.cf0
            d_2214_cf0_ = d_2213___mcc_h1_
            (globalState).f26 = self.f35
            d_2215_v15_: _dafny.Array
            def lambda97_(d_2216_i1_):
                return (self).f32

            init64_ = lambda97_
            nw394_ = _dafny.Array(None, 28)
            for i0_64_ in range(nw394_.length(0)):
                nw394_[i0_64_] = init64_(i0_64_)
            d_2215_v15_ = nw394_
            d_2217_v16_: _dafny.Seq
            d_2217_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erxeyj"))
            d_2218_v17_: _dafny.Seq
            d_2218_v17_ = _dafny.SeqWithoutIsStrInference([(self).f30, len(d_2217_v16_)])
            d_2219_v18_: _dafny.Set
            d_2219_v18_ = _dafny.Set({len(d_2218_v17_)})
            rhs365_ = d_2215_v15_
            rhs366_ = len((((self).fm4(self.f28, d_2219_v18_, (self).f30, d_2214_cf0_, globalState)) | (d_2219_v18_)) | ((d_2219_v18_) - (d_2219_v18_)))
            rhs367_ = d_2196_v1_
            rhs368_ = ((self).f30) + ((self).f34)
            lhs297_ = globalState
            d_2215_v15_ = rhs365_
            r1 = rhs366_
            d_2196_v1_ = rhs367_
            lhs297_.f13 = rhs368_
            d_2220_v19_: _dafny.Array
            nw395_ = _dafny.Array(None, 24)
            d_2220_v19_ = nw395_
            d_2221_v20_: _dafny.Map
            d_2221_v20_ = _dafny.Map({-720: 212})
            d_2222_v21_: _dafny.Map
            d_2222_v21_ = _dafny.Map({d_2220_v19_: len((default__.fm9(self.f36, len(d_2221_v20_), (self).f34, self.f35, globalState)) + (d_2218_v17_))})
            d_2222_v21_ = (d_2222_v21_).set((D2_DC5(d_2220_v19_)).cf8, 206)
            d_2214_cf0_ = self.f31
        d_2223_v22_: _dafny.Array
        def lambda98_(d_2224_i2_):
            return (_dafny.SeqWithoutIsStrInference([self.f36 for d_2225_i3_ in range(default__.abs(866))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxsini")))

        init65_ = lambda98_
        nw396_ = _dafny.Array(None, 6)
        for i0_65_ in range(nw396_.length(0)):
            nw396_[i0_65_] = init65_(i0_65_)
        d_2223_v22_ = nw396_
        index332_ = default__.safeIndex(190, (d_2223_v22_).length(0))
        (d_2223_v22_)[index332_] = False
        index333_ = default__.safeIndex(190, (d_2223_v22_).length(0))
        def lambda99_(source36_):
            if source36_.is_DC1:
                d_2226___mcc_h2_ = source36_.cf1
                d_2227_cf1_ = d_2226___mcc_h2_
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2228_i4_ in range(default__.abs(627))])
            elif True:
                d_2229___mcc_h3_ = source36_.cf0
                d_2230_cf0_ = d_2229___mcc_h3_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srxfo"))) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_2231_i5_ in range(default__.abs(465))]))

        rhs369_ = lambda99_(D0_DC0(self.f29))
        rhs370_ = self.f31
        lhs298_ = d_2223_v22_
        lhs299_ = default__.safeIndex(190, (d_2223_v22_).length(0))
        r2 = rhs369_
        lhs298_[lhs299_] = rhs370_
        d_2232_v23_: _dafny.Seq
        d_2232_v23_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f34, 106), (self).f34])
        d_2233_v24_: _dafny.Seq
        d_2233_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
        rhs371_ = d_2232_v23_
        rhs372_ = not(not(self.f28))
        rhs373_ = d_2233_v24_
        d_2232_v23_ = rhs371_
        r0 = rhs372_
        r2 = rhs373_
        d_2234_i6_: int
        d_2234_i6_ = 0
        with _dafny.label("8"):
            while not((_dafny.SeqWithoutIsStrInference([self.f36 for d_2270_i7_ in range(default__.abs(241))])) < ((self).fm3((d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))], (d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))], globalState))):
                with _dafny.c_label("8"):
                    if (d_2234_i6_) >= (100):
                        raise _dafny.Break("8")
                    d_2234_i6_ = (d_2234_i6_) + (1)
                    d_2235_v25_: _dafny.Set
                    d_2235_v25_ = _dafny.Set({False, self.f35, self.f29})
                    d_2235_v25_ = _dafny.Set({self.f35, ((self).f30) > ((self).f34), (d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))], (self.f28) not in (d_2235_v25_)})
                    d_2236_v26_: D0
                    d_2236_v26_ = D0_DC0(True)
                    d_2237_v27_: _dafny.Array
                    nw397_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                    d_2237_v27_ = nw397_
                    d_2238_v28_: _dafny.Map
                    d_2238_v28_ = _dafny.Map({False: (d_2237_v27_ if (d_2236_v26_).cf0 else d_2237_v27_)})
                    d_2238_v28_ = (d_2238_v28_).set(self.f28, (d_2237_v27_ if (d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))] else ((d_2238_v28_)[self.f35] if (self.f35) in (d_2238_v28_) else d_2237_v27_)))
                    (globalState).f21 = d_2223_v22_
                    d_2239_v29_: _dafny.Array
                    nw398_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                    d_2239_v29_ = nw398_
                    d_2240_v30_: C0
                    nw399_ = C0()
                    nw399_.ctor__(d_2239_v29_)
                    d_2240_v30_ = nw399_
                    d_2241_v31_: D1
                    d_2241_v31_ = D1_DC3(self.f35, ((self).f30) * ((self).f34), d_2240_v30_, 929)
                    source37_ = d_2241_v31_
                    if source37_.is_DC3:
                        d_2242___mcc_h4_ = source37_.cf3
                        d_2243___mcc_h5_ = source37_.cf4
                        d_2244___mcc_h6_ = source37_.cf5
                        d_2245___mcc_h7_ = source37_.cf6
                        d_2246_cf6_ = d_2245___mcc_h7_
                        d_2247_cf5_ = d_2244___mcc_h6_
                        d_2248_cf4_ = d_2243___mcc_h5_
                        d_2249_cf3_ = d_2242___mcc_h4_
                        d_2250_v32_: _dafny.Map
                        d_2250_v32_ = _dafny.Map({((self).f30) * ((0) - (d_2246_cf6_)): (self).f34})
                        d_2251_v33_: D2
                        d_2251_v33_ = D2_DC6((self).f34, (d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))], len(d_2233_v24_), self.f29, (self).f30)
                        d_2250_v32_ = (d_2250_v32_).set(d_2246_cf6_, (d_2251_v33_).cf9)
                        (globalState).f1 = (_dafny.Set({default__.fm11((self).f33, self.f35, True, (self).f34, globalState), d_2235_v25_, d_2235_v25_})).ispropersubset(default__.fm10(d_2235_v25_, _dafny.CodePoint('y'), d_2248_cf4_, not(self.f31), globalState))
                        r2 = ((d_2233_v24_) + (d_2233_v24_)).set(default__.safeIndex(d_2248_cf4_, len((d_2233_v24_) + (d_2233_v24_))), _dafny.CodePoint('r'))
                        (globalState).f12 = self.f31
                    elif source37_.is_DC2:
                        d_2252___mcc_h8_ = source37_.cf2
                        d_2253_cf2_ = d_2252___mcc_h8_
                        d_2254_v34_: D2
                        d_2254_v34_ = D2_DC6((self).f30, self.f35, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len((self).f32), (self).f30]): (self).f34})), not(default__.fm1(True, (self).f30, globalState)), (self).f34)
                        d_2255_v35_: _dafny.Map
                        d_2255_v35_ = _dafny.Map({self.f29: (self).f30})
                        d_2254_v34_ = default__.fm12(self.f36, (0) - (((self).f30 if True else (self).f30)), d_2255_v35_, (d_2233_v24_) < (d_2233_v24_), globalState)
                        d_2256_v36_: _dafny.Seq
                        d_2256_v36_ = _dafny.SeqWithoutIsStrInference([d_2255_v35_])
                        d_2257_v37_: D3
                        d_2257_v37_ = D3_DC7(_dafny.SeqWithoutIsStrInference([_dafny.Map({self.f28: (0) - ((self).f34)}) for d_2258_i8_ in range(default__.abs(-446))]))
                        d_2259_v38_: _dafny.Array
                        nw400_ = _dafny.Array(None, 6)
                        nw400_[int(0)] = d_2256_v36_
                        nw400_[int(1)] = (d_2257_v37_).cf14
                        nw400_[int(2)] = d_2256_v36_
                        nw400_[int(3)] = d_2256_v36_
                        nw400_[int(4)] = d_2256_v36_
                        nw400_[int(5)] = d_2256_v36_
                        d_2259_v38_ = nw400_
                        index334_ = default__.safeIndex(613, (d_2259_v38_).length(0))
                        (d_2259_v38_)[index334_] = d_2256_v36_
                        index335_ = default__.safeIndex(190, (d_2223_v22_).length(0))
                        index336_ = default__.safeIndex(613, (d_2259_v38_).length(0))
                        rhs374_ = ((self).fm8(self.f31, _dafny.CodePoint('p'), self.f31, globalState)) > ((self).f30)
                        rhs375_ = (d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))]
                        rhs376_ = _dafny.CodePoint('d')
                        rhs377_ = (self).f30
                        rhs378_ = ((d_2256_v36_) + (d_2256_v36_)) + (d_2256_v36_)
                        lhs300_ = d_2223_v22_
                        lhs301_ = default__.safeIndex(190, (d_2223_v22_).length(0))
                        lhs302_ = self
                        lhs303_ = globalState
                        lhs304_ = d_2259_v38_
                        lhs305_ = default__.safeIndex(613, (d_2259_v38_).length(0))
                        r0 = rhs374_
                        lhs300_[lhs301_] = rhs375_
                        lhs302_.f36 = rhs376_
                        lhs303_.f13 = rhs377_
                        lhs304_[lhs305_] = rhs378_
                        d_2255_v35_ = (d_2255_v35_).set((d_2223_v22_)[default__.safeIndex(190, (d_2223_v22_).length(0))], (self).f34)
                        d_2240_v30_ = d_2240_v30_
                    elif True:
                        d_2260___mcc_h9_ = source37_.cf7
                        d_2261_cf7_ = d_2260___mcc_h9_
                        (globalState).f26 = self.f31
                        d_2262_v39_: _dafny.MultiSet
                        d_2262_v39_ = _dafny.MultiSet([self.f36, self.f36, self.f36])
                        d_2263_v40_: D3
                        d_2263_v40_ = D3_DC8((d_2262_v39_) | (_dafny.MultiSet([self.f36, self.f36, self.f36])), True, (self).f30)
                        d_2264_v41_: _dafny.Map
                        d_2264_v41_ = _dafny.Map({_dafny.Map({len((self).f33): (self).f30}): (self).f34})
                        d_2265_v43_: _dafny.Map
                        d_2265_v43_ = _dafny.Map({(self).f30: (self).f30})
                        d_2266_v44_: _dafny.Set
                        d_2266_v44_ = _dafny.Set({d_2265_v43_})
                        d_2267_v46_: _dafny.MultiSet
                        d_2267_v46_ = _dafny.MultiSet([d_2235_v25_])
                        def iife135_():
                            coll61_ = _dafny.Map()
                            compr_61_: _dafny.Map
                            for compr_61_ in (d_2266_v44_).Elements:
                                d_2268_v42_: _dafny.Map = compr_61_
                                if (d_2268_v42_) in (d_2266_v44_):
                                    def iife136_():
                                        coll62_ = _dafny.Set()
                                        compr_62_: int
                                        for compr_62_ in _dafny.IntegerRange(590, 316):
                                            d_2269_v45_: int = compr_62_
                                            if ((590) <= (d_2269_v45_)) and ((d_2269_v45_) < (316)):
                                                coll62_ = coll62_.union(_dafny.Set([(d_2269_v45_) + (len(d_2232_v23_))]))
                                        return _dafny.Set(coll62_)
                                    coll61_[d_2268_v42_] = len(iife136_()
                                    )
                            return _dafny.Map(coll61_)
                        d_2263_v40_ = default__.fm13(len((d_2264_v41_) | (iife135_()
                        )), (self).f34, (self).f34, (d_2267_v46_).intersection(d_2267_v46_), globalState)
                        (globalState).f13 = (d_2232_v23_)[default__.safeIndex((self).f34, len(d_2232_v23_))]
                        (globalState).f1 = False
                    pass
            pass
        d_2271_i9_: int
        d_2271_i9_ = 0
        with _dafny.label("9"):
            while self.f31:
                with _dafny.c_label("9"):
                    if (d_2271_i9_) >= (100):
                        raise _dafny.Break("9")
                    d_2271_i9_ = (d_2271_i9_) + (1)
                    (self).f31 = not((self).fm7(_dafny.MultiSet([(self).f34]), (self).f30, globalState))
                    r1 = ((self).f30) + ((0) - (len(_dafny.Map({_dafny.CodePoint('u'): self.f29}))))
                    d_2272_v47_: _dafny.MultiSet
                    d_2272_v47_ = _dafny.MultiSet([default__.fm14((self).f30, False, 427, default__.fm15((self).f34, (self).f30, globalState), globalState), self.f36])
                    d_2273_v48_: _dafny.Set
                    d_2273_v48_ = _dafny.Set({d_2272_v47_})
                    d_2273_v48_ = _dafny.Set({d_2272_v47_})
                    d_2274_v49_: _dafny.Array
                    nw401_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                    d_2274_v49_ = nw401_
                    d_2275_v50_: C0
                    nw402_ = C0()
                    nw402_.ctor__(d_2274_v49_)
                    d_2275_v50_ = nw402_
                    pass
            pass
        r0 = self.f31
        r1 = (self).f34
        r2 = ((self).fm3(((self).f32)[default__.safeIndex((self).f34, len((self).f32))], not(True), globalState)) + ((d_2233_v24_) + (d_2233_v24_))
        return r0, r1, r2

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: D0 = D0.default()()
        r2: D0 = D0.default()()
        r3: bool = False
        d_2276_v0_: _dafny.Array
        def lambda100_(d_2277_i0_):
            return self.f31

        init66_ = lambda100_
        nw403_ = _dafny.Array(None, 16)
        for i0_66_ in range(nw403_.length(0)):
            nw403_[i0_66_] = init66_(i0_66_)
        d_2276_v0_ = nw403_
        d_2278_v1_: _dafny.MultiSet
        d_2278_v1_ = _dafny.MultiSet([self.f36, self.f36, self.f36, self.f36])
        d_2279_v2_: D3
        d_2279_v2_ = D3_DC8(d_2278_v1_, self.f28, (self).f34)
        index337_ = default__.safeIndex(695, (d_2276_v0_).length(0))
        (d_2276_v0_)[index337_] = ((self).f30) == ((d_2279_v2_).cf17)
        index338_ = default__.safeIndex(695, (d_2276_v0_).length(0))
        (d_2276_v0_)[index338_] = self.f29
        d_2280_i1_: int
        d_2280_i1_ = 0
        with _dafny.label("10"):
            while self.f35:
                with _dafny.c_label("10"):
                    if (d_2280_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_2280_i1_ = (d_2280_i1_) + (1)
                    d_2281_v3_: D3
                    d_2281_v3_ = D3_DC9(not(self.f28), self.f35, d_2276_v0_, self.f28)
                    d_2282_v4_: _dafny.Seq
                    d_2282_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kddokf"))
                    index339_ = default__.safeIndex(695, (d_2276_v0_).length(0))
                    rhs379_ = D3_DC9(self.f28, (not(self.f35)) not in (_dafny.SeqWithoutIsStrInference([False, (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], self.f28, (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], self.f35])), d_2276_v0_, (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))])
                    rhs380_ = default__.safeDivisionInt(len((d_2282_v4_) + (d_2282_v4_)), p0)
                    rhs381_ = not (self.f28) or (self.f28)
                    rhs382_ = d_2281_v3_
                    rhs383_ = self.f31
                    lhs306_ = globalState
                    lhs307_ = d_2276_v0_
                    lhs308_ = default__.safeIndex(695, (d_2276_v0_).length(0))
                    lhs309_ = self
                    d_2281_v3_ = rhs379_
                    lhs306_.f13 = rhs380_
                    lhs307_[lhs308_] = rhs381_
                    d_2281_v3_ = rhs382_
                    lhs309_.f31 = rhs383_
                    d_2283_v5_: _dafny.MultiSet
                    d_2283_v5_ = _dafny.MultiSet([p2])
                    d_2284_v6_: _dafny.Map
                    d_2284_v6_ = _dafny.Map({(d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]: len((self).f32)})
                    (globalState).f26 = not ((d_2283_v5_).ispropersubset(d_2283_v5_)) or ((d_2284_v6_) != (d_2284_v6_))
                    d_2285_v7_: _dafny.Array
                    def lambda101_(d_2286_p0_):
                        def lambda102_(d_2287_i2_):
                            return default__.safeModuloInt(d_2287_i2_, d_2286_p0_)

                        return lambda102_

                    init67_ = lambda101_(p0)
                    nw404_ = _dafny.Array(None, 10)
                    for i0_67_ in range(nw404_.length(0)):
                        nw404_[i0_67_] = init67_(i0_67_)
                    d_2285_v7_ = nw404_
                    d_2285_v7_ = d_2285_v7_
                    (globalState).f1 = not((self.f31 if self.f28 else (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]))
                    pass
            pass
        d_2288_v8_: _dafny.Array
        nw405_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
        d_2288_v8_ = nw405_
        d_2289_v9_: _dafny.Seq
        d_2289_v9_ = _dafny.SeqWithoutIsStrInference([d_2288_v8_])
        d_2290_v10_: _dafny.Set
        d_2290_v10_ = _dafny.Set({(d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]})
        d_2291_v11_: _dafny.MultiSet
        d_2291_v11_ = _dafny.MultiSet([d_2290_v10_])
        d_2292_v12_: _dafny.Array
        nw406_ = _dafny.Array(None, 28)
        nw406_[int(0)] = (d_2289_v9_)[default__.safeIndex(((d_2291_v11_)[_dafny.Set({(d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]})] if (_dafny.Set({(d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]})) in (d_2291_v11_) else (self).f34), len(d_2289_v9_))]
        nw406_[int(1)] = d_2288_v8_
        nw406_[int(2)] = d_2288_v8_
        nw406_[int(3)] = d_2288_v8_
        nw406_[int(4)] = d_2288_v8_
        nw406_[int(5)] = d_2288_v8_
        nw406_[int(6)] = d_2288_v8_
        nw406_[int(7)] = d_2288_v8_
        nw406_[int(8)] = d_2288_v8_
        nw406_[int(9)] = (d_2288_v8_ if self.f31 else d_2288_v8_)
        nw406_[int(10)] = d_2288_v8_
        nw406_[int(11)] = d_2288_v8_
        nw406_[int(12)] = d_2288_v8_
        nw406_[int(13)] = d_2288_v8_
        nw406_[int(14)] = d_2288_v8_
        nw406_[int(15)] = d_2288_v8_
        nw406_[int(16)] = d_2288_v8_
        nw406_[int(17)] = d_2288_v8_
        nw406_[int(18)] = d_2288_v8_
        nw406_[int(19)] = d_2288_v8_
        nw406_[int(20)] = (d_2288_v8_ if self.f28 else d_2288_v8_)
        nw406_[int(21)] = d_2288_v8_
        nw406_[int(22)] = d_2288_v8_
        nw406_[int(23)] = d_2288_v8_
        nw406_[int(24)] = d_2288_v8_
        nw406_[int(25)] = d_2288_v8_
        nw406_[int(26)] = d_2288_v8_
        nw406_[int(27)] = d_2288_v8_
        d_2292_v12_ = nw406_
        index340_ = default__.safeIndex(469, (d_2292_v12_).length(0))
        (d_2292_v12_)[index340_] = d_2288_v8_
        index341_ = default__.safeIndex(469, (d_2292_v12_).length(0))
        (d_2292_v12_)[index341_] = d_2288_v8_
        d_2293_i3_: int
        d_2293_i3_ = 0
        with _dafny.label("11"):
            while True:
                with _dafny.c_label("11"):
                    if (d_2293_i3_) >= (100):
                        raise _dafny.Break("11")
                    d_2293_i3_ = (d_2293_i3_) + (1)
                    index342_ = default__.safeIndex(695, (d_2276_v0_).length(0))
                    (d_2276_v0_)[index342_] = (((self).f34) - (p1)) <= ((self).f34)
                    d_2294_v13_: _dafny.MultiSet
                    d_2294_v13_ = _dafny.MultiSet([self.f35, self.f28])
                    d_2295_v14_: _dafny.Seq
                    d_2295_v14_ = _dafny.SeqWithoutIsStrInference([d_2294_v13_])
                    d_2296_v15_: _dafny.Array
                    nw407_ = _dafny.Array(None, 14)
                    nw407_[int(0)] = d_2294_v13_
                    nw407_[int(1)] = (d_2294_v13_ if self.f28 else d_2294_v13_)
                    nw407_[int(2)] = d_2294_v13_
                    nw407_[int(3)] = _dafny.MultiSet([self.f31, ((self).f32)[default__.safeIndex((self).f34, len((self).f32))], self.f29, self.f28, self.f28])
                    nw407_[int(4)] = d_2294_v13_
                    nw407_[int(5)] = d_2294_v13_
                    nw407_[int(6)] = d_2294_v13_
                    nw407_[int(7)] = (d_2294_v13_ if self.f31 else d_2294_v13_)
                    nw407_[int(8)] = (d_2295_v14_)[default__.safeIndex(p2, len(d_2295_v14_))]
                    nw407_[int(9)] = _dafny.MultiSet([((self).f32)[default__.safeIndex((self).f30, len((self).f32))]])
                    nw407_[int(10)] = (d_2294_v13_) | (d_2294_v13_)
                    nw407_[int(11)] = _dafny.MultiSet([self.f35, self.f35, (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], self.f28, False])
                    nw407_[int(12)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(self.f31)]))
                    nw407_[int(13)] = d_2294_v13_
                    d_2296_v15_ = nw407_
                    index343_ = default__.safeIndex(490, (d_2296_v15_).length(0))
                    (d_2296_v15_)[index343_] = (d_2294_v13_).set(False, default__.abs((0) - ((self).f30)))
                    d_2297_v16_: _dafny.Seq
                    d_2297_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unmvjnend"))
                    d_2298_v17_: _dafny.Map
                    d_2298_v17_ = _dafny.Map({self.f29: d_2297_v16_})
                    index344_ = default__.safeIndex(490, (d_2296_v15_).length(0))
                    (d_2296_v15_)[index344_] = _dafny.MultiSet([self.f29, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfc"))) == (((d_2298_v17_)[(d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]] if ((d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))]) in (d_2298_v17_) else d_2297_v16_))])
                    d_2299_v18_: _dafny.Array
                    nw408_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                    d_2299_v18_ = nw408_
                    d_2300_v19_: _dafny.Map
                    d_2300_v19_ = _dafny.Map({self.f35: d_2299_v18_})
                    d_2301_v20_: C0
                    nw409_ = C0()
                    nw409_.ctor__(((d_2300_v19_)[self.f29] if (self.f29) in (d_2300_v19_) else d_2299_v18_))
                    d_2301_v20_ = nw409_
                    d_2302_v21_: _dafny.Array
                    def lambda103_(d_2303_i4_):
                        return (self).f32

                    init68_ = lambda103_
                    nw410_ = _dafny.Array(None, 13)
                    for i0_68_ in range(nw410_.length(0)):
                        nw410_[i0_68_] = init68_(i0_68_)
                    d_2302_v21_ = nw410_
                    d_2304_v22_: D1
                    d_2304_v22_ = D1_DC3((d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], p0, d_2301_v20_, (self).f30)
                    d_2305_v23_: T4
                    nw411_ = C3()
                    nw411_.ctor__(len((self).f32), d_2302_v21_, self.f36, (0) - ((self).f30), (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], _dafny.SeqWithoutIsStrInference([self.f31, (d_2304_v22_).cf3]), _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f31])]), len((self).f32), self.f35, self.f28, not(self.f29))
                    d_2305_v23_ = nw411_
                    d_2306_v24_: _dafny.Map
                    d_2306_v24_ = _dafny.Map({d_2305_v23_: p1})
                    d_2307_v25_: _dafny.Map
                    d_2307_v25_ = _dafny.Map({(self.f36 if self.f28 else self.f36): ((d_2306_v24_)[d_2305_v23_] if (d_2305_v23_) in (d_2306_v24_) else (0) - (p2))})
                    d_2308_v26_: _dafny.Seq
                    d_2308_v26_ = _dafny.SeqWithoutIsStrInference([(self).f30])
                    d_2309_v27_: _dafny.MultiSet
                    d_2309_v27_ = _dafny.MultiSet([(d_2308_v26_)[default__.safeIndex((self).f30, len(d_2308_v26_))], (d_2305_v23_).f34, (d_2305_v23_).f30])
                    d_2307_v25_ = (d_2307_v25_).set(default__.fm34((self).f34, p0, ((d_2309_v27_)[default__.fm46((d_2305_v23_).f34, p1, self.f31, globalState)] if (default__.fm46((d_2305_v23_).f34, p1, self.f31, globalState)) in (d_2309_v27_) else p2), False, globalState), (d_2305_v23_).f34)
                    pass
            pass
        d_2310_v28_: _dafny.Array
        nw412_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_2310_v28_ = nw412_
        d_2311_v29_: _dafny.Array
        nw413_ = _dafny.Array(_dafny.Set({}), 25)
        d_2311_v29_ = nw413_
        d_2312_v30_: D22
        d_2312_v30_ = D22_DC54(d_2311_v29_)
        index345_ = default__.safeIndex(891, (d_2310_v28_).length(0))
        (d_2310_v28_)[index345_] = (d_2312_v30_).cf91
        index346_ = default__.safeIndex(891, (d_2310_v28_).length(0))
        (d_2310_v28_)[index346_] = d_2311_v29_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2276_v0_).length(0)):
            d_2313_i5_: int = guard_loop_5_
            if (True) and (((0) <= (d_2313_i5_)) and ((d_2313_i5_) < ((d_2276_v0_).length(0)))):
                (d_2276_v0_)[(d_2313_i5_)] = False
        d_2314_v31_: _dafny.Seq
        d_2314_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yciqas"))
        d_2315_v32_: D16
        d_2315_v32_ = D16_DC40(d_2314_v31_, p1)
        d_2316_v33_: D16
        d_2316_v33_ = D16_DC41(d_2315_v32_)
        d_2317_v34_: _dafny.MultiSet
        d_2317_v34_ = _dafny.MultiSet([d_2316_v33_, d_2316_v33_])
        r0 = ((_dafny.MultiSet([d_2316_v33_])).set(d_2316_v33_, default__.abs((self).f30))).ispropersubset((d_2317_v34_).set(d_2316_v33_, default__.abs(p1)))
        d_2318_v35_: _dafny.MultiSet
        d_2318_v35_ = _dafny.MultiSet([self.f28, self.f28, self.f28, (d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], True])
        r1 = D0_DC1(d_2318_v35_)
        d_2319_v36_: D0
        d_2319_v36_ = D0_DC0(not((d_2290_v10_) != (_dafny.Set({(d_2276_v0_)[default__.safeIndex(695, (d_2276_v0_).length(0))], self.f31, self.f28}))))
        r2 = d_2319_v36_
        r3 = self.f28
        return r0, r1, r2, r3

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_2320_v0_: _dafny.Set
        d_2320_v0_ = _dafny.Set({-512, (self).f30})
        d_2321_v1_: D7
        d_2321_v1_ = D7_DC19(p0, 608, (self).f30)
        d_2322_v2_: _dafny.Array
        nw414_ = _dafny.Array(None, 8)
        nw414_[int(0)] = (self).f34
        nw414_[int(1)] = (self).f30
        nw414_[int(2)] = (d_2321_v1_).cf38
        nw414_[int(3)] = (self).f30
        nw414_[int(4)] = (self).f34
        nw414_[int(5)] = ((self).f30 if default__.fm1(p0, (self).f30, globalState) else (self).f34)
        nw414_[int(6)] = (0) - ((self).f30)
        nw414_[int(7)] = 135
        d_2322_v2_ = nw414_
        index347_ = default__.safeIndex(458, (d_2322_v2_).length(0))
        (d_2322_v2_)[index347_] = (self).f30
        d_2323_v3_: _dafny.Map
        d_2323_v3_ = _dafny.Map({self.f35: len((self).f32)})
        d_2324_v4_: D5
        d_2324_v4_ = D5_DC13(p0, d_2323_v3_, self.f36)
        d_2325_v5_: _dafny.MultiSet
        d_2325_v5_ = _dafny.MultiSet([d_2323_v3_, d_2323_v3_])
        d_2326_v6_: _dafny.Seq
        d_2326_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vilf"))
        d_2327_v7_: _dafny.MultiSet
        d_2327_v7_ = _dafny.MultiSet([len(d_2326_v6_)])
        d_2328_v8_: _dafny.Map
        d_2328_v8_ = _dafny.Map({self.f28: p1})
        d_2329_v9_: _dafny.Seq
        d_2329_v9_ = _dafny.SeqWithoutIsStrInference([(self).f34, len(d_2328_v8_)])
        d_2330_v10_: _dafny.MultiSet
        d_2330_v10_ = _dafny.MultiSet([p1])
        pat_let_tv26_ = d_2320_v0_
        pat_let_tv27_ = d_2320_v0_
        pat_let_tv28_ = d_2320_v0_
        index348_ = default__.safeIndex(458, (d_2322_v2_).length(0))
        def lambda104_(source38_):
            if source38_.is_DC13:
                d_2331___mcc_h0_ = source38_.cf27
                d_2332___mcc_h1_ = source38_.cf28
                d_2333___mcc_h2_ = source38_.cf29
                d_2334_cf29_ = d_2333___mcc_h2_
                d_2335_cf28_ = d_2332___mcc_h1_
                d_2336_cf27_ = d_2331___mcc_h0_
                return _dafny.Set({(self).f30, -653})
            elif source38_.is_DC14:
                d_2337___mcc_h3_ = source38_.cf30
                d_2338___mcc_h4_ = source38_.cf31
                d_2339_cf31_ = d_2338___mcc_h4_
                d_2340_cf30_ = d_2337___mcc_h3_
                return pat_let_tv26_
            elif True:
                d_2341___mcc_h5_ = source38_.cf26
                d_2342_cf26_ = d_2341___mcc_h5_
                if self.f35:
                    return pat_let_tv27_
                elif True:
                    return pat_let_tv28_

        rhs384_ = p0
        rhs385_ = lambda104_(d_2324_v4_)
        rhs386_ = ((0) - (((d_2325_v5_).set((_dafny.Map({False: (self).f30})).set(False, (self).f34), default__.abs(len(_dafny.SeqWithoutIsStrInference([(self).fm7(default__.fm49(d_2327_v7_, self.f36, _dafny.SeqWithoutIsStrInference([d_2329_v9_]), globalState), (self).f30, globalState)]))))).cardinality)) + ((d_2330_v10_).cardinality)
        lhs310_ = self
        lhs311_ = d_2322_v2_
        lhs312_ = default__.safeIndex(458, (d_2322_v2_).length(0))
        lhs310_.f31 = rhs384_
        d_2320_v0_ = rhs385_
        lhs311_[lhs312_] = rhs386_
        if not(p0):
            d_2343_v11_: _dafny.Map
            d_2343_v11_ = _dafny.Map({self.f35: d_2329_v9_})
            d_2344_v12_: _dafny.Seq
            d_2344_v12_ = _dafny.SeqWithoutIsStrInference([(len(d_2343_v11_)) > (len((d_2329_v9_).set(default__.safeIndex((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], len(d_2329_v9_)), len(d_2326_v6_)))), self.f29])
            d_2344_v12_ = default__.fm0(((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]) - ((self).f34), self.f31, p0, ((self).f34) + (len((self).fm3(self.f29, self.f29, globalState))), globalState)
            (self).f29 = self.f35
            d_2345_v13_: _dafny.Seq
            d_2345_v13_ = _dafny.SeqWithoutIsStrInference([d_2320_v0_])
            d_2320_v0_ = ((d_2345_v13_)[default__.safeIndex((_dafny.MultiSet([self.f28, self.f28])).cardinality, len(d_2345_v13_))]).intersection((d_2320_v0_) | (d_2320_v0_))
            d_2346_v14_: _dafny.MultiSet
            d_2346_v14_ = _dafny.MultiSet([(d_2330_v10_).set(p0, default__.abs(178))])
            d_2347_v15_: _dafny.MultiSet
            d_2347_v15_ = d_2346_v14_
            source39_ = d_2347_v15_
            d_2348___mcc_h6_ = source39_
            d_2349_cf80_ = d_2348___mcc_h6_
            d_2350_v16_: _dafny.Array
            nw415_ = _dafny.Array(_dafny.MultiSet({}), 22)
            d_2350_v16_ = nw415_
            index349_ = default__.safeIndex(163, (d_2350_v16_).length(0))
            (d_2350_v16_)[index349_] = d_2330_v10_
            d_2351_v17_: _dafny.Seq
            d_2351_v17_ = _dafny.SeqWithoutIsStrInference([d_2330_v10_, d_2330_v10_])
            index350_ = default__.safeIndex(163, (d_2350_v16_).length(0))
            (d_2350_v16_)[index350_] = (d_2351_v17_)[default__.safeIndex(((self).f30) - ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]), len(d_2351_v17_))]
            (self).f36 = (self.f36 if self.f29 else self.f36)
            d_2352_v18_: _dafny.Array
            nw416_ = _dafny.Array(D7.default()(), 21)
            d_2352_v18_ = nw416_
            d_2353_v19_: _dafny.Array
            d_2353_v19_ = d_2352_v18_
            d_2353_v19_ = d_2353_v19_
            (globalState).f12 = p1
            d_2354_v20_: _dafny.Array
            def lambda105_(d_2355_i0_):
                return self.f29

            init69_ = lambda105_
            nw417_ = _dafny.Array(None, 11)
            for i0_69_ in range(nw417_.length(0)):
                nw417_[i0_69_] = init69_(i0_69_)
            d_2354_v20_ = nw417_
            d_2356_v21_: D6
            d_2356_v21_ = D6_DC16((D3_DC9(default__.fm1(p0, (self).f30, globalState), default__.fm1(self.f35, (self).f30, globalState), d_2354_v20_, False)).cf20)
            d_2357_v22_: D20
            d_2357_v22_ = D20_DC50(D20_DC49(self.f35, self.f36))
            d_2358_v23_: D3
            d_2358_v23_ = D3_DC9(default__.fm1(p1, (self).f34, globalState), False, d_2354_v20_, p0)
            d_2359_v24_: _dafny.Seq
            d_2359_v24_ = _dafny.SeqWithoutIsStrInference([(d_2358_v23_).cf20])
            d_2360_v25_: C1
            nw418_ = C1()
            nw418_.ctor__(self.f28, p1, p1)
            d_2360_v25_ = nw418_
            d_2361_v26_: D10
            d_2361_v26_ = D10_DC28(d_2360_v25_)
            pat_let_tv29_ = d_2360_v25_
            pat_let_tv30_ = d_2360_v25_
            d_2362_v27_: _dafny.Array
            nw419_ = _dafny.Array(None, 21)
            nw419_[int(0)] = d_2361_v26_
            nw419_[int(1)] = d_2361_v26_
            nw419_[int(2)] = D10_DC28(d_2360_v25_)
            nw419_[int(3)] = d_2361_v26_
            nw419_[int(4)] = d_2361_v26_
            nw419_[int(5)] = d_2361_v26_
            nw419_[int(6)] = d_2361_v26_
            nw419_[int(7)] = d_2361_v26_
            nw419_[int(8)] = D10_DC28(d_2360_v25_)
            nw419_[int(9)] = d_2361_v26_
            nw419_[int(10)] = D10_DC28(d_2360_v25_)
            nw419_[int(11)] = d_2361_v26_
            nw419_[int(12)] = d_2361_v26_
            def iife137_(_pat_let37_0):
                def iife138_(d_2363_dt__update__tmp_h0_):
                    def iife139_(_pat_let38_0):
                        def iife140_(d_2364_dt__update_hcf55_h0_):
                            return D10_DC28(d_2364_dt__update_hcf55_h0_)
                        return iife140_(_pat_let38_0)
                    return iife139_(pat_let_tv29_)
                return iife138_(_pat_let37_0)
            nw419_[int(13)] = iife137_(d_2361_v26_)
            def iife141_(_pat_let39_0):
                def iife142_(d_2365_dt__update__tmp_h1_):
                    def iife143_(_pat_let40_0):
                        def iife144_(d_2366_dt__update_hcf55_h1_):
                            return D10_DC28(d_2366_dt__update_hcf55_h1_)
                        return iife144_(_pat_let40_0)
                    return iife143_(pat_let_tv30_)
                return iife142_(_pat_let39_0)
            nw419_[int(14)] = iife141_(d_2361_v26_)
            nw419_[int(15)] = d_2361_v26_
            nw419_[int(16)] = d_2361_v26_
            nw419_[int(17)] = d_2361_v26_
            nw419_[int(18)] = d_2361_v26_
            nw419_[int(19)] = d_2361_v26_
            nw419_[int(20)] = D10_DC28(d_2360_v25_)
            d_2362_v27_ = nw419_
            d_2367_v28_: D20
            d_2367_v28_ = D20_DC48((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], d_2362_v27_, self.f28)
            d_2368_v29_: D20
            d_2368_v29_ = D20_DC50(d_2367_v28_)
            pat_let_tv31_ = d_2359_v24_
            pat_let_tv32_ = d_2359_v24_
            def iife145_(_pat_let41_0):
                def iife146_(d_2369_dt__update__tmp_h2_):
                    def iife147_(_pat_let42_0):
                        def iife148_(d_2370_dt__update_hcf33_h0_):
                            return D6_DC16(d_2370_dt__update_hcf33_h0_)
                        return iife148_(_pat_let42_0)
                    return iife147_((pat_let_tv31_)[default__.safeIndex((self).f34, len(pat_let_tv32_))])
                return iife146_(_pat_let41_0)
            rhs387_ = iife145_((d_2356_v21_ if p0 else d_2356_v21_))
            rhs388_ = (self.f36) in (d_2326_v6_)
            rhs389_ = (D20_DC50(d_2368_v29_) if True else d_2357_v22_)
            rhs390_ = default__.safeModuloInt((self).f34, -651)
            lhs313_ = globalState
            d_2356_v21_ = rhs387_
            lhs313_.f26 = rhs388_
            d_2357_v22_ = rhs389_
            r1 = rhs390_
        elif True:
            index351_ = default__.safeIndex(458, (d_2322_v2_).length(0))
            rhs391_ = d_2326_v6_
            rhs392_ = (0) - ((0) - ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]))
            lhs314_ = d_2322_v2_
            lhs315_ = default__.safeIndex(458, (d_2322_v2_).length(0))
            d_2326_v6_ = rhs391_
            lhs314_[lhs315_] = rhs392_
            d_2371_v30_: _dafny.Map
            d_2371_v30_ = _dafny.Map({self.f36: 717})
            d_2372_v31_: _dafny.Map
            d_2372_v31_ = _dafny.Map({d_2371_v30_: d_2326_v6_})
            (globalState).f1 = ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]) >= (default__.safeModuloInt(len(d_2372_v31_), (self).f34))
            if self.f35:
                d_2329_v9_ = ((d_2329_v9_) + (d_2329_v9_)) + ((d_2329_v9_) + (d_2329_v9_))
                d_2373_v32_: _dafny.Array
                def lambda106_(d_2374_i1_):
                    return True

                init70_ = lambda106_
                nw420_ = _dafny.Array(None, 6)
                for i0_70_ in range(nw420_.length(0)):
                    nw420_[i0_70_] = init70_(i0_70_)
                d_2373_v32_ = nw420_
                d_2375_v33_: _dafny.Map
                d_2375_v33_ = _dafny.Map({self.f36: d_2328_v8_})
                index352_ = default__.safeIndex(330, (d_2373_v32_).length(0))
                (d_2373_v32_)[index352_] = (default__.fm0((self).f34, True, p0, (self).f34, globalState)) < (default__.fm0(len(((d_2375_v33_)[_dafny.CodePoint('e')] if (_dafny.CodePoint('e')) in (d_2375_v33_) else d_2328_v8_)), False, p0, (self).f30, globalState))
                d_2376_v34_: _dafny.Map
                d_2376_v34_ = _dafny.Map({(self).f34: self.f36})
                d_2377_v35_: _dafny.Array
                def lambda107_(d_2378_i2_):
                    return self.f36

                init71_ = lambda107_
                nw421_ = _dafny.Array(None, 1)
                for i0_71_ in range(nw421_.length(0)):
                    nw421_[i0_71_] = init71_(i0_71_)
                d_2377_v35_ = nw421_
                d_2379_v36_: C0
                nw422_ = C0()
                nw422_.ctor__(d_2377_v35_)
                d_2379_v36_ = nw422_
                d_2380_v37_: D9
                d_2380_v37_ = D9_DC27(d_2379_v36_, d_2326_v6_)
                d_2381_v38_: _dafny.Map
                d_2381_v38_ = _dafny.Map({d_2380_v37_: self.f29})
                d_2382_v39_: C6
                nw423_ = C6()
                nw423_.ctor__(((d_2376_v34_)[default__.fm46(len(d_2381_v38_), (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], self.f28, globalState)] if (default__.fm46(len(d_2381_v38_), (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], self.f28, globalState)) in (d_2376_v34_) else self.f36), (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], not(not(p1)), (self).f32, (self).f33, (self).f30, p0, p0, self.f35)
                d_2382_v39_ = nw423_
                d_2383_v40_: _dafny.Set
                d_2383_v40_ = _dafny.Set({d_2382_v39_, d_2382_v39_})
                d_2384_v41_: _dafny.Map
                d_2384_v41_ = _dafny.Map({default__.safeDivisionInt(len(d_2383_v40_), (self).f30): p0})
                index353_ = default__.safeIndex(330, (d_2373_v32_).length(0))
                (d_2373_v32_)[index353_] = ((d_2384_v41_)[(self).f34] if ((self).f34) in (d_2384_v41_) else p0)
                (globalState).f21 = d_2373_v32_
                r2 = (0) - ((self).f30)
                d_2385_v42_: _dafny.Array
                nw424_ = _dafny.Array(None, 28)
                d_2385_v42_ = nw424_
                d_2386_v43_: _dafny.Map
                d_2386_v43_ = _dafny.Map({p1: d_2385_v42_})
                d_2386_v43_ = (d_2386_v43_).set(self.f29, d_2385_v42_)
            elif True:
                d_2387_v44_: _dafny.Array
                def lambda108_(d_2388_i3_):
                    return self.f28

                init72_ = lambda108_
                nw425_ = _dafny.Array(None, 18)
                for i0_72_ in range(nw425_.length(0)):
                    nw425_[i0_72_] = init72_(i0_72_)
                d_2387_v44_ = nw425_
                index354_ = default__.safeIndex(932, (d_2387_v44_).length(0))
                (d_2387_v44_)[index354_] = self.f28
                d_2389_v45_: _dafny.Seq
                d_2389_v45_ = _dafny.SeqWithoutIsStrInference([d_2322_v2_])
                index355_ = default__.safeIndex(932, (d_2387_v44_).length(0))
                rhs393_ = default__.fm1((self.f28) or (((self).f32)[default__.safeIndex((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], len((self).f32))]), 96, globalState)
                rhs394_ = (d_2389_v45_) < (d_2389_v45_)
                rhs395_ = default__.fm1(self.f29, default__.safeModuloInt((self).f34, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]), globalState)
                rhs396_ = self.f31
                rhs397_ = (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]
                lhs316_ = globalState
                lhs317_ = d_2387_v44_
                lhs318_ = default__.safeIndex(932, (d_2387_v44_).length(0))
                lhs319_ = globalState
                lhs320_ = self
                lhs321_ = globalState
                lhs316_.f26 = rhs393_
                lhs317_[lhs318_] = rhs394_
                lhs319_.f26 = rhs395_
                lhs320_.f28 = rhs396_
                lhs321_.f13 = rhs397_
                d_2390_v47_: _dafny.Set
                def iife149_():
                    coll63_ = _dafny.Set()
                    compr_63_: int
                    for compr_63_ in (d_2320_v0_).Elements:
                        d_2391_v46_: int = compr_63_
                        if (d_2391_v46_) in (d_2320_v0_):
                            coll63_ = coll63_.union(_dafny.Set([default__.safeModuloInt(d_2391_v46_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), 672, len(_dafny.SeqWithoutIsStrInference([not(not(True))])), 866])))]))
                    return _dafny.Set(coll63_)
                d_2390_v47_ = iife149_()
                
                d_2392_v48_: _dafny.Array
                nw426_ = _dafny.Array(None, 7)
                nw426_[int(0)] = d_2390_v47_
                nw426_[int(1)] = d_2390_v47_
                nw426_[int(2)] = d_2320_v0_
                nw426_[int(3)] = d_2390_v47_
                nw426_[int(4)] = (d_2390_v47_ if False else d_2390_v47_)
                nw426_[int(5)] = d_2390_v47_
                nw426_[int(6)] = d_2390_v47_
                d_2392_v48_ = nw426_
                index356_ = default__.safeIndex(830, (d_2392_v48_).length(0))
                (d_2392_v48_)[index356_] = d_2390_v47_
                index357_ = default__.safeIndex(830, (d_2392_v48_).length(0))
                rhs398_ = 292
                rhs399_ = d_2326_v6_
                rhs400_ = not(True)
                rhs401_ = default__.fm63((self.f28) and ((d_2387_v44_)[default__.safeIndex(932, (d_2387_v44_).length(0))]), p1, p0, globalState)
                lhs322_ = globalState
                lhs323_ = globalState
                lhs324_ = d_2392_v48_
                lhs325_ = default__.safeIndex(830, (d_2392_v48_).length(0))
                lhs322_.f13 = rhs398_
                d_2326_v6_ = rhs399_
                lhs323_.f1 = rhs400_
                lhs324_[lhs325_] = rhs401_
                d_2393_v49_: _dafny.Seq
                d_2393_v49_ = _dafny.SeqWithoutIsStrInference([((self).f32) + ((self).f32)])
                index358_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                rhs402_ = 759
                rhs403_ = ((self).f30) + (((self).f30) * ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]))
                rhs404_ = (self).f34
                rhs405_ = self.f36
                rhs406_ = (d_2393_v49_) + ((self).f33)
                lhs326_ = globalState
                lhs327_ = d_2322_v2_
                lhs328_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                lhs329_ = self
                r1 = rhs402_
                lhs326_.f13 = rhs403_
                lhs327_[lhs328_] = rhs404_
                lhs329_.f36 = rhs405_
                d_2393_v49_ = rhs406_
                d_2394_v50_: _dafny.Seq
                d_2394_v50_ = _dafny.SeqWithoutIsStrInference([True])
                d_2394_v50_ = d_2394_v50_
                d_2395_v51_: D4
                d_2395_v51_ = D4_DC11(len(d_2328_v8_), (d_2387_v44_)[default__.safeIndex(932, (d_2387_v44_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tuwhsf"))))
                index359_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                (d_2322_v2_)[index359_] = (((d_2395_v51_).cf25) - ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))])) + ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))])
            d_2396_v52_: D0
            d_2396_v52_ = D0_DC0(p1)
            d_2397_v53_: _dafny.Set
            d_2397_v53_ = _dafny.Set({p0, (d_2396_v52_).cf0})
            d_2398_v54_: _dafny.Seq
            d_2398_v54_ = _dafny.SeqWithoutIsStrInference([d_2397_v53_])
            d_2399_v55_: bool
            d_2400_v56_: D0
            d_2401_v57_: D0
            d_2402_v58_: bool
            out73_: bool
            out74_: D0
            out75_: D0
            out76_: bool
            out73_, out74_, out75_, out76_ = (self).m2((self).f30, (self).f34, (self).f30, d_2398_v54_, globalState)
            d_2399_v55_ = out73_
            d_2400_v56_ = out74_
            d_2401_v57_ = out75_
            d_2402_v58_ = out76_
            d_2403_v59_: _dafny.Array
            nw427_ = _dafny.Array(False, 25)
            d_2403_v59_ = nw427_
            index360_ = default__.safeIndex(104, (d_2403_v59_).length(0))
            (d_2403_v59_)[index360_] = self.f28
            d_2404_v60_: _dafny.Map
            d_2404_v60_ = _dafny.Map({(self).f34: ((d_2330_v10_)[self.f28] if (self.f28) in (d_2330_v10_) else (self).f34)})
            index361_ = default__.safeIndex(104, (d_2403_v59_).length(0))
            (d_2403_v59_)[index361_] = ((self).f32)[default__.safeIndex((len(d_2404_v60_)) + ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]), len((self).f32))]
        if p1:
            r0 = p0
            if self.f31:
                d_2405_v61_: _dafny.Array
                def lambda109_(d_2406_i4_):
                    return (self.f28) == (self.f29)

                init73_ = lambda109_
                nw428_ = _dafny.Array(None, 25)
                for i0_73_ in range(nw428_.length(0)):
                    nw428_[i0_73_] = init73_(i0_73_)
                d_2405_v61_ = nw428_
                (globalState).f21 = d_2405_v61_
                d_2407_v62_: D16
                d_2407_v62_ = D16_DC40(d_2326_v6_, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))])
                d_2408_v63_: T1
                nw429_ = C7()
                nw429_.ctor__((self).f34, (d_2327_v7_).intersection(d_2327_v7_), (self).f30, ((d_2407_v62_).cf73) < (d_2326_v6_), self.f29, self.f28)
                d_2408_v63_ = nw429_
                d_2408_v63_ = d_2408_v63_
                d_2409_v64_: _dafny.MultiSet
                d_2409_v64_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([self.f36])])
                index362_ = default__.safeIndex(872, (d_2405_v61_).length(0))
                (d_2405_v61_)[index362_] = (d_2409_v64_).ispropersubset(d_2409_v64_)
                d_2410_v65_: D3
                d_2410_v65_ = D3_DC9(d_2408_v63_.f28, self.f31, d_2405_v61_, p1)
                d_2411_v66_: _dafny.Set
                d_2411_v66_ = _dafny.Set({not(p1), self.f28, not(d_2408_v63_.f28)})
                d_2412_v67_: _dafny.Map
                d_2412_v67_ = _dafny.Map({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]: self.f31})
                d_2413_v69_: _dafny.Set
                def iife150_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in (_dafny.Set({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]})).Elements:
                        d_2414_v68_: int = compr_64_
                        if (d_2414_v68_) in (_dafny.Set({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]})):
                            coll64_[(d_2414_v68_) * ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))])] = d_2408_v63_.f28
                    return _dafny.Map(coll64_)
                d_2413_v69_ = _dafny.Set({d_2412_v67_, iife150_()
                })
                d_2415_v70_: _dafny.Seq
                d_2415_v70_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({self.f35, False})) != (d_2411_v66_), (d_2413_v69_).isdisjoint(d_2413_v69_), not(True)])
                d_2416_v71_: _dafny.Array
                nw430_ = _dafny.Array(None, 8)
                nw430_[int(0)] = self.f36
                nw430_[int(1)] = _dafny.CodePoint('a')
                nw430_[int(2)] = self.f36
                nw430_[int(3)] = _dafny.CodePoint('s')
                nw430_[int(4)] = self.f36
                nw430_[int(5)] = self.f36
                nw430_[int(6)] = (d_2326_v6_)[default__.safeIndex((self).f34, len(d_2326_v6_))]
                nw430_[int(7)] = self.f36
                d_2416_v71_ = nw430_
                d_2417_v72_: _dafny.Seq
                d_2417_v72_ = _dafny.SeqWithoutIsStrInference([(self).f33, _dafny.SeqWithoutIsStrInference([(self).f32 for d_2418_i5_ in range(default__.abs(414))])])
                d_2419_v73_: T3
                nw431_ = C5()
                nw431_.ctor__(d_2408_v63_.f29, (self).f34, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], p1, (self).f32, _dafny.SeqWithoutIsStrInference([(self).f32 for d_2420_i6_ in range(default__.abs(210))]), (d_2408_v63_).f30, not(p0), self.f31, p1)
                d_2419_v73_ = nw431_
                d_2421_v74_: _dafny.Map
                d_2421_v74_ = _dafny.Map({d_2419_v73_: len(d_2323_v3_)})
                d_2422_v75_: T3
                nw432_ = C5()
                nw432_.ctor__(self.f35, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], (d_2408_v63_).f30, self.f35, d_2415_v70_, (d_2417_v72_)[default__.safeIndex(len(d_2421_v74_), len(d_2417_v72_))], len(d_2415_v70_), d_2419_v73_.f28, p0, p1)
                d_2422_v75_ = nw432_
                d_2423_v76_: _dafny.Map
                d_2423_v76_ = _dafny.Map({d_2416_v71_: d_2422_v75_})
                pat_let_tv33_ = d_2405_v61_
                pat_let_tv34_ = d_2405_v61_
                index363_ = default__.safeIndex(872, (d_2405_v61_).length(0))
                def iife151_(_pat_let43_0):
                    def iife152_(d_2424_dt__update__tmp_h3_):
                        def iife153_(_pat_let44_0):
                            def iife154_(d_2425_dt__update_hcf20_h0_):
                                return D3_DC9((d_2424_dt__update__tmp_h3_).cf18, (d_2424_dt__update__tmp_h3_).cf19, d_2425_dt__update_hcf20_h0_, (d_2424_dt__update__tmp_h3_).cf21)
                            return iife154_(_pat_let44_0)
                        return iife153_((pat_let_tv33_ if self.f29 else pat_let_tv34_))
                    return iife152_(_pat_let43_0)
                rhs407_ = (self.f29) or (p0)
                rhs408_ = iife151_(d_2410_v65_)
                rhs409_ = default__.safeDivisionInt((self).f34, len(d_2423_v76_))
                rhs410_ = (d_2422_v75_).f32
                lhs330_ = d_2405_v61_
                lhs331_ = default__.safeIndex(872, (d_2405_v61_).length(0))
                lhs330_[lhs331_] = rhs407_
                d_2410_v65_ = rhs408_
                r2 = rhs409_
                d_2415_v70_ = rhs410_
                index364_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                (d_2322_v2_)[index364_] = len((d_2411_v66_).intersection((_dafny.Set({(d_2405_v61_)[default__.safeIndex(872, (d_2405_v61_).length(0))], d_2408_v63_.f29, d_2408_v63_.f31, d_2408_v63_.f29, (self).fm7((d_2327_v7_).set((0) - ((self).f34), default__.abs(154)), (d_2422_v75_).f34, globalState)})) - (d_2411_v66_)))
                d_2426_v77_: C4
                nw433_ = C4()
                nw433_.ctor__((d_2422_v75_).f30, self.f28, (d_2422_v75_).f32, (d_2419_v73_).f33, len(d_2412_v67_), d_2419_v73_.f31, d_2422_v75_.f31, self.f35)
                d_2426_v77_ = nw433_
                d_2427_v78_: _dafny.MultiSet
                d_2427_v78_ = _dafny.MultiSet([d_2426_v77_, d_2426_v77_])
                d_2428_v79_: _dafny.Array
                nw434_ = _dafny.Array(D10.default()(), 6)
                d_2428_v79_ = nw434_
                d_2429_v80_: D20
                d_2429_v80_ = D20_DC48((d_2427_v78_).cardinality, d_2428_v79_, d_2422_v75_.f29)
                (globalState).f13 = (d_2429_v80_).cf82
            elif True:
                d_2430_v81_: _dafny.Map
                d_2430_v81_ = _dafny.Map({d_2330_v10_: p1})
                d_2431_v82_: _dafny.Array
                def lambda110_(d_2432_i7_):
                    return (self).f32

                init74_ = lambda110_
                nw435_ = _dafny.Array(None, 1)
                for i0_74_ in range(nw435_.length(0)):
                    nw435_[i0_74_] = init74_(i0_74_)
                d_2431_v82_ = nw435_
                d_2433_v83_: _dafny.Set
                d_2433_v83_ = _dafny.Set({((d_2328_v8_)[p0] if (p0) in (d_2328_v8_) else self.f29), self.f28})
                d_2434_v84_: C3
                nw436_ = C3()
                nw436_.ctor__(len(d_2433_v83_), d_2431_v82_, self.f36, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], False, (self).f32, (self).f33, len(d_2320_v0_), default__.fm1(p1, (self).f30, globalState), True, p1)
                d_2434_v84_ = nw436_
                d_2435_v85_: _dafny.Map
                d_2435_v85_ = _dafny.Map({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]: d_2434_v84_})
                d_2436_v86_: _dafny.Seq
                d_2436_v86_ = _dafny.SeqWithoutIsStrInference([d_2434_v84_, ((d_2435_v85_)[(self).f30] if ((self).f30) in (d_2435_v85_) else d_2434_v84_)])
                d_2437_v87_: _dafny.Map
                d_2437_v87_ = _dafny.Map({len(_dafny.Map({self.f28: self.f36})): (self).f34})
                d_2438_v88_: C3
                nw437_ = C3()
                nw437_.ctor__(len(d_2430_v81_), d_2431_v82_, self.f36, ((self).f34) - ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]), (self).fm7(_dafny.MultiSet([(0) - ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))])]), len(d_2436_v86_), globalState), (self).f32, (self).f33, ((d_2437_v87_)[(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]] if ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]) in (d_2437_v87_) else (d_2434_v84_).f44), self.f28, ((self).f34) <= ((d_2434_v84_).f44), (d_2330_v10_).ispropersubset(d_2330_v10_))
                d_2438_v88_ = nw437_
                d_2439_v89_: C4
                nw438_ = C4()
                nw438_.ctor__(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]])), (d_2434_v84_).f44), (213) <= (452), (self).f32, (_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32, (self).f32])) + ((self).f33), (d_2434_v84_).f44, not (False) or (p1), self.f28, p1)
                d_2439_v89_ = nw438_
                index365_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                (d_2322_v2_)[index365_] = (0) - ((d_2434_v84_).f44)
                (globalState).f26 = (((self).f30) == ((self).f34)) == ((d_2439_v89_).fm6(globalState))
                d_2440_v90_: _dafny.Seq
                d_2440_v90_ = _dafny.SeqWithoutIsStrInference([d_2327_v7_, d_2327_v7_])
                d_2441_v91_: T1
                nw439_ = C7()
                nw439_.ctor__((d_2434_v84_).f44, (d_2440_v90_)[default__.safeIndex((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], len(d_2440_v90_))], len((d_2326_v6_).set(default__.safeIndex((d_2434_v84_).f44, len(d_2326_v6_)), self.f36)), self.f28, p1, (d_2439_v89_).fm44((self).f30, globalState))
                d_2441_v91_ = nw439_
                nw440_ = C7()
                nw440_.ctor__((0) - (default__.safeDivisionInt((d_2441_v91_).f30, (d_2438_v88_).f44)), default__.fm25(globalState), (d_2441_v91_).f30, True, ((d_2434_v84_).f44) != ((d_2438_v88_).f44), d_2441_v91_.f28)
                d_2441_v91_ = nw440_
            d_2328_v8_ = (d_2328_v8_).set(p0, not(not(self.f31)))
            (globalState).f13 = (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]
            d_2442_v92_: D0
            d_2442_v92_ = D0_DC1(d_2330_v10_)
            d_2443_v93_: _dafny.Map
            d_2443_v93_ = _dafny.Map({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]: d_2442_v92_})
            d_2444_v94_: _dafny.Set
            d_2444_v94_ = _dafny.Set({self.f28})
            d_2445_v96_: _dafny.Array
            nw441_ = _dafny.Array(None, 15)
            nw441_[int(0)] = (d_2443_v93_).set(len(d_2444_v94_), d_2442_v92_)
            nw441_[int(1)] = d_2443_v93_
            nw441_[int(2)] = d_2443_v93_
            nw441_[int(3)] = d_2443_v93_
            nw441_[int(4)] = default__.fm64((self).f30, self.f35, (0) - ((self).f34), self.f28, globalState)
            nw441_[int(5)] = (d_2443_v93_ if self.f31 else d_2443_v93_)
            nw441_[int(6)] = d_2443_v93_
            nw441_[int(7)] = _dafny.Map({(0) - ((0) - ((self).f34)): d_2442_v92_})
            nw441_[int(8)] = d_2443_v93_
            nw441_[int(9)] = d_2443_v93_
            def iife155_():
                coll65_ = _dafny.Map()
                compr_65_: int
                for compr_65_ in (d_2329_v9_).Elements:
                    d_2446_v95_: int = compr_65_
                    if (d_2446_v95_) in (d_2329_v9_):
                        coll65_[(d_2446_v95_) - ((self).f30)] = D0_DC1(_dafny.MultiSet([self.f28]))
                return _dafny.Map(coll65_)
            nw441_[int(10)] = iife155_()
            
            nw441_[int(11)] = (d_2443_v93_) | (d_2443_v93_)
            nw441_[int(12)] = d_2443_v93_
            nw441_[int(13)] = (d_2443_v93_) | (d_2443_v93_)
            nw441_[int(14)] = (d_2443_v93_) | (d_2443_v93_)
            d_2445_v96_ = nw441_
            index366_ = default__.safeIndex(549, (d_2445_v96_).length(0))
            (d_2445_v96_)[index366_] = d_2443_v93_
            d_2447_v97_: D8
            d_2447_v97_ = D8_DC22((self).f30, len((d_2326_v6_).set(default__.safeIndex((self).f34, len(d_2326_v6_)), self.f36)))
            d_2448_v98_: _dafny.Map
            d_2448_v98_ = _dafny.Map({self.f28: d_2447_v97_})
            d_2449_v100_: _dafny.Map
            d_2449_v100_ = _dafny.Map({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]: p0})
            index367_ = default__.safeIndex(549, (d_2445_v96_).length(0))
            def iife156_():
                coll66_ = _dafny.Map()
                compr_66_: int
                for compr_66_ in (d_2449_v100_).keys.Elements:
                    d_2450_v99_: int = compr_66_
                    if (d_2450_v99_) in (d_2449_v100_):
                        coll66_[default__.safeDivisionInt(d_2450_v99_, (self).f30)] = d_2442_v92_
                return _dafny.Map(coll66_)
            rhs411_ = (self).f34
            rhs412_ = (d_2443_v93_) | (iife156_()
            )
            rhs413_ = (self).f34
            rhs414_ = (d_2448_v98_ if p1 else (default__.fm65((self).f34, self.f35, (self).f34, globalState)) | ((_dafny.Map({self.f35: d_2447_v97_})).set(self.f35, D8_DC22((self).f30, (self).f34))))
            rhs415_ = self.f28
            lhs332_ = d_2445_v96_
            lhs333_ = default__.safeIndex(549, (d_2445_v96_).length(0))
            lhs334_ = globalState
            lhs335_ = globalState
            r1 = rhs411_
            lhs332_[lhs333_] = rhs412_
            lhs334_.f13 = rhs413_
            d_2448_v98_ = rhs414_
            lhs335_.f1 = rhs415_
        elif True:
            if self.f28:
                (globalState).f13 = (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]
                (self).f36 = (d_2326_v6_)[default__.safeIndex((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], len(d_2326_v6_))]
                def lambda111_(d_2451_p1_):
                    def lambda112_(d_2452_i8_):
                        return d_2451_p1_

                    return lambda112_

                init75_ = lambda111_(p1)
                nw442_ = _dafny.Array(None, 23)
                for i0_75_ in range(nw442_.length(0)):
                    nw442_[i0_75_] = init75_(i0_75_)
                (globalState).f21 = nw442_
                d_2453_v101_: D0
                d_2453_v101_ = D0_DC1(d_2330_v10_)
                d_2454_v102_: _dafny.Map
                d_2454_v102_ = _dafny.Map({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]: self.f31})
                d_2455_v103_: _dafny.Seq
                d_2455_v103_ = _dafny.SeqWithoutIsStrInference([d_2454_v102_])
                d_2456_v104_: _dafny.Set
                d_2456_v104_ = _dafny.Set({False})
                d_2457_v105_: D2
                d_2457_v105_ = D2_DC6((self).f34, self.f35, len(d_2456_v104_), self.f35, len(_dafny.Map({self.f31: (self).f30})))
                index368_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                rhs416_ = default__.fm15(len(d_2329_v9_), len(d_2455_v103_), globalState)
                rhs417_ = (_dafny.Map({self.f35: (d_2457_v105_).cf10})) | (d_2328_v8_)
                rhs418_ = 121
                rhs419_ = ((d_2326_v6_) + (d_2326_v6_)) <= (d_2326_v6_)
                lhs336_ = d_2322_v2_
                lhs337_ = default__.safeIndex(458, (d_2322_v2_).length(0))
                d_2453_v101_ = rhs416_
                d_2328_v8_ = rhs417_
                lhs336_[lhs337_] = rhs418_
                r3 = rhs419_
                (self).f36 = self.f36
            elif True:
                d_2458_v106_: _dafny.Map
                d_2458_v106_ = _dafny.Map({d_2322_v2_: (self).fm4(p0, d_2320_v0_, (self).f30, p1, globalState)})
                d_2459_v107_: _dafny.Map
                d_2459_v107_ = _dafny.Map({(self).f30: self.f31})
                d_2460_v108_: T2
                nw443_ = C8()
                nw443_.ctor__(p1, self.f36, len(d_2326_v6_), self.f31, (self).f32, (self).f33, len(d_2329_v9_), False, self.f28, self.f35)
                d_2460_v108_ = nw443_
                d_2461_v109_: _dafny.MultiSet
                d_2461_v109_ = _dafny.MultiSet([d_2460_v108_, d_2460_v108_, d_2460_v108_, d_2460_v108_])
                d_2462_v110_: T1
                nw444_ = C7()
                nw444_.ctor__((len(d_2458_v106_)) - ((d_2327_v7_).cardinality), (d_2327_v7_).intersection(d_2327_v7_), (0) - (len(_dafny.SeqWithoutIsStrInference([self.f36 for d_2463_i9_ in range(default__.abs(-618))]))), ((d_2459_v107_)[len((d_2326_v6_).set(default__.safeIndex((0) - ((self).f30), len(d_2326_v6_)), self.f36))] if (len((d_2326_v6_).set(default__.safeIndex((0) - ((self).f30), len(d_2326_v6_)), self.f36))) in (d_2459_v107_) else self.f28), (self.f31) in ((self).f32), (d_2461_v109_).isdisjoint(d_2461_v109_))
                d_2462_v110_ = nw444_
                d_2462_v110_ = d_2462_v110_
                d_2464_v111_: _dafny.Array
                def lambda113_(d_2465_i10_):
                    return True

                init76_ = lambda113_
                nw445_ = _dafny.Array(None, 6)
                for i0_76_ in range(nw445_.length(0)):
                    nw445_[i0_76_] = init76_(i0_76_)
                d_2464_v111_ = nw445_
                index369_ = default__.safeIndex(527, (d_2464_v111_).length(0))
                (d_2464_v111_)[index369_] = p1
                index370_ = default__.safeIndex(527, (d_2464_v111_).length(0))
                (d_2464_v111_)[index370_] = ((d_2462_v110_.f31) or (not(d_2462_v110_.f28))) and (d_2462_v110_.f31)
                d_2466_v112_: C4
                nw446_ = C4()
                nw446_.ctor__(default__.safeDivisionInt((D10_DC31(d_2460_v108_.f29, self.f28, (d_2462_v110_).f30)).cf62, 754), True, (self).f32, (d_2460_v108_).f33, (d_2460_v108_).f30, (d_2464_v111_)[default__.safeIndex(527, (d_2464_v111_).length(0))], (d_2323_v3_) in (d_2325_v5_), not(self.f31))
                d_2466_v112_ = nw446_
                d_2467_v113_: _dafny.Seq
                d_2467_v113_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(d_2464_v111_)[default__.safeIndex(527, (d_2464_v111_).length(0))]])).set(d_2462_v110_.f31, default__.abs(default__.fm24(d_2462_v110_.f31, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], globalState))), (_dafny.MultiSet([d_2460_v108_.f28, d_2462_v110_.f29])).intersection((_dafny.MultiSet([self.f28])).set(d_2460_v108_.f28, default__.abs((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]))), (d_2330_v10_).set(d_2462_v110_.f29, default__.abs((d_2460_v108_).f30))])
                d_2467_v113_ = _dafny.SeqWithoutIsStrInference([(d_2330_v10_ if d_2462_v110_.f29 else (d_2330_v10_).set(False, default__.abs((self).f30)))])
                d_2468_v114_: T3
                nw447_ = C4()
                nw447_.ctor__(default__.safeDivisionInt((d_2460_v108_).f30, (self).f30), p1, (self).f32, (self).f33, (0) - ((self).f30), d_2460_v108_.f28, self.f29, False)
                d_2468_v114_ = nw447_
                rhs420_ = d_2468_v114_
                rhs421_ = d_2326_v6_
                d_2468_v114_ = rhs420_
                d_2326_v6_ = rhs421_
            d_2469_v115_: _dafny.Array
            nw448_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
            d_2469_v115_ = nw448_
            index371_ = default__.safeIndex(349, (d_2469_v115_).length(0))
            (d_2469_v115_)[index371_] = d_2326_v6_
            d_2470_v116_: D16
            d_2470_v116_ = D16_DC40((self).fm3(self.f29, p0, globalState), (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))])
            index372_ = default__.safeIndex(349, (d_2469_v115_).length(0))
            (d_2469_v115_)[index372_] = (d_2470_v116_).cf73
            (globalState).f1 = self.f31
            d_2471_v117_: _dafny.Array
            nw449_ = _dafny.Array(False, 24)
            d_2471_v117_ = nw449_
            index373_ = default__.safeIndex(904, (d_2471_v117_).length(0))
            (d_2471_v117_)[index373_] = ((self).f32)[default__.safeIndex((self).f34, len((self).f32))]
            d_2472_v118_: _dafny.Array
            def lambda114_(d_2473_i11_):
                return self.f36

            init77_ = lambda114_
            nw450_ = _dafny.Array(None, 21)
            for i0_77_ in range(nw450_.length(0)):
                nw450_[i0_77_] = init77_(i0_77_)
            d_2472_v118_ = nw450_
            index374_ = default__.safeIndex(786, (d_2472_v118_).length(0))
            (d_2472_v118_)[index374_] = self.f36
            d_2474_v119_: D0
            d_2474_v119_ = D0_DC1(d_2330_v10_)
            d_2475_v120_: _dafny.Map
            d_2475_v120_ = _dafny.Map({-393: self.f28})
            pat_let_tv35_ = d_2330_v10_
            index375_ = default__.safeIndex(904, (d_2471_v117_).length(0))
            index376_ = default__.safeIndex(786, (d_2472_v118_).length(0))
            index377_ = default__.safeIndex(458, (d_2322_v2_).length(0))
            def iife157_(_pat_let45_0):
                def iife158_(d_2476_dt__update__tmp_h4_):
                    def iife159_(_pat_let46_0):
                        def iife160_(d_2477_dt__update_hcf1_h0_):
                            return D0_DC1(d_2477_dt__update_hcf1_h0_)
                        return iife160_(_pat_let46_0)
                    return iife159_(pat_let_tv35_)
                return iife158_(_pat_let45_0)
            rhs422_ = self.f28
            rhs423_ = default__.fm14((self).f34, not (self.f29) or (p0), default__.fm39(self.f28, _dafny.MultiSet([(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]]), self.f35, globalState), iife157_(d_2474_v119_), globalState)
            rhs424_ = (len((_dafny.SeqWithoutIsStrInference([self.f36 for d_2478_i12_ in range(default__.abs(614))])) + ((d_2469_v115_)[default__.safeIndex(349, (d_2469_v115_).length(0))]))) + ((self).f34)
            rhs425_ = (self).f34
            rhs426_ = (len(((d_2475_v120_).set(743, False)).set((0) - ((self).f30), self.f35))) + ((self).f34)
            lhs338_ = d_2471_v117_
            lhs339_ = default__.safeIndex(904, (d_2471_v117_).length(0))
            lhs340_ = d_2472_v118_
            lhs341_ = default__.safeIndex(786, (d_2472_v118_).length(0))
            lhs342_ = d_2322_v2_
            lhs343_ = default__.safeIndex(458, (d_2322_v2_).length(0))
            lhs338_[lhs339_] = rhs422_
            lhs340_[lhs341_] = rhs423_
            lhs342_[lhs343_] = rhs424_
            r1 = rhs425_
            r1 = rhs426_
            d_2322_v2_ = d_2322_v2_
        d_2479_v121_: C5
        nw451_ = C5()
        nw451_.ctor__(p1, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], ((d_2327_v7_)[(self).f34] if ((self).f34) in (d_2327_v7_) else 444), self.f28, (self).f32, (self).f33, (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))], self.f31, self.f31, not(self.f29))
        d_2479_v121_ = nw451_
        d_2480_v122_: _dafny.Set
        d_2480_v122_ = d_2320_v0_
        d_2480_v122_ = (d_2480_v122_ if not(p1) else d_2480_v122_)
        d_2481_v123_: _dafny.Map
        d_2481_v123_ = _dafny.Map({(d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]: (d_2328_v8_).set(p0, p0)})
        d_2482_v124_: _dafny.Array
        nw452_ = _dafny.Array(_dafny.Array(None, 0), 25)
        d_2482_v124_ = nw452_
        d_2483_v125_: _dafny.Map
        d_2483_v125_ = _dafny.Map({d_2482_v124_: (self).f33})
        d_2484_v126_: D14
        d_2484_v126_ = D14_DC36((self).f33)
        d_2485_v127_: _dafny.Set
        d_2485_v127_ = _dafny.Set({p1})
        d_2486_v128_: C9
        nw453_ = C9()
        nw453_.ctor__((0) - ((d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]), d_2481_v123_, self.f36, len((self).f32), ((self).f34) <= ((self).f30), (self).f32, (((d_2483_v125_)[d_2482_v124_] if (d_2482_v124_) in (d_2483_v125_) else (d_2484_v126_).cf67)) + (_dafny.SeqWithoutIsStrInference([(self).f32 for d_2487_i13_ in range(default__.abs(890))])), (self).f30, (d_2479_v121_).f47, (d_2485_v127_).issubset(default__.fm11((self).f33, p1, (d_2479_v121_).f47, 195, globalState)), (d_2321_v1_) in (_dafny.SeqWithoutIsStrInference([d_2321_v1_])))
        d_2486_v128_ = nw453_
        r0 = (d_2479_v121_).fm6(globalState)
        r1 = ((d_2323_v3_)[(d_2479_v121_).f47] if ((d_2479_v121_).f47) in (d_2323_v3_) else (len((self).f32)) - ((self).f34))
        r2 = (d_2322_v2_)[default__.safeIndex(458, (d_2322_v2_).length(0))]
        d_2488_v129_: _dafny.Set
        d_2488_v129_ = _dafny.Set({d_2322_v2_, d_2322_v2_})
        r3 = (self).fm7(_dafny.MultiSet(d_2329_v9_), len((d_2488_v129_) - (d_2488_v129_)), globalState)
        return r0, r1, r2, r3

