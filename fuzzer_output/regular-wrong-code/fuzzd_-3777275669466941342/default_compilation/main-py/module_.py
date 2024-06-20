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
    def fm2(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (-456)])) + (_dafny.SeqWithoutIsStrInference([(D23_DC64((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tk")))), len(_dafny.SeqWithoutIsStrInference([66])), True, len(_dafny.Set({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlogvbcg")): (_dafny.MultiSet([True])).cardinality})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))})), len(_dafny.Map({True: _dafny.CodePoint('s')})))).cf94]))) + ((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(False)])).cardinality, 691])) + (_dafny.SeqWithoutIsStrInference([578, 706, 887, (_dafny.MultiSet([True])).cardinality])))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (0) - (len(_dafny.Map({not((_dafny.CodePoint('a')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkoew")))): len((_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])})))})))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        if False:
            return D2_DC5(True)
        elif True:
            return D2_DC5(not(True))

    @staticmethod
    def fm8(globalState):
        return (0) - (((0) - ((-640 if True else len(_dafny.Map({False: True})))) if False else default__.safeModuloInt(840, 273)))

    @staticmethod
    def fm9(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.Map({-872: 407})).keys.Elements:
                d_0_v0_: int = compr_0_
                if (d_0_v0_) in (_dafny.Map({-872: 407})):
                    coll0_ = coll0_.union(_dafny.Set([(d_0_v0_) - (692)]))
            return _dafny.Set(coll0_)
        return ((_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([170]))), -314, len(_dafny.Map({_dafny.CodePoint('r'): 438}))})).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))}))) - ((_dafny.Set({-324})).intersection(iife0_()
        ))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm12(p0, p1, globalState):
        return (0) - ((len(_dafny.Map({_dafny.MultiSet([-191]): _dafny.CodePoint('b')}))) * ((0) - ((852) + ((0) - (len(_dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')])): len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([True]))}))}))})))))))

    @staticmethod
    def fm13(globalState):
        return (642) <= (((_dafny.MultiSet([D32_DC88()])).cardinality) - (-489))

    @staticmethod
    def fm14(globalState):
        return ((_dafny.MultiSet([(0) - (len(_dafny.Set({True}))), 20])) | (_dafny.MultiSet([666]))).intersection(_dafny.MultiSet([976, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1_i0_ in range(default__.abs(195))])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))).cardinality, -915, 66]))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        if ((D17_DC44(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2_i0_ in range(default__.abs(462))]))).cf62) in (_dafny.Set({not(True)})):
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: _dafny.MultiSet
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([730]) for d_3_i1_ in range(default__.abs(425))])).Elements:
                    d_4_v0_: _dafny.MultiSet = compr_1_
                    if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([730]) for d_3_i1_ in range(default__.abs(425))])):
                        coll1_[d_4_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gylpbdqx"))
                return _dafny.Map(coll1_)
            return (iife1_()
            ) | (_dafny.Map({_dafny.MultiSet([-387, -726]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkyku"))}))
        elif True:
            return _dafny.Map({_dafny.MultiSet([766]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psfuts"))})

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return ((len(_dafny.Map({True: True})) if True else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_5_i0_ in range(default__.abs(953))])))) <= (len((D17_DC44(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dk")))).cf63))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(254, 90):
                d_7_v0_: int = compr_2_
                if ((254) <= (d_7_v0_)) and ((d_7_v0_) < (90)):
                    coll2_[(d_7_v0_) + (643)] = 149
            return _dafny.Map(coll2_)
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(409, 475):
                    d_9_v2_: int = compr_5_
                    if ((409) <= (d_9_v2_)) and ((d_9_v2_) < (475)):
                        coll5_[(d_9_v2_) + (-664)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_10_i3_ in range(default__.abs(360))]))
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(409, 475):
                    d_9_v2_: int = compr_4_
                    if ((409) <= (d_9_v2_)) and ((d_9_v2_) < (475)):
                        coll4_[(d_9_v2_) + (-664)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_10_i3_ in range(default__.abs(360))]))
                return _dafny.Map(coll4_)
            compr_3_: int
            for compr_3_ in (_dafny.Set({len(iife4_()
            )})).Elements:
                d_11_v1_: int = compr_3_
                if (d_11_v1_) in (_dafny.Set({len(iife5_()
                )})):
                    coll3_[(d_11_v1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmttox"))))] = False
            return _dafny.Map(coll3_)
        return (_dafny.SeqWithoutIsStrInference([(D22_DC60(863, True, False, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_6_i0_ in range(default__.abs(404))]))])): False}))]))).cf86, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(_dafny.MultiSet([len(iife2_()
)])).cardinality: _dafny.Map({not(False): -450})})) for d_8_i1_ in range(default__.abs(368))]))).cardinality, (_dafny.MultiSet([_dafny.Set({-551, len(_dafny.SeqWithoutIsStrInference([len(iife3_()
) for d_12_i2_ in range(default__.abs(499))]))})])).cardinality])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_13_i4_ in range(default__.abs(91))])), (_dafny.MultiSet([12])).cardinality]))

    @staticmethod
    def fm20(p0, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsgjti")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_14_i0_ in range(default__.abs(368))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boxkrn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk"))])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrtacxhux")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqnbx")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_15_i1_ in range(default__.abs(113))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdx"))]))

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([17]), _dafny.SeqWithoutIsStrInference([480])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({-678: _dafny.CodePoint('e')})): _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxootepfd")))})}))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([26, 95]))]) for d_16_i0_ in range(default__.abs(-311))])))

    @staticmethod
    def fm24(p0, p1, globalState):
        return _dafny.MultiSet([(len(_dafny.Map({False: D31_DC84(_dafny.Map({610: 210}))}))) + ((0) - (-930))])

    @staticmethod
    def fm25(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([-777])).Elements:
                d_17_v0_: int = compr_6_
                if (d_17_v0_) in (_dafny.SeqWithoutIsStrInference([-777])):
                    coll6_[default__.safeModuloInt(d_17_v0_, 980)] = -731
            return _dafny.Map(coll6_)
        return (len((_dafny.Map({198: 528})) | (iife6_()
        ))) + (-769)

    @staticmethod
    def fm26(p0, globalState):
        return (D4_DC10(False, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: _dafny.CodePoint('p')})])), 563)).cf13

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        return default__.safeDivisionInt(default__.safeDivisionInt(len(_dafny.Map({len(_dafny.Map({len(_dafny.Map({True: False})): True})): False})), 10), len((_dafny.Map({False: 424})) | (_dafny.Map({(D30_DC83(True)).cf132: 552}))))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        source0_ = (D12_DC33(D12_DC31(_dafny.Set({51, len(_dafny.SeqWithoutIsStrInference([True]))}))) if False else D12_DC33(D12_DC31(_dafny.Set({-779, len(_dafny.SeqWithoutIsStrInference([False]))}))))
        if source0_.is_DC32:
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_18_i0_ in range(default__.abs(865))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jssfxmpcl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdgwhlntb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uebuddcnv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbepvkl"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_19_i1_ in range(default__.abs(559))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkrixrc"))]))
        elif source0_.is_DC31:
            d_20___mcc_h0_ = source0_.cf48
            d_21_cf48_ = d_20___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jeaff"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_22_i2_ in range(default__.abs(840))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxoqh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alotwajny"))]))
        elif True:
            d_23___mcc_h1_ = source0_.cf49
            d_24_cf49_ = d_23___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdut"))])

    @staticmethod
    def fm29(p0, p1, globalState):
        return ((_dafny.Map({True: False})) | (_dafny.Map({False: True}))) not in ((_dafny.Map({_dafny.Map({False: not(False)}): len(_dafny.Map({len(_dafny.Map({False: True})): 919}))})) | ((D33_DC89(_dafny.Map({_dafny.Map({False: not(False)}): -485}))).cf139))

    @staticmethod
    def fm35(globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: str
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('n')])).Elements:
                d_25_v0_: str = compr_7_
                if (d_25_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('n')])):
                    coll7_[d_25_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nydoscix"))
            return _dafny.Map(coll7_)
        return _dafny.Set({((0) - (-850)) + (-176), len(_dafny.Map({len(iife7_()
        ): 970})), len((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_26_i0_ in range(default__.abs(719))])): _dafny.CodePoint('q')})])), (0) - (-695)})) - (_dafny.Set({276})))})

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        return _dafny.Set({-582})

    @staticmethod
    def fm38(globalState):
        return ((_dafny.MultiSet([False])).ispropersubset(_dafny.MultiSet([True, True]))) == (False)

    @staticmethod
    def fm39(p0, globalState):
        if not(False):
            return D4_DC9(_dafny.MultiSet([1, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gt"))))]))
        elif True:
            return D4_DC9(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([470])).cardinality, 236])), 785, -853]))

    @staticmethod
    def fm40(globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mj")): len(_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btxruxdpk"))))))]))}) if True else _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sw")): -744}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xityw")): 258}))

    @staticmethod
    def fm41(p0, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_27_i0_ in range(default__.abs(498))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('o')])))

    @staticmethod
    def fm42(p0, p1, globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        if False:
            return D8_DC21(False)
        elif True:
            return D8_DC21(True)

    @staticmethod
    def fm44(p0, globalState):
        if True:
            return D5_DC13(D2_DC4(_dafny.CodePoint('m')))
        elif True:
            return D5_DC13(D2_DC4(_dafny.CodePoint('q')))

    @staticmethod
    def fm45(globalState):
        def iife8_():
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(361, 71):
                    d_28_v1_: int = compr_10_
                    if ((361) <= (d_28_v1_)) and ((d_28_v1_) < (71)):
                        coll10_[default__.safeDivisionInt(d_28_v1_, (_dafny.MultiSet([True, False])).cardinality)] = (0) - (len(_dafny.Map({True: False})))
                return _dafny.Map(coll10_)
            coll8_ = _dafny.Map()
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(361, 71):
                    d_28_v1_: int = compr_9_
                    if ((361) <= (d_28_v1_)) and ((d_28_v1_) < (71)):
                        coll9_[default__.safeDivisionInt(d_28_v1_, (_dafny.MultiSet([True, False])).cardinality)] = (0) - (len(_dafny.Map({True: False})))
                return _dafny.Map(coll9_)
            compr_8_: int
            for compr_8_ in (iife9_()
            ).keys.Elements:
                d_29_v0_: int = compr_8_
                if (d_29_v0_) in (iife10_()
                ):
                    coll8_[default__.safeDivisionInt(d_29_v0_, -298)] = False
            return _dafny.Map(coll8_)
        return (D10_DC26(iife8_()
)).cf39

    @staticmethod
    def fm46(p0, globalState):
        source1_ = (D20_DC53(D20_DC52()) if False else D20_DC53(D20_DC50(_dafny.SeqWithoutIsStrInference([D10_DC29(D10_DC27(False, False, False, -387, True)) for d_30_i0_ in range(default__.abs(419))]))))
        if source1_.is_DC51:
            d_31___mcc_h0_ = source1_.cf75
            d_32___mcc_h1_ = source1_.cf76
            d_33___mcc_h2_ = source1_.cf77
            d_34_cf77_ = d_33___mcc_h2_
            d_35_cf76_ = d_32___mcc_h1_
            d_36_cf75_ = d_31___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_34_cf77_]), _dafny.SeqWithoutIsStrInference([d_35_cf76_])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_35_cf76_, d_34_cf77_]) for d_37_i1_ in range(default__.abs(61))])))
        elif source1_.is_DC52:
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([not(False)])]))
        elif source1_.is_DC50:
            d_38___mcc_h3_ = source1_.cf74
            d_39_cf74_ = d_38___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([not(not(False)), True, True]), _dafny.SeqWithoutIsStrInference([not(True)])])
        elif True:
            d_40___mcc_h4_ = source1_.cf78
            d_41_cf78_ = d_40___mcc_h4_
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])

    @staticmethod
    def fm47(globalState):
        return ((_dafny.Set({False, True, True, True})).intersection(_dafny.Set({False, True}))) - ((_dafny.Set({True})).intersection(_dafny.Set({False})))

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcbb")))])).Elements:
                d_42_v0_: int = compr_11_
                if (d_42_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcbb")))])):
                    coll11_[(d_42_v0_) + (54)] = False
            return _dafny.Map(coll11_)
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({False: _dafny.Set({653})})): not(True)}), _dafny.Map({660: True})])) + (_dafny.SeqWithoutIsStrInference([iife11_()
 for d_43_i0_ in range(default__.abs(193))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-93: (D2_DC5(False)).cf8}), _dafny.Map({(_dafny.MultiSet([-103, -347])).cardinality: False})]))

    @staticmethod
    def fm49(globalState):
        if (len(_dafny.Map({False: False}))) != (439):
            return D7_DC19(not(False), False)
        elif True:
            def iife12_():
                def iife14_():
                    coll14_ = _dafny.Set()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(847, 343):
                        d_47_v0_: int = compr_14_
                        if ((847) <= (d_47_v0_)) and ((d_47_v0_) < (343)):
                            coll14_ = coll14_.union(_dafny.Set([(d_47_v0_) + ((_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_45_i0_ in range(default__.abs(494))]))}))])).cardinality)]))
                    return _dafny.Set(coll14_)
                coll12_ = _dafny.Set()
                def iife13_():
                    coll13_ = _dafny.Set()
                    compr_13_: int
                    for compr_13_ in _dafny.IntegerRange(847, 343):
                        d_44_v0_: int = compr_13_
                        if ((847) <= (d_44_v0_)) and ((d_44_v0_) < (343)):
                            coll13_ = coll13_.union(_dafny.Set([(d_44_v0_) + ((_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_45_i0_ in range(default__.abs(494))]))}))])).cardinality)]))
                    return _dafny.Set(coll13_)
                compr_12_: int
                for compr_12_ in (iife13_()
                ).Elements:
                    d_46_v1_: int = compr_12_
                    if (d_46_v1_) in (iife14_()
                    ):
                        coll12_ = coll12_.union(_dafny.Set([(d_46_v1_) + (-58)]))
                return _dafny.Set(coll12_)
            return D7_DC19(True, not((D1_DC2(924, len(iife12_()
), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwlvydblv")), True)).cf5))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        return (D7_DC16(_dafny.MultiSet([D1_DC2(len(_dafny.SeqWithoutIsStrInference([169])), -20, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpihwnp")), True)]))).cf21

    @staticmethod
    def fm51(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvdk"))

    @staticmethod
    def fm54(p0, p1, p2, p3, globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in _dafny.IntegerRange(-464, 680):
                d_48_v0_: int = compr_15_
                if ((-464) <= (d_48_v0_)) and ((d_48_v0_) < (680)):
                    coll15_[default__.safeModuloInt(d_48_v0_, len(_dafny.Set({True, False, False, False, True})))] = 61
            return _dafny.Map(coll15_)
        return (_dafny.Map({-266: 368})) | (iife15_()
        )

    @staticmethod
    def fm55(p0, p1, globalState):
        return ((_dafny.Map({_dafny.CodePoint('i'): False}) if False else _dafny.Map({_dafny.CodePoint('j'): False}))) | ((_dafny.Map({_dafny.CodePoint('w'): False})) | (_dafny.Map({_dafny.CodePoint('u'): False})))

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        return ((_dafny.Map({not(True): (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlfxecaqd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thjwen")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tepskhntj"))])))})) | (_dafny.Map({True: -175}))) | (_dafny.Map({not(not(not(True))): len(_dafny.Set({_dafny.CodePoint('b')}))}))

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        if not(False):
            return _dafny.MultiSet([not(False)])
        elif True:
            return _dafny.MultiSet([not(not(False)), True, False])

    @staticmethod
    def fm58(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.CodePoint('v'), (_dafny.CodePoint('s') if False else _dafny.CodePoint('e')), _dafny.CodePoint('n')})

    @staticmethod
    def fm59(p0, p1, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({True: True}))) | ((_dafny.Map({False: True})) | (_dafny.Map({True: not(False)})))

    @staticmethod
    def fm60(p0, p1, p2, p3, globalState):
        return D13_DC35()

    @staticmethod
    def fm61(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([False])}))) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([True, True])}))

    @staticmethod
    def fm62(globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True)])) + ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm63(p0, p1, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: _dafny.Seq
            for compr_16_ in ((_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([457, len(_dafny.Set({834, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnpl")))}))]))).cardinality]): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([14]): _dafny.MultiSet([False, False])}))).keys.Elements:
                d_49_v0_: _dafny.Seq = compr_16_
                if (d_49_v0_) in ((_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([457, len(_dafny.Set({834, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnpl")))}))]))).cardinality]): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([14]): _dafny.MultiSet([False, False])}))):
                    coll16_[d_49_v0_] = 219
            return _dafny.Map(coll16_)
        return iife16_()
        

    @staticmethod
    def fm64(p0, p1, globalState):
        source2_ = D12_DC33(D12_DC32())
        if source2_.is_DC32:
            return _dafny.Map({(D22_DC59(not(True), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgpdy")), len(_dafny.SeqWithoutIsStrInference([False])))).cf85: _dafny.CodePoint('l')})
        elif source2_.is_DC31:
            d_50___mcc_h0_ = source2_.cf48
            d_51_cf48_ = d_50___mcc_h0_
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(746, 807):
                    d_52_v0_: int = compr_17_
                    if ((746) <= (d_52_v0_)) and ((d_52_v0_) < (807)):
                        coll17_[default__.safeModuloInt(d_52_v0_, 501)] = 59
                return _dafny.Map(coll17_)
            return _dafny.Map({len(iife17_()
            ): _dafny.CodePoint('o')})
        elif True:
            d_53___mcc_h1_ = source2_.cf49
            d_54_cf49_ = d_53___mcc_h1_
            return _dafny.Map({514: _dafny.CodePoint('l')})

    @staticmethod
    def fm65(p0, p1, globalState):
        return D6_DC15(not((len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((_dafny.MultiSet([True])).cardinality)) for d_55_i0_ in range(default__.abs(-372))]))}))) <= ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])), 392, 372, 494]))))))

    @staticmethod
    def fm66(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(False), False])) + (_dafny.SeqWithoutIsStrInference([True, False, not(False)]))

    @staticmethod
    def fm67(p0, globalState):
        return D7_DC16((_dafny.MultiSet([D1_DC2(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_56_i0_ in range(default__.abs(3))])), 430, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hj")), True)])) | (_dafny.MultiSet([D1_DC2((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_57_i1_ in range(default__.abs(324))]))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False}))])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edkowsg")), False)])))

    @staticmethod
    def fm68(p0, globalState):
        return D9_DC25(not((len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): 509}))) < (828)), 90, 592, (not(False)) and (False), (0) - ((0) - ((0) - (-701))))

    @staticmethod
    def fm69(globalState):
        return D15_DC41((D15_DC38(True, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_58_i0_ in range(default__.abs(202))])), len(_dafny.SeqWithoutIsStrInference([328, 381, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uutnpr")))]))) if not(True) else D15_DC41(D15_DC41(D15_DC38(True, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([553])), 351])).cardinality, 245)))))

    @staticmethod
    def fm70(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([D10_DC29(D10_DC29(D10_DC27(False, False, False, 249, False)))])) + (_dafny.SeqWithoutIsStrInference([D10_DC29(D10_DC28(_dafny.Map({False: True})))]))) + ((_dafny.SeqWithoutIsStrInference([D10_DC29(D10_DC28(_dafny.Map({False: True}))) for d_59_i0_ in range(default__.abs(674))])) + (_dafny.SeqWithoutIsStrInference([D10_DC29(D10_DC29(D10_DC28(_dafny.Map({True: True})))), D10_DC29(D10_DC26(_dafny.Map({517: False}))), D10_DC29(D10_DC27(False, True, True, -975, False)), D10_DC29(D10_DC29(D10_DC28(_dafny.Map({True: False})))), D10_DC29(D10_DC27(False, False, not(True), (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-718])])).cardinality, False))])))

    @staticmethod
    def fm71(p0, globalState):
        return (_dafny.Map({_dafny.CodePoint('t'): D10_DC28(_dafny.Map({False: False}))})) | (_dafny.Map({_dafny.CodePoint('g'): D10_DC28(_dafny.Map({True: False}))}))

    @staticmethod
    def fm72(p0, globalState):
        return _dafny.Map({(773) == (len(_dafny.Set({False}))): _dafny.CodePoint('d')})

    @staticmethod
    def fm73(globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: _dafny.Seq
            for compr_18_ in (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('k')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i'), _dafny.CodePoint('t')]) for d_60_i0_ in range(default__.abs(259))])))).Elements:
                d_61_v0_: _dafny.Seq = compr_18_
                if (d_61_v0_) in (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('k')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i'), _dafny.CodePoint('t')]) for d_60_i0_ in range(default__.abs(259))])))):
                    coll18_[d_61_v0_] = (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_62_i1_ in range(default__.abs(-654))])), (0) - (len(_dafny.Map({D15_DC40(not(False), True): _dafny.Map({True: len(_dafny.Set({True, False}))})})))})).ispropersubset(_dafny.Set({375, 115}))
            return _dafny.Map(coll18_)
        return iife18_()
        

    @staticmethod
    def fm74(p0, p1, p2, p3, globalState):
        return D10_DC27((_dafny.SeqWithoutIsStrInference([True])) < (_dafny.SeqWithoutIsStrInference([True, False, False])), True, not(False), default__.safeModuloInt(406, (0) - ((_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))): False})), 708, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-116, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cg")))})), 514, (_dafny.MultiSet([True])).cardinality])): False}))])).cardinality)), (_dafny.MultiSet([593, 338, 444, 452])).issubset(_dafny.MultiSet([(0) - (-655)])))

    @staticmethod
    def fm75(p0, globalState):
        return D15_DC38((_dafny.Set({False, True})).issubset(_dafny.Set({not(True), not(False)})), len(_dafny.SeqWithoutIsStrInference([False, False, not(False), True, not(not(not(True)))])), (-404 if True else (0) - (len(_dafny.Set({True, True, False})))))

    @staticmethod
    def fm76(p0, p1, p2, globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(519, 867):
                d_63_v0_: int = compr_19_
                if ((519) <= (d_63_v0_)) and ((d_63_v0_) < (867)):
                    coll19_[(d_63_v0_) - (208)] = (_dafny.SeqWithoutIsStrInference([True, not(True)]) if not(False) else _dafny.SeqWithoutIsStrInference([False]))
            return _dafny.Map(coll19_)
        return iife19_()
        

    @staticmethod
    def fm77(globalState):
        return D9_DC23(_dafny.MultiSet([False, True, True]))

    @staticmethod
    def fm78(p0, p1, p2, p3, globalState):
        return D22_DC59(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "venkdua")), default__.safeModuloInt(945, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doga"))): len(_dafny.Map({not(False): 103}))}))))

    @staticmethod
    def fm79(p0, p1, p2, p3, globalState):
        return D20_DC53(D20_DC51((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([707, 859])))), (D6_DC15(False)).cf20, not(False)))

    @staticmethod
    def fm80(p0, globalState):
        return D2_DC4(_dafny.CodePoint('c'))

    @staticmethod
    def fm81(p0, p1, globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: D2
            for compr_20_ in (_dafny.Set({D2_DC6(D2_DC5(not(True)))})).Elements:
                d_64_v0_: D2 = compr_20_
                if (d_64_v0_) in (_dafny.Set({D2_DC6(D2_DC5(not(True)))})):
                    def iife21_():
                        coll21_ = _dafny.Set()
                        compr_21_: int
                        for compr_21_ in _dafny.IntegerRange(758, 53):
                            d_65_v1_: int = compr_21_
                            if ((758) <= (d_65_v1_)) and ((d_65_v1_) < (53)):
                                coll21_ = coll21_.union(_dafny.Set([(d_65_v1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fd"))))]))
                        return _dafny.Set(coll21_)
                    coll20_[d_64_v0_] = iife21_()

            return _dafny.Map(coll20_)
        if (D2_DC6(D2_DC6(D2_DC4(_dafny.CodePoint('c'))))) not in ((D35_DC93(iife20_()
)).cf149):
            return D25_DC69(889, len(_dafny.Map({_dafny.Map({-490: True}): 562})), -858)
        elif True:
            return D25_DC69(len(_dafny.Set({False})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i'), _dafny.CodePoint('q'), _dafny.CodePoint('g'), _dafny.CodePoint('k')])), (0) - ((0) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([434]))})))))

    @staticmethod
    def fm82(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([922, -160])) + (_dafny.SeqWithoutIsStrInference([98]))

    @staticmethod
    def fm83(p0, p1, globalState):
        source3_ = D20_DC51(len(_dafny.Map({False: False})), True, False)
        if source3_.is_DC51:
            d_66___mcc_h0_ = source3_.cf75
            d_67___mcc_h1_ = source3_.cf76
            d_68___mcc_h2_ = source3_.cf77
            d_69_cf77_ = d_68___mcc_h2_
            d_70_cf76_ = d_67___mcc_h1_
            d_71_cf75_ = d_66___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: d_69_cf77_})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_69_cf77_: d_69_cf77_})]))
        elif source3_.is_DC52:
            return (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}) for d_72_i0_ in range(default__.abs(521))]))
        elif source3_.is_DC50:
            d_73___mcc_h3_ = source3_.cf74
            d_74_cf74_ = d_73___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}) for d_75_i1_ in range(default__.abs(-741))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}), _dafny.Map({True: True}), _dafny.Map({True: False})]))
        elif True:
            d_76___mcc_h4_ = source3_.cf78
            d_77_cf78_ = d_76___mcc_h4_
            return _dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}) for d_78_i2_ in range(default__.abs(-506))])

    @staticmethod
    def fm84(globalState):
        def iife22_():
            def iife24_():
                coll24_ = _dafny.Set()
                compr_24_: _dafny.Map
                for compr_24_ in (_dafny.MultiSet([_dafny.Map({not(False): False})])).Elements:
                    d_81_v1_: _dafny.Map = compr_24_
                    if (d_81_v1_) in (_dafny.MultiSet([_dafny.Map({not(False): False})])):
                        coll24_ = coll24_.union(_dafny.Set([d_81_v1_]))
                return _dafny.Set(coll24_)
            coll22_ = _dafny.Map()
            def iife23_():
                coll23_ = _dafny.Set()
                compr_23_: _dafny.Map
                for compr_23_ in (_dafny.MultiSet([_dafny.Map({not(False): False})])).Elements:
                    d_79_v1_: _dafny.Map = compr_23_
                    if (d_79_v1_) in (_dafny.MultiSet([_dafny.Map({not(False): False})])):
                        coll23_ = coll23_.union(_dafny.Set([d_79_v1_]))
                return _dafny.Set(coll23_)
            compr_22_: _dafny.Map
            for compr_22_ in (iife23_()
            ).Elements:
                d_80_v0_: _dafny.Map = compr_22_
                if (d_80_v0_) in (iife24_()
                ):
                    coll22_[d_80_v0_] = _dafny.Map({len(_dafny.Map({True: 666})): (_dafny.MultiSet([False])).cardinality})
            return _dafny.Map(coll22_)
        return (D36_DC97(iife22_()
)).cf157

    @staticmethod
    def fm85(p0, p1, globalState):
        return D1_DC2((len(_dafny.Map({710: D10_DC29(D10_DC26(_dafny.Map({116: False})))}))) + (-167), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jssg"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcebl")), (-705) < (589))

    @staticmethod
    def fm86(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xy"))): (_dafny.Set({not(False)})) - (_dafny.Set({True, False}))})

    @staticmethod
    def fm87(p0, p1, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: D9
            for compr_25_ in (_dafny.SeqWithoutIsStrInference([D9_DC23(_dafny.MultiSet([False])), D9_DC23(_dafny.MultiSet([False, True, False]))])).Elements:
                d_82_v0_: D9 = compr_25_
                if (d_82_v0_) in (_dafny.SeqWithoutIsStrInference([D9_DC23(_dafny.MultiSet([False])), D9_DC23(_dafny.MultiSet([False, True, False]))])):
                    coll25_[d_82_v0_] = not(not(not(False)))
            return _dafny.Map(coll25_)
        return (iife25_()
        ) | (_dafny.Map({D9_DC23(_dafny.MultiSet([False])): True}))

    @staticmethod
    def m0(p0, p1, p2, globalState):
        d_83_v0_: _dafny.Seq
        d_83_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chpsjrs"))
        d_84_v1_: D17
        d_84_v1_ = D17_DC44(p2, d_83_v0_)
        d_85_v2_: _dafny.Array
        nw0_ = _dafny.Array(None, 10)
        nw0_[int(0)] = (d_83_v0_) + (d_83_v0_)
        nw0_[int(1)] = ((d_83_v0_) + (d_83_v0_)).set(default__.safeIndex(p1, len((d_83_v0_) + (d_83_v0_))), default__.fm42(p1, p1, globalState))
        nw0_[int(2)] = (d_83_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "inmioisg")))
        nw0_[int(3)] = d_83_v0_
        nw0_[int(4)] = d_83_v0_
        nw0_[int(5)] = d_83_v0_
        nw0_[int(6)] = d_83_v0_
        nw0_[int(7)] = d_83_v0_
        nw0_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_86_i0_ in range(default__.abs(-148))])
        nw0_[int(9)] = ((d_83_v0_).set(default__.safeIndex(p1, len(d_83_v0_)), _dafny.CodePoint('t'))) + ((d_84_v1_).cf63)
        d_85_v2_ = nw0_
        index0_ = default__.safeIndex(763, (d_85_v2_).length(0))
        (d_85_v2_)[index0_] = d_83_v0_
        index1_ = default__.safeIndex(763, (d_85_v2_).length(0))
        (d_85_v2_)[index1_] = d_83_v0_
        hi0_ = p1
        for d_87_i1_ in range(p1, hi0_):
            (globalState).f24 = (0) - (((0) - (d_87_i1_)) * (-236))
            d_88_v3_: str
            d_88_v3_ = _dafny.CodePoint('j')
            (globalState).f23 = d_88_v3_
            (globalState).f24 = 600
            d_89_v4_: _dafny.Set
            d_89_v4_ = _dafny.Set({p2, p2, p2})
            d_90_v5_: _dafny.Map
            d_90_v5_ = _dafny.Map({(d_87_i1_) + (d_87_i1_): len(d_89_v4_)})
            d_90_v5_ = (_dafny.Map({d_87_i1_: d_87_i1_})) | (d_90_v5_)
        if (p2) == ((p1) != (p1)):
            (globalState).f21 = not(p2)
            (globalState).f1 = p1
            d_91_v7_: _dafny.MultiSet
            d_91_v7_ = _dafny.MultiSet([True])
            d_92_v8_: D9
            d_92_v8_ = D9_DC23(d_91_v7_)
            d_93_v9_: _dafny.Map
            d_93_v9_ = _dafny.Map({d_92_v8_: p1})
            d_94_v10_: _dafny.Map
            d_94_v10_ = _dafny.Map({d_92_v8_: p2})
            d_95_v12_: _dafny.Seq
            d_95_v12_ = _dafny.SeqWithoutIsStrInference([d_92_v8_])
            d_96_v13_: _dafny.Seq
            def iife26_():
                coll26_ = _dafny.Map()
                compr_26_: D9
                for compr_26_ in (d_93_v9_).keys.Elements:
                    d_97_v6_: D9 = compr_26_
                    if (d_97_v6_) in (d_93_v9_):
                        coll26_[d_97_v6_] = p2
                return _dafny.Map(coll26_)
            def iife27_():
                coll27_ = _dafny.Map()
                compr_27_: D9
                for compr_27_ in (d_95_v12_).Elements:
                    d_98_v11_: D9 = compr_27_
                    if (d_98_v11_) in (d_95_v12_):
                        coll27_[d_98_v11_] = p2
                return _dafny.Map(coll27_)
            d_96_v13_ = _dafny.SeqWithoutIsStrInference([iife26_()
            , d_94_v10_, iife27_()
            ])
            d_99_v14_: str
            d_99_v14_ = _dafny.CodePoint('h')
            d_100_v15_: _dafny.Map
            d_100_v15_ = _dafny.Map({d_99_v14_: p2})
            d_101_v16_: _dafny.Array
            nw1_ = _dafny.Array(None, 28)
            nw1_[int(0)] = p2
            nw1_[int(1)] = True
            nw1_[int(2)] = p2
            nw1_[int(3)] = p2
            nw1_[int(4)] = True
            nw1_[int(5)] = p2
            nw1_[int(6)] = p2
            nw1_[int(7)] = not(p2)
            nw1_[int(8)] = p2
            nw1_[int(9)] = p2
            nw1_[int(10)] = p2
            nw1_[int(11)] = p2
            nw1_[int(12)] = p2
            nw1_[int(13)] = p2
            nw1_[int(14)] = p2
            nw1_[int(15)] = p2
            nw1_[int(16)] = p2
            nw1_[int(17)] = p2
            nw1_[int(18)] = p2
            nw1_[int(19)] = p2
            nw1_[int(20)] = p2
            nw1_[int(21)] = ((d_100_v15_)[_dafny.CodePoint('l')] if (_dafny.CodePoint('l')) in (d_100_v15_) else p2)
            nw1_[int(22)] = p2
            nw1_[int(23)] = p2
            nw1_[int(24)] = p2
            nw1_[int(25)] = p2
            nw1_[int(26)] = True
            nw1_[int(27)] = p2
            d_101_v16_ = nw1_
            d_102_v17_: _dafny.Array
            nw2_ = _dafny.Array(None, 14)
            nw2_[int(0)] = d_101_v16_
            nw2_[int(1)] = d_101_v16_
            nw2_[int(2)] = d_101_v16_
            nw2_[int(3)] = d_101_v16_
            nw2_[int(4)] = d_101_v16_
            nw2_[int(5)] = d_101_v16_
            nw2_[int(6)] = d_101_v16_
            nw2_[int(7)] = d_101_v16_
            nw2_[int(8)] = d_101_v16_
            nw2_[int(9)] = d_101_v16_
            nw2_[int(10)] = d_101_v16_
            nw2_[int(11)] = d_101_v16_
            nw2_[int(12)] = d_101_v16_
            nw2_[int(13)] = d_101_v16_
            d_102_v17_ = nw2_
            d_103_v18_: _dafny.Seq
            d_103_v18_ = _dafny.SeqWithoutIsStrInference([d_83_v0_, d_83_v0_])
            d_104_v19_: C6
            nw3_ = C6()
            nw3_.ctor__(d_102_v17_, d_103_v18_, p2, p2)
            d_104_v19_ = nw3_
            d_105_v20_: _dafny.Map
            d_105_v20_ = _dafny.Map({d_104_v19_: (d_96_v13_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_92_v8_: True}), default__.fm87(p2, p1, globalState)]))})
            d_106_v21_: _dafny.Set
            d_106_v21_ = _dafny.Set({(p2 if True else p2), not (p2) or (p2), p2})
            rhs0_ = ((d_105_v20_)[d_104_v19_] if (d_104_v19_) in (d_105_v20_) else d_96_v13_)
            rhs1_ = d_106_v21_
            rhs2_ = len((d_83_v0_) + (d_83_v0_))
            rhs3_ = (0) - (p1)
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = globalState
            d_96_v13_ = rhs0_
            lhs0_.f26 = rhs1_
            lhs1_.f1 = rhs2_
            lhs2_.f24 = rhs3_
            (globalState).f22 = p2
            d_107_v22_: C12
            nw4_ = C12()
            nw4_.ctor__()
            d_107_v22_ = nw4_
            d_107_v22_ = d_107_v22_
        elif True:
            d_108_v23_: _dafny.Array
            def lambda0_(d_109_p0_, d_110_p2_):
                def lambda1_(d_111_i2_):
                    return _dafny.Map({_dafny.CodePoint('b'): ((d_109_p0_)[False] if (False) in (d_109_p0_) else d_110_p2_)})

                return lambda1_

            init0_ = lambda0_(p0, p2)
            nw5_ = _dafny.Array(None, 4)
            for i0_0_ in range(nw5_.length(0)):
                nw5_[i0_0_] = init0_(i0_0_)
            d_108_v23_ = nw5_
            d_112_v24_: _dafny.Seq
            d_112_v24_ = _dafny.SeqWithoutIsStrInference([d_108_v23_, d_108_v23_])
            d_113_v25_: _dafny.Array
            nw6_ = _dafny.Array(None, 9)
            nw6_[int(0)] = (d_108_v23_ if p2 else d_108_v23_)
            nw6_[int(1)] = (d_108_v23_ if False else d_108_v23_)
            nw6_[int(2)] = ((d_112_v24_)[default__.safeIndex(p1, len(d_112_v24_))] if p2 else d_108_v23_)
            nw6_[int(3)] = d_108_v23_
            nw6_[int(4)] = d_108_v23_
            nw6_[int(5)] = d_108_v23_
            nw6_[int(6)] = d_108_v23_
            nw6_[int(7)] = d_108_v23_
            nw6_[int(8)] = d_108_v23_
            d_113_v25_ = nw6_
            index2_ = default__.safeIndex(333, (d_113_v25_).length(0))
            (d_113_v25_)[index2_] = d_108_v23_
            d_114_v26_: str
            d_114_v26_ = _dafny.CodePoint('w')
            d_115_v27_: T1
            nw7_ = C3()
            nw7_.ctor__(p2, p2, p2, default__.fm27(p1, d_114_v26_, p1, globalState))
            d_115_v27_ = nw7_
            d_116_v28_: _dafny.Map
            d_116_v28_ = _dafny.Map({d_115_v27_: 987})
            d_117_v29_: _dafny.Array
            nw8_ = _dafny.Array(None, 4)
            nw8_[int(0)] = ((d_116_v28_)[d_115_v27_] if (d_115_v27_) in (d_116_v28_) else p1)
            nw8_[int(1)] = p1
            nw8_[int(2)] = p1
            nw8_[int(3)] = p1
            d_117_v29_ = nw8_
            index3_ = default__.safeIndex(30, (d_117_v29_).length(0))
            (d_117_v29_)[index3_] = p1
            index4_ = default__.safeIndex(333, (d_113_v25_).length(0))
            index5_ = default__.safeIndex(30, (d_117_v29_).length(0))
            rhs4_ = d_108_v23_
            rhs5_ = p1
            lhs3_ = d_113_v25_
            lhs4_ = default__.safeIndex(333, (d_113_v25_).length(0))
            lhs5_ = d_117_v29_
            lhs6_ = default__.safeIndex(30, (d_117_v29_).length(0))
            lhs3_[lhs4_] = rhs4_
            lhs5_[lhs6_] = rhs5_
            d_118_v30_: _dafny.MultiSet
            d_118_v30_ = _dafny.MultiSet([p2, p2, p2, p2])
            d_119_v31_: T0
            nw9_ = C5()
            nw9_.ctor__(p2, d_114_v26_, ((d_118_v30_).cardinality) >= (p1), (p2) or (p2))
            d_119_v31_ = nw9_
            d_120_v32_: D15
            d_120_v32_ = D15_DC39((d_85_v2_)[default__.safeIndex(763, (d_85_v2_).length(0))])
            d_121_v33_: _dafny.MultiSet
            d_121_v33_ = _dafny.MultiSet([d_120_v32_])
            d_122_v34_: _dafny.Map
            d_122_v34_ = _dafny.Map({(d_121_v33_).issubset(d_121_v33_): ((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]) == ((0) - ((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]))})
            d_123_v35_: C13
            nw10_ = C13()
            nw10_.ctor__()
            d_123_v35_ = nw10_
            d_124_v36_: _dafny.Map
            d_124_v36_ = _dafny.Map({d_123_v35_: d_115_v27_})
            rhs6_ = default__.fm12(default__.fm8(globalState), (d_123_v35_) in (d_124_v36_), globalState)
            rhs7_ = d_119_v31_
            rhs8_ = d_115_v27_
            rhs9_ = (d_122_v34_) | (p0)
            lhs7_ = globalState
            lhs7_.f24 = rhs6_
            d_119_v31_ = rhs7_
            d_115_v27_ = rhs8_
            d_122_v34_ = rhs9_
            d_125_v37_: C0
            nw11_ = C0()
            nw11_.ctor__()
            d_125_v37_ = nw11_
            if d_119_v31_.f28:
                d_126_v38_: _dafny.MultiSet
                d_126_v38_ = _dafny.MultiSet([d_115_v27_])
                d_127_v39_: _dafny.Seq
                d_127_v39_ = _dafny.SeqWithoutIsStrInference([d_119_v31_.f28, d_119_v31_.f28])
                d_128_v40_: _dafny.Array
                def lambda2_(d_129_v31_):
                    def lambda3_(d_130_i3_):
                        return (d_129_v31_).f29

                    return lambda3_

                init1_ = lambda2_(d_119_v31_)
                nw12_ = _dafny.Array(None, 26)
                for i0_1_ in range(nw12_.length(0)):
                    nw12_[i0_1_] = init1_(i0_1_)
                d_128_v40_ = nw12_
                d_131_v41_: _dafny.Map
                d_131_v41_ = _dafny.Map({d_128_v40_: d_128_v40_})
                d_132_v42_: D5
                d_132_v42_ = D5_DC12(((d_131_v41_)[d_128_v40_] if (d_128_v40_) in (d_131_v41_) else d_128_v40_))
                d_133_v43_: _dafny.Array
                nw13_ = _dafny.Array(None, 9)
                nw13_[int(0)] = d_128_v40_
                nw13_[int(1)] = d_128_v40_
                nw13_[int(2)] = d_128_v40_
                nw13_[int(3)] = d_128_v40_
                nw13_[int(4)] = (d_132_v42_).cf17
                nw13_[int(5)] = d_128_v40_
                nw13_[int(6)] = d_128_v40_
                nw13_[int(7)] = d_128_v40_
                nw13_[int(8)] = d_128_v40_
                d_133_v43_ = nw13_
                d_134_v44_: _dafny.Map
                d_134_v44_ = _dafny.Map({(_dafny.MultiSet([d_126_v38_])).set(d_126_v38_, default__.abs(default__.fm3((d_127_v39_)[default__.safeIndex((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))], len(d_127_v39_))], (d_119_v31_).f29, (d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))], globalState))): d_133_v43_})
                d_134_v44_ = (d_134_v44_).set(_dafny.MultiSet([d_126_v38_, d_126_v38_]), d_133_v43_)
                d_135_v45_: _dafny.Array
                nw14_ = _dafny.Array(None, 16)
                d_135_v45_ = nw14_
                d_136_v46_: D29
                d_136_v46_ = D29_DC78(d_135_v45_)
                d_136_v46_ = d_136_v46_
                d_137_v47_: _dafny.Map
                d_137_v47_ = _dafny.Map({-225: d_83_v0_})
                d_138_v48_: C6
                nw15_ = C6()
                nw15_.ctor__(d_133_v43_, _dafny.SeqWithoutIsStrInference([d_83_v0_ for d_139_i4_ in range(default__.abs(31))]), (d_119_v31_).f29, (d_119_v31_).f29)
                d_138_v48_ = nw15_
                d_140_v49_: _dafny.Seq
                d_140_v49_ = _dafny.SeqWithoutIsStrInference([d_138_v48_, d_138_v48_, d_138_v48_, d_138_v48_, d_138_v48_])
                (globalState).f27 = not(not((d_114_v26_) not in (((d_137_v47_)[len(d_140_v49_)] if (len(d_140_v49_)) in (d_137_v47_) else _dafny.SeqWithoutIsStrInference([d_114_v26_ for d_141_i5_ in range(default__.abs(400))])))))
                d_142_v50_: _dafny.MultiSet
                d_142_v50_ = _dafny.MultiSet([(d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]])
                d_143_v51_: _dafny.Set
                d_143_v51_ = _dafny.Set({(d_119_v31_).f29, d_119_v31_.f28})
                d_144_v52_: _dafny.Seq
                d_144_v52_ = _dafny.SeqWithoutIsStrInference([(0) - (p1), len(d_143_v51_)])
                d_145_v53_: C9
                nw16_ = C9()
                nw16_.ctor__((d_142_v50_) != (_dafny.MultiSet(d_144_v52_)), not(False))
                d_145_v53_ = nw16_
                (globalState).f24 = default__.fm27((0) - (default__.safeModuloInt((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))], (d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))])), d_114_v26_, (d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))], globalState)
            elif True:
                (globalState).f1 = p1
                d_146_v54_: _dafny.Map
                d_146_v54_ = _dafny.Map({(d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]: d_117_v29_})
                d_146_v54_ = (d_146_v54_).set((len(_dafny.Map({d_125_v37_: d_114_v26_}))) * ((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]), d_117_v29_)
                d_147_v55_: _dafny.Seq
                d_147_v55_ = _dafny.SeqWithoutIsStrInference([d_115_v27_])
                d_148_v56_: _dafny.Array
                nw17_ = _dafny.Array(None, 18)
                nw17_[int(0)] = d_115_v27_
                nw17_[int(1)] = d_115_v27_
                nw17_[int(2)] = d_115_v27_
                nw17_[int(3)] = d_115_v27_
                nw17_[int(4)] = d_115_v27_
                nw17_[int(5)] = d_115_v27_
                nw17_[int(6)] = d_115_v27_
                nw17_[int(7)] = (d_115_v27_ if d_119_v31_.f28 else d_115_v27_)
                nw17_[int(8)] = d_115_v27_
                nw17_[int(9)] = d_115_v27_
                nw17_[int(10)] = d_115_v27_
                nw17_[int(11)] = d_115_v27_
                nw17_[int(12)] = (d_147_v55_)[default__.safeIndex(len(d_83_v0_), len(d_147_v55_))]
                nw17_[int(13)] = d_115_v27_
                nw17_[int(14)] = (d_147_v55_)[default__.safeIndex(p1, len(d_147_v55_))]
                nw17_[int(15)] = d_115_v27_
                nw17_[int(16)] = d_115_v27_
                nw17_[int(17)] = d_115_v27_
                d_148_v56_ = nw17_
                d_149_v57_: D25
                d_149_v57_ = D25_DC68(d_115_v27_)
                index6_ = default__.safeIndex(778, (d_148_v56_).length(0))
                (d_148_v56_)[index6_] = (d_149_v57_).cf105
                index7_ = default__.safeIndex(778, (d_148_v56_).length(0))
                (d_148_v56_)[index7_] = d_115_v27_
                (globalState).f27 = default__.fm38(globalState)
                d_150_v58_: _dafny.Array
                nw18_ = _dafny.Array(False, 1)
                d_150_v58_ = nw18_
                d_151_v59_: D5
                d_151_v59_ = D5_DC12(d_150_v58_)
                d_152_v60_: _dafny.Map
                d_152_v60_ = _dafny.Map({d_119_v31_.f28: d_150_v58_})
                d_153_v61_: _dafny.Array
                nw19_ = _dafny.Array(None, 18)
                nw19_[int(0)] = d_150_v58_
                nw19_[int(1)] = d_150_v58_
                nw19_[int(2)] = (d_150_v58_ if p2 else (d_151_v59_).cf17)
                nw19_[int(3)] = d_150_v58_
                nw19_[int(4)] = d_150_v58_
                nw19_[int(5)] = d_150_v58_
                nw19_[int(6)] = d_150_v58_
                nw19_[int(7)] = ((d_152_v60_)[(d_119_v31_).f29] if ((d_119_v31_).f29) in (d_152_v60_) else d_150_v58_)
                nw19_[int(8)] = d_150_v58_
                nw19_[int(9)] = (d_150_v58_ if d_119_v31_.f28 else d_150_v58_)
                nw19_[int(10)] = d_150_v58_
                nw19_[int(11)] = d_150_v58_
                nw19_[int(12)] = d_150_v58_
                nw19_[int(13)] = (d_150_v58_ if p2 else d_150_v58_)
                nw19_[int(14)] = d_150_v58_
                nw19_[int(15)] = d_150_v58_
                nw19_[int(16)] = d_150_v58_
                nw19_[int(17)] = d_150_v58_
                d_153_v61_ = nw19_
                index8_ = default__.safeIndex(588, (d_153_v61_).length(0))
                (d_153_v61_)[index8_] = d_150_v58_
                index9_ = default__.safeIndex(588, (d_153_v61_).length(0))
                (d_153_v61_)[index9_] = d_150_v58_
            d_154_v62_: _dafny.Set
            d_154_v62_ = _dafny.Set({5})
            if not(not((d_154_v62_).issubset(d_154_v62_))):
                d_155_v63_: _dafny.Set
                d_155_v63_ = _dafny.Set({d_83_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_156_i6_ in range(default__.abs(982))]), (d_83_v0_) + ((d_85_v2_)[default__.safeIndex(763, (d_85_v2_).length(0))])})
                d_157_v64_: _dafny.Map
                d_157_v64_ = _dafny.Map({False: (p1) - ((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))])})
                rhs10_ = d_155_v63_
                rhs11_ = not(d_119_v31_.f28)
                rhs12_ = d_157_v64_
                rhs13_ = (d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]
                lhs8_ = globalState
                lhs9_ = globalState
                d_155_v63_ = rhs10_
                lhs8_.f21 = rhs11_
                d_157_v64_ = rhs12_
                lhs9_.f1 = rhs13_
                (d_119_v31_).f28 = (_dafny.SeqWithoutIsStrInference([len(d_83_v0_)])) == (_dafny.SeqWithoutIsStrInference([(d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))], (d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_154_v62_]))]))
                d_158_v65_: _dafny.Array
                nw20_ = _dafny.Array(False, 8)
                d_158_v65_ = nw20_
                def lambda4_(d_159_v31_):
                    def lambda5_(d_160_i7_):
                        return (d_159_v31_).f29

                    return lambda5_

                init2_ = lambda4_(d_119_v31_)
                nw21_ = _dafny.Array(None, 7)
                for i0_2_ in range(nw21_.length(0)):
                    nw21_[i0_2_] = init2_(i0_2_)
                d_158_v65_ = nw21_
                d_161_v66_: _dafny.Array
                nw22_ = _dafny.Array(D7.default()(), 6)
                d_161_v66_ = nw22_
                d_161_v66_ = d_161_v66_
                d_162_v67_: D26
                d_162_v67_ = D26_DC72(d_117_v29_, (d_119_v31_).f29)
                d_163_v68_: _dafny.Seq
                d_163_v68_ = _dafny.SeqWithoutIsStrInference([d_162_v67_, d_162_v67_, d_162_v67_])
                rhs14_ = ((p0)[d_119_v31_.f28] if (d_119_v31_.f28) in (p0) else True)
                rhs15_ = d_114_v26_
                rhs16_ = (p1) + (len(d_163_v68_))
                lhs10_ = globalState
                lhs11_ = globalState
                lhs12_ = globalState
                lhs10_.f27 = rhs14_
                lhs11_.f23 = rhs15_
                lhs12_.f24 = rhs16_
            elif True:
                (globalState).f22 = (d_119_v31_).f29
                d_164_v69_: C3
                nw23_ = C3()
                nw23_.ctor__((d_119_v31_).f29, not(p2), d_119_v31_.f28, (d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))])
                d_164_v69_ = nw23_
                d_165_v70_: _dafny.Map
                d_165_v70_ = _dafny.Map({p1: d_164_v69_})
                d_166_v71_: _dafny.Array
                nw24_ = _dafny.Array(None, 17)
                nw24_[int(0)] = d_164_v69_
                nw24_[int(1)] = d_164_v69_
                nw24_[int(2)] = ((d_165_v70_)[(d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]] if ((d_117_v29_)[default__.safeIndex(30, (d_117_v29_).length(0))]) in (d_165_v70_) else d_164_v69_)
                nw24_[int(3)] = d_164_v69_
                nw24_[int(4)] = d_164_v69_
                nw24_[int(5)] = d_164_v69_
                nw24_[int(6)] = d_164_v69_
                nw24_[int(7)] = d_164_v69_
                nw24_[int(8)] = d_164_v69_
                nw24_[int(9)] = d_164_v69_
                nw24_[int(10)] = d_164_v69_
                nw24_[int(11)] = d_164_v69_
                nw24_[int(12)] = d_164_v69_
                nw24_[int(13)] = d_164_v69_
                nw24_[int(14)] = d_164_v69_
                nw24_[int(15)] = d_164_v69_
                nw24_[int(16)] = d_164_v69_
                d_166_v71_ = nw24_
                index10_ = default__.safeIndex(534, (d_166_v71_).length(0))
                (d_166_v71_)[index10_] = d_164_v69_
                index11_ = default__.safeIndex(534, (d_166_v71_).length(0))
                (d_166_v71_)[index11_] = d_164_v69_
                d_167_v72_: C7
                nw25_ = C7()
                nw25_.ctor__((d_119_v31_.f28 if d_119_v31_.f28 else (d_164_v69_).f40), (d_119_v31_).f29)
                d_167_v72_ = nw25_
                d_167_v72_ = d_167_v72_
                d_168_v73_: _dafny.Seq
                d_168_v73_ = _dafny.SeqWithoutIsStrInference([(d_119_v31_).f29, (d_119_v31_).f29])
                d_169_v74_: _dafny.Map
                d_169_v74_ = _dafny.Map({p1: (d_168_v73_)[default__.safeIndex(957, len(d_168_v73_))]})
                d_170_v75_: D10
                d_170_v75_ = D10_DC27(d_119_v31_.f28, (d_164_v69_).f40, d_119_v31_.f28, 28, True)
                pat_let_tv0_ = d_117_v29_
                pat_let_tv1_ = d_117_v29_
                d_171_v76_: _dafny.Map
                def iife28_(_pat_let0_0):
                    def iife29_(d_172_dt__update__tmp_h0_):
                        def iife30_(_pat_let1_0):
                            def iife31_(d_173_dt__update_hcf43_h0_):
                                return D10_DC27((d_172_dt__update__tmp_h0_).cf40, (d_172_dt__update__tmp_h0_).cf41, (d_172_dt__update__tmp_h0_).cf42, d_173_dt__update_hcf43_h0_, (d_172_dt__update__tmp_h0_).cf44)
                            return iife31_(_pat_let1_0)
                        return iife30_((pat_let_tv1_)[default__.safeIndex(30, (pat_let_tv0_).length(0))])
                    return iife29_(_pat_let0_0)
                d_171_v76_ = _dafny.Map({d_169_v74_: iife28_(d_170_v75_)})
                d_171_v76_ = (d_171_v76_).set(d_169_v74_, d_170_v75_)
                index12_ = default__.safeIndex(30, (d_117_v29_).length(0))
                (d_117_v29_)[index12_] = default__.fm27(811, d_114_v26_, p1, globalState)
        if p2:
            (globalState).f1 = p1
            (globalState).f22 = p2
            (globalState).f22 = p2
            d_174_v77_: _dafny.Map
            d_174_v77_ = _dafny.Map({p2: p1})
            d_175_v78_: _dafny.Map
            d_175_v78_ = _dafny.Map({p2: d_174_v77_})
            (globalState).f1 = (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p1: p2}) for d_176_i8_ in range(default__.abs(557))])) if (p2 if p2 else p2) else default__.fm3(p2, not(p2), len(((d_175_v78_)[p2] if (p2) in (d_175_v78_) else _dafny.Map({p2: p1}))), globalState))
            d_177_v79_: _dafny.Array
            nw26_ = _dafny.Array(int(0), 5)
            d_177_v79_ = nw26_
            index13_ = default__.safeIndex(519, (d_177_v79_).length(0))
            (d_177_v79_)[index13_] = p1
            d_178_v80_: _dafny.Map
            d_178_v80_ = _dafny.Map({p1: p1})
            d_179_v81_: _dafny.MultiSet
            d_179_v81_ = _dafny.MultiSet([len(d_178_v80_)])
            index14_ = default__.safeIndex(519, (d_177_v79_).length(0))
            (d_177_v79_)[index14_] = default__.safeDivisionInt(p1, ((d_179_v81_)[p1] if (p1) in (d_179_v81_) else p1))
        elif True:
            d_180_v82_: str
            d_180_v82_ = _dafny.CodePoint('n')
            d_181_v83_: _dafny.Map
            d_181_v83_ = _dafny.Map({d_180_v82_: p1})
            (globalState).f1 = ((d_181_v83_)[d_180_v82_] if (d_180_v82_) in (d_181_v83_) else (873) * (p1))
            (globalState).f1 = p1
            d_182_v84_: _dafny.Array
            nw27_ = _dafny.Array(False, 18)
            d_182_v84_ = nw27_
            index15_ = default__.safeIndex(153, (d_182_v84_).length(0))
            (d_182_v84_)[index15_] = p2
            index16_ = default__.safeIndex(153, (d_182_v84_).length(0))
            (d_182_v84_)[index16_] = not ((p1) == (-230)) or (p2)
            d_183_v85_: _dafny.Array
            nw28_ = _dafny.Array(int(0), 27)
            d_183_v85_ = nw28_
            d_183_v85_ = d_183_v85_
            index17_ = default__.safeIndex(951, (d_183_v85_).length(0))
            (d_183_v85_)[index17_] = (p1) + (p1)
            index18_ = default__.safeIndex(951, (d_183_v85_).length(0))
            (d_183_v85_)[index18_] = p1
        index19_ = default__.safeIndex(763, (d_85_v2_).length(0))
        (d_85_v2_)[index19_] = (d_85_v2_)[default__.safeIndex(763, (d_85_v2_).length(0))]
        d_184_v86_: D27
        d_184_v86_ = D27_DC75((p2) and (p2))
        d_184_v86_ = D27_DC75(p2)

    @staticmethod
    def Main(noArgsParameter__):
        d_185_v0_: _dafny.Array
        nw29_ = _dafny.Array(int(0), 12)
        d_185_v0_ = nw29_
        d_186_v1_: bool
        d_186_v1_ = True
        d_187_v2_: _dafny.Array
        nw30_ = _dafny.Array(None, 28)
        nw30_[int(0)] = False
        nw30_[int(1)] = d_186_v1_
        nw30_[int(2)] = d_186_v1_
        nw30_[int(3)] = d_186_v1_
        nw30_[int(4)] = d_186_v1_
        nw30_[int(5)] = d_186_v1_
        nw30_[int(6)] = d_186_v1_
        nw30_[int(7)] = d_186_v1_
        nw30_[int(8)] = d_186_v1_
        nw30_[int(9)] = d_186_v1_
        nw30_[int(10)] = d_186_v1_
        nw30_[int(11)] = d_186_v1_
        nw30_[int(12)] = d_186_v1_
        nw30_[int(13)] = d_186_v1_
        nw30_[int(14)] = d_186_v1_
        nw30_[int(15)] = d_186_v1_
        nw30_[int(16)] = d_186_v1_
        nw30_[int(17)] = True
        nw30_[int(18)] = d_186_v1_
        nw30_[int(19)] = d_186_v1_
        nw30_[int(20)] = d_186_v1_
        nw30_[int(21)] = d_186_v1_
        nw30_[int(22)] = d_186_v1_
        nw30_[int(23)] = d_186_v1_
        nw30_[int(24)] = d_186_v1_
        nw30_[int(25)] = d_186_v1_
        nw30_[int(26)] = d_186_v1_
        nw30_[int(27)] = d_186_v1_
        d_187_v2_ = nw30_
        d_188_v3_: _dafny.Map
        d_188_v3_ = _dafny.Map({d_187_v2_: d_186_v1_})
        d_189_v4_: _dafny.Array
        nw31_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_189_v4_ = nw31_
        d_190_v5_: _dafny.Seq
        d_190_v5_ = _dafny.SeqWithoutIsStrInference([d_186_v1_, d_186_v1_, d_186_v1_, d_186_v1_])
        d_191_v6_: _dafny.Set
        d_191_v6_ = _dafny.Set({d_186_v1_, d_186_v1_, d_186_v1_, d_186_v1_})
        d_192_v7_: int
        d_192_v7_ = 432
        d_193_v8_: _dafny.Seq
        d_193_v8_ = _dafny.SeqWithoutIsStrInference([d_192_v7_, -988, d_192_v7_])
        d_194_v9_: _dafny.Map
        d_194_v9_ = _dafny.Map({d_193_v8_: d_192_v7_})
        d_195_globalState_: GlobalState
        nw32_ = GlobalState()
        nw32_.ctor__(d_185_v0_, 593, _dafny.CodePoint('l'), _dafny.CodePoint('v'), 430, True, d_188_v3_, _dafny.CodePoint('o'), False, d_189_v4_, 899, d_190_v5_, -433, True, d_191_v6_, -753, 44, d_193_v8_, 575, d_194_v9_, -726, True, True, _dafny.CodePoint('j'), -158, False, d_191_v6_, False)
        d_195_globalState_ = nw32_
        hi1_ = -166
        for d_196_i0_ in range(d_192_v7_, hi1_):
            d_197_v10_: _dafny.Map
            d_197_v10_ = _dafny.Map({True: d_186_v1_})
            d_198_v11_: _dafny.Map
            d_198_v11_ = _dafny.Map({d_187_v2_: d_197_v10_})
            default__.m0(((d_198_v11_)[d_187_v2_] if (d_187_v2_) in (d_198_v11_) else d_197_v10_), d_192_v7_, d_186_v1_, d_195_globalState_)
            index20_ = default__.safeIndex(317, (d_185_v0_).length(0))
            (d_185_v0_)[index20_] = d_192_v7_
            index21_ = default__.safeIndex(317, (d_185_v0_).length(0))
            (d_185_v0_)[index21_] = (d_192_v7_ if d_186_v1_ else default__.safeModuloInt(-330, (0) - (d_196_i0_)))
            (d_195_globalState_).f1 = (d_185_v0_)[default__.safeIndex(317, (d_185_v0_).length(0))]
            d_199_v12_: _dafny.Map
            d_199_v12_ = _dafny.Map({False: d_192_v7_})
            d_199_v12_ = (d_199_v12_).set(d_186_v1_, (d_185_v0_)[default__.safeIndex(317, (d_185_v0_).length(0))])
        if ((d_192_v7_) != (d_192_v7_)) == (d_186_v1_):
            rhs17_ = ((d_191_v6_ if d_186_v1_ else d_191_v6_)).intersection((d_191_v6_).intersection(d_191_v6_))
            rhs18_ = 999
            lhs13_ = d_195_globalState_
            lhs14_ = d_195_globalState_
            lhs13_.f26 = rhs17_
            lhs14_.f24 = rhs18_
            d_200_v13_: C6
            nw33_ = C6()
            nw33_.ctor__(d_189_v4_, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_201_i2_ in range(default__.abs(617))]) for d_202_i1_ in range(default__.abs(169))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_203_i4_ in range(default__.abs(760))]) for d_204_i3_ in range(default__.abs(582))])), d_186_v1_, d_186_v1_)
            d_200_v13_ = nw33_
            d_205_v14_: str
            d_205_v14_ = _dafny.CodePoint('d')
            (d_195_globalState_).f23 = d_205_v14_
            d_206_v15_: _dafny.Array
            nw34_ = _dafny.Array(D12.default()(), 9)
            d_206_v15_ = nw34_
            d_207_v16_: D12
            d_207_v16_ = D12_DC32()
            index22_ = default__.safeIndex(708, (d_206_v15_).length(0))
            (d_206_v15_)[index22_] = d_207_v16_
            d_208_v17_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_208_v17_ = nw35_
            d_209_v18_: T0
            nw36_ = C3()
            nw36_.ctor__(d_186_v1_, d_186_v1_, d_186_v1_, d_192_v7_)
            d_209_v18_ = nw36_
            d_210_v19_: D8
            d_210_v19_ = D8_DC20(d_209_v18_)
            d_211_v20_: _dafny.Seq
            d_211_v20_ = _dafny.SeqWithoutIsStrInference([d_209_v18_, d_209_v18_])
            d_212_v21_: _dafny.Array
            nw37_ = _dafny.Array(None, 15)
            nw37_[int(0)] = d_209_v18_
            nw37_[int(1)] = (D8_DC20(d_209_v18_)).cf26
            nw37_[int(2)] = d_209_v18_
            nw37_[int(3)] = d_209_v18_
            nw37_[int(4)] = d_209_v18_
            nw37_[int(5)] = d_209_v18_
            nw37_[int(6)] = d_209_v18_
            nw37_[int(7)] = (d_210_v19_).cf26
            nw37_[int(8)] = d_209_v18_
            nw37_[int(9)] = (d_211_v20_)[default__.safeIndex(461, len(d_211_v20_))]
            nw37_[int(10)] = d_209_v18_
            nw37_[int(11)] = d_209_v18_
            nw37_[int(12)] = d_209_v18_
            nw37_[int(13)] = d_209_v18_
            nw37_[int(14)] = d_209_v18_
            d_212_v21_ = nw37_
            index23_ = default__.safeIndex(605, (d_208_v17_).length(0))
            (d_208_v17_)[index23_] = d_212_v21_
            d_213_v22_: _dafny.MultiSet
            d_213_v22_ = _dafny.MultiSet([default__.fm25((0) - (d_192_v7_), d_192_v7_, d_195_globalState_)])
            index24_ = default__.safeIndex(708, (d_206_v15_).length(0))
            index25_ = default__.safeIndex(605, (d_208_v17_).length(0))
            rhs19_ = d_207_v16_
            rhs20_ = not((d_209_v18_).f29)
            rhs21_ = d_212_v21_
            rhs22_ = d_213_v22_
            lhs15_ = d_206_v15_
            lhs16_ = default__.safeIndex(708, (d_206_v15_).length(0))
            lhs17_ = d_195_globalState_
            lhs18_ = d_208_v17_
            lhs19_ = default__.safeIndex(605, (d_208_v17_).length(0))
            lhs15_[lhs16_] = rhs19_
            lhs17_.f13 = rhs20_
            lhs18_[lhs19_] = rhs21_
            d_213_v22_ = rhs22_
            (d_195_globalState_).f13 = ((0) - (d_192_v7_)) >= (default__.fm12((0) - (d_192_v7_), not(d_209_v18_.f28), d_195_globalState_))
        elif True:
            d_186_v1_ = d_186_v1_
            d_214_v23_: _dafny.Map
            d_214_v23_ = _dafny.Map({d_186_v1_: False})
            d_215_v24_: _dafny.Map
            d_215_v24_ = _dafny.Map({d_192_v7_: d_186_v1_})
            d_216_v25_: _dafny.Set
            d_216_v25_ = _dafny.Set({len(d_215_v24_), d_192_v7_})
            d_217_v26_: _dafny.Seq
            d_217_v26_ = _dafny.SeqWithoutIsStrInference([d_216_v25_])
            d_218_v27_: _dafny.Seq
            d_218_v27_ = _dafny.SeqWithoutIsStrInference([(d_217_v26_)[default__.safeIndex(808, len(d_217_v26_))], d_216_v25_])
            default__.m0(d_214_v23_, (len(d_218_v27_) if not(d_186_v1_) else d_192_v7_), ((d_215_v24_)[d_192_v7_] if (d_192_v7_) in (d_215_v24_) else d_186_v1_), d_195_globalState_)
            (d_195_globalState_).f13 = (d_217_v26_) < (_dafny.SeqWithoutIsStrInference([d_216_v25_, _dafny.Set({d_192_v7_})]))
            (d_195_globalState_).f21 = d_186_v1_
            d_219_v28_: _dafny.Array
            def lambda6_(d_220_i5_):
                return _dafny.CodePoint('h')

            init3_ = lambda6_
            nw38_ = _dafny.Array(None, 9)
            for i0_3_ in range(nw38_.length(0)):
                nw38_[i0_3_] = init3_(i0_3_)
            d_219_v28_ = nw38_
            d_221_v29_: _dafny.Array
            d_221_v29_ = d_219_v28_
            d_222_v30_: _dafny.Array
            nw39_ = _dafny.Array(None, 14)
            nw39_[int(0)] = d_221_v29_
            nw39_[int(1)] = d_221_v29_
            nw39_[int(2)] = d_221_v29_
            nw39_[int(3)] = d_221_v29_
            nw39_[int(4)] = d_221_v29_
            nw39_[int(5)] = d_221_v29_
            nw39_[int(6)] = d_221_v29_
            nw39_[int(7)] = d_221_v29_
            nw39_[int(8)] = d_221_v29_
            nw39_[int(9)] = d_221_v29_
            nw39_[int(10)] = (d_221_v29_)
            nw39_[int(11)] = d_221_v29_
            nw39_[int(12)] = d_221_v29_
            nw39_[int(13)] = (d_221_v29_ if d_186_v1_ else d_221_v29_)
            d_222_v30_ = nw39_
            d_222_v30_ = d_222_v30_
        if d_186_v1_:
            d_223_v31_: _dafny.Map
            d_223_v31_ = _dafny.Map({d_186_v1_: d_193_v8_})
            d_224_v32_: _dafny.Set
            d_224_v32_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dek"))) for d_225_i6_ in range(default__.abs(-367))]), ((d_223_v31_)[d_186_v1_] if (d_186_v1_) in (d_223_v31_) else d_193_v8_)})
            (d_195_globalState_).f1 = default__.safeDivisionInt(len(d_224_v32_), 125)
            d_226_v33_: _dafny.Map
            d_226_v33_ = _dafny.Map({d_186_v1_: d_186_v1_})
            default__.m0(d_226_v33_, (d_192_v7_) + (d_192_v7_), d_186_v1_, d_195_globalState_)
            d_227_v34_: _dafny.Seq
            d_227_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "liafj"))
            d_228_v35_: _dafny.Set
            d_228_v35_ = _dafny.Set({len(d_227_v34_)})
            d_229_v36_: D12
            d_229_v36_ = D12_DC31((_dafny.Set({d_192_v7_})) - (d_228_v35_))
            source4_ = d_229_v36_
            if source4_.is_DC32:
                (d_195_globalState_).f27 = d_186_v1_
                d_230_v37_: _dafny.MultiSet
                d_230_v37_ = _dafny.MultiSet([d_192_v7_, len(d_191_v6_), d_192_v7_])
                default__.m0(d_226_v33_, ((d_230_v37_)[(_dafny.MultiSet([not(d_186_v1_)])).cardinality] if ((_dafny.MultiSet([not(d_186_v1_)])).cardinality) in (d_230_v37_) else d_192_v7_), d_186_v1_, d_195_globalState_)
                d_231_v38_: C7
                nw40_ = C7()
                nw40_.ctor__((d_190_v5_)[default__.safeIndex(len(d_194_v9_), len(d_190_v5_))], not(default__.fm10(default__.fm38(d_195_globalState_), d_192_v7_, d_186_v1_, d_195_globalState_)))
                d_231_v38_ = nw40_
                nw41_ = C7()
                nw41_.ctor__(d_186_v1_, d_186_v1_)
                d_231_v38_ = nw41_
                d_232_v39_: C0
                nw42_ = C0()
                nw42_.ctor__()
                d_232_v39_ = nw42_
                d_233_v40_: str
                d_233_v40_ = _dafny.CodePoint('n')
                d_234_v41_: _dafny.Map
                d_234_v41_ = _dafny.Map({(d_230_v37_).cardinality: d_233_v40_})
                d_235_v42_: _dafny.Array
                nw43_ = _dafny.Array(None, 14)
                nw43_[int(0)] = _dafny.CodePoint('c')
                nw43_[int(1)] = d_233_v40_
                nw43_[int(2)] = d_233_v40_
                nw43_[int(3)] = _dafny.CodePoint('b')
                nw43_[int(4)] = d_233_v40_
                nw43_[int(5)] = (d_227_v34_)[default__.safeIndex(d_192_v7_, len(d_227_v34_))]
                nw43_[int(6)] = d_233_v40_
                nw43_[int(7)] = ((d_234_v41_)[d_192_v7_] if (d_192_v7_) in (d_234_v41_) else d_233_v40_)
                nw43_[int(8)] = _dafny.CodePoint('k')
                nw43_[int(9)] = d_233_v40_
                nw43_[int(10)] = (d_227_v34_)[default__.safeIndex(default__.fm12(d_192_v7_, d_186_v1_, d_195_globalState_), len(d_227_v34_))]
                nw43_[int(11)] = d_233_v40_
                nw43_[int(12)] = d_233_v40_
                nw43_[int(13)] = _dafny.CodePoint('c')
                d_235_v42_ = nw43_
                index26_ = default__.safeIndex(851, (d_235_v42_).length(0))
                (d_235_v42_)[index26_] = d_233_v40_
                d_236_v43_: _dafny.Seq
                d_236_v43_ = _dafny.SeqWithoutIsStrInference([d_229_v36_, d_229_v36_])
                d_237_v44_: _dafny.Map
                d_237_v44_ = _dafny.Map({(d_236_v43_ if d_186_v1_ else _dafny.SeqWithoutIsStrInference([d_229_v36_ for d_238_i7_ in range(default__.abs(477))])): d_186_v1_})
                index27_ = default__.safeIndex(851, (d_235_v42_).length(0))
                rhs23_ = d_232_v39_
                rhs24_ = d_186_v1_
                rhs25_ = d_233_v40_
                rhs26_ = ((d_237_v44_)[d_236_v43_] if (d_236_v43_) in (d_237_v44_) else (d_190_v5_)[default__.safeIndex(((d_230_v37_)[d_192_v7_] if (d_192_v7_) in (d_230_v37_) else default__.fm12(len(d_227_v34_), d_186_v1_, d_195_globalState_)), len(d_190_v5_))])
                rhs27_ = ((d_190_v5_) + (d_190_v5_)) + ((d_190_v5_) + (d_190_v5_))
                lhs20_ = d_195_globalState_
                lhs21_ = d_235_v42_
                lhs22_ = default__.safeIndex(851, (d_235_v42_).length(0))
                lhs23_ = d_195_globalState_
                d_232_v39_ = rhs23_
                lhs20_.f22 = rhs24_
                lhs21_[lhs22_] = rhs25_
                lhs23_.f27 = rhs26_
                d_190_v5_ = rhs27_
            elif source4_.is_DC31:
                d_239___mcc_h0_ = source4_.cf48
                d_240_cf48_ = d_239___mcc_h0_
                d_240_cf48_ = d_228_v35_
                index28_ = default__.safeIndex(241, (d_185_v0_).length(0))
                (d_185_v0_)[index28_] = 579
                d_241_v45_: str
                d_241_v45_ = _dafny.CodePoint('s')
                d_242_v46_: _dafny.Map
                d_242_v46_ = _dafny.Map({d_187_v2_: d_226_v33_})
                index29_ = default__.safeIndex(241, (d_185_v0_).length(0))
                rhs28_ = d_241_v45_
                rhs29_ = d_187_v2_
                rhs30_ = (0) - ((0) - ((d_193_v8_)[default__.safeIndex(d_192_v7_, len(d_193_v8_))]))
                rhs31_ = (d_192_v7_) * (len(d_227_v34_))
                rhs32_ = (d_190_v5_)[default__.safeIndex(len(d_242_v46_), len(d_190_v5_))]
                lhs24_ = d_195_globalState_
                lhs25_ = d_185_v0_
                lhs26_ = default__.safeIndex(241, (d_185_v0_).length(0))
                lhs27_ = d_195_globalState_
                lhs24_.f23 = rhs28_
                d_187_v2_ = rhs29_
                lhs25_[lhs26_] = rhs30_
                lhs27_.f24 = rhs31_
                d_186_v1_ = rhs32_
                (d_195_globalState_).f1 = d_192_v7_
                (d_195_globalState_).f1 = default__.safeModuloInt((0) - ((len(d_227_v34_)) + (104)), ((d_185_v0_)[default__.safeIndex(241, (d_185_v0_).length(0))] if d_186_v1_ else (d_185_v0_)[default__.safeIndex(241, (d_185_v0_).length(0))]))
            elif True:
                d_243___mcc_h1_ = source4_.cf49
                d_244_cf49_ = d_243___mcc_h1_
                d_245_v47_: _dafny.Map
                d_245_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umblovet")): False})
                (d_195_globalState_).f24 = len(d_245_v47_)
                d_246_v48_: D8
                d_246_v48_ = D8_DC21(d_186_v1_)
                (d_195_globalState_).f13 = (not(((d_246_v48_).cf27) and (d_186_v1_)) if d_186_v1_ else not (True) or (d_186_v1_))
                d_186_v1_ = not (True) or (d_186_v1_)
                d_247_v49_: C10
                nw44_ = C10()
                nw44_.ctor__(d_227_v34_, d_192_v7_)
                d_247_v49_ = nw44_
            d_248_v50_: _dafny.Map
            d_248_v50_ = _dafny.Map({d_187_v2_: d_187_v2_})
            d_248_v50_ = d_248_v50_
            d_249_v51_: C8
            nw45_ = C8()
            nw45_.ctor__()
            d_249_v51_ = nw45_
        elif True:
            d_250_v52_: _dafny.Array
            nw46_ = _dafny.Array(None, 24)
            d_250_v52_ = nw46_
            d_251_v53_: T1
            nw47_ = C4()
            nw47_.ctor__(d_186_v1_, d_186_v1_, d_186_v1_, d_186_v1_)
            d_251_v53_ = nw47_
            index30_ = default__.safeIndex(897, (d_250_v52_).length(0))
            (d_250_v52_)[index30_] = d_251_v53_
            index31_ = default__.safeIndex(897, (d_250_v52_).length(0))
            (d_250_v52_)[index31_] = d_251_v53_
            d_252_v54_: str
            d_252_v54_ = _dafny.CodePoint('q')
            d_253_v55_: _dafny.MultiSet
            d_253_v55_ = _dafny.MultiSet([d_186_v1_, d_186_v1_])
            d_192_v7_ = default__.safeDivisionInt(default__.fm27(d_192_v7_, d_252_v54_, d_192_v7_, d_195_globalState_), (d_253_v55_).cardinality)
            index32_ = default__.safeIndex(747, (d_187_v2_).length(0))
            (d_187_v2_)[index32_] = (d_186_v1_) and (d_186_v1_)
            index33_ = default__.safeIndex(747, (d_187_v2_).length(0))
            (d_187_v2_)[index33_] = (d_186_v1_) == (d_186_v1_)
            d_254_v56_: C8
            nw48_ = C8()
            nw48_.ctor__()
            d_254_v56_ = nw48_
            d_254_v56_ = d_254_v56_
            d_255_v57_: _dafny.Seq
            d_255_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgs"))
            d_256_v58_: _dafny.Map
            d_256_v58_ = _dafny.Map({d_192_v7_: D3_DC8(d_255_v57_)})
            d_257_v59_: D3
            d_257_v59_ = D3_DC8(_dafny.SeqWithoutIsStrInference([d_252_v54_ for d_258_i8_ in range(default__.abs(-34))]))
            d_256_v58_ = _dafny.Map({d_192_v7_: d_257_v59_})
        if (False if True else ((d_190_v5_).set(default__.safeIndex(d_192_v7_, len(d_190_v5_)), d_186_v1_)) < (d_190_v5_)):
            if not(d_186_v1_):
                d_259_v60_: _dafny.Seq
                d_259_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnhiys"))
                d_259_v60_ = d_259_v60_
                d_260_v61_: _dafny.Map
                d_260_v61_ = _dafny.Map({d_186_v1_: d_186_v1_})
                d_261_v62_: _dafny.Map
                d_261_v62_ = _dafny.Map({d_186_v1_: d_260_v61_})
                default__.m0((((d_261_v62_)[False] if (False) in (d_261_v62_) else default__.fm59(d_186_v1_, d_190_v5_, d_195_globalState_)) if d_186_v1_ else _dafny.Map({d_186_v1_: not(d_186_v1_)})), d_192_v7_, d_186_v1_, d_195_globalState_)
                d_262_v63_: str
                d_262_v63_ = _dafny.CodePoint('r')
                d_263_v64_: _dafny.Map
                d_263_v64_ = _dafny.Map({d_262_v63_: d_192_v7_})
                d_263_v64_ = (d_263_v64_).set(default__.fm42(d_192_v7_, 471, d_195_globalState_), d_192_v7_)
                d_264_v65_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.Seq({}), 27)
                d_264_v65_ = nw49_
                d_265_v66_: T1
                nw50_ = C4()
                nw50_.ctor__(d_186_v1_, d_186_v1_, d_186_v1_, d_186_v1_)
                d_265_v66_ = nw50_
                d_266_v67_: _dafny.Seq
                d_266_v67_ = _dafny.SeqWithoutIsStrInference([d_265_v66_, d_265_v66_])
                index34_ = default__.safeIndex(445, (d_264_v65_).length(0))
                (d_264_v65_)[index34_] = _dafny.SeqWithoutIsStrInference([(d_266_v67_)[default__.safeIndex(d_192_v7_, len(d_266_v67_))]])
                index35_ = default__.safeIndex(445, (d_264_v65_).length(0))
                rhs33_ = (d_266_v67_) + (d_266_v67_)
                rhs34_ = d_192_v7_
                lhs28_ = d_264_v65_
                lhs29_ = default__.safeIndex(445, (d_264_v65_).length(0))
                lhs30_ = d_195_globalState_
                lhs28_[lhs29_] = rhs33_
                lhs30_.f24 = rhs34_
                d_267_v68_: _dafny.Array
                nw51_ = _dafny.Array(None, 6)
                d_267_v68_ = nw51_
                d_268_v69_: T0
                nw52_ = C7()
                nw52_.ctor__(d_186_v1_, not(d_186_v1_))
                d_268_v69_ = nw52_
                index36_ = default__.safeIndex(744, (d_267_v68_).length(0))
                (d_267_v68_)[index36_] = d_268_v69_
                d_269_v70_: _dafny.Map
                d_269_v70_ = _dafny.Map({True: len(d_259_v60_)})
                d_270_v71_: _dafny.Map
                d_270_v71_ = _dafny.Map({(d_268_v69_).f29: d_269_v70_})
                index37_ = default__.safeIndex(744, (d_267_v68_).length(0))
                rhs35_ = (d_268_v69_ if d_186_v1_ else d_268_v69_)
                rhs36_ = (d_269_v70_) == ((d_269_v70_ if not((d_268_v69_).f29) else ((d_270_v71_)[(d_268_v69_).f29] if ((d_268_v69_).f29) in (d_270_v71_) else _dafny.Map({d_268_v69_.f28: d_192_v7_}))))
                lhs31_ = d_267_v68_
                lhs32_ = default__.safeIndex(744, (d_267_v68_).length(0))
                lhs33_ = d_195_globalState_
                lhs31_[lhs32_] = rhs35_
                lhs33_.f13 = rhs36_
            elif True:
                d_271_v72_: _dafny.Array
                nw53_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_271_v72_ = nw53_
                index38_ = default__.safeIndex(292, (d_271_v72_).length(0))
                (d_271_v72_)[index38_] = (default__.fm80(d_186_v1_, d_195_globalState_)).cf7
                d_272_v73_: _dafny.Seq
                d_272_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_273_v74_: _dafny.Map
                d_273_v74_ = _dafny.Map({d_272_v73_: d_186_v1_})
                d_274_v75_: _dafny.Map
                d_274_v75_ = _dafny.Map({d_186_v1_: d_186_v1_})
                index39_ = default__.safeIndex(56, (d_187_v2_).length(0))
                (d_187_v2_)[index39_] = ((d_273_v74_)[d_272_v73_] if (d_272_v73_) in (d_273_v74_) else default__.fm10(d_186_v1_, d_192_v7_, ((d_274_v75_)[d_186_v1_] if (d_186_v1_) in (d_274_v75_) else False), d_195_globalState_))
                index40_ = default__.safeIndex(104, (d_189_v4_).length(0))
                (d_189_v4_)[index40_] = d_187_v2_
                d_275_v76_: str
                d_275_v76_ = _dafny.CodePoint('q')
                index41_ = default__.safeIndex(292, (d_271_v72_).length(0))
                index42_ = default__.safeIndex(56, (d_187_v2_).length(0))
                index43_ = default__.safeIndex(104, (d_189_v4_).length(0))
                rhs37_ = d_275_v76_
                rhs38_ = d_186_v1_
                rhs39_ = d_187_v2_
                lhs34_ = d_271_v72_
                lhs35_ = default__.safeIndex(292, (d_271_v72_).length(0))
                lhs36_ = d_187_v2_
                lhs37_ = default__.safeIndex(56, (d_187_v2_).length(0))
                lhs38_ = d_189_v4_
                lhs39_ = default__.safeIndex(104, (d_189_v4_).length(0))
                lhs34_[lhs35_] = rhs37_
                lhs36_[lhs37_] = rhs38_
                lhs38_[lhs39_] = rhs39_
                d_192_v7_ = default__.safeDivisionInt(d_192_v7_, d_192_v7_)
                (d_195_globalState_).f27 = d_186_v1_
                d_276_v77_: _dafny.Seq
                d_276_v77_ = _dafny.SeqWithoutIsStrInference([d_185_v0_])
                d_277_v78_: T0
                nw54_ = C11()
                nw54_.ctor__(d_276_v77_, (d_187_v2_)[default__.safeIndex(56, (d_187_v2_).length(0))], d_186_v1_)
                d_277_v78_ = nw54_
                d_278_v79_: D8
                d_278_v79_ = D8_DC20(d_277_v78_)
                (d_195_globalState_).f1 = ((d_192_v7_ if (d_187_v2_)[default__.safeIndex(56, (d_187_v2_).length(0))] else d_192_v7_)) + (((D9_DC24(d_275_v76_, (d_187_v2_)[default__.safeIndex(56, (d_187_v2_).length(0))], 281, d_278_v79_, (d_277_v78_).f29)).cf31) - (347))
                d_279_v80_: _dafny.Map
                d_279_v80_ = _dafny.Map({d_192_v7_: d_277_v78_.f28})
                d_280_v81_: _dafny.Map
                d_280_v81_ = _dafny.Map({d_279_v80_: default__.fm38(d_195_globalState_)})
                rhs40_ = d_192_v7_
                rhs41_ = d_280_v81_
                rhs42_ = d_185_v0_
                lhs40_ = d_195_globalState_
                lhs40_.f24 = rhs40_
                d_280_v81_ = rhs41_
                d_185_v0_ = rhs42_
            d_281_v82_: _dafny.Seq
            d_281_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftv"))
            d_282_v83_: D22
            d_282_v83_ = D22_DC59(d_186_v1_, d_281_v82_, d_192_v7_)
            (d_195_globalState_).f24 = default__.safeModuloInt((d_282_v83_).cf85, d_192_v7_)
            d_283_v84_: _dafny.Seq
            d_283_v84_ = _dafny.SeqWithoutIsStrInference([d_185_v0_])
            d_284_v85_: _dafny.MultiSet
            d_284_v85_ = _dafny.MultiSet([len((d_283_v84_).set(default__.safeIndex(d_192_v7_, len(d_283_v84_)), d_185_v0_))])
            d_285_v86_: _dafny.Map
            d_285_v86_ = _dafny.Map({len((d_193_v8_).set(default__.safeIndex(d_192_v7_, len(d_193_v8_)), d_192_v7_)): (d_284_v85_).cardinality})
            (d_195_globalState_).f22 = (d_192_v7_) != (((d_285_v86_)[d_192_v7_] if (d_192_v7_) in (d_285_v86_) else d_192_v7_))
            (d_195_globalState_).f21 = (d_192_v7_) <= (d_192_v7_)
            d_286_v87_: _dafny.Map
            d_286_v87_ = _dafny.Map({True: d_186_v1_})
            default__.m0((d_286_v87_) | (default__.fm59(d_186_v1_, d_190_v5_, d_195_globalState_)), d_192_v7_, d_186_v1_, d_195_globalState_)
        elif True:
            d_287_v88_: T0
            nw55_ = C9()
            nw55_.ctor__(d_186_v1_, d_186_v1_)
            d_287_v88_ = nw55_
            d_288_v89_: D8
            d_288_v89_ = D8_DC20(d_287_v88_)
            d_289_v90_: _dafny.Map
            d_289_v90_ = _dafny.Map({(0) - (d_192_v7_): d_287_v88_})
            d_290_v91_: _dafny.Seq
            d_290_v91_ = _dafny.SeqWithoutIsStrInference([(d_288_v89_).cf26, d_287_v88_, ((d_289_v90_)[d_192_v7_] if (d_192_v7_) in (d_289_v90_) else d_287_v88_)])
            d_287_v88_ = (d_290_v91_)[default__.safeIndex(d_192_v7_, len(d_290_v91_))]
            (d_287_v88_).f28 = d_287_v88_.f28
            d_291_v92_: C4
            nw56_ = C4()
            nw56_.ctor__((d_287_v88_).f29, True, not(default__.fm18(d_192_v7_, d_192_v7_, len(default__.fm47(d_195_globalState_)), d_195_globalState_)), (d_287_v88_).f29)
            d_291_v92_ = nw56_
            d_292_v93_: _dafny.Seq
            d_292_v93_ = _dafny.SeqWithoutIsStrInference([d_291_v92_, d_291_v92_, d_291_v92_])
            d_292_v93_ = d_292_v93_
            (d_195_globalState_).f22 = not(not(d_186_v1_))
            d_293_v94_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_293_v94_ = nw57_
            d_294_v95_: str
            d_294_v95_ = _dafny.CodePoint('t')
            d_295_v96_: _dafny.Seq
            d_295_v96_ = _dafny.SeqWithoutIsStrInference([d_294_v95_, _dafny.CodePoint('s'), d_294_v95_])
            index44_ = default__.safeIndex(224, (d_293_v94_).length(0))
            (d_293_v94_)[index44_] = d_295_v96_
            index45_ = default__.safeIndex(224, (d_293_v94_).length(0))
            (d_293_v94_)[index45_] = d_295_v96_
        d_296_v97_: _dafny.Map
        d_296_v97_ = _dafny.Map({d_192_v7_: d_186_v1_})
        def iife32_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in _dafny.IntegerRange(-394, 786):
                d_297_v98_: int = compr_28_
                if ((-394) <= (d_297_v98_)) and ((d_297_v98_) < (786)):
                    coll28_[default__.safeModuloInt(d_297_v98_, d_192_v7_)] = d_186_v1_
            return _dafny.Map(coll28_)
        d_296_v97_ = (d_296_v97_) | (iife32_()
        )
        d_298_i9_: int
        d_298_i9_ = 0
        with _dafny.label("0"):
            while not(d_186_v1_):
                with _dafny.c_label("0"):
                    if (d_298_i9_) >= (100):
                        raise _dafny.Break("0")
                    d_298_i9_ = (d_298_i9_) + (1)
                    d_299_v99_: _dafny.Set
                    d_299_v99_ = _dafny.Set({d_187_v2_, d_187_v2_, d_187_v2_, d_187_v2_, d_187_v2_})
                    d_300_v100_: _dafny.Set
                    d_300_v100_ = d_299_v99_
                    d_301_v101_: _dafny.Array
                    nw58_ = _dafny.Array(None, 10)
                    nw58_[int(0)] = _dafny.Set({d_187_v2_, d_187_v2_, d_187_v2_, d_187_v2_})
                    nw58_[int(1)] = d_299_v99_
                    nw58_[int(2)] = d_299_v99_
                    nw58_[int(3)] = ((d_300_v100_)) - (d_299_v99_)
                    nw58_[int(4)] = d_299_v99_
                    nw58_[int(5)] = d_299_v99_
                    nw58_[int(6)] = d_299_v99_
                    nw58_[int(7)] = d_299_v99_
                    nw58_[int(8)] = d_299_v99_
                    nw58_[int(9)] = d_299_v99_
                    d_301_v101_ = nw58_
                    index46_ = default__.safeIndex(669, (d_301_v101_).length(0))
                    (d_301_v101_)[index46_] = d_299_v99_
                    index47_ = default__.safeIndex(669, (d_301_v101_).length(0))
                    (d_301_v101_)[index47_] = ((d_299_v99_) - (d_299_v99_)) - (d_299_v99_)
                    d_302_v102_: T2
                    nw59_ = C3()
                    nw59_.ctor__(d_186_v1_, d_186_v1_, (d_190_v5_)[default__.safeIndex(d_192_v7_, len(d_190_v5_))], d_192_v7_)
                    d_302_v102_ = nw59_
                    nw60_ = C3()
                    nw60_.ctor__((331) != ((0) - ((0) - ((d_302_v102_).f39))), d_186_v1_, d_186_v1_, (d_302_v102_).f39)
                    d_302_v102_ = nw60_
                    d_303_v103_: _dafny.Seq
                    d_303_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxu"))
                    d_304_v104_: _dafny.Seq
                    d_304_v104_ = _dafny.SeqWithoutIsStrInference([d_303_v103_])
                    d_305_v105_: _dafny.Set
                    d_305_v105_ = _dafny.Set({d_191_v6_})
                    d_306_v106_: T0
                    nw61_ = C6()
                    nw61_.ctor__(d_189_v4_, d_304_v104_, not(d_186_v1_), ((d_296_v97_)[len(d_305_v105_)] if (len(d_305_v105_)) in (d_296_v97_) else d_186_v1_))
                    d_306_v106_ = nw61_
                    d_307_v107_: _dafny.Seq
                    d_307_v107_ = _dafny.SeqWithoutIsStrInference([d_306_v106_, d_306_v106_])
                    d_308_v108_: _dafny.MultiSet
                    d_308_v108_ = _dafny.MultiSet([d_306_v106_.f28, d_306_v106_.f28, True])
                    d_309_v109_: _dafny.Map
                    d_309_v109_ = _dafny.Map({d_303_v103_: (d_302_v102_).f39})
                    rhs43_ = (len(d_190_v5_) if not (True) or (d_306_v106_.f28) else ((d_308_v108_)[d_186_v1_] if (d_186_v1_) in (d_308_v108_) else ((d_309_v109_)[d_303_v103_] if (d_303_v103_) in (d_309_v109_) else (d_302_v102_).f39)))
                    rhs44_ = (d_192_v7_) <= ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncwjvnjw")))) * (d_192_v7_))
                    rhs45_ = (((d_307_v107_).set(default__.safeIndex(d_192_v7_, len(d_307_v107_)), d_306_v106_)) + (d_307_v107_) if d_306_v106_.f28 else d_307_v107_)
                    lhs41_ = d_195_globalState_
                    d_192_v7_ = rhs43_
                    lhs41_.f21 = rhs44_
                    d_307_v107_ = rhs45_
                    d_310_v110_: T1
                    nw62_ = C3()
                    nw62_.ctor__(d_186_v1_, d_186_v1_, (d_306_v106_).f29, d_192_v7_)
                    d_310_v110_ = nw62_
                    d_311_v111_: _dafny.Map
                    d_311_v111_ = _dafny.Map({len(d_191_v6_): (d_302_v102_).f39})
                    d_312_v112_: _dafny.Map
                    d_312_v112_ = _dafny.Map({d_192_v7_: d_311_v111_})
                    d_313_v113_: _dafny.Map
                    d_313_v113_ = _dafny.Map({d_310_v110_: d_312_v112_})
                    d_314_v114_: D18
                    d_314_v114_ = D18_DC47((((d_313_v113_)[d_310_v110_] if (d_310_v110_) in (d_313_v113_) else _dafny.Map({(d_302_v102_).f39: d_311_v111_}))) | (d_312_v112_))
                    d_315_v115_: D25
                    d_315_v115_ = D25_DC69(default__.fm8(d_195_globalState_), 262, (d_302_v102_).f39)
                    d_316_v116_: _dafny.MultiSet
                    d_316_v116_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_192_v7_, default__.fm25(d_192_v7_, (d_302_v102_).f39, d_195_globalState_)])), d_192_v7_, (d_302_v102_).f39])
                    d_317_v117_: D4
                    d_317_v117_ = D4_DC9(d_316_v116_)
                    d_318_v118_: _dafny.Seq
                    d_318_v118_ = _dafny.SeqWithoutIsStrInference([d_317_v117_, d_317_v117_])
                    d_319_v119_: _dafny.Set
                    d_319_v119_ = _dafny.Set({d_192_v7_, len(d_318_v118_)})
                    d_320_v120_: _dafny.Map
                    d_320_v120_ = _dafny.Map({617: d_319_v119_})
                    rhs46_ = d_314_v114_
                    rhs47_ = default__.fm81(d_306_v106_.f28, default__.safeDivisionInt((d_302_v102_).f39, d_192_v7_), d_195_globalState_)
                    rhs48_ = (d_316_v116_) != ((d_316_v116_) - (_dafny.MultiSet([(d_302_v102_).f39])))
                    rhs49_ = (d_302_v102_).fm30(((d_320_v120_)[12] if (12) in (d_320_v120_) else _dafny.Set({d_192_v7_})), d_195_globalState_)
                    lhs42_ = d_195_globalState_
                    d_314_v114_ = rhs46_
                    d_315_v115_ = rhs47_
                    d_186_v1_ = rhs48_
                    lhs42_.f1 = rhs49_
                    pass
            pass
        if d_186_v1_:
            d_321_v121_: _dafny.Map
            d_321_v121_ = _dafny.Map({d_191_v6_: d_186_v1_})
            d_322_v122_: _dafny.Map
            d_322_v122_ = _dafny.Map({((d_321_v121_)[d_191_v6_] if (d_191_v6_) in (d_321_v121_) else d_186_v1_): (d_192_v7_ if d_186_v1_ else (0) - (d_192_v7_))})
            d_322_v122_ = (d_322_v122_).set(((0) - (d_192_v7_)) != (d_192_v7_), (0) - (d_192_v7_))
            (d_195_globalState_).f22 = (d_190_v5_)[default__.safeIndex(d_192_v7_, len(d_190_v5_))]
            index48_ = default__.safeIndex(817, (d_187_v2_).length(0))
            (d_187_v2_)[index48_] = d_186_v1_
            index49_ = default__.safeIndex(817, (d_187_v2_).length(0))
            (d_187_v2_)[index49_] = default__.fm10(d_186_v1_, (d_193_v8_)[default__.safeIndex(411, len(d_193_v8_))], (d_186_v1_) == (d_186_v1_), d_195_globalState_)
            d_323_v123_: _dafny.MultiSet
            d_323_v123_ = _dafny.MultiSet([d_190_v5_])
            d_324_v124_: D15
            d_324_v124_ = D15_DC37(d_323_v123_)
            (d_195_globalState_).f1 = ((d_324_v124_).cf52).cardinality
            d_325_v125_: _dafny.Array
            nw63_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_325_v125_ = nw63_
            d_326_v126_: _dafny.Seq
            d_326_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reijvxd"))
            d_327_v127_: _dafny.Array
            nw64_ = _dafny.Array(None, 5)
            nw64_[int(0)] = (D22_DC60(len(d_326_v126_), d_186_v1_, (d_187_v2_)[default__.safeIndex(817, (d_187_v2_).length(0))], _dafny.SeqWithoutIsStrInference([len(d_190_v5_), d_192_v7_]))).cf89
            nw64_[int(1)] = d_193_v8_
            nw64_[int(2)] = _dafny.SeqWithoutIsStrInference([d_192_v7_ for d_328_i10_ in range(default__.abs(68))])
            nw64_[int(3)] = d_193_v8_
            nw64_[int(4)] = _dafny.SeqWithoutIsStrInference([d_192_v7_ for d_329_i11_ in range(default__.abs(455))])
            d_327_v127_ = nw64_
            index50_ = default__.safeIndex(554, (d_325_v125_).length(0))
            (d_325_v125_)[index50_] = d_327_v127_
            index51_ = default__.safeIndex(554, (d_325_v125_).length(0))
            (d_325_v125_)[index51_] = d_327_v127_
        elif True:
            (d_195_globalState_).f1 = default__.safeDivisionInt(d_192_v7_, len(d_191_v6_))
            d_330_v128_: _dafny.Seq
            d_330_v128_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njytyjmu"))
            (d_195_globalState_).f24 = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwcyiw"))) + ((d_330_v128_) + (d_330_v128_)))
            if (default__.safeModuloInt(d_192_v7_, d_192_v7_)) >= (d_192_v7_):
                d_331_v129_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.Map({}), 19)
                d_331_v129_ = nw65_
                d_332_v130_: _dafny.Map
                d_332_v130_ = _dafny.Map({d_186_v1_: _dafny.MultiSet([d_192_v7_])})
                index52_ = default__.safeIndex(836, (d_331_v129_).length(0))
                (d_331_v129_)[index52_] = (d_332_v130_) | (d_332_v130_)
                index53_ = default__.safeIndex(836, (d_331_v129_).length(0))
                (d_331_v129_)[index53_] = d_332_v130_
                d_333_v131_: D2
                d_333_v131_ = D2_DC5(d_186_v1_)
                d_334_v132_: _dafny.MultiSet
                d_334_v132_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_186_v1_]), d_190_v5_, d_190_v5_])
                (d_195_globalState_).f22 = default__.fm10((default__.fm29(d_334_v132_, d_192_v7_, d_195_globalState_) if (d_333_v131_).cf8 else (d_190_v5_)[default__.safeIndex(d_192_v7_, len(d_190_v5_))]), (len(d_191_v6_)) + (d_192_v7_), d_186_v1_, d_195_globalState_)
                d_335_v133_: C12
                nw66_ = C12()
                nw66_.ctor__()
                d_335_v133_ = nw66_
                d_336_v134_: str
                d_336_v134_ = _dafny.CodePoint('n')
                d_337_v135_: T0
                nw67_ = C5()
                nw67_.ctor__(d_186_v1_, d_336_v134_, d_186_v1_, d_186_v1_)
                d_337_v135_ = nw67_
                d_338_v136_: _dafny.Array
                nw68_ = _dafny.Array(None, 4)
                nw68_[int(0)] = d_337_v135_
                nw68_[int(1)] = d_337_v135_
                nw68_[int(2)] = (d_337_v135_ if (d_337_v135_).f29 else d_337_v135_)
                nw68_[int(3)] = d_337_v135_
                d_338_v136_ = nw68_
                index54_ = default__.safeIndex(326, (d_338_v136_).length(0))
                (d_338_v136_)[index54_] = d_337_v135_
                d_339_v137_: _dafny.MultiSet
                d_339_v137_ = _dafny.MultiSet([d_192_v7_, (0) - (d_192_v7_), (_dafny.MultiSet([d_186_v1_])).cardinality])
                d_340_v138_: D1
                d_340_v138_ = D1_DC2((d_193_v8_)[default__.safeIndex(d_192_v7_, len(d_193_v8_))], d_192_v7_, d_330_v128_, not(d_186_v1_))
                d_341_v139_: _dafny.Set
                d_341_v139_ = _dafny.Set({D1_DC2(default__.fm27(d_192_v7_, d_336_v134_, (d_339_v137_).cardinality, d_195_globalState_), d_192_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), False), d_340_v138_, d_340_v138_, d_340_v138_, d_340_v138_})
                index55_ = default__.safeIndex(326, (d_338_v136_).length(0))
                rhs50_ = d_337_v135_
                rhs51_ = len(d_341_v139_)
                rhs52_ = d_192_v7_
                lhs43_ = d_338_v136_
                lhs44_ = default__.safeIndex(326, (d_338_v136_).length(0))
                lhs45_ = d_195_globalState_
                lhs46_ = d_195_globalState_
                lhs43_[lhs44_] = rhs50_
                lhs45_.f24 = rhs51_
                lhs46_.f24 = rhs52_
                d_342_v140_: _dafny.Seq
                d_342_v140_ = _dafny.SeqWithoutIsStrInference([d_330_v128_])
                d_343_v141_: _dafny.Map
                d_343_v141_ = _dafny.Map({(d_342_v140_)[default__.safeIndex(d_192_v7_, len(d_342_v140_))]: (d_337_v135_).f29})
                d_344_v142_: D17
                d_344_v142_ = D17_DC43(d_343_v141_)
                d_345_v143_: _dafny.Set
                d_345_v143_ = _dafny.Set({d_344_v142_})
                d_345_v143_ = (d_345_v143_).intersection(d_345_v143_)
            elif True:
                (d_195_globalState_).f1 = d_192_v7_
                (d_195_globalState_).f1 = d_192_v7_
                d_346_v144_: _dafny.Map
                d_346_v144_ = _dafny.Map({d_186_v1_: not(d_186_v1_)})
                default__.m0(d_346_v144_, d_192_v7_, False, d_195_globalState_)
                d_347_v145_: C2
                nw69_ = C2()
                nw69_.ctor__()
                d_347_v145_ = nw69_
                d_348_v146_: C7
                nw70_ = C7()
                nw70_.ctor__(d_186_v1_, (d_186_v1_) and (False))
                d_348_v146_ = nw70_
            (d_195_globalState_).f21 = d_186_v1_
            d_349_v147_: str
            d_349_v147_ = _dafny.CodePoint('k')
            d_330_v128_ = _dafny.SeqWithoutIsStrInference([d_349_v147_, d_349_v147_])
        d_350_v148_: C7
        nw71_ = C7()
        nw71_.ctor__(d_186_v1_, d_186_v1_)
        d_350_v148_ = nw71_
        d_351_v149_: _dafny.MultiSet
        d_351_v149_ = _dafny.MultiSet([d_350_v148_])
        d_352_v150_: D23
        d_352_v150_ = D23_DC64(d_192_v7_, d_192_v7_, d_186_v1_, d_192_v7_, (d_351_v149_).cardinality)
        d_353_v151_: _dafny.Map
        d_353_v151_ = _dafny.Map({d_186_v1_: d_192_v7_})
        hi2_ = ((d_353_v151_)[d_186_v1_] if (d_186_v1_) in (d_353_v151_) else d_192_v7_)
        for d_354_i12_ in range((d_352_v150_).cf94, hi2_):
            (d_195_globalState_).f1 = d_192_v7_
            d_355_v152_: _dafny.Seq
            d_355_v152_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgftpxt"))
            d_356_v153_: _dafny.Map
            d_356_v153_ = _dafny.Map({d_355_v152_: d_186_v1_})
            d_357_v154_: _dafny.Seq
            d_357_v154_ = _dafny.SeqWithoutIsStrInference([d_356_v153_, d_356_v153_, d_356_v153_, d_356_v153_, d_356_v153_])
            d_358_v155_: D17
            d_358_v155_ = D17_DC43((d_356_v153_ if d_186_v1_ else (d_357_v154_)[default__.safeIndex(d_192_v7_, len(d_357_v154_))]))
            d_359_v156_: C8
            nw72_ = C8()
            nw72_.ctor__()
            d_359_v156_ = nw72_
            d_360_v157_: _dafny.Array
            nw73_ = _dafny.Array(None, 3)
            d_360_v157_ = nw73_
            pat_let_tv2_ = d_360_v157_
            d_361_v158_: _dafny.Array
            nw74_ = _dafny.Array(None, 9)
            nw74_[int(0)] = d_360_v157_
            nw74_[int(1)] = d_360_v157_
            nw74_[int(2)] = d_360_v157_
            def iife33_(_pat_let2_0):
                def iife34_(d_362_dt__update__tmp_h1_):
                    def iife35_(_pat_let3_0):
                        def iife36_(d_363_dt__update_hcf119_h0_):
                            return D29_DC78(d_363_dt__update_hcf119_h0_)
                        return iife36_(_pat_let3_0)
                    return iife35_(pat_let_tv2_)
                return iife34_(_pat_let2_0)
            nw74_[int(3)] = (iife33_(D29_DC78(d_360_v157_))).cf119
            nw74_[int(4)] = d_360_v157_
            nw74_[int(5)] = d_360_v157_
            nw74_[int(6)] = d_360_v157_
            nw74_[int(7)] = d_360_v157_
            nw74_[int(8)] = d_360_v157_
            d_361_v158_ = nw74_
            index56_ = default__.safeIndex(150, (d_361_v158_).length(0))
            (d_361_v158_)[index56_] = d_360_v157_
            d_364_v159_: _dafny.Seq
            d_364_v159_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))])
            d_365_v160_: C6
            nw75_ = C6()
            nw75_.ctor__(d_189_v4_, d_364_v159_, d_186_v1_, ((d_356_v153_)[d_355_v152_] if (d_355_v152_) in (d_356_v153_) else d_186_v1_))
            d_365_v160_ = nw75_
            d_366_v161_: _dafny.Map
            d_366_v161_ = _dafny.Map({(D30_DC80(d_365_v160_)).cf125: d_186_v1_})
            d_367_v162_: str
            d_367_v162_ = _dafny.CodePoint('u')
            index57_ = default__.safeIndex(150, (d_361_v158_).length(0))
            rhs53_ = d_358_v155_
            rhs54_ = d_359_v156_
            rhs55_ = d_360_v157_
            rhs56_ = (len((d_366_v161_ if d_186_v1_ else _dafny.Map({d_365_v160_: d_186_v1_})))) >= (d_192_v7_)
            rhs57_ = ((d_367_v162_ if d_186_v1_ else d_367_v162_) if d_186_v1_ else default__.fm42(29, d_354_i12_, d_195_globalState_))
            lhs47_ = d_361_v158_
            lhs48_ = default__.safeIndex(150, (d_361_v158_).length(0))
            lhs49_ = d_195_globalState_
            lhs50_ = d_195_globalState_
            d_358_v155_ = rhs53_
            d_359_v156_ = rhs54_
            lhs47_[lhs48_] = rhs55_
            lhs49_.f21 = rhs56_
            lhs50_.f23 = rhs57_
            d_353_v151_ = (d_353_v151_).set(d_186_v1_, d_354_i12_)
            nw76_ = _dafny.Array(int(0), 19)
            d_185_v0_ = nw76_
        d_368_v163_: C12
        nw77_ = C12()
        nw77_.ctor__()
        d_368_v163_ = nw77_
        d_368_v163_ = d_368_v163_
        d_369_v164_: _dafny.Set
        d_369_v164_ = _dafny.Set({d_192_v7_})
        d_370_v165_: _dafny.MultiSet
        d_370_v165_ = _dafny.MultiSet([len(d_369_v164_), d_192_v7_, d_192_v7_])
        d_371_v166_: _dafny.MultiSet
        d_371_v166_ = _dafny.MultiSet([(d_370_v165_).cardinality, d_192_v7_, ((d_353_v151_)[d_186_v1_] if (d_186_v1_) in (d_353_v151_) else d_192_v7_), (0) - (d_192_v7_)])
        d_372_v167_: D4
        d_372_v167_ = D4_DC9(d_370_v165_)
        source5_ = d_372_v167_
        if source5_.is_DC10:
            d_373___mcc_h2_ = source5_.cf13
            d_374___mcc_h3_ = source5_.cf14
            d_375___mcc_h4_ = source5_.cf15
            d_376_cf15_ = d_375___mcc_h4_
            d_377_cf14_ = d_374___mcc_h3_
            d_378_cf13_ = d_373___mcc_h2_
            (d_195_globalState_).f24 = 694
            d_379_v168_: C9
            nw78_ = C9()
            nw78_.ctor__(d_186_v1_, d_378_cf13_)
            d_379_v168_ = nw78_
            d_380_v169_: _dafny.Map
            d_380_v169_ = _dafny.Map({d_186_v1_: d_379_v168_})
            d_380_v169_ = (d_380_v169_).set(d_186_v1_, d_379_v168_)
            d_192_v7_ = d_192_v7_
            (d_195_globalState_).f21 = default__.fm13(d_195_globalState_)
        elif source5_.is_DC9:
            d_381___mcc_h5_ = source5_.cf12
            d_382_cf12_ = d_381___mcc_h5_
            d_383_v170_: str
            d_383_v170_ = _dafny.CodePoint('o')
            d_384_v171_: _dafny.Map
            d_384_v171_ = _dafny.Map({d_186_v1_: d_383_v170_})
            d_384_v171_ = (d_384_v171_).set(d_186_v1_, d_383_v170_)
            d_192_v7_ = 984
            d_385_v172_: T0
            nw79_ = C1()
            nw79_.ctor__(False, d_186_v1_)
            d_385_v172_ = nw79_
            d_386_v173_: _dafny.Seq
            d_386_v173_ = _dafny.SeqWithoutIsStrInference([d_385_v172_])
            d_387_v174_: D30
            d_387_v174_ = D30_DC82((0) - (len(d_386_v173_)), d_192_v7_, d_385_v172_)
            d_388_v175_: _dafny.Array
            nw80_ = _dafny.Array(None, 4)
            nw80_[int(0)] = d_387_v174_
            nw80_[int(1)] = d_387_v174_
            nw80_[int(2)] = D30_DC82(len(d_191_v6_), d_192_v7_, d_385_v172_)
            nw80_[int(3)] = d_387_v174_
            d_388_v175_ = nw80_
            index58_ = default__.safeIndex(209, (d_388_v175_).length(0))
            (d_388_v175_)[index58_] = d_387_v174_
            index59_ = default__.safeIndex(209, (d_388_v175_).length(0))
            (d_388_v175_)[index59_] = d_387_v174_
            index60_ = default__.safeIndex(962, (d_187_v2_).length(0))
            (d_187_v2_)[index60_] = d_385_v172_.f28
            index61_ = default__.safeIndex(962, (d_187_v2_).length(0))
            (d_187_v2_)[index61_] = not(d_385_v172_.f28)
        elif True:
            d_389___mcc_h6_ = source5_.cf16
            d_390_cf16_ = d_389___mcc_h6_
            d_391_v176_: C3
            nw81_ = C3()
            nw81_.ctor__(default__.fm13(d_195_globalState_), d_186_v1_, d_186_v1_, d_192_v7_)
            d_391_v176_ = nw81_
            d_392_v177_: _dafny.Array
            nw82_ = _dafny.Array(None, 14)
            nw82_[int(0)] = d_391_v176_
            nw82_[int(1)] = d_391_v176_
            nw82_[int(2)] = d_391_v176_
            nw82_[int(3)] = d_391_v176_
            nw82_[int(4)] = d_391_v176_
            nw82_[int(5)] = d_391_v176_
            nw82_[int(6)] = d_391_v176_
            nw82_[int(7)] = d_391_v176_
            nw82_[int(8)] = d_391_v176_
            nw82_[int(9)] = d_391_v176_
            nw82_[int(10)] = d_391_v176_
            nw82_[int(11)] = d_391_v176_
            nw82_[int(12)] = d_391_v176_
            nw82_[int(13)] = d_391_v176_
            d_392_v177_ = nw82_
            d_393_v178_: D21
            d_393_v178_ = D21_DC54(d_392_v177_)
            source6_ = d_393_v178_
            if source6_.is_DC55:
                d_394___mcc_h7_ = source6_.cf80
                d_395_cf80_ = d_394___mcc_h7_
                (d_195_globalState_).f21 = (-651) < (d_192_v7_)
                d_396_v179_: _dafny.Seq
                d_396_v179_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rk"))
                d_397_v180_: C10
                nw83_ = C10()
                nw83_.ctor__(d_396_v179_, d_192_v7_)
                d_397_v180_ = nw83_
                d_398_v181_: _dafny.Array
                nw84_ = _dafny.Array(None, 24)
                nw84_[int(0)] = d_397_v180_
                nw84_[int(1)] = d_397_v180_
                nw84_[int(2)] = d_397_v180_
                nw84_[int(3)] = d_397_v180_
                nw84_[int(4)] = d_397_v180_
                nw84_[int(5)] = d_397_v180_
                nw84_[int(6)] = d_397_v180_
                nw84_[int(7)] = d_397_v180_
                nw84_[int(8)] = d_397_v180_
                nw84_[int(9)] = d_397_v180_
                nw84_[int(10)] = d_397_v180_
                nw84_[int(11)] = d_397_v180_
                nw84_[int(12)] = d_397_v180_
                nw84_[int(13)] = d_397_v180_
                nw84_[int(14)] = d_397_v180_
                nw84_[int(15)] = d_397_v180_
                nw84_[int(16)] = d_397_v180_
                nw84_[int(17)] = d_397_v180_
                nw84_[int(18)] = d_397_v180_
                nw84_[int(19)] = d_397_v180_
                nw84_[int(20)] = d_397_v180_
                nw84_[int(21)] = d_397_v180_
                nw84_[int(22)] = d_397_v180_
                nw84_[int(23)] = d_397_v180_
                d_398_v181_ = nw84_
                d_399_v182_: _dafny.Map
                d_399_v182_ = _dafny.Map({(d_192_v7_) + (d_192_v7_): d_398_v181_})
                d_400_v183_: _dafny.MultiSet
                d_400_v183_ = _dafny.MultiSet([d_396_v179_, ((d_397_v180_).f31) + ((d_397_v180_).f31)])
                d_401_v184_: D30
                d_401_v184_ = D30_DC83((d_391_v176_).f40)
                d_402_v185_: _dafny.Seq
                d_402_v185_ = _dafny.SeqWithoutIsStrInference([d_401_v184_])
                d_403_v186_: _dafny.Seq
                d_403_v186_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_401_v184_, d_401_v184_]), _dafny.SeqWithoutIsStrInference([D30_DC83((d_391_v176_).f40) for d_404_i13_ in range(default__.abs(835))]), d_402_v185_])
                d_405_v187_: _dafny.MultiSet
                d_405_v187_ = _dafny.MultiSet([(d_403_v186_)[default__.safeIndex(d_192_v7_, len(d_403_v186_))], d_402_v185_])
                rhs58_ = d_192_v7_
                rhs59_ = (0) - ((((d_397_v180_).f32) * ((d_397_v180_).f32) if (d_405_v187_) != (d_405_v187_) else len((d_190_v5_) + (d_190_v5_))))
                rhs60_ = d_399_v182_
                rhs61_ = default__.fm20(d_186_v1_, d_195_globalState_)
                lhs51_ = d_195_globalState_
                lhs51_.f24 = rhs58_
                d_192_v7_ = rhs59_
                d_399_v182_ = rhs60_
                d_400_v183_ = rhs61_
                d_406_v188_: D1
                d_406_v188_ = D1_DC2(d_192_v7_, (d_397_v180_).f32, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_407_i14_ in range(default__.abs(736))]), (d_391_v176_).f40)
                d_408_v189_: _dafny.Map
                d_408_v189_ = _dafny.Map({d_186_v1_: d_406_v188_})
                d_409_v190_: _dafny.Seq
                d_410_v191_: bool
                d_411_v192_: bool
                d_412_v193_: bool
                out0_: _dafny.Seq
                out1_: bool
                out2_: bool
                out3_: bool
                out0_, out1_, out2_, out3_ = (d_397_v180_).m7(d_192_v7_, (d_397_v180_).f32, ((0) - (d_192_v7_)) <= (d_192_v7_), d_408_v189_, d_195_globalState_)
                d_409_v190_ = out0_
                d_410_v191_ = out1_
                d_411_v192_ = out2_
                d_412_v193_ = out3_
                d_413_v195_: _dafny.Map
                d_413_v195_ = _dafny.Map({(d_391_v176_).f40: d_369_v164_})
                d_414_v196_: bool
                d_415_v197_: _dafny.Map
                out4_: bool
                out5_: _dafny.Map
                def iife37_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(216, 716):
                        d_416_v194_: int = compr_29_
                        if ((216) <= (d_416_v194_)) and ((d_416_v194_) < (716)):
                            coll29_ = coll29_.union(_dafny.Set([(d_416_v194_) * ((d_397_v180_).f32)]))
                    return _dafny.Set(coll29_)
                out4_, out5_ = (d_368_v163_).m3((iife37_()
                ).intersection(((d_413_v195_)[default__.fm38(d_195_globalState_)] if (default__.fm38(d_195_globalState_)) in (d_413_v195_) else d_369_v164_)), default__.fm3((d_391_v176_).f40, (d_391_v176_).f40, d_192_v7_, d_195_globalState_), d_195_globalState_)
                d_414_v196_ = out4_
                d_415_v197_ = out5_
            elif source6_.is_DC56:
                (d_195_globalState_).f22 = (d_391_v176_).f40
                d_417_v198_: C3
                nw85_ = C3()
                nw85_.ctor__(not((d_391_v176_).f40), (d_391_v176_).f40, (d_391_v176_).f40, (d_193_v8_)[default__.safeIndex(253, len(d_193_v8_))])
                d_417_v198_ = nw85_
                d_418_v199_: _dafny.Array
                nw86_ = _dafny.Array(D10.default()(), 16)
                d_418_v199_ = nw86_
                d_419_v200_: _dafny.Map
                d_419_v200_ = _dafny.Map({d_192_v7_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))})
                d_420_v201_: _dafny.Map
                d_420_v201_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_418_v199_]): ((d_419_v200_)[d_192_v7_] if (d_192_v7_) in (d_419_v200_) else default__.fm8(d_195_globalState_))})
                d_421_v202_: _dafny.Seq
                d_421_v202_ = _dafny.SeqWithoutIsStrInference([d_418_v199_, d_418_v199_])
                d_422_v203_: _dafny.Set
                d_422_v203_ = _dafny.Set({d_185_v0_, d_185_v0_, d_185_v0_})
                d_420_v201_ = (d_420_v201_).set((d_421_v202_) + (d_421_v202_), len(d_422_v203_))
                d_423_v204_: str
                d_423_v204_ = _dafny.CodePoint('f')
                d_424_v205_: int
                d_425_v206_: bool
                d_426_v207_: _dafny.Seq
                d_427_v208_: int
                out6_: int
                out7_: bool
                out8_: _dafny.Seq
                out9_: int
                out6_, out7_, out8_, out9_ = (d_391_v176_).m18(d_423_v204_, _dafny.SeqWithoutIsStrInference([d_423_v204_ for d_428_i15_ in range(default__.abs(-751))]), d_195_globalState_)
                d_424_v205_ = out6_
                d_425_v206_ = out7_
                d_426_v207_ = out8_
                d_427_v208_ = out9_
            elif source6_.is_DC54:
                d_429___mcc_h8_ = source6_.cf79
                d_430_cf79_ = d_429___mcc_h8_
                rhs62_ = (d_391_v176_).f40
                rhs63_ = not((d_391_v176_).f40)
                rhs64_ = d_192_v7_
                lhs52_ = d_195_globalState_
                lhs53_ = d_195_globalState_
                lhs52_.f22 = rhs62_
                d_186_v1_ = rhs63_
                lhs53_.f1 = rhs64_
                d_431_v209_: _dafny.Seq
                d_432_v210_: _dafny.Array
                out10_: _dafny.Seq
                out11_: _dafny.Array
                out10_, out11_ = (d_368_v163_).m4(d_195_globalState_)
                d_431_v209_ = out10_
                d_432_v210_ = out11_
                d_433_v211_: C2
                nw87_ = C2()
                nw87_.ctor__()
                d_433_v211_ = nw87_
                d_434_v212_: _dafny.Map
                d_434_v212_ = _dafny.Map({d_368_v163_: d_186_v1_})
                d_435_v213_: _dafny.Seq
                d_435_v213_ = _dafny.SeqWithoutIsStrInference([d_434_v212_, d_434_v212_])
                d_436_v214_: _dafny.Seq
                d_436_v214_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wc"))
                (d_195_globalState_).f1 = default__.safeModuloInt(d_192_v7_, default__.safeDivisionInt(len((d_435_v213_).set(default__.safeIndex(d_192_v7_, len(d_435_v213_)), d_434_v212_)), len(d_436_v214_)))
            elif True:
                d_437___mcc_h9_ = source6_.cf81
                d_438_cf81_ = d_437___mcc_h9_
                d_439_v215_: str
                d_439_v215_ = _dafny.CodePoint('q')
                d_440_v216_: _dafny.Seq
                d_440_v216_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwyssicy"))
                d_441_v217_: int
                d_442_v218_: bool
                d_443_v219_: _dafny.Seq
                d_444_v220_: int
                out12_: int
                out13_: bool
                out14_: _dafny.Seq
                out15_: int
                out12_, out13_, out14_, out15_ = (d_391_v176_).m18(d_439_v215_, d_440_v216_, d_195_globalState_)
                d_441_v217_ = out12_
                d_442_v218_ = out13_
                d_443_v219_ = out14_
                d_444_v220_ = out15_
                d_445_v221_: _dafny.Array
                def lambda7_(d_446_v164_):
                    def lambda8_(d_447_i16_):
                        return (d_446_v164_) - (_dafny.Set({14}))

                    return lambda8_

                init4_ = lambda7_(d_369_v164_)
                nw88_ = _dafny.Array(None, 18)
                for i0_4_ in range(nw88_.length(0)):
                    nw88_[i0_4_] = init4_(i0_4_)
                d_445_v221_ = nw88_
                index62_ = default__.safeIndex(894, (d_445_v221_).length(0))
                (d_445_v221_)[index62_] = (d_369_v164_).intersection(_dafny.Set({len(d_440_v216_)}))
                index63_ = default__.safeIndex(894, (d_445_v221_).length(0))
                (d_445_v221_)[index63_] = ((d_369_v164_ if False else d_369_v164_)) - (d_369_v164_)
                (d_195_globalState_).f27 = not((d_391_v176_).f40)
                d_448_v222_: C8
                nw89_ = C8()
                nw89_.ctor__()
                d_448_v222_ = nw89_
            d_369_v164_ = d_369_v164_
            if d_186_v1_:
                d_186_v1_ = (d_190_v5_)[default__.safeIndex(d_192_v7_, len(d_190_v5_))]
                d_449_v223_: _dafny.Array
                nw90_ = _dafny.Array(None, 12)
                d_449_v223_ = nw90_
                d_450_v224_: C5
                nw91_ = C5()
                nw91_.ctor__(d_186_v1_, default__.fm42(d_192_v7_, d_192_v7_, d_195_globalState_), d_186_v1_, (d_391_v176_).f40)
                d_450_v224_ = nw91_
                index64_ = default__.safeIndex(462, (d_449_v223_).length(0))
                (d_449_v223_)[index64_] = d_450_v224_
                index65_ = default__.safeIndex(462, (d_449_v223_).length(0))
                (d_449_v223_)[index65_] = d_450_v224_
                d_353_v151_ = (d_353_v151_).set(d_450_v224_.f35, d_192_v7_)
                (d_195_globalState_).f21 = d_450_v224_.f35
                d_451_v225_: C2
                nw92_ = C2()
                nw92_.ctor__()
                d_451_v225_ = nw92_
                d_452_v226_: _dafny.Seq
                d_452_v226_ = _dafny.SeqWithoutIsStrInference([d_190_v5_, d_190_v5_, d_190_v5_, d_190_v5_])
                index66_ = default__.safeIndex(112, (d_187_v2_).length(0))
                (d_187_v2_)[index66_] = default__.fm29(_dafny.MultiSet(d_452_v226_), d_192_v7_, d_195_globalState_)
                index67_ = default__.safeIndex(112, (d_187_v2_).length(0))
                rhs65_ = d_451_v225_
                rhs66_ = (d_391_v176_).f40
                lhs54_ = d_187_v2_
                lhs55_ = default__.safeIndex(112, (d_187_v2_).length(0))
                d_451_v225_ = rhs65_
                lhs54_[lhs55_] = rhs66_
            elif True:
                (d_195_globalState_).f24 = d_192_v7_
                index68_ = default__.safeIndex(999, (d_185_v0_).length(0))
                (d_185_v0_)[index68_] = d_192_v7_
                d_453_v227_: _dafny.Seq
                d_453_v227_ = _dafny.SeqWithoutIsStrInference([d_185_v0_])
                d_454_v228_: _dafny.Seq
                d_454_v228_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnmt"))
                d_455_v229_: _dafny.Map
                d_455_v229_ = _dafny.Map({d_454_v228_: d_192_v7_})
                d_456_v230_: T0
                nw93_ = C7()
                nw93_.ctor__((d_391_v176_).f40, ((d_296_v97_)[d_192_v7_] if (d_192_v7_) in (d_296_v97_) else (d_391_v176_).f40))
                d_456_v230_ = nw93_
                d_457_v231_: _dafny.Set
                d_457_v231_ = _dafny.Set({d_456_v230_})
                index69_ = default__.safeIndex(999, (d_185_v0_).length(0))
                rhs67_ = (d_192_v7_) == ((d_192_v7_) - (d_192_v7_))
                rhs68_ = (-430) + (((d_455_v229_)[_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_458_i17_ in range(default__.abs(-510))])] if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_459_i17_ in range(default__.abs(-510))])) in (d_455_v229_) else d_192_v7_))
                rhs69_ = d_453_v227_
                rhs70_ = len((((d_454_v228_) + (d_454_v228_)).set(default__.safeIndex(d_192_v7_, len((d_454_v228_) + (d_454_v228_))), default__.fm42(len(d_457_v231_), d_192_v7_, d_195_globalState_))) + (d_454_v228_))
                lhs56_ = d_195_globalState_
                lhs57_ = d_185_v0_
                lhs58_ = default__.safeIndex(999, (d_185_v0_).length(0))
                lhs59_ = d_195_globalState_
                lhs56_.f22 = rhs67_
                lhs57_[lhs58_] = rhs68_
                d_453_v227_ = rhs69_
                lhs59_.f24 = rhs70_
                d_460_v232_: C0
                nw94_ = C0()
                nw94_.ctor__()
                d_460_v232_ = nw94_
                index70_ = default__.safeIndex(999, (d_185_v0_).length(0))
                (d_185_v0_)[index70_] = ((_dafny.MultiSet(d_193_v8_)).cardinality) - (94)
                (d_195_globalState_).f17 = (_dafny.SeqWithoutIsStrInference([d_192_v7_ for d_461_i18_ in range(default__.abs(684))])) + (d_193_v8_)
            (d_195_globalState_).f24 = (((0) - (len(d_191_v6_))) - (d_192_v7_) if d_186_v1_ else d_192_v7_)
        if True:
            d_462_v233_: str
            d_462_v233_ = _dafny.CodePoint('k')
            d_463_v234_: _dafny.Map
            d_463_v234_ = _dafny.Map({default__.fm82(d_462_v233_, not(d_186_v1_), d_192_v7_, d_195_globalState_): d_187_v2_})
            d_464_v235_: _dafny.Seq
            d_464_v235_ = d_193_v8_
            d_465_v236_: _dafny.Seq
            d_465_v236_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqtenen"))
            d_463_v234_ = (d_463_v234_).set(d_464_v235_, (d_187_v2_ if (d_190_v5_)[default__.safeIndex(len((d_465_v236_).set(default__.safeIndex(d_192_v7_, len(d_465_v236_)), d_462_v233_)), len(d_190_v5_))] else d_187_v2_))
            d_466_v237_: _dafny.Map
            d_466_v237_ = _dafny.Map({932: d_192_v7_})
            d_467_v238_: _dafny.Map
            d_467_v238_ = _dafny.Map({_dafny.Map({d_186_v1_: d_186_v1_}): d_466_v237_})
            d_468_v240_: _dafny.Map
            d_468_v240_ = _dafny.Map({d_186_v1_: d_186_v1_})
            d_469_v241_: D31
            d_469_v241_ = D31_DC84(d_466_v237_)
            d_470_v242_: _dafny.Array
            nw95_ = _dafny.Array(None, 14)
            nw95_[int(0)] = (d_467_v238_) | (d_467_v238_)
            nw95_[int(1)] = (d_467_v238_) | (d_467_v238_)
            nw95_[int(2)] = d_467_v238_
            nw95_[int(3)] = _dafny.Map({_dafny.Map({d_186_v1_: not(d_186_v1_)}): d_466_v237_})
            nw95_[int(4)] = d_467_v238_
            def iife38_():
                coll30_ = _dafny.Map()
                compr_30_: _dafny.Map
                for compr_30_ in (default__.fm83(d_462_v233_, not(d_186_v1_), d_195_globalState_)).Elements:
                    d_471_v239_: _dafny.Map = compr_30_
                    if (d_471_v239_) in (default__.fm83(d_462_v233_, not(d_186_v1_), d_195_globalState_)):
                        coll30_[d_471_v239_] = d_466_v237_
                return _dafny.Map(coll30_)
            nw95_[int(5)] = iife38_()
            
            nw95_[int(6)] = d_467_v238_
            nw95_[int(7)] = d_467_v238_
            nw95_[int(8)] = (d_467_v238_) | (d_467_v238_)
            nw95_[int(9)] = (d_467_v238_) | (d_467_v238_)
            nw95_[int(10)] = (d_467_v238_ if d_186_v1_ else _dafny.Map({d_468_v240_: (d_469_v241_).cf133}))
            nw95_[int(11)] = d_467_v238_
            nw95_[int(12)] = d_467_v238_
            nw95_[int(13)] = d_467_v238_
            d_470_v242_ = nw95_
            index71_ = default__.safeIndex(519, (d_470_v242_).length(0))
            (d_470_v242_)[index71_] = d_467_v238_
            d_472_v244_: _dafny.Seq
            d_472_v244_ = _dafny.SeqWithoutIsStrInference([d_468_v240_])
            d_473_v245_: _dafny.Seq
            def iife39_():
                coll31_ = _dafny.Map()
                compr_31_: _dafny.Map
                for compr_31_ in (d_472_v244_).Elements:
                    d_474_v243_: _dafny.Map = compr_31_
                    if (d_474_v243_) in (d_472_v244_):
                        coll31_[d_474_v243_] = d_466_v237_
                return _dafny.Map(coll31_)
            d_473_v245_ = _dafny.SeqWithoutIsStrInference([default__.fm84(d_195_globalState_), d_467_v238_, (d_467_v238_) | (iife39_()
            )])
            index72_ = default__.safeIndex(519, (d_470_v242_).length(0))
            (d_470_v242_)[index72_] = (d_473_v245_)[default__.safeIndex(536, len(d_473_v245_))]
            (d_195_globalState_).f21 = d_186_v1_
            d_475_v246_: _dafny.MultiSet
            d_475_v246_ = _dafny.MultiSet([d_190_v5_])
            if not(((not(d_186_v1_) if d_186_v1_ else not(default__.fm29(d_475_v246_, d_192_v7_, d_195_globalState_)))) == (d_186_v1_)):
                d_187_v2_ = d_187_v2_
                (d_195_globalState_).f24 = d_192_v7_
                (d_195_globalState_).f23 = d_462_v233_
                index73_ = default__.safeIndex(299, (d_187_v2_).length(0))
                (d_187_v2_)[index73_] = d_186_v1_
                d_476_v247_: _dafny.Set
                d_476_v247_ = _dafny.Set({d_185_v0_})
                index74_ = default__.safeIndex(299, (d_187_v2_).length(0))
                (d_187_v2_)[index74_] = (d_476_v247_).isdisjoint(d_476_v247_)
                d_477_v248_: _dafny.MultiSet
                d_477_v248_ = _dafny.MultiSet([(d_187_v2_)[default__.safeIndex(299, (d_187_v2_).length(0))]])
                index75_ = default__.safeIndex(299, (d_187_v2_).length(0))
                (d_187_v2_)[index75_] = not((_dafny.MultiSet((d_190_v5_) + (_dafny.SeqWithoutIsStrInference([(d_187_v2_)[default__.safeIndex(299, (d_187_v2_).length(0))], (d_187_v2_)[default__.safeIndex(299, (d_187_v2_).length(0))]])))).ispropersubset((d_477_v248_) - (d_477_v248_)))
            elif True:
                d_185_v0_ = d_185_v0_
                (d_195_globalState_).f24 = len(d_465_v236_)
                d_186_v1_ = default__.fm18((d_192_v7_) * (d_192_v7_), d_192_v7_, d_192_v7_, d_195_globalState_)
                d_478_v249_: _dafny.Seq
                d_478_v249_ = _dafny.SeqWithoutIsStrInference([d_185_v0_, d_185_v0_, d_185_v0_])
                d_479_v250_: D13
                d_479_v250_ = D13_DC35()
                d_480_v251_: _dafny.Map
                d_480_v251_ = _dafny.Map({d_479_v250_: d_186_v1_})
                d_481_v252_: _dafny.Map
                d_481_v252_ = _dafny.Map({len(_dafny.Map({d_186_v1_: d_186_v1_})): (d_371_v166_).set(d_192_v7_, default__.abs(len(d_468_v240_)))})
                d_482_v253_: C11
                nw96_ = C11()
                nw96_.ctor__(d_478_v249_, (len(d_480_v251_)) not in (d_481_v252_), d_186_v1_)
                d_482_v253_ = nw96_
                (d_195_globalState_).f1 = 334
            if (d_370_v165_).ispropersubset((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_192_v7_]))).intersection(d_370_v165_)):
                (d_195_globalState_).f22 = d_186_v1_
                (d_195_globalState_).f27 = d_186_v1_
                d_483_v254_: _dafny.Seq
                d_483_v254_ = _dafny.SeqWithoutIsStrInference([d_185_v0_, d_185_v0_])
                d_484_v255_: T0
                nw97_ = C11()
                nw97_.ctor__(d_483_v254_, d_186_v1_, not(d_186_v1_))
                d_484_v255_ = nw97_
                d_485_v256_: D8
                d_485_v256_ = D8_DC20(d_484_v255_)
                d_486_v257_: D9
                d_486_v257_ = D9_DC24(d_462_v233_, d_186_v1_, len(d_468_v240_), d_485_v256_, d_484_v255_.f28)
                d_487_v258_: _dafny.Array
                nw98_ = _dafny.Array(None, 26)
                nw98_[int(0)] = (d_486_v257_).cf29
                nw98_[int(1)] = _dafny.CodePoint('b')
                nw98_[int(2)] = d_462_v233_
                nw98_[int(3)] = d_462_v233_
                nw98_[int(4)] = default__.fm42(d_192_v7_, d_192_v7_, d_195_globalState_)
                nw98_[int(5)] = d_462_v233_
                nw98_[int(6)] = d_462_v233_
                nw98_[int(7)] = d_462_v233_
                nw98_[int(8)] = d_462_v233_
                nw98_[int(9)] = d_462_v233_
                nw98_[int(10)] = d_462_v233_
                nw98_[int(11)] = d_462_v233_
                nw98_[int(12)] = d_462_v233_
                nw98_[int(13)] = d_462_v233_
                nw98_[int(14)] = d_462_v233_
                nw98_[int(15)] = d_462_v233_
                nw98_[int(16)] = d_462_v233_
                nw98_[int(17)] = d_462_v233_
                nw98_[int(18)] = d_462_v233_
                nw98_[int(19)] = (d_462_v233_ if (d_484_v255_).f29 else d_462_v233_)
                nw98_[int(20)] = d_462_v233_
                nw98_[int(21)] = (d_465_v236_)[default__.safeIndex(d_192_v7_, len(d_465_v236_))]
                nw98_[int(22)] = d_462_v233_
                nw98_[int(23)] = d_462_v233_
                nw98_[int(24)] = d_462_v233_
                nw98_[int(25)] = d_462_v233_
                d_487_v258_ = nw98_
                index76_ = default__.safeIndex(764, (d_185_v0_).length(0))
                (d_185_v0_)[index76_] = default__.safeDivisionInt(d_192_v7_, d_192_v7_)
                d_488_v259_: _dafny.Map
                d_488_v259_ = _dafny.Map({d_466_v237_: default__.fm42(d_192_v7_, d_192_v7_, d_195_globalState_)})
                index77_ = default__.safeIndex(764, (d_185_v0_).length(0))
                rhs71_ = d_487_v258_
                rhs72_ = ((d_192_v7_) * (d_192_v7_)) * (d_192_v7_)
                rhs73_ = (default__.safeModuloInt(d_192_v7_, len(d_465_v236_))) - (d_192_v7_)
                rhs74_ = (len(d_193_v8_)) - ((d_352_v150_).cf96)
                rhs75_ = d_488_v259_
                lhs60_ = d_195_globalState_
                lhs61_ = d_195_globalState_
                lhs62_ = d_185_v0_
                lhs63_ = default__.safeIndex(764, (d_185_v0_).length(0))
                d_487_v258_ = rhs71_
                lhs60_.f24 = rhs72_
                lhs61_.f1 = rhs73_
                lhs62_[lhs63_] = rhs74_
                d_488_v259_ = rhs75_
                (d_195_globalState_).f27 = not (d_186_v1_) or (d_484_v255_.f28)
                d_489_v260_: _dafny.MultiSet
                d_489_v260_ = _dafny.MultiSet([False])
                d_490_v261_: T2
                nw99_ = C3()
                nw99_.ctor__((d_484_v255_).f29, (d_484_v255_).f29, d_484_v255_.f28, (d_185_v0_)[default__.safeIndex(764, (d_185_v0_).length(0))])
                d_490_v261_ = nw99_
                d_491_v262_: _dafny.Map
                d_491_v262_ = _dafny.Map({(d_489_v260_).ispropersubset(d_489_v260_): d_490_v261_})
                d_491_v262_ = ((d_491_v262_) | (d_491_v262_)) | (d_491_v262_)
            elif True:
                index78_ = default__.safeIndex(656, (d_185_v0_).length(0))
                (d_185_v0_)[index78_] = d_192_v7_
                index79_ = default__.safeIndex(656, (d_185_v0_).length(0))
                (d_185_v0_)[index79_] = d_192_v7_
                d_492_v263_: D20
                d_492_v263_ = D20_DC52()
                d_493_v264_: _dafny.Map
                d_493_v264_ = _dafny.Map({d_186_v1_: d_492_v263_})
                d_493_v264_ = d_493_v264_
                d_494_v265_: C3
                nw100_ = C3()
                nw100_.ctor__(d_186_v1_, d_186_v1_, d_186_v1_, ((d_185_v0_)[default__.safeIndex(656, (d_185_v0_).length(0))] if not(d_186_v1_) else (0) - (d_192_v7_)))
                d_494_v265_ = nw100_
                (d_195_globalState_).f17 = d_193_v8_
                d_193_v8_ = (d_193_v8_) + ((d_193_v8_) + (d_193_v8_))
        elif True:
            (d_195_globalState_).f13 = d_186_v1_
            d_495_v266_: _dafny.Seq
            d_495_v266_ = _dafny.SeqWithoutIsStrInference([(d_190_v5_).set(default__.safeIndex(d_192_v7_, len(d_190_v5_)), d_186_v1_)])
            d_496_v267_: _dafny.Map
            d_496_v267_ = _dafny.Map({default__.fm29(_dafny.MultiSet([default__.fm62(d_195_globalState_), (d_495_v266_)[default__.safeIndex(d_192_v7_, len(d_495_v266_))], d_190_v5_]), d_192_v7_, d_195_globalState_): d_191_v6_})
            d_191_v6_ = ((d_496_v267_)[(d_369_v164_).ispropersubset(d_369_v164_)] if ((d_369_v164_).ispropersubset(d_369_v164_)) in (d_496_v267_) else (d_191_v6_) - (d_191_v6_))
            (d_195_globalState_).f24 = d_192_v7_
            (d_195_globalState_).f21 = d_186_v1_
            d_497_v268_: C12
            nw101_ = C12()
            nw101_.ctor__()
            d_497_v268_ = nw101_
        d_498_v269_: C9
        nw102_ = C9()
        nw102_.ctor__(d_186_v1_, (d_190_v5_)[default__.safeIndex(((d_371_v166_)[d_192_v7_] if (d_192_v7_) in (d_371_v166_) else d_192_v7_), len(d_190_v5_))])
        d_498_v269_ = nw102_
        d_499_v270_: _dafny.Seq
        d_499_v270_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
        d_500_v271_: D1
        d_500_v271_ = D1_DC2(63, d_192_v7_, d_499_v270_, d_186_v1_)
        d_501_v273_: D7
        def iife40_():
            coll32_ = _dafny.Map()
            compr_32_: int
            for compr_32_ in _dafny.IntegerRange(95, 914):
                d_502_v272_: int = compr_32_
                if ((95) <= (d_502_v272_)) and ((d_502_v272_) < (914)):
                    coll32_[(d_502_v272_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_503_i19_ in range(default__.abs(785))])))] = d_192_v7_
            return _dafny.Map(coll32_)
        d_501_v273_ = D7_DC16((_dafny.MultiSet([d_500_v271_, d_500_v271_, d_500_v271_, d_500_v271_, default__.fm85(d_186_v1_, iife40_()
, d_195_globalState_)])).intersection(_dafny.MultiSet([d_500_v271_, d_500_v271_, D1_DC2(d_192_v7_, d_192_v7_, d_499_v270_, d_186_v1_), d_500_v271_, d_500_v271_])))
        source7_ = d_501_v273_
        if source7_.is_DC17:
            d_504___mcc_h10_ = source7_.cf22
            d_505_cf22_ = d_504___mcc_h10_
            d_506_v274_: _dafny.Seq
            d_507_v275_: _dafny.Array
            out16_: _dafny.Seq
            out17_: _dafny.Array
            out16_, out17_ = (d_368_v163_).m4(d_195_globalState_)
            d_506_v274_ = out16_
            d_507_v275_ = out17_
            d_508_v276_: str
            d_508_v276_ = _dafny.CodePoint('x')
            d_509_v277_: _dafny.Seq
            d_509_v277_ = _dafny.SeqWithoutIsStrInference([d_499_v270_, d_499_v270_, d_499_v270_, d_499_v270_, _dafny.SeqWithoutIsStrInference([d_508_v276_, d_508_v276_, d_508_v276_])])
            if (d_509_v277_) < (_dafny.SeqWithoutIsStrInference([d_499_v270_ for d_510_i20_ in range(default__.abs(515))])):
                (d_195_globalState_).f24 = (693) - (d_192_v7_)
                (d_195_globalState_).f21 = d_186_v1_
                d_499_v270_ = d_499_v270_
                d_511_v278_: _dafny.Array
                nw103_ = _dafny.Array(D1.default()(), 16)
                d_511_v278_ = nw103_
                rhs76_ = d_511_v278_
                rhs77_ = (0) - (d_192_v7_)
                rhs78_ = 201
                rhs79_ = ((d_296_v97_)[257] if (257) in (d_296_v97_) else False)
                lhs64_ = d_195_globalState_
                lhs65_ = d_195_globalState_
                lhs66_ = d_195_globalState_
                d_511_v278_ = rhs76_
                lhs64_.f1 = rhs77_
                lhs65_.f24 = rhs78_
                lhs66_.f13 = rhs79_
                d_192_v7_ = d_192_v7_
            elif True:
                d_512_v279_: D29
                d_512_v279_ = D29_DC79(d_192_v7_, d_505_cf22_, d_505_cf22_, d_187_v2_, d_192_v7_)
                d_513_v280_: _dafny.Map
                d_513_v280_ = _dafny.Map({d_505_cf22_: (d_512_v279_).cf123})
                index80_ = default__.safeIndex(213, (d_507_v275_).length(0))
                (d_507_v275_)[index80_] = len(d_513_v280_)
                index81_ = default__.safeIndex(213, (d_507_v275_).length(0))
                (d_507_v275_)[index81_] = (((0) - (d_192_v7_)) - (d_192_v7_)) + (default__.safeDivisionInt(d_192_v7_, (0) - (len((d_499_v270_).set(default__.safeIndex(len(d_353_v151_), len(d_499_v270_)), d_508_v276_)))))
                (d_195_globalState_).f1 = (d_507_v275_)[default__.safeIndex(213, (d_507_v275_).length(0))]
                d_514_v281_: _dafny.Map
                d_514_v281_ = _dafny.Map({d_192_v7_: d_192_v7_})
                d_514_v281_ = (d_514_v281_).set(d_192_v7_, (0) - (d_192_v7_))
                d_515_v282_: _dafny.Array
                nw104_ = _dafny.Array(_dafny.Map({}), 25)
                d_515_v282_ = nw104_
                index82_ = default__.safeIndex(83, (d_515_v282_).length(0))
                (d_515_v282_)[index82_] = default__.fm86(d_186_v1_, (_dafny.MultiSet(d_190_v5_)).cardinality, _dafny.SeqWithoutIsStrInference([len(d_193_v8_), 109]), d_505_cf22_, d_195_globalState_)
                d_516_v284_: _dafny.MultiSet
                d_516_v284_ = _dafny.MultiSet([d_499_v270_, d_499_v270_])
                d_517_v285_: _dafny.Map
                d_517_v285_ = _dafny.Map({d_499_v270_: _dafny.Set({False})})
                index83_ = default__.safeIndex(83, (d_515_v282_).length(0))
                def iife41_():
                    coll33_ = _dafny.Map()
                    compr_33_: _dafny.Seq
                    for compr_33_ in (d_516_v284_).Elements:
                        d_518_v283_: _dafny.Seq = compr_33_
                        if (d_518_v283_) in (d_516_v284_):
                            coll33_[d_518_v283_] = d_191_v6_
                    return _dafny.Map(coll33_)
                (d_515_v282_)[index83_] = ((iife41_()
                ).set(d_499_v270_, d_191_v6_)) | (d_517_v285_)
                d_353_v151_ = d_353_v151_
            d_519_v286_: D17
            d_519_v286_ = D17_DC44(d_186_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pswum")))
            d_520_v287_: _dafny.Map
            d_520_v287_ = _dafny.Map({d_508_v276_: (0) - (d_192_v7_)})
            d_521_v288_: _dafny.Set
            d_521_v288_ = _dafny.Set({d_499_v270_, _dafny.SeqWithoutIsStrInference([d_508_v276_ for d_522_i21_ in range(default__.abs(703))])})
            d_523_v289_: _dafny.Map
            d_523_v289_ = _dafny.Map({(d_519_v286_).cf62: default__.fm18(d_192_v7_, len(d_520_v287_), len(d_521_v288_), d_195_globalState_)})
            default__.m0(d_523_v289_, d_192_v7_, d_505_cf22_, d_195_globalState_)
            index84_ = default__.safeIndex(901, (d_187_v2_).length(0))
            (d_187_v2_)[index84_] = (d_190_v5_)[default__.safeIndex(d_192_v7_, len(d_190_v5_))]
            index85_ = default__.safeIndex(901, (d_187_v2_).length(0))
            (d_187_v2_)[index85_] = d_186_v1_
        elif source7_.is_DC18:
            d_524___mcc_h11_ = source7_.cf23
            d_525_cf23_ = d_524___mcc_h11_
            rhs80_ = d_369_v164_
            rhs81_ = d_186_v1_
            lhs67_ = d_195_globalState_
            d_369_v164_ = rhs80_
            lhs67_.f22 = rhs81_
            d_526_v290_: _dafny.Seq
            d_526_v290_ = _dafny.SeqWithoutIsStrInference([d_185_v0_])
            d_527_v291_: _dafny.Array
            nw105_ = _dafny.Array(None, 14)
            nw105_[int(0)] = d_185_v0_
            nw105_[int(1)] = d_185_v0_
            nw105_[int(2)] = d_185_v0_
            nw105_[int(3)] = (d_526_v290_)[default__.safeIndex((0) - (d_192_v7_), len(d_526_v290_))]
            nw105_[int(4)] = d_185_v0_
            nw105_[int(5)] = d_185_v0_
            nw105_[int(6)] = d_185_v0_
            nw105_[int(7)] = d_185_v0_
            nw105_[int(8)] = d_185_v0_
            nw105_[int(9)] = (d_185_v0_ if d_186_v1_ else d_185_v0_)
            nw105_[int(10)] = d_185_v0_
            nw105_[int(11)] = d_185_v0_
            nw105_[int(12)] = d_185_v0_
            nw105_[int(13)] = d_185_v0_
            d_527_v291_ = nw105_
            index86_ = default__.safeIndex(227, (d_527_v291_).length(0))
            (d_527_v291_)[index86_] = d_185_v0_
            index87_ = default__.safeIndex(227, (d_527_v291_).length(0))
            (d_527_v291_)[index87_] = d_185_v0_
            d_528_v292_: str
            d_528_v292_ = _dafny.CodePoint('v')
            d_529_v293_: _dafny.MultiSet
            d_529_v293_ = _dafny.MultiSet([d_528_v292_, _dafny.CodePoint('d')])
            d_186_v1_ = ((d_529_v293_) - (d_529_v293_)).issubset(_dafny.MultiSet([d_528_v292_, d_528_v292_, d_528_v292_]))
            (d_195_globalState_).f1 = d_192_v7_
        elif source7_.is_DC19:
            d_530___mcc_h12_ = source7_.cf24
            d_531___mcc_h13_ = source7_.cf25
            d_532_cf25_ = d_531___mcc_h13_
            d_533_cf24_ = d_530___mcc_h12_
            if d_533_cf24_:
                d_192_v7_ = len(_dafny.Set({d_190_v5_}))
                d_534_v294_: str
                d_534_v294_ = _dafny.CodePoint('a')
                d_535_v295_: _dafny.Set
                d_535_v295_ = _dafny.Set({d_534_v294_})
                d_536_v297_: _dafny.Map
                def iife42_():
                    coll34_ = _dafny.Set()
                    compr_34_: str
                    for compr_34_ in (d_535_v295_).Elements:
                        d_537_v296_: str = compr_34_
                        if (d_537_v296_) in (d_535_v295_):
                            coll34_ = coll34_.union(_dafny.Set([d_537_v296_]))
                    return _dafny.Set(coll34_)
                d_536_v297_ = _dafny.Map({d_187_v2_: iife42_()
                })
                d_538_v298_: _dafny.Map
                d_538_v298_ = _dafny.Map({d_186_v1_: d_187_v2_})
                d_536_v297_ = (d_536_v297_).set((d_187_v2_ if d_186_v1_ else ((d_538_v298_)[d_532_cf25_] if (d_532_cf25_) in (d_538_v298_) else (D5_DC12(d_187_v2_)).cf17)), d_535_v295_)
                d_539_v299_: bool
                d_540_v300_: _dafny.Map
                out18_: bool
                out19_: _dafny.Map
                out18_, out19_ = (d_368_v163_).m3(d_369_v164_, d_192_v7_, d_195_globalState_)
                d_539_v299_ = out18_
                d_540_v300_ = out19_
                d_541_v301_: _dafny.Seq
                d_541_v301_ = _dafny.SeqWithoutIsStrInference([d_499_v270_, d_499_v270_, d_499_v270_])
                d_542_v302_: C6
                nw106_ = C6()
                nw106_.ctor__(d_189_v4_, d_541_v301_, d_186_v1_, False)
                d_542_v302_ = nw106_
                d_543_v303_: D30
                d_543_v303_ = D30_DC80(d_542_v302_)
                pat_let_tv3_ = d_542_v302_
                pat_let_tv4_ = d_542_v302_
                d_544_v304_: _dafny.Array
                nw107_ = _dafny.Array(None, 16)
                nw107_[int(0)] = d_543_v303_
                nw107_[int(1)] = d_543_v303_
                nw107_[int(2)] = (d_543_v303_ if d_532_cf25_ else d_543_v303_)
                nw107_[int(3)] = d_543_v303_
                nw107_[int(4)] = d_543_v303_
                nw107_[int(5)] = d_543_v303_
                nw107_[int(6)] = d_543_v303_
                nw107_[int(7)] = D30_DC80(d_542_v302_)
                def iife43_(_pat_let4_0):
                    def iife44_(d_545_dt__update__tmp_h2_):
                        def iife45_(_pat_let5_0):
                            def iife46_(d_546_dt__update_hcf125_h0_):
                                return D30_DC80(d_546_dt__update_hcf125_h0_)
                            return iife46_(_pat_let5_0)
                        return iife45_(pat_let_tv3_)
                    return iife44_(_pat_let4_0)
                nw107_[int(8)] = iife43_(d_543_v303_)
                nw107_[int(9)] = (d_543_v303_ if d_533_cf24_ else d_543_v303_)
                nw107_[int(10)] = D30_DC80(d_542_v302_)
                nw107_[int(11)] = D30_DC80(d_542_v302_)
                nw107_[int(12)] = D30_DC80(d_542_v302_)
                nw107_[int(13)] = d_543_v303_
                nw107_[int(14)] = d_543_v303_
                def iife47_(_pat_let6_0):
                    def iife48_(d_547_dt__update__tmp_h3_):
                        def iife49_(_pat_let7_0):
                            def iife50_(d_548_dt__update_hcf125_h1_):
                                return D30_DC80(d_548_dt__update_hcf125_h1_)
                            return iife50_(_pat_let7_0)
                        return iife49_(pat_let_tv4_)
                    return iife48_(_pat_let6_0)
                nw107_[int(15)] = iife47_(d_543_v303_)
                d_544_v304_ = nw107_
                index88_ = default__.safeIndex(289, (d_544_v304_).length(0))
                (d_544_v304_)[index88_] = d_543_v303_
                index89_ = default__.safeIndex(289, (d_544_v304_).length(0))
                (d_544_v304_)[index89_] = d_543_v303_
                d_549_v305_: C2
                nw108_ = C2()
                nw108_.ctor__()
                d_549_v305_ = nw108_
                d_550_v306_: _dafny.Set
                d_550_v306_ = _dafny.Set({d_549_v305_, d_549_v305_, d_549_v305_, d_549_v305_, d_549_v305_})
                rhs82_ = d_192_v7_
                rhs83_ = d_533_cf24_
                rhs84_ = d_550_v306_
                lhs68_ = d_195_globalState_
                lhs68_.f1 = rhs82_
                d_186_v1_ = rhs83_
                d_550_v306_ = rhs84_
            elif True:
                (d_195_globalState_).f24 = len(((d_499_v270_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_551_i22_ in range(default__.abs(853))]))) + (d_499_v270_))
                d_552_v307_: bool
                out20_: bool
                out20_ = (d_350_v148_).m5(d_192_v7_, d_195_globalState_)
                d_552_v307_ = out20_
                d_553_v308_: C0
                nw109_ = C0()
                nw109_.ctor__()
                d_553_v308_ = nw109_
                d_554_v309_: _dafny.Array
                nw110_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_554_v309_ = nw110_
                index90_ = default__.safeIndex(186, (d_554_v309_).length(0))
                (d_554_v309_)[index90_] = d_499_v270_
                index91_ = default__.safeIndex(186, (d_554_v309_).length(0))
                (d_554_v309_)[index91_] = d_499_v270_
                (d_195_globalState_).f1 = d_192_v7_
            d_555_v311_: _dafny.MultiSet
            d_555_v311_ = _dafny.MultiSet([d_187_v2_, d_187_v2_, d_187_v2_])
            d_556_v312_: _dafny.Map
            d_556_v312_ = _dafny.Map({d_499_v270_: (d_555_v311_).cardinality})
            d_557_v313_: _dafny.Map
            def iife51_():
                coll35_ = _dafny.Map()
                compr_35_: _dafny.Seq
                for compr_35_ in (d_556_v312_).keys.Elements:
                    d_558_v310_: _dafny.Seq = compr_35_
                    if (d_558_v310_) in (d_556_v312_):
                        coll35_[d_558_v310_] = (_dafny.MultiSet([d_186_v1_, d_533_cf24_])).cardinality
                return _dafny.Map(coll35_)
            d_557_v313_ = _dafny.Map({d_499_v270_: len(iife51_()
            )})
            d_559_v314_: _dafny.Map
            d_559_v314_ = _dafny.Map({d_192_v7_: len(d_557_v313_)})
            rhs85_ = ((d_559_v314_)[d_192_v7_] if (d_192_v7_) in (d_559_v314_) else d_192_v7_)
            rhs86_ = d_187_v2_
            lhs69_ = d_195_globalState_
            lhs69_.f24 = rhs85_
            d_187_v2_ = rhs86_
            d_560_v315_: _dafny.Array
            d_560_v315_ = d_185_v0_
            d_561_v316_: C11
            nw111_ = C11()
            nw111_.ctor__(_dafny.SeqWithoutIsStrInference([(d_560_v315_)]), d_186_v1_, d_186_v1_)
            d_561_v316_ = nw111_
            (d_195_globalState_).f1 = len(_dafny.SeqWithoutIsStrInference([d_561_v316_, d_561_v316_, d_561_v316_, d_561_v316_, d_561_v316_]))
            d_499_v270_ = d_499_v270_
        elif True:
            d_562___mcc_h14_ = source7_.cf21
            d_563_cf21_ = d_562___mcc_h14_
            d_564_v317_: D29
            d_564_v317_ = D29_DC79(-397, d_186_v1_, d_186_v1_, d_187_v2_, d_192_v7_)
            d_565_v318_: str
            d_565_v318_ = _dafny.CodePoint('t')
            pat_let_tv5_ = d_187_v2_
            pat_let_tv6_ = d_186_v1_
            pat_let_tv7_ = d_192_v7_
            pat_let_tv8_ = d_192_v7_
            pat_let_tv9_ = d_565_v318_
            pat_let_tv10_ = d_192_v7_
            pat_let_tv11_ = d_195_globalState_
            pat_let_tv12_ = d_186_v1_
            def iife53_(_pat_let9_0):
                def iife54_(d_566_dt__update__tmp_h4_):
                    def iife55_(_pat_let10_0):
                        def iife56_(d_567_dt__update_hcf123_h0_):
                            def iife57_(_pat_let11_0):
                                def iife58_(d_568_dt__update_hcf122_h0_):
                                    def iife59_(_pat_let12_0):
                                        def iife60_(d_569_dt__update_hcf120_h0_):
                                            return D29_DC79(d_569_dt__update_hcf120_h0_, (d_566_dt__update__tmp_h4_).cf121, d_568_dt__update_hcf122_h0_, d_567_dt__update_hcf123_h0_, (d_566_dt__update__tmp_h4_).cf124)
                                        return iife60_(_pat_let12_0)
                                    return iife59_(pat_let_tv7_)
                                return iife58_(_pat_let11_0)
                            return iife57_(pat_let_tv6_)
                        return iife56_(_pat_let10_0)
                    return iife55_(pat_let_tv5_)
                return iife54_(_pat_let9_0)
            def iife52_(_pat_let8_0):
                def iife61_(d_570_dt__update__tmp_h5_):
                    def iife62_(_pat_let13_0):
                        def iife63_(d_571_dt__update_hcf124_h0_):
                            def iife64_(_pat_let14_0):
                                def iife65_(d_572_dt__update_hcf121_h0_):
                                    return D29_DC79((d_570_dt__update__tmp_h5_).cf120, d_572_dt__update_hcf121_h0_, (d_570_dt__update__tmp_h5_).cf122, (d_570_dt__update__tmp_h5_).cf123, d_571_dt__update_hcf124_h0_)
                                return iife65_(_pat_let14_0)
                            return iife64_(pat_let_tv12_)
                        return iife63_(_pat_let13_0)
                    return iife62_(default__.fm27(pat_let_tv8_, pat_let_tv9_, pat_let_tv10_, pat_let_tv11_))
                return iife61_(_pat_let8_0)
            d_564_v317_ = iife52_(iife53_(d_564_v317_))
            d_573_v319_: C5
            nw112_ = C5()
            nw112_.ctor__(d_186_v1_, d_565_v318_, d_186_v1_, True)
            d_573_v319_ = nw112_
            d_574_v320_: _dafny.Map
            d_574_v320_ = _dafny.Map({d_186_v1_: d_185_v0_})
            d_575_v321_: _dafny.Array
            nw113_ = _dafny.Array(None, 23)
            nw113_[int(0)] = d_185_v0_
            nw113_[int(1)] = (d_185_v0_ if d_573_v319_.f35 else d_185_v0_)
            nw113_[int(2)] = d_185_v0_
            nw113_[int(3)] = d_185_v0_
            nw113_[int(4)] = d_185_v0_
            nw113_[int(5)] = d_185_v0_
            nw113_[int(6)] = d_185_v0_
            nw113_[int(7)] = d_185_v0_
            nw113_[int(8)] = d_185_v0_
            nw113_[int(9)] = d_185_v0_
            nw113_[int(10)] = (d_185_v0_ if d_573_v319_.f35 else d_185_v0_)
            nw113_[int(11)] = d_185_v0_
            nw113_[int(12)] = d_185_v0_
            nw113_[int(13)] = d_185_v0_
            nw113_[int(14)] = d_185_v0_
            nw113_[int(15)] = d_185_v0_
            nw113_[int(16)] = d_185_v0_
            nw113_[int(17)] = d_185_v0_
            nw113_[int(18)] = d_185_v0_
            nw113_[int(19)] = d_185_v0_
            nw113_[int(20)] = ((d_574_v320_)[not(False)] if (not(False)) in (d_574_v320_) else d_185_v0_)
            nw113_[int(21)] = d_185_v0_
            nw113_[int(22)] = d_185_v0_
            d_575_v321_ = nw113_
            nw114_ = _dafny.Array(None, 7)
            nw114_[int(0)] = d_185_v0_
            nw114_[int(1)] = d_185_v0_
            nw114_[int(2)] = d_185_v0_
            nw114_[int(3)] = d_185_v0_
            nw114_[int(4)] = d_185_v0_
            nw114_[int(5)] = d_185_v0_
            nw114_[int(6)] = d_185_v0_
            d_575_v321_ = nw114_
            d_576_v322_: D2
            d_576_v322_ = D2_DC4(d_573_v319_.f36)
            source8_ = d_576_v322_
            if source8_.is_DC5:
                d_577___mcc_h15_ = source8_.cf8
                d_578_cf8_ = d_577___mcc_h15_
                d_579_v323_: T1
                nw115_ = C2()
                nw115_.ctor__()
                d_579_v323_ = nw115_
                d_580_v324_: D25
                d_580_v324_ = D25_DC68(d_579_v323_)
                pat_let_tv13_ = d_579_v323_
                def iife66_(_pat_let15_0):
                    def iife67_(d_581_dt__update__tmp_h6_):
                        def iife68_(_pat_let16_0):
                            def iife69_(d_582_dt__update_hcf105_h0_):
                                return D25_DC68(d_582_dt__update_hcf105_h0_)
                            return iife69_(_pat_let16_0)
                        return iife68_(pat_let_tv13_)
                    return iife67_(_pat_let15_0)
                d_580_v324_ = iife66_(d_580_v324_)
                (d_195_globalState_).f1 = default__.safeModuloInt((d_192_v7_ if d_573_v319_.f35 else len(d_499_v270_)), (d_192_v7_) - (291))
                d_193_v8_ = (d_193_v8_) + (default__.fm19((0) - (d_192_v7_), d_578_cf8_, d_573_v319_.f35, d_186_v1_, d_195_globalState_))
                (d_579_v323_).m13((d_573_v319_.f35 if default__.fm26(len(d_369_v164_), d_195_globalState_) else d_573_v319_.f35), d_369_v164_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwgq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_583_i23_ in range(default__.abs(570))])), not (not(d_578_cf8_)) or (d_578_cf8_), d_195_globalState_)
            elif source8_.is_DC4:
                d_584___mcc_h16_ = source8_.cf7
                d_585_cf7_ = d_584___mcc_h16_
                d_186_v1_ = d_573_v319_.f35
                (d_195_globalState_).f1 = d_192_v7_
                (d_195_globalState_).f21 = (d_573_v319_.f35 if d_573_v319_.f35 else d_573_v319_.f35)
                d_586_v325_: _dafny.Seq
                d_586_v325_ = _dafny.SeqWithoutIsStrInference([d_185_v0_])
                d_587_v326_: T0
                nw116_ = C11()
                nw116_.ctor__(d_586_v325_, d_186_v1_, d_573_v319_.f35)
                d_587_v326_ = nw116_
                d_588_v327_: D8
                d_588_v327_ = D8_DC20(d_587_v326_)
                d_589_v328_: D9
                d_589_v328_ = D9_DC24(d_573_v319_.f36, d_186_v1_, d_192_v7_, d_588_v327_, d_587_v326_.f28)
                (d_573_v319_).f35 = (default__.fm25((d_589_v328_).cf31, (0) - (d_192_v7_), d_195_globalState_)) > (315)
            elif True:
                d_590___mcc_h17_ = source8_.cf9
                d_591_cf9_ = d_590___mcc_h17_
                d_185_v0_ = (d_185_v0_ if (d_573_v319_.f35 if d_573_v319_.f35 else d_573_v319_.f35) else d_185_v0_)
                index92_ = default__.safeIndex(690, (d_185_v0_).length(0))
                (d_185_v0_)[index92_] = d_192_v7_
                index93_ = default__.safeIndex(690, (d_185_v0_).length(0))
                rhs87_ = (d_192_v7_ if (d_190_v5_) != ((_dafny.SeqWithoutIsStrInference([d_186_v1_, d_186_v1_])).set(default__.safeIndex(d_192_v7_, len(_dafny.SeqWithoutIsStrInference([d_186_v1_, d_186_v1_]))), d_573_v319_.f35)) else 493)
                rhs88_ = 621
                rhs89_ = (((d_370_v165_)[d_192_v7_] if (d_192_v7_) in (d_370_v165_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdhhsof")))) if d_573_v319_.f35 else d_192_v7_)
                lhs70_ = d_195_globalState_
                lhs71_ = d_195_globalState_
                lhs72_ = d_185_v0_
                lhs73_ = default__.safeIndex(690, (d_185_v0_).length(0))
                lhs70_.f1 = rhs87_
                lhs71_.f1 = rhs88_
                lhs72_[lhs73_] = rhs89_
                index94_ = default__.safeIndex(690, (d_185_v0_).length(0))
                rhs90_ = d_192_v7_
                rhs91_ = (d_185_v0_)[default__.safeIndex(690, (d_185_v0_).length(0))]
                rhs92_ = d_573_v319_.f35
                lhs74_ = d_185_v0_
                lhs75_ = default__.safeIndex(690, (d_185_v0_).length(0))
                lhs76_ = d_573_v319_
                lhs74_[lhs75_] = rhs90_
                d_192_v7_ = rhs91_
                lhs76_.f35 = rhs92_
                d_592_v329_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_592_v329_ = nw117_
                d_593_v330_: _dafny.Array
                d_593_v330_ = d_592_v329_
                d_593_v330_ = d_592_v329_
        d_192_v7_ = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sprt"))) + (default__.fm51(d_192_v7_, d_192_v7_, d_195_globalState_))) + (d_499_v270_))
        d_594_v331_: D15
        d_594_v331_ = D15_DC38(d_186_v1_, (d_192_v7_) - (547), len(_dafny.SeqWithoutIsStrInference([d_186_v1_, d_186_v1_, d_186_v1_])))
        d_595_v332_: str
        d_595_v332_ = _dafny.CodePoint('f')
        d_596_v333_: D2
        d_596_v333_ = D2_DC4(d_595_v332_)
        d_597_v334_: D2
        d_597_v334_ = D2_DC6(d_596_v333_)
        pat_let_tv14_ = d_594_v331_
        pat_let_tv15_ = d_353_v151_
        pat_let_tv16_ = d_594_v331_
        pat_let_tv17_ = d_594_v331_
        def lambda9_(source9_):
            if source9_.is_DC5:
                d_598___mcc_h18_ = source9_.cf8
                d_599_cf8_ = d_598___mcc_h18_
                d_600_dt__update__tmp_h7_ = pat_let_tv14_
                d_601_dt__update_hcf54_h0_ = len(pat_let_tv15_)
                d_602_dt__update_hcf53_h0_ = False
                return D15_DC38(d_602_dt__update_hcf53_h0_, d_601_dt__update_hcf54_h0_, (d_600_dt__update__tmp_h7_).cf55)
            elif source9_.is_DC4:
                d_603___mcc_h19_ = source9_.cf7
                d_604_cf7_ = d_603___mcc_h19_
                return pat_let_tv16_
            elif True:
                d_605___mcc_h20_ = source9_.cf9
                d_606_cf9_ = d_605___mcc_h20_
                return pat_let_tv17_

        d_594_v331_ = lambda9_(d_597_v334_)
        d_607_v335_: _dafny.Map
        d_607_v335_ = _dafny.Map({d_186_v1_: d_186_v1_})
        d_608_v336_: C1
        nw118_ = C1()
        nw118_.ctor__(d_186_v1_, d_186_v1_)
        d_608_v336_ = nw118_
        d_609_v337_: _dafny.Seq
        d_609_v337_ = _dafny.SeqWithoutIsStrInference([d_608_v336_, d_608_v336_])
        d_610_v338_: D17
        d_610_v338_ = D17_DC45(d_607_v335_, not((d_192_v7_) < (len(d_609_v337_))))
        source10_ = d_610_v338_
        if source10_.is_DC44:
            d_611___mcc_h21_ = source10_.cf62
            d_612___mcc_h22_ = source10_.cf63
            d_613_cf63_ = d_612___mcc_h22_
            d_614_cf62_ = d_611___mcc_h21_
            (d_195_globalState_).f1 = ((d_192_v7_) + (d_192_v7_)) * (d_192_v7_)
            index95_ = default__.safeIndex(725, (d_185_v0_).length(0))
            (d_185_v0_)[index95_] = len(d_193_v8_)
            d_615_v339_: _dafny.Map
            d_615_v339_ = _dafny.Map({d_192_v7_: d_499_v270_})
            index96_ = default__.safeIndex(725, (d_185_v0_).length(0))
            (d_185_v0_)[index96_] = (0) - (len((d_615_v339_) | (_dafny.Map({d_192_v7_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdacmygc"))}))))
            d_616_v340_: C12
            nw119_ = C12()
            nw119_.ctor__()
            d_616_v340_ = nw119_
            d_617_v341_: T2
            nw120_ = C3()
            nw120_.ctor__(d_614_cf62_, d_186_v1_, True, len(d_369_v164_))
            d_617_v341_ = nw120_
            d_618_v342_: _dafny.Map
            d_618_v342_ = _dafny.Map({d_617_v341_: (D18_DC48(d_186_v1_, d_192_v7_, d_189_v4_, d_186_v1_, d_595_v332_)).cf68})
            (d_195_globalState_).f22 = (d_618_v342_) != (d_618_v342_)
        elif source10_.is_DC45:
            d_619___mcc_h23_ = source10_.cf64
            d_620___mcc_h24_ = source10_.cf65
            d_621_cf65_ = d_620___mcc_h24_
            d_622_cf64_ = d_619___mcc_h23_
            d_623_v343_: _dafny.Map
            d_623_v343_ = _dafny.Map({509: d_192_v7_})
            d_353_v151_ = (d_353_v151_).set((d_192_v7_) > (len(_dafny.SeqWithoutIsStrInference([len(d_623_v343_)]))), 512)
            index97_ = default__.safeIndex(722, (d_185_v0_).length(0))
            (d_185_v0_)[index97_] = (d_192_v7_) * (-5)
            index98_ = default__.safeIndex(91, (d_187_v2_).length(0))
            (d_187_v2_)[index98_] = d_186_v1_
            d_624_v344_: _dafny.Seq
            d_624_v344_ = _dafny.SeqWithoutIsStrInference([(d_499_v270_).set(default__.safeIndex(d_192_v7_, len(d_499_v270_)), d_595_v332_)])
            index99_ = default__.safeIndex(722, (d_185_v0_).length(0))
            index100_ = default__.safeIndex(91, (d_187_v2_).length(0))
            rhs93_ = d_192_v7_
            rhs94_ = d_192_v7_
            rhs95_ = (0) - (((0) - (d_192_v7_) if (d_498_v269_).fm11(True, d_624_v344_, d_192_v7_, d_195_globalState_) else default__.safeDivisionInt(d_192_v7_, d_192_v7_)))
            rhs96_ = (d_191_v6_) != (d_191_v6_)
            lhs77_ = d_195_globalState_
            lhs78_ = d_195_globalState_
            lhs79_ = d_185_v0_
            lhs80_ = default__.safeIndex(722, (d_185_v0_).length(0))
            lhs81_ = d_187_v2_
            lhs82_ = default__.safeIndex(91, (d_187_v2_).length(0))
            lhs77_.f24 = rhs93_
            lhs78_.f24 = rhs94_
            lhs79_[lhs80_] = rhs95_
            lhs81_[lhs82_] = rhs96_
            index101_ = default__.safeIndex(722, (d_185_v0_).length(0))
            (d_185_v0_)[index101_] = default__.fm8(d_195_globalState_)
            (d_195_globalState_).f1 = ((d_193_v8_)[default__.safeIndex((0) - ((0) - (len(d_499_v270_))), len(d_193_v8_))]) + (d_192_v7_)
        elif source10_.is_DC43:
            d_625___mcc_h25_ = source10_.cf61
            d_626_cf61_ = d_625___mcc_h25_
            d_627_v346_: _dafny.Seq
            d_627_v346_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qh"))])
            d_628_v347_: D17
            def iife70_():
                coll36_ = _dafny.Map()
                compr_36_: _dafny.Seq
                for compr_36_ in (d_627_v346_).Elements:
                    d_629_v345_: _dafny.Seq = compr_36_
                    if (d_629_v345_) in (d_627_v346_):
                        coll36_[d_629_v345_] = d_186_v1_
                return _dafny.Map(coll36_)
            d_628_v347_ = D17_DC43(iife70_()
)
            d_628_v347_ = d_628_v347_
            d_630_v348_: _dafny.MultiSet
            d_630_v348_ = _dafny.MultiSet([d_498_v269_])
            d_631_v349_: _dafny.Map
            d_631_v349_ = _dafny.Map({d_186_v1_: d_630_v348_})
            d_632_v350_: T2
            nw121_ = C3()
            nw121_.ctor__(not(d_186_v1_), d_186_v1_, d_186_v1_, (0) - ((0) - ((((d_631_v349_)[d_186_v1_] if (d_186_v1_) in (d_631_v349_) else d_630_v348_)).cardinality)))
            d_632_v350_ = nw121_
            d_633_v351_: _dafny.Seq
            d_633_v351_ = _dafny.SeqWithoutIsStrInference([d_632_v350_, d_632_v350_])
            d_634_v352_: D32
            d_634_v352_ = D32_DC87(d_632_v350_)
            d_635_v353_: _dafny.Map
            d_635_v353_ = _dafny.Map({d_186_v1_: d_632_v350_})
            d_636_v354_: _dafny.Array
            nw122_ = _dafny.Array(None, 28)
            nw122_[int(0)] = d_632_v350_
            nw122_[int(1)] = d_632_v350_
            nw122_[int(2)] = d_632_v350_
            nw122_[int(3)] = d_632_v350_
            nw122_[int(4)] = d_632_v350_
            nw122_[int(5)] = d_632_v350_
            nw122_[int(6)] = d_632_v350_
            nw122_[int(7)] = (d_633_v351_)[default__.safeIndex(d_192_v7_, len(d_633_v351_))]
            nw122_[int(8)] = d_632_v350_
            nw122_[int(9)] = d_632_v350_
            nw122_[int(10)] = d_632_v350_
            nw122_[int(11)] = d_632_v350_
            nw122_[int(12)] = d_632_v350_
            nw122_[int(13)] = (d_634_v352_).cf138
            nw122_[int(14)] = d_632_v350_
            nw122_[int(15)] = d_632_v350_
            nw122_[int(16)] = d_632_v350_
            nw122_[int(17)] = d_632_v350_
            nw122_[int(18)] = d_632_v350_
            nw122_[int(19)] = d_632_v350_
            nw122_[int(20)] = d_632_v350_
            nw122_[int(21)] = d_632_v350_
            nw122_[int(22)] = d_632_v350_
            nw122_[int(23)] = d_632_v350_
            nw122_[int(24)] = d_632_v350_
            nw122_[int(25)] = ((d_635_v353_)[d_186_v1_] if (d_186_v1_) in (d_635_v353_) else d_632_v350_)
            nw122_[int(26)] = d_632_v350_
            nw122_[int(27)] = d_632_v350_
            d_636_v354_ = nw122_
            d_637_v355_: _dafny.Map
            d_637_v355_ = _dafny.Map({d_636_v354_: d_186_v1_})
            d_638_v356_: C4
            nw123_ = C4()
            nw123_.ctor__(d_186_v1_, (d_190_v5_)[default__.safeIndex((d_632_v350_).f39, len(d_190_v5_))], d_186_v1_, d_186_v1_)
            d_638_v356_ = nw123_
            d_637_v355_ = (d_637_v355_).set(d_636_v354_, (d_190_v5_)[default__.safeIndex(len(_dafny.Map({len(d_499_v270_): d_638_v356_})), len(d_190_v5_))])
            rhs97_ = (d_638_v356_).f37
            rhs98_ = (((d_499_v270_) + (d_499_v270_)) + ((d_499_v270_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "seiep"))))).set(default__.safeIndex((0) - (d_192_v7_), len(((d_499_v270_) + (d_499_v270_)) + ((d_499_v270_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "seiep")))))), d_595_v332_)
            lhs83_ = d_195_globalState_
            lhs83_.f13 = rhs97_
            d_499_v270_ = rhs98_
            (d_195_globalState_).f21 = ((d_632_v350_).f39) >= (len((_dafny.Map({(d_638_v356_).f38: d_192_v7_})) | (d_353_v151_)))
        elif True:
            d_639___mcc_h26_ = source10_.cf66
            d_640_cf66_ = d_639___mcc_h26_
            d_641_v357_: _dafny.Seq
            d_641_v357_ = _dafny.SeqWithoutIsStrInference([d_350_v148_])
            d_642_v358_: _dafny.Map
            d_642_v358_ = _dafny.Map({(0) - ((-181) * (d_192_v7_)): len(d_641_v357_)})
            (d_195_globalState_).f24 = ((d_642_v358_)[(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))).set(default__.safeIndex(d_192_v7_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))), d_595_v332_))) + (len(d_499_v270_))] if ((len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))).set(default__.safeIndex(d_192_v7_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))), d_595_v332_))) + (len(d_499_v270_))) in (d_642_v358_) else d_192_v7_)
            if (d_192_v7_) < (d_192_v7_):
                (d_195_globalState_).f21 = (d_190_v5_)[default__.safeIndex(((0) - ((0) - (d_192_v7_)) if d_186_v1_ else len((d_499_v270_).set(default__.safeIndex(d_192_v7_, len(d_499_v270_)), d_595_v332_))), len(d_190_v5_))]
                d_643_v359_: C3
                nw124_ = C3()
                nw124_.ctor__(not (default__.fm26(d_192_v7_, d_195_globalState_)) or (d_186_v1_), d_186_v1_, d_186_v1_, d_192_v7_)
                d_643_v359_ = nw124_
                (d_195_globalState_).f24 = 878
                d_644_v360_: bool
                out21_: bool
                out21_ = (d_498_v269_).m5(d_192_v7_, d_195_globalState_)
                d_644_v360_ = out21_
                d_645_v361_: D2
                d_645_v361_ = D2_DC5(d_644_v360_)
                d_646_v362_: _dafny.MultiSet
                d_646_v362_ = _dafny.MultiSet([(d_643_v359_).f40, not((d_643_v359_).f40), ((d_643_v359_).fm32(d_191_v6_, ((d_353_v151_)[d_644_v360_] if (d_644_v360_) in (d_353_v151_) else len(d_369_v164_)), (d_643_v359_).f40, d_195_globalState_)) and (default__.fm10(d_186_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_643_v359_).f40, d_186_v1_, (d_643_v359_).f40, d_644_v360_]) for d_647_i24_ in range(default__.abs(163))])), (d_643_v359_).f40, d_195_globalState_)), (d_645_v361_).cf8])
                d_648_v364_: D2
                d_648_v364_ = D2_DC4(d_595_v332_)
                d_649_v365_: D5
                d_649_v365_ = D5_DC13(d_648_v364_)
                d_650_v366_: _dafny.Map
                d_650_v366_ = _dafny.Map({d_649_v365_: d_644_v360_})
                d_651_v367_: _dafny.Map
                d_651_v367_ = _dafny.Map({D5_DC13(d_648_v364_): d_192_v7_})
                d_652_v368_: _dafny.Set
                def iife71_():
                    coll37_ = _dafny.Map()
                    compr_37_: D5
                    for compr_37_ in (d_650_v366_).keys.Elements:
                        d_653_v363_: D5 = compr_37_
                        if (d_653_v363_) in (d_650_v366_):
                            coll37_[d_653_v363_] = len(d_190_v5_)
                    return _dafny.Map(coll37_)
                d_652_v368_ = _dafny.Set({iife71_()
                , d_651_v367_})
                rhs99_ = (246) != (d_192_v7_)
                rhs100_ = _dafny.Map({default__.fm13(d_195_globalState_): d_186_v1_})
                rhs101_ = default__.fm57(d_595_v332_, 39, d_192_v7_, d_195_globalState_)
                rhs102_ = (0) - ((0) - ((0) - (default__.safeModuloInt(d_192_v7_, len(d_652_v368_)))))
                lhs84_ = d_195_globalState_
                lhs85_ = d_195_globalState_
                lhs84_.f27 = rhs99_
                d_607_v335_ = rhs100_
                d_646_v362_ = rhs101_
                lhs85_.f24 = rhs102_
            elif True:
                d_654_v369_: _dafny.Map
                d_654_v369_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_192_v7_])) + (_dafny.SeqWithoutIsStrInference([len(d_190_v5_) for d_655_i25_ in range(default__.abs(631))])): _dafny.CodePoint('p')})
                index102_ = default__.safeIndex(616, (d_185_v0_).length(0))
                (d_185_v0_)[index102_] = default__.fm25(764, len(d_499_v270_), d_195_globalState_)
                d_656_v370_: D22
                d_656_v370_ = D22_DC59((d_186_v1_ if default__.fm29(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_186_v1_, d_186_v1_])]), 285, d_195_globalState_) else not(d_186_v1_)), (d_499_v270_) + (d_499_v270_), (0) - (d_192_v7_))
                pat_let_tv18_ = d_296_v97_
                pat_let_tv19_ = d_192_v7_
                pat_let_tv20_ = d_195_globalState_
                pat_let_tv21_ = d_192_v7_
                pat_let_tv22_ = d_195_globalState_
                pat_let_tv23_ = d_296_v97_
                pat_let_tv24_ = d_186_v1_
                index103_ = default__.safeIndex(616, (d_185_v0_).length(0))
                def iife72_(_pat_let17_0):
                    def iife73_(d_657_dt__update__tmp_h8_):
                        def iife74_(_pat_let18_0):
                            def iife75_(d_658_dt__update_hcf83_h0_):
                                return D22_DC59(d_658_dt__update_hcf83_h0_, (d_657_dt__update__tmp_h8_).cf84, (d_657_dt__update__tmp_h8_).cf85)
                            return iife75_(_pat_let18_0)
                        return iife74_(not(((pat_let_tv18_)[default__.fm12(pat_let_tv19_, not(False), pat_let_tv20_)] if (default__.fm12(pat_let_tv21_, not(False), pat_let_tv22_)) in (pat_let_tv23_) else pat_let_tv24_)))
                    return iife73_(_pat_let17_0)
                rhs103_ = d_654_v369_
                rhs104_ = d_192_v7_
                rhs105_ = iife72_(d_656_v370_)
                lhs86_ = d_185_v0_
                lhs87_ = default__.safeIndex(616, (d_185_v0_).length(0))
                d_654_v369_ = rhs103_
                lhs86_[lhs87_] = rhs104_
                d_656_v370_ = rhs105_
                (d_195_globalState_).f1 = (d_185_v0_)[default__.safeIndex(616, (d_185_v0_).length(0))]
                d_659_v371_: bool
                out22_: bool
                out22_ = (d_350_v148_).m5(default__.fm8(d_195_globalState_), d_195_globalState_)
                d_659_v371_ = out22_
                d_660_v372_: _dafny.Array
                nw125_ = _dafny.Array(None, 14)
                d_660_v372_ = nw125_
                d_661_v373_: D21
                d_661_v373_ = D21_DC54(d_660_v372_)
                d_661_v373_ = d_661_v373_
                d_499_v270_ = ((_dafny.SeqWithoutIsStrInference([d_595_v332_ for d_662_i26_ in range(default__.abs(16))])).set(default__.safeIndex(d_192_v7_, len(_dafny.SeqWithoutIsStrInference([d_595_v332_ for d_663_i26_ in range(default__.abs(16))]))), d_595_v332_)) + ((_dafny.SeqWithoutIsStrInference([d_595_v332_ for d_664_i27_ in range(default__.abs(-590))])) + (d_499_v270_))
            (d_195_globalState_).f13 = not (d_186_v1_) or (d_186_v1_)
            default__.m0((d_607_v335_) | (_dafny.Map({d_186_v1_: d_186_v1_})), d_192_v7_, d_186_v1_, d_195_globalState_)
        _dafny.print(_dafny.string_of((d_185_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_186_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v2_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_188_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v5_) == (_dafny.SeqWithoutIsStrInference([True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_191_v6_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_192_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v8_) == (_dafny.SeqWithoutIsStrInference([432, -988, 432, 432, -988, 432, 432, -988, 432]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v9_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([432, -988, 432]): 432}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_195_globalState_).f0)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_195_globalState_).f0)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_195_globalState_).f0)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_195_globalState_).f6)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_195_globalState_).f11) == (_dafny.SeqWithoutIsStrInference([True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_195_globalState_).f14) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_.f17) == (_dafny.SeqWithoutIsStrInference([432, -988, 432]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_.f19) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([432, -988, 432]): 432}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_globalState_.f26) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_globalState_.f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v97_) == (_dafny.Map({432: True, 38: True, 39: True, 40: True, 41: True, 42: True, 43: True, 44: True, 45: True, 46: True, 47: True, 48: True, 49: True, 50: True, 51: True, 52: True, 53: True, 54: True, 55: True, 56: True, 57: True, 58: True, 59: True, 60: True, 61: True, 62: True, 63: True, 64: True, 65: True, 66: True, 67: True, 68: True, 69: True, 70: True, 71: True, 72: True, 73: True, 74: True, 75: True, 76: True, 77: True, 78: True, 79: True, 80: True, 81: True, 82: True, 83: True, 84: True, 85: True, 86: True, 87: True, 88: True, 89: True, 90: True, 91: True, 92: True, 93: True, 94: True, 95: True, 96: True, 97: True, 98: True, 99: True, 100: True, 101: True, 102: True, 103: True, 104: True, 105: True, 106: True, 107: True, 108: True, 109: True, 110: True, 111: True, 112: True, 113: True, 114: True, 115: True, 116: True, 117: True, 118: True, 119: True, 120: True, 121: True, 122: True, 123: True, 124: True, 125: True, 126: True, 127: True, 128: True, 129: True, 130: True, 131: True, 132: True, 133: True, 134: True, 135: True, 136: True, 137: True, 138: True, 139: True, 140: True, 141: True, 142: True, 143: True, 144: True, 145: True, 146: True, 147: True, 148: True, 149: True, 150: True, 151: True, 152: True, 153: True, 154: True, 155: True, 156: True, 157: True, 158: True, 159: True, 160: True, 161: True, 162: True, 163: True, 164: True, 165: True, 166: True, 167: True, 168: True, 169: True, 170: True, 171: True, 172: True, 173: True, 174: True, 175: True, 176: True, 177: True, 178: True, 179: True, 180: True, 181: True, 182: True, 183: True, 184: True, 185: True, 186: True, 187: True, 188: True, 189: True, 190: True, 191: True, 192: True, 193: True, 194: True, 195: True, 196: True, 197: True, 198: True, 199: True, 200: True, 201: True, 202: True, 203: True, 204: True, 205: True, 206: True, 207: True, 208: True, 209: True, 210: True, 211: True, 212: True, 213: True, 214: True, 215: True, 216: True, 217: True, 218: True, 219: True, 220: True, 221: True, 222: True, 223: True, 224: True, 225: True, 226: True, 227: True, 228: True, 229: True, 230: True, 231: True, 232: True, 233: True, 234: True, 235: True, 236: True, 237: True, 238: True, 239: True, 240: True, 241: True, 242: True, 243: True, 244: True, 245: True, 246: True, 247: True, 248: True, 249: True, 250: True, 251: True, 252: True, 253: True, 254: True, 255: True, 256: True, 257: True, 258: True, 259: True, 260: True, 261: True, 262: True, 263: True, 264: True, 265: True, 266: True, 267: True, 268: True, 269: True, 270: True, 271: True, 272: True, 273: True, 274: True, 275: True, 276: True, 277: True, 278: True, 279: True, 280: True, 281: True, 282: True, 283: True, 284: True, 285: True, 286: True, 287: True, 288: True, 289: True, 290: True, 291: True, 292: True, 293: True, 294: True, 295: True, 296: True, 297: True, 298: True, 299: True, 300: True, 301: True, 302: True, 303: True, 304: True, 305: True, 306: True, 307: True, 308: True, 309: True, 310: True, 311: True, 312: True, 313: True, 314: True, 315: True, 316: True, 317: True, 318: True, 319: True, 320: True, 321: True, 322: True, 323: True, 324: True, 325: True, 326: True, 327: True, 328: True, 329: True, 330: True, 331: True, 332: True, 333: True, 334: True, 335: True, 336: True, 337: True, 338: True, 339: True, 340: True, 341: True, 342: True, 343: True, 344: True, 345: True, 346: True, 347: True, 348: True, 349: True, 350: True, 351: True, 352: True, 353: True, 354: True, 355: True, 356: True, 357: True, 358: True, 359: True, 360: True, 361: True, 362: True, 363: True, 364: True, 365: True, 366: True, 367: True, 368: True, 369: True, 370: True, 371: True, 372: True, 373: True, 374: True, 375: True, 376: True, 377: True, 378: True, 379: True, 380: True, 381: True, 382: True, 383: True, 384: True, 385: True, 386: True, 387: True, 388: True, 389: True, 390: True, 391: True, 392: True, 393: True, 394: True, 395: True, 396: True, 397: True, 398: True, 399: True, 400: True, 401: True, 402: True, 403: True, 404: True, 405: True, 406: True, 407: True, 408: True, 409: True, 410: True, 411: True, 412: True, 413: True, 414: True, 415: True, 416: True, 417: True, 418: True, 419: True, 420: True, 421: True, 422: True, 423: True, 424: True, 425: True, 426: True, 427: True, 428: True, 429: True, 430: True, 431: True, 0: True, 1: True, 2: True, 3: True, 4: True, 5: True, 6: True, 7: True, 8: True, 9: True, 10: True, 11: True, 12: True, 13: True, 14: True, 15: True, 16: True, 17: True, 18: True, 19: True, 20: True, 21: True, 22: True, 23: True, 24: True, 25: True, 26: True, 27: True, 28: True, 29: True, 30: True, 31: True, 32: True, 33: True, 34: True, 35: True, 36: True, 37: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_298_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_351_v149_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v150_).cf93))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v150_).cf94))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v150_).cf95))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v150_).cf96))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v150_).cf97))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_353_v151_) == (_dafny.Map({True: 512}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_369_v164_) == (_dafny.Set({432}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_370_v165_) == (_dafny.MultiSet([1, 432, 432]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_371_v166_) == (_dafny.MultiSet([3, 432, 432, -432]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_372_v167_).cf12) == (_dafny.MultiSet([1, 432, 432]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_499_v270_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_500_v271_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_500_v271_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_500_v271_).cf4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_500_v271_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_501_v273_).cf21) == (_dafny.MultiSet([D1_DC2(63, 984, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), True), D1_DC2(63, 984, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), True), D1_DC2(63, 984, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), True), D1_DC2(63, 984, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), True)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_594_v331_).cf53))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_594_v331_).cf54))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_594_v331_).cf55))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_595_v332_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_596_v333_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_597_v334_).cf9).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_607_v335_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_609_v337_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_610_v338_).cf64) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_610_v338_).cf65))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
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
        return lambda: D1_DC2(int(0), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {self.cf4.VerbatimString(True)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
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
        return lambda: D2_DC5(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({self.cf11.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC10(D4, NamedTuple('DC10', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(D2.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC15(D6, NamedTuple('DC15', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf20 == __o.cf20
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
        return lambda: D7_DC17(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(_dafny.CodePoint('D'), False, int(0), D8.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC27(False, False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC27(D10, NamedTuple('DC27', [('cf40', Any), ('cf41', Any), ('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC30(D11, NamedTuple('DC30', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)

class D12_DC32(D12, NamedTuple('DC32', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC35()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC35(D13, NamedTuple('DC35', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC36(D14, NamedTuple('DC36', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC38(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D15_DC41)

class D15_DC38(D15, NamedTuple('DC38', [('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC39(D15, NamedTuple('DC39', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({self.cf56.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC41(D15, NamedTuple('DC41', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC41({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC41) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)

class D16_DC42(D16, NamedTuple('DC42', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC44(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC44(D17, NamedTuple('DC44', [('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf62)}, {self.cf63.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC45(D17, NamedTuple('DC45', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC43(D17, NamedTuple('DC43', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC48(False, int(0), _dafny.Array(None, 0), False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D18_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)

class D18_DC48(D18, NamedTuple('DC48', [('cf68', Any), ('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC48({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC48) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC47(D18, NamedTuple('DC47', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D19_DC49)

class D19_DC49(D19, NamedTuple('DC49', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC49({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC49) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC51(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D20_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)

class D20_DC51(D20, NamedTuple('DC51', [('cf75', Any), ('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC51({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC51) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC52(D20, NamedTuple('DC52', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC50(D20, NamedTuple('DC50', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC53(D20, NamedTuple('DC53', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC55(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)

class D21_DC55(D21, NamedTuple('DC55', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC56(D21, NamedTuple('DC56', [])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56)
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC57(D21, NamedTuple('DC57', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC59(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D22_DC60)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D22_DC61)

class D22_DC59(D22, NamedTuple('DC59', [('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf83)}, {self.cf84.VerbatimString(True)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC60(D22, NamedTuple('DC60', [('cf86', Any), ('cf87', Any), ('cf88', Any), ('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC60({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC60) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC58(D22, NamedTuple('DC58', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC61(D22, NamedTuple('DC61', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC61({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC61) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC63(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D23_DC63)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D23_DC64)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D23_DC62)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D23_DC65)

class D23_DC63(D23, NamedTuple('DC63', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC63({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC63) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC64(D23, NamedTuple('DC64', [('cf93', Any), ('cf94', Any), ('cf95', Any), ('cf96', Any), ('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC64({_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC64) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC62(D23, NamedTuple('DC62', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC62({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC62) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC65(D23, NamedTuple('DC65', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC65({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC65) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC67(False, D5.default()(), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D24_DC67)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D24_DC66)

class D24_DC67(D24, NamedTuple('DC67', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any), ('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC67({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC67) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC66(D24, NamedTuple('DC66', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC66({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC66) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC69(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D25_DC69)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D25_DC68)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D25_DC70)

class D25_DC69(D25, NamedTuple('DC69', [('cf106', Any), ('cf107', Any), ('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC69({_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC69) and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC68(D25, NamedTuple('DC68', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC68({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC68) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC70(D25, NamedTuple('DC70', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC70({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC70) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC72(_dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D26_DC72)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D26_DC73)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D26_DC71)

class D26_DC72(D26, NamedTuple('DC72', [('cf111', Any), ('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC72({_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC72) and self.cf111 == __o.cf111 and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC73(D26, NamedTuple('DC73', [('cf113', Any), ('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC73({_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC73) and self.cf113 == __o.cf113 and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC71(D26, NamedTuple('DC71', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC71({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC71) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC75(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D27_DC75)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D27_DC74)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D27_DC76)

class D27_DC75(D27, NamedTuple('DC75', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC75({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC75) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC74(D27, NamedTuple('DC74', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC74({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC74) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC76(D27, NamedTuple('DC76', [('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC76({_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC76) and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D28_DC77)

class D28_DC77(D28, NamedTuple('DC77', [('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC77({_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC77) and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC79(int(0), False, False, _dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D29_DC79)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D29_DC78)

class D29_DC79(D29, NamedTuple('DC79', [('cf120', Any), ('cf121', Any), ('cf122', Any), ('cf123', Any), ('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC79({_dafny.string_of(self.cf120)}, {_dafny.string_of(self.cf121)}, {_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC79) and self.cf120 == __o.cf120 and self.cf121 == __o.cf121 and self.cf122 == __o.cf122 and self.cf123 == __o.cf123 and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC78(D29, NamedTuple('DC78', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC78({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC78) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC81(False, False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D30_DC81)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D30_DC82)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D30_DC83)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D30_DC80)

class D30_DC81(D30, NamedTuple('DC81', [('cf126', Any), ('cf127', Any), ('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC81({_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC81) and self.cf126 == __o.cf126 and self.cf127 == __o.cf127 and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC82(D30, NamedTuple('DC82', [('cf129', Any), ('cf130', Any), ('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC82({_dafny.string_of(self.cf129)}, {_dafny.string_of(self.cf130)}, {_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC82) and self.cf129 == __o.cf129 and self.cf130 == __o.cf130 and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC83(D30, NamedTuple('DC83', [('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC83({_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC83) and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC80(D30, NamedTuple('DC80', [('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC80({_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC80) and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: D31_DC85(None, _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D31_DC85)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D31_DC84)
    @property
    def is_DC86(self) -> bool:
        return isinstance(self, D31_DC86)

class D31_DC85(D31, NamedTuple('DC85', [('cf134', Any), ('cf135', Any), ('cf136', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC85({_dafny.string_of(self.cf134)}, {_dafny.string_of(self.cf135)}, {_dafny.string_of(self.cf136)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC85) and self.cf134 == __o.cf134 and self.cf135 == __o.cf135 and self.cf136 == __o.cf136
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC84(D31, NamedTuple('DC84', [('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC84({_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC84) and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC86(D31, NamedTuple('DC86', [('cf137', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC86({_dafny.string_of(self.cf137)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC86) and self.cf137 == __o.cf137
    def __hash__(self) -> int:
        return super().__hash__()


class D32:
    @classmethod
    def default(cls, ):
        return lambda: D32_DC88()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC88(self) -> bool:
        return isinstance(self, D32_DC88)
    @property
    def is_DC87(self) -> bool:
        return isinstance(self, D32_DC87)

class D32_DC88(D32, NamedTuple('DC88', [])):
    def __dafnystr__(self) -> str:
        return f'D32.DC88'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC88)
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC87(D32, NamedTuple('DC87', [('cf138', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC87({_dafny.string_of(self.cf138)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC87) and self.cf138 == __o.cf138
    def __hash__(self) -> int:
        return super().__hash__()


class D33:
    @classmethod
    def default(cls, ):
        return lambda: D33_DC90(_dafny.Seq({}), _dafny.Set({}), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC90(self) -> bool:
        return isinstance(self, D33_DC90)
    @property
    def is_DC91(self) -> bool:
        return isinstance(self, D33_DC91)
    @property
    def is_DC89(self) -> bool:
        return isinstance(self, D33_DC89)

class D33_DC90(D33, NamedTuple('DC90', [('cf140', Any), ('cf141', Any), ('cf142', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC90({_dafny.string_of(self.cf140)}, {_dafny.string_of(self.cf141)}, {_dafny.string_of(self.cf142)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC90) and self.cf140 == __o.cf140 and self.cf141 == __o.cf141 and self.cf142 == __o.cf142
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC91(D33, NamedTuple('DC91', [('cf143', Any), ('cf144', Any), ('cf145', Any), ('cf146', Any), ('cf147', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC91({_dafny.string_of(self.cf143)}, {_dafny.string_of(self.cf144)}, {_dafny.string_of(self.cf145)}, {_dafny.string_of(self.cf146)}, {_dafny.string_of(self.cf147)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC91) and self.cf143 == __o.cf143 and self.cf144 == __o.cf144 and self.cf145 == __o.cf145 and self.cf146 == __o.cf146 and self.cf147 == __o.cf147
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC89(D33, NamedTuple('DC89', [('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC89({_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC89) and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()


class D34:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC92(self) -> bool:
        return isinstance(self, D34_DC92)

class D34_DC92(D34, NamedTuple('DC92', [('cf148', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC92({_dafny.string_of(self.cf148)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC92) and self.cf148 == __o.cf148
    def __hash__(self) -> int:
        return super().__hash__()


class D35:
    @classmethod
    def default(cls, ):
        return lambda: D35_DC94(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC94(self) -> bool:
        return isinstance(self, D35_DC94)
    @property
    def is_DC95(self) -> bool:
        return isinstance(self, D35_DC95)
    @property
    def is_DC96(self) -> bool:
        return isinstance(self, D35_DC96)
    @property
    def is_DC93(self) -> bool:
        return isinstance(self, D35_DC93)

class D35_DC94(D35, NamedTuple('DC94', [('cf150', Any), ('cf151', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC94({_dafny.string_of(self.cf150)}, {_dafny.string_of(self.cf151)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC94) and self.cf150 == __o.cf150 and self.cf151 == __o.cf151
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC95(D35, NamedTuple('DC95', [('cf152', Any), ('cf153', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC95({_dafny.string_of(self.cf152)}, {_dafny.string_of(self.cf153)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC95) and self.cf152 == __o.cf152 and self.cf153 == __o.cf153
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC96(D35, NamedTuple('DC96', [('cf154', Any), ('cf155', Any), ('cf156', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC96({_dafny.string_of(self.cf154)}, {_dafny.string_of(self.cf155)}, {_dafny.string_of(self.cf156)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC96) and self.cf154 == __o.cf154 and self.cf155 == __o.cf155 and self.cf156 == __o.cf156
    def __hash__(self) -> int:
        return super().__hash__()

class D35_DC93(D35, NamedTuple('DC93', [('cf149', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC93({_dafny.string_of(self.cf149)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC93) and self.cf149 == __o.cf149
    def __hash__(self) -> int:
        return super().__hash__()


class D36:
    @classmethod
    def default(cls, ):
        return lambda: D36_DC98(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC98(self) -> bool:
        return isinstance(self, D36_DC98)
    @property
    def is_DC97(self) -> bool:
        return isinstance(self, D36_DC97)

class D36_DC98(D36, NamedTuple('DC98', [('cf158', Any)])):
    def __dafnystr__(self) -> str:
        return f'D36.DC98({_dafny.string_of(self.cf158)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D36_DC98) and self.cf158 == __o.cf158
    def __hash__(self) -> int:
        return super().__hash__()

class D36_DC97(D36, NamedTuple('DC97', [('cf157', Any)])):
    def __dafnystr__(self) -> str:
        return f'D36.DC97({_dafny.string_of(self.cf157)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D36_DC97) and self.cf157 == __o.cf157
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
    def m5(self, p0, globalState):
        pass


class T1:
    pass
    def m13(self, p0, p1, p2, p3, globalState):
        pass


class T2:
    pass
    def m15(self, p0, globalState):
        pass

    def m16(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f13: bool = False
        self.f17: _dafny.Seq = _dafny.Seq({})
        self.f19: _dafny.Map = _dafny.Map({})
        self.f21: bool = False
        self.f22: bool = False
        self.f23: str = _dafny.CodePoint('D')
        self.f24: int = int(0)
        self.f26: _dafny.Set = _dafny.Set({})
        self.f27: bool = False
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f2: str = _dafny.CodePoint('D')
        self._f3: str = _dafny.CodePoint('D')
        self._f4: int = int(0)
        self._f5: bool = False
        self._f6: _dafny.Map = _dafny.Map({})
        self._f7: str = _dafny.CodePoint('D')
        self._f8: bool = False
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f10: int = int(0)
        self._f11: _dafny.Seq = _dafny.Seq({})
        self._f12: int = int(0)
        self._f14: _dafny.Set = _dafny.Set({})
        self._f15: int = int(0)
        self._f16: int = int(0)
        self._f18: int = int(0)
        self._f20: int = int(0)
        self._f25: bool = False
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
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self).f22 = f22
        (self).f23 = f23
        (self).f24 = f24
        (self)._f25 = f25
        (self).f26 = f26
        (self).f27 = f27

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
    def f12(self):
        return self._f12
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
    def f18(self):
        return self._f18
    @property
    def f20(self):
        return self._f20
    @property
    def f25(self):
        return self._f25

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm33(self, p0, p1, p2, globalState):
        return True

    def fm34(self, p0, p1, globalState):
        return True


class C1(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
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
    def ctor__(self, f28, f29):
        (self).f28 = f28
        (self)._f29 = f29

    def fm36(self, p0, p1, globalState):
        return (0) - ((207) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtbl"))]))))

    def m5(self, p0, globalState):
        r0: bool = False
        d_665_v0_: _dafny.Seq
        d_665_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaknucla"))
        d_666_v1_: D1
        d_666_v1_ = D1_DC2(default__.safeDivisionInt(p0, p0), p0, d_665_v0_, (self).f29)
        source11_ = d_666_v1_
        if source11_.is_DC2:
            d_667___mcc_h0_ = source11_.cf2
            d_668___mcc_h1_ = source11_.cf3
            d_669___mcc_h2_ = source11_.cf4
            d_670___mcc_h3_ = source11_.cf5
            d_671_cf5_ = d_670___mcc_h3_
            d_672_cf4_ = d_669___mcc_h2_
            d_673_cf3_ = d_668___mcc_h1_
            d_674_cf2_ = d_667___mcc_h0_
            d_675_v2_: _dafny.Set
            d_675_v2_ = _dafny.Set({self.f28, d_671_cf5_})
            (globalState).f26 = (d_675_v2_) | (d_675_v2_)
            d_676_v3_: _dafny.Map
            d_676_v3_ = _dafny.Map({False: d_673_cf3_})
            d_677_v4_: _dafny.Set
            d_677_v4_ = _dafny.Set({d_665_v0_})
            (globalState).f24 = (((d_676_v3_)[not((self).f29)] if (not((self).f29)) in (d_676_v3_) else (self).fm36(d_674_cf2_, d_677_v4_, globalState))) + (((0) - (d_674_cf2_)) - ((0) - (len(d_665_v0_))))
            d_678_v5_: _dafny.MultiSet
            d_678_v5_ = _dafny.MultiSet([True, self.f28, d_671_cf5_, self.f28, d_671_cf5_])
            d_679_v6_: str
            d_679_v6_ = _dafny.CodePoint('y')
            d_680_v7_: _dafny.Set
            d_680_v7_ = _dafny.Set({d_679_v6_})
            d_681_v9_: _dafny.MultiSet
            def iife76_():
                coll38_ = _dafny.Set()
                compr_38_: str
                for compr_38_ in (d_680_v7_).Elements:
                    d_682_v8_: str = compr_38_
                    if (d_682_v8_) in (d_680_v7_):
                        coll38_ = coll38_.union(_dafny.Set([d_682_v8_]))
                return _dafny.Set(coll38_)
            d_681_v9_ = _dafny.MultiSet([iife76_()
            ])
            d_683_v10_: _dafny.Array
            nw126_ = _dafny.Array(None, 10)
            nw126_[int(0)] = d_671_cf5_
            nw126_[int(1)] = (self).f29
            nw126_[int(2)] = d_671_cf5_
            nw126_[int(3)] = d_671_cf5_
            nw126_[int(4)] = d_671_cf5_
            nw126_[int(5)] = (_dafny.Set({len(d_672_cf4_)})).isdisjoint(default__.fm37((d_678_v5_).cardinality, d_671_cf5_, self.f28, globalState))
            nw126_[int(6)] = (not(default__.fm38(globalState))) == (d_671_cf5_)
            nw126_[int(7)] = self.f28
            nw126_[int(8)] = d_671_cf5_
            nw126_[int(9)] = (d_681_v9_).ispropersubset(d_681_v9_)
            d_683_v10_ = nw126_
            index104_ = default__.safeIndex(979, (d_683_v10_).length(0))
            (d_683_v10_)[index104_] = (self).f29
            d_684_v11_: _dafny.Array
            nw127_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_684_v11_ = nw127_
            d_685_v12_: _dafny.Map
            d_685_v12_ = _dafny.Map({p0: d_684_v11_})
            d_686_v13_: _dafny.Array
            nw128_ = _dafny.Array(None, 24)
            nw128_[int(0)] = (d_684_v11_ if d_671_cf5_ else d_684_v11_)
            nw128_[int(1)] = d_684_v11_
            nw128_[int(2)] = d_684_v11_
            nw128_[int(3)] = d_684_v11_
            nw128_[int(4)] = d_684_v11_
            nw128_[int(5)] = d_684_v11_
            nw128_[int(6)] = ((d_685_v12_)[d_674_cf2_] if (d_674_cf2_) in (d_685_v12_) else d_684_v11_)
            nw128_[int(7)] = d_684_v11_
            nw128_[int(8)] = d_684_v11_
            nw128_[int(9)] = d_684_v11_
            nw128_[int(10)] = d_684_v11_
            nw128_[int(11)] = d_684_v11_
            nw128_[int(12)] = d_684_v11_
            nw128_[int(13)] = d_684_v11_
            nw128_[int(14)] = d_684_v11_
            nw128_[int(15)] = d_684_v11_
            nw128_[int(16)] = d_684_v11_
            nw128_[int(17)] = d_684_v11_
            nw128_[int(18)] = d_684_v11_
            nw128_[int(19)] = d_684_v11_
            nw128_[int(20)] = d_684_v11_
            nw128_[int(21)] = d_684_v11_
            nw128_[int(22)] = d_684_v11_
            nw128_[int(23)] = d_684_v11_
            d_686_v13_ = nw128_
            index105_ = default__.safeIndex(925, (d_686_v13_).length(0))
            (d_686_v13_)[index105_] = d_684_v11_
            d_687_v14_: _dafny.Seq
            d_687_v14_ = _dafny.SeqWithoutIsStrInference([d_684_v11_, d_684_v11_, d_684_v11_])
            index106_ = default__.safeIndex(979, (d_683_v10_).length(0))
            index107_ = default__.safeIndex(925, (d_686_v13_).length(0))
            rhs106_ = (self).f29
            rhs107_ = (d_687_v14_)[default__.safeIndex((-473 if (self).f29 else d_673_cf3_), len(d_687_v14_))]
            lhs88_ = d_683_v10_
            lhs89_ = default__.safeIndex(979, (d_683_v10_).length(0))
            lhs90_ = d_686_v13_
            lhs91_ = default__.safeIndex(925, (d_686_v13_).length(0))
            lhs88_[lhs89_] = rhs106_
            lhs90_[lhs91_] = rhs107_
            d_674_cf2_ = d_674_cf2_
        elif source11_.is_DC1:
            d_688___mcc_h4_ = source11_.cf1
            d_689_cf1_ = d_688___mcc_h4_
            d_690_v15_: C0
            nw129_ = C0()
            nw129_.ctor__()
            d_690_v15_ = nw129_
            (globalState).f24 = (0) - (p0)
            d_691_v16_: _dafny.Seq
            d_691_v16_ = _dafny.SeqWithoutIsStrInference([True, (self).f29, self.f28, True])
            d_691_v16_ = d_691_v16_
            d_692_v17_: _dafny.Array
            def lambda10_(d_693_i0_):
                return self.f28

            init5_ = lambda10_
            nw130_ = _dafny.Array(None, 8)
            for i0_5_ in range(nw130_.length(0)):
                nw130_[i0_5_] = init5_(i0_5_)
            d_692_v17_ = nw130_
            d_694_v18_: _dafny.Array
            nw131_ = _dafny.Array(None, 6)
            nw131_[int(0)] = (0) - (p0)
            nw131_[int(1)] = p0
            nw131_[int(2)] = p0
            nw131_[int(3)] = 616
            nw131_[int(4)] = p0
            nw131_[int(5)] = 835
            d_694_v18_ = nw131_
            d_695_v19_: _dafny.Map
            d_695_v19_ = _dafny.Map({d_692_v17_: d_694_v18_})
            d_696_v20_: _dafny.Map
            d_696_v20_ = _dafny.Map({p0: p0})
            d_697_v21_: _dafny.MultiSet
            d_697_v21_ = _dafny.MultiSet([527])
            d_698_v22_: _dafny.Array
            nw132_ = _dafny.Array(None, 29)
            nw132_[int(0)] = len(d_695_v19_)
            nw132_[int(1)] = p0
            nw132_[int(2)] = p0
            nw132_[int(3)] = len(d_696_v20_)
            nw132_[int(4)] = p0
            nw132_[int(5)] = p0
            nw132_[int(6)] = len(d_665_v0_)
            nw132_[int(7)] = p0
            nw132_[int(8)] = (d_697_v21_).cardinality
            nw132_[int(9)] = p0
            nw132_[int(10)] = p0
            nw132_[int(11)] = p0
            nw132_[int(12)] = p0
            nw132_[int(13)] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f28])): len(d_691_v16_)}))
            nw132_[int(14)] = (0) - (p0)
            nw132_[int(15)] = p0
            nw132_[int(16)] = (0) - (len(d_691_v16_))
            nw132_[int(17)] = p0
            nw132_[int(18)] = p0
            nw132_[int(19)] = (0) - (p0)
            nw132_[int(20)] = p0
            nw132_[int(21)] = p0
            nw132_[int(22)] = p0
            nw132_[int(23)] = p0
            nw132_[int(24)] = p0
            nw132_[int(25)] = p0
            nw132_[int(26)] = p0
            nw132_[int(27)] = (self).fm36(p0, _dafny.Set({d_665_v0_}), globalState)
            nw132_[int(28)] = (_dafny.MultiSet(d_665_v0_)).cardinality
            d_698_v22_ = nw132_
            d_699_v23_: _dafny.Array
            d_699_v23_ = d_698_v22_
            source12_ = d_699_v23_
            d_700___mcc_h6_ = source12_
            d_701_cf0_ = d_700___mcc_h6_
            d_702_v24_: _dafny.Set
            d_702_v24_ = _dafny.Set({p0, p0, p0, 947, len(_dafny.Set({(self).f29}))})
            d_703_v27_: str
            d_703_v27_ = _dafny.CodePoint('t')
            d_704_v29_: _dafny.Array
            nw133_ = _dafny.Array(None, 8)
            nw133_[int(0)] = (d_702_v24_) | (_dafny.Set({p0}))
            def iife77_():
                coll39_ = _dafny.Set()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(-12, 940):
                    d_705_v25_: int = compr_39_
                    if ((-12) <= (d_705_v25_)) and ((d_705_v25_) < (940)):
                        coll39_ = coll39_.union(_dafny.Set([(d_705_v25_) * (p0)]))
                return _dafny.Set(coll39_)
            def iife78_():
                coll40_ = _dafny.Set()
                compr_40_: int
                for compr_40_ in _dafny.IntegerRange(878, 820):
                    d_706_v26_: int = compr_40_
                    if ((878) <= (d_706_v26_)) and ((d_706_v26_) < (820)):
                        coll40_ = coll40_.union(_dafny.Set([(d_706_v26_) * ((D4_DC10(False, (_dafny.MultiSet([(self).f29, self.f28])).cardinality, p0)).cf14)]))
                return _dafny.Set(coll40_)
            nw133_[int(1)] = (iife77_()
             if self.f28 else iife78_()
            )
            nw133_[int(2)] = (d_702_v24_) | (d_702_v24_)
            nw133_[int(3)] = default__.fm37((0) - (p0), (self).f29, self.f28, globalState)
            def iife79_():
                coll41_ = _dafny.Set()
                compr_41_: int
                for compr_41_ in (_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([d_703_v27_])).cardinality), p0])).Elements:
                    d_707_v28_: int = compr_41_
                    if (d_707_v28_) in (_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([d_703_v27_])).cardinality), p0])):
                        coll41_ = coll41_.union(_dafny.Set([default__.safeModuloInt(d_707_v28_, len(_dafny.SeqWithoutIsStrInference([331 for d_708_i1_ in range(default__.abs(407))])))]))
                return _dafny.Set(coll41_)
            nw133_[int(4)] = (d_702_v24_) | (iife79_()
            )
            nw133_[int(5)] = (d_702_v24_) - (d_702_v24_)
            nw133_[int(6)] = d_702_v24_
            nw133_[int(7)] = d_702_v24_
            d_704_v29_ = nw133_
            index108_ = default__.safeIndex(666, (d_704_v29_).length(0))
            (d_704_v29_)[index108_] = d_702_v24_
            d_709_v30_: _dafny.Map
            d_709_v30_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([True, self.f28])})
            d_710_v31_: D4
            d_710_v31_ = D4_DC9(d_697_v21_)
            d_711_v32_: _dafny.Map
            d_711_v32_ = _dafny.Map({(self).f29: p0})
            d_712_v33_: _dafny.Map
            d_712_v33_ = _dafny.Map({d_710_v31_: d_711_v32_})
            d_713_v34_: _dafny.Seq
            d_713_v34_ = _dafny.SeqWithoutIsStrInference([d_709_v30_, d_709_v30_])
            index109_ = default__.safeIndex(666, (d_704_v29_).length(0))
            rhs108_ = (self.f28) == (False)
            rhs109_ = d_703_v27_
            rhs110_ = _dafny.Set({p0, (p0) * (len(_dafny.SeqWithoutIsStrInference([d_703_v27_ for d_714_i2_ in range(default__.abs(827))]))), (len(d_665_v0_)) + (len(((d_712_v33_)[D4_DC9(d_697_v21_)] if (D4_DC9(d_697_v21_)) in (d_712_v33_) else _dafny.Map({self.f28: p0}))))})
            rhs111_ = (d_697_v21_) | (d_697_v21_)
            rhs112_ = ((d_709_v30_) | (d_709_v30_)) | ((d_713_v34_)[default__.safeIndex(p0, len(d_713_v34_))])
            lhs92_ = globalState
            lhs93_ = globalState
            lhs94_ = d_704_v29_
            lhs95_ = default__.safeIndex(666, (d_704_v29_).length(0))
            lhs92_.f21 = rhs108_
            lhs93_.f23 = rhs109_
            lhs94_[lhs95_] = rhs110_
            d_697_v21_ = rhs111_
            d_709_v30_ = rhs112_
            d_694_v18_ = d_701_cf0_
            (globalState).f1 = p0
            (globalState).f1 = default__.safeDivisionInt(len(d_665_v0_), (p0) + (p0))
        elif True:
            d_715___mcc_h5_ = source11_.cf6
            d_716_cf6_ = d_715___mcc_h5_
            d_717_v35_: C0
            nw134_ = C0()
            nw134_.ctor__()
            d_717_v35_ = nw134_
            d_718_v36_: _dafny.Array
            nw135_ = _dafny.Array(_dafny.Map({}), 24)
            d_718_v36_ = nw135_
            d_719_v37_: _dafny.Map
            d_719_v37_ = _dafny.Map({D6_DC14(d_718_v36_): -612})
            d_720_v38_: D6
            d_720_v38_ = D6_DC14(d_718_v36_)
            d_719_v37_ = _dafny.Map({d_720_v38_: p0})
            d_721_v39_: C0
            nw136_ = C0()
            nw136_.ctor__()
            d_721_v39_ = nw136_
            (globalState).f24 = p0
        d_722_v40_: _dafny.Array
        def lambda11_(d_723_i4_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omsof"))

        init6_ = lambda11_
        nw137_ = _dafny.Array(None, 15)
        for i0_6_ in range(nw137_.length(0)):
            nw137_[i0_6_] = init6_(i0_6_)
        d_722_v40_ = nw137_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_722_v40_).length(0)):
            d_724_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_724_i3_)) and ((d_724_i3_) < ((d_722_v40_).length(0)))):
                (d_722_v40_)[(d_724_i3_)] = (d_665_v0_) + (d_665_v0_)
        d_725_v41_: _dafny.Array
        nw138_ = _dafny.Array(int(0), 10)
        d_725_v41_ = nw138_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_725_v41_).length(0)):
            d_726_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_726_i5_)) and ((d_726_i5_) < ((d_725_v41_).length(0)))):
                (d_725_v41_)[(d_726_i5_)] = default__.safeDivisionInt(d_726_i5_, p0)
        if ((self).f29) == ((self).f29):
            d_727_v42_: D4
            d_727_v42_ = D4_DC11(default__.fm39((self).f29, globalState))
            source13_ = d_727_v42_
            if source13_.is_DC10:
                d_728___mcc_h7_ = source13_.cf13
                d_729___mcc_h8_ = source13_.cf14
                d_730___mcc_h9_ = source13_.cf15
                d_731_cf15_ = d_730___mcc_h9_
                d_732_cf14_ = d_729___mcc_h8_
                d_733_cf13_ = d_728___mcc_h7_
                (globalState).f1 = (0) - ((p0) - (((0) - (p0)) * ((0) - (len((D3_DC8(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_734_i6_ in range(default__.abs(416))]))).cf11)))))
                d_665_v0_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxrbrj"))) + (d_665_v0_)) + ((d_665_v0_) + (d_665_v0_))
                d_732_cf14_ = default__.safeModuloInt(d_731_cf15_, p0)
                d_735_v43_: _dafny.Map
                d_735_v43_ = _dafny.Map({d_733_cf13_: self.f28})
                d_735_v43_ = (d_735_v43_).set((self).f29, self.f28)
            elif source13_.is_DC9:
                d_736___mcc_h10_ = source13_.cf12
                d_737_cf12_ = d_736___mcc_h10_
                d_738_v44_: _dafny.Map
                d_738_v44_ = _dafny.Map({default__.fm40(globalState): 572})
                d_739_v45_: _dafny.Map
                d_739_v45_ = _dafny.Map({d_665_v0_: (0) - (p0)})
                d_740_v46_: _dafny.Seq
                d_740_v46_ = _dafny.SeqWithoutIsStrInference([len(d_665_v0_), p0])
                d_738_v44_ = (d_738_v44_).set((d_739_v45_) | (d_739_v45_), (d_740_v46_)[default__.safeIndex(504, len(d_740_v46_))])
                d_741_v47_: _dafny.Array
                nw139_ = _dafny.Array(False, 13)
                d_741_v47_ = nw139_
                d_741_v47_ = d_741_v47_
                d_742_v48_: D2
                d_742_v48_ = D2_DC5(self.f28)
                d_743_v49_: _dafny.Map
                d_743_v49_ = _dafny.Map({d_742_v48_: self.f28})
                d_744_v50_: _dafny.Map
                d_744_v50_ = _dafny.Map({(0) - (default__.safeDivisionInt(p0, p0)): d_743_v49_})
                d_744_v50_ = (_dafny.Map({p0: d_743_v49_})) | ((d_744_v50_).set(p0, d_743_v49_))
                (globalState).f1 = p0
            elif True:
                d_745___mcc_h11_ = source13_.cf16
                d_746_cf16_ = d_745___mcc_h11_
                d_747_v51_: _dafny.Map
                d_747_v51_ = _dafny.Map({(self).f29: (self).f29})
                d_748_v52_: str
                d_748_v52_ = _dafny.CodePoint('r')
                d_747_v51_ = (d_747_v51_).set((d_748_v52_) not in (d_665_v0_), (self).f29)
                (globalState).f24 = p0
                d_749_v53_: _dafny.Set
                d_749_v53_ = _dafny.Set({not(self.f28), (self).f29})
                d_750_v54_: _dafny.MultiSet
                d_750_v54_ = _dafny.MultiSet([d_749_v53_, d_749_v53_])
                d_751_v55_: _dafny.Map
                d_751_v55_ = _dafny.Map({p0: not(True)})
                d_752_v56_: _dafny.Seq
                d_752_v56_ = _dafny.SeqWithoutIsStrInference([not((self).f29), ((d_751_v55_)[p0] if (p0) in (d_751_v55_) else self.f28), self.f28, (self).f29])
                (globalState).f24 = (0) - ((0) - ((((d_750_v54_).set(d_749_v53_, default__.abs(len(d_752_v56_)))).set(_dafny.Set({(self).f29}), default__.abs((p0) + ((0) - (p0))))).cardinality))
                (globalState).f1 = p0
            d_753_v57_: _dafny.Seq
            d_753_v57_ = _dafny.SeqWithoutIsStrInference([d_725_v41_])
            d_754_v58_: _dafny.Array
            nw140_ = _dafny.Array(None, 18)
            nw140_[int(0)] = d_725_v41_
            nw140_[int(1)] = d_725_v41_
            nw140_[int(2)] = d_725_v41_
            nw140_[int(3)] = d_725_v41_
            nw140_[int(4)] = d_725_v41_
            nw140_[int(5)] = d_725_v41_
            nw140_[int(6)] = d_725_v41_
            nw140_[int(7)] = d_725_v41_
            nw140_[int(8)] = d_725_v41_
            nw140_[int(9)] = d_725_v41_
            nw140_[int(10)] = d_725_v41_
            nw140_[int(11)] = d_725_v41_
            nw140_[int(12)] = d_725_v41_
            nw140_[int(13)] = d_725_v41_
            nw140_[int(14)] = d_725_v41_
            nw140_[int(15)] = d_725_v41_
            nw140_[int(16)] = d_725_v41_
            nw140_[int(17)] = (d_753_v57_)[default__.safeIndex(p0, len(d_753_v57_))]
            d_754_v58_ = nw140_
            d_754_v58_ = d_754_v58_
            d_755_v59_: C0
            nw141_ = C0()
            nw141_.ctor__()
            d_755_v59_ = nw141_
            nw142_ = C0()
            nw142_.ctor__()
            d_755_v59_ = nw142_
            (globalState).f1 = p0
            d_756_v60_: str
            d_756_v60_ = _dafny.CodePoint('s')
            d_757_v61_: _dafny.Map
            d_757_v61_ = _dafny.Map({23: d_665_v0_})
            d_758_v62_: _dafny.Map
            d_758_v62_ = _dafny.Map({(d_756_v60_) in (((d_757_v61_)[len(d_665_v0_)] if (len(d_665_v0_)) in (d_757_v61_) else _dafny.SeqWithoutIsStrInference([d_756_v60_ for d_759_i7_ in range(default__.abs(922))]))): p0})
            d_758_v62_ = d_758_v62_
        elif True:
            (globalState).f1 = default__.safeModuloInt(147, (p0) + (p0))
            (self).f28 = (self).f29
            if self.f28:
                rhs113_ = default__.safeDivisionInt(-628, len((d_665_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jf")))))
                rhs114_ = p0
                rhs115_ = p0
                rhs116_ = (len((d_665_v0_) + (d_665_v0_))) * (p0)
                lhs96_ = globalState
                lhs97_ = globalState
                lhs98_ = globalState
                lhs99_ = globalState
                lhs96_.f1 = rhs113_
                lhs97_.f24 = rhs114_
                lhs98_.f1 = rhs115_
                lhs99_.f24 = rhs116_
                d_760_v63_: _dafny.Map
                d_760_v63_ = _dafny.Map({(self).f29: (self).f29})
                default__.m0(d_760_v63_, (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))) + (d_665_v0_))), (self).f29, globalState)
                (globalState).f22 = default__.fm38(globalState)
                d_761_v64_: _dafny.Array
                nw143_ = _dafny.Array(False, 3)
                d_761_v64_ = nw143_
                d_762_v65_: _dafny.Seq
                d_762_v65_ = _dafny.SeqWithoutIsStrInference([(d_666_v1_).cf2])
                d_763_v66_: _dafny.MultiSet
                d_763_v66_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0]), d_762_v65_])
                index110_ = default__.safeIndex(817, (d_761_v64_).length(0))
                (d_761_v64_)[index110_] = (d_763_v66_).issubset(d_763_v66_)
                index111_ = default__.safeIndex(817, (d_761_v64_).length(0))
                (d_761_v64_)[index111_] = (p0) == (p0)
                d_764_v67_: _dafny.Set
                d_764_v67_ = _dafny.Set({(self).f29, not((d_761_v64_)[default__.safeIndex(817, (d_761_v64_).length(0))])})
                d_765_v68_: _dafny.Map
                d_765_v68_ = _dafny.Map({self.f28: d_764_v67_})
                d_765_v68_ = d_765_v68_
            elif True:
                d_766_v69_: _dafny.MultiSet
                d_766_v69_ = _dafny.MultiSet([p0, p0])
                (globalState).f1 = (0) - ((d_766_v69_).cardinality)
                d_767_v70_: _dafny.Seq
                d_767_v70_ = _dafny.SeqWithoutIsStrInference([not((self).f29)])
                d_768_v71_: _dafny.Seq
                d_768_v71_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_769_v72_: str
                d_769_v72_ = _dafny.CodePoint('g')
                d_770_v73_: _dafny.Array
                nw144_ = _dafny.Array(None, 11)
                nw144_[int(0)] = self.f28
                nw144_[int(1)] = not(default__.fm38(globalState))
                nw144_[int(2)] = self.f28
                nw144_[int(3)] = (d_767_v70_)[default__.safeIndex(p0, len(d_767_v70_))]
                nw144_[int(4)] = self.f28
                nw144_[int(5)] = ((d_665_v0_).set(default__.safeIndex((d_768_v71_)[default__.safeIndex(p0, len(d_768_v71_))], len(d_665_v0_)), d_769_v72_)) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))
                nw144_[int(6)] = (self).f29
                nw144_[int(7)] = True
                nw144_[int(8)] = self.f28
                nw144_[int(9)] = ((self).f29 if self.f28 else (self).f29)
                nw144_[int(10)] = default__.fm38(globalState)
                d_770_v73_ = nw144_
                d_770_v73_ = d_770_v73_
                d_771_v74_: _dafny.MultiSet
                d_771_v74_ = _dafny.MultiSet([d_666_v1_, d_666_v1_, d_666_v1_, D1_DC2(p0, -719, d_665_v0_, self.f28)])
                d_772_v75_: D7
                d_772_v75_ = D7_DC16(_dafny.MultiSet([d_666_v1_]))
                d_773_v76_: _dafny.Seq
                d_773_v76_ = _dafny.SeqWithoutIsStrInference([d_771_v74_, d_771_v74_, d_771_v74_, (d_772_v75_).cf21])
                d_774_v77_: _dafny.Seq
                d_774_v77_ = _dafny.SeqWithoutIsStrInference([d_666_v1_, d_666_v1_, d_666_v1_, d_666_v1_])
                d_775_v78_: _dafny.Map
                d_775_v78_ = _dafny.Map({not((self).f29): d_771_v74_})
                d_776_v79_: _dafny.Map
                d_776_v79_ = _dafny.Map({p0: (D7_DC16(d_771_v74_)).cf21})
                d_773_v76_ = _dafny.SeqWithoutIsStrInference([d_771_v74_, _dafny.MultiSet(d_774_v77_), ((d_771_v74_).set(d_666_v1_, default__.abs(380))) - (d_771_v74_), (d_771_v74_).intersection(((d_775_v78_)[(self).f29] if ((self).f29) in (d_775_v78_) else ((d_776_v79_)[p0] if (p0) in (d_776_v79_) else d_771_v74_))), d_771_v74_])
                d_777_v80_: _dafny.Map
                d_777_v80_ = _dafny.Map({(self).f29: (self).f29})
                (globalState).f1 = (0) - (default__.safeDivisionInt(p0, len(d_777_v80_)))
                default__.m0((d_777_v80_) | (d_777_v80_), (0) - (default__.safeModuloInt(p0, (0) - (len(_dafny.Map({self.f28: (self).f29}))))), (self).f29, globalState)
            (globalState).f1 = p0
            d_778_v81_: _dafny.Array
            nw145_ = _dafny.Array(None, 17)
            nw145_[int(0)] = self.f28
            nw145_[int(1)] = self.f28
            nw145_[int(2)] = (self).f29
            nw145_[int(3)] = (self).f29
            nw145_[int(4)] = (self).f29
            nw145_[int(5)] = self.f28
            nw145_[int(6)] = self.f28
            nw145_[int(7)] = self.f28
            nw145_[int(8)] = (self).f29
            nw145_[int(9)] = False
            nw145_[int(10)] = (self).f29
            nw145_[int(11)] = not(self.f28)
            nw145_[int(12)] = self.f28
            nw145_[int(13)] = (d_666_v1_).cf5
            nw145_[int(14)] = self.f28
            nw145_[int(15)] = self.f28
            nw145_[int(16)] = not((self).f29)
            d_778_v81_ = nw145_
            d_779_v82_: D5
            d_779_v82_ = D5_DC12(d_778_v81_)
            source14_ = d_779_v82_
            if source14_.is_DC13:
                d_780___mcc_h12_ = source14_.cf18
                d_781_cf18_ = d_780___mcc_h12_
                (globalState).f1 = (-824) + ((0) - (len(d_665_v0_)))
                (globalState).f24 = len(d_665_v0_)
                d_782_v83_: C0
                nw146_ = C0()
                nw146_.ctor__()
                d_782_v83_ = nw146_
                d_783_v84_: _dafny.Array
                nw147_ = _dafny.Array(_dafny.Set({}), 21)
                d_783_v84_ = nw147_
                d_784_v85_: _dafny.Map
                d_784_v85_ = _dafny.Map({p0: d_783_v84_})
                d_784_v85_ = (d_784_v85_).set(p0, d_783_v84_)
            elif True:
                d_785___mcc_h13_ = source14_.cf17
                d_786_cf17_ = d_785___mcc_h13_
                (globalState).f1 = p0
                index112_ = default__.safeIndex(601, (d_725_v41_).length(0))
                (d_725_v41_)[index112_] = 298
                d_787_v86_: _dafny.Array
                def lambda12_(d_788_i8_):
                    return (_dafny.Set({(self).f29}) if (self).f29 else _dafny.Set({(self).f29, True, not(self.f28)}))

                init7_ = lambda12_
                nw148_ = _dafny.Array(None, 28)
                for i0_7_ in range(nw148_.length(0)):
                    nw148_[i0_7_] = init7_(i0_7_)
                d_787_v86_ = nw148_
                d_789_v87_: _dafny.Set
                d_789_v87_ = _dafny.Set({self.f28})
                index113_ = default__.safeIndex(63, (d_787_v86_).length(0))
                (d_787_v86_)[index113_] = d_789_v87_
                d_790_v88_: D4
                d_790_v88_ = D4_DC10((d_666_v1_).cf5, p0, p0)
                index114_ = default__.safeIndex(601, (d_725_v41_).length(0))
                index115_ = default__.safeIndex(63, (d_787_v86_).length(0))
                rhs117_ = ((d_790_v88_).cf14) - (p0)
                rhs118_ = p0
                rhs119_ = d_789_v87_
                lhs100_ = d_725_v41_
                lhs101_ = default__.safeIndex(601, (d_725_v41_).length(0))
                lhs102_ = globalState
                lhs103_ = d_787_v86_
                lhs104_ = default__.safeIndex(63, (d_787_v86_).length(0))
                lhs100_[lhs101_] = rhs117_
                lhs102_.f24 = rhs118_
                lhs103_[lhs104_] = rhs119_
                d_791_v89_: C0
                nw149_ = C0()
                nw149_.ctor__()
                d_791_v89_ = nw149_
                d_792_v90_: _dafny.Array
                nw150_ = _dafny.Array(_dafny.MultiSet({}), 28)
                d_792_v90_ = nw150_
                d_793_v91_: str
                d_793_v91_ = _dafny.CodePoint('j')
                d_794_v92_: _dafny.MultiSet
                d_794_v92_ = _dafny.MultiSet([d_793_v91_, d_793_v91_, d_793_v91_, d_793_v91_, (d_665_v0_)[default__.safeIndex((0) - ((d_725_v41_)[default__.safeIndex(601, (d_725_v41_).length(0))]), len(d_665_v0_))]])
                index116_ = default__.safeIndex(288, (d_792_v90_).length(0))
                (d_792_v90_)[index116_] = d_794_v92_
                d_795_v93_: _dafny.Map
                d_795_v93_ = _dafny.Map({self.f28: (d_725_v41_)[default__.safeIndex(601, (d_725_v41_).length(0))]})
                d_796_v94_: _dafny.Map
                d_796_v94_ = _dafny.Map({not (False) or (False): (d_795_v93_) | (d_795_v93_)})
                d_797_v95_: _dafny.Set
                d_797_v95_ = _dafny.Set({p0, (d_725_v41_)[default__.safeIndex(601, (d_725_v41_).length(0))]})
                d_798_v96_: _dafny.Set
                d_798_v96_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx"))})
                index117_ = default__.safeIndex(288, (d_792_v90_).length(0))
                rhs120_ = (default__.fm41((self).f29, globalState)).intersection(_dafny.MultiSet([default__.fm42((self).fm36(len(d_797_v95_), d_798_v96_, globalState), len((d_787_v86_)[default__.safeIndex(63, (d_787_v86_).length(0))]), globalState)]))
                rhs121_ = d_796_v94_
                rhs122_ = d_665_v0_
                lhs105_ = d_792_v90_
                lhs106_ = default__.safeIndex(288, (d_792_v90_).length(0))
                lhs105_[lhs106_] = rhs120_
                d_796_v94_ = rhs121_
                d_665_v0_ = rhs122_
        d_799_v97_: _dafny.Map
        d_799_v97_ = _dafny.Map({p0: self.f28})
        d_800_v98_: D4
        d_800_v98_ = D4_DC10(self.f28, len(d_665_v0_), (0) - (len(d_799_v97_)))
        (globalState).f1 = (d_800_v98_).cf14
        (globalState).f27 = (self).f29
        d_801_v99_: _dafny.Seq
        d_801_v99_ = _dafny.SeqWithoutIsStrInference([p0])
        d_802_v100_: _dafny.Map
        d_802_v100_ = _dafny.Map({not((self).f29): d_801_v99_})
        d_803_v101_: _dafny.Set
        d_803_v101_ = _dafny.Set({(0) - (len(d_802_v100_)), p0})
        d_804_v102_: _dafny.Seq
        d_804_v102_ = _dafny.SeqWithoutIsStrInference([(self).f29])
        d_805_v103_: _dafny.Map
        d_805_v103_ = _dafny.Map({len(d_804_v102_): p0})
        d_806_v104_: _dafny.Map
        d_806_v104_ = _dafny.Map({d_801_v99_: d_805_v103_})
        d_807_v105_: _dafny.Map
        d_807_v105_ = _dafny.Map({(0) - (p0): len(d_806_v104_)})
        d_808_v107_: _dafny.MultiSet
        def iife80_():
            coll42_ = _dafny.Set()
            compr_42_: int
            for compr_42_ in (default__.fm37(len(d_805_v103_), self.f28, self.f28, globalState)).Elements:
                d_809_v106_: int = compr_42_
                if (d_809_v106_) in (default__.fm37(len(d_805_v103_), self.f28, self.f28, globalState)):
                    coll42_ = coll42_.union(_dafny.Set([(d_809_v106_) + (211)]))
            return _dafny.Set(coll42_)
        d_808_v107_ = _dafny.MultiSet([d_803_v101_, iife80_()
        ])
        r0 = (d_808_v107_) == (d_808_v107_)
        return r0


class C2(T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def fm52(self, p0, globalState):
        return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([43]))])) - (_dafny.MultiSet([len(_dafny.Set({847}))]))).intersection((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qypw")))}))]))])).intersection(_dafny.MultiSet([-529, 489])))

    def fm53(self, globalState):
        def iife81_():
            coll43_ = _dafny.Map()
            compr_43_: int
            for compr_43_ in _dafny.IntegerRange(226, 776):
                d_810_v0_: int = compr_43_
                if ((226) <= (d_810_v0_)) and ((d_810_v0_) < (776)):
                    coll43_[default__.safeModuloInt(d_810_v0_, 513)] = False
            return _dafny.Map(coll43_)
        def iife82_():
            coll44_ = _dafny.Map()
            compr_44_: int
            for compr_44_ in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, False])).cardinality, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([809])): _dafny.CodePoint('g')}))])).Elements:
                d_811_v1_: int = compr_44_
                if (d_811_v1_) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, False])).cardinality, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([809])): _dafny.CodePoint('g')}))])):
                    coll44_[(d_811_v1_) * (377)] = False
            return _dafny.Map(coll44_)
        return ((_dafny.MultiSet([iife81_()
        ])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D10_DC26(iife82_()
)).cf39 for d_812_i0_ in range(default__.abs(468))])))).isdisjoint(_dafny.MultiSet([_dafny.Map({(_dafny.MultiSet([True])).cardinality: True})]))

    def m13(self, p0, p1, p2, p3, globalState):
        d_813_v0_: _dafny.Array
        nw151_ = _dafny.Array(_dafny.CodePoint('D'), 4)
        d_813_v0_ = nw151_
        d_814_v1_: _dafny.Array
        d_814_v1_ = d_813_v0_
        d_815_v2_: _dafny.Array
        nw152_ = _dafny.Array(None, 21)
        nw152_[int(0)] = d_813_v0_
        nw152_[int(1)] = d_813_v0_
        nw152_[int(2)] = (d_814_v1_)
        nw152_[int(3)] = d_813_v0_
        nw152_[int(4)] = d_813_v0_
        nw152_[int(5)] = d_813_v0_
        nw152_[int(6)] = d_813_v0_
        nw152_[int(7)] = (d_813_v0_ if p0 else d_813_v0_)
        nw152_[int(8)] = d_813_v0_
        nw152_[int(9)] = d_813_v0_
        nw152_[int(10)] = d_813_v0_
        nw152_[int(11)] = d_813_v0_
        nw152_[int(12)] = d_813_v0_
        nw152_[int(13)] = d_813_v0_
        nw152_[int(14)] = d_813_v0_
        nw152_[int(15)] = (d_813_v0_ if p3 else d_813_v0_)
        nw152_[int(16)] = d_813_v0_
        nw152_[int(17)] = (d_813_v0_ if p0 else d_813_v0_)
        nw152_[int(18)] = d_813_v0_
        nw152_[int(19)] = d_813_v0_
        nw152_[int(20)] = d_813_v0_
        d_815_v2_ = nw152_
        d_815_v2_ = (d_815_v2_ if not(p0) else d_815_v2_)
        d_816_v3_: str
        d_816_v3_ = _dafny.CodePoint('s')
        d_817_v4_: _dafny.MultiSet
        d_817_v4_ = _dafny.MultiSet([d_816_v3_, _dafny.CodePoint('u')])
        d_818_v5_: int
        d_818_v5_ = 987
        (globalState).f24 = default__.safeModuloInt(((d_817_v4_).cardinality) * (d_818_v5_), (d_818_v5_) * (d_818_v5_))
        d_819_v6_: D6
        d_819_v6_ = D6_DC15(p3)
        d_820_v7_: C1
        nw153_ = C1()
        nw153_.ctor__(p0, (d_819_v6_).cf20)
        d_820_v7_ = nw153_
        d_821_v8_: D8
        d_821_v8_ = D8_DC21(p3)
        if (d_821_v8_).cf27:
            (globalState).f22 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glef"))) != (p2)
            (globalState).f1 = d_818_v5_
            d_822_v9_: _dafny.Set
            d_822_v9_ = _dafny.Set({p3})
            d_823_v10_: _dafny.Map
            d_823_v10_ = _dafny.Map({p0: not(p3)})
            d_824_v11_: _dafny.Map
            d_824_v11_ = _dafny.Map({p3: d_818_v5_})
            d_825_v12_: _dafny.MultiSet
            d_825_v12_ = _dafny.MultiSet([p3, True])
            d_826_v13_: _dafny.Seq
            d_826_v13_ = _dafny.SeqWithoutIsStrInference([p3, p3])
            d_827_v14_: _dafny.Seq
            d_827_v14_ = _dafny.SeqWithoutIsStrInference([d_818_v5_, len(_dafny.Set({d_818_v5_, d_818_v5_, d_818_v5_})), d_818_v5_])
            d_828_v15_: _dafny.Array
            nw154_ = _dafny.Array(None, 17)
            nw154_[int(0)] = d_818_v5_
            nw154_[int(1)] = d_818_v5_
            nw154_[int(2)] = d_818_v5_
            nw154_[int(3)] = d_818_v5_
            nw154_[int(4)] = len(d_822_v9_)
            nw154_[int(5)] = default__.safeDivisionInt(d_818_v5_, (0) - (d_818_v5_))
            nw154_[int(6)] = d_818_v5_
            nw154_[int(7)] = (d_818_v5_ if p3 else d_818_v5_)
            nw154_[int(8)] = d_818_v5_
            nw154_[int(9)] = default__.safeDivisionInt(d_818_v5_, d_818_v5_)
            nw154_[int(10)] = (0) - (d_818_v5_)
            nw154_[int(11)] = len(d_823_v10_)
            nw154_[int(12)] = ((d_824_v11_)[p0] if (p0) in (d_824_v11_) else d_818_v5_)
            nw154_[int(13)] = (d_825_v12_).cardinality
            nw154_[int(14)] = len((d_826_v13_) + ((d_826_v13_).set(default__.safeIndex(d_818_v5_, len(d_826_v13_)), p3)))
            nw154_[int(15)] = d_818_v5_
            nw154_[int(16)] = (d_827_v14_)[default__.safeIndex(d_818_v5_, len(d_827_v14_))]
            d_828_v15_ = nw154_
            index118_ = default__.safeIndex(855, (d_828_v15_).length(0))
            (d_828_v15_)[index118_] = default__.safeDivisionInt(d_818_v5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfidqb"))))
            index119_ = default__.safeIndex(855, (d_828_v15_).length(0))
            def iife83_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(989, 565):
                    d_829_v16_: int = compr_45_
                    if ((989) <= (d_829_v16_)) and ((d_829_v16_) < (565)):
                        coll45_[default__.safeModuloInt(d_829_v16_, -746)] = d_818_v5_
                return _dafny.Map(coll45_)
            (d_828_v15_)[index119_] = len(iife83_()
            )
            (globalState).f1 = d_818_v5_
            d_830_v17_: D9
            d_830_v17_ = D9_DC25(p0, d_818_v5_, d_818_v5_, p0, -685)
            (globalState).f24 = (0) - ((d_830_v17_).cf35)
        elif True:
            d_831_v18_: _dafny.Array
            nw155_ = _dafny.Array(None, 5)
            nw155_[int(0)] = not (p0) or (p0)
            nw155_[int(1)] = p0
            nw155_[int(2)] = True
            nw155_[int(3)] = p3
            nw155_[int(4)] = p3
            d_831_v18_ = nw155_
            d_832_v19_: _dafny.MultiSet
            d_832_v19_ = _dafny.MultiSet([d_818_v5_])
            d_833_v20_: _dafny.Array
            nw156_ = _dafny.Array(int(0), 12)
            d_833_v20_ = nw156_
            d_834_v21_: _dafny.Map
            d_834_v21_ = _dafny.Map({d_818_v5_: d_833_v20_})
            index120_ = default__.safeIndex(792, (d_831_v18_).length(0))
            (d_831_v18_)[index120_] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_818_v5_, -141]))).set(len(d_834_v21_), default__.abs(d_818_v5_))).issubset(d_832_v19_)
            index121_ = default__.safeIndex(792, (d_831_v18_).length(0))
            (d_831_v18_)[index121_] = (d_818_v5_) > ((d_832_v19_).cardinality)
            d_818_v5_ = d_818_v5_
            d_818_v5_ = d_818_v5_
            d_835_v22_: _dafny.Map
            d_835_v22_ = _dafny.Map({d_818_v5_: (d_831_v18_)[default__.safeIndex(792, (d_831_v18_).length(0))]})
            d_836_v23_: D10
            d_836_v23_ = D10_DC26((d_835_v22_) | (d_835_v22_))
            d_836_v23_ = D10_DC26(d_835_v22_)
            (globalState).f1 = d_818_v5_
        d_837_v24_: _dafny.Array
        nw157_ = _dafny.Array(int(0), 9)
        d_837_v24_ = nw157_
        d_838_v25_: _dafny.Array
        d_838_v25_ = d_837_v24_
        d_839_v26_: _dafny.Map
        d_839_v26_ = _dafny.Map({d_813_v0_: d_838_v25_})
        d_839_v26_ = (d_839_v26_).set(d_813_v0_, d_838_v25_)
        d_840_v27_: _dafny.Array
        nw158_ = _dafny.Array(False, 22)
        d_840_v27_ = nw158_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_840_v27_).length(0)):
            d_841_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_841_i0_)) and ((d_841_i0_) < ((d_840_v27_).length(0)))):
                (d_840_v27_)[(d_841_i0_)] = p0


class C3(T1, T0, T2):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f39: int = int(0)
        self._f40: bool = False
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
    @property
    def f39(self):
        return self._f39
    def ctor__(self, f40, f28, f29, f39):
        (self)._f40 = f40
        (self).f28 = f28
        (self)._f29 = f29
        (self)._f39 = f39

    def fm30(self, p0, globalState):
        return (self).f39

    def fm31(self, globalState):
        return self.f28

    def fm32(self, p0, p1, p2, globalState):
        return self.f28

    def m13(self, p0, p1, p2, p3, globalState):
        if (self).f29:
            d_842_v0_: _dafny.Map
            d_842_v0_ = _dafny.Map({(self).f40: p3})
            d_843_v1_: _dafny.Map
            d_843_v1_ = _dafny.Map({(self).f39: (self).f39})
            d_844_v2_: _dafny.MultiSet
            d_844_v2_ = _dafny.MultiSet([d_843_v1_])
            (globalState).f13 = ((d_842_v0_)[(d_844_v2_).ispropersubset(d_844_v2_)] if ((d_844_v2_).ispropersubset(d_844_v2_)) in (d_842_v0_) else p0)
            d_845_v3_: _dafny.Seq
            d_845_v3_ = _dafny.SeqWithoutIsStrInference([p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpx")), p2])
            d_846_v4_: _dafny.Map
            d_846_v4_ = _dafny.Map({(d_845_v3_) + (d_845_v3_): 841})
            d_846_v4_ = (d_846_v4_).set(d_845_v3_, (0) - ((self).f39))
            d_847_v5_: C0
            nw159_ = C0()
            nw159_.ctor__()
            d_847_v5_ = nw159_
            d_848_v8_: _dafny.Array
            def lambda13_(d_849_i0_):
                def iife84_():
                    coll46_ = _dafny.Set()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(567, 294):
                        d_850_v6_: int = compr_46_
                        if ((567) <= (d_850_v6_)) and ((d_850_v6_) < (294)):
                            coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_850_v6_, (self).f39)]))
                    return _dafny.Set(coll46_)
                def iife85_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in (_dafny.MultiSet([(self).f39, (self).f39])).Elements:
                        d_851_v7_: int = compr_47_
                        if (d_851_v7_) in (_dafny.MultiSet([(self).f39, (self).f39])):
                            coll47_ = coll47_.union(_dafny.Set([default__.safeDivisionInt(d_851_v7_, 514)]))
                    return _dafny.Set(coll47_)
                return (iife84_()
                ).intersection(iife85_()
                )

            init8_ = lambda13_
            nw160_ = _dafny.Array(None, 1)
            for i0_8_ in range(nw160_.length(0)):
                nw160_[i0_8_] = init8_(i0_8_)
            d_848_v8_ = nw160_
            index122_ = default__.safeIndex(711, (d_848_v8_).length(0))
            (d_848_v8_)[index122_] = p1
            index123_ = default__.safeIndex(711, (d_848_v8_).length(0))
            (d_848_v8_)[index123_] = default__.fm35(globalState)
            d_852_v9_: T0
            nw161_ = C1()
            nw161_.ctor__(p0, (self).f29)
            d_852_v9_ = nw161_
            d_853_v10_: D8
            d_853_v10_ = D8_DC20(d_852_v9_)
            d_854_v11_: D8
            d_854_v11_ = D8_DC20((d_853_v10_).cf26)
            rhs123_ = (d_853_v10_).cf26
            rhs124_ = p0
            rhs125_ = (self).f39
            rhs126_ = (self).f39
            rhs127_ = default__.fm42((self).f39, (self).f39, globalState)
            lhs107_ = globalState
            lhs108_ = globalState
            lhs109_ = globalState
            lhs110_ = globalState
            d_852_v9_ = rhs123_
            lhs107_.f22 = rhs124_
            lhs108_.f1 = rhs125_
            lhs109_.f24 = rhs126_
            lhs110_.f23 = rhs127_
        elif True:
            d_855_v12_: str
            d_855_v12_ = _dafny.CodePoint('q')
            (globalState).f23 = d_855_v12_
            (globalState).f22 = (not ((self).f29) or (p0)) or (self.f28)
            if self.f28:
                (globalState).f24 = (self).f39
                d_856_v13_: D8
                d_856_v13_ = D8_DC21((self).f40)
                d_856_v13_ = default__.fm43(p0, ((self).f39) != ((_dafny.MultiSet([_dafny.CodePoint('w')])).cardinality), (self).fm30(p1, globalState), globalState)
                d_857_v14_: _dafny.Seq
                d_857_v14_ = _dafny.SeqWithoutIsStrInference([p2])
                d_858_v15_: _dafny.Seq
                d_858_v15_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_855_v12_ for d_859_i1_ in range(default__.abs(819))]), (d_857_v14_)[default__.safeIndex((self).f39, len(d_857_v14_))]])
                d_860_v16_: _dafny.Map
                d_860_v16_ = _dafny.Map({p2: ((self).f39) > (len(d_858_v15_))})
                d_860_v16_ = (d_860_v16_).set(p2, not(p0))
                (globalState).f1 = (self).f39
                d_861_v17_: _dafny.Seq
                d_861_v17_ = _dafny.SeqWithoutIsStrInference([p0, self.f28])
                (globalState).f1 = default__.safeModuloInt((self).fm30(p1, globalState), (0) - (len(d_861_v17_)))
            elif True:
                d_862_v18_: C1
                nw162_ = C1()
                nw162_.ctor__(p3, True)
                d_862_v18_ = nw162_
                d_862_v18_ = d_862_v18_
                d_863_v19_: _dafny.Seq
                d_863_v19_ = _dafny.SeqWithoutIsStrInference([(p1).ispropersubset(p1), (self).f40])
                (globalState).f1 = (_dafny.MultiSet(d_863_v19_)).cardinality
                d_864_v20_: D1
                d_864_v20_ = D1_DC2(len(_dafny.SeqWithoutIsStrInference([(self).f40, not(self.f28)])), (self).f39, p2, not(self.f28))
                d_865_v21_: _dafny.Map
                d_865_v21_ = _dafny.Map({len(_dafny.Map({(self).f40: _dafny.Map({p0: (self).f29})})): d_864_v20_})
                d_865_v21_ = (d_865_v21_).set(((self).f39) * (-937), D1_DC2((self).f39, (self).f39, p2, self.f28))
                d_866_v22_: _dafny.Seq
                d_866_v22_ = _dafny.SeqWithoutIsStrInference([(self).f39, (self).f39])
                d_867_v23_: _dafny.Seq
                d_867_v23_ = _dafny.SeqWithoutIsStrInference([(self).f39, len(d_866_v22_)])
                d_868_v24_: C1
                nw163_ = C1()
                nw163_.ctor__(self.f28, not(((self).f39) not in (d_867_v23_)))
                d_868_v24_ = nw163_
                (globalState).f23 = d_855_v12_
            d_869_v25_: _dafny.MultiSet
            d_869_v25_ = _dafny.MultiSet([not(p3)])
            d_870_v26_: _dafny.Map
            d_870_v26_ = _dafny.Map({((0) - ((self).f39)) + ((self).f39): d_869_v25_})
            d_870_v26_ = d_870_v26_
            d_871_v27_: _dafny.Seq
            d_871_v27_ = _dafny.SeqWithoutIsStrInference([p3])
            d_872_v28_: _dafny.Map
            d_872_v28_ = _dafny.Map({(self).f39: (d_871_v27_)[default__.safeIndex(-687, len(d_871_v27_))]})
            d_873_v29_: int
            d_874_v30_: int
            d_875_v31_: int
            d_876_v32_: bool
            out23_: int
            out24_: int
            out25_: int
            out26_: bool
            out23_, out24_, out25_, out26_ = (self).m17((len(d_872_v28_)) > ((self).fm30(p1, globalState)), default__.fm44(False, globalState), self.f28, globalState)
            d_873_v29_ = out23_
            d_874_v30_ = out24_
            d_875_v31_ = out25_
            d_876_v32_ = out26_
        d_877_v33_: C0
        nw164_ = C0()
        nw164_.ctor__()
        d_877_v33_ = nw164_
        d_878_v34_: _dafny.Map
        d_878_v34_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not((self).f29)]): d_877_v33_})
        d_879_v35_: str
        d_879_v35_ = _dafny.CodePoint('v')
        d_880_v36_: _dafny.Map
        d_880_v36_ = _dafny.Map({(self).f39: _dafny.SeqWithoutIsStrInference([d_879_v35_])})
        d_881_v37_: _dafny.Seq
        d_881_v37_ = _dafny.SeqWithoutIsStrInference([(self).f39])
        rhs128_ = (d_878_v34_).set(_dafny.SeqWithoutIsStrInference([(d_877_v33_).fm33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbpbcjqon")), d_880_v36_, _dafny.MultiSet([504]), globalState), p3, (self).f29, False]), d_877_v33_)
        rhs129_ = default__.fm42((self).f39, (self).f39, globalState)
        rhs130_ = (d_881_v37_)[default__.safeIndex(((self).f39) * ((self).f39), len(d_881_v37_))]
        rhs131_ = 432
        lhs111_ = globalState
        lhs112_ = globalState
        lhs113_ = globalState
        d_878_v34_ = rhs128_
        lhs111_.f23 = rhs129_
        lhs112_.f1 = rhs130_
        lhs113_.f24 = rhs131_
        d_882_v38_: D2
        d_882_v38_ = D2_DC5((d_877_v33_).fm33(p2, _dafny.Map({(self).f39: p2}), _dafny.MultiSet(d_881_v37_), globalState))
        d_883_v39_: D2
        d_883_v39_ = D2_DC6(d_882_v38_)
        pat_let_tv25_ = p0
        pat_let_tv26_ = d_881_v37_
        pat_let_tv27_ = d_879_v35_
        pat_let_tv28_ = d_879_v35_
        def lambda14_(source15_):
            if source15_.is_DC13:
                d_884___mcc_h0_ = source15_.cf18
                d_885_cf18_ = d_884___mcc_h0_
                return not(pat_let_tv25_)
            elif True:
                d_886___mcc_h1_ = source15_.cf17
                d_887_cf17_ = d_886___mcc_h1_
                return (len(pat_let_tv26_)) != ((self).f39)

        def lambda15_(source16_):
            if source16_.is_DC5:
                d_888___mcc_h2_ = source16_.cf8
                d_889_cf8_ = d_888___mcc_h2_
                return pat_let_tv27_
            elif source16_.is_DC4:
                d_890___mcc_h3_ = source16_.cf7
                d_891_cf7_ = d_890___mcc_h3_
                return pat_let_tv28_
            elif True:
                d_892___mcc_h4_ = source16_.cf9
                d_893_cf9_ = d_892___mcc_h4_
                return _dafny.CodePoint('l')

        rhs132_ = False
        rhs133_ = default__.fm42((self).f39, ((self).f39) + ((self).f39), globalState)
        rhs134_ = not(lambda14_(default__.fm44((self).f29, globalState)))
        rhs135_ = lambda15_(d_883_v39_)
        lhs114_ = globalState
        lhs115_ = globalState
        lhs116_ = globalState
        lhs114_.f22 = rhs132_
        lhs115_.f23 = rhs133_
        lhs116_.f13 = rhs134_
        d_879_v35_ = rhs135_
        source17_ = default__.fm39(p0, globalState)
        if source17_.is_DC10:
            d_894___mcc_h5_ = source17_.cf13
            d_895___mcc_h6_ = source17_.cf14
            d_896___mcc_h7_ = source17_.cf15
            d_897_cf15_ = d_896___mcc_h7_
            d_898_cf14_ = d_895___mcc_h6_
            d_899_cf13_ = d_894___mcc_h5_
            d_900_v40_: D8
            d_900_v40_ = D8_DC22()
            source18_ = d_900_v40_
            if source18_.is_DC21:
                d_901___mcc_h10_ = source18_.cf27
                d_902_cf27_ = d_901___mcc_h10_
                d_903_v41_: _dafny.Map
                d_903_v41_ = _dafny.Map({(self).f39: d_902_cf27_})
                d_903_v41_ = default__.fm45(globalState)
                d_904_v42_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.MultiSet({}), 23)
                d_904_v42_ = nw165_
                d_905_v43_: _dafny.MultiSet
                d_905_v43_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_879_v35_ for d_906_i2_ in range(default__.abs(788))]), p2])
                index124_ = default__.safeIndex(111, (d_904_v42_).length(0))
                (d_904_v42_)[index124_] = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")), p2, p2, p2])).intersection(d_905_v43_)
                d_907_v44_: _dafny.Set
                d_907_v44_ = _dafny.Set({self.f28, (self).f40, (self).fm31(globalState), self.f28})
                d_908_v45_: _dafny.Map
                d_908_v45_ = _dafny.Map({len(d_907_v44_): (d_905_v43_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxwqrag")), default__.abs(d_897_cf15_))})
                d_909_v46_: _dafny.Array
                nw166_ = _dafny.Array(int(0), 12)
                d_909_v46_ = nw166_
                d_910_v47_: _dafny.Array
                d_910_v47_ = d_909_v46_
                d_911_v48_: _dafny.MultiSet
                d_911_v48_ = _dafny.MultiSet([d_910_v47_])
                d_912_v49_: _dafny.Map
                d_912_v49_ = _dafny.Map({d_911_v48_: d_905_v43_})
                index125_ = default__.safeIndex(111, (d_904_v42_).length(0))
                (d_904_v42_)[index125_] = (((d_908_v45_)[len(p2)] if (len(p2)) in (d_908_v45_) else d_905_v43_)) - (((d_912_v49_)[_dafny.MultiSet([d_910_v47_, d_910_v47_])] if (_dafny.MultiSet([d_910_v47_, d_910_v47_])) in (d_912_v49_) else d_905_v43_))
                (globalState).f1 = (0) - (d_898_cf14_)
                d_913_v50_: _dafny.Seq
                d_913_v50_ = _dafny.SeqWithoutIsStrInference([d_899_cf13_, p0])
                d_914_v51_: _dafny.Map
                d_914_v51_ = _dafny.Map({d_913_v50_: (self).f39})
                d_915_v52_: _dafny.Map
                d_915_v52_ = _dafny.Map({d_898_cf14_: d_877_v33_})
                d_916_v53_: _dafny.Map
                d_916_v53_ = _dafny.Map({(d_881_v37_)[default__.safeIndex(len(d_915_v52_), len(d_881_v37_))]: d_897_cf15_})
                d_914_v51_ = (d_914_v51_).set(d_913_v50_, ((d_916_v53_)[d_898_cf14_] if (d_898_cf14_) in (d_916_v53_) else (self).f39))
            elif source18_.is_DC22:
                d_917_v54_: _dafny.MultiSet
                d_917_v54_ = _dafny.MultiSet([D1_DC2(d_898_cf14_, (self).f39, p2, p0)])
                d_918_v55_: D7
                d_918_v55_ = D7_DC16(d_917_v54_)
                d_919_v56_: _dafny.Seq
                d_919_v56_ = _dafny.SeqWithoutIsStrInference([d_918_v55_, d_918_v55_])
                (globalState).f22 = (d_919_v56_) == ((d_919_v56_) + (d_919_v56_))
                d_920_v57_: _dafny.Seq
                d_920_v57_ = _dafny.SeqWithoutIsStrInference([True])
                d_921_v58_: _dafny.Array
                nw167_ = _dafny.Array(None, 25)
                nw167_[int(0)] = True
                nw167_[int(1)] = p3
                nw167_[int(2)] = (self).f29
                nw167_[int(3)] = d_899_cf13_
                nw167_[int(4)] = p3
                nw167_[int(5)] = (self).f40
                nw167_[int(6)] = (d_920_v57_)[default__.safeIndex(d_897_cf15_, len(d_920_v57_))]
                nw167_[int(7)] = p0
                nw167_[int(8)] = p0
                nw167_[int(9)] = default__.fm38(globalState)
                nw167_[int(10)] = p0
                nw167_[int(11)] = (self).f29
                nw167_[int(12)] = (self).f40
                nw167_[int(13)] = (self).f29
                nw167_[int(14)] = self.f28
                nw167_[int(15)] = p0
                nw167_[int(16)] = d_899_cf13_
                nw167_[int(17)] = self.f28
                nw167_[int(18)] = (self).f29
                nw167_[int(19)] = p0
                nw167_[int(20)] = d_899_cf13_
                nw167_[int(21)] = d_899_cf13_
                nw167_[int(22)] = (self).f29
                nw167_[int(23)] = self.f28
                nw167_[int(24)] = p0
                d_921_v58_ = nw167_
                d_922_v59_: _dafny.Map
                d_922_v59_ = _dafny.Map({d_921_v58_: (self).f40})
                d_922_v59_ = (d_922_v59_).set(d_921_v58_, not(d_899_cf13_))
                d_899_cf13_ = ((d_920_v57_).set(default__.safeIndex(d_898_cf14_, len(d_920_v57_)), self.f28)) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f29, p3, True, p3]) for d_923_i3_ in range(default__.abs(960))]))
                d_924_v60_: _dafny.MultiSet
                d_924_v60_ = _dafny.MultiSet([586, d_897_cf15_])
                d_925_v61_: _dafny.Map
                d_925_v61_ = _dafny.Map({(34) > (309): (d_924_v60_ if (d_877_v33_).fm33(p2, d_880_v36_, d_924_v60_, globalState) else _dafny.MultiSet([d_898_cf14_, len(p2)]))})
                d_925_v61_ = (d_925_v61_).set(((self).f40 if self.f28 else (self).f40), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f39 for d_926_i4_ in range(default__.abs(714))])))
            elif True:
                d_927___mcc_h11_ = source18_.cf26
                d_928_cf26_ = d_927___mcc_h11_
                d_929_v62_: _dafny.Map
                d_929_v62_ = _dafny.Map({p0: (0) - ((self).f39)})
                (globalState).f13 = (len((d_929_v62_) | (d_929_v62_))) >= ((self).fm30(p1, globalState))
                (globalState).f13 = ((self).f39) == (d_898_cf14_)
                d_930_v63_: _dafny.Array
                nw168_ = _dafny.Array(int(0), 1)
                d_930_v63_ = nw168_
                d_931_v64_: D1
                d_931_v64_ = D1_DC2((_dafny.MultiSet([d_897_cf15_])).cardinality, (0) - ((self).f39), p2, (self).f40)
                index126_ = default__.safeIndex(362, (d_930_v63_).length(0))
                (d_930_v63_)[index126_] = (0) - ((d_931_v64_).cf3)
                d_932_v65_: _dafny.Seq
                d_932_v65_ = _dafny.SeqWithoutIsStrInference([d_930_v63_])
                d_933_v66_: _dafny.Map
                d_933_v66_ = _dafny.Map({d_899_cf13_: d_879_v35_})
                index127_ = default__.safeIndex(362, (d_930_v63_).length(0))
                rhs136_ = len(d_932_v65_)
                rhs137_ = ((self).f39) > ((d_898_cf14_ if self.f28 else d_897_cf15_))
                rhs138_ = len(d_933_v66_)
                lhs117_ = globalState
                lhs118_ = globalState
                lhs119_ = d_930_v63_
                lhs120_ = default__.safeIndex(362, (d_930_v63_).length(0))
                lhs117_.f1 = rhs136_
                lhs118_.f27 = rhs137_
                lhs119_[lhs120_] = rhs138_
                d_934_v67_: _dafny.Map
                d_934_v67_ = _dafny.Map({(_dafny.MultiSet(default__.fm46(d_928_cf26_.f28, globalState))).cardinality: (self).f40})
                d_934_v67_ = (d_934_v67_).set((0) - (d_897_cf15_), d_928_cf26_.f28)
            d_935_v68_: _dafny.Array
            def lambda16_(d_936_i5_):
                return (d_936_i5_) * (-661)

            init9_ = lambda16_
            nw169_ = _dafny.Array(None, 16)
            for i0_9_ in range(nw169_.length(0)):
                nw169_[i0_9_] = init9_(i0_9_)
            d_935_v68_ = nw169_
            index128_ = default__.safeIndex(438, (d_935_v68_).length(0))
            (d_935_v68_)[index128_] = ((self).f39) * (d_898_cf14_)
            index129_ = default__.safeIndex(438, (d_935_v68_).length(0))
            (d_935_v68_)[index129_] = len(default__.fm40(globalState))
            if self.f28:
                (globalState).f21 = d_899_cf13_
                (globalState).f1 = (d_935_v68_)[default__.safeIndex(438, (d_935_v68_).length(0))]
                (globalState).f1 = ((self).fm30(p1, globalState)) + (len(p2))
                d_937_v69_: _dafny.Map
                d_937_v69_ = _dafny.Map({322: (self).f40})
                d_938_v70_: _dafny.Seq
                d_938_v70_ = _dafny.SeqWithoutIsStrInference([self.f28, ((d_937_v69_)[530] if (530) in (d_937_v69_) else p3), self.f28, ((self).f40) or ((self).f40)])
                d_939_v72_: D6
                d_939_v72_ = D6_DC15(self.f28)
                d_940_v73_: _dafny.Seq
                d_940_v73_ = _dafny.SeqWithoutIsStrInference([d_939_v72_])
                d_941_v74_: _dafny.Seq
                def iife86_():
                    coll48_ = _dafny.Map()
                    compr_48_: D6
                    for compr_48_ in ((d_940_v73_).set(default__.safeIndex(201, len(d_940_v73_)), d_939_v72_)).Elements:
                        d_942_v71_: D6 = compr_48_
                        if (d_942_v71_) in ((d_940_v73_).set(default__.safeIndex(201, len(d_940_v73_)), d_939_v72_)):
                            coll48_[d_942_v71_] = d_879_v35_
                    return _dafny.Map(coll48_)
                d_941_v74_ = _dafny.SeqWithoutIsStrInference([iife86_()
                ])
                d_943_v75_: _dafny.Seq
                d_943_v75_ = _dafny.SeqWithoutIsStrInference([(d_941_v74_)[default__.safeIndex(350, len(d_941_v74_))]])
                d_944_v76_: _dafny.Map
                d_944_v76_ = _dafny.Map({d_939_v72_: d_879_v35_})
                d_945_v77_: _dafny.MultiSet
                d_945_v77_ = _dafny.MultiSet([d_944_v76_, d_944_v76_])
                rhs139_ = ((_dafny.MultiSet(d_941_v74_)).set(d_944_v76_, default__.abs(704))).ispropersubset(d_945_v77_)
                rhs140_ = (d_938_v70_) + (d_938_v70_)
                rhs141_ = (d_935_v68_)
                lhs121_ = globalState
                lhs121_.f27 = rhs139_
                d_938_v70_ = rhs140_
                d_935_v68_ = rhs141_
                d_946_v78_: _dafny.Map
                d_946_v78_ = _dafny.Map({not(d_899_cf13_): not(d_899_cf13_)})
                d_947_v79_: _dafny.Seq
                d_947_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjgokfknf"))
                rhs142_ = True
                rhs143_ = d_946_v78_
                rhs144_ = (p1).ispropersubset((p1) - (_dafny.Set({(self).f39})))
                rhs145_ = p2
                rhs146_ = (self).f40
                lhs122_ = globalState
                lhs123_ = globalState
                lhs122_.f13 = rhs142_
                d_946_v78_ = rhs143_
                d_899_cf13_ = rhs144_
                d_947_v79_ = rhs145_
                lhs123_.f22 = rhs146_
            elif True:
                (globalState).f1 = (self).f39
                d_948_v80_: _dafny.Map
                d_948_v80_ = _dafny.Map({685: d_897_cf15_})
                d_948_v80_ = (d_948_v80_).set((self).fm30(p1, globalState), d_898_cf14_)
                (globalState).f1 = (self).fm30(p1, globalState)
                d_949_v81_: _dafny.Array
                nw170_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_949_v81_ = nw170_
                nw171_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_949_v81_ = nw171_
                d_950_v82_: C0
                nw172_ = C0()
                nw172_.ctor__()
                d_950_v82_ = nw172_
            d_898_cf14_ = default__.safeDivisionInt(len(p2), len(p2))
        elif source17_.is_DC9:
            d_951___mcc_h8_ = source17_.cf12
            d_952_cf12_ = d_951___mcc_h8_
            d_953_v83_: _dafny.Array
            nw173_ = _dafny.Array(_dafny.Set({}), 11)
            d_953_v83_ = nw173_
            d_954_v84_: _dafny.Seq
            d_954_v84_ = _dafny.SeqWithoutIsStrInference([d_953_v83_, d_953_v83_])
            d_955_v85_: _dafny.Array
            nw174_ = _dafny.Array(None, 19)
            nw174_[int(0)] = d_953_v83_
            nw174_[int(1)] = d_953_v83_
            nw174_[int(2)] = d_953_v83_
            nw174_[int(3)] = d_953_v83_
            nw174_[int(4)] = d_953_v83_
            nw174_[int(5)] = d_953_v83_
            nw174_[int(6)] = d_953_v83_
            nw174_[int(7)] = (d_954_v84_)[default__.safeIndex((self).f39, len(d_954_v84_))]
            nw174_[int(8)] = d_953_v83_
            nw174_[int(9)] = d_953_v83_
            nw174_[int(10)] = d_953_v83_
            nw174_[int(11)] = d_953_v83_
            nw174_[int(12)] = d_953_v83_
            nw174_[int(13)] = d_953_v83_
            nw174_[int(14)] = d_953_v83_
            nw174_[int(15)] = d_953_v83_
            nw174_[int(16)] = d_953_v83_
            nw174_[int(17)] = d_953_v83_
            nw174_[int(18)] = d_953_v83_
            d_955_v85_ = nw174_
            index130_ = default__.safeIndex(28, (d_955_v85_).length(0))
            (d_955_v85_)[index130_] = d_953_v83_
            index131_ = default__.safeIndex(28, (d_955_v85_).length(0))
            (d_955_v85_)[index131_] = d_953_v83_
            (globalState).f24 = (self).f39
            (globalState).f24 = (self).f39
            d_956_v86_: D6
            d_956_v86_ = D6_DC15(((p2).set(default__.safeIndex((self).f39, len(p2)), d_879_v35_)) == (p2))
            source19_ = d_956_v86_
            if source19_.is_DC15:
                d_957___mcc_h12_ = source19_.cf20
                d_958_cf20_ = d_957___mcc_h12_
                (globalState).f24 = (0) - ((self).f39)
                d_959_v87_: _dafny.MultiSet
                d_959_v87_ = _dafny.MultiSet([p0])
                d_960_v88_: D4
                d_960_v88_ = D4_DC10((self).f40, ((d_952_cf12_)[(d_959_v87_).cardinality] if ((d_959_v87_).cardinality) in (d_952_cf12_) else (self).f39), len(d_881_v37_))
                (globalState).f24 = (d_960_v88_).cf14
                (globalState).f27 = not(p3)
                d_961_v89_: _dafny.Array
                nw175_ = _dafny.Array(int(0), 1)
                d_961_v89_ = nw175_
                d_962_v90_: _dafny.Map
                d_962_v90_ = _dafny.Map({p3: p0})
                d_963_v91_: D8
                d_963_v91_ = D8_DC21(self.f28)
                d_964_v92_: _dafny.Array
                nw176_ = _dafny.Array(None, 25)
                nw176_[int(0)] = (self).f29
                nw176_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnqlniwi"))) <= (p2)
                nw176_[int(2)] = (self).f40
                nw176_[int(3)] = d_958_cf20_
                nw176_[int(4)] = p0
                nw176_[int(5)] = p0
                nw176_[int(6)] = ((self).f39) == ((self).f39)
                nw176_[int(7)] = not((d_952_cf12_).issubset(_dafny.MultiSet([(self).f39])))
                nw176_[int(8)] = (self).f29
                nw176_[int(9)] = False
                nw176_[int(10)] = p3
                nw176_[int(11)] = (d_879_v35_) not in (p2)
                nw176_[int(12)] = d_958_cf20_
                nw176_[int(13)] = (self.f28 if (self).f40 else d_958_cf20_)
                nw176_[int(14)] = p0
                nw176_[int(15)] = ((self).f39) != ((0) - ((d_952_cf12_).cardinality))
                nw176_[int(16)] = p0
                nw176_[int(17)] = d_958_cf20_
                nw176_[int(18)] = (d_877_v33_).fm33(p2, d_880_v36_, d_952_cf12_, globalState)
                nw176_[int(19)] = ((d_962_v90_)[p0] if (p0) in (d_962_v90_) else (self).f29)
                nw176_[int(20)] = p0
                nw176_[int(21)] = p3
                nw176_[int(22)] = ((self).f39) != ((self).f39)
                nw176_[int(23)] = ((d_963_v91_).cf27) and ((self).f40)
                nw176_[int(24)] = (self).f29
                d_964_v92_ = nw176_
                rhs147_ = True
                rhs148_ = (d_961_v89_ if ((self).f29 if p0 else p0) else d_961_v89_)
                rhs149_ = d_964_v92_
                rhs150_ = (self).f39
                rhs151_ = p0
                lhs124_ = globalState
                lhs125_ = globalState
                lhs126_ = self
                lhs124_.f22 = rhs147_
                d_961_v89_ = rhs148_
                d_964_v92_ = rhs149_
                lhs125_.f1 = rhs150_
                lhs126_.f28 = rhs151_
            elif True:
                d_965___mcc_h13_ = source19_.cf19
                d_966_cf19_ = d_965___mcc_h13_
                (globalState).f22 = p3
                (globalState).f22 = not(self.f28)
                (globalState).f13 = not(self.f28)
                d_967_v93_: _dafny.Array
                nw177_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_967_v93_ = nw177_
                d_967_v93_ = d_967_v93_
        elif True:
            d_968___mcc_h9_ = source17_.cf16
            d_969_cf16_ = d_968___mcc_h9_
            if p0:
                d_970_v94_: _dafny.MultiSet
                d_970_v94_ = _dafny.MultiSet([self.f28])
                d_971_v95_: _dafny.Map
                d_971_v95_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbs")))): (self).f39})
                d_972_v96_: _dafny.Array
                nw178_ = _dafny.Array(None, 10)
                nw178_[int(0)] = (self).f39
                nw178_[int(1)] = (0) - ((219) * ((self).f39))
                nw178_[int(2)] = len(d_881_v37_)
                nw178_[int(3)] = (self).fm30(p1, globalState)
                nw178_[int(4)] = (749) * ((d_970_v94_).cardinality)
                nw178_[int(5)] = (0) - ((self).f39)
                nw178_[int(6)] = 500
                nw178_[int(7)] = (0) - ((self).f39)
                nw178_[int(8)] = (self).fm30(p1, globalState)
                nw178_[int(9)] = (0) - ((71) - ((0) - (((d_971_v95_)[(self).f39] if ((self).f39) in (d_971_v95_) else 883))))
                d_972_v96_ = nw178_
                index132_ = default__.safeIndex(270, (d_972_v96_).length(0))
                (d_972_v96_)[index132_] = (self).f39
                index133_ = default__.safeIndex(270, (d_972_v96_).length(0))
                (d_972_v96_)[index133_] = (self).f39
                index134_ = default__.safeIndex(270, (d_972_v96_).length(0))
                (d_972_v96_)[index134_] = (self).f39
                d_973_v97_: C0
                nw179_ = C0()
                nw179_.ctor__()
                d_973_v97_ = nw179_
                d_974_v98_: _dafny.Seq
                d_974_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldfydewvd"))
                d_974_v98_ = p2
                index135_ = default__.safeIndex(270, (d_972_v96_).length(0))
                (d_972_v96_)[index135_] = (d_972_v96_)[default__.safeIndex(270, (d_972_v96_).length(0))]
            elif True:
                d_975_v99_: D3
                d_975_v99_ = D3_DC8(p2)
                d_975_v99_ = d_975_v99_
                d_976_v100_: C1
                nw180_ = C1()
                nw180_.ctor__(self.f28, self.f28)
                d_976_v100_ = nw180_
                d_977_v101_: C0
                nw181_ = C0()
                nw181_.ctor__()
                d_977_v101_ = nw181_
                (globalState).f22 = True
                (globalState).f22 = ((self).f39) <= ((-68) + (len(_dafny.Map({p2: (self).f29}))))
            d_978_v102_: _dafny.Array
            nw182_ = _dafny.Array(False, 12)
            d_978_v102_ = nw182_
            d_979_v103_: D5
            d_979_v103_ = D5_DC12(d_978_v102_)
            d_979_v103_ = d_979_v103_
            d_980_v104_: _dafny.Map
            d_980_v104_ = _dafny.Map({(self).f39: (self).f39})
            d_981_v105_: _dafny.Seq
            d_981_v105_ = _dafny.SeqWithoutIsStrInference([d_980_v104_])
            (globalState).f22 = (default__.safeModuloInt(len(d_981_v105_), (self).f39)) > ((self).f39)
            (globalState).f24 = ((self).f39 if (self).f29 else (self).f39)
        d_982_v106_: D1
        d_982_v106_ = D1_DC2((self).f39, (self).f39, p2, p0)
        d_983_v107_: _dafny.MultiSet
        d_983_v107_ = _dafny.MultiSet([d_982_v106_, d_982_v106_])
        d_984_v108_: D7
        d_984_v108_ = D7_DC16(d_983_v107_)
        source20_ = d_984_v108_
        if source20_.is_DC17:
            d_985___mcc_h14_ = source20_.cf22
            d_986_cf22_ = d_985___mcc_h14_
            d_987_v109_: _dafny.Map
            d_987_v109_ = _dafny.Map({(self).f29: (self).f39})
            d_988_v110_: _dafny.Map
            d_988_v110_ = _dafny.Map({p2: len(d_987_v109_)})
            (globalState).f1 = ((d_988_v110_)[p2] if (p2) in (d_988_v110_) else (self).f39)
            d_987_v109_ = (d_987_v109_).set((self).f29, ((self).f39) - ((self).f39))
            d_989_v111_: _dafny.Map
            d_989_v111_ = _dafny.Map({True: (self).f29})
            default__.m0(d_989_v111_, (self).f39, default__.fm38(globalState), globalState)
            d_990_v112_: _dafny.Seq
            d_990_v112_ = _dafny.SeqWithoutIsStrInference([p3])
            (self).f28 = not ((d_986_cf22_) or ((self).f40)) or ((d_990_v112_)[default__.safeIndex(326, len(d_990_v112_))])
        elif source20_.is_DC18:
            d_991___mcc_h15_ = source20_.cf23
            d_992_cf23_ = d_991___mcc_h15_
            (globalState).f1 = (_dafny.MultiSet([(self).f39, (self).f39, (d_881_v37_)[default__.safeIndex((self).f39, len(d_881_v37_))]])).cardinality
            d_993_v113_: _dafny.Seq
            d_993_v113_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwnrxndg"))
            d_994_v114_: _dafny.Array
            nw183_ = _dafny.Array(int(0), 28)
            d_994_v114_ = nw183_
            d_995_v115_: _dafny.Seq
            d_995_v115_ = _dafny.SeqWithoutIsStrInference([d_994_v114_])
            d_996_v116_: _dafny.Map
            d_996_v116_ = _dafny.Map({self.f28: d_879_v35_})
            d_997_v117_: _dafny.Map
            d_997_v117_ = _dafny.Map({False: len(d_996_v116_)})
            d_998_v118_: _dafny.Array
            nw184_ = _dafny.Array(None, 15)
            nw184_[int(0)] = len(d_995_v115_)
            nw184_[int(1)] = (self).f39
            nw184_[int(2)] = (self).f39
            nw184_[int(3)] = (self).f39
            nw184_[int(4)] = (self).f39
            nw184_[int(5)] = (self).f39
            nw184_[int(6)] = (len(d_993_v113_)) * (179)
            nw184_[int(7)] = default__.safeDivisionInt(((d_997_v117_)[(self).f29] if ((self).f29) in (d_997_v117_) else (self).f39), (self).f39)
            nw184_[int(8)] = (self).f39
            nw184_[int(9)] = 917
            nw184_[int(10)] = (self).f39
            nw184_[int(11)] = (self).f39
            nw184_[int(12)] = (self).fm30(p1, globalState)
            nw184_[int(13)] = (self).f39
            nw184_[int(14)] = (self).fm30(p1, globalState)
            d_998_v118_ = nw184_
            index136_ = default__.safeIndex(285, (d_998_v118_).length(0))
            (d_998_v118_)[index136_] = default__.safeModuloInt((self).f39, 602)
            index137_ = default__.safeIndex(285, (d_998_v118_).length(0))
            rhs152_ = (p2) + (p2)
            rhs153_ = (self).f39
            rhs154_ = d_992_cf23_
            lhs127_ = d_998_v118_
            lhs128_ = default__.safeIndex(285, (d_998_v118_).length(0))
            d_993_v113_ = rhs152_
            lhs127_[lhs128_] = rhs153_
            d_992_cf23_ = rhs154_
            d_999_v119_: _dafny.MultiSet
            d_999_v119_ = _dafny.MultiSet([p3])
            d_1000_v120_: D9
            d_1000_v120_ = D9_DC23(d_999_v119_)
            d_1001_v121_: _dafny.Map
            d_1001_v121_ = _dafny.Map({(d_1000_v120_).cf28: (self).f29})
            d_1002_v122_: _dafny.Array
            nw185_ = _dafny.Array(None, 5)
            nw185_[int(0)] = p3
            nw185_[int(1)] = True
            nw185_[int(2)] = p3
            nw185_[int(3)] = self.f28
            nw185_[int(4)] = True
            d_1002_v122_ = nw185_
            d_1003_v123_: _dafny.Map
            d_1003_v123_ = _dafny.Map({((d_1001_v121_)[d_999_v119_] if (d_999_v119_) in (d_1001_v121_) else False): d_1002_v122_})
            d_1003_v123_ = ((d_1003_v123_) | (d_1003_v123_)) | (d_1003_v123_)
            d_1004_v124_: C0
            nw186_ = C0()
            nw186_.ctor__()
            d_1004_v124_ = nw186_
        elif source20_.is_DC19:
            d_1005___mcc_h16_ = source20_.cf24
            d_1006___mcc_h17_ = source20_.cf25
            d_1007_cf25_ = d_1006___mcc_h17_
            d_1008_cf24_ = d_1005___mcc_h16_
            (globalState).f1 = (self).f39
            if p0:
                d_1009_v125_: _dafny.Array
                def lambda17_(d_1010_p2_):
                    def lambda18_(d_1011_i6_):
                        return D3_DC8(d_1010_p2_)

                    return lambda18_

                init10_ = lambda17_(p2)
                nw187_ = _dafny.Array(None, 24)
                for i0_10_ in range(nw187_.length(0)):
                    nw187_[i0_10_] = init10_(i0_10_)
                d_1009_v125_ = nw187_
                d_1009_v125_ = d_1009_v125_
                (globalState).f1 = default__.safeModuloInt((self).f39, (self).f39)
                (globalState).f22 = (self).f40
                d_1012_v126_: _dafny.Array
                nw188_ = _dafny.Array(int(0), 12)
                d_1012_v126_ = nw188_
                d_1013_v127_: _dafny.MultiSet
                d_1013_v127_ = _dafny.MultiSet([not(not(not(d_1008_cf24_)))])
                index138_ = default__.safeIndex(786, (d_1012_v126_).length(0))
                (d_1012_v126_)[index138_] = ((d_1013_v127_)[True] if (True) in (d_1013_v127_) else (self).f39)
                d_1014_v128_: _dafny.Set
                d_1014_v128_ = _dafny.Set({p0})
                d_1015_v129_: _dafny.Seq
                d_1015_v129_ = _dafny.SeqWithoutIsStrInference([(self).f40, p3, p3])
                d_1016_v130_: _dafny.Map
                d_1016_v130_ = _dafny.Map({970: False})
                d_1017_v131_: D8
                d_1017_v131_ = D8_DC21(p3)
                d_1018_v132_: _dafny.Map
                d_1018_v132_ = _dafny.Map({p0: p0})
                d_1019_v133_: _dafny.Array
                nw189_ = _dafny.Array(None, 25)
                nw189_[int(0)] = ((self).f39) >= ((self).f39)
                nw189_[int(1)] = not((d_1008_cf24_) == ((self).f40))
                nw189_[int(2)] = d_1008_cf24_
                nw189_[int(3)] = p0
                nw189_[int(4)] = True
                nw189_[int(5)] = p3
                nw189_[int(6)] = (d_1014_v128_).ispropersubset(d_1014_v128_)
                nw189_[int(7)] = (not((d_877_v33_).fm34(not(not(not(d_1007_cf25_))), self.f28, globalState))) in (d_1015_v129_)
                nw189_[int(8)] = not(((self).f39) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrypxvc")))))
                nw189_[int(9)] = self.f28
                nw189_[int(10)] = d_1007_cf25_
                nw189_[int(11)] = p3
                nw189_[int(12)] = p0
                nw189_[int(13)] = ((self).f39) in (d_1016_v130_)
                nw189_[int(14)] = (self).f29
                nw189_[int(15)] = d_1007_cf25_
                nw189_[int(16)] = ((self).f39) == ((self).f39)
                nw189_[int(17)] = (d_1017_v131_).cf27
                nw189_[int(18)] = p0
                nw189_[int(19)] = True
                nw189_[int(20)] = (d_1014_v128_).ispropersubset(default__.fm47(globalState))
                nw189_[int(21)] = p0
                nw189_[int(22)] = (self).f29
                nw189_[int(23)] = ((d_1018_v132_)[d_1008_cf24_] if (d_1008_cf24_) in (d_1018_v132_) else self.f28)
                nw189_[int(24)] = (len(_dafny.SeqWithoutIsStrInference([(self).f39]))) in (d_880_v36_)
                d_1019_v133_ = nw189_
                index139_ = default__.safeIndex(62, (d_1019_v133_).length(0))
                (d_1019_v133_)[index139_] = not(self.f28)
                index140_ = default__.safeIndex(470, (d_1012_v126_).length(0))
                (d_1012_v126_)[index140_] = (self).f39
                index141_ = default__.safeIndex(39, (d_1012_v126_).length(0))
                (d_1012_v126_)[index141_] = ((self).f39) - ((self).f39)
                index142_ = default__.safeIndex(786, (d_1012_v126_).length(0))
                index143_ = default__.safeIndex(62, (d_1019_v133_).length(0))
                index144_ = default__.safeIndex(470, (d_1012_v126_).length(0))
                index145_ = default__.safeIndex(39, (d_1012_v126_).length(0))
                rhs155_ = default__.safeModuloInt(((self).f39) * (len(p2)), -866)
                rhs156_ = d_1008_cf24_
                rhs157_ = (self).f39
                rhs158_ = len((p2) + (p2))
                lhs129_ = d_1012_v126_
                lhs130_ = default__.safeIndex(786, (d_1012_v126_).length(0))
                lhs131_ = d_1019_v133_
                lhs132_ = default__.safeIndex(62, (d_1019_v133_).length(0))
                lhs133_ = d_1012_v126_
                lhs134_ = default__.safeIndex(470, (d_1012_v126_).length(0))
                lhs135_ = d_1012_v126_
                lhs136_ = default__.safeIndex(39, (d_1012_v126_).length(0))
                lhs129_[lhs130_] = rhs155_
                lhs131_[lhs132_] = rhs156_
                lhs133_[lhs134_] = rhs157_
                lhs135_[lhs136_] = rhs158_
                (globalState).f13 = not(not(d_1007_cf25_))
            elif True:
                d_1020_v134_: _dafny.Array
                nw190_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_1020_v134_ = nw190_
                d_1020_v134_ = (d_1020_v134_ if (-840) > ((self).f39) else d_1020_v134_)
                d_1021_v135_: _dafny.Array
                def lambda19_(d_1022_i7_):
                    return _dafny.Set({(self).f39})

                init11_ = lambda19_
                nw191_ = _dafny.Array(None, 4)
                for i0_11_ in range(nw191_.length(0)):
                    nw191_[i0_11_] = init11_(i0_11_)
                d_1021_v135_ = nw191_
                d_1021_v135_ = d_1021_v135_
                d_1023_v136_: _dafny.Map
                d_1023_v136_ = _dafny.Map({(self).f39: (self).f39})
                d_1023_v136_ = d_1023_v136_
                d_1024_v137_: _dafny.Array
                nw192_ = _dafny.Array(_dafny.Map({}), 10)
                d_1024_v137_ = nw192_
                d_1024_v137_ = d_1024_v137_
                d_1025_v138_: _dafny.Array
                def lambda20_(d_1026_i8_):
                    return default__.safeDivisionInt(d_1026_i8_, (self).f39)

                init12_ = lambda20_
                nw193_ = _dafny.Array(None, 24)
                for i0_12_ in range(nw193_.length(0)):
                    nw193_[i0_12_] = init12_(i0_12_)
                d_1025_v138_ = nw193_
                index146_ = default__.safeIndex(329, (d_1025_v138_).length(0))
                (d_1025_v138_)[index146_] = (self).f39
                index147_ = default__.safeIndex(329, (d_1025_v138_).length(0))
                (d_1025_v138_)[index147_] = (self).f39
            d_1027_v139_: _dafny.Seq
            d_1027_v139_ = _dafny.SeqWithoutIsStrInference([d_1007_cf25_])
            if (d_1027_v139_)[default__.safeIndex((195) + ((self).f39), len(d_1027_v139_))]:
                d_1028_v140_: C0
                nw194_ = C0()
                nw194_.ctor__()
                d_1028_v140_ = nw194_
                d_1029_v141_: _dafny.Map
                d_1029_v141_ = _dafny.Map({not((self).f40): (self).f39})
                d_1029_v141_ = (d_1029_v141_).set((d_877_v33_).fm34(p3, p0, globalState), len(p2))
                (globalState).f22 = (d_1008_cf24_ if p3 else self.f28)
                d_1030_v142_: _dafny.Array
                nw195_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_1030_v142_ = nw195_
                d_1031_v143_: D1
                d_1031_v143_ = D1_DC1(d_1030_v142_)
                d_1032_v144_: _dafny.Seq
                d_1032_v144_ = _dafny.SeqWithoutIsStrInference([d_1031_v143_])
                d_1033_v145_: _dafny.Array
                nw196_ = _dafny.Array(None, 5)
                nw196_[int(0)] = self.f28
                nw196_[int(1)] = p3
                nw196_[int(2)] = p0
                nw196_[int(3)] = (d_1031_v143_) in (d_1032_v144_)
                nw196_[int(4)] = True
                d_1033_v145_ = nw196_
                index148_ = default__.safeIndex(704, (d_1033_v145_).length(0))
                (d_1033_v145_)[index148_] = ((self).f39) < ((0) - (-571))
                index149_ = default__.safeIndex(704, (d_1033_v145_).length(0))
                (d_1033_v145_)[index149_] = not(False)
                (globalState).f22 = (len(p2)) == ((self).f39)
            elif True:
                d_1034_v146_: _dafny.MultiSet
                d_1034_v146_ = _dafny.MultiSet([d_1007_cf25_])
                d_1034_v146_ = d_1034_v146_
                (globalState).f1 = len(d_880_v36_)
                (globalState).f24 = (self).f39
                d_1035_v147_: C1
                nw197_ = C1()
                nw197_.ctor__(self.f28, self.f28)
                d_1035_v147_ = nw197_
                (globalState).f27 = p0
            d_1036_v148_: _dafny.Map
            d_1036_v148_ = _dafny.Map({d_1008_cf24_: p1})
            def iife87_():
                coll49_ = _dafny.Set()
                compr_49_: int
                for compr_49_ in _dafny.IntegerRange(-629, 870):
                    d_1037_v149_: int = compr_49_
                    if ((-629) <= (d_1037_v149_)) and ((d_1037_v149_) < (870)):
                        coll49_ = coll49_.union(_dafny.Set([default__.safeModuloInt(d_1037_v149_, (self).f39)]))
                return _dafny.Set(coll49_)
            (globalState).f27 = ((iife87_()
            ) | (p1)).ispropersubset(((d_1036_v148_)[(self).f29] if ((self).f29) in (d_1036_v148_) else p1))
        elif True:
            d_1038___mcc_h18_ = source20_.cf21
            d_1039_cf21_ = d_1038___mcc_h18_
            d_1040_v150_: T0
            nw198_ = C1()
            nw198_.ctor__(self.f28, p0)
            d_1040_v150_ = nw198_
            d_1041_v151_: D8
            d_1041_v151_ = D8_DC20(d_1040_v150_)
            d_1042_v152_: D9
            d_1042_v152_ = D9_DC24(d_879_v35_, not(not(self.f28)), (self).f39, d_1041_v151_, (self).f29)
            (globalState).f24 = (d_1042_v152_).cf31
            d_1043_v153_: _dafny.MultiSet
            d_1043_v153_ = _dafny.MultiSet([d_1040_v150_.f28])
            d_1044_v154_: _dafny.Seq
            d_1044_v154_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f39: p3})])
            rhs159_ = (((self).f39) - (len(p2))) - ((550) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktmavtb")))))
            rhs160_ = ((d_879_v35_ if self.f28 else d_879_v35_) if (default__.fm48(d_1043_v153_, (self).f40, (self).f39, globalState)) <= (d_1044_v154_) else _dafny.CodePoint('k'))
            lhs137_ = globalState
            lhs138_ = globalState
            lhs137_.f24 = rhs159_
            lhs138_.f23 = rhs160_
            (globalState).f22 = ((self).f39) > ((self).f39)
            d_1045_v155_: _dafny.Map
            d_1045_v155_ = _dafny.Map({d_879_v35_: (self).f29})
            d_1046_v156_: D3
            d_1046_v156_ = D3_DC8(p2)
            d_1047_v157_: _dafny.Map
            d_1047_v157_ = _dafny.Map({d_1046_v156_: (self).f39})
            d_1045_v155_ = (d_1045_v155_).set(default__.fm42((self).f39, ((d_1047_v157_)[d_1046_v156_] if (d_1046_v156_) in (d_1047_v157_) else (self).f39), globalState), (self).f40)
        (globalState).f24 = (self).f39

    def m5(self, p0, globalState):
        r0: bool = False
        d_1048_v0_: str
        d_1048_v0_ = _dafny.CodePoint('e')
        d_1049_v1_: _dafny.Set
        d_1049_v1_ = _dafny.Set({self.f28, self.f28})
        d_1050_v2_: _dafny.Array
        nw199_ = _dafny.Array(None, 28)
        nw199_[int(0)] = d_1048_v0_
        nw199_[int(1)] = d_1048_v0_
        nw199_[int(2)] = d_1048_v0_
        nw199_[int(3)] = d_1048_v0_
        nw199_[int(4)] = d_1048_v0_
        nw199_[int(5)] = d_1048_v0_
        nw199_[int(6)] = d_1048_v0_
        nw199_[int(7)] = d_1048_v0_
        nw199_[int(8)] = d_1048_v0_
        nw199_[int(9)] = d_1048_v0_
        nw199_[int(10)] = _dafny.CodePoint('m')
        nw199_[int(11)] = default__.fm42((self).f39, p0, globalState)
        nw199_[int(12)] = d_1048_v0_
        nw199_[int(13)] = d_1048_v0_
        nw199_[int(14)] = d_1048_v0_
        nw199_[int(15)] = d_1048_v0_
        nw199_[int(16)] = d_1048_v0_
        nw199_[int(17)] = (d_1048_v0_ if (self).fm32(d_1049_v1_, p0, (self).f29, globalState) else d_1048_v0_)
        nw199_[int(18)] = d_1048_v0_
        nw199_[int(19)] = d_1048_v0_
        nw199_[int(20)] = _dafny.CodePoint('v')
        nw199_[int(21)] = _dafny.CodePoint('u')
        nw199_[int(22)] = d_1048_v0_
        nw199_[int(23)] = d_1048_v0_
        nw199_[int(24)] = d_1048_v0_
        nw199_[int(25)] = d_1048_v0_
        nw199_[int(26)] = d_1048_v0_
        nw199_[int(27)] = d_1048_v0_
        d_1050_v2_ = nw199_
        d_1050_v2_ = d_1050_v2_
        if (self).f40:
            d_1051_v3_: D7
            d_1051_v3_ = D7_DC19(self.f28, self.f28)
            d_1052_v4_: _dafny.Set
            d_1052_v4_ = _dafny.Set({d_1051_v3_})
            rhs161_ = p0
            rhs162_ = (default__.fm49(globalState)) not in (d_1052_v4_)
            rhs163_ = False
            lhs139_ = globalState
            lhs140_ = globalState
            lhs141_ = globalState
            lhs139_.f24 = rhs161_
            lhs140_.f21 = rhs162_
            lhs141_.f22 = rhs163_
            d_1053_v5_: _dafny.Seq
            d_1053_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhiro"))
            (globalState).f22 = not((d_1053_v5_) <= ((d_1053_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dd")))))
            d_1054_v6_: _dafny.Set
            d_1054_v6_ = _dafny.Set({p0})
            (globalState).f24 = len(((d_1054_v6_) - (d_1054_v6_)) - ((d_1054_v6_).intersection(_dafny.Set({p0}))))
            d_1055_v7_: _dafny.Array
            def lambda21_(d_1056_i0_):
                return (_dafny.SeqWithoutIsStrInference([(D9_DC25((self).f29, (0) - ((self).f39), len(_dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])), False, (self).f39)).cf34]) if self.f28 else _dafny.SeqWithoutIsStrInference([(self).f29, (self).f40]))

            init13_ = lambda21_
            nw200_ = _dafny.Array(None, 20)
            for i0_13_ in range(nw200_.length(0)):
                nw200_[i0_13_] = init13_(i0_13_)
            d_1055_v7_ = nw200_
            index150_ = default__.safeIndex(332, (d_1055_v7_).length(0))
            (d_1055_v7_)[index150_] = _dafny.SeqWithoutIsStrInference([not(not(not((self).f40))), self.f28])
            d_1057_v8_: _dafny.Seq
            d_1057_v8_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            index151_ = default__.safeIndex(332, (d_1055_v7_).length(0))
            rhs164_ = ((_dafny.SeqWithoutIsStrInference([(self).f29, True, self.f28])) + (_dafny.SeqWithoutIsStrInference([(self).f29, (self).f29, (self).f29]))) + (d_1057_v8_)
            rhs165_ = (self).fm30(_dafny.Set({p0}), globalState)
            lhs142_ = d_1055_v7_
            lhs143_ = default__.safeIndex(332, (d_1055_v7_).length(0))
            lhs144_ = globalState
            lhs142_[lhs143_] = rhs164_
            lhs144_.f24 = rhs165_
            (globalState).f1 = p0
        elif True:
            d_1058_v9_: _dafny.Seq
            d_1058_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgq"))
            (globalState).f24 = default__.safeModuloInt(((self).f39) + ((self).f39), len((d_1058_v9_) + (d_1058_v9_)))
            (globalState).f21 = not((self).f40)
            d_1059_v10_: _dafny.Array
            nw201_ = _dafny.Array(int(0), 3)
            d_1059_v10_ = nw201_
            d_1060_v11_: _dafny.MultiSet
            d_1060_v11_ = _dafny.MultiSet([d_1059_v10_])
            d_1061_v12_: _dafny.MultiSet
            d_1061_v12_ = _dafny.MultiSet([(self).f40])
            d_1062_v13_: _dafny.Set
            d_1062_v13_ = _dafny.Set({(self).f39, (self).f39})
            if ((d_1060_v11_) == (d_1060_v11_)) in ((d_1061_v12_).set((self).f40, default__.abs((self).fm30(d_1062_v13_, globalState)))):
                d_1063_v14_: _dafny.Seq
                d_1063_v14_ = _dafny.SeqWithoutIsStrInference([d_1059_v10_])
                d_1064_v15_: _dafny.Map
                d_1064_v15_ = _dafny.Map({(d_1063_v14_)[default__.safeIndex(p0, len(d_1063_v14_))]: ((self).f29) and (not(self.f28))})
                r0 = not(((d_1064_v15_)[d_1059_v10_] if (d_1059_v10_) in (d_1064_v15_) else self.f28))
                (globalState).f21 = (self).f40
                d_1065_v16_: _dafny.MultiSet
                d_1065_v16_ = _dafny.MultiSet([p0])
                d_1066_v17_: _dafny.Array
                nw202_ = _dafny.Array(None, 19)
                nw202_[int(0)] = (self).f40
                nw202_[int(1)] = ((self).fm30(default__.fm35(globalState), globalState)) < ((self).f39)
                nw202_[int(2)] = self.f28
                nw202_[int(3)] = True
                nw202_[int(4)] = self.f28
                nw202_[int(5)] = ((d_1065_v16_).cardinality) == (((d_1061_v12_).set((self).f29, default__.abs((self).f39))).cardinality)
                nw202_[int(6)] = self.f28
                nw202_[int(7)] = ((d_1058_v9_)[default__.safeIndex((self).f39, len(d_1058_v9_))]) in (d_1058_v9_)
                nw202_[int(8)] = (self).f40
                nw202_[int(9)] = (self).fm31(globalState)
                nw202_[int(10)] = (d_1048_v0_) not in (d_1058_v9_)
                nw202_[int(11)] = (self).f40
                nw202_[int(12)] = (self).f40
                nw202_[int(13)] = not((self).f29)
                nw202_[int(14)] = (self).f40
                nw202_[int(15)] = (self).f40
                nw202_[int(16)] = (self).f40
                nw202_[int(17)] = (self).f29
                nw202_[int(18)] = self.f28
                d_1066_v17_ = nw202_
                d_1067_v18_: _dafny.Seq
                d_1067_v18_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1068_v19_: _dafny.Seq
                d_1068_v19_ = _dafny.SeqWithoutIsStrInference([self.f28, (self).f40])
                nw203_ = _dafny.Array(None, 27)
                nw203_[int(0)] = (self).f40
                nw203_[int(1)] = self.f28
                nw203_[int(2)] = (self).f29
                nw203_[int(3)] = (d_1067_v18_) == (_dafny.SeqWithoutIsStrInference([(self).f39 for d_1069_i1_ in range(default__.abs(732))]))
                nw203_[int(4)] = ((d_1068_v19_)[default__.safeIndex(p0, len(d_1068_v19_))]) or ((self).f40)
                nw203_[int(5)] = (d_1068_v19_)[default__.safeIndex((self).f39, len(d_1068_v19_))]
                nw203_[int(6)] = self.f28
                nw203_[int(7)] = (self).f29
                nw203_[int(8)] = (d_1061_v12_).isdisjoint(_dafny.MultiSet(d_1068_v19_))
                nw203_[int(9)] = self.f28
                nw203_[int(10)] = self.f28
                nw203_[int(11)] = (p0) >= ((self).f39)
                nw203_[int(12)] = (self).f40
                nw203_[int(13)] = ((self).f39) == ((self).f39)
                nw203_[int(14)] = ((self).f40) == (False)
                nw203_[int(15)] = True
                nw203_[int(16)] = self.f28
                nw203_[int(17)] = self.f28
                nw203_[int(18)] = (self).f40
                nw203_[int(19)] = True
                nw203_[int(20)] = True
                nw203_[int(21)] = (self.f28) not in (d_1049_v1_)
                nw203_[int(22)] = not((self).f29)
                nw203_[int(23)] = (d_1049_v1_).ispropersubset(d_1049_v1_)
                nw203_[int(24)] = (d_1058_v9_) <= (d_1058_v9_)
                nw203_[int(25)] = (len(d_1068_v19_)) <= (p0)
                nw203_[int(26)] = (self).f40
                d_1066_v17_ = nw203_
                (globalState).f24 = default__.safeDivisionInt((self).f39, (0) - ((0) - ((len(_dafny.Map({not((self).f40): True}))) * (p0))))
                d_1070_v20_: C1
                nw204_ = C1()
                nw204_.ctor__((self).f40, not(False))
                d_1070_v20_ = nw204_
            elif True:
                d_1071_v21_: _dafny.Array
                d_1071_v21_ = d_1059_v10_
                rhs166_ = (d_1048_v0_) not in (d_1058_v9_)
                rhs167_ = d_1071_v21_
                rhs168_ = p0
                lhs145_ = globalState
                lhs146_ = globalState
                lhs145_.f13 = rhs166_
                d_1071_v21_ = rhs167_
                lhs146_.f1 = rhs168_
                d_1072_v22_: _dafny.MultiSet
                d_1072_v22_ = _dafny.MultiSet([p0, p0])
                index152_ = default__.safeIndex(936, (d_1050_v2_).length(0))
                (d_1050_v2_)[index152_] = d_1048_v0_
                index153_ = default__.safeIndex(936, (d_1050_v2_).length(0))
                rhs169_ = (d_1049_v1_).ispropersubset((d_1049_v1_) | (d_1049_v1_))
                rhs170_ = d_1050_v2_
                rhs171_ = d_1072_v22_
                rhs172_ = (p0) - (p0)
                rhs173_ = d_1048_v0_
                lhs147_ = globalState
                lhs148_ = globalState
                lhs149_ = d_1050_v2_
                lhs150_ = default__.safeIndex(936, (d_1050_v2_).length(0))
                lhs147_.f27 = rhs169_
                d_1050_v2_ = rhs170_
                d_1072_v22_ = rhs171_
                lhs148_.f24 = rhs172_
                lhs149_[lhs150_] = rhs173_
                d_1073_v23_: _dafny.Map
                d_1073_v23_ = _dafny.Map({(self).f40: p0})
                (globalState).f1 = ((d_1073_v23_)[(self).f29] if ((self).f29) in (d_1073_v23_) else len(_dafny.SeqWithoutIsStrInference([self.f28])))
                (globalState).f13 = ((self).f39) <= ((self).f39)
                d_1074_v24_: _dafny.Seq
                d_1074_v24_ = _dafny.SeqWithoutIsStrInference([self.f28])
                d_1075_v25_: _dafny.Map
                d_1075_v25_ = _dafny.Map({len(d_1049_v1_): (self).f39})
                d_1076_v26_: _dafny.Seq
                d_1076_v26_ = _dafny.SeqWithoutIsStrInference([len(d_1075_v25_), (self).f39])
                def iife88_():
                    coll50_ = _dafny.Set()
                    compr_50_: int
                    for compr_50_ in (d_1076_v26_).Elements:
                        d_1077_v27_: int = compr_50_
                        if (d_1077_v27_) in (d_1076_v26_):
                            coll50_ = coll50_.union(_dafny.Set([(d_1077_v27_) * (885)]))
                    return _dafny.Set(coll50_)
                (globalState).f24 = default__.safeDivisionInt(len(d_1074_v24_), ((self).f39) * ((self).fm30(iife88_()
                , globalState)))
            d_1078_v28_: _dafny.Array
            nw205_ = _dafny.Array(False, 10)
            d_1078_v28_ = nw205_
            d_1079_v29_: _dafny.Seq
            d_1079_v29_ = _dafny.SeqWithoutIsStrInference([d_1078_v28_])
            (globalState).f21 = not((_dafny.SeqWithoutIsStrInference([d_1078_v28_])) == ((_dafny.SeqWithoutIsStrInference([d_1078_v28_, d_1078_v28_, d_1078_v28_, d_1078_v28_, d_1078_v28_])) + (d_1079_v29_)))
            d_1080_v30_: D9
            d_1080_v30_ = D9_DC23((d_1061_v12_) | (d_1061_v12_))
            pat_let_tv29_ = d_1061_v12_
            pat_let_tv30_ = p0
            pat_let_tv31_ = d_1061_v12_
            pat_let_tv32_ = d_1061_v12_
            def iife90_(_pat_let20_0):
                def iife91_(d_1081_dt__update__tmp_h0_):
                    def iife92_(_pat_let21_0):
                        def iife93_(d_1082_dt__update_hcf28_h0_):
                            return D9_DC23(d_1082_dt__update_hcf28_h0_)
                        return iife93_(_pat_let21_0)
                    return iife92_((pat_let_tv29_).set((self).f29, default__.abs(pat_let_tv30_)))
                return iife91_(_pat_let20_0)
            def iife89_(_pat_let19_0):
                def iife94_(d_1083_dt__update__tmp_h1_):
                    def iife95_(_pat_let22_0):
                        def iife96_(d_1084_dt__update_hcf28_h1_):
                            return D9_DC23(d_1084_dt__update_hcf28_h1_)
                        return iife96_(_pat_let22_0)
                    return iife95_((pat_let_tv31_) - (pat_let_tv32_))
                return iife94_(_pat_let19_0)
            d_1080_v30_ = iife89_(iife90_(d_1080_v30_))
        d_1085_v31_: _dafny.Seq
        d_1085_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guwyok"))
        d_1085_v31_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgrqj"))) + (d_1085_v31_)
        d_1086_v32_: _dafny.Map
        d_1086_v32_ = _dafny.Map({p0: p0})
        d_1086_v32_ = (d_1086_v32_).set((p0 if self.f28 else (self).f39), (self).fm30(_dafny.Set({(self).f39}), globalState))
        d_1087_v33_: D4
        d_1087_v33_ = D4_DC10(self.f28, (self).f39, p0)
        d_1088_v34_: _dafny.Map
        d_1088_v34_ = _dafny.Map({d_1087_v33_: (p0) * ((self).f39)})
        d_1089_v35_: _dafny.Set
        d_1089_v35_ = _dafny.Set({(0) - ((self).f39)})
        d_1088_v34_ = (d_1088_v34_).set(d_1087_v33_, (self).fm30(d_1089_v35_, globalState))
        d_1090_v36_: _dafny.Seq
        d_1090_v36_ = _dafny.SeqWithoutIsStrInference([p0])
        (globalState).f27 = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsm")) if self.f28 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbyelp"))))) >= (default__.safeDivisionInt(len(d_1090_v36_), p0))
        r0 = (self).f40
        return r0

    def m15(self, p0, globalState):
        source21_ = D8_DC21(True)
        if source21_.is_DC21:
            d_1091___mcc_h0_ = source21_.cf27
            d_1092_cf27_ = d_1091___mcc_h0_
            d_1093_v0_: C1
            nw206_ = C1()
            nw206_.ctor__((not((self).f40) if (self).f29 else True), self.f28)
            d_1093_v0_ = nw206_
            d_1094_v1_: C0
            nw207_ = C0()
            nw207_.ctor__()
            d_1094_v1_ = nw207_
            d_1095_v2_: D2
            d_1095_v2_ = D2_DC6(D2_DC5((self).f40))
            d_1096_v3_: _dafny.Map
            d_1096_v3_ = _dafny.Map({(self).f29: d_1095_v2_})
            d_1097_v4_: _dafny.Seq
            d_1097_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1096_v3_), (self).f39, (self).f39])
            (globalState).f21 = ((d_1097_v4_) + (d_1097_v4_)) != (d_1097_v4_)
            d_1098_v5_: _dafny.Seq
            d_1098_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_1099_v6_: _dafny.MultiSet
            d_1099_v6_ = _dafny.MultiSet([D1_DC2((self).f39, -673, d_1098_v5_, True)])
            d_1100_v7_: D7
            d_1100_v7_ = D7_DC16((d_1099_v6_).intersection(d_1099_v6_))
            pat_let_tv33_ = globalState
            def iife97_(_pat_let23_0):
                def iife98_(d_1101_dt__update__tmp_h0_):
                    def iife99_(_pat_let24_0):
                        def iife100_(d_1102_dt__update_hcf21_h0_):
                            return D7_DC16(d_1102_dt__update_hcf21_h0_)
                        return iife100_(_pat_let24_0)
                    return iife99_(default__.fm50(False, False, (self).f39, (self).f39, pat_let_tv33_))
                return iife98_(_pat_let23_0)
            d_1100_v7_ = (d_1100_v7_ if True else iife97_(d_1100_v7_))
        elif source21_.is_DC22:
            d_1103_v8_: _dafny.Seq
            d_1103_v8_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1104_v9_: _dafny.Seq
            d_1104_v9_ = _dafny.SeqWithoutIsStrInference([(self).f39, (self).f39])
            rhs174_ = (self).f29
            rhs175_ = (self).f29
            rhs176_ = (self).f29
            rhs177_ = (self).f39
            rhs178_ = ((_dafny.SeqWithoutIsStrInference([(self).f39, len(d_1103_v8_)])).set(default__.safeIndex((self).f39, len(_dafny.SeqWithoutIsStrInference([(self).f39, len(d_1103_v8_)]))), (self).f39)) + (d_1104_v9_)
            lhs151_ = globalState
            lhs152_ = globalState
            lhs153_ = globalState
            lhs154_ = globalState
            lhs155_ = globalState
            lhs151_.f22 = rhs174_
            lhs152_.f13 = rhs175_
            lhs153_.f22 = rhs176_
            lhs154_.f24 = rhs177_
            lhs155_.f17 = rhs178_
            (globalState).f1 = (self).f39
            (globalState).f1 = (0) - (default__.safeModuloInt(((self).f39) + (848), ((self).f39) * (919)))
            (globalState).f27 = not ((self).f29) or (p0)
        elif True:
            d_1105___mcc_h1_ = source21_.cf26
            d_1106_cf26_ = d_1105___mcc_h1_
            d_1107_v10_: _dafny.Array
            def lambda22_(d_1108_i0_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdy"))

            init14_ = lambda22_
            nw208_ = _dafny.Array(None, 11)
            for i0_14_ in range(nw208_.length(0)):
                nw208_[i0_14_] = init14_(i0_14_)
            d_1107_v10_ = nw208_
            d_1109_v11_: _dafny.Seq
            d_1109_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avfada"))
            index154_ = default__.safeIndex(315, (d_1107_v10_).length(0))
            (d_1107_v10_)[index154_] = (d_1109_v11_) + (d_1109_v11_)
            d_1110_v12_: _dafny.MultiSet
            d_1110_v12_ = _dafny.MultiSet([default__.fm51((self).f39, (self).f39, globalState)])
            d_1111_v13_: _dafny.Seq
            d_1111_v13_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1112_v14_: _dafny.MultiSet
            d_1112_v14_ = _dafny.MultiSet([(self).f39, (self).f39])
            d_1113_v15_: _dafny.Seq
            d_1113_v15_ = _dafny.SeqWithoutIsStrInference([(self).f39, (self).f39])
            d_1114_v16_: _dafny.Seq
            d_1114_v16_ = _dafny.SeqWithoutIsStrInference([(self).f39, len((_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([not((self).f40), d_1106_cf26_.f28])).cardinality) for d_1115_i1_ in range(default__.abs(-170))]) if self.f28 else d_1113_v15_)), default__.safeDivisionInt((self).f39, (self).f39)])
            index155_ = default__.safeIndex(315, (d_1107_v10_).length(0))
            rhs179_ = (d_1109_v11_) + (d_1109_v11_)
            rhs180_ = (d_1111_v13_)[default__.safeIndex(((d_1112_v14_)[len(d_1111_v13_)] if (len(d_1111_v13_)) in (d_1112_v14_) else len(d_1109_v11_)), len(d_1111_v13_))]
            rhs181_ = default__.safeDivisionInt((self).f39, len(d_1111_v13_))
            rhs182_ = (d_1113_v15_)[default__.safeIndex(((self).f39 if not(not(p0)) else len(d_1113_v15_)), len(d_1113_v15_))]
            rhs183_ = _dafny.MultiSet([d_1109_v11_])
            lhs156_ = d_1107_v10_
            lhs157_ = default__.safeIndex(315, (d_1107_v10_).length(0))
            lhs158_ = d_1106_cf26_
            lhs159_ = globalState
            lhs160_ = globalState
            lhs156_[lhs157_] = rhs179_
            lhs158_.f28 = rhs180_
            lhs159_.f24 = rhs181_
            lhs160_.f24 = rhs182_
            d_1110_v12_ = rhs183_
            (globalState).f13 = (self).f29
            d_1116_v17_: str
            d_1116_v17_ = _dafny.CodePoint('i')
            d_1117_v18_: _dafny.Set
            d_1117_v18_ = _dafny.Set({d_1116_v17_})
            d_1118_v19_: _dafny.Set
            d_1118_v19_ = _dafny.Set({(self).f39, len(d_1117_v18_)})
            (globalState).f22 = (d_1118_v19_).ispropersubset(d_1118_v19_)
            d_1119_v20_: T1
            nw209_ = C2()
            nw209_.ctor__()
            d_1119_v20_ = nw209_
            d_1120_v21_: _dafny.Map
            d_1120_v21_ = _dafny.Map({(self).f39: d_1119_v20_})
            d_1120_v21_ = (d_1120_v21_).set((self).f39, d_1119_v20_)
        d_1121_v22_: _dafny.Array
        def lambda23_(d_1122_i2_):
            return (d_1122_i2_) * ((self).f39)

        init15_ = lambda23_
        nw210_ = _dafny.Array(None, 17)
        for i0_15_ in range(nw210_.length(0)):
            nw210_[i0_15_] = init15_(i0_15_)
        d_1121_v22_ = nw210_
        index156_ = default__.safeIndex(684, (d_1121_v22_).length(0))
        (d_1121_v22_)[index156_] = (self).f39
        d_1123_v23_: _dafny.Set
        d_1123_v23_ = _dafny.Set({(self).f39})
        d_1124_v24_: D12
        d_1124_v24_ = D12_DC31(d_1123_v23_)
        index157_ = default__.safeIndex(684, (d_1121_v22_).length(0))
        (d_1121_v22_)[index157_] = (self).fm30((d_1124_v24_).cf48, globalState)
        d_1125_v25_: _dafny.Array
        nw211_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_1125_v25_ = nw211_
        d_1126_v26_: _dafny.Seq
        d_1126_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axamid"))
        index158_ = default__.safeIndex(51, (d_1125_v25_).length(0))
        (d_1125_v25_)[index158_] = d_1126_v26_
        index159_ = default__.safeIndex(51, (d_1125_v25_).length(0))
        (d_1125_v25_)[index159_] = d_1126_v26_
        d_1127_v27_: _dafny.Array
        d_1127_v27_ = d_1121_v22_
        d_1128_v28_: _dafny.Set
        d_1128_v28_ = _dafny.Set({(d_1127_v27_), d_1121_v22_})
        d_1128_v28_ = d_1128_v28_
        d_1129_v29_: D7
        d_1129_v29_ = D7_DC19((self).f29, True)
        source22_ = d_1129_v29_
        if source22_.is_DC17:
            d_1130___mcc_h2_ = source22_.cf22
            d_1131_cf22_ = d_1130___mcc_h2_
            d_1131_cf22_ = (self).f40
            d_1132_v30_: _dafny.Seq
            d_1132_v30_ = _dafny.SeqWithoutIsStrInference([(self).f39, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (self).f39, (self).f39])
            (globalState).f17 = (d_1132_v30_ if (self).f29 else d_1132_v30_)
            d_1133_v31_: _dafny.Map
            d_1133_v31_ = _dafny.Map({(d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]: p0})
            d_1134_v33_: _dafny.Seq
            d_1134_v33_ = _dafny.SeqWithoutIsStrInference([d_1131_cf22_])
            d_1135_v34_: _dafny.Seq
            d_1135_v34_ = _dafny.SeqWithoutIsStrInference([d_1134_v33_])
            d_1136_v37_: _dafny.Array
            nw212_ = _dafny.Array(None, 21)
            nw212_[int(0)] = d_1133_v31_
            nw212_[int(1)] = d_1133_v31_
            nw212_[int(2)] = (d_1133_v31_) | (d_1133_v31_)
            nw212_[int(3)] = d_1133_v31_
            nw212_[int(4)] = _dafny.Map({(self).f39: self.f28})
            nw212_[int(5)] = d_1133_v31_
            nw212_[int(6)] = _dafny.Map({(self).f39: False})
            nw212_[int(7)] = d_1133_v31_
            def iife101_():
                coll51_ = _dafny.Map()
                compr_51_: int
                for compr_51_ in _dafny.IntegerRange(654, 570):
                    d_1137_v32_: int = compr_51_
                    if ((654) <= (d_1137_v32_)) and ((d_1137_v32_) < (570)):
                        coll51_[default__.safeDivisionInt(d_1137_v32_, (self).f39)] = False
                return _dafny.Map(coll51_)
            nw212_[int(8)] = iife101_()
            
            nw212_[int(9)] = _dafny.Map({(0) - ((self).f39): self.f28})
            nw212_[int(10)] = (d_1133_v31_) | (d_1133_v31_)
            nw212_[int(11)] = d_1133_v31_
            nw212_[int(12)] = d_1133_v31_
            nw212_[int(13)] = (d_1133_v31_).set(len(d_1135_v34_), self.f28)
            nw212_[int(14)] = d_1133_v31_
            nw212_[int(15)] = _dafny.Map({(self).f39: (self).f29})
            nw212_[int(16)] = d_1133_v31_
            nw212_[int(17)] = (d_1133_v31_).set((self).f39, True)
            def iife102_():
                coll52_ = _dafny.Map()
                compr_52_: int
                for compr_52_ in _dafny.IntegerRange(599, 425):
                    d_1138_v35_: int = compr_52_
                    if ((599) <= (d_1138_v35_)) and ((d_1138_v35_) < (425)):
                        coll52_[default__.safeModuloInt(d_1138_v35_, len(_dafny.Map({d_1131_cf22_: (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]})))] = False
                return _dafny.Map(coll52_)
            nw212_[int(18)] = iife102_()
            
            def iife103_():
                coll53_ = _dafny.Map()
                compr_53_: int
                for compr_53_ in (_dafny.SeqWithoutIsStrInference([(0) - ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]) for d_1139_i3_ in range(default__.abs(202))])).Elements:
                    d_1140_v36_: int = compr_53_
                    if (d_1140_v36_) in (_dafny.SeqWithoutIsStrInference([(0) - ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]) for d_1139_i3_ in range(default__.abs(202))])):
                        coll53_[default__.safeDivisionInt(d_1140_v36_, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])] = (self).f40
                return _dafny.Map(coll53_)
            nw212_[int(19)] = ((iife103_()
            ).set((self).f39, False)) | ((d_1133_v31_).set(len(d_1126_v26_), (self).f40))
            nw212_[int(20)] = (_dafny.Map({len((d_1125_v25_)[default__.safeIndex(51, (d_1125_v25_).length(0))]): True})) | (d_1133_v31_)
            d_1136_v37_ = nw212_
            index160_ = default__.safeIndex(942, (d_1136_v37_).length(0))
            (d_1136_v37_)[index160_] = d_1133_v31_
            index161_ = default__.safeIndex(942, (d_1136_v37_).length(0))
            (d_1136_v37_)[index161_] = d_1133_v31_
            d_1141_v38_: str
            d_1141_v38_ = _dafny.CodePoint('i')
            d_1142_v39_: T0
            nw213_ = C1()
            nw213_.ctor__(d_1131_cf22_, (self).f29)
            d_1142_v39_ = nw213_
            d_1143_v40_: _dafny.Map
            d_1143_v40_ = _dafny.Map({(self).f40: d_1142_v39_})
            d_1144_v41_: D8
            d_1144_v41_ = D8_DC20(((d_1143_v40_)[(self).f40] if ((self).f40) in (d_1143_v40_) else d_1142_v39_))
            d_1145_v42_: D9
            d_1145_v42_ = D9_DC24(d_1141_v38_, (self).f29, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], d_1144_v41_, p0)
            d_1146_v43_: _dafny.Map
            d_1146_v43_ = _dafny.Map({p0: (d_1145_v42_).cf31})
            d_1146_v43_ = ((d_1146_v43_) | (_dafny.Map({(self).f29: len(d_1134_v33_)}))) | (d_1146_v43_)
        elif source22_.is_DC18:
            d_1147___mcc_h3_ = source22_.cf23
            d_1148_cf23_ = d_1147___mcc_h3_
            (globalState).f21 = (False) or (self.f28)
            index162_ = default__.safeIndex(684, (d_1121_v22_).length(0))
            (d_1121_v22_)[index162_] = (self).f39
            index163_ = default__.safeIndex(684, (d_1121_v22_).length(0))
            rhs184_ = -361
            rhs185_ = self.f28
            lhs161_ = d_1121_v22_
            lhs162_ = default__.safeIndex(684, (d_1121_v22_).length(0))
            lhs161_[lhs162_] = rhs184_
            d_1148_cf23_ = rhs185_
            d_1149_v44_: _dafny.Array
            nw214_ = _dafny.Array(False, 4)
            d_1149_v44_ = nw214_
            index164_ = default__.safeIndex(846, (d_1149_v44_).length(0))
            (d_1149_v44_)[index164_] = d_1148_cf23_
            index165_ = default__.safeIndex(846, (d_1149_v44_).length(0))
            (d_1149_v44_)[index165_] = (self).f40
        elif source22_.is_DC19:
            d_1150___mcc_h4_ = source22_.cf24
            d_1151___mcc_h5_ = source22_.cf25
            d_1152_cf25_ = d_1151___mcc_h5_
            d_1153_cf24_ = d_1150___mcc_h4_
            source23_ = D6_DC15(d_1153_cf24_)
            if source23_.is_DC15:
                d_1154___mcc_h7_ = source23_.cf20
                d_1155_cf20_ = d_1154___mcc_h7_
                d_1156_v45_: _dafny.Map
                d_1156_v45_ = _dafny.Map({(self).f29: (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]})
                d_1157_v46_: _dafny.Map
                d_1157_v46_ = _dafny.Map({-648: (self).f29})
                d_1158_v47_: _dafny.Map
                d_1158_v47_ = _dafny.Map({((d_1156_v45_)[(self).f29] if ((self).f29) in (d_1156_v45_) else (self).f39): ((d_1157_v46_)[(self).f39] if ((self).f39) in (d_1157_v46_) else d_1152_cf25_)})
                d_1159_v48_: _dafny.Seq
                d_1159_v48_ = _dafny.SeqWithoutIsStrInference([(d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]])
                (globalState).f21 = ((len(d_1157_v46_) if not((self).f29) else (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])) in (d_1159_v48_)
                d_1160_v49_: C0
                nw215_ = C0()
                nw215_.ctor__()
                d_1160_v49_ = nw215_
                d_1161_v50_: _dafny.Map
                d_1161_v50_ = _dafny.Map({p0: d_1121_v22_})
                d_1162_v51_: _dafny.Map
                d_1162_v51_ = _dafny.Map({d_1161_v50_: (self).f40})
                d_1163_v52_: _dafny.MultiSet
                d_1163_v52_ = _dafny.MultiSet([(self).f39, (self).f39])
                d_1162_v51_ = (d_1162_v51_).set((d_1161_v50_) | (d_1161_v50_), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((d_1125_v25_)[default__.safeIndex(51, (d_1125_v25_).length(0))])]))).ispropersubset(d_1163_v52_))
                (globalState).f1 = ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]) * (len((d_1126_v26_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1164_i4_ in range(default__.abs(-196))]))))
            elif True:
                d_1165___mcc_h8_ = source23_.cf19
                d_1166_cf19_ = d_1165___mcc_h8_
                d_1121_v22_ = (d_1121_v22_ if p0 else d_1121_v22_)
                d_1167_v53_: _dafny.Set
                d_1167_v53_ = _dafny.Set({True})
                d_1168_v54_: _dafny.MultiSet
                d_1168_v54_ = _dafny.MultiSet([d_1121_v22_])
                (globalState).f21 = (self).fm32(d_1167_v53_, 528, (d_1168_v54_).issubset(d_1168_v54_), globalState)
                d_1169_v55_: _dafny.Seq
                d_1169_v55_ = _dafny.SeqWithoutIsStrInference([(self).f39])
                d_1170_v56_: _dafny.Map
                d_1170_v56_ = _dafny.Map({(d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]: len(d_1169_v55_)})
                d_1171_v57_: _dafny.Map
                d_1171_v57_ = _dafny.Map({not((self).f29): p0})
                d_1172_v59_: str
                d_1172_v59_ = _dafny.CodePoint('e')
                d_1173_v60_: D2
                d_1173_v60_ = D2_DC4(d_1172_v59_)
                d_1174_v61_: D2
                d_1174_v61_ = D2_DC6(d_1173_v60_)
                d_1175_v62_: D2
                d_1175_v62_ = D2_DC6(d_1174_v61_)
                d_1176_v63_: D2
                d_1176_v63_ = D2_DC6(d_1174_v61_)
                d_1177_v64_: D2
                d_1177_v64_ = D2_DC6(d_1173_v60_)
                d_1178_v65_: D2
                d_1178_v65_ = D2_DC6(d_1176_v63_)
                d_1179_v67_: _dafny.Array
                nw216_ = _dafny.Array(None, 27)
                nw216_[int(0)] = d_1170_v56_
                nw216_[int(1)] = default__.fm54(d_1153_cf24_, (self).f39, len(d_1171_v57_), (self).f40, globalState)
                nw216_[int(2)] = (d_1170_v56_) | (d_1170_v56_)
                nw216_[int(3)] = d_1170_v56_
                nw216_[int(4)] = (d_1170_v56_) | (d_1170_v56_)
                nw216_[int(5)] = d_1170_v56_
                nw216_[int(6)] = d_1170_v56_
                nw216_[int(7)] = d_1170_v56_
                nw216_[int(8)] = _dafny.Map({len(_dafny.Set({(self).f40, (self).f40, (self).f29, False, (self).f40})): ((d_1170_v56_)[(d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]] if ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]) in (d_1170_v56_) else (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])})
                nw216_[int(9)] = default__.fm54((self).f40, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (self).f29, globalState)
                nw216_[int(10)] = d_1170_v56_
                nw216_[int(11)] = (_dafny.Map({(self).f39: (self).f39})).set((self).f39, (self).f39)
                def iife104_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in (_dafny.Map({(self).f39: p0})).keys.Elements:
                        d_1180_v58_: int = compr_54_
                        if (d_1180_v58_) in (_dafny.Map({(self).f39: p0})):
                            coll54_[default__.safeModuloInt(d_1180_v58_, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])] = (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]
                    return _dafny.Map(coll54_)
                nw216_[int(12)] = iife104_()
                
                nw216_[int(13)] = d_1170_v56_
                nw216_[int(14)] = (d_1170_v56_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scfyumh"))), (0) - ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]))
                nw216_[int(15)] = (d_1170_v56_) | (default__.fm54((self).fm31(globalState), (self).f39, (0) - ((self).f39), d_1153_cf24_, globalState))
                nw216_[int(16)] = (_dafny.Map({(d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]: (self).f39})) | (d_1170_v56_)
                nw216_[int(17)] = default__.fm54((self).f29, len(_dafny.SeqWithoutIsStrInference([d_1178_v65_])), (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], False, globalState)
                def iife105_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in _dafny.IntegerRange(891, -20):
                        d_1181_v66_: int = compr_55_
                        if ((891) <= (d_1181_v66_)) and ((d_1181_v66_) < (-20)):
                            coll55_[(d_1181_v66_) - ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])] = 349
                    return _dafny.Map(coll55_)
                nw216_[int(18)] = iife105_()
                
                nw216_[int(19)] = (d_1170_v56_) | (d_1170_v56_)
                nw216_[int(20)] = d_1170_v56_
                nw216_[int(21)] = d_1170_v56_
                nw216_[int(22)] = d_1170_v56_
                nw216_[int(23)] = _dafny.Map({896: (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]})
                nw216_[int(24)] = d_1170_v56_
                nw216_[int(25)] = d_1170_v56_
                nw216_[int(26)] = d_1170_v56_
                d_1179_v67_ = nw216_
                index166_ = default__.safeIndex(731, (d_1179_v67_).length(0))
                (d_1179_v67_)[index166_] = default__.fm54(d_1152_cf25_, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], not(d_1152_cf25_), globalState)
                index167_ = default__.safeIndex(731, (d_1179_v67_).length(0))
                (d_1179_v67_)[index167_] = (_dafny.Map({(self).f39: len(_dafny.Map({d_1152_cf25_: len(default__.fm55(d_1153_cf24_, d_1153_cf24_, globalState))}))})) | ((d_1170_v56_) | (d_1170_v56_))
                (globalState).f21 = (self).f29
            (globalState).f22 = d_1153_cf24_
            (globalState).f1 = (self).f39
            d_1182_v68_: _dafny.Map
            d_1182_v68_ = _dafny.Map({p0: (self).f39})
            d_1183_v69_: _dafny.Seq
            d_1183_v69_ = _dafny.SeqWithoutIsStrInference([self.f28, (self).f29])
            d_1182_v68_ = default__.fm56(default__.fm39(p0, globalState), d_1152_cf25_, len(d_1183_v69_), globalState)
        elif True:
            d_1184___mcc_h6_ = source22_.cf21
            d_1185_cf21_ = d_1184___mcc_h6_
            (globalState).f1 = 671
            (globalState).f21 = self.f28
            d_1186_v70_: str
            d_1186_v70_ = _dafny.CodePoint('a')
            d_1187_v71_: D5
            d_1187_v71_ = D5_DC13(D2_DC4(d_1186_v70_))
            d_1188_v72_: int
            d_1189_v73_: int
            d_1190_v74_: int
            d_1191_v75_: bool
            out27_: int
            out28_: int
            out29_: int
            out30_: bool
            out27_, out28_, out29_, out30_ = (self).m17((self).f40, (d_1187_v71_ if self.f28 else d_1187_v71_), (self).f40, globalState)
            d_1188_v72_ = out27_
            d_1189_v73_ = out28_
            d_1190_v74_ = out29_
            d_1191_v75_ = out30_
            d_1192_v76_: _dafny.Array
            def lambda24_(d_1193_v75_):
                def lambda25_(d_1194_i5_):
                    return d_1193_v75_

                return lambda25_

            init16_ = lambda24_(d_1191_v75_)
            nw217_ = _dafny.Array(None, 9)
            for i0_16_ in range(nw217_.length(0)):
                nw217_[i0_16_] = init16_(i0_16_)
            d_1192_v76_ = nw217_
            d_1195_v77_: _dafny.Array
            nw218_ = _dafny.Array(None, 5)
            nw218_[int(0)] = d_1192_v76_
            nw218_[int(1)] = d_1192_v76_
            nw218_[int(2)] = d_1192_v76_
            nw218_[int(3)] = d_1192_v76_
            nw218_[int(4)] = d_1192_v76_
            d_1195_v77_ = nw218_
            index168_ = default__.safeIndex(979, (d_1195_v77_).length(0))
            (d_1195_v77_)[index168_] = d_1192_v76_
            index169_ = default__.safeIndex(979, (d_1195_v77_).length(0))
            (d_1195_v77_)[index169_] = d_1192_v76_
        if (self).f29:
            (globalState).f13 = self.f28
            d_1196_v78_: _dafny.Array
            nw219_ = _dafny.Array(False, 12)
            d_1196_v78_ = nw219_
            index170_ = default__.safeIndex(568, (d_1196_v78_).length(0))
            (d_1196_v78_)[index170_] = self.f28
            index171_ = default__.safeIndex(568, (d_1196_v78_).length(0))
            (d_1196_v78_)[index171_] = self.f28
            d_1197_v79_: str
            d_1197_v79_ = _dafny.CodePoint('o')
            (globalState).f23 = d_1197_v79_
            (globalState).f24 = default__.safeModuloInt((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], -377)
            d_1198_v80_: _dafny.Map
            d_1198_v80_ = _dafny.Map({default__.fm38(globalState): self.f28})
            d_1198_v80_ = (d_1198_v80_).set((self).f29, (d_1196_v78_)[default__.safeIndex(568, (d_1196_v78_).length(0))])
        elif True:
            d_1199_v81_: D8
            d_1199_v81_ = D8_DC21(p0)
            d_1200_v82_: str
            d_1200_v82_ = _dafny.CodePoint('o')
            d_1201_v83_: D2
            d_1201_v83_ = D2_DC4(d_1200_v82_)
            d_1202_v84_: D5
            d_1202_v84_ = D5_DC13(d_1201_v83_)
            d_1203_v85_: int
            d_1204_v86_: int
            d_1205_v87_: int
            d_1206_v88_: bool
            out31_: int
            out32_: int
            out33_: int
            out34_: bool
            out31_, out32_, out33_, out34_ = (self).m17((d_1199_v81_).cf27, d_1202_v84_, p0, globalState)
            d_1203_v85_ = out31_
            d_1204_v86_ = out32_
            d_1205_v87_ = out33_
            d_1206_v88_ = out34_
            if (self).f29:
                d_1207_v89_: _dafny.Array
                nw220_ = _dafny.Array(None, 3)
                nw220_[int(0)] = (self).f40
                nw220_[int(1)] = d_1206_v88_
                nw220_[int(2)] = (self).f40
                d_1207_v89_ = nw220_
                d_1208_v90_: _dafny.Map
                d_1208_v90_ = _dafny.Map({d_1207_v89_: d_1207_v89_})
                d_1209_v91_: _dafny.MultiSet
                d_1209_v91_ = _dafny.MultiSet([d_1207_v89_, (((d_1208_v90_)[d_1207_v89_] if (d_1207_v89_) in (d_1208_v90_) else d_1207_v89_) if (self).f29 else d_1207_v89_)])
                d_1210_v92_: _dafny.Array
                def lambda26_(d_1211_p0_):
                    def lambda27_(d_1212_i6_):
                        return _dafny.Set({(self).f40, (self).f40, (self).f29, (D2_DC5(d_1211_p0_)).cf8})

                    return lambda27_

                init17_ = lambda26_(p0)
                nw221_ = _dafny.Array(None, 24)
                for i0_17_ in range(nw221_.length(0)):
                    nw221_[i0_17_] = init17_(i0_17_)
                d_1210_v92_ = nw221_
                rhs186_ = ((self).f29 if self.f28 else (self).f29)
                rhs187_ = ((d_1125_v25_)[default__.safeIndex(51, (d_1125_v25_).length(0))]) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gucpxn"))) + (d_1126_v26_))
                rhs188_ = ((_dafny.MultiSet([d_1207_v89_, d_1207_v89_])) | (d_1209_v91_)).intersection((d_1209_v91_).set(d_1207_v89_, default__.abs((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])))
                rhs189_ = d_1210_v92_
                lhs163_ = globalState
                lhs164_ = globalState
                lhs163_.f27 = rhs186_
                lhs164_.f27 = rhs187_
                d_1209_v91_ = rhs188_
                d_1210_v92_ = rhs189_
                (globalState).f1 = (self).f39
                d_1213_v93_: _dafny.Map
                d_1213_v93_ = _dafny.Map({(self).f40: True})
                d_1214_v94_: _dafny.Map
                d_1214_v94_ = _dafny.Map({(self).f40: ((d_1213_v93_)[True] if (True) in (d_1213_v93_) else (self).fm31(globalState))})
                (globalState).f1 = default__.safeDivisionInt(110, (0) - (default__.safeModuloInt((self).fm30(d_1123_v23_, globalState), len(d_1213_v93_))))
                d_1205_v87_ = (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))]
                d_1215_v95_: _dafny.Map
                d_1215_v95_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsd"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlsxv")))})
                d_1216_v96_: _dafny.Seq
                d_1216_v96_ = _dafny.SeqWithoutIsStrInference([d_1215_v95_])
                d_1217_v98_: D8
                d_1217_v98_ = D8_DC22()
                d_1218_v99_: _dafny.Map
                d_1218_v99_ = _dafny.Map({d_1217_v98_: (self).f40})
                d_1219_v100_: _dafny.Map
                d_1219_v100_ = _dafny.Map({261: d_1218_v99_})
                def iife106_():
                    coll56_ = _dafny.Map()
                    compr_56_: D8
                    for compr_56_ in (((d_1219_v100_)[d_1205_v87_] if (d_1205_v87_) in (d_1219_v100_) else d_1218_v99_)).keys.Elements:
                        d_1220_v97_: D8 = compr_56_
                        if (d_1220_v97_) in (((d_1219_v100_)[d_1205_v87_] if (d_1205_v87_) in (d_1219_v100_) else d_1218_v99_)):
                            coll56_[d_1220_v97_] = 649
                    return _dafny.Map(coll56_)
                d_1203_v85_ = (d_1203_v85_ if (408) not in ((d_1216_v96_)[default__.safeIndex(len(iife106_()
                ), len(d_1216_v96_))]) else 155)
            elif True:
                d_1221_v101_: _dafny.Map
                d_1221_v101_ = _dafny.Map({(self).f29: d_1206_v88_})
                d_1222_v102_: _dafny.Set
                d_1222_v102_ = _dafny.Set({d_1221_v101_})
                d_1223_v103_: _dafny.Array
                nw222_ = _dafny.Array(None, 15)
                nw222_[int(0)] = d_1201_v83_
                nw222_[int(1)] = D2_DC4(d_1200_v82_)
                nw222_[int(2)] = d_1201_v83_
                nw222_[int(3)] = d_1201_v83_
                nw222_[int(4)] = d_1201_v83_
                nw222_[int(5)] = D2_DC4(d_1200_v82_)
                nw222_[int(6)] = d_1201_v83_
                nw222_[int(7)] = d_1201_v83_
                nw222_[int(8)] = D2_DC4(default__.fm42((self).f39, (d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))], globalState))
                nw222_[int(9)] = d_1201_v83_
                nw222_[int(10)] = D2_DC4(d_1200_v82_)
                nw222_[int(11)] = d_1201_v83_
                nw222_[int(12)] = d_1201_v83_
                nw222_[int(13)] = d_1201_v83_
                nw222_[int(14)] = D2_DC4(d_1200_v82_)
                d_1223_v103_ = nw222_
                index172_ = default__.safeIndex(814, (d_1223_v103_).length(0))
                (d_1223_v103_)[index172_] = (d_1201_v83_ if (self).f29 else d_1201_v83_)
                index173_ = default__.safeIndex(814, (d_1223_v103_).length(0))
                rhs190_ = not (p0) or (False)
                rhs191_ = d_1222_v102_
                rhs192_ = d_1201_v83_
                rhs193_ = (_dafny.CodePoint('w') if p0 else (d_1200_v82_ if ((d_1221_v101_)[(self).f40] if ((self).f40) in (d_1221_v101_) else (self).f29) else ((d_1125_v25_)[default__.safeIndex(51, (d_1125_v25_).length(0))])[default__.safeIndex(d_1203_v85_, len((d_1125_v25_)[default__.safeIndex(51, (d_1125_v25_).length(0))]))]))
                lhs165_ = self
                lhs166_ = d_1223_v103_
                lhs167_ = default__.safeIndex(814, (d_1223_v103_).length(0))
                lhs168_ = globalState
                lhs165_.f28 = rhs190_
                d_1222_v102_ = rhs191_
                lhs166_[lhs167_] = rhs192_
                lhs168_.f23 = rhs193_
                d_1224_v104_: _dafny.MultiSet
                d_1224_v104_ = _dafny.MultiSet([(self).f40])
                d_1225_v105_: _dafny.Set
                d_1225_v105_ = _dafny.Set({(self).f40})
                d_1226_v106_: _dafny.MultiSet
                d_1226_v106_ = _dafny.MultiSet([not((d_1224_v104_).issubset(d_1224_v104_)), not ((self).f29) or ((self).fm32(d_1225_v105_, (0) - ((_dafny.MultiSet([(self).f39, d_1203_v85_])).cardinality), (self).f40, globalState)), (self).f29, d_1206_v88_])
                index174_ = default__.safeIndex(51, (d_1125_v25_).length(0))
                rhs194_ = (self).f39
                rhs195_ = (_dafny.SeqWithoutIsStrInference([d_1200_v82_ for d_1227_i7_ in range(default__.abs(59))])) < (d_1126_v26_)
                rhs196_ = d_1206_v88_
                rhs197_ = (d_1226_v106_).cardinality
                rhs198_ = d_1126_v26_
                lhs169_ = globalState
                lhs170_ = globalState
                lhs171_ = globalState
                lhs172_ = d_1125_v25_
                lhs173_ = default__.safeIndex(51, (d_1125_v25_).length(0))
                d_1205_v87_ = rhs194_
                lhs169_.f22 = rhs195_
                lhs170_.f27 = rhs196_
                lhs171_.f1 = rhs197_
                lhs172_[lhs173_] = rhs198_
                rhs199_ = d_1203_v85_
                rhs200_ = 563
                rhs201_ = d_1205_v87_
                rhs202_ = (0) - ((0) - ((0) - ((d_1121_v22_)[default__.safeIndex(684, (d_1121_v22_).length(0))])))
                rhs203_ = (self).f29
                lhs174_ = globalState
                lhs175_ = globalState
                d_1204_v86_ = rhs199_
                lhs174_.f1 = rhs200_
                d_1205_v87_ = rhs201_
                d_1203_v85_ = rhs202_
                lhs175_.f13 = rhs203_
                d_1228_v107_: _dafny.Seq
                d_1228_v107_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1228_v107_ = _dafny.SeqWithoutIsStrInference([(d_1203_v85_) in (d_1123_v23_)])
                d_1229_v108_: C2
                nw223_ = C2()
                nw223_.ctor__()
                d_1229_v108_ = nw223_
            index175_ = default__.safeIndex(684, (d_1121_v22_).length(0))
            (d_1121_v22_)[index175_] = (self).fm30(d_1123_v23_, globalState)
            if (self).f40:
                d_1230_v109_: _dafny.Map
                d_1230_v109_ = _dafny.Map({d_1121_v22_: d_1205_v87_})
                d_1231_v110_: _dafny.Array
                nw224_ = _dafny.Array(None, 1)
                nw224_[int(0)] = d_1230_v109_
                d_1231_v110_ = nw224_
                index176_ = default__.safeIndex(834, (d_1231_v110_).length(0))
                (d_1231_v110_)[index176_] = d_1230_v109_
                index177_ = default__.safeIndex(834, (d_1231_v110_).length(0))
                (d_1231_v110_)[index177_] = d_1230_v109_
                d_1232_v111_: _dafny.Set
                d_1232_v111_ = _dafny.Set({(self).f40, (self).f40})
                (globalState).f1 = default__.safeDivisionInt((d_1205_v87_) + (len(d_1232_v111_)), d_1204_v86_)
                d_1233_v112_: _dafny.Array
                nw225_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1233_v112_ = nw225_
                d_1234_v113_: _dafny.Array
                def lambda28_(d_1235_i8_):
                    return self.f28

                init18_ = lambda28_
                nw226_ = _dafny.Array(None, 7)
                for i0_18_ in range(nw226_.length(0)):
                    nw226_[i0_18_] = init18_(i0_18_)
                d_1234_v113_ = nw226_
                index178_ = default__.safeIndex(163, (d_1233_v112_).length(0))
                (d_1233_v112_)[index178_] = (D5_DC12(d_1234_v113_)).cf17
                index179_ = default__.safeIndex(163, (d_1233_v112_).length(0))
                nw227_ = _dafny.Array(False, 2)
                (d_1233_v112_)[index179_] = nw227_
                d_1236_v114_: _dafny.Seq
                d_1236_v114_ = _dafny.SeqWithoutIsStrInference([d_1203_v85_])
                d_1237_v115_: D4
                d_1237_v115_ = D4_DC9((_dafny.MultiSet(d_1236_v114_)).set(d_1203_v85_, default__.abs(354)))
                d_1238_v116_: D4
                d_1238_v116_ = D4_DC11(d_1237_v115_)
                d_1238_v116_ = D4_DC11(d_1237_v115_)
                d_1126_v26_ = _dafny.SeqWithoutIsStrInference([d_1200_v82_ for d_1239_i9_ in range(default__.abs(-196))])
            elif True:
                d_1240_v117_: _dafny.Array
                nw228_ = _dafny.Array(False, 13)
                d_1240_v117_ = nw228_
                index180_ = default__.safeIndex(317, (d_1240_v117_).length(0))
                (d_1240_v117_)[index180_] = d_1206_v88_
                index181_ = default__.safeIndex(317, (d_1240_v117_).length(0))
                (d_1240_v117_)[index181_] = self.f28
                d_1241_v118_: _dafny.Seq
                d_1241_v118_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1200_v82_ for d_1242_i10_ in range(default__.abs(870))])])
                d_1243_v119_: _dafny.MultiSet
                d_1243_v119_ = _dafny.MultiSet([d_1203_v85_, d_1205_v87_])
                index182_ = default__.safeIndex(684, (d_1121_v22_).length(0))
                (d_1121_v22_)[index182_] = (d_1204_v86_ if self.f28 else (len((d_1241_v118_)[default__.safeIndex(((d_1243_v119_)[d_1203_v85_] if (d_1203_v85_) in (d_1243_v119_) else (self).f39), len(d_1241_v118_))]) if self.f28 else d_1204_v86_))
                d_1244_v120_: _dafny.MultiSet
                d_1244_v120_ = _dafny.MultiSet([(self).f40])
                index183_ = default__.safeIndex(317, (d_1240_v117_).length(0))
                (d_1240_v117_)[index183_] = (((d_1244_v120_).set(p0, default__.abs(len((d_1125_v25_)[default__.safeIndex(51, (d_1125_v25_).length(0))])))) - (d_1244_v120_)).issubset(d_1244_v120_)
                index184_ = default__.safeIndex(684, (d_1121_v22_).length(0))
                (d_1121_v22_)[index184_] = d_1203_v85_
                index185_ = default__.safeIndex(684, (d_1121_v22_).length(0))
                (d_1121_v22_)[index185_] = 766
            d_1245_v121_: _dafny.Array
            nw229_ = _dafny.Array(_dafny.Map({}), 21)
            d_1245_v121_ = nw229_
            d_1246_v122_: _dafny.Map
            d_1246_v122_ = _dafny.Map({d_1206_v88_: d_1206_v88_})
            index186_ = default__.safeIndex(983, (d_1245_v121_).length(0))
            (d_1245_v121_)[index186_] = d_1246_v122_
            index187_ = default__.safeIndex(983, (d_1245_v121_).length(0))
            (d_1245_v121_)[index187_] = d_1246_v122_

    def m16(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        r2: bool = False
        if True:
            d_1247_v0_: _dafny.Map
            d_1247_v0_ = _dafny.Map({476: p0})
            d_1248_v1_: _dafny.Array
            nw230_ = _dafny.Array(int(0), 17)
            d_1248_v1_ = nw230_
            d_1249_v2_: _dafny.Map
            d_1249_v2_ = _dafny.Map({d_1248_v1_: (self).f39})
            d_1250_v3_: _dafny.Array
            nw231_ = _dafny.Array(None, 4)
            nw231_[int(0)] = _dafny.Map({(self).f39: (self).f40})
            nw231_[int(1)] = _dafny.Map({205: not(p1)})
            nw231_[int(2)] = d_1247_v0_
            nw231_[int(3)] = _dafny.Map({len(d_1249_v2_): p1})
            d_1250_v3_ = nw231_
            index188_ = default__.safeIndex(707, (d_1250_v3_).length(0))
            (d_1250_v3_)[index188_] = (d_1247_v0_) | (d_1247_v0_)
            index189_ = default__.safeIndex(707, (d_1250_v3_).length(0))
            (d_1250_v3_)[index189_] = d_1247_v0_
            d_1251_v4_: _dafny.Seq
            d_1251_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xaaob"))
            d_1252_v5_: str
            d_1252_v5_ = _dafny.CodePoint('b')
            d_1251_v4_ = (p2).set(default__.safeIndex((self).f39, len(p2)), d_1252_v5_)
            d_1253_v6_: _dafny.MultiSet
            d_1253_v6_ = _dafny.MultiSet([(self).f39, (self).f39])
            (globalState).f21 = not((((d_1253_v6_ if p0 else d_1253_v6_)).cardinality) == (default__.safeDivisionInt((0) - ((self).f39), len(d_1251_v4_))))
            d_1252_v5_ = _dafny.CodePoint('i')
            d_1254_v7_: D10
            d_1254_v7_ = D10_DC27(((d_1247_v0_)[(d_1253_v6_).cardinality] if ((d_1253_v6_).cardinality) in (d_1247_v0_) else (self).f40), self.f28, (self).f29, (self).f39, p1)
            d_1255_v8_: _dafny.Map
            d_1255_v8_ = _dafny.Map({((d_1254_v7_).cf43) >= ((self).f39): (self).f29})
            d_1256_v9_: D10
            d_1256_v9_ = D10_DC28(d_1255_v8_)
            d_1255_v8_ = (d_1256_v9_).cf45
        elif True:
            (globalState).f22 = not(p0)
            d_1257_v10_: _dafny.Map
            d_1257_v10_ = _dafny.Map({self.f28: (self).f39})
            d_1258_v11_: _dafny.MultiSet
            d_1258_v11_ = _dafny.MultiSet([p1, (self).f29])
            d_1257_v10_ = (d_1257_v10_).set(((self).f39) < ((self).f39), (((d_1258_v11_)[True] if (True) in (d_1258_v11_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1259_i0_ in range(default__.abs(-800))])))) - ((self).f39))
            rhs204_ = (p2) <= (p2)
            rhs205_ = (self).f39
            lhs176_ = globalState
            lhs177_ = globalState
            lhs176_.f13 = rhs204_
            lhs177_.f1 = rhs205_
            d_1260_v12_: C1
            nw232_ = C1()
            nw232_.ctor__(p0, (self).f40)
            d_1260_v12_ = nw232_
            d_1261_v13_: _dafny.Array
            nw233_ = _dafny.Array(int(0), 8)
            d_1261_v13_ = nw233_
            index190_ = default__.safeIndex(286, (d_1261_v13_).length(0))
            (d_1261_v13_)[index190_] = (self).f39
            d_1262_v14_: _dafny.Set
            d_1262_v14_ = _dafny.Set({(self).f40, self.f28})
            d_1263_v15_: _dafny.Seq
            d_1263_v15_ = _dafny.SeqWithoutIsStrInference([(self).f39, (self).f39])
            d_1264_v16_: _dafny.Seq
            d_1264_v16_ = _dafny.SeqWithoutIsStrInference([D7_DC18((self).f40)])
            d_1265_v17_: _dafny.Array
            nw234_ = _dafny.Array(None, 27)
            nw234_[int(0)] = p0
            nw234_[int(1)] = ((self).f39) >= ((self).f39)
            nw234_[int(2)] = ((self).f39) <= (6)
            nw234_[int(3)] = not ((self).fm32(d_1262_v14_, (d_1263_v15_)[default__.safeIndex((self).f39, len(d_1263_v15_))], self.f28, globalState)) or ((self).f40)
            nw234_[int(4)] = p1
            nw234_[int(5)] = ((self).f39) != ((self).f39)
            nw234_[int(6)] = not (p1) or (p0)
            nw234_[int(7)] = (self).f29
            nw234_[int(8)] = (self).fm32(_dafny.Set({self.f28}), (self).f39, (self).f40, globalState)
            nw234_[int(9)] = p0
            nw234_[int(10)] = (self).f40
            nw234_[int(11)] = p1
            nw234_[int(12)] = (self).f40
            nw234_[int(13)] = p0
            nw234_[int(14)] = (self).f40
            nw234_[int(15)] = p1
            nw234_[int(16)] = False
            nw234_[int(17)] = (self).f29
            nw234_[int(18)] = (d_1264_v16_) < (d_1264_v16_)
            nw234_[int(19)] = p1
            nw234_[int(20)] = (p2) == (default__.fm51((self).f39, (self).f39, globalState))
            nw234_[int(21)] = (self).f29
            nw234_[int(22)] = (self).f40
            nw234_[int(23)] = (self).f29
            nw234_[int(24)] = p0
            nw234_[int(25)] = self.f28
            nw234_[int(26)] = True
            d_1265_v17_ = nw234_
            index191_ = default__.safeIndex(249, (d_1265_v17_).length(0))
            (d_1265_v17_)[index191_] = (self).f29
            index192_ = default__.safeIndex(286, (d_1261_v13_).length(0))
            index193_ = default__.safeIndex(249, (d_1265_v17_).length(0))
            rhs206_ = (self).f39
            rhs207_ = d_1265_v17_
            rhs208_ = (self).f29
            rhs209_ = p0
            lhs178_ = d_1261_v13_
            lhs179_ = default__.safeIndex(286, (d_1261_v13_).length(0))
            lhs180_ = d_1265_v17_
            lhs181_ = default__.safeIndex(249, (d_1265_v17_).length(0))
            lhs182_ = globalState
            lhs178_[lhs179_] = rhs206_
            d_1265_v17_ = rhs207_
            lhs180_[lhs181_] = rhs208_
            lhs182_.f27 = rhs209_
        hi3_ = (self).f39
        for d_1266_i1_ in range(default__.safeModuloInt((self).f39, (self).f39), hi3_):
            d_1267_v18_: _dafny.Array
            nw235_ = _dafny.Array(False, 25)
            d_1267_v18_ = nw235_
            index194_ = default__.safeIndex(175, (d_1267_v18_).length(0))
            (d_1267_v18_)[index194_] = (self).f29
            d_1268_v19_: _dafny.MultiSet
            d_1268_v19_ = _dafny.MultiSet([d_1266_i1_])
            index195_ = default__.safeIndex(175, (d_1267_v18_).length(0))
            (d_1267_v18_)[index195_] = (((self).f39) - ((self).f39)) not in (default__.fm37((d_1268_v19_).cardinality, p0, p1, globalState))
            d_1269_v20_: _dafny.Seq
            d_1269_v20_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            d_1269_v20_ = d_1269_v20_
            (globalState).f27 = (d_1269_v20_)[default__.safeIndex(988, len(d_1269_v20_))]
            if (d_1267_v18_)[default__.safeIndex(175, (d_1267_v18_).length(0))]:
                (globalState).f24 = (0) - (default__.safeDivisionInt((d_1266_i1_) * (814), d_1266_i1_))
                d_1270_v21_: _dafny.Array
                def lambda29_(d_1271_p0_):
                    def lambda30_(d_1272_i2_):
                        return D9_DC23(_dafny.MultiSet([d_1271_p0_]))

                    return lambda30_

                init19_ = lambda29_(p0)
                nw236_ = _dafny.Array(None, 17)
                for i0_19_ in range(nw236_.length(0)):
                    nw236_[i0_19_] = init19_(i0_19_)
                d_1270_v21_ = nw236_
                d_1273_v22_: str
                d_1273_v22_ = _dafny.CodePoint('a')
                d_1274_v23_: _dafny.Map
                d_1274_v23_ = _dafny.Map({d_1266_i1_: d_1266_i1_})
                d_1275_v24_: _dafny.Map
                d_1275_v24_ = _dafny.Map({d_1274_v23_: (0) - (-610)})
                d_1276_v25_: D9
                d_1276_v25_ = D9_DC23(default__.fm57(d_1273_v22_, len(d_1275_v24_), d_1266_i1_, globalState))
                index196_ = default__.safeIndex(631, (d_1270_v21_).length(0))
                (d_1270_v21_)[index196_] = d_1276_v25_
                index197_ = default__.safeIndex(631, (d_1270_v21_).length(0))
                (d_1270_v21_)[index197_] = d_1276_v25_
                d_1277_v26_: _dafny.Seq
                d_1277_v26_ = _dafny.SeqWithoutIsStrInference([d_1267_v18_, d_1267_v18_])
                d_1278_v27_: _dafny.Map
                d_1278_v27_ = _dafny.Map({D5_DC12(d_1267_v18_): (d_1277_v26_) + (d_1277_v26_)})
                d_1279_v28_: D5
                d_1279_v28_ = D5_DC12(d_1267_v18_)
                d_1278_v27_ = (d_1278_v27_).set(d_1279_v28_, (d_1277_v26_) + (_dafny.SeqWithoutIsStrInference([d_1267_v18_])))
                (globalState).f24 = ((d_1268_v19_)[(self).f39] if ((self).f39) in (d_1268_v19_) else d_1266_i1_)
                r2 = ((self).fm31(globalState) if (self).f29 else not (p0) or (False))
            elif True:
                d_1280_v29_: D7
                d_1280_v29_ = D7_DC18(self.f28)
                d_1280_v29_ = d_1280_v29_
                r0 = (self).f39
                d_1281_v30_: _dafny.Seq
                d_1281_v30_ = _dafny.SeqWithoutIsStrInference([d_1266_i1_])
                (globalState).f24 = (d_1281_v30_)[default__.safeIndex(d_1266_i1_, len(d_1281_v30_))]
                d_1282_v31_: _dafny.Map
                d_1282_v31_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([(d_1267_v18_)[default__.safeIndex(175, (d_1267_v18_).length(0))], (d_1267_v18_)[default__.safeIndex(175, (d_1267_v18_).length(0))], (d_1267_v18_)[default__.safeIndex(175, (d_1267_v18_).length(0))]])})
                d_1282_v31_ = ((d_1282_v31_) | (d_1282_v31_)) | (_dafny.Map({(self).f40: d_1269_v20_}))
                d_1283_v32_: _dafny.Array
                nw237_ = _dafny.Array(int(0), 13)
                d_1283_v32_ = nw237_
                index198_ = default__.safeIndex(70, (d_1283_v32_).length(0))
                (d_1283_v32_)[index198_] = (self).f39
                d_1284_v33_: T0
                nw238_ = C1()
                nw238_.ctor__(p0, p1)
                d_1284_v33_ = nw238_
                d_1285_v34_: D8
                d_1285_v34_ = D8_DC20(d_1284_v33_)
                d_1286_v35_: D8
                d_1286_v35_ = D8_DC20((d_1285_v34_).cf26)
                d_1287_v36_: D9
                d_1287_v36_ = D9_DC24(_dafny.CodePoint('s'), p1, d_1266_i1_, d_1286_v35_, (self).f40)
                d_1288_v37_: _dafny.Set
                d_1288_v37_ = _dafny.Set({(d_1287_v36_).cf31})
                index199_ = default__.safeIndex(70, (d_1283_v32_).length(0))
                rhs210_ = len(p2)
                rhs211_ = _dafny.MultiSet([d_1266_i1_, d_1266_i1_, (self).f39, d_1266_i1_])
                rhs212_ = (self).fm30((d_1288_v37_ if not(default__.fm38(globalState)) else d_1288_v37_), globalState)
                rhs213_ = d_1283_v32_
                lhs183_ = d_1283_v32_
                lhs184_ = default__.safeIndex(70, (d_1283_v32_).length(0))
                lhs185_ = globalState
                lhs183_[lhs184_] = rhs210_
                d_1268_v19_ = rhs211_
                lhs185_.f24 = rhs212_
                d_1283_v32_ = rhs213_
        d_1289_v38_: _dafny.Map
        d_1289_v38_ = _dafny.Map({p2: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urnsxrft"))) != (p2)})
        d_1289_v38_ = (d_1289_v38_).set(p2, (self).f29)
        if True:
            if p0:
                (globalState).f13 = (self).f40
                d_1290_v39_: _dafny.Seq
                d_1290_v39_ = _dafny.SeqWithoutIsStrInference([p2, p2, default__.fm51((self).f39, (self).f39, globalState), p2, p2])
                d_1291_v40_: D2
                d_1291_v40_ = D2_DC6(D2_DC5(p0))
                d_1292_v41_: D2
                d_1292_v41_ = D2_DC6(d_1291_v40_)
                d_1293_v42_: _dafny.Seq
                d_1293_v42_ = _dafny.SeqWithoutIsStrInference([d_1292_v41_, d_1292_v41_])
                rhs214_ = (d_1293_v42_) != (_dafny.SeqWithoutIsStrInference([D2_DC6(d_1291_v40_) for d_1294_i3_ in range(default__.abs(581))]))
                rhs215_ = (self).f39
                rhs216_ = 751
                rhs217_ = d_1290_v39_
                lhs186_ = globalState
                lhs187_ = globalState
                lhs188_ = globalState
                lhs186_.f27 = rhs214_
                lhs187_.f1 = rhs215_
                lhs188_.f1 = rhs216_
                d_1290_v39_ = rhs217_
                d_1295_v43_: _dafny.Array
                nw239_ = _dafny.Array(int(0), 18)
                d_1295_v43_ = nw239_
                d_1296_v44_: str
                d_1296_v44_ = _dafny.CodePoint('l')
                d_1297_v45_: _dafny.Set
                d_1297_v45_ = _dafny.Set({(self).f39, (self).f39})
                nw240_ = _dafny.Array(None, 7)
                nw240_[int(0)] = (self).f39
                nw240_[int(1)] = (self).f39
                nw240_[int(2)] = len((p2) + (p2))
                nw240_[int(3)] = (0) - (len(_dafny.SeqWithoutIsStrInference([default__.fm42((self).f39, len(p2), globalState), d_1296_v44_])))
                nw240_[int(4)] = default__.safeModuloInt((self).f39, 7)
                nw240_[int(5)] = (self).fm30(d_1297_v45_, globalState)
                nw240_[int(6)] = default__.safeModuloInt((self).f39, 424)
                d_1295_v43_ = nw240_
                d_1298_v46_: _dafny.Map
                d_1298_v46_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1292_v41_])): (self).f29})
                rhs218_ = d_1296_v44_
                rhs219_ = (d_1296_v44_ if (self).f29 else d_1296_v44_)
                rhs220_ = d_1298_v46_
                lhs189_ = globalState
                r1 = rhs218_
                lhs189_.f23 = rhs219_
                d_1298_v46_ = rhs220_
                d_1299_v47_: _dafny.Set
                d_1299_v47_ = _dafny.Set({p0})
                (globalState).f24 = (((self).f39 if (self).f29 else len(d_1299_v47_)) if (self).f29 else (0) - ((self).f39))
            elif True:
                (globalState).f27 = self.f28
                (self).f28 = (self).f29
                d_1300_v48_: _dafny.Array
                nw241_ = _dafny.Array(False, 4)
                d_1300_v48_ = nw241_
                index200_ = default__.safeIndex(64, (d_1300_v48_).length(0))
                (d_1300_v48_)[index200_] = default__.fm38(globalState)
                d_1301_v49_: _dafny.Seq
                d_1301_v49_ = _dafny.SeqWithoutIsStrInference([d_1300_v48_])
                d_1302_v50_: _dafny.Seq
                d_1302_v50_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1303_v51_: _dafny.Set
                d_1303_v51_ = _dafny.Set({d_1302_v50_, d_1302_v50_, d_1302_v50_, d_1302_v50_, d_1302_v50_})
                d_1304_v52_: _dafny.Map
                d_1304_v52_ = _dafny.Map({(self).f29: d_1303_v51_})
                d_1305_v53_: _dafny.Set
                d_1305_v53_ = _dafny.Set({p0, (self).f40, True})
                d_1306_v54_: _dafny.Set
                d_1306_v54_ = _dafny.Set({p0, (self).fm32(d_1305_v53_, (self).f39, self.f28, globalState)})
                d_1307_v55_: _dafny.Set
                d_1307_v55_ = _dafny.Set({(self).f39})
                d_1308_v56_: _dafny.Map
                d_1308_v56_ = _dafny.Map({(self).f29: len(_dafny.SeqWithoutIsStrInference([(self).fm30(d_1307_v55_, globalState), (self).f39]))})
                d_1309_v57_: _dafny.Seq
                d_1309_v57_ = _dafny.SeqWithoutIsStrInference([(self).f39, (0) - ((self).f39), (self).f39, -492, (self).f39])
                d_1310_v58_: _dafny.Seq
                d_1310_v58_ = _dafny.SeqWithoutIsStrInference([d_1309_v57_])
                index201_ = default__.safeIndex(64, (d_1300_v48_).length(0))
                rhs221_ = p1
                rhs222_ = len(((d_1304_v52_)[(self).fm32(d_1305_v53_, (self).f39, False, globalState)] if ((self).fm32(d_1305_v53_, (self).f39, False, globalState)) in (d_1304_v52_) else (d_1303_v51_) - (d_1303_v51_)))
                rhs223_ = d_1301_v49_
                rhs224_ = default__.safeDivisionInt(default__.safeDivisionInt(len(p2), (self).f39), (0) - (((d_1308_v56_)[self.f28] if (self.f28) in (d_1308_v56_) else len((d_1310_v58_)[default__.safeIndex((self).f39, len(d_1310_v58_))]))))
                lhs190_ = d_1300_v48_
                lhs191_ = default__.safeIndex(64, (d_1300_v48_).length(0))
                lhs192_ = globalState
                lhs193_ = globalState
                lhs190_[lhs191_] = rhs221_
                lhs192_.f1 = rhs222_
                d_1301_v49_ = rhs223_
                lhs193_.f1 = rhs224_
                d_1311_v59_: str
                d_1311_v59_ = _dafny.CodePoint('v')
                d_1312_v60_: _dafny.Map
                d_1312_v60_ = _dafny.Map({(self).f39: d_1311_v59_})
                d_1312_v60_ = (d_1312_v60_).set(-885, d_1311_v59_)
                d_1313_v61_: C0
                nw242_ = C0()
                nw242_.ctor__()
                d_1313_v61_ = nw242_
            if self.f28:
                d_1314_v63_: _dafny.MultiSet
                d_1314_v63_ = _dafny.MultiSet([(self).f39])
                d_1315_v64_: _dafny.Map
                d_1315_v64_ = _dafny.Map({(self).f39: d_1314_v63_})
                d_1316_v65_: _dafny.Array
                nw243_ = _dafny.Array(None, 21)
                nw243_[int(0)] = (self).f40
                nw243_[int(1)] = p0
                nw243_[int(2)] = p1
                nw243_[int(3)] = p1
                nw243_[int(4)] = p0
                nw243_[int(5)] = p1
                nw243_[int(6)] = not((self).f40)
                nw243_[int(7)] = p0
                nw243_[int(8)] = self.f28
                nw243_[int(9)] = p0
                nw243_[int(10)] = p1
                nw243_[int(11)] = self.f28
                nw243_[int(12)] = (self).f29
                nw243_[int(13)] = p1
                nw243_[int(14)] = p1
                nw243_[int(15)] = p0
                nw243_[int(16)] = (self).f40
                nw243_[int(17)] = not((self).f29)
                nw243_[int(18)] = self.f28
                nw243_[int(19)] = p0
                nw243_[int(20)] = (self).f40
                d_1316_v65_ = nw243_
                d_1317_v66_: _dafny.Map
                d_1317_v66_ = _dafny.Map({d_1316_v65_: p0})
                d_1318_v67_: _dafny.Map
                d_1318_v67_ = _dafny.Map({self.f28: len(d_1317_v66_)})
                d_1319_v68_: _dafny.Seq
                d_1319_v68_ = _dafny.SeqWithoutIsStrInference([(self).f39, (self).f39, (self).f39, ((d_1318_v67_)[(self).f40] if ((self).f40) in (d_1318_v67_) else (self).f39)])
                d_1320_v69_: _dafny.Map
                d_1320_v69_ = _dafny.Map({d_1319_v68_: d_1315_v64_})
                d_1321_v70_: _dafny.Array
                nw244_ = _dafny.Array(None, 7)
                def iife107_():
                    coll57_ = _dafny.Map()
                    compr_57_: int
                    for compr_57_ in _dafny.IntegerRange(656, 449):
                        d_1322_v62_: int = compr_57_
                        if ((656) <= (d_1322_v62_)) and ((d_1322_v62_) < (449)):
                            coll57_[default__.safeModuloInt(d_1322_v62_, (self).f39)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([702, len(_dafny.Set({(self).f40}))]))
                    return _dafny.Map(coll57_)
                nw244_[int(0)] = iife107_()
                
                nw244_[int(1)] = (D13_DC34(d_1315_v64_)).cf50
                nw244_[int(2)] = (d_1315_v64_) | (d_1315_v64_)
                nw244_[int(3)] = (d_1315_v64_) | (d_1315_v64_)
                nw244_[int(4)] = d_1315_v64_
                nw244_[int(5)] = d_1315_v64_
                nw244_[int(6)] = ((d_1320_v69_)[d_1319_v68_] if (d_1319_v68_) in (d_1320_v69_) else _dafny.Map({(self).f39: d_1314_v63_}))
                d_1321_v70_ = nw244_
                index202_ = default__.safeIndex(462, (d_1321_v70_).length(0))
                (d_1321_v70_)[index202_] = d_1315_v64_
                index203_ = default__.safeIndex(462, (d_1321_v70_).length(0))
                (d_1321_v70_)[index203_] = d_1315_v64_
                d_1323_v71_: _dafny.MultiSet
                d_1323_v71_ = _dafny.MultiSet([(self).f40, p0])
                d_1324_v72_: _dafny.Set
                d_1324_v72_ = _dafny.Set({((_dafny.MultiSet([True]) if False else d_1323_v71_)).cardinality})
                d_1325_v73_: _dafny.Seq
                d_1325_v73_ = _dafny.SeqWithoutIsStrInference([d_1324_v72_, _dafny.Set({(self).f39, len(_dafny.SeqWithoutIsStrInference([p0]))}), (d_1324_v72_).intersection(d_1324_v72_), d_1324_v72_])
                d_1324_v72_ = (d_1325_v73_)[default__.safeIndex((self).f39, len(d_1325_v73_))]
                d_1323_v71_ = d_1323_v71_
                d_1318_v67_ = (d_1318_v67_).set(True, (self).f39)
                d_1326_v74_: D7
                d_1326_v74_ = D7_DC17(p1)
                d_1327_v75_: _dafny.Map
                d_1327_v75_ = _dafny.Map({d_1326_v74_: default__.fm42((self).f39, 272, globalState)})
                d_1328_v76_: str
                d_1328_v76_ = _dafny.CodePoint('u')
                d_1327_v75_ = (d_1327_v75_).set(d_1326_v74_, d_1328_v76_)
            elif True:
                (globalState).f22 = not(self.f28)
                d_1329_v77_: C0
                nw245_ = C0()
                nw245_.ctor__()
                d_1329_v77_ = nw245_
                d_1330_v78_: D2
                d_1330_v78_ = D2_DC5(p0)
                d_1331_v79_: _dafny.Seq
                d_1331_v79_ = _dafny.SeqWithoutIsStrInference([p0, (d_1330_v78_).cf8, p1])
                d_1332_v80_: D1
                d_1332_v80_ = D1_DC2((self).f39, (self).f39, p2, self.f28)
                d_1333_v81_: _dafny.Map
                d_1333_v81_ = _dafny.Map({(d_1332_v80_).cf5: _dafny.SeqWithoutIsStrInference([(self).f39 for d_1334_i4_ in range(default__.abs(506))])})
                r2 = not((d_1331_v79_)[default__.safeIndex(len(d_1333_v81_), len(d_1331_v79_))])
                d_1335_v82_: str
                d_1335_v82_ = _dafny.CodePoint('o')
                d_1336_v83_: _dafny.Set
                d_1336_v83_ = _dafny.Set({d_1335_v82_})
                (globalState).f13 = (d_1336_v83_).issubset(default__.fm58(p0, p0, (self).f39, p1, globalState))
                d_1337_v84_: _dafny.Array
                def lambda31_(d_1338_i5_):
                    return (d_1338_i5_) + ((self).f39)

                init20_ = lambda31_
                nw246_ = _dafny.Array(None, 27)
                for i0_20_ in range(nw246_.length(0)):
                    nw246_[i0_20_] = init20_(i0_20_)
                d_1337_v84_ = nw246_
                d_1339_v85_: _dafny.Seq
                d_1339_v85_ = _dafny.SeqWithoutIsStrInference([d_1337_v84_])
                (globalState).f24 = len((d_1339_v85_) + (d_1339_v85_))
            d_1340_v86_: _dafny.Set
            d_1340_v86_ = _dafny.Set({(0) - ((self).f39), (self).f39})
            rhs225_ = ((d_1340_v86_) - (default__.fm35(globalState))).ispropersubset(d_1340_v86_)
            rhs226_ = self.f28
            rhs227_ = ((0) - (default__.safeModuloInt((self).f39, (self).f39))) + ((self).f39)
            lhs194_ = globalState
            lhs195_ = globalState
            lhs196_ = globalState
            lhs194_.f27 = rhs225_
            lhs195_.f13 = rhs226_
            lhs196_.f1 = rhs227_
            d_1341_v87_: _dafny.Seq
            d_1341_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oexdvydna"))
            d_1341_v87_ = p2
            d_1342_v88_: _dafny.Array
            nw247_ = _dafny.Array(_dafny.CodePoint('D'), 18)
            d_1342_v88_ = nw247_
            d_1343_v89_: _dafny.Seq
            d_1343_v89_ = _dafny.SeqWithoutIsStrInference([d_1342_v88_, d_1342_v88_, d_1342_v88_, d_1342_v88_])
            d_1344_v90_: _dafny.Array
            nw248_ = _dafny.Array(None, 24)
            nw248_[int(0)] = d_1342_v88_
            nw248_[int(1)] = d_1342_v88_
            nw248_[int(2)] = d_1342_v88_
            nw248_[int(3)] = d_1342_v88_
            nw248_[int(4)] = d_1342_v88_
            nw248_[int(5)] = d_1342_v88_
            nw248_[int(6)] = d_1342_v88_
            nw248_[int(7)] = d_1342_v88_
            nw248_[int(8)] = d_1342_v88_
            nw248_[int(9)] = d_1342_v88_
            nw248_[int(10)] = d_1342_v88_
            nw248_[int(11)] = d_1342_v88_
            nw248_[int(12)] = d_1342_v88_
            nw248_[int(13)] = d_1342_v88_
            nw248_[int(14)] = d_1342_v88_
            nw248_[int(15)] = d_1342_v88_
            nw248_[int(16)] = d_1342_v88_
            nw248_[int(17)] = d_1342_v88_
            nw248_[int(18)] = (d_1343_v89_)[default__.safeIndex((self).f39, len(d_1343_v89_))]
            nw248_[int(19)] = d_1342_v88_
            nw248_[int(20)] = d_1342_v88_
            nw248_[int(21)] = d_1342_v88_
            nw248_[int(22)] = d_1342_v88_
            nw248_[int(23)] = d_1342_v88_
            d_1344_v90_ = nw248_
            d_1344_v90_ = d_1344_v90_
        elif True:
            (globalState).f1 = (0) - ((self).f39)
            r0 = (self).f39
            d_1345_v91_: D7
            d_1345_v91_ = D7_DC18(p1)
            d_1346_v92_: _dafny.Seq
            d_1346_v92_ = _dafny.SeqWithoutIsStrInference([(self).f40, (d_1345_v91_).cf23])
            rhs228_ = d_1346_v92_
            rhs229_ = len((p2) + (p2))
            lhs197_ = globalState
            d_1346_v92_ = rhs228_
            lhs197_.f1 = rhs229_
            d_1347_v93_: _dafny.MultiSet
            d_1347_v93_ = _dafny.MultiSet([((self).f39) + ((self).f39), (self).f39, (self).f39, 766, default__.safeDivisionInt(340, 281)])
            (globalState).f24 = (d_1347_v93_).cardinality
            if False:
                d_1348_v94_: _dafny.Set
                d_1348_v94_ = _dafny.Set({self.f28})
                d_1349_v95_: _dafny.Array
                nw249_ = _dafny.Array(None, 12)
                nw249_[int(0)] = p0
                nw249_[int(1)] = self.f28
                nw249_[int(2)] = (self).f40
                nw249_[int(3)] = ((self).f40) and (False)
                nw249_[int(4)] = p0
                nw249_[int(5)] = (self).fm32(_dafny.Set({self.f28, (self).f29}), (self).f39, p0, globalState)
                nw249_[int(6)] = p0
                nw249_[int(7)] = (not(not(p1)) if not((self).f29) else not((self).f40))
                nw249_[int(8)] = not((self).fm32(d_1348_v94_, 537, p0, globalState))
                nw249_[int(9)] = False
                nw249_[int(10)] = False
                nw249_[int(11)] = (self).f29
                d_1349_v95_ = nw249_
                index204_ = default__.safeIndex(668, (d_1349_v95_).length(0))
                (d_1349_v95_)[index204_] = (self).f40
                index205_ = default__.safeIndex(668, (d_1349_v95_).length(0))
                (d_1349_v95_)[index205_] = not ((self).f29) or ((self).f29)
                d_1350_v96_: _dafny.Seq
                d_1350_v96_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1351_i7_ in range(default__.abs(452))]))), (self).f39])
                d_1352_v97_: _dafny.Map
                d_1352_v97_ = _dafny.Map({((self).f39) > (len(_dafny.Map({(self).f39: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1353_i6_ in range(default__.abs(852))]))}))): (d_1350_v96_) + (d_1350_v96_)})
                d_1354_v98_: _dafny.MultiSet
                d_1354_v98_ = _dafny.MultiSet([self.f28])
                d_1355_v99_: _dafny.Array
                nw250_ = _dafny.Array(None, 17)
                nw250_[int(0)] = (self).f39
                nw250_[int(1)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1356_i8_ in range(default__.abs(106))]))
                nw250_[int(2)] = (self).f39
                nw250_[int(3)] = (self).f39
                nw250_[int(4)] = (self).f39
                nw250_[int(5)] = (self).f39
                nw250_[int(6)] = (self).f39
                nw250_[int(7)] = (self).f39
                nw250_[int(8)] = (self).f39
                nw250_[int(9)] = (self).f39
                nw250_[int(10)] = (self).f39
                nw250_[int(11)] = (d_1354_v98_).cardinality
                nw250_[int(12)] = len(d_1346_v92_)
                nw250_[int(13)] = (0) - ((self).f39)
                nw250_[int(14)] = (self).f39
                nw250_[int(15)] = 213
                nw250_[int(16)] = (self).f39
                d_1355_v99_ = nw250_
                d_1357_v100_: _dafny.Map
                d_1357_v100_ = _dafny.Map({d_1355_v99_: d_1352_v97_})
                d_1352_v97_ = ((d_1357_v100_)[d_1355_v99_] if (d_1355_v99_) in (d_1357_v100_) else d_1352_v97_)
                d_1358_v101_: _dafny.Set
                d_1358_v101_ = _dafny.Set({(d_1354_v98_).cardinality, (self).f39})
                (globalState).f24 = (self).fm30(d_1358_v101_, globalState)
                d_1359_v102_: _dafny.Array
                nw251_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1359_v102_ = nw251_
                d_1360_v103_: _dafny.Array
                nw252_ = _dafny.Array(_dafny.Map({}), 25)
                d_1360_v103_ = nw252_
                index206_ = default__.safeIndex(398, (d_1359_v102_).length(0))
                (d_1359_v102_)[index206_] = d_1360_v103_
                index207_ = default__.safeIndex(398, (d_1359_v102_).length(0))
                (d_1359_v102_)[index207_] = d_1360_v103_
                index208_ = default__.safeIndex(612, (d_1355_v99_).length(0))
                (d_1355_v99_)[index208_] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f39) for d_1361_i9_ in range(default__.abs(223))])), (self).f39)
                index209_ = default__.safeIndex(612, (d_1355_v99_).length(0))
                (d_1355_v99_)[index209_] = (self).f39
            elif True:
                d_1362_v104_: _dafny.Map
                d_1362_v104_ = _dafny.Map({(self).f39: (p0) and (self.f28)})
                d_1363_v105_: _dafny.Array
                def lambda32_(d_1364_p0_):
                    def lambda33_(d_1365_i10_):
                        return d_1364_p0_

                    return lambda33_

                init21_ = lambda32_(p0)
                nw253_ = _dafny.Array(None, 1)
                for i0_21_ in range(nw253_.length(0)):
                    nw253_[i0_21_] = init21_(i0_21_)
                d_1363_v105_ = nw253_
                d_1366_v106_: _dafny.Seq
                d_1366_v106_ = _dafny.SeqWithoutIsStrInference([d_1363_v105_, d_1363_v105_, d_1363_v105_, d_1363_v105_])
                d_1367_v107_: _dafny.Set
                d_1367_v107_ = _dafny.Set({(self).f39})
                d_1368_v108_: _dafny.Set
                d_1368_v108_ = _dafny.Set({(d_1347_v93_).set((self).fm30(_dafny.Set({len(d_1366_v106_), (self).fm30(d_1367_v107_, globalState)}), globalState), default__.abs(len(_dafny.SeqWithoutIsStrInference([382 for d_1369_i11_ in range(default__.abs(-164))]))))})
                (globalState).f21 = ((d_1362_v104_)[((self).f39) * ((self).f39)] if (((self).f39) * ((self).f39)) in (d_1362_v104_) else (_dafny.Set({d_1347_v93_, _dafny.MultiSet([(self).f39]), d_1347_v93_})).ispropersubset(d_1368_v108_))
                d_1370_v109_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Map({}), 22)
                d_1370_v109_ = nw254_
                d_1371_v110_: _dafny.Seq
                d_1371_v110_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1346_v92_))])
                d_1372_v111_: _dafny.Map
                d_1372_v111_ = _dafny.Map({d_1370_v109_: ((self).f39) * ((d_1371_v110_)[default__.safeIndex((self).f39, len(d_1371_v110_))])})
                d_1373_v112_: _dafny.Seq
                d_1373_v112_ = _dafny.SeqWithoutIsStrInference([d_1370_v109_])
                d_1372_v111_ = (d_1372_v111_).set((d_1373_v112_)[default__.safeIndex((self).f39, len(d_1373_v112_))], (self).f39)
                d_1374_v113_: _dafny.Array
                nw255_ = _dafny.Array(int(0), 21)
                d_1374_v113_ = nw255_
                d_1375_v115_: _dafny.MultiSet
                def iife108_():
                    coll58_ = _dafny.Set()
                    compr_58_: int
                    for compr_58_ in _dafny.IntegerRange(123, 232):
                        d_1376_v114_: int = compr_58_
                        if ((123) <= (d_1376_v114_)) and ((d_1376_v114_) < (232)):
                            coll58_ = coll58_.union(_dafny.Set([(d_1376_v114_) + (len(p2))]))
                    return _dafny.Set(coll58_)
                d_1375_v115_ = _dafny.MultiSet([default__.fm35(globalState), iife108_()
                ])
                index210_ = default__.safeIndex(548, (d_1374_v113_).length(0))
                (d_1374_v113_)[index210_] = ((d_1375_v115_)[d_1367_v107_] if (d_1367_v107_) in (d_1375_v115_) else (self).f39)
                index211_ = default__.safeIndex(548, (d_1374_v113_).length(0))
                (d_1374_v113_)[index211_] = (self).f39
                d_1377_v116_: _dafny.MultiSet
                d_1377_v116_ = _dafny.MultiSet([p0, not(p1), p1, p1, (self).f40])
                d_1378_v117_: _dafny.Array
                nw256_ = _dafny.Array(None, 1)
                nw256_[int(0)] = (d_1377_v116_) - ((d_1377_v116_).set((self).f29, default__.abs((d_1374_v113_)[default__.safeIndex(548, (d_1374_v113_).length(0))])))
                d_1378_v117_ = nw256_
                index212_ = default__.safeIndex(310, (d_1378_v117_).length(0))
                (d_1378_v117_)[index212_] = d_1377_v116_
                index213_ = default__.safeIndex(310, (d_1378_v117_).length(0))
                (d_1378_v117_)[index213_] = ((d_1377_v116_) - (d_1377_v116_)).intersection(d_1377_v116_)
                (globalState).f27 = not((self).f40)
        r2 = p1
        d_1379_v118_: _dafny.Array
        def lambda34_(d_1380_i12_):
            return default__.safeDivisionInt(d_1380_i12_, (self).f39)

        init22_ = lambda34_
        nw257_ = _dafny.Array(None, 25)
        for i0_22_ in range(nw257_.length(0)):
            nw257_[i0_22_] = init22_(i0_22_)
        d_1379_v118_ = nw257_
        index214_ = default__.safeIndex(651, (d_1379_v118_).length(0))
        (d_1379_v118_)[index214_] = (self).f39
        index215_ = default__.safeIndex(651, (d_1379_v118_).length(0))
        (d_1379_v118_)[index215_] = 221
        r0 = (self).f39
        d_1381_v119_: str
        d_1381_v119_ = _dafny.CodePoint('x')
        r1 = d_1381_v119_
        d_1382_v120_: _dafny.Seq
        d_1382_v120_ = _dafny.SeqWithoutIsStrInference([(d_1379_v118_)[default__.safeIndex(651, (d_1379_v118_).length(0))], (d_1379_v118_)[default__.safeIndex(651, (d_1379_v118_).length(0))], (self).f39, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixtb"))), 576])
        def iife109_():
            coll59_ = _dafny.Set()
            compr_59_: int
            for compr_59_ in (d_1382_v120_).Elements:
                d_1383_v121_: int = compr_59_
                if (d_1383_v121_) in (d_1382_v120_):
                    coll59_ = coll59_.union(_dafny.Set([default__.safeModuloInt(d_1383_v121_, len(_dafny.SeqWithoutIsStrInference([False])))]))
            return _dafny.Set(coll59_)
        r2 = (D9_DC25(self.f28, (self).f39, (self).fm30(iife109_()
, globalState), p0, (0) - ((d_1379_v118_)[default__.safeIndex(651, (d_1379_v118_).length(0))]))).cf37
        return r0, r1, r2

    def m17(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_1384_v0_: _dafny.Seq
        d_1384_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvg"))
        d_1384_v0_ = (d_1384_v0_) + ((d_1384_v0_) + (d_1384_v0_))
        d_1385_v1_: _dafny.Array
        def lambda35_(d_1386_i1_):
            return (d_1386_i1_) * ((self).f39)

        init23_ = lambda35_
        nw258_ = _dafny.Array(None, 27)
        for i0_23_ in range(nw258_.length(0)):
            nw258_[i0_23_] = init23_(i0_23_)
        d_1385_v1_ = nw258_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1385_v1_).length(0)):
            d_1387_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1387_i0_)) and ((d_1387_i0_) < ((d_1385_v1_).length(0)))):
                (d_1385_v1_)[(d_1387_i0_)] = (d_1387_i0_) + ((0) - ((self).f39))
        d_1388_v2_: _dafny.Map
        d_1388_v2_ = _dafny.Map({p2: p0})
        d_1389_v3_: _dafny.Seq
        d_1389_v3_ = _dafny.SeqWithoutIsStrInference([d_1388_v2_])
        d_1390_v4_: _dafny.Set
        d_1390_v4_ = _dafny.Set({(self).f39, (self).f39})
        d_1391_v5_: D12
        d_1391_v5_ = D12_DC33(D12_DC31(d_1390_v4_))
        pat_let_tv34_ = d_1389_v3_
        pat_let_tv35_ = d_1388_v2_
        pat_let_tv36_ = d_1384_v0_
        pat_let_tv37_ = d_1389_v3_
        pat_let_tv38_ = d_1388_v2_
        pat_let_tv39_ = d_1389_v3_
        pat_let_tv40_ = d_1388_v2_
        pat_let_tv41_ = d_1384_v0_
        pat_let_tv42_ = d_1389_v3_
        pat_let_tv43_ = d_1388_v2_
        pat_let_tv44_ = d_1388_v2_
        pat_let_tv45_ = d_1389_v3_
        pat_let_tv46_ = d_1388_v2_
        pat_let_tv47_ = d_1388_v2_
        pat_let_tv48_ = d_1388_v2_
        def lambda36_(source24_):
            if source24_.is_DC32:
                return (((pat_let_tv34_) + (_dafny.SeqWithoutIsStrInference([pat_let_tv35_ for d_1392_i2_ in range(default__.abs(204))]))).set(default__.safeIndex(len(pat_let_tv36_), len((pat_let_tv37_) + (_dafny.SeqWithoutIsStrInference([pat_let_tv38_ for d_1393_i2_ in range(default__.abs(204))])))), _dafny.Map({self.f28: (self).f29}))).set(default__.safeIndex((self).f39, len(((pat_let_tv39_) + (_dafny.SeqWithoutIsStrInference([pat_let_tv40_ for d_1394_i2_ in range(default__.abs(204))]))).set(default__.safeIndex(len(pat_let_tv41_), len((pat_let_tv42_) + (_dafny.SeqWithoutIsStrInference([pat_let_tv43_ for d_1395_i2_ in range(default__.abs(204))])))), _dafny.Map({self.f28: (self).f29})))), pat_let_tv44_)
            elif source24_.is_DC31:
                d_1396___mcc_h0_ = source24_.cf48
                d_1397_cf48_ = d_1396___mcc_h0_
                return pat_let_tv45_
            elif True:
                d_1398___mcc_h1_ = source24_.cf49
                d_1399_cf49_ = d_1398___mcc_h1_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv46_, pat_let_tv47_, pat_let_tv48_, _dafny.Map({self.f28: (self).f40})])

        rhs230_ = not((self).f29)
        rhs231_ = 901
        rhs232_ = lambda36_(d_1391_v5_)
        rhs233_ = default__.safeDivisionInt((self).f39, (self).f39)
        lhs198_ = globalState
        lhs199_ = globalState
        lhs198_.f22 = rhs230_
        r1 = rhs231_
        d_1389_v3_ = rhs232_
        lhs199_.f1 = rhs233_
        d_1400_v6_: D7
        d_1400_v6_ = D7_DC17(self.f28)
        d_1401_i3_: int
        d_1401_i3_ = 0
        with _dafny.label("1"):
            while (d_1400_v6_).cf22:
                with _dafny.c_label("1"):
                    if (d_1401_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_1401_i3_ = (d_1401_i3_) + (1)
                    d_1402_v7_: C2
                    nw259_ = C2()
                    nw259_.ctor__()
                    d_1402_v7_ = nw259_
                    d_1403_v8_: D7
                    d_1403_v8_ = D7_DC18(p2)
                    d_1404_v9_: _dafny.Map
                    d_1404_v9_ = _dafny.Map({p0: -692})
                    d_1405_v10_: str
                    d_1405_v10_ = _dafny.CodePoint('c')
                    d_1406_v11_: _dafny.Map
                    d_1406_v11_ = _dafny.Map({len(d_1404_v9_): d_1405_v10_})
                    d_1407_v12_: _dafny.MultiSet
                    d_1407_v12_ = _dafny.MultiSet([d_1390_v4_])
                    rhs234_ = (d_1384_v0_).set(default__.safeIndex(-985, len(d_1384_v0_)), ((d_1406_v11_)[len(_dafny.SeqWithoutIsStrInference([d_1405_v10_ for d_1408_i4_ in range(default__.abs(788))]))] if (len(_dafny.SeqWithoutIsStrInference([d_1405_v10_ for d_1409_i4_ in range(default__.abs(788))]))) in (d_1406_v11_) else d_1405_v10_))
                    rhs235_ = (default__.safeModuloInt((0) - ((self).f39), (self).f39)) - ((self).f39)
                    rhs236_ = (d_1407_v12_).isdisjoint(d_1407_v12_)
                    rhs237_ = d_1403_v8_
                    rhs238_ = (not(False)) == ((self).f29)
                    lhs200_ = globalState
                    lhs201_ = globalState
                    d_1384_v0_ = rhs234_
                    r2 = rhs235_
                    lhs200_.f21 = rhs236_
                    d_1403_v8_ = rhs237_
                    lhs201_.f27 = rhs238_
                    d_1410_v13_: _dafny.Array
                    def lambda37_(d_1411_v0_):
                        def lambda38_(d_1412_i5_):
                            return d_1411_v0_

                        return lambda38_

                    init24_ = lambda37_(d_1384_v0_)
                    nw260_ = _dafny.Array(None, 14)
                    for i0_24_ in range(nw260_.length(0)):
                        nw260_[i0_24_] = init24_(i0_24_)
                    d_1410_v13_ = nw260_
                    d_1413_v14_: _dafny.Map
                    d_1413_v14_ = _dafny.Map({d_1410_v13_: d_1384_v0_})
                    d_1414_v15_: _dafny.Map
                    d_1414_v15_ = _dafny.Map({(self).f29: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))})
                    d_1415_v16_: _dafny.Map
                    d_1415_v16_ = _dafny.Map({(self).f39: ((d_1413_v14_)[d_1410_v13_] if (d_1410_v13_) in (d_1413_v14_) else ((d_1414_v15_)[p2] if (p2) in (d_1414_v15_) else d_1384_v0_))})
                    (globalState).f1 = len(d_1415_v16_)
                    if p2:
                        d_1416_v17_: _dafny.Map
                        d_1416_v17_ = _dafny.Map({d_1405_v10_: (d_1405_v10_) in (d_1384_v0_)})
                        d_1417_v18_: _dafny.Array
                        def lambda39_(d_1418_v10_):
                            def lambda40_(d_1419_i6_):
                                return d_1418_v10_

                            return lambda40_

                        init25_ = lambda39_(d_1405_v10_)
                        nw261_ = _dafny.Array(None, 23)
                        for i0_25_ in range(nw261_.length(0)):
                            nw261_[i0_25_] = init25_(i0_25_)
                        d_1417_v18_ = nw261_
                        d_1420_v20_: _dafny.Array
                        d_1420_v20_ = d_1417_v18_
                        def iife110_():
                            coll60_ = _dafny.Map()
                            compr_60_: str
                            for compr_60_ in (((d_1384_v0_).set(default__.safeIndex((0) - ((self).f39), len(d_1384_v0_)), d_1405_v10_)) + (d_1384_v0_)).Elements:
                                d_1421_v19_: str = compr_60_
                                if (d_1421_v19_) in (((d_1384_v0_).set(default__.safeIndex((0) - ((self).f39), len(d_1384_v0_)), d_1405_v10_)) + (d_1384_v0_)):
                                    coll60_[d_1421_v19_] = (self).f40
                            return _dafny.Map(coll60_)
                        rhs239_ = p0
                        rhs240_ = iife110_()

                        rhs241_ = (d_1420_v20_)
                        rhs242_ = d_1405_v10_
                        lhs202_ = globalState
                        lhs203_ = globalState
                        lhs202_.f22 = rhs239_
                        d_1416_v17_ = rhs240_
                        d_1417_v18_ = rhs241_
                        lhs203_.f23 = rhs242_
                        d_1422_v21_: C0
                        nw262_ = C0()
                        nw262_.ctor__()
                        d_1422_v21_ = nw262_
                        d_1423_v22_: _dafny.Map
                        d_1423_v22_ = _dafny.Map({(self).f39: not((self).f29)})
                        d_1424_v23_: _dafny.MultiSet
                        d_1424_v23_ = _dafny.MultiSet([((0) - (-411)) * (len(d_1423_v22_)), default__.safeModuloInt((0) - ((self).f39), -690), (self).f39, 88, (self).fm30(d_1390_v4_, globalState)])
                        d_1424_v23_ = (d_1424_v23_) - (_dafny.MultiSet([len(d_1388_v2_), (self).f39]))
                        d_1425_v24_: _dafny.Array
                        nw263_ = _dafny.Array(_dafny.Array(None, 0), 6)
                        d_1425_v24_ = nw263_
                        d_1426_v25_: _dafny.Seq
                        d_1426_v25_ = _dafny.SeqWithoutIsStrInference([(self).f39, -834, len(_dafny.SeqWithoutIsStrInference([d_1405_v10_ for d_1427_i7_ in range(default__.abs(-310))]))])
                        d_1428_v26_: _dafny.Array
                        nw264_ = _dafny.Array(None, 7)
                        nw264_[int(0)] = _dafny.SeqWithoutIsStrInference([(self).f39])
                        nw264_[int(1)] = d_1426_v25_
                        nw264_[int(2)] = (d_1426_v25_).set(default__.safeIndex((self).f39, len(d_1426_v25_)), (self).f39)
                        nw264_[int(3)] = d_1426_v25_
                        nw264_[int(4)] = d_1426_v25_
                        nw264_[int(5)] = d_1426_v25_
                        nw264_[int(6)] = d_1426_v25_
                        d_1428_v26_ = nw264_
                        index216_ = default__.safeIndex(701, (d_1425_v24_).length(0))
                        (d_1425_v24_)[index216_] = d_1428_v26_
                        index217_ = default__.safeIndex(701, (d_1425_v24_).length(0))
                        (d_1425_v24_)[index217_] = d_1428_v26_
                        d_1423_v22_ = (d_1423_v22_).set((0) - ((self).f39), (d_1422_v21_).fm33(d_1384_v0_, d_1415_v16_, (d_1402_v7_).fm52(d_1384_v0_, globalState), globalState))
                    elif True:
                        r3 = p0
                        d_1429_v28_: _dafny.Array
                        def lambda41_(d_1430_v4_):
                            def lambda42_(d_1431_i8_):
                                def iife111_():
                                    coll61_ = _dafny.Set()
                                    compr_61_: int
                                    for compr_61_ in _dafny.IntegerRange(867, 782):
                                        d_1432_v27_: int = compr_61_
                                        if ((867) <= (d_1432_v27_)) and ((d_1432_v27_) < (782)):
                                            coll61_ = coll61_.union(_dafny.Set([(d_1432_v27_) * ((self).f39)]))
                                    return _dafny.Set(coll61_)
                                return (d_1430_v4_).intersection(iife111_()
                                )

                            return lambda42_

                        init26_ = lambda41_(d_1390_v4_)
                        nw265_ = _dafny.Array(None, 12)
                        for i0_26_ in range(nw265_.length(0)):
                            nw265_[i0_26_] = init26_(i0_26_)
                        d_1429_v28_ = nw265_
                        d_1433_v29_: _dafny.Seq
                        d_1433_v29_ = _dafny.SeqWithoutIsStrInference([d_1390_v4_, d_1390_v4_])
                        d_1434_v30_: _dafny.MultiSet
                        d_1434_v30_ = _dafny.MultiSet([(self).f39])
                        index218_ = default__.safeIndex(994, (d_1429_v28_).length(0))
                        def iife112_():
                            coll62_ = _dafny.Set()
                            compr_62_: int
                            for compr_62_ in (d_1434_v30_).Elements:
                                d_1435_v31_: int = compr_62_
                                if (d_1435_v31_) in (d_1434_v30_):
                                    coll62_ = coll62_.union(_dafny.Set([(d_1435_v31_) - (537)]))
                            return _dafny.Set(coll62_)
                        (d_1429_v28_)[index218_] = ((d_1433_v29_)[default__.safeIndex((self).f39, len(d_1433_v29_))]).intersection(iife112_()
                        )
                        index219_ = default__.safeIndex(994, (d_1429_v28_).length(0))
                        (d_1429_v28_)[index219_] = d_1390_v4_
                        index220_ = default__.safeIndex(557, (d_1385_v1_).length(0))
                        (d_1385_v1_)[index220_] = (0) - ((self).f39)
                        index221_ = default__.safeIndex(557, (d_1385_v1_).length(0))
                        (d_1385_v1_)[index221_] = (self).f39
                        index222_ = default__.safeIndex(557, (d_1385_v1_).length(0))
                        (d_1385_v1_)[index222_] = (d_1385_v1_)[default__.safeIndex(557, (d_1385_v1_).length(0))]
                        (globalState).f22 = self.f28
                    pass
            pass
        pat_let_tv49_ = d_1390_v4_
        def lambda43_(source25_):
            if source25_.is_DC5:
                d_1436___mcc_h2_ = source25_.cf8
                d_1437_cf8_ = d_1436___mcc_h2_
                return (self).f39
            elif source25_.is_DC4:
                d_1438___mcc_h3_ = source25_.cf7
                d_1439_cf7_ = d_1438___mcc_h3_
                return (951) - ((self).f39)
            elif True:
                d_1440___mcc_h4_ = source25_.cf9
                d_1441_cf9_ = d_1440___mcc_h4_
                return len(pat_let_tv49_)

        (globalState).f1 = lambda43_(D2_DC5(p0))
        d_1442_v32_: _dafny.Set
        d_1442_v32_ = _dafny.Set({p0, self.f28, self.f28, default__.fm38(globalState)})
        (globalState).f1 = len(d_1442_v32_)
        d_1443_v33_: _dafny.Map
        d_1443_v33_ = _dafny.Map({(self).f40: d_1442_v32_})
        r0 = default__.safeModuloInt((self).f39, (0) - (len(d_1443_v33_)))
        r1 = (0) - ((self).f39)
        d_1444_v34_: _dafny.Map
        d_1444_v34_ = _dafny.Map({False: (self).f39})
        r2 = default__.safeDivisionInt((self).f39, (((d_1444_v34_)[p2] if (p2) in (d_1444_v34_) else (self).f39)) - ((self).f39))
        pat_let_tv50_ = p2
        pat_let_tv51_ = d_1444_v34_
        pat_let_tv52_ = d_1444_v34_
        def lambda44_(source26_):
            if source26_.is_DC35:
                return pat_let_tv50_
            elif True:
                d_1445___mcc_h5_ = source26_.cf50
                d_1446_cf50_ = d_1445___mcc_h5_
                return (pat_let_tv51_) == (pat_let_tv52_)

        r3 = lambda44_(D13_DC35())
        return r0, r1, r2, r3

    def m18(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: int = int(0)
        d_1447_v0_: _dafny.Set
        d_1447_v0_ = _dafny.Set({self.f28})
        d_1448_v1_: _dafny.Map
        d_1448_v1_ = _dafny.Map({not((self).f40): (self).f39})
        d_1449_v2_: _dafny.Array
        nw266_ = _dafny.Array(None, 7)
        nw266_[int(0)] = self.f28
        nw266_[int(1)] = (self).f29
        nw266_[int(2)] = (self).fm32(d_1447_v0_, (0) - (len(p1)), (self).fm31(globalState), globalState)
        nw266_[int(3)] = True
        nw266_[int(4)] = ((self).f39) > (745)
        nw266_[int(5)] = (self).f40
        nw266_[int(6)] = ((self).f40) not in (d_1448_v1_)
        d_1449_v2_ = nw266_
        index223_ = default__.safeIndex(665, (d_1449_v2_).length(0))
        (d_1449_v2_)[index223_] = (self).f29
        index224_ = default__.safeIndex(665, (d_1449_v2_).length(0))
        (d_1449_v2_)[index224_] = (self).f40
        d_1449_v2_ = (d_1449_v2_ if ((self).f29) or ((self).f29) else d_1449_v2_)
        index225_ = default__.safeIndex(665, (d_1449_v2_).length(0))
        (d_1449_v2_)[index225_] = not (self.f28) or (default__.fm38(globalState))
        index226_ = default__.safeIndex(665, (d_1449_v2_).length(0))
        (d_1449_v2_)[index226_] = (self).f29
        d_1450_v3_: _dafny.MultiSet
        d_1450_v3_ = _dafny.MultiSet([(self).f39])
        d_1451_v4_: _dafny.Seq
        d_1451_v4_ = _dafny.SeqWithoutIsStrInference([211])
        index227_ = default__.safeIndex(665, (d_1449_v2_).length(0))
        rhs243_ = ((self).f39) == ((self).f39)
        rhs244_ = (default__.safeModuloInt((d_1450_v3_).cardinality, len(d_1451_v4_)) if (d_1449_v2_)[default__.safeIndex(665, (d_1449_v2_).length(0))] else (self).f39)
        rhs245_ = default__.fm51((self).f39, (self).f39, globalState)
        rhs246_ = p0
        lhs204_ = d_1449_v2_
        lhs205_ = default__.safeIndex(665, (d_1449_v2_).length(0))
        lhs206_ = globalState
        lhs204_[lhs205_] = rhs243_
        r0 = rhs244_
        r2 = rhs245_
        lhs206_.f23 = rhs246_
        d_1452_v5_: D10
        d_1452_v5_ = D10_DC28(default__.fm59((d_1449_v2_)[default__.safeIndex(665, (d_1449_v2_).length(0))], _dafny.SeqWithoutIsStrInference([(self).f40]), globalState))
        d_1453_v6_: D10
        d_1453_v6_ = D10_DC29(d_1452_v5_)
        d_1454_v7_: _dafny.Map
        d_1454_v7_ = _dafny.Map({d_1453_v6_: d_1449_v2_})
        d_1454_v7_ = d_1454_v7_
        r0 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnsunvi")))) - ((self).f39)
        r1 = (self).fm31(globalState)
        r2 = p1
        d_1455_v8_: _dafny.Seq
        d_1455_v8_ = _dafny.SeqWithoutIsStrInference([(d_1449_v2_)[default__.safeIndex(665, (d_1449_v2_).length(0))], self.f28])
        d_1456_v9_: _dafny.Map
        d_1456_v9_ = _dafny.Map({774: True})
        d_1457_v10_: D10
        d_1457_v10_ = D10_DC27((d_1455_v8_)[default__.safeIndex((self).f39, len(d_1455_v8_))], (self).f29, self.f28, len(d_1456_v9_), False)
        r3 = (d_1457_v10_).cf43
        return r0, r1, r2, r3

    @property
    def f40(self):
        return self._f40

class C4(T0, T1):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f37: bool = False
        self._f38: bool = False
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
    def ctor__(self, f37, f38, f28, f29):
        (self)._f37 = f37
        (self)._f38 = f38
        (self).f28 = f28
        (self)._f29 = f29

    def m5(self, p0, globalState):
        r0: bool = False
        d_1458_v0_: _dafny.Array
        nw267_ = _dafny.Array(None, 10)
        nw267_[int(0)] = self.f28
        nw267_[int(1)] = not((self).f37)
        nw267_[int(2)] = (self).f38
        nw267_[int(3)] = (self).f37
        nw267_[int(4)] = self.f28
        nw267_[int(5)] = (self).f37
        nw267_[int(6)] = True
        nw267_[int(7)] = False
        nw267_[int(8)] = self.f28
        nw267_[int(9)] = (self).f38
        d_1458_v0_ = nw267_
        d_1459_v1_: D5
        d_1459_v1_ = D5_DC12(d_1458_v0_)
        d_1460_v2_: _dafny.Map
        d_1460_v2_ = _dafny.Map({(self).f38: d_1458_v0_})
        d_1461_v3_: _dafny.Array
        nw268_ = _dafny.Array(None, 29)
        nw268_[int(0)] = d_1458_v0_
        nw268_[int(1)] = d_1458_v0_
        nw268_[int(2)] = d_1458_v0_
        nw268_[int(3)] = d_1458_v0_
        nw268_[int(4)] = d_1458_v0_
        nw268_[int(5)] = (d_1459_v1_).cf17
        nw268_[int(6)] = ((d_1460_v2_)[not((self).f38)] if (not((self).f38)) in (d_1460_v2_) else d_1458_v0_)
        nw268_[int(7)] = d_1458_v0_
        nw268_[int(8)] = d_1458_v0_
        nw268_[int(9)] = d_1458_v0_
        nw268_[int(10)] = d_1458_v0_
        nw268_[int(11)] = d_1458_v0_
        nw268_[int(12)] = d_1458_v0_
        nw268_[int(13)] = d_1458_v0_
        nw268_[int(14)] = d_1458_v0_
        nw268_[int(15)] = d_1458_v0_
        nw268_[int(16)] = d_1458_v0_
        nw268_[int(17)] = d_1458_v0_
        nw268_[int(18)] = d_1458_v0_
        nw268_[int(19)] = d_1458_v0_
        nw268_[int(20)] = d_1458_v0_
        nw268_[int(21)] = d_1458_v0_
        nw268_[int(22)] = d_1458_v0_
        nw268_[int(23)] = d_1458_v0_
        nw268_[int(24)] = d_1458_v0_
        nw268_[int(25)] = d_1458_v0_
        nw268_[int(26)] = d_1458_v0_
        nw268_[int(27)] = d_1458_v0_
        nw268_[int(28)] = (d_1458_v0_ if (self).f29 else d_1458_v0_)
        d_1461_v3_ = nw268_
        index228_ = default__.safeIndex(389, (d_1461_v3_).length(0))
        (d_1461_v3_)[index228_] = d_1458_v0_
        index229_ = default__.safeIndex(389, (d_1461_v3_).length(0))
        (d_1461_v3_)[index229_] = d_1458_v0_
        hi4_ = p0
        for d_1462_i0_ in range(p0, hi4_):
            d_1463_v4_: str
            d_1463_v4_ = _dafny.CodePoint('h')
            d_1464_v5_: _dafny.Seq
            d_1464_v5_ = _dafny.SeqWithoutIsStrInference([(d_1462_i0_) + (p0), -764, default__.fm27(927, d_1463_v4_, p0, globalState)])
            (globalState).f17 = d_1464_v5_
            d_1465_v6_: _dafny.Seq
            d_1465_v6_ = _dafny.SeqWithoutIsStrInference([self.f28])
            d_1466_v7_: _dafny.Map
            d_1466_v7_ = _dafny.Map({d_1463_v4_: (d_1465_v6_) + (d_1465_v6_)})
            d_1467_v9_: _dafny.Map
            d_1467_v9_ = _dafny.Map({d_1462_i0_: (self).f38})
            def iife113_():
                coll63_ = _dafny.Map()
                compr_63_: str
                for compr_63_ in (_dafny.Map({d_1463_v4_: len(d_1467_v9_)})).keys.Elements:
                    d_1468_v8_: str = compr_63_
                    if (d_1468_v8_) in (_dafny.Map({d_1463_v4_: len(d_1467_v9_)})):
                        coll63_[d_1468_v8_] = _dafny.SeqWithoutIsStrInference([(self).f38])
                return _dafny.Map(coll63_)
            d_1466_v7_ = (iife113_()
             if self.f28 else (d_1466_v7_) | (_dafny.Map({d_1463_v4_: d_1465_v6_})))
            d_1469_v10_: _dafny.Seq
            d_1469_v10_ = _dafny.SeqWithoutIsStrInference([d_1458_v0_, d_1458_v0_, d_1458_v0_, d_1458_v0_])
            d_1470_v11_: _dafny.MultiSet
            d_1470_v11_ = _dafny.MultiSet([True])
            d_1471_v12_: _dafny.Map
            d_1471_v12_ = _dafny.Map({(self).f37: d_1470_v11_})
            d_1472_v13_: _dafny.Seq
            d_1472_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qr"))
            d_1473_v14_: _dafny.Array
            def lambda45_(d_1474_i0_):
                def lambda46_(d_1475_i1_):
                    return default__.safeModuloInt(d_1475_i1_, d_1474_i0_)

                return lambda46_

            init27_ = lambda45_(d_1462_i0_)
            nw269_ = _dafny.Array(None, 23)
            for i0_27_ in range(nw269_.length(0)):
                nw269_[i0_27_] = init27_(i0_27_)
            d_1473_v14_ = nw269_
            d_1476_v15_: _dafny.Map
            d_1476_v15_ = _dafny.Map({d_1472_v13_: (self).f38})
            d_1477_v17_: D4
            d_1477_v17_ = D4_DC10((self).f29, p0, (0) - (p0))
            d_1478_v18_: _dafny.Map
            d_1478_v18_ = _dafny.Map({self.f28: (self).f29})
            pat_let_tv53_ = p0
            d_1479_v19_: _dafny.Array
            nw270_ = _dafny.Array(None, 25)
            nw270_[int(0)] = d_1462_i0_
            nw270_[int(1)] = d_1462_i0_
            nw270_[int(2)] = p0
            nw270_[int(3)] = len((d_1471_v12_) | (d_1471_v12_))
            nw270_[int(4)] = (p0) * (d_1462_i0_)
            nw270_[int(5)] = (len(d_1472_v13_)) + (d_1462_i0_)
            nw270_[int(6)] = d_1462_i0_
            nw270_[int(7)] = (472) + (len((_dafny.Map({-880: d_1473_v14_})).set(954, d_1473_v14_)))
            def iife114_():
                coll64_ = _dafny.Map()
                compr_64_: _dafny.Seq
                for compr_64_ in (_dafny.SeqWithoutIsStrInference([d_1472_v13_ for d_1480_i2_ in range(default__.abs(152))])).Elements:
                    d_1481_v16_: _dafny.Seq = compr_64_
                    if (d_1481_v16_) in (_dafny.SeqWithoutIsStrInference([d_1472_v13_ for d_1480_i2_ in range(default__.abs(152))])):
                        coll64_[d_1481_v16_] = self.f28
                return _dafny.Map(coll64_)
            nw270_[int(8)] = len((d_1476_v15_) | (iife114_()
            ))
            nw270_[int(9)] = d_1462_i0_
            nw270_[int(10)] = (0) - (d_1462_i0_)
            nw270_[int(11)] = len(d_1465_v6_)
            def iife115_(_pat_let25_0):
                def iife116_(d_1482_dt__update__tmp_h0_):
                    def iife117_(_pat_let26_0):
                        def iife118_(d_1483_dt__update_hcf15_h0_):
                            return D4_DC10((d_1482_dt__update__tmp_h0_).cf13, (d_1482_dt__update__tmp_h0_).cf14, d_1483_dt__update_hcf15_h0_)
                        return iife118_(_pat_let26_0)
                    return iife117_(pat_let_tv53_)
                return iife116_(_pat_let25_0)
            nw270_[int(12)] = len(default__.fm28(iife115_(d_1477_v17_), p0, d_1463_v4_, d_1462_i0_, globalState))
            nw270_[int(13)] = ((_dafny.MultiSet(d_1465_v6_)) - (d_1470_v11_)).cardinality
            nw270_[int(14)] = (p0) * (d_1462_i0_)
            nw270_[int(15)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bodqajyph"))) + (d_1472_v13_))
            nw270_[int(16)] = default__.fm27(p0, d_1463_v4_, d_1462_i0_, globalState)
            nw270_[int(17)] = p0
            nw270_[int(18)] = (_dafny.MultiSet([d_1462_i0_])).cardinality
            nw270_[int(19)] = p0
            nw270_[int(20)] = p0
            nw270_[int(21)] = d_1462_i0_
            nw270_[int(22)] = -444
            nw270_[int(23)] = (p0) + (len(d_1478_v18_))
            nw270_[int(24)] = -775
            d_1479_v19_ = nw270_
            index230_ = default__.safeIndex(647, (d_1473_v14_).length(0))
            (d_1473_v14_)[index230_] = d_1462_i0_
            index231_ = default__.safeIndex(647, (d_1473_v14_).length(0))
            rhs247_ = d_1469_v10_
            rhs248_ = d_1462_i0_
            rhs249_ = False
            rhs250_ = d_1470_v11_
            lhs207_ = d_1473_v14_
            lhs208_ = default__.safeIndex(647, (d_1473_v14_).length(0))
            lhs209_ = globalState
            d_1469_v10_ = rhs247_
            lhs207_[lhs208_] = rhs248_
            lhs209_.f27 = rhs249_
            d_1470_v11_ = rhs250_
            (globalState).f24 = (p0) * (d_1462_i0_)
        rhs251_ = not(not (self.f28) or ((self).f37))
        rhs252_ = not((p0) > (p0))
        lhs210_ = globalState
        lhs210_.f21 = rhs251_
        r0 = rhs252_
        (globalState).f24 = default__.fm27(p0, _dafny.CodePoint('m'), p0, globalState)
        d_1458_v0_ = d_1458_v0_
        (globalState).f21 = ((self).f29) or (not (self.f28) or ((self).f38))
        r0 = (self).f29
        return r0

    def m13(self, p0, p1, p2, p3, globalState):
        d_1484_v0_: bool
        out35_: bool
        out35_ = (self).m14(p3, globalState)
        d_1484_v0_ = out35_
        d_1485_v1_: int
        d_1485_v1_ = 359
        d_1486_v2_: str
        d_1486_v2_ = _dafny.CodePoint('b')
        if (d_1485_v1_) >= (len((p2).set(default__.safeIndex(451, len(p2)), d_1486_v2_))):
            d_1487_v3_: _dafny.Seq
            d_1487_v3_ = _dafny.SeqWithoutIsStrInference([self.f28])
            d_1488_v4_: _dafny.Array
            nw271_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1488_v4_ = nw271_
            d_1489_v5_: D1
            d_1489_v5_ = D1_DC1(d_1488_v4_)
            d_1490_v6_: _dafny.Set
            d_1490_v6_ = _dafny.Set({d_1489_v5_})
            d_1491_v7_: _dafny.Map
            d_1491_v7_ = _dafny.Map({d_1487_v3_: d_1490_v6_})
            d_1491_v7_ = d_1491_v7_
            (globalState).f1 = d_1485_v1_
            d_1487_v3_ = d_1487_v3_
            d_1492_v8_: _dafny.Array
            nw272_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1492_v8_ = nw272_
            d_1493_v9_: _dafny.Array
            def lambda47_(d_1494_v0_, d_1495_p3_):
                def lambda48_(d_1496_i0_):
                    return _dafny.Map({d_1494_v0_: d_1495_p3_})

                return lambda48_

            init28_ = lambda47_(d_1484_v0_, p3)
            nw273_ = _dafny.Array(None, 6)
            for i0_28_ in range(nw273_.length(0)):
                nw273_[i0_28_] = init28_(i0_28_)
            d_1493_v9_ = nw273_
            index232_ = default__.safeIndex(125, (d_1492_v8_).length(0))
            (d_1492_v8_)[index232_] = d_1493_v9_
            d_1497_v10_: D6
            d_1497_v10_ = D6_DC14(d_1493_v9_)
            pat_let_tv54_ = d_1493_v9_
            index233_ = default__.safeIndex(125, (d_1492_v8_).length(0))
            def iife119_(_pat_let27_0):
                def iife120_(d_1498_dt__update__tmp_h0_):
                    def iife121_(_pat_let28_0):
                        def iife122_(d_1499_dt__update_hcf19_h0_):
                            return D6_DC14(d_1499_dt__update_hcf19_h0_)
                        return iife122_(_pat_let28_0)
                    return iife121_(pat_let_tv54_)
                return iife120_(_pat_let27_0)
            (d_1492_v8_)[index233_] = (iife119_(d_1497_v10_)).cf19
            d_1500_v11_: _dafny.Array
            nw274_ = _dafny.Array(False, 9)
            d_1500_v11_ = nw274_
            index234_ = default__.safeIndex(287, (d_1500_v11_).length(0))
            (d_1500_v11_)[index234_] = False
            d_1501_v12_: _dafny.Map
            d_1501_v12_ = _dafny.Map({d_1485_v1_: (self).f38})
            d_1502_v13_: _dafny.MultiSet
            d_1502_v13_ = _dafny.MultiSet([self.f28, not((self).f38), p3, p0, p0])
            index235_ = default__.safeIndex(287, (d_1500_v11_).length(0))
            (d_1500_v11_)[index235_] = ((d_1501_v12_)[((d_1502_v13_) | (d_1502_v13_)).cardinality] if (((d_1502_v13_) | (d_1502_v13_)).cardinality) in (d_1501_v12_) else (self).f38)
        elif True:
            d_1503_v14_: _dafny.Array
            nw275_ = _dafny.Array(False, 26)
            d_1503_v14_ = nw275_
            index236_ = default__.safeIndex(384, (d_1503_v14_).length(0))
            (d_1503_v14_)[index236_] = (d_1485_v1_) <= (d_1485_v1_)
            index237_ = default__.safeIndex(384, (d_1503_v14_).length(0))
            rhs253_ = d_1503_v14_
            rhs254_ = p3
            rhs255_ = (d_1485_v1_) * (901)
            rhs256_ = d_1484_v0_
            rhs257_ = p3
            lhs211_ = globalState
            lhs212_ = globalState
            lhs213_ = d_1503_v14_
            lhs214_ = default__.safeIndex(384, (d_1503_v14_).length(0))
            d_1503_v14_ = rhs253_
            lhs211_.f13 = rhs254_
            lhs212_.f24 = rhs255_
            d_1484_v0_ = rhs256_
            lhs213_[lhs214_] = rhs257_
            d_1504_v15_: D2
            d_1504_v15_ = D2_DC5(p3)
            d_1505_v16_: _dafny.MultiSet
            d_1505_v16_ = _dafny.MultiSet([p3, p0])
            d_1506_v17_: _dafny.Map
            d_1506_v17_ = _dafny.Map({d_1484_v0_: ((_dafny.MultiSet([(d_1504_v15_).cf8])).set((d_1503_v14_)[default__.safeIndex(384, (d_1503_v14_).length(0))], default__.abs(d_1485_v1_))).isdisjoint(d_1505_v16_)})
            (globalState).f21 = ((d_1506_v17_)[p3] if (p3) in (d_1506_v17_) else ((self).f38) and (self.f28))
            d_1507_v18_: _dafny.Map
            d_1507_v18_ = _dafny.Map({not((self).f37): _dafny.SeqWithoutIsStrInference([d_1484_v0_, p3])})
            d_1508_v19_: _dafny.Seq
            d_1508_v19_ = _dafny.SeqWithoutIsStrInference([(self).f38, (self).f29])
            d_1507_v18_ = (d_1507_v18_ if (d_1503_v14_)[default__.safeIndex(384, (d_1503_v14_).length(0))] else _dafny.Map({(d_1508_v19_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))), len(d_1508_v19_))]: d_1508_v19_}))
            d_1504_v15_ = d_1504_v15_
            index238_ = default__.safeIndex(384, (d_1503_v14_).length(0))
            (d_1503_v14_)[index238_] = self.f28
        d_1509_v20_: _dafny.Array
        nw276_ = _dafny.Array(_dafny.Seq({}), 23)
        d_1509_v20_ = nw276_
        d_1510_v21_: _dafny.Map
        d_1510_v21_ = _dafny.Map({d_1509_v20_: d_1484_v0_})
        d_1510_v21_ = (d_1510_v21_).set(d_1509_v20_, p0)
        d_1511_i1_: int
        d_1511_i1_ = 0
        with _dafny.label("2"):
            while (self).f37:
                with _dafny.c_label("2"):
                    if (d_1511_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_1511_i1_ = (d_1511_i1_) + (1)
                    d_1485_v1_ = d_1485_v1_
                    if (self).f37:
                        d_1512_v22_: _dafny.Seq
                        d_1512_v22_ = _dafny.SeqWithoutIsStrInference([d_1484_v0_])
                        d_1513_v23_: _dafny.MultiSet
                        d_1513_v23_ = _dafny.MultiSet([d_1512_v22_, d_1512_v22_])
                        d_1514_v25_: _dafny.Map
                        d_1514_v25_ = _dafny.Map({len(d_1512_v22_): p0})
                        def iife123_():
                            coll65_ = _dafny.Map()
                            compr_65_: int
                            for compr_65_ in _dafny.IntegerRange(987, 610):
                                d_1515_v24_: int = compr_65_
                                if ((987) <= (d_1515_v24_)) and ((d_1515_v24_) < (610)):
                                    coll65_[default__.safeModuloInt(d_1515_v24_, d_1485_v1_)] = (self).f29
                            return _dafny.Map(coll65_)
                        (globalState).f22 = default__.fm29(d_1513_v23_, len((iife123_()
                        ) | (d_1514_v25_)), globalState)
                        d_1516_v26_: _dafny.Array
                        nw277_ = _dafny.Array(_dafny.Array(None, 0), 9)
                        d_1516_v26_ = nw277_
                        d_1517_v27_: _dafny.Array
                        nw278_ = _dafny.Array(D4.default()(), 16)
                        d_1517_v27_ = nw278_
                        index239_ = default__.safeIndex(10, (d_1516_v26_).length(0))
                        (d_1516_v26_)[index239_] = d_1517_v27_
                        index240_ = default__.safeIndex(10, (d_1516_v26_).length(0))
                        (d_1516_v26_)[index240_] = d_1517_v27_
                        d_1518_v28_: _dafny.Seq
                        d_1518_v28_ = _dafny.SeqWithoutIsStrInference([d_1485_v1_])
                        d_1519_v29_: _dafny.MultiSet
                        d_1519_v29_ = _dafny.MultiSet([len(d_1518_v28_), 937])
                        d_1520_v30_: C3
                        nw279_ = C3()
                        nw279_.ctor__(True, p0, ((self).f38 if False else p3), (d_1519_v29_).cardinality)
                        d_1520_v30_ = nw279_
                        d_1485_v1_ = ((0) - (d_1485_v1_)) * (d_1485_v1_)
                        (globalState).f1 = d_1485_v1_
                    elif True:
                        d_1521_v31_: _dafny.Array
                        def lambda49_(d_1522_i2_):
                            return False

                        init29_ = lambda49_
                        nw280_ = _dafny.Array(None, 24)
                        for i0_29_ in range(nw280_.length(0)):
                            nw280_[i0_29_] = init29_(i0_29_)
                        d_1521_v31_ = nw280_
                        index241_ = default__.safeIndex(579, (d_1521_v31_).length(0))
                        (d_1521_v31_)[index241_] = self.f28
                        index242_ = default__.safeIndex(579, (d_1521_v31_).length(0))
                        (d_1521_v31_)[index242_] = not(False)
                        d_1523_v32_: _dafny.Seq
                        d_1523_v32_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                        d_1524_v33_: _dafny.Array
                        def lambda50_(d_1525_v1_, d_1526_v31_):
                            def lambda51_(d_1527_i3_):
                                return (d_1527_i3_) + ((_dafny.MultiSet([_dafny.Map({len(_dafny.Map({len(_dafny.Map({(self).f37: d_1525_v1_})): d_1525_v1_})): True}), _dafny.Map({d_1525_v1_: (d_1526_v31_)[default__.safeIndex(579, (d_1526_v31_).length(0))]})])).cardinality)

                            return lambda51_

                        init30_ = lambda50_(d_1485_v1_, d_1521_v31_)
                        nw281_ = _dafny.Array(None, 10)
                        for i0_30_ in range(nw281_.length(0)):
                            nw281_[i0_30_] = init30_(i0_30_)
                        d_1524_v33_ = nw281_
                        d_1528_v34_: _dafny.Map
                        d_1528_v34_ = _dafny.Map({(d_1521_v31_)[default__.safeIndex(579, (d_1521_v31_).length(0))]: d_1524_v33_})
                        d_1529_v35_: D3
                        d_1529_v35_ = D3_DC7(d_1528_v34_)
                        d_1530_v36_: _dafny.Map
                        d_1530_v36_ = _dafny.Map({d_1529_v35_: d_1486_v2_})
                        d_1531_v37_: _dafny.Array
                        nw282_ = _dafny.Array(None, 10)
                        nw282_[int(0)] = d_1486_v2_
                        nw282_[int(1)] = (p2)[default__.safeIndex(609, len(p2))]
                        nw282_[int(2)] = default__.fm42(-160, d_1485_v1_, globalState)
                        nw282_[int(3)] = (d_1486_v2_ if (d_1523_v32_)[default__.safeIndex(len(p2), len(d_1523_v32_))] else d_1486_v2_)
                        nw282_[int(4)] = d_1486_v2_
                        nw282_[int(5)] = d_1486_v2_
                        nw282_[int(6)] = d_1486_v2_
                        nw282_[int(7)] = d_1486_v2_
                        nw282_[int(8)] = ((d_1530_v36_)[d_1529_v35_] if (d_1529_v35_) in (d_1530_v36_) else d_1486_v2_)
                        nw282_[int(9)] = d_1486_v2_
                        d_1531_v37_ = nw282_
                        index243_ = default__.safeIndex(196, (d_1531_v37_).length(0))
                        (d_1531_v37_)[index243_] = d_1486_v2_
                        index244_ = default__.safeIndex(196, (d_1531_v37_).length(0))
                        (d_1531_v37_)[index244_] = d_1486_v2_
                        index245_ = default__.safeIndex(579, (d_1521_v31_).length(0))
                        (d_1521_v31_)[index245_] = self.f28
                        d_1532_v38_: bool
                        out36_: bool
                        out36_ = (self).m14((self).f38, globalState)
                        d_1532_v38_ = out36_
                        (globalState).f24 = d_1485_v1_
                    d_1533_v39_: _dafny.Array
                    nw283_ = _dafny.Array(_dafny.Array(None, 0), 6)
                    d_1533_v39_ = nw283_
                    d_1534_v40_: _dafny.Seq
                    d_1534_v40_ = _dafny.SeqWithoutIsStrInference([d_1485_v1_])
                    d_1535_v41_: _dafny.Array
                    nw284_ = _dafny.Array(None, 11)
                    nw284_[int(0)] = d_1485_v1_
                    nw284_[int(1)] = d_1485_v1_
                    nw284_[int(2)] = d_1485_v1_
                    nw284_[int(3)] = len(p2)
                    nw284_[int(4)] = default__.fm27(d_1485_v1_, _dafny.CodePoint('w'), d_1485_v1_, globalState)
                    nw284_[int(5)] = d_1485_v1_
                    nw284_[int(6)] = d_1485_v1_
                    nw284_[int(7)] = d_1485_v1_
                    nw284_[int(8)] = d_1485_v1_
                    nw284_[int(9)] = d_1485_v1_
                    nw284_[int(10)] = len(d_1534_v40_)
                    d_1535_v41_ = nw284_
                    index246_ = default__.safeIndex(627, (d_1533_v39_).length(0))
                    (d_1533_v39_)[index246_] = d_1535_v41_
                    d_1536_v42_: _dafny.Map
                    d_1536_v42_ = _dafny.Map({p2: (self).f38})
                    index247_ = default__.safeIndex(627, (d_1533_v39_).length(0))
                    rhs258_ = ((d_1536_v42_)[p2] if (p2) in (d_1536_v42_) else p3)
                    rhs259_ = default__.fm27(d_1485_v1_, _dafny.CodePoint('w'), d_1485_v1_, globalState)
                    rhs260_ = (self).f38
                    rhs261_ = d_1535_v41_
                    lhs215_ = globalState
                    lhs216_ = d_1533_v39_
                    lhs217_ = default__.safeIndex(627, (d_1533_v39_).length(0))
                    d_1484_v0_ = rhs258_
                    lhs215_.f1 = rhs259_
                    d_1484_v0_ = rhs260_
                    lhs216_[lhs217_] = rhs261_
                    if not(self.f28):
                        d_1484_v0_ = not((self).f37)
                        d_1537_v43_: _dafny.MultiSet
                        d_1537_v43_ = _dafny.MultiSet([not(not(not((self).f38))), (p1).issubset(default__.fm35(globalState))])
                        d_1537_v43_ = (d_1537_v43_).intersection((d_1537_v43_) | (d_1537_v43_))
                        d_1538_v44_: D13
                        d_1538_v44_ = D13_DC35()
                        d_1539_v45_: _dafny.Array
                        def lambda52_(d_1540_p2_):
                            def lambda53_(d_1541_i4_):
                                return d_1540_p2_

                            return lambda53_

                        init31_ = lambda52_(p2)
                        nw285_ = _dafny.Array(None, 15)
                        for i0_31_ in range(nw285_.length(0)):
                            nw285_[i0_31_] = init31_(i0_31_)
                        d_1539_v45_ = nw285_
                        index248_ = default__.safeIndex(447, (d_1539_v45_).length(0))
                        (d_1539_v45_)[index248_] = (p2) + (p2)
                        d_1542_v46_: D8
                        d_1542_v46_ = D8_DC22()
                        d_1543_v47_: _dafny.MultiSet
                        d_1543_v47_ = _dafny.MultiSet([d_1485_v1_, d_1485_v1_])
                        d_1544_v48_: _dafny.Array
                        nw286_ = _dafny.Array(None, 20)
                        nw286_[int(0)] = d_1486_v2_
                        nw286_[int(1)] = d_1486_v2_
                        nw286_[int(2)] = d_1486_v2_
                        nw286_[int(3)] = d_1486_v2_
                        nw286_[int(4)] = d_1486_v2_
                        nw286_[int(5)] = d_1486_v2_
                        nw286_[int(6)] = d_1486_v2_
                        nw286_[int(7)] = _dafny.CodePoint('h')
                        nw286_[int(8)] = d_1486_v2_
                        nw286_[int(9)] = d_1486_v2_
                        nw286_[int(10)] = d_1486_v2_
                        nw286_[int(11)] = d_1486_v2_
                        nw286_[int(12)] = d_1486_v2_
                        nw286_[int(13)] = d_1486_v2_
                        nw286_[int(14)] = d_1486_v2_
                        nw286_[int(15)] = d_1486_v2_
                        nw286_[int(16)] = d_1486_v2_
                        nw286_[int(17)] = d_1486_v2_
                        nw286_[int(18)] = _dafny.CodePoint('w')
                        nw286_[int(19)] = d_1486_v2_
                        d_1544_v48_ = nw286_
                        d_1545_v49_: _dafny.Array
                        d_1545_v49_ = d_1544_v48_
                        d_1546_v50_: _dafny.Set
                        d_1546_v50_ = _dafny.Set({d_1544_v48_})
                        d_1547_v51_: _dafny.Map
                        d_1547_v51_ = _dafny.Map({(d_1543_v47_).cardinality: d_1546_v50_})
                        index249_ = default__.safeIndex(447, (d_1539_v45_).length(0))
                        rhs262_ = default__.fm60(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neq")), (d_1534_v40_) < (d_1534_v40_), not((self).f37), (d_1485_v1_) * ((_dafny.MultiSet([p1, p1])).cardinality), globalState)
                        rhs263_ = p2
                        rhs264_ = d_1486_v2_
                        rhs265_ = d_1542_v46_
                        rhs266_ = (_dafny.Map({d_1485_v1_: d_1546_v50_})) | (d_1547_v51_)
                        lhs218_ = d_1539_v45_
                        lhs219_ = default__.safeIndex(447, (d_1539_v45_).length(0))
                        lhs220_ = globalState
                        d_1538_v44_ = rhs262_
                        lhs218_[lhs219_] = rhs263_
                        lhs220_.f23 = rhs264_
                        d_1542_v46_ = rhs265_
                        d_1547_v51_ = rhs266_
                        d_1548_v52_: _dafny.Set
                        d_1548_v52_ = _dafny.Set({(self).f29})
                        rhs267_ = default__.safeDivisionInt(d_1485_v1_, d_1485_v1_)
                        rhs268_ = d_1548_v52_
                        lhs221_ = globalState
                        lhs222_ = globalState
                        lhs221_.f24 = rhs267_
                        lhs222_.f26 = rhs268_
                        d_1549_v53_: _dafny.Map
                        d_1549_v53_ = _dafny.Map({p3: d_1485_v1_})
                        d_1549_v53_ = d_1549_v53_
                    elif True:
                        (globalState).f1 = 919
                        (globalState).f1 = d_1485_v1_
                        d_1550_v54_: D7
                        d_1550_v54_ = D7_DC17(self.f28)
                        d_1551_v55_: _dafny.Array
                        nw287_ = _dafny.Array(False, 17)
                        d_1551_v55_ = nw287_
                        d_1552_v56_: _dafny.Seq
                        d_1552_v56_ = _dafny.SeqWithoutIsStrInference([False, p0, (self).f37])
                        d_1553_v57_: _dafny.MultiSet
                        d_1553_v57_ = _dafny.MultiSet([d_1552_v56_, d_1552_v56_])
                        index250_ = default__.safeIndex(703, (d_1551_v55_).length(0))
                        (d_1551_v55_)[index250_] = default__.fm29(d_1553_v57_, d_1485_v1_, globalState)
                        d_1554_v58_: _dafny.Array
                        nw288_ = _dafny.Array(_dafny.Array(None, 0), 12)
                        d_1554_v58_ = nw288_
                        d_1555_v59_: _dafny.MultiSet
                        d_1555_v59_ = _dafny.MultiSet([self.f28, d_1484_v0_])
                        index251_ = default__.safeIndex(703, (d_1551_v55_).length(0))
                        rhs269_ = default__.fm27((d_1555_v59_).cardinality, d_1486_v2_, (d_1485_v1_) + (len(default__.fm61(d_1485_v1_, d_1485_v1_, (self).f38, d_1484_v0_, globalState))), globalState)
                        rhs270_ = _dafny.CodePoint('b')
                        rhs271_ = d_1550_v54_
                        rhs272_ = d_1484_v0_
                        rhs273_ = d_1554_v58_
                        lhs223_ = globalState
                        lhs224_ = globalState
                        lhs225_ = d_1551_v55_
                        lhs226_ = default__.safeIndex(703, (d_1551_v55_).length(0))
                        lhs223_.f1 = rhs269_
                        lhs224_.f23 = rhs270_
                        d_1550_v54_ = rhs271_
                        lhs225_[lhs226_] = rhs272_
                        d_1554_v58_ = rhs273_
                        d_1536_v42_ = (d_1536_v42_).set(p2, p0)
                        index252_ = default__.safeIndex(703, (d_1551_v55_).length(0))
                        (d_1551_v55_)[index252_] = (self).f37
                    pass
            pass
        (globalState).f1 = default__.safeDivisionInt(d_1485_v1_, (697) * (d_1485_v1_))
        d_1556_v60_: _dafny.Array
        nw289_ = _dafny.Array(False, 20)
        d_1556_v60_ = nw289_
        d_1557_v61_: _dafny.Set
        d_1557_v61_ = _dafny.Set({d_1556_v60_, d_1556_v60_, d_1556_v60_})
        d_1558_v62_: _dafny.Seq
        d_1558_v62_ = _dafny.SeqWithoutIsStrInference([(self).f37])
        if ((d_1557_v61_).isdisjoint(d_1557_v61_)) not in ((d_1558_v62_) + (d_1558_v62_)):
            d_1509_v20_ = d_1509_v20_
            d_1559_v63_: _dafny.Seq
            d_1559_v63_ = _dafny.SeqWithoutIsStrInference([d_1485_v1_])
            (globalState).f17 = (d_1559_v63_) + (_dafny.SeqWithoutIsStrInference([d_1485_v1_]))
            d_1560_v64_: _dafny.Array
            nw290_ = _dafny.Array(int(0), 9)
            d_1560_v64_ = nw290_
            d_1561_v65_: _dafny.Map
            d_1561_v65_ = _dafny.Map({(self).f29: d_1560_v64_})
            source27_ = D3_DC7((d_1561_v65_).set(p3, d_1560_v64_))
            if source27_.is_DC8:
                d_1562___mcc_h0_ = source27_.cf11
                d_1563_cf11_ = d_1562___mcc_h0_
                index253_ = default__.safeIndex(905, (d_1560_v64_).length(0))
                (d_1560_v64_)[index253_] = d_1485_v1_
                index254_ = default__.safeIndex(905, (d_1560_v64_).length(0))
                rhs274_ = d_1485_v1_
                rhs275_ = (self).f29
                lhs227_ = d_1560_v64_
                lhs228_ = default__.safeIndex(905, (d_1560_v64_).length(0))
                lhs229_ = globalState
                lhs227_[lhs228_] = rhs274_
                lhs229_.f27 = rhs275_
                d_1564_v66_: _dafny.Array
                def lambda54_(d_1565_v2_):
                    def lambda55_(d_1566_i5_):
                        return d_1565_v2_

                    return lambda55_

                init32_ = lambda54_(d_1486_v2_)
                nw291_ = _dafny.Array(None, 3)
                for i0_32_ in range(nw291_.length(0)):
                    nw291_[i0_32_] = init32_(i0_32_)
                d_1564_v66_ = nw291_
                d_1567_v67_: _dafny.Map
                d_1567_v67_ = _dafny.Map({d_1486_v2_: d_1564_v66_})
                d_1567_v67_ = (d_1567_v67_).set(_dafny.CodePoint('b'), d_1564_v66_)
                (globalState).f24 = len(p1)
                d_1568_v68_: D1
                d_1568_v68_ = D1_DC2((d_1560_v64_)[default__.safeIndex(905, (d_1560_v64_).length(0))], (d_1560_v64_)[default__.safeIndex(905, (d_1560_v64_).length(0))], d_1563_cf11_, (self).f29)
                d_1569_v69_: D1
                d_1569_v69_ = D1_DC3(d_1568_v68_)
                d_1570_v70_: _dafny.MultiSet
                d_1570_v70_ = _dafny.MultiSet([d_1569_v69_])
                d_1571_v71_: C3
                nw292_ = C3()
                nw292_.ctor__(p3, p0, p0, default__.safeModuloInt((0) - ((0) - ((d_1570_v70_).cardinality)), len(d_1558_v62_)))
                d_1571_v71_ = nw292_
            elif True:
                d_1572___mcc_h1_ = source27_.cf10
                d_1573_cf10_ = d_1572___mcc_h1_
                d_1574_v72_: _dafny.Map
                d_1574_v72_ = _dafny.Map({p1: d_1560_v64_})
                d_1574_v72_ = (d_1574_v72_).set(p1, (d_1560_v64_ if self.f28 else d_1560_v64_))
                d_1575_v73_: _dafny.Seq
                d_1575_v73_ = _dafny.SeqWithoutIsStrInference([d_1486_v2_, d_1486_v2_, d_1486_v2_, d_1486_v2_, (d_1486_v2_ if (self).f38 else default__.fm42(d_1485_v1_, 282, globalState))])
                d_1576_v74_: _dafny.Map
                d_1576_v74_ = _dafny.Map({d_1486_v2_: d_1484_v0_})
                rhs276_ = p2
                rhs277_ = not (not (p3) or (d_1484_v0_)) or (d_1484_v0_)
                rhs278_ = (d_1558_v62_)[default__.safeIndex((d_1485_v1_ if True else d_1485_v1_), len(d_1558_v62_))]
                rhs279_ = (len(d_1576_v74_)) - (d_1485_v1_)
                rhs280_ = d_1485_v1_
                lhs230_ = globalState
                lhs231_ = globalState
                lhs232_ = globalState
                lhs233_ = globalState
                d_1575_v73_ = rhs276_
                lhs230_.f22 = rhs277_
                lhs231_.f13 = rhs278_
                lhs232_.f24 = rhs279_
                lhs233_.f1 = rhs280_
                (globalState).f1 = (0) - (d_1485_v1_)
                d_1577_v75_: C2
                nw293_ = C2()
                nw293_.ctor__()
                d_1577_v75_ = nw293_
            d_1578_v76_: _dafny.Map
            d_1578_v76_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not(self.f28), p3, d_1484_v0_]): (self).f29})
            d_1579_v77_: _dafny.Map
            d_1579_v77_ = _dafny.Map({(d_1578_v76_ if (self).f37 else d_1578_v76_): False})
            d_1580_v78_: _dafny.Seq
            d_1580_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])
            d_1581_v79_: _dafny.Set
            d_1581_v79_ = _dafny.Set({(self).f38})
            d_1579_v77_ = (d_1579_v77_).set(d_1578_v76_, not (default__.fm29(_dafny.MultiSet(d_1580_v78_), len(d_1581_v79_), globalState)) or (d_1484_v0_))
            (globalState).f27 = (d_1558_v62_) != (d_1558_v62_)
        elif True:
            d_1582_v80_: C2
            nw294_ = C2()
            nw294_.ctor__()
            d_1582_v80_ = nw294_
            (globalState).f13 = not(((0) - (d_1485_v1_)) > ((d_1485_v1_ if (d_1558_v62_)[default__.safeIndex(d_1485_v1_, len(d_1558_v62_))] else d_1485_v1_)))
            (globalState).f1 = d_1485_v1_
            d_1583_v81_: D7
            d_1583_v81_ = D7_DC18((len(p2)) > (d_1485_v1_))
            d_1583_v81_ = D7_DC18(p0)
            d_1558_v62_ = (d_1558_v62_) + (d_1558_v62_)

    def m14(self, p0, globalState):
        r0: bool = False
        d_1584_i0_: int
        d_1584_i0_ = 0
        with _dafny.label("3"):
            while p0:
                with _dafny.c_label("3"):
                    if (d_1584_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_1584_i0_ = (d_1584_i0_) + (1)
                    if not(not(not (p0) or ((self).f38))):
                        d_1585_v0_: _dafny.Array
                        def lambda56_(d_1586_i1_):
                            return ((D7_DC17((self).f37)).cf22 if (self).f37 else (self).f38)

                        init33_ = lambda56_
                        nw295_ = _dafny.Array(None, 21)
                        for i0_33_ in range(nw295_.length(0)):
                            nw295_[i0_33_] = init33_(i0_33_)
                        d_1585_v0_ = nw295_
                        index255_ = default__.safeIndex(826, (d_1585_v0_).length(0))
                        (d_1585_v0_)[index255_] = (self).f29
                        index256_ = default__.safeIndex(826, (d_1585_v0_).length(0))
                        (d_1585_v0_)[index256_] = (self).f38
                        d_1587_v1_: int
                        d_1587_v1_ = -592
                        (globalState).f24 = d_1587_v1_
                        d_1588_v2_: C2
                        nw296_ = C2()
                        nw296_.ctor__()
                        d_1588_v2_ = nw296_
                        (globalState).f1 = (len(_dafny.SeqWithoutIsStrInference([d_1587_v1_, 319, d_1587_v1_]))) + (d_1587_v1_)
                        d_1589_v3_: _dafny.Seq
                        d_1589_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                        d_1590_v4_: _dafny.Set
                        d_1590_v4_ = _dafny.Set({d_1589_v3_, (d_1589_v3_) + (d_1589_v3_)})
                        d_1590_v4_ = d_1590_v4_
                    elif True:
                        d_1591_v5_: int
                        d_1591_v5_ = 93
                        (globalState).f24 = (0) - ((647) - (d_1591_v5_))
                        (globalState).f24 = default__.safeDivisionInt(d_1591_v5_, (d_1591_v5_) * (d_1591_v5_))
                        d_1592_v6_: _dafny.Array
                        nw297_ = _dafny.Array(int(0), 24)
                        d_1592_v6_ = nw297_
                        d_1593_v7_: _dafny.Seq
                        d_1593_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdwqcqsh"))
                        d_1594_v8_: _dafny.Set
                        d_1594_v8_ = _dafny.Set({len(d_1593_v7_), d_1591_v5_, d_1591_v5_, 164})
                        index257_ = default__.safeIndex(627, (d_1592_v6_).length(0))
                        (d_1592_v6_)[index257_] = len(d_1594_v8_)
                        index258_ = default__.safeIndex(627, (d_1592_v6_).length(0))
                        (d_1592_v6_)[index258_] = (d_1591_v5_) - (d_1591_v5_)
                        d_1593_v7_ = d_1593_v7_
                        d_1595_v9_: _dafny.Array
                        nw298_ = _dafny.Array(_dafny.Array(None, 0), 22)
                        d_1595_v9_ = nw298_
                        d_1596_v10_: D1
                        d_1596_v10_ = D1_DC1(d_1595_v9_)
                        d_1597_v11_: D1
                        d_1597_v11_ = D1_DC3(d_1596_v10_)
                        d_1598_v12_: D1
                        d_1598_v12_ = D1_DC3(D1_DC3(d_1596_v10_))
                        d_1599_v13_: _dafny.Map
                        d_1599_v13_ = _dafny.Map({d_1594_v8_: d_1598_v12_})
                        d_1600_v14_: _dafny.Map
                        d_1600_v14_ = _dafny.Map({p0: d_1599_v13_})
                        d_1601_v15_: _dafny.MultiSet
                        d_1601_v15_ = _dafny.MultiSet([default__.fm62(globalState)])
                        d_1602_v17_: _dafny.Map
                        def iife124_():
                            coll66_ = _dafny.Set()
                            compr_66_: int
                            for compr_66_ in _dafny.IntegerRange(-239, 353):
                                d_1603_v16_: int = compr_66_
                                if ((-239) <= (d_1603_v16_)) and ((d_1603_v16_) < (353)):
                                    coll66_ = coll66_.union(_dafny.Set([default__.safeDivisionInt(d_1603_v16_, (d_1592_v6_)[default__.safeIndex(627, (d_1592_v6_).length(0))])]))
                            return _dafny.Set(coll66_)
                        d_1602_v17_ = _dafny.Map({(self).f29: iife124_()
                        })
                        d_1604_v19_: D7
                        d_1604_v19_ = D7_DC19(self.f28, (self).f29)
                        d_1605_v20_: _dafny.Map
                        def iife125_():
                            coll67_ = _dafny.Set()
                            compr_67_: int
                            for compr_67_ in _dafny.IntegerRange(894, 2):
                                d_1606_v18_: int = compr_67_
                                if ((894) <= (d_1606_v18_)) and ((d_1606_v18_) < (2)):
                                    coll67_ = coll67_.union(_dafny.Set([(d_1606_v18_) + (454)]))
                            return _dafny.Set(coll67_)
                        def iife126_():
                            coll68_ = _dafny.Set()
                            compr_68_: int
                            for compr_68_ in _dafny.IntegerRange(894, 2):
                                d_1607_v18_: int = compr_68_
                                if ((894) <= (d_1607_v18_)) and ((d_1607_v18_) < (2)):
                                    coll68_ = coll68_.union(_dafny.Set([(d_1607_v18_) + (454)]))
                            return _dafny.Set(coll68_)
                        d_1605_v20_ = _dafny.Map({len(((d_1600_v14_)[default__.fm29(d_1601_v15_, (0) - (len(((d_1602_v17_)[p0] if (p0) in (d_1602_v17_) else iife125_()
                        ))), globalState)] if (default__.fm29(d_1601_v15_, (0) - (len(((d_1602_v17_)[p0] if (p0) in (d_1602_v17_) else iife126_()
                        ))), globalState)) in (d_1600_v14_) else d_1599_v13_)): not((d_1604_v19_).cf24)})
                        d_1605_v20_ = (d_1605_v20_).set(d_1591_v5_, self.f28)
                    d_1608_v21_: int
                    d_1608_v21_ = -480
                    (globalState).f1 = d_1608_v21_
                    d_1609_v22_: _dafny.Seq
                    d_1609_v22_ = _dafny.SeqWithoutIsStrInference([p0, (self).f37])
                    d_1610_v23_: _dafny.MultiSet
                    d_1610_v23_ = _dafny.MultiSet([d_1609_v22_])
                    d_1611_v24_: D4
                    d_1611_v24_ = D4_DC10(default__.fm29(d_1610_v23_, d_1608_v21_, globalState), d_1608_v21_, d_1608_v21_)
                    d_1612_v25_: _dafny.Map
                    d_1612_v25_ = _dafny.Map({(d_1611_v24_).cf13: d_1608_v21_})
                    d_1612_v25_ = (d_1612_v25_).set((d_1608_v21_) < (d_1608_v21_), default__.safeDivisionInt(d_1608_v21_, d_1608_v21_))
                    d_1613_v26_: D2
                    d_1613_v26_ = D2_DC5(self.f28)
                    if (d_1613_v26_).cf8:
                        d_1614_v27_: C0
                        nw299_ = C0()
                        nw299_.ctor__()
                        d_1614_v27_ = nw299_
                        d_1612_v25_ = (d_1612_v25_).set((self).f29, 698)
                        d_1615_v28_: C2
                        nw300_ = C2()
                        nw300_.ctor__()
                        d_1615_v28_ = nw300_
                        d_1616_v29_: _dafny.MultiSet
                        d_1616_v29_ = _dafny.MultiSet([(self).f38])
                        (globalState).f17 = _dafny.SeqWithoutIsStrInference([d_1608_v21_, -861, d_1608_v21_, (d_1608_v21_) * (((d_1616_v29_)[(self).f38] if ((self).f38) in (d_1616_v29_) else d_1608_v21_)), d_1608_v21_])
                        (globalState).f24 = (0) - (d_1608_v21_)
                    elif True:
                        d_1617_v30_: D10
                        d_1617_v30_ = D10_DC29(D10_DC28(_dafny.Map({(self).f37: self.f28})))
                        rhs281_ = d_1608_v21_
                        rhs282_ = len(_dafny.Set({d_1617_v30_, d_1617_v30_, d_1617_v30_}))
                        rhs283_ = (self).f29
                        lhs234_ = globalState
                        lhs235_ = globalState
                        lhs236_ = globalState
                        lhs234_.f24 = rhs281_
                        lhs235_.f24 = rhs282_
                        lhs236_.f27 = rhs283_
                        d_1618_v31_: _dafny.Set
                        d_1618_v31_ = _dafny.Set({self.f28})
                        d_1619_v32_: _dafny.Seq
                        d_1619_v32_ = _dafny.SeqWithoutIsStrInference([d_1618_v31_])
                        d_1619_v32_ = (_dafny.SeqWithoutIsStrInference([d_1618_v31_])) + ((d_1619_v32_) + (d_1619_v32_))
                        d_1620_v33_: _dafny.Seq
                        d_1620_v33_ = d_1609_v22_
                        r0 = ((d_1609_v22_) + ((d_1620_v33_))) == (d_1609_v22_)
                        d_1621_v34_: _dafny.Array
                        nw301_ = _dafny.Array(_dafny.Map({}), 21)
                        d_1621_v34_ = nw301_
                        d_1622_v35_: _dafny.Map
                        d_1622_v35_ = _dafny.Map({(self).f38: (self).f29})
                        index259_ = default__.safeIndex(277, (d_1621_v34_).length(0))
                        (d_1621_v34_)[index259_] = d_1622_v35_
                        index260_ = default__.safeIndex(277, (d_1621_v34_).length(0))
                        (d_1621_v34_)[index260_] = d_1622_v35_
                        d_1623_v36_: str
                        d_1623_v36_ = _dafny.CodePoint('k')
                        d_1624_v37_: _dafny.Map
                        d_1624_v37_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1625_i2_ in range(default__.abs(620))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))), d_1623_v36_)): d_1608_v21_})
                        d_1626_v38_: _dafny.Map
                        d_1626_v38_ = _dafny.Map({d_1624_v37_: d_1608_v21_})
                        d_1626_v38_ = (d_1626_v38_).set(d_1624_v37_, d_1608_v21_)
                    pass
            pass
        d_1627_v39_: _dafny.Seq
        d_1627_v39_ = _dafny.SeqWithoutIsStrInference([self.f28])
        d_1628_v40_: str
        d_1628_v40_ = _dafny.CodePoint('a')
        d_1629_v41_: _dafny.Map
        d_1629_v41_ = _dafny.Map({d_1628_v40_: (self).f37})
        d_1630_v42_: _dafny.MultiSet
        d_1630_v42_ = _dafny.MultiSet([d_1629_v41_])
        (globalState).f27 = (d_1627_v39_)[default__.safeIndex((d_1630_v42_).cardinality, len(d_1627_v39_))]
        d_1631_v43_: _dafny.Seq
        d_1631_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_1632_v44_: int
        d_1632_v44_ = 594
        (globalState).f1 = (len((d_1631_v43_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arnimlkw"))))) - (default__.safeModuloInt(d_1632_v44_, d_1632_v44_))
        d_1633_i3_: int
        d_1633_i3_ = 0
        with _dafny.label("4"):
            while not (p0) or (((self).f37 if (self).f37 else (self).f37)):
                with _dafny.c_label("4"):
                    if (d_1633_i3_) >= (100):
                        raise _dafny.Break("4")
                    d_1633_i3_ = (d_1633_i3_) + (1)
                    (globalState).f1 = d_1632_v44_
                    d_1634_v45_: _dafny.Set
                    d_1634_v45_ = _dafny.Set({(self).f38, self.f28})
                    d_1635_v46_: _dafny.Map
                    d_1635_v46_ = _dafny.Map({(self).f38: d_1634_v45_})
                    d_1636_v47_: _dafny.Map
                    d_1636_v47_ = _dafny.Map({p0: (0) - (d_1632_v44_)})
                    d_1637_v48_: _dafny.Array
                    nw302_ = _dafny.Array(None, 7)
                    nw302_[int(0)] = d_1632_v44_
                    nw302_[int(1)] = d_1632_v44_
                    nw302_[int(2)] = len((d_1635_v46_) | (d_1635_v46_))
                    nw302_[int(3)] = len((d_1631_v43_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjrhmfa"))).set(default__.safeIndex((0) - (d_1632_v44_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjrhmfa")))), d_1628_v40_)))
                    nw302_[int(4)] = ((d_1636_v47_)[(self).f38] if ((self).f38) in (d_1636_v47_) else -107)
                    nw302_[int(5)] = d_1632_v44_
                    nw302_[int(6)] = d_1632_v44_
                    d_1637_v48_ = nw302_
                    index261_ = default__.safeIndex(465, (d_1637_v48_).length(0))
                    (d_1637_v48_)[index261_] = d_1632_v44_
                    index262_ = default__.safeIndex(465, (d_1637_v48_).length(0))
                    (d_1637_v48_)[index262_] = default__.safeDivisionInt(d_1632_v44_, d_1632_v44_)
                    d_1638_v49_: _dafny.Array
                    def lambda57_(d_1639_i4_):
                        return (self).f29

                    init34_ = lambda57_
                    nw303_ = _dafny.Array(None, 13)
                    for i0_34_ in range(nw303_.length(0)):
                        nw303_[i0_34_] = init34_(i0_34_)
                    d_1638_v49_ = nw303_
                    d_1640_v50_: _dafny.Seq
                    d_1640_v50_ = _dafny.SeqWithoutIsStrInference([d_1638_v49_, d_1638_v49_, d_1638_v49_])
                    d_1641_v51_: _dafny.Map
                    d_1641_v51_ = _dafny.Map({(d_1638_v49_ if (self).f38 else (d_1640_v50_)[default__.safeIndex(d_1632_v44_, len(d_1640_v50_))]): (-240) != (d_1632_v44_)})
                    d_1642_v52_: _dafny.Map
                    d_1642_v52_ = _dafny.Map({d_1632_v44_: (self).f38})
                    d_1641_v51_ = (d_1641_v51_).set((d_1638_v49_ if ((d_1642_v52_)[(d_1637_v48_)[default__.safeIndex(465, (d_1637_v48_).length(0))]] if ((d_1637_v48_)[default__.safeIndex(465, (d_1637_v48_).length(0))]) in (d_1642_v52_) else p0) else d_1638_v49_), p0)
                    d_1643_v53_: _dafny.Array
                    nw304_ = _dafny.Array(None, 2)
                    nw304_[int(0)] = d_1627_v39_
                    nw304_[int(1)] = d_1627_v39_
                    d_1643_v53_ = nw304_
                    d_1644_v54_: C2
                    nw305_ = C2()
                    nw305_.ctor__()
                    d_1644_v54_ = nw305_
                    d_1645_v55_: _dafny.Map
                    d_1645_v55_ = _dafny.Map({(self).f29: d_1644_v54_})
                    index263_ = default__.safeIndex(11, (d_1643_v53_).length(0))
                    (d_1643_v53_)[index263_] = (d_1627_v39_).set(default__.safeIndex(len(d_1645_v55_), len(d_1627_v39_)), not((self).f29))
                    d_1646_v56_: _dafny.MultiSet
                    d_1646_v56_ = _dafny.MultiSet([(self).f38])
                    d_1647_v57_: _dafny.Set
                    d_1647_v57_ = _dafny.Set({d_1646_v56_})
                    d_1648_v58_: C3
                    nw306_ = C3()
                    nw306_.ctor__((self).f29, (d_1647_v57_).ispropersubset(d_1647_v57_), (d_1632_v44_) != (d_1632_v44_), default__.fm27(d_1632_v44_, d_1628_v40_, d_1632_v44_, globalState))
                    d_1648_v58_ = nw306_
                    index264_ = default__.safeIndex(11, (d_1643_v53_).length(0))
                    rhs284_ = d_1627_v39_
                    rhs285_ = d_1648_v58_
                    rhs286_ = ((d_1642_v52_)[d_1632_v44_] if (d_1632_v44_) in (d_1642_v52_) else (d_1644_v54_).fm53(globalState))
                    rhs287_ = d_1631_v43_
                    rhs288_ = (d_1637_v48_)[default__.safeIndex(465, (d_1637_v48_).length(0))]
                    lhs237_ = d_1643_v53_
                    lhs238_ = default__.safeIndex(11, (d_1643_v53_).length(0))
                    lhs239_ = self
                    lhs240_ = globalState
                    lhs237_[lhs238_] = rhs284_
                    d_1648_v58_ = rhs285_
                    lhs239_.f28 = rhs286_
                    d_1631_v43_ = rhs287_
                    lhs240_.f1 = rhs288_
                    pass
            pass
        d_1649_v59_: _dafny.Set
        d_1649_v59_ = _dafny.Set({self.f28, self.f28, self.f28})
        (globalState).f1 = len(d_1649_v59_)
        d_1650_v60_: _dafny.Array
        nw307_ = _dafny.Array(None, 15)
        nw307_[int(0)] = (d_1631_v43_ if (self).f38 else default__.fm51(d_1632_v44_, d_1632_v44_, globalState))
        nw307_[int(1)] = d_1631_v43_
        nw307_[int(2)] = d_1631_v43_
        nw307_[int(3)] = (d_1631_v43_) + (d_1631_v43_)
        nw307_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1628_v40_])
        nw307_[int(5)] = (d_1631_v43_).set(default__.safeIndex(d_1632_v44_, len(d_1631_v43_)), d_1628_v40_)
        nw307_[int(6)] = d_1631_v43_
        nw307_[int(7)] = ((d_1631_v43_) + (d_1631_v43_)).set(default__.safeIndex(d_1632_v44_, len((d_1631_v43_) + (d_1631_v43_))), d_1628_v40_)
        nw307_[int(8)] = (d_1631_v43_) + (default__.fm51(d_1632_v44_, d_1632_v44_, globalState))
        nw307_[int(9)] = d_1631_v43_
        nw307_[int(10)] = _dafny.SeqWithoutIsStrInference([d_1628_v40_ for d_1651_i6_ in range(default__.abs(585))])
        nw307_[int(11)] = (_dafny.SeqWithoutIsStrInference([d_1628_v40_])) + (d_1631_v43_)
        nw307_[int(12)] = d_1631_v43_
        nw307_[int(13)] = d_1631_v43_
        nw307_[int(14)] = d_1631_v43_
        d_1650_v60_ = nw307_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1650_v60_).length(0)):
            d_1652_i5_: int = guard_loop_4_
            if (True) and (((0) <= (d_1652_i5_)) and ((d_1652_i5_) < ((d_1650_v60_).length(0)))):
                (d_1650_v60_)[(d_1652_i5_)] = d_1631_v43_
        r0 = True
        return r0

    @property
    def f37(self):
        return self._f37
    @property
    def f38(self):
        return self._f38

class C5(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self.f35: bool = False
        self.f36: str = _dafny.CodePoint('D')
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
    def ctor__(self, f35, f36, f28, f29):
        (self).f35 = f35
        (self).f36 = f36
        (self).f28 = f28
        (self)._f29 = f29

    def m5(self, p0, globalState):
        r0: bool = False
        d_1653_v0_: _dafny.Seq
        d_1653_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjjo"))
        d_1654_v1_: _dafny.Seq
        d_1654_v1_ = _dafny.SeqWithoutIsStrInference([(True) == (self.f28)])
        d_1655_v2_: _dafny.Map
        d_1655_v2_ = _dafny.Map({self.f35: (self).f29})
        d_1656_v3_: _dafny.Map
        d_1656_v3_ = _dafny.Map({not (self.f28) or (self.f35): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ek"))})
        d_1657_v4_: D3
        d_1657_v4_ = D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlrdo")))
        d_1658_v5_: _dafny.MultiSet
        d_1658_v5_ = _dafny.MultiSet([False, self.f35, self.f35, (self).f29, self.f35])
        pat_let_tv55_ = d_1654_v1_
        pat_let_tv56_ = d_1654_v1_
        pat_let_tv57_ = p0
        pat_let_tv58_ = d_1654_v1_
        def lambda58_(source28_):
            if source28_.is_DC8:
                d_1659___mcc_h0_ = source28_.cf11
                d_1660_cf11_ = d_1659___mcc_h0_
                return (pat_let_tv55_) + (_dafny.SeqWithoutIsStrInference([(self).f29, (self).f29]))
            elif True:
                d_1661___mcc_h1_ = source28_.cf10
                d_1662_cf10_ = d_1661___mcc_h1_
                return (_dafny.SeqWithoutIsStrInference([self.f35])) + (_dafny.SeqWithoutIsStrInference([False, (pat_let_tv56_)[default__.safeIndex(pat_let_tv57_, len(pat_let_tv58_))]]))

        rhs289_ = ((d_1656_v3_)[self.f35] if (self.f35) in (d_1656_v3_) else d_1653_v0_)
        rhs290_ = lambda58_(d_1657_v4_)
        rhs291_ = self.f28
        rhs292_ = _dafny.Map({((d_1658_v5_).cardinality) >= (p0): self.f28})
        lhs241_ = self
        d_1653_v0_ = rhs289_
        d_1654_v1_ = rhs290_
        lhs241_.f35 = rhs291_
        d_1655_v2_ = rhs292_
        d_1663_v6_: _dafny.Array
        nw308_ = _dafny.Array(_dafny.CodePoint('D'), 5)
        d_1663_v6_ = nw308_
        d_1663_v6_ = d_1663_v6_
        d_1664_v7_: _dafny.Array
        def lambda59_(d_1665_p0_):
            def lambda60_(d_1666_i0_):
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1665_p0_, d_1665_p0_]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1665_p0_, d_1665_p0_, (_dafny.MultiSet([d_1665_p0_])).cardinality, d_1665_p0_, len(_dafny.Map({self.f28: self.f35}))])) for d_1667_i1_ in range(default__.abs(603))])])

            return lambda60_

        init35_ = lambda59_(p0)
        nw309_ = _dafny.Array(None, 8)
        for i0_35_ in range(nw309_.length(0)):
            nw309_[i0_35_] = init35_(i0_35_)
        d_1664_v7_ = nw309_
        index265_ = default__.safeIndex(158, (d_1664_v7_).length(0))
        (d_1664_v7_)[index265_] = default__.fm23(((d_1658_v5_)[self.f28] if (self.f28) in (d_1658_v5_) else p0), self.f36, p0, globalState)
        d_1668_v8_: _dafny.Seq
        d_1668_v8_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1669_v9_: _dafny.Seq
        d_1669_v9_ = _dafny.SeqWithoutIsStrInference([d_1668_v8_, d_1668_v8_])
        index266_ = default__.safeIndex(158, (d_1664_v7_).length(0))
        (d_1664_v7_)[index266_] = (d_1669_v9_) + (d_1669_v9_)
        d_1670_v10_: _dafny.MultiSet
        d_1670_v10_ = _dafny.MultiSet([394, (0) - (p0), len(d_1653_v0_), p0, p0])
        d_1671_v11_: D4
        d_1671_v11_ = D4_DC9(d_1670_v10_)
        d_1672_v12_: _dafny.Set
        d_1672_v12_ = _dafny.Set({D4_DC9(default__.fm24(_dafny.Set({(_dafny.MultiSet([self.f35])).cardinality, p0}), self.f35, globalState)), d_1671_v11_, d_1671_v11_, d_1671_v11_, d_1671_v11_})
        if (d_1672_v12_).ispropersubset(d_1672_v12_):
            d_1668_v8_ = d_1668_v8_
            d_1673_v13_: _dafny.Array
            nw310_ = _dafny.Array(None, 16)
            nw310_[int(0)] = self.f35
            nw310_[int(1)] = (self).f29
            nw310_[int(2)] = (self).f29
            nw310_[int(3)] = False
            nw310_[int(4)] = (d_1654_v1_)[default__.safeIndex(len(d_1654_v1_), len(d_1654_v1_))]
            nw310_[int(5)] = (self).f29
            nw310_[int(6)] = (self).f29
            nw310_[int(7)] = self.f35
            nw310_[int(8)] = self.f35
            nw310_[int(9)] = self.f35
            nw310_[int(10)] = self.f35
            nw310_[int(11)] = self.f28
            nw310_[int(12)] = (self).f29
            nw310_[int(13)] = (self).f29
            nw310_[int(14)] = self.f35
            nw310_[int(15)] = (self).f29
            d_1673_v13_ = nw310_
            index267_ = default__.safeIndex(105, (d_1673_v13_).length(0))
            (d_1673_v13_)[index267_] = self.f35
            d_1674_v14_: _dafny.Array
            def lambda61_(d_1675_p0_):
                def lambda62_(d_1676_i2_):
                    return default__.safeModuloInt(d_1676_i2_, (0) - (d_1675_p0_))

                return lambda62_

            init36_ = lambda61_(p0)
            nw311_ = _dafny.Array(None, 4)
            for i0_36_ in range(nw311_.length(0)):
                nw311_[i0_36_] = init36_(i0_36_)
            d_1674_v14_ = nw311_
            d_1677_v15_: _dafny.Array
            d_1677_v15_ = d_1674_v14_
            d_1678_v16_: _dafny.MultiSet
            d_1678_v16_ = _dafny.MultiSet([d_1677_v15_, d_1674_v14_, d_1677_v15_])
            index268_ = default__.safeIndex(105, (d_1673_v13_).length(0))
            (d_1673_v13_)[index268_] = (d_1678_v16_).issubset(d_1678_v16_)
            d_1679_v17_: _dafny.Map
            d_1679_v17_ = _dafny.Map({self.f35: p0})
            (globalState).f24 = ((0) - (len(d_1679_v17_))) - (default__.fm25(p0, p0, globalState))
            if self.f28:
                (globalState).f27 = True
                d_1680_v18_: _dafny.Array
                nw312_ = _dafny.Array(_dafny.MultiSet({}), 14)
                d_1680_v18_ = nw312_
                d_1681_v19_: D2
                d_1681_v19_ = D2_DC5(self.f35)
                d_1682_v20_: _dafny.Set
                d_1682_v20_ = _dafny.Set({default__.fm26(p0, globalState), (d_1673_v13_)[default__.safeIndex(105, (d_1673_v13_).length(0))]})
                d_1683_v21_: _dafny.Seq
                d_1683_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_1673_v13_)[default__.safeIndex(105, (d_1673_v13_).length(0))]}), d_1682_v20_])
                d_1684_v22_: _dafny.Set
                d_1684_v22_ = _dafny.Set({(d_1681_v19_).cf8, self.f35, default__.fm26(len(d_1683_v21_), globalState)})
                d_1685_v23_: _dafny.MultiSet
                d_1685_v23_ = _dafny.MultiSet([d_1684_v22_])
                index269_ = default__.safeIndex(806, (d_1680_v18_).length(0))
                (d_1680_v18_)[index269_] = d_1685_v23_
                d_1686_v24_: _dafny.Set
                d_1686_v24_ = _dafny.Set({d_1674_v14_})
                d_1687_v25_: _dafny.Set
                d_1687_v25_ = _dafny.Set({d_1654_v1_, d_1654_v1_})
                index270_ = default__.safeIndex(806, (d_1680_v18_).length(0))
                rhs293_ = (d_1685_v23_).intersection((_dafny.MultiSet([d_1682_v20_])) | (d_1685_v23_))
                rhs294_ = d_1686_v24_
                rhs295_ = not(not(((_dafny.Set({_dafny.SeqWithoutIsStrInference([self.f28])})) | (_dafny.Set({d_1654_v1_}))).isdisjoint(d_1687_v25_)))
                lhs242_ = d_1680_v18_
                lhs243_ = default__.safeIndex(806, (d_1680_v18_).length(0))
                lhs242_[lhs243_] = rhs293_
                d_1686_v24_ = rhs294_
                r0 = rhs295_
                default__.m0((d_1655_v2_) | (d_1655_v2_), (d_1668_v8_)[default__.safeIndex(425, len(d_1668_v8_))], (self).f29, globalState)
                d_1686_v24_ = d_1686_v24_
                (globalState).f13 = self.f35
            elif True:
                d_1688_v26_: C3
                nw313_ = C3()
                nw313_.ctor__((p0) < (p0), not((self).f29), (p0) <= (p0), p0)
                d_1688_v26_ = nw313_
                (globalState).f22 = ((d_1673_v13_)[default__.safeIndex(105, (d_1673_v13_).length(0))] if (p0) == (p0) else (True if self.f28 else self.f28))
                (self).f35 = self.f28
                d_1689_v27_: _dafny.Set
                d_1689_v27_ = _dafny.Set({d_1653_v0_, (d_1653_v0_) + (d_1653_v0_), d_1653_v0_, d_1653_v0_, (d_1653_v0_) + (d_1653_v0_)})
                d_1689_v27_ = (d_1689_v27_) | ((d_1689_v27_) | (d_1689_v27_))
                d_1690_v28_: _dafny.Map
                d_1690_v28_ = _dafny.Map({self.f36: self.f36})
                d_1690_v28_ = (d_1690_v28_).set(self.f36, self.f36)
            d_1670_v10_ = ((d_1670_v10_) | (d_1670_v10_)) | ((d_1670_v10_) | ((d_1671_v11_).cf12))
        elif True:
            d_1691_v29_: _dafny.Seq
            d_1691_v29_ = _dafny.SeqWithoutIsStrInference([d_1655_v2_])
            d_1691_v29_ = d_1691_v29_
            (globalState).f27 = (self).f29
            if True:
                d_1692_v30_: _dafny.Seq
                d_1692_v30_ = _dafny.SeqWithoutIsStrInference([d_1654_v1_, d_1654_v1_, d_1654_v1_])
                d_1693_v31_: _dafny.Array
                nw314_ = _dafny.Array(None, 3)
                nw314_[int(0)] = (0) - (p0)
                nw314_[int(1)] = len((d_1692_v30_)[default__.safeIndex(p0, len(d_1692_v30_))])
                nw314_[int(2)] = p0
                d_1693_v31_ = nw314_
                d_1694_v32_: D1
                d_1694_v32_ = D1_DC2(p0, len(d_1668_v8_), d_1653_v0_, not(self.f35))
                index271_ = default__.safeIndex(148, (d_1693_v31_).length(0))
                (d_1693_v31_)[index271_] = (d_1694_v32_).cf3
                index272_ = default__.safeIndex(148, (d_1693_v31_).length(0))
                (d_1693_v31_)[index272_] = p0
                (globalState).f23 = self.f36
                d_1695_v33_: _dafny.Map
                d_1695_v33_ = _dafny.Map({self.f36: (d_1693_v31_)[default__.safeIndex(148, (d_1693_v31_).length(0))]})
                d_1695_v33_ = (d_1695_v33_).set(self.f36, default__.safeDivisionInt((d_1693_v31_)[default__.safeIndex(148, (d_1693_v31_).length(0))], p0))
                d_1696_v34_: _dafny.Map
                d_1696_v34_ = _dafny.Map({(d_1693_v31_)[default__.safeIndex(148, (d_1693_v31_).length(0))]: p0})
                (globalState).f1 = len((d_1696_v34_) | (d_1696_v34_))
                d_1697_v35_: _dafny.Map
                d_1697_v35_ = _dafny.Map({(d_1693_v31_)[default__.safeIndex(148, (d_1693_v31_).length(0))]: _dafny.MultiSet([(d_1693_v31_)[default__.safeIndex(148, (d_1693_v31_).length(0))], (d_1693_v31_)[default__.safeIndex(148, (d_1693_v31_).length(0))]])})
                d_1698_v36_: _dafny.Map
                d_1698_v36_ = _dafny.Map({(self).f29: d_1697_v35_})
                d_1699_v37_: D13
                d_1699_v37_ = D13_DC34(((d_1698_v36_)[self.f35] if (self.f35) in (d_1698_v36_) else d_1697_v35_))
                d_1699_v37_ = d_1699_v37_
            elif True:
                (globalState).f27 = default__.fm29(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f29])]), p0, globalState)
                (globalState).f27 = not(((self.f35 if False else False) if self.f35 else self.f28))
                d_1700_v38_: C4
                nw315_ = C4()
                nw315_.ctor__(self.f28, self.f35, self.f28, (self.f28) or (self.f28))
                d_1700_v38_ = nw315_
                default__.m0(_dafny.Map({self.f28: (d_1700_v38_).f38}), p0, (d_1700_v38_).f38, globalState)
                (globalState).f27 = (d_1654_v1_)[default__.safeIndex(p0, len(d_1654_v1_))]
            (globalState).f1 = (((0) - (p0)) * (p0)) + ((0) - (p0))
            source29_ = D2_DC4(self.f36)
            if source29_.is_DC5:
                d_1701___mcc_h2_ = source29_.cf8
                d_1702_cf8_ = d_1701___mcc_h2_
                d_1703_v39_: _dafny.Array
                nw316_ = _dafny.Array(_dafny.Map({}), 8)
                d_1703_v39_ = nw316_
                d_1704_v40_: _dafny.Map
                d_1704_v40_ = _dafny.Map({p0: self.f28})
                index273_ = default__.safeIndex(611, (d_1703_v39_).length(0))
                (d_1703_v39_)[index273_] = _dafny.Map({len(d_1704_v40_): p0})
                d_1705_v41_: _dafny.Map
                d_1705_v41_ = _dafny.Map({p0: p0})
                index274_ = default__.safeIndex(611, (d_1703_v39_).length(0))
                def iife127_():
                    coll69_ = _dafny.Map()
                    compr_69_: int
                    for compr_69_ in _dafny.IntegerRange(564, 621):
                        d_1706_v42_: int = compr_69_
                        if ((564) <= (d_1706_v42_)) and ((d_1706_v42_) < (621)):
                            coll69_[default__.safeDivisionInt(d_1706_v42_, p0)] = len(d_1655_v2_)
                    return _dafny.Map(coll69_)
                rhs296_ = self.f28
                rhs297_ = ((d_1705_v41_) | (iife127_()
                ) if (self).f29 else default__.fm54(self.f35, (0) - (p0), p0, d_1702_cf8_, globalState))
                rhs298_ = (((d_1668_v8_) + (d_1668_v8_)) + (d_1668_v8_)).set(default__.safeIndex(default__.safeModuloInt(p0, (0) - (p0)), len(((d_1668_v8_) + (d_1668_v8_)) + (d_1668_v8_))), p0)
                rhs299_ = self.f35
                lhs244_ = self
                lhs245_ = d_1703_v39_
                lhs246_ = default__.safeIndex(611, (d_1703_v39_).length(0))
                lhs247_ = globalState
                lhs248_ = globalState
                lhs244_.f28 = rhs296_
                lhs245_[lhs246_] = rhs297_
                lhs247_.f17 = rhs298_
                lhs248_.f21 = rhs299_
                (globalState).f23 = self.f36
                (globalState).f1 = default__.safeModuloInt(p0, 979)
                d_1707_v43_: _dafny.Set
                d_1707_v43_ = _dafny.Set({d_1653_v0_, _dafny.SeqWithoutIsStrInference([self.f36 for d_1708_i3_ in range(default__.abs(911))]), d_1653_v0_, d_1653_v0_})
                d_1707_v43_ = _dafny.Set({(d_1653_v0_) + (d_1653_v0_)})
            elif source29_.is_DC4:
                d_1709___mcc_h3_ = source29_.cf7
                d_1710_cf7_ = d_1709___mcc_h3_
                d_1711_v44_: _dafny.MultiSet
                d_1711_v44_ = _dafny.MultiSet([(d_1668_v8_) + (d_1668_v8_)])
                (globalState).f24 = ((d_1711_v44_)[d_1668_v8_] if (d_1668_v8_) in (d_1711_v44_) else default__.safeModuloInt(len(d_1653_v0_), p0))
                d_1712_v45_: C1
                nw317_ = C1()
                nw317_.ctor__(self.f35, not(self.f28))
                d_1712_v45_ = nw317_
                d_1713_v46_: _dafny.Array
                nw318_ = _dafny.Array(int(0), 9)
                d_1713_v46_ = nw318_
                d_1714_v47_: _dafny.Map
                d_1714_v47_ = _dafny.Map({self.f36: len(d_1668_v8_)})
                d_1715_v48_: _dafny.Set
                d_1715_v48_ = _dafny.Set({p0, p0})
                d_1716_v49_: D12
                d_1716_v49_ = D12_DC31(d_1715_v48_)
                d_1717_v50_: _dafny.Array
                d_1717_v50_ = d_1663_v6_
                d_1718_v51_: _dafny.Set
                d_1718_v51_ = _dafny.Set({d_1717_v50_})
                d_1719_v52_: _dafny.MultiSet
                d_1719_v52_ = _dafny.MultiSet([d_1654_v1_, d_1654_v1_])
                d_1720_v53_: _dafny.Array
                nw319_ = _dafny.Array(None, 21)
                nw319_[int(0)] = (d_1653_v0_) < (d_1653_v0_)
                nw319_[int(1)] = (self).f29
                nw319_[int(2)] = self.f35
                nw319_[int(3)] = (p0) >= (p0)
                nw319_[int(4)] = self.f28
                nw319_[int(5)] = (d_1653_v0_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpcsblw")))
                nw319_[int(6)] = (self.f36) in (d_1714_v47_)
                nw319_[int(7)] = ((self).f29) or (self.f28)
                nw319_[int(8)] = self.f28
                nw319_[int(9)] = not(self.f35)
                nw319_[int(10)] = self.f35
                nw319_[int(11)] = self.f28
                nw319_[int(12)] = (self).f29
                nw319_[int(13)] = (self).f29
                nw319_[int(14)] = (d_1715_v48_).ispropersubset((d_1716_v49_).cf48)
                nw319_[int(15)] = self.f28
                nw319_[int(16)] = (p0) < (p0)
                nw319_[int(17)] = (d_1718_v51_).ispropersubset(d_1718_v51_)
                nw319_[int(18)] = ((self).f29) == (self.f28)
                nw319_[int(19)] = default__.fm29((D15_DC37(d_1719_v52_)).cf52, p0, globalState)
                nw319_[int(20)] = True
                d_1720_v53_ = nw319_
                rhs300_ = d_1713_v46_
                rhs301_ = self.f28
                rhs302_ = d_1720_v53_
                lhs249_ = globalState
                d_1713_v46_ = rhs300_
                lhs249_.f27 = rhs301_
                d_1720_v53_ = rhs302_
                d_1721_v54_: _dafny.MultiSet
                d_1721_v54_ = _dafny.MultiSet([d_1653_v0_])
                d_1722_v56_: C3
                nw320_ = C3()
                def iife128_():
                    coll70_ = _dafny.Set()
                    compr_70_: _dafny.Seq
                    for compr_70_ in (d_1721_v54_).Elements:
                        d_1723_v55_: _dafny.Seq = compr_70_
                        if (d_1723_v55_) in (d_1721_v54_):
                            coll70_ = coll70_.union(_dafny.Set([d_1723_v55_]))
                    return _dafny.Set(coll70_)
                nw320_.ctor__(self.f28, not ((self).f29) or (self.f28), (d_1654_v1_)[default__.safeIndex(p0, len(d_1654_v1_))], ((d_1712_v45_).fm36(len(default__.fm45(globalState)), iife128_()
                , globalState)) - (p0))
                d_1722_v56_ = nw320_
            elif True:
                d_1724___mcc_h4_ = source29_.cf9
                d_1725_cf9_ = d_1724___mcc_h4_
                d_1726_v57_: _dafny.Array
                nw321_ = _dafny.Array(None, 7)
                nw321_[int(0)] = self.f28
                nw321_[int(1)] = (self).f29
                nw321_[int(2)] = False
                nw321_[int(3)] = (self).f29
                nw321_[int(4)] = self.f35
                nw321_[int(5)] = self.f35
                nw321_[int(6)] = (self).f29
                d_1726_v57_ = nw321_
                d_1727_v58_: _dafny.Map
                d_1727_v58_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([704 for d_1728_i4_ in range(default__.abs(-471))]): p0})
                d_1729_v59_: _dafny.Map
                d_1729_v59_ = _dafny.Map({d_1726_v57_: d_1727_v58_})
                d_1730_v60_: D2
                d_1730_v60_ = D2_DC4(self.f36)
                d_1731_v61_: D2
                d_1731_v61_ = D2_DC4((d_1730_v60_).cf7)
                d_1732_v62_: D5
                d_1732_v62_ = D5_DC13(d_1730_v60_)
                (globalState).f19 = (((d_1729_v59_)[d_1726_v57_] if (d_1726_v57_) in (d_1729_v59_) else default__.fm63(d_1732_v62_, len(d_1653_v0_), globalState))) | (d_1727_v58_)
                (globalState).f24 = p0
                d_1656_v3_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "firud"))})
                default__.m0(d_1655_v2_, (0) - (p0), True, globalState)
        d_1733_v63_: _dafny.Map
        d_1733_v63_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajggj")): (d_1653_v0_)[default__.safeIndex((d_1670_v10_).cardinality, len(d_1653_v0_))]})
        d_1734_v65_: _dafny.Seq
        d_1734_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyscfksv"))])
        def iife129_():
            coll71_ = _dafny.Map()
            compr_71_: _dafny.Seq
            for compr_71_ in (d_1734_v65_).Elements:
                d_1735_v64_: _dafny.Seq = compr_71_
                if (d_1735_v64_) in (d_1734_v65_):
                    coll71_[d_1735_v64_] = self.f36
            return _dafny.Map(coll71_)
        d_1655_v2_ = (d_1655_v2_).set((d_1733_v63_) != (iife129_()
        ), (d_1654_v1_)[default__.safeIndex(len(d_1653_v0_), len(d_1654_v1_))])
        default__.m0((d_1655_v2_) | (d_1655_v2_), ((0) - ((0) - (p0))) * (p0), not (self.f35) or ((self).f29), globalState)
        r0 = self.f28
        return r0

    def m12(self, p0, p1, globalState):
        d_1736_v0_: _dafny.Array
        nw322_ = _dafny.Array(int(0), 17)
        d_1736_v0_ = nw322_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1736_v0_).length(0)):
            d_1737_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_1737_i0_)) and ((d_1737_i0_) < ((d_1736_v0_).length(0)))):
                (d_1736_v0_)[(d_1737_i0_)] = (d_1737_i0_) + (p0)
        (globalState).f24 = (0) - (p1)
        hi5_ = (p1 if self.f35 else p0)
        for d_1738_i1_ in range(p1, hi5_):
            d_1739_v1_: _dafny.Seq
            d_1739_v1_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            d_1740_v2_: _dafny.Seq
            d_1740_v2_ = d_1739_v1_
            d_1741_v3_: D4
            d_1741_v3_ = D4_DC10(self.f35, d_1738_i1_, len((d_1740_v2_)))
            d_1742_v4_: _dafny.Seq
            d_1742_v4_ = _dafny.SeqWithoutIsStrInference([(d_1741_v3_).cf15, d_1738_i1_])
            d_1743_v5_: _dafny.MultiSet
            d_1743_v5_ = _dafny.MultiSet([default__.fm26((d_1742_v4_)[default__.safeIndex(-875, len(d_1742_v4_))], globalState), self.f28, self.f35])
            index275_ = default__.safeIndex(91, (d_1736_v0_).length(0))
            (d_1736_v0_)[index275_] = p1
            index276_ = default__.safeIndex(91, (d_1736_v0_).length(0))
            rhs303_ = d_1743_v5_
            rhs304_ = (p0) - (d_1738_i1_)
            rhs305_ = (d_1739_v1_) + (d_1739_v1_)
            rhs306_ = _dafny.CodePoint('k')
            rhs307_ = (0) - ((0) - (default__.fm25(d_1738_i1_, p0, globalState)))
            lhs250_ = d_1736_v0_
            lhs251_ = default__.safeIndex(91, (d_1736_v0_).length(0))
            lhs252_ = globalState
            lhs253_ = globalState
            d_1743_v5_ = rhs303_
            lhs250_[lhs251_] = rhs304_
            d_1739_v1_ = rhs305_
            lhs252_.f23 = rhs306_
            lhs253_.f24 = rhs307_
            if (self.f35) and (self.f35):
                d_1744_v6_: _dafny.Seq
                d_1744_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxo"))
                (globalState).f1 = len((d_1744_v6_) + ((_dafny.SeqWithoutIsStrInference([self.f36 for d_1745_i2_ in range(default__.abs(154))]) if self.f28 else d_1744_v6_)))
                d_1746_v7_: _dafny.Set
                d_1746_v7_ = _dafny.Set({self.f35, self.f28, self.f28, self.f28, self.f35})
                (globalState).f1 = len(d_1746_v7_)
                (globalState).f27 = self.f35
                d_1747_v8_: _dafny.Seq
                d_1747_v8_ = _dafny.SeqWithoutIsStrInference([d_1746_v7_, (d_1746_v7_ if self.f28 else d_1746_v7_), d_1746_v7_])
                d_1747_v8_ = (_dafny.SeqWithoutIsStrInference([d_1746_v7_ for d_1748_i3_ in range(default__.abs(-776))])) + (d_1747_v8_)
                d_1749_v9_: _dafny.Array
                def lambda63_(d_1750_v0_):
                    def lambda64_(d_1751_i4_):
                        return (_dafny.Map({(self).f29: _dafny.Map({True: (d_1750_v0_)[default__.safeIndex(91, (d_1750_v0_).length(0))]})})) | (_dafny.Map({(self).f29: _dafny.Map({self.f28: len(_dafny.Set({self.f28}))})}))

                    return lambda64_

                init37_ = lambda63_(d_1736_v0_)
                nw323_ = _dafny.Array(None, 14)
                for i0_37_ in range(nw323_.length(0)):
                    nw323_[i0_37_] = init37_(i0_37_)
                d_1749_v9_ = nw323_
                d_1752_v10_: _dafny.Map
                d_1752_v10_ = _dafny.Map({self.f35: p1})
                d_1753_v11_: _dafny.Map
                d_1753_v11_ = _dafny.Map({self.f28: d_1752_v10_})
                index277_ = default__.safeIndex(480, (d_1749_v9_).length(0))
                (d_1749_v9_)[index277_] = d_1753_v11_
                index278_ = default__.safeIndex(480, (d_1749_v9_).length(0))
                (d_1749_v9_)[index278_] = d_1753_v11_
            elif True:
                d_1754_v12_: _dafny.Seq
                d_1754_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnlgjibeg"))
                d_1755_v13_: D3
                d_1755_v13_ = D3_DC8(d_1754_v12_)
                d_1756_v14_: _dafny.Set
                d_1756_v14_ = _dafny.Set({d_1755_v13_, d_1755_v13_})
                d_1757_v15_: _dafny.Seq
                d_1757_v15_ = _dafny.SeqWithoutIsStrInference([d_1756_v14_])
                d_1758_v16_: _dafny.Map
                d_1758_v16_ = _dafny.Map({d_1757_v15_: d_1738_i1_})
                d_1759_v17_: _dafny.Set
                d_1759_v17_ = _dafny.Set({self.f35})
                d_1760_v18_: _dafny.Map
                d_1760_v18_ = _dafny.Map({len(d_1759_v17_): self.f36})
                d_1761_v19_: _dafny.Map
                d_1761_v19_ = _dafny.Map({(d_1736_v0_)[default__.safeIndex(91, (d_1736_v0_).length(0))]: self.f35})
                d_1762_v20_: _dafny.MultiSet
                d_1762_v20_ = _dafny.MultiSet([_dafny.Map({(0) - (d_1738_i1_): self.f36}), d_1760_v18_, default__.fm64(len(d_1761_v19_), (0) - (d_1738_i1_), globalState)])
                d_1758_v16_ = (d_1758_v16_).set((d_1757_v15_ if self.f35 else d_1757_v15_), ((d_1762_v20_).cardinality) + (p0))
                d_1754_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqd"))
                d_1763_v21_: _dafny.Map
                d_1763_v21_ = _dafny.Map({self.f28: len(_dafny.Map({(d_1736_v0_)[default__.safeIndex(91, (d_1736_v0_).length(0))]: 234}))})
                d_1764_v22_: _dafny.MultiSet
                d_1764_v22_ = _dafny.MultiSet([len(d_1739_v1_)])
                d_1765_v23_: C4
                nw324_ = C4()
                nw324_.ctor__((847) < (p0), (p1) > (((d_1763_v21_)[False] if (False) in (d_1763_v21_) else len(d_1742_v4_))), self.f28, (d_1764_v22_).issubset(default__.fm24(default__.fm35(globalState), self.f28, globalState)))
                d_1765_v23_ = nw324_
                d_1766_v24_: _dafny.Map
                d_1766_v24_ = _dafny.Map({default__.fm51(p0, (d_1736_v0_)[default__.safeIndex(91, (d_1736_v0_).length(0))], globalState): self.f28})
                index279_ = default__.safeIndex(91, (d_1736_v0_).length(0))
                (d_1736_v0_)[index279_] = ((d_1736_v0_)[default__.safeIndex(91, (d_1736_v0_).length(0))] if not(((d_1766_v24_)[d_1754_v12_] if (d_1754_v12_) in (d_1766_v24_) else (d_1765_v23_).f37)) else p0)
                (globalState).f24 = p1
            (globalState).f27 = self.f28
            d_1767_v25_: T0
            nw325_ = C3()
            nw325_.ctor__(self.f28, self.f28, not(self.f35), (0) - (d_1738_i1_))
            d_1767_v25_ = nw325_
            d_1768_v26_: D8
            d_1768_v26_ = D8_DC20(d_1767_v25_)
            d_1768_v26_ = d_1768_v26_
        d_1769_v27_: _dafny.Map
        d_1769_v27_ = _dafny.Map({p1: p0})
        d_1770_v29_: _dafny.Map
        d_1770_v29_ = _dafny.Map({(self).f29: self.f35})
        d_1771_v30_: _dafny.Seq
        d_1771_v30_ = _dafny.SeqWithoutIsStrInference([d_1770_v29_])
        d_1772_v31_: _dafny.Seq
        d_1772_v31_ = _dafny.SeqWithoutIsStrInference([(d_1771_v30_)[default__.safeIndex(p0, len(d_1771_v30_))]])
        d_1773_v32_: _dafny.MultiSet
        def iife130_():
            coll72_ = _dafny.Map()
            compr_72_: _dafny.Map
            for compr_72_ in (d_1772_v31_).Elements:
                d_1774_v28_: _dafny.Map = compr_72_
                if (d_1774_v28_) in (d_1772_v31_):
                    coll72_[d_1774_v28_] = p1
            return _dafny.Map(coll72_)
        d_1773_v32_ = _dafny.MultiSet([len(d_1769_v27_), len(iife130_()
        ), p1])
        hi6_ = ((d_1773_v32_) | (d_1773_v32_)).cardinality
        for d_1775_i5_ in range(397, hi6_):
            d_1776_v33_: _dafny.Map
            d_1776_v33_ = _dafny.Map({self.f28: -277})
            default__.m0(_dafny.Map({self.f35: self.f28}), default__.safeDivisionInt(p0, default__.fm25(p1, 485, globalState)), (p0) <= (len(d_1776_v33_)), globalState)
            d_1777_v34_: _dafny.Map
            d_1777_v34_ = _dafny.Map({(p1) + ((0) - (d_1775_i5_)): (default__.fm65((self).f29, self.f35, globalState)).cf20})
            d_1778_v35_: _dafny.Array
            def lambda65_(d_1779_i6_):
                return not(self.f35)

            init38_ = lambda65_
            nw326_ = _dafny.Array(None, 27)
            for i0_38_ in range(nw326_.length(0)):
                nw326_[i0_38_] = init38_(i0_38_)
            d_1778_v35_ = nw326_
            index280_ = default__.safeIndex(916, (d_1778_v35_).length(0))
            (d_1778_v35_)[index280_] = not(default__.fm38(globalState))
            index281_ = default__.safeIndex(406, (d_1736_v0_).length(0))
            (d_1736_v0_)[index281_] = d_1775_i5_
            d_1780_v36_: _dafny.Array
            def lambda66_(d_1781_i7_):
                return D2_DC4(self.f36)

            init39_ = lambda66_
            nw327_ = _dafny.Array(None, 10)
            for i0_39_ in range(nw327_.length(0)):
                nw327_[i0_39_] = init39_(i0_39_)
            d_1780_v36_ = nw327_
            d_1782_v37_: D2
            d_1782_v37_ = D2_DC4(self.f36)
            index282_ = default__.safeIndex(763, (d_1780_v36_).length(0))
            (d_1780_v36_)[index282_] = d_1782_v37_
            d_1783_v39_: _dafny.Seq
            d_1783_v39_ = _dafny.SeqWithoutIsStrInference([self.f28])
            index283_ = default__.safeIndex(916, (d_1778_v35_).length(0))
            index284_ = default__.safeIndex(406, (d_1736_v0_).length(0))
            index285_ = default__.safeIndex(763, (d_1780_v36_).length(0))
            def iife131_():
                coll73_ = _dafny.Map()
                compr_73_: int
                for compr_73_ in _dafny.IntegerRange(614, -847):
                    d_1784_v38_: int = compr_73_
                    if ((614) <= (d_1784_v38_)) and ((d_1784_v38_) < (-847)):
                        coll73_[(d_1784_v38_) - (p1)] = self.f35
                return _dafny.Map(coll73_)
            rhs308_ = iife131_()

            rhs309_ = not(self.f35)
            rhs310_ = (d_1783_v39_) != ((d_1783_v39_) + (d_1783_v39_))
            rhs311_ = p1
            rhs312_ = d_1782_v37_
            lhs254_ = d_1778_v35_
            lhs255_ = default__.safeIndex(916, (d_1778_v35_).length(0))
            lhs256_ = globalState
            lhs257_ = d_1736_v0_
            lhs258_ = default__.safeIndex(406, (d_1736_v0_).length(0))
            lhs259_ = d_1780_v36_
            lhs260_ = default__.safeIndex(763, (d_1780_v36_).length(0))
            d_1777_v34_ = rhs308_
            lhs254_[lhs255_] = rhs309_
            lhs256_.f13 = rhs310_
            lhs257_[lhs258_] = rhs311_
            lhs259_[lhs260_] = rhs312_
            index286_ = default__.safeIndex(406, (d_1736_v0_).length(0))
            rhs313_ = p0
            rhs314_ = d_1775_i5_
            rhs315_ = (d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))]
            lhs261_ = d_1736_v0_
            lhs262_ = default__.safeIndex(406, (d_1736_v0_).length(0))
            lhs263_ = globalState
            lhs264_ = globalState
            lhs261_[lhs262_] = rhs313_
            lhs263_.f1 = rhs314_
            lhs264_.f22 = rhs315_
            source30_ = D10_DC28(d_1770_v29_)
            if source30_.is_DC27:
                d_1785___mcc_h0_ = source30_.cf40
                d_1786___mcc_h1_ = source30_.cf41
                d_1787___mcc_h2_ = source30_.cf42
                d_1788___mcc_h3_ = source30_.cf43
                d_1789___mcc_h4_ = source30_.cf44
                d_1790_cf44_ = d_1789___mcc_h4_
                d_1791_cf43_ = d_1788___mcc_h3_
                d_1792_cf42_ = d_1787___mcc_h2_
                d_1793_cf41_ = d_1786___mcc_h1_
                d_1794_cf40_ = d_1785___mcc_h0_
                d_1795_v40_: _dafny.Map
                d_1795_v40_ = _dafny.Map({self.f35: self.f36})
                d_1796_v41_: _dafny.Map
                d_1796_v41_ = _dafny.Map({(d_1795_v40_) | (d_1795_v40_): p1})
                d_1796_v41_ = d_1796_v41_
                (globalState).f22 = d_1792_cf42_
                d_1797_v42_: _dafny.Array
                def lambda67_(d_1798_v34_):
                    def lambda68_(d_1799_i8_):
                        return default__.safeModuloInt(d_1799_i8_, len(d_1798_v34_))

                    return lambda68_

                init40_ = lambda67_(d_1777_v34_)
                nw328_ = _dafny.Array(None, 26)
                for i0_40_ in range(nw328_.length(0)):
                    nw328_[i0_40_] = init40_(i0_40_)
                d_1797_v42_ = nw328_
                d_1797_v42_ = d_1797_v42_
                (globalState).f22 = (self.f36) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1800_i9_ in range(default__.abs(417))]))
            elif source30_.is_DC28:
                d_1801___mcc_h5_ = source30_.cf45
                d_1802_cf45_ = d_1801___mcc_h5_
                d_1803_v43_: C1
                nw329_ = C1()
                nw329_.ctor__(default__.fm38(globalState), not((d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))]))
                d_1803_v43_ = nw329_
                d_1804_v44_: _dafny.Set
                d_1804_v44_ = _dafny.Set({d_1803_v43_})
                d_1805_v45_: _dafny.Seq
                d_1805_v45_ = _dafny.SeqWithoutIsStrInference([d_1804_v44_, d_1804_v44_])
                d_1806_v46_: _dafny.Map
                d_1806_v46_ = _dafny.Map({self.f28: d_1804_v44_})
                (globalState).f22 = ((d_1804_v44_) | (d_1804_v44_)) in ((d_1805_v45_).set(default__.safeIndex(d_1775_i5_, len(d_1805_v45_)), ((d_1806_v46_)[(d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))]] if ((d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))]) in (d_1806_v46_) else d_1804_v44_)))
                index287_ = default__.safeIndex(916, (d_1778_v35_).length(0))
                (d_1778_v35_)[index287_] = ((self).f29 if (True) and ((self).f29) else (self).f29)
                d_1807_v47_: _dafny.Array
                nw330_ = _dafny.Array(_dafny.Seq({}), 5)
                d_1807_v47_ = nw330_
                d_1808_v48_: _dafny.Seq
                d_1808_v48_ = d_1783_v39_
                index288_ = default__.safeIndex(944, (d_1807_v47_).length(0))
                (d_1807_v47_)[index288_] = d_1808_v48_
                d_1809_v49_: _dafny.Seq
                d_1809_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kytrfhqy"))
                index289_ = default__.safeIndex(944, (d_1807_v47_).length(0))
                (d_1807_v47_)[index289_] = default__.fm66(self.f36, p0, len((_dafny.SeqWithoutIsStrInference([d_1809_v49_])).set(default__.safeIndex((0) - (p1), len(_dafny.SeqWithoutIsStrInference([d_1809_v49_]))), d_1809_v49_)), -123, globalState)
                d_1810_v50_: _dafny.Array
                nw331_ = _dafny.Array(D7.default()(), 22)
                d_1810_v50_ = nw331_
                index290_ = default__.safeIndex(487, (d_1810_v50_).length(0))
                (d_1810_v50_)[index290_] = default__.fm67(self.f35, globalState)
                d_1811_v51_: _dafny.Map
                d_1811_v51_ = _dafny.Map({self.f36: self.f35})
                index291_ = default__.safeIndex(487, (d_1810_v50_).length(0))
                (d_1810_v50_)[index291_] = default__.fm67((False if True else ((d_1811_v51_)[self.f36] if (self.f36) in (d_1811_v51_) else (self).f29)), globalState)
            elif source30_.is_DC26:
                d_1812___mcc_h6_ = source30_.cf39
                d_1813_cf39_ = d_1812___mcc_h6_
                d_1814_v52_: _dafny.Array
                def lambda69_(d_1815_v35_):
                    def lambda70_(d_1816_i10_):
                        return D15_DC41(D15_DC40((d_1815_v35_)[default__.safeIndex(916, (d_1815_v35_).length(0))], (self).f29))

                    return lambda70_

                init41_ = lambda69_(d_1778_v35_)
                nw332_ = _dafny.Array(None, 1)
                for i0_41_ in range(nw332_.length(0)):
                    nw332_[i0_41_] = init41_(i0_41_)
                d_1814_v52_ = nw332_
                d_1817_v53_: _dafny.Array
                nw333_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_1817_v53_ = nw333_
                d_1818_v54_: _dafny.Map
                d_1818_v54_ = _dafny.Map({d_1814_v52_: d_1817_v53_})
                d_1818_v54_ = (d_1818_v54_).set(d_1814_v52_, d_1817_v53_)
                d_1819_v55_: _dafny.Seq
                d_1819_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
                d_1820_v56_: _dafny.Seq
                d_1820_v56_ = _dafny.SeqWithoutIsStrInference([d_1819_v55_])
                d_1821_v57_: _dafny.Array
                nw334_ = _dafny.Array(None, 5)
                nw334_[int(0)] = d_1820_v56_
                nw334_[int(1)] = (d_1820_v56_ if (self).f29 else d_1820_v56_)
                nw334_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_1819_v55_])) + (_dafny.SeqWithoutIsStrInference([d_1819_v55_, d_1819_v55_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsuxt")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vevqhheot"))).set(default__.safeIndex(default__.fm27(d_1775_i5_, self.f36, d_1775_i5_, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vevqhheot")))), self.f36), d_1819_v55_]))
                nw334_[int(3)] = d_1820_v56_
                nw334_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_1819_v55_).set(default__.safeIndex(d_1775_i5_, len(d_1819_v55_)), self.f36)])
                d_1821_v57_ = nw334_
                index292_ = default__.safeIndex(195, (d_1821_v57_).length(0))
                (d_1821_v57_)[index292_] = _dafny.SeqWithoutIsStrInference([d_1819_v55_ for d_1822_i11_ in range(default__.abs(0))])
                index293_ = default__.safeIndex(195, (d_1821_v57_).length(0))
                (d_1821_v57_)[index293_] = d_1820_v56_
                d_1823_v64_: _dafny.Array
                nw335_ = _dafny.Array(None, 24)
                nw335_[int(0)] = default__.fm54(self.f28, (d_1736_v0_)[default__.safeIndex(406, (d_1736_v0_).length(0))], d_1775_i5_, False, globalState)
                def iife132_():
                    coll74_ = _dafny.Map()
                    compr_74_: int
                    for compr_74_ in (_dafny.SeqWithoutIsStrInference([p0 for d_1824_i12_ in range(default__.abs(65))])).Elements:
                        d_1825_v58_: int = compr_74_
                        if (d_1825_v58_) in (_dafny.SeqWithoutIsStrInference([p0 for d_1824_i12_ in range(default__.abs(65))])):
                            coll74_[(d_1825_v58_) + (p0)] = (d_1736_v0_)[default__.safeIndex(406, (d_1736_v0_).length(0))]
                    return _dafny.Map(coll74_)
                nw335_[int(1)] = (iife132_()
                ) | (d_1769_v27_)
                nw335_[int(2)] = d_1769_v27_
                def iife133_():
                    coll75_ = _dafny.Map()
                    compr_75_: int
                    for compr_75_ in _dafny.IntegerRange(-141, 153):
                        d_1826_v59_: int = compr_75_
                        if ((-141) <= (d_1826_v59_)) and ((d_1826_v59_) < (153)):
                            coll75_[(d_1826_v59_) + (p0)] = len(_dafny.SeqWithoutIsStrInference([len(d_1783_v39_), d_1775_i5_]))
                    return _dafny.Map(coll75_)
                nw335_[int(3)] = (d_1769_v27_ if default__.fm26(p0, globalState) else iife133_()
                )
                nw335_[int(4)] = d_1769_v27_
                nw335_[int(5)] = _dafny.Map({p1: p0})
                def iife134_():
                    coll76_ = _dafny.Map()
                    compr_76_: int
                    for compr_76_ in _dafny.IntegerRange(271, 917):
                        d_1827_v60_: int = compr_76_
                        if ((271) <= (d_1827_v60_)) and ((d_1827_v60_) < (917)):
                            coll76_[(d_1827_v60_) + (d_1775_i5_)] = p0
                    return _dafny.Map(coll76_)
                nw335_[int(6)] = iife134_()
                
                nw335_[int(7)] = d_1769_v27_
                nw335_[int(8)] = d_1769_v27_
                nw335_[int(9)] = d_1769_v27_
                def iife135_():
                    def iife137_():
                        coll79_ = _dafny.Map()
                        compr_79_: int
                        for compr_79_ in _dafny.IntegerRange(-613, -196):
                            d_1828_v62_: int = compr_79_
                            if ((-613) <= (d_1828_v62_)) and ((d_1828_v62_) < (-196)):
                                coll79_[(d_1828_v62_) + (d_1775_i5_)] = d_1819_v55_
                        return _dafny.Map(coll79_)
                    coll77_ = _dafny.Map()
                    def iife136_():
                        coll78_ = _dafny.Map()
                        compr_78_: int
                        for compr_78_ in _dafny.IntegerRange(-613, -196):
                            d_1828_v62_: int = compr_78_
                            if ((-613) <= (d_1828_v62_)) and ((d_1828_v62_) < (-196)):
                                coll78_[(d_1828_v62_) + (d_1775_i5_)] = d_1819_v55_
                        return _dafny.Map(coll78_)
                    compr_77_: int
                    for compr_77_ in (iife136_()
                    ).keys.Elements:
                        d_1829_v61_: int = compr_77_
                        if (d_1829_v61_) in (iife137_()
                        ):
                            coll77_[(d_1829_v61_) - (len(d_1783_v39_))] = p0
                    return _dafny.Map(coll77_)
                nw335_[int(10)] = (iife135_()
                 if (d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))] else d_1769_v27_)
                nw335_[int(11)] = d_1769_v27_
                nw335_[int(12)] = _dafny.Map({p0: len(d_1819_v55_)})
                nw335_[int(13)] = (d_1769_v27_) | (d_1769_v27_)
                nw335_[int(14)] = d_1769_v27_
                nw335_[int(15)] = d_1769_v27_
                nw335_[int(16)] = d_1769_v27_
                nw335_[int(17)] = d_1769_v27_
                nw335_[int(18)] = d_1769_v27_
                nw335_[int(19)] = d_1769_v27_
                def iife138_():
                    coll80_ = _dafny.Map()
                    compr_80_: int
                    for compr_80_ in _dafny.IntegerRange(-474, 949):
                        d_1830_v63_: int = compr_80_
                        if ((-474) <= (d_1830_v63_)) and ((d_1830_v63_) < (949)):
                            coll80_[default__.safeModuloInt(d_1830_v63_, (d_1736_v0_)[default__.safeIndex(406, (d_1736_v0_).length(0))])] = 499
                    return _dafny.Map(coll80_)
                nw335_[int(20)] = iife138_()
                
                nw335_[int(21)] = d_1769_v27_
                nw335_[int(22)] = d_1769_v27_
                nw335_[int(23)] = _dafny.Map({(0) - (d_1775_i5_): p0})
                d_1823_v64_ = nw335_
                index294_ = default__.safeIndex(406, (d_1736_v0_).length(0))
                rhs316_ = d_1823_v64_
                rhs317_ = d_1775_i5_
                rhs318_ = p0
                rhs319_ = p0
                rhs320_ = self.f36
                lhs265_ = d_1736_v0_
                lhs266_ = default__.safeIndex(406, (d_1736_v0_).length(0))
                lhs267_ = globalState
                lhs268_ = globalState
                lhs269_ = globalState
                d_1823_v64_ = rhs316_
                lhs265_[lhs266_] = rhs317_
                lhs267_.f1 = rhs318_
                lhs268_.f1 = rhs319_
                lhs269_.f23 = rhs320_
                d_1831_v65_: _dafny.MultiSet
                d_1831_v65_ = _dafny.MultiSet([default__.fm26(512, globalState), True, True, (d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))], self.f28])
                d_1831_v65_ = d_1831_v65_
            elif True:
                d_1832___mcc_h7_ = source30_.cf46
                d_1833_cf46_ = d_1832___mcc_h7_
                d_1834_v66_: _dafny.Seq
                d_1834_v66_ = _dafny.SeqWithoutIsStrInference([d_1775_i5_, 24])
                (globalState).f17 = (d_1834_v66_) + (_dafny.SeqWithoutIsStrInference([d_1775_i5_ for d_1835_i13_ in range(default__.abs(848))]))
                d_1836_v67_: _dafny.Seq
                d_1836_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayjrff"))
                d_1837_v68_: _dafny.Map
                d_1837_v68_ = _dafny.Map({len(d_1836_v67_): d_1836_v67_})
                d_1838_v69_: _dafny.Set
                d_1838_v69_ = _dafny.Set({len(d_1837_v68_), len(d_1769_v27_), p0})
                d_1839_v70_: D12
                d_1839_v70_ = D12_DC31(d_1838_v69_)
                (globalState).f1 = len((d_1838_v69_) | ((d_1839_v70_).cf48))
                d_1840_v71_: _dafny.Array
                nw336_ = _dafny.Array(D7.default()(), 13)
                d_1840_v71_ = nw336_
                d_1840_v71_ = d_1840_v71_
                d_1841_v72_: _dafny.MultiSet
                d_1841_v72_ = _dafny.MultiSet([d_1834_v66_, d_1834_v66_])
                d_1842_v73_: _dafny.Map
                d_1842_v73_ = _dafny.Map({d_1841_v72_: d_1778_v35_})
                d_1843_v74_: _dafny.Map
                d_1843_v74_ = _dafny.Map({self.f35: d_1778_v35_})
                d_1844_v75_: _dafny.Map
                d_1844_v75_ = _dafny.Map({d_1776_v33_: d_1773_v32_})
                d_1845_v76_: _dafny.Map
                d_1845_v76_ = _dafny.Map({len(d_1844_v75_): d_1778_v35_})
                d_1846_v77_: _dafny.Array
                nw337_ = _dafny.Array(None, 9)
                nw337_[int(0)] = d_1778_v35_
                nw337_[int(1)] = ((d_1842_v73_)[d_1841_v72_] if (d_1841_v72_) in (d_1842_v73_) else d_1778_v35_)
                nw337_[int(2)] = ((d_1843_v74_)[(d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))]] if ((d_1778_v35_)[default__.safeIndex(916, (d_1778_v35_).length(0))]) in (d_1843_v74_) else d_1778_v35_)
                nw337_[int(3)] = d_1778_v35_
                nw337_[int(4)] = d_1778_v35_
                nw337_[int(5)] = d_1778_v35_
                nw337_[int(6)] = d_1778_v35_
                nw337_[int(7)] = ((d_1845_v76_)[p0] if (p0) in (d_1845_v76_) else d_1778_v35_)
                nw337_[int(8)] = d_1778_v35_
                d_1846_v77_ = nw337_
                index295_ = default__.safeIndex(71, (d_1846_v77_).length(0))
                (d_1846_v77_)[index295_] = d_1778_v35_
                index296_ = default__.safeIndex(71, (d_1846_v77_).length(0))
                (d_1846_v77_)[index296_] = d_1778_v35_
        d_1847_v78_: _dafny.MultiSet
        d_1847_v78_ = _dafny.MultiSet([d_1769_v27_])
        d_1848_v79_: _dafny.Seq
        d_1848_v79_ = _dafny.SeqWithoutIsStrInference([True])
        d_1849_v80_: _dafny.MultiSet
        d_1849_v80_ = _dafny.MultiSet([d_1848_v79_])
        d_1850_v81_: _dafny.Map
        d_1850_v81_ = _dafny.Map({(d_1847_v78_).cardinality: d_1849_v80_})
        if default__.fm29(((d_1850_v81_)[p0] if (p0) in (d_1850_v81_) else d_1849_v80_), p0, globalState):
            d_1851_v82_: _dafny.MultiSet
            d_1851_v82_ = _dafny.MultiSet([self.f36, self.f36])
            d_1852_v83_: _dafny.Map
            d_1852_v83_ = _dafny.Map({self.f35: d_1851_v82_})
            d_1853_v84_: T0
            nw338_ = C3()
            nw338_.ctor__((self).f29, self.f28, self.f28, p0)
            d_1853_v84_ = nw338_
            d_1854_v85_: D8
            d_1854_v85_ = D8_DC20(d_1853_v84_)
            d_1855_v86_: _dafny.MultiSet
            d_1855_v86_ = _dafny.MultiSet([d_1853_v84_.f28])
            d_1856_v87_: _dafny.Map
            d_1856_v87_ = _dafny.Map({self.f35: p0})
            d_1857_v88_: D9
            d_1857_v88_ = D9_DC24(self.f36, not(self.f28), ((d_1855_v86_)[self.f28] if (self.f28) in (d_1855_v86_) else ((d_1856_v87_)[d_1853_v84_.f28] if (d_1853_v84_.f28) in (d_1856_v87_) else p1)), d_1854_v85_, (self).f29)
            d_1858_v89_: _dafny.Map
            d_1858_v89_ = _dafny.Map({p1: d_1857_v88_})
            d_1859_v90_: _dafny.Map
            d_1859_v90_ = _dafny.Map({(d_1853_v84_).f29: d_1858_v89_})
            d_1860_v91_: _dafny.Array
            nw339_ = _dafny.Array(None, 23)
            nw339_[int(0)] = _dafny.Map({188: D9_DC24(self.f36, True, len(d_1852_v83_), d_1854_v85_, not((d_1853_v84_).f29))})
            nw339_[int(1)] = d_1858_v89_
            nw339_[int(2)] = d_1858_v89_
            nw339_[int(3)] = (d_1858_v89_) | (d_1858_v89_)
            nw339_[int(4)] = d_1858_v89_
            nw339_[int(5)] = d_1858_v89_
            nw339_[int(6)] = d_1858_v89_
            nw339_[int(7)] = (d_1858_v89_) | (d_1858_v89_)
            nw339_[int(8)] = d_1858_v89_
            nw339_[int(9)] = d_1858_v89_
            nw339_[int(10)] = d_1858_v89_
            nw339_[int(11)] = d_1858_v89_
            nw339_[int(12)] = d_1858_v89_
            nw339_[int(13)] = d_1858_v89_
            nw339_[int(14)] = (d_1858_v89_) | (_dafny.Map({p1: D9_DC24(self.f36, self.f28, p0, d_1854_v85_, False)}))
            nw339_[int(15)] = (d_1858_v89_) | (d_1858_v89_)
            nw339_[int(16)] = (d_1858_v89_) | (d_1858_v89_)
            nw339_[int(17)] = (d_1858_v89_ if d_1853_v84_.f28 else d_1858_v89_)
            nw339_[int(18)] = ((d_1859_v90_)[False] if (False) in (d_1859_v90_) else d_1858_v89_)
            nw339_[int(19)] = d_1858_v89_
            nw339_[int(20)] = (d_1858_v89_) | (_dafny.Map({default__.fm25((0) - (p0), p1, globalState): d_1857_v88_}))
            nw339_[int(21)] = d_1858_v89_
            nw339_[int(22)] = d_1858_v89_
            d_1860_v91_ = nw339_
            d_1860_v91_ = d_1860_v91_
            d_1861_v92_: _dafny.Array
            d_1861_v92_ = d_1736_v0_
            d_1862_v93_: _dafny.Map
            d_1862_v93_ = _dafny.Map({(d_1861_v92_): len(d_1770_v29_)})
            d_1862_v93_ = (d_1862_v93_) | (d_1862_v93_)
            (globalState).f24 = p1
            d_1863_v94_: _dafny.Array
            def lambda71_(d_1864_i14_):
                return ((self).f29) and ((self).f29)

            init42_ = lambda71_
            nw340_ = _dafny.Array(None, 27)
            for i0_42_ in range(nw340_.length(0)):
                nw340_[i0_42_] = init42_(i0_42_)
            d_1863_v94_ = nw340_
            d_1863_v94_ = d_1863_v94_
            d_1862_v93_ = (d_1862_v93_).set(d_1736_v0_, p0)
        elif True:
            d_1865_v95_: D8
            d_1865_v95_ = D8_DC21(self.f35)
            source31_ = d_1865_v95_
            if source31_.is_DC21:
                d_1866___mcc_h8_ = source31_.cf27
                d_1867_cf27_ = d_1866___mcc_h8_
                d_1868_v96_: _dafny.Set
                d_1868_v96_ = _dafny.Set({self.f28, self.f35})
                d_1869_v97_: _dafny.Seq
                d_1869_v97_ = _dafny.SeqWithoutIsStrInference([p1, p0, 488, len(d_1868_v96_)])
                default__.m0(d_1770_v29_, (p1) + (p0), (len(d_1869_v97_)) > (p1), globalState)
                (globalState).f23 = _dafny.CodePoint('n')
                (globalState).f21 = (self).f29
                d_1870_v98_: _dafny.Array
                def lambda72_(d_1871_i15_):
                    return self.f35

                init43_ = lambda72_
                nw341_ = _dafny.Array(None, 22)
                for i0_43_ in range(nw341_.length(0)):
                    nw341_[i0_43_] = init43_(i0_43_)
                d_1870_v98_ = nw341_
                index297_ = default__.safeIndex(299, (d_1870_v98_).length(0))
                (d_1870_v98_)[index297_] = (self).f29
                index298_ = default__.safeIndex(299, (d_1870_v98_).length(0))
                (d_1870_v98_)[index298_] = self.f28
            elif source31_.is_DC22:
                d_1872_v99_: _dafny.Seq
                d_1872_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjybm"))
                d_1873_v100_: _dafny.Array
                nw342_ = _dafny.Array(False, 22)
                d_1873_v100_ = nw342_
                d_1874_v101_: D5
                d_1874_v101_ = D5_DC12(d_1873_v100_)
                rhs321_ = p1
                rhs322_ = (default__.safeModuloInt(len(d_1872_v99_), len(_dafny.SeqWithoutIsStrInference([d_1874_v101_, D5_DC12(d_1873_v100_)])))) + (len((_dafny.SeqWithoutIsStrInference([(self).f29, (self).f29, False, (self).f29])) + (_dafny.SeqWithoutIsStrInference([self.f28]))))
                rhs323_ = not ((self.f35 if self.f35 else self.f28)) or (False)
                rhs324_ = (d_1771_v30_) + (d_1771_v30_)
                rhs325_ = (self).f29
                lhs270_ = globalState
                lhs271_ = globalState
                lhs272_ = self
                lhs273_ = globalState
                lhs270_.f1 = rhs321_
                lhs271_.f1 = rhs322_
                lhs272_.f35 = rhs323_
                d_1771_v30_ = rhs324_
                lhs273_.f13 = rhs325_
                (globalState).f27 = (True) == ((self).f29)
                d_1875_v102_: _dafny.Array
                def lambda73_(d_1876_p1_, d_1877_p0_, d_1878_v99_):
                    def lambda74_(d_1879_i16_):
                        return D1_DC2(d_1876_p1_, (0) - (d_1877_p0_), d_1878_v99_, not((self).f29))

                    return lambda74_

                init44_ = lambda73_(p1, p0, d_1872_v99_)
                nw343_ = _dafny.Array(None, 27)
                for i0_44_ in range(nw343_.length(0)):
                    nw343_[i0_44_] = init44_(i0_44_)
                d_1875_v102_ = nw343_
                d_1880_v103_: _dafny.Array
                def lambda75_(d_1881_v79_):
                    def lambda76_(d_1882_i17_):
                        return (d_1881_v79_)

                    return lambda76_

                init45_ = lambda75_(d_1848_v79_)
                nw344_ = _dafny.Array(None, 22)
                for i0_45_ in range(nw344_.length(0)):
                    nw344_[i0_45_] = init45_(i0_45_)
                d_1880_v103_ = nw344_
                index299_ = default__.safeIndex(197, (d_1880_v103_).length(0))
                (d_1880_v103_)[index299_] = ((d_1848_v79_) + (d_1848_v79_)).set(default__.safeIndex(174, len((d_1848_v79_) + (d_1848_v79_))), False)
                d_1883_v104_: _dafny.Map
                d_1883_v104_ = _dafny.Map({(0) - (-552): d_1848_v79_})
                d_1884_v105_: _dafny.Map
                d_1884_v105_ = _dafny.Map({False: p1})
                index300_ = default__.safeIndex(197, (d_1880_v103_).length(0))
                rhs326_ = d_1875_v102_
                rhs327_ = ((((d_1883_v104_)[((d_1884_v105_)[(self).f29] if ((self).f29) in (d_1884_v105_) else p0)] if (((d_1884_v105_)[(self).f29] if ((self).f29) in (d_1884_v105_) else p0)) in (d_1883_v104_) else d_1848_v79_)).set(default__.safeIndex(p0, len(((d_1883_v104_)[((d_1884_v105_)[(self).f29] if ((self).f29) in (d_1884_v105_) else p0)] if (((d_1884_v105_)[(self).f29] if ((self).f29) in (d_1884_v105_) else p0)) in (d_1883_v104_) else d_1848_v79_))), not(self.f35))) + (d_1848_v79_)
                rhs328_ = not(default__.fm38(globalState))
                rhs329_ = (self).f29
                lhs274_ = d_1880_v103_
                lhs275_ = default__.safeIndex(197, (d_1880_v103_).length(0))
                lhs276_ = self
                lhs277_ = globalState
                d_1875_v102_ = rhs326_
                lhs274_[lhs275_] = rhs327_
                lhs276_.f35 = rhs328_
                lhs277_.f27 = rhs329_
                (globalState).f21 = ((d_1872_v99_) == (d_1872_v99_) if self.f28 else self.f28)
            elif True:
                d_1885___mcc_h9_ = source31_.cf26
                d_1886_cf26_ = d_1885___mcc_h9_
                d_1887_v106_: _dafny.Map
                d_1887_v106_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cufuvkasq"))})
                d_1888_v107_: _dafny.Seq
                d_1888_v107_ = _dafny.SeqWithoutIsStrInference([len(d_1887_v106_)])
                d_1889_v108_: _dafny.MultiSet
                d_1889_v108_ = _dafny.MultiSet([d_1886_cf26_.f28])
                d_1890_v109_: _dafny.Array
                nw345_ = _dafny.Array(None, 6)
                nw345_[int(0)] = self.f35
                nw345_[int(1)] = (d_1888_v107_) <= (d_1888_v107_)
                nw345_[int(2)] = ((self).f29 if self.f35 else self.f35)
                nw345_[int(3)] = (d_1889_v108_).ispropersubset(d_1889_v108_)
                nw345_[int(4)] = self.f35
                nw345_[int(5)] = (d_1886_cf26_).f29
                d_1890_v109_ = nw345_
                index301_ = default__.safeIndex(935, (d_1890_v109_).length(0))
                (d_1890_v109_)[index301_] = (self).f29
                d_1891_v110_: _dafny.Seq
                d_1891_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
                d_1892_v111_: _dafny.Seq
                d_1892_v111_ = _dafny.SeqWithoutIsStrInference([d_1891_v110_])
                d_1893_v112_: _dafny.Map
                d_1893_v112_ = _dafny.Map({self.f36: d_1892_v111_})
                index302_ = default__.safeIndex(935, (d_1890_v109_).length(0))
                (d_1890_v109_)[index302_] = (((d_1893_v112_)[self.f36] if (self.f36) in (d_1893_v112_) else d_1892_v111_)) <= (d_1892_v111_)
                d_1894_v113_: D10
                d_1894_v113_ = D10_DC28(_dafny.Map({self.f35: (self).f29}))
                d_1895_v114_: D10
                d_1895_v114_ = D10_DC29(d_1894_v113_)
                d_1895_v114_ = d_1895_v114_
                d_1896_v115_: _dafny.Set
                d_1896_v115_ = _dafny.Set({d_1886_cf26_.f28})
                (globalState).f24 = len(d_1896_v115_)
                (globalState).f24 = p1
            d_1897_v116_: C1
            nw346_ = C1()
            nw346_.ctor__(self.f35, self.f35)
            d_1897_v116_ = nw346_
            if not(((d_1770_v29_)[(self).f29] if ((self).f29) in (d_1770_v29_) else False)):
                d_1898_v117_: _dafny.Array
                def lambda77_(d_1899_i18_):
                    return False

                init46_ = lambda77_
                nw347_ = _dafny.Array(None, 9)
                for i0_46_ in range(nw347_.length(0)):
                    nw347_[i0_46_] = init46_(i0_46_)
                d_1898_v117_ = nw347_
                index303_ = default__.safeIndex(196, (d_1898_v117_).length(0))
                (d_1898_v117_)[index303_] = (self).f29
                index304_ = default__.safeIndex(196, (d_1898_v117_).length(0))
                (d_1898_v117_)[index304_] = self.f35
                d_1898_v117_ = d_1898_v117_
                d_1900_v118_: _dafny.MultiSet
                d_1900_v118_ = _dafny.MultiSet([self.f28, True])
                d_1901_v119_: _dafny.Map
                d_1901_v119_ = _dafny.Map({(d_1900_v118_).cardinality: d_1898_v117_})
                d_1902_v120_: _dafny.Map
                d_1902_v120_ = _dafny.Map({(d_1770_v29_).set((d_1898_v117_)[default__.safeIndex(196, (d_1898_v117_).length(0))], self.f28): d_1898_v117_})
                d_1901_v119_ = (d_1901_v119_).set(p1, ((d_1902_v120_)[d_1770_v29_] if (d_1770_v29_) in (d_1902_v120_) else d_1898_v117_))
                (globalState).f24 = p0
                d_1903_v121_: _dafny.Map
                d_1903_v121_ = _dafny.Map({p1: (False if self.f28 else self.f35)})
                d_1903_v121_ = (d_1903_v121_).set(p1, (self).f29)
            elif True:
                d_1770_v29_ = (_dafny.Map({not(self.f35): (self).f29})) | (d_1770_v29_)
                d_1904_v122_: _dafny.Seq
                d_1904_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clceujd"))
                d_1905_v123_: _dafny.Map
                d_1905_v123_ = _dafny.Map({d_1904_v122_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ronehhkb"))})
                (globalState).f24 = len(((d_1905_v123_)[(d_1904_v122_) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_1906_i19_ in range(default__.abs(576))]))] if ((d_1904_v122_) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_1907_i19_ in range(default__.abs(576))]))) in (d_1905_v123_) else (d_1904_v122_) + (_dafny.SeqWithoutIsStrInference([self.f36 for d_1908_i20_ in range(default__.abs(615))]))))
                d_1909_v124_: _dafny.Map
                d_1909_v124_ = _dafny.Map({False: p1})
                d_1909_v124_ = (d_1909_v124_).set((self.f35) and (not(self.f28)), p0)
                (globalState).f21 = self.f35
                d_1910_v125_: _dafny.Map
                d_1910_v125_ = _dafny.Map({p1: d_1904_v122_})
                d_1911_v126_: _dafny.Map
                d_1911_v126_ = _dafny.Map({(0) - (((_dafny.MultiSet([p1])).cardinality) - (p1)): d_1910_v125_})
                d_1911_v126_ = (d_1911_v126_).set(default__.fm25(p0, 511, globalState), (d_1910_v125_) | (d_1910_v125_))
            index305_ = default__.safeIndex(587, (d_1736_v0_).length(0))
            (d_1736_v0_)[index305_] = default__.safeModuloInt(p0, p1)
            index306_ = default__.safeIndex(587, (d_1736_v0_).length(0))
            (d_1736_v0_)[index306_] = 509
            d_1912_v127_: C2
            nw348_ = C2()
            nw348_.ctor__()
            d_1912_v127_ = nw348_
        d_1913_v128_: D10
        d_1913_v128_ = D10_DC28((d_1770_v29_) | (d_1770_v29_))
        d_1913_v128_ = d_1913_v128_


class C6(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self._f33: _dafny.Array = _dafny.Array(None, 0)
        self._f34: _dafny.Seq = _dafny.Seq({})
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
    def ctor__(self, f33, f34, f28, f29):
        (self)._f33 = f33
        (self)._f34 = f34
        (self).f28 = f28
        (self)._f29 = f29

    def fm21(self, p0, p1, globalState):
        return (_dafny.Set({len(_dafny.Map({(self).f29: -221}))})).isdisjoint(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({552: 256})) for d_1914_i0_ in range(default__.abs(527))])), len(_dafny.Set({self.f28}))}))

    def fm22(self, p0, globalState):
        return (self).f29

    def m5(self, p0, globalState):
        r0: bool = False
        d_1915_v0_: C0
        nw349_ = C0()
        nw349_.ctor__()
        d_1915_v0_ = nw349_
        hi7_ = p0
        for d_1916_i0_ in range(p0, hi7_):
            d_1917_v1_: _dafny.Set
            d_1917_v1_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbjqccpa"))})
            d_1918_v2_: _dafny.Seq
            d_1918_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bf"))
            d_1919_v3_: _dafny.Set
            d_1919_v3_ = _dafny.Set({d_1918_v2_, d_1918_v2_})
            if ((d_1917_v1_) | (d_1917_v1_)).issubset(((d_1919_v3_)) - (d_1917_v1_)):
                d_1920_v4_: D1
                d_1920_v4_ = D1_DC2(len(d_1918_v2_), d_1916_i0_, d_1918_v2_, (self).f29)
                d_1921_v5_: _dafny.MultiSet
                d_1921_v5_ = _dafny.MultiSet([d_1920_v4_])
                d_1922_v6_: D7
                d_1922_v6_ = D7_DC16((d_1921_v5_) - (d_1921_v5_))
                d_1922_v6_ = D7_DC16(_dafny.MultiSet([d_1920_v4_, d_1920_v4_]))
                rhs330_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jc"))) < (d_1918_v2_)
                rhs331_ = d_1916_i0_
                rhs332_ = self.f28
                lhs278_ = globalState
                lhs279_ = globalState
                lhs280_ = globalState
                lhs278_.f13 = rhs330_
                lhs279_.f24 = rhs331_
                lhs280_.f22 = rhs332_
                d_1923_v7_: _dafny.Seq
                d_1923_v7_ = _dafny.SeqWithoutIsStrInference([d_1916_i0_, p0])
                (globalState).f17 = (d_1923_v7_) + (_dafny.SeqWithoutIsStrInference([d_1916_i0_ for d_1924_i1_ in range(default__.abs(772))]))
                d_1925_v8_: D2
                d_1925_v8_ = D2_DC5(self.f28)
                d_1926_v9_: D4
                d_1926_v9_ = D4_DC10(self.f28, p0, p0)
                d_1927_v10_: D4
                d_1927_v10_ = D4_DC11(d_1926_v9_)
                def iife139_(_pat_let29_0):
                    def iife140_(d_1928_dt__update__tmp_h0_):
                        def iife141_(_pat_let30_0):
                            def iife142_(d_1929_dt__update_hcf8_h0_):
                                return D2_DC5(d_1929_dt__update_hcf8_h0_)
                            return iife142_(_pat_let30_0)
                        return iife141_((self).f29)
                    return iife140_(_pat_let29_0)
                rhs333_ = iife139_(d_1925_v8_)
                rhs334_ = d_1927_v10_
                rhs335_ = d_1917_v1_
                d_1925_v8_ = rhs333_
                d_1927_v10_ = rhs334_
                d_1919_v3_ = rhs335_
                d_1930_v11_: _dafny.Array
                nw350_ = _dafny.Array(False, 18)
                d_1930_v11_ = nw350_
                index307_ = default__.safeIndex(454, (d_1930_v11_).length(0))
                (d_1930_v11_)[index307_] = (default__.fm25(d_1916_i0_, p0, globalState)) > (153)
                index308_ = default__.safeIndex(454, (d_1930_v11_).length(0))
                (d_1930_v11_)[index308_] = (self).f29
            elif True:
                d_1931_v12_: _dafny.Seq
                d_1931_v12_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                d_1932_v13_: _dafny.Map
                d_1932_v13_ = _dafny.Map({(self).f29: (d_1931_v12_)[default__.safeIndex(d_1916_i0_, len(d_1931_v12_))]})
                default__.m0(d_1932_v13_, (p0) * (p0), (self).f29, globalState)
                d_1931_v12_ = d_1931_v12_
                d_1933_v14_: D10
                d_1933_v14_ = D10_DC27((self).f29, self.f28, False, 613, self.f28)
                d_1934_v15_: _dafny.Map
                d_1934_v15_ = _dafny.Map({(d_1918_v2_)[default__.safeIndex((d_1933_v14_).cf43, len(d_1918_v2_))]: p0})
                d_1935_v16_: str
                d_1935_v16_ = _dafny.CodePoint('n')
                d_1934_v15_ = (d_1934_v15_).set(d_1935_v16_, d_1916_i0_)
                d_1936_v17_: _dafny.Seq
                d_1936_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))])
                d_1936_v17_ = (_dafny.SeqWithoutIsStrInference([d_1918_v2_ for d_1937_i2_ in range(default__.abs(75))])) + ((self).f34)
                (globalState).f21 = self.f28
            d_1938_v18_: _dafny.Seq
            d_1938_v18_ = _dafny.SeqWithoutIsStrInference([len(d_1918_v2_), d_1916_i0_])
            d_1939_v19_: _dafny.Seq
            d_1939_v19_ = _dafny.SeqWithoutIsStrInference([(self).f29, self.f28])
            (globalState).f1 = default__.safeModuloInt(len(_dafny.Map({len(d_1938_v18_): d_1916_i0_})), len(d_1939_v19_))
            d_1940_v20_: _dafny.Set
            d_1940_v20_ = _dafny.Set({False, (self).f29, default__.fm26(93, globalState)})
            d_1941_v21_: _dafny.Map
            d_1941_v21_ = _dafny.Map({p0: (self).f29})
            d_1942_v22_: _dafny.Array
            nw351_ = _dafny.Array(None, 24)
            nw351_[int(0)] = False
            nw351_[int(1)] = ((self).f29 if self.f28 else False)
            nw351_[int(2)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1943_i3_ in range(default__.abs(954))])) == (d_1918_v2_)
            nw351_[int(3)] = (p0) > (d_1916_i0_)
            nw351_[int(4)] = self.f28
            nw351_[int(5)] = self.f28
            nw351_[int(6)] = (d_1916_i0_) > (p0)
            nw351_[int(7)] = (self).f29
            nw351_[int(8)] = (d_1940_v20_).isdisjoint(d_1940_v20_)
            nw351_[int(9)] = (self).f29
            nw351_[int(10)] = False
            nw351_[int(11)] = ((d_1941_v21_)[d_1916_i0_] if (d_1916_i0_) in (d_1941_v21_) else self.f28)
            nw351_[int(12)] = (self).f29
            nw351_[int(13)] = not(self.f28)
            nw351_[int(14)] = (self).f29
            nw351_[int(15)] = self.f28
            nw351_[int(16)] = (self).f29
            nw351_[int(17)] = (self).f29
            nw351_[int(18)] = self.f28
            nw351_[int(19)] = ((self).f29 if self.f28 else not(True))
            nw351_[int(20)] = (self).f29
            nw351_[int(21)] = not ((self).f29) or ((self).f29)
            nw351_[int(22)] = (self).f29
            nw351_[int(23)] = (self.f28) not in (d_1939_v19_)
            d_1942_v22_ = nw351_
            index309_ = default__.safeIndex(860, (d_1942_v22_).length(0))
            (d_1942_v22_)[index309_] = (self).f29
            index310_ = default__.safeIndex(860, (d_1942_v22_).length(0))
            (d_1942_v22_)[index310_] = not ((self).f29) or ((True) and (self.f28))
            d_1944_v23_: _dafny.Map
            d_1944_v23_ = _dafny.Map({d_1916_i0_: d_1918_v2_})
            d_1945_v24_: _dafny.MultiSet
            d_1945_v24_ = _dafny.MultiSet([p0])
            d_1946_v25_: C5
            nw352_ = C5()
            nw352_.ctor__((d_1915_v0_).fm33(d_1918_v2_, d_1944_v23_, d_1945_v24_, globalState), (d_1918_v2_)[default__.safeIndex(524, len(d_1918_v2_))], self.f28, (d_1942_v22_)[default__.safeIndex(860, (d_1942_v22_).length(0))])
            d_1946_v25_ = nw352_
        d_1947_v26_: _dafny.Array
        nw353_ = _dafny.Array(_dafny.Map({}), 28)
        d_1947_v26_ = nw353_
        d_1948_v27_: _dafny.Map
        d_1948_v27_ = _dafny.Map({p0: self.f28})
        index311_ = default__.safeIndex(458, (d_1947_v26_).length(0))
        (d_1947_v26_)[index311_] = ((d_1948_v27_).set(p0, self.f28)) | (d_1948_v27_)
        index312_ = default__.safeIndex(458, (d_1947_v26_).length(0))
        def iife143_():
            coll81_ = _dafny.Map()
            compr_81_: int
            for compr_81_ in _dafny.IntegerRange(240, 641):
                d_1949_v28_: int = compr_81_
                if ((240) <= (d_1949_v28_)) and ((d_1949_v28_) < (641)):
                    coll81_[default__.safeModuloInt(d_1949_v28_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocdte"))))] = True
            return _dafny.Map(coll81_)
        (d_1947_v26_)[index312_] = (((d_1948_v27_).set(p0, self.f28)) | ((_dafny.Map({p0: False})).set(p0, self.f28)) if self.f28 else iife143_()
        )
        d_1950_v29_: str
        d_1950_v29_ = _dafny.CodePoint('t')
        (globalState).f24 = (default__.safeDivisionInt(default__.fm27(p0, d_1950_v29_, p0, globalState), p0) if (self).f29 else (p0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktfofded")))))
        if ((False) and (self.f28) if self.f28 else (p0) <= (296)):
            d_1951_v30_: _dafny.Array
            def lambda78_(d_1952_i4_):
                return _dafny.MultiSet([(self).f29, not((self).f29), self.f28, (self).f29, self.f28])

            init47_ = lambda78_
            nw354_ = _dafny.Array(None, 23)
            for i0_47_ in range(nw354_.length(0)):
                nw354_[i0_47_] = init47_(i0_47_)
            d_1951_v30_ = nw354_
            d_1953_v31_: _dafny.MultiSet
            d_1953_v31_ = _dafny.MultiSet([d_1951_v30_, d_1951_v30_, d_1951_v30_])
            d_1954_v32_: _dafny.Set
            d_1954_v32_ = _dafny.Set({p0})
            d_1955_v33_: _dafny.Seq
            d_1955_v33_ = _dafny.SeqWithoutIsStrInference([d_1950_v29_, d_1950_v29_])
            rhs336_ = ((d_1953_v31_)[d_1951_v30_] if (d_1951_v30_) in (d_1953_v31_) else (0) - (len(d_1954_v32_)))
            rhs337_ = (d_1955_v33_)[default__.safeIndex(len(d_1954_v32_), len(d_1955_v33_))]
            lhs281_ = globalState
            lhs282_ = globalState
            lhs281_.f24 = rhs336_
            lhs282_.f23 = rhs337_
            d_1956_v34_: _dafny.MultiSet
            d_1956_v34_ = _dafny.MultiSet([self.f28])
            d_1957_v35_: D9
            d_1957_v35_ = D9_DC23(d_1956_v34_)
            source32_ = d_1957_v35_
            if source32_.is_DC24:
                d_1958___mcc_h0_ = source32_.cf29
                d_1959___mcc_h1_ = source32_.cf30
                d_1960___mcc_h2_ = source32_.cf31
                d_1961___mcc_h3_ = source32_.cf32
                d_1962___mcc_h4_ = source32_.cf33
                d_1963_cf33_ = d_1962___mcc_h4_
                d_1964_cf32_ = d_1961___mcc_h3_
                d_1965_cf31_ = d_1960___mcc_h2_
                d_1966_cf30_ = d_1959___mcc_h1_
                d_1967_cf29_ = d_1958___mcc_h0_
                d_1968_v36_: _dafny.Array
                nw355_ = _dafny.Array(_dafny.Map({}), 5)
                d_1968_v36_ = nw355_
                d_1969_v37_: D2
                d_1969_v37_ = D2_DC4(d_1950_v29_)
                rhs338_ = d_1968_v36_
                rhs339_ = d_1969_v37_
                rhs340_ = d_1955_v33_
                d_1968_v36_ = rhs338_
                d_1969_v37_ = rhs339_
                d_1955_v33_ = rhs340_
                d_1970_v38_: _dafny.Seq
                d_1970_v38_ = _dafny.SeqWithoutIsStrInference([d_1966_cf30_])
                d_1971_v39_: _dafny.Array
                nw356_ = _dafny.Array(None, 3)
                nw356_[int(0)] = d_1963_cf33_
                nw356_[int(1)] = not(not((d_1956_v34_).ispropersubset(_dafny.MultiSet(d_1970_v38_))))
                nw356_[int(2)] = (d_1963_cf33_ if (self).f29 else d_1966_cf30_)
                d_1971_v39_ = nw356_
                index313_ = default__.safeIndex(613, (d_1971_v39_).length(0))
                (d_1971_v39_)[index313_] = not(not (False) or (False))
                d_1972_v40_: _dafny.MultiSet
                d_1972_v40_ = _dafny.MultiSet([d_1965_cf31_])
                index314_ = default__.safeIndex(613, (d_1971_v39_).length(0))
                (d_1971_v39_)[index314_] = (d_1915_v0_).fm34((self).fm22(d_1972_v40_, globalState), (d_1956_v34_) != (d_1956_v34_), globalState)
                (globalState).f27 = d_1963_cf33_
                d_1970_v38_ = d_1970_v38_
            elif source32_.is_DC25:
                d_1973___mcc_h5_ = source32_.cf34
                d_1974___mcc_h6_ = source32_.cf35
                d_1975___mcc_h7_ = source32_.cf36
                d_1976___mcc_h8_ = source32_.cf37
                d_1977___mcc_h9_ = source32_.cf38
                d_1978_cf38_ = d_1977___mcc_h9_
                d_1979_cf37_ = d_1976___mcc_h8_
                d_1980_cf36_ = d_1975___mcc_h7_
                d_1981_cf35_ = d_1974___mcc_h6_
                d_1982_cf34_ = d_1973___mcc_h5_
                d_1983_v42_: _dafny.Array
                def lambda79_(d_1984_v33_, d_1985_v29_):
                    def lambda80_(d_1986_i5_):
                        def iife144_():
                            coll82_ = _dafny.Map()
                            compr_82_: _dafny.Seq
                            for compr_82_ in ((self).f34).Elements:
                                d_1987_v41_: _dafny.Seq = compr_82_
                                if (d_1987_v41_) in ((self).f34):
                                    coll82_[d_1987_v41_] = d_1985_v29_
                            return _dafny.Map(coll82_)
                        return (_dafny.Map({d_1984_v33_: d_1985_v29_})) | (iife144_()
                        )

                    return lambda80_

                init48_ = lambda79_(d_1955_v33_, d_1950_v29_)
                nw357_ = _dafny.Array(None, 2)
                for i0_48_ in range(nw357_.length(0)):
                    nw357_[i0_48_] = init48_(i0_48_)
                d_1983_v42_ = nw357_
                d_1988_v43_: _dafny.Map
                d_1988_v43_ = _dafny.Map({d_1955_v33_: d_1950_v29_})
                index315_ = default__.safeIndex(727, (d_1983_v42_).length(0))
                (d_1983_v42_)[index315_] = (d_1988_v43_).set(((self).f34)[default__.safeIndex(len(_dafny.Map({True: len(d_1954_v32_)})), len((self).f34))], d_1950_v29_)
                index316_ = default__.safeIndex(727, (d_1983_v42_).length(0))
                (d_1983_v42_)[index316_] = (d_1988_v43_) | (d_1988_v43_)
                d_1989_v44_: _dafny.Array
                def lambda81_(d_1990_v33_, d_1991_cf35_, d_1992_v29_):
                    def lambda82_(d_1993_i6_):
                        return (_dafny.MultiSet([(d_1990_v33_).set(default__.safeIndex(d_1991_cf35_, len(d_1990_v33_)), d_1992_v29_), d_1990_v33_, d_1990_v33_])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1992_v29_ for d_1994_i7_ in range(default__.abs(345))])]))

                    return lambda82_

                init49_ = lambda81_(d_1955_v33_, d_1981_cf35_, d_1950_v29_)
                nw358_ = _dafny.Array(None, 25)
                for i0_49_ in range(nw358_.length(0)):
                    nw358_[i0_49_] = init49_(i0_49_)
                d_1989_v44_ = nw358_
                d_1995_v45_: _dafny.MultiSet
                d_1995_v45_ = _dafny.MultiSet([d_1955_v33_, d_1955_v33_, ((self).f34)[default__.safeIndex(d_1981_cf35_, len((self).f34))], d_1955_v33_])
                index317_ = default__.safeIndex(642, (d_1989_v44_).length(0))
                (d_1989_v44_)[index317_] = d_1995_v45_
                index318_ = default__.safeIndex(642, (d_1989_v44_).length(0))
                (d_1989_v44_)[index318_] = (d_1995_v45_).intersection(d_1995_v45_)
                d_1996_v47_: _dafny.Seq
                d_1996_v47_ = _dafny.SeqWithoutIsStrInference([len(d_1955_v33_)])
                d_1997_v48_: _dafny.MultiSet
                d_1997_v48_ = _dafny.MultiSet([len(d_1996_v47_), 498])
                d_1998_v49_: _dafny.Seq
                d_1998_v49_ = _dafny.SeqWithoutIsStrInference([d_1997_v48_])
                d_1999_v50_: C4
                nw359_ = C4()
                def iife145_():
                    coll83_ = _dafny.Map()
                    compr_83_: int
                    for compr_83_ in _dafny.IntegerRange(-877, 613):
                        d_2000_v46_: int = compr_83_
                        if ((-877) <= (d_2000_v46_)) and ((d_2000_v46_) < (613)):
                            coll83_[default__.safeDivisionInt(d_2000_v46_, d_1978_cf38_)] = d_1955_v33_
                    return _dafny.Map(coll83_)
                nw359_.ctor__((p0) <= (d_1978_cf38_), d_1982_cf34_, (d_1915_v0_).fm33(d_1955_v33_, iife145_()
                , (d_1998_v49_)[default__.safeIndex(d_1981_cf35_, len(d_1998_v49_))], globalState), d_1982_cf34_)
                d_1999_v50_ = nw359_
                d_2001_v51_: C2
                nw360_ = C2()
                nw360_.ctor__()
                d_2001_v51_ = nw360_
            elif True:
                d_2002___mcc_h10_ = source32_.cf28
                d_2003_cf28_ = d_2002___mcc_h10_
                d_2004_v52_: T0
                nw361_ = C1()
                nw361_.ctor__(self.f28, self.f28)
                d_2004_v52_ = nw361_
                d_2005_v53_: D8
                d_2005_v53_ = D8_DC20(d_2004_v52_)
                d_2006_v54_: D9
                d_2006_v54_ = D9_DC24(d_1950_v29_, ((d_1948_v27_)[p0] if (p0) in (d_1948_v27_) else not(self.f28)), default__.fm27(493, d_1950_v29_, p0, globalState), d_2005_v53_, False)
                d_2007_v55_: _dafny.Map
                d_2007_v55_ = _dafny.Map({(d_2006_v54_).cf29: p0})
                d_2007_v55_ = (d_2007_v55_).set(d_1950_v29_, p0)
                d_2008_v56_: _dafny.Array
                nw362_ = _dafny.Array(False, 11)
                d_2008_v56_ = nw362_
                index319_ = default__.safeIndex(959, (d_2008_v56_).length(0))
                (d_2008_v56_)[index319_] = (d_2004_v52_).f29
                index320_ = default__.safeIndex(959, (d_2008_v56_).length(0))
                (d_2008_v56_)[index320_] = (self).f29
                d_2009_v57_: _dafny.Array
                def lambda83_(d_2010_v33_):
                    def lambda84_(d_2011_i8_):
                        return d_2010_v33_

                    return lambda84_

                init50_ = lambda83_(d_1955_v33_)
                nw363_ = _dafny.Array(None, 20)
                for i0_50_ in range(nw363_.length(0)):
                    nw363_[i0_50_] = init50_(i0_50_)
                d_2009_v57_ = nw363_
                index321_ = default__.safeIndex(396, (d_2009_v57_).length(0))
                (d_2009_v57_)[index321_] = d_1955_v33_
                index322_ = default__.safeIndex(396, (d_2009_v57_).length(0))
                (d_2009_v57_)[index322_] = _dafny.SeqWithoutIsStrInference([d_1950_v29_ for d_2012_i9_ in range(default__.abs(357))])
                d_2013_v58_: _dafny.Map
                d_2013_v58_ = _dafny.Map({d_2004_v52_.f28: self.f28})
                d_2014_v59_: _dafny.Seq
                d_2014_v59_ = _dafny.SeqWithoutIsStrInference([d_2013_v58_])
                d_2015_v60_: _dafny.Map
                d_2015_v60_ = _dafny.Map({(d_2009_v57_)[default__.safeIndex(396, (d_2009_v57_).length(0))]: ((d_2014_v59_)[default__.safeIndex(p0, len(d_2014_v59_))]) == ((d_2013_v58_).set(False, (d_2008_v56_)[default__.safeIndex(959, (d_2008_v56_).length(0))]))})
                d_2015_v60_ = (D17_DC43(d_2015_v60_)).cf61
            d_2016_v61_: D2
            d_2016_v61_ = D2_DC4(d_1950_v29_)
            d_2017_v62_: D5
            d_2017_v62_ = D5_DC13(d_2016_v61_)
            d_2017_v62_ = default__.fm44((self.f28 if self.f28 else self.f28), globalState)
            (globalState).f1 = p0
            d_2018_v64_: _dafny.Seq
            d_2018_v64_ = _dafny.SeqWithoutIsStrInference([len(default__.fm45(globalState))])
            d_2019_v65_: _dafny.MultiSet
            def iife146_():
                coll84_ = _dafny.Map()
                compr_84_: int
                for compr_84_ in (d_2018_v64_).Elements:
                    d_2020_v63_: int = compr_84_
                    if (d_2020_v63_) in (d_2018_v64_):
                        coll84_[(d_2020_v63_) + ((0) - (p0))] = (self).f29
                return _dafny.Map(coll84_)
            d_2019_v65_ = _dafny.MultiSet([(d_1947_v26_)[default__.safeIndex(458, (d_1947_v26_).length(0))], iife146_()
            ])
            d_2021_v66_: _dafny.Seq
            d_2021_v66_ = _dafny.SeqWithoutIsStrInference([d_2019_v65_, d_2019_v65_, (d_2019_v65_) - (d_2019_v65_)])
            d_2019_v65_ = (d_2021_v66_)[default__.safeIndex(len(d_1955_v33_), len(d_2021_v66_))]
        elif True:
            d_2022_v67_: _dafny.Array
            def lambda85_(d_2023_i10_):
                return (self).f29

            init51_ = lambda85_
            nw364_ = _dafny.Array(None, 13)
            for i0_51_ in range(nw364_.length(0)):
                nw364_[i0_51_] = init51_(i0_51_)
            d_2022_v67_ = nw364_
            nw365_ = _dafny.Array(False, 18)
            d_2022_v67_ = nw365_
            (globalState).f24 = p0
            d_2024_v68_: _dafny.Set
            d_2025_v69_: int
            d_2026_v70_: bool
            out37_: _dafny.Set
            out38_: int
            out39_: bool
            out37_, out38_, out39_ = (self).m11((0) - (default__.fm27(p0, d_1950_v29_, -812, globalState)), (self).f29, p0, p0, globalState)
            d_2024_v68_ = out37_
            d_2025_v69_ = out38_
            d_2026_v70_ = out39_
            d_2027_v71_: D12
            d_2027_v71_ = D12_DC33(D12_DC32())
            rhs341_ = d_2027_v71_
            rhs342_ = default__.safeDivisionInt(p0, (0) - (d_2025_v69_))
            rhs343_ = d_1950_v29_
            lhs283_ = globalState
            lhs284_ = globalState
            d_2027_v71_ = rhs341_
            lhs283_.f24 = rhs342_
            lhs284_.f23 = rhs343_
            if not(d_2026_v70_):
                (globalState).f22 = d_2026_v70_
                index323_ = default__.safeIndex(480, (d_2022_v67_).length(0))
                (d_2022_v67_)[index323_] = (self.f28 if self.f28 else (self).f29)
                d_2028_v72_: _dafny.Seq
                d_2028_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                index324_ = default__.safeIndex(480, (d_2022_v67_).length(0))
                rhs344_ = d_1950_v29_
                rhs345_ = ((p0) == ((_dafny.MultiSet([p0])).cardinality)) == (self.f28)
                rhs346_ = p0
                rhs347_ = ((d_2028_v72_) + (d_2028_v72_) if (self).f29 else d_2028_v72_)
                rhs348_ = default__.safeDivisionInt(d_2025_v69_, (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "empckknky"))) if (self).f29 else d_2025_v69_)))
                lhs285_ = globalState
                lhs286_ = d_2022_v67_
                lhs287_ = default__.safeIndex(480, (d_2022_v67_).length(0))
                lhs288_ = globalState
                lhs285_.f23 = rhs344_
                lhs286_[lhs287_] = rhs345_
                lhs288_.f1 = rhs346_
                d_2028_v72_ = rhs347_
                d_2025_v69_ = rhs348_
                d_2028_v72_ = d_2028_v72_
                d_2029_v73_: _dafny.Map
                d_2029_v73_ = _dafny.Map({d_2025_v69_: d_2028_v72_})
                d_2030_v74_: _dafny.MultiSet
                d_2030_v74_ = _dafny.MultiSet([p0])
                d_2031_v75_: _dafny.Seq
                d_2031_v75_ = _dafny.SeqWithoutIsStrInference([d_2026_v70_, (d_2022_v67_)[default__.safeIndex(480, (d_2022_v67_).length(0))], self.f28, (d_1915_v0_).fm33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpfjt")), d_2029_v73_, ((d_2030_v74_).set(len(d_2028_v72_), default__.abs(d_2025_v69_))).set(p0, default__.abs(p0)), globalState)])
                d_2032_v76_: _dafny.Map
                d_2032_v76_ = _dafny.Map({d_2028_v72_: (self).f29})
                rhs349_ = (58) - (len(_dafny.Map({len(d_2032_v76_): self.f28})))
                rhs350_ = (default__.safeDivisionInt(p0, (0) - (d_2025_v69_)) if not((self).f29) else d_2025_v69_)
                rhs351_ = d_2026_v70_
                rhs352_ = _dafny.SeqWithoutIsStrInference([(self).fm22(d_2030_v74_, globalState), (self).f29, (((d_1947_v26_)[default__.safeIndex(458, (d_1947_v26_).length(0))])[d_2025_v69_] if (d_2025_v69_) in ((d_1947_v26_)[default__.safeIndex(458, (d_1947_v26_).length(0))]) else True), (d_2022_v67_)[default__.safeIndex(480, (d_2022_v67_).length(0))], (False if d_2026_v70_ else not(self.f28))])
                lhs289_ = globalState
                lhs290_ = globalState
                lhs291_ = globalState
                lhs289_.f24 = rhs349_
                lhs290_.f24 = rhs350_
                lhs291_.f27 = rhs351_
                d_2031_v75_ = rhs352_
                d_2033_v77_: _dafny.Array
                nw366_ = _dafny.Array(None, 4)
                nw366_[int(0)] = d_1950_v29_
                nw366_[int(1)] = d_1950_v29_
                nw366_[int(2)] = _dafny.CodePoint('o')
                nw366_[int(3)] = d_1950_v29_
                d_2033_v77_ = nw366_
                d_2034_v78_: T0
                nw367_ = C4()
                nw367_.ctor__(self.f28, (d_2022_v67_)[default__.safeIndex(480, (d_2022_v67_).length(0))], d_2026_v70_, not(d_2026_v70_))
                d_2034_v78_ = nw367_
                d_2035_v79_: _dafny.Map
                d_2035_v79_ = _dafny.Map({d_2033_v77_: D9_DC24(_dafny.CodePoint('k'), d_2026_v70_, d_2025_v69_, D8_DC20(d_2034_v78_), d_2034_v78_.f28)})
                d_2036_v80_: D8
                d_2036_v80_ = D8_DC20(d_2034_v78_)
                d_2037_v81_: D9
                d_2037_v81_ = D9_DC24(d_1950_v29_, self.f28, (0) - (p0), d_2036_v80_, not((d_2022_v67_)[default__.safeIndex(480, (d_2022_v67_).length(0))]))
                d_2035_v79_ = (d_2035_v79_).set(d_2033_v77_, d_2037_v81_)
            elif True:
                index325_ = default__.safeIndex(70, (d_2022_v67_).length(0))
                (d_2022_v67_)[index325_] = (d_2025_v69_) < (d_2025_v69_)
                d_2038_v82_: _dafny.Set
                d_2038_v82_ = _dafny.Set({d_2022_v67_})
                index326_ = default__.safeIndex(70, (d_2022_v67_).length(0))
                (d_2022_v67_)[index326_] = (d_2038_v82_).issubset(_dafny.Set({d_2022_v67_}))
                d_2039_v83_: _dafny.Seq
                d_2039_v83_ = _dafny.SeqWithoutIsStrInference([(d_2022_v67_)[default__.safeIndex(70, (d_2022_v67_).length(0))]])
                d_2040_v84_: C1
                nw368_ = C1()
                nw368_.ctor__((d_2022_v67_)[default__.safeIndex(70, (d_2022_v67_).length(0))], ((d_2039_v83_)[default__.safeIndex(d_2025_v69_, len(d_2039_v83_))]) or (not((self).f29)))
                d_2040_v84_ = nw368_
                (globalState).f27 = self.f28
                d_2041_v85_: _dafny.Seq
                d_2041_v85_ = _dafny.SeqWithoutIsStrInference([(d_2039_v83_) + (d_2039_v83_), (d_2039_v83_) + (d_2039_v83_), default__.fm62(globalState)])
                d_2041_v85_ = default__.fm46(True, globalState)
                d_2042_v86_: _dafny.Map
                d_2042_v86_ = _dafny.Map({d_2026_v70_: d_2025_v69_})
                d_2043_v87_: _dafny.Seq
                d_2043_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pokql"))
                d_2044_v88_: _dafny.Map
                d_2044_v88_ = _dafny.Map({len(d_2043_v87_): _dafny.Map({(self).f29: d_2025_v69_})})
                d_2042_v86_ = (d_2042_v86_) | (((d_2044_v88_)[d_2025_v69_] if (d_2025_v69_) in (d_2044_v88_) else d_2042_v86_))
        d_2045_v89_: _dafny.Array
        def lambda86_(d_2046_p0_):
            def lambda87_(d_2047_i11_):
                return (d_2047_i11_) - (d_2046_p0_)

            return lambda87_

        init52_ = lambda86_(p0)
        nw369_ = _dafny.Array(None, 2)
        for i0_52_ in range(nw369_.length(0)):
            nw369_[i0_52_] = init52_(i0_52_)
        d_2045_v89_ = nw369_
        nw370_ = _dafny.Array(int(0), 14)
        d_2045_v89_ = nw370_
        d_2048_v90_: _dafny.MultiSet
        d_2048_v90_ = _dafny.MultiSet([self.f28])
        d_2049_v91_: _dafny.Map
        d_2049_v91_ = _dafny.Map({(self).f29: (d_2048_v90_).cardinality})
        d_2050_v92_: _dafny.Seq
        d_2050_v92_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f29: (0) - (p0)}), d_2049_v91_])
        d_2051_v93_: _dafny.MultiSet
        d_2051_v93_ = _dafny.MultiSet([d_2049_v91_, d_2049_v91_])
        r0 = (_dafny.MultiSet(d_2050_v92_)).isdisjoint(d_2051_v93_)
        return r0

    def m11(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: bool = False
        d_2052_i0_: int
        d_2052_i0_ = 0
        with _dafny.label("5"):
            while True:
                with _dafny.c_label("5"):
                    if (d_2052_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_2052_i0_ = (d_2052_i0_) + (1)
                    (self).f28 = not(p1)
                    (globalState).f13 = (self).f29
                    d_2053_v0_: str
                    d_2053_v0_ = _dafny.CodePoint('p')
                    (globalState).f23 = d_2053_v0_
                    d_2054_v1_: D17
                    d_2054_v1_ = D17_DC44(self.f28, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugbkwxfv")))
                    d_2055_v2_: _dafny.Seq
                    d_2055_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ies"))
                    def iife147_(_pat_let31_0):
                        def iife148_(d_2056_dt__update__tmp_h0_):
                            def iife149_(_pat_let32_0):
                                def iife150_(d_2057_dt__update_hcf62_h0_):
                                    return D17_DC44(d_2057_dt__update_hcf62_h0_, (d_2056_dt__update__tmp_h0_).cf63)
                                return iife150_(_pat_let32_0)
                            return iife149_(False)
                        return iife148_(_pat_let31_0)
                    d_2054_v1_ = iife147_(D17_DC44(p1, d_2055_v2_))
                    pass
            pass
        d_2058_v3_: _dafny.Array
        def lambda88_(d_2059_p1_, d_2060_p0_, d_2061_p3_):
            def lambda89_(d_2062_i2_):
                return (_dafny.Map({d_2059_p1_: d_2060_p0_})) | (_dafny.Map({(self).f29: d_2061_p3_}))

            return lambda89_

        init53_ = lambda88_(p1, p0, p3)
        nw371_ = _dafny.Array(None, 12)
        for i0_53_ in range(nw371_.length(0)):
            nw371_[i0_53_] = init53_(i0_53_)
        d_2058_v3_ = nw371_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2058_v3_).length(0)):
            d_2063_i1_: int = guard_loop_6_
            if (True) and (((0) <= (d_2063_i1_)) and ((d_2063_i1_) < ((d_2058_v3_).length(0)))):
                (d_2058_v3_)[(d_2063_i1_)] = _dafny.Map({p1: (0) - (p2)})
        d_2064_v4_: _dafny.Array
        nw372_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_2064_v4_ = nw372_
        d_2065_v5_: _dafny.Array
        nw373_ = _dafny.Array(int(0), 4)
        d_2065_v5_ = nw373_
        index327_ = default__.safeIndex(223, (d_2064_v4_).length(0))
        (d_2064_v4_)[index327_] = d_2065_v5_
        index328_ = default__.safeIndex(223, (d_2064_v4_).length(0))
        (d_2064_v4_)[index328_] = d_2065_v5_
        index329_ = default__.safeIndex(690, (d_2065_v5_).length(0))
        (d_2065_v5_)[index329_] = p0
        d_2066_v6_: _dafny.Seq
        d_2066_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
        pat_let_tv59_ = p2
        pat_let_tv60_ = p3
        index330_ = default__.safeIndex(690, (d_2065_v5_).length(0))
        def lambda90_(source33_):
            if source33_.is_DC24:
                d_2067___mcc_h0_ = source33_.cf29
                d_2068___mcc_h1_ = source33_.cf30
                d_2069___mcc_h2_ = source33_.cf31
                d_2070___mcc_h3_ = source33_.cf32
                d_2071___mcc_h4_ = source33_.cf33
                d_2072_cf33_ = d_2071___mcc_h4_
                d_2073_cf32_ = d_2070___mcc_h3_
                d_2074_cf31_ = d_2069___mcc_h2_
                d_2075_cf30_ = d_2068___mcc_h1_
                d_2076_cf29_ = d_2067___mcc_h0_
                return pat_let_tv59_
            elif source33_.is_DC25:
                d_2077___mcc_h5_ = source33_.cf34
                d_2078___mcc_h6_ = source33_.cf35
                d_2079___mcc_h7_ = source33_.cf36
                d_2080___mcc_h8_ = source33_.cf37
                d_2081___mcc_h9_ = source33_.cf38
                d_2082_cf38_ = d_2081___mcc_h9_
                d_2083_cf37_ = d_2080___mcc_h8_
                d_2084_cf36_ = d_2079___mcc_h7_
                d_2085_cf35_ = d_2078___mcc_h6_
                d_2086_cf34_ = d_2077___mcc_h5_
                return default__.safeModuloInt(d_2084_cf36_, len(_dafny.Map({pat_let_tv60_: not(d_2083_cf37_)})))
            elif True:
                d_2087___mcc_h10_ = source33_.cf28
                d_2088_cf28_ = d_2087___mcc_h10_
                return 901

        (d_2065_v5_)[index330_] = lambda90_(default__.fm68(d_2066_v6_, globalState))
        index331_ = default__.safeIndex(690, (d_2065_v5_).length(0))
        (d_2065_v5_)[index331_] = p0
        d_2089_v7_: C3
        nw374_ = C3()
        nw374_.ctor__(self.f28, default__.fm26(p2, globalState), ((self).f29) or (not(p1)), (d_2065_v5_)[default__.safeIndex(690, (d_2065_v5_).length(0))])
        d_2089_v7_ = nw374_
        d_2090_v8_: str
        d_2090_v8_ = _dafny.CodePoint('g')
        d_2091_v9_: _dafny.Set
        d_2091_v9_ = _dafny.Set({p2, p2})
        d_2092_v10_: _dafny.Set
        d_2092_v10_ = _dafny.Set({((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgg"))).set(default__.safeIndex(len(_dafny.Map({p1: self.f28})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgg")))), d_2090_v8_)).set(default__.safeIndex(p2, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgg"))).set(default__.safeIndex(len(_dafny.Map({p1: self.f28})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgg")))), d_2090_v8_))), default__.fm42(-874, default__.fm27(p2, d_2090_v8_, len(d_2091_v9_), globalState), globalState)), d_2066_v6_})
        r0 = (d_2092_v10_).intersection(_dafny.Set({(default__.fm51(p3, (d_2065_v5_)[default__.safeIndex(690, (d_2065_v5_).length(0))], globalState)).set(default__.safeIndex(p3, len(default__.fm51(p3, (d_2065_v5_)[default__.safeIndex(690, (d_2065_v5_).length(0))], globalState))), d_2090_v8_)}))
        d_2093_v11_: _dafny.Seq
        d_2093_v11_ = _dafny.SeqWithoutIsStrInference([p2, p0])
        d_2094_v12_: _dafny.Seq
        d_2094_v12_ = _dafny.SeqWithoutIsStrInference([False])
        r1 = (len((d_2093_v11_) + (d_2093_v11_))) + ((len(d_2094_v12_)) - (p0))
        r2 = not((_dafny.Set({p0})).ispropersubset(d_2091_v9_))
        return r0, r1, r2

    @property
    def f33(self):
        return self._f33
    @property
    def f34(self):
        return self._f34

class C7(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
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
    def ctor__(self, f28, f29):
        (self).f28 = f28
        (self)._f29 = f29

    def fm15(self, globalState):
        def iife151_():
            coll85_ = _dafny.Set()
            compr_85_: int
            for compr_85_ in _dafny.IntegerRange(552, 721):
                d_2095_v0_: int = compr_85_
                if ((552) <= (d_2095_v0_)) and ((d_2095_v0_) < (721)):
                    coll85_ = coll85_.union(_dafny.Set([(d_2095_v0_) * ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nafxl")))])).cardinality)]))
            return _dafny.Set(coll85_)
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f29])): iife151_()
        })

    def fm16(self, p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irhbac"))), -663)

    def m5(self, p0, globalState):
        r0: bool = False
        (self).f28 = (self).f29
        (globalState).f24 = p0
        (globalState).f27 = self.f28
        hi8_ = p0
        for d_2096_i0_ in range((0) - (p0), hi8_):
            d_2097_v0_: _dafny.Array
            nw375_ = _dafny.Array(int(0), 13)
            d_2097_v0_ = nw375_
            d_2098_v1_: str
            d_2098_v1_ = _dafny.CodePoint('j')
            rhs353_ = d_2098_v1_
            rhs354_ = d_2097_v0_
            rhs355_ = 525
            rhs356_ = d_2097_v0_
            lhs292_ = globalState
            lhs293_ = globalState
            lhs292_.f23 = rhs353_
            d_2097_v0_ = rhs354_
            lhs293_.f24 = rhs355_
            d_2097_v0_ = rhs356_
            d_2097_v0_ = d_2097_v0_
            d_2099_v2_: D2
            d_2099_v2_ = D2_DC4(d_2098_v1_)
            d_2100_v3_: D2
            d_2100_v3_ = D2_DC6(d_2099_v2_)
            d_2101_v4_: D2
            d_2101_v4_ = D2_DC6(d_2099_v2_)
            source34_ = d_2101_v4_
            if source34_.is_DC5:
                d_2102___mcc_h0_ = source34_.cf8
                d_2103_cf8_ = d_2102___mcc_h0_
                d_2104_v5_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2104_v5_ = nw376_
                index332_ = default__.safeIndex(725, (d_2104_v5_).length(0))
                (d_2104_v5_)[index332_] = d_2097_v0_
                index333_ = default__.safeIndex(725, (d_2104_v5_).length(0))
                (d_2104_v5_)[index333_] = d_2097_v0_
                d_2105_v6_: _dafny.Array
                def lambda91_(d_2106_i0_):
                    def lambda92_(d_2107_i1_):
                        return (_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2106_i0_])): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft"))})) | (_dafny.Map({(D4_DC9(_dafny.MultiSet([d_2106_i0_]))).cf12: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "haaakjkkx"))}))

                    return lambda92_

                init54_ = lambda91_(d_2096_i0_)
                nw377_ = _dafny.Array(None, 16)
                for i0_54_ in range(nw377_.length(0)):
                    nw377_[i0_54_] = init54_(i0_54_)
                d_2105_v6_ = nw377_
                d_2108_v7_: _dafny.MultiSet
                d_2108_v7_ = _dafny.MultiSet([d_2096_i0_])
                d_2109_v8_: _dafny.Seq
                d_2109_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssc"))
                d_2110_v9_: _dafny.Map
                d_2110_v9_ = _dafny.Map({d_2108_v7_: d_2109_v8_})
                index334_ = default__.safeIndex(742, (d_2105_v6_).length(0))
                (d_2105_v6_)[index334_] = d_2110_v9_
                d_2111_v10_: _dafny.Seq
                d_2111_v10_ = _dafny.SeqWithoutIsStrInference([d_2096_i0_])
                index335_ = default__.safeIndex(742, (d_2105_v6_).length(0))
                (d_2105_v6_)[index335_] = ((d_2110_v9_) | (d_2110_v9_)) | ((default__.fm17((self).fm16(p0, len(d_2111_v10_), _dafny.SeqWithoutIsStrInference([d_2098_v1_ for d_2112_i2_ in range(default__.abs(991))]), d_2096_i0_, globalState), d_2103_cf8_, not((self).f29), globalState)) | (d_2110_v9_))
                (globalState).f24 = (p0) - (p0)
                (globalState).f24 = 94
            elif source34_.is_DC4:
                d_2113___mcc_h1_ = source34_.cf7
                d_2114_cf7_ = d_2113___mcc_h1_
                d_2115_v11_: _dafny.Seq
                d_2115_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxyxbn"))
                d_2116_v12_: _dafny.MultiSet
                d_2116_v12_ = _dafny.MultiSet([d_2115_v11_])
                d_2117_v13_: _dafny.Map
                d_2117_v13_ = _dafny.Map({d_2096_i0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wogimllxw"))})
                d_2118_v14_: _dafny.Map
                d_2118_v14_ = _dafny.Map({p0: ((d_2116_v12_)[((d_2117_v13_)[d_2096_i0_] if (d_2096_i0_) in (d_2117_v13_) else d_2115_v11_)] if (((d_2117_v13_)[d_2096_i0_] if (d_2096_i0_) in (d_2117_v13_) else d_2115_v11_)) in (d_2116_v12_) else d_2096_i0_)})
                index336_ = default__.safeIndex(258, (d_2097_v0_).length(0))
                (d_2097_v0_)[index336_] = len(d_2118_v14_)
                index337_ = default__.safeIndex(258, (d_2097_v0_).length(0))
                (d_2097_v0_)[index337_] = (self).fm16(p0, (self).fm16(d_2096_i0_, p0, (d_2115_v11_).set(default__.safeIndex(p0, len(d_2115_v11_)), d_2114_cf7_), 746, globalState), d_2115_v11_, (0) - (p0), globalState)
                d_2119_v15_: _dafny.MultiSet
                d_2119_v15_ = _dafny.MultiSet([p0, len(d_2115_v11_)])
                d_2120_v16_: D1
                d_2120_v16_ = D1_DC2(20, p0, (d_2115_v11_).set(default__.safeIndex(p0, len(d_2115_v11_)), d_2114_cf7_), True)
                d_2121_v17_: _dafny.Set
                d_2121_v17_ = _dafny.Set({(d_2120_v16_).cf5, self.f28})
                d_2122_v18_: _dafny.Map
                d_2122_v18_ = _dafny.Map({d_2119_v15_: len(d_2121_v17_)})
                d_2123_v19_: _dafny.Seq
                d_2123_v19_ = _dafny.SeqWithoutIsStrInference([d_2096_i0_, (d_2097_v0_)[default__.safeIndex(258, (d_2097_v0_).length(0))]])
                rhs357_ = default__.fm18(len(_dafny.Map({(d_2097_v0_)[default__.safeIndex(258, (d_2097_v0_).length(0))]: False})), p0, default__.safeModuloInt(((d_2122_v18_)[d_2119_v15_] if (d_2119_v15_) in (d_2122_v18_) else d_2096_i0_), (d_2097_v0_)[default__.safeIndex(258, (d_2097_v0_).length(0))]), globalState)
                rhs358_ = ((0) - (p0)) * ((d_2123_v19_)[default__.safeIndex((d_2097_v0_)[default__.safeIndex(258, (d_2097_v0_).length(0))], len(d_2123_v19_))])
                rhs359_ = self.f28
                rhs360_ = ((self).f29) or (self.f28)
                lhs294_ = self
                lhs295_ = globalState
                lhs296_ = globalState
                lhs294_.f28 = rhs357_
                lhs295_.f24 = rhs358_
                r0 = rhs359_
                lhs296_.f27 = rhs360_
                rhs361_ = -751
                rhs362_ = p0
                lhs297_ = globalState
                lhs298_ = globalState
                lhs297_.f24 = rhs361_
                lhs298_.f1 = rhs362_
                d_2124_v20_: _dafny.Array
                nw378_ = _dafny.Array(False, 17)
                d_2124_v20_ = nw378_
                d_2125_v21_: _dafny.Seq
                d_2125_v21_ = _dafny.SeqWithoutIsStrInference([d_2124_v20_, d_2124_v20_, d_2124_v20_])
                d_2124_v20_ = (d_2125_v21_)[default__.safeIndex(p0, len(d_2125_v21_))]
            elif True:
                d_2126___mcc_h2_ = source34_.cf9
                d_2127_cf9_ = d_2126___mcc_h2_
                d_2128_v22_: _dafny.Array
                def lambda93_(d_2129_i3_):
                    return True

                init55_ = lambda93_
                nw379_ = _dafny.Array(None, 28)
                for i0_55_ in range(nw379_.length(0)):
                    nw379_[i0_55_] = init55_(i0_55_)
                d_2128_v22_ = nw379_
                index338_ = default__.safeIndex(170, (d_2128_v22_).length(0))
                (d_2128_v22_)[index338_] = (self).f29
                index339_ = default__.safeIndex(170, (d_2128_v22_).length(0))
                (d_2128_v22_)[index339_] = not((p0) > (default__.safeModuloInt(d_2096_i0_, p0)))
                (globalState).f23 = d_2098_v1_
                d_2130_v23_: _dafny.MultiSet
                d_2130_v23_ = _dafny.MultiSet([804])
                d_2131_v24_: _dafny.MultiSet
                d_2131_v24_ = _dafny.MultiSet([d_2130_v23_, d_2130_v23_])
                (globalState).f13 = ((_dafny.MultiSet([d_2130_v23_, _dafny.MultiSet([p0, d_2096_i0_])])) | (d_2131_v24_)).ispropersubset((d_2131_v24_) - (d_2131_v24_))
                d_2132_v25_: D2
                d_2132_v25_ = D2_DC4(d_2098_v1_)
                d_2132_v25_ = d_2132_v25_
            (globalState).f1 = d_2096_i0_
        d_2133_v26_: _dafny.Array
        def lambda94_(d_2134_p0_):
            def lambda95_(d_2135_i4_):
                return (d_2135_i4_) - (d_2134_p0_)

            return lambda95_

        init56_ = lambda94_(p0)
        nw380_ = _dafny.Array(None, 15)
        for i0_56_ in range(nw380_.length(0)):
            nw380_[i0_56_] = init56_(i0_56_)
        d_2133_v26_ = nw380_
        d_2136_v27_: _dafny.Map
        d_2136_v27_ = _dafny.Map({self.f28: len(default__.fm19(p0, self.f28, self.f28, (self).f29, globalState))})
        index340_ = default__.safeIndex(869, (d_2133_v26_).length(0))
        (d_2133_v26_)[index340_] = len(d_2136_v27_)
        index341_ = default__.safeIndex(869, (d_2133_v26_).length(0))
        (d_2133_v26_)[index341_] = p0
        if (self).f29:
            (globalState).f13 = (self).f29
            d_2137_v28_: _dafny.Array
            nw381_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_2137_v28_ = nw381_
            d_2138_v29_: D1
            d_2138_v29_ = D1_DC3(D1_DC1(d_2137_v28_))
            d_2139_v30_: D1
            d_2139_v30_ = D1_DC3(d_2138_v29_)
            source35_ = d_2139_v30_
            if source35_.is_DC2:
                d_2140___mcc_h3_ = source35_.cf2
                d_2141___mcc_h4_ = source35_.cf3
                d_2142___mcc_h5_ = source35_.cf4
                d_2143___mcc_h6_ = source35_.cf5
                d_2144_cf5_ = d_2143___mcc_h6_
                d_2145_cf4_ = d_2142___mcc_h5_
                d_2146_cf3_ = d_2141___mcc_h4_
                d_2147_cf2_ = d_2140___mcc_h3_
                d_2148_v31_: _dafny.MultiSet
                d_2148_v31_ = _dafny.MultiSet([d_2145_cf4_])
                (globalState).f13 = ((d_2148_v31_).intersection(d_2148_v31_)) != (default__.fm20(d_2144_cf5_, globalState))
                d_2149_v32_: _dafny.Map
                d_2149_v32_ = _dafny.Map({d_2144_cf5_: self.f28})
                d_2150_v33_: _dafny.Seq
                d_2150_v33_ = _dafny.SeqWithoutIsStrInference([d_2149_v32_, d_2149_v32_, d_2149_v32_])
                d_2151_v34_: _dafny.Seq
                d_2151_v34_ = _dafny.SeqWithoutIsStrInference([(d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]])
                d_2152_v35_: _dafny.Map
                d_2152_v35_ = _dafny.Map({self.f28: d_2151_v34_})
                index342_ = default__.safeIndex(869, (d_2133_v26_).length(0))
                rhs363_ = d_2147_cf2_
                rhs364_ = len(d_2150_v33_)
                rhs365_ = (p0) + (d_2146_cf3_)
                rhs366_ = d_2147_cf2_
                rhs367_ = ((((d_2152_v35_)[d_2144_cf5_] if (d_2144_cf5_) in (d_2152_v35_) else d_2151_v34_)) + (_dafny.SeqWithoutIsStrInference([p0 for d_2153_i5_ in range(default__.abs(860))]))) + (d_2151_v34_)
                lhs299_ = globalState
                lhs300_ = globalState
                lhs301_ = d_2133_v26_
                lhs302_ = default__.safeIndex(869, (d_2133_v26_).length(0))
                lhs303_ = globalState
                lhs304_ = globalState
                lhs299_.f1 = rhs363_
                lhs300_.f1 = rhs364_
                lhs301_[lhs302_] = rhs365_
                lhs303_.f24 = rhs366_
                lhs304_.f17 = rhs367_
                d_2154_v36_: str
                d_2154_v36_ = _dafny.CodePoint('s')
                d_2155_v37_: _dafny.Map
                d_2155_v37_ = _dafny.Map({(-474) - (d_2147_cf2_): d_2154_v36_})
                d_2156_v39_: _dafny.Map
                d_2156_v39_ = _dafny.Map({d_2147_cf2_: len(d_2145_cf4_)})
                def iife152_():
                    coll86_ = _dafny.Map()
                    compr_86_: int
                    for compr_86_ in (d_2156_v39_).keys.Elements:
                        d_2157_v38_: int = compr_86_
                        if (d_2157_v38_) in (d_2156_v39_):
                            coll86_[(d_2157_v38_) * (p0)] = d_2154_v36_
                    return _dafny.Map(coll86_)
                d_2155_v37_ = ((d_2155_v37_) | (iife152_()
                )) | ((d_2155_v37_) | (d_2155_v37_))
                d_2156_v39_ = (d_2156_v39_).set(-565, d_2146_cf3_)
            elif source35_.is_DC1:
                d_2158___mcc_h7_ = source35_.cf1
                d_2159_cf1_ = d_2158___mcc_h7_
                d_2160_v40_: _dafny.Seq
                d_2160_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvfftmt"))
                d_2136_v27_ = (d_2136_v27_).set(not ((self).f29) or (default__.fm18((d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], -319, p0, globalState)), len((d_2160_v40_) + (d_2160_v40_)))
                d_2161_v41_: _dafny.Map
                d_2161_v41_ = _dafny.Map({default__.fm19(546, self.f28, self.f28, (self).f29, globalState): self.f28})
                d_2162_v42_: _dafny.MultiSet
                d_2162_v42_ = _dafny.MultiSet([(d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], (d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], (d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]])
                d_2163_v43_: _dafny.Array
                nw382_ = _dafny.Array(False, 19)
                d_2163_v43_ = nw382_
                d_2164_v44_: _dafny.Array
                nw383_ = _dafny.Array(None, 2)
                nw383_[int(0)] = d_2163_v43_
                nw383_[int(1)] = d_2163_v43_
                d_2164_v44_ = nw383_
                d_2165_v45_: _dafny.Seq
                d_2165_v45_ = _dafny.SeqWithoutIsStrInference([d_2160_v40_])
                d_2166_v46_: T0
                nw384_ = C6()
                nw384_.ctor__(d_2164_v44_, d_2165_v45_, (self).f29, self.f28)
                d_2166_v46_ = nw384_
                d_2167_v47_: _dafny.Set
                d_2167_v47_ = _dafny.Set({d_2166_v46_, d_2166_v46_})
                d_2168_v48_: _dafny.Seq
                d_2168_v48_ = _dafny.SeqWithoutIsStrInference([((d_2162_v42_)[(d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]] if ((d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]) in (d_2162_v42_) else p0), len(d_2167_v47_)])
                d_2161_v41_ = (d_2161_v41_).set(d_2168_v48_, True)
                d_2133_v26_ = d_2133_v26_
                (globalState).f13 = self.f28
            elif True:
                d_2169___mcc_h8_ = source35_.cf6
                d_2170_cf6_ = d_2169___mcc_h8_
                (globalState).f22 = (self.f28 if self.f28 else self.f28)
                d_2171_v49_: _dafny.Array
                nw385_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_2171_v49_ = nw385_
                d_2172_v50_: _dafny.Seq
                d_2172_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cetnqbmts"))
                d_2173_v51_: D3
                d_2173_v51_ = D3_DC8(d_2172_v50_)
                d_2174_v52_: _dafny.Map
                d_2174_v52_ = _dafny.Map({self.f28: (self).f29})
                d_2175_v53_: str
                d_2175_v53_ = _dafny.CodePoint('g')
                d_2176_v54_: _dafny.Seq
                d_2176_v54_ = _dafny.SeqWithoutIsStrInference([(d_2173_v51_).cf11, d_2172_v50_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayic")), (default__.fm51(len(d_2174_v52_), len(d_2172_v50_), globalState)).set(default__.safeIndex((d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], len(default__.fm51(len(d_2174_v52_), len(d_2172_v50_), globalState))), d_2175_v53_)])
                d_2177_v55_: C6
                nw386_ = C6()
                nw386_.ctor__(d_2171_v49_, d_2176_v54_, ((self).f29) or ((self).f29), True)
                d_2177_v55_ = nw386_
                d_2178_v56_: _dafny.Set
                d_2178_v56_ = _dafny.Set({len(_dafny.Map({(self).f29: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvbl")))}))})
                d_2179_v57_: _dafny.Seq
                d_2179_v57_ = _dafny.SeqWithoutIsStrInference([d_2178_v56_, d_2178_v56_])
                d_2180_v58_: _dafny.Map
                d_2180_v58_ = _dafny.Map({(0) - (len(d_2178_v56_)): self.f28})
                (globalState).f13 = (((d_2179_v57_)[default__.safeIndex(286, len(d_2179_v57_))]) - (_dafny.Set({(d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], len(d_2172_v50_), len(d_2180_v58_)}))).isdisjoint(d_2178_v56_)
                d_2181_v59_: _dafny.Map
                d_2181_v59_ = _dafny.Map({(d_2172_v50_) + (d_2172_v50_): 271})
                d_2182_v60_: D2
                d_2182_v60_ = D2_DC4(d_2175_v53_)
                d_2183_v61_: _dafny.Set
                d_2183_v61_ = _dafny.Set({(d_2177_v55_).fm21(d_2182_v60_, len(_dafny.SeqWithoutIsStrInference([(d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))] for d_2184_i6_ in range(default__.abs(359))])), globalState)})
                d_2185_v62_: _dafny.Map
                d_2185_v62_ = _dafny.Map({d_2175_v53_: d_2172_v50_})
                d_2186_v63_: _dafny.Map
                d_2186_v63_ = _dafny.Map({((d_2185_v62_)[d_2175_v53_] if (d_2175_v53_) in (d_2185_v62_) else d_2172_v50_): (d_2172_v50_) + (d_2172_v50_)})
                rhs368_ = False
                rhs369_ = d_2183_v61_
                rhs370_ = ((d_2186_v63_)[_dafny.SeqWithoutIsStrInference([(d_2172_v50_)[default__.safeIndex(p0, len(d_2172_v50_))] for d_2187_i7_ in range(default__.abs(158))])] if (_dafny.SeqWithoutIsStrInference([(d_2172_v50_)[default__.safeIndex(p0, len(d_2172_v50_))] for d_2188_i7_ in range(default__.abs(158))])) in (d_2186_v63_) else d_2172_v50_)
                rhs371_ = d_2181_v59_
                rhs372_ = (self.f28 if (self).f29 else ((self).f29) or (True))
                lhs305_ = globalState
                lhs306_ = globalState
                lhs307_ = globalState
                lhs305_.f13 = rhs368_
                lhs306_.f26 = rhs369_
                d_2172_v50_ = rhs370_
                d_2181_v59_ = rhs371_
                lhs307_.f27 = rhs372_
            index343_ = default__.safeIndex(869, (d_2133_v26_).length(0))
            (d_2133_v26_)[index343_] = (d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]
            d_2189_v64_: str
            d_2189_v64_ = _dafny.CodePoint('r')
            (globalState).f23 = d_2189_v64_
            d_2190_v65_: _dafny.Seq
            d_2190_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlxtec"))
            d_2191_v66_: _dafny.MultiSet
            d_2191_v66_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "buumhtlxo"))), (d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], len(d_2190_v65_)])
            d_2192_v67_: D4
            d_2192_v67_ = D4_DC11(D4_DC9(d_2191_v66_))
            d_2193_v68_: _dafny.Map
            d_2193_v68_ = _dafny.Map({d_2192_v67_: (self).f29})
            d_2194_v69_: _dafny.Array
            def lambda96_(d_2195_i8_):
                return not(self.f28)

            init57_ = lambda96_
            nw387_ = _dafny.Array(None, 11)
            for i0_57_ in range(nw387_.length(0)):
                nw387_[i0_57_] = init57_(i0_57_)
            d_2194_v69_ = nw387_
            d_2196_v70_: _dafny.MultiSet
            d_2196_v70_ = _dafny.MultiSet([d_2194_v69_])
            d_2193_v68_ = (d_2193_v68_).set(d_2192_v67_, (_dafny.MultiSet([d_2194_v69_])).ispropersubset(d_2196_v70_))
        elif True:
            d_2197_v71_: _dafny.Set
            d_2197_v71_ = _dafny.Set({self.f28, self.f28, (self).f29})
            d_2198_v72_: _dafny.Map
            d_2198_v72_ = _dafny.Map({d_2133_v26_: (self).f29})
            d_2199_v73_: _dafny.Map
            d_2199_v73_ = _dafny.Map({(self).f29: d_2198_v72_})
            d_2200_v74_: _dafny.Map
            d_2200_v74_ = _dafny.Map({len(((d_2199_v73_)[(self).f29] if ((self).f29) in (d_2199_v73_) else d_2198_v72_)): self.f28})
            d_2201_v75_: _dafny.Map
            d_2201_v75_ = _dafny.Map({((d_2200_v74_)[p0] if (p0) in (d_2200_v74_) else self.f28): (self).f29})
            d_2202_v76_: _dafny.Set
            d_2202_v76_ = _dafny.Set({(p0) == ((d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]), ((self).f29) not in (d_2197_v71_), ((d_2201_v75_)[not(self.f28)] if (not(self.f28)) in (d_2201_v75_) else True)})
            (globalState).f26 = d_2197_v71_
            d_2203_v77_: _dafny.Seq
            d_2203_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbjjtmai"))
            d_2204_v78_: _dafny.MultiSet
            d_2204_v78_ = _dafny.MultiSet([d_2203_v77_])
            index344_ = default__.safeIndex(869, (d_2133_v26_).length(0))
            (d_2133_v26_)[index344_] = default__.safeDivisionInt((d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))], ((d_2204_v78_)[d_2203_v77_] if (d_2203_v77_) in (d_2204_v78_) else len(d_2203_v77_)))
            d_2205_v79_: C0
            nw388_ = C0()
            nw388_.ctor__()
            d_2205_v79_ = nw388_
            (globalState).f21 = (p0) > (((d_2133_v26_)[default__.safeIndex(869, (d_2133_v26_).length(0))]) - (p0))
            d_2206_v80_: _dafny.Seq
            d_2206_v80_ = _dafny.SeqWithoutIsStrInference([False, (self).f29, not((self).f29)])
            (globalState).f1 = len((_dafny.SeqWithoutIsStrInference([(self).f29]) if True else (_dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])) + (d_2206_v80_)))
        r0 = (not((self).f29)) or ((self).f29)
        return r0


class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def m9(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_2207_v0_: _dafny.Array
        nw389_ = _dafny.Array(False, 25)
        d_2207_v0_ = nw389_
        index345_ = default__.safeIndex(907, (d_2207_v0_).length(0))
        (d_2207_v0_)[index345_] = False
        d_2208_v1_: _dafny.MultiSet
        d_2208_v1_ = _dafny.MultiSet([p1])
        d_2209_v2_: _dafny.Map
        d_2209_v2_ = _dafny.Map({p1: d_2208_v1_})
        d_2210_v3_: bool
        d_2210_v3_ = True
        index346_ = default__.safeIndex(907, (d_2207_v0_).length(0))
        (d_2207_v0_)[index346_] = ((p1) < (len((d_2209_v2_).set(266, d_2208_v1_)))) and (d_2210_v3_)
        d_2211_v4_: _dafny.Array
        def lambda97_(d_2212_p1_):
            def lambda98_(d_2213_i0_):
                return (d_2213_i0_) - (d_2212_p1_)

            return lambda98_

        init58_ = lambda97_(p1)
        nw390_ = _dafny.Array(None, 29)
        for i0_58_ in range(nw390_.length(0)):
            nw390_[i0_58_] = init58_(i0_58_)
        d_2211_v4_ = nw390_
        index347_ = default__.safeIndex(62, (d_2211_v4_).length(0))
        (d_2211_v4_)[index347_] = p1
        d_2214_v5_: _dafny.Seq
        d_2214_v5_ = _dafny.SeqWithoutIsStrInference([(d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))]])
        d_2215_v6_: _dafny.Map
        d_2215_v6_ = _dafny.Map({d_2214_v5_: not(d_2210_v3_)})
        index348_ = default__.safeIndex(907, (d_2207_v0_).length(0))
        index349_ = default__.safeIndex(62, (d_2211_v4_).length(0))
        rhs373_ = (0) - ((p1) * (p1))
        rhs374_ = (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))]
        rhs375_ = not (default__.fm13(globalState)) or (not(True))
        rhs376_ = default__.safeDivisionInt(p1, p1)
        rhs377_ = len(d_2215_v6_)
        lhs308_ = globalState
        lhs309_ = globalState
        lhs310_ = d_2207_v0_
        lhs311_ = default__.safeIndex(907, (d_2207_v0_).length(0))
        lhs312_ = globalState
        lhs313_ = d_2211_v4_
        lhs314_ = default__.safeIndex(62, (d_2211_v4_).length(0))
        lhs308_.f1 = rhs373_
        lhs309_.f13 = rhs374_
        lhs310_[lhs311_] = rhs375_
        lhs312_.f1 = rhs376_
        lhs313_[lhs314_] = rhs377_
        (globalState).f1 = ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) - ((default__.fm14(globalState)).cardinality)
        d_2216_v7_: str
        d_2216_v7_ = _dafny.CodePoint('o')
        d_2217_v8_: C3
        nw391_ = C3()
        nw391_.ctor__((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], len(p0))
        d_2217_v8_ = nw391_
        d_2218_v9_: _dafny.Seq
        d_2218_v9_ = _dafny.SeqWithoutIsStrInference([d_2217_v8_, d_2217_v8_, d_2217_v8_])
        d_2219_v10_: _dafny.Seq
        d_2219_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), d_2216_v7_])
        d_2220_v11_: D9
        d_2220_v11_ = D9_DC25(d_2210_v3_, (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], p1, (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], len(d_2219_v10_))
        d_2221_v12_: C1
        nw392_ = C1()
        nw392_.ctor__((default__.fm27(p1, d_2216_v7_, len(d_2218_v9_), globalState)) < ((d_2220_v11_).cf35), (d_2217_v8_).f40)
        d_2221_v12_ = nw392_
        if False:
            source36_ = default__.fm69(globalState)
            if source36_.is_DC38:
                d_2222___mcc_h0_ = source36_.cf53
                d_2223___mcc_h1_ = source36_.cf54
                d_2224___mcc_h2_ = source36_.cf55
                d_2225_cf55_ = d_2224___mcc_h2_
                d_2226_cf54_ = d_2223___mcc_h1_
                d_2227_cf53_ = d_2222___mcc_h0_
                d_2228_v13_: _dafny.Map
                d_2228_v13_ = _dafny.Map({not((d_2217_v8_).f40): (d_2219_v10_) + (d_2219_v10_)})
                rhs378_ = d_2227_cf53_
                rhs379_ = d_2228_v13_
                lhs315_ = globalState
                lhs315_.f21 = rhs378_
                d_2228_v13_ = rhs379_
                d_2229_v14_: _dafny.Map
                d_2229_v14_ = _dafny.Map({d_2225_cf55_: False})
                d_2229_v14_ = (d_2229_v14_).set(len((d_2219_v10_) + (d_2219_v10_)), (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))])
                d_2230_v15_: _dafny.Map
                d_2230_v15_ = _dafny.Map({d_2210_v3_: d_2226_cf54_})
                d_2230_v15_ = (d_2230_v15_).set(d_2210_v3_, len(d_2219_v10_))
                d_2231_v16_: _dafny.Array
                nw393_ = _dafny.Array(_dafny.Map({}), 15)
                d_2231_v16_ = nw393_
                d_2231_v16_ = d_2231_v16_
            elif source36_.is_DC39:
                d_2232___mcc_h3_ = source36_.cf56
                d_2233_cf56_ = d_2232___mcc_h3_
                d_2234_v18_: _dafny.Set
                d_2234_v18_ = _dafny.Set({not(d_2210_v3_)})
                index350_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                def iife153_():
                    coll87_ = _dafny.Set()
                    compr_87_: int
                    for compr_87_ in _dafny.IntegerRange(44, 64):
                        d_2235_v17_: int = compr_87_
                        if ((44) <= (d_2235_v17_)) and ((d_2235_v17_) < (64)):
                            coll87_ = coll87_.union(_dafny.Set([default__.safeModuloInt(d_2235_v17_, (0) - ((p0)[default__.safeIndex(p1, len(p0))]))]))
                    return _dafny.Set(coll87_)
                (d_2211_v4_)[index350_] = default__.safeDivisionInt(len(iife153_()
                ), default__.safeDivisionInt(763, (0) - (len(d_2234_v18_))))
                d_2233_cf56_ = (d_2233_cf56_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yp")))
                d_2236_v19_: C7
                nw394_ = C7()
                nw394_.ctor__((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], d_2210_v3_)
                d_2236_v19_ = nw394_
                d_2237_v20_: C7
                nw395_ = C7()
                nw395_.ctor__(((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))]) == (d_2210_v3_), (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))])
                d_2237_v20_ = nw395_
            elif source36_.is_DC40:
                d_2238___mcc_h4_ = source36_.cf57
                d_2239___mcc_h5_ = source36_.cf58
                d_2240_cf58_ = d_2239___mcc_h5_
                d_2241_cf57_ = d_2238___mcc_h4_
                d_2242_v21_: _dafny.Map
                d_2242_v21_ = _dafny.Map({(d_2217_v8_).f40: p1})
                d_2243_v22_: _dafny.Map
                d_2243_v22_ = _dafny.Map({d_2241_cf57_: not((d_2217_v8_).f40)})
                d_2242_v21_ = (d_2242_v21_).set((d_2217_v8_).f40, len((d_2243_v22_).set((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], False)))
                d_2244_v23_: _dafny.Map
                d_2244_v23_ = _dafny.Map({((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) * (p1): (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))]})
                d_2244_v23_ = d_2244_v23_
                d_2242_v21_ = (d_2242_v21_).set((d_2217_v8_).f40, p1)
                d_2245_v24_: _dafny.Array
                nw396_ = _dafny.Array(None, 20)
                nw396_[int(0)] = d_2207_v0_
                nw396_[int(1)] = d_2207_v0_
                nw396_[int(2)] = d_2207_v0_
                nw396_[int(3)] = d_2207_v0_
                nw396_[int(4)] = d_2207_v0_
                nw396_[int(5)] = d_2207_v0_
                nw396_[int(6)] = d_2207_v0_
                nw396_[int(7)] = d_2207_v0_
                nw396_[int(8)] = d_2207_v0_
                nw396_[int(9)] = d_2207_v0_
                nw396_[int(10)] = d_2207_v0_
                nw396_[int(11)] = d_2207_v0_
                nw396_[int(12)] = d_2207_v0_
                nw396_[int(13)] = d_2207_v0_
                nw396_[int(14)] = d_2207_v0_
                nw396_[int(15)] = d_2207_v0_
                nw396_[int(16)] = d_2207_v0_
                nw396_[int(17)] = d_2207_v0_
                nw396_[int(18)] = d_2207_v0_
                nw396_[int(19)] = d_2207_v0_
                d_2245_v24_ = nw396_
                index351_ = default__.safeIndex(7, (d_2245_v24_).length(0))
                (d_2245_v24_)[index351_] = d_2207_v0_
                index352_ = default__.safeIndex(7, (d_2245_v24_).length(0))
                (d_2245_v24_)[index352_] = d_2207_v0_
            elif source36_.is_DC37:
                d_2246___mcc_h6_ = source36_.cf52
                d_2247_cf52_ = d_2246___mcc_h6_
                d_2248_v25_: _dafny.Map
                d_2248_v25_ = _dafny.Map({d_2219_v10_: (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]})
                d_2249_v26_: D15
                d_2249_v26_ = D15_DC40((d_2217_v8_).f40, (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))])
                d_2250_v27_: _dafny.Set
                d_2250_v27_ = _dafny.Set({d_2249_v26_})
                (globalState).f13 = (default__.safeModuloInt(p1, len(d_2219_v10_))) > (default__.safeModuloInt(len(d_2248_v25_), len(d_2250_v27_)))
                d_2251_v28_: _dafny.Set
                d_2251_v28_ = _dafny.Set({(d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]})
                (globalState).f1 = (0) - ((len(d_2214_v5_)) - (((d_2208_v1_)[(d_2217_v8_).fm30(d_2251_v28_, globalState)] if ((d_2217_v8_).fm30(d_2251_v28_, globalState)) in (d_2208_v1_) else (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))])))
                (globalState).f24 = (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]
                (globalState).f27 = (((d_2208_v1_).set((0) - ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]), default__.abs((0) - ((d_2208_v1_).cardinality))) if False else d_2208_v1_)).ispropersubset(_dafny.MultiSet([p1]))
            elif True:
                d_2252___mcc_h7_ = source36_.cf59
                d_2253_cf59_ = d_2252___mcc_h7_
                (globalState).f24 = (p0)[default__.safeIndex(p1, len(p0))]
                (globalState).f1 = (p1) * ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))])
                (globalState).f13 = d_2210_v3_
                d_2254_v29_: _dafny.Set
                d_2254_v29_ = _dafny.Set({(d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))]})
                index353_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                rhs380_ = (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]
                rhs381_ = (len(d_2254_v29_)) * (p1)
                lhs316_ = globalState
                lhs317_ = d_2211_v4_
                lhs318_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                lhs316_.f1 = rhs380_
                lhs317_[lhs318_] = rhs381_
            d_2255_v30_: _dafny.Seq
            d_2255_v30_ = _dafny.SeqWithoutIsStrInference([d_2214_v5_])
            d_2256_v31_: _dafny.Map
            d_2256_v31_ = _dafny.Map({p1: d_2214_v5_})
            d_2257_v32_: _dafny.Set
            d_2257_v32_ = _dafny.Set({d_2219_v10_})
            d_2258_v33_: _dafny.Array
            nw397_ = _dafny.Array(None, 9)
            nw397_[int(0)] = d_2214_v5_
            nw397_[int(1)] = d_2214_v5_
            nw397_[int(2)] = d_2214_v5_
            nw397_[int(3)] = d_2214_v5_
            nw397_[int(4)] = (d_2255_v30_)[default__.safeIndex((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], len(d_2255_v30_))]
            nw397_[int(5)] = default__.fm62(globalState)
            nw397_[int(6)] = d_2214_v5_
            nw397_[int(7)] = ((d_2255_v30_)[default__.safeIndex(p1, len(d_2255_v30_))] if (d_2217_v8_).f40 else d_2214_v5_)
            nw397_[int(8)] = ((d_2256_v31_)[(d_2221_v12_).fm36((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], d_2257_v32_, globalState)] if ((d_2221_v12_).fm36((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], d_2257_v32_, globalState)) in (d_2256_v31_) else d_2214_v5_)
            d_2258_v33_ = nw397_
            index354_ = default__.safeIndex(624, (d_2258_v33_).length(0))
            (d_2258_v33_)[index354_] = (d_2214_v5_) + (d_2214_v5_)
            index355_ = default__.safeIndex(624, (d_2258_v33_).length(0))
            (d_2258_v33_)[index355_] = _dafny.SeqWithoutIsStrInference([(p1) >= ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))])])
            d_2259_v34_: D2
            d_2259_v34_ = D2_DC5(d_2210_v3_)
            d_2260_v35_: D2
            d_2260_v35_ = D2_DC6(d_2259_v34_)
            d_2261_v36_: D2
            d_2261_v36_ = D2_DC6(d_2259_v34_)
            source37_ = d_2261_v36_
            if source37_.is_DC5:
                d_2262___mcc_h8_ = source37_.cf8
                d_2263_cf8_ = d_2262___mcc_h8_
                d_2264_v38_: _dafny.Map
                def iife154_():
                    coll88_ = _dafny.Set()
                    compr_88_: _dafny.Seq
                    for compr_88_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdk"))])).Elements:
                        d_2265_v37_: _dafny.Seq = compr_88_
                        if (d_2265_v37_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdk"))])):
                            coll88_ = coll88_.union(_dafny.Set([d_2265_v37_]))
                    return _dafny.Set(coll88_)
                d_2264_v38_ = _dafny.Map({(0) - (p1): (d_2221_v12_).fm36((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], iife154_()
                , globalState)})
                d_2266_v39_: D18
                d_2266_v39_ = D18_DC47(_dafny.Map({(d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]: d_2264_v38_}))
                (globalState).f24 = len((d_2266_v39_).cf67)
                d_2267_v40_: C1
                nw398_ = C1()
                nw398_.ctor__(d_2210_v3_, ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) != (len(d_2219_v10_)))
                d_2267_v40_ = nw398_
                d_2268_v41_: D17
                d_2268_v41_ = D17_DC44(d_2263_cf8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bknklrvx")))
                d_2269_v42_: _dafny.Seq
                d_2269_v42_ = _dafny.SeqWithoutIsStrInference([d_2268_v41_])
                d_2270_v43_: _dafny.Map
                d_2270_v43_ = _dafny.Map({((d_2258_v33_)[default__.safeIndex(624, (d_2258_v33_).length(0))])[default__.safeIndex(p1, len((d_2258_v33_)[default__.safeIndex(624, (d_2258_v33_).length(0))]))]: (_dafny.SeqWithoutIsStrInference([d_2268_v41_])) + (d_2269_v42_)})
                d_2269_v42_ = ((d_2270_v43_)[(d_2217_v8_).f40] if ((d_2217_v8_).f40) in (d_2270_v43_) else (d_2269_v42_) + (d_2269_v42_))
                index356_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                index357_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                rhs382_ = ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) - (p1)
                rhs383_ = (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]
                rhs384_ = (p1) * (p1)
                rhs385_ = d_2263_cf8_
                lhs319_ = d_2211_v4_
                lhs320_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                lhs321_ = d_2211_v4_
                lhs322_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                lhs323_ = globalState
                lhs324_ = globalState
                lhs319_[lhs320_] = rhs382_
                lhs321_[lhs322_] = rhs383_
                lhs323_.f1 = rhs384_
                lhs324_.f22 = rhs385_
            elif source37_.is_DC4:
                d_2271___mcc_h9_ = source37_.cf7
                d_2272_cf7_ = d_2271___mcc_h9_
                d_2273_v44_: _dafny.Array
                nw399_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_2273_v44_ = nw399_
                index358_ = default__.safeIndex(287, (d_2273_v44_).length(0))
                (d_2273_v44_)[index358_] = d_2211_v4_
                d_2274_v45_: _dafny.Map
                d_2274_v45_ = _dafny.Map({True: (d_2217_v8_).f40})
                d_2275_v46_: _dafny.Seq
                d_2275_v46_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_2210_v3_: d_2210_v3_})) | (default__.fm59((d_2217_v8_).f40, (d_2258_v33_)[default__.safeIndex(624, (d_2258_v33_).length(0))], globalState)), d_2274_v45_])
                d_2276_v47_: _dafny.Map
                d_2276_v47_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2272_cf7_ for d_2277_i1_ in range(default__.abs(174))])): default__.fm13(globalState)})
                d_2278_v49_: _dafny.MultiSet
                d_2278_v49_ = _dafny.MultiSet([(d_2217_v8_).f40, d_2210_v3_, (d_2217_v8_).f40, (d_2217_v8_).f40, (d_2217_v8_).f40])
                d_2279_v50_: _dafny.Map
                def iife155_():
                    coll89_ = _dafny.Map()
                    compr_89_: _dafny.MultiSet
                    for compr_89_ in (_dafny.Map({d_2278_v49_: p1})).keys.Elements:
                        d_2280_v48_: _dafny.MultiSet = compr_89_
                        if (d_2280_v48_) in (_dafny.Map({d_2278_v49_: p1})):
                            coll89_[d_2280_v48_] = len(_dafny.Map({d_2272_cf7_: (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]}))
                    return _dafny.Map(coll89_)
                d_2279_v50_ = _dafny.Map({(d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))]: default__.fm25(len((d_2276_v47_).set(p1, (d_2217_v8_).f40)), len(iife155_()
                ), globalState)})
                index359_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                index360_ = default__.safeIndex(287, (d_2273_v44_).length(0))
                rhs386_ = (len(d_2219_v10_)) <= (default__.fm27((0) - (p1), d_2272_cf7_, p1, globalState))
                rhs387_ = len((d_2275_v46_).set(default__.safeIndex(((d_2279_v50_)[d_2210_v3_] if (d_2210_v3_) in (d_2279_v50_) else len(p0)), len(d_2275_v46_)), d_2274_v45_))
                rhs388_ = d_2211_v4_
                lhs325_ = globalState
                lhs326_ = d_2211_v4_
                lhs327_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                lhs328_ = d_2273_v44_
                lhs329_ = default__.safeIndex(287, (d_2273_v44_).length(0))
                lhs325_.f22 = rhs386_
                lhs326_[lhs327_] = rhs387_
                lhs328_[lhs329_] = rhs388_
                d_2281_v51_: _dafny.Seq
                d_2281_v51_ = _dafny.SeqWithoutIsStrInference([d_2207_v0_, d_2207_v0_])
                d_2207_v0_ = (d_2281_v51_)[default__.safeIndex(403, len(d_2281_v51_))]
                d_2282_v52_: _dafny.Seq
                d_2282_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2210_v3_]), d_2278_v49_, d_2278_v49_])
                d_2283_v53_: C4
                nw400_ = C4()
                nw400_.ctor__(((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) > (p1), ((d_2282_v52_)[default__.safeIndex((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], len(d_2282_v52_))]) == (d_2278_v49_), (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], (d_2217_v8_).f40)
                d_2283_v53_ = nw400_
                d_2284_v54_: _dafny.Map
                d_2284_v54_ = _dafny.Map({(d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]: d_2207_v0_})
                d_2284_v54_ = (d_2284_v54_).set((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], d_2207_v0_)
            elif True:
                d_2285___mcc_h10_ = source37_.cf9
                d_2286_cf9_ = d_2285___mcc_h10_
                d_2287_v55_: D2
                d_2287_v55_ = D2_DC4(d_2216_v7_)
                d_2287_v55_ = d_2287_v55_
                d_2288_v56_: C0
                nw401_ = C0()
                nw401_.ctor__()
                d_2288_v56_ = nw401_
                d_2289_v57_: int
                d_2290_v58_: str
                d_2291_v59_: bool
                out40_: int
                out41_: str
                out42_: bool
                out40_, out41_, out42_ = (d_2217_v8_).m16((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], (d_2217_v8_).f40, d_2219_v10_, d_2258_v33_, globalState)
                d_2289_v57_ = out40_
                d_2290_v58_ = out41_
                d_2291_v59_ = out42_
                index361_ = default__.safeIndex(907, (d_2207_v0_).length(0))
                rhs389_ = ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) - ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))])
                rhs390_ = d_2291_v59_
                lhs330_ = globalState
                lhs331_ = d_2207_v0_
                lhs332_ = default__.safeIndex(907, (d_2207_v0_).length(0))
                lhs330_.f24 = rhs389_
                lhs331_[lhs332_] = rhs390_
            d_2292_v60_: _dafny.MultiSet
            d_2292_v60_ = _dafny.MultiSet([((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))] if (d_2217_v8_).f40 else (d_2217_v8_).f40), (d_2217_v8_).f40])
            index362_ = default__.safeIndex(62, (d_2211_v4_).length(0))
            rhs391_ = _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([(d_2217_v8_).f40])) + (d_2214_v5_)) + ((d_2258_v33_)[default__.safeIndex(624, (d_2258_v33_).length(0))]))
            rhs392_ = p1
            rhs393_ = ((0) - ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]) if (d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))] else default__.safeModuloInt(p1, p1))
            rhs394_ = not((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))])
            lhs333_ = globalState
            lhs334_ = d_2211_v4_
            lhs335_ = default__.safeIndex(62, (d_2211_v4_).length(0))
            lhs336_ = globalState
            d_2292_v60_ = rhs391_
            lhs333_.f1 = rhs392_
            lhs334_[lhs335_] = rhs393_
            lhs336_.f13 = rhs394_
            d_2293_v61_: _dafny.Map
            d_2293_v61_ = _dafny.Map({p1: d_2210_v3_})
            d_2294_v62_: _dafny.Map
            d_2294_v62_ = _dafny.Map({d_2219_v10_: ((d_2293_v61_)[p1] if (p1) in (d_2293_v61_) else (d_2217_v8_).f40)})
            d_2295_v63_: D17
            d_2295_v63_ = D17_DC43(d_2294_v62_)
            d_2295_v63_ = d_2295_v63_
        elif True:
            d_2216_v7_ = default__.fm42((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], p1, globalState)
            d_2296_v64_: _dafny.Map
            d_2296_v64_ = _dafny.Map({d_2216_v7_: (d_2217_v8_).f40})
            d_2296_v64_ = d_2296_v64_
            d_2297_v65_: _dafny.Array
            nw402_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_2297_v65_ = nw402_
            d_2298_v66_: D18
            d_2298_v66_ = D18_DC48(not (not(d_2210_v3_)) or (d_2210_v3_), (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], d_2297_v65_, (not(d_2210_v3_) if (d_2217_v8_).f40 else d_2210_v3_), d_2216_v7_)
            d_2299_v67_: _dafny.Array
            def lambda99_(d_2300_v10_):
                def lambda100_(d_2301_i2_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cs"))) + (d_2300_v10_)

                return lambda100_

            init59_ = lambda99_(d_2219_v10_)
            nw403_ = _dafny.Array(None, 2)
            for i0_59_ in range(nw403_.length(0)):
                nw403_[i0_59_] = init59_(i0_59_)
            d_2299_v67_ = nw403_
            d_2302_v68_: _dafny.Seq
            d_2302_v68_ = _dafny.SeqWithoutIsStrInference([d_2219_v10_])
            d_2303_v69_: T0
            nw404_ = C6()
            nw404_.ctor__(d_2297_v65_, d_2302_v68_, (d_2217_v8_).f40, ((d_2217_v8_).f40) and (True))
            d_2303_v69_ = nw404_
            d_2304_v70_: _dafny.Map
            d_2304_v70_ = _dafny.Map({not((d_2303_v69_).f29): p1})
            d_2305_v71_: _dafny.Map
            d_2305_v71_ = _dafny.Map({275: (p0)[default__.safeIndex(p1, len(p0))]})
            rhs395_ = d_2298_v66_
            rhs396_ = (d_2299_v67_ if d_2303_v69_.f28 else d_2299_v67_)
            rhs397_ = default__.safeDivisionInt(default__.safeDivisionInt((0) - ((0) - (p1)), (0) - (len(d_2304_v70_))), ((d_2305_v71_)[p1] if (p1) in (d_2305_v71_) else (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]))
            rhs398_ = d_2303_v69_
            lhs337_ = globalState
            d_2298_v66_ = rhs395_
            d_2299_v67_ = rhs396_
            lhs337_.f24 = rhs397_
            d_2303_v69_ = rhs398_
            (globalState).f22 = d_2303_v69_.f28
            source38_ = D8_DC21(((d_2217_v8_).f40 if d_2303_v69_.f28 else (d_2217_v8_).f40))
            if source38_.is_DC21:
                d_2306___mcc_h11_ = source38_.cf27
                d_2307_cf27_ = d_2306___mcc_h11_
                index363_ = default__.safeIndex(907, (d_2207_v0_).length(0))
                (d_2207_v0_)[index363_] = (d_2214_v5_)[default__.safeIndex((0) - ((0) - ((d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))])), len(d_2214_v5_))]
                index364_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                (d_2211_v4_)[index364_] = p1
                d_2302_v68_ = default__.fm28(D4_DC10((d_2303_v69_).f29, p1, len(d_2219_v10_)), p1, d_2216_v7_, (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))], globalState)
                d_2308_v72_: _dafny.Map
                d_2308_v72_ = _dafny.Map({(D4_DC9((d_2208_v1_).set(len(p0), default__.abs(p1)))).cf12: (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]})
                index365_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                (d_2211_v4_)[index365_] = default__.fm27(((d_2308_v72_)[d_2208_v1_] if (d_2208_v1_) in (d_2308_v72_) else (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]), _dafny.CodePoint('h'), p1, globalState)
            elif source38_.is_DC22:
                (globalState).f1 = (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]
                d_2219_v10_ = (d_2219_v10_) + (d_2219_v10_)
                d_2309_v73_: _dafny.Set
                d_2309_v73_ = _dafny.Set({d_2216_v7_, _dafny.CodePoint('v'), d_2216_v7_, d_2216_v7_})
                d_2310_v74_: _dafny.Map
                d_2310_v74_ = _dafny.Map({d_2207_v0_: d_2210_v3_})
                index366_ = default__.safeIndex(62, (d_2211_v4_).length(0))
                (d_2211_v4_)[index366_] = ((d_2305_v71_)[(len(d_2309_v73_)) - (len(d_2310_v74_))] if ((len(d_2309_v73_)) - (len(d_2310_v74_))) in (d_2305_v71_) else p1)
                d_2311_v75_: _dafny.Map
                d_2311_v75_ = _dafny.Map({(d_2217_v8_).f40: d_2216_v7_})
                d_2312_v76_: _dafny.Map
                d_2312_v76_ = _dafny.Map({d_2311_v75_: d_2214_v5_})
                d_2313_v77_: _dafny.Seq
                d_2313_v77_ = d_2214_v5_
                d_2314_v78_: _dafny.Seq
                d_2314_v78_ = _dafny.SeqWithoutIsStrInference([((d_2312_v76_)[d_2311_v75_] if (d_2311_v75_) in (d_2312_v76_) else (d_2313_v77_))])
                (globalState).f24 = (0) - (len((d_2314_v78_)[default__.safeIndex(592, len(d_2314_v78_))]))
            elif True:
                d_2315___mcc_h12_ = source38_.cf26
                d_2316_cf26_ = d_2315___mcc_h12_
                (globalState).f1 = (d_2211_v4_)[default__.safeIndex(62, (d_2211_v4_).length(0))]
                (globalState).f23 = d_2216_v7_
                index367_ = default__.safeIndex(383, (d_2299_v67_).length(0))
                (d_2299_v67_)[index367_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvyhgniu")))
                index368_ = default__.safeIndex(383, (d_2299_v67_).length(0))
                (d_2299_v67_)[index368_] = (d_2302_v68_)[default__.safeIndex(p1, len(d_2302_v68_))]
                d_2317_v79_: _dafny.MultiSet
                d_2317_v79_ = _dafny.MultiSet([False, d_2316_cf26_.f28])
                d_2317_v79_ = d_2317_v79_
        d_2318_v80_: _dafny.Array
        def lambda101_(d_2319_v2_):
            def lambda102_(d_2320_i3_):
                return D13_DC34(d_2319_v2_)

            return lambda102_

        init60_ = lambda101_(d_2209_v2_)
        nw405_ = _dafny.Array(None, 12)
        for i0_60_ in range(nw405_.length(0)):
            nw405_[i0_60_] = init60_(i0_60_)
        d_2318_v80_ = nw405_
        d_2318_v80_ = d_2318_v80_
        d_2321_v81_: _dafny.Seq
        d_2321_v81_ = d_2214_v5_
        r0 = default__.fm59((d_2207_v0_)[default__.safeIndex(907, (d_2207_v0_).length(0))], (d_2321_v81_), globalState)
        return r0

    def m10(self, p0, p1, globalState):
        d_2322_v0_: D2
        d_2322_v0_ = D2_DC5(p1)
        d_2323_v1_: D2
        d_2323_v1_ = D2_DC6(d_2322_v0_)
        d_2323_v1_ = d_2323_v1_
        d_2324_v3_: _dafny.Array
        def lambda103_(d_2325_p1_):
            def lambda104_(d_2326_i0_):
                def iife156_():
                    coll90_ = _dafny.Map()
                    compr_90_: D15
                    for compr_90_ in (_dafny.SeqWithoutIsStrInference([D15_DC40(d_2325_p1_, d_2325_p1_), D15_DC40(d_2325_p1_, d_2325_p1_)])).Elements:
                        d_2327_v2_: D15 = compr_90_
                        if (d_2327_v2_) in (_dafny.SeqWithoutIsStrInference([D15_DC40(d_2325_p1_, d_2325_p1_), D15_DC40(d_2325_p1_, d_2325_p1_)])):
                            coll90_[d_2327_v2_] = d_2325_p1_
                    return _dafny.Map(coll90_)
                return iife156_()
                

            return lambda104_

        init61_ = lambda103_(p1)
        nw406_ = _dafny.Array(None, 7)
        for i0_61_ in range(nw406_.length(0)):
            nw406_[i0_61_] = init61_(i0_61_)
        d_2324_v3_ = nw406_
        d_2328_v4_: _dafny.Map
        d_2328_v4_ = _dafny.Map({p1: True})
        d_2329_v5_: _dafny.Array
        nw407_ = _dafny.Array(None, 10)
        nw407_[int(0)] = p1
        nw407_[int(1)] = p1
        nw407_[int(2)] = ((d_2328_v4_)[p1] if (p1) in (d_2328_v4_) else p1)
        nw407_[int(3)] = p1
        nw407_[int(4)] = p1
        nw407_[int(5)] = p1
        nw407_[int(6)] = not(p1)
        nw407_[int(7)] = p1
        nw407_[int(8)] = p1
        nw407_[int(9)] = p1
        d_2329_v5_ = nw407_
        d_2330_v6_: _dafny.Map
        d_2330_v6_ = _dafny.Map({d_2324_v3_: d_2329_v5_})
        d_2330_v6_ = (d_2330_v6_).set(d_2324_v3_, d_2329_v5_)
        d_2331_v7_: _dafny.Array
        nw408_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
        d_2331_v7_ = nw408_
        d_2332_v8_: _dafny.Seq
        d_2332_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypl"))
        index369_ = default__.safeIndex(81, (d_2331_v7_).length(0))
        (d_2331_v7_)[index369_] = d_2332_v8_
        d_2333_v9_: str
        d_2333_v9_ = _dafny.CodePoint('h')
        index370_ = default__.safeIndex(81, (d_2331_v7_).length(0))
        (d_2331_v7_)[index370_] = (d_2332_v8_).set(default__.safeIndex((len(d_2332_v8_)) + (default__.fm25(p0, p0, globalState)), len(d_2332_v8_)), d_2333_v9_)
        d_2334_v10_: _dafny.Array
        def lambda105_(d_2335_v9_, d_2336_p0_):
            def lambda106_(d_2337_i2_):
                return (d_2337_i2_) + (len(_dafny.Map({d_2335_v9_: len(_dafny.Map({d_2336_p0_: d_2335_v9_}))})))

            return lambda106_

        init62_ = lambda105_(d_2333_v9_, p0)
        nw409_ = _dafny.Array(None, 17)
        for i0_62_ in range(nw409_.length(0)):
            nw409_[i0_62_] = init62_(i0_62_)
        d_2334_v10_ = nw409_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2334_v10_).length(0)):
            d_2338_i1_: int = guard_loop_7_
            if (True) and (((0) <= (d_2338_i1_)) and ((d_2338_i1_) < ((d_2334_v10_).length(0)))):
                (d_2334_v10_)[(d_2338_i1_)] = (d_2338_i1_) - (p0)
        d_2339_v11_: _dafny.Map
        d_2339_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): p0})
        d_2340_v12_: _dafny.Seq
        d_2340_v12_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2341_v13_: _dafny.Map
        d_2341_v13_ = _dafny.Map({((d_2339_v11_)[d_2340_v12_] if (d_2340_v12_) in (d_2339_v11_) else p0): False})
        d_2341_v13_ = d_2341_v13_
        d_2342_v14_: _dafny.Seq
        d_2342_v14_ = _dafny.SeqWithoutIsStrInference([d_2334_v10_, d_2334_v10_, d_2334_v10_])
        (globalState).f24 = len((d_2342_v14_) + ((_dafny.SeqWithoutIsStrInference([d_2334_v10_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_2334_v10_]))), d_2334_v10_)))


class C9(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
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
    def ctor__(self, f28, f29):
        (self).f28 = f28
        (self)._f29 = f29

    def fm11(self, p0, p1, p2, globalState):
        return self.f28

    def m5(self, p0, globalState):
        r0: bool = False
        if (self).f29:
            d_2343_v0_: _dafny.Array
            def lambda107_(d_2344_p0_):
                def lambda108_(d_2345_i0_):
                    return (_dafny.Map({(self).f29: (0) - ((0) - (d_2344_p0_))})) | (_dafny.Map({not((self).f29): d_2344_p0_}))

                return lambda108_

            init63_ = lambda107_(p0)
            nw410_ = _dafny.Array(None, 5)
            for i0_63_ in range(nw410_.length(0)):
                nw410_[i0_63_] = init63_(i0_63_)
            d_2343_v0_ = nw410_
            d_2346_v1_: _dafny.Map
            d_2346_v1_ = _dafny.Map({(self).f29: p0})
            index371_ = default__.safeIndex(661, (d_2343_v0_).length(0))
            (d_2343_v0_)[index371_] = d_2346_v1_
            index372_ = default__.safeIndex(661, (d_2343_v0_).length(0))
            (d_2343_v0_)[index372_] = (d_2346_v1_) | (d_2346_v1_)
            d_2347_v2_: _dafny.Seq
            d_2347_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({397})])
            if (p0) <= ((0) - ((((d_2343_v0_)[default__.safeIndex(661, (d_2343_v0_).length(0))])[self.f28] if (self.f28) in ((d_2343_v0_)[default__.safeIndex(661, (d_2343_v0_).length(0))]) else len(d_2347_v2_)))):
                d_2348_v3_: _dafny.Map
                d_2348_v3_ = _dafny.Map({self.f28: (self).f29})
                d_2348_v3_ = (d_2348_v3_).set(self.f28, True)
                d_2349_v4_: _dafny.Seq
                d_2349_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikj"))
                d_2350_v5_: _dafny.Seq
                d_2350_v5_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")) for d_2351_i1_ in range(default__.abs(65))])).set(default__.safeIndex(467, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")) for d_2352_i1_ in range(default__.abs(65))]))), d_2349_v4_)])
                d_2353_v6_: _dafny.Seq
                d_2353_v6_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])
                d_2348_v3_ = (d_2348_v3_).set((self).f29, ((self).fm11(self.f28, (d_2350_v5_)[default__.safeIndex(len(d_2346_v1_), len(d_2350_v5_))], p0, globalState)) in (d_2353_v6_))
                d_2354_v7_: str
                d_2354_v7_ = _dafny.CodePoint('e')
                d_2355_v8_: _dafny.Map
                d_2355_v8_ = _dafny.Map({(p0) != (p0): d_2354_v7_})
                d_2355_v8_ = (d_2355_v8_).set(self.f28, d_2354_v7_)
                (globalState).f24 = p0
                d_2356_v9_: _dafny.Set
                d_2356_v9_ = _dafny.Set({p0, -740, p0})
                d_2357_v10_: _dafny.Array
                nw411_ = _dafny.Array(int(0), 26)
                d_2357_v10_ = nw411_
                d_2358_v11_: _dafny.Map
                d_2358_v11_ = _dafny.Map({(self).f29: d_2357_v10_})
                index373_ = default__.safeIndex(135, (d_2357_v10_).length(0))
                (d_2357_v10_)[index373_] = (p0) + (p0)
                d_2359_v12_: D3
                d_2359_v12_ = D3_DC7(d_2358_v11_)
                index374_ = default__.safeIndex(135, (d_2357_v10_).length(0))
                rhs399_ = d_2356_v9_
                rhs400_ = ((d_2359_v12_).cf10) | ((d_2358_v11_) | (d_2358_v11_))
                rhs401_ = p0
                lhs338_ = d_2357_v10_
                lhs339_ = default__.safeIndex(135, (d_2357_v10_).length(0))
                d_2356_v9_ = rhs399_
                d_2358_v11_ = rhs400_
                lhs338_[lhs339_] = rhs401_
            elif True:
                d_2360_v13_: _dafny.Seq
                d_2360_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngv"))
                d_2361_v14_: _dafny.Seq
                d_2361_v14_ = _dafny.SeqWithoutIsStrInference([d_2360_v13_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcknd"))])
                d_2362_v15_: _dafny.Seq
                d_2362_v15_ = _dafny.SeqWithoutIsStrInference([(self).fm11((self).f29, d_2361_v14_, p0, globalState)])
                d_2363_v16_: _dafny.Set
                d_2363_v16_ = _dafny.Set({d_2362_v15_})
                d_2364_v17_: _dafny.Map
                d_2364_v17_ = _dafny.Map({default__.fm12(p0, (self).f29, globalState): (self).f29})
                d_2365_v18_: _dafny.Map
                d_2365_v18_ = _dafny.Map({d_2362_v15_: ((d_2364_v17_)[p0] if (p0) in (d_2364_v17_) else self.f28)})
                d_2366_v20_: _dafny.Seq
                d_2366_v20_ = _dafny.SeqWithoutIsStrInference([d_2362_v15_, d_2362_v15_, d_2362_v15_])
                def iife157_():
                    coll91_ = _dafny.Set()
                    compr_91_: _dafny.Seq
                    for compr_91_ in (d_2365_v18_).keys.Elements:
                        d_2367_v19_: _dafny.Seq = compr_91_
                        if (d_2367_v19_) in (d_2365_v18_):
                            coll91_ = coll91_.union(_dafny.Set([d_2367_v19_]))
                    return _dafny.Set(coll91_)
                def iife158_():
                    coll92_ = _dafny.Set()
                    compr_92_: _dafny.Seq
                    for compr_92_ in (d_2366_v20_).Elements:
                        d_2368_v21_: _dafny.Seq = compr_92_
                        if (d_2368_v21_) in (d_2366_v20_):
                            coll92_ = coll92_.union(_dafny.Set([d_2368_v21_]))
                    return _dafny.Set(coll92_)
                rhs402_ = (iife157_()
                ) | (iife158_()
                )
                rhs403_ = self.f28
                rhs404_ = d_2360_v13_
                lhs340_ = globalState
                d_2363_v16_ = rhs402_
                lhs340_.f27 = rhs403_
                d_2360_v13_ = rhs404_
                d_2361_v14_ = d_2361_v14_
                (globalState).f1 = len((d_2346_v1_) | (((d_2343_v0_)[default__.safeIndex(661, (d_2343_v0_).length(0))]).set((self).f29, p0)))
                d_2369_v22_: _dafny.Array
                nw412_ = _dafny.Array(None, 27)
                nw412_[int(0)] = (self).f29
                nw412_[int(1)] = False
                nw412_[int(2)] = self.f28
                nw412_[int(3)] = self.f28
                nw412_[int(4)] = not((self).f29)
                nw412_[int(5)] = self.f28
                nw412_[int(6)] = self.f28
                nw412_[int(7)] = True
                nw412_[int(8)] = self.f28
                nw412_[int(9)] = not((self).f29)
                nw412_[int(10)] = self.f28
                nw412_[int(11)] = not((self).f29)
                nw412_[int(12)] = False
                nw412_[int(13)] = self.f28
                nw412_[int(14)] = not((self).f29)
                nw412_[int(15)] = self.f28
                nw412_[int(16)] = self.f28
                nw412_[int(17)] = self.f28
                nw412_[int(18)] = self.f28
                nw412_[int(19)] = self.f28
                nw412_[int(20)] = self.f28
                nw412_[int(21)] = (self).f29
                nw412_[int(22)] = (d_2362_v15_)[default__.safeIndex(p0, len(d_2362_v15_))]
                nw412_[int(23)] = self.f28
                nw412_[int(24)] = (self).f29
                nw412_[int(25)] = not((self).f29)
                nw412_[int(26)] = (self).f29
                d_2369_v22_ = nw412_
                d_2370_v23_: _dafny.Map
                d_2370_v23_ = _dafny.Map({589: d_2369_v22_})
                d_2371_v24_: D1
                d_2371_v24_ = D1_DC2(618, p0, d_2360_v13_, self.f28)
                d_2372_v25_: _dafny.Set
                d_2372_v25_ = _dafny.Set({(0) - ((d_2371_v24_).cf3), p0})
                d_2373_v26_: _dafny.Map
                d_2373_v26_ = _dafny.Map({d_2372_v25_: self.f28})
                d_2370_v23_ = (d_2370_v23_).set(len(d_2373_v26_), d_2369_v22_)
                (globalState).f24 = p0
            (globalState).f13 = (self).f29
            d_2374_v27_: _dafny.Array
            def lambda109_(d_2375_i2_):
                return default__.safeDivisionInt(d_2375_i2_, 548)

            init64_ = lambda109_
            nw413_ = _dafny.Array(None, 19)
            for i0_64_ in range(nw413_.length(0)):
                nw413_[i0_64_] = init64_(i0_64_)
            d_2374_v27_ = nw413_
            index375_ = default__.safeIndex(722, (d_2374_v27_).length(0))
            (d_2374_v27_)[index375_] = p0
            d_2376_v28_: _dafny.Seq
            d_2376_v28_ = _dafny.SeqWithoutIsStrInference([default__.fm12(p0, False, globalState), p0])
            index376_ = default__.safeIndex(722, (d_2374_v27_).length(0))
            (d_2374_v27_)[index376_] = len(((d_2376_v28_) + (d_2376_v28_)).set(default__.safeIndex(default__.safeModuloInt(21, 605), len((d_2376_v28_) + (d_2376_v28_))), (0) - (default__.safeDivisionInt(p0, p0))))
            d_2377_v29_: C2
            nw414_ = C2()
            nw414_.ctor__()
            d_2377_v29_ = nw414_
        elif True:
            (globalState).f27 = self.f28
            (globalState).f13 = default__.fm38(globalState)
            d_2378_v30_: _dafny.MultiSet
            d_2378_v30_ = _dafny.MultiSet([(self).f29, self.f28])
            d_2379_v31_: _dafny.Array
            def lambda110_(d_2380_p0_):
                def lambda111_(d_2381_i3_):
                    return default__.safeDivisionInt(d_2381_i3_, d_2380_p0_)

                return lambda111_

            init65_ = lambda110_(p0)
            nw415_ = _dafny.Array(None, 9)
            for i0_65_ in range(nw415_.length(0)):
                nw415_[i0_65_] = init65_(i0_65_)
            d_2379_v31_ = nw415_
            d_2382_v32_: _dafny.Set
            d_2382_v32_ = _dafny.Set({d_2379_v31_})
            d_2383_v33_: _dafny.Set
            d_2383_v33_ = _dafny.Set({128})
            d_2384_v34_: _dafny.Seq
            d_2384_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdwiwcd"))
            d_2385_v35_: _dafny.Array
            nw416_ = _dafny.Array(None, 15)
            nw416_[int(0)] = (self).f29
            nw416_[int(1)] = (d_2378_v30_).issubset(d_2378_v30_)
            nw416_[int(2)] = self.f28
            nw416_[int(3)] = (d_2382_v32_).isdisjoint(d_2382_v32_)
            nw416_[int(4)] = self.f28
            nw416_[int(5)] = (self).f29
            nw416_[int(6)] = True
            nw416_[int(7)] = (d_2383_v33_).ispropersubset(d_2383_v33_)
            nw416_[int(8)] = not(not((d_2384_v34_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2386_i4_ in range(default__.abs(-13))]))))
            nw416_[int(9)] = (self).f29
            nw416_[int(10)] = (self).f29
            nw416_[int(11)] = self.f28
            nw416_[int(12)] = (self).f29
            nw416_[int(13)] = self.f28
            nw416_[int(14)] = (self.f28) and ((self).f29)
            d_2385_v35_ = nw416_
            d_2387_v36_: _dafny.Seq
            d_2387_v36_ = _dafny.SeqWithoutIsStrInference([d_2384_v34_, d_2384_v34_])
            index377_ = default__.safeIndex(293, (d_2385_v35_).length(0))
            (d_2385_v35_)[index377_] = (self).fm11(not(not(self.f28)), d_2387_v36_, p0, globalState)
            index378_ = default__.safeIndex(45, (d_2379_v31_).length(0))
            (d_2379_v31_)[index378_] = (p0) + (p0)
            d_2388_v37_: _dafny.Seq
            d_2388_v37_ = _dafny.SeqWithoutIsStrInference([self.f28])
            d_2389_v38_: _dafny.Seq
            d_2389_v38_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(self).f29, self.f28])) + (d_2388_v37_)])
            d_2390_v39_: C1
            nw417_ = C1()
            nw417_.ctor__(not((self).f29), False)
            d_2390_v39_ = nw417_
            d_2391_v40_: _dafny.Seq
            d_2391_v40_ = _dafny.SeqWithoutIsStrInference([p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmnjrin")))])
            d_2392_v41_: _dafny.MultiSet
            d_2392_v41_ = _dafny.MultiSet([len(d_2391_v40_), p0, p0])
            d_2393_v42_: _dafny.Map
            d_2393_v42_ = _dafny.Map({d_2390_v39_: d_2392_v41_})
            index379_ = default__.safeIndex(293, (d_2385_v35_).length(0))
            index380_ = default__.safeIndex(45, (d_2379_v31_).length(0))
            rhs405_ = (p0) > (p0)
            rhs406_ = p0
            rhs407_ = (((d_2393_v42_)[d_2390_v39_] if (d_2390_v39_) in (d_2393_v42_) else d_2392_v41_)).issubset(d_2392_v41_)
            rhs408_ = (0) - (default__.fm27(p0, default__.fm42(375, p0, globalState), (0) - (p0), globalState))
            rhs409_ = (d_2389_v38_) + ((default__.fm46(self.f28, globalState)) + ((d_2389_v38_).set(default__.safeIndex((0) - ((0) - ((0) - (p0))), len(d_2389_v38_)), _dafny.SeqWithoutIsStrInference([self.f28, (self).f29]))))
            lhs341_ = d_2385_v35_
            lhs342_ = default__.safeIndex(293, (d_2385_v35_).length(0))
            lhs343_ = d_2379_v31_
            lhs344_ = default__.safeIndex(45, (d_2379_v31_).length(0))
            lhs345_ = globalState
            lhs346_ = globalState
            lhs341_[lhs342_] = rhs405_
            lhs343_[lhs344_] = rhs406_
            lhs345_.f22 = rhs407_
            lhs346_.f1 = rhs408_
            d_2389_v38_ = rhs409_
            d_2385_v35_ = d_2385_v35_
            index381_ = default__.safeIndex(45, (d_2379_v31_).length(0))
            (d_2379_v31_)[index381_] = default__.safeDivisionInt(default__.fm25((0) - ((d_2391_v40_)[default__.safeIndex(-413, len(d_2391_v40_))]), (d_2379_v31_)[default__.safeIndex(45, (d_2379_v31_).length(0))], globalState), ((d_2379_v31_)[default__.safeIndex(45, (d_2379_v31_).length(0))]) - (p0))
        d_2394_v43_: _dafny.Seq
        d_2394_v43_ = _dafny.SeqWithoutIsStrInference([False])
        d_2395_v44_: D15
        d_2395_v44_ = D15_DC37(_dafny.MultiSet([d_2394_v43_, d_2394_v43_]))
        source39_ = d_2395_v44_
        if source39_.is_DC38:
            d_2396___mcc_h0_ = source39_.cf53
            d_2397___mcc_h1_ = source39_.cf54
            d_2398___mcc_h2_ = source39_.cf55
            d_2399_cf55_ = d_2398___mcc_h2_
            d_2400_cf54_ = d_2397___mcc_h1_
            d_2401_cf53_ = d_2396___mcc_h0_
            (globalState).f22 = (self).f29
            d_2402_v45_: _dafny.Seq
            d_2402_v45_ = _dafny.SeqWithoutIsStrInference([d_2400_cf54_ for d_2403_i5_ in range(default__.abs(190))])
            d_2404_v46_: str
            d_2404_v46_ = _dafny.CodePoint('x')
            d_2405_v47_: _dafny.Array
            nw418_ = _dafny.Array(None, 7)
            nw418_[int(0)] = default__.fm42(p0, p0, globalState)
            nw418_[int(1)] = d_2404_v46_
            nw418_[int(2)] = d_2404_v46_
            nw418_[int(3)] = d_2404_v46_
            nw418_[int(4)] = d_2404_v46_
            nw418_[int(5)] = d_2404_v46_
            nw418_[int(6)] = d_2404_v46_
            d_2405_v47_ = nw418_
            d_2406_v48_: _dafny.Array
            nw419_ = _dafny.Array(None, 7)
            nw419_[int(0)] = 953
            nw419_[int(1)] = p0
            nw419_[int(2)] = p0
            nw419_[int(3)] = d_2400_cf54_
            nw419_[int(4)] = 78
            nw419_[int(5)] = d_2400_cf54_
            nw419_[int(6)] = p0
            d_2406_v48_ = nw419_
            d_2407_v49_: _dafny.Map
            d_2407_v49_ = _dafny.Map({d_2405_v47_: d_2406_v48_})
            d_2408_v50_: _dafny.Map
            d_2408_v50_ = _dafny.Map({(d_2402_v45_): (d_2407_v49_) | (_dafny.Map({d_2405_v47_: d_2406_v48_}))})
            d_2409_v51_: _dafny.Seq
            d_2409_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pn"))
            d_2410_v52_: _dafny.Seq
            d_2410_v52_ = _dafny.SeqWithoutIsStrInference([len(d_2409_v51_)])
            d_2408_v50_ = (d_2408_v50_).set(d_2410_v52_, (d_2407_v49_ if False else d_2407_v49_))
            d_2411_v53_: C0
            nw420_ = C0()
            nw420_.ctor__()
            d_2411_v53_ = nw420_
            rhs410_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgbxmv"))
            rhs411_ = len((d_2394_v43_) + (default__.fm62(globalState)))
            lhs347_ = globalState
            d_2409_v51_ = rhs410_
            lhs347_.f24 = rhs411_
        elif source39_.is_DC39:
            d_2412___mcc_h3_ = source39_.cf56
            d_2413_cf56_ = d_2412___mcc_h3_
            d_2414_v54_: _dafny.Map
            d_2414_v54_ = _dafny.Map({self.f28: p0})
            d_2415_v55_: _dafny.Set
            d_2415_v55_ = _dafny.Set({d_2414_v54_, _dafny.Map({self.f28: p0})})
            (globalState).f1 = ((-985) * (len(d_2415_v55_))) - (p0)
            rhs412_ = (self).f29
            rhs413_ = p0
            lhs348_ = self
            lhs349_ = globalState
            lhs348_.f28 = rhs412_
            lhs349_.f1 = rhs413_
            d_2413_cf56_ = d_2413_cf56_
            (globalState).f27 = not((self).f29)
        elif source39_.is_DC40:
            d_2416___mcc_h4_ = source39_.cf57
            d_2417___mcc_h5_ = source39_.cf58
            d_2418_cf58_ = d_2417___mcc_h5_
            d_2419_cf57_ = d_2416___mcc_h4_
            d_2420_v56_: _dafny.Map
            d_2420_v56_ = _dafny.Map({p0: self.f28})
            (globalState).f1 = len(d_2420_v56_)
            d_2421_v57_: _dafny.Set
            d_2421_v57_ = _dafny.Set({d_2419_cf57_, not((self).f29)})
            (globalState).f21 = ((d_2421_v57_) - (d_2421_v57_)).issubset(d_2421_v57_)
            d_2422_v58_: str
            d_2422_v58_ = _dafny.CodePoint('d')
            d_2423_v59_: _dafny.Set
            d_2423_v59_ = _dafny.Set({p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvyfwjsi"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvyfwjsi")))), d_2422_v58_))})
            d_2424_v60_: D12
            d_2424_v60_ = D12_DC31(d_2423_v59_)
            pat_let_tv61_ = d_2423_v59_
            def iife159_(_pat_let33_0):
                def iife160_(d_2425_dt__update__tmp_h0_):
                    def iife161_(_pat_let34_0):
                        def iife162_(d_2426_dt__update_hcf48_h0_):
                            return D12_DC31(d_2426_dt__update_hcf48_h0_)
                        return iife162_(_pat_let34_0)
                    return iife161_(pat_let_tv61_)
                return iife160_(_pat_let33_0)
            source40_ = (iife159_(d_2424_v60_) if (self).f29 else d_2424_v60_)
            if source40_.is_DC32:
                d_2427_v61_: _dafny.Array
                def lambda112_(d_2428_p0_):
                    def lambda113_(d_2429_i6_):
                        return default__.safeDivisionInt(d_2429_i6_, (0) - (d_2428_p0_))

                    return lambda113_

                init66_ = lambda112_(p0)
                nw421_ = _dafny.Array(None, 12)
                for i0_66_ in range(nw421_.length(0)):
                    nw421_[i0_66_] = init66_(i0_66_)
                d_2427_v61_ = nw421_
                d_2430_v62_: _dafny.Map
                d_2430_v62_ = _dafny.Map({p0: d_2427_v61_})
                d_2427_v61_ = (d_2427_v61_ if (p0) >= (p0) else ((d_2430_v62_)[p0] if (p0) in (d_2430_v62_) else d_2427_v61_))
                (globalState).f22 = not(d_2418_cf58_)
                (globalState).f24 = 420
                d_2431_v63_: _dafny.Array
                nw422_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                d_2431_v63_ = nw422_
                d_2432_v64_: _dafny.Seq
                d_2432_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tygcp"))
                d_2433_v65_: _dafny.Map
                d_2433_v65_ = _dafny.Map({d_2431_v63_: len(d_2432_v64_)})
                d_2433_v65_ = (d_2433_v65_).set(d_2431_v63_, p0)
            elif source40_.is_DC31:
                d_2434___mcc_h8_ = source40_.cf48
                d_2435_cf48_ = d_2434___mcc_h8_
                d_2436_v66_: C8
                nw423_ = C8()
                nw423_.ctor__()
                d_2436_v66_ = nw423_
                (globalState).f1 = p0
                d_2394_v43_ = ((d_2394_v43_) + (d_2394_v43_)) + ((d_2394_v43_) + (d_2394_v43_))
                d_2437_v67_: _dafny.Map
                d_2437_v67_ = _dafny.Map({d_2419_cf57_: d_2422_v58_})
                (globalState).f23 = ((d_2437_v67_)[not(not((self).f29))] if (not(not((self).f29))) in (d_2437_v67_) else default__.fm42(p0, p0, globalState))
            elif True:
                d_2438___mcc_h9_ = source40_.cf49
                d_2439_cf49_ = d_2438___mcc_h9_
                d_2440_v68_: _dafny.Seq
                d_2440_v68_ = _dafny.SeqWithoutIsStrInference([d_2422_v58_, d_2422_v58_])
                d_2441_v69_: D9
                d_2441_v69_ = D9_DC25((self).f29, p0, len(d_2420_v56_), not(self.f28), len(d_2440_v68_))
                d_2442_v70_: _dafny.MultiSet
                d_2442_v70_ = _dafny.MultiSet([853])
                d_2443_v71_: _dafny.Array
                nw424_ = _dafny.Array(None, 17)
                nw424_[int(0)] = (not(d_2419_cf57_) if not(self.f28) else d_2418_cf58_)
                nw424_[int(1)] = (d_2441_v69_).cf34
                nw424_[int(2)] = ((self).f29 if d_2418_cf58_ else d_2419_cf57_)
                nw424_[int(3)] = self.f28
                nw424_[int(4)] = self.f28
                nw424_[int(5)] = d_2419_cf57_
                nw424_[int(6)] = default__.fm18(p0, default__.fm25(p0, len(d_2394_v43_), globalState), p0, globalState)
                nw424_[int(7)] = (p0) <= (p0)
                nw424_[int(8)] = not(((d_2442_v70_).set(868, default__.abs(p0))).isdisjoint(default__.fm24(d_2423_v59_, d_2418_cf58_, globalState)))
                nw424_[int(9)] = d_2418_cf58_
                nw424_[int(10)] = d_2419_cf57_
                nw424_[int(11)] = not(not(d_2418_cf58_))
                nw424_[int(12)] = self.f28
                nw424_[int(13)] = False
                nw424_[int(14)] = (d_2394_v43_)[default__.safeIndex(p0, len(d_2394_v43_))]
                nw424_[int(15)] = (self).f29
                nw424_[int(16)] = (self).f29
                d_2443_v71_ = nw424_
                index382_ = default__.safeIndex(593, (d_2443_v71_).length(0))
                (d_2443_v71_)[index382_] = d_2419_cf57_
                index383_ = default__.safeIndex(593, (d_2443_v71_).length(0))
                (d_2443_v71_)[index383_] = (d_2394_v43_)[default__.safeIndex(p0, len(d_2394_v43_))]
                d_2444_v72_: _dafny.Seq
                d_2444_v72_ = _dafny.SeqWithoutIsStrInference([303, len(d_2440_v68_)])
                d_2445_v73_: D3
                d_2445_v73_ = D3_DC8(d_2440_v68_)
                d_2446_v74_: _dafny.Map
                d_2446_v74_ = _dafny.Map({(d_2444_v72_)[default__.safeIndex(p0, len(d_2444_v72_))]: len((d_2445_v73_).cf11)})
                (globalState).f1 = (p0) - (len((d_2446_v74_).set(p0, len(d_2440_v68_))))
                d_2447_v75_: C4
                nw425_ = C4()
                nw425_.ctor__((not(d_2419_cf57_) if True else (d_2443_v71_)[default__.safeIndex(593, (d_2443_v71_).length(0))]), not(self.f28), d_2419_cf57_, not((_dafny.MultiSet(d_2394_v43_)).ispropersubset(_dafny.MultiSet([(self).f29, (self).f29, d_2418_cf58_]))))
                d_2447_v75_ = nw425_
                d_2448_v76_: _dafny.Array
                nw426_ = _dafny.Array(None, 25)
                nw426_[int(0)] = default__.safeModuloInt(p0, -414)
                nw426_[int(1)] = p0
                nw426_[int(2)] = (0) - (p0)
                nw426_[int(3)] = p0
                nw426_[int(4)] = p0
                nw426_[int(5)] = (0) - (default__.fm12(p0, not((d_2447_v75_).f38), globalState))
                nw426_[int(6)] = 592
                nw426_[int(7)] = default__.safeModuloInt(len(d_2444_v72_), p0)
                nw426_[int(8)] = p0
                nw426_[int(9)] = ((D1_DC2(p0, len(_dafny.Map({True: p0})), d_2440_v68_, not(d_2419_cf57_))).cf3 if True else p0)
                nw426_[int(10)] = (p0) + (p0)
                nw426_[int(11)] = p0
                nw426_[int(12)] = p0
                nw426_[int(13)] = p0
                nw426_[int(14)] = (p0) + (p0)
                nw426_[int(15)] = (p0) + ((_dafny.MultiSet([len(d_2444_v72_), 398])).cardinality)
                nw426_[int(16)] = (default__.fm25(p0, p0, globalState)) - (len(d_2440_v68_))
                nw426_[int(17)] = p0
                nw426_[int(18)] = p0
                nw426_[int(19)] = len((d_2440_v68_).set(default__.safeIndex(319, len(d_2440_v68_)), (d_2440_v68_)[default__.safeIndex(p0, len(d_2440_v68_))]))
                nw426_[int(20)] = p0
                nw426_[int(21)] = p0
                nw426_[int(22)] = p0
                nw426_[int(23)] = p0
                nw426_[int(24)] = (p0) * (p0)
                d_2448_v76_ = nw426_
                index384_ = default__.safeIndex(86, (d_2448_v76_).length(0))
                (d_2448_v76_)[index384_] = p0
                d_2449_v77_: _dafny.Map
                d_2449_v77_ = _dafny.Map({d_2448_v76_: ((d_2446_v74_)[(_dafny.MultiSet(default__.fm19(-465, (self).f29, (d_2443_v71_)[default__.safeIndex(593, (d_2443_v71_).length(0))], (d_2447_v75_).f38, globalState))).cardinality] if ((_dafny.MultiSet(default__.fm19(-465, (self).f29, (d_2443_v71_)[default__.safeIndex(593, (d_2443_v71_).length(0))], (d_2447_v75_).f38, globalState))).cardinality) in (d_2446_v74_) else p0)})
                index385_ = default__.safeIndex(86, (d_2448_v76_).length(0))
                (d_2448_v76_)[index385_] = len((d_2449_v77_) | (d_2449_v77_))
            d_2450_v78_: D15
            d_2450_v78_ = D15_DC38((self).f29, p0, p0)
            (globalState).f1 = (d_2450_v78_).cf55
        elif source39_.is_DC37:
            d_2451___mcc_h6_ = source39_.cf52
            d_2452_cf52_ = d_2451___mcc_h6_
            d_2453_v79_: _dafny.Array
            nw427_ = _dafny.Array(False, 9)
            d_2453_v79_ = nw427_
            d_2453_v79_ = d_2453_v79_
            d_2454_v80_: _dafny.Map
            d_2454_v80_ = _dafny.Map({self.f28: p0})
            d_2455_v81_: str
            d_2455_v81_ = _dafny.CodePoint('f')
            d_2456_v82_: _dafny.Map
            d_2456_v82_ = _dafny.Map({p0: p0})
            d_2457_v83_: _dafny.Seq
            d_2457_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kutxyxep"))
            d_2458_v84_: _dafny.Set
            d_2458_v84_ = _dafny.Set({d_2457_v83_})
            (globalState).f24 = ((0) - (default__.safeModuloInt(p0, ((d_2454_v80_)[self.f28] if (self.f28) in (d_2454_v80_) else (default__.fm57(d_2455_v81_, len(d_2456_v82_), p0, globalState)).cardinality))) if (self).f29 else len(d_2458_v84_))
            d_2459_v85_: _dafny.Set
            d_2459_v85_ = _dafny.Set({self.f28, ((self).f29 if False else (self).f29), self.f28, (True) == (self.f28)})
            (globalState).f26 = d_2459_v85_
            d_2460_v86_: _dafny.Seq
            d_2460_v86_ = _dafny.SeqWithoutIsStrInference([d_2459_v85_])
            (self).f28 = not(((_dafny.Set({not(self.f28)})).intersection(d_2459_v85_)).ispropersubset((d_2460_v86_)[default__.safeIndex(p0, len(d_2460_v86_))]))
        elif True:
            d_2461___mcc_h7_ = source39_.cf59
            d_2462_cf59_ = d_2461___mcc_h7_
            d_2463_v87_: _dafny.Set
            d_2463_v87_ = _dafny.Set({p0})
            if (d_2463_v87_).issubset((d_2463_v87_).intersection(d_2463_v87_)):
                d_2464_v88_: _dafny.Seq
                d_2464_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alu"))
                d_2465_v89_: _dafny.Seq
                d_2465_v89_ = _dafny.SeqWithoutIsStrInference([d_2464_v88_, d_2464_v88_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vooct"))])
                (globalState).f24 = (p0 if (self).fm11(self.f28, d_2465_v89_, (0) - (p0), globalState) else (p0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2466_i7_ in range(default__.abs(99))]))))
                d_2467_v90_: _dafny.Map
                d_2467_v90_ = _dafny.Map({(self).f29: p0})
                (globalState).f1 = ((p0) + (len(d_2467_v90_)) if self.f28 else p0)
                d_2468_v91_: _dafny.Array
                nw428_ = _dafny.Array(int(0), 25)
                d_2468_v91_ = nw428_
                d_2468_v91_ = d_2468_v91_
                (globalState).f27 = self.f28
                d_2469_v92_: _dafny.Map
                d_2469_v92_ = _dafny.Map({p0: self.f28})
                d_2470_v93_: _dafny.Map
                d_2470_v93_ = _dafny.Map({917: len(d_2469_v92_)})
                (globalState).f1 = len(((d_2470_v93_) | (d_2470_v93_)) | (d_2470_v93_))
            elif True:
                d_2471_v94_: C7
                nw429_ = C7()
                nw429_.ctor__(not(default__.fm26(p0, globalState)), (self).f29)
                d_2471_v94_ = nw429_
                d_2472_v95_: _dafny.Map
                d_2472_v95_ = _dafny.Map({(self).f29: False})
                default__.m0(d_2472_v95_, len(d_2394_v43_), self.f28, globalState)
                d_2473_v96_: _dafny.Seq
                d_2473_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tl"))
                d_2474_v97_: _dafny.Map
                d_2474_v97_ = _dafny.Map({self.f28: d_2473_v96_})
                d_2475_v98_: _dafny.Array
                nw430_ = _dafny.Array(None, 13)
                nw430_[int(0)] = (d_2473_v96_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2476_i8_ in range(default__.abs(567))]))
                nw430_[int(1)] = d_2473_v96_
                nw430_[int(2)] = d_2473_v96_
                nw430_[int(3)] = d_2473_v96_
                nw430_[int(4)] = (d_2473_v96_) + (d_2473_v96_)
                nw430_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqq"))
                nw430_[int(6)] = default__.fm51(p0, 156, globalState)
                nw430_[int(7)] = d_2473_v96_
                nw430_[int(8)] = d_2473_v96_
                nw430_[int(9)] = d_2473_v96_
                nw430_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uacmegsp"))
                nw430_[int(11)] = ((d_2474_v97_)[self.f28] if (self.f28) in (d_2474_v97_) else d_2473_v96_)
                nw430_[int(12)] = d_2473_v96_
                d_2475_v98_ = nw430_
                index386_ = default__.safeIndex(500, (d_2475_v98_).length(0))
                (d_2475_v98_)[index386_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "du"))
                index387_ = default__.safeIndex(500, (d_2475_v98_).length(0))
                (d_2475_v98_)[index387_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lu"))
                d_2477_v99_: _dafny.Array
                def lambda114_(d_2478_v95_):
                    def lambda115_(d_2479_i9_):
                        return d_2478_v95_

                    return lambda115_

                init67_ = lambda114_(d_2472_v95_)
                nw431_ = _dafny.Array(None, 22)
                for i0_67_ in range(nw431_.length(0)):
                    nw431_[i0_67_] = init67_(i0_67_)
                d_2477_v99_ = nw431_
                index388_ = default__.safeIndex(738, (d_2477_v99_).length(0))
                (d_2477_v99_)[index388_] = d_2472_v95_
                d_2480_v100_: _dafny.Set
                d_2480_v100_ = _dafny.Set({(self).f29})
                index389_ = default__.safeIndex(738, (d_2477_v99_).length(0))
                (d_2477_v99_)[index389_] = _dafny.Map({self.f28: (d_2480_v100_).isdisjoint(d_2480_v100_)})
                d_2481_v101_: _dafny.Array
                nw432_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_2481_v101_ = nw432_
                nw433_ = _dafny.Array(_dafny.MultiSet({}), 4)
                d_2481_v101_ = nw433_
            d_2482_v102_: _dafny.Array
            nw434_ = _dafny.Array(False, 5)
            d_2482_v102_ = nw434_
            d_2483_v103_: _dafny.Array
            nw435_ = _dafny.Array(None, 16)
            nw435_[int(0)] = d_2482_v102_
            nw435_[int(1)] = d_2482_v102_
            nw435_[int(2)] = d_2482_v102_
            nw435_[int(3)] = d_2482_v102_
            nw435_[int(4)] = d_2482_v102_
            nw435_[int(5)] = d_2482_v102_
            nw435_[int(6)] = d_2482_v102_
            nw435_[int(7)] = d_2482_v102_
            nw435_[int(8)] = d_2482_v102_
            nw435_[int(9)] = d_2482_v102_
            nw435_[int(10)] = d_2482_v102_
            nw435_[int(11)] = d_2482_v102_
            nw435_[int(12)] = d_2482_v102_
            nw435_[int(13)] = d_2482_v102_
            nw435_[int(14)] = d_2482_v102_
            nw435_[int(15)] = d_2482_v102_
            d_2483_v103_ = nw435_
            d_2484_v104_: _dafny.Seq
            d_2484_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqpf"))
            d_2485_v105_: _dafny.Seq
            d_2485_v105_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2486_i10_ in range(default__.abs(539))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2487_i11_ in range(default__.abs(145))]), d_2484_v104_])
            d_2488_v106_: C6
            nw436_ = C6()
            nw436_.ctor__(d_2483_v103_, d_2485_v105_, (self).f29, self.f28)
            d_2488_v106_ = nw436_
            d_2489_v107_: _dafny.Map
            d_2489_v107_ = _dafny.Map({d_2488_v106_: (self).f29})
            d_2490_v108_: _dafny.MultiSet
            d_2490_v108_ = _dafny.MultiSet([p0])
            (globalState).f1 = default__.safeDivisionInt(len(d_2489_v107_), (0) - (((d_2490_v108_)[p0] if (p0) in (d_2490_v108_) else p0)))
            d_2491_v109_: str
            d_2491_v109_ = _dafny.CodePoint('q')
            d_2492_v110_: _dafny.Seq
            d_2492_v110_ = _dafny.SeqWithoutIsStrInference([default__.fm27(p0, d_2491_v109_, p0, globalState)])
            def iife163_():
                coll93_ = _dafny.Set()
                compr_93_: int
                for compr_93_ in _dafny.IntegerRange(64, 456):
                    d_2493_v111_: int = compr_93_
                    if ((64) <= (d_2493_v111_)) and ((d_2493_v111_) < (456)):
                        coll93_ = coll93_.union(_dafny.Set([(d_2493_v111_) + (p0)]))
                return _dafny.Set(coll93_)
            (globalState).f24 = (len(d_2492_v110_)) * (((d_2490_v108_)[p0] if (p0) in (d_2490_v108_) else len(iife163_()
            )))
            d_2494_v112_: _dafny.Set
            d_2494_v112_ = _dafny.Set({self.f28, (self).f29})
            d_2495_v113_: _dafny.Seq
            d_2495_v113_ = _dafny.SeqWithoutIsStrInference([d_2494_v112_, d_2494_v112_, d_2494_v112_, d_2494_v112_, d_2494_v112_])
            d_2496_v114_: _dafny.Map
            d_2496_v114_ = _dafny.Map({p0: (d_2495_v113_)[default__.safeIndex(len(d_2492_v110_), len(d_2495_v113_))]})
            d_2497_v115_: _dafny.Map
            d_2497_v115_ = _dafny.Map({(d_2496_v114_) | (_dafny.Map({p0: d_2494_v112_})): p0})
            d_2497_v115_ = (d_2497_v115_).set(d_2496_v114_, p0)
        d_2498_v116_: _dafny.Map
        d_2498_v116_ = _dafny.Map({493: (self).f29})
        d_2499_v117_: T2
        nw437_ = C3()
        nw437_.ctor__((self).f29, self.f28, (self).f29, p0)
        d_2499_v117_ = nw437_
        d_2500_v118_: _dafny.Seq
        d_2500_v118_ = _dafny.SeqWithoutIsStrInference([d_2499_v117_, d_2499_v117_, d_2499_v117_, d_2499_v117_])
        d_2498_v116_ = (d_2498_v116_).set(len(d_2500_v118_), (self).f29)
        d_2501_i12_: int
        d_2501_i12_ = 0
        with _dafny.label("6"):
            while default__.fm18(p0, (d_2499_v117_).f39, p0, globalState):
                with _dafny.c_label("6"):
                    if (d_2501_i12_) >= (100):
                        raise _dafny.Break("6")
                    d_2501_i12_ = (d_2501_i12_) + (1)
                    d_2502_v119_: _dafny.Seq
                    d_2502_v119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl"))
                    d_2503_v120_: _dafny.Seq
                    d_2503_v120_ = _dafny.SeqWithoutIsStrInference([d_2502_v119_, default__.fm51(p0, 956, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxmjiy"))])
                    d_2504_v121_: _dafny.MultiSet
                    d_2504_v121_ = _dafny.MultiSet([d_2502_v119_, d_2502_v119_, d_2502_v119_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wipp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iahq"))])
                    d_2505_v122_: _dafny.Array
                    nw438_ = _dafny.Array(None, 12)
                    nw438_[int(0)] = (self).f29
                    nw438_[int(1)] = self.f28
                    nw438_[int(2)] = default__.fm38(globalState)
                    nw438_[int(3)] = (self).fm11((self).f29, d_2503_v120_, 74, globalState)
                    nw438_[int(4)] = self.f28
                    nw438_[int(5)] = (default__.fm20(self.f28, globalState)).ispropersubset(d_2504_v121_)
                    nw438_[int(6)] = (self).f29
                    nw438_[int(7)] = not(((self).f29) or ((self).f29))
                    nw438_[int(8)] = ((self).f29) and ((self).f29)
                    nw438_[int(9)] = not(((self).f29 if True else True))
                    nw438_[int(10)] = self.f28
                    nw438_[int(11)] = not (self.f28) or (self.f28)
                    d_2505_v122_ = nw438_
                    d_2505_v122_ = d_2505_v122_
                    d_2506_v123_: _dafny.Seq
                    d_2506_v123_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f28]), d_2394_v43_])
                    (globalState).f24 = len((_dafny.SeqWithoutIsStrInference([(self).f29, self.f28])) + ((d_2394_v43_ if default__.fm29(_dafny.MultiSet(d_2506_v123_), 202, globalState) else d_2394_v43_)))
                    d_2505_v122_ = d_2505_v122_
                    (globalState).f13 = self.f28
                    pass
            pass
        if (D7_DC17(False)).cf22:
            d_2507_v124_: _dafny.Array
            nw439_ = _dafny.Array(False, 28)
            d_2507_v124_ = nw439_
            d_2508_v125_: _dafny.Map
            d_2508_v125_ = _dafny.Map({self.f28: d_2507_v124_})
            rhs414_ = (self).f29
            rhs415_ = -703
            rhs416_ = ((_dafny.Map({self.f28: d_2507_v124_})) | (d_2508_v125_)) | ((d_2508_v125_) | (_dafny.Map({default__.fm13(globalState): d_2507_v124_})))
            rhs417_ = self.f28
            lhs350_ = globalState
            lhs351_ = globalState
            r0 = rhs414_
            lhs350_.f1 = rhs415_
            d_2508_v125_ = rhs416_
            lhs351_.f22 = rhs417_
            d_2509_v126_: _dafny.Seq
            d_2509_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lce"))
            d_2510_v127_: _dafny.Map
            d_2510_v127_ = _dafny.Map({(self).f29: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tobgi"))})
            d_2509_v126_ = (((d_2510_v127_)[(self).f29] if ((self).f29) in (d_2510_v127_) else d_2509_v126_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ondgjo")))
            d_2511_v128_: D10
            d_2511_v128_ = D10_DC27((self).f29, (self).f29, True, (d_2499_v117_).f39, True)
            d_2512_v129_: _dafny.Map
            d_2512_v129_ = _dafny.Map({d_2509_v126_: d_2511_v128_})
            d_2513_v130_: D10
            d_2513_v130_ = D10_DC29(((d_2512_v129_)[d_2509_v126_] if (d_2509_v126_) in (d_2512_v129_) else D10_DC27(True, self.f28, (self).f29, p0, (self).f29)))
            d_2514_v131_: _dafny.Seq
            d_2514_v131_ = _dafny.SeqWithoutIsStrInference([d_2513_v130_, D10_DC29(d_2511_v128_)])
            d_2515_v132_: _dafny.MultiSet
            d_2515_v132_ = _dafny.MultiSet([(d_2499_v117_).f39])
            pat_let_tv62_ = d_2511_v128_
            d_2516_v133_: _dafny.Array
            nw440_ = _dafny.Array(None, 28)
            nw440_[int(0)] = d_2514_v131_
            nw440_[int(1)] = d_2514_v131_
            nw440_[int(2)] = d_2514_v131_
            nw440_[int(3)] = d_2514_v131_
            nw440_[int(4)] = d_2514_v131_
            nw440_[int(5)] = default__.fm70(p0, globalState)
            nw440_[int(6)] = d_2514_v131_
            def iife164_(_pat_let35_0):
                def iife165_(d_2517_dt__update__tmp_h1_):
                    def iife166_(_pat_let36_0):
                        def iife167_(d_2518_dt__update_hcf46_h0_):
                            return D10_DC29(d_2518_dt__update_hcf46_h0_)
                        return iife167_(_pat_let36_0)
                    return iife166_(pat_let_tv62_)
                return iife165_(_pat_let35_0)
            nw440_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_2513_v130_, d_2513_v130_, d_2513_v130_, d_2513_v130_, iife164_(d_2513_v130_)])) + (d_2514_v131_)
            nw440_[int(8)] = d_2514_v131_
            nw440_[int(9)] = d_2514_v131_
            nw440_[int(10)] = _dafny.SeqWithoutIsStrInference([d_2513_v130_ for d_2519_i13_ in range(default__.abs(910))])
            nw440_[int(11)] = d_2514_v131_
            nw440_[int(12)] = (d_2514_v131_) + ((D20_DC50(d_2514_v131_)).cf74)
            nw440_[int(13)] = d_2514_v131_
            nw440_[int(14)] = d_2514_v131_
            nw440_[int(15)] = default__.fm70((d_2515_v132_).cardinality, globalState)
            nw440_[int(16)] = d_2514_v131_
            nw440_[int(17)] = _dafny.SeqWithoutIsStrInference([d_2513_v130_])
            nw440_[int(18)] = (d_2514_v131_ if (self).f29 else _dafny.SeqWithoutIsStrInference([d_2513_v130_, d_2513_v130_]))
            nw440_[int(19)] = d_2514_v131_
            nw440_[int(20)] = (d_2514_v131_) + (d_2514_v131_)
            nw440_[int(21)] = d_2514_v131_
            nw440_[int(22)] = (d_2514_v131_) + (d_2514_v131_)
            nw440_[int(23)] = d_2514_v131_
            nw440_[int(24)] = d_2514_v131_
            nw440_[int(25)] = _dafny.SeqWithoutIsStrInference([d_2513_v130_ for d_2520_i14_ in range(default__.abs(53))])
            nw440_[int(26)] = (default__.fm70(269, globalState) if True else d_2514_v131_)
            nw440_[int(27)] = d_2514_v131_
            d_2516_v133_ = nw440_
            d_2516_v133_ = d_2516_v133_
            d_2521_v134_: _dafny.MultiSet
            d_2521_v134_ = _dafny.MultiSet([(self).f29])
            d_2522_v135_: _dafny.Map
            d_2522_v135_ = _dafny.Map({d_2507_v124_: (d_2521_v134_).ispropersubset((d_2521_v134_).set((self).f29, default__.abs((d_2499_v117_).f39)))})
            d_2522_v135_ = d_2522_v135_
            d_2523_v136_: str
            d_2523_v136_ = _dafny.CodePoint('o')
            (globalState).f24 = default__.fm25(p0, default__.fm27((d_2499_v117_).f39, d_2523_v136_, len(d_2509_v126_), globalState), globalState)
        elif True:
            if (self).f29:
                (globalState).f24 = p0
                d_2524_v137_: _dafny.Map
                d_2524_v137_ = _dafny.Map({(self).f29: (self).f29})
                d_2525_v138_: _dafny.Set
                d_2525_v138_ = _dafny.Set({False, (self).f29})
                default__.m0((d_2524_v137_) | (default__.fm59((self).f29, _dafny.SeqWithoutIsStrInference([default__.fm38(globalState)]), globalState)), len(d_2525_v138_), (self).f29, globalState)
                (globalState).f24 = (d_2499_v117_).f39
                d_2526_v139_: D10
                d_2526_v139_ = D10_DC28(d_2524_v137_)
                d_2527_v140_: _dafny.Map
                d_2527_v140_ = _dafny.Map({_dafny.CodePoint('y'): d_2526_v139_})
                (globalState).f1 = len((default__.fm71((d_2499_v117_).f39, globalState)) | (d_2527_v140_))
                d_2528_v141_: C8
                nw441_ = C8()
                nw441_.ctor__()
                d_2528_v141_ = nw441_
            elif True:
                d_2529_v142_: _dafny.Array
                nw442_ = _dafny.Array(None, 3)
                d_2529_v142_ = nw442_
                d_2530_v143_: D21
                d_2530_v143_ = D21_DC54(d_2529_v142_)
                d_2531_v144_: _dafny.MultiSet
                d_2531_v144_ = _dafny.MultiSet([d_2529_v142_, ((d_2530_v143_).cf79 if (self).f29 else d_2529_v142_), d_2529_v142_, d_2529_v142_, d_2529_v142_])
                d_2531_v144_ = d_2531_v144_
                d_2532_v145_: _dafny.Map
                d_2532_v145_ = _dafny.Map({self.f28: (d_2499_v117_).f39})
                d_2532_v145_ = (d_2532_v145_).set((self).f29, (d_2499_v117_).f39)
                d_2533_v146_: _dafny.Array
                def lambda116_(d_2534_i15_):
                    return self.f28

                init68_ = lambda116_
                nw443_ = _dafny.Array(None, 11)
                for i0_68_ in range(nw443_.length(0)):
                    nw443_[i0_68_] = init68_(i0_68_)
                d_2533_v146_ = nw443_
                d_2535_v147_: _dafny.Seq
                d_2535_v147_ = _dafny.SeqWithoutIsStrInference([d_2533_v146_])
                d_2536_v148_: C4
                nw444_ = C4()
                nw444_.ctor__(self.f28, (len(d_2535_v147_)) <= ((d_2499_v117_).f39), not (not(self.f28)) or (self.f28), False)
                d_2536_v148_ = nw444_
                (globalState).f27 = (self).f29
                d_2537_v149_: _dafny.Map
                d_2537_v149_ = _dafny.Map({self.f28: (d_2536_v148_).f37})
                d_2538_v150_: D17
                d_2538_v150_ = D17_DC45(d_2537_v149_, self.f28)
                rhs418_ = d_2538_v150_
                rhs419_ = (d_2538_v150_).cf65
                lhs352_ = globalState
                d_2538_v150_ = rhs418_
                lhs352_.f27 = rhs419_
            (globalState).f1 = (d_2499_v117_).f39
            d_2539_v151_: _dafny.Seq
            d_2539_v151_ = _dafny.SeqWithoutIsStrInference([(d_2499_v117_).f39, (0) - (p0)])
            d_2540_v152_: D4
            d_2540_v152_ = D4_DC10(not(self.f28), (d_2499_v117_).f39, len(d_2539_v151_))
            (globalState).f24 = default__.safeModuloInt((d_2540_v152_).cf14, (842) * ((d_2499_v117_).f39))
            d_2541_v153_: str
            d_2541_v153_ = _dafny.CodePoint('v')
            (globalState).f24 = (0) - (default__.fm27((d_2499_v117_).f39, d_2541_v153_, (d_2499_v117_).f39, globalState))
            (globalState).f24 = (719) - (p0)
        (globalState).f21 = self.f28
        r0 = (self).f29
        return r0


class C10:
    def  __init__(self):
        self._f31: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f32: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self, f31, f32):
        (self)._f31 = f31
        (self)._f32 = f32

    def m7(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: bool = False
        r3: bool = False
        d_2542_v0_: _dafny.Array
        nw445_ = _dafny.Array(_dafny.MultiSet({}), 17)
        d_2542_v0_ = nw445_
        d_2543_v1_: _dafny.Seq
        d_2543_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        d_2544_v2_: _dafny.MultiSet
        d_2544_v2_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p2: 316})) for d_2545_i0_ in range(default__.abs(819))])), (self).f32, len(d_2543_v1_)])])
        index390_ = default__.safeIndex(587, (d_2542_v0_).length(0))
        (d_2542_v0_)[index390_] = d_2544_v2_
        index391_ = default__.safeIndex(587, (d_2542_v0_).length(0))
        (d_2542_v0_)[index391_] = d_2544_v2_
        if p2:
            rhs420_ = _dafny.SeqWithoutIsStrInference([default__.fm8(globalState), p0, (self).f32])
            rhs421_ = p2
            r0 = rhs420_
            r2 = rhs421_
            d_2546_v3_: _dafny.Array
            nw446_ = _dafny.Array(False, 18)
            d_2546_v3_ = nw446_
            d_2547_v4_: _dafny.Seq
            d_2547_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysqwm"))
            d_2548_v5_: D2
            d_2548_v5_ = D2_DC5(p2)
            d_2549_v6_: _dafny.Map
            d_2549_v6_ = _dafny.Map({d_2548_v5_: d_2546_v3_})
            d_2550_v7_: str
            d_2550_v7_ = _dafny.CodePoint('w')
            d_2551_v8_: _dafny.Set
            d_2551_v8_ = _dafny.Set({p2, not(not(p2)), p2})
            d_2552_v9_: _dafny.Set
            d_2552_v9_ = _dafny.Set({(self).f32, p0})
            d_2553_v10_: _dafny.Map
            d_2553_v10_ = _dafny.Map({d_2551_v8_: d_2552_v9_})
            d_2554_v11_: _dafny.Map
            d_2554_v11_ = _dafny.Map({len((d_2553_v10_).set(d_2551_v8_, d_2552_v9_)): (self).f32})
            d_2555_v12_: _dafny.Map
            d_2555_v12_ = _dafny.Map({p2: p1})
            d_2556_v13_: D1
            d_2556_v13_ = D1_DC2(len(d_2555_v12_), -574, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2557_i1_ in range(default__.abs(-692))]), p2)
            d_2558_v14_: D1
            d_2558_v14_ = D1_DC2((d_2556_v13_).cf2, len(d_2543_v1_), (self).f31, p2)
            rhs422_ = ((d_2549_v6_)[d_2548_v5_] if (d_2548_v5_) in (d_2549_v6_) else d_2546_v3_)
            rhs423_ = ((438) - (p1)) - ((self).f32)
            rhs424_ = (((self).f31) + ((d_2547_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yd"))))).set(default__.safeIndex((self).f32, len(((self).f31) + ((d_2547_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yd")))))), d_2550_v7_)
            rhs425_ = ((d_2554_v11_)[p1] if (p1) in (d_2554_v11_) else (self).f32)
            rhs426_ = (len(d_2547_v4_)) - ((d_2556_v13_).cf3)
            lhs353_ = globalState
            lhs354_ = globalState
            lhs355_ = globalState
            d_2546_v3_ = rhs422_
            lhs353_.f24 = rhs423_
            d_2547_v4_ = rhs424_
            lhs354_.f24 = rhs425_
            lhs355_.f24 = rhs426_
            (globalState).f1 = p0
            (globalState).f22 = (p0) != ((0) - (p0))
            d_2559_v15_: _dafny.MultiSet
            d_2559_v15_ = _dafny.MultiSet([p0, p1, 247, p0])
            d_2560_v16_: _dafny.Map
            d_2560_v16_ = _dafny.Map({(d_2559_v15_).cardinality: p2})
            d_2561_v17_: _dafny.Seq
            d_2561_v17_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2562_v18_: _dafny.Seq
            d_2562_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0, p1, p0, p0})])
            d_2563_v19_: _dafny.Array
            nw447_ = _dafny.Array(None, 22)
            nw447_[int(0)] = (len(d_2560_v16_) if p2 else p1)
            nw447_[int(1)] = len((d_2560_v16_) | (_dafny.Map({(0) - (p1): p2})))
            nw447_[int(2)] = (self).f32
            nw447_[int(3)] = (180) * (default__.fm8(globalState))
            nw447_[int(4)] = p0
            nw447_[int(5)] = p0
            nw447_[int(6)] = (0) - ((self).f32)
            nw447_[int(7)] = len(d_2547_v4_)
            nw447_[int(8)] = default__.safeModuloInt(len(_dafny.Map({p2: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "raidwy"))))})), p1)
            nw447_[int(9)] = 907
            nw447_[int(10)] = (len(d_2561_v17_)) * (p0)
            nw447_[int(11)] = 670
            nw447_[int(12)] = p0
            nw447_[int(13)] = len((d_2562_v18_) + (_dafny.SeqWithoutIsStrInference([d_2552_v9_, d_2552_v9_, d_2552_v9_])))
            nw447_[int(14)] = 169
            nw447_[int(15)] = default__.fm8(globalState)
            nw447_[int(16)] = p0
            nw447_[int(17)] = p0
            nw447_[int(18)] = len(_dafny.SeqWithoutIsStrInference([False, False, p2]))
            nw447_[int(19)] = 821
            nw447_[int(20)] = p1
            nw447_[int(21)] = (_dafny.MultiSet(d_2561_v17_)).cardinality
            d_2563_v19_ = nw447_
            d_2564_v20_: _dafny.MultiSet
            d_2564_v20_ = _dafny.MultiSet([p2])
            d_2565_v21_: _dafny.Map
            d_2565_v21_ = _dafny.Map({_dafny.Set({(d_2564_v20_).cardinality}): p1})
            index392_ = default__.safeIndex(662, (d_2563_v19_).length(0))
            (d_2563_v19_)[index392_] = len((_dafny.Map({default__.fm9(default__.fm10(p2, (self).f32, p2, globalState), globalState): (self).f32})) | (d_2565_v21_))
            index393_ = default__.safeIndex(662, (d_2563_v19_).length(0))
            rhs427_ = (default__.fm10(p2, default__.fm8(globalState), p2, globalState)) == (p2)
            rhs428_ = p1
            lhs356_ = globalState
            lhs357_ = d_2563_v19_
            lhs358_ = default__.safeIndex(662, (d_2563_v19_).length(0))
            lhs356_.f13 = rhs427_
            lhs357_[lhs358_] = rhs428_
        elif True:
            d_2566_v22_: C8
            nw448_ = C8()
            nw448_.ctor__()
            d_2566_v22_ = nw448_
            d_2567_v23_: _dafny.Array
            def lambda117_(d_2568_p2_):
                def lambda118_(d_2569_i2_):
                    return not(d_2568_p2_)

                return lambda118_

            init69_ = lambda117_(p2)
            nw449_ = _dafny.Array(None, 24)
            for i0_69_ in range(nw449_.length(0)):
                nw449_[i0_69_] = init69_(i0_69_)
            d_2567_v23_ = nw449_
            index394_ = default__.safeIndex(536, (d_2567_v23_).length(0))
            (d_2567_v23_)[index394_] = p2
            index395_ = default__.safeIndex(536, (d_2567_v23_).length(0))
            (d_2567_v23_)[index395_] = (((self).f31) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khxjcw")))) < ((self).f31)
            d_2570_v24_: _dafny.Array
            nw450_ = _dafny.Array(int(0), 3)
            d_2570_v24_ = nw450_
            d_2570_v24_ = d_2570_v24_
            d_2571_v25_: _dafny.Array
            nw451_ = _dafny.Array(None, 7)
            nw451_[int(0)] = d_2570_v24_
            nw451_[int(1)] = d_2570_v24_
            nw451_[int(2)] = d_2570_v24_
            nw451_[int(3)] = d_2570_v24_
            nw451_[int(4)] = d_2570_v24_
            nw451_[int(5)] = (d_2570_v24_ if (d_2567_v23_)[default__.safeIndex(536, (d_2567_v23_).length(0))] else d_2570_v24_)
            nw451_[int(6)] = d_2570_v24_
            d_2571_v25_ = nw451_
            index396_ = default__.safeIndex(318, (d_2570_v24_).length(0))
            (d_2570_v24_)[index396_] = len((self).f31)
            d_2572_v26_: D7
            d_2572_v26_ = D7_DC18((p0) >= (len((self).f31)))
            d_2573_v27_: _dafny.Array
            def lambda119_(d_2574_i3_):
                return D12_DC31(_dafny.Set({(self).f32}))

            init70_ = lambda119_
            nw452_ = _dafny.Array(None, 10)
            for i0_70_ in range(nw452_.length(0)):
                nw452_[i0_70_] = init70_(i0_70_)
            d_2573_v27_ = nw452_
            d_2575_v28_: _dafny.Set
            d_2575_v28_ = _dafny.Set({125, 703})
            d_2576_v29_: D12
            d_2576_v29_ = D12_DC31(d_2575_v28_)
            index397_ = default__.safeIndex(175, (d_2573_v27_).length(0))
            (d_2573_v27_)[index397_] = d_2576_v29_
            d_2577_v30_: str
            d_2577_v30_ = _dafny.CodePoint('b')
            index398_ = default__.safeIndex(318, (d_2570_v24_).length(0))
            index399_ = default__.safeIndex(175, (d_2573_v27_).length(0))
            rhs429_ = d_2571_v25_
            rhs430_ = len((((self).f31) + ((self).f31) if not (False) or ((d_2567_v23_)[default__.safeIndex(536, (d_2567_v23_).length(0))]) else (((self).f31).set(default__.safeIndex(p1, len((self).f31)), d_2577_v30_)) + ((self).f31)))
            rhs431_ = default__.safeDivisionInt(-806, p0)
            rhs432_ = (D7_DC18(p2) if (d_2567_v23_)[default__.safeIndex(536, (d_2567_v23_).length(0))] else d_2572_v26_)
            rhs433_ = d_2576_v29_
            lhs359_ = globalState
            lhs360_ = d_2570_v24_
            lhs361_ = default__.safeIndex(318, (d_2570_v24_).length(0))
            lhs362_ = d_2573_v27_
            lhs363_ = default__.safeIndex(175, (d_2573_v27_).length(0))
            d_2571_v25_ = rhs429_
            lhs359_.f24 = rhs430_
            lhs360_[lhs361_] = rhs431_
            d_2572_v26_ = rhs432_
            lhs362_[lhs363_] = rhs433_
            (globalState).f1 = default__.safeModuloInt(default__.safeModuloInt(p1, (d_2570_v24_)[default__.safeIndex(318, (d_2570_v24_).length(0))]), p1)
        d_2543_v1_ = (_dafny.SeqWithoutIsStrInference([p2, p2])) + (d_2543_v1_)
        d_2578_i4_: int
        d_2578_i4_ = 0
        with _dafny.label("7"):
            while p2:
                with _dafny.c_label("7"):
                    if (d_2578_i4_) >= (100):
                        raise _dafny.Break("7")
                    d_2578_i4_ = (d_2578_i4_) + (1)
                    d_2579_v31_: C9
                    nw453_ = C9()
                    nw453_.ctor__(not(((self).f32) <= ((0) - (len(default__.fm47(globalState))))), not(p2))
                    d_2579_v31_ = nw453_
                    d_2580_v32_: _dafny.Map
                    d_2580_v32_ = _dafny.Map({(self).f32: _dafny.MultiSet([p1, len((self).f31)])})
                    d_2581_v33_: D13
                    d_2581_v33_ = D13_DC34(d_2580_v32_)
                    source41_ = d_2581_v33_
                    if source41_.is_DC35:
                        d_2582_v34_: _dafny.MultiSet
                        d_2582_v34_ = _dafny.MultiSet([p0])
                        d_2583_v35_: _dafny.Array
                        nw454_ = _dafny.Array(None, 3)
                        nw454_[int(0)] = len((self).f31)
                        nw454_[int(1)] = len(default__.fm51(181, p1, globalState))
                        nw454_[int(2)] = ((d_2582_v34_)[p1] if (p1) in (d_2582_v34_) else p1)
                        d_2583_v35_ = nw454_
                        d_2584_v36_: _dafny.Seq
                        d_2584_v36_ = _dafny.SeqWithoutIsStrInference([(self).f32])
                        d_2585_v37_: _dafny.Map
                        d_2585_v37_ = _dafny.Map({p2: (self).f32})
                        index400_ = default__.safeIndex(17, (d_2583_v35_).length(0))
                        (d_2583_v35_)[index400_] = len(_dafny.Set({(d_2584_v36_)[default__.safeIndex(p1, len(d_2584_v36_))], len(default__.fm72(_dafny.SeqWithoutIsStrInference([d_2585_v37_, d_2585_v37_]), globalState))}))
                        index401_ = default__.safeIndex(17, (d_2583_v35_).length(0))
                        (d_2583_v35_)[index401_] = (self).f32
                        index402_ = default__.safeIndex(17, (d_2583_v35_).length(0))
                        (d_2583_v35_)[index402_] = 321
                        d_2586_v38_: _dafny.Map
                        d_2586_v38_ = _dafny.Map({p2: p2})
                        d_2587_v39_: _dafny.Map
                        d_2587_v39_ = _dafny.Map({(0) - ((d_2582_v34_).cardinality): d_2586_v38_})
                        d_2588_v40_: _dafny.Map
                        d_2588_v40_ = _dafny.Map({(d_2583_v35_)[default__.safeIndex(17, (d_2583_v35_).length(0))]: p0})
                        default__.m0(((d_2587_v39_)[default__.fm27((self).f32, default__.fm42((d_2583_v35_)[default__.safeIndex(17, (d_2583_v35_).length(0))], (self).f32, globalState), p0, globalState)] if (default__.fm27((self).f32, default__.fm42((d_2583_v35_)[default__.safeIndex(17, (d_2583_v35_).length(0))], (self).f32, globalState), p0, globalState)) in (d_2587_v39_) else d_2586_v38_), (0) - (((d_2588_v40_)[(self).f32] if ((self).f32) in (d_2588_v40_) else p1)), p2, globalState)
                        d_2589_v41_: _dafny.Map
                        d_2589_v41_ = _dafny.Map({571: p2})
                        (globalState).f27 = (d_2589_v41_) != (d_2589_v41_)
                    elif True:
                        d_2590___mcc_h0_ = source41_.cf50
                        d_2591_cf50_ = d_2590___mcc_h0_
                        (globalState).f1 = default__.safeModuloInt(p0, 383)
                        d_2592_v42_: _dafny.Array
                        nw455_ = _dafny.Array(None, 15)
                        nw455_[int(0)] = p2
                        nw455_[int(1)] = p2
                        nw455_[int(2)] = p2
                        nw455_[int(3)] = p2
                        nw455_[int(4)] = p2
                        nw455_[int(5)] = p2
                        nw455_[int(6)] = p2
                        nw455_[int(7)] = p2
                        nw455_[int(8)] = False
                        nw455_[int(9)] = p2
                        nw455_[int(10)] = p2
                        nw455_[int(11)] = p2
                        nw455_[int(12)] = p2
                        nw455_[int(13)] = p2
                        nw455_[int(14)] = p2
                        d_2592_v42_ = nw455_
                        d_2593_v43_: _dafny.Seq
                        d_2593_v43_ = _dafny.SeqWithoutIsStrInference([d_2592_v42_, d_2592_v42_])
                        d_2594_v44_: _dafny.Array
                        nw456_ = _dafny.Array(None, 23)
                        nw456_[int(0)] = d_2592_v42_
                        nw456_[int(1)] = d_2592_v42_
                        nw456_[int(2)] = d_2592_v42_
                        nw456_[int(3)] = d_2592_v42_
                        nw456_[int(4)] = d_2592_v42_
                        nw456_[int(5)] = d_2592_v42_
                        nw456_[int(6)] = d_2592_v42_
                        nw456_[int(7)] = d_2592_v42_
                        nw456_[int(8)] = d_2592_v42_
                        nw456_[int(9)] = d_2592_v42_
                        nw456_[int(10)] = d_2592_v42_
                        nw456_[int(11)] = d_2592_v42_
                        nw456_[int(12)] = d_2592_v42_
                        nw456_[int(13)] = d_2592_v42_
                        nw456_[int(14)] = d_2592_v42_
                        nw456_[int(15)] = d_2592_v42_
                        nw456_[int(16)] = d_2592_v42_
                        nw456_[int(17)] = d_2592_v42_
                        nw456_[int(18)] = d_2592_v42_
                        nw456_[int(19)] = d_2592_v42_
                        nw456_[int(20)] = d_2592_v42_
                        nw456_[int(21)] = (d_2593_v43_)[default__.safeIndex(default__.fm27((self).f32, _dafny.CodePoint('k'), (self).f32, globalState), len(d_2593_v43_))]
                        nw456_[int(22)] = d_2592_v42_
                        d_2594_v44_ = nw456_
                        d_2595_v45_: str
                        d_2595_v45_ = _dafny.CodePoint('p')
                        d_2596_v46_: D18
                        d_2596_v46_ = D18_DC48(p2, default__.fm25((self).f32, default__.fm25(p0, (0) - (p0), globalState), globalState), d_2594_v44_, not(not(p2)), d_2595_v45_)
                        (globalState).f21 = (d_2596_v46_).cf71
                        d_2597_v47_: _dafny.Array
                        nw457_ = _dafny.Array(int(0), 4)
                        d_2597_v47_ = nw457_
                        d_2598_v48_: _dafny.Array
                        d_2598_v48_ = d_2597_v47_
                        d_2599_v49_: _dafny.Map
                        d_2599_v49_ = _dafny.Map({p1: d_2598_v48_})
                        d_2599_v49_ = (d_2599_v49_).set(p0, d_2597_v47_)
                        d_2597_v47_ = d_2597_v47_
                    d_2600_v50_: C9
                    nw458_ = C9()
                    nw458_.ctor__(p2, p2)
                    d_2600_v50_ = nw458_
                    d_2601_v51_: _dafny.Map
                    d_2601_v51_ = _dafny.Map({p1: (self).f32})
                    (globalState).f21 = (((d_2601_v51_)[p0] if (p0) in (d_2601_v51_) else default__.fm12(p0, p2, globalState))) < (len(_dafny.Set({True})))
                    pass
            pass
        d_2602_v52_: _dafny.Array
        nw459_ = _dafny.Array(int(0), 2)
        d_2602_v52_ = nw459_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2602_v52_).length(0)):
            d_2603_i5_: int = guard_loop_8_
            if (True) and (((0) <= (d_2603_i5_)) and ((d_2603_i5_) < ((d_2602_v52_).length(0)))):
                (d_2602_v52_)[(d_2603_i5_)] = default__.safeModuloInt(d_2603_i5_, (self).f32)
        d_2604_v53_: str
        d_2604_v53_ = _dafny.CodePoint('e')
        d_2605_v54_: _dafny.Set
        d_2605_v54_ = _dafny.Set({True})
        d_2606_v55_: _dafny.Seq
        d_2606_v55_ = _dafny.SeqWithoutIsStrInference([len(d_2605_v54_), p1])
        d_2607_v56_: _dafny.Seq
        d_2607_v56_ = _dafny.SeqWithoutIsStrInference([(self).f32, (d_2606_v55_)[default__.safeIndex(p1, len(d_2606_v55_))]])
        (globalState).f24 = (default__.fm27(p0, d_2604_v53_, len((d_2606_v55_).set(default__.safeIndex((self).f32, len(d_2606_v55_)), p1)), globalState)) - ((self).f32)
        r0 = d_2607_v56_
        r1 = p2
        r2 = True
        r3 = p2
        return r0, r1, r2, r3

    def m8(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_2608_v0_: T2
        nw460_ = C3()
        nw460_.ctor__(p1, p1, default__.fm13(globalState), (self).f32)
        d_2608_v0_ = nw460_
        d_2609_v1_: _dafny.Seq
        d_2609_v1_ = _dafny.SeqWithoutIsStrInference([d_2608_v0_])
        d_2610_v2_: _dafny.Map
        d_2610_v2_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([d_2608_v0_, d_2608_v0_])})
        (globalState).f1 = len(((d_2609_v1_) + (_dafny.SeqWithoutIsStrInference([d_2608_v0_]))) + (((d_2610_v2_)[p1] if (p1) in (d_2610_v2_) else d_2609_v1_)))
        d_2611_v3_: _dafny.Map
        d_2611_v3_ = _dafny.Map({p1: (d_2608_v0_).f39})
        d_2612_v4_: _dafny.Seq
        d_2612_v4_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2611_v3_ = (d_2611_v3_).set(not(p1), len(d_2612_v4_))
        d_2613_v5_: _dafny.Array
        def lambda120_(d_2614_v0_, d_2615_p0_):
            def lambda121_(d_2616_i0_):
                return (_dafny.MultiSet([(d_2614_v0_).f39, (d_2614_v0_).f39, (self).f32, d_2615_p0_])).intersection(_dafny.MultiSet([357]))

            return lambda121_

        init71_ = lambda120_(d_2608_v0_, p0)
        nw461_ = _dafny.Array(None, 8)
        for i0_71_ in range(nw461_.length(0)):
            nw461_[i0_71_] = init71_(i0_71_)
        d_2613_v5_ = nw461_
        rhs434_ = ((len(default__.fm73(globalState))) - (len(d_2611_v3_))) + (default__.safeModuloInt(default__.fm8(globalState), (d_2608_v0_).f39))
        rhs435_ = (self).f32
        rhs436_ = (p1 if p1 else False)
        rhs437_ = (self).f32
        rhs438_ = d_2613_v5_
        lhs364_ = globalState
        r0 = rhs434_
        lhs364_.f1 = rhs435_
        r1 = rhs436_
        r0 = rhs437_
        d_2613_v5_ = rhs438_
        d_2617_v6_: T0
        nw462_ = C1()
        nw462_.ctor__(p1, p1)
        d_2617_v6_ = nw462_
        d_2618_v7_: D8
        d_2618_v7_ = D8_DC20(d_2617_v6_)
        d_2619_v8_: D8
        d_2619_v8_ = D8_DC20((d_2618_v7_).cf26)
        source42_ = d_2618_v7_
        if source42_.is_DC21:
            d_2620___mcc_h0_ = source42_.cf27
            d_2621_cf27_ = d_2620___mcc_h0_
            d_2622_v9_: _dafny.MultiSet
            d_2622_v9_ = _dafny.MultiSet([(d_2617_v6_).f29, default__.fm26(416, globalState)])
            d_2623_v10_: D9
            d_2623_v10_ = D9_DC23((d_2622_v9_) | (d_2622_v9_))
            d_2624_v11_: _dafny.Seq
            d_2624_v11_ = _dafny.SeqWithoutIsStrInference([d_2622_v9_, d_2622_v9_, d_2622_v9_])
            d_2625_v12_: _dafny.Map
            d_2625_v12_ = _dafny.Map({p0: p0})
            d_2626_v13_: _dafny.MultiSet
            d_2626_v13_ = _dafny.MultiSet([(d_2608_v0_).f39])
            rhs439_ = (self).f32
            rhs440_ = D9_DC23((d_2624_v11_)[default__.safeIndex(len(d_2625_v12_), len(d_2624_v11_))])
            rhs441_ = (d_2608_v0_).f39
            rhs442_ = not((d_2626_v13_) != (d_2626_v13_))
            lhs365_ = globalState
            r0 = rhs439_
            d_2623_v10_ = rhs440_
            r0 = rhs441_
            lhs365_.f13 = rhs442_
            (globalState).f21 = p1
            d_2627_v14_: _dafny.Map
            d_2627_v14_ = _dafny.Map({False: d_2617_v6_.f28})
            d_2627_v14_ = (d_2627_v14_).set((d_2617_v6_).f29, not(not((len((self).f31)) != ((0) - ((d_2608_v0_).f39)))))
            (globalState).f1 = p0
        elif source42_.is_DC22:
            if (d_2617_v6_).f29:
                (globalState).f17 = _dafny.SeqWithoutIsStrInference([((_dafny.MultiSet([(d_2608_v0_).f39])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_2612_v4_)])))).cardinality for d_2628_i1_ in range(default__.abs(444))])
                d_2629_v15_: _dafny.Array
                nw463_ = _dafny.Array(D17.default()(), 3)
                d_2629_v15_ = nw463_
                d_2630_v16_: D22
                d_2630_v16_ = D22_DC58(d_2629_v15_)
                pat_let_tv63_ = d_2629_v15_
                def iife168_(_pat_let37_0):
                    def iife169_(d_2631_dt__update__tmp_h0_):
                        def iife170_(_pat_let38_0):
                            def iife171_(d_2632_dt__update_hcf82_h0_):
                                return D22_DC58(d_2632_dt__update_hcf82_h0_)
                            return iife171_(_pat_let38_0)
                        return iife170_(pat_let_tv63_)
                    return iife169_(_pat_let37_0)
                d_2629_v15_ = (iife168_(d_2630_v16_)).cf82
                (globalState).f1 = (d_2608_v0_).f39
                (globalState).f21 = p1
                d_2633_v17_: C2
                nw464_ = C2()
                nw464_.ctor__()
                d_2633_v17_ = nw464_
            elif True:
                d_2634_v18_: D22
                d_2634_v18_ = D22_DC59((d_2617_v6_).f29, (self).f31, (self).f32)
                d_2635_v19_: D4
                d_2635_v19_ = D4_DC10(d_2617_v6_.f28, (self).f32, (self).f32)
                d_2636_v20_: _dafny.Map
                d_2636_v20_ = _dafny.Map({d_2634_v18_: default__.fm56(d_2635_v19_, p1, (self).f32, globalState)})
                d_2611_v3_ = ((((d_2636_v20_)[d_2634_v18_] if (d_2634_v18_) in (d_2636_v20_) else d_2611_v3_)).set((d_2617_v6_).f29, p0)).set((len(_dafny.SeqWithoutIsStrInference([d_2617_v6_.f28]))) != ((d_2608_v0_).f39), (0) - ((d_2608_v0_).f39))
                d_2637_v21_: str
                d_2637_v21_ = _dafny.CodePoint('f')
                d_2638_v22_: D9
                d_2638_v22_ = D9_DC24(_dafny.CodePoint('u'), p1, (d_2608_v0_).f39, d_2619_v8_, False)
                d_2639_v23_: _dafny.Array
                nw465_ = _dafny.Array(None, 23)
                nw465_[int(0)] = d_2637_v21_
                nw465_[int(1)] = d_2637_v21_
                nw465_[int(2)] = d_2637_v21_
                nw465_[int(3)] = _dafny.CodePoint('h')
                nw465_[int(4)] = default__.fm42(len(_dafny.SeqWithoutIsStrInference([d_2637_v21_])), 950, globalState)
                nw465_[int(5)] = d_2637_v21_
                nw465_[int(6)] = d_2637_v21_
                nw465_[int(7)] = d_2637_v21_
                nw465_[int(8)] = d_2637_v21_
                nw465_[int(9)] = d_2637_v21_
                nw465_[int(10)] = d_2637_v21_
                nw465_[int(11)] = _dafny.CodePoint('c')
                nw465_[int(12)] = ((self).f31)[default__.safeIndex(p0, len((self).f31))]
                nw465_[int(13)] = d_2637_v21_
                nw465_[int(14)] = _dafny.CodePoint('w')
                nw465_[int(15)] = d_2637_v21_
                nw465_[int(16)] = d_2637_v21_
                nw465_[int(17)] = d_2637_v21_
                nw465_[int(18)] = d_2637_v21_
                nw465_[int(19)] = d_2637_v21_
                nw465_[int(20)] = _dafny.CodePoint('b')
                nw465_[int(21)] = default__.fm42((d_2608_v0_).f39, len((self).f31), globalState)
                nw465_[int(22)] = (d_2638_v22_).cf29
                d_2639_v23_ = nw465_
                d_2640_v24_: _dafny.Array
                nw466_ = _dafny.Array(D17.default()(), 29)
                d_2640_v24_ = nw466_
                d_2641_v25_: D22
                d_2641_v25_ = D22_DC58(d_2640_v24_)
                d_2642_v26_: _dafny.Seq
                d_2642_v26_ = _dafny.SeqWithoutIsStrInference([d_2641_v25_, d_2641_v25_, D22_DC58(d_2640_v24_)])
                d_2643_v27_: _dafny.MultiSet
                d_2643_v27_ = _dafny.MultiSet([D22_DC58(d_2640_v24_)])
                d_2644_v28_: _dafny.Seq
                d_2644_v28_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_2642_v26_)) | (d_2643_v27_)])
                d_2645_v29_: _dafny.MultiSet
                d_2645_v29_ = _dafny.MultiSet([d_2612_v4_])
                d_2646_v30_: _dafny.MultiSet
                d_2646_v30_ = _dafny.MultiSet([default__.fm29(d_2645_v29_, p0, globalState)])
                d_2647_v31_: _dafny.Set
                d_2647_v31_ = _dafny.Set({(d_2608_v0_).f39, default__.fm8(globalState)})
                d_2648_v32_: _dafny.Map
                d_2648_v32_ = _dafny.Map({d_2617_v6_.f28: (d_2617_v6_).f29})
                d_2649_v33_: D10
                d_2649_v33_ = D10_DC28(d_2648_v32_)
                d_2650_v34_: _dafny.Map
                d_2650_v34_ = _dafny.Map({d_2649_v33_: d_2639_v23_})
                rhs443_ = default__.safeModuloInt(543, ((d_2646_v30_)[(d_2617_v6_).f29] if ((d_2617_v6_).f29) in (d_2646_v30_) else (self).f32))
                rhs444_ = len((d_2647_v31_) | (d_2647_v31_))
                rhs445_ = ((d_2650_v34_)[d_2649_v33_] if (d_2649_v33_) in (d_2650_v34_) else d_2639_v23_)
                rhs446_ = d_2644_v28_
                rhs447_ = default__.fm38(globalState)
                lhs366_ = globalState
                lhs367_ = globalState
                lhs368_ = d_2617_v6_
                lhs366_.f24 = rhs443_
                lhs367_.f24 = rhs444_
                d_2639_v23_ = rhs445_
                d_2644_v28_ = rhs446_
                lhs368_.f28 = rhs447_
                (globalState).f21 = default__.fm18(((d_2608_v0_).f39) - ((self).f32), ((d_2646_v30_)[d_2617_v6_.f28] if (d_2617_v6_.f28) in (d_2646_v30_) else p0), (d_2608_v0_).fm30(d_2647_v31_, globalState), globalState)
                (globalState).f24 = 809
                r0 = 736
            d_2651_v35_: D21
            d_2651_v35_ = D21_DC56()
            d_2652_v36_: D21
            d_2652_v36_ = D21_DC57(D21_DC57(d_2651_v35_))
            d_2652_v36_ = d_2652_v36_
            d_2653_v37_: _dafny.Seq
            d_2653_v37_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_2608_v0_).f39)])
            d_2654_v38_: D1
            d_2654_v38_ = D1_DC2((self).f32, (d_2608_v0_).f39, (self).f31, d_2617_v6_.f28)
            d_2655_v39_: _dafny.Map
            d_2655_v39_ = _dafny.Map({(len(d_2653_v37_)) - (p0): (d_2654_v38_).cf2})
            d_2655_v39_ = (d_2655_v39_).set(p0, (d_2608_v0_).f39)
            d_2656_v40_: _dafny.Seq
            d_2656_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjgiawls"))
            d_2657_v41_: _dafny.Array
            def lambda122_(d_2658_v0_):
                def lambda123_(d_2659_i2_):
                    return (d_2659_i2_) + ((d_2658_v0_).f39)

                return lambda123_

            init72_ = lambda122_(d_2608_v0_)
            nw467_ = _dafny.Array(None, 11)
            for i0_72_ in range(nw467_.length(0)):
                nw467_[i0_72_] = init72_(i0_72_)
            d_2657_v41_ = nw467_
            d_2660_v42_: _dafny.Map
            d_2660_v42_ = _dafny.Map({((d_2611_v3_)[(d_2617_v6_).f29] if ((d_2617_v6_).f29) in (d_2611_v3_) else (d_2608_v0_).f39): (d_2617_v6_).f29})
            d_2661_v43_: _dafny.Map
            d_2661_v43_ = _dafny.Map({d_2660_v42_: 635})
            d_2662_v44_: _dafny.Set
            d_2662_v44_ = _dafny.Set({len(d_2661_v43_)})
            index403_ = default__.safeIndex(561, (d_2657_v41_).length(0))
            (d_2657_v41_)[index403_] = len(d_2662_v44_)
            d_2663_v45_: str
            d_2663_v45_ = _dafny.CodePoint('b')
            index404_ = default__.safeIndex(561, (d_2657_v41_).length(0))
            rhs448_ = ((self).f32 if d_2617_v6_.f28 else len((self).f31))
            rhs449_ = (((d_2656_v40_) + ((self).f31)) + (((self).f31).set(default__.safeIndex((d_2608_v0_).fm30(d_2662_v44_, globalState), len((self).f31)), d_2663_v45_))).set(default__.safeIndex((0) - ((d_2608_v0_).f39), len(((d_2656_v40_) + ((self).f31)) + (((self).f31).set(default__.safeIndex((d_2608_v0_).fm30(d_2662_v44_, globalState), len((self).f31)), d_2663_v45_)))), d_2663_v45_)
            rhs450_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgatqfqj"))) + (d_2656_v40_)
            rhs451_ = (len((d_2656_v40_) + (_dafny.SeqWithoutIsStrInference([(d_2656_v40_)[default__.safeIndex((0) - ((d_2608_v0_).f39), len(d_2656_v40_))] for d_2664_i3_ in range(default__.abs(818))])))) * ((d_2608_v0_).f39)
            rhs452_ = d_2617_v6_.f28
            lhs369_ = globalState
            lhs370_ = d_2657_v41_
            lhs371_ = default__.safeIndex(561, (d_2657_v41_).length(0))
            lhs369_.f24 = rhs448_
            d_2656_v40_ = rhs449_
            d_2656_v40_ = rhs450_
            lhs370_[lhs371_] = rhs451_
            r2 = rhs452_
        elif True:
            d_2665___mcc_h1_ = source42_.cf26
            d_2666_cf26_ = d_2665___mcc_h1_
            d_2667_v46_: str
            d_2667_v46_ = _dafny.CodePoint('s')
            if (d_2667_v46_) not in (_dafny.SeqWithoutIsStrInference([d_2667_v46_ for d_2668_i4_ in range(default__.abs(116))])):
                d_2669_v47_: C1
                nw468_ = C1()
                nw468_.ctor__(p1, ((d_2666_cf26_).f29) == ((d_2617_v6_).f29))
                d_2669_v47_ = nw468_
                (globalState).f1 = len(((self).f31) + ((self).f31))
                (globalState).f23 = default__.fm42(default__.fm12((d_2608_v0_).f39, False, globalState), (d_2608_v0_).f39, globalState)
                (globalState).f27 = (d_2666_cf26_).f29
                d_2670_v48_: _dafny.Array
                nw469_ = _dafny.Array(_dafny.Set({}), 2)
                d_2670_v48_ = nw469_
                d_2671_v49_: _dafny.Set
                d_2671_v49_ = _dafny.Set({(d_2608_v0_).f39})
                index405_ = default__.safeIndex(811, (d_2670_v48_).length(0))
                (d_2670_v48_)[index405_] = (d_2671_v49_) | (d_2671_v49_)
                index406_ = default__.safeIndex(811, (d_2670_v48_).length(0))
                (d_2670_v48_)[index406_] = (_dafny.Set({(d_2608_v0_).f39})) - (default__.fm35(globalState))
            elif True:
                d_2672_v50_: _dafny.Array
                nw470_ = _dafny.Array(None, 4)
                nw470_[int(0)] = (d_2666_cf26_).f29
                nw470_[int(1)] = (d_2617_v6_).f29
                nw470_[int(2)] = p1
                nw470_[int(3)] = d_2617_v6_.f28
                d_2672_v50_ = nw470_
                d_2673_v51_: _dafny.Seq
                d_2673_v51_ = _dafny.SeqWithoutIsStrInference([d_2672_v50_, d_2672_v50_])
                d_2674_v52_: D5
                d_2674_v52_ = D5_DC12(d_2672_v50_)
                d_2675_v53_: _dafny.Array
                nw471_ = _dafny.Array(None, 27)
                nw471_[int(0)] = d_2672_v50_
                nw471_[int(1)] = d_2672_v50_
                nw471_[int(2)] = d_2672_v50_
                nw471_[int(3)] = d_2672_v50_
                nw471_[int(4)] = d_2672_v50_
                nw471_[int(5)] = d_2672_v50_
                nw471_[int(6)] = d_2672_v50_
                nw471_[int(7)] = d_2672_v50_
                nw471_[int(8)] = d_2672_v50_
                nw471_[int(9)] = d_2672_v50_
                nw471_[int(10)] = d_2672_v50_
                nw471_[int(11)] = d_2672_v50_
                nw471_[int(12)] = d_2672_v50_
                nw471_[int(13)] = d_2672_v50_
                nw471_[int(14)] = d_2672_v50_
                nw471_[int(15)] = d_2672_v50_
                nw471_[int(16)] = d_2672_v50_
                nw471_[int(17)] = d_2672_v50_
                nw471_[int(18)] = (d_2673_v51_)[default__.safeIndex(-422, len(d_2673_v51_))]
                nw471_[int(19)] = d_2672_v50_
                nw471_[int(20)] = d_2672_v50_
                nw471_[int(21)] = d_2672_v50_
                nw471_[int(22)] = (d_2674_v52_).cf17
                nw471_[int(23)] = d_2672_v50_
                nw471_[int(24)] = d_2672_v50_
                nw471_[int(25)] = d_2672_v50_
                nw471_[int(26)] = d_2672_v50_
                d_2675_v53_ = nw471_
                d_2676_v54_: _dafny.Seq
                d_2676_v54_ = _dafny.SeqWithoutIsStrInference([(self).f31])
                d_2677_v55_: C6
                nw472_ = C6()
                nw472_.ctor__(d_2675_v53_, d_2676_v54_, d_2617_v6_.f28, (d_2666_cf26_.f28) in (d_2612_v4_))
                d_2677_v55_ = nw472_
                (d_2617_v6_).f28 = False
                d_2678_v56_: _dafny.Array
                def lambda124_(d_2679_cf26_, d_2680_v6_):
                    def lambda125_(d_2681_i5_):
                        return (_dafny.Map({(d_2679_cf26_).f29: (d_2680_v6_).f29})) | (_dafny.Map({d_2680_v6_.f28: d_2680_v6_.f28}))

                    return lambda125_

                init73_ = lambda124_(d_2666_cf26_, d_2617_v6_)
                nw473_ = _dafny.Array(None, 15)
                for i0_73_ in range(nw473_.length(0)):
                    nw473_[i0_73_] = init73_(i0_73_)
                d_2678_v56_ = nw473_
                d_2678_v56_ = d_2678_v56_
                d_2682_v57_: D21
                d_2682_v57_ = D21_DC56()
                d_2683_v58_: D21
                d_2683_v58_ = D21_DC57(d_2682_v57_)
                d_2684_v59_: _dafny.Array
                def lambda126_(d_2685_v0_):
                    def lambda127_(d_2686_i6_):
                        return default__.safeDivisionInt(d_2686_i6_, (d_2685_v0_).f39)

                    return lambda127_

                init74_ = lambda126_(d_2608_v0_)
                nw474_ = _dafny.Array(None, 23)
                for i0_74_ in range(nw474_.length(0)):
                    nw474_[i0_74_] = init74_(i0_74_)
                d_2684_v59_ = nw474_
                d_2687_v60_: _dafny.Seq
                d_2687_v60_ = _dafny.SeqWithoutIsStrInference([d_2684_v59_, d_2684_v59_, d_2684_v59_, d_2684_v59_, d_2684_v59_])
                d_2688_v61_: _dafny.Map
                d_2688_v61_ = _dafny.Map({d_2684_v59_: (d_2687_v60_)[default__.safeIndex((self).f32, len(d_2687_v60_))]})
                rhs453_ = ((d_2688_v61_)[d_2684_v59_] if (d_2684_v59_) in (d_2688_v61_) else d_2684_v59_)
                rhs454_ = d_2683_v58_
                r3 = rhs453_
                d_2683_v58_ = rhs454_
                d_2689_v62_: _dafny.MultiSet
                d_2689_v62_ = _dafny.MultiSet([(d_2608_v0_).f39, 404])
                d_2690_v63_: D4
                d_2690_v63_ = D4_DC9((d_2689_v62_).set(p0, default__.abs(p0)))
                d_2691_v64_: C6
                nw475_ = C6()
                nw475_.ctor__((d_2675_v53_ if not(d_2666_cf26_.f28) else (d_2677_v55_).f33), (_dafny.SeqWithoutIsStrInference([(self).f31 for d_2692_i7_ in range(default__.abs(-557))])) + (d_2676_v54_), d_2617_v6_.f28, not((d_2677_v55_).fm22((d_2690_v63_).cf12, globalState)))
                d_2691_v64_ = nw475_
            d_2693_v65_: _dafny.Map
            d_2693_v65_ = _dafny.Map({200: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blawixpn"))})
            d_2694_v66_: _dafny.Array
            def lambda128_(d_2695_v0_, d_2696_v46_):
                def lambda129_(d_2697_i8_):
                    return (((self).f31) + ((self).f31)).set(default__.safeIndex((d_2695_v0_).f39, len(((self).f31) + ((self).f31))), d_2696_v46_)

                return lambda129_

            init75_ = lambda128_(d_2608_v0_, d_2667_v46_)
            nw476_ = _dafny.Array(None, 5)
            for i0_75_ in range(nw476_.length(0)):
                nw476_[i0_75_] = init75_(i0_75_)
            d_2694_v66_ = nw476_
            index407_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            (d_2694_v66_)[index407_] = (self).f31
            d_2698_v67_: _dafny.Array
            nw477_ = _dafny.Array(_dafny.Set({}), 16)
            d_2698_v67_ = nw477_
            d_2699_v68_: _dafny.Map
            d_2699_v68_ = _dafny.Map({d_2698_v67_: d_2693_v65_})
            index408_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            def iife172_():
                coll94_ = _dafny.Map()
                compr_94_: int
                for compr_94_ in _dafny.IntegerRange(498, 124):
                    d_2700_v69_: int = compr_94_
                    if ((498) <= (d_2700_v69_)) and ((d_2700_v69_) < (124)):
                        coll94_[default__.safeModuloInt(d_2700_v69_, (self).f32)] = _dafny.SeqWithoutIsStrInference([d_2667_v46_ for d_2701_i9_ in range(default__.abs(693))])
                return _dafny.Map(coll94_)
            rhs455_ = ((d_2699_v68_)[d_2698_v67_] if (d_2698_v67_) in (d_2699_v68_) else (_dafny.Map({(d_2608_v0_).f39: (self).f31})) | (iife172_()
            ))
            rhs456_ = (((self).f31) + ((self).f31)) + (default__.fm51((d_2608_v0_).f39, (self).f32, globalState))
            lhs372_ = d_2694_v66_
            lhs373_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            d_2693_v65_ = rhs455_
            lhs372_[lhs373_] = rhs456_
            d_2702_v70_: _dafny.Array
            nw478_ = _dafny.Array(int(0), 16)
            d_2702_v70_ = nw478_
            index409_ = default__.safeIndex(377, (d_2702_v70_).length(0))
            (d_2702_v70_)[index409_] = 207
            d_2703_v71_: D17
            d_2703_v71_ = D17_DC44(p1, _dafny.SeqWithoutIsStrInference([d_2667_v46_ for d_2704_i10_ in range(default__.abs(943))]))
            d_2705_v72_: _dafny.Map
            d_2705_v72_ = _dafny.Map({len(default__.fm62(globalState)): True})
            d_2706_v73_: _dafny.Set
            d_2706_v73_ = _dafny.Set({d_2705_v72_, d_2705_v72_})
            index410_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            index411_ = default__.safeIndex(377, (d_2702_v70_).length(0))
            index412_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            rhs457_ = d_2667_v46_
            rhs458_ = ((d_2703_v71_).cf63).set(default__.safeIndex(((0) - ((d_2608_v0_).fm30(_dafny.Set({(d_2608_v0_).f39}), globalState))) - ((d_2608_v0_).f39), len((d_2703_v71_).cf63)), d_2667_v46_)
            rhs459_ = (len(d_2706_v73_)) != ((d_2608_v0_).f39)
            rhs460_ = p0
            rhs461_ = (D15_DC39(((d_2694_v66_)[default__.safeIndex(372, (d_2694_v66_).length(0))]).set(default__.safeIndex((d_2608_v0_).f39, len((d_2694_v66_)[default__.safeIndex(372, (d_2694_v66_).length(0))])), _dafny.CodePoint('j')))).cf56
            lhs374_ = globalState
            lhs375_ = d_2694_v66_
            lhs376_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            lhs377_ = d_2666_cf26_
            lhs378_ = d_2702_v70_
            lhs379_ = default__.safeIndex(377, (d_2702_v70_).length(0))
            lhs380_ = d_2694_v66_
            lhs381_ = default__.safeIndex(372, (d_2694_v66_).length(0))
            lhs374_.f23 = rhs457_
            lhs375_[lhs376_] = rhs458_
            lhs377_.f28 = rhs459_
            lhs378_[lhs379_] = rhs460_
            lhs380_[lhs381_] = rhs461_
            d_2707_v74_: _dafny.Map
            d_2707_v74_ = _dafny.Map({d_2702_v70_: (d_2612_v4_)[default__.safeIndex((d_2608_v0_).f39, len(d_2612_v4_))]})
            d_2707_v74_ = (d_2707_v74_).set(d_2702_v70_, d_2617_v6_.f28)
        r0 = p0
        d_2708_v75_: _dafny.Array
        def lambda130_(d_2709_v0_):
            def lambda131_(d_2710_i11_):
                return default__.safeDivisionInt(d_2710_i11_, (d_2709_v0_).f39)

            return lambda131_

        init76_ = lambda130_(d_2608_v0_)
        nw479_ = _dafny.Array(None, 20)
        for i0_76_ in range(nw479_.length(0)):
            nw479_[i0_76_] = init76_(i0_76_)
        d_2708_v75_ = nw479_
        d_2711_v76_: _dafny.Set
        d_2711_v76_ = _dafny.Set({(self).f32, (self).f32})
        index413_ = default__.safeIndex(290, (d_2708_v75_).length(0))
        (d_2708_v75_)[index413_] = (d_2608_v0_).fm30(d_2711_v76_, globalState)
        index414_ = default__.safeIndex(290, (d_2708_v75_).length(0))
        (d_2708_v75_)[index414_] = (self).f32
        r0 = (0) - ((len(d_2711_v76_)) - (((d_2708_v75_)[default__.safeIndex(290, (d_2708_v75_).length(0))]) * ((self).f32)))
        r1 = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2712_i12_ in range(default__.abs(-786))])) != ((self).f31)) == ((d_2617_v6_).f29)
        r2 = p1
        r3 = d_2708_v75_
        return r0, r1, r2, r3

    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32

class C11(T0):
    def  __init__(self):
        self._f28: bool = False
        self._f29: bool = False
        self.f30: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f28(self):
        return self._f28
    @f28.setter
    def f28(self, value):
        self._f28 = value
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f30, f28, f29):
        (self).f30 = f30
        (self).f28 = f28
        (self)._f29 = f29

    def fm7(self, p0, p1, globalState):
        return (self).f29

    def m5(self, p0, globalState):
        r0: bool = False
        d_2713_v0_: _dafny.Map
        d_2713_v0_ = _dafny.Map({_dafny.Map({(self).f29: (self).f29}): 505})
        d_2714_v1_: C1
        nw480_ = C1()
        nw480_.ctor__((self).f29, (d_2713_v0_) == (d_2713_v0_))
        d_2714_v1_ = nw480_
        d_2715_v2_: _dafny.Array
        nw481_ = _dafny.Array(False, 12)
        d_2715_v2_ = nw481_
        index415_ = default__.safeIndex(338, (d_2715_v2_).length(0))
        (d_2715_v2_)[index415_] = self.f28
        index416_ = default__.safeIndex(338, (d_2715_v2_).length(0))
        (d_2715_v2_)[index416_] = ((self).f29 if (self).f29 else (self).f29)
        (globalState).f24 = p0
        d_2716_v3_: _dafny.Seq
        d_2716_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gp"))
        d_2717_v4_: _dafny.Map
        d_2717_v4_ = _dafny.Map({default__.fm51(p0, p0, globalState): p0})
        d_2718_i0_: int
        d_2718_i0_ = 0
        with _dafny.label("8"):
            while (d_2716_v3_) not in ((d_2717_v4_) | (d_2717_v4_)):
                with _dafny.c_label("8"):
                    if (d_2718_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_2718_i0_ = (d_2718_i0_) + (1)
                    if (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]:
                        d_2719_v5_: _dafny.Map
                        d_2719_v5_ = _dafny.Map({p0: (self).f29})
                        d_2720_v6_: D10
                        d_2720_v6_ = D10_DC26(d_2719_v5_)
                        d_2721_v7_: _dafny.MultiSet
                        d_2721_v7_ = _dafny.MultiSet([(d_2720_v6_).cf39])
                        d_2721_v7_ = (d_2721_v7_) - ((d_2721_v7_).intersection(d_2721_v7_))
                        (globalState).f24 = (p0) * (p0)
                        (globalState).f1 = p0
                        d_2722_v8_: _dafny.Set
                        d_2722_v8_ = _dafny.Set({d_2716_v3_})
                        d_2723_v9_: _dafny.Seq
                        d_2723_v9_ = _dafny.SeqWithoutIsStrInference([(d_2714_v1_).fm36((_dafny.MultiSet([(self).f29])).cardinality, d_2722_v8_, globalState)])
                        d_2724_v10_: _dafny.MultiSet
                        d_2724_v10_ = _dafny.MultiSet([(0) - ((d_2723_v9_)[default__.safeIndex(771, len(d_2723_v9_))]), p0])
                        d_2725_v11_: C0
                        nw482_ = C0()
                        nw482_.ctor__()
                        d_2725_v11_ = nw482_
                        d_2726_v12_: _dafny.Set
                        d_2726_v12_ = _dafny.Set({d_2725_v11_, d_2725_v11_})
                        d_2727_v13_: _dafny.MultiSet
                        d_2727_v13_ = _dafny.MultiSet([(p0) * ((d_2724_v10_).cardinality), len(d_2726_v12_), p0, p0, default__.fm25(p0, p0, globalState)])
                        d_2728_v14_: _dafny.Array
                        def lambda132_(d_2729_i1_):
                            return default__.safeModuloInt(d_2729_i1_, 158)

                        init77_ = lambda132_
                        nw483_ = _dafny.Array(None, 20)
                        for i0_77_ in range(nw483_.length(0)):
                            nw483_[i0_77_] = init77_(i0_77_)
                        d_2728_v14_ = nw483_
                        d_2730_v15_: _dafny.Map
                        d_2730_v15_ = _dafny.Map({True: (self).f29})
                        d_2731_v16_: _dafny.Map
                        d_2731_v16_ = _dafny.Map({d_2728_v14_: d_2730_v15_})
                        d_2732_v17_: _dafny.Map
                        d_2732_v17_ = _dafny.Map({d_2716_v3_: (d_2731_v16_).set(d_2728_v14_, d_2730_v15_)})
                        rhs462_ = (0) - (len((((d_2732_v17_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gj"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gj"))) in (d_2732_v17_) else d_2731_v16_)) | (d_2731_v16_)))
                        rhs463_ = (d_2727_v13_ if (self).f29 else d_2724_v10_)
                        rhs464_ = (self).f29
                        lhs382_ = globalState
                        lhs383_ = globalState
                        lhs382_.f24 = rhs462_
                        d_2724_v10_ = rhs463_
                        lhs383_.f21 = rhs464_
                        index417_ = default__.safeIndex(327, (d_2728_v14_).length(0))
                        (d_2728_v14_)[index417_] = p0
                        index418_ = default__.safeIndex(896, (d_2728_v14_).length(0))
                        (d_2728_v14_)[index418_] = (0) - (p0)
                        d_2733_v18_: _dafny.Seq
                        d_2733_v18_ = _dafny.SeqWithoutIsStrInference([(d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))], (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]])
                        d_2734_v19_: _dafny.Map
                        d_2734_v19_ = _dafny.Map({default__.fm29(_dafny.MultiSet([d_2733_v18_, d_2733_v18_, d_2733_v18_, d_2733_v18_]), p0, globalState): p0})
                        index419_ = default__.safeIndex(327, (d_2728_v14_).length(0))
                        index420_ = default__.safeIndex(896, (d_2728_v14_).length(0))
                        rhs465_ = ((d_2734_v19_)[(self).f29] if ((self).f29) in (d_2734_v19_) else (p0) + (p0))
                        rhs466_ = (p0) * (p0)
                        rhs467_ = False
                        rhs468_ = p0
                        rhs469_ = (self).f29
                        lhs384_ = d_2728_v14_
                        lhs385_ = default__.safeIndex(327, (d_2728_v14_).length(0))
                        lhs386_ = globalState
                        lhs387_ = globalState
                        lhs388_ = d_2728_v14_
                        lhs389_ = default__.safeIndex(896, (d_2728_v14_).length(0))
                        lhs390_ = globalState
                        lhs384_[lhs385_] = rhs465_
                        lhs386_.f24 = rhs466_
                        lhs387_.f13 = rhs467_
                        lhs388_[lhs389_] = rhs468_
                        lhs390_.f27 = rhs469_
                    elif True:
                        (globalState).f27 = (871) <= (p0)
                        r0 = (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]
                        (globalState).f13 = (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]
                        d_2715_v2_ = d_2715_v2_
                        index421_ = default__.safeIndex(338, (d_2715_v2_).length(0))
                        rhs470_ = p0
                        rhs471_ = p0
                        rhs472_ = (self).f29
                        lhs391_ = globalState
                        lhs392_ = globalState
                        lhs393_ = d_2715_v2_
                        lhs394_ = default__.safeIndex(338, (d_2715_v2_).length(0))
                        lhs391_.f1 = rhs470_
                        lhs392_.f24 = rhs471_
                        lhs393_[lhs394_] = rhs472_
                    d_2735_v20_: _dafny.Array
                    nw484_ = _dafny.Array(int(0), 7)
                    d_2735_v20_ = nw484_
                    d_2736_v21_: _dafny.Map
                    d_2736_v21_ = _dafny.Map({p0: d_2735_v20_})
                    d_2737_v22_: _dafny.MultiSet
                    d_2737_v22_ = _dafny.MultiSet([not ((self).f29) or (not(not((self).f29)))])
                    d_2738_v23_: _dafny.Set
                    d_2738_v23_ = _dafny.Set({self.f28, self.f28, (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]})
                    d_2739_v24_: _dafny.Set
                    d_2739_v24_ = _dafny.Set({p0})
                    rhs473_ = ((D23_DC62(d_2736_v21_)).cf91) | ((d_2736_v21_) | (d_2736_v21_))
                    rhs474_ = ((d_2738_v23_) - (_dafny.Set({False, default__.fm13(globalState)}))).isdisjoint(d_2738_v23_)
                    rhs475_ = (d_2737_v22_ if ((self).f29 if not(not((d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))])) else (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]) else d_2737_v22_)
                    rhs476_ = default__.safeModuloInt(p0, p0)
                    rhs477_ = (default__.safeDivisionInt(p0, p0)) in (d_2739_v24_)
                    lhs395_ = globalState
                    lhs396_ = globalState
                    lhs397_ = globalState
                    d_2736_v21_ = rhs473_
                    lhs395_.f21 = rhs474_
                    d_2737_v22_ = rhs475_
                    lhs396_.f1 = rhs476_
                    lhs397_.f21 = rhs477_
                    d_2740_v25_: _dafny.Array
                    nw485_ = _dafny.Array(_dafny.MultiSet({}), 25)
                    d_2740_v25_ = nw485_
                    index422_ = default__.safeIndex(319, (d_2740_v25_).length(0))
                    (d_2740_v25_)[index422_] = d_2737_v22_
                    d_2741_v26_: str
                    d_2741_v26_ = _dafny.CodePoint('g')
                    d_2742_v27_: _dafny.MultiSet
                    d_2742_v27_ = _dafny.MultiSet([p0])
                    d_2743_v28_: _dafny.Map
                    d_2743_v28_ = _dafny.Map({default__.fm42(len(_dafny.SeqWithoutIsStrInference([p0])), len(d_2716_v3_), globalState): (self).f29})
                    d_2744_v29_: _dafny.Seq
                    d_2744_v29_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                    index423_ = default__.safeIndex(319, (d_2740_v25_).length(0))
                    (d_2740_v25_)[index423_] = ((default__.fm57(d_2741_v26_, p0, ((d_2742_v27_).set(len(d_2743_v28_), default__.abs(p0))).cardinality, globalState) if default__.fm38(globalState) else (d_2737_v22_).set((d_2744_v29_)[default__.safeIndex(len(d_2716_v3_), len(d_2744_v29_))], default__.abs(50)))).intersection((_dafny.MultiSet([(d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))]]) if (d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))] else _dafny.MultiSet(d_2744_v29_)))
                    d_2741_v26_ = d_2741_v26_
                    pass
            pass
        d_2745_v30_: C9
        nw486_ = C9()
        nw486_.ctor__(self.f28, (self).f29)
        d_2745_v30_ = nw486_
        nw487_ = C9()
        nw487_.ctor__((d_2715_v2_)[default__.safeIndex(338, (d_2715_v2_).length(0))], False)
        d_2745_v30_ = nw487_
        d_2746_v31_: D5
        d_2746_v31_ = D5_DC12(d_2715_v2_)
        d_2746_v31_ = d_2746_v31_
        r0 = (self).f29
        return r0

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Seq = _dafny.Seq({})
        r3: D1 = D1.default()()
        if p3:
            d_2747_v0_: int
            d_2747_v0_ = 366
            d_2748_v1_: _dafny.MultiSet
            d_2748_v1_ = _dafny.MultiSet([d_2747_v0_])
            d_2749_v2_: _dafny.Seq
            d_2749_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2747_v0_, d_2747_v0_, d_2747_v0_, d_2747_v0_]), d_2748_v1_])
            d_2749_v2_ = (d_2749_v2_) + (d_2749_v2_)
            index424_ = default__.safeIndex(966, (p2).length(0))
            (p2)[index424_] = False
            index425_ = default__.safeIndex(966, (p2).length(0))
            (p2)[index425_] = (d_2747_v0_) != (d_2747_v0_)
            d_2750_v3_: _dafny.Array
            nw488_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_2750_v3_ = nw488_
            d_2751_v4_: _dafny.Array
            def lambda133_(d_2752_p0_):
                def lambda134_(d_2753_i0_):
                    return (d_2753_i0_) - ((0) - (len(d_2752_p0_)))

                return lambda134_

            init78_ = lambda133_(p0)
            nw489_ = _dafny.Array(None, 2)
            for i0_78_ in range(nw489_.length(0)):
                nw489_[i0_78_] = init78_(i0_78_)
            d_2751_v4_ = nw489_
            index426_ = default__.safeIndex(173, (d_2751_v4_).length(0))
            (d_2751_v4_)[index426_] = (0) - (d_2747_v0_)
            index427_ = default__.safeIndex(255, (d_2751_v4_).length(0))
            (d_2751_v4_)[index427_] = d_2747_v0_
            d_2754_v5_: str
            d_2754_v5_ = _dafny.CodePoint('n')
            d_2755_v6_: _dafny.Seq
            d_2755_v6_ = _dafny.SeqWithoutIsStrInference([(self).f29, True, self.f28])
            index428_ = default__.safeIndex(173, (d_2751_v4_).length(0))
            index429_ = default__.safeIndex(255, (d_2751_v4_).length(0))
            rhs478_ = d_2750_v3_
            rhs479_ = d_2747_v0_
            rhs480_ = d_2754_v5_
            rhs481_ = default__.safeModuloInt(d_2747_v0_, len(d_2755_v6_))
            rhs482_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg"))).set(default__.safeIndex(d_2747_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg")))), d_2754_v5_))
            lhs398_ = globalState
            lhs399_ = d_2751_v4_
            lhs400_ = default__.safeIndex(173, (d_2751_v4_).length(0))
            lhs401_ = d_2751_v4_
            lhs402_ = default__.safeIndex(255, (d_2751_v4_).length(0))
            d_2750_v3_ = rhs478_
            d_2747_v0_ = rhs479_
            lhs398_.f23 = rhs480_
            lhs399_[lhs400_] = rhs481_
            lhs401_[lhs402_] = rhs482_
            d_2756_v7_: _dafny.Seq
            d_2756_v7_ = _dafny.SeqWithoutIsStrInference([p0])
            index430_ = default__.safeIndex(173, (d_2751_v4_).length(0))
            (d_2751_v4_)[index430_] = len(((d_2756_v7_)[default__.safeIndex(d_2747_v0_, len(d_2756_v7_))]) + (p0))
            if (p2)[default__.safeIndex(966, (p2).length(0))]:
                index431_ = default__.safeIndex(966, (p2).length(0))
                rhs483_ = (self).f29
                rhs484_ = ((0) - (d_2747_v0_)) != ((d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))])
                rhs485_ = d_2754_v5_
                lhs403_ = globalState
                lhs404_ = p2
                lhs405_ = default__.safeIndex(966, (p2).length(0))
                lhs406_ = globalState
                lhs403_.f22 = rhs483_
                lhs404_[lhs405_] = rhs484_
                lhs406_.f23 = rhs485_
                d_2757_v8_: _dafny.Map
                d_2757_v8_ = _dafny.Map({p0: (self).f29})
                index432_ = default__.safeIndex(966, (p2).length(0))
                index433_ = default__.safeIndex(173, (d_2751_v4_).length(0))
                rhs486_ = ((d_2757_v8_)[(default__.fm51((d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))], 305, globalState)) + (p0)] if ((default__.fm51((d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))], 305, globalState)) + (p0)) in (d_2757_v8_) else (self).f29)
                rhs487_ = d_2747_v0_
                lhs407_ = p2
                lhs408_ = default__.safeIndex(966, (p2).length(0))
                lhs409_ = d_2751_v4_
                lhs410_ = default__.safeIndex(173, (d_2751_v4_).length(0))
                lhs407_[lhs408_] = rhs486_
                lhs409_[lhs410_] = rhs487_
                r2 = d_2755_v6_
                d_2758_v9_: _dafny.Seq
                d_2758_v9_ = _dafny.SeqWithoutIsStrInference([len(p0), default__.safeDivisionInt(len(d_2755_v6_), (d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))])])
                (globalState).f17 = d_2758_v9_
                d_2759_v10_: C8
                nw490_ = C8()
                nw490_.ctor__()
                d_2759_v10_ = nw490_
            elif True:
                index434_ = default__.safeIndex(173, (d_2751_v4_).length(0))
                (d_2751_v4_)[index434_] = d_2747_v0_
                (globalState).f1 = d_2747_v0_
                d_2760_v11_: _dafny.Map
                d_2760_v11_ = _dafny.Map({(d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))]: False})
                d_2761_v12_: C3
                nw491_ = C3()
                nw491_.ctor__(self.f28, ((d_2760_v11_)[default__.fm27(len(self.f30), _dafny.CodePoint('v'), (d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))], globalState)] if (default__.fm27(len(self.f30), _dafny.CodePoint('v'), (d_2751_v4_)[default__.safeIndex(173, (d_2751_v4_).length(0))], globalState)) in (d_2760_v11_) else self.f28), ((d_2760_v11_)[-207] if (-207) in (d_2760_v11_) else self.f28), (0) - ((default__.fm74(default__.fm47(globalState), (self).f29, self.f28, True, globalState)).cf43))
                d_2761_v12_ = nw491_
                (self).f28 = not(p3)
                (globalState).f22 = (self).f29
        elif True:
            d_2762_v13_: _dafny.Seq
            d_2762_v13_ = _dafny.SeqWithoutIsStrInference([(self).f29, False])
            d_2763_v14_: _dafny.Map
            d_2763_v14_ = _dafny.Map({self.f28: (d_2762_v13_) + ((d_2762_v13_).set(default__.safeIndex(483, len(d_2762_v13_)), self.f28))})
            d_2763_v14_ = (d_2763_v14_).set((self).f29, d_2762_v13_)
            if not(self.f28):
                d_2764_v15_: C9
                nw492_ = C9()
                nw492_.ctor__(p3, self.f28)
                d_2764_v15_ = nw492_
                d_2765_v16_: _dafny.Array
                def lambda135_(d_2766_p0_):
                    def lambda136_(d_2767_i1_):
                        return (d_2766_p0_ if (self).f29 else d_2766_p0_)

                    return lambda136_

                init79_ = lambda135_(p0)
                nw493_ = _dafny.Array(None, 24)
                for i0_79_ in range(nw493_.length(0)):
                    nw493_[i0_79_] = init79_(i0_79_)
                d_2765_v16_ = nw493_
                index435_ = default__.safeIndex(577, (d_2765_v16_).length(0))
                (d_2765_v16_)[index435_] = p0
                index436_ = default__.safeIndex(577, (d_2765_v16_).length(0))
                (d_2765_v16_)[index436_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oe"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2768_i2_ in range(default__.abs(361))]))
                d_2769_v17_: int
                d_2769_v17_ = 362
                (globalState).f22 = (d_2769_v17_) != (d_2769_v17_)
                (globalState).f13 = p3
                d_2770_v18_: _dafny.Array
                nw494_ = _dafny.Array(int(0), 9)
                d_2770_v18_ = nw494_
                index437_ = default__.safeIndex(467, (d_2770_v18_).length(0))
                (d_2770_v18_)[index437_] = d_2769_v17_
                index438_ = default__.safeIndex(467, (d_2770_v18_).length(0))
                (d_2770_v18_)[index438_] = d_2769_v17_
            elif True:
                rhs488_ = _dafny.CodePoint('b')
                rhs489_ = (d_2762_v13_) + (_dafny.SeqWithoutIsStrInference([not(self.f28), self.f28, (self).f29]))
                rhs490_ = (self).f29
                lhs411_ = globalState
                lhs412_ = globalState
                lhs411_.f23 = rhs488_
                d_2762_v13_ = rhs489_
                lhs412_.f22 = rhs490_
                d_2771_v19_: _dafny.Array
                nw495_ = _dafny.Array(_dafny.Map({}), 25)
                d_2771_v19_ = nw495_
                d_2772_v20_: _dafny.Seq
                d_2772_v20_ = _dafny.SeqWithoutIsStrInference([d_2771_v19_])
                d_2773_v21_: int
                d_2773_v21_ = 332
                d_2771_v19_ = (d_2772_v20_)[default__.safeIndex(d_2773_v21_, len(d_2772_v20_))]
                d_2774_v22_: _dafny.Seq
                d_2774_v22_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_2775_v23_: _dafny.Map
                d_2775_v23_ = _dafny.Map({d_2774_v22_: d_2773_v21_})
                d_2775_v23_ = (d_2775_v23_).set(d_2774_v22_, d_2773_v21_)
                d_2776_v24_: _dafny.Map
                d_2776_v24_ = _dafny.Map({p1: _dafny.Map({self.f28: not((self).f29)})})
                d_2777_v25_: D5
                d_2777_v25_ = D5_DC12(p2)
                d_2778_v26_: _dafny.Map
                d_2778_v26_ = _dafny.Map({self.f28: p3})
                d_2779_v27_: T0
                nw496_ = C7()
                nw496_.ctor__(False, (self).f29)
                d_2779_v27_ = nw496_
                d_2780_v28_: _dafny.Set
                d_2780_v28_ = _dafny.Set({d_2779_v27_})
                d_2781_v29_: _dafny.Map
                d_2781_v29_ = _dafny.Map({d_2780_v28_: _dafny.Map({p2: _dafny.Map({d_2779_v27_.f28: (self).f29})})})
                d_2776_v24_ = ((d_2776_v24_).set((d_2777_v25_).cf17, d_2778_v26_)) | (((d_2781_v29_)[d_2780_v28_] if (d_2780_v28_) in (d_2781_v29_) else d_2776_v24_))
                d_2782_v30_: D7
                d_2782_v30_ = D7_DC17(d_2779_v27_.f28)
                def iife173_(_pat_let39_0):
                    def iife174_(d_2783_dt__update__tmp_h0_):
                        def iife175_(_pat_let40_0):
                            def iife176_(d_2784_dt__update_hcf22_h0_):
                                return D7_DC17(d_2784_dt__update_hcf22_h0_)
                            return iife176_(_pat_let40_0)
                        return iife175_((self).f29)
                    return iife174_(_pat_let39_0)
                d_2782_v30_ = iife173_(d_2782_v30_)
            d_2785_v31_: _dafny.Map
            d_2785_v31_ = _dafny.Map({p3: (self).f29})
            d_2785_v31_ = (d_2785_v31_).set(((self).f29 if p3 else p3), (not((self).f29) if (self).f29 else True))
            d_2786_v32_: int
            d_2786_v32_ = 240
            d_2787_v33_: _dafny.Map
            d_2787_v33_ = _dafny.Map({p3: default__.fm25(d_2786_v32_, d_2786_v32_, globalState)})
            source43_ = D20_DC51((len(d_2787_v33_)) * (default__.fm25(997, len(p0), globalState)), not (not(self.f28)) or (p3), (self.f28) or ((self).f29))
            if source43_.is_DC51:
                d_2788___mcc_h0_ = source43_.cf75
                d_2789___mcc_h1_ = source43_.cf76
                d_2790___mcc_h2_ = source43_.cf77
                d_2791_cf77_ = d_2790___mcc_h2_
                d_2792_cf76_ = d_2789___mcc_h1_
                d_2793_cf75_ = d_2788___mcc_h0_
                d_2792_cf76_ = not ((self).fm7(d_2793_cf75_, d_2792_cf76_, globalState)) or (p3)
                d_2785_v31_ = d_2785_v31_
                d_2794_v34_: _dafny.Array
                nw497_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_2794_v34_ = nw497_
                d_2794_v34_ = (d_2794_v34_ if (self).f29 else d_2794_v34_)
                d_2795_v35_: str
                d_2795_v35_ = _dafny.CodePoint('o')
                d_2796_v36_: D2
                d_2796_v36_ = D2_DC4(d_2795_v35_)
                d_2797_v37_: _dafny.Seq
                def iife177_(_pat_let41_0):
                    def iife178_(d_2798_dt__update__tmp_h1_):
                        def iife179_(_pat_let42_0):
                            def iife180_(d_2799_dt__update_hcf7_h0_):
                                return D2_DC4(d_2799_dt__update_hcf7_h0_)
                            return iife180_(_pat_let42_0)
                        return iife179_(_dafny.CodePoint('a'))
                    return iife178_(_pat_let41_0)
                d_2797_v37_ = _dafny.SeqWithoutIsStrInference([(iife177_(d_2796_v36_)).cf7])
                d_2800_v38_: _dafny.Set
                d_2800_v38_ = _dafny.Set({False, d_2792_cf76_, d_2792_cf76_, self.f28, (self).f29})
                d_2797_v37_ = ((_dafny.SeqWithoutIsStrInference([d_2795_v35_])).set(default__.safeIndex(len(d_2800_v38_), len(_dafny.SeqWithoutIsStrInference([d_2795_v35_]))), _dafny.CodePoint('b'))) + (_dafny.SeqWithoutIsStrInference([d_2795_v35_ for d_2801_i3_ in range(default__.abs(671))]))
            elif source43_.is_DC52:
                d_2762_v13_ = d_2762_v13_
                (globalState).f1 = d_2786_v32_
                d_2802_v39_: _dafny.Array
                nw498_ = _dafny.Array(None, 7)
                nw498_[int(0)] = d_2786_v32_
                nw498_[int(1)] = d_2786_v32_
                nw498_[int(2)] = d_2786_v32_
                nw498_[int(3)] = d_2786_v32_
                nw498_[int(4)] = d_2786_v32_
                nw498_[int(5)] = d_2786_v32_
                nw498_[int(6)] = d_2786_v32_
                d_2802_v39_ = nw498_
                index439_ = default__.safeIndex(3, (d_2802_v39_).length(0))
                (d_2802_v39_)[index439_] = d_2786_v32_
                index440_ = default__.safeIndex(3, (d_2802_v39_).length(0))
                (d_2802_v39_)[index440_] = d_2786_v32_
                (globalState).f27 = p3
            elif source43_.is_DC50:
                d_2803___mcc_h3_ = source43_.cf74
                d_2804_cf74_ = d_2803___mcc_h3_
                d_2805_v40_: str
                d_2805_v40_ = _dafny.CodePoint('h')
                d_2806_v41_: C5
                nw499_ = C5()
                nw499_.ctor__(False, d_2805_v40_, p3, not(default__.fm13(globalState)))
                d_2806_v41_ = nw499_
                d_2807_v42_: D24
                d_2807_v42_ = D24_DC66(d_2806_v41_)
                d_2806_v41_ = (d_2806_v41_ if (d_2806_v41_.f35) == (True) else (d_2807_v42_).cf99)
                index441_ = default__.safeIndex(840, (p1).length(0))
                (p1)[index441_] = p3
                index442_ = default__.safeIndex(840, (p1).length(0))
                (p1)[index442_] = (d_2806_v41_.f35) == (self.f28)
                (globalState).f13 = self.f28
                d_2808_v43_: _dafny.Array
                def lambda137_(d_2809_p3_):
                    def lambda138_(d_2810_i4_):
                        return d_2809_p3_

                    return lambda138_

                init80_ = lambda137_(p3)
                nw500_ = _dafny.Array(None, 19)
                for i0_80_ in range(nw500_.length(0)):
                    nw500_[i0_80_] = init80_(i0_80_)
                d_2808_v43_ = nw500_
                d_2808_v43_ = p2
            elif True:
                d_2811___mcc_h4_ = source43_.cf78
                d_2812_cf78_ = d_2811___mcc_h4_
                (globalState).f1 = -294
                d_2813_v44_: _dafny.Map
                d_2813_v44_ = _dafny.Map({d_2786_v32_: True})
                d_2814_v45_: D1
                d_2814_v45_ = D1_DC2(len(d_2813_v44_), d_2786_v32_, p0, True)
                d_2815_v46_: D1
                d_2815_v46_ = D1_DC3(d_2814_v45_)
                d_2816_v47_: D1
                d_2816_v47_ = D1_DC3(d_2815_v46_)
                r3 = d_2816_v47_
                d_2817_v48_: str
                d_2817_v48_ = _dafny.CodePoint('w')
                d_2818_v49_: _dafny.Map
                d_2818_v49_ = _dafny.Map({d_2817_v48_: d_2786_v32_})
                d_2819_v50_: _dafny.Seq
                d_2819_v50_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2818_v49_ = (d_2818_v49_).set(d_2817_v48_, (len(_dafny.Map({d_2786_v32_: default__.fm18(d_2786_v32_, (_dafny.MultiSet([self.f28, p3])).cardinality, len((d_2819_v50_)[default__.safeIndex(249, len(d_2819_v50_))]), globalState)})) if self.f28 else d_2786_v32_))
                d_2820_v51_: _dafny.Array
                nw501_ = _dafny.Array(int(0), 19)
                d_2820_v51_ = nw501_
                index443_ = default__.safeIndex(736, (d_2820_v51_).length(0))
                (d_2820_v51_)[index443_] = d_2786_v32_
                index444_ = default__.safeIndex(736, (d_2820_v51_).length(0))
                (d_2820_v51_)[index444_] = d_2786_v32_
            index445_ = default__.safeIndex(247, (p2).length(0))
            (p2)[index445_] = p3
            index446_ = default__.safeIndex(247, (p2).length(0))
            (p2)[index446_] = (self.f28 if (self).f29 else self.f28)
        d_2821_v52_: _dafny.MultiSet
        d_2821_v52_ = _dafny.MultiSet([not((self).f29), (self).f29])
        if ((d_2821_v52_) - (d_2821_v52_)).ispropersubset(_dafny.MultiSet([True, p3])):
            d_2822_v53_: _dafny.Array
            nw502_ = _dafny.Array(_dafny.Seq({}), 1)
            d_2822_v53_ = nw502_
            d_2823_v54_: _dafny.Seq
            d_2823_v54_ = _dafny.SeqWithoutIsStrInference([len(p0)])
            index447_ = default__.safeIndex(103, (d_2822_v53_).length(0))
            (d_2822_v53_)[index447_] = d_2823_v54_
            index448_ = default__.safeIndex(103, (d_2822_v53_).length(0))
            (d_2822_v53_)[index448_] = (d_2823_v54_ if False else d_2823_v54_)
            (self).f28 = self.f28
            d_2824_v55_: C7
            nw503_ = C7()
            nw503_.ctor__(p3, self.f28)
            d_2824_v55_ = nw503_
            d_2825_v56_: _dafny.Map
            d_2825_v56_ = _dafny.Map({p3: d_2823_v54_})
            d_2825_v56_ = (d_2825_v56_ if (self).f29 else d_2825_v56_)
            d_2826_v57_: D7
            d_2826_v57_ = D7_DC17(p3)
            d_2827_v58_: _dafny.Seq
            d_2827_v58_ = _dafny.SeqWithoutIsStrInference([self.f28, (self).f29])
            d_2828_v59_: int
            d_2828_v59_ = 734
            d_2826_v57_ = D7_DC17((d_2827_v58_)[default__.safeIndex(d_2828_v59_, len(d_2827_v58_))])
        elif True:
            d_2829_v60_: int
            d_2829_v60_ = 52
            (globalState).f24 = default__.fm25(d_2829_v60_, d_2829_v60_, globalState)
            (globalState).f1 = len((p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2830_i5_ in range(default__.abs(144))])))
            (globalState).f1 = 969
            (globalState).f22 = (self).f29
            d_2831_v61_: _dafny.Map
            d_2831_v61_ = _dafny.Map({d_2829_v60_: d_2829_v60_})
            d_2832_v62_: _dafny.Map
            d_2832_v62_ = _dafny.Map({d_2829_v60_: self.f28})
            d_2831_v61_ = (d_2831_v61_).set(default__.safeModuloInt(d_2829_v60_, len((d_2832_v62_).set(d_2829_v60_, (self).f29))), d_2829_v60_)
        d_2833_v63_: int
        d_2833_v63_ = -700
        d_2834_v64_: D4
        d_2834_v64_ = D4_DC10(self.f28, d_2833_v63_, default__.fm25(d_2833_v63_, d_2833_v63_, globalState))
        d_2835_v65_: _dafny.Map
        d_2835_v65_ = _dafny.Map({default__.fm56(d_2834_v64_, p3, d_2833_v63_, globalState): _dafny.Map({(self).f29: True})})
        d_2835_v65_ = d_2835_v65_
        d_2836_v66_: _dafny.Array
        nw504_ = _dafny.Array(_dafny.Set({}), 7)
        d_2836_v66_ = nw504_
        d_2837_v67_: _dafny.Set
        d_2837_v67_ = _dafny.Set({d_2833_v63_})
        index449_ = default__.safeIndex(529, (d_2836_v66_).length(0))
        (d_2836_v66_)[index449_] = d_2837_v67_
        d_2838_v69_: _dafny.Map
        d_2838_v69_ = _dafny.Map({self.f28: d_2833_v63_})
        d_2839_v70_: _dafny.Seq
        d_2839_v70_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29, p3])
        d_2840_v71_: _dafny.Seq
        d_2840_v71_ = _dafny.SeqWithoutIsStrInference([(d_2839_v70_)[default__.safeIndex(d_2833_v63_, len(d_2839_v70_))]])
        index450_ = default__.safeIndex(529, (d_2836_v66_).length(0))
        def iife181_():
            def iife183_():
                coll97_ = _dafny.Map()
                compr_97_: int
                for compr_97_ in (_dafny.MultiSet([len(d_2838_v69_), len(d_2839_v70_), -425, len(p0)])).Elements:
                    d_2841_v68_: int = compr_97_
                    if (d_2841_v68_) in (_dafny.MultiSet([len(d_2838_v69_), len(d_2839_v70_), -425, len(p0)])):
                        coll97_[(d_2841_v68_) - (d_2833_v63_)] = p3
                return _dafny.Map(coll97_)
            coll95_ = _dafny.Set()
            def iife182_():
                coll96_ = _dafny.Map()
                compr_96_: int
                for compr_96_ in (_dafny.MultiSet([len(d_2838_v69_), len(d_2839_v70_), -425, len(p0)])).Elements:
                    d_2841_v68_: int = compr_96_
                    if (d_2841_v68_) in (_dafny.MultiSet([len(d_2838_v69_), len(d_2839_v70_), -425, len(p0)])):
                        coll96_[(d_2841_v68_) - (d_2833_v63_)] = p3
                return _dafny.Map(coll96_)
            compr_95_: int
            for compr_95_ in (iife182_()
            ).keys.Elements:
                d_2842_v72_: int = compr_95_
                if (d_2842_v72_) in (iife183_()
                ):
                    coll95_ = coll95_.union(_dafny.Set([(d_2842_v72_) + (307)]))
            return _dafny.Set(coll95_)
        (d_2836_v66_)[index450_] = (iife181_()
        ) - ((d_2837_v67_) - (_dafny.Set({d_2833_v63_, d_2833_v63_})))
        index451_ = default__.safeIndex(678, (p1).length(0))
        (p1)[index451_] = True
        d_2843_v73_: _dafny.Array
        nw505_ = _dafny.Array(int(0), 3)
        d_2843_v73_ = nw505_
        index452_ = default__.safeIndex(557, (d_2843_v73_).length(0))
        (d_2843_v73_)[index452_] = 980
        index453_ = default__.safeIndex(678, (p1).length(0))
        index454_ = default__.safeIndex(557, (d_2843_v73_).length(0))
        rhs491_ = (d_2833_v63_) > (len(p0))
        rhs492_ = (d_2833_v63_) - (d_2833_v63_)
        lhs413_ = p1
        lhs414_ = default__.safeIndex(678, (p1).length(0))
        lhs415_ = d_2843_v73_
        lhs416_ = default__.safeIndex(557, (d_2843_v73_).length(0))
        lhs413_[lhs414_] = rhs491_
        lhs415_[lhs416_] = rhs492_
        index455_ = default__.safeIndex(557, (d_2843_v73_).length(0))
        (d_2843_v73_)[index455_] = len(default__.fm51(((d_2843_v73_)[default__.safeIndex(557, (d_2843_v73_).length(0))] if self.f28 else d_2833_v63_), (default__.fm14(globalState)).cardinality, globalState))
        r0 = self.f28
        d_2844_v74_: _dafny.Seq
        d_2844_v74_ = _dafny.SeqWithoutIsStrInference([(d_2843_v73_)[default__.safeIndex(557, (d_2843_v73_).length(0))]])
        d_2845_v75_: _dafny.MultiSet
        d_2845_v75_ = _dafny.MultiSet([(d_2843_v73_)[default__.safeIndex(557, (d_2843_v73_).length(0))], (d_2843_v73_)[default__.safeIndex(557, (d_2843_v73_).length(0))], len(d_2838_v69_), (d_2833_v63_ if self.f28 else d_2833_v63_), (len(d_2844_v74_)) + ((d_2843_v73_)[default__.safeIndex(557, (d_2843_v73_).length(0))])])
        r1 = d_2845_v75_
        r2 = (d_2840_v71_) + ((_dafny.SeqWithoutIsStrInference([p3])) + (d_2839_v70_))
        d_2846_v76_: _dafny.Array
        nw506_ = _dafny.Array(None, 9)
        nw506_[int(0)] = d_2843_v73_
        nw506_[int(1)] = d_2843_v73_
        nw506_[int(2)] = d_2843_v73_
        nw506_[int(3)] = d_2843_v73_
        nw506_[int(4)] = d_2843_v73_
        nw506_[int(5)] = d_2843_v73_
        nw506_[int(6)] = d_2843_v73_
        nw506_[int(7)] = d_2843_v73_
        nw506_[int(8)] = d_2843_v73_
        d_2846_v76_ = nw506_
        d_2847_v77_: D1
        d_2847_v77_ = D1_DC1(d_2846_v76_)
        d_2848_v78_: D1
        d_2848_v78_ = D1_DC3(d_2847_v77_)
        r3 = d_2848_v78_
        return r0, r1, r2, r3


class C12:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    def ctor__(self):
        pass
        pass

    def fm4(self, p0, p1, p2, globalState):
        return False

    def fm5(self, p0, p1, globalState):
        return (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtlcdyk"))), -440}))) - ((len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([True]))}))) * (len(_dafny.Map({D2_DC4(_dafny.CodePoint('l')): _dafny.CodePoint('u')}))))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_2849_v0_: str
        d_2849_v0_ = _dafny.CodePoint('f')
        d_2850_v1_: D2
        d_2850_v1_ = D2_DC4(d_2849_v0_)
        pat_let_tv64_ = p1
        pat_let_tv65_ = p1
        def lambda139_(source44_):
            if source44_.is_DC5:
                d_2851___mcc_h0_ = source44_.cf8
                d_2852_cf8_ = d_2851___mcc_h0_
                if d_2852_cf8_:
                    return not(not(d_2852_cf8_))
                elif True:
                    return d_2852_cf8_
            elif source44_.is_DC4:
                d_2853___mcc_h1_ = source44_.cf7
                d_2854_cf7_ = d_2853___mcc_h1_
                return not((not(True) if False else False))
            elif True:
                d_2855___mcc_h2_ = source44_.cf9
                d_2856_cf9_ = d_2855___mcc_h2_
                return (_dafny.Map({False: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([pat_let_tv64_ for d_2857_i0_ in range(default__.abs(812))]))).cardinality})) == (_dafny.Map({True: pat_let_tv65_}))

        (globalState).f13 = lambda139_(d_2850_v1_)
        d_2858_v2_: _dafny.Array
        nw507_ = _dafny.Array(int(0), 24)
        d_2858_v2_ = nw507_
        index456_ = default__.safeIndex(457, (d_2858_v2_).length(0))
        (d_2858_v2_)[index456_] = p1
        d_2859_v3_: bool
        d_2859_v3_ = True
        d_2860_v4_: _dafny.Set
        d_2860_v4_ = _dafny.Set({d_2859_v3_})
        index457_ = default__.safeIndex(457, (d_2858_v2_).length(0))
        (d_2858_v2_)[index457_] = (len(d_2860_v4_)) - (p1)
        if d_2859_v3_:
            (globalState).f21 = not(d_2859_v3_)
            d_2861_v5_: _dafny.Array
            d_2861_v5_ = d_2858_v2_
            d_2862_v6_: _dafny.Array
            nw508_ = _dafny.Array(None, 21)
            nw508_[int(0)] = d_2861_v5_
            nw508_[int(1)] = d_2858_v2_
            nw508_[int(2)] = d_2861_v5_
            nw508_[int(3)] = d_2861_v5_
            nw508_[int(4)] = d_2861_v5_
            nw508_[int(5)] = d_2861_v5_
            nw508_[int(6)] = d_2861_v5_
            nw508_[int(7)] = d_2861_v5_
            nw508_[int(8)] = d_2861_v5_
            nw508_[int(9)] = d_2861_v5_
            nw508_[int(10)] = d_2861_v5_
            nw508_[int(11)] = d_2861_v5_
            nw508_[int(12)] = d_2858_v2_
            nw508_[int(13)] = d_2861_v5_
            nw508_[int(14)] = d_2858_v2_
            nw508_[int(15)] = d_2861_v5_
            nw508_[int(16)] = d_2861_v5_
            nw508_[int(17)] = d_2861_v5_
            nw508_[int(18)] = d_2861_v5_
            nw508_[int(19)] = d_2861_v5_
            nw508_[int(20)] = d_2861_v5_
            d_2862_v6_ = nw508_
            d_2863_v7_: _dafny.Map
            d_2863_v7_ = _dafny.Map({(d_2862_v6_ if d_2859_v3_ else d_2862_v6_): p1})
            d_2863_v7_ = (d_2863_v7_).set(d_2862_v6_, (0) - (p1))
            d_2864_v8_: _dafny.Seq
            d_2864_v8_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_2865_v9_: _dafny.MultiSet
            d_2865_v9_ = _dafny.MultiSet([-703, (d_2864_v8_)[default__.safeIndex((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], len(d_2864_v8_))]])
            d_2866_v10_: _dafny.Seq
            d_2866_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p1, p1, p1, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]])])
            index458_ = default__.safeIndex(457, (d_2858_v2_).length(0))
            (d_2858_v2_)[index458_] = (0) - ((self).fm5((d_2865_v9_).cardinality, d_2866_v10_, globalState))
            (globalState).f23 = (d_2850_v1_).cf7
            d_2867_v11_: _dafny.Seq
            d_2867_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vq"))
            d_2868_v12_: _dafny.Map
            d_2868_v12_ = _dafny.Map({_dafny.Map({d_2859_v3_: (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]}): d_2867_v11_})
            source45_ = default__.fm6((d_2868_v12_) | (d_2868_v12_), ((self).fm5((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], _dafny.SeqWithoutIsStrInference([d_2865_v9_ for d_2869_i1_ in range(default__.abs(69))]), globalState)) > (len(d_2864_v8_)), 853, d_2859_v3_, globalState)
            if source45_.is_DC5:
                d_2870___mcc_h3_ = source45_.cf8
                d_2871_cf8_ = d_2870___mcc_h3_
                d_2872_v13_: _dafny.Seq
                d_2872_v13_ = _dafny.SeqWithoutIsStrInference([d_2858_v2_, d_2858_v2_, d_2858_v2_, d_2858_v2_, d_2858_v2_])
                d_2873_v14_: _dafny.Seq
                d_2873_v14_ = _dafny.SeqWithoutIsStrInference([(d_2872_v13_)[default__.safeIndex((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], len(d_2872_v13_))]])
                d_2874_v15_: C11
                nw509_ = C11()
                nw509_.ctor__(d_2873_v14_, (d_2867_v11_) <= (d_2867_v11_), d_2871_cf8_)
                d_2874_v15_ = nw509_
                index459_ = default__.safeIndex(457, (d_2858_v2_).length(0))
                rhs493_ = p1
                rhs494_ = d_2859_v3_
                rhs495_ = (d_2871_cf8_) == (d_2859_v3_)
                rhs496_ = p1
                lhs417_ = globalState
                lhs418_ = globalState
                lhs419_ = d_2858_v2_
                lhs420_ = default__.safeIndex(457, (d_2858_v2_).length(0))
                lhs417_.f1 = rhs493_
                r0 = rhs494_
                lhs418_.f13 = rhs495_
                lhs419_[lhs420_] = rhs496_
                d_2875_v16_: _dafny.Seq
                d_2875_v16_ = _dafny.SeqWithoutIsStrInference([not(not(d_2859_v3_)), d_2871_cf8_])
                d_2876_v17_: _dafny.Array
                nw510_ = _dafny.Array(None, 22)
                nw510_[int(0)] = d_2859_v3_
                nw510_[int(1)] = (d_2865_v9_).issubset(d_2865_v9_)
                nw510_[int(2)] = d_2871_cf8_
                nw510_[int(3)] = default__.fm26(p1, globalState)
                nw510_[int(4)] = False
                nw510_[int(5)] = d_2871_cf8_
                nw510_[int(6)] = d_2859_v3_
                nw510_[int(7)] = d_2871_cf8_
                nw510_[int(8)] = d_2871_cf8_
                nw510_[int(9)] = (default__.fm13(globalState) if d_2859_v3_ else d_2859_v3_)
                nw510_[int(10)] = d_2871_cf8_
                nw510_[int(11)] = d_2871_cf8_
                nw510_[int(12)] = d_2871_cf8_
                nw510_[int(13)] = d_2871_cf8_
                nw510_[int(14)] = d_2859_v3_
                nw510_[int(15)] = d_2871_cf8_
                nw510_[int(16)] = d_2871_cf8_
                nw510_[int(17)] = d_2859_v3_
                nw510_[int(18)] = d_2871_cf8_
                nw510_[int(19)] = (d_2875_v16_) != (d_2875_v16_)
                nw510_[int(20)] = (d_2871_cf8_ if d_2871_cf8_ else d_2859_v3_)
                nw510_[int(21)] = d_2871_cf8_
                d_2876_v17_ = nw510_
                index460_ = default__.safeIndex(481, (d_2876_v17_).length(0))
                (d_2876_v17_)[index460_] = (len(d_2867_v11_)) == ((default__.fm75((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], globalState)).cf55)
                index461_ = default__.safeIndex(481, (d_2876_v17_).length(0))
                (d_2876_v17_)[index461_] = d_2859_v3_
                d_2877_v18_: _dafny.Seq
                d_2877_v18_ = _dafny.SeqWithoutIsStrInference([d_2876_v17_])
                d_2878_v19_: _dafny.Array
                nw511_ = _dafny.Array(None, 16)
                nw511_[int(0)] = d_2876_v17_
                nw511_[int(1)] = d_2876_v17_
                nw511_[int(2)] = d_2876_v17_
                nw511_[int(3)] = d_2876_v17_
                nw511_[int(4)] = d_2876_v17_
                nw511_[int(5)] = d_2876_v17_
                nw511_[int(6)] = d_2876_v17_
                nw511_[int(7)] = d_2876_v17_
                nw511_[int(8)] = d_2876_v17_
                nw511_[int(9)] = d_2876_v17_
                nw511_[int(10)] = d_2876_v17_
                nw511_[int(11)] = d_2876_v17_
                nw511_[int(12)] = d_2876_v17_
                nw511_[int(13)] = d_2876_v17_
                nw511_[int(14)] = (d_2877_v18_)[default__.safeIndex(len(d_2867_v11_), len(d_2877_v18_))]
                nw511_[int(15)] = d_2876_v17_
                d_2878_v19_ = nw511_
                d_2879_v20_: D18
                d_2879_v20_ = D18_DC48(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pycci"))), d_2878_v19_, d_2859_v3_, d_2849_v0_)
                d_2880_v21_: _dafny.Map
                d_2880_v21_ = _dafny.Map({(d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]: d_2879_v20_})
                (globalState).f21 = default__.fm18((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], (0) - (len(d_2880_v21_)), p1, globalState)
            elif source45_.is_DC4:
                d_2881___mcc_h4_ = source45_.cf7
                d_2882_cf7_ = d_2881___mcc_h4_
                d_2883_v22_: _dafny.MultiSet
                d_2883_v22_ = _dafny.MultiSet([p0, p0])
                d_2884_v23_: _dafny.Map
                d_2884_v23_ = _dafny.Map({p1: d_2883_v22_})
                d_2884_v23_ = (d_2884_v23_).set(len(default__.fm51(p1, (d_2865_v9_).cardinality, globalState)), (d_2883_v22_) - (d_2883_v22_))
                d_2885_v24_: C4
                nw512_ = C4()
                nw512_.ctor__(d_2859_v3_, False, d_2859_v3_, d_2859_v3_)
                d_2885_v24_ = nw512_
                d_2886_v25_: _dafny.Seq
                d_2886_v25_ = d_2864_v8_
                d_2887_v26_: D15
                d_2887_v26_ = D15_DC38((d_2885_v24_).f37, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])
                (globalState).f17 = (((d_2886_v25_ if d_2859_v3_ else d_2886_v25_))).set(default__.safeIndex(p1, len(((d_2886_v25_ if d_2859_v3_ else d_2886_v25_)))), (0) - ((d_2887_v26_).cf55))
                d_2888_v27_: _dafny.Array
                def lambda140_(d_2889_i2_):
                    return False

                init81_ = lambda140_
                nw513_ = _dafny.Array(None, 23)
                for i0_81_ in range(nw513_.length(0)):
                    nw513_[i0_81_] = init81_(i0_81_)
                d_2888_v27_ = nw513_
                d_2890_v28_: _dafny.Seq
                d_2890_v28_ = _dafny.SeqWithoutIsStrInference([d_2888_v27_, d_2888_v27_])
                d_2890_v28_ = d_2890_v28_
            elif True:
                d_2891___mcc_h5_ = source45_.cf9
                d_2892_cf9_ = d_2891___mcc_h5_
                (globalState).f1 = (((d_2865_v9_).set((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], default__.abs(((d_2865_v9_)[p1] if (p1) in (d_2865_v9_) else (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])))) | (d_2865_v9_)).cardinality
                index462_ = default__.safeIndex(457, (d_2858_v2_).length(0))
                (d_2858_v2_)[index462_] = (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]
                d_2893_v29_: _dafny.Map
                d_2893_v29_ = _dafny.Map({(d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]: (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]})
                d_2893_v29_ = (d_2893_v29_).set((self).fm5((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], d_2866_v10_, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etsix"))))
                d_2894_v30_: _dafny.Array
                nw514_ = _dafny.Array(_dafny.Map({}), 11)
                d_2894_v30_ = nw514_
                index463_ = default__.safeIndex(384, (d_2894_v30_).length(0))
                (d_2894_v30_)[index463_] = d_2893_v29_
                index464_ = default__.safeIndex(384, (d_2894_v30_).length(0))
                (d_2894_v30_)[index464_] = (d_2893_v29_) | (d_2893_v29_)
        elif True:
            d_2895_v31_: _dafny.Map
            d_2895_v31_ = _dafny.Map({(0) - ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]): False})
            d_2896_v32_: _dafny.MultiSet
            d_2896_v32_ = _dafny.MultiSet([_dafny.Map({p1: d_2859_v3_}), d_2895_v31_, d_2895_v31_, d_2895_v31_, d_2895_v31_])
            d_2897_v34_: _dafny.Seq
            def iife184_():
                coll98_ = _dafny.Map()
                compr_98_: int
                for compr_98_ in _dafny.IntegerRange(-839, 771):
                    d_2898_v33_: int = compr_98_
                    if ((-839) <= (d_2898_v33_)) and ((d_2898_v33_) < (771)):
                        coll98_[default__.safeModuloInt(d_2898_v33_, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])] = d_2859_v3_
                return _dafny.Map(coll98_)
            d_2897_v34_ = _dafny.SeqWithoutIsStrInference([(iife184_()
            ).set(649, d_2859_v3_), d_2895_v31_])
            d_2899_v35_: _dafny.Seq
            d_2899_v35_ = _dafny.SeqWithoutIsStrInference([(d_2896_v32_) != (_dafny.MultiSet(d_2897_v34_))])
            d_2899_v35_ = d_2899_v35_
            d_2900_v36_: _dafny.Seq
            d_2900_v36_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2901_v37_: _dafny.Seq
            d_2901_v37_ = _dafny.SeqWithoutIsStrInference([d_2900_v36_, _dafny.SeqWithoutIsStrInference([(0) - (p1)]), d_2900_v36_])
            d_2901_v37_ = d_2901_v37_
            (globalState).f22 = d_2859_v3_
            d_2902_v38_: _dafny.Array
            nw515_ = _dafny.Array(False, 3)
            d_2902_v38_ = nw515_
            index465_ = default__.safeIndex(483, (d_2902_v38_).length(0))
            (d_2902_v38_)[index465_] = (p1) > ((0) - (p1))
            d_2903_v39_: _dafny.MultiSet
            d_2903_v39_ = _dafny.MultiSet([d_2859_v3_, d_2859_v3_])
            d_2904_v40_: _dafny.Map
            d_2904_v40_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2903_v39_, d_2903_v39_]): d_2859_v3_})
            d_2905_v41_: _dafny.Seq
            d_2905_v41_ = _dafny.SeqWithoutIsStrInference([d_2903_v39_, d_2903_v39_, d_2903_v39_, ((d_2903_v39_).set(d_2859_v3_, default__.abs((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]))).set(d_2859_v3_, default__.abs(675))])
            index466_ = default__.safeIndex(483, (d_2902_v38_).length(0))
            (d_2902_v38_)[index466_] = ((d_2904_v40_)[d_2905_v41_] if (d_2905_v41_) in (d_2904_v40_) else not(True))
            d_2906_v42_: _dafny.Map
            d_2906_v42_ = _dafny.Map({d_2859_v3_: d_2859_v3_})
            default__.m0((_dafny.Map({d_2859_v3_: (d_2902_v38_)[default__.safeIndex(483, (d_2902_v38_).length(0))]})) | (d_2906_v42_), (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], not ((d_2899_v35_)[default__.safeIndex(p1, len(d_2899_v35_))]) or (True), globalState)
        if default__.fm38(globalState):
            (globalState).f1 = ((0) - (default__.safeModuloInt(len(p0), p1)) if d_2859_v3_ else len((d_2860_v4_) | (d_2860_v4_)))
            (globalState).f1 = p1
            index467_ = default__.safeIndex(457, (d_2858_v2_).length(0))
            (d_2858_v2_)[index467_] = (0) - ((0) - ((715) * ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])))
            index468_ = default__.safeIndex(457, (d_2858_v2_).length(0))
            (d_2858_v2_)[index468_] = default__.fm27(len(d_2860_v4_), _dafny.CodePoint('r'), ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))] if d_2859_v3_ else -72), globalState)
            d_2907_v43_: _dafny.Seq
            d_2907_v43_ = _dafny.SeqWithoutIsStrInference([(d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]])
            d_2908_v44_: _dafny.Set
            d_2908_v44_ = _dafny.Set({d_2907_v43_})
            (globalState).f13 = (d_2908_v44_).isdisjoint(d_2908_v44_)
        elif True:
            d_2909_v45_: _dafny.Seq
            d_2909_v45_ = _dafny.SeqWithoutIsStrInference([d_2859_v3_])
            d_2910_v46_: _dafny.Seq
            d_2910_v46_ = _dafny.SeqWithoutIsStrInference([(d_2909_v45_)[default__.safeIndex((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], len(d_2909_v45_))]])
            d_2911_v47_: _dafny.Map
            d_2911_v47_ = _dafny.Map({d_2859_v3_: p1})
            d_2912_v48_: _dafny.Seq
            d_2912_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndogeuhko"))
            (globalState).f27 = (((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]) + (len(d_2910_v46_))) > (default__.safeModuloInt(len(d_2911_v47_), len(d_2912_v48_)))
            (globalState).f24 = (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]
            d_2910_v46_ = (d_2910_v46_) + (d_2910_v46_)
            d_2913_v49_: D20
            d_2913_v49_ = D20_DC51(p1, d_2859_v3_, d_2859_v3_)
            pat_let_tv66_ = d_2858_v2_
            pat_let_tv67_ = d_2858_v2_
            index469_ = default__.safeIndex(457, (d_2858_v2_).length(0))
            def iife185_(_pat_let43_0):
                def iife186_(d_2914_dt__update__tmp_h0_):
                    def iife187_(_pat_let44_0):
                        def iife188_(d_2915_dt__update_hcf75_h0_):
                            return D20_DC51(d_2915_dt__update_hcf75_h0_, (d_2914_dt__update__tmp_h0_).cf76, (d_2914_dt__update__tmp_h0_).cf77)
                        return iife188_(_pat_let44_0)
                    return iife187_((pat_let_tv67_)[default__.safeIndex(457, (pat_let_tv66_).length(0))])
                return iife186_(_pat_let43_0)
            (d_2858_v2_)[index469_] = ((iife185_(d_2913_v49_)).cf75) - (len(_dafny.SeqWithoutIsStrInference([(d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))] for d_2916_i3_ in range(default__.abs(500))])))
            d_2917_v50_: _dafny.Map
            d_2917_v50_ = _dafny.Map({d_2859_v3_: (d_2912_v48_) != (d_2912_v48_)})
            if ((d_2917_v50_)[not (d_2859_v3_) or (d_2859_v3_)] if (not (d_2859_v3_) or (d_2859_v3_)) in (d_2917_v50_) else d_2859_v3_):
                d_2918_v51_: _dafny.Seq
                d_2918_v51_ = _dafny.SeqWithoutIsStrInference([d_2910_v46_, d_2910_v46_])
                index470_ = default__.safeIndex(457, (d_2858_v2_).length(0))
                (d_2858_v2_)[index470_] = (0) - (len((_dafny.SeqWithoutIsStrInference([d_2910_v46_, d_2910_v46_, _dafny.SeqWithoutIsStrInference([d_2859_v3_, default__.fm26((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], globalState)]), _dafny.SeqWithoutIsStrInference([d_2859_v3_, d_2859_v3_, d_2859_v3_])])) + (d_2918_v51_)))
                d_2919_v52_: C9
                nw516_ = C9()
                nw516_.ctor__(True, d_2859_v3_)
                d_2919_v52_ = nw516_
                d_2920_v53_: C0
                nw517_ = C0()
                nw517_.ctor__()
                d_2920_v53_ = nw517_
                d_2921_v54_: _dafny.Map
                d_2921_v54_ = _dafny.Map({(p1) - (p1): ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]) + ((0) - (p1))})
                d_2921_v54_ = (d_2921_v54_).set(default__.safeModuloInt(len(d_2909_v45_), (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]), p1)
                (globalState).f22 = d_2859_v3_
            elif True:
                d_2922_v55_: _dafny.Array
                def lambda141_(d_2923_v0_):
                    def lambda142_(d_2924_i4_):
                        return d_2923_v0_

                    return lambda142_

                init82_ = lambda141_(d_2849_v0_)
                nw518_ = _dafny.Array(None, 19)
                for i0_82_ in range(nw518_.length(0)):
                    nw518_[i0_82_] = init82_(i0_82_)
                d_2922_v55_ = nw518_
                index471_ = default__.safeIndex(933, (d_2922_v55_).length(0))
                (d_2922_v55_)[index471_] = d_2849_v0_
                index472_ = default__.safeIndex(933, (d_2922_v55_).length(0))
                (d_2922_v55_)[index472_] = _dafny.CodePoint('g')
                (globalState).f22 = (p1) >= (((0) - ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])) - (p1))
                d_2925_v56_: _dafny.Array
                def lambda143_(d_2926_v50_, d_2927_v3_):
                    def lambda144_(d_2928_i5_):
                        return D17_DC45(d_2926_v50_, d_2927_v3_)

                    return lambda144_

                init83_ = lambda143_(d_2917_v50_, d_2859_v3_)
                nw519_ = _dafny.Array(None, 21)
                for i0_83_ in range(nw519_.length(0)):
                    nw519_[i0_83_] = init83_(i0_83_)
                d_2925_v56_ = nw519_
                d_2929_v57_: _dafny.Map
                d_2929_v57_ = _dafny.Map({p1: D22_DC58(d_2925_v56_)})
                d_2930_v58_: _dafny.Map
                d_2930_v58_ = _dafny.Map({True: d_2912_v48_})
                d_2929_v57_ = (d_2929_v57_).set((len(d_2930_v58_)) + (len(d_2911_v47_)), D22_DC58(d_2925_v56_))
                d_2931_v59_: _dafny.MultiSet
                d_2931_v59_ = _dafny.MultiSet([(d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]])
                rhs497_ = d_2931_v59_
                rhs498_ = (((0) - ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])) == (p1)) and (d_2859_v3_)
                lhs421_ = globalState
                d_2931_v59_ = rhs497_
                lhs421_.f13 = rhs498_
                d_2932_v60_: _dafny.Map
                d_2932_v60_ = _dafny.Map({p1: default__.fm42(p1, p1, globalState)})
                d_2933_v61_: D22
                d_2933_v61_ = D22_DC59(d_2859_v3_, d_2912_v48_, len(d_2932_v60_))
                d_2934_v62_: D22
                d_2934_v62_ = D22_DC61(d_2933_v61_)
                d_2934_v62_ = d_2934_v62_
        hi9_ = p1
        for d_2935_i6_ in range(((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))] if d_2859_v3_ else (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]), hi9_):
            (globalState).f13 = d_2859_v3_
            d_2936_v63_: _dafny.Set
            d_2936_v63_ = _dafny.Set({d_2858_v2_})
            d_2936_v63_ = d_2936_v63_
            d_2937_v64_: _dafny.MultiSet
            d_2937_v64_ = _dafny.MultiSet([d_2935_i6_, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]])
            d_2938_v65_: _dafny.Seq
            d_2938_v65_ = _dafny.SeqWithoutIsStrInference([d_2935_i6_])
            d_2939_v66_: _dafny.Seq
            d_2939_v66_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_2938_v65_)).cardinality, d_2935_i6_])
            d_2940_v67_: D4
            d_2940_v67_ = D4_DC10(d_2859_v3_, p1, (d_2939_v66_)[default__.safeIndex((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], len(d_2939_v66_))])
            d_2941_v68_: _dafny.Seq
            d_2941_v68_ = _dafny.SeqWithoutIsStrInference([not(False)])
            d_2942_v69_: _dafny.MultiSet
            d_2942_v69_ = _dafny.MultiSet([d_2859_v3_, d_2859_v3_])
            d_2943_v70_: _dafny.Array
            nw520_ = _dafny.Array(None, 17)
            nw520_[int(0)] = not((True) or (d_2859_v3_))
            nw520_[int(1)] = d_2859_v3_
            nw520_[int(2)] = d_2859_v3_
            nw520_[int(3)] = (p1) not in (d_2937_v64_)
            nw520_[int(4)] = d_2859_v3_
            nw520_[int(5)] = d_2859_v3_
            nw520_[int(6)] = d_2859_v3_
            nw520_[int(7)] = d_2859_v3_
            nw520_[int(8)] = not(True)
            nw520_[int(9)] = (d_2940_v67_).cf13
            nw520_[int(10)] = (d_2941_v68_)[default__.safeIndex(118, len(d_2941_v68_))]
            nw520_[int(11)] = d_2859_v3_
            nw520_[int(12)] = d_2859_v3_
            nw520_[int(13)] = d_2859_v3_
            nw520_[int(14)] = d_2859_v3_
            nw520_[int(15)] = d_2859_v3_
            nw520_[int(16)] = (d_2942_v69_).issubset(d_2942_v69_)
            d_2943_v70_ = nw520_
            index473_ = default__.safeIndex(482, (d_2943_v70_).length(0))
            (d_2943_v70_)[index473_] = d_2859_v3_
            index474_ = default__.safeIndex(482, (d_2943_v70_).length(0))
            (d_2943_v70_)[index474_] = (((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]) * (d_2935_i6_)) == (len(p0))
            d_2944_v71_: _dafny.Map
            d_2944_v71_ = _dafny.Map({len(_dafny.Map({(d_2943_v70_)[default__.safeIndex(482, (d_2943_v70_).length(0))]: p1})): p1})
            d_2945_v72_: _dafny.Map
            d_2945_v72_ = _dafny.Map({p1: d_2939_v66_})
            d_2946_v73_: _dafny.Seq
            d_2946_v73_ = _dafny.SeqWithoutIsStrInference([d_2944_v71_])
            index475_ = default__.safeIndex(482, (d_2943_v70_).length(0))
            rhs499_ = (_dafny.SeqWithoutIsStrInference([p1, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]])) <= (((((d_2945_v72_)[d_2935_i6_] if (d_2935_i6_) in (d_2945_v72_) else d_2938_v65_)).set(default__.safeIndex(p1, len(((d_2945_v72_)[d_2935_i6_] if (d_2935_i6_) in (d_2945_v72_) else d_2938_v65_))), 718)) + (d_2938_v65_))
            rhs500_ = (d_2946_v73_)[default__.safeIndex((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], len(d_2946_v73_))]
            rhs501_ = (0) - (d_2935_i6_)
            lhs422_ = d_2943_v70_
            lhs423_ = default__.safeIndex(482, (d_2943_v70_).length(0))
            lhs424_ = globalState
            lhs422_[lhs423_] = rhs499_
            d_2944_v71_ = rhs500_
            lhs424_.f1 = rhs501_
        hi10_ = ((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))]) - (p1)
        for d_2947_i7_ in range((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], hi10_):
            d_2948_v74_: _dafny.Seq
            d_2948_v74_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2949_v75_: _dafny.Map
            d_2949_v75_ = _dafny.Map({p1: d_2947_i7_})
            d_2950_v76_: _dafny.MultiSet
            d_2950_v76_ = _dafny.MultiSet([len(d_2949_v75_)])
            d_2951_v77_: _dafny.MultiSet
            d_2951_v77_ = _dafny.MultiSet([not(True)])
            (globalState).f17 = ((_dafny.SeqWithoutIsStrInference([(d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))] for d_2952_i8_ in range(default__.abs(279))])) + (d_2948_v74_) if d_2859_v3_ else (_dafny.SeqWithoutIsStrInference([p1, (d_2950_v76_).cardinality, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], (default__.fm57(d_2849_v0_, (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], d_2947_i7_, globalState)).cardinality, d_2947_i7_])) + (_dafny.SeqWithoutIsStrInference([(d_2951_v77_).cardinality, p1])))
            d_2953_v78_: _dafny.Array
            nw521_ = _dafny.Array(None, 28)
            d_2953_v78_ = nw521_
            d_2954_v79_: C9
            nw522_ = C9()
            nw522_.ctor__(d_2859_v3_, d_2859_v3_)
            d_2954_v79_ = nw522_
            index476_ = default__.safeIndex(864, (d_2953_v78_).length(0))
            (d_2953_v78_)[index476_] = d_2954_v79_
            index477_ = default__.safeIndex(864, (d_2953_v78_).length(0))
            (d_2953_v78_)[index477_] = d_2954_v79_
            d_2955_v80_: _dafny.Seq
            d_2955_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghbkroay"))
            index478_ = default__.safeIndex(457, (d_2858_v2_).length(0))
            (d_2858_v2_)[index478_] = default__.fm27((0) - (d_2947_i7_), default__.fm42(len(d_2955_v80_), d_2947_i7_, globalState), (0) - (default__.safeModuloInt((d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))], (d_2858_v2_)[default__.safeIndex(457, (d_2858_v2_).length(0))])), globalState)
            d_2956_v81_: _dafny.Array
            nw523_ = _dafny.Array(False, 10)
            d_2956_v81_ = nw523_
            rhs502_ = default__.fm8(globalState)
            rhs503_ = d_2947_i7_
            rhs504_ = p1
            rhs505_ = d_2956_v81_
            lhs425_ = globalState
            lhs426_ = globalState
            lhs427_ = globalState
            lhs425_.f24 = rhs502_
            lhs426_.f24 = rhs503_
            lhs427_.f24 = rhs504_
            d_2956_v81_ = rhs505_
        r0 = not(d_2859_v3_)
        d_2957_v82_: _dafny.Map
        d_2957_v82_ = _dafny.Map({p1: d_2859_v3_})
        r1 = (default__.fm45(globalState)) | (d_2957_v82_)
        return r0, r1

    def m4(self, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2958_v0_: bool
        d_2958_v0_ = True
        d_2959_v1_: _dafny.MultiSet
        d_2959_v1_ = _dafny.MultiSet([False, d_2958_v0_, d_2958_v0_, d_2958_v0_, d_2958_v0_])
        d_2960_v2_: _dafny.Seq
        d_2960_v2_ = _dafny.SeqWithoutIsStrInference([d_2959_v1_])
        d_2961_v3_: _dafny.Array
        nw524_ = _dafny.Array(None, 13)
        nw524_[int(0)] = d_2959_v1_
        nw524_[int(1)] = d_2959_v1_
        nw524_[int(2)] = (d_2960_v2_)[default__.safeIndex(-544, len(d_2960_v2_))]
        nw524_[int(3)] = d_2959_v1_
        nw524_[int(4)] = d_2959_v1_
        nw524_[int(5)] = d_2959_v1_
        nw524_[int(6)] = d_2959_v1_
        nw524_[int(7)] = d_2959_v1_
        nw524_[int(8)] = d_2959_v1_
        nw524_[int(9)] = d_2959_v1_
        nw524_[int(10)] = d_2959_v1_
        nw524_[int(11)] = d_2959_v1_
        nw524_[int(12)] = d_2959_v1_
        d_2961_v3_ = nw524_
        d_2962_v4_: D23
        d_2962_v4_ = D23_DC63(d_2961_v3_)
        source46_ = d_2962_v4_
        if source46_.is_DC63:
            d_2963___mcc_h0_ = source46_.cf92
            d_2964_cf92_ = d_2963___mcc_h0_
            d_2965_v5_: _dafny.Seq
            d_2965_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoytvc"))
            (globalState).f1 = len(d_2965_v5_)
            d_2966_v6_: int
            d_2966_v6_ = 462
            d_2967_v7_: _dafny.MultiSet
            d_2967_v7_ = _dafny.MultiSet([d_2966_v6_])
            d_2968_v8_: _dafny.Set
            d_2968_v8_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pywf")), _dafny.SeqWithoutIsStrInference([(d_2965_v5_)[default__.safeIndex(len(_dafny.Map({d_2958_v0_: -144})), len(d_2965_v5_))] for d_2969_i0_ in range(default__.abs(-495))])})
            d_2970_v9_: _dafny.Set
            d_2970_v9_ = d_2968_v8_
            d_2971_v10_: _dafny.Set
            d_2971_v10_ = _dafny.Set({d_2970_v9_, d_2970_v9_, d_2968_v8_, d_2970_v9_, d_2970_v9_})
            pat_let_tv68_ = d_2967_v7_
            pat_let_tv69_ = d_2971_v10_
            def iife189_(_pat_let45_0):
                def iife190_(d_2972_dt__update__tmp_h1_):
                    def iife191_(_pat_let46_0):
                        def iife192_(d_2973_dt__update_hcf12_h0_):
                            return D4_DC9(d_2973_dt__update_hcf12_h0_)
                        return iife192_(_pat_let46_0)
                    return iife191_((pat_let_tv68_).intersection(_dafny.MultiSet([len(pat_let_tv69_)])))
                return iife190_(_pat_let45_0)
            source47_ = iife189_(D4_DC9(d_2967_v7_))
            if source47_.is_DC10:
                d_2974___mcc_h8_ = source47_.cf13
                d_2975___mcc_h9_ = source47_.cf14
                d_2976___mcc_h10_ = source47_.cf15
                d_2977_cf15_ = d_2976___mcc_h10_
                d_2978_cf14_ = d_2975___mcc_h9_
                d_2979_cf13_ = d_2974___mcc_h8_
                d_2980_v11_: _dafny.Map
                d_2980_v11_ = _dafny.Map({d_2965_v5_: d_2979_cf13_})
                d_2980_v11_ = (d_2980_v11_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjfwtjqbb")) if d_2958_v0_ else d_2965_v5_), d_2979_cf13_)
                d_2966_v6_ = d_2978_cf14_
                d_2981_v12_: _dafny.Set
                d_2981_v12_ = _dafny.Set({d_2958_v0_})
                d_2982_v13_: _dafny.Seq
                d_2982_v13_ = _dafny.SeqWithoutIsStrInference([d_2981_v12_])
                (globalState).f24 = ((len((d_2982_v13_)[default__.safeIndex(d_2966_v6_, len(d_2982_v13_))])) * (d_2977_cf15_)) + (len(_dafny.Set({d_2979_cf13_, d_2979_cf13_, d_2979_cf13_})))
                d_2977_cf15_ = d_2977_cf15_
            elif source47_.is_DC9:
                d_2983___mcc_h11_ = source47_.cf12
                d_2984_cf12_ = d_2983___mcc_h11_
                d_2985_v14_: _dafny.Array
                def lambda145_(d_2986_v0_):
                    def lambda146_(d_2987_i1_):
                        return (_dafny.SeqWithoutIsStrInference([d_2986_v0_, d_2986_v0_, (D15_DC40(d_2986_v0_, not(d_2986_v0_))).cf58])) + (_dafny.SeqWithoutIsStrInference([d_2986_v0_]))

                    return lambda146_

                init84_ = lambda145_(d_2958_v0_)
                nw525_ = _dafny.Array(None, 25)
                for i0_84_ in range(nw525_.length(0)):
                    nw525_[i0_84_] = init84_(i0_84_)
                d_2985_v14_ = nw525_
                d_2988_v15_: _dafny.Seq
                d_2988_v15_ = _dafny.SeqWithoutIsStrInference([True])
                d_2989_v16_: _dafny.Seq
                d_2989_v16_ = _dafny.SeqWithoutIsStrInference([(d_2988_v15_)[default__.safeIndex(d_2966_v6_, len(d_2988_v15_))]])
                index479_ = default__.safeIndex(7, (d_2985_v14_).length(0))
                (d_2985_v14_)[index479_] = d_2989_v16_
                d_2990_v17_: _dafny.Array
                nw526_ = _dafny.Array(None, 6)
                nw526_[int(0)] = d_2965_v5_
                nw526_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkeimlfqp"))
                nw526_[int(2)] = d_2965_v5_
                nw526_[int(3)] = d_2965_v5_
                nw526_[int(4)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2991_i2_ in range(default__.abs(363))])).set(default__.safeIndex(d_2966_v6_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2992_i2_ in range(default__.abs(363))]))), (d_2965_v5_)[default__.safeIndex(d_2966_v6_, len(d_2965_v5_))])
                nw526_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2993_i3_ in range(default__.abs(-882))])
                d_2990_v17_ = nw526_
                d_2994_v18_: _dafny.Map
                d_2994_v18_ = _dafny.Map({d_2990_v17_: d_2989_v16_})
                index480_ = default__.safeIndex(7, (d_2985_v14_).length(0))
                (d_2985_v14_)[index480_] = ((d_2994_v18_)[d_2990_v17_] if (d_2990_v17_) in (d_2994_v18_) else (d_2988_v15_) + (d_2988_v15_))
                d_2995_v19_: _dafny.Array
                nw527_ = _dafny.Array(D15.default()(), 25)
                d_2995_v19_ = nw527_
                d_2996_v20_: D15
                d_2996_v20_ = D15_DC40(d_2958_v0_, d_2958_v0_)
                index481_ = default__.safeIndex(914, (d_2995_v19_).length(0))
                (d_2995_v19_)[index481_] = d_2996_v20_
                d_2997_v21_: _dafny.Array
                nw528_ = _dafny.Array(None, 5)
                nw528_[int(0)] = d_2966_v6_
                nw528_[int(1)] = d_2966_v6_
                nw528_[int(2)] = d_2966_v6_
                nw528_[int(3)] = -46
                nw528_[int(4)] = d_2966_v6_
                d_2997_v21_ = nw528_
                d_2998_v22_: _dafny.Set
                d_2998_v22_ = _dafny.Set({(d_2997_v21_ if d_2958_v0_ else d_2997_v21_)})
                d_2999_v23_: _dafny.Map
                d_2999_v23_ = _dafny.Map({75: d_2998_v22_})
                d_3000_v24_: _dafny.Map
                d_3000_v24_ = _dafny.Map({d_2958_v0_: d_2966_v6_})
                d_3001_v25_: _dafny.Map
                d_3001_v25_ = _dafny.Map({d_2958_v0_: ((d_2999_v23_)[len(d_3000_v24_)] if (len(d_3000_v24_)) in (d_2999_v23_) else d_2998_v22_)})
                index482_ = default__.safeIndex(914, (d_2995_v19_).length(0))
                rhs506_ = d_2958_v0_
                rhs507_ = (D23_DC64(default__.fm25(d_2966_v6_, d_2966_v6_, globalState), len(d_2965_v5_), True, d_2966_v6_, d_2966_v6_)).cf95
                rhs508_ = d_2996_v20_
                rhs509_ = (len(d_2965_v5_)) > (d_2966_v6_)
                rhs510_ = (((d_3001_v25_)[d_2958_v0_] if (d_2958_v0_) in (d_3001_v25_) else d_2998_v22_)).intersection(d_2998_v22_)
                lhs428_ = globalState
                lhs429_ = d_2995_v19_
                lhs430_ = default__.safeIndex(914, (d_2995_v19_).length(0))
                lhs431_ = globalState
                lhs428_.f22 = rhs506_
                d_2958_v0_ = rhs507_
                lhs429_[lhs430_] = rhs508_
                lhs431_.f21 = rhs509_
                d_2998_v22_ = rhs510_
                d_3002_v26_: C1
                nw529_ = C1()
                nw529_.ctor__(d_2958_v0_, d_2958_v0_)
                d_3002_v26_ = nw529_
                (globalState).f24 = d_2966_v6_
            elif True:
                d_3003___mcc_h12_ = source47_.cf16
                d_3004_cf16_ = d_3003___mcc_h12_
                d_3005_v27_: _dafny.Array
                def lambda147_(d_3006_v0_):
                    def lambda148_(d_3007_i4_):
                        return (d_3007_i4_) + (len(_dafny.Map({d_3006_v0_: _dafny.CodePoint('g')})))

                    return lambda148_

                init85_ = lambda147_(d_2958_v0_)
                nw530_ = _dafny.Array(None, 29)
                for i0_85_ in range(nw530_.length(0)):
                    nw530_[i0_85_] = init85_(i0_85_)
                d_3005_v27_ = nw530_
                d_3008_v28_: _dafny.Map
                d_3008_v28_ = _dafny.Map({d_2966_v6_: d_3005_v27_})
                d_3009_v29_: str
                d_3009_v29_ = _dafny.CodePoint('w')
                d_3010_v30_: _dafny.Map
                d_3010_v30_ = _dafny.Map({False: d_2966_v6_})
                d_3011_v31_: _dafny.Seq
                d_3011_v31_ = _dafny.SeqWithoutIsStrInference([d_2958_v0_, True, True])
                d_3012_v32_: _dafny.Seq
                d_3012_v32_ = _dafny.SeqWithoutIsStrInference([d_2966_v6_, d_2966_v6_, len(d_3011_v31_), d_2966_v6_, d_2966_v6_])
                d_3013_v33_: D4
                d_3013_v33_ = D4_DC10(d_2958_v0_, (d_2959_v1_).cardinality, d_2966_v6_)
                d_3014_v34_: _dafny.Set
                d_3014_v34_ = _dafny.Set({d_2958_v0_, d_2958_v0_})
                d_3015_v35_: _dafny.Array
                nw531_ = _dafny.Array(None, 24)
                nw531_[int(0)] = _dafny.Map({d_2958_v0_: default__.fm27(len(d_3008_v28_), d_3009_v29_, d_2966_v6_, globalState)})
                nw531_[int(1)] = (d_3010_v30_) | (d_3010_v30_)
                nw531_[int(2)] = d_3010_v30_
                nw531_[int(3)] = d_3010_v30_
                nw531_[int(4)] = d_3010_v30_
                nw531_[int(5)] = (d_3010_v30_) | (d_3010_v30_)
                nw531_[int(6)] = d_3010_v30_
                nw531_[int(7)] = d_3010_v30_
                nw531_[int(8)] = (d_3010_v30_) | (d_3010_v30_)
                nw531_[int(9)] = d_3010_v30_
                nw531_[int(10)] = (d_3010_v30_) | (d_3010_v30_)
                nw531_[int(11)] = (_dafny.Map({d_2958_v0_: (default__.fm57(d_3009_v29_, d_2966_v6_, d_2966_v6_, globalState)).cardinality})) | (d_3010_v30_)
                nw531_[int(12)] = d_3010_v30_
                nw531_[int(13)] = d_3010_v30_
                nw531_[int(14)] = d_3010_v30_
                nw531_[int(15)] = d_3010_v30_
                nw531_[int(16)] = default__.fm56(D4_DC10(d_2958_v0_, d_2966_v6_, d_2966_v6_), default__.fm13(globalState), len(d_3012_v32_), globalState)
                nw531_[int(17)] = d_3010_v30_
                nw531_[int(18)] = _dafny.Map({d_2958_v0_: d_2966_v6_})
                nw531_[int(19)] = d_3010_v30_
                nw531_[int(20)] = default__.fm56(d_3013_v33_, d_2958_v0_, len(_dafny.Map({len(d_3014_v34_): d_3013_v33_})), globalState)
                nw531_[int(21)] = d_3010_v30_
                nw531_[int(22)] = d_3010_v30_
                nw531_[int(23)] = (d_3010_v30_) | (d_3010_v30_)
                d_3015_v35_ = nw531_
                index483_ = default__.safeIndex(773, (d_3015_v35_).length(0))
                (d_3015_v35_)[index483_] = d_3010_v30_
                d_3016_v36_: _dafny.Map
                d_3016_v36_ = _dafny.Map({d_2958_v0_: d_3010_v30_})
                d_3017_v37_: _dafny.Map
                d_3017_v37_ = _dafny.Map({d_2966_v6_: ((d_3016_v36_)[d_2958_v0_] if (d_2958_v0_) in (d_3016_v36_) else d_3010_v30_)})
                index484_ = default__.safeIndex(773, (d_3015_v35_).length(0))
                (d_3015_v35_)[index484_] = ((d_3017_v37_)[(len(d_3011_v31_)) - (d_2966_v6_)] if ((len(d_3011_v31_)) - (d_2966_v6_)) in (d_3017_v37_) else d_3010_v30_)
                d_3018_v38_: _dafny.Array
                nw532_ = _dafny.Array(False, 4)
                d_3018_v38_ = nw532_
                d_3018_v38_ = d_3018_v38_
                d_3019_v39_: _dafny.Seq
                d_3019_v39_ = _dafny.SeqWithoutIsStrInference([d_2967_v7_])
                (globalState).f1 = (self).fm5(98, (d_3019_v39_).set(default__.safeIndex(d_2966_v6_, len(d_3019_v39_)), _dafny.MultiSet(d_3012_v32_)), globalState)
                d_3011_v31_ = d_3011_v31_
            rhs511_ = (d_2966_v6_) <= ((d_2966_v6_) * (d_2966_v6_))
            rhs512_ = d_2966_v6_
            lhs432_ = globalState
            lhs433_ = globalState
            lhs432_.f21 = rhs511_
            lhs433_.f24 = rhs512_
            (globalState).f1 = d_2966_v6_
        elif source46_.is_DC64:
            d_3020___mcc_h1_ = source46_.cf93
            d_3021___mcc_h2_ = source46_.cf94
            d_3022___mcc_h3_ = source46_.cf95
            d_3023___mcc_h4_ = source46_.cf96
            d_3024___mcc_h5_ = source46_.cf97
            d_3025_cf97_ = d_3024___mcc_h5_
            d_3026_cf96_ = d_3023___mcc_h4_
            d_3027_cf95_ = d_3022___mcc_h3_
            d_3028_cf94_ = d_3021___mcc_h2_
            d_3029_cf93_ = d_3020___mcc_h1_
            d_3030_v40_: C0
            nw533_ = C0()
            nw533_.ctor__()
            d_3030_v40_ = nw533_
            d_3031_v41_: _dafny.Seq
            d_3031_v41_ = _dafny.SeqWithoutIsStrInference([d_3030_v40_, d_3030_v40_])
            d_3032_v42_: _dafny.Map
            d_3032_v42_ = _dafny.Map({d_3031_v41_: d_2958_v0_})
            d_3032_v42_ = d_3032_v42_
            d_3033_v43_: _dafny.Map
            d_3033_v43_ = _dafny.Map({False: d_3027_cf95_})
            d_3033_v43_ = d_3033_v43_
            d_3034_v44_: _dafny.Seq
            d_3034_v44_ = _dafny.SeqWithoutIsStrInference([d_3027_cf95_, d_2958_v0_, d_2958_v0_])
            d_3035_v45_: _dafny.Seq
            d_3035_v45_ = _dafny.SeqWithoutIsStrInference([d_3034_v44_, ((d_3034_v44_).set(default__.safeIndex(840, len(d_3034_v44_)), d_2958_v0_)).set(default__.safeIndex(d_3025_cf97_, len((d_3034_v44_).set(default__.safeIndex(840, len(d_3034_v44_)), d_2958_v0_))), d_2958_v0_)])
            d_3036_v46_: _dafny.Map
            d_3036_v46_ = _dafny.Map({(d_3025_cf97_) + (len(d_3035_v45_)): d_3034_v44_})
            def iife193_():
                coll99_ = _dafny.Map()
                compr_99_: int
                for compr_99_ in _dafny.IntegerRange(-777, 412):
                    d_3037_v47_: int = compr_99_
                    if ((-777) <= (d_3037_v47_)) and ((d_3037_v47_) < (412)):
                        coll99_[(d_3037_v47_) - (len(d_3034_v44_))] = d_3034_v44_
                return _dafny.Map(coll99_)
            d_3036_v46_ = (iife193_()
            ) | ((d_3036_v46_) | (_dafny.Map({d_3025_cf97_: d_3034_v44_})))
            d_3038_v48_: _dafny.Seq
            d_3038_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgqb"))
            d_3039_v49_: _dafny.Map
            d_3039_v49_ = _dafny.Map({d_3027_cf95_: len(d_3038_v48_)})
            d_3040_v50_: _dafny.Map
            d_3040_v50_ = _dafny.Map({d_3026_cf96_: len(d_3039_v49_)})
            d_3041_v51_: _dafny.Map
            d_3041_v51_ = _dafny.Map({(d_3040_v50_) | (d_3040_v50_): d_3027_cf95_})
            d_3041_v51_ = (d_3041_v51_) | (d_3041_v51_)
        elif source46_.is_DC62:
            d_3042___mcc_h6_ = source46_.cf91
            d_3043_cf91_ = d_3042___mcc_h6_
            d_3044_v52_: _dafny.Map
            d_3044_v52_ = _dafny.Map({d_2958_v0_: d_2958_v0_})
            d_3045_v53_: _dafny.Array
            nw534_ = _dafny.Array(None, 1)
            nw534_[int(0)] = d_3044_v52_
            d_3045_v53_ = nw534_
            d_3046_v54_: _dafny.Map
            d_3046_v54_ = _dafny.Map({d_3045_v53_: d_2958_v0_})
            d_3046_v54_ = (d_3046_v54_).set(d_3045_v53_, d_2958_v0_)
            d_3047_v55_: _dafny.Array
            nw535_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_3047_v55_ = nw535_
            d_3048_v56_: _dafny.Array
            def lambda149_(d_3049_v0_):
                def lambda150_(d_3050_i5_):
                    return d_3049_v0_

                return lambda150_

            init86_ = lambda149_(d_2958_v0_)
            nw536_ = _dafny.Array(None, 12)
            for i0_86_ in range(nw536_.length(0)):
                nw536_[i0_86_] = init86_(i0_86_)
            d_3048_v56_ = nw536_
            index485_ = default__.safeIndex(904, (d_3047_v55_).length(0))
            (d_3047_v55_)[index485_] = d_3048_v56_
            index486_ = default__.safeIndex(904, (d_3047_v55_).length(0))
            (d_3047_v55_)[index486_] = d_3048_v56_
            (globalState).f27 = d_2958_v0_
            d_3051_v57_: _dafny.Array
            def lambda151_(d_3052_i6_):
                return default__.safeDivisionInt(d_3052_i6_, -622)

            init87_ = lambda151_
            nw537_ = _dafny.Array(None, 9)
            for i0_87_ in range(nw537_.length(0)):
                nw537_[i0_87_] = init87_(i0_87_)
            d_3051_v57_ = nw537_
            d_3053_v58_: int
            d_3053_v58_ = 796
            d_3054_v59_: _dafny.Seq
            d_3054_v59_ = _dafny.SeqWithoutIsStrInference([d_3053_v58_, d_3053_v58_, d_3053_v58_])
            index487_ = default__.safeIndex(705, (d_3051_v57_).length(0))
            (d_3051_v57_)[index487_] = ((_dafny.MultiSet(d_3054_v59_)).cardinality) * (d_3053_v58_)
            index488_ = default__.safeIndex(705, (d_3051_v57_).length(0))
            (d_3051_v57_)[index488_] = d_3053_v58_
        elif True:
            d_3055___mcc_h7_ = source46_.cf98
            d_3056_cf98_ = d_3055___mcc_h7_
            d_3057_v60_: _dafny.Array
            nw538_ = _dafny.Array(int(0), 4)
            d_3057_v60_ = nw538_
            d_3058_v61_: int
            d_3058_v61_ = -179
            index489_ = default__.safeIndex(34, (d_3057_v60_).length(0))
            (d_3057_v60_)[index489_] = d_3058_v61_
            index490_ = default__.safeIndex(34, (d_3057_v60_).length(0))
            (d_3057_v60_)[index490_] = ((d_2959_v1_).cardinality) + (default__.safeDivisionInt((0) - ((0) - (d_3058_v61_)), d_3058_v61_))
            (globalState).f1 = (d_3057_v60_)[default__.safeIndex(34, (d_3057_v60_).length(0))]
            d_3059_v62_: _dafny.Seq
            d_3059_v62_ = _dafny.SeqWithoutIsStrInference([(self).fm4(False, _dafny.CodePoint('t'), d_2958_v0_, globalState), d_2958_v0_])
            d_3059_v62_ = d_3059_v62_
            index491_ = default__.safeIndex(34, (d_3057_v60_).length(0))
            (d_3057_v60_)[index491_] = (d_3057_v60_)[default__.safeIndex(34, (d_3057_v60_).length(0))]
        d_3060_v63_: _dafny.Seq
        d_3060_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrrnydk"))
        d_3060_v63_ = d_3060_v63_
        d_3061_v64_: int
        d_3061_v64_ = 818
        d_3062_v65_: _dafny.Map
        d_3062_v65_ = _dafny.Map({d_2958_v0_: d_3061_v64_})
        (globalState).f22 = (len(d_3062_v65_)) != (d_3061_v64_)
        hi11_ = d_3061_v64_
        for d_3063_i7_ in range(d_3061_v64_, hi11_):
            d_3064_v66_: _dafny.Seq
            d_3064_v66_ = _dafny.SeqWithoutIsStrInference([d_3061_v64_])
            d_3065_v67_: _dafny.Map
            d_3065_v67_ = _dafny.Map({d_3060_v63_: d_3063_i7_})
            d_3066_v68_: _dafny.Array
            nw539_ = _dafny.Array(int(0), 14)
            d_3066_v68_ = nw539_
            d_3067_v69_: _dafny.Set
            d_3067_v69_ = _dafny.Set({d_3066_v68_, d_3066_v68_})
            d_3068_v70_: _dafny.MultiSet
            d_3068_v70_ = _dafny.MultiSet([d_3063_i7_, len(d_3067_v69_), d_3063_i7_, d_3063_i7_])
            d_3069_v71_: _dafny.Seq
            d_3069_v71_ = _dafny.SeqWithoutIsStrInference([d_3068_v70_, d_3068_v70_])
            (globalState).f22 = ((0) - ((0) - (default__.safeModuloInt(704, d_3063_i7_)))) != ((d_3064_v66_)[default__.safeIndex((self).fm5(((d_3065_v67_)[_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_3070_i8_ in range(default__.abs(736))])] if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_3071_i8_ in range(default__.abs(736))])) in (d_3065_v67_) else d_3061_v64_), d_3069_v71_, globalState), len(d_3064_v66_))])
            d_3072_v72_: _dafny.Array
            def lambda152_(d_3073_v1_):
                def lambda153_(d_3074_i9_):
                    return (d_3073_v1_).ispropersubset(d_3073_v1_)

                return lambda153_

            init88_ = lambda152_(d_2959_v1_)
            nw540_ = _dafny.Array(None, 9)
            for i0_88_ in range(nw540_.length(0)):
                nw540_[i0_88_] = init88_(i0_88_)
            d_3072_v72_ = nw540_
            index492_ = default__.safeIndex(107, (d_3072_v72_).length(0))
            (d_3072_v72_)[index492_] = d_2958_v0_
            index493_ = default__.safeIndex(107, (d_3072_v72_).length(0))
            (d_3072_v72_)[index493_] = (d_2958_v0_) or (d_2958_v0_)
            d_3075_v73_: C9
            nw541_ = C9()
            nw541_.ctor__((d_3072_v72_)[default__.safeIndex(107, (d_3072_v72_).length(0))], (d_3072_v72_)[default__.safeIndex(107, (d_3072_v72_).length(0))])
            d_3075_v73_ = nw541_
            d_3076_v74_: _dafny.Map
            d_3076_v74_ = _dafny.Map({d_3063_i7_: d_3075_v73_})
            d_3076_v74_ = (d_3076_v74_).set(d_3063_i7_, d_3075_v73_)
            d_3077_v75_: _dafny.Seq
            d_3077_v75_ = _dafny.SeqWithoutIsStrInference([d_3060_v63_])
            d_3078_v76_: _dafny.Seq
            d_3078_v76_ = _dafny.SeqWithoutIsStrInference([(d_3072_v72_)[default__.safeIndex(107, (d_3072_v72_).length(0))]])
            d_3079_v77_: _dafny.Map
            d_3079_v77_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_3080_i10_ in range(default__.abs(531))])])) + (d_3077_v75_)): d_3078_v76_})
            d_3081_v78_: _dafny.MultiSet
            d_3081_v78_ = _dafny.MultiSet([d_3078_v76_, d_3078_v76_])
            d_3082_v79_: D15
            d_3082_v79_ = D15_DC37(d_3081_v78_)
            d_3083_v80_: _dafny.Map
            d_3083_v80_ = _dafny.Map({default__.fm42(d_3063_i7_, d_3063_i7_, globalState): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3084_i11_ in range(default__.abs(-244))])})
            rhs513_ = default__.fm76((d_3060_v63_) == (d_3060_v63_), default__.fm29((d_3082_v79_).cf52, d_3061_v64_, globalState), (d_3072_v72_)[default__.safeIndex(107, (d_3072_v72_).length(0))], globalState)
            rhs514_ = _dafny.Map({(d_3072_v72_)[default__.safeIndex(107, (d_3072_v72_).length(0))]: (d_3061_v64_) + (len(d_3083_v80_))})
            rhs515_ = False
            rhs516_ = (-736) - (d_3063_i7_)
            lhs434_ = globalState
            lhs435_ = globalState
            d_3079_v77_ = rhs513_
            d_3062_v65_ = rhs514_
            lhs434_.f22 = rhs515_
            lhs435_.f1 = rhs516_
        hi12_ = d_3061_v64_
        for d_3085_i12_ in range((d_3061_v64_) * (d_3061_v64_), hi12_):
            d_3086_v81_: _dafny.MultiSet
            d_3086_v81_ = _dafny.MultiSet([default__.fm51(d_3085_i12_, d_3085_i12_, globalState)])
            d_3087_v82_: _dafny.Map
            d_3087_v82_ = _dafny.Map({False: d_3086_v81_})
            d_3088_v83_: _dafny.Map
            d_3088_v83_ = _dafny.Map({(d_3086_v81_).intersection(((d_3087_v82_)[d_2958_v0_] if (d_2958_v0_) in (d_3087_v82_) else (default__.fm20((D21_DC55(d_2958_v0_)).cf80, globalState)).set(d_3060_v63_, default__.abs(d_3085_i12_)))): d_2958_v0_})
            d_3089_v84_: _dafny.Seq
            d_3089_v84_ = _dafny.SeqWithoutIsStrInference([d_3060_v63_])
            d_3088_v83_ = (d_3088_v83_).set(_dafny.MultiSet((d_3089_v84_).set(default__.safeIndex(d_3085_i12_, len(d_3089_v84_)), d_3060_v63_)), not(d_2958_v0_))
            (globalState).f24 = d_3061_v64_
            d_3090_v85_: _dafny.Set
            d_3090_v85_ = _dafny.Set({len(d_3060_v63_)})
            d_3091_v86_: _dafny.Set
            d_3091_v86_ = _dafny.Set({default__.fm18(d_3061_v64_, d_3085_i12_, 600, globalState), d_2958_v0_})
            d_3092_v87_: bool
            d_3093_v88_: _dafny.Map
            out43_: bool
            out44_: _dafny.Map
            out43_, out44_ = (self).m3((d_3090_v85_) - (d_3090_v85_), (len(d_3091_v86_) if False else (_dafny.MultiSet([_dafny.Set({d_3061_v64_})])).cardinality), globalState)
            d_3092_v87_ = out43_
            d_3093_v88_ = out44_
            d_3094_v89_: C3
            nw542_ = C3()
            nw542_.ctor__(d_2958_v0_, d_3092_v87_, d_2958_v0_, (d_3085_i12_) - (d_3085_i12_))
            d_3094_v89_ = nw542_
        d_3061_v64_ = d_3061_v64_
        d_3095_v90_: _dafny.Array
        nw543_ = _dafny.Array(_dafny.Array(None, 0), 7)
        d_3095_v90_ = nw543_
        d_3096_v91_: D1
        d_3096_v91_ = D1_DC1(d_3095_v90_)
        d_3097_v92_: _dafny.Seq
        d_3097_v92_ = _dafny.SeqWithoutIsStrInference([d_3096_v91_, d_3096_v91_, d_3096_v91_, D1_DC1(d_3095_v90_)])
        pat_let_tv70_ = d_3095_v90_
        def iife194_(_pat_let47_0):
            def iife195_(d_3098_dt__update__tmp_h2_):
                def iife196_(_pat_let48_0):
                    def iife197_(d_3099_dt__update_hcf1_h0_):
                        return D1_DC1(d_3099_dt__update_hcf1_h0_)
                    return iife197_(_pat_let48_0)
                return iife196_(pat_let_tv70_)
            return iife195_(_pat_let47_0)
        r0 = ((d_3097_v92_).set(default__.safeIndex(d_3061_v64_, len(d_3097_v92_)), iife194_(d_3096_v91_))) + (d_3097_v92_)
        d_3100_v93_: _dafny.Array
        def lambda154_(d_3101_v64_):
            def lambda155_(d_3102_i13_):
                return (d_3102_i13_) + (d_3101_v64_)

            return lambda155_

        init89_ = lambda154_(d_3061_v64_)
        nw544_ = _dafny.Array(None, 26)
        for i0_89_ in range(nw544_.length(0)):
            nw544_[i0_89_] = init89_(i0_89_)
        d_3100_v93_ = nw544_
        r1 = (d_3100_v93_ if d_2958_v0_ else d_3100_v93_)
        return r0, r1


class C13:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        if True:
            return False
        elif True:
            return True

    def fm1(self, globalState):
        return _dafny.CodePoint('d')

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_3103_v0_: _dafny.Array
        nw545_ = _dafny.Array(int(0), 25)
        d_3103_v0_ = nw545_
        index494_ = default__.safeIndex(696, (d_3103_v0_).length(0))
        (d_3103_v0_)[index494_] = p0
        index495_ = default__.safeIndex(696, (d_3103_v0_).length(0))
        (d_3103_v0_)[index495_] = (p0) * (default__.safeDivisionInt(p0, p0))
        d_3104_v1_: bool
        d_3104_v1_ = False
        d_3105_v2_: _dafny.Seq
        d_3105_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_3106_v3_: _dafny.Map
        d_3106_v3_ = _dafny.Map({not(d_3104_v1_): (len(d_3105_v2_)) == (p0)})
        d_3106_v3_ = (d_3106_v3_).set(d_3104_v1_, ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) <= ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]))
        d_3107_v4_: _dafny.Seq
        d_3107_v4_ = _dafny.SeqWithoutIsStrInference([d_3105_v2_, d_3105_v2_, d_3105_v2_])
        d_3107_v4_ = (_dafny.SeqWithoutIsStrInference([d_3105_v2_, d_3105_v2_]) if not (d_3104_v1_) or (d_3104_v1_) else (d_3107_v4_) + (d_3107_v4_))
        if d_3104_v1_:
            d_3108_v5_: _dafny.Array
            nw546_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_3108_v5_ = nw546_
            d_3109_v6_: _dafny.Map
            d_3109_v6_ = _dafny.Map({d_3104_v1_: (d_3108_v5_ if (self).fm0(d_3104_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_3110_i0_ in range(default__.abs(108))]), globalState) else d_3108_v5_)})
            d_3109_v6_ = (d_3109_v6_).set(d_3104_v1_, d_3108_v5_)
            d_3111_v7_: _dafny.MultiSet
            d_3111_v7_ = _dafny.MultiSet([p0])
            d_3112_v8_: _dafny.Array
            d_3112_v8_ = d_3103_v0_
            d_3113_v9_: _dafny.Array
            nw547_ = _dafny.Array(None, 29)
            nw547_[int(0)] = d_3103_v0_
            nw547_[int(1)] = d_3103_v0_
            nw547_[int(2)] = d_3103_v0_
            nw547_[int(3)] = d_3103_v0_
            nw547_[int(4)] = d_3103_v0_
            nw547_[int(5)] = d_3103_v0_
            nw547_[int(6)] = d_3103_v0_
            nw547_[int(7)] = d_3103_v0_
            nw547_[int(8)] = d_3103_v0_
            nw547_[int(9)] = d_3103_v0_
            nw547_[int(10)] = d_3103_v0_
            nw547_[int(11)] = d_3103_v0_
            nw547_[int(12)] = d_3103_v0_
            nw547_[int(13)] = d_3103_v0_
            nw547_[int(14)] = d_3103_v0_
            nw547_[int(15)] = d_3103_v0_
            nw547_[int(16)] = d_3103_v0_
            nw547_[int(17)] = d_3103_v0_
            nw547_[int(18)] = d_3103_v0_
            nw547_[int(19)] = d_3103_v0_
            nw547_[int(20)] = d_3103_v0_
            nw547_[int(21)] = d_3103_v0_
            nw547_[int(22)] = d_3103_v0_
            nw547_[int(23)] = d_3103_v0_
            nw547_[int(24)] = (d_3112_v8_)
            nw547_[int(25)] = d_3103_v0_
            nw547_[int(26)] = d_3103_v0_
            nw547_[int(27)] = d_3103_v0_
            nw547_[int(28)] = d_3103_v0_
            d_3113_v9_ = nw547_
            d_3114_v10_: D1
            d_3114_v10_ = D1_DC1(d_3113_v9_)
            d_3115_v11_: str
            d_3115_v11_ = _dafny.CodePoint('h')
            d_3116_v12_: _dafny.Map
            d_3116_v12_ = _dafny.Map({d_3104_v1_: d_3115_v11_})
            d_3117_v13_: _dafny.Map
            d_3117_v13_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (len(d_3116_v12_))]): d_3113_v9_})
            d_3118_v14_: _dafny.Array
            nw548_ = _dafny.Array(None, 29)
            nw548_[int(0)] = d_3113_v9_
            nw548_[int(1)] = (d_3114_v10_).cf1
            nw548_[int(2)] = d_3113_v9_
            nw548_[int(3)] = d_3113_v9_
            nw548_[int(4)] = d_3113_v9_
            nw548_[int(5)] = d_3113_v9_
            nw548_[int(6)] = d_3113_v9_
            nw548_[int(7)] = d_3113_v9_
            nw548_[int(8)] = d_3113_v9_
            nw548_[int(9)] = d_3113_v9_
            nw548_[int(10)] = d_3113_v9_
            nw548_[int(11)] = d_3113_v9_
            nw548_[int(12)] = d_3113_v9_
            nw548_[int(13)] = d_3113_v9_
            nw548_[int(14)] = d_3113_v9_
            nw548_[int(15)] = ((d_3117_v13_)[d_3105_v2_] if (d_3105_v2_) in (d_3117_v13_) else d_3113_v9_)
            nw548_[int(16)] = d_3113_v9_
            nw548_[int(17)] = d_3113_v9_
            nw548_[int(18)] = d_3113_v9_
            nw548_[int(19)] = d_3113_v9_
            nw548_[int(20)] = d_3113_v9_
            nw548_[int(21)] = d_3113_v9_
            nw548_[int(22)] = d_3113_v9_
            nw548_[int(23)] = d_3113_v9_
            nw548_[int(24)] = d_3113_v9_
            nw548_[int(25)] = d_3113_v9_
            nw548_[int(26)] = d_3113_v9_
            nw548_[int(27)] = d_3113_v9_
            nw548_[int(28)] = d_3113_v9_
            d_3118_v14_ = nw548_
            index496_ = default__.safeIndex(821, (d_3118_v14_).length(0))
            (d_3118_v14_)[index496_] = d_3113_v9_
            d_3119_v15_: _dafny.Map
            d_3119_v15_ = _dafny.Map({d_3104_v1_: (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]})
            d_3120_v16_: _dafny.MultiSet
            d_3120_v16_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]])).set(default__.safeIndex((d_3105_v2_)[default__.safeIndex(len(d_3119_v15_), len(d_3105_v2_))], len(_dafny.SeqWithoutIsStrInference([(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]]))), (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]), d_3105_v2_, default__.fm2(_dafny.SeqWithoutIsStrInference([p0 for d_3121_i1_ in range(default__.abs(68))]), d_3104_v1_, globalState), d_3105_v2_])
            d_3122_v17_: _dafny.Map
            d_3122_v17_ = _dafny.Map({568: (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]})
            d_3123_v18_: _dafny.Map
            d_3123_v18_ = _dafny.Map({default__.safeModuloInt((d_3120_v16_).cardinality, len(d_3122_v17_)): d_3113_v9_})
            d_3124_v19_: _dafny.Seq
            d_3124_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "logqeu"))
            index497_ = default__.safeIndex(821, (d_3118_v14_).length(0))
            rhs517_ = d_3111_v7_
            rhs518_ = d_3104_v1_
            rhs519_ = (d_3104_v1_ if d_3104_v1_ else d_3104_v1_)
            rhs520_ = ((d_3123_v18_)[((d_3105_v2_)[default__.safeIndex(p0, len(d_3105_v2_))]) - (default__.fm3(d_3104_v1_, (self).fm0(d_3104_v1_, d_3124_v19_, globalState), p0, globalState))] if (((d_3105_v2_)[default__.safeIndex(p0, len(d_3105_v2_))]) - (default__.fm3(d_3104_v1_, (self).fm0(d_3104_v1_, d_3124_v19_, globalState), p0, globalState))) in (d_3123_v18_) else (d_3113_v9_ if False else d_3113_v9_))
            lhs436_ = globalState
            lhs437_ = globalState
            lhs438_ = d_3118_v14_
            lhs439_ = default__.safeIndex(821, (d_3118_v14_).length(0))
            d_3111_v7_ = rhs517_
            lhs436_.f22 = rhs518_
            lhs437_.f22 = rhs519_
            lhs438_[lhs439_] = rhs520_
            d_3125_v20_: _dafny.MultiSet
            d_3125_v20_ = _dafny.MultiSet([d_3106_v3_, _dafny.Map({d_3104_v1_: d_3104_v1_})])
            d_3126_v21_: _dafny.Seq
            d_3126_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_3106_v3_]), d_3125_v20_, d_3125_v20_])
            r3 = (d_3125_v20_).isdisjoint((d_3126_v21_)[default__.safeIndex(p0, len(d_3126_v21_))])
            if d_3104_v1_:
                (globalState).f27 = d_3104_v1_
                (globalState).f24 = ((0) - ((d_3105_v2_)[default__.safeIndex(p0, len(d_3105_v2_))])) * (p0)
                d_3127_v22_: D2
                d_3127_v22_ = D2_DC4((self).fm1(globalState))
                d_3128_v23_: _dafny.Array
                nw549_ = _dafny.Array(None, 5)
                nw549_[int(0)] = d_3115_v11_
                nw549_[int(1)] = (d_3127_v22_).cf7
                nw549_[int(2)] = d_3115_v11_
                nw549_[int(3)] = _dafny.CodePoint('v')
                nw549_[int(4)] = d_3115_v11_
                d_3128_v23_ = nw549_
                index498_ = default__.safeIndex(769, (d_3128_v23_).length(0))
                (d_3128_v23_)[index498_] = d_3115_v11_
                d_3129_v24_: _dafny.Seq
                d_3129_v24_ = _dafny.SeqWithoutIsStrInference([d_3113_v9_, (d_3114_v10_).cf1])
                pat_let_tv71_ = d_3118_v14_
                pat_let_tv72_ = d_3118_v14_
                pat_let_tv73_ = d_3118_v14_
                pat_let_tv74_ = d_3118_v14_
                d_3130_v25_: _dafny.Array
                nw550_ = _dafny.Array(None, 26)
                def iife198_(_pat_let49_0):
                    def iife199_(d_3131_dt__update__tmp_h0_):
                        def iife200_(_pat_let50_0):
                            def iife201_(d_3132_dt__update_hcf1_h0_):
                                return D1_DC1(d_3132_dt__update_hcf1_h0_)
                            return iife201_(_pat_let50_0)
                        return iife200_((pat_let_tv72_)[default__.safeIndex(821, (pat_let_tv71_).length(0))])
                    return iife199_(_pat_let49_0)
                nw550_[int(0)] = iife198_(D1_DC1(d_3113_v9_))
                nw550_[int(1)] = d_3114_v10_
                nw550_[int(2)] = D1_DC1(d_3113_v9_)
                nw550_[int(3)] = d_3114_v10_
                nw550_[int(4)] = d_3114_v10_
                nw550_[int(5)] = d_3114_v10_
                nw550_[int(6)] = d_3114_v10_
                nw550_[int(7)] = D1_DC1((d_3118_v14_)[default__.safeIndex(821, (d_3118_v14_).length(0))])
                nw550_[int(8)] = d_3114_v10_
                def iife202_(_pat_let51_0):
                    def iife203_(d_3133_dt__update__tmp_h1_):
                        def iife204_(_pat_let52_0):
                            def iife205_(d_3134_dt__update_hcf1_h1_):
                                return D1_DC1(d_3134_dt__update_hcf1_h1_)
                            return iife205_(_pat_let52_0)
                        return iife204_((pat_let_tv74_)[default__.safeIndex(821, (pat_let_tv73_).length(0))])
                    return iife203_(_pat_let51_0)
                nw550_[int(9)] = iife202_(D1_DC1((d_3118_v14_)[default__.safeIndex(821, (d_3118_v14_).length(0))]))
                nw550_[int(10)] = d_3114_v10_
                nw550_[int(11)] = D1_DC1((d_3129_v24_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_3115_v11_ for d_3135_i2_ in range(default__.abs(-36))])), len(d_3129_v24_))])
                nw550_[int(12)] = d_3114_v10_
                nw550_[int(13)] = D1_DC1((d_3118_v14_)[default__.safeIndex(821, (d_3118_v14_).length(0))])
                nw550_[int(14)] = d_3114_v10_
                nw550_[int(15)] = D1_DC1((d_3118_v14_)[default__.safeIndex(821, (d_3118_v14_).length(0))])
                nw550_[int(16)] = d_3114_v10_
                nw550_[int(17)] = d_3114_v10_
                nw550_[int(18)] = d_3114_v10_
                nw550_[int(19)] = d_3114_v10_
                nw550_[int(20)] = d_3114_v10_
                nw550_[int(21)] = d_3114_v10_
                nw550_[int(22)] = d_3114_v10_
                nw550_[int(23)] = D1_DC1((d_3118_v14_)[default__.safeIndex(821, (d_3118_v14_).length(0))])
                nw550_[int(24)] = d_3114_v10_
                nw550_[int(25)] = d_3114_v10_
                d_3130_v25_ = nw550_
                d_3136_v26_: _dafny.MultiSet
                d_3136_v26_ = _dafny.MultiSet([d_3130_v25_, d_3130_v25_, d_3130_v25_])
                d_3137_v27_: _dafny.Map
                d_3137_v27_ = _dafny.Map({d_3115_v11_: (d_3136_v26_).set(d_3130_v25_, default__.abs(p0))})
                index499_ = default__.safeIndex(769, (d_3128_v23_).length(0))
                rhs521_ = not (d_3104_v1_) or ((p0) >= (p0))
                rhs522_ = p0
                rhs523_ = d_3115_v11_
                rhs524_ = (((d_3137_v27_)[d_3115_v11_] if (d_3115_v11_) in (d_3137_v27_) else d_3136_v26_)).ispropersubset(d_3136_v26_)
                rhs525_ = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
                lhs440_ = globalState
                lhs441_ = globalState
                lhs442_ = d_3128_v23_
                lhs443_ = default__.safeIndex(769, (d_3128_v23_).length(0))
                lhs444_ = globalState
                lhs440_.f27 = rhs521_
                lhs441_.f24 = rhs522_
                lhs442_[lhs443_] = rhs523_
                d_3104_v1_ = rhs524_
                lhs444_.f1 = rhs525_
                d_3138_v28_: _dafny.Seq
                d_3138_v28_ = _dafny.SeqWithoutIsStrInference([d_3103_v0_])
                d_3139_v29_: _dafny.Seq
                d_3139_v29_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_])
                d_3140_v30_: C11
                nw551_ = C11()
                nw551_.ctor__(d_3138_v28_, ((_dafny.MultiSet([(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]])).cardinality) < (len(d_3139_v29_)), d_3104_v1_)
                d_3140_v30_ = nw551_
                d_3141_v31_: _dafny.Set
                d_3141_v31_ = _dafny.Set({d_3104_v1_, False, d_3104_v1_, d_3104_v1_})
                d_3142_v32_: D22
                d_3142_v32_ = D22_DC59(d_3104_v1_, d_3124_v19_, len(d_3141_v31_))
                d_3143_v33_: _dafny.Array
                nw552_ = _dafny.Array(None, 29)
                nw552_[int(0)] = d_3104_v1_
                nw552_[int(1)] = True
                nw552_[int(2)] = not(d_3104_v1_)
                nw552_[int(3)] = d_3104_v1_
                nw552_[int(4)] = False
                nw552_[int(5)] = d_3104_v1_
                nw552_[int(6)] = d_3104_v1_
                nw552_[int(7)] = False
                nw552_[int(8)] = d_3104_v1_
                nw552_[int(9)] = d_3104_v1_
                nw552_[int(10)] = d_3104_v1_
                nw552_[int(11)] = d_3104_v1_
                nw552_[int(12)] = False
                nw552_[int(13)] = d_3104_v1_
                nw552_[int(14)] = d_3104_v1_
                nw552_[int(15)] = d_3104_v1_
                nw552_[int(16)] = d_3104_v1_
                nw552_[int(17)] = d_3104_v1_
                nw552_[int(18)] = d_3104_v1_
                nw552_[int(19)] = (d_3142_v32_).cf83
                nw552_[int(20)] = False
                nw552_[int(21)] = d_3104_v1_
                nw552_[int(22)] = d_3104_v1_
                nw552_[int(23)] = d_3104_v1_
                nw552_[int(24)] = d_3104_v1_
                nw552_[int(25)] = d_3104_v1_
                nw552_[int(26)] = not(d_3104_v1_)
                nw552_[int(27)] = d_3104_v1_
                nw552_[int(28)] = d_3104_v1_
                d_3143_v33_ = nw552_
                d_3144_v34_: _dafny.Seq
                d_3144_v34_ = _dafny.SeqWithoutIsStrInference([d_3143_v33_])
                (globalState).f24 = (len(d_3144_v34_)) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
            elif True:
                (globalState).f21 = d_3104_v1_
                (globalState).f1 = p0
                def iife206_():
                    coll100_ = _dafny.Map()
                    compr_100_: int
                    for compr_100_ in ((D22_DC60(p0, d_3104_v1_, d_3104_v1_, d_3105_v2_)).cf89).Elements:
                        d_3145_v35_: int = compr_100_
                        if (d_3145_v35_) in ((D22_DC60(p0, d_3104_v1_, d_3104_v1_, d_3105_v2_)).cf89):
                            coll100_[(d_3145_v35_) * (p0)] = d_3124_v19_
                    return _dafny.Map(coll100_)
                d_3124_v19_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhjmmdhm"))) + (_dafny.SeqWithoutIsStrInference([(d_3124_v19_)[default__.safeIndex(len(iife206_()
), len(d_3124_v19_))] for d_3146_i3_ in range(default__.abs(578))]))
                d_3147_v36_: _dafny.Array
                nw553_ = _dafny.Array(None, 2)
                nw553_[int(0)] = d_3115_v11_
                nw553_[int(1)] = d_3115_v11_
                d_3147_v36_ = nw553_
                d_3148_v37_: _dafny.Map
                d_3148_v37_ = _dafny.Map({d_3147_v36_: (d_3104_v1_ if d_3104_v1_ else d_3104_v1_)})
                d_3148_v37_ = d_3148_v37_
                d_3149_v38_: C8
                nw554_ = C8()
                nw554_.ctor__()
                d_3149_v38_ = nw554_
            d_3150_v39_: _dafny.Set
            d_3150_v39_ = _dafny.Set({True})
            d_3151_v40_: _dafny.Map
            d_3151_v40_ = _dafny.Map({p0: d_3150_v39_})
            if ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) >= (len(((d_3151_v40_)[(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]] if ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) in (d_3151_v40_) else d_3150_v39_))):
                d_3152_v41_: T2
                nw555_ = C3()
                nw555_.ctor__(d_3104_v1_, d_3104_v1_, d_3104_v1_, (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
                d_3152_v41_ = nw555_
                d_3153_v42_: _dafny.Seq
                d_3153_v42_ = _dafny.SeqWithoutIsStrInference([d_3152_v41_, d_3152_v41_])
                d_3111_v7_ = (d_3111_v7_).set(len(d_3153_v42_), default__.abs(len((d_3124_v19_) + (d_3124_v19_))))
                index500_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                (d_3103_v0_)[index500_] = (p0) - (default__.safeModuloInt(p0, 269))
                d_3154_v43_: _dafny.MultiSet
                d_3154_v43_ = _dafny.MultiSet([d_3104_v1_, not(d_3104_v1_)])
                d_3155_v44_: D9
                d_3155_v44_ = D9_DC23(d_3154_v43_)
                d_3156_v45_: _dafny.MultiSet
                d_3156_v45_ = _dafny.MultiSet([d_3155_v44_, default__.fm77(globalState)])
                d_3157_v46_: _dafny.Map
                d_3157_v46_ = _dafny.Map({d_3104_v1_: d_3156_v45_})
                d_3158_v47_: _dafny.Map
                d_3158_v47_ = _dafny.Map({True: d_3157_v46_})
                d_3158_v47_ = (d_3158_v47_).set(d_3104_v1_, d_3157_v46_)
                d_3159_v48_: _dafny.Set
                d_3159_v48_ = _dafny.Set({853})
                (globalState).f1 = len((d_3159_v48_).intersection(d_3159_v48_))
                default__.m0(d_3106_v3_, len(d_3124_v19_), d_3104_v1_, globalState)
            elif True:
                (globalState).f1 = p0
                (globalState).f13 = default__.fm26(p0, globalState)
                r1 = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
                (globalState).f1 = (d_3105_v2_)[default__.safeIndex(p0, len(d_3105_v2_))]
                d_3106_v3_ = (d_3106_v3_).set(d_3104_v1_, not (d_3104_v1_) or (d_3104_v1_))
        elif True:
            (globalState).f1 = ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
            d_3160_v49_: _dafny.Seq
            d_3160_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            d_3161_v50_: _dafny.Array
            def lambda156_(d_3162_i4_):
                return D2_DC6(D2_DC4(_dafny.CodePoint('h')))

            init90_ = lambda156_
            nw556_ = _dafny.Array(None, 5)
            for i0_90_ in range(nw556_.length(0)):
                nw556_[i0_90_] = init90_(i0_90_)
            d_3161_v50_ = nw556_
            d_3163_v51_: C8
            nw557_ = C8()
            nw557_.ctor__()
            d_3163_v51_ = nw557_
            d_3164_v52_: _dafny.Map
            d_3164_v52_ = _dafny.Map({d_3103_v0_: p0})
            index501_ = default__.safeIndex(696, (d_3103_v0_).length(0))
            rhs526_ = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
            rhs527_ = d_3160_v49_
            rhs528_ = d_3161_v50_
            rhs529_ = d_3163_v51_
            rhs530_ = (len(d_3164_v52_)) + ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
            lhs445_ = globalState
            lhs446_ = d_3103_v0_
            lhs447_ = default__.safeIndex(696, (d_3103_v0_).length(0))
            lhs445_.f24 = rhs526_
            d_3160_v49_ = rhs527_
            d_3161_v50_ = rhs528_
            d_3163_v51_ = rhs529_
            lhs446_[lhs447_] = rhs530_
            d_3165_v53_: _dafny.Array
            def lambda157_(d_3166_v1_):
                def lambda158_(d_3167_i5_):
                    return d_3166_v1_

                return lambda158_

            init91_ = lambda157_(d_3104_v1_)
            nw558_ = _dafny.Array(None, 11)
            for i0_91_ in range(nw558_.length(0)):
                nw558_[i0_91_] = init91_(i0_91_)
            d_3165_v53_ = nw558_
            r0 = d_3165_v53_
            d_3168_v54_: _dafny.Array
            nw559_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_3168_v54_ = nw559_
            d_3169_v55_: D10
            d_3169_v55_ = D10_DC28(d_3106_v3_)
            d_3170_v56_: _dafny.Array
            nw560_ = _dafny.Array(None, 14)
            nw560_[int(0)] = d_3169_v55_
            nw560_[int(1)] = d_3169_v55_
            nw560_[int(2)] = d_3169_v55_
            nw560_[int(3)] = d_3169_v55_
            nw560_[int(4)] = d_3169_v55_
            nw560_[int(5)] = d_3169_v55_
            nw560_[int(6)] = d_3169_v55_
            nw560_[int(7)] = d_3169_v55_
            nw560_[int(8)] = d_3169_v55_
            nw560_[int(9)] = d_3169_v55_
            nw560_[int(10)] = d_3169_v55_
            nw560_[int(11)] = d_3169_v55_
            nw560_[int(12)] = d_3169_v55_
            nw560_[int(13)] = D10_DC28(d_3106_v3_)
            d_3170_v56_ = nw560_
            index502_ = default__.safeIndex(136, (d_3168_v54_).length(0))
            (d_3168_v54_)[index502_] = d_3170_v56_
            d_3171_v57_: D21
            d_3171_v57_ = D21_DC55(d_3104_v1_)
            d_3172_v58_: _dafny.Map
            d_3172_v58_ = _dafny.Map({d_3171_v57_: len(_dafny.Map({(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]: p0}))})
            index503_ = default__.safeIndex(136, (d_3168_v54_).length(0))
            rhs531_ = d_3104_v1_
            rhs532_ = d_3170_v56_
            rhs533_ = d_3172_v58_
            lhs448_ = globalState
            lhs449_ = d_3168_v54_
            lhs450_ = default__.safeIndex(136, (d_3168_v54_).length(0))
            lhs448_.f27 = rhs531_
            lhs449_[lhs450_] = rhs532_
            d_3172_v58_ = rhs533_
            index504_ = default__.safeIndex(147, (d_3165_v53_).length(0))
            (d_3165_v53_)[index504_] = d_3104_v1_
            d_3173_v59_: _dafny.Map
            d_3173_v59_ = _dafny.Map({default__.fm12(p0, True, globalState): d_3104_v1_})
            d_3174_v60_: _dafny.Seq
            d_3174_v60_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_])
            d_3175_v61_: _dafny.Seq
            d_3175_v61_ = d_3174_v60_
            index505_ = default__.safeIndex(147, (d_3165_v53_).length(0))
            (d_3165_v53_)[index505_] = not((((d_3173_v59_)[(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]] if ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) in (d_3173_v59_) else not(d_3104_v1_))) in ((d_3175_v61_)))
        if d_3104_v1_:
            d_3176_v62_: str
            d_3176_v62_ = _dafny.CodePoint('o')
            (globalState).f23 = d_3176_v62_
            d_3177_v63_: D7
            d_3177_v63_ = D7_DC18(d_3104_v1_)
            d_3178_v64_: _dafny.MultiSet
            d_3178_v64_ = _dafny.MultiSet([(d_3177_v63_).cf23])
            d_3179_v65_: _dafny.Seq
            d_3179_v65_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_, d_3104_v1_])
            d_3180_v66_: _dafny.MultiSet
            d_3180_v66_ = _dafny.MultiSet([p0, 256, p0, (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]])
            d_3181_v67_: _dafny.Map
            d_3181_v67_ = _dafny.Map({d_3104_v1_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygscpwya")))})
            d_3182_v68_: T2
            nw561_ = C3()
            nw561_.ctor__(d_3104_v1_, d_3104_v1_, (d_3179_v65_)[default__.safeIndex(((d_3180_v66_)[len(_dafny.Set({d_3181_v67_, d_3181_v67_}))] if (len(_dafny.Set({d_3181_v67_, d_3181_v67_}))) in (d_3180_v66_) else p0), len(d_3179_v65_))], p0)
            d_3182_v68_ = nw561_
            d_3183_v69_: _dafny.Map
            d_3183_v69_ = _dafny.Map({((0) - (p0)) + ((d_3178_v64_).cardinality): d_3182_v68_})
            d_3183_v69_ = ((d_3183_v69_) | (d_3183_v69_)) | (d_3183_v69_)
            (globalState).f21 = d_3104_v1_
            d_3184_v70_: C3
            nw562_ = C3()
            nw562_.ctor__(d_3104_v1_, d_3104_v1_, d_3104_v1_, default__.safeDivisionInt((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], 245))
            d_3184_v70_ = nw562_
            (globalState).f21 = d_3104_v1_
        elif True:
            d_3185_v71_: _dafny.Seq
            d_3185_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrgdvbry"))
            d_3186_v72_: _dafny.Map
            d_3186_v72_ = _dafny.Map({d_3185_v71_: True})
            d_3187_v73_: D17
            d_3187_v73_ = D17_DC43((_dafny.Map({d_3185_v71_: d_3104_v1_})) | (d_3186_v72_))
            source48_ = d_3187_v73_
            if source48_.is_DC44:
                d_3188___mcc_h0_ = source48_.cf62
                d_3189___mcc_h1_ = source48_.cf63
                d_3190_cf63_ = d_3189___mcc_h1_
                d_3191_cf62_ = d_3188___mcc_h0_
                (globalState).f24 = p0
                d_3192_v74_: C9
                nw563_ = C9()
                nw563_.ctor__(d_3104_v1_, d_3191_cf62_)
                d_3192_v74_ = nw563_
                d_3193_v75_: str
                d_3193_v75_ = _dafny.CodePoint('j')
                d_3194_v76_: C5
                nw564_ = C5()
                nw564_.ctor__(not(default__.fm18((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], p0, len(d_3185_v71_), globalState)), d_3193_v75_, d_3191_cf62_, default__.fm13(globalState))
                d_3194_v76_ = nw564_
                d_3195_v77_: D24
                d_3195_v77_ = D24_DC66(d_3194_v76_)
                d_3195_v77_ = d_3195_v77_
                (globalState).f17 = d_3105_v2_
            elif source48_.is_DC45:
                d_3196___mcc_h2_ = source48_.cf64
                d_3197___mcc_h3_ = source48_.cf65
                d_3198_cf65_ = d_3197___mcc_h3_
                d_3199_cf64_ = d_3196___mcc_h2_
                d_3200_v78_: _dafny.Seq
                d_3200_v78_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_])
                d_3201_v79_: T2
                nw565_ = C3()
                nw565_.ctor__(d_3104_v1_, (d_3200_v78_)[default__.safeIndex(391, len(d_3200_v78_))], d_3104_v1_, 629)
                d_3201_v79_ = nw565_
                d_3201_v79_ = d_3201_v79_
                index506_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                (d_3103_v0_)[index506_] = ((d_3201_v79_).f39) + (p0)
                d_3202_v80_: _dafny.Seq
                d_3202_v80_ = _dafny.SeqWithoutIsStrInference([d_3103_v0_])
                d_3203_v81_: C11
                nw566_ = C11()
                nw566_.ctor__(d_3202_v80_, d_3104_v1_, d_3198_cf65_)
                d_3203_v81_ = nw566_
                d_3204_v82_: _dafny.Array
                nw567_ = _dafny.Array(_dafny.Map({}), 17)
                d_3204_v82_ = nw567_
                d_3205_v83_: str
                d_3205_v83_ = _dafny.CodePoint('y')
                d_3206_v84_: T0
                nw568_ = C7()
                nw568_.ctor__(d_3104_v1_, d_3104_v1_)
                d_3206_v84_ = nw568_
                d_3207_v85_: D8
                d_3207_v85_ = D8_DC20(d_3206_v84_)
                d_3208_v86_: D9
                d_3208_v86_ = D9_DC24(d_3205_v83_, d_3104_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anqepprbb"))), d_3207_v85_, (d_3206_v84_).f29)
                d_3209_v87_: C0
                nw569_ = C0()
                nw569_.ctor__()
                d_3209_v87_ = nw569_
                d_3210_v88_: _dafny.Map
                d_3210_v88_ = _dafny.Map({not((d_3208_v86_).cf30): d_3209_v87_})
                index507_ = default__.safeIndex(657, (d_3204_v82_).length(0))
                (d_3204_v82_)[index507_] = d_3210_v88_
                d_3211_v89_: _dafny.Array
                nw570_ = _dafny.Array(_dafny.Set({}), 19)
                d_3211_v89_ = nw570_
                index508_ = default__.safeIndex(657, (d_3204_v82_).length(0))
                rhs534_ = not (d_3104_v1_) or (d_3104_v1_)
                rhs535_ = d_3210_v88_
                rhs536_ = d_3211_v89_
                lhs451_ = globalState
                lhs452_ = d_3204_v82_
                lhs453_ = default__.safeIndex(657, (d_3204_v82_).length(0))
                lhs451_.f13 = rhs534_
                lhs452_[lhs453_] = rhs535_
                d_3211_v89_ = rhs536_
            elif source48_.is_DC43:
                d_3212___mcc_h4_ = source48_.cf61
                d_3213_cf61_ = d_3212___mcc_h4_
                (globalState).f24 = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
                (globalState).f22 = not(d_3104_v1_)
                index509_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                (d_3103_v0_)[index509_] = p0
                d_3214_v90_: _dafny.Seq
                d_3214_v90_ = _dafny.SeqWithoutIsStrInference([False, d_3104_v1_])
                default__.m0(default__.fm59(True, d_3214_v90_, globalState), (len(d_3185_v71_)) + (p0), (len(d_3106_v3_)) >= (p0), globalState)
            elif True:
                d_3215___mcc_h5_ = source48_.cf66
                d_3216_cf66_ = d_3215___mcc_h5_
                (globalState).f21 = (default__.safeDivisionInt(len(d_3105_v2_), (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])) == ((0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]))
                (globalState).f22 = d_3104_v1_
                d_3217_v91_: _dafny.Array
                nw571_ = _dafny.Array(False, 18)
                d_3217_v91_ = nw571_
                d_3218_v92_: _dafny.Seq
                d_3218_v92_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_, d_3104_v1_])
                index510_ = default__.safeIndex(908, (d_3217_v91_).length(0))
                (d_3217_v91_)[index510_] = (d_3218_v92_)[default__.safeIndex((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], len(d_3218_v92_))]
                index511_ = default__.safeIndex(908, (d_3217_v91_).length(0))
                (d_3217_v91_)[index511_] = not (d_3104_v1_) or (default__.fm10(d_3104_v1_, p0, d_3104_v1_, globalState))
                (globalState).f24 = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
            d_3219_v93_: _dafny.Array
            def lambda159_(d_3220_i6_):
                return True

            init92_ = lambda159_
            nw572_ = _dafny.Array(None, 25)
            for i0_92_ in range(nw572_.length(0)):
                nw572_[i0_92_] = init92_(i0_92_)
            d_3219_v93_ = nw572_
            d_3221_v94_: _dafny.Set
            d_3221_v94_ = _dafny.Set({d_3219_v93_})
            (globalState).f21 = (d_3221_v94_) == (d_3221_v94_)
            index512_ = default__.safeIndex(696, (d_3103_v0_).length(0))
            (d_3103_v0_)[index512_] = default__.safeDivisionInt((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], p0)
            d_3222_v95_: T0
            nw573_ = C3()
            nw573_.ctor__((self).fm0(d_3104_v1_, d_3185_v71_, globalState), d_3104_v1_, d_3104_v1_, (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
            d_3222_v95_ = nw573_
            d_3223_v96_: _dafny.Map
            d_3223_v96_ = _dafny.Map({d_3104_v1_: d_3222_v95_})
            d_3224_v97_: _dafny.Set
            d_3224_v97_ = _dafny.Set({default__.fm8(globalState)})
            d_3223_v96_ = (d_3223_v96_).set((d_3224_v97_).issubset(d_3224_v97_), d_3222_v95_)
            d_3225_v98_: C10
            nw574_ = C10()
            nw574_.ctor__(d_3185_v71_, default__.safeDivisionInt(len(d_3106_v3_), p0))
            d_3225_v98_ = nw574_
        d_3226_v99_: str
        d_3226_v99_ = _dafny.CodePoint('n')
        d_3227_v100_: D7
        d_3227_v100_ = D7_DC18((d_3226_v99_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rl"))))
        source49_ = d_3227_v100_
        if source49_.is_DC17:
            d_3228___mcc_h6_ = source49_.cf22
            d_3229_cf22_ = d_3228___mcc_h6_
            index513_ = default__.safeIndex(696, (d_3103_v0_).length(0))
            (d_3103_v0_)[index513_] = (0) - (default__.fm3(d_3104_v1_, d_3104_v1_, default__.safeModuloInt(p0, len(d_3105_v2_)), globalState))
            d_3230_v101_: _dafny.Seq
            d_3230_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sf"))
            (globalState).f13 = ((d_3230_v101_) == (d_3230_v101_) if (d_3229_cf22_ if not(d_3104_v1_) else d_3229_cf22_) else d_3229_cf22_)
            d_3229_cf22_ = (p0) != ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
            (globalState).f21 = d_3229_cf22_
        elif source49_.is_DC18:
            d_3231___mcc_h7_ = source49_.cf23
            d_3232_cf23_ = d_3231___mcc_h7_
            d_3233_v102_: _dafny.Seq
            d_3233_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vep"))
            d_3233_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqgmtr"))
            d_3234_v103_: _dafny.Array
            nw575_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_3234_v103_ = nw575_
            d_3235_v104_: D18
            d_3235_v104_ = D18_DC48(d_3104_v1_, len(d_3233_v102_), d_3234_v103_, d_3232_cf23_, d_3226_v99_)
            d_3236_v105_: _dafny.Array
            nw576_ = _dafny.Array(None, 14)
            nw576_[int(0)] = d_3226_v99_
            nw576_[int(1)] = (d_3235_v104_).cf72
            nw576_[int(2)] = d_3226_v99_
            nw576_[int(3)] = _dafny.CodePoint('d')
            nw576_[int(4)] = d_3226_v99_
            nw576_[int(5)] = d_3226_v99_
            nw576_[int(6)] = d_3226_v99_
            nw576_[int(7)] = d_3226_v99_
            nw576_[int(8)] = d_3226_v99_
            nw576_[int(9)] = d_3226_v99_
            nw576_[int(10)] = d_3226_v99_
            nw576_[int(11)] = d_3226_v99_
            nw576_[int(12)] = d_3226_v99_
            nw576_[int(13)] = d_3226_v99_
            d_3236_v105_ = nw576_
            d_3237_v106_: _dafny.MultiSet
            d_3237_v106_ = _dafny.MultiSet([d_3236_v105_, d_3236_v105_])
            d_3238_v107_: _dafny.Seq
            d_3238_v107_ = _dafny.SeqWithoutIsStrInference([d_3237_v106_])
            if ((d_3237_v106_).intersection(d_3237_v106_)).ispropersubset((d_3238_v107_)[default__.safeIndex((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], len(d_3238_v107_))]):
                (globalState).f13 = d_3104_v1_
                d_3239_v108_: _dafny.Map
                d_3239_v108_ = _dafny.Map({d_3232_cf23_: len(_dafny.SeqWithoutIsStrInference([d_3226_v99_ for d_3240_i7_ in range(default__.abs(-658))]))})
                d_3241_v109_: _dafny.Map
                d_3241_v109_ = _dafny.Map({677: d_3239_v108_})
                d_3242_v110_: C7
                nw577_ = C7()
                nw577_.ctor__(d_3104_v1_, (len(d_3241_v109_)) <= ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]))
                d_3242_v110_ = nw577_
                d_3243_v111_: _dafny.Array
                nw578_ = _dafny.Array(_dafny.MultiSet({}), 11)
                d_3243_v111_ = nw578_
                d_3244_v112_: _dafny.Array
                nw579_ = _dafny.Array(None, 8)
                nw579_[int(0)] = d_3104_v1_
                nw579_[int(1)] = d_3232_cf23_
                nw579_[int(2)] = d_3232_cf23_
                nw579_[int(3)] = d_3104_v1_
                nw579_[int(4)] = d_3104_v1_
                nw579_[int(5)] = not(d_3104_v1_)
                nw579_[int(6)] = True
                nw579_[int(7)] = d_3104_v1_
                d_3244_v112_ = nw579_
                index514_ = default__.safeIndex(782, (d_3243_v111_).length(0))
                (d_3243_v111_)[index514_] = (_dafny.MultiSet([d_3244_v112_, d_3244_v112_, (D5_DC12(d_3244_v112_)).cf17, d_3244_v112_, d_3244_v112_])).intersection(_dafny.MultiSet([d_3244_v112_, d_3244_v112_, d_3244_v112_]))
                d_3245_v113_: _dafny.MultiSet
                d_3245_v113_ = _dafny.MultiSet([d_3244_v112_, d_3244_v112_])
                index515_ = default__.safeIndex(782, (d_3243_v111_).length(0))
                (d_3243_v111_)[index515_] = ((d_3245_v113_).intersection((d_3245_v113_).set(d_3244_v112_, default__.abs((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])))).set(d_3244_v112_, default__.abs((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]))
                index516_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                rhs537_ = (0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
                rhs538_ = 301
                rhs539_ = d_3104_v1_
                lhs454_ = d_3103_v0_
                lhs455_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                lhs456_ = globalState
                lhs454_[lhs455_] = rhs537_
                r1 = rhs538_
                lhs456_.f21 = rhs539_
                index517_ = default__.safeIndex(444, (d_3244_v112_).length(0))
                (d_3244_v112_)[index517_] = (d_3232_cf23_ if d_3232_cf23_ else d_3232_cf23_)
                index518_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                index519_ = default__.safeIndex(444, (d_3244_v112_).length(0))
                rhs540_ = -767
                rhs541_ = d_3104_v1_
                lhs457_ = d_3103_v0_
                lhs458_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                lhs459_ = d_3244_v112_
                lhs460_ = default__.safeIndex(444, (d_3244_v112_).length(0))
                lhs457_[lhs458_] = rhs540_
                lhs459_[lhs460_] = rhs541_
            elif True:
                d_3233_v102_ = (d_3233_v102_) + (d_3233_v102_)
                d_3246_v114_: _dafny.Array
                nw580_ = _dafny.Array(None, 25)
                d_3246_v114_ = nw580_
                d_3247_v115_: _dafny.Seq
                d_3247_v115_ = _dafny.SeqWithoutIsStrInference([d_3103_v0_, d_3103_v0_, d_3103_v0_, d_3103_v0_, d_3103_v0_])
                d_3248_v116_: T0
                nw581_ = C11()
                nw581_.ctor__(d_3247_v115_, d_3104_v1_, d_3104_v1_)
                d_3248_v116_ = nw581_
                index520_ = default__.safeIndex(968, (d_3246_v114_).length(0))
                (d_3246_v114_)[index520_] = d_3248_v116_
                index521_ = default__.safeIndex(968, (d_3246_v114_).length(0))
                (d_3246_v114_)[index521_] = d_3248_v116_
                d_3249_v117_: _dafny.Seq
                d_3249_v117_ = _dafny.SeqWithoutIsStrInference([d_3106_v3_])
                d_3250_v118_: _dafny.MultiSet
                d_3250_v118_ = _dafny.MultiSet([(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]])
                d_3251_v119_: _dafny.Seq
                d_3251_v119_ = _dafny.SeqWithoutIsStrInference([(d_3249_v117_)[default__.safeIndex((d_3250_v118_).cardinality, len(d_3249_v117_))], d_3106_v3_, d_3106_v3_])
                d_3252_v120_: _dafny.Set
                d_3252_v120_ = _dafny.Set({not((D22_DC59((d_3248_v116_).f29, d_3233_v102_, len((d_3251_v119_)[default__.safeIndex((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], len(d_3251_v119_))]))).cf83)})
                rhs542_ = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
                rhs543_ = (not((d_3252_v120_).ispropersubset(d_3252_v120_))) or ((d_3248_v116_).f29)
                lhs461_ = globalState
                lhs462_ = globalState
                lhs461_.f24 = rhs542_
                lhs462_.f27 = rhs543_
                (globalState).f22 = (d_3232_cf23_) and (d_3248_v116_.f28)
                d_3253_v121_: _dafny.Array
                nw582_ = _dafny.Array(False, 16)
                d_3253_v121_ = nw582_
                index522_ = default__.safeIndex(216, (d_3253_v121_).length(0))
                (d_3253_v121_)[index522_] = not (d_3248_v116_.f28) or (d_3232_cf23_)
                d_3254_v122_: C1
                nw583_ = C1()
                nw583_.ctor__((d_3248_v116_).f29, (d_3248_v116_).f29)
                d_3254_v122_ = nw583_
                d_3255_v123_: _dafny.Seq
                d_3255_v123_ = _dafny.SeqWithoutIsStrInference([d_3254_v122_])
                d_3256_v124_: _dafny.Map
                d_3256_v124_ = _dafny.Map({d_3255_v123_: (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]})
                d_3257_v125_: _dafny.Set
                d_3257_v125_ = _dafny.Set({default__.fm3(False, d_3248_v116_.f28, -445, globalState)})
                index523_ = default__.safeIndex(216, (d_3253_v121_).length(0))
                (d_3253_v121_)[index523_] = (len((d_3256_v124_) | (_dafny.Map({d_3255_v123_: (0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])})))) < (len(d_3257_v125_))
            d_3104_v1_ = d_3232_cf23_
            d_3258_v126_: _dafny.Array
            def lambda160_(d_3259_v1_, d_3260_v99_, d_3261_v3_, d_3262_p0_):
                def lambda161_(d_3263_i8_):
                    return D22_DC59(d_3259_v1_, (_dafny.SeqWithoutIsStrInference([d_3260_v99_ for d_3264_i9_ in range(default__.abs(809))])).set(default__.safeIndex(len(d_3261_v3_), len(_dafny.SeqWithoutIsStrInference([d_3260_v99_ for d_3265_i9_ in range(default__.abs(809))]))), d_3260_v99_), d_3262_p0_)

                return lambda161_

            init93_ = lambda160_(d_3104_v1_, d_3226_v99_, d_3106_v3_, p0)
            nw584_ = _dafny.Array(None, 25)
            for i0_93_ in range(nw584_.length(0)):
                nw584_[i0_93_] = init93_(i0_93_)
            d_3258_v126_ = nw584_
            d_3266_v127_: D22
            d_3266_v127_ = D22_DC59(True, d_3233_v102_, -203)
            index524_ = default__.safeIndex(862, (d_3258_v126_).length(0))
            (d_3258_v126_)[index524_] = d_3266_v127_
            d_3267_v128_: _dafny.Set
            d_3267_v128_ = _dafny.Set({d_3104_v1_})
            index525_ = default__.safeIndex(862, (d_3258_v126_).length(0))
            rhs544_ = _dafny.CodePoint('j')
            rhs545_ = default__.fm78(False, -955, (_dafny.Set({d_3232_cf23_})).intersection(d_3267_v128_), d_3232_cf23_, globalState)
            rhs546_ = d_3104_v1_
            lhs463_ = globalState
            lhs464_ = d_3258_v126_
            lhs465_ = default__.safeIndex(862, (d_3258_v126_).length(0))
            lhs466_ = globalState
            lhs463_.f23 = rhs544_
            lhs464_[lhs465_] = rhs545_
            lhs466_.f22 = rhs546_
        elif source49_.is_DC19:
            d_3268___mcc_h8_ = source49_.cf24
            d_3269___mcc_h9_ = source49_.cf25
            d_3270_cf25_ = d_3269___mcc_h9_
            d_3271_cf24_ = d_3268___mcc_h8_
            d_3272_v129_: C2
            nw585_ = C2()
            nw585_.ctor__()
            d_3272_v129_ = nw585_
            d_3273_v130_: _dafny.Map
            d_3273_v130_ = _dafny.Map({((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) >= (len(default__.fm51(len(_dafny.SeqWithoutIsStrInference([d_3226_v99_ for d_3274_i10_ in range(default__.abs(353))])), p0, globalState))): -929})
            d_3275_v131_: _dafny.Seq
            d_3275_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_3273_v130_ = (d_3273_v130_).set((d_3275_v131_) <= (d_3275_v131_), (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
            d_3276_v132_: _dafny.Seq
            d_3276_v132_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_, d_3104_v1_])
            d_3277_v133_: _dafny.Array
            nw586_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_3277_v133_ = nw586_
            d_3278_v134_: D18
            d_3278_v134_ = D18_DC48(False, (_dafny.MultiSet([d_3270_cf25_])).cardinality, d_3277_v133_, False, d_3226_v99_)
            if not((self).fm0((d_3276_v132_)[default__.safeIndex(p0, len(d_3276_v132_))], (d_3275_v131_) + ((d_3275_v131_).set(default__.safeIndex((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], len(d_3275_v131_)), (d_3278_v134_).cf72)), globalState)):
                d_3279_v135_: _dafny.Map
                d_3279_v135_ = _dafny.Map({p0: d_3104_v1_})
                d_3280_v136_: _dafny.MultiSet
                d_3280_v136_ = _dafny.MultiSet([d_3279_v135_])
                d_3104_v1_ = (((d_3280_v136_).set(d_3279_v135_, default__.abs((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]))) | (_dafny.MultiSet([d_3279_v135_]))).isdisjoint(d_3280_v136_)
                d_3281_v137_: _dafny.Array
                nw587_ = _dafny.Array(False, 4)
                d_3281_v137_ = nw587_
                index526_ = default__.safeIndex(781, (d_3281_v137_).length(0))
                (d_3281_v137_)[index526_] = d_3271_cf24_
                d_3282_v138_: D15
                d_3282_v138_ = D15_DC40(d_3271_cf24_, d_3271_cf24_)
                index527_ = default__.safeIndex(781, (d_3281_v137_).length(0))
                (d_3281_v137_)[index527_] = ((d_3282_v138_) in (_dafny.SeqWithoutIsStrInference([d_3282_v138_]))) == (d_3271_cf24_)
                (globalState).f1 = p0
                d_3283_v139_: _dafny.MultiSet
                d_3283_v139_ = _dafny.MultiSet([d_3105_v2_, d_3105_v2_, _dafny.SeqWithoutIsStrInference([(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], p0])])
                (globalState).f1 = default__.safeDivisionInt((d_3283_v139_).cardinality, p0)
                r2 = ((p0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]) if not(not(d_3104_v1_)) else default__.safeModuloInt((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], 456))
            elif True:
                (globalState).f13 = default__.fm26((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], globalState)
                d_3273_v130_ = d_3273_v130_
                d_3284_v140_: _dafny.MultiSet
                d_3284_v140_ = _dafny.MultiSet([len(d_3276_v132_)])
                d_3271_cf24_ = ((d_3284_v140_).set(len(d_3275_v131_), default__.abs((0) - (p0)))).ispropersubset(d_3284_v140_)
                r2 = (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]
                d_3285_v141_: _dafny.Map
                d_3285_v141_ = _dafny.Map({d_3226_v99_: d_3271_cf24_})
                d_3286_v142_: C3
                nw588_ = C3()
                nw588_.ctor__(((d_3285_v141_)[d_3226_v99_] if (d_3226_v99_) in (d_3285_v141_) else d_3104_v1_), d_3270_cf25_, not(False), p0)
                d_3286_v142_ = nw588_
                d_3287_v143_: _dafny.Set
                d_3287_v143_ = _dafny.Set({d_3104_v1_})
                d_3288_v144_: _dafny.Map
                d_3288_v144_ = _dafny.Map({d_3286_v142_: len(d_3287_v143_)})
                d_3289_v145_: _dafny.Map
                d_3289_v145_ = _dafny.Map({d_3288_v144_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ot"))})
                d_3289_v145_ = (d_3289_v145_).set(d_3288_v144_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilm")))
            d_3290_v146_: D2
            d_3290_v146_ = D2_DC4(_dafny.CodePoint('d'))
            d_3291_v147_: D24
            d_3291_v147_ = D24_DC67(d_3104_v1_, D5_DC13(d_3290_v146_), d_3271_cf24_, d_3270_cf25_, True)
            pat_let_tv75_ = d_3104_v1_
            def iife207_(_pat_let53_0):
                def iife208_(d_3292_dt__update__tmp_h2_):
                    def iife209_(_pat_let54_0):
                        def iife210_(d_3293_dt__update_hcf103_h0_):
                            def iife211_(_pat_let55_0):
                                def iife212_(d_3294_dt__update_hcf104_h0_):
                                    return D24_DC67((d_3292_dt__update__tmp_h2_).cf100, (d_3292_dt__update__tmp_h2_).cf101, (d_3292_dt__update__tmp_h2_).cf102, d_3293_dt__update_hcf103_h0_, d_3294_dt__update_hcf104_h0_)
                                return iife212_(_pat_let55_0)
                            return iife211_(False)
                        return iife210_(_pat_let54_0)
                    return iife209_(pat_let_tv75_)
                return iife208_(_pat_let53_0)
            (globalState).f24 = default__.safeDivisionInt((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], (len(_dafny.Set({iife207_(d_3291_v147_), d_3291_v147_}))) * (p0))
        elif True:
            d_3295___mcc_h10_ = source49_.cf21
            d_3296_cf21_ = d_3295___mcc_h10_
            if d_3104_v1_:
                d_3297_v148_: _dafny.Map
                d_3297_v148_ = _dafny.Map({(d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]: d_3104_v1_})
                d_3298_v149_: D10
                d_3298_v149_ = D10_DC26(d_3297_v148_)
                d_3299_v150_: D10
                d_3299_v150_ = D10_DC29(d_3298_v149_)
                d_3300_v151_: _dafny.Seq
                d_3300_v151_ = _dafny.SeqWithoutIsStrInference([d_3299_v150_, d_3299_v150_, d_3299_v150_, d_3299_v150_, d_3299_v150_])
                d_3301_v152_: D20
                d_3301_v152_ = D20_DC53(D20_DC50(d_3300_v151_))
                d_3302_v153_: D20
                d_3302_v153_ = D20_DC53(d_3301_v152_)
                pat_let_tv76_ = d_3301_v152_
                d_3303_v154_: _dafny.Array
                nw589_ = _dafny.Array(None, 15)
                nw589_[int(0)] = d_3302_v153_
                nw589_[int(1)] = d_3302_v153_
                nw589_[int(2)] = d_3302_v153_
                nw589_[int(3)] = d_3302_v153_
                nw589_[int(4)] = d_3302_v153_
                nw589_[int(5)] = default__.fm79((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], d_3104_v1_, default__.fm3(d_3104_v1_, d_3104_v1_, p0, globalState), p0, globalState)
                nw589_[int(6)] = d_3302_v153_
                def iife213_(_pat_let56_0):
                    def iife214_(d_3304_dt__update__tmp_h3_):
                        def iife215_(_pat_let57_0):
                            def iife216_(d_3305_dt__update_hcf78_h0_):
                                return D20_DC53(d_3305_dt__update_hcf78_h0_)
                            return iife216_(_pat_let57_0)
                        return iife215_(pat_let_tv76_)
                    return iife214_(_pat_let56_0)
                nw589_[int(7)] = iife213_(D20_DC53(d_3301_v152_))
                nw589_[int(8)] = d_3302_v153_
                nw589_[int(9)] = d_3302_v153_
                nw589_[int(10)] = D20_DC53(d_3301_v152_)
                nw589_[int(11)] = d_3302_v153_
                nw589_[int(12)] = d_3302_v153_
                nw589_[int(13)] = d_3302_v153_
                nw589_[int(14)] = d_3302_v153_
                d_3303_v154_ = nw589_
                index528_ = default__.safeIndex(790, (d_3303_v154_).length(0))
                (d_3303_v154_)[index528_] = d_3302_v153_
                index529_ = default__.safeIndex(790, (d_3303_v154_).length(0))
                (d_3303_v154_)[index529_] = d_3302_v153_
                d_3306_v155_: _dafny.Seq
                d_3306_v155_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_, d_3104_v1_, d_3104_v1_])
                d_3307_v156_: _dafny.Map
                d_3307_v156_ = _dafny.Map({(0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]): (d_3306_v155_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neaexvxqb"))), len(d_3306_v155_)), d_3104_v1_)})
                d_3307_v156_ = (d_3307_v156_).set(default__.safeDivisionInt(p0, (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]), (d_3306_v155_) + (default__.fm62(globalState)))
                (globalState).f24 = (p0) + ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
                d_3308_v157_: _dafny.Seq
                d_3308_v157_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hidqd"))
                d_3308_v157_ = d_3308_v157_
                d_3309_v158_: _dafny.Array
                def lambda162_(d_3310_v1_):
                    def lambda163_(d_3311_i11_):
                        return d_3310_v1_

                    return lambda163_

                init94_ = lambda162_(d_3104_v1_)
                nw590_ = _dafny.Array(None, 16)
                for i0_94_ in range(nw590_.length(0)):
                    nw590_[i0_94_] = init94_(i0_94_)
                d_3309_v158_ = nw590_
                d_3312_v159_: _dafny.Set
                d_3312_v159_ = _dafny.Set({True, d_3104_v1_})
                index530_ = default__.safeIndex(924, (d_3309_v158_).length(0))
                (d_3309_v158_)[index530_] = (d_3312_v159_).isdisjoint(d_3312_v159_)
                index531_ = default__.safeIndex(924, (d_3309_v158_).length(0))
                (d_3309_v158_)[index531_] = d_3104_v1_
            elif True:
                d_3313_v160_: _dafny.Seq
                d_3313_v160_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fx"))
                d_3314_v161_: T1
                nw591_ = C3()
                nw591_.ctor__(d_3104_v1_, d_3104_v1_, d_3104_v1_, p0)
                d_3314_v161_ = nw591_
                d_3315_v162_: D25
                d_3315_v162_ = D25_DC68(d_3314_v161_)
                d_3316_v163_: _dafny.Array
                nw592_ = _dafny.Array(None, 27)
                nw592_[int(0)] = d_3314_v161_
                nw592_[int(1)] = d_3314_v161_
                nw592_[int(2)] = d_3314_v161_
                nw592_[int(3)] = d_3314_v161_
                nw592_[int(4)] = d_3314_v161_
                nw592_[int(5)] = d_3314_v161_
                nw592_[int(6)] = d_3314_v161_
                nw592_[int(7)] = d_3314_v161_
                nw592_[int(8)] = d_3314_v161_
                nw592_[int(9)] = d_3314_v161_
                nw592_[int(10)] = d_3314_v161_
                nw592_[int(11)] = d_3314_v161_
                nw592_[int(12)] = d_3314_v161_
                nw592_[int(13)] = d_3314_v161_
                nw592_[int(14)] = d_3314_v161_
                nw592_[int(15)] = d_3314_v161_
                nw592_[int(16)] = d_3314_v161_
                nw592_[int(17)] = d_3314_v161_
                nw592_[int(18)] = d_3314_v161_
                nw592_[int(19)] = (d_3315_v162_).cf105
                nw592_[int(20)] = d_3314_v161_
                nw592_[int(21)] = (d_3315_v162_).cf105
                nw592_[int(22)] = d_3314_v161_
                nw592_[int(23)] = d_3314_v161_
                nw592_[int(24)] = d_3314_v161_
                nw592_[int(25)] = d_3314_v161_
                nw592_[int(26)] = d_3314_v161_
                d_3316_v163_ = nw592_
                d_3317_v164_: _dafny.Map
                d_3317_v164_ = _dafny.Map({(d_3313_v160_) + ((d_3313_v160_).set(default__.safeIndex((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], len(d_3313_v160_)), _dafny.CodePoint('o'))): d_3316_v163_})
                d_3317_v164_ = d_3317_v164_
                (globalState).f13 = default__.fm13(globalState)
                d_3318_v165_: _dafny.Array
                nw593_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_3318_v165_ = nw593_
                index532_ = default__.safeIndex(53, (d_3318_v165_).length(0))
                (d_3318_v165_)[index532_] = d_3226_v99_
                index533_ = default__.safeIndex(53, (d_3318_v165_).length(0))
                (d_3318_v165_)[index533_] = d_3226_v99_
                d_3319_v166_: _dafny.Array
                nw594_ = _dafny.Array(None, 4)
                nw594_[int(0)] = d_3104_v1_
                nw594_[int(1)] = d_3104_v1_
                nw594_[int(2)] = d_3104_v1_
                nw594_[int(3)] = d_3104_v1_
                d_3319_v166_ = nw594_
                index534_ = default__.safeIndex(282, (d_3319_v166_).length(0))
                (d_3319_v166_)[index534_] = d_3104_v1_
                index535_ = default__.safeIndex(282, (d_3319_v166_).length(0))
                (d_3319_v166_)[index535_] = d_3104_v1_
                (globalState).f21 = d_3104_v1_
            (globalState).f27 = d_3104_v1_
            d_3320_v167_: _dafny.Map
            d_3320_v167_ = _dafny.Map({d_3103_v0_: 691})
            (globalState).f27 = default__.fm26(len((d_3320_v167_) | (d_3320_v167_)), globalState)
            source50_ = d_3227_v100_
            if source50_.is_DC17:
                d_3321___mcc_h11_ = source50_.cf22
                d_3322_cf22_ = d_3321___mcc_h11_
                d_3323_v168_: int
                d_3324_v169_: int
                d_3325_v170_: bool
                d_3326_v171_: int
                out45_: int
                out46_: int
                out47_: bool
                out48_: int
                out45_, out46_, out47_, out48_ = (self).m2(globalState)
                d_3323_v168_ = out45_
                d_3324_v169_ = out46_
                d_3325_v170_ = out47_
                d_3326_v171_ = out48_
                (globalState).f27 = d_3104_v1_
                r2 = d_3326_v171_
                d_3327_v172_: C7
                nw595_ = C7()
                nw595_.ctor__(not(d_3322_cf22_), d_3322_cf22_)
                d_3327_v172_ = nw595_
            elif source50_.is_DC18:
                d_3328___mcc_h12_ = source50_.cf23
                d_3329_cf23_ = d_3328___mcc_h12_
                d_3330_v173_: C9
                nw596_ = C9()
                nw596_.ctor__(d_3329_cf23_, d_3104_v1_)
                d_3330_v173_ = nw596_
                d_3330_v173_ = d_3330_v173_
                d_3331_v174_: bool
                out49_: bool
                out49_ = (d_3330_v173_).m5((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], globalState)
                d_3331_v174_ = out49_
                index536_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                (d_3103_v0_)[index536_] = p0
                d_3332_v175_: _dafny.Seq
                d_3332_v175_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mloh"))
                d_3333_v176_: _dafny.MultiSet
                d_3333_v176_ = _dafny.MultiSet([d_3329_cf23_])
                d_3334_v177_: _dafny.Map
                d_3334_v177_ = _dafny.Map({((d_3333_v176_)[d_3329_cf23_] if (d_3329_cf23_) in (d_3333_v176_) else (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]): d_3331_v174_})
                d_3335_v178_: _dafny.Map
                d_3335_v178_ = _dafny.Map({((d_3334_v177_)[p0] if (p0) in (d_3334_v177_) else False): (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]})
                (globalState).f24 = len((d_3332_v175_) + (default__.fm51(((d_3335_v178_)[d_3104_v1_] if (d_3104_v1_) in (d_3335_v178_) else (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]), p0, globalState)))
            elif source50_.is_DC19:
                d_3336___mcc_h13_ = source50_.cf24
                d_3337___mcc_h14_ = source50_.cf25
                d_3338_cf25_ = d_3337___mcc_h14_
                d_3339_cf24_ = d_3336___mcc_h13_
                (globalState).f13 = False
                index537_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                (d_3103_v0_)[index537_] = 604
                (globalState).f27 = d_3339_cf24_
                d_3340_v179_: D23
                d_3340_v179_ = D23_DC64(-654, p0, d_3339_cf24_, (0) - (p0), (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
                pat_let_tv77_ = p0
                pat_let_tv78_ = p0
                pat_let_tv79_ = d_3103_v0_
                pat_let_tv80_ = d_3103_v0_
                def iife217_(_pat_let58_0):
                    def iife218_(d_3341_dt__update__tmp_h4_):
                        def iife219_(_pat_let59_0):
                            def iife220_(d_3342_dt__update_hcf97_h0_):
                                def iife221_(_pat_let60_0):
                                    def iife222_(d_3343_dt__update_hcf93_h0_):
                                        def iife223_(_pat_let61_0):
                                            def iife224_(d_3344_dt__update_hcf96_h0_):
                                                return D23_DC64(d_3343_dt__update_hcf93_h0_, (d_3341_dt__update__tmp_h4_).cf94, (d_3341_dt__update__tmp_h4_).cf95, d_3344_dt__update_hcf96_h0_, d_3342_dt__update_hcf97_h0_)
                                            return iife224_(_pat_let61_0)
                                        return iife223_((pat_let_tv80_)[default__.safeIndex(696, (pat_let_tv79_).length(0))])
                                    return iife222_(_pat_let60_0)
                                return iife221_(pat_let_tv78_)
                            return iife220_(_pat_let59_0)
                        return iife219_(pat_let_tv77_)
                    return iife218_(_pat_let58_0)
                (globalState).f22 = ((len(_dafny.Map({d_3339_cf24_: not(d_3339_cf24_)}))) == (p0) if ((0) - (len(_dafny.Map({d_3226_v99_: (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]})))) <= (p0) else ((iife217_(d_3340_v179_)).cf97) >= (default__.fm27((0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))]), d_3226_v99_, (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], globalState)))
            elif True:
                d_3345___mcc_h15_ = source50_.cf21
                d_3346_cf21_ = d_3345___mcc_h15_
                d_3347_v180_: _dafny.Array
                nw597_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_3347_v180_ = nw597_
                d_3347_v180_ = d_3347_v180_
                d_3348_v181_: _dafny.Seq
                d_3348_v181_ = _dafny.SeqWithoutIsStrInference([d_3104_v1_])
                index538_ = default__.safeIndex(696, (d_3103_v0_).length(0))
                (d_3103_v0_)[index538_] = (p0) - (len((d_3348_v181_) + (d_3348_v181_)))
                rhs547_ = 846
                rhs548_ = (0) - (-122)
                rhs549_ = d_3104_v1_
                rhs550_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhtsfluwm")))) for d_3349_i12_ in range(default__.abs(382))])
                lhs467_ = globalState
                lhs468_ = globalState
                lhs469_ = globalState
                lhs470_ = globalState
                lhs467_.f24 = rhs547_
                lhs468_.f24 = rhs548_
                lhs469_.f22 = rhs549_
                lhs470_.f17 = rhs550_
                r2 = (0) - ((d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))])
        d_3350_v182_: _dafny.Array
        nw598_ = _dafny.Array(False, 20)
        d_3350_v182_ = nw598_
        d_3351_v183_: D5
        d_3351_v183_ = D5_DC12(d_3350_v182_)
        r0 = (d_3351_v183_).cf17
        d_3352_v184_: C5
        nw599_ = C5()
        nw599_.ctor__(False, default__.fm42(p0, (d_3103_v0_)[default__.safeIndex(696, (d_3103_v0_).length(0))], globalState), d_3104_v1_, d_3104_v1_)
        d_3352_v184_ = nw599_
        d_3353_v185_: _dafny.MultiSet
        d_3353_v185_ = _dafny.MultiSet([d_3352_v184_])
        r1 = (d_3353_v185_).cardinality
        d_3354_v186_: _dafny.Seq
        d_3354_v186_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjf"))
        r2 = len(d_3354_v186_)
        r3 = d_3104_v1_
        return r0, r1, r2, r3

    def m2(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_3355_v0_: _dafny.Array
        nw600_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_3355_v0_ = nw600_
        d_3356_v1_: int
        d_3356_v1_ = 156
        d_3357_v2_: _dafny.Seq
        d_3357_v2_ = _dafny.SeqWithoutIsStrInference([True])
        d_3358_v3_: _dafny.Map
        d_3358_v3_ = _dafny.Map({len(d_3357_v2_): d_3356_v1_})
        d_3359_v4_: _dafny.Seq
        d_3359_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvjaxm"))
        d_3360_v5_: _dafny.Array
        nw601_ = _dafny.Array(None, 21)
        nw601_[int(0)] = d_3356_v1_
        nw601_[int(1)] = d_3356_v1_
        nw601_[int(2)] = d_3356_v1_
        nw601_[int(3)] = -407
        nw601_[int(4)] = len(d_3357_v2_)
        nw601_[int(5)] = len((d_3358_v3_).set(d_3356_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iaa")))))
        nw601_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjvp")))
        nw601_[int(7)] = d_3356_v1_
        nw601_[int(8)] = d_3356_v1_
        nw601_[int(9)] = d_3356_v1_
        nw601_[int(10)] = (0) - (d_3356_v1_)
        nw601_[int(11)] = d_3356_v1_
        nw601_[int(12)] = d_3356_v1_
        nw601_[int(13)] = d_3356_v1_
        nw601_[int(14)] = 375
        nw601_[int(15)] = len(d_3359_v4_)
        nw601_[int(16)] = -263
        nw601_[int(17)] = 754
        nw601_[int(18)] = -473
        nw601_[int(19)] = d_3356_v1_
        nw601_[int(20)] = -912
        d_3360_v5_ = nw601_
        index539_ = default__.safeIndex(7, (d_3355_v0_).length(0))
        (d_3355_v0_)[index539_] = d_3360_v5_
        index540_ = default__.safeIndex(7, (d_3355_v0_).length(0))
        (d_3355_v0_)[index540_] = d_3360_v5_
        d_3361_v6_: _dafny.Array
        nw602_ = _dafny.Array(_dafny.CodePoint('D'), 18)
        d_3361_v6_ = nw602_
        d_3362_v7_: bool
        d_3362_v7_ = True
        d_3363_v8_: _dafny.Set
        d_3363_v8_ = _dafny.Set({d_3356_v1_, d_3356_v1_})
        d_3364_v9_: _dafny.Seq
        d_3364_v9_ = _dafny.SeqWithoutIsStrInference([d_3356_v1_, len(d_3363_v8_), len(d_3359_v4_), d_3356_v1_, d_3356_v1_])
        d_3365_v10_: _dafny.Map
        d_3365_v10_ = _dafny.Map({d_3362_v7_: d_3364_v9_})
        index541_ = default__.safeIndex(508, (d_3361_v6_).length(0))
        (d_3361_v6_)[index541_] = (d_3359_v4_)[default__.safeIndex(len(((d_3365_v10_)[True] if (True) in (d_3365_v10_) else d_3364_v9_)), len(d_3359_v4_))]
        d_3366_v11_: str
        d_3366_v11_ = _dafny.CodePoint('d')
        index542_ = default__.safeIndex(508, (d_3361_v6_).length(0))
        (d_3361_v6_)[index542_] = (d_3366_v11_ if not(d_3362_v7_) else d_3366_v11_)
        d_3367_v12_: _dafny.Set
        d_3367_v12_ = _dafny.Set({not(d_3362_v7_)})
        if (d_3367_v12_) == (_dafny.Set({d_3362_v7_, default__.fm26(d_3356_v1_, globalState), False})):
            (globalState).f24 = (0) - (d_3356_v1_)
            d_3368_v13_: _dafny.MultiSet
            d_3368_v13_ = _dafny.MultiSet([d_3357_v2_, d_3357_v2_, d_3357_v2_])
            d_3369_v14_: _dafny.MultiSet
            d_3369_v14_ = _dafny.MultiSet([d_3362_v7_, d_3362_v7_, d_3362_v7_, default__.fm29(d_3368_v13_, len(d_3364_v9_), globalState), d_3362_v7_])
            if ((self).fm0(False, default__.fm51((d_3369_v14_).cardinality, d_3356_v1_, globalState), globalState) if d_3362_v7_ else d_3362_v7_):
                d_3370_v15_: T1
                nw603_ = C3()
                nw603_.ctor__(d_3362_v7_, not(not((d_3357_v2_)[default__.safeIndex(43, len(d_3357_v2_))])), not (d_3362_v7_) or (False), default__.fm27(len(d_3359_v4_), d_3366_v11_, d_3356_v1_, globalState))
                d_3370_v15_ = nw603_
                d_3370_v15_ = d_3370_v15_
                (globalState).f27 = d_3362_v7_
                d_3371_v16_: C9
                nw604_ = C9()
                nw604_.ctor__(d_3362_v7_, d_3362_v7_)
                d_3371_v16_ = nw604_
                d_3372_v17_: _dafny.Seq
                d_3372_v17_ = _dafny.SeqWithoutIsStrInference([d_3371_v16_, d_3371_v16_])
                d_3373_v18_: _dafny.Map
                d_3373_v18_ = _dafny.Map({d_3356_v1_: d_3362_v7_})
                d_3374_v19_: _dafny.Array
                nw605_ = _dafny.Array(None, 18)
                nw605_[int(0)] = d_3371_v16_
                nw605_[int(1)] = d_3371_v16_
                nw605_[int(2)] = d_3371_v16_
                nw605_[int(3)] = d_3371_v16_
                nw605_[int(4)] = (d_3372_v17_)[default__.safeIndex(d_3356_v1_, len(d_3372_v17_))]
                nw605_[int(5)] = d_3371_v16_
                nw605_[int(6)] = d_3371_v16_
                nw605_[int(7)] = d_3371_v16_
                nw605_[int(8)] = d_3371_v16_
                nw605_[int(9)] = d_3371_v16_
                nw605_[int(10)] = d_3371_v16_
                nw605_[int(11)] = d_3371_v16_
                nw605_[int(12)] = d_3371_v16_
                nw605_[int(13)] = d_3371_v16_
                nw605_[int(14)] = d_3371_v16_
                nw605_[int(15)] = d_3371_v16_
                nw605_[int(16)] = (d_3371_v16_ if ((d_3373_v18_)[d_3356_v1_] if (d_3356_v1_) in (d_3373_v18_) else default__.fm29(d_3368_v13_, default__.fm27((0) - (d_3356_v1_), d_3366_v11_, d_3356_v1_, globalState), globalState)) else d_3371_v16_)
                nw605_[int(17)] = (d_3371_v16_ if d_3362_v7_ else (d_3372_v17_)[default__.safeIndex(d_3356_v1_, len(d_3372_v17_))])
                d_3374_v19_ = nw605_
                d_3374_v19_ = d_3374_v19_
                d_3375_v20_: C5
                nw606_ = C5()
                nw606_.ctor__((d_3362_v7_) not in (default__.fm59(d_3362_v7_, d_3357_v2_, globalState)), _dafny.CodePoint('u'), d_3362_v7_, d_3362_v7_)
                d_3375_v20_ = nw606_
                d_3371_v16_ = d_3371_v16_
            elif True:
                d_3376_v21_: _dafny.Array
                nw607_ = _dafny.Array(_dafny.Map({}), 14)
                d_3376_v21_ = nw607_
                d_3377_v22_: _dafny.Map
                d_3377_v22_ = _dafny.Map({d_3362_v7_: d_3356_v1_})
                index543_ = default__.safeIndex(700, (d_3376_v21_).length(0))
                (d_3376_v21_)[index543_] = d_3377_v22_
                index544_ = default__.safeIndex(700, (d_3376_v21_).length(0))
                (d_3376_v21_)[index544_] = ((d_3377_v22_) | (d_3377_v22_) if d_3362_v7_ else _dafny.Map({d_3362_v7_: d_3356_v1_}))
                d_3378_v23_: _dafny.Map
                d_3378_v23_ = _dafny.Map({_dafny.Map({d_3362_v7_: d_3362_v7_}): d_3369_v14_})
                d_3379_v24_: _dafny.Map
                d_3379_v24_ = _dafny.Map({d_3362_v7_: True})
                d_3380_v25_: _dafny.Seq
                d_3380_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_3362_v7_, d_3362_v7_])])
                d_3378_v23_ = (d_3378_v23_).set((d_3379_v24_) | (_dafny.Map({d_3362_v7_: d_3362_v7_})), ((d_3380_v25_)[default__.safeIndex(d_3356_v1_, len(d_3380_v25_))]) | (d_3369_v14_))
                d_3381_v26_: _dafny.Array
                nw608_ = _dafny.Array(False, 16)
                d_3381_v26_ = nw608_
                index545_ = default__.safeIndex(387, (d_3381_v26_).length(0))
                (d_3381_v26_)[index545_] = (d_3377_v22_) == (((d_3376_v21_)[default__.safeIndex(700, (d_3376_v21_).length(0))]).set(d_3362_v7_, -153))
                index546_ = default__.safeIndex(958, (d_3381_v26_).length(0))
                (d_3381_v26_)[index546_] = d_3362_v7_
                d_3382_v27_: C2
                nw609_ = C2()
                nw609_.ctor__()
                d_3382_v27_ = nw609_
                index547_ = default__.safeIndex(387, (d_3381_v26_).length(0))
                index548_ = default__.safeIndex(958, (d_3381_v26_).length(0))
                rhs551_ = not((-485) == (default__.fm12(d_3356_v1_, d_3362_v7_, globalState)))
                rhs552_ = d_3362_v7_
                rhs553_ = d_3359_v4_
                rhs554_ = d_3382_v27_
                lhs471_ = d_3381_v26_
                lhs472_ = default__.safeIndex(387, (d_3381_v26_).length(0))
                lhs473_ = d_3381_v26_
                lhs474_ = default__.safeIndex(958, (d_3381_v26_).length(0))
                lhs471_[lhs472_] = rhs551_
                lhs473_[lhs474_] = rhs552_
                d_3359_v4_ = rhs553_
                d_3382_v27_ = rhs554_
                (globalState).f24 = d_3356_v1_
                index549_ = default__.safeIndex(198, (d_3360_v5_).length(0))
                (d_3360_v5_)[index549_] = d_3356_v1_
                index550_ = default__.safeIndex(198, (d_3360_v5_).length(0))
                (d_3360_v5_)[index550_] = d_3356_v1_
            d_3383_v28_: D12
            d_3383_v28_ = D12_DC31(_dafny.Set({d_3356_v1_}))
            d_3384_v29_: D12
            d_3384_v29_ = D12_DC33(d_3383_v28_)
            source51_ = d_3384_v29_
            if source51_.is_DC32:
                d_3357_v2_ = _dafny.SeqWithoutIsStrInference([d_3362_v7_])
                d_3385_v30_: _dafny.Map
                d_3385_v30_ = _dafny.Map({(d_3361_v6_)[default__.safeIndex(508, (d_3361_v6_).length(0))]: d_3356_v1_})
                d_3385_v30_ = (d_3385_v30_).set((d_3361_v6_)[default__.safeIndex(508, (d_3361_v6_).length(0))], d_3356_v1_)
                r2 = (d_3356_v1_) < (d_3356_v1_)
                d_3386_v31_: _dafny.Seq
                d_3386_v31_ = _dafny.SeqWithoutIsStrInference([(d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]])
                d_3387_v32_: C11
                nw610_ = C11()
                nw610_.ctor__((d_3386_v31_).set(default__.safeIndex(d_3356_v1_, len(d_3386_v31_)), d_3360_v5_), (not(d_3362_v7_)) == (d_3362_v7_), (d_3362_v7_ if d_3362_v7_ else d_3362_v7_))
                d_3387_v32_ = nw610_
            elif source51_.is_DC31:
                d_3388___mcc_h0_ = source51_.cf48
                d_3389_cf48_ = d_3388___mcc_h0_
                (globalState).f24 = (0) - (d_3356_v1_)
                d_3390_v33_: _dafny.Array
                nw611_ = _dafny.Array(_dafny.Map({}), 8)
                d_3390_v33_ = nw611_
                d_3391_v34_: _dafny.Map
                d_3391_v34_ = _dafny.Map({(d_3390_v33_ if default__.fm10(d_3362_v7_, d_3356_v1_, d_3362_v7_, globalState) else d_3390_v33_): d_3356_v1_})
                d_3391_v34_ = (d_3391_v34_).set(d_3390_v33_, 93)
                d_3366_v11_ = d_3366_v11_
                (globalState).f13 = d_3362_v7_
            elif True:
                d_3392___mcc_h1_ = source51_.cf49
                d_3393_cf49_ = d_3392___mcc_h1_
                (globalState).f22 = d_3362_v7_
                d_3394_v35_: _dafny.Map
                d_3394_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_3366_v11_ for d_3395_i0_ in range(default__.abs(101))]): d_3362_v7_})
                d_3396_v36_: D17
                d_3396_v36_ = D17_DC43(d_3394_v35_)
                d_3396_v36_ = d_3396_v36_
                r2 = d_3362_v7_
                d_3397_v37_: C0
                nw612_ = C0()
                nw612_.ctor__()
                d_3397_v37_ = nw612_
            d_3398_v38_: _dafny.Array
            nw613_ = _dafny.Array(False, 26)
            d_3398_v38_ = nw613_
            index551_ = default__.safeIndex(787, (d_3398_v38_).length(0))
            (d_3398_v38_)[index551_] = (default__.fm18(default__.fm8(globalState), d_3356_v1_, d_3356_v1_, globalState)) == (d_3362_v7_)
            index552_ = default__.safeIndex(787, (d_3398_v38_).length(0))
            (d_3398_v38_)[index552_] = d_3362_v7_
            (globalState).f1 = ((0) - (d_3356_v1_)) - (d_3356_v1_)
        elif True:
            if (d_3362_v7_ if d_3362_v7_ else (d_3362_v7_) and (d_3362_v7_)):
                (globalState).f22 = not(((default__.fm25(d_3356_v1_, (0) - (d_3356_v1_), globalState)) - (d_3356_v1_)) > (d_3356_v1_))
                arr0_ = (d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]
                index553_ = default__.safeIndex(647, ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0))
                arr0_[index553_] = 350
                arr1_ = (d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]
                index554_ = default__.safeIndex(647, ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0))
                arr1_[index554_] = len((d_3357_v2_) + (d_3357_v2_))
                d_3399_v39_: _dafny.Array
                nw614_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_3399_v39_ = nw614_
                index555_ = default__.safeIndex(395, (d_3399_v39_).length(0))
                (d_3399_v39_)[index555_] = d_3360_v5_
                index556_ = default__.safeIndex(395, (d_3399_v39_).length(0))
                (d_3399_v39_)[index556_] = d_3360_v5_
                d_3400_v40_: _dafny.Array
                def lambda164_(d_3401_v7_):
                    def lambda165_(d_3402_i1_):
                        return d_3401_v7_

                    return lambda165_

                init95_ = lambda164_(d_3362_v7_)
                nw615_ = _dafny.Array(None, 22)
                for i0_95_ in range(nw615_.length(0)):
                    nw615_[i0_95_] = init95_(i0_95_)
                d_3400_v40_ = nw615_
                d_3403_v41_: _dafny.Map
                d_3403_v41_ = _dafny.Map({d_3400_v40_: (0) - ((d_3356_v1_ if True else d_3356_v1_))})
                d_3403_v41_ = (d_3403_v41_).set(d_3400_v40_, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqkpb"))), ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))])[default__.safeIndex(647, ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0))]))
                d_3404_v42_: _dafny.Map
                d_3404_v42_ = _dafny.Map({d_3362_v7_: d_3356_v1_})
                r2 = ((0) - ((((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))])[default__.safeIndex(647, ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0))]) - (len(d_3404_v42_)))) <= ((0) - (len(_dafny.Map({(0) - (((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))])[default__.safeIndex(647, ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0))]): (0) - ((0) - (d_3356_v1_))}))))
            elif True:
                (globalState).f24 = d_3356_v1_
                d_3405_v43_: _dafny.MultiSet
                d_3405_v43_ = _dafny.MultiSet([d_3361_v6_, d_3361_v6_])
                d_3406_v44_: _dafny.MultiSet
                d_3406_v44_ = _dafny.MultiSet([d_3359_v4_])
                r2 = ((d_3405_v43_).set(d_3361_v6_, default__.abs((d_3406_v44_).cardinality))).issubset((d_3405_v43_).set(d_3361_v6_, default__.abs(d_3356_v1_)))
                d_3407_v45_: _dafny.Map
                d_3407_v45_ = _dafny.Map({d_3356_v1_: d_3355_v0_})
                d_3407_v45_ = (d_3407_v45_).set((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))) * (d_3356_v1_), d_3355_v0_)
                r0 = d_3356_v1_
                d_3408_v46_: C2
                nw616_ = C2()
                nw616_.ctor__()
                d_3408_v46_ = nw616_
                d_3408_v46_ = d_3408_v46_
            d_3409_v47_: C9
            nw617_ = C9()
            nw617_.ctor__(d_3362_v7_, d_3362_v7_)
            d_3409_v47_ = nw617_
            index557_ = default__.safeIndex(461, (d_3360_v5_).length(0))
            (d_3360_v5_)[index557_] = d_3356_v1_
            index558_ = default__.safeIndex(461, (d_3360_v5_).length(0))
            rhs555_ = (0) - (-479)
            rhs556_ = default__.fm62(globalState)
            lhs475_ = d_3360_v5_
            lhs476_ = default__.safeIndex(461, (d_3360_v5_).length(0))
            lhs475_[lhs476_] = rhs555_
            d_3357_v2_ = rhs556_
            (globalState).f23 = (d_3366_v11_ if d_3362_v7_ else (d_3361_v6_)[default__.safeIndex(508, (d_3361_v6_).length(0))])
            if default__.fm26(-668, globalState):
                (globalState).f1 = 836
                d_3410_v48_: _dafny.Map
                d_3410_v48_ = _dafny.Map({d_3362_v7_: len((d_3357_v2_).set(default__.safeIndex(d_3356_v1_, len(d_3357_v2_)), d_3362_v7_))})
                d_3411_v49_: _dafny.Map
                d_3411_v49_ = _dafny.Map({(d_3364_v9_) <= (d_3364_v9_): ((d_3410_v48_)[d_3362_v7_] if (d_3362_v7_) in (d_3410_v48_) else (d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))])})
                d_3412_v50_: _dafny.Array
                def lambda166_(d_3413_v7_):
                    def lambda167_(d_3414_i2_):
                        return d_3413_v7_

                    return lambda167_

                init96_ = lambda166_(d_3362_v7_)
                nw618_ = _dafny.Array(None, 25)
                for i0_96_ in range(nw618_.length(0)):
                    nw618_[i0_96_] = init96_(i0_96_)
                d_3412_v50_ = nw618_
                index559_ = default__.safeIndex(155, (d_3412_v50_).length(0))
                (d_3412_v50_)[index559_] = d_3362_v7_
                d_3415_v51_: _dafny.Map
                d_3415_v51_ = _dafny.Map({d_3409_v47_: d_3362_v7_})
                index560_ = default__.safeIndex(643, (d_3412_v50_).length(0))
                (d_3412_v50_)[index560_] = ((d_3415_v51_)[d_3409_v47_] if (d_3409_v47_) in (d_3415_v51_) else d_3362_v7_)
                d_3416_v52_: _dafny.Seq
                d_3416_v52_ = _dafny.SeqWithoutIsStrInference([d_3411_v49_, d_3410_v48_])
                d_3417_v53_: D9
                d_3417_v53_ = D9_DC25(d_3362_v7_, d_3356_v1_, d_3356_v1_, d_3362_v7_, -954)
                d_3418_v54_: _dafny.Map
                d_3418_v54_ = _dafny.Map({(d_3417_v53_).cf37: d_3359_v4_})
                d_3419_v55_: _dafny.Map
                d_3419_v55_ = _dafny.Map({d_3362_v7_: d_3362_v7_})
                index561_ = default__.safeIndex(155, (d_3412_v50_).length(0))
                index562_ = default__.safeIndex(643, (d_3412_v50_).length(0))
                rhs557_ = ((d_3411_v49_ if d_3362_v7_ else _dafny.Map({d_3362_v7_: (d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))]}))) | ((d_3416_v52_)[default__.safeIndex((d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))], len(d_3416_v52_))])
                rhs558_ = (d_3366_v11_) in ((d_3359_v4_) + (d_3359_v4_))
                rhs559_ = ((d_3361_v6_)[default__.safeIndex(508, (d_3361_v6_).length(0))]) in (((d_3418_v54_)[(d_3417_v53_).cf37] if ((d_3417_v53_).cf37) in (d_3418_v54_) else d_3359_v4_))
                rhs560_ = ((d_3419_v55_)[False] if (False) in (d_3419_v55_) else d_3362_v7_)
                lhs477_ = d_3412_v50_
                lhs478_ = default__.safeIndex(155, (d_3412_v50_).length(0))
                lhs479_ = d_3412_v50_
                lhs480_ = default__.safeIndex(643, (d_3412_v50_).length(0))
                d_3410_v48_ = rhs557_
                lhs477_[lhs478_] = rhs558_
                lhs479_[lhs480_] = rhs559_
                r2 = rhs560_
                d_3420_v56_: _dafny.Map
                d_3420_v56_ = _dafny.Map({default__.safeDivisionInt((d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))], (0) - ((0) - (d_3356_v1_))): not((d_3412_v50_)[default__.safeIndex(155, (d_3412_v50_).length(0))])})
                d_3420_v56_ = (d_3420_v56_).set((len(d_3367_v12_)) + ((d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))]), d_3362_v7_)
                (globalState).f21 = d_3362_v7_
                d_3421_v57_: D17
                d_3421_v57_ = D17_DC45(d_3419_v55_, d_3362_v7_)
                d_3422_v58_: _dafny.Array
                nw619_ = _dafny.Array(None, 15)
                nw619_[int(0)] = d_3421_v57_
                nw619_[int(1)] = d_3421_v57_
                nw619_[int(2)] = d_3421_v57_
                nw619_[int(3)] = d_3421_v57_
                nw619_[int(4)] = D17_DC45(d_3419_v55_, (d_3412_v50_)[default__.safeIndex(155, (d_3412_v50_).length(0))])
                nw619_[int(5)] = d_3421_v57_
                nw619_[int(6)] = d_3421_v57_
                nw619_[int(7)] = d_3421_v57_
                nw619_[int(8)] = D17_DC45(d_3419_v55_, (d_3412_v50_)[default__.safeIndex(155, (d_3412_v50_).length(0))])
                nw619_[int(9)] = d_3421_v57_
                nw619_[int(10)] = d_3421_v57_
                nw619_[int(11)] = d_3421_v57_
                nw619_[int(12)] = D17_DC45(_dafny.Map({(d_3412_v50_)[default__.safeIndex(155, (d_3412_v50_).length(0))]: d_3362_v7_}), (d_3412_v50_)[default__.safeIndex(155, (d_3412_v50_).length(0))])
                nw619_[int(13)] = d_3421_v57_
                nw619_[int(14)] = d_3421_v57_
                d_3422_v58_ = nw619_
                d_3423_v59_: D22
                d_3423_v59_ = D22_DC58(d_3422_v58_)
                d_3424_v60_: _dafny.Map
                d_3424_v60_ = _dafny.Map({d_3423_v59_: d_3362_v7_})
                d_3424_v60_ = (d_3424_v60_).set(d_3423_v59_, d_3362_v7_)
            elif True:
                d_3425_v61_: D15
                d_3425_v61_ = D15_DC38(d_3362_v7_, d_3356_v1_, (d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))])
                d_3426_v62_: C10
                nw620_ = C10()
                nw620_.ctor__(d_3359_v4_, -702)
                d_3426_v62_ = nw620_
                d_3427_v63_: _dafny.Map
                d_3427_v63_ = _dafny.Map({d_3425_v61_: d_3426_v62_})
                d_3427_v63_ = (d_3427_v63_).set(d_3425_v61_, d_3426_v62_)
                d_3428_v64_: _dafny.Array
                nw621_ = _dafny.Array(False, 6)
                d_3428_v64_ = nw621_
                d_3429_v65_: _dafny.MultiSet
                d_3429_v65_ = _dafny.MultiSet([d_3428_v64_, d_3428_v64_])
                d_3430_v66_: _dafny.MultiSet
                d_3430_v66_ = _dafny.MultiSet([True, d_3362_v7_])
                d_3431_v67_: _dafny.Seq
                d_3431_v67_ = _dafny.SeqWithoutIsStrInference([(d_3429_v65_).set(d_3428_v64_, default__.abs((d_3430_v66_).cardinality)), d_3429_v65_, d_3429_v65_])
                d_3432_v68_: D26
                d_3432_v68_ = D26_DC71(d_3429_v65_)
                d_3433_v69_: _dafny.Array
                nw622_ = _dafny.Array(None, 21)
                nw622_[int(0)] = d_3429_v65_
                nw622_[int(1)] = (_dafny.MultiSet([d_3428_v64_, d_3428_v64_])).intersection((d_3429_v65_).set(d_3428_v64_, default__.abs(d_3356_v1_)))
                nw622_[int(2)] = d_3429_v65_
                nw622_[int(3)] = d_3429_v65_
                nw622_[int(4)] = d_3429_v65_
                nw622_[int(5)] = d_3429_v65_
                nw622_[int(6)] = _dafny.MultiSet([d_3428_v64_])
                nw622_[int(7)] = d_3429_v65_
                nw622_[int(8)] = (d_3429_v65_) - (d_3429_v65_)
                nw622_[int(9)] = ((_dafny.MultiSet([d_3428_v64_])).set(d_3428_v64_, default__.abs((d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))])) if d_3362_v7_ else d_3429_v65_)
                nw622_[int(10)] = _dafny.MultiSet([d_3428_v64_, d_3428_v64_])
                nw622_[int(11)] = (d_3429_v65_).intersection(d_3429_v65_)
                nw622_[int(12)] = d_3429_v65_
                nw622_[int(13)] = (d_3431_v67_)[default__.safeIndex((d_3426_v62_).f32, len(d_3431_v67_))]
                nw622_[int(14)] = (d_3429_v65_).set(d_3428_v64_, default__.abs((d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))]))
                nw622_[int(15)] = d_3429_v65_
                nw622_[int(16)] = _dafny.MultiSet([d_3428_v64_])
                nw622_[int(17)] = (d_3432_v68_).cf110
                nw622_[int(18)] = d_3429_v65_
                nw622_[int(19)] = d_3429_v65_
                nw622_[int(20)] = (d_3429_v65_) - (d_3429_v65_)
                d_3433_v69_ = nw622_
                index563_ = default__.safeIndex(452, (d_3433_v69_).length(0))
                (d_3433_v69_)[index563_] = (d_3429_v65_).intersection(d_3429_v65_)
                index564_ = default__.safeIndex(452, (d_3433_v69_).length(0))
                (d_3433_v69_)[index564_] = (d_3429_v65_) - (d_3429_v65_)
                d_3434_v70_: _dafny.Set
                d_3434_v70_ = _dafny.Set({d_3428_v64_, d_3428_v64_})
                d_3435_v71_: _dafny.Seq
                d_3435_v71_ = _dafny.SeqWithoutIsStrInference([d_3434_v70_, _dafny.Set({d_3428_v64_, d_3428_v64_}), d_3434_v70_, d_3434_v70_, d_3434_v70_])
                index565_ = default__.safeIndex(461, (d_3360_v5_).length(0))
                rhs561_ = (d_3361_v6_)[default__.safeIndex(508, (d_3361_v6_).length(0))]
                rhs562_ = (0) - (default__.fm25((d_3426_v62_).f32, len((d_3435_v71_)[default__.safeIndex((d_3360_v5_)[default__.safeIndex(461, (d_3360_v5_).length(0))], len(d_3435_v71_))]), globalState))
                lhs481_ = globalState
                lhs482_ = d_3360_v5_
                lhs483_ = default__.safeIndex(461, (d_3360_v5_).length(0))
                lhs481_.f23 = rhs561_
                lhs482_[lhs483_] = rhs562_
                d_3366_v11_ = d_3366_v11_
                d_3436_v72_: _dafny.Set
                d_3436_v72_ = _dafny.Set({default__.fm51((d_3426_v62_).f32, (0) - (len(d_3367_v12_)), globalState)})
                d_3436_v72_ = d_3436_v72_
        _ingredientsd_0 = []
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, ((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0)):
            d_3437_i3_: int = guard_loop_9_
            if (True) and (((0) <= (d_3437_i3_)) and ((d_3437_i3_) < (((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]).length(0)))):
                _ingredientsd_0.append(((d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))], int(d_3437_i3_), default__.safeModuloInt(d_3437_i3_, (d_3356_v1_ if d_3362_v7_ else len(d_3364_v9_)))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_3438_v73_: _dafny.Seq
        d_3438_v73_ = _dafny.SeqWithoutIsStrInference([d_3360_v5_])
        d_3439_v74_: D27
        d_3439_v74_ = D27_DC74(d_3438_v73_)
        d_3440_v75_: T0
        nw623_ = C11()
        nw623_.ctor__(((d_3438_v73_) + ((d_3439_v74_).cf115)).set(default__.safeIndex(346, len((d_3438_v73_) + ((d_3439_v74_).cf115))), (d_3355_v0_)[default__.safeIndex(7, (d_3355_v0_).length(0))]), (d_3356_v1_) != (d_3356_v1_), d_3362_v7_)
        d_3440_v75_ = nw623_
        nw624_ = C7()
        nw624_.ctor__(d_3440_v75_.f28, (d_3440_v75_).f29)
        d_3440_v75_ = nw624_
        d_3441_v76_: C3
        nw625_ = C3()
        nw625_.ctor__(d_3362_v7_, (d_3440_v75_).f29, (len(d_3359_v4_)) != (d_3356_v1_), d_3356_v1_)
        d_3441_v76_ = nw625_
        r0 = d_3356_v1_
        r1 = d_3356_v1_
        r2 = (d_3362_v7_ if d_3440_v75_.f28 else d_3440_v75_.f28)
        d_3442_v77_: _dafny.Map
        d_3442_v77_ = _dafny.Map({(d_3441_v76_).f40: d_3356_v1_})
        d_3443_v78_: _dafny.MultiSet
        d_3443_v78_ = _dafny.MultiSet([d_3356_v1_, d_3356_v1_])
        d_3444_v79_: _dafny.Map
        d_3444_v79_ = _dafny.Map({d_3356_v1_: d_3443_v78_})
        r3 = (652) + (default__.safeModuloInt(((d_3442_v77_)[d_3440_v75_.f28] if (d_3440_v75_.f28) in (d_3442_v77_) else -83), (((d_3444_v79_)[d_3356_v1_] if (d_3356_v1_) in (d_3444_v79_) else d_3443_v78_)).cardinality))
        return r0, r1, r2, r3

