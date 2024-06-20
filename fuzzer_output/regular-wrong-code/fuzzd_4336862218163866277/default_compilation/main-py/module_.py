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
        return False

    @staticmethod
    def fm1(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bur"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "noxdr")))

    @staticmethod
    def fm2(p0, p1, globalState):
        return default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iu")))) * (694), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return _dafny.Set({(_dafny.Map({False: 90})) | (_dafny.Map({False: -787}))})

    @staticmethod
    def fm11(globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({False: True}))) | ((_dafny.Map({True: True})) | (_dafny.Map({False: not(True)})))

    @staticmethod
    def fm12(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True]) if True else _dafny.SeqWithoutIsStrInference([False]))) + (_dafny.SeqWithoutIsStrInference([False, True]))

    @staticmethod
    def fm15(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgdqdqqcq")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcmmbgxd")))

    @staticmethod
    def fm16(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxsrnkgpo"))])).Elements:
                d_0_v0_: _dafny.Seq = compr_0_
                if (d_0_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxsrnkgpo"))])):
                    coll0_[d_0_v0_] = False
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1_i0_ in range(default__.abs(283))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukufam"))])).Elements:
                d_2_v1_: _dafny.Seq = compr_1_
                if (d_2_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1_i0_ in range(default__.abs(283))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukufam"))])):
                    coll1_[d_2_v1_] = not(not(False))
            return _dafny.Map(coll1_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rchljvs")): True})) | (iife0_()
        )) | (iife1_()
        )

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        source0_ = D2_DC3(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djvdjhw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqfwcrmm"))]))
        if source0_.is_DC4:
            d_3___mcc_h0_ = source0_.cf6
            d_4___mcc_h1_ = source0_.cf7
            d_5___mcc_h2_ = source0_.cf8
            d_6_cf8_ = d_5___mcc_h2_
            d_7_cf7_ = d_4___mcc_h1_
            d_8_cf6_ = d_3___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_6_cf8_]))])) + (_dafny.SeqWithoutIsStrInference([d_6_cf8_, (0) - (len(_dafny.Map({_dafny.Map({len(_dafny.Map({False: len(_dafny.Set({True}))})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfcbcpfd")))}): False})))]))
        elif source0_.is_DC5:
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.Map({119: (D8_DC16(True, (D16_DC39(len(_dafny.SeqWithoutIsStrInference([True, False])), True, _dafny.Map({True: True}), _dafny.CodePoint('s'))).cf62)).cf24})).keys.Elements:
                    d_9_v0_: int = compr_2_
                    if (d_9_v0_) in (_dafny.Map({119: (D8_DC16(True, (D16_DC39(len(_dafny.SeqWithoutIsStrInference([True, False])), True, _dafny.Map({True: True}), _dafny.CodePoint('s'))).cf62)).cf24})):
                        coll2_[default__.safeDivisionInt(d_9_v0_, len(_dafny.SeqWithoutIsStrInference([333])))] = _dafny.MultiSet([807])
                return _dafny.Map(coll2_)
            return _dafny.SeqWithoutIsStrInference([len(iife2_()
            )])
        elif True:
            d_10___mcc_h3_ = source0_.cf5
            d_11_cf5_ = d_10___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_12_i1_ in range(default__.abs(952))]))) for d_13_i0_ in range(default__.abs(732))])

    @staticmethod
    def fm18(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_14_i0_ in range(default__.abs(177))]))])).Elements:
                d_15_v0_: int = compr_3_
                if (d_15_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_14_i0_ in range(default__.abs(177))]))])):
                    coll3_ = coll3_.union(_dafny.Set([(d_15_v0_) + (443)]))
            return _dafny.Set(coll3_)
        return ((_dafny.Set({950})).intersection(_dafny.Set({len(iife3_()
        )}))) | ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('k')]))})) | (_dafny.Set({len(_dafny.Set({not(False)})), 778})))

    @staticmethod
    def fm19(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([False])

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        source1_ = D22_DC51(_dafny.Set({103}))
        if source1_.is_DC52:
            return _dafny.CodePoint('q')
        elif source1_.is_DC53:
            d_16___mcc_h0_ = source1_.cf87
            d_17___mcc_h1_ = source1_.cf88
            d_18___mcc_h2_ = source1_.cf89
            d_19___mcc_h3_ = source1_.cf90
            d_20_cf90_ = d_19___mcc_h3_
            d_21_cf89_ = d_18___mcc_h2_
            d_22_cf88_ = d_17___mcc_h1_
            d_23_cf87_ = d_16___mcc_h0_
            return _dafny.CodePoint('h')
        elif True:
            d_24___mcc_h4_ = source1_.cf86
            d_25_cf86_ = d_24___mcc_h4_
            return _dafny.CodePoint('j')

    @staticmethod
    def fm22(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(40, -988):
                d_26_v0_: int = compr_4_
                if ((40) <= (d_26_v0_)) and ((d_26_v0_) < (-988)):
                    coll4_[(d_26_v0_) * (144)] = True
            return _dafny.Map(coll4_)
        return ((D6_DC9(iife4_()
)).cf12) | ((_dafny.Map({58: not(True)})) | (_dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})): False})))

    @staticmethod
    def fm23(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([not(False), True]), _dafny.SeqWithoutIsStrInference([not(not(True)), False, True])])

    @staticmethod
    def fm24(p0, globalState):
        if False:
            return _dafny.MultiSet([False, False])
        elif True:
            return _dafny.MultiSet([True])

    @staticmethod
    def fm26(globalState):
        return _dafny.Set({373})

    @staticmethod
    def fm27(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(951, 703):
                d_27_v0_: int = compr_5_
                if ((951) <= (d_27_v0_)) and ((d_27_v0_) < (703)):
                    coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_27_v0_, 631)]))
            return _dafny.Set(coll5_)
        return _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len(_dafny.Set({D13_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "livdrujf")), len(_dafny.Map({False: D6_DC9(_dafny.Map({(0) - (len(iife5_()
)): False}))})), True)})), (0) - (-719)) for d_28_i0_ in range(default__.abs(12))])

    @staticmethod
    def fm28(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: str
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('b')])).Elements:
                d_29_v0_: str = compr_6_
                if (d_29_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('b')])):
                    coll6_ = coll6_.union(_dafny.Set([d_29_v0_]))
            return _dafny.Set(coll6_)
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: str
            for compr_7_ in (_dafny.Map({_dafny.CodePoint('n'): False})).keys.Elements:
                d_30_v1_: str = compr_7_
                if (d_30_v1_) in (_dafny.Map({_dafny.CodePoint('n'): False})):
                    coll7_ = coll7_.union(_dafny.Set([d_30_v1_]))
            return _dafny.Set(coll7_)
        return _dafny.Set({(_dafny.Set({_dafny.CodePoint('e'), _dafny.CodePoint('o'), _dafny.CodePoint('f'), _dafny.CodePoint('q')}) if False else _dafny.Set({_dafny.CodePoint('b'), _dafny.CodePoint('p')})), (iife6_()
        ) - (_dafny.Set({_dafny.CodePoint('w'), _dafny.CodePoint('f'), _dafny.CodePoint('a'), _dafny.CodePoint('t'), _dafny.CodePoint('o')})), (_dafny.Set({_dafny.CodePoint('a'), _dafny.CodePoint('w'), _dafny.CodePoint('f'), _dafny.CodePoint('t'), _dafny.CodePoint('k')})) - (_dafny.Set({_dafny.CodePoint('e'), _dafny.CodePoint('x'), _dafny.CodePoint('n')})), iife7_()
        })

    @staticmethod
    def fm29(globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (_dafny.SeqWithoutIsStrInference([-482 for d_31_i0_ in range(default__.abs(96))])).Elements:
                d_32_v0_: int = compr_8_
                if (d_32_v0_) in (_dafny.SeqWithoutIsStrInference([-482 for d_31_i0_ in range(default__.abs(96))])):
                    coll8_[default__.safeDivisionInt(d_32_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuy"))))] = True
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: _dafny.Set
            for compr_9_ in (_dafny.MultiSet([_dafny.Set({473, 328}), _dafny.Set({464}), _dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))), len(_dafny.SeqWithoutIsStrInference([True, False, True, False, True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tloyglr")))}))})])).Elements:
                d_33_v1_: _dafny.Set = compr_9_
                if (d_33_v1_) in (_dafny.MultiSet([_dafny.Set({473, 328}), _dafny.Set({464}), _dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))), len(_dafny.SeqWithoutIsStrInference([True, False, True, False, True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tloyglr")))}))})])):
                    coll9_ = coll9_.union(_dafny.Set([d_33_v1_]))
            return _dafny.Set(coll9_)
        return ((_dafny.Set({D9_DC18(_dafny.Map({not(False): len(_dafny.Map({len(_dafny.Map({len(iife8_()
): False})): True}))}))})) | (_dafny.Set({D9_DC18(_dafny.Map({True: -716}))}))) - ((_dafny.Set({D9_DC18(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hucnnesk")))})), D9_DC18(_dafny.Map({False: 759})), D9_DC18(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))})), D9_DC18(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(iife9_()
), len(_dafny.Map({True: True})), 772]))})), D9_DC18(_dafny.Map({True: 832}))})) - (_dafny.Set({D9_DC18(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))}))})))

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({-526: _dafny.Map({not(False): True})})) | (_dafny.Map({144: _dafny.Map({False: True})}))) | (_dafny.Map({354: _dafny.Map({False: True})}))

    @staticmethod
    def fm31(p0, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference([368]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([741 for d_34_i1_ in range(default__.abs(-107))])) for d_35_i0_ in range(default__.abs(-154))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbacfmeic")): 500})), (0) - (-351), 576, 622]), _dafny.SeqWithoutIsStrInference([901 for d_36_i2_ in range(default__.abs(961))])})

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([True, False]))))]))).Elements:
                d_37_v0_: int = compr_10_
                if (d_37_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([True, False]))))]))):
                    coll10_[(d_37_v0_) * (len(_dafny.Map({True: True})))] = (D9_DC18(_dafny.Map({True: 841}))).cf26
            return _dafny.Map(coll10_)
        return ((_dafny.Map({len(_dafny.Set({True})): _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([313, 538]))]))})})) | (iife10_()
        )) | (_dafny.Map({703: _dafny.Map({True: len(_dafny.Map({True: 669}))})}))

    @staticmethod
    def fm33(p0, globalState):
        if not(False):
            return D6_DC10((D11_DC24(413, _dafny.SeqWithoutIsStrInference([True]), _dafny.CodePoint('f'))).cf35)
        elif True:
            return D6_DC10(_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm34(p0, p1, globalState):
        source2_ = D22_DC52()
        if source2_.is_DC52:
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "re"))), len(_dafny.Map({_dafny.CodePoint('q'): True}))])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dppgedr")))])).cardinality for d_38_i0_ in range(default__.abs(674))]))
        elif source2_.is_DC53:
            d_39___mcc_h0_ = source2_.cf87
            d_40___mcc_h1_ = source2_.cf88
            d_41___mcc_h2_ = source2_.cf89
            d_42___mcc_h3_ = source2_.cf90
            d_43_cf90_ = d_42___mcc_h3_
            d_44_cf89_ = d_41___mcc_h2_
            d_45_cf88_ = d_40___mcc_h1_
            d_46_cf87_ = d_39___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([d_46_cf87_])) + (_dafny.SeqWithoutIsStrInference([921 for d_47_i1_ in range(default__.abs(140))]))
        elif True:
            d_48___mcc_h4_ = source2_.cf86
            d_49_cf86_ = d_48___mcc_h4_
            return _dafny.SeqWithoutIsStrInference([len(_dafny.Map({-538: True})) for d_50_i2_ in range(default__.abs(-42))])

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        source3_ = D6_DC9(_dafny.Map({19: not(True)}))
        if source3_.is_DC10:
            d_51___mcc_h0_ = source3_.cf13
            d_52_cf13_ = d_51___mcc_h0_
            return D7_DC13(_dafny.Set({False}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcxj"))), len(_dafny.Map({len(_dafny.Set({628})): -166})))
        elif source3_.is_DC11:
            d_53___mcc_h1_ = source3_.cf14
            d_54___mcc_h2_ = source3_.cf15
            d_55___mcc_h3_ = source3_.cf16
            d_56_cf16_ = d_55___mcc_h3_
            d_57_cf15_ = d_54___mcc_h2_
            d_58_cf14_ = d_53___mcc_h1_
            return D7_DC13(_dafny.Set({True, True}), d_57_cf15_, d_56_cf16_)
        elif True:
            d_59___mcc_h4_ = source3_.cf12
            d_60_cf12_ = d_59___mcc_h4_
            return D7_DC13(_dafny.Set({False}), 461, -56)

    @staticmethod
    def fm37(globalState):
        return _dafny.MultiSet([(len(_dafny.Set({863}))) == (len(_dafny.Set({False, True}))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dexnes"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sijt"))), not((369) != (len(_dafny.Set({False})))), False])

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return D0_DC0(default__.safeDivisionInt((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_61_i0_ in range(default__.abs(278))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrpag")))])).cardinality, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), 566, len(_dafny.SeqWithoutIsStrInference([True, True, True, True]))]))))

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        return _dafny.CodePoint('a')

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        return _dafny.Set({D11_DC23(False, 575, 896), D11_DC23(True, len(_dafny.Map({False: 169})), 427)})

    @staticmethod
    def fm41(globalState):
        source4_ = D15_DC36(4, False, True, True, 388)
        if source4_.is_DC34:
            return _dafny.Set({False, False, not((D20_DC45(True, True)).cf76), not(not(False)), True})
        elif source4_.is_DC35:
            d_62___mcc_h0_ = source4_.cf50
            d_63___mcc_h1_ = source4_.cf51
            d_64___mcc_h2_ = source4_.cf52
            d_65___mcc_h3_ = source4_.cf53
            d_66___mcc_h4_ = source4_.cf54
            d_67_cf54_ = d_66___mcc_h4_
            d_68_cf53_ = d_65___mcc_h3_
            d_69_cf52_ = d_64___mcc_h2_
            d_70_cf51_ = d_63___mcc_h1_
            d_71_cf50_ = d_62___mcc_h0_
            return ((d_70_cf51_).f12) | ((d_70_cf51_).f12)
        elif source4_.is_DC36:
            d_72___mcc_h5_ = source4_.cf55
            d_73___mcc_h6_ = source4_.cf56
            d_74___mcc_h7_ = source4_.cf57
            d_75___mcc_h8_ = source4_.cf58
            d_76___mcc_h9_ = source4_.cf59
            d_77_cf59_ = d_76___mcc_h9_
            d_78_cf58_ = d_75___mcc_h8_
            d_79_cf57_ = d_74___mcc_h7_
            d_80_cf56_ = d_73___mcc_h6_
            d_81_cf55_ = d_72___mcc_h5_
            return (_dafny.Set({d_78_cf58_, d_79_cf57_})) - (_dafny.Set({d_78_cf58_}))
        elif source4_.is_DC33:
            d_82___mcc_h10_ = source4_.cf49
            d_83_cf49_ = d_82___mcc_h10_
            return _dafny.Set({True})
        elif True:
            d_84___mcc_h11_ = source4_.cf60
            d_85_cf60_ = d_84___mcc_h11_
            return _dafny.Set({not(not(True)), not(True)})

    @staticmethod
    def fm42(p0, globalState):
        return D9_DC19()

    @staticmethod
    def fm43(p0, globalState):
        def iife11_():
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(633, -258):
                    d_88_v0_: int = compr_13_
                    if ((633) <= (d_88_v0_)) and ((d_88_v0_) < (-258)):
                        coll13_[default__.safeModuloInt(d_88_v0_, len(_dafny.SeqWithoutIsStrInference([True])))] = len(_dafny.Map({True: True}))
                return _dafny.Map(coll13_)
            coll11_ = _dafny.Set()
            def iife12_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(633, -258):
                    d_88_v0_: int = compr_12_
                    if ((633) <= (d_88_v0_)) and ((d_88_v0_) < (-258)):
                        coll12_[default__.safeModuloInt(d_88_v0_, len(_dafny.SeqWithoutIsStrInference([True])))] = len(_dafny.Map({True: True}))
                return _dafny.Map(coll12_)
            compr_11_: int
            for compr_11_ in (_dafny.SeqWithoutIsStrInference([347, len(iife12_()
            )])).Elements:
                d_89_v1_: int = compr_11_
                if (d_89_v1_) in (_dafny.SeqWithoutIsStrInference([347, len(iife13_()
                )])):
                    def iife14_():
                        coll14_ = _dafny.Map()
                        compr_14_: _dafny.Seq
                        for compr_14_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wra")): 574})).keys.Elements:
                            d_91_v2_: _dafny.Seq = compr_14_
                            if (d_91_v2_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wra")): 574})):
                                coll14_[d_91_v2_] = False
                        return _dafny.Map(coll14_)
                    coll11_ = coll11_.union(_dafny.Set([(d_89_v1_) - ((_dafny.MultiSet([634, len(_dafny.Map({_dafny.CodePoint('y'): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_90_i2_ in range(default__.abs(993))]))})), len(iife14_()
)])).cardinality)]))
            return _dafny.Set(coll11_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bha")): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([703]))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpxjw")): _dafny.MultiSet([-382])}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_86_i0_ in range(default__.abs(529))]): _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_87_i1_ in range(default__.abs(697))])), len(iife11_()
        )])}))

    @staticmethod
    def fm44(p0, p1, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: D8
            for compr_15_ in (_dafny.MultiSet([D8_DC17(D8_DC16(False, len(_dafny.Map({True: _dafny.CodePoint('p')}))))])).Elements:
                d_92_v0_: D8 = compr_15_
                if (d_92_v0_) in (_dafny.MultiSet([D8_DC17(D8_DC16(False, len(_dafny.Map({True: _dafny.CodePoint('p')}))))])):
                    coll15_[d_92_v0_] = (0) - (len(_dafny.Set({not((D21_DC48(False)).cf78), not(False)})))
            return _dafny.Map(coll15_)
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: D8
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([D8_DC17(D8_DC15(_dafny.Map({False: False}))) for d_93_i0_ in range(default__.abs(727))])).Elements:
                d_94_v1_: D8 = compr_16_
                if (d_94_v1_) in (_dafny.SeqWithoutIsStrInference([D8_DC17(D8_DC15(_dafny.Map({False: False}))) for d_93_i0_ in range(default__.abs(727))])):
                    coll16_[d_94_v1_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.Set({True}))})]))).cardinality
            return _dafny.Map(coll16_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: D8
            for compr_17_ in (_dafny.SeqWithoutIsStrInference([D8_DC17(D8_DC17(D8_DC16(False, len(_dafny.SeqWithoutIsStrInference([456])))))])).Elements:
                d_95_v2_: D8 = compr_17_
                if (d_95_v2_) in (_dafny.SeqWithoutIsStrInference([D8_DC17(D8_DC17(D8_DC16(False, len(_dafny.SeqWithoutIsStrInference([456])))))])):
                    coll17_[d_95_v2_] = 988
            return _dafny.Map(coll17_)
        return _dafny.SeqWithoutIsStrInference([(iife15_()
        ) | (iife16_()
        ), iife17_()
        , (_dafny.Map({D8_DC17(D8_DC15(_dafny.Map({False: True}))): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(True), True, True, False})) for d_96_i1_ in range(default__.abs(819))]))})) | (_dafny.Map({D8_DC17(D8_DC15(_dafny.Map({True: not(True)}))): -665}))])

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([-462])) | ((_dafny.MultiSet([-710])) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxu")))])))

    @staticmethod
    def fm46(globalState):
        def iife18_():
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(945, 197):
                    d_97_v1_: int = compr_20_
                    if ((945) <= (d_97_v1_)) and ((d_97_v1_) < (197)):
                        coll20_[(d_97_v1_) - (len(_dafny.Map({True: False})))] = len(_dafny.Map({False: False}))
                return _dafny.Map(coll20_)
            coll18_ = _dafny.Map()
            def iife19_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in _dafny.IntegerRange(945, 197):
                    d_97_v1_: int = compr_19_
                    if ((945) <= (d_97_v1_)) and ((d_97_v1_) < (197)):
                        coll19_[(d_97_v1_) - (len(_dafny.Map({True: False})))] = len(_dafny.Map({False: False}))
                return _dafny.Map(coll19_)
            compr_18_: int
            for compr_18_ in (iife19_()
            ).keys.Elements:
                d_98_v0_: int = compr_18_
                if (d_98_v0_) in (iife20_()
                ):
                    coll18_[(d_98_v0_) - (-664)] = (_dafny.MultiSet([False])).cardinality
            return _dafny.Map(coll18_)
        return (((D16_DC38(iife18_()
)).cf61) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.Map({173: 261}))}))) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([666])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvgkqkf")))})) | (_dafny.Map({len(_dafny.Map({False: not(True)})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))})))

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: _dafny.Map
            for compr_21_ in (_dafny.Map({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([not(True)]))}): len(_dafny.Map({False: True}))})).keys.Elements:
                d_99_v0_: _dafny.Map = compr_21_
                if (d_99_v0_) in (_dafny.Map({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([not(True)]))}): len(_dafny.Map({False: True}))})):
                    coll21_[d_99_v0_] = len(_dafny.Set({True}))
            return _dafny.Map(coll21_)
        return D13_DC27(_dafny.Map({D8_DC17(D8_DC16(False, len(iife21_()
))): 777}))

    @staticmethod
    def fm48(p0, globalState):
        if False:
            return D8_DC15(_dafny.Map({True: not(True)}))
        elif True:
            return D8_DC15(_dafny.Map({not(False): False}))

    @staticmethod
    def fm49(p0, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_100_i0_ in range(default__.abs(564))]))) > (len(_dafny.Set({False, False, False, not(True), True}))):
            return D0_DC1(True, 817, True)
        elif True:
            return D0_DC1(not(True), len(_dafny.SeqWithoutIsStrInference([True])), False)

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        return D11_DC23((True if False else not(False)), (-90 if True else 147), len((_dafny.Set({True})) - (_dafny.Set({False}))))

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: str
            for compr_22_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('w'), _dafny.CodePoint('f')])).Elements:
                d_101_v0_: str = compr_22_
                if (d_101_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('w'), _dafny.CodePoint('f')])):
                    coll22_ = coll22_.union(_dafny.Set([d_101_v0_]))
            return _dafny.Set(coll22_)
        return ((_dafny.SeqWithoutIsStrInference([D14_DC32(D14_DC31(True, 177, 233))])) + (_dafny.SeqWithoutIsStrInference([D14_DC32(D14_DC32(D14_DC32(D14_DC31(False, 71, 802)))), D14_DC32(D14_DC32(D14_DC32(D14_DC31(False, -926, 123)))), D14_DC32(D14_DC32(D14_DC31(True, len(_dafny.Map({True: False})), (0) - (len(iife22_()
)))))]))) + (_dafny.SeqWithoutIsStrInference([D14_DC32(D14_DC30(_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "immtempc"))}))), D14_DC32(D14_DC31(True, 903, (_dafny.MultiSet([len(_dafny.Map({False: -189}))])).cardinality))]))

    @staticmethod
    def fm52(p0, p1, globalState):
        def iife23_():
            def iife25_():
                coll25_ = _dafny.Map()
                compr_25_: D0
                for compr_25_ in (_dafny.SeqWithoutIsStrInference([D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkuwoejgm"))), False), D0_DC1(False, 45, not(False)), D0_DC1(not(False), len(_dafny.SeqWithoutIsStrInference([568])), not(not(False)))])).Elements:
                    d_102_v1_: D0 = compr_25_
                    if (d_102_v1_) in (_dafny.SeqWithoutIsStrInference([D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkuwoejgm"))), False), D0_DC1(False, 45, not(False)), D0_DC1(not(False), len(_dafny.SeqWithoutIsStrInference([568])), not(not(False)))])):
                        coll25_[d_102_v1_] = _dafny.Set({(_dafny.MultiSet([True])).cardinality})
                return _dafny.Map(coll25_)
            coll23_ = _dafny.Map()
            def iife24_():
                coll24_ = _dafny.Map()
                compr_24_: D0
                for compr_24_ in (_dafny.SeqWithoutIsStrInference([D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkuwoejgm"))), False), D0_DC1(False, 45, not(False)), D0_DC1(not(False), len(_dafny.SeqWithoutIsStrInference([568])), not(not(False)))])).Elements:
                    d_102_v1_: D0 = compr_24_
                    if (d_102_v1_) in (_dafny.SeqWithoutIsStrInference([D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkuwoejgm"))), False), D0_DC1(False, 45, not(False)), D0_DC1(not(False), len(_dafny.SeqWithoutIsStrInference([568])), not(not(False)))])):
                        coll24_[d_102_v1_] = _dafny.Set({(_dafny.MultiSet([True])).cardinality})
                return _dafny.Map(coll24_)
            compr_23_: D0
            for compr_23_ in (iife24_()
            ).keys.Elements:
                d_103_v0_: D0 = compr_23_
                if (d_103_v0_) in (iife25_()
                ):
                    coll23_[d_103_v0_] = True
            return _dafny.Map(coll23_)
        return (_dafny.Map({D0_DC1(not(False), 949, False): False})) | ((iife23_()
        ) | (_dafny.Map({D0_DC1(True, (0) - (-656), True): False})))

    @staticmethod
    def fm53(p0, p1, p2, globalState):
        return _dafny.MultiSet([_dafny.Set({False}), _dafny.Set({not(not(True)), not(not(False))})])

    @staticmethod
    def fm54(p0, globalState):
        def iife26_():
            coll26_ = _dafny.Map()
            compr_26_: _dafny.Map
            for compr_26_ in (_dafny.Map({_dafny.Map({not(not(True)): _dafny.CodePoint('h')}): D6_DC9(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_104_i0_ in range(default__.abs(768))])): True}))})).keys.Elements:
                d_105_v0_: _dafny.Map = compr_26_
                if (d_105_v0_) in (_dafny.Map({_dafny.Map({not(not(True)): _dafny.CodePoint('h')}): D6_DC9(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_104_i0_ in range(default__.abs(768))])): True}))})):
                    coll26_[d_105_v0_] = False
            return _dafny.Map(coll26_)
        return _dafny.Map({_dafny.CodePoint('l'): len(iife26_()
        )})

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        def iife27_():
            coll27_ = _dafny.Map()
            compr_27_: _dafny.Set
            for compr_27_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Set({False})), -889}) for d_106_i0_ in range(default__.abs(464))])).Elements:
                d_107_v0_: _dafny.Set = compr_27_
                if (d_107_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Set({False})), -889}) for d_106_i0_ in range(default__.abs(464))])):
                    coll27_[d_107_v0_] = _dafny.Map({True: False})
            return _dafny.Map(coll27_)
        return (_dafny.Map({_dafny.Set({len(_dafny.Map({len(_dafny.Set({not(True), False})): 294}))}): _dafny.Map({False: False})})) | (iife27_()
        )

    @staticmethod
    def fm56(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([125, (_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('p')])).cardinality, len((_dafny.Map({True: _dafny.CodePoint('s')}))), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: True}), _dafny.Map({True: (D15_DC36(len(_dafny.Set({878})), True, True, True, 714)).cf57})]))), 750]) for d_108_i0_ in range(default__.abs(501))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-861]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgbtjar"))), 24]), _dafny.SeqWithoutIsStrInference([63 for d_109_i1_ in range(default__.abs(97))])]))

    @staticmethod
    def Main(noArgsParameter__):
        d_110_v0_: bool
        d_110_v0_ = False
        d_111_v1_: int
        d_111_v1_ = 754
        d_112_v2_: _dafny.Seq
        d_112_v2_ = _dafny.SeqWithoutIsStrInference([True, d_110_v0_, d_110_v0_])
        d_113_v3_: _dafny.Map
        d_113_v3_ = _dafny.Map({d_111_v1_: len(d_112_v2_)})
        d_114_v4_: _dafny.Map
        d_114_v4_ = _dafny.Map({d_110_v0_: ((d_113_v3_)[d_111_v1_] if (d_111_v1_) in (d_113_v3_) else d_111_v1_)})
        d_115_v5_: _dafny.Array
        nw0_ = _dafny.Array(None, 8)
        nw0_[int(0)] = d_114_v4_
        nw0_[int(1)] = d_114_v4_
        nw0_[int(2)] = d_114_v4_
        nw0_[int(3)] = _dafny.Map({d_110_v0_: d_111_v1_})
        nw0_[int(4)] = d_114_v4_
        nw0_[int(5)] = _dafny.Map({d_110_v0_: d_111_v1_})
        nw0_[int(6)] = d_114_v4_
        nw0_[int(7)] = d_114_v4_
        d_115_v5_ = nw0_
        d_116_globalState_: GlobalState
        nw1_ = GlobalState()
        nw1_.ctor__(False, 92, 431, -645, False, d_115_v5_, True, True, False, -357, True)
        d_116_globalState_ = nw1_
        d_117_i0_: int
        d_117_i0_ = 0
        with _dafny.label("0"):
            while d_110_v0_:
                with _dafny.c_label("0"):
                    if (d_117_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_117_i0_ = (d_117_i0_) + (1)
                    (d_116_globalState_).f6 = d_110_v0_
                    if not(d_110_v0_):
                        d_118_v6_: D0
                        d_118_v6_ = D0_DC1(d_110_v0_, d_111_v1_, d_110_v0_)
                        d_119_v7_: _dafny.Seq
                        d_119_v7_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_110_v0_, d_110_v0_}))])
                        d_120_v8_: _dafny.MultiSet
                        d_120_v8_ = _dafny.MultiSet([d_111_v1_, d_111_v1_, d_111_v1_, (0) - (d_111_v1_), d_111_v1_])
                        d_121_v9_: _dafny.Set
                        d_121_v9_ = _dafny.Set({((d_120_v8_)[d_111_v1_] if (d_111_v1_) in (d_120_v8_) else 664)})
                        d_122_v10_: _dafny.Seq
                        d_122_v10_ = _dafny.SeqWithoutIsStrInference([d_111_v1_, (d_118_v6_).cf2, d_111_v1_, len(d_119_v7_), len(d_121_v9_)])
                        rhs0_ = not((_dafny.SeqWithoutIsStrInference([d_111_v1_, d_111_v1_])) == (d_119_v7_))
                        rhs1_ = d_110_v0_
                        rhs2_ = not (False) or (d_110_v0_)
                        lhs0_ = d_116_globalState_
                        lhs1_ = d_116_globalState_
                        lhs0_.f7 = rhs0_
                        lhs1_.f7 = rhs1_
                        d_110_v0_ = rhs2_
                        d_123_v11_: _dafny.Array
                        nw2_ = _dafny.Array(False, 26)
                        d_123_v11_ = nw2_
                        d_123_v11_ = d_123_v11_
                        d_124_v12_: _dafny.Seq
                        d_124_v12_ = _dafny.SeqWithoutIsStrInference([d_123_v11_])
                        d_125_v13_: _dafny.Set
                        d_125_v13_ = _dafny.Set({default__.fm0(d_116_globalState_)})
                        d_126_v14_: _dafny.Map
                        d_126_v14_ = _dafny.Map({d_111_v1_: d_125_v13_})
                        d_127_v15_: _dafny.Array
                        nw3_ = _dafny.Array(None, 17)
                        nw3_[int(0)] = (0) - ((len(d_119_v7_)) * (len(d_113_v3_)))
                        nw3_[int(1)] = d_111_v1_
                        nw3_[int(2)] = d_111_v1_
                        nw3_[int(3)] = d_111_v1_
                        nw3_[int(4)] = default__.safeDivisionInt(d_111_v1_, d_111_v1_)
                        nw3_[int(5)] = len(d_124_v12_)
                        nw3_[int(6)] = (0) - (len(d_126_v14_))
                        nw3_[int(7)] = d_111_v1_
                        nw3_[int(8)] = d_111_v1_
                        nw3_[int(9)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_128_i1_ in range(default__.abs(753))]))
                        nw3_[int(10)] = d_111_v1_
                        nw3_[int(11)] = len((d_112_v2_) + (d_112_v2_))
                        nw3_[int(12)] = ((0) - (d_111_v1_)) * (len(d_113_v3_))
                        nw3_[int(13)] = d_111_v1_
                        nw3_[int(14)] = d_111_v1_
                        nw3_[int(15)] = (d_111_v1_ if False else d_111_v1_)
                        nw3_[int(16)] = d_111_v1_
                        d_127_v15_ = nw3_
                        index0_ = default__.safeIndex(552, (d_127_v15_).length(0))
                        (d_127_v15_)[index0_] = d_111_v1_
                        d_129_v16_: D0
                        d_129_v16_ = D0_DC0(d_111_v1_)
                        pat_let_tv0_ = d_111_v1_
                        d_130_v17_: _dafny.Array
                        nw4_ = _dafny.Array(None, 12)
                        nw4_[int(0)] = D0_DC0(d_111_v1_)
                        nw4_[int(1)] = d_129_v16_
                        nw4_[int(2)] = d_129_v16_
                        nw4_[int(3)] = d_129_v16_
                        nw4_[int(4)] = d_129_v16_
                        nw4_[int(5)] = d_129_v16_
                        nw4_[int(6)] = D0_DC0(d_111_v1_)
                        nw4_[int(7)] = d_129_v16_
                        nw4_[int(8)] = d_129_v16_
                        nw4_[int(9)] = d_129_v16_
                        nw4_[int(10)] = d_129_v16_
                        def iife28_(_pat_let0_0):
                            def iife29_(d_131_dt__update__tmp_h0_):
                                def iife30_(_pat_let1_0):
                                    def iife31_(d_132_dt__update_hcf0_h0_):
                                        return D0_DC0(d_132_dt__update_hcf0_h0_)
                                    return iife31_(_pat_let1_0)
                                return iife30_(pat_let_tv0_)
                            return iife29_(_pat_let0_0)
                        nw4_[int(11)] = iife28_(d_129_v16_)
                        d_130_v17_ = nw4_
                        d_133_v18_: _dafny.MultiSet
                        d_133_v18_ = _dafny.MultiSet([d_130_v17_, d_130_v17_])
                        index1_ = default__.safeIndex(552, (d_127_v15_).length(0))
                        rhs3_ = d_110_v0_
                        rhs4_ = (d_133_v18_).ispropersubset(((d_133_v18_).set(d_130_v17_, default__.abs(d_111_v1_))) - (d_133_v18_))
                        rhs5_ = (0) - ((d_111_v1_) - (d_111_v1_))
                        lhs2_ = d_116_globalState_
                        lhs3_ = d_116_globalState_
                        lhs4_ = d_127_v15_
                        lhs5_ = default__.safeIndex(552, (d_127_v15_).length(0))
                        lhs2_.f6 = rhs3_
                        lhs3_.f6 = rhs4_
                        lhs4_[lhs5_] = rhs5_
                        rhs6_ = d_127_v15_
                        rhs7_ = (len(d_119_v7_)) > (d_111_v1_)
                        rhs8_ = ((d_127_v15_)[default__.safeIndex(552, (d_127_v15_).length(0))]) * (d_111_v1_)
                        lhs6_ = d_116_globalState_
                        d_127_v15_ = rhs6_
                        lhs6_.f6 = rhs7_
                        d_111_v1_ = rhs8_
                        index2_ = default__.safeIndex(552, (d_127_v15_).length(0))
                        (d_127_v15_)[index2_] = d_111_v1_
                    elif True:
                        d_134_v19_: _dafny.Seq
                        d_134_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prcfca"))
                        d_135_v20_: _dafny.Map
                        d_135_v20_ = _dafny.Map({d_110_v0_: d_114_v4_})
                        d_136_v21_: _dafny.Map
                        d_136_v21_ = _dafny.Map({(default__.fm1(d_116_globalState_)) + (d_134_v19_): default__.fm2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xngayfcy"))), len(d_135_v20_), d_116_globalState_)})
                        d_137_v22_: _dafny.Seq
                        d_137_v22_ = _dafny.SeqWithoutIsStrInference([d_111_v1_, d_111_v1_])
                        d_136_v21_ = (d_136_v21_).set((d_134_v19_ if d_110_v0_ else d_134_v19_), (d_137_v22_)[default__.safeIndex((0) - (d_111_v1_), len(d_137_v22_))])
                        d_138_v23_: _dafny.MultiSet
                        d_138_v23_ = _dafny.MultiSet([default__.fm2(d_111_v1_, d_111_v1_, d_116_globalState_)])
                        d_139_v24_: _dafny.Map
                        d_139_v24_ = _dafny.Map({d_138_v23_: d_110_v0_})
                        d_139_v24_ = (d_139_v24_).set(_dafny.MultiSet(d_137_v22_), d_110_v0_)
                        d_140_v25_: _dafny.Map
                        d_140_v25_ = _dafny.Map({False: d_110_v0_})
                        d_141_v26_: _dafny.Set
                        d_141_v26_ = _dafny.Set({_dafny.Map({d_110_v0_: d_110_v0_})})
                        d_142_v27_: _dafny.Map
                        d_142_v27_ = _dafny.Map({d_111_v1_: not(d_110_v0_)})
                        d_143_v28_: D0
                        d_143_v28_ = D0_DC1(not(d_110_v0_), len(_dafny.Map({d_111_v1_: d_110_v0_})), d_110_v0_)
                        d_144_v29_: _dafny.Array
                        nw5_ = _dafny.Array(None, 24)
                        nw5_[int(0)] = d_110_v0_
                        nw5_[int(1)] = (_dafny.Set({d_140_v25_, _dafny.Map({d_110_v0_: d_110_v0_})})).issubset(d_141_v26_)
                        nw5_[int(2)] = d_110_v0_
                        nw5_[int(3)] = d_110_v0_
                        nw5_[int(4)] = not(d_110_v0_)
                        nw5_[int(5)] = d_110_v0_
                        nw5_[int(6)] = (d_137_v22_) <= (d_137_v22_)
                        nw5_[int(7)] = d_110_v0_
                        nw5_[int(8)] = not(not(True))
                        nw5_[int(9)] = (d_138_v23_).issubset(d_138_v23_)
                        nw5_[int(10)] = (True) or (d_110_v0_)
                        nw5_[int(11)] = d_110_v0_
                        nw5_[int(12)] = d_110_v0_
                        nw5_[int(13)] = (d_111_v1_) < (135)
                        nw5_[int(14)] = True
                        nw5_[int(15)] = d_110_v0_
                        nw5_[int(16)] = default__.fm0(d_116_globalState_)
                        nw5_[int(17)] = d_110_v0_
                        nw5_[int(18)] = (930) == (d_111_v1_)
                        nw5_[int(19)] = d_110_v0_
                        nw5_[int(20)] = ((d_142_v27_)[d_111_v1_] if (d_111_v1_) in (d_142_v27_) else d_110_v0_)
                        nw5_[int(21)] = (d_143_v28_).cf1
                        nw5_[int(22)] = d_110_v0_
                        nw5_[int(23)] = (d_142_v27_) != (_dafny.Map({len(_dafny.Map({d_111_v1_: (d_112_v2_)[default__.safeIndex(d_111_v1_, len(d_112_v2_))]})): not(default__.fm0(d_116_globalState_))}))
                        d_144_v29_ = nw5_
                        index3_ = default__.safeIndex(152, (d_144_v29_).length(0))
                        (d_144_v29_)[index3_] = d_110_v0_
                        index4_ = default__.safeIndex(152, (d_144_v29_).length(0))
                        (d_144_v29_)[index4_] = d_110_v0_
                        d_145_v30_: str
                        d_145_v30_ = _dafny.CodePoint('v')
                        d_146_v31_: C2
                        nw6_ = C2()
                        nw6_.ctor__(d_145_v30_, d_111_v1_)
                        d_146_v31_ = nw6_
                        d_147_v32_: bool
                        d_148_v33_: bool
                        d_149_v34_: bool
                        d_150_v35_: int
                        out0_: bool
                        out1_: bool
                        out2_: bool
                        out3_: int
                        out0_, out1_, out2_, out3_ = (d_146_v31_).m8(d_116_globalState_)
                        d_147_v32_ = out0_
                        d_148_v33_ = out1_
                        d_149_v34_ = out2_
                        d_150_v35_ = out3_
                    if d_110_v0_:
                        d_151_v36_: _dafny.Set
                        d_151_v36_ = _dafny.Set({d_110_v0_})
                        d_152_v37_: _dafny.Set
                        d_152_v37_ = _dafny.Set({d_110_v0_, d_110_v0_, d_110_v0_, (d_112_v2_)[default__.safeIndex(len(d_151_v36_), len(d_112_v2_))], d_110_v0_})
                        d_153_v38_: _dafny.Seq
                        d_153_v38_ = _dafny.SeqWithoutIsStrInference([d_151_v36_])
                        d_154_v39_: _dafny.Seq
                        d_154_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eua"))
                        d_155_v40_: D7
                        d_155_v40_ = D7_DC13((d_153_v38_)[default__.safeIndex(len(d_154_v39_), len(d_153_v38_))], d_111_v1_, d_111_v1_)
                        pat_let_tv1_ = d_111_v1_
                        pat_let_tv2_ = d_111_v1_
                        pat_let_tv3_ = d_116_globalState_
                        def iife32_(_pat_let2_0):
                            def iife33_(d_156_dt__update__tmp_h1_):
                                def iife34_(_pat_let3_0):
                                    def iife35_(d_157_dt__update_hcf19_h0_):
                                        return D7_DC13((d_156_dt__update__tmp_h1_).cf18, d_157_dt__update_hcf19_h0_, (d_156_dt__update__tmp_h1_).cf20)
                                    return iife35_(_pat_let3_0)
                                return iife34_(default__.fm2(pat_let_tv1_, pat_let_tv2_, pat_let_tv3_))
                            return iife33_(_pat_let2_0)
                        d_111_v1_ = (d_111_v1_) * ((iife32_(d_155_v40_)).cf19)
                        d_154_v39_ = default__.fm1(d_116_globalState_)
                        d_158_v41_: str
                        d_158_v41_ = _dafny.CodePoint('n')
                        d_159_v42_: _dafny.Array
                        def lambda0_(d_160_v41_):
                            def lambda1_(d_161_i2_):
                                return d_160_v41_

                            return lambda1_

                        init0_ = lambda0_(d_158_v41_)
                        nw7_ = _dafny.Array(None, 8)
                        for i0_0_ in range(nw7_.length(0)):
                            nw7_[i0_0_] = init0_(i0_0_)
                        d_159_v42_ = nw7_
                        d_162_v43_: D21
                        d_162_v43_ = D21_DC50(d_159_v42_, (0) - (d_111_v1_), d_110_v0_)
                        d_163_v44_: _dafny.Seq
                        d_163_v44_ = _dafny.SeqWithoutIsStrInference([d_154_v39_])
                        d_164_v45_: _dafny.Seq
                        d_164_v45_ = _dafny.SeqWithoutIsStrInference([d_111_v1_, d_111_v1_, d_111_v1_, d_111_v1_, len(d_163_v44_)])
                        d_165_v46_: _dafny.Array
                        nw8_ = _dafny.Array(None, 24)
                        nw8_[int(0)] = True
                        nw8_[int(1)] = (d_158_v41_) not in (d_154_v39_)
                        nw8_[int(2)] = default__.fm0(d_116_globalState_)
                        nw8_[int(3)] = d_110_v0_
                        nw8_[int(4)] = default__.fm0(d_116_globalState_)
                        nw8_[int(5)] = d_110_v0_
                        nw8_[int(6)] = (d_162_v43_).cf85
                        nw8_[int(7)] = (d_112_v2_)[default__.safeIndex(d_111_v1_, len(d_112_v2_))]
                        nw8_[int(8)] = d_110_v0_
                        nw8_[int(9)] = not(d_110_v0_)
                        nw8_[int(10)] = (d_111_v1_) > ((d_164_v45_)[default__.safeIndex((_dafny.MultiSet([d_110_v0_])).cardinality, len(d_164_v45_))])
                        nw8_[int(11)] = d_110_v0_
                        nw8_[int(12)] = d_110_v0_
                        nw8_[int(13)] = default__.fm0(d_116_globalState_)
                        nw8_[int(14)] = d_110_v0_
                        nw8_[int(15)] = (d_111_v1_) == (305)
                        nw8_[int(16)] = not(d_110_v0_)
                        nw8_[int(17)] = d_110_v0_
                        nw8_[int(18)] = True
                        nw8_[int(19)] = (not(d_110_v0_)) and (d_110_v0_)
                        nw8_[int(20)] = True
                        nw8_[int(21)] = default__.fm0(d_116_globalState_)
                        nw8_[int(22)] = d_110_v0_
                        nw8_[int(23)] = d_110_v0_
                        d_165_v46_ = nw8_
                        d_165_v46_ = d_165_v46_
                        d_166_v47_: _dafny.MultiSet
                        d_166_v47_ = _dafny.MultiSet([(d_111_v1_) <= (d_111_v1_)])
                        d_167_v48_: _dafny.Map
                        d_167_v48_ = _dafny.Map({d_110_v0_: d_110_v0_})
                        d_168_v49_: _dafny.Map
                        d_168_v49_ = _dafny.Map({d_110_v0_: d_167_v48_})
                        d_169_v50_: D16
                        d_169_v50_ = D16_DC39(d_111_v1_, d_110_v0_, ((d_168_v49_)[d_110_v0_] if (d_110_v0_) in (d_168_v49_) else d_167_v48_), d_158_v41_)
                        d_170_v51_: D6
                        d_170_v51_ = D6_DC10(d_112_v2_)
                        d_171_v52_: D19
                        d_171_v52_ = D19_DC43(d_110_v0_, (d_169_v50_).cf65, d_110_v0_, d_170_v51_, d_154_v39_)
                        rhs9_ = (d_171_v52_).cf70
                        rhs10_ = d_110_v0_
                        rhs11_ = d_166_v47_
                        d_158_v41_ = rhs9_
                        d_110_v0_ = rhs10_
                        d_166_v47_ = rhs11_
                        d_172_v53_: _dafny.Array
                        def lambda2_(d_173_v1_):
                            def lambda3_(d_174_i3_):
                                return (d_174_i3_) * (d_173_v1_)

                            return lambda3_

                        init1_ = lambda2_(d_111_v1_)
                        nw9_ = _dafny.Array(None, 24)
                        for i0_1_ in range(nw9_.length(0)):
                            nw9_[i0_1_] = init1_(i0_1_)
                        d_172_v53_ = nw9_
                        d_175_v54_: _dafny.Array
                        nw10_ = _dafny.Array(None, 29)
                        nw10_[int(0)] = d_172_v53_
                        nw10_[int(1)] = d_172_v53_
                        nw10_[int(2)] = d_172_v53_
                        nw10_[int(3)] = d_172_v53_
                        nw10_[int(4)] = d_172_v53_
                        nw10_[int(5)] = d_172_v53_
                        nw10_[int(6)] = d_172_v53_
                        nw10_[int(7)] = d_172_v53_
                        nw10_[int(8)] = d_172_v53_
                        nw10_[int(9)] = d_172_v53_
                        nw10_[int(10)] = d_172_v53_
                        nw10_[int(11)] = d_172_v53_
                        nw10_[int(12)] = d_172_v53_
                        nw10_[int(13)] = d_172_v53_
                        nw10_[int(14)] = d_172_v53_
                        nw10_[int(15)] = d_172_v53_
                        nw10_[int(16)] = d_172_v53_
                        nw10_[int(17)] = d_172_v53_
                        nw10_[int(18)] = d_172_v53_
                        nw10_[int(19)] = d_172_v53_
                        nw10_[int(20)] = d_172_v53_
                        nw10_[int(21)] = d_172_v53_
                        nw10_[int(22)] = d_172_v53_
                        nw10_[int(23)] = d_172_v53_
                        nw10_[int(24)] = d_172_v53_
                        nw10_[int(25)] = d_172_v53_
                        nw10_[int(26)] = d_172_v53_
                        nw10_[int(27)] = d_172_v53_
                        nw10_[int(28)] = d_172_v53_
                        d_175_v54_ = nw10_
                        d_176_v55_: _dafny.Array
                        d_176_v55_ = d_175_v54_
                        d_177_v56_: _dafny.Array
                        nw11_ = _dafny.Array(None, 23)
                        nw11_[int(0)] = d_175_v54_
                        nw11_[int(1)] = d_176_v55_
                        nw11_[int(2)] = d_176_v55_
                        nw11_[int(3)] = d_176_v55_
                        nw11_[int(4)] = d_175_v54_
                        nw11_[int(5)] = d_176_v55_
                        nw11_[int(6)] = d_176_v55_
                        nw11_[int(7)] = d_176_v55_
                        nw11_[int(8)] = d_176_v55_
                        nw11_[int(9)] = d_176_v55_
                        nw11_[int(10)] = d_176_v55_
                        nw11_[int(11)] = d_176_v55_
                        nw11_[int(12)] = d_176_v55_
                        nw11_[int(13)] = d_175_v54_
                        nw11_[int(14)] = d_175_v54_
                        nw11_[int(15)] = d_176_v55_
                        nw11_[int(16)] = d_176_v55_
                        nw11_[int(17)] = d_176_v55_
                        nw11_[int(18)] = d_176_v55_
                        nw11_[int(19)] = d_175_v54_
                        nw11_[int(20)] = d_175_v54_
                        nw11_[int(21)] = d_176_v55_
                        nw11_[int(22)] = d_176_v55_
                        d_177_v56_ = nw11_
                        d_178_v57_: _dafny.Map
                        d_178_v57_ = _dafny.Map({d_177_v56_: d_172_v53_})
                        d_179_v58_: D11
                        d_179_v58_ = D11_DC22(d_178_v57_)
                        d_179_v58_ = d_179_v58_
                    elif True:
                        d_180_v59_: str
                        d_180_v59_ = _dafny.CodePoint('k')
                        d_181_v60_: _dafny.Seq
                        d_181_v60_ = _dafny.SeqWithoutIsStrInference([d_180_v59_, d_180_v59_, d_180_v59_, d_180_v59_, d_180_v59_])
                        (d_116_globalState_).f1 = (_dafny.MultiSet(d_181_v60_)).cardinality
                        d_110_v0_ = d_110_v0_
                        d_182_v61_: _dafny.MultiSet
                        d_182_v61_ = _dafny.MultiSet([d_112_v2_])
                        d_183_v62_: _dafny.Array
                        nw12_ = _dafny.Array(None, 16)
                        nw12_[int(0)] = (d_182_v61_) | (d_182_v61_)
                        nw12_[int(1)] = (_dafny.MultiSet([d_112_v2_, d_112_v2_, d_112_v2_])) - (_dafny.MultiSet([d_112_v2_, d_112_v2_, d_112_v2_, d_112_v2_, _dafny.SeqWithoutIsStrInference([not(d_110_v0_)])]))
                        nw12_[int(2)] = d_182_v61_
                        nw12_[int(3)] = (d_182_v61_).set(d_112_v2_, default__.abs(default__.fm2(d_111_v1_, d_111_v1_, d_116_globalState_)))
                        nw12_[int(4)] = d_182_v61_
                        nw12_[int(5)] = _dafny.MultiSet([d_112_v2_])
                        nw12_[int(6)] = (d_182_v61_) - (d_182_v61_)
                        nw12_[int(7)] = d_182_v61_
                        nw12_[int(8)] = d_182_v61_
                        nw12_[int(9)] = d_182_v61_
                        nw12_[int(10)] = (d_182_v61_).intersection(_dafny.MultiSet([d_112_v2_, (d_112_v2_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_110_v0_])), len(d_112_v2_)), True)]))
                        nw12_[int(11)] = d_182_v61_
                        nw12_[int(12)] = d_182_v61_
                        nw12_[int(13)] = d_182_v61_
                        nw12_[int(14)] = d_182_v61_
                        nw12_[int(15)] = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_110_v0_])])) | (d_182_v61_)
                        d_183_v62_ = nw12_
                        index5_ = default__.safeIndex(861, (d_183_v62_).length(0))
                        (d_183_v62_)[index5_] = (_dafny.MultiSet([d_112_v2_, d_112_v2_])).intersection(d_182_v61_)
                        d_184_v63_: _dafny.Seq
                        d_184_v63_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_110_v0_, d_110_v0_, True, d_110_v0_, d_110_v0_])])
                        d_185_v64_: _dafny.MultiSet
                        d_185_v64_ = _dafny.MultiSet([False])
                        d_186_v65_: _dafny.MultiSet
                        d_186_v65_ = _dafny.MultiSet([d_111_v1_, d_111_v1_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymwi"))).set(default__.safeIndex(d_111_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymwi")))), d_180_v59_))])
                        d_187_v66_: _dafny.Map
                        d_187_v66_ = _dafny.Map({(d_186_v65_).cardinality: d_112_v2_})
                        index6_ = default__.safeIndex(861, (d_183_v62_).length(0))
                        (d_183_v62_)[index6_] = (_dafny.MultiSet(((d_184_v63_) + (_dafny.SeqWithoutIsStrInference([d_112_v2_]))).set(default__.safeIndex((d_185_v64_).cardinality, len((d_184_v63_) + (_dafny.SeqWithoutIsStrInference([d_112_v2_])))), ((d_187_v66_)[d_111_v1_] if (d_111_v1_) in (d_187_v66_) else d_112_v2_)))).set(d_112_v2_, default__.abs(default__.fm2(d_111_v1_, d_111_v1_, d_116_globalState_)))
                        d_188_v67_: _dafny.Array
                        nw13_ = _dafny.Array(None, 7)
                        nw13_[int(0)] = d_110_v0_
                        nw13_[int(1)] = (d_110_v0_) not in (d_112_v2_)
                        nw13_[int(2)] = not(d_110_v0_)
                        nw13_[int(3)] = (d_110_v0_ if d_110_v0_ else d_110_v0_)
                        nw13_[int(4)] = d_110_v0_
                        nw13_[int(5)] = d_110_v0_
                        nw13_[int(6)] = d_110_v0_
                        d_188_v67_ = nw13_
                        index7_ = default__.safeIndex(244, (d_188_v67_).length(0))
                        (d_188_v67_)[index7_] = d_110_v0_
                        d_189_v68_: _dafny.Set
                        d_189_v68_ = _dafny.Set({d_110_v0_, d_110_v0_})
                        d_190_v69_: T1
                        nw14_ = C8()
                        nw14_.ctor__(d_110_v0_, d_189_v68_, (d_186_v65_).cardinality)
                        d_190_v69_ = nw14_
                        d_191_v70_: _dafny.Set
                        d_191_v70_ = _dafny.Set({d_190_v69_})
                        index8_ = default__.safeIndex(244, (d_188_v67_).length(0))
                        rhs12_ = default__.safeDivisionInt(len(d_112_v2_), (0) - ((len(d_191_v70_)) - (d_111_v1_)))
                        rhs13_ = ((d_113_v3_)[((_dafny.MultiSet([(d_190_v69_).f11])).cardinality) - ((d_190_v69_).f11)] if (((_dafny.MultiSet([(d_190_v69_).f11])).cardinality) - ((d_190_v69_).f11)) in (d_113_v3_) else (0) - (len(_dafny.Set({d_110_v0_}))))
                        rhs14_ = (default__.safeDivisionInt(d_111_v1_, d_111_v1_)) > (d_111_v1_)
                        lhs7_ = d_116_globalState_
                        lhs8_ = d_116_globalState_
                        lhs9_ = d_188_v67_
                        lhs10_ = default__.safeIndex(244, (d_188_v67_).length(0))
                        lhs7_.f1 = rhs12_
                        lhs8_.f1 = rhs13_
                        lhs9_[lhs10_] = rhs14_
                        d_192_v71_: D0
                        d_192_v71_ = D0_DC1(False, (d_190_v69_).f11, d_110_v0_)
                        d_193_v72_: _dafny.Array
                        nw15_ = _dafny.Array(None, 1)
                        nw15_[int(0)] = default__.safeDivisionInt(d_111_v1_, (d_192_v71_).cf2)
                        d_193_v72_ = nw15_
                        index9_ = default__.safeIndex(415, (d_193_v72_).length(0))
                        (d_193_v72_)[index9_] = default__.safeModuloInt((d_190_v69_).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "day"))))
                        index10_ = default__.safeIndex(415, (d_193_v72_).length(0))
                        rhs15_ = default__.safeModuloInt((d_190_v69_).f11, d_111_v1_)
                        rhs16_ = d_180_v59_
                        rhs17_ = (d_190_v69_).f11
                        lhs11_ = d_193_v72_
                        lhs12_ = default__.safeIndex(415, (d_193_v72_).length(0))
                        lhs11_[lhs12_] = rhs15_
                        d_180_v59_ = rhs16_
                        d_111_v1_ = rhs17_
                    d_194_v73_: _dafny.Seq
                    d_194_v73_ = _dafny.SeqWithoutIsStrInference([d_111_v1_])
                    d_195_v74_: _dafny.MultiSet
                    d_195_v74_ = _dafny.MultiSet([d_111_v1_])
                    d_196_v75_: _dafny.Array
                    nw16_ = _dafny.Array(None, 4)
                    nw16_[int(0)] = (_dafny.MultiSet(d_194_v73_)).set(d_111_v1_, default__.abs(d_111_v1_))
                    nw16_[int(1)] = _dafny.MultiSet([d_111_v1_, d_111_v1_])
                    nw16_[int(2)] = d_195_v74_
                    nw16_[int(3)] = d_195_v74_
                    d_196_v75_ = nw16_
                    d_197_v76_: _dafny.Array
                    d_197_v76_ = d_196_v75_
                    d_197_v76_ = d_197_v76_
                    pass
            pass
        d_198_v77_: _dafny.Seq
        d_198_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyaxksq"))
        d_199_v78_: D13
        d_199_v78_ = D13_DC28((d_198_v77_) + (d_198_v77_), (0) - (default__.safeDivisionInt(d_111_v1_, 668)), d_110_v0_)
        pat_let_tv4_ = d_110_v0_
        pat_let_tv5_ = d_198_v77_
        def iife36_(_pat_let4_0):
            def iife37_(d_200_dt__update__tmp_h3_):
                def iife38_(_pat_let5_0):
                    def iife39_(d_201_dt__update_hcf42_h0_):
                        def iife40_(_pat_let6_0):
                            def iife41_(d_202_dt__update_hcf40_h0_):
                                return D13_DC28(d_202_dt__update_hcf40_h0_, (d_200_dt__update__tmp_h3_).cf41, d_201_dt__update_hcf42_h0_)
                            return iife41_(_pat_let6_0)
                        return iife40_(pat_let_tv5_)
                    return iife39_(_pat_let5_0)
                return iife38_(pat_let_tv4_)
            return iife37_(_pat_let4_0)
        d_199_v78_ = iife36_(d_199_v78_)
        if d_110_v0_:
            d_203_v79_: _dafny.Array
            def lambda4_(d_204_v1_):
                def lambda5_(d_205_i4_):
                    return (d_205_i4_) + (d_204_v1_)

                return lambda5_

            init2_ = lambda4_(d_111_v1_)
            nw17_ = _dafny.Array(None, 2)
            for i0_2_ in range(nw17_.length(0)):
                nw17_[i0_2_] = init2_(i0_2_)
            d_203_v79_ = nw17_
            d_206_v80_: D19
            d_206_v80_ = D19_DC42(d_203_v79_)
            d_203_v79_ = (d_206_v80_).cf68
            d_207_v81_: _dafny.Array
            nw18_ = _dafny.Array(None, 17)
            d_207_v81_ = nw18_
            d_208_v82_: _dafny.MultiSet
            d_208_v82_ = _dafny.MultiSet([d_111_v1_, d_111_v1_])
            d_209_v83_: _dafny.Set
            d_209_v83_ = _dafny.Set({d_110_v0_, d_110_v0_})
            d_210_v84_: C4
            nw19_ = C4()
            nw19_.ctor__(d_208_v82_, d_110_v0_, d_209_v83_, len(d_113_v3_))
            d_210_v84_ = nw19_
            d_211_v85_: _dafny.Map
            d_211_v85_ = _dafny.Map({d_203_v79_: d_210_v84_})
            index11_ = default__.safeIndex(109, (d_207_v81_).length(0))
            (d_207_v81_)[index11_] = ((d_211_v85_)[d_203_v79_] if (d_203_v79_) in (d_211_v85_) else d_210_v84_)
            d_212_v86_: _dafny.Seq
            d_212_v86_ = _dafny.SeqWithoutIsStrInference([len(d_112_v2_)])
            index12_ = default__.safeIndex(109, (d_207_v81_).length(0))
            rhs18_ = d_210_v84_
            rhs19_ = (d_210_v84_).f17
            rhs20_ = (d_212_v86_) + (d_212_v86_)
            lhs13_ = d_207_v81_
            lhs14_ = default__.safeIndex(109, (d_207_v81_).length(0))
            lhs15_ = d_116_globalState_
            lhs13_[lhs14_] = rhs18_
            lhs15_.f7 = rhs19_
            d_212_v86_ = rhs20_
            d_213_v87_: C1
            nw20_ = C1()
            nw20_.ctor__(d_110_v0_, (d_210_v84_).f17)
            d_213_v87_ = nw20_
            d_214_v88_: _dafny.Set
            d_214_v88_ = _dafny.Set({d_213_v87_, d_213_v87_, d_213_v87_})
            if (default__.safeModuloInt(len(d_214_v88_), d_111_v1_)) == (d_111_v1_):
                (d_116_globalState_).f6 = (d_210_v84_).f17
                (d_116_globalState_).f1 = len((_dafny.SeqWithoutIsStrInference([(d_210_v84_).f17])) + (d_112_v2_))
                d_215_v89_: _dafny.Array
                nw21_ = _dafny.Array(None, 11)
                nw21_[int(0)] = d_198_v77_
                nw21_[int(1)] = default__.fm1(d_116_globalState_)
                nw21_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_216_i5_ in range(default__.abs(722))])
                nw21_[int(3)] = d_198_v77_
                nw21_[int(4)] = d_198_v77_
                nw21_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_217_i6_ in range(default__.abs(181))])
                nw21_[int(6)] = d_198_v77_
                nw21_[int(7)] = default__.fm1(d_116_globalState_)
                nw21_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltloi"))
                nw21_[int(9)] = d_198_v77_
                nw21_[int(10)] = d_198_v77_
                d_215_v89_ = nw21_
                index13_ = default__.safeIndex(893, (d_215_v89_).length(0))
                (d_215_v89_)[index13_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brktfp"))
                d_218_v90_: _dafny.MultiSet
                d_218_v90_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tk"))) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_219_i7_ in range(default__.abs(121))]))])
                d_220_v91_: D24
                d_220_v91_ = D24_DC58((d_210_v84_).f16)
                d_221_v92_: _dafny.Map
                d_221_v92_ = _dafny.Map({(d_220_v91_).cf96: _dafny.Map({len(d_212_v86_): default__.fm2(d_111_v1_, len(d_112_v2_), d_116_globalState_)})})
                index14_ = default__.safeIndex(893, (d_215_v89_).length(0))
                rhs21_ = (d_198_v77_) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_222_i8_ in range(default__.abs(-425))])) + (d_198_v77_))
                rhs22_ = (d_198_v77_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_223_i9_ in range(default__.abs(792))]))
                rhs23_ = (d_210_v84_).fm4((d_111_v1_) * (d_111_v1_), d_221_v92_, d_111_v1_, d_110_v0_, d_116_globalState_)
                rhs24_ = d_111_v1_
                rhs25_ = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(d_213_v87_).f18])) + (d_112_v2_))
                lhs16_ = d_215_v89_
                lhs17_ = default__.safeIndex(893, (d_215_v89_).length(0))
                lhs18_ = d_116_globalState_
                lhs19_ = d_116_globalState_
                d_198_v77_ = rhs21_
                lhs16_[lhs17_] = rhs22_
                lhs18_.f7 = rhs23_
                lhs19_.f1 = rhs24_
                d_218_v90_ = rhs25_
                d_224_v93_: str
                d_224_v93_ = _dafny.CodePoint('a')
                d_225_v94_: _dafny.Map
                d_225_v94_ = _dafny.Map({d_111_v1_: (d_210_v84_).f16})
                d_226_v95_: _dafny.Map
                d_226_v95_ = _dafny.Map({len(d_225_v94_): (d_213_v87_).f19})
                d_227_v96_: _dafny.Map
                d_227_v96_ = _dafny.Map({d_224_v93_: d_226_v95_})
                d_228_v97_: _dafny.Array
                d_229_v98_: _dafny.Map
                out4_: _dafny.Array
                out5_: _dafny.Map
                out4_, out5_ = (d_210_v84_).m1(len(((d_227_v96_)[d_224_v93_] if (d_224_v93_) in (d_227_v96_) else d_226_v95_)), d_116_globalState_)
                d_228_v97_ = out4_
                d_229_v98_ = out5_
                rhs26_ = len(d_198_v77_)
                rhs27_ = (d_111_v1_) + (d_111_v1_)
                rhs28_ = (0) - (default__.safeDivisionInt(d_111_v1_, (0) - (d_111_v1_)))
                rhs29_ = 271
                rhs30_ = (d_213_v87_).f18
                lhs20_ = d_116_globalState_
                lhs21_ = d_116_globalState_
                lhs22_ = d_116_globalState_
                lhs23_ = d_116_globalState_
                lhs24_ = d_116_globalState_
                lhs20_.f1 = rhs26_
                lhs21_.f1 = rhs27_
                lhs22_.f1 = rhs28_
                lhs23_.f1 = rhs29_
                lhs24_.f7 = rhs30_
            elif True:
                d_230_v99_: _dafny.Map
                d_230_v99_ = _dafny.Map({(d_213_v87_).f18: (d_213_v87_).f18})
                d_230_v99_ = (d_230_v99_).set(True, (d_213_v87_).f18)
                d_231_v100_: _dafny.Array
                d_232_v101_: _dafny.Map
                out6_: _dafny.Array
                out7_: _dafny.Map
                out6_, out7_ = (d_210_v84_).m1(d_111_v1_, d_116_globalState_)
                d_231_v100_ = out6_
                d_232_v101_ = out7_
                index15_ = default__.safeIndex(565, (d_203_v79_).length(0))
                (d_203_v79_)[index15_] = default__.safeDivisionInt(d_111_v1_, d_111_v1_)
                d_233_v102_: _dafny.Set
                d_233_v102_ = _dafny.Set({d_111_v1_})
                index16_ = default__.safeIndex(565, (d_203_v79_).length(0))
                (d_203_v79_)[index16_] = (d_210_v84_).fm3(d_111_v1_, ((0) - (len(d_233_v102_))) > (default__.fm2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))), d_111_v1_, d_116_globalState_)), (d_213_v87_).f19, d_116_globalState_)
                d_234_v103_: str
                d_234_v103_ = _dafny.CodePoint('p')
                d_235_v104_: T0
                nw22_ = C2()
                nw22_.ctor__(d_234_v103_, d_111_v1_)
                d_235_v104_ = nw22_
                d_236_v105_: _dafny.Array
                def lambda6_(d_237_v87_):
                    def lambda7_(d_238_i10_):
                        return (d_237_v87_).f19

                    return lambda7_

                init3_ = lambda6_(d_213_v87_)
                nw23_ = _dafny.Array(None, 14)
                for i0_3_ in range(nw23_.length(0)):
                    nw23_[i0_3_] = init3_(i0_3_)
                d_236_v105_ = nw23_
                d_239_v106_: _dafny.Map
                d_239_v106_ = _dafny.Map({d_236_v105_: len(d_212_v86_)})
                nw24_ = C2()
                nw24_.ctor__(d_234_v103_, ((d_239_v106_)[d_236_v105_] if (d_236_v105_) in (d_239_v106_) else d_111_v1_))
                d_235_v104_ = nw24_
                index17_ = default__.safeIndex(565, (d_203_v79_).length(0))
                (d_203_v79_)[index17_] = 146
            d_240_v107_: _dafny.Array
            d_241_v108_: _dafny.Map
            out8_: _dafny.Array
            out9_: _dafny.Map
            out8_, out9_ = (d_210_v84_).m1(d_111_v1_, d_116_globalState_)
            d_240_v107_ = out8_
            d_241_v108_ = out9_
            index18_ = default__.safeIndex(414, (d_203_v79_).length(0))
            (d_203_v79_)[index18_] = (d_210_v84_).fm6(d_110_v0_, d_111_v1_, d_116_globalState_)
            index19_ = default__.safeIndex(414, (d_203_v79_).length(0))
            rhs31_ = (d_111_v1_) - (default__.safeDivisionInt(len(d_198_v77_), d_111_v1_))
            rhs32_ = d_111_v1_
            rhs33_ = (d_111_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))))
            lhs25_ = d_203_v79_
            lhs26_ = default__.safeIndex(414, (d_203_v79_).length(0))
            lhs27_ = d_116_globalState_
            lhs25_[lhs26_] = rhs31_
            d_111_v1_ = rhs32_
            lhs27_.f1 = rhs33_
        elif True:
            d_242_v109_: _dafny.MultiSet
            d_242_v109_ = _dafny.MultiSet([d_111_v1_])
            d_243_v110_: _dafny.Set
            d_243_v110_ = _dafny.Set({d_110_v0_})
            d_244_v111_: C3
            nw25_ = C3()
            nw25_.ctor__((d_111_v1_ if d_110_v0_ else len(d_114_v4_)), d_111_v1_, default__.safeModuloInt((d_242_v109_).cardinality, d_111_v1_), d_243_v110_)
            d_244_v111_ = nw25_
            d_244_v111_ = d_244_v111_
            d_245_v112_: _dafny.Array
            nw26_ = _dafny.Array(int(0), 29)
            d_245_v112_ = nw26_
            d_246_v113_: str
            d_246_v113_ = _dafny.CodePoint('r')
            d_247_v114_: _dafny.Map
            d_247_v114_ = _dafny.Map({d_245_v112_: d_246_v113_})
            index20_ = default__.safeIndex(485, (d_245_v112_).length(0))
            (d_245_v112_)[index20_] = (d_244_v111_.f23) - (((d_113_v3_)[len(d_247_v114_)] if (len(d_247_v114_)) in (d_113_v3_) else (d_244_v111_).f22))
            index21_ = default__.safeIndex(485, (d_245_v112_).length(0))
            (d_245_v112_)[index21_] = (d_244_v111_).f22
            d_111_v1_ = default__.safeDivisionInt(-630, ((d_244_v111_).f22 if d_110_v0_ else (0) - ((d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))])))
            d_248_v115_: _dafny.Array
            d_249_v116_: _dafny.Map
            out10_: _dafny.Array
            out11_: _dafny.Map
            out10_, out11_ = (d_244_v111_).m1((d_244_v111_).f22, d_116_globalState_)
            d_248_v115_ = out10_
            d_249_v116_ = out11_
            d_250_v117_: _dafny.Seq
            d_250_v117_ = _dafny.SeqWithoutIsStrInference([d_111_v1_])
            source5_ = D7_DC13((d_243_v110_) | (d_243_v110_), (d_244_v111_).f22, len(d_250_v117_))
            if source5_.is_DC13:
                d_251___mcc_h0_ = source5_.cf18
                d_252___mcc_h1_ = source5_.cf19
                d_253___mcc_h2_ = source5_.cf20
                d_254_cf20_ = d_253___mcc_h2_
                d_255_cf19_ = d_252___mcc_h1_
                d_256_cf18_ = d_251___mcc_h0_
                index22_ = default__.safeIndex(485, (d_245_v112_).length(0))
                (d_245_v112_)[index22_] = (d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))]
                d_257_v118_: D22
                d_257_v118_ = D22_DC52()
                d_257_v118_ = d_257_v118_
                d_258_v119_: _dafny.Map
                d_258_v119_ = _dafny.Map({d_110_v0_: True})
                d_259_v120_: _dafny.Set
                d_259_v120_ = _dafny.Set({((d_242_v109_)[default__.fm2((d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))], (d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))], d_116_globalState_)] if (default__.fm2((d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))], (d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))], d_116_globalState_)) in (d_242_v109_) else d_254_cf20_), d_244_v111_.f23, (866) + (72), default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([d_111_v1_ for d_260_i11_ in range(default__.abs(807))])).set(default__.safeIndex(len(d_258_v119_), len(_dafny.SeqWithoutIsStrInference([d_111_v1_ for d_261_i11_ in range(default__.abs(807))]))), (0) - (d_244_v111_.f23))), d_111_v1_), len(d_258_v119_)})
                d_259_v120_ = _dafny.Set({d_111_v1_})
                d_262_v122_: _dafny.Seq
                def iife42_():
                    coll28_ = _dafny.Set()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(113, -958):
                        d_263_v121_: int = compr_28_
                        if ((113) <= (d_263_v121_)) and ((d_263_v121_) < (-958)):
                            coll28_ = coll28_.union(_dafny.Set([(d_263_v121_) * (d_111_v1_)]))
                    return _dafny.Set(coll28_)
                d_262_v122_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({((d_242_v109_)[(d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))]] if ((d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))]) in (d_242_v109_) else (d_244_v111_).f22)}), _dafny.Set({261}), iife42_()
                ])
                d_262_v122_ = d_262_v122_
            elif source5_.is_DC12:
                d_264___mcc_h3_ = source5_.cf17
                d_265_cf17_ = d_264___mcc_h3_
                d_266_v123_: _dafny.Map
                d_266_v123_ = _dafny.Map({d_110_v0_: (d_198_v77_) + (d_198_v77_)})
                d_266_v123_ = (d_266_v123_).set(False, (d_198_v77_) + (d_198_v77_))
                (d_116_globalState_).f6 = not((default__.fm2(len(d_198_v77_), d_244_v111_.f23, d_116_globalState_)) < ((d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))]))
                d_267_v124_: _dafny.Array
                nw27_ = _dafny.Array(False, 26)
                d_267_v124_ = nw27_
                index23_ = default__.safeIndex(915, (d_267_v124_).length(0))
                (d_267_v124_)[index23_] = d_110_v0_
                index24_ = default__.safeIndex(915, (d_267_v124_).length(0))
                (d_267_v124_)[index24_] = default__.fm0(d_116_globalState_)
                d_268_v125_: C4
                nw28_ = C4()
                nw28_.ctor__(d_242_v109_, False, d_243_v110_, (d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))])
                d_268_v125_ = nw28_
                d_269_v126_: D25
                d_269_v126_ = D25_DC61(d_268_v125_)
                d_268_v125_ = (d_269_v126_).cf98
            elif True:
                d_270___mcc_h4_ = source5_.cf21
                d_271_cf21_ = d_270___mcc_h4_
                d_272_v127_: D6
                d_272_v127_ = D6_DC10(_dafny.SeqWithoutIsStrInference([False, d_110_v0_]))
                d_273_v128_: D19
                d_273_v128_ = D19_DC43(d_110_v0_, d_246_v113_, True, d_272_v127_, d_198_v77_)
                d_274_v129_: _dafny.Array
                nw29_ = _dafny.Array(None, 10)
                nw29_[int(0)] = True
                nw29_[int(1)] = d_110_v0_
                nw29_[int(2)] = not((d_244_v111_).fm5(D0_DC0(d_244_v111_.f23), d_112_v2_, d_111_v1_, (d_273_v128_).cf70, d_116_globalState_))
                nw29_[int(3)] = d_110_v0_
                nw29_[int(4)] = d_110_v0_
                nw29_[int(5)] = True
                nw29_[int(6)] = d_110_v0_
                nw29_[int(7)] = d_110_v0_
                nw29_[int(8)] = False
                nw29_[int(9)] = False
                d_274_v129_ = nw29_
                d_275_v130_: _dafny.Set
                d_275_v130_ = _dafny.Set({d_274_v129_, d_274_v129_, d_274_v129_})
                d_275_v130_ = (_dafny.Set({d_274_v129_, d_274_v129_, d_274_v129_})) | (d_275_v130_)
                d_111_v1_ = d_111_v1_
                def iife43_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in (default__.fm27(d_110_v0_, d_110_v0_, d_116_globalState_)).Elements:
                        d_276_v131_: int = compr_29_
                        if (d_276_v131_) in (default__.fm27(d_110_v0_, d_110_v0_, d_116_globalState_)):
                            coll29_ = coll29_.union(_dafny.Set([default__.safeDivisionInt(d_276_v131_, -307)]))
                    return _dafny.Set(coll29_)
                (d_116_globalState_).f6 = (d_110_v0_ if d_110_v0_ else (d_244_v111_.f23) not in (iife43_()
                ))
                d_277_v132_: C7
                nw30_ = C7()
                nw30_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_244_v111_.f23, d_111_v1_]) for d_278_i12_ in range(default__.abs(966))]), d_243_v110_, (d_245_v112_)[default__.safeIndex(485, (d_245_v112_).length(0))])
                d_277_v132_ = nw30_
        d_279_v133_: str
        d_279_v133_ = _dafny.CodePoint('o')
        d_280_v134_: D6
        d_280_v134_ = D6_DC10(_dafny.SeqWithoutIsStrInference([d_110_v0_, d_110_v0_]))
        (d_116_globalState_).f7 = (D19_DC43(False, d_279_v133_, d_110_v0_, d_280_v134_, d_198_v77_)).cf71
        d_281_i13_: int
        d_281_i13_ = 0
        with _dafny.label("1"):
            while d_110_v0_:
                with _dafny.c_label("1"):
                    if (d_281_i13_) >= (100):
                        raise _dafny.Break("1")
                    d_281_i13_ = (d_281_i13_) + (1)
                    d_282_v135_: C5
                    nw31_ = C5()
                    nw31_.ctor__()
                    d_282_v135_ = nw31_
                    rhs34_ = ((_dafny.SeqWithoutIsStrInference([d_110_v0_, True, d_110_v0_])) + (_dafny.SeqWithoutIsStrInference([d_110_v0_]))) < (d_112_v2_)
                    rhs35_ = d_111_v1_
                    rhs36_ = d_111_v1_
                    rhs37_ = d_282_v135_
                    rhs38_ = default__.safeDivisionInt((d_111_v1_ if True else d_111_v1_), d_111_v1_)
                    lhs28_ = d_116_globalState_
                    lhs29_ = d_116_globalState_
                    lhs30_ = d_116_globalState_
                    lhs31_ = d_116_globalState_
                    lhs28_.f6 = rhs34_
                    lhs29_.f1 = rhs35_
                    lhs30_.f1 = rhs36_
                    d_282_v135_ = rhs37_
                    lhs31_.f1 = rhs38_
                    d_283_v136_: _dafny.Set
                    d_283_v136_ = _dafny.Set({d_110_v0_, d_110_v0_, d_110_v0_, d_110_v0_})
                    d_284_v137_: C7
                    nw32_ = C7()
                    nw32_.ctor__((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_110_v0_]))]) for d_285_i14_ in range(default__.abs(269))])) + (default__.fm56(d_111_v1_, d_111_v1_, d_116_globalState_)), d_283_v136_, (d_111_v1_) + (d_111_v1_))
                    d_284_v137_ = nw32_
                    d_284_v137_ = d_284_v137_
                    d_286_v138_: _dafny.Array
                    nw33_ = _dafny.Array(int(0), 18)
                    d_286_v138_ = nw33_
                    index25_ = default__.safeIndex(226, (d_286_v138_).length(0))
                    (d_286_v138_)[index25_] = (0) - (len(_dafny.Set({d_111_v1_})))
                    d_287_v139_: C0
                    nw34_ = C0()
                    nw34_.ctor__(d_112_v2_, d_111_v1_)
                    d_287_v139_ = nw34_
                    d_288_v140_: _dafny.MultiSet
                    d_288_v140_ = _dafny.MultiSet([d_287_v139_, d_287_v139_])
                    index26_ = default__.safeIndex(226, (d_286_v138_).length(0))
                    (d_286_v138_)[index26_] = default__.safeDivisionInt(d_111_v1_, ((d_288_v140_).intersection(d_288_v140_)).cardinality)
                    (d_116_globalState_).f7 = d_110_v0_
                    pass
            pass
        (d_116_globalState_).f7 = True
        (d_116_globalState_).f1 = len((_dafny.SeqWithoutIsStrInference([d_279_v133_ for d_289_i15_ in range(default__.abs(812))])) + ((d_198_v77_).set(default__.safeIndex(d_111_v1_, len(d_198_v77_)), d_279_v133_)))
        d_290_v141_: _dafny.Array
        nw35_ = _dafny.Array(int(0), 29)
        d_290_v141_ = nw35_
        index27_ = default__.safeIndex(23, (d_290_v141_).length(0))
        (d_290_v141_)[index27_] = len(d_112_v2_)
        d_291_v142_: _dafny.Seq
        d_291_v142_ = _dafny.SeqWithoutIsStrInference([d_111_v1_])
        d_292_v143_: _dafny.Seq
        d_292_v143_ = _dafny.SeqWithoutIsStrInference([d_111_v1_, (0) - (d_111_v1_), len(d_291_v142_), d_111_v1_])
        index28_ = default__.safeIndex(23, (d_290_v141_).length(0))
        rhs39_ = d_111_v1_
        rhs40_ = not(not((d_291_v142_) != (_dafny.SeqWithoutIsStrInference([d_111_v1_ for d_293_i16_ in range(default__.abs(-502))]))))
        lhs32_ = d_290_v141_
        lhs33_ = default__.safeIndex(23, (d_290_v141_).length(0))
        lhs34_ = d_116_globalState_
        lhs32_[lhs33_] = rhs39_
        lhs34_.f7 = rhs40_
        d_294_v145_: D8
        d_294_v145_ = D8_DC16(True, d_111_v1_)
        d_295_v146_: D8
        d_295_v146_ = D8_DC17(d_294_v145_)
        d_296_v147_: _dafny.Seq
        d_296_v147_ = _dafny.SeqWithoutIsStrInference([d_295_v146_, d_295_v146_])
        d_297_v148_: _dafny.Map
        def iife44_():
            coll30_ = _dafny.Map()
            compr_30_: D8
            for compr_30_ in (_dafny.MultiSet((d_296_v147_).set(default__.safeIndex((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], len(d_296_v147_)), d_295_v146_))).Elements:
                d_298_v144_: D8 = compr_30_
                if (d_298_v144_) in (_dafny.MultiSet((d_296_v147_).set(default__.safeIndex((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], len(d_296_v147_)), d_295_v146_))):
                    coll30_[d_298_v144_] = (d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]
            return _dafny.Map(coll30_)
        d_297_v148_ = _dafny.Map({d_279_v133_: iife44_()
        })
        d_299_v149_: _dafny.Map
        d_299_v149_ = _dafny.Map({D8_DC17(d_294_v145_): default__.fm2((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], len(d_292_v143_), d_116_globalState_)})
        d_300_v150_: D13
        d_300_v150_ = D13_DC27(((d_297_v148_)[d_279_v133_] if (d_279_v133_) in (d_297_v148_) else d_299_v149_))
        d_301_v151_: _dafny.Map
        d_301_v151_ = _dafny.Map({(d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]: d_300_v150_})
        d_301_v151_ = (d_301_v151_).set((0) - ((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]), d_300_v150_)
        d_302_v152_: _dafny.MultiSet
        d_302_v152_ = _dafny.MultiSet([d_111_v1_])
        d_303_v153_: C4
        nw36_ = C4()
        nw36_.ctor__(d_302_v152_, not(d_110_v0_), _dafny.Set({d_110_v0_}), (d_111_v1_ if d_110_v0_ else (d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]))
        d_303_v153_ = nw36_
        d_304_v154_: _dafny.Array
        nw37_ = _dafny.Array(False, 11)
        d_304_v154_ = nw37_
        index29_ = default__.safeIndex(805, (d_304_v154_).length(0))
        (d_304_v154_)[index29_] = d_110_v0_
        index30_ = default__.safeIndex(805, (d_304_v154_).length(0))
        rhs41_ = d_110_v0_
        rhs42_ = (d_303_v153_ if (d_303_v153_).f17 else d_303_v153_)
        rhs43_ = (d_303_v153_).f17
        rhs44_ = (d_303_v153_).f17
        lhs35_ = d_116_globalState_
        lhs36_ = d_116_globalState_
        lhs37_ = d_304_v154_
        lhs38_ = default__.safeIndex(805, (d_304_v154_).length(0))
        lhs35_.f7 = rhs41_
        d_303_v153_ = rhs42_
        lhs36_.f7 = rhs43_
        lhs37_[lhs38_] = rhs44_
        (d_116_globalState_).f1 = default__.safeModuloInt(d_111_v1_, 223)
        d_305_v155_: C5
        nw38_ = C5()
        nw38_.ctor__()
        d_305_v155_ = nw38_
        d_306_v156_: _dafny.Seq
        d_306_v156_ = _dafny.SeqWithoutIsStrInference([d_305_v155_])
        d_307_v157_: D20
        d_307_v157_ = D20_DC44((d_306_v156_)[default__.safeIndex(len(d_198_v77_), len(d_306_v156_))])
        d_307_v157_ = d_307_v157_
        d_308_v159_: _dafny.Map
        d_308_v159_ = _dafny.Map({True: True})
        d_309_v160_: _dafny.Set
        d_309_v160_ = _dafny.Set({_dafny.MultiSet([len(d_308_v159_), d_111_v1_])})
        d_310_v161_: _dafny.MultiSet
        def iife45_():
            coll31_ = _dafny.Map()
            compr_31_: _dafny.MultiSet
            for compr_31_ in (d_309_v160_).Elements:
                d_311_v158_: _dafny.MultiSet = compr_31_
                if (d_311_v158_) in (d_309_v160_):
                    coll31_[d_311_v158_] = d_113_v3_
            return _dafny.Map(coll31_)
        d_310_v161_ = _dafny.MultiSet([(d_303_v153_).fm4(d_111_v1_, iife45_()
        , d_111_v1_, (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], d_116_globalState_), (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))]])
        d_312_v163_: D22
        def iife46_():
            coll32_ = _dafny.Set()
            compr_32_: int
            for compr_32_ in (_dafny.Map({(d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]: (d_303_v153_).f17})).keys.Elements:
                d_313_v162_: int = compr_32_
                if (d_313_v162_) in (_dafny.Map({(d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]: (d_303_v153_).f17})):
                    coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_313_v162_, -545)]))
            return _dafny.Set(coll32_)
        d_312_v163_ = D22_DC51(iife46_()
)
        d_314_v164_: _dafny.Seq
        d_314_v164_ = _dafny.SeqWithoutIsStrInference([d_312_v163_])
        (d_116_globalState_).f6 = (len(d_291_v142_)) > (default__.safeModuloInt(((d_310_v161_)[(d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))]] if ((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))]) in (d_310_v161_) else (0) - (d_111_v1_)), (_dafny.MultiSet(d_314_v164_)).cardinality))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_290_v141_).length(0)):
            d_315_i17_: int = guard_loop_0_
            if (True) and (((0) <= (d_315_i17_)) and ((d_315_i17_) < ((d_290_v141_).length(0)))):
                def iife47_():
                    coll33_ = _dafny.Set()
                    compr_33_: str
                    for compr_33_ in (_dafny.MultiSet([d_279_v133_])).Elements:
                        d_316_v165_: str = compr_33_
                        if (d_316_v165_) in (_dafny.MultiSet([d_279_v133_])):
                            coll33_ = coll33_.union(_dafny.Set([d_316_v165_]))
                    return _dafny.Set(coll33_)
                (d_290_v141_)[(d_315_i17_)] = (d_315_i17_) - (default__.safeDivisionInt(len(_dafny.Set({len(iife47_()
                ), d_111_v1_})), len(d_112_v2_)))
        d_317_v166_: D24
        d_317_v166_ = D24_DC58((d_303_v153_).f16)
        d_318_i18_: int
        d_318_i18_ = 0
        with _dafny.label("2"):
            pat_let_tv6_ = d_292_v143_
            pat_let_tv7_ = d_292_v143_
            pat_let_tv8_ = d_111_v1_
            pat_let_tv9_ = d_304_v154_
            pat_let_tv10_ = d_304_v154_
            pat_let_tv11_ = d_292_v143_
            pat_let_tv12_ = d_111_v1_
            pat_let_tv13_ = d_292_v143_
            pat_let_tv14_ = d_290_v141_
            pat_let_tv15_ = d_290_v141_
            pat_let_tv16_ = d_290_v141_
            pat_let_tv17_ = d_290_v141_
            pat_let_tv18_ = d_111_v1_
            pat_let_tv19_ = d_290_v141_
            pat_let_tv20_ = d_290_v141_
            def lambda8_(source6_):
                if source6_.is_DC59:
                    d_340___mcc_h5_ = source6_.cf97
                    d_341_cf97_ = d_340___mcc_h5_
                    return (_dafny.SeqWithoutIsStrInference([pat_let_tv6_, pat_let_tv7_])) != (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([pat_let_tv8_])]))
                elif source6_.is_DC60:
                    return (pat_let_tv10_)[default__.safeIndex(805, (pat_let_tv9_).length(0))]
                elif True:
                    d_342___mcc_h6_ = source6_.cf96
                    d_343_cf96_ = d_342___mcc_h6_
                    return ((pat_let_tv11_).set(default__.safeIndex(pat_let_tv12_, len(pat_let_tv13_)), (pat_let_tv15_)[default__.safeIndex(23, (pat_let_tv14_).length(0))])) < (_dafny.SeqWithoutIsStrInference([(pat_let_tv17_)[default__.safeIndex(23, (pat_let_tv16_).length(0))], pat_let_tv18_, (pat_let_tv20_)[default__.safeIndex(23, (pat_let_tv19_).length(0))]]))

            while lambda8_((D24_DC58((d_303_v153_).f16) if (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))] else d_317_v166_)):
                with _dafny.c_label("2"):
                    if (d_318_i18_) >= (100):
                        raise _dafny.Break("2")
                    d_318_i18_ = (d_318_i18_) + (1)
                    d_319_v167_: C0
                    nw39_ = C0()
                    nw39_.ctor__((d_112_v2_) + (d_112_v2_), default__.safeModuloInt((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], ((d_303_v153_).f16).cardinality))
                    d_319_v167_ = nw39_
                    d_320_v168_: _dafny.Array
                    nw40_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_320_v168_ = nw40_
                    d_321_v169_: _dafny.Array
                    nw41_ = _dafny.Array(None, 27)
                    nw41_[int(0)] = d_198_v77_
                    nw41_[int(1)] = d_198_v77_
                    nw41_[int(2)] = default__.fm1(d_116_globalState_)
                    nw41_[int(3)] = d_198_v77_
                    nw41_[int(4)] = d_198_v77_
                    nw41_[int(5)] = d_198_v77_
                    nw41_[int(6)] = d_198_v77_
                    nw41_[int(7)] = d_198_v77_
                    nw41_[int(8)] = d_198_v77_
                    nw41_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smkfbwtdu"))
                    nw41_[int(10)] = _dafny.SeqWithoutIsStrInference([d_279_v133_ for d_322_i19_ in range(default__.abs(667))])
                    nw41_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfcvhqiw"))
                    nw41_[int(12)] = _dafny.SeqWithoutIsStrInference([d_279_v133_ for d_323_i20_ in range(default__.abs(849))])
                    nw41_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjwo"))
                    nw41_[int(14)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_324_i21_ in range(default__.abs(996))])
                    nw41_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnhdyskyq"))
                    nw41_[int(16)] = d_198_v77_
                    nw41_[int(17)] = _dafny.SeqWithoutIsStrInference([d_279_v133_ for d_325_i22_ in range(default__.abs(-3))])
                    nw41_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykantl"))
                    nw41_[int(19)] = d_198_v77_
                    nw41_[int(20)] = d_198_v77_
                    nw41_[int(21)] = (d_198_v77_).set(default__.safeIndex((d_319_v167_).f21, len(d_198_v77_)), d_279_v133_)
                    nw41_[int(22)] = _dafny.SeqWithoutIsStrInference([d_279_v133_ for d_326_i23_ in range(default__.abs(780))])
                    nw41_[int(23)] = d_198_v77_
                    nw41_[int(24)] = d_198_v77_
                    nw41_[int(25)] = d_198_v77_
                    nw41_[int(26)] = d_198_v77_
                    d_321_v169_ = nw41_
                    index31_ = default__.safeIndex(832, (d_320_v168_).length(0))
                    (d_320_v168_)[index31_] = d_321_v169_
                    index32_ = default__.safeIndex(832, (d_320_v168_).length(0))
                    (d_320_v168_)[index32_] = d_321_v169_
                    d_198_v77_ = ((d_198_v77_) + (d_198_v77_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpbqxjl"))) + (d_198_v77_))
                    if (d_303_v153_).f17:
                        d_327_v170_: _dafny.Seq
                        d_327_v170_ = _dafny.SeqWithoutIsStrInference([d_310_v161_, d_310_v161_, d_310_v161_, d_310_v161_])
                        d_328_v171_: _dafny.Set
                        d_328_v171_ = _dafny.Set({(d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))]})
                        d_329_v172_: _dafny.Array
                        nw42_ = _dafny.Array(None, 25)
                        nw42_[int(0)] = d_310_v161_
                        nw42_[int(1)] = _dafny.MultiSet([(d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], (d_303_v153_).f17])
                        nw42_[int(2)] = (_dafny.MultiSet([not((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))])])) - (d_310_v161_)
                        nw42_[int(3)] = d_310_v161_
                        nw42_[int(4)] = d_310_v161_
                        nw42_[int(5)] = _dafny.MultiSet(d_112_v2_)
                        nw42_[int(6)] = (default__.fm37(d_116_globalState_)) | (d_310_v161_)
                        nw42_[int(7)] = d_310_v161_
                        nw42_[int(8)] = d_310_v161_
                        nw42_[int(9)] = d_310_v161_
                        nw42_[int(10)] = _dafny.MultiSet([(d_319_v167_).fm21(True, d_111_v1_, d_116_globalState_), (d_303_v153_).f17, (d_303_v153_).f17])
                        nw42_[int(11)] = _dafny.MultiSet([(d_303_v153_).f17])
                        nw42_[int(12)] = d_310_v161_
                        nw42_[int(13)] = (d_310_v161_) - (d_310_v161_)
                        nw42_[int(14)] = d_310_v161_
                        nw42_[int(15)] = _dafny.MultiSet([(d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], (d_303_v153_).f17, not(d_110_v0_)])
                        nw42_[int(16)] = d_310_v161_
                        nw42_[int(17)] = (d_310_v161_ if not((d_303_v153_).f17) else _dafny.MultiSet(d_112_v2_))
                        nw42_[int(18)] = (d_327_v170_)[default__.safeIndex(d_111_v1_, len(d_327_v170_))]
                        nw42_[int(19)] = d_310_v161_
                        nw42_[int(20)] = (d_310_v161_) | (d_310_v161_)
                        nw42_[int(21)] = (d_310_v161_) - (default__.fm37(d_116_globalState_))
                        nw42_[int(22)] = (_dafny.MultiSet([d_110_v0_, not((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))])])) - (d_310_v161_)
                        nw42_[int(23)] = d_310_v161_
                        nw42_[int(24)] = (d_310_v161_).set((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], default__.abs(len(d_328_v171_)))
                        d_329_v172_ = nw42_
                        index33_ = default__.safeIndex(352, (d_329_v172_).length(0))
                        (d_329_v172_)[index33_] = d_310_v161_
                        index34_ = default__.safeIndex(352, (d_329_v172_).length(0))
                        (d_329_v172_)[index34_] = (d_310_v161_) | (_dafny.MultiSet([(d_303_v153_).f17]))
                        d_330_v173_: _dafny.Map
                        d_330_v173_ = _dafny.Map({d_114_v4_: not ((d_303_v153_).f17) or ((d_303_v153_).f17)})
                        rhs45_ = (default__.safeModuloInt((d_292_v143_)[default__.safeIndex((d_319_v167_).f21, len(d_292_v143_))], (d_319_v167_).f21)) + (default__.safeModuloInt((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], (d_319_v167_).f21))
                        rhs46_ = (d_303_v153_).f17
                        rhs47_ = ((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))] if (d_303_v153_).f17 else not((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))]))
                        rhs48_ = ((d_319_v167_).f21) - (len(d_198_v77_))
                        rhs49_ = d_330_v173_
                        lhs39_ = d_116_globalState_
                        lhs40_ = d_116_globalState_
                        lhs41_ = d_116_globalState_
                        lhs39_.f1 = rhs45_
                        d_110_v0_ = rhs46_
                        lhs40_.f6 = rhs47_
                        lhs41_.f1 = rhs48_
                        d_330_v173_ = rhs49_
                        d_331_v174_: _dafny.Array
                        nw43_ = _dafny.Array(None, 29)
                        d_331_v174_ = nw43_
                        d_332_v175_: C8
                        nw44_ = C8()
                        nw44_.ctor__(not((d_303_v153_).f17), d_328_v171_, len(d_198_v77_))
                        d_332_v175_ = nw44_
                        index35_ = default__.safeIndex(515, (d_331_v174_).length(0))
                        (d_331_v174_)[index35_] = d_332_v175_
                        index36_ = default__.safeIndex(515, (d_331_v174_).length(0))
                        (d_331_v174_)[index36_] = d_332_v175_
                        d_333_v176_: C3
                        nw45_ = C3()
                        nw45_.ctor__(d_111_v1_, 282, (len(d_198_v77_)) * ((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]), d_328_v171_)
                        d_333_v176_ = nw45_
                        d_334_v177_: _dafny.Array
                        d_335_v178_: _dafny.Map
                        out12_: _dafny.Array
                        out13_: _dafny.Map
                        out12_, out13_ = (d_332_v175_).m1((d_333_v176_.f23) - ((d_332_v175_).fm3((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], d_116_globalState_)), d_116_globalState_)
                        d_334_v177_ = out12_
                        d_335_v178_ = out13_
                    elif True:
                        (d_116_globalState_).f7 = not (d_110_v0_) or ((d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))])
                        d_336_v179_: _dafny.Array
                        nw46_ = _dafny.Array(_dafny.Array(None, 0), 9)
                        d_336_v179_ = nw46_
                        d_337_v180_: _dafny.Array
                        nw47_ = _dafny.Array(None, 6)
                        d_337_v180_ = nw47_
                        index37_ = default__.safeIndex(726, (d_336_v179_).length(0))
                        (d_336_v179_)[index37_] = d_337_v180_
                        d_338_v181_: _dafny.Seq
                        d_338_v181_ = _dafny.SeqWithoutIsStrInference([d_337_v180_])
                        d_339_v182_: _dafny.Map
                        d_339_v182_ = _dafny.Map({(d_303_v153_).f17: d_291_v142_})
                        index38_ = default__.safeIndex(726, (d_336_v179_).length(0))
                        index39_ = default__.safeIndex(23, (d_290_v141_).length(0))
                        rhs50_ = (d_338_v181_)[default__.safeIndex((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], len(d_338_v181_))]
                        rhs51_ = (0) - ((len(d_339_v182_)) - (((d_319_v167_).f21) + ((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))])))
                        rhs52_ = d_111_v1_
                        lhs42_ = d_336_v179_
                        lhs43_ = default__.safeIndex(726, (d_336_v179_).length(0))
                        lhs44_ = d_116_globalState_
                        lhs45_ = d_290_v141_
                        lhs46_ = default__.safeIndex(23, (d_290_v141_).length(0))
                        lhs42_[lhs43_] = rhs50_
                        lhs44_.f1 = rhs51_
                        lhs45_[lhs46_] = rhs52_
                        index40_ = default__.safeIndex(23, (d_290_v141_).length(0))
                        rhs53_ = (d_305_v155_).fm9((d_319_v167_).f21, default__.fm2(d_111_v1_, 325, d_116_globalState_), (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))], d_116_globalState_)
                        rhs54_ = d_290_v141_
                        rhs55_ = (d_319_v167_).f21
                        rhs56_ = (d_303_v153_).fm4(default__.safeDivisionInt(d_111_v1_, d_111_v1_), _dafny.Map({(d_303_v153_).f16: d_113_v3_}), d_111_v1_, not ((d_303_v153_).f17) or ((d_303_v153_).f17), d_116_globalState_)
                        lhs47_ = d_116_globalState_
                        lhs48_ = d_290_v141_
                        lhs49_ = default__.safeIndex(23, (d_290_v141_).length(0))
                        lhs50_ = d_116_globalState_
                        lhs47_.f1 = rhs53_
                        d_290_v141_ = rhs54_
                        lhs48_[lhs49_] = rhs55_
                        lhs50_.f6 = rhs56_
                        (d_116_globalState_).f1 = (0) - ((((d_319_v167_).f21 if (d_303_v153_).f17 else (d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]) if (d_303_v153_).f17 else d_111_v1_))
                        rhs57_ = (d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]
                        rhs58_ = ((d_114_v4_)[d_110_v0_] if (d_110_v0_) in (d_114_v4_) else d_111_v1_)
                        rhs59_ = (d_303_v153_).f17
                        rhs60_ = ((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))]) != ((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))])
                        lhs51_ = d_116_globalState_
                        lhs52_ = d_116_globalState_
                        lhs53_ = d_116_globalState_
                        lhs51_.f1 = rhs57_
                        lhs52_.f1 = rhs58_
                        lhs53_.f7 = rhs59_
                        d_110_v0_ = rhs60_
                    pass
            pass
        d_344_i24_: int
        d_344_i24_ = 0
        with _dafny.label("3"):
            while (d_304_v154_)[default__.safeIndex(805, (d_304_v154_).length(0))]:
                with _dafny.c_label("3"):
                    if (d_344_i24_) >= (100):
                        raise _dafny.Break("3")
                    d_344_i24_ = (d_344_i24_) + (1)
                    d_345_v183_: _dafny.MultiSet
                    out14_: _dafny.MultiSet
                    out14_ = (d_303_v153_).m0(d_279_v133_, d_116_globalState_)
                    d_345_v183_ = out14_
                    d_305_v155_ = d_305_v155_
                    d_290_v141_ = d_290_v141_
                    index41_ = default__.safeIndex(23, (d_290_v141_).length(0))
                    (d_290_v141_)[index41_] = (0) - ((0) - ((0) - (default__.safeModuloInt((d_290_v141_)[default__.safeIndex(23, (d_290_v141_).length(0))], (d_111_v1_ if True else d_111_v1_)))))
                    pass
            pass
        _dafny.print(_dafny.string_of(d_110_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_111_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v2_) == (_dafny.SeqWithoutIsStrInference([True, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v3_) == (_dafny.Map({754: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v4_) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[0]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[1]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[2]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[3]) == (_dafny.Map({False: 754}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[4]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[5]) == (_dafny.Map({False: 754}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[6]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v5_)[7]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_116_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[0]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[1]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[2]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[3]) == (_dafny.Map({False: 754}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[4]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[5]) == (_dafny.Map({False: 754}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[6]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_globalState_.f5)[7]) == (_dafny.Map({False: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_116_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_116_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_117_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_198_v77_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_199_v78_).cf40).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v78_).cf41))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v78_).cf42))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_279_v133_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_280_v134_).cf13) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_281_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v141_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v142_) == (_dafny.SeqWithoutIsStrInference([630]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v143_) == (_dafny.SeqWithoutIsStrInference([630, -630, 1, 630]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v145_).cf23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v145_).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_295_v146_).cf25).cf23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_295_v146_).cf25).cf24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v147_) == (_dafny.SeqWithoutIsStrInference([D8_DC17(D8_DC16(True, 630)), D8_DC17(D8_DC16(True, 630))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v148_) == (_dafny.Map({_dafny.CodePoint('o'): _dafny.Map({D8_DC17(D8_DC16(True, 630)): 630})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_299_v149_) == (_dafny.Map({D8_DC17(D8_DC16(True, 630)): 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v150_).cf39) == (_dafny.Map({D8_DC17(D8_DC16(True, 630)): 630}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_301_v151_) == (_dafny.Map({630: D13_DC27(_dafny.Map({D8_DC17(D8_DC16(True, 630)): 630})), -630: D13_DC27(_dafny.Map({D8_DC17(D8_DC16(True, 630)): 630}))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v152_) == (_dafny.MultiSet([630]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_303_v153_).f16) == (_dafny.MultiSet([630]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_303_v153_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v154_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_306_v156_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_308_v159_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_v160_) == (_dafny.Set({_dafny.MultiSet([1, 630])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v161_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_312_v163_).cf86) == (_dafny.Set({85}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_314_v164_) == (_dafny.SeqWithoutIsStrInference([D22_DC51(_dafny.Set({85}))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_317_v166_).cf96) == (_dafny.MultiSet([630]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_318_i18_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_344_i24_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
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
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({self.cf4.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC6(D3, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)

class D4_DC7(D4, NamedTuple('DC7', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D5_DC8)

class D5_DC8(D5, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC8({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC10(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D6_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D6_DC9)

class D6_DC10(D6, NamedTuple('DC10', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC10({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC10) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC11(D6, NamedTuple('DC11', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC9(D6, NamedTuple('DC9', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC9({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC9) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC13(_dafny.Set({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D7_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D7_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)

class D7_DC13(D7, NamedTuple('DC13', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC13({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC13) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC12(D7, NamedTuple('DC12', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC12({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC12) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC16(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D8_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)

class D8_DC16(D8, NamedTuple('DC16', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC15(D8, NamedTuple('DC15', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC15({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC15) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC19()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)

class D9_DC19(D9, NamedTuple('DC19', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC18(D9, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC21(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D10_DC20)

class D10_DC21(D10, NamedTuple('DC21', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC20(D10, NamedTuple('DC20', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC20({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC20) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC23(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D11_DC22)

class D11_DC23(D11, NamedTuple('DC23', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC22(D11, NamedTuple('DC22', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC22({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC22) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC26(_dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D12_DC25)

class D12_DC26(D12, NamedTuple('DC26', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC25(D12, NamedTuple('DC25', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC25({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC25) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D13_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)

class D13_DC28(D13, NamedTuple('DC28', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({self.cf40.VerbatimString(True)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC27(D13, NamedTuple('DC27', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC27({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC27) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC31(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC31(D14, NamedTuple('DC31', [('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC30(D14, NamedTuple('DC30', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC32(D14, NamedTuple('DC32', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC34()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D15_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)

class D15_DC34(D15, NamedTuple('DC34', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC35(D15, NamedTuple('DC35', [('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC36(D15, NamedTuple('DC36', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC33(D15, NamedTuple('DC33', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC39(int(0), False, _dafny.Map({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)

class D16_DC39(D16, NamedTuple('DC39', [('cf62', Any), ('cf63', Any), ('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC38(D16, NamedTuple('DC38', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC40(D17, NamedTuple('DC40', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D18_DC41)

class D18_DC41(D18, NamedTuple('DC41', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC43(False, _dafny.CodePoint('D'), False, D6.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D19_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D19_DC42)

class D19_DC43(D19, NamedTuple('DC43', [('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC43({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {self.cf73.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC43) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC42(D19, NamedTuple('DC42', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC42({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC42) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC45(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D20_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D20_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D20_DC44)

class D20_DC45(D20, NamedTuple('DC45', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC45({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC45) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC46(D20, NamedTuple('DC46', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC46'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC46)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC44(D20, NamedTuple('DC44', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC44({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC44) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC48(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D21_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D21_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D21_DC50)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D21_DC47)

class D21_DC48(D21, NamedTuple('DC48', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC48({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC48) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC49(D21, NamedTuple('DC49', [('cf79', Any), ('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC49({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC49) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC50(D21, NamedTuple('DC50', [('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC50({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC50) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC47(D21, NamedTuple('DC47', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC47({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC47) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC52()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D22_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D22_DC51)

class D22_DC52(D22, NamedTuple('DC52', [])):
    def __dafnystr__(self) -> str:
        return f'D22.DC52'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC52)
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC53(D22, NamedTuple('DC53', [('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {self.cf89.VerbatimString(True)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC51(D22, NamedTuple('DC51', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC51({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC51) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC55(_dafny.Seq({}), _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D23_DC56)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D23_DC54)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)

class D23_DC55(D23, NamedTuple('DC55', [('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC56(D23, NamedTuple('DC56', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC56'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC56)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC54(D23, NamedTuple('DC54', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC54({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC54) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC57(D23, NamedTuple('DC57', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC59(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D24_DC59)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D24_DC58)

class D24_DC59(D24, NamedTuple('DC59', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC59({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC59) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC60(D24, NamedTuple('DC60', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC58(D24, NamedTuple('DC58', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC58({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC58) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC62(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D25_DC62)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D25_DC61)

class D25_DC62(D25, NamedTuple('DC62', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC62({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC62) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC61(D25, NamedTuple('DC61', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC61({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC61) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D26_DC63)

class D26_DC63(D26, NamedTuple('DC63', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC63({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC63) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m0(self, p0, globalState):
        pass

    def m1(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f5: _dafny.Array = _dafny.Array(None, 0)
        self.f6: bool = False
        self.f7: bool = False
        self._f0: bool = False
        self._f2: int = int(0)
        self._f3: int = int(0)
        self._f4: bool = False
        self._f8: bool = False
        self._f9: int = int(0)
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10

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
        self._f20: _dafny.Seq = _dafny.Seq({})
        self._f21: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f20, f21):
        (self)._f20 = f20
        (self)._f21 = f21

    def fm21(self, p0, p1, globalState):
        if (_dafny.MultiSet([False, True])) != (_dafny.MultiSet([True])):
            return (not(not(True))) and (False)
        elif True:
            return False

    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C1:
    def  __init__(self):
        self._f18: bool = False
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f18, f19):
        (self)._f18 = f18
        (self)._f19 = f19

    def fm13(self, p0, p1, p2, p3, globalState):
        return (self).f18

    def fm14(self, p0, globalState):
        return (self).f18

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_346_v0_: _dafny.MultiSet
        d_346_v0_ = _dafny.MultiSet([(self).f18, (self).f18, (self).f18])
        d_347_v1_: _dafny.Set
        d_347_v1_ = _dafny.Set({default__.fm15(p2, globalState)})
        d_348_v3_: _dafny.Set
        def iife48_():
            coll34_ = _dafny.Set()
            compr_34_: _dafny.Seq
            for compr_34_ in (default__.fm16(p2, len(p3), globalState)).keys.Elements:
                d_349_v2_: _dafny.Seq = compr_34_
                if (d_349_v2_) in (default__.fm16(p2, len(p3), globalState)):
                    coll34_ = coll34_.union(_dafny.Set([d_349_v2_]))
            return _dafny.Set(coll34_)
        d_348_v3_ = _dafny.Set({d_347_v1_, d_347_v1_, iife48_()
        })
        d_350_v4_: _dafny.Seq
        d_350_v4_ = _dafny.SeqWithoutIsStrInference([p1])
        d_351_v5_: _dafny.Map
        d_351_v5_ = _dafny.Map({(self).f19: p1})
        d_352_v6_: _dafny.Map
        d_352_v6_ = _dafny.Map({d_351_v5_: (self).f19})
        d_353_v8_: D2
        d_353_v8_ = D2_DC5()
        d_354_v9_: _dafny.Map
        d_354_v9_ = _dafny.Map({p2: (0) - (p1)})
        d_355_v10_: _dafny.Seq
        d_355_v10_ = _dafny.SeqWithoutIsStrInference([(default__.fm17(d_352_v6_, p1, D2_DC5(), (self).f18, globalState)).set(default__.safeIndex(len(p0), len(default__.fm17(d_352_v6_, p1, D2_DC5(), (self).f18, globalState))), len((d_354_v9_).set(-171, p1))), d_350_v4_, default__.fm17(d_352_v6_, 415, d_353_v8_, (self).f19, globalState)])
        d_356_v11_: _dafny.Seq
        def iife49_():
            coll35_ = _dafny.Map()
            compr_35_: int
            for compr_35_ in _dafny.IntegerRange(-770, -776):
                d_357_v7_: int = compr_35_
                if ((-770) <= (d_357_v7_)) and ((d_357_v7_) < (-776)):
                    coll35_[(d_357_v7_) * (p2)] = (self).f18
            return _dafny.Map(coll35_)
        d_356_v11_ = _dafny.SeqWithoutIsStrInference([d_350_v4_, default__.fm17(d_352_v6_, len(iife49_()
        ), d_353_v8_, (self).f18, globalState), d_350_v4_, (d_355_v10_)[default__.safeIndex(p2, len(d_355_v10_))]])
        d_358_v12_: D2
        out15_: D2
        out15_ = (self).m7(d_346_v0_, d_348_v3_, (d_356_v11_)[default__.safeIndex(default__.fm2(p2, 450, globalState), len(d_356_v11_))], globalState)
        d_358_v12_ = out15_
        d_359_v13_: _dafny.Seq
        d_359_v13_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f18, (p2) < (p2), (self).f18])
        if (d_359_v13_)[default__.safeIndex(p2, len(d_359_v13_))]:
            d_360_v14_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.Map({}), 3)
            d_360_v14_ = nw48_
            d_360_v14_ = d_360_v14_
            d_361_v15_: _dafny.Array
            nw49_ = _dafny.Array(_dafny.Map({}), 13)
            d_361_v15_ = nw49_
            d_362_v16_: str
            d_362_v16_ = _dafny.CodePoint('k')
            d_363_v17_: _dafny.Array
            nw50_ = _dafny.Array(False, 14)
            d_363_v17_ = nw50_
            d_364_v18_: _dafny.Set
            d_364_v18_ = _dafny.Set({d_363_v17_})
            d_365_v19_: _dafny.Array
            nw51_ = _dafny.Array(None, 21)
            nw51_[int(0)] = default__.fm1(globalState)
            nw51_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_366_i0_ in range(default__.abs(149))])
            nw51_[int(2)] = p0
            nw51_[int(3)] = (p0).set(default__.safeIndex(p2, len(p0)), d_362_v16_)
            nw51_[int(4)] = p0
            nw51_[int(5)] = p0
            nw51_[int(6)] = p0
            nw51_[int(7)] = p0
            nw51_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpilsnhk"))
            nw51_[int(9)] = p0
            nw51_[int(10)] = p0
            nw51_[int(11)] = p0
            nw51_[int(12)] = p0
            nw51_[int(13)] = p0
            nw51_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gn"))
            nw51_[int(15)] = p0
            nw51_[int(16)] = p0
            nw51_[int(17)] = (p0).set(default__.safeIndex((d_350_v4_)[default__.safeIndex(len(d_364_v18_), len(d_350_v4_))], len(p0)), _dafny.CodePoint('a'))
            nw51_[int(18)] = p0
            nw51_[int(19)] = p0
            nw51_[int(20)] = p0
            d_365_v19_ = nw51_
            d_367_v20_: _dafny.Map
            d_367_v20_ = _dafny.Map({d_361_v15_: d_365_v19_})
            d_367_v20_ = (d_367_v20_).set(d_361_v15_, (d_365_v19_ if (d_359_v13_)[default__.safeIndex(len(d_354_v9_), len(d_359_v13_))] else d_365_v19_))
            d_368_v21_: _dafny.Array
            nw52_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_368_v21_ = nw52_
            d_369_v22_: _dafny.Set
            d_369_v22_ = _dafny.Set({p1})
            d_370_v23_: _dafny.Array
            nw53_ = _dafny.Array(None, 3)
            nw53_[int(0)] = d_369_v22_
            nw53_[int(1)] = d_369_v22_
            nw53_[int(2)] = default__.fm18(len(default__.fm19(True, (self).f18, globalState)), globalState)
            d_370_v23_ = nw53_
            index42_ = default__.safeIndex(337, (d_368_v21_).length(0))
            (d_368_v21_)[index42_] = d_370_v23_
            index43_ = default__.safeIndex(337, (d_368_v21_).length(0))
            nw54_ = _dafny.Array(None, 3)
            nw54_[int(0)] = d_369_v22_
            nw54_[int(1)] = d_369_v22_
            nw54_[int(2)] = d_369_v22_
            (d_368_v21_)[index43_] = nw54_
            d_371_v24_: _dafny.Set
            d_371_v24_ = _dafny.Set({(self).f19})
            d_372_v25_: _dafny.Map
            d_372_v25_ = _dafny.Map({d_371_v24_: (self).f19})
            (globalState).f1 = len(d_372_v25_)
            (globalState).f7 = (self).f18
        elif True:
            d_373_v26_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_373_v26_ = nw55_
            d_374_v27_: _dafny.Array
            def lambda9_(d_375_i1_):
                return False

            init4_ = lambda9_
            nw56_ = _dafny.Array(None, 8)
            for i0_4_ in range(nw56_.length(0)):
                nw56_[i0_4_] = init4_(i0_4_)
            d_374_v27_ = nw56_
            index44_ = default__.safeIndex(759, (d_373_v26_).length(0))
            (d_373_v26_)[index44_] = d_374_v27_
            index45_ = default__.safeIndex(759, (d_373_v26_).length(0))
            (d_373_v26_)[index45_] = d_374_v27_
            (globalState).f6 = default__.fm0(globalState)
            d_376_v28_: str
            d_376_v28_ = _dafny.CodePoint('f')
            d_377_v29_: _dafny.Set
            d_377_v29_ = _dafny.Set({p2})
            d_378_v30_: _dafny.Array
            nw57_ = _dafny.Array(None, 5)
            nw57_[int(0)] = d_376_v28_
            nw57_[int(1)] = d_376_v28_
            nw57_[int(2)] = d_376_v28_
            nw57_[int(3)] = d_376_v28_
            nw57_[int(4)] = d_376_v28_
            d_378_v30_ = nw57_
            index46_ = default__.safeIndex(830, (d_378_v30_).length(0))
            (d_378_v30_)[index46_] = d_376_v28_
            index47_ = default__.safeIndex(510, (d_378_v30_).length(0))
            (d_378_v30_)[index47_] = d_376_v28_
            d_379_v31_: _dafny.MultiSet
            d_379_v31_ = _dafny.MultiSet([D2_DC4(len(p0), len(p0), p2)])
            d_380_v32_: _dafny.Map
            d_380_v32_ = _dafny.Map({d_379_v31_: d_377_v29_})
            index48_ = default__.safeIndex(830, (d_378_v30_).length(0))
            index49_ = default__.safeIndex(510, (d_378_v30_).length(0))
            rhs61_ = d_376_v28_
            rhs62_ = ((d_380_v32_)[_dafny.MultiSet([d_358_v12_, d_358_v12_, d_358_v12_])] if (_dafny.MultiSet([d_358_v12_, d_358_v12_, d_358_v12_])) in (d_380_v32_) else d_377_v29_)
            rhs63_ = default__.fm20(d_359_v13_, (p1) - (936), (self).f18, globalState)
            rhs64_ = d_376_v28_
            lhs54_ = d_378_v30_
            lhs55_ = default__.safeIndex(830, (d_378_v30_).length(0))
            lhs56_ = d_378_v30_
            lhs57_ = default__.safeIndex(510, (d_378_v30_).length(0))
            d_376_v28_ = rhs61_
            d_377_v29_ = rhs62_
            lhs54_[lhs55_] = rhs63_
            lhs56_[lhs57_] = rhs64_
            d_374_v27_ = d_374_v27_
            d_381_v33_: _dafny.Seq
            d_381_v33_ = d_359_v13_
            d_382_v34_: C0
            nw58_ = C0()
            nw58_.ctor__((d_381_v33_), (default__.fm2(p1, p2, globalState)) * (353))
            d_382_v34_ = nw58_
        (globalState).f6 = (self).f19
        d_383_v35_: D0
        d_383_v35_ = D0_DC0(p1)
        source7_ = d_383_v35_
        if source7_.is_DC1:
            d_384___mcc_h0_ = source7_.cf1
            d_385___mcc_h1_ = source7_.cf2
            d_386___mcc_h2_ = source7_.cf3
            d_387_cf3_ = d_386___mcc_h2_
            d_388_cf2_ = d_385___mcc_h1_
            d_389_cf1_ = d_384___mcc_h0_
            (globalState).f7 = (d_359_v13_)[default__.safeIndex(-413, len(d_359_v13_))]
            d_390_v36_: _dafny.MultiSet
            d_390_v36_ = _dafny.MultiSet([d_388_cf2_, (0) - (p1)])
            d_387_cf3_ = (((d_390_v36_)[(0) - (len(p3))] if ((0) - (len(p3))) in (d_390_v36_) else p2)) < (d_388_cf2_)
            d_391_v37_: C0
            nw59_ = C0()
            nw59_.ctor__((d_359_v13_) + (d_359_v13_), default__.safeDivisionInt(906, d_388_cf2_))
            d_391_v37_ = nw59_
            d_392_v38_: _dafny.Map
            d_392_v38_ = _dafny.Map({d_353_v8_: len((default__.fm22(p0, True, globalState)) | (_dafny.Map({p1: (self).f19})))})
            d_393_v39_: _dafny.Seq
            d_393_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jr"))
            d_394_v40_: _dafny.Array
            def lambda10_(d_395_v37_):
                def lambda11_(d_396_i2_):
                    return (d_396_i2_) - ((d_395_v37_).f21)

                return lambda11_

            init5_ = lambda10_(d_391_v37_)
            nw60_ = _dafny.Array(None, 5)
            for i0_5_ in range(nw60_.length(0)):
                nw60_[i0_5_] = init5_(i0_5_)
            d_394_v40_ = nw60_
            index50_ = default__.safeIndex(368, (d_394_v40_).length(0))
            (d_394_v40_)[index50_] = (p2) + (d_388_cf2_)
            d_397_v41_: str
            d_397_v41_ = _dafny.CodePoint('v')
            index51_ = default__.safeIndex(368, (d_394_v40_).length(0))
            rhs65_ = d_392_v38_
            rhs66_ = ((default__.fm1(globalState)) + (d_393_v39_)).set(default__.safeIndex((0) - ((d_388_cf2_ if (self).f18 else (0) - ((d_358_v12_).cf8))), len((default__.fm1(globalState)) + (d_393_v39_))), d_397_v41_)
            rhs67_ = d_358_v12_
            rhs68_ = p1
            rhs69_ = p1
            lhs58_ = d_394_v40_
            lhs59_ = default__.safeIndex(368, (d_394_v40_).length(0))
            d_392_v38_ = rhs65_
            d_393_v39_ = rhs66_
            d_358_v12_ = rhs67_
            lhs58_[lhs59_] = rhs68_
            r0 = rhs69_
        elif True:
            d_398___mcc_h3_ = source7_.cf0
            d_399_cf0_ = d_398___mcc_h3_
            d_400_v42_: _dafny.Array
            def lambda12_(d_401_v0_, d_402_p1_):
                def lambda13_(d_403_i3_):
                    return (d_403_i3_) - (((d_401_v0_)[(self).f19] if ((self).f19) in (d_401_v0_) else d_402_p1_))

                return lambda13_

            init6_ = lambda12_(d_346_v0_, p1)
            nw61_ = _dafny.Array(None, 29)
            for i0_6_ in range(nw61_.length(0)):
                nw61_[i0_6_] = init6_(i0_6_)
            d_400_v42_ = nw61_
            d_404_v43_: _dafny.MultiSet
            d_404_v43_ = _dafny.MultiSet([d_400_v42_])
            d_405_v44_: _dafny.Seq
            d_405_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_406_i4_ in range(default__.abs(973))])])
            d_407_v45_: _dafny.Map
            d_407_v45_ = _dafny.Map({(d_404_v43_).intersection(d_404_v43_): d_405_v44_})
            d_407_v45_ = (d_407_v45_).set(d_404_v43_, d_405_v44_)
            (globalState).f7 = (((d_350_v4_)[default__.safeIndex(len(d_351_v5_), len(d_350_v4_))]) + (d_399_cf0_)) <= (d_399_cf0_)
            d_408_v46_: _dafny.MultiSet
            d_408_v46_ = _dafny.MultiSet([d_350_v4_])
            d_409_v47_: _dafny.Array
            nw62_ = _dafny.Array(False, 17)
            d_409_v47_ = nw62_
            d_410_v48_: _dafny.MultiSet
            d_410_v48_ = _dafny.MultiSet([d_409_v47_, d_409_v47_, d_409_v47_])
            d_411_v49_: _dafny.Map
            d_411_v49_ = _dafny.Map({(self).f19: (self).f18})
            d_412_v50_: _dafny.Array
            nw63_ = _dafny.Array(None, 24)
            nw63_[int(0)] = (self).f19
            nw63_[int(1)] = (d_408_v46_).isdisjoint(d_408_v46_)
            nw63_[int(2)] = (self).f18
            nw63_[int(3)] = not ((self).f18) or ((self).f19)
            nw63_[int(4)] = not((self).f18)
            nw63_[int(5)] = (self).f18
            nw63_[int(6)] = (self).f19
            nw63_[int(7)] = (self).f18
            nw63_[int(8)] = (self).f18
            nw63_[int(9)] = (p3) == (p3)
            nw63_[int(10)] = (self).f18
            nw63_[int(11)] = (self).f19
            nw63_[int(12)] = False
            nw63_[int(13)] = (self).f19
            nw63_[int(14)] = (d_410_v48_).ispropersubset(_dafny.MultiSet([d_409_v47_]))
            nw63_[int(15)] = (self).f18
            nw63_[int(16)] = (self).f19
            nw63_[int(17)] = (default__.fm1(globalState)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kppg")))
            nw63_[int(18)] = (self).f18
            nw63_[int(19)] = (self).f19
            nw63_[int(20)] = (self).f18
            nw63_[int(21)] = (self).f18
            nw63_[int(22)] = not((len(d_411_v49_)) >= (p1))
            nw63_[int(23)] = (self).f18
            d_412_v50_ = nw63_
            index52_ = default__.safeIndex(839, (d_409_v47_).length(0))
            (d_409_v47_)[index52_] = (self).f19
            d_413_v51_: _dafny.Map
            d_413_v51_ = _dafny.Map({(self).f18: d_346_v0_})
            d_414_v52_: _dafny.Map
            d_414_v52_ = _dafny.Map({len(default__.fm18(p2, globalState)): _dafny.MultiSet(d_359_v13_)})
            d_415_v53_: _dafny.MultiSet
            d_415_v53_ = ((d_413_v51_)[(self).f19] if ((self).f19) in (d_413_v51_) else d_346_v0_)
            d_416_v54_: _dafny.Array
            nw64_ = _dafny.Array(None, 28)
            nw64_[int(0)] = d_346_v0_
            nw64_[int(1)] = d_346_v0_
            nw64_[int(2)] = (d_346_v0_ if (self).f19 else _dafny.MultiSet(d_359_v13_))
            nw64_[int(3)] = (d_346_v0_) - (_dafny.MultiSet([(self).f18]))
            nw64_[int(4)] = d_346_v0_
            nw64_[int(5)] = (d_346_v0_) | (d_346_v0_)
            nw64_[int(6)] = d_346_v0_
            nw64_[int(7)] = ((d_413_v51_)[not((self).f18)] if (not((self).f18)) in (d_413_v51_) else d_346_v0_)
            nw64_[int(8)] = _dafny.MultiSet([(self).f19])
            nw64_[int(9)] = d_346_v0_
            nw64_[int(10)] = _dafny.MultiSet([(self).f19, not((self).f19), (self).f18, (self).f19])
            nw64_[int(11)] = d_346_v0_
            nw64_[int(12)] = (_dafny.MultiSet([(self).f19, False]) if False else d_346_v0_)
            nw64_[int(13)] = d_346_v0_
            nw64_[int(14)] = ((d_414_v52_)[d_399_cf0_] if (d_399_cf0_) in (d_414_v52_) else _dafny.MultiSet(d_359_v13_))
            nw64_[int(15)] = d_346_v0_
            nw64_[int(16)] = d_346_v0_
            nw64_[int(17)] = (d_346_v0_) - ((d_415_v53_))
            nw64_[int(18)] = d_346_v0_
            nw64_[int(19)] = (d_346_v0_) | ((_dafny.MultiSet([(self).f18])).set((self).f19, default__.abs(727)))
            nw64_[int(20)] = d_346_v0_
            nw64_[int(21)] = d_346_v0_
            nw64_[int(22)] = (_dafny.MultiSet([(self).f19, (self).f19])).intersection((d_346_v0_).set(not((self).f18), default__.abs(p2)))
            nw64_[int(23)] = d_346_v0_
            nw64_[int(24)] = d_346_v0_
            nw64_[int(25)] = d_346_v0_
            nw64_[int(26)] = d_346_v0_
            nw64_[int(27)] = d_346_v0_
            d_416_v54_ = nw64_
            index53_ = default__.safeIndex(635, (d_416_v54_).length(0))
            (d_416_v54_)[index53_] = (d_346_v0_).set((self).f18, default__.abs(p2))
            d_417_v55_: D0
            d_417_v55_ = D0_DC1(not((self).f18), p2, (self).f19)
            d_418_v56_: str
            d_418_v56_ = _dafny.CodePoint('t')
            d_419_v57_: _dafny.Array
            nw65_ = _dafny.Array(None, 11)
            nw65_[int(0)] = (d_418_v56_ if (d_417_v55_).cf1 else default__.fm20(d_359_v13_, d_399_cf0_, (self).f19, globalState))
            nw65_[int(1)] = _dafny.CodePoint('t')
            nw65_[int(2)] = _dafny.CodePoint('b')
            nw65_[int(3)] = _dafny.CodePoint('x')
            nw65_[int(4)] = _dafny.CodePoint('r')
            nw65_[int(5)] = d_418_v56_
            nw65_[int(6)] = _dafny.CodePoint('q')
            nw65_[int(7)] = _dafny.CodePoint('e')
            nw65_[int(8)] = d_418_v56_
            nw65_[int(9)] = d_418_v56_
            nw65_[int(10)] = d_418_v56_
            d_419_v57_ = nw65_
            index54_ = default__.safeIndex(754, (d_419_v57_).length(0))
            (d_419_v57_)[index54_] = d_418_v56_
            d_420_v58_: _dafny.Map
            d_420_v58_ = _dafny.Map({(self).f19: p0})
            index55_ = default__.safeIndex(839, (d_409_v47_).length(0))
            index56_ = default__.safeIndex(635, (d_416_v54_).length(0))
            index57_ = default__.safeIndex(754, (d_419_v57_).length(0))
            rhs70_ = d_399_cf0_
            rhs71_ = not((705) >= ((d_399_cf0_) + (d_399_cf0_)))
            rhs72_ = ((p0) + (p0)) != (((d_420_v58_)[(self).f18] if ((self).f18) in (d_420_v58_) else _dafny.SeqWithoutIsStrInference([d_418_v56_ for d_421_i5_ in range(default__.abs(91))])))
            rhs73_ = d_346_v0_
            rhs74_ = d_418_v56_
            lhs60_ = globalState
            lhs61_ = d_409_v47_
            lhs62_ = default__.safeIndex(839, (d_409_v47_).length(0))
            lhs63_ = globalState
            lhs64_ = d_416_v54_
            lhs65_ = default__.safeIndex(635, (d_416_v54_).length(0))
            lhs66_ = d_419_v57_
            lhs67_ = default__.safeIndex(754, (d_419_v57_).length(0))
            lhs60_.f1 = rhs70_
            lhs61_[lhs62_] = rhs71_
            lhs63_.f6 = rhs72_
            lhs64_[lhs65_] = rhs73_
            lhs66_[lhs67_] = rhs74_
            (globalState).f6 = (d_409_v47_)[default__.safeIndex(839, (d_409_v47_).length(0))]
        (globalState).f1 = p2
        d_422_v59_: D0
        d_422_v59_ = D0_DC1(((self).f18) == ((self).f19), p1, not ((self).f18) or ((self).f18))
        pat_let_tv21_ = p1
        def iife50_(_pat_let7_0):
            def iife51_(d_423_dt__update__tmp_h0_):
                def iife52_(_pat_let8_0):
                    def iife53_(d_424_dt__update_hcf1_h0_):
                        def iife54_(_pat_let9_0):
                            def iife55_(d_425_dt__update_hcf2_h0_):
                                return D0_DC1(d_424_dt__update_hcf1_h0_, d_425_dt__update_hcf2_h0_, (d_423_dt__update__tmp_h0_).cf3)
                            return iife55_(_pat_let9_0)
                        return iife54_(default__.safeModuloInt(473, pat_let_tv21_))
                    return iife53_(_pat_let8_0)
                return iife52_((self).f18)
            return iife51_(_pat_let7_0)
        d_422_v59_ = iife50_(d_422_v59_)
        r0 = p1
        d_426_v60_: D6
        d_426_v60_ = D6_DC9(p3)
        pat_let_tv22_ = p3
        d_427_v61_: _dafny.Array
        nw66_ = _dafny.Array(None, 10)
        nw66_[int(0)] = p3
        nw66_[int(1)] = (d_426_v60_).cf12
        nw66_[int(2)] = (p3) | (_dafny.Map({p1: (self).f19}))
        nw66_[int(3)] = p3
        nw66_[int(4)] = p3
        nw66_[int(5)] = (p3) | (p3)
        nw66_[int(6)] = p3
        def iife56_(_pat_let10_0):
            def iife57_(d_428_dt__update__tmp_h1_):
                def iife58_(_pat_let11_0):
                    def iife59_(d_429_dt__update_hcf12_h0_):
                        return D6_DC9(d_429_dt__update_hcf12_h0_)
                    return iife59_(_pat_let11_0)
                return iife58_(pat_let_tv22_)
            return iife57_(_pat_let10_0)
        nw66_[int(7)] = ((iife56_(d_426_v60_)).cf12) | ((p3).set(p2, (self).f19))
        nw66_[int(8)] = (p3) | (p3)
        nw66_[int(9)] = p3
        d_427_v61_ = nw66_
        r1 = d_427_v61_
        return r0, r1

    def m7(self, p0, p1, p2, globalState):
        r0: D2 = D2.default()()
        d_430_v0_: _dafny.Seq
        d_430_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpix"))
        d_431_i0_: int
        d_431_i0_ = 0
        with _dafny.label("4"):
            while (default__.fm1(globalState)) <= ((d_430_v0_) + (d_430_v0_)):
                with _dafny.c_label("4"):
                    if (d_431_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_431_i0_ = (d_431_i0_) + (1)
                    d_432_v1_: int
                    d_432_v1_ = 239
                    (globalState).f6 = ((self).f18 if ((self).f19) == ((self).f19) else (d_432_v1_) < (d_432_v1_))
                    d_433_v2_: str
                    d_433_v2_ = _dafny.CodePoint('s')
                    d_433_v2_ = d_433_v2_
                    d_434_v3_: _dafny.Map
                    d_434_v3_ = _dafny.Map({False: d_433_v2_})
                    d_433_v2_ = ((d_434_v3_)[(self).f19] if ((self).f19) in (d_434_v3_) else d_433_v2_)
                    d_435_v4_: D0
                    d_435_v4_ = D0_DC1((self).f19, d_432_v1_, ((self).f18) in (p0))
                    d_436_v5_: _dafny.Array
                    def lambda14_(d_437_v1_):
                        def lambda15_(d_438_i1_):
                            return default__.safeDivisionInt(d_438_i1_, d_437_v1_)

                        return lambda15_

                    init7_ = lambda14_(d_432_v1_)
                    nw67_ = _dafny.Array(None, 24)
                    for i0_7_ in range(nw67_.length(0)):
                        nw67_[i0_7_] = init7_(i0_7_)
                    d_436_v5_ = nw67_
                    index58_ = default__.safeIndex(750, (d_436_v5_).length(0))
                    (d_436_v5_)[index58_] = d_432_v1_
                    d_439_v6_: _dafny.MultiSet
                    d_439_v6_ = _dafny.MultiSet([794])
                    d_440_v7_: _dafny.Map
                    d_440_v7_ = _dafny.Map({d_439_v6_: d_432_v1_})
                    d_441_v8_: _dafny.Seq
                    d_441_v8_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                    d_442_v9_: _dafny.Seq
                    d_442_v9_ = _dafny.SeqWithoutIsStrInference([d_430_v0_, d_430_v0_, d_430_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efknn")), _dafny.SeqWithoutIsStrInference([d_433_v2_ for d_443_i2_ in range(default__.abs(735))])])
                    index59_ = default__.safeIndex(750, (d_436_v5_).length(0))
                    rhs75_ = D0_DC1((self).f19, default__.safeDivisionInt(len(default__.fm23(d_432_v1_, d_432_v1_, globalState)), d_432_v1_), (self).f19)
                    rhs76_ = (0) - (d_432_v1_)
                    rhs77_ = (0) - (((d_440_v7_)[d_439_v6_] if (d_439_v6_) in (d_440_v7_) else len(d_441_v8_)))
                    rhs78_ = len((d_430_v0_ if (self).f19 else (d_442_v9_)[default__.safeIndex(326, len(d_442_v9_))]))
                    lhs68_ = globalState
                    lhs69_ = d_436_v5_
                    lhs70_ = default__.safeIndex(750, (d_436_v5_).length(0))
                    d_435_v4_ = rhs75_
                    lhs68_.f1 = rhs76_
                    lhs69_[lhs70_] = rhs77_
                    d_432_v1_ = rhs78_
                    pass
            pass
        d_444_v10_: int
        d_444_v10_ = 907
        d_445_v12_: _dafny.Array
        nw68_ = _dafny.Array(None, 5)
        nw68_[int(0)] = default__.fm2(d_444_v10_, d_444_v10_, globalState)
        nw68_[int(1)] = d_444_v10_
        nw68_[int(2)] = len(d_430_v0_)
        def iife60_():
            coll36_ = _dafny.Map()
            compr_36_: int
            for compr_36_ in _dafny.IntegerRange(180, 711):
                d_446_v11_: int = compr_36_
                if ((180) <= (d_446_v11_)) and ((d_446_v11_) < (711)):
                    coll36_[(d_446_v11_) + (len(d_430_v0_))] = len(_dafny.Set({d_444_v10_}))
            return _dafny.Map(coll36_)
        nw68_[int(3)] = len(iife60_()
        )
        nw68_[int(4)] = d_444_v10_
        d_445_v12_ = nw68_
        d_447_v13_: _dafny.Set
        d_447_v13_ = _dafny.Set({d_445_v12_})
        d_447_v13_ = d_447_v13_
        d_448_v14_: _dafny.Map
        d_448_v14_ = _dafny.Map({(self).f19: d_444_v10_})
        d_449_v15_: _dafny.Set
        d_449_v15_ = _dafny.Set({not((self).f19), (self).f18})
        d_444_v10_ = default__.safeModuloInt(((d_448_v14_)[(self).f18] if ((self).f18) in (d_448_v14_) else d_444_v10_), default__.safeModuloInt(len(d_449_v15_), d_444_v10_))
        (globalState).f6 = (self).f19
        d_450_i3_: int
        d_450_i3_ = 0
        with _dafny.label("5"):
            while (self).f18:
                with _dafny.c_label("5"):
                    if (d_450_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_450_i3_ = (d_450_i3_) + (1)
                    d_451_v16_: _dafny.Set
                    d_451_v16_ = _dafny.Set({789})
                    d_452_v17_: D7
                    d_452_v17_ = D7_DC12(p2)
                    d_444_v10_ = default__.fm2(len((_dafny.Set({d_444_v10_, d_444_v10_, d_444_v10_, d_444_v10_, d_444_v10_})) | (d_451_v16_)), (len((d_452_v17_).cf17)) - (d_444_v10_), globalState)
                    d_453_v18_: D2
                    d_453_v18_ = D2_DC4(d_444_v10_, d_444_v10_, d_444_v10_)
                    (globalState).f1 = default__.fm2(default__.fm2(d_444_v10_, d_444_v10_, globalState), (d_444_v10_) * ((d_453_v18_).cf6), globalState)
                    d_454_v19_: _dafny.Array
                    def lambda16_(d_455_v16_):
                        def lambda17_(d_456_i4_):
                            return d_455_v16_

                        return lambda17_

                    init8_ = lambda16_(d_451_v16_)
                    nw69_ = _dafny.Array(None, 3)
                    for i0_8_ in range(nw69_.length(0)):
                        nw69_[i0_8_] = init8_(i0_8_)
                    d_454_v19_ = nw69_
                    index60_ = default__.safeIndex(400, (d_454_v19_).length(0))
                    (d_454_v19_)[index60_] = d_451_v16_
                    d_457_v20_: _dafny.Map
                    d_457_v20_ = _dafny.Map({(self).f19: default__.fm18(519, globalState)})
                    index61_ = default__.safeIndex(400, (d_454_v19_).length(0))
                    (d_454_v19_)[index61_] = (((d_457_v20_)[(self).f18] if ((self).f18) in (d_457_v20_) else d_451_v16_)) | (d_451_v16_)
                    (globalState).f6 = (self).f19
                    pass
            pass
        d_458_v21_: _dafny.Array
        nw70_ = _dafny.Array(_dafny.Map({}), 1)
        d_458_v21_ = nw70_
        d_458_v21_ = d_458_v21_
        d_459_v22_: D0
        d_459_v22_ = D0_DC1((self).f19, d_444_v10_, False)
        pat_let_tv23_ = d_444_v10_
        pat_let_tv24_ = d_444_v10_
        pat_let_tv25_ = d_444_v10_
        def lambda18_(source8_):
            d_460___mcc_h0_ = source8_
            d_461_cf11_ = d_460___mcc_h0_
            return D2_DC4(pat_let_tv23_, pat_let_tv24_, pat_let_tv25_)

        r0 = lambda18_(default__.fm24((d_459_v22_).cf3, globalState))
        return r0

    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19

class C2(T0):
    def  __init__(self):
        self._f11: int = int(0)
        self._f24: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f24, f11):
        (self)._f24 = f24
        (self)._f11 = f11

    def fm3(self, p0, p1, p2, globalState):
        return ((self).f11) + ((self).f11)

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(self).f24 for d_462_i0_ in range(default__.abs(337))])) + (_dafny.SeqWithoutIsStrInference([(self).f24 for d_463_i1_ in range(default__.abs(428))]))) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pykmo"))) + (_dafny.SeqWithoutIsStrInference([(self).f24 for d_464_i2_ in range(default__.abs(-42))])))

    def m8(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_465_v0_: bool
        d_465_v0_ = False
        d_466_v1_: _dafny.Map
        d_466_v1_ = _dafny.Map({-775: (self).fm3((self).f11, not(d_465_v0_), d_465_v0_, globalState)})
        d_467_v2_: D8
        d_467_v2_ = D8_DC16(True, ((d_466_v1_)[(self).f11] if ((self).f11) in (d_466_v1_) else (self).f11))
        (globalState).f1 = ((d_467_v2_).cf24) - (((self).f11) + ((self).f11))
        d_468_v3_: _dafny.Seq
        d_468_v3_ = _dafny.SeqWithoutIsStrInference([d_465_v0_, d_465_v0_])
        d_469_v4_: C0
        nw71_ = C0()
        nw71_.ctor__((d_468_v3_) + (d_468_v3_), ((self).f11) * (687))
        d_469_v4_ = nw71_
        d_470_v5_: _dafny.Map
        d_470_v5_ = _dafny.Map({(0) - ((d_469_v4_).f21): not(d_465_v0_)})
        d_471_v6_: D6
        d_471_v6_ = D6_DC9((d_470_v5_) | ((d_470_v5_).set((d_469_v4_).f21, d_465_v0_)))
        source9_ = d_471_v6_
        if source9_.is_DC10:
            d_472___mcc_h0_ = source9_.cf13
            d_473_cf13_ = d_472___mcc_h0_
            d_474_v7_: _dafny.Seq
            d_474_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpp"))
            d_475_v8_: _dafny.Seq
            d_475_v8_ = _dafny.SeqWithoutIsStrInference([len(d_474_v7_)])
            d_476_v9_: D7
            d_476_v9_ = D7_DC12(d_475_v8_)
            d_477_v10_: _dafny.Map
            d_477_v10_ = _dafny.Map({False: d_465_v0_})
            d_478_v11_: _dafny.Map
            d_478_v11_ = _dafny.Map({d_477_v10_: d_465_v0_})
            d_479_v12_: _dafny.Map
            d_479_v12_ = _dafny.Map({len(d_478_v11_): d_475_v8_})
            d_480_v13_: _dafny.Array
            nw72_ = _dafny.Array(None, 7)
            nw72_[int(0)] = d_475_v8_
            nw72_[int(1)] = d_475_v8_
            nw72_[int(2)] = d_475_v8_
            nw72_[int(3)] = (d_476_v9_).cf17
            nw72_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_469_v4_).f21 for d_481_i0_ in range(default__.abs(364))])
            nw72_[int(5)] = d_475_v8_
            nw72_[int(6)] = (d_475_v8_) + (((d_479_v12_)[len((d_469_v4_).f20)] if (len((d_469_v4_).f20)) in (d_479_v12_) else d_475_v8_))
            d_480_v13_ = nw72_
            rhs79_ = ((self).f11) < ((0) - ((0) - ((d_469_v4_).f21)))
            rhs80_ = (d_480_v13_ if d_465_v0_ else d_480_v13_)
            rhs81_ = d_465_v0_
            lhs71_ = globalState
            lhs72_ = globalState
            lhs71_.f6 = rhs79_
            d_480_v13_ = rhs80_
            lhs72_.f6 = rhs81_
            d_482_v14_: _dafny.MultiSet
            d_482_v14_ = _dafny.MultiSet([d_465_v0_])
            if (d_465_v0_) not in (d_482_v14_):
                d_483_v15_: _dafny.Map
                d_483_v15_ = _dafny.Map({d_465_v0_: (d_469_v4_).f21})
                r3 = default__.safeDivisionInt((self).fm3(len((d_469_v4_).f20), d_465_v0_, d_465_v0_, globalState), ((d_483_v15_)[d_465_v0_] if (d_465_v0_) in (d_483_v15_) else (self).f11))
                d_484_v16_: _dafny.Set
                d_484_v16_ = _dafny.Set({(self).f24})
                d_485_v17_: _dafny.Set
                d_485_v17_ = _dafny.Set({d_484_v16_})
                d_486_v18_: _dafny.Map
                d_486_v18_ = _dafny.Map({d_474_v7_: d_485_v17_})
                d_486_v18_ = (d_486_v18_).set(d_474_v7_, default__.fm28(d_465_v0_, -68, globalState))
                d_487_v19_: _dafny.Array
                nw73_ = _dafny.Array(int(0), 20)
                d_487_v19_ = nw73_
                d_488_v20_: _dafny.Array
                nw74_ = _dafny.Array(None, 13)
                nw74_[int(0)] = d_487_v19_
                nw74_[int(1)] = d_487_v19_
                nw74_[int(2)] = d_487_v19_
                nw74_[int(3)] = d_487_v19_
                nw74_[int(4)] = d_487_v19_
                nw74_[int(5)] = d_487_v19_
                nw74_[int(6)] = d_487_v19_
                nw74_[int(7)] = d_487_v19_
                nw74_[int(8)] = d_487_v19_
                nw74_[int(9)] = d_487_v19_
                nw74_[int(10)] = d_487_v19_
                nw74_[int(11)] = d_487_v19_
                nw74_[int(12)] = (d_487_v19_ if False else d_487_v19_)
                d_488_v20_ = nw74_
                index62_ = default__.safeIndex(951, (d_488_v20_).length(0))
                (d_488_v20_)[index62_] = d_487_v19_
                d_489_v21_: _dafny.MultiSet
                d_489_v21_ = _dafny.MultiSet([(self).f11, (d_469_v4_).f21])
                index63_ = default__.safeIndex(951, (d_488_v20_).length(0))
                rhs82_ = (0) - ((d_469_v4_).f21)
                rhs83_ = ((self).f11) + (default__.safeDivisionInt(((d_489_v21_)[(d_469_v4_).f21] if ((d_469_v4_).f21) in (d_489_v21_) else ((d_466_v1_)[(d_469_v4_).f21] if ((d_469_v4_).f21) in (d_466_v1_) else (d_469_v4_).f21)), (self).f11))
                rhs84_ = d_487_v19_
                lhs73_ = globalState
                lhs74_ = globalState
                lhs75_ = d_488_v20_
                lhs76_ = default__.safeIndex(951, (d_488_v20_).length(0))
                lhs73_.f1 = rhs82_
                lhs74_.f1 = rhs83_
                lhs75_[lhs76_] = rhs84_
                (globalState).f7 = d_465_v0_
                d_490_v22_: _dafny.Set
                d_490_v22_ = _dafny.Set({(d_465_v0_ if d_465_v0_ else False)})
                d_491_v23_: str
                d_491_v23_ = _dafny.CodePoint('v')
                rhs85_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.Map({len(d_474_v7_): d_465_v0_}))) - ((0) - (((d_466_v1_)[968] if (968) in (d_466_v1_) else (d_469_v4_).f21))) for d_492_i1_ in range(default__.abs(249))])
                rhs86_ = d_490_v22_
                rhs87_ = ((self).f24 if d_465_v0_ else d_491_v23_)
                rhs88_ = (d_469_v4_).f21
                lhs77_ = globalState
                d_475_v8_ = rhs85_
                d_490_v22_ = rhs86_
                d_491_v23_ = rhs87_
                lhs77_.f1 = rhs88_
            elif True:
                d_470_v5_ = (d_470_v5_).set((default__.fm2((0) - ((d_469_v4_).f21), (0) - ((self).f11), globalState) if d_465_v0_ else -92), d_465_v0_)
                d_493_v24_: _dafny.Set
                d_493_v24_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21 for d_494_i2_ in range(default__.abs(260))]) if d_465_v0_ else d_475_v8_)})
                d_493_v24_ = d_493_v24_
                d_495_v25_: str
                d_495_v25_ = _dafny.CodePoint('i')
                d_496_v26_: _dafny.Set
                d_496_v26_ = _dafny.Set({D9_DC18(_dafny.Map({d_465_v0_: (d_469_v4_).f21}))})
                d_497_v27_: _dafny.Array
                nw75_ = _dafny.Array(None, 9)
                nw75_[int(0)] = d_496_v26_
                nw75_[int(1)] = d_496_v26_
                nw75_[int(2)] = d_496_v26_
                nw75_[int(3)] = d_496_v26_
                nw75_[int(4)] = (d_496_v26_) - (d_496_v26_)
                nw75_[int(5)] = d_496_v26_
                nw75_[int(6)] = default__.fm29(globalState)
                nw75_[int(7)] = d_496_v26_
                nw75_[int(8)] = d_496_v26_
                d_497_v27_ = nw75_
                index64_ = default__.safeIndex(660, (d_497_v27_).length(0))
                (d_497_v27_)[index64_] = d_496_v26_
                index65_ = default__.safeIndex(660, (d_497_v27_).length(0))
                rhs89_ = (self).f24
                rhs90_ = d_496_v26_
                lhs78_ = d_497_v27_
                lhs79_ = default__.safeIndex(660, (d_497_v27_).length(0))
                d_495_v25_ = rhs89_
                lhs78_[lhs79_] = rhs90_
                d_498_v28_: D8
                d_498_v28_ = D8_DC15(d_477_v10_)
                d_499_v29_: D8
                d_499_v29_ = D8_DC17(d_498_v28_)
                d_499_v29_ = d_499_v29_
                d_500_v30_: C0
                nw76_ = C0()
                nw76_.ctor__(_dafny.SeqWithoutIsStrInference([d_465_v0_]), default__.safeDivisionInt((d_469_v4_).f21, (d_469_v4_).f21))
                d_500_v30_ = nw76_
            d_501_v31_: _dafny.Array
            nw77_ = _dafny.Array(int(0), 22)
            d_501_v31_ = nw77_
            index66_ = default__.safeIndex(744, (d_501_v31_).length(0))
            (d_501_v31_)[index66_] = (d_469_v4_).f21
            index67_ = default__.safeIndex(744, (d_501_v31_).length(0))
            (d_501_v31_)[index67_] = default__.safeModuloInt((d_469_v4_).f21, (d_467_v2_).cf24)
            r3 = (default__.safeModuloInt(359, (d_469_v4_).f21)) * ((self).f11)
        elif source9_.is_DC11:
            d_502___mcc_h1_ = source9_.cf14
            d_503___mcc_h2_ = source9_.cf15
            d_504___mcc_h3_ = source9_.cf16
            d_505_cf16_ = d_504___mcc_h3_
            d_506_cf15_ = d_503___mcc_h2_
            d_507_cf14_ = d_502___mcc_h1_
            d_508_v32_: _dafny.Seq
            d_508_v32_ = _dafny.SeqWithoutIsStrInference([(self).f11, d_507_cf14_])
            d_509_v33_: _dafny.Array
            nw78_ = _dafny.Array(None, 8)
            nw78_[int(0)] = (self).f11
            nw78_[int(1)] = (d_469_v4_).f21
            nw78_[int(2)] = 552
            nw78_[int(3)] = len(d_508_v32_)
            nw78_[int(4)] = d_505_cf16_
            nw78_[int(5)] = 683
            nw78_[int(6)] = (0) - ((self).f11)
            nw78_[int(7)] = d_507_cf14_
            d_509_v33_ = nw78_
            d_510_v34_: _dafny.Map
            d_510_v34_ = _dafny.Map({(d_509_v33_ if d_465_v0_ else d_509_v33_): d_465_v0_})
            d_510_v34_ = (d_510_v34_).set(d_509_v33_, (d_469_v4_).fm21(d_465_v0_, 777, globalState))
            d_511_v35_: _dafny.Map
            d_511_v35_ = _dafny.Map({len(d_466_v1_): d_509_v33_})
            d_511_v35_ = d_511_v35_
            d_512_v36_: _dafny.Map
            d_512_v36_ = _dafny.Map({(len(d_508_v32_)) * ((d_469_v4_).f21): (self).f24})
            d_513_v37_: D7
            d_513_v37_ = D7_DC12(d_508_v32_)
            d_514_v39_: _dafny.MultiSet
            d_514_v39_ = _dafny.MultiSet([557, d_505_cf16_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwljmxlt")))])
            d_515_v40_: _dafny.Map
            d_515_v40_ = _dafny.Map({d_465_v0_: d_469_v4_})
            def iife61_():
                coll37_ = _dafny.Set()
                compr_37_: _dafny.Seq
                for compr_37_ in (_dafny.SeqWithoutIsStrInference([(d_513_v37_).cf17, _dafny.SeqWithoutIsStrInference([144])])).Elements:
                    d_516_v38_: _dafny.Seq = compr_37_
                    if (d_516_v38_) in (_dafny.SeqWithoutIsStrInference([(d_513_v37_).cf17, _dafny.SeqWithoutIsStrInference([144])])):
                        coll37_ = coll37_.union(_dafny.Set([d_516_v38_]))
                return _dafny.Set(coll37_)
            rhs91_ = (((_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21])).set(default__.safeIndex((0) - ((d_469_v4_).f21), len(_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21]))), (self).f11)) + (_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21]))) in (iife61_()
            )
            rhs92_ = default__.safeModuloInt((d_469_v4_).f21, ((d_514_v39_).cardinality) + ((0) - ((d_469_v4_).f21)))
            rhs93_ = d_512_v36_
            rhs94_ = (d_465_v0_) not in ((d_515_v40_) | (d_515_v40_))
            rhs95_ = ((d_469_v4_).f20)[default__.safeIndex((0) - (-724), len((d_469_v4_).f20))]
            lhs80_ = globalState
            lhs81_ = globalState
            lhs82_ = globalState
            lhs80_.f7 = rhs91_
            lhs81_.f1 = rhs92_
            d_512_v36_ = rhs93_
            lhs82_.f6 = rhs94_
            r0 = rhs95_
            d_507_cf14_ = (d_469_v4_).f21
        elif True:
            d_517___mcc_h4_ = source9_.cf12
            d_518_cf12_ = d_517___mcc_h4_
            if (d_465_v0_) or (((self).f11) <= ((d_469_v4_).f21)):
                (globalState).f1 = (d_469_v4_).f21
                (globalState).f7 = False
                d_519_v41_: _dafny.Array
                def lambda19_(d_520_i3_):
                    return (d_520_i3_) * ((self).f11)

                init9_ = lambda19_
                nw79_ = _dafny.Array(None, 2)
                for i0_9_ in range(nw79_.length(0)):
                    nw79_[i0_9_] = init9_(i0_9_)
                d_519_v41_ = nw79_
                d_521_v42_: _dafny.Seq
                d_521_v42_ = _dafny.SeqWithoutIsStrInference([d_519_v41_, d_519_v41_, d_519_v41_])
                d_522_v43_: _dafny.Seq
                d_522_v43_ = _dafny.SeqWithoutIsStrInference([(d_521_v42_)[default__.safeIndex((d_469_v4_).f21, len(d_521_v42_))], d_519_v41_, d_519_v41_])
                d_523_v44_: _dafny.Seq
                d_523_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "miujifbj"))
                d_524_v45_: _dafny.Seq
                d_524_v45_ = _dafny.SeqWithoutIsStrInference([(d_469_v4_).f21])
                d_525_v46_: _dafny.Seq
                d_525_v46_ = _dafny.SeqWithoutIsStrInference([d_524_v45_])
                d_526_v47_: D9
                d_526_v47_ = D9_DC19()
                d_527_v48_: _dafny.Map
                d_527_v48_ = _dafny.Map({_dafny.MultiSet([(d_469_v4_).f21]): d_466_v1_})
                d_528_v50_: _dafny.Map
                d_528_v50_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfr"))): _dafny.Set({(self).f11})})
                d_529_v52_: _dafny.Set
                d_529_v52_ = _dafny.Set({(d_469_v4_).f21, (self).f11, (d_469_v4_).f21})
                d_530_v53_: _dafny.Array
                nw80_ = _dafny.Array(None, 20)
                nw80_[int(0)] = (d_465_v0_ if not(d_465_v0_) else d_465_v0_)
                nw80_[int(1)] = d_465_v0_
                nw80_[int(2)] = (d_525_v46_) == (d_525_v46_)
                nw80_[int(3)] = d_465_v0_
                nw80_[int(4)] = (d_469_v4_).fm21(d_465_v0_, (d_469_v4_).f21, globalState)
                nw80_[int(5)] = d_465_v0_
                nw80_[int(6)] = ((d_469_v4_).f21) in (default__.fm30((d_469_v4_).f21, False, (d_469_v4_).f21, 930, globalState))
                nw80_[int(7)] = True
                nw80_[int(8)] = d_465_v0_
                nw80_[int(9)] = (d_526_v47_) not in (_dafny.SeqWithoutIsStrInference([d_526_v47_]))
                nw80_[int(10)] = d_465_v0_
                def iife62_():
                    coll38_ = _dafny.Map()
                    compr_38_: _dafny.Seq
                    for compr_38_ in (default__.fm31((self).f24, globalState)).Elements:
                        d_531_v49_: _dafny.Seq = compr_38_
                        if (d_531_v49_) in (default__.fm31((self).f24, globalState)):
                            coll38_[d_531_v49_] = d_465_v0_
                    return _dafny.Map(coll38_)
                nw80_[int(11)] = (self).fm4((self).f11, d_527_v48_, len(iife62_()
                ), False, globalState)
                nw80_[int(12)] = ((d_468_v3_)[default__.safeIndex((self).f11, len(d_468_v3_))] if d_465_v0_ else d_465_v0_)
                nw80_[int(13)] = False
                def iife63_():
                    coll39_ = _dafny.Set()
                    compr_39_: int
                    for compr_39_ in (_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21])).Elements:
                        d_532_v51_: int = compr_39_
                        if (d_532_v51_) in (_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21])):
                            coll39_ = coll39_.union(_dafny.Set([(d_532_v51_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knsv"))))]))
                    return _dafny.Set(coll39_)
                nw80_[int(14)] = not((d_529_v52_).issubset(((d_528_v50_)[(d_469_v4_).f21] if ((d_469_v4_).f21) in (d_528_v50_) else iife63_()
                )))
                nw80_[int(15)] = d_465_v0_
                nw80_[int(16)] = False
                nw80_[int(17)] = d_465_v0_
                nw80_[int(18)] = ((d_469_v4_).f20)[default__.safeIndex((0) - ((self).f11), len((d_469_v4_).f20))]
                nw80_[int(19)] = d_465_v0_
                d_530_v53_ = nw80_
                index68_ = default__.safeIndex(597, (d_530_v53_).length(0))
                (d_530_v53_)[index68_] = not(d_465_v0_)
                index69_ = default__.safeIndex(597, (d_530_v53_).length(0))
                rhs96_ = ((d_521_v42_) + (d_522_v43_)) + (d_522_v43_)
                rhs97_ = (default__.fm1(globalState)) + ((d_523_v44_ if not(True) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtq"))))
                rhs98_ = True
                rhs99_ = (d_469_v4_).f21
                lhs83_ = d_530_v53_
                lhs84_ = default__.safeIndex(597, (d_530_v53_).length(0))
                lhs85_ = globalState
                d_522_v43_ = rhs96_
                d_523_v44_ = rhs97_
                lhs83_[lhs84_] = rhs98_
                lhs85_.f1 = rhs99_
                d_533_v54_: _dafny.Map
                d_533_v54_ = _dafny.Map({(d_530_v53_)[default__.safeIndex(597, (d_530_v53_).length(0))]: len(d_466_v1_)})
                d_534_v55_: _dafny.Map
                d_534_v55_ = _dafny.Map({((d_469_v4_).f21) * (len(d_523_v44_)): d_533_v54_})
                d_535_v57_: _dafny.MultiSet
                d_535_v57_ = _dafny.MultiSet([(d_469_v4_).f21])
                index70_ = default__.safeIndex(597, (d_530_v53_).length(0))
                index71_ = default__.safeIndex(597, (d_530_v53_).length(0))
                def iife64_():
                    coll40_ = _dafny.Map()
                    compr_40_: _dafny.MultiSet
                    for compr_40_ in (_dafny.Set({d_535_v57_})).Elements:
                        d_536_v56_: _dafny.MultiSet = compr_40_
                        if (d_536_v56_) in (_dafny.Set({d_535_v57_})):
                            coll40_[d_536_v56_] = d_466_v1_
                    return _dafny.Map(coll40_)
                rhs100_ = ((d_530_v53_)[default__.safeIndex(597, (d_530_v53_).length(0))] if (self).fm4((d_469_v4_).f21, d_527_v48_, (self).f11, (self).fm4(635, iife64_()
                , (d_469_v4_).f21, (d_530_v53_)[default__.safeIndex(597, (d_530_v53_).length(0))], globalState), globalState) else (d_530_v53_)[default__.safeIndex(597, (d_530_v53_).length(0))])
                rhs101_ = not ((d_470_v5_) != (d_518_cf12_)) or ((d_465_v0_ if d_465_v0_ else (d_530_v53_)[default__.safeIndex(597, (d_530_v53_).length(0))]))
                rhs102_ = d_465_v0_
                rhs103_ = (default__.fm32((d_469_v4_).f21, (d_469_v4_).f21, (d_469_v4_).f20, (d_469_v4_).f21, globalState)) | (d_534_v55_)
                rhs104_ = (d_530_v53_)[default__.safeIndex(597, (d_530_v53_).length(0))]
                lhs86_ = d_530_v53_
                lhs87_ = default__.safeIndex(597, (d_530_v53_).length(0))
                lhs88_ = d_530_v53_
                lhs89_ = default__.safeIndex(597, (d_530_v53_).length(0))
                r2 = rhs100_
                r0 = rhs101_
                lhs86_[lhs87_] = rhs102_
                d_534_v55_ = rhs103_
                lhs88_[lhs89_] = rhs104_
                d_537_v58_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.Set({}), 6)
                d_537_v58_ = nw81_
                d_538_v59_: _dafny.Map
                d_538_v59_ = _dafny.Map({d_537_v58_: d_530_v53_})
                d_538_v59_ = (d_538_v59_).set(d_537_v58_, d_530_v53_)
            elif True:
                (globalState).f6 = (not(d_465_v0_)) and (((d_469_v4_).f21) in (_dafny.MultiSet([(d_469_v4_).f21])))
                d_539_v60_: _dafny.Map
                d_539_v60_ = _dafny.Map({d_465_v0_: _dafny.SeqWithoutIsStrInference([(d_469_v4_).f21, (d_469_v4_).f21])})
                d_540_v61_: _dafny.Seq
                d_540_v61_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])
                d_541_v62_: _dafny.Seq
                d_541_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f24 for d_542_i4_ in range(default__.abs(442))])])
                d_543_v63_: D7
                d_543_v63_ = D7_DC12(_dafny.SeqWithoutIsStrInference([len(d_540_v61_), len((d_541_v62_)[default__.safeIndex((self).f11, len(d_541_v62_))])]))
                d_544_v64_: _dafny.Seq
                d_544_v64_ = _dafny.SeqWithoutIsStrInference([(self).f11, len(d_540_v61_), (d_469_v4_).f21])
                pat_let_tv26_ = d_544_v64_
                def iife65_(_pat_let12_0):
                    def iife66_(d_545_dt__update__tmp_h0_):
                        def iife67_(_pat_let13_0):
                            def iife68_(d_546_dt__update_hcf17_h0_):
                                return D7_DC12(d_546_dt__update_hcf17_h0_)
                            return iife68_(_pat_let13_0)
                        return iife67_(pat_let_tv26_)
                    return iife66_(_pat_let12_0)
                d_539_v60_ = (d_539_v60_).set(default__.fm0(globalState), ((iife65_(d_543_v63_)).cf17) + (d_544_v64_))
                d_547_v66_: _dafny.MultiSet
                def iife69_():
                    coll41_ = _dafny.Set()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(891, 318):
                        d_548_v65_: int = compr_41_
                        if ((891) <= (d_548_v65_)) and ((d_548_v65_) < (318)):
                            coll41_ = coll41_.union(_dafny.Set([(d_548_v65_) * (len(_dafny.SeqWithoutIsStrInference([(d_469_v4_).f21 for d_549_i5_ in range(default__.abs(338))])))]))
                    return _dafny.Set(coll41_)
                d_547_v66_ = _dafny.MultiSet([iife69_()
                ])
                d_547_v66_ = (d_547_v66_) - (d_547_v66_)
                d_550_v67_: _dafny.MultiSet
                d_550_v67_ = _dafny.MultiSet([(d_469_v4_).f21, (d_469_v4_).f21, (d_469_v4_).f21])
                d_550_v67_ = d_550_v67_
                (globalState).f1 = ((d_469_v4_).f21) - ((d_469_v4_).f21)
            d_551_v68_: _dafny.Map
            d_551_v68_ = _dafny.Map({not(not(d_465_v0_)): (self).f11})
            d_552_v69_: _dafny.Set
            d_552_v69_ = _dafny.Set({(self).fm3(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbmbrcsg"))).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbmbrcsg")))), (self).f24)), d_465_v0_, d_465_v0_, globalState), (d_469_v4_).f21, default__.fm2((d_469_v4_).f21, (self).f11, globalState), len(d_551_v68_)})
            d_552_v69_ = d_552_v69_
            d_553_v70_: str
            d_553_v70_ = _dafny.CodePoint('d')
            d_553_v70_ = ((self).f24 if d_465_v0_ else d_553_v70_)
            d_554_v71_: D9
            d_554_v71_ = D9_DC19()
            d_554_v71_ = d_554_v71_
        d_555_v72_: _dafny.Seq
        d_555_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
        d_556_v73_: _dafny.Map
        d_556_v73_ = _dafny.Map({d_555_v72_: d_465_v0_})
        r2 = (d_556_v73_) == (d_556_v73_)
        d_557_v74_: _dafny.MultiSet
        d_557_v74_ = _dafny.MultiSet([(self).f11, (d_469_v4_).f21])
        d_558_v76_: _dafny.Map
        def iife70_():
            coll42_ = _dafny.Map()
            compr_42_: int
            for compr_42_ in _dafny.IntegerRange(679, 599):
                d_559_v75_: int = compr_42_
                if ((679) <= (d_559_v75_)) and ((d_559_v75_) < (599)):
                    coll42_[default__.safeDivisionInt(d_559_v75_, (D7_DC13(_dafny.Set({d_465_v0_}), (self).f11, (d_469_v4_).f21)).cf20)] = len(d_555_v72_)
            return _dafny.Map(coll42_)
        d_558_v76_ = _dafny.Map({d_557_v74_: iife70_()
        })
        d_560_v77_: D0
        d_560_v77_ = D0_DC1(d_465_v0_, (self).f11, d_465_v0_)
        r2 = not((self).fm4(-452, d_558_v76_, (d_469_v4_).f21, (d_560_v77_).cf1, globalState))
        d_561_v78_: _dafny.Set
        d_561_v78_ = _dafny.Set({False, d_465_v0_})
        d_562_v79_: _dafny.Map
        d_562_v79_ = _dafny.Map({(self).f24: d_561_v78_})
        d_563_v80_: D7
        d_563_v80_ = D7_DC13(((d_562_v79_)[(self).f24] if ((self).f24) in (d_562_v79_) else d_561_v78_), (d_469_v4_).f21, (self).f11)
        d_564_v81_: _dafny.Map
        d_564_v81_ = _dafny.Map({(self).f24: d_465_v0_})
        d_563_v80_ = D7_DC13((_dafny.Set({d_465_v0_, ((d_564_v81_)[(self).f24] if ((self).f24) in (d_564_v81_) else d_465_v0_)})) | (d_561_v78_), (d_469_v4_).f21, 71)
        d_565_v82_: _dafny.Seq
        d_565_v82_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_566_v83_: _dafny.Seq
        d_566_v83_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_565_v82_: d_465_v0_}))])
        d_567_v84_: _dafny.Array
        nw82_ = _dafny.Array(int(0), 5)
        d_567_v84_ = nw82_
        d_568_v85_: _dafny.Map
        d_568_v85_ = _dafny.Map({(self).f11: d_567_v84_})
        r0 = (d_469_v4_).fm21((_dafny.MultiSet([(self).f11])).ispropersubset(d_557_v74_), default__.safeModuloInt(len(d_565_v82_), len(d_568_v85_)), globalState)
        r1 = ((d_556_v73_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbqkmva"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnu")))] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbqkmva"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnu")))) in (d_556_v73_) else d_465_v0_)
        d_569_v86_: _dafny.Seq
        d_569_v86_ = _dafny.SeqWithoutIsStrInference([d_468_v3_])
        d_570_v87_: _dafny.Seq
        d_570_v87_ = _dafny.SeqWithoutIsStrInference([d_569_v86_, _dafny.SeqWithoutIsStrInference([d_468_v3_ for d_571_i6_ in range(default__.abs(308))]), d_569_v86_, d_569_v86_])
        r2 = (((self).f11) * ((d_469_v4_).f21)) < (len((d_570_v87_)[default__.safeIndex((d_469_v4_).f21, len(d_570_v87_))]))
        r3 = ((d_469_v4_).f21) - ((self).f11)
        return r0, r1, r2, r3

    def m9(self, p0, p1, globalState):
        r0: D7 = D7.default()()
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_572_i0_: int
        d_572_i0_ = 0
        with _dafny.label("6"):
            while p0:
                with _dafny.c_label("6"):
                    if (d_572_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_572_i0_ = (d_572_i0_) + (1)
                    d_573_v0_: D0
                    d_573_v0_ = D0_DC1(p0, (self).f11, default__.fm0(globalState))
                    pat_let_tv27_ = p0
                    pat_let_tv28_ = p0
                    pat_let_tv29_ = p0
                    def iife71_(_pat_let14_0):
                        def iife72_(d_574_dt__update__tmp_h0_):
                            def iife73_(_pat_let15_0):
                                def iife74_(d_575_dt__update_hcf1_h0_):
                                    return D0_DC1(d_575_dt__update_hcf1_h0_, (d_574_dt__update__tmp_h0_).cf2, (d_574_dt__update__tmp_h0_).cf3)
                                return iife74_(_pat_let15_0)
                            return iife73_((pat_let_tv27_ if pat_let_tv28_ else pat_let_tv29_))
                        return iife72_(_pat_let14_0)
                    source10_ = iife71_(d_573_v0_)
                    if source10_.is_DC1:
                        d_576___mcc_h0_ = source10_.cf1
                        d_577___mcc_h1_ = source10_.cf2
                        d_578___mcc_h2_ = source10_.cf3
                        d_579_cf3_ = d_578___mcc_h2_
                        d_580_cf2_ = d_577___mcc_h1_
                        d_581_cf1_ = d_576___mcc_h0_
                        (globalState).f6 = False
                        d_582_v1_: bool
                        d_583_v2_: bool
                        d_584_v3_: bool
                        d_585_v4_: int
                        out16_: bool
                        out17_: bool
                        out18_: bool
                        out19_: int
                        out16_, out17_, out18_, out19_ = (self).m8(globalState)
                        d_582_v1_ = out16_
                        d_583_v2_ = out17_
                        d_584_v3_ = out18_
                        d_585_v4_ = out19_
                        d_586_v5_: _dafny.Array
                        nw83_ = _dafny.Array(_dafny.Set({}), 5)
                        d_586_v5_ = nw83_
                        d_586_v5_ = d_586_v5_
                        d_587_v6_: D9
                        d_587_v6_ = D9_DC19()
                        d_587_v6_ = d_587_v6_
                    elif True:
                        d_588___mcc_h3_ = source10_.cf0
                        d_589_cf0_ = d_588___mcc_h3_
                        d_590_v7_: _dafny.Map
                        d_590_v7_ = _dafny.Map({p0: p0})
                        (globalState).f6 = (len((d_590_v7_) | (d_590_v7_))) > (d_589_cf0_)
                        d_591_v8_: _dafny.Map
                        d_591_v8_ = _dafny.Map({d_589_cf0_: d_589_cf0_})
                        d_592_v9_: _dafny.Seq
                        d_592_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_591_v8_ for d_593_i1_ in range(default__.abs(-762))]), _dafny.SeqWithoutIsStrInference([d_591_v8_ for d_594_i2_ in range(default__.abs(-221))])])
                        d_595_v10_: _dafny.Array
                        nw84_ = _dafny.Array(None, 11)
                        nw84_[int(0)] = d_589_cf0_
                        nw84_[int(1)] = (self).f11
                        nw84_[int(2)] = (self).f11
                        nw84_[int(3)] = (len(d_591_v8_)) + (len((d_592_v9_)[default__.safeIndex((self).f11, len(d_592_v9_))]))
                        nw84_[int(4)] = d_589_cf0_
                        nw84_[int(5)] = (self).f11
                        nw84_[int(6)] = default__.safeDivisionInt((self).f11, (self).f11)
                        nw84_[int(7)] = (0) - (d_589_cf0_)
                        nw84_[int(8)] = len(_dafny.Set({not(p0)}))
                        nw84_[int(9)] = d_589_cf0_
                        nw84_[int(10)] = (self).f11
                        d_595_v10_ = nw84_
                        d_596_v11_: _dafny.Set
                        d_596_v11_ = _dafny.Set({d_589_cf0_})
                        d_597_v12_: _dafny.Map
                        d_597_v12_ = _dafny.Map({d_589_cf0_: d_596_v11_})
                        rhs105_ = d_595_v10_
                        rhs106_ = ((0) - (d_589_cf0_)) - ((self).f11)
                        rhs107_ = p0
                        rhs108_ = (len(((d_597_v12_)[d_589_cf0_] if (d_589_cf0_) in (d_597_v12_) else _dafny.Set({d_589_cf0_, (self).f11})))) != ((self).f11)
                        lhs90_ = globalState
                        lhs91_ = globalState
                        lhs92_ = globalState
                        d_595_v10_ = rhs105_
                        lhs90_.f1 = rhs106_
                        lhs91_.f7 = rhs107_
                        lhs92_.f7 = rhs108_
                        d_598_v13_: _dafny.MultiSet
                        d_598_v13_ = _dafny.MultiSet([d_589_cf0_])
                        d_599_v14_: _dafny.Map
                        d_599_v14_ = _dafny.Map({(_dafny.MultiSet([(self).f11])) | (d_598_v13_): p0})
                        (globalState).f7 = ((d_599_v14_)[(d_598_v13_) - (_dafny.MultiSet([(self).f11]))] if ((d_598_v13_) - (_dafny.MultiSet([(self).f11]))) in (d_599_v14_) else p0)
                        d_600_v15_: _dafny.Seq
                        d_600_v15_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                        d_601_v16_: _dafny.Map
                        d_601_v16_ = _dafny.Map({p1: False})
                        d_602_v17_: C0
                        nw85_ = C0()
                        nw85_.ctor__(d_600_v15_, default__.fm2(d_589_cf0_, len(d_601_v16_), globalState))
                        d_602_v17_ = nw85_
                        d_603_v18_: _dafny.Map
                        d_603_v18_ = _dafny.Map({p0: d_602_v17_})
                        d_604_v19_: _dafny.Map
                        d_604_v19_ = _dafny.Map({((d_603_v18_)[p0] if (p0) in (d_603_v18_) else d_602_v17_): d_589_cf0_})
                        (globalState).f1 = ((d_604_v19_)[d_602_v17_] if (d_602_v17_) in (d_604_v19_) else 893)
                    d_605_v20_: _dafny.Array
                    nw86_ = _dafny.Array(_dafny.Array(None, 0), 13)
                    d_605_v20_ = nw86_
                    d_606_v21_: _dafny.Array
                    nw87_ = _dafny.Array(int(0), 25)
                    d_606_v21_ = nw87_
                    index72_ = default__.safeIndex(402, (d_605_v20_).length(0))
                    (d_605_v20_)[index72_] = d_606_v21_
                    d_607_v22_: _dafny.Seq
                    d_607_v22_ = _dafny.SeqWithoutIsStrInference([d_606_v21_, d_606_v21_])
                    index73_ = default__.safeIndex(402, (d_605_v20_).length(0))
                    (d_605_v20_)[index73_] = (d_607_v22_)[default__.safeIndex((self).f11, len(d_607_v22_))]
                    d_608_v23_: _dafny.Seq
                    d_608_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vb"))
                    d_609_v24_: _dafny.Seq
                    d_609_v24_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_610_v25_: _dafny.MultiSet
                    d_610_v25_ = _dafny.MultiSet([(d_609_v24_ if p0 else d_609_v24_), d_609_v24_, (d_609_v24_) + (d_609_v24_), d_609_v24_])
                    d_611_v26_: _dafny.MultiSet
                    d_611_v26_ = _dafny.MultiSet([p0])
                    d_612_v27_: _dafny.Map
                    d_612_v27_ = _dafny.Map({(self).f11: d_611_v26_})
                    d_613_v29_: _dafny.Map
                    d_613_v29_ = _dafny.Map({(self).f11: (self).f11})
                    d_614_v30_: _dafny.Map
                    def iife75_():
                        coll43_ = _dafny.Set()
                        compr_43_: str
                        for compr_43_ in ((d_608_v23_).set(default__.safeIndex((self).f11, len(d_608_v23_)), (self).f24)).Elements:
                            d_615_v28_: str = compr_43_
                            if (d_615_v28_) in ((d_608_v23_).set(default__.safeIndex((self).f11, len(d_608_v23_)), (self).f24)):
                                coll43_ = coll43_.union(_dafny.Set([d_615_v28_]))
                        return _dafny.Set(coll43_)
                    d_614_v30_ = _dafny.Map({_dafny.MultiSet([len(iife75_()
                    )]): d_613_v29_})
                    rhs109_ = (((d_608_v23_).set(default__.safeIndex(len((d_612_v27_).set((self).f11, _dafny.MultiSet([p0, p0]))), len(d_608_v23_)), (self).f24)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjdgtkgo")))) + ((d_608_v23_) + (d_608_v23_))
                    rhs110_ = not (True) or (default__.fm0(globalState))
                    rhs111_ = (d_610_v25_ if (p0) and (p0) else (d_610_v25_) - (_dafny.MultiSet([d_609_v24_, d_609_v24_, d_609_v24_, d_609_v24_, d_609_v24_])))
                    rhs112_ = ((d_611_v26_)[(self).fm4((self).f11, d_614_v30_, (self).f11, p0, globalState)] if ((self).fm4((self).f11, d_614_v30_, (self).f11, p0, globalState)) in (d_611_v26_) else (self).f11)
                    lhs93_ = globalState
                    lhs94_ = globalState
                    d_608_v23_ = rhs109_
                    lhs93_.f7 = rhs110_
                    d_610_v25_ = rhs111_
                    lhs94_.f1 = rhs112_
                    (globalState).f7 = p0
                    pass
            pass
        d_616_v31_: _dafny.MultiSet
        d_616_v31_ = _dafny.MultiSet([(self).f11])
        d_617_v32_: _dafny.Seq
        d_617_v32_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
        hi0_ = 770
        for d_618_i3_ in range(((d_616_v31_) - (_dafny.MultiSet(d_617_v32_))).cardinality, hi0_):
            (globalState).f6 = p0
            d_619_v33_: D8
            d_619_v33_ = D8_DC16(p0, d_618_i3_)
            source11_ = d_619_v33_
            if source11_.is_DC16:
                d_620___mcc_h4_ = source11_.cf23
                d_621___mcc_h5_ = source11_.cf24
                d_622_cf24_ = d_621___mcc_h5_
                d_623_cf23_ = d_620___mcc_h4_
                d_624_v34_: _dafny.Array
                def lambda20_(d_625_i4_):
                    return default__.safeModuloInt(d_625_i4_, (self).f11)

                init10_ = lambda20_
                nw88_ = _dafny.Array(None, 4)
                for i0_10_ in range(nw88_.length(0)):
                    nw88_[i0_10_] = init10_(i0_10_)
                d_624_v34_ = nw88_
                d_626_v35_: _dafny.Map
                d_626_v35_ = _dafny.Map({d_624_v34_: p0})
                d_627_v36_: _dafny.Seq
                d_627_v36_ = _dafny.SeqWithoutIsStrInference([p0, d_623_cf23_])
                d_628_v37_: _dafny.MultiSet
                d_628_v37_ = _dafny.MultiSet([(d_627_v36_)[default__.safeIndex(d_622_cf24_, len(d_627_v36_))], not(p0)])
                d_629_v38_: _dafny.MultiSet
                d_629_v38_ = d_628_v37_
                d_630_v39_: _dafny.Set
                d_630_v39_ = _dafny.Set({p0, p0})
                d_631_v40_: _dafny.Map
                d_631_v40_ = _dafny.Map({True: (_dafny.Set({d_623_cf23_, default__.fm0(globalState), d_623_cf23_, p0, p0})).ispropersubset(d_630_v39_)})
                d_632_v41_: D10
                d_632_v41_ = D10_DC20(d_626_v35_)
                rhs113_ = ((d_631_v40_)[d_623_cf23_] if (d_623_cf23_) in (d_631_v40_) else False)
                rhs114_ = (d_626_v35_) | ((d_632_v41_).cf27)
                rhs115_ = (self).f11
                rhs116_ = d_629_v38_
                lhs95_ = globalState
                lhs96_ = globalState
                lhs95_.f6 = rhs113_
                d_626_v35_ = rhs114_
                lhs96_.f1 = rhs115_
                d_629_v38_ = rhs116_
                d_633_v42_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_633_v42_ = nw89_
                d_634_v43_: _dafny.Array
                d_634_v43_ = d_633_v42_
                d_635_v44_: _dafny.Array
                nw90_ = _dafny.Array(None, 12)
                nw90_[int(0)] = d_633_v42_
                nw90_[int(1)] = d_634_v43_
                nw90_[int(2)] = d_634_v43_
                nw90_[int(3)] = d_634_v43_
                nw90_[int(4)] = d_634_v43_
                nw90_[int(5)] = d_634_v43_
                nw90_[int(6)] = d_634_v43_
                nw90_[int(7)] = d_634_v43_
                nw90_[int(8)] = d_634_v43_
                nw90_[int(9)] = d_633_v42_
                nw90_[int(10)] = d_634_v43_
                nw90_[int(11)] = d_634_v43_
                d_635_v44_ = nw90_
                d_636_v45_: _dafny.Map
                d_636_v45_ = _dafny.Map({d_635_v44_: d_624_v34_})
                d_637_v46_: _dafny.Map
                d_637_v46_ = _dafny.Map({d_636_v45_: p0})
                d_638_v47_: D11
                d_638_v47_ = D11_DC22((d_636_v45_).set(d_635_v44_, d_624_v34_))
                d_637_v46_ = (d_637_v46_).set((d_636_v45_) | ((d_638_v47_).cf30), p0)
                d_639_v48_: _dafny.Map
                d_639_v48_ = _dafny.Map({(self).f24: (d_618_i3_) - (d_622_cf24_)})
                index74_ = default__.safeIndex(313, (d_624_v34_).length(0))
                (d_624_v34_)[index74_] = -915
                index75_ = default__.safeIndex(313, (d_624_v34_).length(0))
                rhs117_ = _dafny.Map({(self).f24: 196})
                rhs118_ = (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f24 for d_640_i5_ in range(default__.abs(414))])))
                rhs119_ = (0) - (((self).f11) * (630))
                lhs97_ = d_624_v34_
                lhs98_ = default__.safeIndex(313, (d_624_v34_).length(0))
                d_639_v48_ = rhs117_
                lhs97_[lhs98_] = rhs118_
                d_622_cf24_ = rhs119_
                d_641_v49_: _dafny.Map
                d_641_v49_ = _dafny.Map({d_633_v42_: d_622_cf24_})
                d_641_v49_ = (d_641_v49_).set((d_634_v43_ if d_623_cf23_ else d_634_v43_), (0) - (default__.safeModuloInt(d_622_cf24_, (self).fm3((self).f11, d_623_cf23_, p0, globalState))))
            elif source11_.is_DC15:
                d_642___mcc_h6_ = source11_.cf22
                d_643_cf22_ = d_642___mcc_h6_
                d_617_v32_ = d_617_v32_
                d_644_v50_: _dafny.Seq
                d_644_v50_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])
                d_645_v51_: D6
                d_645_v51_ = D6_DC10((d_644_v50_) + (d_644_v50_))
                d_645_v51_ = d_645_v51_
                d_646_v52_: D11
                d_646_v52_ = D11_DC23(False, d_618_i3_, 662)
                d_647_v53_: _dafny.Seq
                d_647_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                d_648_v54_: _dafny.MultiSet
                d_648_v54_ = _dafny.MultiSet([p0, not(p0)])
                d_649_v55_: _dafny.Array
                nw91_ = _dafny.Array(None, 27)
                nw91_[int(0)] = (d_616_v31_) | (d_616_v31_)
                nw91_[int(1)] = _dafny.MultiSet([(self).f11])
                nw91_[int(2)] = d_616_v31_
                nw91_[int(3)] = d_616_v31_
                nw91_[int(4)] = (d_616_v31_).intersection(d_616_v31_)
                nw91_[int(5)] = _dafny.MultiSet([d_618_i3_, len(_dafny.SeqWithoutIsStrInference([d_618_i3_])), d_618_i3_, (self).f11, (self).f11])
                nw91_[int(6)] = d_616_v31_
                nw91_[int(7)] = (_dafny.MultiSet([(d_646_v52_).cf33])) - (d_616_v31_)
                nw91_[int(8)] = d_616_v31_
                nw91_[int(9)] = (d_616_v31_ if p0 else d_616_v31_)
                nw91_[int(10)] = d_616_v31_
                nw91_[int(11)] = d_616_v31_
                nw91_[int(12)] = d_616_v31_
                nw91_[int(13)] = d_616_v31_
                nw91_[int(14)] = (d_616_v31_).intersection(d_616_v31_)
                nw91_[int(15)] = d_616_v31_
                nw91_[int(16)] = (d_616_v31_).set(default__.fm2(888, len(d_647_v53_), globalState), default__.abs((self).f11))
                nw91_[int(17)] = ((_dafny.MultiSet([58, (self).f11])).set(d_618_i3_, default__.abs(d_618_i3_))).intersection(d_616_v31_)
                nw91_[int(18)] = d_616_v31_
                nw91_[int(19)] = d_616_v31_
                nw91_[int(20)] = d_616_v31_
                nw91_[int(21)] = d_616_v31_
                nw91_[int(22)] = d_616_v31_
                nw91_[int(23)] = ((d_616_v31_).set((self).f11, default__.abs((d_648_v54_).cardinality))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_618_i3_ for d_650_i6_ in range(default__.abs(436))])))
                nw91_[int(24)] = d_616_v31_
                nw91_[int(25)] = _dafny.MultiSet([d_618_i3_, d_618_i3_])
                nw91_[int(26)] = _dafny.MultiSet([(self).f11])
                d_649_v55_ = nw91_
                d_649_v55_ = d_649_v55_
                d_651_v56_: C0
                nw92_ = C0()
                nw92_.ctor__(d_644_v50_, ((self).f11) - (-174))
                d_651_v56_ = nw92_
            elif True:
                d_652___mcc_h7_ = source11_.cf25
                d_653_cf25_ = d_652___mcc_h7_
                d_654_v57_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.Map({}), 10)
                d_654_v57_ = nw93_
                d_655_v58_: _dafny.MultiSet
                d_655_v58_ = _dafny.MultiSet([(self).f24])
                d_656_v59_: _dafny.Array
                nw94_ = _dafny.Array(int(0), 20)
                d_656_v59_ = nw94_
                d_657_v60_: _dafny.Seq
                d_657_v60_ = _dafny.SeqWithoutIsStrInference([d_656_v59_])
                d_658_v61_: _dafny.Map
                d_658_v61_ = _dafny.Map({d_655_v58_: (d_657_v60_)[default__.safeIndex((self).f11, len(d_657_v60_))]})
                index76_ = default__.safeIndex(774, (d_654_v57_).length(0))
                (d_654_v57_)[index76_] = d_658_v61_
                index77_ = default__.safeIndex(774, (d_654_v57_).length(0))
                rhs120_ = ((d_658_v61_) | (d_658_v61_)) | ((d_658_v61_).set(d_655_v58_, d_656_v59_))
                rhs121_ = len(((d_657_v60_) + (d_657_v60_)) + (d_657_v60_))
                rhs122_ = p0
                rhs123_ = (d_617_v32_) + ((_dafny.SeqWithoutIsStrInference([(self).f11])) + (d_617_v32_))
                lhs99_ = d_654_v57_
                lhs100_ = default__.safeIndex(774, (d_654_v57_).length(0))
                lhs101_ = globalState
                lhs102_ = globalState
                lhs99_[lhs100_] = rhs120_
                lhs101_.f1 = rhs121_
                lhs102_.f7 = rhs122_
                d_617_v32_ = rhs123_
                d_659_v62_: C1
                nw95_ = C1()
                nw95_.ctor__(p0, not((default__.fm2((self).f11, d_618_i3_, globalState)) > (d_618_i3_)))
                d_659_v62_ = nw95_
                d_660_v63_: _dafny.Seq
                d_660_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgsyppx"))
                d_661_v64_: _dafny.MultiSet
                d_661_v64_ = _dafny.MultiSet([d_660_v63_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gg"))])
                d_662_v65_: D2
                d_662_v65_ = D2_DC3(d_661_v64_)
                d_663_v66_: _dafny.Seq
                d_663_v66_ = _dafny.SeqWithoutIsStrInference([(d_659_v62_).f19, (d_659_v62_).fm13(d_662_v65_, default__.fm0(globalState), (d_659_v62_).f18, (self).f11, globalState)])
                d_664_v67_: D6
                d_664_v67_ = D6_DC10(d_663_v66_)
                pat_let_tv30_ = d_663_v66_
                pat_let_tv31_ = d_659_v62_
                pat_let_tv32_ = d_659_v62_
                d_665_v68_: _dafny.Array
                nw96_ = _dafny.Array(None, 25)
                nw96_[int(0)] = d_664_v67_
                nw96_[int(1)] = D6_DC10(d_663_v66_)
                nw96_[int(2)] = D6_DC10((d_664_v67_).cf13)
                nw96_[int(3)] = d_664_v67_
                nw96_[int(4)] = d_664_v67_
                nw96_[int(5)] = D6_DC10(d_663_v66_)
                nw96_[int(6)] = default__.fm33(p0, globalState)
                nw96_[int(7)] = d_664_v67_
                nw96_[int(8)] = d_664_v67_
                nw96_[int(9)] = d_664_v67_
                nw96_[int(10)] = d_664_v67_
                nw96_[int(11)] = d_664_v67_
                nw96_[int(12)] = d_664_v67_
                nw96_[int(13)] = d_664_v67_
                nw96_[int(14)] = d_664_v67_
                def iife76_(_pat_let16_0):
                    def iife77_(d_666_dt__update__tmp_h1_):
                        def iife78_(_pat_let17_0):
                            def iife79_(d_667_dt__update_hcf13_h0_):
                                return D6_DC10(d_667_dt__update_hcf13_h0_)
                            return iife79_(_pat_let17_0)
                        return iife78_(pat_let_tv30_)
                    return iife77_(_pat_let16_0)
                nw96_[int(15)] = iife76_(d_664_v67_)
                nw96_[int(16)] = D6_DC10(d_663_v66_)
                nw96_[int(17)] = d_664_v67_
                nw96_[int(18)] = D6_DC10(d_663_v66_)
                nw96_[int(19)] = d_664_v67_
                nw96_[int(20)] = D6_DC10(_dafny.SeqWithoutIsStrInference([(d_659_v62_).f19, (d_659_v62_).f19]))
                def iife80_(_pat_let18_0):
                    def iife81_(d_668_dt__update__tmp_h2_):
                        def iife82_(_pat_let19_0):
                            def iife83_(d_669_dt__update_hcf13_h1_):
                                return D6_DC10(d_669_dt__update_hcf13_h1_)
                            return iife83_(_pat_let19_0)
                        return iife82_(_dafny.SeqWithoutIsStrInference([(pat_let_tv31_).f19, not((pat_let_tv32_).f18)]))
                    return iife81_(_pat_let18_0)
                nw96_[int(21)] = iife80_(d_664_v67_)
                nw96_[int(22)] = d_664_v67_
                nw96_[int(23)] = d_664_v67_
                nw96_[int(24)] = d_664_v67_
                d_665_v68_ = nw96_
                index78_ = default__.safeIndex(876, (d_665_v68_).length(0))
                (d_665_v68_)[index78_] = D6_DC10(d_663_v66_)
                index79_ = default__.safeIndex(876, (d_665_v68_).length(0))
                (d_665_v68_)[index79_] = d_664_v67_
                d_670_v69_: _dafny.Map
                d_670_v69_ = _dafny.Map({p0: d_656_v59_})
                d_671_v70_: _dafny.Map
                d_671_v70_ = _dafny.Map({(not(not((d_659_v62_).f18))) in (d_670_v69_): (p0) and ((d_659_v62_).f19)})
                d_672_v71_: _dafny.Set
                d_672_v71_ = _dafny.Set({(d_659_v62_).f19})
                d_671_v70_ = (d_671_v70_).set(((self).f11) > (429), (d_672_v71_).ispropersubset(d_672_v71_))
            d_673_v72_: _dafny.Map
            d_673_v72_ = _dafny.Map({default__.fm0(globalState): (d_617_v32_) + (d_617_v32_)})
            (globalState).f1 = len(((d_673_v72_)[p0] if (p0) in (d_673_v72_) else d_617_v32_))
            d_674_v73_: C1
            nw97_ = C1()
            nw97_.ctor__(p0, p0)
            d_674_v73_ = nw97_
        d_675_v74_: _dafny.Seq
        d_675_v74_ = _dafny.SeqWithoutIsStrInference([p0])
        d_676_v75_: C0
        nw98_ = C0()
        nw98_.ctor__((d_675_v74_) + (d_675_v74_), ((self).f11) - ((self).f11))
        d_676_v75_ = nw98_
        d_677_v76_: _dafny.Map
        d_677_v76_ = _dafny.Map({(self).f11: p0})
        d_678_v77_: D8
        d_678_v77_ = D8_DC16(((d_677_v76_)[(d_676_v75_).f21] if ((d_676_v75_).f21) in (d_677_v76_) else p0), (d_676_v75_).f21)
        pat_let_tv33_ = p0
        pat_let_tv34_ = p0
        def lambda21_(source12_):
            if source12_.is_DC16:
                d_679___mcc_h17_ = source12_.cf23
                d_680___mcc_h18_ = source12_.cf24
                d_681_cf24_ = d_680___mcc_h18_
                d_682_cf23_ = d_679___mcc_h17_
                return True
            elif source12_.is_DC15:
                d_683___mcc_h19_ = source12_.cf22
                d_684_cf22_ = d_683___mcc_h19_
                return True
            elif True:
                d_685___mcc_h20_ = source12_.cf25
                d_686_cf25_ = d_685___mcc_h20_
                return (pat_let_tv33_) == (False)

        def iife84_(_pat_let20_0):
            def iife85_(d_687_dt__update__tmp_h3_):
                def iife86_(_pat_let21_0):
                    def iife87_(d_688_dt__update_hcf23_h0_):
                        return D8_DC16(d_688_dt__update_hcf23_h0_, (d_687_dt__update__tmp_h3_).cf24)
                    return iife87_(_pat_let21_0)
                return iife86_(pat_let_tv34_)
            return iife85_(_pat_let20_0)
        if lambda21_(iife84_(d_678_v77_)):
            d_689_v78_: D2
            d_689_v78_ = D2_DC5()
            source13_ = d_689_v78_
            if source13_.is_DC4:
                d_690___mcc_h8_ = source13_.cf6
                d_691___mcc_h9_ = source13_.cf7
                d_692___mcc_h10_ = source13_.cf8
                d_693_cf8_ = d_692___mcc_h10_
                d_694_cf7_ = d_691___mcc_h9_
                d_695_cf6_ = d_690___mcc_h8_
                d_696_v79_: _dafny.Array
                def lambda22_(d_697_v77_):
                    def lambda23_(d_698_i7_):
                        return d_697_v77_

                    return lambda23_

                init11_ = lambda22_(d_678_v77_)
                nw99_ = _dafny.Array(None, 26)
                for i0_11_ in range(nw99_.length(0)):
                    nw99_[i0_11_] = init11_(i0_11_)
                d_696_v79_ = nw99_
                d_699_v80_: D12
                d_699_v80_ = D12_DC25(d_696_v79_)
                pat_let_tv35_ = d_696_v79_
                def iife88_(_pat_let22_0):
                    def iife89_(d_700_dt__update__tmp_h4_):
                        def iife90_(_pat_let23_0):
                            def iife91_(d_701_dt__update_hcf37_h0_):
                                return D12_DC25(d_701_dt__update_hcf37_h0_)
                            return iife91_(_pat_let23_0)
                        return iife90_(pat_let_tv35_)
                    return iife89_(_pat_let22_0)
                rhs124_ = default__.safeDivisionInt(default__.safeModuloInt((d_676_v75_).f21, len(_dafny.Map({p0: p0}))), (d_676_v75_).f21)
                rhs125_ = (iife88_(d_699_v80_)).cf37
                lhs103_ = globalState
                lhs103_.f1 = rhs124_
                d_696_v79_ = rhs125_
                d_702_v81_: _dafny.Array
                nw100_ = _dafny.Array(int(0), 14)
                d_702_v81_ = nw100_
                d_703_v82_: _dafny.MultiSet
                d_703_v82_ = _dafny.MultiSet([p0, p0])
                index80_ = default__.safeIndex(8, (d_702_v81_).length(0))
                (d_702_v81_)[index80_] = (d_703_v82_).cardinality
                index81_ = default__.safeIndex(8, (d_702_v81_).length(0))
                (d_702_v81_)[index81_] = (self).f11
                r1 = p1
                d_704_v83_: _dafny.Array
                nw101_ = _dafny.Array(None, 10)
                nw101_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_676_v75_).f21, (d_676_v75_).f21, 737])
                nw101_[int(1)] = d_617_v32_
                nw101_[int(2)] = _dafny.SeqWithoutIsStrInference([d_695_cf6_])
                nw101_[int(3)] = d_617_v32_
                nw101_[int(4)] = (_dafny.SeqWithoutIsStrInference([(d_676_v75_).f21])) + (_dafny.SeqWithoutIsStrInference([d_694_cf7_]))
                nw101_[int(5)] = d_617_v32_
                nw101_[int(6)] = d_617_v32_
                nw101_[int(7)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f24])) for d_705_i8_ in range(default__.abs(296))])
                nw101_[int(8)] = d_617_v32_
                nw101_[int(9)] = d_617_v32_
                d_704_v83_ = nw101_
                index82_ = default__.safeIndex(95, (d_704_v83_).length(0))
                (d_704_v83_)[index82_] = (_dafny.SeqWithoutIsStrInference([d_695_cf6_ for d_706_i9_ in range(default__.abs(47))])) + (default__.fm34(p0, D6_DC11(100, d_694_cf7_, (d_676_v75_).f21), globalState))
                index83_ = default__.safeIndex(95, (d_704_v83_).length(0))
                (d_704_v83_)[index83_] = ((_dafny.SeqWithoutIsStrInference([(0) - ((d_702_v81_)[default__.safeIndex(8, (d_702_v81_).length(0))])])) + (d_617_v32_)) + (d_617_v32_)
            elif source13_.is_DC5:
                d_707_v84_: _dafny.Set
                d_707_v84_ = _dafny.Set({p0, False})
                d_708_v85_: _dafny.Array
                nw102_ = _dafny.Array(None, 3)
                nw102_[int(0)] = d_707_v84_
                nw102_[int(1)] = (d_707_v84_) - (_dafny.Set({p0}))
                nw102_[int(2)] = (d_707_v84_) | (d_707_v84_)
                d_708_v85_ = nw102_
                index84_ = default__.safeIndex(418, (d_708_v85_).length(0))
                (d_708_v85_)[index84_] = d_707_v84_
                d_709_v86_: _dafny.Seq
                d_709_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cym"))
                d_710_v87_: D6
                d_710_v87_ = D6_DC11((self).f11, (d_676_v75_).f21, (0) - (len(d_709_v86_)))
                index85_ = default__.safeIndex(418, (d_708_v85_).length(0))
                rhs126_ = default__.fm34((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))) < ((d_676_v75_).f21), d_710_v87_, globalState)
                rhs127_ = d_707_v84_
                lhs104_ = d_708_v85_
                lhs105_ = default__.safeIndex(418, (d_708_v85_).length(0))
                d_617_v32_ = rhs126_
                lhs104_[lhs105_] = rhs127_
                index86_ = default__.safeIndex(182, (p1).length(0))
                (p1)[index86_] = p0
                index87_ = default__.safeIndex(182, (p1).length(0))
                (p1)[index87_] = not(p0)
                (globalState).f1 = default__.safeDivisionInt((len(d_617_v32_)) * ((0) - ((self).f11)), (48) * ((d_676_v75_).f21))
                d_711_v88_: _dafny.Map
                d_711_v88_ = _dafny.Map({(p1)[default__.safeIndex(182, (p1).length(0))]: p0})
                d_712_v89_: _dafny.Array
                nw103_ = _dafny.Array(None, 23)
                nw103_[int(0)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(1)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(2)] = False
                nw103_[int(3)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(4)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(5)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(6)] = p0
                nw103_[int(7)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(8)] = p0
                nw103_[int(9)] = not (p0) or (p0)
                nw103_[int(10)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(11)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(12)] = (p0 if (p1)[default__.safeIndex(182, (p1).length(0))] else not(p0))
                nw103_[int(13)] = (d_676_v75_).fm21((p1)[default__.safeIndex(182, (p1).length(0))], (d_676_v75_).f21, globalState)
                nw103_[int(14)] = (p1)[default__.safeIndex(182, (p1).length(0))]
                nw103_[int(15)] = False
                nw103_[int(16)] = False
                nw103_[int(17)] = (p0) and ((p1)[default__.safeIndex(182, (p1).length(0))])
                nw103_[int(18)] = (True if p0 else (p1)[default__.safeIndex(182, (p1).length(0))])
                nw103_[int(19)] = p0
                nw103_[int(20)] = (d_676_v75_).fm21(p0, (d_710_v87_).cf14, globalState)
                nw103_[int(21)] = ((d_711_v88_)[p0] if (p0) in (d_711_v88_) else not((p1)[default__.safeIndex(182, (p1).length(0))]))
                nw103_[int(22)] = p0
                d_712_v89_ = nw103_
                r1 = d_712_v89_
            elif True:
                d_713___mcc_h11_ = source13_.cf5
                d_714_cf5_ = d_713___mcc_h11_
                d_715_v90_: _dafny.Seq
                d_715_v90_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, True]), _dafny.SeqWithoutIsStrInference([p0])])
                d_716_v91_: _dafny.Map
                d_716_v91_ = _dafny.Map({p0: (self).f24})
                d_717_v92_: C0
                nw104_ = C0()
                nw104_.ctor__(((d_715_v90_)[default__.safeIndex((self).f11, len(d_715_v90_))]).set(default__.safeIndex(len(d_716_v91_), len((d_715_v90_)[default__.safeIndex((self).f11, len(d_715_v90_))])), p0), (self).f11)
                d_717_v92_ = nw104_
                d_718_v93_: _dafny.Map
                d_718_v93_ = _dafny.Map({p0: default__.fm1(globalState)})
                d_719_v94_: _dafny.Seq
                d_719_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxgefovt"))
                d_720_v95_: _dafny.Map
                d_720_v95_ = _dafny.Map({((d_718_v93_)[p0] if (p0) in (d_718_v93_) else d_719_v94_): (d_676_v75_).f21})
                d_721_v96_: _dafny.Map
                d_721_v96_ = _dafny.Map({p0: (d_676_v75_).f21})
                d_720_v95_ = (d_720_v95_).set((d_719_v94_) + (d_719_v94_), len((d_721_v96_) | (_dafny.Map({p0: (d_676_v75_).f21}))))
                d_722_v97_: _dafny.MultiSet
                d_722_v97_ = _dafny.MultiSet([p0, p0])
                (globalState).f1 = (default__.safeDivisionInt((d_676_v75_).f21, (self).f11)) * ((((d_722_v97_)[p0] if (p0) in (d_722_v97_) else (d_676_v75_).f21)) * ((d_717_v92_).f21))
                d_723_v98_: C0
                nw105_ = C0()
                nw105_.ctor__((_dafny.SeqWithoutIsStrInference([p0])) + (((d_676_v75_).f20).set(default__.safeIndex((d_676_v75_).f21, len((d_676_v75_).f20)), p0)), len((d_721_v96_) | (d_721_v96_)))
                d_723_v98_ = nw105_
            (globalState).f1 = (0) - ((d_676_v75_).f21)
            d_724_v99_: _dafny.Seq
            d_724_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uesixv"))
            rhs128_ = (d_676_v75_).f21
            rhs129_ = d_724_v99_
            lhs106_ = globalState
            lhs106_.f1 = rhs128_
            d_724_v99_ = rhs129_
            (globalState).f1 = ((self).f11) + ((self).f11)
            rhs130_ = ((d_676_v75_).f21) <= (((d_676_v75_).f21) * ((d_676_v75_).f21))
            rhs131_ = not(p0)
            lhs107_ = globalState
            lhs108_ = globalState
            lhs107_.f7 = rhs130_
            lhs108_.f6 = rhs131_
        elif True:
            (globalState).f7 = p0
            d_725_v100_: _dafny.Array
            nw106_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_725_v100_ = nw106_
            d_726_v101_: _dafny.Array
            nw107_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_726_v101_ = nw107_
            d_727_v102_: _dafny.Array
            d_727_v102_ = d_726_v101_
            index88_ = default__.safeIndex(693, (d_725_v100_).length(0))
            (d_725_v100_)[index88_] = d_727_v102_
            index89_ = default__.safeIndex(693, (d_725_v100_).length(0))
            (d_725_v100_)[index89_] = d_727_v102_
            d_728_v103_: C0
            nw108_ = C0()
            nw108_.ctor__((d_675_v74_) + (d_675_v74_), (d_676_v75_).f21)
            d_728_v103_ = nw108_
            d_729_v104_: D7
            d_729_v104_ = D7_DC12((d_617_v32_).set(default__.safeIndex((d_728_v103_).f21, len(d_617_v32_)), (d_728_v103_).f21))
            source14_ = d_729_v104_
            if source14_.is_DC13:
                d_730___mcc_h12_ = source14_.cf18
                d_731___mcc_h13_ = source14_.cf19
                d_732___mcc_h14_ = source14_.cf20
                d_733_cf20_ = d_732___mcc_h14_
                d_734_cf19_ = d_731___mcc_h13_
                d_735_cf18_ = d_730___mcc_h12_
                d_736_v105_: bool
                d_737_v106_: bool
                d_738_v107_: bool
                d_739_v108_: int
                out20_: bool
                out21_: bool
                out22_: bool
                out23_: int
                out20_, out21_, out22_, out23_ = (self).m8(globalState)
                d_736_v105_ = out20_
                d_737_v106_ = out21_
                d_738_v107_ = out22_
                d_739_v108_ = out23_
                d_740_v109_: _dafny.Seq
                d_740_v109_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbcbg"))
                rhs132_ = _dafny.SeqWithoutIsStrInference([(d_740_v109_) <= (d_740_v109_)])
                rhs133_ = p1
                d_675_v74_ = rhs132_
                r1 = rhs133_
                (globalState).f6 = p0
                d_738_v107_ = not(d_736_v105_)
            elif source14_.is_DC12:
                d_741___mcc_h15_ = source14_.cf17
                d_742_cf17_ = d_741___mcc_h15_
                (globalState).f7 = not(p0)
                d_743_v110_: C0
                nw109_ = C0()
                nw109_.ctor__((d_728_v103_).f20, (d_676_v75_).f21)
                d_743_v110_ = nw109_
                d_744_v111_: _dafny.Seq
                d_744_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckt"))
                d_744_v111_ = d_744_v111_
                d_745_v112_: _dafny.Seq
                d_745_v112_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({True: (self).f24})])
                d_746_v113_: _dafny.Seq
                d_746_v113_ = d_744_v111_
                d_747_v114_: _dafny.Map
                d_747_v114_ = _dafny.Map({(d_745_v112_)[default__.safeIndex(default__.fm2((d_728_v103_).f21, len(d_744_v111_), globalState), len(d_745_v112_))]: d_746_v113_})
                d_748_v115_: _dafny.Map
                d_748_v115_ = _dafny.Map({False: (self).f24})
                d_747_v114_ = (d_747_v114_).set(d_748_v115_, d_746_v113_)
            elif True:
                d_749___mcc_h16_ = source14_.cf21
                d_750_cf21_ = d_749___mcc_h16_
                d_751_v116_: _dafny.Map
                d_751_v116_ = _dafny.Map({p0: p0})
                (globalState).f7 = (d_728_v103_).fm21(((d_751_v116_)[p0] if (p0) in (d_751_v116_) else p0), (d_676_v75_).f21, globalState)
                (globalState).f7 = p0
                d_752_v117_: _dafny.Set
                d_752_v117_ = _dafny.Set({324, (d_676_v75_).f21, ((d_676_v75_).f21) + ((d_728_v103_).f21)})
                rhs134_ = len(d_752_v117_)
                rhs135_ = (d_675_v74_) == (default__.fm35((d_728_v103_).f21, (d_676_v75_).f21, default__.fm36(p0, p0, (self).f11, globalState), 60, globalState))
                lhs109_ = globalState
                lhs110_ = globalState
                lhs109_.f1 = rhs134_
                lhs110_.f7 = rhs135_
                d_617_v32_ = d_617_v32_
            if (False if p0 else p0):
                def iife92_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in _dafny.IntegerRange(-190, 507):
                        d_754_v118_: int = compr_44_
                        if ((-190) <= (d_754_v118_)) and ((d_754_v118_) < (507)):
                            coll44_[(d_754_v118_) - ((d_676_v75_).f21)] = (0) - ((d_676_v75_).f21)
                    return _dafny.Map(coll44_)
                d_677_v76_ = (d_677_v76_).set(len((_dafny.SeqWithoutIsStrInference([(d_728_v103_).f21 for d_753_i10_ in range(default__.abs(535))])) + (_dafny.SeqWithoutIsStrInference([len(iife92_()
) for d_755_i11_ in range(default__.abs(85))]))), p0)
                (globalState).f1 = (d_728_v103_).f21
                d_756_v119_: _dafny.Seq
                d_756_v119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                d_757_v120_: _dafny.Seq
                d_757_v120_ = _dafny.SeqWithoutIsStrInference([d_756_v119_])
                (globalState).f7 = ((d_756_v119_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lviqvc")))) < ((d_756_v119_) + ((d_757_v120_)[default__.safeIndex((d_676_v75_).f21, len(d_757_v120_))]))
                (globalState).f1 = len(default__.fm1(globalState))
                d_758_v121_: _dafny.MultiSet
                d_758_v121_ = _dafny.MultiSet([(d_756_v119_) < (d_756_v119_), ((self).f24) not in (d_756_v119_)])
                rhs136_ = ((self).fm3(983, True, p0, globalState)) - ((d_676_v75_).f21)
                rhs137_ = ((default__.fm37(globalState) if p0 else _dafny.MultiSet((d_676_v75_).f20))).set((p0 if False else p0), default__.abs((d_728_v103_).f21))
                rhs138_ = not(p0)
                lhs111_ = globalState
                lhs112_ = globalState
                lhs111_.f1 = rhs136_
                d_758_v121_ = rhs137_
                lhs112_.f7 = rhs138_
            elif True:
                (globalState).f6 = p0
                (globalState).f1 = ((0) - ((self).f11) if p0 else (d_728_v103_).f21)
                d_759_v122_: _dafny.Array
                nw110_ = _dafny.Array(int(0), 20)
                d_759_v122_ = nw110_
                d_759_v122_ = d_759_v122_
                d_760_v123_: _dafny.Array
                def lambda24_(d_761_i12_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pswwmeq"))

                init12_ = lambda24_
                nw111_ = _dafny.Array(None, 7)
                for i0_12_ in range(nw111_.length(0)):
                    nw111_[i0_12_] = init12_(i0_12_)
                d_760_v123_ = nw111_
                rhs139_ = d_760_v123_
                rhs140_ = (d_728_v103_).f21
                lhs113_ = globalState
                d_760_v123_ = rhs139_
                lhs113_.f1 = rhs140_
                d_762_v124_: _dafny.MultiSet
                d_762_v124_ = _dafny.MultiSet([d_675_v74_])
                index90_ = default__.safeIndex(874, (d_759_v122_).length(0))
                (d_759_v122_)[index90_] = ((d_762_v124_ if p0 else d_762_v124_)).cardinality
                index91_ = default__.safeIndex(874, (d_759_v122_).length(0))
                (d_759_v122_)[index91_] = default__.fm2((0) - ((self).fm3((d_676_v75_).f21, not(p0), p0, globalState)), (d_676_v75_).f21, globalState)
        d_763_v125_: D6
        d_763_v125_ = D6_DC11(85, (d_676_v75_).f21, (d_676_v75_).f21)
        if (default__.fm34(p0, d_763_v125_, globalState)) == (_dafny.SeqWithoutIsStrInference([(d_676_v75_).f21, (self).f11])):
            (globalState).f1 = (d_676_v75_).f21
            if False:
                d_764_v126_: _dafny.Seq
                d_764_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkdpcyeue"))
                d_765_v127_: _dafny.Map
                d_765_v127_ = _dafny.Map({p0: d_764_v126_})
                (globalState).f1 = (default__.fm38(d_765_v127_, p0, False, globalState)).cf0
                d_766_v128_: _dafny.Array
                def lambda25_(d_767_v126_, d_768_p0_):
                    def lambda26_(d_769_i13_):
                        return (d_769_i13_) * (len(_dafny.Map({len(d_767_v126_): d_768_p0_})))

                    return lambda26_

                init13_ = lambda25_(d_764_v126_, p0)
                nw112_ = _dafny.Array(None, 2)
                for i0_13_ in range(nw112_.length(0)):
                    nw112_[i0_13_] = init13_(i0_13_)
                d_766_v128_ = nw112_
                index92_ = default__.safeIndex(136, (d_766_v128_).length(0))
                (d_766_v128_)[index92_] = (d_676_v75_).f21
                index93_ = default__.safeIndex(474, (d_766_v128_).length(0))
                (d_766_v128_)[index93_] = (d_676_v75_).f21
                index94_ = default__.safeIndex(136, (d_766_v128_).length(0))
                index95_ = default__.safeIndex(474, (d_766_v128_).length(0))
                rhs141_ = (d_676_v75_).f21
                rhs142_ = ((d_676_v75_).f21) * ((d_676_v75_).f21)
                lhs114_ = d_766_v128_
                lhs115_ = default__.safeIndex(136, (d_766_v128_).length(0))
                lhs116_ = d_766_v128_
                lhs117_ = default__.safeIndex(474, (d_766_v128_).length(0))
                lhs114_[lhs115_] = rhs141_
                lhs116_[lhs117_] = rhs142_
                rhs143_ = (0) - ((0) - ((d_766_v128_)[default__.safeIndex(136, (d_766_v128_).length(0))]))
                rhs144_ = p0
                rhs145_ = (d_617_v32_) <= (d_617_v32_)
                rhs146_ = (p0 if not(False) else True)
                lhs118_ = globalState
                lhs119_ = globalState
                lhs120_ = globalState
                lhs121_ = globalState
                lhs118_.f1 = rhs143_
                lhs119_.f7 = rhs144_
                lhs120_.f7 = rhs145_
                lhs121_.f6 = rhs146_
                (globalState).f6 = p0
                (globalState).f1 = (0) - (default__.safeDivisionInt((d_676_v75_).f21, len((_dafny.SeqWithoutIsStrInference([(self).f24 for d_770_i14_ in range(default__.abs(35))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))))))
            elif True:
                d_771_v129_: bool
                d_772_v130_: bool
                d_773_v131_: bool
                d_774_v132_: int
                out24_: bool
                out25_: bool
                out26_: bool
                out27_: int
                out24_, out25_, out26_, out27_ = (self).m8(globalState)
                d_771_v129_ = out24_
                d_772_v130_ = out25_
                d_773_v131_ = out26_
                d_774_v132_ = out27_
                (globalState).f6 = not((d_676_v75_).fm21(False, default__.safeDivisionInt(d_774_v132_, (d_676_v75_).f21), globalState))
                r1 = p1
                d_775_v133_: D11
                d_775_v133_ = D11_DC23(d_772_v130_, (self).f11, (d_676_v75_).f21)
                (globalState).f7 = (d_775_v133_).cf31
                (globalState).f1 = (0) - ((default__.safeDivisionInt((d_676_v75_).f21, (0) - ((d_676_v75_).f21))) + ((self).f11))
            d_776_v134_: _dafny.MultiSet
            d_776_v134_ = _dafny.MultiSet([(self).f24, (self).f24])
            d_777_v135_: _dafny.Seq
            d_777_v135_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vds"))
            d_677_v76_ = (d_677_v76_).set(((d_776_v134_)[(self).f24] if ((self).f24) in (d_776_v134_) else (self).fm3((d_676_v75_).f21, p0, p0, globalState)), (d_777_v135_) != (d_777_v135_))
            d_778_v136_: C0
            nw113_ = C0()
            nw113_.ctor__(d_675_v74_, (d_676_v75_).f21)
            d_778_v136_ = nw113_
            d_779_v137_: C0
            nw114_ = C0()
            nw114_.ctor__(((d_676_v75_).f20 if p0 else d_675_v74_), (d_778_v136_).f21)
            d_779_v137_ = nw114_
        elif True:
            d_780_v138_: C0
            nw115_ = C0()
            nw115_.ctor__((d_676_v75_).f20, default__.safeDivisionInt(default__.fm2((self).f11, len(_dafny.SeqWithoutIsStrInference([(d_676_v75_).f21 for d_781_i15_ in range(default__.abs(451))])), globalState), ((d_616_v31_)[(self).f11] if ((self).f11) in (d_616_v31_) else (d_676_v75_).f21)))
            d_780_v138_ = nw115_
            (globalState).f6 = ((d_676_v75_).fm21(p0, (d_676_v75_).f21, globalState) if p0 else p0)
            d_782_v139_: C0
            nw116_ = C0()
            nw116_.ctor__((d_676_v75_).f20, (d_676_v75_).f21)
            d_782_v139_ = nw116_
            (globalState).f1 = (d_780_v138_).f21
            d_783_v140_: _dafny.Map
            d_783_v140_ = _dafny.Map({((d_780_v138_).f20)[default__.safeIndex((0) - ((self).fm3((0) - ((d_782_v139_).f21), p0, p0, globalState)), len((d_780_v138_).f20))]: p0})
            d_784_v141_: _dafny.MultiSet
            d_784_v141_ = _dafny.MultiSet([p0])
            d_785_v142_: _dafny.Seq
            d_785_v142_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pag"))
            d_786_v143_: _dafny.Array
            nw117_ = _dafny.Array(None, 23)
            nw117_[int(0)] = (d_780_v138_).f21
            nw117_[int(1)] = (d_782_v139_).f21
            nw117_[int(2)] = len((d_780_v138_).f20)
            nw117_[int(3)] = (d_676_v75_).f21
            nw117_[int(4)] = -913
            nw117_[int(5)] = (d_676_v75_).f21
            nw117_[int(6)] = (d_782_v139_).f21
            nw117_[int(7)] = (d_676_v75_).f21
            nw117_[int(8)] = (d_780_v138_).f21
            nw117_[int(9)] = (d_780_v138_).f21
            nw117_[int(10)] = (d_676_v75_).f21
            nw117_[int(11)] = (d_616_v31_).cardinality
            nw117_[int(12)] = (d_780_v138_).f21
            nw117_[int(13)] = (self).f11
            nw117_[int(14)] = (d_780_v138_).f21
            nw117_[int(15)] = (d_676_v75_).f21
            nw117_[int(16)] = len(d_785_v142_)
            nw117_[int(17)] = len(d_617_v32_)
            nw117_[int(18)] = len(d_677_v76_)
            nw117_[int(19)] = len(d_785_v142_)
            nw117_[int(20)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sowrbha")))
            nw117_[int(21)] = (d_676_v75_).f21
            nw117_[int(22)] = (self).f11
            d_786_v143_ = nw117_
            d_787_v144_: _dafny.Seq
            d_787_v144_ = _dafny.SeqWithoutIsStrInference([d_786_v143_])
            d_788_v145_: _dafny.Map
            d_788_v145_ = _dafny.Map({d_617_v32_: p0})
            d_789_v146_: _dafny.Map
            d_789_v146_ = _dafny.Map({((d_616_v31_)[26] if (26) in (d_616_v31_) else (d_676_v75_).f21): default__.fm39(len(d_785_v142_), p0, len(d_617_v32_), globalState)})
            d_790_v147_: _dafny.Array
            nw118_ = _dafny.Array(None, 27)
            nw118_[int(0)] = 660
            nw118_[int(1)] = (len(_dafny.Map({p0: d_783_v140_})) if True else (d_782_v139_).f21)
            nw118_[int(2)] = len(d_675_v74_)
            nw118_[int(3)] = (d_676_v75_).f21
            nw118_[int(4)] = (d_780_v138_).f21
            nw118_[int(5)] = (d_782_v139_).f21
            nw118_[int(6)] = ((d_780_v138_).f21) * ((d_676_v75_).f21)
            nw118_[int(7)] = (d_676_v75_).f21
            nw118_[int(8)] = default__.fm2((self).f11, (d_676_v75_).f21, globalState)
            nw118_[int(9)] = ((d_780_v138_).f21) - ((d_784_v141_).cardinality)
            nw118_[int(10)] = len(d_787_v144_)
            nw118_[int(11)] = (self).f11
            nw118_[int(12)] = (0) - ((d_676_v75_).f21)
            nw118_[int(13)] = (self).f11
            nw118_[int(14)] = 98
            nw118_[int(15)] = (d_676_v75_).f21
            nw118_[int(16)] = len(d_788_v145_)
            nw118_[int(17)] = ((d_676_v75_).f21) + ((d_676_v75_).f21)
            nw118_[int(18)] = default__.safeModuloInt((d_780_v138_).f21, ((_dafny.MultiSet([p0])).set(p0, default__.abs((d_676_v75_).f21))).cardinality)
            nw118_[int(19)] = (self).f11
            nw118_[int(20)] = (d_782_v139_).f21
            nw118_[int(21)] = len(d_789_v146_)
            nw118_[int(22)] = default__.safeModuloInt(983, len(d_785_v142_))
            nw118_[int(23)] = (d_780_v138_).f21
            nw118_[int(24)] = (d_782_v139_).f21
            nw118_[int(25)] = (self).f11
            nw118_[int(26)] = ((d_780_v138_).f21) * ((d_780_v138_).f21)
            d_790_v147_ = nw118_
            d_791_v148_: _dafny.Map
            d_791_v148_ = _dafny.Map({p1: d_786_v143_})
            d_792_v149_: _dafny.Map
            d_792_v149_ = _dafny.Map({len(default__.fm40((d_780_v138_).f20, (self).f24, (self).f11, globalState)): d_790_v147_})
            d_786_v143_ = ((d_791_v148_)[p1] if (p1) in (d_791_v148_) else ((d_792_v149_)[(d_782_v139_).f21] if ((d_782_v139_).f21) in (d_792_v149_) else d_786_v143_))
        d_793_v150_: _dafny.Seq
        d_793_v150_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbi"))
        d_794_v151_: C1
        nw119_ = C1()
        nw119_.ctor__((_dafny.SeqWithoutIsStrInference([(self).f24 for d_795_i16_ in range(default__.abs(-580))])) <= (d_793_v150_), (p0) and (p0))
        d_794_v151_ = nw119_
        d_796_v152_: _dafny.Set
        d_796_v152_ = _dafny.Set({p0})
        d_797_v153_: D7
        d_797_v153_ = D7_DC13((d_796_v152_) | (_dafny.Set({p0})), (d_616_v31_).cardinality, 263)
        r0 = d_797_v153_
        r1 = p1
        return r0, r1

    @property
    def f24(self):
        return self._f24

class C3(T0, T1):
    def  __init__(self):
        self._f12: _dafny.Set = _dafny.Set({})
        self._f11: int = int(0)
        self.f23: int = int(0)
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f12(self):
        return self._f12
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f22, f23, f11, f12):
        (self)._f22 = f22
        (self).f23 = f23
        (self)._f11 = f11
        (self)._f12 = f12

    def fm3(self, p0, p1, p2, globalState):
        return ((self).f11) + ((self).f11)

    def fm4(self, p0, p1, p2, p3, globalState):
        return not(not(not(not((self.f23) <= ((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True, True, False])) + (_dafny.SeqWithoutIsStrInference([not(not(True)), True, False, True])))).cardinality)))))

    def fm5(self, p0, p1, p2, p3, globalState):
        source15_ = D7_DC14(D7_DC14(D7_DC12(_dafny.SeqWithoutIsStrInference([(self).f11 for d_798_i0_ in range(default__.abs(-233))]))))
        if source15_.is_DC13:
            d_799___mcc_h0_ = source15_.cf18
            d_800___mcc_h1_ = source15_.cf19
            d_801___mcc_h2_ = source15_.cf20
            d_802_cf20_ = d_801___mcc_h2_
            d_803_cf19_ = d_800___mcc_h1_
            d_804_cf18_ = d_799___mcc_h0_
            return not(True)
        elif source15_.is_DC12:
            d_805___mcc_h3_ = source15_.cf17
            d_806_cf17_ = d_805___mcc_h3_
            return ((self).f11) != ((self).f11)
        elif True:
            d_807___mcc_h4_ = source15_.cf21
            d_808_cf21_ = d_807___mcc_h4_
            return (True) and (False)

    def fm6(self, p0, p1, globalState):
        return 694

    def fm25(self, p0, p1, p2, p3, globalState):
        return (_dafny.Map({not((D0_DC1(False, 444, True)).cf1): not(False)})) | ((_dafny.Map({False: True})) | ((D8_DC15(_dafny.Map({False: True}))).cf22))

    def m0(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_809_v0_: _dafny.Seq
        d_809_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njxqrukhg"))
        (globalState).f6 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xihlij"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_810_i0_ in range(default__.abs(598))]))) <= (d_809_v0_)
        d_811_v1_: bool
        d_811_v1_ = False
        if d_811_v1_:
            d_812_v2_: _dafny.Seq
            d_812_v2_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            rhs147_ = not(d_811_v1_)
            rhs148_ = ((d_812_v2_)[default__.safeIndex(len(d_809_v0_), len(d_812_v2_))] if False else (self).f11)
            rhs149_ = ((self).f11) <= (((self).f22) - (len(d_809_v0_)))
            lhs122_ = globalState
            lhs123_ = self
            lhs122_.f7 = rhs147_
            lhs123_.f23 = rhs148_
            d_811_v1_ = rhs149_
            d_813_v3_: _dafny.MultiSet
            d_813_v3_ = _dafny.MultiSet([d_811_v1_, d_811_v1_])
            d_814_v4_: _dafny.Seq
            d_814_v4_ = _dafny.SeqWithoutIsStrInference([True, False, d_811_v1_])
            d_815_v5_: D0
            d_815_v5_ = D0_DC0((self).f11)
            d_816_v6_: _dafny.Seq
            d_816_v6_ = _dafny.SeqWithoutIsStrInference([d_812_v2_])
            d_817_v7_: _dafny.Array
            nw120_ = _dafny.Array(int(0), 7)
            d_817_v7_ = nw120_
            d_818_v8_: _dafny.Map
            d_818_v8_ = _dafny.Map({d_817_v7_: d_811_v1_})
            d_819_v9_: _dafny.Array
            nw121_ = _dafny.Array(None, 21)
            nw121_[int(0)] = (d_813_v3_).issubset(d_813_v3_)
            nw121_[int(1)] = d_811_v1_
            nw121_[int(2)] = d_811_v1_
            nw121_[int(3)] = (self).fm5(D0_DC0(-65), d_814_v4_, 533, p0, globalState)
            nw121_[int(4)] = d_811_v1_
            nw121_[int(5)] = ((self).fm6(d_811_v1_, self.f23, globalState)) != ((self).f22)
            nw121_[int(6)] = (False if d_811_v1_ else not(d_811_v1_))
            nw121_[int(7)] = d_811_v1_
            nw121_[int(8)] = d_811_v1_
            nw121_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkcgskmp"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ju")))
            nw121_[int(10)] = (self).fm5(d_815_v5_, d_814_v4_, (self).f11, _dafny.CodePoint('o'), globalState)
            nw121_[int(11)] = (d_816_v6_) <= (_dafny.SeqWithoutIsStrInference([d_812_v2_ for d_820_i1_ in range(default__.abs(76))]))
            nw121_[int(12)] = ((self).f11) >= (self.f23)
            nw121_[int(13)] = (d_811_v1_) and (d_811_v1_)
            nw121_[int(14)] = (d_812_v2_) < ((d_812_v2_).set(default__.safeIndex((self).f11, len(d_812_v2_)), len((d_816_v6_)[default__.safeIndex(self.f23, len(d_816_v6_))])))
            nw121_[int(15)] = not (((d_818_v8_)[d_817_v7_] if (d_817_v7_) in (d_818_v8_) else d_811_v1_)) or (d_811_v1_)
            nw121_[int(16)] = (d_814_v4_)[default__.safeIndex((self).f22, len(d_814_v4_))]
            nw121_[int(17)] = not (False) or (True)
            nw121_[int(18)] = ((self).f22) >= (self.f23)
            nw121_[int(19)] = d_811_v1_
            nw121_[int(20)] = d_811_v1_
            d_819_v9_ = nw121_
            d_821_v10_: _dafny.MultiSet
            d_821_v10_ = _dafny.MultiSet([145, len((self).f12)])
            d_822_v11_: _dafny.Map
            d_822_v11_ = _dafny.Map({d_811_v1_: (self).f11})
            d_823_v12_: _dafny.Map
            d_823_v12_ = _dafny.Map({len(d_822_v11_): len(d_809_v0_)})
            d_824_v13_: _dafny.Map
            d_824_v13_ = _dafny.Map({d_821_v10_: d_823_v12_})
            index96_ = default__.safeIndex(939, (d_819_v9_).length(0))
            (d_819_v9_)[index96_] = (self).fm4(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxsuipcwe"))), d_824_v13_, self.f23, d_811_v1_, globalState)
            index97_ = default__.safeIndex(939, (d_819_v9_).length(0))
            (d_819_v9_)[index97_] = d_811_v1_
            d_825_v14_: C0
            nw122_ = C0()
            nw122_.ctor__(d_814_v4_, ((d_822_v11_)[(d_819_v9_)[default__.safeIndex(939, (d_819_v9_).length(0))]] if ((d_819_v9_)[default__.safeIndex(939, (d_819_v9_).length(0))]) in (d_822_v11_) else (0) - ((self).f22)))
            d_825_v14_ = nw122_
            d_825_v14_ = d_825_v14_
            d_826_v15_: _dafny.Seq
            d_826_v15_ = d_809_v0_
            rhs150_ = d_811_v1_
            rhs151_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nb"))
            lhs124_ = globalState
            lhs124_.f7 = rhs150_
            d_826_v15_ = rhs151_
            d_827_v16_: _dafny.Map
            d_827_v16_ = _dafny.Map({d_813_v3_: default__.fm26(globalState)})
            d_822_v11_ = (d_822_v11_).set(((self).f11) > (len(d_822_v11_)), len(d_827_v16_))
        elif True:
            d_828_v17_: _dafny.Seq
            d_828_v17_ = _dafny.SeqWithoutIsStrInference([self.f23])
            d_829_v18_: D8
            d_829_v18_ = D8_DC16(d_811_v1_, (0) - (self.f23))
            d_830_v19_: D8
            d_830_v19_ = D8_DC17(d_829_v18_)
            d_831_v20_: _dafny.Map
            d_831_v20_ = _dafny.Map({d_828_v17_: d_829_v18_})
            d_832_v21_: D8
            d_832_v21_ = D8_DC17(((d_831_v20_)[_dafny.SeqWithoutIsStrInference([self.f23])] if (_dafny.SeqWithoutIsStrInference([self.f23])) in (d_831_v20_) else d_829_v18_))
            d_832_v21_ = d_832_v21_
            d_833_v22_: D0
            d_833_v22_ = D0_DC0(self.f23)
            d_834_v23_: _dafny.Seq
            d_834_v23_ = _dafny.SeqWithoutIsStrInference([d_811_v1_])
            d_835_v24_: _dafny.MultiSet
            d_835_v24_ = _dafny.MultiSet([(self).f11])
            d_836_v25_: _dafny.MultiSet
            d_836_v25_ = _dafny.MultiSet([d_835_v24_])
            d_837_v26_: _dafny.Seq
            d_837_v26_ = _dafny.SeqWithoutIsStrInference([d_834_v23_])
            d_838_v27_: _dafny.Array
            nw123_ = _dafny.Array(None, 18)
            nw123_[int(0)] = not (d_811_v1_) or (d_811_v1_)
            nw123_[int(1)] = not(d_811_v1_)
            nw123_[int(2)] = d_811_v1_
            nw123_[int(3)] = True
            nw123_[int(4)] = d_811_v1_
            nw123_[int(5)] = (self).fm5(d_833_v22_, d_834_v23_, self.f23, p0, globalState)
            nw123_[int(6)] = (self.f23) > (self.f23)
            nw123_[int(7)] = not(not(d_811_v1_))
            nw123_[int(8)] = d_811_v1_
            nw123_[int(9)] = d_811_v1_
            nw123_[int(10)] = not((_dafny.Set({(d_835_v24_).cardinality})).ispropersubset(_dafny.Set({(self).f11})))
            nw123_[int(11)] = (d_836_v25_).issubset(d_836_v25_)
            nw123_[int(12)] = d_811_v1_
            nw123_[int(13)] = not(d_811_v1_)
            nw123_[int(14)] = d_811_v1_
            nw123_[int(15)] = d_811_v1_
            nw123_[int(16)] = ((0) - (len((d_837_v26_)[default__.safeIndex(self.f23, len(d_837_v26_))]))) == (self.f23)
            nw123_[int(17)] = d_811_v1_
            d_838_v27_ = nw123_
            index98_ = default__.safeIndex(313, (d_838_v27_).length(0))
            (d_838_v27_)[index98_] = d_811_v1_
            index99_ = default__.safeIndex(313, (d_838_v27_).length(0))
            (d_838_v27_)[index99_] = d_811_v1_
            d_839_v28_: _dafny.Map
            d_839_v28_ = _dafny.Map({d_838_v27_: (0) - ((self.f23) + (732))})
            d_839_v28_ = (d_839_v28_).set(d_838_v27_, (len(d_809_v0_)) - ((self).f22))
            d_811_v1_ = True
            d_840_v29_: _dafny.Array
            nw124_ = _dafny.Array(int(0), 26)
            d_840_v29_ = nw124_
            d_841_v30_: _dafny.Set
            d_841_v30_ = _dafny.Set({(self).f22, self.f23, len(_dafny.Set({(self).f11})), self.f23})
            d_842_v31_: _dafny.Map
            d_842_v31_ = _dafny.Map({491: len(d_841_v30_)})
            d_843_v32_: _dafny.Map
            d_843_v32_ = _dafny.Map({d_840_v29_: d_842_v31_})
            d_843_v32_ = (d_843_v32_).set((d_840_v29_ if (d_838_v27_)[default__.safeIndex(313, (d_838_v27_).length(0))] else d_840_v29_), d_842_v31_)
        d_844_i2_: int
        d_844_i2_ = 0
        with _dafny.label("7"):
            while False:
                with _dafny.c_label("7"):
                    if (d_844_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_844_i2_ = (d_844_i2_) + (1)
                    d_845_v33_: C1
                    nw125_ = C1()
                    nw125_.ctor__(d_811_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdveoar"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csqsqpon"))))
                    d_845_v33_ = nw125_
                    (globalState).f6 = False
                    (globalState).f1 = (self).f22
                    (self).f23 = (0) - ((self).f11)
                    pass
            pass
        d_846_v34_: _dafny.Seq
        d_846_v34_ = _dafny.SeqWithoutIsStrInference([True, d_811_v1_])
        d_847_v35_: C0
        nw126_ = C0()
        nw126_.ctor__((d_846_v34_) + (d_846_v34_), (461) - ((self).f11))
        d_847_v35_ = nw126_
        d_848_v36_: _dafny.Set
        d_848_v36_ = _dafny.Set({342})
        d_849_v37_: _dafny.Array
        nw127_ = _dafny.Array(None, 16)
        nw127_[int(0)] = not((len(d_809_v0_)) == (self.f23))
        nw127_[int(1)] = d_811_v1_
        nw127_[int(2)] = (_dafny.Set({(self).f22})).ispropersubset(d_848_v36_)
        nw127_[int(3)] = False
        nw127_[int(4)] = not(d_811_v1_)
        nw127_[int(5)] = (len(_dafny.SeqWithoutIsStrInference([d_811_v1_]))) > (-62)
        nw127_[int(6)] = d_811_v1_
        nw127_[int(7)] = (d_811_v1_ if d_811_v1_ else d_811_v1_)
        nw127_[int(8)] = False
        nw127_[int(9)] = d_811_v1_
        nw127_[int(10)] = d_811_v1_
        nw127_[int(11)] = (d_811_v1_) and (d_811_v1_)
        nw127_[int(12)] = d_811_v1_
        nw127_[int(13)] = d_811_v1_
        nw127_[int(14)] = d_811_v1_
        nw127_[int(15)] = d_811_v1_
        d_849_v37_ = nw127_
        index100_ = default__.safeIndex(765, (d_849_v37_).length(0))
        (d_849_v37_)[index100_] = d_811_v1_
        d_850_v38_: _dafny.Map
        d_850_v38_ = _dafny.Map({d_811_v1_: (d_847_v35_).f21})
        d_851_v39_: _dafny.MultiSet
        d_851_v39_ = _dafny.MultiSet([d_850_v38_])
        d_852_v40_: _dafny.Map
        d_852_v40_ = _dafny.Map({(d_847_v35_).f21: d_851_v39_})
        d_853_v41_: D9
        d_853_v41_ = D9_DC18(d_850_v38_)
        index101_ = default__.safeIndex(765, (d_849_v37_).length(0))
        (d_849_v37_)[index101_] = ((d_851_v39_) - (d_851_v39_)).issubset(((d_852_v40_)[self.f23] if (self.f23) in (d_852_v40_) else _dafny.MultiSet([d_850_v38_, d_850_v38_, (d_853_v41_).cf26])))
        d_854_v42_: _dafny.Map
        d_854_v42_ = _dafny.Map({(d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))]: (d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))]})
        d_855_i3_: int
        d_855_i3_ = 0
        with _dafny.label("8"):
            while ((d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))]) in (d_854_v42_):
                with _dafny.c_label("8"):
                    if (d_855_i3_) >= (100):
                        raise _dafny.Break("8")
                    d_855_i3_ = (d_855_i3_) + (1)
                    if not(d_811_v1_):
                        d_856_v43_: _dafny.MultiSet
                        d_856_v43_ = _dafny.MultiSet([(self).f22, (d_847_v35_).f21, (self).fm3((0) - ((self).f11), d_811_v1_, d_811_v1_, globalState)])
                        d_857_v44_: _dafny.Map
                        d_857_v44_ = _dafny.Map({(d_847_v35_).f21: _dafny.SeqWithoutIsStrInference([p0 for d_858_i4_ in range(default__.abs(-99))])})
                        (self).f23 = (((self).f22) - ((d_847_v35_).f21)) - (default__.safeModuloInt((d_856_v43_).cardinality, (self).fm6((d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))], len(d_857_v44_), globalState)))
                        (globalState).f7 = d_811_v1_
                        d_811_v1_ = (d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))]
                        (globalState).f1 = default__.safeModuloInt((d_847_v35_).f21, self.f23)
                        d_859_v45_: _dafny.MultiSet
                        d_859_v45_ = _dafny.MultiSet([(d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))], d_811_v1_])
                        d_860_v46_: _dafny.Map
                        d_860_v46_ = _dafny.Map({d_859_v45_: d_849_v37_})
                        d_860_v46_ = d_860_v46_
                    elif True:
                        r0 = _dafny.MultiSet([d_811_v1_, (d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))], (d_811_v1_ if (d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))] else d_811_v1_)])
                        d_861_v47_: _dafny.Map
                        d_861_v47_ = _dafny.Map({len(default__.fm26(globalState)): d_811_v1_})
                        d_862_v48_: str
                        d_862_v48_ = _dafny.CodePoint('d')
                        d_863_v49_: _dafny.Map
                        d_863_v49_ = _dafny.Map({_dafny.MultiSet(default__.fm27(d_811_v1_, d_811_v1_, globalState)): False})
                        rhs152_ = (d_863_v49_) == ((d_863_v49_) | (d_863_v49_))
                        rhs153_ = ((self).f11) - (self.f23)
                        rhs154_ = (d_809_v0_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhrhniii"))))
                        rhs155_ = (d_861_v47_) | (d_861_v47_)
                        rhs156_ = p0
                        lhs125_ = globalState
                        lhs126_ = globalState
                        lhs125_.f6 = rhs152_
                        lhs126_.f1 = rhs153_
                        d_809_v0_ = rhs154_
                        d_861_v47_ = rhs155_
                        d_862_v48_ = rhs156_
                        d_864_v50_: _dafny.MultiSet
                        d_864_v50_ = _dafny.MultiSet([False])
                        d_865_v51_: _dafny.Seq
                        d_865_v51_ = _dafny.SeqWithoutIsStrInference([(0) - (((d_864_v50_)[d_811_v1_] if (d_811_v1_) in (d_864_v50_) else self.f23))])
                        d_866_v52_: _dafny.MultiSet
                        d_866_v52_ = _dafny.MultiSet([(d_865_v51_)[default__.safeIndex((self).f22, len(d_865_v51_))], self.f23])
                        d_866_v52_ = d_866_v52_
                        d_867_v53_: _dafny.Map
                        d_867_v53_ = _dafny.Map({d_811_v1_: p0})
                        d_868_v54_: _dafny.Array
                        nw128_ = _dafny.Array(None, 19)
                        nw128_[int(0)] = (self).f11
                        nw128_[int(1)] = len(d_867_v53_)
                        nw128_[int(2)] = self.f23
                        nw128_[int(3)] = 522
                        nw128_[int(4)] = len(d_865_v51_)
                        nw128_[int(5)] = 845
                        nw128_[int(6)] = (self).f11
                        nw128_[int(7)] = (d_847_v35_).f21
                        nw128_[int(8)] = (d_847_v35_).f21
                        nw128_[int(9)] = (self).f22
                        nw128_[int(10)] = (d_847_v35_).f21
                        nw128_[int(11)] = len(d_854_v42_)
                        nw128_[int(12)] = (_dafny.MultiSet([(d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))], d_811_v1_])).cardinality
                        nw128_[int(13)] = self.f23
                        nw128_[int(14)] = len(d_809_v0_)
                        nw128_[int(15)] = (self).f22
                        nw128_[int(16)] = (0) - ((self).f22)
                        nw128_[int(17)] = 138
                        nw128_[int(18)] = (d_847_v35_).f21
                        d_868_v54_ = nw128_
                        d_869_v55_: _dafny.Set
                        d_869_v55_ = _dafny.Set({d_868_v54_})
                        d_869_v55_ = (d_869_v55_).intersection((d_869_v55_) | (d_869_v55_))
                        d_870_v56_: _dafny.Seq
                        d_870_v56_ = d_809_v0_
                        d_871_v57_: _dafny.Seq
                        d_871_v57_ = _dafny.SeqWithoutIsStrInference([d_809_v0_])
                        (globalState).f7 = ((d_870_v56_)) in ((d_871_v57_).set(default__.safeIndex(62, len(d_871_v57_)), _dafny.SeqWithoutIsStrInference([d_862_v48_ for d_872_i5_ in range(default__.abs(537))])))
                    (globalState).f6 = (d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))]
                    d_873_v58_: str
                    d_873_v58_ = _dafny.CodePoint('w')
                    d_873_v58_ = p0
                    d_874_v59_: _dafny.Seq
                    d_874_v59_ = _dafny.SeqWithoutIsStrInference([(d_847_v35_).f21])
                    d_875_v60_: T0
                    nw129_ = C2()
                    nw129_.ctor__(d_873_v58_, len(d_874_v59_))
                    d_875_v60_ = nw129_
                    d_876_v61_: _dafny.Map
                    d_876_v61_ = _dafny.Map({d_875_v60_: d_849_v37_})
                    d_877_v62_: D0
                    d_877_v62_ = D0_DC0(len(d_876_v61_))
                    (globalState).f6 = (self).fm5(d_877_v62_, (_dafny.SeqWithoutIsStrInference([(d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))]])) + ((d_847_v35_).f20), ((self).fm6((d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))], (self).f11, globalState)) * (self.f23), _dafny.CodePoint('e'), globalState)
                    pass
            pass
        d_878_v63_: _dafny.MultiSet
        d_878_v63_ = _dafny.MultiSet([d_811_v1_])
        r0 = ((d_878_v63_).intersection(_dafny.MultiSet([(d_849_v37_)[default__.safeIndex(765, (d_849_v37_).length(0))], d_811_v1_, d_811_v1_, d_811_v1_]))).set((d_811_v1_) and (d_811_v1_), default__.abs((self).f11))
        return r0

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        d_879_v0_: bool
        d_879_v0_ = True
        if d_879_v0_:
            d_880_v1_: str
            d_880_v1_ = _dafny.CodePoint('v')
            d_881_v2_: _dafny.Array
            nw130_ = _dafny.Array(None, 1)
            nw130_[int(0)] = (self).f22
            d_881_v2_ = nw130_
            d_882_v3_: C2
            nw131_ = C2()
            nw131_.ctor__(d_880_v1_, len(_dafny.Map({d_881_v2_: (self).f11})))
            d_882_v3_ = nw131_
            (globalState).f6 = d_879_v0_
            d_883_v4_: _dafny.Set
            d_883_v4_ = _dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('i')})
            (globalState).f1 = (0) - ((self).fm3(len(d_883_v4_), not(d_879_v0_), d_879_v0_, globalState))
            d_884_v5_: _dafny.Set
            d_884_v5_ = _dafny.Set({(self).f22, (283) * (p0), ((self).f22) - (565), p0})
            (globalState).f1 = len(d_884_v5_)
            d_885_v6_: _dafny.Set
            d_885_v6_ = _dafny.Set({d_879_v0_})
            d_885_v6_ = default__.fm41(globalState)
        elif True:
            d_886_v7_: _dafny.Seq
            d_886_v7_ = _dafny.SeqWithoutIsStrInference([d_879_v0_])
            d_887_v8_: _dafny.Seq
            d_887_v8_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_888_v9_: _dafny.Set
            d_888_v9_ = _dafny.Set({d_887_v8_})
            d_889_v10_: _dafny.Map
            d_889_v10_ = _dafny.Map({d_879_v0_: d_888_v9_})
            d_890_v11_: _dafny.Map
            d_890_v11_ = _dafny.Map({d_879_v0_: len(((d_889_v10_)[d_879_v0_] if (d_879_v0_) in (d_889_v10_) else d_888_v9_))})
            d_891_v12_: _dafny.Map
            d_891_v12_ = _dafny.Map({((d_890_v11_)[d_879_v0_] if (d_879_v0_) in (d_890_v11_) else (self).f11): d_886_v7_})
            d_892_v13_: C1
            nw132_ = C1()
            nw132_.ctor__((d_886_v7_) != (((d_891_v12_)[(self).f22] if ((self).f22) in (d_891_v12_) else d_886_v7_)), True)
            d_892_v13_ = nw132_
            (globalState).f7 = d_879_v0_
            d_893_v14_: _dafny.Seq
            d_893_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vea"))
            d_894_v15_: D7
            d_894_v15_ = D7_DC13((self).f12, (self).f11, p0)
            d_895_v16_: C0
            nw133_ = C0()
            nw133_.ctor__((d_886_v7_) + (default__.fm35(len(d_893_v14_), p0, d_894_v15_, (self).f11, globalState)), p0)
            d_895_v16_ = nw133_
            d_896_v17_: _dafny.Array
            def lambda27_(d_897_v0_):
                def lambda28_(d_898_i0_):
                    return d_897_v0_

                return lambda28_

            init14_ = lambda27_(d_879_v0_)
            nw134_ = _dafny.Array(None, 7)
            for i0_14_ in range(nw134_.length(0)):
                nw134_[i0_14_] = init14_(i0_14_)
            d_896_v17_ = nw134_
            rhs157_ = d_895_v16_
            rhs158_ = d_896_v17_
            rhs159_ = d_879_v0_
            rhs160_ = (d_879_v0_ if not(d_879_v0_) else ((d_892_v13_).f18 if (d_892_v13_).f19 else (d_892_v13_).f18))
            rhs161_ = not (default__.fm0(globalState)) or ((d_892_v13_).f19)
            lhs127_ = globalState
            lhs128_ = globalState
            d_895_v16_ = rhs157_
            d_896_v17_ = rhs158_
            d_879_v0_ = rhs159_
            lhs127_.f7 = rhs160_
            lhs128_.f7 = rhs161_
            (globalState).f1 = self.f23
            index102_ = default__.safeIndex(321, (d_896_v17_).length(0))
            (d_896_v17_)[index102_] = d_879_v0_
            index103_ = default__.safeIndex(321, (d_896_v17_).length(0))
            (d_896_v17_)[index103_] = d_879_v0_
        d_899_v18_: _dafny.Seq
        d_899_v18_ = _dafny.SeqWithoutIsStrInference([d_879_v0_])
        d_900_v19_: _dafny.Seq
        d_900_v19_ = _dafny.SeqWithoutIsStrInference([d_879_v0_, (d_899_v18_)[default__.safeIndex(self.f23, len(d_899_v18_))], d_879_v0_, not(d_879_v0_), d_879_v0_])
        d_901_i1_: int
        d_901_i1_ = 0
        with _dafny.label("9"):
            while (d_900_v19_) == (d_899_v18_):
                with _dafny.c_label("9"):
                    if (d_901_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_901_i1_ = (d_901_i1_) + (1)
                    d_902_v20_: str
                    d_902_v20_ = _dafny.CodePoint('u')
                    d_902_v20_ = d_902_v20_
                    d_903_v21_: D0
                    d_903_v21_ = D0_DC0((self).f11)
                    d_904_v22_: _dafny.MultiSet
                    d_904_v22_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bu"))])
                    d_905_v23_: _dafny.MultiSet
                    d_905_v23_ = _dafny.MultiSet([(self).fm5(d_903_v21_, d_900_v19_, (d_904_v22_).cardinality, d_902_v20_, globalState), default__.fm0(globalState), d_879_v0_, d_879_v0_])
                    d_906_v24_: _dafny.MultiSet
                    d_906_v24_ = d_905_v23_
                    source16_ = d_906_v24_
                    d_907___mcc_h0_ = source16_
                    d_908_cf11_ = d_907___mcc_h0_
                    rhs162_ = d_879_v0_
                    rhs163_ = not((d_879_v0_ if d_879_v0_ else not (d_879_v0_) or (False)))
                    lhs129_ = globalState
                    lhs130_ = globalState
                    lhs129_.f7 = rhs162_
                    lhs130_.f6 = rhs163_
                    (globalState).f6 = d_879_v0_
                    d_909_v25_: _dafny.Array
                    nw135_ = _dafny.Array(None, 8)
                    nw135_[int(0)] = d_906_v24_
                    nw135_[int(1)] = d_906_v24_
                    nw135_[int(2)] = d_906_v24_
                    nw135_[int(3)] = d_905_v23_
                    nw135_[int(4)] = d_906_v24_
                    nw135_[int(5)] = d_906_v24_
                    nw135_[int(6)] = d_906_v24_
                    nw135_[int(7)] = d_908_cf11_
                    d_909_v25_ = nw135_
                    index104_ = default__.safeIndex(282, (d_909_v25_).length(0))
                    (d_909_v25_)[index104_] = d_906_v24_
                    d_910_v26_: _dafny.Seq
                    d_910_v26_ = _dafny.SeqWithoutIsStrInference([(self).f22])
                    d_911_v27_: _dafny.Map
                    d_911_v27_ = _dafny.Map({(d_910_v26_)[default__.safeIndex(392, len(d_910_v26_))]: False})
                    d_912_v28_: _dafny.Seq
                    d_912_v28_ = _dafny.SeqWithoutIsStrInference([d_900_v19_, d_899_v18_])
                    d_913_v29_: _dafny.Set
                    d_913_v29_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_879_v0_]), (d_912_v28_)[default__.safeIndex((self).f11, len(d_912_v28_))]})
                    d_914_v31_: _dafny.MultiSet
                    d_914_v31_ = _dafny.MultiSet([p0, 496])
                    d_915_v33_: _dafny.Map
                    def iife93_():
                        coll45_ = _dafny.Map()
                        compr_45_: int
                        for compr_45_ in _dafny.IntegerRange(381, 912):
                            d_916_v32_: int = compr_45_
                            if ((381) <= (d_916_v32_)) and ((d_916_v32_) < (912)):
                                coll45_[(d_916_v32_) - ((self).f11)] = (self).f11
                        return _dafny.Map(coll45_)
                    d_915_v33_ = _dafny.Map({d_914_v31_: iife93_()
                    })
                    d_917_v34_: _dafny.Array
                    nw136_ = _dafny.Array(None, 10)
                    nw136_[int(0)] = d_879_v0_
                    nw136_[int(1)] = d_879_v0_
                    nw136_[int(2)] = d_879_v0_
                    nw136_[int(3)] = d_879_v0_
                    nw136_[int(4)] = d_879_v0_
                    nw136_[int(5)] = d_879_v0_
                    nw136_[int(6)] = d_879_v0_
                    nw136_[int(7)] = d_879_v0_
                    def iife94_():
                        coll46_ = _dafny.Set()
                        compr_46_: _dafny.Seq
                        for compr_46_ in (d_913_v29_).Elements:
                            d_918_v30_: _dafny.Seq = compr_46_
                            if (d_918_v30_) in (d_913_v29_):
                                coll46_ = coll46_.union(_dafny.Set([d_918_v30_]))
                        return _dafny.Set(coll46_)
                    def iife95_():
                        coll47_ = _dafny.Set()
                        compr_47_: _dafny.Seq
                        for compr_47_ in (d_913_v29_).Elements:
                            d_919_v30_: _dafny.Seq = compr_47_
                            if (d_919_v30_) in (d_913_v29_):
                                coll47_ = coll47_.union(_dafny.Set([d_919_v30_]))
                        return _dafny.Set(coll47_)
                    nw136_[int(8)] = ((d_911_v27_)[len(iife94_()
                    )] if (len(iife95_()
                    )) in (d_911_v27_) else (self).fm4(970, d_915_v33_, len(default__.fm26(globalState)), d_879_v0_, globalState))
                    nw136_[int(9)] = (d_914_v31_).ispropersubset(d_914_v31_)
                    d_917_v34_ = nw136_
                    index105_ = default__.safeIndex(172, (d_917_v34_).length(0))
                    (d_917_v34_)[index105_] = d_879_v0_
                    d_920_v35_: C0
                    nw137_ = C0()
                    nw137_.ctor__(d_900_v19_, (self).f22)
                    d_920_v35_ = nw137_
                    index106_ = default__.safeIndex(282, (d_909_v25_).length(0))
                    index107_ = default__.safeIndex(172, (d_917_v34_).length(0))
                    rhs164_ = (d_906_v24_ if True else d_908_cf11_)
                    rhs165_ = d_879_v0_
                    rhs166_ = d_920_v35_
                    lhs131_ = d_909_v25_
                    lhs132_ = default__.safeIndex(282, (d_909_v25_).length(0))
                    lhs133_ = d_917_v34_
                    lhs134_ = default__.safeIndex(172, (d_917_v34_).length(0))
                    lhs131_[lhs132_] = rhs164_
                    lhs133_[lhs134_] = rhs165_
                    d_920_v35_ = rhs166_
                    index108_ = default__.safeIndex(172, (d_917_v34_).length(0))
                    (d_917_v34_)[index108_] = True
                    d_921_v36_: _dafny.Map
                    d_921_v36_ = _dafny.Map({p0: default__.fm26(globalState)})
                    rhs167_ = True
                    rhs168_ = (0) - (((self).f11) - (p0))
                    rhs169_ = (d_921_v36_) != (d_921_v36_)
                    lhs135_ = self
                    lhs136_ = globalState
                    d_879_v0_ = rhs167_
                    lhs135_.f23 = rhs168_
                    lhs136_.f6 = rhs169_
                    d_922_v38_: _dafny.MultiSet
                    def iife96_():
                        coll48_ = _dafny.Map()
                        compr_48_: int
                        for compr_48_ in (_dafny.Set({p0})).Elements:
                            d_923_v37_: int = compr_48_
                            if (d_923_v37_) in (_dafny.Set({p0})):
                                coll48_[(d_923_v37_) - (len(_dafny.Map({d_879_v0_: len(_dafny.SeqWithoutIsStrInference([self.f23, (0) - ((self).f11)]))})))] = d_902_v20_
                        return _dafny.Map(coll48_)
                    d_922_v38_ = _dafny.MultiSet([iife96_()
                    ])
                    d_924_v39_: _dafny.Map
                    d_924_v39_ = _dafny.Map({(self).f22: d_902_v20_})
                    d_925_v40_: _dafny.Seq
                    d_925_v40_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_924_v39_, d_924_v39_]), d_922_v38_])
                    d_926_v41_: _dafny.Set
                    d_926_v41_ = _dafny.Set({d_899_v18_})
                    if not((((d_925_v40_)[default__.safeIndex(len(d_926_v41_), len(d_925_v40_))] if d_879_v0_ else d_922_v38_)).issubset(d_922_v38_)):
                        d_927_v42_: _dafny.Array
                        nw138_ = _dafny.Array(int(0), 17)
                        d_927_v42_ = nw138_
                        d_928_v43_: _dafny.Map
                        d_928_v43_ = _dafny.Map({d_927_v42_: d_879_v0_})
                        d_928_v43_ = d_928_v43_
                        d_929_v44_: _dafny.Array
                        def lambda29_(d_930_v0_):
                            def lambda30_(d_931_i2_):
                                return d_930_v0_

                            return lambda30_

                        init15_ = lambda29_(d_879_v0_)
                        nw139_ = _dafny.Array(None, 5)
                        for i0_15_ in range(nw139_.length(0)):
                            nw139_[i0_15_] = init15_(i0_15_)
                        d_929_v44_ = nw139_
                        d_929_v44_ = (d_929_v44_ if d_879_v0_ else d_929_v44_)
                        (globalState).f7 = ((0) - ((self).f22)) != (self.f23)
                        d_932_v45_: _dafny.Map
                        d_932_v45_ = _dafny.Map({(self).f11: d_879_v0_})
                        d_933_v46_: _dafny.Map
                        d_933_v46_ = _dafny.Map({(0) - ((self).f22): d_932_v45_})
                        d_934_v47_: D7
                        d_934_v47_ = D7_DC13((self).f12, (d_905_v23_).cardinality, self.f23)
                        d_935_v48_: _dafny.Seq
                        d_935_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxka"))
                        d_936_v49_: _dafny.Map
                        d_936_v49_ = _dafny.Map({self.f23: (self).f11})
                        d_937_v50_: C1
                        nw140_ = C1()
                        nw140_.ctor__(d_879_v0_, d_879_v0_)
                        d_937_v50_ = nw140_
                        d_938_v51_: _dafny.Map
                        d_938_v51_ = _dafny.Map({(self).f11: d_937_v50_})
                        rhs170_ = (d_933_v46_) | (_dafny.Map({p0: d_932_v45_}))
                        rhs171_ = ((((default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState)).set(default__.safeIndex(len(d_938_v51_), len(default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState))), (d_937_v50_).f18)) + (d_899_v18_)).set(default__.safeIndex((self).f22, len(((default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState)).set(default__.safeIndex(len(d_938_v51_), len(default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState))), (d_937_v50_).f18)) + (d_899_v18_))), d_879_v0_)).set(default__.safeIndex(self.f23, len((((default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState)).set(default__.safeIndex(len(d_938_v51_), len(default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState))), (d_937_v50_).f18)) + (d_899_v18_)).set(default__.safeIndex((self).f22, len(((default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState)).set(default__.safeIndex(len(d_938_v51_), len(default__.fm35((self).f22, ((d_936_v49_)[p0] if (p0) in (d_936_v49_) else (0) - (self.f23)), D7_DC13((self).f12, p0, p0), self.f23, globalState))), (d_937_v50_).f18)) + (d_899_v18_))), d_879_v0_))), (d_937_v50_).f18)
                        rhs172_ = (d_934_v47_ if (d_937_v50_).f19 else d_934_v47_)
                        rhs173_ = (d_935_v48_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eppqimqu")))
                        d_933_v46_ = rhs170_
                        d_899_v18_ = rhs171_
                        d_934_v47_ = rhs172_
                        d_935_v48_ = rhs173_
                        d_939_v52_: _dafny.Seq
                        d_939_v52_ = _dafny.SeqWithoutIsStrInference([-590, self.f23])
                        d_940_v53_: _dafny.MultiSet
                        d_940_v53_ = _dafny.MultiSet([(self).f22])
                        d_941_v54_: _dafny.Seq
                        d_941_v54_ = _dafny.SeqWithoutIsStrInference([d_939_v52_, d_939_v52_, d_939_v52_, (_dafny.SeqWithoutIsStrInference([p0, (d_940_v53_).cardinality, 894])) + (d_939_v52_), _dafny.SeqWithoutIsStrInference([(self).f11])])
                        d_939_v52_ = (d_941_v54_)[default__.safeIndex((len(d_939_v52_)) * ((self).f11), len(d_941_v54_))]
                    elif True:
                        d_942_v55_: _dafny.Array
                        nw141_ = _dafny.Array(False, 22)
                        d_942_v55_ = nw141_
                        index109_ = default__.safeIndex(414, (d_942_v55_).length(0))
                        (d_942_v55_)[index109_] = d_879_v0_
                        d_943_v56_: _dafny.Seq
                        d_943_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iq"))
                        d_944_v57_: D2
                        d_944_v57_ = D2_DC3(d_904_v22_)
                        index110_ = default__.safeIndex(414, (d_942_v55_).length(0))
                        (d_942_v55_)[index110_] = ((d_944_v57_).cf5).ispropersubset((_dafny.MultiSet([d_943_v56_])) | (d_904_v22_))
                        d_905_v23_ = (d_905_v23_) | ((_dafny.MultiSet([(d_942_v55_)[default__.safeIndex(414, (d_942_v55_).length(0))], (d_942_v55_)[default__.safeIndex(414, (d_942_v55_).length(0))], (d_942_v55_)[default__.safeIndex(414, (d_942_v55_).length(0))], (d_942_v55_)[default__.safeIndex(414, (d_942_v55_).length(0))], d_879_v0_])).intersection(d_905_v23_))
                        d_945_v58_: C2
                        nw142_ = C2()
                        nw142_.ctor__(default__.fm39(175, (d_942_v55_)[default__.safeIndex(414, (d_942_v55_).length(0))], (self).f11, globalState), p0)
                        d_945_v58_ = nw142_
                        d_946_v59_: _dafny.Map
                        d_946_v59_ = _dafny.Map({d_879_v0_: (0) - ((self.f23) * ((self).f22))})
                        d_947_v60_: _dafny.MultiSet
                        d_947_v60_ = _dafny.MultiSet([204])
                        d_948_v61_: _dafny.Map
                        d_948_v61_ = _dafny.Map({(self).f11: (self).f11})
                        d_949_v62_: _dafny.Map
                        d_949_v62_ = _dafny.Map({d_947_v60_: d_948_v61_})
                        d_946_v59_ = (d_946_v59_).set(not((self).fm4(self.f23, d_949_v62_, len(((d_943_v56_).set(default__.safeIndex((self).f11, len(d_943_v56_)), d_902_v20_)).set(default__.safeIndex(117, len((d_943_v56_).set(default__.safeIndex((self).f11, len(d_943_v56_)), d_902_v20_))), d_902_v20_)), True, globalState)), (0) - ((self).f22))
                        index111_ = default__.safeIndex(414, (d_942_v55_).length(0))
                        (d_942_v55_)[index111_] = not(not(d_879_v0_))
                    pass
            pass
        d_950_v63_: _dafny.Set
        d_950_v63_ = _dafny.Set({(self).f11})
        rhs174_ = d_879_v0_
        rhs175_ = default__.safeDivisionInt((self).f22, (self).f11)
        rhs176_ = p0
        rhs177_ = ((self).f22) not in (d_950_v63_)
        lhs137_ = self
        lhs138_ = self
        lhs139_ = globalState
        d_879_v0_ = rhs174_
        lhs137_.f23 = rhs175_
        lhs138_.f23 = rhs176_
        lhs139_.f6 = rhs177_
        d_951_v64_: _dafny.Seq
        d_951_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnfp"))
        d_952_v65_: _dafny.Seq
        d_952_v65_ = _dafny.SeqWithoutIsStrInference([len(d_951_v64_)])
        d_953_v66_: _dafny.MultiSet
        d_953_v66_ = _dafny.MultiSet([(d_952_v65_)[default__.safeIndex((self).f11, len(d_952_v65_))], (self).f11, (self).f11])
        d_954_v67_: _dafny.Map
        d_954_v67_ = _dafny.Map({d_951_v64_: d_953_v66_})
        r1 = d_954_v67_
        d_955_v68_: _dafny.Seq
        d_955_v68_ = _dafny.SeqWithoutIsStrInference([d_899_v18_, d_899_v18_, d_899_v18_])
        d_956_i3_: int
        d_956_i3_ = 0
        with _dafny.label("10"):
            while not (True) or ((d_900_v19_) in (d_955_v68_)):
                with _dafny.c_label("10"):
                    if (d_956_i3_) >= (100):
                        raise _dafny.Break("10")
                    d_956_i3_ = (d_956_i3_) + (1)
                    if d_879_v0_:
                        d_957_v69_: _dafny.Map
                        d_957_v69_ = _dafny.Map({785: d_951_v64_})
                        rhs178_ = (d_879_v0_) == ((_dafny.MultiSet([((d_957_v69_)[(self).f22] if ((self).f22) in (d_957_v69_) else d_951_v64_)])).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_951_v64_ for d_958_i4_ in range(default__.abs(761))]))))
                        rhs179_ = p0
                        rhs180_ = 598
                        lhs140_ = globalState
                        lhs141_ = globalState
                        lhs142_ = globalState
                        lhs140_.f6 = rhs178_
                        lhs141_.f1 = rhs179_
                        lhs142_.f1 = rhs180_
                        (globalState).f1 = (self).f22
                        d_959_v70_: str
                        d_959_v70_ = _dafny.CodePoint('g')
                        d_960_v72_: C0
                        nw143_ = C0()
                        nw143_.ctor__(d_900_v19_, (self).f22)
                        d_960_v72_ = nw143_
                        d_961_v73_: _dafny.Seq
                        d_961_v73_ = _dafny.SeqWithoutIsStrInference([d_960_v72_])
                        d_962_v74_: D0
                        d_962_v74_ = D0_DC0(len(d_961_v73_))
                        d_963_v75_: _dafny.MultiSet
                        d_963_v75_ = _dafny.MultiSet([d_962_v74_, d_962_v74_, d_962_v74_, d_962_v74_, d_962_v74_])
                        def iife97_():
                            coll49_ = _dafny.Map()
                            compr_49_: D0
                            for compr_49_ in (d_963_v75_).Elements:
                                d_964_v71_: D0 = compr_49_
                                if (d_964_v71_) in (d_963_v75_):
                                    coll49_[d_964_v71_] = (self).f11
                            return _dafny.Map(coll49_)
                        d_959_v70_ = default__.fm39(599, default__.fm0(globalState), len(iife97_()
                        ), globalState)
                        (globalState).f7 = (d_899_v18_) == (((d_960_v72_).f20) + (d_899_v18_))
                        d_965_v76_: _dafny.Array
                        def lambda31_(d_966_v0_, d_967_v66_):
                            def lambda32_(d_968_i5_):
                                return (d_967_v66_ if d_966_v0_ else d_967_v66_)

                            return lambda32_

                        init16_ = lambda31_(d_879_v0_, d_953_v66_)
                        nw144_ = _dafny.Array(None, 1)
                        for i0_16_ in range(nw144_.length(0)):
                            nw144_[i0_16_] = init16_(i0_16_)
                        d_965_v76_ = nw144_
                        index112_ = default__.safeIndex(30, (d_965_v76_).length(0))
                        (d_965_v76_)[index112_] = d_953_v66_
                        index113_ = default__.safeIndex(30, (d_965_v76_).length(0))
                        rhs181_ = default__.safeDivisionInt(p0, (self).f22)
                        rhs182_ = (d_951_v64_) + (((d_951_v64_) + (d_951_v64_)).set(default__.safeIndex((d_952_v65_)[default__.safeIndex(len(d_952_v65_), len(d_952_v65_))], len((d_951_v64_) + (d_951_v64_))), d_959_v70_))
                        rhs183_ = self.f23
                        rhs184_ = d_953_v66_
                        lhs143_ = globalState
                        lhs144_ = self
                        lhs145_ = d_965_v76_
                        lhs146_ = default__.safeIndex(30, (d_965_v76_).length(0))
                        lhs143_.f1 = rhs181_
                        d_951_v64_ = rhs182_
                        lhs144_.f23 = rhs183_
                        lhs145_[lhs146_] = rhs184_
                    elif True:
                        (globalState).f1 = (self).f22
                        d_969_v77_: _dafny.Seq
                        d_969_v77_ = _dafny.SeqWithoutIsStrInference([d_952_v65_])
                        d_970_v78_: _dafny.Map
                        d_970_v78_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([self.f23, -896])) in (d_969_v77_): self.f23})
                        d_970_v78_ = d_970_v78_
                        d_971_v79_: _dafny.Map
                        d_971_v79_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_972_i6_ in range(default__.abs(-274))])})
                        d_973_v80_: str
                        d_973_v80_ = _dafny.CodePoint('n')
                        rhs185_ = ((len(d_950_v63_)) - ((self).fm3(p0, False, d_879_v0_, globalState))) * (((self).f11 if d_879_v0_ else (self).f11))
                        rhs186_ = not(not((d_900_v19_)[default__.safeIndex(self.f23, len(d_900_v19_))]))
                        rhs187_ = ((default__.fm1(globalState)) + (_dafny.SeqWithoutIsStrInference([d_973_v80_]))) + (d_951_v64_)
                        rhs188_ = d_971_v79_
                        rhs189_ = True
                        lhs147_ = globalState
                        lhs148_ = globalState
                        lhs147_.f1 = rhs185_
                        lhs148_.f6 = rhs186_
                        d_951_v64_ = rhs187_
                        d_971_v79_ = rhs188_
                        d_879_v0_ = rhs189_
                        (globalState).f1 = (self).f11
                        d_974_v81_: C1
                        nw145_ = C1()
                        nw145_.ctor__(d_879_v0_, (d_953_v66_).isdisjoint(d_953_v66_))
                        d_974_v81_ = nw145_
                    d_975_v82_: _dafny.Map
                    d_975_v82_ = _dafny.Map({self.f23: d_899_v18_})
                    d_976_v83_: C0
                    nw146_ = C0()
                    nw146_.ctor__(((d_975_v82_)[(d_952_v65_)[default__.safeIndex(len(d_952_v65_), len(d_952_v65_))]] if ((d_952_v65_)[default__.safeIndex(len(d_952_v65_), len(d_952_v65_))]) in (d_975_v82_) else d_899_v18_), self.f23)
                    d_976_v83_ = nw146_
                    d_977_v84_: _dafny.Seq
                    d_977_v84_ = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f12, (self).f12, (self).f12])
                    d_977_v84_ = d_977_v84_
                    d_952_v65_ = _dafny.SeqWithoutIsStrInference([(self).f22, p0, (d_976_v83_).f21, self.f23, (self).f22])
                    pass
            pass
        d_978_v85_: _dafny.Set
        d_978_v85_ = _dafny.Set({d_952_v65_})
        d_979_v86_: D9
        d_979_v86_ = D9_DC19()
        d_980_v87_: _dafny.Array
        nw147_ = _dafny.Array(None, 18)
        nw147_[int(0)] = default__.fm42(d_978_v85_, globalState)
        nw147_[int(1)] = d_979_v86_
        nw147_[int(2)] = (d_979_v86_ if d_879_v0_ else d_979_v86_)
        nw147_[int(3)] = d_979_v86_
        nw147_[int(4)] = d_979_v86_
        nw147_[int(5)] = d_979_v86_
        nw147_[int(6)] = d_979_v86_
        nw147_[int(7)] = d_979_v86_
        nw147_[int(8)] = d_979_v86_
        nw147_[int(9)] = d_979_v86_
        nw147_[int(10)] = d_979_v86_
        nw147_[int(11)] = d_979_v86_
        nw147_[int(12)] = d_979_v86_
        nw147_[int(13)] = d_979_v86_
        nw147_[int(14)] = d_979_v86_
        nw147_[int(15)] = d_979_v86_
        nw147_[int(16)] = default__.fm42(d_978_v85_, globalState)
        nw147_[int(17)] = d_979_v86_
        d_980_v87_ = nw147_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_980_v87_).length(0)):
            d_981_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_981_i7_)) and ((d_981_i7_) < ((d_980_v87_).length(0)))):
                (d_980_v87_)[(d_981_i7_)] = ((d_979_v86_ if d_879_v0_ else d_979_v86_) if (self.f23) < (self.f23) else d_979_v86_)
        d_982_v88_: _dafny.Array
        def lambda33_(d_983_v0_):
            def lambda34_(d_984_i8_):
                return _dafny.MultiSet([d_983_v0_])

            return lambda34_

        init17_ = lambda33_(d_879_v0_)
        nw148_ = _dafny.Array(None, 8)
        for i0_17_ in range(nw148_.length(0)):
            nw148_[i0_17_] = init17_(i0_17_)
        d_982_v88_ = nw148_
        d_985_v89_: _dafny.Map
        d_985_v89_ = _dafny.Map({(self).f22: d_982_v88_})
        r0 = ((d_985_v89_)[(0) - (-6)] if ((0) - (-6)) in (d_985_v89_) else d_982_v88_)
        r1 = (d_954_v67_) | (default__.fm43(d_879_v0_, globalState))
        return r0, r1

    @property
    def f22(self):
        return self._f22

class C4(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Set = _dafny.Set({})
        self._f11: int = int(0)
        self._f16: _dafny.MultiSet = _dafny.MultiSet({})
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f12(self):
        return self._f12
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f16, f17, f12, f11):
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f12 = f12
        (self)._f11 = f11

    def fm5(self, p0, p1, p2, p3, globalState):
        return (self).f17

    def fm6(self, p0, p1, globalState):
        return (((self).f16).intersection((self).f16)).cardinality

    def fm3(self, p0, p1, p2, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njlnvedvb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjlvfx")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uocufamr"))))

    def fm4(self, p0, p1, p2, p3, globalState):
        return (self).f17

    def m0(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        (globalState).f7 = (self).f17
        d_986_i0_: int
        d_986_i0_ = 0
        with _dafny.label("11"):
            while (self).f17:
                with _dafny.c_label("11"):
                    if (d_986_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_986_i0_ = (d_986_i0_) + (1)
                    d_987_v0_: _dafny.MultiSet
                    d_987_v0_ = _dafny.MultiSet([(self).f17, (self).f17, (self).f17])
                    r0 = ((d_987_v0_ if (self).f17 else _dafny.MultiSet([(self).f17]))) - (d_987_v0_)
                    d_988_v1_: _dafny.Array
                    nw149_ = _dafny.Array(False, 8)
                    d_988_v1_ = nw149_
                    index114_ = default__.safeIndex(552, (d_988_v1_).length(0))
                    (d_988_v1_)[index114_] = (self).f17
                    d_989_v2_: _dafny.Array
                    nw150_ = _dafny.Array(D2.default()(), 4)
                    d_989_v2_ = nw150_
                    d_990_v3_: _dafny.MultiSet
                    d_990_v3_ = _dafny.MultiSet([d_989_v2_])
                    index115_ = default__.safeIndex(552, (d_988_v1_).length(0))
                    (d_988_v1_)[index115_] = ((_dafny.MultiSet([d_989_v2_, d_989_v2_])) - (d_990_v3_)).isdisjoint(d_990_v3_)
                    (globalState).f7 = False
                    d_991_v4_: _dafny.Array
                    nw151_ = _dafny.Array(_dafny.Set({}), 4)
                    d_991_v4_ = nw151_
                    d_992_v5_: _dafny.Seq
                    d_992_v5_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])
                    d_993_v6_: _dafny.Map
                    d_993_v6_ = _dafny.Map({not((self).f17): (self).f11})
                    d_994_v7_: _dafny.Set
                    d_994_v7_ = _dafny.Set({_dafny.Map({(d_988_v1_)[default__.safeIndex(552, (d_988_v1_).length(0))]: len(_dafny.Map({d_992_v5_: (self).fm3((self).f11, (d_988_v1_)[default__.safeIndex(552, (d_988_v1_).length(0))], (d_988_v1_)[default__.safeIndex(552, (d_988_v1_).length(0))], globalState)}))}), d_993_v6_})
                    index116_ = default__.safeIndex(359, (d_991_v4_).length(0))
                    (d_991_v4_)[index116_] = (d_994_v7_) - (d_994_v7_)
                    d_995_v8_: _dafny.Array
                    nw152_ = _dafny.Array(int(0), 12)
                    d_995_v8_ = nw152_
                    d_996_v9_: _dafny.MultiSet
                    d_996_v9_ = _dafny.MultiSet([d_995_v8_])
                    d_997_v10_: D2
                    d_997_v10_ = D2_DC4((0) - ((self).f11), ((d_996_v9_)[d_995_v8_] if (d_995_v8_) in (d_996_v9_) else (self).f11), (self).f11)
                    d_998_v11_: _dafny.Seq
                    d_998_v11_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
                    d_999_v12_: _dafny.Seq
                    d_999_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
                    d_1000_v13_: _dafny.Seq
                    d_1000_v13_ = _dafny.SeqWithoutIsStrInference([len(d_998_v11_), len(d_992_v5_), len(d_999_v12_)])
                    pat_let_tv36_ = d_998_v11_
                    d_1001_v14_: _dafny.Map
                    def iife98_(_pat_let24_0):
                        def iife99_(d_1002_dt__update__tmp_h0_):
                            def iife100_(_pat_let25_0):
                                def iife101_(d_1003_dt__update_hcf6_h0_):
                                    def iife102_(_pat_let26_0):
                                        def iife103_(d_1004_dt__update_hcf8_h0_):
                                            return D2_DC4(d_1003_dt__update_hcf6_h0_, (d_1002_dt__update__tmp_h0_).cf7, d_1004_dt__update_hcf8_h0_)
                                        return iife103_(_pat_let26_0)
                                    return iife102_(len(pat_let_tv36_))
                                return iife101_(_pat_let25_0)
                            return iife100_((self).f11)
                        return iife99_(_pat_let24_0)
                    d_1001_v14_ = _dafny.Map({(self).f11: iife98_(d_997_v10_)})
                    index117_ = default__.safeIndex(359, (d_991_v4_).length(0))
                    (d_991_v4_)[index117_] = default__.fm10((self).f11, ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smwrdj")) if not((d_988_v1_)[default__.safeIndex(552, (d_988_v1_).length(0))]) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tctiechk")))).set(default__.safeIndex((self).f11, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smwrdj")) if not((d_988_v1_)[default__.safeIndex(552, (d_988_v1_).length(0))]) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tctiechk"))))), _dafny.CodePoint('q')), len(d_1001_v14_), (d_988_v1_)[default__.safeIndex(552, (d_988_v1_).length(0))], globalState)
                    pass
            pass
        d_1005_v15_: _dafny.Array
        nw153_ = _dafny.Array(False, 5)
        d_1005_v15_ = nw153_
        d_1006_v16_: _dafny.Seq
        d_1006_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oc"))
        rhs190_ = (self).fm6((self).f17, default__.safeModuloInt((self).f11, (0) - ((self).f11)), globalState)
        rhs191_ = (self).f11
        rhs192_ = d_1005_v15_
        rhs193_ = d_1005_v15_
        rhs194_ = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjar"))) + (d_1006_v16_))) + (len((_dafny.Map({True: (self).f17})) | (default__.fm11(globalState))))
        lhs149_ = globalState
        lhs150_ = globalState
        lhs151_ = globalState
        lhs149_.f1 = rhs190_
        lhs150_.f1 = rhs191_
        d_1005_v15_ = rhs192_
        d_1005_v15_ = rhs193_
        lhs151_.f1 = rhs194_
        index118_ = default__.safeIndex(627, (d_1005_v15_).length(0))
        (d_1005_v15_)[index118_] = (self).f17
        index119_ = default__.safeIndex(627, (d_1005_v15_).length(0))
        (d_1005_v15_)[index119_] = (self).f17
        (globalState).f6 = (self).f17
        d_1007_v17_: _dafny.Map
        d_1007_v17_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]]): (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]})
        if (default__.fm12((0) - ((self).f11), d_1006_v16_, globalState)) in (d_1007_v17_):
            d_1008_v18_: _dafny.Map
            d_1008_v18_ = _dafny.Map({(self).f17: (self).f11})
            rhs195_ = len(d_1008_v18_)
            rhs196_ = (self).f17
            lhs152_ = globalState
            lhs153_ = globalState
            lhs152_.f1 = rhs195_
            lhs153_.f7 = rhs196_
            d_1009_v19_: _dafny.Map
            d_1009_v19_ = _dafny.Map({(d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]: (self).f17})
            d_1010_v20_: _dafny.Map
            d_1010_v20_ = _dafny.Map({len(d_1006_v16_): (self).f17})
            d_1011_v21_: _dafny.Map
            d_1011_v21_ = _dafny.Map({p0: ((d_1010_v20_)[(self).f11] if ((self).f11) in (d_1010_v20_) else False)})
            d_1012_v22_: _dafny.Seq
            d_1012_v22_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeModuloInt((self).f11, (self).f11)), ((self).f11 if ((d_1009_v19_)[(self).f17] if ((self).f17) in (d_1009_v19_) else (self).f17) else (self).f11), (self).f11, (0) - (len((((d_1011_v21_).set(p0, (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))])).set(_dafny.CodePoint('w'), (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))])) | (d_1011_v21_))), default__.safeDivisionInt((self).f11, 923)])
            d_1013_v23_: _dafny.Array
            nw154_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1013_v23_ = nw154_
            d_1014_v24_: _dafny.Array
            nw155_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_1014_v24_ = nw155_
            index120_ = default__.safeIndex(488, (d_1013_v23_).length(0))
            (d_1013_v23_)[index120_] = d_1014_v24_
            d_1015_v26_: _dafny.Map
            d_1015_v26_ = _dafny.Map({(self).f16: (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]})
            d_1016_v28_: _dafny.Array
            d_1016_v28_ = d_1014_v24_
            index121_ = default__.safeIndex(488, (d_1013_v23_).length(0))
            def iife104_():
                coll50_ = _dafny.Map()
                compr_50_: _dafny.MultiSet
                for compr_50_ in (d_1015_v26_).keys.Elements:
                    d_1017_v25_: _dafny.MultiSet = compr_50_
                    if (d_1017_v25_) in (d_1015_v26_):
                        def iife105_():
                            coll51_ = _dafny.Map()
                            compr_51_: int
                            for compr_51_ in _dafny.IntegerRange(257, 278):
                                d_1018_v27_: int = compr_51_
                                if ((257) <= (d_1018_v27_)) and ((d_1018_v27_) < (278)):
                                    coll51_[(d_1018_v27_) * (len((self).f12))] = len(_dafny.SeqWithoutIsStrInference([(self).f17, (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]]))
                            return _dafny.Map(coll51_)
                        coll50_[d_1017_v25_] = iife105_()

                return _dafny.Map(coll50_)
            rhs197_ = (d_1012_v22_) + (d_1012_v22_)
            rhs198_ = (d_1012_v22_)[default__.safeIndex((self).f11, len(d_1012_v22_))]
            rhs199_ = (d_1014_v24_ if ((self).fm4((self).f11, iife104_()
            , (self).f11, (self).f17, globalState) if (self).f17 else (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]) else (d_1016_v28_))
            rhs200_ = (d_1005_v15_ if (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))] else d_1005_v15_)
            lhs154_ = globalState
            lhs155_ = d_1013_v23_
            lhs156_ = default__.safeIndex(488, (d_1013_v23_).length(0))
            d_1012_v22_ = rhs197_
            lhs154_.f1 = rhs198_
            lhs155_[lhs156_] = rhs199_
            d_1005_v15_ = rhs200_
            if not((d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]):
                d_1019_v29_: int
                d_1020_v30_: bool
                d_1021_v31_: int
                out28_: int
                out29_: bool
                out30_: int
                out28_, out29_, out30_ = (self).m5((d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], not ((self).f17) or (True), (self).f11, globalState)
                d_1019_v29_ = out28_
                d_1020_v30_ = out29_
                d_1021_v31_ = out30_
                d_1006_v16_ = default__.fm1(globalState)
                index122_ = default__.safeIndex(627, (d_1005_v15_).length(0))
                (d_1005_v15_)[index122_] = (self).f17
                d_1022_v32_: _dafny.Seq
                d_1022_v32_ = _dafny.SeqWithoutIsStrInference([(d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], (self).f17, (self).f17])
                d_1023_v33_: C0
                nw156_ = C0()
                nw156_.ctor__(d_1022_v32_, len(_dafny.SeqWithoutIsStrInference([(0) - (len(d_1010_v20_)) for d_1024_i1_ in range(default__.abs(648))])))
                d_1023_v33_ = nw156_
                d_1025_v34_: T1
                nw157_ = C3()
                nw157_.ctor__(len(d_1006_v16_), 729, d_1021_v31_, (self).f12)
                d_1025_v34_ = nw157_
                d_1026_v35_: _dafny.Seq
                d_1026_v35_ = _dafny.SeqWithoutIsStrInference([d_1025_v34_, d_1025_v34_, d_1025_v34_, d_1025_v34_])
                index123_ = default__.safeIndex(627, (d_1005_v15_).length(0))
                rhs201_ = default__.safeDivisionInt((self).f11, (_dafny.MultiSet([len(d_1026_v35_)])).cardinality)
                rhs202_ = (((self).f16)[(self).f11] if ((self).f11) in ((self).f16) else (d_1021_v31_) * (d_1019_v29_))
                rhs203_ = not(not(((d_1010_v20_)[(d_1025_v34_).f11] if ((d_1025_v34_).f11) in (d_1010_v20_) else d_1020_v30_)))
                rhs204_ = d_1020_v30_
                lhs157_ = globalState
                lhs158_ = globalState
                lhs159_ = d_1005_v15_
                lhs160_ = default__.safeIndex(627, (d_1005_v15_).length(0))
                d_1021_v31_ = rhs201_
                lhs157_.f1 = rhs202_
                lhs158_.f6 = rhs203_
                lhs159_[lhs160_] = rhs204_
            elif True:
                d_1027_v36_: _dafny.Array
                def lambda35_(d_1028_v16_):
                    def lambda36_(d_1029_i2_):
                        return (d_1028_v16_)

                    return lambda36_

                init18_ = lambda35_(d_1006_v16_)
                nw158_ = _dafny.Array(None, 25)
                for i0_18_ in range(nw158_.length(0)):
                    nw158_[i0_18_] = init18_(i0_18_)
                d_1027_v36_ = nw158_
                d_1027_v36_ = d_1027_v36_
                d_1030_v37_: _dafny.Map
                d_1030_v37_ = _dafny.Map({d_1012_v22_: (self).f11})
                d_1031_v38_: _dafny.Map
                d_1031_v38_ = _dafny.Map({d_1030_v37_: (self).f11})
                (globalState).f1 = ((0) - ((self).f11)) * (((d_1031_v38_)[d_1030_v37_] if (d_1030_v37_) in (d_1031_v38_) else (self).f11))
                (globalState).f6 = (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]
                d_1032_v39_: C2
                nw159_ = C2()
                nw159_.ctor__(_dafny.CodePoint('d'), default__.safeDivisionInt((self).f11, (self).f11))
                d_1032_v39_ = nw159_
                (globalState).f6 = ((self).f11) in ((self).f16)
            d_1033_v41_: _dafny.Set
            def iife106_():
                coll52_ = _dafny.Map()
                compr_52_: int
                for compr_52_ in (_dafny.SeqWithoutIsStrInference([(self).f11 for d_1034_i3_ in range(default__.abs(-494))])).Elements:
                    d_1035_v40_: int = compr_52_
                    if (d_1035_v40_) in (_dafny.SeqWithoutIsStrInference([(self).f11 for d_1034_i3_ in range(default__.abs(-494))])):
                        coll52_[(d_1035_v40_) - ((_dafny.MultiSet([(self).f11])).cardinality)] = True
                return _dafny.Map(coll52_)
            d_1033_v41_ = _dafny.Set({len(d_1012_v22_), len(iife106_()
            )})
            (globalState).f1 = len((d_1033_v41_) | (d_1033_v41_))
            d_1036_v42_: C3
            nw160_ = C3()
            nw160_.ctor__(len(d_1006_v16_), (self).f11, (self).f11, (self).f12)
            d_1036_v42_ = nw160_
            d_1037_v43_: _dafny.Seq
            d_1037_v43_ = _dafny.SeqWithoutIsStrInference([d_1036_v42_, d_1036_v42_])
            (globalState).f1 = len(d_1037_v43_)
        elif True:
            d_1038_v44_: _dafny.Array
            nw161_ = _dafny.Array(_dafny.CodePoint('D'), 7)
            d_1038_v44_ = nw161_
            d_1039_v45_: _dafny.Seq
            d_1039_v45_ = _dafny.SeqWithoutIsStrInference([(d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]])
            d_1040_v46_: D11
            d_1040_v46_ = D11_DC24((self).f11, d_1039_v45_, p0)
            index124_ = default__.safeIndex(627, (d_1005_v15_).length(0))
            rhs205_ = (d_1040_v46_).cf34
            rhs206_ = (self).f11
            rhs207_ = (self).f17
            rhs208_ = (self).f11
            rhs209_ = d_1038_v44_
            lhs161_ = globalState
            lhs162_ = globalState
            lhs163_ = d_1005_v15_
            lhs164_ = default__.safeIndex(627, (d_1005_v15_).length(0))
            lhs165_ = globalState
            lhs161_.f1 = rhs205_
            lhs162_.f1 = rhs206_
            lhs163_[lhs164_] = rhs207_
            lhs165_.f1 = rhs208_
            d_1038_v44_ = rhs209_
            d_1041_v47_: _dafny.Map
            d_1041_v47_ = _dafny.Map({((self).f11) + ((self).f11): ((self).f11) + ((self).f11)})
            d_1041_v47_ = (d_1041_v47_).set((self).f11, (self).f11)
            d_1042_v48_: _dafny.Array
            nw162_ = _dafny.Array(int(0), 1)
            d_1042_v48_ = nw162_
            d_1043_v49_: D6
            d_1043_v49_ = D6_DC11((self).f11, (self).f11, (0) - ((self).f11))
            d_1044_v50_: _dafny.Map
            d_1044_v50_ = _dafny.Map({d_1043_v49_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "go"))).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "go")))), p0)})
            d_1045_v51_: _dafny.MultiSet
            d_1045_v51_ = _dafny.MultiSet([((d_1044_v50_)[d_1043_v49_] if (d_1043_v49_) in (d_1044_v50_) else d_1006_v16_)])
            index125_ = default__.safeIndex(728, (d_1042_v48_).length(0))
            (d_1042_v48_)[index125_] = (d_1045_v51_).cardinality
            index126_ = default__.safeIndex(728, (d_1042_v48_).length(0))
            (d_1042_v48_)[index126_] = (-940) * ((0) - ((self).f11))
            d_1046_v52_: _dafny.Set
            d_1046_v52_ = _dafny.Set({(self).f11, (self).f11})
            if (((self).f11) - (len(d_1046_v52_))) > (len(d_1006_v16_)):
                (globalState).f1 = -492
                d_1047_v53_: _dafny.MultiSet
                d_1047_v53_ = _dafny.MultiSet([(d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]])
                d_1048_v54_: C1
                nw163_ = C1()
                nw163_.ctor__((_dafny.MultiSet(d_1039_v45_)).issubset(d_1047_v53_), (self).f17)
                d_1048_v54_ = nw163_
                d_1049_v55_: _dafny.Seq
                d_1049_v55_ = _dafny.SeqWithoutIsStrInference([(d_1042_v48_)[default__.safeIndex(728, (d_1042_v48_).length(0))], -48, (d_1042_v48_)[default__.safeIndex(728, (d_1042_v48_).length(0))]])
                (globalState).f7 = (default__.fm0(globalState)) and (((d_1049_v55_)[default__.safeIndex((self).f11, len(d_1049_v55_))]) > ((self).f11))
                index127_ = default__.safeIndex(627, (d_1005_v15_).length(0))
                (d_1005_v15_)[index127_] = (self).f17
                (globalState).f7 = (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]
            elif True:
                d_1050_v56_: _dafny.Map
                d_1050_v56_ = _dafny.Map({(self).f17: (self).f17})
                d_1051_v57_: D8
                d_1051_v57_ = D8_DC15(d_1050_v56_)
                d_1052_v58_: D8
                d_1052_v58_ = D8_DC17(d_1051_v57_)
                d_1053_v59_: _dafny.Map
                d_1053_v59_ = _dafny.Map({d_1052_v58_: (d_1042_v48_)[default__.safeIndex(728, (d_1042_v48_).length(0))]})
                d_1054_v60_: _dafny.Seq
                d_1054_v60_ = _dafny.SeqWithoutIsStrInference([d_1053_v59_, d_1053_v59_, d_1053_v59_, d_1053_v59_])
                d_1055_v61_: _dafny.MultiSet
                d_1055_v61_ = _dafny.MultiSet([default__.fm0(globalState)])
                d_1056_v62_: D13
                d_1056_v62_ = D13_DC27(d_1053_v59_)
                d_1057_v63_: _dafny.Map
                d_1057_v63_ = _dafny.Map({(self).f17: d_1054_v60_})
                d_1058_v64_: _dafny.Map
                d_1058_v64_ = _dafny.Map({11: (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]})
                d_1059_v65_: _dafny.Array
                nw164_ = _dafny.Array(None, 15)
                nw164_[int(0)] = d_1054_v60_
                nw164_[int(1)] = d_1054_v60_
                nw164_[int(2)] = ((d_1054_v60_).set(default__.safeIndex(((d_1055_v61_)[True] if (True) in (d_1055_v61_) else (self).f11), len(d_1054_v60_)), d_1053_v59_)) + (d_1054_v60_)
                nw164_[int(3)] = d_1054_v60_
                nw164_[int(4)] = d_1054_v60_
                nw164_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1053_v59_, d_1053_v59_, _dafny.Map({d_1052_v58_: (self).f11})])
                nw164_[int(6)] = d_1054_v60_
                nw164_[int(7)] = ((_dafny.SeqWithoutIsStrInference([(d_1056_v62_).cf39, d_1053_v59_, (d_1056_v62_).cf39])) + (_dafny.SeqWithoutIsStrInference([d_1053_v59_]))).set(default__.safeIndex((self).f11, len((_dafny.SeqWithoutIsStrInference([(d_1056_v62_).cf39, d_1053_v59_, (d_1056_v62_).cf39])) + (_dafny.SeqWithoutIsStrInference([d_1053_v59_])))), d_1053_v59_)
                nw164_[int(8)] = ((d_1057_v63_)[((d_1058_v64_)[(self).f11] if ((self).f11) in (d_1058_v64_) else (self).f17)] if (((d_1058_v64_)[(self).f11] if ((self).f11) in (d_1058_v64_) else (self).f17)) in (d_1057_v63_) else default__.fm44(not((self).f17), (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], globalState))
                nw164_[int(9)] = ((default__.fm44((self).f17, (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], globalState)).set(default__.safeIndex(960, len(default__.fm44((self).f17, (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], globalState))), d_1053_v59_)) + (d_1054_v60_)
                nw164_[int(10)] = (d_1054_v60_) + (d_1054_v60_)
                nw164_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1053_v59_, d_1053_v59_])
                nw164_[int(12)] = default__.fm44(not((self).f17), (self).f17, globalState)
                nw164_[int(13)] = (d_1054_v60_) + (d_1054_v60_)
                nw164_[int(14)] = _dafny.SeqWithoutIsStrInference([(d_1053_v59_).set(D8_DC17(d_1051_v57_), (0) - (-8))])
                d_1059_v65_ = nw164_
                index128_ = default__.safeIndex(956, (d_1059_v65_).length(0))
                (d_1059_v65_)[index128_] = ((d_1057_v63_)[not((d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))])] if (not((d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))])) in (d_1057_v63_) else _dafny.SeqWithoutIsStrInference([d_1053_v59_ for d_1060_i4_ in range(default__.abs(-255))]))
                index129_ = default__.safeIndex(956, (d_1059_v65_).length(0))
                (d_1059_v65_)[index129_] = d_1054_v60_
                d_1061_v66_: _dafny.Array
                def lambda37_(d_1062_v58_):
                    def lambda38_(d_1063_i5_):
                        return d_1062_v58_

                    return lambda38_

                init19_ = lambda37_(d_1052_v58_)
                nw165_ = _dafny.Array(None, 23)
                for i0_19_ in range(nw165_.length(0)):
                    nw165_[i0_19_] = init19_(i0_19_)
                d_1061_v66_ = nw165_
                index130_ = default__.safeIndex(451, (d_1061_v66_).length(0))
                (d_1061_v66_)[index130_] = d_1052_v58_
                index131_ = default__.safeIndex(451, (d_1061_v66_).length(0))
                (d_1061_v66_)[index131_] = d_1052_v58_
                d_1005_v15_ = d_1005_v15_
                (globalState).f6 = not(((self).f11) >= ((d_1042_v48_)[default__.safeIndex(728, (d_1042_v48_).length(0))]))
                d_1064_v67_: _dafny.Array
                nw166_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_1064_v67_ = nw166_
                d_1065_v68_: _dafny.Map
                d_1065_v68_ = _dafny.Map({(self).f17: d_1064_v67_})
                (globalState).f1 = len((d_1065_v68_).set(((self).f17) or ((self).f17), d_1064_v67_))
            (globalState).f7 = (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]
        d_1066_v70_: _dafny.Map
        d_1066_v70_ = _dafny.Map({((self).f16).set((self).f11, default__.abs((self).f11)): (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))]})
        d_1067_v72_: _dafny.MultiSet
        def iife107_():
            coll53_ = _dafny.Map()
            compr_53_: _dafny.MultiSet
            for compr_53_ in (d_1066_v70_).keys.Elements:
                d_1068_v69_: _dafny.MultiSet = compr_53_
                if (d_1068_v69_) in (d_1066_v70_):
                    def iife108_():
                        coll54_ = _dafny.Map()
                        compr_54_: int
                        for compr_54_ in _dafny.IntegerRange(958, 898):
                            d_1069_v71_: int = compr_54_
                            if ((958) <= (d_1069_v71_)) and ((d_1069_v71_) < (898)):
                                coll54_[default__.safeModuloInt(d_1069_v71_, (self).f11)] = -113
                        return _dafny.Map(coll54_)
                    coll53_[d_1068_v69_] = iife108_()

            return _dafny.Map(coll53_)
        d_1067_v72_ = _dafny.MultiSet([(self).fm4((self).f11, iife107_()
        , 893, (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], globalState)])
        d_1070_v73_: D0
        d_1070_v73_ = D0_DC0((self).f11)
        d_1071_v74_: _dafny.Seq
        d_1071_v74_ = _dafny.SeqWithoutIsStrInference([True, (d_1005_v15_)[default__.safeIndex(627, (d_1005_v15_).length(0))], (self).f17, (self).f17])
        r0 = (d_1067_v72_).set((self).fm5(d_1070_v73_, d_1071_v74_, (self).f11, _dafny.CodePoint('x'), globalState), default__.abs(((self).f11) * (len((self).f12))))
        return r0

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        d_1072_i0_: int
        d_1072_i0_ = 0
        with _dafny.label("12"):
            while (self).f17:
                with _dafny.c_label("12"):
                    if (d_1072_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_1072_i0_ = (d_1072_i0_) + (1)
                    d_1073_v0_: str
                    d_1073_v0_ = _dafny.CodePoint('c')
                    d_1074_v1_: C2
                    nw167_ = C2()
                    nw167_.ctor__(d_1073_v0_, 813)
                    d_1074_v1_ = nw167_
                    (globalState).f1 = (self).f11
                    d_1075_v2_: _dafny.Seq
                    d_1075_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulwlw"))
                    d_1076_v3_: _dafny.Seq
                    d_1076_v3_ = _dafny.SeqWithoutIsStrInference([True])
                    d_1077_v4_: _dafny.Array
                    nw168_ = _dafny.Array(None, 9)
                    nw168_[int(0)] = default__.fm12((self).f11, d_1075_v2_, globalState)
                    nw168_[int(1)] = (d_1076_v3_).set(default__.safeIndex((self).f11, len(d_1076_v3_)), (self).f17)
                    nw168_[int(2)] = d_1076_v3_
                    nw168_[int(3)] = d_1076_v3_
                    nw168_[int(4)] = (d_1076_v3_).set(default__.safeIndex((self).f11, len(d_1076_v3_)), (self).f17)
                    nw168_[int(5)] = d_1076_v3_
                    nw168_[int(6)] = (d_1076_v3_) + (d_1076_v3_)
                    nw168_[int(7)] = d_1076_v3_
                    nw168_[int(8)] = d_1076_v3_
                    d_1077_v4_ = nw168_
                    index132_ = default__.safeIndex(346, (d_1077_v4_).length(0))
                    (d_1077_v4_)[index132_] = d_1076_v3_
                    index133_ = default__.safeIndex(346, (d_1077_v4_).length(0))
                    (d_1077_v4_)[index133_] = (d_1076_v3_) + (d_1076_v3_)
                    d_1078_v5_: _dafny.Map
                    d_1078_v5_ = _dafny.Map({p0: (self).f17})
                    d_1079_v6_: _dafny.Map
                    d_1079_v6_ = _dafny.Map({(self).f17: (self).f17})
                    d_1080_v7_: _dafny.Array
                    nw169_ = _dafny.Array(None, 9)
                    nw169_[int(0)] = len(d_1078_v5_)
                    nw169_[int(1)] = p0
                    nw169_[int(2)] = (self).f11
                    nw169_[int(3)] = p0
                    nw169_[int(4)] = len(d_1079_v6_)
                    nw169_[int(5)] = 815
                    nw169_[int(6)] = (self).f11
                    nw169_[int(7)] = len((self).f12)
                    nw169_[int(8)] = 745
                    d_1080_v7_ = nw169_
                    d_1081_v8_: _dafny.Map
                    d_1081_v8_ = _dafny.Map({d_1080_v7_: not(((self).f11) >= (len((d_1077_v4_)[default__.safeIndex(346, (d_1077_v4_).length(0))])))})
                    d_1081_v8_ = (d_1081_v8_).set(d_1080_v7_, (self).f17)
                    pass
            pass
        d_1082_v9_: _dafny.Seq
        d_1082_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        d_1083_v10_: C3
        nw170_ = C3()
        nw170_.ctor__((D13_DC28(d_1082_v9_, (self).f11, (self).f17)).cf41, (p0) + (len(d_1082_v9_)), default__.fm2(p0, p0, globalState), (self).f12)
        d_1083_v10_ = nw170_
        if (False) or (((self).f17) or ((self).f17)):
            d_1084_v11_: _dafny.Seq
            d_1084_v11_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16])
            if (_dafny.SeqWithoutIsStrInference([(self).f16 for d_1085_i1_ in range(default__.abs(185))])) < (d_1084_v11_):
                d_1086_v13_: _dafny.Seq
                d_1086_v13_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_1083_v10_).f22)])
                d_1087_v14_: _dafny.Set
                d_1087_v14_ = _dafny.Set({(self).f11, (d_1083_v10_).f22})
                d_1088_v15_: _dafny.Seq
                d_1088_v15_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                d_1089_v16_: D0
                d_1089_v16_ = D0_DC0(-405)
                d_1090_v17_: _dafny.Map
                d_1090_v17_ = _dafny.Map({811: (self).f11})
                d_1091_v18_: str
                d_1091_v18_ = _dafny.CodePoint('u')
                d_1092_v19_: _dafny.Array
                nw171_ = _dafny.Array(None, 28)
                nw171_[int(0)] = (self).f17
                nw171_[int(1)] = (self).f17
                nw171_[int(2)] = not((self).f17)
                nw171_[int(3)] = (d_1083_v10_.f23) >= ((0) - (d_1083_v10_.f23))
                def iife109_():
                    coll55_ = _dafny.Set()
                    compr_55_: _dafny.MultiSet
                    for compr_55_ in (d_1084_v11_).Elements:
                        d_1093_v12_: _dafny.MultiSet = compr_55_
                        if (d_1093_v12_) in (d_1084_v11_):
                            coll55_ = coll55_.union(_dafny.Set([d_1093_v12_]))
                    return _dafny.Set(coll55_)
                nw171_[int(4)] = ((self).f11) != (len(iife109_()
                ))
                nw171_[int(5)] = (d_1082_v9_) <= (d_1082_v9_)
                nw171_[int(6)] = not ((self).f17) or (not(not(not((self).f17))))
                nw171_[int(7)] = (d_1086_v13_) < (d_1086_v13_)
                nw171_[int(8)] = (self).f17
                nw171_[int(9)] = (self).f17
                nw171_[int(10)] = (_dafny.Set({p0})).isdisjoint(d_1087_v14_)
                nw171_[int(11)] = (self).f17
                nw171_[int(12)] = (self).f17
                nw171_[int(13)] = True
                nw171_[int(14)] = ((self).f17) not in (d_1088_v15_)
                nw171_[int(15)] = (self).f17
                nw171_[int(16)] = (self).f17
                nw171_[int(17)] = (self).f17
                nw171_[int(18)] = (self).f17
                nw171_[int(19)] = (self).f17
                nw171_[int(20)] = (self).f17
                nw171_[int(21)] = not(True)
                nw171_[int(22)] = (self).f17
                nw171_[int(23)] = (d_1083_v10_).fm5(d_1089_v16_, d_1088_v15_, (d_1086_v13_)[default__.safeIndex(len(d_1090_v17_), len(d_1086_v13_))], d_1091_v18_, globalState)
                nw171_[int(24)] = not ((self).f17) or (True)
                nw171_[int(25)] = (self).f17
                nw171_[int(26)] = (self).f17
                nw171_[int(27)] = (len(d_1082_v9_)) > (d_1083_v10_.f23)
                d_1092_v19_ = nw171_
                index134_ = default__.safeIndex(355, (d_1092_v19_).length(0))
                (d_1092_v19_)[index134_] = (self).fm4((d_1083_v10_).fm6((self).f17, d_1083_v10_.f23, globalState), _dafny.Map({(self).f16: d_1090_v17_}), (d_1083_v10_).f22, True, globalState)
                index135_ = default__.safeIndex(355, (d_1092_v19_).length(0))
                (d_1092_v19_)[index135_] = (default__.safeModuloInt(len(d_1082_v9_), d_1083_v10_.f23)) != (d_1083_v10_.f23)
                d_1091_v18_ = d_1091_v18_
                d_1094_v20_: _dafny.Array
                nw172_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_1094_v20_ = nw172_
                d_1094_v20_ = d_1094_v20_
                d_1095_v21_: _dafny.Seq
                d_1095_v21_ = d_1082_v9_
                d_1095_v21_ = d_1095_v21_
                index136_ = default__.safeIndex(355, (d_1092_v19_).length(0))
                (d_1092_v19_)[index136_] = (d_1083_v10_).fm5(d_1089_v16_, d_1088_v15_, p0, default__.fm39(p0, False, p0, globalState), globalState)
            elif True:
                d_1096_v22_: _dafny.Map
                d_1096_v22_ = _dafny.Map({d_1083_v10_.f23: (d_1083_v10_).f22})
                d_1097_v23_: _dafny.Seq
                d_1097_v23_ = _dafny.SeqWithoutIsStrInference([len(d_1096_v22_)])
                d_1098_v24_: _dafny.Seq
                d_1098_v24_ = _dafny.SeqWithoutIsStrInference([((d_1083_v10_).f22) + (len(d_1097_v23_))])
                d_1099_v25_: D10
                d_1099_v25_ = D10_DC21((self).f17, (self).f17)
                d_1100_v28_: _dafny.Map
                d_1100_v28_ = _dafny.Map({False: (self).f17})
                def iife110_():
                    def iife112_():
                        coll58_ = _dafny.Set()
                        compr_58_: _dafny.Map
                        for compr_58_ in (_dafny.Set({_dafny.Map({(self).f17: (self).f17})})).Elements:
                            d_1103_v27_: _dafny.Map = compr_58_
                            if (d_1103_v27_) in (_dafny.Set({_dafny.Map({(self).f17: (self).f17})})):
                                coll58_ = coll58_.union(_dafny.Set([d_1103_v27_]))
                        return _dafny.Set(coll58_)
                    coll56_ = _dafny.Map()
                    def iife111_():
                        coll57_ = _dafny.Set()
                        compr_57_: _dafny.Map
                        for compr_57_ in (_dafny.Set({_dafny.Map({(self).f17: (self).f17})})).Elements:
                            d_1101_v27_: _dafny.Map = compr_57_
                            if (d_1101_v27_) in (_dafny.Set({_dafny.Map({(self).f17: (self).f17})})):
                                coll57_ = coll57_.union(_dafny.Set([d_1101_v27_]))
                        return _dafny.Set(coll57_)
                    compr_56_: _dafny.Map
                    for compr_56_ in (iife111_()
                    ).Elements:
                        d_1102_v26_: _dafny.Map = compr_56_
                        if (d_1102_v26_) in (iife112_()
                        ):
                            coll56_[d_1102_v26_] = d_1083_v10_.f23
                    return _dafny.Map(coll56_)
                rhs210_ = (d_1083_v10_).f22
                rhs211_ = ((_dafny.SeqWithoutIsStrInference([len(iife110_()
) for d_1104_i2_ in range(default__.abs(78))])) + (d_1098_v24_)) + ((_dafny.SeqWithoutIsStrInference([(d_1083_v10_).f22, (d_1083_v10_).f22, len(_dafny.SeqWithoutIsStrInference([d_1083_v10_.f23, d_1083_v10_.f23, d_1083_v10_.f23])), -313])).set(default__.safeIndex(len(d_1100_v28_), len(_dafny.SeqWithoutIsStrInference([(d_1083_v10_).f22, (d_1083_v10_).f22, len(_dafny.SeqWithoutIsStrInference([d_1083_v10_.f23, d_1083_v10_.f23, d_1083_v10_.f23])), -313]))), (d_1083_v10_).f22))
                rhs212_ = d_1099_v25_
                lhs166_ = d_1083_v10_
                lhs166_.f23 = rhs210_
                d_1098_v24_ = rhs211_
                d_1099_v25_ = rhs212_
                d_1105_v29_: C3
                nw173_ = C3()
                nw173_.ctor__((0) - ((self).fm3((self).f11, (self).f17, True, globalState)), (self).f11, default__.safeDivisionInt(len(d_1082_v9_), (0) - (len(d_1100_v28_))), (self).f12)
                d_1105_v29_ = nw173_
                d_1106_v30_: _dafny.Array
                nw174_ = _dafny.Array(None, 25)
                d_1106_v30_ = nw174_
                d_1107_v31_: _dafny.MultiSet
                d_1107_v31_ = _dafny.MultiSet([(self).f17])
                d_1108_v32_: T1
                nw175_ = C3()
                nw175_.ctor__((d_1107_v31_).cardinality, (d_1105_v29_).f22, len(d_1082_v9_), (self).f12)
                d_1108_v32_ = nw175_
                index137_ = default__.safeIndex(824, (d_1106_v30_).length(0))
                (d_1106_v30_)[index137_] = d_1108_v32_
                d_1109_v33_: _dafny.Seq
                d_1109_v33_ = _dafny.SeqWithoutIsStrInference([d_1082_v9_])
                d_1110_v34_: _dafny.Seq
                d_1110_v34_ = _dafny.SeqWithoutIsStrInference([d_1097_v23_, d_1097_v23_])
                d_1111_v35_: _dafny.Array
                nw176_ = _dafny.Array(None, 7)
                nw176_[int(0)] = d_1105_v29_.f23
                nw176_[int(1)] = (0) - (d_1083_v10_.f23)
                nw176_[int(2)] = (p0) + ((d_1083_v10_).f22)
                nw176_[int(3)] = (d_1083_v10_).f22
                nw176_[int(4)] = len(d_1109_v33_)
                nw176_[int(5)] = ((d_1083_v10_).f22 if (self).f17 else d_1083_v10_.f23)
                nw176_[int(6)] = default__.safeModuloInt((d_1108_v32_).f11, len(d_1110_v34_))
                d_1111_v35_ = nw176_
                d_1112_v36_: str
                d_1112_v36_ = _dafny.CodePoint('o')
                d_1113_v37_: _dafny.Map
                d_1113_v37_ = _dafny.Map({d_1112_v36_: (self).f11})
                index138_ = default__.safeIndex(95, (d_1111_v35_).length(0))
                (d_1111_v35_)[index138_] = len((d_1113_v37_) | (d_1113_v37_))
                d_1114_v39_: _dafny.Set
                d_1114_v39_ = _dafny.Set({(d_1082_v9_)[default__.safeIndex((d_1083_v10_).f22, len(d_1082_v9_))], d_1112_v36_, d_1112_v36_})
                d_1115_v40_: _dafny.Map
                def iife113_():
                    coll59_ = _dafny.Map()
                    compr_59_: str
                    for compr_59_ in (d_1114_v39_).Elements:
                        d_1116_v38_: str = compr_59_
                        if (d_1116_v38_) in (d_1114_v39_):
                            coll59_[d_1116_v38_] = (self).f17
                    return _dafny.Map(coll59_)
                d_1115_v40_ = _dafny.Map({iife113_()
                : (self).f17})
                d_1117_v41_: _dafny.Map
                d_1117_v41_ = _dafny.Map({d_1112_v36_: not((self).f17)})
                index139_ = default__.safeIndex(824, (d_1106_v30_).length(0))
                index140_ = default__.safeIndex(95, (d_1111_v35_).length(0))
                rhs213_ = d_1108_v32_
                rhs214_ = (0) - ((d_1108_v32_).f11)
                rhs215_ = len(d_1096_v22_)
                rhs216_ = (d_1115_v40_) | ((_dafny.Map({d_1117_v41_: (self).f17})).set((d_1117_v41_).set(d_1112_v36_, (self).f17), (self).f17))
                lhs167_ = d_1106_v30_
                lhs168_ = default__.safeIndex(824, (d_1106_v30_).length(0))
                lhs169_ = d_1083_v10_
                lhs170_ = d_1111_v35_
                lhs171_ = default__.safeIndex(95, (d_1111_v35_).length(0))
                lhs167_[lhs168_] = rhs213_
                lhs169_.f23 = rhs214_
                lhs170_[lhs171_] = rhs215_
                d_1115_v40_ = rhs216_
                d_1118_v42_: _dafny.Array
                nw177_ = _dafny.Array(False, 4)
                d_1118_v42_ = nw177_
                d_1119_v43_: D0
                d_1119_v43_ = D0_DC0((d_1111_v35_)[default__.safeIndex(95, (d_1111_v35_).length(0))])
                d_1120_v44_: _dafny.Set
                d_1120_v44_ = _dafny.Set({(d_1111_v35_)[default__.safeIndex(95, (d_1111_v35_).length(0))], (d_1119_v43_).cf0})
                index141_ = default__.safeIndex(522, (d_1118_v42_).length(0))
                (d_1118_v42_)[index141_] = (d_1083_v10_.f23) == (len(d_1120_v44_))
                index142_ = default__.safeIndex(522, (d_1118_v42_).length(0))
                (d_1118_v42_)[index142_] = (self).f17
                d_1121_v45_: _dafny.Map
                d_1121_v45_ = _dafny.Map({(0) - ((self).f11): ((self).f17 if (d_1118_v42_)[default__.safeIndex(522, (d_1118_v42_).length(0))] else (self).f17)})
                d_1121_v45_ = (d_1121_v45_).set(len(d_1082_v9_), default__.fm0(globalState))
            d_1122_v46_: C2
            nw178_ = C2()
            nw178_.ctor__(_dafny.CodePoint('g'), (self).f11)
            d_1122_v46_ = nw178_
            d_1123_v47_: _dafny.Seq
            d_1123_v47_ = _dafny.SeqWithoutIsStrInference([d_1083_v10_.f23, len(d_1082_v9_), d_1083_v10_.f23])
            (globalState).f6 = ((self).f17 if (default__.fm45(False, (self).f17, (d_1083_v10_).f22, len(d_1082_v9_), globalState)).ispropersubset(_dafny.MultiSet(d_1123_v47_)) else (self).f17)
            d_1124_v48_: D13
            d_1124_v48_ = D13_DC28(d_1082_v9_, -728, (self).f17)
            pat_let_tv37_ = globalState
            pat_let_tv38_ = d_1082_v9_
            def iife114_(_pat_let27_0):
                def iife115_(d_1125_dt__update__tmp_h0_):
                    def iife116_(_pat_let28_0):
                        def iife117_(d_1126_dt__update_hcf40_h0_):
                            return D13_DC28(d_1126_dt__update_hcf40_h0_, (d_1125_dt__update__tmp_h0_).cf41, (d_1125_dt__update__tmp_h0_).cf42)
                        return iife117_(_pat_let28_0)
                    return iife116_((default__.fm1(pat_let_tv37_)) + (pat_let_tv38_))
                return iife115_(_pat_let27_0)
            d_1124_v48_ = iife114_(D13_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsb")), (self).f11, (self).f17))
            (globalState).f6 = not(not(((self).f17) or (True)))
        elif True:
            (globalState).f1 = default__.safeDivisionInt((d_1083_v10_).f22, p0)
            d_1127_v49_: str
            d_1127_v49_ = _dafny.CodePoint('u')
            d_1128_v50_: _dafny.Seq
            d_1128_v50_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])
            d_1129_v51_: _dafny.Seq
            d_1129_v51_ = _dafny.SeqWithoutIsStrInference([d_1082_v9_, d_1082_v9_, _dafny.SeqWithoutIsStrInference([d_1127_v49_ for d_1130_i3_ in range(default__.abs(730))])])
            d_1131_v52_: _dafny.Map
            d_1131_v52_ = _dafny.Map({(self).f17: d_1083_v10_.f23})
            d_1132_v53_: _dafny.Seq
            d_1132_v53_ = _dafny.SeqWithoutIsStrInference([d_1083_v10_.f23, (d_1083_v10_).f22])
            d_1127_v49_ = default__.fm20((d_1128_v50_) + (d_1128_v50_), default__.safeModuloInt((d_1083_v10_).f22, d_1083_v10_.f23), (_dafny.SeqWithoutIsStrInference([len((d_1129_v51_)[default__.safeIndex(len(d_1131_v52_), len(d_1129_v51_))])])) == (d_1132_v53_), globalState)
            d_1133_v54_: _dafny.Array
            def lambda39_(d_1134_i4_):
                return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f11]))

            init20_ = lambda39_
            nw179_ = _dafny.Array(None, 13)
            for i0_20_ in range(nw179_.length(0)):
                nw179_[i0_20_] = init20_(i0_20_)
            d_1133_v54_ = nw179_
            d_1133_v54_ = d_1133_v54_
            d_1135_v55_: _dafny.Map
            d_1135_v55_ = _dafny.Map({(self).f11: (self).f17})
            d_1136_v56_: _dafny.Seq
            d_1136_v56_ = _dafny.SeqWithoutIsStrInference([d_1135_v55_])
            d_1137_v57_: C3
            nw180_ = C3()
            nw180_.ctor__(len((d_1136_v56_)[default__.safeIndex((d_1083_v10_).f22, len(d_1136_v56_))]), (self).f11, d_1083_v10_.f23, (self).f12)
            d_1137_v57_ = nw180_
            rhs217_ = 24
            rhs218_ = not(((self).f17) and (not((self).f17)))
            lhs172_ = d_1083_v10_
            lhs173_ = globalState
            lhs172_.f23 = rhs217_
            lhs173_.f7 = rhs218_
        d_1138_v58_: _dafny.Seq
        d_1138_v58_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])
        d_1139_v59_: _dafny.Seq
        d_1139_v59_ = _dafny.SeqWithoutIsStrInference([len(d_1138_v58_), d_1083_v10_.f23, (self).f11])
        d_1140_v60_: _dafny.MultiSet
        d_1140_v60_ = _dafny.MultiSet([(self).f17])
        d_1141_v61_: _dafny.MultiSet
        d_1141_v61_ = d_1140_v60_
        d_1142_v62_: _dafny.Array
        nw181_ = _dafny.Array(None, 16)
        nw181_[int(0)] = (self).f11
        nw181_[int(1)] = (d_1139_v59_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0])), len(d_1139_v59_))]
        nw181_[int(2)] = p0
        nw181_[int(3)] = ((d_1083_v10_).f22) - ((self).f11)
        nw181_[int(4)] = (0) - (len(d_1138_v58_))
        nw181_[int(5)] = (self).f11
        nw181_[int(6)] = (self).f11
        nw181_[int(7)] = ((self).f11) + (p0)
        nw181_[int(8)] = ((d_1141_v61_)).cardinality
        nw181_[int(9)] = (self).f11
        nw181_[int(10)] = d_1083_v10_.f23
        nw181_[int(11)] = len(default__.fm11(globalState))
        nw181_[int(12)] = (d_1083_v10_).f22
        nw181_[int(13)] = (d_1083_v10_).fm3(p0, (self).f17, (self).f17, globalState)
        nw181_[int(14)] = p0
        nw181_[int(15)] = (self).f11
        d_1142_v62_ = nw181_
        index143_ = default__.safeIndex(641, (d_1142_v62_).length(0))
        (d_1142_v62_)[index143_] = 255
        index144_ = default__.safeIndex(641, (d_1142_v62_).length(0))
        rhs219_ = (0) - (((d_1140_v60_)[(self).f17] if ((self).f17) in (d_1140_v60_) else d_1083_v10_.f23))
        rhs220_ = (d_1083_v10_).f22
        rhs221_ = 297
        lhs174_ = d_1142_v62_
        lhs175_ = default__.safeIndex(641, (d_1142_v62_).length(0))
        lhs176_ = d_1083_v10_
        lhs177_ = d_1083_v10_
        lhs174_[lhs175_] = rhs219_
        lhs176_.f23 = rhs220_
        lhs177_.f23 = rhs221_
        index145_ = default__.safeIndex(641, (d_1142_v62_).length(0))
        (d_1142_v62_)[index145_] = 600
        (d_1083_v10_).f23 = (d_1083_v10_).f22
        d_1143_v63_: _dafny.Array
        nw182_ = _dafny.Array(_dafny.MultiSet({}), 18)
        d_1143_v63_ = nw182_
        r0 = d_1143_v63_
        d_1144_v64_: _dafny.Map
        d_1144_v64_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1145_i5_ in range(default__.abs(825))]): (self).f16})
        r1 = d_1144_v64_
        return r0, r1

    def m5(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        if p1:
            d_1146_v0_: _dafny.Array
            def lambda40_(d_1147_p1_):
                def lambda41_(d_1148_i0_):
                    return d_1147_p1_

                return lambda41_

            init21_ = lambda40_(p1)
            nw183_ = _dafny.Array(None, 15)
            for i0_21_ in range(nw183_.length(0)):
                nw183_[i0_21_] = init21_(i0_21_)
            d_1146_v0_ = nw183_
            d_1149_v1_: _dafny.Set
            d_1149_v1_ = _dafny.Set({default__.fm46(globalState)})
            d_1150_v2_: _dafny.Map
            d_1150_v2_ = _dafny.Map({_dafny.MultiSet([d_1146_v0_]): d_1149_v1_})
            d_1151_v3_: _dafny.MultiSet
            d_1151_v3_ = _dafny.MultiSet([d_1146_v0_, d_1146_v0_, d_1146_v0_])
            d_1150_v2_ = (d_1150_v2_).set(d_1151_v3_, (d_1149_v1_) | (d_1149_v1_))
            d_1152_v5_: _dafny.Map
            d_1152_v5_ = _dafny.Map({p0: p1})
            d_1153_v6_: D8
            d_1153_v6_ = D8_DC15(d_1152_v5_)
            d_1154_v7_: D8
            d_1154_v7_ = D8_DC17(d_1153_v6_)
            d_1155_v8_: D13
            def iife118_():
                coll60_ = _dafny.Map()
                compr_60_: D8
                for compr_60_ in (_dafny.SeqWithoutIsStrInference([d_1154_v7_])).Elements:
                    d_1156_v4_: D8 = compr_60_
                    if (d_1156_v4_) in (_dafny.SeqWithoutIsStrInference([d_1154_v7_])):
                        coll60_[d_1156_v4_] = (self).f11
                return _dafny.Map(coll60_)
            d_1155_v8_ = D13_DC27(iife118_()
)
            d_1157_v9_: D13
            d_1157_v9_ = D13_DC29(d_1155_v8_)
            d_1158_v10_: D13
            d_1158_v10_ = D13_DC29(d_1155_v8_)
            d_1159_v11_: _dafny.Map
            d_1159_v11_ = _dafny.Map({d_1154_v7_: 992})
            d_1158_v10_ = D13_DC29(D13_DC27(d_1159_v11_))
            (globalState).f7 = p1
            d_1160_v12_: _dafny.Seq
            d_1160_v12_ = _dafny.SeqWithoutIsStrInference([p1, p0, p1])
            d_1161_v13_: _dafny.Seq
            d_1161_v13_ = _dafny.SeqWithoutIsStrInference([not((_dafny.SeqWithoutIsStrInference([p1])) <= (d_1160_v12_)), p0, (not(p0) if p0 else p0), not(((self).f16).issubset((self).f16)), p0])
            d_1162_v14_: _dafny.Map
            d_1162_v14_ = _dafny.Map({p0: d_1152_v5_})
            d_1163_v15_: _dafny.Seq
            d_1163_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aw"))
            d_1160_v12_ = default__.fm12(default__.safeModuloInt((0) - ((self).f11), len(d_1162_v14_)), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1164_i1_ in range(default__.abs(-390))])) + (d_1163_v15_), globalState)
            d_1165_v16_: C3
            nw184_ = C3()
            nw184_.ctor__((self).f11, (0) - (((self).f11) - (267)), p2, (self).f12)
            d_1165_v16_ = nw184_
        elif True:
            d_1166_v17_: C3
            nw185_ = C3()
            nw185_.ctor__(p2, -116, (((self).f16)[p2] if (p2) in ((self).f16) else p2), (self).f12)
            d_1166_v17_ = nw185_
            d_1167_v18_: _dafny.Array
            def lambda42_(d_1168_p1_, d_1169_p0_):
                def lambda43_(d_1170_i2_):
                    return _dafny.Map({D11_DC24(2, _dafny.SeqWithoutIsStrInference([d_1168_p1_, d_1169_p0_]), _dafny.CodePoint('n')): d_1168_p1_})

                return lambda43_

            init22_ = lambda42_(p1, p0)
            nw186_ = _dafny.Array(None, 24)
            for i0_22_ in range(nw186_.length(0)):
                nw186_[i0_22_] = init22_(i0_22_)
            d_1167_v18_ = nw186_
            d_1171_v19_: str
            d_1171_v19_ = _dafny.CodePoint('f')
            d_1172_v20_: D11
            d_1172_v20_ = D11_DC24(p2, _dafny.SeqWithoutIsStrInference([p1, p1]), d_1171_v19_)
            d_1173_v21_: _dafny.Map
            d_1173_v21_ = _dafny.Map({d_1172_v20_: p1})
            index146_ = default__.safeIndex(326, (d_1167_v18_).length(0))
            (d_1167_v18_)[index146_] = d_1173_v21_
            index147_ = default__.safeIndex(326, (d_1167_v18_).length(0))
            (d_1167_v18_)[index147_] = (_dafny.Map({d_1172_v20_: p1})) | ((d_1173_v21_) | (d_1173_v21_))
            d_1174_v22_: D0
            d_1174_v22_ = D0_DC1((self).f17, d_1166_v17_.f23, not(p1))
            d_1175_v23_: _dafny.Seq
            d_1175_v23_ = _dafny.SeqWithoutIsStrInference([(d_1166_v17_).f22])
            d_1176_v24_: D7
            d_1176_v24_ = D7_DC12(d_1175_v23_)
            d_1177_v25_: _dafny.Map
            d_1177_v25_ = _dafny.Map({default__.fm39((d_1166_v17_).f22, not(p1), (d_1174_v22_).cf2, globalState): d_1176_v24_})
            d_1177_v25_ = (d_1177_v25_).set(d_1171_v19_, d_1176_v24_)
            (d_1166_v17_).f23 = (0) - ((d_1166_v17_).f22)
            d_1178_v26_: C2
            nw187_ = C2()
            nw187_.ctor__(d_1171_v19_, -968)
            d_1178_v26_ = nw187_
        d_1179_v27_: _dafny.Seq
        d_1179_v27_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1180_v28_: D6
        d_1180_v28_ = D6_DC10(d_1179_v27_)
        pat_let_tv39_ = p2
        pat_let_tv40_ = p2
        pat_let_tv41_ = p2
        def lambda44_(source17_):
            if source17_.is_DC10:
                d_1181___mcc_h0_ = source17_.cf13
                d_1182_cf13_ = d_1181___mcc_h0_
                return pat_let_tv39_
            elif source17_.is_DC11:
                d_1183___mcc_h1_ = source17_.cf14
                d_1184___mcc_h2_ = source17_.cf15
                d_1185___mcc_h3_ = source17_.cf16
                d_1186_cf16_ = d_1185___mcc_h3_
                d_1187_cf15_ = d_1184___mcc_h2_
                d_1188_cf14_ = d_1183___mcc_h1_
                return default__.safeDivisionInt(437, (self).f11)
            elif True:
                d_1189___mcc_h4_ = source17_.cf12
                d_1190_cf12_ = d_1189___mcc_h4_
                return (len(_dafny.Map({(self).f17: pat_let_tv40_}))) * (pat_let_tv41_)

        r0 = (0) - (lambda44_(d_1180_v28_))
        r0 = p2
        d_1191_v29_: _dafny.Array
        def lambda45_(d_1192_i3_):
            return (self).f17

        init23_ = lambda45_
        nw188_ = _dafny.Array(None, 9)
        for i0_23_ in range(nw188_.length(0)):
            nw188_[i0_23_] = init23_(i0_23_)
        d_1191_v29_ = nw188_
        nw189_ = _dafny.Array(False, 5)
        d_1191_v29_ = nw189_
        d_1193_v30_: _dafny.Array
        nw190_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_1193_v30_ = nw190_
        d_1194_v31_: _dafny.Array
        def lambda46_(d_1195_i4_):
            return D2_DC5()

        init24_ = lambda46_
        nw191_ = _dafny.Array(None, 6)
        for i0_24_ in range(nw191_.length(0)):
            nw191_[i0_24_] = init24_(i0_24_)
        d_1194_v31_ = nw191_
        index148_ = default__.safeIndex(17, (d_1193_v30_).length(0))
        (d_1193_v30_)[index148_] = d_1194_v31_
        index149_ = default__.safeIndex(17, (d_1193_v30_).length(0))
        rhs222_ = d_1194_v31_
        rhs223_ = p1
        rhs224_ = (self).f17
        rhs225_ = p2
        rhs226_ = 93
        lhs178_ = d_1193_v30_
        lhs179_ = default__.safeIndex(17, (d_1193_v30_).length(0))
        lhs180_ = globalState
        lhs181_ = globalState
        lhs178_[lhs179_] = rhs222_
        r1 = rhs223_
        lhs180_.f6 = rhs224_
        r2 = rhs225_
        lhs181_.f1 = rhs226_
        d_1196_i5_: int
        d_1196_i5_ = 0
        with _dafny.label("13"):
            while (self).f17:
                with _dafny.c_label("13"):
                    if (d_1196_i5_) >= (100):
                        raise _dafny.Break("13")
                    d_1196_i5_ = (d_1196_i5_) + (1)
                    (globalState).f1 = (self).f11
                    d_1197_v32_: C0
                    nw192_ = C0()
                    nw192_.ctor__((d_1180_v28_).cf13, (self).f11)
                    d_1197_v32_ = nw192_
                    d_1198_v33_: _dafny.Map
                    d_1198_v33_ = _dafny.Map({p2: d_1197_v32_})
                    d_1199_v34_: _dafny.Array
                    nw193_ = _dafny.Array(int(0), 16)
                    d_1199_v34_ = nw193_
                    d_1200_v35_: _dafny.Map
                    d_1200_v35_ = _dafny.Map({len(d_1198_v33_): d_1199_v34_})
                    d_1200_v35_ = (d_1200_v35_).set(p2, d_1199_v34_)
                    d_1201_v36_: str
                    d_1201_v36_ = _dafny.CodePoint('b')
                    d_1201_v36_ = d_1201_v36_
                    (globalState).f1 = ((0) - ((len(_dafny.Set({593, p2}))) - (p2))) + (len((d_1197_v32_).f20))
                    pass
            pass
        d_1202_v37_: _dafny.MultiSet
        d_1202_v37_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f17]), _dafny.SeqWithoutIsStrInference([p1, (self).f17]), d_1179_v27_, d_1179_v27_])
        r0 = (0) - (((d_1202_v37_).set(_dafny.SeqWithoutIsStrInference([(self).f17, p0, True]), default__.abs((0) - ((self).f11)))).cardinality)
        r1 = (p0) not in (((self).f12) - (default__.fm41(globalState)))
        d_1203_v38_: _dafny.MultiSet
        d_1203_v38_ = _dafny.MultiSet([p0, True, False, p1])
        d_1204_v39_: _dafny.MultiSet
        d_1204_v39_ = _dafny.MultiSet([d_1203_v38_])
        r2 = (((d_1204_v39_)[((d_1203_v38_).set(True, default__.abs(p2))).set(False, default__.abs(p2))] if (((d_1203_v38_).set(True, default__.abs(p2))).set(False, default__.abs(p2))) in (d_1204_v39_) else p2) if p1 else 14)
        return r0, r1, r2

    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C5:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def fm9(self, p0, p1, p2, globalState):
        return (0) - (((539) + ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usdxyfb"))))))) - (933))

    def m4(self, p0, p1, globalState):
        d_1205_v0_: int
        d_1205_v0_ = 826
        (globalState).f1 = d_1205_v0_
        d_1206_v1_: bool
        d_1206_v1_ = True
        d_1207_v2_: str
        d_1207_v2_ = _dafny.CodePoint('k')
        d_1208_v3_: _dafny.Map
        d_1208_v3_ = _dafny.Map({d_1206_v1_: (_dafny.MultiSet([len(_dafny.Map({d_1205_v0_: d_1207_v2_})), d_1205_v0_])).ispropersubset(_dafny.MultiSet([(0) - (d_1205_v0_)]))})
        if ((d_1208_v3_)[d_1206_v1_] if (d_1206_v1_) in (d_1208_v3_) else d_1206_v1_):
            d_1209_v4_: _dafny.Map
            d_1209_v4_ = _dafny.Map({d_1205_v0_: 821})
            d_1209_v4_ = (d_1209_v4_).set(d_1205_v0_, (0) - (d_1205_v0_))
            d_1210_v5_: _dafny.Seq
            d_1210_v5_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_])
            (globalState).f6 = (d_1206_v1_) and ((d_1210_v5_)[default__.safeIndex(d_1205_v0_, len(d_1210_v5_))])
            d_1211_v6_: _dafny.Map
            d_1211_v6_ = _dafny.Map({d_1207_v2_: d_1205_v0_})
            d_1212_v7_: C0
            nw194_ = C0()
            nw194_.ctor__(d_1210_v5_, ((d_1211_v6_)[d_1207_v2_] if (d_1207_v2_) in (d_1211_v6_) else d_1205_v0_))
            d_1212_v7_ = nw194_
            d_1213_v8_: _dafny.Array
            nw195_ = _dafny.Array(None, 28)
            d_1213_v8_ = nw195_
            d_1214_v9_: _dafny.Set
            d_1214_v9_ = _dafny.Set({False})
            d_1215_v10_: T0
            nw196_ = C3()
            nw196_.ctor__(611, (0) - (d_1205_v0_), d_1205_v0_, d_1214_v9_)
            d_1215_v10_ = nw196_
            index150_ = default__.safeIndex(629, (d_1213_v8_).length(0))
            (d_1213_v8_)[index150_] = d_1215_v10_
            index151_ = default__.safeIndex(629, (d_1213_v8_).length(0))
            (d_1213_v8_)[index151_] = d_1215_v10_
            (globalState).f1 = len(((p0) + (p0)) + (p0))
        elif True:
            (globalState).f1 = d_1205_v0_
            d_1216_v11_: _dafny.Seq
            d_1216_v11_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(d_1205_v0_, d_1205_v0_), d_1205_v0_])
            d_1216_v11_ = _dafny.SeqWithoutIsStrInference([d_1205_v0_])
            d_1217_v12_: C3
            nw197_ = C3()
            nw197_.ctor__(d_1205_v0_, d_1205_v0_, d_1205_v0_, _dafny.Set({d_1206_v1_}))
            d_1217_v12_ = nw197_
            d_1218_v13_: _dafny.Array
            nw198_ = _dafny.Array(None, 17)
            nw198_[int(0)] = d_1217_v12_
            nw198_[int(1)] = d_1217_v12_
            nw198_[int(2)] = d_1217_v12_
            nw198_[int(3)] = (d_1217_v12_ if d_1206_v1_ else d_1217_v12_)
            nw198_[int(4)] = d_1217_v12_
            nw198_[int(5)] = d_1217_v12_
            nw198_[int(6)] = d_1217_v12_
            nw198_[int(7)] = d_1217_v12_
            nw198_[int(8)] = d_1217_v12_
            nw198_[int(9)] = d_1217_v12_
            nw198_[int(10)] = d_1217_v12_
            nw198_[int(11)] = d_1217_v12_
            nw198_[int(12)] = (d_1217_v12_ if not(not(not(True))) else d_1217_v12_)
            nw198_[int(13)] = d_1217_v12_
            nw198_[int(14)] = (d_1217_v12_ if d_1206_v1_ else d_1217_v12_)
            nw198_[int(15)] = d_1217_v12_
            nw198_[int(16)] = d_1217_v12_
            d_1218_v13_ = nw198_
            index152_ = default__.safeIndex(819, (d_1218_v13_).length(0))
            (d_1218_v13_)[index152_] = d_1217_v12_
            index153_ = default__.safeIndex(819, (d_1218_v13_).length(0))
            (d_1218_v13_)[index153_] = (d_1217_v12_ if default__.fm0(globalState) else d_1217_v12_)
            d_1219_v14_: D0
            d_1219_v14_ = D0_DC0(d_1217_v12_.f23)
            pat_let_tv42_ = d_1205_v0_
            def iife119_(_pat_let29_0):
                def iife120_(d_1220_dt__update__tmp_h0_):
                    def iife121_(_pat_let30_0):
                        def iife122_(d_1221_dt__update_hcf0_h0_):
                            return D0_DC0(d_1221_dt__update_hcf0_h0_)
                        return iife122_(_pat_let30_0)
                    return iife121_(pat_let_tv42_)
                return iife120_(_pat_let29_0)
            d_1219_v14_ = iife119_(d_1219_v14_)
            d_1222_v15_: _dafny.Seq
            d_1222_v15_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_, d_1206_v1_])
            d_1223_v16_: _dafny.Seq
            d_1223_v16_ = _dafny.SeqWithoutIsStrInference([d_1222_v15_])
            (globalState).f1 = (len(d_1223_v16_)) * ((0) - (d_1205_v0_))
        d_1224_v17_: D13
        d_1224_v17_ = D13_DC28(p0, d_1205_v0_, d_1206_v1_)
        if (d_1224_v17_).cf42:
            d_1225_v18_: _dafny.Array
            def lambda47_(d_1226_v3_):
                def lambda48_(d_1227_i0_):
                    return (d_1226_v3_) | (d_1226_v3_)

                return lambda48_

            init25_ = lambda47_(d_1208_v3_)
            nw199_ = _dafny.Array(None, 16)
            for i0_25_ in range(nw199_.length(0)):
                nw199_[i0_25_] = init25_(i0_25_)
            d_1225_v18_ = nw199_
            d_1225_v18_ = d_1225_v18_
            (globalState).f1 = 498
            d_1228_v19_: _dafny.MultiSet
            d_1228_v19_ = _dafny.MultiSet([default__.fm0(globalState), d_1206_v1_])
            (globalState).f7 = not ((d_1228_v19_).isdisjoint(d_1228_v19_)) or (d_1206_v1_)
            (globalState).f6 = d_1206_v1_
            d_1229_v20_: _dafny.Array
            nw200_ = _dafny.Array(int(0), 24)
            d_1229_v20_ = nw200_
            index154_ = default__.safeIndex(665, (d_1229_v20_).length(0))
            (d_1229_v20_)[index154_] = d_1205_v0_
            d_1230_v21_: _dafny.Seq
            d_1230_v21_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_])
            d_1231_v22_: _dafny.MultiSet
            d_1231_v22_ = _dafny.MultiSet([d_1230_v21_, d_1230_v21_, d_1230_v21_, ((d_1230_v21_).set(default__.safeIndex(len(p0), len(d_1230_v21_)), d_1206_v1_)) + ((d_1230_v21_).set(default__.safeIndex(d_1205_v0_, len(d_1230_v21_)), d_1206_v1_))])
            index155_ = default__.safeIndex(665, (d_1229_v20_).length(0))
            (d_1229_v20_)[index155_] = ((d_1231_v22_)[(d_1230_v21_ if not(d_1206_v1_) else _dafny.SeqWithoutIsStrInference([d_1206_v1_]))] if ((d_1230_v21_ if not(d_1206_v1_) else _dafny.SeqWithoutIsStrInference([d_1206_v1_]))) in (d_1231_v22_) else d_1205_v0_)
        elif True:
            d_1232_v23_: _dafny.Array
            def lambda49_(d_1233_v0_):
                def lambda50_(d_1234_i1_):
                    return default__.safeDivisionInt(d_1234_i1_, d_1233_v0_)

                return lambda50_

            init26_ = lambda49_(d_1205_v0_)
            nw201_ = _dafny.Array(None, 6)
            for i0_26_ in range(nw201_.length(0)):
                nw201_[i0_26_] = init26_(i0_26_)
            d_1232_v23_ = nw201_
            d_1235_v24_: _dafny.Map
            d_1235_v24_ = _dafny.Map({d_1232_v23_: True})
            d_1236_v25_: D10
            d_1236_v25_ = D10_DC20(d_1235_v24_)
            source18_ = d_1236_v25_
            if source18_.is_DC21:
                d_1237___mcc_h0_ = source18_.cf28
                d_1238___mcc_h1_ = source18_.cf29
                d_1239_cf29_ = d_1238___mcc_h1_
                d_1240_cf28_ = d_1237___mcc_h0_
                d_1241_v26_: _dafny.MultiSet
                d_1241_v26_ = _dafny.MultiSet([default__.fm2(d_1205_v0_, len(p0), globalState)])
                d_1241_v26_ = d_1241_v26_
                d_1242_v27_: _dafny.Seq
                d_1242_v27_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_])
                d_1243_v28_: D6
                d_1243_v28_ = D6_DC10(d_1242_v27_)
                d_1243_v28_ = (d_1243_v28_ if ((d_1208_v3_)[d_1239_cf29_] if (d_1239_cf29_) in (d_1208_v3_) else default__.fm0(globalState)) else d_1243_v28_)
                d_1244_v29_: _dafny.Array
                def lambda51_(d_1245_cf28_, d_1246_p0_, d_1247_v1_):
                    def lambda52_(d_1248_i2_):
                        return ((D14_DC30(_dafny.Map({d_1245_cf28_: d_1246_p0_}))).cf44) | (_dafny.Map({d_1247_v1_: d_1246_p0_}))

                    return lambda52_

                init27_ = lambda51_(d_1240_cf28_, p0, d_1206_v1_)
                nw202_ = _dafny.Array(None, 29)
                for i0_27_ in range(nw202_.length(0)):
                    nw202_[i0_27_] = init27_(i0_27_)
                d_1244_v29_ = nw202_
                d_1249_v30_: _dafny.Map
                d_1249_v30_ = _dafny.Map({d_1206_v1_: p0})
                index156_ = default__.safeIndex(956, (d_1244_v29_).length(0))
                (d_1244_v29_)[index156_] = d_1249_v30_
                d_1250_v31_: _dafny.Seq
                d_1250_v31_ = _dafny.SeqWithoutIsStrInference([d_1249_v30_])
                index157_ = default__.safeIndex(956, (d_1244_v29_).length(0))
                (d_1244_v29_)[index157_] = (d_1250_v31_)[default__.safeIndex(d_1205_v0_, len(d_1250_v31_))]
                d_1251_v32_: _dafny.Set
                d_1251_v32_ = _dafny.Set({len((p0) + (p0))})
                d_1252_v33_: _dafny.Map
                d_1252_v33_ = _dafny.Map({d_1205_v0_: d_1205_v0_})
                d_1253_v34_: _dafny.Map
                d_1253_v34_ = _dafny.Map({d_1207_v2_: default__.fm0(globalState)})
                d_1254_v36_: _dafny.Map
                d_1254_v36_ = _dafny.Map({d_1207_v2_: d_1205_v0_})
                d_1255_v37_: _dafny.Array
                nw203_ = _dafny.Array(None, 29)
                nw203_[int(0)] = not (d_1240_cf28_) or (d_1240_cf28_)
                nw203_[int(1)] = d_1206_v1_
                nw203_[int(2)] = ((self).fm9((0) - (len(d_1252_v33_)), -717, d_1206_v1_, globalState)) < (d_1205_v0_)
                def iife123_():
                    coll61_ = _dafny.Map()
                    compr_61_: str
                    for compr_61_ in (d_1254_v36_).keys.Elements:
                        d_1256_v35_: str = compr_61_
                        if (d_1256_v35_) in (d_1254_v36_):
                            coll61_[d_1256_v35_] = d_1239_cf29_
                    return _dafny.Map(coll61_)
                nw203_[int(3)] = (d_1253_v34_) not in (_dafny.SeqWithoutIsStrInference([iife123_()
                , d_1253_v34_]))
                nw203_[int(4)] = not((d_1242_v27_)[default__.safeIndex(d_1205_v0_, len(d_1242_v27_))])
                nw203_[int(5)] = d_1240_cf28_
                nw203_[int(6)] = (d_1239_cf29_ if not(d_1239_cf29_) else (d_1242_v27_)[default__.safeIndex(d_1205_v0_, len(d_1242_v27_))])
                nw203_[int(7)] = True
                nw203_[int(8)] = False
                nw203_[int(9)] = (d_1240_cf28_ if d_1206_v1_ else d_1206_v1_)
                nw203_[int(10)] = not(d_1240_cf28_)
                nw203_[int(11)] = not (d_1239_cf29_) or (default__.fm0(globalState))
                nw203_[int(12)] = not((p0) < (p0))
                nw203_[int(13)] = d_1206_v1_
                nw203_[int(14)] = False
                nw203_[int(15)] = d_1206_v1_
                nw203_[int(16)] = (d_1206_v1_) and (d_1240_cf28_)
                nw203_[int(17)] = d_1206_v1_
                nw203_[int(18)] = d_1240_cf28_
                nw203_[int(19)] = (d_1239_cf29_ if d_1206_v1_ else d_1206_v1_)
                nw203_[int(20)] = True
                nw203_[int(21)] = ((d_1208_v3_)[d_1206_v1_] if (d_1206_v1_) in (d_1208_v3_) else d_1206_v1_)
                nw203_[int(22)] = (p0) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ley")))
                nw203_[int(23)] = d_1240_cf28_
                nw203_[int(24)] = d_1240_cf28_
                nw203_[int(25)] = not(d_1206_v1_)
                nw203_[int(26)] = (d_1206_v1_) and (d_1240_cf28_)
                nw203_[int(27)] = (d_1240_cf28_ if d_1240_cf28_ else True)
                nw203_[int(28)] = d_1239_cf29_
                d_1255_v37_ = nw203_
                index158_ = default__.safeIndex(645, (d_1255_v37_).length(0))
                (d_1255_v37_)[index158_] = d_1206_v1_
                index159_ = default__.safeIndex(645, (d_1255_v37_).length(0))
                rhs227_ = (d_1251_v32_) - (d_1251_v32_)
                rhs228_ = default__.fm0(globalState)
                lhs182_ = d_1255_v37_
                lhs183_ = default__.safeIndex(645, (d_1255_v37_).length(0))
                d_1251_v32_ = rhs227_
                lhs182_[lhs183_] = rhs228_
            elif True:
                d_1257___mcc_h2_ = source18_.cf27
                d_1258_cf27_ = d_1257___mcc_h2_
                d_1259_v38_: _dafny.Seq
                d_1259_v38_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_])
                d_1260_v39_: C1
                nw204_ = C1()
                nw204_.ctor__(d_1206_v1_, (d_1259_v38_)[default__.safeIndex(len(p0), len(d_1259_v38_))])
                d_1260_v39_ = nw204_
                (globalState).f7 = (d_1260_v39_).f19
                d_1261_v40_: D8
                d_1261_v40_ = D8_DC16(d_1206_v1_, (self).fm9(d_1205_v0_, d_1205_v0_, d_1206_v1_, globalState))
                d_1262_v41_: D8
                d_1262_v41_ = D8_DC17(d_1261_v40_)
                d_1263_v42_: D8
                d_1263_v42_ = D8_DC17(d_1261_v40_)
                d_1264_v43_: _dafny.Map
                d_1264_v43_ = _dafny.Map({d_1263_v42_: d_1205_v0_})
                d_1265_v44_: D13
                d_1265_v44_ = D13_DC27(d_1264_v43_)
                d_1265_v44_ = d_1265_v44_
                (globalState).f7 = True
            (globalState).f1 = d_1205_v0_
            d_1266_v45_: _dafny.Seq
            d_1266_v45_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1205_v0_)])
            d_1206_v1_ = (d_1266_v45_) < (_dafny.SeqWithoutIsStrInference([d_1205_v0_]))
            d_1267_v46_: _dafny.Seq
            d_1267_v46_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_, not(d_1206_v1_), d_1206_v1_, d_1206_v1_])
            d_1268_v47_: _dafny.Seq
            d_1268_v47_ = d_1267_v46_
            source19_ = d_1267_v46_
            d_1269___mcc_h3_ = source19_
            d_1270_cf10_ = d_1269___mcc_h3_
            index160_ = default__.safeIndex(251, (d_1232_v23_).length(0))
            (d_1232_v23_)[index160_] = d_1205_v0_
            index161_ = default__.safeIndex(251, (d_1232_v23_).length(0))
            rhs229_ = d_1205_v0_
            rhs230_ = d_1205_v0_
            rhs231_ = (0) - (len(d_1266_v45_))
            lhs184_ = globalState
            lhs185_ = globalState
            lhs186_ = d_1232_v23_
            lhs187_ = default__.safeIndex(251, (d_1232_v23_).length(0))
            lhs184_.f1 = rhs229_
            lhs185_.f1 = rhs230_
            lhs186_[lhs187_] = rhs231_
            d_1271_v48_: D8
            d_1271_v48_ = D8_DC16(d_1206_v1_, d_1205_v0_)
            d_1271_v48_ = d_1271_v48_
            d_1272_v49_: C1
            nw205_ = C1()
            nw205_.ctor__(d_1206_v1_, d_1206_v1_)
            d_1272_v49_ = nw205_
            d_1273_v50_: _dafny.Map
            d_1273_v50_ = _dafny.Map({(d_1232_v23_)[default__.safeIndex(251, (d_1232_v23_).length(0))]: d_1205_v0_})
            d_1274_v51_: C1
            nw206_ = C1()
            nw206_.ctor__((p0) <= (p0), (d_1273_v50_) != (d_1273_v50_))
            d_1274_v51_ = nw206_
            d_1275_v52_: _dafny.Map
            d_1275_v52_ = _dafny.Map({d_1205_v0_: d_1232_v23_})
            d_1275_v52_ = (d_1275_v52_).set(d_1205_v0_, d_1232_v23_)
        d_1276_i3_: int
        d_1276_i3_ = 0
        with _dafny.label("14"):
            while not((d_1205_v0_) != (d_1205_v0_)):
                with _dafny.c_label("14"):
                    if (d_1276_i3_) >= (100):
                        raise _dafny.Break("14")
                    d_1276_i3_ = (d_1276_i3_) + (1)
                    d_1277_v53_: _dafny.Array
                    nw207_ = _dafny.Array(_dafny.Map({}), 21)
                    d_1277_v53_ = nw207_
                    index162_ = default__.safeIndex(440, (d_1277_v53_).length(0))
                    (d_1277_v53_)[index162_] = d_1208_v3_
                    index163_ = default__.safeIndex(440, (d_1277_v53_).length(0))
                    (d_1277_v53_)[index163_] = default__.fm11(globalState)
                    if not (d_1206_v1_) or ((d_1206_v1_) == (d_1206_v1_)):
                        d_1278_v54_: C2
                        nw208_ = C2()
                        nw208_.ctor__(d_1207_v2_, d_1205_v0_)
                        d_1278_v54_ = nw208_
                        d_1279_v55_: _dafny.Map
                        d_1279_v55_ = _dafny.Map({d_1206_v1_: d_1278_v54_})
                        d_1280_v56_: _dafny.Array
                        def lambda53_(d_1281_v1_):
                            def lambda54_(d_1282_i4_):
                                return d_1281_v1_

                            return lambda54_

                        init28_ = lambda53_(d_1206_v1_)
                        nw209_ = _dafny.Array(None, 14)
                        for i0_28_ in range(nw209_.length(0)):
                            nw209_[i0_28_] = init28_(i0_28_)
                        d_1280_v56_ = nw209_
                        d_1283_v57_: _dafny.Array
                        nw210_ = _dafny.Array(None, 13)
                        nw210_[int(0)] = d_1280_v56_
                        nw210_[int(1)] = d_1280_v56_
                        nw210_[int(2)] = d_1280_v56_
                        nw210_[int(3)] = d_1280_v56_
                        nw210_[int(4)] = d_1280_v56_
                        nw210_[int(5)] = d_1280_v56_
                        nw210_[int(6)] = d_1280_v56_
                        nw210_[int(7)] = d_1280_v56_
                        nw210_[int(8)] = d_1280_v56_
                        nw210_[int(9)] = d_1280_v56_
                        nw210_[int(10)] = d_1280_v56_
                        nw210_[int(11)] = d_1280_v56_
                        nw210_[int(12)] = d_1280_v56_
                        d_1283_v57_ = nw210_
                        index164_ = default__.safeIndex(45, (d_1280_v56_).length(0))
                        (d_1280_v56_)[index164_] = d_1206_v1_
                        index165_ = default__.safeIndex(45, (d_1280_v56_).length(0))
                        rhs232_ = d_1279_v55_
                        rhs233_ = d_1283_v57_
                        rhs234_ = d_1205_v0_
                        rhs235_ = d_1206_v1_
                        lhs188_ = d_1280_v56_
                        lhs189_ = default__.safeIndex(45, (d_1280_v56_).length(0))
                        d_1279_v55_ = rhs232_
                        d_1283_v57_ = rhs233_
                        d_1205_v0_ = rhs234_
                        lhs188_[lhs189_] = rhs235_
                        d_1284_v58_: _dafny.Seq
                        d_1284_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnmuw"))
                        d_1284_v58_ = p0
                        (globalState).f7 = (d_1280_v56_)[default__.safeIndex(45, (d_1280_v56_).length(0))]
                        d_1285_v59_: _dafny.Seq
                        d_1285_v59_ = _dafny.SeqWithoutIsStrInference([d_1284_v58_, p0, p0, p0])
                        (globalState).f1 = (d_1278_v54_).fm3(len(d_1285_v59_), (d_1205_v0_) > (d_1205_v0_), False, globalState)
                        d_1286_v60_: _dafny.MultiSet
                        d_1286_v60_ = _dafny.MultiSet([d_1205_v0_])
                        d_1287_v61_: _dafny.Set
                        d_1287_v61_ = _dafny.Set({d_1206_v1_})
                        d_1288_v62_: D7
                        d_1288_v62_ = D7_DC13(d_1287_v61_, d_1205_v0_, d_1205_v0_)
                        d_1289_v63_: C4
                        nw211_ = C4()
                        nw211_.ctor__(d_1286_v60_, False, (d_1288_v62_).cf18, (d_1205_v0_) - (d_1205_v0_))
                        d_1289_v63_ = nw211_
                    elif True:
                        d_1290_v64_: _dafny.MultiSet
                        d_1290_v64_ = _dafny.MultiSet([default__.fm2(d_1205_v0_, d_1205_v0_, globalState)])
                        d_1205_v0_ = (0) - (default__.safeModuloInt(d_1205_v0_, default__.safeModuloInt(d_1205_v0_, ((d_1290_v64_)[d_1205_v0_] if (d_1205_v0_) in (d_1290_v64_) else -984))))
                        d_1291_v65_: D9
                        d_1291_v65_ = D9_DC19()
                        d_1291_v65_ = d_1291_v65_
                        d_1206_v1_ = d_1206_v1_
                        d_1292_v66_: _dafny.Seq
                        d_1292_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxlji"))
                        d_1292_v66_ = ((d_1292_v66_) + (p0)).set(default__.safeIndex(d_1205_v0_, len((d_1292_v66_) + (p0))), d_1207_v2_)
                        d_1293_v67_: _dafny.Array
                        nw212_ = _dafny.Array(False, 2)
                        d_1293_v67_ = nw212_
                        d_1294_v68_: _dafny.Set
                        d_1294_v68_ = _dafny.Set({True, d_1206_v1_, d_1206_v1_})
                        index166_ = default__.safeIndex(909, (d_1293_v67_).length(0))
                        (d_1293_v67_)[index166_] = (d_1294_v68_).isdisjoint(_dafny.Set({d_1206_v1_}))
                        d_1295_v69_: _dafny.Seq
                        d_1295_v69_ = _dafny.SeqWithoutIsStrInference([d_1205_v0_])
                        d_1296_v70_: _dafny.Set
                        d_1296_v70_ = _dafny.Set({(0) - ((0) - ((_dafny.MultiSet(d_1295_v69_)).cardinality)), d_1205_v0_, d_1205_v0_, d_1205_v0_})
                        d_1297_v71_: _dafny.Array
                        nw213_ = _dafny.Array(None, 5)
                        nw213_[int(0)] = (0) - (d_1205_v0_)
                        nw213_[int(1)] = d_1205_v0_
                        nw213_[int(2)] = (d_1205_v0_) * (d_1205_v0_)
                        nw213_[int(3)] = d_1205_v0_
                        nw213_[int(4)] = (d_1205_v0_) + (len(d_1296_v70_))
                        d_1297_v71_ = nw213_
                        index167_ = default__.safeIndex(571, (d_1297_v71_).length(0))
                        (d_1297_v71_)[index167_] = d_1205_v0_
                        d_1298_v72_: _dafny.Seq
                        d_1298_v72_ = _dafny.SeqWithoutIsStrInference([not(d_1206_v1_), d_1206_v1_])
                        d_1299_v73_: _dafny.Map
                        d_1299_v73_ = _dafny.Map({d_1205_v0_: d_1298_v72_})
                        d_1300_v74_: _dafny.Map
                        d_1300_v74_ = _dafny.Map({(0) - ((self).fm9(len(d_1299_v73_), (0) - ((0) - (len(d_1294_v68_))), d_1206_v1_, globalState)): (0) - (d_1205_v0_)})
                        d_1301_v75_: _dafny.Seq
                        d_1301_v75_ = _dafny.SeqWithoutIsStrInference([d_1300_v74_])
                        d_1302_v76_: _dafny.MultiSet
                        d_1302_v76_ = _dafny.MultiSet([default__.fm39(d_1205_v0_, d_1206_v1_, d_1205_v0_, globalState)])
                        index168_ = default__.safeIndex(909, (d_1293_v67_).length(0))
                        index169_ = default__.safeIndex(571, (d_1297_v71_).length(0))
                        rhs236_ = (d_1302_v76_).isdisjoint(_dafny.MultiSet((d_1292_v66_) + (default__.fm1(globalState))))
                        rhs237_ = not (d_1206_v1_) or (((d_1298_v72_).set(default__.safeIndex(d_1205_v0_, len(d_1298_v72_)), d_1206_v1_)) < (((d_1299_v73_)[d_1205_v0_] if (d_1205_v0_) in (d_1299_v73_) else d_1298_v72_)))
                        rhs238_ = (0) - (d_1205_v0_)
                        rhs239_ = (default__.safeModuloInt(d_1205_v0_, -527)) * ((338) - (len(p0)))
                        rhs240_ = (d_1301_v75_).set(default__.safeIndex(d_1205_v0_, len(d_1301_v75_)), (d_1300_v74_) | (d_1300_v74_))
                        lhs190_ = d_1293_v67_
                        lhs191_ = default__.safeIndex(909, (d_1293_v67_).length(0))
                        lhs192_ = d_1297_v71_
                        lhs193_ = default__.safeIndex(571, (d_1297_v71_).length(0))
                        lhs194_ = globalState
                        d_1206_v1_ = rhs236_
                        lhs190_[lhs191_] = rhs237_
                        lhs192_[lhs193_] = rhs238_
                        lhs194_.f1 = rhs239_
                        d_1301_v75_ = rhs240_
                    (globalState).f7 = not(d_1206_v1_)
                    if d_1206_v1_:
                        d_1303_v77_: _dafny.Seq
                        d_1303_v77_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_, d_1206_v1_])
                        d_1304_v78_: _dafny.Map
                        d_1304_v78_ = _dafny.Map({d_1205_v0_: d_1205_v0_})
                        d_1305_v79_: _dafny.Array
                        d_1305_v79_ = p1
                        d_1306_v80_: _dafny.Set
                        d_1306_v80_ = _dafny.Set({d_1207_v2_})
                        d_1307_v81_: _dafny.Map
                        d_1307_v81_ = _dafny.Map({d_1305_v79_: d_1306_v80_})
                        d_1308_v82_: _dafny.Array
                        nw214_ = _dafny.Array(None, 10)
                        nw214_[int(0)] = d_1205_v0_
                        nw214_[int(1)] = -279
                        nw214_[int(2)] = d_1205_v0_
                        nw214_[int(3)] = d_1205_v0_
                        nw214_[int(4)] = d_1205_v0_
                        nw214_[int(5)] = (0) - (d_1205_v0_)
                        nw214_[int(6)] = d_1205_v0_
                        nw214_[int(7)] = len((d_1303_v77_).set(default__.safeIndex(default__.fm2(d_1205_v0_, len(d_1304_v78_), globalState), len(d_1303_v77_)), not(d_1206_v1_)))
                        nw214_[int(8)] = d_1205_v0_
                        nw214_[int(9)] = len(((d_1307_v81_)[d_1305_v79_] if (d_1305_v79_) in (d_1307_v81_) else _dafny.Set({d_1207_v2_})))
                        d_1308_v82_ = nw214_
                        index170_ = default__.safeIndex(102, (p1).length(0))
                        (p1)[index170_] = d_1308_v82_
                        index171_ = default__.safeIndex(102, (p1).length(0))
                        (p1)[index171_] = d_1308_v82_
                        d_1309_v83_: _dafny.Array
                        nw215_ = _dafny.Array(False, 14)
                        d_1309_v83_ = nw215_
                        d_1310_v84_: _dafny.Map
                        d_1310_v84_ = _dafny.Map({d_1206_v1_: d_1309_v83_})
                        d_1311_v86_: D15
                        d_1311_v86_ = D15_DC33(d_1310_v84_)
                        d_1312_v87_: _dafny.Map
                        def iife124_():
                            coll62_ = _dafny.Map()
                            compr_62_: int
                            for compr_62_ in _dafny.IntegerRange(-81, 373):
                                d_1313_v85_: int = compr_62_
                                if ((-81) <= (d_1313_v85_)) and ((d_1313_v85_) < (373)):
                                    coll62_[default__.safeModuloInt(d_1313_v85_, len(_dafny.Map({d_1206_v1_: d_1205_v0_})))] = (_dafny.MultiSet([d_1205_v0_])).cardinality
                            return _dafny.Map(coll62_)
                        d_1312_v87_ = _dafny.Map({iife124_()
                        : (d_1311_v86_).cf49})
                        d_1314_v88_: D16
                        d_1314_v88_ = D16_DC38(d_1304_v78_)
                        d_1310_v84_ = ((d_1312_v87_)[(d_1314_v88_).cf61] if ((d_1314_v88_).cf61) in (d_1312_v87_) else d_1310_v84_)
                        d_1315_v89_: _dafny.Seq
                        d_1315_v89_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1205_v0_), -243])
                        (globalState).f7 = (len(d_1315_v89_)) < (d_1205_v0_)
                        arr0_ = (p1)[default__.safeIndex(102, (p1).length(0))]
                        index172_ = default__.safeIndex(724, ((p1)[default__.safeIndex(102, (p1).length(0))]).length(0))
                        arr0_[index172_] = 161
                        arr1_ = (p1)[default__.safeIndex(102, (p1).length(0))]
                        index173_ = default__.safeIndex(724, ((p1)[default__.safeIndex(102, (p1).length(0))]).length(0))
                        rhs241_ = (d_1205_v0_) * (d_1205_v0_)
                        rhs242_ = d_1207_v2_
                        lhs195_ = (p1)[default__.safeIndex(102, (p1).length(0))]
                        lhs196_ = default__.safeIndex(724, ((p1)[default__.safeIndex(102, (p1).length(0))]).length(0))
                        lhs195_[lhs196_] = rhs241_
                        d_1207_v2_ = rhs242_
                        d_1316_v90_: _dafny.Map
                        d_1316_v90_ = _dafny.Map({len((d_1208_v3_) | ((d_1277_v53_)[default__.safeIndex(440, (d_1277_v53_).length(0))])): d_1315_v89_})
                        d_1316_v90_ = d_1316_v90_
                    elif True:
                        (globalState).f1 = d_1205_v0_
                        d_1317_v91_: _dafny.Array
                        nw216_ = _dafny.Array(False, 7)
                        d_1317_v91_ = nw216_
                        index174_ = default__.safeIndex(477, (d_1317_v91_).length(0))
                        (d_1317_v91_)[index174_] = d_1206_v1_
                        d_1318_v92_: _dafny.Map
                        d_1318_v92_ = _dafny.Map({d_1205_v0_: not(False)})
                        d_1319_v93_: _dafny.Map
                        d_1319_v93_ = _dafny.Map({d_1206_v1_: d_1205_v0_})
                        d_1320_v94_: _dafny.Seq
                        d_1320_v94_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_])
                        d_1321_v95_: _dafny.Seq
                        d_1321_v95_ = _dafny.SeqWithoutIsStrInference([d_1205_v0_])
                        d_1322_v96_: _dafny.Array
                        nw217_ = _dafny.Array(None, 17)
                        nw217_[int(0)] = (0) - ((d_1205_v0_) + (755))
                        nw217_[int(1)] = default__.safeModuloInt(len(d_1318_v92_), d_1205_v0_)
                        nw217_[int(2)] = (0) - (d_1205_v0_)
                        nw217_[int(3)] = d_1205_v0_
                        nw217_[int(4)] = (len(p0)) * (d_1205_v0_)
                        nw217_[int(5)] = ((d_1319_v93_)[d_1206_v1_] if (d_1206_v1_) in (d_1319_v93_) else (0) - (d_1205_v0_))
                        nw217_[int(6)] = d_1205_v0_
                        nw217_[int(7)] = 784
                        nw217_[int(8)] = len(d_1320_v94_)
                        nw217_[int(9)] = d_1205_v0_
                        nw217_[int(10)] = d_1205_v0_
                        nw217_[int(11)] = d_1205_v0_
                        nw217_[int(12)] = (d_1205_v0_) + (len(d_1320_v94_))
                        nw217_[int(13)] = (d_1321_v95_)[default__.safeIndex(d_1205_v0_, len(d_1321_v95_))]
                        nw217_[int(14)] = len(d_1320_v94_)
                        nw217_[int(15)] = len(p0)
                        nw217_[int(16)] = (0) - (d_1205_v0_)
                        d_1322_v96_ = nw217_
                        d_1323_v97_: _dafny.Array
                        nw218_ = _dafny.Array(_dafny.Array(None, 0), 14)
                        d_1323_v97_ = nw218_
                        index175_ = default__.safeIndex(306, (d_1323_v97_).length(0))
                        (d_1323_v97_)[index175_] = d_1317_v91_
                        d_1324_v98_: _dafny.Seq
                        d_1324_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isyb"))
                        d_1325_v99_: _dafny.Set
                        d_1325_v99_ = _dafny.Set({d_1205_v0_, d_1205_v0_})
                        index176_ = default__.safeIndex(477, (d_1317_v91_).length(0))
                        index177_ = default__.safeIndex(306, (d_1323_v97_).length(0))
                        rhs243_ = False
                        rhs244_ = (d_1322_v96_ if not(d_1206_v1_) else d_1322_v96_)
                        rhs245_ = d_1317_v91_
                        rhs246_ = (default__.fm1(globalState)).set(default__.safeIndex(len((d_1325_v99_) - (d_1325_v99_)), len(default__.fm1(globalState))), d_1207_v2_)
                        lhs197_ = d_1317_v91_
                        lhs198_ = default__.safeIndex(477, (d_1317_v91_).length(0))
                        lhs199_ = d_1323_v97_
                        lhs200_ = default__.safeIndex(306, (d_1323_v97_).length(0))
                        lhs197_[lhs198_] = rhs243_
                        d_1322_v96_ = rhs244_
                        lhs199_[lhs200_] = rhs245_
                        d_1324_v98_ = rhs246_
                        d_1326_v100_: _dafny.Set
                        d_1326_v100_ = _dafny.Set({(d_1323_v97_)[default__.safeIndex(306, (d_1323_v97_).length(0))], (d_1323_v97_)[default__.safeIndex(306, (d_1323_v97_).length(0))]})
                        d_1327_v101_: _dafny.Map
                        d_1327_v101_ = _dafny.Map({(d_1326_v100_) | (d_1326_v100_): d_1205_v0_})
                        d_1327_v101_ = (d_1327_v101_).set(d_1326_v100_, len(d_1319_v93_))
                        d_1322_v96_ = d_1322_v96_
                        d_1328_v102_: C0
                        nw219_ = C0()
                        nw219_.ctor__(d_1320_v94_, (0) - ((d_1205_v0_ if not(d_1206_v1_) else d_1205_v0_)))
                        d_1328_v102_ = nw219_
                    pass
            pass
        d_1329_v103_: _dafny.Map
        d_1329_v103_ = _dafny.Map({False: d_1205_v0_})
        d_1330_v104_: _dafny.Set
        d_1330_v104_ = _dafny.Set({(0) - (len(d_1329_v103_)), d_1205_v0_, d_1205_v0_})
        d_1331_v105_: _dafny.MultiSet
        d_1331_v105_ = _dafny.MultiSet([d_1205_v0_])
        (globalState).f1 = ((d_1205_v0_ if d_1206_v1_ else len(d_1330_v104_))) + (((d_1331_v105_)[d_1205_v0_] if (d_1205_v0_) in (d_1331_v105_) else d_1205_v0_))
        d_1332_v106_: _dafny.Seq
        d_1332_v106_ = _dafny.SeqWithoutIsStrInference([d_1206_v1_, d_1206_v1_])
        d_1333_v107_: C2
        nw220_ = C2()
        nw220_.ctor__(d_1207_v2_, d_1205_v0_)
        d_1333_v107_ = nw220_
        d_1334_v108_: _dafny.MultiSet
        d_1334_v108_ = _dafny.MultiSet([d_1333_v107_])
        d_1335_v109_: _dafny.Array
        nw221_ = _dafny.Array(None, 14)
        nw221_[int(0)] = d_1205_v0_
        nw221_[int(1)] = len(d_1329_v103_)
        nw221_[int(2)] = len(d_1332_v106_)
        nw221_[int(3)] = len(d_1329_v103_)
        nw221_[int(4)] = 83
        nw221_[int(5)] = 347
        nw221_[int(6)] = d_1205_v0_
        nw221_[int(7)] = d_1205_v0_
        nw221_[int(8)] = d_1205_v0_
        nw221_[int(9)] = d_1205_v0_
        nw221_[int(10)] = 956
        nw221_[int(11)] = d_1205_v0_
        nw221_[int(12)] = (0) - (d_1205_v0_)
        nw221_[int(13)] = (d_1334_v108_).cardinality
        d_1335_v109_ = nw221_
        d_1336_v110_: _dafny.Array
        nw222_ = _dafny.Array(_dafny.Array(None, 0), 9)
        d_1336_v110_ = nw222_
        d_1337_v111_: _dafny.Map
        d_1337_v111_ = _dafny.Map({d_1335_v109_: d_1336_v110_})
        (globalState).f7 = (not(d_1206_v1_)) and ((d_1335_v109_) in (d_1337_v111_))


class C6(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Set = _dafny.Set({})
        self._f11: int = int(0)
        self._f15: D2 = D2.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f12(self):
        return self._f12
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f15, f12, f11):
        (self)._f15 = f15
        (self)._f12 = f12
        (self)._f11 = f11

    def fm5(self, p0, p1, p2, p3, globalState):
        if False:
            return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoywkyi")))) <= ((self).f11)
        elif True:
            return not(not(True))

    def fm6(self, p0, p1, globalState):
        return default__.safeModuloInt((self).f11, len((_dafny.Map({False: True}) if not(not(True)) else _dafny.Map({False: False}))))

    def fm3(self, p0, p1, p2, globalState):
        def iife125_():
            coll63_ = _dafny.Map()
            compr_63_: str
            for compr_63_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('y')])).Elements:
                d_1338_v0_: str = compr_63_
                if (d_1338_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('y')])):
                    coll63_[d_1338_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhdueatrb"))
            return _dafny.Map(coll63_)
        return (len(iife125_()
        )) - ((0) - (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1339_i0_ in range(default__.abs(-288))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1340_i1_ in range(default__.abs(941))])))))

    def fm4(self, p0, p1, p2, p3, globalState):
        source20_ = (self).f15
        if source20_.is_DC4:
            d_1341___mcc_h0_ = source20_.cf6
            d_1342___mcc_h1_ = source20_.cf7
            d_1343___mcc_h2_ = source20_.cf8
            d_1344_cf8_ = d_1343___mcc_h2_
            d_1345_cf7_ = d_1342___mcc_h1_
            d_1346_cf6_ = d_1341___mcc_h0_
            return True
        elif source20_.is_DC5:
            return (D0_DC1(True, 252, True)).cf1
        elif True:
            d_1347___mcc_h3_ = source20_.cf5
            d_1348_cf5_ = d_1347___mcc_h3_
            return ((self).f12).isdisjoint(_dafny.Set({(D0_DC1(False, (self).f11, True)).cf3}))

    def m0(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_1349_v0_: _dafny.Array
        nw223_ = _dafny.Array(_dafny.Set({}), 8)
        d_1349_v0_ = nw223_
        d_1350_v1_: _dafny.Array
        nw224_ = _dafny.Array(False, 13)
        d_1350_v1_ = nw224_
        d_1351_v2_: _dafny.Set
        d_1351_v2_ = _dafny.Set({d_1350_v1_})
        index178_ = default__.safeIndex(911, (d_1349_v0_).length(0))
        (d_1349_v0_)[index178_] = d_1351_v2_
        index179_ = default__.safeIndex(911, (d_1349_v0_).length(0))
        (d_1349_v0_)[index179_] = d_1351_v2_
        d_1352_v3_: bool
        d_1352_v3_ = True
        d_1353_i0_: int
        d_1353_i0_ = 0
        with _dafny.label("15"):
            while d_1352_v3_:
                with _dafny.c_label("15"):
                    if (d_1353_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_1353_i0_ = (d_1353_i0_) + (1)
                    d_1354_v4_: _dafny.Map
                    d_1354_v4_ = _dafny.Map({d_1352_v3_: d_1352_v3_})
                    d_1354_v4_ = d_1354_v4_
                    d_1355_v5_: _dafny.MultiSet
                    d_1355_v5_ = _dafny.MultiSet([(self).f11])
                    d_1356_v6_: _dafny.Map
                    d_1356_v6_ = _dafny.Map({d_1355_v5_: not (d_1352_v3_) or (d_1352_v3_)})
                    (globalState).f1 = len(d_1356_v6_)
                    d_1352_v3_ = default__.fm0(globalState)
                    (globalState).f6 = d_1352_v3_
                    pass
            pass
        (globalState).f1 = (self).f11
        if (d_1352_v3_ if d_1352_v3_ else d_1352_v3_):
            d_1357_v7_: _dafny.Map
            d_1357_v7_ = _dafny.Map({(self).f11: 194})
            d_1358_v8_: _dafny.Array
            nw225_ = _dafny.Array(None, 6)
            nw225_[int(0)] = -194
            nw225_[int(1)] = (self).f11
            nw225_[int(2)] = (self).fm3((self).f11, d_1352_v3_, d_1352_v3_, globalState)
            nw225_[int(3)] = (self).f11
            nw225_[int(4)] = (self).f11
            nw225_[int(5)] = default__.safeModuloInt(((d_1357_v7_)[(self).f11] if ((self).f11) in (d_1357_v7_) else (self).f11), (self).f11)
            d_1358_v8_ = nw225_
            index180_ = default__.safeIndex(89, (d_1358_v8_).length(0))
            (d_1358_v8_)[index180_] = (0) - (default__.safeModuloInt((self).f11, (self).f11))
            d_1359_v9_: _dafny.Seq
            d_1359_v9_ = _dafny.SeqWithoutIsStrInference([not(d_1352_v3_), d_1352_v3_])
            d_1360_v10_: _dafny.Seq
            d_1360_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f11: d_1352_v3_})])
            d_1361_v11_: _dafny.Seq
            d_1361_v11_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11, (self).f11])
            d_1362_v12_: _dafny.Map
            d_1362_v12_ = _dafny.Map({len((self).f12): False})
            index181_ = default__.safeIndex(89, (d_1358_v8_).length(0))
            rhs247_ = (self).f11
            rhs248_ = (d_1359_v9_)[default__.safeIndex(len(((d_1360_v10_)[default__.safeIndex((self).fm3(len(d_1361_v11_), d_1352_v3_, d_1352_v3_, globalState), len(d_1360_v10_))]) | (d_1362_v12_)), len(d_1359_v9_))]
            rhs249_ = d_1352_v3_
            lhs201_ = d_1358_v8_
            lhs202_ = default__.safeIndex(89, (d_1358_v8_).length(0))
            lhs203_ = globalState
            lhs204_ = globalState
            lhs201_[lhs202_] = rhs247_
            lhs203_.f7 = rhs248_
            lhs204_.f6 = rhs249_
            d_1363_v13_: _dafny.Array
            nw226_ = _dafny.Array(_dafny.Seq({}), 17)
            d_1363_v13_ = nw226_
            index182_ = default__.safeIndex(862, (d_1363_v13_).length(0))
            (d_1363_v13_)[index182_] = d_1359_v9_
            index183_ = default__.safeIndex(862, (d_1363_v13_).length(0))
            (d_1363_v13_)[index183_] = d_1359_v9_
            d_1364_v14_: _dafny.Map
            d_1364_v14_ = _dafny.Map({False: (self).f11})
            (globalState).f1 = (default__.fm2(384, (d_1358_v8_)[default__.safeIndex(89, (d_1358_v8_).length(0))], globalState)) + ((0) - (((d_1357_v7_)[(d_1358_v8_)[default__.safeIndex(89, (d_1358_v8_).length(0))]] if ((d_1358_v8_)[default__.safeIndex(89, (d_1358_v8_).length(0))]) in (d_1357_v7_) else len(d_1364_v14_))))
            d_1365_v15_: _dafny.Seq
            d_1365_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irx"))
            d_1365_v15_ = (d_1365_v15_).set(default__.safeIndex(382, len(d_1365_v15_)), _dafny.CodePoint('m'))
            d_1366_v16_: C2
            nw227_ = C2()
            nw227_.ctor__(p0, ((self).f11) - ((self).f11))
            d_1366_v16_ = nw227_
        elif True:
            d_1367_v17_: _dafny.Map
            d_1367_v17_ = _dafny.Map({d_1352_v3_: d_1352_v3_})
            d_1368_v18_: D8
            d_1368_v18_ = D8_DC15(d_1367_v17_)
            d_1369_v19_: D8
            d_1369_v19_ = D8_DC17(D8_DC17(d_1368_v18_))
            d_1370_v20_: _dafny.Map
            d_1370_v20_ = _dafny.Map({d_1369_v19_: (self).f11})
            d_1371_v21_: D13
            d_1371_v21_ = D13_DC27(d_1370_v20_)
            d_1372_v22_: D13
            d_1372_v22_ = D13_DC29(d_1371_v21_)
            source21_ = d_1372_v22_
            if source21_.is_DC28:
                d_1373___mcc_h0_ = source21_.cf40
                d_1374___mcc_h1_ = source21_.cf41
                d_1375___mcc_h2_ = source21_.cf42
                d_1376_cf42_ = d_1375___mcc_h2_
                d_1377_cf41_ = d_1374___mcc_h1_
                d_1378_cf40_ = d_1373___mcc_h0_
                d_1379_v23_: _dafny.MultiSet
                d_1379_v23_ = _dafny.MultiSet([d_1376_cf42_])
                d_1380_v24_: _dafny.Map
                d_1380_v24_ = _dafny.Map({(d_1379_v23_).cardinality: d_1376_cf42_})
                d_1381_v25_: _dafny.Map
                d_1381_v25_ = _dafny.Map({d_1352_v3_: D12_DC26(d_1380_v24_)})
                d_1382_v26_: _dafny.MultiSet
                d_1382_v26_ = _dafny.MultiSet([d_1381_v25_, d_1381_v25_])
                d_1383_v27_: _dafny.Seq
                d_1383_v27_ = _dafny.SeqWithoutIsStrInference([d_1377_cf41_, 262])
                d_1384_v28_: _dafny.Map
                d_1384_v28_ = _dafny.Map({not(d_1376_cf42_): d_1377_cf41_})
                d_1385_v29_: _dafny.Array
                nw228_ = _dafny.Array(None, 20)
                nw228_[int(0)] = (self).f11
                nw228_[int(1)] = default__.safeDivisionInt(d_1377_cf41_, 328)
                nw228_[int(2)] = (self).f11
                nw228_[int(3)] = (self).f11
                nw228_[int(4)] = d_1377_cf41_
                nw228_[int(5)] = 6
                nw228_[int(6)] = d_1377_cf41_
                nw228_[int(7)] = (self).f11
                nw228_[int(8)] = ((0) - (((d_1382_v26_)[d_1381_v25_] if (d_1381_v25_) in (d_1382_v26_) else (self).f11))) + ((self).f11)
                nw228_[int(9)] = (self).f11
                nw228_[int(10)] = len((d_1383_v27_) + (_dafny.SeqWithoutIsStrInference([(self).f11 for d_1386_i1_ in range(default__.abs(371))])))
                nw228_[int(11)] = default__.safeDivisionInt((self).f11, ((d_1384_v28_)[d_1352_v3_] if (d_1352_v3_) in (d_1384_v28_) else (self).f11))
                nw228_[int(12)] = d_1377_cf41_
                nw228_[int(13)] = d_1377_cf41_
                nw228_[int(14)] = ((d_1379_v23_).cardinality) - ((self).fm6(True, (self).f11, globalState))
                nw228_[int(15)] = d_1377_cf41_
                nw228_[int(16)] = (self).f11
                nw228_[int(17)] = d_1377_cf41_
                nw228_[int(18)] = d_1377_cf41_
                nw228_[int(19)] = 798
                d_1385_v29_ = nw228_
                index184_ = default__.safeIndex(221, (d_1385_v29_).length(0))
                (d_1385_v29_)[index184_] = (self).f11
                d_1387_v30_: _dafny.Map
                d_1387_v30_ = _dafny.Map({d_1378_cf40_: (d_1379_v23_).cardinality})
                d_1388_v31_: _dafny.Map
                d_1388_v31_ = _dafny.Map({not(d_1376_cf42_): d_1378_cf40_})
                index185_ = default__.safeIndex(221, (d_1385_v29_).length(0))
                (d_1385_v29_)[index185_] = ((d_1387_v30_)[((d_1388_v31_)[d_1352_v3_] if (d_1352_v3_) in (d_1388_v31_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1389_i2_ in range(default__.abs(-548))]))] if (((d_1388_v31_)[d_1352_v3_] if (d_1352_v3_) in (d_1388_v31_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1390_i2_ in range(default__.abs(-548))]))) in (d_1387_v30_) else (self).fm3((self).f11, d_1376_cf42_, d_1352_v3_, globalState))
                d_1391_v32_: _dafny.Map
                d_1391_v32_ = _dafny.Map({d_1385_v29_: d_1352_v3_})
                d_1392_v33_: D10
                d_1392_v33_ = D10_DC20(d_1391_v32_)
                d_1393_v34_: _dafny.MultiSet
                d_1393_v34_ = _dafny.MultiSet([d_1392_v33_, d_1392_v33_])
                d_1393_v34_ = d_1393_v34_
                d_1394_v35_: C1
                nw229_ = C1()
                nw229_.ctor__(False, d_1376_cf42_)
                d_1394_v35_ = nw229_
                d_1395_v36_: _dafny.Seq
                d_1395_v36_ = _dafny.SeqWithoutIsStrInference([d_1394_v35_])
                index186_ = default__.safeIndex(221, (d_1385_v29_).length(0))
                rhs250_ = (len(d_1380_v24_)) * (d_1377_cf41_)
                rhs251_ = (0) - (((0) - ((0) - ((d_1377_cf41_) * ((d_1385_v29_)[default__.safeIndex(221, (d_1385_v29_).length(0))])))) - (d_1377_cf41_))
                rhs252_ = (d_1394_v35_).f19
                rhs253_ = d_1395_v36_
                lhs205_ = d_1385_v29_
                lhs206_ = default__.safeIndex(221, (d_1385_v29_).length(0))
                lhs207_ = globalState
                lhs208_ = globalState
                lhs205_[lhs206_] = rhs250_
                lhs207_.f1 = rhs251_
                lhs208_.f6 = rhs252_
                d_1395_v36_ = rhs253_
                (globalState).f1 = (self).f11
            elif source21_.is_DC27:
                d_1396___mcc_h3_ = source21_.cf39
                d_1397_cf39_ = d_1396___mcc_h3_
                d_1398_v37_: _dafny.Seq
                d_1398_v37_ = _dafny.SeqWithoutIsStrInference([d_1352_v3_, d_1352_v3_, d_1352_v3_])
                d_1399_v38_: _dafny.Map
                d_1399_v38_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f11])) + (_dafny.SeqWithoutIsStrInference([(self).f11, (0) - ((self).f11), len(d_1398_v37_)])): (self).f11})
                d_1399_v38_ = d_1399_v38_
                d_1400_v39_: _dafny.Seq
                d_1400_v39_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_1401_v40_: _dafny.Seq
                d_1401_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdlkv"))
                d_1402_v41_: _dafny.Map
                d_1402_v41_ = _dafny.Map({default__.fm2(78, (self).f11, globalState): (self).f11})
                d_1403_v42_: _dafny.MultiSet
                d_1403_v42_ = _dafny.MultiSet([(self).f11])
                d_1404_v43_: _dafny.Map
                d_1404_v43_ = _dafny.Map({d_1403_v42_: d_1402_v41_})
                d_1405_v44_: _dafny.Array
                nw230_ = _dafny.Array(None, 16)
                nw230_[int(0)] = (self).f11
                nw230_[int(1)] = ((self).f11) * ((self).f11)
                nw230_[int(2)] = (self).f11
                nw230_[int(3)] = -574
                nw230_[int(4)] = (self).f11
                nw230_[int(5)] = (self).f11
                nw230_[int(6)] = (0) - ((len(d_1400_v39_)) + ((self).f11))
                nw230_[int(7)] = (self).f11
                nw230_[int(8)] = (len(d_1401_v40_)) * ((self).f11)
                nw230_[int(9)] = len(d_1367_v17_)
                nw230_[int(10)] = (self).f11
                nw230_[int(11)] = default__.safeModuloInt(len(d_1402_v41_), (self).f11)
                nw230_[int(12)] = (self).f11
                nw230_[int(13)] = (self).f11
                nw230_[int(14)] = (0) - ((d_1400_v39_)[default__.safeIndex((self).f11, len(d_1400_v39_))])
                nw230_[int(15)] = ((self).fm3((0) - ((0) - ((self).f11)), (self).fm4((self).f11, d_1404_v43_, (self).f11, d_1352_v3_, globalState), d_1352_v3_, globalState)) - (len(_dafny.Map({(self).f11: d_1352_v3_})))
                d_1405_v44_ = nw230_
                index187_ = default__.safeIndex(666, (d_1405_v44_).length(0))
                (d_1405_v44_)[index187_] = ((d_1403_v42_) | (d_1403_v42_)).cardinality
                index188_ = default__.safeIndex(17, (d_1405_v44_).length(0))
                (d_1405_v44_)[index188_] = (self).f11
                d_1406_v45_: _dafny.Map
                d_1406_v45_ = _dafny.Map({d_1350_v1_: (self).f11})
                index189_ = default__.safeIndex(666, (d_1405_v44_).length(0))
                index190_ = default__.safeIndex(17, (d_1405_v44_).length(0))
                rhs254_ = default__.safeDivisionInt(len(d_1398_v37_), (self).f11)
                rhs255_ = (((d_1406_v45_)[d_1350_v1_] if (d_1350_v1_) in (d_1406_v45_) else (self).f11)) - ((self).f11)
                lhs209_ = d_1405_v44_
                lhs210_ = default__.safeIndex(666, (d_1405_v44_).length(0))
                lhs211_ = d_1405_v44_
                lhs212_ = default__.safeIndex(17, (d_1405_v44_).length(0))
                lhs209_[lhs210_] = rhs254_
                lhs211_[lhs212_] = rhs255_
                d_1407_v46_: D10
                d_1407_v46_ = D10_DC21(((d_1367_v17_)[d_1352_v3_] if (d_1352_v3_) in (d_1367_v17_) else d_1352_v3_), d_1352_v3_)
                d_1408_v47_: _dafny.Map
                d_1408_v47_ = _dafny.Map({d_1407_v46_: _dafny.Map({(d_1405_v44_)[default__.safeIndex(666, (d_1405_v44_).length(0))]: (self).f11})})
                d_1408_v47_ = (d_1408_v47_).set(d_1407_v46_, d_1402_v41_)
                d_1409_v48_: _dafny.Array
                nw231_ = _dafny.Array(D12.default()(), 25)
                d_1409_v48_ = nw231_
                d_1410_v49_: _dafny.Array
                def lambda55_(d_1411_v3_):
                    def lambda56_(d_1412_i3_):
                        return D8_DC16(d_1411_v3_, -553)

                    return lambda56_

                init29_ = lambda55_(d_1352_v3_)
                nw232_ = _dafny.Array(None, 10)
                for i0_29_ in range(nw232_.length(0)):
                    nw232_[i0_29_] = init29_(i0_29_)
                d_1410_v49_ = nw232_
                d_1413_v50_: D12
                d_1413_v50_ = D12_DC25(d_1410_v49_)
                index191_ = default__.safeIndex(864, (d_1409_v48_).length(0))
                (d_1409_v48_)[index191_] = d_1413_v50_
                index192_ = default__.safeIndex(864, (d_1409_v48_).length(0))
                (d_1409_v48_)[index192_] = D12_DC25(d_1410_v49_)
            elif True:
                d_1414___mcc_h4_ = source21_.cf43
                d_1415_cf43_ = d_1414___mcc_h4_
                d_1416_v51_: _dafny.Seq
                d_1416_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_1416_v51_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djqq"))) + (d_1416_v51_)
                d_1417_v52_: str
                d_1417_v52_ = _dafny.CodePoint('k')
                d_1417_v52_ = p0
                d_1418_v53_: D15
                d_1418_v53_ = D15_DC36((self).f11, d_1352_v3_, d_1352_v3_, d_1352_v3_, (0) - ((self).f11))
                (globalState).f1 = (self).fm6((d_1418_v53_).cf58, (self).f11, globalState)
                d_1417_v52_ = p0
            d_1419_v54_: D15
            d_1419_v54_ = D15_DC34()
            d_1420_v55_: _dafny.Set
            d_1420_v55_ = _dafny.Set({d_1419_v54_, d_1419_v54_, d_1419_v54_})
            source22_ = default__.fm47(d_1420_v55_, d_1352_v3_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(d_1352_v3_)]) for d_1421_i4_ in range(default__.abs(407))]), d_1419_v54_, globalState)
            if source22_.is_DC28:
                d_1422___mcc_h5_ = source22_.cf40
                d_1423___mcc_h6_ = source22_.cf41
                d_1424___mcc_h7_ = source22_.cf42
                d_1425_cf42_ = d_1424___mcc_h7_
                d_1426_cf41_ = d_1423___mcc_h6_
                d_1427_cf40_ = d_1422___mcc_h5_
                d_1428_v56_: D8
                d_1428_v56_ = D8_DC15(d_1367_v17_)
                d_1429_v57_: _dafny.MultiSet
                d_1429_v57_ = _dafny.MultiSet([d_1352_v3_, d_1425_cf42_])
                d_1430_v58_: _dafny.Map
                d_1430_v58_ = _dafny.Map({d_1427_cf40_: ((d_1429_v57_)[d_1352_v3_] if (d_1352_v3_) in (d_1429_v57_) else d_1426_cf41_)})
                d_1431_v60_: _dafny.Seq
                d_1431_v60_ = _dafny.SeqWithoutIsStrInference([d_1427_cf40_, d_1427_cf40_, d_1427_cf40_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhsif")), d_1427_cf40_])
                d_1432_v61_: T1
                nw233_ = C4()
                nw233_.ctor__(_dafny.MultiSet([(self).f11]), d_1352_v3_, _dafny.Set({d_1352_v3_}), (self).f11)
                d_1432_v61_ = nw233_
                d_1433_v62_: _dafny.Array
                nw234_ = _dafny.Array(None, 17)
                nw234_[int(0)] = d_1432_v61_
                nw234_[int(1)] = d_1432_v61_
                nw234_[int(2)] = d_1432_v61_
                nw234_[int(3)] = d_1432_v61_
                nw234_[int(4)] = d_1432_v61_
                nw234_[int(5)] = d_1432_v61_
                nw234_[int(6)] = d_1432_v61_
                nw234_[int(7)] = d_1432_v61_
                nw234_[int(8)] = d_1432_v61_
                nw234_[int(9)] = d_1432_v61_
                nw234_[int(10)] = d_1432_v61_
                nw234_[int(11)] = d_1432_v61_
                nw234_[int(12)] = d_1432_v61_
                nw234_[int(13)] = d_1432_v61_
                nw234_[int(14)] = d_1432_v61_
                nw234_[int(15)] = d_1432_v61_
                nw234_[int(16)] = d_1432_v61_
                d_1433_v62_ = nw234_
                d_1434_v63_: _dafny.MultiSet
                d_1434_v63_ = _dafny.MultiSet([d_1433_v62_, d_1433_v62_, d_1433_v62_])
                d_1435_v64_: _dafny.Seq
                d_1435_v64_ = _dafny.SeqWithoutIsStrInference([(self).f11, d_1426_cf41_])
                d_1436_v65_: _dafny.Map
                d_1436_v65_ = _dafny.Map({(self).f11: d_1429_v57_})
                d_1437_v66_: _dafny.Array
                nw235_ = _dafny.Array(None, 29)
                nw235_[int(0)] = d_1426_cf41_
                nw235_[int(1)] = (self).fm3((self).f11, d_1352_v3_, d_1352_v3_, globalState)
                nw235_[int(2)] = default__.safeModuloInt((0) - (d_1426_cf41_), (self).f11)
                nw235_[int(3)] = (self).f11
                nw235_[int(4)] = (d_1426_cf41_) * ((self).f11)
                nw235_[int(5)] = (self).fm6(d_1425_cf42_, d_1426_cf41_, globalState)
                nw235_[int(6)] = d_1426_cf41_
                nw235_[int(7)] = d_1426_cf41_
                def iife126_():
                    coll64_ = _dafny.Map()
                    compr_64_: _dafny.Seq
                    for compr_64_ in (d_1431_v60_).Elements:
                        d_1438_v59_: _dafny.Seq = compr_64_
                        if (d_1438_v59_) in (d_1431_v60_):
                            coll64_[d_1438_v59_] = 510
                    return _dafny.Map(coll64_)
                nw235_[int(8)] = len((d_1430_v58_) | (iife126_()
                ))
                nw235_[int(9)] = (self).f11
                nw235_[int(10)] = d_1426_cf41_
                nw235_[int(11)] = (self).f11
                nw235_[int(12)] = (self).f11
                nw235_[int(13)] = (self).f11
                nw235_[int(14)] = (0) - (((d_1434_v63_).cardinality) - ((self).f11))
                nw235_[int(15)] = default__.safeDivisionInt(-181, (d_1432_v61_).fm3(len(d_1427_cf40_), d_1425_cf42_, d_1425_cf42_, globalState))
                nw235_[int(16)] = (d_1432_v61_).f11
                nw235_[int(17)] = 438
                nw235_[int(18)] = (self).fm3((0) - ((d_1432_v61_).f11), d_1352_v3_, not(d_1352_v3_), globalState)
                nw235_[int(19)] = (d_1426_cf41_) * ((self).f11)
                nw235_[int(20)] = len((_dafny.Map({False: _dafny.Map({d_1352_v3_: d_1425_cf42_})})).set(d_1352_v3_, d_1367_v17_))
                nw235_[int(21)] = len(d_1367_v17_)
                nw235_[int(22)] = len(d_1435_v64_)
                nw235_[int(23)] = d_1426_cf41_
                nw235_[int(24)] = (self).f11
                nw235_[int(25)] = len((self).f12)
                nw235_[int(26)] = (d_1426_cf41_) + ((self).f11)
                nw235_[int(27)] = len(d_1436_v65_)
                nw235_[int(28)] = d_1426_cf41_
                d_1437_v66_ = nw235_
                d_1439_v67_: _dafny.MultiSet
                d_1439_v67_ = _dafny.MultiSet([(d_1426_cf41_) + ((self).f11)])
                d_1440_v68_: _dafny.Map
                d_1440_v68_ = _dafny.Map({d_1426_cf41_: d_1426_cf41_})
                d_1441_v69_: _dafny.Seq
                d_1441_v69_ = _dafny.SeqWithoutIsStrInference([not(d_1352_v3_), d_1352_v3_])
                rhs256_ = ((d_1439_v67_)[(0) - (((d_1440_v68_)[-405] if (-405) in (d_1440_v68_) else (self).f11))] if ((0) - (((d_1440_v68_)[-405] if (-405) in (d_1440_v68_) else (self).f11))) in (d_1439_v67_) else (0) - (default__.safeModuloInt(762, len(d_1441_v69_))))
                rhs257_ = len((d_1427_cf40_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1442_i5_ in range(default__.abs(443))])))
                rhs258_ = default__.fm48(458, globalState)
                rhs259_ = (d_1425_cf42_) == (d_1425_cf42_)
                rhs260_ = d_1437_v66_
                lhs213_ = globalState
                lhs214_ = globalState
                lhs215_ = globalState
                lhs213_.f1 = rhs256_
                lhs214_.f1 = rhs257_
                d_1428_v56_ = rhs258_
                lhs215_.f7 = rhs259_
                d_1437_v66_ = rhs260_
                d_1352_v3_ = True
                d_1443_v70_: _dafny.Array
                nw236_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_1443_v70_ = nw236_
                index193_ = default__.safeIndex(980, (d_1443_v70_).length(0))
                (d_1443_v70_)[index193_] = p0
                index194_ = default__.safeIndex(980, (d_1443_v70_).length(0))
                (d_1443_v70_)[index194_] = p0
                d_1444_v71_: D6
                d_1444_v71_ = D6_DC10(d_1441_v69_)
                rhs261_ = d_1425_cf42_
                rhs262_ = d_1444_v71_
                lhs216_ = globalState
                lhs216_.f6 = rhs261_
                d_1444_v71_ = rhs262_
            elif source22_.is_DC27:
                d_1445___mcc_h8_ = source22_.cf39
                d_1446_cf39_ = d_1445___mcc_h8_
                index195_ = default__.safeIndex(744, (d_1350_v1_).length(0))
                (d_1350_v1_)[index195_] = (d_1352_v3_) == (d_1352_v3_)
                d_1447_v72_: _dafny.Seq
                d_1447_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctfp"))
                d_1448_v73_: D8
                d_1448_v73_ = D8_DC16(d_1352_v3_, len(d_1447_v72_))
                index196_ = default__.safeIndex(744, (d_1350_v1_).length(0))
                (d_1350_v1_)[index196_] = (d_1448_v73_).cf23
                d_1449_v74_: _dafny.Seq
                d_1449_v74_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_1449_v74_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11, (self).f11])
                d_1450_v75_: _dafny.Map
                d_1450_v75_ = _dafny.Map({d_1352_v3_: 879})
                d_1451_v76_: _dafny.Seq
                d_1451_v76_ = _dafny.SeqWithoutIsStrInference([d_1450_v75_])
                d_1452_v77_: _dafny.Map
                d_1452_v77_ = _dafny.Map({default__.fm2((0) - (len(d_1451_v76_)), (self).f11, globalState): d_1449_v74_})
                d_1452_v77_ = (d_1452_v77_).set(((0) - (len(default__.fm1(globalState))) if (d_1350_v1_)[default__.safeIndex(744, (d_1350_v1_).length(0))] else 649), d_1449_v74_)
                d_1453_v78_: _dafny.Array
                nw237_ = _dafny.Array(_dafny.Map({}), 29)
                d_1453_v78_ = nw237_
                d_1454_v79_: _dafny.Map
                d_1454_v79_ = _dafny.Map({(self).f11: d_1453_v78_})
                d_1454_v79_ = (d_1454_v79_).set((self).f11, d_1453_v78_)
            elif True:
                d_1455___mcc_h9_ = source22_.cf43
                d_1456_cf43_ = d_1455___mcc_h9_
                d_1457_v80_: _dafny.Seq
                d_1457_v80_ = _dafny.SeqWithoutIsStrInference([d_1352_v3_])
                d_1458_v81_: _dafny.Set
                d_1458_v81_ = _dafny.Set({p0, p0, p0, default__.fm20(d_1457_v80_, (self).f11, d_1352_v3_, globalState)})
                d_1458_v81_ = d_1458_v81_
                d_1459_v82_: _dafny.Seq
                d_1459_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                d_1460_v83_: _dafny.Seq
                d_1460_v83_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_1461_v84_: _dafny.Array
                nw238_ = _dafny.Array(None, 3)
                nw238_[int(0)] = (self).f11
                nw238_[int(1)] = (self).f11
                nw238_[int(2)] = (self).f11
                d_1461_v84_ = nw238_
                d_1462_v85_: _dafny.MultiSet
                d_1462_v85_ = _dafny.MultiSet([d_1352_v3_, d_1352_v3_])
                d_1463_v86_: _dafny.Map
                d_1463_v86_ = _dafny.Map({len(_dafny.Set({d_1461_v84_, d_1461_v84_})): (d_1462_v85_).cardinality})
                d_1464_v87_: _dafny.Array
                nw239_ = _dafny.Array(None, 15)
                nw239_[int(0)] = ((self).f11) * (len(d_1459_v82_))
                nw239_[int(1)] = (self).f11
                nw239_[int(2)] = (d_1460_v83_)[default__.safeIndex(((d_1463_v86_)[(D11_DC23(default__.fm0(globalState), 19, (self).f11)).cf33] if ((D11_DC23(default__.fm0(globalState), 19, (self).f11)).cf33) in (d_1463_v86_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwt")))), len(d_1460_v83_))]
                nw239_[int(3)] = (self).f11
                nw239_[int(4)] = default__.safeModuloInt((self).f11, (self).f11)
                nw239_[int(5)] = (self).f11
                nw239_[int(6)] = (self).f11
                nw239_[int(7)] = (self).fm3((self).f11, not(d_1352_v3_), d_1352_v3_, globalState)
                nw239_[int(8)] = (self).f11
                nw239_[int(9)] = (self).f11
                nw239_[int(10)] = (self).f11
                nw239_[int(11)] = ((self).f11) + (len(d_1457_v80_))
                nw239_[int(12)] = (self).f11
                nw239_[int(13)] = (self).f11
                nw239_[int(14)] = (self).f11
                d_1464_v87_ = nw239_
                index197_ = default__.safeIndex(585, (d_1461_v84_).length(0))
                (d_1461_v84_)[index197_] = 835
                d_1465_v88_: D14
                d_1465_v88_ = D14_DC31(d_1352_v3_, (self).f11, (self).f11)
                d_1466_v89_: _dafny.MultiSet
                d_1466_v89_ = _dafny.MultiSet([d_1459_v82_])
                index198_ = default__.safeIndex(585, (d_1461_v84_).length(0))
                (d_1461_v84_)[index198_] = default__.safeModuloInt((d_1465_v88_).cf46, default__.safeModuloInt((self).f11, (d_1466_v89_).cardinality))
                d_1467_v90_: _dafny.Array
                nw240_ = _dafny.Array(_dafny.MultiSet({}), 18)
                d_1467_v90_ = nw240_
                d_1468_v91_: _dafny.Array
                d_1468_v91_ = d_1467_v90_
                d_1467_v90_ = (d_1468_v91_)
                (globalState).f1 = (_dafny.MultiSet(((((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_))).set(default__.safeIndex(-764, len((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_)))), d_1352_v3_)).set(default__.safeIndex(default__.safeModuloInt((d_1461_v84_)[default__.safeIndex(585, (d_1461_v84_).length(0))], (d_1461_v84_)[default__.safeIndex(585, (d_1461_v84_).length(0))]), len(((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_))).set(default__.safeIndex(-764, len((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_)))), d_1352_v3_))), ((self).f11) == ((self).f11))).set(default__.safeIndex(-761, len((((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_))).set(default__.safeIndex(-764, len((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_)))), d_1352_v3_)).set(default__.safeIndex(default__.safeModuloInt((d_1461_v84_)[default__.safeIndex(585, (d_1461_v84_).length(0))], (d_1461_v84_)[default__.safeIndex(585, (d_1461_v84_).length(0))]), len(((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_))).set(default__.safeIndex(-764, len((d_1457_v80_) + ((d_1457_v80_) + (d_1457_v80_)))), d_1352_v3_))), ((self).f11) == ((self).f11)))), d_1352_v3_))).cardinality
            d_1469_v92_: _dafny.Array
            nw241_ = _dafny.Array(int(0), 14)
            d_1469_v92_ = nw241_
            index199_ = default__.safeIndex(440, (d_1469_v92_).length(0))
            (d_1469_v92_)[index199_] = (self).f11
            index200_ = default__.safeIndex(440, (d_1469_v92_).length(0))
            (d_1469_v92_)[index200_] = (self).fm3(((self).f11) - ((self).f11), d_1352_v3_, not(True), globalState)
            (globalState).f7 = ((d_1469_v92_)[default__.safeIndex(440, (d_1469_v92_).length(0))]) == ((self).f11)
            d_1470_v93_: _dafny.Seq
            d_1470_v93_ = _dafny.SeqWithoutIsStrInference([d_1469_v92_])
            d_1471_v94_: _dafny.MultiSet
            d_1471_v94_ = _dafny.MultiSet([d_1352_v3_])
            d_1472_v95_: _dafny.Map
            d_1472_v95_ = _dafny.Map({51: d_1352_v3_})
            d_1473_v96_: _dafny.Array
            nw242_ = _dafny.Array(None, 13)
            nw242_[int(0)] = d_1469_v92_
            nw242_[int(1)] = d_1469_v92_
            nw242_[int(2)] = d_1469_v92_
            nw242_[int(3)] = d_1469_v92_
            nw242_[int(4)] = d_1469_v92_
            nw242_[int(5)] = d_1469_v92_
            nw242_[int(6)] = (d_1470_v93_)[default__.safeIndex(((d_1471_v94_)[d_1352_v3_] if (d_1352_v3_) in (d_1471_v94_) else len(d_1472_v95_)), len(d_1470_v93_))]
            nw242_[int(7)] = d_1469_v92_
            nw242_[int(8)] = d_1469_v92_
            nw242_[int(9)] = d_1469_v92_
            nw242_[int(10)] = d_1469_v92_
            nw242_[int(11)] = d_1469_v92_
            nw242_[int(12)] = d_1469_v92_
            d_1473_v96_ = nw242_
            d_1474_v97_: _dafny.Array
            nw243_ = _dafny.Array(None, 8)
            nw243_[int(0)] = d_1473_v96_
            nw243_[int(1)] = d_1473_v96_
            nw243_[int(2)] = d_1473_v96_
            nw243_[int(3)] = d_1473_v96_
            nw243_[int(4)] = (d_1473_v96_)
            nw243_[int(5)] = d_1473_v96_
            nw243_[int(6)] = d_1473_v96_
            nw243_[int(7)] = d_1473_v96_
            d_1474_v97_ = nw243_
            d_1475_v98_: _dafny.Map
            d_1475_v98_ = _dafny.Map({(self).f11: d_1473_v96_})
            index201_ = default__.safeIndex(606, (d_1474_v97_).length(0))
            (d_1474_v97_)[index201_] = ((d_1475_v98_)[(d_1469_v92_)[default__.safeIndex(440, (d_1469_v92_).length(0))]] if ((d_1469_v92_)[default__.safeIndex(440, (d_1469_v92_).length(0))]) in (d_1475_v98_) else d_1473_v96_)
            index202_ = default__.safeIndex(606, (d_1474_v97_).length(0))
            nw244_ = _dafny.Array(_dafny.Array(None, 0), 29)
            (d_1474_v97_)[index202_] = nw244_
        d_1476_v99_: D15
        d_1476_v99_ = D15_DC36((self).f11, d_1352_v3_, d_1352_v3_, d_1352_v3_, (self).f11)
        (globalState).f1 = (0) - ((d_1476_v99_).cf55)
        d_1477_v100_: _dafny.Seq
        d_1477_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dohvyph"))
        hi1_ = len(d_1477_v100_)
        for d_1478_i6_ in range(((self).f11) * ((self).f11), hi1_):
            d_1479_v101_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_1479_v101_ = nw245_
            d_1480_v102_: _dafny.Array
            d_1480_v102_ = d_1479_v101_
            d_1481_v103_: _dafny.Array
            nw246_ = _dafny.Array(None, 8)
            nw246_[int(0)] = d_1480_v102_
            nw246_[int(1)] = d_1480_v102_
            nw246_[int(2)] = d_1479_v101_
            nw246_[int(3)] = d_1480_v102_
            nw246_[int(4)] = d_1480_v102_
            nw246_[int(5)] = d_1479_v101_
            nw246_[int(6)] = d_1480_v102_
            nw246_[int(7)] = d_1480_v102_
            d_1481_v103_ = nw246_
            d_1482_v104_: _dafny.Seq
            d_1482_v104_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f11), d_1478_i6_])
            d_1483_v105_: D0
            d_1483_v105_ = D0_DC0(len(d_1482_v104_))
            d_1484_v106_: _dafny.Seq
            d_1484_v106_ = _dafny.SeqWithoutIsStrInference([not(d_1352_v3_)])
            d_1485_v107_: _dafny.Seq
            def iife127_(_pat_let31_0):
                def iife128_(d_1486_dt__update__tmp_h0_):
                    def iife129_(_pat_let32_0):
                        def iife130_(d_1487_dt__update_hcf0_h0_):
                            return D0_DC0(d_1487_dt__update_hcf0_h0_)
                        return iife130_(_pat_let32_0)
                    return iife129_((self).f11)
                return iife128_(_pat_let31_0)
            d_1485_v107_ = _dafny.SeqWithoutIsStrInference([(self).fm5(iife127_(d_1483_v105_), d_1484_v106_, len(_dafny.SeqWithoutIsStrInference([p0 for d_1488_i7_ in range(default__.abs(481))])), p0, globalState)])
            d_1489_v108_: _dafny.MultiSet
            d_1489_v108_ = _dafny.MultiSet([len(d_1484_v106_)])
            d_1490_v109_: _dafny.Set
            d_1490_v109_ = _dafny.Set({(self).f11, d_1478_i6_})
            d_1491_v110_: _dafny.Map
            d_1491_v110_ = _dafny.Map({d_1352_v3_: 716})
            d_1492_v111_: _dafny.Seq
            d_1492_v111_ = _dafny.SeqWithoutIsStrInference([d_1482_v104_, d_1482_v104_])
            d_1493_v112_: _dafny.Array
            nw247_ = _dafny.Array(None, 21)
            nw247_[int(0)] = 840
            nw247_[int(1)] = len(_dafny.Map({(self).f11: d_1478_i6_}))
            nw247_[int(2)] = d_1478_i6_
            nw247_[int(3)] = (self).f11
            nw247_[int(4)] = (d_1489_v108_).cardinality
            nw247_[int(5)] = len(d_1490_v109_)
            nw247_[int(6)] = (self).f11
            nw247_[int(7)] = (self).f11
            nw247_[int(8)] = d_1478_i6_
            nw247_[int(9)] = ((d_1491_v110_)[d_1352_v3_] if (d_1352_v3_) in (d_1491_v110_) else d_1478_i6_)
            nw247_[int(10)] = d_1478_i6_
            nw247_[int(11)] = (self).f11
            nw247_[int(12)] = default__.fm2((self).f11, d_1478_i6_, globalState)
            nw247_[int(13)] = len(d_1485_v107_)
            nw247_[int(14)] = len(d_1492_v111_)
            nw247_[int(15)] = d_1478_i6_
            nw247_[int(16)] = d_1478_i6_
            nw247_[int(17)] = d_1478_i6_
            nw247_[int(18)] = d_1478_i6_
            nw247_[int(19)] = d_1478_i6_
            nw247_[int(20)] = d_1478_i6_
            d_1493_v112_ = nw247_
            d_1494_v113_: _dafny.Map
            d_1494_v113_ = _dafny.Map({d_1481_v103_: d_1493_v112_})
            d_1495_v114_: D11
            d_1495_v114_ = D11_DC22((d_1494_v113_) | (d_1494_v113_))
            source23_ = d_1495_v114_
            if source23_.is_DC23:
                d_1496___mcc_h10_ = source23_.cf31
                d_1497___mcc_h11_ = source23_.cf32
                d_1498___mcc_h12_ = source23_.cf33
                d_1499_cf33_ = d_1498___mcc_h12_
                d_1500_cf32_ = d_1497___mcc_h11_
                d_1501_cf31_ = d_1496___mcc_h10_
                d_1502_v115_: _dafny.MultiSet
                d_1502_v115_ = _dafny.MultiSet([d_1350_v1_, d_1350_v1_])
                d_1503_v116_: _dafny.Map
                d_1503_v116_ = _dafny.Map({d_1500_cf32_: p0})
                rhs263_ = (len(((d_1477_v100_) + (d_1477_v100_)).set(default__.safeIndex(d_1499_cf33_, len((d_1477_v100_) + (d_1477_v100_))), ((d_1503_v116_)[d_1499_cf33_] if (d_1499_cf33_) in (d_1503_v116_) else p0)))) - (d_1500_cf32_)
                rhs264_ = d_1502_v115_
                lhs217_ = globalState
                lhs217_.f1 = rhs263_
                d_1502_v115_ = rhs264_
                d_1483_v105_ = d_1483_v105_
                d_1489_v108_ = (d_1489_v108_) | (d_1489_v108_)
                d_1504_v117_: C2
                nw248_ = C2()
                nw248_.ctor__((p0 if d_1501_cf31_ else _dafny.CodePoint('n')), (d_1500_cf32_) * (d_1500_cf32_))
                d_1504_v117_ = nw248_
            elif source23_.is_DC24:
                d_1505___mcc_h13_ = source23_.cf34
                d_1506___mcc_h14_ = source23_.cf35
                d_1507___mcc_h15_ = source23_.cf36
                d_1508_cf36_ = d_1507___mcc_h15_
                d_1509_cf35_ = d_1506___mcc_h14_
                d_1510_cf34_ = d_1505___mcc_h13_
                d_1511_v118_: _dafny.Set
                d_1511_v118_ = _dafny.Set({p0})
                d_1511_v118_ = d_1511_v118_
                def lambda57_(d_1512_i8_):
                    return (d_1512_i8_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))))

                init30_ = lambda57_
                nw249_ = _dafny.Array(None, 9)
                for i0_30_ in range(nw249_.length(0)):
                    nw249_[i0_30_] = init30_(i0_30_)
                d_1493_v112_ = nw249_
                index203_ = default__.safeIndex(730, (d_1493_v112_).length(0))
                (d_1493_v112_)[index203_] = d_1478_i6_
                index204_ = default__.safeIndex(730, (d_1493_v112_).length(0))
                (d_1493_v112_)[index204_] = (0) - ((0) - (d_1478_i6_))
                d_1510_cf34_ = 700
            elif True:
                d_1513___mcc_h16_ = source23_.cf30
                d_1514_cf30_ = d_1513___mcc_h16_
                (globalState).f1 = (default__.fm2((self).f11, d_1478_i6_, globalState)) - (d_1478_i6_)
                d_1515_v119_: _dafny.Map
                d_1515_v119_ = _dafny.Map({d_1478_i6_: d_1478_i6_})
                rhs265_ = (d_1482_v104_) + (_dafny.SeqWithoutIsStrInference([(self).f11, d_1478_i6_, len(d_1477_v100_)]))
                rhs266_ = (default__.safeDivisionInt(d_1478_i6_, d_1478_i6_)) >= (d_1478_i6_)
                rhs267_ = d_1493_v112_
                rhs268_ = default__.safeDivisionInt(len(d_1515_v119_), (self).f11)
                rhs269_ = d_1350_v1_
                lhs218_ = globalState
                lhs219_ = globalState
                d_1482_v104_ = rhs265_
                lhs218_.f6 = rhs266_
                d_1493_v112_ = rhs267_
                lhs219_.f1 = rhs268_
                d_1350_v1_ = rhs269_
                d_1516_v120_: D11
                d_1516_v120_ = D11_DC24((self).f11, d_1485_v107_, p0)
                r0 = (_dafny.MultiSet((d_1516_v120_).cf35)) | (_dafny.MultiSet([False, True, d_1352_v3_, d_1352_v3_]))
                d_1517_v122_: _dafny.Seq
                def iife131_():
                    coll65_ = _dafny.Map()
                    compr_65_: int
                    for compr_65_ in (d_1482_v104_).Elements:
                        d_1518_v121_: int = compr_65_
                        if (d_1518_v121_) in (d_1482_v104_):
                            coll65_[(d_1518_v121_) - (d_1478_i6_)] = d_1352_v3_
                    return _dafny.Map(coll65_)
                d_1517_v122_ = _dafny.SeqWithoutIsStrInference([iife131_()
                ])
                d_1519_v123_: D0
                d_1519_v123_ = D0_DC1(d_1352_v3_, len((d_1517_v122_)[default__.safeIndex(d_1478_i6_, len(d_1517_v122_))]), d_1352_v3_)
                d_1520_v124_: _dafny.Array
                nw250_ = _dafny.Array(None, 10)
                nw250_[int(0)] = d_1519_v123_
                nw250_[int(1)] = default__.fm49(d_1352_v3_, globalState)
                nw250_[int(2)] = d_1519_v123_
                nw250_[int(3)] = d_1519_v123_
                nw250_[int(4)] = d_1519_v123_
                nw250_[int(5)] = d_1519_v123_
                nw250_[int(6)] = d_1519_v123_
                nw250_[int(7)] = d_1519_v123_
                nw250_[int(8)] = d_1519_v123_
                nw250_[int(9)] = d_1519_v123_
                d_1520_v124_ = nw250_
                d_1521_v125_: _dafny.MultiSet
                d_1521_v125_ = _dafny.MultiSet([d_1520_v124_])
                d_1522_v126_: _dafny.Seq
                d_1522_v126_ = _dafny.SeqWithoutIsStrInference([d_1520_v124_, d_1520_v124_])
                d_1523_v127_: _dafny.Seq
                d_1523_v127_ = _dafny.SeqWithoutIsStrInference([d_1520_v124_, d_1520_v124_, (d_1522_v126_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0 for d_1524_i9_ in range(default__.abs(-797))])), len(d_1522_v126_))], d_1520_v124_])
                d_1521_v125_ = (_dafny.MultiSet([d_1520_v124_, d_1520_v124_, d_1520_v124_, (d_1523_v127_)[default__.safeIndex((self).f11, len(d_1523_v127_))], d_1520_v124_])) - (d_1521_v125_)
            (globalState).f1 = d_1478_i6_
            d_1525_v128_: _dafny.Map
            d_1525_v128_ = _dafny.Map({d_1352_v3_: p0})
            d_1526_v129_: D11
            d_1526_v129_ = D11_DC23(d_1352_v3_, d_1478_i6_, (self).f11)
            d_1527_v130_: _dafny.Array
            nw251_ = _dafny.Array(None, 17)
            nw251_[int(0)] = p0
            nw251_[int(1)] = (d_1477_v100_)[default__.safeIndex(d_1478_i6_, len(d_1477_v100_))]
            nw251_[int(2)] = _dafny.CodePoint('l')
            nw251_[int(3)] = p0
            nw251_[int(4)] = _dafny.CodePoint('q')
            nw251_[int(5)] = _dafny.CodePoint('e')
            nw251_[int(6)] = ((d_1525_v128_)[(d_1526_v129_).cf31] if ((d_1526_v129_).cf31) in (d_1525_v128_) else p0)
            nw251_[int(7)] = p0
            nw251_[int(8)] = p0
            nw251_[int(9)] = p0
            nw251_[int(10)] = p0
            nw251_[int(11)] = p0
            nw251_[int(12)] = _dafny.CodePoint('v')
            nw251_[int(13)] = p0
            nw251_[int(14)] = (p0 if d_1352_v3_ else p0)
            nw251_[int(15)] = _dafny.CodePoint('l')
            nw251_[int(16)] = p0
            d_1527_v130_ = nw251_
            index205_ = default__.safeIndex(502, (d_1527_v130_).length(0))
            (d_1527_v130_)[index205_] = p0
            index206_ = default__.safeIndex(502, (d_1527_v130_).length(0))
            (d_1527_v130_)[index206_] = p0
            d_1477_v100_ = d_1477_v100_
        d_1528_v131_: _dafny.MultiSet
        d_1528_v131_ = _dafny.MultiSet([d_1352_v3_, d_1352_v3_, True])
        r0 = d_1528_v131_
        return r0

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        (globalState).f1 = default__.safeDivisionInt((self).f11, (0) - (p0))
        d_1529_v0_: _dafny.Array
        nw252_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
        d_1529_v0_ = nw252_
        d_1530_v1_: _dafny.Seq
        d_1530_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sva"))
        index207_ = default__.safeIndex(232, (d_1529_v0_).length(0))
        (d_1529_v0_)[index207_] = d_1530_v1_
        index208_ = default__.safeIndex(232, (d_1529_v0_).length(0))
        (d_1529_v0_)[index208_] = (default__.fm1(globalState)) + ((d_1530_v1_) + (d_1530_v1_))
        d_1531_v2_: bool
        d_1531_v2_ = True
        if not(d_1531_v2_):
            d_1532_v3_: _dafny.Seq
            d_1532_v3_ = _dafny.SeqWithoutIsStrInference([d_1531_v2_])
            d_1533_v4_: _dafny.MultiSet
            d_1533_v4_ = _dafny.MultiSet([d_1531_v2_])
            source24_ = default__.fm50(len(d_1532_v3_), d_1531_v2_, d_1531_v2_, (d_1532_v3_)[default__.safeIndex(((d_1533_v4_)[d_1531_v2_] if (d_1531_v2_) in (d_1533_v4_) else p0), len(d_1532_v3_))], globalState)
            if source24_.is_DC23:
                d_1534___mcc_h0_ = source24_.cf31
                d_1535___mcc_h1_ = source24_.cf32
                d_1536___mcc_h2_ = source24_.cf33
                d_1537_cf33_ = d_1536___mcc_h2_
                d_1538_cf32_ = d_1535___mcc_h1_
                d_1539_cf31_ = d_1534___mcc_h0_
                d_1540_v5_: _dafny.Array
                def lambda58_(d_1541_v2_, d_1542_cf31_):
                    def lambda59_(d_1543_i0_):
                        return (_dafny.Map({d_1541_v2_: d_1541_v2_})) | (_dafny.Map({d_1542_cf31_: False}))

                    return lambda59_

                init31_ = lambda58_(d_1531_v2_, d_1539_cf31_)
                nw253_ = _dafny.Array(None, 13)
                for i0_31_ in range(nw253_.length(0)):
                    nw253_[i0_31_] = init31_(i0_31_)
                d_1540_v5_ = nw253_
                d_1540_v5_ = d_1540_v5_
                (globalState).f7 = d_1539_cf31_
                d_1544_v6_: _dafny.Map
                d_1544_v6_ = _dafny.Map({d_1531_v2_: d_1539_cf31_})
                d_1545_v7_: _dafny.Seq
                d_1545_v7_ = _dafny.SeqWithoutIsStrInference([d_1544_v6_])
                d_1546_v8_: D8
                d_1546_v8_ = D8_DC15((d_1545_v7_)[default__.safeIndex(178, len(d_1545_v7_))])
                d_1546_v8_ = d_1546_v8_
                d_1547_v9_: _dafny.Array
                def lambda60_(d_1548_v3_, d_1549_v1_, d_1550_cf31_):
                    def lambda61_(d_1551_i1_):
                        return (d_1548_v3_).set(default__.safeIndex(len(d_1549_v1_), len(d_1548_v3_)), not(d_1550_cf31_))

                    return lambda61_

                init32_ = lambda60_(d_1532_v3_, d_1530_v1_, d_1539_cf31_)
                nw254_ = _dafny.Array(None, 4)
                for i0_32_ in range(nw254_.length(0)):
                    nw254_[i0_32_] = init32_(i0_32_)
                d_1547_v9_ = nw254_
                index209_ = default__.safeIndex(449, (d_1547_v9_).length(0))
                (d_1547_v9_)[index209_] = d_1532_v3_
                index210_ = default__.safeIndex(449, (d_1547_v9_).length(0))
                rhs270_ = default__.fm19(d_1539_cf31_, not(d_1539_cf31_), globalState)
                rhs271_ = d_1531_v2_
                lhs220_ = d_1547_v9_
                lhs221_ = default__.safeIndex(449, (d_1547_v9_).length(0))
                lhs220_[lhs221_] = rhs270_
                d_1531_v2_ = rhs271_
            elif source24_.is_DC24:
                d_1552___mcc_h3_ = source24_.cf34
                d_1553___mcc_h4_ = source24_.cf35
                d_1554___mcc_h5_ = source24_.cf36
                d_1555_cf36_ = d_1554___mcc_h5_
                d_1556_cf35_ = d_1553___mcc_h4_
                d_1557_cf34_ = d_1552___mcc_h3_
                d_1558_v10_: _dafny.Array
                nw255_ = _dafny.Array(int(0), 17)
                d_1558_v10_ = nw255_
                index211_ = default__.safeIndex(947, (d_1558_v10_).length(0))
                (d_1558_v10_)[index211_] = default__.safeModuloInt((0) - ((self).f11), p0)
                index212_ = default__.safeIndex(947, (d_1558_v10_).length(0))
                (d_1558_v10_)[index212_] = (self).f11
                d_1558_v10_ = d_1558_v10_
                d_1559_v11_: _dafny.Set
                d_1559_v11_ = _dafny.Set({(d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))], d_1530_v1_})
                d_1559_v11_ = d_1559_v11_
                d_1560_v12_: _dafny.Map
                d_1560_v12_ = _dafny.Map({(self).f11: (d_1530_v1_) + ((d_1530_v1_).set(default__.safeIndex(p0, len(d_1530_v1_)), d_1555_cf36_))})
                d_1560_v12_ = (d_1560_v12_).set((self).f11, (d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])
            elif True:
                d_1561___mcc_h6_ = source24_.cf30
                d_1562_cf30_ = d_1561___mcc_h6_
                d_1563_v13_: _dafny.Array
                nw256_ = _dafny.Array(None, 6)
                nw256_[int(0)] = p0
                nw256_[int(1)] = (p0) * (p0)
                nw256_[int(2)] = p0
                nw256_[int(3)] = default__.fm2((self).f11, (_dafny.MultiSet(d_1532_v3_)).cardinality, globalState)
                nw256_[int(4)] = len((self).f12)
                nw256_[int(5)] = default__.safeDivisionInt((self).f11, (0) - ((self).f11))
                d_1563_v13_ = nw256_
                index213_ = default__.safeIndex(334, (d_1563_v13_).length(0))
                (d_1563_v13_)[index213_] = (self).f11
                index214_ = default__.safeIndex(629, (d_1563_v13_).length(0))
                (d_1563_v13_)[index214_] = (self).f11
                d_1564_v14_: _dafny.Array
                nw257_ = _dafny.Array(_dafny.Map({}), 10)
                d_1564_v14_ = nw257_
                d_1565_v15_: _dafny.Map
                d_1565_v15_ = _dafny.Map({d_1564_v14_: len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])})
                d_1566_v16_: _dafny.MultiSet
                d_1566_v16_ = _dafny.MultiSet([(self).f11, (self).f11, (self).f11, (0) - (p0)])
                index215_ = default__.safeIndex(334, (d_1563_v13_).length(0))
                index216_ = default__.safeIndex(629, (d_1563_v13_).length(0))
                rhs272_ = ((d_1565_v15_)[d_1564_v14_] if (d_1564_v14_) in (d_1565_v15_) else len((d_1532_v3_) + (_dafny.SeqWithoutIsStrInference([d_1531_v2_]))))
                rhs273_ = False
                rhs274_ = (0) - (p0)
                rhs275_ = (d_1566_v16_).issubset((d_1566_v16_).set((self).f11, default__.abs((self).f11)))
                lhs222_ = d_1563_v13_
                lhs223_ = default__.safeIndex(334, (d_1563_v13_).length(0))
                lhs224_ = globalState
                lhs225_ = d_1563_v13_
                lhs226_ = default__.safeIndex(629, (d_1563_v13_).length(0))
                lhs227_ = globalState
                lhs222_[lhs223_] = rhs272_
                lhs224_.f7 = rhs273_
                lhs225_[lhs226_] = rhs274_
                lhs227_.f7 = rhs275_
                d_1567_v17_: _dafny.Map
                d_1567_v17_ = _dafny.Map({d_1531_v2_: p0})
                d_1568_v18_: _dafny.Set
                d_1568_v18_ = _dafny.Set({d_1567_v17_})
                d_1568_v18_ = _dafny.Set({d_1567_v17_, d_1567_v17_})
                d_1569_v19_: _dafny.Map
                d_1569_v19_ = _dafny.Map({len((d_1532_v3_) + (d_1532_v3_)): d_1532_v3_})
                d_1570_v20_: C1
                nw258_ = C1()
                nw258_.ctor__(d_1531_v2_, d_1531_v2_)
                d_1570_v20_ = nw258_
                d_1571_v21_: _dafny.Map
                d_1571_v21_ = d_1569_v19_
                d_1572_v22_: D13
                d_1572_v22_ = D13_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqjjbmfo")), (d_1563_v13_)[default__.safeIndex(334, (d_1563_v13_).length(0))], (d_1570_v20_).f18)
                d_1573_v23_: _dafny.Map
                d_1573_v23_ = _dafny.Map({(d_1531_v2_ if (d_1570_v20_).f18 else (d_1570_v20_).f18): d_1570_v20_})
                rhs276_ = ((d_1571_v21_) if d_1531_v2_ else (d_1569_v19_).set((d_1572_v22_).cf41, d_1532_v3_))
                rhs277_ = ((d_1573_v23_)[(d_1570_v20_).f19] if ((d_1570_v20_).f19) in (d_1573_v23_) else d_1570_v20_)
                d_1569_v19_ = rhs276_
                d_1570_v20_ = rhs277_
                d_1574_v24_: _dafny.Seq
                d_1574_v24_ = _dafny.SeqWithoutIsStrInference([(d_1563_v13_)[default__.safeIndex(334, (d_1563_v13_).length(0))], -974])
                d_1575_v25_: _dafny.Map
                d_1575_v25_ = _dafny.Map({(d_1563_v13_)[default__.safeIndex(334, (d_1563_v13_).length(0))]: (len(d_1574_v24_)) * (p0)})
                d_1575_v25_ = (d_1575_v25_).set(799, (d_1563_v13_)[default__.safeIndex(334, (d_1563_v13_).length(0))])
            d_1576_v26_: _dafny.Map
            d_1576_v26_ = _dafny.Map({len(_dafny.Map({True: len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])})): p0})
            d_1577_v27_: D16
            d_1577_v27_ = D16_DC38(d_1576_v26_)
            d_1578_v28_: _dafny.Map
            d_1578_v28_ = _dafny.Map({d_1577_v27_: d_1531_v2_})
            pat_let_tv43_ = d_1576_v26_
            d_1579_v29_: _dafny.Map
            def iife132_(_pat_let33_0):
                def iife133_(d_1580_dt__update__tmp_h0_):
                    def iife134_(_pat_let34_0):
                        def iife135_(d_1581_dt__update_hcf61_h0_):
                            return D16_DC38(d_1581_dt__update_hcf61_h0_)
                        return iife135_(_pat_let34_0)
                    return iife134_(pat_let_tv43_)
                return iife133_(_pat_let33_0)
            d_1579_v29_ = _dafny.Map({d_1531_v2_: _dafny.Map({(self).f12: (d_1578_v28_).set(iife132_(D16_DC38(default__.fm46(globalState))), False)})})
            d_1582_v30_: _dafny.Map
            d_1582_v30_ = _dafny.Map({(self).f12: d_1578_v28_})
            d_1579_v29_ = (d_1579_v29_).set(d_1531_v2_, d_1582_v30_)
            d_1583_v31_: _dafny.Array
            nw259_ = _dafny.Array(None, 13)
            d_1583_v31_ = nw259_
            d_1584_v32_: _dafny.Map
            d_1584_v32_ = _dafny.Map({d_1583_v31_: (0) - ((0) - ((self).f11))})
            d_1585_v33_: _dafny.Array
            nw260_ = _dafny.Array(int(0), 23)
            d_1585_v33_ = nw260_
            index217_ = default__.safeIndex(974, (d_1585_v33_).length(0))
            (d_1585_v33_)[index217_] = (self).f11
            index218_ = default__.safeIndex(974, (d_1585_v33_).length(0))
            rhs278_ = _dafny.Map({d_1583_v31_: (p0) + (p0)})
            rhs279_ = not(d_1531_v2_)
            rhs280_ = default__.safeDivisionInt((self).f11, (self).f11)
            rhs281_ = ((0) - (len(d_1532_v3_)) if d_1531_v2_ else 183)
            lhs228_ = globalState
            lhs229_ = globalState
            lhs230_ = d_1585_v33_
            lhs231_ = default__.safeIndex(974, (d_1585_v33_).length(0))
            d_1584_v32_ = rhs278_
            lhs228_.f7 = rhs279_
            lhs229_.f1 = rhs280_
            lhs230_[lhs231_] = rhs281_
            d_1586_v34_: _dafny.Seq
            d_1586_v34_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            (globalState).f1 = (self).fm3(len(d_1586_v34_), d_1531_v2_, True, globalState)
            (globalState).f7 = d_1531_v2_
        elif True:
            d_1587_v35_: _dafny.Array
            nw261_ = _dafny.Array(int(0), 2)
            d_1587_v35_ = nw261_
            d_1588_v36_: _dafny.Seq
            d_1588_v36_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
            d_1589_v37_: D14
            d_1589_v37_ = D14_DC31(d_1531_v2_, (self).f11, (self).f11)
            d_1590_v38_: D14
            d_1590_v38_ = D14_DC32(d_1589_v37_)
            d_1591_v39_: _dafny.Seq
            d_1591_v39_ = _dafny.SeqWithoutIsStrInference([d_1590_v38_])
            index219_ = default__.safeIndex(594, (d_1587_v35_).length(0))
            (d_1587_v35_)[index219_] = len((default__.fm51(d_1588_v36_, p0, p0, p0, globalState)) + (d_1591_v39_))
            index220_ = default__.safeIndex(594, (d_1587_v35_).length(0))
            (d_1587_v35_)[index220_] = (0) - ((self).f11)
            if d_1531_v2_:
                d_1592_v40_: _dafny.Array
                nw262_ = _dafny.Array(False, 1)
                d_1592_v40_ = nw262_
                index221_ = default__.safeIndex(225, (d_1592_v40_).length(0))
                (d_1592_v40_)[index221_] = d_1531_v2_
                index222_ = default__.safeIndex(225, (d_1592_v40_).length(0))
                (d_1592_v40_)[index222_] = d_1531_v2_
                d_1588_v36_ = d_1588_v36_
                (globalState).f1 = (0) - ((self).f11)
                d_1593_v41_: _dafny.Array
                nw263_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1593_v41_ = nw263_
                index223_ = default__.safeIndex(845, (d_1593_v41_).length(0))
                (d_1593_v41_)[index223_] = d_1592_v40_
                index224_ = default__.safeIndex(845, (d_1593_v41_).length(0))
                (d_1593_v41_)[index224_] = d_1592_v40_
                d_1594_v42_: _dafny.Seq
                d_1594_v42_ = _dafny.SeqWithoutIsStrInference([(d_1592_v40_)[default__.safeIndex(225, (d_1592_v40_).length(0))]])
                d_1595_v43_: C0
                nw264_ = C0()
                nw264_.ctor__(d_1594_v42_, 225)
                d_1595_v43_ = nw264_
            elif True:
                d_1587_v35_ = d_1587_v35_
                (globalState).f1 = (self).f11
                d_1596_v44_: D19
                d_1596_v44_ = D19_DC42(d_1587_v35_)
                d_1597_v45_: _dafny.Array
                nw265_ = _dafny.Array(None, 20)
                nw265_[int(0)] = d_1587_v35_
                nw265_[int(1)] = d_1587_v35_
                nw265_[int(2)] = d_1587_v35_
                nw265_[int(3)] = d_1587_v35_
                nw265_[int(4)] = d_1587_v35_
                nw265_[int(5)] = d_1587_v35_
                nw265_[int(6)] = d_1587_v35_
                nw265_[int(7)] = d_1587_v35_
                nw265_[int(8)] = d_1587_v35_
                nw265_[int(9)] = d_1587_v35_
                nw265_[int(10)] = d_1587_v35_
                nw265_[int(11)] = d_1587_v35_
                nw265_[int(12)] = d_1587_v35_
                nw265_[int(13)] = d_1587_v35_
                nw265_[int(14)] = d_1587_v35_
                nw265_[int(15)] = d_1587_v35_
                nw265_[int(16)] = d_1587_v35_
                nw265_[int(17)] = d_1587_v35_
                nw265_[int(18)] = d_1587_v35_
                nw265_[int(19)] = (d_1596_v44_).cf68
                d_1597_v45_ = nw265_
                index225_ = default__.safeIndex(948, (d_1597_v45_).length(0))
                (d_1597_v45_)[index225_] = d_1587_v35_
                index226_ = default__.safeIndex(948, (d_1597_v45_).length(0))
                (d_1597_v45_)[index226_] = d_1587_v35_
                d_1598_v46_: str
                d_1598_v46_ = _dafny.CodePoint('l')
                d_1599_v47_: _dafny.Seq
                d_1599_v47_ = _dafny.SeqWithoutIsStrInference([d_1531_v2_, d_1531_v2_, d_1531_v2_])
                rhs282_ = (d_1531_v2_ if d_1531_v2_ else (d_1599_v47_)[default__.safeIndex((self).f11, len(d_1599_v47_))])
                rhs283_ = d_1598_v46_
                d_1531_v2_ = rhs282_
                d_1598_v46_ = rhs283_
                (globalState).f1 = p0
            (globalState).f1 = (self).f11
            (globalState).f7 = (d_1531_v2_) or (d_1531_v2_)
            (globalState).f7 = default__.fm0(globalState)
        if (not(d_1531_v2_) if d_1531_v2_ else (p0) != (522)):
            d_1600_v48_: D15
            d_1600_v48_ = D15_DC34()
            d_1600_v48_ = (d_1600_v48_ if d_1531_v2_ else (d_1600_v48_ if d_1531_v2_ else d_1600_v48_))
            if d_1531_v2_:
                (globalState).f6 = (((self).f12) - ((self).f12)).issubset((self).f12)
                d_1530_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcrg"))
                d_1601_v49_: str
                d_1601_v49_ = _dafny.CodePoint('h')
                d_1601_v49_ = (d_1530_v1_)[default__.safeIndex(p0, len(d_1530_v1_))]
                d_1602_v50_: D13
                d_1602_v50_ = D13_DC28((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaqirwuan"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaqirwuan")))), d_1601_v49_), (self).f11, default__.fm0(globalState))
                (globalState).f6 = (d_1602_v50_).cf42
                d_1603_v51_: _dafny.Array
                nw266_ = _dafny.Array(None, 3)
                nw266_[int(0)] = (self).f11
                nw266_[int(1)] = p0
                nw266_[int(2)] = p0
                d_1603_v51_ = nw266_
                d_1604_v52_: _dafny.Seq
                d_1604_v52_ = _dafny.SeqWithoutIsStrInference([d_1603_v51_, d_1603_v51_])
                d_1605_v53_: _dafny.MultiSet
                d_1605_v53_ = _dafny.MultiSet([591])
                d_1606_v54_: _dafny.Map
                d_1606_v54_ = _dafny.Map({d_1531_v2_: d_1531_v2_})
                d_1604_v52_ = (d_1604_v52_).set(default__.safeIndex(((d_1605_v53_)[len(d_1606_v54_)] if (len(d_1606_v54_)) in (d_1605_v53_) else (self).f11), len(d_1604_v52_)), d_1603_v51_)
            elif True:
                d_1607_v55_: _dafny.Array
                nw267_ = _dafny.Array(None, 7)
                d_1607_v55_ = nw267_
                d_1608_v56_: _dafny.Seq
                d_1608_v56_ = _dafny.SeqWithoutIsStrInference([default__.fm52((0) - ((self).f11), (self).f11, globalState)])
                d_1609_v57_: _dafny.Map
                d_1609_v57_ = _dafny.Map({d_1607_v55_: (d_1608_v56_) + (d_1608_v56_)})
                d_1609_v57_ = (d_1609_v57_).set(d_1607_v55_, (_dafny.SeqWithoutIsStrInference([_dafny.Map({D0_DC1(d_1531_v2_, len(_dafny.SeqWithoutIsStrInference([d_1531_v2_, d_1531_v2_])), not(d_1531_v2_)): d_1531_v2_}) for d_1610_i2_ in range(default__.abs(-320))])) + (d_1608_v56_))
                (globalState).f1 = ((self).f11) - (p0)
                d_1611_v58_: _dafny.Seq
                d_1611_v58_ = _dafny.SeqWithoutIsStrInference([p0])
                (globalState).f6 = not((len(d_1611_v58_)) >= (((self).f11) * (p0)))
                (globalState).f1 = (p0) - ((self).f11)
                (globalState).f1 = default__.safeModuloInt(p0, (self).f11)
            d_1612_v59_: _dafny.Seq
            d_1612_v59_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1613_v60_: _dafny.Set
            d_1613_v60_ = _dafny.Set({d_1612_v59_, (d_1612_v59_) + (d_1612_v59_)})
            d_1613_v60_ = (d_1613_v60_) - (d_1613_v60_)
            d_1614_v61_: D0
            d_1614_v61_ = D0_DC0((self).f11)
            d_1615_v62_: _dafny.Seq
            d_1615_v62_ = _dafny.SeqWithoutIsStrInference([False])
            d_1616_v63_: _dafny.MultiSet
            d_1616_v63_ = _dafny.MultiSet([p0])
            d_1617_v64_: str
            d_1617_v64_ = _dafny.CodePoint('r')
            if (self).fm5(d_1614_v61_, d_1615_v62_, (d_1616_v63_).cardinality, d_1617_v64_, globalState):
                (globalState).f7 = d_1531_v2_
                index227_ = default__.safeIndex(232, (d_1529_v0_).length(0))
                (d_1529_v0_)[index227_] = ((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))]) + (((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))]).set(default__.safeIndex(len(d_1612_v59_), len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])), d_1617_v64_))
                d_1618_v65_: _dafny.Array
                nw268_ = _dafny.Array(None, 29)
                d_1618_v65_ = nw268_
                d_1619_v66_: C4
                nw269_ = C4()
                nw269_.ctor__(d_1616_v63_, d_1531_v2_, (self).f12, -299)
                d_1619_v66_ = nw269_
                index228_ = default__.safeIndex(351, (d_1618_v65_).length(0))
                (d_1618_v65_)[index228_] = (d_1619_v66_ if d_1531_v2_ else d_1619_v66_)
                d_1620_v67_: _dafny.MultiSet
                d_1620_v67_ = _dafny.MultiSet([(d_1619_v66_).f17, True, d_1531_v2_, (d_1619_v66_).f17, True])
                index229_ = default__.safeIndex(351, (d_1618_v65_).length(0))
                nw270_ = C4()
                nw270_.ctor__((d_1616_v63_ if False else _dafny.MultiSet([(d_1612_v59_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f11])), len(d_1612_v59_))]])), d_1531_v2_, (self).f12, (0) - ((p0) * (((d_1620_v67_)[True] if (True) in (d_1620_v67_) else p0))))
                (d_1618_v65_)[index229_] = nw270_
                d_1621_v68_: _dafny.Array
                nw271_ = _dafny.Array(int(0), 26)
                d_1621_v68_ = nw271_
                index230_ = default__.safeIndex(865, (d_1621_v68_).length(0))
                (d_1621_v68_)[index230_] = (self).f11
                d_1622_v69_: _dafny.Array
                def lambda62_(d_1623_v2_):
                    def lambda63_(d_1624_i3_):
                        return d_1623_v2_

                    return lambda63_

                init33_ = lambda62_(d_1531_v2_)
                nw272_ = _dafny.Array(None, 1)
                for i0_33_ in range(nw272_.length(0)):
                    nw272_[i0_33_] = init33_(i0_33_)
                d_1622_v69_ = nw272_
                d_1625_v70_: _dafny.Array
                nw273_ = _dafny.Array(None, 3)
                nw273_[int(0)] = d_1622_v69_
                nw273_[int(1)] = d_1622_v69_
                nw273_[int(2)] = d_1622_v69_
                d_1625_v70_ = nw273_
                index231_ = default__.safeIndex(772, (d_1625_v70_).length(0))
                (d_1625_v70_)[index231_] = d_1622_v69_
                d_1626_v71_: _dafny.Map
                d_1626_v71_ = _dafny.Map({len((d_1615_v62_).set(default__.safeIndex(p0, len(d_1615_v62_)), (d_1619_v66_).f17)): p0})
                d_1627_v72_: _dafny.Seq
                d_1627_v72_ = _dafny.SeqWithoutIsStrInference([d_1626_v71_, d_1626_v71_])
                d_1628_v73_: _dafny.Map
                d_1628_v73_ = _dafny.Map({d_1531_v2_: len((d_1627_v72_)[default__.safeIndex(p0, len(d_1627_v72_))])})
                d_1629_v74_: _dafny.Map
                d_1629_v74_ = _dafny.Map({p0: (d_1619_v66_).f17})
                d_1630_v75_: _dafny.Map
                d_1630_v75_ = _dafny.Map({not(d_1531_v2_): d_1531_v2_})
                index232_ = default__.safeIndex(865, (d_1621_v68_).length(0))
                index233_ = default__.safeIndex(772, (d_1625_v70_).length(0))
                rhs284_ = ((d_1628_v73_)[not ((d_1619_v66_).f17) or (((d_1629_v74_)[len(d_1630_v75_)] if (len(d_1630_v75_)) in (d_1629_v74_) else d_1531_v2_))] if (not ((d_1619_v66_).f17) or (((d_1629_v74_)[len(d_1630_v75_)] if (len(d_1630_v75_)) in (d_1629_v74_) else d_1531_v2_))) in (d_1628_v73_) else (self).f11)
                rhs285_ = d_1622_v69_
                rhs286_ = ((d_1619_v66_).f17) or (default__.fm0(globalState))
                rhs287_ = (d_1615_v62_)[default__.safeIndex(p0, len(d_1615_v62_))]
                rhs288_ = (d_1619_v66_).f17
                lhs232_ = d_1621_v68_
                lhs233_ = default__.safeIndex(865, (d_1621_v68_).length(0))
                lhs234_ = d_1625_v70_
                lhs235_ = default__.safeIndex(772, (d_1625_v70_).length(0))
                lhs236_ = globalState
                lhs237_ = globalState
                lhs238_ = globalState
                lhs232_[lhs233_] = rhs284_
                lhs234_[lhs235_] = rhs285_
                lhs236_.f6 = rhs286_
                lhs237_.f7 = rhs287_
                lhs238_.f7 = rhs288_
                d_1631_v76_: _dafny.Array
                nw274_ = _dafny.Array(None, 18)
                nw274_[int(0)] = d_1621_v68_
                nw274_[int(1)] = d_1621_v68_
                nw274_[int(2)] = d_1621_v68_
                nw274_[int(3)] = d_1621_v68_
                nw274_[int(4)] = d_1621_v68_
                nw274_[int(5)] = d_1621_v68_
                nw274_[int(6)] = d_1621_v68_
                nw274_[int(7)] = d_1621_v68_
                nw274_[int(8)] = d_1621_v68_
                nw274_[int(9)] = d_1621_v68_
                nw274_[int(10)] = d_1621_v68_
                nw274_[int(11)] = d_1621_v68_
                nw274_[int(12)] = d_1621_v68_
                nw274_[int(13)] = d_1621_v68_
                nw274_[int(14)] = d_1621_v68_
                nw274_[int(15)] = d_1621_v68_
                nw274_[int(16)] = d_1621_v68_
                nw274_[int(17)] = d_1621_v68_
                d_1631_v76_ = nw274_
                d_1632_v77_: _dafny.Array
                d_1632_v77_ = d_1631_v76_
                d_1633_v78_: _dafny.Array
                nw275_ = _dafny.Array(None, 16)
                nw275_[int(0)] = d_1632_v77_
                nw275_[int(1)] = d_1632_v77_
                nw275_[int(2)] = d_1632_v77_
                nw275_[int(3)] = d_1632_v77_
                nw275_[int(4)] = d_1632_v77_
                nw275_[int(5)] = d_1632_v77_
                nw275_[int(6)] = d_1631_v76_
                nw275_[int(7)] = d_1632_v77_
                nw275_[int(8)] = d_1632_v77_
                nw275_[int(9)] = d_1632_v77_
                nw275_[int(10)] = d_1632_v77_
                nw275_[int(11)] = d_1632_v77_
                nw275_[int(12)] = d_1632_v77_
                nw275_[int(13)] = d_1632_v77_
                nw275_[int(14)] = d_1632_v77_
                nw275_[int(15)] = d_1632_v77_
                d_1633_v78_ = nw275_
                d_1634_v79_: _dafny.Map
                d_1634_v79_ = _dafny.Map({d_1633_v78_: d_1621_v68_})
                d_1635_v80_: D11
                d_1635_v80_ = D11_DC22(d_1634_v79_)
                pat_let_tv44_ = d_1634_v79_
                d_1636_v81_: D11
                def iife136_(_pat_let35_0):
                    def iife137_(d_1637_dt__update__tmp_h1_):
                        def iife138_(_pat_let36_0):
                            def iife139_(d_1638_dt__update_hcf30_h0_):
                                return D11_DC22(d_1638_dt__update_hcf30_h0_)
                            return iife139_(_pat_let36_0)
                        return iife138_(pat_let_tv44_)
                    return iife137_(_pat_let35_0)
                d_1636_v81_ = D11_DC22((iife136_(d_1635_v80_)).cf30)
                arr2_ = (d_1625_v70_)[default__.safeIndex(772, (d_1625_v70_).length(0))]
                index234_ = default__.safeIndex(828, ((d_1625_v70_)[default__.safeIndex(772, (d_1625_v70_).length(0))]).length(0))
                arr2_[index234_] = not ((d_1619_v66_).f17) or ((d_1619_v66_).f17)
                index235_ = default__.safeIndex(865, (d_1621_v68_).length(0))
                arr3_ = (d_1625_v70_)[default__.safeIndex(772, (d_1625_v70_).length(0))]
                index236_ = default__.safeIndex(828, ((d_1625_v70_)[default__.safeIndex(772, (d_1625_v70_).length(0))]).length(0))
                rhs289_ = (default__.safeModuloInt((d_1621_v68_)[default__.safeIndex(865, (d_1621_v68_).length(0))], (d_1619_v66_).fm6((d_1619_v66_).f17, len(d_1626_v71_), globalState))) * ((d_1621_v68_)[default__.safeIndex(865, (d_1621_v68_).length(0))])
                rhs290_ = len(d_1530_v1_)
                rhs291_ = d_1635_v80_
                rhs292_ = (self).f11
                rhs293_ = (d_1619_v66_).f17
                lhs239_ = d_1621_v68_
                lhs240_ = default__.safeIndex(865, (d_1621_v68_).length(0))
                lhs241_ = globalState
                lhs242_ = globalState
                lhs243_ = (d_1625_v70_)[default__.safeIndex(772, (d_1625_v70_).length(0))]
                lhs244_ = default__.safeIndex(828, ((d_1625_v70_)[default__.safeIndex(772, (d_1625_v70_).length(0))]).length(0))
                lhs239_[lhs240_] = rhs289_
                lhs241_.f1 = rhs290_
                d_1635_v80_ = rhs291_
                lhs242_.f1 = rhs292_
                lhs243_[lhs244_] = rhs293_
            elif True:
                d_1615_v62_ = (d_1615_v62_) + (d_1615_v62_)
                d_1639_v82_: _dafny.MultiSet
                d_1639_v82_ = _dafny.MultiSet([d_1531_v2_, d_1531_v2_, d_1531_v2_, d_1531_v2_, d_1531_v2_])
                d_1640_v83_: _dafny.Map
                d_1640_v83_ = _dafny.Map({p0: d_1639_v82_})
                d_1641_v84_: _dafny.Seq
                d_1641_v84_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1531_v2_]), ((d_1640_v83_)[(self).f11] if ((self).f11) in (d_1640_v83_) else d_1639_v82_)])
                d_1641_v84_ = _dafny.SeqWithoutIsStrInference([default__.fm37(globalState)])
                (globalState).f1 = default__.safeModuloInt((self).fm6(True, (self).f11, globalState), p0)
                d_1531_v2_ = d_1531_v2_
                rhs294_ = (self).fm6(d_1531_v2_, len((d_1612_v59_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1642_i4_ in range(default__.abs(95))]))), globalState)
                rhs295_ = (d_1612_v59_)[default__.safeIndex(default__.safeModuloInt(p0, (self).f11), len(d_1612_v59_))]
                rhs296_ = d_1531_v2_
                rhs297_ = 474
                lhs245_ = globalState
                lhs246_ = globalState
                lhs247_ = globalState
                lhs248_ = globalState
                lhs245_.f1 = rhs294_
                lhs246_.f1 = rhs295_
                lhs247_.f7 = rhs296_
                lhs248_.f1 = rhs297_
            d_1643_v85_: _dafny.Map
            d_1643_v85_ = _dafny.Map({p0: (self).f11})
            d_1643_v85_ = d_1643_v85_
        elif True:
            d_1644_v86_: _dafny.Map
            d_1644_v86_ = _dafny.Map({d_1531_v2_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1645_i5_ in range(default__.abs(225))])})
            d_1646_v87_: D14
            d_1646_v87_ = D14_DC30(d_1644_v86_)
            d_1647_v88_: D14
            d_1647_v88_ = D14_DC32(d_1646_v87_)
            d_1648_v89_: D7
            d_1648_v89_ = D7_DC13(((self).f12) - ((self).f12), (0) - ((0) - ((self).f11)), -673)
            index237_ = default__.safeIndex(232, (d_1529_v0_).length(0))
            rhs298_ = d_1530_v1_
            rhs299_ = d_1647_v88_
            rhs300_ = default__.safeModuloInt(p0, len(_dafny.SeqWithoutIsStrInference([(0) - (p0)])))
            rhs301_ = d_1648_v89_
            rhs302_ = default__.safeDivisionInt((self).f11, p0)
            lhs249_ = d_1529_v0_
            lhs250_ = default__.safeIndex(232, (d_1529_v0_).length(0))
            lhs251_ = globalState
            lhs252_ = globalState
            lhs249_[lhs250_] = rhs298_
            d_1647_v88_ = rhs299_
            lhs251_.f1 = rhs300_
            d_1648_v89_ = rhs301_
            lhs252_.f1 = rhs302_
            (globalState).f1 = (self).f11
            d_1649_v90_: D15
            d_1649_v90_ = D15_DC36(p0, d_1531_v2_, d_1531_v2_, d_1531_v2_, (self).f11)
            if (d_1649_v90_).cf56:
                (globalState).f6 = (d_1531_v2_) or (d_1531_v2_)
                (globalState).f1 = 841
                d_1650_v91_: _dafny.MultiSet
                d_1650_v91_ = _dafny.MultiSet([p0])
                d_1651_v92_: T1
                nw276_ = C4()
                nw276_.ctor__(d_1650_v91_, d_1531_v2_, (self).f12, (self).f11)
                d_1651_v92_ = nw276_
                d_1652_v93_: _dafny.Map
                d_1652_v93_ = _dafny.Map({not(d_1531_v2_): d_1651_v92_})
                d_1653_v94_: _dafny.Seq
                d_1653_v94_ = _dafny.SeqWithoutIsStrInference([d_1651_v92_])
                d_1654_v95_: _dafny.Array
                nw277_ = _dafny.Array(None, 26)
                nw277_[int(0)] = d_1651_v92_
                nw277_[int(1)] = d_1651_v92_
                nw277_[int(2)] = d_1651_v92_
                nw277_[int(3)] = d_1651_v92_
                nw277_[int(4)] = d_1651_v92_
                nw277_[int(5)] = d_1651_v92_
                nw277_[int(6)] = d_1651_v92_
                nw277_[int(7)] = d_1651_v92_
                nw277_[int(8)] = d_1651_v92_
                nw277_[int(9)] = d_1651_v92_
                nw277_[int(10)] = d_1651_v92_
                nw277_[int(11)] = d_1651_v92_
                nw277_[int(12)] = d_1651_v92_
                nw277_[int(13)] = d_1651_v92_
                nw277_[int(14)] = d_1651_v92_
                nw277_[int(15)] = d_1651_v92_
                nw277_[int(16)] = d_1651_v92_
                nw277_[int(17)] = d_1651_v92_
                nw277_[int(18)] = d_1651_v92_
                nw277_[int(19)] = d_1651_v92_
                nw277_[int(20)] = ((d_1652_v93_)[d_1531_v2_] if (d_1531_v2_) in (d_1652_v93_) else d_1651_v92_)
                nw277_[int(21)] = d_1651_v92_
                nw277_[int(22)] = d_1651_v92_
                nw277_[int(23)] = d_1651_v92_
                nw277_[int(24)] = (d_1653_v94_)[default__.safeIndex((D0_DC1(d_1531_v2_, (d_1651_v92_).f11, d_1531_v2_)).cf2, len(d_1653_v94_))]
                nw277_[int(25)] = d_1651_v92_
                d_1654_v95_ = nw277_
                d_1654_v95_ = d_1654_v95_
                d_1655_v96_: _dafny.Seq
                d_1655_v96_ = _dafny.SeqWithoutIsStrInference([p0, (d_1651_v92_).f11])
                d_1656_v97_: D6
                d_1656_v97_ = D6_DC11(len(d_1655_v96_), (self).f11, p0)
                d_1655_v96_ = (d_1655_v96_) + (default__.fm34(d_1531_v2_, d_1656_v97_, globalState))
                (globalState).f1 = p0
            elif True:
                (globalState).f6 = (d_1531_v2_) and (default__.fm0(globalState))
                d_1657_v98_: D9
                d_1657_v98_ = D9_DC18(_dafny.Map({d_1531_v2_: p0}))
                d_1658_v99_: _dafny.Seq
                d_1658_v99_ = _dafny.SeqWithoutIsStrInference([d_1531_v2_])
                d_1659_v100_: _dafny.Map
                d_1659_v100_ = _dafny.Map({d_1531_v2_: len(d_1658_v99_)})
                pat_let_tv45_ = d_1659_v100_
                def iife140_(_pat_let37_0):
                    def iife141_(d_1660_dt__update__tmp_h2_):
                        def iife142_(_pat_let38_0):
                            def iife143_(d_1661_dt__update_hcf26_h0_):
                                return D9_DC18(d_1661_dt__update_hcf26_h0_)
                            return iife143_(_pat_let38_0)
                        return iife142_(pat_let_tv45_)
                    return iife141_(_pat_let37_0)
                (globalState).f1 = default__.safeModuloInt(len(_dafny.Map({(self).f11: d_1531_v2_})), len((iife140_(d_1657_v98_)).cf26))
                d_1659_v100_ = d_1659_v100_
                d_1662_v101_: str
                d_1662_v101_ = _dafny.CodePoint('r')
                d_1662_v101_ = d_1662_v101_
                (globalState).f1 = (self).f11
            d_1663_v102_: C1
            nw278_ = C1()
            nw278_.ctor__(d_1531_v2_, not (d_1531_v2_) or (d_1531_v2_))
            d_1663_v102_ = nw278_
            d_1663_v102_ = d_1663_v102_
            d_1664_v103_: str
            d_1664_v103_ = _dafny.CodePoint('s')
            d_1665_v104_: _dafny.Set
            d_1665_v104_ = _dafny.Set({d_1664_v103_, d_1664_v103_, d_1664_v103_, d_1664_v103_, d_1664_v103_})
            d_1666_v105_: _dafny.Map
            d_1666_v105_ = _dafny.Map({len(d_1665_v104_): (d_1663_v102_).f18})
            d_1667_v106_: _dafny.Map
            d_1667_v106_ = _dafny.Map({len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))]): ((d_1666_v105_)[(self).f11] if ((self).f11) in (d_1666_v105_) else d_1531_v2_)})
            d_1668_v107_: _dafny.Map
            d_1668_v107_ = _dafny.Map({d_1531_v2_: (_dafny.MultiSet([d_1531_v2_, d_1531_v2_])).cardinality})
            d_1669_v108_: _dafny.MultiSet
            d_1669_v108_ = _dafny.MultiSet([len(d_1530_v1_)])
            d_1670_v109_: _dafny.Array
            nw279_ = _dafny.Array(None, 16)
            nw279_[int(0)] = False
            nw279_[int(1)] = not (d_1531_v2_) or ((d_1663_v102_).f18)
            nw279_[int(2)] = d_1531_v2_
            nw279_[int(3)] = ((d_1663_v102_).f18) and ((d_1663_v102_).f18)
            nw279_[int(4)] = ((d_1666_v105_)[len(_dafny.SeqWithoutIsStrInference([d_1664_v103_, d_1664_v103_]))] if (len(_dafny.SeqWithoutIsStrInference([d_1664_v103_, d_1664_v103_]))) in (d_1666_v105_) else (d_1663_v102_).f19)
            nw279_[int(5)] = not(not((d_1663_v102_).f19))
            nw279_[int(6)] = ((_dafny.MultiSet([(d_1663_v102_).f18, False])).cardinality) != ((self).f11)
            nw279_[int(7)] = not ((d_1663_v102_).f19) or ((d_1663_v102_).f19)
            nw279_[int(8)] = (d_1663_v102_).f19
            nw279_[int(9)] = (p0) >= ((self).f11)
            nw279_[int(10)] = (d_1663_v102_).f19
            nw279_[int(11)] = (d_1663_v102_).f19
            nw279_[int(12)] = (p0) == (len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))]))
            nw279_[int(13)] = (len(d_1668_v107_)) <= (-68)
            nw279_[int(14)] = (d_1669_v108_).ispropersubset(d_1669_v108_)
            nw279_[int(15)] = not ((d_1663_v102_).f19) or (d_1531_v2_)
            d_1670_v109_ = nw279_
            d_1671_v110_: D14
            d_1671_v110_ = D14_DC31((d_1663_v102_).f18, (self).f11, 341)
            index238_ = default__.safeIndex(174, (d_1670_v109_).length(0))
            (d_1670_v109_)[index238_] = ((self).f11) < ((d_1671_v110_).cf46)
            index239_ = default__.safeIndex(174, (d_1670_v109_).length(0))
            (d_1670_v109_)[index239_] = (d_1663_v102_).f19
        d_1672_v111_: _dafny.Array
        def lambda64_(d_1673_i6_):
            return (d_1673_i6_) * ((self).f11)

        init34_ = lambda64_
        nw280_ = _dafny.Array(None, 22)
        for i0_34_ in range(nw280_.length(0)):
            nw280_[i0_34_] = init34_(i0_34_)
        d_1672_v111_ = nw280_
        index240_ = default__.safeIndex(102, (d_1672_v111_).length(0))
        (d_1672_v111_)[index240_] = p0
        index241_ = default__.safeIndex(102, (d_1672_v111_).length(0))
        (d_1672_v111_)[index241_] = (self).f11
        if d_1531_v2_:
            d_1674_v112_: str
            d_1674_v112_ = _dafny.CodePoint('k')
            d_1675_v113_: _dafny.Seq
            d_1675_v113_ = _dafny.SeqWithoutIsStrInference([d_1531_v2_])
            d_1676_v114_: _dafny.Set
            d_1676_v114_ = _dafny.Set({(d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]})
            d_1677_v115_: _dafny.Map
            d_1677_v115_ = _dafny.Map({d_1531_v2_: d_1531_v2_})
            d_1678_v116_: _dafny.Map
            d_1678_v116_ = _dafny.Map({d_1676_v114_: d_1677_v115_})
            d_1679_v117_: _dafny.MultiSet
            d_1679_v117_ = _dafny.MultiSet([d_1674_v112_, default__.fm20(d_1675_v113_, (0) - ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]), d_1531_v2_, globalState), default__.fm20(d_1675_v113_, len(d_1678_v116_), d_1531_v2_, globalState)])
            (globalState).f1 = (((d_1679_v117_)[d_1674_v112_] if (d_1674_v112_) in (d_1679_v117_) else (self).f11)) * (len(_dafny.SeqWithoutIsStrInference([(d_1675_v113_)[default__.safeIndex((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))], len(d_1675_v113_))], d_1531_v2_])))
            d_1680_v118_: _dafny.Map
            d_1680_v118_ = _dafny.Map({d_1531_v2_: (((self).f15).cf8) * (len(d_1530_v1_))})
            d_1680_v118_ = (d_1680_v118_).set(((self).f11) < (102), (d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))])
            d_1530_v1_ = (d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))]
            (globalState).f1 = (self).f11
            d_1681_v119_: D0
            d_1681_v119_ = D0_DC0((p0) * ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]))
            source25_ = d_1681_v119_
            if source25_.is_DC1:
                d_1682___mcc_h7_ = source25_.cf1
                d_1683___mcc_h8_ = source25_.cf2
                d_1684___mcc_h9_ = source25_.cf3
                d_1685_cf3_ = d_1684___mcc_h9_
                d_1686_cf2_ = d_1683___mcc_h8_
                d_1687_cf1_ = d_1682___mcc_h7_
                d_1688_v120_: _dafny.Map
                d_1688_v120_ = _dafny.Map({(d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]: True})
                (globalState).f1 = (len((d_1688_v120_) | (d_1688_v120_))) - ((p0) * ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]))
                d_1689_v121_: _dafny.MultiSet
                d_1689_v121_ = _dafny.MultiSet([d_1685_cf3_])
                d_1690_v122_: _dafny.Map
                d_1690_v122_ = _dafny.Map({(d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]: (self).f11})
                d_1691_v123_: _dafny.Map
                d_1691_v123_ = _dafny.Map({_dafny.MultiSet([d_1686_cf2_]): d_1690_v122_})
                d_1531_v2_ = ((_dafny.MultiSet([(self).fm4((d_1689_v121_).cardinality, d_1691_v123_, p0, d_1687_cf1_, globalState), d_1687_cf1_])) - (_dafny.MultiSet([d_1685_cf3_]))).ispropersubset(d_1689_v121_)
                d_1675_v113_ = (((d_1675_v113_) + (d_1675_v113_)) + (d_1675_v113_)).set(default__.safeIndex(default__.safeModuloInt(d_1686_cf2_, (_dafny.MultiSet([559])).cardinality), len(((d_1675_v113_) + (d_1675_v113_)) + (d_1675_v113_))), ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]) >= (518))
                d_1692_v124_: _dafny.Map
                d_1692_v124_ = _dafny.Map({d_1674_v112_: d_1531_v2_})
                d_1693_v125_: _dafny.Array
                nw281_ = _dafny.Array(None, 26)
                nw281_[int(0)] = ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]) != (p0)
                nw281_[int(1)] = (default__.fm1(globalState)) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlp")))
                nw281_[int(2)] = (d_1675_v113_)[default__.safeIndex(p0, len(d_1675_v113_))]
                nw281_[int(3)] = not(not (d_1685_cf3_) or (d_1687_cf1_))
                nw281_[int(4)] = ((d_1692_v124_)[d_1674_v112_] if (d_1674_v112_) in (d_1692_v124_) else d_1687_cf1_)
                nw281_[int(5)] = not(True)
                nw281_[int(6)] = (d_1675_v113_)[default__.safeIndex(476, len(d_1675_v113_))]
                nw281_[int(7)] = d_1531_v2_
                nw281_[int(8)] = d_1685_cf3_
                nw281_[int(9)] = d_1685_cf3_
                nw281_[int(10)] = d_1685_cf3_
                nw281_[int(11)] = d_1531_v2_
                nw281_[int(12)] = d_1687_cf1_
                nw281_[int(13)] = d_1685_cf3_
                nw281_[int(14)] = d_1685_cf3_
                nw281_[int(15)] = d_1531_v2_
                nw281_[int(16)] = not((d_1689_v121_).ispropersubset(_dafny.MultiSet([not(d_1531_v2_)])))
                nw281_[int(17)] = d_1531_v2_
                nw281_[int(18)] = d_1687_cf1_
                nw281_[int(19)] = d_1685_cf3_
                nw281_[int(20)] = not(d_1687_cf1_)
                nw281_[int(21)] = d_1531_v2_
                nw281_[int(22)] = (not(d_1687_cf1_)) or (d_1685_cf3_)
                nw281_[int(23)] = d_1685_cf3_
                nw281_[int(24)] = (d_1687_cf1_) or (d_1531_v2_)
                nw281_[int(25)] = d_1531_v2_
                d_1693_v125_ = nw281_
                index242_ = default__.safeIndex(531, (d_1693_v125_).length(0))
                (d_1693_v125_)[index242_] = False
                d_1694_v127_: _dafny.Seq
                d_1694_v127_ = _dafny.SeqWithoutIsStrInference([d_1676_v114_])
                index243_ = default__.safeIndex(531, (d_1693_v125_).length(0))
                index244_ = default__.safeIndex(232, (d_1529_v0_).length(0))
                def iife144_():
                    coll66_ = _dafny.Set()
                    compr_66_: int
                    for compr_66_ in (d_1676_v114_).Elements:
                        d_1695_v126_: int = compr_66_
                        if (d_1695_v126_) in (d_1676_v114_):
                            coll66_ = coll66_.union(_dafny.Set([default__.safeModuloInt(d_1695_v126_, 575)]))
                    return _dafny.Set(coll66_)
                rhs303_ = (iife144_()
                ).isdisjoint((d_1694_v127_)[default__.safeIndex(p0, len(d_1694_v127_))])
                rhs304_ = default__.fm1(globalState)
                rhs305_ = (0) - (default__.fm2(p0, p0, globalState))
                lhs253_ = d_1693_v125_
                lhs254_ = default__.safeIndex(531, (d_1693_v125_).length(0))
                lhs255_ = d_1529_v0_
                lhs256_ = default__.safeIndex(232, (d_1529_v0_).length(0))
                lhs257_ = globalState
                lhs253_[lhs254_] = rhs303_
                lhs255_[lhs256_] = rhs304_
                lhs257_.f1 = rhs305_
            elif True:
                d_1696___mcc_h10_ = source25_.cf0
                d_1697_cf0_ = d_1696___mcc_h10_
                d_1698_v128_: _dafny.Array
                def lambda65_(d_1699_v2_):
                    def lambda66_(d_1700_i7_):
                        return _dafny.MultiSet([not(not(d_1699_v2_)), (D10_DC21(d_1699_v2_, d_1699_v2_)).cf28, not(d_1699_v2_)])

                    return lambda66_

                init35_ = lambda65_(d_1531_v2_)
                nw282_ = _dafny.Array(None, 4)
                for i0_35_ in range(nw282_.length(0)):
                    nw282_[i0_35_] = init35_(i0_35_)
                d_1698_v128_ = nw282_
                rhs306_ = d_1698_v128_
                rhs307_ = _dafny.CodePoint('j')
                rhs308_ = default__.safeDivisionInt((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))], (412 if d_1531_v2_ else p0))
                rhs309_ = (d_1530_v1_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqo")) if d_1531_v2_ else d_1530_v1_))
                lhs258_ = globalState
                lhs259_ = globalState
                d_1698_v128_ = rhs306_
                d_1674_v112_ = rhs307_
                lhs258_.f1 = rhs308_
                lhs259_.f7 = rhs309_
                (globalState).f6 = d_1531_v2_
                d_1701_v129_: _dafny.Array
                nw283_ = _dafny.Array(_dafny.Set({}), 21)
                d_1701_v129_ = nw283_
                rhs310_ = not(d_1531_v2_)
                rhs311_ = d_1701_v129_
                lhs260_ = globalState
                lhs260_.f7 = rhs310_
                d_1701_v129_ = rhs311_
                d_1702_v130_: _dafny.Array
                def lambda67_(d_1703_v2_):
                    def lambda68_(d_1704_i8_):
                        return d_1703_v2_

                    return lambda68_

                init36_ = lambda67_(d_1531_v2_)
                nw284_ = _dafny.Array(None, 2)
                for i0_36_ in range(nw284_.length(0)):
                    nw284_[i0_36_] = init36_(i0_36_)
                d_1702_v130_ = nw284_
                d_1705_v131_: _dafny.Map
                d_1705_v131_ = _dafny.Map({len(_dafny.Map({(self).f11: 812})): (self).f11})
                index245_ = default__.safeIndex(800, (d_1702_v130_).length(0))
                (d_1702_v130_)[index245_] = (d_1705_v131_) == (d_1705_v131_)
                d_1706_v132_: _dafny.Map
                d_1706_v132_ = _dafny.Map({d_1672_v111_: d_1531_v2_})
                d_1707_v133_: D10
                d_1707_v133_ = D10_DC20(d_1706_v132_)
                d_1708_v134_: _dafny.Map
                d_1708_v134_ = _dafny.Map({d_1707_v133_: (p0) - ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))])})
                index246_ = default__.safeIndex(800, (d_1702_v130_).length(0))
                rhs312_ = d_1531_v2_
                rhs313_ = d_1708_v134_
                rhs314_ = d_1672_v111_
                lhs261_ = d_1702_v130_
                lhs262_ = default__.safeIndex(800, (d_1702_v130_).length(0))
                lhs261_[lhs262_] = rhs312_
                d_1708_v134_ = rhs313_
                d_1672_v111_ = rhs314_
        elif True:
            def lambda69_(d_1709_v111_):
                def lambda70_(d_1710_i9_):
                    return (d_1710_i9_) * ((d_1709_v111_)[default__.safeIndex(102, (d_1709_v111_).length(0))])

                return lambda70_

            init37_ = lambda69_(d_1672_v111_)
            nw285_ = _dafny.Array(None, 7)
            for i0_37_ in range(nw285_.length(0)):
                nw285_[i0_37_] = init37_(i0_37_)
            d_1672_v111_ = nw285_
            d_1711_v135_: C5
            nw286_ = C5()
            nw286_.ctor__()
            d_1711_v135_ = nw286_
            d_1712_v136_: _dafny.Array
            nw287_ = _dafny.Array(None, 14)
            nw287_[int(0)] = d_1711_v135_
            nw287_[int(1)] = d_1711_v135_
            nw287_[int(2)] = d_1711_v135_
            nw287_[int(3)] = (D20_DC44(d_1711_v135_)).cf74
            nw287_[int(4)] = d_1711_v135_
            nw287_[int(5)] = d_1711_v135_
            nw287_[int(6)] = d_1711_v135_
            nw287_[int(7)] = d_1711_v135_
            nw287_[int(8)] = d_1711_v135_
            nw287_[int(9)] = d_1711_v135_
            nw287_[int(10)] = d_1711_v135_
            nw287_[int(11)] = d_1711_v135_
            nw287_[int(12)] = d_1711_v135_
            nw287_[int(13)] = d_1711_v135_
            d_1712_v136_ = nw287_
            d_1713_v137_: _dafny.Set
            d_1713_v137_ = _dafny.Set({d_1712_v136_, d_1712_v136_, (d_1712_v136_ if d_1531_v2_ else d_1712_v136_)})
            d_1713_v137_ = ((d_1713_v137_) | (d_1713_v137_)).intersection(_dafny.Set({d_1712_v136_, d_1712_v136_}))
            d_1530_v1_ = (default__.fm1(globalState)) + (d_1530_v1_)
            d_1714_v138_: _dafny.Seq
            d_1714_v138_ = _dafny.SeqWithoutIsStrInference([d_1531_v2_])
            (globalState).f7 = (d_1714_v138_)[default__.safeIndex(287, len(d_1714_v138_))]
            d_1715_v139_: str
            d_1715_v139_ = _dafny.CodePoint('f')
            d_1716_v140_: _dafny.Map
            d_1716_v140_ = _dafny.Map({False: d_1531_v2_})
            d_1717_v141_: _dafny.MultiSet
            d_1717_v141_ = _dafny.MultiSet([d_1716_v140_, d_1716_v140_, d_1716_v140_])
            d_1718_v142_: _dafny.Array
            nw288_ = _dafny.Array(None, 8)
            nw288_[int(0)] = (p0) != (883)
            nw288_[int(1)] = d_1531_v2_
            nw288_[int(2)] = default__.fm0(globalState)
            nw288_[int(3)] = d_1531_v2_
            nw288_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))) == ((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])
            nw288_[int(5)] = d_1531_v2_
            nw288_[int(6)] = (d_1715_v139_) not in ((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])
            nw288_[int(7)] = (d_1717_v141_).isdisjoint(d_1717_v141_)
            d_1718_v142_ = nw288_
            d_1719_v144_: _dafny.MultiSet
            d_1719_v144_ = _dafny.MultiSet([(self).f11, (d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))], p0])
            d_1720_v145_: _dafny.Seq
            d_1720_v145_ = _dafny.SeqWithoutIsStrInference([d_1719_v144_])
            index247_ = default__.safeIndex(687, (d_1718_v142_).length(0))
            def iife145_():
                coll67_ = _dafny.Map()
                compr_67_: _dafny.MultiSet
                for compr_67_ in (d_1720_v145_).Elements:
                    d_1721_v143_: _dafny.MultiSet = compr_67_
                    if (d_1721_v143_) in (d_1720_v145_):
                        coll67_[d_1721_v143_] = _dafny.Map({p0: (0) - (p0)})
                return _dafny.Map(coll67_)
            (d_1718_v142_)[index247_] = (self).fm4((self).f11, iife145_()
            , (self).f11, (d_1714_v138_)[default__.safeIndex(p0, len(d_1714_v138_))], globalState)
            d_1722_v146_: _dafny.Set
            d_1722_v146_ = _dafny.Set({d_1715_v139_})
            index248_ = default__.safeIndex(687, (d_1718_v142_).length(0))
            (d_1718_v142_)[index248_] = (d_1722_v146_).issubset(d_1722_v146_)
        d_1723_v147_: _dafny.MultiSet
        d_1723_v147_ = _dafny.MultiSet([d_1531_v2_, d_1531_v2_, d_1531_v2_])
        d_1724_v148_: _dafny.Set
        d_1724_v148_ = _dafny.Set({len(d_1530_v1_)})
        d_1725_v149_: _dafny.Map
        d_1725_v149_ = _dafny.Map({len(d_1724_v148_): d_1531_v2_})
        d_1726_v150_: D12
        d_1726_v150_ = D12_DC26(d_1725_v149_)
        d_1727_v151_: _dafny.Map
        d_1727_v151_ = _dafny.Map({d_1726_v150_: d_1723_v147_})
        d_1728_v152_: _dafny.Array
        nw289_ = _dafny.Array(None, 7)
        nw289_[int(0)] = d_1723_v147_
        nw289_[int(1)] = d_1723_v147_
        nw289_[int(2)] = d_1723_v147_
        nw289_[int(3)] = _dafny.MultiSet([d_1531_v2_, d_1531_v2_])
        nw289_[int(4)] = d_1723_v147_
        nw289_[int(5)] = d_1723_v147_
        nw289_[int(6)] = ((d_1727_v151_)[D12_DC26(_dafny.Map({(0) - ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]): d_1531_v2_}))] if (D12_DC26(_dafny.Map({(0) - ((d_1672_v111_)[default__.safeIndex(102, (d_1672_v111_).length(0))]): d_1531_v2_}))) in (d_1727_v151_) else d_1723_v147_)
        d_1728_v152_ = nw289_
        r0 = d_1728_v152_
        d_1729_v153_: _dafny.Seq
        d_1729_v153_ = _dafny.SeqWithoutIsStrInference([True, False])
        d_1730_v154_: _dafny.MultiSet
        d_1730_v154_ = _dafny.MultiSet([len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])])
        d_1731_v155_: _dafny.Map
        d_1731_v155_ = _dafny.Map({((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))]).set(default__.safeIndex((self).f11, len((d_1529_v0_)[default__.safeIndex(232, (d_1529_v0_).length(0))])), default__.fm20(d_1729_v153_, len(d_1530_v1_), d_1531_v2_, globalState)): d_1730_v154_})
        r1 = (d_1731_v155_) | ((d_1731_v155_) | (d_1731_v155_))
        return r0, r1

    @property
    def f15(self):
        return self._f15

class C7(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Set = _dafny.Set({})
        self._f11: int = int(0)
        self._f14: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f12(self):
        return self._f12
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f14, f12, f11):
        (self)._f14 = f14
        (self)._f12 = f12
        (self)._f11 = f11

    def fm5(self, p0, p1, p2, p3, globalState):
        return not((D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))), True)).cf1)

    def fm6(self, p0, p1, globalState):
        source26_ = D0_DC0((self).f11)
        if source26_.is_DC1:
            d_1732___mcc_h0_ = source26_.cf1
            d_1733___mcc_h1_ = source26_.cf2
            d_1734___mcc_h2_ = source26_.cf3
            d_1735_cf3_ = d_1734___mcc_h2_
            d_1736_cf2_ = d_1733___mcc_h1_
            d_1737_cf1_ = d_1732___mcc_h0_
            return d_1736_cf2_
        elif True:
            d_1738___mcc_h3_ = source26_.cf0
            d_1739_cf0_ = d_1738___mcc_h3_
            return (self).f11

    def fm3(self, p0, p1, p2, globalState):
        def lambda71_(source27_):
            if source27_.is_DC1:
                d_1740___mcc_h0_ = source27_.cf1
                d_1741___mcc_h1_ = source27_.cf2
                d_1742___mcc_h2_ = source27_.cf3
                d_1743_cf3_ = d_1742___mcc_h2_
                d_1744_cf2_ = d_1741___mcc_h1_
                d_1745_cf1_ = d_1740___mcc_h0_
                return ((self).f11) + ((self).f11)
            elif True:
                d_1746___mcc_h3_ = source27_.cf0
                d_1747_cf0_ = d_1746___mcc_h3_
                return (self).f11

        return (0) - (lambda71_((D0_DC1(not(False), (self).f11, True) if True else D0_DC1(True, (self).f11, False))))

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1748_i0_ in range(default__.abs(97))]))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdsxn")))

    def m0(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_1749_v0_: bool
        d_1749_v0_ = True
        d_1750_v1_: _dafny.MultiSet
        d_1750_v1_ = _dafny.MultiSet([d_1749_v0_, d_1749_v0_, d_1749_v0_, d_1749_v0_])
        d_1751_v2_: _dafny.Map
        d_1751_v2_ = _dafny.Map({d_1749_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lupwwwla"))})
        hi2_ = len(((d_1751_v2_)[d_1749_v0_] if (d_1749_v0_) in (d_1751_v2_) else _dafny.SeqWithoutIsStrInference([p0 for d_1752_i1_ in range(default__.abs(-92))])))
        for d_1753_i0_ in range((d_1750_v1_).cardinality, hi2_):
            d_1754_v3_: _dafny.Seq
            d_1754_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyslj"))
            d_1755_v4_: _dafny.MultiSet
            d_1755_v4_ = _dafny.MultiSet([d_1754_v3_])
            d_1756_v5_: D2
            d_1756_v5_ = D2_DC3(d_1755_v4_)
            d_1757_v6_: _dafny.Seq
            d_1757_v6_ = _dafny.SeqWithoutIsStrInference([d_1749_v0_, d_1749_v0_, not(((d_1756_v5_).cf5).issubset(d_1755_v4_)), (d_1749_v0_) == (False)])
            d_1757_v6_ = (d_1757_v6_).set(default__.safeIndex((self).f11, len(d_1757_v6_)), d_1749_v0_)
            (globalState).f1 = d_1753_i0_
            (globalState).f6 = (_dafny.Set({d_1749_v0_})).issubset((self).f12)
            d_1758_v7_: _dafny.Map
            d_1758_v7_ = _dafny.Map({(self).f11: d_1749_v0_})
            d_1759_v8_: _dafny.Map
            d_1759_v8_ = _dafny.Map({len(((d_1751_v2_)[d_1749_v0_] if (d_1749_v0_) in (d_1751_v2_) else d_1754_v3_)): len(d_1758_v7_)})
            d_1760_v9_: C2
            nw290_ = C2()
            nw290_.ctor__(default__.fm20(d_1757_v6_, (self).f11, d_1749_v0_, globalState), len((d_1759_v8_) | (_dafny.Map({-721: (self).f11}))))
            d_1760_v9_ = nw290_
        d_1761_v10_: _dafny.MultiSet
        d_1761_v10_ = _dafny.MultiSet([(self).f11])
        d_1762_v11_: C4
        nw291_ = C4()
        nw291_.ctor__(d_1761_v10_, d_1749_v0_, ((self).f12).intersection((self).f12), default__.safeModuloInt(-754, (d_1750_v1_).cardinality))
        d_1762_v11_ = nw291_
        hi3_ = (self).f11
        for d_1763_i2_ in range((self).f11, hi3_):
            d_1764_v12_: D9
            d_1764_v12_ = D9_DC19()
            d_1765_v13_: _dafny.Seq
            d_1765_v13_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f11), (self).f11])
            d_1766_v14_: _dafny.Array
            nw292_ = _dafny.Array(None, 15)
            nw292_[int(0)] = d_1764_v12_
            nw292_[int(1)] = d_1764_v12_
            nw292_[int(2)] = (d_1764_v12_ if (d_1762_v11_).f17 else d_1764_v12_)
            nw292_[int(3)] = d_1764_v12_
            nw292_[int(4)] = D9_DC19()
            nw292_[int(5)] = d_1764_v12_
            nw292_[int(6)] = D9_DC19()
            nw292_[int(7)] = d_1764_v12_
            nw292_[int(8)] = d_1764_v12_
            nw292_[int(9)] = d_1764_v12_
            nw292_[int(10)] = d_1764_v12_
            nw292_[int(11)] = default__.fm42(_dafny.Set({d_1765_v13_}), globalState)
            nw292_[int(12)] = D9_DC19()
            nw292_[int(13)] = (d_1764_v12_ if d_1749_v0_ else d_1764_v12_)
            nw292_[int(14)] = d_1764_v12_
            d_1766_v14_ = nw292_
            index249_ = default__.safeIndex(144, (d_1766_v14_).length(0))
            (d_1766_v14_)[index249_] = d_1764_v12_
            d_1767_v15_: D8
            d_1767_v15_ = D8_DC16((d_1762_v11_).f17, d_1763_i2_)
            d_1768_v16_: D14
            d_1768_v16_ = D14_DC31(d_1749_v0_, (0) - (d_1763_i2_), (self).f11)
            d_1769_v17_: _dafny.Seq
            d_1769_v17_ = _dafny.SeqWithoutIsStrInference([(d_1762_v11_).f17, d_1749_v0_, d_1749_v0_])
            d_1770_v18_: _dafny.Array
            nw293_ = _dafny.Array(None, 26)
            nw293_[int(0)] = (537) <= (d_1763_i2_)
            nw293_[int(1)] = (False if (d_1762_v11_).f17 else (d_1762_v11_).f17)
            nw293_[int(2)] = True
            nw293_[int(3)] = (d_1762_v11_).f17
            nw293_[int(4)] = (d_1767_v15_).cf23
            nw293_[int(5)] = ((self).f11) < ((self).f11)
            nw293_[int(6)] = (d_1762_v11_).f17
            nw293_[int(7)] = (d_1762_v11_).f17
            nw293_[int(8)] = (d_1762_v11_).f17
            nw293_[int(9)] = default__.fm0(globalState)
            nw293_[int(10)] = (d_1763_i2_) > (861)
            nw293_[int(11)] = ((self).f11) < ((self).f11)
            nw293_[int(12)] = False
            nw293_[int(13)] = (d_1763_i2_) >= (d_1763_i2_)
            nw293_[int(14)] = (d_1762_v11_).f17
            nw293_[int(15)] = (d_1768_v16_).cf45
            nw293_[int(16)] = (len(_dafny.Set({(d_1762_v11_).f17, (d_1762_v11_).f17}))) != (d_1763_i2_)
            nw293_[int(17)] = (d_1762_v11_).f17
            nw293_[int(18)] = d_1749_v0_
            nw293_[int(19)] = d_1749_v0_
            nw293_[int(20)] = (d_1762_v11_).f17
            nw293_[int(21)] = (d_1762_v11_).f17
            nw293_[int(22)] = (d_1769_v17_)[default__.safeIndex((d_1765_v13_)[default__.safeIndex((0) - (d_1763_i2_), len(d_1765_v13_))], len(d_1769_v17_))]
            nw293_[int(23)] = (d_1762_v11_).f17
            nw293_[int(24)] = (d_1762_v11_).f17
            nw293_[int(25)] = (d_1762_v11_).f17
            d_1770_v18_ = nw293_
            d_1771_v19_: _dafny.Seq
            d_1771_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdckwdveg"))
            d_1772_v20_: _dafny.Seq
            d_1772_v20_ = d_1771_v19_
            d_1773_v21_: _dafny.Map
            d_1773_v21_ = _dafny.Map({d_1772_v20_: (d_1762_v11_).f17})
            index250_ = default__.safeIndex(665, (d_1770_v18_).length(0))
            (d_1770_v18_)[index250_] = ((d_1762_v11_).f17 if ((d_1773_v21_)[d_1772_v20_] if (d_1772_v20_) in (d_1773_v21_) else (d_1762_v11_).f17) else False)
            d_1774_v22_: _dafny.Map
            d_1774_v22_ = _dafny.Map({(d_1762_v11_).f17: (d_1762_v11_).f17})
            index251_ = default__.safeIndex(144, (d_1766_v14_).length(0))
            index252_ = default__.safeIndex(665, (d_1770_v18_).length(0))
            rhs315_ = d_1764_v12_
            rhs316_ = (d_1763_i2_) != ((d_1762_v11_).fm6((d_1762_v11_).f17, default__.fm2(d_1763_i2_, d_1763_i2_, globalState), globalState))
            rhs317_ = (d_1769_v17_)[default__.safeIndex(default__.safeDivisionInt(len(d_1774_v22_), d_1763_i2_), len(d_1769_v17_))]
            rhs318_ = (d_1762_v11_).f17
            lhs263_ = d_1766_v14_
            lhs264_ = default__.safeIndex(144, (d_1766_v14_).length(0))
            lhs265_ = globalState
            lhs266_ = d_1770_v18_
            lhs267_ = default__.safeIndex(665, (d_1770_v18_).length(0))
            lhs268_ = globalState
            lhs263_[lhs264_] = rhs315_
            lhs265_.f6 = rhs316_
            lhs266_[lhs267_] = rhs317_
            lhs268_.f6 = rhs318_
            d_1775_v23_: D16
            d_1775_v23_ = D16_DC39((self).f11, (d_1770_v18_)[default__.safeIndex(665, (d_1770_v18_).length(0))], d_1774_v22_, p0)
            d_1776_v24_: _dafny.Map
            d_1776_v24_ = _dafny.Map({d_1762_v11_: d_1769_v17_})
            d_1777_v25_: T0
            nw294_ = C3()
            nw294_.ctor__(default__.safeDivisionInt((d_1775_v23_).cf62, (self).f11), (0) - ((self).f11), len(d_1776_v24_), ((self).f12) - ((self).f12))
            d_1777_v25_ = nw294_
            index253_ = default__.safeIndex(665, (d_1770_v18_).length(0))
            rhs319_ = not((d_1762_v11_).f17)
            rhs320_ = (d_1762_v11_).f17
            rhs321_ = d_1777_v25_
            rhs322_ = (d_1777_v25_).f11
            lhs269_ = d_1770_v18_
            lhs270_ = default__.safeIndex(665, (d_1770_v18_).length(0))
            lhs271_ = globalState
            d_1749_v0_ = rhs319_
            lhs269_[lhs270_] = rhs320_
            d_1777_v25_ = rhs321_
            lhs271_.f1 = rhs322_
            (globalState).f6 = (d_1771_v19_) != (((d_1771_v19_).set(default__.safeIndex(d_1763_i2_, len(d_1771_v19_)), p0) if not((d_1770_v18_)[default__.safeIndex(665, (d_1770_v18_).length(0))]) else d_1771_v19_))
            d_1778_v26_: D7
            d_1778_v26_ = D7_DC13((self).f12, d_1763_i2_, (len(d_1771_v19_)) - ((self).f11))
            d_1779_v27_: _dafny.Set
            d_1779_v27_ = _dafny.Set({(d_1777_v25_).f11})
            d_1780_v28_: _dafny.Map
            d_1780_v28_ = _dafny.Map({p0: len((d_1771_v19_).set(default__.safeIndex(d_1763_i2_, len(d_1771_v19_)), p0))})
            pat_let_tv46_ = d_1780_v28_
            def iife146_(_pat_let39_0):
                def iife147_(d_1781_dt__update__tmp_h0_):
                    def iife148_(_pat_let40_0):
                        def iife149_(d_1782_dt__update_hcf18_h0_):
                            def iife150_(_pat_let41_0):
                                def iife151_(d_1783_dt__update_hcf20_h0_):
                                    return D7_DC13(d_1782_dt__update_hcf18_h0_, (d_1781_dt__update__tmp_h0_).cf19, d_1783_dt__update_hcf20_h0_)
                                return iife151_(_pat_let41_0)
                            return iife150_(len(pat_let_tv46_))
                        return iife149_(_pat_let40_0)
                    return iife148_((self).f12)
                return iife147_(_pat_let39_0)
            d_1778_v26_ = (default__.fm36((d_1770_v18_)[default__.safeIndex(665, (d_1770_v18_).length(0))], default__.fm0(globalState), (d_1777_v25_).f11, globalState) if (d_1779_v27_).issubset(d_1779_v27_) else iife146_(d_1778_v26_))
        (globalState).f1 = (0) - ((self).f11)
        d_1784_v29_: _dafny.Seq
        d_1784_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sctxpfjb"))
        hi4_ = (-812) - (len(d_1784_v29_))
        for d_1785_i3_ in range((0) - ((self).f11), hi4_):
            d_1786_v30_: _dafny.Array
            d_1787_v31_: _dafny.Map
            out31_: _dafny.Array
            out32_: _dafny.Map
            out31_, out32_ = (d_1762_v11_).m1(d_1785_i3_, globalState)
            d_1786_v30_ = out31_
            d_1787_v31_ = out32_
            d_1788_v32_: _dafny.Seq
            d_1788_v32_ = _dafny.SeqWithoutIsStrInference([(d_1762_v11_).f17, (d_1762_v11_).f17])
            d_1789_v33_: C1
            nw295_ = C1()
            nw295_.ctor__((d_1762_v11_).f17, (self).fm5(D0_DC0((self).f11), d_1788_v32_, d_1785_i3_, p0, globalState))
            d_1789_v33_ = nw295_
            if (d_1789_v33_).f18:
                d_1790_v34_: _dafny.Map
                d_1790_v34_ = _dafny.Map({d_1785_i3_: p0})
                d_1791_v35_: _dafny.Seq
                d_1791_v35_ = _dafny.SeqWithoutIsStrInference([d_1785_i3_])
                d_1792_v36_: _dafny.Seq
                d_1792_v36_ = _dafny.SeqWithoutIsStrInference([len(d_1790_v34_), (d_1785_i3_) - ((self).f11), default__.safeDivisionInt(len(d_1791_v35_), d_1785_i3_)])
                d_1792_v36_ = _dafny.SeqWithoutIsStrInference([d_1785_i3_ for d_1793_i4_ in range(default__.abs(534))])
                d_1794_v37_: _dafny.Array
                def lambda72_(d_1795_i5_):
                    return default__.safeModuloInt(d_1795_i5_, (self).f11)

                init38_ = lambda72_
                nw296_ = _dafny.Array(None, 24)
                for i0_38_ in range(nw296_.length(0)):
                    nw296_[i0_38_] = init38_(i0_38_)
                d_1794_v37_ = nw296_
                index254_ = default__.safeIndex(935, (d_1794_v37_).length(0))
                (d_1794_v37_)[index254_] = ((self).f11) * ((self).f11)
                index255_ = default__.safeIndex(935, (d_1794_v37_).length(0))
                (d_1794_v37_)[index255_] = default__.safeDivisionInt((self).f11, (0) - (len(_dafny.Map({d_1785_i3_: (d_1784_v29_).set(default__.safeIndex(d_1785_i3_, len(d_1784_v29_)), p0)}))))
                d_1796_v38_: _dafny.Set
                d_1796_v38_ = _dafny.Set({(self).fm6((d_1789_v33_).f19, (((d_1762_v11_).f16)[len(d_1788_v32_)] if (len(d_1788_v32_)) in ((d_1762_v11_).f16) else (d_1794_v37_)[default__.safeIndex(935, (d_1794_v37_).length(0))]), globalState), 532, d_1785_i3_})
                d_1797_v39_: _dafny.Map
                d_1797_v39_ = _dafny.Map({(self).f11: (d_1796_v38_) | (d_1796_v38_)})
                d_1796_v38_ = ((d_1797_v39_)[(0) - ((d_1794_v37_)[default__.safeIndex(935, (d_1794_v37_).length(0))])] if ((0) - ((d_1794_v37_)[default__.safeIndex(935, (d_1794_v37_).length(0))])) in (d_1797_v39_) else _dafny.Set({len(default__.fm1(globalState)), (self).f11}))
                pat_let_tv47_ = d_1789_v33_
                pat_let_tv48_ = d_1789_v33_
                pat_let_tv49_ = d_1762_v11_
                def iife152_(_pat_let42_0):
                    def iife153_(d_1798_dt__update__tmp_h1_):
                        def iife154_(_pat_let43_0):
                            def iife155_(d_1799_dt__update_hcf13_h0_):
                                return D6_DC10(d_1799_dt__update_hcf13_h0_)
                            return iife155_(_pat_let43_0)
                        return iife154_(_dafny.SeqWithoutIsStrInference([(pat_let_tv47_).f18, (pat_let_tv48_).f18, (pat_let_tv49_).f17]))
                    return iife153_(_pat_let42_0)
                (globalState).f7 = (_dafny.SeqWithoutIsStrInference([(d_1762_v11_).f17, (D19_DC43((d_1789_v33_).f18, p0, (d_1762_v11_).f17, iife152_(D6_DC10(_dafny.SeqWithoutIsStrInference([(d_1762_v11_).f17]))), d_1784_v29_)).cf69, d_1749_v0_, (d_1788_v32_)[default__.safeIndex(((d_1750_v1_)[(d_1762_v11_).f17] if ((d_1762_v11_).f17) in (d_1750_v1_) else (self).f11), len(d_1788_v32_))]])) != (d_1788_v32_)
                d_1800_v41_: _dafny.Map
                def iife156_():
                    coll68_ = _dafny.Set()
                    compr_68_: str
                    for compr_68_ in (_dafny.SeqWithoutIsStrInference([p0, p0])).Elements:
                        d_1801_v40_: str = compr_68_
                        if (d_1801_v40_) in (_dafny.SeqWithoutIsStrInference([p0, p0])):
                            coll68_ = coll68_.union(_dafny.Set([d_1801_v40_]))
                    return _dafny.Set(coll68_)
                d_1800_v41_ = _dafny.Map({d_1785_i3_: len(iife156_()
                )})
                d_1802_v42_: C3
                nw297_ = C3()
                nw297_.ctor__((self).f11, d_1785_i3_, d_1785_i3_, _dafny.Set({(d_1789_v33_).f19}))
                d_1802_v42_ = nw297_
                d_1803_v43_: _dafny.Map
                d_1803_v43_ = _dafny.Map({d_1800_v41_: d_1802_v42_})
                d_1804_v44_: _dafny.Set
                d_1804_v44_ = _dafny.Set({((d_1803_v43_)[d_1800_v41_] if (d_1800_v41_) in (d_1803_v43_) else d_1802_v42_), d_1802_v42_, d_1802_v42_})
                (globalState).f1 = (d_1785_i3_) - ((0) - (len(d_1804_v44_)))
            elif True:
                d_1805_v45_: _dafny.Set
                d_1805_v45_ = _dafny.Set({(d_1762_v11_).f17, not((d_1789_v33_).fm14((self).f14, globalState))})
                d_1806_v46_: _dafny.Map
                d_1806_v46_ = _dafny.Map({D0_DC0(d_1785_i3_): d_1805_v45_})
                d_1807_v47_: D0
                d_1807_v47_ = D0_DC0((self).f11)
                rhs323_ = ((self).f11) - ((self).f11)
                rhs324_ = d_1749_v0_
                rhs325_ = d_1785_i3_
                rhs326_ = (((d_1806_v46_)[d_1807_v47_] if (d_1807_v47_) in (d_1806_v46_) else default__.fm41(globalState)) if (d_1789_v33_).f19 else d_1805_v45_)
                rhs327_ = default__.safeModuloInt((self).f11, len((d_1784_v29_) + (default__.fm1(globalState))))
                lhs272_ = globalState
                lhs273_ = globalState
                lhs274_ = globalState
                lhs275_ = globalState
                lhs272_.f1 = rhs323_
                lhs273_.f7 = rhs324_
                lhs274_.f1 = rhs325_
                d_1805_v45_ = rhs326_
                lhs275_.f1 = rhs327_
                (globalState).f1 = (self).f11
                d_1808_v48_: _dafny.Set
                d_1808_v48_ = _dafny.Set({p0, default__.fm20(d_1788_v32_, (d_1762_v11_).fm3(d_1785_i3_, (d_1762_v11_).f17, (d_1789_v33_).f19, globalState), (d_1762_v11_).f17, globalState), p0, p0})
                (globalState).f1 = ((self).f11) + (((self).fm6((d_1789_v33_).f18, (self).f11, globalState)) * (len(d_1808_v48_)))
                (globalState).f7 = ((0) - ((self).f11)) > ((self).f11)
                rhs328_ = default__.safeDivisionInt(((self).f11) - ((0) - ((self).f11)), -154)
                rhs329_ = default__.fm0(globalState)
                rhs330_ = (d_1789_v33_).fm14((self).f14, globalState)
                rhs331_ = (d_1784_v29_) + (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stifwiti"))) + (_dafny.SeqWithoutIsStrInference([p0 for d_1809_i6_ in range(default__.abs(265))]))).set(default__.safeIndex(d_1785_i3_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stifwiti"))) + (_dafny.SeqWithoutIsStrInference([p0 for d_1810_i6_ in range(default__.abs(265))])))), p0))
                lhs276_ = globalState
                lhs277_ = globalState
                lhs278_ = globalState
                lhs276_.f1 = rhs328_
                lhs277_.f6 = rhs329_
                lhs278_.f6 = rhs330_
                d_1784_v29_ = rhs331_
            d_1811_v49_: _dafny.Array
            nw298_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_1811_v49_ = nw298_
            index256_ = default__.safeIndex(740, (d_1811_v49_).length(0))
            (d_1811_v49_)[index256_] = d_1786_v30_
            index257_ = default__.safeIndex(740, (d_1811_v49_).length(0))
            (d_1811_v49_)[index257_] = d_1786_v30_
        d_1812_v50_: _dafny.Array
        nw299_ = _dafny.Array(int(0), 21)
        d_1812_v50_ = nw299_
        d_1812_v50_ = d_1812_v50_
        r0 = default__.fm37(globalState)
        return r0

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        d_1813_v0_: bool
        d_1813_v0_ = False
        d_1814_v1_: str
        d_1814_v1_ = _dafny.CodePoint('c')
        d_1815_v2_: _dafny.Set
        d_1815_v2_ = _dafny.Set({d_1814_v1_, d_1814_v1_, d_1814_v1_})
        d_1816_v3_: _dafny.MultiSet
        d_1816_v3_ = _dafny.MultiSet([(self).f12, (self).f12])
        d_1817_i0_: int
        d_1817_i0_ = 0
        with _dafny.label("16"):
            while (default__.fm53((self).f11, d_1813_v0_, (0) - (len(d_1815_v2_)), globalState)) == (d_1816_v3_):
                with _dafny.c_label("16"):
                    if (d_1817_i0_) >= (100):
                        raise _dafny.Break("16")
                    d_1817_i0_ = (d_1817_i0_) + (1)
                    d_1818_v4_: _dafny.Array
                    nw300_ = _dafny.Array(False, 18)
                    d_1818_v4_ = nw300_
                    index258_ = default__.safeIndex(437, (d_1818_v4_).length(0))
                    (d_1818_v4_)[index258_] = d_1813_v0_
                    index259_ = default__.safeIndex(437, (d_1818_v4_).length(0))
                    (d_1818_v4_)[index259_] = d_1813_v0_
                    d_1819_v5_: _dafny.Map
                    d_1819_v5_ = _dafny.Map({619: d_1813_v0_})
                    d_1820_v6_: _dafny.Map
                    d_1820_v6_ = _dafny.Map({p0: (self).f11})
                    (globalState).f7 = (_dafny.Map({(self).f11: len(d_1819_v5_)})) != (((d_1820_v6_).set(p0, -967)).set((self).f11, p0))
                    d_1821_v7_: D0
                    d_1821_v7_ = D0_DC0(p0)
                    d_1822_v8_: _dafny.Map
                    d_1822_v8_ = _dafny.Map({d_1813_v0_: (d_1818_v4_)[default__.safeIndex(437, (d_1818_v4_).length(0))]})
                    d_1823_v9_: D16
                    d_1823_v9_ = D16_DC39(p0, d_1813_v0_, d_1822_v8_, _dafny.CodePoint('m'))
                    d_1824_v10_: _dafny.Seq
                    d_1824_v10_ = _dafny.SeqWithoutIsStrInference([(d_1823_v9_).cf63])
                    index260_ = default__.safeIndex(437, (d_1818_v4_).length(0))
                    (d_1818_v4_)[index260_] = (self).fm5(d_1821_v7_, (d_1824_v10_) + (d_1824_v10_), p0, d_1814_v1_, globalState)
                    if (d_1818_v4_)[default__.safeIndex(437, (d_1818_v4_).length(0))]:
                        (globalState).f1 = p0
                        d_1825_v11_: _dafny.Array
                        nw301_ = _dafny.Array(int(0), 14)
                        d_1825_v11_ = nw301_
                        nw302_ = _dafny.Array(int(0), 19)
                        d_1825_v11_ = nw302_
                        (globalState).f1 = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gh"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gh")))), (d_1814_v1_ if d_1813_v0_ else d_1814_v1_)))
                        d_1819_v5_ = (d_1819_v5_).set((self).f11, True)
                        d_1826_v12_: _dafny.MultiSet
                        d_1826_v12_ = _dafny.MultiSet([d_1813_v0_, (d_1818_v4_)[default__.safeIndex(437, (d_1818_v4_).length(0))]])
                        (globalState).f7 = ((d_1826_v12_).cardinality) < (-682)
                    elif True:
                        (globalState).f1 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psoaf")))
                        d_1827_v13_: _dafny.MultiSet
                        d_1827_v13_ = _dafny.MultiSet([p0, 331])
                        d_1828_v14_: _dafny.Seq
                        d_1828_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acp"))
                        d_1829_v15_: C4
                        nw303_ = C4()
                        nw303_.ctor__(d_1827_v13_, (d_1828_v14_) != (d_1828_v14_), (self).f12, len(_dafny.Map({p0: d_1814_v1_})))
                        d_1829_v15_ = nw303_
                        (globalState).f6 = ((-207) + (p0)) <= (p0)
                        (globalState).f1 = (0) - (p0)
                        d_1830_v16_: _dafny.Array
                        nw304_ = _dafny.Array(int(0), 3)
                        d_1830_v16_ = nw304_
                        index261_ = default__.safeIndex(330, (d_1830_v16_).length(0))
                        (d_1830_v16_)[index261_] = (0) - (p0)
                        index262_ = default__.safeIndex(330, (d_1830_v16_).length(0))
                        (d_1830_v16_)[index262_] = (p0) - (len(_dafny.SeqWithoutIsStrInference([(self).f11])))
                    pass
            pass
        d_1831_v17_: _dafny.MultiSet
        d_1831_v17_ = _dafny.MultiSet([p0])
        d_1832_v18_: T0
        nw305_ = C4()
        nw305_.ctor__(d_1831_v17_, d_1813_v0_, _dafny.Set({d_1813_v0_}), len(_dafny.SeqWithoutIsStrInference([d_1813_v0_, False, d_1813_v0_, d_1813_v0_, d_1813_v0_])))
        d_1832_v18_ = nw305_
        d_1833_v19_: _dafny.Set
        d_1833_v19_ = _dafny.Set({p0, len(_dafny.Set({d_1832_v18_}))})
        d_1834_v20_: _dafny.Map
        d_1834_v20_ = _dafny.Map({d_1833_v19_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eo")))})
        d_1835_v21_: _dafny.Map
        d_1835_v21_ = _dafny.Map({len(d_1834_v20_): default__.fm1(globalState)})
        d_1836_v22_: _dafny.Seq
        d_1836_v22_ = _dafny.SeqWithoutIsStrInference([d_1814_v1_])
        d_1837_v23_: D16
        d_1837_v23_ = D16_DC39((self).f11, d_1813_v0_, _dafny.Map({d_1813_v0_: False}), d_1814_v1_)
        d_1838_v24_: _dafny.Map
        d_1838_v24_ = _dafny.Map({d_1813_v0_: d_1813_v0_})
        d_1839_v25_: _dafny.Seq
        d_1839_v25_ = _dafny.SeqWithoutIsStrInference([d_1838_v24_, d_1838_v24_, d_1838_v24_, _dafny.Map({True: d_1813_v0_})])
        d_1840_v26_: _dafny.Seq
        d_1840_v26_ = _dafny.SeqWithoutIsStrInference([d_1813_v0_])
        d_1841_v27_: _dafny.Map
        d_1841_v27_ = _dafny.Map({d_1813_v0_: 555})
        d_1842_v28_: _dafny.Array
        nw306_ = _dafny.Array(None, 19)
        nw306_[int(0)] = not(d_1813_v0_)
        nw306_[int(1)] = d_1813_v0_
        nw306_[int(2)] = d_1813_v0_
        nw306_[int(3)] = ((self).f11) >= ((self).f11)
        nw306_[int(4)] = ((self).f11) not in (d_1835_v21_)
        nw306_[int(5)] = d_1813_v0_
        nw306_[int(6)] = d_1813_v0_
        nw306_[int(7)] = (d_1836_v22_) < (_dafny.SeqWithoutIsStrInference([(d_1837_v23_).cf65]))
        nw306_[int(8)] = d_1813_v0_
        nw306_[int(9)] = d_1813_v0_
        nw306_[int(10)] = ((d_1832_v18_).f11) < (len(d_1839_v25_))
        nw306_[int(11)] = d_1813_v0_
        nw306_[int(12)] = ((d_1838_v24_)[d_1813_v0_] if (d_1813_v0_) in (d_1838_v24_) else d_1813_v0_)
        nw306_[int(13)] = d_1813_v0_
        nw306_[int(14)] = d_1813_v0_
        nw306_[int(15)] = d_1813_v0_
        nw306_[int(16)] = d_1813_v0_
        nw306_[int(17)] = d_1813_v0_
        nw306_[int(18)] = (d_1840_v26_)[default__.safeIndex((D11_DC23(True, (d_1832_v18_).f11, len(d_1841_v27_))).cf32, len(d_1840_v26_))]
        d_1842_v28_ = nw306_
        index263_ = default__.safeIndex(748, (d_1842_v28_).length(0))
        (d_1842_v28_)[index263_] = d_1813_v0_
        index264_ = default__.safeIndex(748, (d_1842_v28_).length(0))
        (d_1842_v28_)[index264_] = d_1813_v0_
        d_1843_v29_: _dafny.Array
        nw307_ = _dafny.Array(int(0), 28)
        d_1843_v29_ = nw307_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1843_v29_).length(0)):
            d_1844_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_1844_i1_)) and ((d_1844_i1_) < ((d_1843_v29_).length(0)))):
                (d_1843_v29_)[(d_1844_i1_)] = default__.safeDivisionInt(d_1844_i1_, (0) - (p0))
        d_1845_v30_: _dafny.Map
        d_1845_v30_ = _dafny.Map({default__.safeModuloInt(p0, p0): (d_1842_v28_)[default__.safeIndex(748, (d_1842_v28_).length(0))]})
        d_1846_v31_: _dafny.Seq
        d_1846_v31_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_1845_v30_ = (d_1845_v30_).set(((d_1832_v18_).f11) + (len(d_1846_v31_)), d_1813_v0_)
        d_1847_v32_: _dafny.MultiSet
        d_1847_v32_ = _dafny.MultiSet([d_1813_v0_, ((d_1832_v18_).f11) < ((d_1832_v18_).f11)])
        d_1847_v32_ = d_1847_v32_
        (globalState).f1 = (self).f11
        d_1848_v33_: _dafny.Array
        nw308_ = _dafny.Array(_dafny.MultiSet({}), 26)
        d_1848_v33_ = nw308_
        r0 = d_1848_v33_
        d_1849_v34_: _dafny.Map
        d_1849_v34_ = _dafny.Map({d_1836_v22_: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1832_v18_).f11 for d_1850_i2_ in range(default__.abs(-102))]))})
        r1 = (d_1849_v34_) | (d_1849_v34_)
        return r0, r1

    @property
    def f14(self):
        return self._f14

class C8(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Set = _dafny.Set({})
        self._f11: int = int(0)
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f12(self):
        return self._f12
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f13, f12, f11):
        (self)._f13 = f13
        (self)._f12 = f12
        (self)._f11 = f11

    def fm5(self, p0, p1, p2, p3, globalState):
        if (self).f13:
            return (self).f13
        elif True:
            return True

    def fm6(self, p0, p1, globalState):
        return (self).f11

    def fm3(self, p0, p1, p2, globalState):
        return ((_dafny.MultiSet([(self).f11, (self).f11])) - ((_dafny.MultiSet([(self).f11])) | (_dafny.MultiSet([(self).f11])))).cardinality

    def fm4(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypsi"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1851_i0_ in range(default__.abs(3))]))

    def fm7(self, p0, p1, p2, globalState):
        return (self).f11

    def fm8(self, p0, globalState):
        return -6

    def m0(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        if ((self).f13) and ((self).f13):
            d_1852_v0_: D0
            d_1852_v0_ = D0_DC0((self).f11)
            d_1853_v1_: _dafny.Array
            nw309_ = _dafny.Array(None, 17)
            nw309_[int(0)] = (self).f13
            nw309_[int(1)] = (self).f13
            nw309_[int(2)] = (self).f13
            nw309_[int(3)] = (self).f13
            nw309_[int(4)] = (self).f13
            nw309_[int(5)] = (self).fm5(d_1852_v0_, _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13]), (self).f11, p0, globalState)
            nw309_[int(6)] = (self).f13
            nw309_[int(7)] = (self).f13
            nw309_[int(8)] = False
            nw309_[int(9)] = (self).f13
            nw309_[int(10)] = (self).f13
            nw309_[int(11)] = (self).f13
            nw309_[int(12)] = (self).f13
            nw309_[int(13)] = (self).f13
            nw309_[int(14)] = (self).f13
            nw309_[int(15)] = (self).f13
            nw309_[int(16)] = (self).f13
            d_1853_v1_ = nw309_
            d_1854_v2_: _dafny.Seq
            d_1854_v2_ = _dafny.SeqWithoutIsStrInference([d_1853_v1_, d_1853_v1_])
            d_1855_v3_: _dafny.Map
            d_1855_v3_ = _dafny.Map({(self).f13: d_1854_v2_})
            d_1856_v4_: _dafny.Seq
            d_1856_v4_ = _dafny.SeqWithoutIsStrInference([d_1854_v2_])
            d_1857_v5_: _dafny.Seq
            d_1857_v5_ = _dafny.SeqWithoutIsStrInference([(self).f13, not((self).f13), (self).f13])
            d_1858_v6_: _dafny.Array
            nw310_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_1858_v6_ = nw310_
            d_1859_v7_: _dafny.Map
            d_1859_v7_ = _dafny.Map({d_1858_v6_: d_1854_v2_})
            d_1860_v8_: _dafny.Array
            nw311_ = _dafny.Array(None, 27)
            nw311_[int(0)] = (d_1854_v2_) + (d_1854_v2_)
            nw311_[int(1)] = d_1854_v2_
            nw311_[int(2)] = d_1854_v2_
            nw311_[int(3)] = d_1854_v2_
            nw311_[int(4)] = d_1854_v2_
            nw311_[int(5)] = d_1854_v2_
            nw311_[int(6)] = d_1854_v2_
            nw311_[int(7)] = d_1854_v2_
            nw311_[int(8)] = d_1854_v2_
            nw311_[int(9)] = (((d_1855_v3_)[(self).f13] if ((self).f13) in (d_1855_v3_) else ((d_1856_v4_)[default__.safeIndex(len(d_1857_v5_), len(d_1856_v4_))]).set(default__.safeIndex(len(_dafny.Set({(self).f13})), len((d_1856_v4_)[default__.safeIndex(len(d_1857_v5_), len(d_1856_v4_))])), d_1853_v1_))).set(default__.safeIndex((self).f11, len(((d_1855_v3_)[(self).f13] if ((self).f13) in (d_1855_v3_) else ((d_1856_v4_)[default__.safeIndex(len(d_1857_v5_), len(d_1856_v4_))]).set(default__.safeIndex(len(_dafny.Set({(self).f13})), len((d_1856_v4_)[default__.safeIndex(len(d_1857_v5_), len(d_1856_v4_))])), d_1853_v1_)))), (d_1854_v2_)[default__.safeIndex((self).f11, len(d_1854_v2_))])
            nw311_[int(10)] = d_1854_v2_
            nw311_[int(11)] = d_1854_v2_
            nw311_[int(12)] = (d_1854_v2_).set(default__.safeIndex((self).f11, len(d_1854_v2_)), d_1853_v1_)
            nw311_[int(13)] = d_1854_v2_
            nw311_[int(14)] = d_1854_v2_
            nw311_[int(15)] = d_1854_v2_
            nw311_[int(16)] = ((d_1859_v7_)[d_1858_v6_] if (d_1858_v6_) in (d_1859_v7_) else d_1854_v2_)
            nw311_[int(17)] = d_1854_v2_
            nw311_[int(18)] = d_1854_v2_
            nw311_[int(19)] = d_1854_v2_
            nw311_[int(20)] = d_1854_v2_
            nw311_[int(21)] = d_1854_v2_
            nw311_[int(22)] = (d_1854_v2_) + (d_1854_v2_)
            nw311_[int(23)] = d_1854_v2_
            nw311_[int(24)] = d_1854_v2_
            nw311_[int(25)] = _dafny.SeqWithoutIsStrInference([d_1853_v1_])
            nw311_[int(26)] = (d_1854_v2_) + (d_1854_v2_)
            d_1860_v8_ = nw311_
            index265_ = default__.safeIndex(825, (d_1860_v8_).length(0))
            (d_1860_v8_)[index265_] = d_1854_v2_
            index266_ = default__.safeIndex(825, (d_1860_v8_).length(0))
            (d_1860_v8_)[index266_] = d_1854_v2_
            d_1861_v9_: _dafny.Seq
            d_1861_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doukbfk"))
            d_1861_v9_ = d_1861_v9_
            d_1862_v10_: str
            d_1862_v10_ = _dafny.CodePoint('m')
            d_1862_v10_ = d_1862_v10_
            (globalState).f1 = (self).f11
            d_1863_v11_: _dafny.Map
            d_1863_v11_ = _dafny.Map({(self).f13: (self).f11})
            d_1864_v12_: _dafny.Map
            d_1864_v12_ = _dafny.Map({d_1863_v11_: _dafny.Set({default__.fm0(globalState)})})
            d_1865_v13_: C2
            nw312_ = C2()
            nw312_.ctor__(d_1862_v10_, len((d_1864_v12_) | (d_1864_v12_)))
            d_1865_v13_ = nw312_
        elif True:
            d_1866_v14_: _dafny.MultiSet
            d_1866_v14_ = _dafny.MultiSet([(0) - ((self).f11)])
            d_1867_v15_: T0
            nw313_ = C4()
            nw313_.ctor__(d_1866_v14_, (self).f13, (self).f12, (self).f11)
            d_1867_v15_ = nw313_
            d_1868_v16_: _dafny.Map
            d_1868_v16_ = _dafny.Map({(self).f13: d_1867_v15_})
            d_1869_v17_: _dafny.Map
            d_1869_v17_ = _dafny.Map({d_1866_v14_: _dafny.Map({(self).f11: len(d_1868_v16_)})})
            d_1870_v18_: _dafny.MultiSet
            d_1870_v18_ = _dafny.MultiSet([(self).f13, False])
            d_1871_v19_: _dafny.Map
            d_1871_v19_ = _dafny.Map({(self).fm4((self).f11, d_1869_v17_, ((d_1870_v18_)[(self).f13] if ((self).f13) in (d_1870_v18_) else (self).f11), (self).f13, globalState): (d_1867_v15_).f11})
            d_1872_v21_: _dafny.Seq
            d_1872_v21_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iswspx")))])
            d_1873_v22_: _dafny.Array
            nw314_ = _dafny.Array(None, 26)
            nw314_[int(0)] = (self).f11
            nw314_[int(1)] = (self).f11
            nw314_[int(2)] = (d_1867_v15_).f11
            nw314_[int(3)] = (912 if not((self).f13) else (self).f11)
            nw314_[int(4)] = 331
            nw314_[int(5)] = (self).f11
            nw314_[int(6)] = len((self).f12)
            nw314_[int(7)] = default__.safeModuloInt((d_1867_v15_).f11, 382)
            nw314_[int(8)] = (d_1867_v15_).f11
            nw314_[int(9)] = (d_1867_v15_).f11
            nw314_[int(10)] = (d_1867_v15_).f11
            nw314_[int(11)] = (d_1867_v15_).f11
            nw314_[int(12)] = (d_1867_v15_).f11
            nw314_[int(13)] = (d_1867_v15_).f11
            nw314_[int(14)] = (d_1867_v15_).f11
            nw314_[int(15)] = (934) - ((self).f11)
            nw314_[int(16)] = (d_1867_v15_).f11
            def iife157_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in (d_1866_v14_).Elements:
                    d_1874_v20_: int = compr_69_
                    if (d_1874_v20_) in (d_1866_v14_):
                        coll69_[(d_1874_v20_) + ((d_1867_v15_).f11)] = 244
                return _dafny.Map(coll69_)
            nw314_[int(17)] = len(iife157_()
            )
            nw314_[int(18)] = (self).f11
            nw314_[int(19)] = (d_1867_v15_).f11
            nw314_[int(20)] = (d_1867_v15_).f11
            nw314_[int(21)] = (d_1867_v15_).f11
            nw314_[int(22)] = (self).f11
            nw314_[int(23)] = ((d_1867_v15_).f11) - ((self).f11)
            nw314_[int(24)] = (self).f11
            nw314_[int(25)] = ((d_1872_v21_)[default__.safeIndex((d_1867_v15_).f11, len(d_1872_v21_))] if True else 723)
            d_1873_v22_ = nw314_
            index267_ = default__.safeIndex(373, (d_1873_v22_).length(0))
            (d_1873_v22_)[index267_] = len(_dafny.SeqWithoutIsStrInference([p0 for d_1875_i0_ in range(default__.abs(51))]))
            d_1876_v23_: D2
            d_1876_v23_ = D2_DC4(894, (d_1867_v15_).f11, (self).f11)
            d_1877_v24_: _dafny.Seq
            d_1877_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlasr"))
            d_1878_v25_: C6
            nw315_ = C6()
            nw315_.ctor__(d_1876_v23_, _dafny.Set({not(False)}), (0) - (len((d_1877_v24_) + ((default__.fm1(globalState)).set(default__.safeIndex((self).f11, len(default__.fm1(globalState))), p0)))))
            d_1878_v25_ = nw315_
            d_1879_v26_: _dafny.Seq
            d_1879_v26_ = _dafny.SeqWithoutIsStrInference([not((self).f13), (self).f13])
            d_1880_v27_: _dafny.Seq
            d_1880_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twrkpvv"))])
            index268_ = default__.safeIndex(373, (d_1873_v22_).length(0))
            rhs332_ = default__.safeDivisionInt(len(d_1879_v26_), len(d_1872_v21_))
            rhs333_ = (d_1871_v19_) | (d_1871_v19_)
            rhs334_ = ((_dafny.MultiSet([d_1877_v24_])).intersection(_dafny.MultiSet(d_1880_v27_))).cardinality
            rhs335_ = d_1878_v25_
            rhs336_ = default__.safeDivisionInt(((self).f11 if (self).fm4((d_1867_v15_).f11, d_1869_v17_, ((d_1866_v14_)[(self).f11] if ((self).f11) in (d_1866_v14_) else len(d_1879_v26_)), (self).f13, globalState) else (d_1867_v15_).f11), (self).f11)
            lhs279_ = globalState
            lhs280_ = d_1873_v22_
            lhs281_ = default__.safeIndex(373, (d_1873_v22_).length(0))
            lhs282_ = globalState
            lhs279_.f1 = rhs332_
            d_1871_v19_ = rhs333_
            lhs280_[lhs281_] = rhs334_
            d_1878_v25_ = rhs335_
            lhs282_.f1 = rhs336_
            d_1881_v28_: _dafny.Array
            d_1882_v29_: _dafny.Map
            out33_: _dafny.Array
            out34_: _dafny.Map
            out33_, out34_ = (d_1878_v25_).m1((d_1872_v21_)[default__.safeIndex((d_1873_v22_)[default__.safeIndex(373, (d_1873_v22_).length(0))], len(d_1872_v21_))], globalState)
            d_1881_v28_ = out33_
            d_1882_v29_ = out34_
            (globalState).f1 = (self).f11
            d_1883_v30_: _dafny.Seq
            d_1883_v30_ = _dafny.SeqWithoutIsStrInference([d_1872_v21_, d_1872_v21_])
            (globalState).f1 = ((self).f11) + (len((d_1872_v21_ if (self).f13 else (d_1883_v30_)[default__.safeIndex((d_1867_v15_).f11, len(d_1883_v30_))])))
            index269_ = default__.safeIndex(373, (d_1873_v22_).length(0))
            (d_1873_v22_)[index269_] = (0) - ((0) - ((d_1873_v22_)[default__.safeIndex(373, (d_1873_v22_).length(0))]))
        d_1884_v31_: _dafny.Seq
        d_1884_v31_ = _dafny.SeqWithoutIsStrInference([-537])
        d_1885_v32_: D0
        d_1885_v32_ = D0_DC1((self).f13, (self).f11, (self).f13)
        d_1886_v33_: _dafny.Seq
        d_1886_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "liexbvcq"))
        d_1887_v34_: _dafny.Array
        nw316_ = _dafny.Array(False, 28)
        d_1887_v34_ = nw316_
        d_1888_v35_: _dafny.Set
        d_1888_v35_ = _dafny.Set({d_1887_v34_, d_1887_v34_})
        d_1889_v36_: D21
        d_1889_v36_ = D21_DC47(d_1888_v35_)
        d_1890_v37_: _dafny.Array
        nw317_ = _dafny.Array(None, 21)
        nw317_[int(0)] = not((_dafny.SeqWithoutIsStrInference([(self).f11])) < (d_1884_v31_))
        nw317_[int(1)] = (self).f13
        nw317_[int(2)] = (d_1885_v32_).cf3
        nw317_[int(3)] = (p0) in ((d_1886_v33_).set(default__.safeIndex((self).f11, len(d_1886_v33_)), _dafny.CodePoint('q')))
        nw317_[int(4)] = ((d_1889_v36_).cf77) != (_dafny.Set({d_1887_v34_}))
        nw317_[int(5)] = (self).f13
        nw317_[int(6)] = (self).f13
        nw317_[int(7)] = not(((self).f13 if (self).f13 else (self).f13))
        nw317_[int(8)] = ((self).f11) != ((self).f11)
        nw317_[int(9)] = (self).f13
        nw317_[int(10)] = (self).f13
        nw317_[int(11)] = (self).f13
        nw317_[int(12)] = (self).f13
        nw317_[int(13)] = (self).f13
        nw317_[int(14)] = not ((self).f13) or ((self).f13)
        nw317_[int(15)] = (self).f13
        nw317_[int(16)] = (self).f13
        nw317_[int(17)] = (self).f13
        nw317_[int(18)] = (self).f13
        nw317_[int(19)] = (self).f13
        nw317_[int(20)] = (self).f13
        d_1890_v37_ = nw317_
        index270_ = default__.safeIndex(945, (d_1890_v37_).length(0))
        (d_1890_v37_)[index270_] = True
        d_1891_v38_: D7
        d_1891_v38_ = D7_DC14(D7_DC12(d_1884_v31_))
        d_1892_v39_: _dafny.Map
        d_1892_v39_ = _dafny.Map({default__.fm2(len(d_1884_v31_), (self).f11, globalState): not((self).f13)})
        d_1893_v41_: _dafny.MultiSet
        d_1893_v41_ = _dafny.MultiSet([(self).f11, (self).f11, (self).f11])
        d_1894_v42_: _dafny.Set
        def iife158_():
            coll70_ = _dafny.Map()
            compr_70_: int
            for compr_70_ in ((d_1893_v41_).set((self).f11, default__.abs((0) - ((self).f11)))).Elements:
                d_1895_v40_: int = compr_70_
                if (d_1895_v40_) in ((d_1893_v41_).set((self).f11, default__.abs((0) - ((self).f11)))):
                    coll70_[(d_1895_v40_) * ((self).f11)] = not((self).f13)
            return _dafny.Map(coll70_)
        d_1894_v42_ = _dafny.Set({d_1892_v39_, (d_1892_v39_ if (self).f13 else d_1892_v39_), d_1892_v39_, (iife158_()
        ) | (d_1892_v39_), _dafny.Map({len(d_1886_v33_): (self).f13})})
        d_1896_v44_: _dafny.Map
        def iife159_():
            coll71_ = _dafny.Map()
            compr_71_: int
            for compr_71_ in _dafny.IntegerRange(50, -202):
                d_1897_v43_: int = compr_71_
                if ((50) <= (d_1897_v43_)) and ((d_1897_v43_) < (-202)):
                    coll71_[default__.safeDivisionInt(d_1897_v43_, (self).f11)] = (0) - ((self).f11)
            return _dafny.Map(coll71_)
        d_1896_v44_ = _dafny.Map({d_1893_v41_: iife159_()
        })
        d_1898_v45_: _dafny.Seq
        d_1898_v45_ = _dafny.SeqWithoutIsStrInference([d_1884_v31_])
        d_1899_v46_: C7
        nw318_ = C7()
        nw318_.ctor__(d_1898_v45_, (self).f12, len(_dafny.Map({(0) - ((self).f11): True})))
        d_1899_v46_ = nw318_
        d_1900_v47_: _dafny.Set
        d_1900_v47_ = _dafny.Set({d_1899_v46_, d_1899_v46_})
        d_1901_v48_: _dafny.Map
        d_1901_v48_ = _dafny.Map({(self).fm4((self).f11, d_1896_v44_, len(d_1900_v47_), (self).f13, globalState): (self).f11})
        d_1902_v49_: _dafny.Map
        d_1902_v49_ = _dafny.Map({(0) - ((self).fm3(len(d_1901_v48_), (self).f13, (self).f13, globalState)): 452})
        d_1903_v50_: _dafny.Map
        d_1903_v50_ = _dafny.Map({_dafny.MultiSet([len(d_1884_v31_)]): d_1902_v49_})
        d_1904_v52_: _dafny.Map
        def iife160_():
            coll72_ = _dafny.Map()
            compr_72_: int
            for compr_72_ in (_dafny.Map({(self).f11: (self).f13})).keys.Elements:
                d_1905_v51_: int = compr_72_
                if (d_1905_v51_) in (_dafny.Map({(self).f11: (self).f13})):
                    coll72_[default__.safeDivisionInt(d_1905_v51_, (self).f11)] = (self).f13
            return _dafny.Map(coll72_)
        d_1904_v52_ = _dafny.Map({iife160_()
        : (self).f13})
        index271_ = default__.safeIndex(945, (d_1890_v37_).length(0))
        def iife161_():
            coll73_ = _dafny.Set()
            compr_73_: _dafny.Map
            for compr_73_ in (d_1904_v52_).keys.Elements:
                d_1906_v53_: _dafny.Map = compr_73_
                if (d_1906_v53_) in (d_1904_v52_):
                    coll73_ = coll73_.union(_dafny.Set([d_1906_v53_]))
            return _dafny.Set(coll73_)
        rhs337_ = (self).fm4(((self).f11) * ((self).f11), d_1896_v44_, (self).f11, True, globalState)
        rhs338_ = d_1891_v38_
        rhs339_ = iife161_()

        rhs340_ = (self).f11
        lhs283_ = d_1890_v37_
        lhs284_ = default__.safeIndex(945, (d_1890_v37_).length(0))
        lhs285_ = globalState
        lhs283_[lhs284_] = rhs337_
        d_1891_v38_ = rhs338_
        d_1894_v42_ = rhs339_
        lhs285_.f1 = rhs340_
        d_1907_v54_: _dafny.Map
        d_1907_v54_ = _dafny.Map({(self).f11: _dafny.CodePoint('h')})
        d_1908_v55_: _dafny.Seq
        d_1908_v55_ = _dafny.SeqWithoutIsStrInference([(d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], (self).f13])
        d_1909_v56_: _dafny.Array
        nw319_ = _dafny.Array(None, 21)
        nw319_[int(0)] = default__.fm1(globalState)
        nw319_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yur"))
        nw319_[int(2)] = _dafny.SeqWithoutIsStrInference([p0 for d_1910_i1_ in range(default__.abs(677))])
        nw319_[int(3)] = _dafny.SeqWithoutIsStrInference([p0 for d_1911_i2_ in range(default__.abs(527))])
        nw319_[int(4)] = d_1886_v33_
        nw319_[int(5)] = d_1886_v33_
        nw319_[int(6)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivoblt"))).set(default__.safeIndex(448, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivoblt")))), ((d_1907_v54_)[len(d_1908_v55_)] if (len(d_1908_v55_)) in (d_1907_v54_) else p0))
        nw319_[int(7)] = d_1886_v33_
        nw319_[int(8)] = d_1886_v33_
        nw319_[int(9)] = (d_1886_v33_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1912_i3_ in range(default__.abs(716))]))
        nw319_[int(10)] = d_1886_v33_
        nw319_[int(11)] = d_1886_v33_
        nw319_[int(12)] = d_1886_v33_
        nw319_[int(13)] = d_1886_v33_
        nw319_[int(14)] = d_1886_v33_
        nw319_[int(15)] = (d_1886_v33_) + (d_1886_v33_)
        nw319_[int(16)] = d_1886_v33_
        nw319_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugocbljk"))
        nw319_[int(18)] = d_1886_v33_
        nw319_[int(19)] = d_1886_v33_
        nw319_[int(20)] = d_1886_v33_
        d_1909_v56_ = nw319_
        d_1909_v56_ = d_1909_v56_
        if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]:
            (globalState).f7 = not ((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]) or ((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))])
            (globalState).f1 = (self).f11
            (globalState).f6 = (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]
            (globalState).f1 = len(d_1908_v55_)
            index272_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            (d_1890_v37_)[index272_] = (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]
        elif True:
            d_1913_v57_: str
            d_1913_v57_ = _dafny.CodePoint('k')
            d_1914_v58_: _dafny.Array
            nw320_ = _dafny.Array(int(0), 28)
            d_1914_v58_ = nw320_
            d_1915_v59_: _dafny.Set
            d_1915_v59_ = _dafny.Set({d_1914_v58_})
            rhs341_ = (163) * (len(d_1915_v59_))
            rhs342_ = (d_1886_v33_)[default__.safeIndex(len((d_1908_v55_) + (d_1908_v55_)), len(d_1886_v33_))]
            rhs343_ = (self).f11
            lhs286_ = globalState
            lhs287_ = globalState
            lhs286_.f1 = rhs341_
            d_1913_v57_ = rhs342_
            lhs287_.f1 = rhs343_
            d_1916_v60_: D7
            d_1916_v60_ = D7_DC12((d_1884_v31_) + (_dafny.SeqWithoutIsStrInference([(self).f11 for d_1917_i4_ in range(default__.abs(335))])))
            source28_ = d_1916_v60_
            if source28_.is_DC13:
                d_1918___mcc_h0_ = source28_.cf18
                d_1919___mcc_h1_ = source28_.cf19
                d_1920___mcc_h2_ = source28_.cf20
                d_1921_cf20_ = d_1920___mcc_h2_
                d_1922_cf19_ = d_1919___mcc_h1_
                d_1923_cf18_ = d_1918___mcc_h0_
                d_1924_v61_: D8
                d_1924_v61_ = D8_DC16((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], (self).f11)
                d_1925_v62_: C3
                nw321_ = C3()
                nw321_.ctor__(d_1921_cf20_, 36, d_1921_cf20_, (self).f12)
                d_1925_v62_ = nw321_
                d_1926_v63_: _dafny.Map
                d_1926_v63_ = _dafny.Map({d_1925_v62_: d_1925_v62_.f23})
                d_1927_v64_: _dafny.Set
                d_1927_v64_ = _dafny.Set({d_1926_v63_})
                d_1928_v65_: D2
                d_1928_v65_ = D2_DC4((0) - ((d_1924_v61_).cf24), len(d_1927_v64_), d_1925_v62_.f23)
                d_1929_v66_: C6
                nw322_ = C6()
                nw322_.ctor__(d_1928_v65_, ((self).f12).intersection(d_1923_cf18_), (self).f11)
                d_1929_v66_ = nw322_
                d_1930_v67_: _dafny.Seq
                d_1930_v67_ = _dafny.SeqWithoutIsStrInference([d_1914_v58_, d_1914_v58_, d_1914_v58_, d_1914_v58_])
                index273_ = default__.safeIndex(782, (d_1914_v58_).length(0))
                (d_1914_v58_)[index273_] = len(d_1930_v67_)
                index274_ = default__.safeIndex(782, (d_1914_v58_).length(0))
                (d_1914_v58_)[index274_] = default__.safeModuloInt(default__.safeDivisionInt((self).f11, d_1921_cf20_), (0) - ((d_1925_v62_).f22))
                index275_ = default__.safeIndex(782, (d_1914_v58_).length(0))
                (d_1914_v58_)[index275_] = (0) - ((((d_1924_v61_).cf24) * (-219)) * ((self).f11))
                d_1931_v68_: _dafny.Set
                d_1931_v68_ = _dafny.Set({(d_1925_v62_).f22})
                d_1932_v70_: _dafny.Seq
                d_1932_v70_ = _dafny.SeqWithoutIsStrInference([d_1931_v68_, d_1931_v68_, d_1931_v68_])
                d_1933_v71_: _dafny.Map
                d_1933_v71_ = _dafny.Map({d_1913_v57_: (d_1914_v58_)[default__.safeIndex(782, (d_1914_v58_).length(0))]})
                d_1934_v72_: _dafny.Map
                d_1934_v72_ = _dafny.Map({_dafny.MultiSet([default__.fm54(d_1901_v48_, globalState), d_1933_v71_]): d_1931_v68_})
                d_1935_v75_: _dafny.Map
                d_1935_v75_ = _dafny.Map({d_1886_v33_: default__.fm18(len((self).f12), globalState)})
                d_1936_v76_: _dafny.Array
                nw323_ = _dafny.Array(None, 24)
                nw323_[int(0)] = d_1931_v68_
                def iife162_():
                    coll74_ = _dafny.Set()
                    compr_74_: int
                    for compr_74_ in _dafny.IntegerRange(212, 905):
                        d_1937_v69_: int = compr_74_
                        if ((212) <= (d_1937_v69_)) and ((d_1937_v69_) < (905)):
                            coll74_ = coll74_.union(_dafny.Set([default__.safeModuloInt(d_1937_v69_, (self).f11)]))
                    return _dafny.Set(coll74_)
                nw323_[int(1)] = (iife162_()
                ) | (d_1931_v68_)
                nw323_[int(2)] = (d_1931_v68_) | (_dafny.Set({d_1921_cf20_}))
                nw323_[int(3)] = d_1931_v68_
                nw323_[int(4)] = (d_1932_v70_)[default__.safeIndex(d_1922_cf19_, len(d_1932_v70_))]
                nw323_[int(5)] = (D22_DC51(d_1931_v68_)).cf86
                nw323_[int(6)] = _dafny.Set({49, d_1921_cf20_, (d_1925_v62_).f22})
                nw323_[int(7)] = d_1931_v68_
                nw323_[int(8)] = d_1931_v68_
                nw323_[int(9)] = d_1931_v68_
                nw323_[int(10)] = (d_1931_v68_) | (d_1931_v68_)
                nw323_[int(11)] = (d_1931_v68_).intersection(d_1931_v68_)
                nw323_[int(12)] = d_1931_v68_
                def iife163_():
                    coll75_ = _dafny.Map()
                    compr_75_: str
                    for compr_75_ in (_dafny.SeqWithoutIsStrInference([p0 for d_1938_i5_ in range(default__.abs(768))])).Elements:
                        d_1939_v73_: str = compr_75_
                        if (d_1939_v73_) in (_dafny.SeqWithoutIsStrInference([p0 for d_1938_i5_ in range(default__.abs(768))])):
                            coll75_[d_1939_v73_] = d_1925_v62_.f23
                    return _dafny.Map(coll75_)
                def iife164_():
                    coll76_ = _dafny.Map()
                    compr_76_: str
                    for compr_76_ in (_dafny.SeqWithoutIsStrInference([p0 for d_1940_i5_ in range(default__.abs(768))])).Elements:
                        d_1941_v73_: str = compr_76_
                        if (d_1941_v73_) in (_dafny.SeqWithoutIsStrInference([p0 for d_1940_i5_ in range(default__.abs(768))])):
                            coll76_[d_1941_v73_] = d_1925_v62_.f23
                    return _dafny.Map(coll76_)
                nw323_[int(13)] = ((d_1934_v72_)[_dafny.MultiSet([iife163_()
                ])] if (_dafny.MultiSet([iife164_()
                ])) in (d_1934_v72_) else d_1931_v68_)
                nw323_[int(14)] = _dafny.Set({len(d_1886_v33_), (_dafny.MultiSet([(d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]])).cardinality})
                nw323_[int(15)] = default__.fm26(globalState)
                nw323_[int(16)] = d_1931_v68_
                def iife165_():
                    coll77_ = _dafny.Set()
                    compr_77_: int
                    for compr_77_ in _dafny.IntegerRange(-328, 346):
                        d_1942_v74_: int = compr_77_
                        if ((-328) <= (d_1942_v74_)) and ((d_1942_v74_) < (346)):
                            coll77_ = coll77_.union(_dafny.Set([default__.safeDivisionInt(d_1942_v74_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mak"))))]))
                    return _dafny.Set(coll77_)
                nw323_[int(17)] = (iife165_()
                ).intersection(_dafny.Set({d_1925_v62_.f23, (self).f11, (d_1884_v31_)[default__.safeIndex((d_1914_v58_)[default__.safeIndex(782, (d_1914_v58_).length(0))], len(d_1884_v31_))]}))
                nw323_[int(18)] = d_1931_v68_
                nw323_[int(19)] = ((d_1935_v75_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcfn"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcfn"))) in (d_1935_v75_) else d_1931_v68_)
                nw323_[int(20)] = d_1931_v68_
                nw323_[int(21)] = d_1931_v68_
                nw323_[int(22)] = default__.fm26(globalState)
                nw323_[int(23)] = d_1931_v68_
                d_1936_v76_ = nw323_
                d_1936_v76_ = d_1936_v76_
            elif source28_.is_DC12:
                d_1943___mcc_h3_ = source28_.cf17
                d_1944_cf17_ = d_1943___mcc_h3_
                index276_ = default__.safeIndex(549, (d_1914_v58_).length(0))
                (d_1914_v58_)[index276_] = ((self).f11) + ((self).f11)
                d_1945_v77_: _dafny.Seq
                d_1945_v77_ = _dafny.SeqWithoutIsStrInference([d_1892_v39_, _dafny.Map({(self).f11: (self).f13})])
                index277_ = default__.safeIndex(549, (d_1914_v58_).length(0))
                index278_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                rhs344_ = (d_1908_v55_) + (d_1908_v55_)
                rhs345_ = (self).fm3(-429, ((self).f11) in ((d_1945_v77_)[default__.safeIndex((self).f11, len(d_1945_v77_))]), (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], globalState)
                rhs346_ = default__.safeModuloInt((self).f11, (self).f11)
                rhs347_ = ((d_1886_v33_) + (d_1886_v33_)) != (d_1886_v33_)
                rhs348_ = (self).f11
                lhs288_ = d_1914_v58_
                lhs289_ = default__.safeIndex(549, (d_1914_v58_).length(0))
                lhs290_ = globalState
                lhs291_ = d_1890_v37_
                lhs292_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                lhs293_ = globalState
                d_1908_v55_ = rhs344_
                lhs288_[lhs289_] = rhs345_
                lhs290_.f1 = rhs346_
                lhs291_[lhs292_] = rhs347_
                lhs293_.f1 = rhs348_
                d_1886_v33_ = (d_1886_v33_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okrvwip"))) + (d_1886_v33_))
                (globalState).f1 = (0) - ((self).f11)
                d_1946_v78_: _dafny.Map
                d_1946_v78_ = _dafny.Map({(d_1908_v55_) + (d_1908_v55_): (d_1914_v58_)[default__.safeIndex(549, (d_1914_v58_).length(0))]})
                d_1947_v79_: D7
                d_1947_v79_ = D7_DC13((self).f12, len(d_1907_v54_), 485)
                d_1946_v78_ = (d_1946_v78_).set(default__.fm35(len(d_1886_v33_), (d_1914_v58_)[default__.safeIndex(549, (d_1914_v58_).length(0))], d_1947_v79_, (d_1893_v41_).cardinality, globalState), (self).f11)
            elif True:
                d_1948___mcc_h4_ = source28_.cf21
                d_1949_cf21_ = d_1948___mcc_h4_
                d_1886_v33_ = d_1886_v33_
                (globalState).f1 = (self).f11
                d_1950_v80_: D8
                d_1950_v80_ = D8_DC16((self).f13, len(_dafny.Set({(self).f11})))
                d_1951_v81_: D8
                d_1951_v81_ = D8_DC17(d_1950_v80_)
                d_1952_v82_: D8
                d_1952_v82_ = D8_DC17(d_1950_v80_)
                d_1953_v83_: D8
                d_1953_v83_ = D8_DC17(d_1951_v81_)
                d_1954_v84_: _dafny.Map
                d_1954_v84_ = _dafny.Map({(d_1953_v83_ if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))] else d_1953_v83_): _dafny.CodePoint('p')})
                d_1954_v84_ = (d_1954_v84_).set((d_1953_v83_ if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))] else d_1953_v83_), p0)
                d_1955_v85_: _dafny.Set
                d_1955_v85_ = _dafny.Set({_dafny.MultiSet([(d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]])})
                d_1956_v86_: _dafny.Map
                d_1956_v86_ = _dafny.Map({(default__.fm12((self).f11, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwgdxagxs")), globalState)) + (d_1908_v55_): d_1955_v85_})
                d_1956_v86_ = (d_1956_v86_).set((d_1908_v55_) + (_dafny.SeqWithoutIsStrInference([True, (self).f13])), d_1955_v85_)
            d_1957_v87_: _dafny.Map
            d_1957_v87_ = _dafny.Map({d_1886_v33_: (self).f11})
            d_1958_v88_: T0
            nw324_ = C4()
            nw324_.ctor__(_dafny.MultiSet([(self).f11]), (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], (self).f12, (self).f11)
            d_1958_v88_ = nw324_
            d_1959_v89_: _dafny.Map
            d_1959_v89_ = _dafny.Map({len(d_1957_v87_): d_1958_v88_})
            d_1960_v90_: _dafny.Map
            d_1960_v90_ = _dafny.Map({len(d_1902_v49_): d_1907_v54_})
            d_1961_v91_: _dafny.Map
            d_1961_v91_ = _dafny.Map({(self).f13: p0})
            d_1962_v92_: D11
            d_1962_v92_ = D11_DC23(True, (d_1958_v88_).f11, len(d_1961_v91_))
            d_1963_v94_: _dafny.Set
            def iife166_():
                coll78_ = _dafny.Map()
                compr_78_: int
                for compr_78_ in _dafny.IntegerRange(301, 662):
                    d_1964_v93_: int = compr_78_
                    if ((301) <= (d_1964_v93_)) and ((d_1964_v93_) < (662)):
                        coll78_[default__.safeModuloInt(d_1964_v93_, (self).f11)] = p0
                return _dafny.Map(coll78_)
            d_1963_v94_ = _dafny.Set({d_1907_v54_, _dafny.Map({len(default__.fm46(globalState)): _dafny.CodePoint('k')}), (_dafny.Map({(d_1958_v88_).f11: _dafny.CodePoint('p')})).set((d_1962_v92_).cf33, d_1913_v57_), iife166_()
            })
            if (((_dafny.Map({315: d_1913_v57_})).set(len(d_1959_v89_), p0)) | (((d_1960_v90_)[(d_1958_v88_).f11] if ((d_1958_v88_).f11) in (d_1960_v90_) else _dafny.Map({len(d_1908_v55_): p0})))) not in (d_1963_v94_):
                index279_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                (d_1890_v37_)[index279_] = False
                d_1965_v95_: D0
                d_1965_v95_ = D0_DC0((d_1884_v31_)[default__.safeIndex(len(d_1901_v48_), len(d_1884_v31_))])
                pat_let_tv50_ = d_1916_v60_
                index280_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                def iife168_(_pat_let45_0):
                    def iife169_(d_1966_dt__update__tmp_h0_):
                        def iife170_(_pat_let46_0):
                            def iife171_(d_1967_dt__update_hcf0_h0_):
                                return D0_DC0(d_1967_dt__update_hcf0_h0_)
                            return iife171_(_pat_let46_0)
                        return iife170_(128)
                    return iife169_(_pat_let45_0)
                def iife167_(_pat_let44_0):
                    def iife172_(d_1968_dt__update__tmp_h1_):
                        def iife173_(_pat_let47_0):
                            def iife174_(d_1969_dt__update_hcf0_h1_):
                                return D0_DC0(d_1969_dt__update_hcf0_h1_)
                            return iife174_(_pat_let47_0)
                        return iife173_(len((pat_let_tv50_).cf17))
                    return iife172_(_pat_let44_0)
                (d_1890_v37_)[index280_] = (d_1899_v46_).fm5(iife167_(iife168_(d_1965_v95_)), d_1908_v55_, default__.safeDivisionInt((self).f11, 830), d_1913_v57_, globalState)
                (globalState).f1 = (default__.safeDivisionInt((self).f11, (self).f11)) + (684)
                d_1970_v96_: _dafny.Map
                d_1970_v96_ = _dafny.Map({(self).f13: (self).f13})
                d_1971_v97_: D16
                d_1971_v97_ = D16_DC39(len(d_1902_v49_), (self).f13, d_1970_v96_, _dafny.CodePoint('n'))
                d_1972_v98_: C2
                nw325_ = C2()
                nw325_.ctor__((d_1971_v97_).cf65, (d_1958_v88_).f11)
                d_1972_v98_ = nw325_
                d_1973_v99_: _dafny.Map
                d_1973_v99_ = _dafny.Map({d_1972_v98_: (d_1958_v88_).f11})
                d_1973_v99_ = (d_1973_v99_).set(d_1972_v98_, ((d_1958_v88_).f11) - ((d_1958_v88_).f11))
                d_1974_v100_: _dafny.Array
                nw326_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_1974_v100_ = nw326_
                index281_ = default__.safeIndex(857, (d_1974_v100_).length(0))
                (d_1974_v100_)[index281_] = d_1887_v34_
                index282_ = default__.safeIndex(857, (d_1974_v100_).length(0))
                rhs349_ = d_1887_v34_
                rhs350_ = (d_1958_v88_).f11
                lhs294_ = d_1974_v100_
                lhs295_ = default__.safeIndex(857, (d_1974_v100_).length(0))
                lhs296_ = globalState
                lhs294_[lhs295_] = rhs349_
                lhs296_.f1 = rhs350_
            elif True:
                (globalState).f6 = (d_1962_v92_).cf31
                index283_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                (d_1890_v37_)[index283_] = (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]
                (globalState).f1 = (d_1958_v88_).f11
                d_1975_v101_: C2
                nw327_ = C2()
                nw327_.ctor__(_dafny.CodePoint('f'), (d_1958_v88_).f11)
                d_1975_v101_ = nw327_
                pat_let_tv51_ = d_1886_v33_
                def iife175_(_pat_let48_0):
                    def iife176_(d_1976_dt__update__tmp_h2_):
                        def iife177_(_pat_let49_0):
                            def iife178_(d_1977_dt__update_hcf19_h0_):
                                return D7_DC13((d_1976_dt__update__tmp_h2_).cf18, d_1977_dt__update_hcf19_h0_, (d_1976_dt__update__tmp_h2_).cf20)
                            return iife178_(_pat_let49_0)
                        return iife177_(len(pat_let_tv51_))
                    return iife176_(_pat_let48_0)
                (globalState).f1 = ((d_1958_v88_).f11) * ((iife175_(D7_DC13((self).f12, (d_1958_v88_).f11, (d_1958_v88_).f11))).cf20)
            d_1978_v102_: D6
            d_1978_v102_ = D6_DC9(d_1892_v39_)
            pat_let_tv52_ = d_1892_v39_
            def iife179_(_pat_let50_0):
                def iife180_(d_1979_dt__update__tmp_h3_):
                    def iife181_(_pat_let51_0):
                        def iife182_(d_1980_dt__update_hcf12_h0_):
                            return D6_DC9(d_1980_dt__update_hcf12_h0_)
                        return iife182_(_pat_let51_0)
                    return iife181_(pat_let_tv52_)
                return iife180_(_pat_let50_0)
            d_1978_v102_ = iife179_(d_1978_v102_)
            (globalState).f6 = (not((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))])) or ((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))])
        if (((self).f11 if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))] else (self).f11)) <= ((self).f11):
            rhs351_ = (self).f13
            rhs352_ = (False) and ((self).f13)
            rhs353_ = ((self).f11 if ((self).f13) or (not((self).f13)) else (self).f11)
            lhs297_ = globalState
            lhs298_ = globalState
            lhs299_ = globalState
            lhs297_.f7 = rhs351_
            lhs298_.f6 = rhs352_
            lhs299_.f1 = rhs353_
            d_1907_v54_ = (d_1907_v54_).set(569, p0)
            (globalState).f7 = (((self).f11) - ((self).f11)) > ((self).f11)
            index284_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            (d_1890_v37_)[index284_] = (self).f13
            pat_let_tv53_ = d_1888_v35_
            d_1981_v103_: _dafny.MultiSet
            def iife183_(_pat_let52_0):
                def iife184_(d_1982_dt__update__tmp_h4_):
                    def iife185_(_pat_let53_0):
                        def iife186_(d_1983_dt__update_hcf77_h0_):
                            return D21_DC47(d_1983_dt__update_hcf77_h0_)
                        return iife186_(_pat_let53_0)
                    return iife185_(pat_let_tv53_)
                return iife184_(_pat_let52_0)
            d_1981_v103_ = _dafny.MultiSet([d_1889_v36_, d_1889_v36_, d_1889_v36_, d_1889_v36_, iife183_(d_1889_v36_)])
            d_1981_v103_ = (d_1981_v103_).set(d_1889_v36_, default__.abs((self).f11))
        elif True:
            (globalState).f1 = (self).f11
            (globalState).f1 = 112
            d_1984_v104_: D2
            d_1984_v104_ = D2_DC4(121, (self).f11, 728)
            d_1985_v105_: _dafny.Seq
            d_1985_v105_ = _dafny.SeqWithoutIsStrInference([d_1984_v104_])
            d_1986_v106_: _dafny.MultiSet
            d_1986_v106_ = _dafny.MultiSet([(d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], (self).f13])
            (globalState).f1 = ((self).f11) - (len(((d_1985_v105_) + (d_1985_v105_)).set(default__.safeIndex((d_1986_v106_).cardinality, len((d_1985_v105_) + (d_1985_v105_))), D2_DC4(len(d_1884_v31_), (self).f11, (self).f11))))
            d_1987_v107_: _dafny.Seq
            d_1987_v107_ = _dafny.SeqWithoutIsStrInference([d_1908_v55_])
            d_1988_v108_: C0
            nw328_ = C0()
            nw328_.ctor__(((d_1987_v107_)[default__.safeIndex((self).f11, len(d_1987_v107_))]) + (d_1908_v55_), (self).fm7((self).f11, (self).f13, _dafny.MultiSet(d_1908_v55_), globalState))
            d_1988_v108_ = nw328_
            d_1988_v108_ = d_1988_v108_
            (globalState).f1 = 901
        d_1989_v109_: D16
        d_1989_v109_ = D16_DC38(d_1902_v49_)
        def lambda73_(source30_):
            if source30_.is_DC39:
                d_1990___mcc_h17_ = source30_.cf62
                d_1991___mcc_h18_ = source30_.cf63
                d_1992___mcc_h19_ = source30_.cf64
                d_1993___mcc_h20_ = source30_.cf65
                d_1994_cf65_ = d_1993___mcc_h20_
                d_1995_cf64_ = d_1992___mcc_h19_
                d_1996_cf63_ = d_1991___mcc_h18_
                d_1997_cf62_ = d_1990___mcc_h17_
                return D15_DC34()
            elif True:
                d_1998___mcc_h21_ = source30_.cf61
                d_1999_cf61_ = d_1998___mcc_h21_
                return D15_DC34()

        source29_ = lambda73_(d_1989_v109_)
        if source29_.is_DC34:
            (globalState).f1 = ((self).f11 if False else ((0) - ((self).f11) if default__.fm0(globalState) else (self).f11))
            d_2000_v110_: T0
            nw329_ = C2()
            nw329_.ctor__(default__.fm20(_dafny.SeqWithoutIsStrInference([(self).f13, not((self).f13), (self).f13, (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]]), 311, (self).f13, globalState), -613)
            d_2000_v110_ = nw329_
            d_2001_v111_: _dafny.Map
            d_2001_v111_ = _dafny.Map({d_2000_v110_: (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]})
            d_2002_v112_: C2
            nw330_ = C2()
            nw330_.ctor__(p0, (self).f11)
            d_2002_v112_ = nw330_
            d_2003_v113_: _dafny.Set
            d_2003_v113_ = _dafny.Set({d_2002_v112_, d_2002_v112_})
            rhs354_ = (d_2001_v111_) | (d_2001_v111_)
            rhs355_ = not((d_1908_v55_)[default__.safeIndex(len((d_2003_v113_) | (d_2003_v113_)), len(d_1908_v55_))])
            rhs356_ = (self).f11
            rhs357_ = default__.safeDivisionInt(((0) - ((self).f11) if (self).f13 else 25), (self).f11)
            lhs300_ = globalState
            lhs301_ = globalState
            lhs302_ = globalState
            d_2001_v111_ = rhs354_
            lhs300_.f6 = rhs355_
            lhs301_.f1 = rhs356_
            lhs302_.f1 = rhs357_
            d_2004_v114_: _dafny.Map
            d_2004_v114_ = _dafny.Map({d_1989_v109_: (d_2000_v110_).f11})
            d_2004_v114_ = (d_2004_v114_).set(d_1989_v109_, -321)
            d_2005_v115_: _dafny.Array
            def lambda74_(d_2006_v110_, d_2007_v37_):
                def lambda75_(d_2008_i6_):
                    return _dafny.Map({_dafny.Set({(d_2006_v110_).f11}): _dafny.Map({(d_2007_v37_)[default__.safeIndex(945, (d_2007_v37_).length(0))]: (d_2007_v37_)[default__.safeIndex(945, (d_2007_v37_).length(0))]})})

                return lambda75_

            init39_ = lambda74_(d_2000_v110_, d_1890_v37_)
            nw331_ = _dafny.Array(None, 28)
            for i0_39_ in range(nw331_.length(0)):
                nw331_[i0_39_] = init39_(i0_39_)
            d_2005_v115_ = nw331_
            d_2009_v116_: _dafny.Set
            d_2009_v116_ = _dafny.Set({(d_2000_v110_).f11, (d_2000_v110_).f11})
            d_2010_v117_: _dafny.Map
            d_2010_v117_ = _dafny.Map({True: True})
            d_2011_v118_: _dafny.Map
            d_2011_v118_ = _dafny.Map({d_2009_v116_: d_2010_v117_})
            index285_ = default__.safeIndex(96, (d_2005_v115_).length(0))
            (d_2005_v115_)[index285_] = d_2011_v118_
            d_2012_v119_: _dafny.Seq
            d_2012_v119_ = _dafny.SeqWithoutIsStrInference([d_1908_v55_, d_1908_v55_, d_1908_v55_])
            d_2013_v120_: D14
            d_2013_v120_ = D14_DC31(False, (self).f11, 397)
            index286_ = default__.safeIndex(96, (d_2005_v115_).length(0))
            rhs358_ = ((default__.fm55((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], len(d_1886_v33_), (self).f13, globalState)) | (d_2011_v118_)).set(d_2009_v116_, d_2010_v117_)
            rhs359_ = (((d_2012_v119_)[default__.safeIndex((self).f11, len(d_2012_v119_))]) + (_dafny.SeqWithoutIsStrInference([(d_2013_v120_).cf45, (self).f13]))) <= (d_1908_v55_)
            lhs303_ = d_2005_v115_
            lhs304_ = default__.safeIndex(96, (d_2005_v115_).length(0))
            lhs305_ = globalState
            lhs303_[lhs304_] = rhs358_
            lhs305_.f6 = rhs359_
        elif source29_.is_DC35:
            d_2014___mcc_h5_ = source29_.cf50
            d_2015___mcc_h6_ = source29_.cf51
            d_2016___mcc_h7_ = source29_.cf52
            d_2017___mcc_h8_ = source29_.cf53
            d_2018___mcc_h9_ = source29_.cf54
            d_2019_cf54_ = d_2018___mcc_h9_
            d_2020_cf53_ = d_2017___mcc_h8_
            d_2021_cf52_ = d_2016___mcc_h7_
            d_2022_cf51_ = d_2015___mcc_h6_
            d_2023_cf50_ = d_2014___mcc_h5_
            d_2024_v121_: C7
            nw332_ = C7()
            nw332_.ctor__((d_1899_v46_).f14, ((d_2022_cf51_).f12) - (_dafny.Set({d_2023_cf50_, (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], False})), (d_2022_cf51_).f11)
            d_2024_v121_ = nw332_
            d_2025_v122_: D2
            d_2025_v122_ = D2_DC4((self).f11, (self).f11, (self).f11)
            source31_ = d_2025_v122_
            if source31_.is_DC4:
                d_2026___mcc_h22_ = source31_.cf6
                d_2027___mcc_h23_ = source31_.cf7
                d_2028___mcc_h24_ = source31_.cf8
                d_2029_cf8_ = d_2028___mcc_h24_
                d_2030_cf7_ = d_2027___mcc_h23_
                d_2031_cf6_ = d_2026___mcc_h22_
                d_2032_v124_: _dafny.Map
                def iife187_():
                    coll79_ = _dafny.Map()
                    compr_79_: int
                    for compr_79_ in (d_1892_v39_).keys.Elements:
                        d_2033_v123_: int = compr_79_
                        if (d_2033_v123_) in (d_1892_v39_):
                            coll79_[(d_2033_v123_) + (d_2029_cf8_)] = (self).f11
                    return _dafny.Map(coll79_)
                d_2032_v124_ = _dafny.Map({len((default__.fm27((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], globalState)) + (d_1884_v31_)): iife187_()
                })
                d_2032_v124_ = (d_2032_v124_).set(d_2030_cf7_, d_1902_v49_)
                d_2034_v125_: C0
                nw333_ = C0()
                nw333_.ctor__(d_1908_v55_, d_2019_cf54_)
                d_2034_v125_ = nw333_
                (globalState).f6 = (d_2023_cf50_) and (d_2023_cf50_)
                d_2035_v126_: _dafny.Map
                d_2035_v126_ = _dafny.Map({p0: d_2021_cf52_})
                d_2036_v128_: _dafny.Set
                def iife188_():
                    coll80_ = _dafny.Map()
                    compr_80_: str
                    for compr_80_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2037_i7_ in range(default__.abs(203))])).Elements:
                        d_2038_v127_: str = compr_80_
                        if (d_2038_v127_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2037_i7_ in range(default__.abs(203))])):
                            coll80_[d_2038_v127_] = (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]
                    return _dafny.Map(coll80_)
                d_2036_v128_ = _dafny.Set({d_2035_v126_, iife188_()
                , d_2035_v126_})
                d_2039_v129_: C6
                nw334_ = C6()
                nw334_.ctor__(d_2025_v122_, (d_2022_cf51_).f12, default__.safeModuloInt(d_2031_cf6_, len(d_2036_v128_)))
                d_2039_v129_ = nw334_
            elif source31_.is_DC5:
                d_2040_v130_: _dafny.Array
                nw335_ = _dafny.Array(int(0), 1)
                d_2040_v130_ = nw335_
                index287_ = default__.safeIndex(111, (d_2040_v130_).length(0))
                (d_2040_v130_)[index287_] = default__.safeDivisionInt(len(d_1886_v33_), 483)
                index288_ = default__.safeIndex(952, (d_2040_v130_).length(0))
                (d_2040_v130_)[index288_] = (0) - ((d_2022_cf51_).f11)
                index289_ = default__.safeIndex(111, (d_2040_v130_).length(0))
                index290_ = default__.safeIndex(952, (d_2040_v130_).length(0))
                rhs360_ = (d_1884_v31_)[default__.safeIndex(d_2019_cf54_, len(d_1884_v31_))]
                rhs361_ = default__.safeDivisionInt(((self).f11) + ((self).f11), ((d_2022_cf51_).f11) * (len(d_1886_v33_)))
                rhs362_ = 518
                lhs306_ = d_2040_v130_
                lhs307_ = default__.safeIndex(111, (d_2040_v130_).length(0))
                lhs308_ = d_2040_v130_
                lhs309_ = default__.safeIndex(952, (d_2040_v130_).length(0))
                lhs306_[lhs307_] = rhs360_
                lhs308_[lhs309_] = rhs361_
                d_2019_cf54_ = rhs362_
                (globalState).f1 = (d_2040_v130_)[default__.safeIndex(111, (d_2040_v130_).length(0))]
                d_2041_v131_: _dafny.Seq
                d_2041_v131_ = _dafny.SeqWithoutIsStrInference([d_2040_v130_, d_2040_v130_])
                d_2042_v132_: _dafny.Array
                nw336_ = _dafny.Array(None, 23)
                nw336_[int(0)] = d_2040_v130_
                nw336_[int(1)] = d_2040_v130_
                nw336_[int(2)] = d_2040_v130_
                nw336_[int(3)] = d_2040_v130_
                nw336_[int(4)] = d_2040_v130_
                nw336_[int(5)] = d_2040_v130_
                nw336_[int(6)] = d_2040_v130_
                nw336_[int(7)] = d_2040_v130_
                nw336_[int(8)] = d_2040_v130_
                nw336_[int(9)] = d_2040_v130_
                nw336_[int(10)] = (d_2041_v131_)[default__.safeIndex((self).f11, len(d_2041_v131_))]
                nw336_[int(11)] = d_2040_v130_
                nw336_[int(12)] = d_2040_v130_
                nw336_[int(13)] = d_2040_v130_
                nw336_[int(14)] = d_2040_v130_
                nw336_[int(15)] = d_2040_v130_
                nw336_[int(16)] = d_2040_v130_
                nw336_[int(17)] = d_2040_v130_
                nw336_[int(18)] = d_2040_v130_
                nw336_[int(19)] = d_2040_v130_
                nw336_[int(20)] = d_2040_v130_
                nw336_[int(21)] = d_2040_v130_
                nw336_[int(22)] = d_2040_v130_
                d_2042_v132_ = nw336_
                d_2043_v133_: _dafny.Array
                d_2043_v133_ = d_2042_v132_
                d_2043_v133_ = d_2043_v133_
                d_2044_v134_: C6
                nw337_ = C6()
                nw337_.ctor__(D2_DC4((self).fm8(default__.fm39((d_2040_v130_)[default__.safeIndex(111, (d_2040_v130_).length(0))], (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], (d_2022_cf51_).f11, globalState), globalState), (d_2040_v130_)[default__.safeIndex(111, (d_2040_v130_).length(0))], (self).f11), (self).f12, len(d_1886_v33_))
                d_2044_v134_ = nw337_
            elif True:
                d_2045___mcc_h25_ = source31_.cf5
                d_2046_cf5_ = d_2045___mcc_h25_
                d_2047_v135_: _dafny.Array
                nw338_ = _dafny.Array(None, 5)
                nw338_[int(0)] = ((d_2022_cf51_).f11) + ((self).f11)
                nw338_[int(1)] = 289
                nw338_[int(2)] = (d_2022_cf51_).f11
                nw338_[int(3)] = -902
                nw338_[int(4)] = (self).f11
                d_2047_v135_ = nw338_
                index291_ = default__.safeIndex(235, (d_2047_v135_).length(0))
                (d_2047_v135_)[index291_] = (self).f11
                index292_ = default__.safeIndex(235, (d_2047_v135_).length(0))
                rhs363_ = d_1890_v37_
                rhs364_ = 976
                rhs365_ = (d_2022_cf51_).f11
                lhs310_ = d_2047_v135_
                lhs311_ = default__.safeIndex(235, (d_2047_v135_).length(0))
                lhs312_ = globalState
                d_1890_v37_ = rhs363_
                lhs310_[lhs311_] = rhs364_
                lhs312_.f1 = rhs365_
                d_2048_v136_: _dafny.MultiSet
                d_2048_v136_ = _dafny.MultiSet([(self).f13, d_2023_cf50_])
                (globalState).f1 = (default__.safeModuloInt((self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxxic")))) if (d_2023_cf50_) not in (d_2048_v136_) else (self).f11)
                d_2021_cf52_ = ((self).f11) >= (d_2019_cf54_)
                d_2049_v137_: _dafny.Map
                d_2049_v137_ = _dafny.Map({not((self).fm4((d_2022_cf51_).f11, _dafny.Map({d_1893_v41_: d_1902_v49_}), (0) - (default__.fm2(-392, (self).f11, globalState)), (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], globalState)): (d_2047_v135_ if d_2023_cf50_ else d_2047_v135_)})
                def iife189_():
                    coll81_ = _dafny.Map()
                    compr_81_: _dafny.MultiSet
                    for compr_81_ in (_dafny.Set({d_1893_v41_})).Elements:
                        d_2050_v138_: _dafny.MultiSet = compr_81_
                        if (d_2050_v138_) in (_dafny.Set({d_1893_v41_})):
                            coll81_[d_2050_v138_] = d_1902_v49_
                    return _dafny.Map(coll81_)
                d_2049_v137_ = (d_2049_v137_).set((self).fm4((d_2046_cf5_).cardinality, iife189_()
                , 977, (self).f13, globalState), d_2047_v135_)
            (globalState).f7 = (self).fm4((self).f11, (_dafny.Map({(d_1893_v41_).set((d_2022_cf51_).f11, default__.abs((d_2022_cf51_).f11)): _dafny.Map({(d_2022_cf51_).f11: (d_2022_cf51_).f11})})) | (d_1903_v50_), (d_2025_v122_).cf7, False, globalState)
            index293_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            (d_1890_v37_)[index293_] = (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]
        elif source29_.is_DC36:
            d_2051___mcc_h10_ = source29_.cf55
            d_2052___mcc_h11_ = source29_.cf56
            d_2053___mcc_h12_ = source29_.cf57
            d_2054___mcc_h13_ = source29_.cf58
            d_2055___mcc_h14_ = source29_.cf59
            d_2056_cf59_ = d_2055___mcc_h14_
            d_2057_cf58_ = d_2054___mcc_h13_
            d_2058_cf57_ = d_2053___mcc_h12_
            d_2059_cf56_ = d_2052___mcc_h11_
            d_2060_cf55_ = d_2051___mcc_h10_
            if d_2057_cf58_:
                (globalState).f7 = default__.fm0(globalState)
                d_1886_v33_ = (d_1886_v33_) + (d_1886_v33_)
                d_2057_cf58_ = d_2059_cf56_
                index294_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                (d_1890_v37_)[index294_] = (d_2057_cf58_ if d_2059_cf56_ else (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))])
                d_2061_v139_: D8
                d_2061_v139_ = D8_DC16(d_2059_cf56_, d_2060_cf55_)
                d_2062_v140_: D8
                d_2062_v140_ = D8_DC17(d_2061_v139_)
                d_2062_v140_ = d_2062_v140_
            elif True:
                d_2063_v141_: _dafny.Array
                nw339_ = _dafny.Array(int(0), 22)
                d_2063_v141_ = nw339_
                index295_ = default__.safeIndex(453, (d_2063_v141_).length(0))
                (d_2063_v141_)[index295_] = (self).f11
                index296_ = default__.safeIndex(453, (d_2063_v141_).length(0))
                (d_2063_v141_)[index296_] = -571
                (globalState).f1 = d_2056_cf59_
                index297_ = default__.safeIndex(945, (d_1890_v37_).length(0))
                (d_1890_v37_)[index297_] = (default__.fm19((self).f13, False, globalState)) <= ((d_1908_v55_) + (d_1908_v55_))
                d_2064_v142_: C2
                nw340_ = C2()
                nw340_.ctor__(p0, (self).f11)
                d_2064_v142_ = nw340_
                d_1893_v41_ = (_dafny.MultiSet([(self).f11, d_2056_cf59_, d_2056_cf59_])) | (d_1893_v41_)
            d_2065_v143_: _dafny.Map
            d_2065_v143_ = _dafny.Map({(self).f13: d_1886_v33_})
            d_2066_v144_: _dafny.Seq
            d_2066_v144_ = _dafny.SeqWithoutIsStrInference([default__.fm22(((d_2065_v143_)[True] if (True) in (d_2065_v143_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gau"))), d_2057_cf58_, globalState), d_1892_v39_])
            d_2067_v145_: D8
            d_2067_v145_ = D8_DC16(d_2058_cf57_, (len(d_2066_v144_) if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))] else (self).f11))
            d_2068_v146_: _dafny.Map
            d_2068_v146_ = _dafny.Map({((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]) not in (d_1901_v48_): d_2057_cf58_})
            d_2069_v147_: T1
            nw341_ = C3()
            nw341_.ctor__(d_2060_cf55_, d_2060_cf55_, (self).f11, (self).f12)
            d_2069_v147_ = nw341_
            d_2070_v148_: _dafny.Set
            d_2070_v148_ = _dafny.Set({d_2069_v147_})
            d_2071_v149_: _dafny.Seq
            d_2071_v149_ = _dafny.SeqWithoutIsStrInference([d_2070_v148_, d_2070_v148_, d_2070_v148_, d_2070_v148_])
            index298_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            index299_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            rhs366_ = (d_1908_v55_)[default__.safeIndex((d_2056_cf59_) * (d_2060_cf55_), len(d_1908_v55_))]
            rhs367_ = (d_1899_v46_).fm3(((self).f11) + ((self).f11), d_2057_cf58_, not((True if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))] else (self).f13)), globalState)
            rhs368_ = ((self).f13 if d_2057_cf58_ else (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))])
            rhs369_ = D8_DC16((self).f13, (self).f11)
            rhs370_ = ((d_2068_v146_)[d_2058_cf57_] if (d_2058_cf57_) in (d_2068_v146_) else (d_2071_v149_) < (d_2071_v149_))
            lhs313_ = d_1890_v37_
            lhs314_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            lhs315_ = globalState
            lhs316_ = d_1890_v37_
            lhs317_ = default__.safeIndex(945, (d_1890_v37_).length(0))
            lhs313_[lhs314_] = rhs366_
            lhs315_.f1 = rhs367_
            lhs316_[lhs317_] = rhs368_
            d_2067_v145_ = rhs369_
            d_2058_cf57_ = rhs370_
            d_2072_v150_: _dafny.Seq
            d_2072_v150_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2068_v146_])])
            d_2073_v151_: _dafny.Map
            d_2073_v151_ = _dafny.Map({d_1886_v33_: (d_2072_v150_)[default__.safeIndex((self).f11, len(d_2072_v150_))]})
            d_2074_v152_: _dafny.MultiSet
            d_2074_v152_ = _dafny.MultiSet([d_2068_v146_, d_2068_v146_])
            d_2073_v151_ = (d_2073_v151_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fq")), (_dafny.MultiSet([_dafny.Map({d_2058_cf57_: (self).f13})]) if d_2059_cf56_ else d_2074_v152_))
            d_2060_cf55_ = (d_2069_v147_).fm3(len((d_2069_v147_).f12), d_2059_cf56_, (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], globalState)
        elif source29_.is_DC33:
            d_2075___mcc_h15_ = source29_.cf49
            d_2076_cf49_ = d_2075___mcc_h15_
            d_1901_v48_ = (d_1901_v48_).set(False, (self).f11)
            if ((self).f11) >= ((self).f11):
                d_2077_v153_: D21
                d_2077_v153_ = D21_DC48((self).f13)
                d_2078_v154_: _dafny.Map
                def iife190_(_pat_let54_0):
                    def iife191_(d_2079_dt__update__tmp_h5_):
                        def iife192_(_pat_let55_0):
                            def iife193_(d_2080_dt__update_hcf78_h0_):
                                return D21_DC48(d_2080_dt__update_hcf78_h0_)
                            return iife193_(_pat_let55_0)
                        return iife192_((self).f13)
                    return iife191_(_pat_let54_0)
                d_2078_v154_ = _dafny.Map({iife190_(d_2077_v153_): (self).f13})
                (globalState).f1 = default__.safeModuloInt((len(d_2078_v154_)) * (-913), ((self).f11) + (391))
                d_1902_v49_ = (d_1902_v49_) | ((d_1902_v49_) | (d_1902_v49_))
                d_1909_v56_ = d_1909_v56_
                (globalState).f1 = (self).f11
                d_2081_v155_: int
                d_2082_v156_: bool
                out35_: int
                out36_: bool
                out35_, out36_ = (self).m2(default__.safeModuloInt(222, (self).f11), globalState)
                d_2081_v155_ = out35_
                d_2082_v156_ = out36_
            elif True:
                (globalState).f1 = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoj"))) + (d_1886_v33_))
                (globalState).f1 = 328
                (globalState).f1 = 632
                d_2083_v157_: _dafny.Array
                nw342_ = _dafny.Array(_dafny.MultiSet({}), 5)
                d_2083_v157_ = nw342_
                d_2084_v158_: _dafny.MultiSet
                d_2084_v158_ = _dafny.MultiSet([(self).f13])
                index300_ = default__.safeIndex(356, (d_2083_v157_).length(0))
                (d_2083_v157_)[index300_] = d_2084_v158_
                index301_ = default__.safeIndex(356, (d_2083_v157_).length(0))
                (d_2083_v157_)[index301_] = d_2084_v158_
                d_2085_v159_: _dafny.Array
                nw343_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_2085_v159_ = nw343_
                d_2086_v160_: _dafny.Array
                nw344_ = _dafny.Array(None, 9)
                nw344_[int(0)] = (self).f11
                nw344_[int(1)] = 283
                nw344_[int(2)] = (self).f11
                nw344_[int(3)] = (self).f11
                nw344_[int(4)] = (self).f11
                nw344_[int(5)] = (self).f11
                nw344_[int(6)] = (self).f11
                nw344_[int(7)] = (self).f11
                nw344_[int(8)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlq")))
                d_2086_v160_ = nw344_
                d_2087_v161_: _dafny.Map
                d_2087_v161_ = _dafny.Map({d_2085_v159_: d_2086_v160_})
                d_2088_v162_: D11
                d_2088_v162_ = D11_DC22(d_2087_v161_)
                d_2089_v163_: _dafny.MultiSet
                d_2089_v163_ = _dafny.MultiSet([d_2088_v162_])
                d_2090_v164_: D7
                d_2090_v164_ = D7_DC13((self).f12, (d_2089_v163_).cardinality, (self).f11)
                d_2091_v165_: C7
                nw345_ = C7()
                nw345_.ctor__((d_1899_v46_).f14, (d_2090_v164_).cf18, (self).f11)
                d_2091_v165_ = nw345_
            d_2092_v166_: _dafny.Map
            d_2092_v166_ = _dafny.Map({319: d_1892_v39_})
            d_2093_v168_: _dafny.Seq
            def iife194_():
                coll82_ = _dafny.Map()
                compr_82_: int
                for compr_82_ in (d_1884_v31_).Elements:
                    d_2094_v167_: int = compr_82_
                    if (d_2094_v167_) in (d_1884_v31_):
                        coll82_[(d_2094_v167_) + ((self).f11)] = d_1892_v39_
                return _dafny.Map(coll82_)
            d_2093_v168_ = _dafny.SeqWithoutIsStrInference([d_2092_v166_, iife194_()
            ])
            d_2092_v166_ = (d_2092_v166_) | ((d_2093_v168_)[default__.safeIndex((self).f11, len(d_2093_v168_))])
            d_2095_v169_: _dafny.Seq
            d_2095_v169_ = _dafny.SeqWithoutIsStrInference([d_1890_v37_, d_1890_v37_, d_1887_v34_])
            rhs371_ = d_2095_v169_
            rhs372_ = (self).f13
            lhs318_ = globalState
            d_2095_v169_ = rhs371_
            lhs318_.f7 = rhs372_
        elif True:
            d_2096___mcc_h16_ = source29_.cf60
            d_2097_cf60_ = d_2096___mcc_h16_
            (globalState).f1 = ((self).f11 if (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))] else default__.safeDivisionInt((self).f11, len((d_1886_v33_).set(default__.safeIndex(default__.fm2((self).fm6((d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))], 240, globalState), (self).f11, globalState), len(d_1886_v33_)), p0))))
            (globalState).f1 = len(d_1886_v33_)
            d_2098_v170_: int
            d_2099_v171_: bool
            out37_: int
            out38_: bool
            out37_, out38_ = (self).m2(len((d_1886_v33_).set(default__.safeIndex((self).f11, len(d_1886_v33_)), p0)), globalState)
            d_2098_v170_ = out37_
            d_2099_v171_ = out38_
            d_2100_v172_: _dafny.Map
            d_2100_v172_ = _dafny.Map({p0: (d_1890_v37_)[default__.safeIndex(945, (d_1890_v37_).length(0))]})
            d_2100_v172_ = (_dafny.Map({p0: d_2099_v171_})) | (d_2100_v172_)
        r0 = default__.fm37(globalState)
        return r0

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        d_2101_v0_: _dafny.Array
        def lambda76_(d_2102_i0_):
            return default__.safeDivisionInt(d_2102_i0_, (self).f11)

        init40_ = lambda76_
        nw346_ = _dafny.Array(None, 13)
        for i0_40_ in range(nw346_.length(0)):
            nw346_[i0_40_] = init40_(i0_40_)
        d_2101_v0_ = nw346_
        d_2103_v1_: str
        d_2103_v1_ = _dafny.CodePoint('j')
        index302_ = default__.safeIndex(387, (d_2101_v0_).length(0))
        (d_2101_v0_)[index302_] = (0) - ((self).fm8(d_2103_v1_, globalState))
        d_2104_v2_: _dafny.Seq
        d_2104_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
        d_2105_v3_: _dafny.Seq
        d_2105_v3_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        index303_ = default__.safeIndex(387, (d_2101_v0_).length(0))
        (d_2101_v0_)[index303_] = (len((d_2104_v2_) + (d_2104_v2_))) + (len(d_2105_v3_))
        d_2103_v1_ = d_2103_v1_
        d_2106_v4_: _dafny.MultiSet
        d_2106_v4_ = _dafny.MultiSet([(d_2105_v3_)[default__.safeIndex((self).f11, len(d_2105_v3_))]])
        d_2107_v5_: _dafny.Map
        d_2107_v5_ = _dafny.Map({(0) - (p0): (self).fm7((0) - ((d_2106_v4_).cardinality), (self).f13, d_2106_v4_, globalState)})
        hi5_ = len(_dafny.SeqWithoutIsStrInference([d_2103_v1_ for d_2108_i2_ in range(default__.abs(-346))]))
        for d_2109_i1_ in range(((d_2107_v5_)[(d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]] if ((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]) in (d_2107_v5_) else p0), hi5_):
            d_2110_v6_: _dafny.MultiSet
            d_2110_v6_ = _dafny.MultiSet([d_2109_i1_, (d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]])
            def iife195_():
                coll83_ = _dafny.Map()
                compr_83_: int
                for compr_83_ in _dafny.IntegerRange(163, 568):
                    d_2111_v7_: int = compr_83_
                    if ((163) <= (d_2111_v7_)) and ((d_2111_v7_) < (568)):
                        coll83_[(d_2111_v7_) * (252)] = (d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]
                return _dafny.Map(coll83_)
            (globalState).f1 = ((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]) + ((0) - (((self).f11) * ((self).fm3(-774, (self).fm4(d_2109_i1_, _dafny.Map({d_2110_v6_: iife195_()
            }), (self).f11, (self).f13, globalState), (self).f13, globalState))))
            (globalState).f1 = -211
            d_2112_v8_: _dafny.Map
            d_2112_v8_ = _dafny.Map({(self).f13: d_2103_v1_})
            if not ((d_2112_v8_) == (d_2112_v8_)) or (((self).f12).issubset((self).f12)):
                d_2113_v9_: C1
                nw347_ = C1()
                nw347_.ctor__((self).f13, (self).f13)
                d_2113_v9_ = nw347_
                d_2114_v10_: _dafny.Map
                d_2114_v10_ = _dafny.Map({d_2113_v9_: (d_2113_v9_).f18})
                d_2114_v10_ = (d_2114_v10_).set(d_2113_v9_, (self).f13)
                d_2115_v11_: _dafny.Array
                nw348_ = _dafny.Array(None, 20)
                d_2115_v11_ = nw348_
                d_2115_v11_ = d_2115_v11_
                d_2116_v12_: C2
                nw349_ = C2()
                nw349_.ctor__(d_2103_v1_, ((0) - (p0)) - ((0) - (p0)))
                d_2116_v12_ = nw349_
                d_2117_v13_: _dafny.Map
                d_2117_v13_ = _dafny.Map({(d_2113_v9_).f18: d_2110_v6_})
                d_2118_v14_: _dafny.Map
                d_2118_v14_ = _dafny.Map({len(d_2117_v13_): not((d_2113_v9_).f19)})
                def iife196_():
                    coll84_ = _dafny.Map()
                    compr_84_: int
                    for compr_84_ in _dafny.IntegerRange(923, -68):
                        d_2119_v15_: int = compr_84_
                        if ((923) <= (d_2119_v15_)) and ((d_2119_v15_) < (-68)):
                            coll84_[default__.safeModuloInt(d_2119_v15_, p0)] = False
                    return _dafny.Map(coll84_)
                rhs373_ = False
                rhs374_ = (d_2118_v14_) | (iife196_()
                )
                lhs319_ = globalState
                lhs319_.f7 = rhs373_
                d_2118_v14_ = rhs374_
                d_2120_v16_: _dafny.Seq
                d_2120_v16_ = _dafny.SeqWithoutIsStrInference([d_2107_v5_, d_2107_v5_])
                d_2121_v17_: _dafny.Set
                d_2121_v17_ = _dafny.Set({p0, (self).f11})
                d_2122_v18_: _dafny.Map
                d_2122_v18_ = _dafny.Map({(d_2120_v16_)[default__.safeIndex(len(d_2121_v17_), len(d_2120_v16_))]: (d_2116_v12_).f24})
                d_2122_v18_ = d_2122_v18_
            elif True:
                (globalState).f6 = (d_2104_v2_) != (((d_2104_v2_).set(default__.safeIndex(d_2109_i1_, len(d_2104_v2_)), d_2103_v1_)) + (d_2104_v2_))
                d_2123_v19_: _dafny.Array
                def lambda77_(d_2124_i3_):
                    return ((self).f13) and (not((self).f13))

                init41_ = lambda77_
                nw350_ = _dafny.Array(None, 18)
                for i0_41_ in range(nw350_.length(0)):
                    nw350_[i0_41_] = init41_(i0_41_)
                d_2123_v19_ = nw350_
                d_2123_v19_ = d_2123_v19_
                (globalState).f6 = ((self).f12).ispropersubset(((self).f12).intersection(_dafny.Set({(self).f13, (self).f13})))
                d_2104_v2_ = _dafny.SeqWithoutIsStrInference([d_2103_v1_ for d_2125_i4_ in range(default__.abs(175))])
                (globalState).f6 = not ((self).f13) or ((self).f13)
            d_2126_v20_: _dafny.Map
            d_2126_v20_ = _dafny.Map({d_2103_v1_: True})
            d_2127_v21_: _dafny.Seq
            d_2127_v21_ = _dafny.SeqWithoutIsStrInference([d_2126_v20_])
            rhs375_ = (self).f13
            rhs376_ = not(not ((self).f13) or ((self).f13))
            rhs377_ = d_2127_v21_
            rhs378_ = ((self).f13) == (not(((self).f13 if (self).f13 else (self).f13)))
            rhs379_ = (((self).f11) - ((0) - (((d_2110_v6_)[len(d_2104_v2_)] if (len(d_2104_v2_)) in (d_2110_v6_) else (self).f11)))) >= ((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))])
            lhs320_ = globalState
            lhs321_ = globalState
            lhs322_ = globalState
            lhs323_ = globalState
            lhs320_.f7 = rhs375_
            lhs321_.f7 = rhs376_
            d_2127_v21_ = rhs377_
            lhs322_.f6 = rhs378_
            lhs323_.f6 = rhs379_
        hi6_ = (self).f11
        for d_2128_i5_ in range((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))], hi6_):
            d_2129_v22_: _dafny.Seq
            d_2129_v22_ = _dafny.SeqWithoutIsStrInference([len(d_2105_v3_)])
            rhs380_ = _dafny.SeqWithoutIsStrInference([(d_2128_i5_) - ((self).f11) for d_2130_i6_ in range(default__.abs(291))])
            rhs381_ = (d_2128_i5_) - (d_2128_i5_)
            lhs324_ = globalState
            d_2129_v22_ = rhs380_
            lhs324_.f1 = rhs381_
            d_2131_v23_: _dafny.Array
            nw351_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_2131_v23_ = nw351_
            index304_ = default__.safeIndex(510, (d_2131_v23_).length(0))
            (d_2131_v23_)[index304_] = d_2101_v0_
            index305_ = default__.safeIndex(510, (d_2131_v23_).length(0))
            (d_2131_v23_)[index305_] = d_2101_v0_
            d_2132_v24_: _dafny.MultiSet
            d_2132_v24_ = _dafny.MultiSet([(self).f11])
            d_2133_v25_: _dafny.Map
            d_2133_v25_ = _dafny.Map({d_2132_v24_: (0) - ((self).f11)})
            d_2133_v25_ = (d_2133_v25_).set(d_2132_v24_, default__.fm2((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))], (self).f11, globalState))
            d_2134_v26_: _dafny.Map
            d_2134_v26_ = _dafny.Map({(0) - (default__.safeModuloInt(d_2128_i5_, len(default__.fm19((self).f13, (self).f13, globalState)))): d_2104_v2_})
            d_2104_v2_ = (((d_2134_v26_)[(0) - ((self).f11)] if ((0) - ((self).f11)) in (d_2134_v26_) else d_2104_v2_)).set(default__.safeIndex(d_2128_i5_, len(((d_2134_v26_)[(0) - ((self).f11)] if ((0) - ((self).f11)) in (d_2134_v26_) else d_2104_v2_))), d_2103_v1_)
        hi7_ = p0
        for d_2135_i7_ in range((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))], hi7_):
            d_2136_v27_: _dafny.Map
            d_2136_v27_ = _dafny.Map({not((self).f13): (d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]})
            if ((d_2105_v3_)[default__.safeIndex(p0, len(d_2105_v3_))]) == ((d_2136_v27_) == (_dafny.Map({(self).f13: (0) - (d_2135_i7_)}))):
                d_2137_v28_: C5
                nw352_ = C5()
                nw352_.ctor__()
                d_2137_v28_ = nw352_
                (globalState).f7 = (self).f13
                d_2138_v29_: _dafny.Array
                nw353_ = _dafny.Array(False, 25)
                d_2138_v29_ = nw353_
                d_2139_v30_: _dafny.Seq
                d_2139_v30_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_2140_v31_: _dafny.Map
                d_2140_v31_ = _dafny.Map({False: (self).f13})
                d_2141_v32_: D6
                d_2141_v32_ = D6_DC10(d_2105_v3_)
                d_2142_v33_: D19
                d_2142_v33_ = D19_DC43(False, _dafny.CodePoint('r'), (self).f13, d_2141_v32_, d_2104_v2_)
                index306_ = default__.safeIndex(387, (d_2101_v0_).length(0))
                index307_ = default__.safeIndex(387, (d_2101_v0_).length(0))
                rhs382_ = d_2138_v29_
                rhs383_ = (d_2139_v30_)[default__.safeIndex(len(d_2140_v31_), len(d_2139_v30_))]
                rhs384_ = (p0 if (d_2142_v33_).cf69 else (self).f11)
                rhs385_ = ((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]) != ((self).f11)
                lhs325_ = d_2101_v0_
                lhs326_ = default__.safeIndex(387, (d_2101_v0_).length(0))
                lhs327_ = d_2101_v0_
                lhs328_ = default__.safeIndex(387, (d_2101_v0_).length(0))
                lhs329_ = globalState
                d_2138_v29_ = rhs382_
                lhs325_[lhs326_] = rhs383_
                lhs327_[lhs328_] = rhs384_
                lhs329_.f7 = rhs385_
                d_2143_v34_: _dafny.Set
                d_2143_v34_ = _dafny.Set({(self).f11, d_2135_i7_})
                d_2144_v35_: _dafny.MultiSet
                d_2144_v35_ = _dafny.MultiSet([d_2143_v34_, d_2143_v34_])
                d_2145_v36_: _dafny.Map
                d_2145_v36_ = _dafny.Map({((d_2144_v35_)[d_2143_v34_] if (d_2143_v34_) in (d_2144_v35_) else p0): (self).f13})
                d_2146_v37_: C3
                nw354_ = C3()
                nw354_.ctor__(((0) - (p0)) - (len(d_2145_v36_)), default__.safeModuloInt((self).f11, (d_2139_v30_)[default__.safeIndex(len(_dafny.Set({_dafny.Map({(self).f13: (self).f13})})), len(d_2139_v30_))]), len(d_2139_v30_), (self).f12)
                d_2146_v37_ = nw354_
                d_2147_v38_: _dafny.Seq
                d_2147_v38_ = _dafny.SeqWithoutIsStrInference([d_2139_v30_, d_2139_v30_, d_2139_v30_, d_2139_v30_])
                d_2148_v39_: _dafny.MultiSet
                d_2148_v39_ = _dafny.MultiSet([len(d_2104_v2_)])
                d_2149_v40_: _dafny.Map
                d_2149_v40_ = _dafny.Map({d_2148_v39_: d_2107_v5_})
                d_2150_v41_: C7
                nw355_ = C7()
                nw355_.ctor__(d_2147_v38_, (self).f12, (0) - ((len(d_2105_v3_) if not((self).fm4(p0, d_2149_v40_, (d_2146_v37_).f22, (self).f13, globalState)) else (0) - ((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]))))
                d_2150_v41_ = nw355_
                rhs386_ = d_2150_v41_
                rhs387_ = d_2105_v3_
                d_2150_v41_ = rhs386_
                d_2105_v3_ = rhs387_
            elif True:
                d_2151_v42_: _dafny.Set
                d_2151_v42_ = _dafny.Set({d_2104_v2_, d_2104_v2_})
                d_2151_v42_ = (d_2151_v42_ if (self).f13 else d_2151_v42_)
                d_2152_v43_: C5
                nw356_ = C5()
                nw356_.ctor__()
                d_2152_v43_ = nw356_
                d_2153_v44_: _dafny.Array
                nw357_ = _dafny.Array(False, 15)
                d_2153_v44_ = nw357_
                d_2154_v45_: _dafny.Map
                d_2154_v45_ = _dafny.Map({d_2103_v1_: (self).f13})
                index308_ = default__.safeIndex(217, (d_2153_v44_).length(0))
                (d_2153_v44_)[index308_] = ((d_2154_v45_)[d_2103_v1_] if (d_2103_v1_) in (d_2154_v45_) else (self).f13)
                index309_ = default__.safeIndex(217, (d_2153_v44_).length(0))
                (d_2153_v44_)[index309_] = (self).f13
                index310_ = default__.safeIndex(217, (d_2153_v44_).length(0))
                (d_2153_v44_)[index310_] = ((self).fm6((self).f13, (self).f11, globalState)) > (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgekcnj"))) + (d_2104_v2_)))
                (globalState).f7 = not(((d_2153_v44_)[default__.safeIndex(217, (d_2153_v44_).length(0))]) or (False))
            index311_ = default__.safeIndex(387, (d_2101_v0_).length(0))
            (d_2101_v0_)[index311_] = ((self).f11) * ((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))])
            index312_ = default__.safeIndex(387, (d_2101_v0_).length(0))
            (d_2101_v0_)[index312_] = (self).f11
            index313_ = default__.safeIndex(387, (d_2101_v0_).length(0))
            (d_2101_v0_)[index313_] = p0
        d_2155_v46_: _dafny.MultiSet
        d_2155_v46_ = _dafny.MultiSet([629])
        d_2156_v48_: _dafny.Map
        def iife197_():
            coll85_ = _dafny.Map()
            compr_85_: int
            for compr_85_ in _dafny.IntegerRange(737, -125):
                d_2157_v47_: int = compr_85_
                if ((737) <= (d_2157_v47_)) and ((d_2157_v47_) < (-125)):
                    coll85_[default__.safeModuloInt(d_2157_v47_, (_dafny.MultiSet([d_2106_v4_])).cardinality)] = len(_dafny.Map({(d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))]: True}))
            return _dafny.Map(coll85_)
        d_2156_v48_ = _dafny.Map({d_2155_v46_: iife197_()
        })
        d_2158_v49_: _dafny.Map
        d_2158_v49_ = _dafny.Map({default__.fm0(globalState): not((self).f13)})
        d_2159_v50_: _dafny.Array
        nw358_ = _dafny.Array(None, 12)
        nw358_[int(0)] = _dafny.Map({(self).f13: (self).fm4((self).f11, d_2156_v48_, 128, True, globalState)})
        nw358_[int(1)] = _dafny.Map({(self).f13: (self).f13})
        nw358_[int(2)] = d_2158_v49_
        nw358_[int(3)] = d_2158_v49_
        nw358_[int(4)] = d_2158_v49_
        nw358_[int(5)] = (d_2158_v49_) | (d_2158_v49_)
        nw358_[int(6)] = (d_2158_v49_) | (default__.fm11(globalState))
        nw358_[int(7)] = default__.fm11(globalState)
        nw358_[int(8)] = d_2158_v49_
        nw358_[int(9)] = d_2158_v49_
        nw358_[int(10)] = d_2158_v49_
        nw358_[int(11)] = default__.fm11(globalState)
        d_2159_v50_ = nw358_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2159_v50_).length(0)):
            d_2160_i8_: int = guard_loop_3_
            if (True) and (((0) <= (d_2160_i8_)) and ((d_2160_i8_) < ((d_2159_v50_).length(0)))):
                (d_2159_v50_)[(d_2160_i8_)] = d_2158_v49_
        d_2161_v51_: _dafny.Array
        def lambda78_(d_2162_v4_):
            def lambda79_(d_2163_i9_):
                return d_2162_v4_

            return lambda79_

        init42_ = lambda78_(d_2106_v4_)
        nw359_ = _dafny.Array(None, 10)
        for i0_42_ in range(nw359_.length(0)):
            nw359_[i0_42_] = init42_(i0_42_)
        d_2161_v51_ = nw359_
        d_2164_v52_: _dafny.Seq
        d_2164_v52_ = _dafny.SeqWithoutIsStrInference([d_2161_v51_, d_2161_v51_, d_2161_v51_])
        r0 = (d_2164_v52_)[default__.safeIndex((d_2101_v0_)[default__.safeIndex(387, (d_2101_v0_).length(0))], len(d_2164_v52_))]
        d_2165_v53_: _dafny.Map
        d_2165_v53_ = _dafny.Map({d_2104_v2_: d_2155_v46_})
        r1 = (d_2165_v53_ if (d_2105_v3_)[default__.safeIndex((self).fm8(d_2103_v1_, globalState), len(d_2105_v3_))] else (d_2165_v53_) | (d_2165_v53_))
        return r0, r1

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2166_v0_: str
        d_2166_v0_ = _dafny.CodePoint('f')
        d_2167_v1_: _dafny.Seq
        d_2167_v1_ = _dafny.SeqWithoutIsStrInference([d_2166_v0_, d_2166_v0_])
        r1 = (d_2166_v0_) in (d_2167_v1_)
        r1 = (self).f13
        r0 = (self).f11
        if True:
            (globalState).f7 = ((self).f13) == ((self).f13)
            d_2168_v2_: _dafny.Seq
            d_2168_v2_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13, False, (self).f13, (self).f13])
            d_2169_v3_: D6
            d_2169_v3_ = D6_DC10(d_2168_v2_)
            d_2170_v4_: C0
            nw360_ = C0()
            nw360_.ctor__((d_2169_v3_).cf13, (self).f11)
            d_2170_v4_ = nw360_
            d_2171_v5_: D23
            d_2171_v5_ = D23_DC54(d_2170_v4_)
            d_2172_v6_: _dafny.Array
            nw361_ = _dafny.Array(None, 13)
            nw361_[int(0)] = d_2170_v4_
            nw361_[int(1)] = d_2170_v4_
            nw361_[int(2)] = d_2170_v4_
            nw361_[int(3)] = d_2170_v4_
            nw361_[int(4)] = d_2170_v4_
            nw361_[int(5)] = d_2170_v4_
            nw361_[int(6)] = d_2170_v4_
            nw361_[int(7)] = d_2170_v4_
            nw361_[int(8)] = d_2170_v4_
            nw361_[int(9)] = d_2170_v4_
            nw361_[int(10)] = (d_2171_v5_).cf91
            nw361_[int(11)] = d_2170_v4_
            nw361_[int(12)] = d_2170_v4_
            d_2172_v6_ = nw361_
            index314_ = default__.safeIndex(569, (d_2172_v6_).length(0))
            (d_2172_v6_)[index314_] = d_2170_v4_
            index315_ = default__.safeIndex(569, (d_2172_v6_).length(0))
            (d_2172_v6_)[index315_] = d_2170_v4_
            d_2173_v7_: D15
            d_2173_v7_ = D15_DC34()
            source32_ = d_2173_v7_
            if source32_.is_DC34:
                d_2174_v8_: _dafny.Map
                d_2174_v8_ = _dafny.Map({d_2166_v0_: (self).f13})
                d_2174_v8_ = ((d_2174_v8_) | (d_2174_v8_) if (self).f13 else (d_2174_v8_) | ((_dafny.Map({d_2166_v0_: (self).f13})).set(d_2166_v0_, (self).f13)))
                d_2175_v9_: _dafny.Set
                d_2175_v9_ = _dafny.Set({d_2166_v0_, d_2166_v0_})
                d_2176_v10_: _dafny.Seq
                d_2176_v10_ = _dafny.SeqWithoutIsStrInference([d_2175_v9_])
                d_2177_v11_: _dafny.MultiSet
                d_2177_v11_ = _dafny.MultiSet([(d_2176_v10_)[default__.safeIndex(621, len(d_2176_v10_))]])
                r1 = (d_2177_v11_).issubset(_dafny.MultiSet(d_2176_v10_))
                d_2166_v0_ = d_2166_v0_
                rhs388_ = (self).f13
                rhs389_ = (0) - ((d_2170_v4_).f21)
                rhs390_ = d_2166_v0_
                rhs391_ = default__.fm2((d_2170_v4_).f21, (0) - ((d_2170_v4_).f21), globalState)
                lhs330_ = globalState
                lhs331_ = globalState
                lhs332_ = globalState
                lhs330_.f6 = rhs388_
                lhs331_.f1 = rhs389_
                d_2166_v0_ = rhs390_
                lhs332_.f1 = rhs391_
            elif source32_.is_DC35:
                d_2178___mcc_h0_ = source32_.cf50
                d_2179___mcc_h1_ = source32_.cf51
                d_2180___mcc_h2_ = source32_.cf52
                d_2181___mcc_h3_ = source32_.cf53
                d_2182___mcc_h4_ = source32_.cf54
                d_2183_cf54_ = d_2182___mcc_h4_
                d_2184_cf53_ = d_2181___mcc_h3_
                d_2185_cf52_ = d_2180___mcc_h2_
                d_2186_cf51_ = d_2179___mcc_h1_
                d_2187_cf50_ = d_2178___mcc_h0_
                d_2188_v12_: _dafny.Array
                nw362_ = _dafny.Array(False, 11)
                d_2188_v12_ = nw362_
                index316_ = default__.safeIndex(292, (d_2188_v12_).length(0))
                (d_2188_v12_)[index316_] = (self).f13
                index317_ = default__.safeIndex(292, (d_2188_v12_).length(0))
                (d_2188_v12_)[index317_] = (self).f13
                rhs392_ = (self).f13
                rhs393_ = 845
                rhs394_ = d_2185_cf52_
                lhs333_ = globalState
                lhs334_ = globalState
                lhs335_ = globalState
                lhs333_.f7 = rhs392_
                lhs334_.f1 = rhs393_
                lhs335_.f7 = rhs394_
                d_2189_v13_: _dafny.Array
                def lambda80_(d_2190_v4_):
                    def lambda81_(d_2191_i0_):
                        return (d_2191_i0_) - ((d_2190_v4_).f21)

                    return lambda81_

                init43_ = lambda80_(d_2170_v4_)
                nw363_ = _dafny.Array(None, 16)
                for i0_43_ in range(nw363_.length(0)):
                    nw363_[i0_43_] = init43_(i0_43_)
                d_2189_v13_ = nw363_
                d_2192_v14_: _dafny.Set
                d_2192_v14_ = _dafny.Set({d_2189_v13_, d_2189_v13_, d_2189_v13_, d_2189_v13_})
                d_2193_v15_: D2
                d_2193_v15_ = D2_DC4(p0, d_2183_cf54_, len(d_2184_cf53_))
                d_2194_v16_: C6
                nw364_ = C6()
                nw364_.ctor__(d_2193_v15_, _dafny.Set({d_2187_cf50_}), (d_2170_v4_).f21)
                d_2194_v16_ = nw364_
                d_2195_v17_: _dafny.Map
                d_2195_v17_ = _dafny.Map({d_2194_v16_: d_2187_cf50_})
                d_2196_v18_: _dafny.Seq
                d_2196_v18_ = _dafny.SeqWithoutIsStrInference([(d_2170_v4_).f21, len((d_2195_v17_).set(d_2194_v16_, (self).f13))])
                d_2197_v19_: _dafny.Seq
                d_2197_v19_ = _dafny.SeqWithoutIsStrInference([d_2196_v18_, (_dafny.SeqWithoutIsStrInference([p0, d_2183_cf54_, (self).f11])).set(default__.safeIndex((d_2170_v4_).f21, len(_dafny.SeqWithoutIsStrInference([p0, d_2183_cf54_, (self).f11]))), (self).f11), _dafny.SeqWithoutIsStrInference([(d_2186_cf51_).f11]), d_2196_v18_, d_2196_v18_])
                d_2198_v20_: C7
                nw365_ = C7()
                nw365_.ctor__(d_2197_v19_, (self).f12, p0)
                d_2198_v20_ = nw365_
                d_2199_v21_: _dafny.Array
                def lambda82_(d_2200_v0_):
                    def lambda83_(d_2201_i1_):
                        return d_2200_v0_

                    return lambda83_

                init44_ = lambda82_(d_2166_v0_)
                nw366_ = _dafny.Array(None, 10)
                for i0_44_ in range(nw366_.length(0)):
                    nw366_[i0_44_] = init44_(i0_44_)
                d_2199_v21_ = nw366_
                index318_ = default__.safeIndex(622, (d_2199_v21_).length(0))
                (d_2199_v21_)[index318_] = d_2166_v0_
                d_2202_v22_: _dafny.Seq
                d_2202_v22_ = _dafny.SeqWithoutIsStrInference([d_2198_v20_])
                index319_ = default__.safeIndex(622, (d_2199_v21_).length(0))
                rhs395_ = d_2185_cf52_
                rhs396_ = (d_2192_v14_).intersection((_dafny.Set({d_2189_v13_})) | (d_2192_v14_))
                rhs397_ = (d_2202_v22_)[default__.safeIndex((d_2170_v4_).f21, len(d_2202_v22_))]
                rhs398_ = (self).f11
                rhs399_ = d_2166_v0_
                lhs336_ = globalState
                lhs337_ = globalState
                lhs338_ = d_2199_v21_
                lhs339_ = default__.safeIndex(622, (d_2199_v21_).length(0))
                lhs336_.f6 = rhs395_
                d_2192_v14_ = rhs396_
                d_2198_v20_ = rhs397_
                lhs337_.f1 = rhs398_
                lhs338_[lhs339_] = rhs399_
                d_2203_v23_: _dafny.Set
                d_2203_v23_ = _dafny.Set({d_2188_v12_})
                d_2204_v24_: _dafny.Array
                nw367_ = _dafny.Array(None, 10)
                nw367_[int(0)] = d_2203_v23_
                nw367_[int(1)] = (d_2203_v23_).intersection(d_2203_v23_)
                nw367_[int(2)] = (d_2203_v23_) | (d_2203_v23_)
                nw367_[int(3)] = d_2203_v23_
                nw367_[int(4)] = d_2203_v23_
                nw367_[int(5)] = (D21_DC47(d_2203_v23_)).cf77
                nw367_[int(6)] = (d_2203_v23_) - (d_2203_v23_)
                nw367_[int(7)] = (d_2203_v23_) - (d_2203_v23_)
                nw367_[int(8)] = d_2203_v23_
                nw367_[int(9)] = d_2203_v23_
                d_2204_v24_ = nw367_
                index320_ = default__.safeIndex(692, (d_2204_v24_).length(0))
                (d_2204_v24_)[index320_] = d_2203_v23_
                index321_ = default__.safeIndex(692, (d_2204_v24_).length(0))
                (d_2204_v24_)[index321_] = ((d_2203_v23_) | (d_2203_v23_)) | (d_2203_v23_)
            elif source32_.is_DC36:
                d_2205___mcc_h5_ = source32_.cf55
                d_2206___mcc_h6_ = source32_.cf56
                d_2207___mcc_h7_ = source32_.cf57
                d_2208___mcc_h8_ = source32_.cf58
                d_2209___mcc_h9_ = source32_.cf59
                d_2210_cf59_ = d_2209___mcc_h9_
                d_2211_cf58_ = d_2208___mcc_h8_
                d_2212_cf57_ = d_2207___mcc_h7_
                d_2213_cf56_ = d_2206___mcc_h6_
                d_2214_cf55_ = d_2205___mcc_h5_
                d_2215_v25_: _dafny.MultiSet
                d_2215_v25_ = _dafny.MultiSet([default__.fm2((d_2170_v4_).f21, (d_2170_v4_).f21, globalState)])
                d_2216_v26_: _dafny.Map
                d_2216_v26_ = _dafny.Map({d_2212_cf57_: d_2210_cf59_})
                d_2217_v27_: _dafny.Set
                d_2217_v27_ = _dafny.Set({(d_2210_cf59_) in (d_2215_v25_), (((d_2216_v26_)[d_2211_cf58_] if (d_2211_cf58_) in (d_2216_v26_) else d_2210_cf59_)) >= (939), (819) != (369), d_2212_cf57_, d_2213_cf56_})
                d_2218_v28_: D7
                d_2218_v28_ = D7_DC13(default__.fm41(globalState), (d_2170_v4_).f21, (d_2170_v4_).f21)
                pat_let_tv54_ = d_2217_v27_
                def iife198_(_pat_let56_0):
                    def iife199_(d_2219_dt__update__tmp_h0_):
                        def iife200_(_pat_let57_0):
                            def iife201_(d_2220_dt__update_hcf18_h0_):
                                return D7_DC13(d_2220_dt__update_hcf18_h0_, (d_2219_dt__update__tmp_h0_).cf19, (d_2219_dt__update__tmp_h0_).cf20)
                            return iife201_(_pat_let57_0)
                        return iife200_(pat_let_tv54_)
                    return iife199_(_pat_let56_0)
                d_2217_v27_ = (iife198_(d_2218_v28_)).cf18
                d_2166_v0_ = d_2166_v0_
                d_2221_v29_: _dafny.Array
                def lambda84_(d_2222_i2_):
                    return (self).f13

                init45_ = lambda84_
                nw368_ = _dafny.Array(None, 18)
                for i0_45_ in range(nw368_.length(0)):
                    nw368_[i0_45_] = init45_(i0_45_)
                d_2221_v29_ = nw368_
                index322_ = default__.safeIndex(395, (d_2221_v29_).length(0))
                (d_2221_v29_)[index322_] = not(d_2212_cf57_)
                index323_ = default__.safeIndex(395, (d_2221_v29_).length(0))
                (d_2221_v29_)[index323_] = d_2212_cf57_
                d_2167_v1_ = (d_2167_v1_) + (d_2167_v1_)
            elif source32_.is_DC33:
                d_2223___mcc_h10_ = source32_.cf49
                d_2224_cf49_ = d_2223___mcc_h10_
                d_2225_v30_: _dafny.Map
                d_2225_v30_ = _dafny.Map({(0) - ((d_2170_v4_).f21): (d_2170_v4_).f21})
                d_2225_v30_ = (d_2225_v30_).set(574, (self).f11)
                (globalState).f7 = (self).f13
                index324_ = default__.safeIndex(569, (d_2172_v6_).length(0))
                (d_2172_v6_)[index324_] = d_2170_v4_
                d_2226_v31_: _dafny.Array
                nw369_ = _dafny.Array(False, 12)
                d_2226_v31_ = nw369_
                d_2227_v32_: D2
                d_2227_v32_ = D2_DC4((d_2170_v4_).f21, -724, (d_2170_v4_).f21)
                d_2228_v33_: C6
                nw370_ = C6()
                nw370_.ctor__(d_2227_v32_, (self).f12, (d_2170_v4_).f21)
                d_2228_v33_ = nw370_
                d_2229_v34_: _dafny.MultiSet
                d_2229_v34_ = _dafny.MultiSet([d_2228_v33_, d_2228_v33_, d_2228_v33_])
                d_2230_v35_: _dafny.Set
                d_2230_v35_ = _dafny.Set({d_2229_v34_})
                index325_ = default__.safeIndex(79, (d_2226_v31_).length(0))
                (d_2226_v31_)[index325_] = (d_2230_v35_).ispropersubset(d_2230_v35_)
                index326_ = default__.safeIndex(79, (d_2226_v31_).length(0))
                (d_2226_v31_)[index326_] = False
            elif True:
                d_2231___mcc_h11_ = source32_.cf60
                d_2232_cf60_ = d_2231___mcc_h11_
                r1 = (self).f13
                (globalState).f7 = True
                (globalState).f1 = (d_2170_v4_).f21
                d_2233_v36_: D11
                d_2233_v36_ = D11_DC24(994, (d_2170_v4_).f20, d_2166_v0_)
                d_2234_v37_: _dafny.MultiSet
                d_2234_v37_ = _dafny.MultiSet([(self).f13, (self).f13, False])
                d_2235_v38_: _dafny.Seq
                d_2235_v38_ = _dafny.SeqWithoutIsStrInference([(self).fm7(p0, (self).f13, d_2234_v37_, globalState)])
                d_2236_v39_: D15
                d_2236_v39_ = D15_DC36((self).f11, (self).f13, (self).f13, (self).f13, 369)
                d_2237_v40_: _dafny.Map
                d_2237_v40_ = _dafny.Map({d_2166_v0_: (self).f13})
                d_2238_v41_: _dafny.MultiSet
                d_2238_v41_ = _dafny.MultiSet([(self).f11])
                d_2239_v42_: _dafny.Map
                d_2239_v42_ = _dafny.Map({(d_2238_v41_).cardinality: (self).fm6(False, (d_2170_v4_).f21, globalState)})
                d_2240_v43_: D7
                d_2240_v43_ = D7_DC13(default__.fm41(globalState), (self).f11, (d_2170_v4_).f21)
                d_2241_v44_: _dafny.Array
                nw371_ = _dafny.Array(None, 29)
                nw371_[int(0)] = (d_2233_v36_).cf34
                nw371_[int(1)] = len(_dafny.Set({(self).f13}))
                nw371_[int(2)] = (self).fm8(d_2166_v0_, globalState)
                nw371_[int(3)] = (d_2170_v4_).f21
                nw371_[int(4)] = (self).f11
                nw371_[int(5)] = 822
                nw371_[int(6)] = len(d_2235_v38_)
                nw371_[int(7)] = (self).f11
                nw371_[int(8)] = len((d_2167_v1_) + (d_2167_v1_))
                nw371_[int(9)] = ((0) - ((d_2170_v4_).f21)) - (p0)
                nw371_[int(10)] = ((d_2236_v39_).cf55) * ((self).f11)
                nw371_[int(11)] = (0) - ((len(d_2237_v40_)) + (p0))
                nw371_[int(12)] = len((d_2239_v42_) | (_dafny.Map({(d_2170_v4_).f21: (self).f11})))
                nw371_[int(13)] = p0
                nw371_[int(14)] = p0
                nw371_[int(15)] = (d_2234_v37_).cardinality
                nw371_[int(16)] = ((self).f11 if (self).f13 else (self).f11)
                nw371_[int(17)] = p0
                nw371_[int(18)] = p0
                nw371_[int(19)] = ((self).f11 if (d_2170_v4_).fm21((self).f13, len(d_2235_v38_), globalState) else 8)
                nw371_[int(20)] = -315
                nw371_[int(21)] = (d_2234_v37_).cardinality
                nw371_[int(22)] = (self).f11
                nw371_[int(23)] = (d_2240_v43_).cf20
                nw371_[int(24)] = (self).f11
                nw371_[int(25)] = (d_2170_v4_).f21
                nw371_[int(26)] = (d_2170_v4_).f21
                nw371_[int(27)] = default__.safeDivisionInt((0) - ((d_2170_v4_).f21), (d_2170_v4_).f21)
                nw371_[int(28)] = (self).f11
                d_2241_v44_ = nw371_
                index327_ = default__.safeIndex(102, (d_2241_v44_).length(0))
                (d_2241_v44_)[index327_] = p0
                index328_ = default__.safeIndex(102, (d_2241_v44_).length(0))
                (d_2241_v44_)[index328_] = (self).fm6((self).f13, (d_2170_v4_).f21, globalState)
            d_2242_v45_: _dafny.Array
            def lambda85_(d_2243_v1_):
                def lambda86_(d_2244_i3_):
                    return d_2243_v1_

                return lambda86_

            init46_ = lambda85_(d_2167_v1_)
            nw372_ = _dafny.Array(None, 10)
            for i0_46_ in range(nw372_.length(0)):
                nw372_[i0_46_] = init46_(i0_46_)
            d_2242_v45_ = nw372_
            index329_ = default__.safeIndex(806, (d_2242_v45_).length(0))
            (d_2242_v45_)[index329_] = d_2167_v1_
            index330_ = default__.safeIndex(806, (d_2242_v45_).length(0))
            (d_2242_v45_)[index330_] = (d_2167_v1_).set(default__.safeIndex(p0, len(d_2167_v1_)), d_2166_v0_)
            d_2245_v46_: _dafny.Seq
            d_2245_v46_ = _dafny.SeqWithoutIsStrInference([(d_2170_v4_).f21])
            d_2246_v47_: _dafny.Map
            d_2246_v47_ = _dafny.Map({d_2166_v0_: d_2245_v46_})
            d_2247_v48_: _dafny.Seq
            d_2247_v48_ = _dafny.SeqWithoutIsStrInference([d_2245_v46_, ((d_2246_v47_)[_dafny.CodePoint('f')] if (_dafny.CodePoint('f')) in (d_2246_v47_) else d_2245_v46_)])
            d_2248_v49_: C7
            nw373_ = C7()
            nw373_.ctor__(d_2247_v48_, (self).f12, (d_2170_v4_).f21)
            d_2248_v49_ = nw373_
        elif True:
            d_2249_v50_: _dafny.Seq
            d_2249_v50_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13, not((self).f13)])
            d_2250_v51_: _dafny.MultiSet
            d_2250_v51_ = _dafny.MultiSet([d_2249_v50_, d_2249_v50_, d_2249_v50_, d_2249_v50_, d_2249_v50_])
            (globalState).f1 = default__.safeDivisionInt(default__.safeModuloInt(len(d_2167_v1_), (self).f11), (d_2250_v51_).cardinality)
            d_2251_v52_: D2
            d_2251_v52_ = D2_DC4(p0, p0, p0)
            d_2252_v53_: C6
            nw374_ = C6()
            def iife202_(_pat_let58_0):
                def iife203_(d_2253_dt__update__tmp_h1_):
                    def iife204_(_pat_let59_0):
                        def iife205_(d_2254_dt__update_hcf8_h0_):
                            return D2_DC4((d_2253_dt__update__tmp_h1_).cf6, (d_2253_dt__update__tmp_h1_).cf7, d_2254_dt__update_hcf8_h0_)
                        return iife205_(_pat_let59_0)
                    return iife204_((self).f11)
                return iife203_(_pat_let58_0)
            nw374_.ctor__(iife202_(d_2251_v52_), default__.fm41(globalState), p0)
            d_2252_v53_ = nw374_
            d_2255_v54_: D13
            d_2255_v54_ = D13_DC28(d_2167_v1_, p0, (self).f13)
            r0 = ((d_2255_v54_).cf41) - ((self).f11)
            d_2256_v55_: _dafny.Array
            nw375_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_2256_v55_ = nw375_
            d_2257_v56_: _dafny.Array
            nw376_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_2257_v56_ = nw376_
            index331_ = default__.safeIndex(633, (d_2256_v55_).length(0))
            (d_2256_v55_)[index331_] = d_2257_v56_
            index332_ = default__.safeIndex(633, (d_2256_v55_).length(0))
            nw377_ = _dafny.Array(_dafny.Array(None, 0), 8)
            (d_2256_v55_)[index332_] = nw377_
            if (self).f13:
                d_2258_v57_: _dafny.Array
                def lambda87_(d_2259_v0_):
                    def lambda88_(d_2260_i4_):
                        return d_2259_v0_

                    return lambda88_

                init47_ = lambda87_(d_2166_v0_)
                nw378_ = _dafny.Array(None, 3)
                for i0_47_ in range(nw378_.length(0)):
                    nw378_[i0_47_] = init47_(i0_47_)
                d_2258_v57_ = nw378_
                d_2261_v58_: _dafny.Seq
                d_2261_v58_ = _dafny.SeqWithoutIsStrInference([d_2258_v57_, d_2258_v57_, d_2258_v57_])
                d_2262_v59_: _dafny.Map
                d_2262_v59_ = _dafny.Map({(self).f13: d_2258_v57_})
                d_2263_v60_: _dafny.Array
                nw379_ = _dafny.Array(None, 20)
                nw379_[int(0)] = d_2258_v57_
                nw379_[int(1)] = d_2258_v57_
                nw379_[int(2)] = (d_2261_v58_)[default__.safeIndex(p0, len(d_2261_v58_))]
                nw379_[int(3)] = d_2258_v57_
                nw379_[int(4)] = d_2258_v57_
                nw379_[int(5)] = d_2258_v57_
                nw379_[int(6)] = d_2258_v57_
                nw379_[int(7)] = d_2258_v57_
                nw379_[int(8)] = d_2258_v57_
                nw379_[int(9)] = d_2258_v57_
                nw379_[int(10)] = ((d_2262_v59_)[(self).f13] if ((self).f13) in (d_2262_v59_) else (d_2261_v58_)[default__.safeIndex(p0, len(d_2261_v58_))])
                nw379_[int(11)] = d_2258_v57_
                nw379_[int(12)] = d_2258_v57_
                nw379_[int(13)] = d_2258_v57_
                nw379_[int(14)] = (d_2258_v57_ if (self).f13 else d_2258_v57_)
                nw379_[int(15)] = d_2258_v57_
                nw379_[int(16)] = d_2258_v57_
                nw379_[int(17)] = d_2258_v57_
                nw379_[int(18)] = d_2258_v57_
                nw379_[int(19)] = d_2258_v57_
                d_2263_v60_ = nw379_
                index333_ = default__.safeIndex(247, (d_2263_v60_).length(0))
                (d_2263_v60_)[index333_] = d_2258_v57_
                index334_ = default__.safeIndex(247, (d_2263_v60_).length(0))
                (d_2263_v60_)[index334_] = d_2258_v57_
                d_2264_v61_: C2
                nw380_ = C2()
                nw380_.ctor__(_dafny.CodePoint('s'), p0)
                d_2264_v61_ = nw380_
                d_2265_v62_: _dafny.Seq
                d_2265_v62_ = _dafny.SeqWithoutIsStrInference([p0])
                (globalState).f1 = ((p0) - ((d_2265_v62_)[default__.safeIndex((self).f11, len(d_2265_v62_))])) + (default__.safeDivisionInt(p0, len(d_2167_v1_)))
                d_2266_v63_: D7
                d_2266_v63_ = D7_DC12(d_2265_v62_)
                d_2266_v63_ = d_2266_v63_
                d_2267_v64_: _dafny.Array
                nw381_ = _dafny.Array(False, 18)
                d_2267_v64_ = nw381_
                d_2268_v65_: _dafny.Map
                d_2268_v65_ = _dafny.Map({((self).f11) * (len(d_2167_v1_)): d_2267_v64_})
                d_2269_v66_: _dafny.Set
                d_2269_v66_ = _dafny.Set({(self).f11, (self).f11})
                d_2270_v67_: _dafny.MultiSet
                d_2270_v67_ = _dafny.MultiSet([len(d_2269_v66_)])
                rhs400_ = len(d_2268_v65_)
                rhs401_ = not(not((self).f13))
                rhs402_ = (d_2270_v67_).issubset((d_2270_v67_).intersection(_dafny.MultiSet(d_2265_v62_)))
                rhs403_ = default__.safeModuloInt(p0, (self).f11)
                rhs404_ = (self).f13
                lhs340_ = globalState
                lhs341_ = globalState
                lhs342_ = globalState
                lhs343_ = globalState
                r0 = rhs400_
                lhs340_.f6 = rhs401_
                lhs341_.f7 = rhs402_
                lhs342_.f1 = rhs403_
                lhs343_.f6 = rhs404_
            elif True:
                d_2271_v68_: _dafny.Array
                nw382_ = _dafny.Array(False, 25)
                d_2271_v68_ = nw382_
                d_2272_v69_: _dafny.Set
                d_2272_v69_ = _dafny.Set({d_2271_v68_})
                d_2273_v70_: T1
                nw383_ = C6()
                nw383_.ctor__((d_2252_v53_).f15, (self).f12, len(d_2272_v69_))
                d_2273_v70_ = nw383_
                d_2274_v71_: _dafny.Map
                d_2274_v71_ = _dafny.Map({d_2273_v70_: p0})
                d_2275_v72_: _dafny.Seq
                d_2275_v72_ = _dafny.SeqWithoutIsStrInference([len(d_2274_v71_), len(d_2167_v1_)])
                d_2276_v73_: _dafny.Seq
                d_2276_v73_ = _dafny.SeqWithoutIsStrInference([d_2275_v72_])
                d_2277_v74_: _dafny.Map
                d_2277_v74_ = _dafny.Map({(d_2273_v70_).f11: (0) - ((self).f11)})
                d_2278_v75_: C7
                nw384_ = C7()
                nw384_.ctor__(d_2276_v73_, (_dafny.Set({(self).f13})).intersection((self).f12), (((d_2277_v74_)[(0) - (len(d_2167_v1_))] if ((0) - (len(d_2167_v1_))) in (d_2277_v74_) else len(d_2167_v1_))) - ((d_2273_v70_).f11))
                d_2278_v75_ = nw384_
                index335_ = default__.safeIndex(355, (d_2271_v68_).length(0))
                (d_2271_v68_)[index335_] = ((self).f13 if (self).f13 else (self).f13)
                index336_ = default__.safeIndex(355, (d_2271_v68_).length(0))
                (d_2271_v68_)[index336_] = ((d_2273_v70_).f12).ispropersubset(((self).f12) | ((self).f12))
                d_2279_v76_: _dafny.MultiSet
                d_2279_v76_ = _dafny.MultiSet([(self).f13, (d_2271_v68_)[default__.safeIndex(355, (d_2271_v68_).length(0))]])
                d_2280_v77_: _dafny.MultiSet
                d_2280_v77_ = d_2279_v76_
                d_2281_v78_: _dafny.Map
                d_2281_v78_ = _dafny.Map({not((d_2271_v68_)[default__.safeIndex(355, (d_2271_v68_).length(0))]): d_2280_v77_})
                d_2282_v79_: D0
                d_2282_v79_ = D0_DC1((self).f13, (d_2273_v70_).f11, (self).f13)
                d_2283_v80_: _dafny.Seq
                d_2283_v80_ = _dafny.SeqWithoutIsStrInference([d_2282_v79_])
                d_2284_v81_: D8
                d_2284_v81_ = D8_DC16((self).f13, len(d_2283_v80_))
                d_2285_v82_: _dafny.Map
                d_2285_v82_ = _dafny.Map({p0: d_2166_v0_})
                d_2286_v83_: _dafny.Map
                d_2286_v83_ = _dafny.Map({(self).f11: (d_2271_v68_)[default__.safeIndex(355, (d_2271_v68_).length(0))]})
                d_2287_v84_: D7
                d_2287_v84_ = D7_DC13((d_2273_v70_).f12, (self).f11, (self).f11)
                d_2288_v85_: _dafny.Map
                d_2288_v85_ = _dafny.Map({(self).f11: d_2287_v84_})
                d_2289_v86_: D7
                d_2289_v86_ = D7_DC14(((d_2288_v85_)[(self).f11] if ((self).f11) in (d_2288_v85_) else d_2287_v84_))
                d_2290_v87_: _dafny.Map
                d_2290_v87_ = _dafny.Map({(self).f13: (self).f11})
                d_2291_v88_: _dafny.Array
                nw385_ = _dafny.Array(None, 29)
                nw385_[int(0)] = (self).f11
                nw385_[int(1)] = p0
                nw385_[int(2)] = p0
                nw385_[int(3)] = (self).f11
                nw385_[int(4)] = len(d_2281_v78_)
                nw385_[int(5)] = (d_2273_v70_).f11
                nw385_[int(6)] = default__.safeDivisionInt(p0, p0)
                nw385_[int(7)] = ((0) - ((d_2273_v70_).f11)) - ((d_2284_v81_).cf24)
                nw385_[int(8)] = len(d_2285_v82_)
                nw385_[int(9)] = (787) * (p0)
                nw385_[int(10)] = p0
                nw385_[int(11)] = (self).f11
                nw385_[int(12)] = (0) - ((self).f11)
                nw385_[int(13)] = (0) - (((0) - (len(d_2286_v83_)) if (self).f13 else p0))
                nw385_[int(14)] = ((d_2273_v70_).f11 if default__.fm0(globalState) else (d_2278_v75_).fm6((self).f13, (d_2273_v70_).f11, globalState))
                nw385_[int(15)] = (self).f11
                nw385_[int(16)] = len(_dafny.Set({d_2289_v86_}))
                nw385_[int(17)] = default__.safeDivisionInt((self).f11, (d_2279_v76_).cardinality)
                nw385_[int(18)] = p0
                nw385_[int(19)] = default__.safeModuloInt(p0, p0)
                nw385_[int(20)] = -888
                nw385_[int(21)] = (self).f11
                nw385_[int(22)] = p0
                nw385_[int(23)] = (d_2273_v70_).f11
                nw385_[int(24)] = (0) - (((self).f11) + ((d_2273_v70_).f11))
                nw385_[int(25)] = ((self).f11) + ((d_2273_v70_).f11)
                nw385_[int(26)] = (len(d_2286_v83_) if (d_2271_v68_)[default__.safeIndex(355, (d_2271_v68_).length(0))] else len(d_2290_v87_))
                nw385_[int(27)] = p0
                nw385_[int(28)] = p0
                d_2291_v88_ = nw385_
                def lambda89_(d_2292_v70_):
                    def lambda90_(d_2293_i5_):
                        return default__.safeDivisionInt(d_2293_i5_, (d_2292_v70_).f11)

                    return lambda90_

                init48_ = lambda89_(d_2273_v70_)
                nw386_ = _dafny.Array(None, 7)
                for i0_48_ in range(nw386_.length(0)):
                    nw386_[i0_48_] = init48_(i0_48_)
                d_2291_v88_ = nw386_
                r1 = (d_2271_v68_)[default__.safeIndex(355, (d_2271_v68_).length(0))]
                def iife206_():
                    coll86_ = _dafny.Set()
                    compr_86_: int
                    for compr_86_ in (d_2286_v83_).keys.Elements:
                        d_2294_v89_: int = compr_86_
                        if (d_2294_v89_) in (d_2286_v83_):
                            coll86_ = coll86_.union(_dafny.Set([(d_2294_v89_) - (611)]))
                    return _dafny.Set(coll86_)
                (globalState).f1 = ((p0) * (714)) * (len(iife206_()
                ))
        d_2295_v90_: _dafny.Seq
        d_2295_v90_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_2296_v91_: _dafny.MultiSet
        d_2296_v91_ = _dafny.MultiSet([p0, p0, (self).f11])
        d_2297_v92_: _dafny.Seq
        out39_: _dafny.Seq
        out39_ = (self).m3((self).f13, -814, (self).f11, (d_2296_v91_).ispropersubset(_dafny.MultiSet(d_2295_v90_)), globalState)
        d_2297_v92_ = out39_
        d_2298_v93_: _dafny.Seq
        d_2298_v93_ = _dafny.SeqWithoutIsStrInference([default__.fm0(globalState)])
        d_2299_v94_: _dafny.Seq
        d_2299_v94_ = _dafny.SeqWithoutIsStrInference([(d_2298_v93_)[default__.safeIndex(p0, len(d_2298_v93_))], (self).f13])
        d_2300_v95_: D0
        d_2300_v95_ = D0_DC0(len((d_2298_v93_).set(default__.safeIndex(p0, len(d_2298_v93_)), not((self).f13))))
        d_2301_v96_: _dafny.MultiSet
        d_2301_v96_ = _dafny.MultiSet([d_2166_v0_, _dafny.CodePoint('u')])
        d_2302_v97_: _dafny.Seq
        out40_: _dafny.Seq
        out40_ = (self).m3((self).fm5(d_2300_v95_, d_2299_v94_, (d_2301_v96_).cardinality, d_2166_v0_, globalState), (self).f11, (self).f11, (self).f13, globalState)
        d_2302_v97_ = out40_
        r0 = default__.safeDivisionInt(p0, (self).f11)
        r1 = (self).f13
        return r0, r1

    def m3(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2303_v0_: _dafny.Seq
        d_2303_v0_ = _dafny.SeqWithoutIsStrInference([True])
        d_2303_v0_ = (d_2303_v0_) + (d_2303_v0_)
        d_2304_v1_: str
        d_2304_v1_ = _dafny.CodePoint('v')
        d_2305_v2_: _dafny.Map
        d_2305_v2_ = _dafny.Map({d_2304_v1_: p0})
        d_2305_v2_ = d_2305_v2_
        d_2306_v3_: _dafny.Seq
        d_2306_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifvvwmi"))
        if (False if not((self).f13) else not((len(d_2306_v3_)) <= ((self).f11))):
            d_2307_v4_: _dafny.Array
            def lambda91_(d_2308_v0_):
                def lambda92_(d_2309_i0_):
                    return default__.safeModuloInt(d_2309_i0_, len(d_2308_v0_))

                return lambda92_

            init49_ = lambda91_(d_2303_v0_)
            nw387_ = _dafny.Array(None, 16)
            for i0_49_ in range(nw387_.length(0)):
                nw387_[i0_49_] = init49_(i0_49_)
            d_2307_v4_ = nw387_
            d_2310_v5_: _dafny.Map
            d_2310_v5_ = _dafny.Map({d_2307_v4_: p0})
            d_2311_v6_: _dafny.Seq
            d_2311_v6_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_2307_v4_: p0})).set(d_2307_v4_, (self).f13)])
            d_2312_v7_: D19
            d_2312_v7_ = D19_DC42(d_2307_v4_)
            d_2313_v8_: _dafny.Array
            nw388_ = _dafny.Array(None, 15)
            nw388_[int(0)] = d_2310_v5_
            nw388_[int(1)] = _dafny.Map({d_2307_v4_: default__.fm0(globalState)})
            nw388_[int(2)] = (d_2310_v5_ if True else d_2310_v5_)
            nw388_[int(3)] = (d_2310_v5_) | (d_2310_v5_)
            nw388_[int(4)] = (d_2310_v5_) | (d_2310_v5_)
            nw388_[int(5)] = ((d_2311_v6_)[default__.safeIndex((self).f11, len(d_2311_v6_))] if (self).f13 else d_2310_v5_)
            nw388_[int(6)] = d_2310_v5_
            nw388_[int(7)] = (_dafny.Map({(d_2312_v7_).cf68: p3})) | (d_2310_v5_)
            nw388_[int(8)] = (_dafny.Map({d_2307_v4_: p0})) | (d_2310_v5_)
            nw388_[int(9)] = d_2310_v5_
            nw388_[int(10)] = d_2310_v5_
            nw388_[int(11)] = d_2310_v5_
            nw388_[int(12)] = _dafny.Map({d_2307_v4_: (self).f13})
            nw388_[int(13)] = d_2310_v5_
            nw388_[int(14)] = (d_2310_v5_) | (d_2310_v5_)
            d_2313_v8_ = nw388_
            index337_ = default__.safeIndex(292, (d_2313_v8_).length(0))
            (d_2313_v8_)[index337_] = d_2310_v5_
            index338_ = default__.safeIndex(292, (d_2313_v8_).length(0))
            (d_2313_v8_)[index338_] = d_2310_v5_
            d_2314_v9_: _dafny.Set
            d_2314_v9_ = _dafny.Set({(self).f11})
            d_2315_v10_: _dafny.Map
            d_2315_v10_ = _dafny.Map({len(d_2314_v9_): p0})
            d_2316_v11_: _dafny.Seq
            d_2316_v11_ = _dafny.SeqWithoutIsStrInference([d_2315_v10_])
            d_2317_v12_: _dafny.Map
            d_2317_v12_ = _dafny.Map({True: len(d_2316_v11_)})
            d_2318_v13_: _dafny.Array
            nw389_ = _dafny.Array(None, 22)
            nw389_[int(0)] = d_2304_v1_
            nw389_[int(1)] = d_2304_v1_
            nw389_[int(2)] = d_2304_v1_
            nw389_[int(3)] = d_2304_v1_
            nw389_[int(4)] = d_2304_v1_
            nw389_[int(5)] = d_2304_v1_
            nw389_[int(6)] = d_2304_v1_
            nw389_[int(7)] = _dafny.CodePoint('y')
            nw389_[int(8)] = d_2304_v1_
            nw389_[int(9)] = d_2304_v1_
            nw389_[int(10)] = d_2304_v1_
            nw389_[int(11)] = _dafny.CodePoint('k')
            nw389_[int(12)] = default__.fm20(d_2303_v0_, 663, p0, globalState)
            nw389_[int(13)] = d_2304_v1_
            nw389_[int(14)] = d_2304_v1_
            nw389_[int(15)] = d_2304_v1_
            nw389_[int(16)] = d_2304_v1_
            nw389_[int(17)] = d_2304_v1_
            nw389_[int(18)] = d_2304_v1_
            nw389_[int(19)] = d_2304_v1_
            nw389_[int(20)] = d_2304_v1_
            nw389_[int(21)] = d_2304_v1_
            d_2318_v13_ = nw389_
            d_2319_v14_: _dafny.Map
            d_2319_v14_ = _dafny.Map({d_2318_v13_: d_2304_v1_})
            d_2320_v15_: _dafny.MultiSet
            d_2320_v15_ = _dafny.MultiSet([((d_2317_v12_)[(self).f13] if ((self).f13) in (d_2317_v12_) else len(d_2319_v14_)), (0) - ((self).f11)])
            d_2321_v16_: _dafny.Map
            d_2321_v16_ = _dafny.Map({d_2320_v15_: _dafny.Map({p1: 673})})
            d_2322_v17_: _dafny.Map
            d_2322_v17_ = _dafny.Map({len(default__.fm1(globalState)): p2})
            (globalState).f7 = not((self).fm4((0) - (p1), (d_2321_v16_).set(d_2320_v15_, _dafny.Map({(self).f11: (_dafny.MultiSet([p2, p2])).cardinality})), ((d_2322_v17_)[(self).f11] if ((self).f11) in (d_2322_v17_) else -482), p3, globalState))
            d_2307_v4_ = ((d_2307_v4_ if p3 else d_2307_v4_) if ((self).f11) > ((self).f11) else d_2307_v4_)
            d_2318_v13_ = d_2318_v13_
            d_2323_v18_: _dafny.Array
            nw390_ = _dafny.Array(_dafny.Map({}), 28)
            d_2323_v18_ = nw390_
            rhs405_ = p1
            rhs406_ = d_2323_v18_
            rhs407_ = p0
            lhs344_ = globalState
            lhs345_ = globalState
            lhs344_.f1 = rhs405_
            d_2323_v18_ = rhs406_
            lhs345_.f7 = rhs407_
        elif True:
            d_2324_v19_: _dafny.MultiSet
            d_2324_v19_ = _dafny.MultiSet([(self).f11])
            d_2325_v21_: _dafny.Map
            def iife207_():
                coll87_ = _dafny.Map()
                compr_87_: int
                for compr_87_ in _dafny.IntegerRange(-821, 806):
                    d_2326_v20_: int = compr_87_
                    if ((-821) <= (d_2326_v20_)) and ((d_2326_v20_) < (806)):
                        coll87_[(d_2326_v20_) + ((self).f11)] = p1
                return _dafny.Map(coll87_)
            d_2325_v21_ = _dafny.Map({d_2324_v19_: iife207_()
            })
            d_2327_v22_: _dafny.Array
            nw391_ = _dafny.Array(False, 19)
            d_2327_v22_ = nw391_
            d_2328_v23_: _dafny.Seq
            d_2328_v23_ = _dafny.SeqWithoutIsStrInference([d_2327_v22_])
            (globalState).f7 = (self).fm4(len(_dafny.SeqWithoutIsStrInference([(0) - (len((self).f12)) for d_2329_i1_ in range(default__.abs(-837))])), d_2325_v21_, default__.safeModuloInt((self).f11, (self).f11), (d_2328_v23_) < (d_2328_v23_), globalState)
            (globalState).f1 = (self).fm8(d_2304_v1_, globalState)
            d_2330_v24_: _dafny.Seq
            d_2330_v24_ = _dafny.SeqWithoutIsStrInference([25])
            d_2330_v24_ = d_2330_v24_
            d_2331_v25_: D6
            d_2331_v25_ = D6_DC11(p2, 647, (self).f11)
            d_2332_v26_: _dafny.Map
            d_2332_v26_ = _dafny.Map({len(default__.fm34(p3, d_2331_v25_, globalState)): 60})
            d_2333_v27_: _dafny.Seq
            d_2333_v27_ = _dafny.SeqWithoutIsStrInference([d_2332_v26_])
            d_2334_v28_: _dafny.Set
            d_2334_v28_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1])) for d_2335_i2_ in range(default__.abs(284))])).set(default__.safeIndex(340, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1])) for d_2336_i2_ in range(default__.abs(284))]))), p2), d_2330_v24_})
            d_2337_v31_: _dafny.Map
            d_2337_v31_ = _dafny.Map({p0: d_2304_v1_})
            d_2338_v32_: _dafny.Map
            d_2338_v32_ = _dafny.Map({p3: ((_dafny.Map({len(d_2306_v3_): len(d_2337_v31_)})).set(p2, (self).f11)).set(p1, (self).f11)})
            d_2339_v33_: _dafny.Array
            nw392_ = _dafny.Array(None, 19)
            nw392_[int(0)] = (d_2333_v27_)[default__.safeIndex(len(d_2303_v0_), len(d_2333_v27_))]
            nw392_[int(1)] = d_2332_v26_
            nw392_[int(2)] = ((d_2332_v26_).set(-527, len(d_2334_v28_)) if p3 else d_2332_v26_)
            nw392_[int(3)] = (d_2332_v26_) | (_dafny.Map({p2: p2}))
            def iife208_():
                coll88_ = _dafny.Map()
                compr_88_: int
                for compr_88_ in _dafny.IntegerRange(760, 727):
                    d_2340_v29_: int = compr_88_
                    if ((760) <= (d_2340_v29_)) and ((d_2340_v29_) < (727)):
                        coll88_[(d_2340_v29_) + ((self).f11)] = 86
                return _dafny.Map(coll88_)
            nw392_[int(4)] = iife208_()
            
            nw392_[int(5)] = (d_2332_v26_).set(((d_2324_v19_)[p2] if (p2) in (d_2324_v19_) else 103), p2)
            def iife209_():
                coll89_ = _dafny.Map()
                compr_89_: int
                for compr_89_ in _dafny.IntegerRange(2, 656):
                    d_2341_v30_: int = compr_89_
                    if ((2) <= (d_2341_v30_)) and ((d_2341_v30_) < (656)):
                        coll89_[default__.safeModuloInt(d_2341_v30_, (self).f11)] = 533
                return _dafny.Map(coll89_)
            nw392_[int(6)] = iife209_()
            
            nw392_[int(7)] = default__.fm46(globalState)
            nw392_[int(8)] = (_dafny.Map({(self).f11: (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2342_i3_ in range(default__.abs(254))])))}) if not(False) else d_2332_v26_)
            nw392_[int(9)] = default__.fm46(globalState)
            nw392_[int(10)] = _dafny.Map({(self).f11: p2})
            nw392_[int(11)] = (d_2332_v26_) | (d_2332_v26_)
            nw392_[int(12)] = d_2332_v26_
            nw392_[int(13)] = _dafny.Map({len(d_2338_v32_): len(_dafny.SeqWithoutIsStrInference([d_2304_v1_ for d_2343_i4_ in range(default__.abs(704))]))})
            nw392_[int(14)] = (_dafny.Map({default__.fm2(p2, 493, globalState): 256})) | (d_2332_v26_)
            nw392_[int(15)] = (d_2332_v26_) | (d_2332_v26_)
            nw392_[int(16)] = d_2332_v26_
            nw392_[int(17)] = d_2332_v26_
            nw392_[int(18)] = d_2332_v26_
            d_2339_v33_ = nw392_
            index339_ = default__.safeIndex(337, (d_2339_v33_).length(0))
            (d_2339_v33_)[index339_] = d_2332_v26_
            index340_ = default__.safeIndex(337, (d_2339_v33_).length(0))
            rhs408_ = (_dafny.Map({p2: 157})) | ((default__.fm46(globalState)) | (d_2332_v26_))
            rhs409_ = (138) - (default__.safeModuloInt(p2, (0) - ((self).f11)))
            lhs346_ = d_2339_v33_
            lhs347_ = default__.safeIndex(337, (d_2339_v33_).length(0))
            lhs348_ = globalState
            lhs346_[lhs347_] = rhs408_
            lhs348_.f1 = rhs409_
            d_2344_v34_: D6
            d_2344_v34_ = D6_DC10(d_2303_v0_)
            d_2345_v35_: D19
            d_2345_v35_ = D19_DC43(p3, d_2304_v1_, p3, d_2344_v34_, d_2306_v3_)
            d_2346_v36_: _dafny.Map
            d_2346_v36_ = _dafny.Map({(self).f11: d_2306_v3_})
            d_2347_v37_: _dafny.Array
            nw393_ = _dafny.Array(None, 16)
            nw393_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nowe")) if p0 else d_2306_v3_)
            nw393_[int(1)] = d_2306_v3_
            nw393_[int(2)] = d_2306_v3_
            nw393_[int(3)] = d_2306_v3_
            nw393_[int(4)] = d_2306_v3_
            nw393_[int(5)] = d_2306_v3_
            nw393_[int(6)] = (d_2306_v3_) + (d_2306_v3_)
            nw393_[int(7)] = d_2306_v3_
            nw393_[int(8)] = (_dafny.SeqWithoutIsStrInference([d_2304_v1_ for d_2348_i5_ in range(default__.abs(115))])) + ((d_2345_v35_).cf73)
            nw393_[int(9)] = (d_2306_v3_) + (d_2306_v3_)
            nw393_[int(10)] = ((d_2346_v36_)[p2] if (p2) in (d_2346_v36_) else d_2306_v3_)
            nw393_[int(11)] = d_2306_v3_
            nw393_[int(12)] = d_2306_v3_
            nw393_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcgs"))
            nw393_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_2304_v1_ for d_2349_i6_ in range(default__.abs(910))])) + (d_2306_v3_)
            nw393_[int(15)] = (d_2306_v3_) + (d_2306_v3_)
            d_2347_v37_ = nw393_
            index341_ = default__.safeIndex(514, (d_2347_v37_).length(0))
            (d_2347_v37_)[index341_] = (d_2306_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhbyei")))
            index342_ = default__.safeIndex(514, (d_2347_v37_).length(0))
            (d_2347_v37_)[index342_] = d_2306_v3_
        if p3:
            d_2350_v38_: _dafny.MultiSet
            d_2350_v38_ = _dafny.MultiSet([p0, False])
            d_2351_v39_: T1
            nw394_ = C3()
            nw394_.ctor__((self).f11, ((d_2350_v38_)[(self).f13] if ((self).f13) in (d_2350_v38_) else p2), (self).f11, _dafny.Set({p3}))
            d_2351_v39_ = nw394_
            d_2352_v40_: _dafny.Map
            d_2352_v40_ = _dafny.Map({d_2351_v39_: (self).f11})
            d_2353_v41_: _dafny.MultiSet
            d_2353_v41_ = _dafny.MultiSet([d_2352_v40_])
            d_2354_v42_: _dafny.MultiSet
            d_2354_v42_ = _dafny.MultiSet([(d_2351_v39_).f11])
            d_2355_v43_: _dafny.Array
            nw395_ = _dafny.Array(None, 22)
            nw395_[int(0)] = ((d_2353_v41_)[d_2352_v40_] if (d_2352_v40_) in (d_2353_v41_) else (d_2351_v39_).f11)
            nw395_[int(1)] = (d_2351_v39_).f11
            nw395_[int(2)] = p2
            nw395_[int(3)] = default__.safeDivisionInt((d_2351_v39_).f11, (self).f11)
            nw395_[int(4)] = (self).f11
            nw395_[int(5)] = (self).f11
            nw395_[int(6)] = len(((d_2306_v3_).set(default__.safeIndex(p1, len(d_2306_v3_)), d_2304_v1_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lysmecy"))))
            nw395_[int(7)] = p1
            nw395_[int(8)] = ((self).fm7((d_2354_v42_).cardinality, (self).f13, d_2350_v38_, globalState)) - ((self).f11)
            nw395_[int(9)] = p2
            nw395_[int(10)] = ((0) - (p1) if (self).f13 else 710)
            nw395_[int(11)] = (d_2351_v39_).f11
            nw395_[int(12)] = default__.safeDivisionInt(-666, (self).f11)
            nw395_[int(13)] = p2
            nw395_[int(14)] = 93
            nw395_[int(15)] = default__.safeModuloInt(p1, len(d_2306_v3_))
            nw395_[int(16)] = -212
            nw395_[int(17)] = (d_2351_v39_).f11
            nw395_[int(18)] = (self).f11
            nw395_[int(19)] = ((d_2351_v39_).f11) * (-92)
            nw395_[int(20)] = (-196) + (p1)
            nw395_[int(21)] = (self).f11
            d_2355_v43_ = nw395_
            d_2355_v43_ = d_2355_v43_
            d_2356_v44_: _dafny.Array
            d_2357_v45_: _dafny.Map
            out41_: _dafny.Array
            out42_: _dafny.Map
            out41_, out42_ = (d_2351_v39_).m1(p1, globalState)
            d_2356_v44_ = out41_
            d_2357_v45_ = out42_
            rhs410_ = -437
            rhs411_ = d_2304_v1_
            lhs349_ = globalState
            lhs349_.f1 = rhs410_
            d_2304_v1_ = rhs411_
            (globalState).f1 = (939) * (p2)
            (globalState).f1 = len(_dafny.SeqWithoutIsStrInference([p3]))
        elif True:
            d_2358_v47_: _dafny.MultiSet
            d_2358_v47_ = _dafny.MultiSet([d_2306_v3_])
            d_2359_v48_: _dafny.Map
            def iife210_():
                coll90_ = _dafny.Map()
                compr_90_: _dafny.Seq
                for compr_90_ in (d_2358_v47_).Elements:
                    d_2360_v46_: _dafny.Seq = compr_90_
                    if (d_2360_v46_) in (d_2358_v47_):
                        coll90_[d_2360_v46_] = False
                return _dafny.Map(coll90_)
            d_2359_v48_ = _dafny.Map({True: len(iife210_()
            )})
            d_2361_v49_: _dafny.Seq
            d_2361_v49_ = _dafny.SeqWithoutIsStrInference([len(d_2359_v48_)])
            (globalState).f1 = ((0) - ((0) - (default__.safeDivisionInt((self).fm8(default__.fm20(d_2303_v0_, len(d_2361_v49_), p0, globalState), globalState), p1)))) * (p1)
            (globalState).f1 = (self).f11
            d_2362_v50_: D2
            d_2362_v50_ = D2_DC4(default__.fm2((self).f11, (0) - (p1), globalState), p1, p2)
            d_2363_v51_: C6
            nw396_ = C6()
            nw396_.ctor__(d_2362_v50_, (default__.fm41(globalState)) | ((self).f12), p1)
            d_2363_v51_ = nw396_
            d_2364_v52_: _dafny.Map
            d_2364_v52_ = _dafny.Map({p0: d_2359_v48_})
            d_2365_v53_: _dafny.Map
            d_2365_v53_ = _dafny.Map({len((_dafny.Map({p3: (d_2359_v48_).set(p0, p1)})) | (d_2364_v52_)): d_2363_v51_})
            d_2363_v51_ = ((d_2365_v53_)[(self).f11] if ((self).f11) in (d_2365_v53_) else d_2363_v51_)
            d_2366_v54_: _dafny.Array
            nw397_ = _dafny.Array(int(0), 10)
            d_2366_v54_ = nw397_
            d_2367_v55_: _dafny.Map
            d_2367_v55_ = _dafny.Map({d_2366_v54_: p0})
            d_2368_v56_: D10
            d_2368_v56_ = D10_DC20(d_2367_v55_)
            pat_let_tv55_ = d_2367_v55_
            pat_let_tv56_ = d_2367_v55_
            def iife212_(_pat_let61_0):
                def iife213_(d_2369_dt__update__tmp_h0_):
                    def iife214_(_pat_let62_0):
                        def iife215_(d_2370_dt__update_hcf27_h0_):
                            return D10_DC20(d_2370_dt__update_hcf27_h0_)
                        return iife215_(_pat_let62_0)
                    return iife214_(pat_let_tv55_)
                return iife213_(_pat_let61_0)
            def iife211_(_pat_let60_0):
                def iife216_(d_2371_dt__update__tmp_h1_):
                    def iife217_(_pat_let63_0):
                        def iife218_(d_2372_dt__update_hcf27_h1_):
                            return D10_DC20(d_2372_dt__update_hcf27_h1_)
                        return iife218_(_pat_let63_0)
                    return iife217_(pat_let_tv56_)
                return iife216_(_pat_let60_0)
            source33_ = iife211_(iife212_(d_2368_v56_))
            if source33_.is_DC21:
                d_2373___mcc_h0_ = source33_.cf28
                d_2374___mcc_h1_ = source33_.cf29
                d_2375_cf29_ = d_2374___mcc_h1_
                d_2376_cf28_ = d_2373___mcc_h0_
                (globalState).f1 = 178
                d_2377_v57_: _dafny.Array
                nw398_ = _dafny.Array(False, 9)
                d_2377_v57_ = nw398_
                index343_ = default__.safeIndex(788, (d_2377_v57_).length(0))
                (d_2377_v57_)[index343_] = (532) < ((self).f11)
                index344_ = default__.safeIndex(788, (d_2377_v57_).length(0))
                (d_2377_v57_)[index344_] = (d_2306_v3_) != ((d_2306_v3_) + (d_2306_v3_))
                (globalState).f6 = d_2376_cf28_
                index345_ = default__.safeIndex(719, (d_2366_v54_).length(0))
                (d_2366_v54_)[index345_] = (self).f11
                d_2378_v58_: _dafny.Array
                nw399_ = _dafny.Array(None, 27)
                nw399_[int(0)] = d_2377_v57_
                nw399_[int(1)] = d_2377_v57_
                nw399_[int(2)] = d_2377_v57_
                nw399_[int(3)] = d_2377_v57_
                nw399_[int(4)] = d_2377_v57_
                nw399_[int(5)] = d_2377_v57_
                nw399_[int(6)] = d_2377_v57_
                nw399_[int(7)] = d_2377_v57_
                nw399_[int(8)] = d_2377_v57_
                nw399_[int(9)] = d_2377_v57_
                nw399_[int(10)] = d_2377_v57_
                nw399_[int(11)] = d_2377_v57_
                nw399_[int(12)] = d_2377_v57_
                nw399_[int(13)] = d_2377_v57_
                nw399_[int(14)] = d_2377_v57_
                nw399_[int(15)] = d_2377_v57_
                nw399_[int(16)] = d_2377_v57_
                nw399_[int(17)] = d_2377_v57_
                nw399_[int(18)] = d_2377_v57_
                nw399_[int(19)] = d_2377_v57_
                nw399_[int(20)] = d_2377_v57_
                nw399_[int(21)] = d_2377_v57_
                nw399_[int(22)] = d_2377_v57_
                nw399_[int(23)] = d_2377_v57_
                nw399_[int(24)] = d_2377_v57_
                nw399_[int(25)] = d_2377_v57_
                nw399_[int(26)] = d_2377_v57_
                d_2378_v58_ = nw399_
                index346_ = default__.safeIndex(719, (d_2366_v54_).length(0))
                rhs412_ = p2
                rhs413_ = d_2378_v58_
                rhs414_ = d_2366_v54_
                lhs350_ = d_2366_v54_
                lhs351_ = default__.safeIndex(719, (d_2366_v54_).length(0))
                lhs350_[lhs351_] = rhs412_
                d_2378_v58_ = rhs413_
                d_2366_v54_ = rhs414_
            elif True:
                d_2379___mcc_h2_ = source33_.cf27
                d_2380_cf27_ = d_2379___mcc_h2_
                d_2381_v59_: _dafny.MultiSet
                d_2381_v59_ = _dafny.MultiSet([(self).f11, (self).f11])
                d_2382_v60_: C4
                nw400_ = C4()
                nw400_.ctor__(d_2381_v59_, (self).f13, (self).f12, 587)
                d_2382_v60_ = nw400_
                d_2383_v61_: _dafny.Seq
                d_2383_v61_ = _dafny.SeqWithoutIsStrInference([d_2382_v60_, d_2382_v60_])
                d_2384_v62_: _dafny.Map
                d_2384_v62_ = _dafny.Map({(len(d_2383_v61_)) * (len(d_2306_v3_)): ((d_2361_v49_)[default__.safeIndex(p1, len(d_2361_v49_))]) + (p1)})
                d_2384_v62_ = (d_2384_v62_).set(p2, p2)
                d_2385_v63_: _dafny.Map
                d_2385_v63_ = _dafny.Map({(d_2382_v60_).f16: d_2384_v62_})
                d_2386_v64_: _dafny.Array
                nw401_ = _dafny.Array(None, 29)
                nw401_[int(0)] = (d_2382_v60_).f17
                nw401_[int(1)] = p3
                nw401_[int(2)] = False
                nw401_[int(3)] = (p2) <= (len(d_2359_v48_))
                nw401_[int(4)] = True
                nw401_[int(5)] = p3
                nw401_[int(6)] = (d_2382_v60_).f17
                nw401_[int(7)] = (d_2382_v60_).f17
                nw401_[int(8)] = not((p2) <= ((self).f11))
                nw401_[int(9)] = (p1) == ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))))))
                nw401_[int(10)] = not((d_2382_v60_).f17)
                nw401_[int(11)] = (d_2306_v3_) <= (d_2306_v3_)
                nw401_[int(12)] = (d_2306_v3_) < ((d_2306_v3_).set(default__.safeIndex(p1, len(d_2306_v3_)), d_2304_v1_))
                nw401_[int(13)] = not (p3) or (p0)
                nw401_[int(14)] = p0
                nw401_[int(15)] = p3
                nw401_[int(16)] = p0
                nw401_[int(17)] = (d_2382_v60_).fm4(p2, d_2385_v63_, p2, p3, globalState)
                nw401_[int(18)] = (p1) != (674)
                nw401_[int(19)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwylrsvr"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rk")))
                nw401_[int(20)] = False
                nw401_[int(21)] = p3
                nw401_[int(22)] = (p3 if p3 else (self).f13)
                nw401_[int(23)] = (d_2382_v60_).f17
                nw401_[int(24)] = ((self).f11) != ((0) - ((self).f11))
                nw401_[int(25)] = (self).f13
                nw401_[int(26)] = p3
                nw401_[int(27)] = ((self).f13) == ((d_2382_v60_).f17)
                nw401_[int(28)] = (self).f13
                d_2386_v64_ = nw401_
                index347_ = default__.safeIndex(795, (d_2386_v64_).length(0))
                (d_2386_v64_)[index347_] = p0
                index348_ = default__.safeIndex(795, (d_2386_v64_).length(0))
                (d_2386_v64_)[index348_] = (p3) or ((d_2382_v60_).f17)
                (globalState).f1 = p2
                d_2387_v65_: T0
                nw402_ = C3()
                nw402_.ctor__(p2, (0) - (p1), p2, (self).f12)
                d_2387_v65_ = nw402_
                d_2387_v65_ = d_2387_v65_
            d_2388_v66_: _dafny.Array
            nw403_ = _dafny.Array(_dafny.Set({}), 13)
            d_2388_v66_ = nw403_
            index349_ = default__.safeIndex(738, (d_2388_v66_).length(0))
            (d_2388_v66_)[index349_] = _dafny.Set({p0, True, p0})
            index350_ = default__.safeIndex(738, (d_2388_v66_).length(0))
            (d_2388_v66_)[index350_] = (self).f12
        d_2304_v1_ = d_2304_v1_
        (globalState).f6 = not(p0)
        r0 = (_dafny.SeqWithoutIsStrInference([d_2304_v1_ for d_2389_i7_ in range(default__.abs(133))])) + (d_2306_v3_)
        return r0

    @property
    def f13(self):
        return self._f13
