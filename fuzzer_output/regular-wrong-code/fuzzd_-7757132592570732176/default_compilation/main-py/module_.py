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
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - (-258), len(_dafny.SeqWithoutIsStrInference([False])), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohtdbkn")))), 898, len(_dafny.Set({474, len(_dafny.SeqWithoutIsStrInference([-238]))}))])) for d_0_i0_ in range(default__.abs(876))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1_i1_ in range(default__.abs(271))])), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajayh"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "li"))), 980])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2_i2_ in range(default__.abs(-351))]))])) + (_dafny.SeqWithoutIsStrInference([-156 for d_3_i3_ in range(default__.abs(897))])))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D9_DC26(_dafny.CodePoint('o'))).cf48]) for d_4_i0_ in range(default__.abs(713))])).Elements:
                d_5_v0_: _dafny.Seq = compr_0_
                if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D9_DC26(_dafny.CodePoint('o'))).cf48]) for d_4_i0_ in range(default__.abs(713))])):
                    coll0_ = coll0_.union(_dafny.Set([d_5_v0_]))
            return _dafny.Set(coll0_)
        return (not((iife0_()
        ) != (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')])})))) == ((True if False else False))

    @staticmethod
    def fm2(globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(209, -442):
                d_6_v0_: int = compr_1_
                if ((209) <= (d_6_v0_)) and ((d_6_v0_) < (-442)):
                    coll1_[(d_6_v0_) + (-890)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
            return _dafny.Map(coll1_)
        return (0) - (((72) - (792)) + (len(iife1_()
        )))

    @staticmethod
    def fm3(globalState):
        return ((_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypcfvnmt"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: _dafny.CodePoint('t')})) for d_7_i1_ in range(default__.abs(11))]))]) for d_8_i0_ in range(default__.abs(85))])): True})})) | (_dafny.Map({False: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfvw"))): False})}))) | ((_dafny.Map({False: _dafny.Map({(_dafny.MultiSet([not(False), True, False, not(False), False])).cardinality: True})})) | (_dafny.Map({True: _dafny.Map({973: True})})))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return _dafny.Set({((D7_DC21(_dafny.MultiSet([False]))).cf39).ispropersubset(_dafny.MultiSet([True]))})

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([153])

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return D0_DC1(433, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_9_i0_ in range(default__.abs(296))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sogpqka"))), default__.safeModuloInt((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([-122])).cardinality for d_10_i1_ in range(default__.abs(632))]))).cardinality, (0) - (len(_dafny.Set({270})))), not(False))

    @staticmethod
    def fm17(globalState):
        return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality, 97, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_11_i0_ in range(default__.abs(945))])), 929, (0) - (len(_dafny.Map({len(_dafny.Set({False, True})): len(_dafny.Set({(_dafny.MultiSet([590])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enf")))}))})))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.Map({len(_dafny.Map({not(True): False})): len(_dafny.Map({True: _dafny.CodePoint('q')}))}))})), 119]))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: False}))) | (_dafny.Map({not(True): False}))

    @staticmethod
    def fm22(p0, globalState):
        if False:
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_12_i0_ in range(default__.abs(-982))])})
        elif True:
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: _dafny.Seq
                for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjpv"))])).Elements:
                    d_13_v0_: _dafny.Seq = compr_2_
                    if (d_13_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjpv"))])):
                        coll2_ = coll2_.union(_dafny.Set([d_13_v0_]))
                return _dafny.Set(coll2_)
            return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctdyvhdge")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrw"))})).intersection(iife2_()
            )

    @staticmethod
    def fm23(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: _dafny.Map
            for compr_3_ in (_dafny.MultiSet([_dafny.Map({540: len(_dafny.SeqWithoutIsStrInference([-219, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrdi"))): _dafny.CodePoint('t')})), 735, 847, 648]))})])).Elements:
                d_14_v0_: _dafny.Map = compr_3_
                if (d_14_v0_) in (_dafny.MultiSet([_dafny.Map({540: len(_dafny.SeqWithoutIsStrInference([-219, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrdi"))): _dafny.CodePoint('t')})), 735, 847, 648]))})])):
                    coll3_ = coll3_.union(_dafny.Set([d_14_v0_]))
            return _dafny.Set(coll3_)
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([847, 96])).Elements:
                d_15_v1_: int = compr_4_
                if (d_15_v1_) in (_dafny.SeqWithoutIsStrInference([847, 96])):
                    coll4_[(d_15_v1_) + (len(_dafny.Map({_dafny.CodePoint('l'): 298})))] = 403
            return _dafny.Map(coll4_)
        return ((_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): 438}), _dafny.Map({140: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))})})).intersection(iife3_()
        )).intersection(_dafny.Set({iife4_()
        , (D4_DC11(_dafny.Map({len(_dafny.Set({False})): 68}))).cf20, _dafny.Map({720: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mabtyye")))})}))

    @staticmethod
    def fm25(p0, p1, globalState):
        source0_ = D17_DC46(D17_DC44(_dafny.Set({D15_DC41(True)})))
        if source0_.is_DC45:
            d_16___mcc_h0_ = source0_.cf69
            d_17___mcc_h1_ = source0_.cf70
            d_18_cf70_ = d_17___mcc_h1_
            d_19_cf69_ = d_16___mcc_h0_
            return _dafny.CodePoint('p')
        elif source0_.is_DC44:
            d_20___mcc_h2_ = source0_.cf68
            d_21_cf68_ = d_20___mcc_h2_
            return _dafny.CodePoint('y')
        elif True:
            d_22___mcc_h3_ = source0_.cf71
            d_23_cf71_ = d_22___mcc_h3_
            return _dafny.CodePoint('h')

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([not(True)])) | (_dafny.MultiSet([True]))) | (((D7_DC21(_dafny.MultiSet([False]))).cf39) | (_dafny.MultiSet([not(True)])))

    @staticmethod
    def fm31(globalState):
        return _dafny.Map({539: _dafny.Map({(D19_DC51(len(_dafny.SeqWithoutIsStrInference([-19, -441])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))))).cf77: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, not(False)]))])).cardinality})})

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        if (_dafny.Set({False, False, False})) != (_dafny.Set({not(False), True, False, True, (D0_DC3(233, True)).cf8})):
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_24_i0_ in range(default__.abs(632))])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_25_i1_ in range(default__.abs(-849))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hen")))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        source1_ = D20_DC53()
        if source1_.is_DC53:
            return D5_DC15(-438, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntyca")))), 976)
        elif source1_.is_DC52:
            d_26___mcc_h0_ = source1_.cf79
            d_27_cf79_ = d_26___mcc_h0_
            return D5_DC15(-615, (_dafny.MultiSet([False, True])).cardinality, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality)
        elif True:
            d_28___mcc_h1_ = source1_.cf80
            d_29_cf80_ = d_28___mcc_h1_
            return D5_DC15(-728, len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xseseylhc"))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgx"))))

    @staticmethod
    def fm38(p0, p1, globalState):
        source2_ = D12_DC35(-704)
        if source2_.is_DC33:
            d_30___mcc_h0_ = source2_.cf58
            d_31_cf58_ = d_30___mcc_h0_
            return (_dafny.Map({886: d_31_cf58_})) | (_dafny.Map({476: -762}))
        elif source2_.is_DC34:
            d_32___mcc_h1_ = source2_.cf59
            d_33_cf59_ = d_32___mcc_h1_
            return _dafny.Map({535: 167})
        elif source2_.is_DC35:
            d_34___mcc_h2_ = source2_.cf60
            d_35_cf60_ = d_34___mcc_h2_
            return _dafny.Map({391: d_35_cf60_})
        elif True:
            d_36___mcc_h3_ = source2_.cf57
            d_37_cf57_ = d_36___mcc_h3_
            if True:
                return _dafny.Map({844: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_38_i0_ in range(default__.abs(138))]))})
            elif True:
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: int
                    for compr_5_ in (_dafny.Map({-112: 862})).keys.Elements:
                        d_39_v0_: int = compr_5_
                        if (d_39_v0_) in (_dafny.Map({-112: 862})):
                            coll5_[default__.safeModuloInt(d_39_v0_, len(_dafny.Map({50: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlsvjvt")))})))] = len(_dafny.Set({False, True}))
                    return _dafny.Map(coll5_)
                return iife5_()
                

    @staticmethod
    def fm39(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm40(p0, globalState):
        return D3_DC8((_dafny.Map({True: 124})) | (_dafny.Map({not(True): 347})))

    @staticmethod
    def fm41(globalState):
        return _dafny.SeqWithoutIsStrInference([D6_DC18(True) for d_40_i0_ in range(default__.abs(351))])

    @staticmethod
    def fm42(p0, p1, globalState):
        return _dafny.Map({not(False): (len(_dafny.SeqWithoutIsStrInference([False]))) == (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_41_i0_ in range(default__.abs(-576))])): not(True)})))})

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(370, 695):
                d_42_v0_: int = compr_6_
                if ((370) <= (d_42_v0_)) and ((d_42_v0_) < (695)):
                    coll6_[(d_42_v0_) + (956)] = D5_DC15(-552, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "polbyirqn"))), -321)
            return _dafny.Map(coll6_)
        return D6_DC16((iife6_()
) | (_dafny.Map({826: D5_DC15(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_43_i0_ in range(default__.abs(288))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_44_i1_ in range(default__.abs(288))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_45_i2_ in range(default__.abs(877))])))})))

    @staticmethod
    def fm44(globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: D6
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([D6_DC18(True), D6_DC18(True), D6_DC18(False), D6_DC18(True), D6_DC18(False)])).Elements:
                d_46_v0_: D6 = compr_7_
                if (d_46_v0_) in (_dafny.SeqWithoutIsStrInference([D6_DC18(True), D6_DC18(True), D6_DC18(False), D6_DC18(True), D6_DC18(False)])):
                    coll7_ = coll7_.union(_dafny.Set([d_46_v0_]))
            return _dafny.Set(coll7_)
        return ((iife7_()
        ) | (_dafny.Set({D6_DC18(True), D6_DC18(False)}))).intersection(_dafny.Set({D6_DC18(False)}))

    @staticmethod
    def fm45(p0, globalState):
        return (D9_DC26(_dafny.CodePoint('k'))).cf48

    @staticmethod
    def fm46(p0, p1, globalState):
        return D4_DC12((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_47_i0_ in range(default__.abs(21))]))).issubset(_dafny.MultiSet([_dafny.CodePoint('p')])), 336, len(_dafny.Map({not(not(not(not(True)))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vul")))})))

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(633, 325):
                d_48_v0_: int = compr_8_
                if ((633) <= (d_48_v0_)) and ((d_48_v0_) < (325)):
                    coll8_[(d_48_v0_) - (389)] = False
            return _dafny.Map(coll8_)
        return (_dafny.Map({len(_dafny.Map({_dafny.CodePoint('t'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "te"))})): True})) | (iife8_()
        )

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([D0_DC3(-618, not(True))])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D0_DC3(len(_dafny.Map({False: len(_dafny.Map({not(not(True)): True}))})), False)])))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D0_DC3(842, not(True)) for d_49_i0_ in range(default__.abs(989))])))

    @staticmethod
    def fm49(globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: _dafny.Set
            for compr_9_ in (_dafny.Map({_dafny.Set({len(_dafny.Map({not(True): 856}))}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss"))})).keys.Elements:
                d_50_v0_: _dafny.Set = compr_9_
                if (d_50_v0_) in (_dafny.Map({_dafny.Set({len(_dafny.Map({not(True): 856}))}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss"))})):
                    coll9_ = coll9_.union(_dafny.Set([d_50_v0_]))
            return _dafny.Set(coll9_)
        return iife9_()
        

    @staticmethod
    def fm50(globalState):
        def iife10_():
            def iife12_():
                coll12_ = _dafny.Set()
                compr_12_: str
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('s')])).Elements:
                    d_54_v1_: str = compr_12_
                    if (d_54_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('s')])):
                        coll12_ = coll12_.union(_dafny.Set([d_54_v1_]))
                return _dafny.Set(coll12_)
            coll10_ = _dafny.Map()
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: str
                for compr_11_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('s')])).Elements:
                    d_52_v1_: str = compr_11_
                    if (d_52_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('s')])):
                        coll11_ = coll11_.union(_dafny.Set([d_52_v1_]))
                return _dafny.Set(coll11_)
            compr_10_: str
            for compr_10_ in (iife11_()
            ).Elements:
                d_53_v0_: str = compr_10_
                if (d_53_v0_) in (iife12_()
                ):
                    def iife13_():
                        coll13_ = _dafny.Map()
                        compr_13_: str
                        for compr_13_ in (_dafny.Map({_dafny.CodePoint('k'): _dafny.CodePoint('d')})).keys.Elements:
                            d_55_v2_: str = compr_13_
                            if (d_55_v2_) in (_dafny.Map({_dafny.CodePoint('k'): _dafny.CodePoint('d')})):
                                coll13_[d_55_v2_] = False
                        return _dafny.Map(coll13_)
                    coll10_[d_53_v0_] = (_dafny.MultiSet([D18_DC47(iife13_()
), D18_DC47(_dafny.Map({_dafny.CodePoint('a'): True}))])).cardinality
            return _dafny.Map(coll10_)
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxvistgq")))}) if not(False) else _dafny.Set({737, -486}))) - ((_dafny.Set({-904, (_dafny.MultiSet([_dafny.Map({True: False}), _dafny.Map({False: False})])).cardinality, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality for d_51_i0_ in range(default__.abs(994))]))).cardinality})).intersection(_dafny.Set({(0) - (len(iife10_()
        )), -460, (0) - (len(_dafny.SeqWithoutIsStrInference([False]))), (0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality]))}))), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_56_i1_ in range(default__.abs(387))])))})))

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([D17_DC44(_dafny.Set({D15_DC41(True)})) for d_57_i0_ in range(default__.abs(765))]))]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hct")))})) | (_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})) for d_58_i1_ in range(default__.abs(298))])): len(_dafny.SeqWithoutIsStrInference([-586 for d_59_i2_ in range(default__.abs(-268))]))}))) | (_dafny.Map({_dafny.MultiSet([-826, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvbhnjfvo"))), 40, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('s')])).cardinality])): True})), 598]): 424}))

    @staticmethod
    def fm52(globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(643, 200):
                d_61_v0_: int = compr_14_
                if ((643) <= (d_61_v0_)) and ((d_61_v0_) < (200)):
                    coll14_[(d_61_v0_) + (675)] = _dafny.Map({len(_dafny.Set({True})): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_62_i1_ in range(default__.abs(665))]))})
            return _dafny.Map(coll14_)
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in (_dafny.Map({-411: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-818: -633})), (0) - (len(_dafny.Map({935: True})))]))]))})).keys.Elements:
                d_63_v1_: int = compr_15_
                if (d_63_v1_) in (_dafny.Map({-411: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-818: -633})), (0) - (len(_dafny.Map({935: True})))]))]))})):
                    coll15_[(d_63_v1_) - (744)] = 567
            return _dafny.Map(coll15_)
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([213, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sokhgatk")))])).Elements:
                d_64_v2_: int = compr_16_
                if (d_64_v2_) in (_dafny.SeqWithoutIsStrInference([213, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sokhgatk")))])):
                    coll16_[default__.safeDivisionInt(d_64_v2_, (_dafny.MultiSet([-745])).cardinality)] = True
            return _dafny.Map(coll16_)
        return (_dafny.MultiSet([D3_DC9((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: -271})), (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), -945, 163])).cardinality, len(_dafny.SeqWithoutIsStrInference([True]))])).cardinality, len(_dafny.Map({False: True})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdnyphax"))) for d_60_i0_ in range(default__.abs(-984))])), 315}))]))])).cardinality, iife14_()
, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaevsmgns"))), (D17_DC45(648, _dafny.Set({(0) - (-451), (0) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({347, (_dafny.MultiSet([_dafny.CodePoint('k'), _dafny.CodePoint('x')])).cardinality})), 217])])).cardinality)}))).cf69), D3_DC9(582, _dafny.Map({len(_dafny.Set({True})): _dafny.Map({517: len(_dafny.Map({_dafny.CodePoint('o'): False}))})}), -58, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False, True]): len(_dafny.Map({False: False}))})))])).intersection((_dafny.MultiSet([D3_DC9((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, True, not(False), False]))])).cardinality, _dafny.Map({(D19_DC51(160, len(_dafny.SeqWithoutIsStrInference([185])))).cf78: _dafny.Map({-564: (D12_DC33(len(_dafny.SeqWithoutIsStrInference([True])))).cf58})}), 434, len(_dafny.Map({_dafny.CodePoint('d'): _dafny.CodePoint('l')}))), D3_DC9(542, _dafny.Map({246: _dafny.Map({len(_dafny.Map({len(_dafny.Map({5: False})): 759})): 769})}), len(_dafny.Map({True: False})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pf"))))])) - (_dafny.MultiSet([D3_DC9(-617, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([-228])): _dafny.Map({337: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjucomr"))))})}), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hckgvkew"))), (_dafny.MultiSet([False])).cardinality), D3_DC9(-585, _dafny.Map({(0) - (len(_dafny.Set({False}))): iife15_()
}), -490, len(_dafny.Map({391: True}))), D3_DC9(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iduro"))), _dafny.Map({-987: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))): len(_dafny.SeqWithoutIsStrInference([False]))})}), 348, len(iife16_()
))])))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D3_DC9(576, _dafny.Map({len(_dafny.Set({False})): _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvqhguysd"))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_65_i0_ in range(default__.abs(993))]))})}), 377, (0) - ((-266 if False else -394)))

    @staticmethod
    def fm54(p0, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in (_dafny.SeqWithoutIsStrInference([-746])).Elements:
                d_67_v0_: int = compr_17_
                if (d_67_v0_) in (_dafny.SeqWithoutIsStrInference([-746])):
                    coll17_[default__.safeModuloInt(d_67_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xflc"))))] = len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbjpcvxsh")): True}))
            return _dafny.Map(coll17_)
        return _dafny.Map({(len(_dafny.Set({False, True, False}))) not in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([150]))): (_dafny.Map({609: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-711 for d_66_i0_ in range(default__.abs(313))])), len(iife17_()
        )]))).cardinality}) if False else _dafny.Map({len(_dafny.Set({D6_DC19(not(True), True, False, True, _dafny.SeqWithoutIsStrInference([not(True)]))})): 666}))})

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): D7_DC21(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): D7_DC21(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(False))])))}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): D7_DC21(_dafny.MultiSet([False]))}))

    @staticmethod
    def fm56(p0, p1, p2, p3, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(974, -465):
                d_69_v0_: int = compr_18_
                if ((974) <= (d_69_v0_)) and ((d_69_v0_) < (-465)):
                    coll18_[default__.safeModuloInt(d_69_v0_, len(_dafny.Set({False, not(not(True))})))] = False
            return _dafny.Map(coll18_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({-61: True}) for d_68_i0_ in range(default__.abs(311))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({788: False}), iife18_()
        ]))

    @staticmethod
    def fm57(p0, p1, globalState):
        def iife19_():
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(574, -763):
                    d_71_v1_: int = compr_20_
                    if ((574) <= (d_71_v1_)) and ((d_71_v1_) < (-763)):
                        coll20_[(d_71_v1_) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})))] = True
                return _dafny.Map(coll20_)
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(134, 571):
                d_70_v0_: int = compr_19_
                if ((134) <= (d_70_v0_)) and ((d_70_v0_) < (571)):
                    coll19_[(d_70_v0_) - (len(iife20_()
                    ))] = False
            return _dafny.Map(coll19_)
        def iife21_():
            def iife23_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(652, 162):
                    d_72_v3_: int = compr_23_
                    if ((652) <= (d_72_v3_)) and ((d_72_v3_) < (162)):
                        coll23_[default__.safeModuloInt(d_72_v3_, 955)] = False
                return _dafny.Map(coll23_)
            coll21_ = _dafny.Map()
            def iife22_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(652, 162):
                    d_72_v3_: int = compr_22_
                    if ((652) <= (d_72_v3_)) and ((d_72_v3_) < (162)):
                        coll22_[default__.safeModuloInt(d_72_v3_, 955)] = False
                return _dafny.Map(coll22_)
            compr_21_: _dafny.Map
            for compr_21_ in (_dafny.Map({iife22_()
            : True})).keys.Elements:
                d_73_v2_: _dafny.Map = compr_21_
                if (d_73_v2_) in (_dafny.Map({iife23_()
                : True})):
                    coll21_[d_73_v2_] = _dafny.SeqWithoutIsStrInference([False])
            return _dafny.Map(coll21_)
        def iife24_():
            def iife26_():
                coll26_ = _dafny.Map()
                compr_26_: int
                for compr_26_ in (_dafny.Set({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([not(True)]))), 729}))})).Elements:
                    d_74_v5_: int = compr_26_
                    if (d_74_v5_) in (_dafny.Set({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([not(True)]))), 729}))})):
                        coll26_[(d_74_v5_) * (len(_dafny.Map({False: len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))}))})))] = True
                return _dafny.Map(coll26_)
            coll24_ = _dafny.Map()
            def iife25_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in (_dafny.Set({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([not(True)]))), 729}))})).Elements:
                    d_74_v5_: int = compr_25_
                    if (d_74_v5_) in (_dafny.Set({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([not(True)]))), 729}))})):
                        coll25_[(d_74_v5_) * (len(_dafny.Map({False: len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))}))})))] = True
                return _dafny.Map(coll25_)
            compr_24_: _dafny.Map
            for compr_24_ in (_dafny.Map({iife25_()
            : len(_dafny.Set({712}))})).keys.Elements:
                d_75_v4_: _dafny.Map = compr_24_
                if (d_75_v4_) in (_dafny.Map({iife26_()
                : len(_dafny.Set({712}))})):
                    coll24_[d_75_v4_] = _dafny.SeqWithoutIsStrInference([not(False)])
            return _dafny.Map(coll24_)
        return ((_dafny.Map({iife19_()
        : _dafny.SeqWithoutIsStrInference([True])})) | (_dafny.Map({_dafny.Map({48: False}): _dafny.SeqWithoutIsStrInference([False])}))) | ((iife21_()
        ) | (iife24_()
        ))

    @staticmethod
    def fm58(p0, globalState):
        def iife27_():
            coll27_ = _dafny.Map()
            compr_27_: int
            for compr_27_ in _dafny.IntegerRange(949, 919):
                d_76_v0_: int = compr_27_
                if ((949) <= (d_76_v0_)) and ((d_76_v0_) < (919)):
                    coll27_[default__.safeDivisionInt(d_76_v0_, len(_dafny.Set({False, True})))] = (_dafny.MultiSet([False])).cardinality
            return _dafny.Map(coll27_)
        return D15_DC41((_dafny.MultiSet([317, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjorlfrb"))), 636])).issubset(_dafny.MultiSet([(0) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jk")): (_dafny.SeqWithoutIsStrInference([D13_DC37(), D13_DC37()]))}))), 880, len(iife27_()
)])))

    @staticmethod
    def fm59(p0, p1, p2, p3, globalState):
        def iife28_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in (_dafny.Set({170, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuifnsqld"))): 197}))})).Elements:
                d_77_v0_: int = compr_28_
                if (d_77_v0_) in (_dafny.Set({170, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuifnsqld"))): 197}))})):
                    coll28_[(d_77_v0_) * (437)] = 242
            return _dafny.Map(coll28_)
        def iife29_():
            coll29_ = _dafny.Map()
            compr_29_: int
            for compr_29_ in _dafny.IntegerRange(746, -949):
                d_78_v1_: int = compr_29_
                if ((746) <= (d_78_v1_)) and ((d_78_v1_) < (-949)):
                    coll29_[default__.safeDivisionInt(d_78_v1_, 915)] = 607
            return _dafny.Map(coll29_)
        return (_dafny.Map({iife28_()
        : 548})) | ((_dafny.Map({iife29_()
        : 449})) | (_dafny.Map({_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([391]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([701]))]))).cardinality: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcxgjmau")): True})))]))).cardinality}): 737})))

    @staticmethod
    def fm60(p0, globalState):
        return ((_dafny.Map({_dafny.CodePoint('y'): not(False)})) | (_dafny.Map({_dafny.CodePoint('k'): True}))) | (_dafny.Map({_dafny.CodePoint('e'): not(False)}))

    @staticmethod
    def fm61(globalState):
        return D15_DC42(D15_DC41(True))

    @staticmethod
    def fm62(p0, p1, p2, p3, globalState):
        if True:
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(D6_DC18(False)).cf32, False]), _dafny.SeqWithoutIsStrInference([False, True, False, True])])
        elif True:
            return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(True)])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, True])]))

    @staticmethod
    def fm63(globalState):
        def iife30_():
            coll30_ = _dafny.Map()
            compr_30_: int
            for compr_30_ in _dafny.IntegerRange(360, 62):
                d_79_v0_: int = compr_30_
                if ((360) <= (d_79_v0_)) and ((d_79_v0_) < (62)):
                    coll30_[(d_79_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihi"))))] = len(_dafny.SeqWithoutIsStrInference([90]))
            return _dafny.Map(coll30_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([642])), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([iife30_()
        ])), 10, 817]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-504]))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([494, len(_dafny.Map({True: 746})), len(_dafny.SeqWithoutIsStrInference([False, True]))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_80_i1_ in range(default__.abs(914))]))]) for d_81_i0_ in range(default__.abs(-829))])))

    @staticmethod
    def fm64(globalState):
        source3_ = D9_DC26(_dafny.CodePoint('v'))
        if source3_.is_DC27:
            d_82___mcc_h0_ = source3_.cf49
            d_83___mcc_h1_ = source3_.cf50
            d_84___mcc_h2_ = source3_.cf51
            d_85___mcc_h3_ = source3_.cf52
            d_86_cf52_ = d_85___mcc_h3_
            d_87_cf51_ = d_84___mcc_h2_
            d_88_cf50_ = d_83___mcc_h1_
            d_89_cf49_ = d_82___mcc_h0_
            return D18_DC48()
        elif source3_.is_DC26:
            d_90___mcc_h4_ = source3_.cf48
            d_91_cf48_ = d_90___mcc_h4_
            return D18_DC48()
        elif True:
            d_92___mcc_h5_ = source3_.cf53
            d_93_cf53_ = d_92___mcc_h5_
            return D18_DC48()

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        def iife31_():
            coll31_ = _dafny.Set()
            compr_31_: int
            for compr_31_ in (_dafny.MultiSet([451])).Elements:
                d_94_v0_: int = compr_31_
                if (d_94_v0_) in (_dafny.MultiSet([451])):
                    coll31_ = coll31_.union(_dafny.Set([(d_94_v0_) + (408)]))
            return _dafny.Set(coll31_)
        return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})) | (iife31_()
        )

    @staticmethod
    def fm66(p0, globalState):
        return D20_DC53()

    @staticmethod
    def fm67(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.CodePoint('n')})

    @staticmethod
    def fm68(p0, p1, globalState):
        return _dafny.Set({D20_DC53()})

    @staticmethod
    def fm69(p0, p1, p2, p3, globalState):
        def iife32_():
            coll32_ = _dafny.Set()
            compr_32_: int
            for compr_32_ in _dafny.IntegerRange(395, -382):
                d_95_v0_: int = compr_32_
                if ((395) <= (d_95_v0_)) and ((d_95_v0_) < (-382)):
                    coll32_ = coll32_.union(_dafny.Set([(d_95_v0_) - (909)]))
            return _dafny.Set(coll32_)
        return (_dafny.Map({not(not(True)): len((D3_DC8(_dafny.Map({True: len(iife32_()
)}))).cf14)})) | (_dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdbdqhl")))}))

    @staticmethod
    def fm70(globalState):
        def iife33_():
            coll33_ = _dafny.Map()
            compr_33_: str
            for compr_33_ in (_dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('c'), _dafny.CodePoint('b')])).Elements:
                d_96_v0_: str = compr_33_
                if (d_96_v0_) in (_dafny.MultiSet([_dafny.CodePoint('p'), _dafny.CodePoint('c'), _dafny.CodePoint('b')])):
                    coll33_[d_96_v0_] = len((_dafny.Set({False})).intersection(_dafny.Set({False})))
            return _dafny.Map(coll33_)
        return iife33_()
        

    @staticmethod
    def fm71(p0, p1, p2, p3, globalState):
        def iife34_():
            coll34_ = _dafny.Set()
            compr_34_: _dafny.Seq
            for compr_34_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])])).Elements:
                d_98_v0_: _dafny.Seq = compr_34_
                if (d_98_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])])):
                    coll34_ = coll34_.union(_dafny.Set([d_98_v0_]))
            return _dafny.Set(coll34_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([714]), _dafny.SeqWithoutIsStrInference([754 for d_97_i0_ in range(default__.abs(710))])})) - (iife34_()
        )

    @staticmethod
    def m0(globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_99_v0_: int
        d_99_v0_ = 826
        d_100_v1_: _dafny.Seq
        d_100_v1_ = _dafny.SeqWithoutIsStrInference([d_99_v0_])
        d_101_v2_: _dafny.Set
        d_101_v2_ = _dafny.Set({(d_100_v1_)[default__.safeIndex(d_99_v0_, len(d_100_v1_))], d_99_v0_})
        d_102_v3_: _dafny.Array
        def lambda0_(d_103_v0_):
            def lambda1_(d_104_i0_):
                return (d_104_i0_) * (d_103_v0_)

            return lambda1_

        init0_ = lambda0_(d_99_v0_)
        nw0_ = _dafny.Array(None, 15)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_102_v3_ = nw0_
        d_105_v4_: D0
        d_105_v4_ = D0_DC2(len(d_101_v2_), d_102_v3_)
        d_106_v5_: _dafny.Set
        d_106_v5_ = _dafny.Set({d_105_v4_, d_105_v4_})
        source4_ = (d_105_v4_ if (d_106_v5_).issubset(d_106_v5_) else D0_DC2(d_99_v0_, d_102_v3_))
        if source4_.is_DC1:
            d_107___mcc_h0_ = source4_.cf1
            d_108___mcc_h1_ = source4_.cf2
            d_109___mcc_h2_ = source4_.cf3
            d_110___mcc_h3_ = source4_.cf4
            d_111_cf4_ = d_110___mcc_h3_
            d_112_cf3_ = d_109___mcc_h2_
            d_113_cf2_ = d_108___mcc_h1_
            d_114_cf1_ = d_107___mcc_h0_
            if not(not((d_111_cf4_) and (d_111_cf4_))):
                d_115_v6_: str
                d_115_v6_ = _dafny.CodePoint('l')
                d_116_v7_: _dafny.MultiSet
                d_116_v7_ = _dafny.MultiSet([d_115_v6_, d_115_v6_])
                d_117_v8_: _dafny.Set
                d_117_v8_ = _dafny.Set({d_116_v7_, (d_116_v7_) | (d_116_v7_), _dafny.MultiSet([d_115_v6_])})
                d_117_v8_ = d_117_v8_
                (globalState).f8 = d_111_cf4_
                d_118_v9_: _dafny.Array
                def lambda2_(d_119_cf4_):
                    def lambda3_(d_120_i1_):
                        return (_dafny.SeqWithoutIsStrInference([d_119_cf4_, d_119_cf4_, d_119_cf4_, d_119_cf4_])) + (_dafny.SeqWithoutIsStrInference([d_119_cf4_]))

                    return lambda3_

                init1_ = lambda2_(d_111_cf4_)
                nw1_ = _dafny.Array(None, 11)
                for i0_1_ in range(nw1_.length(0)):
                    nw1_[i0_1_] = init1_(i0_1_)
                d_118_v9_ = nw1_
                d_121_v10_: _dafny.Seq
                d_121_v10_ = _dafny.SeqWithoutIsStrInference([d_111_cf4_])
                rhs0_ = (d_121_v10_)[default__.safeIndex(d_99_v0_, len(d_121_v10_))]
                rhs1_ = d_111_cf4_
                rhs2_ = ((0) - (default__.safeModuloInt(829, d_99_v0_))) != (d_99_v0_)
                rhs3_ = (d_99_v0_) * (default__.safeModuloInt(d_112_cf3_, (d_100_v1_)[default__.safeIndex(d_114_cf1_, len(d_100_v1_))]))
                rhs4_ = (d_118_v9_ if d_111_cf4_ else d_118_v9_)
                lhs0_ = globalState
                lhs0_.f8 = rhs0_
                r1 = rhs1_
                r1 = rhs2_
                d_112_cf3_ = rhs3_
                d_118_v9_ = rhs4_
                d_122_v11_: C16
                nw2_ = C16()
                nw2_.ctor__(d_111_cf4_, d_112_cf3_, d_111_cf4_, default__.fm2(globalState), d_114_cf1_)
                d_122_v11_ = nw2_
                (globalState).f8 = d_111_cf4_
            elif True:
                (globalState).f8 = d_111_cf4_
                d_123_v12_: _dafny.Array
                nw3_ = _dafny.Array(False, 7)
                d_123_v12_ = nw3_
                index0_ = default__.safeIndex(772, (d_123_v12_).length(0))
                (d_123_v12_)[index0_] = d_111_cf4_
                d_124_v13_: _dafny.Map
                d_124_v13_ = _dafny.Map({d_99_v0_: d_112_cf3_})
                d_125_v14_: D4
                d_125_v14_ = D4_DC12(d_111_cf4_, len(d_124_v13_), d_99_v0_)
                d_126_v15_: _dafny.Seq
                d_126_v15_ = _dafny.SeqWithoutIsStrInference([(d_125_v14_).cf21, d_111_cf4_])
                d_127_v16_: _dafny.MultiSet
                d_127_v16_ = _dafny.MultiSet([d_111_cf4_, default__.fm1(838, d_111_cf4_, (0) - (d_114_cf1_), globalState)])
                index1_ = default__.safeIndex(772, (d_123_v12_).length(0))
                (d_123_v12_)[index1_] = (d_126_v15_)[default__.safeIndex(((d_127_v16_)[d_111_cf4_] if (d_111_cf4_) in (d_127_v16_) else d_114_cf1_), len(d_126_v15_))]
                d_112_cf3_ = -68
                (globalState).f8 = default__.fm1((d_112_cf3_) - (d_114_cf1_), d_111_cf4_, (d_99_v0_) * (len(d_101_v2_)), globalState)
                index2_ = default__.safeIndex(772, (d_123_v12_).length(0))
                index3_ = default__.safeIndex(772, (d_123_v12_).length(0))
                rhs5_ = not(d_111_cf4_)
                rhs6_ = d_99_v0_
                rhs7_ = (d_123_v12_)[default__.safeIndex(772, (d_123_v12_).length(0))]
                lhs1_ = d_123_v12_
                lhs2_ = default__.safeIndex(772, (d_123_v12_).length(0))
                lhs3_ = d_123_v12_
                lhs4_ = default__.safeIndex(772, (d_123_v12_).length(0))
                lhs1_[lhs2_] = rhs5_
                d_112_cf3_ = rhs6_
                lhs3_[lhs4_] = rhs7_
            d_99_v0_ = d_112_cf3_
            d_128_v17_: _dafny.Array
            def lambda4_(d_129_cf4_):
                def lambda5_(d_130_i2_):
                    return (d_129_cf4_ if not(True) else d_129_cf4_)

                return lambda5_

            init2_ = lambda4_(d_111_cf4_)
            nw4_ = _dafny.Array(None, 10)
            for i0_2_ in range(nw4_.length(0)):
                nw4_[i0_2_] = init2_(i0_2_)
            d_128_v17_ = nw4_
            d_131_v18_: _dafny.MultiSet
            d_131_v18_ = _dafny.MultiSet([False])
            index4_ = default__.safeIndex(662, (d_128_v17_).length(0))
            (d_128_v17_)[index4_] = ((d_131_v18_).cardinality) <= (d_99_v0_)
            d_132_v19_: _dafny.Seq
            d_132_v19_ = _dafny.SeqWithoutIsStrInference([d_111_cf4_])
            index5_ = default__.safeIndex(662, (d_128_v17_).length(0))
            (d_128_v17_)[index5_] = ((d_132_v19_) + ((d_132_v19_).set(default__.safeIndex(len(default__.fm50(globalState)), len(d_132_v19_)), d_111_cf4_))) == (_dafny.SeqWithoutIsStrInference([default__.fm1(d_112_cf3_, d_111_cf4_, d_99_v0_, globalState)]))
            d_133_v20_: _dafny.Seq
            d_133_v20_ = _dafny.SeqWithoutIsStrInference([d_132_v19_])
            d_134_v21_: C7
            nw5_ = C7()
            nw5_.ctor__(len(default__.fm0(globalState)), d_114_cf1_)
            d_134_v21_ = nw5_
            d_135_v22_: _dafny.Map
            d_135_v22_ = _dafny.Map({d_134_v21_: d_132_v19_})
            r1 = (((d_133_v20_)[default__.safeIndex(d_112_cf3_, len(d_133_v20_))]).set(default__.safeIndex(d_99_v0_, len((d_133_v20_)[default__.safeIndex(d_112_cf3_, len(d_133_v20_))])), d_111_cf4_)) not in (_dafny.SeqWithoutIsStrInference([((d_135_v22_)[d_134_v21_] if (d_134_v21_) in (d_135_v22_) else _dafny.SeqWithoutIsStrInference([(d_128_v17_)[default__.safeIndex(662, (d_128_v17_).length(0))]]))]))
        elif source4_.is_DC2:
            d_136___mcc_h4_ = source4_.cf5
            d_137___mcc_h5_ = source4_.cf6
            d_138_cf6_ = d_137___mcc_h5_
            d_139_cf5_ = d_136___mcc_h4_
            d_140_v23_: bool
            d_140_v23_ = True
            d_141_v24_: str
            d_141_v24_ = _dafny.CodePoint('r')
            d_142_v25_: _dafny.MultiSet
            d_142_v25_ = _dafny.MultiSet([d_141_v24_])
            d_143_v26_: _dafny.Map
            d_143_v26_ = _dafny.Map({d_138_cf6_: (d_142_v25_).cardinality})
            d_144_v27_: T0
            nw6_ = C16()
            nw6_.ctor__((False) and (False), d_139_cf5_, d_140_v23_, default__.fm2(globalState), len((d_143_v26_) | (d_143_v26_)))
            d_144_v27_ = nw6_
            d_144_v27_ = d_144_v27_
            d_145_v28_: D20
            d_145_v28_ = D20_DC54(default__.fm66(d_141_v24_, globalState))
            d_145_v28_ = d_145_v28_
            rhs8_ = d_99_v0_
            rhs9_ = d_139_cf5_
            rhs10_ = ((default__.fm0(globalState)) + (d_100_v1_)) == (d_100_v1_)
            d_99_v0_ = rhs8_
            d_99_v0_ = rhs9_
            r1 = rhs10_
            d_146_v29_: _dafny.Map
            d_146_v29_ = _dafny.Map({d_144_v27_.f11: ((0) - (d_139_cf5_)) + (d_99_v0_)})
            d_99_v0_ = len(d_146_v29_)
        elif source4_.is_DC3:
            d_147___mcc_h6_ = source4_.cf7
            d_148___mcc_h7_ = source4_.cf8
            d_149_cf8_ = d_148___mcc_h7_
            d_150_cf7_ = d_147___mcc_h6_
            d_149_cf8_ = d_149_cf8_
            d_151_v30_: _dafny.Array
            nw7_ = _dafny.Array(None, 2)
            nw7_[int(0)] = True
            nw7_[int(1)] = d_149_cf8_
            d_151_v30_ = nw7_
            d_152_v31_: _dafny.MultiSet
            d_152_v31_ = _dafny.MultiSet([d_151_v30_])
            d_153_v32_: _dafny.MultiSet
            d_153_v32_ = _dafny.MultiSet([d_99_v0_, d_99_v0_])
            d_154_v33_: C12
            nw8_ = C12()
            nw8_.ctor__((d_152_v31_).intersection(d_152_v31_), d_99_v0_, (d_153_v32_).issubset(d_153_v32_), d_150_cf7_, d_99_v0_)
            d_154_v33_ = nw8_
            (globalState).f8 = d_149_cf8_
            d_155_v34_: _dafny.Map
            d_155_v34_ = _dafny.Map({(d_154_v33_).f19: D5_DC15((d_154_v33_).f19, (d_154_v33_).f19, 703)})
            d_156_v35_: D6
            d_156_v35_ = D6_DC16(d_155_v34_)
            d_156_v35_ = d_156_v35_
        elif True:
            d_157___mcc_h8_ = source4_.cf0
            d_158_cf0_ = d_157___mcc_h8_
            d_159_v36_: bool
            d_159_v36_ = True
            r1 = d_159_v36_
            d_160_v37_: str
            d_160_v37_ = _dafny.CodePoint('g')
            d_160_v37_ = d_160_v37_
            d_161_v38_: _dafny.MultiSet
            d_161_v38_ = _dafny.MultiSet([(d_105_v4_).cf6, d_102_v3_, d_102_v3_, d_102_v3_])
            d_162_v39_: _dafny.Seq
            d_162_v39_ = _dafny.SeqWithoutIsStrInference([d_161_v38_, (_dafny.MultiSet([d_102_v3_, d_102_v3_])).set(d_102_v3_, default__.abs(-378)), d_161_v38_])
            d_163_v40_: _dafny.Map
            d_163_v40_ = _dafny.Map({(d_162_v39_)[default__.safeIndex(d_99_v0_, len(d_162_v39_))]: d_99_v0_})
            d_164_v41_: _dafny.Map
            d_164_v41_ = _dafny.Map({d_160_v37_: 911})
            d_165_v42_: _dafny.Seq
            d_165_v42_ = _dafny.SeqWithoutIsStrInference([d_159_v36_])
            d_163_v40_ = (d_163_v40_).set((d_161_v38_ if default__.fm1(172, d_159_v36_, len(d_164_v41_), globalState) else d_161_v38_), len((d_165_v42_ if d_159_v36_ else (d_165_v42_).set(default__.safeIndex(d_99_v0_, len(d_165_v42_)), d_159_v36_))))
            d_166_v43_: _dafny.Array
            nw9_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
            d_166_v43_ = nw9_
            index6_ = default__.safeIndex(488, (d_166_v43_).length(0))
            (d_166_v43_)[index6_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhw"))
            d_167_v44_: _dafny.Seq
            d_167_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcb"))
            index7_ = default__.safeIndex(488, (d_166_v43_).length(0))
            (d_166_v43_)[index7_] = d_167_v44_
        d_168_v45_: bool
        d_168_v45_ = True
        d_169_v46_: str
        d_169_v46_ = _dafny.CodePoint('i')
        d_170_v47_: _dafny.Set
        d_170_v47_ = _dafny.Set({d_169_v46_})
        (globalState).f8 = (default__.fm67(d_168_v45_, d_168_v45_, d_99_v0_, globalState)) == (d_170_v47_)
        if default__.fm1(len(d_170_v47_), False, d_99_v0_, globalState):
            if True:
                d_171_v48_: C15
                nw10_ = C15()
                nw10_.ctor__(d_168_v45_, d_99_v0_, d_168_v45_)
                d_171_v48_ = nw10_
                d_172_v49_: _dafny.Seq
                d_172_v49_ = _dafny.SeqWithoutIsStrInference([d_171_v48_])
                d_173_v50_: _dafny.Map
                d_173_v50_ = _dafny.Map({d_172_v49_: d_99_v0_})
                d_174_v51_: _dafny.Map
                d_174_v51_ = _dafny.Map({(d_173_v50_) | (_dafny.Map({d_172_v49_: d_99_v0_})): d_171_v48_.f16})
                d_174_v51_ = (d_174_v51_).set(d_173_v50_, d_171_v48_.f16)
                index8_ = default__.safeIndex(294, (d_102_v3_).length(0))
                (d_102_v3_)[index8_] = d_171_v48_.f17
                index9_ = default__.safeIndex(294, (d_102_v3_).length(0))
                (d_102_v3_)[index9_] = d_99_v0_
                d_175_v52_: _dafny.Array
                nw11_ = _dafny.Array(False, 11)
                d_175_v52_ = nw11_
                d_176_v53_: _dafny.MultiSet
                d_176_v53_ = _dafny.MultiSet([d_175_v52_])
                d_177_v54_: _dafny.Seq
                d_177_v54_ = _dafny.SeqWithoutIsStrInference([d_171_v48_.f16, d_171_v48_.f16])
                d_178_v55_: _dafny.MultiSet
                d_178_v55_ = _dafny.MultiSet([d_168_v45_, (d_177_v54_)[default__.safeIndex(d_171_v48_.f17, len(d_177_v54_))], d_171_v48_.f16, True, d_168_v45_])
                d_179_v56_: T0
                nw12_ = C12()
                nw12_.ctor__(d_176_v53_, ((d_178_v55_).cardinality) * (d_171_v48_.f17), d_168_v45_, d_171_v48_.f17, (d_102_v3_)[default__.safeIndex(294, (d_102_v3_).length(0))])
                d_179_v56_ = nw12_
                d_180_v57_: _dafny.Map
                d_180_v57_ = _dafny.Map({(d_102_v3_)[default__.safeIndex(294, (d_102_v3_).length(0))]: d_175_v52_})
                d_181_v58_: _dafny.Array
                nw13_ = _dafny.Array(None, 17)
                nw13_[int(0)] = d_175_v52_
                nw13_[int(1)] = d_175_v52_
                nw13_[int(2)] = d_175_v52_
                nw13_[int(3)] = d_175_v52_
                nw13_[int(4)] = d_175_v52_
                nw13_[int(5)] = (d_175_v52_ if d_179_v56_.f11 else d_175_v52_)
                nw13_[int(6)] = d_175_v52_
                nw13_[int(7)] = d_175_v52_
                nw13_[int(8)] = d_175_v52_
                nw13_[int(9)] = d_175_v52_
                nw13_[int(10)] = ((d_180_v57_)[(0) - (d_99_v0_)] if ((0) - (d_99_v0_)) in (d_180_v57_) else d_175_v52_)
                nw13_[int(11)] = d_175_v52_
                nw13_[int(12)] = d_175_v52_
                nw13_[int(13)] = d_175_v52_
                nw13_[int(14)] = d_175_v52_
                nw13_[int(15)] = d_175_v52_
                nw13_[int(16)] = d_175_v52_
                d_181_v58_ = nw13_
                rhs11_ = d_179_v56_.f11
                rhs12_ = d_179_v56_
                rhs13_ = d_171_v48_.f16
                rhs14_ = (d_102_v3_)[default__.safeIndex(294, (d_102_v3_).length(0))]
                rhs15_ = d_181_v58_
                lhs5_ = globalState
                lhs6_ = d_171_v48_
                lhs5_.f8 = rhs11_
                d_179_v56_ = rhs12_
                r1 = rhs13_
                lhs6_.f17 = rhs14_
                d_181_v58_ = rhs15_
                index10_ = default__.safeIndex(294, (d_102_v3_).length(0))
                (d_102_v3_)[index10_] = len(d_100_v1_)
                d_182_v59_: _dafny.Map
                d_182_v59_ = _dafny.Map({d_171_v48_.f16: d_179_v56_.f11})
                d_183_v60_: C14
                nw14_ = C14()
                nw14_.ctor__((d_171_v48_.f16) == (((d_182_v59_)[d_179_v56_.f11] if (d_179_v56_.f11) in (d_182_v59_) else not(d_179_v56_.f11))))
                d_183_v60_ = nw14_
            elif True:
                d_184_v61_: C4
                nw15_ = C4()
                nw15_.ctor__((d_99_v0_) * (d_99_v0_), d_168_v45_)
                d_184_v61_ = nw15_
                d_185_v62_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.MultiSet({}), 8)
                d_185_v62_ = nw16_
                d_186_v63_: _dafny.MultiSet
                d_186_v63_ = _dafny.MultiSet([d_168_v45_, True, not(not(d_168_v45_))])
                index11_ = default__.safeIndex(654, (d_185_v62_).length(0))
                (d_185_v62_)[index11_] = d_186_v63_
                d_187_v64_: _dafny.Seq
                d_187_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncihubldw"))
                d_188_v65_: C8
                nw17_ = C8()
                nw17_.ctor__(_dafny.SeqWithoutIsStrInference([d_99_v0_ for d_189_i3_ in range(default__.abs(957))]))
                d_188_v65_ = nw17_
                d_190_v66_: _dafny.Map
                d_190_v66_ = _dafny.Map({d_188_v65_: default__.fm36(d_99_v0_, d_99_v0_, d_169_v46_, globalState)})
                d_191_v67_: _dafny.Map
                d_191_v67_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgjuj")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))})
                d_192_v68_: _dafny.Map
                d_192_v68_ = _dafny.Map({d_99_v0_: d_169_v46_})
                d_193_v69_: _dafny.Array
                nw18_ = _dafny.Array(None, 26)
                nw18_[int(0)] = d_187_v64_
                nw18_[int(1)] = d_187_v64_
                nw18_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkd"))
                nw18_[int(3)] = d_187_v64_
                nw18_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrjw"))
                nw18_[int(5)] = d_187_v64_
                nw18_[int(6)] = d_187_v64_
                nw18_[int(7)] = default__.fm36((d_184_v61_).f25, (d_184_v61_).f25, d_169_v46_, globalState)
                nw18_[int(8)] = d_187_v64_
                nw18_[int(9)] = ((d_190_v66_)[d_188_v65_] if (d_188_v65_) in (d_190_v66_) else ((d_191_v67_)[d_187_v64_] if (d_187_v64_) in (d_191_v67_) else d_187_v64_))
                nw18_[int(10)] = (d_187_v64_).set(default__.safeIndex((d_184_v61_).f25, len(d_187_v64_)), ((d_192_v68_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_192_v68_) else d_169_v46_))
                nw18_[int(11)] = d_187_v64_
                nw18_[int(12)] = d_187_v64_
                nw18_[int(13)] = d_187_v64_
                nw18_[int(14)] = d_187_v64_
                nw18_[int(15)] = d_187_v64_
                nw18_[int(16)] = d_187_v64_
                nw18_[int(17)] = d_187_v64_
                nw18_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                nw18_[int(19)] = (_dafny.SeqWithoutIsStrInference([d_169_v46_ for d_194_i4_ in range(default__.abs(537))])).set(default__.safeIndex((d_184_v61_).f25, len(_dafny.SeqWithoutIsStrInference([d_169_v46_ for d_195_i4_ in range(default__.abs(537))]))), d_169_v46_)
                nw18_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bljdb"))
                nw18_[int(21)] = d_187_v64_
                nw18_[int(22)] = d_187_v64_
                nw18_[int(23)] = d_187_v64_
                nw18_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdgj"))
                nw18_[int(25)] = d_187_v64_
                d_193_v69_ = nw18_
                d_196_v70_: _dafny.Map
                d_196_v70_ = _dafny.Map({d_193_v69_: d_184_v61_})
                d_197_v71_: _dafny.Seq
                d_197_v71_ = _dafny.SeqWithoutIsStrInference([d_186_v63_, d_186_v63_])
                d_198_v72_: _dafny.Array
                nw19_ = _dafny.Array(False, 12)
                d_198_v72_ = nw19_
                d_199_v73_: _dafny.Seq
                d_199_v73_ = _dafny.SeqWithoutIsStrInference([d_198_v72_])
                index12_ = default__.safeIndex(654, (d_185_v62_).length(0))
                rhs16_ = ((d_196_v70_)[d_193_v69_] if (d_193_v69_) in (d_196_v70_) else d_184_v61_)
                rhs17_ = (d_197_v71_)[default__.safeIndex(len(d_199_v73_), len(d_197_v71_))]
                lhs7_ = d_185_v62_
                lhs8_ = default__.safeIndex(654, (d_185_v62_).length(0))
                d_184_v61_ = rhs16_
                lhs7_[lhs8_] = rhs17_
                d_200_v74_: _dafny.MultiSet
                d_200_v74_ = _dafny.MultiSet([d_198_v72_, d_198_v72_])
                d_201_v75_: C12
                nw20_ = C12()
                nw20_.ctor__(d_200_v74_, 211, True, 121, d_99_v0_)
                d_201_v75_ = nw20_
                d_202_v76_: _dafny.Seq
                d_202_v76_ = _dafny.SeqWithoutIsStrInference([d_201_v75_])
                d_99_v0_ = len(d_202_v76_)
                d_99_v0_ = len(d_170_v47_)
                d_99_v0_ = (d_184_v61_).f25
                d_203_v77_: D21
                d_203_v77_ = D21_DC56(True, not(d_168_v45_), d_99_v0_)
                d_99_v0_ = (d_203_v77_).cf84
            if d_168_v45_:
                d_204_v78_: _dafny.Array
                def lambda6_(d_205_v45_):
                    def lambda7_(d_206_i5_):
                        return not((d_205_v45_) or (d_205_v45_))

                    return lambda7_

                init3_ = lambda6_(d_168_v45_)
                nw21_ = _dafny.Array(None, 16)
                for i0_3_ in range(nw21_.length(0)):
                    nw21_[i0_3_] = init3_(i0_3_)
                d_204_v78_ = nw21_
                index13_ = default__.safeIndex(309, (d_204_v78_).length(0))
                (d_204_v78_)[index13_] = d_168_v45_
                d_207_v79_: _dafny.Map
                d_207_v79_ = _dafny.Map({925: d_168_v45_})
                d_208_v80_: _dafny.Map
                d_208_v80_ = _dafny.Map({d_99_v0_: (d_99_v0_) not in (d_207_v79_)})
                index14_ = default__.safeIndex(309, (d_204_v78_).length(0))
                (d_204_v78_)[index14_] = ((d_208_v80_)[default__.safeModuloInt(102, d_99_v0_)] if (default__.safeModuloInt(102, d_99_v0_)) in (d_208_v80_) else d_168_v45_)
                (globalState).f8 = d_168_v45_
                d_99_v0_ = d_99_v0_
                (globalState).f8 = (d_204_v78_)[default__.safeIndex(309, (d_204_v78_).length(0))]
                d_209_v81_: C10
                nw22_ = C10()
                nw22_.ctor__((d_204_v78_)[default__.safeIndex(309, (d_204_v78_).length(0))])
                d_209_v81_ = nw22_
            elif True:
                index15_ = default__.safeIndex(676, (d_102_v3_).length(0))
                (d_102_v3_)[index15_] = d_99_v0_
                index16_ = default__.safeIndex(676, (d_102_v3_).length(0))
                (d_102_v3_)[index16_] = d_99_v0_
                d_210_v82_: _dafny.Map
                d_210_v82_ = _dafny.Map({d_100_v1_: d_168_v45_})
                d_211_v83_: _dafny.Map
                d_211_v83_ = _dafny.Map({(d_102_v3_)[default__.safeIndex(676, (d_102_v3_).length(0))]: 196})
                d_212_v85_: _dafny.Seq
                def iife35_():
                    coll35_ = _dafny.Map()
                    compr_35_: D20
                    for compr_35_ in (default__.fm68(d_169_v46_, d_168_v45_, globalState)).Elements:
                        d_213_v84_: D20 = compr_35_
                        if (d_213_v84_) in (default__.fm68(d_169_v46_, d_168_v45_, globalState)):
                            coll35_[d_213_v84_] = _dafny.MultiSet([_dafny.Map({d_168_v45_: d_168_v45_}), _dafny.Map({d_168_v45_: d_168_v45_})])
                    return _dafny.Map(coll35_)
                d_212_v85_ = _dafny.SeqWithoutIsStrInference([iife35_()
                ])
                d_214_v87_: D20
                d_214_v87_ = D20_DC53()
                d_215_v88_: _dafny.Map
                d_215_v88_ = _dafny.Map({d_168_v45_: d_168_v45_})
                d_216_v89_: _dafny.MultiSet
                d_216_v89_ = _dafny.MultiSet([d_215_v88_])
                d_217_v90_: _dafny.Seq
                d_217_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tf"))
                d_218_v91_: _dafny.MultiSet
                d_218_v91_ = ((d_216_v89_).set(d_215_v88_, default__.abs(len(d_217_v90_)))).set(d_215_v88_, default__.abs(d_99_v0_))
                d_219_v92_: _dafny.Map
                d_219_v92_ = _dafny.Map({d_214_v87_: d_218_v91_})
                d_220_v93_: _dafny.Set
                d_220_v93_ = _dafny.Set({_dafny.Map({d_214_v87_: d_216_v89_}), d_219_v92_})
                d_221_v94_: _dafny.Array
                nw23_ = _dafny.Array(None, 16)
                nw23_[int(0)] = (d_99_v0_) > (len(d_210_v82_))
                nw23_[int(1)] = not(not(not((default__.fm2(globalState)) in (d_211_v83_))))
                nw23_[int(2)] = False
                nw23_[int(3)] = d_168_v45_
                nw23_[int(4)] = not(d_168_v45_)
                nw23_[int(5)] = d_168_v45_
                nw23_[int(6)] = d_168_v45_
                nw23_[int(7)] = (d_99_v0_) >= (980)
                nw23_[int(8)] = False
                nw23_[int(9)] = (d_99_v0_) < (len(d_100_v1_))
                nw23_[int(10)] = (default__.fm2(globalState)) != ((d_102_v3_)[default__.safeIndex(676, (d_102_v3_).length(0))])
                def iife36_():
                    coll36_ = _dafny.Set()
                    compr_36_: _dafny.Map
                    for compr_36_ in (d_212_v85_).Elements:
                        d_222_v86_: _dafny.Map = compr_36_
                        if (d_222_v86_) in (d_212_v85_):
                            coll36_ = coll36_.union(_dafny.Set([d_222_v86_]))
                    return _dafny.Set(coll36_)
                nw23_[int(11)] = (iife36_()
                ).isdisjoint(d_220_v93_)
                nw23_[int(12)] = d_168_v45_
                nw23_[int(13)] = d_168_v45_
                nw23_[int(14)] = (d_99_v0_) not in (d_100_v1_)
                nw23_[int(15)] = d_168_v45_
                d_221_v94_ = nw23_
                index17_ = default__.safeIndex(730, (d_221_v94_).length(0))
                (d_221_v94_)[index17_] = not(d_168_v45_)
                d_223_v95_: _dafny.MultiSet
                d_223_v95_ = _dafny.MultiSet([d_168_v45_])
                index18_ = default__.safeIndex(730, (d_221_v94_).length(0))
                (d_221_v94_)[index18_] = ((d_99_v0_) + ((d_102_v3_)[default__.safeIndex(676, (d_102_v3_).length(0))])) > (((d_223_v95_)[d_168_v45_] if (d_168_v45_) in (d_223_v95_) else d_99_v0_))
                d_224_v96_: D15
                d_224_v96_ = D15_DC41(d_168_v45_)
                d_225_v97_: _dafny.Map
                d_225_v97_ = _dafny.Map({d_169_v46_: d_224_v96_})
                d_225_v97_ = (d_225_v97_).set(d_169_v46_, d_224_v96_)
                d_226_v98_: D19
                d_226_v98_ = D19_DC51(((d_102_v3_)[default__.safeIndex(676, (d_102_v3_).length(0))]) - (len(d_217_v90_)), ((d_223_v95_)[d_168_v45_] if (d_168_v45_) in (d_223_v95_) else (d_102_v3_)[default__.safeIndex(676, (d_102_v3_).length(0))]))
                d_226_v98_ = (d_226_v98_ if (d_223_v95_) != (d_223_v95_) else (d_226_v98_ if (d_221_v94_)[default__.safeIndex(730, (d_221_v94_).length(0))] else d_226_v98_))
                d_227_v99_: _dafny.Array
                nw24_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                d_227_v99_ = nw24_
                index19_ = default__.safeIndex(648, (d_227_v99_).length(0))
                (d_227_v99_)[index19_] = d_169_v46_
                index20_ = default__.safeIndex(676, (d_102_v3_).length(0))
                index21_ = default__.safeIndex(648, (d_227_v99_).length(0))
                index22_ = default__.safeIndex(730, (d_221_v94_).length(0))
                rhs18_ = (d_102_v3_)[default__.safeIndex(676, (d_102_v3_).length(0))]
                rhs19_ = _dafny.CodePoint('w')
                rhs20_ = (False) and ((d_221_v94_)[default__.safeIndex(730, (d_221_v94_).length(0))])
                lhs9_ = d_102_v3_
                lhs10_ = default__.safeIndex(676, (d_102_v3_).length(0))
                lhs11_ = d_227_v99_
                lhs12_ = default__.safeIndex(648, (d_227_v99_).length(0))
                lhs13_ = d_221_v94_
                lhs14_ = default__.safeIndex(730, (d_221_v94_).length(0))
                lhs9_[lhs10_] = rhs18_
                lhs11_[lhs12_] = rhs19_
                lhs13_[lhs14_] = rhs20_
            d_228_v100_: C2
            nw25_ = C2()
            nw25_.ctor__((d_99_v0_) * (d_99_v0_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvjqaslaa"))).set(default__.safeIndex((0) - (d_99_v0_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvjqaslaa")))), d_169_v46_), d_99_v0_, d_99_v0_, default__.fm1(d_99_v0_, False, -216, globalState))
            d_228_v100_ = nw25_
            d_228_v100_ = d_228_v100_
            index23_ = default__.safeIndex(981, (d_102_v3_).length(0))
            (d_102_v3_)[index23_] = -865
            index24_ = default__.safeIndex(981, (d_102_v3_).length(0))
            (d_102_v3_)[index24_] = d_99_v0_
            d_229_v101_: _dafny.Map
            d_229_v101_ = _dafny.Map({(d_228_v100_).f26: d_99_v0_})
            (globalState).f8 = (d_229_v101_) == ((d_229_v101_) | (d_229_v101_))
        elif True:
            d_230_v102_: D12
            d_230_v102_ = D12_DC35(d_99_v0_)
            source5_ = d_230_v102_
            if source5_.is_DC33:
                d_231___mcc_h9_ = source5_.cf58
                d_232_cf58_ = d_231___mcc_h9_
                r1 = d_168_v45_
                d_233_v103_: T0
                nw26_ = C15()
                nw26_.ctor__(d_168_v45_, 747, False)
                d_233_v103_ = nw26_
                d_234_v104_: _dafny.Map
                d_234_v104_ = _dafny.Map({d_168_v45_: d_233_v103_})
                d_235_v105_: _dafny.Set
                d_235_v105_ = _dafny.Set({(D0_DC3(len(d_234_v104_), d_233_v103_.f11)).cf8, d_168_v45_, not(d_233_v103_.f11)})
                (globalState).f8 = (not (True) or (True)) not in (d_235_v105_)
                d_99_v0_ = d_232_cf58_
                d_168_v45_ = default__.fm1(d_99_v0_, not(d_233_v103_.f11), d_232_cf58_, globalState)
            elif source5_.is_DC34:
                d_236___mcc_h10_ = source5_.cf59
                d_237_cf59_ = d_236___mcc_h10_
                d_238_v106_: _dafny.Array
                def lambda8_(d_239_v45_):
                    def lambda9_(d_240_i6_):
                        return d_239_v45_

                    return lambda9_

                init4_ = lambda8_(d_168_v45_)
                nw27_ = _dafny.Array(None, 4)
                for i0_4_ in range(nw27_.length(0)):
                    nw27_[i0_4_] = init4_(i0_4_)
                d_238_v106_ = nw27_
                d_241_v107_: _dafny.Seq
                d_241_v107_ = _dafny.SeqWithoutIsStrInference([d_238_v106_])
                d_241_v107_ = ((d_241_v107_) + (d_241_v107_)) + (_dafny.SeqWithoutIsStrInference([d_238_v106_, d_238_v106_, d_238_v106_, d_238_v106_, d_238_v106_]))
                d_242_v108_: C17
                nw28_ = C17()
                nw28_.ctor__(d_168_v45_)
                d_242_v108_ = nw28_
                d_243_v109_: _dafny.Array
                nw29_ = _dafny.Array(None, 2)
                nw29_[int(0)] = d_242_v108_
                nw29_[int(1)] = d_242_v108_
                d_243_v109_ = nw29_
                d_244_v110_: C17
                d_244_v110_ = d_242_v108_
                index25_ = default__.safeIndex(201, (d_243_v109_).length(0))
                (d_243_v109_)[index25_] = (d_244_v110_)
                index26_ = default__.safeIndex(201, (d_243_v109_).length(0))
                (d_243_v109_)[index26_] = d_242_v108_
                (globalState).f8 = d_168_v45_
                d_168_v45_ = d_168_v45_
            elif source5_.is_DC35:
                d_245___mcc_h11_ = source5_.cf60
                d_246_cf60_ = d_245___mcc_h11_
                d_247_v111_: _dafny.Seq
                d_247_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmrtyw"))
                d_246_cf60_ = len(((d_247_v111_) + (d_247_v111_)) + (_dafny.SeqWithoutIsStrInference([d_169_v46_ for d_248_i7_ in range(default__.abs(591))])))
                d_99_v0_ = d_99_v0_
                d_249_v112_: _dafny.Array
                def lambda10_(d_250_v45_):
                    def lambda11_(d_251_i8_):
                        return d_250_v45_

                    return lambda11_

                init5_ = lambda10_(d_168_v45_)
                nw30_ = _dafny.Array(None, 2)
                for i0_5_ in range(nw30_.length(0)):
                    nw30_[i0_5_] = init5_(i0_5_)
                d_249_v112_ = nw30_
                d_252_v113_: _dafny.MultiSet
                d_252_v113_ = _dafny.MultiSet([d_249_v112_, d_249_v112_])
                d_253_v114_: _dafny.Map
                d_253_v114_ = _dafny.Map({(d_252_v113_).issubset(d_252_v113_): d_246_cf60_})
                d_254_v115_: _dafny.Map
                d_254_v115_ = _dafny.Map({True: _dafny.Map({d_246_cf60_: d_168_v45_})})
                d_255_v116_: C10
                nw31_ = C10()
                nw31_.ctor__(True)
                d_255_v116_ = nw31_
                d_256_v118_: _dafny.Map
                d_256_v118_ = _dafny.Map({d_255_v116_: default__.fm70(globalState)})
                def iife37_():
                    coll37_ = _dafny.Map()
                    compr_37_: str
                    for compr_37_ in (d_170_v47_).Elements:
                        d_257_v117_: str = compr_37_
                        if (d_257_v117_) in (d_170_v47_):
                            coll37_[d_257_v117_] = 343
                    return _dafny.Map(coll37_)
                d_253_v114_ = default__.fm69((d_254_v115_) | (_dafny.Map({d_168_v45_: _dafny.Map({d_99_v0_: d_168_v45_})})), default__.safeModuloInt(d_99_v0_, d_99_v0_), d_246_cf60_, (_dafny.Map({d_255_v116_: iife37_()
                })) != (d_256_v118_), globalState)
                (globalState).f8 = d_168_v45_
            elif True:
                d_258___mcc_h12_ = source5_.cf57
                d_259_cf57_ = d_258___mcc_h12_
                d_260_v119_: _dafny.Seq
                d_260_v119_ = _dafny.SeqWithoutIsStrInference([d_168_v45_, d_168_v45_, d_168_v45_, not(d_168_v45_), d_168_v45_])
                d_260_v119_ = ((d_260_v119_) + (d_260_v119_)) + (d_260_v119_)
                d_261_v120_: _dafny.Array
                nw32_ = _dafny.Array(False, 27)
                d_261_v120_ = nw32_
                d_262_v121_: C4
                nw33_ = C4()
                nw33_.ctor__(d_99_v0_, d_168_v45_)
                d_262_v121_ = nw33_
                d_263_v122_: _dafny.Map
                d_263_v122_ = _dafny.Map({(d_261_v120_ if d_168_v45_ else d_261_v120_): d_262_v121_})
                d_263_v122_ = (d_263_v122_).set(d_261_v120_, d_262_v121_)
                (globalState).f8 = False
                index27_ = default__.safeIndex(597, (d_102_v3_).length(0))
                (d_102_v3_)[index27_] = (d_262_v121_).f25
                index28_ = default__.safeIndex(597, (d_102_v3_).length(0))
                (d_102_v3_)[index28_] = d_99_v0_
            d_264_v123_: _dafny.Seq
            d_264_v123_ = _dafny.SeqWithoutIsStrInference([d_168_v45_])
            r1 = not((d_264_v123_)[default__.safeIndex((0) - (d_99_v0_), len(d_264_v123_))])
            d_264_v123_ = d_264_v123_
            d_265_v124_: _dafny.Array
            def lambda12_(d_266_v45_):
                def lambda13_(d_267_i9_):
                    return d_266_v45_

                return lambda13_

            init6_ = lambda12_(d_168_v45_)
            nw34_ = _dafny.Array(None, 27)
            for i0_6_ in range(nw34_.length(0)):
                nw34_[i0_6_] = init6_(i0_6_)
            d_265_v124_ = nw34_
            d_268_v125_: _dafny.Map
            d_268_v125_ = _dafny.Map({d_265_v124_: d_168_v45_})
            d_269_v126_: T0
            nw35_ = C4()
            nw35_.ctor__(d_99_v0_, d_168_v45_)
            d_269_v126_ = nw35_
            d_270_v127_: _dafny.Map
            d_270_v127_ = _dafny.Map({_dafny.Map({d_269_v126_: not(d_168_v45_)}): d_268_v125_})
            d_271_v128_: _dafny.Map
            d_271_v128_ = _dafny.Map({d_269_v126_: d_269_v126_.f11})
            d_272_v129_: _dafny.Seq
            d_272_v129_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lof"))
            d_273_v130_: _dafny.Seq
            d_273_v130_ = _dafny.SeqWithoutIsStrInference([d_272_v129_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeksf"))])
            d_274_v131_: _dafny.Map
            d_274_v131_ = _dafny.Map({not((d_269_v126_).fm6(d_269_v126_.f11, len(d_272_v129_), len(d_273_v130_), globalState)): _dafny.Map({d_265_v124_: d_168_v45_})})
            d_275_v133_: _dafny.MultiSet
            d_275_v133_ = _dafny.MultiSet([719])
            d_276_v134_: _dafny.Array
            nw36_ = _dafny.Array(None, 20)
            nw36_[int(0)] = d_268_v125_
            nw36_[int(1)] = d_268_v125_
            nw36_[int(2)] = d_268_v125_
            nw36_[int(3)] = (d_268_v125_) | (d_268_v125_)
            nw36_[int(4)] = d_268_v125_
            nw36_[int(5)] = d_268_v125_
            nw36_[int(6)] = d_268_v125_
            nw36_[int(7)] = d_268_v125_
            nw36_[int(8)] = d_268_v125_
            nw36_[int(9)] = d_268_v125_
            nw36_[int(10)] = d_268_v125_
            nw36_[int(11)] = ((d_270_v127_)[(d_271_v128_).set(d_269_v126_, (d_269_v126_).fm6(d_269_v126_.f11, d_99_v0_, d_99_v0_, globalState))] if ((d_271_v128_).set(d_269_v126_, (d_269_v126_).fm6(d_269_v126_.f11, d_99_v0_, d_99_v0_, globalState))) in (d_270_v127_) else d_268_v125_)
            nw36_[int(12)] = (d_268_v125_) | (d_268_v125_)
            nw36_[int(13)] = d_268_v125_
            nw36_[int(14)] = d_268_v125_
            nw36_[int(15)] = d_268_v125_
            nw36_[int(16)] = (d_268_v125_) | (d_268_v125_)
            nw36_[int(17)] = ((d_274_v131_)[False] if (False) in (d_274_v131_) else _dafny.Map({d_265_v124_: True}))
            def iife38_():
                coll38_ = _dafny.Map()
                compr_38_: int
                for compr_38_ in (d_275_v133_).Elements:
                    d_277_v132_: int = compr_38_
                    if (d_277_v132_) in (d_275_v133_):
                        coll38_[default__.safeModuloInt(d_277_v132_, 167)] = d_269_v126_.f11
                return _dafny.Map(coll38_)
            nw36_[int(18)] = _dafny.Map({d_265_v124_: (d_269_v126_).fm6(not(d_168_v45_), len((iife38_()
            ).set(d_99_v0_, False)), d_99_v0_, globalState)})
            nw36_[int(19)] = _dafny.Map({d_265_v124_: d_269_v126_.f11})
            d_276_v134_ = nw36_
            index29_ = default__.safeIndex(139, (d_276_v134_).length(0))
            (d_276_v134_)[index29_] = d_268_v125_
            index30_ = default__.safeIndex(139, (d_276_v134_).length(0))
            (d_276_v134_)[index30_] = (d_268_v125_).set(d_265_v124_, True)
            d_278_v135_: _dafny.MultiSet
            d_278_v135_ = _dafny.MultiSet([d_265_v124_, d_265_v124_, d_265_v124_, d_265_v124_, d_265_v124_])
            d_279_v136_: C12
            nw37_ = C12()
            nw37_.ctor__(d_278_v135_, (d_99_v0_) * (d_99_v0_), d_269_v126_.f11, (d_99_v0_ if True else d_99_v0_), (d_99_v0_) - (len(d_264_v123_)))
            d_279_v136_ = nw37_
        d_280_v137_: _dafny.MultiSet
        d_280_v137_ = _dafny.MultiSet([d_168_v45_, (d_168_v45_ if d_168_v45_ else d_168_v45_), default__.fm1(d_99_v0_, d_168_v45_, 338, globalState)])
        d_281_v138_: _dafny.Seq
        d_281_v138_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyydy"))
        d_282_v139_: _dafny.Seq
        d_282_v139_ = _dafny.SeqWithoutIsStrInference([d_168_v45_, d_168_v45_])
        d_283_v140_: _dafny.MultiSet
        d_283_v140_ = _dafny.MultiSet([len(d_101_v2_), 784, d_99_v0_, len(d_281_v138_), len(d_282_v139_)])
        d_280_v137_ = default__.fm30((_dafny.MultiSet([363, d_99_v0_])).isdisjoint(d_283_v140_), d_168_v45_, d_99_v0_, globalState)
        d_284_v141_: _dafny.Map
        d_284_v141_ = _dafny.Map({537: d_168_v45_})
        d_284_v141_ = (d_284_v141_).set(d_99_v0_, d_168_v45_)
        d_285_v142_: T2
        nw38_ = C2()
        nw38_.ctor__(d_99_v0_, d_281_v138_, d_99_v0_, d_99_v0_, d_168_v45_)
        d_285_v142_ = nw38_
        d_286_v143_: _dafny.Map
        d_286_v143_ = _dafny.Map({d_285_v142_: d_168_v45_})
        if (d_168_v45_) == (((d_286_v143_)[d_285_v142_] if (d_285_v142_) in (d_286_v143_) else d_168_v45_)):
            d_287_v144_: _dafny.Array
            def lambda14_(d_288_v45_):
                def lambda15_(d_289_i10_):
                    return d_288_v45_

                return lambda15_

            init7_ = lambda14_(d_168_v45_)
            nw39_ = _dafny.Array(None, 29)
            for i0_7_ in range(nw39_.length(0)):
                nw39_[i0_7_] = init7_(i0_7_)
            d_287_v144_ = nw39_
            d_290_v145_: _dafny.MultiSet
            d_290_v145_ = _dafny.MultiSet([d_287_v144_])
            d_291_v146_: D21
            d_291_v146_ = D21_DC55(d_290_v145_)
            d_292_v147_: D15
            d_292_v147_ = D15_DC41(default__.fm1((d_285_v142_).f12, d_168_v45_, (d_285_v142_).f12, globalState))
            d_293_v148_: D15
            d_293_v148_ = D15_DC42(d_292_v147_)
            d_294_v149_: _dafny.Array
            d_294_v149_ = d_287_v144_
            d_295_v150_: _dafny.Map
            d_295_v150_ = _dafny.Map({(d_285_v142_).f12: (d_285_v142_).f13})
            rhs21_ = d_291_v146_
            rhs22_ = (d_294_v149_)
            rhs23_ = d_293_v148_
            rhs24_ = ((d_285_v142_).f13) * (((d_295_v150_)[(d_285_v142_).f13] if ((d_285_v142_).f13) in (d_295_v150_) else (d_285_v142_).f13))
            d_291_v146_ = rhs21_
            d_287_v144_ = rhs22_
            d_293_v148_ = rhs23_
            d_99_v0_ = rhs24_
            d_283_v140_ = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_99_v0_, (d_285_v142_).f13]))).set((d_285_v142_).f12, default__.abs((d_285_v142_).f12))) | (default__.fm12((0) - ((d_285_v142_).f12), d_99_v0_, len(_dafny.Set({d_168_v45_, d_168_v45_})), d_168_v45_, globalState))
            d_168_v45_ = d_168_v45_
            d_296_v151_: C3
            nw40_ = C3()
            nw40_.ctor__(((d_285_v142_).f12 if d_168_v45_ else 534), len(d_100_v1_))
            d_296_v151_ = nw40_
            d_99_v0_ = default__.safeDivisionInt(-313, default__.safeModuloInt((d_285_v142_).f13, (d_285_v142_).f13))
        elif True:
            d_297_v152_: C3
            nw41_ = C3()
            nw41_.ctor__((d_285_v142_).f13, (d_285_v142_).f12)
            d_297_v152_ = nw41_
            d_297_v152_ = d_297_v152_
            d_298_v153_: _dafny.Array
            nw42_ = _dafny.Array(None, 20)
            nw42_[int(0)] = (d_285_v142_).f20
            nw42_[int(1)] = ((d_285_v142_).f20) + (d_281_v138_)
            nw42_[int(2)] = d_281_v138_
            nw42_[int(3)] = (d_281_v138_) + (d_281_v138_)
            nw42_[int(4)] = (d_285_v142_).f20
            nw42_[int(5)] = d_281_v138_
            nw42_[int(6)] = (d_281_v138_).set(default__.safeIndex((d_285_v142_).f12, len(d_281_v138_)), d_169_v46_)
            nw42_[int(7)] = ((d_285_v142_).f20 if d_168_v45_ else _dafny.SeqWithoutIsStrInference([d_169_v46_ for d_299_i11_ in range(default__.abs(701))]))
            nw42_[int(8)] = ((d_285_v142_).f20) + ((d_285_v142_).f20)
            nw42_[int(9)] = (d_285_v142_).f20
            nw42_[int(10)] = d_281_v138_
            nw42_[int(11)] = (d_285_v142_).f20
            nw42_[int(12)] = d_281_v138_
            nw42_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjuxrcwi"))
            nw42_[int(14)] = ((d_285_v142_).f20).set(default__.safeIndex((d_285_v142_).f13, len((d_285_v142_).f20)), d_169_v46_)
            nw42_[int(15)] = (d_285_v142_).f20
            nw42_[int(16)] = (d_281_v138_) + ((d_285_v142_).f20)
            nw42_[int(17)] = d_281_v138_
            nw42_[int(18)] = (d_285_v142_).f20
            nw42_[int(19)] = (d_285_v142_).f20
            d_298_v153_ = nw42_
            index31_ = default__.safeIndex(903, (d_298_v153_).length(0))
            (d_298_v153_)[index31_] = ((d_285_v142_).f20) + (_dafny.SeqWithoutIsStrInference([(D9_DC26(d_169_v46_)).cf48 for d_300_i12_ in range(default__.abs(141))]))
            index32_ = default__.safeIndex(903, (d_298_v153_).length(0))
            (d_298_v153_)[index32_] = (d_285_v142_).f20
            if ((D6_DC17(d_168_v45_, d_99_v0_)).cf30) and ((d_168_v45_) == (d_168_v45_)):
                d_301_v154_: _dafny.Map
                d_301_v154_ = _dafny.Map({d_99_v0_: (d_285_v142_).f12})
                d_302_v155_: _dafny.Set
                d_302_v155_ = _dafny.Set({d_100_v1_, _dafny.SeqWithoutIsStrInference([((d_301_v154_)[(d_285_v142_).f12] if ((d_285_v142_).f12) in (d_301_v154_) else -953)])})
                d_168_v45_ = (default__.fm71(d_168_v45_, d_168_v45_, (d_285_v142_).f13, d_168_v45_, globalState)).issubset((d_302_v155_) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_285_v142_).f12, (d_285_v142_).f12, (d_285_v142_).f12, 268, (d_285_v142_).f12])})))
                (globalState).f8 = d_168_v45_
                d_303_v156_: _dafny.Map
                d_303_v156_ = _dafny.Map({((d_285_v142_).f12) <= ((d_285_v142_).f13): d_285_v142_})
                d_303_v156_ = (d_303_v156_).set(d_168_v45_, d_285_v142_)
                d_304_v157_: _dafny.Map
                d_304_v157_ = _dafny.Map({(d_285_v142_).f12: _dafny.SeqWithoutIsStrInference([(d_285_v142_).f13 for d_305_i13_ in range(default__.abs(670))])})
                d_304_v157_ = (d_304_v157_).set(len(d_100_v1_), (_dafny.SeqWithoutIsStrInference([len(d_100_v1_) for d_306_i14_ in range(default__.abs(507))])) + (d_100_v1_))
                (globalState).f8 = (d_297_v152_).fm7(globalState)
            elif True:
                d_99_v0_ = (0) - ((0) - ((d_100_v1_)[default__.safeIndex(59, len(d_100_v1_))]))
                d_168_v45_ = d_168_v45_
                d_307_v158_: _dafny.Map
                d_307_v158_ = _dafny.Map({(d_285_v142_).f12: _dafny.CodePoint('x')})
                d_308_v159_: D6
                d_308_v159_ = D6_DC18(d_168_v45_)
                d_307_v158_ = (d_307_v158_).set((len(_dafny.Map({d_102_v3_: (d_308_v159_).cf32}))) * ((0) - ((d_285_v142_).f12)), d_169_v46_)
                d_168_v45_ = d_168_v45_
                d_309_v160_: T1
                nw43_ = C0()
                nw43_.ctor__(len(d_281_v138_), (d_168_v45_) and (not(d_168_v45_)), (d_285_v142_).f13, (d_285_v142_).f13)
                d_309_v160_ = nw43_
                d_309_v160_ = d_309_v160_
            d_310_v161_: _dafny.Map
            d_310_v161_ = _dafny.Map({d_168_v45_: d_168_v45_})
            d_311_v162_: _dafny.MultiSet
            d_311_v162_ = _dafny.MultiSet([d_310_v161_, d_310_v161_])
            d_312_v163_: _dafny.MultiSet
            d_312_v163_ = d_311_v162_
            d_313_v165_: C1
            nw44_ = C1()
            def iife39_():
                coll39_ = _dafny.Map()
                compr_39_: str
                for compr_39_ in (d_170_v47_).Elements:
                    d_314_v164_: str = compr_39_
                    if (d_314_v164_) in (d_170_v47_):
                        coll39_[d_314_v164_] = (d_285_v142_).f12
                return _dafny.Map(coll39_)
            nw44_.ctor__(len(iife39_()
            ))
            d_313_v165_ = nw44_
            d_315_v166_: _dafny.Array
            def lambda16_(d_316_v1_):
                def lambda17_(d_317_i15_):
                    return d_316_v1_

                return lambda17_

            init8_ = lambda16_(d_100_v1_)
            nw45_ = _dafny.Array(None, 29)
            for i0_8_ in range(nw45_.length(0)):
                nw45_[i0_8_] = init8_(i0_8_)
            d_315_v166_ = nw45_
            index33_ = default__.safeIndex(761, (d_315_v166_).length(0))
            (d_315_v166_)[index33_] = d_100_v1_
            d_318_v167_: C7
            nw46_ = C7()
            nw46_.ctor__((0) - (d_99_v0_), d_99_v0_)
            d_318_v167_ = nw46_
            index34_ = default__.safeIndex(761, (d_315_v166_).length(0))
            rhs25_ = d_312_v163_
            rhs26_ = (d_313_v165_ if not (d_168_v45_) or (d_168_v45_) else d_313_v165_)
            rhs27_ = (_dafny.SeqWithoutIsStrInference([(d_285_v142_).f12, len(_dafny.Map({len(_dafny.Map({d_285_v142_: d_318_v167_})): d_168_v45_})), len(_dafny.Map({(d_285_v142_).f12: d_99_v0_})), (d_285_v142_).f12, (d_285_v142_).f13]))
            lhs15_ = d_315_v166_
            lhs16_ = default__.safeIndex(761, (d_315_v166_).length(0))
            d_312_v163_ = rhs25_
            d_313_v165_ = rhs26_
            lhs15_[lhs16_] = rhs27_
            d_102_v3_ = d_102_v3_
        d_319_v168_: _dafny.Array
        def lambda18_(d_320_v137_):
            def lambda19_(d_321_i16_):
                return d_320_v137_

            return lambda19_

        init9_ = lambda18_(d_280_v137_)
        nw47_ = _dafny.Array(None, 6)
        for i0_9_ in range(nw47_.length(0)):
            nw47_[i0_9_] = init9_(i0_9_)
        d_319_v168_ = nw47_
        r0 = d_319_v168_
        r1 = (not (d_168_v45_) or (True)) == (d_168_v45_)
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_322_v0_: _dafny.Seq
        d_322_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wh"))
        d_323_v1_: int
        d_323_v1_ = -230
        d_324_v2_: _dafny.Seq
        d_324_v2_ = _dafny.SeqWithoutIsStrInference([d_323_v1_])
        d_325_v3_: _dafny.Seq
        d_325_v3_ = _dafny.SeqWithoutIsStrInference([d_324_v2_])
        d_326_v4_: bool
        d_326_v4_ = True
        d_327_v5_: _dafny.Map
        d_327_v5_ = _dafny.Map({19: d_326_v4_})
        d_328_v6_: str
        d_328_v6_ = _dafny.CodePoint('l')
        d_329_v7_: _dafny.Map
        d_329_v7_ = _dafny.Map({d_328_v6_: d_326_v4_})
        d_330_v9_: _dafny.Set
        d_330_v9_ = _dafny.Set({d_328_v6_})
        d_331_v10_: _dafny.Seq
        def iife40_():
            coll40_ = _dafny.Map()
            compr_40_: str
            for compr_40_ in (d_330_v9_).Elements:
                d_332_v8_: str = compr_40_
                if (d_332_v8_) in (d_330_v9_):
                    coll40_[d_332_v8_] = True
            return _dafny.Map(coll40_)
        d_331_v10_ = _dafny.SeqWithoutIsStrInference([iife40_()
        ])
        d_333_v11_: _dafny.Set
        d_333_v11_ = _dafny.Set({_dafny.Map({d_328_v6_: d_326_v4_}), d_329_v7_, _dafny.Map({_dafny.CodePoint('r'): d_326_v4_}), (d_331_v10_)[default__.safeIndex(d_323_v1_, len(d_331_v10_))]})
        d_334_globalState_: GlobalState
        nw48_ = GlobalState()
        nw48_.ctor__(886, 542, False, 596, (d_322_v0_) + (d_322_v0_), (((d_325_v3_)[default__.safeIndex(d_323_v1_, len(d_325_v3_))]).set(default__.safeIndex(d_323_v1_, len((d_325_v3_)[default__.safeIndex(d_323_v1_, len(d_325_v3_))])), d_323_v1_)).set(default__.safeIndex(len(d_327_v5_), len(((d_325_v3_)[default__.safeIndex(d_323_v1_, len(d_325_v3_))]).set(default__.safeIndex(d_323_v1_, len((d_325_v3_)[default__.safeIndex(d_323_v1_, len(d_325_v3_))])), d_323_v1_))), d_323_v1_), 285, 542, True, d_333_v11_, 803)
        d_334_globalState_ = nw48_
        d_335_i0_: int
        d_335_i0_ = 0
        with _dafny.label("0"):
            while d_326_v4_:
                with _dafny.c_label("0"):
                    if (d_335_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_335_i0_ = (d_335_i0_) + (1)
                    d_323_v1_ = d_323_v1_
                    d_336_v12_: _dafny.Array
                    d_337_v13_: bool
                    out0_: _dafny.Array
                    out1_: bool
                    out0_, out1_ = default__.m0(d_334_globalState_)
                    d_336_v12_ = out0_
                    d_337_v13_ = out1_
                    d_338_v14_: _dafny.Seq
                    d_338_v14_ = _dafny.SeqWithoutIsStrInference([d_326_v4_, d_337_v13_, not(d_337_v13_), d_326_v4_, d_337_v13_])
                    d_339_v15_: _dafny.Array
                    nw49_ = _dafny.Array(None, 14)
                    nw49_[int(0)] = d_324_v2_
                    nw49_[int(1)] = _dafny.SeqWithoutIsStrInference([d_323_v1_])
                    nw49_[int(2)] = d_324_v2_
                    nw49_[int(3)] = d_324_v2_
                    nw49_[int(4)] = _dafny.SeqWithoutIsStrInference([d_323_v1_ for d_340_i1_ in range(default__.abs(289))])
                    nw49_[int(5)] = (d_324_v2_).set(default__.safeIndex(d_323_v1_, len(d_324_v2_)), d_323_v1_)
                    nw49_[int(6)] = _dafny.SeqWithoutIsStrInference([d_323_v1_ for d_341_i2_ in range(default__.abs(-352))])
                    nw49_[int(7)] = (d_324_v2_) + (d_324_v2_)
                    nw49_[int(8)] = d_324_v2_
                    nw49_[int(9)] = d_324_v2_
                    nw49_[int(10)] = (d_324_v2_).set(default__.safeIndex(d_323_v1_, len(d_324_v2_)), len(d_338_v14_))
                    nw49_[int(11)] = d_324_v2_
                    nw49_[int(12)] = (d_324_v2_) + (default__.fm0(d_334_globalState_))
                    nw49_[int(13)] = _dafny.SeqWithoutIsStrInference([d_323_v1_, len(d_322_v0_), d_323_v1_, d_323_v1_, d_323_v1_])
                    d_339_v15_ = nw49_
                    d_342_v16_: D0
                    d_342_v16_ = D0_DC0(d_339_v15_)
                    d_339_v15_ = ((d_342_v16_ if d_326_v4_ else d_342_v16_)).cf0
                    d_323_v1_ = d_323_v1_
                    pass
            pass
        d_322_v0_ = d_322_v0_
        d_343_v17_: _dafny.Set
        d_343_v17_ = _dafny.Set({d_326_v4_, d_326_v4_, d_326_v4_})
        d_344_v18_: _dafny.MultiSet
        d_344_v18_ = _dafny.MultiSet([(d_322_v0_).set(default__.safeIndex(len(d_343_v17_), len(d_322_v0_)), d_328_v6_)])
        d_345_v19_: _dafny.Seq
        d_345_v19_ = _dafny.SeqWithoutIsStrInference([d_322_v0_, d_322_v0_])
        d_346_v20_: _dafny.Map
        d_346_v20_ = _dafny.Map({default__.fm1(37, d_326_v4_, d_323_v1_, d_334_globalState_): d_326_v4_})
        d_347_v21_: _dafny.MultiSet
        d_347_v21_ = _dafny.MultiSet([d_323_v1_])
        d_348_v22_: _dafny.Seq
        d_348_v22_ = _dafny.SeqWithoutIsStrInference([d_326_v4_])
        d_349_v23_: _dafny.Array
        nw50_ = _dafny.Array(None, 27)
        nw50_[int(0)] = d_326_v4_
        nw50_[int(1)] = d_326_v4_
        nw50_[int(2)] = d_326_v4_
        nw50_[int(3)] = (d_344_v18_).isdisjoint(_dafny.MultiSet(d_345_v19_))
        nw50_[int(4)] = d_326_v4_
        nw50_[int(5)] = d_326_v4_
        nw50_[int(6)] = d_326_v4_
        nw50_[int(7)] = d_326_v4_
        nw50_[int(8)] = ((d_346_v20_)[d_326_v4_] if (d_326_v4_) in (d_346_v20_) else d_326_v4_)
        nw50_[int(9)] = d_326_v4_
        nw50_[int(10)] = (default__.fm1(d_323_v1_, default__.fm1(d_323_v1_, d_326_v4_, d_323_v1_, d_334_globalState_), d_323_v1_, d_334_globalState_)) or (d_326_v4_)
        nw50_[int(11)] = d_326_v4_
        nw50_[int(12)] = (d_326_v4_) or (d_326_v4_)
        nw50_[int(13)] = d_326_v4_
        nw50_[int(14)] = True
        nw50_[int(15)] = d_326_v4_
        nw50_[int(16)] = d_326_v4_
        nw50_[int(17)] = d_326_v4_
        nw50_[int(18)] = d_326_v4_
        nw50_[int(19)] = (D0_DC3((d_347_v21_).cardinality, True)).cf8
        nw50_[int(20)] = d_326_v4_
        nw50_[int(21)] = (d_348_v22_)[default__.safeIndex(-828, len(d_348_v22_))]
        nw50_[int(22)] = (d_322_v0_) < (d_322_v0_)
        nw50_[int(23)] = d_326_v4_
        nw50_[int(24)] = (d_323_v1_) <= (d_323_v1_)
        nw50_[int(25)] = d_326_v4_
        nw50_[int(26)] = (not(d_326_v4_)) or (default__.fm1(d_323_v1_, d_326_v4_, d_323_v1_, d_334_globalState_))
        d_349_v23_ = nw50_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_349_v23_).length(0)):
            d_350_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_350_i3_)) and ((d_350_i3_) < ((d_349_v23_).length(0)))):
                (d_349_v23_)[(d_350_i3_)] = d_326_v4_
        index35_ = default__.safeIndex(50, (d_349_v23_).length(0))
        (d_349_v23_)[index35_] = d_326_v4_
        index36_ = default__.safeIndex(50, (d_349_v23_).length(0))
        (d_349_v23_)[index36_] = (d_323_v1_) <= ((default__.fm2(d_334_globalState_)) * (d_323_v1_))
        d_351_v24_: _dafny.Map
        d_351_v24_ = _dafny.Map({d_324_v2_: (d_349_v23_)[default__.safeIndex(50, (d_349_v23_).length(0))]})
        d_352_v25_: _dafny.Seq
        d_352_v25_ = _dafny.SeqWithoutIsStrInference([d_351_v24_])
        (d_334_globalState_).f8 = not(not(not (d_326_v4_) or ((_dafny.Map({d_324_v2_: (d_349_v23_)[default__.safeIndex(50, (d_349_v23_).length(0))]})) == ((d_352_v25_)[default__.safeIndex(d_323_v1_, len(d_352_v25_))]))))
        d_328_v6_ = d_328_v6_
        d_323_v1_ = (d_323_v1_ if (d_349_v23_)[default__.safeIndex(50, (d_349_v23_).length(0))] else d_323_v1_)
        hi0_ = default__.safeDivisionInt(d_323_v1_, d_323_v1_)
        for d_353_i4_ in range(d_323_v1_, hi0_):
            d_323_v1_ = 122
            d_348_v22_ = d_348_v22_
            d_354_v26_: _dafny.Set
            d_354_v26_ = _dafny.Set({d_323_v1_, default__.safeDivisionInt(d_353_i4_, d_353_i4_), d_353_i4_, len(d_322_v0_)})
            d_355_v27_: _dafny.Map
            d_355_v27_ = _dafny.Map({d_353_i4_: d_354_v26_})
            d_354_v26_ = (((d_355_v27_)[d_323_v1_] if (d_323_v1_) in (d_355_v27_) else _dafny.Set({d_353_i4_}))).intersection(d_354_v26_)
            d_326_v4_ = (len(default__.fm3(d_334_globalState_))) > (default__.safeDivisionInt(len(_dafny.Map({default__.fm2(d_334_globalState_): d_323_v1_})), d_353_i4_))
        d_323_v1_ = d_323_v1_
        d_348_v22_ = d_348_v22_
        d_356_i5_: int
        d_356_i5_ = 0
        with _dafny.label("1"):
            while d_326_v4_:
                with _dafny.c_label("1"):
                    if (d_356_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_356_i5_ = (d_356_i5_) + (1)
                    d_357_v28_: D0
                    d_357_v28_ = D0_DC3(d_323_v1_, d_326_v4_)
                    d_343_v17_ = default__.fm4(-348, d_357_v28_, default__.safeModuloInt(d_323_v1_, d_323_v1_), _dafny.CodePoint('o'), d_334_globalState_)
                    d_322_v0_ = d_322_v0_
                    index37_ = default__.safeIndex(50, (d_349_v23_).length(0))
                    (d_349_v23_)[index37_] = d_326_v4_
                    d_358_v29_: _dafny.Array
                    nw51_ = _dafny.Array(_dafny.Map({}), 17)
                    d_358_v29_ = nw51_
                    d_359_v30_: _dafny.Map
                    d_359_v30_ = _dafny.Map({d_358_v29_: (58) >= (d_323_v1_)})
                    d_359_v30_ = (d_359_v30_).set(d_358_v29_, not(d_326_v4_))
                    pass
            pass
        d_360_v31_: _dafny.Array
        nw52_ = _dafny.Array(int(0), 27)
        d_360_v31_ = nw52_
        d_360_v31_ = d_360_v31_
        (d_334_globalState_).f8 = (False) == (d_326_v4_)
        d_361_v32_: _dafny.Array
        nw53_ = _dafny.Array(_dafny.Seq({}), 19)
        d_361_v32_ = nw53_
        d_362_v33_: _dafny.Array
        def lambda20_(d_363_v2_):
            def lambda21_(d_364_i6_):
                return d_363_v2_

            return lambda21_

        init10_ = lambda20_(d_324_v2_)
        nw54_ = _dafny.Array(None, 8)
        for i0_10_ in range(nw54_.length(0)):
            nw54_[i0_10_] = init10_(i0_10_)
        d_362_v33_ = nw54_
        d_365_v34_: D0
        d_365_v34_ = D0_DC0(d_362_v33_)
        d_366_v35_: _dafny.Seq
        d_366_v35_ = _dafny.SeqWithoutIsStrInference([d_365_v34_])
        pat_let_tv0_ = d_362_v33_
        pat_let_tv1_ = d_362_v33_
        index38_ = default__.safeIndex(298, (d_361_v32_).length(0))
        def iife41_(_pat_let0_0):
            def iife42_(d_367_dt__update__tmp_h0_):
                def iife43_(_pat_let1_0):
                    def iife44_(d_368_dt__update_hcf0_h0_):
                        return D0_DC0(d_368_dt__update_hcf0_h0_)
                    return iife44_(_pat_let1_0)
                return iife43_(pat_let_tv0_)
            return iife42_(_pat_let0_0)
        def iife45_(_pat_let2_0):
            def iife46_(d_369_dt__update__tmp_h1_):
                def iife47_(_pat_let3_0):
                    def iife48_(d_370_dt__update_hcf0_h1_):
                        return D0_DC0(d_370_dt__update_hcf0_h1_)
                    return iife48_(_pat_let3_0)
                return iife47_(pat_let_tv1_)
            return iife46_(_pat_let2_0)
        (d_361_v32_)[index38_] = (d_366_v35_) + (_dafny.SeqWithoutIsStrInference([iife41_(d_365_v34_), d_365_v34_, iife45_(d_365_v34_)]))
        index39_ = default__.safeIndex(298, (d_361_v32_).length(0))
        (d_361_v32_)[index39_] = _dafny.SeqWithoutIsStrInference([D0_DC0(d_362_v33_), d_365_v34_])
        d_323_v1_ = (0) - (((len(d_343_v17_)) * (d_323_v1_)) * (d_323_v1_))
        d_371_v36_: _dafny.Array
        def lambda22_(d_372_v22_):
            def lambda23_(d_373_i7_):
                return _dafny.MultiSet([d_372_v22_, d_372_v22_])

            return lambda23_

        init11_ = lambda22_(d_348_v22_)
        nw55_ = _dafny.Array(None, 13)
        for i0_11_ in range(nw55_.length(0)):
            nw55_[i0_11_] = init11_(i0_11_)
        d_371_v36_ = nw55_
        d_374_v37_: _dafny.MultiSet
        d_374_v37_ = _dafny.MultiSet([d_348_v22_])
        index40_ = default__.safeIndex(915, (d_371_v36_).length(0))
        (d_371_v36_)[index40_] = d_374_v37_
        index41_ = default__.safeIndex(915, (d_371_v36_).length(0))
        rhs28_ = d_326_v4_
        rhs29_ = d_374_v37_
        rhs30_ = d_323_v1_
        lhs17_ = d_334_globalState_
        lhs18_ = d_371_v36_
        lhs19_ = default__.safeIndex(915, (d_371_v36_).length(0))
        lhs17_.f8 = rhs28_
        lhs18_[lhs19_] = rhs29_
        d_323_v1_ = rhs30_
        _dafny.print((d_322_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_323_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_324_v2_) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_325_v3_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-230])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_326_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v5_) == (_dafny.Map({19: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_328_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_329_v7_) == (_dafny.Map({_dafny.CodePoint('l'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_330_v9_) == (_dafny.Set({_dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_331_v10_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('l'): True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v11_) == (_dafny.Set({_dafny.Map({_dafny.CodePoint('l'): True}), _dafny.Map({_dafny.CodePoint('r'): True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_globalState_).f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_.f5) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_334_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_334_globalState_).f9) == (_dafny.Set({_dafny.Map({_dafny.CodePoint('l'): True}), _dafny.Map({_dafny.CodePoint('r'): True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_334_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_335_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v17_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_344_v18_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wl"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_345_v19_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wh"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_346_v20_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_347_v21_) == (_dafny.MultiSet([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_348_v22_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v23_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_351_v24_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([-230]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v25_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.SeqWithoutIsStrInference([-230]): False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_356_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_361_v32_)[13])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[0]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[1]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[2]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[3]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[4]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[5]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[6]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_362_v33_)[7]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[0]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[1]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[2]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[3]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[4]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[5]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[6]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_365_v34_).cf0)[7]) == (_dafny.SeqWithoutIsStrInference([-230]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_366_v35_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[0]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[1]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[2]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[3]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[4]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[5]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[6]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[7]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[8]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[9]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[10]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[11]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_371_v36_)[12]) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v37_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {self.cf2.VerbatimString(True)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
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
        return lambda: D1_DC5(_dafny.MultiSet({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC5(D1, NamedTuple('DC5', [('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0), _dafny.Map({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(False, int(0), int(0))
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

class D4_DC13(D4, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf24 == __o.cf24
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
        return lambda: D5_DC15(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC22(int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)

class D7_DC22(D7, NamedTuple('DC22', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC23(D7, NamedTuple('DC23', [('cf42', Any), ('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC24(D7, NamedTuple('DC24', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC25(D8, NamedTuple('DC25', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27(False, D7.default()(), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)

class D9_DC27(D9, NamedTuple('DC27', [('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC28(D9, NamedTuple('DC28', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC29(D10, NamedTuple('DC29', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC33(D12, NamedTuple('DC33', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC37()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)

class D13_DC37(D13, NamedTuple('DC37', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)

class D14_DC39(D14, NamedTuple('DC39', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC41(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D15_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)

class D15_DC41(D15, NamedTuple('DC41', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC41({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC41) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)

class D16_DC43(D16, NamedTuple('DC43', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC45(int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC45(D17, NamedTuple('DC45', [('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC44(D17, NamedTuple('DC44', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC48()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D18_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)

class D18_DC48(D18, NamedTuple('DC48', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC48'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC48)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC49(D18, NamedTuple('DC49', [('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC47(D18, NamedTuple('DC47', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC51(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D19_DC50)

class D19_DC51(D19, NamedTuple('DC51', [('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC50(D19, NamedTuple('DC50', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC50({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC50) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC53()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)

class D20_DC53(D20, NamedTuple('DC53', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC52(D20, NamedTuple('DC52', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC56(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)

class D21_DC56(D21, NamedTuple('DC56', [('cf82', Any), ('cf83', Any), ('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC57(D21, NamedTuple('DC57', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)

class D22_DC58(D22, NamedTuple('DC58', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf86 == __o.cf86
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

class D23_DC59(D23, NamedTuple('DC59', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC61(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D24_DC61)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)

class D24_DC61(D24, NamedTuple('DC61', [('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC61({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC61) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC60(D24, NamedTuple('DC60', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D25_DC62)

class D25_DC62(D25, NamedTuple('DC62', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC62({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC62) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D26_DC63)

class D26_DC63(D26, NamedTuple('DC63', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC63({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC63) and self.cf92 == __o.cf92
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

class T1:
    pass

class T2:
    pass
    def m11(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f5: _dafny.Seq = _dafny.Seq({})
        self.f8: bool = False
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f2: bool = False
        self._f3: int = int(0)
        self._f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f6: int = int(0)
        self._f7: int = int(0)
        self._f9: _dafny.Set = _dafny.Set({})
        self._f10: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10

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

class C0(T1):
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: int = int(0)
        self.f29: bool = False
        self._f28: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f28, f29, f12, f13):
        (self)._f28 = f28
        (self).f29 = f29
        (self)._f12 = f12
        (self)._f13 = f13

    def fm7(self, globalState):
        return self.f29

    def fm8(self, p0, p1, p2, globalState):
        return ((D0_DC1(68, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqdyj")), len(_dafny.Set({self.f29})), self.f29)).cf2) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgplbwu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boygyd"))))

    @property
    def f28(self):
        return self._f28

class C1:
    def  __init__(self):
        self.f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f27):
        (self).f27 = f27

    def fm34(self, globalState):
        return (default__.safeDivisionInt(self.f27, self.f27)) - (self.f27)

    def fm35(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, True])), _dafny.MultiSet([True])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False]), _dafny.MultiSet([False, True, False])])))

    def m24(self, globalState):
        d_375_v0_: _dafny.Seq
        d_375_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbgh"))
        d_376_v1_: str
        d_376_v1_ = _dafny.CodePoint('y')
        d_377_v2_: _dafny.Array
        nw56_ = _dafny.Array(None, 28)
        nw56_[int(0)] = d_375_v0_
        nw56_[int(1)] = d_375_v0_
        nw56_[int(2)] = d_375_v0_
        nw56_[int(3)] = d_375_v0_
        nw56_[int(4)] = d_375_v0_
        nw56_[int(5)] = (d_375_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_378_i0_ in range(default__.abs(931))]))
        nw56_[int(6)] = d_375_v0_
        nw56_[int(7)] = d_375_v0_
        nw56_[int(8)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_379_i1_ in range(default__.abs(229))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "deoknfa")))
        nw56_[int(9)] = (d_375_v0_) + (d_375_v0_)
        nw56_[int(10)] = d_375_v0_
        nw56_[int(11)] = d_375_v0_
        nw56_[int(12)] = d_375_v0_
        nw56_[int(13)] = d_375_v0_
        nw56_[int(14)] = default__.fm36(self.f27, len(d_375_v0_), d_376_v1_, globalState)
        nw56_[int(15)] = d_375_v0_
        nw56_[int(16)] = d_375_v0_
        nw56_[int(17)] = d_375_v0_
        nw56_[int(18)] = d_375_v0_
        nw56_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qscv"))
        nw56_[int(20)] = (d_375_v0_) + (d_375_v0_)
        nw56_[int(21)] = d_375_v0_
        nw56_[int(22)] = d_375_v0_
        nw56_[int(23)] = (d_375_v0_) + (d_375_v0_)
        nw56_[int(24)] = d_375_v0_
        nw56_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "beqc"))
        nw56_[int(26)] = d_375_v0_
        nw56_[int(27)] = d_375_v0_
        d_377_v2_ = nw56_
        index42_ = default__.safeIndex(599, (d_377_v2_).length(0))
        (d_377_v2_)[index42_] = d_375_v0_
        d_380_v3_: _dafny.Map
        d_380_v3_ = _dafny.Map({d_376_v1_: d_375_v0_})
        d_381_v4_: bool
        d_381_v4_ = True
        index43_ = default__.safeIndex(599, (d_377_v2_).length(0))
        (d_377_v2_)[index43_] = (((d_380_v3_)[d_376_v1_] if (d_376_v1_) in (d_380_v3_) else default__.fm36(self.f27, (0) - (len(_dafny.Set({d_381_v4_}))), d_376_v1_, globalState))) + (d_375_v0_)
        d_382_v5_: D4
        d_382_v5_ = D4_DC13(self.f27)
        d_382_v5_ = d_382_v5_
        d_383_v6_: C0
        nw57_ = C0()
        nw57_.ctor__(self.f27, d_381_v4_, default__.safeModuloInt(self.f27, self.f27), self.f27)
        d_383_v6_ = nw57_
        d_384_v7_: _dafny.Map
        d_384_v7_ = _dafny.Map({not(d_381_v4_): d_383_v6_.f29})
        if ((d_384_v7_)[d_383_v6_.f29] if (d_383_v6_.f29) in (d_384_v7_) else ((d_383_v6_).f28) >= ((d_383_v6_).f28)):
            d_385_v8_: _dafny.Seq
            d_385_v8_ = _dafny.SeqWithoutIsStrInference([d_383_v6_.f29, d_383_v6_.f29])
            d_386_v9_: _dafny.Seq
            d_386_v9_ = _dafny.SeqWithoutIsStrInference([len(d_385_v8_), len((d_384_v7_).set(d_381_v4_, d_383_v6_.f29))])
            (self).f27 = default__.safeDivisionInt(len((d_386_v9_) + (d_386_v9_)), (d_383_v6_).f28)
            index44_ = default__.safeIndex(599, (d_377_v2_).length(0))
            (d_377_v2_)[index44_] = d_375_v0_
            d_387_v10_: _dafny.Array
            def lambda24_(d_388_v6_):
                def lambda25_(d_389_i2_):
                    return (d_389_i2_) + ((d_388_v6_).f28)

                return lambda25_

            init12_ = lambda24_(d_383_v6_)
            nw58_ = _dafny.Array(None, 11)
            for i0_12_ in range(nw58_.length(0)):
                nw58_[i0_12_] = init12_(i0_12_)
            d_387_v10_ = nw58_
            index45_ = default__.safeIndex(44, (d_387_v10_).length(0))
            (d_387_v10_)[index45_] = self.f27
            index46_ = default__.safeIndex(44, (d_387_v10_).length(0))
            (d_387_v10_)[index46_] = (d_386_v9_)[default__.safeIndex(len(d_385_v8_), len(d_386_v9_))]
            d_390_v11_: _dafny.Map
            d_390_v11_ = _dafny.Map({d_385_v8_: d_383_v6_.f29})
            d_390_v11_ = (d_390_v11_).set(d_385_v8_, d_381_v4_)
            if d_383_v6_.f29:
                index47_ = default__.safeIndex(599, (d_377_v2_).length(0))
                (d_377_v2_)[index47_] = (d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))]
                (d_383_v6_).f29 = d_383_v6_.f29
                index48_ = default__.safeIndex(599, (d_377_v2_).length(0))
                (d_377_v2_)[index48_] = (d_375_v0_) + ((d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))])
                d_391_v12_: _dafny.Array
                def lambda26_(d_392_v6_):
                    def lambda27_(d_393_i3_):
                        return d_392_v6_.f29

                    return lambda27_

                init13_ = lambda26_(d_383_v6_)
                nw59_ = _dafny.Array(None, 6)
                for i0_13_ in range(nw59_.length(0)):
                    nw59_[i0_13_] = init13_(i0_13_)
                d_391_v12_ = nw59_
                d_394_v13_: _dafny.MultiSet
                d_394_v13_ = _dafny.MultiSet([(d_381_v4_) and (d_383_v6_.f29)])
                d_395_v14_: _dafny.Map
                d_395_v14_ = _dafny.Map({416: self.f27})
                d_396_v15_: _dafny.MultiSet
                d_396_v15_ = _dafny.MultiSet([(0) - ((D3_DC9((d_383_v6_).f28, _dafny.Map({(0) - ((d_387_v10_)[default__.safeIndex(44, (d_387_v10_).length(0))]): d_395_v14_}), (d_383_v6_).f28, (d_387_v10_)[default__.safeIndex(44, (d_387_v10_).length(0))])).cf17)])
                d_397_v16_: _dafny.Map
                d_397_v16_ = _dafny.Map({d_396_v15_: (d_383_v6_).f28})
                def iife49_():
                    coll41_ = _dafny.Set()
                    compr_41_: _dafny.MultiSet
                    for compr_41_ in (d_397_v16_).keys.Elements:
                        d_398_v17_: _dafny.MultiSet = compr_41_
                        if (d_398_v17_) in (d_397_v16_):
                            coll41_ = coll41_.union(_dafny.Set([d_398_v17_]))
                    return _dafny.Set(coll41_)
                rhs31_ = d_391_v12_
                rhs32_ = d_391_v12_
                rhs33_ = (_dafny.MultiSet([d_381_v4_])) | ((d_394_v13_) - (_dafny.MultiSet(d_385_v8_)))
                rhs34_ = (d_375_v0_)[default__.safeIndex(((d_383_v6_).f28) + (len(iife49_()
                )), len(d_375_v0_))]
                d_391_v12_ = rhs31_
                d_391_v12_ = rhs32_
                d_394_v13_ = rhs33_
                d_376_v1_ = rhs34_
                (self).f27 = -182
            elif True:
                d_399_v18_: _dafny.Set
                d_399_v18_ = _dafny.Set({d_383_v6_.f29})
                d_400_v19_: _dafny.Seq
                d_400_v19_ = _dafny.SeqWithoutIsStrInference([d_399_v18_])
                d_401_v20_: _dafny.Map
                d_401_v20_ = _dafny.Map({d_399_v18_: d_383_v6_.f29})
                d_402_v21_: _dafny.MultiSet
                d_402_v21_ = _dafny.MultiSet([(d_387_v10_)[default__.safeIndex(44, (d_387_v10_).length(0))], len(d_401_v20_)])
                d_403_v22_: _dafny.Array
                nw60_ = _dafny.Array(None, 11)
                nw60_[int(0)] = False
                nw60_[int(1)] = d_383_v6_.f29
                nw60_[int(2)] = ((d_400_v19_)[default__.safeIndex((d_383_v6_).f28, len(d_400_v19_))]).issubset(d_399_v18_)
                nw60_[int(3)] = (d_399_v18_) == (d_399_v18_)
                nw60_[int(4)] = (d_399_v18_).issubset(d_399_v18_)
                nw60_[int(5)] = not(default__.fm1((d_383_v6_).f28, d_383_v6_.f29, (d_387_v10_)[default__.safeIndex(44, (d_387_v10_).length(0))], globalState))
                nw60_[int(6)] = d_381_v4_
                nw60_[int(7)] = (d_402_v21_).ispropersubset(d_402_v21_)
                nw60_[int(8)] = d_383_v6_.f29
                nw60_[int(9)] = d_383_v6_.f29
                nw60_[int(10)] = not(not(d_381_v4_))
                d_403_v22_ = nw60_
                index49_ = default__.safeIndex(151, (d_403_v22_).length(0))
                (d_403_v22_)[index49_] = ((d_383_v6_).f28) != ((d_383_v6_).f28)
                index50_ = default__.safeIndex(151, (d_403_v22_).length(0))
                rhs35_ = not(((d_383_v6_).f28) >= ((d_387_v10_)[default__.safeIndex(44, (d_387_v10_).length(0))]))
                rhs36_ = self.f27
                lhs20_ = d_403_v22_
                lhs21_ = default__.safeIndex(151, (d_403_v22_).length(0))
                lhs22_ = self
                lhs20_[lhs21_] = rhs35_
                lhs22_.f27 = rhs36_
                d_404_v23_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_404_v23_ = nw61_
                d_404_v23_ = d_404_v23_
                d_387_v10_ = d_387_v10_
                d_377_v2_ = d_377_v2_
                d_405_v24_: _dafny.Map
                d_405_v24_ = _dafny.Map({d_383_v6_.f29: (d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))]})
                rhs37_ = (d_381_v4_) not in ((d_405_v24_) | (d_405_v24_))
                rhs38_ = d_383_v6_.f29
                lhs23_ = globalState
                lhs23_.f8 = rhs37_
                d_381_v4_ = rhs38_
        elif True:
            d_406_v25_: _dafny.Array
            nw62_ = _dafny.Array(False, 6)
            d_406_v25_ = nw62_
            index51_ = default__.safeIndex(311, (d_406_v25_).length(0))
            (d_406_v25_)[index51_] = (d_383_v6_).fm7(globalState)
            d_407_v26_: D0
            d_407_v26_ = D0_DC3(len(_dafny.SeqWithoutIsStrInference([d_383_v6_.f29, True])), False)
            d_408_v27_: _dafny.Map
            d_408_v27_ = _dafny.Map({d_407_v26_: d_383_v6_.f29})
            d_409_v28_: _dafny.Seq
            d_409_v28_ = _dafny.SeqWithoutIsStrInference([(d_383_v6_).f28])
            d_410_v29_: _dafny.Seq
            d_410_v29_ = _dafny.SeqWithoutIsStrInference([((d_408_v27_)[D0_DC3(len(((d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))]).set(default__.safeIndex(self.f27, len((d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))])), d_376_v1_)), default__.fm1(len(_dafny.Map({True: len(d_409_v28_)})), d_383_v6_.f29, (d_383_v6_).f28, globalState))] if (D0_DC3(len(((d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))]).set(default__.safeIndex(self.f27, len((d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))])), d_376_v1_)), default__.fm1(len(_dafny.Map({True: len(d_409_v28_)})), d_383_v6_.f29, (d_383_v6_).f28, globalState))) in (d_408_v27_) else d_383_v6_.f29), d_381_v4_, d_383_v6_.f29])
            index52_ = default__.safeIndex(311, (d_406_v25_).length(0))
            (d_406_v25_)[index52_] = (d_410_v29_)[default__.safeIndex((self.f27) + ((d_383_v6_).f28), len(d_410_v29_))]
            if ((d_377_v2_)[default__.safeIndex(599, (d_377_v2_).length(0))]) <= (d_375_v0_):
                d_411_v30_: _dafny.Map
                d_411_v30_ = _dafny.Map({d_383_v6_.f29: (d_383_v6_).f28})
                d_412_v31_: C0
                nw63_ = C0()
                nw63_.ctor__((d_383_v6_).f28, d_381_v4_, (d_383_v6_).f28, len((d_411_v30_) | (d_411_v30_)))
                d_412_v31_ = nw63_
                d_406_v25_ = d_406_v25_
                d_413_v32_: _dafny.Array
                def lambda28_(d_414_v31_):
                    def lambda29_(d_415_i4_):
                        return (d_415_i4_) * ((d_414_v31_).f28)

                    return lambda29_

                init14_ = lambda28_(d_412_v31_)
                nw64_ = _dafny.Array(None, 14)
                for i0_14_ in range(nw64_.length(0)):
                    nw64_[i0_14_] = init14_(i0_14_)
                d_413_v32_ = nw64_
                d_416_v33_: _dafny.Map
                d_416_v33_ = _dafny.Map({d_383_v6_.f29: d_413_v32_})
                d_417_v34_: _dafny.Map
                d_417_v34_ = _dafny.Map({d_416_v33_: d_406_v25_})
                d_417_v34_ = (d_417_v34_).set(d_416_v33_, d_406_v25_)
                d_418_v35_: _dafny.Array
                nw65_ = _dafny.Array(None, 4)
                nw65_[int(0)] = (d_410_v29_) + (_dafny.SeqWithoutIsStrInference([(d_406_v25_)[default__.safeIndex(311, (d_406_v25_).length(0))]]))
                nw65_[int(1)] = _dafny.SeqWithoutIsStrInference([False])
                nw65_[int(2)] = (d_410_v29_).set(default__.safeIndex(self.f27, len(d_410_v29_)), True)
                nw65_[int(3)] = d_410_v29_
                d_418_v35_ = nw65_
                index53_ = default__.safeIndex(310, (d_418_v35_).length(0))
                (d_418_v35_)[index53_] = (d_410_v29_) + (d_410_v29_)
                d_419_v36_: _dafny.Set
                d_419_v36_ = _dafny.Set({d_381_v4_})
                index54_ = default__.safeIndex(310, (d_418_v35_).length(0))
                index55_ = default__.safeIndex(311, (d_406_v25_).length(0))
                rhs39_ = ((d_383_v6_).f28) - ((d_383_v6_).f28)
                rhs40_ = (d_410_v29_) + ((d_410_v29_) + (d_410_v29_))
                rhs41_ = (d_383_v6_.f29) or (default__.fm1(self.f27, False, len(d_419_v36_), globalState))
                lhs24_ = self
                lhs25_ = d_418_v35_
                lhs26_ = default__.safeIndex(310, (d_418_v35_).length(0))
                lhs27_ = d_406_v25_
                lhs28_ = default__.safeIndex(311, (d_406_v25_).length(0))
                lhs24_.f27 = rhs39_
                lhs25_[lhs26_] = rhs40_
                lhs27_[lhs28_] = rhs41_
                d_420_v37_: _dafny.MultiSet
                d_420_v37_ = _dafny.MultiSet([(d_412_v31_).f28])
                d_420_v37_ = (d_420_v37_) - (((D1_DC5(_dafny.MultiSet([(d_383_v6_).f28]), d_381_v4_)).cf10) - (_dafny.MultiSet([(d_383_v6_).f28])))
            elif True:
                d_421_v38_: _dafny.Array
                def lambda30_(d_422_v6_):
                    def lambda31_(d_423_i5_):
                        return default__.safeModuloInt(d_423_i5_, (d_422_v6_).f28)

                    return lambda31_

                init15_ = lambda30_(d_383_v6_)
                nw66_ = _dafny.Array(None, 24)
                for i0_15_ in range(nw66_.length(0)):
                    nw66_[i0_15_] = init15_(i0_15_)
                d_421_v38_ = nw66_
                d_424_v39_: D0
                d_424_v39_ = D0_DC2(len(d_409_v28_), d_421_v38_)
                (self).f27 = (d_424_v39_).cf5
                d_425_v40_: _dafny.Map
                d_425_v40_ = _dafny.Map({self.f27: (d_383_v6_).f28})
                pat_let_tv2_ = d_425_v40_
                d_426_v41_: _dafny.Seq
                def iife50_(_pat_let4_0):
                    def iife51_(d_427_dt__update__tmp_h0_):
                        def iife52_(_pat_let5_0):
                            def iife53_(d_428_dt__update_hcf20_h0_):
                                return D4_DC11(d_428_dt__update_hcf20_h0_)
                            return iife53_(_pat_let5_0)
                        return iife52_(pat_let_tv2_)
                    return iife51_(_pat_let4_0)
                d_426_v41_ = _dafny.SeqWithoutIsStrInference([iife50_(D4_DC11(d_425_v40_))])
                d_429_v42_: _dafny.MultiSet
                d_429_v42_ = _dafny.MultiSet([d_383_v6_.f29])
                d_430_v43_: D4
                d_430_v43_ = D4_DC11(_dafny.Map({(d_429_v42_).cardinality: (d_383_v6_).f28}))
                d_426_v41_ = ((d_426_v41_) + (_dafny.SeqWithoutIsStrInference([D4_DC11(d_425_v40_) for d_431_i6_ in range(default__.abs(81))]))) + ((d_426_v41_) + ((d_426_v41_).set(default__.safeIndex(self.f27, len(d_426_v41_)), d_430_v43_)))
                d_432_v44_: _dafny.Array
                nw67_ = _dafny.Array(_dafny.Map({}), 14)
                d_432_v44_ = nw67_
                d_433_v45_: _dafny.Map
                d_433_v45_ = _dafny.Map({d_410_v29_: (d_383_v6_).f28})
                index56_ = default__.safeIndex(417, (d_432_v44_).length(0))
                (d_432_v44_)[index56_] = d_433_v45_
                index57_ = default__.safeIndex(417, (d_432_v44_).length(0))
                (d_432_v44_)[index57_] = (d_433_v45_) | (d_433_v45_)
                d_434_v46_: C0
                nw68_ = C0()
                nw68_.ctor__(self.f27, d_383_v6_.f29, len(d_409_v28_), -556)
                d_434_v46_ = nw68_
                (self).f27 = (-286) - ((self.f27) + ((d_383_v6_).f28))
            d_435_v47_: D0
            d_435_v47_ = D0_DC1((d_383_v6_).f28, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbumnuk")), 285, d_381_v4_)
            index58_ = default__.safeIndex(599, (d_377_v2_).length(0))
            (d_377_v2_)[index58_] = ((d_375_v0_) + (((d_435_v47_).cf2) + (d_375_v0_))).set(default__.safeIndex(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_383_v6_).f28: d_383_v6_.f29})) for d_436_i7_ in range(default__.abs(644))]))})), len((d_375_v0_) + (((d_435_v47_).cf2) + (d_375_v0_)))), d_376_v1_)
            d_437_v48_: _dafny.Set
            d_437_v48_ = _dafny.Set({self.f27, 921, self.f27, (self.f27) * ((d_383_v6_).f28)})
            d_437_v48_ = d_437_v48_
            d_438_v49_: _dafny.Array
            def lambda32_(d_439_i8_):
                return (d_439_i8_) + (self.f27)

            init16_ = lambda32_
            nw69_ = _dafny.Array(None, 28)
            for i0_16_ in range(nw69_.length(0)):
                nw69_[i0_16_] = init16_(i0_16_)
            d_438_v49_ = nw69_
            d_440_v50_: _dafny.Seq
            d_440_v50_ = _dafny.SeqWithoutIsStrInference([d_438_v49_, d_438_v49_, d_438_v49_])
            d_441_v51_: D1
            d_441_v51_ = D1_DC4(d_440_v50_)
            d_442_v52_: _dafny.Map
            d_442_v52_ = _dafny.Map({(d_383_v6_).f28: d_441_v51_})
            d_442_v52_ = (d_442_v52_).set((d_383_v6_).f28, d_441_v51_)
        d_443_v53_: _dafny.Map
        d_443_v53_ = _dafny.Map({self.f27: d_383_v6_.f29})
        d_444_v54_: _dafny.Set
        d_444_v54_ = _dafny.Set({d_383_v6_.f29})
        d_445_v55_: D0
        d_445_v55_ = D0_DC1(len(d_443_v53_), d_375_v0_, len(d_444_v54_), d_383_v6_.f29)
        d_446_v56_: _dafny.MultiSet
        d_446_v56_ = _dafny.MultiSet([self.f27])
        d_447_v57_: _dafny.Seq
        d_447_v57_ = _dafny.SeqWithoutIsStrInference([d_446_v56_, d_446_v56_])
        pat_let_tv3_ = d_375_v0_
        pat_let_tv4_ = d_377_v2_
        pat_let_tv5_ = d_377_v2_
        pat_let_tv6_ = d_377_v2_
        pat_let_tv7_ = d_377_v2_
        pat_let_tv8_ = d_375_v0_
        pat_let_tv9_ = d_377_v2_
        pat_let_tv10_ = d_377_v2_
        pat_let_tv11_ = d_377_v2_
        pat_let_tv12_ = d_377_v2_
        def lambda33_(source6_):
            if source6_.is_DC1:
                d_448___mcc_h0_ = source6_.cf1
                d_449___mcc_h1_ = source6_.cf2
                d_450___mcc_h2_ = source6_.cf3
                d_451___mcc_h3_ = source6_.cf4
                d_452_cf4_ = d_451___mcc_h3_
                d_453_cf3_ = d_450___mcc_h2_
                d_454_cf2_ = d_449___mcc_h1_
                d_455_cf1_ = d_448___mcc_h0_
                return pat_let_tv3_
            elif source6_.is_DC2:
                d_456___mcc_h4_ = source6_.cf5
                d_457___mcc_h5_ = source6_.cf6
                d_458_cf6_ = d_457___mcc_h5_
                d_459_cf5_ = d_456___mcc_h4_
                return (pat_let_tv5_)[default__.safeIndex(599, (pat_let_tv4_).length(0))]
            elif source6_.is_DC3:
                d_460___mcc_h6_ = source6_.cf7
                d_461___mcc_h7_ = source6_.cf8
                d_462_cf8_ = d_461___mcc_h7_
                d_463_cf7_ = d_460___mcc_h6_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            elif True:
                d_464___mcc_h8_ = source6_.cf0
                d_465_cf0_ = d_464___mcc_h8_
                return (pat_let_tv7_)[default__.safeIndex(599, (pat_let_tv6_).length(0))]

        def lambda34_(source7_):
            if source7_.is_DC1:
                d_466___mcc_h9_ = source7_.cf1
                d_467___mcc_h10_ = source7_.cf2
                d_468___mcc_h11_ = source7_.cf3
                d_469___mcc_h12_ = source7_.cf4
                d_470_cf4_ = d_469___mcc_h12_
                d_471_cf3_ = d_468___mcc_h11_
                d_472_cf2_ = d_467___mcc_h10_
                d_473_cf1_ = d_466___mcc_h9_
                return pat_let_tv8_
            elif source7_.is_DC2:
                d_474___mcc_h13_ = source7_.cf5
                d_475___mcc_h14_ = source7_.cf6
                d_476_cf6_ = d_475___mcc_h14_
                d_477_cf5_ = d_474___mcc_h13_
                return (pat_let_tv10_)[default__.safeIndex(599, (pat_let_tv9_).length(0))]
            elif source7_.is_DC3:
                d_478___mcc_h15_ = source7_.cf7
                d_479___mcc_h16_ = source7_.cf8
                d_480_cf8_ = d_479___mcc_h16_
                d_481_cf7_ = d_478___mcc_h15_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            elif True:
                d_482___mcc_h17_ = source7_.cf0
                d_483_cf0_ = d_482___mcc_h17_
                return (pat_let_tv12_)[default__.safeIndex(599, (pat_let_tv11_).length(0))]

        d_375_v0_ = (lambda33_(d_445_v55_)).set(default__.safeIndex((((d_446_v56_).set(self.f27, default__.abs((_dafny.MultiSet(d_447_v57_)).cardinality))).cardinality) * ((d_383_v6_).f28), len(lambda34_(d_445_v55_))), d_376_v1_)
        d_484_v58_: _dafny.Array
        nw70_ = _dafny.Array(int(0), 21)
        d_484_v58_ = nw70_
        index59_ = default__.safeIndex(999, (d_484_v58_).length(0))
        (d_484_v58_)[index59_] = (d_383_v6_).f28
        d_485_v59_: _dafny.Array
        def lambda35_(d_486_v6_):
            def lambda36_(d_487_i9_):
                return (d_486_v6_.f29) or (d_486_v6_.f29)

            return lambda36_

        init17_ = lambda35_(d_383_v6_)
        nw71_ = _dafny.Array(None, 17)
        for i0_17_ in range(nw71_.length(0)):
            nw71_[i0_17_] = init17_(i0_17_)
        d_485_v59_ = nw71_
        d_488_v60_: _dafny.Seq
        d_488_v60_ = _dafny.SeqWithoutIsStrInference([(d_443_v53_) == (d_443_v53_), (d_375_v0_) < (d_375_v0_), d_381_v4_, not (d_383_v6_.f29) or (d_381_v4_), (self.f27) != (971)])
        index60_ = default__.safeIndex(999, (d_484_v58_).length(0))
        rhs42_ = len(d_488_v60_)
        rhs43_ = d_485_v59_
        rhs44_ = (0) - ((d_383_v6_).f28)
        lhs29_ = d_484_v58_
        lhs30_ = default__.safeIndex(999, (d_484_v58_).length(0))
        lhs31_ = self
        lhs29_[lhs30_] = rhs42_
        d_485_v59_ = rhs43_
        lhs31_.f27 = rhs44_


class C2(T2, T0, T1):
    def  __init__(self):
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: int = int(0)
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f26: int = int(0)
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
    @property
    def f13(self):
        return self._f13
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f26, f20, f12, f13, f11):
        (self)._f26 = f26
        (self)._f20 = f20
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f11 = f11

    def fm18(self, p0, p1, p2, p3, globalState):
        return ((self).f12) > (((_dafny.MultiSet([self.f11, self.f11])).intersection(_dafny.MultiSet([self.f11, self.f11]))).cardinality)

    def fm7(self, globalState):
        return self.f11

    def fm8(self, p0, p1, p2, globalState):
        return ((self).f20) + (((self).f20) + ((self).f20))

    def fm5(self, p0, globalState):
        source8_ = (D1_DC5(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f12, (self).f26, (self).f13])), self.f11) if False else D1_DC5(_dafny.MultiSet([len(_dafny.Map({self.f11: (self).f13})), -86]), False))
        if source8_.is_DC5:
            d_489___mcc_h0_ = source8_.cf10
            d_490___mcc_h1_ = source8_.cf11
            d_491_cf11_ = d_490___mcc_h1_
            d_492_cf10_ = d_489___mcc_h0_
            def iife54_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(129, -59):
                    d_493_v0_: int = compr_42_
                    if ((129) <= (d_493_v0_)) and ((d_493_v0_) < (-59)):
                        coll42_[(d_493_v0_) * (750)] = self.f11
                return _dafny.Map(coll42_)
            return D0_DC1(len(iife54_()
), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nthyvsrum")), 788, d_491_cf11_)
        elif source8_.is_DC4:
            d_494___mcc_h2_ = source8_.cf9
            d_495_cf9_ = d_494___mcc_h2_
            return D0_DC1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "au"))), (self).f20, (self).f26, self.f11)
        elif True:
            d_496___mcc_h3_ = source8_.cf12
            d_497_cf12_ = d_496___mcc_h3_
            return D0_DC1((self).f26, (self).f20, 575, not(self.f11))

    def fm6(self, p0, p1, p2, globalState):
        return not(self.f11)

    def fm32(self, p0, globalState):
        source9_ = D0_DC3(len(((self).f20).set(default__.safeIndex((self).f12, len((self).f20)), _dafny.CodePoint('o'))), self.f11)
        if source9_.is_DC1:
            d_498___mcc_h0_ = source9_.cf1
            d_499___mcc_h1_ = source9_.cf2
            d_500___mcc_h2_ = source9_.cf3
            d_501___mcc_h3_ = source9_.cf4
            d_502_cf4_ = d_501___mcc_h3_
            d_503_cf3_ = d_500___mcc_h2_
            d_504_cf2_ = d_499___mcc_h1_
            d_505_cf1_ = d_498___mcc_h0_
            if d_502_cf4_:
                def iife55_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(65, 338):
                        d_506_v0_: int = compr_43_
                        if ((65) <= (d_506_v0_)) and ((d_506_v0_) < (338)):
                            coll43_[default__.safeModuloInt(d_506_v0_, (self).f26)] = (self).f26
                    return _dafny.Map(coll43_)
                return _dafny.Map({iife55_()
                : self.f11})
            elif True:
                return _dafny.Map({_dafny.Map({d_503_cf3_: len(_dafny.Map({False: (self).f12}))}): self.f11})
        elif source9_.is_DC2:
            d_507___mcc_h4_ = source9_.cf5
            d_508___mcc_h5_ = source9_.cf6
            d_509_cf6_ = d_508___mcc_h5_
            d_510_cf5_ = d_507___mcc_h4_
            return _dafny.Map({_dafny.Map({(_dafny.MultiSet([True, self.f11, True, self.f11, self.f11])).cardinality: -522}): self.f11})
        elif source9_.is_DC3:
            d_511___mcc_h6_ = source9_.cf7
            d_512___mcc_h7_ = source9_.cf8
            d_513_cf8_ = d_512___mcc_h7_
            d_514_cf7_ = d_511___mcc_h6_
            def iife56_():
                def iife58_():
                    coll46_ = _dafny.Map()
                    compr_46_: _dafny.Map
                    for compr_46_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_514_cf7_: (self).f13})])).Elements:
                        d_515_v2_: _dafny.Map = compr_46_
                        if (d_515_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_514_cf7_: (self).f13})])):
                            coll46_[d_515_v2_] = False
                    return _dafny.Map(coll46_)
                coll44_ = _dafny.Map()
                def iife57_():
                    coll45_ = _dafny.Map()
                    compr_45_: _dafny.Map
                    for compr_45_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_514_cf7_: (self).f13})])).Elements:
                        d_515_v2_: _dafny.Map = compr_45_
                        if (d_515_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_514_cf7_: (self).f13})])):
                            coll45_[d_515_v2_] = False
                    return _dafny.Map(coll45_)
                compr_44_: _dafny.Map
                for compr_44_ in (iife57_()
                ).keys.Elements:
                    d_516_v1_: _dafny.Map = compr_44_
                    if (d_516_v1_) in (iife58_()
                    ):
                        coll44_[d_516_v1_] = self.f11
                return _dafny.Map(coll44_)
            return iife56_()
            
        elif True:
            d_517___mcc_h8_ = source9_.cf0
            d_518_cf0_ = d_517___mcc_h8_
            def iife59_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in _dafny.IntegerRange(809, 501):
                    d_519_v3_: int = compr_47_
                    if ((809) <= (d_519_v3_)) and ((d_519_v3_) < (501)):
                        coll47_[default__.safeDivisionInt(d_519_v3_, (self).f26)] = (self).f13
                return _dafny.Map(coll47_)
            return (_dafny.Map({_dafny.Map({(self).f12: (self).f12}): self.f11})) | (_dafny.Map({iife59_()
            : self.f11}))

    def fm33(self, p0, globalState):
        source10_ = D3_DC10(D3_DC10(D3_DC9((self).f13, _dafny.Map({(self).f26: (D4_DC11(_dafny.Map({(self).f12: (self).f26}))).cf20}), (self).f26, len(_dafny.Map({self.f11: self.f11})))))
        if source10_.is_DC9:
            d_520___mcc_h0_ = source10_.cf15
            d_521___mcc_h1_ = source10_.cf16
            d_522___mcc_h2_ = source10_.cf17
            d_523___mcc_h3_ = source10_.cf18
            d_524_cf18_ = d_523___mcc_h3_
            d_525_cf17_ = d_522___mcc_h2_
            d_526_cf16_ = d_521___mcc_h1_
            d_527_cf15_ = d_520___mcc_h0_
            return self.f11
        elif source10_.is_DC8:
            d_528___mcc_h4_ = source10_.cf14
            d_529_cf14_ = d_528___mcc_h4_
            return self.f11
        elif True:
            d_530___mcc_h5_ = source10_.cf19
            d_531_cf19_ = d_530___mcc_h5_
            return not(True)

    def m11(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        if self.f11:
            d_532_v0_: C1
            nw72_ = C1()
            nw72_.ctor__(p1)
            d_532_v0_ = nw72_
            (globalState).f8 = self.f11
            d_533_v1_: _dafny.Array
            nw73_ = _dafny.Array(False, 16)
            d_533_v1_ = nw73_
            d_534_v2_: _dafny.Map
            d_534_v2_ = _dafny.Map({(self.f11 if self.f11 else self.f11): d_533_v1_})
            d_534_v2_ = (d_534_v2_).set(self.f11, d_533_v1_)
            d_535_v3_: _dafny.Set
            d_535_v3_ = _dafny.Set({self.f11})
            d_536_v4_: _dafny.Set
            d_536_v4_ = _dafny.Set({(self).f12})
            d_537_v5_: _dafny.Seq
            d_537_v5_ = _dafny.SeqWithoutIsStrInference([len(d_536_v4_), p1])
            d_538_v6_: _dafny.MultiSet
            d_538_v6_ = _dafny.MultiSet([p1])
            d_535_v3_ = ((d_535_v3_) - (_dafny.Set({self.f11})) if (self).fm6(self.f11, (d_537_v5_)[default__.safeIndex(324, len(d_537_v5_))], (d_538_v6_).cardinality, globalState) else _dafny.Set({self.f11}))
            (self).f11 = (self).fm33(((self).f12) - (p0), globalState)
        elif True:
            d_539_v7_: str
            d_539_v7_ = _dafny.CodePoint('s')
            d_540_v8_: _dafny.Array
            def lambda37_(d_541_i0_):
                return default__.safeDivisionInt(d_541_i0_, (self).f13)

            init18_ = lambda37_
            nw74_ = _dafny.Array(None, 13)
            for i0_18_ in range(nw74_.length(0)):
                nw74_[i0_18_] = init18_(i0_18_)
            d_540_v8_ = nw74_
            rhs45_ = (((self).f12) >= (p2)) == (self.f11)
            rhs46_ = (_dafny.CodePoint('i') if not(self.f11) else d_539_v7_)
            rhs47_ = p3
            lhs32_ = self
            lhs32_.f11 = rhs45_
            d_539_v7_ = rhs46_
            d_540_v8_ = rhs47_
            d_542_v9_: _dafny.Map
            d_542_v9_ = _dafny.Map({p0: _dafny.Map({self.f11: p1})})
            rhs48_ = default__.safeModuloInt(p2, (0) - (p0))
            rhs49_ = (d_542_v9_) | (d_542_v9_)
            r1 = rhs48_
            d_542_v9_ = rhs49_
            (self).f11 = True
            d_543_v10_: _dafny.Map
            d_543_v10_ = _dafny.Map({d_539_v7_: self.f11})
            d_544_v11_: _dafny.Set
            d_544_v11_ = _dafny.Set({(self).f12})
            (self).f11 = ((self).fm8((self).f13, d_543_v10_, d_544_v11_, globalState)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwb")))
            d_545_v12_: _dafny.Seq
            d_545_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
            d_545_v12_ = ((self).f20) + ((self).f20)
        d_546_v13_: _dafny.Array
        nw75_ = _dafny.Array(None, 5)
        d_546_v13_ = nw75_
        d_547_v14_: _dafny.MultiSet
        d_547_v14_ = _dafny.MultiSet([d_546_v13_, d_546_v13_])
        d_547_v14_ = d_547_v14_
        d_548_v15_: D4
        d_548_v15_ = D4_DC12(self.f11, p1, p0)
        d_549_v16_: _dafny.Map
        d_549_v16_ = _dafny.Map({self.f11: self.f11})
        pat_let_tv13_ = p1
        d_550_v17_: _dafny.Map
        def iife60_(_pat_let6_0):
            def iife61_(d_551_dt__update__tmp_h0_):
                def iife62_(_pat_let7_0):
                    def iife63_(d_553_dt__update_hcf22_h0_):
                        def iife64_(_pat_let8_0):
                            def iife65_(d_554_dt__update_hcf21_h0_):
                                return D4_DC12(d_554_dt__update_hcf21_h0_, d_553_dt__update_hcf22_h0_, (d_551_dt__update__tmp_h0_).cf23)
                            return iife65_(_pat_let8_0)
                        return iife64_(self.f11)
                    return iife63_(_pat_let7_0)
                return iife62_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv13_ for d_552_i1_ in range(default__.abs(-544))])))
            return iife61_(_pat_let6_0)
        d_550_v17_ = _dafny.Map({(iife60_(d_548_v15_)).cf21: len((d_549_v16_ if self.f11 else d_549_v16_))})
        d_555_v18_: _dafny.Seq
        d_555_v18_ = _dafny.SeqWithoutIsStrInference([d_550_v17_, d_550_v17_])
        d_556_v19_: D0
        d_556_v19_ = D0_DC1(p1, (self).f20, (self).f26, self.f11)
        d_550_v17_ = ((d_555_v18_)[default__.safeIndex((self).f26, len(d_555_v18_))] if self.f11 else (_dafny.Map({self.f11: len((d_556_v19_).cf2)})) | (d_550_v17_))
        d_557_v20_: _dafny.Array
        nw76_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_557_v20_ = nw76_
        d_558_v21_: _dafny.Map
        d_558_v21_ = _dafny.Map({not(self.f11): d_557_v20_})
        d_559_v22_: _dafny.Map
        d_559_v22_ = _dafny.Map({((d_558_v21_)[self.f11] if (self.f11) in (d_558_v21_) else d_557_v20_): self.f11})
        if not((self.f11) == (((d_559_v22_)[d_557_v20_] if (d_557_v20_) in (d_559_v22_) else self.f11))):
            index61_ = default__.safeIndex(204, (p3).length(0))
            (p3)[index61_] = (0) - ((self).f26)
            index62_ = default__.safeIndex(204, (p3).length(0))
            (p3)[index62_] = p0
            if (self.f11) and (self.f11):
                r1 = default__.safeModuloInt((self).f13, (self).f26)
                r1 = (self).f12
                (globalState).f8 = not(self.f11)
                d_560_v23_: _dafny.Array
                def lambda38_(d_561_i2_):
                    return self.f11

                init19_ = lambda38_
                nw77_ = _dafny.Array(None, 22)
                for i0_19_ in range(nw77_.length(0)):
                    nw77_[i0_19_] = init19_(i0_19_)
                d_560_v23_ = nw77_
                index63_ = default__.safeIndex(638, (d_560_v23_).length(0))
                (d_560_v23_)[index63_] = self.f11
                d_562_v24_: _dafny.Array
                def lambda39_(d_563_v15_):
                    def lambda40_(d_564_i3_):
                        return d_563_v15_

                    return lambda40_

                init20_ = lambda39_(d_548_v15_)
                nw78_ = _dafny.Array(None, 4)
                for i0_20_ in range(nw78_.length(0)):
                    nw78_[i0_20_] = init20_(i0_20_)
                d_562_v24_ = nw78_
                index64_ = default__.safeIndex(963, (d_562_v24_).length(0))
                (d_562_v24_)[index64_] = d_548_v15_
                d_565_v25_: _dafny.Map
                d_565_v25_ = _dafny.Map({330: self.f11})
                index65_ = default__.safeIndex(638, (d_560_v23_).length(0))
                index66_ = default__.safeIndex(204, (p3).length(0))
                index67_ = default__.safeIndex(963, (d_562_v24_).length(0))
                rhs50_ = self.f11
                rhs51_ = (len((_dafny.Map({(p3)[default__.safeIndex(204, (p3).length(0))]: self.f11})) | (d_565_v25_))) * (p2)
                rhs52_ = (((self).f12) < (767)) == (not(self.f11))
                rhs53_ = (d_548_v15_ if True else D4_DC12(self.f11, len((self).f20), (0) - (p2)))
                rhs54_ = (d_560_v23_ if self.f11 else d_560_v23_)
                lhs33_ = d_560_v23_
                lhs34_ = default__.safeIndex(638, (d_560_v23_).length(0))
                lhs35_ = p3
                lhs36_ = default__.safeIndex(204, (p3).length(0))
                lhs37_ = globalState
                lhs38_ = d_562_v24_
                lhs39_ = default__.safeIndex(963, (d_562_v24_).length(0))
                lhs33_[lhs34_] = rhs50_
                lhs35_[lhs36_] = rhs51_
                lhs37_.f8 = rhs52_
                lhs38_[lhs39_] = rhs53_
                d_560_v23_ = rhs54_
                d_566_v26_: _dafny.Seq
                d_566_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_566_v26_ = d_566_v26_
            elif True:
                r1 = 758
                index68_ = default__.safeIndex(204, (p3).length(0))
                (p3)[index68_] = (0) - (default__.safeDivisionInt((self).f13, ((self).f13) * (p2)))
                (self).f11 = self.f11
                d_567_v27_: _dafny.Array
                nw79_ = _dafny.Array(int(0), 4)
                d_567_v27_ = nw79_
                nw80_ = _dafny.Array(None, 10)
                nw80_[int(0)] = (p3)[default__.safeIndex(204, (p3).length(0))]
                nw80_[int(1)] = (p2) * (p2)
                nw80_[int(2)] = (p3)[default__.safeIndex(204, (p3).length(0))]
                nw80_[int(3)] = (0) - ((0) - ((self).f12))
                nw80_[int(4)] = 505
                nw80_[int(5)] = p1
                nw80_[int(6)] = p2
                nw80_[int(7)] = 344
                nw80_[int(8)] = (self).f26
                nw80_[int(9)] = len(_dafny.Set({self.f11}))
                d_567_v27_ = nw80_
                d_568_v28_: _dafny.Array
                nw81_ = _dafny.Array(False, 10)
                d_568_v28_ = nw81_
                r0 = d_568_v28_
            if self.f11:
                d_549_v16_ = (d_549_v16_).set(((self).f20) != ((self).f20), self.f11)
                r1 = ((-753) - (default__.fm2(globalState))) - (p0)
                d_569_v29_: _dafny.MultiSet
                d_569_v29_ = _dafny.MultiSet([p2])
                d_570_v30_: _dafny.MultiSet
                d_570_v30_ = _dafny.MultiSet([(d_569_v29_).isdisjoint(d_569_v29_), self.f11, self.f11, self.f11, not((self).fm18(self.f11, self.f11, (self).f26, p1, globalState))])
                d_570_v30_ = d_570_v30_
                d_571_v31_: _dafny.Seq
                d_571_v31_ = _dafny.SeqWithoutIsStrInference([len((d_549_v16_) | (d_549_v16_))])
                (globalState).f5 = d_571_v31_
                d_572_v32_: C0
                nw82_ = C0()
                nw82_.ctor__(default__.fm2(globalState), self.f11, (p3)[default__.safeIndex(204, (p3).length(0))], p2)
                d_572_v32_ = nw82_
            elif True:
                r1 = (self).f12
                d_573_v33_: _dafny.Array
                def lambda41_(d_574_i4_):
                    return (self).f20

                init21_ = lambda41_
                nw83_ = _dafny.Array(None, 14)
                for i0_21_ in range(nw83_.length(0)):
                    nw83_[i0_21_] = init21_(i0_21_)
                d_573_v33_ = nw83_
                d_575_v34_: _dafny.Map
                d_575_v34_ = _dafny.Map({(self).f20: d_573_v33_})
                d_575_v34_ = (d_575_v34_).set((self).f20, (D5_DC14(d_573_v33_)).cf25)
                d_576_v35_: str
                d_576_v35_ = _dafny.CodePoint('s')
                d_577_v36_: _dafny.Map
                d_577_v36_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_578_i5_ in range(default__.abs(873))])).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_579_i5_ in range(default__.abs(873))]))), d_576_v35_): (self).f20})
                d_577_v36_ = d_577_v36_
                d_580_v37_: D3
                d_580_v37_ = D3_DC8(d_550_v17_)
                pat_let_tv14_ = d_550_v17_
                d_581_v38_: _dafny.Map
                def iife66_(_pat_let9_0):
                    def iife67_(d_582_dt__update__tmp_h1_):
                        def iife68_(_pat_let10_0):
                            def iife69_(d_583_dt__update_hcf14_h0_):
                                return D3_DC8(d_583_dt__update_hcf14_h0_)
                            return iife69_(_pat_let10_0)
                        return iife68_(pat_let_tv14_)
                    return iife67_(_pat_let9_0)
                d_581_v38_ = _dafny.Map({self.f11: iife66_(d_580_v37_)})
                d_584_v39_: _dafny.MultiSet
                d_584_v39_ = _dafny.MultiSet([d_581_v38_, d_581_v38_])
                d_585_v40_: _dafny.Map
                d_585_v40_ = _dafny.Map({d_584_v39_: d_576_v35_})
                d_585_v40_ = (d_585_v40_).set(d_584_v39_, d_576_v35_)
                d_586_v41_: _dafny.Seq
                d_586_v41_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                d_587_v42_: _dafny.MultiSet
                d_587_v42_ = _dafny.MultiSet([(d_586_v41_).set(default__.safeIndex((p3)[default__.safeIndex(204, (p3).length(0))], len(d_586_v41_)), p0), d_586_v41_])
                d_588_v43_: _dafny.MultiSet
                d_588_v43_ = _dafny.MultiSet([(d_587_v42_).cardinality, p2, -503, 235])
                d_549_v16_ = (d_549_v16_).set(True, (d_588_v43_).ispropersubset(d_588_v43_))
            d_589_v44_: _dafny.Seq
            d_589_v44_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11, not(self.f11)])
            d_590_v45_: C1
            nw84_ = C1()
            nw84_.ctor__(len(d_589_v44_))
            d_590_v45_ = nw84_
            d_591_v46_: _dafny.Array
            nw85_ = _dafny.Array(D0.default()(), 16)
            d_591_v46_ = nw85_
            d_592_v47_: _dafny.Array
            def lambda42_(d_593_i6_):
                return _dafny.SeqWithoutIsStrInference([652 for d_594_i7_ in range(default__.abs(-821))])

            init22_ = lambda42_
            nw86_ = _dafny.Array(None, 4)
            for i0_22_ in range(nw86_.length(0)):
                nw86_[i0_22_] = init22_(i0_22_)
            d_592_v47_ = nw86_
            d_595_v48_: D0
            d_595_v48_ = D0_DC0(d_592_v47_)
            index69_ = default__.safeIndex(990, (d_591_v46_).length(0))
            (d_591_v46_)[index69_] = d_595_v48_
            index70_ = default__.safeIndex(990, (d_591_v46_).length(0))
            (d_591_v46_)[index70_] = D0_DC0((d_595_v48_).cf0)
        elif True:
            r1 = (default__.fm2(globalState)) - (p2)
            (self).f11 = not(((self).f12) > (len(((self).f20) + ((self).f20))))
            (globalState).f8 = (self).fm7(globalState)
            d_596_v49_: _dafny.MultiSet
            d_596_v49_ = _dafny.MultiSet([(self).f26, 583, p2, (self).f12])
            d_597_v50_: _dafny.Map
            d_597_v50_ = _dafny.Map({(d_596_v49_).intersection(d_596_v49_): self.f11})
            d_598_v51_: _dafny.Seq
            d_598_v51_ = _dafny.SeqWithoutIsStrInference([self.f11, not(self.f11)])
            d_599_v52_: _dafny.Map
            d_599_v52_ = _dafny.Map({d_598_v51_: -494})
            d_600_v53_: _dafny.Seq
            d_600_v53_ = _dafny.SeqWithoutIsStrInference([849, (self).f26, ((d_599_v52_)[_dafny.SeqWithoutIsStrInference([self.f11, self.f11])] if (_dafny.SeqWithoutIsStrInference([self.f11, self.f11])) in (d_599_v52_) else (self).f12), (self).f12])
            d_601_v54_: _dafny.Seq
            d_601_v54_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_602_i8_ in range(default__.abs(350))])])
            d_597_v50_ = (d_597_v50_).set(_dafny.MultiSet((d_600_v53_) + (d_600_v53_)), (d_601_v54_) < (d_601_v54_))
            if self.f11:
                d_603_v55_: _dafny.Array
                nw87_ = _dafny.Array(None, 1)
                nw87_[int(0)] = (d_600_v53_)[default__.safeIndex((self).f12, len(d_600_v53_))]
                d_603_v55_ = nw87_
                d_604_v56_: _dafny.Seq
                d_604_v56_ = _dafny.SeqWithoutIsStrInference([p3, p3, p3, p3])
                d_603_v55_ = (d_604_v56_)[default__.safeIndex((p1) + (175), len(d_604_v56_))]
                d_603_v55_ = d_603_v55_
                (self).f11 = self.f11
                d_605_v57_: _dafny.MultiSet
                d_605_v57_ = _dafny.MultiSet([True])
                d_550_v17_ = (d_550_v17_).set((d_605_v57_).issubset(d_605_v57_), p2)
                d_606_v58_: _dafny.Map
                d_606_v58_ = _dafny.Map({(self).f13: p2})
                d_607_v59_: C1
                nw88_ = C1()
                nw88_.ctor__(((self).f12) * (len(d_606_v58_)))
                d_607_v59_ = nw88_
            elif True:
                d_608_v60_: _dafny.Map
                d_608_v60_ = _dafny.Map({default__.safeDivisionInt((self).f12, (0) - ((self).f13)): default__.fm2(globalState)})
                d_609_v61_: str
                d_609_v61_ = _dafny.CodePoint('f')
                d_608_v60_ = (d_608_v60_).set(((self).f13) * (p0), len(((self).f20).set(default__.safeIndex(2, len((self).f20)), d_609_v61_)))
                (globalState).f5 = d_600_v53_
                index71_ = default__.safeIndex(439, (p3).length(0))
                (p3)[index71_] = (self).f26
                index72_ = default__.safeIndex(439, (p3).length(0))
                (p3)[index72_] = (0) - ((self).f13)
                d_610_v62_: _dafny.Set
                d_610_v62_ = _dafny.Set({self.f11})
                d_611_v63_: _dafny.Array
                nw89_ = _dafny.Array(None, 17)
                nw89_[int(0)] = self.f11
                nw89_[int(1)] = self.f11
                nw89_[int(2)] = self.f11
                nw89_[int(3)] = self.f11
                nw89_[int(4)] = self.f11
                nw89_[int(5)] = self.f11
                nw89_[int(6)] = self.f11
                nw89_[int(7)] = self.f11
                nw89_[int(8)] = (self).fm6(self.f11, len((d_549_v16_).set(self.f11, self.f11)), 964, globalState)
                nw89_[int(9)] = not(self.f11)
                nw89_[int(10)] = self.f11
                nw89_[int(11)] = self.f11
                nw89_[int(12)] = self.f11
                nw89_[int(13)] = self.f11
                nw89_[int(14)] = self.f11
                nw89_[int(15)] = self.f11
                nw89_[int(16)] = self.f11
                d_611_v63_ = nw89_
                d_612_v64_: _dafny.Seq
                d_612_v64_ = _dafny.SeqWithoutIsStrInference([(self).f12, len(d_610_v62_)])
                d_613_v65_: _dafny.Map
                d_613_v65_ = _dafny.Map({self.f11: d_612_v64_})
                d_614_v66_: _dafny.Map
                d_614_v66_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_611_v63_, d_611_v63_]): len((d_613_v65_) | (d_613_v65_))})
                d_615_v67_: _dafny.Seq
                d_615_v67_ = _dafny.SeqWithoutIsStrInference([d_611_v63_])
                d_616_v68_: _dafny.Set
                d_616_v68_ = _dafny.Set({-78})
                index73_ = default__.safeIndex(439, (p3).length(0))
                index74_ = default__.safeIndex(439, (p3).length(0))
                rhs55_ = not(((0) - ((p3)[default__.safeIndex(439, (p3).length(0))])) > (p0))
                rhs56_ = (p3)[default__.safeIndex(439, (p3).length(0))]
                rhs57_ = ((d_614_v66_)[d_615_v67_] if (d_615_v67_) in (d_614_v66_) else len(d_598_v51_))
                rhs58_ = d_610_v62_
                rhs59_ = not(((d_610_v62_).ispropersubset(d_610_v62_)) and ((d_616_v68_).issubset(d_616_v68_)))
                lhs40_ = globalState
                lhs41_ = p3
                lhs42_ = default__.safeIndex(439, (p3).length(0))
                lhs43_ = p3
                lhs44_ = default__.safeIndex(439, (p3).length(0))
                lhs45_ = globalState
                lhs40_.f8 = rhs55_
                lhs41_[lhs42_] = rhs56_
                lhs43_[lhs44_] = rhs57_
                d_610_v62_ = rhs58_
                lhs45_.f8 = rhs59_
                d_617_v69_: _dafny.Map
                d_617_v69_ = _dafny.Map({(p1 if self.f11 else p2): d_598_v51_})
                d_617_v69_ = _dafny.Map({p0: d_598_v51_})
        d_618_v70_: D3
        d_618_v70_ = D3_DC8(d_550_v17_)
        d_619_v71_: D3
        d_619_v71_ = D3_DC10(d_618_v70_)
        d_620_v72_: D3
        d_620_v72_ = D3_DC10(D3_DC10(d_618_v70_))
        d_621_v73_: D3
        d_621_v73_ = D3_DC10(d_619_v71_)
        d_622_v74_: D3
        d_622_v74_ = D3_DC10(d_620_v72_)
        d_623_v75_: D3
        d_623_v75_ = D3_DC10(d_620_v72_)
        pat_let_tv15_ = p2
        def lambda43_(source11_):
            if source11_.is_DC9:
                d_624___mcc_h0_ = source11_.cf15
                d_625___mcc_h1_ = source11_.cf16
                d_626___mcc_h2_ = source11_.cf17
                d_627___mcc_h3_ = source11_.cf18
                d_628_cf18_ = d_627___mcc_h3_
                d_629_cf17_ = d_626___mcc_h2_
                d_630_cf16_ = d_625___mcc_h1_
                d_631_cf15_ = d_624___mcc_h0_
                return default__.safeModuloInt(d_628_cf18_, (_dafny.MultiSet([self.f11, self.f11])).cardinality)
            elif source11_.is_DC8:
                d_632___mcc_h4_ = source11_.cf14
                d_633_cf14_ = d_632___mcc_h4_
                return (self).f13
            elif True:
                d_634___mcc_h5_ = source11_.cf19
                d_635_cf19_ = d_634___mcc_h5_
                return pat_let_tv15_

        r1 = lambda43_(d_623_v75_)
        d_636_v76_: _dafny.Array
        d_637_v77_: bool
        out2_: _dafny.Array
        out3_: bool
        out2_, out3_ = default__.m0(globalState)
        d_636_v76_ = out2_
        d_637_v77_ = out3_
        d_638_v78_: _dafny.Array
        def lambda44_(d_639_v77_):
            def lambda45_(d_640_i9_):
                return not (d_639_v77_) or (self.f11)

            return lambda45_

        init23_ = lambda44_(d_637_v77_)
        nw90_ = _dafny.Array(None, 26)
        for i0_23_ in range(nw90_.length(0)):
            nw90_[i0_23_] = init23_(i0_23_)
        d_638_v78_ = nw90_
        r0 = d_638_v78_
        r1 = p1
        r2 = _dafny.Set({p1, (0) - ((self).f12), (self).f12})
        return r0, r1, r2

    def m23(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D3 = D3.default()()
        r2: int = int(0)
        r3: bool = False
        hi1_ = (self).f13
        for d_641_i0_ in range((self).f12, hi1_):
            r0 = (self).f20
            r2 = default__.safeDivisionInt((self).f13, p2)
            d_642_v0_: _dafny.Array
            def lambda46_(d_643_p0_):
                def lambda47_(d_644_i1_):
                    return d_643_p0_

                return lambda47_

            init24_ = lambda46_(p0)
            nw91_ = _dafny.Array(None, 26)
            for i0_24_ in range(nw91_.length(0)):
                nw91_[i0_24_] = init24_(i0_24_)
            d_642_v0_ = nw91_
            nw92_ = _dafny.Array(False, 4)
            d_642_v0_ = nw92_
            (self).f11 = self.f11
        d_645_v1_: _dafny.Map
        d_645_v1_ = _dafny.Map({not(p0): not(self.f11)})
        d_646_i2_: int
        d_646_i2_ = 0
        with _dafny.label("2"):
            while ((d_645_v1_) | (d_645_v1_)) == ((d_645_v1_) | (d_645_v1_)):
                with _dafny.c_label("2"):
                    if (d_646_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_646_i2_ = (d_646_i2_) + (1)
                    d_647_v2_: str
                    d_647_v2_ = _dafny.CodePoint('s')
                    d_648_v3_: D0
                    d_648_v3_ = D0_DC1((self).f13, (self).f20, p1, p0)
                    source12_ = default__.fm37((self).fm6(p0, len(default__.fm36((self).f12, p2, d_647_v2_, globalState)), p2, globalState), (d_648_v3_).cf1, (self).fm33(p2, globalState), p2, globalState)
                    if source12_.is_DC15:
                        d_649___mcc_h0_ = source12_.cf26
                        d_650___mcc_h1_ = source12_.cf27
                        d_651___mcc_h2_ = source12_.cf28
                        d_652_cf28_ = d_651___mcc_h2_
                        d_653_cf27_ = d_650___mcc_h1_
                        d_654_cf26_ = d_649___mcc_h0_
                        d_653_cf27_ = (self).f13
                        d_655_v4_: _dafny.Seq
                        d_655_v4_ = _dafny.SeqWithoutIsStrInference([d_652_cf28_])
                        d_656_v5_: _dafny.Array
                        nw93_ = _dafny.Array(None, 17)
                        nw93_[int(0)] = d_653_cf27_
                        nw93_[int(1)] = 61
                        nw93_[int(2)] = d_654_cf26_
                        nw93_[int(3)] = d_653_cf27_
                        nw93_[int(4)] = (self).f26
                        nw93_[int(5)] = (d_655_v4_)[default__.safeIndex(len(d_645_v1_), len(d_655_v4_))]
                        nw93_[int(6)] = len((self).f20)
                        nw93_[int(7)] = 975
                        nw93_[int(8)] = (self).f12
                        nw93_[int(9)] = (self).f26
                        nw93_[int(10)] = (self).f26
                        nw93_[int(11)] = p1
                        nw93_[int(12)] = d_653_cf27_
                        nw93_[int(13)] = default__.fm2(globalState)
                        nw93_[int(14)] = d_654_cf26_
                        nw93_[int(15)] = (self).f13
                        nw93_[int(16)] = d_652_cf28_
                        d_656_v5_ = nw93_
                        d_657_v6_: _dafny.MultiSet
                        d_657_v6_ = _dafny.MultiSet([d_656_v5_, d_656_v5_, d_656_v5_, d_656_v5_, d_656_v5_])
                        d_658_v7_: _dafny.Seq
                        d_658_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                        d_659_v8_: _dafny.MultiSet
                        d_659_v8_ = _dafny.MultiSet([len(d_658_v7_), (self).f26])
                        d_660_v9_: _dafny.MultiSet
                        d_660_v9_ = _dafny.MultiSet([d_647_v2_, d_647_v2_, d_647_v2_, d_647_v2_, d_647_v2_])
                        d_661_v10_: _dafny.Map
                        d_661_v10_ = _dafny.Map({p0: p2})
                        d_662_v11_: _dafny.Map
                        d_662_v11_ = _dafny.Map({d_648_v3_: (self).f13})
                        d_663_v12_: _dafny.Array
                        nw94_ = _dafny.Array(None, 19)
                        nw94_[int(0)] = (self).f13
                        nw94_[int(1)] = d_652_cf28_
                        nw94_[int(2)] = (((d_657_v6_).set(d_656_v5_, default__.abs((self).f26))) | (d_657_v6_)).cardinality
                        nw94_[int(3)] = ((self).f26) * (len((self).f20))
                        nw94_[int(4)] = d_653_cf27_
                        nw94_[int(5)] = p2
                        nw94_[int(6)] = p2
                        nw94_[int(7)] = (self).f26
                        nw94_[int(8)] = d_654_cf26_
                        nw94_[int(9)] = default__.safeModuloInt(p1, d_654_cf26_)
                        nw94_[int(10)] = default__.safeDivisionInt((0) - (p2), default__.fm2(globalState))
                        nw94_[int(11)] = (0) - (((self).f13 if self.f11 else (0) - (d_653_cf27_)))
                        nw94_[int(12)] = ((d_659_v8_) - (d_659_v8_)).cardinality
                        nw94_[int(13)] = p1
                        nw94_[int(14)] = ((d_660_v9_)[d_647_v2_] if (d_647_v2_) in (d_660_v9_) else (self).f13)
                        nw94_[int(15)] = ((d_661_v10_)[self.f11] if (self.f11) in (d_661_v10_) else (0) - (p1))
                        nw94_[int(16)] = (len(d_658_v7_)) * ((self).f13)
                        nw94_[int(17)] = default__.safeDivisionInt(p1, ((d_662_v11_)[d_648_v3_] if (d_648_v3_) in (d_662_v11_) else d_653_cf27_))
                        nw94_[int(18)] = (self).f26
                        d_663_v12_ = nw94_
                        nw95_ = _dafny.Array(int(0), 27)
                        d_663_v12_ = nw95_
                        d_664_v13_: _dafny.Set
                        d_664_v13_ = _dafny.Set({(self).f13, (self).f13})
                        d_665_v14_: _dafny.Map
                        d_665_v14_ = _dafny.Map({((self).f13) * (d_653_cf27_): (d_664_v13_) - (d_664_v13_)})
                        d_665_v14_ = d_665_v14_
                        d_663_v12_ = d_656_v5_
                    elif True:
                        d_666___mcc_h3_ = source12_.cf25
                        d_667_cf25_ = d_666___mcc_h3_
                        d_668_v15_: _dafny.Array
                        nw96_ = _dafny.Array(None, 4)
                        nw96_[int(0)] = (self).f12
                        nw96_[int(1)] = len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f26])) for d_669_i3_ in range(default__.abs(726))]))
                        nw96_[int(2)] = (0) - ((self).f13)
                        nw96_[int(3)] = (self).f13
                        d_668_v15_ = nw96_
                        d_670_v16_: _dafny.Map
                        d_670_v16_ = _dafny.Map({d_668_v15_: (_dafny.MultiSet([len(_dafny.Set({p0}))])).cardinality})
                        d_671_v17_: _dafny.MultiSet
                        d_671_v17_ = _dafny.MultiSet([((d_670_v16_)[d_668_v15_] if (d_668_v15_) in (d_670_v16_) else (self).f12)])
                        d_672_v18_: _dafny.MultiSet
                        d_672_v18_ = _dafny.MultiSet([((d_671_v17_)[p2] if (p2) in (d_671_v17_) else p2)])
                        d_673_v19_: C0
                        nw97_ = C0()
                        nw97_.ctor__((self).f12, p0, (self).f13, ((self).f12) - (((d_671_v17_)[(0) - (-963)] if ((0) - (-963)) in (d_671_v17_) else p1)))
                        d_673_v19_ = nw97_
                        d_673_v19_ = (d_673_v19_ if self.f11 else (d_673_v19_ if not(d_673_v19_.f29) else d_673_v19_))
                        d_647_v2_ = d_647_v2_
                        d_674_v20_: _dafny.MultiSet
                        d_674_v20_ = _dafny.MultiSet([self.f11])
                        d_674_v20_ = (d_674_v20_) | (d_674_v20_)
                        d_647_v2_ = d_647_v2_
                    d_675_v21_: _dafny.Array
                    def lambda48_(d_676_v2_):
                        def lambda49_(d_677_i4_):
                            return _dafny.Map({d_676_v2_: self.f11})

                        return lambda49_

                    init25_ = lambda48_(d_647_v2_)
                    nw98_ = _dafny.Array(None, 19)
                    for i0_25_ in range(nw98_.length(0)):
                        nw98_[i0_25_] = init25_(i0_25_)
                    d_675_v21_ = nw98_
                    d_678_v22_: _dafny.Map
                    d_678_v22_ = _dafny.Map({d_647_v2_: self.f11})
                    index75_ = default__.safeIndex(157, (d_675_v21_).length(0))
                    (d_675_v21_)[index75_] = (d_678_v22_) | ((_dafny.Map({d_647_v2_: self.f11})).set(d_647_v2_, p0))
                    d_679_v23_: _dafny.Map
                    d_679_v23_ = _dafny.Map({p0: (self).f12})
                    d_680_v24_: _dafny.Array
                    nw99_ = _dafny.Array(None, 20)
                    nw99_[int(0)] = (self).f13
                    nw99_[int(1)] = p2
                    nw99_[int(2)] = len(d_679_v23_)
                    nw99_[int(3)] = p2
                    nw99_[int(4)] = (self).f13
                    nw99_[int(5)] = (self).f12
                    nw99_[int(6)] = (self).f12
                    nw99_[int(7)] = (self).f12
                    nw99_[int(8)] = p2
                    nw99_[int(9)] = p1
                    nw99_[int(10)] = (self).f13
                    nw99_[int(11)] = p1
                    nw99_[int(12)] = (self).f12
                    nw99_[int(13)] = (self).f12
                    nw99_[int(14)] = p2
                    nw99_[int(15)] = (self).f13
                    nw99_[int(16)] = p2
                    nw99_[int(17)] = (self).f26
                    nw99_[int(18)] = (self).f13
                    nw99_[int(19)] = p2
                    d_680_v24_ = nw99_
                    d_681_v25_: _dafny.Map
                    d_681_v25_ = _dafny.Map({p1: d_680_v24_})
                    d_682_v26_: _dafny.Map
                    d_682_v26_ = _dafny.Map({((d_681_v25_)[p2] if (p2) in (d_681_v25_) else d_680_v24_): ((self).f12 if p0 else (self).f12)})
                    d_683_v27_: _dafny.Map
                    d_683_v27_ = _dafny.Map({928: (self).f12})
                    index76_ = default__.safeIndex(157, (d_675_v21_).length(0))
                    rhs60_ = default__.safeModuloInt((self).f12, len((default__.fm38(d_647_v2_, d_647_v2_, globalState) if self.f11 else d_683_v27_)))
                    rhs61_ = d_678_v22_
                    rhs62_ = (d_682_v26_) | (d_682_v26_)
                    lhs46_ = d_675_v21_
                    lhs47_ = default__.safeIndex(157, (d_675_v21_).length(0))
                    r2 = rhs60_
                    lhs46_[lhs47_] = rhs61_
                    d_682_v26_ = rhs62_
                    r2 = -497
                    d_684_v28_: C0
                    nw100_ = C0()
                    nw100_.ctor__(p2, self.f11, (self).f26, 712)
                    d_684_v28_ = nw100_
                    pass
            pass
        d_685_v29_: str
        d_685_v29_ = _dafny.CodePoint('f')
        d_686_v30_: _dafny.Map
        d_686_v30_ = _dafny.Map({p2: d_685_v29_})
        d_687_v32_: _dafny.MultiSet
        def iife70_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(225, 317):
                d_688_v31_: int = compr_48_
                if ((225) <= (d_688_v31_)) and ((d_688_v31_) < (317)):
                    coll48_[(d_688_v31_) + (len(_dafny.SeqWithoutIsStrInference([p0])))] = d_685_v29_
            return _dafny.Map(coll48_)
        d_687_v32_ = _dafny.MultiSet([iife70_()
        ])
        if not(not (False) or ((d_686_v30_) in (d_687_v32_))):
            d_689_v33_: _dafny.Map
            d_689_v33_ = _dafny.Map({(self).f20: p1})
            d_690_v34_: C0
            nw101_ = C0()
            nw101_.ctor__((p2) + (((d_689_v33_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfxotwq"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfxotwq"))) in (d_689_v33_) else (0) - ((self).f13))), self.f11, len((self).f20), (len((self).f20)) - ((self).f26))
            d_690_v34_ = nw101_
            d_691_v35_: _dafny.Map
            d_691_v35_ = _dafny.Map({(0) - ((self).f13): (self).f12})
            d_692_v36_: _dafny.Seq
            d_692_v36_ = _dafny.SeqWithoutIsStrInference([p0, (p2) > ((self).f13), (58) <= (((d_691_v35_)[p1] if (p1) in (d_691_v35_) else p1)), d_690_v34_.f29])
            if (d_692_v36_)[default__.safeIndex(p2, len(d_692_v36_))]:
                r2 = 91
                d_693_v37_: _dafny.Set
                d_693_v37_ = _dafny.Set({p0})
                d_694_v38_: _dafny.Map
                d_694_v38_ = _dafny.Map({_dafny.CodePoint('g'): d_693_v37_})
                d_695_v39_: C0
                nw102_ = C0()
                nw102_.ctor__((self).f26, d_690_v34_.f29, len(d_694_v38_), 839)
                d_695_v39_ = nw102_
                d_696_v40_: _dafny.Array
                d_697_v41_: bool
                out4_: _dafny.Array
                out5_: bool
                out4_, out5_ = default__.m0(globalState)
                d_696_v40_ = out4_
                d_697_v41_ = out5_
                d_698_v42_: _dafny.Map
                d_698_v42_ = _dafny.Map({p1: self.f11})
                d_699_v43_: _dafny.Seq
                d_699_v43_ = _dafny.SeqWithoutIsStrInference([d_698_v42_, d_698_v42_])
                d_700_v44_: _dafny.Array
                nw103_ = _dafny.Array(None, 16)
                nw103_[int(0)] = 983
                nw103_[int(1)] = p1
                nw103_[int(2)] = (self).f12
                nw103_[int(3)] = p1
                nw103_[int(4)] = (self).f12
                nw103_[int(5)] = len(d_692_v36_)
                nw103_[int(6)] = (d_690_v34_).f28
                nw103_[int(7)] = (d_690_v34_).f28
                nw103_[int(8)] = 717
                nw103_[int(9)] = len(d_699_v43_)
                nw103_[int(10)] = (d_695_v39_).f28
                nw103_[int(11)] = 242
                nw103_[int(12)] = ((self).f26) + ((d_695_v39_).f28)
                nw103_[int(13)] = (d_690_v34_).f28
                nw103_[int(14)] = (d_690_v34_).f28
                nw103_[int(15)] = 248
                d_700_v44_ = nw103_
                d_700_v44_ = d_700_v44_
                r2 = (self).f13
            elif True:
                d_701_v45_: _dafny.Array
                nw104_ = _dafny.Array(int(0), 26)
                d_701_v45_ = nw104_
                index77_ = default__.safeIndex(819, (d_701_v45_).length(0))
                (d_701_v45_)[index77_] = (self).f26
                index78_ = default__.safeIndex(819, (d_701_v45_).length(0))
                (d_701_v45_)[index78_] = (self).f12
                (globalState).f8 = self.f11
                index79_ = default__.safeIndex(819, (d_701_v45_).length(0))
                (d_701_v45_)[index79_] = (d_690_v34_).f28
                (d_690_v34_).f29 = not(p0)
                r2 = ((self).f12) + ((d_690_v34_).f28)
            d_702_v46_: _dafny.Array
            def lambda50_(d_703_v34_):
                def lambda51_(d_704_i5_):
                    return d_703_v34_.f29

                return lambda51_

            init26_ = lambda50_(d_690_v34_)
            nw105_ = _dafny.Array(None, 11)
            for i0_26_ in range(nw105_.length(0)):
                nw105_[i0_26_] = init26_(i0_26_)
            d_702_v46_ = nw105_
            d_705_v47_: _dafny.Map
            d_705_v47_ = _dafny.Map({(d_685_v29_ if d_690_v34_.f29 else d_685_v29_): 798})
            d_706_v48_: _dafny.Seq
            d_706_v48_ = _dafny.SeqWithoutIsStrInference([p1])
            rhs63_ = d_702_v46_
            rhs64_ = ((self).f20) + ((self).f20)
            rhs65_ = ((d_705_v47_)[d_685_v29_] if (d_685_v29_) in (d_705_v47_) else len(d_706_v48_))
            d_702_v46_ = rhs63_
            r0 = rhs64_
            r2 = rhs65_
            if d_690_v34_.f29:
                d_707_v49_: C0
                nw106_ = C0()
                nw106_.ctor__((0) - (default__.safeDivisionInt((d_690_v34_).f28, (self).f13)), d_690_v34_.f29, (self).f26, (d_706_v48_)[default__.safeIndex((self).f26, len(d_706_v48_))])
                d_707_v49_ = nw106_
                (d_707_v49_).f29 = self.f11
                r2 = ((d_707_v49_).f28) * ((self).f12)
                d_708_v50_: _dafny.Map
                d_708_v50_ = _dafny.Map({self.f11: (self).f26})
                d_709_v51_: _dafny.Set
                d_709_v51_ = _dafny.Set({len(d_708_v50_), -61})
                d_709_v51_ = d_709_v51_
                r2 = (-438 if d_707_v49_.f29 else (self).f26)
            elif True:
                rhs66_ = len(_dafny.SeqWithoutIsStrInference([(d_690_v34_).f28, p2]))
                rhs67_ = (d_692_v36_) + ((d_692_v36_) + (default__.fm39(d_690_v34_.f29, False, globalState)))
                r2 = rhs66_
                d_692_v36_ = rhs67_
                d_710_v53_: _dafny.Array
                def lambda52_(d_711_v34_):
                    def lambda53_(d_712_i6_):
                        def iife71_():
                            coll49_ = _dafny.Map()
                            compr_49_: int
                            for compr_49_ in _dafny.IntegerRange(916, 751):
                                d_713_v52_: int = compr_49_
                                if ((916) <= (d_713_v52_)) and ((d_713_v52_) < (751)):
                                    coll49_[default__.safeModuloInt(d_713_v52_, 781)] = (self).f20
                            return _dafny.Map(coll49_)
                        return (D6_DC16(_dafny.Map({(d_711_v34_).f28: D5_DC15((d_711_v34_).f28, (d_711_v34_).f28, len(iife71_()
))}))).cf29

                    return lambda53_

                init27_ = lambda52_(d_690_v34_)
                nw107_ = _dafny.Array(None, 27)
                for i0_27_ in range(nw107_.length(0)):
                    nw107_[i0_27_] = init27_(i0_27_)
                d_710_v53_ = nw107_
                d_714_v54_: _dafny.MultiSet
                d_714_v54_ = _dafny.MultiSet([p1])
                d_715_v55_: D1
                d_715_v55_ = D1_DC5((d_714_v54_) | (d_714_v54_), (len(default__.fm36((d_690_v34_).f28, (d_690_v34_).f28, d_685_v29_, globalState))) >= ((0) - (p2)))
                d_716_v56_: D4
                d_716_v56_ = D4_DC13((self).f13)
                rhs68_ = d_685_v29_
                rhs69_ = default__.fm36(((d_690_v34_).f28) * ((d_716_v56_).cf24), 466, d_685_v29_, globalState)
                rhs70_ = d_710_v53_
                rhs71_ = D1_DC5(d_714_v54_, p0)
                rhs72_ = len((self).f20)
                d_685_v29_ = rhs68_
                r0 = rhs69_
                d_710_v53_ = rhs70_
                d_715_v55_ = rhs71_
                r2 = rhs72_
                d_717_v57_: _dafny.Map
                d_717_v57_ = _dafny.Map({not(d_690_v34_.f29): p2})
                d_717_v57_ = (d_717_v57_).set((self.f11) not in (d_717_v57_), p2)
                (self).f11 = p0
                d_691_v35_ = (d_691_v35_).set(((self).f13) + ((self).f26), (d_690_v34_).f28)
            d_718_v58_: _dafny.Set
            d_718_v58_ = _dafny.Set({False})
            rhs73_ = (self.f11 if (d_692_v36_)[default__.safeIndex(469, len(d_692_v36_))] else p0)
            rhs74_ = (self).f26
            rhs75_ = _dafny.Set({(d_692_v36_)[default__.safeIndex(p1, len(d_692_v36_))]})
            rhs76_ = d_685_v29_
            lhs48_ = globalState
            lhs48_.f8 = rhs73_
            r2 = rhs74_
            d_718_v58_ = rhs75_
            d_685_v29_ = rhs76_
        elif True:
            (self).f11 = p0
            if not((self).fm7(globalState)):
                d_719_v59_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 24)
                d_719_v59_ = nw108_
                index80_ = default__.safeIndex(991, (d_719_v59_).length(0))
                (d_719_v59_)[index80_] = (self).f26
                index81_ = default__.safeIndex(991, (d_719_v59_).length(0))
                (d_719_v59_)[index81_] = ((self).f12) - ((self).f12)
                d_720_v60_: _dafny.Array
                d_721_v61_: bool
                out6_: _dafny.Array
                out7_: bool
                out6_, out7_ = default__.m0(globalState)
                d_720_v60_ = out6_
                d_721_v61_ = out7_
                d_722_v62_: _dafny.MultiSet
                d_722_v62_ = _dafny.MultiSet([(self).f13])
                d_723_v63_: _dafny.Map
                d_723_v63_ = _dafny.Map({d_722_v62_: (self).f13})
                d_724_v64_: _dafny.Map
                d_724_v64_ = _dafny.Map({(self).fm33(p2, globalState): p2})
                index82_ = default__.safeIndex(991, (d_719_v59_).length(0))
                rhs77_ = (self).f26
                rhs78_ = (((d_723_v63_)[(D1_DC5(d_722_v62_, d_721_v61_)).cf10] if ((D1_DC5(d_722_v62_, d_721_v61_)).cf10) in (d_723_v63_) else (self).f12)) >= ((0) - (((d_724_v64_)[p0] if (p0) in (d_724_v64_) else p2)))
                rhs79_ = d_721_v61_
                rhs80_ = p0
                rhs81_ = False
                lhs49_ = d_719_v59_
                lhs50_ = default__.safeIndex(991, (d_719_v59_).length(0))
                lhs51_ = self
                lhs49_[lhs50_] = rhs77_
                lhs51_.f11 = rhs78_
                r3 = rhs79_
                d_721_v61_ = rhs80_
                r3 = rhs81_
                index83_ = default__.safeIndex(991, (d_719_v59_).length(0))
                (d_719_v59_)[index83_] = (self).f13
                d_725_v65_: D4
                d_725_v65_ = D4_DC13((0) - ((d_719_v59_)[default__.safeIndex(991, (d_719_v59_).length(0))]))
                d_725_v65_ = d_725_v65_
            elif True:
                d_685_v29_ = d_685_v29_
                d_726_v66_: _dafny.MultiSet
                d_726_v66_ = _dafny.MultiSet([self.f11])
                d_727_v67_: _dafny.Set
                d_727_v67_ = _dafny.Set({(d_726_v66_).cardinality, len(d_645_v1_), (self).f26, (0) - ((self).f26), default__.fm2(globalState)})
                d_728_v68_: _dafny.Seq
                d_728_v68_ = _dafny.SeqWithoutIsStrInference([p1, (self).f13])
                def iife72_():
                    coll50_ = _dafny.Set()
                    compr_50_: int
                    for compr_50_ in (d_728_v68_).Elements:
                        d_729_v69_: int = compr_50_
                        if (d_729_v69_) in (d_728_v68_):
                            coll50_ = coll50_.union(_dafny.Set([(d_729_v69_) + (956)]))
                    return _dafny.Set(coll50_)
                def iife73_():
                    coll51_ = _dafny.Set()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(896, 705):
                        d_730_v70_: int = compr_51_
                        if ((896) <= (d_730_v70_)) and ((d_730_v70_) < (705)):
                            coll51_ = coll51_.union(_dafny.Set([(d_730_v70_) + ((self).f26)]))
                    return _dafny.Set(coll51_)
                def iife74_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(352, 698):
                        d_731_v71_: int = compr_52_
                        if ((352) <= (d_731_v71_)) and ((d_731_v71_) < (698)):
                            coll52_ = coll52_.union(_dafny.Set([(d_731_v71_) * (len(_dafny.Map({D4_DC12(self.f11, len((self).f20), len(_dafny.SeqWithoutIsStrInference([self.f11]))): d_726_v66_})))]))
                    return _dafny.Set(coll52_)
                d_727_v67_ = ((iife72_()
                ) | (iife73_()
                )).intersection(iife74_()
                )
                d_732_v72_: _dafny.Array
                def lambda54_(d_733_i7_):
                    return not(self.f11)

                init28_ = lambda54_
                nw109_ = _dafny.Array(None, 12)
                for i0_28_ in range(nw109_.length(0)):
                    nw109_[i0_28_] = init28_(i0_28_)
                d_732_v72_ = nw109_
                d_734_v73_: _dafny.Seq
                d_734_v73_ = _dafny.SeqWithoutIsStrInference([d_732_v72_, d_732_v72_])
                d_732_v72_ = (d_734_v73_)[default__.safeIndex((0) - (p1), len(d_734_v73_))]
                r2 = len(((_dafny.SeqWithoutIsStrInference([d_685_v29_ for d_735_i8_ in range(default__.abs(60))])) + (((self).f20) + ((self).f20))).set(default__.safeIndex(p2, len((_dafny.SeqWithoutIsStrInference([d_685_v29_ for d_736_i8_ in range(default__.abs(60))])) + (((self).f20) + ((self).f20)))), ((self).f20)[default__.safeIndex((self).f12, len((self).f20))]))
                d_737_v74_: _dafny.Seq
                d_737_v74_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11, p0])
                d_738_v75_: _dafny.Array
                def lambda55_(d_739_v68_):
                    def lambda56_(d_740_i9_):
                        return (d_740_i9_) - ((d_739_v68_)[default__.safeIndex(341, len(d_739_v68_))])

                    return lambda56_

                init29_ = lambda55_(d_728_v68_)
                nw110_ = _dafny.Array(None, 23)
                for i0_29_ in range(nw110_.length(0)):
                    nw110_[i0_29_] = init29_(i0_29_)
                d_738_v75_ = nw110_
                d_741_v76_: _dafny.Array
                nw111_ = _dafny.Array(None, 22)
                nw111_[int(0)] = d_738_v75_
                nw111_[int(1)] = d_738_v75_
                nw111_[int(2)] = d_738_v75_
                nw111_[int(3)] = d_738_v75_
                nw111_[int(4)] = d_738_v75_
                nw111_[int(5)] = d_738_v75_
                nw111_[int(6)] = d_738_v75_
                nw111_[int(7)] = d_738_v75_
                nw111_[int(8)] = d_738_v75_
                nw111_[int(9)] = d_738_v75_
                nw111_[int(10)] = d_738_v75_
                nw111_[int(11)] = d_738_v75_
                nw111_[int(12)] = (d_738_v75_ if self.f11 else d_738_v75_)
                nw111_[int(13)] = d_738_v75_
                nw111_[int(14)] = d_738_v75_
                nw111_[int(15)] = d_738_v75_
                nw111_[int(16)] = d_738_v75_
                nw111_[int(17)] = d_738_v75_
                nw111_[int(18)] = d_738_v75_
                nw111_[int(19)] = d_738_v75_
                nw111_[int(20)] = d_738_v75_
                nw111_[int(21)] = d_738_v75_
                d_741_v76_ = nw111_
                index84_ = default__.safeIndex(705, (d_741_v76_).length(0))
                (d_741_v76_)[index84_] = d_738_v75_
                d_742_v77_: D6
                d_742_v77_ = D6_DC18(p0)
                index85_ = default__.safeIndex(705, (d_741_v76_).length(0))
                rhs82_ = self.f11
                rhs83_ = True
                rhs84_ = self.f11
                rhs85_ = ((_dafny.SeqWithoutIsStrInference([p0, p0, p0, (d_742_v77_).cf32])) + (d_737_v74_)).set(default__.safeIndex((self).f13, len((_dafny.SeqWithoutIsStrInference([p0, p0, p0, (d_742_v77_).cf32])) + (d_737_v74_))), self.f11)
                rhs86_ = d_738_v75_
                lhs52_ = globalState
                lhs53_ = globalState
                lhs54_ = globalState
                lhs55_ = d_741_v76_
                lhs56_ = default__.safeIndex(705, (d_741_v76_).length(0))
                lhs52_.f8 = rhs82_
                lhs53_.f8 = rhs83_
                lhs54_.f8 = rhs84_
                d_737_v74_ = rhs85_
                lhs55_[lhs56_] = rhs86_
            d_743_v78_: _dafny.Set
            d_743_v78_ = _dafny.Set({self.f11})
            d_744_v79_: _dafny.Seq
            d_744_v79_ = _dafny.SeqWithoutIsStrInference([(self).f12, len(d_743_v78_)])
            d_745_v80_: _dafny.Set
            d_745_v80_ = _dafny.Set({d_744_v79_, d_744_v79_})
            d_746_v81_: _dafny.Map
            d_746_v81_ = _dafny.Map({self.f11: d_744_v79_})
            (globalState).f8 = ((_dafny.Set({((d_746_v81_)[p0] if (p0) in (d_746_v81_) else d_744_v79_), d_744_v79_, _dafny.SeqWithoutIsStrInference([(self).f13 for d_747_i10_ in range(default__.abs(860))])})).intersection(d_745_v80_)).ispropersubset(d_745_v80_)
            r2 = default__.fm2(globalState)
            d_748_v82_: _dafny.Seq
            d_748_v82_ = _dafny.SeqWithoutIsStrInference([p0])
            d_748_v82_ = ((d_748_v82_) + (d_748_v82_)) + (d_748_v82_)
        r3 = p0
        d_749_v83_: _dafny.Map
        d_749_v83_ = _dafny.Map({p0: 619})
        d_750_v84_: D3
        d_750_v84_ = D3_DC8(d_749_v83_)
        d_750_v84_ = default__.fm40((d_685_v29_) in (((self).f20).set(default__.safeIndex((self).f12, len((self).f20)), _dafny.CodePoint('y'))), globalState)
        d_751_v85_: _dafny.Map
        d_751_v85_ = _dafny.Map({(self).f20: ((self).f26) == (905)})
        d_752_v87_: _dafny.Map
        d_752_v87_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_685_v29_ for d_753_i11_ in range(default__.abs(-311))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))), len(_dafny.SeqWithoutIsStrInference([d_685_v29_ for d_754_i11_ in range(default__.abs(-311))]))), d_685_v29_): (self).f13})
        def iife75_():
            coll53_ = _dafny.Map()
            compr_53_: _dafny.Seq
            for compr_53_ in ((d_752_v87_) | (d_752_v87_)).keys.Elements:
                d_755_v86_: _dafny.Seq = compr_53_
                if (d_755_v86_) in ((d_752_v87_) | (d_752_v87_)):
                    coll53_[d_755_v86_] = not (p0) or (self.f11)
            return _dafny.Map(coll53_)
        d_751_v85_ = iife75_()
        
        r0 = (self).f20
        d_756_v88_: _dafny.Map
        d_756_v88_ = _dafny.Map({(self).f12: (self).f12})
        d_757_v89_: _dafny.Map
        d_757_v89_ = _dafny.Map({p2: d_756_v88_})
        d_758_v90_: C0
        nw112_ = C0()
        nw112_.ctor__(len((self).f20), False, (self).f13, (self).f13)
        d_758_v90_ = nw112_
        d_759_v91_: D3
        d_759_v91_ = D3_DC10(D3_DC9(p1, d_757_v89_, default__.fm2(globalState), len(_dafny.Map({(self).fm7(globalState): d_758_v90_}))))
        d_760_v92_: D3
        d_760_v92_ = D3_DC8(d_749_v83_)
        pat_let_tv16_ = d_760_v92_
        def iife76_(_pat_let11_0):
            def iife77_(d_761_dt__update__tmp_h0_):
                def iife78_(_pat_let12_0):
                    def iife79_(d_762_dt__update_hcf19_h0_):
                        return D3_DC10(d_762_dt__update_hcf19_h0_)
                    return iife79_(_pat_let12_0)
                return iife78_(pat_let_tv16_)
            return iife77_(_pat_let11_0)
        r1 = iife76_(d_759_v91_)
        d_763_v93_: _dafny.Map
        d_763_v93_ = _dafny.Map({d_758_v90_.f29: (self).f20})
        d_764_v94_: _dafny.Seq
        d_764_v94_ = _dafny.SeqWithoutIsStrInference([p1, (d_758_v90_).f28, 385, len(((d_763_v93_)[p0] if (p0) in (d_763_v93_) else (self).f20))])
        d_765_v95_: _dafny.Map
        d_765_v95_ = _dafny.Map({((0) - (len(d_764_v94_))) * ((self).f13): d_758_v90_.f29})
        r2 = len(d_765_v95_)
        r3 = (p1) <= (240)
        return r0, r1, r2, r3

    @property
    def f26(self):
        return self._f26

class C3(T1):
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def fm7(self, globalState):
        return (D0_DC3((self).f13, not(True))).cf8

    def fm8(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qei"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_766_i0_ in range(default__.abs(-489))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lotyl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "st"))))

    def m22(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: _dafny.Set = _dafny.Set({})
        d_767_v0_: _dafny.Map
        d_767_v0_ = _dafny.Map({(self).f13: 455})
        d_768_v1_: _dafny.Map
        d_768_v1_ = _dafny.Map({(self).f13: d_767_v0_})
        d_769_v2_: D3
        d_769_v2_ = D3_DC9((self).f13, d_768_v1_, (self).f12, (self).f12)
        d_770_v3_: D3
        d_770_v3_ = D3_DC10(d_769_v2_)
        d_771_v4_: D3
        d_771_v4_ = D3_DC10(d_769_v2_)
        source13_ = d_771_v4_
        if source13_.is_DC9:
            d_772___mcc_h0_ = source13_.cf15
            d_773___mcc_h1_ = source13_.cf16
            d_774___mcc_h2_ = source13_.cf17
            d_775___mcc_h3_ = source13_.cf18
            d_776_cf18_ = d_775___mcc_h3_
            d_777_cf17_ = d_774___mcc_h2_
            d_778_cf16_ = d_773___mcc_h1_
            d_779_cf15_ = d_772___mcc_h0_
            d_780_v5_: str
            d_780_v5_ = _dafny.CodePoint('j')
            d_781_v6_: _dafny.Set
            d_781_v6_ = _dafny.Set({d_780_v5_, d_780_v5_, d_780_v5_})
            if not((_dafny.Set({_dafny.CodePoint('g'), d_780_v5_, d_780_v5_, d_780_v5_, d_780_v5_})).isdisjoint(d_781_v6_)):
                d_777_cf17_ = (self).f13
                d_782_v7_: bool
                d_782_v7_ = False
                d_783_v8_: _dafny.Map
                d_783_v8_ = _dafny.Map({not(d_782_v7_): d_782_v7_})
                d_784_v9_: _dafny.Array
                nw113_ = _dafny.Array(None, 3)
                nw113_[int(0)] = d_782_v7_
                nw113_[int(1)] = d_782_v7_
                nw113_[int(2)] = d_782_v7_
                d_784_v9_ = nw113_
                d_785_v10_: _dafny.Map
                d_785_v10_ = _dafny.Map({len(d_783_v8_): d_784_v9_})
                d_786_v11_: C0
                nw114_ = C0()
                nw114_.ctor__(len((d_785_v10_) | (d_785_v10_)), d_782_v7_, 463, d_777_cf17_)
                d_786_v11_ = nw114_
                d_787_v12_: _dafny.Array
                def lambda57_(d_788_v11_):
                    def lambda58_(d_789_i0_):
                        return (d_789_i0_) - ((d_788_v11_).f28)

                    return lambda58_

                init30_ = lambda57_(d_786_v11_)
                nw115_ = _dafny.Array(None, 11)
                for i0_30_ in range(nw115_.length(0)):
                    nw115_[i0_30_] = init30_(i0_30_)
                d_787_v12_ = nw115_
                r1 = d_787_v12_
                d_790_v13_: _dafny.Seq
                d_790_v13_ = _dafny.SeqWithoutIsStrInference([not(d_786_v11_.f29)])
                d_791_v14_: _dafny.MultiSet
                d_791_v14_ = _dafny.MultiSet([(self).f13])
                d_792_v15_: _dafny.Array
                nw116_ = _dafny.Array(None, 21)
                nw116_[int(0)] = _dafny.SeqWithoutIsStrInference([d_782_v7_, d_782_v7_])
                nw116_[int(1)] = d_790_v13_
                nw116_[int(2)] = (d_790_v13_ if d_786_v11_.f29 else d_790_v13_)
                nw116_[int(3)] = d_790_v13_
                nw116_[int(4)] = _dafny.SeqWithoutIsStrInference([True, False, d_782_v7_, d_786_v11_.f29])
                nw116_[int(5)] = d_790_v13_
                nw116_[int(6)] = d_790_v13_
                nw116_[int(7)] = _dafny.SeqWithoutIsStrInference([d_782_v7_, d_782_v7_])
                nw116_[int(8)] = d_790_v13_
                nw116_[int(9)] = d_790_v13_
                nw116_[int(10)] = d_790_v13_
                nw116_[int(11)] = d_790_v13_
                nw116_[int(12)] = d_790_v13_
                nw116_[int(13)] = (d_790_v13_) + (d_790_v13_)
                nw116_[int(14)] = d_790_v13_
                nw116_[int(15)] = _dafny.SeqWithoutIsStrInference([d_782_v7_, d_786_v11_.f29])
                nw116_[int(16)] = (d_790_v13_).set(default__.safeIndex((d_786_v11_).f28, len(d_790_v13_)), d_786_v11_.f29)
                nw116_[int(17)] = d_790_v13_
                nw116_[int(18)] = d_790_v13_
                nw116_[int(19)] = d_790_v13_
                nw116_[int(20)] = ((d_790_v13_) + (d_790_v13_)).set(default__.safeIndex(((p0).set(d_791_v14_, default__.abs(d_777_cf17_))).cardinality, len((d_790_v13_) + (d_790_v13_))), d_782_v7_)
                d_792_v15_ = nw116_
                d_793_v16_: _dafny.Seq
                d_793_v16_ = _dafny.SeqWithoutIsStrInference([d_776_cf18_, 579, (d_786_v11_).f28, (self).f12, d_777_cf17_])
                d_794_v17_: _dafny.MultiSet
                d_794_v17_ = _dafny.MultiSet([False])
                rhs87_ = d_792_v15_
                rhs88_ = (default__.fm2(globalState)) - (d_779_cf15_)
                rhs89_ = ((d_793_v16_) + (d_793_v16_)) < (_dafny.SeqWithoutIsStrInference([d_776_cf18_]))
                rhs90_ = ((_dafny.MultiSet([d_782_v7_])).ispropersubset(d_794_v17_)) == ((d_793_v16_) == (d_793_v16_))
                lhs57_ = globalState
                lhs58_ = globalState
                d_792_v15_ = rhs87_
                d_777_cf17_ = rhs88_
                lhs57_.f8 = rhs89_
                lhs58_.f8 = rhs90_
                d_795_v19_: _dafny.Map
                def iife80_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(601, 209):
                        d_796_v18_: int = compr_54_
                        if ((601) <= (d_796_v18_)) and ((d_796_v18_) < (209)):
                            coll54_[(d_796_v18_) - (len(d_793_v16_))] = d_782_v7_
                    return _dafny.Map(coll54_)
                d_795_v19_ = _dafny.Map({(d_779_cf15_) <= (d_776_cf18_): iife80_()
                })
                d_797_v20_: _dafny.Map
                d_797_v20_ = _dafny.Map({(0) - ((0) - ((d_786_v11_).f28)): d_786_v11_.f29})
                d_795_v19_ = (d_795_v19_).set(d_782_v7_, d_797_v20_)
            elif True:
                d_798_v23_: _dafny.Array
                def lambda59_(d_799_cf18_, d_800_cf15_):
                    def lambda60_(d_801_i1_):
                        def iife81_():
                            coll55_ = _dafny.Map()
                            compr_55_: int
                            for compr_55_ in (_dafny.SeqWithoutIsStrInference([d_799_cf18_])).Elements:
                                d_802_v21_: int = compr_55_
                                if (d_802_v21_) in (_dafny.SeqWithoutIsStrInference([d_799_cf18_])):
                                    coll55_[(d_802_v21_) + (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([(self).f13])})))] = _dafny.SeqWithoutIsStrInference([D6_DC18(False), D6_DC18(not(True))])
                            return _dafny.Map(coll55_)
                        def iife82_():
                            coll56_ = _dafny.Map()
                            compr_56_: int
                            for compr_56_ in (_dafny.Set({(self).f13, d_800_cf15_, d_800_cf15_, d_800_cf15_, len(_dafny.SeqWithoutIsStrInference([True, False]))})).Elements:
                                d_803_v22_: int = compr_56_
                                if (d_803_v22_) in (_dafny.Set({(self).f13, d_800_cf15_, d_800_cf15_, d_800_cf15_, len(_dafny.SeqWithoutIsStrInference([True, False]))})):
                                    coll56_[(d_803_v22_) + (d_799_cf18_)] = _dafny.SeqWithoutIsStrInference([D6_DC18(False)])
                            return _dafny.Map(coll56_)
                        return (iife81_()
                        ) | (iife82_()
                        )

                    return lambda60_

                init31_ = lambda59_(d_776_cf18_, d_779_cf15_)
                nw117_ = _dafny.Array(None, 20)
                for i0_31_ in range(nw117_.length(0)):
                    nw117_[i0_31_] = init31_(i0_31_)
                d_798_v23_ = nw117_
                d_804_v24_: _dafny.Map
                d_804_v24_ = _dafny.Map({455: default__.fm41(globalState)})
                index86_ = default__.safeIndex(938, (d_798_v23_).length(0))
                (d_798_v23_)[index86_] = d_804_v24_
                index87_ = default__.safeIndex(938, (d_798_v23_).length(0))
                (d_798_v23_)[index87_] = (d_804_v24_) | (d_804_v24_)
                d_805_v25_: bool
                d_805_v25_ = False
                d_806_v26_: _dafny.Array
                nw118_ = _dafny.Array(None, 6)
                nw118_[int(0)] = default__.fm1((self).f12, d_805_v25_, len(_dafny.SeqWithoutIsStrInference([d_780_v5_ for d_807_i2_ in range(default__.abs(80))])), globalState)
                nw118_[int(1)] = d_805_v25_
                nw118_[int(2)] = d_805_v25_
                nw118_[int(3)] = d_805_v25_
                nw118_[int(4)] = d_805_v25_
                nw118_[int(5)] = d_805_v25_
                d_806_v26_ = nw118_
                index88_ = default__.safeIndex(225, (d_806_v26_).length(0))
                (d_806_v26_)[index88_] = d_805_v25_
                d_808_v27_: _dafny.MultiSet
                d_808_v27_ = _dafny.MultiSet([d_779_cf15_, -893])
                d_809_v28_: _dafny.Array
                nw119_ = _dafny.Array(None, 8)
                nw119_[int(0)] = len(_dafny.SeqWithoutIsStrInference([((d_808_v27_)[d_777_cf17_] if (d_777_cf17_) in (d_808_v27_) else d_777_cf17_)]))
                nw119_[int(1)] = d_779_cf15_
                nw119_[int(2)] = (self).f12
                nw119_[int(3)] = d_779_cf15_
                nw119_[int(4)] = 338
                nw119_[int(5)] = (self).f13
                nw119_[int(6)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))), _dafny.CodePoint('w')))
                nw119_[int(7)] = 625
                d_809_v28_ = nw119_
                d_810_v29_: _dafny.Array
                nw120_ = _dafny.Array(None, 7)
                nw120_[int(0)] = d_809_v28_
                nw120_[int(1)] = d_809_v28_
                nw120_[int(2)] = d_809_v28_
                nw120_[int(3)] = d_809_v28_
                nw120_[int(4)] = (d_809_v28_ if d_805_v25_ else d_809_v28_)
                nw120_[int(5)] = d_809_v28_
                nw120_[int(6)] = d_809_v28_
                d_810_v29_ = nw120_
                index89_ = default__.safeIndex(950, (d_810_v29_).length(0))
                (d_810_v29_)[index89_] = d_809_v28_
                index90_ = default__.safeIndex(225, (d_806_v26_).length(0))
                index91_ = default__.safeIndex(950, (d_810_v29_).length(0))
                rhs91_ = d_805_v25_
                rhs92_ = not(d_805_v25_)
                rhs93_ = d_809_v28_
                lhs59_ = globalState
                lhs60_ = d_806_v26_
                lhs61_ = default__.safeIndex(225, (d_806_v26_).length(0))
                lhs62_ = d_810_v29_
                lhs63_ = default__.safeIndex(950, (d_810_v29_).length(0))
                lhs59_.f8 = rhs91_
                lhs60_[lhs61_] = rhs92_
                lhs62_[lhs63_] = rhs93_
                d_811_v30_: _dafny.Map
                d_811_v30_ = _dafny.Map({(d_806_v26_)[default__.safeIndex(225, (d_806_v26_).length(0))]: d_776_cf18_})
                d_812_v31_: _dafny.MultiSet
                d_812_v31_ = _dafny.MultiSet([d_811_v30_])
                d_813_v32_: C0
                nw121_ = C0()
                nw121_.ctor__(d_776_cf18_, (d_806_v26_)[default__.safeIndex(225, (d_806_v26_).length(0))], d_779_cf15_, (d_812_v31_).cardinality)
                d_813_v32_ = nw121_
                d_814_v33_: _dafny.Array
                nw122_ = _dafny.Array(_dafny.Map({}), 15)
                d_814_v33_ = nw122_
                d_815_v34_: _dafny.Array
                nw123_ = _dafny.Array(None, 4)
                nw123_[int(0)] = d_814_v33_
                nw123_[int(1)] = d_814_v33_
                nw123_[int(2)] = d_814_v33_
                nw123_[int(3)] = d_814_v33_
                d_815_v34_ = nw123_
                index92_ = default__.safeIndex(450, (d_815_v34_).length(0))
                (d_815_v34_)[index92_] = d_814_v33_
                index93_ = default__.safeIndex(450, (d_815_v34_).length(0))
                (d_815_v34_)[index93_] = d_814_v33_
                d_816_v35_: _dafny.Seq
                d_816_v35_ = _dafny.SeqWithoutIsStrInference([(D0_DC2(d_779_cf15_, d_809_v28_)).cf5])
                d_779_cf15_ = (len((d_816_v35_) + (d_816_v35_))) + (92)
            d_817_v36_: bool
            d_817_v36_ = False
            d_818_v37_: _dafny.Map
            d_818_v37_ = _dafny.Map({d_817_v36_: (self).f13})
            d_819_v38_: D3
            d_819_v38_ = D3_DC8(d_818_v37_)
            pat_let_tv17_ = d_817_v36_
            pat_let_tv18_ = d_776_cf18_
            pat_let_tv19_ = d_818_v37_
            d_820_v39_: _dafny.Map
            def iife84_(_pat_let14_0):
                def iife85_(d_821_dt__update__tmp_h0_):
                    def iife86_(_pat_let15_0):
                        def iife87_(d_822_dt__update_hcf14_h0_):
                            return D3_DC8(d_822_dt__update_hcf14_h0_)
                        return iife87_(_pat_let15_0)
                    return iife86_(_dafny.Map({pat_let_tv17_: pat_let_tv18_}))
                return iife85_(_pat_let14_0)
            def iife83_(_pat_let13_0):
                def iife88_(d_823_dt__update__tmp_h1_):
                    def iife89_(_pat_let16_0):
                        def iife90_(d_824_dt__update_hcf14_h1_):
                            return D3_DC8(d_824_dt__update_hcf14_h1_)
                        return iife90_(_pat_let16_0)
                    return iife89_(pat_let_tv19_)
                return iife88_(_pat_let13_0)
            d_820_v39_ = _dafny.Map({(0) - (d_776_cf18_): iife83_(iife84_(d_819_v38_))})
            d_820_v39_ = (d_820_v39_).set(default__.fm2(globalState), d_819_v38_)
            d_825_v40_: _dafny.Array
            def lambda61_(d_826_v36_):
                def lambda62_(d_827_i3_):
                    return _dafny.Map({d_826_v36_: d_826_v36_})

                return lambda62_

            init32_ = lambda61_(d_817_v36_)
            nw124_ = _dafny.Array(None, 7)
            for i0_32_ in range(nw124_.length(0)):
                nw124_[i0_32_] = init32_(i0_32_)
            d_825_v40_ = nw124_
            d_828_v41_: _dafny.Map
            d_828_v41_ = _dafny.Map({True: d_817_v36_})
            index94_ = default__.safeIndex(504, (d_825_v40_).length(0))
            (d_825_v40_)[index94_] = (d_828_v41_ if d_817_v36_ else d_828_v41_)
            index95_ = default__.safeIndex(504, (d_825_v40_).length(0))
            (d_825_v40_)[index95_] = (d_828_v41_) | (default__.fm42(d_779_cf15_, (self).f12, globalState))
            d_829_v42_: _dafny.Array
            nw125_ = _dafny.Array(False, 2)
            d_829_v42_ = nw125_
            index96_ = default__.safeIndex(789, (d_829_v42_).length(0))
            (d_829_v42_)[index96_] = ((0) - (d_779_cf15_)) == (d_779_cf15_)
            d_830_v43_: _dafny.Array
            nw126_ = _dafny.Array(int(0), 10)
            d_830_v43_ = nw126_
            d_831_v44_: _dafny.Seq
            d_831_v44_ = _dafny.SeqWithoutIsStrInference([len(d_767_v0_)])
            index97_ = default__.safeIndex(789, (d_829_v42_).length(0))
            rhs94_ = (self).f12
            rhs95_ = ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_830_v43_})), (self).f13, d_776_cf18_, 476, (self).f13])) + (default__.fm0(globalState))) == ((d_831_v44_) + (d_831_v44_))
            lhs64_ = d_829_v42_
            lhs65_ = default__.safeIndex(789, (d_829_v42_).length(0))
            d_776_cf18_ = rhs94_
            lhs64_[lhs65_] = rhs95_
        elif source13_.is_DC8:
            d_832___mcc_h4_ = source13_.cf14
            d_833_cf14_ = d_832___mcc_h4_
            d_834_v45_: _dafny.Array
            nw127_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_834_v45_ = nw127_
            d_835_v46_: _dafny.Array
            nw128_ = _dafny.Array(int(0), 29)
            d_835_v46_ = nw128_
            index98_ = default__.safeIndex(789, (d_834_v45_).length(0))
            (d_834_v45_)[index98_] = d_835_v46_
            index99_ = default__.safeIndex(789, (d_834_v45_).length(0))
            (d_834_v45_)[index99_] = d_835_v46_
            d_836_v47_: C0
            nw129_ = C0()
            nw129_.ctor__(-25, not(False), (self).f12, (len(_dafny.SeqWithoutIsStrInference([True]))) - ((self).f12))
            d_836_v47_ = nw129_
            d_837_v48_: str
            d_837_v48_ = _dafny.CodePoint('a')
            d_838_v49_: _dafny.Set
            d_838_v49_ = _dafny.Set({d_837_v48_})
            d_839_v50_: _dafny.Map
            d_839_v50_ = _dafny.Map({(d_836_v47_).f28: d_838_v49_})
            d_840_v51_: _dafny.Seq
            d_840_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_841_v52_: D6
            d_841_v52_ = D6_DC20(default__.fm43((self).f13, d_836_v47_.f29, len(((d_839_v50_)[297] if (297) in (d_839_v50_) else d_838_v49_)), d_840_v51_, globalState))
            d_842_v53_: D6
            d_842_v53_ = D6_DC18(False)
            d_843_v54_: _dafny.Set
            d_843_v54_ = _dafny.Set({d_842_v53_})
            d_844_v55_: D6
            d_844_v55_ = D6_DC18(d_836_v47_.f29)
            d_845_v56_: D6
            d_845_v56_ = D6_DC20(d_844_v55_)
            rhs96_ = d_836_v47_.f29
            rhs97_ = (d_840_v51_) == (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqio"))) + (d_840_v51_)).set(default__.safeIndex((d_836_v47_).f28, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqio"))) + (d_840_v51_))), d_837_v48_))
            rhs98_ = (d_843_v54_).issubset(default__.fm44(globalState))
            rhs99_ = D6_DC20(d_845_v56_)
            lhs66_ = globalState
            r0 = rhs96_
            r0 = rhs97_
            lhs66_.f8 = rhs98_
            d_841_v52_ = rhs99_
            if d_836_v47_.f29:
                d_846_v57_: int
                d_846_v57_ = 494
                d_847_v58_: _dafny.Seq
                d_847_v58_ = _dafny.SeqWithoutIsStrInference([d_835_v46_])
                d_848_v59_: D1
                d_848_v59_ = D1_DC6(D1_DC4(d_847_v58_))
                d_849_v60_: _dafny.Seq
                d_849_v60_ = _dafny.SeqWithoutIsStrInference([d_848_v59_])
                d_850_v61_: _dafny.Map
                d_850_v61_ = _dafny.Map({(self).f13: d_849_v60_})
                d_851_v62_: _dafny.Seq
                d_851_v62_ = _dafny.SeqWithoutIsStrInference([d_846_v57_])
                d_852_v63_: _dafny.Seq
                d_852_v63_ = _dafny.SeqWithoutIsStrInference([len(d_851_v62_)])
                d_853_v64_: _dafny.Seq
                d_853_v64_ = _dafny.SeqWithoutIsStrInference([len(d_851_v62_), (d_836_v47_).f28])
                d_854_v65_: D4
                d_854_v65_ = D4_DC12(not(default__.fm1((self).f12, default__.fm1(len(d_852_v63_), (d_836_v47_).fm7(globalState), default__.fm2(globalState), globalState), len(d_853_v64_), globalState)), len(d_833_cf14_), (d_836_v47_).f28)
                rhs100_ = ((d_836_v47_).f28) * (571)
                rhs101_ = ((self).f13) - ((0) - ((d_836_v47_).f28))
                rhs102_ = (d_854_v65_).cf22
                rhs103_ = d_850_v61_
                rhs104_ = d_840_v51_
                d_846_v57_ = rhs100_
                d_846_v57_ = rhs101_
                d_846_v57_ = rhs102_
                d_850_v61_ = rhs103_
                d_840_v51_ = rhs104_
                d_855_v66_: D5
                d_855_v66_ = D5_DC15((d_836_v47_).f28, (d_836_v47_).f28, 156)
                pat_let_tv20_ = d_836_v47_
                pat_let_tv21_ = d_833_cf14_
                pat_let_tv22_ = d_836_v47_
                pat_let_tv23_ = d_767_v0_
                d_856_v67_: _dafny.Array
                nw130_ = _dafny.Array(None, 5)
                def iife92_(_pat_let18_0):
                    def iife93_(d_857_dt__update__tmp_h2_):
                        def iife94_(_pat_let19_0):
                            def iife95_(d_858_dt__update_hcf26_h0_):
                                return D5_DC15(d_858_dt__update_hcf26_h0_, (d_857_dt__update__tmp_h2_).cf27, (d_857_dt__update__tmp_h2_).cf28)
                            return iife95_(_pat_let19_0)
                        return iife94_((pat_let_tv20_).f28)
                    return iife93_(_pat_let18_0)
                def iife91_(_pat_let17_0):
                    def iife96_(d_859_dt__update__tmp_h3_):
                        def iife97_(_pat_let20_0):
                            def iife98_(d_860_dt__update_hcf26_h1_):
                                def iife99_(_pat_let21_0):
                                    def iife100_(d_861_dt__update_hcf28_h0_):
                                        return D5_DC15(d_860_dt__update_hcf26_h1_, (d_859_dt__update__tmp_h3_).cf27, d_861_dt__update_hcf28_h0_)
                                    return iife100_(_pat_let21_0)
                                return iife99_((pat_let_tv22_).f28)
                            return iife98_(_pat_let20_0)
                        return iife97_((0) - (len(pat_let_tv21_)))
                    return iife96_(_pat_let17_0)
                nw130_[int(0)] = iife91_(iife92_(d_855_v66_))
                nw130_[int(1)] = d_855_v66_
                nw130_[int(2)] = D5_DC15((d_836_v47_).f28, (self).f12, (d_836_v47_).f28)
                nw130_[int(3)] = D5_DC15((self).f13, (d_836_v47_).f28, 119)
                def iife101_(_pat_let22_0):
                    def iife102_(d_862_dt__update__tmp_h4_):
                        def iife103_(_pat_let23_0):
                            def iife104_(d_863_dt__update_hcf28_h1_):
                                return D5_DC15((d_862_dt__update__tmp_h4_).cf26, (d_862_dt__update__tmp_h4_).cf27, d_863_dt__update_hcf28_h1_)
                            return iife104_(_pat_let23_0)
                        return iife103_((0) - (len(pat_let_tv23_)))
                    return iife102_(_pat_let22_0)
                nw130_[int(4)] = iife101_(d_855_v66_)
                d_856_v67_ = nw130_
                rhs105_ = d_856_v67_
                rhs106_ = d_836_v47_.f29
                rhs107_ = (d_852_v63_)[default__.safeIndex(876, len(d_852_v63_))]
                d_856_v67_ = rhs105_
                r0 = rhs106_
                d_846_v57_ = rhs107_
                d_864_v68_: _dafny.Map
                d_864_v68_ = _dafny.Map({-872: d_835_v46_})
                d_865_v69_: D3
                d_865_v69_ = D3_DC9((self).f13, d_768_v1_, len(d_853_v64_), (d_836_v47_).f28)
                r1 = ((d_864_v68_)[(d_851_v62_)[default__.safeIndex((d_865_v69_).cf17, len(d_851_v62_))]] if ((d_851_v62_)[default__.safeIndex((d_865_v69_).cf17, len(d_851_v62_))]) in (d_864_v68_) else (d_847_v58_)[default__.safeIndex((d_836_v47_).f28, len(d_847_v58_))])
                (globalState).f8 = d_836_v47_.f29
                d_866_v70_: D1
                d_866_v70_ = D1_DC4((d_847_v58_) + (d_847_v58_))
                d_866_v70_ = d_866_v70_
            elif True:
                d_837_v48_ = default__.fm45((self).f12, globalState)
                d_867_v71_: _dafny.Set
                d_867_v71_ = _dafny.Set({not(d_836_v47_.f29), True})
                d_868_v72_: _dafny.Set
                d_868_v72_ = _dafny.Set({(d_836_v47_).f28, len(d_833_cf14_), (d_836_v47_).f28, (self).f13, default__.fm2(globalState)})
                d_840_v51_ = ((d_836_v47_).fm8(len(d_867_v71_), _dafny.Map({d_837_v48_: d_836_v47_.f29}), d_868_v72_, globalState)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_869_i4_ in range(default__.abs(332))]))
                d_870_v73_: _dafny.Map
                d_870_v73_ = _dafny.Map({(0) - ((d_836_v47_).f28): d_836_v47_.f29})
                d_871_v74_: _dafny.Seq
                d_871_v74_ = _dafny.SeqWithoutIsStrInference([True])
                d_872_v75_: _dafny.MultiSet
                d_872_v75_ = _dafny.MultiSet([((d_870_v73_)[(d_836_v47_).f28] if ((d_836_v47_).f28) in (d_870_v73_) else d_836_v47_.f29), (d_836_v47_).fm7(globalState), d_836_v47_.f29, (d_871_v74_)[default__.safeIndex((self).f12, len(d_871_v74_))], d_836_v47_.f29])
                d_873_v76_: D7
                d_873_v76_ = D7_DC21((d_872_v75_).set(d_836_v47_.f29, default__.abs((self).f13)))
                d_872_v75_ = (d_873_v76_).cf39
                d_874_v77_: int
                d_874_v77_ = 239
                d_875_v78_: _dafny.Array
                nw131_ = _dafny.Array(False, 6)
                d_875_v78_ = nw131_
                d_876_v79_: _dafny.MultiSet
                d_876_v79_ = _dafny.MultiSet([(_dafny.Map({d_836_v47_.f29: d_836_v47_.f29})).set(d_836_v47_.f29, d_836_v47_.f29)])
                d_877_v80_: _dafny.Set
                d_877_v80_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ceucd")), d_840_v51_})
                d_878_v81_: _dafny.Seq
                d_878_v81_ = _dafny.SeqWithoutIsStrInference([d_875_v78_, d_875_v78_])
                d_879_v82_: _dafny.Map
                d_879_v82_ = _dafny.Map({d_836_v47_.f29: True})
                d_880_v83_: _dafny.MultiSet
                d_880_v83_ = _dafny.MultiSet([d_879_v82_])
                rhs108_ = _dafny.SeqWithoutIsStrInference([not((D7_DC23(d_834_v45_, d_871_v74_, d_836_v47_.f29, len(d_877_v80_))).cf44), d_836_v47_.f29, not (d_836_v47_.f29) or (d_836_v47_.f29)])
                rhs109_ = (False if d_836_v47_.f29 else (self).fm7(globalState))
                rhs110_ = default__.fm2(globalState)
                rhs111_ = (d_878_v81_)[default__.safeIndex((0) - ((self).f12), len(d_878_v81_))]
                rhs112_ = ((d_880_v83_)).intersection(_dafny.MultiSet([d_879_v82_]))
                lhs67_ = d_836_v47_
                d_871_v74_ = rhs108_
                lhs67_.f29 = rhs109_
                d_874_v77_ = rhs110_
                d_875_v78_ = rhs111_
                d_876_v79_ = rhs112_
                d_881_v84_: _dafny.Array
                nw132_ = _dafny.Array(_dafny.MultiSet({}), 20)
                d_881_v84_ = nw132_
                d_882_v85_: _dafny.Map
                d_882_v85_ = _dafny.Map({d_881_v84_: d_871_v74_})
                d_871_v74_ = ((d_882_v85_)[d_881_v84_] if (d_881_v84_) in (d_882_v85_) else _dafny.SeqWithoutIsStrInference([d_836_v47_.f29, d_836_v47_.f29, d_836_v47_.f29]))
        elif True:
            d_883___mcc_h5_ = source13_.cf19
            d_884_cf19_ = d_883___mcc_h5_
            d_885_v86_: int
            d_885_v86_ = -242
            d_886_v87_: _dafny.Array
            nw133_ = _dafny.Array(False, 28)
            d_886_v87_ = nw133_
            index100_ = default__.safeIndex(920, (d_886_v87_).length(0))
            (d_886_v87_)[index100_] = True
            d_887_v88_: _dafny.Seq
            d_887_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghpgbut"))
            index101_ = default__.safeIndex(920, (d_886_v87_).length(0))
            rhs113_ = (self).f12
            rhs114_ = (default__.safeModuloInt(-326, (self).f13)) * ((-870) + (default__.fm2(globalState)))
            rhs115_ = (d_885_v86_) >= (len(d_887_v88_))
            lhs68_ = d_886_v87_
            lhs69_ = default__.safeIndex(920, (d_886_v87_).length(0))
            d_885_v86_ = rhs113_
            d_885_v86_ = rhs114_
            lhs68_[lhs69_] = rhs115_
            d_886_v87_ = d_886_v87_
            d_888_v89_: T0
            nw134_ = C2()
            nw134_.ctor__((-247 if (d_886_v87_)[default__.safeIndex(920, (d_886_v87_).length(0))] else -417), (d_887_v88_) + (_dafny.SeqWithoutIsStrInference([(d_887_v88_)[default__.safeIndex((self).f12, len(d_887_v88_))] for d_889_i5_ in range(default__.abs(870))])), (self).f13, d_885_v86_, (463) < (d_885_v86_))
            d_888_v89_ = nw134_
            d_888_v89_ = d_888_v89_
            d_890_v90_: _dafny.Seq
            d_890_v90_ = _dafny.SeqWithoutIsStrInference([d_888_v89_.f11, False, (d_886_v87_)[default__.safeIndex(920, (d_886_v87_).length(0))]])
            d_891_v91_: _dafny.MultiSet
            d_891_v91_ = _dafny.MultiSet([d_890_v90_, d_890_v90_, _dafny.SeqWithoutIsStrInference([d_888_v89_.f11])])
            d_892_v92_: _dafny.Map
            d_892_v92_ = _dafny.Map({d_891_v91_: (self).f12})
            d_892_v92_ = (d_892_v92_).set(d_891_v91_, (self).f12)
        d_893_v93_: bool
        d_893_v93_ = False
        d_894_v94_: _dafny.MultiSet
        d_894_v94_ = _dafny.MultiSet([d_893_v93_, d_893_v93_])
        d_895_v95_: _dafny.Seq
        d_895_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfsiu"))
        d_896_v96_: _dafny.Seq
        d_896_v96_ = _dafny.SeqWithoutIsStrInference([d_893_v93_, d_893_v93_])
        d_897_v97_: _dafny.Seq
        d_897_v97_ = _dafny.SeqWithoutIsStrInference([len(d_896_v96_)])
        d_898_v98_: _dafny.Set
        d_898_v98_ = _dafny.Set({len(d_897_v97_), (self).f13})
        d_899_v99_: D6
        d_899_v99_ = D6_DC19(True, not((d_894_v94_).isdisjoint(d_894_v94_)), (d_898_v98_).ispropersubset(_dafny.Set({len((_dafny.Map({default__.fm2(globalState): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_900_i6_ in range(default__.abs(760))])})).set((self).f13, d_895_v95_)), (self).f12, (self).f13})), d_893_v93_, (d_896_v96_) + (d_896_v96_))
        source14_ = d_899_v99_
        if source14_.is_DC17:
            d_901___mcc_h6_ = source14_.cf30
            d_902___mcc_h7_ = source14_.cf31
            d_903_cf31_ = d_902___mcc_h7_
            d_904_cf30_ = d_901___mcc_h6_
            d_903_cf31_ = (0) - ((default__.safeDivisionInt((self).f13, 403)) * ((self).f12))
            rhs116_ = ((d_897_v97_)[default__.safeIndex((self).f13, len(d_897_v97_))]) > ((self).f12)
            rhs117_ = len(d_897_v97_)
            rhs118_ = len(d_898_v98_)
            lhs70_ = globalState
            lhs70_.f8 = rhs116_
            d_903_cf31_ = rhs117_
            d_903_cf31_ = rhs118_
            d_905_v100_: _dafny.Array
            def lambda63_(d_906_v96_):
                def lambda64_(d_907_i7_):
                    return (d_907_i7_) - (len(d_906_v96_))

                return lambda64_

            init33_ = lambda63_(d_896_v96_)
            nw135_ = _dafny.Array(None, 25)
            for i0_33_ in range(nw135_.length(0)):
                nw135_[i0_33_] = init33_(i0_33_)
            d_905_v100_ = nw135_
            d_908_v101_: _dafny.Seq
            d_908_v101_ = _dafny.SeqWithoutIsStrInference([d_905_v100_, d_905_v100_])
            d_909_v102_: _dafny.Map
            d_909_v102_ = _dafny.Map({d_904_cf30_: d_905_v100_})
            d_910_v103_: _dafny.Array
            nw136_ = _dafny.Array(None, 29)
            nw136_[int(0)] = d_905_v100_
            nw136_[int(1)] = d_905_v100_
            nw136_[int(2)] = d_905_v100_
            nw136_[int(3)] = d_905_v100_
            nw136_[int(4)] = (d_908_v101_)[default__.safeIndex(317, len(d_908_v101_))]
            nw136_[int(5)] = d_905_v100_
            nw136_[int(6)] = d_905_v100_
            nw136_[int(7)] = d_905_v100_
            nw136_[int(8)] = d_905_v100_
            nw136_[int(9)] = d_905_v100_
            nw136_[int(10)] = d_905_v100_
            nw136_[int(11)] = d_905_v100_
            nw136_[int(12)] = d_905_v100_
            nw136_[int(13)] = d_905_v100_
            nw136_[int(14)] = d_905_v100_
            nw136_[int(15)] = d_905_v100_
            nw136_[int(16)] = (d_908_v101_)[default__.safeIndex(len(d_898_v98_), len(d_908_v101_))]
            nw136_[int(17)] = d_905_v100_
            nw136_[int(18)] = d_905_v100_
            nw136_[int(19)] = ((d_909_v102_)[not(not(d_893_v93_))] if (not(not(d_893_v93_))) in (d_909_v102_) else d_905_v100_)
            nw136_[int(20)] = d_905_v100_
            nw136_[int(21)] = d_905_v100_
            nw136_[int(22)] = d_905_v100_
            nw136_[int(23)] = d_905_v100_
            nw136_[int(24)] = d_905_v100_
            nw136_[int(25)] = d_905_v100_
            nw136_[int(26)] = d_905_v100_
            nw136_[int(27)] = d_905_v100_
            nw136_[int(28)] = d_905_v100_
            d_910_v103_ = nw136_
            d_911_v104_: D7
            d_911_v104_ = D7_DC22((self).f12, d_910_v103_)
            d_912_v105_: _dafny.Map
            d_912_v105_ = _dafny.Map({(self).f12: (D7_DC22(d_903_cf31_, d_910_v103_) if d_893_v93_ else d_911_v104_)})
            d_912_v105_ = (d_912_v105_).set((d_897_v97_)[default__.safeIndex(((d_894_v94_)[d_893_v93_] if (d_893_v93_) in (d_894_v94_) else (self).f12), len(d_897_v97_))], d_911_v104_)
            d_913_v106_: _dafny.Map
            d_913_v106_ = _dafny.Map({True: -733})
            d_914_v107_: str
            d_914_v107_ = _dafny.CodePoint('q')
            d_915_v108_: _dafny.Map
            d_915_v108_ = _dafny.Map({d_914_v107_: d_893_v93_})
            d_916_v109_: D7
            d_916_v109_ = D7_DC23(d_910_v103_, d_896_v96_, d_893_v93_, (self).f13)
            d_917_v110_: C2
            nw137_ = C2()
            nw137_.ctor__((len(d_913_v106_)) * (d_903_cf31_), (self).fm8((self).f13, d_915_v108_, _dafny.Set({(self).f12}), globalState), (d_916_v109_).cf45, 4, (d_896_v96_)[default__.safeIndex((self).f13, len(d_896_v96_))])
            d_917_v110_ = nw137_
        elif source14_.is_DC18:
            d_918___mcc_h8_ = source14_.cf32
            d_919_cf32_ = d_918___mcc_h8_
            d_920_v111_: _dafny.Set
            d_920_v111_ = _dafny.Set({d_919_cf32_, d_919_cf32_})
            d_920_v111_ = (d_920_v111_).intersection(d_920_v111_)
            d_919_cf32_ = (_dafny.SeqWithoutIsStrInference([d_893_v93_])) == (d_896_v96_)
            d_921_v112_: int
            d_921_v112_ = 158
            d_921_v112_ = d_921_v112_
            d_922_v113_: _dafny.Map
            d_922_v113_ = _dafny.Map({d_895_v95_: d_896_v96_})
            d_923_v115_: _dafny.Map
            d_923_v115_ = _dafny.Map({d_919_cf32_: (self).f12})
            d_924_v116_: str
            d_924_v116_ = _dafny.CodePoint('h')
            d_925_v117_: _dafny.MultiSet
            d_925_v117_ = _dafny.MultiSet([(d_895_v95_).set(default__.safeIndex(len(d_923_v115_), len(d_895_v95_)), d_924_v116_)])
            def iife105_():
                coll57_ = _dafny.Map()
                compr_57_: _dafny.Seq
                for compr_57_ in (d_925_v117_).Elements:
                    d_926_v114_: _dafny.Seq = compr_57_
                    if (d_926_v114_) in (d_925_v117_):
                        coll57_[d_926_v114_] = d_896_v96_
                return _dafny.Map(coll57_)
            d_922_v113_ = iife105_()
            
        elif source14_.is_DC19:
            d_927___mcc_h9_ = source14_.cf33
            d_928___mcc_h10_ = source14_.cf34
            d_929___mcc_h11_ = source14_.cf35
            d_930___mcc_h12_ = source14_.cf36
            d_931___mcc_h13_ = source14_.cf37
            d_932_cf37_ = d_931___mcc_h13_
            d_933_cf36_ = d_930___mcc_h12_
            d_934_cf35_ = d_929___mcc_h11_
            d_935_cf34_ = d_928___mcc_h10_
            d_936_cf33_ = d_927___mcc_h9_
            d_937_v118_: C0
            nw138_ = C0()
            nw138_.ctor__((636) - (402), d_893_v93_, (self).f13, default__.safeDivisionInt(788, 350))
            d_937_v118_ = nw138_
            d_938_v119_: int
            d_938_v119_ = 716
            d_938_v119_ = (self).f12
            d_939_v120_: _dafny.Array
            nw139_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_939_v120_ = nw139_
            d_940_v121_: D7
            d_940_v121_ = D7_DC22((d_937_v118_).f28, (d_939_v120_ if d_893_v93_ else d_939_v120_))
            rhs119_ = d_933_cf36_
            rhs120_ = d_897_v97_
            rhs121_ = d_940_v121_
            r0 = rhs119_
            d_897_v97_ = rhs120_
            d_940_v121_ = rhs121_
            d_941_v122_: _dafny.Array
            nw140_ = _dafny.Array(False, 20)
            d_941_v122_ = nw140_
            d_942_v123_: _dafny.Seq
            d_942_v123_ = _dafny.SeqWithoutIsStrInference([d_941_v122_, d_941_v122_])
            d_941_v122_ = (d_942_v123_)[default__.safeIndex(len(d_895_v95_), len(d_942_v123_))]
        elif source14_.is_DC16:
            d_943___mcc_h14_ = source14_.cf29
            d_944_cf29_ = d_943___mcc_h14_
            d_945_v124_: _dafny.Map
            d_945_v124_ = _dafny.Map({not(d_893_v93_): d_893_v93_})
            d_946_v125_: _dafny.MultiSet
            d_946_v125_ = _dafny.MultiSet([d_945_v124_, d_945_v124_])
            d_947_v126_: _dafny.MultiSet
            d_947_v126_ = d_946_v125_
            d_947_v126_ = d_947_v126_
            d_948_v127_: _dafny.Array
            nw141_ = _dafny.Array(int(0), 27)
            d_948_v127_ = nw141_
            index102_ = default__.safeIndex(76, (d_948_v127_).length(0))
            (d_948_v127_)[index102_] = (self).f13
            d_949_v128_: _dafny.MultiSet
            d_949_v128_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_893_v93_, d_893_v93_, d_893_v93_])), (self).f13])
            d_950_v129_: str
            d_950_v129_ = _dafny.CodePoint('d')
            d_951_v130_: _dafny.Set
            d_951_v130_ = _dafny.Set({d_950_v129_, d_950_v129_})
            d_952_v131_: _dafny.Map
            d_952_v131_ = _dafny.Map({d_951_v130_: len(d_896_v96_)})
            d_953_v132_: D4
            d_953_v132_ = D4_DC11(_dafny.Map({(self).f13: (self).f12}))
            index103_ = default__.safeIndex(76, (d_948_v127_).length(0))
            (d_948_v127_)[index103_] = ((d_949_v128_)[(584) - (len((d_952_v131_).set(d_951_v130_, len(_dafny.SeqWithoutIsStrInference([d_953_v132_, D4_DC11((d_767_v0_).set(((d_894_v94_)[d_893_v93_] if (d_893_v93_) in (d_894_v94_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajqx")))), (self).f13))])))))] if ((584) - (len((d_952_v131_).set(d_951_v130_, len(_dafny.SeqWithoutIsStrInference([d_953_v132_, D4_DC11((d_767_v0_).set(((d_894_v94_)[d_893_v93_] if (d_893_v93_) in (d_894_v94_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajqx")))), (self).f13))])))))) in (d_949_v128_) else default__.safeModuloInt((self).f13, (self).f12))
            (globalState).f8 = d_893_v93_
            index104_ = default__.safeIndex(76, (d_948_v127_).length(0))
            rhs122_ = d_897_v97_
            rhs123_ = (d_948_v127_)[default__.safeIndex(76, (d_948_v127_).length(0))]
            lhs71_ = globalState
            lhs72_ = d_948_v127_
            lhs73_ = default__.safeIndex(76, (d_948_v127_).length(0))
            lhs71_.f5 = rhs122_
            lhs72_[lhs73_] = rhs123_
        elif True:
            d_954___mcc_h15_ = source14_.cf38
            d_955_cf38_ = d_954___mcc_h15_
            d_956_v133_: C0
            nw142_ = C0()
            nw142_.ctor__((self).f13, default__.fm1((self).f13, d_893_v93_, (self).f13, globalState), 960, ((self).f13) - ((self).f13))
            d_956_v133_ = nw142_
            d_957_v134_: _dafny.Map
            d_957_v134_ = _dafny.Map({d_893_v93_: ((d_956_v133_).f28) + ((self).f13)})
            d_958_v135_: _dafny.Map
            d_958_v135_ = _dafny.Map({(d_956_v133_).f28: d_957_v134_})
            d_957_v134_ = (((d_958_v135_)[(d_956_v133_).f28] if ((d_956_v133_).f28) in (d_958_v135_) else d_957_v134_)).set(False, ((self).f12) + (default__.fm2(globalState)))
            (d_956_v133_).f29 = (d_893_v93_) == ((d_893_v93_) or (d_893_v93_))
            d_959_v136_: int
            d_959_v136_ = 758
            d_960_v137_: _dafny.Set
            d_960_v137_ = _dafny.Set({d_956_v133_.f29, d_956_v133_.f29, d_893_v93_})
            d_961_v138_: _dafny.Seq
            d_961_v138_ = _dafny.SeqWithoutIsStrInference([(d_960_v137_) - (_dafny.Set({False, d_893_v93_})), d_960_v137_])
            d_959_v136_ = len(d_961_v138_)
        r0 = d_893_v93_
        d_962_v139_: int
        d_962_v139_ = 974
        d_962_v139_ = (self).f12
        d_963_v140_: _dafny.Map
        d_963_v140_ = _dafny.Map({(_dafny.MultiSet(d_897_v97_)).set((self).f12, default__.abs((self).f13)): len(d_897_v97_)})
        d_962_v139_ = len(((d_897_v97_).set(default__.safeIndex(6, len(d_897_v97_)), (self).f13)) + (_dafny.SeqWithoutIsStrInference([d_962_v139_, (0) - (len(d_963_v140_)), (self).f13])))
        d_964_v141_: D6
        d_964_v141_ = D6_DC20(D6_DC17(d_893_v93_, (self).f12))
        d_965_v142_: str
        d_965_v142_ = _dafny.CodePoint('r')
        d_966_v143_: D0
        d_966_v143_ = D0_DC1((self).f12, d_895_v95_, (self).f12, True)
        d_967_v144_: _dafny.Seq
        d_967_v144_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_965_v142_]), (d_966_v143_).cf2])
        d_968_v145_: _dafny.Array
        def lambda65_(d_969_i8_):
            return (d_969_i8_) - ((self).f13)

        init34_ = lambda65_
        nw143_ = _dafny.Array(None, 24)
        for i0_34_ in range(nw143_.length(0)):
            nw143_[i0_34_] = init34_(i0_34_)
        d_968_v145_ = nw143_
        d_970_v146_: _dafny.Array
        nw144_ = _dafny.Array(None, 5)
        nw144_[int(0)] = d_968_v145_
        nw144_[int(1)] = d_968_v145_
        nw144_[int(2)] = d_968_v145_
        nw144_[int(3)] = d_968_v145_
        nw144_[int(4)] = d_968_v145_
        d_970_v146_ = nw144_
        d_971_v147_: D7
        d_971_v147_ = D7_DC22((self).f13, d_970_v146_)
        index105_ = default__.safeIndex(829, (d_968_v145_).length(0))
        (d_968_v145_)[index105_] = (d_897_v97_)[default__.safeIndex(len(_dafny.Map({d_893_v93_: d_971_v147_})), len(d_897_v97_))]
        index106_ = default__.safeIndex(829, (d_968_v145_).length(0))
        rhs124_ = not(d_893_v93_)
        rhs125_ = ((self).f13) < ((self).f12)
        rhs126_ = default__.fm43((0) - ((self).f13), (d_893_v93_) and (d_893_v93_), (self).f12, (d_895_v95_) + (d_895_v95_), globalState)
        rhs127_ = _dafny.SeqWithoutIsStrInference([d_895_v95_])
        rhs128_ = d_962_v139_
        lhs74_ = d_968_v145_
        lhs75_ = default__.safeIndex(829, (d_968_v145_).length(0))
        d_893_v93_ = rhs124_
        d_893_v93_ = rhs125_
        d_964_v141_ = rhs126_
        d_967_v144_ = rhs127_
        lhs74_[lhs75_] = rhs128_
        r0 = d_893_v93_
        r1 = d_968_v145_
        d_972_v148_: _dafny.Array
        nw145_ = _dafny.Array(_dafny.Map({}), 16)
        d_972_v148_ = nw145_
        r2 = d_972_v148_
        d_973_v149_: _dafny.Seq
        d_973_v149_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D7_DC21(d_894_v94_), D7_DC21(d_894_v94_), D7_DC21(d_894_v94_)])) for d_974_i9_ in range(default__.abs(898))])
        r3 = _dafny.Set({d_973_v149_})
        return r0, r1, r2, r3


class C4(T0):
    def  __init__(self):
        self._f11: bool = False
        self._f25: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    def ctor__(self, f25, f11):
        (self)._f25 = f25
        (self).f11 = f11

    def fm5(self, p0, globalState):
        return D0_DC1((self).f25, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsdsks")), (self).f25, self.f11)

    def fm6(self, p0, p1, p2, globalState):
        return ((self).f25) > ((self).f25)

    def fm28(self, p0, globalState):
        return not(self.f11)

    def fm29(self, p0, p1, globalState):
        if self.f11:
            return (not(self.f11)) and (self.f11)
        elif True:
            return self.f11

    def m21(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_975_v0_: _dafny.Seq
        d_975_v0_ = _dafny.SeqWithoutIsStrInference([False])
        d_976_v1_: _dafny.Map
        d_976_v1_ = _dafny.Map({(d_975_v0_) + (_dafny.SeqWithoutIsStrInference([True, self.f11])): p2})
        d_976_v1_ = (_dafny.Map({d_975_v0_: len(d_975_v0_)})) | ((d_976_v1_) | (d_976_v1_))
        d_977_v3_: _dafny.Map
        def iife106_():
            coll58_ = _dafny.Map()
            compr_58_: int
            for compr_58_ in _dafny.IntegerRange(665, 886):
                d_978_v2_: int = compr_58_
                if ((665) <= (d_978_v2_)) and ((d_978_v2_) < (886)):
                    coll58_[default__.safeModuloInt(d_978_v2_, p3)] = p3
            return _dafny.Map(coll58_)
        d_977_v3_ = _dafny.Map({p2: iife106_()
        })
        d_979_v4_: D3
        d_979_v4_ = D3_DC9((self).f25, d_977_v3_, (self).f25, (self).f25)
        d_980_v5_: D3
        d_980_v5_ = D3_DC10(d_979_v4_)
        d_981_i0_: int
        d_981_i0_ = 0
        with _dafny.label("3"):
            pat_let_tv24_ = p2
            pat_let_tv25_ = p0
            pat_let_tv26_ = p0
            def lambda66_(source15_):
                if source15_.is_DC9:
                    d_987___mcc_h0_ = source15_.cf15
                    d_988___mcc_h1_ = source15_.cf16
                    d_989___mcc_h2_ = source15_.cf17
                    d_990___mcc_h3_ = source15_.cf18
                    d_991_cf18_ = d_990___mcc_h3_
                    d_992_cf17_ = d_989___mcc_h2_
                    d_993_cf16_ = d_988___mcc_h1_
                    d_994_cf15_ = d_987___mcc_h0_
                    def iife107_():
                        coll59_ = _dafny.Set()
                        compr_59_: int
                        for compr_59_ in _dafny.IntegerRange(649, 528):
                            d_995_v6_: int = compr_59_
                            if ((649) <= (d_995_v6_)) and ((d_995_v6_) < (528)):
                                coll59_ = coll59_.union(_dafny.Set([(d_995_v6_) + (pat_let_tv24_)]))
                        return _dafny.Set(coll59_)
                    return (len(iife107_()
                    )) == ((self).f25)
                elif source15_.is_DC8:
                    d_996___mcc_h4_ = source15_.cf14
                    d_997_cf14_ = d_996___mcc_h4_
                    return self.f11
                elif True:
                    d_998___mcc_h5_ = source15_.cf19
                    d_999_cf19_ = d_998___mcc_h5_
                    return (pat_let_tv25_) != (pat_let_tv26_)

            while lambda66_(d_980_v5_):
                with _dafny.c_label("3"):
                    if (d_981_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_981_i0_ = (d_981_i0_) + (1)
                    d_982_v7_: D0
                    d_982_v7_ = D0_DC3(p2, self.f11)
                    r1 = ((d_982_v7_).cf7) * ((self).f25)
                    d_983_v8_: _dafny.Seq
                    d_983_v8_ = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])
                    r1 = (333) - (len(d_983_v8_))
                    d_984_v9_: _dafny.Array
                    nw146_ = _dafny.Array(_dafny.MultiSet({}), 18)
                    d_984_v9_ = nw146_
                    d_985_v10_: _dafny.Seq
                    d_985_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccejxpoh"))
                    d_986_v11_: _dafny.Seq
                    d_986_v11_ = _dafny.SeqWithoutIsStrInference([d_984_v9_])
                    rhs129_ = (0) - (default__.fm2(globalState))
                    rhs130_ = (d_986_v11_)[default__.safeIndex((len(d_985_v10_)) * (p3), len(d_986_v11_))]
                    rhs131_ = (p2) - (p3)
                    rhs132_ = default__.safeModuloInt(p2, default__.fm2(globalState))
                    rhs133_ = (p0) + (p0)
                    r0 = rhs129_
                    d_984_v9_ = rhs130_
                    r1 = rhs131_
                    r1 = rhs132_
                    d_985_v10_ = rhs133_
                    r1 = len(d_985_v10_)
                    pass
            pass
        d_1000_v12_: _dafny.Map
        d_1000_v12_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p3, p3, p2]): p3})
        d_1001_v13_: _dafny.MultiSet
        d_1001_v13_ = _dafny.MultiSet([self.f11, (self).fm6(self.f11, (self).f25, len(d_1000_v12_), globalState)])
        if not((default__.fm30(self.f11, not(self.f11), (d_1001_v13_).cardinality, globalState)).ispropersubset(default__.fm30(self.f11, False, (self).f25, globalState))):
            d_1002_v14_: _dafny.Map
            d_1002_v14_ = _dafny.Map({default__.fm2(globalState): 970})
            d_1003_v15_: _dafny.Map
            d_1003_v15_ = _dafny.Map({self.f11: self.f11})
            d_1004_v16_: _dafny.Array
            nw147_ = _dafny.Array(None, 14)
            nw147_[int(0)] = len(d_1002_v14_)
            nw147_[int(1)] = p3
            nw147_[int(2)] = (self).f25
            nw147_[int(3)] = p2
            nw147_[int(4)] = p3
            nw147_[int(5)] = p3
            nw147_[int(6)] = p2
            nw147_[int(7)] = -549
            nw147_[int(8)] = len(d_1002_v14_)
            nw147_[int(9)] = (0) - (len(d_1003_v15_))
            nw147_[int(10)] = p2
            nw147_[int(11)] = p2
            nw147_[int(12)] = -768
            nw147_[int(13)] = p2
            d_1004_v16_ = nw147_
            d_1005_v17_: _dafny.Array
            nw148_ = _dafny.Array(None, 7)
            nw148_[int(0)] = d_1004_v16_
            nw148_[int(1)] = d_1004_v16_
            nw148_[int(2)] = d_1004_v16_
            nw148_[int(3)] = d_1004_v16_
            nw148_[int(4)] = d_1004_v16_
            nw148_[int(5)] = d_1004_v16_
            nw148_[int(6)] = d_1004_v16_
            d_1005_v17_ = nw148_
            index107_ = default__.safeIndex(399, (d_1005_v17_).length(0))
            (d_1005_v17_)[index107_] = d_1004_v16_
            index108_ = default__.safeIndex(399, (d_1005_v17_).length(0))
            (d_1005_v17_)[index108_] = d_1004_v16_
            d_1006_v18_: _dafny.Seq
            d_1006_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
            d_1007_v19_: _dafny.Seq
            d_1007_v19_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1008_v20_: _dafny.Seq
            d_1008_v20_ = _dafny.SeqWithoutIsStrInference([(d_1007_v19_)[default__.safeIndex((D0_DC2(p3, d_1004_v16_)).cf5, len(d_1007_v19_))], d_1006_v18_, (d_1006_v18_).set(default__.safeIndex(p3, len(d_1006_v18_)), (p0)[default__.safeIndex(default__.fm2(globalState), len(p0))])])
            d_1006_v18_ = (d_1008_v20_)[default__.safeIndex(default__.safeDivisionInt((self).f25, p3), len(d_1008_v20_))]
            source16_ = D0_DC3(p2, self.f11)
            if source16_.is_DC1:
                d_1009___mcc_h6_ = source16_.cf1
                d_1010___mcc_h7_ = source16_.cf2
                d_1011___mcc_h8_ = source16_.cf3
                d_1012___mcc_h9_ = source16_.cf4
                d_1013_cf4_ = d_1012___mcc_h9_
                d_1014_cf3_ = d_1011___mcc_h8_
                d_1015_cf2_ = d_1010___mcc_h7_
                d_1016_cf1_ = d_1009___mcc_h6_
                d_1017_v21_: _dafny.Map
                d_1017_v21_ = _dafny.Map({self.f11: 167})
                d_1018_v22_: _dafny.Seq
                d_1018_v22_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmgibfqjo")))])
                rhs134_ = (((d_1017_v21_)[d_1013_cf4_] if (d_1013_cf4_) in (d_1017_v21_) else len(d_975_v0_))) + (len(p0))
                rhs135_ = ((d_1001_v13_).set(self.f11, default__.abs((d_1018_v22_)[default__.safeIndex(len(d_1015_cf2_), len(d_1018_v22_))]))).isdisjoint(d_1001_v13_)
                rhs136_ = (self.f11 if d_1013_cf4_ else self.f11)
                lhs76_ = globalState
                lhs77_ = globalState
                r0 = rhs134_
                lhs76_.f8 = rhs135_
                lhs77_.f8 = rhs136_
                d_1019_v23_: _dafny.Set
                d_1019_v23_ = _dafny.Set({d_1014_cf3_})
                d_1020_v24_: _dafny.Map
                d_1020_v24_ = _dafny.Map({d_1013_cf4_: default__.fm0(globalState)})
                d_1013_cf4_ = (len(d_1019_v23_)) != ((0) - (len(((d_1020_v24_)[((d_1003_v15_)[d_1013_cf4_] if (d_1013_cf4_) in (d_1003_v15_) else self.f11)] if (((d_1003_v15_)[d_1013_cf4_] if (d_1013_cf4_) in (d_1003_v15_) else self.f11)) in (d_1020_v24_) else d_1018_v22_))))
                d_1021_v25_: _dafny.Array
                def lambda67_(d_1022_v0_):
                    def lambda68_(d_1023_i1_):
                        return d_1022_v0_

                    return lambda68_

                init35_ = lambda67_(d_975_v0_)
                nw149_ = _dafny.Array(None, 25)
                for i0_35_ in range(nw149_.length(0)):
                    nw149_[i0_35_] = init35_(i0_35_)
                d_1021_v25_ = nw149_
                d_1021_v25_ = d_1021_v25_
                d_1013_cf4_ = self.f11
            elif source16_.is_DC2:
                d_1024___mcc_h10_ = source16_.cf5
                d_1025___mcc_h11_ = source16_.cf6
                d_1026_cf6_ = d_1025___mcc_h11_
                d_1027_cf5_ = d_1024___mcc_h10_
                d_1028_v26_: D3
                d_1028_v26_ = D3_DC9(p3, d_977_v3_, (self).f25, (self).f25)
                d_1028_v26_ = D3_DC9(p2, default__.fm31(globalState), p3, (0) - (d_1027_cf5_))
                d_1029_v27_: _dafny.Seq
                d_1029_v27_ = _dafny.SeqWithoutIsStrInference([d_1004_v16_])
                d_1030_v28_: D1
                d_1030_v28_ = D1_DC4(d_1029_v27_)
                d_1031_v29_: D1
                d_1031_v29_ = D1_DC6(d_1030_v28_)
                d_1031_v29_ = d_1031_v29_
                d_1027_cf5_ = ((d_1001_v13_)[((d_1003_v15_)[not(self.f11)] if (not(self.f11)) in (d_1003_v15_) else self.f11)] if (((d_1003_v15_)[not(self.f11)] if (not(self.f11)) in (d_1003_v15_) else self.f11)) in (d_1001_v13_) else default__.safeModuloInt(d_1027_cf5_, p2))
                d_1032_v30_: _dafny.Seq
                d_1032_v30_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1033_v31_: C2
                nw150_ = C2()
                nw150_.ctor__((p2 if self.f11 else (self).f25), p0, (0) - ((d_1032_v30_)[default__.safeIndex(p2, len(d_1032_v30_))]), 803, not(True))
                d_1033_v31_ = nw150_
            elif source16_.is_DC3:
                d_1034___mcc_h12_ = source16_.cf7
                d_1035___mcc_h13_ = source16_.cf8
                d_1036_cf8_ = d_1035___mcc_h13_
                d_1037_cf7_ = d_1034___mcc_h12_
                d_1002_v14_ = (d_1002_v14_).set(p2, p3)
                d_1038_v32_: _dafny.Array
                nw151_ = _dafny.Array(_dafny.Set({}), 25)
                d_1038_v32_ = nw151_
                d_1038_v32_ = d_1038_v32_
                d_1039_v33_: _dafny.Array
                d_1040_v34_: bool
                out8_: _dafny.Array
                out9_: bool
                out8_, out9_ = default__.m0(globalState)
                d_1039_v33_ = out8_
                d_1040_v34_ = out9_
                arr0_ = (d_1005_v17_)[default__.safeIndex(399, (d_1005_v17_).length(0))]
                index109_ = default__.safeIndex(322, ((d_1005_v17_)[default__.safeIndex(399, (d_1005_v17_).length(0))]).length(0))
                arr0_[index109_] = (0) - ((self).f25)
                arr1_ = (d_1005_v17_)[default__.safeIndex(399, (d_1005_v17_).length(0))]
                index110_ = default__.safeIndex(322, ((d_1005_v17_)[default__.safeIndex(399, (d_1005_v17_).length(0))]).length(0))
                arr1_[index110_] = (p2) * (d_1037_cf7_)
            elif True:
                d_1041___mcc_h14_ = source16_.cf0
                d_1042_cf0_ = d_1041___mcc_h14_
                d_1006_v18_ = p0
                d_1043_v35_: D0
                d_1043_v35_ = D0_DC1(p3, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1044_i2_ in range(default__.abs(761))]), p3, self.f11)
                d_1045_v36_: _dafny.Seq
                d_1045_v36_ = _dafny.SeqWithoutIsStrInference([p3, p2, p3, len(p0)])
                d_1046_v37_: D0
                d_1046_v37_ = D0_DC2(p2, d_1004_v16_)
                d_1047_v38_: T0
                nw152_ = C2()
                nw152_.ctor__((self).f25, ((d_1043_v35_).cf2) + (d_1006_v18_), default__.fm2(globalState), (self).f25, (len(d_1045_v36_)) < ((0) - (len(_dafny.Map({self.f11: d_1046_v37_})))))
                d_1047_v38_ = nw152_
                nw153_ = C2()
                nw153_.ctor__(p2, d_1006_v18_, ((self).f25 if self.f11 else 143), (self).f25, self.f11)
                d_1047_v38_ = nw153_
                d_1048_v39_: _dafny.Set
                d_1048_v39_ = _dafny.Set({d_1047_v38_.f11})
                d_1049_v40_: _dafny.MultiSet
                d_1049_v40_ = _dafny.MultiSet([d_1048_v39_, d_1048_v39_, d_1048_v39_, d_1048_v39_])
                d_1003_v15_ = (d_1003_v15_).set((d_1049_v40_).ispropersubset(d_1049_v40_), False)
                d_1050_v41_: _dafny.Array
                d_1051_v42_: bool
                out10_: _dafny.Array
                out11_: bool
                out10_, out11_ = default__.m0(globalState)
                d_1050_v41_ = out10_
                d_1051_v42_ = out11_
            d_1052_v43_: _dafny.Set
            d_1052_v43_ = _dafny.Set({p3})
            r0 = default__.safeDivisionInt(default__.safeDivisionInt((0) - (len(d_1052_v43_)), p3), len((p0) + (d_1006_v18_)))
            if not(self.f11):
                (globalState).f8 = not(not((self).fm6(not((p2) == (p3)), len(_dafny.Set({self.f11, False, self.f11, self.f11, self.f11})), (p3) + (764), globalState)))
                r0 = (self).f25
                (globalState).f8 = (self.f11) and ((self).fm29(self.f11, len(_dafny.SeqWithoutIsStrInference([len(p0)])), globalState))
                index111_ = default__.safeIndex(174, (p1).length(0))
                (p1)[index111_] = self.f11
                index112_ = default__.safeIndex(174, (p1).length(0))
                (p1)[index112_] = (not(not(not(self.f11))) if default__.fm1(p2, self.f11, 803, globalState) else (self.f11) == (self.f11))
                d_1053_v44_: _dafny.Map
                d_1053_v44_ = _dafny.Map({p2: (p1)[default__.safeIndex(174, (p1).length(0))]})
                d_1054_v45_: _dafny.Map
                d_1054_v45_ = _dafny.Map({d_1053_v44_: not(self.f11)})
                d_1055_v46_: _dafny.Array
                nw154_ = _dafny.Array(False, 5)
                d_1055_v46_ = nw154_
                d_1056_v47_: _dafny.Map
                d_1056_v47_ = _dafny.Map({(d_1054_v45_) | (_dafny.Map({_dafny.Map({(self).f25: (p1)[default__.safeIndex(174, (p1).length(0))]}): not((self).fm6(not(not(self.f11)), p2, p3, globalState))})): d_1055_v46_})
                d_1056_v47_ = (d_1056_v47_).set(d_1054_v45_, d_1055_v46_)
            elif True:
                d_1057_v48_: _dafny.Seq
                d_1057_v48_ = _dafny.SeqWithoutIsStrInference([p2])
                r1 = (d_1057_v48_)[default__.safeIndex(((self).f25) * ((self).f25), len(d_1057_v48_))]
                d_1058_v49_: _dafny.Seq
                d_1058_v49_ = _dafny.SeqWithoutIsStrInference([d_1004_v16_, (d_1005_v17_)[default__.safeIndex(399, (d_1005_v17_).length(0))], d_1004_v16_])
                d_1059_v50_: D1
                d_1059_v50_ = D1_DC4(_dafny.SeqWithoutIsStrInference([(d_1005_v17_)[default__.safeIndex(399, (d_1005_v17_).length(0))], d_1004_v16_]))
                pat_let_tv27_ = d_1058_v49_
                pat_let_tv28_ = d_1058_v49_
                d_1060_v51_: _dafny.Seq
                def iife109_(_pat_let25_0):
                    def iife110_(d_1061_dt__update__tmp_h0_):
                        def iife111_(_pat_let26_0):
                            def iife112_(d_1062_dt__update_hcf9_h0_):
                                return D1_DC4(d_1062_dt__update_hcf9_h0_)
                            return iife112_(_pat_let26_0)
                        return iife111_(pat_let_tv27_)
                    return iife110_(_pat_let25_0)
                def iife108_(_pat_let24_0):
                    def iife113_(d_1063_dt__update__tmp_h1_):
                        def iife114_(_pat_let27_0):
                            def iife115_(d_1064_dt__update_hcf9_h1_):
                                return D1_DC4(d_1064_dt__update_hcf9_h1_)
                            return iife115_(_pat_let27_0)
                        return iife114_(pat_let_tv28_)
                    return iife113_(_pat_let24_0)
                d_1060_v51_ = _dafny.SeqWithoutIsStrInference([D1_DC4(d_1058_v49_), iife108_(iife109_(d_1059_v50_))])
                d_1060_v51_ = d_1060_v51_
                (self).f11 = (default__.fm46(self.f11, (d_975_v0_)[default__.safeIndex(102, len(d_975_v0_))], globalState)).cf21
                d_1065_v52_: _dafny.Array
                nw155_ = _dafny.Array(_dafny.Map({}), 27)
                d_1065_v52_ = nw155_
                d_1066_v53_: _dafny.Array
                nw156_ = _dafny.Array(_dafny.Seq({}), 16)
                d_1066_v53_ = nw156_
                d_1067_v54_: D0
                d_1067_v54_ = D0_DC0(d_1066_v53_)
                d_1068_v55_: _dafny.Seq
                d_1068_v55_ = _dafny.SeqWithoutIsStrInference([d_1067_v54_, d_1067_v54_])
                index113_ = default__.safeIndex(434, (d_1065_v52_).length(0))
                (d_1065_v52_)[index113_] = _dafny.Map({d_1068_v55_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "va"))})
                d_1069_v56_: _dafny.Map
                d_1069_v56_ = _dafny.Map({d_1068_v55_: d_1006_v18_})
                index114_ = default__.safeIndex(434, (d_1065_v52_).length(0))
                (d_1065_v52_)[index114_] = (d_1069_v56_) | (_dafny.Map({d_1068_v55_: p0}))
                index115_ = default__.safeIndex(399, (d_1005_v17_).length(0))
                (d_1005_v17_)[index115_] = d_1004_v16_
        elif True:
            r1 = (self).f25
            d_1070_v57_: C2
            nw157_ = C2()
            nw157_.ctor__((0) - (len(default__.fm39(self.f11, self.f11, globalState))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njlurjslh")), 688, p2, not(self.f11))
            d_1070_v57_ = nw157_
            d_1071_v58_: _dafny.Map
            d_1071_v58_ = _dafny.Map({p2: self.f11})
            d_1072_v59_: str
            d_1072_v59_ = _dafny.CodePoint('d')
            rhs137_ = (_dafny.Map({p2: self.f11})) | ((d_1071_v58_) | (d_1071_v58_))
            rhs138_ = d_1072_v59_
            d_1071_v58_ = rhs137_
            d_1072_v59_ = rhs138_
            d_1073_v60_: _dafny.Set
            d_1073_v60_ = _dafny.Set({d_975_v0_})
            d_1074_v61_: C2
            nw158_ = C2()
            nw158_.ctor__((d_1070_v57_).f26, p0, ((d_1070_v57_).f26) * (p2), len((_dafny.Set({d_975_v0_})).intersection(d_1073_v60_)), self.f11)
            d_1074_v61_ = nw158_
            d_1071_v58_ = _dafny.Map({-739: self.f11})
        r1 = p3
        hi2_ = (self).f25
        for d_1075_i3_ in range((default__.fm2(globalState)) - (p3), hi2_):
            if True:
                d_1076_v62_: _dafny.Array
                nw159_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_1076_v62_ = nw159_
                d_1076_v62_ = d_1076_v62_
                (globalState).f8 = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1077_i4_ in range(default__.abs(148))]) if self.f11 else p0)) == ((p0 if self.f11 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1078_i5_ in range(default__.abs(83))])))
                d_1079_v63_: _dafny.Map
                d_1079_v63_ = _dafny.Map({p1: (self).f25})
                d_1080_v64_: _dafny.Set
                d_1080_v64_ = _dafny.Set({len(p0)})
                (globalState).f8 = (((d_1079_v63_)[p1] if (p1) in (d_1079_v63_) else len(d_1080_v64_))) < (-92)
                d_1081_v65_: str
                d_1081_v65_ = _dafny.CodePoint('a')
                d_1081_v65_ = (d_1081_v65_ if self.f11 else d_1081_v65_)
                d_1082_v66_: _dafny.Seq
                d_1082_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pq"))
                d_1082_v66_ = d_1082_v66_
            elif True:
                d_1083_v67_: _dafny.Array
                def lambda69_(d_1084_i3_):
                    def lambda70_(d_1085_i6_):
                        return (d_1085_i6_) + (d_1084_i3_)

                    return lambda70_

                init36_ = lambda69_(d_1075_i3_)
                nw160_ = _dafny.Array(None, 9)
                for i0_36_ in range(nw160_.length(0)):
                    nw160_[i0_36_] = init36_(i0_36_)
                d_1083_v67_ = nw160_
                index116_ = default__.safeIndex(104, (d_1083_v67_).length(0))
                (d_1083_v67_)[index116_] = p2
                index117_ = default__.safeIndex(104, (d_1083_v67_).length(0))
                (d_1083_v67_)[index117_] = p3
                index118_ = default__.safeIndex(77, (p1).length(0))
                (p1)[index118_] = self.f11
                index119_ = default__.safeIndex(77, (p1).length(0))
                (p1)[index119_] = (((d_1083_v67_)[default__.safeIndex(104, (d_1083_v67_).length(0))]) * (p2)) == (((0) - (d_1075_i3_)) - (d_1075_i3_))
                index120_ = default__.safeIndex(77, (p1).length(0))
                (p1)[index120_] = ((p1)[default__.safeIndex(77, (p1).length(0))]) and (self.f11)
                d_1086_v68_: _dafny.Array
                nw161_ = _dafny.Array(_dafny.CodePoint('D'), 4)
                d_1086_v68_ = nw161_
                d_1087_v69_: str
                d_1087_v69_ = _dafny.CodePoint('l')
                index121_ = default__.safeIndex(207, (d_1086_v68_).length(0))
                (d_1086_v68_)[index121_] = d_1087_v69_
                d_1088_v70_: _dafny.Array
                nw162_ = _dafny.Array(_dafny.Seq({}), 1)
                d_1088_v70_ = nw162_
                d_1089_v71_: _dafny.Seq
                d_1089_v71_ = _dafny.SeqWithoutIsStrInference([d_1075_i3_])
                d_1090_v72_: _dafny.MultiSet
                d_1090_v72_ = _dafny.MultiSet([531, p3])
                d_1091_v73_: _dafny.Array
                nw163_ = _dafny.Array(None, 6)
                nw163_[int(0)] = d_1089_v71_
                nw163_[int(1)] = d_1089_v71_
                nw163_[int(2)] = d_1089_v71_
                nw163_[int(3)] = d_1089_v71_
                nw163_[int(4)] = _dafny.SeqWithoutIsStrInference([p2, d_1075_i3_, len(_dafny.SeqWithoutIsStrInference([-349 for d_1092_i7_ in range(default__.abs(91))])), (d_1089_v71_)[default__.safeIndex((d_1090_v72_).cardinality, len(d_1089_v71_))]])
                nw163_[int(5)] = d_1089_v71_
                d_1091_v73_ = nw163_
                d_1093_v74_: D0
                d_1093_v74_ = D0_DC0(d_1091_v73_)
                d_1094_v75_: _dafny.Set
                d_1094_v75_ = _dafny.Set({d_1093_v74_})
                index122_ = default__.safeIndex(77, (p1).length(0))
                index123_ = default__.safeIndex(207, (d_1086_v68_).length(0))
                index124_ = default__.safeIndex(104, (d_1083_v67_).length(0))
                rhs139_ = (d_1094_v75_).issubset(d_1094_v75_)
                rhs140_ = d_1087_v69_
                rhs141_ = d_1088_v70_
                rhs142_ = (d_1083_v67_)[default__.safeIndex(104, (d_1083_v67_).length(0))]
                rhs143_ = (p0) != (p0)
                lhs78_ = p1
                lhs79_ = default__.safeIndex(77, (p1).length(0))
                lhs80_ = d_1086_v68_
                lhs81_ = default__.safeIndex(207, (d_1086_v68_).length(0))
                lhs82_ = d_1083_v67_
                lhs83_ = default__.safeIndex(104, (d_1083_v67_).length(0))
                lhs84_ = self
                lhs78_[lhs79_] = rhs139_
                lhs80_[lhs81_] = rhs140_
                d_1088_v70_ = rhs141_
                lhs82_[lhs83_] = rhs142_
                lhs84_.f11 = rhs143_
                d_1095_v76_: D0
                d_1095_v76_ = D0_DC2(d_1075_i3_, d_1083_v67_)
                d_1095_v76_ = d_1095_v76_
            d_1096_v77_: _dafny.Set
            d_1096_v77_ = _dafny.Set({self.f11})
            d_1097_v78_: _dafny.Array
            nw164_ = _dafny.Array(int(0), 28)
            d_1097_v78_ = nw164_
            index125_ = default__.safeIndex(729, (d_1097_v78_).length(0))
            (d_1097_v78_)[index125_] = (0) - (p2)
            d_1098_v79_: _dafny.Seq
            d_1098_v79_ = _dafny.SeqWithoutIsStrInference([len((p0) + (p0))])
            index126_ = default__.safeIndex(729, (d_1097_v78_).length(0))
            rhs144_ = ((d_1096_v77_) | (d_1096_v77_)).intersection(d_1096_v77_)
            rhs145_ = (d_1098_v79_)[default__.safeIndex(p3, len(d_1098_v79_))]
            rhs146_ = self.f11
            rhs147_ = (self).f25
            lhs85_ = d_1097_v78_
            lhs86_ = default__.safeIndex(729, (d_1097_v78_).length(0))
            lhs87_ = globalState
            d_1096_v77_ = rhs144_
            lhs85_[lhs86_] = rhs145_
            lhs87_.f8 = rhs146_
            r1 = rhs147_
            index127_ = default__.safeIndex(182, (p1).length(0))
            (p1)[index127_] = self.f11
            index128_ = default__.safeIndex(182, (p1).length(0))
            (p1)[index128_] = self.f11
            index129_ = default__.safeIndex(182, (p1).length(0))
            (p1)[index129_] = (p1)[default__.safeIndex(182, (p1).length(0))]
        d_1099_v80_: _dafny.MultiSet
        d_1099_v80_ = _dafny.MultiSet([p3, (self).f25, (self).f25])
        d_1100_v81_: str
        d_1100_v81_ = _dafny.CodePoint('e')
        d_1101_v82_: _dafny.Set
        d_1101_v82_ = _dafny.Set({self.f11})
        d_1102_v83_: _dafny.Seq
        d_1102_v83_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_1103_v84_: D5
        d_1103_v84_ = D5_DC15(len(d_1102_v83_), p2, p3)
        d_1104_v85_: _dafny.Map
        d_1104_v85_ = _dafny.Map({len(d_1101_v82_): d_1103_v84_})
        d_1105_v86_: D6
        d_1105_v86_ = D6_DC20(D6_DC16(d_1104_v85_))
        d_1106_v87_: D6
        d_1106_v87_ = D6_DC20(d_1105_v86_)
        d_1107_v88_: D6
        d_1107_v88_ = D6_DC20(d_1105_v86_)
        d_1108_v89_: _dafny.Map
        d_1108_v89_ = _dafny.Map({default__.fm2(globalState): p3})
        pat_let_tv29_ = d_1099_v80_
        pat_let_tv30_ = d_1099_v80_
        pat_let_tv31_ = d_1099_v80_
        pat_let_tv32_ = d_1099_v80_
        pat_let_tv33_ = d_1102_v83_
        pat_let_tv34_ = d_1102_v83_
        pat_let_tv35_ = d_975_v0_
        def lambda71_(source17_):
            if source17_.is_DC17:
                d_1109___mcc_h15_ = source17_.cf30
                d_1110___mcc_h16_ = source17_.cf31
                d_1111_cf31_ = d_1110___mcc_h16_
                d_1112_cf30_ = d_1109___mcc_h15_
                return pat_let_tv29_
            elif source17_.is_DC18:
                d_1113___mcc_h17_ = source17_.cf32
                d_1114_cf32_ = d_1113___mcc_h17_
                return (pat_let_tv30_).intersection(_dafny.MultiSet([699]))
            elif source17_.is_DC19:
                d_1115___mcc_h18_ = source17_.cf33
                d_1116___mcc_h19_ = source17_.cf34
                d_1117___mcc_h20_ = source17_.cf35
                d_1118___mcc_h21_ = source17_.cf36
                d_1119___mcc_h22_ = source17_.cf37
                d_1120_cf37_ = d_1119___mcc_h22_
                d_1121_cf36_ = d_1118___mcc_h21_
                d_1122_cf35_ = d_1117___mcc_h20_
                d_1123_cf34_ = d_1116___mcc_h19_
                d_1124_cf33_ = d_1115___mcc_h18_
                return pat_let_tv31_
            elif source17_.is_DC16:
                d_1125___mcc_h23_ = source17_.cf29
                d_1126_cf29_ = d_1125___mcc_h23_
                return pat_let_tv32_
            elif True:
                d_1127___mcc_h24_ = source17_.cf38
                d_1128_cf38_ = d_1127___mcc_h24_
                return _dafny.MultiSet((pat_let_tv33_) + (pat_let_tv34_))

        def iife116_(_pat_let28_0):
            def iife117_(d_1129_dt__update__tmp_h2_):
                def iife118_(_pat_let29_0):
                    def iife119_(d_1130_dt__update_hcf38_h0_):
                        return D6_DC20(d_1130_dt__update_hcf38_h0_)
                    return iife119_(_pat_let29_0)
                return iife118_(D6_DC19(self.f11, self.f11, self.f11, self.f11, pat_let_tv35_))
            return iife117_(_pat_let28_0)
        rhs148_ = p3
        rhs149_ = lambda71_((d_1107_v88_ if self.f11 else iife116_(d_1107_v88_)))
        rhs150_ = len(d_1108_v89_)
        rhs151_ = d_1100_v81_
        r1 = rhs148_
        d_1099_v80_ = rhs149_
        r0 = rhs150_
        d_1100_v81_ = rhs151_
        r0 = p3
        r1 = p3
        d_1131_v90_: _dafny.Seq
        d_1131_v90_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        r2 = (d_1131_v90_)[default__.safeIndex(default__.safeDivisionInt((d_1099_v80_).cardinality, len(d_975_v0_)), len(d_1131_v90_))]
        return r0, r1, r2

    @property
    def f25(self):
        return self._f25

class C5:
    def  __init__(self):
        self._f24: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f24):
        (self)._f24 = f24

    def fm26(self, p0, p1, p2, p3, globalState):
        def iife120_():
            coll60_ = _dafny.Set()
            compr_60_: int
            for compr_60_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), True]))).cardinality])) for d_1132_i0_ in range(default__.abs(966))])).Elements:
                d_1133_v0_: int = compr_60_
                if (d_1133_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), True]))).cardinality])) for d_1132_i0_ in range(default__.abs(966))])):
                    coll60_ = coll60_.union(_dafny.Set([default__.safeDivisionInt(d_1133_v0_, 11)]))
            return _dafny.Set(coll60_)
        def iife121_():
            coll61_ = _dafny.Set()
            compr_61_: int
            for compr_61_ in _dafny.IntegerRange(-629, 423):
                d_1134_v1_: int = compr_61_
                if ((-629) <= (d_1134_v1_)) and ((d_1134_v1_) < (423)):
                    coll61_ = coll61_.union(_dafny.Set([(d_1134_v1_) * ((self).f24)]))
            return _dafny.Set(coll61_)
        return (iife120_()
        ) | ((iife121_()
        ).intersection(_dafny.Set({(self).f24})))

    def fm27(self, p0, p1, p2, globalState):
        return (default__.safeModuloInt((self).f24, (self).f24)) > (((self).f24) - ((D0_DC1((self).f24, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgoabqmce")), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ir"))), True)).cf1))

    def m20(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1135_v0_: bool
        d_1135_v0_ = True
        d_1136_v1_: _dafny.Map
        d_1136_v1_ = _dafny.Map({d_1135_v0_: (self).f24})
        d_1137_v2_: D3
        d_1137_v2_ = D3_DC8(d_1136_v1_)
        source18_ = d_1137_v2_
        if source18_.is_DC9:
            d_1138___mcc_h0_ = source18_.cf15
            d_1139___mcc_h1_ = source18_.cf16
            d_1140___mcc_h2_ = source18_.cf17
            d_1141___mcc_h3_ = source18_.cf18
            d_1142_cf18_ = d_1141___mcc_h3_
            d_1143_cf17_ = d_1140___mcc_h2_
            d_1144_cf16_ = d_1139___mcc_h1_
            d_1145_cf15_ = d_1138___mcc_h0_
            (globalState).f8 = (default__.fm2(globalState)) >= (default__.fm2(globalState))
            d_1146_v3_: _dafny.Seq
            d_1146_v3_ = _dafny.SeqWithoutIsStrInference([not(d_1135_v0_), d_1135_v0_])
            d_1147_v4_: _dafny.Seq
            d_1147_v4_ = _dafny.SeqWithoutIsStrInference([d_1146_v3_])
            d_1148_v5_: C0
            nw165_ = C0()
            nw165_.ctor__(d_1145_cf15_, d_1135_v0_, (_dafny.MultiSet(d_1147_v4_)).cardinality, d_1142_cf18_)
            d_1148_v5_ = nw165_
            d_1149_v6_: _dafny.Array
            nw166_ = _dafny.Array(_dafny.MultiSet({}), 25)
            d_1149_v6_ = nw166_
            d_1150_v8_: _dafny.MultiSet
            def iife122_():
                coll62_ = _dafny.Map()
                compr_62_: int
                for compr_62_ in _dafny.IntegerRange(326, 901):
                    d_1151_v7_: int = compr_62_
                    if ((326) <= (d_1151_v7_)) and ((d_1151_v7_) < (901)):
                        coll62_[default__.safeModuloInt(d_1151_v7_, (d_1148_v5_).f28)] = d_1148_v5_.f29
                return _dafny.Map(coll62_)
            d_1150_v8_ = _dafny.MultiSet([len(iife122_()
            )])
            index130_ = default__.safeIndex(332, (d_1149_v6_).length(0))
            (d_1149_v6_)[index130_] = d_1150_v8_
            d_1152_v9_: _dafny.Array
            nw167_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_1152_v9_ = nw167_
            d_1153_v10_: D7
            d_1153_v10_ = D7_DC23(d_1152_v9_, d_1146_v3_, False, -32)
            d_1154_v11_: _dafny.Seq
            d_1154_v11_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            index131_ = default__.safeIndex(332, (d_1149_v6_).length(0))
            rhs152_ = (d_1150_v8_) - ((_dafny.MultiSet(d_1154_v11_)).set(len(d_1146_v3_), default__.abs(d_1145_cf15_)))
            rhs153_ = d_1142_cf18_
            rhs154_ = d_1153_v10_
            rhs155_ = (d_1148_v5_.f29) and (d_1135_v0_)
            lhs88_ = d_1149_v6_
            lhs89_ = default__.safeIndex(332, (d_1149_v6_).length(0))
            lhs90_ = globalState
            lhs88_[lhs89_] = rhs152_
            r0 = rhs153_
            d_1153_v10_ = rhs154_
            lhs90_.f8 = rhs155_
            d_1153_v10_ = d_1153_v10_
        elif source18_.is_DC8:
            d_1155___mcc_h4_ = source18_.cf14
            d_1156_cf14_ = d_1155___mcc_h4_
            if d_1135_v0_:
                d_1157_v12_: _dafny.Array
                nw168_ = _dafny.Array(None, 7)
                nw168_[int(0)] = not(((self).f24) != ((self).f24))
                nw168_[int(1)] = d_1135_v0_
                nw168_[int(2)] = d_1135_v0_
                nw168_[int(3)] = d_1135_v0_
                nw168_[int(4)] = d_1135_v0_
                nw168_[int(5)] = not(((self).f24) <= ((self).f24))
                nw168_[int(6)] = d_1135_v0_
                d_1157_v12_ = nw168_
                d_1158_v13_: _dafny.Array
                nw169_ = _dafny.Array(int(0), 7)
                d_1158_v13_ = nw169_
                index132_ = default__.safeIndex(452, (d_1158_v13_).length(0))
                (d_1158_v13_)[index132_] = (self).f24
                d_1159_v14_: _dafny.Seq
                d_1159_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihtnsyxyb"))
                index133_ = default__.safeIndex(452, (d_1158_v13_).length(0))
                rhs156_ = d_1157_v12_
                rhs157_ = d_1135_v0_
                rhs158_ = (self).f24
                rhs159_ = ((self).f24) - (len(d_1159_v14_))
                lhs91_ = globalState
                lhs92_ = d_1158_v13_
                lhs93_ = default__.safeIndex(452, (d_1158_v13_).length(0))
                d_1157_v12_ = rhs156_
                lhs91_.f8 = rhs157_
                lhs92_[lhs93_] = rhs158_
                r0 = rhs159_
                d_1160_v15_: _dafny.Map
                d_1160_v15_ = _dafny.Map({False: d_1135_v0_})
                d_1160_v15_ = (d_1160_v15_).set(not (d_1135_v0_) or (d_1135_v0_), d_1135_v0_)
                d_1161_v16_: _dafny.Seq
                d_1161_v16_ = _dafny.SeqWithoutIsStrInference([not(d_1135_v0_)])
                d_1162_v17_: D6
                d_1162_v17_ = D6_DC19(d_1135_v0_, d_1135_v0_, d_1135_v0_, d_1135_v0_, d_1161_v16_)
                d_1163_v18_: D6
                d_1163_v18_ = D6_DC20(d_1162_v17_)
                rhs160_ = default__.safeDivisionInt((self).f24, (d_1158_v13_)[default__.safeIndex(452, (d_1158_v13_).length(0))])
                rhs161_ = d_1163_v18_
                rhs162_ = len(d_1160_v15_)
                r0 = rhs160_
                d_1163_v18_ = rhs161_
                r0 = rhs162_
                d_1164_v19_: _dafny.MultiSet
                d_1164_v19_ = _dafny.MultiSet([d_1135_v0_])
                d_1165_v20_: _dafny.Map
                d_1165_v20_ = _dafny.Map({d_1164_v19_: d_1135_v0_})
                d_1166_v21_: _dafny.Map
                d_1166_v21_ = _dafny.Map({d_1165_v20_: d_1135_v0_})
                d_1167_v23_: _dafny.Seq
                d_1167_v23_ = _dafny.SeqWithoutIsStrInference([d_1164_v19_, d_1164_v19_])
                def iife123_():
                    coll63_ = _dafny.Map()
                    compr_63_: _dafny.MultiSet
                    for compr_63_ in (d_1167_v23_).Elements:
                        d_1168_v22_: _dafny.MultiSet = compr_63_
                        if (d_1168_v22_) in (d_1167_v23_):
                            coll63_[d_1168_v22_] = d_1135_v0_
                    return _dafny.Map(coll63_)
                d_1166_v21_ = (d_1166_v21_).set(iife123_()
                , d_1135_v0_)
                d_1169_v24_: _dafny.Array
                nw170_ = _dafny.Array(None, 6)
                nw170_[int(0)] = d_1158_v13_
                nw170_[int(1)] = d_1158_v13_
                nw170_[int(2)] = d_1158_v13_
                nw170_[int(3)] = d_1158_v13_
                nw170_[int(4)] = d_1158_v13_
                nw170_[int(5)] = d_1158_v13_
                d_1169_v24_ = nw170_
                d_1170_v25_: D7
                d_1170_v25_ = D7_DC24(D7_DC22((d_1158_v13_)[default__.safeIndex(452, (d_1158_v13_).length(0))], d_1169_v24_))
                d_1171_v26_: _dafny.Map
                d_1171_v26_ = _dafny.Map({(self).f24: d_1169_v24_})
                d_1172_v27_: _dafny.Seq
                d_1172_v27_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                pat_let_tv36_ = d_1171_v26_
                pat_let_tv37_ = d_1172_v27_
                pat_let_tv38_ = d_1158_v13_
                pat_let_tv39_ = d_1158_v13_
                pat_let_tv40_ = d_1172_v27_
                pat_let_tv41_ = d_1135_v0_
                pat_let_tv42_ = d_1172_v27_
                pat_let_tv43_ = d_1158_v13_
                pat_let_tv44_ = d_1158_v13_
                pat_let_tv45_ = d_1172_v27_
                pat_let_tv46_ = d_1135_v0_
                pat_let_tv47_ = d_1171_v26_
                pat_let_tv48_ = d_1169_v24_
                pat_let_tv49_ = d_1161_v16_
                pat_let_tv50_ = d_1135_v0_
                pat_let_tv51_ = d_1159_v14_
                def iife124_(_pat_let30_0):
                    def iife125_(d_1173_dt__update__tmp_h0_):
                        def iife126_(_pat_let31_0):
                            def iife127_(d_1174_dt__update_hcf46_h0_):
                                return D7_DC24(d_1174_dt__update_hcf46_h0_)
                            return iife127_(_pat_let31_0)
                        return iife126_(D7_DC23(((pat_let_tv36_)[len(_dafny.Map({(pat_let_tv37_)[default__.safeIndex((pat_let_tv39_)[default__.safeIndex(452, (pat_let_tv38_).length(0))], len(pat_let_tv40_))]: not(pat_let_tv41_)}))] if (len(_dafny.Map({(pat_let_tv42_)[default__.safeIndex((pat_let_tv44_)[default__.safeIndex(452, (pat_let_tv43_).length(0))], len(pat_let_tv45_))]: not(pat_let_tv46_)}))) in (pat_let_tv47_) else pat_let_tv48_), pat_let_tv49_, pat_let_tv50_, len(pat_let_tv51_)))
                    return iife125_(_pat_let30_0)
                d_1170_v25_ = iife124_(D7_DC24(D7_DC21(d_1164_v19_)))
            elif True:
                d_1175_v28_: _dafny.Set
                d_1175_v28_ = _dafny.Set({d_1135_v0_, True, d_1135_v0_})
                d_1175_v28_ = d_1175_v28_
                d_1176_v29_: _dafny.Array
                nw171_ = _dafny.Array(_dafny.Seq({}), 16)
                d_1176_v29_ = nw171_
                d_1177_v30_: _dafny.Map
                d_1177_v30_ = _dafny.Map({d_1176_v29_: (d_1135_v0_) and (d_1135_v0_)})
                d_1177_v30_ = (d_1177_v30_).set(d_1176_v29_, d_1135_v0_)
                d_1178_v31_: _dafny.Seq
                d_1178_v31_ = _dafny.SeqWithoutIsStrInference([d_1135_v0_])
                d_1179_v32_: _dafny.Map
                d_1179_v32_ = _dafny.Map({D0_DC3((_dafny.MultiSet(d_1178_v31_)).cardinality, d_1135_v0_): _dafny.Set({(self).f24})})
                d_1180_v33_: D0
                d_1180_v33_ = D0_DC3((self).f24, d_1135_v0_)
                d_1181_v34_: _dafny.Set
                d_1181_v34_ = _dafny.Set({(self).f24})
                d_1182_v35_: _dafny.Set
                d_1182_v35_ = _dafny.Set({len(d_1181_v34_)})
                d_1179_v32_ = (d_1179_v32_) | ((d_1179_v32_).set(d_1180_v33_, d_1182_v35_))
                d_1183_v36_: _dafny.MultiSet
                d_1183_v36_ = _dafny.MultiSet([(self).f24, (self).f24])
                d_1184_v37_: _dafny.Seq
                d_1184_v37_ = _dafny.SeqWithoutIsStrInference([258])
                d_1185_v38_: _dafny.Seq
                d_1185_v38_ = _dafny.SeqWithoutIsStrInference([(d_1184_v37_)[default__.safeIndex(((d_1183_v36_)[473] if (473) in (d_1183_v36_) else (self).f24), len(d_1184_v37_))]])
                d_1186_v39_: D1
                d_1186_v39_ = D1_DC5(((d_1183_v36_).set((self).f24, default__.abs((self).f24))) - (_dafny.MultiSet(d_1184_v37_)), d_1135_v0_)
                pat_let_tv52_ = d_1135_v0_
                pat_let_tv53_ = d_1135_v0_
                def iife128_(_pat_let32_0):
                    def iife129_(d_1187_dt__update__tmp_h1_):
                        def iife130_(_pat_let33_0):
                            def iife131_(d_1188_dt__update_hcf11_h0_):
                                return D1_DC5((d_1187_dt__update__tmp_h1_).cf10, d_1188_dt__update_hcf11_h0_)
                            return iife131_(_pat_let33_0)
                        return iife130_((pat_let_tv52_) or (pat_let_tv53_))
                    return iife129_(_pat_let32_0)
                d_1186_v39_ = iife128_(d_1186_v39_)
                (globalState).f8 = (d_1183_v36_).issubset(d_1183_v36_)
            d_1189_v40_: _dafny.Seq
            d_1189_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvfyoivjb"))
            (globalState).f8 = not ((d_1189_v40_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llh")))) or (d_1135_v0_)
            (globalState).f8 = ((self).f24) == (-218)
            d_1190_v41_: C2
            nw172_ = C2()
            nw172_.ctor__((self).f24, d_1189_v40_, (self).f24, (self).f24, d_1135_v0_)
            d_1190_v41_ = nw172_
        elif True:
            d_1191___mcc_h5_ = source18_.cf19
            d_1192_cf19_ = d_1191___mcc_h5_
            r0 = (self).f24
            d_1193_v42_: _dafny.Array
            def lambda72_(d_1194_i0_):
                return False

            init37_ = lambda72_
            nw173_ = _dafny.Array(None, 24)
            for i0_37_ in range(nw173_.length(0)):
                nw173_[i0_37_] = init37_(i0_37_)
            d_1193_v42_ = nw173_
            index134_ = default__.safeIndex(286, (d_1193_v42_).length(0))
            (d_1193_v42_)[index134_] = d_1135_v0_
            index135_ = default__.safeIndex(286, (d_1193_v42_).length(0))
            (d_1193_v42_)[index135_] = d_1135_v0_
            d_1195_v43_: _dafny.Array
            def lambda73_(d_1196_i1_):
                return D5_DC15((self).f24, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1197_i2_ in range(default__.abs(272))])).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1198_i2_ in range(default__.abs(272))]))), _dafny.CodePoint('b'))), (self).f24)

            init38_ = lambda73_
            nw174_ = _dafny.Array(None, 1)
            for i0_38_ in range(nw174_.length(0)):
                nw174_[i0_38_] = init38_(i0_38_)
            d_1195_v43_ = nw174_
            d_1199_v44_: _dafny.MultiSet
            d_1199_v44_ = _dafny.MultiSet([d_1195_v43_])
            d_1200_v45_: _dafny.Map
            d_1200_v45_ = _dafny.Map({(self).f24: (d_1199_v44_) | (d_1199_v44_)})
            d_1201_v46_: _dafny.Set
            d_1201_v46_ = _dafny.Set({(self).f24})
            d_1202_v47_: _dafny.Map
            d_1202_v47_ = _dafny.Map({(self).f24: len(d_1201_v46_)})
            d_1200_v45_ = (d_1200_v45_).set(((d_1202_v47_)[-360] if (-360) in (d_1202_v47_) else (self).f24), d_1199_v44_)
            d_1203_v48_: D5
            d_1203_v48_ = D5_DC15(len(_dafny.Map({default__.fm1(len(_dafny.SeqWithoutIsStrInference([(d_1193_v42_)[default__.safeIndex(286, (d_1193_v42_).length(0))], d_1135_v0_, d_1135_v0_, d_1135_v0_, (d_1193_v42_)[default__.safeIndex(286, (d_1193_v42_).length(0))]])), (d_1193_v42_)[default__.safeIndex(286, (d_1193_v42_).length(0))], (self).f24, globalState): len(_dafny.SeqWithoutIsStrInference([(self).f24]))})), (self).f24, (self).f24)
            d_1204_v49_: _dafny.Map
            d_1204_v49_ = _dafny.Map({(self).f24: d_1203_v48_})
            d_1205_v50_: D6
            d_1205_v50_ = D6_DC16(d_1204_v49_)
            d_1206_v51_: D6
            d_1206_v51_ = D6_DC20(d_1205_v50_)
            d_1207_v52_: _dafny.Seq
            d_1207_v52_ = _dafny.SeqWithoutIsStrInference([D6_DC20(d_1205_v50_), d_1206_v51_, d_1206_v51_])
            r0 = len(d_1207_v52_)
        d_1208_v53_: _dafny.Seq
        d_1208_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atdgxf"))
        d_1209_v54_: D0
        d_1209_v54_ = D0_DC1((self).f24, d_1208_v53_, (self).f24, not(d_1135_v0_))
        if (d_1209_v54_).cf4:
            d_1210_v55_: _dafny.Map
            d_1210_v55_ = _dafny.Map({d_1135_v0_: d_1135_v0_})
            r0 = len((((d_1210_v55_).set(d_1135_v0_, d_1135_v0_)) | (d_1210_v55_)) | (d_1210_v55_))
            d_1211_v56_: str
            d_1211_v56_ = _dafny.CodePoint('r')
            r1 = len(((d_1208_v53_).set(default__.safeIndex((self).f24, len(d_1208_v53_)), d_1211_v56_)).set(default__.safeIndex((0) - (((0) - (default__.fm2(globalState))) - ((self).f24)), len((d_1208_v53_).set(default__.safeIndex((self).f24, len(d_1208_v53_)), d_1211_v56_))), d_1211_v56_))
            d_1212_v57_: D9
            d_1212_v57_ = D9_DC26(d_1211_v56_)
            (globalState).f8 = not(((d_1212_v57_).cf48) in (d_1208_v53_))
            d_1213_v58_: _dafny.MultiSet
            d_1213_v58_ = _dafny.MultiSet([d_1211_v56_, d_1211_v56_])
            d_1213_v58_ = (d_1213_v58_).intersection(d_1213_v58_)
            d_1135_v0_ = True
        elif True:
            d_1214_v59_: _dafny.Seq
            d_1214_v59_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            d_1215_v60_: D3
            d_1215_v60_ = D3_DC8(d_1136_v1_)
            d_1216_v61_: _dafny.Seq
            d_1216_v61_ = _dafny.SeqWithoutIsStrInference([D3_DC8(d_1136_v1_), d_1215_v60_])
            d_1217_v62_: D3
            d_1217_v62_ = D3_DC10((d_1216_v61_)[default__.safeIndex(len(d_1214_v59_), len(d_1216_v61_))])
            d_1218_v63_: _dafny.Map
            d_1218_v63_ = _dafny.Map({d_1217_v62_: d_1136_v1_})
            d_1219_v64_: _dafny.Array
            nw175_ = _dafny.Array(None, 20)
            nw175_[int(0)] = d_1136_v1_
            nw175_[int(1)] = d_1136_v1_
            nw175_[int(2)] = d_1136_v1_
            nw175_[int(3)] = d_1136_v1_
            nw175_[int(4)] = d_1136_v1_
            nw175_[int(5)] = (d_1136_v1_) | (d_1136_v1_)
            nw175_[int(6)] = d_1136_v1_
            nw175_[int(7)] = d_1136_v1_
            nw175_[int(8)] = d_1136_v1_
            nw175_[int(9)] = d_1136_v1_
            nw175_[int(10)] = d_1136_v1_
            nw175_[int(11)] = _dafny.Map({d_1135_v0_: (self).f24})
            nw175_[int(12)] = d_1136_v1_
            nw175_[int(13)] = _dafny.Map({d_1135_v0_: len(d_1214_v59_)})
            nw175_[int(14)] = d_1136_v1_
            nw175_[int(15)] = d_1136_v1_
            nw175_[int(16)] = ((d_1218_v63_)[d_1217_v62_] if (d_1217_v62_) in (d_1218_v63_) else d_1136_v1_)
            nw175_[int(17)] = d_1136_v1_
            nw175_[int(18)] = d_1136_v1_
            nw175_[int(19)] = (d_1136_v1_) | (d_1136_v1_)
            d_1219_v64_ = nw175_
            index136_ = default__.safeIndex(521, (d_1219_v64_).length(0))
            (d_1219_v64_)[index136_] = d_1136_v1_
            index137_ = default__.safeIndex(521, (d_1219_v64_).length(0))
            (d_1219_v64_)[index137_] = d_1136_v1_
            d_1135_v0_ = d_1135_v0_
            d_1220_v65_: _dafny.Array
            nw176_ = _dafny.Array(False, 7)
            d_1220_v65_ = nw176_
            index138_ = default__.safeIndex(933, (d_1220_v65_).length(0))
            (d_1220_v65_)[index138_] = d_1135_v0_
            index139_ = default__.safeIndex(933, (d_1220_v65_).length(0))
            (d_1220_v65_)[index139_] = d_1135_v0_
            (globalState).f5 = (_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, default__.fm2(globalState), (self).f24, (self).f24])) + (d_1214_v59_)
            index140_ = default__.safeIndex(933, (d_1220_v65_).length(0))
            (d_1220_v65_)[index140_] = (d_1135_v0_) and ((d_1135_v0_ if d_1135_v0_ else (d_1220_v65_)[default__.safeIndex(933, (d_1220_v65_).length(0))]))
        d_1221_v66_: _dafny.Seq
        d_1221_v66_ = _dafny.SeqWithoutIsStrInference([d_1135_v0_])
        d_1222_v67_: _dafny.Set
        d_1222_v67_ = _dafny.Set({d_1135_v0_, d_1135_v0_})
        d_1223_v68_: _dafny.Map
        d_1223_v68_ = _dafny.Map({_dafny.Set({(self).f24, len(d_1221_v66_), len(d_1222_v67_)}): (d_1208_v53_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omygfqai")))})
        d_1224_v69_: _dafny.Set
        d_1224_v69_ = _dafny.Set({-491, (self).f24, (self).f24})
        d_1225_v70_: _dafny.MultiSet
        d_1225_v70_ = _dafny.MultiSet([(self).f24, (self).f24, (self).f24, (self).f24])
        d_1226_v71_: str
        d_1226_v71_ = _dafny.CodePoint('j')
        d_1223_v68_ = (d_1223_v68_).set(d_1224_v69_, (d_1208_v53_).set(default__.safeIndex(((d_1225_v70_)[483] if (483) in (d_1225_v70_) else (self).f24), len(d_1208_v53_)), d_1226_v71_))
        hi3_ = ((self).f24) - (179)
        for d_1227_i3_ in range((self).f24, hi3_):
            d_1228_v72_: _dafny.Array
            nw177_ = _dafny.Array(_dafny.CodePoint('D'), 20)
            d_1228_v72_ = nw177_
            d_1228_v72_ = d_1228_v72_
            r1 = default__.safeModuloInt((self).f24, (self).f24)
            d_1229_v73_: _dafny.Array
            def lambda74_(d_1230_i4_):
                return False

            init39_ = lambda74_
            nw178_ = _dafny.Array(None, 6)
            for i0_39_ in range(nw178_.length(0)):
                nw178_[i0_39_] = init39_(i0_39_)
            d_1229_v73_ = nw178_
            d_1231_v74_: _dafny.Array
            nw179_ = _dafny.Array(None, 13)
            nw179_[int(0)] = d_1229_v73_
            nw179_[int(1)] = d_1229_v73_
            nw179_[int(2)] = d_1229_v73_
            nw179_[int(3)] = d_1229_v73_
            nw179_[int(4)] = d_1229_v73_
            nw179_[int(5)] = d_1229_v73_
            nw179_[int(6)] = d_1229_v73_
            nw179_[int(7)] = d_1229_v73_
            nw179_[int(8)] = d_1229_v73_
            nw179_[int(9)] = d_1229_v73_
            nw179_[int(10)] = d_1229_v73_
            nw179_[int(11)] = d_1229_v73_
            nw179_[int(12)] = d_1229_v73_
            d_1231_v74_ = nw179_
            d_1232_v75_: _dafny.Map
            d_1232_v75_ = _dafny.Map({d_1224_v69_: d_1231_v74_})
            d_1232_v75_ = (d_1232_v75_).set(_dafny.Set({(self).f24, (self).f24, d_1227_i3_}), d_1231_v74_)
            d_1233_v76_: _dafny.Set
            d_1233_v76_ = _dafny.Set({default__.fm47(d_1135_v0_, d_1135_v0_, d_1136_v1_, d_1227_i3_, globalState)})
            d_1234_v77_: _dafny.Array
            nw180_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_1234_v77_ = nw180_
            pat_let_tv54_ = d_1234_v77_
            d_1235_v78_: D9
            def iife132_(_pat_let34_0):
                def iife133_(d_1236_dt__update__tmp_h2_):
                    def iife134_(_pat_let35_0):
                        def iife135_(d_1237_dt__update_hcf41_h0_):
                            return D7_DC22((d_1236_dt__update__tmp_h2_).cf40, d_1237_dt__update_hcf41_h0_)
                        return iife135_(_pat_let35_0)
                    return iife134_(pat_let_tv54_)
                return iife133_(_pat_let34_0)
            d_1235_v78_ = D9_DC27(d_1135_v0_, iife132_(D7_DC22(len(_dafny.SeqWithoutIsStrInference([(self).f24, len(d_1233_v76_), (self).f24])), d_1234_v77_)), (self).f24, d_1135_v0_)
            (globalState).f8 = (self).fm27((d_1235_v78_).cf51, d_1208_v53_, (-919) == (len(d_1221_v66_)), globalState)
        d_1238_v79_: _dafny.Map
        d_1238_v79_ = _dafny.Map({not(True): d_1135_v0_})
        d_1239_v80_: D6
        d_1239_v80_ = D6_DC18(d_1135_v0_)
        d_1240_v81_: D6
        d_1240_v81_ = D6_DC19(d_1135_v0_, ((d_1238_v79_)[d_1135_v0_] if (d_1135_v0_) in (d_1238_v79_) else (d_1239_v80_).cf32), not(default__.fm1((self).f24, d_1135_v0_, (self).f24, globalState)), d_1135_v0_, _dafny.SeqWithoutIsStrInference([d_1135_v0_, d_1135_v0_, d_1135_v0_, d_1135_v0_, d_1135_v0_]))
        rhs163_ = (d_1135_v0_ if (self).fm27((self).f24, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldu")), d_1135_v0_, globalState) else d_1135_v0_)
        rhs164_ = D6_DC19(d_1135_v0_, not(True), d_1135_v0_, not(not((d_1135_v0_) == (d_1135_v0_))), d_1221_v66_)
        rhs165_ = d_1208_v53_
        rhs166_ = ((self).f24) + ((self).f24)
        rhs167_ = (355) <= ((self).f24)
        d_1135_v0_ = rhs163_
        d_1240_v81_ = rhs164_
        d_1208_v53_ = rhs165_
        r1 = rhs166_
        d_1135_v0_ = rhs167_
        d_1241_v83_: D3
        def iife136_():
            coll64_ = _dafny.Map()
            compr_64_: int
            for compr_64_ in _dafny.IntegerRange(-620, 554):
                d_1242_v82_: int = compr_64_
                if ((-620) <= (d_1242_v82_)) and ((d_1242_v82_) < (554)):
                    coll64_[(d_1242_v82_) * ((self).f24)] = _dafny.Map({(self).f24: (self).f24})
            return _dafny.Map(coll64_)
        d_1241_v83_ = D3_DC9((self).f24, iife136_()
, len(d_1221_v66_), default__.fm2(globalState))
        d_1243_v84_: D3
        d_1243_v84_ = D3_DC10(d_1241_v83_)
        d_1244_v85_: D3
        d_1244_v85_ = D3_DC10(d_1241_v83_)
        pat_let_tv55_ = d_1243_v84_
        def iife137_(_pat_let36_0):
            def iife138_(d_1245_dt__update__tmp_h3_):
                def iife139_(_pat_let37_0):
                    def iife140_(d_1246_dt__update_hcf19_h0_):
                        return D3_DC10(d_1246_dt__update_hcf19_h0_)
                    return iife140_(_pat_let37_0)
                return iife139_(pat_let_tv55_)
            return iife138_(_pat_let36_0)
        source19_ = iife137_(d_1244_v85_)
        if source19_.is_DC9:
            d_1247___mcc_h6_ = source19_.cf15
            d_1248___mcc_h7_ = source19_.cf16
            d_1249___mcc_h8_ = source19_.cf17
            d_1250___mcc_h9_ = source19_.cf18
            d_1251_cf18_ = d_1250___mcc_h9_
            d_1252_cf17_ = d_1249___mcc_h8_
            d_1253_cf16_ = d_1248___mcc_h7_
            d_1254_cf15_ = d_1247___mcc_h6_
            d_1255_v86_: _dafny.Map
            d_1255_v86_ = _dafny.Map({d_1254_cf15_: d_1135_v0_})
            d_1256_v87_: _dafny.Map
            d_1256_v87_ = _dafny.Map({d_1254_cf15_: (self).f24})
            d_1257_v88_: _dafny.Array
            nw181_ = _dafny.Array(int(0), 6)
            d_1257_v88_ = nw181_
            d_1258_v89_: _dafny.Array
            nw182_ = _dafny.Array(None, 25)
            nw182_[int(0)] = d_1257_v88_
            nw182_[int(1)] = d_1257_v88_
            nw182_[int(2)] = d_1257_v88_
            nw182_[int(3)] = d_1257_v88_
            nw182_[int(4)] = d_1257_v88_
            nw182_[int(5)] = d_1257_v88_
            nw182_[int(6)] = d_1257_v88_
            nw182_[int(7)] = d_1257_v88_
            nw182_[int(8)] = d_1257_v88_
            nw182_[int(9)] = d_1257_v88_
            nw182_[int(10)] = d_1257_v88_
            nw182_[int(11)] = d_1257_v88_
            nw182_[int(12)] = d_1257_v88_
            nw182_[int(13)] = d_1257_v88_
            nw182_[int(14)] = d_1257_v88_
            nw182_[int(15)] = d_1257_v88_
            nw182_[int(16)] = d_1257_v88_
            nw182_[int(17)] = d_1257_v88_
            nw182_[int(18)] = d_1257_v88_
            nw182_[int(19)] = d_1257_v88_
            nw182_[int(20)] = d_1257_v88_
            nw182_[int(21)] = d_1257_v88_
            nw182_[int(22)] = d_1257_v88_
            nw182_[int(23)] = d_1257_v88_
            nw182_[int(24)] = d_1257_v88_
            d_1258_v89_ = nw182_
            d_1259_v90_: D7
            d_1259_v90_ = D7_DC22(d_1252_cf17_, d_1258_v89_)
            d_1260_v91_: D9
            d_1260_v91_ = D9_DC27(d_1135_v0_, d_1259_v90_, (self).f24, True)
            d_1261_v92_: _dafny.Array
            nw183_ = _dafny.Array(None, 25)
            nw183_[int(0)] = d_1135_v0_
            nw183_[int(1)] = False
            nw183_[int(2)] = d_1135_v0_
            nw183_[int(3)] = False
            nw183_[int(4)] = d_1135_v0_
            nw183_[int(5)] = d_1135_v0_
            nw183_[int(6)] = d_1135_v0_
            nw183_[int(7)] = d_1135_v0_
            nw183_[int(8)] = d_1135_v0_
            nw183_[int(9)] = d_1135_v0_
            nw183_[int(10)] = not (d_1135_v0_) or (d_1135_v0_)
            nw183_[int(11)] = d_1135_v0_
            nw183_[int(12)] = d_1135_v0_
            nw183_[int(13)] = d_1135_v0_
            nw183_[int(14)] = (d_1252_cf17_) == (d_1251_cf18_)
            nw183_[int(15)] = (d_1209_v54_).cf4
            nw183_[int(16)] = d_1135_v0_
            nw183_[int(17)] = d_1135_v0_
            nw183_[int(18)] = d_1135_v0_
            nw183_[int(19)] = d_1135_v0_
            nw183_[int(20)] = ((d_1255_v86_)[len(d_1256_v87_)] if (len(d_1256_v87_)) in (d_1255_v86_) else True)
            nw183_[int(21)] = (self).fm27(d_1254_cf15_, (D0_DC1(((d_1225_v70_)[(self).f24] if ((self).f24) in (d_1225_v70_) else (self).f24), d_1208_v53_, d_1252_cf17_, not(d_1135_v0_))).cf2, (d_1260_v91_).cf52, globalState)
            nw183_[int(22)] = ((d_1240_v81_).cf36) == (d_1135_v0_)
            nw183_[int(23)] = d_1135_v0_
            nw183_[int(24)] = d_1135_v0_
            d_1261_v92_ = nw183_
            index141_ = default__.safeIndex(957, (d_1261_v92_).length(0))
            (d_1261_v92_)[index141_] = True
            index142_ = default__.safeIndex(957, (d_1261_v92_).length(0))
            (d_1261_v92_)[index142_] = (d_1135_v0_) not in (d_1221_v66_)
            d_1262_v93_: C2
            nw184_ = C2()
            nw184_.ctor__(d_1252_cf17_, _dafny.SeqWithoutIsStrInference([d_1226_v71_ for d_1263_i5_ in range(default__.abs(369))]), (self).f24, 943, d_1135_v0_)
            d_1262_v93_ = nw184_
            d_1264_v94_: _dafny.Seq
            d_1264_v94_ = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState)])
            d_1254_cf15_ = (0) - ((default__.safeDivisionInt(d_1251_cf18_, (d_1264_v94_)[default__.safeIndex((self).f24, len(d_1264_v94_))])) * (default__.fm2(globalState)))
            r1 = default__.safeModuloInt((0) - (d_1252_cf17_), (self).f24)
        elif source19_.is_DC8:
            d_1265___mcc_h10_ = source19_.cf14
            d_1266_cf14_ = d_1265___mcc_h10_
            d_1267_v95_: _dafny.Array
            nw185_ = _dafny.Array(None, 26)
            nw185_[int(0)] = d_1135_v0_
            nw185_[int(1)] = d_1135_v0_
            nw185_[int(2)] = d_1135_v0_
            nw185_[int(3)] = d_1135_v0_
            nw185_[int(4)] = True
            nw185_[int(5)] = d_1135_v0_
            nw185_[int(6)] = True
            nw185_[int(7)] = d_1135_v0_
            nw185_[int(8)] = d_1135_v0_
            nw185_[int(9)] = d_1135_v0_
            nw185_[int(10)] = d_1135_v0_
            nw185_[int(11)] = d_1135_v0_
            nw185_[int(12)] = True
            nw185_[int(13)] = d_1135_v0_
            nw185_[int(14)] = d_1135_v0_
            nw185_[int(15)] = d_1135_v0_
            nw185_[int(16)] = not(d_1135_v0_)
            nw185_[int(17)] = d_1135_v0_
            nw185_[int(18)] = d_1135_v0_
            nw185_[int(19)] = d_1135_v0_
            nw185_[int(20)] = d_1135_v0_
            nw185_[int(21)] = True
            nw185_[int(22)] = d_1135_v0_
            nw185_[int(23)] = d_1135_v0_
            nw185_[int(24)] = d_1135_v0_
            nw185_[int(25)] = d_1135_v0_
            d_1267_v95_ = nw185_
            (globalState).f8 = (_dafny.SeqWithoutIsStrInference([d_1267_v95_, d_1267_v95_])) <= (p0)
            if (not(d_1135_v0_)) and (d_1135_v0_):
                d_1268_v96_: _dafny.Array
                d_1269_v97_: bool
                out12_: _dafny.Array
                out13_: bool
                out12_, out13_ = default__.m0(globalState)
                d_1268_v96_ = out12_
                d_1269_v97_ = out13_
                index143_ = default__.safeIndex(146, (d_1267_v95_).length(0))
                (d_1267_v95_)[index143_] = d_1269_v97_
                index144_ = default__.safeIndex(146, (d_1267_v95_).length(0))
                (d_1267_v95_)[index144_] = default__.fm1(315, not (d_1269_v97_) or (d_1135_v0_), ((self).f24) * (len((_dafny.Map({d_1269_v97_: (self).f24})).set(d_1269_v97_, (self).f24))), globalState)
                d_1270_v98_: C1
                nw186_ = C1()
                nw186_.ctor__(58)
                d_1270_v98_ = nw186_
                r1 = (self).f24
                d_1271_v99_: C1
                nw187_ = C1()
                nw187_.ctor__(len(d_1208_v53_))
                d_1271_v99_ = nw187_
            elif True:
                r0 = default__.fm2(globalState)
                r0 = (0) - ((self).f24)
                d_1272_v100_: _dafny.Map
                d_1272_v100_ = _dafny.Map({d_1238_v79_: default__.safeModuloInt((self).f24, (self).f24)})
                d_1273_v101_: _dafny.Seq
                d_1273_v101_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
                d_1274_v102_: _dafny.MultiSet
                d_1274_v102_ = _dafny.MultiSet([d_1273_v101_, _dafny.SeqWithoutIsStrInference([(self).f24 for d_1275_i6_ in range(default__.abs(315))])])
                d_1276_v103_: D0
                d_1276_v103_ = D0_DC3(((d_1274_v102_)[_dafny.SeqWithoutIsStrInference([262])] if (_dafny.SeqWithoutIsStrInference([262])) in (d_1274_v102_) else (self).f24), d_1135_v0_)
                d_1277_v104_: _dafny.MultiSet
                d_1277_v104_ = _dafny.MultiSet([d_1276_v103_, (d_1276_v103_ if d_1135_v0_ else d_1276_v103_), d_1276_v103_])
                rhs168_ = (d_1272_v100_) | (d_1272_v100_)
                rhs169_ = (d_1277_v104_) | (default__.fm48(len(_dafny.SeqWithoutIsStrInference([d_1135_v0_, d_1135_v0_])), d_1226_v71_, len(d_1221_v66_), d_1135_v0_, globalState))
                d_1272_v100_ = rhs168_
                d_1277_v104_ = rhs169_
                (globalState).f8 = d_1135_v0_
                r0 = (self).f24
            d_1135_v0_ = (d_1135_v0_ if d_1135_v0_ else d_1135_v0_)
            d_1278_v105_: _dafny.Map
            d_1278_v105_ = _dafny.Map({(self).f24: (self).f24})
            r1 = ((((d_1278_v105_)[(self).f24] if ((self).f24) in (d_1278_v105_) else (self).f24) if d_1135_v0_ else 485) if (d_1226_v71_) in (d_1208_v53_) else ((self).f24) + ((0) - ((self).f24)))
        elif True:
            d_1279___mcc_h11_ = source19_.cf19
            d_1280_cf19_ = d_1279___mcc_h11_
            d_1226_v71_ = d_1226_v71_
            d_1281_v106_: _dafny.Map
            d_1281_v106_ = _dafny.Map({(0) - (-671): 201})
            d_1136_v1_ = (_dafny.Map({d_1135_v0_: len(d_1281_v106_)})) | ((d_1136_v1_).set(d_1135_v0_, (self).f24))
            d_1282_v107_: D9
            d_1282_v107_ = D9_DC26((d_1208_v53_)[default__.safeIndex(len(d_1224_v69_), len(d_1208_v53_))])
            source20_ = d_1282_v107_
            if source20_.is_DC27:
                d_1283___mcc_h12_ = source20_.cf49
                d_1284___mcc_h13_ = source20_.cf50
                d_1285___mcc_h14_ = source20_.cf51
                d_1286___mcc_h15_ = source20_.cf52
                d_1287_cf52_ = d_1286___mcc_h15_
                d_1288_cf51_ = d_1285___mcc_h14_
                d_1289_cf50_ = d_1284___mcc_h13_
                d_1290_cf49_ = d_1283___mcc_h12_
                r0 = (0) - (((d_1288_cf51_) * (d_1288_cf51_)) * (default__.fm2(globalState)))
                d_1136_v1_ = (d_1136_v1_).set(True, (self).f24)
                d_1291_v108_: C4
                nw188_ = C4()
                nw188_.ctor__(default__.safeModuloInt(default__.fm2(globalState), len(_dafny.SeqWithoutIsStrInference([True]))), d_1135_v0_)
                d_1291_v108_ = nw188_
                d_1292_v109_: _dafny.Array
                nw189_ = _dafny.Array(int(0), 3)
                d_1292_v109_ = nw189_
                d_1293_v110_: _dafny.Array
                nw190_ = _dafny.Array(None, 25)
                nw190_[int(0)] = d_1292_v109_
                nw190_[int(1)] = d_1292_v109_
                nw190_[int(2)] = d_1292_v109_
                nw190_[int(3)] = d_1292_v109_
                nw190_[int(4)] = d_1292_v109_
                nw190_[int(5)] = d_1292_v109_
                nw190_[int(6)] = d_1292_v109_
                nw190_[int(7)] = d_1292_v109_
                nw190_[int(8)] = d_1292_v109_
                nw190_[int(9)] = d_1292_v109_
                nw190_[int(10)] = d_1292_v109_
                nw190_[int(11)] = d_1292_v109_
                nw190_[int(12)] = d_1292_v109_
                nw190_[int(13)] = (D0_DC2(d_1288_cf51_, d_1292_v109_)).cf6
                nw190_[int(14)] = d_1292_v109_
                nw190_[int(15)] = d_1292_v109_
                nw190_[int(16)] = d_1292_v109_
                nw190_[int(17)] = d_1292_v109_
                nw190_[int(18)] = d_1292_v109_
                nw190_[int(19)] = d_1292_v109_
                nw190_[int(20)] = d_1292_v109_
                nw190_[int(21)] = d_1292_v109_
                nw190_[int(22)] = d_1292_v109_
                nw190_[int(23)] = d_1292_v109_
                nw190_[int(24)] = d_1292_v109_
                d_1293_v110_ = nw190_
                d_1294_v111_: D7
                d_1294_v111_ = D7_DC23(d_1293_v110_, d_1221_v66_, not(d_1287_cf52_), (self).f24)
                index145_ = default__.safeIndex(34, (d_1292_v109_).length(0))
                (d_1292_v109_)[index145_] = (d_1294_v111_).cf45
                index146_ = default__.safeIndex(34, (d_1292_v109_).length(0))
                (d_1292_v109_)[index146_] = (d_1291_v108_).f25
            elif source20_.is_DC26:
                d_1295___mcc_h16_ = source20_.cf48
                d_1296_cf48_ = d_1295___mcc_h16_
                d_1297_v112_: _dafny.Array
                def lambda75_(d_1298_i7_):
                    return (d_1298_i7_) * ((self).f24)

                init40_ = lambda75_
                nw191_ = _dafny.Array(None, 25)
                for i0_40_ in range(nw191_.length(0)):
                    nw191_[i0_40_] = init40_(i0_40_)
                d_1297_v112_ = nw191_
                d_1299_v113_: _dafny.Seq
                d_1299_v113_ = _dafny.SeqWithoutIsStrInference([d_1297_v112_, d_1297_v112_])
                d_1300_v114_: C3
                nw192_ = C3()
                nw192_.ctor__(len(d_1299_v113_), (self).f24)
                d_1300_v114_ = nw192_
                d_1301_v115_: D0
                d_1301_v115_ = D0_DC3((self).f24, d_1135_v0_)
                d_1302_v116_: _dafny.Array
                def lambda76_(d_1303_v0_):
                    def lambda77_(d_1304_i8_):
                        return not(d_1303_v0_)

                    return lambda77_

                init41_ = lambda76_(d_1135_v0_)
                nw193_ = _dafny.Array(None, 27)
                for i0_41_ in range(nw193_.length(0)):
                    nw193_[i0_41_] = init41_(i0_41_)
                d_1302_v116_ = nw193_
                d_1305_v117_: _dafny.Seq
                d_1305_v117_ = _dafny.SeqWithoutIsStrInference([d_1302_v116_])
                rhs170_ = d_1301_v115_
                rhs171_ = d_1305_v117_
                d_1301_v115_ = rhs170_
                d_1305_v117_ = rhs171_
                d_1306_v119_: _dafny.Seq
                d_1306_v119_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f24: (self).f24}))])
                d_1307_v120_: _dafny.Seq
                def iife141_():
                    coll65_ = _dafny.Map()
                    compr_65_: int
                    for compr_65_ in (d_1306_v119_).Elements:
                        d_1308_v118_: int = compr_65_
                        if (d_1308_v118_) in (d_1306_v119_):
                            coll65_[(d_1308_v118_) + ((self).f24)] = d_1135_v0_
                    return _dafny.Map(coll65_)
                d_1307_v120_ = _dafny.SeqWithoutIsStrInference([len(iife141_()
                ), (0) - ((self).f24), (self).f24, (self).f24, (self).f24])
                (globalState).f5 = (d_1306_v119_) + (d_1307_v120_)
                d_1309_v121_: _dafny.Array
                nw194_ = _dafny.Array(None, 13)
                nw194_[int(0)] = d_1302_v116_
                nw194_[int(1)] = d_1302_v116_
                nw194_[int(2)] = d_1302_v116_
                nw194_[int(3)] = d_1302_v116_
                nw194_[int(4)] = d_1302_v116_
                nw194_[int(5)] = d_1302_v116_
                nw194_[int(6)] = d_1302_v116_
                nw194_[int(7)] = d_1302_v116_
                nw194_[int(8)] = (d_1302_v116_ if not(d_1135_v0_) else d_1302_v116_)
                nw194_[int(9)] = d_1302_v116_
                nw194_[int(10)] = (d_1305_v117_)[default__.safeIndex(375, len(d_1305_v117_))]
                nw194_[int(11)] = d_1302_v116_
                nw194_[int(12)] = d_1302_v116_
                d_1309_v121_ = nw194_
                d_1309_v121_ = (d_1309_v121_ if d_1135_v0_ else d_1309_v121_)
            elif True:
                d_1310___mcc_h17_ = source20_.cf53
                d_1311_cf53_ = d_1310___mcc_h17_
                d_1312_v122_: _dafny.Array
                nw195_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_1312_v122_ = nw195_
                index147_ = default__.safeIndex(9, (d_1312_v122_).length(0))
                (d_1312_v122_)[index147_] = d_1208_v53_
                index148_ = default__.safeIndex(9, (d_1312_v122_).length(0))
                (d_1312_v122_)[index148_] = _dafny.SeqWithoutIsStrInference([d_1226_v71_ for d_1313_i9_ in range(default__.abs(-765))])
                d_1226_v71_ = ((d_1312_v122_)[default__.safeIndex(9, (d_1312_v122_).length(0))])[default__.safeIndex(default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdjws"))).set(default__.safeIndex(-193, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdjws")))), d_1226_v71_)), (self).f24), len((d_1312_v122_)[default__.safeIndex(9, (d_1312_v122_).length(0))]))]
                d_1314_v123_: _dafny.Array
                nw196_ = _dafny.Array(False, 4)
                d_1314_v123_ = nw196_
                d_1315_v124_: _dafny.Seq
                d_1315_v124_ = _dafny.SeqWithoutIsStrInference([d_1314_v123_])
                d_1315_v124_ = (d_1315_v124_) + (d_1315_v124_)
                d_1316_v125_: _dafny.Array
                d_1317_v126_: bool
                out14_: _dafny.Array
                out15_: bool
                out14_, out15_ = default__.m0(globalState)
                d_1316_v125_ = out14_
                d_1317_v126_ = out15_
            d_1318_v127_: D5
            d_1318_v127_ = D5_DC15((self).f24, (self).f24, (self).f24)
            d_1319_v128_: _dafny.Map
            d_1319_v128_ = _dafny.Map({(self).f24: d_1318_v127_})
            d_1320_v129_: D6
            d_1320_v129_ = D6_DC16(d_1319_v128_)
            d_1321_v130_: D6
            d_1321_v130_ = D6_DC20(D6_DC20(d_1320_v129_))
            source21_ = d_1321_v130_
            if source21_.is_DC17:
                d_1322___mcc_h18_ = source21_.cf30
                d_1323___mcc_h19_ = source21_.cf31
                d_1324_cf31_ = d_1323___mcc_h19_
                d_1325_cf30_ = d_1322___mcc_h18_
                d_1326_v131_: _dafny.Array
                nw197_ = _dafny.Array(None, 8)
                nw197_[int(0)] = d_1325_cf30_
                nw197_[int(1)] = (d_1224_v69_).isdisjoint(d_1224_v69_)
                nw197_[int(2)] = d_1135_v0_
                nw197_[int(3)] = d_1135_v0_
                nw197_[int(4)] = d_1135_v0_
                nw197_[int(5)] = d_1135_v0_
                nw197_[int(6)] = d_1135_v0_
                nw197_[int(7)] = (d_1325_cf30_) == (d_1325_cf30_)
                d_1326_v131_ = nw197_
                d_1326_v131_ = d_1326_v131_
                d_1327_v132_: _dafny.Map
                d_1327_v132_ = _dafny.Map({_dafny.Map({d_1324_cf31_: d_1135_v0_}): default__.fm49(globalState)})
                d_1328_v133_: _dafny.Set
                d_1328_v133_ = _dafny.Set({d_1224_v69_})
                d_1327_v132_ = (d_1327_v132_).set(default__.fm47(not((self).fm27(d_1324_cf31_, default__.fm36(d_1324_cf31_, d_1324_cf31_, d_1226_v71_, globalState), False, globalState)), d_1325_cf30_, d_1136_v1_, len(d_1208_v53_), globalState), d_1328_v133_)
                d_1329_v134_: _dafny.Array
                def lambda78_(d_1330_i10_):
                    return default__.safeDivisionInt(d_1330_i10_, 744)

                init42_ = lambda78_
                nw198_ = _dafny.Array(None, 18)
                for i0_42_ in range(nw198_.length(0)):
                    nw198_[i0_42_] = init42_(i0_42_)
                d_1329_v134_ = nw198_
                index149_ = default__.safeIndex(205, (d_1329_v134_).length(0))
                (d_1329_v134_)[index149_] = ((0) - (d_1324_cf31_)) + (d_1324_cf31_)
                index150_ = default__.safeIndex(205, (d_1329_v134_).length(0))
                (d_1329_v134_)[index150_] = (self).f24
                r2 = ((d_1208_v53_) + (d_1208_v53_)) + ((d_1208_v53_) + (d_1208_v53_))
            elif source21_.is_DC18:
                d_1331___mcc_h20_ = source21_.cf32
                d_1332_cf32_ = d_1331___mcc_h20_
                d_1333_v135_: _dafny.Array
                d_1334_v136_: bool
                out16_: _dafny.Array
                out17_: bool
                out16_, out17_ = default__.m0(globalState)
                d_1333_v135_ = out16_
                d_1334_v136_ = out17_
                (globalState).f8 = d_1332_cf32_
                d_1335_v137_: _dafny.Array
                nw199_ = _dafny.Array(int(0), 29)
                d_1335_v137_ = nw199_
                index151_ = default__.safeIndex(819, (d_1335_v137_).length(0))
                (d_1335_v137_)[index151_] = ((self).f24) * ((self).f24)
                index152_ = default__.safeIndex(819, (d_1335_v137_).length(0))
                (d_1335_v137_)[index152_] = (self).f24
                (globalState).f8 = (d_1135_v0_ if True else d_1332_cf32_)
            elif source21_.is_DC19:
                d_1336___mcc_h21_ = source21_.cf33
                d_1337___mcc_h22_ = source21_.cf34
                d_1338___mcc_h23_ = source21_.cf35
                d_1339___mcc_h24_ = source21_.cf36
                d_1340___mcc_h25_ = source21_.cf37
                d_1341_cf37_ = d_1340___mcc_h25_
                d_1342_cf36_ = d_1339___mcc_h24_
                d_1343_cf35_ = d_1338___mcc_h23_
                d_1344_cf34_ = d_1337___mcc_h22_
                d_1345_cf33_ = d_1336___mcc_h21_
                d_1346_v138_: _dafny.Map
                d_1346_v138_ = _dafny.Map({(self).f24: d_1208_v53_})
                d_1208_v53_ = ((d_1346_v138_)[(self).f24] if ((self).f24) in (d_1346_v138_) else d_1208_v53_)
                d_1347_v139_: _dafny.Seq
                d_1347_v139_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1348_v140_: _dafny.Array
                nw200_ = _dafny.Array(None, 4)
                nw200_[int(0)] = d_1347_v139_
                nw200_[int(1)] = _dafny.SeqWithoutIsStrInference([default__.fm2(globalState), (self).f24, (self).f24, (d_1225_v70_).cardinality, (self).f24])
                nw200_[int(2)] = d_1347_v139_
                nw200_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_1347_v139_)[default__.safeIndex((self).f24, len(d_1347_v139_))] for d_1349_i11_ in range(default__.abs(96))])
                d_1348_v140_ = nw200_
                d_1350_v141_: D0
                d_1350_v141_ = D0_DC0(d_1348_v140_)
                pat_let_tv56_ = d_1348_v140_
                def iife142_(_pat_let38_0):
                    def iife143_(d_1351_dt__update__tmp_h4_):
                        def iife144_(_pat_let39_0):
                            def iife145_(d_1352_dt__update_hcf0_h0_):
                                return D0_DC0(d_1352_dt__update_hcf0_h0_)
                            return iife145_(_pat_let39_0)
                        return iife144_(pat_let_tv56_)
                    return iife143_(_pat_let38_0)
                d_1350_v141_ = iife142_(d_1350_v141_)
                r0 = (self).f24
                d_1353_v142_: D4
                d_1353_v142_ = D4_DC13((_dafny.MultiSet(d_1341_cf37_)).cardinality)
                d_1354_v143_: _dafny.Map
                d_1354_v143_ = _dafny.Map({(0) - ((self).f24): d_1353_v142_})
                d_1355_v144_: _dafny.Map
                d_1355_v144_ = _dafny.Map({d_1343_cf35_: d_1354_v143_})
                d_1355_v144_ = d_1355_v144_
            elif source21_.is_DC16:
                d_1356___mcc_h26_ = source21_.cf29
                d_1357_cf29_ = d_1356___mcc_h26_
                d_1358_v145_: _dafny.Seq
                d_1358_v145_ = _dafny.SeqWithoutIsStrInference([len(d_1221_v66_), (self).f24, (self).f24, default__.fm2(globalState), default__.fm2(globalState)])
                (globalState).f8 = (((self).f24) - ((self).f24)) in (d_1358_v145_)
                d_1359_v146_: _dafny.Array
                def lambda79_(d_1360_i12_):
                    return ((self).f24) > ((self).f24)

                init43_ = lambda79_
                nw201_ = _dafny.Array(None, 6)
                for i0_43_ in range(nw201_.length(0)):
                    nw201_[i0_43_] = init43_(i0_43_)
                d_1359_v146_ = nw201_
                d_1361_v147_: _dafny.Map
                d_1361_v147_ = _dafny.Map({len(d_1208_v53_): d_1135_v0_})
                d_1362_v148_: _dafny.Map
                d_1362_v148_ = _dafny.Map({(self).f24: d_1358_v145_})
                index153_ = default__.safeIndex(219, (d_1359_v146_).length(0))
                (d_1359_v146_)[index153_] = ((d_1361_v147_)[(0) - ((self).f24)] if ((0) - ((self).f24)) in (d_1361_v147_) else default__.fm1(671, d_1135_v0_, len(d_1362_v148_), globalState))
                index154_ = default__.safeIndex(219, (d_1359_v146_).length(0))
                (d_1359_v146_)[index154_] = False
                d_1363_v149_: _dafny.Array
                nw202_ = _dafny.Array(int(0), 22)
                d_1363_v149_ = nw202_
                d_1363_v149_ = d_1363_v149_
                d_1364_v150_: _dafny.Array
                nw203_ = _dafny.Array(None, 3)
                nw203_[int(0)] = (_dafny.SeqWithoutIsStrInference([d_1135_v0_])) + (d_1221_v66_)
                nw203_[int(1)] = d_1221_v66_
                nw203_[int(2)] = (d_1221_v66_) + (d_1221_v66_)
                d_1364_v150_ = nw203_
                index155_ = default__.safeIndex(877, (d_1364_v150_).length(0))
                (d_1364_v150_)[index155_] = (d_1221_v66_) + (default__.fm39((d_1359_v146_)[default__.safeIndex(219, (d_1359_v146_).length(0))], (d_1359_v146_)[default__.safeIndex(219, (d_1359_v146_).length(0))], globalState))
                d_1365_v151_: _dafny.Map
                d_1365_v151_ = _dafny.Map({d_1226_v71_: d_1221_v66_})
                index156_ = default__.safeIndex(877, (d_1364_v150_).length(0))
                (d_1364_v150_)[index156_] = (d_1221_v66_) + (((d_1365_v151_)[d_1226_v71_] if (d_1226_v71_) in (d_1365_v151_) else default__.fm39((d_1359_v146_)[default__.safeIndex(219, (d_1359_v146_).length(0))], d_1135_v0_, globalState)))
            elif True:
                d_1366___mcc_h27_ = source21_.cf38
                d_1367_cf38_ = d_1366___mcc_h27_
                d_1368_v152_: _dafny.Array
                nw204_ = _dafny.Array(False, 22)
                d_1368_v152_ = nw204_
                d_1369_v153_: _dafny.Map
                d_1369_v153_ = _dafny.Map({d_1135_v0_: d_1226_v71_})
                index157_ = default__.safeIndex(977, (d_1368_v152_).length(0))
                (d_1368_v152_)[index157_] = (d_1135_v0_) not in (d_1369_v153_)
                index158_ = default__.safeIndex(977, (d_1368_v152_).length(0))
                (d_1368_v152_)[index158_] = d_1135_v0_
                (globalState).f8 = not(d_1135_v0_)
                d_1370_v154_: D1
                d_1370_v154_ = D1_DC5(_dafny.MultiSet([(self).f24, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okfmdc"))).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okfmdc")))), d_1226_v71_))]), d_1135_v0_)
                d_1225_v70_ = (d_1225_v70_) - ((d_1370_v154_).cf10)
                r0 = (len(d_1208_v53_)) + ((self).f24)
        r0 = (self).f24
        r1 = (self).f24
        d_1371_v155_: _dafny.MultiSet
        d_1371_v155_ = _dafny.MultiSet([d_1238_v79_, d_1238_v79_, (_dafny.Map({d_1135_v0_: (d_1221_v66_)[default__.safeIndex((self).f24, len(d_1221_v66_))]})).set(d_1135_v0_, d_1135_v0_)])
        d_1372_v156_: _dafny.MultiSet
        d_1372_v156_ = d_1371_v155_
        pat_let_tv57_ = d_1208_v53_
        pat_let_tv58_ = d_1208_v53_
        def lambda80_(source22_):
            d_1373___mcc_h28_ = source22_
            d_1374_cf47_ = d_1373___mcc_h28_
            return (pat_let_tv57_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sywamaofb")))

        def lambda81_(source23_):
            d_1375___mcc_h29_ = source23_
            d_1376_cf47_ = d_1375___mcc_h29_
            return (pat_let_tv58_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sywamaofb")))

        r2 = (lambda80_(d_1372_v156_)).set(default__.safeIndex((0) - ((self).f24), len(lambda81_(d_1372_v156_))), ((d_1208_v53_)[default__.safeIndex((self).f24, len(d_1208_v53_))] if d_1135_v0_ else d_1226_v71_))
        return r0, r1, r2

    @property
    def f24(self):
        return self._f24

class C6:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm24(self, p0, globalState):
        if (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfqdouu"))])).isdisjoint(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjnse")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfgxspivk"))])):
            return _dafny.Map({_dafny.CodePoint('t'): -283})
        elif True:
            def iife146_():
                coll66_ = _dafny.Map()
                compr_66_: str
                for compr_66_ in (_dafny.Map({_dafny.CodePoint('w'): len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([339])).cardinality, 153, 441, (0) - (-446)]))})).keys.Elements:
                    d_1377_v0_: str = compr_66_
                    if (d_1377_v0_) in (_dafny.Map({_dafny.CodePoint('w'): len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([339])).cardinality, 153, 441, (0) - (-446)]))})):
                        coll66_[d_1377_v0_] = 431
                return _dafny.Map(coll66_)
            def iife147_():
                coll67_ = _dafny.Map()
                compr_67_: str
                for compr_67_ in (_dafny.Map({_dafny.CodePoint('t'): True})).keys.Elements:
                    d_1378_v1_: str = compr_67_
                    if (d_1378_v1_) in (_dafny.Map({_dafny.CodePoint('t'): True})):
                        coll67_[d_1378_v1_] = -180
                return _dafny.Map(coll67_)
            return (iife146_()
            ) | (iife147_()
            )

    def m18(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: D0 = D0.default()()
        d_1379_v0_: bool
        d_1379_v0_ = False
        d_1380_v1_: int
        d_1380_v1_ = 602
        d_1381_v2_: _dafny.Seq
        d_1381_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hybdjuc"))
        d_1382_v3_: D0
        d_1382_v3_ = D0_DC1(d_1380_v1_, d_1381_v2_, d_1380_v1_, default__.fm1(d_1380_v1_, d_1379_v0_, d_1380_v1_, globalState))
        d_1383_v4_: _dafny.Map
        d_1383_v4_ = _dafny.Map({d_1379_v0_: d_1382_v3_})
        d_1384_v5_: _dafny.Seq
        d_1384_v5_ = _dafny.SeqWithoutIsStrInference([d_1379_v0_])
        d_1383_v4_ = (d_1383_v4_).set((d_1384_v5_)[default__.safeIndex((_dafny.MultiSet([True])).cardinality, len(d_1384_v5_))], d_1382_v3_)
        d_1379_v0_ = d_1379_v0_
        d_1385_i0_: int
        d_1385_i0_ = 0
        with _dafny.label("4"):
            while d_1379_v0_:
                with _dafny.c_label("4"):
                    if (d_1385_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_1385_i0_ = (d_1385_i0_) + (1)
                    r1 = d_1380_v1_
                    d_1386_v6_: _dafny.Array
                    nw205_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_1386_v6_ = nw205_
                    nw206_ = _dafny.Array(_dafny.Array(None, 0), 4)
                    d_1386_v6_ = nw206_
                    r0 = (399) != (default__.safeModuloInt(d_1380_v1_, d_1380_v1_))
                    r0 = d_1379_v0_
                    pass
            pass
        if d_1379_v0_:
            d_1380_v1_ = 183
            d_1387_v7_: _dafny.Array
            nw207_ = _dafny.Array(None, 12)
            nw207_[int(0)] = (d_1380_v1_) != (d_1380_v1_)
            nw207_[int(1)] = default__.fm1(d_1380_v1_, d_1379_v0_, (0) - (d_1380_v1_), globalState)
            nw207_[int(2)] = d_1379_v0_
            nw207_[int(3)] = (d_1384_v5_)[default__.safeIndex(d_1380_v1_, len(d_1384_v5_))]
            nw207_[int(4)] = (d_1379_v0_ if d_1379_v0_ else not(d_1379_v0_))
            nw207_[int(5)] = d_1379_v0_
            nw207_[int(6)] = d_1379_v0_
            nw207_[int(7)] = d_1379_v0_
            nw207_[int(8)] = d_1379_v0_
            nw207_[int(9)] = d_1379_v0_
            nw207_[int(10)] = d_1379_v0_
            nw207_[int(11)] = d_1379_v0_
            d_1387_v7_ = nw207_
            index159_ = default__.safeIndex(966, (d_1387_v7_).length(0))
            (d_1387_v7_)[index159_] = d_1379_v0_
            index160_ = default__.safeIndex(966, (d_1387_v7_).length(0))
            (d_1387_v7_)[index160_] = default__.fm1(d_1380_v1_, d_1379_v0_, d_1380_v1_, globalState)
            (globalState).f8 = default__.fm1(842, (d_1387_v7_)[default__.safeIndex(966, (d_1387_v7_).length(0))], d_1380_v1_, globalState)
            d_1388_v8_: _dafny.Map
            d_1388_v8_ = _dafny.Map({(d_1380_v1_) * (d_1380_v1_): (d_1387_v7_)[default__.safeIndex(966, (d_1387_v7_).length(0))]})
            d_1388_v8_ = (d_1388_v8_).set(d_1380_v1_, (d_1387_v7_)[default__.safeIndex(966, (d_1387_v7_).length(0))])
            d_1389_v9_: _dafny.Map
            d_1389_v9_ = _dafny.Map({_dafny.CodePoint('a'): default__.fm2(globalState)})
            d_1390_v10_: str
            d_1390_v10_ = _dafny.CodePoint('o')
            d_1391_v11_: _dafny.Map
            d_1391_v11_ = _dafny.Map({((d_1381_v2_).set(default__.safeIndex(len(d_1389_v9_), len(d_1381_v2_)), d_1390_v10_)) + (d_1381_v2_): d_1380_v1_})
            d_1391_v11_ = (d_1391_v11_).set(d_1381_v2_, d_1380_v1_)
        elif True:
            (globalState).f8 = d_1379_v0_
            d_1392_v12_: str
            d_1392_v12_ = _dafny.CodePoint('t')
            d_1393_v13_: _dafny.Map
            d_1393_v13_ = _dafny.Map({d_1380_v1_: d_1392_v12_})
            d_1394_v14_: _dafny.Map
            d_1394_v14_ = _dafny.Map({d_1380_v1_: d_1379_v0_})
            d_1395_v15_: _dafny.Map
            d_1395_v15_ = _dafny.Map({len(d_1394_v14_): d_1379_v0_})
            d_1396_v16_: _dafny.MultiSet
            d_1396_v16_ = _dafny.MultiSet([_dafny.Map({len(d_1393_v13_): d_1379_v0_}), d_1394_v14_])
            d_1397_v17_: _dafny.Set
            d_1397_v17_ = _dafny.Set({d_1379_v0_, d_1379_v0_})
            d_1398_v18_: _dafny.Map
            d_1398_v18_ = _dafny.Map({default__.fm25((d_1396_v16_).cardinality, len(d_1397_v17_), globalState): False})
            d_1379_v0_ = (d_1398_v18_) != ((d_1398_v18_) | (d_1398_v18_))
            d_1399_v19_: _dafny.Array
            def lambda82_(d_1400_v1_):
                def lambda83_(d_1401_i1_):
                    return (d_1401_i1_) - (d_1400_v1_)

                return lambda83_

            init44_ = lambda82_(d_1380_v1_)
            nw208_ = _dafny.Array(None, 18)
            for i0_44_ in range(nw208_.length(0)):
                nw208_[i0_44_] = init44_(i0_44_)
            d_1399_v19_ = nw208_
            index161_ = default__.safeIndex(972, (d_1399_v19_).length(0))
            (d_1399_v19_)[index161_] = (len(d_1381_v2_)) - (d_1380_v1_)
            index162_ = default__.safeIndex(972, (d_1399_v19_).length(0))
            (d_1399_v19_)[index162_] = d_1380_v1_
            d_1402_v20_: _dafny.Array
            def lambda84_(d_1403_v2_):
                def lambda85_(d_1404_i2_):
                    return d_1403_v2_

                return lambda85_

            init45_ = lambda84_(d_1381_v2_)
            nw209_ = _dafny.Array(None, 12)
            for i0_45_ in range(nw209_.length(0)):
                nw209_[i0_45_] = init45_(i0_45_)
            d_1402_v20_ = nw209_
            d_1405_v21_: _dafny.Map
            d_1405_v21_ = _dafny.Map({(d_1399_v19_)[default__.safeIndex(972, (d_1399_v19_).length(0))]: d_1402_v20_})
            d_1405_v21_ = (d_1405_v21_).set((0) - ((d_1399_v19_)[default__.safeIndex(972, (d_1399_v19_).length(0))]), d_1402_v20_)
            d_1406_v22_: _dafny.Seq
            d_1406_v22_ = _dafny.SeqWithoutIsStrInference([952])
            d_1407_v23_: _dafny.Map
            d_1407_v23_ = _dafny.Map({d_1379_v0_: d_1406_v22_})
            (globalState).f5 = (((d_1407_v23_)[d_1379_v0_] if (d_1379_v0_) in (d_1407_v23_) else _dafny.SeqWithoutIsStrInference([d_1380_v1_, (d_1399_v19_)[default__.safeIndex(972, (d_1399_v19_).length(0))], (d_1399_v19_)[default__.safeIndex(972, (d_1399_v19_).length(0))], (d_1399_v19_)[default__.safeIndex(972, (d_1399_v19_).length(0))], d_1380_v1_]))) + (d_1406_v22_)
        if d_1379_v0_:
            d_1408_v24_: str
            d_1408_v24_ = _dafny.CodePoint('w')
            d_1409_v25_: _dafny.Map
            d_1409_v25_ = _dafny.Map({d_1408_v24_: d_1380_v1_})
            d_1410_v26_: _dafny.Map
            d_1410_v26_ = _dafny.Map({d_1380_v1_: d_1379_v0_})
            d_1411_v27_: _dafny.Seq
            d_1411_v27_ = _dafny.SeqWithoutIsStrInference([d_1410_v26_, (d_1410_v26_).set(d_1380_v1_, not(d_1379_v0_)), d_1410_v26_])
            d_1412_v28_: C0
            nw210_ = C0()
            nw210_.ctor__(d_1380_v1_, d_1379_v0_, (0) - (len(d_1409_v25_)), (_dafny.MultiSet(d_1411_v27_)).cardinality)
            d_1412_v28_ = nw210_
            d_1413_v29_: _dafny.MultiSet
            d_1413_v29_ = _dafny.MultiSet([d_1379_v0_, d_1379_v0_])
            d_1414_v30_: _dafny.Seq
            d_1414_v30_ = _dafny.SeqWithoutIsStrInference([d_1413_v29_, d_1413_v29_])
            d_1415_v31_: _dafny.Map
            d_1415_v31_ = _dafny.Map({(d_1379_v0_ if d_1412_v28_.f29 else not(True)): (d_1414_v30_)[default__.safeIndex((d_1412_v28_).f28, len(d_1414_v30_))]})
            d_1415_v31_ = (d_1415_v31_ if d_1379_v0_ else d_1415_v31_)
            (globalState).f8 = d_1412_v28_.f29
            d_1416_v32_: _dafny.Map
            d_1416_v32_ = _dafny.Map({(d_1412_v28_).f28: default__.safeModuloInt(183, -824)})
            d_1416_v32_ = (d_1416_v32_).set((d_1412_v28_).f28, 682)
            r1 = default__.safeDivisionInt(d_1380_v1_, d_1380_v1_)
        elif True:
            d_1417_v33_: str
            d_1417_v33_ = _dafny.CodePoint('c')
            d_1418_v34_: _dafny.Map
            d_1418_v34_ = _dafny.Map({d_1417_v33_: d_1380_v1_})
            d_1419_v35_: _dafny.Set
            d_1419_v35_ = _dafny.Set({len(d_1381_v2_), ((d_1418_v34_)[_dafny.CodePoint('t')] if (_dafny.CodePoint('t')) in (d_1418_v34_) else 257)})
            d_1420_v36_: _dafny.Array
            nw211_ = _dafny.Array(int(0), 25)
            d_1420_v36_ = nw211_
            d_1421_v37_: _dafny.MultiSet
            d_1421_v37_ = _dafny.MultiSet([d_1379_v0_, d_1379_v0_, d_1379_v0_])
            d_1422_v38_: _dafny.Array
            out18_: _dafny.Array
            out18_ = (self).m19(d_1419_v35_, d_1420_v36_, d_1380_v1_, d_1421_v37_, globalState)
            d_1422_v38_ = out18_
            d_1380_v1_ = (d_1380_v1_) * ((d_1380_v1_ if d_1379_v0_ else 286))
            index163_ = default__.safeIndex(199, (d_1422_v38_).length(0))
            (d_1422_v38_)[index163_] = d_1379_v0_
            d_1423_v39_: _dafny.Map
            d_1423_v39_ = _dafny.Map({d_1422_v38_: d_1379_v0_})
            index164_ = default__.safeIndex(199, (d_1422_v38_).length(0))
            rhs172_ = (d_1379_v0_) == ((d_1380_v1_) <= (d_1380_v1_))
            rhs173_ = d_1379_v0_
            rhs174_ = d_1379_v0_
            rhs175_ = _dafny.SeqWithoutIsStrInference([not(((d_1423_v39_)[d_1422_v38_] if (d_1422_v38_) in (d_1423_v39_) else d_1379_v0_)), d_1379_v0_])
            rhs176_ = (d_1380_v1_) > ((d_1380_v1_) - (d_1380_v1_))
            lhs94_ = globalState
            lhs95_ = d_1422_v38_
            lhs96_ = default__.safeIndex(199, (d_1422_v38_).length(0))
            lhs94_.f8 = rhs172_
            d_1379_v0_ = rhs173_
            r0 = rhs174_
            d_1384_v5_ = rhs175_
            lhs95_[lhs96_] = rhs176_
            d_1417_v33_ = d_1417_v33_
            d_1424_v40_: _dafny.Array
            def lambda86_(d_1425_v2_):
                def lambda87_(d_1426_i3_):
                    return d_1425_v2_

                return lambda87_

            init46_ = lambda86_(d_1381_v2_)
            nw212_ = _dafny.Array(None, 12)
            for i0_46_ in range(nw212_.length(0)):
                nw212_[i0_46_] = init46_(i0_46_)
            d_1424_v40_ = nw212_
            d_1427_v41_: D5
            d_1427_v41_ = D5_DC14(d_1424_v40_)
            d_1428_v42_: _dafny.Map
            d_1428_v42_ = _dafny.Map({d_1427_v41_: (d_1422_v38_)[default__.safeIndex(199, (d_1422_v38_).length(0))]})
            d_1429_v43_: _dafny.Seq
            d_1429_v43_ = _dafny.SeqWithoutIsStrInference([(0) - (len((d_1428_v42_).set(d_1427_v41_, (d_1422_v38_)[default__.safeIndex(199, (d_1422_v38_).length(0))]))), d_1380_v1_, d_1380_v1_])
            d_1430_v44_: _dafny.Map
            d_1430_v44_ = _dafny.Map({d_1417_v33_: ((d_1429_v43_)[default__.safeIndex(d_1380_v1_, len(d_1429_v43_))]) <= (d_1380_v1_)})
            d_1431_v45_: _dafny.MultiSet
            d_1431_v45_ = _dafny.MultiSet([d_1380_v1_, d_1380_v1_, d_1380_v1_])
            d_1430_v44_ = (d_1430_v44_).set(d_1417_v33_, (((d_1431_v45_)[d_1380_v1_] if (d_1380_v1_) in (d_1431_v45_) else d_1380_v1_)) < (d_1380_v1_))
        r0 = (d_1380_v1_) > (d_1380_v1_)
        r0 = d_1379_v0_
        d_1432_v46_: _dafny.Map
        d_1432_v46_ = _dafny.Map({True: d_1380_v1_})
        d_1433_v47_: _dafny.Map
        d_1433_v47_ = _dafny.Map({default__.fm50(globalState): (len(_dafny.SeqWithoutIsStrInference([len(d_1432_v46_)]))) + (d_1380_v1_)})
        d_1434_v48_: _dafny.Set
        d_1434_v48_ = _dafny.Set({d_1380_v1_})
        r1 = ((d_1433_v47_)[d_1434_v48_] if (d_1434_v48_) in (d_1433_v47_) else d_1380_v1_)
        d_1435_v49_: _dafny.Array
        nw213_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_1435_v49_ = nw213_
        d_1436_v50_: D7
        d_1436_v50_ = D7_DC22(d_1380_v1_, d_1435_v49_)
        d_1437_v51_: D9
        d_1437_v51_ = D9_DC27(d_1379_v0_, d_1436_v50_, d_1380_v1_, (d_1384_v5_)[default__.safeIndex(d_1380_v1_, len(d_1384_v5_))])
        d_1438_v52_: _dafny.Array
        nw214_ = _dafny.Array(None, 4)
        nw214_[int(0)] = d_1379_v0_
        nw214_[int(1)] = (d_1437_v51_).cf49
        nw214_[int(2)] = d_1379_v0_
        nw214_[int(3)] = not(d_1379_v0_)
        d_1438_v52_ = nw214_
        d_1439_v53_: _dafny.MultiSet
        d_1439_v53_ = _dafny.MultiSet([d_1438_v52_])
        r2 = d_1439_v53_
        d_1440_v54_: _dafny.Array
        def lambda88_(d_1441_v1_):
            def lambda89_(d_1442_i4_):
                return _dafny.SeqWithoutIsStrInference([d_1441_v1_])

            return lambda89_

        init47_ = lambda88_(d_1380_v1_)
        nw215_ = _dafny.Array(None, 17)
        for i0_47_ in range(nw215_.length(0)):
            nw215_[i0_47_] = init47_(i0_47_)
        d_1440_v54_ = nw215_
        r3 = D0_DC0(d_1440_v54_)
        return r0, r1, r2, r3

    def m19(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_1443_v0_: _dafny.Array
        nw216_ = _dafny.Array(_dafny.Array(None, 0), 24)
        d_1443_v0_ = nw216_
        d_1443_v0_ = d_1443_v0_
        d_1444_v1_: bool
        d_1444_v1_ = True
        d_1445_v2_: _dafny.MultiSet
        d_1445_v2_ = _dafny.MultiSet([p2])
        d_1446_v3_: str
        d_1446_v3_ = _dafny.CodePoint('i')
        d_1447_v4_: _dafny.Array
        nw217_ = _dafny.Array(None, 17)
        nw217_[int(0)] = d_1444_v1_
        nw217_[int(1)] = default__.fm1(p2, d_1444_v1_, -777, globalState)
        nw217_[int(2)] = d_1444_v1_
        nw217_[int(3)] = d_1444_v1_
        nw217_[int(4)] = (94) < (181)
        nw217_[int(5)] = d_1444_v1_
        nw217_[int(6)] = d_1444_v1_
        nw217_[int(7)] = d_1444_v1_
        nw217_[int(8)] = not (d_1444_v1_) or (d_1444_v1_)
        nw217_[int(9)] = d_1444_v1_
        nw217_[int(10)] = d_1444_v1_
        nw217_[int(11)] = default__.fm1(len(default__.fm51((d_1445_v2_).cardinality, p2, p2, d_1446_v3_, globalState)), d_1444_v1_, p2, globalState)
        nw217_[int(12)] = d_1444_v1_
        nw217_[int(13)] = d_1444_v1_
        nw217_[int(14)] = d_1444_v1_
        nw217_[int(15)] = d_1444_v1_
        nw217_[int(16)] = (p2) <= (p2)
        d_1447_v4_ = nw217_
        index165_ = default__.safeIndex(601, (d_1447_v4_).length(0))
        (d_1447_v4_)[index165_] = not(d_1444_v1_)
        index166_ = default__.safeIndex(601, (d_1447_v4_).length(0))
        (d_1447_v4_)[index166_] = d_1444_v1_
        index167_ = default__.safeIndex(601, (d_1447_v4_).length(0))
        (d_1447_v4_)[index167_] = True
        d_1448_v5_: _dafny.Seq
        d_1448_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        if (d_1448_v5_) != ((d_1448_v5_ if (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))] else d_1448_v5_)):
            d_1449_v6_: _dafny.Set
            d_1449_v6_ = _dafny.Set({p2, len((d_1448_v5_) + (d_1448_v5_))})
            d_1450_v7_: _dafny.Set
            d_1450_v7_ = p0
            d_1449_v6_ = (d_1450_v7_)
            d_1451_v8_: D4
            d_1451_v8_ = D4_DC13(473)
            d_1452_v9_: _dafny.Map
            d_1452_v9_ = _dafny.Map({d_1451_v8_: _dafny.MultiSet([d_1444_v1_])})
            d_1453_v10_: _dafny.Map
            d_1453_v10_ = _dafny.Map({(_dafny.Map({D4_DC13(461): p3}) if d_1444_v1_ else d_1452_v9_): default__.safeModuloInt(len((d_1448_v5_).set(default__.safeIndex(((p3)[(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]] if ((d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]) in (p3) else p2), len(d_1448_v5_)), d_1446_v3_)), p2)})
            d_1453_v10_ = (d_1453_v10_).set(d_1452_v9_, (p2) + (p2))
            d_1454_v11_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1454_v11_ = nw218_
            d_1455_v12_: _dafny.Array
            def lambda90_(d_1456_v1_, d_1457_v4_):
                def lambda91_(d_1458_i0_):
                    return _dafny.Set({d_1456_v1_, not((d_1457_v4_)[default__.safeIndex(601, (d_1457_v4_).length(0))])})

                return lambda91_

            init48_ = lambda90_(d_1444_v1_, d_1447_v4_)
            nw219_ = _dafny.Array(None, 6)
            for i0_48_ in range(nw219_.length(0)):
                nw219_[i0_48_] = init48_(i0_48_)
            d_1455_v12_ = nw219_
            index168_ = default__.safeIndex(798, (d_1454_v11_).length(0))
            (d_1454_v11_)[index168_] = d_1455_v12_
            index169_ = default__.safeIndex(798, (d_1454_v11_).length(0))
            (d_1454_v11_)[index169_] = d_1455_v12_
            index170_ = default__.safeIndex(227, (p1).length(0))
            (p1)[index170_] = default__.fm2(globalState)
            d_1459_v13_: _dafny.Map
            d_1459_v13_ = _dafny.Map({p2: p1})
            d_1460_v14_: _dafny.Map
            d_1460_v14_ = _dafny.Map({d_1444_v1_: True})
            d_1461_v15_: _dafny.Seq
            d_1461_v15_ = _dafny.SeqWithoutIsStrInference([(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))], not((d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]), (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))], d_1444_v1_])
            d_1462_v16_: _dafny.Set
            d_1462_v16_ = _dafny.Set({d_1460_v14_, (d_1460_v14_ if (d_1461_v15_)[default__.safeIndex(len(d_1448_v5_), len(d_1461_v15_))] else d_1460_v14_), _dafny.Map({(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]: d_1444_v1_})})
            d_1463_v17_: _dafny.Seq
            d_1463_v17_ = _dafny.SeqWithoutIsStrInference([p2, 889])
            d_1464_v18_: _dafny.Set
            d_1464_v18_ = _dafny.Set({d_1444_v1_, d_1444_v1_})
            index171_ = default__.safeIndex(227, (p1).length(0))
            rhs177_ = (d_1461_v15_)[default__.safeIndex(default__.safeModuloInt(p2, len(d_1463_v17_)), len(d_1461_v15_))]
            rhs178_ = len(d_1464_v18_)
            rhs179_ = d_1459_v13_
            rhs180_ = d_1462_v16_
            rhs181_ = (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]
            lhs97_ = globalState
            lhs98_ = p1
            lhs99_ = default__.safeIndex(227, (p1).length(0))
            lhs100_ = globalState
            lhs97_.f8 = rhs177_
            lhs98_[lhs99_] = rhs178_
            d_1459_v13_ = rhs179_
            d_1462_v16_ = rhs180_
            lhs100_.f8 = rhs181_
            index172_ = default__.safeIndex(227, (p1).length(0))
            (p1)[index172_] = (d_1463_v17_)[default__.safeIndex(p2, len(d_1463_v17_))]
        elif True:
            d_1465_v19_: int
            d_1465_v19_ = 699
            d_1465_v19_ = 970
            d_1466_v20_: _dafny.Seq
            d_1466_v20_ = _dafny.SeqWithoutIsStrInference([d_1444_v1_])
            d_1467_v21_: D6
            d_1467_v21_ = D6_DC19(d_1444_v1_, d_1444_v1_, (d_1466_v20_)[default__.safeIndex(p2, len(d_1466_v20_))], d_1444_v1_, d_1466_v20_)
            d_1444_v1_ = not((d_1467_v21_).cf33)
            index173_ = default__.safeIndex(601, (d_1447_v4_).length(0))
            (d_1447_v4_)[index173_] = (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]
            if d_1444_v1_:
                d_1468_v22_: _dafny.Set
                d_1468_v22_ = _dafny.Set({d_1444_v1_, not((d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))])})
                d_1469_v23_: _dafny.Map
                d_1469_v23_ = _dafny.Map({(_dafny.Set({(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))], d_1444_v1_})) - (d_1468_v22_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaswmji"))})
                d_1469_v23_ = (d_1469_v23_).set((d_1468_v22_) - (d_1468_v22_), (_dafny.SeqWithoutIsStrInference([d_1446_v3_ for d_1470_i1_ in range(default__.abs(198))])) + (d_1448_v5_))
                index174_ = default__.safeIndex(557, (p1).length(0))
                (p1)[index174_] = d_1465_v19_
                index175_ = default__.safeIndex(557, (p1).length(0))
                (p1)[index175_] = p2
                d_1471_v24_: _dafny.Array
                nw220_ = _dafny.Array(D6.default()(), 19)
                d_1471_v24_ = nw220_
                d_1472_v25_: _dafny.MultiSet
                d_1472_v25_ = _dafny.MultiSet([d_1471_v24_])
                d_1473_v26_: D11
                d_1473_v26_ = D11_DC30(d_1472_v25_)
                rhs182_ = d_1446_v3_
                rhs183_ = (d_1472_v25_) | (((d_1472_v25_).set(d_1471_v24_, default__.abs(p2))).intersection((d_1473_v26_).cf55))
                d_1446_v3_ = rhs182_
                d_1472_v25_ = rhs183_
                index176_ = default__.safeIndex(601, (d_1447_v4_).length(0))
                (d_1447_v4_)[index176_] = (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]
                d_1474_v27_: _dafny.Map
                d_1474_v27_ = _dafny.Map({(p1)[default__.safeIndex(557, (p1).length(0))]: d_1446_v3_})
                d_1475_v28_: _dafny.Map
                d_1475_v28_ = _dafny.Map({(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]: len(_dafny.Set({len(d_1474_v27_), (p1)[default__.safeIndex(557, (p1).length(0))]}))})
                index177_ = default__.safeIndex(601, (d_1447_v4_).length(0))
                (d_1447_v4_)[index177_] = (len(d_1475_v28_)) == (d_1465_v19_)
            elif True:
                d_1476_v29_: _dafny.Map
                d_1476_v29_ = _dafny.Map({d_1448_v5_: (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]})
                d_1477_v30_: _dafny.Map
                d_1477_v30_ = _dafny.Map({((d_1476_v29_)[d_1448_v5_] if (d_1448_v5_) in (d_1476_v29_) else (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]): d_1465_v19_})
                d_1477_v30_ = d_1477_v30_
                d_1478_v31_: _dafny.Seq
                d_1478_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssylvffrk"))])
                d_1479_v32_: D0
                d_1479_v32_ = D0_DC1(p2, d_1448_v5_, p2, not(d_1444_v1_))
                d_1478_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1446_v3_ for d_1480_i2_ in range(default__.abs(-907))]), d_1448_v5_, (d_1479_v32_).cf2, d_1448_v5_, d_1448_v5_])
                d_1465_v19_ = d_1465_v19_
                d_1465_v19_ = p2
                index178_ = default__.safeIndex(532, (p1).length(0))
                (p1)[index178_] = d_1465_v19_
                d_1481_v33_: _dafny.Seq
                d_1481_v33_ = _dafny.SeqWithoutIsStrInference([d_1466_v20_])
                d_1482_v34_: _dafny.Seq
                d_1482_v34_ = _dafny.SeqWithoutIsStrInference([((p3)[d_1444_v1_] if (d_1444_v1_) in (p3) else d_1465_v19_)])
                d_1483_v35_: C2
                nw221_ = C2()
                nw221_.ctor__(p2, d_1448_v5_, p2, (d_1482_v34_)[default__.safeIndex(p2, len(d_1482_v34_))], d_1444_v1_)
                d_1483_v35_ = nw221_
                d_1484_v36_: _dafny.MultiSet
                d_1484_v36_ = _dafny.MultiSet([d_1483_v35_])
                index179_ = default__.safeIndex(532, (p1).length(0))
                rhs184_ = p2
                rhs185_ = (0) - ((p2) * ((len(d_1448_v5_)) + (len(d_1481_v33_))))
                rhs186_ = default__.safeDivisionInt(d_1465_v19_, default__.safeDivisionInt(p2, ((d_1484_v36_)[d_1483_v35_] if (d_1483_v35_) in (d_1484_v36_) else d_1465_v19_)))
                rhs187_ = ((p2) * (p2)) - (d_1465_v19_)
                lhs101_ = p1
                lhs102_ = default__.safeIndex(532, (p1).length(0))
                lhs101_[lhs102_] = rhs184_
                d_1465_v19_ = rhs185_
                d_1465_v19_ = rhs186_
                d_1465_v19_ = rhs187_
            index180_ = default__.safeIndex(681, (p1).length(0))
            (p1)[index180_] = p2
            d_1485_v37_: _dafny.Map
            d_1485_v37_ = _dafny.Map({(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]: (_dafny.MultiSet([d_1444_v1_, d_1444_v1_])).set(True, default__.abs(p2))})
            d_1486_v38_: D12
            d_1486_v38_ = D12_DC32(d_1485_v37_)
            index181_ = default__.safeIndex(681, (p1).length(0))
            rhs188_ = (d_1465_v19_) + (d_1465_v19_)
            rhs189_ = (d_1486_v38_).cf57
            lhs103_ = p1
            lhs104_ = default__.safeIndex(681, (p1).length(0))
            lhs103_[lhs104_] = rhs188_
            d_1485_v37_ = rhs189_
        d_1487_v39_: D1
        d_1487_v39_ = D1_DC4(_dafny.SeqWithoutIsStrInference([p1, p1]))
        d_1488_v40_: D1
        d_1488_v40_ = D1_DC6(d_1487_v39_)
        source24_ = d_1488_v40_
        if source24_.is_DC5:
            d_1489___mcc_h0_ = source24_.cf10
            d_1490___mcc_h1_ = source24_.cf11
            d_1491_cf11_ = d_1490___mcc_h1_
            d_1492_cf10_ = d_1489___mcc_h0_
            (globalState).f8 = d_1491_cf11_
            d_1493_v41_: _dafny.Map
            d_1493_v41_ = _dafny.Map({546: p2})
            d_1494_v43_: D3
            def iife148_():
                coll68_ = _dafny.Map()
                compr_68_: int
                for compr_68_ in (d_1445_v2_).Elements:
                    d_1495_v42_: int = compr_68_
                    if (d_1495_v42_) in (d_1445_v2_):
                        coll68_[(d_1495_v42_) * (p2)] = d_1493_v41_
                return _dafny.Map(coll68_)
            d_1494_v43_ = D3_DC9(p2, iife148_()
, len(d_1448_v5_), p2)
            d_1496_v44_: _dafny.MultiSet
            d_1496_v44_ = _dafny.MultiSet([D3_DC9(p2, _dafny.Map({p2: d_1493_v41_}), p2, p2), d_1494_v43_, d_1494_v43_])
            (globalState).f8 = (default__.fm52(globalState)).issubset(d_1496_v44_)
            d_1497_v45_: _dafny.Set
            d_1497_v45_ = _dafny.Set({d_1444_v1_})
            (globalState).f8 = not ((d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]) or ((d_1497_v45_).isdisjoint(d_1497_v45_))
            d_1498_v46_: int
            d_1498_v46_ = -948
            d_1498_v46_ = (p2 if d_1444_v1_ else p2)
        elif source24_.is_DC4:
            d_1499___mcc_h2_ = source24_.cf9
            d_1500_cf9_ = d_1499___mcc_h2_
            d_1501_v47_: C0
            nw222_ = C0()
            nw222_.ctor__(default__.safeModuloInt(p2, p2), d_1444_v1_, p2, 483)
            d_1501_v47_ = nw222_
            d_1502_v48_: _dafny.Map
            d_1502_v48_ = _dafny.Map({True: True})
            d_1503_v49_: C2
            nw223_ = C2()
            nw223_.ctor__(default__.safeDivisionInt(p2, p2), default__.fm36((d_1501_v47_).f28, (d_1501_v47_).f28, d_1446_v3_, globalState), (d_1501_v47_).f28, (0) - (p2), ((d_1502_v48_)[(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]] if ((d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]) in (d_1502_v48_) else d_1444_v1_))
            d_1503_v49_ = nw223_
            d_1444_v1_ = ((_dafny.MultiSet([(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))], d_1501_v47_.f29])) != (p3)) and (d_1501_v47_.f29)
            d_1504_v50_: _dafny.Map
            d_1504_v50_ = _dafny.Map({((d_1503_v49_).f26) + (732): default__.fm1(p2, (d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))], 469, globalState)})
            d_1504_v50_ = (d_1504_v50_).set(default__.safeDivisionInt((d_1501_v47_).f28, (d_1503_v49_).f26), ((d_1501_v47_).f28) <= (p2))
        elif True:
            d_1505___mcc_h3_ = source24_.cf12
            d_1506_cf12_ = d_1505___mcc_h3_
            d_1507_v51_: _dafny.Array
            nw224_ = _dafny.Array(_dafny.CodePoint('D'), 28)
            d_1507_v51_ = nw224_
            index182_ = default__.safeIndex(287, (d_1507_v51_).length(0))
            (d_1507_v51_)[index182_] = d_1446_v3_
            index183_ = default__.safeIndex(287, (d_1507_v51_).length(0))
            (d_1507_v51_)[index183_] = _dafny.CodePoint('v')
            d_1508_v52_: int
            d_1508_v52_ = 605
            d_1508_v52_ = p2
            d_1509_v53_: _dafny.Map
            d_1509_v53_ = _dafny.Map({False: not(d_1444_v1_)})
            d_1510_v54_: _dafny.Seq
            d_1510_v54_ = _dafny.SeqWithoutIsStrInference([(p2) + (p2), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivvcj"))), len(d_1509_v53_)), default__.safeModuloInt((0) - (p2), p2)])
            d_1511_v55_: _dafny.Set
            d_1511_v55_ = _dafny.Set({(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))], d_1444_v1_, d_1444_v1_})
            d_1508_v52_ = (d_1510_v54_)[default__.safeIndex(len((d_1511_v55_).intersection(_dafny.Set({d_1444_v1_}))), len(d_1510_v54_))]
            d_1511_v55_ = d_1511_v55_
        d_1448_v5_ = (((d_1448_v5_) + (d_1448_v5_)) + (d_1448_v5_)).set(default__.safeIndex((p2) - (p2), len(((d_1448_v5_) + (d_1448_v5_)) + (d_1448_v5_))), d_1446_v3_)
        d_1512_v56_: _dafny.Seq
        d_1512_v56_ = _dafny.SeqWithoutIsStrInference([(d_1447_v4_)[default__.safeIndex(601, (d_1447_v4_).length(0))]])
        d_1513_v57_: D0
        d_1513_v57_ = D0_DC1(default__.fm2(globalState), d_1448_v5_, p2, (d_1512_v56_)[default__.safeIndex(p2, len(d_1512_v56_))])
        r0 = (d_1447_v4_ if not((d_1513_v57_).cf4) else d_1447_v4_)
        return r0


class C7(T1):
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def fm7(self, globalState):
        return not(False)

    def fm8(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ce"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xawihm")))

    def m16(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        hi4_ = (self).f13
        for d_1514_i0_ in range((self).f13, hi4_):
            d_1515_v0_: _dafny.Map
            d_1515_v0_ = _dafny.Map({p2: ((self).f12) <= (default__.fm2(globalState))})
            d_1515_v0_ = d_1515_v0_
            d_1516_v1_: int
            d_1516_v1_ = 577
            d_1516_v1_ = (0) - (d_1516_v1_)
            d_1517_v2_: _dafny.Seq
            d_1517_v2_ = _dafny.SeqWithoutIsStrInference([p0, p0, True, True])
            index184_ = default__.safeIndex(272, (p2).length(0))
            (p2)[index184_] = not((not(p0)) == ((d_1517_v2_)[default__.safeIndex(d_1514_i0_, len(d_1517_v2_))]))
            index185_ = default__.safeIndex(272, (p2).length(0))
            (p2)[index185_] = not(p0)
            d_1518_v3_: _dafny.Array
            def lambda92_(d_1519_v1_, d_1520_p2_):
                def lambda93_(d_1521_i1_):
                    return ((D3_DC8(_dafny.Map({False: d_1519_v1_}))).cf14) | (_dafny.Map({(d_1520_p2_)[default__.safeIndex(272, (d_1520_p2_).length(0))]: (self).f12}))

                return lambda93_

            init49_ = lambda92_(d_1516_v1_, p2)
            nw225_ = _dafny.Array(None, 13)
            for i0_49_ in range(nw225_.length(0)):
                nw225_[i0_49_] = init49_(i0_49_)
            d_1518_v3_ = nw225_
            d_1522_v4_: _dafny.Map
            d_1522_v4_ = _dafny.Map({(p2)[default__.safeIndex(272, (p2).length(0))]: (0) - (p3)})
            index186_ = default__.safeIndex(874, (d_1518_v3_).length(0))
            (d_1518_v3_)[index186_] = d_1522_v4_
            index187_ = default__.safeIndex(874, (d_1518_v3_).length(0))
            (d_1518_v3_)[index187_] = (d_1522_v4_) | (d_1522_v4_)
        d_1523_i2_: int
        d_1523_i2_ = 0
        with _dafny.label("5"):
            while p0:
                with _dafny.c_label("5"):
                    if (d_1523_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_1523_i2_ = (d_1523_i2_) + (1)
                    d_1524_v5_: int
                    d_1524_v5_ = -181
                    d_1524_v5_ = default__.safeDivisionInt(p3, default__.fm2(globalState))
                    d_1525_v6_: _dafny.Map
                    d_1525_v6_ = _dafny.Map({p0: -437})
                    d_1526_v7_: D3
                    d_1526_v7_ = D3_DC8((d_1525_v6_).set(p0, -342))
                    d_1527_v8_: D3
                    d_1527_v8_ = D3_DC10(d_1526_v7_)
                    d_1528_v9_: D3
                    d_1528_v9_ = D3_DC10(d_1526_v7_)
                    d_1529_v10_: D3
                    d_1529_v10_ = D3_DC10(d_1527_v8_)
                    source25_ = d_1529_v10_
                    if source25_.is_DC9:
                        d_1530___mcc_h0_ = source25_.cf15
                        d_1531___mcc_h1_ = source25_.cf16
                        d_1532___mcc_h2_ = source25_.cf17
                        d_1533___mcc_h3_ = source25_.cf18
                        d_1534_cf18_ = d_1533___mcc_h3_
                        d_1535_cf17_ = d_1532___mcc_h2_
                        d_1536_cf16_ = d_1531___mcc_h1_
                        d_1537_cf15_ = d_1530___mcc_h0_
                        d_1538_v11_: D0
                        d_1538_v11_ = D0_DC1((0) - (p3), p1, (self).f13, p0)
                        d_1535_cf17_ = (d_1524_v5_) - ((0) - ((d_1538_v11_).cf3))
                        r1 = (p2 if True else p2)
                        d_1539_v12_: bool
                        d_1540_v13_: int
                        d_1541_v14_: _dafny.Seq
                        out19_: bool
                        out20_: int
                        out21_: _dafny.Seq
                        out19_, out20_, out21_ = (self).m17(((d_1525_v6_)[p0] if (p0) in (d_1525_v6_) else (self).f13), p0, globalState)
                        d_1539_v12_ = out19_
                        d_1540_v13_ = out20_
                        d_1541_v14_ = out21_
                        d_1542_v15_: D0
                        d_1542_v15_ = D0_DC3(801, p0)
                        d_1539_v12_ = (d_1542_v15_).cf8
                    elif source25_.is_DC8:
                        d_1543___mcc_h4_ = source25_.cf14
                        d_1544_cf14_ = d_1543___mcc_h4_
                        d_1545_v16_: _dafny.Seq
                        d_1545_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otl"))
                        d_1546_v17_: _dafny.Set
                        d_1546_v17_ = _dafny.Set({(self).f13})
                        d_1547_v18_: _dafny.Seq
                        d_1547_v18_ = _dafny.SeqWithoutIsStrInference([d_1524_v5_, -121, (0) - (-505), (self).f13, p3])
                        d_1548_v19_: str
                        d_1548_v19_ = _dafny.CodePoint('h')
                        d_1549_v20_: _dafny.Map
                        d_1549_v20_ = _dafny.Map({d_1548_v19_: p0})
                        d_1550_v22_: _dafny.Map
                        def iife149_():
                            coll69_ = _dafny.Set()
                            compr_69_: int
                            for compr_69_ in (d_1547_v18_).Elements:
                                d_1551_v21_: int = compr_69_
                                if (d_1551_v21_) in (d_1547_v18_):
                                    coll69_ = coll69_.union(_dafny.Set([(d_1551_v21_) * (len(_dafny.SeqWithoutIsStrInference([True])))]))
                            return _dafny.Set(coll69_)
                        d_1550_v22_ = _dafny.Map({d_1546_v17_: (self).fm8(len((d_1547_v18_)), d_1549_v20_, iife149_()
                        , globalState)})
                        d_1545_v16_ = ((d_1550_v22_)[d_1546_v17_] if (d_1546_v17_) in (d_1550_v22_) else (p1) + (p1))
                        d_1524_v5_ = (0) - ((self).f13)
                        d_1524_v5_ = (self).f12
                        index188_ = default__.safeIndex(73, (p2).length(0))
                        (p2)[index188_] = ((self).f13) == ((self).f13)
                        d_1552_v23_: _dafny.MultiSet
                        d_1552_v23_ = _dafny.MultiSet([p0])
                        d_1553_v24_: _dafny.Seq
                        d_1553_v24_ = _dafny.SeqWithoutIsStrInference([p0])
                        index189_ = default__.safeIndex(73, (p2).length(0))
                        (p2)[index189_] = (p0) not in ((d_1552_v23_) | (_dafny.MultiSet(d_1553_v24_)))
                    elif True:
                        d_1554___mcc_h5_ = source25_.cf19
                        d_1555_cf19_ = d_1554___mcc_h5_
                        d_1556_v25_: _dafny.Seq
                        d_1556_v25_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ob")))])
                        rhs190_ = (0) - ((self).f13)
                        rhs191_ = (d_1524_v5_) * ((self).f12)
                        rhs192_ = True
                        rhs193_ = not(p0)
                        rhs194_ = d_1556_v25_
                        lhs105_ = globalState
                        lhs106_ = globalState
                        lhs107_ = globalState
                        d_1524_v5_ = rhs190_
                        d_1524_v5_ = rhs191_
                        lhs105_.f8 = rhs192_
                        lhs106_.f8 = rhs193_
                        lhs107_.f5 = rhs194_
                        d_1524_v5_ = ((d_1556_v25_)[default__.safeIndex(p3, len(d_1556_v25_))]) * (p3)
                        d_1557_v26_: _dafny.Set
                        d_1557_v26_ = _dafny.Set({True, p0})
                        d_1558_v27_: _dafny.Seq
                        d_1558_v27_ = _dafny.SeqWithoutIsStrInference([p0, p0, (d_1557_v26_).issubset(d_1557_v26_)])
                        rhs195_ = p0
                        rhs196_ = p0
                        rhs197_ = default__.safeModuloInt((self).f13, (self).f12)
                        rhs198_ = (0) - ((0) - (p3))
                        rhs199_ = len(d_1558_v27_)
                        lhs108_ = globalState
                        lhs109_ = globalState
                        lhs108_.f8 = rhs195_
                        lhs109_.f8 = rhs196_
                        d_1524_v5_ = rhs197_
                        d_1524_v5_ = rhs198_
                        d_1524_v5_ = rhs199_
                        (globalState).f8 = p0
                    d_1559_v29_: _dafny.Map
                    def iife150_():
                        coll70_ = _dafny.Map()
                        compr_70_: int
                        for compr_70_ in _dafny.IntegerRange(615, -716):
                            d_1560_v28_: int = compr_70_
                            if ((615) <= (d_1560_v28_)) and ((d_1560_v28_) < (-716)):
                                coll70_[(d_1560_v28_) - (p3)] = (self).f12
                        return _dafny.Map(coll70_)
                    d_1559_v29_ = _dafny.Map({p3: iife150_()
                    })
                    d_1561_v30_: _dafny.Set
                    d_1561_v30_ = _dafny.Set({(self).f13, default__.fm2(globalState), p3, (self).f12})
                    d_1562_v31_: D3
                    d_1562_v31_ = D3_DC9((self).f12, d_1559_v29_, d_1524_v5_, len(d_1561_v30_))
                    d_1563_v32_: _dafny.Seq
                    d_1563_v32_ = _dafny.SeqWithoutIsStrInference([d_1524_v5_, p3, (d_1562_v31_).cf17, default__.fm2(globalState)])
                    (globalState).f5 = ((d_1563_v32_) + (_dafny.SeqWithoutIsStrInference([(self).f13, (self).f12]))) + (d_1563_v32_)
                    (globalState).f8 = p0
                    pass
            pass
        d_1564_v33_: D1
        d_1564_v33_ = D1_DC5(_dafny.MultiSet([(self).f13]), p0)
        d_1565_v34_: D1
        d_1565_v34_ = D1_DC6(d_1564_v33_)
        d_1566_v35_: D1
        d_1566_v35_ = D1_DC6(d_1565_v34_)
        d_1566_v35_ = d_1566_v35_
        d_1567_v36_: _dafny.Seq
        d_1567_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "paujkw"))
        d_1567_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hy"))
        d_1568_v37_: _dafny.Array
        d_1569_v38_: bool
        out22_: _dafny.Array
        out23_: bool
        out22_, out23_ = default__.m0(globalState)
        d_1568_v37_ = out22_
        d_1569_v38_ = out23_
        d_1570_v40_: _dafny.Seq
        d_1570_v40_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        def iife151_():
            coll71_ = _dafny.Map()
            compr_71_: int
            for compr_71_ in (d_1570_v40_).Elements:
                d_1571_v39_: int = compr_71_
                if (d_1571_v39_) in (d_1570_v40_):
                    coll71_[default__.safeDivisionInt(d_1571_v39_, 285)] = p3
            return _dafny.Map(coll71_)
        hi5_ = len(default__.fm22(iife151_()
        , globalState))
        for d_1572_i3_ in range((self).f13, hi5_):
            (globalState).f8 = not(p0)
            d_1573_v41_: int
            d_1573_v41_ = 905
            d_1573_v41_ = d_1573_v41_
            d_1574_v42_: _dafny.MultiSet
            d_1574_v42_ = _dafny.MultiSet([d_1573_v41_])
            rhs200_ = d_1569_v38_
            rhs201_ = (0) - ((self).f12)
            rhs202_ = (not(p0) if (_dafny.MultiSet([(self).f12])).isdisjoint(d_1574_v42_) else not(p0))
            rhs203_ = (p1) + (p1)
            lhs110_ = globalState
            lhs110_.f8 = rhs200_
            d_1573_v41_ = rhs201_
            d_1569_v38_ = rhs202_
            d_1567_v36_ = rhs203_
            d_1573_v41_ = (((self).f12) + (len(p1))) + (d_1572_i3_)
        d_1575_v43_: _dafny.Seq
        d_1575_v43_ = _dafny.SeqWithoutIsStrInference([True])
        d_1576_v44_: _dafny.MultiSet
        d_1576_v44_ = _dafny.MultiSet([p3, p3])
        d_1577_v45_: _dafny.MultiSet
        d_1577_v45_ = _dafny.MultiSet([(D0_DC1((d_1576_v44_).cardinality, d_1567_v36_, 485, not(False))).cf1, p3])
        d_1578_v46_: _dafny.Array
        nw226_ = _dafny.Array(None, 14)
        nw226_[int(0)] = p3
        nw226_[int(1)] = p3
        nw226_[int(2)] = len(d_1575_v43_)
        nw226_[int(3)] = ((self).f12 if d_1569_v38_ else p3)
        nw226_[int(4)] = (self).f13
        nw226_[int(5)] = ((D1_DC5((d_1576_v44_).set(p3, default__.abs((self).f12)), d_1569_v38_)).cf10).cardinality
        nw226_[int(6)] = p3
        nw226_[int(7)] = -5
        nw226_[int(8)] = (self).f13
        nw226_[int(9)] = (p3) + (90)
        nw226_[int(10)] = (self).f13
        nw226_[int(11)] = p3
        nw226_[int(12)] = ((self).f12) + ((self).f13)
        nw226_[int(13)] = p3
        d_1578_v46_ = nw226_
        r0 = d_1578_v46_
        r1 = p2
        return r0, r1

    def m17(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1579_v0_: _dafny.Map
        d_1579_v0_ = _dafny.Map({(self).f13: p0})
        d_1580_v1_: _dafny.Map
        d_1580_v1_ = _dafny.Map({p0: d_1579_v0_})
        d_1581_v3_: str
        d_1581_v3_ = _dafny.CodePoint('a')
        d_1582_v4_: _dafny.Map
        d_1582_v4_ = _dafny.Map({d_1581_v3_: p1})
        d_1583_v5_: _dafny.Set
        d_1583_v5_ = _dafny.Set({p0})
        d_1584_v6_: _dafny.Seq
        d_1584_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agru"))
        d_1585_v7_: _dafny.Map
        def iife152_():
            coll72_ = _dafny.Map()
            compr_72_: int
            for compr_72_ in _dafny.IntegerRange(243, 483):
                d_1586_v2_: int = compr_72_
                if ((243) <= (d_1586_v2_)) and ((d_1586_v2_) < (483)):
                    coll72_[(d_1586_v2_) - ((self).f12)] = (self).f13
            return _dafny.Map(coll72_)
        d_1585_v7_ = _dafny.Map({_dafny.Set({d_1579_v0_, ((d_1580_v1_)[p0] if (p0) in (d_1580_v1_) else iife152_()
        )}): ((self).fm8((self).f13, d_1582_v4_, d_1583_v5_, globalState)) < (d_1584_v6_)})
        d_1587_v8_: _dafny.Set
        d_1587_v8_ = _dafny.Set({(d_1579_v0_).set((self).f12, (0) - ((self).f12))})
        d_1585_v7_ = (d_1585_v7_).set((d_1587_v8_) | (default__.fm23(p1, -496, globalState)), True)
        d_1588_v9_: _dafny.Seq
        d_1588_v9_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        d_1589_v10_: _dafny.Seq
        d_1589_v10_ = (d_1588_v9_) + (_dafny.SeqWithoutIsStrInference([897, 70, 924]))
        source26_ = d_1589_v10_
        d_1590___mcc_h0_ = source26_
        d_1591_cf13_ = d_1590___mcc_h0_
        d_1579_v0_ = (d_1579_v0_).set((self).f13, (self).f12)
        if p1:
            r0 = p1
            r1 = len(d_1591_cf13_)
            r1 = (self).f13
            def iife153_():
                coll73_ = _dafny.Set()
                compr_73_: int
                for compr_73_ in _dafny.IntegerRange(524, 131):
                    d_1592_v11_: int = compr_73_
                    if ((524) <= (d_1592_v11_)) and ((d_1592_v11_) < (131)):
                        coll73_ = coll73_.union(_dafny.Set([(d_1592_v11_) * ((self).f13)]))
                return _dafny.Set(coll73_)
            r1 = len(iife153_()
            )
            d_1593_v12_: C2
            nw227_ = C2()
            nw227_.ctor__((self).f12, default__.fm36((self).f12, 804, _dafny.CodePoint('d'), globalState), len(d_1584_v6_), (self).f12, (d_1583_v5_).issubset(d_1583_v5_))
            d_1593_v12_ = nw227_
        elif True:
            d_1594_v13_: _dafny.Map
            d_1594_v13_ = _dafny.Map({((d_1588_v9_).set(default__.safeIndex(645, len(d_1588_v9_)), (self).f13)).set(default__.safeIndex((self).f13, len((d_1588_v9_).set(default__.safeIndex(645, len(d_1588_v9_)), (self).f13))), len(d_1584_v6_)): p1})
            d_1594_v13_ = d_1594_v13_
            r1 = default__.safeDivisionInt(p0, p0)
            (globalState).f8 = ((len(_dafny.Set({d_1581_v3_, d_1581_v3_})) if p1 else (self).f12)) != ((self).f12)
            d_1595_v14_: _dafny.Array
            nw228_ = _dafny.Array(int(0), 1)
            d_1595_v14_ = nw228_
            index190_ = default__.safeIndex(453, (d_1595_v14_).length(0))
            (d_1595_v14_)[index190_] = len(default__.fm0(globalState))
            index191_ = default__.safeIndex(453, (d_1595_v14_).length(0))
            (d_1595_v14_)[index191_] = (self).f13
            d_1596_v15_: _dafny.Array
            nw229_ = _dafny.Array(D3.default()(), 11)
            d_1596_v15_ = nw229_
            d_1597_v16_: _dafny.Map
            d_1597_v16_ = _dafny.Map({p1: (self).f13})
            d_1598_v17_: D3
            d_1598_v17_ = D3_DC8(d_1597_v16_)
            index192_ = default__.safeIndex(772, (d_1596_v15_).length(0))
            (d_1596_v15_)[index192_] = d_1598_v17_
            index193_ = default__.safeIndex(772, (d_1596_v15_).length(0))
            (d_1596_v15_)[index193_] = d_1598_v17_
        (globalState).f8 = p1
        d_1599_v18_: _dafny.Map
        d_1599_v18_ = _dafny.Map({p1: default__.fm36(p0, 854, d_1581_v3_, globalState)})
        d_1600_v19_: C2
        nw230_ = C2()
        nw230_.ctor__(default__.safeModuloInt((0) - ((self).f12), (self).f12), ((d_1599_v18_)[p1] if (p1) in (d_1599_v18_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oblpvb"))), p0, (self).f13, p1)
        d_1600_v19_ = nw230_
        d_1601_v20_: _dafny.Map
        d_1601_v20_ = _dafny.Map({(self).f13: p1})
        def iife154_():
            coll74_ = _dafny.Map()
            compr_74_: int
            for compr_74_ in (d_1601_v20_).keys.Elements:
                d_1602_v21_: int = compr_74_
                if (d_1602_v21_) in (d_1601_v20_):
                    coll74_[(d_1602_v21_) * ((self).f12)] = p1
            return _dafny.Map(coll74_)
        d_1601_v20_ = (iife154_()
        ) | (_dafny.Map({871: p1}))
        r1 = p0
        d_1603_v22_: _dafny.Array
        nw231_ = _dafny.Array(D7.default()(), 1)
        d_1603_v22_ = nw231_
        rhs204_ = (self).f13
        rhs205_ = (self).f13
        rhs206_ = (not (not(p1)) or (p1) if not(p1) else p1)
        rhs207_ = d_1603_v22_
        rhs208_ = (d_1581_v3_) in (d_1584_v6_)
        lhs111_ = globalState
        r1 = rhs204_
        r1 = rhs205_
        lhs111_.f8 = rhs206_
        d_1603_v22_ = rhs207_
        r0 = rhs208_
        d_1604_v23_: _dafny.Array
        nw232_ = _dafny.Array(None, 3)
        nw232_[int(0)] = not (p1) or (p1)
        nw232_[int(1)] = p1
        nw232_[int(2)] = p1
        d_1604_v23_ = nw232_
        d_1604_v23_ = d_1604_v23_
        d_1605_v24_: _dafny.Seq
        d_1605_v24_ = _dafny.SeqWithoutIsStrInference([p1])
        r0 = (d_1605_v24_)[default__.safeIndex((self).f12, len(d_1605_v24_))]
        r1 = p0
        r2 = d_1584_v6_
        return r0, r1, r2


class C8:
    def  __init__(self):
        self._f23: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f23):
        (self)._f23 = f23

    def fm21(self, p0, p1, globalState):
        return D0_DC3((len(_dafny.Set({False, True}))) * (18), (len(_dafny.Map({_dafny.Map({297: (_dafny.MultiSet([492])).cardinality}): 144}))) != (-162))

    def m14(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: D0 = D0.default()()
        d_1606_v0_: bool
        d_1606_v0_ = True
        (globalState).f8 = d_1606_v0_
        d_1607_v1_: int
        d_1607_v1_ = 556
        d_1607_v1_ = d_1607_v1_
        d_1608_i0_: int
        d_1608_i0_ = 0
        with _dafny.label("6"):
            while default__.fm1((0) - (d_1607_v1_), d_1606_v0_, (0) - (d_1607_v1_), globalState):
                with _dafny.c_label("6"):
                    if (d_1608_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1608_i0_ = (d_1608_i0_) + (1)
                    d_1609_v2_: _dafny.Seq
                    d_1609_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1606_v0_, not(d_1606_v0_)])])
                    d_1610_v3_: _dafny.Seq
                    d_1610_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypssnf"))
                    d_1611_v4_: C0
                    nw233_ = C0()
                    nw233_.ctor__(((0) - (d_1607_v1_)) + (len(default__.fm36(d_1607_v1_, d_1607_v1_, default__.fm45(len(d_1609_v2_), globalState), globalState))), d_1606_v0_, default__.safeDivisionInt(len(d_1610_v3_), d_1607_v1_), d_1607_v1_)
                    d_1611_v4_ = nw233_
                    d_1612_v5_: _dafny.Set
                    d_1612_v5_ = _dafny.Set({d_1611_v4_.f29, d_1606_v0_})
                    d_1613_v6_: _dafny.MultiSet
                    d_1613_v6_ = _dafny.MultiSet([len(d_1612_v5_)])
                    d_1614_v7_: D1
                    d_1614_v7_ = D1_DC5(d_1613_v6_, d_1606_v0_)
                    d_1615_v8_: _dafny.Array
                    nw234_ = _dafny.Array(None, 3)
                    nw234_[int(0)] = _dafny.MultiSet([d_1607_v1_])
                    nw234_[int(1)] = (d_1614_v7_).cf10
                    nw234_[int(2)] = d_1613_v6_
                    d_1615_v8_ = nw234_
                    d_1616_v9_: _dafny.Map
                    d_1616_v9_ = _dafny.Map({d_1615_v8_: not(d_1611_v4_.f29)})
                    d_1616_v9_ = (d_1616_v9_).set(d_1615_v8_, d_1606_v0_)
                    d_1617_v10_: _dafny.MultiSet
                    d_1617_v10_ = _dafny.MultiSet([d_1606_v0_, d_1611_v4_.f29, False])
                    rhs209_ = d_1606_v0_
                    rhs210_ = d_1610_v3_
                    rhs211_ = ((d_1611_v4_).f28) - (((0) - (d_1607_v1_)) - ((d_1611_v4_).f28))
                    rhs212_ = d_1606_v0_
                    rhs213_ = d_1617_v10_
                    lhs112_ = globalState
                    d_1606_v0_ = rhs209_
                    d_1610_v3_ = rhs210_
                    d_1607_v1_ = rhs211_
                    lhs112_.f8 = rhs212_
                    d_1617_v10_ = rhs213_
                    d_1607_v1_ = (d_1611_v4_).f28
                    pass
            pass
        d_1618_v11_: _dafny.Map
        d_1618_v11_ = _dafny.Map({d_1606_v0_: d_1606_v0_})
        d_1619_v12_: _dafny.MultiSet
        d_1619_v12_ = _dafny.MultiSet([d_1618_v11_, d_1618_v11_])
        d_1620_v13_: _dafny.MultiSet
        d_1620_v13_ = d_1619_v12_
        d_1621_v14_: _dafny.Map
        d_1621_v14_ = _dafny.Map({d_1620_v13_: _dafny.CodePoint('r')})
        d_1622_v15_: _dafny.Set
        d_1622_v15_ = _dafny.Set({d_1621_v14_})
        d_1622_v15_ = (d_1622_v15_).intersection((_dafny.Set({d_1621_v14_}) if d_1606_v0_ else d_1622_v15_))
        d_1623_i1_: int
        d_1623_i1_ = 0
        with _dafny.label("7"):
            while d_1606_v0_:
                with _dafny.c_label("7"):
                    if (d_1623_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_1623_i1_ = (d_1623_i1_) + (1)
                    d_1624_v16_: _dafny.Array
                    nw235_ = _dafny.Array(_dafny.Map({}), 7)
                    d_1624_v16_ = nw235_
                    d_1625_v17_: _dafny.Array
                    def lambda94_(d_1626_v0_):
                        def lambda95_(d_1627_i2_):
                            return d_1626_v0_

                        return lambda95_

                    init50_ = lambda94_(d_1606_v0_)
                    nw236_ = _dafny.Array(None, 6)
                    for i0_50_ in range(nw236_.length(0)):
                        nw236_[i0_50_] = init50_(i0_50_)
                    d_1625_v17_ = nw236_
                    d_1628_v18_: _dafny.Map
                    d_1628_v18_ = _dafny.Map({d_1625_v17_: d_1606_v0_})
                    index194_ = default__.safeIndex(525, (d_1624_v16_).length(0))
                    (d_1624_v16_)[index194_] = d_1628_v18_
                    index195_ = default__.safeIndex(525, (d_1624_v16_).length(0))
                    (d_1624_v16_)[index195_] = _dafny.Map({d_1625_v17_: d_1606_v0_})
                    if ((d_1607_v1_) == (d_1607_v1_)) and (d_1606_v0_):
                        d_1607_v1_ = ((self).f23)[default__.safeIndex(267, len((self).f23))]
                        d_1629_v19_: str
                        d_1629_v19_ = _dafny.CodePoint('v')
                        d_1630_v20_: _dafny.Seq
                        d_1630_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jia"))
                        d_1631_v21_: _dafny.Array
                        def lambda96_(d_1632_v1_):
                            def lambda97_(d_1633_i3_):
                                return (d_1633_i3_) + (d_1632_v1_)

                            return lambda97_

                        init51_ = lambda96_(d_1607_v1_)
                        nw237_ = _dafny.Array(None, 23)
                        for i0_51_ in range(nw237_.length(0)):
                            nw237_[i0_51_] = init51_(i0_51_)
                        d_1631_v21_ = nw237_
                        rhs214_ = ((d_1629_v19_ if True else d_1629_v19_)) not in (d_1630_v20_)
                        rhs215_ = d_1606_v0_
                        rhs216_ = d_1631_v21_
                        rhs217_ = d_1607_v1_
                        lhs113_ = globalState
                        r0 = rhs214_
                        lhs113_.f8 = rhs215_
                        r1 = rhs216_
                        d_1607_v1_ = rhs217_
                        d_1607_v1_ = d_1607_v1_
                        d_1634_v22_: _dafny.Seq
                        d_1634_v22_ = _dafny.SeqWithoutIsStrInference([d_1606_v0_, d_1606_v0_])
                        d_1635_v23_: T2
                        nw238_ = C2()
                        nw238_.ctor__(d_1607_v1_, d_1630_v20_, len(d_1634_v22_), (d_1607_v1_ if True else d_1607_v1_), not(d_1606_v0_))
                        d_1635_v23_ = nw238_
                        rhs218_ = d_1635_v23_
                        rhs219_ = d_1606_v0_
                        lhs114_ = globalState
                        d_1635_v23_ = rhs218_
                        lhs114_.f8 = rhs219_
                        d_1636_v24_: C6
                        nw239_ = C6()
                        nw239_.ctor__()
                        d_1636_v24_ = nw239_
                    elif True:
                        rhs220_ = (d_1606_v0_ if d_1606_v0_ else d_1606_v0_)
                        rhs221_ = d_1607_v1_
                        lhs115_ = globalState
                        lhs115_.f8 = rhs220_
                        d_1607_v1_ = rhs221_
                        d_1637_v25_: _dafny.Seq
                        d_1637_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxdmehy"))
                        d_1638_v26_: str
                        d_1638_v26_ = _dafny.CodePoint('t')
                        d_1639_v27_: _dafny.Map
                        d_1639_v27_ = _dafny.Map({True: d_1637_v25_})
                        d_1640_v28_: _dafny.Map
                        d_1640_v28_ = _dafny.Map({d_1638_v26_: d_1607_v1_})
                        d_1641_v29_: _dafny.Seq
                        d_1641_v29_ = _dafny.SeqWithoutIsStrInference([False])
                        d_1642_v30_: _dafny.Array
                        nw240_ = _dafny.Array(None, 21)
                        nw240_[int(0)] = (d_1637_v25_) + (d_1637_v25_)
                        nw240_[int(1)] = d_1637_v25_
                        nw240_[int(2)] = ((d_1637_v25_).set(default__.safeIndex(d_1607_v1_, len(d_1637_v25_)), d_1638_v26_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyrd")))
                        nw240_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_1638_v26_ for d_1643_i4_ in range(default__.abs(731))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1644_i5_ in range(default__.abs(-275))]))
                        nw240_[int(4)] = (d_1637_v25_) + (d_1637_v25_)
                        nw240_[int(5)] = d_1637_v25_
                        nw240_[int(6)] = d_1637_v25_
                        nw240_[int(7)] = ((d_1639_v27_)[d_1606_v0_] if (d_1606_v0_) in (d_1639_v27_) else _dafny.SeqWithoutIsStrInference([d_1638_v26_ for d_1645_i6_ in range(default__.abs(34))]))
                        nw240_[int(8)] = (d_1637_v25_) + (d_1637_v25_)
                        nw240_[int(9)] = (d_1637_v25_) + (d_1637_v25_)
                        nw240_[int(10)] = d_1637_v25_
                        nw240_[int(11)] = d_1637_v25_
                        nw240_[int(12)] = default__.fm36(len(d_1640_v28_), d_1607_v1_, default__.fm45(d_1607_v1_, globalState), globalState)
                        nw240_[int(13)] = d_1637_v25_
                        nw240_[int(14)] = d_1637_v25_
                        nw240_[int(15)] = (d_1637_v25_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrlgffk")))
                        nw240_[int(16)] = d_1637_v25_
                        nw240_[int(17)] = default__.fm36(d_1607_v1_, len(d_1641_v29_), d_1638_v26_, globalState)
                        nw240_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_1638_v26_ for d_1646_i7_ in range(default__.abs(-427))])) + (d_1637_v25_)
                        nw240_[int(19)] = d_1637_v25_
                        nw240_[int(20)] = d_1637_v25_
                        d_1642_v30_ = nw240_
                        d_1642_v30_ = d_1642_v30_
                        d_1647_v31_: D6
                        d_1647_v31_ = D6_DC19(d_1606_v0_, d_1606_v0_, d_1606_v0_, d_1606_v0_, d_1641_v29_)
                        index196_ = default__.safeIndex(447, (d_1625_v17_).length(0))
                        (d_1625_v17_)[index196_] = (d_1641_v29_) == ((d_1647_v31_).cf37)
                        d_1648_v32_: _dafny.Array
                        nw241_ = _dafny.Array(int(0), 16)
                        d_1648_v32_ = nw241_
                        index197_ = default__.safeIndex(858, (d_1648_v32_).length(0))
                        (d_1648_v32_)[index197_] = d_1607_v1_
                        d_1649_v33_: _dafny.Set
                        d_1649_v33_ = _dafny.Set({d_1607_v1_, d_1607_v1_, d_1607_v1_})
                        d_1650_v34_: _dafny.Set
                        d_1650_v34_ = _dafny.Set({d_1649_v33_, d_1649_v33_})
                        d_1651_v35_: _dafny.Seq
                        d_1651_v35_ = _dafny.SeqWithoutIsStrInference([d_1650_v34_])
                        d_1652_v36_: _dafny.MultiSet
                        d_1652_v36_ = _dafny.MultiSet([len(d_1641_v29_)])
                        d_1653_v37_: _dafny.Map
                        d_1653_v37_ = _dafny.Map({d_1606_v0_: d_1607_v1_})
                        d_1654_v38_: _dafny.Map
                        d_1654_v38_ = _dafny.Map({d_1607_v1_: default__.fm1(d_1607_v1_, d_1606_v0_, len(d_1653_v37_), globalState)})
                        index198_ = default__.safeIndex(447, (d_1625_v17_).length(0))
                        index199_ = default__.safeIndex(858, (d_1648_v32_).length(0))
                        rhs222_ = _dafny.CodePoint('l')
                        rhs223_ = (((d_1624_v16_)[default__.safeIndex(525, (d_1624_v16_).length(0))])[d_1625_v17_] if (d_1625_v17_) in ((d_1624_v16_)[default__.safeIndex(525, (d_1624_v16_).length(0))]) else d_1606_v0_)
                        rhs224_ = default__.fm1(d_1607_v1_, d_1606_v0_, (0) - (len((d_1651_v35_)[default__.safeIndex(d_1607_v1_, len(d_1651_v35_))])), globalState)
                        rhs225_ = (_dafny.CodePoint('i') if True else (d_1638_v26_ if default__.fm1((d_1652_v36_).cardinality, ((d_1654_v38_)[d_1607_v1_] if (d_1607_v1_) in (d_1654_v38_) else d_1606_v0_), d_1607_v1_, globalState) else d_1638_v26_))
                        rhs226_ = d_1607_v1_
                        lhs116_ = d_1625_v17_
                        lhs117_ = default__.safeIndex(447, (d_1625_v17_).length(0))
                        lhs118_ = globalState
                        lhs119_ = d_1648_v32_
                        lhs120_ = default__.safeIndex(858, (d_1648_v32_).length(0))
                        d_1638_v26_ = rhs222_
                        lhs116_[lhs117_] = rhs223_
                        lhs118_.f8 = rhs224_
                        d_1638_v26_ = rhs225_
                        lhs119_[lhs120_] = rhs226_
                        d_1655_v39_: _dafny.Map
                        d_1655_v39_ = _dafny.Map({(d_1648_v32_)[default__.safeIndex(858, (d_1648_v32_).length(0))]: d_1638_v26_})
                        d_1656_v40_: C7
                        nw242_ = C7()
                        nw242_.ctor__((d_1648_v32_)[default__.safeIndex(858, (d_1648_v32_).length(0))], ((self).f23)[default__.safeIndex(len(d_1655_v39_), len((self).f23))])
                        d_1656_v40_ = nw242_
                        index200_ = default__.safeIndex(858, (d_1648_v32_).length(0))
                        (d_1648_v32_)[index200_] = (d_1648_v32_)[default__.safeIndex(858, (d_1648_v32_).length(0))]
                    if d_1606_v0_:
                        d_1657_v41_: _dafny.MultiSet
                        d_1657_v41_ = _dafny.MultiSet([d_1606_v0_, not(d_1606_v0_), d_1606_v0_])
                        d_1658_v42_: D7
                        d_1658_v42_ = D7_DC21(d_1657_v41_)
                        d_1659_v43_: D7
                        d_1659_v43_ = D7_DC24(d_1658_v42_)
                        d_1660_v44_: _dafny.Map
                        d_1660_v44_ = _dafny.Map({not(d_1606_v0_): d_1658_v42_})
                        pat_let_tv59_ = d_1658_v42_
                        def iife155_(_pat_let40_0):
                            def iife156_(d_1661_dt__update__tmp_h0_):
                                def iife157_(_pat_let41_0):
                                    def iife158_(d_1662_dt__update_hcf46_h0_):
                                        return D7_DC24(d_1662_dt__update_hcf46_h0_)
                                    return iife158_(_pat_let41_0)
                                return iife157_(pat_let_tv59_)
                            return iife156_(_pat_let40_0)
                        d_1659_v43_ = iife155_(D7_DC24(((d_1660_v44_)[d_1606_v0_] if (d_1606_v0_) in (d_1660_v44_) else d_1658_v42_)))
                        d_1663_v45_: _dafny.Seq
                        d_1663_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgvpeq"))
                        d_1664_v46_: T0
                        nw243_ = C2()
                        nw243_.ctor__(d_1607_v1_, d_1663_v45_, d_1607_v1_, d_1607_v1_, d_1606_v0_)
                        d_1664_v46_ = nw243_
                        d_1665_v47_: _dafny.Seq
                        d_1665_v47_ = _dafny.SeqWithoutIsStrInference([d_1664_v46_, d_1664_v46_])
                        d_1666_v48_: _dafny.Array
                        nw244_ = _dafny.Array(None, 6)
                        nw244_[int(0)] = d_1607_v1_
                        nw244_[int(1)] = d_1607_v1_
                        nw244_[int(2)] = len(d_1665_v47_)
                        nw244_[int(3)] = (0) - (d_1607_v1_)
                        nw244_[int(4)] = d_1607_v1_
                        nw244_[int(5)] = d_1607_v1_
                        d_1666_v48_ = nw244_
                        d_1667_v49_: _dafny.Map
                        d_1667_v49_ = _dafny.Map({d_1666_v48_: -758})
                        d_1667_v49_ = (d_1667_v49_).set(d_1666_v48_, len(d_1663_v45_))
                        d_1668_v50_: str
                        d_1668_v50_ = _dafny.CodePoint('x')
                        d_1669_v51_: _dafny.Map
                        d_1669_v51_ = _dafny.Map({d_1663_v45_: d_1668_v50_})
                        d_1669_v51_ = (d_1669_v51_).set((d_1663_v45_) + (d_1663_v45_), d_1668_v50_)
                        d_1606_v0_ = d_1606_v0_
                        d_1670_v52_: _dafny.MultiSet
                        d_1670_v52_ = _dafny.MultiSet([d_1607_v1_, d_1607_v1_, d_1607_v1_, d_1607_v1_])
                        d_1670_v52_ = d_1670_v52_
                    elif True:
                        (globalState).f8 = d_1606_v0_
                        d_1671_v53_: C5
                        nw245_ = C5()
                        nw245_.ctor__(((0) - (default__.fm2(globalState)) if default__.fm1(d_1607_v1_, True, default__.fm2(globalState), globalState) else len((self).f23)))
                        d_1671_v53_ = nw245_
                        d_1672_v54_: _dafny.Map
                        d_1672_v54_ = _dafny.Map({d_1606_v0_: _dafny.SeqWithoutIsStrInference([d_1606_v0_])})
                        d_1673_v55_: str
                        d_1673_v55_ = _dafny.CodePoint('q')
                        d_1674_v56_: _dafny.Map
                        d_1674_v56_ = _dafny.Map({d_1673_v55_: d_1607_v1_})
                        d_1675_v57_: _dafny.MultiSet
                        d_1675_v57_ = _dafny.MultiSet([d_1674_v56_])
                        d_1676_v58_: _dafny.Seq
                        d_1676_v58_ = _dafny.SeqWithoutIsStrInference([False])
                        d_1607_v1_ = (0) - (len(((d_1672_v54_)[not((_dafny.MultiSet([d_1674_v56_, d_1674_v56_])).issubset(d_1675_v57_))] if (not((_dafny.MultiSet([d_1674_v56_, d_1674_v56_])).issubset(d_1675_v57_))) in (d_1672_v54_) else (d_1676_v58_) + (d_1676_v58_))))
                        d_1607_v1_ = (d_1671_v53_).f24
                        d_1607_v1_ = (d_1671_v53_).f24
                    d_1607_v1_ = default__.fm2(globalState)
                    pass
            pass
        d_1677_v59_: _dafny.MultiSet
        d_1677_v59_ = _dafny.MultiSet([d_1606_v0_, d_1606_v0_, False, d_1606_v0_])
        d_1678_v60_: D4
        d_1678_v60_ = D4_DC13(758)
        pat_let_tv60_ = d_1677_v59_
        pat_let_tv61_ = d_1677_v59_
        pat_let_tv62_ = d_1677_v59_
        def lambda98_(source27_):
            if source27_.is_DC12:
                d_1679___mcc_h0_ = source27_.cf21
                d_1680___mcc_h1_ = source27_.cf22
                d_1681___mcc_h2_ = source27_.cf23
                d_1682_cf23_ = d_1681___mcc_h2_
                d_1683_cf22_ = d_1680___mcc_h1_
                d_1684_cf21_ = d_1679___mcc_h0_
                return pat_let_tv60_
            elif source27_.is_DC13:
                d_1685___mcc_h3_ = source27_.cf24
                d_1686_cf24_ = d_1685___mcc_h3_
                return pat_let_tv61_
            elif True:
                d_1687___mcc_h4_ = source27_.cf20
                d_1688_cf20_ = d_1687___mcc_h4_
                return pat_let_tv62_

        d_1677_v59_ = lambda98_(d_1678_v60_)
        r0 = d_1606_v0_
        d_1689_v61_: _dafny.Array
        nw246_ = _dafny.Array(int(0), 21)
        d_1689_v61_ = nw246_
        r1 = d_1689_v61_
        d_1690_v62_: D0
        d_1690_v62_ = D0_DC3(d_1607_v1_, d_1606_v0_)
        d_1691_v63_: _dafny.Seq
        d_1691_v63_ = _dafny.SeqWithoutIsStrInference([d_1606_v0_])
        pat_let_tv63_ = d_1677_v59_
        pat_let_tv64_ = d_1677_v59_
        pat_let_tv65_ = d_1691_v63_
        pat_let_tv66_ = d_1607_v1_
        pat_let_tv67_ = d_1691_v63_
        pat_let_tv68_ = d_1606_v0_
        def iife159_(_pat_let42_0):
            def iife160_(d_1692_dt__update__tmp_h1_):
                def iife161_(_pat_let43_0):
                    def iife162_(d_1693_dt__update_hcf7_h0_):
                        return D0_DC3(d_1693_dt__update_hcf7_h0_, (d_1692_dt__update__tmp_h1_).cf8)
                    return iife162_(_pat_let43_0)
                return iife161_(((pat_let_tv63_)[True] if (True) in (pat_let_tv64_) else len((pat_let_tv65_).set(default__.safeIndex(pat_let_tv66_, len(pat_let_tv67_)), pat_let_tv68_))))
            return iife160_(_pat_let42_0)
        r2 = iife159_(d_1690_v62_)
        return r0, r1, r2

    def m15(self, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: int = int(0)
        r3: _dafny.Map = _dafny.Map({})
        d_1694_v0_: _dafny.Array
        def lambda99_(d_1695_i0_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dslysgcu"))

        init52_ = lambda99_
        nw247_ = _dafny.Array(None, 11)
        for i0_52_ in range(nw247_.length(0)):
            nw247_[i0_52_] = init52_(i0_52_)
        d_1694_v0_ = nw247_
        d_1696_v1_: _dafny.Seq
        d_1696_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "barxrs"))
        d_1697_v2_: _dafny.Seq
        d_1697_v2_ = _dafny.SeqWithoutIsStrInference([d_1696_v1_, d_1696_v1_])
        d_1698_v3_: int
        d_1698_v3_ = 168
        index201_ = default__.safeIndex(668, (d_1694_v0_).length(0))
        (d_1694_v0_)[index201_] = ((d_1697_v2_)[default__.safeIndex(d_1698_v3_, len(d_1697_v2_))]) + (d_1696_v1_)
        d_1699_v4_: bool
        d_1699_v4_ = True
        d_1700_v5_: _dafny.Seq
        d_1700_v5_ = _dafny.SeqWithoutIsStrInference([d_1699_v4_, False, d_1699_v4_, (d_1698_v3_) == (d_1698_v3_), d_1699_v4_])
        d_1701_v6_: _dafny.Array
        nw248_ = _dafny.Array(int(0), 8)
        d_1701_v6_ = nw248_
        index202_ = default__.safeIndex(329, (d_1701_v6_).length(0))
        (d_1701_v6_)[index202_] = d_1698_v3_
        d_1702_v7_: _dafny.Map
        d_1702_v7_ = _dafny.Map({d_1699_v4_: 455})
        index203_ = default__.safeIndex(668, (d_1694_v0_).length(0))
        index204_ = default__.safeIndex(329, (d_1701_v6_).length(0))
        rhs227_ = d_1696_v1_
        rhs228_ = default__.fm39(not(d_1699_v4_), d_1699_v4_, globalState)
        rhs229_ = (default__.safeModuloInt(d_1698_v3_, 390)) <= (136)
        rhs230_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1703_i1_ in range(default__.abs(293))]))
        rhs231_ = default__.safeModuloInt(765, (d_1698_v3_) * (((self).f23)[default__.safeIndex(len((d_1702_v7_).set(d_1699_v4_, d_1698_v3_)), len((self).f23))]))
        lhs121_ = d_1694_v0_
        lhs122_ = default__.safeIndex(668, (d_1694_v0_).length(0))
        lhs123_ = globalState
        lhs124_ = d_1701_v6_
        lhs125_ = default__.safeIndex(329, (d_1701_v6_).length(0))
        lhs121_[lhs122_] = rhs227_
        d_1700_v5_ = rhs228_
        lhs123_.f8 = rhs229_
        lhs124_[lhs125_] = rhs230_
        d_1698_v3_ = rhs231_
        d_1704_v8_: _dafny.Map
        d_1704_v8_ = _dafny.Map({(d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]: D3_DC8(d_1702_v7_)})
        d_1705_v9_: _dafny.Set
        d_1705_v9_ = _dafny.Set({d_1699_v4_, d_1699_v4_, d_1699_v4_})
        d_1706_v10_: _dafny.Set
        d_1706_v10_ = _dafny.Set({len(d_1705_v9_), len((self).f23)})
        d_1707_v11_: _dafny.Map
        d_1707_v11_ = _dafny.Map({(d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]: 548})
        d_1708_v12_: _dafny.Map
        d_1708_v12_ = _dafny.Map({d_1698_v3_: d_1707_v11_})
        d_1709_v13_: D3
        d_1709_v13_ = D3_DC10(((d_1704_v8_)[len(d_1706_v10_)] if (len(d_1706_v10_)) in (d_1704_v8_) else D3_DC9(d_1698_v3_, d_1708_v12_, (d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))], ((self).f23)[default__.safeIndex(d_1698_v3_, len((self).f23))])))
        d_1709_v13_ = d_1709_v13_
        d_1710_v14_: _dafny.Array
        nw249_ = _dafny.Array(False, 28)
        d_1710_v14_ = nw249_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1710_v14_).length(0)):
            d_1711_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_1711_i2_)) and ((d_1711_i2_) < ((d_1710_v14_).length(0)))):
                _ingredientsd_0.append((d_1710_v14_, int(d_1711_i2_), ((d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]) <= (d_1698_v3_)))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        index205_ = default__.safeIndex(329, (d_1701_v6_).length(0))
        (d_1701_v6_)[index205_] = d_1698_v3_
        r2 = (d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]
        d_1712_v15_: _dafny.MultiSet
        d_1712_v15_ = _dafny.MultiSet([d_1698_v3_])
        d_1713_v16_: _dafny.Map
        d_1713_v16_ = _dafny.Map({(d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]: d_1712_v15_})
        r1 = (0) - (((d_1712_v15_)[(d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]] if ((d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]) in (d_1712_v15_) else default__.safeModuloInt(len(d_1713_v16_), ((d_1707_v11_)[(d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]] if ((d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]) in (d_1707_v11_) else d_1698_v3_))))
        d_1714_v17_: _dafny.MultiSet
        d_1714_v17_ = _dafny.MultiSet([True, d_1699_v4_])
        d_1715_v18_: _dafny.Map
        d_1715_v18_ = _dafny.Map({d_1699_v4_: d_1714_v17_})
        r0 = (((d_1715_v18_)[not(d_1699_v4_)] if (not(d_1699_v4_)) in (d_1715_v18_) else _dafny.MultiSet(d_1700_v5_))).intersection((_dafny.MultiSet([d_1699_v4_, d_1699_v4_, d_1699_v4_, d_1699_v4_, d_1699_v4_])) | (d_1714_v17_))
        r1 = ((0) - (d_1698_v3_)) - (default__.safeModuloInt(-548, (d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))]))
        r2 = default__.safeDivisionInt(len(((d_1694_v0_)[default__.safeIndex(668, (d_1694_v0_).length(0))]) + ((d_1694_v0_)[default__.safeIndex(668, (d_1694_v0_).length(0))])), d_1698_v3_)
        d_1716_v19_: _dafny.Map
        d_1716_v19_ = _dafny.Map({d_1699_v4_: d_1699_v4_})
        d_1717_v20_: _dafny.Map
        d_1717_v20_ = _dafny.Map({((d_1716_v19_)[d_1699_v4_] if (d_1699_v4_) in (d_1716_v19_) else d_1699_v4_): d_1699_v4_})
        r3 = ((d_1717_v20_ if d_1699_v4_ else default__.fm42((d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))], (d_1701_v6_)[default__.safeIndex(329, (d_1701_v6_).length(0))], globalState))) | (_dafny.Map({False: d_1699_v4_}))
        return r0, r1, r2, r3

    @property
    def f23(self):
        return self._f23

class C9(T1, T0, T2):
    def  __init__(self):
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: int = int(0)
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f21: _dafny.Set = _dafny.Set({})
        self.f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f21, f22, f12, f13, f11, f20):
        (self).f21 = f21
        (self).f22 = f22
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f11 = f11
        (self)._f20 = f20

    def fm7(self, globalState):
        def iife163_():
            coll75_ = _dafny.Set()
            compr_75_: _dafny.Set
            for compr_75_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({self.f11, self.f11, False, self.f11, not(self.f11)}), _dafny.Set({self.f11, self.f11})])).Elements:
                d_1718_v0_: _dafny.Set = compr_75_
                if (d_1718_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({self.f11, self.f11, False, self.f11, not(self.f11)}), _dafny.Set({self.f11, self.f11})])):
                    coll75_ = coll75_.union(_dafny.Set([d_1718_v0_]))
            return _dafny.Set(coll75_)
        return (_dafny.Set({_dafny.Set({self.f11, self.f11, not(self.f11), self.f11})})).ispropersubset((_dafny.Set({_dafny.Set({self.f11, self.f11}), _dafny.Set({True})})) - (iife163_()
        ))

    def fm8(self, p0, p1, p2, globalState):
        if self.f11:
            return (self).f20
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1719_i0_ in range(default__.abs(912))])

    def fm5(self, p0, globalState):
        return D0_DC1((self).f12, (self).f20, (self).f12, (_dafny.Set({len(_dafny.Set({self.f11, self.f11, self.f11}))})).issubset(_dafny.Set({(self).f13})))

    def fm6(self, p0, p1, p2, globalState):
        return not(self.f11)

    def fm18(self, p0, p1, p2, p3, globalState):
        return self.f11

    def fm19(self, p0, p1, globalState):
        return _dafny.CodePoint('f')

    def m11(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        d_1720_v0_: _dafny.Seq
        d_1720_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "almbhp"))
        d_1720_v0_ = d_1720_v0_
        d_1721_v1_: _dafny.Map
        d_1721_v1_ = _dafny.Map({self.f11: not(self.f11)})
        d_1722_v2_: _dafny.Map
        d_1722_v2_ = _dafny.Map({self.f22: self.f11})
        d_1723_v3_: _dafny.Map
        d_1723_v3_ = _dafny.Map({p1: ((d_1721_v1_)[self.f11] if (self.f11) in (d_1721_v1_) else ((d_1722_v2_)[p0] if (p0) in (d_1722_v2_) else self.f11))})
        d_1724_v4_: _dafny.MultiSet
        d_1724_v4_ = _dafny.MultiSet([self.f11, self.f11, self.f11])
        d_1725_v5_: _dafny.Seq
        d_1725_v5_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f11]))])
        d_1726_v6_: _dafny.Set
        d_1726_v6_ = _dafny.Set({d_1725_v5_, d_1725_v5_})
        d_1722_v2_ = (d_1722_v2_).set((((d_1724_v4_)[self.f11] if (self.f11) in (d_1724_v4_) else len(d_1726_v6_))) + ((0) - (p1)), not((self).fm7(globalState)))
        if ((0) - (default__.safeModuloInt((self).f12, self.f22))) != (self.f22):
            d_1727_v7_: D0
            d_1727_v7_ = D0_DC3((p0) * (p2), self.f11)
            d_1728_v8_: _dafny.Seq
            d_1728_v8_ = _dafny.SeqWithoutIsStrInference([False, self.f11])
            d_1729_v9_: _dafny.Array
            nw250_ = _dafny.Array(None, 24)
            nw250_[int(0)] = self.f11
            nw250_[int(1)] = self.f11
            nw250_[int(2)] = self.f11
            nw250_[int(3)] = self.f11
            nw250_[int(4)] = (not(self.f11)) and (self.f11)
            nw250_[int(5)] = (_dafny.Map({not(self.f11): self.f11})) != (d_1721_v1_)
            nw250_[int(6)] = self.f11
            nw250_[int(7)] = self.f11
            nw250_[int(8)] = self.f11
            nw250_[int(9)] = self.f11
            nw250_[int(10)] = self.f11
            nw250_[int(11)] = self.f11
            nw250_[int(12)] = True
            nw250_[int(13)] = self.f11
            nw250_[int(14)] = (True) or (not(self.f11))
            nw250_[int(15)] = self.f11
            nw250_[int(16)] = (d_1728_v8_)[default__.safeIndex((self).f13, len(d_1728_v8_))]
            nw250_[int(17)] = self.f11
            nw250_[int(18)] = (self.f11) or (self.f11)
            nw250_[int(19)] = (self.f11) or (self.f11)
            nw250_[int(20)] = False
            nw250_[int(21)] = self.f11
            nw250_[int(22)] = ((d_1722_v2_)[(self).f13] if ((self).f13) in (d_1722_v2_) else self.f11)
            nw250_[int(23)] = not(not (self.f11) or (self.f11))
            d_1729_v9_ = nw250_
            rhs232_ = d_1729_v9_
            rhs233_ = d_1727_v7_
            r0 = rhs232_
            d_1727_v7_ = rhs233_
            d_1730_v10_: _dafny.Map
            d_1730_v10_ = _dafny.Map({self.f11: len(d_1720_v0_)})
            d_1731_v11_: _dafny.Map
            d_1731_v11_ = _dafny.Map({len(d_1730_v10_): p1})
            d_1732_v12_: _dafny.MultiSet
            d_1732_v12_ = _dafny.MultiSet([((d_1731_v11_)[p1] if (p1) in (d_1731_v11_) else len((self).f20))])
            d_1733_v13_: D1
            d_1733_v13_ = D1_DC5(d_1732_v12_, self.f11)
            d_1734_v14_: D1
            d_1734_v14_ = D1_DC6(d_1733_v13_)
            d_1735_v15_: _dafny.Seq
            d_1735_v15_ = _dafny.SeqWithoutIsStrInference([d_1734_v14_, d_1734_v14_])
            d_1736_v16_: str
            d_1736_v16_ = _dafny.CodePoint('m')
            d_1737_v17_: _dafny.Map
            d_1737_v17_ = _dafny.Map({d_1735_v15_: d_1736_v16_})
            d_1738_v18_: _dafny.Map
            d_1738_v18_ = _dafny.Map({886: d_1737_v17_})
            rhs234_ = ((p0) - (len(default__.fm20(self.f22, len((self).f20), len(d_1721_v1_), globalState)))) + ((0) - ((0) - ((0) - ((self).f13))))
            rhs235_ = ((d_1732_v12_)[default__.safeDivisionInt(p2, p0)] if (default__.safeDivisionInt(p2, p0)) in (d_1732_v12_) else (self).f12)
            rhs236_ = ((self.f11 if self.f11 else self.f11)) == (self.f11)
            rhs237_ = (p2) * (p2)
            rhs238_ = len((((d_1738_v18_)[(self).f12] if ((self).f12) in (d_1738_v18_) else d_1737_v17_)) | ((d_1737_v17_) | (d_1737_v17_)))
            lhs126_ = self
            lhs127_ = self
            r1 = rhs234_
            r1 = rhs235_
            lhs126_.f11 = rhs236_
            r1 = rhs237_
            lhs127_.f22 = rhs238_
            d_1739_v19_: _dafny.Set
            d_1739_v19_ = _dafny.Set({(self).f13})
            d_1740_v20_: _dafny.Seq
            d_1740_v20_ = _dafny.SeqWithoutIsStrInference([(d_1739_v19_) | (d_1739_v19_), d_1739_v19_, _dafny.Set({(self).f13})])
            index206_ = default__.safeIndex(98, (d_1729_v9_).length(0))
            (d_1729_v9_)[index206_] = not(not(self.f11))
            d_1741_v21_: _dafny.MultiSet
            d_1741_v21_ = _dafny.MultiSet([p3])
            d_1742_v22_: _dafny.Map
            d_1742_v22_ = _dafny.Map({(d_1741_v21_).cardinality: d_1739_v19_})
            d_1743_v23_: _dafny.Seq
            d_1743_v23_ = _dafny.SeqWithoutIsStrInference([d_1729_v9_, d_1729_v9_])
            d_1744_v24_: _dafny.Set
            d_1744_v24_ = _dafny.Set({(d_1743_v23_)[default__.safeIndex(p0, len(d_1743_v23_))]})
            d_1745_v25_: _dafny.Map
            d_1745_v25_ = _dafny.Map({d_1736_v16_: self.f11})
            d_1746_v26_: _dafny.Map
            d_1746_v26_ = _dafny.Map({d_1725_v5_: (self).fm8(len(d_1744_v24_), d_1745_v25_, d_1739_v19_, globalState)})
            index207_ = default__.safeIndex(98, (d_1729_v9_).length(0))
            rhs239_ = ((d_1740_v20_) + ((_dafny.SeqWithoutIsStrInference([d_1739_v19_, d_1739_v19_])) + (d_1740_v20_))).set(default__.safeIndex(default__.fm2(globalState), len((d_1740_v20_) + ((_dafny.SeqWithoutIsStrInference([d_1739_v19_, d_1739_v19_])) + (d_1740_v20_)))), (((d_1742_v22_)[(self).f13] if ((self).f13) in (d_1742_v22_) else _dafny.Set({((d_1724_v4_)[self.f11] if (self.f11) in (d_1724_v4_) else 577), p1}))) | (d_1739_v19_))
            rhs240_ = self.f11
            rhs241_ = len(d_1746_v26_)
            rhs242_ = (d_1728_v8_)[default__.safeIndex(len(d_1720_v0_), len(d_1728_v8_))]
            rhs243_ = (0) - (p2)
            lhs128_ = d_1729_v9_
            lhs129_ = default__.safeIndex(98, (d_1729_v9_).length(0))
            lhs130_ = self
            lhs131_ = self
            lhs132_ = self
            d_1740_v20_ = rhs239_
            lhs128_[lhs129_] = rhs240_
            lhs130_.f22 = rhs241_
            lhs131_.f11 = rhs242_
            lhs132_.f22 = rhs243_
            if self.f11:
                index208_ = default__.safeIndex(669, (p3).length(0))
                (p3)[index208_] = self.f22
                index209_ = default__.safeIndex(669, (p3).length(0))
                (p3)[index209_] = ((_dafny.MultiSet([self.f11, (d_1729_v9_)[default__.safeIndex(98, (d_1729_v9_).length(0))]])).cardinality) + ((self).f13)
                d_1747_v27_: _dafny.Map
                d_1747_v27_ = _dafny.Map({self.f11: d_1720_v0_})
                d_1748_v28_: D0
                d_1748_v28_ = D0_DC1(len(_dafny.Set({len(d_1725_v5_)})), (self).f20, (self).f13, (d_1729_v9_)[default__.safeIndex(98, (d_1729_v9_).length(0))])
                d_1720_v0_ = ((d_1747_v27_)[False] if (False) in (d_1747_v27_) else (d_1748_v28_).cf2)
                r1 = p0
                d_1749_v29_: _dafny.Map
                d_1749_v29_ = _dafny.Map({d_1722_v2_: d_1741_v21_})
                d_1749_v29_ = (d_1749_v29_).set(d_1723_v3_, d_1741_v21_)
                d_1750_v30_: _dafny.Array
                def lambda100_(d_1751_i0_):
                    return (d_1751_i0_) + (self.f22)

                init53_ = lambda100_
                nw251_ = _dafny.Array(None, 16)
                for i0_53_ in range(nw251_.length(0)):
                    nw251_[i0_53_] = init53_(i0_53_)
                d_1750_v30_ = nw251_
                d_1752_v31_: _dafny.Seq
                d_1752_v31_ = _dafny.SeqWithoutIsStrInference([d_1750_v30_])
                d_1753_v32_: _dafny.Array
                nw252_ = _dafny.Array(None, 18)
                nw252_[int(0)] = d_1752_v31_
                nw252_[int(1)] = d_1752_v31_
                nw252_[int(2)] = (d_1752_v31_) + (d_1752_v31_)
                nw252_[int(3)] = d_1752_v31_
                nw252_[int(4)] = (d_1752_v31_) + (d_1752_v31_)
                nw252_[int(5)] = d_1752_v31_
                nw252_[int(6)] = d_1752_v31_
                nw252_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_1750_v30_, d_1750_v30_])) + (d_1752_v31_)
                nw252_[int(8)] = (d_1752_v31_) + (d_1752_v31_)
                nw252_[int(9)] = d_1752_v31_
                nw252_[int(10)] = d_1752_v31_
                nw252_[int(11)] = d_1752_v31_
                nw252_[int(12)] = d_1752_v31_
                nw252_[int(13)] = d_1752_v31_
                nw252_[int(14)] = (d_1752_v31_) + (d_1752_v31_)
                nw252_[int(15)] = d_1752_v31_
                nw252_[int(16)] = (d_1752_v31_).set(default__.safeIndex(p1, len(d_1752_v31_)), d_1750_v30_)
                nw252_[int(17)] = (_dafny.SeqWithoutIsStrInference([d_1750_v30_])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_1750_v30_]))), d_1750_v30_)
                d_1753_v32_ = nw252_
                index210_ = default__.safeIndex(407, (d_1753_v32_).length(0))
                (d_1753_v32_)[index210_] = (d_1752_v31_) + (_dafny.SeqWithoutIsStrInference([d_1750_v30_, d_1750_v30_]))
                index211_ = default__.safeIndex(407, (d_1753_v32_).length(0))
                (d_1753_v32_)[index211_] = _dafny.SeqWithoutIsStrInference([d_1750_v30_, d_1750_v30_, d_1750_v30_, d_1750_v30_, d_1750_v30_])
            elif True:
                (globalState).f8 = (d_1728_v8_)[default__.safeIndex((self).f13, len(d_1728_v8_))]
                d_1754_v33_: _dafny.Seq
                d_1754_v33_ = d_1725_v5_
                d_1755_v34_: _dafny.Map
                d_1755_v34_ = _dafny.Map({d_1754_v33_: self.f22})
                d_1755_v34_ = (d_1755_v34_).set(d_1754_v33_, p2)
                d_1756_v35_: C2
                nw253_ = C2()
                nw253_.ctor__((len(_dafny.SeqWithoutIsStrInference([self.f22])) if self.f11 else p2), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqwuuwg")), default__.safeModuloInt(self.f22, default__.fm2(globalState)), (d_1732_v12_).cardinality, not((p2) <= (len((self).f20))))
                d_1756_v35_ = nw253_
                index212_ = default__.safeIndex(98, (d_1729_v9_).length(0))
                (d_1729_v9_)[index212_] = self.f11
                d_1757_v36_: D6
                d_1757_v36_ = D6_DC17(False, p1)
                d_1758_v37_: D0
                d_1758_v37_ = D0_DC1((d_1757_v36_).cf31, (self).f20, (d_1756_v35_).f26, (d_1729_v9_)[default__.safeIndex(98, (d_1729_v9_).length(0))])
                d_1759_v38_: _dafny.Array
                nw254_ = _dafny.Array(None, 4)
                nw254_[int(0)] = d_1729_v9_
                nw254_[int(1)] = d_1729_v9_
                nw254_[int(2)] = d_1729_v9_
                nw254_[int(3)] = d_1729_v9_
                d_1759_v38_ = nw254_
                d_1760_v39_: bool
                d_1761_v40_: bool
                out24_: bool
                out25_: bool
                out24_, out25_ = (self).m12((self).f20, d_1758_v37_, d_1759_v38_, globalState)
                d_1760_v39_ = out24_
                d_1761_v40_ = out25_
            r1 = (self).f12
        elif True:
            if self.f11:
                d_1762_v41_: _dafny.Map
                d_1762_v41_ = _dafny.Map({self.f22: -956})
                d_1762_v41_ = ((d_1762_v41_) | (_dafny.Map({(self).f13: self.f22}))) | (d_1762_v41_)
                d_1763_v42_: _dafny.Seq
                d_1763_v42_ = _dafny.SeqWithoutIsStrInference([False])
                (globalState).f8 = not((((d_1763_v42_) + ((_dafny.SeqWithoutIsStrInference([self.f11, self.f11])).set(default__.safeIndex(self.f22, len(_dafny.SeqWithoutIsStrInference([self.f11, self.f11]))), self.f11))).set(default__.safeIndex(len(d_1763_v42_), len((d_1763_v42_) + ((_dafny.SeqWithoutIsStrInference([self.f11, self.f11])).set(default__.safeIndex(self.f22, len(_dafny.SeqWithoutIsStrInference([self.f11, self.f11]))), self.f11)))), self.f11)) <= (d_1763_v42_))
                d_1764_v43_: _dafny.Map
                d_1764_v43_ = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference([False, self.f11])})
                d_1765_v44_: _dafny.Map
                d_1765_v44_ = _dafny.Map({p3: d_1764_v43_})
                d_1765_v44_ = (d_1765_v44_).set(p3, d_1764_v43_)
                d_1766_v45_: _dafny.Array
                def lambda101_(d_1767_i1_):
                    return D6_DC18(not(self.f11))

                init54_ = lambda101_
                nw255_ = _dafny.Array(None, 25)
                for i0_54_ in range(nw255_.length(0)):
                    nw255_[i0_54_] = init54_(i0_54_)
                d_1766_v45_ = nw255_
                d_1768_v46_: D6
                d_1768_v46_ = D6_DC18(not(self.f11))
                index213_ = default__.safeIndex(92, (d_1766_v45_).length(0))
                (d_1766_v45_)[index213_] = d_1768_v46_
                index214_ = default__.safeIndex(92, (d_1766_v45_).length(0))
                (d_1766_v45_)[index214_] = d_1768_v46_
                d_1769_v47_: _dafny.Array
                nw256_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_1769_v47_ = nw256_
                d_1769_v47_ = d_1769_v47_
            elif True:
                (self).f11 = self.f11
                d_1770_v48_: _dafny.Array
                nw257_ = _dafny.Array(False, 8)
                d_1770_v48_ = nw257_
                index215_ = default__.safeIndex(932, (d_1770_v48_).length(0))
                (d_1770_v48_)[index215_] = self.f11
                index216_ = default__.safeIndex(932, (d_1770_v48_).length(0))
                (d_1770_v48_)[index216_] = self.f11
                index217_ = default__.safeIndex(233, (p3).length(0))
                (p3)[index217_] = (self).f12
                index218_ = default__.safeIndex(233, (p3).length(0))
                (p3)[index218_] = (694) + (-748)
                (globalState).f8 = (d_1770_v48_)[default__.safeIndex(932, (d_1770_v48_).length(0))]
                d_1771_v49_: C2
                nw258_ = C2()
                nw258_.ctor__(p1, d_1720_v0_, default__.fm2(globalState), (p2) + (p0), self.f11)
                d_1771_v49_ = nw258_
            if self.f11:
                d_1772_v50_: _dafny.Array
                nw259_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_1772_v50_ = nw259_
                d_1773_v51_: _dafny.Seq
                d_1773_v51_ = _dafny.SeqWithoutIsStrInference([False])
                r1 = default__.safeModuloInt((0) - ((D7_DC23(d_1772_v50_, d_1773_v51_, self.f11, self.f22)).cf45), ((0) - (p0)) - (p0))
                r1 = default__.safeDivisionInt(default__.safeDivisionInt(p0, p0), (self).f13)
                d_1774_v52_: _dafny.Array
                def lambda102_(d_1775_i2_):
                    return _dafny.CodePoint('g')

                init55_ = lambda102_
                nw260_ = _dafny.Array(None, 22)
                for i0_55_ in range(nw260_.length(0)):
                    nw260_[i0_55_] = init55_(i0_55_)
                d_1774_v52_ = nw260_
                d_1774_v52_ = d_1774_v52_
                (globalState).f8 = False
                index219_ = default__.safeIndex(878, (p3).length(0))
                (p3)[index219_] = p1
                index220_ = default__.safeIndex(878, (p3).length(0))
                (p3)[index220_] = self.f22
            elif True:
                r1 = (0) - ((p2 if self.f11 else p0))
                d_1724_v4_ = d_1724_v4_
                d_1720_v0_ = (self).f20
                index221_ = default__.safeIndex(253, (p3).length(0))
                (p3)[index221_] = p2
                index222_ = default__.safeIndex(253, (p3).length(0))
                rhs244_ = self.f11
                rhs245_ = self.f22
                rhs246_ = (0) - (p1)
                lhs133_ = self
                lhs134_ = self
                lhs135_ = p3
                lhs136_ = default__.safeIndex(253, (p3).length(0))
                lhs133_.f11 = rhs244_
                lhs134_.f22 = rhs245_
                lhs135_[lhs136_] = rhs246_
                d_1776_v53_: _dafny.Array
                nw261_ = _dafny.Array(False, 22)
                d_1776_v53_ = nw261_
                index223_ = default__.safeIndex(581, (d_1776_v53_).length(0))
                (d_1776_v53_)[index223_] = True
                d_1777_v54_: str
                d_1777_v54_ = _dafny.CodePoint('n')
                d_1778_v55_: _dafny.Map
                d_1778_v55_ = _dafny.Map({(self).f13: p1})
                d_1779_v56_: D0
                d_1779_v56_ = D0_DC1((self).f13, (self).f20, len(d_1778_v55_), self.f11)
                index224_ = default__.safeIndex(581, (d_1776_v53_).length(0))
                rhs247_ = ((d_1720_v0_) + ((d_1720_v0_).set(default__.safeIndex((0) - (self.f22), len(d_1720_v0_)), d_1777_v54_))) + ((d_1720_v0_) + ((d_1779_v56_).cf2))
                rhs248_ = (d_1720_v0_ if False else (self).f20)
                rhs249_ = (_dafny.CodePoint('i')) in (_dafny.SeqWithoutIsStrInference([d_1777_v54_ for d_1780_i3_ in range(default__.abs(394))]))
                lhs137_ = d_1776_v53_
                lhs138_ = default__.safeIndex(581, (d_1776_v53_).length(0))
                d_1720_v0_ = rhs247_
                d_1720_v0_ = rhs248_
                lhs137_[lhs138_] = rhs249_
            d_1781_v57_: _dafny.Set
            d_1781_v57_ = _dafny.Set({self.f11})
            d_1782_v58_: _dafny.Seq
            d_1782_v58_ = _dafny.SeqWithoutIsStrInference([d_1720_v0_])
            rhs250_ = (self.f22) == (default__.fm2(globalState))
            rhs251_ = (self).f13
            rhs252_ = ((self).f12) == (p2)
            rhs253_ = ((d_1782_v58_)[default__.safeIndex((0) - (self.f22), len(d_1782_v58_))]) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1783_i4_ in range(default__.abs(713))]))
            rhs254_ = d_1781_v57_
            lhs139_ = globalState
            lhs140_ = self
            lhs139_.f8 = rhs250_
            r1 = rhs251_
            lhs140_.f11 = rhs252_
            d_1720_v0_ = rhs253_
            d_1781_v57_ = rhs254_
            d_1784_v59_: str
            d_1784_v59_ = _dafny.CodePoint('p')
            d_1785_v60_: _dafny.Array
            nw262_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_1785_v60_ = nw262_
            d_1786_v61_: D7
            d_1786_v61_ = D7_DC22(p0, d_1785_v60_)
            d_1787_v62_: D9
            d_1787_v62_ = D9_DC27(not(self.f11), d_1786_v61_, p0, self.f11)
            d_1788_v63_: _dafny.Seq
            d_1788_v63_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11])
            d_1789_v64_: _dafny.Map
            d_1789_v64_ = _dafny.Map({_dafny.Map({_dafny.Set({d_1784_v59_}): self.f11}): ((d_1787_v62_).cf49) or ((d_1788_v63_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))), len(d_1788_v63_))])})
            d_1790_v65_: _dafny.Set
            d_1790_v65_ = _dafny.Set({d_1784_v59_, d_1784_v59_})
            d_1791_v66_: _dafny.Map
            d_1791_v66_ = _dafny.Map({d_1790_v65_: self.f11})
            d_1792_v68_: _dafny.MultiSet
            d_1792_v68_ = _dafny.MultiSet([d_1784_v59_])
            d_1793_v70_: _dafny.Seq
            def iife164_():
                coll76_ = _dafny.Set()
                compr_76_: str
                for compr_76_ in (d_1792_v68_).Elements:
                    d_1794_v69_: str = compr_76_
                    if (d_1794_v69_) in (d_1792_v68_):
                        coll76_ = coll76_.union(_dafny.Set([d_1794_v69_]))
                return _dafny.Set(coll76_)
            d_1793_v70_ = _dafny.SeqWithoutIsStrInference([iife164_()
            ])
            def iife165_():
                coll77_ = _dafny.Map()
                compr_77_: _dafny.Set
                for compr_77_ in (d_1793_v70_).Elements:
                    d_1795_v67_: _dafny.Set = compr_77_
                    if (d_1795_v67_) in (d_1793_v70_):
                        coll77_[d_1795_v67_] = self.f11
                return _dafny.Map(coll77_)
            d_1789_v64_ = (d_1789_v64_).set((d_1791_v66_) | (iife165_()
            ), self.f11)
            d_1796_v71_: D0
            d_1796_v71_ = D0_DC1(p1, (self).f20, (self).f12, (default__.fm46(self.f11, self.f11, globalState)).cf21)
            d_1797_v72_: _dafny.Array
            nw263_ = _dafny.Array(False, 6)
            d_1797_v72_ = nw263_
            d_1798_v73_: _dafny.Array
            nw264_ = _dafny.Array(None, 3)
            nw264_[int(0)] = d_1797_v72_
            nw264_[int(1)] = d_1797_v72_
            nw264_[int(2)] = d_1797_v72_
            d_1798_v73_ = nw264_
            d_1799_v74_: _dafny.Seq
            d_1799_v74_ = _dafny.SeqWithoutIsStrInference([d_1798_v73_, d_1798_v73_, d_1798_v73_, d_1798_v73_, d_1798_v73_])
            d_1800_v75_: bool
            d_1801_v76_: bool
            out26_: bool
            out27_: bool
            out26_, out27_ = (self).m12(d_1720_v0_, d_1796_v71_, (d_1799_v74_)[default__.safeIndex(self.f22, len(d_1799_v74_))], globalState)
            d_1800_v75_ = out26_
            d_1801_v76_ = out27_
        hi6_ = len(d_1720_v0_)
        for d_1802_i5_ in range(p1, hi6_):
            d_1803_v77_: _dafny.Seq
            d_1803_v77_ = _dafny.SeqWithoutIsStrInference([self.f11, True, self.f11])
            d_1804_v78_: _dafny.Seq
            d_1804_v78_ = _dafny.SeqWithoutIsStrInference([(d_1803_v77_)[default__.safeIndex((_dafny.MultiSet(d_1725_v5_)).cardinality, len(d_1803_v77_))], self.f11])
            d_1805_v79_: _dafny.Array
            nw265_ = _dafny.Array(None, 11)
            nw265_[int(0)] = False
            nw265_[int(1)] = (True) or ((self).fm7(globalState))
            nw265_[int(2)] = self.f11
            nw265_[int(3)] = self.f11
            nw265_[int(4)] = self.f11
            nw265_[int(5)] = self.f11
            nw265_[int(6)] = (self.f11 if self.f11 else self.f11)
            nw265_[int(7)] = (self.f11 if (d_1803_v77_)[default__.safeIndex(p0, len(d_1803_v77_))] else self.f11)
            nw265_[int(8)] = not(self.f11)
            nw265_[int(9)] = self.f11
            nw265_[int(10)] = (self.f11) or (not(self.f11))
            d_1805_v79_ = nw265_
            index225_ = default__.safeIndex(600, (d_1805_v79_).length(0))
            (d_1805_v79_)[index225_] = self.f11
            d_1806_v80_: _dafny.Set
            d_1806_v80_ = _dafny.Set({not(self.f11), self.f11})
            d_1807_v81_: _dafny.Map
            d_1807_v81_ = _dafny.Map({d_1806_v80_: p1})
            index226_ = default__.safeIndex(600, (d_1805_v79_).length(0))
            (d_1805_v79_)[index226_] = (((d_1807_v81_)[d_1806_v80_] if (d_1806_v80_) in (d_1807_v81_) else self.f22)) < ((self.f22) + (len(d_1721_v1_)))
            index227_ = default__.safeIndex(600, (d_1805_v79_).length(0))
            (d_1805_v79_)[index227_] = (d_1805_v79_)[default__.safeIndex(600, (d_1805_v79_).length(0))]
            index228_ = default__.safeIndex(600, (d_1805_v79_).length(0))
            (d_1805_v79_)[index228_] = (_dafny.SeqWithoutIsStrInference([-205 for d_1808_i6_ in range(default__.abs(733))])) == (d_1725_v5_)
            d_1809_v82_: D0
            d_1809_v82_ = D0_DC2(len(_dafny.SeqWithoutIsStrInference([self.f11])), p3)
            d_1810_v83_: _dafny.Seq
            d_1810_v83_ = _dafny.SeqWithoutIsStrInference([p3, p3])
            d_1811_v84_: _dafny.Array
            nw266_ = _dafny.Array(None, 29)
            nw266_[int(0)] = p3
            nw266_[int(1)] = p3
            nw266_[int(2)] = p3
            nw266_[int(3)] = p3
            nw266_[int(4)] = p3
            nw266_[int(5)] = p3
            nw266_[int(6)] = p3
            nw266_[int(7)] = p3
            nw266_[int(8)] = p3
            nw266_[int(9)] = p3
            nw266_[int(10)] = p3
            nw266_[int(11)] = p3
            nw266_[int(12)] = p3
            nw266_[int(13)] = p3
            nw266_[int(14)] = p3
            nw266_[int(15)] = p3
            nw266_[int(16)] = p3
            nw266_[int(17)] = p3
            nw266_[int(18)] = p3
            nw266_[int(19)] = (d_1809_v82_).cf6
            nw266_[int(20)] = p3
            nw266_[int(21)] = p3
            nw266_[int(22)] = (d_1810_v83_)[default__.safeIndex(self.f22, len(d_1810_v83_))]
            nw266_[int(23)] = p3
            nw266_[int(24)] = p3
            nw266_[int(25)] = p3
            nw266_[int(26)] = p3
            nw266_[int(27)] = p3
            nw266_[int(28)] = p3
            d_1811_v84_ = nw266_
            d_1812_v85_: D7
            d_1812_v85_ = D7_DC22((self).f12, d_1811_v84_)
            pat_let_tv69_ = d_1811_v84_
            def iife166_(_pat_let44_0):
                def iife167_(d_1813_dt__update__tmp_h0_):
                    def iife168_(_pat_let45_0):
                        def iife169_(d_1814_dt__update_hcf41_h0_):
                            return D7_DC22((d_1813_dt__update__tmp_h0_).cf40, d_1814_dt__update_hcf41_h0_)
                        return iife169_(_pat_let45_0)
                    return iife168_(pat_let_tv69_)
                return iife167_(_pat_let44_0)
            source28_ = iife166_(d_1812_v85_)
            if source28_.is_DC22:
                d_1815___mcc_h0_ = source28_.cf40
                d_1816___mcc_h1_ = source28_.cf41
                d_1817_cf41_ = d_1816___mcc_h1_
                d_1818_cf40_ = d_1815___mcc_h0_
                d_1819_v86_: C8
                nw267_ = C8()
                nw267_.ctor__((d_1725_v5_) + (d_1725_v5_))
                d_1819_v86_ = nw267_
                d_1818_cf40_ = ((d_1819_v86_).f23)[default__.safeIndex((p2 if (d_1805_v79_)[default__.safeIndex(600, (d_1805_v79_).length(0))] else (self).f12), len((d_1819_v86_).f23))]
                d_1820_v87_: _dafny.Map
                d_1820_v87_ = _dafny.Map({d_1720_v0_: d_1805_v79_})
                d_1821_v88_: _dafny.Seq
                d_1821_v88_ = _dafny.SeqWithoutIsStrInference([d_1805_v79_])
                d_1820_v87_ = (d_1820_v87_).set(((self).f20) + ((self).f20), (d_1821_v88_)[default__.safeIndex((self).f13, len(d_1821_v88_))])
                d_1822_v89_: _dafny.MultiSet
                d_1822_v89_ = _dafny.MultiSet([_dafny.Map({d_1802_i5_: not((d_1805_v79_)[default__.safeIndex(600, (d_1805_v79_).length(0))])})])
                d_1823_v91_: _dafny.Map
                def iife170_():
                    coll78_ = _dafny.Set()
                    compr_78_: _dafny.Map
                    for compr_78_ in (d_1822_v89_).Elements:
                        d_1824_v90_: _dafny.Map = compr_78_
                        if (d_1824_v90_) in (d_1822_v89_):
                            coll78_ = coll78_.union(_dafny.Set([d_1824_v90_]))
                    return _dafny.Set(coll78_)
                d_1823_v91_ = _dafny.Map({len(iife170_()
                ): self.f22})
                d_1825_v92_: D0
                d_1825_v92_ = D0_DC3((self).f12, (d_1805_v79_)[default__.safeIndex(600, (d_1805_v79_).length(0))])
                d_1826_v93_: str
                d_1826_v93_ = _dafny.CodePoint('c')
                d_1827_v94_: _dafny.MultiSet
                d_1827_v94_ = _dafny.MultiSet([(self).f13, d_1802_i5_])
                d_1828_v95_: _dafny.Array
                nw268_ = _dafny.Array(None, 18)
                nw268_[int(0)] = (self).f12
                nw268_[int(1)] = len((self).f20)
                nw268_[int(2)] = len((d_1823_v91_) | (d_1823_v91_))
                nw268_[int(3)] = len((d_1806_v80_).intersection(default__.fm4(d_1818_cf40_, d_1825_v92_, len(d_1721_v1_), d_1826_v93_, globalState)))
                nw268_[int(4)] = len((d_1720_v0_) + (d_1720_v0_))
                nw268_[int(5)] = p0
                nw268_[int(6)] = (d_1827_v94_).cardinality
                nw268_[int(7)] = self.f22
                nw268_[int(8)] = d_1818_cf40_
                nw268_[int(9)] = (self).f12
                nw268_[int(10)] = p2
                nw268_[int(11)] = d_1802_i5_
                nw268_[int(12)] = (self.f22) + (p1)
                nw268_[int(13)] = (((d_1827_v94_)[p0] if (p0) in (d_1827_v94_) else self.f22)) * (781)
                nw268_[int(14)] = d_1802_i5_
                nw268_[int(15)] = p0
                nw268_[int(16)] = self.f22
                nw268_[int(17)] = p1
                d_1828_v95_ = nw268_
                d_1828_v95_ = p3
            elif source28_.is_DC23:
                d_1829___mcc_h2_ = source28_.cf42
                d_1830___mcc_h3_ = source28_.cf43
                d_1831___mcc_h4_ = source28_.cf44
                d_1832___mcc_h5_ = source28_.cf45
                d_1833_cf45_ = d_1832___mcc_h5_
                d_1834_cf44_ = d_1831___mcc_h4_
                d_1835_cf43_ = d_1830___mcc_h3_
                d_1836_cf42_ = d_1829___mcc_h2_
                index229_ = default__.safeIndex(600, (d_1805_v79_).length(0))
                (d_1805_v79_)[index229_] = d_1834_cf44_
                d_1837_v96_: _dafny.Array
                def lambda103_(d_1838_i7_):
                    return default__.safeModuloInt(d_1838_i7_, (_dafny.MultiSet([(self).f20])).cardinality)

                init56_ = lambda103_
                nw269_ = _dafny.Array(None, 2)
                for i0_56_ in range(nw269_.length(0)):
                    nw269_[i0_56_] = init56_(i0_56_)
                d_1837_v96_ = nw269_
                rhs255_ = d_1837_v96_
                rhs256_ = d_1720_v0_
                d_1837_v96_ = rhs255_
                d_1720_v0_ = rhs256_
                d_1839_v97_: D1
                d_1839_v97_ = D1_DC4(d_1810_v83_)
                d_1840_v98_: D1
                d_1840_v98_ = D1_DC6(d_1839_v97_)
                d_1841_v99_: _dafny.Array
                nw270_ = _dafny.Array(None, 3)
                nw270_[int(0)] = d_1840_v98_
                nw270_[int(1)] = d_1840_v98_
                nw270_[int(2)] = D1_DC6(d_1839_v97_)
                d_1841_v99_ = nw270_
                index230_ = default__.safeIndex(621, (d_1841_v99_).length(0))
                (d_1841_v99_)[index230_] = d_1840_v98_
                index231_ = default__.safeIndex(621, (d_1841_v99_).length(0))
                (d_1841_v99_)[index231_] = d_1840_v98_
                r1 = ((self).f13) - (p1)
            elif source28_.is_DC21:
                d_1842___mcc_h6_ = source28_.cf39
                d_1843_cf39_ = d_1842___mcc_h6_
                (self).f22 = (p2) - (p1)
                index232_ = default__.safeIndex(526, (p3).length(0))
                (p3)[index232_] = len(_dafny.SeqWithoutIsStrInference([self.f22, (self).f13]))
                index233_ = default__.safeIndex(526, (p3).length(0))
                (p3)[index233_] = len((d_1722_v2_) | (d_1723_v3_))
                d_1844_v100_: D7
                d_1844_v100_ = D7_DC24(D7_DC23(d_1811_v84_, d_1804_v78_, (d_1805_v79_)[default__.safeIndex(600, (d_1805_v79_).length(0))], p2))
                d_1845_v101_: _dafny.Array
                nw271_ = _dafny.Array(None, 8)
                nw271_[int(0)] = d_1844_v100_
                nw271_[int(1)] = d_1844_v100_
                nw271_[int(2)] = d_1844_v100_
                nw271_[int(3)] = d_1844_v100_
                nw271_[int(4)] = d_1844_v100_
                nw271_[int(5)] = d_1844_v100_
                nw271_[int(6)] = d_1844_v100_
                nw271_[int(7)] = d_1844_v100_
                d_1845_v101_ = nw271_
                index234_ = default__.safeIndex(303, (d_1845_v101_).length(0))
                (d_1845_v101_)[index234_] = d_1844_v100_
                index235_ = default__.safeIndex(303, (d_1845_v101_).length(0))
                rhs257_ = (self).f20
                rhs258_ = d_1803_v77_
                rhs259_ = d_1844_v100_
                lhs141_ = d_1845_v101_
                lhs142_ = default__.safeIndex(303, (d_1845_v101_).length(0))
                d_1720_v0_ = rhs257_
                d_1804_v78_ = rhs258_
                lhs141_[lhs142_] = rhs259_
                d_1846_v102_: _dafny.Array
                nw272_ = _dafny.Array(int(0), 8)
                d_1846_v102_ = nw272_
                d_1847_v103_: _dafny.Map
                d_1847_v103_ = _dafny.Map({d_1846_v102_: (0) - (default__.fm2(globalState))})
                index236_ = default__.safeIndex(526, (p3).length(0))
                (p3)[index236_] = len(d_1847_v103_)
            elif True:
                d_1848___mcc_h7_ = source28_.cf46
                d_1849_cf46_ = d_1848___mcc_h7_
                r1 = default__.safeDivisionInt(p0, (len(d_1803_v77_)) - ((self).f13))
                d_1850_v104_: _dafny.Seq
                d_1850_v104_ = _dafny.SeqWithoutIsStrInference([d_1805_v79_, d_1805_v79_, d_1805_v79_])
                d_1851_v105_: D3
                d_1851_v105_ = D3_DC8(_dafny.Map({self.f11: d_1802_i5_}))
                d_1852_v106_: _dafny.Map
                d_1852_v106_ = _dafny.Map({(d_1805_v79_)[default__.safeIndex(600, (d_1805_v79_).length(0))]: p1})
                r0 = (d_1850_v104_)[default__.safeIndex(len(((d_1851_v105_).cf14) | (d_1852_v106_)), len(d_1850_v104_))]
                d_1853_v107_: str
                d_1853_v107_ = _dafny.CodePoint('s')
                d_1854_v108_: _dafny.Map
                d_1854_v108_ = _dafny.Map({((self).f12) - (p1): d_1853_v107_})
                d_1854_v108_ = _dafny.Map({d_1802_i5_: d_1853_v107_})
                d_1855_v109_: C1
                nw273_ = C1()
                nw273_.ctor__(p0)
                d_1855_v109_ = nw273_
        d_1856_v110_: _dafny.Array
        nw274_ = _dafny.Array(False, 26)
        d_1856_v110_ = nw274_
        index237_ = default__.safeIndex(273, (d_1856_v110_).length(0))
        (d_1856_v110_)[index237_] = (226) >= (p0)
        d_1857_v111_: _dafny.MultiSet
        d_1857_v111_ = _dafny.MultiSet([d_1725_v5_])
        index238_ = default__.safeIndex(273, (d_1856_v110_).length(0))
        (d_1856_v110_)[index238_] = not((d_1857_v111_).issubset(d_1857_v111_))
        if (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))]:
            (self).f11 = (self).fm18(self.f11, (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], p1, p0, globalState)
            d_1858_v112_: C8
            nw275_ = C8()
            nw275_.ctor__((d_1725_v5_) + (d_1725_v5_))
            d_1858_v112_ = nw275_
            index239_ = default__.safeIndex(273, (d_1856_v110_).length(0))
            (d_1856_v110_)[index239_] = not(((self).f20) < ((self).f20))
            r1 = (self).f12
            d_1859_v113_: _dafny.Map
            d_1859_v113_ = _dafny.Map({self.f22: p2})
            source29_ = default__.fm53((d_1859_v113_) | (d_1859_v113_), default__.safeModuloInt(p0, (self).f12), globalState)
            if source29_.is_DC9:
                d_1860___mcc_h8_ = source29_.cf15
                d_1861___mcc_h9_ = source29_.cf16
                d_1862___mcc_h10_ = source29_.cf17
                d_1863___mcc_h11_ = source29_.cf18
                d_1864_cf18_ = d_1863___mcc_h11_
                d_1865_cf17_ = d_1862___mcc_h10_
                d_1866_cf16_ = d_1861___mcc_h9_
                d_1867_cf15_ = d_1860___mcc_h8_
                d_1868_v114_: _dafny.MultiSet
                d_1868_v114_ = _dafny.MultiSet([(self).f12, p1])
                index240_ = default__.safeIndex(273, (d_1856_v110_).length(0))
                (d_1856_v110_)[index240_] = not(not(((0) - (default__.safeModuloInt((d_1868_v114_).cardinality, p1))) != ((0) - (d_1865_cf17_))))
                d_1869_v115_: _dafny.Array
                nw276_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1869_v115_ = nw276_
                d_1870_v116_: D7
                d_1870_v116_ = D7_DC22(p0, d_1869_v115_)
                d_1871_v117_: D9
                d_1871_v117_ = D9_DC27(True, d_1870_v116_, -105, (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))])
                d_1872_v118_: D9
                d_1872_v118_ = D9_DC28(D9_DC28(d_1871_v117_))
                d_1872_v118_ = D9_DC28(d_1871_v117_)
                d_1873_v119_: str
                d_1873_v119_ = _dafny.CodePoint('w')
                d_1874_v120_: _dafny.Array
                nw277_ = _dafny.Array(None, 14)
                nw277_[int(0)] = d_1873_v119_
                nw277_[int(1)] = d_1873_v119_
                nw277_[int(2)] = d_1873_v119_
                nw277_[int(3)] = d_1873_v119_
                nw277_[int(4)] = _dafny.CodePoint('w')
                nw277_[int(5)] = d_1873_v119_
                nw277_[int(6)] = d_1873_v119_
                nw277_[int(7)] = d_1873_v119_
                nw277_[int(8)] = d_1873_v119_
                nw277_[int(9)] = d_1873_v119_
                nw277_[int(10)] = d_1873_v119_
                nw277_[int(11)] = _dafny.CodePoint('c')
                nw277_[int(12)] = _dafny.CodePoint('t')
                nw277_[int(13)] = d_1873_v119_
                d_1874_v120_ = nw277_
                d_1875_v121_: D13
                d_1875_v121_ = D13_DC36(d_1874_v120_)
                d_1876_v122_: _dafny.Array
                nw278_ = _dafny.Array(None, 6)
                nw278_[int(0)] = (d_1875_v121_).cf61
                nw278_[int(1)] = d_1874_v120_
                nw278_[int(2)] = d_1874_v120_
                nw278_[int(3)] = (d_1875_v121_).cf61
                nw278_[int(4)] = d_1874_v120_
                nw278_[int(5)] = d_1874_v120_
                d_1876_v122_ = nw278_
                index241_ = default__.safeIndex(327, (d_1876_v122_).length(0))
                (d_1876_v122_)[index241_] = d_1874_v120_
                index242_ = default__.safeIndex(327, (d_1876_v122_).length(0))
                (d_1876_v122_)[index242_] = d_1874_v120_
                (globalState).f5 = _dafny.SeqWithoutIsStrInference([d_1865_cf17_ for d_1877_i8_ in range(default__.abs(914))])
            elif source29_.is_DC8:
                d_1878___mcc_h12_ = source29_.cf14
                d_1879_cf14_ = d_1878___mcc_h12_
                (globalState).f8 = ((self).fm6((d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], 732, self.f22, globalState)) == (self.f11)
                d_1880_v123_: _dafny.Array
                def lambda104_(d_1881_v2_):
                    def lambda105_(d_1882_i9_):
                        return d_1881_v2_

                    return lambda105_

                init57_ = lambda104_(d_1722_v2_)
                nw279_ = _dafny.Array(None, 9)
                for i0_57_ in range(nw279_.length(0)):
                    nw279_[i0_57_] = init57_(i0_57_)
                d_1880_v123_ = nw279_
                index243_ = default__.safeIndex(460, (d_1880_v123_).length(0))
                (d_1880_v123_)[index243_] = d_1723_v3_
                index244_ = default__.safeIndex(460, (d_1880_v123_).length(0))
                (d_1880_v123_)[index244_] = d_1722_v2_
                d_1883_v124_: _dafny.Array
                def lambda106_(d_1884_i10_):
                    return (d_1884_i10_) - ((self).f13)

                init58_ = lambda106_
                nw280_ = _dafny.Array(None, 11)
                for i0_58_ in range(nw280_.length(0)):
                    nw280_[i0_58_] = init58_(i0_58_)
                d_1883_v124_ = nw280_
                d_1883_v124_ = (d_1883_v124_ if (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))] else p3)
                d_1885_v125_: _dafny.Array
                nw281_ = _dafny.Array(_dafny.MultiSet({}), 1)
                d_1885_v125_ = nw281_
                index245_ = default__.safeIndex(102, (d_1885_v125_).length(0))
                (d_1885_v125_)[index245_] = d_1724_v4_
                d_1886_v126_: _dafny.Map
                d_1886_v126_ = _dafny.Map({p0: d_1725_v5_})
                d_1887_v127_: _dafny.Map
                d_1887_v127_ = _dafny.Map({len(d_1886_v126_): _dafny.MultiSet([self.f11])})
                d_1888_v128_: _dafny.MultiSet
                d_1888_v128_ = _dafny.MultiSet([-915])
                index246_ = default__.safeIndex(102, (d_1885_v125_).length(0))
                (d_1885_v125_)[index246_] = ((((d_1887_v127_)[(d_1888_v128_).cardinality] if ((d_1888_v128_).cardinality) in (d_1887_v127_) else d_1724_v4_)) | (d_1724_v4_)).set((_dafny.MultiSet([self.f11])).ispropersubset(default__.fm30(not((d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))]), (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], 545, globalState)), default__.abs((p1) * (default__.fm2(globalState))))
            elif True:
                d_1889___mcc_h13_ = source29_.cf19
                d_1890_cf19_ = d_1889___mcc_h13_
                d_1891_v129_: _dafny.Array
                nw282_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_1891_v129_ = nw282_
                index247_ = default__.safeIndex(513, (d_1891_v129_).length(0))
                (d_1891_v129_)[index247_] = (self).f20
                d_1892_v130_: _dafny.Seq
                d_1892_v130_ = _dafny.SeqWithoutIsStrInference([(d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))]])
                d_1893_v131_: _dafny.Array
                nw283_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_1893_v131_ = nw283_
                d_1894_v132_: D7
                d_1894_v132_ = D7_DC22(len(d_1892_v130_), d_1893_v131_)
                d_1895_v133_: D9
                d_1895_v133_ = D9_DC27(self.f11, d_1894_v132_, -258, self.f11)
                d_1896_v134_: D9
                d_1896_v134_ = D9_DC28(d_1895_v133_)
                index248_ = default__.safeIndex(513, (d_1891_v129_).length(0))
                rhs260_ = (self).f20
                rhs261_ = len(_dafny.Map({d_1896_v134_: self.f22}))
                lhs143_ = d_1891_v129_
                lhs144_ = default__.safeIndex(513, (d_1891_v129_).length(0))
                lhs143_[lhs144_] = rhs260_
                r1 = rhs261_
                index249_ = default__.safeIndex(273, (d_1856_v110_).length(0))
                (d_1856_v110_)[index249_] = not(((d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))]) and ((p2) >= ((self).f13)))
                d_1897_v135_: _dafny.Array
                nw284_ = _dafny.Array(D5.default()(), 23)
                d_1897_v135_ = nw284_
                d_1898_v136_: D5
                d_1898_v136_ = D5_DC14(d_1891_v129_)
                index250_ = default__.safeIndex(569, (d_1897_v135_).length(0))
                (d_1897_v135_)[index250_] = d_1898_v136_
                index251_ = default__.safeIndex(569, (d_1897_v135_).length(0))
                (d_1897_v135_)[index251_] = d_1898_v136_
                d_1899_v137_: str
                d_1899_v137_ = _dafny.CodePoint('e')
                d_1900_v138_: _dafny.Set
                d_1900_v138_ = _dafny.Set({_dafny.CodePoint('o'), d_1899_v137_, d_1899_v137_, d_1899_v137_})
                d_1901_v139_: _dafny.Map
                d_1901_v139_ = _dafny.Map({d_1900_v138_: p1})
                d_1902_v140_: _dafny.Seq
                d_1902_v140_ = _dafny.SeqWithoutIsStrInference([d_1720_v0_, d_1720_v0_])
                d_1903_v141_: D0
                d_1903_v141_ = D0_DC1((0) - (self.f22), (d_1902_v140_)[default__.safeIndex((self).f12, len(d_1902_v140_))], (self).f12, (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))])
                d_1901_v139_ = (d_1901_v139_).set(d_1900_v138_, ((d_1903_v141_).cf3) * (self.f22))
        elif True:
            d_1904_v142_: _dafny.Array
            nw285_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1904_v142_ = nw285_
            index252_ = default__.safeIndex(983, (d_1904_v142_).length(0))
            (d_1904_v142_)[index252_] = d_1856_v110_
            d_1905_v143_: _dafny.Array
            d_1905_v143_ = d_1856_v110_
            index253_ = default__.safeIndex(983, (d_1904_v142_).length(0))
            (d_1904_v142_)[index253_] = (d_1905_v143_)
            d_1906_v144_: _dafny.Set
            d_1906_v144_ = _dafny.Set({(d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], True})
            index254_ = default__.safeIndex(273, (d_1856_v110_).length(0))
            (d_1856_v110_)[index254_] = (default__.fm2(globalState)) >= (len((d_1906_v144_) | (_dafny.Set({(d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], self.f11}))))
            (self).f22 = ((p2 if (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))] else (self).f12)) - ((p2) * (p2))
            d_1907_v145_: _dafny.Seq
            d_1907_v145_ = _dafny.SeqWithoutIsStrInference([(d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))]])
            index255_ = default__.safeIndex(983, (d_1904_v142_).length(0))
            rhs262_ = (d_1907_v145_) + (_dafny.SeqWithoutIsStrInference([(d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))]]))
            rhs263_ = (len((d_1906_v144_) | (d_1906_v144_)) if (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))] else self.f22)
            rhs264_ = not(((d_1724_v4_).intersection(d_1724_v4_)).isdisjoint((default__.fm30((d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], self.f22, globalState)).set(self.f11, default__.abs(877))))
            rhs265_ = not (not(True)) or ((d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))])
            rhs266_ = ((d_1904_v142_)[default__.safeIndex(983, (d_1904_v142_).length(0))] if (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))] else (d_1904_v142_)[default__.safeIndex(983, (d_1904_v142_).length(0))])
            lhs145_ = self
            lhs146_ = globalState
            lhs147_ = d_1904_v142_
            lhs148_ = default__.safeIndex(983, (d_1904_v142_).length(0))
            d_1907_v145_ = rhs262_
            r1 = rhs263_
            lhs145_.f11 = rhs264_
            lhs146_.f8 = rhs265_
            lhs147_[lhs148_] = rhs266_
            (self).m13(((self).f12) + (p1), (d_1856_v110_)[default__.safeIndex(273, (d_1856_v110_).length(0))], p3, globalState)
        r0 = d_1856_v110_
        r1 = (p1) + (((self).f12) + (p2))
        d_1908_v147_: _dafny.Set
        d_1908_v147_ = _dafny.Set({(self).f13})
        def iife171_():
            coll79_ = _dafny.Set()
            compr_79_: int
            for compr_79_ in _dafny.IntegerRange(552, 597):
                d_1909_v146_: int = compr_79_
                if ((552) <= (d_1909_v146_)) and ((d_1909_v146_) < (597)):
                    coll79_ = coll79_.union(_dafny.Set([default__.safeDivisionInt(d_1909_v146_, len(d_1720_v0_))]))
            return _dafny.Set(coll79_)
        def iife172_():
            coll80_ = _dafny.Set()
            compr_80_: int
            for compr_80_ in (d_1725_v5_).Elements:
                d_1910_v148_: int = compr_80_
                if (d_1910_v148_) in (d_1725_v5_):
                    coll80_ = coll80_.union(_dafny.Set([default__.safeModuloInt(d_1910_v148_, len(_dafny.Map({True: not(True)})))]))
            return _dafny.Set(coll80_)
        r2 = ((iife171_()
        ).intersection(d_1908_v147_)) | (iife172_()
        )
        return r0, r1, r2

    def m12(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        rhs267_ = not(False)
        rhs268_ = (self.f22) * (default__.safeModuloInt(self.f22, self.f22))
        lhs149_ = self
        lhs150_ = self
        lhs149_.f11 = rhs267_
        lhs150_.f22 = rhs268_
        d_1911_v0_: _dafny.Array
        nw286_ = _dafny.Array(False, 7)
        d_1911_v0_ = nw286_
        index256_ = default__.safeIndex(646, (p2).length(0))
        (p2)[index256_] = d_1911_v0_
        d_1912_v1_: _dafny.Seq
        d_1912_v1_ = _dafny.SeqWithoutIsStrInference([(False) and (self.f11)])
        index257_ = default__.safeIndex(646, (p2).length(0))
        rhs269_ = not(self.f11)
        rhs270_ = (d_1912_v1_)[default__.safeIndex(default__.safeDivisionInt((self).f12, self.f22), len(d_1912_v1_))]
        rhs271_ = d_1911_v0_
        lhs151_ = p2
        lhs152_ = default__.safeIndex(646, (p2).length(0))
        r0 = rhs269_
        r1 = rhs270_
        lhs151_[lhs152_] = rhs271_
        d_1913_v2_: _dafny.Array
        nw287_ = _dafny.Array(D3.default()(), 15)
        d_1913_v2_ = nw287_
        d_1914_v3_: _dafny.Map
        d_1914_v3_ = _dafny.Map({self.f11: (self).f13})
        d_1915_v4_: D3
        d_1915_v4_ = D3_DC8(d_1914_v3_)
        index258_ = default__.safeIndex(425, (d_1913_v2_).length(0))
        (d_1913_v2_)[index258_] = d_1915_v4_
        index259_ = default__.safeIndex(425, (d_1913_v2_).length(0))
        (d_1913_v2_)[index259_] = d_1915_v4_
        d_1916_v5_: _dafny.Map
        d_1916_v5_ = _dafny.Map({(self).f13: _dafny.Set({(self).f12, self.f22, self.f22})})
        d_1917_v6_: _dafny.Seq
        d_1917_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f22])])
        d_1918_v7_: _dafny.Set
        d_1918_v7_ = _dafny.Set({(0) - (len(d_1912_v1_)), (self).f12})
        d_1916_v5_ = (d_1916_v5_).set(((d_1917_v6_)[default__.safeIndex((self).f12, len(d_1917_v6_))]).cardinality, (d_1918_v7_) | (_dafny.Set({835, (self).f12})))
        hi7_ = (0) - (((0) - ((self).f13)) - ((self).f13))
        for d_1919_i0_ in range((self).f13, hi7_):
            d_1920_v8_: str
            d_1920_v8_ = _dafny.CodePoint('b')
            d_1920_v8_ = d_1920_v8_
            if False:
                d_1921_v9_: C7
                nw288_ = C7()
                nw288_.ctor__(self.f22, 868)
                d_1921_v9_ = nw288_
                (self).f22 = default__.fm2(globalState)
                r1 = not (((self).f13) < ((self).f13)) or (self.f11)
                d_1922_v10_: _dafny.MultiSet
                d_1922_v10_ = _dafny.MultiSet([self.f11, self.f11])
                d_1923_v11_: D12
                d_1923_v11_ = D12_DC32(_dafny.Map({False: d_1922_v10_}))
                d_1923_v11_ = d_1923_v11_
                arr2_ = (p2)[default__.safeIndex(646, (p2).length(0))]
                index260_ = default__.safeIndex(428, ((p2)[default__.safeIndex(646, (p2).length(0))]).length(0))
                arr2_[index260_] = (d_1918_v7_).issubset(_dafny.Set({(self).f13}))
                arr3_ = (p2)[default__.safeIndex(646, (p2).length(0))]
                index261_ = default__.safeIndex(428, ((p2)[default__.safeIndex(646, (p2).length(0))]).length(0))
                arr3_[index261_] = self.f11
            elif True:
                (self).f22 = 368
                d_1924_v12_: _dafny.Map
                d_1924_v12_ = _dafny.Map({len(d_1912_v1_): self.f22})
                d_1925_v13_: _dafny.Map
                d_1925_v13_ = _dafny.Map({False: (d_1924_v12_).set((self).f13, d_1919_i0_)})
                rhs272_ = d_1920_v8_
                rhs273_ = default__.fm54((not(False) if (d_1912_v1_)[default__.safeIndex((0) - (len(d_1918_v7_)), len(d_1912_v1_))] else self.f11), globalState)
                d_1920_v8_ = rhs272_
                d_1925_v13_ = rhs273_
                d_1926_v14_: _dafny.Array
                nw289_ = _dafny.Array(int(0), 25)
                d_1926_v14_ = nw289_
                index262_ = default__.safeIndex(531, (d_1926_v14_).length(0))
                (d_1926_v14_)[index262_] = default__.safeDivisionInt(len((D0_DC1(382, p0, (self).f13, True)).cf2), self.f22)
                d_1927_v15_: _dafny.Array
                nw290_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_1927_v15_ = nw290_
                d_1928_v16_: _dafny.Array
                def lambda107_(d_1929_i1_):
                    return _dafny.MultiSet([(self).f12])

                init59_ = lambda107_
                nw291_ = _dafny.Array(None, 26)
                for i0_59_ in range(nw291_.length(0)):
                    nw291_[i0_59_] = init59_(i0_59_)
                d_1928_v16_ = nw291_
                index263_ = default__.safeIndex(566, (d_1927_v15_).length(0))
                (d_1927_v15_)[index263_] = d_1928_v16_
                index264_ = default__.safeIndex(340, (d_1926_v14_).length(0))
                (d_1926_v14_)[index264_] = (self).f12
                index265_ = default__.safeIndex(531, (d_1926_v14_).length(0))
                index266_ = default__.safeIndex(566, (d_1927_v15_).length(0))
                index267_ = default__.safeIndex(340, (d_1926_v14_).length(0))
                rhs274_ = (self).f13
                rhs275_ = (default__.fm50(globalState)) == (d_1918_v7_)
                rhs276_ = d_1928_v16_
                rhs277_ = d_1920_v8_
                rhs278_ = default__.fm2(globalState)
                lhs153_ = d_1926_v14_
                lhs154_ = default__.safeIndex(531, (d_1926_v14_).length(0))
                lhs155_ = self
                lhs156_ = d_1927_v15_
                lhs157_ = default__.safeIndex(566, (d_1927_v15_).length(0))
                lhs158_ = d_1926_v14_
                lhs159_ = default__.safeIndex(340, (d_1926_v14_).length(0))
                lhs153_[lhs154_] = rhs274_
                lhs155_.f11 = rhs275_
                lhs156_[lhs157_] = rhs276_
                d_1920_v8_ = rhs277_
                lhs158_[lhs159_] = rhs278_
                d_1920_v8_ = default__.fm45(d_1919_i0_, globalState)
                d_1930_v17_: _dafny.Array
                nw292_ = _dafny.Array(_dafny.Seq({}), 4)
                d_1930_v17_ = nw292_
                d_1931_v18_: _dafny.Seq
                d_1931_v18_ = _dafny.SeqWithoutIsStrInference([(self).f13, 653])
                index268_ = default__.safeIndex(542, (d_1930_v17_).length(0))
                (d_1930_v17_)[index268_] = d_1931_v18_
                index269_ = default__.safeIndex(542, (d_1930_v17_).length(0))
                (d_1930_v17_)[index269_] = d_1931_v18_
            if self.f11:
                d_1932_v19_: _dafny.Map
                d_1932_v19_ = _dafny.Map({d_1920_v8_: _dafny.CodePoint('g')})
                d_1932_v19_ = (d_1932_v19_).set(d_1920_v8_, _dafny.CodePoint('e'))
                d_1933_v20_: _dafny.Array
                nw293_ = _dafny.Array(_dafny.Seq({}), 19)
                d_1933_v20_ = nw293_
                d_1934_v21_: _dafny.Seq
                d_1934_v21_ = _dafny.SeqWithoutIsStrInference([self.f22, self.f22])
                d_1935_v22_: _dafny.Array
                nw294_ = _dafny.Array(int(0), 10)
                d_1935_v22_ = nw294_
                d_1936_v23_: _dafny.MultiSet
                d_1936_v23_ = _dafny.MultiSet([d_1935_v22_])
                d_1937_v24_: _dafny.Map
                d_1937_v24_ = _dafny.Map({(self).f13: (0) - ((d_1936_v23_).cardinality)})
                d_1938_v25_: _dafny.Set
                d_1938_v25_ = _dafny.Set({(self).f20})
                d_1939_v26_: _dafny.Array
                nw295_ = _dafny.Array(None, 22)
                nw295_[int(0)] = default__.fm2(globalState)
                nw295_[int(1)] = len(d_1914_v3_)
                nw295_[int(2)] = (self).f12
                nw295_[int(3)] = (self).f12
                nw295_[int(4)] = (self).f12
                nw295_[int(5)] = len(d_1934_v21_)
                nw295_[int(6)] = (self).f12
                nw295_[int(7)] = (self).f13
                nw295_[int(8)] = (self).f13
                nw295_[int(9)] = self.f22
                nw295_[int(10)] = 992
                nw295_[int(11)] = d_1919_i0_
                nw295_[int(12)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1940_i2_ in range(default__.abs(-454))]))
                nw295_[int(13)] = d_1919_i0_
                nw295_[int(14)] = (self).f12
                nw295_[int(15)] = (self).f13
                nw295_[int(16)] = self.f22
                nw295_[int(17)] = len(d_1937_v24_)
                nw295_[int(18)] = len(d_1938_v25_)
                nw295_[int(19)] = len((self).f20)
                nw295_[int(20)] = (self).f12
                nw295_[int(21)] = len(d_1912_v1_)
                d_1939_v26_ = nw295_
                d_1941_v27_: _dafny.Seq
                d_1941_v27_ = _dafny.SeqWithoutIsStrInference([d_1939_v26_])
                d_1942_v28_: _dafny.Map
                d_1942_v28_ = _dafny.Map({0: d_1941_v27_})
                index270_ = default__.safeIndex(993, (d_1933_v20_).length(0))
                (d_1933_v20_)[index270_] = ((d_1942_v28_)[(self).f13] if ((self).f13) in (d_1942_v28_) else (d_1941_v27_).set(default__.safeIndex(-133, len(d_1941_v27_)), d_1935_v22_))
                index271_ = default__.safeIndex(993, (d_1933_v20_).length(0))
                (d_1933_v20_)[index271_] = (d_1941_v27_) + (d_1941_v27_)
                d_1943_v29_: _dafny.Map
                d_1943_v29_ = _dafny.Map({d_1914_v3_: not(self.f11)})
                d_1944_v31_: D5
                d_1944_v31_ = D5_DC15(d_1919_i0_, d_1919_i0_, (self).f13)
                d_1945_v32_: D6
                d_1945_v32_ = D6_DC16(_dafny.Map({(self).f13: d_1944_v31_}))
                d_1946_v33_: D6
                d_1946_v33_ = D6_DC20(d_1945_v32_)
                d_1947_v34_: D6
                d_1947_v34_ = D6_DC20(d_1945_v32_)
                d_1948_v35_: _dafny.Seq
                d_1948_v35_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f11: (self).f12}), d_1914_v3_])
                d_1949_v36_: _dafny.Map
                d_1949_v36_ = _dafny.Map({d_1947_v34_: d_1948_v35_})
                def iife173_():
                    coll81_ = _dafny.Map()
                    compr_81_: _dafny.Map
                    for compr_81_ in (((d_1949_v36_)[d_1947_v34_] if (d_1947_v34_) in (d_1949_v36_) else _dafny.SeqWithoutIsStrInference([d_1914_v3_]))).Elements:
                        d_1950_v30_: _dafny.Map = compr_81_
                        if (d_1950_v30_) in (((d_1949_v36_)[d_1947_v34_] if (d_1947_v34_) in (d_1949_v36_) else _dafny.SeqWithoutIsStrInference([d_1914_v3_]))):
                            coll81_[d_1950_v30_] = (d_1919_i0_) in (d_1934_v21_)
                    return _dafny.Map(coll81_)
                d_1943_v29_ = iife173_()
                
                def iife174_():
                    coll82_ = _dafny.Set()
                    compr_82_: int
                    for compr_82_ in (d_1937_v24_).keys.Elements:
                        d_1951_v37_: int = compr_82_
                        if (d_1951_v37_) in (d_1937_v24_):
                            coll82_ = coll82_.union(_dafny.Set([(d_1951_v37_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1952_i3_ in range(default__.abs(201))]))))]))
                    return _dafny.Set(coll82_)
                d_1918_v7_ = (d_1918_v7_).intersection(iife174_()
                )
                d_1953_v38_: _dafny.Array
                nw296_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_1953_v38_ = nw296_
                d_1953_v38_ = d_1953_v38_
            elif True:
                d_1954_v39_: _dafny.Map
                d_1954_v39_ = _dafny.Map({default__.fm2(globalState): self.f11})
                (self).f22 = len(d_1954_v39_)
                index272_ = default__.safeIndex(64, (d_1911_v0_).length(0))
                (d_1911_v0_)[index272_] = self.f11
                index273_ = default__.safeIndex(64, (d_1911_v0_).length(0))
                (d_1911_v0_)[index273_] = (True if self.f11 else (self.f11 if self.f11 else self.f11))
                d_1955_v40_: _dafny.Array
                nw297_ = _dafny.Array(int(0), 3)
                d_1955_v40_ = nw297_
                d_1956_v41_: _dafny.Seq
                d_1956_v41_ = _dafny.SeqWithoutIsStrInference([d_1955_v40_])
                d_1957_v42_: _dafny.Array
                nw298_ = _dafny.Array(None, 9)
                nw298_[int(0)] = d_1955_v40_
                nw298_[int(1)] = (d_1956_v41_)[default__.safeIndex(self.f22, len(d_1956_v41_))]
                nw298_[int(2)] = d_1955_v40_
                nw298_[int(3)] = d_1955_v40_
                nw298_[int(4)] = d_1955_v40_
                nw298_[int(5)] = d_1955_v40_
                nw298_[int(6)] = d_1955_v40_
                nw298_[int(7)] = d_1955_v40_
                nw298_[int(8)] = d_1955_v40_
                d_1957_v42_ = nw298_
                d_1958_v43_: D9
                d_1958_v43_ = D9_DC27(self.f11, D7_DC22((0) - ((D5_DC15(len(d_1918_v7_), d_1919_i0_, (self).f12)).cf27), d_1957_v42_), (self.f22) * ((self).f13), False)
                d_1958_v43_ = d_1958_v43_
                (globalState).f8 = (d_1911_v0_)[default__.safeIndex(64, (d_1911_v0_).length(0))]
                d_1959_v44_: _dafny.Array
                d_1960_v45_: bool
                out28_: _dafny.Array
                out29_: bool
                out28_, out29_ = default__.m0(globalState)
                d_1959_v44_ = out28_
                d_1960_v45_ = out29_
            if (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1919_i0_ for d_1961_i4_ in range(default__.abs(271))]))).cardinality) * ((self).f13)) > (default__.fm2(globalState)):
                d_1962_v46_: _dafny.Array
                nw299_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_1962_v46_ = nw299_
                d_1962_v46_ = p2
                (self).f22 = (self).f13
                arr4_ = (p2)[default__.safeIndex(646, (p2).length(0))]
                index274_ = default__.safeIndex(863, ((p2)[default__.safeIndex(646, (p2).length(0))]).length(0))
                arr4_[index274_] = self.f11
                arr5_ = (p2)[default__.safeIndex(646, (p2).length(0))]
                index275_ = default__.safeIndex(863, ((p2)[default__.safeIndex(646, (p2).length(0))]).length(0))
                arr5_[index275_] = not (True) or (True)
                (self).f22 = self.f22
                d_1963_v47_: _dafny.Map
                d_1963_v47_ = _dafny.Map({d_1911_v0_: (len(p0)) * (self.f22)})
                d_1963_v47_ = ((d_1963_v47_) | (d_1963_v47_)) | (_dafny.Map({d_1911_v0_: 940}))
            elif True:
                (globalState).f8 = self.f11
                (self).f22 = (self).f12
                d_1964_v48_: _dafny.Array
                nw300_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_1964_v48_ = nw300_
                index276_ = default__.safeIndex(694, (d_1964_v48_).length(0))
                (d_1964_v48_)[index276_] = p0
                d_1965_v49_: C1
                nw301_ = C1()
                nw301_.ctor__((self).f12)
                d_1965_v49_ = nw301_
                d_1966_v50_: _dafny.Seq
                d_1966_v50_ = _dafny.SeqWithoutIsStrInference([d_1965_v49_, d_1965_v49_, d_1965_v49_])
                index277_ = default__.safeIndex(646, (p2).length(0))
                index278_ = default__.safeIndex(694, (d_1964_v48_).length(0))
                rhs279_ = (p2)[default__.safeIndex(646, (p2).length(0))]
                rhs280_ = (self).f20
                rhs281_ = _dafny.SeqWithoutIsStrInference([d_1965_v49_, d_1965_v49_, d_1965_v49_])
                rhs282_ = -688
                lhs160_ = p2
                lhs161_ = default__.safeIndex(646, (p2).length(0))
                lhs162_ = d_1964_v48_
                lhs163_ = default__.safeIndex(694, (d_1964_v48_).length(0))
                lhs164_ = d_1965_v49_
                lhs160_[lhs161_] = rhs279_
                lhs162_[lhs163_] = rhs280_
                d_1966_v50_ = rhs281_
                lhs164_.f27 = rhs282_
                (globalState).f8 = (self.f11) == (self.f11)
                rhs283_ = (self.f11) or (self.f11)
                rhs284_ = (0) - ((self).f12)
                lhs165_ = d_1965_v49_
                r1 = rhs283_
                lhs165_.f27 = rhs284_
        hi8_ = 157
        for d_1967_i5_ in range((self).f12, hi8_):
            d_1968_v51_: _dafny.Array
            d_1969_v52_: bool
            out30_: _dafny.Array
            out31_: bool
            out30_, out31_ = default__.m0(globalState)
            d_1968_v51_ = out30_
            d_1969_v52_ = out31_
            (self).f22 = -302
            d_1970_v53_: _dafny.Seq
            d_1970_v53_ = _dafny.SeqWithoutIsStrInference([(self).f12, d_1967_i5_])
            d_1971_v54_: str
            d_1971_v54_ = _dafny.CodePoint('j')
            d_1972_v55_: _dafny.Array
            nw302_ = _dafny.Array(None, 13)
            nw302_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omgjypkmn"))
            nw302_[int(1)] = (self).f20
            nw302_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpqqegm"))
            nw302_[int(3)] = default__.fm36((self).f12, len(d_1970_v53_), (D9_DC26(d_1971_v54_)).cf48, globalState)
            nw302_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1971_v54_ for d_1973_i6_ in range(default__.abs(853))])
            nw302_[int(5)] = (self).f20
            nw302_[int(6)] = p0
            nw302_[int(7)] = (self).f20
            nw302_[int(8)] = (self).f20
            nw302_[int(9)] = (self).f20
            nw302_[int(10)] = _dafny.SeqWithoutIsStrInference([d_1971_v54_ for d_1974_i7_ in range(default__.abs(767))])
            nw302_[int(11)] = p0
            nw302_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnwrb"))
            d_1972_v55_ = nw302_
            d_1975_v56_: D5
            d_1975_v56_ = D5_DC14((D5_DC14(d_1972_v55_)).cf25)
            d_1976_v57_: _dafny.MultiSet
            d_1976_v57_ = _dafny.MultiSet([d_1967_i5_])
            d_1977_v58_: _dafny.Map
            d_1977_v58_ = _dafny.Map({_dafny.MultiSet([d_1969_v52_]): d_1976_v57_})
            d_1978_v60_: _dafny.MultiSet
            d_1978_v60_ = _dafny.MultiSet([False, not((self).fm6(self.f11, d_1967_i5_, d_1967_i5_, globalState))])
            d_1979_v61_: _dafny.Set
            d_1979_v61_ = _dafny.Set({d_1978_v60_})
            d_1980_v62_: _dafny.Map
            def iife175_():
                coll83_ = _dafny.Map()
                compr_83_: _dafny.MultiSet
                for compr_83_ in (d_1979_v61_).Elements:
                    d_1981_v59_: _dafny.MultiSet = compr_83_
                    if (d_1981_v59_) in (d_1979_v61_):
                        coll83_[d_1981_v59_] = d_1976_v57_
                return _dafny.Map(coll83_)
            d_1980_v62_ = _dafny.Map({d_1975_v56_: (d_1977_v58_) == (iife175_()
            )})
            d_1980_v62_ = (d_1980_v62_) | (d_1980_v62_)
            d_1969_v52_ = (d_1969_v52_ if self.f11 else self.f11)
        r0 = self.f11
        d_1982_v63_: _dafny.Seq
        d_1982_v63_ = _dafny.SeqWithoutIsStrInference([self.f22, (self).f12, 174, (0) - ((self).f13)])
        d_1983_v64_: D4
        d_1983_v64_ = D4_DC12(self.f11, default__.fm2(globalState), len(d_1982_v63_))
        r1 = (d_1983_v64_).cf21
        return r0, r1

    def m13(self, p0, p1, p2, globalState):
        d_1984_v0_: _dafny.Seq
        d_1984_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        index279_ = default__.safeIndex(821, (p2).length(0))
        (p2)[index279_] = len((d_1984_v0_) + (_dafny.SeqWithoutIsStrInference([self.f11, self.f11])))
        index280_ = default__.safeIndex(821, (p2).length(0))
        (p2)[index280_] = default__.fm2(globalState)
        d_1985_v1_: _dafny.Set
        d_1985_v1_ = _dafny.Set({313})
        d_1986_v2_: _dafny.Seq
        d_1986_v2_ = _dafny.SeqWithoutIsStrInference([d_1985_v1_, d_1985_v1_, d_1985_v1_, d_1985_v1_, d_1985_v1_])
        d_1987_v3_: _dafny.Set
        d_1987_v3_ = (d_1986_v2_)[default__.safeIndex(-741, len(d_1986_v2_))]
        source30_ = d_1987_v3_
        d_1988___mcc_h0_ = source30_
        d_1989_cf54_ = d_1988___mcc_h0_
        d_1990_v4_: _dafny.MultiSet
        d_1990_v4_ = _dafny.MultiSet([p1])
        d_1991_v5_: D7
        d_1991_v5_ = D7_DC21(d_1990_v4_)
        d_1992_v6_: D7
        d_1992_v6_ = D7_DC24(d_1991_v5_)
        pat_let_tv70_ = d_1991_v5_
        def iife176_(_pat_let46_0):
            def iife177_(d_1993_dt__update__tmp_h0_):
                def iife178_(_pat_let47_0):
                    def iife179_(d_1994_dt__update_hcf46_h0_):
                        return D7_DC24(d_1994_dt__update_hcf46_h0_)
                    return iife179_(_pat_let47_0)
                return iife178_(pat_let_tv70_)
            return iife177_(_pat_let46_0)
        d_1992_v6_ = iife176_(d_1992_v6_)
        if False:
            (globalState).f8 = (self.f11) == (self.f11)
            d_1995_v7_: C0
            nw303_ = C0()
            nw303_.ctor__((self).f13, p1, (self).f13, self.f22)
            d_1995_v7_ = nw303_
            d_1996_v8_: _dafny.Map
            d_1996_v8_ = _dafny.Map({p1: d_1995_v7_})
            d_1997_v9_: _dafny.Set
            d_1997_v9_ = _dafny.Set({d_1996_v8_})
            d_1998_v10_: _dafny.Map
            d_1998_v10_ = _dafny.Map({True: d_1997_v9_})
            d_1999_v11_: _dafny.MultiSet
            d_1999_v11_ = _dafny.MultiSet([self.f22])
            d_2000_v12_: _dafny.Map
            d_2000_v12_ = _dafny.Map({self.f11: d_1999_v11_})
            d_2001_v13_: _dafny.Array
            nw304_ = _dafny.Array(None, 13)
            nw304_[int(0)] = True
            nw304_[int(1)] = self.f11
            nw304_[int(2)] = ((p2)[default__.safeIndex(821, (p2).length(0))]) == (self.f22)
            nw304_[int(3)] = (d_1996_v8_) not in (((d_1998_v10_)[d_1995_v7_.f29] if (d_1995_v7_.f29) in (d_1998_v10_) else d_1997_v9_))
            nw304_[int(4)] = (_dafny.MultiSet([(d_1995_v7_).f28])).issubset(((d_2000_v12_)[self.f11] if (self.f11) in (d_2000_v12_) else d_1999_v11_))
            nw304_[int(5)] = self.f11
            nw304_[int(6)] = p1
            nw304_[int(7)] = not (p1) or (not(p1))
            nw304_[int(8)] = (self.f22) <= ((self).f12)
            nw304_[int(9)] = (_dafny.CodePoint('d')) in ((self).f20)
            nw304_[int(10)] = ((self).f12) != ((self).f12)
            nw304_[int(11)] = ((self).f13) <= ((self).f12)
            nw304_[int(12)] = False
            d_2001_v13_ = nw304_
            index281_ = default__.safeIndex(979, (d_2001_v13_).length(0))
            (d_2001_v13_)[index281_] = False
            index282_ = default__.safeIndex(979, (d_2001_v13_).length(0))
            (d_2001_v13_)[index282_] = not(p1)
            d_2002_v14_: _dafny.Map
            d_2002_v14_ = _dafny.Map({(_dafny.CodePoint('w')) not in ((self).f20): 35})
            d_2003_v15_: _dafny.Map
            d_2003_v15_ = _dafny.Map({d_1995_v7_.f29: p1})
            d_2002_v14_ = (d_2002_v14_).set(((d_2003_v15_)[True] if (True) in (d_2003_v15_) else d_1995_v7_.f29), ((self).f13 if d_1995_v7_.f29 else self.f22))
            d_2004_v16_: _dafny.Seq
            d_2004_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pybpnsivf"))
            d_2004_v16_ = (self).f20
            index283_ = default__.safeIndex(821, (p2).length(0))
            (p2)[index283_] = default__.fm2(globalState)
        elif True:
            index284_ = default__.safeIndex(821, (p2).length(0))
            (p2)[index284_] = default__.fm2(globalState)
            d_2005_v17_: C8
            nw305_ = C8()
            nw305_.ctor__(_dafny.SeqWithoutIsStrInference([(self).f12 for d_2006_i0_ in range(default__.abs(305))]))
            d_2005_v17_ = nw305_
            (globalState).f8 = not(p1)
            d_2007_v18_: _dafny.Seq
            d_2007_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlq"))
            d_2008_v19_: _dafny.MultiSet
            d_2008_v19_ = _dafny.MultiSet([457])
            d_2009_v20_: _dafny.MultiSet
            d_2009_v20_ = _dafny.MultiSet([((d_2008_v19_)[self.f22] if (self.f22) in (d_2008_v19_) else 737)])
            d_2010_v21_: D11
            d_2010_v21_ = D11_DC31(self.f11)
            d_2011_v22_: _dafny.Map
            d_2011_v22_ = _dafny.Map({d_2010_v21_: (p2)[default__.safeIndex(821, (p2).length(0))]})
            index285_ = default__.safeIndex(821, (p2).length(0))
            index286_ = default__.safeIndex(821, (p2).length(0))
            rhs285_ = (p2)[default__.safeIndex(821, (p2).length(0))]
            rhs286_ = (p2)[default__.safeIndex(821, (p2).length(0))]
            rhs287_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uk"))
            rhs288_ = self.f11
            rhs289_ = default__.safeDivisionInt(((d_2009_v20_)[len(d_2011_v22_)] if (len(d_2011_v22_)) in (d_2009_v20_) else len(d_2007_v18_)), ((self).f13) - (p0))
            lhs166_ = p2
            lhs167_ = default__.safeIndex(821, (p2).length(0))
            lhs168_ = self
            lhs169_ = self
            lhs170_ = p2
            lhs171_ = default__.safeIndex(821, (p2).length(0))
            lhs166_[lhs167_] = rhs285_
            lhs168_.f22 = rhs286_
            d_2007_v18_ = rhs287_
            lhs169_.f11 = rhs288_
            lhs170_[lhs171_] = rhs289_
            (self).f22 = default__.safeModuloInt((self).f13, default__.safeModuloInt(p0, p0))
        d_2012_v23_: _dafny.Map
        d_2012_v23_ = _dafny.Map({not(self.f11): (p2)[default__.safeIndex(821, (p2).length(0))]})
        (globalState).f8 = (p1) in (d_2012_v23_)
        d_2013_v24_: _dafny.Map
        d_2013_v24_ = _dafny.Map({d_1990_v4_: p1})
        d_2014_v25_: _dafny.MultiSet
        d_2014_v25_ = _dafny.MultiSet([(p2)[default__.safeIndex(821, (p2).length(0))], len(d_2013_v24_), (self).f12, ((self).f13 if p1 else len((self).f20))])
        (self).f22 = ((d_2014_v25_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_2014_v25_) else default__.safeDivisionInt(38, p0))
        hi9_ = (self).f13
        for d_2015_i1_ in range((self).f13, hi9_):
            d_2016_v26_: _dafny.MultiSet
            d_2016_v26_ = _dafny.MultiSet([(p2)[default__.safeIndex(821, (p2).length(0))], (p2)[default__.safeIndex(821, (p2).length(0))]])
            d_2017_v27_: str
            d_2017_v27_ = _dafny.CodePoint('j')
            d_2018_v28_: _dafny.Map
            d_2018_v28_ = _dafny.Map({d_2017_v27_: self.f11})
            d_2019_v29_: _dafny.Map
            d_2019_v29_ = _dafny.Map({p1: d_2015_i1_})
            d_2020_v30_: _dafny.Seq
            d_2020_v30_ = _dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(821, (p2).length(0))], (self).f12])
            d_2021_v31_: _dafny.Map
            d_2021_v31_ = _dafny.Map({self.f11: _dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(821, (p2).length(0))] for d_2022_i2_ in range(default__.abs(569))])})
            d_2023_v32_: _dafny.Map
            d_2023_v32_ = _dafny.Map({(self).f13: (self).f12})
            d_2024_v33_: _dafny.MultiSet
            d_2024_v33_ = _dafny.MultiSet([d_2023_v32_, d_2023_v32_])
            d_2025_v34_: _dafny.Map
            d_2025_v34_ = _dafny.Map({(p2)[default__.safeIndex(821, (p2).length(0))]: p1})
            d_2026_v35_: _dafny.Set
            d_2026_v35_ = _dafny.Set({p1})
            d_2027_v36_: _dafny.Array
            nw306_ = _dafny.Array(None, 23)
            nw306_[int(0)] = self.f11
            nw306_[int(1)] = False
            nw306_[int(2)] = (_dafny.MultiSet([(p2)[default__.safeIndex(821, (p2).length(0))], (self).f12])).ispropersubset(d_2016_v26_)
            nw306_[int(3)] = self.f11
            nw306_[int(4)] = self.f11
            nw306_[int(5)] = ((d_2018_v28_)[d_2017_v27_] if (d_2017_v27_) in (d_2018_v28_) else self.f11)
            nw306_[int(6)] = (len(d_2019_v29_)) > ((p2)[default__.safeIndex(821, (p2).length(0))])
            nw306_[int(7)] = (d_2020_v30_) != (d_2020_v30_)
            nw306_[int(8)] = True
            nw306_[int(9)] = (len(d_1984_v0_)) not in (((d_2021_v31_)[self.f11] if (self.f11) in (d_2021_v31_) else d_2020_v30_))
            nw306_[int(10)] = p1
            nw306_[int(11)] = self.f11
            nw306_[int(12)] = not (p1) or (self.f11)
            nw306_[int(13)] = self.f11
            nw306_[int(14)] = p1
            nw306_[int(15)] = True
            nw306_[int(16)] = default__.fm1(p0, self.f11, (self).f13, globalState)
            nw306_[int(17)] = not(p1)
            nw306_[int(18)] = (d_2024_v33_).issubset(d_2024_v33_)
            nw306_[int(19)] = ((d_2025_v34_)[len(_dafny.Map({p1: (p2)[default__.safeIndex(821, (p2).length(0))]}))] if (len(_dafny.Map({p1: (p2)[default__.safeIndex(821, (p2).length(0))]}))) in (d_2025_v34_) else self.f11)
            nw306_[int(20)] = (d_2026_v35_) != (d_2026_v35_)
            nw306_[int(21)] = not((d_2023_v32_) != ((d_2023_v32_).set(320, len((self).f20))))
            nw306_[int(22)] = p1
            d_2027_v36_ = nw306_
            d_2028_v37_: _dafny.Array
            nw307_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_2028_v37_ = nw307_
            d_2029_v38_: D7
            d_2029_v38_ = D7_DC23(d_2028_v37_, d_1984_v0_, p1, (self).f12)
            index287_ = default__.safeIndex(805, (d_2027_v36_).length(0))
            (d_2027_v36_)[index287_] = not ((d_2029_v38_).cf44) or (True)
            index288_ = default__.safeIndex(805, (d_2027_v36_).length(0))
            (d_2027_v36_)[index288_] = p1
            index289_ = default__.safeIndex(821, (p2).length(0))
            (p2)[index289_] = self.f22
            d_2019_v29_ = (d_2019_v29_).set((d_2027_v36_)[default__.safeIndex(805, (d_2027_v36_).length(0))], ((0) - ((self).f13)) - ((self).f13))
            index290_ = default__.safeIndex(821, (p2).length(0))
            rhs290_ = (len((self).f20)) + ((308) - (len((self).fm8((0) - (len((self).f20)), _dafny.Map({d_2017_v27_: p1}), d_1985_v1_, globalState))))
            rhs291_ = ((len((self).f20)) * ((self).f13)) - (default__.fm2(globalState))
            rhs292_ = self.f11
            lhs172_ = p2
            lhs173_ = default__.safeIndex(821, (p2).length(0))
            lhs174_ = self
            lhs175_ = globalState
            lhs172_[lhs173_] = rhs290_
            lhs174_.f22 = rhs291_
            lhs175_.f8 = rhs292_
        d_2030_v39_: _dafny.Array
        nw308_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
        d_2030_v39_ = nw308_
        d_2031_v40_: D5
        d_2031_v40_ = D5_DC14(d_2030_v39_)
        d_2032_v41_: _dafny.MultiSet
        d_2032_v41_ = _dafny.MultiSet([d_2031_v40_])
        d_2033_v42_: _dafny.Set
        d_2033_v42_ = _dafny.Set({_dafny.MultiSet([d_2031_v40_, d_2031_v40_]), d_2032_v41_})
        d_2034_v43_: _dafny.Seq
        d_2034_v43_ = _dafny.SeqWithoutIsStrInference([(0) - ((p2)[default__.safeIndex(821, (p2).length(0))])])
        d_2035_v44_: _dafny.MultiSet
        d_2035_v44_ = _dafny.MultiSet([(self).f20])
        d_2036_v45_: _dafny.Set
        d_2036_v45_ = _dafny.Set({self.f11})
        d_2037_v46_: T0
        nw309_ = C4()
        nw309_.ctor__(531, p1)
        d_2037_v46_ = nw309_
        d_2038_v47_: _dafny.Map
        d_2038_v47_ = _dafny.Map({d_2037_v46_: (p2)[default__.safeIndex(821, (p2).length(0))]})
        d_2039_v48_: _dafny.Array
        nw310_ = _dafny.Array(None, 20)
        nw310_[int(0)] = p0
        nw310_[int(1)] = (0) - (len(d_2033_v42_))
        nw310_[int(2)] = default__.safeDivisionInt((self).f13, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggj"))))
        nw310_[int(3)] = (0) - (default__.fm2(globalState))
        nw310_[int(4)] = (len(d_2034_v43_)) - (((d_2035_v44_)[(self).f20] if ((self).f20) in (d_2035_v44_) else (self).f13))
        nw310_[int(5)] = self.f22
        nw310_[int(6)] = (0) - (p0)
        nw310_[int(7)] = default__.safeDivisionInt((self).f13, len(d_2036_v45_))
        nw310_[int(8)] = (self).f13
        nw310_[int(9)] = (self).f13
        nw310_[int(10)] = (self).f13
        nw310_[int(11)] = (0) - ((p2)[default__.safeIndex(821, (p2).length(0))])
        nw310_[int(12)] = default__.safeDivisionInt((self).f13, p0)
        nw310_[int(13)] = (self).f13
        nw310_[int(14)] = (self).f12
        nw310_[int(15)] = len((self).f20)
        nw310_[int(16)] = self.f22
        nw310_[int(17)] = p0
        nw310_[int(18)] = ((d_2038_v47_)[d_2037_v46_] if (d_2037_v46_) in (d_2038_v47_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uidytrhx"))))
        nw310_[int(19)] = p0
        d_2039_v48_ = nw310_
        d_2039_v48_ = d_2039_v48_
        if self.f11:
            index291_ = default__.safeIndex(821, (p2).length(0))
            (p2)[index291_] = self.f22
            d_2040_v49_: str
            d_2040_v49_ = _dafny.CodePoint('i')
            d_2040_v49_ = d_2040_v49_
            (globalState).f8 = (len(default__.fm36(self.f22, (self).f13, d_2040_v49_, globalState))) != ((0) - ((_dafny.MultiSet([p1])).cardinality))
            d_2041_v50_: _dafny.Map
            d_2041_v50_ = _dafny.Map({len((self).f20): d_2037_v46_.f11})
            d_2041_v50_ = (d_2041_v50_).set(default__.safeDivisionInt(default__.fm2(globalState), (self).f13), (d_2037_v46_.f11 if d_2037_v46_.f11 else not(self.f11)))
            d_2042_v51_: D5
            d_2042_v51_ = D5_DC15((self).f12, (self).f12, p0)
            d_2043_v52_: _dafny.Map
            d_2043_v52_ = _dafny.Map({(self).f13: d_2042_v51_})
            d_2044_v53_: D6
            d_2044_v53_ = D6_DC16(d_2043_v52_)
            d_2045_v54_: _dafny.Set
            d_2045_v54_ = _dafny.Set({D6_DC16(d_2043_v52_), d_2044_v53_})
            d_2046_v55_: _dafny.Map
            d_2046_v55_ = _dafny.Map({not(not(d_2037_v46_.f11)): len(d_2045_v54_)})
            d_2047_v56_: D0
            d_2047_v56_ = D0_DC1(len((d_2034_v43_).set(default__.safeIndex(len(d_1984_v0_), len(d_2034_v43_)), (self).f12)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), (self).f12, d_2037_v46_.f11)
            d_2048_v57_: _dafny.Array
            nw311_ = _dafny.Array(None, 21)
            nw311_[int(0)] = d_2037_v46_.f11
            nw311_[int(1)] = d_2037_v46_.f11
            nw311_[int(2)] = p1
            nw311_[int(3)] = ((self).f13) == ((self).f13)
            nw311_[int(4)] = self.f11
            nw311_[int(5)] = True
            nw311_[int(6)] = d_2037_v46_.f11
            nw311_[int(7)] = p1
            nw311_[int(8)] = (self).fm7(globalState)
            nw311_[int(9)] = (d_2046_v55_) != (_dafny.Map({d_2037_v46_.f11: self.f22}))
            nw311_[int(10)] = d_2037_v46_.f11
            nw311_[int(11)] = False
            nw311_[int(12)] = not(default__.fm1(self.f22, self.f11, 906, globalState))
            nw311_[int(13)] = d_2037_v46_.f11
            nw311_[int(14)] = d_2037_v46_.f11
            nw311_[int(15)] = self.f11
            nw311_[int(16)] = d_2037_v46_.f11
            def iife180_(_pat_let48_0):
                def iife181_(d_2049_dt__update__tmp_h1_):
                    def iife182_(_pat_let49_0):
                        def iife183_(d_2050_dt__update_hcf2_h0_):
                            return D0_DC1((d_2049_dt__update__tmp_h1_).cf1, d_2050_dt__update_hcf2_h0_, (d_2049_dt__update__tmp_h1_).cf3, (d_2049_dt__update__tmp_h1_).cf4)
                        return iife183_(_pat_let49_0)
                    return iife182_((self).f20)
                return iife181_(_pat_let48_0)
            nw311_[int(17)] = not((iife180_(d_2047_v56_)).cf4)
            nw311_[int(18)] = (d_2036_v45_) == (d_2036_v45_)
            nw311_[int(19)] = not(not(((d_2041_v50_)[p0] if (p0) in (d_2041_v50_) else d_2037_v46_.f11)))
            nw311_[int(20)] = p1
            d_2048_v57_ = nw311_
            index292_ = default__.safeIndex(983, (d_2048_v57_).length(0))
            (d_2048_v57_)[index292_] = not(self.f11)
            index293_ = default__.safeIndex(983, (d_2048_v57_).length(0))
            (d_2048_v57_)[index293_] = False
        elif True:
            d_2051_v58_: _dafny.Seq
            d_2051_v58_ = _dafny.SeqWithoutIsStrInference([(self).f20])
            d_2051_v58_ = d_2051_v58_
            d_2052_v59_: D7
            d_2052_v59_ = D7_DC21(_dafny.MultiSet([self.f11, True]))
            d_2053_v60_: _dafny.MultiSet
            d_2053_v60_ = _dafny.MultiSet([d_2037_v46_.f11, d_2037_v46_.f11, p1, self.f11])
            pat_let_tv71_ = d_2053_v60_
            d_2054_v61_: _dafny.Map
            def iife184_(_pat_let50_0):
                def iife185_(d_2055_dt__update__tmp_h2_):
                    def iife186_(_pat_let51_0):
                        def iife187_(d_2056_dt__update_hcf39_h0_):
                            return D7_DC21(d_2056_dt__update_hcf39_h0_)
                        return iife187_(_pat_let51_0)
                    return iife186_(pat_let_tv71_)
                return iife185_(_pat_let50_0)
            d_2054_v61_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2037_v46_.f11]): iife184_(d_2052_v59_)})
            d_2057_v62_: str
            d_2057_v62_ = _dafny.CodePoint('v')
            d_2058_v63_: _dafny.Map
            d_2058_v63_ = _dafny.Map({_dafny.MultiSet([_dafny.CodePoint('i'), d_2057_v62_, d_2057_v62_]): (self).f13})
            d_2059_v64_: _dafny.MultiSet
            d_2059_v64_ = _dafny.MultiSet([d_2057_v62_])
            d_2060_v65_: _dafny.Array
            nw312_ = _dafny.Array(None, 2)
            nw312_[int(0)] = d_2058_v63_
            nw312_[int(1)] = (_dafny.Map({d_2059_v64_: (p2)[default__.safeIndex(821, (p2).length(0))]})) | (d_2058_v63_)
            d_2060_v65_ = nw312_
            index294_ = default__.safeIndex(75, (d_2060_v65_).length(0))
            (d_2060_v65_)[index294_] = d_2058_v63_
            d_2061_v66_: D15
            d_2061_v66_ = D15_DC40(default__.fm55(831, (self).f13, p0, globalState))
            d_2062_v67_: _dafny.Map
            d_2062_v67_ = _dafny.Map({(self).f13: (self).f20})
            d_2063_v68_: _dafny.Map
            d_2063_v68_ = _dafny.Map({(self).f13: 398})
            d_2064_v69_: _dafny.Map
            d_2064_v69_ = _dafny.Map({(self.f22) + (len(d_2062_v67_)): ((d_2063_v68_)[(self).f12] if ((self).f12) in (d_2063_v68_) else (self).f12)})
            d_2065_v70_: _dafny.Map
            d_2065_v70_ = _dafny.Map({d_2037_v46_.f11: d_2037_v46_.f11})
            index295_ = default__.safeIndex(75, (d_2060_v65_).length(0))
            rhs293_ = ((D15_DC40(_dafny.Map({d_1984_v0_: d_2052_v59_})) if d_2037_v46_.f11 else d_2061_v66_)).cf64
            rhs294_ = d_2058_v63_
            rhs295_ = d_2057_v62_
            rhs296_ = ((d_2064_v69_)[len(d_2065_v70_)] if (len(d_2065_v70_)) in (d_2064_v69_) else len((self).f20))
            rhs297_ = (d_2037_v46_.f11) not in ((d_2065_v70_ if d_2037_v46_.f11 else _dafny.Map({self.f11: d_2037_v46_.f11})))
            lhs176_ = d_2060_v65_
            lhs177_ = default__.safeIndex(75, (d_2060_v65_).length(0))
            lhs178_ = self
            lhs179_ = d_2037_v46_
            d_2054_v61_ = rhs293_
            lhs176_[lhs177_] = rhs294_
            d_2057_v62_ = rhs295_
            lhs178_.f22 = rhs296_
            lhs179_.f11 = rhs297_
            d_2066_v71_: _dafny.MultiSet
            d_2066_v71_ = _dafny.MultiSet([default__.safeDivisionInt(p0, (self).f13), (self).f13])
            d_2066_v71_ = ((_dafny.MultiSet([(p2)[default__.safeIndex(821, (p2).length(0))]])) - (d_2066_v71_)) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([525, (self).f13, (p2)[default__.safeIndex(821, (p2).length(0))], (0) - ((self).f12)])))
            d_2067_v72_: _dafny.Array
            nw313_ = _dafny.Array(_dafny.Map({}), 9)
            d_2067_v72_ = nw313_
            d_2068_v73_: _dafny.Seq
            d_2068_v73_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f22])])
            index296_ = default__.safeIndex(574, (d_2067_v72_).length(0))
            (d_2067_v72_)[index296_] = _dafny.Map({(self).f12: (d_2068_v73_)[default__.safeIndex((p2)[default__.safeIndex(821, (p2).length(0))], len(d_2068_v73_))]})
            d_2069_v74_: _dafny.Map
            d_2069_v74_ = _dafny.Map({(self).f13: (d_2034_v43_) + (_dafny.SeqWithoutIsStrInference([p0]))})
            index297_ = default__.safeIndex(574, (d_2067_v72_).length(0))
            rhs298_ = len(_dafny.Map({(self).f13: not (d_2037_v46_.f11) or (d_2037_v46_.f11)}))
            rhs299_ = d_2069_v74_
            rhs300_ = p1
            rhs301_ = (p0) * (202)
            rhs302_ = d_2057_v62_
            lhs180_ = self
            lhs181_ = d_2067_v72_
            lhs182_ = default__.safeIndex(574, (d_2067_v72_).length(0))
            lhs183_ = d_2037_v46_
            lhs184_ = self
            lhs180_.f22 = rhs298_
            lhs181_[lhs182_] = rhs299_
            lhs183_.f11 = rhs300_
            lhs184_.f22 = rhs301_
            d_2057_v62_ = rhs302_
            d_2070_v75_: _dafny.Seq
            d_2070_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdeejae"))
            d_2070_v75_ = (self).f20
        d_2071_v76_: _dafny.Array
        def lambda108_(d_2072_v0_):
            def lambda109_(d_2073_i3_):
                return _dafny.SeqWithoutIsStrInference([(d_2072_v0_)[default__.safeIndex((_dafny.MultiSet([_dafny.CodePoint('y')])).cardinality, len(d_2072_v0_))], not((D15_DC41(self.f11)).cf65)])

            return lambda109_

        init60_ = lambda108_(d_1984_v0_)
        nw314_ = _dafny.Array(None, 25)
        for i0_60_ in range(nw314_.length(0)):
            nw314_[i0_60_] = init60_(i0_60_)
        d_2071_v76_ = nw314_
        d_2071_v76_ = d_2071_v76_


class C10(T0):
    def  __init__(self):
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    def ctor__(self, f11):
        (self).f11 = f11

    def fm5(self, p0, globalState):
        return D0_DC1(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2074_i0_ in range(default__.abs(361))])), -38])), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecqjml"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fojk"))), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omya"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfg")))), True)

    def fm6(self, p0, p1, p2, globalState):
        return self.f11

    def fm16(self, p0, p1, globalState):
        return (not (not(self.f11)) or ((D0_DC3((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f11])) for d_2075_i0_ in range(default__.abs(325))])))).cardinality, self.f11)).cf8)) not in (_dafny.Map({self.f11: (0) - (len(_dafny.Map({self.f11: self.f11})))}))

    def m9(self, p0, p1, globalState):
        d_2076_v0_: _dafny.Array
        def lambda110_(d_2077_i0_):
            return self.f11

        init61_ = lambda110_
        nw315_ = _dafny.Array(None, 25)
        for i0_61_ in range(nw315_.length(0)):
            nw315_[i0_61_] = init61_(i0_61_)
        d_2076_v0_ = nw315_
        d_2078_v1_: _dafny.Set
        d_2078_v1_ = _dafny.Set({p0, p0})
        index298_ = default__.safeIndex(0, (d_2076_v0_).length(0))
        (d_2076_v0_)[index298_] = (d_2078_v1_).isdisjoint(d_2078_v1_)
        index299_ = default__.safeIndex(0, (d_2076_v0_).length(0))
        (d_2076_v0_)[index299_] = self.f11
        d_2079_i1_: int
        d_2079_i1_ = 0
        with _dafny.label("8"):
            while (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]:
                with _dafny.c_label("8"):
                    if (d_2079_i1_) >= (100):
                        raise _dafny.Break("8")
                    d_2079_i1_ = (d_2079_i1_) + (1)
                    index300_ = default__.safeIndex(0, (d_2076_v0_).length(0))
                    (d_2076_v0_)[index300_] = (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]
                    d_2080_v2_: _dafny.MultiSet
                    d_2080_v2_ = _dafny.MultiSet([232])
                    (globalState).f8 = (self).fm16(p0, ((d_2080_v2_).cardinality) + (p0), globalState)
                    d_2081_v3_: _dafny.Seq
                    d_2081_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                    d_2082_v4_: _dafny.Seq
                    d_2082_v4_ = (d_2081_v3_).set(default__.safeIndex(p0, len(d_2081_v3_)), p0)
                    d_2083_v5_: _dafny.Map
                    d_2083_v5_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2084_i2_ in range(default__.abs(-638))]): (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]})
                    d_2085_v6_: _dafny.Map
                    d_2085_v6_ = _dafny.Map({default__.fm0(globalState): len(d_2083_v5_)})
                    d_2086_v7_: _dafny.Seq
                    d_2086_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptqylp"))
                    d_2085_v6_ = (d_2085_v6_).set(_dafny.SeqWithoutIsStrInference([p0 for d_2087_i3_ in range(default__.abs(142))]), ((0) - (len(d_2086_v7_))) * (p0))
                    d_2088_v8_: _dafny.Map
                    d_2088_v8_ = _dafny.Map({d_2078_v1_: (d_2080_v2_).cardinality})
                    d_2088_v8_ = (d_2088_v8_).set(d_2078_v1_, p0)
                    pass
            pass
        d_2089_v9_: _dafny.Map
        d_2089_v9_ = _dafny.Map({p0: (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]})
        d_2090_v10_: _dafny.MultiSet
        d_2090_v10_ = _dafny.MultiSet([default__.fm2(globalState), -777])
        d_2091_v11_: _dafny.Seq
        d_2091_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(0) - ((0) - (((d_2090_v10_)[p0] if (p0) in (d_2090_v10_) else p0))): (self).fm6((d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], p0, p0, globalState)}), d_2089_v9_, (_dafny.Map({len(d_2078_v1_): self.f11})).set(p0, False)])
        d_2089_v9_ = (d_2091_v11_)[default__.safeIndex(default__.safeModuloInt(p0, -618), len(d_2091_v11_))]
        d_2092_v12_: _dafny.Seq
        d_2092_v12_ = _dafny.SeqWithoutIsStrInference([self.f11])
        d_2093_v13_: _dafny.Map
        d_2093_v13_ = _dafny.Map({(self).fm6((d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], len(d_2092_v12_), p0, globalState): len(d_2078_v1_)})
        if (d_2092_v12_)[default__.safeIndex(len((_dafny.Map({not((d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]): p0})) | (d_2093_v13_)), len(d_2092_v12_))]:
            d_2094_v14_: int
            d_2094_v14_ = 724
            d_2095_v15_: str
            d_2095_v15_ = _dafny.CodePoint('n')
            d_2096_v16_: _dafny.Map
            d_2096_v16_ = _dafny.Map({d_2095_v15_: self.f11})
            rhs303_ = p0
            rhs304_ = not(((d_2094_v14_) > (p0) if not (self.f11) or (((d_2089_v9_)[d_2094_v14_] if (d_2094_v14_) in (d_2089_v9_) else True)) else not (self.f11) or (((d_2096_v16_)[d_2095_v15_] if (d_2095_v15_) in (d_2096_v16_) else False))))
            rhs305_ = d_2094_v14_
            rhs306_ = 182
            rhs307_ = (d_2094_v14_) >= (p0)
            lhs185_ = globalState
            lhs186_ = self
            d_2094_v14_ = rhs303_
            lhs185_.f8 = rhs304_
            d_2094_v14_ = rhs305_
            d_2094_v14_ = rhs306_
            lhs186_.f11 = rhs307_
            d_2094_v14_ = d_2094_v14_
            if self.f11:
                d_2094_v14_ = d_2094_v14_
                d_2097_v17_: _dafny.Map
                d_2097_v17_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p0, p0])})
                d_2098_v18_: _dafny.Seq
                d_2098_v18_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2099_v19_: _dafny.Seq
                d_2099_v19_ = ((d_2097_v17_)[p0] if (p0) in (d_2097_v17_) else d_2098_v18_)
                d_2100_v20_: _dafny.Array
                nw316_ = _dafny.Array(None, 27)
                nw316_[int(0)] = d_2099_v19_
                nw316_[int(1)] = _dafny.SeqWithoutIsStrInference([(0) - (p0) for d_2101_i4_ in range(default__.abs(944))])
                nw316_[int(2)] = d_2099_v19_
                nw316_[int(3)] = d_2099_v19_
                nw316_[int(4)] = d_2099_v19_
                nw316_[int(5)] = d_2099_v19_
                nw316_[int(6)] = d_2099_v19_
                nw316_[int(7)] = d_2099_v19_
                nw316_[int(8)] = d_2099_v19_
                nw316_[int(9)] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2099_v19_: D0_DC3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))), (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))])})) for d_2102_i5_ in range(default__.abs(804))])
                nw316_[int(10)] = d_2099_v19_
                nw316_[int(11)] = d_2099_v19_
                nw316_[int(12)] = d_2099_v19_
                nw316_[int(13)] = d_2099_v19_
                nw316_[int(14)] = d_2099_v19_
                nw316_[int(15)] = d_2099_v19_
                nw316_[int(16)] = d_2099_v19_
                nw316_[int(17)] = _dafny.SeqWithoutIsStrInference([p0, d_2094_v14_, d_2094_v14_])
                nw316_[int(18)] = d_2099_v19_
                nw316_[int(19)] = d_2099_v19_
                nw316_[int(20)] = d_2098_v18_
                nw316_[int(21)] = d_2098_v18_
                nw316_[int(22)] = d_2099_v19_
                nw316_[int(23)] = d_2099_v19_
                nw316_[int(24)] = d_2099_v19_
                nw316_[int(25)] = d_2098_v18_
                nw316_[int(26)] = d_2099_v19_
                d_2100_v20_ = nw316_
                d_2103_v21_: _dafny.Map
                d_2104_v22_: int
                d_2105_v23_: _dafny.Map
                out32_: _dafny.Map
                out33_: int
                out34_: _dafny.Map
                out32_, out33_, out34_ = (self).m10(self.f11, d_2100_v20_, p0, globalState)
                d_2103_v21_ = out32_
                d_2104_v22_ = out33_
                d_2105_v23_ = out34_
                d_2094_v14_ = (0) - (default__.fm2(globalState))
                (self).f11 = (False) not in (d_2092_v12_)
                (globalState).f8 = not(((self.f11) or (self.f11)) or (self.f11))
            elif True:
                d_2106_v24_: _dafny.MultiSet
                d_2106_v24_ = _dafny.MultiSet([self.f11, self.f11, not(not((d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]))])
                d_2107_v25_: _dafny.Seq
                d_2107_v25_ = _dafny.SeqWithoutIsStrInference([d_2106_v24_])
                rhs308_ = (0) - (d_2094_v14_)
                rhs309_ = (d_2107_v25_) + (d_2107_v25_)
                d_2094_v14_ = rhs308_
                d_2107_v25_ = rhs309_
                d_2108_v26_: _dafny.Map
                d_2108_v26_ = _dafny.Map({(d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]: d_2095_v15_})
                d_2109_v27_: _dafny.Seq
                d_2109_v27_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_2108_v26_))])
                d_2110_v28_: C8
                nw317_ = C8()
                nw317_.ctor__(d_2109_v27_)
                d_2110_v28_ = nw317_
                d_2111_v29_: _dafny.Seq
                d_2111_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvdlfmlk"))
                d_2112_v30_: _dafny.Seq
                d_2112_v30_ = _dafny.SeqWithoutIsStrInference([d_2111_v29_, (d_2111_v29_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sslfvi")))])
                d_2112_v30_ = _dafny.SeqWithoutIsStrInference([d_2111_v29_])
                index301_ = default__.safeIndex(59, (p1).length(0))
                (p1)[index301_] = default__.fm2(globalState)
                index302_ = default__.safeIndex(59, (p1).length(0))
                (p1)[index302_] = default__.safeDivisionInt((len(d_2111_v29_)) * (p0), (0) - (d_2094_v14_))
                (globalState).f5 = (d_2110_v28_).f23
            d_2094_v14_ = p0
            d_2113_v31_: C6
            nw318_ = C6()
            nw318_.ctor__()
            d_2113_v31_ = nw318_
        elif True:
            d_2114_v32_: _dafny.Seq
            d_2114_v32_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2115_v33_: _dafny.Map
            d_2115_v33_ = _dafny.Map({(d_2114_v32_)[default__.safeIndex(p0, len(d_2114_v32_))]: d_2076_v0_})
            d_2115_v33_ = (d_2115_v33_).set((d_2114_v32_)[default__.safeIndex(918, len(d_2114_v32_))], d_2076_v0_)
            d_2116_v34_: int
            d_2116_v34_ = 725
            d_2117_v35_: _dafny.Map
            d_2117_v35_ = _dafny.Map({d_2116_v34_: p0})
            d_2118_v36_: D11
            d_2118_v36_ = D11_DC31((d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))])
            d_2119_v37_: _dafny.Seq
            d_2119_v37_ = _dafny.SeqWithoutIsStrInference([d_2118_v36_])
            d_2116_v34_ = (0) - (default__.safeModuloInt((((d_2117_v35_)[len(d_2119_v37_)] if (len(d_2119_v37_)) in (d_2117_v35_) else (0) - (p0))) * (d_2116_v34_), len(d_2092_v12_)))
            d_2120_v38_: str
            d_2120_v38_ = _dafny.CodePoint('r')
            d_2120_v38_ = d_2120_v38_
            d_2116_v34_ = (0) - (d_2116_v34_)
            d_2121_v39_: _dafny.Map
            d_2121_v39_ = _dafny.Map({self.f11: self.f11})
            (globalState).f8 = ((len(d_2121_v39_)) - (551)) <= ((-17) * (d_2116_v34_))
        if ((d_2092_v12_) < (d_2092_v12_) if ((d_2089_v9_)[len(d_2093_v13_)] if (len(d_2093_v13_)) in (d_2089_v9_) else (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]) else self.f11):
            d_2122_v40_: C6
            nw319_ = C6()
            nw319_.ctor__()
            d_2122_v40_ = nw319_
            d_2123_v41_: _dafny.Seq
            d_2123_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgnjjve"))
            d_2124_v42_: _dafny.Map
            d_2124_v42_ = _dafny.Map({(default__.fm2(globalState)) * (len(d_2123_v41_)): default__.safeDivisionInt(p0, p0)})
            d_2124_v42_ = (d_2124_v42_).set((p0) + ((0) - (p0)), (p0) - (p0))
            d_2125_v43_: _dafny.Set
            d_2125_v43_ = _dafny.Set({d_2124_v42_, d_2124_v42_, d_2124_v42_})
            d_2125_v43_ = (_dafny.Set({d_2124_v42_, d_2124_v42_, d_2124_v42_})) | (d_2125_v43_)
            d_2126_v44_: C4
            nw320_ = C4()
            nw320_.ctor__(default__.safeDivisionInt(457, p0), default__.fm1((0) - (p0), self.f11, p0, globalState))
            d_2126_v44_ = nw320_
            d_2127_v45_: _dafny.Seq
            d_2127_v45_ = _dafny.SeqWithoutIsStrInference([(d_2126_v44_).f25])
            d_2128_v46_: _dafny.Seq
            d_2128_v46_ = d_2127_v45_
            d_2129_v47_: _dafny.Map
            d_2129_v47_ = _dafny.Map({(d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]: d_2128_v46_})
            d_2129_v47_ = (d_2129_v47_).set(self.f11, d_2128_v46_)
        elif True:
            index303_ = default__.safeIndex(0, (d_2076_v0_).length(0))
            (d_2076_v0_)[index303_] = not(self.f11)
            d_2130_v48_: _dafny.Map
            d_2130_v48_ = _dafny.Map({d_2089_v9_: d_2092_v12_})
            d_2131_v50_: _dafny.Map
            def iife188_():
                coll84_ = _dafny.Map()
                compr_84_: int
                for compr_84_ in _dafny.IntegerRange(-830, 911):
                    d_2132_v49_: int = compr_84_
                    if ((-830) <= (d_2132_v49_)) and ((d_2132_v49_) < (911)):
                        coll84_[default__.safeDivisionInt(d_2132_v49_, p0)] = (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]
                return _dafny.Map(coll84_)
            d_2131_v50_ = _dafny.Map({d_2076_v0_: (d_2130_v48_).set(iife188_()
            , _dafny.SeqWithoutIsStrInference([(d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]]))})
            d_2133_v53_: _dafny.Map
            d_2133_v53_ = _dafny.Map({d_2089_v9_: p0})
            d_2134_v55_: _dafny.Seq
            def iife189_():
                coll85_ = _dafny.Map()
                compr_85_: _dafny.Map
                for compr_85_ in (default__.fm56(p0, (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], self.f11, (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], globalState)).Elements:
                    d_2135_v54_: _dafny.Map = compr_85_
                    if (d_2135_v54_) in (default__.fm56(p0, (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], self.f11, (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))], globalState)):
                        coll85_[d_2135_v54_] = d_2092_v12_
                return _dafny.Map(coll85_)
            d_2134_v55_ = _dafny.SeqWithoutIsStrInference([d_2130_v48_, d_2130_v48_, d_2130_v48_, iife189_()
            ])
            d_2136_v56_: _dafny.Array
            nw321_ = _dafny.Array(None, 25)
            nw321_[int(0)] = d_2130_v48_
            nw321_[int(1)] = (d_2130_v48_) | (((d_2131_v50_)[d_2076_v0_] if (d_2076_v0_) in (d_2131_v50_) else _dafny.Map({d_2089_v9_: _dafny.SeqWithoutIsStrInference([self.f11])})))
            def iife190_():
                coll86_ = _dafny.Map()
                compr_86_: int
                for compr_86_ in (d_2090_v10_).Elements:
                    d_2137_v51_: int = compr_86_
                    if (d_2137_v51_) in (d_2090_v10_):
                        coll86_[(d_2137_v51_) - (p0)] = self.f11
                return _dafny.Map(coll86_)
            nw321_[int(2)] = _dafny.Map({iife190_()
            : _dafny.SeqWithoutIsStrInference([(d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]])})
            nw321_[int(3)] = d_2130_v48_
            nw321_[int(4)] = d_2130_v48_
            nw321_[int(5)] = d_2130_v48_
            nw321_[int(6)] = _dafny.Map({d_2089_v9_: d_2092_v12_})
            def iife191_():
                coll87_ = _dafny.Map()
                compr_87_: _dafny.Map
                for compr_87_ in (d_2133_v53_).keys.Elements:
                    d_2138_v52_: _dafny.Map = compr_87_
                    if (d_2138_v52_) in (d_2133_v53_):
                        coll87_[d_2138_v52_] = d_2092_v12_
                return _dafny.Map(coll87_)
            nw321_[int(7)] = (iife191_()
            ) | (d_2130_v48_)
            nw321_[int(8)] = d_2130_v48_
            nw321_[int(9)] = d_2130_v48_
            nw321_[int(10)] = d_2130_v48_
            nw321_[int(11)] = _dafny.Map({d_2089_v9_: d_2092_v12_})
            nw321_[int(12)] = ((_dafny.Map({d_2089_v9_: d_2092_v12_}))) | (d_2130_v48_)
            nw321_[int(13)] = d_2130_v48_
            nw321_[int(14)] = (d_2134_v55_)[default__.safeIndex(p0, len(d_2134_v55_))]
            nw321_[int(15)] = d_2130_v48_
            nw321_[int(16)] = d_2130_v48_
            nw321_[int(17)] = _dafny.Map({_dafny.Map({default__.fm2(globalState): (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))]}): d_2092_v12_})
            nw321_[int(18)] = d_2130_v48_
            nw321_[int(19)] = d_2130_v48_
            nw321_[int(20)] = (d_2130_v48_ if (d_2076_v0_)[default__.safeIndex(0, (d_2076_v0_).length(0))] else d_2130_v48_)
            nw321_[int(21)] = d_2130_v48_
            nw321_[int(22)] = d_2130_v48_
            nw321_[int(23)] = d_2130_v48_
            nw321_[int(24)] = d_2130_v48_
            d_2136_v56_ = nw321_
            index304_ = default__.safeIndex(93, (d_2136_v56_).length(0))
            (d_2136_v56_)[index304_] = default__.fm57(-671, self.f11, globalState)
            d_2139_v57_: int
            d_2139_v57_ = -177
            d_2140_v58_: _dafny.Seq
            d_2140_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehm"))
            index305_ = default__.safeIndex(93, (d_2136_v56_).length(0))
            rhs310_ = d_2130_v48_
            rhs311_ = (d_2140_v58_) < (d_2140_v58_)
            rhs312_ = (p0) + (d_2139_v57_)
            rhs313_ = self.f11
            lhs187_ = d_2136_v56_
            lhs188_ = default__.safeIndex(93, (d_2136_v56_).length(0))
            lhs189_ = self
            lhs190_ = globalState
            lhs187_[lhs188_] = rhs310_
            lhs189_.f11 = rhs311_
            d_2139_v57_ = rhs312_
            lhs190_.f8 = rhs313_
            (globalState).f8 = True
            d_2139_v57_ = p0
            index306_ = default__.safeIndex(138, (p1).length(0))
            (p1)[index306_] = p0
            index307_ = default__.safeIndex(138, (p1).length(0))
            (p1)[index307_] = (0) - (d_2139_v57_)
        index308_ = default__.safeIndex(0, (d_2076_v0_).length(0))
        (d_2076_v0_)[index308_] = self.f11

    def m10(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        d_2141_i0_: int
        d_2141_i0_ = 0
        with _dafny.label("9"):
            while self.f11:
                with _dafny.c_label("9"):
                    if (d_2141_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_2141_i0_ = (d_2141_i0_) + (1)
                    d_2142_v0_: _dafny.Array
                    nw322_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                    d_2142_v0_ = nw322_
                    d_2143_v1_: _dafny.Seq
                    d_2143_v1_ = _dafny.SeqWithoutIsStrInference([self.f11])
                    rhs314_ = ((p2) != (p2)) or (p0)
                    rhs315_ = d_2142_v0_
                    rhs316_ = (d_2143_v1_) + (d_2143_v1_)
                    rhs317_ = self.f11
                    rhs318_ = self.f11
                    lhs191_ = self
                    lhs192_ = globalState
                    lhs193_ = globalState
                    lhs191_.f11 = rhs314_
                    d_2142_v0_ = rhs315_
                    d_2143_v1_ = rhs316_
                    lhs192_.f8 = rhs317_
                    lhs193_.f8 = rhs318_
                    d_2144_v2_: _dafny.Map
                    d_2144_v2_ = _dafny.Map({p0: p0})
                    d_2144_v2_ = (d_2144_v2_).set(True, p0)
                    d_2145_v3_: _dafny.Seq
                    d_2145_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                    (globalState).f8 = (self).fm6((default__.fm1(len(d_2145_v3_), p0, p2, globalState)) == (p0), 678, p2, globalState)
                    d_2146_v4_: _dafny.Map
                    d_2146_v4_ = _dafny.Map({len(d_2143_v1_): p2})
                    d_2147_v5_: _dafny.Map
                    d_2147_v5_ = _dafny.Map({len(d_2146_v4_): d_2146_v4_})
                    d_2148_v6_: D3
                    d_2148_v6_ = D3_DC9(p2, d_2147_v5_, (0) - (len(_dafny.SeqWithoutIsStrInference([p2 for d_2149_i1_ in range(default__.abs(478))]))), (0) - (len(d_2145_v3_)))
                    d_2150_v7_: _dafny.Map
                    d_2150_v7_ = _dafny.Map({d_2148_v6_: d_2145_v3_})
                    pat_let_tv72_ = p2
                    pat_let_tv73_ = p2
                    def iife192_(_pat_let52_0):
                        def iife193_(d_2151_dt__update__tmp_h0_):
                            def iife194_(_pat_let53_0):
                                def iife195_(d_2152_dt__update_hcf17_h0_):
                                    def iife196_(_pat_let54_0):
                                        def iife197_(d_2153_dt__update_hcf18_h0_):
                                            return D3_DC9((d_2151_dt__update__tmp_h0_).cf15, (d_2151_dt__update__tmp_h0_).cf16, d_2152_dt__update_hcf17_h0_, d_2153_dt__update_hcf18_h0_)
                                        return iife197_(_pat_let54_0)
                                    return iife196_(pat_let_tv73_)
                                return iife195_(_pat_let53_0)
                            return iife194_(pat_let_tv72_)
                        return iife193_(_pat_let52_0)
                    d_2150_v7_ = (d_2150_v7_).set(iife192_(d_2148_v6_), default__.fm36(p2, len(_dafny.Map({self.f11: self.f11})), _dafny.CodePoint('r'), globalState))
                    pass
            pass
        d_2154_v8_: _dafny.Array
        nw323_ = _dafny.Array(int(0), 26)
        d_2154_v8_ = nw323_
        d_2155_v9_: _dafny.Array
        nw324_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
        d_2155_v9_ = nw324_
        d_2156_v10_: _dafny.Map
        d_2156_v10_ = _dafny.Map({d_2154_v8_: (d_2155_v9_ if p0 else d_2155_v9_)})
        d_2156_v10_ = (d_2156_v10_).set(d_2154_v8_, d_2155_v9_)
        d_2157_v11_: _dafny.Array
        d_2158_v12_: bool
        out35_: _dafny.Array
        out36_: bool
        out35_, out36_ = default__.m0(globalState)
        d_2157_v11_ = out35_
        d_2158_v12_ = out36_
        d_2159_v13_: str
        d_2159_v13_ = _dafny.CodePoint('x')
        d_2160_v14_: _dafny.Array
        nw325_ = _dafny.Array(False, 7)
        d_2160_v14_ = nw325_
        d_2161_v15_: _dafny.Array
        d_2161_v15_ = d_2160_v14_
        d_2162_v16_: _dafny.Seq
        d_2162_v16_ = _dafny.SeqWithoutIsStrInference([d_2160_v14_, (d_2161_v15_)])
        rhs319_ = d_2159_v13_
        rhs320_ = (d_2162_v16_) == (d_2162_v16_)
        rhs321_ = (not(self.f11)) or (self.f11)
        lhs194_ = globalState
        lhs195_ = globalState
        d_2159_v13_ = rhs319_
        lhs194_.f8 = rhs320_
        lhs195_.f8 = rhs321_
        d_2163_v17_: D6
        d_2163_v17_ = D6_DC17(p0, p2)
        d_2164_v18_: D6
        d_2164_v18_ = D6_DC20(d_2163_v17_)
        d_2165_v19_: _dafny.Map
        d_2165_v19_ = _dafny.Map({d_2164_v18_: 515})
        r1 = ((d_2165_v19_)[d_2164_v18_] if (d_2164_v18_) in (d_2165_v19_) else p2)
        d_2166_v20_: _dafny.Seq
        d_2166_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwra"))
        d_2167_v21_: T0
        nw326_ = C2()
        nw326_.ctor__(p2, d_2166_v20_, p2, default__.safeDivisionInt(p2, p2), p0)
        d_2167_v21_ = nw326_
        d_2167_v21_ = d_2167_v21_
        d_2168_v22_: _dafny.Seq
        d_2168_v22_ = _dafny.SeqWithoutIsStrInference([d_2158_v12_, not(d_2167_v21_.f11), p0])
        d_2169_v23_: _dafny.MultiSet
        d_2169_v23_ = _dafny.MultiSet([default__.fm39(d_2158_v12_, p0, globalState), d_2168_v22_, d_2168_v22_, d_2168_v22_, d_2168_v22_])
        d_2170_v24_: _dafny.Map
        d_2170_v24_ = _dafny.Map({((d_2169_v23_)[d_2168_v22_] if (d_2168_v22_) in (d_2169_v23_) else p2): d_2166_v20_})
        r0 = d_2170_v24_
        r1 = p2
        d_2171_v25_: _dafny.Map
        d_2171_v25_ = _dafny.Map({d_2160_v14_: default__.fm1(p2, p0, (0) - (p2), globalState)})
        r2 = ((_dafny.Map({d_2160_v14_: self.f11})) | (d_2171_v25_)) | ((d_2171_v25_) | (d_2171_v25_))
        return r0, r1, r2


class C11(T1):
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def fm7(self, globalState):
        return not((not(False)) == ((not(not((D0_DC3((self).f13, not(False))).cf8))) == (False)))

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqlbnbxw"))

    def fm14(self, globalState):
        source31_ = D1_DC5(_dafny.MultiSet([(self).f13, (self).f12]), not(True))
        if source31_.is_DC5:
            d_2172___mcc_h0_ = source31_.cf10
            d_2173___mcc_h1_ = source31_.cf11
            d_2174_cf11_ = d_2173___mcc_h1_
            d_2175_cf10_ = d_2172___mcc_h0_
            return (self).f12
        elif source31_.is_DC4:
            d_2176___mcc_h2_ = source31_.cf9
            d_2177_cf9_ = d_2176___mcc_h2_
            return (0) - ((0) - ((0) - (default__.safeModuloInt((self).f13, (self).f13))))
        elif True:
            d_2178___mcc_h3_ = source31_.cf12
            d_2179_cf12_ = d_2178___mcc_h3_
            return (self).f13

    def m8(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        d_2180_v0_: bool
        d_2180_v0_ = False
        d_2181_v1_: _dafny.Map
        d_2181_v1_ = _dafny.Map({(self).f13: d_2180_v0_})
        d_2182_v2_: _dafny.Seq
        d_2182_v2_ = _dafny.SeqWithoutIsStrInference([d_2180_v0_])
        d_2183_v3_: _dafny.Map
        d_2183_v3_ = _dafny.Map({d_2180_v0_: d_2180_v0_})
        d_2184_v4_: _dafny.Array
        nw327_ = _dafny.Array(int(0), 21)
        d_2184_v4_ = nw327_
        d_2185_v5_: _dafny.Map
        d_2185_v5_ = _dafny.Map({d_2180_v0_: (self).f12})
        d_2186_v6_: _dafny.Map
        d_2186_v6_ = _dafny.Map({((d_2181_v1_)[806] if (806) in (d_2181_v1_) else d_2180_v0_): d_2185_v5_})
        d_2187_v7_: _dafny.Map
        d_2187_v7_ = _dafny.Map({d_2184_v4_: len(((d_2186_v6_)[d_2180_v0_] if (d_2180_v0_) in (d_2186_v6_) else _dafny.Map({d_2180_v0_: (self).f13})))})
        d_2188_v8_: str
        d_2188_v8_ = _dafny.CodePoint('a')
        d_2189_v9_: _dafny.MultiSet
        d_2189_v9_ = _dafny.MultiSet([_dafny.CodePoint('i'), d_2188_v8_])
        d_2190_v10_: _dafny.Array
        nw328_ = _dafny.Array(None, 27)
        nw328_[int(0)] = (self).f12
        nw328_[int(1)] = (self).f12
        nw328_[int(2)] = (self).f12
        nw328_[int(3)] = (self).f13
        nw328_[int(4)] = len(d_2181_v1_)
        nw328_[int(5)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2191_i0_ in range(default__.abs(825))]))
        nw328_[int(6)] = (self).f13
        nw328_[int(7)] = (0) - ((self).f13)
        nw328_[int(8)] = (_dafny.MultiSet([(self).f12])).cardinality
        nw328_[int(9)] = (self).f13
        nw328_[int(10)] = ((self).f13 if d_2180_v0_ else (default__.fm15(d_2180_v0_, (self).f13, d_2182_v2_, d_2180_v0_, globalState)).cf1)
        nw328_[int(11)] = len((d_2183_v3_) | (d_2183_v3_))
        nw328_[int(12)] = ((self).f13) * (len(d_2187_v7_))
        nw328_[int(13)] = ((0) - ((0) - ((self).f13)) if True else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fj"))))
        nw328_[int(14)] = len((_dafny.Map({d_2180_v0_: (self).f12}) if d_2180_v0_ else _dafny.Map({d_2180_v0_: (self).f12})))
        nw328_[int(15)] = ((self).f13) * ((self).f13)
        nw328_[int(16)] = (self).f12
        nw328_[int(17)] = 182
        nw328_[int(18)] = (self).f12
        nw328_[int(19)] = (self).f12
        nw328_[int(20)] = (self).f12
        nw328_[int(21)] = ((self).f13) * ((self).f13)
        nw328_[int(22)] = (self).f12
        nw328_[int(23)] = -69
        nw328_[int(24)] = (d_2189_v9_).cardinality
        nw328_[int(25)] = 230
        nw328_[int(26)] = ((self).f12) + ((self).f12)
        d_2190_v10_ = nw328_
        index309_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        (d_2190_v10_)[index309_] = default__.safeModuloInt((self).f12, (self).f12)
        index310_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        (d_2190_v10_)[index310_] = (self).f12
        d_2192_v11_: _dafny.Seq
        d_2192_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scvtajpa"))
        index311_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        (d_2190_v10_)[index311_] = len(d_2192_v11_)
        d_2193_v12_: _dafny.Array
        def lambda111_(d_2194_v8_):
            def lambda112_(d_2195_i1_):
                return _dafny.SeqWithoutIsStrInference([d_2194_v8_ for d_2196_i2_ in range(default__.abs(799))])

            return lambda112_

        init62_ = lambda111_(d_2188_v8_)
        nw329_ = _dafny.Array(None, 22)
        for i0_62_ in range(nw329_.length(0)):
            nw329_[i0_62_] = init62_(i0_62_)
        d_2193_v12_ = nw329_
        index312_ = default__.safeIndex(145, (d_2193_v12_).length(0))
        (d_2193_v12_)[index312_] = d_2192_v11_
        index313_ = default__.safeIndex(271, (d_2193_v12_).length(0))
        (d_2193_v12_)[index313_] = d_2192_v11_
        d_2197_v13_: _dafny.MultiSet
        d_2197_v13_ = _dafny.MultiSet([d_2180_v0_])
        d_2198_v14_: _dafny.MultiSet
        d_2198_v14_ = _dafny.MultiSet([(d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))], (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))]])
        pat_let_tv74_ = d_2190_v10_
        pat_let_tv75_ = d_2190_v10_
        pat_let_tv76_ = d_2180_v0_
        pat_let_tv77_ = d_2180_v0_
        pat_let_tv78_ = d_2180_v0_
        index314_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        index315_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        index316_ = default__.safeIndex(145, (d_2193_v12_).length(0))
        index317_ = default__.safeIndex(271, (d_2193_v12_).length(0))
        def lambda113_(source32_):
            if source32_.is_DC1:
                d_2199___mcc_h0_ = source32_.cf1
                d_2200___mcc_h1_ = source32_.cf2
                d_2201___mcc_h2_ = source32_.cf3
                d_2202___mcc_h3_ = source32_.cf4
                d_2203_cf4_ = d_2202___mcc_h3_
                d_2204_cf3_ = d_2201___mcc_h2_
                d_2205_cf2_ = d_2200___mcc_h1_
                d_2206_cf1_ = d_2199___mcc_h0_
                return (d_2204_cf3_) - (-932)
            elif source32_.is_DC2:
                d_2207___mcc_h4_ = source32_.cf5
                d_2208___mcc_h5_ = source32_.cf6
                d_2209_cf6_ = d_2208___mcc_h5_
                d_2210_cf5_ = d_2207___mcc_h4_
                return (pat_let_tv75_)[default__.safeIndex(746, (pat_let_tv74_).length(0))]
            elif source32_.is_DC3:
                d_2211___mcc_h6_ = source32_.cf7
                d_2212___mcc_h7_ = source32_.cf8
                d_2213_cf8_ = d_2212___mcc_h7_
                d_2214_cf7_ = d_2211___mcc_h6_
                return (self).f12
            elif True:
                d_2215___mcc_h8_ = source32_.cf0
                d_2216_cf0_ = d_2215___mcc_h8_
                return -906

        def lambda114_(source33_):
            if source33_.is_DC5:
                d_2217___mcc_h9_ = source33_.cf10
                d_2218___mcc_h10_ = source33_.cf11
                d_2219_cf11_ = d_2218___mcc_h10_
                d_2220_cf10_ = d_2217___mcc_h9_
                return pat_let_tv76_
            elif source33_.is_DC4:
                d_2221___mcc_h11_ = source33_.cf9
                d_2222_cf9_ = d_2221___mcc_h11_
                return pat_let_tv77_
            elif True:
                d_2223___mcc_h12_ = source33_.cf12
                d_2224_cf12_ = d_2223___mcc_h12_
                return pat_let_tv78_

        rhs322_ = (self).f12
        rhs323_ = lambda113_(D0_DC1((0) - (((d_2197_v13_).set(d_2180_v0_, default__.abs((self).f12))).cardinality), d_2192_v11_, (self).f13, d_2180_v0_))
        rhs324_ = d_2192_v11_
        rhs325_ = lambda114_(D1_DC5(d_2198_v14_, d_2180_v0_))
        rhs326_ = d_2192_v11_
        lhs196_ = d_2190_v10_
        lhs197_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        lhs198_ = d_2190_v10_
        lhs199_ = default__.safeIndex(746, (d_2190_v10_).length(0))
        lhs200_ = d_2193_v12_
        lhs201_ = default__.safeIndex(145, (d_2193_v12_).length(0))
        lhs202_ = d_2193_v12_
        lhs203_ = default__.safeIndex(271, (d_2193_v12_).length(0))
        lhs196_[lhs197_] = rhs322_
        lhs198_[lhs199_] = rhs323_
        lhs200_[lhs201_] = rhs324_
        d_2180_v0_ = rhs325_
        lhs202_[lhs203_] = rhs326_
        d_2225_i3_: int
        d_2225_i3_ = 0
        with _dafny.label("10"):
            while (self).fm7(globalState):
                with _dafny.c_label("10"):
                    if (d_2225_i3_) >= (100):
                        raise _dafny.Break("10")
                    d_2225_i3_ = (d_2225_i3_) + (1)
                    d_2226_v15_: _dafny.Map
                    d_2226_v15_ = _dafny.Map({d_2188_v8_: default__.fm1((d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))], d_2180_v0_, (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))], globalState)})
                    d_2227_v16_: _dafny.Map
                    d_2227_v16_ = _dafny.Map({(d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))]: (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))]})
                    d_2228_v17_: _dafny.Set
                    d_2228_v17_ = _dafny.Set({len(d_2227_v16_), (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))]})
                    d_2229_v18_: _dafny.Set
                    d_2229_v18_ = _dafny.Set({len((self).fm8((self).f12, d_2226_v15_, d_2228_v17_, globalState))})
                    d_2230_v19_: _dafny.Seq
                    d_2230_v19_ = _dafny.SeqWithoutIsStrInference([d_2198_v14_])
                    d_2231_v20_: _dafny.Array
                    nw330_ = _dafny.Array(None, 10)
                    nw330_[int(0)] = d_2198_v14_
                    nw330_[int(1)] = d_2198_v14_
                    nw330_[int(2)] = _dafny.MultiSet([len(d_2229_v18_), (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))], (self).f12, len(d_2182_v2_)])
                    nw330_[int(3)] = d_2198_v14_
                    nw330_[int(4)] = (d_2230_v19_)[default__.safeIndex((self).f12, len(d_2230_v19_))]
                    nw330_[int(5)] = d_2198_v14_
                    nw330_[int(6)] = d_2198_v14_
                    nw330_[int(7)] = (d_2198_v14_).set((self).f12, default__.abs((self).f13))
                    nw330_[int(8)] = d_2198_v14_
                    nw330_[int(9)] = (d_2230_v19_)[default__.safeIndex((d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))], len(d_2230_v19_))]
                    d_2231_v20_ = nw330_
                    d_2232_v21_: _dafny.Map
                    d_2232_v21_ = _dafny.Map({d_2180_v0_: d_2231_v20_})
                    d_2232_v21_ = (d_2232_v21_).set(d_2180_v0_, d_2231_v20_)
                    d_2192_v11_ = (((d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))]) + ((d_2192_v11_).set(default__.safeIndex(589, len(d_2192_v11_)), d_2188_v8_))) + (d_2192_v11_)
                    d_2233_v22_: C3
                    nw331_ = C3()
                    nw331_.ctor__((self).f13, (self).f13)
                    d_2233_v22_ = nw331_
                    d_2189_v9_ = (d_2189_v9_).set(d_2188_v8_, default__.abs(len((_dafny.Map({(self).f12: (self).f13})) | (d_2227_v16_))))
                    pass
            pass
        hi10_ = (self).f12
        for d_2234_i4_ in range(((self).f12 if d_2180_v0_ else (self).f13), hi10_):
            d_2235_v23_: _dafny.Map
            d_2235_v23_ = _dafny.Map({271: d_2182_v2_})
            d_2236_v24_: _dafny.Array
            nw332_ = _dafny.Array(None, 14)
            nw332_[int(0)] = d_2182_v2_
            nw332_[int(1)] = d_2182_v2_
            nw332_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2180_v0_])
            nw332_[int(3)] = d_2182_v2_
            nw332_[int(4)] = _dafny.SeqWithoutIsStrInference([(self).fm7(globalState), d_2180_v0_])
            nw332_[int(5)] = (d_2182_v2_) + (d_2182_v2_)
            nw332_[int(6)] = d_2182_v2_
            nw332_[int(7)] = d_2182_v2_
            nw332_[int(8)] = d_2182_v2_
            nw332_[int(9)] = ((d_2182_v2_).set(default__.safeIndex(d_2234_i4_, len(d_2182_v2_)), default__.fm1(144, d_2180_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwqyqwpb"))), globalState))) + (d_2182_v2_)
            nw332_[int(10)] = d_2182_v2_
            nw332_[int(11)] = (((((d_2235_v23_)[d_2234_i4_] if (d_2234_i4_) in (d_2235_v23_) else _dafny.SeqWithoutIsStrInference([d_2180_v0_, False]))).set(default__.safeIndex((0) - ((self).f12), len(((d_2235_v23_)[d_2234_i4_] if (d_2234_i4_) in (d_2235_v23_) else _dafny.SeqWithoutIsStrInference([d_2180_v0_, False])))), d_2180_v0_)) + (d_2182_v2_)).set(default__.safeIndex((self).f12, len(((((d_2235_v23_)[d_2234_i4_] if (d_2234_i4_) in (d_2235_v23_) else _dafny.SeqWithoutIsStrInference([d_2180_v0_, False]))).set(default__.safeIndex((0) - ((self).f12), len(((d_2235_v23_)[d_2234_i4_] if (d_2234_i4_) in (d_2235_v23_) else _dafny.SeqWithoutIsStrInference([d_2180_v0_, False])))), d_2180_v0_)) + (d_2182_v2_))), d_2180_v0_)
            nw332_[int(12)] = d_2182_v2_
            nw332_[int(13)] = d_2182_v2_
            d_2236_v24_ = nw332_
            d_2237_v25_: _dafny.Map
            d_2237_v25_ = _dafny.Map({len(((d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))]).set(default__.safeIndex((self).f12, len((d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))])), d_2188_v8_)): len(_dafny.SeqWithoutIsStrInference([d_2234_i4_, (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))], (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))]]))})
            d_2238_v26_: _dafny.Seq
            d_2238_v26_ = _dafny.SeqWithoutIsStrInference([d_2184_v4_, d_2190_v10_, d_2184_v4_])
            index318_ = default__.safeIndex(746, (d_2190_v10_).length(0))
            index319_ = default__.safeIndex(746, (d_2190_v10_).length(0))
            rhs327_ = d_2236_v24_
            rhs328_ = (len(d_2237_v25_) if d_2180_v0_ else (self).f12)
            rhs329_ = len((d_2238_v26_) + ((_dafny.SeqWithoutIsStrInference([d_2190_v10_])) + (d_2238_v26_)))
            lhs204_ = d_2190_v10_
            lhs205_ = default__.safeIndex(746, (d_2190_v10_).length(0))
            lhs206_ = d_2190_v10_
            lhs207_ = default__.safeIndex(746, (d_2190_v10_).length(0))
            d_2236_v24_ = rhs327_
            lhs204_[lhs205_] = rhs328_
            lhs206_[lhs207_] = rhs329_
            d_2239_v27_: C2
            nw333_ = C2()
            nw333_.ctor__((self).f13, (d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))], (d_2234_i4_) + (len(_dafny.SeqWithoutIsStrInference([d_2180_v0_, d_2180_v0_]))), ((self).f12) - ((self).f13), d_2180_v0_)
            d_2239_v27_ = nw333_
            def iife198_():
                def iife200_():
                    coll90_ = _dafny.Map()
                    compr_90_: _dafny.Seq
                    for compr_90_ in (_dafny.SeqWithoutIsStrInference([(d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arocdqm")), (d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))]])).Elements:
                        d_2240_v29_: _dafny.Seq = compr_90_
                        if (d_2240_v29_) in (_dafny.SeqWithoutIsStrInference([(d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arocdqm")), (d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))]])):
                            coll90_[d_2240_v29_] = len(d_2237_v25_)
                    return _dafny.Map(coll90_)
                coll88_ = _dafny.Map()
                def iife199_():
                    coll89_ = _dafny.Map()
                    compr_89_: _dafny.Seq
                    for compr_89_ in (_dafny.SeqWithoutIsStrInference([(d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arocdqm")), (d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))]])).Elements:
                        d_2240_v29_: _dafny.Seq = compr_89_
                        if (d_2240_v29_) in (_dafny.SeqWithoutIsStrInference([(d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arocdqm")), (d_2193_v12_)[default__.safeIndex(145, (d_2193_v12_).length(0))]])):
                            coll89_[d_2240_v29_] = len(d_2237_v25_)
                    return _dafny.Map(coll89_)
                compr_88_: _dafny.Seq
                for compr_88_ in (iife199_()
                ).keys.Elements:
                    d_2241_v28_: _dafny.Seq = compr_88_
                    if (d_2241_v28_) in (iife200_()
                    ):
                        coll88_[d_2241_v28_] = _dafny.Set({(d_2239_v27_).f26})
                return _dafny.Map(coll88_)
            d_2180_v0_ = (d_2192_v11_) not in (iife198_()
            )
            d_2242_v30_: D9
            d_2242_v30_ = D9_DC26(d_2188_v8_)
            d_2243_v31_: _dafny.Array
            nw334_ = _dafny.Array(None, 7)
            nw334_[int(0)] = d_2188_v8_
            nw334_[int(1)] = d_2188_v8_
            nw334_[int(2)] = d_2188_v8_
            nw334_[int(3)] = d_2188_v8_
            nw334_[int(4)] = (d_2242_v30_).cf48
            nw334_[int(5)] = d_2188_v8_
            nw334_[int(6)] = _dafny.CodePoint('q')
            d_2243_v31_ = nw334_
            d_2244_v32_: _dafny.MultiSet
            d_2244_v32_ = _dafny.MultiSet([d_2184_v4_, d_2190_v10_])
            index320_ = default__.safeIndex(746, (d_2190_v10_).length(0))
            rhs330_ = (d_2190_v10_)[default__.safeIndex(746, (d_2190_v10_).length(0))]
            rhs331_ = (d_2243_v31_ if (d_2244_v32_).issubset(d_2244_v32_) else d_2243_v31_)
            lhs208_ = d_2190_v10_
            lhs209_ = default__.safeIndex(746, (d_2190_v10_).length(0))
            lhs208_[lhs209_] = rhs330_
            d_2243_v31_ = rhs331_
        r1 = d_2180_v0_
        d_2245_v33_: D6
        d_2245_v33_ = D6_DC19(d_2180_v0_, d_2180_v0_, d_2180_v0_, True, d_2182_v2_)
        pat_let_tv79_ = d_2180_v0_
        pat_let_tv80_ = d_2185_v5_
        pat_let_tv81_ = d_2190_v10_
        pat_let_tv82_ = d_2190_v10_
        pat_let_tv83_ = d_2190_v10_
        pat_let_tv84_ = d_2190_v10_
        pat_let_tv85_ = d_2190_v10_
        pat_let_tv86_ = d_2190_v10_
        pat_let_tv87_ = d_2192_v11_
        pat_let_tv88_ = d_2185_v5_
        pat_let_tv89_ = d_2190_v10_
        pat_let_tv90_ = d_2190_v10_
        def lambda115_(source34_):
            if source34_.is_DC17:
                d_2246___mcc_h13_ = source34_.cf30
                d_2247___mcc_h14_ = source34_.cf31
                d_2248_cf31_ = d_2247___mcc_h14_
                d_2249_cf30_ = d_2246___mcc_h13_
                if pat_let_tv79_:
                    return _dafny.SeqWithoutIsStrInference([(self).f12])
                elif True:
                    return _dafny.SeqWithoutIsStrInference([(self).f12, len(pat_let_tv80_), (pat_let_tv82_)[default__.safeIndex(746, (pat_let_tv81_).length(0))], (0) - (d_2248_cf31_), d_2248_cf31_])
            elif source34_.is_DC18:
                d_2250___mcc_h15_ = source34_.cf32
                d_2251_cf32_ = d_2250___mcc_h15_
                if not((D4_DC12(d_2251_cf32_, (self).f13, (pat_let_tv84_)[default__.safeIndex(746, (pat_let_tv83_).length(0))])).cf21):
                    return _dafny.SeqWithoutIsStrInference([(self).f13 for d_2252_i5_ in range(default__.abs(-196))])
                elif True:
                    return _dafny.SeqWithoutIsStrInference([-548, (self).f12])
            elif source34_.is_DC19:
                d_2253___mcc_h16_ = source34_.cf33
                d_2254___mcc_h17_ = source34_.cf34
                d_2255___mcc_h18_ = source34_.cf35
                d_2256___mcc_h19_ = source34_.cf36
                d_2257___mcc_h20_ = source34_.cf37
                d_2258_cf37_ = d_2257___mcc_h20_
                d_2259_cf36_ = d_2256___mcc_h19_
                d_2260_cf35_ = d_2255___mcc_h18_
                d_2261_cf34_ = d_2254___mcc_h17_
                d_2262_cf33_ = d_2253___mcc_h16_
                return _dafny.SeqWithoutIsStrInference([(pat_let_tv86_)[default__.safeIndex(746, (pat_let_tv85_).length(0))], len(pat_let_tv87_)])
            elif source34_.is_DC16:
                d_2263___mcc_h21_ = source34_.cf29
                d_2264_cf29_ = d_2263___mcc_h21_
                return _dafny.SeqWithoutIsStrInference([len(pat_let_tv88_), (self).f12])
            elif True:
                d_2265___mcc_h22_ = source34_.cf38
                d_2266_cf38_ = d_2265___mcc_h22_
                return (_dafny.SeqWithoutIsStrInference([(pat_let_tv90_)[default__.safeIndex(746, (pat_let_tv89_).length(0))] for d_2267_i6_ in range(default__.abs(826))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f13])), (self).f13]))

        r0 = lambda115_(d_2245_v33_)
        r1 = d_2180_v0_
        return r0, r1


class C12(T0, T1):
    def  __init__(self):
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: int = int(0)
        self.f18: _dafny.MultiSet = _dafny.MultiSet({})
        self._f19: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f18, f19, f11, f12, f13):
        (self).f18 = f18
        (self)._f19 = f19
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13

    def fm5(self, p0, globalState):
        source35_ = D0_DC3(len(_dafny.SeqWithoutIsStrInference([(self).f19, (self).f19, (self).f13, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gw"))), (self).f13])), self.f11)
        if source35_.is_DC1:
            d_2268___mcc_h0_ = source35_.cf1
            d_2269___mcc_h1_ = source35_.cf2
            d_2270___mcc_h2_ = source35_.cf3
            d_2271___mcc_h3_ = source35_.cf4
            d_2272_cf4_ = d_2271___mcc_h3_
            d_2273_cf3_ = d_2270___mcc_h2_
            d_2274_cf2_ = d_2269___mcc_h1_
            d_2275_cf1_ = d_2268___mcc_h0_
            def iife201_():
                coll91_ = _dafny.Map()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(772, 95):
                    d_2276_v0_: int = compr_91_
                    if ((772) <= (d_2276_v0_)) and ((d_2276_v0_) < (95)):
                        coll91_[(d_2276_v0_) + (d_2273_cf3_)] = _dafny.SeqWithoutIsStrInference([self.f11])
                return _dafny.Map(coll91_)
            return D0_DC1(len(iife201_()
), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyqxy"))).set(default__.safeIndex((self).f19, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyqxy")))), _dafny.CodePoint('x')), (self).f12, False)
        elif source35_.is_DC2:
            d_2277___mcc_h4_ = source35_.cf5
            d_2278___mcc_h5_ = source35_.cf6
            d_2279_cf6_ = d_2278___mcc_h5_
            d_2280_cf5_ = d_2277___mcc_h4_
            return D0_DC1(494, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccyu")), 785, self.f11)
        elif source35_.is_DC3:
            d_2281___mcc_h6_ = source35_.cf7
            d_2282___mcc_h7_ = source35_.cf8
            d_2283_cf8_ = d_2282___mcc_h7_
            d_2284_cf7_ = d_2281___mcc_h6_
            return D0_DC1(len(_dafny.SeqWithoutIsStrInference([self.f11, self.f11])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apji")), (self).f12, self.f11)
        elif True:
            d_2285___mcc_h8_ = source35_.cf0
            d_2286_cf0_ = d_2285___mcc_h8_
            if self.f11:
                return D0_DC1((self).f12, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "niwbwv")), (self).f13, True)
            elif True:
                return D0_DC1((self).f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kidqydpry")), (self).f13, self.f11)

    def fm6(self, p0, p1, p2, globalState):
        return self.f11

    def fm7(self, globalState):
        return self.f11

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))

    def fm13(self, p0, globalState):
        return self.f11

    def m7(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2287_v0_: _dafny.Seq
        d_2287_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2288_v1_: C4
        nw335_ = C4()
        nw335_.ctor__(len((d_2287_v0_) + (d_2287_v0_)), True)
        d_2288_v1_ = nw335_
        index321_ = default__.safeIndex(893, (p0).length(0))
        (p0)[index321_] = (d_2288_v1_).f25
        index322_ = default__.safeIndex(893, (p0).length(0))
        (p0)[index322_] = (0) - ((0) - (default__.safeDivisionInt(default__.fm2(globalState), (self).f13)))
        if self.f11:
            (globalState).f8 = not(default__.fm1(p1, True, ((self).f12) + ((p0)[default__.safeIndex(893, (p0).length(0))]), globalState))
            d_2289_v2_: _dafny.Map
            d_2289_v2_ = _dafny.Map({self.f11: self.f11})
            d_2290_v3_: _dafny.Set
            d_2290_v3_ = _dafny.Set({self.f11, False, self.f11})
            d_2291_v4_: _dafny.Map
            d_2291_v4_ = _dafny.Map({17: d_2290_v3_})
            (globalState).f8 = ((self.f11 if self.f11 else self.f11) if (not((self).fm6(((d_2289_v2_)[True] if (True) in (d_2289_v2_) else default__.fm1((p0)[default__.safeIndex(893, (p0).length(0))], self.f11, (p0)[default__.safeIndex(893, (p0).length(0))], globalState)), (0) - ((d_2288_v1_).f25), len(d_2291_v4_), globalState))) or (True) else not(((0) - ((p0)[default__.safeIndex(893, (p0).length(0))])) < ((self).f12)))
            d_2292_v5_: _dafny.Array
            def lambda116_(d_2293_i0_):
                return (d_2293_i0_) * (139)

            init63_ = lambda116_
            nw336_ = _dafny.Array(None, 10)
            for i0_63_ in range(nw336_.length(0)):
                nw336_[i0_63_] = init63_(i0_63_)
            d_2292_v5_ = nw336_
            d_2292_v5_ = (d_2292_v5_ if self.f11 else p0)
            d_2294_v6_: C5
            nw337_ = C5()
            nw337_.ctor__((self).f13)
            d_2294_v6_ = nw337_
            if (self.f11 if not (not(self.f11)) or (self.f11) else True):
                d_2295_v7_: _dafny.MultiSet
                d_2295_v7_ = _dafny.MultiSet([self.f11, self.f11, self.f11, self.f11, self.f11])
                d_2296_v8_: _dafny.Map
                d_2296_v8_ = _dafny.Map({self.f11: d_2295_v7_})
                d_2297_v9_: _dafny.Map
                d_2297_v9_ = _dafny.Map({(len(d_2296_v8_)) - ((d_2288_v1_).f25): (self.f11 if self.f11 else self.f11)})
                d_2297_v9_ = (d_2297_v9_).set((0) - ((0) - (((d_2288_v1_).f25) - ((d_2288_v1_).f25))), self.f11)
                d_2298_v10_: C1
                nw338_ = C1()
                nw338_.ctor__((0) - (default__.safeModuloInt((d_2288_v1_).f25, (self).f12)))
                d_2298_v10_ = nw338_
                (globalState).f8 = not(self.f11)
                (globalState).f8 = not(self.f11)
                d_2299_v11_: _dafny.Map
                d_2299_v11_ = _dafny.Map({(d_2288_v1_).f25: d_2292_v5_})
                d_2299_v11_ = (d_2299_v11_) | (_dafny.Map({(d_2298_v10_).fm34(globalState): d_2292_v5_}))
            elif True:
                r0 = ((self).f19) + (((d_2287_v0_)[default__.safeIndex(481, len(d_2287_v0_))] if default__.fm1((d_2288_v1_).f25, self.f11, -204, globalState) else (self).f19))
                index323_ = default__.safeIndex(893, (p0).length(0))
                (p0)[index323_] = (13) + ((self).f19)
                d_2300_v12_: _dafny.Seq
                d_2300_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvld"))
                d_2301_v13_: str
                d_2301_v13_ = _dafny.CodePoint('j')
                d_2302_v14_: T2
                nw339_ = C2()
                nw339_.ctor__(default__.fm2(globalState), ((d_2300_v12_) + (d_2300_v12_)).set(default__.safeIndex((d_2288_v1_).f25, len((d_2300_v12_) + (d_2300_v12_))), d_2301_v13_), 422, ((0) - ((self).f12)) - (len(d_2300_v12_)), not(self.f11))
                d_2302_v14_ = nw339_
                d_2302_v14_ = d_2302_v14_
                d_2303_v15_: _dafny.Map
                d_2303_v15_ = _dafny.Map({self.f11: (0) - ((p0)[default__.safeIndex(893, (p0).length(0))])})
                d_2304_v16_: _dafny.Map
                d_2304_v16_ = _dafny.Map({(p0)[default__.safeIndex(893, (p0).length(0))]: default__.safeModuloInt(((d_2303_v15_)[(d_2294_v6_).fm27((d_2288_v1_).f25, _dafny.SeqWithoutIsStrInference([d_2301_v13_ for d_2305_i1_ in range(default__.abs(855))]), self.f11, globalState)] if ((d_2294_v6_).fm27((d_2288_v1_).f25, _dafny.SeqWithoutIsStrInference([d_2301_v13_ for d_2306_i1_ in range(default__.abs(855))]), self.f11, globalState)) in (d_2303_v15_) else (d_2288_v1_).f25), len(d_2303_v15_))})
                d_2307_v17_: _dafny.MultiSet
                d_2307_v17_ = _dafny.MultiSet([self.f11, self.f11, not(self.f11)])
                d_2308_v18_: _dafny.Seq
                d_2308_v18_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11])
                d_2304_v16_ = (d_2304_v16_).set(((d_2307_v17_).cardinality) - ((d_2288_v1_).f25), len(d_2308_v18_))
                (globalState).f8 = not (self.f11) or (self.f11)
        elif True:
            (globalState).f8 = self.f11
            d_2309_v19_: _dafny.MultiSet
            d_2309_v19_ = _dafny.MultiSet([(self).f12])
            d_2310_v20_: str
            d_2310_v20_ = _dafny.CodePoint('e')
            d_2311_v21_: _dafny.Seq
            d_2311_v21_ = _dafny.SeqWithoutIsStrInference([d_2310_v20_, d_2310_v20_, d_2310_v20_, d_2310_v20_, d_2310_v20_])
            d_2312_v22_: D1
            d_2312_v22_ = D1_DC5(((d_2309_v19_).set((d_2288_v1_).f25, default__.abs(len(d_2311_v21_)))).intersection(d_2309_v19_), self.f11)
            d_2312_v22_ = d_2312_v22_
            d_2313_v23_: _dafny.Array
            nw340_ = _dafny.Array(None, 3)
            nw340_[int(0)] = self.f11
            nw340_[int(1)] = False
            nw340_[int(2)] = (d_2309_v19_).isdisjoint(d_2309_v19_)
            d_2313_v23_ = nw340_
            d_2314_v24_: D9
            d_2314_v24_ = D9_DC26(d_2310_v20_)
            d_2315_v25_: _dafny.Map
            d_2315_v25_ = _dafny.Map({self.f11: True})
            d_2316_v26_: D0
            d_2316_v26_ = D0_DC2(len(d_2315_v25_), p0)
            d_2317_v27_: _dafny.Seq
            d_2317_v27_ = _dafny.SeqWithoutIsStrInference([d_2311_v21_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oakqppwdg")), default__.fm36(len(d_2311_v21_), (self).f19, d_2310_v20_, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwsviefmw"))])
            d_2318_v28_: _dafny.MultiSet
            d_2318_v28_ = _dafny.MultiSet([self.f11, self.f11])
            d_2319_v29_: _dafny.Seq
            d_2319_v29_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11, self.f11])
            d_2320_v30_: _dafny.Map
            d_2320_v30_ = _dafny.Map({(self).f19: self.f11})
            d_2321_v32_: _dafny.Map
            def iife202_():
                coll92_ = _dafny.Set()
                compr_92_: int
                for compr_92_ in (d_2320_v30_).keys.Elements:
                    d_2322_v31_: int = compr_92_
                    if (d_2322_v31_) in (d_2320_v30_):
                        coll92_ = coll92_.union(_dafny.Set([(d_2322_v31_) * (len(_dafny.SeqWithoutIsStrInference([False])))]))
                return _dafny.Set(coll92_)
            d_2321_v32_ = _dafny.Map({d_2310_v20_: (d_2319_v29_)[default__.safeIndex(len(iife202_()
            ), len(d_2319_v29_))]})
            d_2323_v33_: _dafny.Set
            d_2323_v33_ = _dafny.Set({(p0)[default__.safeIndex(893, (p0).length(0))]})
            d_2324_v34_: _dafny.Array
            nw341_ = _dafny.Array(None, 28)
            nw341_[int(0)] = ((d_2311_v21_) + (d_2311_v21_)).set(default__.safeIndex((0) - (default__.fm2(globalState)), len((d_2311_v21_) + (d_2311_v21_))), (d_2314_v24_).cf48)
            nw341_[int(1)] = d_2311_v21_
            nw341_[int(2)] = d_2311_v21_
            nw341_[int(3)] = ((d_2311_v21_).set(default__.safeIndex((d_2316_v26_).cf5, len(d_2311_v21_)), d_2310_v20_) if self.f11 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llmohotq")))
            nw341_[int(4)] = (d_2317_v27_)[default__.safeIndex((0) - ((d_2288_v1_).f25), len(d_2317_v27_))]
            nw341_[int(5)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2325_i2_ in range(default__.abs(21))])).set(default__.safeIndex(default__.fm2(globalState), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2326_i2_ in range(default__.abs(21))]))), _dafny.CodePoint('o'))
            nw341_[int(6)] = d_2311_v21_
            nw341_[int(7)] = (((d_2311_v21_).set(default__.safeIndex(((d_2318_v28_)[self.f11] if (self.f11) in (d_2318_v28_) else (self).f12), len(d_2311_v21_)), d_2310_v20_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vx")))).set(default__.safeIndex((self).f12, len(((d_2311_v21_).set(default__.safeIndex(((d_2318_v28_)[self.f11] if (self.f11) in (d_2318_v28_) else (self).f12), len(d_2311_v21_)), d_2310_v20_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vx"))))), _dafny.CodePoint('q'))
            nw341_[int(8)] = _dafny.SeqWithoutIsStrInference([d_2310_v20_ for d_2327_i3_ in range(default__.abs(619))])
            nw341_[int(9)] = _dafny.SeqWithoutIsStrInference([d_2310_v20_ for d_2328_i4_ in range(default__.abs(-35))])
            nw341_[int(10)] = d_2311_v21_
            nw341_[int(11)] = d_2311_v21_
            nw341_[int(12)] = d_2311_v21_
            nw341_[int(13)] = d_2311_v21_
            nw341_[int(14)] = d_2311_v21_
            nw341_[int(15)] = d_2311_v21_
            nw341_[int(16)] = d_2311_v21_
            nw341_[int(17)] = d_2311_v21_
            nw341_[int(18)] = d_2311_v21_
            nw341_[int(19)] = d_2311_v21_
            nw341_[int(20)] = d_2311_v21_
            nw341_[int(21)] = d_2311_v21_
            nw341_[int(22)] = (d_2311_v21_) + (_dafny.SeqWithoutIsStrInference([d_2310_v20_ for d_2329_i5_ in range(default__.abs(153))]))
            nw341_[int(23)] = d_2311_v21_
            nw341_[int(24)] = ((self).fm8((d_2288_v1_).f25, d_2321_v32_, d_2323_v33_, globalState)).set(default__.safeIndex((0) - ((d_2288_v1_).f25), len((self).fm8((d_2288_v1_).f25, d_2321_v32_, d_2323_v33_, globalState))), d_2310_v20_)
            nw341_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tc"))
            nw341_[int(26)] = (d_2311_v21_ if self.f11 else d_2311_v21_)
            nw341_[int(27)] = d_2311_v21_
            d_2324_v34_ = nw341_
            index324_ = default__.safeIndex(772, (d_2324_v34_).length(0))
            (d_2324_v34_)[index324_] = d_2311_v21_
            d_2330_v35_: _dafny.Map
            d_2330_v35_ = _dafny.Map({(d_2288_v1_).f25: (d_2288_v1_).f25})
            d_2331_v36_: _dafny.Map
            d_2331_v36_ = _dafny.Map({375: d_2330_v35_})
            d_2332_v37_: D3
            d_2332_v37_ = D3_DC9((self).f19, d_2331_v36_, p1, (self).f13)
            index325_ = default__.safeIndex(772, (d_2324_v34_).length(0))
            rhs332_ = d_2313_v23_
            rhs333_ = d_2311_v21_
            rhs334_ = d_2310_v20_
            rhs335_ = ((d_2311_v21_) + (d_2311_v21_)) == (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxs"))).set(default__.safeIndex((d_2288_v1_).f25, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxs")))), _dafny.CodePoint('t')) if True else d_2311_v21_))
            rhs336_ = ((len(_dafny.Set({d_2332_v37_}))) + ((d_2288_v1_).f25)) - (p1)
            lhs210_ = d_2324_v34_
            lhs211_ = default__.safeIndex(772, (d_2324_v34_).length(0))
            lhs212_ = self
            d_2313_v23_ = rhs332_
            lhs210_[lhs211_] = rhs333_
            d_2310_v20_ = rhs334_
            lhs212_.f11 = rhs335_
            r0 = rhs336_
            index326_ = default__.safeIndex(893, (p0).length(0))
            (p0)[index326_] = (0) - ((self).f19)
            index327_ = default__.safeIndex(893, (p0).length(0))
            (p0)[index327_] = (default__.safeModuloInt((d_2287_v0_)[default__.safeIndex((self).f13, len(d_2287_v0_))], (self).f13)) * ((self).f19)
        d_2333_v38_: _dafny.Array
        nw342_ = _dafny.Array(False, 12)
        d_2333_v38_ = nw342_
        index328_ = default__.safeIndex(886, (d_2333_v38_).length(0))
        (d_2333_v38_)[index328_] = self.f11
        d_2334_v39_: _dafny.Array
        def lambda117_(d_2335_i6_):
            return default__.safeDivisionInt(d_2335_i6_, (self).f19)

        init64_ = lambda117_
        nw343_ = _dafny.Array(None, 25)
        for i0_64_ in range(nw343_.length(0)):
            nw343_[i0_64_] = init64_(i0_64_)
        d_2334_v39_ = nw343_
        index329_ = default__.safeIndex(886, (d_2333_v38_).length(0))
        rhs337_ = self.f11
        rhs338_ = d_2334_v39_
        lhs213_ = d_2333_v38_
        lhs214_ = default__.safeIndex(886, (d_2333_v38_).length(0))
        lhs213_[lhs214_] = rhs337_
        d_2334_v39_ = rhs338_
        d_2336_v40_: D12
        d_2336_v40_ = D12_DC35(default__.safeDivisionInt(p1, default__.fm2(globalState)))
        source36_ = d_2336_v40_
        if source36_.is_DC33:
            d_2337___mcc_h0_ = source36_.cf58
            d_2338_cf58_ = d_2337___mcc_h0_
            d_2334_v39_ = d_2334_v39_
            d_2339_v41_: C10
            nw344_ = C10()
            nw344_.ctor__(self.f11)
            d_2339_v41_ = nw344_
            d_2339_v41_ = (d_2339_v41_ if (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))] else d_2339_v41_)
            d_2340_v42_: C0
            nw345_ = C0()
            nw345_.ctor__((p0)[default__.safeIndex(893, (p0).length(0))], self.f11, (d_2288_v1_).f25, p1)
            d_2340_v42_ = nw345_
            d_2341_v43_: _dafny.Map
            d_2341_v43_ = _dafny.Map({self.f11: d_2340_v42_})
            d_2342_v44_: _dafny.Seq
            d_2342_v44_ = _dafny.SeqWithoutIsStrInference([d_2340_v42_.f29])
            d_2343_v45_: _dafny.MultiSet
            d_2343_v45_ = _dafny.MultiSet([len(_dafny.Map({d_2341_v43_: (d_2342_v44_)[default__.safeIndex((d_2288_v1_).f25, len(d_2342_v44_))]}))])
            d_2344_v46_: _dafny.Seq
            d_2344_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdwgotbd"))
            d_2345_v47_: D0
            d_2345_v47_ = D0_DC1(p1, d_2344_v46_, (d_2288_v1_).f25, d_2340_v42_.f29)
            d_2346_v48_: _dafny.Map
            d_2346_v48_ = _dafny.Map({len(_dafny.Set({(d_2345_v47_).cf2, d_2344_v46_})): (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]})
            index330_ = default__.safeIndex(886, (d_2333_v38_).length(0))
            (d_2333_v38_)[index330_] = not (d_2340_v42_.f29) or (((d_2343_v45_).cardinality) >= (len(d_2346_v48_)))
            d_2344_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjd"))
        elif source36_.is_DC34:
            d_2347___mcc_h1_ = source36_.cf59
            d_2348_cf59_ = d_2347___mcc_h1_
            d_2349_v49_: _dafny.MultiSet
            d_2349_v49_ = _dafny.MultiSet([((p0)[default__.safeIndex(893, (p0).length(0))]) < (len(d_2287_v0_)), False])
            index331_ = default__.safeIndex(893, (p0).length(0))
            index332_ = default__.safeIndex(893, (p0).length(0))
            rhs339_ = ((d_2349_v49_)[self.f11] if (self.f11) in (d_2349_v49_) else p1)
            rhs340_ = p1
            lhs215_ = p0
            lhs216_ = default__.safeIndex(893, (p0).length(0))
            lhs217_ = p0
            lhs218_ = default__.safeIndex(893, (p0).length(0))
            lhs215_[lhs216_] = rhs339_
            lhs217_[lhs218_] = rhs340_
            d_2350_v50_: _dafny.Map
            d_2350_v50_ = _dafny.Map({(self).f13: d_2333_v38_})
            d_2351_v51_: _dafny.Array
            d_2351_v51_ = d_2333_v38_
            d_2352_v52_: _dafny.Array
            nw346_ = _dafny.Array(None, 18)
            nw346_[int(0)] = d_2333_v38_
            nw346_[int(1)] = d_2333_v38_
            nw346_[int(2)] = ((d_2350_v50_)[264] if (264) in (d_2350_v50_) else d_2333_v38_)
            nw346_[int(3)] = d_2333_v38_
            nw346_[int(4)] = d_2333_v38_
            nw346_[int(5)] = d_2333_v38_
            nw346_[int(6)] = d_2333_v38_
            nw346_[int(7)] = (d_2333_v38_ if self.f11 else d_2333_v38_)
            nw346_[int(8)] = d_2333_v38_
            nw346_[int(9)] = d_2333_v38_
            nw346_[int(10)] = d_2333_v38_
            nw346_[int(11)] = d_2333_v38_
            nw346_[int(12)] = d_2333_v38_
            nw346_[int(13)] = d_2333_v38_
            nw346_[int(14)] = d_2333_v38_
            nw346_[int(15)] = (d_2351_v51_)
            nw346_[int(16)] = d_2333_v38_
            nw346_[int(17)] = (d_2333_v38_)
            d_2352_v52_ = nw346_
            d_2352_v52_ = d_2352_v52_
            index333_ = default__.safeIndex(893, (p0).length(0))
            (p0)[index333_] = default__.safeDivisionInt(((d_2288_v1_).f25) - ((self).f12), (d_2288_v1_).f25)
            d_2353_v53_: str
            d_2353_v53_ = _dafny.CodePoint('j')
            d_2354_v54_: _dafny.Seq
            d_2354_v54_ = _dafny.SeqWithoutIsStrInference([d_2353_v53_])
            d_2355_v55_: _dafny.Map
            d_2355_v55_ = _dafny.Map({d_2349_v49_: 415})
            d_2356_v57_: _dafny.Seq
            d_2356_v57_ = _dafny.SeqWithoutIsStrInference([p0, d_2348_cf59_])
            d_2357_v58_: _dafny.Map
            d_2357_v58_ = _dafny.Map({len(d_2356_v57_): p1})
            index334_ = default__.safeIndex(893, (p0).length(0))
            def iife203_():
                coll93_ = _dafny.Map()
                compr_93_: _dafny.MultiSet
                for compr_93_ in (_dafny.SeqWithoutIsStrInference([d_2349_v49_, _dafny.MultiSet([(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]])])).Elements:
                    d_2358_v56_: _dafny.MultiSet = compr_93_
                    if (d_2358_v56_) in (_dafny.SeqWithoutIsStrInference([d_2349_v49_, _dafny.MultiSet([(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]])])):
                        coll93_[d_2358_v56_] = (d_2288_v1_).f25
                return _dafny.Map(coll93_)
            rhs341_ = (self).f13
            rhs342_ = d_2354_v54_
            rhs343_ = iife203_()

            rhs344_ = d_2287_v0_
            rhs345_ = (((self).f13) + ((d_2288_v1_).f25)) * ((((d_2357_v58_)[(d_2288_v1_).f25] if ((d_2288_v1_).f25) in (d_2357_v58_) else len(d_2357_v58_))) - ((d_2288_v1_).f25))
            lhs219_ = p0
            lhs220_ = default__.safeIndex(893, (p0).length(0))
            lhs219_[lhs220_] = rhs341_
            d_2354_v54_ = rhs342_
            d_2355_v55_ = rhs343_
            d_2287_v0_ = rhs344_
            r0 = rhs345_
        elif source36_.is_DC35:
            d_2359___mcc_h2_ = source36_.cf60
            d_2360_cf60_ = d_2359___mcc_h2_
            d_2334_v39_ = d_2334_v39_
            d_2361_v59_: _dafny.Array
            d_2362_v60_: bool
            out37_: _dafny.Array
            out38_: bool
            out37_, out38_ = default__.m0(globalState)
            d_2361_v59_ = out37_
            d_2362_v60_ = out38_
            index335_ = default__.safeIndex(893, (p0).length(0))
            (p0)[index335_] = default__.safeModuloInt((d_2288_v1_).f25, (self).f13)
            d_2363_v61_: _dafny.Seq
            d_2363_v61_ = _dafny.SeqWithoutIsStrInference([p0, d_2334_v39_])
            d_2364_v62_: _dafny.Array
            nw347_ = _dafny.Array(None, 8)
            nw347_[int(0)] = (d_2363_v61_)[default__.safeIndex(len(default__.fm50(globalState)), len(d_2363_v61_))]
            nw347_[int(1)] = d_2334_v39_
            nw347_[int(2)] = p0
            nw347_[int(3)] = p0
            nw347_[int(4)] = d_2334_v39_
            nw347_[int(5)] = p0
            nw347_[int(6)] = d_2334_v39_
            nw347_[int(7)] = p0
            d_2364_v62_ = nw347_
            index336_ = default__.safeIndex(203, (d_2364_v62_).length(0))
            (d_2364_v62_)[index336_] = p0
            d_2365_v63_: _dafny.Seq
            d_2365_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hate"))
            d_2366_v64_: _dafny.Set
            d_2366_v64_ = _dafny.Set({(d_2288_v1_).f25, (self).f19, p1, p1})
            d_2367_v66_: str
            d_2367_v66_ = _dafny.CodePoint('j')
            d_2368_v67_: _dafny.Map
            d_2368_v67_ = _dafny.Map({d_2367_v66_: self.f11})
            d_2369_v68_: _dafny.Map
            d_2369_v68_ = _dafny.Map({(d_2288_v1_).f25: (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]})
            index337_ = default__.safeIndex(203, (d_2364_v62_).length(0))
            index338_ = default__.safeIndex(893, (p0).length(0))
            index339_ = default__.safeIndex(893, (p0).length(0))
            def iife204_():
                coll94_ = _dafny.Set()
                compr_94_: int
                for compr_94_ in (d_2366_v64_).Elements:
                    d_2370_v65_: int = compr_94_
                    if (d_2370_v65_) in (d_2366_v64_):
                        coll94_ = coll94_.union(_dafny.Set([default__.safeModuloInt(d_2370_v65_, len(_dafny.SeqWithoutIsStrInference([False, True])))]))
                return _dafny.Set(coll94_)
            rhs346_ = p0
            rhs347_ = (d_2288_v1_).f25
            rhs348_ = (self).fm8(default__.safeDivisionInt((self).f12, len(iife204_()
            )), (d_2368_v67_) | (_dafny.Map({d_2367_v66_: self.f11})), _dafny.Set({p1, (p0)[default__.safeIndex(893, (p0).length(0))], (self).f19, len(d_2369_v68_), (0) - (len(d_2369_v68_))}), globalState)
            rhs349_ = default__.fm2(globalState)
            lhs221_ = d_2364_v62_
            lhs222_ = default__.safeIndex(203, (d_2364_v62_).length(0))
            lhs223_ = p0
            lhs224_ = default__.safeIndex(893, (p0).length(0))
            lhs225_ = p0
            lhs226_ = default__.safeIndex(893, (p0).length(0))
            lhs221_[lhs222_] = rhs346_
            lhs223_[lhs224_] = rhs347_
            d_2365_v63_ = rhs348_
            lhs225_[lhs226_] = rhs349_
        elif True:
            d_2371___mcc_h3_ = source36_.cf57
            d_2372_cf57_ = d_2371___mcc_h3_
            if False:
                d_2333_v38_ = d_2333_v38_
                d_2373_v69_: _dafny.Seq
                d_2373_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsy"))
                d_2374_v70_: _dafny.Map
                d_2374_v70_ = _dafny.Map({-635: d_2373_v69_})
                d_2375_v71_: str
                d_2375_v71_ = _dafny.CodePoint('s')
                d_2376_v72_: _dafny.Map
                d_2376_v72_ = _dafny.Map({d_2375_v71_: self.f11})
                d_2377_v73_: _dafny.Set
                d_2377_v73_ = _dafny.Set({p1, (self).f19})
                d_2374_v70_ = (d_2374_v70_).set((d_2288_v1_).f25, ((self).fm8((d_2288_v1_).f25, d_2376_v72_, d_2377_v73_, globalState) if (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))] else d_2373_v69_))
                d_2378_v74_: C6
                nw348_ = C6()
                nw348_.ctor__()
                d_2378_v74_ = nw348_
                d_2379_v75_: _dafny.Array
                nw349_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_2379_v75_ = nw349_
                d_2380_v76_: D4
                d_2380_v76_ = D4_DC12(self.f11, -251, (p0)[default__.safeIndex(893, (p0).length(0))])
                d_2381_v77_: _dafny.Seq
                d_2381_v77_ = _dafny.SeqWithoutIsStrInference([self.f11])
                d_2382_v78_: D7
                d_2382_v78_ = D7_DC23(d_2379_v75_, d_2381_v77_, (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))], (p0)[default__.safeIndex(893, (p0).length(0))])
                d_2379_v75_ = ((d_2382_v78_).cf42 if (d_2380_v76_).cf21 else d_2379_v75_)
                d_2383_v79_: _dafny.Map
                d_2383_v79_ = _dafny.Map({d_2373_v69_: (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]})
                (self).f11 = (d_2373_v69_) in (d_2383_v79_)
            elif True:
                index340_ = default__.safeIndex(886, (d_2333_v38_).length(0))
                (d_2333_v38_)[index340_] = (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]
                d_2384_v80_: str
                d_2384_v80_ = _dafny.CodePoint('a')
                d_2385_v81_: _dafny.Map
                d_2385_v81_ = _dafny.Map({(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]: False})
                d_2386_v82_: _dafny.Set
                d_2386_v82_ = _dafny.Set({d_2385_v81_})
                d_2387_v83_: _dafny.Map
                d_2387_v83_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2384_v80_ for d_2388_i7_ in range(default__.abs(472))])): default__.fm2(globalState)})
                d_2389_v84_: _dafny.Set
                d_2389_v84_ = _dafny.Set({default__.safeDivisionInt((0) - (len(d_2386_v82_)), (d_2288_v1_).f25), (d_2288_v1_).f25, default__.fm2(globalState), len((d_2387_v83_ if not(self.f11) else d_2387_v83_)), ((self).f12) + ((d_2288_v1_).f25)})
                d_2390_v85_: _dafny.Map
                d_2390_v85_ = _dafny.Map({(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]: d_2384_v80_})
                rhs350_ = ((d_2390_v85_)[(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]] if ((d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]) in (d_2390_v85_) else d_2384_v80_)
                rhs351_ = (_dafny.Set({(self).f12, (d_2288_v1_).f25, p1, (p0)[default__.safeIndex(893, (p0).length(0))]})) | (default__.fm50(globalState))
                d_2384_v80_ = rhs350_
                d_2389_v84_ = rhs351_
                d_2391_v86_: _dafny.Map
                d_2391_v86_ = _dafny.Map({len(d_2385_v81_): self.f11})
                d_2391_v86_ = (d_2391_v86_).set(default__.safeModuloInt(-393, (d_2288_v1_).f25), self.f11)
                d_2392_v87_: _dafny.Set
                d_2392_v87_ = _dafny.Set({(p0 if (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))] else p0), d_2334_v39_, p0, p0})
                d_2392_v87_ = d_2392_v87_
                r1 = not(not((self).fm13(default__.fm1((d_2288_v1_).f25, (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))], (self).f13, globalState), globalState)))
            if not(self.f11):
                d_2393_v88_: C6
                nw350_ = C6()
                nw350_.ctor__()
                d_2393_v88_ = nw350_
                d_2394_v89_: _dafny.Map
                d_2394_v89_ = _dafny.Map({self.f11: (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]})
                d_2395_v90_: str
                d_2395_v90_ = _dafny.CodePoint('d')
                d_2396_v91_: _dafny.Seq
                d_2396_v91_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), default__.fm45(len(d_2394_v89_), globalState), default__.fm45((d_2288_v1_).f25, globalState), d_2395_v90_, d_2395_v90_])
                d_2396_v91_ = _dafny.SeqWithoutIsStrInference([d_2395_v90_ for d_2397_i8_ in range(default__.abs(280))])
                d_2398_v92_: _dafny.MultiSet
                d_2398_v92_ = _dafny.MultiSet([d_2395_v90_])
                d_2399_v93_: _dafny.Array
                nw351_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2399_v93_ = nw351_
                d_2400_v94_: D7
                d_2400_v94_ = D7_DC22(((d_2398_v92_)[_dafny.CodePoint('u')] if (_dafny.CodePoint('u')) in (d_2398_v92_) else 235), d_2399_v93_)
                d_2401_v95_: D7
                d_2401_v95_ = D7_DC24(d_2400_v94_)
                d_2401_v95_ = d_2401_v95_
                d_2402_v96_: _dafny.Array
                d_2403_v97_: bool
                out39_: _dafny.Array
                out40_: bool
                out39_, out40_ = default__.m0(globalState)
                d_2402_v96_ = out39_
                d_2403_v97_ = out40_
                d_2404_v98_: _dafny.Set
                d_2404_v98_ = _dafny.Set({default__.fm58(754, globalState), D15_DC41(d_2403_v97_)})
                d_2405_v99_: D17
                d_2405_v99_ = D17_DC44(d_2404_v98_)
                index341_ = default__.safeIndex(886, (d_2333_v38_).length(0))
                (d_2333_v38_)[index341_] = ((d_2405_v99_).cf68).issubset(d_2404_v98_)
            elif True:
                d_2406_v100_: str
                d_2406_v100_ = _dafny.CodePoint('t')
                d_2407_v101_: _dafny.Map
                d_2407_v101_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2406_v100_]): (p0)[default__.safeIndex(893, (p0).length(0))]})
                d_2408_v103_: _dafny.Map
                d_2408_v103_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2406_v100_]): self.f11})
                def iife205_():
                    coll95_ = _dafny.Map()
                    compr_95_: _dafny.Seq
                    for compr_95_ in (d_2408_v103_).keys.Elements:
                        d_2409_v102_: _dafny.Seq = compr_95_
                        if (d_2409_v102_) in (d_2408_v103_):
                            coll95_[d_2409_v102_] = (d_2288_v1_).f25
                    return _dafny.Map(coll95_)
                r1 = (d_2407_v101_) == ((iife205_()
                ) | (d_2407_v101_))
                d_2410_v104_: _dafny.Map
                d_2410_v104_ = _dafny.Map({(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]: p1})
                d_2411_v105_: _dafny.Map
                d_2411_v105_ = _dafny.Map({d_2410_v104_: (d_2288_v1_).f25})
                d_2412_v106_: _dafny.Seq
                d_2412_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuhbaxora"))
                d_2411_v105_ = (d_2411_v105_).set((d_2410_v104_) | (d_2410_v104_), ((d_2288_v1_).f25) + (len(d_2412_v106_)))
                d_2413_v107_: _dafny.Array
                nw352_ = _dafny.Array(None, 27)
                nw352_[int(0)] = p0
                nw352_[int(1)] = p0
                nw352_[int(2)] = p0
                nw352_[int(3)] = p0
                nw352_[int(4)] = d_2334_v39_
                nw352_[int(5)] = p0
                nw352_[int(6)] = d_2334_v39_
                nw352_[int(7)] = d_2334_v39_
                nw352_[int(8)] = p0
                nw352_[int(9)] = p0
                nw352_[int(10)] = p0
                nw352_[int(11)] = d_2334_v39_
                nw352_[int(12)] = d_2334_v39_
                nw352_[int(13)] = p0
                nw352_[int(14)] = d_2334_v39_
                nw352_[int(15)] = d_2334_v39_
                nw352_[int(16)] = p0
                nw352_[int(17)] = p0
                nw352_[int(18)] = d_2334_v39_
                nw352_[int(19)] = p0
                nw352_[int(20)] = d_2334_v39_
                nw352_[int(21)] = d_2334_v39_
                nw352_[int(22)] = d_2334_v39_
                nw352_[int(23)] = d_2334_v39_
                nw352_[int(24)] = p0
                nw352_[int(25)] = p0
                nw352_[int(26)] = d_2334_v39_
                d_2413_v107_ = nw352_
                d_2414_v108_: _dafny.Seq
                d_2414_v108_ = _dafny.SeqWithoutIsStrInference([not(True)])
                (globalState).f8 = (D7_DC23(d_2413_v107_, d_2414_v108_, True, (d_2288_v1_).f25)).cf44
                d_2415_v109_: _dafny.Array
                nw353_ = _dafny.Array(None, 4)
                nw353_[int(0)] = d_2333_v38_
                nw353_[int(1)] = d_2333_v38_
                nw353_[int(2)] = d_2333_v38_
                nw353_[int(3)] = d_2333_v38_
                d_2415_v109_ = nw353_
                rhs352_ = (d_2288_v1_).f25
                rhs353_ = d_2415_v109_
                r0 = rhs352_
                d_2415_v109_ = rhs353_
                index342_ = default__.safeIndex(828, (d_2415_v109_).length(0))
                (d_2415_v109_)[index342_] = d_2333_v38_
                d_2416_v110_: _dafny.MultiSet
                d_2416_v110_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_2288_v1_).f25 for d_2417_i9_ in range(default__.abs(849))])])
                index343_ = default__.safeIndex(893, (p0).length(0))
                index344_ = default__.safeIndex(828, (d_2415_v109_).length(0))
                rhs354_ = self.f11
                rhs355_ = self.f11
                rhs356_ = (d_2416_v110_).ispropersubset(d_2416_v110_)
                rhs357_ = ((p0)[default__.safeIndex(893, (p0).length(0))] if (d_2414_v108_)[default__.safeIndex((p0)[default__.safeIndex(893, (p0).length(0))], len(d_2414_v108_))] else (d_2288_v1_).f25)
                rhs358_ = (d_2333_v38_ if (self.f11 if (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))] else (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]) else d_2333_v38_)
                lhs227_ = self
                lhs228_ = self
                lhs229_ = globalState
                lhs230_ = p0
                lhs231_ = default__.safeIndex(893, (p0).length(0))
                lhs232_ = d_2415_v109_
                lhs233_ = default__.safeIndex(828, (d_2415_v109_).length(0))
                lhs227_.f11 = rhs354_
                lhs228_.f11 = rhs355_
                lhs229_.f8 = rhs356_
                lhs230_[lhs231_] = rhs357_
                lhs232_[lhs233_] = rhs358_
            d_2418_v111_: _dafny.Map
            d_2418_v111_ = _dafny.Map({(d_2288_v1_).f25: 719})
            d_2419_v112_: _dafny.Map
            d_2419_v112_ = _dafny.Map({d_2418_v111_: (d_2287_v0_)[default__.safeIndex((d_2288_v1_).f25, len(d_2287_v0_))]})
            d_2420_v113_: _dafny.Seq
            d_2420_v113_ = _dafny.SeqWithoutIsStrInference([d_2418_v111_, d_2418_v111_])
            d_2419_v112_ = (_dafny.Map({(d_2420_v113_)[default__.safeIndex((self).f13, len(d_2420_v113_))]: (p0)[default__.safeIndex(893, (p0).length(0))]})) | ((d_2419_v112_) | (default__.fm59((self).f13, (d_2288_v1_).f25, self.f11, (d_2288_v1_).f25, globalState)))
            d_2421_v114_: _dafny.Map
            d_2421_v114_ = _dafny.Map({_dafny.Set({self.f11}): True})
            d_2422_v115_: _dafny.Set
            d_2422_v115_ = _dafny.Set({(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]})
            d_2421_v114_ = (d_2421_v114_).set((d_2422_v115_).intersection(_dafny.Set({(d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]})), (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))])
        (globalState).f8 = (d_2333_v38_)[default__.safeIndex(886, (d_2333_v38_).length(0))]
        r0 = (0) - (((self).f19) + ((p0)[default__.safeIndex(893, (p0).length(0))]))
        d_2423_v116_: _dafny.Seq
        d_2423_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iowpq"))
        r1 = (d_2423_v116_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvs")))
        return r0, r1

    @property
    def f19(self):
        return self._f19

class C13:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    def ctor__(self):
        pass
        pass

    def m6(self, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        r3: _dafny.Seq = _dafny.Seq({})
        d_2424_v0_: int
        d_2424_v0_ = 388
        hi11_ = d_2424_v0_
        for d_2425_i0_ in range((0) - (d_2424_v0_), hi11_):
            d_2426_v1_: str
            d_2426_v1_ = _dafny.CodePoint('l')
            d_2427_v2_: _dafny.Seq
            d_2427_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            r0 = (d_2426_v1_) in (d_2427_v2_)
            r0 = ((-350) - (d_2425_i0_)) != (default__.fm2(globalState))
            d_2428_v3_: bool
            d_2428_v3_ = False
            d_2429_v4_: _dafny.MultiSet
            d_2429_v4_ = _dafny.MultiSet([d_2425_i0_])
            d_2430_v5_: D0
            d_2430_v5_ = D0_DC1(d_2425_i0_, d_2427_v2_, d_2425_i0_, d_2428_v3_)
            (globalState).f8 = default__.fm1(d_2425_i0_, not((d_2428_v3_) and (d_2428_v3_)), ((d_2429_v4_)[(d_2430_v5_).cf3] if ((d_2430_v5_).cf3) in (d_2429_v4_) else (0) - (d_2425_i0_)), globalState)
            d_2427_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpaumjqq"))
        d_2431_v6_: bool
        d_2431_v6_ = False
        d_2432_v7_: D0
        d_2432_v7_ = D0_DC3(d_2424_v0_, d_2431_v6_)
        source37_ = d_2432_v7_
        if source37_.is_DC1:
            d_2433___mcc_h0_ = source37_.cf1
            d_2434___mcc_h1_ = source37_.cf2
            d_2435___mcc_h2_ = source37_.cf3
            d_2436___mcc_h3_ = source37_.cf4
            d_2437_cf4_ = d_2436___mcc_h3_
            d_2438_cf3_ = d_2435___mcc_h2_
            d_2439_cf2_ = d_2434___mcc_h1_
            d_2440_cf1_ = d_2433___mcc_h0_
            d_2441_v8_: _dafny.Seq
            d_2441_v8_ = _dafny.SeqWithoutIsStrInference([d_2424_v0_])
            d_2440_cf1_ = (d_2440_cf1_) * (len(d_2441_v8_))
            d_2442_v9_: _dafny.Map
            d_2442_v9_ = _dafny.Map({d_2440_cf1_: d_2431_v6_})
            d_2431_v6_ = ((d_2442_v9_)[default__.safeDivisionInt(d_2440_cf1_, d_2438_cf3_)] if (default__.safeDivisionInt(d_2440_cf1_, d_2438_cf3_)) in (d_2442_v9_) else not(d_2431_v6_))
            if d_2437_cf4_:
                d_2443_v10_: str
                d_2443_v10_ = _dafny.CodePoint('r')
                rhs359_ = d_2431_v6_
                rhs360_ = _dafny.CodePoint('y')
                rhs361_ = d_2440_cf1_
                lhs234_ = globalState
                lhs234_.f8 = rhs359_
                d_2443_v10_ = rhs360_
                d_2424_v0_ = rhs361_
                d_2444_v11_: _dafny.MultiSet
                d_2444_v11_ = _dafny.MultiSet([d_2431_v6_, d_2431_v6_])
                d_2445_v12_: _dafny.Set
                d_2445_v12_ = _dafny.Set({(0) - ((0) - (((d_2444_v11_)[d_2437_cf4_] if (d_2437_cf4_) in (d_2444_v11_) else default__.fm2(globalState)))), d_2424_v0_, -530, len(d_2442_v9_), d_2438_cf3_})
                d_2446_v13_: _dafny.Set
                d_2446_v13_ = _dafny.Set({d_2437_cf4_, not(d_2431_v6_)})
                d_2447_v14_: _dafny.MultiSet
                d_2447_v14_ = _dafny.MultiSet([d_2446_v13_, d_2446_v13_])
                d_2448_v15_: _dafny.Array
                def lambda118_(d_2449_v0_):
                    def lambda119_(d_2450_i1_):
                        return default__.safeModuloInt(d_2450_i1_, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jf"))): d_2449_v0_})))

                    return lambda119_

                init65_ = lambda118_(d_2424_v0_)
                nw354_ = _dafny.Array(None, 4)
                for i0_65_ in range(nw354_.length(0)):
                    nw354_[i0_65_] = init65_(i0_65_)
                d_2448_v15_ = nw354_
                d_2451_v16_: D0
                d_2451_v16_ = D0_DC2((d_2447_v14_).cardinality, d_2448_v15_)
                d_2452_v17_: _dafny.Seq
                d_2452_v17_ = _dafny.SeqWithoutIsStrInference([d_2431_v6_, d_2431_v6_, d_2431_v6_])
                rhs362_ = ((d_2445_v12_) - (_dafny.Set({d_2440_cf1_}))).intersection(d_2445_v12_)
                rhs363_ = (d_2451_v16_).cf5
                rhs364_ = (d_2452_v17_) <= (d_2452_v17_)
                lhs235_ = globalState
                d_2445_v12_ = rhs362_
                r2 = rhs363_
                lhs235_.f8 = rhs364_
                d_2453_v18_: _dafny.Array
                nw355_ = _dafny.Array(False, 8)
                d_2453_v18_ = nw355_
                index345_ = default__.safeIndex(621, (d_2453_v18_).length(0))
                (d_2453_v18_)[index345_] = d_2437_cf4_
                index346_ = default__.safeIndex(621, (d_2453_v18_).length(0))
                (d_2453_v18_)[index346_] = d_2431_v6_
                d_2431_v6_ = (d_2453_v18_)[default__.safeIndex(621, (d_2453_v18_).length(0))]
                d_2454_v19_: D0
                d_2454_v19_ = D0_DC1((0) - (d_2424_v0_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igwwlgl")), d_2440_cf1_, d_2437_cf4_)
                r0 = ((d_2454_v19_).cf2) < (d_2439_cf2_)
            elif True:
                d_2455_v20_: C3
                nw356_ = C3()
                nw356_.ctor__(d_2438_cf3_, -62)
                d_2455_v20_ = nw356_
                d_2456_v21_: _dafny.Array
                def lambda120_(d_2457_v0_):
                    def lambda121_(d_2458_i2_):
                        return (d_2458_i2_) * (d_2457_v0_)

                    return lambda121_

                init66_ = lambda120_(d_2424_v0_)
                nw357_ = _dafny.Array(None, 29)
                for i0_66_ in range(nw357_.length(0)):
                    nw357_[i0_66_] = init66_(i0_66_)
                d_2456_v21_ = nw357_
                index347_ = default__.safeIndex(931, (d_2456_v21_).length(0))
                (d_2456_v21_)[index347_] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpsgkquyo"))), d_2438_cf3_)
                index348_ = default__.safeIndex(931, (d_2456_v21_).length(0))
                (d_2456_v21_)[index348_] = d_2438_cf3_
                d_2459_v22_: _dafny.Array
                nw358_ = _dafny.Array(_dafny.Set({}), 15)
                d_2459_v22_ = nw358_
                d_2459_v22_ = d_2459_v22_
                d_2437_cf4_ = d_2431_v6_
                d_2460_v23_: _dafny.Set
                d_2460_v23_ = _dafny.Set({len(d_2439_cf2_)})
                d_2461_v24_: D17
                d_2461_v24_ = D17_DC45((0) - (len(d_2441_v8_)), d_2460_v23_)
                r2 = (d_2461_v24_).cf69
            d_2462_v25_: _dafny.Array
            nw359_ = _dafny.Array(D7.default()(), 28)
            d_2462_v25_ = nw359_
            d_2463_v26_: str
            d_2463_v26_ = _dafny.CodePoint('j')
            d_2464_v27_: _dafny.MultiSet
            d_2464_v27_ = _dafny.MultiSet([d_2437_cf4_, d_2437_cf4_, True])
            d_2465_v28_: _dafny.Map
            d_2465_v28_ = _dafny.Map({_dafny.Set({d_2463_v26_}): d_2464_v27_})
            d_2466_v29_: D9
            d_2466_v29_ = D9_DC26(d_2463_v26_)
            d_2467_v30_: _dafny.Set
            d_2467_v30_ = _dafny.Set({(d_2466_v29_).cf48, d_2463_v26_})
            d_2468_v31_: D7
            d_2468_v31_ = D7_DC24(D7_DC21(((d_2465_v28_)[d_2467_v30_] if (d_2467_v30_) in (d_2465_v28_) else d_2464_v27_)))
            index349_ = default__.safeIndex(62, (d_2462_v25_).length(0))
            (d_2462_v25_)[index349_] = D7_DC24(D7_DC24((d_2468_v31_).cf46))
            index350_ = default__.safeIndex(62, (d_2462_v25_).length(0))
            (d_2462_v25_)[index350_] = d_2468_v31_
        elif source37_.is_DC2:
            d_2469___mcc_h4_ = source37_.cf5
            d_2470___mcc_h5_ = source37_.cf6
            d_2471_cf6_ = d_2470___mcc_h5_
            d_2472_cf5_ = d_2469___mcc_h4_
            d_2473_v32_: D6
            d_2473_v32_ = D6_DC18(d_2431_v6_)
            d_2473_v32_ = d_2473_v32_
            d_2474_v33_: _dafny.Set
            d_2474_v33_ = _dafny.Set({d_2431_v6_})
            d_2475_v34_: str
            d_2475_v34_ = _dafny.CodePoint('c')
            d_2474_v33_ = (default__.fm4(992, d_2432_v7_, d_2472_cf5_, d_2475_v34_, globalState)) | (_dafny.Set({d_2431_v6_, d_2431_v6_, not(d_2431_v6_), d_2431_v6_, default__.fm1(d_2424_v0_, d_2431_v6_, d_2424_v0_, globalState)}))
            d_2476_v35_: _dafny.Seq
            d_2476_v35_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2472_cf5_)])
            d_2477_v36_: _dafny.Seq
            d_2477_v36_ = d_2476_v35_
            d_2478_v37_: _dafny.Set
            d_2478_v37_ = _dafny.Set({d_2471_cf6_})
            d_2479_v38_: _dafny.Map
            d_2479_v38_ = _dafny.Map({d_2477_v36_: len(d_2478_v37_)})
            d_2480_v40_: _dafny.Seq
            d_2480_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqem"))
            d_2481_v41_: _dafny.Map
            d_2481_v41_ = _dafny.Map({d_2480_v40_: d_2472_cf5_})
            d_2482_v42_: _dafny.Map
            d_2482_v42_ = _dafny.Map({default__.fm1(d_2424_v0_, d_2431_v6_, d_2424_v0_, globalState): d_2481_v41_})
            d_2483_v43_: _dafny.Map
            d_2483_v43_ = _dafny.Map({d_2431_v6_: d_2431_v6_})
            d_2484_v44_: T2
            nw360_ = C9()
            def iife206_():
                coll96_ = _dafny.Set()
                compr_96_: _dafny.Seq
                for compr_96_ in (d_2479_v38_).keys.Elements:
                    d_2485_v39_: _dafny.Seq = compr_96_
                    if (d_2485_v39_) in (d_2479_v38_):
                        coll96_ = coll96_.union(_dafny.Set([d_2485_v39_]))
                return _dafny.Set(coll96_)
            nw360_.ctor__(iife206_()
            , (len(((d_2482_v42_)[not(d_2431_v6_)] if (not(d_2431_v6_)) in (d_2482_v42_) else d_2481_v41_))) * (len(d_2483_v43_)), d_2424_v0_, d_2472_cf5_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))) != (d_2480_v40_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxog")))
            d_2484_v44_ = nw360_
            index351_ = default__.safeIndex(96, (d_2471_cf6_).length(0))
            (d_2471_cf6_)[index351_] = 804
            index352_ = default__.safeIndex(96, (d_2471_cf6_).length(0))
            rhs365_ = d_2484_v44_
            rhs366_ = (d_2471_cf6_ if d_2431_v6_ else (d_2471_cf6_ if True else d_2471_cf6_))
            rhs367_ = d_2431_v6_
            rhs368_ = ((default__.fm2(globalState)) * (d_2424_v0_)) + (d_2424_v0_)
            rhs369_ = d_2431_v6_
            lhs236_ = globalState
            lhs237_ = d_2471_cf6_
            lhs238_ = default__.safeIndex(96, (d_2471_cf6_).length(0))
            d_2484_v44_ = rhs365_
            d_2471_cf6_ = rhs366_
            lhs236_.f8 = rhs367_
            lhs237_[lhs238_] = rhs368_
            d_2431_v6_ = rhs369_
            d_2431_v6_ = d_2431_v6_
        elif source37_.is_DC3:
            d_2486___mcc_h6_ = source37_.cf7
            d_2487___mcc_h7_ = source37_.cf8
            d_2488_cf8_ = d_2487___mcc_h7_
            d_2489_cf7_ = d_2486___mcc_h6_
            d_2490_v45_: _dafny.Seq
            d_2490_v45_ = _dafny.SeqWithoutIsStrInference([d_2431_v6_, d_2488_cf8_, d_2488_cf8_])
            r3 = (d_2490_v45_) + (d_2490_v45_)
            d_2491_v46_: _dafny.Seq
            d_2491_v46_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2424_v0_)])
            d_2492_v48_: _dafny.Map
            def iife207_():
                coll97_ = _dafny.Map()
                compr_97_: int
                for compr_97_ in _dafny.IntegerRange(407, 236):
                    d_2493_v47_: int = compr_97_
                    if ((407) <= (d_2493_v47_)) and ((d_2493_v47_) < (236)):
                        coll97_[(d_2493_v47_) - (d_2424_v0_)] = d_2488_cf8_
                return _dafny.Map(coll97_)
            d_2492_v48_ = _dafny.Map({d_2491_v46_: iife207_()
            })
            d_2494_v49_: _dafny.Map
            d_2494_v49_ = _dafny.Map({453: d_2431_v6_})
            d_2495_v50_: _dafny.Array
            nw361_ = _dafny.Array(None, 3)
            nw361_[int(0)] = d_2492_v48_
            nw361_[int(1)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([617]): (d_2494_v49_).set(d_2424_v0_, d_2431_v6_)})
            nw361_[int(2)] = d_2492_v48_
            d_2495_v50_ = nw361_
            index353_ = default__.safeIndex(799, (d_2495_v50_).length(0))
            (d_2495_v50_)[index353_] = (d_2492_v48_).set(d_2491_v46_, (d_2494_v49_).set(-790, True))
            d_2496_v51_: _dafny.Map
            d_2496_v51_ = _dafny.Map({d_2488_cf8_: d_2424_v0_})
            d_2497_v52_: _dafny.Set
            d_2497_v52_ = _dafny.Set({d_2494_v49_})
            d_2498_v53_: str
            d_2498_v53_ = _dafny.CodePoint('x')
            d_2499_v54_: _dafny.Map
            d_2499_v54_ = _dafny.Map({d_2498_v53_: d_2489_cf7_})
            d_2500_v55_: _dafny.Map
            d_2500_v55_ = _dafny.Map({d_2488_cf8_: True})
            index354_ = default__.safeIndex(799, (d_2495_v50_).length(0))
            rhs370_ = ((d_2492_v48_ if d_2431_v6_ else _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2489_cf7_ for d_2501_i3_ in range(default__.abs(57))]): default__.fm47(True, d_2488_cf8_, d_2496_v51_, d_2424_v0_, globalState)}))).set(d_2491_v46_, default__.fm47(True, d_2488_cf8_, _dafny.Map({d_2488_cf8_: len(d_2497_v52_)}), d_2489_cf7_, globalState))
            rhs371_ = (d_2431_v6_) in ((d_2490_v45_) + (d_2490_v45_))
            rhs372_ = d_2431_v6_
            rhs373_ = (default__.safeModuloInt(d_2424_v0_, d_2489_cf7_)) * (((d_2499_v54_)[d_2498_v53_] if (d_2498_v53_) in (d_2499_v54_) else d_2489_cf7_))
            rhs374_ = (d_2500_v55_) | ((d_2500_v55_) | (d_2500_v55_))
            lhs239_ = d_2495_v50_
            lhs240_ = default__.safeIndex(799, (d_2495_v50_).length(0))
            lhs241_ = globalState
            lhs239_[lhs240_] = rhs370_
            lhs241_.f8 = rhs371_
            d_2431_v6_ = rhs372_
            d_2489_cf7_ = rhs373_
            r1 = rhs374_
            if d_2431_v6_:
                d_2502_v56_: C1
                nw362_ = C1()
                nw362_.ctor__(d_2424_v0_)
                d_2502_v56_ = nw362_
                d_2503_v57_: _dafny.Array
                nw363_ = _dafny.Array(False, 19)
                d_2503_v57_ = nw363_
                d_2504_v58_: _dafny.Map
                d_2504_v58_ = _dafny.Map({d_2503_v57_: d_2431_v6_})
                d_2505_v59_: _dafny.Map
                d_2505_v59_ = _dafny.Map({d_2491_v46_: (d_2504_v58_) == (d_2504_v58_)})
                d_2505_v59_ = (d_2505_v59_).set((_dafny.SeqWithoutIsStrInference([d_2502_v56_.f27, d_2424_v0_])).set(default__.safeIndex(len(_dafny.Set({d_2502_v56_.f27})), len(_dafny.SeqWithoutIsStrInference([d_2502_v56_.f27, d_2424_v0_]))), len(_dafny.SeqWithoutIsStrInference([d_2498_v53_ for d_2506_i4_ in range(default__.abs(901))]))), (d_2490_v45_) == (d_2490_v45_))
                (globalState).f8 = (d_2490_v45_)[default__.safeIndex(d_2489_cf7_, len(d_2490_v45_))]
                d_2431_v6_ = d_2488_cf8_
                d_2498_v53_ = _dafny.CodePoint('n')
            elif True:
                r2 = d_2489_cf7_
                d_2507_v60_: _dafny.Set
                d_2507_v60_ = _dafny.Set({False, d_2488_cf8_})
                r2 = (d_2424_v0_) + (len(d_2507_v60_))
                d_2508_v61_: _dafny.Map
                d_2508_v61_ = _dafny.Map({d_2424_v0_: d_2498_v53_})
                d_2509_v62_: _dafny.Map
                d_2509_v62_ = _dafny.Map({d_2508_v61_: default__.fm2(globalState)})
                def iife208_():
                    coll98_ = _dafny.Map()
                    compr_98_: int
                    for compr_98_ in (d_2494_v49_).keys.Elements:
                        d_2510_v63_: int = compr_98_
                        if (d_2510_v63_) in (d_2494_v49_):
                            coll98_[(d_2510_v63_) - ((d_2491_v46_)[default__.safeIndex(d_2424_v0_, len(d_2491_v46_))])] = d_2498_v53_
                    return _dafny.Map(coll98_)
                d_2509_v62_ = (d_2509_v62_).set((d_2508_v61_) | (iife208_()
                ), (0) - (d_2424_v0_))
                d_2511_v64_: _dafny.Seq
                d_2511_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nivaav"))
                d_2512_v65_: _dafny.Array
                nw364_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_2512_v65_ = nw364_
                d_2513_v66_: D7
                d_2513_v66_ = D7_DC23(d_2512_v65_, d_2490_v45_, d_2488_cf8_, d_2424_v0_)
                d_2514_v67_: _dafny.Map
                d_2514_v67_ = _dafny.Map({(0) - ((d_2424_v0_ if d_2488_cf8_ else len(d_2511_v64_))): d_2513_v66_})
                d_2514_v67_ = (d_2514_v67_).set(d_2489_cf7_, d_2513_v66_)
                d_2515_v68_: _dafny.Array
                nw365_ = _dafny.Array(False, 22)
                d_2515_v68_ = nw365_
                d_2515_v68_ = d_2515_v68_
            d_2516_v69_: _dafny.Array
            def lambda122_(d_2517_v45_, d_2518_v6_):
                def lambda123_(d_2519_i5_):
                    return D15_DC40(_dafny.Map({d_2517_v45_: D7_DC21(_dafny.MultiSet([d_2518_v6_]))}))

                return lambda123_

            init67_ = lambda122_(d_2490_v45_, d_2431_v6_)
            nw366_ = _dafny.Array(None, 7)
            for i0_67_ in range(nw366_.length(0)):
                nw366_[i0_67_] = init67_(i0_67_)
            d_2516_v69_ = nw366_
            d_2520_v70_: D7
            d_2520_v70_ = D7_DC21(default__.fm30(d_2431_v6_, d_2431_v6_, d_2489_cf7_, globalState))
            d_2521_v71_: _dafny.Map
            d_2521_v71_ = _dafny.Map({d_2490_v45_: d_2520_v70_})
            d_2522_v72_: D15
            d_2522_v72_ = D15_DC40(d_2521_v71_)
            index355_ = default__.safeIndex(562, (d_2516_v69_).length(0))
            (d_2516_v69_)[index355_] = d_2522_v72_
            pat_let_tv91_ = d_2489_cf7_
            pat_let_tv92_ = d_2424_v0_
            pat_let_tv93_ = globalState
            index356_ = default__.safeIndex(562, (d_2516_v69_).length(0))
            def iife209_(_pat_let55_0):
                def iife210_(d_2523_dt__update__tmp_h0_):
                    def iife211_(_pat_let56_0):
                        def iife212_(d_2524_dt__update_hcf64_h0_):
                            return D15_DC40(d_2524_dt__update_hcf64_h0_)
                        return iife212_(_pat_let56_0)
                    return iife211_(default__.fm55(488, pat_let_tv91_, pat_let_tv92_, pat_let_tv93_))
                return iife210_(_pat_let55_0)
            (d_2516_v69_)[index356_] = iife209_(d_2522_v72_)
        elif True:
            d_2525___mcc_h8_ = source37_.cf0
            d_2526_cf0_ = d_2525___mcc_h8_
            d_2527_v73_: _dafny.MultiSet
            d_2527_v73_ = _dafny.MultiSet([(d_2424_v0_) + (d_2424_v0_), (d_2424_v0_ if not(not(d_2431_v6_)) else len(_dafny.Map({d_2424_v0_: d_2424_v0_}))), (0) - (d_2424_v0_)])
            d_2528_v74_: _dafny.MultiSet
            d_2528_v74_ = _dafny.MultiSet([d_2431_v6_, True])
            d_2527_v73_ = ((d_2527_v73_) - (d_2527_v73_) if (d_2528_v74_) == ((d_2528_v74_).set(d_2431_v6_, default__.abs(d_2424_v0_))) else (d_2527_v73_).intersection((d_2527_v73_).set(d_2424_v0_, default__.abs(84))))
            d_2529_v75_: _dafny.Array
            d_2530_v76_: bool
            out41_: _dafny.Array
            out42_: bool
            out41_, out42_ = default__.m0(globalState)
            d_2529_v75_ = out41_
            d_2530_v76_ = out42_
            (globalState).f8 = ((default__.fm2(globalState)) - (d_2424_v0_)) != (((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja"))))) * (d_2424_v0_))
            d_2531_v77_: _dafny.Map
            d_2531_v77_ = _dafny.Map({d_2530_v76_: d_2431_v6_})
            d_2532_v78_: _dafny.Array
            nw367_ = _dafny.Array(None, 26)
            nw367_[int(0)] = d_2424_v0_
            nw367_[int(1)] = d_2424_v0_
            nw367_[int(2)] = d_2424_v0_
            nw367_[int(3)] = d_2424_v0_
            nw367_[int(4)] = d_2424_v0_
            nw367_[int(5)] = d_2424_v0_
            nw367_[int(6)] = len(d_2531_v77_)
            nw367_[int(7)] = d_2424_v0_
            nw367_[int(8)] = d_2424_v0_
            nw367_[int(9)] = len(d_2531_v77_)
            nw367_[int(10)] = 718
            nw367_[int(11)] = d_2424_v0_
            nw367_[int(12)] = d_2424_v0_
            nw367_[int(13)] = 329
            nw367_[int(14)] = 895
            nw367_[int(15)] = d_2424_v0_
            nw367_[int(16)] = 825
            nw367_[int(17)] = -284
            nw367_[int(18)] = d_2424_v0_
            nw367_[int(19)] = d_2424_v0_
            nw367_[int(20)] = d_2424_v0_
            nw367_[int(21)] = default__.fm2(globalState)
            nw367_[int(22)] = d_2424_v0_
            nw367_[int(23)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2533_i6_ in range(default__.abs(985))]))
            nw367_[int(24)] = 468
            nw367_[int(25)] = 115
            d_2532_v78_ = nw367_
            d_2534_v79_: _dafny.Array
            nw368_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_2534_v79_ = nw368_
            d_2535_v80_: _dafny.Seq
            d_2535_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gka"))
            d_2536_v81_: _dafny.Map
            d_2536_v81_ = _dafny.Map({d_2535_v80_: d_2431_v6_})
            d_2537_v82_: _dafny.Seq
            d_2537_v82_ = _dafny.SeqWithoutIsStrInference([d_2530_v76_, ((d_2536_v81_)[d_2535_v80_] if (d_2535_v80_) in (d_2536_v81_) else d_2530_v76_)])
            d_2538_v83_: D7
            d_2538_v83_ = D7_DC23(d_2534_v79_, d_2537_v82_, d_2530_v76_, d_2424_v0_)
            d_2539_v84_: _dafny.Map
            d_2539_v84_ = _dafny.Map({d_2532_v78_: d_2538_v83_})
            d_2540_v85_: _dafny.Map
            d_2540_v85_ = _dafny.Map({d_2424_v0_: d_2537_v82_})
            d_2541_v86_: _dafny.Map
            d_2541_v86_ = _dafny.Map({d_2540_v85_: d_2530_v76_})
            d_2542_v87_: _dafny.Seq
            d_2542_v87_ = _dafny.SeqWithoutIsStrInference([len(d_2531_v77_)])
            d_2543_v88_: _dafny.MultiSet
            d_2543_v88_ = _dafny.MultiSet([d_2542_v87_])
            rhs375_ = (len(d_2539_v84_)) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2544_i7_ in range(default__.abs(545))])))
            rhs376_ = (((d_2528_v74_)[d_2431_v6_] if (d_2431_v6_) in (d_2528_v74_) else (0) - (d_2424_v0_))) > (d_2424_v0_)
            rhs377_ = d_2431_v6_
            rhs378_ = not(((d_2541_v86_)[d_2540_v85_] if (d_2540_v85_) in (d_2541_v86_) else not((d_2543_v88_).isdisjoint(_dafny.MultiSet([d_2542_v87_])))))
            rhs379_ = d_2530_v76_
            lhs242_ = globalState
            r2 = rhs375_
            r0 = rhs376_
            lhs242_.f8 = rhs377_
            d_2530_v76_ = rhs378_
            d_2530_v76_ = rhs379_
        d_2545_i8_: int
        d_2545_i8_ = 0
        with _dafny.label("11"):
            while (d_2431_v6_ if (d_2431_v6_ if d_2431_v6_ else True) else d_2431_v6_):
                with _dafny.c_label("11"):
                    if (d_2545_i8_) >= (100):
                        raise _dafny.Break("11")
                    d_2545_i8_ = (d_2545_i8_) + (1)
                    d_2546_v89_: _dafny.Map
                    d_2546_v89_ = _dafny.Map({d_2431_v6_: d_2431_v6_})
                    d_2547_v90_: _dafny.Set
                    d_2547_v90_ = _dafny.Set({d_2431_v6_})
                    d_2548_v91_: C11
                    nw369_ = C11()
                    nw369_.ctor__(((0) - (d_2424_v0_) if d_2431_v6_ else len(d_2546_v89_)), len(d_2547_v90_))
                    d_2548_v91_ = nw369_
                    d_2549_v92_: _dafny.Seq
                    d_2549_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jf"))
                    d_2550_v93_: str
                    d_2550_v93_ = _dafny.CodePoint('u')
                    d_2551_v94_: C2
                    nw370_ = C2()
                    nw370_.ctor__(d_2424_v0_, d_2549_v92_, d_2424_v0_, len(default__.fm4(d_2424_v0_, d_2432_v7_, d_2424_v0_, d_2550_v93_, globalState)), d_2431_v6_)
                    d_2551_v94_ = nw370_
                    d_2552_v95_: _dafny.Map
                    d_2552_v95_ = _dafny.Map({d_2551_v94_: not (d_2431_v6_) or (d_2431_v6_)})
                    d_2552_v95_ = (d_2552_v95_).set(d_2551_v94_, d_2431_v6_)
                    d_2553_v96_: _dafny.Map
                    d_2553_v96_ = _dafny.Map({d_2424_v0_: d_2431_v6_})
                    d_2554_v97_: _dafny.Map
                    d_2554_v97_ = _dafny.Map({((d_2553_v96_)[d_2424_v0_] if (d_2424_v0_) in (d_2553_v96_) else d_2431_v6_): 831})
                    d_2555_v98_: _dafny.Map
                    d_2555_v98_ = _dafny.Map({d_2431_v6_: (len(d_2554_v97_)) * (d_2424_v0_)})
                    d_2555_v98_ = d_2555_v98_
                    d_2424_v0_ = (d_2424_v0_) * ((d_2551_v94_).f26)
                    pass
            pass
        d_2556_v99_: _dafny.Array
        def lambda124_(d_2557_v0_):
            def lambda125_(d_2558_i9_):
                return (d_2558_i9_) - (len(_dafny.Map({d_2557_v0_: d_2557_v0_})))

            return lambda125_

        init68_ = lambda124_(d_2424_v0_)
        nw371_ = _dafny.Array(None, 20)
        for i0_68_ in range(nw371_.length(0)):
            nw371_[i0_68_] = init68_(i0_68_)
        d_2556_v99_ = nw371_
        d_2559_v100_: _dafny.Array
        nw372_ = _dafny.Array(None, 26)
        nw372_[int(0)] = d_2556_v99_
        nw372_[int(1)] = d_2556_v99_
        nw372_[int(2)] = d_2556_v99_
        nw372_[int(3)] = d_2556_v99_
        nw372_[int(4)] = d_2556_v99_
        nw372_[int(5)] = d_2556_v99_
        nw372_[int(6)] = d_2556_v99_
        nw372_[int(7)] = d_2556_v99_
        nw372_[int(8)] = d_2556_v99_
        nw372_[int(9)] = d_2556_v99_
        nw372_[int(10)] = d_2556_v99_
        nw372_[int(11)] = d_2556_v99_
        nw372_[int(12)] = d_2556_v99_
        nw372_[int(13)] = d_2556_v99_
        nw372_[int(14)] = d_2556_v99_
        nw372_[int(15)] = d_2556_v99_
        nw372_[int(16)] = d_2556_v99_
        nw372_[int(17)] = d_2556_v99_
        nw372_[int(18)] = d_2556_v99_
        nw372_[int(19)] = d_2556_v99_
        nw372_[int(20)] = d_2556_v99_
        nw372_[int(21)] = d_2556_v99_
        nw372_[int(22)] = d_2556_v99_
        nw372_[int(23)] = d_2556_v99_
        nw372_[int(24)] = d_2556_v99_
        nw372_[int(25)] = d_2556_v99_
        d_2559_v100_ = nw372_
        d_2560_v101_: _dafny.Seq
        d_2560_v101_ = _dafny.SeqWithoutIsStrInference([d_2431_v6_, d_2431_v6_, d_2431_v6_])
        d_2561_v102_: D7
        d_2561_v102_ = D7_DC23(d_2559_v100_, (d_2560_v101_) + (d_2560_v101_), d_2431_v6_, d_2424_v0_)
        source38_ = d_2561_v102_
        if source38_.is_DC22:
            d_2562___mcc_h9_ = source38_.cf40
            d_2563___mcc_h10_ = source38_.cf41
            d_2564_cf41_ = d_2563___mcc_h10_
            d_2565_cf40_ = d_2562___mcc_h9_
            d_2566_v103_: _dafny.Map
            d_2566_v103_ = _dafny.Map({d_2556_v99_: (d_2424_v0_ if d_2431_v6_ else d_2424_v0_)})
            d_2567_v104_: D15
            d_2567_v104_ = D15_DC41(d_2431_v6_)
            d_2568_v105_: D15
            d_2568_v105_ = D15_DC42(d_2567_v104_)
            d_2569_v106_: D15
            d_2569_v106_ = D15_DC42(d_2568_v105_)
            d_2570_v107_: _dafny.Map
            d_2570_v107_ = _dafny.Map({d_2569_v106_: d_2424_v0_})
            d_2566_v103_ = (d_2566_v103_).set(d_2556_v99_, (0) - (len(d_2570_v107_)))
            if not ((d_2431_v6_ if d_2431_v6_ else d_2431_v6_)) or (not((len(_dafny.Map({d_2424_v0_: default__.fm1(d_2424_v0_, d_2431_v6_, d_2424_v0_, globalState)}))) == (d_2565_cf40_))):
                d_2571_v108_: _dafny.Array
                nw373_ = _dafny.Array(_dafny.Seq({}), 27)
                d_2571_v108_ = nw373_
                index357_ = default__.safeIndex(18, (d_2571_v108_).length(0))
                (d_2571_v108_)[index357_] = (d_2560_v101_) + ((_dafny.SeqWithoutIsStrInference([d_2431_v6_])).set(default__.safeIndex(d_2424_v0_, len(_dafny.SeqWithoutIsStrInference([d_2431_v6_]))), d_2431_v6_))
                d_2572_v109_: _dafny.Seq
                d_2572_v109_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oevt"))
                d_2573_v110_: _dafny.Map
                d_2573_v110_ = _dafny.Map({d_2431_v6_: d_2431_v6_})
                d_2574_v111_: _dafny.Seq
                d_2574_v111_ = _dafny.SeqWithoutIsStrInference([d_2573_v110_, d_2573_v110_])
                d_2575_v112_: _dafny.Seq
                d_2575_v112_ = _dafny.SeqWithoutIsStrInference([len((d_2572_v109_) + (d_2572_v109_)), 798, len((d_2574_v111_)[default__.safeIndex(d_2424_v0_, len(d_2574_v111_))])])
                index358_ = default__.safeIndex(18, (d_2571_v108_).length(0))
                rhs380_ = ((d_2565_cf40_) > (d_2424_v0_)) or (d_2431_v6_)
                rhs381_ = d_2575_v112_
                rhs382_ = not (d_2431_v6_) or ((d_2431_v6_ if d_2431_v6_ else d_2431_v6_))
                rhs383_ = d_2560_v101_
                lhs243_ = globalState
                lhs244_ = globalState
                lhs245_ = globalState
                lhs246_ = d_2571_v108_
                lhs247_ = default__.safeIndex(18, (d_2571_v108_).length(0))
                lhs243_.f8 = rhs380_
                lhs244_.f5 = rhs381_
                lhs245_.f8 = rhs382_
                lhs246_[lhs247_] = rhs383_
                d_2431_v6_ = d_2431_v6_
                r2 = d_2565_cf40_
                d_2576_v113_: _dafny.Array
                nw374_ = _dafny.Array(False, 2)
                d_2576_v113_ = nw374_
                d_2576_v113_ = d_2576_v113_
                index359_ = default__.safeIndex(736, (d_2556_v99_).length(0))
                (d_2556_v99_)[index359_] = d_2424_v0_
                d_2577_v114_: _dafny.MultiSet
                d_2577_v114_ = _dafny.MultiSet([d_2431_v6_, d_2431_v6_, d_2431_v6_])
                d_2578_v115_: _dafny.Seq
                d_2578_v115_ = _dafny.SeqWithoutIsStrInference([d_2577_v114_, (d_2577_v114_).set(True, default__.abs(d_2565_cf40_)), _dafny.MultiSet(d_2560_v101_)])
                index360_ = default__.safeIndex(736, (d_2556_v99_).length(0))
                (d_2556_v99_)[index360_] = len(d_2578_v115_)
            elif True:
                index361_ = default__.safeIndex(464, (d_2556_v99_).length(0))
                (d_2556_v99_)[index361_] = d_2424_v0_
                index362_ = default__.safeIndex(549, (d_2556_v99_).length(0))
                (d_2556_v99_)[index362_] = d_2565_cf40_
                d_2579_v116_: _dafny.Seq
                d_2579_v116_ = _dafny.SeqWithoutIsStrInference([924])
                d_2580_v117_: _dafny.Array
                nw375_ = _dafny.Array(None, 12)
                nw375_[int(0)] = d_2431_v6_
                nw375_[int(1)] = (d_2431_v6_ if d_2431_v6_ else d_2431_v6_)
                nw375_[int(2)] = not((d_2431_v6_) == (False))
                nw375_[int(3)] = d_2431_v6_
                nw375_[int(4)] = d_2431_v6_
                nw375_[int(5)] = d_2431_v6_
                nw375_[int(6)] = d_2431_v6_
                nw375_[int(7)] = d_2431_v6_
                nw375_[int(8)] = default__.fm1(d_2424_v0_, d_2431_v6_, d_2424_v0_, globalState)
                nw375_[int(9)] = d_2431_v6_
                nw375_[int(10)] = default__.fm1(d_2565_cf40_, d_2431_v6_, len(d_2579_v116_), globalState)
                nw375_[int(11)] = (d_2565_cf40_) <= (d_2424_v0_)
                d_2580_v117_ = nw375_
                index363_ = default__.safeIndex(959, (d_2580_v117_).length(0))
                (d_2580_v117_)[index363_] = d_2431_v6_
                d_2581_v118_: D9
                d_2581_v118_ = D9_DC26(_dafny.CodePoint('p'))
                d_2582_v119_: _dafny.Map
                d_2582_v119_ = _dafny.Map({(d_2581_v118_).cf48: False})
                d_2583_v120_: str
                d_2583_v120_ = _dafny.CodePoint('l')
                d_2584_v121_: _dafny.Array
                nw376_ = _dafny.Array(None, 13)
                nw376_[int(0)] = (d_2582_v119_) | (d_2582_v119_)
                nw376_[int(1)] = (d_2582_v119_).set(d_2583_v120_, not(d_2431_v6_))
                nw376_[int(2)] = d_2582_v119_
                nw376_[int(3)] = (D18_DC47(d_2582_v119_)).cf72
                nw376_[int(4)] = d_2582_v119_
                nw376_[int(5)] = default__.fm60(d_2431_v6_, globalState)
                nw376_[int(6)] = d_2582_v119_
                nw376_[int(7)] = default__.fm60(d_2431_v6_, globalState)
                nw376_[int(8)] = d_2582_v119_
                nw376_[int(9)] = d_2582_v119_
                nw376_[int(10)] = d_2582_v119_
                nw376_[int(11)] = d_2582_v119_
                nw376_[int(12)] = d_2582_v119_
                d_2584_v121_ = nw376_
                index364_ = default__.safeIndex(464, (d_2556_v99_).length(0))
                index365_ = default__.safeIndex(549, (d_2556_v99_).length(0))
                index366_ = default__.safeIndex(959, (d_2580_v117_).length(0))
                rhs384_ = (default__.safeDivisionInt((0) - (d_2565_cf40_), d_2565_cf40_)) - ((0) - (d_2565_cf40_))
                rhs385_ = d_2424_v0_
                rhs386_ = d_2431_v6_
                rhs387_ = d_2584_v121_
                lhs248_ = d_2556_v99_
                lhs249_ = default__.safeIndex(464, (d_2556_v99_).length(0))
                lhs250_ = d_2556_v99_
                lhs251_ = default__.safeIndex(549, (d_2556_v99_).length(0))
                lhs252_ = d_2580_v117_
                lhs253_ = default__.safeIndex(959, (d_2580_v117_).length(0))
                lhs248_[lhs249_] = rhs384_
                lhs250_[lhs251_] = rhs385_
                lhs252_[lhs253_] = rhs386_
                d_2584_v121_ = rhs387_
                d_2585_v122_: _dafny.Map
                d_2585_v122_ = _dafny.Map({default__.safeDivisionInt(d_2424_v0_, d_2565_cf40_): d_2424_v0_})
                d_2585_v122_ = (d_2585_v122_).set(118, default__.safeDivisionInt(-368, d_2424_v0_))
                (globalState).f8 = ((d_2580_v117_)[default__.safeIndex(959, (d_2580_v117_).length(0))]) and (False)
                d_2586_v123_: _dafny.MultiSet
                d_2586_v123_ = _dafny.MultiSet([d_2565_cf40_, len(d_2585_v122_)])
                d_2587_v124_: _dafny.MultiSet
                d_2587_v124_ = _dafny.MultiSet([d_2586_v123_])
                d_2588_v125_: _dafny.Seq
                d_2588_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qm"))
                d_2589_v126_: C2
                nw377_ = C2()
                nw377_.ctor__((d_2587_v124_).cardinality, d_2588_v125_, d_2565_cf40_, (d_2556_v99_)[default__.safeIndex(464, (d_2556_v99_).length(0))], (d_2580_v117_)[default__.safeIndex(959, (d_2580_v117_).length(0))])
                d_2589_v126_ = nw377_
                d_2590_v127_: _dafny.MultiSet
                d_2590_v127_ = _dafny.MultiSet([((d_2556_v99_)[default__.safeIndex(464, (d_2556_v99_).length(0))]) > ((d_2589_v126_).f26)])
                d_2590_v127_ = d_2590_v127_
            index367_ = default__.safeIndex(794, (d_2559_v100_).length(0))
            (d_2559_v100_)[index367_] = d_2556_v99_
            index368_ = default__.safeIndex(794, (d_2559_v100_).length(0))
            (d_2559_v100_)[index368_] = (d_2556_v99_ if not(d_2431_v6_) else d_2556_v99_)
            d_2591_v128_: _dafny.Seq
            d_2591_v128_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqbfsoug"))
            d_2591_v128_ = (D0_DC1(default__.fm2(globalState), d_2591_v128_, d_2424_v0_, d_2431_v6_)).cf2
        elif source38_.is_DC23:
            d_2592___mcc_h11_ = source38_.cf42
            d_2593___mcc_h12_ = source38_.cf43
            d_2594___mcc_h13_ = source38_.cf44
            d_2595___mcc_h14_ = source38_.cf45
            d_2596_cf45_ = d_2595___mcc_h14_
            d_2597_cf44_ = d_2594___mcc_h13_
            d_2598_cf43_ = d_2593___mcc_h12_
            d_2599_cf42_ = d_2592___mcc_h11_
            d_2600_v129_: _dafny.Seq
            d_2600_v129_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adgdyssy"))
            d_2601_v130_: _dafny.Seq
            d_2601_v130_ = _dafny.SeqWithoutIsStrInference([-629])
            d_2602_v131_: _dafny.MultiSet
            d_2602_v131_ = _dafny.MultiSet([d_2424_v0_])
            d_2603_v132_: _dafny.Seq
            d_2603_v132_ = _dafny.SeqWithoutIsStrInference([d_2596_cf45_, (d_2601_v130_)[default__.safeIndex(d_2424_v0_, len(d_2601_v130_))], d_2596_cf45_, len(d_2560_v101_), ((d_2602_v131_)[d_2424_v0_] if (d_2424_v0_) in (d_2602_v131_) else d_2424_v0_)])
            d_2604_v133_: C2
            nw378_ = C2()
            nw378_.ctor__((d_2424_v0_) + (d_2596_cf45_), d_2600_v129_, len(d_2601_v130_), default__.safeModuloInt((d_2601_v130_)[default__.safeIndex(d_2596_cf45_, len(d_2601_v130_))], d_2596_cf45_), d_2597_cf44_)
            d_2604_v133_ = nw378_
            d_2605_v134_: _dafny.Map
            d_2605_v134_ = _dafny.Map({(d_2604_v133_).f26: d_2600_v129_})
            d_2605_v134_ = (d_2605_v134_).set((d_2604_v133_).f26, d_2600_v129_)
            if d_2431_v6_:
                d_2597_cf44_ = d_2431_v6_
                d_2606_v135_: _dafny.Map
                d_2606_v135_ = _dafny.Map({749: d_2597_cf44_})
                d_2607_v136_: D0
                d_2607_v136_ = D0_DC1(d_2596_cf45_, d_2600_v129_, (d_2604_v133_).f26, d_2597_cf44_)
                d_2608_v137_: str
                d_2608_v137_ = _dafny.CodePoint('n')
                d_2609_v138_: _dafny.Set
                d_2609_v138_ = _dafny.Set({d_2424_v0_})
                d_2610_v139_: _dafny.Array
                nw379_ = _dafny.Array(D0.default()(), 16)
                d_2610_v139_ = nw379_
                d_2611_v140_: _dafny.Set
                d_2611_v140_ = _dafny.Set({d_2610_v139_})
                d_2612_v141_: _dafny.Map
                d_2612_v141_ = _dafny.Map({d_2597_cf44_: len(d_2611_v140_)})
                pat_let_tv94_ = d_2600_v129_
                pat_let_tv95_ = d_2600_v129_
                pat_let_tv96_ = globalState
                pat_let_tv97_ = globalState
                pat_let_tv98_ = d_2608_v137_
                pat_let_tv99_ = globalState
                pat_let_tv100_ = d_2424_v0_
                pat_let_tv101_ = d_2597_cf44_
                d_2613_v142_: _dafny.Array
                nw380_ = _dafny.Array(None, 27)
                nw380_[int(0)] = D0_DC1((d_2601_v130_)[default__.safeIndex(370, len(d_2601_v130_))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oabhtcmf")), len(d_2606_v135_), d_2597_cf44_)
                def iife213_(_pat_let57_0):
                    def iife214_(d_2614_dt__update__tmp_h1_):
                        def iife215_(_pat_let58_0):
                            def iife216_(d_2615_dt__update_hcf3_h0_):
                                def iife217_(_pat_let59_0):
                                    def iife218_(d_2616_dt__update_hcf2_h0_):
                                        return D0_DC1((d_2614_dt__update__tmp_h1_).cf1, d_2616_dt__update_hcf2_h0_, d_2615_dt__update_hcf3_h0_, (d_2614_dt__update__tmp_h1_).cf4)
                                    return iife218_(_pat_let59_0)
                                return iife217_(pat_let_tv95_)
                            return iife216_(_pat_let58_0)
                        return iife215_((0) - (len(pat_let_tv94_)))
                    return iife214_(_pat_let57_0)
                nw380_[int(1)] = iife213_(d_2607_v136_)
                nw380_[int(2)] = (d_2604_v133_).fm5(d_2596_cf45_, globalState)
                nw380_[int(3)] = d_2607_v136_
                nw380_[int(4)] = d_2607_v136_
                nw380_[int(5)] = d_2607_v136_
                def iife219_(_pat_let60_0):
                    def iife220_(d_2617_dt__update__tmp_h2_):
                        def iife221_(_pat_let61_0):
                            def iife222_(d_2618_dt__update_hcf2_h1_):
                                def iife223_(_pat_let62_0):
                                    def iife224_(d_2619_dt__update_hcf1_h0_):
                                        return D0_DC1(d_2619_dt__update_hcf1_h0_, d_2618_dt__update_hcf2_h1_, (d_2617_dt__update__tmp_h2_).cf3, (d_2617_dt__update__tmp_h2_).cf4)
                                    return iife224_(_pat_let62_0)
                                return iife223_(pat_let_tv100_)
                            return iife222_(_pat_let61_0)
                        return iife221_(default__.fm36(default__.fm2(pat_let_tv96_), default__.fm2(pat_let_tv97_), pat_let_tv98_, pat_let_tv99_))
                    return iife220_(_pat_let60_0)
                nw380_[int(6)] = iife219_(d_2607_v136_)
                nw380_[int(7)] = d_2607_v136_
                nw380_[int(8)] = d_2607_v136_
                nw380_[int(9)] = d_2607_v136_
                def iife225_(_pat_let63_0):
                    def iife226_(d_2620_dt__update__tmp_h3_):
                        def iife227_(_pat_let64_0):
                            def iife228_(d_2621_dt__update_hcf4_h0_):
                                return D0_DC1((d_2620_dt__update__tmp_h3_).cf1, (d_2620_dt__update__tmp_h3_).cf2, (d_2620_dt__update__tmp_h3_).cf3, d_2621_dt__update_hcf4_h0_)
                            return iife228_(_pat_let64_0)
                        return iife227_(not(pat_let_tv101_))
                    return iife226_(_pat_let63_0)
                nw380_[int(10)] = (iife225_(d_2607_v136_) if d_2431_v6_ else D0_DC1((d_2604_v133_).f26, d_2600_v129_, len(d_2609_v138_), d_2431_v6_))
                nw380_[int(11)] = D0_DC1(len(default__.fm0(globalState)), d_2600_v129_, 774, d_2597_cf44_)
                nw380_[int(12)] = d_2607_v136_
                nw380_[int(13)] = d_2607_v136_
                nw380_[int(14)] = D0_DC1(len(d_2606_v135_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfcbcyxby")), 133, d_2431_v6_)
                nw380_[int(15)] = D0_DC1(d_2596_cf45_, d_2600_v129_, len(_dafny.Map({d_2597_cf44_: d_2431_v6_})), d_2431_v6_)
                nw380_[int(16)] = d_2607_v136_
                nw380_[int(17)] = d_2607_v136_
                nw380_[int(18)] = d_2607_v136_
                nw380_[int(19)] = d_2607_v136_
                nw380_[int(20)] = d_2607_v136_
                nw380_[int(21)] = default__.fm15(d_2597_cf44_, (d_2604_v133_).f26, _dafny.SeqWithoutIsStrInference([d_2431_v6_]), d_2431_v6_, globalState)
                nw380_[int(22)] = d_2607_v136_
                nw380_[int(23)] = (d_2604_v133_).fm5(d_2424_v0_, globalState)
                nw380_[int(24)] = D0_DC1(len(d_2600_v129_), d_2600_v129_, (0) - ((d_2604_v133_).f26), d_2431_v6_)
                nw380_[int(25)] = D0_DC1(len(d_2612_v141_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du"))).set(default__.safeIndex((d_2604_v133_).f26, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du")))), d_2608_v137_), d_2596_cf45_, d_2431_v6_)
                nw380_[int(26)] = default__.fm15(d_2431_v6_, (d_2604_v133_).f26, d_2598_cf43_, d_2431_v6_, globalState)
                d_2613_v142_ = nw380_
                index369_ = default__.safeIndex(214, (d_2613_v142_).length(0))
                (d_2613_v142_)[index369_] = D0_DC1(len(d_2600_v129_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")), d_2596_cf45_, d_2431_v6_)
                index370_ = default__.safeIndex(214, (d_2613_v142_).length(0))
                (d_2613_v142_)[index370_] = d_2607_v136_
                d_2622_v143_: _dafny.MultiSet
                d_2622_v143_ = _dafny.MultiSet([d_2600_v129_])
                (globalState).f8 = not((d_2622_v143_).isdisjoint((d_2622_v143_) | (d_2622_v143_)))
                d_2623_v144_: _dafny.Set
                d_2623_v144_ = _dafny.Set({(_dafny.MultiSet([d_2424_v0_, d_2424_v0_, d_2424_v0_])).isdisjoint(d_2602_v131_)})
                rhs388_ = (d_2560_v101_)[default__.safeIndex(((0) - (d_2424_v0_)) * ((d_2604_v133_).f26), len(d_2560_v101_))]
                rhs389_ = d_2602_v131_
                rhs390_ = _dafny.Map({(d_2623_v144_).isdisjoint(d_2623_v144_): d_2597_cf44_})
                rhs391_ = (d_2623_v144_).intersection((d_2623_v144_) | (d_2623_v144_))
                r0 = rhs388_
                d_2602_v131_ = rhs389_
                r1 = rhs390_
                d_2623_v144_ = rhs391_
                d_2556_v99_ = d_2556_v99_
            elif True:
                d_2624_v145_: _dafny.Array
                nw381_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_2624_v145_ = nw381_
                d_2625_v146_: _dafny.Map
                d_2625_v146_ = _dafny.Map({not(not (False) or (d_2431_v6_)): d_2624_v145_})
                d_2625_v146_ = (d_2625_v146_).set(d_2597_cf44_, d_2624_v145_)
                d_2596_cf45_ = d_2424_v0_
                (globalState).f8 = (d_2604_v133_).fm6(d_2431_v6_, (d_2604_v133_).f26, (d_2604_v133_).f26, globalState)
                index371_ = default__.safeIndex(47, (d_2556_v99_).length(0))
                (d_2556_v99_)[index371_] = 328
                index372_ = default__.safeIndex(47, (d_2556_v99_).length(0))
                (d_2556_v99_)[index372_] = (default__.safeDivisionInt((d_2604_v133_).f26, -509)) + ((d_2604_v133_).f26)
                d_2596_cf45_ = len(d_2603_v132_)
            d_2626_v147_: _dafny.Array
            nw382_ = _dafny.Array(False, 18)
            d_2626_v147_ = nw382_
            d_2627_v148_: _dafny.MultiSet
            d_2627_v148_ = _dafny.MultiSet([d_2626_v147_, d_2626_v147_, d_2626_v147_, d_2626_v147_])
            d_2628_v149_: C12
            nw383_ = C12()
            nw383_.ctor__((d_2627_v148_).intersection(d_2627_v148_), d_2424_v0_, d_2597_cf44_, d_2596_cf45_, default__.safeDivisionInt(d_2596_cf45_, (d_2604_v133_).f26))
            d_2628_v149_ = nw383_
        elif source38_.is_DC21:
            d_2629___mcc_h15_ = source38_.cf39
            d_2630_cf39_ = d_2629___mcc_h15_
            d_2631_v150_: C10
            nw384_ = C10()
            nw384_.ctor__(False)
            d_2631_v150_ = nw384_
            d_2632_v151_: _dafny.Set
            d_2632_v151_ = _dafny.Set({d_2631_v150_, d_2631_v150_})
            d_2633_v152_: _dafny.Map
            d_2633_v152_ = _dafny.Map({(d_2632_v151_) | (d_2632_v151_): d_2424_v0_})
            d_2633_v152_ = d_2633_v152_
            d_2634_v153_: _dafny.Map
            d_2634_v153_ = _dafny.Map({d_2424_v0_: d_2431_v6_})
            d_2635_v154_: _dafny.Seq
            d_2635_v154_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2424_v0_: not(d_2431_v6_)})), (0) - (d_2424_v0_)])
            d_2636_v155_: C0
            nw385_ = C0()
            nw385_.ctor__(d_2424_v0_, not((d_2424_v0_) == (len(d_2634_v153_))), (0) - (d_2424_v0_), (d_2635_v154_)[default__.safeIndex(853, len(d_2635_v154_))])
            d_2636_v155_ = nw385_
            r0 = d_2636_v155_.f29
            d_2637_v156_: _dafny.Array
            nw386_ = _dafny.Array(False, 8)
            d_2637_v156_ = nw386_
            def lambda126_(d_2638_v155_):
                def lambda127_(d_2639_i10_):
                    return d_2638_v155_.f29

                return lambda127_

            init69_ = lambda126_(d_2636_v155_)
            nw387_ = _dafny.Array(None, 11)
            for i0_69_ in range(nw387_.length(0)):
                nw387_[i0_69_] = init69_(i0_69_)
            d_2637_v156_ = nw387_
        elif True:
            d_2640___mcc_h16_ = source38_.cf46
            d_2641_cf46_ = d_2640___mcc_h16_
            d_2642_v157_: _dafny.Array
            def lambda128_(d_2643_v6_):
                def lambda129_(d_2644_i11_):
                    return d_2643_v6_

                return lambda129_

            init70_ = lambda128_(d_2431_v6_)
            nw388_ = _dafny.Array(None, 7)
            for i0_70_ in range(nw388_.length(0)):
                nw388_[i0_70_] = init70_(i0_70_)
            d_2642_v157_ = nw388_
            d_2645_v158_: _dafny.Array
            def lambda130_(d_2646_i12_):
                return D6_DC18(True)

            init71_ = lambda130_
            nw389_ = _dafny.Array(None, 14)
            for i0_71_ in range(nw389_.length(0)):
                nw389_[i0_71_] = init71_(i0_71_)
            d_2645_v158_ = nw389_
            d_2647_v159_: _dafny.MultiSet
            d_2647_v159_ = _dafny.MultiSet([d_2645_v158_, d_2645_v158_])
            d_2648_v160_: D11
            d_2648_v160_ = D11_DC30(d_2647_v159_)
            d_2649_v161_: _dafny.MultiSet
            d_2649_v161_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2648_v160_, d_2648_v160_, d_2648_v160_, d_2648_v160_, d_2648_v160_])])
            d_2650_v162_: _dafny.Seq
            d_2650_v162_ = _dafny.SeqWithoutIsStrInference([d_2648_v160_, d_2648_v160_])
            d_2651_v163_: _dafny.Map
            d_2651_v163_ = _dafny.Map({d_2642_v157_: (d_2649_v161_).set(d_2650_v162_, default__.abs(len(d_2560_v101_)))})
            r0 = (default__.fm1(d_2424_v0_, d_2431_v6_, d_2424_v0_, globalState) if (((d_2651_v163_)[d_2642_v157_] if (d_2642_v157_) in (d_2651_v163_) else d_2649_v161_)).issubset(d_2649_v161_) else d_2431_v6_)
            d_2652_v164_: D7
            d_2652_v164_ = D7_DC23(d_2559_v100_, d_2560_v101_, True, d_2424_v0_)
            d_2653_v165_: D7
            d_2653_v165_ = D7_DC24(d_2652_v164_)
            d_2654_v166_: D7
            d_2654_v166_ = D7_DC24(d_2652_v164_)
            d_2655_v167_: _dafny.Seq
            d_2655_v167_ = _dafny.SeqWithoutIsStrInference([d_2654_v166_])
            d_2656_v168_: _dafny.Map
            d_2656_v168_ = _dafny.Map({d_2655_v167_: d_2424_v0_})
            d_2657_v169_: _dafny.Map
            d_2657_v169_ = _dafny.Map({d_2656_v168_: d_2642_v157_})
            d_2642_v157_ = ((d_2657_v169_)[d_2656_v168_] if (d_2656_v168_) in (d_2657_v169_) else d_2642_v157_)
            d_2658_v170_: _dafny.Seq
            d_2658_v170_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cefp"))
            (globalState).f8 = (d_2658_v170_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))
            d_2659_v171_: _dafny.Seq
            d_2659_v171_ = _dafny.SeqWithoutIsStrInference([d_2424_v0_, d_2424_v0_])
            d_2660_v172_: _dafny.Map
            d_2660_v172_ = _dafny.Map({d_2659_v171_: d_2424_v0_})
            index373_ = default__.safeIndex(325, (d_2556_v99_).length(0))
            (d_2556_v99_)[index373_] = len((d_2660_v172_).set(d_2659_v171_, d_2424_v0_))
            d_2661_v174_: _dafny.Set
            def iife229_():
                coll99_ = _dafny.Set()
                compr_99_: int
                for compr_99_ in _dafny.IntegerRange(968, 401):
                    d_2662_v173_: int = compr_99_
                    if ((968) <= (d_2662_v173_)) and ((d_2662_v173_) < (401)):
                        coll99_ = coll99_.union(_dafny.Set([default__.safeDivisionInt(d_2662_v173_, d_2424_v0_)]))
                return _dafny.Set(coll99_)
            d_2661_v174_ = iife229_()
            
            index374_ = default__.safeIndex(806, (d_2642_v157_).length(0))
            (d_2642_v157_)[index374_] = d_2431_v6_
            d_2663_v175_: _dafny.Set
            d_2663_v175_ = _dafny.Set({-621, d_2424_v0_})
            index375_ = default__.safeIndex(325, (d_2556_v99_).length(0))
            index376_ = default__.safeIndex(806, (d_2642_v157_).length(0))
            rhs392_ = (default__.safeDivisionInt(d_2424_v0_, d_2424_v0_)) * (d_2424_v0_)
            rhs393_ = d_2663_v175_
            rhs394_ = d_2431_v6_
            rhs395_ = False
            rhs396_ = (d_2424_v0_) - (d_2424_v0_)
            lhs254_ = d_2556_v99_
            lhs255_ = default__.safeIndex(325, (d_2556_v99_).length(0))
            lhs256_ = d_2642_v157_
            lhs257_ = default__.safeIndex(806, (d_2642_v157_).length(0))
            lhs254_[lhs255_] = rhs392_
            d_2661_v174_ = rhs393_
            d_2431_v6_ = rhs394_
            lhs256_[lhs257_] = rhs395_
            r2 = rhs396_
        d_2664_v176_: _dafny.Seq
        d_2664_v176_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbbw"))
        d_2665_v177_: _dafny.Map
        d_2665_v177_ = _dafny.Map({d_2424_v0_: 855})
        d_2666_v178_: _dafny.Map
        d_2666_v178_ = _dafny.Map({_dafny.CodePoint('w'): len(d_2665_v177_)})
        d_2667_v179_: _dafny.Map
        d_2667_v179_ = _dafny.Map({d_2666_v178_: not(d_2431_v6_)})
        d_2668_v180_: C2
        nw390_ = C2()
        nw390_.ctor__(default__.safeDivisionInt(297, d_2424_v0_), d_2664_v176_, (0) - (len(d_2667_v179_)), d_2424_v0_, d_2431_v6_)
        d_2668_v180_ = nw390_
        d_2669_v181_: _dafny.Array
        def lambda131_(d_2670_v176_):
            def lambda132_(d_2671_i13_):
                return d_2670_v176_

            return lambda132_

        init72_ = lambda131_(d_2664_v176_)
        nw391_ = _dafny.Array(None, 25)
        for i0_72_ in range(nw391_.length(0)):
            nw391_[i0_72_] = init72_(i0_72_)
        d_2669_v181_ = nw391_
        d_2672_v182_: _dafny.Map
        d_2672_v182_ = _dafny.Map({d_2669_v181_: d_2431_v6_})
        d_2672_v182_ = (d_2672_v182_).set(d_2669_v181_, d_2431_v6_)
        d_2673_v183_: _dafny.Map
        d_2673_v183_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")): not(d_2431_v6_)})
        d_2674_v185_: _dafny.Map
        d_2674_v185_ = _dafny.Map({len(_dafny.Map({d_2424_v0_: (d_2668_v180_).f26})): d_2431_v6_})
        d_2675_v186_: str
        d_2675_v186_ = _dafny.CodePoint('f')
        d_2676_v187_: _dafny.Seq
        d_2676_v187_ = _dafny.SeqWithoutIsStrInference([default__.fm36((d_2668_v180_).f26, (d_2668_v180_).f26, d_2675_v186_, globalState), d_2664_v176_, d_2664_v176_])
        d_2677_v188_: _dafny.Set
        d_2677_v188_ = _dafny.Set({d_2424_v0_, d_2424_v0_})
        def iife230_():
            coll100_ = _dafny.Set()
            compr_100_: _dafny.Seq
            for compr_100_ in (d_2673_v183_).keys.Elements:
                d_2678_v184_: _dafny.Seq = compr_100_
                if (d_2678_v184_) in (d_2673_v183_):
                    coll100_ = coll100_.union(_dafny.Set([d_2678_v184_]))
            return _dafny.Set(coll100_)
        def iife231_():
            coll101_ = _dafny.Set()
            compr_101_: _dafny.Seq
            for compr_101_ in (d_2673_v183_).keys.Elements:
                d_2679_v184_: _dafny.Seq = compr_101_
                if (d_2679_v184_) in (d_2673_v183_):
                    coll101_ = coll101_.union(_dafny.Set([d_2679_v184_]))
            return _dafny.Set(coll101_)
        r0 = ((0) - (len((((d_2676_v187_)[default__.safeIndex((d_2668_v180_).f26, len(d_2676_v187_))] if default__.fm1(len(iife230_()
        ), d_2431_v6_, (0) - (len(d_2674_v185_)), globalState) else d_2664_v176_)).set(default__.safeIndex((d_2668_v180_).f26, len(((d_2676_v187_)[default__.safeIndex((d_2668_v180_).f26, len(d_2676_v187_))] if default__.fm1(len(iife231_()
        ), d_2431_v6_, (0) - (len(d_2674_v185_)), globalState) else d_2664_v176_))), d_2675_v186_)))) not in ((_dafny.Set({(d_2668_v180_).f26})) | (d_2677_v188_))
        d_2680_v189_: _dafny.Map
        d_2680_v189_ = _dafny.Map({d_2431_v6_: d_2431_v6_})
        r1 = (d_2680_v189_) | (d_2680_v189_)
        r2 = (d_2668_v180_).f26
        r3 = ((d_2560_v101_) + ((d_2560_v101_).set(default__.safeIndex(d_2424_v0_, len(d_2560_v101_)), False))) + (d_2560_v101_)
        return r0, r1, r2, r3


class C14(T0):
    def  __init__(self):
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C14"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    def ctor__(self, f11):
        (self).f11 = f11

    def fm5(self, p0, globalState):
        return D0_DC1((len(_dafny.Set({self.f11}))) + (len(_dafny.Map({self.f11: 158}))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxliobhom"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2681_i0_ in range(default__.abs(953))])), (len(_dafny.SeqWithoutIsStrInference([self.f11, self.f11, (D0_DC3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbesth"))), self.f11)).cf8, self.f11, False]))) + ((0) - (-157)), False)

    def fm6(self, p0, p1, p2, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference([False, True]))) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2682_i0_ in range(default__.abs(867))])))) > (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehqwd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "viqa")))))

    def fm11(self, p0, p1, p2, globalState):
        return ((_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebhb"))))])) | (_dafny.MultiSet([-492]))).isdisjoint(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-929]))]))

    def m5(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2683_v0_: int
        d_2683_v0_ = -41
        d_2684_v1_: _dafny.Set
        d_2684_v1_ = _dafny.Set({p0})
        d_2685_v2_: _dafny.Seq
        d_2685_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhd"))
        d_2686_v3_: _dafny.Map
        d_2686_v3_ = _dafny.Map({d_2683_v0_: d_2683_v0_})
        d_2687_v4_: _dafny.Array
        nw392_ = _dafny.Array(None, 28)
        nw392_[int(0)] = d_2683_v0_
        nw392_[int(1)] = d_2683_v0_
        nw392_[int(2)] = len(d_2684_v1_)
        nw392_[int(3)] = 844
        nw392_[int(4)] = 555
        nw392_[int(5)] = len((D1_DC4(p2)).cf9)
        nw392_[int(6)] = (0) - (d_2683_v0_)
        nw392_[int(7)] = len(d_2685_v2_)
        nw392_[int(8)] = d_2683_v0_
        nw392_[int(9)] = d_2683_v0_
        nw392_[int(10)] = d_2683_v0_
        nw392_[int(11)] = 734
        nw392_[int(12)] = d_2683_v0_
        nw392_[int(13)] = d_2683_v0_
        nw392_[int(14)] = d_2683_v0_
        nw392_[int(15)] = d_2683_v0_
        nw392_[int(16)] = (0) - (d_2683_v0_)
        nw392_[int(17)] = d_2683_v0_
        nw392_[int(18)] = len(_dafny.Map({d_2683_v0_: self.f11}))
        nw392_[int(19)] = ((d_2686_v3_)[d_2683_v0_] if (d_2683_v0_) in (d_2686_v3_) else d_2683_v0_)
        nw392_[int(20)] = 223
        nw392_[int(21)] = d_2683_v0_
        nw392_[int(22)] = (0) - (d_2683_v0_)
        nw392_[int(23)] = d_2683_v0_
        nw392_[int(24)] = d_2683_v0_
        nw392_[int(25)] = d_2683_v0_
        nw392_[int(26)] = d_2683_v0_
        nw392_[int(27)] = d_2683_v0_
        d_2687_v4_ = nw392_
        hi12_ = (D0_DC2((0) - (d_2683_v0_), d_2687_v4_)).cf5
        for d_2688_i0_ in range(d_2683_v0_, hi12_):
            r1 = self.f11
            d_2689_v5_: str
            d_2689_v5_ = _dafny.CodePoint('v')
            d_2690_v6_: _dafny.Map
            d_2690_v6_ = _dafny.Map({d_2686_v3_: d_2689_v5_})
            d_2691_v7_: _dafny.MultiSet
            d_2691_v7_ = _dafny.MultiSet([self.f11, p0])
            d_2692_v8_: _dafny.MultiSet
            d_2692_v8_ = _dafny.MultiSet([len(_dafny.Map({len(d_2690_v6_): (d_2691_v7_).cardinality})), d_2688_i0_])
            r1 = (default__.fm12(d_2683_v0_, d_2683_v0_, ((d_2692_v8_)[(d_2691_v7_).cardinality] if ((d_2691_v7_).cardinality) in (d_2692_v8_) else d_2683_v0_), p0, globalState)).issubset((_dafny.MultiSet(default__.fm0(globalState))).intersection(d_2692_v8_))
            d_2693_v9_: C1
            nw393_ = C1()
            nw393_.ctor__(len(d_2685_v2_))
            d_2693_v9_ = nw393_
            pat_let_tv102_ = d_2687_v4_
            def iife234_(_pat_let67_0):
                def iife235_(d_2694_dt__update__tmp_h0_):
                    def iife236_(_pat_let68_0):
                        def iife237_(d_2695_dt__update_hcf58_h0_):
                            return D12_DC33(d_2695_dt__update_hcf58_h0_)
                        return iife237_(_pat_let68_0)
                    return iife236_(d_2688_i0_)
                return iife235_(_pat_let67_0)
            def iife233_(_pat_let66_0):
                def iife238_(d_2696_dt__update__tmp_h1_):
                    def iife239_(_pat_let69_0):
                        def iife240_(d_2697_dt__update_hcf58_h1_):
                            return D12_DC33(d_2697_dt__update_hcf58_h1_)
                        return iife240_(_pat_let69_0)
                    return iife239_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv102_])))
                return iife238_(_pat_let66_0)
            def iife232_(_pat_let65_0):
                def iife241_(d_2698_dt__update__tmp_h2_):
                    def iife242_(_pat_let70_0):
                        def iife243_(d_2699_dt__update_hcf58_h2_):
                            return D12_DC33(d_2699_dt__update_hcf58_h2_)
                        return iife243_(_pat_let70_0)
                    return iife242_(226)
                return iife241_(_pat_let65_0)
            source39_ = iife232_(iife233_(iife234_(D12_DC33(d_2683_v0_))))
            if source39_.is_DC33:
                d_2700___mcc_h0_ = source39_.cf58
                d_2701_cf58_ = d_2700___mcc_h0_
                r1 = p0
                index377_ = default__.safeIndex(307, (d_2687_v4_).length(0))
                (d_2687_v4_)[index377_] = d_2688_i0_
                index378_ = default__.safeIndex(307, (d_2687_v4_).length(0))
                (d_2687_v4_)[index378_] = (0) - (default__.fm2(globalState))
                (globalState).f8 = p0
                (globalState).f8 = False
            elif source39_.is_DC34:
                d_2702___mcc_h1_ = source39_.cf59
                d_2703_cf59_ = d_2702___mcc_h1_
                d_2704_v10_: _dafny.Array
                nw394_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2704_v10_ = nw394_
                d_2705_v11_: D7
                d_2705_v11_ = D7_DC22(-520, d_2704_v10_)
                d_2706_v12_: _dafny.Seq
                d_2706_v12_ = _dafny.SeqWithoutIsStrInference([D7_DC22(548, d_2704_v10_), d_2705_v11_, d_2705_v11_, d_2705_v11_, d_2705_v11_])
                d_2706_v12_ = d_2706_v12_
                d_2707_v13_: _dafny.Map
                d_2707_v13_ = _dafny.Map({d_2692_v8_: self.f11})
                d_2708_v14_: _dafny.Seq
                d_2708_v14_ = _dafny.SeqWithoutIsStrInference([d_2693_v9_.f27])
                d_2709_v15_: _dafny.Map
                d_2709_v15_ = _dafny.Map({p0: _dafny.MultiSet(d_2708_v14_)})
                d_2707_v13_ = (d_2707_v13_).set(((d_2709_v15_)[(self).fm6(p0, d_2693_v9_.f27, 618, globalState)] if ((self).fm6(p0, d_2693_v9_.f27, 618, globalState)) in (d_2709_v15_) else d_2692_v8_), p0)
                (d_2693_v9_).m24(globalState)
                (d_2693_v9_).m24(globalState)
            elif source39_.is_DC35:
                d_2710___mcc_h2_ = source39_.cf60
                d_2711_cf60_ = d_2710___mcc_h2_
                d_2712_v16_: _dafny.Seq
                d_2712_v16_ = _dafny.SeqWithoutIsStrInference([self.f11, not(False), p0, self.f11, self.f11])
                d_2713_v17_: D6
                d_2713_v17_ = D6_DC19(p0, self.f11, p0, p0, d_2712_v16_)
                d_2714_v18_: _dafny.Map
                d_2714_v18_ = _dafny.Map({d_2689_v5_: (0) - (d_2688_i0_)})
                d_2715_v19_: _dafny.Map
                d_2715_v19_ = _dafny.Map({((d_2691_v7_)[(d_2713_v17_).cf34] if ((d_2713_v17_).cf34) in (d_2691_v7_) else len(d_2714_v18_)): p0})
                d_2712_v16_ = _dafny.SeqWithoutIsStrInference([((d_2715_v19_)[len(p1)] if (len(p1)) in (d_2715_v19_) else p0), p0, self.f11, p0])
                d_2687_v4_ = d_2687_v4_
                (d_2693_v9_).f27 = d_2693_v9_.f27
                d_2716_v20_: _dafny.Array
                def lambda133_(d_2717_cf60_):
                    def lambda134_(d_2718_i1_):
                        return (d_2717_cf60_) <= (790)

                    return lambda134_

                init73_ = lambda133_(d_2711_cf60_)
                nw395_ = _dafny.Array(None, 26)
                for i0_73_ in range(nw395_.length(0)):
                    nw395_[i0_73_] = init73_(i0_73_)
                d_2716_v20_ = nw395_
                index379_ = default__.safeIndex(242, (d_2716_v20_).length(0))
                (d_2716_v20_)[index379_] = p0
                index380_ = default__.safeIndex(242, (d_2716_v20_).length(0))
                (d_2716_v20_)[index380_] = p0
            elif True:
                d_2719___mcc_h3_ = source39_.cf57
                d_2720_cf57_ = d_2719___mcc_h3_
                d_2721_v21_: C11
                nw396_ = C11()
                nw396_.ctor__((len(d_2685_v2_) if self.f11 else (d_2692_v8_).cardinality), d_2693_v9_.f27)
                d_2721_v21_ = nw396_
                d_2722_v22_: _dafny.Map
                d_2722_v22_ = _dafny.Map({D1_DC5(d_2692_v8_, self.f11): _dafny.SeqWithoutIsStrInference([d_2689_v5_ for d_2723_i2_ in range(default__.abs(96))])})
                d_2724_v23_: _dafny.Array
                def lambda135_(d_2725_v5_):
                    def lambda136_(d_2726_i3_):
                        return (D9_DC26(d_2725_v5_)).cf48

                    return lambda136_

                init74_ = lambda135_(d_2689_v5_)
                nw397_ = _dafny.Array(None, 19)
                for i0_74_ in range(nw397_.length(0)):
                    nw397_[i0_74_] = init74_(i0_74_)
                d_2724_v23_ = nw397_
                index381_ = default__.safeIndex(547, (d_2724_v23_).length(0))
                (d_2724_v23_)[index381_] = _dafny.CodePoint('g')
                d_2727_v24_: _dafny.Array
                def lambda137_(d_2728_p0_):
                    def lambda138_(d_2729_i4_):
                        return d_2728_p0_

                    return lambda138_

                init75_ = lambda137_(p0)
                nw398_ = _dafny.Array(None, 1)
                for i0_75_ in range(nw398_.length(0)):
                    nw398_[i0_75_] = init75_(i0_75_)
                d_2727_v24_ = nw398_
                d_2730_v25_: D1
                d_2730_v25_ = D1_DC5(_dafny.MultiSet([d_2688_i0_, d_2693_v9_.f27]), not(p0))
                d_2731_v26_: _dafny.Set
                d_2731_v26_ = _dafny.Set({d_2684_v1_})
                index382_ = default__.safeIndex(547, (d_2724_v23_).length(0))
                rhs397_ = self.f11
                rhs398_ = _dafny.Map({d_2730_v25_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whowf"))) + (d_2685_v2_)})
                rhs399_ = ((d_2731_v26_) | (d_2731_v26_)).issubset(d_2731_v26_)
                rhs400_ = (d_2689_v5_ if self.f11 else d_2689_v5_)
                rhs401_ = d_2727_v24_
                lhs258_ = globalState
                lhs259_ = self
                lhs260_ = d_2724_v23_
                lhs261_ = default__.safeIndex(547, (d_2724_v23_).length(0))
                lhs258_.f8 = rhs397_
                d_2722_v22_ = rhs398_
                lhs259_.f11 = rhs399_
                lhs260_[lhs261_] = rhs400_
                d_2727_v24_ = rhs401_
                d_2685_v2_ = d_2685_v2_
                d_2732_v27_: _dafny.Map
                d_2732_v27_ = _dafny.Map({(d_2724_v23_)[default__.safeIndex(547, (d_2724_v23_).length(0))]: self.f11})
                d_2685_v2_ = ((d_2721_v21_).fm8(d_2693_v9_.f27, d_2732_v27_, p1, globalState)) + (_dafny.SeqWithoutIsStrInference([d_2689_v5_ for d_2733_i5_ in range(default__.abs(699))]))
        d_2734_v28_: _dafny.Array
        def lambda139_(d_2735_i6_):
            return _dafny.CodePoint('t')

        init76_ = lambda139_
        nw399_ = _dafny.Array(None, 18)
        for i0_76_ in range(nw399_.length(0)):
            nw399_[i0_76_] = init76_(i0_76_)
        d_2734_v28_ = nw399_
        d_2736_v29_: str
        d_2736_v29_ = _dafny.CodePoint('s')
        d_2737_v30_: _dafny.Map
        d_2737_v30_ = _dafny.Map({self.f11: d_2736_v29_})
        index383_ = default__.safeIndex(399, (d_2734_v28_).length(0))
        (d_2734_v28_)[index383_] = ((d_2737_v30_)[self.f11] if (self.f11) in (d_2737_v30_) else d_2736_v29_)
        index384_ = default__.safeIndex(399, (d_2734_v28_).length(0))
        (d_2734_v28_)[index384_] = d_2736_v29_
        d_2738_v32_: _dafny.Seq
        d_2738_v32_ = _dafny.SeqWithoutIsStrInference([d_2685_v2_, d_2685_v2_])
        d_2739_v33_: _dafny.Map
        d_2739_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wueuik")): False})
        d_2740_i7_: int
        d_2740_i7_ = 0
        with _dafny.label("12"):
            def iife245_():
                coll103_ = _dafny.Map()
                compr_103_: _dafny.Seq
                for compr_103_ in (d_2738_v32_).Elements:
                    d_2753_v31_: _dafny.Seq = compr_103_
                    if (d_2753_v31_) in (d_2738_v32_):
                        coll103_[d_2753_v31_] = p0
                return _dafny.Map(coll103_)
            while not ((iife245_()
            ) != (d_2739_v33_)) or (self.f11):
                with _dafny.c_label("12"):
                    if (d_2740_i7_) >= (100):
                        raise _dafny.Break("12")
                    d_2740_i7_ = (d_2740_i7_) + (1)
                    (globalState).f8 = self.f11
                    d_2741_v34_: C10
                    nw400_ = C10()
                    nw400_.ctor__(p0)
                    d_2741_v34_ = nw400_
                    d_2742_v35_: _dafny.MultiSet
                    d_2742_v35_ = _dafny.MultiSet([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))], (d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))]])
                    d_2743_v36_: _dafny.Map
                    d_2743_v36_ = _dafny.Map({d_2742_v35_: p0})
                    d_2744_v37_: D19
                    d_2744_v37_ = D19_DC50(d_2743_v36_)
                    d_2745_v39_: _dafny.Array
                    nw401_ = _dafny.Array(None, 6)
                    nw401_[int(0)] = (_dafny.Map({_dafny.MultiSet([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))]]): p0})).set(_dafny.MultiSet([_dafny.CodePoint('c')]), False)
                    nw401_[int(1)] = (d_2744_v37_).cf76
                    nw401_[int(2)] = (d_2744_v37_).cf76
                    nw401_[int(3)] = _dafny.Map({d_2742_v35_: default__.fm1(len(_dafny.SeqWithoutIsStrInference([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))] for d_2746_i8_ in range(default__.abs(-464))])), p0, (_dafny.MultiSet([not(self.f11), p0])).cardinality, globalState)})
                    nw401_[int(4)] = (_dafny.Map({_dafny.MultiSet(d_2685_v2_): p0})) | (d_2743_v36_)
                    def iife244_():
                        coll102_ = _dafny.Map()
                        compr_102_: _dafny.MultiSet
                        for compr_102_ in ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))]]), (d_2742_v35_).set(d_2736_v29_, default__.abs(d_2683_v0_))])).set(default__.safeIndex(d_2683_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))]]), (d_2742_v35_).set(d_2736_v29_, default__.abs(d_2683_v0_))]))), d_2742_v35_)).Elements:
                            d_2747_v38_: _dafny.MultiSet = compr_102_
                            if (d_2747_v38_) in ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))]]), (d_2742_v35_).set(d_2736_v29_, default__.abs(d_2683_v0_))])).set(default__.safeIndex(d_2683_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2734_v28_)[default__.safeIndex(399, (d_2734_v28_).length(0))]]), (d_2742_v35_).set(d_2736_v29_, default__.abs(d_2683_v0_))]))), d_2742_v35_)):
                                coll102_[d_2747_v38_] = False
                        return _dafny.Map(coll102_)
                    nw401_[int(5)] = (d_2743_v36_) | (iife244_()
                    )
                    d_2745_v39_ = nw401_
                    index385_ = default__.safeIndex(682, (d_2745_v39_).length(0))
                    (d_2745_v39_)[index385_] = (d_2743_v36_).set(d_2742_v35_, p0)
                    d_2748_v40_: D15
                    d_2748_v40_ = D15_DC41(self.f11)
                    index386_ = default__.safeIndex(682, (d_2745_v39_).length(0))
                    rhs402_ = self.f11
                    rhs403_ = d_2687_v4_
                    rhs404_ = d_2743_v36_
                    rhs405_ = (d_2748_v40_).cf65
                    rhs406_ = d_2685_v2_
                    lhs262_ = d_2745_v39_
                    lhs263_ = default__.safeIndex(682, (d_2745_v39_).length(0))
                    lhs264_ = globalState
                    r1 = rhs402_
                    d_2687_v4_ = rhs403_
                    lhs262_[lhs263_] = rhs404_
                    lhs264_.f8 = rhs405_
                    d_2685_v2_ = rhs406_
                    d_2749_v41_: _dafny.Array
                    nw402_ = _dafny.Array(False, 19)
                    d_2749_v41_ = nw402_
                    index387_ = default__.safeIndex(805, (d_2749_v41_).length(0))
                    (d_2749_v41_)[index387_] = p0
                    d_2750_v42_: _dafny.Map
                    d_2750_v42_ = _dafny.Map({p0: d_2683_v0_})
                    d_2751_v43_: _dafny.Seq
                    d_2751_v43_ = _dafny.SeqWithoutIsStrInference([d_2750_v42_])
                    d_2752_v44_: _dafny.MultiSet
                    d_2752_v44_ = _dafny.MultiSet([False])
                    index388_ = default__.safeIndex(805, (d_2749_v41_).length(0))
                    rhs407_ = ((d_2751_v43_) + (d_2751_v43_)) != ((d_2751_v43_) + (d_2751_v43_))
                    rhs408_ = (d_2683_v0_ if (d_2752_v44_).issubset(_dafny.MultiSet([p0])) else d_2683_v0_)
                    rhs409_ = self.f11
                    rhs410_ = d_2685_v2_
                    rhs411_ = ((d_2683_v0_) * (len(_dafny.SeqWithoutIsStrInference([d_2683_v0_])))) * (d_2683_v0_)
                    lhs265_ = d_2749_v41_
                    lhs266_ = default__.safeIndex(805, (d_2749_v41_).length(0))
                    lhs267_ = self
                    lhs265_[lhs266_] = rhs407_
                    r0 = rhs408_
                    lhs267_.f11 = rhs409_
                    d_2685_v2_ = rhs410_
                    d_2683_v0_ = rhs411_
                    pass
            pass
        d_2754_v45_: D4
        d_2754_v45_ = D4_DC12(self.f11, -585, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2755_i9_ in range(default__.abs(-289))])))
        d_2756_v46_: _dafny.Map
        d_2756_v46_ = _dafny.Map({self.f11: d_2683_v0_})
        d_2757_v48_: _dafny.Set
        def iife246_():
            coll104_ = _dafny.Set()
            compr_104_: int
            for compr_104_ in _dafny.IntegerRange(-622, 857):
                d_2758_v47_: int = compr_104_
                if ((-622) <= (d_2758_v47_)) and ((d_2758_v47_) < (857)):
                    coll104_ = coll104_.union(_dafny.Set([(d_2758_v47_) - (d_2683_v0_)]))
            return _dafny.Set(coll104_)
        d_2757_v48_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([default__.fm46(p0, self.f11, globalState), d_2754_v45_, D4_DC12(self.f11, ((d_2756_v46_)[p0] if (p0) in (d_2756_v46_) else len(iife246_()
)), d_2683_v0_), d_2754_v45_, d_2754_v45_])})
        d_2757_v48_ = (d_2757_v48_) - (d_2757_v48_)
        d_2759_v49_: _dafny.MultiSet
        d_2759_v49_ = _dafny.MultiSet([d_2687_v4_, d_2687_v4_])
        (self).f11 = not((d_2687_v4_) not in ((d_2759_v49_).set(d_2687_v4_, default__.abs(d_2683_v0_))))
        d_2760_v50_: D15
        d_2760_v50_ = D15_DC41(self.f11)
        d_2761_v51_: D15
        d_2761_v51_ = D15_DC42(d_2760_v50_)
        d_2762_v52_: _dafny.Array
        nw403_ = _dafny.Array(None, 17)
        nw403_[int(0)] = d_2761_v51_
        nw403_[int(1)] = d_2761_v51_
        nw403_[int(2)] = d_2761_v51_
        nw403_[int(3)] = d_2761_v51_
        nw403_[int(4)] = d_2761_v51_
        nw403_[int(5)] = D15_DC42(d_2760_v50_)
        nw403_[int(6)] = d_2761_v51_
        nw403_[int(7)] = d_2761_v51_
        nw403_[int(8)] = d_2761_v51_
        nw403_[int(9)] = d_2761_v51_
        nw403_[int(10)] = D15_DC42(d_2760_v50_)
        nw403_[int(11)] = default__.fm61(globalState)
        nw403_[int(12)] = d_2761_v51_
        nw403_[int(13)] = d_2761_v51_
        nw403_[int(14)] = d_2761_v51_
        nw403_[int(15)] = d_2761_v51_
        nw403_[int(16)] = d_2761_v51_
        d_2762_v52_ = nw403_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_2762_v52_).length(0)):
            d_2763_i10_: int = guard_loop_2_
            if (True) and (((0) <= (d_2763_i10_)) and ((d_2763_i10_) < ((d_2762_v52_).length(0)))):
                (d_2762_v52_)[(d_2763_i10_)] = d_2761_v51_
        r0 = d_2683_v0_
        r1 = p0
        return r0, r1


class C15(T0):
    def  __init__(self):
        self._f11: bool = False
        self.f16: bool = False
        self.f17: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C15"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    def ctor__(self, f16, f17, f11):
        (self).f16 = f16
        (self).f17 = f17
        (self).f11 = f11

    def fm5(self, p0, globalState):
        source40_ = D0_DC3(self.f17, self.f11)
        if source40_.is_DC1:
            d_2764___mcc_h0_ = source40_.cf1
            d_2765___mcc_h1_ = source40_.cf2
            d_2766___mcc_h2_ = source40_.cf3
            d_2767___mcc_h3_ = source40_.cf4
            d_2768_cf4_ = d_2767___mcc_h3_
            d_2769_cf3_ = d_2766___mcc_h2_
            d_2770_cf2_ = d_2765___mcc_h1_
            d_2771_cf1_ = d_2764___mcc_h0_
            return D0_DC1(622, d_2770_cf2_, len(_dafny.SeqWithoutIsStrInference([d_2771_cf1_, len(d_2770_cf2_)])), self.f11)
        elif source40_.is_DC2:
            d_2772___mcc_h4_ = source40_.cf5
            d_2773___mcc_h5_ = source40_.cf6
            d_2774_cf6_ = d_2773___mcc_h5_
            d_2775_cf5_ = d_2772___mcc_h4_
            return D0_DC1(d_2775_cf5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxbwjen")), d_2775_cf5_, not(self.f11))
        elif source40_.is_DC3:
            d_2776___mcc_h6_ = source40_.cf7
            d_2777___mcc_h7_ = source40_.cf8
            d_2778_cf8_ = d_2777___mcc_h7_
            d_2779_cf7_ = d_2776___mcc_h6_
            return D0_DC1((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2778_cf8_, self.f16]))).cardinality, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtyog")), len(_dafny.SeqWithoutIsStrInference([d_2778_cf8_])), not(self.f16))
        elif True:
            d_2780___mcc_h8_ = source40_.cf0
            d_2781_cf0_ = d_2780___mcc_h8_
            return D0_DC1(self.f17, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuxniq")), 715, self.f16)

    def fm6(self, p0, p1, p2, globalState):
        return not(self.f16)

    def fm10(self, p0, p1, p2, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('h')])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')]))))

    def m4(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        d_2782_v0_: _dafny.Array
        def lambda140_(d_2783_i0_):
            return self.f11

        init77_ = lambda140_
        nw404_ = _dafny.Array(None, 16)
        for i0_77_ in range(nw404_.length(0)):
            nw404_[i0_77_] = init77_(i0_77_)
        d_2782_v0_ = nw404_
        d_2784_v1_: _dafny.MultiSet
        d_2784_v1_ = _dafny.MultiSet([d_2782_v0_])
        d_2785_v2_: _dafny.Seq
        d_2785_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtcekccsk"))
        d_2786_v3_: T0
        nw405_ = C12()
        nw405_.ctor__(d_2784_v1_, len(_dafny.Set({self.f17, p0})), self.f16, len(_dafny.Set({len(d_2785_v2_)})), 372)
        d_2786_v3_ = nw405_
        d_2787_v4_: D6
        d_2787_v4_ = D6_DC17(d_2786_v3_.f11, (0) - (self.f17))
        pat_let_tv103_ = d_2786_v3_
        pat_let_tv104_ = d_2785_v2_
        pat_let_tv105_ = d_2785_v2_
        pat_let_tv106_ = d_2785_v2_
        pat_let_tv107_ = p1
        pat_let_tv108_ = p0
        pat_let_tv109_ = d_2786_v3_
        pat_let_tv110_ = p2
        pat_let_tv111_ = p1
        pat_let_tv112_ = p2
        def lambda141_(source41_):
            if source41_.is_DC17:
                d_2788___mcc_h0_ = source41_.cf30
                d_2789___mcc_h1_ = source41_.cf31
                d_2790_cf31_ = d_2789___mcc_h1_
                d_2791_cf30_ = d_2788___mcc_h0_
                if self.f16:
                    return pat_let_tv103_.f11
                elif True:
                    return d_2791_cf30_
            elif source41_.is_DC18:
                d_2792___mcc_h2_ = source41_.cf32
                d_2793_cf32_ = d_2792___mcc_h2_
                return (pat_let_tv104_) in (_dafny.MultiSet([pat_let_tv105_, pat_let_tv106_]))
            elif source41_.is_DC19:
                d_2794___mcc_h3_ = source41_.cf33
                d_2795___mcc_h4_ = source41_.cf34
                d_2796___mcc_h5_ = source41_.cf35
                d_2797___mcc_h6_ = source41_.cf36
                d_2798___mcc_h7_ = source41_.cf37
                d_2799_cf37_ = d_2798___mcc_h7_
                d_2800_cf36_ = d_2797___mcc_h6_
                d_2801_cf35_ = d_2796___mcc_h5_
                d_2802_cf34_ = d_2795___mcc_h4_
                d_2803_cf33_ = d_2794___mcc_h3_
                return (_dafny.MultiSet([pat_let_tv107_])).issubset(_dafny.MultiSet([len(_dafny.Map({pat_let_tv108_: self.f17}))]))
            elif source41_.is_DC16:
                d_2804___mcc_h8_ = source41_.cf29
                d_2805_cf29_ = d_2804___mcc_h8_
                return not(self.f16)
            elif True:
                d_2806___mcc_h9_ = source41_.cf38
                d_2807_cf38_ = d_2806___mcc_h9_
                return (pat_let_tv109_.f11) not in (_dafny.Map({self.f11: (pat_let_tv110_)[default__.safeIndex(pat_let_tv111_, len(pat_let_tv112_))]}))

        rhs412_ = self.f16
        rhs413_ = not(lambda141_(d_2787_v4_))
        rhs414_ = default__.fm2(globalState)
        rhs415_ = d_2786_v3_
        lhs268_ = self
        lhs269_ = self
        lhs268_.f11 = rhs412_
        lhs269_.f11 = rhs413_
        r1 = rhs414_
        d_2786_v3_ = rhs415_
        d_2808_i1_: int
        d_2808_i1_ = 0
        with _dafny.label("13"):
            while d_2786_v3_.f11:
                with _dafny.c_label("13"):
                    if (d_2808_i1_) >= (100):
                        raise _dafny.Break("13")
                    d_2808_i1_ = (d_2808_i1_) + (1)
                    index389_ = default__.safeIndex(529, (d_2782_v0_).length(0))
                    (d_2782_v0_)[index389_] = self.f16
                    index390_ = default__.safeIndex(529, (d_2782_v0_).length(0))
                    (d_2782_v0_)[index390_] = (default__.fm2(globalState)) == (p0)
                    d_2809_v5_: _dafny.Seq
                    d_2809_v5_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_2810_v6_: C3
                    nw406_ = C3()
                    nw406_.ctor__((d_2809_v5_)[default__.safeIndex(p1, len(d_2809_v5_))], p1)
                    d_2810_v6_ = nw406_
                    d_2811_v7_: _dafny.Set
                    d_2811_v7_ = _dafny.Set({self.f16})
                    d_2812_v8_: _dafny.Map
                    d_2812_v8_ = _dafny.Map({p1: (p2).set(default__.safeIndex(self.f17, len(p2)), d_2786_v3_.f11)})
                    d_2813_v9_: _dafny.MultiSet
                    d_2813_v9_ = _dafny.MultiSet([p2, ((d_2812_v8_)[p0] if (p0) in (d_2812_v8_) else p2)])
                    d_2814_v10_: str
                    d_2814_v10_ = _dafny.CodePoint('i')
                    d_2815_v11_: _dafny.Map
                    d_2815_v11_ = _dafny.Map({d_2786_v3_.f11: d_2814_v10_})
                    d_2816_v12_: _dafny.Array
                    nw407_ = _dafny.Array(None, 16)
                    nw407_[int(0)] = len(d_2811_v7_)
                    nw407_[int(1)] = p1
                    nw407_[int(2)] = p1
                    nw407_[int(3)] = p1
                    nw407_[int(4)] = p1
                    nw407_[int(5)] = self.f17
                    nw407_[int(6)] = (0) - (p1)
                    nw407_[int(7)] = 522
                    nw407_[int(8)] = p0
                    nw407_[int(9)] = p0
                    nw407_[int(10)] = ((d_2813_v9_) | (default__.fm62(self.f16, self.f16, self.f11, not(self.f16), globalState))).cardinality
                    nw407_[int(11)] = len(d_2815_v11_)
                    nw407_[int(12)] = p0
                    nw407_[int(13)] = p0
                    nw407_[int(14)] = (self.f17) + (self.f17)
                    nw407_[int(15)] = 205
                    d_2816_v12_ = nw407_
                    index391_ = default__.safeIndex(625, (d_2816_v12_).length(0))
                    (d_2816_v12_)[index391_] = self.f17
                    d_2817_v13_: _dafny.Map
                    d_2817_v13_ = _dafny.Map({d_2786_v3_.f11: p0})
                    d_2818_v14_: _dafny.Map
                    d_2818_v14_ = _dafny.Map({self.f16: d_2817_v13_})
                    d_2819_v15_: _dafny.Map
                    d_2819_v15_ = _dafny.Map({self.f17: d_2786_v3_.f11})
                    d_2820_v16_: _dafny.Map
                    d_2820_v16_ = _dafny.Map({((d_2818_v14_)[False] if (False) in (d_2818_v14_) else d_2817_v13_): d_2819_v15_})
                    index392_ = default__.safeIndex(625, (d_2816_v12_).length(0))
                    (d_2816_v12_)[index392_] = len(d_2820_v16_)
                    d_2821_v17_: C2
                    nw408_ = C2()
                    nw408_.ctor__(self.f17, d_2785_v2_, self.f17, self.f17, (d_2782_v0_)[default__.safeIndex(529, (d_2782_v0_).length(0))])
                    d_2821_v17_ = nw408_
                    d_2822_v18_: _dafny.Map
                    d_2822_v18_ = _dafny.Map({d_2786_v3_.f11: d_2821_v17_})
                    d_2823_v19_: _dafny.Array
                    nw409_ = _dafny.Array(None, 22)
                    nw409_[int(0)] = d_2821_v17_
                    nw409_[int(1)] = d_2821_v17_
                    nw409_[int(2)] = ((d_2822_v18_)[d_2786_v3_.f11] if (d_2786_v3_.f11) in (d_2822_v18_) else d_2821_v17_)
                    nw409_[int(3)] = d_2821_v17_
                    nw409_[int(4)] = d_2821_v17_
                    nw409_[int(5)] = d_2821_v17_
                    nw409_[int(6)] = d_2821_v17_
                    nw409_[int(7)] = d_2821_v17_
                    nw409_[int(8)] = d_2821_v17_
                    nw409_[int(9)] = d_2821_v17_
                    nw409_[int(10)] = d_2821_v17_
                    nw409_[int(11)] = d_2821_v17_
                    nw409_[int(12)] = d_2821_v17_
                    nw409_[int(13)] = d_2821_v17_
                    nw409_[int(14)] = d_2821_v17_
                    nw409_[int(15)] = d_2821_v17_
                    nw409_[int(16)] = d_2821_v17_
                    nw409_[int(17)] = d_2821_v17_
                    nw409_[int(18)] = d_2821_v17_
                    nw409_[int(19)] = d_2821_v17_
                    nw409_[int(20)] = d_2821_v17_
                    nw409_[int(21)] = d_2821_v17_
                    d_2823_v19_ = nw409_
                    index393_ = default__.safeIndex(124, (d_2823_v19_).length(0))
                    (d_2823_v19_)[index393_] = d_2821_v17_
                    index394_ = default__.safeIndex(124, (d_2823_v19_).length(0))
                    (d_2823_v19_)[index394_] = d_2821_v17_
                    pass
            pass
        d_2824_v20_: _dafny.MultiSet
        d_2824_v20_ = _dafny.MultiSet([54])
        d_2825_v21_: _dafny.Seq
        d_2825_v21_ = _dafny.SeqWithoutIsStrInference([d_2824_v20_])
        d_2826_v22_: _dafny.Map
        d_2826_v22_ = _dafny.Map({self.f16: self.f11})
        d_2827_v23_: _dafny.Seq
        d_2827_v23_ = _dafny.SeqWithoutIsStrInference([-518, p1])
        d_2828_v24_: _dafny.Seq
        d_2828_v24_ = _dafny.SeqWithoutIsStrInference([len(d_2826_v22_), (0) - (len(d_2827_v23_)), (0) - (-60)])
        d_2829_v25_: _dafny.Set
        d_2829_v25_ = _dafny.Set({self.f11, d_2786_v3_.f11})
        d_2830_v26_: _dafny.Seq
        d_2830_v26_ = _dafny.SeqWithoutIsStrInference([d_2829_v25_])
        d_2831_v27_: _dafny.Map
        d_2831_v27_ = _dafny.Map({self.f16: self.f17})
        d_2832_v28_: _dafny.Map
        d_2832_v28_ = _dafny.Map({(p2)[default__.safeIndex(p0, len(p2))]: (d_2825_v21_).set(default__.safeIndex(self.f17, len(d_2825_v21_)), _dafny.MultiSet(d_2828_v24_))})
        d_2833_v29_: _dafny.Map
        d_2833_v29_ = _dafny.Map({p1: d_2825_v21_})
        d_2834_v30_: _dafny.Array
        nw410_ = _dafny.Array(None, 24)
        nw410_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p1])])
        nw410_[int(1)] = d_2825_v21_
        nw410_[int(2)] = d_2825_v21_
        nw410_[int(3)] = (d_2825_v21_) + (d_2825_v21_)
        nw410_[int(4)] = d_2825_v21_
        nw410_[int(5)] = d_2825_v21_
        nw410_[int(6)] = (d_2825_v21_).set(default__.safeIndex(self.f17, len(d_2825_v21_)), _dafny.MultiSet(d_2828_v24_))
        nw410_[int(7)] = d_2825_v21_
        nw410_[int(8)] = d_2825_v21_
        nw410_[int(9)] = d_2825_v21_
        nw410_[int(10)] = (d_2825_v21_) + ((d_2825_v21_).set(default__.safeIndex(len((d_2830_v26_)[default__.safeIndex((0) - (len(d_2831_v27_)), len(d_2830_v26_))]), len(d_2825_v21_)), d_2824_v20_))
        nw410_[int(11)] = ((d_2832_v28_)[self.f16] if (self.f16) in (d_2832_v28_) else d_2825_v21_)
        nw410_[int(12)] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_2828_v24_) for d_2835_i2_ in range(default__.abs(462))])
        nw410_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_2824_v20_, (D1_DC5(_dafny.MultiSet([len(d_2785_v2_)]), False)).cf10])) + (d_2825_v21_)
        nw410_[int(14)] = d_2825_v21_
        nw410_[int(15)] = d_2825_v21_
        nw410_[int(16)] = (d_2825_v21_) + (_dafny.SeqWithoutIsStrInference([d_2824_v20_, d_2824_v20_]))
        nw410_[int(17)] = d_2825_v21_
        nw410_[int(18)] = d_2825_v21_
        nw410_[int(19)] = d_2825_v21_
        nw410_[int(20)] = (((d_2833_v29_)[p0] if (p0) in (d_2833_v29_) else d_2825_v21_)) + (d_2825_v21_)
        nw410_[int(21)] = default__.fm63(globalState)
        nw410_[int(22)] = d_2825_v21_
        nw410_[int(23)] = d_2825_v21_
        d_2834_v30_ = nw410_
        index395_ = default__.safeIndex(873, (d_2834_v30_).length(0))
        (d_2834_v30_)[index395_] = (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p0]), d_2824_v20_])) + (_dafny.SeqWithoutIsStrInference([d_2824_v20_ for d_2836_i3_ in range(default__.abs(989))]))
        d_2837_v31_: str
        d_2837_v31_ = _dafny.CodePoint('y')
        index396_ = default__.safeIndex(873, (d_2834_v30_).length(0))
        rhs416_ = _dafny.SeqWithoutIsStrInference([((D1_DC5(d_2824_v20_, self.f16)).cf10) - (_dafny.MultiSet(d_2828_v24_)) for d_2838_i4_ in range(default__.abs(-703))])
        rhs417_ = False
        rhs418_ = d_2837_v31_
        rhs419_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2839_i5_ in range(default__.abs(762))])) + (d_2785_v2_)
        lhs270_ = d_2834_v30_
        lhs271_ = default__.safeIndex(873, (d_2834_v30_).length(0))
        lhs272_ = globalState
        lhs270_[lhs271_] = rhs416_
        lhs272_.f8 = rhs417_
        d_2837_v31_ = rhs418_
        d_2785_v2_ = rhs419_
        d_2840_i6_: int
        d_2840_i6_ = 0
        with _dafny.label("14"):
            while self.f16:
                with _dafny.c_label("14"):
                    if (d_2840_i6_) >= (100):
                        raise _dafny.Break("14")
                    d_2840_i6_ = (d_2840_i6_) + (1)
                    if (d_2829_v25_).ispropersubset(d_2829_v25_):
                        d_2782_v0_ = d_2782_v0_
                        d_2841_v32_: C2
                        nw411_ = C2()
                        nw411_.ctor__(p1, (d_2785_v2_) + (d_2785_v2_), (0) - (p0), p0, d_2786_v3_.f11)
                        d_2841_v32_ = nw411_
                        r1 = (0) - ((0) - (default__.fm2(globalState)))
                        d_2842_v34_: _dafny.Map
                        d_2842_v34_ = _dafny.Map({self.f17: d_2786_v3_.f11})
                        d_2843_v35_: _dafny.Map
                        d_2843_v35_ = _dafny.Map({(d_2842_v34_).set(p1, d_2786_v3_.f11): d_2828_v24_})
                        d_2844_v36_: _dafny.Map
                        def iife247_():
                            coll105_ = _dafny.Map()
                            compr_105_: _dafny.Map
                            for compr_105_ in (d_2843_v35_).keys.Elements:
                                d_2845_v33_: _dafny.Map = compr_105_
                                if (d_2845_v33_) in (d_2843_v35_):
                                    coll105_[d_2845_v33_] = p2
                            return _dafny.Map(coll105_)
                        d_2844_v36_ = iife247_()
                        
                        d_2844_v36_ = d_2844_v36_
                        (self).f17 = p0
                    elif True:
                        d_2846_v37_: C1
                        nw412_ = C1()
                        nw412_.ctor__(self.f17)
                        d_2846_v37_ = nw412_
                        (self).f17 = (d_2846_v37_).fm34(globalState)
                        r1 = (self.f17) + (p0)
                        (self).f11 = d_2786_v3_.f11
                        (self).f17 = (len(d_2828_v24_)) * (p0)
                    d_2847_v38_: _dafny.Seq
                    d_2847_v38_ = _dafny.SeqWithoutIsStrInference([d_2785_v2_])
                    source42_ = default__.fm43(p0, (self.f16 if self.f16 else d_2786_v3_.f11), p0, (d_2847_v38_)[default__.safeIndex(82, len(d_2847_v38_))], globalState)
                    if source42_.is_DC17:
                        d_2848___mcc_h10_ = source42_.cf30
                        d_2849___mcc_h11_ = source42_.cf31
                        d_2850_cf31_ = d_2849___mcc_h11_
                        d_2851_cf30_ = d_2848___mcc_h10_
                        d_2852_v39_: _dafny.Seq
                        d_2852_v39_ = d_2828_v24_
                        d_2853_v40_: _dafny.Set
                        d_2853_v40_ = _dafny.Set({d_2852_v39_})
                        d_2854_v41_: _dafny.Map
                        d_2854_v41_ = _dafny.Map({self.f17: False})
                        d_2855_v42_: T2
                        nw413_ = C9()
                        nw413_.ctor__(d_2853_v40_, p1, self.f17, (p1) + (p0), (p1) <= (len(d_2854_v41_)), default__.fm36(p1, p0, d_2837_v31_, globalState))
                        d_2855_v42_ = nw413_
                        d_2855_v42_ = d_2855_v42_
                        d_2837_v31_ = d_2837_v31_
                        r1 = p1
                        d_2856_v43_: _dafny.Array
                        d_2857_v44_: bool
                        out43_: _dafny.Array
                        out44_: bool
                        out43_, out44_ = default__.m0(globalState)
                        d_2856_v43_ = out43_
                        d_2857_v44_ = out44_
                    elif source42_.is_DC18:
                        d_2858___mcc_h12_ = source42_.cf32
                        d_2859_cf32_ = d_2858___mcc_h12_
                        (globalState).f5 = (d_2828_v24_) + ((d_2828_v24_) + (_dafny.SeqWithoutIsStrInference([p1])))
                        d_2860_v45_: _dafny.MultiSet
                        d_2860_v45_ = _dafny.MultiSet([d_2786_v3_.f11, self.f11])
                        (self).f17 = (p1 if (131) <= (((d_2860_v45_)[d_2786_v3_.f11] if (d_2786_v3_.f11) in (d_2860_v45_) else p0)) else p1)
                        d_2861_v46_: _dafny.Array
                        nw414_ = _dafny.Array(None, 4)
                        nw414_[int(0)] = p0
                        nw414_[int(1)] = p1
                        nw414_[int(2)] = self.f17
                        nw414_[int(3)] = p1
                        d_2861_v46_ = nw414_
                        index397_ = default__.safeIndex(915, (d_2861_v46_).length(0))
                        (d_2861_v46_)[index397_] = self.f17
                        index398_ = default__.safeIndex(915, (d_2861_v46_).length(0))
                        (d_2861_v46_)[index398_] = (default__.safeDivisionInt(p1, p0)) + (p1)
                        d_2862_v47_: D7
                        d_2862_v47_ = D7_DC21((_dafny.MultiSet([False, not(d_2786_v3_.f11), self.f11])).set(False, default__.abs(self.f17)))
                        d_2863_v48_: D7
                        d_2863_v48_ = D7_DC24(d_2862_v47_)
                        index399_ = default__.safeIndex(915, (d_2861_v46_).length(0))
                        index400_ = default__.safeIndex(915, (d_2861_v46_).length(0))
                        rhs420_ = p0
                        rhs421_ = len(((p2) + (p2)) + ((p2) + (p2)))
                        rhs422_ = d_2863_v48_
                        rhs423_ = p0
                        rhs424_ = p0
                        lhs273_ = d_2861_v46_
                        lhs274_ = default__.safeIndex(915, (d_2861_v46_).length(0))
                        lhs275_ = self
                        lhs276_ = d_2861_v46_
                        lhs277_ = default__.safeIndex(915, (d_2861_v46_).length(0))
                        lhs273_[lhs274_] = rhs420_
                        lhs275_.f17 = rhs421_
                        d_2863_v48_ = rhs422_
                        r1 = rhs423_
                        lhs276_[lhs277_] = rhs424_
                    elif source42_.is_DC19:
                        d_2864___mcc_h13_ = source42_.cf33
                        d_2865___mcc_h14_ = source42_.cf34
                        d_2866___mcc_h15_ = source42_.cf35
                        d_2867___mcc_h16_ = source42_.cf36
                        d_2868___mcc_h17_ = source42_.cf37
                        d_2869_cf37_ = d_2868___mcc_h17_
                        d_2870_cf36_ = d_2867___mcc_h16_
                        d_2871_cf35_ = d_2866___mcc_h15_
                        d_2872_cf34_ = d_2865___mcc_h14_
                        d_2873_cf33_ = d_2864___mcc_h13_
                        d_2874_v49_: _dafny.Seq
                        d_2874_v49_ = _dafny.SeqWithoutIsStrInference([d_2782_v0_, d_2782_v0_, d_2782_v0_, d_2782_v0_, d_2782_v0_])
                        d_2875_v50_: D4
                        d_2875_v50_ = D4_DC12(d_2870_cf36_, len(d_2785_v2_), p0)
                        d_2876_v51_: _dafny.Array
                        nw415_ = _dafny.Array(int(0), 9)
                        d_2876_v51_ = nw415_
                        d_2877_v52_: _dafny.Array
                        nw416_ = _dafny.Array(None, 26)
                        nw416_[int(0)] = d_2876_v51_
                        nw416_[int(1)] = d_2876_v51_
                        nw416_[int(2)] = d_2876_v51_
                        nw416_[int(3)] = d_2876_v51_
                        nw416_[int(4)] = d_2876_v51_
                        nw416_[int(5)] = d_2876_v51_
                        nw416_[int(6)] = d_2876_v51_
                        nw416_[int(7)] = d_2876_v51_
                        nw416_[int(8)] = d_2876_v51_
                        nw416_[int(9)] = d_2876_v51_
                        nw416_[int(10)] = d_2876_v51_
                        nw416_[int(11)] = d_2876_v51_
                        nw416_[int(12)] = d_2876_v51_
                        nw416_[int(13)] = d_2876_v51_
                        nw416_[int(14)] = d_2876_v51_
                        nw416_[int(15)] = d_2876_v51_
                        nw416_[int(16)] = d_2876_v51_
                        nw416_[int(17)] = d_2876_v51_
                        nw416_[int(18)] = d_2876_v51_
                        nw416_[int(19)] = d_2876_v51_
                        nw416_[int(20)] = d_2876_v51_
                        nw416_[int(21)] = d_2876_v51_
                        nw416_[int(22)] = d_2876_v51_
                        nw416_[int(23)] = d_2876_v51_
                        nw416_[int(24)] = d_2876_v51_
                        nw416_[int(25)] = d_2876_v51_
                        d_2877_v52_ = nw416_
                        d_2878_v53_: D7
                        d_2878_v53_ = D7_DC22(self.f17, d_2877_v52_)
                        d_2879_v54_: _dafny.MultiSet
                        d_2879_v54_ = _dafny.MultiSet([d_2871_cf35_, d_2870_cf36_])
                        d_2880_v55_: D7
                        d_2880_v55_ = D7_DC21(d_2879_v54_)
                        d_2881_v56_: _dafny.Set
                        d_2881_v56_ = _dafny.Set({D7_DC21(_dafny.MultiSet([self.f11])), d_2880_v55_, d_2880_v55_})
                        d_2882_v57_: _dafny.Map
                        d_2882_v57_ = _dafny.Map({d_2881_v56_: d_2837_v31_})
                        d_2883_v58_: _dafny.Seq
                        d_2883_v58_ = d_2828_v24_
                        d_2884_v59_: _dafny.Set
                        d_2884_v59_ = _dafny.Set({d_2883_v58_, d_2883_v58_, d_2883_v58_, d_2883_v58_})
                        d_2885_v60_: C9
                        nw417_ = C9()
                        nw417_.ctor__(d_2884_v59_, p0, (d_2879_v54_).cardinality, p1, d_2786_v3_.f11, d_2785_v2_)
                        d_2885_v60_ = nw417_
                        d_2886_v61_: _dafny.Map
                        d_2886_v61_ = _dafny.Map({d_2885_v60_: d_2872_cf34_})
                        d_2887_v62_: _dafny.Set
                        d_2887_v62_ = _dafny.Set({len(d_2882_v57_), len(d_2886_v61_)})
                        pat_let_tv113_ = p1
                        pat_let_tv114_ = p0
                        pat_let_tv115_ = p1
                        d_2888_v63_: _dafny.Array
                        nw418_ = _dafny.Array(None, 14)
                        nw418_[int(0)] = d_2875_v50_
                        nw418_[int(1)] = d_2875_v50_
                        def iife249_(_pat_let72_0):
                            def iife250_(d_2889_dt__update__tmp_h0_):
                                def iife251_(_pat_let73_0):
                                    def iife252_(d_2890_dt__update_hcf22_h0_):
                                        return D4_DC12((d_2889_dt__update__tmp_h0_).cf21, d_2890_dt__update_hcf22_h0_, (d_2889_dt__update__tmp_h0_).cf23)
                                    return iife252_(_pat_let73_0)
                                return iife251_(self.f17)
                            return iife250_(_pat_let72_0)
                        def iife248_(_pat_let71_0):
                            def iife253_(d_2891_dt__update__tmp_h1_):
                                def iife254_(_pat_let74_0):
                                    def iife255_(d_2892_dt__update_hcf23_h0_):
                                        return D4_DC12((d_2891_dt__update__tmp_h1_).cf21, (d_2891_dt__update__tmp_h1_).cf22, d_2892_dt__update_hcf23_h0_)
                                    return iife255_(_pat_let74_0)
                                return iife254_((_dafny.MultiSet([pat_let_tv113_, pat_let_tv114_])).cardinality)
                            return iife253_(_pat_let71_0)
                        nw418_[int(2)] = iife248_(iife249_(d_2875_v50_))
                        nw418_[int(3)] = (d_2875_v50_ if d_2870_cf36_ else d_2875_v50_)
                        nw418_[int(4)] = D4_DC12(self.f16, self.f17, self.f17)
                        def iife256_(_pat_let75_0):
                            def iife257_(d_2893_dt__update__tmp_h2_):
                                def iife258_(_pat_let76_0):
                                    def iife259_(d_2894_dt__update_hcf23_h1_):
                                        return D4_DC12((d_2893_dt__update__tmp_h2_).cf21, (d_2893_dt__update__tmp_h2_).cf22, d_2894_dt__update_hcf23_h1_)
                                    return iife259_(_pat_let76_0)
                                return iife258_(pat_let_tv115_)
                            return iife257_(_pat_let75_0)
                        nw418_[int(5)] = iife256_(d_2875_v50_)
                        nw418_[int(6)] = d_2875_v50_
                        nw418_[int(7)] = D4_DC12(d_2786_v3_.f11, p0, p0)
                        nw418_[int(8)] = D4_DC12(self.f11, (D9_DC27(d_2786_v3_.f11, d_2878_v53_, p1, d_2873_cf33_)).cf51, len(d_2887_v62_))
                        nw418_[int(9)] = d_2875_v50_
                        nw418_[int(10)] = d_2875_v50_
                        nw418_[int(11)] = d_2875_v50_
                        nw418_[int(12)] = d_2875_v50_
                        nw418_[int(13)] = d_2875_v50_
                        d_2888_v63_ = nw418_
                        d_2895_v64_: _dafny.Map
                        d_2895_v64_ = _dafny.Map({d_2885_v60_.f22: d_2870_cf36_})
                        d_2896_v65_: _dafny.Map
                        d_2896_v65_ = _dafny.Map({d_2895_v64_: d_2872_cf34_})
                        d_2897_v66_: D20
                        d_2897_v66_ = D20_DC52(d_2896_v65_)
                        rhs425_ = len(((d_2897_v66_).cf79) | (d_2896_v65_))
                        rhs426_ = (_dafny.SeqWithoutIsStrInference([d_2782_v0_])) + (_dafny.SeqWithoutIsStrInference([d_2782_v0_, d_2782_v0_, d_2782_v0_, d_2782_v0_, d_2782_v0_]))
                        rhs427_ = d_2888_v63_
                        r1 = rhs425_
                        d_2874_v49_ = rhs426_
                        d_2888_v63_ = rhs427_
                        d_2898_v67_: C5
                        nw419_ = C5()
                        nw419_.ctor__(d_2885_v60_.f22)
                        d_2898_v67_ = nw419_
                        index401_ = default__.safeIndex(49, (d_2782_v0_).length(0))
                        (d_2782_v0_)[index401_] = d_2872_cf34_
                        d_2899_v68_: _dafny.Array
                        def lambda142_(d_2900_p1_, d_2901_v31_):
                            def lambda143_(d_2902_i7_):
                                return _dafny.SeqWithoutIsStrInference([d_2900_p1_, len(_dafny.SeqWithoutIsStrInference([d_2901_v31_ for d_2903_i8_ in range(default__.abs(868))]))])

                            return lambda143_

                        init78_ = lambda142_(p1, d_2837_v31_)
                        nw420_ = _dafny.Array(None, 9)
                        for i0_78_ in range(nw420_.length(0)):
                            nw420_[i0_78_] = init78_(i0_78_)
                        d_2899_v68_ = nw420_
                        d_2904_v69_: _dafny.Map
                        d_2904_v69_ = _dafny.Map({d_2885_v60_.f22: p2})
                        index402_ = default__.safeIndex(49, (d_2782_v0_).length(0))
                        rhs428_ = ((p2) + (d_2869_cf37_)) < (((d_2904_v69_)[self.f17] if (self.f17) in (d_2904_v69_) else p2))
                        rhs429_ = (d_2869_cf37_)[default__.safeIndex(p0, len(d_2869_cf37_))]
                        rhs430_ = d_2899_v68_
                        rhs431_ = not(d_2870_cf36_)
                        rhs432_ = d_2885_v60_.f22
                        lhs278_ = d_2786_v3_
                        lhs279_ = d_2782_v0_
                        lhs280_ = default__.safeIndex(49, (d_2782_v0_).length(0))
                        lhs281_ = self
                        lhs282_ = self
                        lhs278_.f11 = rhs428_
                        lhs279_[lhs280_] = rhs429_
                        d_2899_v68_ = rhs430_
                        lhs281_.f11 = rhs431_
                        lhs282_.f17 = rhs432_
                        d_2905_v70_: _dafny.Map
                        d_2905_v70_ = _dafny.Map({p1: _dafny.Set({d_2786_v3_.f11})})
                        d_2905_v70_ = _dafny.Map({p1: (d_2829_v25_).intersection(_dafny.Set({self.f11}))})
                    elif source42_.is_DC16:
                        d_2906___mcc_h18_ = source42_.cf29
                        d_2907_cf29_ = d_2906___mcc_h18_
                        d_2785_v2_ = d_2785_v2_
                        (d_2786_v3_).f11 = self.f16
                        d_2908_v71_: D1
                        d_2908_v71_ = D1_DC5(d_2824_v20_, d_2786_v3_.f11)
                        d_2909_v72_: D1
                        d_2909_v72_ = D1_DC6(d_2908_v71_)
                        d_2910_v73_: _dafny.Array
                        nw421_ = _dafny.Array(int(0), 28)
                        d_2910_v73_ = nw421_
                        d_2911_v74_: _dafny.Map
                        d_2911_v74_ = _dafny.Map({p1: p1})
                        d_2912_v75_: _dafny.Map
                        d_2912_v75_ = _dafny.Map({d_2837_v31_: default__.fm50(globalState)})
                        d_2913_v76_: _dafny.Set
                        d_2913_v76_ = _dafny.Set({p0, len(d_2831_v27_), (d_2824_v20_).cardinality, self.f17})
                        index403_ = default__.safeIndex(654, (d_2910_v73_).length(0))
                        (d_2910_v73_)[index403_] = (((d_2911_v74_)[self.f17] if (self.f17) in (d_2911_v74_) else len(((d_2912_v75_)[d_2837_v31_] if (d_2837_v31_) in (d_2912_v75_) else d_2913_v76_)))) + (self.f17)
                        d_2914_v77_: D21
                        d_2914_v77_ = D21_DC55(d_2784_v1_)
                        d_2915_v78_: C12
                        nw422_ = C12()
                        nw422_.ctor__((d_2914_v77_).cf81, self.f17, False, p0, -840)
                        d_2915_v78_ = nw422_
                        d_2916_v79_: _dafny.Map
                        d_2916_v79_ = _dafny.Map({self.f11: _dafny.Map({p0: (d_2915_v78_).f19})})
                        d_2917_v80_: _dafny.Array
                        nw423_ = _dafny.Array(None, 3)
                        nw423_[int(0)] = ((d_2916_v79_)[True] if (True) in (d_2916_v79_) else (d_2911_v74_).set((d_2915_v78_).f19, default__.fm2(globalState)))
                        nw423_[int(1)] = _dafny.Map({self.f17: 352})
                        nw423_[int(2)] = d_2911_v74_
                        d_2917_v80_ = nw423_
                        index404_ = default__.safeIndex(216, (d_2917_v80_).length(0))
                        (d_2917_v80_)[index404_] = _dafny.Map({len(d_2785_v2_): (d_2915_v78_).f19})
                        d_2918_v81_: C12
                        d_2918_v81_ = d_2915_v78_
                        index405_ = default__.safeIndex(654, (d_2910_v73_).length(0))
                        index406_ = default__.safeIndex(216, (d_2917_v80_).length(0))
                        rhs433_ = d_2909_v72_
                        rhs434_ = p0
                        rhs435_ = (d_2918_v81_)
                        rhs436_ = ((d_2911_v74_) | (default__.fm38(d_2837_v31_, d_2837_v31_, globalState))) | ((d_2911_v74_).set(p0, p1))
                        lhs283_ = d_2910_v73_
                        lhs284_ = default__.safeIndex(654, (d_2910_v73_).length(0))
                        lhs285_ = d_2917_v80_
                        lhs286_ = default__.safeIndex(216, (d_2917_v80_).length(0))
                        d_2909_v72_ = rhs433_
                        lhs283_[lhs284_] = rhs434_
                        d_2915_v78_ = rhs435_
                        lhs285_[lhs286_] = rhs436_
                        index407_ = default__.safeIndex(654, (d_2910_v73_).length(0))
                        (d_2910_v73_)[index407_] = (0) - (((d_2824_v20_) - ((d_2824_v20_) - (d_2824_v20_))).cardinality)
                    elif True:
                        d_2919___mcc_h19_ = source42_.cf38
                        d_2920_cf38_ = d_2919___mcc_h19_
                        index408_ = default__.safeIndex(392, (d_2782_v0_).length(0))
                        (d_2782_v0_)[index408_] = d_2786_v3_.f11
                        index409_ = default__.safeIndex(392, (d_2782_v0_).length(0))
                        (d_2782_v0_)[index409_] = d_2786_v3_.f11
                        d_2921_v82_: _dafny.Array
                        nw424_ = _dafny.Array(None, 5)
                        nw424_[int(0)] = (722 if self.f16 else self.f17)
                        nw424_[int(1)] = (p0) - (p1)
                        nw424_[int(2)] = p1
                        nw424_[int(3)] = p1
                        nw424_[int(4)] = 805
                        d_2921_v82_ = nw424_
                        index410_ = default__.safeIndex(759, (d_2921_v82_).length(0))
                        (d_2921_v82_)[index410_] = p0
                        index411_ = default__.safeIndex(759, (d_2921_v82_).length(0))
                        (d_2921_v82_)[index411_] = p1
                        (globalState).f8 = ((p2).set(default__.safeIndex(len(d_2785_v2_), len(p2)), self.f11)) < (default__.fm39((p2)[default__.safeIndex(self.f17, len(p2))], self.f11, globalState))
                        index412_ = default__.safeIndex(759, (d_2921_v82_).length(0))
                        (d_2921_v82_)[index412_] = (((0) - ((d_2921_v82_)[default__.safeIndex(759, (d_2921_v82_).length(0))])) - (p1)) * ((d_2921_v82_)[default__.safeIndex(759, (d_2921_v82_).length(0))])
                    if not (d_2786_v3_.f11) or (d_2786_v3_.f11):
                        d_2922_v83_: _dafny.Set
                        d_2922_v83_ = _dafny.Set({549})
                        d_2923_v84_: _dafny.Set
                        d_2923_v84_ = _dafny.Set({d_2922_v83_, d_2922_v83_, _dafny.Set({len(p2)}), d_2922_v83_, default__.fm50(globalState)})
                        def iife260_():
                            coll106_ = _dafny.Set()
                            compr_106_: int
                            for compr_106_ in _dafny.IntegerRange(-143, 279):
                                d_2924_v85_: int = compr_106_
                                if ((-143) <= (d_2924_v85_)) and ((d_2924_v85_) < (279)):
                                    coll106_ = coll106_.union(_dafny.Set([(d_2924_v85_) * (586)]))
                            return _dafny.Set(coll106_)
                        (globalState).f8 = ((d_2923_v84_).intersection(_dafny.Set({iife260_()
                        }))).issubset(d_2923_v84_)
                        d_2925_v86_: C0
                        nw425_ = C0()
                        nw425_.ctor__(p1, d_2786_v3_.f11, p1, p0)
                        d_2925_v86_ = nw425_
                        index413_ = default__.safeIndex(358, (d_2782_v0_).length(0))
                        (d_2782_v0_)[index413_] = not(not (self.f16) or (d_2786_v3_.f11))
                        index414_ = default__.safeIndex(358, (d_2782_v0_).length(0))
                        (d_2782_v0_)[index414_] = self.f16
                        d_2926_v87_: _dafny.Map
                        d_2926_v87_ = _dafny.Map({(0) - ((d_2925_v86_).f28): self.f17})
                        rhs437_ = True
                        rhs438_ = d_2926_v87_
                        rhs439_ = self.f17
                        rhs440_ = ((d_2785_v2_) + (d_2785_v2_)) + (d_2785_v2_)
                        lhs287_ = d_2925_v86_
                        lhs287_.f29 = rhs437_
                        d_2926_v87_ = rhs438_
                        r1 = rhs439_
                        d_2785_v2_ = rhs440_
                        d_2927_v88_: _dafny.MultiSet
                        d_2927_v88_ = _dafny.MultiSet([d_2826_v22_])
                        d_2928_v89_: _dafny.MultiSet
                        d_2928_v89_ = (d_2927_v88_) - (d_2927_v88_)
                        d_2928_v89_ = d_2928_v89_
                    elif True:
                        d_2837_v31_ = d_2837_v31_
                        d_2929_v90_: _dafny.Array
                        nw426_ = _dafny.Array(_dafny.Array(None, 0), 15)
                        d_2929_v90_ = nw426_
                        d_2930_v91_: D7
                        d_2930_v91_ = D7_DC23(d_2929_v90_, p2, self.f16, len(d_2827_v23_))
                        d_2931_v92_: _dafny.Array
                        nw427_ = _dafny.Array(None, 12)
                        nw427_[int(0)] = default__.safeModuloInt(self.f17, self.f17)
                        nw427_[int(1)] = p1
                        nw427_[int(2)] = p1
                        nw427_[int(3)] = self.f17
                        nw427_[int(4)] = default__.safeDivisionInt((0) - (p0), p0)
                        nw427_[int(5)] = (len(p2)) * ((d_2930_v91_).cf45)
                        nw427_[int(6)] = p0
                        nw427_[int(7)] = p1
                        nw427_[int(8)] = p1
                        nw427_[int(9)] = len(d_2785_v2_)
                        nw427_[int(10)] = p0
                        nw427_[int(11)] = (0) - (default__.safeModuloInt(p0, self.f17))
                        d_2931_v92_ = nw427_
                        index415_ = default__.safeIndex(790, (d_2931_v92_).length(0))
                        (d_2931_v92_)[index415_] = len(d_2785_v2_)
                        d_2932_v93_: C6
                        nw428_ = C6()
                        nw428_.ctor__()
                        d_2932_v93_ = nw428_
                        d_2933_v94_: _dafny.Map
                        d_2933_v94_ = _dafny.Map({d_2932_v93_: _dafny.Map({d_2786_v3_.f11: p1})})
                        index416_ = default__.safeIndex(790, (d_2931_v92_).length(0))
                        rhs441_ = (d_2831_v27_) | (((d_2933_v94_)[d_2932_v93_] if (d_2932_v93_) in (d_2933_v94_) else d_2831_v27_))
                        rhs442_ = d_2782_v0_
                        rhs443_ = (p0) + (p0)
                        lhs288_ = d_2931_v92_
                        lhs289_ = default__.safeIndex(790, (d_2931_v92_).length(0))
                        d_2831_v27_ = rhs441_
                        d_2782_v0_ = rhs442_
                        lhs288_[lhs289_] = rhs443_
                        r1 = (0) - ((d_2931_v92_)[default__.safeIndex(790, (d_2931_v92_).length(0))])
                        d_2934_v95_: _dafny.MultiSet
                        d_2934_v95_ = _dafny.MultiSet([d_2786_v3_.f11, d_2786_v3_.f11, self.f16])
                        d_2935_v96_: _dafny.Set
                        d_2935_v96_ = _dafny.Set({((d_2931_v92_)[default__.safeIndex(790, (d_2931_v92_).length(0))]) - ((d_2931_v92_)[default__.safeIndex(790, (d_2931_v92_).length(0))]), p1, ((d_2831_v27_)[(self).fm6(self.f16, (d_2931_v92_)[default__.safeIndex(790, (d_2931_v92_).length(0))], self.f17, globalState)] if ((self).fm6(self.f16, (d_2931_v92_)[default__.safeIndex(790, (d_2931_v92_).length(0))], self.f17, globalState)) in (d_2831_v27_) else (d_2934_v95_).cardinality)})
                        index417_ = default__.safeIndex(790, (d_2931_v92_).length(0))
                        rhs444_ = not(d_2786_v3_.f11)
                        rhs445_ = (0) - (self.f17)
                        rhs446_ = (d_2935_v96_) | (d_2935_v96_)
                        lhs290_ = d_2786_v3_
                        lhs291_ = d_2931_v92_
                        lhs292_ = default__.safeIndex(790, (d_2931_v92_).length(0))
                        lhs290_.f11 = rhs444_
                        lhs291_[lhs292_] = rhs445_
                        d_2935_v96_ = rhs446_
                        d_2936_v97_: _dafny.Map
                        d_2936_v97_ = _dafny.Map({d_2837_v31_: (d_2829_v25_) | (d_2829_v25_)})
                        d_2936_v97_ = (d_2936_v97_).set(d_2837_v31_, d_2829_v25_)
                    d_2937_v98_: _dafny.Map
                    d_2937_v98_ = _dafny.Map({(self).fm6(self.f11, p1, p0, globalState): p2})
                    d_2938_v99_: _dafny.Array
                    nw429_ = _dafny.Array(int(0), 5)
                    d_2938_v99_ = nw429_
                    d_2939_v100_: _dafny.Map
                    d_2939_v100_ = _dafny.Map({d_2938_v99_: not(self.f16)})
                    d_2940_v101_: _dafny.Map
                    d_2940_v101_ = _dafny.Map({len(((d_2937_v98_)[self.f11] if (self.f11) in (d_2937_v98_) else p2)): d_2939_v100_})
                    d_2940_v101_ = (d_2940_v101_).set(p0, _dafny.Map({d_2938_v99_: d_2786_v3_.f11}))
                    pass
            pass
        d_2941_v105_: _dafny.Array
        def lambda144_(d_2942_v27_, d_2943_p0_, d_2944_v3_):
            def lambda145_(d_2945_i9_):
                def iife261_():
                    coll107_ = _dafny.Map()
                    compr_107_: int
                    for compr_107_ in _dafny.IntegerRange(5, -382):
                        d_2946_v102_: int = compr_107_
                        if ((5) <= (d_2946_v102_)) and ((d_2946_v102_) < (-382)):
                            coll107_[(d_2946_v102_) - (94)] = len(_dafny.Map({True: D3_DC8(d_2942_v27_)}))
                    return _dafny.Map(coll107_)
                def iife262_():
                    def iife264_():
                        coll110_ = _dafny.Map()
                        compr_110_: int
                        for compr_110_ in _dafny.IntegerRange(897, 68):
                            d_2947_v104_: int = compr_110_
                            if ((897) <= (d_2947_v104_)) and ((d_2947_v104_) < (68)):
                                coll110_[default__.safeModuloInt(d_2947_v104_, d_2943_p0_)] = d_2944_v3_.f11
                        return _dafny.Map(coll110_)
                    coll108_ = _dafny.Map()
                    def iife263_():
                        coll109_ = _dafny.Map()
                        compr_109_: int
                        for compr_109_ in _dafny.IntegerRange(897, 68):
                            d_2947_v104_: int = compr_109_
                            if ((897) <= (d_2947_v104_)) and ((d_2947_v104_) < (68)):
                                coll109_[default__.safeModuloInt(d_2947_v104_, d_2943_p0_)] = d_2944_v3_.f11
                        return _dafny.Map(coll109_)
                    compr_108_: int
                    for compr_108_ in (iife263_()
                    ).keys.Elements:
                        d_2948_v103_: int = compr_108_
                        if (d_2948_v103_) in (iife264_()
                        ):
                            coll108_[default__.safeDivisionInt(d_2948_v103_, self.f17)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfj")))
                    return _dafny.Map(coll108_)
                return (iife261_()
                ) | (iife262_()
                )

            return lambda145_

        init79_ = lambda144_(d_2831_v27_, p0, d_2786_v3_)
        nw430_ = _dafny.Array(None, 25)
        for i0_79_ in range(nw430_.length(0)):
            nw430_[i0_79_] = init79_(i0_79_)
        d_2941_v105_ = nw430_
        d_2949_v106_: _dafny.Array
        nw431_ = _dafny.Array(int(0), 23)
        d_2949_v106_ = nw431_
        d_2950_v107_: _dafny.Map
        d_2950_v107_ = _dafny.Map({d_2949_v106_: p1})
        pat_let_tv116_ = p1
        pat_let_tv117_ = d_2786_v3_
        pat_let_tv118_ = p0
        def lambda146_(source43_):
            if source43_.is_DC41:
                d_2953___mcc_h20_ = source43_.cf65
                d_2954_cf65_ = d_2953___mcc_h20_
                return pat_let_tv116_
            elif source43_.is_DC40:
                d_2955___mcc_h21_ = source43_.cf64
                d_2956_cf64_ = d_2955___mcc_h21_
                return (_dafny.MultiSet([pat_let_tv117_.f11])).cardinality
            elif True:
                d_2957___mcc_h22_ = source43_.cf66
                d_2958_cf66_ = d_2957___mcc_h22_
                return pat_let_tv118_

        rhs447_ = d_2941_v105_
        rhs448_ = (((0) - (p1) if d_2786_v3_.f11 else self.f17) if self.f16 else p0)
        rhs449_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nosncc"))) + (_dafny.SeqWithoutIsStrInference([d_2837_v31_ for d_2951_i10_ in range(default__.abs(548))]))).set(default__.safeIndex(len(d_2950_v107_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nosncc"))) + (_dafny.SeqWithoutIsStrInference([d_2837_v31_ for d_2952_i10_ in range(default__.abs(548))])))), d_2837_v31_)
        rhs450_ = (0) - (lambda146_(D15_DC41(self.f11)))
        lhs293_ = self
        d_2941_v105_ = rhs447_
        lhs293_.f17 = rhs448_
        d_2785_v2_ = rhs449_
        r1 = rhs450_
        d_2959_v108_: _dafny.Array
        nw432_ = _dafny.Array(_dafny.CodePoint('D'), 17)
        d_2959_v108_ = nw432_
        index418_ = default__.safeIndex(56, (d_2959_v108_).length(0))
        (d_2959_v108_)[index418_] = default__.fm25(self.f17, self.f17, globalState)
        index419_ = default__.safeIndex(56, (d_2959_v108_).length(0))
        rhs451_ = default__.fm2(globalState)
        rhs452_ = d_2782_v0_
        rhs453_ = d_2837_v31_
        rhs454_ = d_2786_v3_.f11
        lhs294_ = d_2959_v108_
        lhs295_ = default__.safeIndex(56, (d_2959_v108_).length(0))
        lhs296_ = globalState
        r1 = rhs451_
        d_2782_v0_ = rhs452_
        lhs294_[lhs295_] = rhs453_
        lhs296_.f8 = rhs454_
        d_2960_v109_: _dafny.Map
        d_2960_v109_ = _dafny.Map({d_2828_v24_: p0})
        d_2961_v110_: D0
        d_2961_v110_ = D0_DC2((p0) * (len(d_2960_v109_)), d_2949_v106_)
        r0 = d_2961_v110_
        d_2962_v111_: _dafny.Map
        d_2962_v111_ = _dafny.Map({d_2785_v2_: p2})
        r1 = len((p2) + (((d_2962_v111_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqst"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqst"))) in (d_2962_v111_) else _dafny.SeqWithoutIsStrInference([self.f11, self.f16]))))
        d_2963_v112_: C5
        nw433_ = C5()
        nw433_.ctor__(len(d_2829_v25_))
        d_2963_v112_ = nw433_
        d_2964_v113_: _dafny.Set
        d_2964_v113_ = _dafny.Set({d_2963_v112_})
        d_2965_v114_: _dafny.Map
        d_2965_v114_ = _dafny.Map({True: d_2831_v27_})
        d_2966_v115_: _dafny.Map
        d_2966_v115_ = _dafny.Map({d_2964_v113_: (d_2965_v114_) | (d_2965_v114_)})
        r2 = ((d_2966_v115_)[d_2964_v113_] if (d_2964_v113_) in (d_2966_v115_) else d_2965_v114_)
        return r0, r1, r2


class C16(T0, T1):
    def  __init__(self):
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: int = int(0)
        self._f14: bool = False
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C16"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f14, f15, f11, f12, f13):
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13

    def fm5(self, p0, globalState):
        return D0_DC1((self).f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agcumto")), default__.safeModuloInt((0) - ((self).f15), (self).f12), self.f11)

    def fm6(self, p0, p1, p2, globalState):
        return not((self).f14)

    def fm7(self, globalState):
        return not((917) != (default__.safeModuloInt((self).f13, (self).f15)))

    def fm8(self, p0, p1, p2, globalState):
        if self.f11:
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2967_i0_ in range(default__.abs(108))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_2968_i1_ in range(default__.abs(542))]))
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcgbxqc"))

    def fm9(self, p0, p1, globalState):
        return ((self).f13) < ((self).f15)

    def m3(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: D0 = D0.default()()
        r3: _dafny.Map = _dafny.Map({})
        d_2969_v0_: _dafny.Array
        def lambda147_(d_2970_i0_):
            return (d_2970_i0_) - ((self).f15)

        init80_ = lambda147_
        nw434_ = _dafny.Array(None, 7)
        for i0_80_ in range(nw434_.length(0)):
            nw434_[i0_80_] = init80_(i0_80_)
        d_2969_v0_ = nw434_
        d_2971_v1_: _dafny.Array
        nw435_ = _dafny.Array(None, 11)
        nw435_[int(0)] = d_2969_v0_
        nw435_[int(1)] = d_2969_v0_
        nw435_[int(2)] = d_2969_v0_
        nw435_[int(3)] = d_2969_v0_
        nw435_[int(4)] = d_2969_v0_
        nw435_[int(5)] = d_2969_v0_
        nw435_[int(6)] = d_2969_v0_
        nw435_[int(7)] = d_2969_v0_
        nw435_[int(8)] = d_2969_v0_
        nw435_[int(9)] = d_2969_v0_
        nw435_[int(10)] = d_2969_v0_
        d_2971_v1_ = nw435_
        d_2972_v2_: _dafny.Map
        d_2972_v2_ = _dafny.Map({-465: True})
        d_2973_v3_: T2
        nw436_ = C2()
        nw436_.ctor__((self).f15, p0, (self).f12, len(d_2972_v2_), True)
        d_2973_v3_ = nw436_
        d_2974_v4_: _dafny.Seq
        d_2974_v4_ = _dafny.SeqWithoutIsStrInference([d_2973_v3_])
        d_2975_v5_: _dafny.Map
        d_2975_v5_ = _dafny.Map({len(d_2974_v4_): (d_2973_v3_).f13})
        d_2976_v6_: D7
        d_2976_v6_ = D7_DC23(d_2971_v1_, _dafny.SeqWithoutIsStrInference([self.f11]), (self).f14, len(d_2975_v5_))
        d_2977_v7_: C11
        nw437_ = C11()
        nw437_.ctor__((self).f15, (d_2976_v6_).cf45)
        d_2977_v7_ = nw437_
        d_2978_v8_: C14
        nw438_ = C14()
        nw438_.ctor__(not((self).f14))
        d_2978_v8_ = nw438_
        d_2979_v9_: _dafny.MultiSet
        d_2979_v9_ = _dafny.MultiSet([d_2978_v8_])
        d_2980_v10_: _dafny.Set
        d_2980_v10_ = _dafny.Set({d_2979_v9_})
        d_2981_i1_: int
        d_2981_i1_ = 0
        with _dafny.label("15"):
            while (d_2980_v10_).issubset(_dafny.Set({_dafny.MultiSet([d_2978_v8_])})):
                with _dafny.c_label("15"):
                    if (d_2981_i1_) >= (100):
                        raise _dafny.Break("15")
                    d_2981_i1_ = (d_2981_i1_) + (1)
                    (self).f11 = self.f11
                    d_2982_v11_: int
                    d_2982_v11_ = 963
                    d_2982_v11_ = (d_2973_v3_).f13
                    d_2982_v11_ = (0) - ((0) - ((self).f13))
                    if self.f11:
                        index420_ = default__.safeIndex(715, (d_2969_v0_).length(0))
                        (d_2969_v0_)[index420_] = default__.safeModuloInt((d_2973_v3_).f12, (d_2973_v3_).f12)
                        index421_ = default__.safeIndex(715, (d_2969_v0_).length(0))
                        (d_2969_v0_)[index421_] = (d_2977_v7_).fm14(globalState)
                        d_2983_v12_: _dafny.Seq
                        d_2983_v12_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                        d_2983_v12_ = d_2983_v12_
                        d_2984_v13_: _dafny.Set
                        d_2984_v13_ = _dafny.Set({838, (self).f13})
                        d_2985_v14_: _dafny.Map
                        d_2985_v14_ = _dafny.Map({not(((d_2973_v3_).f12) == (len(d_2984_v13_))): (d_2973_v3_).f20})
                        d_2985_v14_ = (d_2985_v14_).set((self).f14, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2986_i2_ in range(default__.abs(-176))])).set(default__.safeIndex((d_2973_v3_).f12, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2987_i2_ in range(default__.abs(-176))]))), (p0)[default__.safeIndex(len(d_2983_v12_), len(p0))]))
                        d_2988_v15_: _dafny.Array
                        nw439_ = _dafny.Array(_dafny.Array(None, 0), 3)
                        d_2988_v15_ = nw439_
                        d_2989_v16_: _dafny.Array
                        def lambda148_(d_2990_i3_):
                            return self.f11

                        init81_ = lambda148_
                        nw440_ = _dafny.Array(None, 20)
                        for i0_81_ in range(nw440_.length(0)):
                            nw440_[i0_81_] = init81_(i0_81_)
                        d_2989_v16_ = nw440_
                        index422_ = default__.safeIndex(414, (d_2988_v15_).length(0))
                        (d_2988_v15_)[index422_] = d_2989_v16_
                        index423_ = default__.safeIndex(715, (d_2969_v0_).length(0))
                        index424_ = default__.safeIndex(414, (d_2988_v15_).length(0))
                        rhs455_ = self.f11
                        rhs456_ = (d_2973_v3_).f13
                        rhs457_ = d_2989_v16_
                        rhs458_ = self.f11
                        lhs297_ = globalState
                        lhs298_ = d_2969_v0_
                        lhs299_ = default__.safeIndex(715, (d_2969_v0_).length(0))
                        lhs300_ = d_2988_v15_
                        lhs301_ = default__.safeIndex(414, (d_2988_v15_).length(0))
                        lhs302_ = globalState
                        lhs297_.f8 = rhs455_
                        lhs298_[lhs299_] = rhs456_
                        lhs300_[lhs301_] = rhs457_
                        lhs302_.f8 = rhs458_
                        d_2991_v17_: _dafny.Map
                        d_2991_v17_ = _dafny.Map({_dafny.CodePoint('x'): (d_2973_v3_).f12})
                        d_2992_v18_: str
                        d_2992_v18_ = _dafny.CodePoint('n')
                        d_2991_v17_ = (d_2991_v17_).set(d_2992_v18_, (d_2973_v3_).f12)
                    elif True:
                        d_2993_v19_: _dafny.Array
                        def lambda149_(d_2994_v3_, d_2995_v2_):
                            def lambda150_(d_2996_i4_):
                                return _dafny.SeqWithoutIsStrInference([(d_2994_v3_).f12, (self).f13, len(d_2995_v2_)])

                            return lambda150_

                        init82_ = lambda149_(d_2973_v3_, d_2972_v2_)
                        nw441_ = _dafny.Array(None, 8)
                        for i0_82_ in range(nw441_.length(0)):
                            nw441_[i0_82_] = init82_(i0_82_)
                        d_2993_v19_ = nw441_
                        d_2997_v20_: _dafny.Seq
                        d_2997_v20_ = _dafny.SeqWithoutIsStrInference([len((d_2973_v3_).f20)])
                        index425_ = default__.safeIndex(707, (d_2993_v19_).length(0))
                        (d_2993_v19_)[index425_] = (d_2997_v20_) + (d_2997_v20_)
                        index426_ = default__.safeIndex(707, (d_2993_v19_).length(0))
                        (d_2993_v19_)[index426_] = d_2997_v20_
                        d_2998_v21_: str
                        d_2998_v21_ = _dafny.CodePoint('w')
                        d_2999_v22_: C5
                        nw442_ = C5()
                        nw442_.ctor__(((d_2973_v3_).f12 if (self).f14 else len(_dafny.SeqWithoutIsStrInference([d_2998_v21_, d_2998_v21_, d_2998_v21_]))))
                        d_2999_v22_ = nw442_
                        index427_ = default__.safeIndex(76, (d_2969_v0_).length(0))
                        (d_2969_v0_)[index427_] = 709
                        index428_ = default__.safeIndex(76, (d_2969_v0_).length(0))
                        (d_2969_v0_)[index428_] = (0) - (((self).f12) - (((self).f13 if self.f11 else (_dafny.MultiSet([d_2982_v11_, (d_2999_v22_).f24])).cardinality)))
                        d_3000_v23_: _dafny.Array
                        nw443_ = _dafny.Array(False, 7)
                        d_3000_v23_ = nw443_
                        d_3001_v24_: _dafny.MultiSet
                        d_3001_v24_ = _dafny.MultiSet([d_3000_v23_])
                        d_3002_v25_: C12
                        nw444_ = C12()
                        nw444_.ctor__(d_3001_v24_, (self).f12, self.f11, (d_2969_v0_)[default__.safeIndex(76, (d_2969_v0_).length(0))], (d_2973_v3_).f12)
                        d_3002_v25_ = nw444_
                        d_3003_v26_: C12
                        d_3003_v26_ = (d_3002_v25_ if self.f11 else d_3002_v25_)
                        d_3004_v27_: C10
                        nw445_ = C10()
                        nw445_.ctor__((self).f14)
                        d_3004_v27_ = nw445_
                        d_3005_v28_: _dafny.Map
                        d_3005_v28_ = _dafny.Map({d_3004_v27_: d_3002_v25_})
                        d_3003_v26_ = ((d_3005_v28_)[d_3004_v27_] if (d_3004_v27_) in (d_3005_v28_) else d_3002_v25_)
                        d_3006_v29_: D6
                        d_3006_v29_ = D6_DC18(self.f11)
                        d_3007_v30_: _dafny.Array
                        nw446_ = _dafny.Array(None, 1)
                        nw446_[int(0)] = d_3006_v29_
                        d_3007_v30_ = nw446_
                        d_3008_v31_: _dafny.MultiSet
                        d_3008_v31_ = _dafny.MultiSet([d_3007_v30_])
                        d_3009_v32_: D11
                        d_3009_v32_ = D11_DC30((d_3008_v31_) - (d_3008_v31_))
                        d_3010_v33_: _dafny.Array
                        nw447_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_3010_v33_ = nw447_
                        index429_ = default__.safeIndex(905, (d_3000_v23_).length(0))
                        (d_3000_v23_)[index429_] = self.f11
                        d_3011_v34_: _dafny.Map
                        d_3011_v34_ = _dafny.Map({d_3000_v23_: (0) - ((self).f13)})
                        d_3012_v35_: _dafny.Set
                        d_3012_v35_ = _dafny.Set({(self).f13, (d_2973_v3_).f13})
                        d_3013_v37_: _dafny.MultiSet
                        d_3013_v37_ = _dafny.MultiSet([d_2998_v21_])
                        d_3014_v38_: _dafny.MultiSet
                        d_3014_v38_ = _dafny.MultiSet([(self).f15, len(d_2972_v2_), (self).f13, ((d_3013_v37_)[d_2998_v21_] if (d_2998_v21_) in (d_3013_v37_) else (self).f12), (d_2977_v7_).fm14(globalState)])
                        d_3015_v39_: T0
                        nw448_ = C15()
                        nw448_.ctor__(self.f11, (d_2973_v3_).f12, self.f11)
                        d_3015_v39_ = nw448_
                        d_3016_v40_: _dafny.Map
                        d_3016_v40_ = _dafny.Map({self.f11: d_3015_v39_})
                        index430_ = default__.safeIndex(76, (d_2969_v0_).length(0))
                        index431_ = default__.safeIndex(905, (d_3000_v23_).length(0))
                        def iife265_():
                            coll111_ = _dafny.Set()
                            compr_111_: int
                            for compr_111_ in _dafny.IntegerRange(-330, 581):
                                d_3017_v36_: int = compr_111_
                                if ((-330) <= (d_3017_v36_)) and ((d_3017_v36_) < (581)):
                                    coll111_ = coll111_.union(_dafny.Set([default__.safeDivisionInt(d_3017_v36_, (self).f12)]))
                            return _dafny.Set(coll111_)
                        rhs459_ = (((d_2993_v19_)[default__.safeIndex(707, (d_2993_v19_).length(0))]) + (default__.fm0(globalState))) < ((d_2993_v19_)[default__.safeIndex(707, (d_2993_v19_).length(0))])
                        rhs460_ = d_3009_v32_
                        rhs461_ = d_3010_v33_
                        rhs462_ = len((d_3011_v34_) | (d_3011_v34_))
                        rhs463_ = (d_2978_v8_).fm6((d_3012_v35_) == (iife265_()
                        ), (d_2973_v3_).f13, ((d_3014_v38_)[(0) - (len(d_3016_v40_))] if ((0) - (len(d_3016_v40_))) in (d_3014_v38_) else (d_2973_v3_).f13), globalState)
                        lhs303_ = globalState
                        lhs304_ = d_2969_v0_
                        lhs305_ = default__.safeIndex(76, (d_2969_v0_).length(0))
                        lhs306_ = d_3000_v23_
                        lhs307_ = default__.safeIndex(905, (d_3000_v23_).length(0))
                        lhs303_.f8 = rhs459_
                        d_3009_v32_ = rhs460_
                        d_3010_v33_ = rhs461_
                        lhs304_[lhs305_] = rhs462_
                        lhs306_[lhs307_] = rhs463_
                    pass
            pass
        hi13_ = (d_2973_v3_).f13
        for d_3018_i5_ in range((self).f12, hi13_):
            index432_ = default__.safeIndex(663, (d_2969_v0_).length(0))
            (d_2969_v0_)[index432_] = (self).f12
            index433_ = default__.safeIndex(663, (d_2969_v0_).length(0))
            (d_2969_v0_)[index433_] = (self).f15
            d_3019_v41_: _dafny.Array
            nw449_ = _dafny.Array(False, 15)
            d_3019_v41_ = nw449_
            d_3019_v41_ = d_3019_v41_
            if (d_2973_v3_).fm7(globalState):
                d_3020_v42_: _dafny.Array
                def lambda151_(d_3021_i6_):
                    return (d_3021_i6_) + ((self).f13)

                init83_ = lambda151_
                nw450_ = _dafny.Array(None, 13)
                for i0_83_ in range(nw450_.length(0)):
                    nw450_[i0_83_] = init83_(i0_83_)
                d_3020_v42_ = nw450_
                d_3020_v42_ = d_3020_v42_
                d_3022_v43_: _dafny.Set
                d_3022_v43_ = _dafny.Set({(d_2973_v3_).f13})
                d_3023_v44_: _dafny.Map
                d_3023_v44_ = _dafny.Map({d_3020_v42_: d_3022_v43_})
                d_3022_v43_ = ((d_3023_v44_)[(d_3020_v42_ if (self).f14 else d_3020_v42_)] if ((d_3020_v42_ if (self).f14 else d_3020_v42_)) in (d_3023_v44_) else (d_3022_v43_) | (d_3022_v43_))
                d_3024_v45_: _dafny.Set
                d_3024_v45_ = _dafny.Set({self.f11})
                d_3025_v46_: D19
                d_3025_v46_ = D19_DC51(len(d_3024_v45_), (d_2973_v3_).f13)
                d_3026_v47_: _dafny.Map
                d_3026_v47_ = _dafny.Map({self.f11: d_3025_v46_})
                d_3026_v47_ = (d_3026_v47_).set((self).f14, D19_DC51((self).f13, (d_2973_v3_).f13))
                d_3027_v48_: _dafny.Seq
                d_3027_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sutgkklap"))
                d_3028_v49_: _dafny.Map
                d_3028_v49_ = _dafny.Map({self.f11: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_3029_i7_ in range(default__.abs(596))]))})
                d_3030_v50_: str
                d_3030_v50_ = _dafny.CodePoint('p')
                index434_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                rhs464_ = default__.safeDivisionInt(((self).f12) * ((self).f12), (d_2973_v3_).f12)
                rhs465_ = (default__.fm47(self.f11, False, d_3028_v49_, (d_2973_v3_).f12, globalState)) | (d_2972_v2_)
                rhs466_ = (self).f14
                rhs467_ = (((d_2969_v0_)[default__.safeIndex(663, (d_2969_v0_).length(0))]) + ((self).f15)) != ((default__.fm2(globalState) if (self).f14 else (d_2973_v3_).f13))
                rhs468_ = (d_3027_v48_) + ((((d_2973_v3_).f20) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggglnxi")))).set(default__.safeIndex((d_2973_v3_).f13, len(((d_2973_v3_).f20) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggglnxi"))))), d_3030_v50_))
                lhs308_ = d_2969_v0_
                lhs309_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                lhs308_[lhs309_] = rhs464_
                r3 = rhs465_
                r0 = rhs466_
                r0 = rhs467_
                d_3027_v48_ = rhs468_
                d_3031_v51_: _dafny.Seq
                d_3031_v51_ = _dafny.SeqWithoutIsStrInference([self.f11, (self).f14, (self).f14, True, self.f11])
                (self).f11 = ((self.f11) in (d_3028_v49_)) not in ((d_3031_v51_) + ((_dafny.SeqWithoutIsStrInference([self.f11, self.f11])).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference([self.f11, self.f11]))), True)))
            elif True:
                index435_ = default__.safeIndex(445, (d_3019_v41_).length(0))
                (d_3019_v41_)[index435_] = ((self).f14 if (self).fm9((self).f14, (self).fm7(globalState), globalState) else not(self.f11))
                index436_ = default__.safeIndex(445, (d_3019_v41_).length(0))
                (d_3019_v41_)[index436_] = self.f11
                d_3032_v52_: _dafny.Array
                nw451_ = _dafny.Array(_dafny.MultiSet({}), 19)
                d_3032_v52_ = nw451_
                d_3033_v53_: _dafny.MultiSet
                d_3033_v53_ = _dafny.MultiSet([default__.fm2(globalState)])
                index437_ = default__.safeIndex(552, (d_3032_v52_).length(0))
                (d_3032_v52_)[index437_] = d_3033_v53_
                d_3034_v54_: _dafny.Array
                nw452_ = _dafny.Array(D0.default()(), 5)
                d_3034_v54_ = nw452_
                index438_ = default__.safeIndex(856, (d_3034_v54_).length(0))
                (d_3034_v54_)[index438_] = (p1 if (d_3019_v41_)[default__.safeIndex(445, (d_3019_v41_).length(0))] else p1)
                d_3035_v55_: _dafny.Seq
                d_3035_v55_ = _dafny.SeqWithoutIsStrInference([self.f11, True, (self).f14])
                d_3036_v56_: _dafny.Map
                d_3036_v56_ = _dafny.Map({len(d_3035_v55_): d_3033_v53_})
                index439_ = default__.safeIndex(445, (d_3019_v41_).length(0))
                index440_ = default__.safeIndex(552, (d_3032_v52_).length(0))
                index441_ = default__.safeIndex(445, (d_3019_v41_).length(0))
                index442_ = default__.safeIndex(856, (d_3034_v54_).length(0))
                rhs469_ = (self).f14
                rhs470_ = ((d_3036_v56_)[default__.safeModuloInt(777, (self).f12)] if (default__.safeModuloInt(777, (self).f12)) in (d_3036_v56_) else (d_3033_v53_).set((d_2973_v3_).f12, default__.abs(len(d_3035_v55_))))
                rhs471_ = (self).f14
                rhs472_ = p1
                lhs310_ = d_3019_v41_
                lhs311_ = default__.safeIndex(445, (d_3019_v41_).length(0))
                lhs312_ = d_3032_v52_
                lhs313_ = default__.safeIndex(552, (d_3032_v52_).length(0))
                lhs314_ = d_3019_v41_
                lhs315_ = default__.safeIndex(445, (d_3019_v41_).length(0))
                lhs316_ = d_3034_v54_
                lhs317_ = default__.safeIndex(856, (d_3034_v54_).length(0))
                lhs310_[lhs311_] = rhs469_
                lhs312_[lhs313_] = rhs470_
                lhs314_[lhs315_] = rhs471_
                lhs316_[lhs317_] = rhs472_
                d_3037_v57_: _dafny.Array
                def lambda152_(d_3038_i8_):
                    return False

                init84_ = lambda152_
                nw453_ = _dafny.Array(None, 2)
                for i0_84_ in range(nw453_.length(0)):
                    nw453_[i0_84_] = init84_(i0_84_)
                d_3037_v57_ = nw453_
                d_3039_v58_: _dafny.Map
                d_3039_v58_ = _dafny.Map({(self).f14: (self).f14})
                d_3037_v57_ = (d_3037_v57_ if ((d_3039_v58_)[(d_3019_v41_)[default__.safeIndex(445, (d_3019_v41_).length(0))]] if ((d_3019_v41_)[default__.safeIndex(445, (d_3019_v41_).length(0))]) in (d_3039_v58_) else self.f11) else d_3037_v57_)
                d_3040_v59_: T1
                nw454_ = C3()
                nw454_.ctor__((d_2973_v3_).f13, len((d_3035_v55_) + (d_3035_v55_)))
                d_3040_v59_ = nw454_
                d_3041_v60_: _dafny.Seq
                d_3041_v60_ = _dafny.SeqWithoutIsStrInference([d_3035_v55_])
                d_3042_v61_: C10
                nw455_ = C10()
                nw455_.ctor__(self.f11)
                d_3042_v61_ = nw455_
                d_3043_v62_: D6
                d_3043_v62_ = D6_DC19((self).f14, False, False, self.f11, (d_3041_v60_)[default__.safeIndex(len(_dafny.Map({d_3042_v61_: (self).f14})), len(d_3041_v60_))])
                d_3044_v63_: _dafny.Map
                d_3044_v63_ = _dafny.Map({(self).f14: d_3043_v62_})
                nw456_ = C7()
                nw456_.ctor__(len((d_3044_v63_) | (_dafny.Map({(d_3019_v41_)[default__.safeIndex(445, (d_3019_v41_).length(0))]: d_3043_v62_}))), default__.safeModuloInt(((d_3033_v53_)[(self).f12] if ((self).f12) in (d_3033_v53_) else (d_2973_v3_).f12), (d_2973_v3_).f12))
                d_3040_v59_ = nw456_
                index443_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                (d_2969_v0_)[index443_] = (((d_3032_v52_)[default__.safeIndex(552, (d_3032_v52_).length(0))])[(d_2973_v3_).f13] if ((d_2973_v3_).f13) in ((d_3032_v52_)[default__.safeIndex(552, (d_3032_v52_).length(0))]) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bthuoibuq"))))
            if self.f11:
                index444_ = default__.safeIndex(361, (d_3019_v41_).length(0))
                (d_3019_v41_)[index444_] = ((d_2973_v3_).f12) == ((self).f13)
                d_3045_v64_: _dafny.Set
                d_3045_v64_ = _dafny.Set({self.f11, self.f11})
                d_3046_v65_: _dafny.Map
                d_3046_v65_ = _dafny.Map({d_2972_v2_: (self).f14})
                d_3047_v66_: D20
                d_3047_v66_ = D20_DC52(d_3046_v65_)
                d_3048_v67_: D20
                d_3048_v67_ = D20_DC54(d_3047_v66_)
                d_3049_v68_: _dafny.MultiSet
                d_3049_v68_ = _dafny.MultiSet([d_3048_v67_])
                d_3050_v69_: _dafny.Seq
                d_3050_v69_ = _dafny.SeqWithoutIsStrInference([(self).f12, (d_2973_v3_).f12, 572, (d_3049_v68_).cardinality, (d_2973_v3_).f13])
                d_3051_v70_: _dafny.Seq
                d_3051_v70_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
                d_3052_v71_: _dafny.Set
                d_3052_v71_ = _dafny.Set({d_3051_v70_, d_3051_v70_})
                index445_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                index446_ = default__.safeIndex(361, (d_3019_v41_).length(0))
                index447_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                index448_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                rhs473_ = (d_2973_v3_).f13
                rhs474_ = (self).f14
                rhs475_ = (len((d_3045_v64_) - (_dafny.Set({self.f11, True})))) - ((d_2973_v3_).f13)
                rhs476_ = (default__.safeDivisionInt((d_2973_v3_).f13, (d_3050_v69_)[default__.safeIndex((self).f15, len(d_3050_v69_))]) if (self.f11 if (self).f14 else self.f11) else 847)
                rhs477_ = (d_3051_v70_) not in (d_3052_v71_)
                lhs318_ = d_2969_v0_
                lhs319_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                lhs320_ = d_3019_v41_
                lhs321_ = default__.safeIndex(361, (d_3019_v41_).length(0))
                lhs322_ = d_2969_v0_
                lhs323_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                lhs324_ = d_2969_v0_
                lhs325_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                lhs326_ = globalState
                lhs318_[lhs319_] = rhs473_
                lhs320_[lhs321_] = rhs474_
                lhs322_[lhs323_] = rhs475_
                lhs324_[lhs325_] = rhs476_
                lhs326_.f8 = rhs477_
                d_3053_v72_: _dafny.Array
                nw457_ = _dafny.Array(int(0), 11)
                d_3053_v72_ = nw457_
                index449_ = default__.safeIndex(138, (d_2971_v1_).length(0))
                (d_2971_v1_)[index449_] = d_3053_v72_
                index450_ = default__.safeIndex(138, (d_2971_v1_).length(0))
                (d_2971_v1_)[index450_] = d_3053_v72_
                d_2975_v5_ = (d_2975_v5_).set((self).f12, ((self).f13 if self.f11 else len(p0)))
                index451_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                (d_2969_v0_)[index451_] = default__.safeModuloInt((d_2973_v3_).f13, -80)
                index452_ = default__.safeIndex(361, (d_3019_v41_).length(0))
                (d_3019_v41_)[index452_] = (d_3019_v41_)[default__.safeIndex(361, (d_3019_v41_).length(0))]
            elif True:
                d_3054_v73_: _dafny.MultiSet
                d_3054_v73_ = _dafny.MultiSet([(d_2973_v3_).f12])
                d_3055_v74_: _dafny.Array
                nw458_ = _dafny.Array(None, 2)
                nw458_[int(0)] = (d_3054_v73_).cardinality
                nw458_[int(1)] = (d_2973_v3_).f13
                d_3055_v74_ = nw458_
                d_3056_v75_: _dafny.MultiSet
                d_3056_v75_ = _dafny.MultiSet([(D0_DC2((self).f12, d_3055_v74_)).cf6, d_3055_v74_])
                index453_ = default__.safeIndex(663, (d_2969_v0_).length(0))
                (d_2969_v0_)[index453_] = (0) - (((d_3056_v75_)[d_3055_v74_] if (d_3055_v74_) in (d_3056_v75_) else ((d_3054_v73_)[(d_2973_v3_).f13] if ((d_2973_v3_).f13) in (d_3054_v73_) else (d_2973_v3_).f12)))
                d_3057_v76_: _dafny.Seq
                d_3057_v76_ = _dafny.SeqWithoutIsStrInference([(d_2973_v3_).f20, (d_2973_v3_).f20])
                d_3058_v78_: _dafny.Map
                d_3058_v78_ = _dafny.Map({not(self.f11): (d_2969_v0_)[default__.safeIndex(663, (d_2969_v0_).length(0))]})
                d_3059_v79_: _dafny.Seq
                d_3059_v79_ = _dafny.SeqWithoutIsStrInference([len(d_3058_v78_)])
                d_3060_v80_: _dafny.Seq
                d_3060_v80_ = _dafny.SeqWithoutIsStrInference([self.f11])
                d_3061_v81_: _dafny.Set
                d_3061_v81_ = _dafny.Set({not((d_3060_v80_)[default__.safeIndex((d_2973_v3_).f12, len(d_3060_v80_))]), self.f11})
                d_3062_v82_: _dafny.Map
                def iife266_():
                    coll112_ = _dafny.Map()
                    compr_112_: int
                    for compr_112_ in (d_3059_v79_).Elements:
                        d_3063_v77_: int = compr_112_
                        if (d_3063_v77_) in (d_3059_v79_):
                            coll112_[default__.safeModuloInt(d_3063_v77_, (self).f15)] = len(_dafny.Set({self.f11, self.f11, self.f11}))
                    return _dafny.Map(coll112_)
                d_3062_v82_ = _dafny.Map({(d_3057_v76_)[default__.safeIndex(len(iife266_()
                ), len(d_3057_v76_))]: d_3061_v81_})
                d_3062_v82_ = (d_3062_v82_) | (d_3062_v82_)
                d_3064_v83_: _dafny.Array
                def lambda153_(d_3065_p0_):
                    def lambda154_(d_3066_i9_):
                        return d_3065_p0_

                    return lambda154_

                init85_ = lambda153_(p0)
                nw459_ = _dafny.Array(None, 24)
                for i0_85_ in range(nw459_.length(0)):
                    nw459_[i0_85_] = init85_(i0_85_)
                d_3064_v83_ = nw459_
                index454_ = default__.safeIndex(872, (d_3064_v83_).length(0))
                (d_3064_v83_)[index454_] = ((d_2973_v3_).f20) + (p0)
                index455_ = default__.safeIndex(872, (d_3064_v83_).length(0))
                (d_3064_v83_)[index455_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3067_i10_ in range(default__.abs(67))])
                d_2978_v8_ = d_2978_v8_
                d_3068_v84_: D7
                d_3068_v84_ = D7_DC22((d_2973_v3_).f13, d_2971_v1_)
                pat_let_tv119_ = d_2969_v0_
                pat_let_tv120_ = d_2969_v0_
                d_3069_v85_: _dafny.Array
                nw460_ = _dafny.Array(None, 15)
                def iife267_(_pat_let77_0):
                    def iife268_(d_3070_dt__update__tmp_h1_):
                        def iife269_(_pat_let78_0):
                            def iife270_(d_3071_dt__update_hcf40_h0_):
                                return D7_DC22(d_3071_dt__update_hcf40_h0_, (d_3070_dt__update__tmp_h1_).cf41)
                            return iife270_(_pat_let78_0)
                        return iife269_((pat_let_tv120_)[default__.safeIndex(663, (pat_let_tv119_).length(0))])
                    return iife268_(_pat_let77_0)
                nw460_[int(0)] = iife267_(d_3068_v84_)
                def iife271_(_pat_let79_0):
                    def iife272_(d_3072_dt__update__tmp_h2_):
                        def iife273_(_pat_let80_0):
                            def iife274_(d_3073_dt__update_hcf40_h1_):
                                return D7_DC22(d_3073_dt__update_hcf40_h1_, (d_3072_dt__update__tmp_h2_).cf41)
                            return iife274_(_pat_let80_0)
                        return iife273_((self).f12)
                    return iife272_(_pat_let79_0)
                nw460_[int(1)] = iife271_(D7_DC22((d_2973_v3_).f13, d_2971_v1_))
                nw460_[int(2)] = D7_DC22((d_2977_v7_).fm14(globalState), d_2971_v1_)
                nw460_[int(3)] = d_3068_v84_
                nw460_[int(4)] = d_3068_v84_
                nw460_[int(5)] = D7_DC22((self).f12, d_2971_v1_)
                nw460_[int(6)] = D7_DC22((d_2973_v3_).f13, d_2971_v1_)
                nw460_[int(7)] = d_3068_v84_
                nw460_[int(8)] = d_3068_v84_
                nw460_[int(9)] = d_3068_v84_
                nw460_[int(10)] = d_3068_v84_
                nw460_[int(11)] = d_3068_v84_
                nw460_[int(12)] = d_3068_v84_
                nw460_[int(13)] = d_3068_v84_
                nw460_[int(14)] = d_3068_v84_
                d_3069_v85_ = nw460_
                d_3074_v86_: _dafny.MultiSet
                d_3074_v86_ = _dafny.MultiSet([d_3069_v85_])
                d_3074_v86_ = d_3074_v86_
        d_3075_i11_: int
        d_3075_i11_ = 0
        with _dafny.label("16"):
            while not (self.f11) or ((self).f14):
                with _dafny.c_label("16"):
                    if (d_3075_i11_) >= (100):
                        raise _dafny.Break("16")
                    d_3075_i11_ = (d_3075_i11_) + (1)
                    d_3076_v87_: _dafny.Array
                    nw461_ = _dafny.Array(None, 2)
                    nw461_[int(0)] = (self).f14
                    nw461_[int(1)] = self.f11
                    d_3076_v87_ = nw461_
                    index456_ = default__.safeIndex(510, (d_3076_v87_).length(0))
                    (d_3076_v87_)[index456_] = False
                    d_3077_v88_: D0
                    d_3077_v88_ = D0_DC1((self).f15, (d_2973_v3_).f20, (d_2973_v3_).f12, True)
                    index457_ = default__.safeIndex(510, (d_3076_v87_).length(0))
                    (d_3076_v87_)[index457_] = (((d_2973_v3_).f20 if self.f11 else (d_2973_v3_).f20)) <= ((d_3077_v88_).cf2)
                    index458_ = default__.safeIndex(5, (d_2971_v1_).length(0))
                    (d_2971_v1_)[index458_] = d_2969_v0_
                    index459_ = default__.safeIndex(5, (d_2971_v1_).length(0))
                    def lambda155_(d_3078_v3_):
                        def lambda156_(d_3079_i12_):
                            return default__.safeModuloInt(d_3079_i12_, (d_3078_v3_).f13)

                        return lambda156_

                    init86_ = lambda155_(d_2973_v3_)
                    nw462_ = _dafny.Array(None, 26)
                    for i0_86_ in range(nw462_.length(0)):
                        nw462_[i0_86_] = init86_(i0_86_)
                    (d_2971_v1_)[index459_] = nw462_
                    d_3080_v89_: C3
                    nw463_ = C3()
                    nw463_.ctor__(((d_2973_v3_).f13) - ((self).f12), 743)
                    d_3080_v89_ = nw463_
                    d_3081_v90_: D20
                    d_3081_v90_ = D20_DC53()
                    d_3081_v90_ = (d_3081_v90_ if (self).f14 else d_3081_v90_)
                    pass
            pass
        d_3082_v91_: _dafny.Seq
        d_3082_v91_ = _dafny.SeqWithoutIsStrInference([(d_2973_v3_).f13])
        d_3083_v92_: C8
        nw464_ = C8()
        nw464_.ctor__(d_3082_v91_)
        d_3083_v92_ = nw464_
        d_3084_v93_: int
        d_3084_v93_ = 736
        d_3085_v94_: _dafny.Map
        d_3085_v94_ = _dafny.Map({_dafny.Map({(self).f13: (d_2973_v3_).f12}): d_2969_v0_})
        d_3084_v93_ = ((0) - ((d_2973_v3_).f13)) + (len(d_3085_v94_))
        r0 = self.f11
        d_3086_v95_: _dafny.MultiSet
        d_3086_v95_ = _dafny.MultiSet([not(self.f11), (d_2978_v8_).fm6((self).f14, (d_2977_v7_).fm14(globalState), (self).f12, globalState), self.f11])
        r1 = d_3086_v95_
        d_3087_v96_: _dafny.Array
        nw465_ = _dafny.Array(_dafny.Seq({}), 27)
        d_3087_v96_ = nw465_
        r2 = D0_DC0(d_3087_v96_)
        r3 = ((d_2972_v2_) | (d_2972_v2_)) | ((d_2972_v2_ if (self).f14 else d_2972_v2_))
        return r0, r1, r2, r3

    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15

class C17(T0):
    def  __init__(self):
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C17"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    def ctor__(self, f11):
        (self).f11 = f11

    def fm5(self, p0, globalState):
        def iife275_():
            def iife277_():
                coll115_ = _dafny.Map()
                compr_115_: int
                for compr_115_ in _dafny.IntegerRange(-642, 865):
                    d_3088_v1_: int = compr_115_
                    if ((-642) <= (d_3088_v1_)) and ((d_3088_v1_) < (865)):
                        coll115_[(d_3088_v1_) + (504)] = _dafny.SeqWithoutIsStrInference([self.f11])
                return _dafny.Map(coll115_)
            coll113_ = _dafny.Map()
            def iife276_():
                coll114_ = _dafny.Map()
                compr_114_: int
                for compr_114_ in _dafny.IntegerRange(-642, 865):
                    d_3088_v1_: int = compr_114_
                    if ((-642) <= (d_3088_v1_)) and ((d_3088_v1_) < (865)):
                        coll114_[(d_3088_v1_) + (504)] = _dafny.SeqWithoutIsStrInference([self.f11])
                return _dafny.Map(coll114_)
            compr_113_: int
            for compr_113_ in (iife276_()
            ).keys.Elements:
                d_3089_v0_: int = compr_113_
                if (d_3089_v0_) in (iife277_()
                ):
                    coll113_[(d_3089_v0_) + (770)] = -852
            return _dafny.Map(coll113_)
        return D0_DC1(len((_dafny.Set({-82})).intersection(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vldiolxvg"))), len(_dafny.SeqWithoutIsStrInference([262, 497]))])).cardinality, 390, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-67: 526})), 122, len(_dafny.Map({_dafny.Map({self.f11: 725}): -479})), len(_dafny.Set({(0) - ((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f11, self.f11]))).cardinality)), 36}))]))}))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnpq")), (0) - ((0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahuf"))), len(iife275_()
)))), self.f11)

    def fm6(self, p0, p1, p2, globalState):
        return self.f11

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        r3: bool = False
        d_3090_v0_: _dafny.Map
        d_3090_v0_ = _dafny.Map({p0: p0})
        if (((d_3090_v0_)[self.f11] if (self.f11) in (d_3090_v0_) else p2)) and (p0):
            d_3091_v1_: _dafny.Map
            d_3091_v1_ = _dafny.Map({default__.safeDivisionInt(p1, -950): p2})
            d_3091_v1_ = (d_3091_v1_).set(p1, p2)
            d_3092_v2_: _dafny.Seq
            d_3092_v2_ = _dafny.SeqWithoutIsStrInference([p1])
            d_3093_v3_: _dafny.Seq
            d_3093_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgsgelena"))
            d_3094_v4_: C2
            nw466_ = C2()
            nw466_.ctor__(len((d_3092_v2_) + (d_3092_v2_)), d_3093_v3_, p1, len(d_3092_v2_), self.f11)
            d_3094_v4_ = nw466_
            d_3095_v5_: _dafny.Array
            nw467_ = _dafny.Array(_dafny.Seq({}), 23)
            d_3095_v5_ = nw467_
            index460_ = default__.safeIndex(608, (d_3095_v5_).length(0))
            (d_3095_v5_)[index460_] = _dafny.SeqWithoutIsStrInference([p0, (d_3094_v4_).fm7(globalState), self.f11])
            d_3096_v6_: _dafny.Seq
            d_3096_v6_ = _dafny.SeqWithoutIsStrInference([False, (d_3094_v4_).fm33((d_3094_v4_).f26, globalState)])
            index461_ = default__.safeIndex(608, (d_3095_v5_).length(0))
            (d_3095_v5_)[index461_] = (d_3096_v6_) + ((d_3096_v6_) + (_dafny.SeqWithoutIsStrInference([p2, self.f11])))
            d_3097_v7_: str
            d_3097_v7_ = _dafny.CodePoint('v')
            d_3098_v8_: _dafny.Map
            d_3098_v8_ = _dafny.Map({d_3097_v7_: p0})
            d_3099_v9_: _dafny.Set
            d_3099_v9_ = _dafny.Set({157, p1})
            rhs478_ = (d_3094_v4_).fm8((d_3094_v4_).f26, d_3098_v8_, d_3099_v9_, globalState)
            rhs479_ = (d_3092_v2_) < ((_dafny.SeqWithoutIsStrInference([(d_3094_v4_).f26])) + (d_3092_v2_))
            lhs327_ = globalState
            d_3093_v3_ = rhs478_
            lhs327_.f8 = rhs479_
            r0 = p0
        elif True:
            d_3100_v10_: D7
            d_3100_v10_ = D7_DC21(_dafny.MultiSet([(self).fm6(self.f11, p1, p1, globalState), p0]))
            d_3101_v11_: _dafny.Map
            d_3101_v11_ = _dafny.Map({d_3100_v10_: not (p0) or (p0)})
            d_3102_v12_: _dafny.Array
            nw468_ = _dafny.Array(_dafny.Seq({}), 23)
            d_3102_v12_ = nw468_
            d_3103_v13_: D0
            d_3103_v13_ = D0_DC0(d_3102_v12_)
            d_3104_v14_: _dafny.Map
            d_3104_v14_ = _dafny.Map({d_3103_v13_: p0})
            d_3105_v15_: _dafny.Seq
            d_3105_v15_ = _dafny.SeqWithoutIsStrInference([d_3104_v14_, d_3104_v14_, _dafny.Map({d_3103_v13_: self.f11})])
            d_3106_v16_: _dafny.Map
            d_3106_v16_ = _dafny.Map({p2: len((d_3105_v15_)[default__.safeIndex(p1, len(d_3105_v15_))])})
            d_3107_v17_: C4
            nw469_ = C4()
            nw469_.ctor__(p1, p0)
            d_3107_v17_ = nw469_
            d_3108_v19_: _dafny.Map
            d_3108_v19_ = _dafny.Map({(d_3107_v17_).f25: p1})
            d_3109_v20_: _dafny.MultiSet
            d_3109_v20_ = _dafny.MultiSet([True])
            d_3110_v21_: _dafny.Seq
            d_3110_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            d_3111_v22_: _dafny.Array
            nw470_ = _dafny.Array(None, 22)
            nw470_[int(0)] = -444
            nw470_[int(1)] = p1
            nw470_[int(2)] = len(_dafny.Map({p0: d_3107_v17_}))
            nw470_[int(3)] = 374
            nw470_[int(4)] = p1
            def iife278_():
                coll116_ = _dafny.Map()
                compr_116_: int
                for compr_116_ in (d_3108_v19_).keys.Elements:
                    d_3112_v18_: int = compr_116_
                    if (d_3112_v18_) in (d_3108_v19_):
                        coll116_[default__.safeDivisionInt(d_3112_v18_, (d_3107_v17_).f25)] = p0
                return _dafny.Map(coll116_)
            nw470_[int(5)] = len(iife278_()
            )
            nw470_[int(6)] = p1
            nw470_[int(7)] = 287
            nw470_[int(8)] = (d_3109_v20_).cardinality
            nw470_[int(9)] = (d_3107_v17_).f25
            nw470_[int(10)] = p1
            nw470_[int(11)] = (d_3107_v17_).f25
            nw470_[int(12)] = p1
            nw470_[int(13)] = (d_3107_v17_).f25
            nw470_[int(14)] = len(d_3110_v21_)
            nw470_[int(15)] = p1
            nw470_[int(16)] = p1
            nw470_[int(17)] = (0) - ((d_3107_v17_).f25)
            nw470_[int(18)] = (d_3107_v17_).f25
            nw470_[int(19)] = (0) - (p1)
            nw470_[int(20)] = p1
            nw470_[int(21)] = (d_3107_v17_).f25
            d_3111_v22_ = nw470_
            d_3113_v23_: _dafny.Array
            nw471_ = _dafny.Array(None, 4)
            nw471_[int(0)] = d_3111_v22_
            nw471_[int(1)] = d_3111_v22_
            nw471_[int(2)] = d_3111_v22_
            nw471_[int(3)] = d_3111_v22_
            d_3113_v23_ = nw471_
            d_3114_v24_: D7
            d_3114_v24_ = D7_DC22((_dafny.MultiSet([p1, p1, -252, p1])).cardinality, d_3113_v23_)
            d_3115_v25_: _dafny.Map
            d_3115_v25_ = _dafny.Map({p3: p0})
            rhs480_ = ((d_3106_v16_)[(p2) in (d_3090_v0_)] if ((p2) in (d_3090_v0_)) in (d_3106_v16_) else p1)
            rhs481_ = (d_3101_v11_) | (d_3101_v11_)
            rhs482_ = (D9_DC27(p0, d_3114_v24_, len(d_3115_v25_), self.f11)).cf51
            rhs483_ = (d_3107_v17_).f25
            r1 = rhs480_
            d_3101_v11_ = rhs481_
            r1 = rhs482_
            r1 = rhs483_
            d_3116_v26_: str
            d_3116_v26_ = _dafny.CodePoint('m')
            d_3117_v27_: _dafny.Map
            d_3117_v27_ = _dafny.Map({(self.f11 if p0 else True): d_3116_v26_})
            r2 = ((d_3117_v27_)[p0] if (p0) in (d_3117_v27_) else default__.fm45(p1, globalState))
            d_3118_v28_: _dafny.Seq
            d_3118_v28_ = _dafny.SeqWithoutIsStrInference([(d_3107_v17_).f25, p1])
            rhs484_ = p0
            rhs485_ = ((d_3118_v28_ if p2 else d_3118_v28_)) + ((d_3118_v28_) + (_dafny.SeqWithoutIsStrInference([p1, (d_3107_v17_).f25])))
            rhs486_ = p1
            rhs487_ = ((d_3107_v17_).f25) + ((0) - (((d_3109_v20_) | (d_3109_v20_)).cardinality))
            rhs488_ = (d_3118_v28_)[default__.safeIndex((d_3107_v17_).f25, len(d_3118_v28_))]
            lhs328_ = globalState
            lhs329_ = globalState
            lhs328_.f8 = rhs484_
            lhs329_.f5 = rhs485_
            r1 = rhs486_
            r1 = rhs487_
            r1 = rhs488_
            d_3119_v29_: D0
            d_3119_v29_ = D0_DC2(p1, d_3111_v22_)
            r1 = (((d_3119_v29_).cf5 if False else p1)) - ((d_3107_v17_).f25)
            index462_ = default__.safeIndex(326, (d_3111_v22_).length(0))
            (d_3111_v22_)[index462_] = default__.safeModuloInt(((d_3106_v16_)[True] if (True) in (d_3106_v16_) else p1), (0) - (p1))
            d_3120_v30_: _dafny.Map
            d_3120_v30_ = _dafny.Map({len(_dafny.Map({p2: p0})): p0})
            d_3121_v31_: _dafny.Map
            d_3121_v31_ = _dafny.Map({d_3120_v30_: p2})
            d_3122_v32_: D20
            d_3122_v32_ = D20_DC52(d_3121_v31_)
            d_3123_v33_: _dafny.Set
            d_3123_v33_ = _dafny.Set({p0})
            d_3124_v34_: _dafny.Map
            d_3124_v34_ = _dafny.Map({d_3122_v32_: (d_3123_v33_).issubset(d_3123_v33_)})
            d_3125_v35_: _dafny.Set
            d_3125_v35_ = _dafny.Set({(d_3107_v17_).f25, (d_3107_v17_).f25})
            d_3126_v36_: D17
            d_3126_v36_ = D17_DC45(len(d_3118_v28_), d_3125_v35_)
            d_3127_v37_: C11
            nw472_ = C11()
            nw472_.ctor__(default__.fm2(globalState), default__.fm2(globalState))
            d_3127_v37_ = nw472_
            d_3128_v38_: _dafny.Seq
            d_3128_v38_ = _dafny.SeqWithoutIsStrInference([d_3127_v37_])
            d_3129_v39_: _dafny.MultiSet
            d_3129_v39_ = _dafny.MultiSet([937, default__.safeModuloInt(((d_3108_v19_)[(d_3107_v17_).f25] if ((d_3107_v17_).f25) in (d_3108_v19_) else p1), (d_3107_v17_).f25), ((d_3107_v17_).f25) * ((d_3126_v36_).cf69), len(d_3128_v38_)])
            index463_ = default__.safeIndex(326, (d_3111_v22_).length(0))
            rhs489_ = (d_3127_v37_).fm14(globalState)
            rhs490_ = self.f11
            rhs491_ = d_3124_v34_
            rhs492_ = (default__.fm12((d_3107_v17_).f25, (d_3107_v17_).f25, p1, p2, globalState) if self.f11 else (d_3129_v39_) - (d_3129_v39_))
            rhs493_ = not(p2)
            lhs330_ = d_3111_v22_
            lhs331_ = default__.safeIndex(326, (d_3111_v22_).length(0))
            lhs332_ = globalState
            lhs330_[lhs331_] = rhs489_
            r0 = rhs490_
            d_3124_v34_ = rhs491_
            d_3129_v39_ = rhs492_
            lhs332_.f8 = rhs493_
        d_3130_v40_: _dafny.Set
        d_3130_v40_ = _dafny.Set({p1, (_dafny.MultiSet([p2])).cardinality, p1, p1})
        d_3131_v41_: _dafny.Array
        nw473_ = _dafny.Array(None, 24)
        d_3131_v41_ = nw473_
        d_3132_v42_: _dafny.Seq
        d_3132_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "od"))
        rhs494_ = (d_3130_v40_).intersection(d_3130_v40_)
        rhs495_ = (0) - (default__.safeModuloInt(default__.safeModuloInt(len(d_3132_v42_), p1), p1))
        rhs496_ = ((814) == (-31)) == (p2)
        rhs497_ = d_3131_v41_
        lhs333_ = globalState
        d_3130_v40_ = rhs494_
        r1 = rhs495_
        lhs333_.f8 = rhs496_
        d_3131_v41_ = rhs497_
        hi14_ = default__.safeModuloInt(p1, p1)
        for d_3133_i0_ in range(p1, hi14_):
            r1 = d_3133_i0_
            if (p1) >= (d_3133_i0_):
                index464_ = default__.safeIndex(98, (p3).length(0))
                (p3)[index464_] = p2
                index465_ = default__.safeIndex(98, (p3).length(0))
                (p3)[index465_] = False
                d_3134_v43_: bool
                d_3135_v44_: bool
                d_3136_v45_: int
                out45_: bool
                out46_: bool
                out47_: int
                out45_, out46_, out47_ = (self).m2(globalState)
                d_3134_v43_ = out45_
                d_3135_v44_ = out46_
                d_3136_v45_ = out47_
                d_3137_v46_: bool
                d_3138_v47_: bool
                d_3139_v48_: int
                out48_: bool
                out49_: bool
                out50_: int
                out48_, out49_, out50_ = (self).m2(globalState)
                d_3137_v46_ = out48_
                d_3138_v47_ = out49_
                d_3139_v48_ = out50_
                d_3140_v49_: str
                d_3140_v49_ = _dafny.CodePoint('e')
                d_3141_v50_: _dafny.Set
                d_3141_v50_ = _dafny.Set({d_3140_v49_, d_3140_v49_})
                d_3141_v50_ = d_3141_v50_
                d_3142_v51_: D9
                d_3142_v51_ = D9_DC26(d_3140_v49_)
                d_3142_v51_ = d_3142_v51_
            elif True:
                r1 = d_3133_i0_
                d_3143_v52_: _dafny.Map
                d_3143_v52_ = _dafny.Map({p0: d_3133_i0_})
                d_3143_v52_ = (d_3143_v52_).set(p0, default__.safeDivisionInt(d_3133_i0_, d_3133_i0_))
                index466_ = default__.safeIndex(903, (p3).length(0))
                (p3)[index466_] = (p2) and (not(p0))
                d_3144_v53_: str
                d_3144_v53_ = _dafny.CodePoint('x')
                index467_ = default__.safeIndex(191, (p3).length(0))
                (p3)[index467_] = (d_3144_v53_) in (d_3132_v42_)
                d_3145_v54_: _dafny.Seq
                d_3145_v54_ = _dafny.SeqWithoutIsStrInference([d_3132_v42_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhlhneom")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wktwjoaxr"))])
                index468_ = default__.safeIndex(903, (p3).length(0))
                index469_ = default__.safeIndex(191, (p3).length(0))
                rhs498_ = self.f11
                rhs499_ = (d_3145_v54_) == ((d_3145_v54_ if self.f11 else d_3145_v54_))
                rhs500_ = (p2 if (self.f11) or (self.f11) else p0)
                lhs334_ = globalState
                lhs335_ = p3
                lhs336_ = default__.safeIndex(903, (p3).length(0))
                lhs337_ = p3
                lhs338_ = default__.safeIndex(191, (p3).length(0))
                lhs334_.f8 = rhs498_
                lhs335_[lhs336_] = rhs499_
                lhs337_[lhs338_] = rhs500_
                d_3146_v55_: _dafny.Seq
                d_3146_v55_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.fm2(globalState)), p1])
                d_3147_v56_: _dafny.MultiSet
                d_3147_v56_ = _dafny.MultiSet([(d_3132_v42_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhfpf"))), p2, default__.fm1(len(d_3146_v55_), True, d_3133_i0_, globalState)])
                d_3148_v57_: _dafny.Map
                d_3148_v57_ = _dafny.Map({p0: d_3147_v56_})
                d_3147_v56_ = (d_3147_v56_) | (((d_3148_v57_)[(p3)[default__.safeIndex(903, (p3).length(0))]] if ((p3)[default__.safeIndex(903, (p3).length(0))]) in (d_3148_v57_) else d_3147_v56_))
                r1 = (d_3133_i0_) - (d_3133_i0_)
            d_3149_v58_: _dafny.Set
            d_3149_v58_ = d_3130_v40_
            d_3149_v58_ = d_3149_v58_
            d_3150_v59_: _dafny.Map
            d_3150_v59_ = _dafny.Map({p1: p0})
            d_3151_v60_: _dafny.Set
            d_3151_v60_ = _dafny.Set({(p0 if p0 else p2), p0, ((d_3150_v59_)[p1] if (p1) in (d_3150_v59_) else self.f11)})
            d_3152_v61_: D0
            d_3152_v61_ = D0_DC3(p1, self.f11)
            d_3153_v62_: _dafny.MultiSet
            d_3153_v62_ = _dafny.MultiSet([p2, self.f11, p0, p0])
            d_3154_v63_: str
            d_3154_v63_ = _dafny.CodePoint('q')
            d_3151_v60_ = default__.fm4(d_3133_i0_, d_3152_v61_, ((d_3153_v62_)[self.f11] if (self.f11) in (d_3153_v62_) else d_3133_i0_), d_3154_v63_, globalState)
        d_3155_v64_: D20
        d_3155_v64_ = D20_DC53()
        d_3156_v65_: _dafny.Array
        nw474_ = _dafny.Array(None, 5)
        nw474_[int(0)] = d_3155_v64_
        nw474_[int(1)] = d_3155_v64_
        nw474_[int(2)] = d_3155_v64_
        nw474_[int(3)] = d_3155_v64_
        nw474_[int(4)] = d_3155_v64_
        d_3156_v65_ = nw474_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_3156_v65_).length(0)):
            d_3157_i1_: int = guard_loop_3_
            if (True) and (((0) <= (d_3157_i1_)) and ((d_3157_i1_) < ((d_3156_v65_).length(0)))):
                (d_3156_v65_)[(d_3157_i1_)] = d_3155_v64_
        d_3158_v66_: _dafny.Set
        d_3158_v66_ = _dafny.Set({p2})
        d_3159_v67_: _dafny.Map
        d_3159_v67_ = _dafny.Map({d_3158_v66_: p2})
        d_3160_v68_: str
        d_3160_v68_ = _dafny.CodePoint('y')
        d_3161_v69_: _dafny.Seq
        d_3161_v69_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_3162_v70_: _dafny.Map
        d_3162_v70_ = _dafny.Map({p2: d_3161_v69_})
        d_3163_v71_: _dafny.Map
        d_3163_v71_ = _dafny.Map({_dafny.CodePoint('k'): len((_dafny.SeqWithoutIsStrInference([d_3160_v68_ for d_3164_i2_ in range(default__.abs(72))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_3160_v68_ for d_3165_i2_ in range(default__.abs(72))]))), d_3160_v68_))})
        d_3166_v72_: _dafny.Array
        nw475_ = _dafny.Array(None, 23)
        nw475_[int(0)] = 834
        nw475_[int(1)] = p1
        nw475_[int(2)] = p1
        nw475_[int(3)] = (p1 if default__.fm1(p1, p0, p1, globalState) else p1)
        nw475_[int(4)] = 969
        nw475_[int(5)] = p1
        nw475_[int(6)] = p1
        nw475_[int(7)] = default__.safeDivisionInt(len(d_3132_v42_), (0) - (p1))
        nw475_[int(8)] = p1
        nw475_[int(9)] = (p1) - (default__.fm2(globalState))
        nw475_[int(10)] = (0) - (len(d_3159_v67_))
        nw475_[int(11)] = p1
        nw475_[int(12)] = (p1 if p0 else p1)
        nw475_[int(13)] = 727
        nw475_[int(14)] = len(_dafny.SeqWithoutIsStrInference([d_3160_v68_, d_3160_v68_]))
        nw475_[int(15)] = len(d_3132_v42_)
        nw475_[int(16)] = len(d_3162_v70_)
        nw475_[int(17)] = (p1) - (p1)
        nw475_[int(18)] = p1
        nw475_[int(19)] = (p1) + (p1)
        nw475_[int(20)] = len(d_3132_v42_)
        nw475_[int(21)] = p1
        nw475_[int(22)] = len(d_3163_v71_)
        d_3166_v72_ = nw475_
        d_3166_v72_ = d_3166_v72_
        (self).f11 = (d_3130_v40_).issubset(d_3130_v40_)
        r0 = (p1) > (p1)
        r1 = p1
        r2 = d_3160_v68_
        d_3167_v73_: T0
        nw476_ = C4()
        nw476_.ctor__(p1, self.f11)
        d_3167_v73_ = nw476_
        d_3168_v74_: _dafny.Set
        d_3168_v74_ = _dafny.Set({d_3167_v73_, d_3167_v73_})
        d_3169_v75_: _dafny.Seq
        d_3169_v75_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_3167_v73_})])
        r3 = (d_3168_v74_).isdisjoint((d_3169_v75_)[default__.safeIndex(p1, len(d_3169_v75_))])
        return r0, r1, r2, r3

    def m2(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_3170_v0_: int
        d_3170_v0_ = 585
        hi15_ = d_3170_v0_
        for d_3171_i0_ in range(d_3170_v0_, hi15_):
            (globalState).f8 = self.f11
            d_3172_v1_: str
            d_3172_v1_ = _dafny.CodePoint('b')
            d_3173_v2_: _dafny.Map
            d_3173_v2_ = _dafny.Map({self.f11: _dafny.CodePoint('p')})
            d_3174_v3_: _dafny.Set
            d_3174_v3_ = _dafny.Set({_dafny.Map({self.f11: d_3172_v1_}), d_3173_v2_, d_3173_v2_, d_3173_v2_})
            d_3175_v4_: _dafny.MultiSet
            d_3175_v4_ = _dafny.MultiSet([((d_3173_v2_).set(self.f11, d_3172_v1_)).set(self.f11, d_3172_v1_), d_3173_v2_, d_3173_v2_, (_dafny.Map({self.f11: d_3172_v1_})).set(self.f11, d_3172_v1_), d_3173_v2_])
            def iife279_():
                coll117_ = _dafny.Set()
                compr_117_: _dafny.Map
                for compr_117_ in (d_3175_v4_).Elements:
                    d_3176_v5_: _dafny.Map = compr_117_
                    if (d_3176_v5_) in (d_3175_v4_):
                        coll117_ = coll117_.union(_dafny.Set([d_3176_v5_]))
                return _dafny.Set(coll117_)
            (globalState).f8 = (d_3174_v3_) != ((_dafny.Set({_dafny.Map({self.f11: d_3172_v1_})})) | (iife279_()
            ))
            d_3177_v6_: _dafny.Array
            d_3178_v7_: bool
            out51_: _dafny.Array
            out52_: bool
            out51_, out52_ = default__.m0(globalState)
            d_3177_v6_ = out51_
            d_3178_v7_ = out52_
            rhs501_ = d_3178_v7_
            rhs502_ = (d_3171_i0_) - (d_3170_v0_)
            lhs339_ = globalState
            lhs339_.f8 = rhs501_
            r2 = rhs502_
        d_3179_v8_: str
        d_3179_v8_ = _dafny.CodePoint('k')
        d_3180_v9_: _dafny.MultiSet
        d_3180_v9_ = _dafny.MultiSet([d_3179_v8_])
        r2 = default__.safeModuloInt(default__.fm2(globalState), ((d_3180_v9_)).cardinality)
        d_3181_i1_: int
        d_3181_i1_ = 0
        with _dafny.label("17"):
            while self.f11:
                with _dafny.c_label("17"):
                    if (d_3181_i1_) >= (100):
                        raise _dafny.Break("17")
                    d_3181_i1_ = (d_3181_i1_) + (1)
                    d_3170_v0_ = d_3170_v0_
                    if self.f11:
                        r1 = (self.f11) == ((d_3170_v0_) != (465))
                        d_3182_v10_: D19
                        d_3182_v10_ = D19_DC51(d_3170_v0_, d_3170_v0_)
                        d_3183_v11_: C7
                        nw477_ = C7()
                        nw477_.ctor__(((d_3182_v10_).cf77 if True else d_3170_v0_), (d_3170_v0_) + (d_3170_v0_))
                        d_3183_v11_ = nw477_
                        d_3184_v12_: _dafny.Array
                        nw478_ = _dafny.Array(False, 22)
                        d_3184_v12_ = nw478_
                        index470_ = default__.safeIndex(636, (d_3184_v12_).length(0))
                        (d_3184_v12_)[index470_] = self.f11
                        d_3185_v13_: _dafny.Seq
                        d_3185_v13_ = _dafny.SeqWithoutIsStrInference([d_3184_v12_])
                        d_3186_v14_: _dafny.Map
                        d_3186_v14_ = _dafny.Map({d_3170_v0_: (len(d_3185_v13_)) * (d_3170_v0_)})
                        d_3187_v15_: _dafny.Seq
                        d_3187_v15_ = _dafny.SeqWithoutIsStrInference([self.f11])
                        d_3188_v16_: _dafny.Seq
                        d_3188_v16_ = _dafny.SeqWithoutIsStrInference([(d_3187_v15_) + (d_3187_v15_), (d_3187_v15_) + (d_3187_v15_), d_3187_v15_, d_3187_v15_, d_3187_v15_])
                        d_3189_v17_: _dafny.Seq
                        d_3189_v17_ = _dafny.SeqWithoutIsStrInference([d_3188_v16_])
                        index471_ = default__.safeIndex(636, (d_3184_v12_).length(0))
                        rhs503_ = self.f11
                        rhs504_ = d_3186_v14_
                        rhs505_ = ((d_3188_v16_) + (_dafny.SeqWithoutIsStrInference([d_3187_v15_ for d_3190_i2_ in range(default__.abs(340))]))) + ((d_3189_v17_)[default__.safeIndex(d_3170_v0_, len(d_3189_v17_))])
                        lhs340_ = d_3184_v12_
                        lhs341_ = default__.safeIndex(636, (d_3184_v12_).length(0))
                        lhs340_[lhs341_] = rhs503_
                        d_3186_v14_ = rhs504_
                        d_3188_v16_ = rhs505_
                        d_3170_v0_ = default__.safeModuloInt((0) - (d_3170_v0_), 424)
                        d_3191_v18_: _dafny.Array
                        def lambda157_(d_3192_i3_):
                            return D18_DC48()

                        init87_ = lambda157_
                        nw479_ = _dafny.Array(None, 16)
                        for i0_87_ in range(nw479_.length(0)):
                            nw479_[i0_87_] = init87_(i0_87_)
                        d_3191_v18_ = nw479_
                        index472_ = default__.safeIndex(200, (d_3191_v18_).length(0))
                        (d_3191_v18_)[index472_] = default__.fm64(globalState)
                        index473_ = default__.safeIndex(200, (d_3191_v18_).length(0))
                        (d_3191_v18_)[index473_] = default__.fm64(globalState)
                    elif True:
                        d_3193_v19_: _dafny.Set
                        d_3193_v19_ = _dafny.Set({-628, d_3170_v0_, d_3170_v0_, d_3170_v0_, d_3170_v0_})
                        d_3194_v20_: _dafny.Map
                        d_3194_v20_ = _dafny.Map({d_3193_v19_: d_3179_v8_})
                        d_3195_v21_: _dafny.Map
                        d_3195_v21_ = _dafny.Map({d_3170_v0_: self.f11})
                        (globalState).f8 = not(default__.fm1(len((d_3194_v20_) | (d_3194_v20_)), (((d_3195_v21_)[d_3170_v0_] if (d_3170_v0_) in (d_3195_v21_) else (self).fm6(self.f11, d_3170_v0_, (0) - (d_3170_v0_), globalState)) if self.f11 else self.f11), d_3170_v0_, globalState))
                        d_3196_v22_: _dafny.Seq
                        d_3196_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eintml"))
                        d_3197_v23_: _dafny.Seq
                        d_3197_v23_ = _dafny.SeqWithoutIsStrInference([(d_3196_v22_) + (d_3196_v22_)])
                        d_3196_v22_ = (d_3197_v23_)[default__.safeIndex(d_3170_v0_, len(d_3197_v23_))]
                        d_3198_v24_: C3
                        nw480_ = C3()
                        nw480_.ctor__(d_3170_v0_, default__.safeModuloInt(d_3170_v0_, d_3170_v0_))
                        d_3198_v24_ = nw480_
                        d_3199_v25_: _dafny.Seq
                        d_3199_v25_ = _dafny.SeqWithoutIsStrInference([d_3193_v19_])
                        rhs506_ = default__.safeDivisionInt(d_3170_v0_, d_3170_v0_)
                        rhs507_ = d_3196_v22_
                        rhs508_ = self.f11
                        rhs509_ = d_3198_v24_
                        rhs510_ = ((d_3199_v25_) + (d_3199_v25_)).set(default__.safeIndex(d_3170_v0_, len((d_3199_v25_) + (d_3199_v25_))), d_3193_v19_)
                        r2 = rhs506_
                        d_3196_v22_ = rhs507_
                        r0 = rhs508_
                        d_3198_v24_ = rhs509_
                        d_3199_v25_ = rhs510_
                        d_3200_v26_: _dafny.Seq
                        d_3200_v26_ = _dafny.SeqWithoutIsStrInference([585, (0) - (d_3170_v0_)])
                        d_3201_v27_: C5
                        nw481_ = C5()
                        nw481_.ctor__((d_3200_v26_)[default__.safeIndex(len(_dafny.Map({d_3179_v8_: d_3170_v0_})), len(d_3200_v26_))])
                        d_3201_v27_ = nw481_
                        d_3179_v8_ = default__.fm25((d_3201_v27_).f24, d_3170_v0_, globalState)
                    d_3170_v0_ = d_3170_v0_
                    (globalState).f8 = self.f11
                    pass
            pass
        d_3202_v28_: _dafny.Set
        d_3202_v28_ = _dafny.Set({d_3170_v0_})
        d_3203_v29_: _dafny.Set
        d_3203_v29_ = d_3202_v28_
        d_3204_v30_: _dafny.Array
        nw482_ = _dafny.Array(False, 1)
        d_3204_v30_ = nw482_
        d_3205_v31_: _dafny.Set
        d_3205_v31_ = _dafny.Set({d_3204_v30_})
        d_3206_v32_: _dafny.MultiSet
        d_3206_v32_ = _dafny.MultiSet([len(d_3205_v31_)])
        d_3207_v33_: D5
        d_3207_v33_ = D5_DC15(d_3170_v0_, d_3170_v0_, len(_dafny.Set({d_3170_v0_})))
        d_3208_v34_: _dafny.Map
        d_3208_v34_ = _dafny.Map({(0) - ((d_3206_v32_).cardinality): d_3207_v33_})
        d_3209_v35_: _dafny.Seq
        d_3209_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idy"))
        d_3210_v36_: _dafny.Map
        d_3210_v36_ = _dafny.Map({d_3170_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roaishuns"))})
        d_3203_v29_ = default__.fm65(D6_DC16(d_3208_v34_), d_3209_v35_, d_3210_v36_, globalState)
        d_3211_i4_: int
        d_3211_i4_ = 0
        with _dafny.label("18"):
            while default__.fm1(d_3170_v0_, self.f11, d_3170_v0_, globalState):
                with _dafny.c_label("18"):
                    if (d_3211_i4_) >= (100):
                        raise _dafny.Break("18")
                    d_3211_i4_ = (d_3211_i4_) + (1)
                    d_3212_v37_: _dafny.Map
                    d_3212_v37_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kupca")): d_3204_v30_})
                    d_3213_v38_: _dafny.Map
                    d_3213_v38_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_3179_v8_ for d_3214_i5_ in range(default__.abs(508))]): d_3212_v37_})
                    d_3215_v39_: D24
                    d_3215_v39_ = D24_DC60(d_3212_v37_)
                    d_3216_v40_: _dafny.Seq
                    d_3216_v40_ = _dafny.SeqWithoutIsStrInference([d_3212_v37_, d_3212_v37_, d_3212_v37_])
                    d_3217_v41_: _dafny.Array
                    nw483_ = _dafny.Array(None, 29)
                    nw483_[int(0)] = d_3212_v37_
                    nw483_[int(1)] = d_3212_v37_
                    nw483_[int(2)] = d_3212_v37_
                    nw483_[int(3)] = d_3212_v37_
                    nw483_[int(4)] = d_3212_v37_
                    nw483_[int(5)] = d_3212_v37_
                    nw483_[int(6)] = (d_3212_v37_) | (d_3212_v37_)
                    nw483_[int(7)] = d_3212_v37_
                    nw483_[int(8)] = d_3212_v37_
                    nw483_[int(9)] = ((d_3213_v38_)[default__.fm36(d_3170_v0_, d_3170_v0_, d_3179_v8_, globalState)] if (default__.fm36(d_3170_v0_, d_3170_v0_, d_3179_v8_, globalState)) in (d_3213_v38_) else d_3212_v37_)
                    nw483_[int(10)] = (_dafny.Map({d_3209_v35_: d_3204_v30_})) | (d_3212_v37_)
                    nw483_[int(11)] = d_3212_v37_
                    nw483_[int(12)] = d_3212_v37_
                    nw483_[int(13)] = d_3212_v37_
                    nw483_[int(14)] = d_3212_v37_
                    nw483_[int(15)] = (d_3212_v37_) | (d_3212_v37_)
                    nw483_[int(16)] = d_3212_v37_
                    nw483_[int(17)] = (d_3212_v37_) | (d_3212_v37_)
                    nw483_[int(18)] = (d_3215_v39_).cf88
                    nw483_[int(19)] = ((d_3215_v39_).cf88 if not(self.f11) else d_3212_v37_)
                    nw483_[int(20)] = d_3212_v37_
                    nw483_[int(21)] = (d_3212_v37_) | (d_3212_v37_)
                    nw483_[int(22)] = (d_3216_v40_)[default__.safeIndex(d_3170_v0_, len(d_3216_v40_))]
                    nw483_[int(23)] = d_3212_v37_
                    nw483_[int(24)] = _dafny.Map({d_3209_v35_: d_3204_v30_})
                    nw483_[int(25)] = d_3212_v37_
                    nw483_[int(26)] = d_3212_v37_
                    nw483_[int(27)] = (d_3212_v37_) | (d_3212_v37_)
                    nw483_[int(28)] = (d_3212_v37_ if self.f11 else d_3212_v37_)
                    d_3217_v41_ = nw483_
                    index474_ = default__.safeIndex(4, (d_3217_v41_).length(0))
                    (d_3217_v41_)[index474_] = d_3212_v37_
                    index475_ = default__.safeIndex(4, (d_3217_v41_).length(0))
                    (d_3217_v41_)[index475_] = d_3212_v37_
                    d_3218_v42_: _dafny.MultiSet
                    d_3218_v42_ = _dafny.MultiSet([self.f11, self.f11])
                    d_3219_v43_: _dafny.Seq
                    d_3219_v43_ = _dafny.SeqWithoutIsStrInference([self.f11, (_dafny.MultiSet([self.f11])).issubset(d_3218_v42_), self.f11, False])
                    d_3219_v43_ = (d_3219_v43_) + (d_3219_v43_)
                    if self.f11:
                        d_3220_v44_: C14
                        nw484_ = C14()
                        nw484_.ctor__(self.f11)
                        d_3220_v44_ = nw484_
                        (globalState).f8 = self.f11
                        d_3221_v45_: _dafny.Array
                        nw485_ = _dafny.Array(_dafny.Map({}), 18)
                        d_3221_v45_ = nw485_
                        d_3222_v46_: C15
                        nw486_ = C15()
                        nw486_.ctor__(self.f11, d_3170_v0_, self.f11)
                        d_3222_v46_ = nw486_
                        d_3223_v47_: _dafny.Map
                        d_3223_v47_ = _dafny.Map({d_3222_v46_: d_3218_v42_})
                        index476_ = default__.safeIndex(890, (d_3221_v45_).length(0))
                        (d_3221_v45_)[index476_] = d_3223_v47_
                        index477_ = default__.safeIndex(890, (d_3221_v45_).length(0))
                        (d_3221_v45_)[index477_] = d_3223_v47_
                        d_3224_v48_: _dafny.Array
                        nw487_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                        d_3224_v48_ = nw487_
                        d_3225_v49_: _dafny.Map
                        d_3225_v49_ = _dafny.Map({(d_3224_v48_) in (_dafny.Map({d_3224_v48_: d_3222_v46_.f17})): d_3170_v0_})
                        d_3225_v49_ = (d_3225_v49_).set((self.f11) == ((self).fm6(default__.fm1(d_3170_v0_, default__.fm1(d_3170_v0_, True, d_3222_v46_.f17, globalState), d_3170_v0_, globalState), d_3170_v0_, (d_3206_v32_).cardinality, globalState)), (0) - (d_3170_v0_))
                        (self).f11 = self.f11
                    elif True:
                        d_3226_v50_: C6
                        nw488_ = C6()
                        nw488_.ctor__()
                        d_3226_v50_ = nw488_
                        d_3226_v50_ = d_3226_v50_
                        rhs511_ = self.f11
                        rhs512_ = d_3170_v0_
                        lhs342_ = self
                        lhs342_.f11 = rhs511_
                        d_3170_v0_ = rhs512_
                        d_3227_v51_: _dafny.Map
                        d_3227_v51_ = _dafny.Map({not(self.f11): d_3170_v0_})
                        d_3228_v52_: _dafny.Map
                        d_3228_v52_ = _dafny.Map({d_3170_v0_: self.f11})
                        d_3229_v53_: _dafny.Set
                        d_3229_v53_ = _dafny.Set({self.f11})
                        d_3230_v54_: _dafny.Array
                        nw489_ = _dafny.Array(None, 20)
                        nw489_[int(0)] = len((d_3209_v35_) + ((_dafny.SeqWithoutIsStrInference([d_3179_v8_ for d_3231_i6_ in range(default__.abs(24))])).set(default__.safeIndex(d_3170_v0_, len(_dafny.SeqWithoutIsStrInference([d_3179_v8_ for d_3232_i6_ in range(default__.abs(24))]))), d_3179_v8_)))
                        nw489_[int(1)] = d_3170_v0_
                        nw489_[int(2)] = d_3170_v0_
                        nw489_[int(3)] = d_3170_v0_
                        nw489_[int(4)] = default__.safeDivisionInt(((d_3206_v32_)[d_3170_v0_] if (d_3170_v0_) in (d_3206_v32_) else d_3170_v0_), d_3170_v0_)
                        nw489_[int(5)] = (d_3218_v42_).cardinality
                        nw489_[int(6)] = (396) + (d_3170_v0_)
                        nw489_[int(7)] = d_3170_v0_
                        nw489_[int(8)] = default__.fm2(globalState)
                        nw489_[int(9)] = d_3170_v0_
                        nw489_[int(10)] = ((d_3227_v51_)[((d_3228_v52_)[d_3170_v0_] if (d_3170_v0_) in (d_3228_v52_) else not(self.f11))] if (((d_3228_v52_)[d_3170_v0_] if (d_3170_v0_) in (d_3228_v52_) else not(self.f11))) in (d_3227_v51_) else d_3170_v0_)
                        nw489_[int(11)] = default__.fm2(globalState)
                        nw489_[int(12)] = d_3170_v0_
                        nw489_[int(13)] = len(default__.fm0(globalState))
                        nw489_[int(14)] = d_3170_v0_
                        nw489_[int(15)] = ((0) - (d_3170_v0_)) - (len(d_3229_v53_))
                        nw489_[int(16)] = d_3170_v0_
                        nw489_[int(17)] = len((default__.fm36(d_3170_v0_, d_3170_v0_, _dafny.CodePoint('c'), globalState)) + (d_3209_v35_))
                        nw489_[int(18)] = d_3170_v0_
                        nw489_[int(19)] = (d_3206_v32_).cardinality
                        d_3230_v54_ = nw489_
                        index478_ = default__.safeIndex(676, (d_3230_v54_).length(0))
                        (d_3230_v54_)[index478_] = d_3170_v0_
                        index479_ = default__.safeIndex(676, (d_3230_v54_).length(0))
                        (d_3230_v54_)[index479_] = d_3170_v0_
                        (globalState).f8 = (d_3170_v0_) > (len(d_3227_v51_))
                        d_3233_v55_: _dafny.Seq
                        d_3233_v55_ = _dafny.SeqWithoutIsStrInference([(181) * (len(d_3227_v51_)), (d_3230_v54_)[default__.safeIndex(676, (d_3230_v54_).length(0))], (d_3230_v54_)[default__.safeIndex(676, (d_3230_v54_).length(0))]])
                        (globalState).f5 = d_3233_v55_
                    d_3170_v0_ = d_3170_v0_
                    pass
            pass
        d_3234_i7_: int
        d_3234_i7_ = 0
        with _dafny.label("19"):
            while (d_3170_v0_) > (d_3170_v0_):
                with _dafny.c_label("19"):
                    if (d_3234_i7_) >= (100):
                        raise _dafny.Break("19")
                    d_3234_i7_ = (d_3234_i7_) + (1)
                    d_3235_v56_: _dafny.Seq
                    d_3235_v56_ = _dafny.SeqWithoutIsStrInference([not(self.f11)])
                    d_3236_v57_: C16
                    nw490_ = C16()
                    nw490_.ctor__(self.f11, d_3170_v0_, False, d_3170_v0_, d_3170_v0_)
                    d_3236_v57_ = nw490_
                    d_3237_v58_: _dafny.Seq
                    d_3237_v58_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_3236_v57_: self.f11}))])
                    rhs513_ = (d_3235_v56_) <= (_dafny.SeqWithoutIsStrInference([True]))
                    rhs514_ = ((d_3170_v0_) + (d_3170_v0_)) == (((0) - (len(d_3237_v58_))) * ((d_3236_v57_).f15))
                    rhs515_ = self.f11
                    rhs516_ = (len(d_3237_v58_)) - (d_3170_v0_)
                    lhs343_ = globalState
                    r1 = rhs513_
                    r1 = rhs514_
                    lhs343_.f8 = rhs515_
                    r2 = rhs516_
                    d_3238_v59_: _dafny.Seq
                    d_3238_v59_ = _dafny.SeqWithoutIsStrInference([d_3209_v35_])
                    d_3210_v36_ = (d_3210_v36_).set(d_3170_v0_, (d_3238_v59_)[default__.safeIndex(d_3170_v0_, len(d_3238_v59_))])
                    d_3239_v60_: _dafny.Map
                    d_3239_v60_ = _dafny.Map({d_3170_v0_: d_3170_v0_})
                    d_3240_v61_: _dafny.Seq
                    d_3240_v61_ = (d_3237_v58_).set(default__.safeIndex(d_3170_v0_, len(d_3237_v58_)), len(d_3239_v60_))
                    d_3241_v62_: _dafny.Set
                    d_3241_v62_ = _dafny.Set({d_3240_v61_, d_3240_v61_, d_3240_v61_})
                    d_3242_v63_: C9
                    nw491_ = C9()
                    nw491_.ctor__(d_3241_v62_, 908, d_3170_v0_, len(d_3209_v35_), (d_3236_v57_).f14, d_3209_v35_)
                    d_3242_v63_ = nw491_
                    d_3243_v64_: _dafny.Set
                    d_3243_v64_ = _dafny.Set({d_3242_v63_, d_3242_v63_, d_3242_v63_, d_3242_v63_})
                    (self).f11 = (d_3243_v64_) == ((d_3243_v64_) | (d_3243_v64_))
                    d_3244_v65_: C6
                    nw492_ = C6()
                    nw492_.ctor__()
                    d_3244_v65_ = nw492_
                    pass
            pass
        r0 = (default__.safeModuloInt(d_3170_v0_, len(d_3209_v35_))) < (d_3170_v0_)
        r1 = ((d_3170_v0_) - (d_3170_v0_)) >= (905)
        r2 = (default__.fm2(globalState)) - (d_3170_v0_)
        return r0, r1, r2

