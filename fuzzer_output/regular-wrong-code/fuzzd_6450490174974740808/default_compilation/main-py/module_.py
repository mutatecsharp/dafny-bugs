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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Set
            for compr_0_ in (_dafny.MultiSet([_dafny.Set({True})])).Elements:
                d_0_v0_: _dafny.Set = compr_0_
                if (d_0_v0_) in (_dafny.MultiSet([_dafny.Set({True})])):
                    coll0_[d_0_v0_] = False
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Set
            for compr_1_ in (_dafny.MultiSet([_dafny.Set({False})])).Elements:
                d_1_v1_: _dafny.Set = compr_1_
                if (d_1_v1_) in (_dafny.MultiSet([_dafny.Set({False})])):
                    coll1_[d_1_v1_] = False
            return _dafny.Map(coll1_)
        return len(((iife0_()
        ) | (iife1_()
        )) | (_dafny.Map({_dafny.Set({False, not(not(True))}): not(False)})))

    @staticmethod
    def fm1(globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False, False, False, not(False)])))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brkfsxkwu"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2_i0_ in range(default__.abs(235))]))

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return ((_dafny.Map({False: len(_dafny.Map({True: True}))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykyytibe")))}))) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([749]))}))

    @staticmethod
    def fm8(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.MultiSet
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([972]), _dafny.MultiSet([169])])).Elements:
                d_3_v0_: _dafny.MultiSet = compr_2_
                if (d_3_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([972]), _dafny.MultiSet([169])])):
                    coll2_[d_3_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            return _dafny.Map(coll2_)
        return ((iife2_()
        ) | (_dafny.Map({(D19_DC39(_dafny.MultiSet([(_dafny.MultiSet([971, 664])).cardinality]))).cf60: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))}))) | (_dafny.Map({_dafny.MultiSet([-308]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkclbdj"))}))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({783}), _dafny.Set({290})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Set({-640})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({-318})])))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True)])) != (_dafny.SeqWithoutIsStrInference([not(False)]))

    @staticmethod
    def fm20(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.Set({-761})).Elements:
                d_4_v0_: int = compr_3_
                if (d_4_v0_) in (_dafny.Set({-761})):
                    coll3_[(d_4_v0_) + (949)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvghivqib")))
            return _dafny.Map(coll3_)
        return (_dafny.MultiSet([iife3_()
        , _dafny.Map({-361: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_5_i0_ in range(default__.abs(510))]))}), _dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, not(False)])): not(False)})): -874})])) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.Map({587: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbojidmhb")))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-55: (0) - (-384)})]))))

    @staticmethod
    def fm21(globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, True])])).Elements:
                d_6_v0_: _dafny.Seq = compr_4_
                if (d_6_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, True])])):
                    coll4_[d_6_v0_] = True
            return _dafny.Map(coll4_)
        return not(not(not (True) or (not((len(_dafny.SeqWithoutIsStrInference([False, True, not(True)]))) < (len(_dafny.Map({len(iife4_()
        ): not(True)})))))))

    @staticmethod
    def fm22(p0, globalState):
        return (_dafny.Map({(_dafny.MultiSet([344, 378, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utdhwcb")))])).cardinality: True})) | ((_dafny.Map({960: True})) | (_dafny.Map({(0) - (-518): False})))

    @staticmethod
    def fm25(p0, globalState):
        return ((_dafny.Map({not(False): -48})) | (_dafny.Map({False: -685}))) | ((_dafny.Map({True: (_dafny.MultiSet([465, len(_dafny.SeqWithoutIsStrInference([not(False)]))])).cardinality})) | (_dafny.Map({not(True): -478})))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: str
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])).Elements:
                d_7_v0_: str = compr_5_
                if (d_7_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])):
                    coll5_[d_7_v0_] = not(False)
            return _dafny.Map(coll5_)
        return (((D28_DC59(_dafny.Map({_dafny.CodePoint('e'): False}))).cf91 if not(False) else iife5_()
        )) | (_dafny.Map({_dafny.CodePoint('p'): True}))

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return not(not ((True if True else False)) or ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axni"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_8_i0_ in range(default__.abs(330))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_9_i1_ in range(default__.abs(293))])))

    @staticmethod
    def fm29(globalState):
        return D3_DC3(not(True))

    @staticmethod
    def fm30(p0, globalState):
        return (_dafny.MultiSet([409])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bm"))), -783]))

    @staticmethod
    def fm31(p0, p1, globalState):
        return not(((_dafny.SeqWithoutIsStrInference([214, len(_dafny.SeqWithoutIsStrInference([(0) - (-799)]))])) + (_dafny.SeqWithoutIsStrInference([928]))) == (_dafny.SeqWithoutIsStrInference([-772])))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        source0_ = D6_DC10()
        if source0_.is_DC10:
            return _dafny.Map({False: 733})
        elif True:
            d_10___mcc_h0_ = source0_.cf14
            d_11_cf14_ = d_10___mcc_h0_
            return _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olpebd")))})

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('c')

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: str
            for compr_6_ in (_dafny.Map({_dafny.CodePoint('d'): True})).keys.Elements:
                d_13_v0_: str = compr_6_
                if (d_13_v0_) in (_dafny.Map({_dafny.CodePoint('d'): True})):
                    coll6_[d_13_v0_] = False
            return _dafny.Map(coll6_)
        return (_dafny.SeqWithoutIsStrInference([len((D29_DC63(_dafny.Map({-409: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyixwxt"))}))).cf100) for d_12_i0_ in range(default__.abs(638))])) + (_dafny.SeqWithoutIsStrInference([len(iife6_()
        ), 800, 95, -392]))

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oihg")))), -997, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gekldkc")))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjogb"))) for d_14_i0_ in range(default__.abs(995))]))

    @staticmethod
    def fm39(globalState):
        return _dafny.CodePoint('r')

    @staticmethod
    def fm40(p0, globalState):
        return ((_dafny.Map({True: not(True)})) | (_dafny.Map({False: not(False)}))) | ((_dafny.Map({not(False): not(False)})) | (_dafny.Map({True: True})))

    @staticmethod
    def fm41(p0, p1, globalState):
        return (_dafny.Map({not(True): _dafny.Map({False: False})})) | (_dafny.Map({False: _dafny.Map({False: False})}))

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('u') if False else _dafny.CodePoint('r')) for d_15_i0_ in range(default__.abs(33))])

    @staticmethod
    def fm43(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(280, 197):
                d_16_v0_: int = compr_7_
                if ((280) <= (d_16_v0_)) and ((d_16_v0_) < (197)):
                    coll7_[default__.safeModuloInt(d_16_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcqe"))))] = _dafny.CodePoint('j')
            return _dafny.Map(coll7_)
        return not ((_dafny.Map({-928: False})) == (_dafny.Map({len(iife7_()
        ): False}))) or ((_dafny.MultiSet([_dafny.CodePoint('a')])).ispropersubset(_dafny.MultiSet([_dafny.CodePoint('m')])))

    @staticmethod
    def fm44(p0, p1, globalState):
        return (_dafny.Set({False, True})) - ((_dafny.Set({True, False, True})) - (_dafny.Set({False, False})))

    @staticmethod
    def fm45(p0, globalState):
        return _dafny.Set({_dafny.CodePoint('e')})

    @staticmethod
    def fm46(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: str
            for compr_8_ in (_dafny.Map({_dafny.CodePoint('d'): False})).keys.Elements:
                d_17_v0_: str = compr_8_
                if (d_17_v0_) in (_dafny.Map({_dafny.CodePoint('d'): False})):
                    coll8_[d_17_v0_] = 732
            return _dafny.Map(coll8_)
        return ((_dafny.Map({_dafny.CodePoint('u'): len(_dafny.SeqWithoutIsStrInference([False]))})) | (_dafny.Map({_dafny.CodePoint('q'): 822}))) | ((iife8_()
        ) | (_dafny.Map({_dafny.CodePoint('t'): len(_dafny.Map({False: False}))})))

    @staticmethod
    def fm47(p0, globalState):
        return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Map({_dafny.SeqWithoutIsStrInference([111]): len(_dafny.Set({not(False)}))})))

    @staticmethod
    def fm48(globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: _dafny.Map
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}) for d_18_i0_ in range(default__.abs(216))])).Elements:
                d_19_v0_: _dafny.Map = compr_9_
                if (d_19_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}) for d_18_i0_ in range(default__.abs(216))])):
                    coll9_[d_19_v0_] = 519
            return _dafny.Map(coll9_)
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: _dafny.Map
            for compr_10_ in (_dafny.Set({_dafny.Map({True: True})})).Elements:
                d_20_v1_: _dafny.Map = compr_10_
                if (d_20_v1_) in (_dafny.Set({_dafny.Map({True: True})})):
                    coll10_[d_20_v1_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fratralci")))
            return _dafny.Map(coll10_)
        return ((_dafny.Map({_dafny.Map({not(True): True}): 209})) | (iife9_()
        )) | ((_dafny.Map({_dafny.Map({True: True}): -941})) | (iife10_()
        ))

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(753, 841):
                d_21_v0_: int = compr_11_
                if ((753) <= (d_21_v0_)) and ((d_21_v0_) < (841)):
                    coll11_ = coll11_.union(_dafny.Set([(d_21_v0_) * (len(_dafny.Set({-474})))]))
            return _dafny.Set(coll11_)
        return (iife11_()
        ) - ((_dafny.Set({len(_dafny.Map({False: not(False)})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yattoed"))), -345, len(_dafny.SeqWithoutIsStrInference([756 for d_22_i0_ in range(default__.abs(445))])), -986})) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gx")))})))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: _dafny.MultiSet
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True, True, False, True]), _dafny.MultiSet([not(False)])])).Elements:
                d_23_v0_: _dafny.MultiSet = compr_12_
                if (d_23_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True, True, False, True]), _dafny.MultiSet([not(False)])])):
                    coll12_ = coll12_.union(_dafny.Set([d_23_v0_]))
            return _dafny.Set(coll12_)
        return ((_dafny.Set({_dafny.MultiSet([True])})) - (_dafny.Set({_dafny.MultiSet([not(True)]), _dafny.MultiSet([True])}))) - (iife12_()
        )

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return _dafny.Map({_dafny.Set({(_dafny.MultiSet([-37])).cardinality}): (False if False else True)})

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, False]))

    @staticmethod
    def fm53(p0, globalState):
        return D3_DC3(False)

    @staticmethod
    def fm54(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_24_i0_ in range(default__.abs(20))]))})])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): 994})]) if False else _dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): len(_dafny.Set({970}))}) for d_25_i1_ in range(default__.abs(353))])))

    @staticmethod
    def fm55(p0, p1, globalState):
        return D16_DC32((_dafny.MultiSet([True, True])).intersection(_dafny.MultiSet([True, True, False, True])))

    @staticmethod
    def fm56(p0, p1, globalState):
        return D14_DC27(_dafny.Map({_dafny.Map({False: _dafny.Map({False: _dafny.Map({False: True})})}): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpjbrs")))}))

    @staticmethod
    def fm57(globalState):
        return (_dafny.Set({_dafny.Map({False: not(not(False))})})).intersection(_dafny.Set({_dafny.Map({not(not(not(False))): True}), _dafny.Map({False: False}), _dafny.Map({False: False})}))

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        if False:
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uua"))): D4_DC7(_dafny.SeqWithoutIsStrInference([True]), len(_dafny.Map({len(_dafny.Map({320: _dafny.Set({False})})): len(_dafny.Map({581: len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True), False])): False})): _dafny.CodePoint('y')}))}))})))})) | (_dafny.Map({-772: D4_DC7(_dafny.SeqWithoutIsStrInference([False, not(False)]), 953)}))
        elif not(False):
            return _dafny.Map({504: D4_DC7(_dafny.SeqWithoutIsStrInference([not(True)]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lakodemt"))))})
        elif True:
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmiwuxmt"))): D4_DC7(_dafny.SeqWithoutIsStrInference([True]), -934)})

    @staticmethod
    def fm59(globalState):
        return D4_DC7(_dafny.SeqWithoutIsStrInference([True]), len(_dafny.Set({not(False), not(True), False, False})))

    @staticmethod
    def fm60(p0, p1, p2, globalState):
        return D11_DC20()

    @staticmethod
    def fm61(p0, p1, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: _dafny.Seq
            for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lt")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkt"))])).Elements:
                d_26_v0_: _dafny.Seq = compr_13_
                if (d_26_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lt")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkt"))])):
                    coll13_[d_26_v0_] = True
            return _dafny.Map(coll13_)
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: _dafny.Seq
            for compr_14_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_27_i0_ in range(default__.abs(863))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mh"))])).Elements:
                d_28_v1_: _dafny.Seq = compr_14_
                if (d_28_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_27_i0_ in range(default__.abs(863))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mh"))])):
                    coll14_[d_28_v1_] = not(False)
            return _dafny.Map(coll14_)
        return ((iife13_()
        ) | (iife14_()
        )) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "issb")): True})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_29_i1_ in range(default__.abs(373))]): True})))

    @staticmethod
    def fm62(globalState):
        return ((_dafny.SeqWithoutIsStrInference([D3_DC3(False), D3_DC3(False)])) + (_dafny.SeqWithoutIsStrInference([D3_DC3(False) for d_30_i0_ in range(default__.abs(926))]))) + (_dafny.SeqWithoutIsStrInference([D3_DC3(False) for d_31_i1_ in range(default__.abs(928))]))

    @staticmethod
    def fm63(globalState):
        def iife15_():
            coll15_ = _dafny.Set()
            compr_15_: int
            for compr_15_ in _dafny.IntegerRange(504, 758):
                d_32_v0_: int = compr_15_
                if ((504) <= (d_32_v0_)) and ((d_32_v0_) < (758)):
                    coll15_ = coll15_.union(_dafny.Set([(d_32_v0_) * (len(_dafny.SeqWithoutIsStrInference([True])))]))
            return _dafny.Set(coll15_)
        return D19_DC39((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([584, len((D28_DC60(40, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([377, 135, (0) - ((_dafny.MultiSet([False, True])).cardinality)])), len(iife15_()
), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfsij")))}))).cf93)]))) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_33_i0_ in range(default__.abs(339))]))])))

    @staticmethod
    def fm64(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: _dafny.CodePoint('m')})) | (_dafny.Map({True: _dafny.CodePoint('e')}))) | ((_dafny.Map({False: _dafny.CodePoint('b')})) | (_dafny.Map({True: _dafny.CodePoint('i')})))

    @staticmethod
    def fm65(p0, globalState):
        return D3_DC4((_dafny.Map({_dafny.CodePoint('c'): False})) == (_dafny.Map({_dafny.CodePoint('c'): not(not(False))})), len(_dafny.Set({False, not(not(False))})), default__.safeModuloInt(445, 979), (len(_dafny.Map({137: not(True)}))) < ((_dafny.MultiSet([_dafny.MultiSet([True, not(True)]), _dafny.MultiSet([True, False])])).cardinality), (_dafny.MultiSet([False, not(False), not(True), True, False])).ispropersubset(_dafny.MultiSet([not(False), (D16_DC33(False, False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_34_i0_ in range(default__.abs(287))])), (_dafny.MultiSet([False])).cardinality])), 262])), False)).cf48])))

    @staticmethod
    def fm66(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([False, not(not(False))]))): -355})})

    @staticmethod
    def fm67(p0, globalState):
        return D7_DC12()

    @staticmethod
    def fm68(p0, p1, globalState):
        return D6_DC9(_dafny.CodePoint('h'))

    @staticmethod
    def fm69(p0, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: _dafny.Seq
            for compr_16_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vioord")): len(_dafny.SeqWithoutIsStrInference([(D6_DC9(_dafny.CodePoint('k'))).cf14 for d_35_i0_ in range(default__.abs(763))]))})).keys.Elements:
                d_36_v0_: _dafny.Seq = compr_16_
                if (d_36_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vioord")): len(_dafny.SeqWithoutIsStrInference([(D6_DC9(_dafny.CodePoint('k'))).cf14 for d_35_i0_ in range(default__.abs(763))]))})):
                    coll16_[d_36_v0_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgrs")))
            return _dafny.Map(coll16_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: _dafny.Seq
            for compr_17_ in (_dafny.Map({(D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oixmk")))).cf29: _dafny.CodePoint('m')})).keys.Elements:
                d_37_v1_: _dafny.Seq = compr_17_
                if (d_37_v1_) in (_dafny.Map({(D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oixmk")))).cf29: _dafny.CodePoint('m')})):
                    coll17_[d_37_v1_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "illhgq")))
            return _dafny.Map(coll17_)
        return (iife16_()
        ) | ((D13_DC25(iife17_()
)).cf38)

    @staticmethod
    def fm70(p0, p1, p2, globalState):
        return _dafny.Map({862: len(_dafny.Set({True}))})

    @staticmethod
    def fm71(p0, p1, p2, p3, globalState):
        return D16_DC33(False, False, (D25_DC52(791)).cf79, not(False))

    @staticmethod
    def m16(p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: int = int(0)
        d_38_i0_: int
        d_38_i0_ = 0
        with _dafny.label("0"):
            while (p1) < (p1):
                with _dafny.c_label("0"):
                    if (d_38_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_38_i0_ = (d_38_i0_) + (1)
                    d_39_v0_: _dafny.Map
                    d_39_v0_ = _dafny.Map({p3: True})
                    d_40_v1_: _dafny.MultiSet
                    d_40_v1_ = _dafny.MultiSet([p3, ((d_39_v0_)[p3] if (p3) in (d_39_v0_) else p3), p3, p3, p3])
                    d_41_v2_: _dafny.Seq
                    d_41_v2_ = _dafny.SeqWithoutIsStrInference([True, True])
                    d_42_v3_: _dafny.Set
                    d_42_v3_ = _dafny.Set({not(p3)})
                    d_43_v4_: _dafny.MultiSet
                    d_43_v4_ = _dafny.MultiSet([p1])
                    d_44_v5_: _dafny.Map
                    d_44_v5_ = _dafny.Map({default__.fm39(globalState): p3})
                    d_45_v6_: _dafny.Array
                    nw0_ = _dafny.Array(None, 8)
                    nw0_[int(0)] = p1
                    nw0_[int(1)] = p1
                    nw0_[int(2)] = (d_43_v4_).cardinality
                    nw0_[int(3)] = p1
                    nw0_[int(4)] = p1
                    nw0_[int(5)] = p1
                    nw0_[int(6)] = p1
                    nw0_[int(7)] = len(d_44_v5_)
                    d_45_v6_ = nw0_
                    d_46_v7_: _dafny.Map
                    d_46_v7_ = _dafny.Map({len(d_42_v3_): d_45_v6_})
                    d_47_v8_: _dafny.Array
                    nw1_ = _dafny.Array(None, 16)
                    nw1_[int(0)] = p3
                    nw1_[int(1)] = (d_40_v1_).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p3])))
                    nw1_[int(2)] = p3
                    nw1_[int(3)] = p3
                    nw1_[int(4)] = not(p3)
                    nw1_[int(5)] = (d_41_v2_)[default__.safeIndex(-168, len(d_41_v2_))]
                    nw1_[int(6)] = p3
                    nw1_[int(7)] = p3
                    nw1_[int(8)] = p3
                    nw1_[int(9)] = p3
                    nw1_[int(10)] = p3
                    nw1_[int(11)] = p3
                    nw1_[int(12)] = p3
                    nw1_[int(13)] = (p1) not in (d_46_v7_)
                    nw1_[int(14)] = not (p3) or (p3)
                    nw1_[int(15)] = p3
                    d_47_v8_ = nw1_
                    index0_ = default__.safeIndex(485, (d_47_v8_).length(0))
                    (d_47_v8_)[index0_] = p3
                    index1_ = default__.safeIndex(485, (d_47_v8_).length(0))
                    rhs0_ = False
                    rhs1_ = (0) - ((0) - (p1))
                    lhs0_ = d_47_v8_
                    lhs1_ = default__.safeIndex(485, (d_47_v8_).length(0))
                    lhs2_ = globalState
                    lhs0_[lhs1_] = rhs0_
                    lhs2_.f7 = rhs1_
                    d_48_v9_: str
                    d_48_v9_ = _dafny.CodePoint('f')
                    d_48_v9_ = d_48_v9_
                    d_49_v10_: _dafny.Array
                    nw2_ = _dafny.Array(_dafny.Array(None, 0), 6)
                    d_49_v10_ = nw2_
                    d_50_v11_: _dafny.Map
                    d_50_v11_ = _dafny.Map({p1: p3})
                    d_51_v12_: _dafny.Seq
                    d_51_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouvug"))
                    d_52_v13_: C3
                    nw3_ = C3()
                    nw3_.ctor__(len(d_50_v11_), d_51_v12_, not((d_47_v8_)[default__.safeIndex(485, (d_47_v8_).length(0))]), p3)
                    d_52_v13_ = nw3_
                    d_53_v14_: _dafny.Array
                    nw4_ = _dafny.Array(None, 7)
                    nw4_[int(0)] = d_52_v13_
                    nw4_[int(1)] = d_52_v13_
                    nw4_[int(2)] = d_52_v13_
                    nw4_[int(3)] = d_52_v13_
                    nw4_[int(4)] = d_52_v13_
                    nw4_[int(5)] = d_52_v13_
                    nw4_[int(6)] = d_52_v13_
                    d_53_v14_ = nw4_
                    index2_ = default__.safeIndex(439, (d_49_v10_).length(0))
                    (d_49_v10_)[index2_] = d_53_v14_
                    index3_ = default__.safeIndex(439, (d_49_v10_).length(0))
                    (d_49_v10_)[index3_] = d_53_v14_
                    d_54_v15_: D3
                    d_54_v15_ = D3_DC3(False)
                    d_55_v16_: C2
                    nw5_ = C2()
                    nw5_.ctor__(d_54_v15_, (d_47_v8_)[default__.safeIndex(485, (d_47_v8_).length(0))], p3)
                    d_55_v16_ = nw5_
                    d_56_v17_: _dafny.Seq
                    d_56_v17_ = _dafny.SeqWithoutIsStrInference([d_55_v16_])
                    index4_ = default__.safeIndex(485, (d_47_v8_).length(0))
                    (d_47_v8_)[index4_] = (not((False if default__.fm21(globalState) else not((d_47_v8_)[default__.safeIndex(485, (d_47_v8_).length(0))]))) if (p1) < (len(d_56_v17_)) else False)
                    pass
            pass
        d_57_v18_: _dafny.Array
        nw6_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_57_v18_ = nw6_
        d_57_v18_ = d_57_v18_
        d_58_v19_: _dafny.Array
        def lambda0_(d_59_i1_):
            return False

        init0_ = lambda0_
        nw7_ = _dafny.Array(None, 23)
        for i0_0_ in range(nw7_.length(0)):
            nw7_[i0_0_] = init0_(i0_0_)
        d_58_v19_ = nw7_
        d_60_v20_: _dafny.Seq
        d_60_v20_ = _dafny.SeqWithoutIsStrInference([d_58_v19_])
        d_60_v20_ = _dafny.SeqWithoutIsStrInference([d_58_v19_, d_58_v19_, d_58_v19_, d_58_v19_, d_58_v19_])
        d_61_v21_: C13
        nw8_ = C13()
        nw8_.ctor__()
        d_61_v21_ = nw8_
        d_62_v22_: bool
        d_62_v22_ = False
        d_62_v22_ = default__.fm17(d_62_v22_, p1, len((p2) + (_dafny.SeqWithoutIsStrInference([default__.fm0(not(p3), globalState)]))), p1, globalState)
        d_63_v23_: _dafny.Seq
        d_63_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfxq"))
        d_64_v24_: C3
        nw9_ = C3()
        nw9_.ctor__(p1, d_63_v23_, True, (D3_DC3(not(True))).cf3)
        d_64_v24_ = nw9_
        d_65_v25_: _dafny.Array
        nw10_ = _dafny.Array(_dafny.Set({}), 25)
        d_65_v25_ = nw10_
        d_66_v26_: C0
        nw11_ = C0()
        nw11_.ctor__(d_65_v25_, p1)
        d_66_v26_ = nw11_
        d_67_v27_: _dafny.Map
        d_67_v27_ = _dafny.Map({d_64_v24_: d_66_v26_})
        d_68_v28_: C12
        nw12_ = C12()
        nw12_.ctor__((d_64_v24_) not in (d_67_v27_), d_66_v26_.f27)
        d_68_v28_ = nw12_
        r0 = d_58_v19_
        r1 = (d_63_v23_).set(default__.safeIndex(p1, len(d_63_v23_)), p0)
        d_69_v29_: _dafny.MultiSet
        d_69_v29_ = _dafny.MultiSet([p3])
        d_70_v30_: _dafny.Map
        d_70_v30_ = _dafny.Map({p1: d_69_v29_})
        r2 = (((d_69_v29_) | (((d_70_v30_)[p1] if (p1) in (d_70_v30_) else d_69_v29_))) - (d_69_v29_)).cardinality
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_71_v0_: _dafny.Array
        def lambda1_(d_72_i0_):
            return True

        init1_ = lambda1_
        nw13_ = _dafny.Array(None, 11)
        for i0_1_ in range(nw13_.length(0)):
            nw13_[i0_1_] = init1_(i0_1_)
        d_71_v0_ = nw13_
        d_73_v1_: _dafny.Array
        def lambda2_(d_74_i1_):
            return (d_74_i1_) * (616)

        init2_ = lambda2_
        nw14_ = _dafny.Array(None, 18)
        for i0_2_ in range(nw14_.length(0)):
            nw14_[i0_2_] = init2_(i0_2_)
        d_73_v1_ = nw14_
        d_75_v2_: _dafny.MultiSet
        d_75_v2_ = _dafny.MultiSet([d_73_v1_, d_73_v1_, d_73_v1_])
        d_76_globalState_: GlobalState
        nw15_ = GlobalState()
        nw15_.ctor__(150, -360, 288, 359, True, 773, False, 4, 763, 290, True, d_71_v0_, d_75_v2_, 105)
        d_76_globalState_ = nw15_
        d_77_v3_: bool
        d_77_v3_ = True
        d_78_v4_: _dafny.Array
        nw16_ = _dafny.Array(_dafny.Seq({}), 5)
        d_78_v4_ = nw16_
        d_79_v5_: int
        d_79_v5_ = 63
        index5_ = default__.safeIndex(310, (d_78_v4_).length(0))
        (d_78_v4_)[index5_] = (_dafny.SeqWithoutIsStrInference([d_79_v5_])) + (_dafny.SeqWithoutIsStrInference([d_79_v5_, d_79_v5_]))
        d_80_v6_: _dafny.Map
        d_80_v6_ = _dafny.Map({d_79_v5_: d_77_v3_})
        d_81_v7_: _dafny.Map
        d_81_v7_ = _dafny.Map({d_77_v3_: d_77_v3_})
        d_82_v8_: _dafny.MultiSet
        d_82_v8_ = _dafny.MultiSet([len(d_81_v7_), d_79_v5_])
        d_83_v9_: _dafny.Seq
        d_83_v9_ = _dafny.SeqWithoutIsStrInference([(0) - (d_79_v5_), (d_82_v8_).cardinality, (0) - (d_79_v5_), d_79_v5_, d_79_v5_])
        d_84_v10_: str
        d_84_v10_ = _dafny.CodePoint('c')
        d_85_v11_: _dafny.Map
        d_85_v11_ = _dafny.Map({d_79_v5_: d_84_v10_})
        index6_ = default__.safeIndex(310, (d_78_v4_).length(0))
        rhs2_ = ((d_80_v6_)[(0) - (d_79_v5_)] if ((0) - (d_79_v5_)) in (d_80_v6_) else d_77_v3_)
        rhs3_ = (d_79_v5_) - (d_79_v5_)
        rhs4_ = d_77_v3_
        rhs5_ = (d_83_v9_) + ((_dafny.SeqWithoutIsStrInference([len(d_85_v11_)])).set(default__.safeIndex(d_79_v5_, len(_dafny.SeqWithoutIsStrInference([len(d_85_v11_)]))), d_79_v5_))
        rhs6_ = (((_dafny.MultiSet(d_83_v9_)).cardinality) != (d_79_v5_) if d_77_v3_ else d_77_v3_)
        lhs3_ = d_76_globalState_
        lhs4_ = d_78_v4_
        lhs5_ = default__.safeIndex(310, (d_78_v4_).length(0))
        d_77_v3_ = rhs2_
        lhs3_.f7 = rhs3_
        d_77_v3_ = rhs4_
        lhs4_[lhs5_] = rhs5_
        d_77_v3_ = rhs6_
        d_86_v12_: _dafny.Seq
        d_86_v12_ = _dafny.SeqWithoutIsStrInference([d_77_v3_, d_77_v3_, d_77_v3_])
        d_87_v13_: _dafny.Seq
        d_87_v13_ = _dafny.SeqWithoutIsStrInference([d_86_v12_])
        d_79_v5_ = len((d_87_v13_) + (_dafny.SeqWithoutIsStrInference([d_86_v12_])))
        d_88_v14_: _dafny.Seq
        d_88_v14_ = _dafny.SeqWithoutIsStrInference([d_73_v1_])
        d_89_v15_: _dafny.Set
        d_89_v15_ = _dafny.Set({d_77_v3_, d_77_v3_})
        d_90_v16_: _dafny.Map
        d_90_v16_ = _dafny.Map({d_73_v1_: d_89_v15_})
        d_91_v17_: _dafny.MultiSet
        d_91_v17_ = _dafny.MultiSet([((d_88_v14_)[default__.safeIndex(d_79_v5_, len(d_88_v14_))]) in (d_90_v16_), not(d_77_v3_), False])
        d_92_v18_: _dafny.Map
        d_92_v18_ = _dafny.Map({default__.fm0(d_77_v3_, d_76_globalState_): d_91_v17_})
        d_93_v19_: _dafny.Seq
        d_93_v19_ = _dafny.SeqWithoutIsStrInference([d_91_v17_, _dafny.MultiSet([d_77_v3_, d_77_v3_]), d_91_v17_])
        d_91_v17_ = (((d_92_v18_)[-359] if (-359) in (d_92_v18_) else (d_93_v19_)[default__.safeIndex(d_79_v5_, len(d_93_v19_))])) - ((d_91_v17_) | (default__.fm1(d_76_globalState_)))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_73_v1_).length(0)):
            d_94_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_94_i2_)) and ((d_94_i2_) < ((d_73_v1_).length(0)))):
                (d_73_v1_)[(d_94_i2_)] = default__.safeDivisionInt(d_94_i2_, d_79_v5_)
        if False:
            d_95_v20_: C13
            nw17_ = C13()
            nw17_.ctor__()
            d_95_v20_ = nw17_
            d_77_v3_ = (default__.fm31(d_77_v3_, d_77_v3_, d_76_globalState_)) and (d_77_v3_)
            d_96_v21_: C5
            nw18_ = C5()
            nw18_.ctor__(d_77_v3_, d_77_v3_)
            d_96_v21_ = nw18_
            d_96_v21_ = (d_96_v21_ if not(d_77_v3_) else d_96_v21_)
            d_97_v22_: _dafny.Map
            d_97_v22_ = _dafny.Map({d_79_v5_: d_79_v5_})
            d_98_v23_: C10
            nw19_ = C10()
            nw19_.ctor__(d_77_v3_, d_77_v3_, len(d_97_v22_))
            d_98_v23_ = nw19_
            d_79_v5_ = (d_79_v5_) - (d_79_v5_)
        elif True:
            index7_ = default__.safeIndex(905, (d_71_v0_).length(0))
            (d_71_v0_)[index7_] = not(False)
            index8_ = default__.safeIndex(905, (d_71_v0_).length(0))
            rhs7_ = d_84_v10_
            rhs8_ = not(not (False) or (d_77_v3_))
            rhs9_ = not(True)
            lhs6_ = d_71_v0_
            lhs7_ = default__.safeIndex(905, (d_71_v0_).length(0))
            d_84_v10_ = rhs7_
            d_77_v3_ = rhs8_
            lhs6_[lhs7_] = rhs9_
            d_99_v24_: _dafny.Array
            nw20_ = _dafny.Array(D16.default()(), 13)
            d_99_v24_ = nw20_
            index9_ = default__.safeIndex(572, (d_99_v24_).length(0))
            (d_99_v24_)[index9_] = D16_DC33(d_77_v3_, False, d_79_v5_, d_77_v3_)
            d_100_v25_: _dafny.Map
            d_100_v25_ = _dafny.Map({d_79_v5_: d_79_v5_})
            index10_ = default__.safeIndex(572, (d_99_v24_).length(0))
            (d_99_v24_)[index10_] = default__.fm71(d_100_v25_, default__.fm52(not(d_77_v3_), d_79_v5_, d_79_v5_, d_79_v5_, d_76_globalState_), 741, (d_89_v15_) | (_dafny.Set({not(d_77_v3_)})), d_76_globalState_)
            if ((d_81_v7_)[(d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]] if ((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]) in (d_81_v7_) else (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]):
                d_101_v26_: _dafny.Array
                nw21_ = _dafny.Array(_dafny.Seq({}), 28)
                d_101_v26_ = nw21_
                d_102_v27_: _dafny.Seq
                d_102_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxwae"))
                d_103_v28_: _dafny.Map
                d_103_v28_ = _dafny.Map({d_102_v27_: len((_dafny.SeqWithoutIsStrInference([d_84_v10_ for d_104_i3_ in range(default__.abs(916))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwfis"))), len(_dafny.SeqWithoutIsStrInference([d_84_v10_ for d_105_i3_ in range(default__.abs(916))]))), d_84_v10_))})
                d_106_v29_: D13
                d_106_v29_ = D13_DC25(d_103_v28_)
                d_107_v30_: _dafny.Seq
                d_107_v30_ = _dafny.SeqWithoutIsStrInference([D13_DC25(d_103_v28_), d_106_v29_])
                index11_ = default__.safeIndex(973, (d_101_v26_).length(0))
                (d_101_v26_)[index11_] = d_107_v30_
                index12_ = default__.safeIndex(973, (d_101_v26_).length(0))
                (d_101_v26_)[index12_] = d_107_v30_
                d_77_v3_ = False
                d_108_v31_: _dafny.Array
                def lambda3_(d_109_i4_):
                    return True

                init3_ = lambda3_
                nw22_ = _dafny.Array(None, 9)
                for i0_3_ in range(nw22_.length(0)):
                    nw22_[i0_3_] = init3_(i0_3_)
                d_108_v31_ = nw22_
                d_108_v31_ = d_108_v31_
                d_103_v28_ = (d_103_v28_).set(d_102_v27_, d_79_v5_)
                d_110_v32_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_110_v32_ = nw23_
                d_111_v33_: D25
                d_111_v33_ = D25_DC51(d_110_v32_)
                d_110_v32_ = (d_110_v32_ if (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))] else (d_111_v33_).cf78)
            elif True:
                index13_ = default__.safeIndex(794, (d_73_v1_).length(0))
                (d_73_v1_)[index13_] = d_79_v5_
                index14_ = default__.safeIndex(794, (d_73_v1_).length(0))
                rhs10_ = (d_79_v5_) - ((0) - (d_79_v5_))
                rhs11_ = (931) >= (d_79_v5_)
                rhs12_ = d_81_v7_
                lhs8_ = d_73_v1_
                lhs9_ = default__.safeIndex(794, (d_73_v1_).length(0))
                lhs8_[lhs9_] = rhs10_
                d_77_v3_ = rhs11_
                d_81_v7_ = rhs12_
                d_112_v34_: C13
                nw24_ = C13()
                nw24_.ctor__()
                d_112_v34_ = nw24_
                d_112_v34_ = d_112_v34_
                d_79_v5_ = default__.safeDivisionInt(d_79_v5_, d_79_v5_)
                d_113_v35_: _dafny.Array
                nw25_ = _dafny.Array(None, 29)
                nw25_[int(0)] = d_77_v3_
                nw25_[int(1)] = not(d_77_v3_)
                nw25_[int(2)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(3)] = d_77_v3_
                nw25_[int(4)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(5)] = d_77_v3_
                nw25_[int(6)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(7)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(8)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(9)] = d_77_v3_
                nw25_[int(10)] = False
                nw25_[int(11)] = d_77_v3_
                nw25_[int(12)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(13)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(14)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(15)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(16)] = default__.fm31(d_77_v3_, d_77_v3_, d_76_globalState_)
                nw25_[int(17)] = default__.fm43((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], d_76_globalState_)
                nw25_[int(18)] = d_77_v3_
                nw25_[int(19)] = not((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))])
                nw25_[int(20)] = True
                nw25_[int(21)] = d_77_v3_
                nw25_[int(22)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(23)] = d_77_v3_
                nw25_[int(24)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(25)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(26)] = (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]
                nw25_[int(27)] = d_77_v3_
                nw25_[int(28)] = d_77_v3_
                d_113_v35_ = nw25_
                d_114_v36_: _dafny.Seq
                d_114_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                d_115_v37_: _dafny.Map
                d_115_v37_ = _dafny.Map({d_113_v35_: d_114_v36_})
                d_116_v38_: D12
                d_116_v38_ = D12_DC22(d_114_v36_)
                d_115_v37_ = (d_115_v37_).set(d_113_v35_, ((d_116_v38_).cf29) + (_dafny.SeqWithoutIsStrInference([d_84_v10_ for d_117_i5_ in range(default__.abs(15))])))
                d_118_v39_: _dafny.Array
                nw26_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_118_v39_ = nw26_
                d_119_v40_: _dafny.Array
                nw27_ = _dafny.Array(None, 3)
                d_119_v40_ = nw27_
                index15_ = default__.safeIndex(503, (d_118_v39_).length(0))
                (d_118_v39_)[index15_] = d_119_v40_
                index16_ = default__.safeIndex(503, (d_118_v39_).length(0))
                (d_118_v39_)[index16_] = d_119_v40_
            d_120_v41_: _dafny.Seq
            d_120_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfsifht"))
            if not((((d_120_v41_ if (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))] else d_120_v41_)).set(default__.safeIndex(d_79_v5_, len((d_120_v41_ if (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))] else d_120_v41_))), d_84_v10_)) <= (d_120_v41_)):
                d_121_v42_: C13
                nw28_ = C13()
                nw28_.ctor__()
                d_121_v42_ = nw28_
                d_122_v43_: _dafny.MultiSet
                d_122_v43_ = _dafny.MultiSet([d_87_v13_, d_87_v13_, d_87_v13_])
                d_123_v44_: _dafny.Seq
                d_123_v44_ = _dafny.SeqWithoutIsStrInference([d_77_v3_, d_77_v3_])
                d_124_v45_: D4
                d_124_v45_ = D4_DC7(d_123_v44_, d_79_v5_)
                (d_76_globalState_).f2 = ((d_122_v43_)[_dafny.SeqWithoutIsStrInference([d_86_v12_])] if (_dafny.SeqWithoutIsStrInference([d_86_v12_])) in (d_122_v43_) else (0) - (((d_124_v45_).cf12) * (d_79_v5_)))
                d_80_v6_ = (d_80_v6_).set(default__.fm0((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], d_76_globalState_), (_dafny.SeqWithoutIsStrInference([d_84_v10_ for d_125_i6_ in range(default__.abs(602))])) == (d_120_v41_))
                d_80_v6_ = (d_80_v6_).set((d_79_v5_) - (len(d_83_v9_)), d_77_v3_)
                d_126_v46_: _dafny.Set
                d_126_v46_ = _dafny.Set({d_79_v5_, d_79_v5_, 442})
                d_127_v47_: _dafny.Array
                nw29_ = _dafny.Array(False, 14)
                d_127_v47_ = nw29_
                d_128_v48_: D9
                d_128_v48_ = D9_DC15((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], d_126_v46_, d_73_v1_, d_127_v47_)
                pat_let_tv0_ = d_73_v1_
                pat_let_tv1_ = d_127_v47_
                pat_let_tv2_ = d_127_v47_
                pat_let_tv3_ = d_73_v1_
                d_129_v50_: _dafny.Array
                nw30_ = _dafny.Array(None, 14)
                nw30_[int(0)] = d_128_v48_
                nw30_[int(1)] = d_128_v48_
                def iife18_(_pat_let0_0):
                    def iife19_(d_130_dt__update__tmp_h0_):
                        def iife20_(_pat_let1_0):
                            def iife21_(d_131_dt__update_hcf21_h0_):
                                def iife22_(_pat_let2_0):
                                    def iife23_(d_132_dt__update_hcf22_h0_):
                                        return D9_DC15((d_130_dt__update__tmp_h0_).cf18, (d_130_dt__update__tmp_h0_).cf19, (d_130_dt__update__tmp_h0_).cf20, d_131_dt__update_hcf21_h0_, d_132_dt__update_hcf22_h0_)
                                    return iife23_(_pat_let2_0)
                                return iife22_(pat_let_tv1_)
                            return iife21_(_pat_let1_0)
                        return iife20_(pat_let_tv0_)
                    return iife19_(_pat_let0_0)
                nw30_[int(2)] = iife18_(D9_DC15(d_77_v3_, (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], d_126_v46_, d_73_v1_, d_127_v47_))
                def iife24_():
                    coll18_ = _dafny.Set()
                    compr_18_: int
                    for compr_18_ in (d_100_v25_).keys.Elements:
                        d_133_v49_: int = compr_18_
                        if (d_133_v49_) in (d_100_v25_):
                            coll18_ = coll18_.union(_dafny.Set([default__.safeModuloInt(d_133_v49_, len(_dafny.Set({True, True, False, False, not(True)})))]))
                    return _dafny.Set(coll18_)
                nw30_[int(3)] = D9_DC15((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], default__.fm43(False, d_76_globalState_), iife24_()
, d_73_v1_, d_127_v47_)
                nw30_[int(4)] = d_128_v48_
                nw30_[int(5)] = d_128_v48_
                nw30_[int(6)] = d_128_v48_
                nw30_[int(7)] = d_128_v48_
                nw30_[int(8)] = d_128_v48_
                nw30_[int(9)] = d_128_v48_
                def iife25_(_pat_let3_0):
                    def iife26_(d_134_dt__update__tmp_h1_):
                        def iife27_(_pat_let4_0):
                            def iife28_(d_135_dt__update_hcf22_h1_):
                                def iife29_(_pat_let5_0):
                                    def iife30_(d_136_dt__update_hcf19_h0_):
                                        def iife31_(_pat_let6_0):
                                            def iife32_(d_137_dt__update_hcf21_h1_):
                                                return D9_DC15((d_134_dt__update__tmp_h1_).cf18, d_136_dt__update_hcf19_h0_, (d_134_dt__update__tmp_h1_).cf20, d_137_dt__update_hcf21_h1_, d_135_dt__update_hcf22_h1_)
                                            return iife32_(_pat_let6_0)
                                        return iife31_(pat_let_tv3_)
                                    return iife30_(_pat_let5_0)
                                return iife29_(False)
                            return iife28_(_pat_let4_0)
                        return iife27_(pat_let_tv2_)
                    return iife26_(_pat_let3_0)
                nw30_[int(10)] = iife25_(d_128_v48_)
                nw30_[int(11)] = d_128_v48_
                nw30_[int(12)] = d_128_v48_
                nw30_[int(13)] = d_128_v48_
                d_129_v50_ = nw30_
                d_138_v51_: D18
                d_138_v51_ = D18_DC37(d_129_v50_)
                d_138_v51_ = d_138_v51_
            elif True:
                d_120_v41_ = ((d_120_v41_) + (d_120_v41_)) + (d_120_v41_)
                d_139_v52_: _dafny.Array
                nw31_ = _dafny.Array(None, 25)
                d_139_v52_ = nw31_
                d_140_v53_: _dafny.Array
                nw32_ = _dafny.Array(None, 26)
                nw32_[int(0)] = _dafny.Set({d_77_v3_})
                nw32_[int(1)] = d_89_v15_
                nw32_[int(2)] = _dafny.Set({d_77_v3_})
                nw32_[int(3)] = d_89_v15_
                nw32_[int(4)] = _dafny.Set({(d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]})
                nw32_[int(5)] = d_89_v15_
                nw32_[int(6)] = d_89_v15_
                nw32_[int(7)] = d_89_v15_
                nw32_[int(8)] = _dafny.Set({d_77_v3_, False, d_77_v3_, False, not(d_77_v3_)})
                nw32_[int(9)] = d_89_v15_
                nw32_[int(10)] = d_89_v15_
                nw32_[int(11)] = d_89_v15_
                nw32_[int(12)] = d_89_v15_
                nw32_[int(13)] = d_89_v15_
                nw32_[int(14)] = _dafny.Set({d_77_v3_})
                nw32_[int(15)] = d_89_v15_
                nw32_[int(16)] = d_89_v15_
                nw32_[int(17)] = _dafny.Set({d_77_v3_})
                nw32_[int(18)] = d_89_v15_
                nw32_[int(19)] = d_89_v15_
                nw32_[int(20)] = d_89_v15_
                nw32_[int(21)] = d_89_v15_
                nw32_[int(22)] = default__.fm44(d_77_v3_, False, d_76_globalState_)
                nw32_[int(23)] = d_89_v15_
                nw32_[int(24)] = d_89_v15_
                nw32_[int(25)] = d_89_v15_
                d_140_v53_ = nw32_
                d_141_v54_: C0
                nw33_ = C0()
                nw33_.ctor__(d_140_v53_, len((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))]))
                d_141_v54_ = nw33_
                index17_ = default__.safeIndex(482, (d_139_v52_).length(0))
                (d_139_v52_)[index17_] = d_141_v54_
                index18_ = default__.safeIndex(482, (d_139_v52_).length(0))
                (d_139_v52_)[index18_] = d_141_v54_
                d_142_v55_: _dafny.Array
                nw34_ = _dafny.Array(False, 24)
                d_142_v55_ = nw34_
                d_143_v56_: D11
                d_143_v56_ = D11_DC19(d_100_v25_)
                d_144_v57_: _dafny.Map
                d_144_v57_ = _dafny.Map({d_142_v55_: d_143_v56_})
                d_144_v57_ = (d_144_v57_).set(d_142_v55_, d_143_v56_)
                d_145_v59_: _dafny.Set
                def iife33_():
                    coll19_ = _dafny.Set()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(137, 987):
                        d_146_v58_: int = compr_19_
                        if ((137) <= (d_146_v58_)) and ((d_146_v58_) < (987)):
                            coll19_ = coll19_.union(_dafny.Set([(d_146_v58_) - (d_141_v54_.f27)]))
                    return _dafny.Set(coll19_)
                d_145_v59_ = _dafny.Set({len(iife33_()
                )})
                d_147_v60_: C2
                nw35_ = C2()
                nw35_.ctor__(D3_DC3(d_77_v3_), (d_145_v59_).isdisjoint(d_145_v59_), d_77_v3_)
                d_147_v60_ = nw35_
                index19_ = default__.safeIndex(905, (d_71_v0_).length(0))
                index20_ = default__.safeIndex(905, (d_71_v0_).length(0))
                rhs13_ = d_77_v3_
                rhs14_ = d_77_v3_
                rhs15_ = (d_79_v5_) != (d_79_v5_)
                rhs16_ = d_77_v3_
                rhs17_ = d_147_v60_
                lhs10_ = d_71_v0_
                lhs11_ = default__.safeIndex(905, (d_71_v0_).length(0))
                lhs12_ = d_71_v0_
                lhs13_ = default__.safeIndex(905, (d_71_v0_).length(0))
                d_77_v3_ = rhs13_
                lhs10_[lhs11_] = rhs14_
                d_77_v3_ = rhs15_
                lhs12_[lhs13_] = rhs16_
                d_147_v60_ = rhs17_
                d_148_v61_: _dafny.Map
                d_148_v61_ = _dafny.Map({d_77_v3_: ((0) - (len(d_83_v9_))) - (697)})
                d_149_v62_: D12
                d_149_v62_ = D12_DC22(d_120_v41_)
                d_150_v63_: _dafny.Seq
                d_150_v63_ = _dafny.SeqWithoutIsStrInference([d_149_v62_, D12_DC22(d_120_v41_), d_149_v62_])
                d_79_v5_ = ((d_148_v61_)[((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]) not in (d_148_v61_)] if (((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]) not in (d_148_v61_)) in (d_148_v61_) else len((d_150_v63_ if d_77_v3_ else d_150_v63_)))
            if d_77_v3_:
                d_151_v64_: C9
                nw36_ = C9()
                nw36_.ctor__()
                d_151_v64_ = nw36_
                d_152_v65_: _dafny.Set
                d_152_v65_ = _dafny.Set({d_151_v64_, d_151_v64_, d_151_v64_, d_151_v64_, d_151_v64_})
                d_153_v66_: C6
                nw37_ = C6()
                nw37_.ctor__(d_79_v5_, default__.fm43((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))], d_76_globalState_), d_77_v3_, (_dafny.Set({d_151_v64_, d_151_v64_})).ispropersubset(d_152_v65_), ((0) - ((0) - (d_79_v5_))) * (len(d_83_v9_)))
                d_153_v66_ = nw37_
                d_154_v67_: C10
                nw38_ = C10()
                nw38_.ctor__((d_153_v66_).f20, not(((d_153_v66_).f20) and (True)), (d_153_v66_).f19)
                d_154_v67_ = nw38_
                d_155_v68_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Set({}), 11)
                d_155_v68_ = nw39_
                d_156_v69_: C0
                nw40_ = C0()
                nw40_.ctor__(d_155_v68_, d_79_v5_)
                d_156_v69_ = nw40_
                d_157_v70_: _dafny.Map
                d_157_v70_ = _dafny.Map({d_79_v5_: d_156_v69_})
                d_158_v71_: _dafny.Set
                d_158_v71_ = _dafny.Set({d_79_v5_, (d_153_v66_).f19, len(d_120_v41_), len(d_120_v41_), len(d_157_v70_)})
                d_159_v72_: _dafny.Map
                d_159_v72_ = _dafny.Map({(d_153_v66_).f20: d_158_v71_})
                d_160_v73_: C11
                nw41_ = C11()
                nw41_.ctor__(d_156_v69_.f27, (d_153_v66_).f20, (d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))])
                d_160_v73_ = nw41_
                index21_ = default__.safeIndex(905, (d_71_v0_).length(0))
                rhs18_ = d_159_v72_
                rhs19_ = (d_153_v66_).f20
                rhs20_ = d_79_v5_
                rhs21_ = d_160_v73_
                rhs22_ = (d_81_v7_) | (d_81_v7_)
                lhs14_ = d_71_v0_
                lhs15_ = default__.safeIndex(905, (d_71_v0_).length(0))
                d_159_v72_ = rhs18_
                lhs14_[lhs15_] = rhs19_
                d_79_v5_ = rhs20_
                d_160_v73_ = rhs21_
                d_81_v7_ = rhs22_
                d_77_v3_ = (d_120_v41_) < ((d_154_v67_).fm9(d_79_v5_, _dafny.SeqWithoutIsStrInference([d_158_v71_, d_158_v71_]), _dafny.CodePoint('s'), d_156_v69_.f27, d_76_globalState_))
                d_161_v74_: _dafny.Map
                d_161_v74_ = _dafny.Map({(d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]: 24})
                d_161_v74_ = (d_161_v74_).set(((d_71_v0_)[default__.safeIndex(905, (d_71_v0_).length(0))]) and ((d_153_v66_).f20), default__.safeDivisionInt(d_156_v69_.f27, len(_dafny.Set({d_156_v69_.f27}))))
            elif True:
                d_162_v75_: C13
                nw42_ = C13()
                nw42_.ctor__()
                d_162_v75_ = nw42_
                d_163_v76_: _dafny.MultiSet
                d_163_v76_ = _dafny.MultiSet([d_162_v75_])
                d_164_v77_: _dafny.Seq
                d_164_v77_ = _dafny.SeqWithoutIsStrInference([d_162_v75_])
                d_77_v3_ = (d_163_v76_).isdisjoint(_dafny.MultiSet(d_164_v77_))
                d_165_v78_: _dafny.Map
                d_165_v78_ = _dafny.Map({d_79_v5_: d_120_v41_})
                d_77_v3_ = (d_120_v41_) != ((((d_165_v78_)[964] if (964) in (d_165_v78_) else d_120_v41_)) + (_dafny.SeqWithoutIsStrInference([d_84_v10_ for d_166_i7_ in range(default__.abs(68))])))
                (d_76_globalState_).f2 = d_79_v5_
                index22_ = default__.safeIndex(905, (d_71_v0_).length(0))
                (d_71_v0_)[index22_] = (default__.safeModuloInt(-188, d_79_v5_)) != ((d_79_v5_ if d_77_v3_ else d_79_v5_))
                d_167_v79_: C13
                nw43_ = C13()
                nw43_.ctor__()
                d_167_v79_ = nw43_
        def lambda4_(d_168_v5_):
            def lambda5_(d_169_i8_):
                return (d_169_i8_) + (d_168_v5_)

            return lambda5_

        init4_ = lambda4_(d_79_v5_)
        nw44_ = _dafny.Array(None, 5)
        for i0_4_ in range(nw44_.length(0)):
            nw44_[i0_4_] = init4_(i0_4_)
        d_73_v1_ = nw44_
        source1_ = D25_DC53()
        if source1_.is_DC52:
            d_170___mcc_h0_ = source1_.cf79
            d_171_cf79_ = d_170___mcc_h0_
            d_77_v3_ = d_77_v3_
            d_77_v3_ = d_77_v3_
            d_77_v3_ = ((0) - (((d_82_v8_)[d_171_cf79_] if (d_171_cf79_) in (d_82_v8_) else d_79_v5_))) < (d_171_cf79_)
            d_172_v80_: _dafny.Seq
            d_172_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_173_v81_: _dafny.Map
            d_173_v81_ = _dafny.Map({d_77_v3_: (0) - (len(d_172_v80_))})
            d_174_v82_: _dafny.MultiSet
            d_174_v82_ = _dafny.MultiSet([default__.fm38(d_77_v3_, d_171_cf79_, d_171_cf79_, d_76_globalState_), _dafny.SeqWithoutIsStrInference([len(d_173_v81_), d_79_v5_])])
            d_80_v6_ = (_dafny.Map({len((d_172_v80_).set(default__.safeIndex((d_174_v82_).cardinality, len(d_172_v80_)), d_84_v10_)): d_77_v3_})) | (d_80_v6_)
        elif source1_.is_DC53:
            d_175_v83_: T2
            nw45_ = C7()
            nw45_.ctor__(d_77_v3_, (d_77_v3_ if d_77_v3_ else d_77_v3_), 997)
            d_175_v83_ = nw45_
            d_175_v83_ = d_175_v83_
            (d_175_v83_).f16 = not (True) or (d_77_v3_)
            (d_175_v83_).f16 = default__.fm17(d_175_v83_.f16, default__.safeModuloInt(d_79_v5_, d_79_v5_), (d_79_v5_) * (159), len(_dafny.Map({d_79_v5_: d_79_v5_})), d_76_globalState_)
            if not(d_77_v3_):
                d_176_v84_: D21
                d_176_v84_ = D21_DC44(d_79_v5_, not(not((d_175_v83_).f17)))
                d_177_v85_: _dafny.Seq
                d_177_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewiw"))
                d_178_v86_: T0
                nw46_ = C6()
                nw46_.ctor__(d_79_v5_, d_77_v3_, (d_175_v83_).f17, (d_175_v83_).f17, (0) - (d_79_v5_))
                d_178_v86_ = nw46_
                d_179_v87_: _dafny.Set
                d_179_v87_ = _dafny.Set({d_178_v86_})
                d_180_v88_: _dafny.Map
                d_180_v88_ = _dafny.Map({(d_175_v83_).f17: len(d_179_v87_)})
                d_181_v89_: int
                d_182_v90_: bool
                d_183_v91_: _dafny.Seq
                d_184_v92_: bool
                out0_: int
                out1_: bool
                out2_: _dafny.Seq
                out3_: bool
                out0_, out1_, out2_, out3_ = (d_175_v83_).m6((d_175_v83_).f17, (d_176_v84_).cf70, default__.safeDivisionInt(776, d_79_v5_), default__.fm34(len(d_177_v85_), len(d_180_v88_), (d_175_v83_).f17, -499, d_76_globalState_), d_76_globalState_)
                d_181_v89_ = out0_
                d_182_v90_ = out1_
                d_183_v91_ = out2_
                d_184_v92_ = out3_
                d_185_v93_: _dafny.MultiSet
                d_185_v93_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_182_v90_])])
                d_186_v94_: C11
                nw47_ = C11()
                nw47_.ctor__(d_181_v89_, (d_185_v93_) == (d_185_v93_), not(d_184_v92_))
                d_186_v94_ = nw47_
                (d_76_globalState_).f2 = 158
                d_187_v95_: _dafny.Map
                d_187_v95_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_188_i9_ in range(default__.abs(339))])): len(d_86_v12_)})
                d_189_v101_: D11
                d_189_v101_ = D11_DC19(d_187_v95_)
                d_190_v102_: _dafny.Array
                nw48_ = _dafny.Array(None, 22)
                nw48_[int(0)] = (d_187_v95_ if d_175_v83_.f16 else d_187_v95_)
                nw48_[int(1)] = _dafny.Map({(d_91_v17_).cardinality: d_178_v86_.f14})
                nw48_[int(2)] = (_dafny.Map({d_79_v5_: len(d_187_v95_)})) | (d_187_v95_)
                nw48_[int(3)] = _dafny.Map({d_178_v86_.f14: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjhe")))})
                def iife34_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(156, 112):
                        d_191_v96_: int = compr_20_
                        if ((156) <= (d_191_v96_)) and ((d_191_v96_) < (112)):
                            coll20_[(d_191_v96_) + (d_181_v89_)] = (0) - (d_181_v89_)
                    return _dafny.Map(coll20_)
                def iife35_():
                    coll21_ = _dafny.Map()
                    compr_21_: int
                    for compr_21_ in ((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))]).Elements:
                        d_192_v97_: int = compr_21_
                        if (d_192_v97_) in ((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))]):
                            coll21_[default__.safeDivisionInt(d_192_v97_, d_178_v86_.f14)] = len(_dafny.SeqWithoutIsStrInference([d_175_v83_.f16, d_77_v3_, (d_175_v83_).f17]))
                    return _dafny.Map(coll21_)
                nw48_[int(4)] = (iife34_()
                ) | (iife35_()
                )
                nw48_[int(5)] = (d_187_v95_) | (d_187_v95_)
                def iife36_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in (d_83_v9_).Elements:
                        d_193_v98_: int = compr_22_
                        if (d_193_v98_) in (d_83_v9_):
                            coll22_[(d_193_v98_) - (d_181_v89_)] = d_178_v86_.f14
                    return _dafny.Map(coll22_)
                nw48_[int(6)] = iife36_()
                
                nw48_[int(7)] = _dafny.Map({d_178_v86_.f14: d_181_v89_})
                def iife37_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in ((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))]).Elements:
                        d_194_v99_: int = compr_23_
                        if (d_194_v99_) in ((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))]):
                            coll23_[default__.safeModuloInt(d_194_v99_, d_178_v86_.f14)] = d_181_v89_
                    return _dafny.Map(coll23_)
                nw48_[int(8)] = (d_187_v95_) | (iife37_()
                )
                nw48_[int(9)] = (d_187_v95_) | (d_187_v95_)
                nw48_[int(10)] = d_187_v95_
                nw48_[int(11)] = (d_187_v95_) | (d_187_v95_)
                nw48_[int(12)] = d_187_v95_
                def iife38_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(986, 286):
                        d_195_v100_: int = compr_24_
                        if ((986) <= (d_195_v100_)) and ((d_195_v100_) < (286)):
                            coll24_[(d_195_v100_) * (len(d_187_v95_))] = d_178_v86_.f14
                    return _dafny.Map(coll24_)
                nw48_[int(13)] = iife38_()
                
                nw48_[int(14)] = d_187_v95_
                nw48_[int(15)] = _dafny.Map({d_181_v89_: d_181_v89_})
                nw48_[int(16)] = (d_187_v95_) | (_dafny.Map({len(d_83_v9_): d_79_v5_}))
                nw48_[int(17)] = (d_189_v101_).cf27
                nw48_[int(18)] = d_187_v95_
                nw48_[int(19)] = (d_175_v83_).fm3(d_76_globalState_)
                nw48_[int(20)] = d_187_v95_
                nw48_[int(21)] = d_187_v95_
                d_190_v102_ = nw48_
                d_196_v104_: _dafny.Set
                d_196_v104_ = _dafny.Set({(0) - (len(d_177_v85_))})
                d_197_v106_: _dafny.Map
                d_197_v106_ = _dafny.Map({(d_175_v83_).f17: (d_189_v101_).cf27})
                d_198_v107_: _dafny.Array
                nw49_ = _dafny.Array(D9.default()(), 15)
                d_198_v107_ = nw49_
                d_199_v108_: _dafny.Seq
                d_199_v108_ = _dafny.SeqWithoutIsStrInference([d_198_v107_, d_198_v107_])
                d_200_v109_: D12
                d_200_v109_ = D12_DC24(d_184_v92_, d_86_v12_, d_199_v108_, (d_175_v83_).f17, _dafny.SeqWithoutIsStrInference([D3_DC3(d_175_v83_.f16) for d_201_i10_ in range(default__.abs(852))]))
                nw50_ = _dafny.Array(None, 26)
                nw50_[int(0)] = _dafny.Map({d_79_v5_: d_181_v89_})
                nw50_[int(1)] = d_187_v95_
                nw50_[int(2)] = d_187_v95_
                nw50_[int(3)] = d_187_v95_
                nw50_[int(4)] = d_187_v95_
                nw50_[int(5)] = d_187_v95_
                nw50_[int(6)] = d_187_v95_
                def iife39_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in (d_196_v104_).Elements:
                        d_202_v103_: int = compr_25_
                        if (d_202_v103_) in (d_196_v104_):
                            coll25_[(d_202_v103_) - (d_181_v89_)] = (0) - ((0) - (d_178_v86_.f14))
                    return _dafny.Map(coll25_)
                nw50_[int(7)] = iife39_()
                
                nw50_[int(8)] = (d_187_v95_) | (d_187_v95_)
                nw50_[int(9)] = d_187_v95_
                nw50_[int(10)] = d_187_v95_
                def iife40_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(157, -96):
                        d_203_v105_: int = compr_26_
                        if ((157) <= (d_203_v105_)) and ((d_203_v105_) < (-96)):
                            coll26_[default__.safeDivisionInt(d_203_v105_, d_178_v86_.f14)] = d_178_v86_.f14
                    return _dafny.Map(coll26_)
                nw50_[int(11)] = iife40_()
                
                nw50_[int(12)] = d_187_v95_
                nw50_[int(13)] = (d_187_v95_) | (_dafny.Map({d_181_v89_: d_181_v89_}))
                nw50_[int(14)] = (d_187_v95_) | (d_187_v95_)
                nw50_[int(15)] = d_187_v95_
                nw50_[int(16)] = d_187_v95_
                nw50_[int(17)] = d_187_v95_
                nw50_[int(18)] = d_187_v95_
                nw50_[int(19)] = d_187_v95_
                nw50_[int(20)] = (d_187_v95_) | (((d_197_v106_)[(d_200_v109_).cf36] if ((d_200_v109_).cf36) in (d_197_v106_) else d_187_v95_))
                nw50_[int(21)] = (d_187_v95_) | (d_187_v95_)
                nw50_[int(22)] = d_187_v95_
                nw50_[int(23)] = (_dafny.Map({d_79_v5_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rceafxpqm")))})) | (d_187_v95_)
                nw50_[int(24)] = _dafny.Map({d_178_v86_.f14: d_178_v86_.f14})
                nw50_[int(25)] = d_187_v95_
                d_190_v102_ = nw50_
                index23_ = default__.safeIndex(102, (d_73_v1_).length(0))
                (d_73_v1_)[index23_] = d_79_v5_
                d_204_v110_: _dafny.Map
                d_204_v110_ = _dafny.Map({False: d_79_v5_})
                index24_ = default__.safeIndex(102, (d_73_v1_).length(0))
                rhs23_ = (0) - ((d_175_v83_).fm12(727, d_76_globalState_))
                rhs24_ = default__.fm21(d_76_globalState_)
                rhs25_ = d_178_v86_
                rhs26_ = ((d_187_v95_)[d_178_v86_.f14] if (d_178_v86_.f14) in (d_187_v95_) else (0) - (d_79_v5_))
                rhs27_ = d_204_v110_
                lhs16_ = d_76_globalState_
                lhs17_ = d_73_v1_
                lhs18_ = default__.safeIndex(102, (d_73_v1_).length(0))
                lhs16_.f2 = rhs23_
                d_77_v3_ = rhs24_
                d_178_v86_ = rhs25_
                lhs17_[lhs18_] = rhs26_
                d_204_v110_ = rhs27_
            elif True:
                d_205_v111_: _dafny.Map
                d_205_v111_ = _dafny.Map({d_175_v83_.f16: -58})
                d_206_v112_: _dafny.Seq
                d_206_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_207_v113_: C3
                nw51_ = C3()
                nw51_.ctor__(352, d_206_v112_, d_175_v83_.f16, d_175_v83_.f16)
                d_207_v113_ = nw51_
                d_208_v114_: _dafny.Map
                d_208_v114_ = _dafny.Map({d_205_v111_: d_207_v113_})
                d_209_v115_: _dafny.Seq
                d_209_v115_ = _dafny.SeqWithoutIsStrInference([((d_208_v114_)[d_205_v111_] if (d_205_v111_) in (d_208_v114_) else d_207_v113_), d_207_v113_])
                (d_76_globalState_).f7 = len(d_209_v115_)
                d_210_v116_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                d_210_v116_ = nw52_
                d_211_v117_: C4
                nw53_ = C4()
                nw53_.ctor__(True, (d_175_v83_).f17, d_175_v83_.f16)
                d_211_v117_ = nw53_
                d_212_v118_: _dafny.Map
                d_212_v118_ = _dafny.Map({d_211_v117_: d_210_v116_})
                d_210_v116_ = ((d_212_v118_)[d_211_v117_] if (d_211_v117_) in (d_212_v118_) else d_210_v116_)
                rhs28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkc"))
                rhs29_ = (d_86_v12_)[default__.safeIndex(default__.safeDivisionInt((d_207_v113_).f22, (d_207_v113_).f22), len(d_86_v12_))]
                rhs30_ = d_175_v83_.f16
                lhs19_ = d_175_v83_
                d_206_v112_ = rhs28_
                d_77_v3_ = rhs29_
                lhs19_.f16 = rhs30_
                d_213_v119_: D21
                d_213_v119_ = D21_DC44(d_79_v5_, d_175_v83_.f16)
                d_214_v120_: _dafny.Map
                d_214_v120_ = _dafny.Map({(d_207_v113_).f22: d_213_v119_})
                d_215_v121_: C7
                nw54_ = C7()
                nw54_.ctor__(True, (len(d_214_v120_)) >= (((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))])[default__.safeIndex(d_79_v5_, len((d_78_v4_)[default__.safeIndex(310, (d_78_v4_).length(0))]))]), (d_207_v113_).f22)
                d_215_v121_ = nw54_
                index25_ = default__.safeIndex(245, (d_73_v1_).length(0))
                (d_73_v1_)[index25_] = ((0) - (d_79_v5_)) * (len(d_83_v9_))
                index26_ = default__.safeIndex(245, (d_73_v1_).length(0))
                (d_73_v1_)[index26_] = (0) - ((d_207_v113_).f22)
        elif source1_.is_DC51:
            d_216___mcc_h1_ = source1_.cf78
            d_217_cf78_ = d_216___mcc_h1_
            d_218_v122_: _dafny.Seq
            d_218_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "it"))
            d_219_v123_: C8
            nw55_ = C8()
            nw55_.ctor__(d_218_v122_, d_77_v3_, d_77_v3_, 569)
            d_219_v123_ = nw55_
            d_220_v124_: _dafny.Seq
            d_220_v124_ = _dafny.SeqWithoutIsStrInference([d_219_v123_, d_219_v123_, d_219_v123_, d_219_v123_])
            d_221_v125_: _dafny.Set
            d_221_v125_ = _dafny.Set({d_219_v123_, d_219_v123_, d_219_v123_, d_219_v123_, (d_220_v124_)[default__.safeIndex(d_79_v5_, len(d_220_v124_))]})
            d_222_v126_: _dafny.Set
            d_222_v126_ = _dafny.Set({len(d_81_v7_)})
            d_223_v127_: D9
            d_223_v127_ = D9_DC15((d_219_v123_) in (d_221_v125_), d_77_v3_, d_222_v126_, d_73_v1_, d_71_v0_)
            d_223_v127_ = d_223_v127_
            (d_76_globalState_).f2 = d_79_v5_
            d_224_v128_: D19
            d_224_v128_ = D19_DC39(_dafny.MultiSet([d_79_v5_]))
            d_225_v129_: _dafny.Map
            d_225_v129_ = _dafny.Map({d_224_v128_: d_77_v3_})
            d_226_v130_: D26
            d_226_v130_ = D26_DC55(d_225_v129_)
            d_227_v131_: _dafny.Array
            nw56_ = _dafny.Array(None, 15)
            nw56_[int(0)] = d_225_v129_
            nw56_[int(1)] = (d_226_v130_).cf81
            nw56_[int(2)] = d_225_v129_
            nw56_[int(3)] = d_225_v129_
            nw56_[int(4)] = d_225_v129_
            nw56_[int(5)] = (d_225_v129_) | (d_225_v129_)
            nw56_[int(6)] = d_225_v129_
            nw56_[int(7)] = (d_225_v129_).set(d_224_v128_, d_77_v3_)
            nw56_[int(8)] = (d_225_v129_) | (_dafny.Map({d_224_v128_: d_77_v3_}))
            nw56_[int(9)] = d_225_v129_
            nw56_[int(10)] = d_225_v129_
            nw56_[int(11)] = d_225_v129_
            nw56_[int(12)] = d_225_v129_
            nw56_[int(13)] = d_225_v129_
            nw56_[int(14)] = d_225_v129_
            d_227_v131_ = nw56_
            d_228_v133_: _dafny.Seq
            def iife41_():
                coll27_ = _dafny.Map()
                compr_27_: D19
                for compr_27_ in (_dafny.MultiSet([d_224_v128_])).Elements:
                    d_229_v132_: D19 = compr_27_
                    if (d_229_v132_) in (_dafny.MultiSet([d_224_v128_])):
                        coll27_[d_229_v132_] = d_77_v3_
                return _dafny.Map(coll27_)
            d_228_v133_ = _dafny.SeqWithoutIsStrInference([iife41_()
            ])
            index27_ = default__.safeIndex(676, (d_227_v131_).length(0))
            (d_227_v131_)[index27_] = ((d_228_v133_)[default__.safeIndex((d_219_v123_).fm12(default__.fm0(d_77_v3_, d_76_globalState_), d_76_globalState_), len(d_228_v133_))]) | (d_225_v129_)
            d_230_v134_: C5
            nw57_ = C5()
            nw57_.ctor__((not(d_77_v3_) if d_77_v3_ else d_77_v3_), (d_77_v3_) or (d_77_v3_))
            d_230_v134_ = nw57_
            index28_ = default__.safeIndex(676, (d_227_v131_).length(0))
            rhs31_ = d_225_v129_
            rhs32_ = (d_79_v5_) != ((d_79_v5_) + (len(d_218_v122_)))
            rhs33_ = d_230_v134_
            lhs20_ = d_227_v131_
            lhs21_ = default__.safeIndex(676, (d_227_v131_).length(0))
            lhs20_[lhs21_] = rhs31_
            d_77_v3_ = rhs32_
            d_230_v134_ = rhs33_
            rhs34_ = d_77_v3_
            rhs35_ = (0) - (default__.safeDivisionInt((d_219_v123_).fm12(d_79_v5_, d_76_globalState_), (len(_dafny.SeqWithoutIsStrInference([d_77_v3_])) if False else d_79_v5_)))
            rhs36_ = d_77_v3_
            lhs22_ = d_76_globalState_
            d_77_v3_ = rhs34_
            lhs22_.f7 = rhs35_
            d_77_v3_ = rhs36_
        elif True:
            d_231___mcc_h2_ = source1_.cf80
            d_232_cf80_ = d_231___mcc_h2_
            d_77_v3_ = (d_77_v3_ if d_77_v3_ else d_77_v3_)
            if d_77_v3_:
                d_77_v3_ = d_77_v3_
                d_233_v135_: _dafny.Map
                d_233_v135_ = _dafny.Map({not(d_77_v3_): d_79_v5_})
                index29_ = default__.safeIndex(355, (d_73_v1_).length(0))
                (d_73_v1_)[index29_] = default__.safeDivisionInt(d_79_v5_, ((d_233_v135_)[d_77_v3_] if (d_77_v3_) in (d_233_v135_) else d_79_v5_))
                index30_ = default__.safeIndex(355, (d_73_v1_).length(0))
                (d_73_v1_)[index30_] = d_79_v5_
                (d_76_globalState_).f2 = ((d_73_v1_)[default__.safeIndex(355, (d_73_v1_).length(0))]) - (default__.safeModuloInt(d_79_v5_, (d_73_v1_)[default__.safeIndex(355, (d_73_v1_).length(0))]))
                d_234_v136_: _dafny.Array
                def lambda6_(d_235_v7_):
                    def lambda7_(d_236_i11_):
                        return (d_236_i11_) + (len(d_235_v7_))

                    return lambda7_

                init5_ = lambda6_(d_81_v7_)
                nw58_ = _dafny.Array(None, 21)
                for i0_5_ in range(nw58_.length(0)):
                    nw58_[i0_5_] = init5_(i0_5_)
                d_234_v136_ = nw58_
                d_237_v137_: _dafny.Map
                d_237_v137_ = _dafny.Map({(0) - (d_79_v5_): d_234_v136_})
                d_238_v138_: _dafny.Map
                d_238_v138_ = _dafny.Map({default__.safeModuloInt(default__.fm0(d_77_v3_, d_76_globalState_), len(d_83_v9_)): ((d_237_v137_)[(d_73_v1_)[default__.safeIndex(355, (d_73_v1_).length(0))]] if ((d_73_v1_)[default__.safeIndex(355, (d_73_v1_).length(0))]) in (d_237_v137_) else d_234_v136_)})
                d_239_v139_: T1
                nw59_ = C6()
                nw59_.ctor__(478, not(d_77_v3_), d_77_v3_, default__.fm21(d_76_globalState_), d_79_v5_)
                d_239_v139_ = nw59_
                d_240_v140_: _dafny.Map
                d_240_v140_ = _dafny.Map({d_239_v139_: d_238_v138_})
                d_237_v137_ = ((d_240_v140_)[d_239_v139_] if (d_239_v139_) in (d_240_v140_) else d_237_v137_)
                d_241_v141_: C13
                nw60_ = C13()
                nw60_.ctor__()
                d_241_v141_ = nw60_
            elif True:
                d_242_v142_: _dafny.Set
                d_242_v142_ = _dafny.Set({d_79_v5_, d_79_v5_})
                d_243_v143_: _dafny.Map
                d_243_v143_ = _dafny.Map({d_77_v3_: d_242_v142_})
                d_244_v144_: _dafny.Array
                d_245_v145_: _dafny.Seq
                d_246_v146_: int
                out4_: _dafny.Array
                out5_: _dafny.Seq
                out6_: int
                out4_, out5_, out6_ = default__.m16(d_84_v10_, default__.safeModuloInt(d_79_v5_, d_79_v5_), d_83_v9_, (len(d_243_v143_)) > (len((d_86_v12_).set(default__.safeIndex(d_79_v5_, len(d_86_v12_)), d_77_v3_))), d_76_globalState_)
                d_244_v144_ = out4_
                d_245_v145_ = out5_
                d_246_v146_ = out6_
                d_77_v3_ = d_77_v3_
                d_77_v3_ = d_77_v3_
                d_77_v3_ = True
                d_247_v147_: _dafny.Map
                d_247_v147_ = _dafny.Map({not(d_77_v3_): d_79_v5_})
                d_248_v148_: _dafny.Map
                d_248_v148_ = _dafny.Map({d_247_v147_: default__.safeModuloInt(d_246_v146_, d_246_v146_)})
                d_248_v148_ = (d_248_v148_).set(d_247_v147_, len(_dafny.SeqWithoutIsStrInference([((d_91_v17_)[d_77_v3_] if (d_77_v3_) in (d_91_v17_) else -741) for d_249_i12_ in range(default__.abs(244))])))
            d_77_v3_ = d_77_v3_
            d_250_v149_: _dafny.Array
            nw61_ = _dafny.Array(None, 17)
            d_250_v149_ = nw61_
            d_251_v150_: C12
            nw62_ = C12()
            nw62_.ctor__(d_77_v3_, 827)
            d_251_v150_ = nw62_
            index31_ = default__.safeIndex(398, (d_250_v149_).length(0))
            (d_250_v149_)[index31_] = d_251_v150_
            d_252_v151_: _dafny.Map
            d_252_v151_ = _dafny.Map({(0) - (d_79_v5_): d_251_v150_})
            index32_ = default__.safeIndex(398, (d_250_v149_).length(0))
            (d_250_v149_)[index32_] = ((d_252_v151_)[default__.safeModuloInt(d_79_v5_, d_79_v5_)] if (default__.safeModuloInt(d_79_v5_, d_79_v5_)) in (d_252_v151_) else d_251_v150_)
        index33_ = default__.safeIndex(179, (d_73_v1_).length(0))
        (d_73_v1_)[index33_] = (default__.fm0(d_77_v3_, d_76_globalState_) if d_77_v3_ else 17)
        index34_ = default__.safeIndex(179, (d_73_v1_).length(0))
        (d_73_v1_)[index34_] = d_79_v5_
        d_253_v152_: D16
        d_253_v152_ = D16_DC32(d_91_v17_)
        hi0_ = d_79_v5_
        for d_254_i13_ in range(((d_253_v152_).cf46).cardinality, hi0_):
            d_255_v153_: _dafny.Seq
            d_255_v153_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbx"))
            d_256_v154_: C3
            nw63_ = C3()
            nw63_.ctor__(d_254_i13_, d_255_v153_, d_77_v3_, d_77_v3_)
            d_256_v154_ = nw63_
            d_257_v155_: _dafny.Map
            d_257_v155_ = _dafny.Map({d_256_v154_: d_254_i13_})
            d_258_v156_: C7
            nw64_ = C7()
            nw64_.ctor__(d_77_v3_, d_77_v3_, 932)
            d_258_v156_ = nw64_
            d_259_v157_: _dafny.Set
            d_259_v157_ = _dafny.Set({d_258_v156_})
            d_260_v158_: _dafny.Map
            d_260_v158_ = _dafny.Map({((d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]) + (((d_257_v155_)[d_256_v154_] if (d_256_v154_) in (d_257_v155_) else d_79_v5_)): (d_259_v157_) - (d_259_v157_)})
            index35_ = default__.safeIndex(179, (d_73_v1_).length(0))
            rhs37_ = (0) - (d_79_v5_)
            rhs38_ = len(d_260_v158_)
            lhs23_ = d_73_v1_
            lhs24_ = default__.safeIndex(179, (d_73_v1_).length(0))
            lhs25_ = d_76_globalState_
            lhs23_[lhs24_] = rhs37_
            lhs25_.f7 = rhs38_
            d_261_v159_: _dafny.Set
            d_261_v159_ = _dafny.Set({-72})
            d_262_v160_: D9
            d_262_v160_ = D9_DC15(not(d_77_v3_), d_77_v3_, d_261_v159_, d_73_v1_, d_71_v0_)
            d_263_v161_: _dafny.Seq
            d_263_v161_ = _dafny.SeqWithoutIsStrInference([d_262_v160_, d_262_v160_])
            d_264_v162_: _dafny.Map
            d_264_v162_ = _dafny.Map({(d_77_v3_ if d_77_v3_ else d_77_v3_): d_263_v161_})
            d_264_v162_ = (d_264_v162_).set(d_77_v3_, ((d_263_v161_) + (d_263_v161_)).set(default__.safeIndex((d_256_v154_).f22, len((d_263_v161_) + (d_263_v161_))), d_262_v160_))
            d_71_v0_ = d_71_v0_
            d_265_v163_: C13
            nw65_ = C13()
            nw65_.ctor__()
            d_265_v163_ = nw65_
        d_266_v164_: D15
        d_266_v164_ = D15_DC30(d_84_v10_)
        d_266_v164_ = d_266_v164_
        d_267_v165_: D27
        d_267_v165_ = D27_DC57(_dafny.SeqWithoutIsStrInference([d_80_v6_]))
        index36_ = default__.safeIndex(179, (d_73_v1_).length(0))
        (d_73_v1_)[index36_] = len((d_267_v165_).cf87)
        d_268_v167_: _dafny.Seq
        def iife42_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in (_dafny.MultiSet(d_83_v9_)).Elements:
                d_269_v166_: int = compr_28_
                if (d_269_v166_) in (_dafny.MultiSet(d_83_v9_)):
                    coll28_[(d_269_v166_) - (d_79_v5_)] = d_77_v3_
            return _dafny.Map(coll28_)
        d_268_v167_ = _dafny.SeqWithoutIsStrInference([iife42_()
        ])
        (d_76_globalState_).f2 = len(d_268_v167_)
        index37_ = default__.safeIndex(627, (d_71_v0_).length(0))
        (d_71_v0_)[index37_] = (d_77_v3_) or (d_77_v3_)
        index38_ = default__.safeIndex(627, (d_71_v0_).length(0))
        (d_71_v0_)[index38_] = d_77_v3_
        hi1_ = (0) - ((d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))])
        for d_270_i14_ in range(default__.fm0(not((d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))]), d_76_globalState_), hi1_):
            d_271_v168_: _dafny.Seq
            d_271_v168_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
            d_272_v169_: _dafny.Array
            d_273_v170_: _dafny.Seq
            d_274_v171_: int
            out7_: _dafny.Array
            out8_: _dafny.Seq
            out9_: int
            out7_, out8_, out9_ = default__.m16(d_84_v10_, 361, d_83_v9_, (d_271_v168_) <= (d_271_v168_), d_76_globalState_)
            d_272_v169_ = out7_
            d_273_v170_ = out8_
            d_274_v171_ = out9_
            (d_76_globalState_).f2 = default__.safeModuloInt(d_270_i14_, d_79_v5_)
            d_275_v172_: _dafny.Map
            d_275_v172_ = _dafny.Map({(d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]: d_79_v5_})
            d_276_v173_: _dafny.Array
            nw66_ = _dafny.Array(D9.default()(), 14)
            d_276_v173_ = nw66_
            d_277_v174_: D18
            d_277_v174_ = D18_DC37(d_276_v173_)
            source2_ = ((D18_DC37(d_276_v173_) if d_77_v3_ else d_277_v174_) if (d_275_v172_) != (d_275_v172_) else d_277_v174_)
            if source2_.is_DC38:
                d_278___mcc_h3_ = source2_.cf57
                d_279___mcc_h4_ = source2_.cf58
                d_280___mcc_h5_ = source2_.cf59
                d_281_cf59_ = d_280___mcc_h5_
                d_282_cf58_ = d_279___mcc_h4_
                d_283_cf57_ = d_278___mcc_h3_
                d_284_v175_: _dafny.Array
                nw67_ = _dafny.Array(_dafny.Seq({}), 18)
                d_284_v175_ = nw67_
                d_285_v176_: _dafny.Set
                d_285_v176_ = _dafny.Set({-789})
                rhs39_ = d_284_v175_
                rhs40_ = ((d_285_v176_).isdisjoint(d_285_v176_) if (d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))] else ((d_80_v6_)[d_283_cf57_] if (d_283_cf57_) in (d_80_v6_) else True))
                d_284_v175_ = rhs39_
                d_77_v3_ = rhs40_
                d_286_v177_: _dafny.Map
                d_286_v177_ = _dafny.Map({(d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))]: d_73_v1_})
                d_286_v177_ = (d_286_v177_).set((d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))], d_73_v1_)
                d_287_v178_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_287_v178_ = nw68_
                index39_ = default__.safeIndex(35, (d_287_v178_).length(0))
                (d_287_v178_)[index39_] = d_73_v1_
                index40_ = default__.safeIndex(35, (d_287_v178_).length(0))
                rhs41_ = default__.fm0((not(default__.fm43(d_282_cf58_, d_76_globalState_))) and (d_77_v3_), d_76_globalState_)
                rhs42_ = d_73_v1_
                lhs26_ = d_287_v178_
                lhs27_ = default__.safeIndex(35, (d_287_v178_).length(0))
                d_79_v5_ = rhs41_
                lhs26_[lhs27_] = rhs42_
                (d_76_globalState_).f7 = default__.fm0((d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))], d_76_globalState_)
            elif True:
                d_288___mcc_h6_ = source2_.cf56
                d_289_cf56_ = d_288___mcc_h6_
                d_290_v179_: C6
                nw69_ = C6()
                nw69_.ctor__(d_274_v171_, d_77_v3_, (d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))], (len(_dafny.Map({d_77_v3_: 574}))) <= (992), (d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))])
                d_290_v179_ = nw69_
                d_91_v17_ = d_91_v17_
                d_291_v180_: C13
                nw70_ = C13()
                nw70_.ctor__()
                d_291_v180_ = nw70_
                d_292_v181_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.Map({}), 24)
                d_292_v181_ = nw71_
                d_293_v182_: _dafny.Map
                d_293_v182_ = _dafny.Map({d_91_v17_: d_77_v3_})
                index41_ = default__.safeIndex(296, (d_292_v181_).length(0))
                (d_292_v181_)[index41_] = d_293_v182_
                d_294_v183_: D16
                d_294_v183_ = D16_DC33((d_290_v179_).f20, d_77_v3_, (d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))], d_77_v3_)
                d_295_v184_: _dafny.Seq
                d_295_v184_ = _dafny.SeqWithoutIsStrInference([d_294_v183_])
                d_296_v185_: C9
                nw72_ = C9()
                nw72_.ctor__()
                d_296_v185_ = nw72_
                d_297_v186_: _dafny.Map
                d_297_v186_ = _dafny.Map({d_79_v5_: d_296_v185_})
                d_298_v187_: _dafny.Set
                d_298_v187_ = _dafny.Set({((d_297_v186_)[(d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]] if ((d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]) in (d_297_v186_) else d_296_v185_)})
                d_299_v188_: _dafny.Seq
                d_299_v188_ = _dafny.SeqWithoutIsStrInference([d_295_v184_, d_295_v184_, (d_295_v184_).set(default__.safeIndex(17, len(d_295_v184_)), d_294_v183_)])
                pat_let_tv4_ = d_274_v171_
                index42_ = default__.safeIndex(296, (d_292_v181_).length(0))
                def iife43_(_pat_let7_0):
                    def iife44_(d_300_dt__update__tmp_h2_):
                        def iife45_(_pat_let8_0):
                            def iife46_(d_301_dt__update_hcf49_h0_):
                                return D16_DC33((d_300_dt__update__tmp_h2_).cf47, (d_300_dt__update__tmp_h2_).cf48, d_301_dt__update_hcf49_h0_, (d_300_dt__update__tmp_h2_).cf50)
                            return iife46_(_pat_let8_0)
                        return iife45_(pat_let_tv4_)
                    return iife44_(_pat_let7_0)
                rhs43_ = default__.safeModuloInt(((d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]) + (d_274_v171_), -573)
                rhs44_ = d_291_v180_
                rhs45_ = d_293_v182_
                rhs46_ = ((_dafny.SeqWithoutIsStrInference([iife43_(d_294_v183_)])) + (d_295_v184_)) + ((d_299_v188_)[default__.safeIndex(len(d_271_v168_), len(d_299_v188_))])
                rhs47_ = _dafny.Set({d_296_v185_, d_296_v185_, d_296_v185_})
                lhs28_ = d_76_globalState_
                lhs29_ = d_292_v181_
                lhs30_ = default__.safeIndex(296, (d_292_v181_).length(0))
                lhs28_.f2 = rhs43_
                d_291_v180_ = rhs44_
                lhs29_[lhs30_] = rhs45_
                d_295_v184_ = rhs46_
                d_298_v187_ = rhs47_
                d_302_v189_: int
                d_303_v190_: bool
                out10_: int
                out11_: bool
                out10_, out11_ = (d_296_v185_).m5(d_76_globalState_)
                d_302_v189_ = out10_
                d_303_v190_ = out11_
            index43_ = default__.safeIndex(627, (d_71_v0_).length(0))
            (d_71_v0_)[index43_] = not ((d_86_v12_)[default__.safeIndex(d_270_i14_, len(d_86_v12_))]) or ((d_71_v0_)[default__.safeIndex(627, (d_71_v0_).length(0))])
        d_304_v191_: _dafny.Array
        def lambda8_(d_305_v15_):
            def lambda9_(d_306_i16_):
                return d_305_v15_

            return lambda9_

        init6_ = lambda8_(d_89_v15_)
        nw73_ = _dafny.Array(None, 13)
        for i0_6_ in range(nw73_.length(0)):
            nw73_[i0_6_] = init6_(i0_6_)
        d_304_v191_ = nw73_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_304_v191_).length(0)):
            d_307_i15_: int = guard_loop_1_
            if (True) and (((0) <= (d_307_i15_)) and ((d_307_i15_) < ((d_304_v191_).length(0)))):
                _ingredientsd_0.append((d_304_v191_, int(d_307_i15_), _dafny.Set({d_77_v3_, (_dafny.Set({d_79_v5_, (d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]})).issubset(_dafny.Set({(d_73_v1_)[default__.safeIndex(179, (d_73_v1_).length(0))]}))})))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        (d_76_globalState_).f2 = 115
        _dafny.print(_dafny.string_of((d_71_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v2_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_76_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_76_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_globalState_).f12).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_77_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_78_v4_)[0]) == (_dafny.SeqWithoutIsStrInference([-63, 2, -63, 63, 63, 63]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_79_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_80_v6_) == (_dafny.Map({63: True, 3: False, -4: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v7_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v8_) == (_dafny.MultiSet([1, 63]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v9_) == (_dafny.SeqWithoutIsStrInference([-63, 2, -63, 63, 63]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_v11_) == (_dafny.Map({63: _dafny.CodePoint('c')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_86_v12_) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v13_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_88_v14_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v15_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_90_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v17_) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v18_) == (_dafny.Map({3: _dafny.MultiSet([True, False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_93_v19_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, False, False]), _dafny.MultiSet([True, True]), _dafny.MultiSet([True, False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v152_).cf46) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v164_).cf44))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_267_v165_).cf87) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({63: True, 3: False, -4: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v167_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({-64: True, 1: True, 62: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[0]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[1]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[2]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[3]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[4]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[5]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[6]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[7]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[8]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[9]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[10]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[11]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v191_)[12]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
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
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

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
        return lambda: int(0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D2_DC2)

class D2_DC2(D2, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC4(False, int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D3_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D3_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)

class D3_DC4(D3, NamedTuple('DC4', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC3(D3, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC7(_dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D4_DC6)

class D4_DC7(D4, NamedTuple('DC7', [('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC6(D4, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D5_DC8)

class D5_DC8(D5, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC10()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D6_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D6_DC9)

class D6_DC10(D6, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC9(D6, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC12()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D7_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D7_DC11)

class D7_DC12(D7, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC11(D7, NamedTuple('DC11', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC11({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC11) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D8_DC13)

class D8_DC13(D8, NamedTuple('DC13', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC13({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC13) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC15(False, False, _dafny.Set({}), _dafny.Array(None, 0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D9_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D9_DC16)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D9_DC14)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D9_DC17)

class D9_DC15(D9, NamedTuple('DC15', [('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC15({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC15) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC16(D9, NamedTuple('DC16', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC16({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC16) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC14(D9, NamedTuple('DC14', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC14({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC14) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC17(D9, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D10_DC18)

class D10_DC18(D10, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC20()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D11_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D11_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D11_DC21)

class D11_DC20(D11, NamedTuple('DC20', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC20'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC20)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC19(D11, NamedTuple('DC19', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC19({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC19) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC21(D11, NamedTuple('DC21', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC21({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC21) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC23(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D12_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D12_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D12_DC22)

class D12_DC23(D12, NamedTuple('DC23', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC23({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC23) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC24(D12, NamedTuple('DC24', [('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC24({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC24) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC22(D12, NamedTuple('DC22', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC22({self.cf29.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC22) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC26()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D13_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D13_DC25)

class D13_DC26(D13, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC25(D13, NamedTuple('DC25', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC25({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC25) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC28(int(0), False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D14_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D14_DC27)

class D14_DC28(D14, NamedTuple('DC28', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC28({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC28) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC27(D14, NamedTuple('DC27', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC27({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC27) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC30(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D15_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D15_DC29)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D15_DC31)

class D15_DC30(D15, NamedTuple('DC30', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC30({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC30) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC29(D15, NamedTuple('DC29', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC29({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC29) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC31(D15, NamedTuple('DC31', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC31({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC31) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC33(False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D16_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D16_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D16_DC32)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC33(D16, NamedTuple('DC33', [('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC33({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC33) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC34(D16, NamedTuple('DC34', [('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC34({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC34) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC32(D16, NamedTuple('DC32', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC32({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC32) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D17_DC36)

class D17_DC36(D17, NamedTuple('DC36', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC36({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC36) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC38(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D18_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D18_DC37)

class D18_DC38(D18, NamedTuple('DC38', [('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC38({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC38) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC37(D18, NamedTuple('DC37', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC37({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC37) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC40(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D19_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D19_DC39)

class D19_DC40(D19, NamedTuple('DC40', [('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC40({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC40) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC39(D19, NamedTuple('DC39', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC39({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC39) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC42(int(0), False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D20_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D20_DC41)

class D20_DC42(D20, NamedTuple('DC42', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC42({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC42) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC41(D20, NamedTuple('DC41', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC41({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC41) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC44(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D21_DC44)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D21_DC43)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D21_DC45)

class D21_DC44(D21, NamedTuple('DC44', [('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC44({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC44) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC43(D21, NamedTuple('DC43', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC43({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC43) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC45(D21, NamedTuple('DC45', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC45({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC45) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D22_DC46)

class D22_DC46(D22, NamedTuple('DC46', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC46({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC46) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC48(False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D23_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D23_DC47)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D23_DC49)

class D23_DC48(D23, NamedTuple('DC48', [('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC48({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC48) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC47(D23, NamedTuple('DC47', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC47({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC47) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC49(D23, NamedTuple('DC49', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC49({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC49) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D24_DC50)

class D24_DC50(D24, NamedTuple('DC50', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC50({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC50) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC52(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D25_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D25_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D25_DC51)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D25_DC54)

class D25_DC52(D25, NamedTuple('DC52', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC52({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC52) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC53(D25, NamedTuple('DC53', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC53'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC53)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC51(D25, NamedTuple('DC51', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC51({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC51) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC54(D25, NamedTuple('DC54', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC54({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC54) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC56(int(0), None, _dafny.Seq({}), _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D26_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D26_DC55)

class D26_DC56(D26, NamedTuple('DC56', [('cf82', Any), ('cf83', Any), ('cf84', Any), ('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC56({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {self.cf86.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC56) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC55(D26, NamedTuple('DC55', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC55({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC55) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC58(_dafny.Array(None, 0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D27_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D27_DC57)

class D27_DC58(D27, NamedTuple('DC58', [('cf88', Any), ('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC58({_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC58) and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC57(D27, NamedTuple('DC57', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC57({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC57) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC60(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D28_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D28_DC61)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D28_DC59)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D28_DC62)

class D28_DC60(D28, NamedTuple('DC60', [('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC60({_dafny.string_of(self.cf92)}, {self.cf93.VerbatimString(True)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC60) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC61(D28, NamedTuple('DC61', [('cf95', Any), ('cf96', Any), ('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC61({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC61) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC59(D28, NamedTuple('DC59', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC59({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC59) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC62(D28, NamedTuple('DC62', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC62({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC62) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC64()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D29_DC64)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D29_DC63)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D29_DC65)

class D29_DC64(D29, NamedTuple('DC64', [])):
    def __dafnystr__(self) -> str:
        return f'D29.DC64'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC64)
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC63(D29, NamedTuple('DC63', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC63({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC63) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC65(D29, NamedTuple('DC65', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC65({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC65) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value

class T1:
    pass
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    def m2(self, p0, p1, globalState):
        pass

    def m3(self, p0, p1, p2, p3, globalState):
        pass


class T2:
    pass
    def m6(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self.f7: int = int(0)
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f3: int = int(0)
        self._f4: bool = False
        self._f5: int = int(0)
        self._f6: bool = False
        self._f8: int = int(0)
        self._f9: int = int(0)
        self._f10: bool = False
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self._f12: _dafny.MultiSet = _dafny.MultiSet({})
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13

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
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
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
    def f13(self):
        return self._f13

class C0:
    def  __init__(self):
        self.f26: _dafny.Array = _dafny.Array(None, 0)
        self.f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f26, f27):
        (self).f26 = f26
        (self).f27 = f27

    def fm36(self, p0, p1, p2, globalState):
        return self.f27

    def fm37(self, p0, globalState):
        return (D9_DC16(self.f27, not(False))).cf24

    def m15(self, p0, p1, p2, p3, globalState):
        d_308_v0_: _dafny.Array
        def lambda10_(d_309_i0_):
            return True

        init7_ = lambda10_
        nw74_ = _dafny.Array(None, 14)
        for i0_7_ in range(nw74_.length(0)):
            nw74_[i0_7_] = init7_(i0_7_)
        d_308_v0_ = nw74_
        index44_ = default__.safeIndex(253, (d_308_v0_).length(0))
        (d_308_v0_)[index44_] = (self.f27) <= (520)
        d_310_v1_: bool
        d_310_v1_ = True
        index45_ = default__.safeIndex(253, (d_308_v0_).length(0))
        (d_308_v0_)[index45_] = d_310_v1_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_308_v0_).length(0)):
            d_311_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_311_i1_)) and ((d_311_i1_) < ((d_308_v0_).length(0)))):
                (d_308_v0_)[(d_311_i1_)] = not((p2) < (self.f27))
        d_312_v2_: _dafny.MultiSet
        d_312_v2_ = _dafny.MultiSet([False])
        d_313_v3_: _dafny.Seq
        d_313_v3_ = _dafny.SeqWithoutIsStrInference([(d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))], d_310_v1_])
        d_314_v4_: int
        d_314_v4_ = len(d_313_v3_)
        d_315_v5_: _dafny.Set
        d_315_v5_ = _dafny.Set({((d_312_v2_)[(self).fm37((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))], globalState)] if ((self).fm37((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))], globalState)) in (d_312_v2_) else (d_314_v4_)), self.f27, p2})
        d_316_v6_: _dafny.Array
        def lambda11_(d_317_p2_):
            def lambda12_(d_318_i2_):
                return default__.safeDivisionInt(d_318_i2_, d_317_p2_)

            return lambda12_

        init8_ = lambda11_(p2)
        nw75_ = _dafny.Array(None, 14)
        for i0_8_ in range(nw75_.length(0)):
            nw75_[i0_8_] = init8_(i0_8_)
        d_316_v6_ = nw75_
        d_319_v7_: D9
        d_319_v7_ = D9_DC15(False, (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))], d_315_v5_, d_316_v6_, d_308_v0_)
        d_319_v7_ = d_319_v7_
        d_320_v8_: D7
        d_320_v8_ = D7_DC12()
        d_321_i3_: int
        d_321_i3_ = 0
        with _dafny.label("1"):
            pat_let_tv5_ = d_315_v5_
            pat_let_tv6_ = d_310_v1_
            def lambda13_(source3_):
                if source3_.is_DC12:
                    return (pat_let_tv5_) != (_dafny.Set({self.f27}))
                elif True:
                    d_324___mcc_h0_ = source3_.cf15
                    d_325_cf15_ = d_324___mcc_h0_
                    return pat_let_tv6_

            while lambda13_(d_320_v8_):
                with _dafny.c_label("1"):
                    if (d_321_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_321_i3_ = (d_321_i3_) + (1)
                    index46_ = default__.safeIndex(253, (d_308_v0_).length(0))
                    rhs48_ = True
                    rhs49_ = (0) - (default__.fm0((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))], globalState))
                    lhs31_ = d_308_v0_
                    lhs32_ = default__.safeIndex(253, (d_308_v0_).length(0))
                    lhs33_ = globalState
                    lhs31_[lhs32_] = rhs48_
                    lhs33_.f7 = rhs49_
                    d_322_v9_: D9
                    d_322_v9_ = D9_DC16(len(p3), (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))])
                    index47_ = default__.safeIndex(253, (d_308_v0_).length(0))
                    (d_308_v0_)[index47_] = not(not(not((d_322_v9_).cf24)))
                    d_310_v1_ = (self.f27) >= (p2)
                    d_323_v10_: _dafny.Array
                    nw76_ = _dafny.Array(None, 17)
                    nw76_[int(0)] = d_308_v0_
                    nw76_[int(1)] = d_308_v0_
                    nw76_[int(2)] = d_308_v0_
                    nw76_[int(3)] = d_308_v0_
                    nw76_[int(4)] = d_308_v0_
                    nw76_[int(5)] = d_308_v0_
                    nw76_[int(6)] = d_308_v0_
                    nw76_[int(7)] = d_308_v0_
                    nw76_[int(8)] = d_308_v0_
                    nw76_[int(9)] = d_308_v0_
                    nw76_[int(10)] = d_308_v0_
                    nw76_[int(11)] = d_308_v0_
                    nw76_[int(12)] = (d_308_v0_ if (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))] else d_308_v0_)
                    nw76_[int(13)] = d_308_v0_
                    nw76_[int(14)] = d_308_v0_
                    nw76_[int(15)] = d_308_v0_
                    nw76_[int(16)] = d_308_v0_
                    d_323_v10_ = nw76_
                    d_323_v10_ = (d_323_v10_ if d_310_v1_ else (d_323_v10_ if d_310_v1_ else d_323_v10_))
                    pass
            pass
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_308_v0_).length(0)):
            d_326_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_326_i4_)) and ((d_326_i4_) < ((d_308_v0_).length(0)))):
                (d_308_v0_)[(d_326_i4_)] = (D9_DC16(self.f27, d_310_v1_)).cf24
        d_327_v11_: _dafny.Map
        d_327_v11_ = _dafny.Map({p2: (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))]})
        d_328_i5_: int
        d_328_i5_ = 0
        with _dafny.label("2"):
            while (p2) not in (d_327_v11_):
                with _dafny.c_label("2"):
                    if (d_328_i5_) >= (100):
                        raise _dafny.Break("2")
                    d_328_i5_ = (d_328_i5_) + (1)
                    if not((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))]):
                        d_315_v5_ = d_315_v5_
                        d_329_v12_: _dafny.Map
                        d_329_v12_ = _dafny.Map({self.f27: d_308_v0_})
                        d_330_v13_: _dafny.Array
                        nw77_ = _dafny.Array(None, 4)
                        nw77_[int(0)] = d_329_v12_
                        nw77_[int(1)] = d_329_v12_
                        nw77_[int(2)] = d_329_v12_
                        nw77_[int(3)] = d_329_v12_
                        d_330_v13_ = nw77_
                        index48_ = default__.safeIndex(483, (d_330_v13_).length(0))
                        (d_330_v13_)[index48_] = (d_329_v12_) | (d_329_v12_)
                        d_331_v14_: _dafny.Seq
                        d_331_v14_ = _dafny.SeqWithoutIsStrInference([self.f27, len(p3)])
                        index49_ = default__.safeIndex(483, (d_330_v13_).length(0))
                        (d_330_v13_)[index49_] = _dafny.Map({len(d_331_v14_): d_308_v0_})
                        d_310_v1_ = False
                        d_332_v15_: _dafny.Map
                        d_332_v15_ = _dafny.Map({p2: d_331_v14_})
                        d_332_v15_ = (d_332_v15_).set((d_312_v2_).cardinality, (((p1)[d_310_v1_] if (d_310_v1_) in (p1) else _dafny.SeqWithoutIsStrInference([self.f27]))) + (d_331_v14_))
                        (globalState).f7 = len(d_315_v5_)
                    elif True:
                        d_310_v1_ = ((d_312_v2_).set(not(not((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))])), default__.abs(p2))).issubset(d_312_v2_)
                        d_333_v16_: _dafny.Array
                        def lambda14_(d_334_v4_):
                            def lambda15_(d_335_i6_):
                                return d_334_v4_

                            return lambda15_

                        init9_ = lambda14_(d_314_v4_)
                        nw78_ = _dafny.Array(None, 28)
                        for i0_9_ in range(nw78_.length(0)):
                            nw78_[i0_9_] = init9_(i0_9_)
                        d_333_v16_ = nw78_
                        index50_ = default__.safeIndex(559, (d_333_v16_).length(0))
                        (d_333_v16_)[index50_] = d_314_v4_
                        index51_ = default__.safeIndex(559, (d_333_v16_).length(0))
                        (d_333_v16_)[index51_] = (p2) - (self.f27)
                        index52_ = default__.safeIndex(253, (d_308_v0_).length(0))
                        (d_308_v0_)[index52_] = d_310_v1_
                        d_336_v17_: D3
                        d_336_v17_ = D3_DC4(d_310_v1_, self.f27, self.f27, (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))], (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))])
                        (self).f27 = (d_336_v17_).cf5
                        d_327_v11_ = (d_327_v11_).set(366, d_310_v1_)
                    d_337_v18_: _dafny.Seq
                    d_337_v18_ = _dafny.SeqWithoutIsStrInference([self.f27])
                    d_338_v19_: _dafny.Seq
                    d_338_v19_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(len(d_313_v3_), p2), self.f27, (d_337_v18_)[default__.safeIndex(self.f27, len(d_337_v18_))], 821, p2])
                    d_339_v20_: _dafny.Map
                    d_339_v20_ = _dafny.Map({d_310_v1_: d_310_v1_})
                    rhs50_ = p2
                    rhs51_ = ((default__.fm38(False, p2, p2, globalState)).set(default__.safeIndex(p2, len(default__.fm38(False, p2, p2, globalState))), len(_dafny.Map({self.f27: d_310_v1_})))) + (d_338_v19_)
                    rhs52_ = _dafny.Map({(0) - (p2): (((d_339_v20_)[not((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))])] if (not((d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))])) in (d_339_v20_) else d_310_v1_) if (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))] else (d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))])})
                    lhs34_ = globalState
                    lhs34_.f7 = rhs50_
                    d_337_v18_ = rhs51_
                    d_327_v11_ = rhs52_
                    d_340_v21_: _dafny.Map
                    d_340_v21_ = _dafny.Map({_dafny.Map({(self).fm37(d_310_v1_, globalState): p2}): p2})
                    d_341_v22_: _dafny.Map
                    d_341_v22_ = _dafny.Map({(d_308_v0_)[default__.safeIndex(253, (d_308_v0_).length(0))]: (d_337_v18_)[default__.safeIndex(p2, len(d_337_v18_))]})
                    d_340_v21_ = (d_340_v21_).set(d_341_v22_, self.f27)
                    d_342_v23_: _dafny.Array
                    def lambda16_(d_343_i7_):
                        return _dafny.CodePoint('y')

                    init10_ = lambda16_
                    nw79_ = _dafny.Array(None, 12)
                    for i0_10_ in range(nw79_.length(0)):
                        nw79_[i0_10_] = init10_(i0_10_)
                    d_342_v23_ = nw79_
                    d_344_v24_: str
                    d_344_v24_ = _dafny.CodePoint('e')
                    index53_ = default__.safeIndex(81, (d_342_v23_).length(0))
                    (d_342_v23_)[index53_] = d_344_v24_
                    index54_ = default__.safeIndex(81, (d_342_v23_).length(0))
                    (d_342_v23_)[index54_] = d_344_v24_
                    pass
            pass


class C1(T0):
    def  __init__(self):
        self._f14: int = int(0)
        self.f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    def ctor__(self, f25, f14):
        (self).f25 = f25
        (self).f14 = f14

    def m13(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: str = _dafny.CodePoint('D')
        r3: _dafny.Map = _dafny.Map({})
        d_345_v0_: _dafny.Array
        nw80_ = _dafny.Array(False, 6)
        d_345_v0_ = nw80_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_345_v0_).length(0)):
            d_346_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_346_i0_)) and ((d_346_i0_) < ((d_345_v0_).length(0)))):
                (d_345_v0_)[(d_346_i0_)] = (p0) or (p0)
        d_347_v1_: _dafny.Array
        def lambda17_(d_348_p1_):
            def lambda18_(d_349_i1_):
                return (_dafny.MultiSet([len(_dafny.Map({self.f14: _dafny.SeqWithoutIsStrInference([d_348_p1_])}))])) - (_dafny.MultiSet([self.f14, self.f14]))

            return lambda18_

        init11_ = lambda17_(p1)
        nw81_ = _dafny.Array(None, 17)
        for i0_11_ in range(nw81_.length(0)):
            nw81_[i0_11_] = init11_(i0_11_)
        d_347_v1_ = nw81_
        d_350_v2_: _dafny.Seq
        d_350_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njtfyb"))
        d_351_v3_: _dafny.Map
        d_351_v3_ = _dafny.Map({p0: len(d_350_v2_)})
        d_352_v4_: _dafny.MultiSet
        d_352_v4_ = _dafny.MultiSet([len(d_351_v3_), self.f14])
        index55_ = default__.safeIndex(456, (d_347_v1_).length(0))
        (d_347_v1_)[index55_] = (d_352_v4_).intersection(d_352_v4_)
        index56_ = default__.safeIndex(456, (d_347_v1_).length(0))
        (d_347_v1_)[index56_] = d_352_v4_
        d_353_v5_: _dafny.Array
        def lambda19_(d_354_p0_, d_355_p1_):
            def lambda20_(d_356_i2_):
                return _dafny.Set({d_354_p0_, d_355_p1_})

            return lambda20_

        init12_ = lambda19_(p0, p1)
        nw82_ = _dafny.Array(None, 27)
        for i0_12_ in range(nw82_.length(0)):
            nw82_[i0_12_] = init12_(i0_12_)
        d_353_v5_ = nw82_
        d_357_v6_: C0
        nw83_ = C0()
        nw83_.ctor__(d_353_v5_, self.f14)
        d_357_v6_ = nw83_
        d_358_v7_: _dafny.Map
        d_358_v7_ = _dafny.Map({d_350_v2_: True})
        d_359_v8_: D3
        d_359_v8_ = D3_DC4(p0, 517, d_357_v6_.f27, True, ((d_358_v7_)[d_350_v2_] if (d_350_v2_) in (d_358_v7_) else p1))
        hi2_ = d_357_v6_.f27
        for d_360_i3_ in range((d_359_v8_).cf5, hi2_):
            (self).f25 = p0
            d_361_v9_: _dafny.Map
            d_361_v9_ = _dafny.Map({(d_357_v6_).fm37(not(p1), globalState): default__.fm39(globalState)})
            d_362_v10_: str
            d_362_v10_ = _dafny.CodePoint('y')
            d_361_v9_ = (d_361_v9_).set(self.f25, d_362_v10_)
            d_363_v11_: _dafny.Map
            d_363_v11_ = _dafny.Map({self.f25: self.f25})
            d_364_v12_: _dafny.Map
            d_364_v12_ = _dafny.Map({p1: d_363_v11_})
            d_365_v13_: _dafny.Map
            d_365_v13_ = _dafny.Map({p0: d_363_v11_})
            r1 = (True) not in ((((d_364_v12_)).set(p1, default__.fm40(p1, globalState))) | (d_365_v13_))
            (self).f25 = p1
        d_366_v14_: str
        d_366_v14_ = _dafny.CodePoint('o')
        d_367_v15_: _dafny.Map
        d_367_v15_ = _dafny.Map({d_366_v14_: self.f14})
        d_368_v16_: D9
        d_368_v16_ = D9_DC16(len((d_351_v3_).set(self.f25, 968)), True)
        hi3_ = self.f14
        for d_369_i4_ in range(((d_367_v15_)[d_366_v14_] if (d_366_v14_) in (d_367_v15_) else default__.fm0((d_368_v16_).cf24, globalState)), hi3_):
            (self).f25 = self.f25
            d_370_v17_: _dafny.Seq
            d_370_v17_ = _dafny.SeqWithoutIsStrInference([p0, p1])
            d_371_v18_: _dafny.Map
            d_371_v18_ = _dafny.Map({len(d_370_v17_): -186})
            d_372_v19_: _dafny.Seq
            d_372_v19_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(d_357_v6_.f27, -108), len((d_371_v18_) | ((d_371_v18_).set(self.f14, (0) - (d_357_v6_.f27)))), (0) - (d_357_v6_.f27), d_357_v6_.f27])
            rhs53_ = d_357_v6_.f27
            rhs54_ = _dafny.SeqWithoutIsStrInference([d_357_v6_.f27])
            rhs55_ = default__.fm38((d_368_v16_).cf24, (0) - (d_357_v6_.f27), default__.safeDivisionInt(d_357_v6_.f27, d_369_i4_), globalState)
            lhs35_ = globalState
            lhs35_.f2 = rhs53_
            d_372_v19_ = rhs54_
            d_372_v19_ = rhs55_
            if not(p1):
                (self).f25 = not(p1)
                d_373_v20_: _dafny.Map
                d_373_v20_ = _dafny.Map({d_366_v14_: self.f25})
                d_374_v21_: _dafny.Seq
                d_374_v21_ = _dafny.SeqWithoutIsStrInference([d_370_v17_])
                d_375_v22_: _dafny.MultiSet
                d_375_v22_ = _dafny.MultiSet([p0, p1])
                d_376_v23_: _dafny.Array
                nw84_ = _dafny.Array(None, 21)
                nw84_[int(0)] = default__.safeDivisionInt(self.f14, d_357_v6_.f27)
                nw84_[int(1)] = (d_357_v6_.f27) + (default__.fm0(self.f25, globalState))
                nw84_[int(2)] = (default__.fm0(((d_373_v20_)[d_366_v14_] if (d_366_v14_) in (d_373_v20_) else self.f25), globalState)) + (-340)
                nw84_[int(3)] = 902
                nw84_[int(4)] = d_357_v6_.f27
                nw84_[int(5)] = (0) - (((d_351_v3_)[p1] if (p1) in (d_351_v3_) else d_357_v6_.f27))
                nw84_[int(6)] = len(d_351_v3_)
                nw84_[int(7)] = len(d_350_v2_)
                nw84_[int(8)] = (len((d_374_v21_)[default__.safeIndex(d_357_v6_.f27, len(d_374_v21_))])) - (d_357_v6_.f27)
                nw84_[int(9)] = d_357_v6_.f27
                nw84_[int(10)] = (d_359_v8_).cf5
                nw84_[int(11)] = d_357_v6_.f27
                nw84_[int(12)] = ((d_371_v18_)[default__.fm0(p1, globalState)] if (default__.fm0(p1, globalState)) in (d_371_v18_) else (d_375_v22_).cardinality)
                nw84_[int(13)] = d_357_v6_.f27
                nw84_[int(14)] = len(d_372_v19_)
                nw84_[int(15)] = d_357_v6_.f27
                nw84_[int(16)] = d_357_v6_.f27
                nw84_[int(17)] = d_357_v6_.f27
                nw84_[int(18)] = self.f14
                nw84_[int(19)] = (d_357_v6_.f27) - (-472)
                nw84_[int(20)] = d_357_v6_.f27
                d_376_v23_ = nw84_
                d_376_v23_ = d_376_v23_
                d_377_v24_: _dafny.Map
                d_377_v24_ = _dafny.Map({p1: _dafny.Map({self.f25: self.f25})})
                d_378_v25_: _dafny.Map
                d_378_v25_ = d_377_v24_
                d_379_v26_: _dafny.Map
                d_379_v26_ = _dafny.Map({self.f25: p0})
                d_380_v27_: _dafny.Array
                nw85_ = _dafny.Array(None, 28)
                nw85_[int(0)] = d_378_v25_
                nw85_[int(1)] = d_377_v24_
                nw85_[int(2)] = default__.fm41(p1, p1, globalState)
                nw85_[int(3)] = (d_378_v25_ if p1 else d_378_v25_)
                nw85_[int(4)] = d_378_v25_
                nw85_[int(5)] = _dafny.Map({self.f25: d_379_v26_})
                nw85_[int(6)] = d_378_v25_
                nw85_[int(7)] = d_378_v25_
                nw85_[int(8)] = d_378_v25_
                nw85_[int(9)] = _dafny.Map({p0: d_379_v26_})
                nw85_[int(10)] = d_378_v25_
                nw85_[int(11)] = d_378_v25_
                nw85_[int(12)] = d_378_v25_
                nw85_[int(13)] = d_378_v25_
                nw85_[int(14)] = d_378_v25_
                nw85_[int(15)] = _dafny.Map({p1: d_379_v26_})
                nw85_[int(16)] = d_378_v25_
                nw85_[int(17)] = default__.fm41(p0, p0, globalState)
                nw85_[int(18)] = d_378_v25_
                nw85_[int(19)] = d_377_v24_
                nw85_[int(20)] = d_378_v25_
                nw85_[int(21)] = d_378_v25_
                nw85_[int(22)] = d_378_v25_
                nw85_[int(23)] = d_378_v25_
                nw85_[int(24)] = d_378_v25_
                nw85_[int(25)] = default__.fm41(p0, self.f25, globalState)
                nw85_[int(26)] = d_378_v25_
                nw85_[int(27)] = (d_378_v25_ if p0 else d_378_v25_)
                d_380_v27_ = nw85_
                index57_ = default__.safeIndex(484, (d_380_v27_).length(0))
                (d_380_v27_)[index57_] = d_377_v24_
                index58_ = default__.safeIndex(484, (d_380_v27_).length(0))
                (d_380_v27_)[index58_] = d_378_v25_
                (globalState).f7 = self.f14
                d_381_v28_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_381_v28_ = nw86_
                index59_ = default__.safeIndex(273, (d_381_v28_).length(0))
                (d_381_v28_)[index59_] = d_376_v23_
                index60_ = default__.safeIndex(672, (d_381_v28_).length(0))
                (d_381_v28_)[index60_] = (d_376_v23_ if self.f25 else d_376_v23_)
                d_382_v29_: D6
                d_382_v29_ = D6_DC9(d_366_v14_)
                d_383_v31_: _dafny.Map
                def iife47_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in (_dafny.SeqWithoutIsStrInference([d_357_v6_.f27])).Elements:
                        d_384_v30_: int = compr_29_
                        if (d_384_v30_) in (_dafny.SeqWithoutIsStrInference([d_357_v6_.f27])):
                            coll29_ = coll29_.union(_dafny.Set([(d_384_v30_) - (668)]))
                    return _dafny.Set(coll29_)
                d_383_v31_ = _dafny.Map({iife47_()
                : d_350_v2_})
                d_385_v32_: _dafny.Set
                d_385_v32_ = _dafny.Set({d_357_v6_.f27, d_357_v6_.f27})
                d_386_v33_: _dafny.Array
                nw87_ = _dafny.Array(None, 24)
                nw87_[int(0)] = (d_350_v2_).set(default__.safeIndex(d_357_v6_.f27, len(d_350_v2_)), (d_382_v29_).cf14)
                nw87_[int(1)] = (d_350_v2_) + (d_350_v2_)
                nw87_[int(2)] = d_350_v2_
                nw87_[int(3)] = (d_350_v2_) + (d_350_v2_)
                nw87_[int(4)] = _dafny.SeqWithoutIsStrInference([d_366_v14_ for d_387_i5_ in range(default__.abs(595))])
                nw87_[int(5)] = d_350_v2_
                nw87_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slao"))
                nw87_[int(7)] = (default__.fm42(p1, (_dafny.MultiSet([d_357_v6_.f27, d_357_v6_.f27, d_357_v6_.f27])).cardinality, d_357_v6_.f27, p1, globalState)) + (d_350_v2_)
                nw87_[int(8)] = (d_350_v2_).set(default__.safeIndex(self.f14, len(d_350_v2_)), d_366_v14_)
                nw87_[int(9)] = d_350_v2_
                nw87_[int(10)] = d_350_v2_
                nw87_[int(11)] = (d_350_v2_) + (default__.fm42(p0, d_357_v6_.f27, d_357_v6_.f27, p1, globalState))
                nw87_[int(12)] = (d_350_v2_) + (d_350_v2_)
                nw87_[int(13)] = d_350_v2_
                nw87_[int(14)] = d_350_v2_
                nw87_[int(15)] = (d_350_v2_) + (d_350_v2_)
                nw87_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sub"))
                nw87_[int(17)] = (_dafny.SeqWithoutIsStrInference([d_366_v14_ for d_388_i6_ in range(default__.abs(21))])).set(default__.safeIndex(-560, len(_dafny.SeqWithoutIsStrInference([d_366_v14_ for d_389_i6_ in range(default__.abs(21))]))), d_366_v14_)
                nw87_[int(18)] = (d_350_v2_) + (d_350_v2_)
                nw87_[int(19)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrtitlk"))).set(default__.safeIndex(d_369_i4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrtitlk")))), d_366_v14_)
                nw87_[int(20)] = (_dafny.SeqWithoutIsStrInference([d_366_v14_ for d_390_i7_ in range(default__.abs(605))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytmlhni")))
                nw87_[int(21)] = d_350_v2_
                nw87_[int(22)] = ((d_383_v31_)[d_385_v32_] if (d_385_v32_) in (d_383_v31_) else default__.fm42(p0, d_357_v6_.f27, 453, False, globalState))
                nw87_[int(23)] = d_350_v2_
                d_386_v33_ = nw87_
                index61_ = default__.safeIndex(918, (d_386_v33_).length(0))
                (d_386_v33_)[index61_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwcpbhs"))
                index62_ = default__.safeIndex(704, (d_376_v23_).length(0))
                (d_376_v23_)[index62_] = len(d_350_v2_)
                d_391_v34_: _dafny.Seq
                d_391_v34_ = _dafny.SeqWithoutIsStrInference([d_350_v2_])
                d_392_v35_: _dafny.Map
                d_392_v35_ = _dafny.Map({len((d_370_v17_)): d_345_v0_})
                index63_ = default__.safeIndex(273, (d_381_v28_).length(0))
                index64_ = default__.safeIndex(672, (d_381_v28_).length(0))
                index65_ = default__.safeIndex(918, (d_386_v33_).length(0))
                index66_ = default__.safeIndex(704, (d_376_v23_).length(0))
                rhs56_ = d_376_v23_
                rhs57_ = d_376_v23_
                rhs58_ = (((d_391_v34_)[default__.safeIndex(d_357_v6_.f27, len(d_391_v34_))] if p0 else _dafny.SeqWithoutIsStrInference([d_366_v14_ for d_393_i8_ in range(default__.abs(-320))]))) + (d_350_v2_)
                rhs59_ = (0) - (-799)
                rhs60_ = len(d_392_v35_)
                lhs36_ = d_381_v28_
                lhs37_ = default__.safeIndex(273, (d_381_v28_).length(0))
                lhs38_ = d_381_v28_
                lhs39_ = default__.safeIndex(672, (d_381_v28_).length(0))
                lhs40_ = d_386_v33_
                lhs41_ = default__.safeIndex(918, (d_386_v33_).length(0))
                lhs42_ = globalState
                lhs43_ = d_376_v23_
                lhs44_ = default__.safeIndex(704, (d_376_v23_).length(0))
                lhs36_[lhs37_] = rhs56_
                lhs38_[lhs39_] = rhs57_
                lhs40_[lhs41_] = rhs58_
                lhs42_.f7 = rhs59_
                lhs43_[lhs44_] = rhs60_
            elif True:
                (globalState).f7 = d_357_v6_.f27
                d_394_v36_: _dafny.Set
                d_394_v36_ = _dafny.Set({d_369_i4_})
                d_395_v37_: _dafny.Array
                def lambda21_(d_396_i9_):
                    return (d_396_i9_) * (self.f14)

                init13_ = lambda21_
                nw88_ = _dafny.Array(None, 24)
                for i0_13_ in range(nw88_.length(0)):
                    nw88_[i0_13_] = init13_(i0_13_)
                d_395_v37_ = nw88_
                d_397_v38_: _dafny.Map
                d_397_v38_ = _dafny.Map({d_369_i4_: d_345_v0_})
                d_398_v39_: D9
                d_398_v39_ = D9_DC15(p0, p0, d_394_v36_, d_395_v37_, ((d_397_v38_)[d_357_v6_.f27] if (d_357_v6_.f27) in (d_397_v38_) else d_345_v0_))
                r1 = (self.f25 if False else (d_398_v39_).cf19)
                (self).f25 = (D3_DC3(not(p1))).cf3
                d_399_v40_: _dafny.Array
                def lambda22_(d_400_v3_, d_401_v6_):
                    def lambda23_(d_402_i10_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_400_v3_) for d_403_i11_ in range(default__.abs(118))]), _dafny.SeqWithoutIsStrInference([d_401_v6_.f27])])

                    return lambda23_

                init14_ = lambda22_(d_351_v3_, d_357_v6_)
                nw89_ = _dafny.Array(None, 8)
                for i0_14_ in range(nw89_.length(0)):
                    nw89_[i0_14_] = init14_(i0_14_)
                d_399_v40_ = nw89_
                d_404_v41_: _dafny.Seq
                d_404_v41_ = _dafny.SeqWithoutIsStrInference([d_372_v19_])
                index67_ = default__.safeIndex(100, (d_399_v40_).length(0))
                (d_399_v40_)[index67_] = d_404_v41_
                d_405_v42_: _dafny.Array
                def lambda24_(d_406_p0_):
                    def lambda25_(d_407_i12_):
                        return _dafny.Map({d_406_p0_: self.f25})

                    return lambda25_

                init15_ = lambda24_(p0)
                nw90_ = _dafny.Array(None, 6)
                for i0_15_ in range(nw90_.length(0)):
                    nw90_[i0_15_] = init15_(i0_15_)
                d_405_v42_ = nw90_
                index68_ = default__.safeIndex(100, (d_399_v40_).length(0))
                rhs61_ = d_404_v41_
                rhs62_ = d_405_v42_
                rhs63_ = d_395_v37_
                rhs64_ = d_357_v6_.f27
                rhs65_ = self.f25
                lhs45_ = d_399_v40_
                lhs46_ = default__.safeIndex(100, (d_399_v40_).length(0))
                lhs47_ = globalState
                lhs45_[lhs46_] = rhs61_
                d_405_v42_ = rhs62_
                d_395_v37_ = rhs63_
                lhs47_.f7 = rhs64_
                r1 = rhs65_
                r2 = d_366_v14_
            d_408_v43_: _dafny.Map
            d_408_v43_ = _dafny.Map({d_372_v19_: -50})
            d_409_v44_: _dafny.Set
            d_409_v44_ = _dafny.Set({p1, p1, p1, p0})
            (self).f25 = (d_369_i4_) != (((d_408_v43_)[_dafny.SeqWithoutIsStrInference([349 for d_410_i13_ in range(default__.abs(274))])] if (_dafny.SeqWithoutIsStrInference([349 for d_411_i13_ in range(default__.abs(274))])) in (d_408_v43_) else len(d_409_v44_)))
        d_412_v45_: _dafny.Set
        d_412_v45_ = _dafny.Set({self.f14, self.f14})
        d_413_v46_: _dafny.MultiSet
        d_413_v46_ = _dafny.MultiSet([p0, p1, p1, (self.f14) < (len(d_412_v45_)), p0])
        d_414_v47_: _dafny.MultiSet
        d_414_v47_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_366_v14_ for d_415_i14_ in range(default__.abs(802))])])
        rhs66_ = (self.f14) * ((0) - (((d_413_v46_)[p1] if (p1) in (d_413_v46_) else len(_dafny.Map({self.f25: p1})))))
        rhs67_ = (d_366_v14_ if self.f25 else d_366_v14_)
        rhs68_ = not (p1) or ((d_414_v47_).issubset(d_414_v47_))
        rhs69_ = d_357_v6_.f27
        rhs70_ = d_413_v46_
        lhs48_ = globalState
        r0 = rhs66_
        r2 = rhs67_
        r1 = rhs68_
        lhs48_.f2 = rhs69_
        d_413_v46_ = rhs70_
        r0 = d_357_v6_.f27
        r1 = not(not(not (self.f25) or (not(p0))))
        r2 = d_366_v14_
        r3 = d_351_v3_
        return r0, r1, r2, r3

    def m14(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        (globalState).f7 = (0) - (self.f14)
        d_416_v0_: _dafny.Seq
        d_416_v0_ = _dafny.SeqWithoutIsStrInference([self.f25])
        d_417_v1_: _dafny.Set
        d_417_v1_ = _dafny.Set({self.f25, self.f25})
        d_418_v2_: _dafny.MultiSet
        d_418_v2_ = _dafny.MultiSet([self.f14])
        d_419_v3_: _dafny.Set
        d_419_v3_ = _dafny.Set({self.f14})
        d_420_v4_: _dafny.Map
        d_420_v4_ = _dafny.Map({self.f14: self.f14})
        d_421_v5_: _dafny.Array
        nw91_ = _dafny.Array(None, 6)
        nw91_[int(0)] = -4
        nw91_[int(1)] = len(d_420_v4_)
        nw91_[int(2)] = (0) - (self.f14)
        nw91_[int(3)] = 225
        nw91_[int(4)] = self.f14
        nw91_[int(5)] = self.f14
        d_421_v5_ = nw91_
        d_422_v6_: _dafny.Array
        def lambda26_(d_423_i0_):
            return self.f25

        init16_ = lambda26_
        nw92_ = _dafny.Array(None, 3)
        for i0_16_ in range(nw92_.length(0)):
            nw92_[i0_16_] = init16_(i0_16_)
        d_422_v6_ = nw92_
        d_424_v7_: D9
        d_424_v7_ = D9_DC15(True, self.f25, d_419_v3_, d_421_v5_, d_422_v6_)
        d_425_v8_: _dafny.Seq
        d_425_v8_ = _dafny.SeqWithoutIsStrInference([d_424_v7_])
        d_426_v9_: _dafny.Map
        d_426_v9_ = _dafny.Map({d_425_v8_: self.f14})
        d_427_v10_: _dafny.Seq
        d_427_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgnjpywx"))
        d_428_v11_: _dafny.Map
        d_428_v11_ = _dafny.Map({self.f25: self.f25})
        d_429_v12_: _dafny.Seq
        d_429_v12_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14, self.f14, self.f14])
        d_430_v13_: _dafny.Array
        nw93_ = _dafny.Array(None, 23)
        nw93_[int(0)] = len(d_416_v0_)
        nw93_[int(1)] = self.f14
        nw93_[int(2)] = (len(d_417_v1_)) * (105)
        nw93_[int(3)] = (0) - (len(d_417_v1_))
        nw93_[int(4)] = (0) - ((d_418_v2_).cardinality)
        nw93_[int(5)] = self.f14
        nw93_[int(6)] = (0) - (self.f14)
        nw93_[int(7)] = self.f14
        nw93_[int(8)] = self.f14
        nw93_[int(9)] = ((d_426_v9_)[(d_425_v8_).set(default__.safeIndex(self.f14, len(d_425_v8_)), d_424_v7_)] if ((d_425_v8_).set(default__.safeIndex(self.f14, len(d_425_v8_)), d_424_v7_)) in (d_426_v9_) else len(d_427_v10_))
        nw93_[int(10)] = (self.f14) + (907)
        nw93_[int(11)] = self.f14
        nw93_[int(12)] = self.f14
        nw93_[int(13)] = len(d_428_v11_)
        nw93_[int(14)] = self.f14
        nw93_[int(15)] = (self.f14) - (len(d_427_v10_))
        nw93_[int(16)] = self.f14
        nw93_[int(17)] = (0) - ((0) - (len(d_427_v10_)))
        nw93_[int(18)] = self.f14
        nw93_[int(19)] = len((d_416_v0_).set(default__.safeIndex(202, len(d_416_v0_)), self.f25))
        nw93_[int(20)] = default__.safeModuloInt(len(d_429_v12_), self.f14)
        nw93_[int(21)] = self.f14
        nw93_[int(22)] = 113
        d_430_v13_ = nw93_
        index69_ = default__.safeIndex(779, (d_421_v5_).length(0))
        (d_421_v5_)[index69_] = default__.safeModuloInt(self.f14, self.f14)
        index70_ = default__.safeIndex(779, (d_421_v5_).length(0))
        (d_421_v5_)[index70_] = len(d_416_v0_)
        d_424_v7_ = d_424_v7_
        hi4_ = self.f14
        for d_431_i1_ in range((0) - (len(d_416_v0_)), hi4_):
            d_432_v14_: _dafny.Array
            nw94_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_432_v14_ = nw94_
            index71_ = default__.safeIndex(8, (d_432_v14_).length(0))
            (d_432_v14_)[index71_] = d_421_v5_
            index72_ = default__.safeIndex(8, (d_432_v14_).length(0))
            (d_432_v14_)[index72_] = d_430_v13_
            def lambda27_(d_433_i2_):
                return (d_433_i2_) * (444)

            init17_ = lambda27_
            nw95_ = _dafny.Array(None, 13)
            for i0_17_ in range(nw95_.length(0)):
                nw95_[i0_17_] = init17_(i0_17_)
            d_421_v5_ = nw95_
            index73_ = default__.safeIndex(779, (d_421_v5_).length(0))
            (d_421_v5_)[index73_] = (d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))]
            (globalState).f7 = 548
        d_434_v15_: _dafny.Map
        d_434_v15_ = _dafny.Map({d_424_v7_: d_417_v1_})
        rhs71_ = len((D11_DC19((d_420_v4_).set((0) - (self.f14), (d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))]))).cf27)
        rhs72_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymnrs"))).set(default__.safeIndex((d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymnrs")))), _dafny.CodePoint('r'))), (d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))]) for d_435_i3_ in range(default__.abs(96))])
        rhs73_ = d_422_v6_
        rhs74_ = default__.fm43(self.f25, globalState)
        rhs75_ = default__.fm42(True, (d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))], (-970) - (self.f14), (d_417_v1_).issubset(((d_434_v15_)[D9_DC15(self.f25, self.f25, d_419_v3_, d_430_v13_, d_422_v6_)] if (D9_DC15(self.f25, self.f25, d_419_v3_, d_430_v13_, d_422_v6_)) in (d_434_v15_) else d_417_v1_)), globalState)
        lhs49_ = self
        lhs49_.f14 = rhs71_
        d_429_v12_ = rhs72_
        d_422_v6_ = rhs73_
        r2 = rhs74_
        d_427_v10_ = rhs75_
        (self).f14 = default__.safeModuloInt((d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))], (d_421_v5_)[default__.safeIndex(779, (d_421_v5_).length(0))])
        d_436_v16_: _dafny.Array
        def lambda28_(d_437_i4_):
            return _dafny.CodePoint('c')

        init18_ = lambda28_
        nw96_ = _dafny.Array(None, 10)
        for i0_18_ in range(nw96_.length(0)):
            nw96_[i0_18_] = init18_(i0_18_)
        d_436_v16_ = nw96_
        r0 = d_436_v16_
        r1 = -994
        d_438_v17_: D7
        d_438_v17_ = D7_DC12()
        def lambda29_(source4_):
            if source4_.is_DC12:
                if self.f25:
                    return self.f25
                elif True:
                    return self.f25
            elif True:
                d_439___mcc_h0_ = source4_.cf15
                d_440_cf15_ = d_439___mcc_h0_
                return self.f25

        r2 = lambda29_(d_438_v17_)
        d_441_v18_: _dafny.Set
        d_441_v18_ = _dafny.Set({d_422_v6_})
        d_442_v19_: int
        d_442_v19_ = len(d_441_v18_)
        r3 = (0) - ((d_442_v19_))
        return r0, r1, r2, r3


class C2(T1):
    def  __init__(self):
        self._f16: bool = False
        self._f17: bool = False
        self.f24: D3 = D3.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
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
        (self).f24 = f24
        (self).f16 = f16
        (self)._f17 = f17

    def fm3(self, globalState):
        def iife48_():
            coll30_ = _dafny.Map()
            compr_30_: int
            for compr_30_ in _dafny.IntegerRange(-864, 515):
                d_443_v0_: int = compr_30_
                if ((-864) <= (d_443_v0_)) and ((d_443_v0_) < (515)):
                    coll30_[(d_443_v0_) * (832)] = 330
            return _dafny.Map(coll30_)
        def iife49_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in _dafny.IntegerRange(374, 74):
                d_444_v1_: int = compr_31_
                if ((374) <= (d_444_v1_)) and ((d_444_v1_) < (74)):
                    coll31_[default__.safeDivisionInt(d_444_v1_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jypfe"))), 105, (_dafny.MultiSet([(self).f17, self.f16])).cardinality, 562, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_445_i0_ in range(default__.abs(640))]))}))])))] = (0) - ((0) - (-325))
            return _dafny.Map(coll31_)
        def iife50_():
            def iife52_():
                coll34_ = _dafny.Map()
                compr_34_: _dafny.Seq
                for compr_34_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_446_i1_ in range(default__.abs(161))]): (self).f17})).keys.Elements:
                    d_447_v2_: _dafny.Seq = compr_34_
                    if (d_447_v2_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_446_i1_ in range(default__.abs(161))]): (self).f17})):
                        coll34_[d_447_v2_] = 669
                return _dafny.Map(coll34_)
            coll32_ = _dafny.Set()
            def iife51_():
                coll33_ = _dafny.Map()
                compr_33_: _dafny.Seq
                for compr_33_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_446_i1_ in range(default__.abs(161))]): (self).f17})).keys.Elements:
                    d_447_v2_: _dafny.Seq = compr_33_
                    if (d_447_v2_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_446_i1_ in range(default__.abs(161))]): (self).f17})):
                        coll33_[d_447_v2_] = 669
                return _dafny.Map(coll33_)
            compr_32_: _dafny.Seq
            for compr_32_ in (iife51_()
            ).keys.Elements:
                d_448_v3_: _dafny.Seq = compr_32_
                if (d_448_v3_) in (iife52_()
                ):
                    coll32_ = coll32_.union(_dafny.Set([d_448_v3_]))
            return _dafny.Set(coll32_)
        return ((iife48_()
        ) | (iife49_()
        )) | (_dafny.Map({-423: len(iife50_()
        )}))

    def fm33(self, globalState):
        source5_ = D6_DC9(_dafny.CodePoint('k'))
        if source5_.is_DC10:
            return _dafny.SeqWithoutIsStrInference([False, (self).f17])
        elif True:
            d_449___mcc_h0_ = source5_.cf14
            d_450_cf14_ = d_449___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([self.f16])

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_451_v0_: _dafny.Array
        nw97_ = _dafny.Array(_dafny.CodePoint('D'), 19)
        d_451_v0_ = nw97_
        d_452_v1_: _dafny.Map
        d_452_v1_ = _dafny.Map({p0: p0})
        index74_ = default__.safeIndex(349, (d_451_v0_).length(0))
        (d_451_v0_)[index74_] = default__.fm34(p0, ((d_452_v1_)[p0] if (p0) in (d_452_v1_) else p0), p1, p0, globalState)
        d_453_v2_: str
        d_453_v2_ = _dafny.CodePoint('x')
        index75_ = default__.safeIndex(349, (d_451_v0_).length(0))
        (d_451_v0_)[index75_] = d_453_v2_
        d_454_i0_: int
        d_454_i0_ = 0
        with _dafny.label("3"):
            while True:
                with _dafny.c_label("3"):
                    if (d_454_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_454_i0_ = (d_454_i0_) + (1)
                    d_455_v3_: T0
                    nw98_ = C1()
                    nw98_.ctor__(p1, p0)
                    d_455_v3_ = nw98_
                    d_456_v4_: _dafny.MultiSet
                    d_456_v4_ = _dafny.MultiSet([d_455_v3_, d_455_v3_, d_455_v3_])
                    d_457_v5_: _dafny.Map
                    d_458_v6_: bool
                    out12_: _dafny.Map
                    out13_: bool
                    out12_, out13_ = (self).m12((self).f17, len(default__.fm35(p0, False, p0, (d_451_v0_)[default__.safeIndex(349, (d_451_v0_).length(0))], globalState)), d_456_v4_, not((self).f17), globalState)
                    d_457_v5_ = out12_
                    d_458_v6_ = out13_
                    d_459_v7_: _dafny.Array
                    nw99_ = _dafny.Array(_dafny.Array(None, 0), 24)
                    d_459_v7_ = nw99_
                    d_460_v8_: _dafny.Seq
                    d_460_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kiji"))
                    d_461_v9_: _dafny.Map
                    d_461_v9_ = _dafny.Map({d_458_v6_: d_455_v3_.f14})
                    d_462_v10_: _dafny.MultiSet
                    d_462_v10_ = _dafny.MultiSet([p1])
                    d_463_v11_: _dafny.Set
                    d_463_v11_ = _dafny.Set({p0})
                    d_464_v12_: _dafny.Seq
                    d_464_v12_ = _dafny.SeqWithoutIsStrInference([((d_462_v10_)[self.f16] if (self.f16) in (d_462_v10_) else 419), len(d_463_v11_)])
                    d_465_v13_: _dafny.Set
                    d_465_v13_ = _dafny.Set({(d_464_v12_)[default__.safeIndex(d_455_v3_.f14, len(d_464_v12_))]})
                    d_466_v14_: _dafny.Array
                    nw100_ = _dafny.Array(None, 28)
                    nw100_[int(0)] = p0
                    nw100_[int(1)] = p0
                    nw100_[int(2)] = d_455_v3_.f14
                    nw100_[int(3)] = 748
                    nw100_[int(4)] = 525
                    nw100_[int(5)] = p0
                    nw100_[int(6)] = p0
                    nw100_[int(7)] = len(d_452_v1_)
                    nw100_[int(8)] = (0) - (len(d_460_v8_))
                    nw100_[int(9)] = p0
                    nw100_[int(10)] = 724
                    nw100_[int(11)] = p0
                    nw100_[int(12)] = p0
                    nw100_[int(13)] = len(d_452_v1_)
                    nw100_[int(14)] = (0) - (-858)
                    nw100_[int(15)] = p0
                    nw100_[int(16)] = p0
                    nw100_[int(17)] = d_455_v3_.f14
                    nw100_[int(18)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_461_v9_]))).cardinality
                    nw100_[int(19)] = len(_dafny.Set({d_455_v3_.f14, d_455_v3_.f14, d_455_v3_.f14, len(d_460_v8_), p0}))
                    nw100_[int(20)] = p0
                    nw100_[int(21)] = -302
                    nw100_[int(22)] = d_455_v3_.f14
                    nw100_[int(23)] = (0) - (p0)
                    nw100_[int(24)] = d_455_v3_.f14
                    nw100_[int(25)] = len(d_465_v13_)
                    nw100_[int(26)] = p0
                    nw100_[int(27)] = d_455_v3_.f14
                    d_466_v14_ = nw100_
                    index76_ = default__.safeIndex(926, (d_459_v7_).length(0))
                    (d_459_v7_)[index76_] = d_466_v14_
                    index77_ = default__.safeIndex(926, (d_459_v7_).length(0))
                    nw101_ = _dafny.Array(int(0), 26)
                    (d_459_v7_)[index77_] = nw101_
                    d_467_v15_: _dafny.Set
                    d_467_v15_ = _dafny.Set({(self.f24).cf3, not(p1)})
                    (self).f16 = not((d_467_v15_).isdisjoint(default__.fm44(p1, self.f16, globalState)))
                    d_468_v16_: _dafny.Array
                    nw102_ = _dafny.Array(_dafny.Set({}), 6)
                    d_468_v16_ = nw102_
                    d_469_v17_: int
                    d_469_v17_ = p0
                    d_470_v18_: _dafny.Map
                    d_470_v18_ = _dafny.Map({d_468_v16_: (d_469_v17_)})
                    d_470_v18_ = (d_470_v18_).set(d_468_v16_, default__.fm0((self).f17, globalState))
                    pass
            pass
        d_471_v19_: _dafny.Array
        nw103_ = _dafny.Array(False, 6)
        d_471_v19_ = nw103_
        d_472_v20_: _dafny.Array
        nw104_ = _dafny.Array(int(0), 3)
        d_472_v20_ = nw104_
        d_473_v21_: _dafny.Map
        d_473_v21_ = _dafny.Map({_dafny.Map({p1: 333}): p0})
        d_474_v22_: _dafny.Seq
        d_474_v22_ = _dafny.SeqWithoutIsStrInference([(self).f17, p1, p1])
        d_475_v23_: _dafny.Array
        d_475_v23_ = d_471_v19_
        d_476_v24_: _dafny.Map
        d_476_v24_ = _dafny.Map({(self).f17: p0})
        d_477_v25_: _dafny.Map
        d_477_v25_ = _dafny.Map({p0: _dafny.Map({d_476_v24_: p0})})
        d_478_v26_: _dafny.Seq
        d_478_v26_ = d_474_v22_
        rhs76_ = (d_471_v19_)
        rhs77_ = default__.fm43(p1, globalState)
        rhs78_ = d_472_v20_
        rhs79_ = ((d_477_v25_)[p0] if (p0) in (d_477_v25_) else d_473_v21_)
        rhs80_ = (d_478_v26_)
        lhs50_ = self
        d_471_v19_ = rhs76_
        lhs50_.f16 = rhs77_
        d_472_v20_ = rhs78_
        d_473_v21_ = rhs79_
        d_474_v22_ = rhs80_
        d_479_v27_: _dafny.Set
        d_479_v27_ = _dafny.Set({d_453_v2_})
        d_479_v27_ = default__.fm45((d_451_v0_)[default__.safeIndex(349, (d_451_v0_).length(0))], globalState)
        r0 = (0) - (default__.safeModuloInt(p0, p0))
        d_480_v28_: _dafny.Map
        d_480_v28_ = _dafny.Map({p0: self.f16})
        d_481_v29_: _dafny.Seq
        d_481_v29_ = _dafny.SeqWithoutIsStrInference([d_480_v28_])
        d_482_i1_: int
        d_482_i1_ = 0
        with _dafny.label("4"):
            def lambda30_(source6_):
                if source6_.is_DC12:
                    if (self).f17:
                        return (self).f17
                    elif True:
                        return self.f16
                elif True:
                    d_499___mcc_h0_ = source6_.cf15
                    d_500_cf15_ = d_499___mcc_h0_
                    return (self).f17

            while lambda30_(D7_DC11((((d_481_v29_)[default__.safeIndex(p0, len(d_481_v29_))]).set(p0, (self).f17)).set(p0, self.f16))):
                with _dafny.c_label("4"):
                    if (d_482_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_482_i1_ = (d_482_i1_) + (1)
                    d_483_v30_: _dafny.Array
                    nw105_ = _dafny.Array(_dafny.Set({}), 6)
                    d_483_v30_ = nw105_
                    d_484_v31_: C0
                    nw106_ = C0()
                    nw106_.ctor__((d_483_v30_ if False else d_483_v30_), -730)
                    d_484_v31_ = nw106_
                    d_485_v32_: _dafny.Map
                    d_485_v32_ = _dafny.Map({d_453_v2_: (0) - (p0)})
                    d_485_v32_ = ((default__.fm46(self.f16, False, globalState)) | (d_485_v32_)) | (d_485_v32_)
                    if p1:
                        d_486_v33_: _dafny.Seq
                        d_486_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftj"))
                        d_487_v34_: _dafny.Array
                        nw107_ = _dafny.Array(None, 9)
                        nw107_[int(0)] = d_486_v33_
                        nw107_[int(1)] = d_486_v33_
                        nw107_[int(2)] = d_486_v33_
                        nw107_[int(3)] = d_486_v33_
                        nw107_[int(4)] = d_486_v33_
                        nw107_[int(5)] = d_486_v33_
                        nw107_[int(6)] = d_486_v33_
                        nw107_[int(7)] = (d_486_v33_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utmfmeu")))
                        nw107_[int(8)] = _dafny.SeqWithoutIsStrInference([(D6_DC9(d_453_v2_)).cf14 for d_488_i2_ in range(default__.abs(925))])
                        d_487_v34_ = nw107_
                        d_487_v34_ = d_487_v34_
                        d_489_v35_: _dafny.Map
                        d_489_v35_ = _dafny.Map({(self).f17: (d_451_v0_)[default__.safeIndex(349, (d_451_v0_).length(0))]})
                        d_490_v36_: _dafny.Map
                        d_490_v36_ = _dafny.Map({d_489_v35_: (d_484_v31_.f27) + (p0)})
                        d_490_v36_ = (d_490_v36_) | (d_490_v36_)
                        d_491_v37_: _dafny.Array
                        nw108_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_491_v37_ = nw108_
                        index78_ = default__.safeIndex(743, (d_491_v37_).length(0))
                        (d_491_v37_)[index78_] = d_487_v34_
                        index79_ = default__.safeIndex(743, (d_491_v37_).length(0))
                        (d_491_v37_)[index79_] = d_487_v34_
                        d_492_v38_: D3
                        d_492_v38_ = D3_DC4(not((D3_DC3(p1)).cf3), d_484_v31_.f27, len(d_480_v28_), True, True)
                        d_492_v38_ = d_492_v38_
                        (self).f16 = not((p0) <= (-834))
                    elif True:
                        (self).f16 = ((self).f17 if p1 else (self).f17)
                        (d_484_v31_).f27 = (0) - (-26)
                        d_493_v39_: _dafny.Seq
                        d_493_v39_ = _dafny.SeqWithoutIsStrInference([d_475_v23_])
                        d_494_v40_: _dafny.Set
                        d_494_v40_ = _dafny.Set({d_476_v24_, d_476_v24_})
                        rhs81_ = ((_dafny.Set({d_476_v24_, d_476_v24_})) - (d_494_v40_)).isdisjoint(d_494_v40_)
                        rhs82_ = (d_493_v39_) + (d_493_v39_)
                        lhs51_ = self
                        lhs51_.f16 = rhs81_
                        d_493_v39_ = rhs82_
                        d_495_v41_: _dafny.Array
                        nw109_ = _dafny.Array(None, 9)
                        d_495_v41_ = nw109_
                        d_495_v41_ = d_495_v41_
                        (globalState).f7 = ((0) - (d_484_v31_.f27)) + ((372) + (((d_452_v1_)[p0] if (p0) in (d_452_v1_) else d_484_v31_.f27)))
                    d_496_v42_: _dafny.MultiSet
                    d_496_v42_ = _dafny.MultiSet([(self).f17])
                    d_497_v43_: _dafny.Seq
                    d_497_v43_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: p1})), d_484_v31_.f27])
                    d_498_v44_: _dafny.Map
                    d_498_v44_ = _dafny.Map({d_484_v31_.f27: d_497_v43_})
                    (globalState).f7 = default__.safeModuloInt((d_496_v42_).cardinality, ((d_497_v43_)[default__.safeIndex(len(d_498_v44_), len(d_497_v43_))]) * ((d_497_v43_)[default__.safeIndex(190, len(d_497_v43_))]))
                    pass
            pass
        r0 = default__.safeModuloInt(default__.fm0(p1, globalState), default__.fm0(self.f16, globalState))
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_501_v0_: _dafny.Seq
        d_501_v0_ = _dafny.SeqWithoutIsStrInference([True])
        d_502_v1_: _dafny.Map
        d_502_v1_ = _dafny.Map({p0: d_501_v0_})
        (self).f16 = (len(d_502_v1_)) == (p3)
        d_503_v2_: _dafny.Array
        nw110_ = _dafny.Array(int(0), 15)
        d_503_v2_ = nw110_
        d_503_v2_ = d_503_v2_
        d_504_v3_: _dafny.Array
        nw111_ = _dafny.Array(None, 25)
        nw111_[int(0)] = p2
        nw111_[int(1)] = self.f16
        nw111_[int(2)] = (self.f24).cf3
        nw111_[int(3)] = (self).f17
        nw111_[int(4)] = True
        nw111_[int(5)] = p2
        nw111_[int(6)] = self.f16
        nw111_[int(7)] = self.f16
        nw111_[int(8)] = False
        nw111_[int(9)] = p2
        nw111_[int(10)] = self.f16
        nw111_[int(11)] = (self).f17
        nw111_[int(12)] = self.f16
        nw111_[int(13)] = p2
        nw111_[int(14)] = p2
        nw111_[int(15)] = self.f16
        nw111_[int(16)] = p2
        nw111_[int(17)] = p2
        nw111_[int(18)] = (self).f17
        nw111_[int(19)] = (self).f17
        nw111_[int(20)] = not((self).f17)
        nw111_[int(21)] = self.f16
        nw111_[int(22)] = p2
        nw111_[int(23)] = self.f16
        nw111_[int(24)] = p2
        d_504_v3_ = nw111_
        d_505_v4_: _dafny.Array
        nw112_ = _dafny.Array(None, 25)
        nw112_[int(0)] = d_504_v3_
        nw112_[int(1)] = d_504_v3_
        nw112_[int(2)] = d_504_v3_
        nw112_[int(3)] = d_504_v3_
        nw112_[int(4)] = d_504_v3_
        nw112_[int(5)] = d_504_v3_
        nw112_[int(6)] = d_504_v3_
        nw112_[int(7)] = d_504_v3_
        nw112_[int(8)] = d_504_v3_
        nw112_[int(9)] = d_504_v3_
        nw112_[int(10)] = d_504_v3_
        nw112_[int(11)] = d_504_v3_
        nw112_[int(12)] = d_504_v3_
        nw112_[int(13)] = d_504_v3_
        nw112_[int(14)] = d_504_v3_
        nw112_[int(15)] = d_504_v3_
        nw112_[int(16)] = d_504_v3_
        nw112_[int(17)] = d_504_v3_
        nw112_[int(18)] = d_504_v3_
        nw112_[int(19)] = d_504_v3_
        nw112_[int(20)] = d_504_v3_
        nw112_[int(21)] = d_504_v3_
        nw112_[int(22)] = d_504_v3_
        nw112_[int(23)] = d_504_v3_
        nw112_[int(24)] = d_504_v3_
        d_505_v4_ = nw112_
        d_505_v4_ = d_505_v4_
        hi5_ = 71
        for d_506_i0_ in range(815, hi5_):
            d_507_v5_: C1
            nw113_ = C1()
            nw113_.ctor__(p2, p3)
            d_507_v5_ = nw113_
            d_508_v6_: _dafny.Array
            nw114_ = _dafny.Array(_dafny.Set({}), 12)
            d_508_v6_ = nw114_
            d_509_v7_: _dafny.Set
            d_509_v7_ = _dafny.Set({p0})
            index80_ = default__.safeIndex(902, (d_508_v6_).length(0))
            (d_508_v6_)[index80_] = d_509_v7_
            index81_ = default__.safeIndex(902, (d_508_v6_).length(0))
            (d_508_v6_)[index81_] = (d_509_v7_) | (_dafny.Set({p0, p0}))
            (self).f16 = default__.fm43((self).f17, globalState)
            index82_ = default__.safeIndex(570, (d_503_v2_).length(0))
            (d_503_v2_)[index82_] = (0) - (d_506_i0_)
            d_510_v8_: int
            d_510_v8_ = default__.safeDivisionInt(590, p3)
            d_511_v9_: _dafny.Map
            d_511_v9_ = _dafny.Map({p3: d_504_v3_})
            d_512_v10_: _dafny.Seq
            d_512_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p3: p2})])
            d_513_v11_: _dafny.Map
            d_513_v11_ = _dafny.Map({-48: d_507_v5_.f25})
            d_514_v12_: _dafny.Seq
            d_514_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktd"))
            index83_ = default__.safeIndex(570, (d_503_v2_).length(0))
            rhs83_ = len(d_511_v9_)
            rhs84_ = not(((d_512_v10_)[default__.safeIndex(p1, len(d_512_v10_))]) != (d_513_v11_))
            rhs85_ = len(d_514_v12_)
            rhs86_ = (0) - (p1)
            rhs87_ = default__.fm47((p1) * (p3), globalState)
            lhs52_ = globalState
            lhs53_ = d_507_v5_
            lhs54_ = globalState
            lhs55_ = d_503_v2_
            lhs56_ = default__.safeIndex(570, (d_503_v2_).length(0))
            lhs52_.f2 = rhs83_
            lhs53_.f25 = rhs84_
            lhs54_.f2 = rhs85_
            lhs55_[lhs56_] = rhs86_
            d_510_v8_ = rhs87_
        r0 = default__.safeDivisionInt(p1, p3)
        hi6_ = p3
        for d_515_i1_ in range(p3, hi6_):
            d_516_v13_: _dafny.Array
            nw115_ = _dafny.Array(None, 5)
            nw115_[int(0)] = p0
            nw115_[int(1)] = p0
            nw115_[int(2)] = p0
            nw115_[int(3)] = _dafny.CodePoint('i')
            nw115_[int(4)] = p0
            d_516_v13_ = nw115_
            index84_ = default__.safeIndex(521, (d_516_v13_).length(0))
            (d_516_v13_)[index84_] = p0
            index85_ = default__.safeIndex(521, (d_516_v13_).length(0))
            (d_516_v13_)[index85_] = p0
            d_517_v14_: _dafny.MultiSet
            d_517_v14_ = _dafny.MultiSet([p2, (self).f17, p2])
            d_518_v15_: _dafny.Seq
            d_518_v15_ = _dafny.SeqWithoutIsStrInference([d_517_v14_])
            index86_ = default__.safeIndex(459, (d_503_v2_).length(0))
            (d_503_v2_)[index86_] = d_515_i1_
            index87_ = default__.safeIndex(459, (d_503_v2_).length(0))
            rhs88_ = (d_518_v15_) + ((d_518_v15_) + (d_518_v15_))
            rhs89_ = 321
            rhs90_ = (648) - (p1)
            rhs91_ = (p1) + (655)
            lhs57_ = d_503_v2_
            lhs58_ = default__.safeIndex(459, (d_503_v2_).length(0))
            d_518_v15_ = rhs88_
            lhs57_[lhs58_] = rhs89_
            r0 = rhs90_
            r0 = rhs91_
            d_519_v16_: _dafny.Map
            d_519_v16_ = _dafny.Map({False: p0})
            if ((self).f17 if ((self).f17) in (d_519_v16_) else (d_501_v0_)[default__.safeIndex((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p3, p3]))).cardinality, len(d_501_v0_))]):
                (globalState).f7 = default__.safeModuloInt((d_503_v2_)[default__.safeIndex(459, (d_503_v2_).length(0))], (p1 if self.f16 else -740))
                d_520_v17_: _dafny.MultiSet
                d_520_v17_ = _dafny.MultiSet([p1])
                d_521_v18_: _dafny.Map
                d_521_v18_ = _dafny.Map({(p3) + (p1): d_520_v17_})
                d_521_v18_ = d_521_v18_
                d_522_v19_: _dafny.Seq
                d_522_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awfmvcj"))
                d_522_v19_ = (d_522_v19_ if not (False) or (self.f16) else d_522_v19_)
                (self).f16 = (d_501_v0_)[default__.safeIndex((d_503_v2_)[default__.safeIndex(459, (d_503_v2_).length(0))], len(d_501_v0_))]
                d_523_v20_: _dafny.MultiSet
                d_523_v20_ = _dafny.MultiSet([_dafny.CodePoint('w'), _dafny.CodePoint('c')])
                d_524_v21_: T0
                nw116_ = C1()
                nw116_.ctor__(p2, p3)
                d_524_v21_ = nw116_
                d_525_v22_: _dafny.MultiSet
                d_525_v22_ = _dafny.MultiSet([d_524_v21_])
                d_526_v23_: _dafny.Set
                d_526_v23_ = _dafny.Set({(self).f17})
                d_527_v24_: _dafny.Seq
                d_527_v24_ = _dafny.SeqWithoutIsStrInference([len(d_526_v23_), d_524_v21_.f14])
                d_528_v25_: _dafny.Set
                d_528_v25_ = _dafny.Set({(d_503_v2_)[default__.safeIndex(459, (d_503_v2_).length(0))], p3})
                d_529_v26_: _dafny.Array
                nw117_ = _dafny.Array(int(0), 19)
                d_529_v26_ = nw117_
                d_530_v27_: D9
                d_530_v27_ = D9_DC15(p2, p2, d_528_v25_, d_529_v26_, d_504_v3_)
                d_531_v28_: _dafny.Map
                d_532_v29_: bool
                out14_: _dafny.Map
                out15_: bool
                out14_, out15_ = (self).m12(((self).f17 if self.f16 else self.f16), (default__.fm0(self.f16, globalState)) * ((d_523_v20_).cardinality), ((d_525_v22_).set(d_524_v21_, default__.abs(len(d_527_v24_)))).intersection(d_525_v22_), (d_528_v25_).ispropersubset((d_530_v27_).cf20), globalState)
                d_531_v28_ = out14_
                d_532_v29_ = out15_
            elif True:
                d_533_v30_: _dafny.Map
                d_533_v30_ = _dafny.Map({p3: (d_516_v13_)[default__.safeIndex(521, (d_516_v13_).length(0))]})
                d_534_v31_: _dafny.Seq
                d_534_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atptjlkbt"))
                d_535_v32_: _dafny.Seq
                d_535_v32_ = _dafny.SeqWithoutIsStrInference([p1, d_515_i1_, len(d_534_v31_)])
                d_536_v33_: _dafny.Map
                d_536_v33_ = _dafny.Map({d_533_v30_: d_535_v32_})
                d_536_v33_ = d_536_v33_
                d_537_v34_: _dafny.Set
                d_537_v34_ = _dafny.Set({p2})
                d_538_v35_: _dafny.Array
                nw118_ = _dafny.Array(None, 13)
                nw118_[int(0)] = _dafny.Set({not((self).f17), self.f16, self.f16, p2, (self).f17})
                nw118_[int(1)] = d_537_v34_
                nw118_[int(2)] = d_537_v34_
                nw118_[int(3)] = _dafny.Set({p2, (self).f17})
                nw118_[int(4)] = d_537_v34_
                nw118_[int(5)] = _dafny.Set({p2, p2})
                nw118_[int(6)] = d_537_v34_
                nw118_[int(7)] = d_537_v34_
                nw118_[int(8)] = d_537_v34_
                nw118_[int(9)] = d_537_v34_
                nw118_[int(10)] = d_537_v34_
                nw118_[int(11)] = d_537_v34_
                nw118_[int(12)] = d_537_v34_
                d_538_v35_ = nw118_
                d_539_v36_: C0
                nw119_ = C0()
                nw119_.ctor__(d_538_v35_, (d_515_i1_) * (p3))
                d_539_v36_ = nw119_
                d_540_v37_: _dafny.Array
                def lambda31_(d_541_v0_):
                    def lambda32_(d_542_i2_):
                        return (d_542_i2_) * (len(_dafny.Map({len(d_541_v0_): self.f16})))

                    return lambda32_

                init19_ = lambda31_(d_501_v0_)
                nw120_ = _dafny.Array(None, 12)
                for i0_19_ in range(nw120_.length(0)):
                    nw120_[i0_19_] = init19_(i0_19_)
                d_540_v37_ = nw120_
                d_540_v37_ = d_540_v37_
                d_543_v38_: _dafny.Map
                d_543_v38_ = _dafny.Map({(0) - ((d_503_v2_)[default__.safeIndex(459, (d_503_v2_).length(0))]): p2})
                d_544_v39_: _dafny.MultiSet
                d_544_v39_ = _dafny.MultiSet([default__.fm42(((d_543_v38_)[d_515_i1_] if (d_515_i1_) in (d_543_v38_) else True), d_539_v36_.f27, d_539_v36_.f27, self.f16, globalState)])
                d_544_v39_ = d_544_v39_
                index88_ = default__.safeIndex(67, (d_504_v3_).length(0))
                (d_504_v3_)[index88_] = default__.fm43(self.f16, globalState)
                d_545_v40_: _dafny.Seq
                d_545_v40_ = _dafny.SeqWithoutIsStrInference([d_504_v3_, d_504_v3_])
                d_546_v41_: _dafny.Array
                d_546_v41_ = (d_545_v40_)[default__.safeIndex((d_503_v2_)[default__.safeIndex(459, (d_503_v2_).length(0))], len(d_545_v40_))]
                index89_ = default__.safeIndex(521, (d_516_v13_).length(0))
                index90_ = default__.safeIndex(67, (d_504_v3_).length(0))
                index91_ = default__.safeIndex(459, (d_503_v2_).length(0))
                rhs92_ = p0
                rhs93_ = p2
                rhs94_ = (0) - (len((D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsgrpsbcd")))).cf29))
                rhs95_ = d_504_v3_
                lhs59_ = d_516_v13_
                lhs60_ = default__.safeIndex(521, (d_516_v13_).length(0))
                lhs61_ = d_504_v3_
                lhs62_ = default__.safeIndex(67, (d_504_v3_).length(0))
                lhs63_ = d_503_v2_
                lhs64_ = default__.safeIndex(459, (d_503_v2_).length(0))
                lhs59_[lhs60_] = rhs92_
                lhs61_[lhs62_] = rhs93_
                lhs63_[lhs64_] = rhs94_
                d_546_v41_ = rhs95_
            (globalState).f2 = p1
        r0 = (default__.fm0(p2, globalState)) - (p3)
        return r0

    def m12(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_547_v0_: _dafny.Map
        d_547_v0_ = _dafny.Map({True: p0})
        d_548_v1_: _dafny.Map
        d_548_v1_ = _dafny.Map({self.f16: d_547_v0_})
        d_549_v2_: _dafny.Map
        d_549_v2_ = d_548_v1_
        d_549_v2_ = d_549_v2_
        r1 = p0
        if default__.fm43(p3, globalState):
            d_550_v3_: _dafny.Array
            nw121_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_550_v3_ = nw121_
            d_551_v4_: _dafny.Seq
            d_551_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yktx"))
            index92_ = default__.safeIndex(545, (d_550_v3_).length(0))
            (d_550_v3_)[index92_] = (d_551_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))
            d_552_v5_: str
            d_552_v5_ = _dafny.CodePoint('s')
            index93_ = default__.safeIndex(545, (d_550_v3_).length(0))
            (d_550_v3_)[index93_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilx"))) + (((d_551_v4_).set(default__.safeIndex(p1, len(d_551_v4_)), d_552_v5_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))
            r1 = (self).f17
            d_553_v6_: _dafny.Array
            nw122_ = _dafny.Array(_dafny.Set({}), 1)
            d_553_v6_ = nw122_
            d_554_v7_: C0
            nw123_ = C0()
            nw123_.ctor__(d_553_v6_, p1)
            d_554_v7_ = nw123_
            d_555_v9_: _dafny.Seq
            d_555_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, not(p0), (self).f17])])
            d_556_v10_: _dafny.MultiSet
            d_556_v10_ = _dafny.MultiSet([(self).f17])
            def iife53_():
                coll35_ = _dafny.Map()
                compr_35_: _dafny.Seq
                for compr_35_ in (d_555_v9_).Elements:
                    d_557_v8_: _dafny.Seq = compr_35_
                    if (d_557_v8_) in (d_555_v9_):
                        coll35_[d_557_v8_] = False
                return _dafny.Map(coll35_)
            (globalState).f7 = (len(iife53_()
            )) - (((d_556_v10_)[True] if (True) in (d_556_v10_) else p1))
            if (self).f17:
                d_558_v11_: _dafny.MultiSet
                d_558_v11_ = _dafny.MultiSet([-551])
                d_559_v12_: _dafny.Seq
                d_559_v12_ = _dafny.SeqWithoutIsStrInference([d_558_v11_])
                d_560_v13_: _dafny.Array
                nw124_ = _dafny.Array(None, 12)
                nw124_[int(0)] = (self).f17
                nw124_[int(1)] = not((self).f17)
                nw124_[int(2)] = p3
                nw124_[int(3)] = p3
                nw124_[int(4)] = (self.f24).cf3
                nw124_[int(5)] = (self).f17
                nw124_[int(6)] = p0
                nw124_[int(7)] = False
                nw124_[int(8)] = p0
                nw124_[int(9)] = p0
                nw124_[int(10)] = p0
                nw124_[int(11)] = self.f16
                d_560_v13_ = nw124_
                d_561_v14_: _dafny.Map
                d_561_v14_ = _dafny.Map({d_560_v13_: p1})
                d_562_v15_: _dafny.Set
                d_562_v15_ = _dafny.Set({not((self).f17)})
                d_563_v16_: _dafny.Set
                d_563_v16_ = _dafny.Set({(d_558_v11_).cardinality, (len(d_562_v15_)) - (p1)})
                d_564_v17_: _dafny.Map
                d_564_v17_ = _dafny.Map({d_554_v7_.f27: (d_554_v7_.f27) + (784)})
                d_565_v18_: _dafny.Seq
                d_565_v18_ = _dafny.SeqWithoutIsStrInference([951])
                d_566_v20_: _dafny.Map
                d_566_v20_ = _dafny.Map({d_554_v7_.f27: p0})
                def iife54_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in (d_566_v20_).keys.Elements:
                        d_567_v19_: int = compr_36_
                        if (d_567_v19_) in (d_566_v20_):
                            coll36_[default__.safeDivisionInt(d_567_v19_, ((d_556_v10_)[(self).f17] if ((self).f17) in (d_556_v10_) else len(d_551_v4_)))] = d_554_v7_.f27
                    return _dafny.Map(coll36_)
                rhs96_ = d_559_v12_
                rhs97_ = ((d_561_v14_) | (d_561_v14_)) | (d_561_v14_)
                rhs98_ = (d_547_v0_).set(self.f16, (self).f17)
                rhs99_ = ((d_563_v16_).intersection(d_563_v16_)).intersection((d_563_v16_) | (d_563_v16_))
                rhs100_ = ((d_564_v17_).set((d_565_v18_)[default__.safeIndex(len((_dafny.Map({p1: d_560_v13_})).set(p1, d_560_v13_)), len(d_565_v18_))], 831)) | (iife54_()
                )
                d_559_v12_ = rhs96_
                d_561_v14_ = rhs97_
                d_547_v0_ = rhs98_
                d_563_v16_ = rhs99_
                d_564_v17_ = rhs100_
                d_568_v21_: C0
                nw125_ = C0()
                nw125_.ctor__(d_553_v6_, d_554_v7_.f27)
                d_568_v21_ = nw125_
                (self).f16 = not(self.f16)
                d_569_v22_: _dafny.Map
                d_569_v22_ = _dafny.Map({(d_550_v3_)[default__.safeIndex(545, (d_550_v3_).length(0))]: not(not(not (not(False)) or (p0)))})
                r1 = ((d_569_v22_)[(d_550_v3_)[default__.safeIndex(545, (d_550_v3_).length(0))]] if ((d_550_v3_)[default__.safeIndex(545, (d_550_v3_).length(0))]) in (d_569_v22_) else (self).f17)
                d_570_v23_: _dafny.MultiSet
                d_570_v23_ = _dafny.MultiSet([d_552_v5_, d_552_v5_])
                (self).f16 = ((d_552_v5_) not in (d_570_v23_) if self.f16 else p0)
            elif True:
                (globalState).f7 = p1
                d_571_v24_: D9
                d_571_v24_ = D9_DC16(d_554_v7_.f27, p0)
                d_572_v25_: _dafny.Seq
                d_572_v25_ = _dafny.SeqWithoutIsStrInference([(self).f17, self.f16])
                d_573_v26_: _dafny.Array
                nw126_ = _dafny.Array(None, 6)
                nw126_[int(0)] = p1
                nw126_[int(1)] = (d_571_v24_).cf23
                nw126_[int(2)] = p1
                nw126_[int(3)] = len(d_572_v25_)
                nw126_[int(4)] = p1
                nw126_[int(5)] = p1
                d_573_v26_ = nw126_
                d_574_v27_: _dafny.Map
                d_574_v27_ = _dafny.Map({(p1) * (len(d_551_v4_)): d_573_v26_})
                d_574_v27_ = (d_574_v27_).set(len(d_572_v25_), d_573_v26_)
                d_575_v28_: _dafny.Map
                d_575_v28_ = _dafny.Map({True: d_551_v4_})
                d_576_v29_: _dafny.MultiSet
                d_576_v29_ = _dafny.MultiSet([((d_575_v28_)[p3] if (p3) in (d_575_v28_) else d_551_v4_), (d_550_v3_)[default__.safeIndex(545, (d_550_v3_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imiolrx"))])
                d_577_v30_: _dafny.Map
                d_577_v30_ = _dafny.Map({d_554_v7_.f27: d_554_v7_.f27})
                d_578_v31_: _dafny.Map
                d_578_v31_ = _dafny.Map({(d_576_v29_).issubset(d_576_v29_): d_577_v30_})
                d_578_v31_ = (d_578_v31_).set(self.f16, d_577_v30_)
                d_579_v32_: _dafny.Seq
                d_579_v32_ = d_572_v25_
                d_580_v33_: _dafny.Set
                d_580_v33_ = _dafny.Set({not(p3)})
                d_581_v34_: _dafny.Set
                d_581_v34_ = _dafny.Set({len(d_580_v33_), d_554_v7_.f27, 627})
                d_582_v35_: D4
                d_582_v35_ = D4_DC7(d_572_v25_, len(d_581_v34_))
                d_583_v36_: _dafny.Map
                d_583_v36_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxnylel")): p1})
                d_584_v37_: D13
                d_584_v37_ = D13_DC25(d_583_v36_)
                d_585_v38_: _dafny.Map
                d_585_v38_ = _dafny.Map({d_547_v0_: len((d_584_v37_).cf38)})
                d_586_v39_: _dafny.Map
                d_586_v39_ = _dafny.Map({(0) - ((d_582_v35_).cf12): d_585_v38_})
                d_586_v39_ = (d_586_v39_).set(p1, (default__.fm48(globalState)).set(d_547_v0_, d_554_v7_.f27))
                index94_ = default__.safeIndex(545, (d_550_v3_).length(0))
                (d_550_v3_)[index94_] = ((d_551_v4_) + ((d_550_v3_)[default__.safeIndex(545, (d_550_v3_).length(0))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouxahjk")))
        elif True:
            d_587_v40_: C1
            nw127_ = C1()
            nw127_.ctor__(self.f16, p1)
            d_587_v40_ = nw127_
            d_588_v41_: _dafny.Array
            nw128_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_588_v41_ = nw128_
            d_589_v42_: _dafny.Seq
            d_589_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
            index95_ = default__.safeIndex(385, (d_588_v41_).length(0))
            (d_588_v41_)[index95_] = d_589_v42_
            d_590_v43_: _dafny.Seq
            d_590_v43_ = _dafny.SeqWithoutIsStrInference([(self).f17])
            d_591_v44_: _dafny.Set
            d_591_v44_ = _dafny.Set({len(d_589_v42_), default__.fm0((self).f17, globalState)})
            d_592_v45_: _dafny.MultiSet
            d_592_v45_ = _dafny.MultiSet([d_591_v44_, d_591_v44_])
            d_593_v46_: _dafny.MultiSet
            d_593_v46_ = _dafny.MultiSet([d_587_v40_.f25, (self).f17, d_587_v40_.f25])
            d_594_v47_: _dafny.Array
            nw129_ = _dafny.Array(None, 3)
            nw129_[int(0)] = False
            nw129_[int(1)] = (d_593_v46_).issubset(_dafny.MultiSet((d_590_v43_).set(default__.safeIndex((d_592_v45_).cardinality, len(d_590_v43_)), self.f16)))
            nw129_[int(2)] = (p0) or (p0)
            d_594_v47_ = nw129_
            index96_ = default__.safeIndex(749, (d_594_v47_).length(0))
            (d_594_v47_)[index96_] = (_dafny.MultiSet(d_590_v43_)).isdisjoint(d_593_v46_)
            d_595_v48_: str
            d_595_v48_ = _dafny.CodePoint('u')
            d_596_v49_: _dafny.Map
            d_596_v49_ = _dafny.Map({d_595_v48_: p1})
            d_597_v50_: _dafny.Seq
            d_597_v50_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            index97_ = default__.safeIndex(385, (d_588_v41_).length(0))
            index98_ = default__.safeIndex(749, (d_594_v47_).length(0))
            rhs101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbphfkp"))
            rhs102_ = (p1) >= (default__.fm0(not(p3), globalState))
            rhs103_ = (0) - (default__.safeDivisionInt(((d_596_v49_)[_dafny.CodePoint('m')] if (_dafny.CodePoint('m')) in (d_596_v49_) else p1), (d_597_v50_)[default__.safeIndex(p1, len(d_597_v50_))]))
            rhs104_ = len(d_589_v42_)
            lhs65_ = d_588_v41_
            lhs66_ = default__.safeIndex(385, (d_588_v41_).length(0))
            lhs67_ = d_594_v47_
            lhs68_ = default__.safeIndex(749, (d_594_v47_).length(0))
            lhs69_ = globalState
            lhs70_ = globalState
            lhs65_[lhs66_] = rhs101_
            lhs67_[lhs68_] = rhs102_
            lhs69_.f2 = rhs103_
            lhs70_.f7 = rhs104_
            if d_587_v40_.f25:
                d_598_v51_: _dafny.Array
                def lambda33_(d_599_v40_):
                    def lambda34_(d_600_i0_):
                        return _dafny.Set({d_599_v40_.f25})

                    return lambda34_

                init20_ = lambda33_(d_587_v40_)
                nw130_ = _dafny.Array(None, 2)
                for i0_20_ in range(nw130_.length(0)):
                    nw130_[i0_20_] = init20_(i0_20_)
                d_598_v51_ = nw130_
                d_601_v52_: C0
                nw131_ = C0()
                nw131_.ctor__(d_598_v51_, (0) - (p1))
                d_601_v52_ = nw131_
                d_602_v53_: _dafny.Set
                d_602_v53_ = _dafny.Set({p3, True})
                d_603_v54_: _dafny.Map
                d_603_v54_ = _dafny.Map({default__.fm42(d_587_v40_.f25, p1, p1, d_587_v40_.f25, globalState): (d_602_v53_) | (d_602_v53_)})
                d_603_v54_ = (d_603_v54_).set((d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))], d_602_v53_)
                d_604_v55_: _dafny.MultiSet
                d_604_v55_ = _dafny.MultiSet([d_589_v42_, (d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))]])
                d_605_v56_: _dafny.Map
                d_605_v56_ = _dafny.Map({(self).f17: p1})
                d_606_v57_: _dafny.Map
                d_606_v57_ = _dafny.Map({not((d_594_v47_)[default__.safeIndex(749, (d_594_v47_).length(0))]): _dafny.CodePoint('k')})
                d_607_v58_: _dafny.Seq
                d_607_v58_ = _dafny.SeqWithoutIsStrInference([(self).f17, d_587_v40_.f25, True])
                d_608_v59_: _dafny.Map
                d_608_v59_ = _dafny.Map({d_601_v52_.f27: (d_593_v46_).cardinality})
                d_609_v60_: _dafny.Array
                nw132_ = _dafny.Array(None, 22)
                nw132_[int(0)] = (d_604_v55_).cardinality
                nw132_[int(1)] = d_601_v52_.f27
                nw132_[int(2)] = d_601_v52_.f27
                nw132_[int(3)] = len((d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))])
                nw132_[int(4)] = ((d_605_v56_)[(self).f17] if ((self).f17) in (d_605_v56_) else p1)
                nw132_[int(5)] = ((d_596_v49_)[d_595_v48_] if (d_595_v48_) in (d_596_v49_) else p1)
                nw132_[int(6)] = (len(((d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))]).set(default__.safeIndex(len(d_606_v57_), len((d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))])), d_595_v48_))) * (d_601_v52_.f27)
                nw132_[int(7)] = len((d_607_v58_))
                nw132_[int(8)] = (d_601_v52_.f27) * (d_601_v52_.f27)
                nw132_[int(9)] = d_601_v52_.f27
                nw132_[int(10)] = (p1) + (p1)
                nw132_[int(11)] = 249
                nw132_[int(12)] = d_601_v52_.f27
                nw132_[int(13)] = (0) - (len(d_602_v53_))
                nw132_[int(14)] = p1
                nw132_[int(15)] = d_601_v52_.f27
                nw132_[int(16)] = len((d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))])
                nw132_[int(17)] = p1
                nw132_[int(18)] = ((0) - (d_601_v52_.f27)) - (d_601_v52_.f27)
                nw132_[int(19)] = ((_dafny.MultiSet([d_608_v59_, d_608_v59_])).cardinality if (d_594_v47_)[default__.safeIndex(749, (d_594_v47_).length(0))] else d_601_v52_.f27)
                nw132_[int(20)] = default__.safeDivisionInt((0) - (d_601_v52_.f27), p1)
                nw132_[int(21)] = 637
                d_609_v60_ = nw132_
                index99_ = default__.safeIndex(421, (d_609_v60_).length(0))
                (d_609_v60_)[index99_] = d_601_v52_.f27
                index100_ = default__.safeIndex(421, (d_609_v60_).length(0))
                (d_609_v60_)[index100_] = d_601_v52_.f27
                d_610_v61_: _dafny.Map
                d_610_v61_ = _dafny.Map({p1: default__.fm43((self).f17, globalState)})
                d_611_v62_: _dafny.Map
                d_611_v62_ = _dafny.Map({((d_610_v61_)[d_601_v52_.f27] if (d_601_v52_.f27) in (d_610_v61_) else self.f16): d_609_v60_})
                d_611_v62_ = (d_611_v62_).set(False, d_609_v60_)
                (globalState).f7 = (d_609_v60_)[default__.safeIndex(421, (d_609_v60_).length(0))]
            elif True:
                d_612_v63_: _dafny.Set
                d_612_v63_ = _dafny.Set({(self).f17, self.f16})
                d_613_v64_: _dafny.Array
                nw133_ = _dafny.Array(None, 17)
                nw133_[int(0)] = p1
                nw133_[int(1)] = p1
                nw133_[int(2)] = p1
                nw133_[int(3)] = p1
                nw133_[int(4)] = -868
                nw133_[int(5)] = (0) - (p1)
                nw133_[int(6)] = p1
                nw133_[int(7)] = len(d_612_v63_)
                nw133_[int(8)] = p1
                nw133_[int(9)] = p1
                nw133_[int(10)] = p1
                nw133_[int(11)] = len(d_589_v42_)
                nw133_[int(12)] = p1
                nw133_[int(13)] = len(d_597_v50_)
                nw133_[int(14)] = p1
                nw133_[int(15)] = p1
                nw133_[int(16)] = p1
                d_613_v64_ = nw133_
                d_614_v65_: _dafny.Map
                d_614_v65_ = _dafny.Map({_dafny.MultiSet([d_613_v64_]): d_595_v48_})
                d_615_v66_: _dafny.MultiSet
                d_615_v66_ = _dafny.MultiSet([d_613_v64_])
                d_614_v65_ = (d_614_v65_).set(d_615_v66_, (d_595_v48_ if (d_590_v43_)[default__.safeIndex(p1, len(d_590_v43_))] else d_595_v48_))
                (globalState).f7 = (p1) * ((d_593_v46_).cardinality)
                (globalState).f7 = ((p1) - (p1)) + (p1)
                d_616_v67_: int
                d_617_v68_: bool
                d_618_v69_: str
                d_619_v70_: _dafny.Map
                out16_: int
                out17_: bool
                out18_: str
                out19_: _dafny.Map
                out16_, out17_, out18_, out19_ = (d_587_v40_).m13(not(not(not(not(p0)))), p3, globalState)
                d_616_v67_ = out16_
                d_617_v68_ = out17_
                d_618_v69_ = out18_
                d_619_v70_ = out19_
                (globalState).f7 = d_616_v67_
            d_620_v71_: _dafny.Set
            d_620_v71_ = _dafny.Set({d_595_v48_})
            d_620_v71_ = d_620_v71_
            d_621_v72_: _dafny.Map
            d_621_v72_ = _dafny.Map({p1: p1})
            d_622_v73_: _dafny.Array
            def lambda35_(d_623_p1_):
                def lambda36_(d_624_i1_):
                    return (d_624_i1_) * ((0) - (d_623_p1_))

                return lambda36_

            init21_ = lambda35_(p1)
            nw134_ = _dafny.Array(None, 17)
            for i0_21_ in range(nw134_.length(0)):
                nw134_[i0_21_] = init21_(i0_21_)
            d_622_v73_ = nw134_
            d_625_v74_: _dafny.MultiSet
            d_625_v74_ = _dafny.MultiSet([d_622_v73_, d_622_v73_, d_622_v73_, d_622_v73_, d_622_v73_])
            d_626_v75_: _dafny.Array
            nw135_ = _dafny.Array(None, 11)
            nw135_[int(0)] = p1
            nw135_[int(1)] = p1
            nw135_[int(2)] = len(d_621_v72_)
            nw135_[int(3)] = (0) - (p1)
            nw135_[int(4)] = (d_625_v74_).cardinality
            nw135_[int(5)] = p1
            nw135_[int(6)] = p1
            nw135_[int(7)] = p1
            nw135_[int(8)] = 330
            nw135_[int(9)] = (p1) - (p1)
            nw135_[int(10)] = len((d_588_v41_)[default__.safeIndex(385, (d_588_v41_).length(0))])
            d_626_v75_ = nw135_
            index101_ = default__.safeIndex(554, (d_622_v73_).length(0))
            (d_622_v73_)[index101_] = p1
            index102_ = default__.safeIndex(554, (d_622_v73_).length(0))
            (d_622_v73_)[index102_] = p1
        d_627_v76_: _dafny.Map
        d_627_v76_ = _dafny.Map({len(_dafny.Map({not(default__.fm43(p3, globalState)): p1})): p3})
        hi7_ = len(d_627_v76_)
        for d_628_i2_ in range(177, hi7_):
            d_629_v77_: _dafny.Seq
            d_629_v77_ = _dafny.SeqWithoutIsStrInference([p0])
            d_630_v78_: _dafny.Map
            d_630_v78_ = _dafny.Map({d_628_i2_: (d_629_v77_) + (d_629_v77_)})
            d_630_v78_ = (d_630_v78_ if p0 else d_630_v78_)
            d_631_v79_: C1
            nw136_ = C1()
            nw136_.ctor__(p3, (p1) + (879))
            d_631_v79_ = nw136_
            (globalState).f2 = p1
            d_632_v80_: _dafny.Array
            def lambda37_(d_633_v79_):
                def lambda38_(d_634_i3_):
                    return _dafny.Set({d_633_v79_.f25, self.f16})

                return lambda38_

            init22_ = lambda37_(d_631_v79_)
            nw137_ = _dafny.Array(None, 17)
            for i0_22_ in range(nw137_.length(0)):
                nw137_[i0_22_] = init22_(i0_22_)
            d_632_v80_ = nw137_
            d_635_v81_: C0
            nw138_ = C0()
            nw138_.ctor__(d_632_v80_, default__.safeDivisionInt(p1, d_628_i2_))
            d_635_v81_ = nw138_
        d_636_v82_: _dafny.Map
        d_636_v82_ = _dafny.Map({self.f16: p1})
        hi8_ = p1
        for d_637_i4_ in range(len(d_636_v82_), hi8_):
            (self).f16 = default__.fm43(p3, globalState)
            d_638_v83_: _dafny.Seq
            d_638_v83_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
            d_638_v83_ = d_638_v83_
            (globalState).f7 = p1
            d_639_v84_: _dafny.MultiSet
            d_639_v84_ = _dafny.MultiSet([not((self).f17), self.f16, False])
            d_640_v85_: _dafny.Map
            d_640_v85_ = _dafny.Map({(d_547_v0_) | (d_547_v0_): (d_639_v84_).issubset(d_639_v84_)})
            d_641_v86_: _dafny.Seq
            d_641_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afrpdrn"))
            if not(((d_640_v85_)[(d_547_v0_ if (self).f17 else d_547_v0_)] if ((d_547_v0_ if (self).f17 else d_547_v0_)) in (d_640_v85_) else (d_641_v86_) <= (d_641_v86_))):
                (globalState).f7 = p1
                (globalState).f7 = len(d_641_v86_)
                d_642_v87_: _dafny.Map
                d_642_v87_ = _dafny.Map({False: d_549_v2_})
                d_643_v88_: _dafny.Map
                d_643_v88_ = _dafny.Map({d_642_v87_: d_637_i4_})
                d_644_v89_: D14
                d_644_v89_ = D14_DC27(d_643_v88_)
                d_643_v88_ = ((d_644_v89_).cf39) | ((d_643_v88_) | (d_643_v88_))
                (globalState).f2 = d_637_i4_
                d_645_v90_: str
                d_645_v90_ = _dafny.CodePoint('e')
                d_646_v91_: _dafny.Map
                d_646_v91_ = _dafny.Map({p3: default__.fm35((d_638_v83_)[default__.safeIndex(p1, len(d_638_v83_))], p3, len(d_547_v0_), d_645_v90_, globalState)})
                d_647_v92_: _dafny.Array
                nw139_ = _dafny.Array(None, 5)
                nw139_[int(0)] = d_637_i4_
                nw139_[int(1)] = len(d_646_v91_)
                nw139_[int(2)] = (_dafny.MultiSet([d_639_v84_])).cardinality
                nw139_[int(3)] = 842
                nw139_[int(4)] = (p1) - (p1)
                d_647_v92_ = nw139_
                index103_ = default__.safeIndex(74, (d_647_v92_).length(0))
                (d_647_v92_)[index103_] = 850
                index104_ = default__.safeIndex(74, (d_647_v92_).length(0))
                (d_647_v92_)[index104_] = (0) - (p1)
            elif True:
                d_648_v93_: _dafny.Array
                nw140_ = _dafny.Array(_dafny.Map({}), 6)
                d_648_v93_ = nw140_
                d_649_v94_: _dafny.Map
                d_649_v94_ = _dafny.Map({len(d_641_v86_): (0) - (d_637_i4_)})
                index105_ = default__.safeIndex(610, (d_648_v93_).length(0))
                (d_648_v93_)[index105_] = d_649_v94_
                index106_ = default__.safeIndex(610, (d_648_v93_).length(0))
                (d_648_v93_)[index106_] = d_649_v94_
                (self).f16 = (p3) and (p3)
                d_650_v95_: _dafny.Set
                d_650_v95_ = _dafny.Set({(self).f17, p3})
                d_651_v96_: _dafny.Map
                d_651_v96_ = _dafny.Map({p3: d_650_v95_})
                d_652_v97_: _dafny.Seq
                d_652_v97_ = _dafny.SeqWithoutIsStrInference([d_650_v95_])
                d_653_v98_: _dafny.Array
                nw141_ = _dafny.Array(None, 26)
                nw141_[int(0)] = ((d_651_v96_)[self.f16] if (self.f16) in (d_651_v96_) else (d_652_v97_)[default__.safeIndex(d_637_i4_, len(d_652_v97_))])
                nw141_[int(1)] = d_650_v95_
                nw141_[int(2)] = _dafny.Set({True})
                nw141_[int(3)] = d_650_v95_
                nw141_[int(4)] = d_650_v95_
                nw141_[int(5)] = d_650_v95_
                nw141_[int(6)] = d_650_v95_
                nw141_[int(7)] = d_650_v95_
                nw141_[int(8)] = d_650_v95_
                nw141_[int(9)] = d_650_v95_
                nw141_[int(10)] = _dafny.Set({p0, p0, self.f16, default__.fm43((self).f17, globalState)})
                nw141_[int(11)] = d_650_v95_
                nw141_[int(12)] = d_650_v95_
                nw141_[int(13)] = d_650_v95_
                nw141_[int(14)] = d_650_v95_
                nw141_[int(15)] = d_650_v95_
                nw141_[int(16)] = d_650_v95_
                nw141_[int(17)] = d_650_v95_
                nw141_[int(18)] = _dafny.Set({((d_547_v0_)[p3] if (p3) in (d_547_v0_) else p0)})
                nw141_[int(19)] = _dafny.Set({p3})
                nw141_[int(20)] = _dafny.Set({(self).f17})
                nw141_[int(21)] = d_650_v95_
                nw141_[int(22)] = d_650_v95_
                nw141_[int(23)] = d_650_v95_
                nw141_[int(24)] = _dafny.Set({p0, p3})
                nw141_[int(25)] = default__.fm44(self.f16, self.f16, globalState)
                d_653_v98_ = nw141_
                d_654_v99_: C0
                nw142_ = C0()
                nw142_.ctor__(d_653_v98_, d_637_i4_)
                d_654_v99_ = nw142_
                r1 = ((_dafny.CodePoint('w')) in (d_641_v86_) if False else p3)
                d_655_v100_: _dafny.Array
                nw143_ = _dafny.Array(False, 27)
                d_655_v100_ = nw143_
                index107_ = default__.safeIndex(426, (d_655_v100_).length(0))
                (d_655_v100_)[index107_] = p0
                index108_ = default__.safeIndex(426, (d_655_v100_).length(0))
                (d_655_v100_)[index108_] = p3
        d_656_v101_: _dafny.Set
        d_656_v101_ = _dafny.Set({p1})
        def iife55_():
            coll37_ = _dafny.Set()
            compr_37_: int
            for compr_37_ in (d_656_v101_).Elements:
                d_657_v102_: int = compr_37_
                if (d_657_v102_) in (d_656_v101_):
                    coll37_ = coll37_.union(_dafny.Set([default__.safeDivisionInt(d_657_v102_, len(_dafny.SeqWithoutIsStrInference([False])))]))
            return _dafny.Set(coll37_)
        r1 = (default__.fm49(p1, p1, self.f16, p1, globalState)) == ((iife55_()
        ) | (d_656_v101_))
        d_658_v103_: _dafny.MultiSet
        d_658_v103_ = _dafny.MultiSet([p0])
        d_659_v104_: _dafny.Map
        d_659_v104_ = _dafny.Map({p3: d_658_v103_})
        r0 = (d_659_v104_) | (d_659_v104_)
        r1 = p0
        return r0, r1


class C3(T1):
    def  __init__(self):
        self._f16: bool = False
        self._f17: bool = False
        self._f22: int = int(0)
        self._f23: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
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
    def ctor__(self, f22, f23, f16, f17):
        (self)._f22 = f22
        (self)._f23 = f23
        (self).f16 = f16
        (self)._f17 = f17

    def fm3(self, globalState):
        return (_dafny.Map({len(_dafny.Map({len(_dafny.Map({self.f16: (self).f17})): len((self).f23)})): (self).f22})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f22])): 391}))

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_660_v0_: str
        d_660_v0_ = _dafny.CodePoint('w')
        def lambda39_(source7_):
            if source7_.is_DC10:
                if self.f16:
                    return (self).f17
                elif True:
                    return True
            elif True:
                d_661___mcc_h0_ = source7_.cf14
                d_662_cf14_ = d_661___mcc_h0_
                return (self).f17

        (self).f16 = lambda39_(D6_DC9(d_660_v0_))
        d_663_v1_: _dafny.Seq
        d_663_v1_ = _dafny.SeqWithoutIsStrInference([(self).f17])
        d_664_v2_: _dafny.Seq
        d_664_v2_ = d_663_v1_
        d_665_v3_: _dafny.Map
        d_665_v3_ = _dafny.Map({(self).f22: (self).f22})
        d_666_v4_: _dafny.Seq
        d_666_v4_ = _dafny.SeqWithoutIsStrInference([p0, ((d_665_v3_)[(self).f22] if ((self).f22) in (d_665_v3_) else p0), p0])
        source8_ = D4_DC7(d_664_v2_, len(_dafny.Map({d_666_v4_: self.f16})))
        if source8_.is_DC7:
            d_667___mcc_h1_ = source8_.cf11
            d_668___mcc_h2_ = source8_.cf12
            d_669_cf12_ = d_668___mcc_h2_
            d_670_cf11_ = d_667___mcc_h1_
            d_671_v5_: _dafny.Array
            nw144_ = _dafny.Array(False, 9)
            d_671_v5_ = nw144_
            d_672_v6_: _dafny.Map
            d_672_v6_ = _dafny.Map({p1: d_671_v5_})
            d_672_v6_ = d_672_v6_
            nw145_ = _dafny.Array(False, 21)
            d_671_v5_ = nw145_
            (globalState).f2 = d_669_cf12_
            rhs105_ = (self).f17
            rhs106_ = d_669_cf12_
            lhs71_ = self
            lhs71_.f16 = rhs105_
            r0 = rhs106_
        elif True:
            d_673___mcc_h3_ = source8_.cf10
            d_674_cf10_ = d_673___mcc_h3_
            d_675_v7_: _dafny.Map
            d_675_v7_ = _dafny.Map({(d_674_cf10_) | ((_dafny.Map({self.f16: (self).f17})).set((self).f17, self.f16)): (0) - (p0)})
            d_675_v7_ = (d_675_v7_).set(d_674_cf10_, p0)
            d_660_v0_ = d_660_v0_
            d_676_v8_: _dafny.Map
            d_676_v8_ = _dafny.Map({(self).f22: (self).f17})
            d_677_v9_: D3
            d_677_v9_ = D3_DC3((self).f17)
            d_678_v10_: _dafny.Array
            nw146_ = _dafny.Array(False, 3)
            d_678_v10_ = nw146_
            d_679_v11_: _dafny.Set
            d_679_v11_ = _dafny.Set({d_678_v10_, d_678_v10_})
            d_680_v12_: D3
            d_680_v12_ = D3_DC4(False, p0, p0, p1, (self).f17)
            d_681_v13_: _dafny.Array
            nw147_ = _dafny.Array(None, 28)
            nw147_[int(0)] = p1
            nw147_[int(1)] = (self).f17
            nw147_[int(2)] = not(p1)
            nw147_[int(3)] = (self).f17
            nw147_[int(4)] = ((d_676_v8_)[p0] if (p0) in (d_676_v8_) else ((d_676_v8_)[p0] if (p0) in (d_676_v8_) else ((d_676_v8_)[p0] if (p0) in (d_676_v8_) else self.f16)))
            nw147_[int(5)] = (self.f16 if p1 else self.f16)
            nw147_[int(6)] = (d_677_v9_).cf3
            nw147_[int(7)] = ((self).f22) < (536)
            nw147_[int(8)] = self.f16
            nw147_[int(9)] = not(self.f16)
            nw147_[int(10)] = True
            nw147_[int(11)] = p1
            nw147_[int(12)] = ((0) - (p0)) >= (len(_dafny.Set({not(self.f16)})))
            nw147_[int(13)] = self.f16
            nw147_[int(14)] = p1
            nw147_[int(15)] = not(p1)
            nw147_[int(16)] = False
            nw147_[int(17)] = (d_663_v1_)[default__.safeIndex(p0, len(d_663_v1_))]
            nw147_[int(18)] = p1
            nw147_[int(19)] = self.f16
            nw147_[int(20)] = p1
            nw147_[int(21)] = (d_679_v11_).isdisjoint(d_679_v11_)
            nw147_[int(22)] = not((self).f17)
            nw147_[int(23)] = self.f16
            nw147_[int(24)] = self.f16
            nw147_[int(25)] = (self).f17
            nw147_[int(26)] = True
            nw147_[int(27)] = ((d_680_v12_).cf6) < (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxvxscggb"))))
            d_681_v13_ = nw147_
            index109_ = default__.safeIndex(717, (d_681_v13_).length(0))
            (d_681_v13_)[index109_] = default__.fm31(p1, p1, globalState)
            d_682_v14_: _dafny.Set
            d_682_v14_ = _dafny.Set({p0})
            d_683_v16_: int
            d_683_v16_ = (self).f22
            d_684_v17_: _dafny.MultiSet
            d_684_v17_ = _dafny.MultiSet([_dafny.Set({(0) - ((d_683_v16_)), p0}), d_682_v14_, d_682_v14_, d_682_v14_])
            index110_ = default__.safeIndex(717, (d_681_v13_).length(0))
            def iife56_():
                coll38_ = _dafny.Set()
                compr_38_: int
                for compr_38_ in _dafny.IntegerRange(122, 68):
                    d_685_v15_: int = compr_38_
                    if ((122) <= (d_685_v15_)) and ((d_685_v15_) < (68)):
                        coll38_ = coll38_.union(_dafny.Set([default__.safeDivisionInt(d_685_v15_, p0)]))
                return _dafny.Set(coll38_)
            rhs107_ = (d_660_v0_ if not(self.f16) else d_660_v0_)
            rhs108_ = default__.fm31((True) and (True), self.f16, globalState)
            rhs109_ = default__.safeModuloInt(len(d_666_v4_), default__.fm0(p1, globalState))
            rhs110_ = _dafny.SeqWithoutIsStrInference([p0, (self).f22, p0, p0])
            rhs111_ = default__.fm31((self).f17, (d_684_v17_).issubset(_dafny.MultiSet([d_682_v14_, iife56_()
            ])), globalState)
            lhs72_ = self
            lhs73_ = globalState
            lhs74_ = d_681_v13_
            lhs75_ = default__.safeIndex(717, (d_681_v13_).length(0))
            d_660_v0_ = rhs107_
            lhs72_.f16 = rhs108_
            lhs73_.f7 = rhs109_
            d_666_v4_ = rhs110_
            lhs74_[lhs75_] = rhs111_
            d_686_v18_: _dafny.Seq
            d_686_v18_ = _dafny.SeqWithoutIsStrInference([d_681_v13_])
            d_687_v19_: _dafny.Array
            nw148_ = _dafny.Array(None, 13)
            nw148_[int(0)] = default__.safeDivisionInt((self).f22, (0) - ((self).f22))
            nw148_[int(1)] = 297
            nw148_[int(2)] = ((self).f22) - ((self).f22)
            nw148_[int(3)] = default__.fm0(self.f16, globalState)
            nw148_[int(4)] = p0
            nw148_[int(5)] = len(d_686_v18_)
            nw148_[int(6)] = (self).f22
            nw148_[int(7)] = p0
            nw148_[int(8)] = default__.fm0((self).f17, globalState)
            nw148_[int(9)] = (self).f22
            nw148_[int(10)] = p0
            nw148_[int(11)] = len(d_663_v1_)
            nw148_[int(12)] = p0
            d_687_v19_ = nw148_
            d_688_v20_: _dafny.Seq
            d_688_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maojmxu"))
            d_689_v21_: _dafny.Seq
            d_689_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: (self).f17}), _dafny.Map({p0: (d_681_v13_)[default__.safeIndex(717, (d_681_v13_).length(0))]})])
            d_690_v22_: D9
            d_690_v22_ = D9_DC15((self).f17, (d_681_v13_)[default__.safeIndex(717, (d_681_v13_).length(0))], _dafny.Set({p0}), d_687_v19_, d_681_v13_)
            rhs112_ = (d_690_v22_).cf21
            rhs113_ = (self).f23
            rhs114_ = d_689_v21_
            d_687_v19_ = rhs112_
            d_688_v20_ = rhs113_
            d_689_v21_ = rhs114_
        if not((p1) and (default__.fm31(True, p1, globalState))):
            d_691_v23_: _dafny.Set
            d_691_v23_ = _dafny.Set({(self).f22})
            d_691_v23_ = d_691_v23_
            d_692_v24_: _dafny.Map
            d_692_v24_ = _dafny.Map({default__.fm32((self).f17, not(p1), self.f16, self.f16, globalState): len((self).f23)})
            d_693_v26_: _dafny.Array
            def lambda40_(d_694_i0_):
                return (d_694_i0_) - ((self).f22)

            init23_ = lambda40_
            nw149_ = _dafny.Array(None, 12)
            for i0_23_ in range(nw149_.length(0)):
                nw149_[i0_23_] = init23_(i0_23_)
            d_693_v26_ = nw149_
            d_695_v27_: _dafny.Map
            d_695_v27_ = _dafny.Map({(self).f22: d_693_v26_})
            d_696_v28_: _dafny.Map
            d_696_v28_ = _dafny.Map({self.f16: (self).f17})
            d_697_v29_: _dafny.Array
            nw150_ = _dafny.Array(False, 1)
            d_697_v29_ = nw150_
            d_698_v30_: _dafny.Seq
            d_698_v30_ = _dafny.SeqWithoutIsStrInference([d_697_v29_, d_697_v29_])
            d_699_v31_: _dafny.Map
            d_699_v31_ = _dafny.Map({len(d_696_v28_): (d_698_v30_)[default__.safeIndex((self).f22, len(d_698_v30_))]})
            d_700_v32_: _dafny.Set
            d_700_v32_ = _dafny.Set({(self).f17, self.f16, not((self).f17)})
            d_701_v33_: D9
            def iife57_():
                coll39_ = _dafny.Set()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(679, 561):
                    d_702_v25_: int = compr_39_
                    if ((679) <= (d_702_v25_)) and ((d_702_v25_) < (561)):
                        coll39_ = coll39_.union(_dafny.Set([(d_702_v25_) - ((self).f22)]))
                return _dafny.Set(coll39_)
            d_701_v33_ = D9_DC15(((_dafny.Map({p1: (self).f22})).set((self).f17, p0)) in (d_692_v24_), (self).f17, iife57_()
, ((d_695_v27_)[(self).f22] if ((self).f22) in (d_695_v27_) else d_693_v26_), ((d_699_v31_)[(0) - (len(d_700_v32_))] if ((0) - (len(d_700_v32_))) in (d_699_v31_) else d_697_v29_))
            d_701_v33_ = D9_DC15(((self).f17) == (p1), default__.fm31(self.f16, self.f16, globalState), d_691_v23_, d_693_v26_, d_697_v29_)
            (globalState).f2 = (0) - (default__.safeModuloInt(p0, (0) - (len(_dafny.SeqWithoutIsStrInference([p1, not(self.f16)])))))
            d_703_v34_: _dafny.Array
            nw151_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_703_v34_ = nw151_
            index111_ = default__.safeIndex(660, (d_703_v34_).length(0))
            (d_703_v34_)[index111_] = d_697_v29_
            index112_ = default__.safeIndex(660, (d_703_v34_).length(0))
            (d_703_v34_)[index112_] = d_697_v29_
            d_704_v35_: _dafny.MultiSet
            d_704_v35_ = _dafny.MultiSet([(self).f22])
            d_705_v36_: _dafny.MultiSet
            d_705_v36_ = _dafny.MultiSet([d_704_v35_])
            (self).f16 = (d_705_v36_).ispropersubset(((d_705_v36_).set(d_704_v35_, default__.abs(330))).set(_dafny.MultiSet([len(d_663_v1_), len((self).f23)]), default__.abs((self).f22)))
        elif True:
            d_706_v37_: _dafny.Array
            nw152_ = _dafny.Array(False, 14)
            d_706_v37_ = nw152_
            d_707_v38_: _dafny.Array
            d_707_v38_ = d_706_v37_
            d_708_v39_: _dafny.Seq
            d_708_v39_ = _dafny.SeqWithoutIsStrInference([d_707_v38_])
            d_709_v40_: D3
            d_709_v40_ = D3_DC3(False)
            d_710_v41_: _dafny.Map
            d_710_v41_ = _dafny.Map({(len(d_708_v39_)) <= (p0): d_709_v40_})
            d_711_v42_: _dafny.Map
            d_711_v42_ = _dafny.Map({(self).f17: p1})
            d_710_v41_ = (d_710_v41_).set(((d_711_v42_)[p1] if (p1) in (d_711_v42_) else p1), d_709_v40_)
            d_712_v43_: _dafny.Array
            def lambda41_(d_713_v0_):
                def lambda42_(d_714_i1_):
                    return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D6_DC9(d_713_v0_)]))

                return lambda42_

            init24_ = lambda41_(d_660_v0_)
            nw153_ = _dafny.Array(None, 27)
            for i0_24_ in range(nw153_.length(0)):
                nw153_[i0_24_] = init24_(i0_24_)
            d_712_v43_ = nw153_
            d_715_v44_: _dafny.Map
            d_715_v44_ = _dafny.Map({(self).f17: (self).f22})
            d_716_v45_: _dafny.Map
            d_716_v45_ = _dafny.Map({d_712_v43_: d_715_v44_})
            d_717_v46_: _dafny.Seq
            d_717_v46_ = _dafny.SeqWithoutIsStrInference([d_712_v43_, d_712_v43_, d_712_v43_])
            d_716_v45_ = (d_716_v45_).set((d_717_v46_)[default__.safeIndex((self).f22, len(d_717_v46_))], default__.fm32(p1, False, self.f16, p1, globalState))
            (self).f16 = False
            d_666_v4_ = d_666_v4_
            d_718_v47_: _dafny.Seq
            d_718_v47_ = (d_666_v4_) + (d_666_v4_)
            d_718_v47_ = ((d_718_v47_ if p1 else d_718_v47_) if default__.fm31((self).f17, default__.fm31(p1, self.f16, globalState), globalState) else d_718_v47_)
        (globalState).f2 = (p0) * ((self).f22)
        if default__.fm31(p1, (self).f17, globalState):
            r0 = default__.fm0(self.f16, globalState)
            d_719_v48_: _dafny.Set
            d_719_v48_ = _dafny.Set({p1})
            (globalState).f7 = (len(d_719_v48_) if p1 else p0)
            (self).f16 = not ((self).f17) or (((self).f22) == (775))
            d_720_v49_: D4
            d_720_v49_ = D4_DC7(d_664_v2_, (self).f22)
            d_721_v50_: _dafny.Array
            nw154_ = _dafny.Array(int(0), 28)
            d_721_v50_ = nw154_
            d_722_v51_: _dafny.Map
            d_722_v51_ = _dafny.Map({d_721_v50_: p0})
            d_723_v52_: _dafny.MultiSet
            d_723_v52_ = _dafny.MultiSet([(self).f22])
            d_724_v53_: _dafny.Map
            d_724_v53_ = _dafny.Map({d_719_v48_: d_723_v52_})
            d_725_v54_: D9
            d_725_v54_ = D9_DC16((self).f22, p1)
            d_726_v55_: _dafny.Array
            nw155_ = _dafny.Array(None, 29)
            nw155_[int(0)] = (0) - (((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdsflqmgt"))))) - (len(d_719_v48_)))
            nw155_[int(1)] = (D9_DC16((self).f22, (self).f17)).cf23
            nw155_[int(2)] = (self).f22
            nw155_[int(3)] = (self).f22
            nw155_[int(4)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsh"))) + ((self).f23))
            nw155_[int(5)] = p0
            nw155_[int(6)] = -190
            nw155_[int(7)] = 20
            nw155_[int(8)] = (self).f22
            nw155_[int(9)] = (p0) * ((0) - (p0))
            nw155_[int(10)] = (self).f22
            nw155_[int(11)] = p0
            nw155_[int(12)] = len(_dafny.Map({d_660_v0_: p1}))
            nw155_[int(13)] = p0
            nw155_[int(14)] = p0
            nw155_[int(15)] = p0
            nw155_[int(16)] = (self).f22
            nw155_[int(17)] = p0
            nw155_[int(18)] = (817 if self.f16 else default__.fm0(p1, globalState))
            nw155_[int(19)] = (self).f22
            nw155_[int(20)] = (self).f22
            nw155_[int(21)] = p0
            nw155_[int(22)] = (self).f22
            nw155_[int(23)] = (self).f22
            nw155_[int(24)] = len(d_666_v4_)
            nw155_[int(25)] = ((d_722_v51_)[d_721_v50_] if (d_721_v50_) in (d_722_v51_) else p0)
            nw155_[int(26)] = default__.safeDivisionInt((d_723_v52_).cardinality, (self).f22)
            nw155_[int(27)] = (((d_724_v53_)[d_719_v48_] if (d_719_v48_) in (d_724_v53_) else (_dafny.MultiSet([p0])).set((d_725_v54_).cf23, default__.abs((self).f22)))).cardinality
            nw155_[int(28)] = default__.safeModuloInt((self).f22, p0)
            d_726_v55_ = nw155_
            index113_ = default__.safeIndex(88, (d_726_v55_).length(0))
            (d_726_v55_)[index113_] = ((self).f22) * (default__.fm0(p1, globalState))
            index114_ = default__.safeIndex(88, (d_726_v55_).length(0))
            rhs115_ = d_720_v49_
            rhs116_ = (self).f22
            lhs76_ = d_726_v55_
            lhs77_ = default__.safeIndex(88, (d_726_v55_).length(0))
            d_720_v49_ = rhs115_
            lhs76_[lhs77_] = rhs116_
            if self.f16:
                (globalState).f7 = p0
                d_727_v56_: C1
                nw156_ = C1()
                nw156_.ctor__(p1, len((self).f23))
                d_727_v56_ = nw156_
                d_728_v57_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.Set({}), 17)
                d_728_v57_ = nw157_
                d_729_v58_: C0
                nw158_ = C0()
                nw158_.ctor__(d_728_v57_, (self).f22)
                d_729_v58_ = nw158_
                (d_727_v56_).f25 = self.f16
                d_730_v59_: _dafny.Array
                def lambda43_(d_731_i2_):
                    return True

                init25_ = lambda43_
                nw159_ = _dafny.Array(None, 12)
                for i0_25_ in range(nw159_.length(0)):
                    nw159_[i0_25_] = init25_(i0_25_)
                d_730_v59_ = nw159_
                index115_ = default__.safeIndex(643, (d_730_v59_).length(0))
                (d_730_v59_)[index115_] = False
                index116_ = default__.safeIndex(643, (d_730_v59_).length(0))
                (d_730_v59_)[index116_] = d_727_v56_.f25
            elif True:
                d_732_v60_: _dafny.Array
                nw160_ = _dafny.Array(False, 1)
                d_732_v60_ = nw160_
                index117_ = default__.safeIndex(264, (d_732_v60_).length(0))
                (d_732_v60_)[index117_] = self.f16
                d_733_v61_: _dafny.MultiSet
                d_733_v61_ = _dafny.MultiSet([d_732_v60_, d_732_v60_, d_732_v60_, d_732_v60_])
                index118_ = default__.safeIndex(264, (d_732_v60_).length(0))
                rhs117_ = _dafny.CodePoint('m')
                rhs118_ = (self).f17
                rhs119_ = ((d_733_v61_) - (d_733_v61_)) != ((_dafny.MultiSet([d_732_v60_, d_732_v60_])) - (d_733_v61_))
                lhs78_ = self
                lhs79_ = d_732_v60_
                lhs80_ = default__.safeIndex(264, (d_732_v60_).length(0))
                d_660_v0_ = rhs117_
                lhs78_.f16 = rhs118_
                lhs79_[lhs80_] = rhs119_
                d_734_v62_: _dafny.MultiSet
                d_734_v62_ = _dafny.MultiSet([d_660_v0_])
                d_735_v63_: _dafny.MultiSet
                d_735_v63_ = _dafny.MultiSet([self.f16, (self).f17])
                index119_ = default__.safeIndex(264, (d_732_v60_).length(0))
                index120_ = default__.safeIndex(88, (d_726_v55_).length(0))
                rhs120_ = (0) - ((0) - ((d_734_v62_).cardinality))
                rhs121_ = p0
                rhs122_ = (d_735_v63_) == (d_735_v63_)
                rhs123_ = p0
                rhs124_ = d_721_v50_
                lhs81_ = globalState
                lhs82_ = d_732_v60_
                lhs83_ = default__.safeIndex(264, (d_732_v60_).length(0))
                lhs84_ = d_726_v55_
                lhs85_ = default__.safeIndex(88, (d_726_v55_).length(0))
                lhs81_.f7 = rhs120_
                r0 = rhs121_
                lhs82_[lhs83_] = rhs122_
                lhs84_[lhs85_] = rhs123_
                d_721_v50_ = rhs124_
                (globalState).f7 = (self).f22
                index121_ = default__.safeIndex(264, (d_732_v60_).length(0))
                (d_732_v60_)[index121_] = not((self).f17)
                d_736_v64_: D3
                d_736_v64_ = D3_DC3(p1)
                d_737_v65_: C2
                nw161_ = C2()
                nw161_.ctor__(d_736_v64_, not(p1), (d_735_v63_) == (d_735_v63_))
                d_737_v65_ = nw161_
        elif True:
            r0 = (p0 if p1 else (0) - ((self).f22))
            d_738_v66_: _dafny.Array
            def lambda44_(d_739_p0_):
                def lambda45_(d_740_i3_):
                    return (d_740_i3_) + (d_739_p0_)

                return lambda45_

            init26_ = lambda44_(p0)
            nw162_ = _dafny.Array(None, 29)
            for i0_26_ in range(nw162_.length(0)):
                nw162_[i0_26_] = init26_(i0_26_)
            d_738_v66_ = nw162_
            d_741_v67_: _dafny.Map
            d_741_v67_ = _dafny.Map({p1: d_738_v66_})
            d_742_v68_: _dafny.Set
            d_742_v68_ = _dafny.Set({self.f16, self.f16, (self).f17, False})
            d_743_v69_: _dafny.Map
            d_743_v69_ = _dafny.Map({((d_741_v67_)[self.f16] if (self.f16) in (d_741_v67_) else d_738_v66_): d_742_v68_})
            d_744_v70_: _dafny.Map
            d_744_v70_ = _dafny.Map({d_666_v4_: len(d_743_v69_)})
            d_745_v71_: _dafny.MultiSet
            d_745_v71_ = _dafny.MultiSet([p1, (self).f17])
            d_746_v72_: _dafny.Map
            d_746_v72_ = _dafny.Map({self.f16: d_745_v71_})
            rhs125_ = d_666_v4_
            rhs126_ = ((d_744_v70_)[_dafny.SeqWithoutIsStrInference([(self).f22, p0, len(d_746_v72_), 633])] if (_dafny.SeqWithoutIsStrInference([(self).f22, p0, len(d_746_v72_), 633])) in (d_744_v70_) else (self).f22)
            d_666_v4_ = rhs125_
            r0 = rhs126_
            d_747_v73_: _dafny.Set
            d_747_v73_ = _dafny.Set({d_745_v71_, d_745_v71_, d_745_v71_, d_745_v71_})
            (self).m11(default__.safeModuloInt(p0, 258), (self).f17, (d_747_v73_) | (default__.fm50((self).f22, p1, p1, globalState)), d_745_v71_, globalState)
            d_748_v74_: C1
            nw163_ = C1()
            nw163_.ctor__(p1, default__.safeDivisionInt(502, p0))
            d_748_v74_ = nw163_
            d_749_v75_: _dafny.Map
            d_749_v75_ = _dafny.Map({self.f16: (self).f17})
            d_749_v75_ = (d_749_v75_).set(False, p1)
        d_750_v76_: _dafny.MultiSet
        d_750_v76_ = _dafny.MultiSet([(self).f17, p1])
        d_751_v77_: D12
        d_751_v77_ = D12_DC23(self.f16, (self).f22, (self).f22)
        if not(((d_750_v76_).cardinality) > ((d_751_v77_).cf31)):
            d_752_v78_: _dafny.Map
            d_752_v78_ = _dafny.Map({not ((self).f17) or ((self).f17): d_666_v4_})
            d_752_v78_ = (d_752_v78_).set(not(not(p1)), d_666_v4_)
            d_753_v79_: D6
            d_753_v79_ = D6_DC10()
            d_754_v80_: _dafny.Map
            d_754_v80_ = _dafny.Map({d_663_v1_: d_753_v79_})
            def iife58_():
                coll40_ = _dafny.Map()
                compr_40_: _dafny.Seq
                for compr_40_ in (_dafny.Map({d_663_v1_: len(_dafny.SeqWithoutIsStrInference([p1, self.f16, p1, (self).f17, (self).f17]))})).keys.Elements:
                    d_755_v81_: _dafny.Seq = compr_40_
                    if (d_755_v81_) in (_dafny.Map({d_663_v1_: len(_dafny.SeqWithoutIsStrInference([p1, self.f16, p1, (self).f17, (self).f17]))})):
                        coll40_[d_755_v81_] = d_753_v79_
                return _dafny.Map(coll40_)
            d_754_v80_ = (iife58_()
            ) | ((d_754_v80_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f17, self.f16]): d_753_v79_})))
            d_756_v82_: D13
            d_756_v82_ = D13_DC26()
            source9_ = d_756_v82_
            if source9_.is_DC26:
                d_757_v83_: _dafny.Set
                d_757_v83_ = _dafny.Set({(self).f22, p0})
                d_758_v84_: _dafny.Map
                d_758_v84_ = _dafny.Map({p0: d_757_v83_})
                d_759_v85_: _dafny.Map
                d_759_v85_ = _dafny.Map({len((d_663_v1_) + (_dafny.SeqWithoutIsStrInference([not(self.f16)]))): d_758_v84_})
                def iife59_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(462, 717):
                        d_760_v86_: int = compr_41_
                        if ((462) <= (d_760_v86_)) and ((d_760_v86_) < (717)):
                            coll41_[(d_760_v86_) * ((self).f22)] = d_757_v83_
                    return _dafny.Map(coll41_)
                d_759_v85_ = (d_759_v85_).set(p0, iife59_()
                )
                d_761_v87_: _dafny.Set
                d_761_v87_ = _dafny.Set({d_750_v76_})
                (self).m11(p0, p1, (d_761_v87_) - (d_761_v87_), d_750_v76_, globalState)
                d_762_v88_: _dafny.Seq
                d_762_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ampyus"))
                d_763_v89_: _dafny.Map
                d_763_v89_ = _dafny.Map({d_660_v0_: (self).f17})
                d_764_v90_: _dafny.Array
                nw164_ = _dafny.Array(None, 26)
                nw164_[int(0)] = (self).f22
                nw164_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))
                nw164_[int(2)] = default__.safeModuloInt((self).f22, (self).f22)
                nw164_[int(3)] = (0) - (p0)
                nw164_[int(4)] = p0
                nw164_[int(5)] = p0
                nw164_[int(6)] = p0
                nw164_[int(7)] = (self).f22
                nw164_[int(8)] = (d_666_v4_)[default__.safeIndex((self).f22, len(d_666_v4_))]
                nw164_[int(9)] = (0) - ((self).f22)
                nw164_[int(10)] = (p0) - (p0)
                nw164_[int(11)] = 671
                nw164_[int(12)] = len(d_763_v89_)
                nw164_[int(13)] = p0
                nw164_[int(14)] = p0
                nw164_[int(15)] = p0
                nw164_[int(16)] = (default__.fm0(p1, globalState) if True else default__.fm0(True, globalState))
                nw164_[int(17)] = len(d_666_v4_)
                nw164_[int(18)] = ((0) - (p0)) + ((0) - (p0))
                nw164_[int(19)] = p0
                nw164_[int(20)] = 26
                nw164_[int(21)] = p0
                nw164_[int(22)] = p0
                nw164_[int(23)] = len(default__.fm42(self.f16, p0, p0, True, globalState))
                nw164_[int(24)] = (self).f22
                nw164_[int(25)] = (self).f22
                d_764_v90_ = nw164_
                index122_ = default__.safeIndex(322, (d_764_v90_).length(0))
                (d_764_v90_)[index122_] = (self).f22
                d_765_v91_: _dafny.Map
                d_765_v91_ = _dafny.Map({default__.fm35(p0, p1, (self).f22, d_660_v0_, globalState): (self).f22})
                d_766_v92_: D12
                d_766_v92_ = D12_DC22(d_762_v88_)
                index123_ = default__.safeIndex(322, (d_764_v90_).length(0))
                rhs127_ = (d_762_v88_) + ((self).f23)
                rhs128_ = (self).f17
                rhs129_ = default__.safeModuloInt(default__.safeModuloInt((self).f22, (0) - ((self).f22)), len((d_765_v91_) | (d_765_v91_)))
                rhs130_ = p0
                rhs131_ = (d_660_v0_) in ((d_766_v92_).cf29)
                lhs86_ = self
                lhs87_ = d_764_v90_
                lhs88_ = default__.safeIndex(322, (d_764_v90_).length(0))
                lhs89_ = self
                d_762_v88_ = rhs127_
                lhs86_.f16 = rhs128_
                lhs87_[lhs88_] = rhs129_
                r0 = rhs130_
                lhs89_.f16 = rhs131_
                r0 = (self).f22
            elif True:
                d_767___mcc_h4_ = source9_.cf38
                d_768_cf38_ = d_767___mcc_h4_
                d_769_v93_: D3
                d_769_v93_ = D3_DC3((self).f17)
                d_770_v94_: C2
                nw165_ = C2()
                nw165_.ctor__(d_769_v93_, (self).f17, p1)
                d_770_v94_ = nw165_
                (self).f16 = p1
                (globalState).f2 = p0
                d_771_v95_: _dafny.Map
                d_771_v95_ = _dafny.Map({False: (self).f22})
                d_772_v96_: _dafny.Seq
                d_772_v96_ = _dafny.SeqWithoutIsStrInference([d_771_v95_, (d_771_v95_) | (d_771_v95_), d_771_v95_, d_771_v95_])
                r0 = len((d_772_v96_)[default__.safeIndex((self).f22, len(d_772_v96_))])
            if p1:
                (globalState).f7 = (d_666_v4_)[default__.safeIndex((self).f22, len(d_666_v4_))]
                d_773_v97_: _dafny.Map
                d_773_v97_ = _dafny.Map({not(((self).f22) == ((self).f22)): d_752_v78_})
                d_773_v97_ = (d_773_v97_).set(True, d_752_v78_)
                (self).f16 = p1
                d_774_v98_: _dafny.Array
                nw166_ = _dafny.Array(int(0), 26)
                d_774_v98_ = nw166_
                index124_ = default__.safeIndex(403, (d_774_v98_).length(0))
                (d_774_v98_)[index124_] = (0) - ((self).f22)
                index125_ = default__.safeIndex(403, (d_774_v98_).length(0))
                (d_774_v98_)[index125_] = (self).f22
                (globalState).f7 = ((d_774_v98_)[default__.safeIndex(403, (d_774_v98_).length(0))] if self.f16 else p0)
            elif True:
                d_775_v99_: _dafny.Set
                d_775_v99_ = _dafny.Set({d_750_v76_})
                (self).m11((self).f22, (self).f17, d_775_v99_, d_750_v76_, globalState)
                (self).f16 = p1
                d_776_v100_: D7
                d_776_v100_ = D7_DC12()
                d_777_v101_: _dafny.Map
                d_777_v101_ = _dafny.Map({not((len(d_666_v4_)) == (402)): d_776_v100_})
                d_777_v101_ = (d_777_v101_).set(not ((d_663_v1_)[default__.safeIndex((0) - (p0), len(d_663_v1_))]) or (False), D7_DC12())
                d_778_v102_: C1
                nw167_ = C1()
                nw167_.ctor__((self).f17, ((self).f22) + ((self).f22))
                d_778_v102_ = nw167_
                d_779_v103_: D3
                d_779_v103_ = D3_DC3(not(True))
                d_780_v104_: C2
                nw168_ = C2()
                nw168_.ctor__(d_779_v103_, (d_778_v102_.f25) or (not((self).f17)), d_778_v102_.f25)
                d_780_v104_ = nw168_
            d_781_v105_: _dafny.Map
            d_781_v105_ = _dafny.Map({((self).f23)[default__.safeIndex(default__.fm0((self).f17, globalState), len((self).f23))]: 536})
            d_782_v106_: _dafny.Map
            d_782_v106_ = _dafny.Map({self.f16: (self).f22})
            d_783_v107_: _dafny.MultiSet
            d_783_v107_ = _dafny.MultiSet([34, p0, len(d_782_v106_), len((self).fm3(globalState))])
            d_781_v105_ = (d_781_v105_) | (_dafny.Map({((self).f23)[default__.safeIndex(-351, len((self).f23))]: (d_783_v107_).cardinality}))
        elif True:
            d_784_v108_: D9
            d_784_v108_ = D9_DC16(457, False)
            d_785_v109_: _dafny.Array
            nw169_ = _dafny.Array(None, 20)
            nw169_[int(0)] = (self).f17
            nw169_[int(1)] = ((self).f22) in (d_666_v4_)
            nw169_[int(2)] = (self).f17
            nw169_[int(3)] = (d_663_v1_)[default__.safeIndex((self).f22, len(d_663_v1_))]
            nw169_[int(4)] = p1
            nw169_[int(5)] = p1
            nw169_[int(6)] = self.f16
            nw169_[int(7)] = (self).f17
            nw169_[int(8)] = self.f16
            nw169_[int(9)] = (d_663_v1_)[default__.safeIndex((self).f22, len(d_663_v1_))]
            nw169_[int(10)] = ((self).f23) <= ((self).f23)
            nw169_[int(11)] = self.f16
            nw169_[int(12)] = not(((self).f22) > (p0))
            nw169_[int(13)] = True
            nw169_[int(14)] = p1
            nw169_[int(15)] = self.f16
            nw169_[int(16)] = (d_784_v108_).cf24
            nw169_[int(17)] = True
            nw169_[int(18)] = p1
            nw169_[int(19)] = (p1 if self.f16 else (self).f17)
            d_785_v109_ = nw169_
            d_786_v110_: _dafny.Seq
            d_786_v110_ = _dafny.SeqWithoutIsStrInference([d_663_v1_])
            index126_ = default__.safeIndex(605, (d_785_v109_).length(0))
            (d_785_v109_)[index126_] = (_dafny.SeqWithoutIsStrInference([d_663_v1_, d_663_v1_])) <= (d_786_v110_)
            index127_ = default__.safeIndex(605, (d_785_v109_).length(0))
            (d_785_v109_)[index127_] = not(self.f16)
            d_787_v111_: _dafny.Set
            d_787_v111_ = _dafny.Set({(self).f17})
            d_788_v112_: _dafny.Seq
            d_788_v112_ = _dafny.SeqWithoutIsStrInference([d_750_v76_])
            d_789_v113_: _dafny.Array
            nw170_ = _dafny.Array(None, 3)
            nw170_[int(0)] = (self).f22
            nw170_[int(1)] = p0
            nw170_[int(2)] = p0
            d_789_v113_ = nw170_
            d_790_v114_: D9
            d_790_v114_ = D9_DC14(d_789_v113_)
            d_791_v115_: _dafny.Seq
            d_791_v115_ = _dafny.SeqWithoutIsStrInference([d_790_v114_, d_790_v114_, d_790_v114_])
            d_792_v116_: _dafny.Map
            d_792_v116_ = _dafny.Map({self.f16: (self).f17})
            d_793_v117_: _dafny.Set
            d_793_v117_ = _dafny.Set({d_792_v116_})
            d_794_v118_: _dafny.Array
            nw171_ = _dafny.Array(None, 23)
            nw171_[int(0)] = len(((self).f23).set(default__.safeIndex(len(d_787_v111_), len((self).f23)), d_660_v0_))
            nw171_[int(1)] = (self).f22
            nw171_[int(2)] = p0
            nw171_[int(3)] = len((self).f23)
            nw171_[int(4)] = (0) - (((self).f22) * ((self).f22))
            nw171_[int(5)] = ((_dafny.MultiSet([not((d_785_v109_)[default__.safeIndex(605, (d_785_v109_).length(0))]), p1])).intersection((d_788_v112_)[default__.safeIndex(450, len(d_788_v112_))])).cardinality
            nw171_[int(6)] = (p0) + ((self).f22)
            nw171_[int(7)] = (self).f22
            nw171_[int(8)] = default__.safeDivisionInt((self).f22, (self).f22)
            nw171_[int(9)] = p0
            nw171_[int(10)] = len(d_791_v115_)
            nw171_[int(11)] = (self).f22
            nw171_[int(12)] = (self).f22
            nw171_[int(13)] = p0
            nw171_[int(14)] = (0) - ((self).f22)
            nw171_[int(15)] = (-897) - (p0)
            nw171_[int(16)] = p0
            nw171_[int(17)] = len((self).f23)
            nw171_[int(18)] = (self).f22
            nw171_[int(19)] = len(d_793_v117_)
            nw171_[int(20)] = p0
            nw171_[int(21)] = (self).f22
            nw171_[int(22)] = (_dafny.MultiSet((d_666_v4_) + ((d_666_v4_).set(default__.safeIndex(p0, len(d_666_v4_)), (d_750_v76_).cardinality)))).cardinality
            d_794_v118_ = nw171_
            nw172_ = _dafny.Array(int(0), 5)
            d_794_v118_ = nw172_
            index128_ = default__.safeIndex(761, (d_789_v113_).length(0))
            (d_789_v113_)[index128_] = p0
            index129_ = default__.safeIndex(761, (d_789_v113_).length(0))
            (d_789_v113_)[index129_] = (self).f22
            if default__.fm31((d_785_v109_)[default__.safeIndex(605, (d_785_v109_).length(0))], (d_750_v76_) == (_dafny.MultiSet(d_663_v1_)), globalState):
                index130_ = default__.safeIndex(605, (d_785_v109_).length(0))
                (d_785_v109_)[index130_] = ((self).f22) != (len(d_666_v4_))
                d_795_v119_: C1
                nw173_ = C1()
                nw173_.ctor__(((d_789_v113_)[default__.safeIndex(761, (d_789_v113_).length(0))]) < (len((self).f23)), ((self).f22) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slvmstdm")))))
                d_795_v119_ = nw173_
                (globalState).f7 = p0
                d_796_v120_: _dafny.Array
                nw174_ = _dafny.Array(_dafny.Map({}), 19)
                d_796_v120_ = nw174_
                d_797_v121_: _dafny.Set
                d_797_v121_ = _dafny.Set({len(d_787_v111_)})
                d_798_v122_: _dafny.Map
                d_798_v122_ = _dafny.Map({d_797_v121_: (self).f17})
                index131_ = default__.safeIndex(437, (d_796_v120_).length(0))
                (d_796_v120_)[index131_] = d_798_v122_
                d_799_v123_: _dafny.Seq
                d_799_v123_ = _dafny.SeqWithoutIsStrInference([d_787_v111_])
                index132_ = default__.safeIndex(437, (d_796_v120_).length(0))
                index133_ = default__.safeIndex(605, (d_785_v109_).length(0))
                rhs132_ = default__.fm43(p1, globalState)
                rhs133_ = default__.fm51(default__.fm0(self.f16, globalState), 631, (d_663_v1_) + (_dafny.SeqWithoutIsStrInference([self.f16])), d_799_v123_, globalState)
                rhs134_ = (p1) == (self.f16)
                lhs90_ = self
                lhs91_ = d_796_v120_
                lhs92_ = default__.safeIndex(437, (d_796_v120_).length(0))
                lhs93_ = d_785_v109_
                lhs94_ = default__.safeIndex(605, (d_785_v109_).length(0))
                lhs90_.f16 = rhs132_
                lhs91_[lhs92_] = rhs133_
                lhs93_[lhs94_] = rhs134_
                d_800_v124_: _dafny.MultiSet
                d_800_v124_ = _dafny.MultiSet([(((d_786_v110_)[default__.safeIndex(p0, len(d_786_v110_))]).set(default__.safeIndex(p0, len((d_786_v110_)[default__.safeIndex(p0, len(d_786_v110_))])), (self).f17)) + (d_663_v1_), d_663_v1_])
                d_801_v125_: _dafny.Seq
                d_801_v125_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axlauykx")), (self).f23, (self).f23])
                index134_ = default__.safeIndex(605, (d_785_v109_).length(0))
                rhs135_ = d_666_v4_
                rhs136_ = self.f16
                rhs137_ = ((d_800_v124_)[(d_663_v1_) + ((d_663_v1_).set(default__.safeIndex(len(d_801_v125_), len(d_663_v1_)), d_795_v119_.f25))] if ((d_663_v1_) + ((d_663_v1_).set(default__.safeIndex(len(d_801_v125_), len(d_663_v1_)), d_795_v119_.f25))) in (d_800_v124_) else p0)
                rhs138_ = default__.fm39(globalState)
                lhs95_ = d_785_v109_
                lhs96_ = default__.safeIndex(605, (d_785_v109_).length(0))
                lhs97_ = globalState
                d_666_v4_ = rhs135_
                lhs95_[lhs96_] = rhs136_
                lhs97_.f7 = rhs137_
                d_660_v0_ = rhs138_
            elif True:
                d_802_v126_: _dafny.Seq
                d_802_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srahniy"))
                d_802_v126_ = ((d_802_v126_ if True else d_802_v126_)) + ((d_802_v126_) + ((self).f23))
                d_663_v1_ = d_663_v1_
                d_803_v127_: _dafny.Set
                d_803_v127_ = _dafny.Set({((self).f23) + (_dafny.SeqWithoutIsStrInference([d_660_v0_ for d_804_i4_ in range(default__.abs(-495))])), (self).f23, (self).f23, ((self).f23) + (d_802_v126_)})
                d_805_v128_: _dafny.Array
                def lambda46_(d_806_i5_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btxaxkf"))

                init27_ = lambda46_
                nw175_ = _dafny.Array(None, 2)
                for i0_27_ in range(nw175_.length(0)):
                    nw175_[i0_27_] = init27_(i0_27_)
                d_805_v128_ = nw175_
                index135_ = default__.safeIndex(517, (d_805_v128_).length(0))
                (d_805_v128_)[index135_] = d_802_v126_
                d_807_v129_: _dafny.MultiSet
                d_807_v129_ = _dafny.MultiSet([d_660_v0_, d_660_v0_])
                index136_ = default__.safeIndex(605, (d_785_v109_).length(0))
                index137_ = default__.safeIndex(517, (d_805_v128_).length(0))
                index138_ = default__.safeIndex(605, (d_785_v109_).length(0))
                rhs139_ = d_803_v127_
                rhs140_ = (p1 if not (not(p1)) or ((d_785_v109_)[default__.safeIndex(605, (d_785_v109_).length(0))]) else (d_807_v129_).ispropersubset(_dafny.MultiSet([_dafny.CodePoint('w'), d_660_v0_])))
                rhs141_ = (self).f23
                rhs142_ = (self).f17
                lhs98_ = d_785_v109_
                lhs99_ = default__.safeIndex(605, (d_785_v109_).length(0))
                lhs100_ = d_805_v128_
                lhs101_ = default__.safeIndex(517, (d_805_v128_).length(0))
                lhs102_ = d_785_v109_
                lhs103_ = default__.safeIndex(605, (d_785_v109_).length(0))
                d_803_v127_ = rhs139_
                lhs98_[lhs99_] = rhs140_
                lhs100_[lhs101_] = rhs141_
                lhs102_[lhs103_] = rhs142_
                index139_ = default__.safeIndex(605, (d_785_v109_).length(0))
                (d_785_v109_)[index139_] = ((p1) and ((self).f17) if (((d_665_v3_)[(d_666_v4_)[default__.safeIndex((self).f22, len(d_666_v4_))]] if ((d_666_v4_)[default__.safeIndex((self).f22, len(d_666_v4_))]) in (d_665_v3_) else 459)) > (p0) else (self).f17)
                (self).f16 = (846) < (p0)
            (globalState).f2 = (p0) - (default__.fm0(p1, globalState))
        r0 = (0) - ((0) - ((192) - (len(((self).f23) + (_dafny.SeqWithoutIsStrInference([d_660_v0_ for d_808_i6_ in range(default__.abs(863))]))))))
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_809_v0_: _dafny.Map
        d_809_v0_ = _dafny.Map({p3: self.f16})
        d_810_v1_: D7
        d_810_v1_ = D7_DC11(d_809_v0_)
        d_811_v2_: _dafny.Seq
        d_811_v2_ = _dafny.SeqWithoutIsStrInference([(self).f17, p2])
        d_812_v3_: D3
        d_812_v3_ = D3_DC4(not((self).f17), len((self).f23), len((d_810_v1_).cf15), (d_811_v2_)[default__.safeIndex(len(d_811_v2_), len(d_811_v2_))], (self).f17)
        source10_ = D3_DC5(d_812_v3_)
        if source10_.is_DC4:
            d_813___mcc_h0_ = source10_.cf4
            d_814___mcc_h1_ = source10_.cf5
            d_815___mcc_h2_ = source10_.cf6
            d_816___mcc_h3_ = source10_.cf7
            d_817___mcc_h4_ = source10_.cf8
            d_818_cf8_ = d_817___mcc_h4_
            d_819_cf7_ = d_816___mcc_h3_
            d_820_cf6_ = d_815___mcc_h2_
            d_821_cf5_ = d_814___mcc_h1_
            d_822_cf4_ = d_813___mcc_h0_
            d_819_cf7_ = d_819_cf7_
            d_823_v4_: _dafny.Map
            d_823_v4_ = _dafny.Map({False: d_820_cf6_})
            (globalState).f7 = len((default__.fm52(d_819_cf7_, len(_dafny.Map({d_818_cf8_: p1})), len(d_823_v4_), (self).f22, globalState)) + (d_811_v2_))
            d_824_v5_: _dafny.Array
            nw176_ = _dafny.Array(None, 9)
            nw176_[int(0)] = p3
            nw176_[int(1)] = d_821_cf5_
            nw176_[int(2)] = 621
            nw176_[int(3)] = p1
            nw176_[int(4)] = p1
            nw176_[int(5)] = d_820_cf6_
            nw176_[int(6)] = (self).f22
            nw176_[int(7)] = (self).f22
            nw176_[int(8)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfxolsjs")))) - (d_821_cf5_)
            d_824_v5_ = nw176_
            index140_ = default__.safeIndex(740, (d_824_v5_).length(0))
            (d_824_v5_)[index140_] = p1
            d_825_v6_: _dafny.Seq
            d_825_v6_ = _dafny.SeqWithoutIsStrInference([(0) - (len((self).f23))])
            index141_ = default__.safeIndex(740, (d_824_v5_).length(0))
            (d_824_v5_)[index141_] = (d_825_v6_)[default__.safeIndex(p1, len(d_825_v6_))]
            d_826_v7_: _dafny.Map
            d_826_v7_ = _dafny.Map({True: d_825_v6_})
            d_826_v7_ = (d_826_v7_).set((d_811_v2_)[default__.safeIndex(d_821_cf5_, len(d_811_v2_))], d_825_v6_)
        elif source10_.is_DC3:
            d_827___mcc_h5_ = source10_.cf3
            d_828_cf3_ = d_827___mcc_h5_
            d_829_v8_: _dafny.Array
            nw177_ = _dafny.Array(None, 9)
            nw177_[int(0)] = d_828_cf3_
            nw177_[int(1)] = (self).f17
            nw177_[int(2)] = (self).f17
            nw177_[int(3)] = True
            nw177_[int(4)] = (p1) == (935)
            nw177_[int(5)] = (self).f17
            nw177_[int(6)] = p2
            nw177_[int(7)] = p2
            nw177_[int(8)] = False
            d_829_v8_ = nw177_
            index142_ = default__.safeIndex(860, (d_829_v8_).length(0))
            (d_829_v8_)[index142_] = default__.fm43((self).f17, globalState)
            index143_ = default__.safeIndex(860, (d_829_v8_).length(0))
            (d_829_v8_)[index143_] = (p1) != (len((self).f23))
            d_830_v9_: C2
            nw178_ = C2()
            nw178_.ctor__(default__.fm53((self).f22, globalState), p2, not((not(p2)) in (default__.fm40((self).f17, globalState))))
            d_830_v9_ = nw178_
            d_831_v10_: str
            d_831_v10_ = _dafny.CodePoint('w')
            d_831_v10_ = d_831_v10_
            d_832_v11_: _dafny.Seq
            d_832_v11_ = _dafny.SeqWithoutIsStrInference([len(default__.fm42(d_828_cf3_, 905, (self).f22, d_828_cf3_, globalState))])
            index144_ = default__.safeIndex(860, (d_829_v8_).length(0))
            (d_829_v8_)[index144_] = (default__.safeModuloInt(len(d_811_v2_), (self).f22)) in (d_832_v11_)
        elif True:
            d_833___mcc_h6_ = source10_.cf9
            d_834_cf9_ = d_833___mcc_h6_
            d_835_v12_: _dafny.Array
            def lambda47_(d_836_p1_):
                def lambda48_(d_837_i0_):
                    return (d_836_p1_) > (-949)

                return lambda48_

            init28_ = lambda47_(p1)
            nw179_ = _dafny.Array(None, 12)
            for i0_28_ in range(nw179_.length(0)):
                nw179_[i0_28_] = init28_(i0_28_)
            d_835_v12_ = nw179_
            index145_ = default__.safeIndex(247, (d_835_v12_).length(0))
            (d_835_v12_)[index145_] = self.f16
            index146_ = default__.safeIndex(247, (d_835_v12_).length(0))
            (d_835_v12_)[index146_] = True
            d_838_v13_: _dafny.Map
            d_838_v13_ = _dafny.Map({d_811_v2_: default__.safeModuloInt(p3, p3)})
            d_838_v13_ = d_838_v13_
            r0 = default__.safeModuloInt((-582) + (p3), p3)
            index147_ = default__.safeIndex(247, (d_835_v12_).length(0))
            (d_835_v12_)[index147_] = (d_835_v12_)[default__.safeIndex(247, (d_835_v12_).length(0))]
        def lambda49_(source11_):
            d_839___mcc_h7_ = source11_
            d_840_cf0_ = d_839___mcc_h7_
            return ((self).f23) + ((self).f23)

        (globalState).f2 = (0) - (len(lambda49_(d_811_v2_)))
        d_841_v14_: _dafny.Array
        def lambda50_(d_842_i1_):
            return default__.safeModuloInt(d_842_i1_, len(_dafny.Set({(self).f22})))

        init29_ = lambda50_
        nw180_ = _dafny.Array(None, 24)
        for i0_29_ in range(nw180_.length(0)):
            nw180_[i0_29_] = init29_(i0_29_)
        d_841_v14_ = nw180_
        d_841_v14_ = d_841_v14_
        d_843_v15_: C1
        nw181_ = C1()
        nw181_.ctor__(not (self.f16) or (not(p2)), (self).f22)
        d_843_v15_ = nw181_
        d_844_v16_: _dafny.Map
        d_844_v16_ = _dafny.Map({p3: (self).f22})
        d_844_v16_ = (d_844_v16_).set(606, p3)
        r0 = (0) - (p3)
        r0 = p3
        return r0

    def m11(self, p0, p1, p2, p3, globalState):
        (self).f16 = not ((self).f17) or (default__.fm31(self.f16, self.f16, globalState))
        d_845_v0_: _dafny.Array
        nw182_ = _dafny.Array(False, 29)
        d_845_v0_ = nw182_
        index148_ = default__.safeIndex(852, (d_845_v0_).length(0))
        (d_845_v0_)[index148_] = (self).f17
        index149_ = default__.safeIndex(852, (d_845_v0_).length(0))
        (d_845_v0_)[index149_] = default__.fm31(self.f16, True, globalState)
        if not ((d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))]) or (self.f16):
            (globalState).f2 = default__.safeModuloInt((self).f22, (p0) + ((self).f22))
            index150_ = default__.safeIndex(852, (d_845_v0_).length(0))
            (d_845_v0_)[index150_] = p1
            d_846_v1_: C1
            nw183_ = C1()
            nw183_.ctor__((d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], (p0) + (p0))
            d_846_v1_ = nw183_
            d_846_v1_ = d_846_v1_
            d_847_v2_: _dafny.Map
            d_847_v2_ = _dafny.Map({False: p1})
            d_848_v3_: _dafny.Map
            d_848_v3_ = _dafny.Map({d_847_v2_: 844})
            (globalState).f7 = ((0) - (((d_848_v3_)[d_847_v2_] if (d_847_v2_) in (d_848_v3_) else (0) - (len((self).f23))))) * (default__.fm0((d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], globalState))
            d_849_v4_: _dafny.Map
            d_849_v4_ = _dafny.Map({p0: (self).f17})
            d_850_v5_: _dafny.Seq
            d_850_v5_ = _dafny.SeqWithoutIsStrInference([d_849_v4_])
            d_851_v6_: D7
            d_851_v6_ = D7_DC11(((d_850_v5_)[default__.safeIndex((self).f22, len(d_850_v5_))]) | (d_849_v4_))
            d_851_v6_ = d_851_v6_
        elif True:
            rhs143_ = self.f16
            rhs144_ = (0) - (p0)
            lhs104_ = self
            lhs105_ = globalState
            lhs104_.f16 = rhs143_
            lhs105_.f7 = rhs144_
            index151_ = default__.safeIndex(852, (d_845_v0_).length(0))
            (d_845_v0_)[index151_] = (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))]
            d_852_v7_: _dafny.Map
            d_852_v7_ = _dafny.Map({(self).f17: _dafny.Map({p0: True})})
            d_853_v8_: _dafny.Map
            d_853_v8_ = _dafny.Map({(self).f22: (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))]})
            pat_let_tv7_ = d_853_v8_
            def iife60_(_pat_let9_0):
                def iife61_(d_854_dt__update__tmp_h0_):
                    def iife62_(_pat_let10_0):
                        def iife63_(d_855_dt__update_hcf15_h0_):
                            return D7_DC11(d_855_dt__update_hcf15_h0_)
                        return iife63_(_pat_let10_0)
                    return iife62_(pat_let_tv7_)
                return iife61_(_pat_let9_0)
            source12_ = iife60_(D7_DC11(((d_852_v7_)[p1] if (p1) in (d_852_v7_) else d_853_v8_)))
            if source12_.is_DC12:
                (globalState).f7 = (self).f22
                d_856_v9_: str
                d_856_v9_ = _dafny.CodePoint('h')
                (self).f16 = ((self).f22) >= (len(((self).f23 if p1 else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_857_i0_ in range(default__.abs(719))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_858_i0_ in range(default__.abs(719))]))), d_856_v9_))))
                index152_ = default__.safeIndex(852, (d_845_v0_).length(0))
                (d_845_v0_)[index152_] = (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))]
                d_859_v10_: _dafny.Array
                def lambda51_(d_860_p1_):
                    def lambda52_(d_861_i1_):
                        return _dafny.Set({d_860_p1_, self.f16})

                    return lambda52_

                init30_ = lambda51_(p1)
                nw184_ = _dafny.Array(None, 24)
                for i0_30_ in range(nw184_.length(0)):
                    nw184_[i0_30_] = init30_(i0_30_)
                d_859_v10_ = nw184_
                d_862_v11_: C0
                nw185_ = C0()
                nw185_.ctor__(d_859_v10_, (self).f22)
                d_862_v11_ = nw185_
            elif True:
                d_863___mcc_h0_ = source12_.cf15
                d_864_cf15_ = d_863___mcc_h0_
                d_865_v12_: _dafny.Seq
                d_865_v12_ = _dafny.SeqWithoutIsStrInference([(self).f17, (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], (self).f17, (True if (self).f17 else False), (p1) or (self.f16)])
                d_866_v13_: _dafny.Seq
                d_866_v13_ = _dafny.SeqWithoutIsStrInference([344, (self).f22])
                d_865_v12_ = default__.fm52((len(_dafny.SeqWithoutIsStrInference([((self).f23)[default__.safeIndex(p0, len((self).f23))] for d_867_i2_ in range(default__.abs(734))]))) < (p0), (self).f22, (len(d_866_v13_) if (self).f17 else (self).f22), default__.fm0(default__.fm31((d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], globalState), globalState), globalState)
                d_868_v14_: _dafny.Array
                def lambda53_(d_869_v12_):
                    def lambda54_(d_870_i3_):
                        return d_869_v12_

                    return lambda54_

                init31_ = lambda53_(d_865_v12_)
                nw186_ = _dafny.Array(None, 15)
                for i0_31_ in range(nw186_.length(0)):
                    nw186_[i0_31_] = init31_(i0_31_)
                d_868_v14_ = nw186_
                index153_ = default__.safeIndex(482, (d_868_v14_).length(0))
                (d_868_v14_)[index153_] = d_865_v12_
                index154_ = default__.safeIndex(482, (d_868_v14_).length(0))
                (d_868_v14_)[index154_] = (d_865_v12_) + (d_865_v12_)
                (self).f16 = (self).f17
                d_871_v15_: _dafny.Set
                d_871_v15_ = _dafny.Set({(d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], False})
                d_872_v16_: D15
                d_872_v16_ = D15_DC29(d_871_v15_)
                (globalState).f7 = len((d_871_v15_) - (((d_872_v16_).cf43) - (d_871_v15_)))
            (globalState).f7 = ((self).f22) - ((self).f22)
            d_873_v17_: _dafny.Set
            d_873_v17_ = _dafny.Set({(self).f17})
            d_874_v18_: D15
            d_874_v18_ = D15_DC29(d_873_v17_)
            d_875_v19_: D15
            d_875_v19_ = D15_DC31(d_874_v18_)
            d_875_v19_ = D15_DC31(d_874_v18_)
        d_876_v20_: _dafny.Map
        d_876_v20_ = _dafny.Map({D13_DC26(): default__.safeDivisionInt((0) - ((self).f22), p0)})
        d_876_v20_ = (d_876_v20_) | (d_876_v20_)
        index155_ = default__.safeIndex(852, (d_845_v0_).length(0))
        (d_845_v0_)[index155_] = (p0) < ((self).f22)
        pat_let_tv8_ = p1
        pat_let_tv9_ = d_845_v0_
        pat_let_tv10_ = d_845_v0_
        index156_ = default__.safeIndex(852, (d_845_v0_).length(0))
        index157_ = default__.safeIndex(852, (d_845_v0_).length(0))
        index158_ = default__.safeIndex(852, (d_845_v0_).length(0))
        def lambda55_(source13_):
            if source13_.is_DC30:
                d_877___mcc_h1_ = source13_.cf44
                d_878_cf44_ = d_877___mcc_h1_
                return pat_let_tv8_
            elif source13_.is_DC29:
                d_879___mcc_h2_ = source13_.cf43
                d_880_cf43_ = d_879___mcc_h2_
                return not((pat_let_tv10_)[default__.safeIndex(852, (pat_let_tv9_).length(0))])
            elif True:
                d_881___mcc_h3_ = source13_.cf45
                d_882_cf45_ = d_881___mcc_h3_
                return self.f16

        rhs145_ = lambda55_(D15_DC29(default__.fm44((self).f17, (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))], globalState)))
        rhs146_ = (0) - ((0) - (((0) - ((p0) - (p0)) if (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))] else (p0) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_883_i4_ in range(default__.abs(-710))]))))))
        rhs147_ = (d_845_v0_)[default__.safeIndex(852, (d_845_v0_).length(0))]
        rhs148_ = default__.fm0(p1, globalState)
        rhs149_ = True
        lhs106_ = d_845_v0_
        lhs107_ = default__.safeIndex(852, (d_845_v0_).length(0))
        lhs108_ = globalState
        lhs109_ = d_845_v0_
        lhs110_ = default__.safeIndex(852, (d_845_v0_).length(0))
        lhs111_ = globalState
        lhs112_ = d_845_v0_
        lhs113_ = default__.safeIndex(852, (d_845_v0_).length(0))
        lhs106_[lhs107_] = rhs145_
        lhs108_.f2 = rhs146_
        lhs109_[lhs110_] = rhs147_
        lhs111_.f2 = rhs148_
        lhs112_[lhs113_] = rhs149_

    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23

class C4(T2, T1):
    def  __init__(self):
        self._f16: bool = False
        self._f17: bool = False
        self._f21: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f21, f16, f17):
        (self)._f21 = f21
        (self).f16 = f16
        (self)._f17 = f17

    def fm12(self, p0, globalState):
        return (0) - ((0) - ((0) - ((0) - (default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsrmhd")) if (self).f17 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv")))), len((_dafny.Set({self.f16, False})) - (_dafny.Set({not(self.f16)}))))))))

    def fm13(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(self).f21, (self).f17])) + (_dafny.SeqWithoutIsStrInference([self.f16]))) + ((_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])) + (_dafny.SeqWithoutIsStrInference([True, not((self).f17)])))

    def fm3(self, globalState):
        def iife64_():
            coll42_ = _dafny.Map()
            compr_42_: int
            for compr_42_ in (_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teahcor")))})): _dafny.SeqWithoutIsStrInference([-549])})).keys.Elements:
                d_884_v0_: int = compr_42_
                if (d_884_v0_) in (_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teahcor")))})): _dafny.SeqWithoutIsStrInference([-549])})):
                    coll42_[(d_884_v0_) - (912)] = -8
            return _dafny.Map(coll42_)
        def iife65_():
            coll43_ = _dafny.Map()
            compr_43_: int
            for compr_43_ in (_dafny.SeqWithoutIsStrInference([(0) - (-808)])).Elements:
                d_885_v1_: int = compr_43_
                if (d_885_v1_) in (_dafny.SeqWithoutIsStrInference([(0) - (-808)])):
                    coll43_[default__.safeModuloInt(d_885_v1_, len(_dafny.SeqWithoutIsStrInference([922, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_886_i0_ in range(default__.abs(309))]))), 380])))] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f21, self.f16, self.f16, (self).f17, True]))])).cardinality
            return _dafny.Map(coll43_)
        return (iife64_()
        ) | ((iife65_()
        ) | (_dafny.Map({len(_dafny.Map({(_dafny.MultiSet([(self).f21, (self).f17, False])).cardinality: len(_dafny.Set({(self).f17, self.f16}))})): (_dafny.MultiSet([951])).cardinality})))

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: bool = False
        d_887_v0_: _dafny.Seq
        d_887_v0_ = _dafny.SeqWithoutIsStrInference([p2])
        d_888_v1_: _dafny.Map
        d_888_v1_ = _dafny.Map({(0) - (len(d_887_v0_)): p2})
        r1 = default__.fm27(p2, p2, (p2) + (len(d_888_v1_)), _dafny.Map({default__.fm0(p1, globalState): p2}), globalState)
        if True:
            r3 = (p2) < ((0) - (len(d_887_v0_)))
            d_889_v2_: _dafny.Seq
            d_889_v2_ = _dafny.SeqWithoutIsStrInference([not(p0), (self).f21, (self).f17, self.f16])
            d_890_v3_: _dafny.Seq
            d_890_v3_ = d_889_v2_
            d_891_v4_: D7
            d_891_v4_ = D7_DC12()
            d_892_v5_: _dafny.Map
            d_892_v5_ = _dafny.Map({d_890_v3_: d_891_v4_})
            if (d_890_v3_) not in (d_892_v5_):
                d_893_v6_: _dafny.Seq
                d_893_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojvcuvnba"))
                r2 = d_893_v6_
                d_894_v7_: _dafny.Array
                nw187_ = _dafny.Array(None, 7)
                nw187_[int(0)] = d_889_v2_
                nw187_[int(1)] = (d_889_v2_) + (d_889_v2_)
                nw187_[int(2)] = (d_889_v2_) + (d_889_v2_)
                nw187_[int(3)] = d_889_v2_
                nw187_[int(4)] = d_889_v2_
                nw187_[int(5)] = d_889_v2_
                nw187_[int(6)] = (d_889_v2_ if (self).f21 else d_889_v2_)
                d_894_v7_ = nw187_
                d_894_v7_ = d_894_v7_
                d_895_v8_: _dafny.Seq
                d_895_v8_ = _dafny.SeqWithoutIsStrInference([d_893_v6_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnpswr"))])
                d_896_v9_: _dafny.Seq
                d_896_v9_ = _dafny.SeqWithoutIsStrInference([(d_895_v8_)[default__.safeIndex(p2, len(d_895_v8_))]])
                d_897_v10_: _dafny.Set
                d_897_v10_ = _dafny.Set({p2})
                d_898_v11_: _dafny.MultiSet
                d_898_v11_ = _dafny.MultiSet([self.f16, (self).f17])
                d_899_v12_: _dafny.Array
                nw188_ = _dafny.Array(None, 13)
                nw188_[int(0)] = self.f16
                nw188_[int(1)] = ((d_896_v9_)[default__.safeIndex(-154, len(d_896_v9_))]) == (d_893_v6_)
                nw188_[int(2)] = (self).f21
                nw188_[int(3)] = (self).f17
                nw188_[int(4)] = (self).f21
                nw188_[int(5)] = (p2) >= (p2)
                nw188_[int(6)] = p1
                nw188_[int(7)] = (p0 if not(True) else (self).f17)
                nw188_[int(8)] = p0
                nw188_[int(9)] = default__.fm27(254, (self).fm12(len(d_897_v10_), globalState), p2, _dafny.Map({p2: p2}), globalState)
                nw188_[int(10)] = p1
                nw188_[int(11)] = not(self.f16)
                nw188_[int(12)] = (d_889_v2_)[default__.safeIndex((d_898_v11_).cardinality, len(d_889_v2_))]
                d_899_v12_ = nw188_
                index159_ = default__.safeIndex(298, (d_899_v12_).length(0))
                (d_899_v12_)[index159_] = p1
                d_900_v13_: _dafny.Array
                nw189_ = _dafny.Array(int(0), 12)
                d_900_v13_ = nw189_
                d_901_v14_: _dafny.Set
                d_901_v14_ = _dafny.Set({p1})
                d_902_v15_: D3
                d_902_v15_ = D3_DC4(False, 42, p2, (self).f21, True)
                d_903_v16_: _dafny.Map
                d_903_v16_ = _dafny.Map({(self).f21: p1})
                index160_ = default__.safeIndex(298, (d_899_v12_).length(0))
                rhs150_ = d_899_v12_
                rhs151_ = ((d_901_v14_) | (d_901_v14_)).ispropersubset(_dafny.Set({False, not(p0)}))
                rhs152_ = default__.safeDivisionInt(p2, 256)
                rhs153_ = default__.fm28(False, _dafny.Map({p2: len(_dafny.Set({(0) - (p2)}))}), d_902_v15_, ((d_903_v16_).set(not(True), not((self).f21))) | (d_903_v16_), globalState)
                rhs154_ = d_900_v13_
                lhs114_ = d_899_v12_
                lhs115_ = default__.safeIndex(298, (d_899_v12_).length(0))
                lhs116_ = globalState
                d_899_v12_ = rhs150_
                lhs114_[lhs115_] = rhs151_
                lhs116_.f2 = rhs152_
                d_893_v6_ = rhs153_
                d_900_v13_ = rhs154_
                (globalState).f7 = p2
                d_904_v17_: D3
                d_904_v17_ = D3_DC3((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgncwsm"))) == (d_893_v6_))
                rhs155_ = (default__.fm29(globalState) if not (p0) or (p0) else d_904_v17_)
                rhs156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebjbv"))
                d_904_v17_ = rhs155_
                d_893_v6_ = rhs156_
            elif True:
                d_905_v18_: _dafny.Map
                d_905_v18_ = _dafny.Map({p0: self.f16})
                d_906_v19_: _dafny.Set
                d_906_v19_ = _dafny.Set({p2})
                d_907_v20_: _dafny.MultiSet
                d_907_v20_ = _dafny.MultiSet([len(default__.fm28((self).f21, d_888_v1_, D3_DC4(p1, p2, p2, (self).f17, (self).f17), d_905_v18_, globalState)), len(d_906_v19_), p2])
                d_908_v21_: _dafny.Array
                nw190_ = _dafny.Array(False, 4)
                d_908_v21_ = nw190_
                d_909_v22_: _dafny.MultiSet
                d_909_v22_ = _dafny.MultiSet([d_908_v21_, d_908_v21_])
                d_910_v23_: _dafny.Map
                d_910_v23_ = _dafny.Map({p2: d_907_v20_})
                d_911_v24_: _dafny.Array
                nw191_ = _dafny.Array(None, 23)
                nw191_[int(0)] = (_dafny.MultiSet([p2, p2])).set((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({p3: _dafny.SeqWithoutIsStrInference([p2, p2, p2, 673])})) for d_912_i0_ in range(default__.abs(-855))]))), default__.abs(p2))
                nw191_[int(1)] = d_907_v20_
                nw191_[int(2)] = (_dafny.MultiSet([p2])) | (_dafny.MultiSet([len(d_887_v0_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bq"))), p2]))
                nw191_[int(3)] = d_907_v20_
                nw191_[int(4)] = d_907_v20_
                nw191_[int(5)] = (d_907_v20_) - ((d_907_v20_).set(default__.fm0(p0, globalState), default__.abs(len(d_889_v2_))))
                nw191_[int(6)] = (d_907_v20_) | (d_907_v20_)
                nw191_[int(7)] = d_907_v20_
                nw191_[int(8)] = d_907_v20_
                nw191_[int(9)] = d_907_v20_
                nw191_[int(10)] = d_907_v20_
                nw191_[int(11)] = _dafny.MultiSet([p2, (d_909_v22_).cardinality, p2])
                nw191_[int(12)] = d_907_v20_
                nw191_[int(13)] = d_907_v20_
                nw191_[int(14)] = (d_907_v20_) | (d_907_v20_)
                nw191_[int(15)] = ((d_910_v23_)[p2] if (p2) in (d_910_v23_) else d_907_v20_)
                nw191_[int(16)] = _dafny.MultiSet([p2])
                nw191_[int(17)] = default__.fm30(self.f16, globalState)
                nw191_[int(18)] = d_907_v20_
                nw191_[int(19)] = _dafny.MultiSet(d_887_v0_)
                nw191_[int(20)] = _dafny.MultiSet([p2, p2])
                nw191_[int(21)] = d_907_v20_
                nw191_[int(22)] = d_907_v20_
                d_911_v24_ = nw191_
                index161_ = default__.safeIndex(225, (d_911_v24_).length(0))
                (d_911_v24_)[index161_] = d_907_v20_
                index162_ = default__.safeIndex(225, (d_911_v24_).length(0))
                rhs157_ = (_dafny.MultiSet(d_887_v0_)) | (d_907_v20_)
                rhs158_ = p2
                lhs117_ = d_911_v24_
                lhs118_ = default__.safeIndex(225, (d_911_v24_).length(0))
                lhs119_ = globalState
                lhs117_[lhs118_] = rhs157_
                lhs119_.f7 = rhs158_
                d_913_v25_: _dafny.Seq
                d_913_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fk"))
                d_914_v26_: C3
                nw192_ = C3()
                nw192_.ctor__(p2, d_913_v25_, False, not (True) or ((self).f17))
                d_914_v26_ = nw192_
                index163_ = default__.safeIndex(16, (d_908_v21_).length(0))
                (d_908_v21_)[index163_] = (self).f21
                index164_ = default__.safeIndex(16, (d_908_v21_).length(0))
                (d_908_v21_)[index164_] = p1
                d_915_v27_: _dafny.Map
                d_915_v27_ = _dafny.Map({not((d_908_v21_)[default__.safeIndex(16, (d_908_v21_).length(0))]): p2})
                d_916_v28_: _dafny.Map
                d_916_v28_ = _dafny.Map({d_915_v27_: len((default__.fm32(p0, True, p1, p0, globalState)) | (d_915_v27_))})
                d_917_v30_: _dafny.MultiSet
                d_917_v30_ = _dafny.MultiSet([_dafny.Map({self.f16: p2})])
                def iife66_():
                    coll44_ = _dafny.Map()
                    compr_44_: _dafny.Map
                    for compr_44_ in (d_917_v30_).Elements:
                        d_918_v29_: _dafny.Map = compr_44_
                        if (d_918_v29_) in (d_917_v30_):
                            coll44_[d_918_v29_] = (d_914_v26_).f22
                    return _dafny.Map(coll44_)
                d_916_v28_ = (d_916_v28_) | ((d_916_v28_) | (iife66_()
                ))
                index165_ = default__.safeIndex(16, (d_908_v21_).length(0))
                (d_908_v21_)[index165_] = self.f16
            (globalState).f7 = p2
            (globalState).f2 = p2
            (globalState).f2 = default__.safeDivisionInt(p2, p2)
        elif True:
            d_919_v31_: _dafny.Map
            d_919_v31_ = _dafny.Map({((self).f17 if p0 else default__.fm27(137, p2, p2, d_888_v1_, globalState)): default__.fm43(self.f16, globalState)})
            d_920_v33_: _dafny.Map
            d_920_v33_ = _dafny.Map({p2: d_888_v1_})
            d_921_v34_: D11
            d_921_v34_ = D11_DC19(d_888_v1_)
            d_922_v35_: _dafny.Seq
            def iife67_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(650, -586):
                    d_923_v32_: int = compr_45_
                    if ((650) <= (d_923_v32_)) and ((d_923_v32_) < (-586)):
                        coll45_[(d_923_v32_) - (-436)] = p2
                return _dafny.Map(coll45_)
            d_922_v35_ = _dafny.SeqWithoutIsStrInference([iife67_()
            , ((d_920_v33_)[p2] if (p2) in (d_920_v33_) else (d_921_v34_).cf27)])
            d_919_v31_ = (d_919_v31_).set(not((d_888_v1_) in (d_922_v35_)), self.f16)
            d_924_v36_: D3
            d_924_v36_ = D3_DC3((self).f17)
            d_925_v37_: C2
            nw193_ = C2()
            nw193_.ctor__(d_924_v36_, True, True)
            d_925_v37_ = nw193_
            (globalState).f2 = default__.fm0(p1, globalState)
            d_926_v38_: _dafny.Seq
            d_926_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gx"))
            r2 = (d_926_v38_) + (d_926_v38_)
            d_927_v39_: _dafny.Set
            d_927_v39_ = _dafny.Set({(self).f17, p0, (self).f17})
            d_928_v40_: D13
            d_928_v40_ = D13_DC26()
            d_929_v41_: _dafny.Set
            d_929_v41_ = _dafny.Set({d_928_v40_})
            d_930_v42_: D3
            d_930_v42_ = D3_DC4((self).f17, (len(d_927_v39_) if default__.fm43((self).f17, globalState) else len(d_929_v41_)), p2, (self).f17, p1)
            d_931_v43_: _dafny.MultiSet
            d_931_v43_ = _dafny.MultiSet([not(((d_919_v31_)[(self).f17] if ((self).f17) in (d_919_v31_) else self.f16))])
            d_932_v44_: _dafny.MultiSet
            d_932_v44_ = _dafny.MultiSet([p2])
            rhs159_ = d_887_v0_
            rhs160_ = (False if p1 else self.f16)
            rhs161_ = (D3_DC4(p0, p2, p2, (self).f17, p1) if ((d_931_v43_).cardinality) < (p2) else D3_DC4(not(default__.fm43(self.f16, globalState)), (d_932_v44_).cardinality, (0) - (p2), p0, p1))
            rhs162_ = p2
            rhs163_ = not (not(False)) or (self.f16)
            lhs120_ = self
            lhs121_ = globalState
            lhs122_ = self
            d_887_v0_ = rhs159_
            lhs120_.f16 = rhs160_
            d_930_v42_ = rhs161_
            lhs121_.f2 = rhs162_
            lhs122_.f16 = rhs163_
        d_933_v45_: _dafny.Seq
        d_933_v45_ = _dafny.SeqWithoutIsStrInference([True])
        d_934_v46_: str
        d_934_v46_ = _dafny.CodePoint('o')
        rhs164_ = ((d_933_v45_) if (p2) in (_dafny.MultiSet([p2])) else d_933_v45_)
        rhs165_ = d_934_v46_
        d_933_v45_ = rhs164_
        d_934_v46_ = rhs165_
        if ((_dafny.MultiSet([(self).f17])).cardinality) == (p2):
            r3 = self.f16
            (globalState).f7 = p2
            d_935_v47_: _dafny.Array
            def lambda56_(d_936_p1_, d_937_p2_, d_938_p0_):
                def lambda57_(d_939_i1_):
                    return D9_DC16((_dafny.MultiSet([d_936_p1_])).cardinality, (D9_DC16(d_937_p2_, d_938_p0_)).cf24)

                return lambda57_

            init32_ = lambda56_(p1, p2, p0)
            nw194_ = _dafny.Array(None, 13)
            for i0_32_ in range(nw194_.length(0)):
                nw194_[i0_32_] = init32_(i0_32_)
            d_935_v47_ = nw194_
            d_935_v47_ = (d_935_v47_ if p1 else d_935_v47_)
            d_934_v46_ = p3
            r0 = p2
        elif True:
            if (self).f21:
                r1 = ((self).f21 if True else (p2) == ((0) - (p2)))
                d_940_v48_: _dafny.Map
                d_940_v48_ = _dafny.Map({d_887_v0_: p2})
                d_941_v49_: _dafny.Map
                d_941_v49_ = _dafny.Map({p2: (self).f21})
                d_942_v50_: _dafny.MultiSet
                d_942_v50_ = _dafny.MultiSet([p2])
                d_943_v51_: _dafny.Array
                nw195_ = _dafny.Array(None, 4)
                nw195_[int(0)] = (self).fm12(p2, globalState)
                nw195_[int(1)] = (0) - ((len(_dafny.Map({len(d_940_v48_): p2}))) * (len(d_941_v49_)))
                nw195_[int(2)] = p2
                nw195_[int(3)] = (p2) - ((d_942_v50_).cardinality)
                d_943_v51_ = nw195_
                index166_ = default__.safeIndex(360, (d_943_v51_).length(0))
                (d_943_v51_)[index166_] = p2
                index167_ = default__.safeIndex(360, (d_943_v51_).length(0))
                (d_943_v51_)[index167_] = p2
                d_944_v53_: _dafny.MultiSet
                d_944_v53_ = _dafny.MultiSet([(self).f17])
                d_945_v54_: _dafny.Seq
                def iife68_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(533, 487):
                        d_946_v52_: int = compr_46_
                        if ((533) <= (d_946_v52_)) and ((d_946_v52_) < (487)):
                            coll46_[default__.safeModuloInt(d_946_v52_, len(d_888_v1_))] = p1
                    return _dafny.Map(coll46_)
                d_945_v54_ = _dafny.SeqWithoutIsStrInference([(d_941_v49_) | (iife68_()
                ), _dafny.Map({(d_944_v53_).cardinality: self.f16})])
                d_945_v54_ = d_945_v54_
                r3 = not(((self).f21) or (not(self.f16)))
                d_947_v55_: _dafny.Array
                nw196_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_947_v55_ = nw196_
                d_947_v55_ = d_947_v55_
            elif True:
                (globalState).f2 = p2
                d_948_v56_: C1
                nw197_ = C1()
                nw197_.ctor__(self.f16, p2)
                d_948_v56_ = nw197_
                d_949_v57_: _dafny.MultiSet
                d_949_v57_ = _dafny.MultiSet([d_948_v56_])
                d_949_v57_ = d_949_v57_
                (self).f16 = True
                d_950_v58_: _dafny.Seq
                d_950_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdfcanlo"))
                d_951_v59_: _dafny.Array
                nw198_ = _dafny.Array(None, 8)
                nw198_[int(0)] = default__.safeModuloInt(len(d_950_v58_), p2)
                nw198_[int(1)] = p2
                nw198_[int(2)] = default__.safeModuloInt(p2, p2)
                nw198_[int(3)] = (0) - ((p2) - (p2))
                nw198_[int(4)] = len(d_950_v58_)
                nw198_[int(5)] = p2
                nw198_[int(6)] = (p2) * (p2)
                nw198_[int(7)] = len(d_950_v58_)
                d_951_v59_ = nw198_
                d_951_v59_ = d_951_v59_
                d_952_v60_: _dafny.Map
                d_952_v60_ = _dafny.Map({p1: p2})
                d_953_v61_: _dafny.Map
                d_953_v61_ = _dafny.Map({(0) - (len(d_952_v60_)): self.f16})
                rhs166_ = d_948_v56_.f25
                rhs167_ = ((self).f21 if (p3) in (d_950_v58_) else False)
                rhs168_ = (d_953_v61_) != (d_953_v61_)
                rhs169_ = (0) - (p2)
                lhs123_ = self
                lhs124_ = d_948_v56_
                lhs125_ = globalState
                r3 = rhs166_
                lhs123_.f16 = rhs167_
                lhs124_.f25 = rhs168_
                lhs125_.f2 = rhs169_
            d_954_v62_: C1
            nw199_ = C1()
            nw199_.ctor__(self.f16, p2)
            d_954_v62_ = nw199_
            d_955_v63_: C1
            nw200_ = C1()
            nw200_.ctor__(d_954_v62_.f25, p2)
            d_955_v63_ = nw200_
            d_956_v64_: D3
            d_956_v64_ = D3_DC3(d_954_v62_.f25)
            d_957_v65_: C2
            nw201_ = C2()
            nw201_.ctor__(d_956_v64_, p1, False)
            d_957_v65_ = nw201_
            rhs170_ = d_954_v62_
            rhs171_ = self.f16
            rhs172_ = d_957_v65_
            lhs126_ = d_954_v62_
            d_955_v63_ = rhs170_
            lhs126_.f25 = rhs171_
            d_957_v65_ = rhs172_
            d_958_v66_: _dafny.Array
            nw202_ = _dafny.Array(int(0), 1)
            d_958_v66_ = nw202_
            d_959_v67_: _dafny.Map
            d_959_v67_ = _dafny.Map({p0: d_958_v66_})
            d_960_v68_: _dafny.Set
            d_960_v68_ = _dafny.Set({d_958_v66_, d_958_v66_, ((d_959_v67_)[default__.fm43(self.f16, globalState)] if (default__.fm43(self.f16, globalState)) in (d_959_v67_) else d_958_v66_)})
            d_960_v68_ = d_960_v68_
            (globalState).f7 = (p2) - (default__.safeDivisionInt(598, p2))
        (self).f16 = default__.fm43((self).f21, globalState)
        d_933_v45_ = d_933_v45_
        r0 = p2
        d_961_v69_: _dafny.MultiSet
        d_961_v69_ = _dafny.MultiSet([D15_DC30(d_934_v46_)])
        d_962_v71_: _dafny.Set
        d_962_v71_ = _dafny.Set({d_961_v69_, d_961_v69_, d_961_v69_, d_961_v69_, _dafny.MultiSet([D15_DC30(d_934_v46_)])})
        def iife69_():
            coll47_ = _dafny.Map()
            compr_47_: _dafny.MultiSet
            for compr_47_ in (d_962_v71_).Elements:
                d_963_v70_: _dafny.MultiSet = compr_47_
                if (d_963_v70_) in (d_962_v71_):
                    coll47_[d_963_v70_] = 374
            return _dafny.Map(coll47_)
        r1 = not ((d_961_v69_) not in (iife69_()
        )) or ((self).f21)
        d_964_v72_: _dafny.Seq
        d_964_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wo"))
        r2 = (d_964_v72_) + (d_964_v72_)
        r3 = True
        return r0, r1, r2, r3

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_965_v0_: _dafny.Array
        nw203_ = _dafny.Array(_dafny.CodePoint('D'), 6)
        d_965_v0_ = nw203_
        d_966_v1_: _dafny.Map
        d_966_v1_ = _dafny.Map({(self).f21: d_965_v0_})
        d_967_v2_: _dafny.Map
        d_967_v2_ = _dafny.Map({p0: p1})
        d_968_v3_: _dafny.Map
        d_968_v3_ = _dafny.Map({d_966_v1_: ((d_967_v2_)[p0] if (p0) in (d_967_v2_) else (self).f21)})
        d_969_v4_: _dafny.Seq
        d_969_v4_ = _dafny.SeqWithoutIsStrInference([self.f16, ((d_968_v3_)[d_966_v1_] if (d_966_v1_) in (d_968_v3_) else self.f16)])
        d_970_v5_: _dafny.Seq
        d_970_v5_ = d_969_v4_
        source14_ = (d_970_v5_ if False else d_970_v5_)
        d_971___mcc_h0_ = source14_
        d_972_cf0_ = d_971___mcc_h0_
        if (self).f21:
            d_973_v6_: T0
            nw204_ = C1()
            nw204_.ctor__(self.f16, default__.fm0((self).f21, globalState))
            d_973_v6_ = nw204_
            d_974_v7_: D14
            d_974_v7_ = D14_DC28(p0, (self).f17, d_973_v6_)
            rhs173_ = default__.fm0(p1, globalState)
            rhs174_ = p1
            rhs175_ = (d_969_v4_)[default__.safeIndex((0) - (p0), len(d_969_v4_))]
            rhs176_ = (d_974_v7_).cf41
            lhs127_ = globalState
            lhs128_ = self
            lhs129_ = self
            lhs130_ = self
            lhs127_.f2 = rhs173_
            lhs128_.f16 = rhs174_
            lhs129_.f16 = rhs175_
            lhs130_.f16 = rhs176_
            d_969_v4_ = d_972_cf0_
            (self).f16 = True
            d_975_v8_: str
            d_975_v8_ = _dafny.CodePoint('f')
            d_975_v8_ = d_975_v8_
            d_976_v9_: bool
            d_977_v10_: D6
            d_978_v11_: bool
            out20_: bool
            out21_: D6
            out22_: bool
            out20_, out21_, out22_ = (self).m10((self).f17, default__.safeModuloInt(d_973_v6_.f14, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_973_v6_.f14])]))), (p0) + (d_973_v6_.f14), p1, globalState)
            d_976_v9_ = out20_
            d_977_v10_ = out21_
            d_978_v11_ = out22_
        elif True:
            d_979_v12_: _dafny.Array
            def lambda58_(d_980_i0_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_981_i1_ in range(default__.abs(257))]))

            init33_ = lambda58_
            nw205_ = _dafny.Array(None, 18)
            for i0_33_ in range(nw205_.length(0)):
                nw205_[i0_33_] = init33_(i0_33_)
            d_979_v12_ = nw205_
            index168_ = default__.safeIndex(985, (d_979_v12_).length(0))
            (d_979_v12_)[index168_] = (self).f17
            index169_ = default__.safeIndex(985, (d_979_v12_).length(0))
            rhs177_ = d_979_v12_
            rhs178_ = self.f16
            lhs131_ = d_979_v12_
            lhs132_ = default__.safeIndex(985, (d_979_v12_).length(0))
            d_979_v12_ = rhs177_
            lhs131_[lhs132_] = rhs178_
            d_982_v13_: C1
            nw206_ = C1()
            nw206_.ctor__(False, p0)
            d_982_v13_ = nw206_
            d_983_v14_: _dafny.Array
            nw207_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
            d_983_v14_ = nw207_
            d_984_v15_: _dafny.Seq
            d_984_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uiwdx"))
            index170_ = default__.safeIndex(412, (d_983_v14_).length(0))
            (d_983_v14_)[index170_] = d_984_v15_
            d_985_v16_: D12
            d_985_v16_ = D12_DC22(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_986_i2_ in range(default__.abs(385))]))
            index171_ = default__.safeIndex(412, (d_983_v14_).length(0))
            (d_983_v14_)[index171_] = (d_985_v16_).cf29
            (globalState).f2 = (p0) + (520)
            index172_ = default__.safeIndex(985, (d_979_v12_).length(0))
            rhs179_ = ((d_967_v2_)[(p0 if not(p1) else p0)] if ((p0 if not(p1) else p0)) in (d_967_v2_) else (self).f21)
            rhs180_ = d_982_v13_.f25
            lhs133_ = self
            lhs134_ = d_979_v12_
            lhs135_ = default__.safeIndex(985, (d_979_v12_).length(0))
            lhs133_.f16 = rhs179_
            lhs134_[lhs135_] = rhs180_
        d_987_v17_: _dafny.Seq
        d_987_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecyugm"))
        d_988_v18_: T0
        nw208_ = C1()
        nw208_.ctor__(self.f16, p0)
        d_988_v18_ = nw208_
        d_989_v19_: D14
        d_989_v19_ = D14_DC28(len(d_987_v17_), (self).f17, d_988_v18_)
        def iife70_(_pat_let11_0):
            def iife71_(d_990_dt__update__tmp_h0_):
                def iife72_(_pat_let12_0):
                    def iife73_(d_991_dt__update_hcf41_h0_):
                        return D14_DC28((d_990_dt__update__tmp_h0_).cf40, d_991_dt__update_hcf41_h0_, (d_990_dt__update__tmp_h0_).cf42)
                    return iife73_(_pat_let12_0)
                return iife72_((self).f17)
            return iife71_(_pat_let11_0)
        source15_ = iife70_(d_989_v19_)
        if source15_.is_DC28:
            d_992___mcc_h1_ = source15_.cf40
            d_993___mcc_h2_ = source15_.cf41
            d_994___mcc_h3_ = source15_.cf42
            d_995_cf42_ = d_994___mcc_h3_
            d_996_cf41_ = d_993___mcc_h2_
            d_997_cf40_ = d_992___mcc_h1_
            d_998_v20_: _dafny.Set
            d_998_v20_ = _dafny.Set({self.f16})
            d_999_v21_: D15
            d_999_v21_ = D15_DC29(d_998_v20_)
            d_1000_v22_: _dafny.Seq
            d_1000_v22_ = _dafny.SeqWithoutIsStrInference([242])
            d_1001_v23_: _dafny.Map
            d_1001_v23_ = _dafny.Map({d_988_v18_.f14: d_995_cf42_.f14})
            pat_let_tv11_ = d_995_cf42_
            pat_let_tv12_ = d_1000_v22_
            pat_let_tv13_ = d_1001_v23_
            pat_let_tv14_ = globalState
            d_1002_v24_: _dafny.Array
            nw209_ = _dafny.Array(None, 10)
            nw209_[int(0)] = d_999_v21_
            nw209_[int(1)] = d_999_v21_
            nw209_[int(2)] = d_999_v21_
            nw209_[int(3)] = D15_DC29(_dafny.Set({d_996_cf41_}))
            nw209_[int(4)] = d_999_v21_
            nw209_[int(5)] = D15_DC29(d_998_v20_)
            nw209_[int(6)] = D15_DC29(d_998_v20_)
            nw209_[int(7)] = D15_DC29(d_998_v20_)
            nw209_[int(8)] = d_999_v21_
            def iife74_(_pat_let13_0):
                def iife75_(d_1003_dt__update__tmp_h1_):
                    def iife76_(_pat_let14_0):
                        def iife77_(d_1004_dt__update_hcf43_h0_):
                            return D15_DC29(d_1004_dt__update_hcf43_h0_)
                        return iife77_(_pat_let14_0)
                    return iife76_(_dafny.Set({default__.fm27(pat_let_tv11_.f14, -518, len(pat_let_tv12_), pat_let_tv13_, pat_let_tv14_)}))
                return iife75_(_pat_let13_0)
            nw209_[int(9)] = iife74_(d_999_v21_)
            d_1002_v24_ = nw209_
            d_1002_v24_ = d_1002_v24_
            d_1005_v25_: _dafny.Array
            nw210_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_1005_v25_ = nw210_
            d_1005_v25_ = d_1005_v25_
            (globalState).f7 = (0) - ((0) - (d_988_v18_.f14))
            d_1006_v26_: C3
            nw211_ = C3()
            nw211_.ctor__(p0, d_987_v17_, p1, d_996_cf41_)
            d_1006_v26_ = nw211_
        elif True:
            d_1007___mcc_h4_ = source15_.cf39
            d_1008_cf39_ = d_1007___mcc_h4_
            d_1009_v27_: D9
            d_1009_v27_ = D9_DC16(p0, (self).f17)
            d_1010_v28_: _dafny.MultiSet
            d_1010_v28_ = _dafny.MultiSet([(d_1009_v27_).cf24, self.f16, (self).f21, (self).f21, (self).f21])
            d_1011_v29_: _dafny.Set
            d_1011_v29_ = _dafny.Set({d_988_v18_.f14})
            d_1012_v30_: _dafny.Map
            d_1012_v30_ = _dafny.Map({(d_1010_v28_).cardinality: d_1011_v29_})
            d_1013_v31_: _dafny.Map
            d_1013_v31_ = _dafny.Map({p0: p0})
            def iife78_():
                coll48_ = _dafny.Set()
                compr_48_: int
                for compr_48_ in (d_1013_v31_).keys.Elements:
                    d_1014_v32_: int = compr_48_
                    if (d_1014_v32_) in (d_1013_v31_):
                        coll48_ = coll48_.union(_dafny.Set([default__.safeDivisionInt(d_1014_v32_, len(_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference([D4_DC6(_dafny.Map({True: not(True)})), D4_DC6(_dafny.Map({False: True}))])))})))]))
                return _dafny.Set(coll48_)
            d_1012_v30_ = (d_1012_v30_).set(default__.fm0(not(self.f16), globalState), iife78_()
            )
            (d_988_v18_).f14 = d_988_v18_.f14
            d_1015_v33_: _dafny.Array
            def lambda59_(d_1016_v17_):
                def lambda60_(d_1017_i3_):
                    return (d_1017_i3_) * (len(d_1016_v17_))

                return lambda60_

            init34_ = lambda59_(d_987_v17_)
            nw212_ = _dafny.Array(None, 20)
            for i0_34_ in range(nw212_.length(0)):
                nw212_[i0_34_] = init34_(i0_34_)
            d_1015_v33_ = nw212_
            index173_ = default__.safeIndex(988, (d_1015_v33_).length(0))
            (d_1015_v33_)[index173_] = d_988_v18_.f14
            d_1018_v34_: D3
            d_1018_v34_ = D3_DC4((d_972_cf0_)[default__.safeIndex(d_988_v18_.f14, len(d_972_cf0_))], -260, d_988_v18_.f14, True, self.f16)
            index174_ = default__.safeIndex(988, (d_1015_v33_).length(0))
            (d_1015_v33_)[index174_] = (d_1018_v34_).cf6
            d_1019_v35_: _dafny.Seq
            d_1019_v35_ = _dafny.SeqWithoutIsStrInference([d_1018_v34_, d_1018_v34_])
            d_1020_v36_: _dafny.Map
            d_1020_v36_ = _dafny.Map({(self).f17: d_1019_v35_})
            d_1020_v36_ = (d_1020_v36_).set(self.f16, d_1019_v35_)
        d_1021_v37_: _dafny.Array
        nw213_ = _dafny.Array(None, 1)
        nw213_[int(0)] = d_988_v18_.f14
        d_1021_v37_ = nw213_
        index175_ = default__.safeIndex(514, (d_1021_v37_).length(0))
        (d_1021_v37_)[index175_] = p0
        d_1022_v38_: str
        d_1022_v38_ = _dafny.CodePoint('a')
        index176_ = default__.safeIndex(589, (d_965_v0_).length(0))
        (d_965_v0_)[index176_] = d_1022_v38_
        d_1023_v39_: _dafny.Seq
        d_1023_v39_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1024_v40_: _dafny.Array
        nw214_ = _dafny.Array(D12.default()(), 9)
        d_1024_v40_ = nw214_
        d_1025_v41_: _dafny.Map
        d_1025_v41_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_988_v18_.f14 for d_1026_i4_ in range(default__.abs(-487))])): d_1024_v40_})
        d_1027_v42_: _dafny.Array
        nw215_ = _dafny.Array(None, 11)
        nw215_[int(0)] = _dafny.Map({(0) - ((_dafny.MultiSet(d_1023_v39_)).cardinality): d_1024_v40_})
        nw215_[int(1)] = d_1025_v41_
        nw215_[int(2)] = d_1025_v41_
        nw215_[int(3)] = (d_1025_v41_) | (d_1025_v41_)
        nw215_[int(4)] = (d_1025_v41_) | (_dafny.Map({p0: d_1024_v40_}))
        nw215_[int(5)] = d_1025_v41_
        nw215_[int(6)] = (d_1025_v41_) | (d_1025_v41_)
        nw215_[int(7)] = d_1025_v41_
        nw215_[int(8)] = d_1025_v41_
        nw215_[int(9)] = d_1025_v41_
        nw215_[int(10)] = _dafny.Map({p0: d_1024_v40_})
        d_1027_v42_ = nw215_
        index177_ = default__.safeIndex(228, (d_1027_v42_).length(0))
        (d_1027_v42_)[index177_] = d_1025_v41_
        index178_ = default__.safeIndex(514, (d_1021_v37_).length(0))
        index179_ = default__.safeIndex(589, (d_965_v0_).length(0))
        index180_ = default__.safeIndex(228, (d_1027_v42_).length(0))
        rhs181_ = d_988_v18_.f14
        rhs182_ = d_1022_v38_
        rhs183_ = d_1025_v41_
        lhs136_ = d_1021_v37_
        lhs137_ = default__.safeIndex(514, (d_1021_v37_).length(0))
        lhs138_ = d_965_v0_
        lhs139_ = default__.safeIndex(589, (d_965_v0_).length(0))
        lhs140_ = d_1027_v42_
        lhs141_ = default__.safeIndex(228, (d_1027_v42_).length(0))
        lhs136_[lhs137_] = rhs181_
        lhs138_[lhs139_] = rhs182_
        lhs140_[lhs141_] = rhs183_
        if ((d_1021_v37_)[default__.safeIndex(514, (d_1021_v37_).length(0))]) in ((_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(21, len(_dafny.SeqWithoutIsStrInference([p0]))), len(d_987_v17_))):
            (globalState).f2 = len(_dafny.SeqWithoutIsStrInference([self.f16, p1, not((self).f17), default__.fm43((d_969_v4_)[default__.safeIndex((d_1021_v37_)[default__.safeIndex(514, (d_1021_v37_).length(0))], len(d_969_v4_))], globalState)]))
            d_1028_v43_: _dafny.Array
            nw216_ = _dafny.Array(None, 21)
            nw216_[int(0)] = (d_972_cf0_)[default__.safeIndex(d_988_v18_.f14, len(d_972_cf0_))]
            nw216_[int(1)] = (self).f21
            nw216_[int(2)] = True
            nw216_[int(3)] = self.f16
            nw216_[int(4)] = (self).f21
            nw216_[int(5)] = (D3_DC4(p1, 257, d_988_v18_.f14, (self).f17, False)).cf4
            nw216_[int(6)] = (d_972_cf0_)[default__.safeIndex(d_988_v18_.f14, len(d_972_cf0_))]
            nw216_[int(7)] = p1
            nw216_[int(8)] = not((self).f17)
            nw216_[int(9)] = (self).f17
            nw216_[int(10)] = ((d_967_v2_)[d_988_v18_.f14] if (d_988_v18_.f14) in (d_967_v2_) else (self).f21)
            nw216_[int(11)] = (self).f21
            nw216_[int(12)] = (self).f21
            nw216_[int(13)] = (self).f21
            nw216_[int(14)] = not((d_969_v4_)[default__.safeIndex(len(d_987_v17_), len(d_969_v4_))])
            nw216_[int(15)] = self.f16
            nw216_[int(16)] = not((self).f17)
            nw216_[int(17)] = False
            nw216_[int(18)] = True
            nw216_[int(19)] = (self).f21
            nw216_[int(20)] = (self).f21
            d_1028_v43_ = nw216_
            d_1029_v44_: _dafny.Map
            d_1029_v44_ = _dafny.Map({d_987_v17_: d_1028_v43_})
            d_1029_v44_ = ((d_1029_v44_) | (d_1029_v44_)) | (d_1029_v44_)
            d_1030_v45_: D4
            d_1030_v45_ = D4_DC7(d_970_v5_, d_988_v18_.f14)
            d_1031_v46_: _dafny.Map
            d_1031_v46_ = _dafny.Map({p1: True})
            d_1032_v47_: _dafny.Map
            d_1032_v47_ = _dafny.Map({len(d_1031_v46_): (d_1023_v39_)[default__.safeIndex(len(d_967_v2_), len(d_1023_v39_))]})
            d_1033_v48_: D7
            d_1033_v48_ = D7_DC12()
            d_1034_v49_: _dafny.Seq
            d_1034_v49_ = _dafny.SeqWithoutIsStrInference([d_1033_v48_, d_1033_v48_, d_1033_v48_])
            d_1035_v50_: _dafny.MultiSet
            d_1035_v50_ = _dafny.MultiSet([d_988_v18_.f14, (d_1021_v37_)[default__.safeIndex(514, (d_1021_v37_).length(0))], (d_1021_v37_)[default__.safeIndex(514, (d_1021_v37_).length(0))]])
            d_1036_v51_: _dafny.Map
            d_1036_v51_ = _dafny.Map({d_1034_v49_: d_1035_v50_})
            d_1037_v52_: D3
            d_1037_v52_ = D3_DC4(not(not((self).f17)), d_988_v18_.f14, (((d_1036_v51_)[(d_1034_v49_).set(default__.safeIndex(p0, len(d_1034_v49_)), d_1033_v48_)] if ((d_1034_v49_).set(default__.safeIndex(p0, len(d_1034_v49_)), d_1033_v48_)) in (d_1036_v51_) else d_1035_v50_)).cardinality, self.f16, default__.fm27((d_1021_v37_)[default__.safeIndex(514, (d_1021_v37_).length(0))], 36, p0, d_1032_v47_, globalState))
            d_987_v17_ = default__.fm28((d_988_v18_.f14) != ((d_1030_v45_).cf12), d_1032_v47_, d_1037_v52_, default__.fm40(not(p1), globalState), globalState)
            index181_ = default__.safeIndex(372, (d_1028_v43_).length(0))
            (d_1028_v43_)[index181_] = self.f16
            index182_ = default__.safeIndex(372, (d_1028_v43_).length(0))
            (d_1028_v43_)[index182_] = p1
            d_1038_v53_: _dafny.Set
            d_1038_v53_ = _dafny.Set({len(_dafny.Map({not((self).f17): (d_1028_v43_)[default__.safeIndex(372, (d_1028_v43_).length(0))]})), 684, len(d_967_v2_), d_988_v18_.f14, (d_1023_v39_)[default__.safeIndex(p0, len(d_1023_v39_))]})
            d_1039_v54_: _dafny.Map
            d_1039_v54_ = _dafny.Map({d_1038_v53_: False})
            d_1040_v55_: _dafny.Seq
            d_1040_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.Set({(d_1037_v52_).cf6, (d_1021_v37_)[default__.safeIndex(514, (d_1021_v37_).length(0))], d_988_v18_.f14}): (self).f17})])
            d_1039_v54_ = ((d_1040_v55_)[default__.safeIndex(len(d_972_cf0_), len(d_1040_v55_))]).set(default__.fm49(-344, -264, (d_1028_v43_)[default__.safeIndex(372, (d_1028_v43_).length(0))], p0, globalState), self.f16)
        elif True:
            (self).f16 = ((self).fm12(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhqnbdg"))), globalState)) != (d_988_v18_.f14)
            d_969_v4_ = d_969_v4_
            d_1022_v38_ = _dafny.CodePoint('r')
            (globalState).f7 = p0
            d_1041_v56_: _dafny.Array
            nw217_ = _dafny.Array(_dafny.Set({}), 16)
            d_1041_v56_ = nw217_
            d_1042_v57_: C0
            nw218_ = C0()
            nw218_.ctor__(d_1041_v56_, d_988_v18_.f14)
            d_1042_v57_ = nw218_
        d_1043_v58_: _dafny.Set
        d_1043_v58_ = _dafny.Set({p0})
        d_1044_v59_: _dafny.Map
        d_1044_v59_ = _dafny.Map({default__.fm0(self.f16, globalState): d_967_v2_})
        d_1045_v60_: _dafny.Seq
        d_1045_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llljm"))
        d_1046_v61_: _dafny.MultiSet
        d_1046_v61_ = _dafny.MultiSet([self.f16, (self).f21])
        d_1047_v62_: D9
        d_1047_v62_ = D9_DC16(p0, (self).f17)
        d_1048_v63_: _dafny.Map
        d_1048_v63_ = _dafny.Map({(d_1047_v62_).cf23: p0})
        d_1049_v64_: _dafny.Array
        nw219_ = _dafny.Array(None, 20)
        nw219_[int(0)] = not(not(True))
        nw219_[int(1)] = (d_1043_v58_) != (default__.fm49(p0, p0, self.f16, p0, globalState))
        nw219_[int(2)] = (self).f17
        nw219_[int(3)] = p1
        nw219_[int(4)] = self.f16
        nw219_[int(5)] = self.f16
        nw219_[int(6)] = not((self).f17)
        nw219_[int(7)] = p1
        nw219_[int(8)] = p1
        nw219_[int(9)] = (self).f17
        nw219_[int(10)] = default__.fm27(len(d_1044_v59_), p0, len((d_1045_v60_).set(default__.safeIndex((d_1046_v61_).cardinality, len(d_1045_v60_)), _dafny.CodePoint('o'))), d_1048_v63_, globalState)
        nw219_[int(11)] = p1
        nw219_[int(12)] = default__.fm27(p0, (self).fm12(p0, globalState), p0, d_1048_v63_, globalState)
        nw219_[int(13)] = (self).f21
        nw219_[int(14)] = (p0) == (p0)
        nw219_[int(15)] = self.f16
        nw219_[int(16)] = self.f16
        nw219_[int(17)] = p1
        nw219_[int(18)] = p1
        nw219_[int(19)] = ((self).f17 if (self).f21 else (self).f21)
        d_1049_v64_ = nw219_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1049_v64_).length(0)):
            d_1050_i5_: int = guard_loop_5_
            if (True) and (((0) <= (d_1050_i5_)) and ((d_1050_i5_) < ((d_1049_v64_).length(0)))):
                (d_1049_v64_)[(d_1050_i5_)] = (p0) == (p0)
        d_1051_v65_: _dafny.Map
        d_1051_v65_ = _dafny.Map({d_1048_v63_: (p0) * ((d_1047_v62_).cf23)})
        d_1051_v65_ = (d_1051_v65_).set(d_1048_v63_, p0)
        rhs184_ = not(p1)
        rhs185_ = p1
        lhs142_ = self
        lhs143_ = self
        lhs142_.f16 = rhs184_
        lhs143_.f16 = rhs185_
        if not(not(False)):
            index183_ = default__.safeIndex(722, (d_1049_v64_).length(0))
            (d_1049_v64_)[index183_] = (self).f17
            index184_ = default__.safeIndex(722, (d_1049_v64_).length(0))
            (d_1049_v64_)[index184_] = self.f16
            (self).f16 = self.f16
            d_1045_v60_ = d_1045_v60_
            d_1052_v66_: _dafny.Set
            d_1052_v66_ = _dafny.Set({(self).f21})
            d_1053_v67_: D12
            d_1053_v67_ = D12_DC23((self).f21, p0, len(d_1052_v66_))
            source16_ = d_1053_v67_
            if source16_.is_DC23:
                d_1054___mcc_h5_ = source16_.cf30
                d_1055___mcc_h6_ = source16_.cf31
                d_1056___mcc_h7_ = source16_.cf32
                d_1057_cf32_ = d_1056___mcc_h7_
                d_1058_cf31_ = d_1055___mcc_h6_
                d_1059_cf30_ = d_1054___mcc_h5_
                d_1059_cf30_ = (self).f17
                index185_ = default__.safeIndex(722, (d_1049_v64_).length(0))
                (d_1049_v64_)[index185_] = self.f16
                d_1060_v68_: _dafny.Array
                def lambda61_(d_1061_cf30_, d_1062_p1_):
                    def lambda62_(d_1063_i6_):
                        return (_dafny.Map({d_1061_cf30_: (self).f17})) | (_dafny.Map({d_1061_cf30_: d_1062_p1_}))

                    return lambda62_

                init35_ = lambda61_(d_1059_cf30_, p1)
                nw220_ = _dafny.Array(None, 28)
                for i0_35_ in range(nw220_.length(0)):
                    nw220_[i0_35_] = init35_(i0_35_)
                d_1060_v68_ = nw220_
                d_1064_v69_: _dafny.Map
                d_1064_v69_ = _dafny.Map({not((self).f17): self.f16})
                index186_ = default__.safeIndex(402, (d_1060_v68_).length(0))
                (d_1060_v68_)[index186_] = (d_1064_v69_) | ((d_1064_v69_).set((d_1049_v64_)[default__.safeIndex(722, (d_1049_v64_).length(0))], (self).f21))
                d_1065_v70_: _dafny.Seq
                d_1065_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1049_v64_)[default__.safeIndex(722, (d_1049_v64_).length(0))]: (d_1049_v64_)[default__.safeIndex(722, (d_1049_v64_).length(0))]}), d_1064_v69_])
                index187_ = default__.safeIndex(402, (d_1060_v68_).length(0))
                (d_1060_v68_)[index187_] = (d_1065_v70_)[default__.safeIndex(len((d_1052_v66_) | (d_1052_v66_)), len(d_1065_v70_))]
                d_1066_v71_: str
                d_1066_v71_ = _dafny.CodePoint('e')
                d_1066_v71_ = d_1066_v71_
            elif source16_.is_DC24:
                d_1067___mcc_h8_ = source16_.cf33
                d_1068___mcc_h9_ = source16_.cf34
                d_1069___mcc_h10_ = source16_.cf35
                d_1070___mcc_h11_ = source16_.cf36
                d_1071___mcc_h12_ = source16_.cf37
                d_1072_cf37_ = d_1071___mcc_h12_
                d_1073_cf36_ = d_1070___mcc_h11_
                d_1074_cf35_ = d_1069___mcc_h10_
                d_1075_cf34_ = d_1068___mcc_h9_
                d_1076_cf33_ = d_1067___mcc_h8_
                d_1077_v72_: _dafny.Array
                def lambda63_(d_1078_v66_):
                    def lambda64_(d_1079_i7_):
                        return d_1078_v66_

                    return lambda64_

                init36_ = lambda63_(d_1052_v66_)
                nw221_ = _dafny.Array(None, 6)
                for i0_36_ in range(nw221_.length(0)):
                    nw221_[i0_36_] = init36_(i0_36_)
                d_1077_v72_ = nw221_
                d_1080_v73_: C0
                nw222_ = C0()
                nw222_.ctor__(d_1077_v72_, p0)
                d_1080_v73_ = nw222_
                (globalState).f7 = d_1080_v73_.f27
                d_1081_v74_: _dafny.Map
                d_1081_v74_ = _dafny.Map({d_1076_cf33_: p0})
                d_1082_v75_: _dafny.MultiSet
                d_1082_v75_ = _dafny.MultiSet([d_1081_v74_, d_1081_v74_, d_1081_v74_, d_1081_v74_, d_1081_v74_])
                d_1083_v76_: _dafny.Seq
                d_1083_v76_ = _dafny.SeqWithoutIsStrInference([default__.fm54(p0, (self).f21, d_1080_v73_.f27, not(True), globalState)])
                d_1084_v77_: _dafny.Seq
                d_1084_v77_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: p0}), d_1081_v74_, d_1081_v74_])
                d_1082_v75_ = ((d_1082_v75_ if False else d_1082_v75_)) | (_dafny.MultiSet(((d_1083_v76_)[default__.safeIndex(d_1080_v73_.f27, len(d_1083_v76_))]) + ((d_1084_v77_).set(default__.safeIndex(p0, len(d_1084_v77_)), d_1081_v74_))))
                d_1085_v78_: _dafny.Array
                nw223_ = _dafny.Array(int(0), 5)
                d_1085_v78_ = nw223_
                index188_ = default__.safeIndex(905, (d_1085_v78_).length(0))
                (d_1085_v78_)[index188_] = d_1080_v73_.f27
                index189_ = default__.safeIndex(905, (d_1085_v78_).length(0))
                rhs186_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1086_i8_ in range(default__.abs(331))])
                rhs187_ = d_1080_v73_.f27
                rhs188_ = self.f16
                lhs144_ = d_1085_v78_
                lhs145_ = default__.safeIndex(905, (d_1085_v78_).length(0))
                d_1045_v60_ = rhs186_
                lhs144_[lhs145_] = rhs187_
                d_1073_cf36_ = rhs188_
            elif True:
                d_1087___mcc_h13_ = source16_.cf29
                d_1088_cf29_ = d_1087___mcc_h13_
                (globalState).f2 = p0
                d_1045_v60_ = d_1045_v60_
                d_1089_v79_: _dafny.Seq
                d_1089_v79_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1090_v80_: T0
                nw224_ = C1()
                nw224_.ctor__(self.f16, p0)
                d_1090_v80_ = nw224_
                d_1091_v81_: D3
                d_1091_v81_ = D3_DC4((D14_DC28(p0, (self).f17, d_1090_v80_)).cf41, d_1090_v80_.f14, len(default__.fm42(False, len(d_1043_v58_), d_1090_v80_.f14, p1, globalState)), (d_1049_v64_)[default__.safeIndex(722, (d_1049_v64_).length(0))], (self).f21)
                d_1092_v82_: _dafny.Map
                d_1092_v82_ = _dafny.Map({(self).f17: self.f16})
                (globalState).f7 = default__.safeDivisionInt((d_1089_v79_)[default__.safeIndex(len(default__.fm28(p1, d_1048_v63_, d_1091_v81_, d_1092_v82_, globalState)), len(d_1089_v79_))], (default__.fm0(self.f16, globalState)) * (len(_dafny.Map({p0: d_1090_v80_.f14}))))
                d_1093_v83_: str
                d_1093_v83_ = _dafny.CodePoint('p')
                d_1094_v84_: _dafny.Array
                nw225_ = _dafny.Array(None, 9)
                nw225_[int(0)] = d_1052_v66_
                nw225_[int(1)] = _dafny.Set({(self).f17})
                nw225_[int(2)] = _dafny.Set({p1})
                nw225_[int(3)] = d_1052_v66_
                nw225_[int(4)] = d_1052_v66_
                nw225_[int(5)] = d_1052_v66_
                nw225_[int(6)] = d_1052_v66_
                nw225_[int(7)] = _dafny.Set({(d_1049_v64_)[default__.safeIndex(722, (d_1049_v64_).length(0))]})
                nw225_[int(8)] = d_1052_v66_
                d_1094_v84_ = nw225_
                d_1095_v85_: C0
                nw226_ = C0()
                nw226_.ctor__(d_1094_v84_, len(d_1089_v79_))
                d_1095_v85_ = nw226_
                d_1096_v86_: _dafny.Map
                d_1096_v86_ = _dafny.Map({d_1095_v85_: d_1093_v83_})
                d_1097_v87_: _dafny.Map
                d_1097_v87_ = _dafny.Map({p0: d_1093_v83_})
                index190_ = default__.safeIndex(722, (d_1049_v64_).length(0))
                rhs189_ = (((d_1096_v86_)[d_1095_v85_] if (d_1095_v85_) in (d_1096_v86_) else d_1093_v83_) if (self).f21 else ((d_1097_v87_)[312] if (312) in (d_1097_v87_) else d_1093_v83_))
                rhs190_ = d_1045_v60_
                rhs191_ = (self).f17
                rhs192_ = not (not((self).f17)) or ((self).f21)
                lhs146_ = self
                lhs147_ = d_1049_v64_
                lhs148_ = default__.safeIndex(722, (d_1049_v64_).length(0))
                d_1093_v83_ = rhs189_
                d_1045_v60_ = rhs190_
                lhs146_.f16 = rhs191_
                lhs147_[lhs148_] = rhs192_
            (globalState).f2 = default__.fm0((d_1049_v64_)[default__.safeIndex(722, (d_1049_v64_).length(0))], globalState)
        elif True:
            (self).f16 = self.f16
            index191_ = default__.safeIndex(973, (d_1049_v64_).length(0))
            (d_1049_v64_)[index191_] = not(p1)
            index192_ = default__.safeIndex(973, (d_1049_v64_).length(0))
            rhs193_ = not(not (True) or (not(p1)))
            rhs194_ = p0
            lhs149_ = d_1049_v64_
            lhs150_ = default__.safeIndex(973, (d_1049_v64_).length(0))
            lhs149_[lhs150_] = rhs193_
            r0 = rhs194_
            index193_ = default__.safeIndex(973, (d_1049_v64_).length(0))
            (d_1049_v64_)[index193_] = self.f16
            d_1098_v88_: str
            d_1098_v88_ = _dafny.CodePoint('q')
            d_1099_v89_: _dafny.Array
            nw227_ = _dafny.Array(_dafny.Set({}), 10)
            d_1099_v89_ = nw227_
            d_1100_v90_: C0
            nw228_ = C0()
            nw228_.ctor__(d_1099_v89_, p0)
            d_1100_v90_ = nw228_
            d_1101_v91_: T0
            nw229_ = C1()
            nw229_.ctor__((d_1049_v64_)[default__.safeIndex(973, (d_1049_v64_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_1100_v90_, d_1100_v90_])))
            d_1101_v91_ = nw229_
            d_1102_v92_: _dafny.Seq
            d_1102_v92_ = _dafny.SeqWithoutIsStrInference([d_1101_v91_, d_1101_v91_, d_1101_v91_, d_1101_v91_])
            d_1103_v93_: C3
            nw230_ = C3()
            nw230_.ctor__(d_1101_v91_.f14, _dafny.SeqWithoutIsStrInference([d_1098_v88_ for d_1104_i9_ in range(default__.abs(354))]), (self).f21, p1)
            d_1103_v93_ = nw230_
            d_1105_v94_: _dafny.Map
            d_1105_v94_ = _dafny.Map({d_1103_v93_: d_1100_v90_.f27})
            d_1106_v95_: _dafny.MultiSet
            d_1106_v95_ = _dafny.MultiSet([d_1101_v91_.f14, 846])
            d_1107_v96_: _dafny.Array
            nw231_ = _dafny.Array(None, 19)
            nw231_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "te")))
            nw231_[int(1)] = p0
            nw231_[int(2)] = p0
            nw231_[int(3)] = len(d_1048_v63_)
            nw231_[int(4)] = p0
            nw231_[int(5)] = p0
            nw231_[int(6)] = p0
            nw231_[int(7)] = len(d_1043_v58_)
            nw231_[int(8)] = p0
            nw231_[int(9)] = p0
            nw231_[int(10)] = len(d_1102_v92_)
            nw231_[int(11)] = d_1101_v91_.f14
            nw231_[int(12)] = len(d_969_v4_)
            nw231_[int(13)] = len(d_1105_v94_)
            nw231_[int(14)] = len((d_1103_v93_).f23)
            nw231_[int(15)] = d_1101_v91_.f14
            nw231_[int(16)] = (d_1106_v95_).cardinality
            nw231_[int(17)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ld")))
            nw231_[int(18)] = d_1100_v90_.f27
            d_1107_v96_ = nw231_
            d_1108_v97_: _dafny.Array
            def lambda65_(d_1109_i10_):
                return False

            init37_ = lambda65_
            nw232_ = _dafny.Array(None, 27)
            for i0_37_ in range(nw232_.length(0)):
                nw232_[i0_37_] = init37_(i0_37_)
            d_1108_v97_ = nw232_
            d_1110_v98_: D9
            d_1110_v98_ = D9_DC15(p1, self.f16, d_1043_v58_, d_1107_v96_, d_1108_v97_)
            pat_let_tv15_ = d_1108_v97_
            pat_let_tv16_ = p1
            d_1111_v99_: _dafny.Array
            nw233_ = _dafny.Array(None, 28)
            nw233_[int(0)] = d_1110_v98_
            nw233_[int(1)] = D9_DC15(self.f16, (self).f17, d_1043_v58_, d_1107_v96_, d_1108_v97_)
            nw233_[int(2)] = d_1110_v98_
            nw233_[int(3)] = d_1110_v98_
            def iife80_(_pat_let16_0):
                def iife81_(d_1112_dt__update__tmp_h2_):
                    def iife82_(_pat_let17_0):
                        def iife83_(d_1113_dt__update_hcf22_h0_):
                            return D9_DC15((d_1112_dt__update__tmp_h2_).cf18, (d_1112_dt__update__tmp_h2_).cf19, (d_1112_dt__update__tmp_h2_).cf20, (d_1112_dt__update__tmp_h2_).cf21, d_1113_dt__update_hcf22_h0_)
                        return iife83_(_pat_let17_0)
                    return iife82_(pat_let_tv15_)
                return iife81_(_pat_let16_0)
            def iife79_(_pat_let15_0):
                def iife84_(d_1114_dt__update__tmp_h3_):
                    def iife85_(_pat_let18_0):
                        def iife86_(d_1115_dt__update_hcf18_h0_):
                            return D9_DC15(d_1115_dt__update_hcf18_h0_, (d_1114_dt__update__tmp_h3_).cf19, (d_1114_dt__update__tmp_h3_).cf20, (d_1114_dt__update__tmp_h3_).cf21, (d_1114_dt__update__tmp_h3_).cf22)
                        return iife86_(_pat_let18_0)
                    return iife85_(pat_let_tv16_)
                return iife84_(_pat_let15_0)
            nw233_[int(4)] = iife79_(iife80_(d_1110_v98_))
            nw233_[int(5)] = d_1110_v98_
            nw233_[int(6)] = d_1110_v98_
            nw233_[int(7)] = d_1110_v98_
            nw233_[int(8)] = D9_DC15((self).f21, p1, d_1043_v58_, d_1107_v96_, d_1108_v97_)
            nw233_[int(9)] = d_1110_v98_
            nw233_[int(10)] = d_1110_v98_
            nw233_[int(11)] = d_1110_v98_
            nw233_[int(12)] = d_1110_v98_
            nw233_[int(13)] = d_1110_v98_
            nw233_[int(14)] = d_1110_v98_
            nw233_[int(15)] = d_1110_v98_
            def iife87_(_pat_let19_0):
                def iife88_(d_1116_dt__update__tmp_h4_):
                    def iife89_(_pat_let20_0):
                        def iife90_(d_1117_dt__update_hcf19_h0_):
                            return D9_DC15((d_1116_dt__update__tmp_h4_).cf18, d_1117_dt__update_hcf19_h0_, (d_1116_dt__update__tmp_h4_).cf20, (d_1116_dt__update__tmp_h4_).cf21, (d_1116_dt__update__tmp_h4_).cf22)
                        return iife90_(_pat_let20_0)
                    return iife89_(True)
                return iife88_(_pat_let19_0)
            nw233_[int(16)] = iife87_(D9_DC15(p1, (self).f17, d_1043_v58_, d_1107_v96_, d_1108_v97_))
            nw233_[int(17)] = d_1110_v98_
            nw233_[int(18)] = d_1110_v98_
            nw233_[int(19)] = d_1110_v98_
            nw233_[int(20)] = d_1110_v98_
            nw233_[int(21)] = d_1110_v98_
            nw233_[int(22)] = D9_DC15((d_1049_v64_)[default__.safeIndex(973, (d_1049_v64_).length(0))], (d_1049_v64_)[default__.safeIndex(973, (d_1049_v64_).length(0))], d_1043_v58_, d_1107_v96_, d_1108_v97_)
            nw233_[int(23)] = d_1110_v98_
            nw233_[int(24)] = d_1110_v98_
            nw233_[int(25)] = D9_DC15((self).f17, (d_1049_v64_)[default__.safeIndex(973, (d_1049_v64_).length(0))], _dafny.Set({(d_1103_v93_).f22}), d_1107_v96_, d_1108_v97_)
            nw233_[int(26)] = d_1110_v98_
            nw233_[int(27)] = D9_DC15(p1, (self).f17, d_1043_v58_, d_1107_v96_, d_1108_v97_)
            d_1111_v99_ = nw233_
            d_1118_v100_: _dafny.Seq
            d_1118_v100_ = _dafny.SeqWithoutIsStrInference([d_1111_v99_, d_1111_v99_, d_1111_v99_])
            d_1119_v101_: D3
            d_1119_v101_ = D3_DC3((d_1049_v64_)[default__.safeIndex(973, (d_1049_v64_).length(0))])
            d_1120_v102_: _dafny.Seq
            d_1120_v102_ = _dafny.SeqWithoutIsStrInference([d_1119_v101_])
            d_1121_v103_: D12
            d_1121_v103_ = D12_DC24(self.f16, (d_969_v4_).set(default__.safeIndex(p0, len(d_969_v4_)), p1), d_1118_v100_, (d_1119_v101_).cf3, d_1120_v102_)
            d_1122_v104_: _dafny.Map
            d_1122_v104_ = _dafny.Map({default__.fm43(not((self).f17), globalState): p1})
            rhs195_ = d_1098_v88_
            rhs196_ = ((d_1121_v103_).cf33) or (self.f16)
            rhs197_ = default__.fm0(((d_1122_v104_)[(self).f17] if ((self).f17) in (d_1122_v104_) else (self).f17), globalState)
            lhs151_ = self
            lhs152_ = globalState
            d_1098_v88_ = rhs195_
            lhs151_.f16 = rhs196_
            lhs152_.f2 = rhs197_
            d_1123_v105_: C2
            nw234_ = C2()
            nw234_.ctor__(d_1119_v101_, (d_969_v4_)[default__.safeIndex((d_1103_v93_).f22, len(d_969_v4_))], p1)
            d_1123_v105_ = nw234_
        d_1124_v106_: D3
        d_1124_v106_ = D3_DC3((self).f21)
        source17_ = d_1124_v106_
        if source17_.is_DC4:
            d_1125___mcc_h14_ = source17_.cf4
            d_1126___mcc_h15_ = source17_.cf5
            d_1127___mcc_h16_ = source17_.cf6
            d_1128___mcc_h17_ = source17_.cf7
            d_1129___mcc_h18_ = source17_.cf8
            d_1130_cf8_ = d_1129___mcc_h18_
            d_1131_cf7_ = d_1128___mcc_h17_
            d_1132_cf6_ = d_1127___mcc_h16_
            d_1133_cf5_ = d_1126___mcc_h15_
            d_1134_cf4_ = d_1125___mcc_h14_
            d_1135_v107_: D16
            d_1135_v107_ = D16_DC32(d_1046_v61_)
            d_1136_v108_: bool
            d_1137_v109_: D6
            d_1138_v110_: bool
            out23_: bool
            out24_: D6
            out25_: bool
            out23_, out24_, out25_ = (self).m10((self).f17, p0, (((d_1135_v107_).cf46).cardinality if self.f16 else 92), d_1134_cf4_, globalState)
            d_1136_v108_ = out23_
            d_1137_v109_ = out24_
            d_1138_v110_ = out25_
            (self).f16 = not((d_1136_v108_) not in (d_969_v4_))
            if (default__.fm42(self.f16, 537, (self).fm12(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtf"))), globalState), True, globalState)) <= (d_1045_v60_):
                d_1139_v111_: _dafny.Array
                def lambda66_(d_1140_v4_):
                    def lambda67_(d_1141_i11_):
                        return d_1140_v4_

                    return lambda67_

                init38_ = lambda66_(d_969_v4_)
                nw235_ = _dafny.Array(None, 21)
                for i0_38_ in range(nw235_.length(0)):
                    nw235_[i0_38_] = init38_(i0_38_)
                d_1139_v111_ = nw235_
                d_1139_v111_ = d_1139_v111_
                d_1142_v112_: _dafny.Map
                d_1142_v112_ = _dafny.Map({p0: d_1045_v60_})
                d_1133_cf5_ = (default__.safeModuloInt(p0, p0)) - (len(d_1142_v112_))
                d_1143_v113_: _dafny.Map
                d_1143_v113_ = _dafny.Map({d_1130_cf8_: p1})
                d_1143_v113_ = (d_1143_v113_).set((d_1130_cf8_ if d_1138_v110_ else self.f16), p1)
                d_1144_v114_: _dafny.MultiSet
                d_1144_v114_ = _dafny.MultiSet([d_1132_cf6_, d_1132_cf6_, ((d_1046_v61_) | (d_1046_v61_)).cardinality])
                d_1145_v115_: _dafny.Map
                d_1145_v115_ = _dafny.Map({d_1144_v114_: p0})
                rhs198_ = d_1131_cf7_
                rhs199_ = d_1144_v114_
                rhs200_ = (0) - (((default__.fm55(True, p0, globalState)).cf46).cardinality)
                rhs201_ = 340
                rhs202_ = ((d_1144_v114_) - (_dafny.MultiSet([len(d_1045_v60_), len(d_1145_v115_)]))).intersection((d_1144_v114_ if d_1134_cf4_ else _dafny.MultiSet([d_1133_cf5_])))
                lhs153_ = globalState
                d_1138_v110_ = rhs198_
                d_1144_v114_ = rhs199_
                lhs153_.f2 = rhs200_
                d_1133_cf5_ = rhs201_
                d_1144_v114_ = rhs202_
                d_1146_v116_: bool
                d_1147_v117_: D6
                d_1148_v118_: bool
                out26_: bool
                out27_: D6
                out28_: bool
                out26_, out27_, out28_ = (self).m10(d_1134_cf4_, d_1133_cf5_, d_1133_cf5_, d_1134_cf4_, globalState)
                d_1146_v116_ = out26_
                d_1147_v117_ = out27_
                d_1148_v118_ = out28_
            elif True:
                index194_ = default__.safeIndex(307, (d_1049_v64_).length(0))
                (d_1049_v64_)[index194_] = self.f16
                index195_ = default__.safeIndex(307, (d_1049_v64_).length(0))
                (d_1049_v64_)[index195_] = ((d_967_v2_)[default__.fm0(d_1131_cf7_, globalState)] if (default__.fm0(d_1131_cf7_, globalState)) in (d_967_v2_) else not(default__.fm31(d_1131_cf7_, d_1134_cf4_, globalState)))
                d_1136_v108_ = p1
                d_1149_v119_: _dafny.Set
                d_1149_v119_ = _dafny.Set({self.f16})
                d_1150_v120_: _dafny.Map
                d_1150_v120_ = _dafny.Map({d_1134_cf4_: p1})
                (globalState).f7 = ((d_1046_v61_)[(len(d_1048_v63_)) <= (len(d_1149_v119_))] if ((len(d_1048_v63_)) <= (len(d_1149_v119_))) in (d_1046_v61_) else len(d_1150_v120_))
                d_1151_v121_: str
                d_1151_v121_ = _dafny.CodePoint('s')
                index196_ = default__.safeIndex(228, (d_965_v0_).length(0))
                (d_965_v0_)[index196_] = d_1151_v121_
                index197_ = default__.safeIndex(228, (d_965_v0_).length(0))
                (d_965_v0_)[index197_] = d_1151_v121_
                d_1152_v122_: C1
                nw236_ = C1()
                nw236_.ctor__(not (p1) or (d_1138_v110_), d_1132_cf6_)
                d_1152_v122_ = nw236_
            d_1153_v123_: _dafny.Array
            nw237_ = _dafny.Array(int(0), 18)
            d_1153_v123_ = nw237_
            index198_ = default__.safeIndex(161, (d_1153_v123_).length(0))
            (d_1153_v123_)[index198_] = default__.safeModuloInt((0) - (p0), len(d_1045_v60_))
            index199_ = default__.safeIndex(161, (d_1153_v123_).length(0))
            (d_1153_v123_)[index199_] = p0
        elif source17_.is_DC3:
            d_1154___mcc_h19_ = source17_.cf3
            d_1155_cf3_ = d_1154___mcc_h19_
            if (self).f17:
                d_1156_v124_: _dafny.MultiSet
                d_1156_v124_ = _dafny.MultiSet([p0])
                d_1157_v125_: _dafny.Map
                d_1157_v125_ = _dafny.Map({True: p0})
                d_1158_v126_: _dafny.MultiSet
                d_1158_v126_ = _dafny.MultiSet([((d_1156_v124_)[p0] if (p0) in (d_1156_v124_) else p0), len(d_1157_v125_)])
                d_1159_v127_: C3
                nw238_ = C3()
                nw238_.ctor__(default__.safeDivisionInt(p0, p0), d_1045_v60_, p1, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1046_v61_).cardinality]))).issubset(d_1158_v126_))
                d_1159_v127_ = nw238_
                d_969_v4_ = d_969_v4_
                d_1160_v128_: _dafny.Array
                nw239_ = _dafny.Array(None, 4)
                nw239_[int(0)] = (d_1159_v127_).f22
                nw239_[int(1)] = (d_1159_v127_).f22
                nw239_[int(2)] = len((d_1159_v127_).f23)
                nw239_[int(3)] = p0
                d_1160_v128_ = nw239_
                d_1161_v129_: _dafny.Map
                d_1161_v129_ = _dafny.Map({(self).f21: d_1160_v128_})
                d_1162_v130_: _dafny.Seq
                d_1162_v130_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1163_v131_: _dafny.Array
                nw240_ = _dafny.Array(None, 15)
                nw240_[int(0)] = (-291 if p1 else len(d_1161_v129_))
                nw240_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsmajdos")))
                nw240_[int(2)] = (d_1159_v127_).f22
                nw240_[int(3)] = ((0) - (p0)) - (len((self).fm13((d_1159_v127_).f22, (self).f21, False, globalState)))
                nw240_[int(4)] = (d_1159_v127_).f22
                nw240_[int(5)] = (d_1159_v127_).f22
                nw240_[int(6)] = (d_1159_v127_).f22
                nw240_[int(7)] = (865) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mksbkds"))))
                nw240_[int(8)] = default__.safeModuloInt((d_1159_v127_).f22, (d_1159_v127_).f22)
                nw240_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjpp")))
                nw240_[int(10)] = (d_1162_v130_)[default__.safeIndex(p0, len(d_1162_v130_))]
                nw240_[int(11)] = (0) - ((0) - (p0))
                nw240_[int(12)] = (d_1159_v127_).f22
                nw240_[int(13)] = len((d_1159_v127_).f23)
                nw240_[int(14)] = p0
                d_1163_v131_ = nw240_
                index200_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                (d_1160_v128_)[index200_] = p0
                index201_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                (d_1160_v128_)[index201_] = p0
                d_1164_v132_: _dafny.Array
                nw241_ = _dafny.Array(None, 1)
                nw241_[int(0)] = d_1049_v64_
                d_1164_v132_ = nw241_
                d_1165_v133_: _dafny.Array
                def lambda68_(d_1166_v60_):
                    def lambda69_(d_1167_i12_):
                        return d_1166_v60_

                    return lambda69_

                init39_ = lambda68_(d_1045_v60_)
                nw242_ = _dafny.Array(None, 26)
                for i0_39_ in range(nw242_.length(0)):
                    nw242_[i0_39_] = init39_(i0_39_)
                d_1165_v133_ = nw242_
                index202_ = default__.safeIndex(708, (d_1165_v133_).length(0))
                (d_1165_v133_)[index202_] = (d_1045_v60_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "catdtpru")))
                index203_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                index204_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                index205_ = default__.safeIndex(708, (d_1165_v133_).length(0))
                rhs203_ = p0
                rhs204_ = d_1164_v132_
                rhs205_ = (d_1159_v127_).f22
                rhs206_ = (d_1045_v60_) + ((d_1159_v127_).f23)
                lhs154_ = d_1160_v128_
                lhs155_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                lhs156_ = d_1160_v128_
                lhs157_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                lhs158_ = d_1165_v133_
                lhs159_ = default__.safeIndex(708, (d_1165_v133_).length(0))
                lhs154_[lhs155_] = rhs203_
                d_1164_v132_ = rhs204_
                lhs156_[lhs157_] = rhs205_
                lhs158_[lhs159_] = rhs206_
                index206_ = default__.safeIndex(57, (d_1049_v64_).length(0))
                (d_1049_v64_)[index206_] = d_1155_cf3_
                index207_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                index208_ = default__.safeIndex(57, (d_1049_v64_).length(0))
                index209_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                rhs207_ = self.f16
                rhs208_ = (d_1159_v127_).f22
                rhs209_ = default__.fm27((d_1160_v128_)[default__.safeIndex(423, (d_1160_v128_).length(0))], (p0) + ((0) - ((0) - (p0))), (0) - (default__.fm0((self).f17, globalState)), d_1048_v63_, globalState)
                rhs210_ = p0
                rhs211_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1168_i13_ in range(default__.abs(883))])
                lhs160_ = d_1160_v128_
                lhs161_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                lhs162_ = d_1049_v64_
                lhs163_ = default__.safeIndex(57, (d_1049_v64_).length(0))
                lhs164_ = d_1160_v128_
                lhs165_ = default__.safeIndex(423, (d_1160_v128_).length(0))
                d_1155_cf3_ = rhs207_
                lhs160_[lhs161_] = rhs208_
                lhs162_[lhs163_] = rhs209_
                lhs164_[lhs165_] = rhs210_
                d_1045_v60_ = rhs211_
            elif True:
                d_1169_v134_: _dafny.Array
                def lambda70_(d_1170_p0_):
                    def lambda71_(d_1171_i14_):
                        return (d_1171_i14_) - (d_1170_p0_)

                    return lambda71_

                init40_ = lambda70_(p0)
                nw243_ = _dafny.Array(None, 21)
                for i0_40_ in range(nw243_.length(0)):
                    nw243_[i0_40_] = init40_(i0_40_)
                d_1169_v134_ = nw243_
                d_1172_v135_: _dafny.Set
                d_1172_v135_ = _dafny.Set({d_1049_v64_})
                d_1173_v136_: T0
                nw244_ = C1()
                nw244_.ctor__((self).f17, len(d_1172_v135_))
                d_1173_v136_ = nw244_
                d_1174_v137_: D14
                d_1174_v137_ = D14_DC28(p0, (self).f21, d_1173_v136_)
                d_1175_v138_: _dafny.Map
                d_1175_v138_ = _dafny.Map({d_1169_v134_: d_1174_v137_})
                d_1175_v138_ = (d_1175_v138_).set(d_1169_v134_, d_1174_v137_)
                rhs212_ = (d_1046_v61_).set(d_1155_cf3_, default__.abs(default__.safeModuloInt(len(_dafny.Map({(d_969_v4_)[default__.safeIndex(d_1173_v136_.f14, len(d_969_v4_))]: (self).f17})), d_1173_v136_.f14)))
                rhs213_ = not(not(d_1155_cf3_))
                rhs214_ = d_1173_v136_.f14
                lhs166_ = self
                lhs167_ = globalState
                d_1046_v61_ = rhs212_
                lhs166_.f16 = rhs213_
                lhs167_.f2 = rhs214_
                (self).f16 = False
                (globalState).f2 = (default__.safeDivisionInt(d_1173_v136_.f14, p0)) * ((0) - (default__.fm0(self.f16, globalState)))
                d_1049_v64_ = d_1049_v64_
            d_1176_v139_: _dafny.Map
            d_1176_v139_ = _dafny.Map({d_1155_cf3_: default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))))})
            d_1176_v139_ = (d_1176_v139_).set(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))))
            d_1177_v140_: T0
            nw245_ = C1()
            nw245_.ctor__(p1, p0)
            d_1177_v140_ = nw245_
            d_1178_v141_: D16
            d_1178_v141_ = D16_DC33(d_1155_cf3_, (self).f17, len(d_1048_v63_), (self).f17)
            d_1179_v142_: D16
            d_1179_v142_ = D16_DC34((self).f17, d_1177_v140_, (d_1178_v141_).cf49)
            d_1180_v143_: _dafny.MultiSet
            d_1180_v143_ = _dafny.MultiSet([d_1179_v142_])
            d_1181_v144_: _dafny.Map
            d_1181_v144_ = _dafny.Map({not(p1): d_1180_v143_})
            d_1182_v145_: _dafny.MultiSet
            d_1182_v145_ = _dafny.MultiSet([p0, p0, 913])
            d_1181_v144_ = (d_1181_v144_).set((d_1182_v145_).issubset(d_1182_v145_), d_1180_v143_)
            d_1183_v146_: _dafny.Map
            d_1183_v146_ = _dafny.Map({d_1155_cf3_: True})
            d_1183_v146_ = (d_1183_v146_).set((self).f21, (self).f21)
        elif True:
            d_1184___mcc_h20_ = source17_.cf9
            d_1185_cf9_ = d_1184___mcc_h20_
            d_1186_v147_: _dafny.Array
            def lambda72_(d_1187_p0_):
                def lambda73_(d_1188_i15_):
                    return (d_1188_i15_) + (d_1187_p0_)

                return lambda73_

            init41_ = lambda72_(p0)
            nw246_ = _dafny.Array(None, 25)
            for i0_41_ in range(nw246_.length(0)):
                nw246_[i0_41_] = init41_(i0_41_)
            d_1186_v147_ = nw246_
            d_1189_v148_: _dafny.Seq
            d_1189_v148_ = _dafny.SeqWithoutIsStrInference([d_1186_v147_, d_1186_v147_, d_1186_v147_])
            (self).f16 = (d_1189_v148_) <= (d_1189_v148_)
            d_1190_v149_: str
            d_1190_v149_ = _dafny.CodePoint('j')
            d_1190_v149_ = d_1190_v149_
            (globalState).f7 = p0
            d_1191_v150_: _dafny.Seq
            d_1191_v150_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1192_v151_: _dafny.Map
            d_1192_v151_ = _dafny.Map({self.f16: len(_dafny.Set({default__.fm0((self).f21, globalState), p0, p0}))})
            d_1193_v152_: _dafny.Map
            d_1193_v152_ = _dafny.Map({len(_dafny.Set({d_1191_v150_, _dafny.SeqWithoutIsStrInference([len(d_1192_v151_), len(d_1048_v63_)])})): _dafny.CodePoint('x')})
            (self).f16 = ((p0) + (len(d_1193_v152_))) != (p0)
        r0 = 536
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_1194_v0_: _dafny.Set
        d_1194_v0_ = _dafny.Set({p3, p1, p3, p3})
        d_1195_v1_: _dafny.Seq
        d_1195_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgfxhsafg"))
        d_1196_v2_: _dafny.Seq
        d_1196_v2_ = _dafny.SeqWithoutIsStrInference([len(d_1195_v1_), p3])
        d_1197_v3_: D3
        d_1197_v3_ = D3_DC4((self).f21, len(d_1195_v1_), p1, (self).f17, not(default__.fm43((self).f17, globalState)))
        d_1198_v4_: _dafny.Array
        nw247_ = _dafny.Array(None, 13)
        nw247_[int(0)] = p1
        nw247_[int(1)] = p3
        nw247_[int(2)] = (0) - (p3)
        nw247_[int(3)] = len(d_1194_v0_)
        nw247_[int(4)] = (0) - ((self).fm12(p1, globalState))
        nw247_[int(5)] = (_dafny.MultiSet(d_1196_v2_)).cardinality
        nw247_[int(6)] = len(d_1195_v1_)
        nw247_[int(7)] = p1
        nw247_[int(8)] = p3
        nw247_[int(9)] = p1
        nw247_[int(10)] = p3
        nw247_[int(11)] = (d_1197_v3_).cf5
        nw247_[int(12)] = 84
        d_1198_v4_ = nw247_
        d_1199_v5_: _dafny.Map
        d_1199_v5_ = _dafny.Map({d_1198_v4_: (p3) < (len(default__.fm49(p3, p3, (self).f21, p1, globalState)))})
        d_1199_v5_ = (d_1199_v5_) | (d_1199_v5_)
        d_1200_v6_: _dafny.Set
        d_1200_v6_ = _dafny.Set({p2, p2, (self).f17, p2, (self).f17})
        d_1201_v7_: _dafny.Map
        d_1201_v7_ = _dafny.Map({(self).f17: 400})
        d_1202_v8_: _dafny.Map
        d_1202_v8_ = _dafny.Map({d_1200_v6_: (d_1201_v7_) | (d_1201_v7_)})
        d_1202_v8_ = (d_1202_v8_).set(d_1200_v6_, d_1201_v7_)
        d_1203_v9_: _dafny.Map
        d_1203_v9_ = _dafny.Map({p3: p1})
        (self).f16 = ((len(d_1195_v1_)) + (len(default__.fm38((self).f21, len(d_1203_v9_), p3, globalState)))) == (352)
        d_1204_v10_: D3
        d_1204_v10_ = D3_DC3(True)
        if (d_1204_v10_).cf3:
            (self).f16 = (self).f21
            d_1205_v11_: _dafny.Map
            d_1205_v11_ = _dafny.Map({d_1195_v1_: self.f16})
            d_1205_v11_ = (d_1205_v11_).set(d_1195_v1_, False)
            (self).f16 = ((self).f17 if self.f16 else not(False))
            d_1206_v12_: _dafny.Seq
            d_1206_v12_ = _dafny.SeqWithoutIsStrInference([self.f16])
            (self).f16 = (d_1206_v12_)[default__.safeIndex(p3, len(d_1206_v12_))]
            d_1207_v13_: _dafny.MultiSet
            d_1207_v13_ = _dafny.MultiSet([p1])
            d_1208_v14_: _dafny.Map
            d_1208_v14_ = _dafny.Map({p3: d_1207_v13_})
            d_1208_v14_ = (d_1208_v14_).set(p3, d_1207_v13_)
        elif True:
            (self).f16 = (self).f17
            d_1209_v15_: _dafny.Map
            d_1209_v15_ = _dafny.Map({True: p0})
            d_1209_v15_ = (d_1209_v15_).set(self.f16, p0)
            d_1210_v16_: _dafny.Seq
            d_1210_v16_ = _dafny.SeqWithoutIsStrInference([default__.fm27(p3, p1, p1, (d_1203_v9_).set(p1, p1), globalState)])
            r0 = (p3) - (len((d_1210_v16_) + (d_1210_v16_)))
            d_1211_v17_: T0
            nw248_ = C1()
            nw248_.ctor__((self).f21, p1)
            d_1211_v17_ = nw248_
            d_1212_v18_: D14
            d_1212_v18_ = D14_DC28(p1, (self).f17, d_1211_v17_)
            index210_ = default__.safeIndex(322, (d_1198_v4_).length(0))
            (d_1198_v4_)[index210_] = (0) - ((d_1212_v18_).cf40)
            index211_ = default__.safeIndex(322, (d_1198_v4_).length(0))
            (d_1198_v4_)[index211_] = (0) - (-460)
            d_1213_v19_: _dafny.Map
            d_1213_v19_ = _dafny.Map({d_1201_v7_: p2})
            d_1214_v20_: _dafny.Seq
            d_1214_v20_ = _dafny.SeqWithoutIsStrInference([d_1195_v1_, d_1195_v1_])
            d_1215_v21_: _dafny.Array
            nw249_ = _dafny.Array(int(0), 10)
            d_1215_v21_ = nw249_
            d_1216_v22_: _dafny.Seq
            d_1216_v22_ = _dafny.SeqWithoutIsStrInference([d_1215_v21_, d_1215_v21_])
            d_1217_v23_: _dafny.Array
            nw250_ = _dafny.Array(None, 18)
            nw250_[int(0)] = (d_1210_v16_)[default__.safeIndex(len(d_1210_v16_), len(d_1210_v16_))]
            nw250_[int(1)] = self.f16
            nw250_[int(2)] = not((self).f21)
            nw250_[int(3)] = self.f16
            nw250_[int(4)] = (self).f21
            nw250_[int(5)] = (self).f21
            nw250_[int(6)] = p2
            nw250_[int(7)] = p2
            nw250_[int(8)] = (self).f17
            nw250_[int(9)] = ((self).f21 if True else ((d_1213_v19_)[d_1201_v7_] if (d_1201_v7_) in (d_1213_v19_) else False))
            nw250_[int(10)] = self.f16
            nw250_[int(11)] = self.f16
            nw250_[int(12)] = p2
            nw250_[int(13)] = (((d_1214_v20_)[default__.safeIndex(d_1211_v17_.f14, len(d_1214_v20_))]).set(default__.safeIndex(len(d_1210_v16_), len((d_1214_v20_)[default__.safeIndex(d_1211_v17_.f14, len(d_1214_v20_))])), p0)) < (d_1195_v1_)
            nw250_[int(14)] = (d_1216_v22_) == (d_1216_v22_)
            nw250_[int(15)] = self.f16
            nw250_[int(16)] = ((self).f17) == ((d_1210_v16_)[default__.safeIndex(p3, len(d_1210_v16_))])
            nw250_[int(17)] = (self).f17
            d_1217_v23_ = nw250_
            nw251_ = _dafny.Array(False, 16)
            d_1217_v23_ = nw251_
        d_1218_v24_: _dafny.Array
        nw252_ = _dafny.Array(_dafny.Set({}), 15)
        d_1218_v24_ = nw252_
        d_1219_v25_: _dafny.Seq
        d_1219_v25_ = _dafny.SeqWithoutIsStrInference([False, self.f16])
        d_1220_v26_: C0
        nw253_ = C0()
        nw253_.ctor__(d_1218_v24_, len(d_1219_v25_))
        d_1220_v26_ = nw253_
        if not(not((self).f17)):
            (self).f16 = self.f16
            d_1221_v27_: _dafny.Map
            d_1221_v27_ = _dafny.Map({d_1198_v4_: (0) - (p1)})
            d_1222_v28_: _dafny.MultiSet
            d_1222_v28_ = _dafny.MultiSet([d_1195_v1_])
            d_1223_v29_: _dafny.Map
            d_1223_v29_ = _dafny.Map({d_1222_v28_: d_1221_v27_})
            d_1224_v30_: _dafny.Map
            d_1224_v30_ = _dafny.Map({p0: d_1221_v27_})
            d_1225_v31_: _dafny.Array
            nw254_ = _dafny.Array(None, 8)
            nw254_[int(0)] = d_1221_v27_
            nw254_[int(1)] = d_1221_v27_
            nw254_[int(2)] = (_dafny.Map({d_1198_v4_: p3})) | (d_1221_v27_)
            nw254_[int(3)] = d_1221_v27_
            nw254_[int(4)] = (d_1221_v27_) | (d_1221_v27_)
            nw254_[int(5)] = ((d_1223_v29_)[_dafny.MultiSet([d_1195_v1_])] if (_dafny.MultiSet([d_1195_v1_])) in (d_1223_v29_) else ((d_1224_v30_)[p0] if (p0) in (d_1224_v30_) else d_1221_v27_))
            nw254_[int(6)] = d_1221_v27_
            nw254_[int(7)] = d_1221_v27_
            d_1225_v31_ = nw254_
            index212_ = default__.safeIndex(502, (d_1225_v31_).length(0))
            (d_1225_v31_)[index212_] = d_1221_v27_
            index213_ = default__.safeIndex(502, (d_1225_v31_).length(0))
            (d_1225_v31_)[index213_] = d_1221_v27_
            d_1226_v32_: _dafny.Array
            nw255_ = _dafny.Array(False, 21)
            d_1226_v32_ = nw255_
            d_1226_v32_ = (d_1226_v32_ if (d_1194_v0_).isdisjoint(default__.fm49(len(d_1194_v0_), p1, not(not((self).f17)), d_1220_v26_.f27, globalState)) else d_1226_v32_)
            (globalState).f7 = 387
            d_1226_v32_ = d_1226_v32_
        elif True:
            (self).f16 = (p1) >= (d_1220_v26_.f27)
            d_1227_v33_: _dafny.Seq
            d_1227_v33_ = _dafny.SeqWithoutIsStrInference([d_1219_v25_])
            d_1228_v34_: T1
            nw256_ = C3()
            nw256_.ctor__((_dafny.MultiSet((d_1227_v33_)[default__.safeIndex(p3, len(d_1227_v33_))])).cardinality, d_1195_v1_, (self).f21, (self).f21)
            d_1228_v34_ = nw256_
            d_1229_v35_: _dafny.Map
            d_1229_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0 for d_1230_i0_ in range(default__.abs(430))]): d_1228_v34_})
            d_1231_v36_: _dafny.Array
            nw257_ = _dafny.Array(None, 21)
            nw257_[int(0)] = d_1229_v35_
            nw257_[int(1)] = d_1229_v35_
            nw257_[int(2)] = (d_1229_v35_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")): d_1228_v34_}))
            nw257_[int(3)] = d_1229_v35_
            nw257_[int(4)] = d_1229_v35_
            nw257_[int(5)] = (d_1229_v35_) | (_dafny.Map({d_1195_v1_: d_1228_v34_}))
            nw257_[int(6)] = (d_1229_v35_).set(d_1195_v1_, d_1228_v34_)
            nw257_[int(7)] = (d_1229_v35_) | (d_1229_v35_)
            nw257_[int(8)] = (d_1229_v35_) | (_dafny.Map({d_1195_v1_: d_1228_v34_}))
            nw257_[int(9)] = d_1229_v35_
            nw257_[int(10)] = d_1229_v35_
            nw257_[int(11)] = d_1229_v35_
            nw257_[int(12)] = d_1229_v35_
            nw257_[int(13)] = (d_1229_v35_) | (_dafny.Map({d_1195_v1_: d_1228_v34_}))
            nw257_[int(14)] = d_1229_v35_
            nw257_[int(15)] = _dafny.Map({d_1195_v1_: d_1228_v34_})
            nw257_[int(16)] = d_1229_v35_
            nw257_[int(17)] = d_1229_v35_
            nw257_[int(18)] = (d_1229_v35_ if p2 else d_1229_v35_)
            nw257_[int(19)] = d_1229_v35_
            nw257_[int(20)] = d_1229_v35_
            d_1231_v36_ = nw257_
            index214_ = default__.safeIndex(612, (d_1231_v36_).length(0))
            (d_1231_v36_)[index214_] = d_1229_v35_
            index215_ = default__.safeIndex(612, (d_1231_v36_).length(0))
            (d_1231_v36_)[index215_] = d_1229_v35_
            d_1232_v37_: D15
            d_1232_v37_ = D15_DC30(p0)
            d_1233_v38_: D15
            d_1233_v38_ = D15_DC31(d_1232_v37_)
            d_1233_v38_ = d_1233_v38_
            (globalState).f2 = d_1220_v26_.f27
            d_1234_v39_: _dafny.Map
            d_1234_v39_ = _dafny.Map({(self).f21: d_1195_v1_})
            d_1195_v1_ = ((d_1234_v39_)[p2] if (p2) in (d_1234_v39_) else d_1195_v1_)
        d_1235_v40_: _dafny.MultiSet
        d_1235_v40_ = _dafny.MultiSet([((self).f17) == ((self).f17), not((self).f21)])
        d_1236_v41_: D15
        d_1236_v41_ = D15_DC29(d_1200_v6_)
        r0 = ((d_1235_v40_)[(864) > (len(d_1200_v6_))] if ((864) > (len(d_1200_v6_))) in (d_1235_v40_) else (len((d_1236_v41_).cf43) if default__.fm43((self).f21, globalState) else p3))
        return r0

    def m10(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: D6 = D6.default()()
        r2: bool = False
        d_1237_v0_: _dafny.Seq
        d_1237_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbxik"))
        d_1238_v1_: _dafny.Seq
        d_1238_v1_ = _dafny.SeqWithoutIsStrInference([p2, len(d_1237_v0_), p2, p1, p1])
        (globalState).f7 = ((d_1238_v1_)[default__.safeIndex(p2, len(d_1238_v1_))]) * (p1)
        d_1239_i0_: int
        d_1239_i0_ = 0
        with _dafny.label("5"):
            while (p1) >= (default__.fm0(p3, globalState)):
                with _dafny.c_label("5"):
                    if (d_1239_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1239_i0_ = (d_1239_i0_) + (1)
                    d_1240_v2_: _dafny.Map
                    d_1240_v2_ = _dafny.Map({(self).f17: p0})
                    d_1241_v3_: _dafny.Map
                    d_1241_v3_ = _dafny.Map({(self).f21: d_1240_v2_})
                    d_1242_v4_: _dafny.Map
                    d_1242_v4_ = _dafny.Map({d_1237_v0_: (0) - (len(d_1241_v3_))})
                    d_1243_v5_: D13
                    d_1243_v5_ = D13_DC25(d_1242_v4_)
                    pat_let_tv17_ = d_1242_v4_
                    def iife91_(_pat_let21_0):
                        def iife92_(d_1244_dt__update__tmp_h0_):
                            def iife93_(_pat_let22_0):
                                def iife94_(d_1245_dt__update_hcf38_h0_):
                                    return D13_DC25(d_1245_dt__update_hcf38_h0_)
                                return iife94_(_pat_let22_0)
                            return iife93_(pat_let_tv17_)
                        return iife92_(_pat_let21_0)
                    source18_ = (d_1243_v5_ if (p2) >= (p2) else iife91_(d_1243_v5_))
                    if source18_.is_DC26:
                        d_1246_v6_: D3
                        d_1246_v6_ = D3_DC3((self).f21)
                        d_1247_v7_: C2
                        nw258_ = C2()
                        nw258_.ctor__(d_1246_v6_, p0, p0)
                        d_1247_v7_ = nw258_
                        d_1238_v1_ = d_1238_v1_
                        (self).f16 = (not(False)) == ((self).f17)
                        d_1248_v8_: _dafny.Array
                        nw259_ = _dafny.Array(False, 9)
                        d_1248_v8_ = nw259_
                        d_1249_v9_: _dafny.Set
                        d_1249_v9_ = _dafny.Set({19})
                        d_1250_v10_: _dafny.Map
                        d_1250_v10_ = _dafny.Map({d_1248_v8_: (d_1249_v9_).intersection(_dafny.Set({(D9_DC16(len(d_1240_v2_), not((self).f21))).cf23}))})
                        rhs215_ = _dafny.Map({d_1248_v8_: d_1249_v9_})
                        rhs216_ = (self).f21
                        rhs217_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlvv"))
                        rhs218_ = p1
                        rhs219_ = True
                        lhs168_ = globalState
                        d_1250_v10_ = rhs215_
                        r2 = rhs216_
                        d_1237_v0_ = rhs217_
                        lhs168_.f7 = rhs218_
                        r0 = rhs219_
                    elif True:
                        d_1251___mcc_h0_ = source18_.cf38
                        d_1252_cf38_ = d_1251___mcc_h0_
                        (globalState).f7 = (p1) * ((p1 if (self).f17 else p1))
                        d_1253_v11_: str
                        d_1253_v11_ = _dafny.CodePoint('q')
                        d_1254_v12_: _dafny.Array
                        nw260_ = _dafny.Array(None, 22)
                        nw260_[int(0)] = d_1253_v11_
                        nw260_[int(1)] = (d_1237_v0_)[default__.safeIndex((0) - (p2), len(d_1237_v0_))]
                        nw260_[int(2)] = d_1253_v11_
                        nw260_[int(3)] = d_1253_v11_
                        nw260_[int(4)] = d_1253_v11_
                        nw260_[int(5)] = d_1253_v11_
                        nw260_[int(6)] = d_1253_v11_
                        nw260_[int(7)] = d_1253_v11_
                        nw260_[int(8)] = d_1253_v11_
                        nw260_[int(9)] = d_1253_v11_
                        nw260_[int(10)] = d_1253_v11_
                        nw260_[int(11)] = d_1253_v11_
                        nw260_[int(12)] = d_1253_v11_
                        nw260_[int(13)] = d_1253_v11_
                        nw260_[int(14)] = d_1253_v11_
                        nw260_[int(15)] = d_1253_v11_
                        nw260_[int(16)] = d_1253_v11_
                        nw260_[int(17)] = d_1253_v11_
                        nw260_[int(18)] = d_1253_v11_
                        nw260_[int(19)] = d_1253_v11_
                        nw260_[int(20)] = d_1253_v11_
                        nw260_[int(21)] = d_1253_v11_
                        d_1254_v12_ = nw260_
                        d_1254_v12_ = d_1254_v12_
                        (globalState).f7 = p1
                        d_1255_v13_: _dafny.Seq
                        d_1255_v13_ = _dafny.SeqWithoutIsStrInference([p3])
                        (globalState).f2 = (p1) * (default__.safeDivisionInt(len(d_1255_v13_), (0) - (p1)))
                    d_1256_v14_: str
                    d_1256_v14_ = _dafny.CodePoint('p')
                    d_1257_v15_: _dafny.Map
                    d_1257_v15_ = _dafny.Map({_dafny.MultiSet([d_1256_v14_]): (self).f17})
                    d_1258_v16_: _dafny.Map
                    d_1258_v16_ = _dafny.Map({d_1257_v15_: _dafny.Map({p3: False})})
                    d_1259_v17_: _dafny.MultiSet
                    d_1259_v17_ = _dafny.MultiSet([d_1256_v14_, _dafny.CodePoint('w'), d_1256_v14_])
                    d_1258_v16_ = (d_1258_v16_).set((d_1257_v15_).set(d_1259_v17_, self.f16), d_1240_v2_)
                    (globalState).f2 = p2
                    (globalState).f7 = len(_dafny.Map({(p1 if p3 else p1): p1}))
                    pass
            pass
        d_1260_v18_: _dafny.Set
        d_1260_v18_ = _dafny.Set({p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgey")))})
        d_1261_v21_: _dafny.Seq
        d_1261_v21_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1262_v22_: _dafny.Array
        nw261_ = _dafny.Array(None, 14)
        nw261_[int(0)] = p1
        nw261_[int(1)] = len(d_1261_v21_)
        nw261_[int(2)] = p1
        nw261_[int(3)] = p1
        nw261_[int(4)] = len(d_1261_v21_)
        nw261_[int(5)] = -92
        nw261_[int(6)] = 12
        nw261_[int(7)] = (0) - (p2)
        nw261_[int(8)] = 170
        nw261_[int(9)] = 903
        nw261_[int(10)] = p1
        nw261_[int(11)] = len(d_1238_v1_)
        nw261_[int(12)] = p2
        nw261_[int(13)] = p2
        d_1262_v22_ = nw261_
        d_1263_v23_: _dafny.Array
        nw262_ = _dafny.Array(False, 27)
        d_1263_v23_ = nw262_
        d_1264_v24_: D9
        def iife95_():
            coll49_ = _dafny.Set()
            compr_49_: int
            for compr_49_ in _dafny.IntegerRange(413, -815):
                d_1265_v20_: int = compr_49_
                if ((413) <= (d_1265_v20_)) and ((d_1265_v20_) < (-815)):
                    coll49_ = coll49_.union(_dafny.Set([(d_1265_v20_) - (p1)]))
            return _dafny.Set(coll49_)
        d_1264_v24_ = D9_DC15(False, (self).f21, iife95_()
, d_1262_v22_, d_1263_v23_)
        d_1266_v25_: _dafny.Map
        d_1266_v25_ = _dafny.Map({p1: d_1263_v23_})
        d_1267_v26_: D9
        d_1267_v26_ = D9_DC15(False, self.f16, (d_1264_v24_).cf20, d_1262_v22_, ((d_1266_v25_)[p1] if (p1) in (d_1266_v25_) else d_1263_v23_))
        d_1268_v28_: _dafny.Array
        nw263_ = _dafny.Array(None, 10)
        def iife96_():
            coll50_ = _dafny.Set()
            compr_50_: int
            for compr_50_ in (_dafny.SeqWithoutIsStrInference([p2])).Elements:
                d_1269_v19_: int = compr_50_
                if (d_1269_v19_) in (_dafny.SeqWithoutIsStrInference([p2])):
                    coll50_ = coll50_.union(_dafny.Set([(d_1269_v19_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aracab"))))]))
            return _dafny.Set(coll50_)
        nw263_[int(0)] = (d_1260_v18_).intersection(iife96_()
        )
        def iife97_():
            coll51_ = _dafny.Set()
            compr_51_: int
            for compr_51_ in _dafny.IntegerRange(474, 449):
                d_1270_v27_: int = compr_51_
                if ((474) <= (d_1270_v27_)) and ((d_1270_v27_) < (449)):
                    coll51_ = coll51_.union(_dafny.Set([(d_1270_v27_) + (p2)]))
            return _dafny.Set(coll51_)
        nw263_[int(1)] = ((d_1264_v24_).cf20) | (iife97_()
        )
        nw263_[int(2)] = default__.fm49(767, p2, self.f16, (0) - (p1), globalState)
        nw263_[int(3)] = d_1260_v18_
        nw263_[int(4)] = (d_1260_v18_) - (d_1260_v18_)
        nw263_[int(5)] = d_1260_v18_
        nw263_[int(6)] = (d_1260_v18_) | (d_1260_v18_)
        nw263_[int(7)] = d_1260_v18_
        nw263_[int(8)] = (d_1260_v18_ if (self).f21 else d_1260_v18_)
        nw263_[int(9)] = d_1260_v18_
        d_1268_v28_ = nw263_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1268_v28_).length(0)):
            d_1271_i1_: int = guard_loop_6_
            if (True) and (((0) <= (d_1271_i1_)) and ((d_1271_i1_) < ((d_1268_v28_).length(0)))):
                (d_1268_v28_)[(d_1271_i1_)] = (d_1260_v18_) - (d_1260_v18_)
        d_1272_v29_: _dafny.MultiSet
        d_1272_v29_ = _dafny.MultiSet([(self).f21])
        d_1273_v30_: D16
        d_1273_v30_ = D16_DC32((d_1272_v29_).intersection(d_1272_v29_))
        source19_ = d_1273_v30_
        if source19_.is_DC33:
            d_1274___mcc_h1_ = source19_.cf47
            d_1275___mcc_h2_ = source19_.cf48
            d_1276___mcc_h3_ = source19_.cf49
            d_1277___mcc_h4_ = source19_.cf50
            d_1278_cf50_ = d_1277___mcc_h4_
            d_1279_cf49_ = d_1276___mcc_h3_
            d_1280_cf48_ = d_1275___mcc_h2_
            d_1281_cf47_ = d_1274___mcc_h1_
            (globalState).f7 = p2
            d_1282_v31_: str
            d_1282_v31_ = _dafny.CodePoint('i')
            d_1283_v32_: _dafny.Map
            d_1283_v32_ = _dafny.Map({d_1278_cf50_: d_1262_v22_})
            rhs220_ = d_1282_v31_
            rhs221_ = d_1279_cf49_
            rhs222_ = (d_1283_v32_).set(p3, d_1262_v22_)
            lhs169_ = globalState
            d_1282_v31_ = rhs220_
            lhs169_.f7 = rhs221_
            d_1283_v32_ = rhs222_
            d_1284_v33_: _dafny.Array
            nw264_ = _dafny.Array(_dafny.CodePoint('D'), 9)
            d_1284_v33_ = nw264_
            index216_ = default__.safeIndex(751, (d_1284_v33_).length(0))
            (d_1284_v33_)[index216_] = (d_1282_v31_ if True else d_1282_v31_)
            index217_ = default__.safeIndex(751, (d_1284_v33_).length(0))
            (d_1284_v33_)[index217_] = d_1282_v31_
            (globalState).f7 = (d_1279_cf49_) * ((len(default__.fm52(d_1281_cf47_, len(d_1260_v18_), d_1279_cf49_, p2, globalState)) if (self).f17 else p1))
        elif source19_.is_DC34:
            d_1285___mcc_h5_ = source19_.cf51
            d_1286___mcc_h6_ = source19_.cf52
            d_1287___mcc_h7_ = source19_.cf53
            d_1288_cf53_ = d_1287___mcc_h7_
            d_1289_cf52_ = d_1286___mcc_h6_
            d_1290_cf51_ = d_1285___mcc_h5_
            index218_ = default__.safeIndex(258, (d_1263_v23_).length(0))
            (d_1263_v23_)[index218_] = d_1290_cf51_
            d_1291_v34_: _dafny.Array
            def lambda74_(d_1292_i2_):
                return _dafny.CodePoint('y')

            init42_ = lambda74_
            nw265_ = _dafny.Array(None, 2)
            for i0_42_ in range(nw265_.length(0)):
                nw265_[i0_42_] = init42_(i0_42_)
            d_1291_v34_ = nw265_
            d_1293_v35_: _dafny.Map
            d_1293_v35_ = _dafny.Map({d_1290_cf51_: d_1291_v34_})
            d_1294_v36_: _dafny.Map
            d_1294_v36_ = _dafny.Map({(self).f17: ((d_1293_v35_)[(self).f21] if ((self).f21) in (d_1293_v35_) else d_1291_v34_)})
            d_1295_v37_: _dafny.Map
            d_1295_v37_ = _dafny.Map({(self).f17: p1})
            index219_ = default__.safeIndex(258, (d_1263_v23_).length(0))
            rhs223_ = ((len(d_1293_v35_)) - (p1)) - (len(d_1261_v21_))
            rhs224_ = default__.fm43((p3) not in (d_1295_v37_), globalState)
            lhs170_ = globalState
            lhs171_ = d_1263_v23_
            lhs172_ = default__.safeIndex(258, (d_1263_v23_).length(0))
            lhs170_.f7 = rhs223_
            lhs171_[lhs172_] = rhs224_
            d_1296_v38_: _dafny.Map
            d_1296_v38_ = _dafny.Map({d_1289_cf52_.f14: len(d_1237_v0_)})
            (d_1289_cf52_).f14 = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f21, True])), ((d_1296_v38_)[d_1288_cf53_] if (d_1288_cf53_) in (d_1296_v38_) else len(d_1238_v1_)))
            d_1297_v39_: _dafny.Array
            nw266_ = _dafny.Array(_dafny.Set({}), 15)
            d_1297_v39_ = nw266_
            d_1298_v40_: C0
            nw267_ = C0()
            nw267_.ctor__(d_1297_v39_, p2)
            d_1298_v40_ = nw267_
            (self).f16 = not((d_1272_v29_).ispropersubset((d_1273_v30_).cf46))
        elif source19_.is_DC32:
            d_1299___mcc_h8_ = source19_.cf46
            d_1300_cf46_ = d_1299___mcc_h8_
            if (self).f17:
                r0 = p0
                index220_ = default__.safeIndex(604, (d_1262_v22_).length(0))
                (d_1262_v22_)[index220_] = p1
                index221_ = default__.safeIndex(604, (d_1262_v22_).length(0))
                (d_1262_v22_)[index221_] = (0) - (default__.safeModuloInt(p1, p2))
                d_1301_v41_: _dafny.Array
                nw268_ = _dafny.Array(int(0), 11)
                d_1301_v41_ = nw268_
                d_1302_v42_: _dafny.Map
                d_1302_v42_ = _dafny.Map({d_1301_v41_: p0})
                d_1303_v43_: _dafny.MultiSet
                d_1303_v43_ = _dafny.MultiSet([d_1301_v41_, d_1301_v41_])
                index222_ = default__.safeIndex(604, (d_1262_v22_).length(0))
                rhs225_ = d_1267_v26_
                rhs226_ = (self).f17
                rhs227_ = (d_1303_v43_).ispropersubset(d_1303_v43_)
                rhs228_ = p2
                rhs229_ = d_1302_v42_
                lhs173_ = d_1262_v22_
                lhs174_ = default__.safeIndex(604, (d_1262_v22_).length(0))
                d_1264_v24_ = rhs225_
                r0 = rhs226_
                r0 = rhs227_
                lhs173_[lhs174_] = rhs228_
                d_1302_v42_ = rhs229_
                (globalState).f7 = (0) - (((d_1262_v22_)[default__.safeIndex(604, (d_1262_v22_).length(0))]) * (p2))
                d_1304_v45_: _dafny.Seq
                def iife98_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(370, -776):
                        d_1305_v44_: int = compr_52_
                        if ((370) <= (d_1305_v44_)) and ((d_1305_v44_) < (-776)):
                            coll52_ = coll52_.union(_dafny.Set([(d_1305_v44_) + (p2)]))
                    return _dafny.Set(coll52_)
                d_1304_v45_ = _dafny.SeqWithoutIsStrInference([iife98_()
                , (d_1260_v18_).intersection(_dafny.Set({p2, p2, len(d_1237_v0_)}))])
                d_1260_v18_ = (d_1304_v45_)[default__.safeIndex((d_1262_v22_)[default__.safeIndex(604, (d_1262_v22_).length(0))], len(d_1304_v45_))]
            elif True:
                d_1306_v46_: D12
                d_1306_v46_ = D12_DC23(p0, p2, 182)
                rhs230_ = ((d_1306_v46_).cf32) == ((d_1238_v1_)[default__.safeIndex(p1, len(d_1238_v1_))])
                rhs231_ = (d_1261_v21_)[default__.safeIndex(p2, len(d_1261_v21_))]
                r0 = rhs230_
                r0 = rhs231_
                r2 = p3
                d_1307_v47_: D15
                d_1307_v47_ = D15_DC30(_dafny.CodePoint('m'))
                d_1308_v48_: str
                d_1308_v48_ = _dafny.CodePoint('b')
                def iife99_(_pat_let23_0):
                    def iife100_(d_1309_dt__update__tmp_h1_):
                        def iife101_(_pat_let24_0):
                            def iife102_(d_1310_dt__update_hcf44_h0_):
                                return D15_DC30(d_1310_dt__update_hcf44_h0_)
                            return iife102_(_pat_let24_0)
                        return iife101_(_dafny.CodePoint('m'))
                    return iife100_(_pat_let23_0)
                d_1307_v47_ = iife99_(D15_DC30(d_1308_v48_))
                d_1311_v49_: _dafny.Map
                d_1311_v49_ = _dafny.Map({len(d_1266_v25_): (d_1261_v21_)[default__.safeIndex(p2, len(d_1261_v21_))]})
                d_1311_v49_ = (d_1311_v49_).set(p2, not((self).f21))
                d_1312_v50_: _dafny.Map
                d_1312_v50_ = _dafny.Map({p2: p1})
                d_1313_v52_: _dafny.Map
                def iife103_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(549, -902):
                        d_1314_v51_: int = compr_53_
                        if ((549) <= (d_1314_v51_)) and ((d_1314_v51_) < (-902)):
                            coll53_[(d_1314_v51_) + (p1)] = p2
                    return _dafny.Map(coll53_)
                d_1313_v52_ = _dafny.Map({(d_1312_v50_) | (iife103_()
                ): default__.fm31(not(True), not(p3), globalState)})
                d_1313_v52_ = d_1313_v52_
            (globalState).f7 = p1
            d_1315_v53_: _dafny.Map
            d_1315_v53_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwnthagln"))]): d_1238_v1_})
            d_1238_v1_ = ((d_1315_v53_)[_dafny.SeqWithoutIsStrInference([d_1237_v0_])] if (_dafny.SeqWithoutIsStrInference([d_1237_v0_])) in (d_1315_v53_) else default__.fm35(p1, True, len(d_1237_v0_), _dafny.CodePoint('m'), globalState))
            d_1316_v54_: _dafny.Map
            d_1316_v54_ = _dafny.Map({default__.safeModuloInt(p2, p2): p2})
            d_1317_v55_: _dafny.MultiSet
            d_1317_v55_ = _dafny.MultiSet([p2])
            d_1316_v54_ = ((_dafny.Map({p1: p2})).set(p1, p2)).set(p2, ((d_1317_v55_)[p1] if (p1) in (d_1317_v55_) else p1))
        elif True:
            d_1318___mcc_h9_ = source19_.cf54
            d_1319_cf54_ = d_1318___mcc_h9_
            index223_ = default__.safeIndex(355, (d_1262_v22_).length(0))
            (d_1262_v22_)[index223_] = p1
            index224_ = default__.safeIndex(355, (d_1262_v22_).length(0))
            (d_1262_v22_)[index224_] = 952
            (self).f16 = not((self).f21)
            d_1237_v0_ = (D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tncsoxvl")))).cf29
            source20_ = default__.fm56(p1, (p2) * (p1), globalState)
            if source20_.is_DC28:
                d_1320___mcc_h10_ = source20_.cf40
                d_1321___mcc_h11_ = source20_.cf41
                d_1322___mcc_h12_ = source20_.cf42
                d_1323_cf42_ = d_1322___mcc_h12_
                d_1324_cf41_ = d_1321___mcc_h11_
                d_1325_cf40_ = d_1320___mcc_h10_
                r2 = (p2) != (d_1323_cf42_.f14)
                d_1326_v56_: _dafny.MultiSet
                d_1326_v56_ = _dafny.MultiSet([d_1263_v23_, d_1263_v23_])
                r2 = ((d_1326_v56_) | (d_1326_v56_)) != (_dafny.MultiSet([d_1263_v23_]))
                r0 = self.f16
                d_1327_v57_: _dafny.Map
                d_1327_v57_ = _dafny.Map({p0: d_1324_cf41_})
                d_1328_v58_: _dafny.Set
                d_1328_v58_ = _dafny.Set({d_1327_v57_})
                d_1329_v60_: str
                d_1329_v60_ = _dafny.CodePoint('b')
                d_1330_v61_: D6
                d_1330_v61_ = D6_DC9(d_1329_v60_)
                def iife104_():
                    coll54_ = _dafny.Set()
                    compr_54_: _dafny.Map
                    for compr_54_ in (_dafny.Map({d_1327_v57_: d_1323_cf42_.f14})).keys.Elements:
                        d_1331_v59_: _dafny.Map = compr_54_
                        if (d_1331_v59_) in (_dafny.Map({d_1327_v57_: d_1323_cf42_.f14})):
                            coll54_ = coll54_.union(_dafny.Set([d_1331_v59_]))
                    return _dafny.Set(coll54_)
                rhs232_ = (iife104_()
                ) | (default__.fm57(globalState))
                rhs233_ = p3
                rhs234_ = (self).f21
                rhs235_ = (d_1330_v61_ if default__.fm27(len(d_1237_v0_), p1, d_1323_cf42_.f14, _dafny.Map({d_1323_cf42_.f14: d_1323_cf42_.f14}), globalState) else d_1330_v61_)
                rhs236_ = default__.fm35(92, (d_1324_cf41_) == (p3), d_1323_cf42_.f14, (d_1329_v60_ if default__.fm43(self.f16, globalState) else (d_1330_v61_).cf14), globalState)
                lhs175_ = self
                d_1328_v58_ = rhs232_
                r0 = rhs233_
                lhs175_.f16 = rhs234_
                r1 = rhs235_
                d_1238_v1_ = rhs236_
            elif True:
                d_1332___mcc_h13_ = source20_.cf39
                d_1333_cf39_ = d_1332___mcc_h13_
                d_1334_v62_: _dafny.Array
                nw269_ = _dafny.Array(int(0), 3)
                d_1334_v62_ = nw269_
                d_1334_v62_ = d_1334_v62_
                index225_ = default__.safeIndex(696, (d_1263_v23_).length(0))
                (d_1263_v23_)[index225_] = self.f16
                index226_ = default__.safeIndex(696, (d_1263_v23_).length(0))
                (d_1263_v23_)[index226_] = not(not(not((d_1272_v29_) != ((d_1272_v29_) | (d_1272_v29_)))))
                nw270_ = _dafny.Array(int(0), 3)
                d_1334_v62_ = nw270_
                index227_ = default__.safeIndex(696, (d_1263_v23_).length(0))
                (d_1263_v23_)[index227_] = (d_1261_v21_)[default__.safeIndex(p2, len(d_1261_v21_))]
        hi9_ = p1
        for d_1335_i3_ in range((0) - ((p1) * (p1)), hi9_):
            d_1261_v21_ = d_1261_v21_
            d_1336_v64_: _dafny.Map
            d_1336_v64_ = _dafny.Map({(self).f21: p2})
            d_1337_v65_: _dafny.Map
            d_1337_v65_ = _dafny.Map({(0) - (len(d_1336_v64_)): d_1238_v1_})
            d_1338_v66_: _dafny.Map
            def iife105_():
                coll55_ = _dafny.Map()
                compr_55_: int
                for compr_55_ in (_dafny.Set({p1, d_1335_i3_})).Elements:
                    d_1339_v63_: int = compr_55_
                    if (d_1339_v63_) in (_dafny.Set({p1, d_1335_i3_})):
                        coll55_[(d_1339_v63_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puh"))))] = d_1238_v1_
                return _dafny.Map(coll55_)
            d_1338_v66_ = _dafny.Map({(iife105_()
            ) | (d_1337_v65_): d_1260_v18_})
            d_1338_v66_ = (d_1338_v66_).set(d_1337_v65_, d_1260_v18_)
            d_1340_v67_: C3
            nw271_ = C3()
            nw271_.ctor__(d_1335_i3_, d_1237_v0_, (d_1335_i3_) == (d_1335_i3_), not((self).f21))
            d_1340_v67_ = nw271_
            (globalState).f2 = (0) - ((len((d_1340_v67_).f23)) + (d_1335_i3_))
        d_1341_v68_: _dafny.MultiSet
        d_1341_v68_ = _dafny.MultiSet([p2])
        if default__.fm31((d_1341_v68_).isdisjoint(d_1341_v68_), (d_1237_v0_) <= (d_1237_v0_), globalState):
            (globalState).f7 = default__.safeDivisionInt((p1) + (p2), p2)
            d_1342_v69_: _dafny.Set
            d_1342_v69_ = _dafny.Set({self.f16})
            (globalState).f7 = default__.safeModuloInt((p2) - (p1), (len(d_1342_v69_)) - ((0) - (len(d_1237_v0_))))
            d_1343_v70_: _dafny.Seq
            d_1343_v70_ = _dafny.SeqWithoutIsStrInference([d_1261_v21_])
            d_1344_v71_: _dafny.Array
            nw272_ = _dafny.Array(None, 17)
            nw272_[int(0)] = d_1261_v21_
            nw272_[int(1)] = d_1261_v21_
            nw272_[int(2)] = default__.fm52((self).f21, -962, p1, 129, globalState)
            nw272_[int(3)] = d_1261_v21_
            nw272_[int(4)] = (d_1261_v21_) + (d_1261_v21_)
            nw272_[int(5)] = default__.fm52((self).f21, len(d_1238_v1_), p1, p2, globalState)
            nw272_[int(6)] = d_1261_v21_
            nw272_[int(7)] = (d_1261_v21_) + ((d_1343_v70_)[default__.safeIndex(p1, len(d_1343_v70_))])
            nw272_[int(8)] = d_1261_v21_
            nw272_[int(9)] = d_1261_v21_
            nw272_[int(10)] = d_1261_v21_
            nw272_[int(11)] = d_1261_v21_
            nw272_[int(12)] = d_1261_v21_
            nw272_[int(13)] = d_1261_v21_
            nw272_[int(14)] = d_1261_v21_
            nw272_[int(15)] = _dafny.SeqWithoutIsStrInference([True])
            nw272_[int(16)] = d_1261_v21_
            d_1344_v71_ = nw272_
            d_1344_v71_ = d_1344_v71_
            d_1345_v72_: str
            d_1345_v72_ = _dafny.CodePoint('g')
            d_1237_v0_ = (((d_1237_v0_).set(default__.safeIndex(p1, len(d_1237_v0_)), d_1345_v72_)) + (d_1237_v0_)) + (d_1237_v0_)
            index228_ = default__.safeIndex(468, (d_1263_v23_).length(0))
            (d_1263_v23_)[index228_] = (self).f17
            index229_ = default__.safeIndex(468, (d_1263_v23_).length(0))
            (d_1263_v23_)[index229_] = False
        elif True:
            d_1346_v74_: _dafny.Seq
            def iife106_():
                coll56_ = _dafny.Set()
                compr_56_: int
                for compr_56_ in (d_1238_v1_).Elements:
                    d_1347_v73_: int = compr_56_
                    if (d_1347_v73_) in (d_1238_v1_):
                        coll56_ = coll56_.union(_dafny.Set([default__.safeDivisionInt(d_1347_v73_, 772)]))
                return _dafny.Set(coll56_)
            d_1346_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_1272_v29_).cardinality}), iife106_()
            , d_1260_v18_])
            d_1348_v75_: _dafny.Seq
            d_1348_v75_ = d_1346_v74_
            d_1349_v76_: _dafny.Map
            d_1349_v76_ = _dafny.Map({not ((self).f17) or (not((self).f21)): (d_1348_v75_)})
            d_1349_v76_ = (d_1349_v76_).set((735) < (p1), d_1346_v74_)
            (globalState).f2 = (334) + (p1)
            r2 = not(True)
            d_1350_v77_: _dafny.Array
            nw273_ = _dafny.Array(_dafny.Set({}), 14)
            d_1350_v77_ = nw273_
            d_1351_v78_: _dafny.Map
            d_1351_v78_ = _dafny.Map({p2: d_1350_v77_})
            d_1352_v79_: C0
            nw274_ = C0()
            nw274_.ctor__((d_1350_v77_ if p0 else ((d_1351_v78_)[p2] if (p2) in (d_1351_v78_) else d_1350_v77_)), p2)
            d_1352_v79_ = nw274_
            d_1353_v80_: _dafny.Map
            d_1353_v80_ = _dafny.Map({p0: d_1352_v79_.f27})
            d_1354_v81_: _dafny.Map
            d_1354_v81_ = (d_1353_v80_) | (d_1353_v80_)
            d_1354_v81_ = d_1354_v81_
        r0 = self.f16
        d_1355_v82_: _dafny.Map
        d_1355_v82_ = _dafny.Map({p2: True})
        d_1356_v83_: D7
        d_1356_v83_ = D7_DC11(d_1355_v82_)
        def lambda75_(source21_):
            if source21_.is_DC12:
                return D6_DC9(_dafny.CodePoint('n'))
            elif True:
                d_1357___mcc_h14_ = source21_.cf15
                d_1358_cf15_ = d_1357___mcc_h14_
                return D6_DC9(_dafny.CodePoint('i'))

        r1 = lambda75_(d_1356_v83_)
        r2 = (self).f21
        return r0, r1, r2

    @property
    def f21(self):
        return self._f21

class C5(T2, T1):
    def  __init__(self):
        self._f16: bool = False
        self._f17: bool = False
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

    def fm12(self, p0, globalState):
        return len(_dafny.Set({(_dafny.Map({-22: (self).f17})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s')])): False}))}))

    def fm13(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([not((self).f17)])

    def fm3(self, globalState):
        return ((_dafny.Map({586: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f17, (self).f17]))).cardinality})) | (_dafny.Map({-351: 589}))) | (_dafny.Map({len(_dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f17])): self.f16}))})): 355}))

    def fm23(self, p0, globalState):
        return False

    def fm24(self, p0, p1, p2, globalState):
        def iife107_():
            coll57_ = _dafny.Map()
            compr_57_: int
            for compr_57_ in _dafny.IntegerRange(226, 398):
                d_1359_v0_: int = compr_57_
                if ((226) <= (d_1359_v0_)) and ((d_1359_v0_) < (398)):
                    coll57_[default__.safeDivisionInt(d_1359_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1360_i0_ in range(default__.abs(40))])))] = (self).f17
            return _dafny.Map(coll57_)
        def iife108_():
            coll58_ = _dafny.Map()
            compr_58_: int
            for compr_58_ in _dafny.IntegerRange(-579, 47):
                d_1361_v1_: int = compr_58_
                if ((-579) <= (d_1361_v1_)) and ((d_1361_v1_) < (47)):
                    coll58_[default__.safeDivisionInt(d_1361_v1_, 190)] = (self).f17
            return _dafny.Map(coll58_)
        def iife109_():
            def iife111_():
                coll61_ = _dafny.Set()
                compr_61_: int
                for compr_61_ in _dafny.IntegerRange(109, 285):
                    d_1364_v3_: int = compr_61_
                    if ((109) <= (d_1364_v3_)) and ((d_1364_v3_) < (285)):
                        coll61_ = coll61_.union(_dafny.Set([(d_1364_v3_) * (len(_dafny.SeqWithoutIsStrInference([(self).f17])))]))
                return _dafny.Set(coll61_)
            coll59_ = _dafny.Map()
            def iife110_():
                coll60_ = _dafny.Set()
                compr_60_: int
                for compr_60_ in _dafny.IntegerRange(109, 285):
                    d_1362_v3_: int = compr_60_
                    if ((109) <= (d_1362_v3_)) and ((d_1362_v3_) < (285)):
                        coll60_ = coll60_.union(_dafny.Set([(d_1362_v3_) * (len(_dafny.SeqWithoutIsStrInference([(self).f17])))]))
                return _dafny.Set(coll60_)
            compr_59_: int
            for compr_59_ in (iife110_()
            ).Elements:
                d_1363_v2_: int = compr_59_
                if (d_1363_v2_) in (iife111_()
                ):
                    coll59_[default__.safeDivisionInt(d_1363_v2_, 595)] = not((self).f17)
            return _dafny.Map(coll59_)
        return _dafny.MultiSet([((D7_DC11(iife107_()
)).cf15) | (iife108_()
        ), (iife109_()
         if self.f16 else _dafny.Map({400: True})), _dafny.Map({392: self.f16}), (_dafny.Map({(_dafny.MultiSet([(self).f17, True])).cardinality: (self).f17})) | (_dafny.Map({len(_dafny.Set({(self).f17})): (self).f17})), _dafny.Map({1: self.f16})])

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: bool = False
        (self).f16 = not(p1)
        d_1365_v0_: _dafny.Map
        d_1365_v0_ = _dafny.Map({((self).f17) == (False): (-772) + (p2)})
        d_1365_v0_ = (d_1365_v0_).set(self.f16, p2)
        d_1366_v1_: _dafny.Seq
        d_1366_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaw"))
        (globalState).f2 = (len(d_1366_v1_)) + ((0) - ((self).fm12(p2, globalState)))
        d_1367_v2_: _dafny.MultiSet
        d_1367_v2_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsjaaalnq"))])
        d_1368_v3_: _dafny.Map
        d_1368_v3_ = _dafny.Map({(p2) + (p2): (d_1367_v2_) | (d_1367_v2_)})
        d_1369_v4_: _dafny.Map
        d_1369_v4_ = _dafny.Map({p2: p1})
        d_1368_v3_ = (d_1368_v3_).set(len((_dafny.Map({p2: p0}) if p0 else d_1369_v4_)), ((d_1367_v2_).set(d_1366_v1_, default__.abs(p2))).intersection(d_1367_v2_))
        d_1370_v5_: _dafny.Array
        nw275_ = _dafny.Array(D3.default()(), 18)
        d_1370_v5_ = nw275_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1370_v5_).length(0)):
            d_1371_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_1371_i0_)) and ((d_1371_i0_) < ((d_1370_v5_).length(0)))):
                (d_1370_v5_)[(d_1371_i0_)] = D3_DC3(self.f16)
        d_1372_v6_: _dafny.Seq
        d_1372_v6_ = _dafny.SeqWithoutIsStrInference([(self).fm23(self.f16, globalState)])
        d_1373_v7_: _dafny.Map
        d_1373_v7_ = _dafny.Map({False: d_1372_v6_})
        d_1373_v7_ = d_1373_v7_
        r0 = (p2) * (p2)
        r1 = (self).f17
        r2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suk"))
        r3 = self.f16
        return r0, r1, r2, r3

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_1374_v0_: _dafny.Array
        nw276_ = _dafny.Array(int(0), 6)
        d_1374_v0_ = nw276_
        index230_ = default__.safeIndex(245, (d_1374_v0_).length(0))
        (d_1374_v0_)[index230_] = p0
        index231_ = default__.safeIndex(245, (d_1374_v0_).length(0))
        rhs237_ = p0
        rhs238_ = ((self).fm23(p1, globalState)) or ((self).f17)
        lhs176_ = d_1374_v0_
        lhs177_ = default__.safeIndex(245, (d_1374_v0_).length(0))
        lhs178_ = self
        lhs176_[lhs177_] = rhs237_
        lhs178_.f16 = rhs238_
        d_1375_v1_: _dafny.Seq
        d_1375_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ma"))
        d_1375_v1_ = d_1375_v1_
        d_1376_i0_: int
        d_1376_i0_ = 0
        with _dafny.label("6"):
            while self.f16:
                with _dafny.c_label("6"):
                    if (d_1376_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1376_i0_ = (d_1376_i0_) + (1)
                    (globalState).f7 = p0
                    d_1377_v2_: _dafny.Map
                    d_1377_v2_ = _dafny.Map({(self).f17: self.f16})
                    d_1378_v3_: _dafny.Map
                    d_1378_v3_ = _dafny.Map({self.f16: ((d_1377_v2_)[p1] if (p1) in (d_1377_v2_) else self.f16)})
                    d_1379_v4_: _dafny.Map
                    d_1379_v4_ = _dafny.Map({len(d_1377_v2_): not(not(False))})
                    d_1380_v5_: D7
                    d_1380_v5_ = D7_DC11(d_1379_v4_)
                    source22_ = d_1380_v5_
                    if source22_.is_DC12:
                        (self).f16 = (self).fm23(p1, globalState)
                        d_1381_v6_: str
                        d_1381_v6_ = _dafny.CodePoint('i')
                        d_1382_v7_: _dafny.Set
                        d_1382_v7_ = _dafny.Set({d_1381_v6_, d_1381_v6_, d_1381_v6_, d_1381_v6_, _dafny.CodePoint('h')})
                        def iife112_():
                            coll62_ = _dafny.Set()
                            compr_62_: str
                            for compr_62_ in (_dafny.MultiSet(d_1375_v1_)).Elements:
                                d_1383_v8_: str = compr_62_
                                if (d_1383_v8_) in (_dafny.MultiSet(d_1375_v1_)):
                                    coll62_ = coll62_.union(_dafny.Set([d_1383_v8_]))
                            return _dafny.Set(coll62_)
                        (globalState).f2 = len(((d_1382_v7_).intersection(iife112_()
                        )).intersection(d_1382_v7_))
                        d_1384_v9_: _dafny.Array
                        nw277_ = _dafny.Array(False, 2)
                        d_1384_v9_ = nw277_
                        d_1385_v10_: _dafny.Array
                        d_1385_v10_ = d_1384_v9_
                        d_1386_v11_: _dafny.Map
                        d_1386_v11_ = _dafny.Map({self.f16: d_1385_v10_})
                        d_1387_v12_: _dafny.Map
                        d_1387_v12_ = _dafny.Map({(self).f17: d_1386_v11_})
                        d_1388_v13_: _dafny.Map
                        d_1388_v13_ = _dafny.Map({d_1384_v9_: ((d_1387_v12_)[(self).f17] if ((self).f17) in (d_1387_v12_) else d_1386_v11_)})
                        rhs239_ = (self.f16) == (p1)
                        rhs240_ = d_1388_v13_
                        rhs241_ = p1
                        lhs179_ = self
                        lhs180_ = self
                        lhs179_.f16 = rhs239_
                        d_1388_v13_ = rhs240_
                        lhs180_.f16 = rhs241_
                        d_1389_v14_: _dafny.Map
                        d_1389_v14_ = _dafny.Map({(self).f17: d_1384_v9_})
                        d_1390_v15_: _dafny.Map
                        d_1390_v15_ = _dafny.Map({((d_1379_v4_)[len(d_1389_v14_)] if (len(d_1389_v14_)) in (d_1379_v4_) else self.f16): (p0) - (len(_dafny.Map({d_1374_v0_: p0})))})
                        d_1391_v16_: _dafny.Map
                        d_1391_v16_ = d_1390_v15_
                        d_1390_v15_ = ((d_1391_v16_)) | (d_1390_v15_)
                    elif True:
                        d_1392___mcc_h0_ = source22_.cf15
                        d_1393_cf15_ = d_1392___mcc_h0_
                        d_1394_v17_: _dafny.Map
                        d_1394_v17_ = _dafny.Map({p1: p0})
                        d_1395_v18_: _dafny.Map
                        d_1395_v18_ = _dafny.Map({d_1375_v1_: len((d_1394_v17_) | (default__.fm25((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], globalState)))})
                        d_1396_v19_: _dafny.Seq
                        d_1396_v19_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                        d_1397_v20_: _dafny.MultiSet
                        d_1397_v20_ = _dafny.MultiSet([d_1396_v19_])
                        d_1398_v21_: str
                        d_1398_v21_ = _dafny.CodePoint('v')
                        d_1399_v22_: _dafny.Map
                        d_1399_v22_ = _dafny.Map({p0: d_1398_v21_})
                        d_1400_v23_: _dafny.Map
                        d_1400_v23_ = _dafny.Map({((d_1397_v20_)[d_1396_v19_] if (d_1396_v19_) in (d_1397_v20_) else p0): default__.fm26((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], p0, len(d_1399_v22_), globalState)})
                        d_1395_v18_ = (d_1395_v18_).set(d_1375_v1_, ((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]) + (len(d_1400_v23_)))
                        d_1401_v24_: _dafny.Map
                        d_1401_v24_ = _dafny.Map({(d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]: d_1375_v1_})
                        d_1401_v24_ = (d_1401_v24_).set((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], d_1375_v1_)
                        d_1402_v25_: _dafny.Array
                        nw278_ = _dafny.Array(_dafny.Set({}), 26)
                        d_1402_v25_ = nw278_
                        d_1403_v26_: C0
                        nw279_ = C0()
                        nw279_.ctor__(d_1402_v25_, len(_dafny.SeqWithoutIsStrInference([d_1398_v21_ for d_1404_i1_ in range(default__.abs(350))])))
                        d_1403_v26_ = nw279_
                        (self).f16 = self.f16
                    (self).f16 = (self).f17
                    d_1405_v27_: _dafny.Seq
                    d_1405_v27_ = _dafny.SeqWithoutIsStrInference([p0, (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], len((self).fm13((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], p1, (self).f17, globalState)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jx")))])
                    d_1406_v28_: _dafny.Seq
                    d_1406_v28_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([p0 for d_1407_i2_ in range(default__.abs(774))])) == (d_1405_v27_), self.f16])
                    d_1406_v28_ = d_1406_v28_
                    pass
            pass
        if (self).f17:
            d_1408_v29_: _dafny.Seq
            d_1408_v29_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1409_v30_: _dafny.Map
            d_1409_v30_ = _dafny.Map({(d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]: len(d_1408_v29_)})
            (self).f16 = (-434) not in (d_1409_v30_)
            d_1410_v31_: str
            d_1410_v31_ = _dafny.CodePoint('y')
            index232_ = default__.safeIndex(245, (d_1374_v0_).length(0))
            rhs242_ = (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]
            rhs243_ = d_1410_v31_
            rhs244_ = ((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]) + ((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))])
            lhs181_ = globalState
            lhs182_ = d_1374_v0_
            lhs183_ = default__.safeIndex(245, (d_1374_v0_).length(0))
            lhs181_.f7 = rhs242_
            d_1410_v31_ = rhs243_
            lhs182_[lhs183_] = rhs244_
            d_1411_v32_: _dafny.Array
            nw280_ = _dafny.Array(None, 14)
            nw280_[int(0)] = d_1375_v1_
            nw280_[int(1)] = d_1375_v1_
            nw280_[int(2)] = d_1375_v1_
            nw280_[int(3)] = d_1375_v1_
            nw280_[int(4)] = d_1375_v1_
            nw280_[int(5)] = d_1375_v1_
            nw280_[int(6)] = (d_1375_v1_) + (d_1375_v1_)
            nw280_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfbtphsjv"))
            nw280_[int(8)] = d_1375_v1_
            nw280_[int(9)] = d_1375_v1_
            nw280_[int(10)] = d_1375_v1_
            nw280_[int(11)] = d_1375_v1_
            nw280_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocouv"))
            nw280_[int(13)] = (d_1375_v1_) + (d_1375_v1_)
            d_1411_v32_ = nw280_
            index233_ = default__.safeIndex(596, (d_1411_v32_).length(0))
            (d_1411_v32_)[index233_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xaxrbysaq"))
            index234_ = default__.safeIndex(596, (d_1411_v32_).length(0))
            (d_1411_v32_)[index234_] = (_dafny.SeqWithoutIsStrInference([d_1410_v31_ for d_1412_i3_ in range(default__.abs(-244))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1413_i4_ in range(default__.abs(308))]))
            d_1414_v33_: C1
            nw281_ = C1()
            nw281_.ctor__(p1, default__.safeModuloInt(p0, len(d_1409_v30_)))
            d_1414_v33_ = nw281_
            d_1415_v34_: _dafny.Set
            d_1415_v34_ = _dafny.Set({d_1410_v31_, d_1410_v31_})
            d_1416_v35_: C4
            nw282_ = C4()
            nw282_.ctor__(p1, (self).f17, (d_1415_v34_).issubset(d_1415_v34_))
            d_1416_v35_ = nw282_
        elif True:
            d_1417_v36_: _dafny.Map
            d_1417_v36_ = _dafny.Map({p0: (self).f17})
            d_1418_v37_: _dafny.Map
            d_1418_v37_ = _dafny.Map({p1: len(d_1375_v1_)})
            (self).f16 = (default__.safeModuloInt((0) - (len(d_1417_v36_)), (0) - (p0))) <= (((d_1418_v37_)[p1] if (p1) in (d_1418_v37_) else (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]))
            d_1419_v38_: _dafny.Seq
            d_1419_v38_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_1420_v39_: _dafny.Seq
            d_1420_v39_ = d_1419_v38_
            d_1421_v40_: D4
            d_1421_v40_ = D4_DC7(d_1420_v39_, 357)
            d_1422_v41_: _dafny.Seq
            d_1422_v41_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1423_v42_: _dafny.Map
            d_1423_v42_ = _dafny.Map({d_1421_v40_: (d_1422_v41_)[default__.safeIndex((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], len(d_1422_v41_))]})
            (globalState).f2 = (0) - (((d_1423_v42_)[d_1421_v40_] if (d_1421_v40_) in (d_1423_v42_) else (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]))
            r0 = default__.fm0(p1, globalState)
            (self).f16 = default__.fm43((len(d_1422_v41_)) > ((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]), globalState)
            d_1424_v43_: D3
            d_1424_v43_ = D3_DC3(self.f16)
            d_1425_v44_: C2
            nw283_ = C2()
            nw283_.ctor__(d_1424_v43_, p1, (self).f17)
            d_1425_v44_ = nw283_
        d_1426_v45_: _dafny.Seq
        d_1426_v45_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1427_v46_: _dafny.Map
        d_1427_v46_ = _dafny.Map({len(d_1426_v45_): p1})
        d_1428_v48_: _dafny.MultiSet
        d_1428_v48_ = _dafny.MultiSet([(d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]])
        d_1429_v49_: D7
        def iife113_():
            coll63_ = _dafny.Map()
            compr_63_: int
            for compr_63_ in (d_1428_v48_).Elements:
                d_1430_v47_: int = compr_63_
                if (d_1430_v47_) in (d_1428_v48_):
                    coll63_[(d_1430_v47_) + ((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))])] = self.f16
            return _dafny.Map(coll63_)
        d_1429_v49_ = D7_DC11(iife113_()
)
        pat_let_tv18_ = d_1374_v0_
        pat_let_tv19_ = d_1374_v0_
        pat_let_tv20_ = d_1374_v0_
        pat_let_tv21_ = d_1374_v0_
        index235_ = default__.safeIndex(245, (d_1374_v0_).length(0))
        def lambda76_(source23_):
            if source23_.is_DC12:
                return (pat_let_tv19_)[default__.safeIndex(245, (pat_let_tv18_).length(0))]
            elif True:
                d_1431___mcc_h1_ = source23_.cf15
                d_1432_cf15_ = d_1431___mcc_h1_
                return (pat_let_tv21_)[default__.safeIndex(245, (pat_let_tv20_).length(0))]

        (d_1374_v0_)[index235_] = lambda76_((D7_DC11(d_1427_v46_) if (self).fm23((self).f17, globalState) else d_1429_v49_))
        d_1433_v50_: T0
        nw284_ = C1()
        nw284_.ctor__((self).f17, p0)
        d_1433_v50_ = nw284_
        d_1434_v51_: D14
        d_1434_v51_ = D14_DC28((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], p1, d_1433_v50_)
        d_1435_v52_: D3
        d_1435_v52_ = D3_DC4(p1, p0, (d_1434_v51_).cf40, (d_1426_v45_) <= ((d_1426_v45_).set(default__.safeIndex(p0, len(d_1426_v45_)), d_1433_v50_.f14)), (self.f16) or (True))
        source24_ = d_1435_v52_
        if source24_.is_DC4:
            d_1436___mcc_h2_ = source24_.cf4
            d_1437___mcc_h3_ = source24_.cf5
            d_1438___mcc_h4_ = source24_.cf6
            d_1439___mcc_h5_ = source24_.cf7
            d_1440___mcc_h6_ = source24_.cf8
            d_1441_cf8_ = d_1440___mcc_h6_
            d_1442_cf7_ = d_1439___mcc_h5_
            d_1443_cf6_ = d_1438___mcc_h4_
            d_1444_cf5_ = d_1437___mcc_h3_
            d_1445_cf4_ = d_1436___mcc_h2_
            index236_ = default__.safeIndex(245, (d_1374_v0_).length(0))
            (d_1374_v0_)[index236_] = d_1443_cf6_
            d_1446_v54_: _dafny.Seq
            def iife114_():
                coll64_ = _dafny.Map()
                compr_64_: int
                for compr_64_ in _dafny.IntegerRange(699, 566):
                    d_1447_v53_: int = compr_64_
                    if ((699) <= (d_1447_v53_)) and ((d_1447_v53_) < (566)):
                        coll64_[(d_1447_v53_) * (len(d_1375_v1_))] = True
                return _dafny.Map(coll64_)
            d_1446_v54_ = _dafny.SeqWithoutIsStrInference([iife114_()
            ])
            d_1448_v55_: _dafny.Map
            d_1448_v55_ = _dafny.Map({(d_1446_v54_)[default__.safeIndex(p0, len(d_1446_v54_))]: d_1444_cf5_})
            d_1449_v56_: _dafny.Array
            nw285_ = _dafny.Array(None, 7)
            nw285_[int(0)] = (d_1448_v55_).set(d_1427_v46_, (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))])
            nw285_[int(1)] = d_1448_v55_
            nw285_[int(2)] = _dafny.Map({(d_1429_v49_).cf15: (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]})
            nw285_[int(3)] = (d_1448_v55_) | (d_1448_v55_)
            nw285_[int(4)] = d_1448_v55_
            nw285_[int(5)] = d_1448_v55_
            nw285_[int(6)] = d_1448_v55_
            d_1449_v56_ = nw285_
            index237_ = default__.safeIndex(224, (d_1449_v56_).length(0))
            (d_1449_v56_)[index237_] = d_1448_v55_
            d_1450_v57_: _dafny.Array
            nw286_ = _dafny.Array(False, 26)
            d_1450_v57_ = nw286_
            index238_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            (d_1450_v57_)[index238_] = (self).f17
            index239_ = default__.safeIndex(224, (d_1449_v56_).length(0))
            index240_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            def iife115_():
                coll65_ = _dafny.Map()
                compr_65_: _dafny.Map
                for compr_65_ in (_dafny.Map({_dafny.Map({d_1433_v50_.f14: True}): 163})).keys.Elements:
                    d_1451_v58_: _dafny.Map = compr_65_
                    if (d_1451_v58_) in (_dafny.Map({_dafny.Map({d_1433_v50_.f14: True}): 163})):
                        coll65_[d_1451_v58_] = (0) - (d_1433_v50_.f14)
                return _dafny.Map(coll65_)
            rhs245_ = ((d_1448_v55_) | (iife115_()
            )) | (d_1448_v55_)
            rhs246_ = self.f16
            lhs184_ = d_1449_v56_
            lhs185_ = default__.safeIndex(224, (d_1449_v56_).length(0))
            lhs186_ = d_1450_v57_
            lhs187_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            lhs184_[lhs185_] = rhs245_
            lhs186_[lhs187_] = rhs246_
            d_1452_v59_: _dafny.Array
            def lambda77_(d_1453_v1_, d_1454_v46_):
                def lambda78_(d_1455_i5_):
                    return (d_1453_v1_).set(default__.safeIndex(len(d_1454_v46_), len(d_1453_v1_)), _dafny.CodePoint('h'))

                return lambda78_

            init43_ = lambda77_(d_1375_v1_, d_1427_v46_)
            nw287_ = _dafny.Array(None, 2)
            for i0_43_ in range(nw287_.length(0)):
                nw287_[i0_43_] = init43_(i0_43_)
            d_1452_v59_ = nw287_
            d_1456_v60_: _dafny.Seq
            d_1456_v60_ = d_1426_v45_
            rhs247_ = d_1433_v50_.f14
            rhs248_ = False
            rhs249_ = ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqlcg")))])) - (d_1428_v48_)).isdisjoint(_dafny.MultiSet((d_1456_v60_)))
            rhs250_ = d_1452_v59_
            lhs188_ = globalState
            lhs188_.f7 = rhs247_
            d_1445_cf4_ = rhs248_
            d_1441_cf8_ = rhs249_
            d_1452_v59_ = rhs250_
            d_1457_v61_: _dafny.Map
            d_1457_v61_ = _dafny.Map({(self).f17: self.f16})
            d_1458_v62_: str
            d_1458_v62_ = _dafny.CodePoint('h')
            index241_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            index242_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            rhs251_ = (d_1375_v1_) + ((d_1375_v1_) + (d_1375_v1_))
            rhs252_ = ((d_1457_v61_)[(not(d_1442_cf7_)) or (self.f16)] if ((not(d_1442_cf7_)) or (self.f16)) in (d_1457_v61_) else (d_1450_v57_)[default__.safeIndex(568, (d_1450_v57_).length(0))])
            rhs253_ = (d_1375_v1_) < ((d_1375_v1_).set(default__.safeIndex((d_1426_v45_)[default__.safeIndex((d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))], len(d_1426_v45_))], len(d_1375_v1_)), d_1458_v62_))
            lhs189_ = d_1450_v57_
            lhs190_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            lhs191_ = d_1450_v57_
            lhs192_ = default__.safeIndex(568, (d_1450_v57_).length(0))
            d_1375_v1_ = rhs251_
            lhs189_[lhs190_] = rhs252_
            lhs191_[lhs192_] = rhs253_
        elif source24_.is_DC3:
            d_1459___mcc_h7_ = source24_.cf3
            d_1460_cf3_ = d_1459___mcc_h7_
            d_1461_v63_: _dafny.Array
            nw288_ = _dafny.Array(None, 8)
            nw288_[int(0)] = d_1460_cf3_
            nw288_[int(1)] = (self).f17
            nw288_[int(2)] = self.f16
            nw288_[int(3)] = p1
            nw288_[int(4)] = not(not(self.f16))
            nw288_[int(5)] = False
            nw288_[int(6)] = p1
            nw288_[int(7)] = self.f16
            d_1461_v63_ = nw288_
            index243_ = default__.safeIndex(340, (d_1461_v63_).length(0))
            (d_1461_v63_)[index243_] = d_1460_cf3_
            d_1462_v64_: _dafny.Seq
            d_1462_v64_ = _dafny.SeqWithoutIsStrInference([p1, False, False])
            index244_ = default__.safeIndex(340, (d_1461_v63_).length(0))
            (d_1461_v63_)[index244_] = default__.fm31((self).f17, (d_1462_v64_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xus"))), len(d_1462_v64_))], globalState)
            index245_ = default__.safeIndex(340, (d_1461_v63_).length(0))
            (d_1461_v63_)[index245_] = (self).f17
            index246_ = default__.safeIndex(245, (d_1374_v0_).length(0))
            (d_1374_v0_)[index246_] = len(d_1426_v45_)
            d_1463_v65_: D9
            d_1463_v65_ = D9_DC16(p0, (d_1461_v63_)[default__.safeIndex(340, (d_1461_v63_).length(0))])
            d_1464_v66_: D9
            d_1464_v66_ = D9_DC17(d_1463_v65_)
            d_1465_v67_: _dafny.Map
            d_1465_v67_ = _dafny.Map({(d_1461_v63_)[default__.safeIndex(340, (d_1461_v63_).length(0))]: d_1460_cf3_})
            d_1466_v68_: _dafny.Map
            d_1466_v68_ = _dafny.Map({not(d_1460_cf3_): d_1465_v67_})
            rhs254_ = D9_DC17(D9_DC16(len(d_1426_v45_), (self).f17))
            rhs255_ = ((d_1466_v68_) | (_dafny.Map({self.f16: d_1465_v67_}))) | (_dafny.Map({not(d_1460_cf3_): d_1465_v67_}))
            d_1464_v66_ = rhs254_
            d_1466_v68_ = rhs255_
        elif True:
            d_1467___mcc_h8_ = source24_.cf9
            d_1468_cf9_ = d_1467___mcc_h8_
            d_1469_v69_: _dafny.Array
            def lambda79_(d_1470_v1_):
                def lambda80_(d_1471_i6_):
                    return (d_1470_v1_) < (d_1470_v1_)

                return lambda80_

            init44_ = lambda79_(d_1375_v1_)
            nw289_ = _dafny.Array(None, 21)
            for i0_44_ in range(nw289_.length(0)):
                nw289_[i0_44_] = init44_(i0_44_)
            d_1469_v69_ = nw289_
            index247_ = default__.safeIndex(263, (d_1469_v69_).length(0))
            (d_1469_v69_)[index247_] = (self).f17
            index248_ = default__.safeIndex(263, (d_1469_v69_).length(0))
            (d_1469_v69_)[index248_] = (self).f17
            d_1374_v0_ = d_1374_v0_
            index249_ = default__.safeIndex(245, (d_1374_v0_).length(0))
            (d_1374_v0_)[index249_] = (d_1374_v0_)[default__.safeIndex(245, (d_1374_v0_).length(0))]
            if True:
                d_1472_v70_: _dafny.MultiSet
                d_1472_v70_ = _dafny.MultiSet([p1, p1, (self).f17])
                d_1473_v71_: _dafny.Seq
                d_1473_v71_ = _dafny.SeqWithoutIsStrInference([self.f16])
                (d_1433_v50_).f14 = ((d_1472_v70_) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(self).f17, (d_1469_v69_)[default__.safeIndex(263, (d_1469_v69_).length(0))], not((self).f17), (self).f17])) + (d_1473_v71_)))).cardinality
                d_1474_v72_: _dafny.Array
                nw290_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1474_v72_ = nw290_
                d_1475_v73_: _dafny.Array
                nw291_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_1475_v73_ = nw291_
                index250_ = default__.safeIndex(442, (d_1474_v72_).length(0))
                (d_1474_v72_)[index250_] = d_1475_v73_
                index251_ = default__.safeIndex(442, (d_1474_v72_).length(0))
                (d_1474_v72_)[index251_] = d_1475_v73_
                d_1476_v74_: _dafny.Seq
                d_1476_v74_ = _dafny.SeqWithoutIsStrInference([d_1374_v0_, (d_1374_v0_ if (self).f17 else d_1374_v0_), d_1374_v0_, d_1374_v0_, d_1374_v0_])
                index252_ = default__.safeIndex(245, (d_1374_v0_).length(0))
                (d_1374_v0_)[index252_] = len(d_1476_v74_)
                d_1477_v75_: C4
                nw292_ = C4()
                nw292_.ctor__(False, self.f16, True)
                d_1477_v75_ = nw292_
                d_1478_v76_: str
                d_1478_v76_ = _dafny.CodePoint('w')
                d_1479_v77_: D6
                d_1479_v77_ = D6_DC9(d_1478_v76_)
                d_1480_v78_: _dafny.Array
                nw293_ = _dafny.Array(None, 5)
                nw293_[int(0)] = d_1479_v77_
                nw293_[int(1)] = d_1479_v77_
                nw293_[int(2)] = d_1479_v77_
                nw293_[int(3)] = d_1479_v77_
                nw293_[int(4)] = d_1479_v77_
                d_1480_v78_ = nw293_
                d_1480_v78_ = d_1480_v78_
            elif True:
                d_1481_v79_: _dafny.Map
                d_1481_v79_ = _dafny.Map({len(d_1375_v1_): d_1375_v1_})
                d_1481_v79_ = (d_1481_v79_).set(d_1433_v50_.f14, d_1375_v1_)
                d_1482_v80_: _dafny.Seq
                d_1482_v80_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                d_1483_v81_: _dafny.Set
                d_1483_v81_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(d_1375_v1_) for d_1484_i7_ in range(default__.abs(603))]))})
                d_1485_v82_: D9
                d_1485_v82_ = D9_DC15((d_1469_v69_)[default__.safeIndex(263, (d_1469_v69_).length(0))], (self).f17, d_1483_v81_, d_1374_v0_, d_1469_v69_)
                d_1486_v83_: _dafny.Array
                nw294_ = _dafny.Array(None, 5)
                nw294_[int(0)] = d_1485_v82_
                nw294_[int(1)] = d_1485_v82_
                nw294_[int(2)] = d_1485_v82_
                nw294_[int(3)] = d_1485_v82_
                nw294_[int(4)] = d_1485_v82_
                d_1486_v83_ = nw294_
                d_1487_v84_: _dafny.Seq
                d_1487_v84_ = _dafny.SeqWithoutIsStrInference([D3_DC3(self.f16)])
                d_1488_v85_: D12
                d_1488_v85_ = D12_DC24(self.f16, d_1482_v80_, _dafny.SeqWithoutIsStrInference([d_1486_v83_]), (self).f17, d_1487_v84_)
                d_1489_v86_: _dafny.Seq
                d_1489_v86_ = _dafny.SeqWithoutIsStrInference([(d_1488_v85_).cf36, (d_1469_v69_)[default__.safeIndex(263, (d_1469_v69_).length(0))], (self).f17])
                (self).f16 = ((476) * (len(d_1482_v80_))) < (p0)
                index253_ = default__.safeIndex(245, (d_1374_v0_).length(0))
                (d_1374_v0_)[index253_] = 140
                d_1426_v45_ = ((d_1426_v45_) + (d_1426_v45_)) + (((_dafny.SeqWithoutIsStrInference([(D16_DC34(True, d_1433_v50_, d_1433_v50_.f14)).cf53, d_1433_v50_.f14])) + (d_1426_v45_)).set(default__.safeIndex(87, len((_dafny.SeqWithoutIsStrInference([(D16_DC34(True, d_1433_v50_, d_1433_v50_.f14)).cf53, d_1433_v50_.f14])) + (d_1426_v45_))), d_1433_v50_.f14))
                d_1469_v69_ = d_1469_v69_
        r0 = default__.safeDivisionInt(p0, d_1433_v50_.f14)
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_1490_v0_: _dafny.Seq
        d_1490_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdkvfv"))
        d_1491_v1_: _dafny.Map
        d_1491_v1_ = _dafny.Map({d_1490_v0_: (p1) != (p1)})
        if ((d_1491_v1_)[d_1490_v0_] if (d_1490_v0_) in (d_1491_v1_) else (self).f17):
            d_1492_v2_: _dafny.MultiSet
            d_1492_v2_ = _dafny.MultiSet([p0])
            d_1493_v3_: _dafny.Map
            d_1493_v3_ = _dafny.Map({p1: len(d_1490_v0_)})
            d_1494_v4_: _dafny.Seq
            d_1494_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm31(self.f16, default__.fm27(len(_dafny.SeqWithoutIsStrInference([(self).f17])), (d_1492_v2_).cardinality, p3, d_1493_v3_, globalState), globalState), self.f16])
            d_1494_v4_ = (d_1494_v4_).set(default__.safeIndex(p1, len(d_1494_v4_)), not(False))
            d_1495_v5_: _dafny.Map
            d_1495_v5_ = _dafny.Map({False: p2})
            d_1496_v6_: D4
            d_1496_v6_ = D4_DC6(d_1495_v5_)
            d_1497_v7_: _dafny.Map
            d_1497_v7_ = _dafny.Map({p3: self.f16})
            d_1498_v8_: _dafny.Array
            nw295_ = _dafny.Array(None, 25)
            nw295_[int(0)] = p1
            nw295_[int(1)] = default__.fm0((self).f17, globalState)
            nw295_[int(2)] = p1
            nw295_[int(3)] = p1
            nw295_[int(4)] = p1
            nw295_[int(5)] = (0) - (p1)
            nw295_[int(6)] = 558
            nw295_[int(7)] = p3
            nw295_[int(8)] = p3
            nw295_[int(9)] = 841
            nw295_[int(10)] = p3
            nw295_[int(11)] = p1
            nw295_[int(12)] = p1
            nw295_[int(13)] = 762
            nw295_[int(14)] = p3
            nw295_[int(15)] = p1
            nw295_[int(16)] = p3
            nw295_[int(17)] = len((d_1496_v6_).cf10)
            nw295_[int(18)] = p3
            nw295_[int(19)] = p1
            nw295_[int(20)] = (self).fm12(865, globalState)
            nw295_[int(21)] = default__.fm0((self).f17, globalState)
            nw295_[int(22)] = 399
            nw295_[int(23)] = len((d_1497_v7_).set(p1, not((self).f17)))
            nw295_[int(24)] = p1
            d_1498_v8_ = nw295_
            d_1499_v9_: _dafny.Map
            d_1499_v9_ = _dafny.Map({d_1490_v0_: d_1498_v8_})
            d_1499_v9_ = (d_1499_v9_).set((d_1490_v0_) + (d_1490_v0_), d_1498_v8_)
            d_1500_v10_: T0
            nw296_ = C1()
            nw296_.ctor__(self.f16, p3)
            d_1500_v10_ = nw296_
            d_1501_v11_: D14
            d_1501_v11_ = D14_DC28(p1, (self).f17, d_1500_v10_)
            d_1502_v12_: _dafny.Seq
            d_1502_v12_ = _dafny.SeqWithoutIsStrInference([p3, len(d_1490_v0_), d_1500_v10_.f14, (0) - (p3)])
            d_1503_v13_: _dafny.Array
            nw297_ = _dafny.Array(False, 29)
            d_1503_v13_ = nw297_
            d_1504_v14_: _dafny.Map
            d_1504_v14_ = _dafny.Map({default__.fm27(p1, (d_1501_v11_).cf40, len(d_1502_v12_), d_1493_v3_, globalState): d_1503_v13_})
            d_1505_v15_: _dafny.Set
            d_1505_v15_ = _dafny.Set({(0) - (len(d_1494_v4_)), p3})
            d_1504_v14_ = (d_1504_v14_).set((d_1505_v15_).ispropersubset(d_1505_v15_), d_1503_v13_)
            index254_ = default__.safeIndex(10, (d_1503_v13_).length(0))
            (d_1503_v13_)[index254_] = (912) > (732)
            index255_ = default__.safeIndex(10, (d_1503_v13_).length(0))
            (d_1503_v13_)[index255_] = not (self.f16) or (False)
            (globalState).f2 = default__.fm0((self).f17, globalState)
        elif True:
            d_1506_v16_: _dafny.Map
            d_1506_v16_ = _dafny.Map({self.f16: p2})
            d_1507_v17_: _dafny.Seq
            d_1507_v17_ = _dafny.SeqWithoutIsStrInference([(self).f17])
            d_1508_v18_: C3
            nw298_ = C3()
            nw298_.ctor__((p1) * (len(d_1490_v0_)), d_1490_v0_, ((d_1506_v16_)[(d_1507_v17_)[default__.safeIndex(len(d_1506_v16_), len(d_1507_v17_))]] if ((d_1507_v17_)[default__.safeIndex(len(d_1506_v16_), len(d_1507_v17_))]) in (d_1506_v16_) else p2), True)
            d_1508_v18_ = nw298_
            d_1509_v19_: _dafny.Array
            nw299_ = _dafny.Array(int(0), 27)
            d_1509_v19_ = nw299_
            index256_ = default__.safeIndex(436, (d_1509_v19_).length(0))
            (d_1509_v19_)[index256_] = p1
            index257_ = default__.safeIndex(436, (d_1509_v19_).length(0))
            rhs256_ = p3
            rhs257_ = d_1490_v0_
            rhs258_ = 563
            lhs193_ = globalState
            lhs194_ = d_1509_v19_
            lhs195_ = default__.safeIndex(436, (d_1509_v19_).length(0))
            lhs193_.f2 = rhs256_
            d_1490_v0_ = rhs257_
            lhs194_[lhs195_] = rhs258_
            if (self).f17:
                d_1510_v20_: _dafny.Array
                def lambda81_(d_1511_i0_):
                    return self.f16

                init45_ = lambda81_
                nw300_ = _dafny.Array(None, 12)
                for i0_45_ in range(nw300_.length(0)):
                    nw300_[i0_45_] = init45_(i0_45_)
                d_1510_v20_ = nw300_
                index258_ = default__.safeIndex(370, (d_1510_v20_).length(0))
                (d_1510_v20_)[index258_] = True
                d_1512_v21_: _dafny.Seq
                d_1512_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1508_v18_).f22: (self).f17})])
                index259_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                index260_ = default__.safeIndex(370, (d_1510_v20_).length(0))
                rhs259_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wikghvh")) if p2 else _dafny.SeqWithoutIsStrInference([p0 for d_1513_i1_ in range(default__.abs(792))]))
                rhs260_ = default__.safeDivisionInt(len(d_1512_v21_), (0) - (p3))
                rhs261_ = len(d_1490_v0_)
                rhs262_ = (self).f17
                rhs263_ = (d_1508_v18_).f22
                lhs196_ = globalState
                lhs197_ = d_1509_v19_
                lhs198_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                lhs199_ = d_1510_v20_
                lhs200_ = default__.safeIndex(370, (d_1510_v20_).length(0))
                d_1490_v0_ = rhs259_
                lhs196_.f7 = rhs260_
                lhs197_[lhs198_] = rhs261_
                lhs199_[lhs200_] = rhs262_
                r0 = rhs263_
                index261_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                (d_1509_v19_)[index261_] = (d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))]
                index262_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                (d_1509_v19_)[index262_] = ((d_1508_v18_).f22 if (self).f17 else (d_1508_v18_).f22)
                d_1514_v22_: _dafny.Map
                d_1514_v22_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0 for d_1515_i2_ in range(default__.abs(-561))])): (d_1490_v0_) + (d_1490_v0_)})
                d_1514_v22_ = (d_1514_v22_).set(683, (d_1508_v18_).f23)
                (globalState).f7 = p3
            elif True:
                index263_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                index264_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                index265_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                rhs264_ = default__.safeDivisionInt((d_1508_v18_).f22, (d_1508_v18_).f22)
                rhs265_ = (d_1508_v18_).f22
                rhs266_ = (d_1508_v18_).f22
                rhs267_ = p1
                lhs201_ = d_1509_v19_
                lhs202_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                lhs203_ = d_1509_v19_
                lhs204_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                lhs205_ = d_1509_v19_
                lhs206_ = default__.safeIndex(436, (d_1509_v19_).length(0))
                lhs207_ = globalState
                lhs201_[lhs202_] = rhs264_
                lhs203_[lhs204_] = rhs265_
                lhs205_[lhs206_] = rhs266_
                lhs207_.f2 = rhs267_
                d_1516_v23_: C4
                nw301_ = C4()
                nw301_.ctor__((self).f17, self.f16, p2)
                d_1516_v23_ = nw301_
                d_1517_v24_: _dafny.MultiSet
                d_1517_v24_ = _dafny.MultiSet([((d_1516_v23_).f21 if True else p2), p2, self.f16])
                rhs268_ = (d_1517_v24_ if not (self.f16) or ((self).f17) else (d_1517_v24_ if p2 else d_1517_v24_))
                rhs269_ = p3
                d_1517_v24_ = rhs268_
                r0 = rhs269_
                d_1518_v25_: _dafny.Set
                d_1518_v25_ = _dafny.Set({-158})
                d_1519_v26_: _dafny.Array
                nw302_ = _dafny.Array(False, 21)
                d_1519_v26_ = nw302_
                d_1520_v27_: D9
                d_1520_v27_ = D9_DC15((self).f17, (self).f17, d_1518_v25_, d_1509_v19_, d_1519_v26_)
                d_1521_v28_: T0
                nw303_ = C1()
                nw303_.ctor__((d_1520_v27_).cf18, p1)
                d_1521_v28_ = nw303_
                d_1522_v29_: D14
                d_1522_v29_ = D14_DC28((d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))], p2, d_1521_v28_)
                d_1518_v25_ = _dafny.Set({(d_1508_v18_).f22, (d_1522_v29_).cf40, len((d_1508_v18_).f23)})
                d_1523_v30_: D3
                d_1523_v30_ = D3_DC3(p2)
                d_1524_v31_: C2
                nw304_ = C2()
                nw304_.ctor__(d_1523_v30_, (d_1516_v23_).f21, not(not((self).f17)))
                d_1524_v31_ = nw304_
            index266_ = default__.safeIndex(436, (d_1509_v19_).length(0))
            (d_1509_v19_)[index266_] = (d_1508_v18_).f22
            d_1525_v32_: _dafny.Set
            d_1525_v32_ = _dafny.Set({(d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))], (d_1508_v18_).f22})
            d_1526_v33_: _dafny.Map
            d_1526_v33_ = _dafny.Map({p3: len(d_1525_v32_)})
            d_1527_v36_: _dafny.Map
            def iife116_():
                coll66_ = _dafny.Map()
                compr_66_: int
                for compr_66_ in _dafny.IntegerRange(8, 230):
                    d_1528_v35_: int = compr_66_
                    if ((8) <= (d_1528_v35_)) and ((d_1528_v35_) < (230)):
                        coll66_[(d_1528_v35_) * ((d_1508_v18_).f22)] = 333
                return _dafny.Map(coll66_)
            d_1527_v36_ = _dafny.Map({not(self.f16): iife116_()
            })
            d_1529_v37_: _dafny.Seq
            d_1529_v37_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1530_v38_: D11
            d_1530_v38_ = D11_DC19(d_1526_v33_)
            d_1531_v39_: _dafny.Array
            nw305_ = _dafny.Array(None, 28)
            nw305_[int(0)] = (d_1526_v33_) | (d_1526_v33_)
            def iife117_():
                coll67_ = _dafny.Map()
                compr_67_: int
                for compr_67_ in _dafny.IntegerRange(736, 62):
                    d_1532_v34_: int = compr_67_
                    if ((736) <= (d_1532_v34_)) and ((d_1532_v34_) < (62)):
                        coll67_[(d_1532_v34_) + ((d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))])] = len((d_1508_v18_).f23)
                return _dafny.Map(coll67_)
            nw305_[int(1)] = (iife117_()
            ) | (_dafny.Map({77: p3}))
            nw305_[int(2)] = (d_1526_v33_) | (((d_1527_v36_)[self.f16] if (self.f16) in (d_1527_v36_) else d_1526_v33_))
            nw305_[int(3)] = d_1526_v33_
            nw305_[int(4)] = (d_1526_v33_) | (d_1526_v33_)
            nw305_[int(5)] = d_1526_v33_
            nw305_[int(6)] = d_1526_v33_
            nw305_[int(7)] = d_1526_v33_
            nw305_[int(8)] = d_1526_v33_
            nw305_[int(9)] = d_1526_v33_
            nw305_[int(10)] = d_1526_v33_
            nw305_[int(11)] = (d_1526_v33_) | (_dafny.Map({p3: 871}))
            nw305_[int(12)] = d_1526_v33_
            nw305_[int(13)] = _dafny.Map({(d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))]: (0) - ((d_1508_v18_).f22)})
            nw305_[int(14)] = (d_1526_v33_) | (d_1526_v33_)
            nw305_[int(15)] = d_1526_v33_
            nw305_[int(16)] = (d_1526_v33_) | (d_1526_v33_)
            nw305_[int(17)] = (d_1526_v33_) | (d_1526_v33_)
            nw305_[int(18)] = d_1526_v33_
            nw305_[int(19)] = d_1526_v33_
            nw305_[int(20)] = (_dafny.Map({p1: (d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))]})).set(len(d_1529_v37_), default__.fm0(self.f16, globalState))
            nw305_[int(21)] = _dafny.Map({(0) - (p1): len((d_1508_v18_).f23)})
            nw305_[int(22)] = d_1526_v33_
            nw305_[int(23)] = (d_1526_v33_) | ((d_1530_v38_).cf27)
            nw305_[int(24)] = ((d_1526_v33_).set(default__.fm0(False, globalState), (d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))])) | (d_1526_v33_)
            nw305_[int(25)] = (d_1526_v33_).set((0) - (len(d_1526_v33_)), -433)
            nw305_[int(26)] = d_1526_v33_
            nw305_[int(27)] = (d_1526_v33_) | (d_1526_v33_)
            d_1531_v39_ = nw305_
            index267_ = default__.safeIndex(394, (d_1531_v39_).length(0))
            (d_1531_v39_)[index267_] = ((d_1526_v33_).set(p1, (d_1509_v19_)[default__.safeIndex(436, (d_1509_v19_).length(0))])) | (d_1526_v33_)
            d_1533_v40_: _dafny.Map
            d_1533_v40_ = _dafny.Map({744: False})
            d_1534_v41_: _dafny.Seq
            d_1534_v41_ = _dafny.SeqWithoutIsStrInference([d_1490_v0_])
            d_1535_v42_: _dafny.Seq
            d_1535_v42_ = _dafny.SeqWithoutIsStrInference([(d_1534_v41_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f16, True, (self).f17, (self).f17])), len(d_1534_v41_))]])
            index268_ = default__.safeIndex(394, (d_1531_v39_).length(0))
            rhs270_ = self.f16
            rhs271_ = _dafny.Map({len(d_1533_v40_): p1})
            rhs272_ = (p2 if self.f16 else not((self).f17))
            rhs273_ = (d_1535_v42_) == (d_1535_v42_)
            rhs274_ = ((d_1508_v18_).f22) - (p1)
            lhs208_ = self
            lhs209_ = d_1531_v39_
            lhs210_ = default__.safeIndex(394, (d_1531_v39_).length(0))
            lhs211_ = self
            lhs212_ = self
            lhs213_ = globalState
            lhs208_.f16 = rhs270_
            lhs209_[lhs210_] = rhs271_
            lhs211_.f16 = rhs272_
            lhs212_.f16 = rhs273_
            lhs213_.f2 = rhs274_
        d_1536_v43_: _dafny.Seq
        d_1536_v43_ = _dafny.SeqWithoutIsStrInference([self.f16])
        d_1537_v44_: _dafny.Set
        d_1537_v44_ = _dafny.Set({False})
        d_1538_v45_: _dafny.Array
        nw306_ = _dafny.Array(None, 18)
        nw306_[int(0)] = (self).fm23(False, globalState)
        nw306_[int(1)] = p2
        nw306_[int(2)] = not(p2)
        nw306_[int(3)] = False
        nw306_[int(4)] = p2
        nw306_[int(5)] = not(not((p1) > (p1)))
        nw306_[int(6)] = p2
        nw306_[int(7)] = False
        nw306_[int(8)] = p2
        nw306_[int(9)] = not(p2)
        nw306_[int(10)] = self.f16
        nw306_[int(11)] = (d_1536_v43_)[default__.safeIndex(p1, len(d_1536_v43_))]
        nw306_[int(12)] = (d_1536_v43_)[default__.safeIndex(p1, len(d_1536_v43_))]
        nw306_[int(13)] = ((self).f17 if (self).f17 else p2)
        nw306_[int(14)] = (False) and (False)
        nw306_[int(15)] = (d_1537_v44_).isdisjoint(d_1537_v44_)
        nw306_[int(16)] = p2
        nw306_[int(17)] = (self).f17
        d_1538_v45_ = nw306_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1538_v45_).length(0)):
            d_1539_i3_: int = guard_loop_8_
            if (True) and (((0) <= (d_1539_i3_)) and ((d_1539_i3_) < ((d_1538_v45_).length(0)))):
                (d_1538_v45_)[(d_1539_i3_)] = True
        d_1540_v46_: _dafny.Map
        d_1540_v46_ = _dafny.Map({p1: 741})
        d_1541_v47_: D3
        d_1541_v47_ = D3_DC4(self.f16, p1, p1, not(self.f16), (self).f17)
        d_1542_v48_: _dafny.Map
        d_1542_v48_ = _dafny.Map({p2: (self).fm23(self.f16, globalState)})
        d_1543_v49_: _dafny.Seq
        d_1543_v49_ = _dafny.SeqWithoutIsStrInference([d_1542_v48_])
        d_1544_v50_: _dafny.Set
        d_1544_v50_ = _dafny.Set({p1, p3, len(default__.fm28((self).f17, d_1540_v46_, d_1541_v47_, (d_1543_v49_)[default__.safeIndex(p1, len(d_1543_v49_))], globalState))})
        d_1545_v51_: _dafny.Map
        d_1545_v51_ = _dafny.Map({len(d_1544_v50_): (self).f17})
        d_1546_v53_: D7
        def iife118_():
            coll68_ = _dafny.Map()
            compr_68_: int
            for compr_68_ in _dafny.IntegerRange(-134, 203):
                d_1547_v52_: int = compr_68_
                if ((-134) <= (d_1547_v52_)) and ((d_1547_v52_) < (203)):
                    coll68_[(d_1547_v52_) + (-376)] = (self).f17
            return _dafny.Map(coll68_)
        d_1546_v53_ = D7_DC11((d_1545_v51_) | (iife118_()
))
        source25_ = d_1546_v53_
        if source25_.is_DC12:
            (globalState).f2 = p3
            d_1548_v54_: str
            d_1548_v54_ = _dafny.CodePoint('d')
            d_1549_v55_: _dafny.Seq
            d_1549_v55_ = _dafny.SeqWithoutIsStrInference([default__.fm42((self).f17, len(d_1544_v50_), p1, (self).f17, globalState)])
            d_1550_v56_: _dafny.Array
            nw307_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1550_v56_ = nw307_
            d_1551_v57_: _dafny.Array
            def lambda82_(d_1552_i4_):
                return (d_1552_i4_) * (((_dafny.MultiSet([(self).f17, (self).f17])).cardinality))

            init46_ = lambda82_
            nw308_ = _dafny.Array(None, 28)
            for i0_46_ in range(nw308_.length(0)):
                nw308_[i0_46_] = init46_(i0_46_)
            d_1551_v57_ = nw308_
            index269_ = default__.safeIndex(192, (d_1550_v56_).length(0))
            (d_1550_v56_)[index269_] = d_1551_v57_
            d_1553_v58_: _dafny.Array
            def lambda83_(d_1554_i5_):
                return _dafny.SeqWithoutIsStrInference([self.f16])

            init47_ = lambda83_
            nw309_ = _dafny.Array(None, 8)
            for i0_47_ in range(nw309_.length(0)):
                nw309_[i0_47_] = init47_(i0_47_)
            d_1553_v58_ = nw309_
            d_1555_v59_: _dafny.Map
            d_1555_v59_ = _dafny.Map({d_1553_v58_: p3})
            index270_ = default__.safeIndex(192, (d_1550_v56_).length(0))
            rhs275_ = p2
            rhs276_ = _dafny.CodePoint('y')
            rhs277_ = d_1549_v55_
            rhs278_ = d_1551_v57_
            rhs279_ = ((d_1555_v59_).set(d_1553_v58_, p1)) | ((d_1555_v59_) | (d_1555_v59_))
            lhs214_ = self
            lhs215_ = d_1550_v56_
            lhs216_ = default__.safeIndex(192, (d_1550_v56_).length(0))
            lhs214_.f16 = rhs275_
            d_1548_v54_ = rhs276_
            d_1549_v55_ = rhs277_
            lhs215_[lhs216_] = rhs278_
            d_1555_v59_ = rhs279_
            r0 = p3
            (self).f16 = (self).f17
        elif True:
            d_1556___mcc_h0_ = source25_.cf15
            d_1557_cf15_ = d_1556___mcc_h0_
            (self).f16 = not ((self).f17) or ((self).f17)
            (globalState).f2 = p3
            d_1558_v60_: _dafny.Array
            nw310_ = _dafny.Array(_dafny.CodePoint('D'), 6)
            d_1558_v60_ = nw310_
            index271_ = default__.safeIndex(347, (d_1558_v60_).length(0))
            (d_1558_v60_)[index271_] = p0
            d_1559_v61_: _dafny.Array
            def lambda84_(d_1560_p1_):
                def lambda85_(d_1561_i6_):
                    return (d_1561_i6_) - (d_1560_p1_)

                return lambda85_

            init48_ = lambda84_(p1)
            nw311_ = _dafny.Array(None, 29)
            for i0_48_ in range(nw311_.length(0)):
                nw311_[i0_48_] = init48_(i0_48_)
            d_1559_v61_ = nw311_
            d_1562_v62_: D9
            d_1562_v62_ = D9_DC15((self).f17, p2, d_1544_v50_, d_1559_v61_, d_1538_v45_)
            d_1563_v63_: _dafny.Array
            nw312_ = _dafny.Array(None, 22)
            nw312_[int(0)] = D9_DC15(p2, p2, _dafny.Set({p3}), d_1559_v61_, d_1538_v45_)
            nw312_[int(1)] = d_1562_v62_
            nw312_[int(2)] = d_1562_v62_
            nw312_[int(3)] = d_1562_v62_
            nw312_[int(4)] = d_1562_v62_
            nw312_[int(5)] = D9_DC15(p2, self.f16, d_1544_v50_, d_1559_v61_, d_1538_v45_)
            nw312_[int(6)] = d_1562_v62_
            nw312_[int(7)] = d_1562_v62_
            nw312_[int(8)] = d_1562_v62_
            nw312_[int(9)] = d_1562_v62_
            nw312_[int(10)] = D9_DC15(p2, False, d_1544_v50_, d_1559_v61_, d_1538_v45_)
            nw312_[int(11)] = d_1562_v62_
            nw312_[int(12)] = d_1562_v62_
            nw312_[int(13)] = d_1562_v62_
            nw312_[int(14)] = d_1562_v62_
            nw312_[int(15)] = d_1562_v62_
            nw312_[int(16)] = d_1562_v62_
            nw312_[int(17)] = d_1562_v62_
            nw312_[int(18)] = d_1562_v62_
            nw312_[int(19)] = d_1562_v62_
            nw312_[int(20)] = d_1562_v62_
            nw312_[int(21)] = d_1562_v62_
            d_1563_v63_ = nw312_
            d_1564_v64_: D18
            d_1564_v64_ = D18_DC37(d_1563_v63_)
            d_1565_v65_: _dafny.Seq
            d_1565_v65_ = _dafny.SeqWithoutIsStrInference([(d_1564_v64_).cf56])
            d_1566_v66_: D12
            d_1566_v66_ = D12_DC24((d_1536_v43_) < (d_1536_v43_), d_1536_v43_, d_1565_v65_, (p2 if self.f16 else p2), _dafny.SeqWithoutIsStrInference([D3_DC3(p2) for d_1567_i7_ in range(default__.abs(445))]))
            index272_ = default__.safeIndex(347, (d_1558_v60_).length(0))
            rhs280_ = (d_1490_v0_)[default__.safeIndex(p3, len(d_1490_v0_))]
            rhs281_ = d_1566_v66_
            lhs217_ = d_1558_v60_
            lhs218_ = default__.safeIndex(347, (d_1558_v60_).length(0))
            lhs217_[lhs218_] = rhs280_
            d_1566_v66_ = rhs281_
            d_1568_v67_: D3
            d_1568_v67_ = D3_DC3((self).fm23(self.f16, globalState))
            d_1569_v68_: C2
            nw313_ = C2()
            nw313_.ctor__(d_1568_v67_, self.f16, (default__.fm49(p3, p3, p2, p1, globalState)).ispropersubset(d_1544_v50_))
            d_1569_v68_ = nw313_
        hi10_ = p3
        for d_1570_i8_ in range(p3, hi10_):
            d_1571_v69_: _dafny.Array
            def lambda86_(d_1572_v43_):
                def lambda87_(d_1573_i9_):
                    return (d_1572_v43_) + (d_1572_v43_)

                return lambda87_

            init49_ = lambda86_(d_1536_v43_)
            nw314_ = _dafny.Array(None, 16)
            for i0_49_ in range(nw314_.length(0)):
                nw314_[i0_49_] = init49_(i0_49_)
            d_1571_v69_ = nw314_
            index273_ = default__.safeIndex(835, (d_1571_v69_).length(0))
            (d_1571_v69_)[index273_] = _dafny.SeqWithoutIsStrInference([self.f16, p2])
            d_1574_v70_: _dafny.Array
            nw315_ = _dafny.Array(int(0), 27)
            d_1574_v70_ = nw315_
            d_1575_v71_: D9
            d_1575_v71_ = D9_DC15((self).f17, p2, d_1544_v50_, d_1574_v70_, d_1538_v45_)
            pat_let_tv22_ = d_1544_v50_
            d_1576_v72_: _dafny.Array
            nw316_ = _dafny.Array(None, 23)
            nw316_[int(0)] = d_1575_v71_
            nw316_[int(1)] = d_1575_v71_
            nw316_[int(2)] = D9_DC15((self).f17, (self).f17, d_1544_v50_, d_1574_v70_, d_1538_v45_)
            nw316_[int(3)] = d_1575_v71_
            nw316_[int(4)] = d_1575_v71_
            nw316_[int(5)] = d_1575_v71_
            nw316_[int(6)] = d_1575_v71_
            nw316_[int(7)] = d_1575_v71_
            nw316_[int(8)] = d_1575_v71_
            nw316_[int(9)] = d_1575_v71_
            def iife119_(_pat_let25_0):
                def iife120_(d_1577_dt__update__tmp_h0_):
                    def iife121_(_pat_let26_0):
                        def iife122_(d_1578_dt__update_hcf19_h0_):
                            def iife123_(_pat_let27_0):
                                def iife124_(d_1579_dt__update_hcf20_h0_):
                                    return D9_DC15((d_1577_dt__update__tmp_h0_).cf18, d_1578_dt__update_hcf19_h0_, d_1579_dt__update_hcf20_h0_, (d_1577_dt__update__tmp_h0_).cf21, (d_1577_dt__update__tmp_h0_).cf22)
                                return iife124_(_pat_let27_0)
                            return iife123_(pat_let_tv22_)
                        return iife122_(_pat_let26_0)
                    return iife121_((self).f17)
                return iife120_(_pat_let25_0)
            nw316_[int(10)] = iife119_(d_1575_v71_)
            nw316_[int(11)] = d_1575_v71_
            nw316_[int(12)] = d_1575_v71_
            nw316_[int(13)] = d_1575_v71_
            nw316_[int(14)] = D9_DC15(p2, (self).f17, d_1544_v50_, d_1574_v70_, d_1538_v45_)
            nw316_[int(15)] = d_1575_v71_
            nw316_[int(16)] = d_1575_v71_
            nw316_[int(17)] = d_1575_v71_
            nw316_[int(18)] = d_1575_v71_
            nw316_[int(19)] = d_1575_v71_
            nw316_[int(20)] = D9_DC15(p2, (self).f17, d_1544_v50_, d_1574_v70_, d_1538_v45_)
            nw316_[int(21)] = d_1575_v71_
            nw316_[int(22)] = d_1575_v71_
            d_1576_v72_ = nw316_
            d_1580_v73_: _dafny.Seq
            d_1580_v73_ = _dafny.SeqWithoutIsStrInference([d_1576_v72_])
            d_1581_v74_: D3
            d_1581_v74_ = D3_DC3(True)
            d_1582_v75_: _dafny.Seq
            d_1582_v75_ = _dafny.SeqWithoutIsStrInference([d_1581_v74_])
            d_1583_v76_: D12
            d_1583_v76_ = D12_DC24(False, d_1536_v43_, d_1580_v73_, p2, d_1582_v75_)
            index274_ = default__.safeIndex(835, (d_1571_v69_).length(0))
            (d_1571_v69_)[index274_] = ((d_1583_v76_ if (self).f17 else D12_DC24(self.f16, d_1536_v43_, d_1580_v73_, (self).f17, d_1582_v75_))).cf34
            d_1584_v77_: C1
            nw317_ = C1()
            nw317_.ctor__(self.f16, d_1570_i8_)
            d_1584_v77_ = nw317_
            (globalState).f2 = len((_dafny.SeqWithoutIsStrInference([p2])) + (_dafny.SeqWithoutIsStrInference([not(True)])))
            d_1585_v78_: _dafny.Seq
            d_1585_v78_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1586_v79_: _dafny.Seq
            d_1586_v79_ = (d_1571_v69_)[default__.safeIndex(835, (d_1571_v69_).length(0))]
            d_1587_v80_: D4
            d_1587_v80_ = D4_DC7(d_1586_v79_, p3)
            d_1588_v81_: _dafny.Map
            d_1588_v81_ = _dafny.Map({(0) - (len(d_1585_v78_)): d_1587_v80_})
            d_1589_v82_: _dafny.Map
            d_1589_v82_ = _dafny.Map({p2: len((default__.fm58(False, p2, 498, globalState)) | (d_1588_v81_))})
            d_1589_v82_ = (d_1589_v82_).set((self).f17, p1)
        d_1590_v83_: C1
        nw318_ = C1()
        nw318_.ctor__((p3) < (p1), p1)
        d_1590_v83_ = nw318_
        hi11_ = p3
        for d_1591_i10_ in range(p1, hi11_):
            d_1592_v84_: _dafny.MultiSet
            d_1592_v84_ = _dafny.MultiSet([p0])
            d_1593_v85_: _dafny.Array
            def lambda88_(d_1594_i11_):
                return default__.safeModuloInt(d_1594_i11_, -761)

            init50_ = lambda88_
            nw319_ = _dafny.Array(None, 10)
            for i0_50_ in range(nw319_.length(0)):
                nw319_[i0_50_] = init50_(i0_50_)
            d_1593_v85_ = nw319_
            d_1595_v86_: D9
            d_1595_v86_ = D9_DC15(default__.fm27((d_1592_v84_).cardinality, d_1591_i10_, p3, _dafny.Map({-678: 9}), globalState), p2, d_1544_v50_, d_1593_v85_, d_1538_v45_)
            source26_ = d_1595_v86_
            if source26_.is_DC15:
                d_1596___mcc_h1_ = source26_.cf18
                d_1597___mcc_h2_ = source26_.cf19
                d_1598___mcc_h3_ = source26_.cf20
                d_1599___mcc_h4_ = source26_.cf21
                d_1600___mcc_h5_ = source26_.cf22
                d_1601_cf22_ = d_1600___mcc_h5_
                d_1602_cf21_ = d_1599___mcc_h4_
                d_1603_cf20_ = d_1598___mcc_h3_
                d_1604_cf19_ = d_1597___mcc_h2_
                d_1605_cf18_ = d_1596___mcc_h1_
                d_1606_v87_: D3
                d_1606_v87_ = D3_DC4(d_1590_v83_.f25, (0) - (d_1591_i10_), d_1591_i10_, not(default__.fm31(d_1590_v83_.f25, True, globalState)), d_1605_cf18_)
                d_1607_v88_: D3
                d_1607_v88_ = D3_DC5(d_1606_v87_)
                d_1608_v89_: D3
                d_1608_v89_ = D3_DC5(d_1607_v88_)
                d_1609_v90_: _dafny.Map
                d_1609_v90_ = _dafny.Map({d_1608_v89_: not(d_1605_cf18_)})
                d_1610_v91_: D3
                d_1610_v91_ = D3_DC4(p2, p1, 965, self.f16, ((d_1609_v90_)[d_1608_v89_] if (d_1608_v89_) in (d_1609_v90_) else d_1604_cf19_))
                d_1611_v92_: D3
                d_1611_v92_ = D3_DC5(d_1607_v88_)
                d_1608_v89_ = d_1608_v89_
                d_1612_v93_: _dafny.Seq
                d_1612_v93_ = d_1536_v43_
                d_1613_v94_: D4
                d_1613_v94_ = D4_DC7(d_1612_v93_, default__.safeDivisionInt(p1, p1))
                d_1613_v94_ = d_1613_v94_
                (globalState).f7 = (-594) - (((d_1540_v46_)[p3] if (p3) in (d_1540_v46_) else len(_dafny.Map({False: d_1590_v83_.f25}))))
                (globalState).f7 = (self).fm12(len(d_1537_v44_), globalState)
            elif source26_.is_DC16:
                d_1614___mcc_h6_ = source26_.cf23
                d_1615___mcc_h7_ = source26_.cf24
                d_1616_cf24_ = d_1615___mcc_h7_
                d_1617_cf23_ = d_1614___mcc_h6_
                d_1618_v95_: _dafny.Array
                nw320_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                d_1618_v95_ = nw320_
                d_1618_v95_ = d_1618_v95_
                index275_ = default__.safeIndex(8, (d_1593_v85_).length(0))
                (d_1593_v85_)[index275_] = 974
                d_1619_v96_: _dafny.Seq
                d_1619_v96_ = _dafny.SeqWithoutIsStrInference([len(d_1490_v0_)])
                d_1620_v97_: T0
                nw321_ = C1()
                nw321_.ctor__(self.f16, len(d_1619_v96_))
                d_1620_v97_ = nw321_
                index276_ = default__.safeIndex(8, (d_1593_v85_).length(0))
                (d_1593_v85_)[index276_] = (default__.safeDivisionInt(p1, len(d_1490_v0_)) if p2 else default__.safeDivisionInt(len(_dafny.Map({default__.fm31(p2, p2, globalState): d_1620_v97_})), p1))
                d_1621_v98_: str
                d_1621_v98_ = _dafny.CodePoint('p')
                d_1621_v98_ = p0
                (globalState).f2 = (p1) + (len(d_1545_v51_))
            elif source26_.is_DC14:
                d_1622___mcc_h8_ = source26_.cf17
                d_1623_cf17_ = d_1622___mcc_h8_
                d_1624_v99_: _dafny.Array
                def lambda89_(d_1625_p3_):
                    def lambda90_(d_1626_i12_):
                        return _dafny.SeqWithoutIsStrInference([d_1625_p3_ for d_1627_i13_ in range(default__.abs(425))])

                    return lambda90_

                init51_ = lambda89_(p3)
                nw322_ = _dafny.Array(None, 23)
                for i0_51_ in range(nw322_.length(0)):
                    nw322_[i0_51_] = init51_(i0_51_)
                d_1624_v99_ = nw322_
                d_1628_v100_: _dafny.Map
                d_1628_v100_ = _dafny.Map({d_1624_v99_: (p1) + (p1)})
                d_1629_v101_: _dafny.MultiSet
                d_1629_v101_ = _dafny.MultiSet([d_1490_v0_])
                d_1628_v100_ = (d_1628_v100_).set(d_1624_v99_, ((d_1629_v101_)[d_1490_v0_] if (d_1490_v0_) in (d_1629_v101_) else len(d_1545_v51_)))
                d_1630_v102_: D9
                d_1630_v102_ = D9_DC16(p3, p2)
                def iife125_(_pat_let28_0):
                    def iife126_(d_1631_dt__update__tmp_h1_):
                        def iife127_(_pat_let29_0):
                            def iife128_(d_1632_dt__update_hcf24_h0_):
                                return D9_DC16((d_1631_dt__update__tmp_h1_).cf23, d_1632_dt__update_hcf24_h0_)
                            return iife128_(_pat_let29_0)
                        return iife127_((self).f17)
                    return iife126_(_pat_let28_0)
                r0 = default__.safeDivisionInt(p1, (0) - ((iife125_(d_1630_v102_)).cf23))
                rhs282_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1633_i14_ in range(default__.abs(524))])
                rhs283_ = (d_1591_i10_) - ((p1) + (p3))
                lhs219_ = globalState
                d_1490_v0_ = rhs282_
                lhs219_.f7 = rhs283_
                (d_1590_v83_).f25 = False
            elif True:
                d_1634___mcc_h9_ = source26_.cf25
                d_1635_cf25_ = d_1634___mcc_h9_
                rhs284_ = p1
                rhs285_ = (default__.fm40(d_1590_v83_.f25, globalState)) | ((d_1542_v48_) | (d_1542_v48_))
                rhs286_ = (d_1538_v45_ if not(not(p2)) else d_1538_v45_)
                lhs220_ = globalState
                lhs220_.f7 = rhs284_
                d_1542_v48_ = rhs285_
                d_1538_v45_ = rhs286_
                d_1636_v103_: _dafny.MultiSet
                d_1636_v103_ = _dafny.MultiSet([d_1591_i10_])
                d_1637_v104_: D19
                d_1637_v104_ = D19_DC39(d_1636_v103_)
                d_1638_v105_: D19
                d_1638_v105_ = D19_DC39((d_1637_v104_).cf60)
                (globalState).f7 = (((d_1636_v103_) | ((d_1638_v105_).cf60)) | ((d_1636_v103_) | (d_1636_v103_))).cardinality
                index277_ = default__.safeIndex(729, (d_1593_v85_).length(0))
                (d_1593_v85_)[index277_] = p1
                d_1639_v106_: _dafny.Array
                nw323_ = _dafny.Array(D9.default()(), 12)
                d_1639_v106_ = nw323_
                d_1640_v107_: _dafny.Seq
                d_1640_v107_ = _dafny.SeqWithoutIsStrInference([d_1639_v106_, d_1639_v106_, d_1639_v106_])
                d_1641_v108_: D3
                d_1641_v108_ = D3_DC3(d_1590_v83_.f25)
                d_1642_v109_: _dafny.Seq
                d_1642_v109_ = _dafny.SeqWithoutIsStrInference([d_1641_v108_, d_1641_v108_, D3_DC3(not(p2)), D3_DC3(True), d_1641_v108_])
                d_1643_v110_: D12
                d_1643_v110_ = D12_DC24((self).f17, d_1536_v43_, d_1640_v107_, d_1590_v83_.f25, d_1642_v109_)
                d_1644_v111_: _dafny.MultiSet
                d_1644_v111_ = _dafny.MultiSet([d_1643_v110_, d_1643_v110_, d_1643_v110_, d_1643_v110_, d_1643_v110_])
                d_1645_v112_: D20
                d_1645_v112_ = D20_DC41((_dafny.MultiSet([d_1643_v110_, d_1643_v110_])).set(d_1643_v110_, default__.abs(len(d_1490_v0_))))
                d_1646_v113_: _dafny.MultiSet
                d_1646_v113_ = _dafny.MultiSet([d_1590_v83_.f25, d_1590_v83_.f25])
                d_1647_v114_: _dafny.Map
                d_1647_v114_ = _dafny.Map({d_1646_v113_: self.f16})
                index278_ = default__.safeIndex(729, (d_1593_v85_).length(0))
                rhs287_ = not(not((d_1536_v43_)[default__.safeIndex(p1, len(d_1536_v43_))]))
                rhs288_ = ((d_1644_v111_) - ((d_1645_v112_).cf64)).ispropersubset(_dafny.MultiSet([d_1643_v110_, D12_DC24(p2, d_1536_v43_, d_1640_v107_, d_1590_v83_.f25, d_1642_v109_)]))
                rhs289_ = default__.safeModuloInt(p3, len((d_1647_v114_) | (d_1647_v114_)))
                lhs221_ = d_1590_v83_
                lhs222_ = d_1590_v83_
                lhs223_ = d_1593_v85_
                lhs224_ = default__.safeIndex(729, (d_1593_v85_).length(0))
                lhs221_.f25 = rhs287_
                lhs222_.f25 = rhs288_
                lhs223_[lhs224_] = rhs289_
                d_1648_v115_: D4
                d_1648_v115_ = D4_DC6(d_1542_v48_)
                d_1649_v116_: T0
                nw324_ = C1()
                nw324_.ctor__(p2, (d_1593_v85_)[default__.safeIndex(729, (d_1593_v85_).length(0))])
                d_1649_v116_ = nw324_
                d_1650_v117_: D16
                d_1650_v117_ = D16_DC34((self).f17, d_1649_v116_, p1)
                d_1651_v118_: D16
                d_1651_v118_ = D16_DC34(p2, (d_1650_v117_).cf52, d_1649_v116_.f14)
                rhs290_ = p1
                rhs291_ = default__.safeDivisionInt(default__.safeDivisionInt(len(d_1542_v48_), len((d_1648_v115_).cf10)), (d_1650_v117_).cf53)
                rhs292_ = p2
                rhs293_ = not(d_1590_v83_.f25)
                rhs294_ = d_1649_v116_.f14
                lhs225_ = globalState
                lhs226_ = self
                lhs227_ = d_1590_v83_
                lhs228_ = globalState
                lhs225_.f7 = rhs290_
                r0 = rhs291_
                lhs226_.f16 = rhs292_
                lhs227_.f25 = rhs293_
                lhs228_.f2 = rhs294_
            d_1652_v119_: _dafny.Map
            d_1652_v119_ = _dafny.Map({d_1592_v84_: d_1538_v45_})
            d_1653_v120_: _dafny.Array
            nw325_ = _dafny.Array(None, 4)
            nw325_[int(0)] = d_1538_v45_
            nw325_[int(1)] = d_1538_v45_
            nw325_[int(2)] = d_1538_v45_
            nw325_[int(3)] = ((d_1652_v119_)[d_1592_v84_] if (d_1592_v84_) in (d_1652_v119_) else d_1538_v45_)
            d_1653_v120_ = nw325_
            d_1653_v120_ = d_1653_v120_
            d_1654_v121_: C4
            nw326_ = C4()
            nw326_.ctor__(d_1590_v83_.f25, self.f16, d_1590_v83_.f25)
            d_1654_v121_ = nw326_
            index279_ = default__.safeIndex(105, (d_1538_v45_).length(0))
            (d_1538_v45_)[index279_] = self.f16
            index280_ = default__.safeIndex(105, (d_1538_v45_).length(0))
            (d_1538_v45_)[index280_] = d_1590_v83_.f25
        r0 = len((d_1536_v43_) + (_dafny.SeqWithoutIsStrInference([(self).f17])))
        return r0


class C6(T1, T0, T2):
    def  __init__(self):
        self._f16: bool = False
        self._f14: int = int(0)
        self._f17: bool = False
        self._f19: int = int(0)
        self._f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f19, f20, f16, f17, f14):
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f16 = f16
        (self)._f17 = f17
        (self).f14 = f14

    def fm3(self, globalState):
        def iife129_():
            coll69_ = _dafny.Map()
            compr_69_: int
            for compr_69_ in _dafny.IntegerRange(425, -562):
                d_1655_v0_: int = compr_69_
                if ((425) <= (d_1655_v0_)) and ((d_1655_v0_) < (-562)):
                    coll69_[(d_1655_v0_) - ((self).f19)] = (self).f19
            return _dafny.Map(coll69_)
        return iife129_()
        

    def fm12(self, p0, globalState):
        return (0) - (default__.safeModuloInt((self).f19, (self).f19))

    def fm13(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([self.f16, (self).f20])) + ((_dafny.SeqWithoutIsStrInference([self.f16])) + (_dafny.SeqWithoutIsStrInference([(self).f20, (self).f20, self.f16])))

    def fm18(self, p0, p1, p2, p3, globalState):
        source27_ = D4_DC6(_dafny.Map({self.f16: (self).f20}))
        if source27_.is_DC7:
            d_1656___mcc_h0_ = source27_.cf11
            d_1657___mcc_h1_ = source27_.cf12
            d_1658_cf12_ = d_1657___mcc_h1_
            d_1659_cf11_ = d_1656___mcc_h0_
            return (d_1658_cf12_) + (self.f14)
        elif True:
            d_1660___mcc_h2_ = source27_.cf10
            d_1661_cf10_ = d_1660___mcc_h2_
            return (self).f19

    def fm19(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(0) - (self.f14), self.f14])

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        hi12_ = ((self).f19) + ((self).f19)
        for d_1662_i0_ in range((self).f19, hi12_):
            d_1663_v0_: _dafny.Seq
            d_1663_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usc"))
            d_1663_v0_ = (d_1663_v0_) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1664_i1_ in range(default__.abs(425))]) if self.f16 else d_1663_v0_))
            d_1665_v1_: _dafny.Array
            nw327_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_1665_v1_ = nw327_
            d_1665_v1_ = d_1665_v1_
            (globalState).f2 = 495
            d_1666_v2_: _dafny.Array
            nw328_ = _dafny.Array(int(0), 17)
            d_1666_v2_ = nw328_
            index281_ = default__.safeIndex(435, (d_1666_v2_).length(0))
            (d_1666_v2_)[index281_] = p0
            d_1667_v3_: D3
            d_1667_v3_ = D3_DC3(p1)
            d_1668_v4_: _dafny.Set
            d_1668_v4_ = _dafny.Set({(d_1667_v3_).cf3, not((self).f17)})
            d_1669_v5_: _dafny.Seq
            d_1669_v5_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1668_v4_)), len(d_1668_v4_)])
            d_1670_v6_: _dafny.Map
            d_1670_v6_ = _dafny.Map({len(d_1669_v5_): self.f14})
            d_1671_v7_: _dafny.MultiSet
            d_1671_v7_ = _dafny.MultiSet([d_1670_v6_, d_1670_v6_, (d_1670_v6_) | (d_1670_v6_)])
            d_1672_v8_: str
            d_1672_v8_ = _dafny.CodePoint('m')
            d_1673_v9_: _dafny.Map
            d_1673_v9_ = _dafny.Map({_dafny.MultiSet([True]): d_1672_v8_})
            d_1674_v10_: _dafny.Seq
            d_1674_v10_ = _dafny.SeqWithoutIsStrInference([(self).f20])
            d_1675_v11_: _dafny.Seq
            d_1675_v11_ = _dafny.SeqWithoutIsStrInference([d_1674_v10_, _dafny.SeqWithoutIsStrInference([p1])])
            index282_ = default__.safeIndex(435, (d_1666_v2_).length(0))
            rhs295_ = default__.safeDivisionInt((0) - ((len(d_1668_v4_)) + (p0)), default__.safeModuloInt(p0, len(d_1673_v9_)))
            rhs296_ = not((D3_DC4(self.f16, (self).f19, p0, (self).f20, False)).cf8)
            rhs297_ = default__.fm20(d_1662_i0_, len((d_1675_v11_)[default__.safeIndex(self.f14, len(d_1675_v11_))]), globalState)
            rhs298_ = False
            lhs229_ = d_1666_v2_
            lhs230_ = default__.safeIndex(435, (d_1666_v2_).length(0))
            lhs231_ = self
            lhs232_ = self
            lhs229_[lhs230_] = rhs295_
            lhs231_.f16 = rhs296_
            d_1671_v7_ = rhs297_
            lhs232_.f16 = rhs298_
        d_1676_i2_: int
        d_1676_i2_ = 0
        with _dafny.label("7"):
            while ((p0) >= (p0)) or (p1):
                with _dafny.c_label("7"):
                    if (d_1676_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_1676_i2_ = (d_1676_i2_) + (1)
                    d_1677_v12_: _dafny.Array
                    nw329_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                    d_1677_v12_ = nw329_
                    d_1678_v13_: _dafny.Map
                    d_1678_v13_ = _dafny.Map({False: default__.fm21(globalState)})
                    d_1679_v14_: _dafny.Map
                    d_1679_v14_ = _dafny.Map({self.f14: d_1678_v13_})
                    d_1680_v15_: _dafny.Seq
                    d_1680_v15_ = _dafny.SeqWithoutIsStrInference([(self).f17, p1])
                    rhs299_ = d_1677_v12_
                    rhs300_ = (self).fm18(True, 50, d_1679_v14_, (self).f19, globalState)
                    rhs301_ = (default__.fm21(globalState)) in (d_1680_v15_)
                    lhs233_ = globalState
                    lhs234_ = self
                    d_1677_v12_ = rhs299_
                    lhs233_.f2 = rhs300_
                    lhs234_.f16 = rhs301_
                    (self).f16 = not ((self).f20) or ((self).f20)
                    d_1681_v16_: str
                    d_1681_v16_ = _dafny.CodePoint('v')
                    d_1682_v17_: D6
                    d_1682_v17_ = D6_DC9(d_1681_v16_)
                    d_1681_v16_ = (d_1682_v17_).cf14
                    (globalState).f7 = p0
                    pass
            pass
        d_1683_v18_: _dafny.Array
        def lambda91_(d_1684_i3_):
            return (self).f20

        init52_ = lambda91_
        nw330_ = _dafny.Array(None, 24)
        for i0_52_ in range(nw330_.length(0)):
            nw330_[i0_52_] = init52_(i0_52_)
        d_1683_v18_ = nw330_
        d_1685_v19_: _dafny.Map
        d_1685_v19_ = _dafny.Map({d_1683_v18_: ((self).f19) - (self.f14)})
        d_1686_v20_: _dafny.Seq
        d_1686_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yc"))
        d_1687_v21_: _dafny.MultiSet
        d_1687_v21_ = _dafny.MultiSet([(self).f19, p0, self.f14])
        d_1688_v22_: _dafny.Seq
        d_1688_v22_ = _dafny.SeqWithoutIsStrInference([len(d_1686_v20_), self.f14, self.f14, (d_1687_v21_).cardinality])
        d_1689_v23_: _dafny.Seq
        d_1689_v23_ = _dafny.SeqWithoutIsStrInference([d_1683_v18_])
        d_1690_v24_: _dafny.Seq
        d_1690_v24_ = _dafny.SeqWithoutIsStrInference([(d_1689_v23_)[default__.safeIndex(self.f14, len(d_1689_v23_))]])
        d_1691_v25_: _dafny.Map
        d_1691_v25_ = _dafny.Map({default__.fm22(not((self).f20), globalState): d_1686_v20_})
        rhs302_ = d_1683_v18_
        rhs303_ = _dafny.Map({(d_1690_v24_)[default__.safeIndex(self.f14, len(d_1690_v24_))]: (0) - (p0)})
        rhs304_ = self.f14
        rhs305_ = (d_1688_v22_) + ((d_1688_v22_) + (_dafny.SeqWithoutIsStrInference([(self).f19 for d_1692_i4_ in range(default__.abs(776))])))
        rhs306_ = default__.safeModuloInt((self).fm12((d_1688_v22_)[default__.safeIndex(len(d_1688_v22_), len(d_1688_v22_))], globalState), len(d_1691_v25_))
        lhs235_ = self
        lhs236_ = globalState
        d_1683_v18_ = rhs302_
        d_1685_v19_ = rhs303_
        lhs235_.f14 = rhs304_
        d_1688_v22_ = rhs305_
        lhs236_.f7 = rhs306_
        d_1693_v26_: _dafny.Array
        def lambda92_(d_1694_i5_):
            return _dafny.CodePoint('j')

        init53_ = lambda92_
        nw331_ = _dafny.Array(None, 2)
        for i0_53_ in range(nw331_.length(0)):
            nw331_[i0_53_] = init53_(i0_53_)
        d_1693_v26_ = nw331_
        d_1695_v27_: str
        d_1695_v27_ = _dafny.CodePoint('k')
        index283_ = default__.safeIndex(53, (d_1693_v26_).length(0))
        (d_1693_v26_)[index283_] = d_1695_v27_
        index284_ = default__.safeIndex(53, (d_1693_v26_).length(0))
        (d_1693_v26_)[index284_] = d_1695_v27_
        source28_ = D6_DC9(_dafny.CodePoint('x'))
        if source28_.is_DC10:
            d_1696_v28_: D6
            d_1696_v28_ = D6_DC10()
            d_1697_v29_: _dafny.Set
            d_1697_v29_ = _dafny.Set({len(d_1686_v20_)})
            d_1698_v30_: _dafny.Set
            d_1698_v30_ = _dafny.Set({d_1697_v29_})
            d_1699_v31_: _dafny.Array
            nw332_ = _dafny.Array(None, 2)
            nw332_[int(0)] = len(d_1698_v30_)
            nw332_[int(1)] = p0
            d_1699_v31_ = nw332_
            d_1700_v32_: _dafny.Seq
            d_1700_v32_ = _dafny.SeqWithoutIsStrInference([d_1699_v31_, d_1699_v31_, d_1699_v31_])
            d_1701_v33_: _dafny.Map
            d_1701_v33_ = _dafny.Map({d_1696_v28_: (d_1700_v32_)[default__.safeIndex(self.f14, len(d_1700_v32_))]})
            d_1701_v33_ = (d_1701_v33_ if True else d_1701_v33_)
            d_1702_v34_: T2
            nw333_ = C4()
            nw333_.ctor__((self).f17, self.f16, p1)
            d_1702_v34_ = nw333_
            d_1703_v35_: _dafny.Seq
            d_1703_v35_ = _dafny.SeqWithoutIsStrInference([d_1702_v34_, d_1702_v34_])
            d_1703_v35_ = (d_1703_v35_) + (d_1703_v35_)
            d_1704_v36_: _dafny.Map
            d_1704_v36_ = _dafny.Map({_dafny.MultiSet(d_1688_v22_): (d_1686_v20_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiktmgg")))})
            d_1705_v37_: _dafny.Map
            d_1705_v37_ = _dafny.Map({p0: p1})
            d_1706_v38_: _dafny.Map
            d_1706_v38_ = _dafny.Map({(d_1693_v26_)[default__.safeIndex(53, (d_1693_v26_).length(0))]: p1})
            rhs307_ = ((d_1704_v36_) | (d_1704_v36_)).set(d_1687_v21_, (d_1686_v20_ if ((d_1705_v37_)[len(d_1706_v38_)] if (len(d_1706_v38_)) in (d_1705_v37_) else p1) else d_1686_v20_))
            rhs308_ = p0
            rhs309_ = len((d_1686_v20_) + (d_1686_v20_))
            lhs237_ = globalState
            lhs238_ = globalState
            d_1704_v36_ = rhs307_
            lhs237_.f2 = rhs308_
            lhs238_.f7 = rhs309_
            d_1707_v39_: int
            d_1708_v40_: bool
            d_1709_v41_: _dafny.Seq
            d_1710_v42_: bool
            out29_: int
            out30_: bool
            out31_: _dafny.Seq
            out32_: bool
            out29_, out30_, out31_, out32_ = (d_1702_v34_).m6((self).f17, not((d_1702_v34_).f17), p0, (d_1693_v26_)[default__.safeIndex(53, (d_1693_v26_).length(0))], globalState)
            d_1707_v39_ = out29_
            d_1708_v40_ = out30_
            d_1709_v41_ = out31_
            d_1710_v42_ = out32_
        elif True:
            d_1711___mcc_h0_ = source28_.cf14
            d_1712_cf14_ = d_1711___mcc_h0_
            d_1713_v43_: _dafny.Map
            d_1713_v43_ = _dafny.Map({(self).f17: p1})
            d_1686_v20_ = (d_1686_v20_) + ((default__.fm42(p1, (self).f19, len(d_1713_v43_), (self).f17, globalState)) + (d_1686_v20_))
            (globalState).f7 = p0
            d_1714_v44_: _dafny.Array
            def lambda93_(d_1715_i6_):
                return (d_1715_i6_) - (self.f14)

            init54_ = lambda93_
            nw334_ = _dafny.Array(None, 1)
            for i0_54_ in range(nw334_.length(0)):
                nw334_[i0_54_] = init54_(i0_54_)
            d_1714_v44_ = nw334_
            d_1714_v44_ = d_1714_v44_
            d_1716_v45_: _dafny.Seq
            d_1716_v45_ = _dafny.SeqWithoutIsStrInference([(self).f20])
            d_1717_v46_: _dafny.Map
            d_1717_v46_ = _dafny.Map({self.f14: (self).f20})
            d_1716_v45_ = default__.fm52(((d_1717_v46_)[self.f14] if (self.f14) in (d_1717_v46_) else p1), (0) - ((p0) - (p0)), len(_dafny.Map({False: (self).f20})), (0) - (default__.fm0((self).f17, globalState)), globalState)
        d_1718_v47_: _dafny.Array
        def lambda94_(d_1719_p0_):
            def lambda95_(d_1720_i7_):
                return default__.safeModuloInt(d_1720_i7_, d_1719_p0_)

            return lambda95_

        init55_ = lambda94_(p0)
        nw335_ = _dafny.Array(None, 16)
        for i0_55_ in range(nw335_.length(0)):
            nw335_[i0_55_] = init55_(i0_55_)
        d_1718_v47_ = nw335_
        d_1718_v47_ = d_1718_v47_
        r0 = self.f14
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_1721_v0_: _dafny.Seq
        d_1721_v0_ = _dafny.SeqWithoutIsStrInference([(self).f17])
        d_1722_v1_: _dafny.Seq
        d_1722_v1_ = d_1721_v0_
        source29_ = d_1722_v1_
        d_1723___mcc_h0_ = source29_
        d_1724_cf0_ = d_1723___mcc_h0_
        d_1725_v2_: _dafny.Array
        def lambda96_(d_1726_i0_):
            return (d_1726_i0_) * ((self).f19)

        init56_ = lambda96_
        nw336_ = _dafny.Array(None, 21)
        for i0_56_ in range(nw336_.length(0)):
            nw336_[i0_56_] = init56_(i0_56_)
        d_1725_v2_ = nw336_
        d_1727_v3_: _dafny.MultiSet
        d_1727_v3_ = _dafny.MultiSet([d_1725_v2_, d_1725_v2_])
        d_1728_v4_: C4
        nw337_ = C4()
        nw337_.ctor__((((d_1727_v3_)[d_1725_v2_] if (d_1725_v2_) in (d_1727_v3_) else self.f14)) != ((self).f19), (False if self.f16 else p2), (self).f17)
        d_1728_v4_ = nw337_
        d_1729_v5_: C1
        nw338_ = C1()
        nw338_.ctor__(self.f16, p3)
        d_1729_v5_ = nw338_
        d_1730_v6_: _dafny.Array
        nw339_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_1730_v6_ = nw339_
        d_1731_v7_: _dafny.Array
        nw340_ = _dafny.Array(D15.default()(), 26)
        d_1731_v7_ = nw340_
        d_1732_v8_: _dafny.Seq
        d_1732_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gumglvsn"))
        d_1733_v9_: _dafny.Seq
        d_1733_v9_ = _dafny.SeqWithoutIsStrInference([p1, ((self).f19) + (p3), p3])
        d_1734_v10_: _dafny.Seq
        d_1734_v10_ = _dafny.SeqWithoutIsStrInference([d_1731_v7_, d_1731_v7_, d_1731_v7_])
        d_1735_v11_: _dafny.Seq
        d_1735_v11_ = _dafny.SeqWithoutIsStrInference([d_1731_v7_, d_1731_v7_, d_1731_v7_, (d_1734_v10_)[default__.safeIndex(p1, len(d_1734_v10_))]])
        rhs310_ = -202
        rhs311_ = ((d_1732_v8_) <= (d_1732_v8_)) or (not((self).f20))
        rhs312_ = (d_1733_v9_)[default__.safeIndex((self).f19, len(d_1733_v9_))]
        rhs313_ = d_1730_v6_
        rhs314_ = (d_1735_v11_)[default__.safeIndex(p3, len(d_1735_v11_))]
        lhs239_ = globalState
        lhs240_ = d_1729_v5_
        lhs241_ = globalState
        lhs239_.f2 = rhs310_
        lhs240_.f25 = rhs311_
        lhs241_.f7 = rhs312_
        d_1730_v6_ = rhs313_
        d_1731_v7_ = rhs314_
        (d_1729_v5_).f25 = self.f16
        d_1736_i1_: int
        d_1736_i1_ = 0
        with _dafny.label("8"):
            while (self).f17:
                with _dafny.c_label("8"):
                    if (d_1736_i1_) >= (100):
                        raise _dafny.Break("8")
                    d_1736_i1_ = (d_1736_i1_) + (1)
                    (self).f14 = p3
                    d_1737_v12_: D9
                    d_1737_v12_ = D9_DC16(self.f14, self.f16)
                    d_1738_v13_: _dafny.Seq
                    d_1738_v13_ = _dafny.SeqWithoutIsStrInference([self.f14, (d_1737_v12_).cf23, p3, (0) - (self.f14), default__.fm0(p2, globalState)])
                    d_1738_v13_ = d_1738_v13_
                    d_1739_v14_: D4
                    d_1739_v14_ = D4_DC7(d_1722_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpuicrr"))))
                    pat_let_tv23_ = p3
                    pat_let_tv24_ = d_1722_v1_
                    pat_let_tv25_ = d_1722_v1_
                    d_1740_v15_: _dafny.Array
                    nw341_ = _dafny.Array(None, 21)
                    nw341_[int(0)] = d_1739_v14_
                    def iife130_(_pat_let30_0):
                        def iife131_(d_1741_dt__update__tmp_h0_):
                            def iife132_(_pat_let31_0):
                                def iife133_(d_1742_dt__update_hcf12_h0_):
                                    return D4_DC7((d_1741_dt__update__tmp_h0_).cf11, d_1742_dt__update_hcf12_h0_)
                                return iife133_(_pat_let31_0)
                            return iife132_(pat_let_tv23_)
                        return iife131_(_pat_let30_0)
                    nw341_[int(1)] = iife130_(D4_DC7(_dafny.SeqWithoutIsStrInference([default__.fm31(p2, (self).f17, globalState)]), p1))
                    nw341_[int(2)] = D4_DC7(d_1722_v1_, (self).f19)
                    nw341_[int(3)] = d_1739_v14_
                    nw341_[int(4)] = d_1739_v14_
                    nw341_[int(5)] = default__.fm59(globalState)
                    nw341_[int(6)] = d_1739_v14_
                    nw341_[int(7)] = D4_DC7(d_1722_v1_, p3)
                    nw341_[int(8)] = D4_DC7(d_1722_v1_, p1)
                    nw341_[int(9)] = d_1739_v14_
                    nw341_[int(10)] = d_1739_v14_
                    nw341_[int(11)] = D4_DC7(d_1722_v1_, (self).f19)
                    nw341_[int(12)] = d_1739_v14_
                    nw341_[int(13)] = D4_DC7(d_1722_v1_, (self).f19)
                    def iife134_(_pat_let32_0):
                        def iife135_(d_1743_dt__update__tmp_h1_):
                            def iife136_(_pat_let33_0):
                                def iife137_(d_1744_dt__update_hcf11_h0_):
                                    return D4_DC7(d_1744_dt__update_hcf11_h0_, (d_1743_dt__update__tmp_h1_).cf12)
                                return iife137_(_pat_let33_0)
                            return iife136_(pat_let_tv24_)
                        return iife135_(_pat_let32_0)
                    nw341_[int(14)] = iife134_(default__.fm59(globalState))
                    nw341_[int(15)] = d_1739_v14_
                    nw341_[int(16)] = d_1739_v14_
                    nw341_[int(17)] = d_1739_v14_
                    def iife138_(_pat_let34_0):
                        def iife139_(d_1745_dt__update__tmp_h2_):
                            def iife140_(_pat_let35_0):
                                def iife141_(d_1746_dt__update_hcf11_h1_):
                                    return D4_DC7(d_1746_dt__update_hcf11_h1_, (d_1745_dt__update__tmp_h2_).cf12)
                                return iife141_(_pat_let35_0)
                            return iife140_(pat_let_tv25_)
                        return iife139_(_pat_let34_0)
                    nw341_[int(18)] = iife138_(d_1739_v14_)
                    nw341_[int(19)] = default__.fm59(globalState)
                    nw341_[int(20)] = d_1739_v14_
                    d_1740_v15_ = nw341_
                    index285_ = default__.safeIndex(896, (d_1740_v15_).length(0))
                    (d_1740_v15_)[index285_] = D4_DC7(d_1722_v1_, p1)
                    index286_ = default__.safeIndex(896, (d_1740_v15_).length(0))
                    (d_1740_v15_)[index286_] = d_1739_v14_
                    d_1747_v16_: _dafny.MultiSet
                    d_1747_v16_ = _dafny.MultiSet([(self).f19])
                    d_1747_v16_ = ((d_1747_v16_).intersection(_dafny.MultiSet((d_1738_v13_).set(default__.safeIndex((self).f19, len(d_1738_v13_)), -990)))).set(p1, default__.abs(p3))
                    pass
            pass
        d_1748_v17_: _dafny.Array
        def lambda97_(d_1749_i2_):
            return (self).f20

        init57_ = lambda97_
        nw342_ = _dafny.Array(None, 1)
        for i0_57_ in range(nw342_.length(0)):
            nw342_[i0_57_] = init57_(i0_57_)
        d_1748_v17_ = nw342_
        d_1750_v18_: _dafny.Seq
        d_1750_v18_ = _dafny.SeqWithoutIsStrInference([p0, _dafny.CodePoint('b')])
        index287_ = default__.safeIndex(245, (d_1748_v17_).length(0))
        (d_1748_v17_)[index287_] = default__.fm27(p3, 599, (_dafny.MultiSet(d_1750_v18_)).cardinality, _dafny.Map({default__.fm0(p2, globalState): len(d_1721_v0_)}), globalState)
        d_1751_v19_: D12
        d_1751_v19_ = D12_DC23((self).f17, (self).fm12(p3, globalState), p3)
        d_1752_v20_: _dafny.Map
        d_1752_v20_ = _dafny.Map({p3: d_1751_v19_})
        index288_ = default__.safeIndex(245, (d_1748_v17_).length(0))
        rhs315_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm"))) < (d_1750_v18_)
        rhs316_ = (self).f20
        rhs317_ = (706) * ((387 if self.f16 else len(d_1752_v20_)))
        rhs318_ = default__.safeModuloInt(self.f14, ((self).f19) * (p1))
        rhs319_ = (not (self.f16) or ((self).f20) if (d_1721_v0_)[default__.safeIndex(p3, len(d_1721_v0_))] else (self).f17)
        lhs242_ = self
        lhs243_ = self
        lhs244_ = globalState
        lhs245_ = self
        lhs246_ = d_1748_v17_
        lhs247_ = default__.safeIndex(245, (d_1748_v17_).length(0))
        lhs242_.f16 = rhs315_
        lhs243_.f16 = rhs316_
        lhs244_.f7 = rhs317_
        lhs245_.f14 = rhs318_
        lhs246_[lhs247_] = rhs319_
        d_1753_v21_: _dafny.Map
        d_1753_v21_ = _dafny.Map({p3: p2})
        (self).f16 = ((self).f17) == (((d_1753_v21_)[(self).f19] if ((self).f19) in (d_1753_v21_) else p2))
        index289_ = default__.safeIndex(245, (d_1748_v17_).length(0))
        (d_1748_v17_)[index289_] = not((self).f20)
        d_1754_v22_: _dafny.MultiSet
        d_1754_v22_ = _dafny.MultiSet([p1])
        d_1754_v22_ = d_1754_v22_
        r0 = (p3) - ((0) - ((0) - (self.f14)))
        return r0

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: bool = False
        (globalState).f7 = self.f14
        d_1755_v0_: _dafny.Seq
        d_1755_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
        hi13_ = default__.safeModuloInt((self).f19, len(_dafny.SeqWithoutIsStrInference([len(d_1755_v0_)])))
        for d_1756_i0_ in range(p2, hi13_):
            d_1757_v1_: D11
            d_1757_v1_ = D11_DC20()
            d_1757_v1_ = default__.fm60((self).f17, (543) * (728), p3, globalState)
            d_1758_v2_: _dafny.Array
            nw343_ = _dafny.Array(_dafny.Map({}), 3)
            d_1758_v2_ = nw343_
            d_1759_v3_: _dafny.Map
            d_1759_v3_ = _dafny.Map({self.f16: len(d_1755_v0_)})
            index290_ = default__.safeIndex(600, (d_1758_v2_).length(0))
            (d_1758_v2_)[index290_] = d_1759_v3_
            index291_ = default__.safeIndex(600, (d_1758_v2_).length(0))
            (d_1758_v2_)[index291_] = (d_1759_v3_) | (_dafny.Map({p0: d_1756_i0_}))
            d_1760_v4_: _dafny.Array
            nw344_ = _dafny.Array(int(0), 29)
            d_1760_v4_ = nw344_
            index292_ = default__.safeIndex(807, (d_1760_v4_).length(0))
            (d_1760_v4_)[index292_] = d_1756_i0_
            index293_ = default__.safeIndex(807, (d_1760_v4_).length(0))
            (d_1760_v4_)[index293_] = -545
            d_1755_v0_ = (d_1755_v0_) + (d_1755_v0_)
        d_1761_v5_: _dafny.Array
        nw345_ = _dafny.Array(False, 24)
        d_1761_v5_ = nw345_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_1761_v5_).length(0)):
            d_1762_i1_: int = guard_loop_9_
            if (True) and (((0) <= (d_1762_i1_)) and ((d_1762_i1_) < ((d_1761_v5_).length(0)))):
                (d_1761_v5_)[(d_1762_i1_)] = not((self.f16 if not((self).f17) else p0))
        hi14_ = (p2) * (p2)
        for d_1763_i2_ in range((self).f19, hi14_):
            d_1764_v6_: _dafny.Map
            d_1764_v6_ = _dafny.Map({d_1755_v0_: d_1763_i2_})
            d_1765_v7_: _dafny.Map
            d_1765_v7_ = _dafny.Map({p1: p2})
            d_1766_v8_: _dafny.MultiSet
            d_1766_v8_ = _dafny.MultiSet([(self).f17, (self).f17])
            d_1764_v6_ = _dafny.Map({(((d_1755_v0_).set(default__.safeIndex(len(d_1765_v7_), len(d_1755_v0_)), p3)).set(default__.safeIndex(((d_1766_v8_)[p1] if (p1) in (d_1766_v8_) else (self).f19), len((d_1755_v0_).set(default__.safeIndex(len(d_1765_v7_), len(d_1755_v0_)), p3))), p3)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hr"))): default__.safeModuloInt((self).f19, d_1763_i2_)})
            d_1767_v9_: _dafny.Map
            d_1767_v9_ = _dafny.Map({d_1755_v0_: not(not((self).f20))})
            d_1768_v10_: _dafny.Seq
            d_1768_v10_ = _dafny.SeqWithoutIsStrInference([False])
            d_1769_v11_: _dafny.Set
            d_1769_v11_ = _dafny.Set({d_1768_v10_, d_1768_v10_})
            d_1770_v12_: _dafny.Array
            nw346_ = _dafny.Array(None, 7)
            nw346_[int(0)] = _dafny.Map({d_1755_v0_: self.f16})
            nw346_[int(1)] = d_1767_v9_
            nw346_[int(2)] = d_1767_v9_
            nw346_[int(3)] = d_1767_v9_
            nw346_[int(4)] = default__.fm61(d_1769_v11_, (d_1766_v8_).cardinality, globalState)
            nw346_[int(5)] = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqpgv")): p1})
            nw346_[int(6)] = d_1767_v9_
            d_1770_v12_ = nw346_
            index294_ = default__.safeIndex(804, (d_1770_v12_).length(0))
            (d_1770_v12_)[index294_] = d_1767_v9_
            d_1771_v13_: _dafny.Map
            d_1771_v13_ = _dafny.Map({(self).f17: (self).f17})
            d_1772_v14_: _dafny.Map
            d_1772_v14_ = _dafny.Map({len(d_1771_v13_): p2})
            d_1773_v15_: _dafny.Map
            d_1773_v15_ = _dafny.Map({((d_1772_v14_)[d_1763_i2_] if (d_1763_i2_) in (d_1772_v14_) else (self).f19): p0})
            d_1774_v16_: C4
            nw347_ = C4()
            nw347_.ctor__(((d_1773_v15_)[self.f14] if (self.f14) in (d_1773_v15_) else (self).f17), self.f16, (self.f16 if p1 else p1))
            d_1774_v16_ = nw347_
            d_1775_v17_: str
            d_1775_v17_ = _dafny.CodePoint('e')
            d_1776_v19_: _dafny.Seq
            d_1776_v19_ = _dafny.SeqWithoutIsStrInference([d_1755_v0_])
            index295_ = default__.safeIndex(804, (d_1770_v12_).length(0))
            def iife142_():
                coll70_ = _dafny.Map()
                compr_70_: _dafny.Seq
                for compr_70_ in (d_1776_v19_).Elements:
                    d_1777_v18_: _dafny.Seq = compr_70_
                    if (d_1777_v18_) in (d_1776_v19_):
                        coll70_[d_1777_v18_] = p1
                return _dafny.Map(coll70_)
            rhs320_ = ((iife142_()
            ) | (d_1767_v9_)) | (default__.fm61(d_1769_v11_, self.f14, globalState))
            rhs321_ = ((d_1771_v13_) | (d_1771_v13_)) | (d_1771_v13_)
            rhs322_ = d_1774_v16_
            rhs323_ = d_1775_v17_
            lhs248_ = d_1770_v12_
            lhs249_ = default__.safeIndex(804, (d_1770_v12_).length(0))
            lhs248_[lhs249_] = rhs320_
            d_1771_v13_ = rhs321_
            d_1774_v16_ = rhs322_
            d_1775_v17_ = rhs323_
            d_1778_v20_: D3
            d_1778_v20_ = D3_DC3(p1)
            d_1779_v21_: C2
            nw348_ = C2()
            nw348_.ctor__(d_1778_v20_, (self).f17, (self).f17)
            d_1779_v21_ = nw348_
            index296_ = default__.safeIndex(738, (d_1761_v5_).length(0))
            (d_1761_v5_)[index296_] = not (p1) or ((self).f20)
            index297_ = default__.safeIndex(738, (d_1761_v5_).length(0))
            (d_1761_v5_)[index297_] = (self).f20
        index298_ = default__.safeIndex(443, (d_1761_v5_).length(0))
        (d_1761_v5_)[index298_] = ((self).f19) != ((0) - ((self).f19))
        d_1780_v22_: _dafny.Map
        d_1780_v22_ = _dafny.Map({p1: (self).f17})
        d_1781_v23_: D12
        d_1781_v23_ = D12_DC22(d_1755_v0_)
        pat_let_tv26_ = p1
        pat_let_tv27_ = p1
        index299_ = default__.safeIndex(443, (d_1761_v5_).length(0))
        def lambda98_(source30_):
            if source30_.is_DC23:
                d_1782___mcc_h0_ = source30_.cf30
                d_1783___mcc_h1_ = source30_.cf31
                d_1784___mcc_h2_ = source30_.cf32
                d_1785_cf32_ = d_1784___mcc_h2_
                d_1786_cf31_ = d_1783___mcc_h1_
                d_1787_cf30_ = d_1782___mcc_h0_
                return pat_let_tv26_
            elif source30_.is_DC24:
                d_1788___mcc_h3_ = source30_.cf33
                d_1789___mcc_h4_ = source30_.cf34
                d_1790___mcc_h5_ = source30_.cf35
                d_1791___mcc_h6_ = source30_.cf36
                d_1792___mcc_h7_ = source30_.cf37
                d_1793_cf37_ = d_1792___mcc_h7_
                d_1794_cf36_ = d_1791___mcc_h6_
                d_1795_cf35_ = d_1790___mcc_h5_
                d_1796_cf34_ = d_1789___mcc_h4_
                d_1797_cf33_ = d_1788___mcc_h3_
                return d_1797_cf33_
            elif True:
                d_1798___mcc_h8_ = source30_.cf29
                d_1799_cf29_ = d_1798___mcc_h8_
                return pat_let_tv27_

        rhs324_ = self.f14
        rhs325_ = False
        rhs326_ = not(lambda98_(d_1781_v23_))
        rhs327_ = (self).f20
        rhs328_ = (d_1780_v22_) | (d_1780_v22_)
        lhs250_ = globalState
        lhs251_ = self
        lhs252_ = d_1761_v5_
        lhs253_ = default__.safeIndex(443, (d_1761_v5_).length(0))
        lhs250_.f7 = rhs324_
        lhs251_.f16 = rhs325_
        r1 = rhs326_
        lhs252_[lhs253_] = rhs327_
        d_1780_v22_ = rhs328_
        d_1800_v24_: _dafny.Array
        def lambda99_(d_1801_p1_, d_1802_p0_):
            def lambda100_(d_1803_i3_):
                return _dafny.MultiSet([len(_dafny.Set({not(not(d_1801_p1_)), d_1802_p0_, d_1801_p1_}))])

            return lambda100_

        init58_ = lambda99_(p1, p0)
        nw349_ = _dafny.Array(None, 27)
        for i0_58_ in range(nw349_.length(0)):
            nw349_[i0_58_] = init58_(i0_58_)
        d_1800_v24_ = nw349_
        d_1804_v25_: _dafny.MultiSet
        d_1804_v25_ = _dafny.MultiSet([(self).f19])
        d_1805_v26_: _dafny.Set
        d_1805_v26_ = _dafny.Set({False})
        d_1806_v27_: _dafny.Seq
        d_1806_v27_ = _dafny.SeqWithoutIsStrInference([self.f14, (self).f19])
        d_1807_v28_: _dafny.Seq
        d_1807_v28_ = _dafny.SeqWithoutIsStrInference([len(d_1806_v27_), (self).f19])
        index300_ = default__.safeIndex(195, (d_1800_v24_).length(0))
        (d_1800_v24_)[index300_] = (d_1804_v25_).set(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1805_v26_)]), default__.fm38(p1, p2, 816, globalState), d_1806_v27_, d_1807_v28_])), default__.abs((self).f19))
        d_1808_v29_: _dafny.MultiSet
        d_1808_v29_ = _dafny.MultiSet([self.f16])
        d_1809_v30_: _dafny.Map
        d_1809_v30_ = _dafny.Map({(d_1808_v29_).issubset(default__.fm1(globalState)): (d_1804_v25_).set(322, default__.abs(len(d_1755_v0_)))})
        index301_ = default__.safeIndex(195, (d_1800_v24_).length(0))
        (d_1800_v24_)[index301_] = ((d_1809_v30_)[(d_1761_v5_)[default__.safeIndex(443, (d_1761_v5_).length(0))]] if ((d_1761_v5_)[default__.safeIndex(443, (d_1761_v5_).length(0))]) in (d_1809_v30_) else _dafny.MultiSet(d_1806_v27_))
        def iife143_():
            coll71_ = _dafny.Map()
            compr_71_: int
            for compr_71_ in _dafny.IntegerRange(-924, 707):
                d_1810_v31_: int = compr_71_
                if ((-924) <= (d_1810_v31_)) and ((d_1810_v31_) < (707)):
                    coll71_[(d_1810_v31_) + (p2)] = d_1780_v22_
            return _dafny.Map(coll71_)
        r0 = ((self).fm18((self).f20, default__.fm0((d_1761_v5_)[default__.safeIndex(443, (d_1761_v5_).length(0))], globalState), iife143_()
        , (0) - ((self).f19), globalState)) + (704)
        d_1811_v32_: _dafny.Set
        d_1811_v32_ = _dafny.Set({p2, p2, (0) - (self.f14), self.f14})
        r1 = not((d_1811_v32_).ispropersubset(d_1811_v32_))
        r2 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dljlsf"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dljlsf")))), p3)
        r3 = (self).f17
        return r0, r1, r2, r3

    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20

class C7(T2, T0, T1):
    def  __init__(self):
        self._f16: bool = False
        self._f14: int = int(0)
        self._f17: bool = False
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
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f16, f17, f14):
        (self).f16 = f16
        (self)._f17 = f17
        (self).f14 = f14

    def fm12(self, p0, globalState):
        return (0) - ((self.f14) + (len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({self.f14: (self).f17}))]), _dafny.SeqWithoutIsStrInference([(D3_DC4(True, self.f14, self.f14, (self).f17, False)).cf5]), _dafny.SeqWithoutIsStrInference([self.f14 for d_1812_i0_ in range(default__.abs(499))]), (_dafny.SeqWithoutIsStrInference([self.f14]))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([91, self.f14, self.f14, self.f14, self.f14]) for d_1813_i1_ in range(default__.abs(504))])))))

    def fm13(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([(self).f17]))) + ((_dafny.SeqWithoutIsStrInference([False, (self).f17])) + (_dafny.SeqWithoutIsStrInference([not(self.f16)])))

    def fm3(self, globalState):
        if self.f16:
            return _dafny.Map({self.f14: self.f14})
        elif True:
            return _dafny.Map({self.f14: self.f14})

    def fm16(self, p0, globalState):
        return self.f14

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: bool = False
        d_1814_v0_: _dafny.Seq
        d_1814_v0_ = _dafny.SeqWithoutIsStrInference([p1, (self).f17])
        d_1815_v1_: _dafny.Set
        d_1815_v1_ = _dafny.Set({p2, p2})
        d_1816_v2_: _dafny.Map
        d_1816_v2_ = _dafny.Map({p2: p1})
        d_1817_v3_: _dafny.Seq
        d_1817_v3_ = _dafny.SeqWithoutIsStrInference([602, p2, p2])
        d_1818_v4_: _dafny.Set
        d_1818_v4_ = _dafny.Set({p1, p0, self.f16, self.f16})
        d_1819_v5_: _dafny.Array
        nw350_ = _dafny.Array(None, 15)
        nw350_[int(0)] = (p2) <= (p2)
        nw350_[int(1)] = (self).f17
        nw350_[int(2)] = p0
        nw350_[int(3)] = not(not(not(p0)))
        nw350_[int(4)] = not((d_1814_v0_) != (d_1814_v0_))
        nw350_[int(5)] = self.f16
        nw350_[int(6)] = (self).f17
        nw350_[int(7)] = p1
        nw350_[int(8)] = (d_1815_v1_).isdisjoint(d_1815_v1_)
        nw350_[int(9)] = not(not(((self).f17 if (self).f17 else ((d_1816_v2_)[(_dafny.MultiSet([self.f16])).cardinality] if ((_dafny.MultiSet([self.f16])).cardinality) in (d_1816_v2_) else self.f16))))
        nw350_[int(10)] = default__.fm17(p1, p2, p2, len(d_1817_v3_), globalState)
        nw350_[int(11)] = (d_1818_v4_).isdisjoint(d_1818_v4_)
        nw350_[int(12)] = (self).f17
        nw350_[int(13)] = self.f16
        nw350_[int(14)] = default__.fm17((self).f17, self.f14, self.f14, (_dafny.MultiSet(d_1814_v0_)).cardinality, globalState)
        d_1819_v5_ = nw350_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_1819_v5_).length(0)):
            d_1820_i0_: int = guard_loop_10_
            if (True) and (((0) <= (d_1820_i0_)) and ((d_1820_i0_) < ((d_1819_v5_).length(0)))):
                (d_1819_v5_)[(d_1820_i0_)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdxkk"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1821_i1_ in range(default__.abs(520))]))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cj")))
        d_1822_i2_: int
        d_1822_i2_ = 0
        with _dafny.label("9"):
            while not(default__.fm17((self).f17, (p2) - (self.f14), default__.safeModuloInt(p2, (0) - (p2)), self.f14, globalState)):
                with _dafny.c_label("9"):
                    if (d_1822_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_1822_i2_ = (d_1822_i2_) + (1)
                    d_1823_v6_: _dafny.Map
                    d_1823_v6_ = _dafny.Map({p1: p1})
                    d_1823_v6_ = (d_1823_v6_).set((self).f17, p1)
                    d_1824_v7_: _dafny.Array
                    def lambda101_(d_1825_p3_):
                        def lambda102_(d_1826_i3_):
                            return _dafny.SeqWithoutIsStrInference([d_1825_p3_ for d_1827_i4_ in range(default__.abs(501))])

                        return lambda102_

                    init59_ = lambda101_(p3)
                    nw351_ = _dafny.Array(None, 27)
                    for i0_59_ in range(nw351_.length(0)):
                        nw351_[i0_59_] = init59_(i0_59_)
                    d_1824_v7_ = nw351_
                    d_1824_v7_ = d_1824_v7_
                    r3 = p0
                    d_1828_v8_: C4
                    nw352_ = C4()
                    nw352_.ctor__(p1, self.f16, p0)
                    d_1828_v8_ = nw352_
                    pass
            pass
        d_1829_v9_: _dafny.Seq
        d_1829_v9_ = _dafny.SeqWithoutIsStrInference([d_1818_v4_, (d_1818_v4_).intersection(d_1818_v4_)])
        d_1829_v9_ = _dafny.SeqWithoutIsStrInference([(d_1818_v4_) - (d_1818_v4_)])
        d_1830_v10_: C4
        nw353_ = C4()
        nw353_.ctor__((p3) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))), self.f16, self.f16)
        d_1830_v10_ = nw353_
        d_1831_v11_: C5
        nw354_ = C5()
        nw354_.ctor__(default__.fm43(p1, globalState), p1)
        d_1831_v11_ = nw354_
        if self.f16:
            if ((d_1830_v10_).f21 if default__.fm21(globalState) else not ((d_1830_v10_).f21) or ((d_1830_v10_).f21)):
                d_1832_v12_: _dafny.Map
                d_1832_v12_ = _dafny.Map({p2: p2})
                d_1832_v12_ = (d_1832_v12_) | ((d_1832_v12_).set(len(d_1832_v12_), p2))
                d_1833_v13_: _dafny.Array
                nw355_ = _dafny.Array(int(0), 12)
                d_1833_v13_ = nw355_
                nw356_ = _dafny.Array(int(0), 13)
                d_1833_v13_ = nw356_
                d_1834_v14_: _dafny.Set
                d_1834_v14_ = _dafny.Set({d_1819_v5_})
                r1 = (d_1834_v14_).ispropersubset(d_1834_v14_)
                d_1817_v3_ = ((d_1817_v3_) + ((_dafny.SeqWithoutIsStrInference([p2])) + (d_1817_v3_))).set(default__.safeIndex(745, len((d_1817_v3_) + ((_dafny.SeqWithoutIsStrInference([p2])) + (d_1817_v3_)))), p2)
                (globalState).f2 = default__.safeModuloInt(p2, self.f14)
            elif True:
                d_1835_v15_: str
                d_1835_v15_ = _dafny.CodePoint('h')
                d_1836_v16_: _dafny.Seq
                d_1836_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiltx"))
                d_1837_v17_: _dafny.Seq
                d_1837_v17_ = _dafny.SeqWithoutIsStrInference([d_1836_v16_])
                rhs329_ = p1
                rhs330_ = p3
                rhs331_ = (d_1830_v10_).f21
                rhs332_ = d_1837_v17_
                rhs333_ = len(d_1814_v0_)
                lhs254_ = globalState
                r3 = rhs329_
                d_1835_v15_ = rhs330_
                r3 = rhs331_
                d_1837_v17_ = rhs332_
                lhs254_.f7 = rhs333_
                r3 = (d_1830_v10_).f21
                r3 = not(p1)
                index302_ = default__.safeIndex(660, (d_1819_v5_).length(0))
                (d_1819_v5_)[index302_] = (d_1830_v10_).f21
                d_1838_v18_: _dafny.Set
                d_1838_v18_ = _dafny.Set({d_1819_v5_})
                index303_ = default__.safeIndex(660, (d_1819_v5_).length(0))
                (d_1819_v5_)[index303_] = not((True if (self).f17 else (d_1817_v3_) == ((_dafny.SeqWithoutIsStrInference([p2, len(d_1838_v18_)])).set(default__.safeIndex(self.f14, len(_dafny.SeqWithoutIsStrInference([p2, len(d_1838_v18_)]))), len(d_1816_v2_)))))
                (globalState).f7 = (p2) + ((p2) + (p2))
            d_1839_v19_: _dafny.Seq
            d_1839_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdpuq"))
            d_1840_v20_: _dafny.Map
            d_1840_v20_ = _dafny.Map({self.f14: self.f14})
            d_1841_v21_: _dafny.Array
            nw357_ = _dafny.Array(None, 6)
            nw357_[int(0)] = d_1839_v19_
            nw357_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_1839_v19_)[default__.safeIndex((d_1817_v3_)[default__.safeIndex(p2, len(d_1817_v3_))], len(d_1839_v19_))] for d_1842_i5_ in range(default__.abs(468))])
            nw357_[int(2)] = default__.fm42(p1, len(d_1817_v3_), 342, (d_1830_v10_).f21, globalState)
            nw357_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydpgyfg"))) + (d_1839_v19_)
            nw357_[int(4)] = d_1839_v19_
            nw357_[int(5)] = default__.fm42(p0, len(d_1817_v3_), ((d_1840_v20_)[p2] if (p2) in (d_1840_v20_) else self.f14), (d_1830_v10_).f21, globalState)
            d_1841_v21_ = nw357_
            index304_ = default__.safeIndex(879, (d_1841_v21_).length(0))
            (d_1841_v21_)[index304_] = d_1839_v19_
            index305_ = default__.safeIndex(879, (d_1841_v21_).length(0))
            rhs334_ = (d_1839_v19_) + (d_1839_v19_)
            rhs335_ = (d_1830_v10_).f21
            rhs336_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgew"))
            lhs255_ = d_1841_v21_
            lhs256_ = default__.safeIndex(879, (d_1841_v21_).length(0))
            lhs257_ = self
            lhs255_[lhs256_] = rhs334_
            lhs257_.f16 = rhs335_
            r2 = rhs336_
            d_1843_v22_: _dafny.Array
            d_1843_v22_ = d_1819_v5_
            d_1844_v23_: _dafny.Map
            d_1844_v23_ = _dafny.Map({(d_1843_v22_): p2})
            d_1844_v23_ = (d_1844_v23_).set(d_1819_v5_, (d_1817_v3_)[default__.safeIndex(self.f14, len(d_1817_v3_))])
            d_1845_v24_: _dafny.Map
            d_1845_v24_ = _dafny.Map({(d_1830_v10_).f21: ((d_1830_v10_).fm12(self.f14, globalState)) > (self.f14)})
            d_1846_v25_: _dafny.Array
            nw358_ = _dafny.Array(None, 9)
            nw358_[int(0)] = p2
            nw358_[int(1)] = p2
            nw358_[int(2)] = len(d_1817_v3_)
            nw358_[int(3)] = 27
            nw358_[int(4)] = p2
            nw358_[int(5)] = p2
            nw358_[int(6)] = (0) - (p2)
            nw358_[int(7)] = p2
            nw358_[int(8)] = self.f14
            d_1846_v25_ = nw358_
            d_1847_v26_: _dafny.Set
            d_1847_v26_ = _dafny.Set({d_1846_v25_})
            index306_ = default__.safeIndex(652, (d_1819_v5_).length(0))
            (d_1819_v5_)[index306_] = (d_1847_v26_).issubset(d_1847_v26_)
            d_1848_v27_: _dafny.Seq
            d_1848_v27_ = _dafny.SeqWithoutIsStrInference([d_1845_v24_, d_1845_v24_, d_1845_v24_, default__.fm40(self.f16, globalState), (d_1845_v24_).set(p0, p0)])
            d_1849_v28_: D4
            d_1849_v28_ = D4_DC6((d_1848_v27_)[default__.safeIndex(self.f14, len(d_1848_v27_))])
            index307_ = default__.safeIndex(652, (d_1819_v5_).length(0))
            rhs337_ = (self).f17
            rhs338_ = self.f14
            rhs339_ = (d_1849_v28_).cf10
            rhs340_ = default__.fm21(globalState)
            rhs341_ = len(((d_1841_v21_)[default__.safeIndex(879, (d_1841_v21_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbl"))))
            lhs258_ = globalState
            lhs259_ = d_1819_v5_
            lhs260_ = default__.safeIndex(652, (d_1819_v5_).length(0))
            lhs261_ = globalState
            r1 = rhs337_
            lhs258_.f7 = rhs338_
            d_1845_v24_ = rhs339_
            lhs259_[lhs260_] = rhs340_
            lhs261_.f7 = rhs341_
            index308_ = default__.safeIndex(652, (d_1819_v5_).length(0))
            (d_1819_v5_)[index308_] = not((self.f14) >= ((d_1831_v11_).fm12((0) - (p2), globalState)))
        elif True:
            d_1850_v29_: _dafny.Seq
            d_1850_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfbavpph"))
            d_1851_v30_: _dafny.Map
            d_1851_v30_ = _dafny.Map({False: p0})
            d_1852_v31_: C3
            nw359_ = C3()
            nw359_.ctor__(self.f14, (d_1850_v29_).set(default__.safeIndex(p2, len(d_1850_v29_)), _dafny.CodePoint('c')), p1, (len(d_1851_v30_)) == (self.f14))
            d_1852_v31_ = nw359_
            index309_ = default__.safeIndex(227, (d_1819_v5_).length(0))
            (d_1819_v5_)[index309_] = (_dafny.SeqWithoutIsStrInference([p3, p3, p3, p3])) <= ((d_1852_v31_).f23)
            d_1853_v32_: _dafny.MultiSet
            d_1853_v32_ = _dafny.MultiSet([p2])
            index310_ = default__.safeIndex(227, (d_1819_v5_).length(0))
            (d_1819_v5_)[index310_] = not((d_1853_v32_).isdisjoint((d_1853_v32_ if p1 else d_1853_v32_)))
            d_1854_v33_: _dafny.Array
            def lambda103_(d_1855_v4_):
                def lambda104_(d_1856_i6_):
                    return d_1855_v4_

                return lambda104_

            init60_ = lambda103_(d_1818_v4_)
            nw360_ = _dafny.Array(None, 19)
            for i0_60_ in range(nw360_.length(0)):
                nw360_[i0_60_] = init60_(i0_60_)
            d_1854_v33_ = nw360_
            d_1857_v34_: C0
            nw361_ = C0()
            nw361_.ctor__(d_1854_v33_, (d_1852_v31_).f22)
            d_1857_v34_ = nw361_
            d_1858_v35_: _dafny.Seq
            d_1858_v35_ = _dafny.SeqWithoutIsStrInference([(d_1816_v2_) | (d_1816_v2_), d_1816_v2_, d_1816_v2_, d_1816_v2_, d_1816_v2_])
            rhs342_ = (self).f17
            rhs343_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: (d_1830_v10_).f21}) for d_1859_i7_ in range(default__.abs(284))])
            r1 = rhs342_
            d_1858_v35_ = rhs343_
            (d_1857_v34_).f27 = self.f14
        r0 = (0) - (p2)
        d_1860_v36_: _dafny.Seq
        d_1860_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nei"))
        d_1861_v37_: _dafny.Seq
        d_1861_v37_ = _dafny.SeqWithoutIsStrInference([d_1860_v36_])
        r1 = ((self).f17) == ((_dafny.SeqWithoutIsStrInference([d_1860_v36_, (d_1860_v36_).set(default__.safeIndex(self.f14, len(d_1860_v36_)), p3), d_1860_v36_])) == (d_1861_v37_))
        r2 = (d_1860_v36_) + (d_1860_v36_)
        d_1862_v38_: _dafny.Map
        d_1862_v38_ = _dafny.Map({(_dafny.MultiSet([p0, (self).f17, True])).set(p1, default__.abs(self.f14)): not(not((self).f17))})
        d_1863_v39_: _dafny.MultiSet
        d_1863_v39_ = _dafny.MultiSet([self.f16, (d_1830_v10_).f21])
        r3 = (((d_1862_v38_)[d_1863_v39_] if (d_1863_v39_) in (d_1862_v38_) else not(default__.fm31((self).f17, p1, globalState))) if (self).f17 else (d_1814_v0_) == (d_1814_v0_))
        return r0, r1, r2, r3

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        if p1:
            d_1864_v0_: _dafny.Array
            nw362_ = _dafny.Array(False, 3)
            d_1864_v0_ = nw362_
            d_1865_v1_: _dafny.Array
            d_1865_v1_ = d_1864_v0_
            d_1866_v2_: str
            d_1866_v2_ = _dafny.CodePoint('b')
            d_1867_v3_: _dafny.Map
            d_1867_v3_ = _dafny.Map({d_1866_v2_: d_1864_v0_})
            d_1868_v4_: _dafny.Array
            nw363_ = _dafny.Array(None, 14)
            nw363_[int(0)] = d_1864_v0_
            nw363_[int(1)] = d_1864_v0_
            nw363_[int(2)] = d_1865_v1_
            nw363_[int(3)] = d_1865_v1_
            nw363_[int(4)] = d_1865_v1_
            nw363_[int(5)] = d_1865_v1_
            nw363_[int(6)] = d_1864_v0_
            nw363_[int(7)] = d_1865_v1_
            nw363_[int(8)] = d_1865_v1_
            nw363_[int(9)] = d_1865_v1_
            nw363_[int(10)] = d_1865_v1_
            nw363_[int(11)] = d_1865_v1_
            nw363_[int(12)] = ((d_1867_v3_)[d_1866_v2_] if (d_1866_v2_) in (d_1867_v3_) else d_1864_v0_)
            nw363_[int(13)] = d_1865_v1_
            d_1868_v4_ = nw363_
            d_1869_v5_: _dafny.Array
            nw364_ = _dafny.Array(None, 11)
            nw364_[int(0)] = d_1868_v4_
            nw364_[int(1)] = d_1868_v4_
            nw364_[int(2)] = (d_1868_v4_ if default__.fm17(self.f16, len(_dafny.Set({p0, p0, self.f14})), p0, self.f14, globalState) else d_1868_v4_)
            nw364_[int(3)] = d_1868_v4_
            nw364_[int(4)] = d_1868_v4_
            nw364_[int(5)] = d_1868_v4_
            nw364_[int(6)] = d_1868_v4_
            nw364_[int(7)] = (d_1868_v4_ if True else d_1868_v4_)
            nw364_[int(8)] = d_1868_v4_
            nw364_[int(9)] = d_1868_v4_
            nw364_[int(10)] = d_1868_v4_
            d_1869_v5_ = nw364_
            index311_ = default__.safeIndex(539, (d_1869_v5_).length(0))
            (d_1869_v5_)[index311_] = d_1868_v4_
            d_1870_v6_: _dafny.Array
            nw365_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
            d_1870_v6_ = nw365_
            d_1871_v7_: _dafny.Seq
            d_1871_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yefkcasyq"))
            index312_ = default__.safeIndex(415, (d_1870_v6_).length(0))
            (d_1870_v6_)[index312_] = d_1871_v7_
            index313_ = default__.safeIndex(539, (d_1869_v5_).length(0))
            index314_ = default__.safeIndex(415, (d_1870_v6_).length(0))
            rhs344_ = d_1868_v4_
            rhs345_ = d_1871_v7_
            rhs346_ = self.f16
            lhs262_ = d_1869_v5_
            lhs263_ = default__.safeIndex(539, (d_1869_v5_).length(0))
            lhs264_ = d_1870_v6_
            lhs265_ = default__.safeIndex(415, (d_1870_v6_).length(0))
            lhs266_ = self
            lhs262_[lhs263_] = rhs344_
            lhs264_[lhs265_] = rhs345_
            lhs266_.f16 = rhs346_
            d_1872_v8_: _dafny.Map
            d_1872_v8_ = _dafny.Map({d_1864_v0_: (d_1864_v0_ if self.f16 else d_1864_v0_)})
            d_1872_v8_ = (d_1872_v8_).set(d_1864_v0_, d_1864_v0_)
            d_1873_v9_: C6
            nw366_ = C6()
            nw366_.ctor__((0) - (self.f14), self.f16, p1, (self.f14) > (p0), p0)
            d_1873_v9_ = nw366_
            d_1873_v9_ = d_1873_v9_
            (globalState).f7 = p0
            d_1874_v10_: _dafny.Array
            nw367_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_1874_v10_ = nw367_
            d_1875_v11_: T1
            nw368_ = C6()
            nw368_.ctor__((d_1873_v9_).f19, not(self.f16), self.f16, not((d_1873_v9_).f20), self.f14)
            d_1875_v11_ = nw368_
            d_1876_v12_: _dafny.Map
            d_1876_v12_ = _dafny.Map({d_1875_v11_: p0})
            d_1877_v13_: _dafny.Map
            d_1877_v13_ = _dafny.Map({not(p1): self.f14})
            d_1878_v14_: _dafny.Set
            d_1878_v14_ = _dafny.Set({(d_1873_v9_).f20, (self).f17})
            d_1879_v15_: _dafny.Map
            d_1879_v15_ = _dafny.Map({self.f14: d_1875_v11_.f16})
            d_1880_v16_: _dafny.Map
            d_1880_v16_ = _dafny.Map({(0) - ((self).fm16(d_1879_v15_, globalState)): d_1866_v2_})
            d_1881_v17_: _dafny.Array
            nw369_ = _dafny.Array(None, 24)
            nw369_[int(0)] = p0
            nw369_[int(1)] = self.f14
            nw369_[int(2)] = (d_1873_v9_).f19
            nw369_[int(3)] = self.f14
            nw369_[int(4)] = self.f14
            nw369_[int(5)] = self.f14
            nw369_[int(6)] = self.f14
            nw369_[int(7)] = -166
            nw369_[int(8)] = len(d_1876_v12_)
            nw369_[int(9)] = p0
            nw369_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxtakl")))
            nw369_[int(11)] = (d_1873_v9_).f19
            nw369_[int(12)] = len(d_1877_v13_)
            nw369_[int(13)] = (d_1873_v9_).f19
            nw369_[int(14)] = len(d_1878_v14_)
            nw369_[int(15)] = (0) - (self.f14)
            nw369_[int(16)] = self.f14
            nw369_[int(17)] = p0
            nw369_[int(18)] = (0) - (self.f14)
            nw369_[int(19)] = self.f14
            nw369_[int(20)] = self.f14
            nw369_[int(21)] = len(d_1880_v16_)
            nw369_[int(22)] = p0
            nw369_[int(23)] = len((d_1870_v6_)[default__.safeIndex(415, (d_1870_v6_).length(0))])
            d_1881_v17_ = nw369_
            index315_ = default__.safeIndex(138, (d_1874_v10_).length(0))
            (d_1874_v10_)[index315_] = d_1881_v17_
            d_1882_v18_: _dafny.Seq
            d_1882_v18_ = _dafny.SeqWithoutIsStrInference([(d_1873_v9_).f19])
            index316_ = default__.safeIndex(138, (d_1874_v10_).length(0))
            def iife144_():
                coll72_ = _dafny.Set()
                compr_72_: int
                for compr_72_ in (d_1882_v18_).Elements:
                    d_1883_v19_: int = compr_72_
                    if (d_1883_v19_) in (d_1882_v18_):
                        coll72_ = coll72_.union(_dafny.Set([default__.safeModuloInt(d_1883_v19_, 82)]))
                return _dafny.Set(coll72_)
            rhs347_ = (D9_DC15(True, (d_1873_v9_).f20, iife144_()
, d_1881_v17_, d_1864_v0_)).cf21
            rhs348_ = (d_1873_v9_).f20
            lhs267_ = d_1874_v10_
            lhs268_ = default__.safeIndex(138, (d_1874_v10_).length(0))
            lhs269_ = d_1875_v11_
            lhs267_[lhs268_] = rhs347_
            lhs269_.f16 = rhs348_
        elif True:
            d_1884_v20_: _dafny.Seq
            d_1884_v20_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (p0)), p0])
            (globalState).f2 = len(d_1884_v20_)
            d_1885_v21_: _dafny.MultiSet
            d_1885_v21_ = _dafny.MultiSet([(self).f17])
            d_1886_v22_: _dafny.Seq
            d_1886_v22_ = _dafny.SeqWithoutIsStrInference([False, (self).f17])
            d_1887_v23_: D16
            d_1887_v23_ = D16_DC32(d_1885_v21_)
            d_1888_v24_: _dafny.Map
            d_1888_v24_ = _dafny.Map({p1: p1})
            d_1889_v25_: _dafny.Array
            nw370_ = _dafny.Array(None, 23)
            nw370_[int(0)] = d_1885_v21_
            nw370_[int(1)] = (d_1885_v21_).intersection(_dafny.MultiSet(d_1886_v22_))
            nw370_[int(2)] = (default__.fm1(globalState)) | (d_1885_v21_)
            nw370_[int(3)] = (_dafny.MultiSet(d_1886_v22_)) - (d_1885_v21_)
            nw370_[int(4)] = d_1885_v21_
            nw370_[int(5)] = _dafny.MultiSet([self.f16, True])
            nw370_[int(6)] = d_1885_v21_
            nw370_[int(7)] = d_1885_v21_
            nw370_[int(8)] = d_1885_v21_
            nw370_[int(9)] = default__.fm1(globalState)
            nw370_[int(10)] = d_1885_v21_
            nw370_[int(11)] = (d_1885_v21_) - (_dafny.MultiSet([False]))
            nw370_[int(12)] = d_1885_v21_
            nw370_[int(13)] = d_1885_v21_
            nw370_[int(14)] = d_1885_v21_
            nw370_[int(15)] = d_1885_v21_
            nw370_[int(16)] = d_1885_v21_
            nw370_[int(17)] = d_1885_v21_
            nw370_[int(18)] = (d_1887_v23_).cf46
            nw370_[int(19)] = (d_1885_v21_) - (d_1885_v21_)
            nw370_[int(20)] = (d_1885_v21_).set(((d_1888_v24_)[self.f16] if (self.f16) in (d_1888_v24_) else self.f16), default__.abs(self.f14))
            nw370_[int(21)] = (_dafny.MultiSet([default__.fm21(globalState), p1, self.f16])).set(self.f16, default__.abs((0) - ((d_1884_v20_)[default__.safeIndex(self.f14, len(d_1884_v20_))])))
            nw370_[int(22)] = (d_1887_v23_).cf46
            d_1889_v25_ = nw370_
            index317_ = default__.safeIndex(840, (d_1889_v25_).length(0))
            (d_1889_v25_)[index317_] = (d_1885_v21_) | (d_1885_v21_)
            index318_ = default__.safeIndex(840, (d_1889_v25_).length(0))
            (d_1889_v25_)[index318_] = _dafny.MultiSet([not (False) or (self.f16)])
            d_1890_v26_: _dafny.Seq
            d_1890_v26_ = d_1884_v20_
            d_1890_v26_ = d_1890_v26_
            d_1891_v27_: _dafny.Seq
            d_1891_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhiffetkh"))
            (self).f16 = (d_1891_v27_) != (d_1891_v27_)
            (self).f16 = self.f16
        d_1892_v28_: _dafny.Array
        nw371_ = _dafny.Array(_dafny.CodePoint('D'), 10)
        d_1892_v28_ = nw371_
        d_1893_v29_: _dafny.Seq
        d_1893_v29_ = _dafny.SeqWithoutIsStrInference([d_1892_v28_])
        d_1893_v29_ = d_1893_v29_
        d_1894_v30_: _dafny.Set
        d_1894_v30_ = _dafny.Set({self.f16, (self).f17})
        hi15_ = len(d_1894_v30_)
        for d_1895_i0_ in range(self.f14, hi15_):
            d_1896_v31_: _dafny.Seq
            d_1896_v31_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1897_v32_: _dafny.Map
            d_1897_v32_ = _dafny.Map({p0: (0) - (d_1895_i0_)})
            d_1898_v33_: str
            d_1898_v33_ = _dafny.CodePoint('h')
            d_1899_v34_: _dafny.Array
            nw372_ = _dafny.Array(None, 10)
            nw372_[int(0)] = p1
            nw372_[int(1)] = (self.f16) in (d_1896_v31_)
            nw372_[int(2)] = default__.fm43(p1, globalState)
            nw372_[int(3)] = (self.f16 if (self).f17 else p1)
            nw372_[int(4)] = self.f16
            nw372_[int(5)] = self.f16
            nw372_[int(6)] = not((self).f17)
            nw372_[int(7)] = p1
            nw372_[int(8)] = (d_1895_i0_) != (len(d_1897_v32_))
            nw372_[int(9)] = (d_1898_v33_) not in (_dafny.SeqWithoutIsStrInference([d_1898_v33_ for d_1900_i1_ in range(default__.abs(-425))]))
            d_1899_v34_ = nw372_
            d_1901_v35_: _dafny.Seq
            d_1901_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlmtgv"))
            d_1902_v36_: _dafny.MultiSet
            d_1902_v36_ = _dafny.MultiSet([d_1901_v35_, d_1901_v35_])
            index319_ = default__.safeIndex(756, (d_1899_v34_).length(0))
            (d_1899_v34_)[index319_] = (d_1902_v36_).issubset(d_1902_v36_)
            index320_ = default__.safeIndex(756, (d_1899_v34_).length(0))
            (d_1899_v34_)[index320_] = ((self).f17) == (self.f16)
            d_1898_v33_ = d_1898_v33_
            (self).f16 = p1
            d_1903_v37_: _dafny.Array
            nw373_ = _dafny.Array(_dafny.MultiSet({}), 12)
            d_1903_v37_ = nw373_
            d_1904_v39_: _dafny.Seq
            d_1904_v39_ = _dafny.SeqWithoutIsStrInference([423, p0])
            d_1905_v40_: _dafny.MultiSet
            def iife145_():
                coll73_ = _dafny.Map()
                compr_73_: int
                for compr_73_ in (d_1904_v39_).Elements:
                    d_1906_v38_: int = compr_73_
                    if (d_1906_v38_) in (d_1904_v39_):
                        coll73_[default__.safeDivisionInt(d_1906_v38_, len(d_1897_v32_))] = D11_DC20()
                return _dafny.Map(coll73_)
            d_1905_v40_ = _dafny.MultiSet([iife145_()
            ])
            index321_ = default__.safeIndex(892, (d_1903_v37_).length(0))
            (d_1903_v37_)[index321_] = (d_1905_v40_) - (d_1905_v40_)
            index322_ = default__.safeIndex(892, (d_1903_v37_).length(0))
            (d_1903_v37_)[index322_] = ((d_1905_v40_) - (d_1905_v40_)) - (d_1905_v40_)
        d_1907_v41_: _dafny.Seq
        d_1907_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "intkfrwlm"))
        d_1907_v41_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmngcvt"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwta"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1908_i2_ in range(default__.abs(277))])))
        hi16_ = self.f14
        for d_1909_i3_ in range(self.f14, hi16_):
            d_1910_v42_: _dafny.Seq
            d_1910_v42_ = _dafny.SeqWithoutIsStrInference([392])
            d_1911_v43_: _dafny.MultiSet
            d_1911_v43_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0, (0) - (self.f14)]), d_1910_v42_])
            d_1912_v44_: _dafny.Map
            d_1912_v44_ = _dafny.Map({d_1911_v43_: (d_1910_v42_)[default__.safeIndex(p0, len(d_1910_v42_))]})
            d_1912_v44_ = (d_1912_v44_).set(d_1911_v43_, p0)
            (self).f16 = ((0) - (d_1909_i3_)) > (default__.fm0((self).f17, globalState))
            (self).f16 = True
            (self).f14 = default__.safeDivisionInt(677, d_1909_i3_)
        d_1907_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1913_i4_ in range(default__.abs(262))])
        r0 = self.f14
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_1914_v0_: _dafny.Array
        nw374_ = _dafny.Array(_dafny.Array(None, 0), 24)
        d_1914_v0_ = nw374_
        d_1915_v1_: _dafny.Map
        d_1915_v1_ = _dafny.Map({p1: p1})
        d_1916_v2_: _dafny.Seq
        d_1916_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnmueq"))
        d_1917_v3_: _dafny.Array
        nw375_ = _dafny.Array(None, 17)
        nw375_[int(0)] = not(p2)
        nw375_[int(1)] = True
        nw375_[int(2)] = (self).f17
        nw375_[int(3)] = default__.fm27(self.f14, p1, self.f14, (d_1915_v1_).set(131, len(d_1916_v2_)), globalState)
        nw375_[int(4)] = default__.fm21(globalState)
        nw375_[int(5)] = True
        nw375_[int(6)] = p2
        nw375_[int(7)] = self.f16
        nw375_[int(8)] = p2
        nw375_[int(9)] = (self).f17
        nw375_[int(10)] = (self).f17
        nw375_[int(11)] = self.f16
        nw375_[int(12)] = not((self).f17)
        nw375_[int(13)] = self.f16
        nw375_[int(14)] = self.f16
        nw375_[int(15)] = False
        nw375_[int(16)] = p2
        d_1917_v3_ = nw375_
        d_1918_v4_: _dafny.Map
        d_1918_v4_ = _dafny.Map({self.f16: d_1917_v3_})
        index323_ = default__.safeIndex(460, (d_1914_v0_).length(0))
        (d_1914_v0_)[index323_] = ((d_1918_v4_)[(self).f17] if ((self).f17) in (d_1918_v4_) else d_1917_v3_)
        index324_ = default__.safeIndex(460, (d_1914_v0_).length(0))
        (d_1914_v0_)[index324_] = d_1917_v3_
        _ingredientsd_1 = []
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, ((d_1914_v0_)[default__.safeIndex(460, (d_1914_v0_).length(0))]).length(0)):
            d_1919_i0_: int = guard_loop_11_
            if (True) and (((0) <= (d_1919_i0_)) and ((d_1919_i0_) < (((d_1914_v0_)[default__.safeIndex(460, (d_1914_v0_).length(0))]).length(0)))):
                _ingredientsd_1.append(((d_1914_v0_)[default__.safeIndex(460, (d_1914_v0_).length(0))], int(d_1919_i0_), self.f16))
        for _tupd_1 in _ingredientsd_1:
            _tupd_1[0][_tupd_1[1]] = _tupd_1[2]
        _ingredientsd_2 = []
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, ((d_1914_v0_)[default__.safeIndex(460, (d_1914_v0_).length(0))]).length(0)):
            d_1920_i1_: int = guard_loop_12_
            if (True) and (((0) <= (d_1920_i1_)) and ((d_1920_i1_) < (((d_1914_v0_)[default__.safeIndex(460, (d_1914_v0_).length(0))]).length(0)))):
                _ingredientsd_2.append(((d_1914_v0_)[default__.safeIndex(460, (d_1914_v0_).length(0))], int(d_1920_i1_), True))
        for _tupd_2 in _ingredientsd_2:
            _tupd_2[0][_tupd_2[1]] = _tupd_2[2]
        d_1921_v5_: _dafny.Array
        nw376_ = _dafny.Array(int(0), 14)
        d_1921_v5_ = nw376_
        d_1921_v5_ = d_1921_v5_
        (self).f14 = (0) - ((len(d_1916_v2_)) + (p1))
        d_1916_v2_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukilfkar")) if p2 else d_1916_v2_)
        d_1922_v6_: _dafny.Seq
        d_1922_v6_ = _dafny.SeqWithoutIsStrInference([self.f16])
        r0 = (len(d_1922_v6_)) * (default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eemycykf"))).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eemycykf")))), _dafny.CodePoint('w'))), self.f14))
        return r0

    def m9(self, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        r3: int = int(0)
        d_1923_v0_: _dafny.Array
        nw377_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_1923_v0_ = nw377_
        d_1924_v1_: _dafny.Array
        nw378_ = _dafny.Array(_dafny.Seq({}), 27)
        d_1924_v1_ = nw378_
        index325_ = default__.safeIndex(691, (d_1923_v0_).length(0))
        (d_1923_v0_)[index325_] = d_1924_v1_
        d_1925_v2_: C4
        nw379_ = C4()
        nw379_.ctor__(self.f16, self.f16, (self).f17)
        d_1925_v2_ = nw379_
        d_1926_v3_: _dafny.Set
        d_1926_v3_ = _dafny.Set({d_1925_v2_, d_1925_v2_, d_1925_v2_, (d_1925_v2_ if False else d_1925_v2_)})
        d_1927_v4_: _dafny.Set
        d_1927_v4_ = _dafny.Set({self.f14})
        index326_ = default__.safeIndex(691, (d_1923_v0_).length(0))
        rhs349_ = d_1924_v1_
        rhs350_ = not((d_1927_v4_).ispropersubset(d_1927_v4_))
        rhs351_ = d_1926_v3_
        lhs270_ = d_1923_v0_
        lhs271_ = default__.safeIndex(691, (d_1923_v0_).length(0))
        lhs270_[lhs271_] = rhs349_
        r0 = rhs350_
        d_1926_v3_ = rhs351_
        d_1928_v5_: _dafny.Array
        def lambda105_(d_1929_i1_):
            return (d_1929_i1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjeefai"))))

        init61_ = lambda105_
        nw380_ = _dafny.Array(None, 11)
        for i0_61_ in range(nw380_.length(0)):
            nw380_[i0_61_] = init61_(i0_61_)
        d_1928_v5_ = nw380_
        guard_loop_13_: int
        for guard_loop_13_ in _dafny.IntegerRange(0, (d_1928_v5_).length(0)):
            d_1930_i0_: int = guard_loop_13_
            if (True) and (((0) <= (d_1930_i0_)) and ((d_1930_i0_) < ((d_1928_v5_).length(0)))):
                (d_1928_v5_)[(d_1930_i0_)] = (d_1930_i0_) - (default__.safeDivisionInt(163, len(_dafny.Map({not((d_1925_v2_).f21): self.f14}))))
        d_1931_v6_: _dafny.Seq
        d_1931_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        d_1931_v6_ = d_1931_v6_
        d_1932_v7_: _dafny.Seq
        d_1932_v7_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1931_v6_)), self.f14, self.f14, self.f14, self.f14])
        d_1933_v8_: _dafny.Map
        d_1933_v8_ = _dafny.Map({default__.fm0(False, globalState): (d_1932_v7_).set(default__.safeIndex(347, len(d_1932_v7_)), self.f14)})
        if (((d_1933_v8_)[default__.fm0((self).f17, globalState)] if (default__.fm0((self).f17, globalState)) in (d_1933_v8_) else d_1932_v7_)) < (((d_1933_v8_)[len(_dafny.Map({self.f16: _dafny.SeqWithoutIsStrInference([self.f14, self.f14])}))] if (len(_dafny.Map({self.f16: _dafny.SeqWithoutIsStrInference([self.f14, self.f14])}))) in (d_1933_v8_) else d_1932_v7_)):
            d_1934_v9_: str
            d_1934_v9_ = _dafny.CodePoint('j')
            (globalState).f7 = (0) - ((len((d_1931_v6_).set(default__.safeIndex(self.f14, len(d_1931_v6_)), d_1934_v9_))) - (self.f14))
            d_1935_v10_: _dafny.Array
            def lambda106_(d_1936_v2_):
                def lambda107_(d_1937_i2_):
                    return (d_1936_v2_).f21

                return lambda107_

            init62_ = lambda106_(d_1925_v2_)
            nw381_ = _dafny.Array(None, 4)
            for i0_62_ in range(nw381_.length(0)):
                nw381_[i0_62_] = init62_(i0_62_)
            d_1935_v10_ = nw381_
            d_1938_v11_: _dafny.MultiSet
            d_1938_v11_ = _dafny.MultiSet([d_1935_v10_])
            d_1939_v12_: C6
            nw382_ = C6()
            nw382_.ctor__(self.f14, (d_1925_v2_).f21, not((d_1925_v2_).f21), ((d_1938_v11_).set(d_1935_v10_, default__.abs(len(d_1931_v6_)))).issubset((d_1938_v11_).set(d_1935_v10_, default__.abs(self.f14))), self.f14)
            d_1939_v12_ = nw382_
            d_1940_v13_: _dafny.Map
            d_1940_v13_ = _dafny.Map({self.f14: d_1935_v10_})
            r2 = not(((d_1939_v12_).f19) in (_dafny.Map({(0) - (len(d_1940_v13_)): d_1928_v5_})))
            r3 = self.f14
            d_1941_v14_: int
            out33_: int
            out33_ = (d_1939_v12_).m3(d_1934_v9_, len(_dafny.SeqWithoutIsStrInference([(d_1939_v12_).f19 for d_1942_i3_ in range(default__.abs(857))])), (self).f17, self.f14, globalState)
            d_1941_v14_ = out33_
        elif True:
            d_1943_v15_: str
            d_1943_v15_ = _dafny.CodePoint('m')
            d_1931_v6_ = (d_1931_v6_).set(default__.safeIndex(455, len(d_1931_v6_)), d_1943_v15_)
            d_1944_v16_: D9
            d_1944_v16_ = D9_DC16(self.f14, self.f16)
            d_1945_v17_: D9
            d_1945_v17_ = D9_DC17(d_1944_v16_)
            d_1946_v18_: _dafny.Map
            d_1946_v18_ = _dafny.Map({d_1945_v17_: self.f16})
            rhs352_ = self.f14
            rhs353_ = (self.f14) != (len(d_1946_v18_))
            rhs354_ = not((self).f17)
            rhs355_ = self.f14
            lhs272_ = globalState
            lhs273_ = self
            lhs274_ = globalState
            lhs272_.f7 = rhs352_
            r2 = rhs353_
            lhs273_.f16 = rhs354_
            lhs274_.f7 = rhs355_
            (globalState).f2 = self.f14
            d_1947_v19_: _dafny.Array
            nw383_ = _dafny.Array(False, 7)
            d_1947_v19_ = nw383_
            d_1948_v20_: _dafny.Map
            d_1948_v20_ = _dafny.Map({self.f16: d_1947_v19_})
            d_1948_v20_ = (d_1948_v20_) | (d_1948_v20_)
            d_1949_v21_: _dafny.Map
            d_1949_v21_ = _dafny.Map({self.f16: self.f16})
            r2 = not(((d_1949_v21_)[(self).f17] if ((self).f17) in (d_1949_v21_) else (d_1925_v2_).f21))
        d_1950_v22_: _dafny.Seq
        d_1950_v22_ = _dafny.SeqWithoutIsStrInference([(d_1925_v2_).f21, self.f16, False, self.f16, (d_1925_v2_).f21])
        hi17_ = self.f14
        for d_1951_i4_ in range(len((d_1950_v22_).set(default__.safeIndex(self.f14, len(d_1950_v22_)), (self).f17)), hi17_):
            r2 = (d_1925_v2_).f21
            d_1952_v23_: _dafny.Map
            d_1952_v23_ = _dafny.Map({False: self.f14})
            d_1953_v24_: C3
            nw384_ = C3()
            nw384_.ctor__(len(d_1931_v6_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blldidk")), (d_1925_v2_).f21, (self).f17)
            d_1953_v24_ = nw384_
            d_1954_v25_: _dafny.Array
            nw385_ = _dafny.Array(None, 6)
            nw385_[int(0)] = d_1953_v24_
            nw385_[int(1)] = d_1953_v24_
            nw385_[int(2)] = d_1953_v24_
            nw385_[int(3)] = d_1953_v24_
            nw385_[int(4)] = d_1953_v24_
            nw385_[int(5)] = d_1953_v24_
            d_1954_v25_ = nw385_
            index327_ = default__.safeIndex(633, (d_1954_v25_).length(0))
            (d_1954_v25_)[index327_] = d_1953_v24_
            index328_ = default__.safeIndex(633, (d_1954_v25_).length(0))
            rhs356_ = d_1928_v5_
            rhs357_ = d_1952_v23_
            rhs358_ = d_1931_v6_
            rhs359_ = d_1953_v24_
            lhs275_ = d_1954_v25_
            lhs276_ = default__.safeIndex(633, (d_1954_v25_).length(0))
            d_1928_v5_ = rhs356_
            d_1952_v23_ = rhs357_
            d_1931_v6_ = rhs358_
            lhs275_[lhs276_] = rhs359_
            d_1955_v26_: _dafny.Map
            d_1955_v26_ = _dafny.Map({self.f14: d_1951_i4_})
            d_1956_v28_: _dafny.Map
            d_1956_v28_ = _dafny.Map({default__.fm0((d_1925_v2_).f21, globalState): (self).f17})
            d_1957_v29_: D11
            def iife146_():
                coll74_ = _dafny.Map()
                compr_74_: int
                for compr_74_ in (d_1956_v28_).keys.Elements:
                    d_1958_v27_: int = compr_74_
                    if (d_1958_v27_) in (d_1956_v28_):
                        coll74_[default__.safeModuloInt(d_1958_v27_, d_1951_i4_)] = 504
                return _dafny.Map(coll74_)
            d_1957_v29_ = D11_DC19((d_1955_v26_) | (iife146_()
))
            source31_ = d_1957_v29_
            if source31_.is_DC20:
                d_1959_v30_: D3
                d_1959_v30_ = D3_DC3(self.f16)
                d_1960_v31_: _dafny.Map
                d_1960_v31_ = _dafny.Map({(d_1925_v2_).f21: (d_1925_v2_).f21})
                d_1961_v32_: C2
                nw386_ = C2()
                def iife147_(_pat_let36_0):
                    def iife148_(d_1962_dt__update__tmp_h0_):
                        def iife149_(_pat_let37_0):
                            def iife150_(d_1963_dt__update_hcf3_h0_):
                                return D3_DC3(d_1963_dt__update_hcf3_h0_)
                            return iife150_(_pat_let37_0)
                        return iife149_(not(not((self).f17)))
                    return iife148_(_pat_let36_0)
                nw386_.ctor__(iife147_(d_1959_v30_), (d_1925_v2_).f21, (((d_1960_v31_)[(d_1925_v2_).f21] if ((d_1925_v2_).f21) in (d_1960_v31_) else (d_1925_v2_).f21)) or ((self).f17))
                d_1961_v32_ = nw386_
                d_1964_v33_: _dafny.Map
                d_1964_v33_ = _dafny.Map({(d_1925_v2_).f21: d_1932_v7_})
                index329_ = default__.safeIndex(372, (d_1928_v5_).length(0))
                (d_1928_v5_)[index329_] = d_1951_i4_
                d_1965_v34_: _dafny.MultiSet
                d_1965_v34_ = _dafny.MultiSet([True])
                d_1966_v35_: str
                d_1966_v35_ = _dafny.CodePoint('y')
                d_1967_v36_: _dafny.MultiSet
                d_1967_v36_ = _dafny.MultiSet([(d_1953_v24_).f22, d_1951_i4_])
                d_1968_v38_: _dafny.Array
                nw387_ = _dafny.Array(None, 22)
                nw387_[int(0)] = (d_1925_v2_).f21
                nw387_[int(1)] = self.f16
                nw387_[int(2)] = False
                nw387_[int(3)] = (self).f17
                nw387_[int(4)] = (d_1925_v2_).f21
                nw387_[int(5)] = self.f16
                nw387_[int(6)] = True
                nw387_[int(7)] = (d_1925_v2_).f21
                nw387_[int(8)] = self.f16
                nw387_[int(9)] = False
                nw387_[int(10)] = (d_1925_v2_).f21
                nw387_[int(11)] = self.f16
                nw387_[int(12)] = (d_1925_v2_).f21
                nw387_[int(13)] = (self).f17
                nw387_[int(14)] = (d_1925_v2_).f21
                nw387_[int(15)] = (d_1925_v2_).f21
                nw387_[int(16)] = self.f16
                nw387_[int(17)] = default__.fm43((d_1925_v2_).f21, globalState)
                nw387_[int(18)] = (d_1925_v2_).f21
                nw387_[int(19)] = (self).f17
                nw387_[int(20)] = (d_1925_v2_).f21
                nw387_[int(21)] = self.f16
                d_1968_v38_ = nw387_
                d_1969_v39_: D9
                def iife151_():
                    coll75_ = _dafny.Set()
                    compr_75_: int
                    for compr_75_ in (d_1927_v4_).Elements:
                        d_1970_v37_: int = compr_75_
                        if (d_1970_v37_) in (d_1927_v4_):
                            coll75_ = coll75_.union(_dafny.Set([default__.safeDivisionInt(d_1970_v37_, 891)]))
                    return _dafny.Set(coll75_)
                d_1969_v39_ = D9_DC15((self).f17, (d_1925_v2_).f21, iife151_()
, d_1928_v5_, d_1968_v38_)
                d_1971_v40_: _dafny.Array
                nw388_ = _dafny.Array(None, 6)
                nw388_[int(0)] = d_1969_v39_
                nw388_[int(1)] = D9_DC15((d_1950_v22_)[default__.safeIndex(len(d_1955_v26_), len(d_1950_v22_))], default__.fm31((d_1925_v2_).f21, (d_1925_v2_).f21, globalState), d_1927_v4_, d_1928_v5_, d_1968_v38_)
                nw388_[int(2)] = d_1969_v39_
                nw388_[int(3)] = d_1969_v39_
                nw388_[int(4)] = d_1969_v39_
                nw388_[int(5)] = d_1969_v39_
                d_1971_v40_ = nw388_
                d_1972_v41_: D12
                d_1972_v41_ = D12_DC24((d_1925_v2_).f21, d_1950_v22_, _dafny.SeqWithoutIsStrInference([d_1971_v40_]), (d_1925_v2_).f21, default__.fm62(globalState))
                d_1973_v42_: _dafny.Array
                nw389_ = _dafny.Array(None, 19)
                nw389_[int(0)] = (d_1950_v22_)[default__.safeIndex(165, len(d_1950_v22_))]
                nw389_[int(1)] = (d_1925_v2_).f21
                nw389_[int(2)] = (_dafny.MultiSet(d_1950_v22_)).isdisjoint(d_1965_v34_)
                nw389_[int(3)] = (self).f17
                nw389_[int(4)] = (d_1925_v2_).f21
                nw389_[int(5)] = default__.fm31((self).f17, self.f16, globalState)
                nw389_[int(6)] = (d_1966_v35_) not in ((d_1953_v24_).f23)
                nw389_[int(7)] = (d_1925_v2_).f21
                nw389_[int(8)] = (d_1925_v2_).f21
                nw389_[int(9)] = self.f16
                nw389_[int(10)] = (self).f17
                nw389_[int(11)] = ((0) - (self.f14)) not in (d_1967_v36_)
                nw389_[int(12)] = self.f16
                nw389_[int(13)] = (d_1972_v41_).cf33
                nw389_[int(14)] = (d_1925_v2_).f21
                nw389_[int(15)] = ((d_1953_v24_).f22) < ((d_1925_v2_).fm12(self.f14, globalState))
                nw389_[int(16)] = (self).f17
                nw389_[int(17)] = True
                nw389_[int(18)] = (d_1925_v2_).f21
                d_1973_v42_ = nw389_
                index330_ = default__.safeIndex(547, (d_1973_v42_).length(0))
                (d_1973_v42_)[index330_] = False
                d_1974_v43_: _dafny.Set
                d_1974_v43_ = _dafny.Set({self.f16})
                index331_ = default__.safeIndex(372, (d_1928_v5_).length(0))
                index332_ = default__.safeIndex(547, (d_1973_v42_).length(0))
                rhs360_ = _dafny.Map({(d_1925_v2_).f21: (d_1932_v7_) + (d_1932_v7_)})
                rhs361_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))), len(d_1932_v7_))
                rhs362_ = (d_1925_v2_).f21
                rhs363_ = ((d_1974_v43_) | (d_1974_v43_)).ispropersubset(d_1974_v43_)
                lhs277_ = d_1928_v5_
                lhs278_ = default__.safeIndex(372, (d_1928_v5_).length(0))
                lhs279_ = d_1973_v42_
                lhs280_ = default__.safeIndex(547, (d_1973_v42_).length(0))
                d_1964_v33_ = rhs360_
                lhs277_[lhs278_] = rhs361_
                lhs279_[lhs280_] = rhs362_
                r2 = rhs363_
                d_1975_v44_: _dafny.Map
                d_1975_v44_ = _dafny.Map({(d_1953_v24_).f23: not((d_1925_v2_).f21)})
                d_1976_v45_: D21
                d_1976_v45_ = D21_DC43(_dafny.MultiSet([d_1966_v35_, d_1966_v35_]))
                d_1977_v46_: _dafny.MultiSet
                d_1977_v46_ = _dafny.MultiSet([d_1966_v35_])
                d_1978_v47_: _dafny.Map
                d_1978_v47_ = _dafny.Map({d_1966_v35_: d_1977_v46_})
                d_1979_v48_: _dafny.Seq
                d_1979_v48_ = _dafny.SeqWithoutIsStrInference([(d_1976_v45_).cf68, ((d_1978_v47_)[_dafny.CodePoint('m')] if (_dafny.CodePoint('m')) in (d_1978_v47_) else _dafny.MultiSet(d_1931_v6_))])
                d_1980_v49_: D21
                d_1980_v49_ = D21_DC44((d_1953_v24_).f22, (d_1925_v2_).f21)
                index333_ = default__.safeIndex(547, (d_1973_v42_).length(0))
                rhs364_ = (d_1953_v24_).f22
                rhs365_ = ((d_1953_v24_).f23) not in (d_1975_v44_)
                rhs366_ = ((d_1979_v48_)[default__.safeIndex(d_1951_i4_, len(d_1979_v48_))]).isdisjoint(d_1977_v46_)
                rhs367_ = ((d_1953_v24_).f22) > (default__.safeModuloInt((d_1980_v49_).cf69, (d_1953_v24_).f22))
                lhs281_ = globalState
                lhs282_ = d_1973_v42_
                lhs283_ = default__.safeIndex(547, (d_1973_v42_).length(0))
                lhs281_.f2 = rhs364_
                r2 = rhs365_
                r0 = rhs366_
                lhs282_[lhs283_] = rhs367_
                def lambda108_(d_1981_v24_, d_1982_v26_):
                    def lambda109_(d_1983_i5_):
                        return ((d_1981_v24_).f22) >= (((d_1982_v26_)[self.f14] if (self.f14) in (d_1982_v26_) else self.f14))

                    return lambda109_

                init63_ = lambda108_(d_1953_v24_, d_1955_v26_)
                nw390_ = _dafny.Array(None, 13)
                for i0_63_ in range(nw390_.length(0)):
                    nw390_[i0_63_] = init63_(i0_63_)
                d_1968_v38_ = nw390_
            elif source31_.is_DC19:
                d_1984___mcc_h0_ = source31_.cf27
                d_1985_cf27_ = d_1984___mcc_h0_
                index334_ = default__.safeIndex(298, (d_1928_v5_).length(0))
                (d_1928_v5_)[index334_] = len(_dafny.SeqWithoutIsStrInference([(d_1953_v24_).f22 for d_1986_i6_ in range(default__.abs(812))]))
                d_1987_v50_: _dafny.MultiSet
                d_1987_v50_ = _dafny.MultiSet([self.f14])
                index335_ = default__.safeIndex(298, (d_1928_v5_).length(0))
                (d_1928_v5_)[index335_] = (d_1987_v50_).cardinality
                (globalState).f7 = (d_1953_v24_).f22
                d_1988_v51_: int
                out34_: int
                out34_ = (d_1925_v2_).m3(((d_1953_v24_).f23)[default__.safeIndex(d_1951_i4_, len((d_1953_v24_).f23))], self.f14, (d_1925_v2_).f21, d_1951_i4_, globalState)
                d_1988_v51_ = out34_
                d_1989_v52_: _dafny.Array
                def lambda110_(d_1990_i7_):
                    return (self).f17

                init64_ = lambda110_
                nw391_ = _dafny.Array(None, 28)
                for i0_64_ in range(nw391_.length(0)):
                    nw391_[i0_64_] = init64_(i0_64_)
                d_1989_v52_ = nw391_
                d_1991_v53_: _dafny.Set
                d_1991_v53_ = _dafny.Set({d_1989_v52_})
                d_1985_cf27_ = (d_1985_cf27_).set(22, len(d_1991_v53_))
            elif True:
                d_1992___mcc_h1_ = source31_.cf28
                d_1993_cf28_ = d_1992___mcc_h1_
                d_1994_v54_: str
                d_1994_v54_ = _dafny.CodePoint('v')
                d_1995_v55_: _dafny.MultiSet
                d_1995_v55_ = _dafny.MultiSet([d_1994_v54_, d_1994_v54_, d_1994_v54_])
                r2 = default__.fm27(((d_1995_v55_)[d_1994_v54_] if (d_1994_v54_) in (d_1995_v55_) else ((d_1995_v55_)[d_1994_v54_] if (d_1994_v54_) in (d_1995_v55_) else self.f14)), (d_1953_v24_).f22, 676, _dafny.Map({(d_1953_v24_).f22: (d_1953_v24_).f22}), globalState)
                d_1996_v56_: D3
                d_1996_v56_ = D3_DC3((d_1925_v2_).f21)
                d_1997_v57_: T1
                nw392_ = C2()
                nw392_.ctor__(d_1996_v56_, (d_1925_v2_).f21, (d_1925_v2_).f21)
                d_1997_v57_ = nw392_
                d_1998_v58_: _dafny.Map
                d_1998_v58_ = _dafny.Map({(d_1950_v22_)[default__.safeIndex((d_1953_v24_).f22, len(d_1950_v22_))]: d_1997_v57_})
                d_1997_v57_ = (d_1997_v57_ if (d_1925_v2_).f21 else (((d_1998_v58_)[(d_1925_v2_).f21] if ((d_1925_v2_).f21) in (d_1998_v58_) else d_1997_v57_) if (d_1925_v2_).f21 else d_1997_v57_))
                (self).f16 = (d_1996_v56_).cf3
                r0 = not(((d_1953_v24_).f22) > (self.f14))
            d_1999_v59_: _dafny.Array
            nw393_ = _dafny.Array(None, 2)
            nw393_[int(0)] = (d_1955_v26_) | (d_1955_v26_)
            nw393_[int(1)] = (d_1955_v26_ if (self).f17 else d_1955_v26_)
            d_1999_v59_ = nw393_
            index336_ = default__.safeIndex(236, (d_1999_v59_).length(0))
            (d_1999_v59_)[index336_] = d_1955_v26_
            index337_ = default__.safeIndex(236, (d_1999_v59_).length(0))
            (d_1999_v59_)[index337_] = (d_1953_v24_).fm3(globalState)
        d_2000_v60_: _dafny.Map
        d_2000_v60_ = _dafny.Map({(self).f17: (self).f17})
        d_2001_v61_: _dafny.MultiSet
        d_2001_v61_ = _dafny.MultiSet([self.f16])
        d_2002_v62_: _dafny.Map
        d_2002_v62_ = _dafny.Map({self.f14: len(d_1950_v22_)})
        d_2003_v63_: _dafny.Map
        d_2003_v63_ = _dafny.Map({False: len(d_1932_v7_)})
        d_2004_v64_: D3
        d_2004_v64_ = D3_DC4((self).f17, self.f14, len(d_2003_v63_), self.f16, False)
        d_2005_v65_: _dafny.Map
        d_2005_v65_ = _dafny.Map({(self).f17: default__.fm28(self.f16, d_2002_v62_, d_2004_v64_, d_2000_v60_, globalState)})
        d_2006_v66_: C6
        nw394_ = C6()
        nw394_.ctor__(self.f14, (((d_2000_v60_)[self.f16] if (self.f16) in (d_2000_v60_) else not((d_1925_v2_).f21))) not in (d_2001_v61_), self.f16, ((d_2005_v65_).set((d_1925_v2_).f21, ((d_2005_v65_)[True] if (True) in (d_2005_v65_) else d_1931_v6_))) == (d_2005_v65_), 375)
        d_2006_v66_ = nw394_
        r0 = (d_1925_v2_).f21
        d_2007_v67_: _dafny.Seq
        d_2007_v67_ = d_1950_v22_
        pat_let_tv28_ = d_2006_v66_
        pat_let_tv29_ = d_2000_v60_
        pat_let_tv30_ = d_2006_v66_
        pat_let_tv31_ = d_2006_v66_
        pat_let_tv32_ = d_2000_v60_
        def lambda111_(source32_):
            d_2008___mcc_h2_ = source32_
            d_2009_cf0_ = d_2008___mcc_h2_
            return _dafny.Set({(self).f17, (pat_let_tv28_).f20, not(not(self.f16)), False, not(((pat_let_tv29_)[(pat_let_tv30_).f20] if ((pat_let_tv31_).f20) in (pat_let_tv32_) else (self).f17))})

        r1 = lambda111_(d_2007_v67_)
        r2 = (d_2006_v66_).f20
        r3 = self.f14
        return r0, r1, r2, r3


class C8(T1, T0, T2):
    def  __init__(self):
        self._f16: bool = False
        self._f14: int = int(0)
        self._f17: bool = False
        self.f18: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
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
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f18, f16, f17, f14):
        (self).f18 = f18
        (self).f16 = f16
        (self)._f17 = f17
        (self).f14 = f14

    def fm3(self, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f14 for d_2010_i0_ in range(default__.abs(88))])): self.f14})

    def fm12(self, p0, globalState):
        return ((0) - (self.f14)) * ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbtospxwi"))) if self.f16 else self.f14))

    def fm13(self, p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(self).f17, self.f16, (self).f17])

    def fm14(self, globalState):
        return self.f16

    def fm15(self, globalState):
        return _dafny.MultiSet([self.f18, self.f18, self.f18, (self.f18) + (self.f18), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kk"))])

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_2011_v0_: _dafny.Seq
        d_2011_v0_ = _dafny.SeqWithoutIsStrInference([True, p1])
        d_2012_v1_: C4
        nw395_ = C4()
        nw395_.ctor__(self.f16, self.f16, (d_2011_v0_)[default__.safeIndex(p0, len(d_2011_v0_))])
        d_2012_v1_ = nw395_
        d_2013_v2_: _dafny.Set
        d_2013_v2_ = _dafny.Set({self.f14, p0, self.f14, self.f14, default__.fm0((self).f17, globalState)})
        (self).f16 = ((default__.fm49(p0, p0, self.f16, self.f14, globalState)).intersection(d_2013_v2_)).isdisjoint(_dafny.Set({p0}))
        d_2014_v3_: _dafny.Seq
        d_2014_v3_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14])
        d_2015_v4_: _dafny.Seq
        d_2015_v4_ = _dafny.SeqWithoutIsStrInference([d_2014_v3_])
        d_2016_i0_: int
        d_2016_i0_ = 0
        with _dafny.label("10"):
            while not(not((d_2014_v3_) <= ((d_2015_v4_)[default__.safeIndex(len(self.f18), len(d_2015_v4_))]))):
                with _dafny.c_label("10"):
                    if (d_2016_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_2016_i0_ = (d_2016_i0_) + (1)
                    (globalState).f7 = p0
                    d_2017_v5_: _dafny.Set
                    d_2017_v5_ = _dafny.Set({(d_2012_v1_).f21, False, (d_2012_v1_).f21, self.f16, self.f16})
                    d_2018_v6_: T0
                    nw396_ = C1()
                    nw396_.ctor__(self.f16, p0)
                    d_2018_v6_ = nw396_
                    d_2019_v7_: _dafny.MultiSet
                    d_2019_v7_ = _dafny.MultiSet([d_2014_v3_])
                    d_2020_v8_: D16
                    d_2020_v8_ = D16_DC34(not(p1), d_2018_v6_, (d_2019_v7_).cardinality)
                    d_2021_v9_: _dafny.Array
                    nw397_ = _dafny.Array(None, 28)
                    nw397_[int(0)] = self.f14
                    nw397_[int(1)] = self.f14
                    nw397_[int(2)] = p0
                    nw397_[int(3)] = (p0) * (len(self.f18))
                    nw397_[int(4)] = p0
                    nw397_[int(5)] = self.f14
                    nw397_[int(6)] = (0) - (p0)
                    nw397_[int(7)] = len(d_2017_v5_)
                    nw397_[int(8)] = default__.safeModuloInt(p0, p0)
                    nw397_[int(9)] = p0
                    nw397_[int(10)] = len((self.f18) + (self.f18))
                    nw397_[int(11)] = (self.f14) - (p0)
                    nw397_[int(12)] = (p0) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2022_i1_ in range(default__.abs(182))])))
                    nw397_[int(13)] = 652
                    nw397_[int(14)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([self.f16, p1])), (0) - (len(d_2014_v3_)))
                    nw397_[int(15)] = self.f14
                    nw397_[int(16)] = p0
                    nw397_[int(17)] = (0) - (len(d_2014_v3_))
                    nw397_[int(18)] = self.f14
                    nw397_[int(19)] = p0
                    nw397_[int(20)] = p0
                    nw397_[int(21)] = (p0) - ((0) - (self.f14))
                    nw397_[int(22)] = 567
                    nw397_[int(23)] = (p0 if (self).f17 else 397)
                    nw397_[int(24)] = (d_2020_v8_).cf53
                    nw397_[int(25)] = d_2018_v6_.f14
                    nw397_[int(26)] = 672
                    nw397_[int(27)] = len((d_2011_v0_) + (d_2011_v0_))
                    d_2021_v9_ = nw397_
                    index338_ = default__.safeIndex(864, (d_2021_v9_).length(0))
                    (d_2021_v9_)[index338_] = d_2018_v6_.f14
                    index339_ = default__.safeIndex(864, (d_2021_v9_).length(0))
                    rhs368_ = (0) - ((self.f14) - (d_2018_v6_.f14))
                    rhs369_ = len(d_2011_v0_)
                    rhs370_ = d_2018_v6_.f14
                    rhs371_ = p1
                    lhs284_ = d_2021_v9_
                    lhs285_ = default__.safeIndex(864, (d_2021_v9_).length(0))
                    lhs286_ = globalState
                    lhs287_ = self
                    lhs288_ = self
                    lhs284_[lhs285_] = rhs368_
                    lhs286_.f2 = rhs369_
                    lhs287_.f14 = rhs370_
                    lhs288_.f16 = rhs371_
                    rhs372_ = (0) - ((d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))])
                    rhs373_ = (d_2014_v3_) + (((_dafny.SeqWithoutIsStrInference([default__.fm0((d_2012_v1_).f21, globalState)])) + (d_2014_v3_)).set(default__.safeIndex((d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))], len((_dafny.SeqWithoutIsStrInference([default__.fm0((d_2012_v1_).f21, globalState)])) + (d_2014_v3_))), (d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))]))
                    r0 = rhs372_
                    d_2014_v3_ = rhs373_
                    if p1:
                        d_2023_v10_: str
                        d_2023_v10_ = _dafny.CodePoint('r')
                        d_2024_v11_: _dafny.Map
                        d_2024_v11_ = _dafny.Map({_dafny.Set({d_2023_v10_, d_2023_v10_}): (d_2012_v1_).f21})
                        d_2024_v11_ = (d_2024_v11_).set(_dafny.Set({_dafny.CodePoint('w'), d_2023_v10_, d_2023_v10_, d_2023_v10_, d_2023_v10_}), self.f16)
                        (self).f16 = (self).f17
                        (globalState).f7 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkxdimuqr")))
                        d_2025_v12_: _dafny.Seq
                        d_2025_v12_ = _dafny.SeqWithoutIsStrInference([self.f18])
                        d_2026_v13_: _dafny.Seq
                        d_2026_v13_ = _dafny.SeqWithoutIsStrInference([(d_2025_v12_)[default__.safeIndex(self.f14, len(d_2025_v12_))], self.f18])
                        d_2027_v14_: _dafny.Array
                        nw398_ = _dafny.Array(None, 5)
                        nw398_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dess"))) + (self.f18)
                        nw398_[int(1)] = self.f18
                        nw398_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlwjpaav"))) + ((d_2026_v13_)[default__.safeIndex((self).fm12(-793, globalState), len(d_2026_v13_))])
                        nw398_[int(3)] = self.f18
                        nw398_[int(4)] = (self.f18) + (self.f18)
                        d_2027_v14_ = nw398_
                        index340_ = default__.safeIndex(546, (d_2027_v14_).length(0))
                        (d_2027_v14_)[index340_] = (self.f18) + (self.f18)
                        index341_ = default__.safeIndex(546, (d_2027_v14_).length(0))
                        (d_2027_v14_)[index341_] = self.f18
                        d_2028_v15_: _dafny.Map
                        d_2028_v15_ = _dafny.Map({(d_2012_v1_).f21: (d_2027_v14_)[default__.safeIndex(546, (d_2027_v14_).length(0))]})
                        d_2029_v16_: _dafny.Map
                        d_2029_v16_ = _dafny.Map({self.f14: d_2023_v10_})
                        d_2030_v17_: _dafny.Map
                        d_2030_v17_ = _dafny.Map({(d_2012_v1_).f21: len((d_2027_v14_)[default__.safeIndex(546, (d_2027_v14_).length(0))])})
                        rhs374_ = (0) - ((0) - ((d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))]))
                        rhs375_ = not (default__.fm21(globalState)) or ((d_2012_v1_).f21)
                        rhs376_ = (0) - (len(((((d_2028_v15_)[(d_2012_v1_).f21] if ((d_2012_v1_).f21) in (d_2028_v15_) else _dafny.SeqWithoutIsStrInference([d_2023_v10_ for d_2031_i2_ in range(default__.abs(352))]))).set(default__.safeIndex((_dafny.MultiSet(d_2014_v3_)).cardinality, len(((d_2028_v15_)[(d_2012_v1_).f21] if ((d_2012_v1_).f21) in (d_2028_v15_) else _dafny.SeqWithoutIsStrInference([d_2023_v10_ for d_2032_i2_ in range(default__.abs(352))])))), ((d_2029_v16_)[556] if (556) in (d_2029_v16_) else d_2023_v10_))) + (self.f18)))
                        rhs377_ = (366 if (d_2030_v17_) == (d_2030_v17_) else (d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))])
                        lhs289_ = globalState
                        lhs290_ = self
                        lhs291_ = d_2018_v6_
                        lhs292_ = self
                        lhs289_.f7 = rhs374_
                        lhs290_.f16 = rhs375_
                        lhs291_.f14 = rhs376_
                        lhs292_.f14 = rhs377_
                    elif True:
                        d_2033_v18_: _dafny.Array
                        nw399_ = _dafny.Array(D9.default()(), 12)
                        d_2033_v18_ = nw399_
                        d_2034_v19_: _dafny.Seq
                        d_2034_v19_ = _dafny.SeqWithoutIsStrInference([d_2033_v18_])
                        d_2035_v20_: D12
                        d_2035_v20_ = D12_DC23(True, (0) - ((d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))]), p0)
                        d_2036_v21_: _dafny.MultiSet
                        d_2036_v21_ = _dafny.MultiSet([(d_2012_v1_).f21, (d_2035_v20_).cf30])
                        d_2037_v22_: _dafny.Map
                        d_2037_v22_ = _dafny.Map({d_2018_v6_.f14: self.f14})
                        d_2038_v23_: D3
                        d_2038_v23_ = D3_DC3((d_2012_v1_).f21)
                        d_2039_v24_: D12
                        d_2039_v24_ = D12_DC24(self.f16, d_2011_v0_, d_2034_v19_, (d_2012_v1_).f21, _dafny.SeqWithoutIsStrInference([D3_DC3(default__.fm27((d_2036_v21_).cardinality, (d_2014_v3_)[default__.safeIndex(self.f14, len(d_2014_v3_))], self.f14, d_2037_v22_, globalState)), d_2038_v23_]))
                        d_2040_v25_: _dafny.Map
                        d_2040_v25_ = _dafny.Map({d_2039_v24_: (d_2012_v1_).f21})
                        d_2041_v26_: _dafny.Map
                        d_2041_v26_ = _dafny.Map({not(p1): d_2040_v25_})
                        d_2042_v27_: _dafny.Map
                        d_2042_v27_ = _dafny.Map({self.f14: d_2040_v25_})
                        d_2043_v28_: _dafny.Array
                        nw400_ = _dafny.Array(None, 15)
                        nw400_[int(0)] = d_2040_v25_
                        nw400_[int(1)] = (d_2040_v25_) | (_dafny.Map({d_2039_v24_: self.f16}))
                        nw400_[int(2)] = (d_2040_v25_) | (d_2040_v25_)
                        nw400_[int(3)] = (d_2040_v25_) | (d_2040_v25_)
                        nw400_[int(4)] = (((d_2041_v26_)[(d_2012_v1_).f21] if ((d_2012_v1_).f21) in (d_2041_v26_) else d_2040_v25_) if p1 else d_2040_v25_)
                        nw400_[int(5)] = d_2040_v25_
                        nw400_[int(6)] = (d_2040_v25_) | (d_2040_v25_)
                        nw400_[int(7)] = (d_2040_v25_) | (d_2040_v25_)
                        nw400_[int(8)] = d_2040_v25_
                        nw400_[int(9)] = _dafny.Map({d_2039_v24_: (d_2011_v0_)[default__.safeIndex((d_2021_v9_)[default__.safeIndex(864, (d_2021_v9_).length(0))], len(d_2011_v0_))]})
                        nw400_[int(10)] = d_2040_v25_
                        nw400_[int(11)] = ((d_2042_v27_)[p0] if (p0) in (d_2042_v27_) else d_2040_v25_)
                        nw400_[int(12)] = d_2040_v25_
                        nw400_[int(13)] = d_2040_v25_
                        nw400_[int(14)] = d_2040_v25_
                        d_2043_v28_ = nw400_
                        nw401_ = _dafny.Array(None, 23)
                        nw401_[int(0)] = d_2040_v25_
                        nw401_[int(1)] = d_2040_v25_
                        nw401_[int(2)] = (d_2040_v25_) | (d_2040_v25_)
                        nw401_[int(3)] = d_2040_v25_
                        nw401_[int(4)] = d_2040_v25_
                        nw401_[int(5)] = (d_2040_v25_) | (d_2040_v25_)
                        nw401_[int(6)] = d_2040_v25_
                        nw401_[int(7)] = d_2040_v25_
                        nw401_[int(8)] = (d_2040_v25_) | (d_2040_v25_)
                        nw401_[int(9)] = d_2040_v25_
                        nw401_[int(10)] = (d_2040_v25_).set(d_2039_v24_, (d_2012_v1_).f21)
                        nw401_[int(11)] = ((d_2040_v25_).set(d_2039_v24_, (d_2012_v1_).f21)).set(d_2039_v24_, not((self).f17))
                        nw401_[int(12)] = (_dafny.Map({d_2039_v24_: (d_2012_v1_).f21})) | (_dafny.Map({d_2039_v24_: self.f16}))
                        nw401_[int(13)] = d_2040_v25_
                        nw401_[int(14)] = d_2040_v25_
                        nw401_[int(15)] = d_2040_v25_
                        nw401_[int(16)] = d_2040_v25_
                        nw401_[int(17)] = d_2040_v25_
                        nw401_[int(18)] = d_2040_v25_
                        nw401_[int(19)] = d_2040_v25_
                        nw401_[int(20)] = d_2040_v25_
                        nw401_[int(21)] = _dafny.Map({d_2039_v24_: (self).f17})
                        nw401_[int(22)] = d_2040_v25_
                        d_2043_v28_ = nw401_
                        (self).f16 = (d_2012_v1_).f21
                        d_2037_v22_ = d_2037_v22_
                        d_2044_v29_: _dafny.Array
                        nw402_ = _dafny.Array(_dafny.Seq({}), 9)
                        d_2044_v29_ = nw402_
                        d_2044_v29_ = d_2044_v29_
                        (globalState).f2 = default__.fm0((d_2012_v1_).f21, globalState)
                    pass
            pass
        d_2045_v30_: _dafny.Array
        nw403_ = _dafny.Array(_dafny.MultiSet({}), 25)
        d_2045_v30_ = nw403_
        index342_ = default__.safeIndex(412, (d_2045_v30_).length(0))
        (d_2045_v30_)[index342_] = default__.fm1(globalState)
        d_2046_v31_: _dafny.MultiSet
        d_2046_v31_ = _dafny.MultiSet([self.f16])
        d_2047_v32_: _dafny.Map
        d_2047_v32_ = _dafny.Map({_dafny.Set({p1}): d_2013_v2_})
        d_2048_v33_: _dafny.Set
        d_2048_v33_ = _dafny.Set({(self).f17, not((d_2012_v1_).f21), (d_2012_v1_).f21, (d_2012_v1_).f21})
        index343_ = default__.safeIndex(412, (d_2045_v30_).length(0))
        rhs378_ = (d_2046_v31_).set((len((d_2047_v32_).set(d_2048_v33_, d_2013_v2_))) < (p0), default__.abs(self.f14))
        rhs379_ = p0
        lhs293_ = d_2045_v30_
        lhs294_ = default__.safeIndex(412, (d_2045_v30_).length(0))
        lhs293_[lhs294_] = rhs378_
        r0 = rhs379_
        (globalState).f7 = default__.safeModuloInt(153, p0)
        d_2049_i3_: int
        d_2049_i3_ = 0
        with _dafny.label("11"):
            while False:
                with _dafny.c_label("11"):
                    if (d_2049_i3_) >= (100):
                        raise _dafny.Break("11")
                    d_2049_i3_ = (d_2049_i3_) + (1)
                    (globalState).f2 = (0) - ((0) - ((p0) * (p0)))
                    (self).f14 = (82) - (len(self.f18))
                    d_2045_v30_ = d_2045_v30_
                    d_2050_v34_: str
                    d_2050_v34_ = _dafny.CodePoint('w')
                    d_2051_v35_: _dafny.Map
                    d_2051_v35_ = _dafny.Map({(self).f17: d_2050_v34_})
                    d_2052_v36_: _dafny.Map
                    d_2052_v36_ = _dafny.Map({(_dafny.MultiSet(d_2011_v0_)).cardinality: (self).f17})
                    d_2053_v37_: _dafny.Map
                    d_2053_v37_ = _dafny.Map({default__.fm17((d_2012_v1_).f21, len(d_2051_v35_), self.f14, len(d_2052_v36_), globalState): d_2050_v34_})
                    d_2054_v38_: _dafny.Map
                    d_2054_v38_ = _dafny.Map({p0: d_2050_v34_})
                    rhs380_ = not ((self).f17) or (not((self).f17))
                    rhs381_ = ((d_2053_v37_)[(d_2054_v38_) == (d_2054_v38_)] if ((d_2054_v38_) == (d_2054_v38_)) in (d_2053_v37_) else d_2050_v34_)
                    rhs382_ = self.f14
                    lhs295_ = self
                    lhs296_ = globalState
                    lhs295_.f16 = rhs380_
                    d_2050_v34_ = rhs381_
                    lhs296_.f7 = rhs382_
                    pass
            pass
        r0 = p0
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_2055_v0_: _dafny.Seq
        d_2055_v0_ = _dafny.SeqWithoutIsStrInference([self.f16, True])
        d_2056_v1_: _dafny.Array
        nw404_ = _dafny.Array(D9.default()(), 26)
        d_2056_v1_ = nw404_
        d_2057_v2_: _dafny.Seq
        d_2057_v2_ = _dafny.SeqWithoutIsStrInference([d_2056_v1_, d_2056_v1_])
        d_2058_v3_: D3
        d_2058_v3_ = D3_DC3((self).f17)
        d_2059_v4_: _dafny.Seq
        d_2059_v4_ = _dafny.SeqWithoutIsStrInference([d_2058_v3_])
        d_2060_v5_: D12
        d_2060_v5_ = D12_DC24(self.f16, d_2055_v0_, d_2057_v2_, p2, d_2059_v4_)
        d_2061_v6_: D20
        d_2061_v6_ = D20_DC41(_dafny.MultiSet([d_2060_v5_]))
        d_2062_v7_: _dafny.Map
        d_2062_v7_ = _dafny.Map({default__.safeModuloInt(p1, (0) - (self.f14)): (d_2061_v6_ if self.f16 else d_2061_v6_)})
        d_2063_v8_: T0
        nw405_ = C6()
        nw405_.ctor__(self.f14, (self).f17, self.f16, (self).f17, p1)
        d_2063_v8_ = nw405_
        d_2064_v9_: _dafny.Map
        d_2064_v9_ = _dafny.Map({self.f14: d_2063_v8_})
        d_2065_v10_: D20
        d_2065_v10_ = D20_DC42(default__.safeModuloInt(393, p3), (p2 if p2 else (D14_DC28(p3, p2, ((d_2064_v9_)[self.f14] if (self.f14) in (d_2064_v9_) else d_2063_v8_))).cf41), d_2063_v8_)
        d_2066_v11_: _dafny.Seq
        d_2066_v11_ = _dafny.SeqWithoutIsStrInference([d_2062_v7_])
        d_2067_v12_: _dafny.MultiSet
        d_2067_v12_ = _dafny.MultiSet([p0])
        d_2068_v13_: _dafny.Map
        d_2068_v13_ = _dafny.Map({(self).f17: (d_2066_v11_)[default__.safeIndex((d_2067_v12_).cardinality, len(d_2066_v11_))]})
        def iife152_(_pat_let38_0):
            def iife153_(d_2069_dt__update__tmp_h0_):
                def iife154_(_pat_let39_0):
                    def iife155_(d_2070_dt__update_hcf65_h0_):
                        return D20_DC42(d_2070_dt__update_hcf65_h0_, (d_2069_dt__update__tmp_h0_).cf66, (d_2069_dt__update__tmp_h0_).cf67)
                    return iife155_(_pat_let39_0)
                return iife154_(len(_dafny.SeqWithoutIsStrInference([(self).f17, True, self.f16])))
            return iife153_(_pat_let38_0)
        rhs383_ = ((d_2062_v7_) | (d_2062_v7_)) | (((d_2068_v13_)[self.f16] if (self.f16) in (d_2068_v13_) else d_2062_v7_))
        rhs384_ = (((self.f18) + (self.f18)).set(default__.safeIndex(p1, len((self.f18) + (self.f18))), default__.fm34(d_2063_v8_.f14, p1, self.f16, p3, globalState))).set(default__.safeIndex(default__.fm0(True, globalState), len(((self.f18) + (self.f18)).set(default__.safeIndex(p1, len((self.f18) + (self.f18))), default__.fm34(d_2063_v8_.f14, p1, self.f16, p3, globalState)))), p0)
        rhs385_ = (self).f17
        rhs386_ = iife152_(d_2065_v10_)
        lhs297_ = self
        lhs298_ = self
        d_2062_v7_ = rhs383_
        lhs297_.f18 = rhs384_
        lhs298_.f16 = rhs385_
        d_2065_v10_ = rhs386_
        (self).f14 = default__.safeModuloInt(p3, self.f14)
        d_2071_i0_: int
        d_2071_i0_ = 0
        with _dafny.label("12"):
            while not (True) or (p2):
                with _dafny.c_label("12"):
                    if (d_2071_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_2071_i0_ = (d_2071_i0_) + (1)
                    d_2072_v14_: _dafny.Map
                    d_2072_v14_ = _dafny.Map({d_2063_v8_.f14: d_2063_v8_.f14})
                    d_2073_v15_: _dafny.Set
                    d_2073_v15_ = _dafny.Set({-126, d_2063_v8_.f14, d_2063_v8_.f14})
                    d_2074_v16_: _dafny.Map
                    d_2074_v16_ = _dafny.Map({d_2073_v15_: self.f14})
                    d_2075_v17_: _dafny.Seq
                    d_2075_v17_ = _dafny.SeqWithoutIsStrInference([p3])
                    d_2076_v18_: _dafny.Map
                    d_2076_v18_ = _dafny.Map({d_2075_v17_: _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])})
                    d_2077_v19_: _dafny.Set
                    d_2077_v19_ = _dafny.Set({self.f16})
                    d_2078_v20_: _dafny.Map
                    d_2078_v20_ = _dafny.Map({self.f16: len(_dafny.Set({self.f14}))})
                    d_2079_v21_: _dafny.Array
                    nw406_ = _dafny.Array(None, 26)
                    nw406_[int(0)] = (p3) - (p3)
                    nw406_[int(1)] = self.f14
                    nw406_[int(2)] = default__.safeModuloInt(len(d_2072_v14_), p3)
                    nw406_[int(3)] = default__.fm0(p2, globalState)
                    nw406_[int(4)] = d_2063_v8_.f14
                    nw406_[int(5)] = p3
                    nw406_[int(6)] = p3
                    nw406_[int(7)] = p1
                    nw406_[int(8)] = p3
                    nw406_[int(9)] = p1
                    nw406_[int(10)] = (len(d_2074_v16_)) - ((self).fm12(p1, globalState))
                    nw406_[int(11)] = default__.safeDivisionInt(p1, p3)
                    nw406_[int(12)] = d_2063_v8_.f14
                    nw406_[int(13)] = p3
                    nw406_[int(14)] = self.f14
                    nw406_[int(15)] = (0) - (len((d_2055_v0_) + (((d_2076_v18_)[d_2075_v17_] if (d_2075_v17_) in (d_2076_v18_) else d_2055_v0_))))
                    nw406_[int(16)] = (p3) - (d_2063_v8_.f14)
                    nw406_[int(17)] = len((d_2077_v19_) - (_dafny.Set({p2, (self).f17})))
                    nw406_[int(18)] = self.f14
                    nw406_[int(19)] = len(d_2075_v17_)
                    nw406_[int(20)] = (p3) - (d_2063_v8_.f14)
                    nw406_[int(21)] = (self.f14) - (728)
                    nw406_[int(22)] = default__.safeModuloInt(p3, len(d_2055_v0_))
                    nw406_[int(23)] = ((0) - (d_2063_v8_.f14)) * (len(d_2055_v0_))
                    nw406_[int(24)] = len(d_2078_v20_)
                    nw406_[int(25)] = d_2063_v8_.f14
                    d_2079_v21_ = nw406_
                    index344_ = default__.safeIndex(490, (d_2079_v21_).length(0))
                    (d_2079_v21_)[index344_] = d_2063_v8_.f14
                    index345_ = default__.safeIndex(490, (d_2079_v21_).length(0))
                    (d_2079_v21_)[index345_] = (len(self.f18) if (self).f17 else default__.fm0(not((self).f17), globalState))
                    d_2080_v22_: _dafny.Array
                    nw407_ = _dafny.Array(_dafny.MultiSet({}), 9)
                    d_2080_v22_ = nw407_
                    d_2081_v23_: _dafny.Array
                    def lambda112_(d_2082_p2_):
                        def lambda113_(d_2083_i1_):
                            return d_2082_p2_

                        return lambda113_

                    init65_ = lambda112_(p2)
                    nw408_ = _dafny.Array(None, 27)
                    for i0_65_ in range(nw408_.length(0)):
                        nw408_[i0_65_] = init65_(i0_65_)
                    d_2081_v23_ = nw408_
                    d_2084_v24_: _dafny.Seq
                    d_2084_v24_ = _dafny.SeqWithoutIsStrInference([d_2080_v22_])
                    d_2085_v25_: _dafny.MultiSet
                    d_2085_v25_ = _dafny.MultiSet([self.f16, self.f16, p2, (self).f17, default__.fm31(self.f16, (self).f17, globalState)])
                    d_2086_v26_: C2
                    nw409_ = C2()
                    nw409_.ctor__(D3_DC3(p2), (self).f17, self.f16)
                    d_2086_v26_ = nw409_
                    d_2087_v27_: _dafny.Seq
                    d_2087_v27_ = _dafny.SeqWithoutIsStrInference([d_2086_v26_, d_2086_v26_, d_2086_v26_, d_2086_v26_, d_2086_v26_])
                    d_2088_v28_: _dafny.Map
                    d_2088_v28_ = _dafny.Map({d_2087_v27_: (d_2081_v23_ if self.f16 else d_2081_v23_)})
                    rhs387_ = (d_2084_v24_)[default__.safeIndex(p1, len(d_2084_v24_))]
                    rhs388_ = default__.safeDivisionInt(((d_2085_v25_).cardinality if False else default__.fm0(self.f16, globalState)), ((d_2075_v17_)[default__.safeIndex(d_2063_v8_.f14, len(d_2075_v17_))]) + (self.f14))
                    rhs389_ = ((d_2088_v28_)[d_2087_v27_] if (d_2087_v27_) in (d_2088_v28_) else d_2081_v23_)
                    lhs299_ = globalState
                    d_2080_v22_ = rhs387_
                    lhs299_.f2 = rhs388_
                    d_2081_v23_ = rhs389_
                    if (self).f17:
                        d_2089_v29_: _dafny.Map
                        d_2089_v29_ = _dafny.Map({True: p2})
                        d_2072_v14_ = (d_2072_v14_).set(len((default__.fm40(p2, globalState)) | (d_2089_v29_)), len(d_2075_v17_))
                        d_2090_v30_: D16
                        d_2090_v30_ = D16_DC33(self.f16, self.f16, (d_2079_v21_)[default__.safeIndex(490, (d_2079_v21_).length(0))], self.f16)
                        d_2091_v31_: _dafny.Map
                        def iife156_(_pat_let40_0):
                            def iife157_(d_2092_dt__update__tmp_h1_):
                                def iife158_(_pat_let41_0):
                                    def iife159_(d_2093_dt__update_hcf50_h0_):
                                        return D16_DC33((d_2092_dt__update__tmp_h1_).cf47, (d_2092_dt__update__tmp_h1_).cf48, (d_2092_dt__update__tmp_h1_).cf49, d_2093_dt__update_hcf50_h0_)
                                    return iife159_(_pat_let41_0)
                                return iife158_(False)
                            return iife157_(_pat_let40_0)
                        d_2091_v31_ = _dafny.Map({iife156_(d_2090_v30_): 970})
                        d_2091_v31_ = (d_2091_v31_).set(d_2090_v30_, ((0) - (p3)) + (self.f14))
                        d_2094_v32_: D4
                        d_2094_v32_ = D4_DC6(_dafny.Map({p2: p2}))
                        d_2095_v33_: C4
                        nw410_ = C4()
                        nw410_.ctor__((self).f17, (self.f16) in ((d_2094_v32_).cf10), (self).f17)
                        d_2095_v33_ = nw410_
                        index346_ = default__.safeIndex(150, (d_2081_v23_).length(0))
                        (d_2081_v23_)[index346_] = False
                        index347_ = default__.safeIndex(150, (d_2081_v23_).length(0))
                        (d_2081_v23_)[index347_] = default__.fm17(False, d_2063_v8_.f14, p3, (d_2079_v21_)[default__.safeIndex(490, (d_2079_v21_).length(0))], globalState)
                        (self).f16 = (d_2095_v33_).f21
                    elif True:
                        d_2096_v34_: _dafny.Seq
                        d_2096_v34_ = _dafny.SeqWithoutIsStrInference([d_2081_v23_, d_2081_v23_])
                        d_2097_v35_: _dafny.Seq
                        d_2097_v35_ = _dafny.SeqWithoutIsStrInference([d_2096_v34_, d_2096_v34_, d_2096_v34_])
                        d_2098_v36_: _dafny.Seq
                        d_2098_v36_ = (d_2097_v35_)[default__.safeIndex(p1, len(d_2097_v35_))]
                        d_2096_v34_ = (d_2098_v36_)
                        d_2099_v37_: C2
                        nw411_ = C2()
                        nw411_.ctor__(d_2086_v26_.f24, default__.fm43(self.f16, globalState), not(p2))
                        d_2099_v37_ = nw411_
                        d_2100_v38_: _dafny.Map
                        d_2100_v38_ = _dafny.Map({d_2077_v19_: (self).f17})
                        d_2101_v39_: _dafny.Map
                        d_2101_v39_ = _dafny.Map({self.f14: True})
                        d_2100_v38_ = (d_2100_v38_).set((_dafny.Set({True, default__.fm17(p2, p3, d_2063_v8_.f14, len(d_2055_v0_), globalState), ((d_2101_v39_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lph")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lph")))) in (d_2101_v39_) else self.f16)}) if (self).f17 else d_2077_v19_), (self).f17)
                        (globalState).f2 = len(self.f18)
                        index348_ = default__.safeIndex(490, (d_2079_v21_).length(0))
                        (d_2079_v21_)[index348_] = len((_dafny.SeqWithoutIsStrInference([p0]) if (d_2085_v25_).isdisjoint(_dafny.MultiSet([(self).f17])) else self.f18))
                    index349_ = default__.safeIndex(211, (d_2081_v23_).length(0))
                    (d_2081_v23_)[index349_] = self.f16
                    d_2102_v40_: _dafny.Seq
                    d_2102_v40_ = _dafny.SeqWithoutIsStrInference([d_2081_v23_, d_2081_v23_, d_2081_v23_, d_2081_v23_, d_2081_v23_])
                    d_2103_v41_: _dafny.Seq
                    d_2103_v41_ = d_2102_v40_
                    index350_ = default__.safeIndex(211, (d_2081_v23_).length(0))
                    rhs390_ = False
                    rhs391_ = d_2103_v41_
                    lhs300_ = d_2081_v23_
                    lhs301_ = default__.safeIndex(211, (d_2081_v23_).length(0))
                    lhs300_[lhs301_] = rhs390_
                    d_2103_v41_ = rhs391_
                    pass
            pass
        if not(self.f16):
            r0 = p3
            d_2104_v42_: _dafny.Array
            nw412_ = _dafny.Array(False, 21)
            d_2104_v42_ = nw412_
            d_2105_v43_: _dafny.Map
            d_2105_v43_ = _dafny.Map({p2: d_2104_v42_})
            d_2106_v44_: _dafny.Set
            d_2106_v44_ = _dafny.Set({((d_2105_v43_)[(self).f17] if ((self).f17) in (d_2105_v43_) else d_2104_v42_)})
            (globalState).f2 = len(d_2106_v44_)
            if not((p3) >= (-292)):
                d_2107_v45_: _dafny.Array
                nw413_ = _dafny.Array(D9.default()(), 6)
                d_2107_v45_ = nw413_
                d_2107_v45_ = d_2107_v45_
                d_2108_v47_: _dafny.Array
                def lambda114_(d_2109_v0_):
                    def lambda115_(d_2110_i2_):
                        def iife160_():
                            coll76_ = _dafny.Set()
                            compr_76_: _dafny.Seq
                            for compr_76_ in (_dafny.SeqWithoutIsStrInference([d_2109_v0_])).Elements:
                                d_2111_v46_: _dafny.Seq = compr_76_
                                if (d_2111_v46_) in (_dafny.SeqWithoutIsStrInference([d_2109_v0_])):
                                    coll76_ = coll76_.union(_dafny.Set([d_2111_v46_]))
                            return _dafny.Set(coll76_)
                        return iife160_()
                        

                    return lambda115_

                init66_ = lambda114_(d_2055_v0_)
                nw414_ = _dafny.Array(None, 28)
                for i0_66_ in range(nw414_.length(0)):
                    nw414_[i0_66_] = init66_(i0_66_)
                d_2108_v47_ = nw414_
                d_2112_v48_: _dafny.Set
                d_2112_v48_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([not(p2)]), d_2055_v0_})
                index351_ = default__.safeIndex(538, (d_2108_v47_).length(0))
                (d_2108_v47_)[index351_] = (d_2112_v48_) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f17]), d_2055_v0_, d_2055_v0_}))
                index352_ = default__.safeIndex(538, (d_2108_v47_).length(0))
                def iife161_():
                    coll77_ = _dafny.Set()
                    compr_77_: _dafny.Seq
                    for compr_77_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p2]), d_2055_v0_])).Elements:
                        d_2113_v49_: _dafny.Seq = compr_77_
                        if (d_2113_v49_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p2]), d_2055_v0_])):
                            coll77_ = coll77_.union(_dafny.Set([d_2113_v49_]))
                    return _dafny.Set(coll77_)
                (d_2108_v47_)[index352_] = (d_2112_v48_ if (d_2063_v8_.f14) != (d_2063_v8_.f14) else iife161_()
                )
                (self).f16 = False
                d_2114_v50_: _dafny.Array
                nw415_ = _dafny.Array(int(0), 17)
                d_2114_v50_ = nw415_
                index353_ = default__.safeIndex(889, (d_2114_v50_).length(0))
                (d_2114_v50_)[index353_] = (d_2063_v8_.f14) + (d_2063_v8_.f14)
                d_2115_v51_: _dafny.Seq
                d_2115_v51_ = _dafny.SeqWithoutIsStrInference([d_2114_v50_, d_2114_v50_])
                index354_ = default__.safeIndex(889, (d_2114_v50_).length(0))
                rhs392_ = d_2063_v8_.f14
                rhs393_ = (self.f18) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lj"))) + (self.f18))
                rhs394_ = (d_2115_v51_)[default__.safeIndex(d_2063_v8_.f14, len(d_2115_v51_))]
                rhs395_ = (self).f17
                rhs396_ = (0) - ((self.f14) - ((0) - ((((self).fm15(globalState)).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0 for d_2116_i3_ in range(default__.abs(495))]), self.f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibuvhh"))]))).cardinality)))
                lhs302_ = d_2063_v8_
                lhs303_ = self
                lhs304_ = self
                lhs305_ = d_2114_v50_
                lhs306_ = default__.safeIndex(889, (d_2114_v50_).length(0))
                lhs302_.f14 = rhs392_
                lhs303_.f18 = rhs393_
                d_2114_v50_ = rhs394_
                lhs304_.f16 = rhs395_
                lhs305_[lhs306_] = rhs396_
                d_2117_v52_: D19
                d_2117_v52_ = D19_DC40(d_2063_v8_.f14, p2, (self.f14) - (d_2063_v8_.f14))
                d_2117_v52_ = d_2117_v52_
            elif True:
                (self).f16 = not(True)
                (globalState).f2 = d_2063_v8_.f14
                d_2118_v53_: D16
                d_2118_v53_ = D16_DC34(False, d_2063_v8_, p3)
                d_2119_v54_: D16
                d_2119_v54_ = D16_DC35(d_2118_v53_)
                d_2120_v55_: D16
                d_2120_v55_ = D16_DC35(d_2118_v53_)
                d_2121_v56_: D16
                d_2121_v56_ = D16_DC35(d_2118_v53_)
                d_2122_v57_: _dafny.Seq
                d_2122_v57_ = _dafny.SeqWithoutIsStrInference([d_2121_v56_, d_2121_v56_])
                d_2123_v58_: T1
                nw416_ = C6()
                nw416_.ctor__(d_2063_v8_.f14, (d_2063_v8_.f14) <= (p1), self.f16, p2, len(d_2122_v57_))
                d_2123_v58_ = nw416_
                d_2123_v58_ = d_2123_v58_
                (d_2123_v58_).f16 = (self).fm14(globalState)
                d_2055_v0_ = d_2055_v0_
            d_2124_v59_: _dafny.Map
            d_2124_v59_ = _dafny.Map({p1: d_2104_v42_})
            d_2125_v60_: _dafny.Array
            nw417_ = _dafny.Array(_dafny.Seq({}), 17)
            d_2125_v60_ = nw417_
            index355_ = default__.safeIndex(509, (d_2125_v60_).length(0))
            (d_2125_v60_)[index355_] = (d_2055_v0_) + (d_2055_v0_)
            index356_ = default__.safeIndex(509, (d_2125_v60_).length(0))
            rhs397_ = _dafny.Map({default__.safeDivisionInt(d_2063_v8_.f14, default__.fm0(self.f16, globalState)): d_2104_v42_})
            rhs398_ = d_2055_v0_
            lhs307_ = d_2125_v60_
            lhs308_ = default__.safeIndex(509, (d_2125_v60_).length(0))
            d_2124_v59_ = rhs397_
            lhs307_[lhs308_] = rhs398_
            d_2126_v61_: D19
            d_2126_v61_ = D19_DC40(d_2063_v8_.f14, self.f16, d_2063_v8_.f14)
            source33_ = d_2126_v61_
            if source33_.is_DC40:
                d_2127___mcc_h0_ = source33_.cf61
                d_2128___mcc_h1_ = source33_.cf62
                d_2129___mcc_h2_ = source33_.cf63
                d_2130_cf63_ = d_2129___mcc_h2_
                d_2131_cf62_ = d_2128___mcc_h1_
                d_2132_cf61_ = d_2127___mcc_h0_
                r0 = p1
                d_2133_v62_: _dafny.Array
                nw418_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_2133_v62_ = nw418_
                index357_ = default__.safeIndex(987, (d_2133_v62_).length(0))
                (d_2133_v62_)[index357_] = d_2104_v42_
                index358_ = default__.safeIndex(987, (d_2133_v62_).length(0))
                (d_2133_v62_)[index358_] = d_2104_v42_
                d_2134_v63_: _dafny.Array
                def lambda116_(d_2135_i4_):
                    return _dafny.Set({self.f16, not((self).f17)})

                init67_ = lambda116_
                nw419_ = _dafny.Array(None, 20)
                for i0_67_ in range(nw419_.length(0)):
                    nw419_[i0_67_] = init67_(i0_67_)
                d_2134_v63_ = nw419_
                d_2136_v64_: C0
                nw420_ = C0()
                nw420_.ctor__(d_2134_v63_, len((self.f18) + (self.f18)))
                d_2136_v64_ = nw420_
                d_2137_v65_: _dafny.Map
                d_2137_v65_ = _dafny.Map({not (self.f16) or (self.f16): ((d_2125_v60_)[default__.safeIndex(509, (d_2125_v60_).length(0))] if self.f16 else d_2055_v0_)})
                d_2137_v65_ = (d_2137_v65_).set(not ((self).f17) or (self.f16), ((d_2055_v0_)) + (d_2055_v0_))
            elif True:
                d_2138___mcc_h3_ = source33_.cf60
                d_2139_cf60_ = d_2138___mcc_h3_
                d_2140_v66_: _dafny.Map
                d_2140_v66_ = _dafny.Map({self.f14: (self).f17})
                d_2141_v67_: _dafny.MultiSet
                d_2141_v67_ = _dafny.MultiSet([self.f16])
                d_2142_v68_: C1
                nw421_ = C1()
                nw421_.ctor__(((d_2140_v66_)[len(_dafny.Set({p1, (d_2141_v67_).cardinality, p1}))] if (len(_dafny.Set({p1, (d_2141_v67_).cardinality, p1}))) in (d_2140_v66_) else False), (self).fm12(d_2063_v8_.f14, globalState))
                d_2142_v68_ = nw421_
                d_2143_v69_: _dafny.Map
                d_2143_v69_ = _dafny.Map({self.f16: (d_2141_v67_).set((self).f17, default__.abs((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])))))})
                d_2144_v70_: _dafny.Set
                d_2144_v70_ = _dafny.Set({d_2142_v68_.f25})
                d_2145_v71_: _dafny.Map
                d_2145_v71_ = _dafny.Map({p2: len(d_2144_v70_)})
                d_2146_v72_: _dafny.Seq
                d_2146_v72_ = _dafny.SeqWithoutIsStrInference([p1])
                d_2143_v69_ = (d_2143_v69_).set((d_2142_v68_.f25 if default__.fm17(self.f16, len(d_2145_v71_), len((d_2146_v72_)), self.f14, globalState) else (self).f17), _dafny.MultiSet([d_2142_v68_.f25, not((self).f17), d_2142_v68_.f25]))
                d_2147_v73_: _dafny.Array
                def lambda117_(d_2148_i5_):
                    return self.f18

                init68_ = lambda117_
                nw422_ = _dafny.Array(None, 11)
                for i0_68_ in range(nw422_.length(0)):
                    nw422_[i0_68_] = init68_(i0_68_)
                d_2147_v73_ = nw422_
                d_2149_v74_: _dafny.Array
                nw423_ = _dafny.Array(int(0), 10)
                d_2149_v74_ = nw423_
                index359_ = default__.safeIndex(684, (d_2149_v74_).length(0))
                (d_2149_v74_)[index359_] = d_2063_v8_.f14
                pat_let_tv33_ = d_2055_v0_
                pat_let_tv34_ = d_2057_v2_
                d_2150_v75_: _dafny.MultiSet
                def iife162_(_pat_let42_0):
                    def iife163_(d_2151_dt__update__tmp_h2_):
                        def iife164_(_pat_let43_0):
                            def iife165_(d_2152_dt__update_hcf34_h0_):
                                def iife166_(_pat_let44_0):
                                    def iife167_(d_2153_dt__update_hcf35_h0_):
                                        return D12_DC24((d_2151_dt__update__tmp_h2_).cf33, d_2152_dt__update_hcf34_h0_, d_2153_dt__update_hcf35_h0_, (d_2151_dt__update__tmp_h2_).cf36, (d_2151_dt__update__tmp_h2_).cf37)
                                    return iife167_(_pat_let44_0)
                                return iife166_(pat_let_tv34_)
                            return iife165_(_pat_let43_0)
                        return iife164_(pat_let_tv33_)
                    return iife163_(_pat_let42_0)
                d_2150_v75_ = _dafny.MultiSet([d_2060_v5_, d_2060_v5_, d_2060_v5_, iife162_(d_2060_v5_)])
                d_2154_v76_: _dafny.Map
                d_2154_v76_ = _dafny.Map({d_2063_v8_.f14: d_2141_v67_})
                index360_ = default__.safeIndex(684, (d_2149_v74_).length(0))
                rhs399_ = d_2147_v73_
                rhs400_ = default__.safeDivisionInt(default__.safeModuloInt(d_2063_v8_.f14, len(d_2154_v76_)), d_2063_v8_.f14)
                rhs401_ = (d_2145_v71_) | (d_2145_v71_)
                rhs402_ = _dafny.MultiSet([D12_DC24(False, _dafny.SeqWithoutIsStrInference([self.f16]), d_2057_v2_, self.f16, d_2059_v4_), d_2060_v5_, d_2060_v5_])
                lhs309_ = d_2149_v74_
                lhs310_ = default__.safeIndex(684, (d_2149_v74_).length(0))
                d_2147_v73_ = rhs399_
                lhs309_[lhs310_] = rhs400_
                d_2145_v71_ = rhs401_
                d_2150_v75_ = rhs402_
                d_2155_v77_: _dafny.Array
                nw424_ = _dafny.Array(_dafny.Seq({}), 6)
                d_2155_v77_ = nw424_
                index361_ = default__.safeIndex(684, (d_2149_v74_).length(0))
                rhs403_ = (self.f18) + ((self.f18) + (self.f18))
                rhs404_ = p1
                rhs405_ = d_2155_v77_
                lhs311_ = self
                lhs312_ = d_2149_v74_
                lhs313_ = default__.safeIndex(684, (d_2149_v74_).length(0))
                lhs311_.f18 = rhs403_
                lhs312_[lhs313_] = rhs404_
                d_2155_v77_ = rhs405_
        elif True:
            rhs406_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsv"))
            rhs407_ = d_2063_v8_.f14
            rhs408_ = d_2063_v8_.f14
            lhs314_ = self
            lhs315_ = self
            lhs316_ = globalState
            lhs314_.f18 = rhs406_
            lhs315_.f14 = rhs407_
            lhs316_.f2 = rhs408_
            d_2156_v78_: _dafny.MultiSet
            d_2156_v78_ = _dafny.MultiSet([d_2063_v8_.f14])
            d_2157_v79_: D19
            d_2157_v79_ = D19_DC39(d_2156_v78_)
            d_2158_v80_: _dafny.Set
            d_2158_v80_ = _dafny.Set({True, not(self.f16)})
            rhs409_ = default__.fm31((True) in (d_2158_v80_), ((self).f17) == ((self).f17), globalState)
            rhs410_ = d_2060_v5_
            rhs411_ = default__.fm63(globalState)
            lhs317_ = self
            lhs317_.f16 = rhs409_
            d_2060_v5_ = rhs410_
            d_2157_v79_ = rhs411_
            (globalState).f2 = d_2063_v8_.f14
            d_2159_v81_: _dafny.Array
            nw425_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_2159_v81_ = nw425_
            index362_ = default__.safeIndex(912, (d_2159_v81_).length(0))
            (d_2159_v81_)[index362_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxckdbjch"))
            index363_ = default__.safeIndex(912, (d_2159_v81_).length(0))
            (d_2159_v81_)[index363_] = self.f18
            if not ((_dafny.SeqWithoutIsStrInference([p0 for d_2160_i6_ in range(default__.abs(377))])) == ((d_2159_v81_)[default__.safeIndex(912, (d_2159_v81_).length(0))])) or (not(not(p2))):
                index364_ = default__.safeIndex(912, (d_2159_v81_).length(0))
                (d_2159_v81_)[index364_] = (d_2159_v81_)[default__.safeIndex(912, (d_2159_v81_).length(0))]
                d_2161_v82_: _dafny.Array
                nw426_ = _dafny.Array(_dafny.MultiSet({}), 20)
                d_2161_v82_ = nw426_
                d_2162_v83_: _dafny.Map
                d_2162_v83_ = _dafny.Map({self.f16: d_2156_v78_})
                index365_ = default__.safeIndex(203, (d_2161_v82_).length(0))
                (d_2161_v82_)[index365_] = ((d_2162_v83_)[self.f16] if (self.f16) in (d_2162_v83_) else d_2156_v78_)
                index366_ = default__.safeIndex(203, (d_2161_v82_).length(0))
                (d_2161_v82_)[index366_] = d_2156_v78_
                (self).f16 = p2
                r0 = d_2063_v8_.f14
                d_2163_v84_: _dafny.Set
                d_2163_v84_ = _dafny.Set({d_2063_v8_.f14, d_2063_v8_.f14})
                d_2164_v85_: _dafny.Seq
                d_2164_v85_ = _dafny.SeqWithoutIsStrInference([d_2163_v84_, d_2163_v84_, d_2163_v84_])
                d_2165_v86_: _dafny.Seq
                d_2165_v86_ = _dafny.SeqWithoutIsStrInference([p3])
                d_2166_v87_: _dafny.Seq
                d_2166_v87_ = _dafny.SeqWithoutIsStrInference([default__.fm0(default__.fm17(p2, len(_dafny.SeqWithoutIsStrInference([p0])), d_2063_v8_.f14, len(d_2165_v86_), globalState), globalState)])
                d_2167_v88_: _dafny.Map
                d_2167_v88_ = _dafny.Map({True: d_2063_v8_.f14})
                d_2168_v89_: _dafny.Map
                d_2168_v89_ = _dafny.Map({d_2063_v8_.f14: d_2063_v8_.f14})
                d_2169_v90_: _dafny.MultiSet
                d_2169_v90_ = _dafny.MultiSet([not(self.f16), p2])
                d_2170_v92_: _dafny.Array
                nw427_ = _dafny.Array(None, 25)
                nw427_[int(0)] = _dafny.Set({len(_dafny.Map({self.f16: (d_2161_v82_)[default__.safeIndex(203, (d_2161_v82_).length(0))]}))})
                nw427_[int(1)] = d_2163_v84_
                nw427_[int(2)] = (d_2164_v85_)[default__.safeIndex(len((d_2055_v0_).set(default__.safeIndex(len(d_2166_v87_), len(d_2055_v0_)), not(self.f16))), len(d_2164_v85_))]
                nw427_[int(3)] = default__.fm49(p3, len(_dafny.Map({True: self.f16})), (self).f17, len((d_2159_v81_)[default__.safeIndex(912, (d_2159_v81_).length(0))]), globalState)
                nw427_[int(4)] = d_2163_v84_
                nw427_[int(5)] = d_2163_v84_
                nw427_[int(6)] = default__.fm49((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p0 for d_2171_i7_ in range(default__.abs(-407))])), d_2063_v8_.f14, self.f14]))), ((d_2167_v88_)[p2] if (p2) in (d_2167_v88_) else ((d_2168_v89_)[(d_2156_v78_).cardinality] if ((d_2156_v78_).cardinality) in (d_2168_v89_) else d_2063_v8_.f14)), self.f16, d_2063_v8_.f14, globalState)
                nw427_[int(7)] = default__.fm49(d_2063_v8_.f14, d_2063_v8_.f14, False, d_2063_v8_.f14, globalState)
                nw427_[int(8)] = _dafny.Set({((d_2169_v90_)[self.f16] if (self.f16) in (d_2169_v90_) else d_2063_v8_.f14)})
                nw427_[int(9)] = d_2163_v84_
                nw427_[int(10)] = d_2163_v84_
                nw427_[int(11)] = d_2163_v84_
                nw427_[int(12)] = d_2163_v84_
                nw427_[int(13)] = d_2163_v84_
                nw427_[int(14)] = d_2163_v84_
                nw427_[int(15)] = (d_2163_v84_) - (_dafny.Set({p1, -151}))
                nw427_[int(16)] = _dafny.Set({len(d_2158_v80_)})
                nw427_[int(17)] = (_dafny.Set({d_2063_v8_.f14}) if self.f16 else d_2163_v84_)
                def iife168_():
                    coll78_ = _dafny.Set()
                    compr_78_: int
                    for compr_78_ in _dafny.IntegerRange(909, 310):
                        d_2172_v91_: int = compr_78_
                        if ((909) <= (d_2172_v91_)) and ((d_2172_v91_) < (310)):
                            coll78_ = coll78_.union(_dafny.Set([default__.safeDivisionInt(d_2172_v91_, len((d_2159_v81_)[default__.safeIndex(912, (d_2159_v81_).length(0))]))]))
                    return _dafny.Set(coll78_)
                nw427_[int(18)] = iife168_()
                
                nw427_[int(19)] = _dafny.Set({p3, 134, len(self.f18)})
                nw427_[int(20)] = (d_2163_v84_).intersection(d_2163_v84_)
                nw427_[int(21)] = d_2163_v84_
                nw427_[int(22)] = d_2163_v84_
                nw427_[int(23)] = default__.fm49(d_2063_v8_.f14, p1, p2, len(_dafny.Set({(self).f17, (self).f17, self.f16, p2})), globalState)
                nw427_[int(24)] = d_2163_v84_
                d_2170_v92_ = nw427_
                d_2173_v93_: _dafny.Map
                d_2173_v93_ = _dafny.Map({self.f14: p2})
                index367_ = default__.safeIndex(841, (d_2170_v92_).length(0))
                (d_2170_v92_)[index367_] = default__.fm49(len(d_2173_v93_), len(_dafny.SeqWithoutIsStrInference([p3])), self.f16, p3, globalState)
                index368_ = default__.safeIndex(841, (d_2170_v92_).length(0))
                (d_2170_v92_)[index368_] = (d_2163_v84_) - ((d_2163_v84_) | (d_2163_v84_))
            elif True:
                rhs412_ = ((d_2063_v8_.f14) - (p3)) < (301)
                rhs413_ = d_2055_v0_
                rhs414_ = d_2063_v8_.f14
                rhs415_ = ((p3) * (p3)) + (d_2063_v8_.f14)
                lhs318_ = self
                lhs319_ = globalState
                lhs320_ = globalState
                lhs318_.f16 = rhs412_
                d_2055_v0_ = rhs413_
                lhs319_.f2 = rhs414_
                lhs320_.f7 = rhs415_
                d_2174_v94_: _dafny.Array
                def lambda118_(d_2175_i8_):
                    return default__.safeModuloInt(d_2175_i8_, 544)

                init69_ = lambda118_
                nw428_ = _dafny.Array(None, 6)
                for i0_69_ in range(nw428_.length(0)):
                    nw428_[i0_69_] = init69_(i0_69_)
                d_2174_v94_ = nw428_
                index369_ = default__.safeIndex(142, (d_2174_v94_).length(0))
                (d_2174_v94_)[index369_] = (p3) * (d_2063_v8_.f14)
                index370_ = default__.safeIndex(142, (d_2174_v94_).length(0))
                (d_2174_v94_)[index370_] = self.f14
                d_2176_v95_: _dafny.Map
                d_2176_v95_ = _dafny.Map({p1: self.f16})
                d_2177_v96_: D12
                d_2177_v96_ = D12_DC23(p2, d_2063_v8_.f14, len(d_2176_v95_))
                d_2178_v97_: _dafny.Map
                d_2178_v97_ = _dafny.Map({len((d_2055_v0_).set(default__.safeIndex((d_2174_v94_)[default__.safeIndex(142, (d_2174_v94_).length(0))], len(d_2055_v0_)), self.f16)): d_2174_v94_})
                d_2179_v99_: _dafny.Seq
                d_2179_v99_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f16: (self).f17})])
                d_2180_v100_: D4
                d_2180_v100_ = D4_DC6(default__.fm40((self).f17, globalState))
                d_2181_v101_: _dafny.Map
                d_2181_v101_ = _dafny.Map({(self).f17: False})
                d_2182_v102_: _dafny.Map
                d_2182_v102_ = _dafny.Map({d_2180_v100_: d_2181_v101_})
                pat_let_tv35_ = d_2174_v94_
                pat_let_tv36_ = d_2174_v94_
                pat_let_tv37_ = d_2178_v97_
                d_2183_v103_: _dafny.Array
                nw429_ = _dafny.Array(None, 21)
                nw429_[int(0)] = (self).f17
                def iife169_(_pat_let45_0):
                    def iife170_(d_2184_dt__update__tmp_h3_):
                        def iife171_(_pat_let46_0):
                            def iife172_(d_2185_dt__update_hcf32_h0_):
                                def iife173_(_pat_let47_0):
                                    def iife174_(d_2186_dt__update_hcf31_h0_):
                                        return D12_DC23((d_2184_dt__update__tmp_h3_).cf30, d_2186_dt__update_hcf31_h0_, d_2185_dt__update_hcf32_h0_)
                                    return iife174_(_pat_let47_0)
                                return iife173_(len(pat_let_tv37_))
                            return iife172_(_pat_let46_0)
                        return iife171_((pat_let_tv36_)[default__.safeIndex(142, (pat_let_tv35_).length(0))])
                    return iife170_(_pat_let45_0)
                nw429_[int(1)] = not ((iife169_(d_2177_v96_)).cf30) or (self.f16)
                nw429_[int(2)] = p2
                def iife175_():
                    coll79_ = _dafny.Map()
                    compr_79_: str
                    for compr_79_ in (_dafny.Map({p0: (0) - (p1)})).keys.Elements:
                        d_2187_v98_: str = compr_79_
                        if (d_2187_v98_) in (_dafny.Map({p0: (0) - (p1)})):
                            coll79_[d_2187_v98_] = d_2063_v8_.f14
                    return _dafny.Map(coll79_)
                nw429_[int(3)] = (len(iife175_()
                )) != (default__.fm0(p2, globalState))
                nw429_[int(4)] = p2
                nw429_[int(5)] = p2
                nw429_[int(6)] = self.f16
                nw429_[int(7)] = (self).f17
                nw429_[int(8)] = not((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: (self).f17}) for d_2188_i9_ in range(default__.abs(293))])) < ((d_2179_v99_).set(default__.safeIndex(d_2063_v8_.f14, len(d_2179_v99_)), (((d_2182_v102_)[D4_DC6(d_2181_v101_)] if (D4_DC6(d_2181_v101_)) in (d_2182_v102_) else d_2181_v101_)).set(self.f16, p2))))
                nw429_[int(9)] = self.f16
                nw429_[int(10)] = (d_2156_v78_).issubset(d_2156_v78_)
                nw429_[int(11)] = (self.f16) not in ((self).fm13(-80, True, (self).f17, globalState))
                nw429_[int(12)] = False
                nw429_[int(13)] = p2
                nw429_[int(14)] = self.f16
                nw429_[int(15)] = not((self).f17)
                nw429_[int(16)] = self.f16
                nw429_[int(17)] = True
                nw429_[int(18)] = self.f16
                nw429_[int(19)] = p2
                nw429_[int(20)] = self.f16
                d_2183_v103_ = nw429_
                index371_ = default__.safeIndex(912, (d_2159_v81_).length(0))
                index372_ = default__.safeIndex(142, (d_2174_v94_).length(0))
                rhs416_ = self.f18
                rhs417_ = d_2063_v8_.f14
                rhs418_ = (0) - ((0) - (self.f14))
                rhs419_ = d_2183_v103_
                rhs420_ = ((p3) * (self.f14)) + (default__.fm0(False, globalState))
                lhs321_ = d_2159_v81_
                lhs322_ = default__.safeIndex(912, (d_2159_v81_).length(0))
                lhs323_ = d_2063_v8_
                lhs324_ = self
                lhs325_ = d_2174_v94_
                lhs326_ = default__.safeIndex(142, (d_2174_v94_).length(0))
                lhs321_[lhs322_] = rhs416_
                lhs323_.f14 = rhs417_
                lhs324_.f14 = rhs418_
                d_2183_v103_ = rhs419_
                lhs325_[lhs326_] = rhs420_
                d_2189_v104_: _dafny.Seq
                d_2189_v104_ = _dafny.SeqWithoutIsStrInference([d_2183_v103_, d_2183_v103_])
                d_2190_v105_: _dafny.Map
                d_2190_v105_ = _dafny.Map({743: (p3 if True else len(d_2189_v104_))})
                d_2190_v105_ = d_2190_v105_
                rhs421_ = (default__.fm0((self).f17, globalState)) == (p3)
                rhs422_ = False
                lhs327_ = self
                lhs328_ = self
                lhs327_.f16 = rhs421_
                lhs328_.f16 = rhs422_
        (d_2063_v8_).f14 = (d_2067_v12_).cardinality
        d_2055_v0_ = (d_2055_v0_) + (d_2055_v0_)
        r0 = (d_2063_v8_.f14) - (d_2063_v8_.f14)
        return r0

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: bool = False
        d_2191_v0_: _dafny.MultiSet
        d_2191_v0_ = _dafny.MultiSet([p0])
        d_2192_v1_: _dafny.Map
        d_2192_v1_ = _dafny.Map({p1: self.f16})
        d_2193_v2_: D15
        d_2193_v2_ = D15_DC30(p3)
        d_2194_v3_: _dafny.MultiSet
        d_2194_v3_ = _dafny.MultiSet([d_2193_v2_])
        (self).f14 = ((d_2191_v0_)[((d_2192_v1_)[p0] if (p0) in (d_2192_v1_) else self.f16)] if (((d_2192_v1_)[p0] if (p0) in (d_2192_v1_) else self.f16)) in (d_2191_v0_) else (0) - ((d_2194_v3_).cardinality))
        hi18_ = self.f14
        for d_2195_i0_ in range(default__.safeDivisionInt(p2, p2), hi18_):
            if p1:
                rhs423_ = default__.fm43((self.f16) or (not(p0)), globalState)
                rhs424_ = (self.f18) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvioansl")))
                rhs425_ = False
                lhs329_ = self
                lhs330_ = self
                lhs329_.f16 = rhs423_
                lhs330_.f18 = rhs424_
                r1 = rhs425_
                d_2196_v4_: _dafny.Set
                d_2196_v4_ = _dafny.Set({p2})
                d_2197_v5_: _dafny.Array
                nw430_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_2197_v5_ = nw430_
                d_2198_v6_: _dafny.Map
                d_2198_v6_ = _dafny.Map({d_2196_v4_: (d_2197_v5_ if (self).f17 else d_2197_v5_)})
                d_2198_v6_ = (d_2198_v6_).set(d_2196_v4_, d_2197_v5_)
                (self).f14 = self.f14
                (globalState).f2 = d_2195_i0_
                d_2199_v7_: _dafny.Seq
                d_2199_v7_ = _dafny.SeqWithoutIsStrInference([(self).f17, p0])
                r3 = not((self.f16) not in (d_2199_v7_))
            elif True:
                d_2200_v8_: _dafny.Seq
                d_2200_v8_ = _dafny.SeqWithoutIsStrInference([(self).fm14(globalState)])
                d_2201_v9_: _dafny.Map
                d_2201_v9_ = _dafny.Map({d_2200_v8_: d_2195_i0_})
                d_2202_v11_: _dafny.Map
                def iife176_():
                    coll80_ = _dafny.Set()
                    compr_80_: _dafny.Seq
                    for compr_80_ in (d_2201_v9_).keys.Elements:
                        d_2203_v10_: _dafny.Seq = compr_80_
                        if (d_2203_v10_) in (d_2201_v9_):
                            coll80_ = coll80_.union(_dafny.Set([d_2203_v10_]))
                    return _dafny.Set(coll80_)
                d_2202_v11_ = _dafny.Map({iife176_()
                : self.f14})
                d_2202_v11_ = (d_2202_v11_).set(_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_2200_v8_)[default__.safeIndex(d_2195_i0_, len(d_2200_v8_))]])}), d_2195_i0_)
                d_2204_v12_: str
                d_2204_v12_ = _dafny.CodePoint('p')
                d_2204_v12_ = p3
                d_2205_v13_: _dafny.Map
                d_2205_v13_ = _dafny.Map({len(d_2192_v1_): p0})
                d_2205_v13_ = (d_2205_v13_).set(self.f14, True)
                d_2206_v14_: _dafny.Array
                nw431_ = _dafny.Array(_dafny.Seq({}), 5)
                d_2206_v14_ = nw431_
                nw432_ = _dafny.Array(_dafny.Seq({}), 24)
                d_2206_v14_ = nw432_
                (globalState).f7 = (d_2195_i0_) - (default__.safeModuloInt(p2, d_2195_i0_))
            d_2207_v15_: _dafny.Array
            nw433_ = _dafny.Array(False, 20)
            d_2207_v15_ = nw433_
            d_2208_v16_: _dafny.MultiSet
            d_2208_v16_ = _dafny.MultiSet([d_2207_v15_])
            if not(not (True) or ((d_2208_v16_).ispropersubset(d_2208_v16_))):
                d_2209_v17_: _dafny.Map
                d_2209_v17_ = _dafny.Map({p2: (0) - (self.f14)})
                d_2210_v18_: _dafny.Set
                d_2210_v18_ = _dafny.Set({d_2209_v17_, d_2209_v17_, (_dafny.Map({p2: d_2195_i0_})) | (d_2209_v17_), _dafny.Map({(0) - (self.f14): (0) - (d_2195_i0_)})})
                d_2210_v18_ = d_2210_v18_
                d_2211_v19_: D11
                d_2211_v19_ = D11_DC19(d_2209_v17_)
                d_2212_v20_: D11
                d_2212_v20_ = D11_DC21(d_2211_v19_)
                d_2213_v21_: D11
                d_2213_v21_ = D11_DC21(d_2212_v20_)
                d_2214_v22_: _dafny.MultiSet
                d_2214_v22_ = _dafny.MultiSet([d_2213_v21_, d_2213_v21_, d_2213_v21_, d_2213_v21_])
                d_2215_v23_: C7
                nw434_ = C7()
                nw434_.ctor__((self).f17, (d_2214_v22_).ispropersubset(d_2214_v22_), self.f14)
                d_2215_v23_ = nw434_
                d_2216_v24_: _dafny.Array
                def lambda119_(d_2217_i0_):
                    def lambda120_(d_2218_i1_):
                        return (d_2218_i1_) + (d_2217_i0_)

                    return lambda120_

                init70_ = lambda119_(d_2195_i0_)
                nw435_ = _dafny.Array(None, 10)
                for i0_70_ in range(nw435_.length(0)):
                    nw435_[i0_70_] = init70_(i0_70_)
                d_2216_v24_ = nw435_
                index373_ = default__.safeIndex(97, (d_2216_v24_).length(0))
                (d_2216_v24_)[index373_] = default__.safeDivisionInt(self.f14, len(self.f18))
                d_2219_v25_: _dafny.Seq
                d_2219_v25_ = _dafny.SeqWithoutIsStrInference([p2, p2, d_2195_i0_, (self).fm12(d_2195_i0_, globalState), d_2195_i0_])
                d_2220_v26_: T0
                nw436_ = C7()
                nw436_.ctor__(p0, (self).f17, len(_dafny.SeqWithoutIsStrInference([(self).f17])))
                d_2220_v26_ = nw436_
                d_2221_v27_: D20
                d_2221_v27_ = D20_DC42(len(_dafny.SeqWithoutIsStrInference([p1])), False, d_2220_v26_)
                d_2222_v28_: _dafny.Map
                d_2222_v28_ = _dafny.Map({len(d_2219_v25_): default__.fm21(globalState)})
                d_2223_v29_: _dafny.Seq
                d_2223_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(default__.fm64(66, d_2195_i0_, default__.fm21(globalState), 852, globalState)): p0}), (d_2222_v28_).set(d_2195_i0_, (self).f17)])
                index374_ = default__.safeIndex(97, (d_2216_v24_).length(0))
                rhs426_ = (self).fm12((d_2219_v25_)[default__.safeIndex(self.f14, len(d_2219_v25_))], globalState)
                rhs427_ = ((self.f14) + (p2)) >= ((d_2195_i0_) * (len(_dafny.SeqWithoutIsStrInference([p3 for d_2224_i2_ in range(default__.abs(631))]))))
                rhs428_ = (default__.safeModuloInt(d_2195_i0_, d_2195_i0_) if p0 else ((d_2221_v27_).cf65) + (len(d_2223_v29_)))
                lhs331_ = globalState
                lhs332_ = d_2216_v24_
                lhs333_ = default__.safeIndex(97, (d_2216_v24_).length(0))
                lhs331_.f7 = rhs426_
                r3 = rhs427_
                lhs332_[lhs333_] = rhs428_
                d_2225_v30_: _dafny.Map
                d_2225_v30_ = _dafny.Map({p3: D20_DC42(-219, p1, d_2220_v26_)})
                d_2226_v31_: D23
                d_2226_v31_ = D23_DC47(d_2225_v30_)
                pat_let_tv38_ = d_2220_v26_
                d_2227_v32_: _dafny.Array
                nw437_ = _dafny.Array(None, 17)
                nw437_[int(0)] = (d_2225_v30_) | (d_2225_v30_)
                nw437_[int(1)] = (d_2225_v30_) | (d_2225_v30_)
                nw437_[int(2)] = d_2225_v30_
                nw437_[int(3)] = (_dafny.Map({_dafny.CodePoint('h'): D20_DC42(d_2220_v26_.f14, (self).f17, d_2220_v26_)})).set(p3, d_2221_v27_)
                nw437_[int(4)] = (d_2225_v30_).set((self.f18)[default__.safeIndex(self.f14, len(self.f18))], d_2221_v27_)
                def iife177_(_pat_let48_0):
                    def iife178_(d_2228_dt__update__tmp_h0_):
                        def iife179_(_pat_let49_0):
                            def iife180_(d_2229_dt__update_hcf67_h0_):
                                return D20_DC42((d_2228_dt__update__tmp_h0_).cf65, (d_2228_dt__update__tmp_h0_).cf66, d_2229_dt__update_hcf67_h0_)
                            return iife180_(_pat_let49_0)
                        return iife179_(pat_let_tv38_)
                    return iife178_(_pat_let48_0)
                nw437_[int(5)] = _dafny.Map({p3: iife177_(d_2221_v27_)})
                nw437_[int(6)] = _dafny.Map({p3: d_2221_v27_})
                nw437_[int(7)] = (d_2225_v30_) | ((d_2226_v31_).cf73)
                nw437_[int(8)] = _dafny.Map({p3: d_2221_v27_})
                nw437_[int(9)] = (d_2225_v30_) | (d_2225_v30_)
                nw437_[int(10)] = d_2225_v30_
                nw437_[int(11)] = _dafny.Map({p3: d_2221_v27_})
                nw437_[int(12)] = d_2225_v30_
                nw437_[int(13)] = d_2225_v30_
                nw437_[int(14)] = d_2225_v30_
                nw437_[int(15)] = (d_2225_v30_) | (d_2225_v30_)
                nw437_[int(16)] = d_2225_v30_
                d_2227_v32_ = nw437_
                index375_ = default__.safeIndex(538, (d_2227_v32_).length(0))
                (d_2227_v32_)[index375_] = d_2225_v30_
                index376_ = default__.safeIndex(538, (d_2227_v32_).length(0))
                rhs429_ = ((self).f17) == (p0)
                rhs430_ = d_2225_v30_
                rhs431_ = (default__.fm0((self).f17, globalState)) + (d_2195_i0_)
                lhs334_ = d_2227_v32_
                lhs335_ = default__.safeIndex(538, (d_2227_v32_).length(0))
                lhs336_ = globalState
                r3 = rhs429_
                lhs334_[lhs335_] = rhs430_
                lhs336_.f2 = rhs431_
                index377_ = default__.safeIndex(222, (d_2207_v15_).length(0))
                (d_2207_v15_)[index377_] = p0
                d_2230_v33_: _dafny.MultiSet
                d_2230_v33_ = _dafny.MultiSet([_dafny.CodePoint('b')])
                index378_ = default__.safeIndex(222, (d_2207_v15_).length(0))
                (d_2207_v15_)[index378_] = not((d_2230_v33_).issubset(d_2230_v33_))
            elif True:
                (self).f16 = p1
                (globalState).f2 = d_2195_i0_
                d_2231_v34_: _dafny.Map
                d_2231_v34_ = _dafny.Map({len(self.f18): (self).f17})
                d_2232_v35_: _dafny.Seq
                d_2232_v35_ = _dafny.SeqWithoutIsStrInference([d_2231_v34_])
                d_2232_v35_ = (d_2232_v35_).set(default__.safeIndex(p2, len(d_2232_v35_)), d_2231_v34_)
                (self).f14 = d_2195_i0_
                (globalState).f7 = p2
            (globalState).f2 = p2
            d_2233_v36_: _dafny.Map
            d_2233_v36_ = _dafny.Map({p0: default__.fm0(self.f16, globalState)})
            d_2234_v37_: _dafny.MultiSet
            d_2234_v37_ = _dafny.MultiSet([((d_2233_v36_)[self.f16] if (self.f16) in (d_2233_v36_) else (0) - (p2)), (d_2195_i0_) * (p2)])
            (self).f14 = ((d_2234_v37_)[(0) - (p2)] if ((0) - (p2)) in (d_2234_v37_) else d_2195_i0_)
        (globalState).f2 = self.f14
        r3 = False
        d_2235_i3_: int
        d_2235_i3_ = 0
        with _dafny.label("13"):
            while ((p2) * (default__.fm0(False, globalState))) <= (self.f14):
                with _dafny.c_label("13"):
                    if (d_2235_i3_) >= (100):
                        raise _dafny.Break("13")
                    d_2235_i3_ = (d_2235_i3_) + (1)
                    d_2236_v38_: _dafny.MultiSet
                    d_2236_v38_ = _dafny.MultiSet([p2, p2])
                    d_2237_v39_: _dafny.Set
                    d_2237_v39_ = _dafny.Set({self.f14, self.f14, self.f14})
                    d_2238_v40_: _dafny.Array
                    nw438_ = _dafny.Array(None, 20)
                    nw438_[int(0)] = p0
                    nw438_[int(1)] = p1
                    nw438_[int(2)] = (self.f14) < (p2)
                    nw438_[int(3)] = self.f16
                    nw438_[int(4)] = (d_2236_v38_) == (_dafny.MultiSet([self.f14, 799]))
                    nw438_[int(5)] = False
                    nw438_[int(6)] = (self).f17
                    nw438_[int(7)] = False
                    nw438_[int(8)] = False
                    nw438_[int(9)] = (self).f17
                    nw438_[int(10)] = self.f16
                    nw438_[int(11)] = p1
                    nw438_[int(12)] = p0
                    nw438_[int(13)] = p0
                    nw438_[int(14)] = (len(d_2237_v39_)) <= (p2)
                    nw438_[int(15)] = self.f16
                    nw438_[int(16)] = p1
                    nw438_[int(17)] = default__.fm21(globalState)
                    nw438_[int(18)] = p1
                    nw438_[int(19)] = not((False) == ((self).f17))
                    d_2238_v40_ = nw438_
                    index379_ = default__.safeIndex(576, (d_2238_v40_).length(0))
                    (d_2238_v40_)[index379_] = not (p1) or (p1)
                    index380_ = default__.safeIndex(576, (d_2238_v40_).length(0))
                    (d_2238_v40_)[index380_] = self.f16
                    d_2239_v41_: _dafny.Array
                    nw439_ = _dafny.Array(D9.default()(), 13)
                    d_2239_v41_ = nw439_
                    d_2240_v42_: _dafny.Array
                    def lambda121_(d_2241_p3_):
                        def lambda122_(d_2242_i4_):
                            return _dafny.SeqWithoutIsStrInference([d_2241_p3_ for d_2243_i5_ in range(default__.abs(-351))])

                        return lambda122_

                    init71_ = lambda121_(p3)
                    nw440_ = _dafny.Array(None, 13)
                    for i0_71_ in range(nw440_.length(0)):
                        nw440_[i0_71_] = init71_(i0_71_)
                    d_2240_v42_ = nw440_
                    index381_ = default__.safeIndex(87, (d_2240_v42_).length(0))
                    (d_2240_v42_)[index381_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xugnav"))
                    index382_ = default__.safeIndex(87, (d_2240_v42_).length(0))
                    rhs432_ = d_2239_v41_
                    rhs433_ = default__.safeModuloInt(self.f14, (p2 if (d_2238_v40_)[default__.safeIndex(576, (d_2238_v40_).length(0))] else (0) - (p2)))
                    rhs434_ = self.f18
                    lhs337_ = globalState
                    lhs338_ = d_2240_v42_
                    lhs339_ = default__.safeIndex(87, (d_2240_v42_).length(0))
                    d_2239_v41_ = rhs432_
                    lhs337_.f2 = rhs433_
                    lhs338_[lhs339_] = rhs434_
                    (self).f16 = (self).f17
                    r0 = ((d_2191_v0_).cardinality if self.f16 else 338)
                    pass
            pass
        hi19_ = self.f14
        for d_2244_i6_ in range(p2, hi19_):
            d_2245_v43_: C1
            nw441_ = C1()
            nw441_.ctor__(not(p0), d_2244_i6_)
            d_2245_v43_ = nw441_
            d_2246_v44_: _dafny.Seq
            d_2246_v44_ = _dafny.SeqWithoutIsStrInference([p0, d_2245_v43_.f25, self.f16])
            d_2247_v45_: C3
            nw442_ = C3()
            nw442_.ctor__(self.f14, (default__.fm42(d_2245_v43_.f25, self.f14, len(_dafny.SeqWithoutIsStrInference([not(d_2245_v43_.f25)])), False, globalState)) + ((self.f18).set(default__.safeIndex(len(d_2246_v44_), len(self.f18)), p3)), not((p0 if self.f16 else True)), self.f16)
            d_2247_v45_ = nw442_
            d_2248_v46_: _dafny.Map
            d_2248_v46_ = _dafny.Map({self.f14: p2})
            d_2249_v47_: _dafny.Map
            d_2249_v47_ = _dafny.Map({p1: len(d_2248_v46_)})
            d_2250_v48_: _dafny.MultiSet
            d_2250_v48_ = _dafny.MultiSet([d_2249_v47_])
            d_2251_v49_: _dafny.Map
            d_2251_v49_ = _dafny.Map({d_2244_i6_: (d_2250_v48_) | (d_2250_v48_)})
            d_2251_v49_ = (d_2251_v49_).set(p2, d_2250_v48_)
            source34_ = default__.fm65((d_2247_v45_).f22, globalState)
            if source34_.is_DC4:
                d_2252___mcc_h0_ = source34_.cf4
                d_2253___mcc_h1_ = source34_.cf5
                d_2254___mcc_h2_ = source34_.cf6
                d_2255___mcc_h3_ = source34_.cf7
                d_2256___mcc_h4_ = source34_.cf8
                d_2257_cf8_ = d_2256___mcc_h4_
                d_2258_cf7_ = d_2255___mcc_h3_
                d_2259_cf6_ = d_2254___mcc_h2_
                d_2260_cf5_ = d_2253___mcc_h1_
                d_2261_cf4_ = d_2252___mcc_h0_
                r0 = (d_2247_v45_).f22
                d_2262_v50_: _dafny.Array
                nw443_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_2262_v50_ = nw443_
                d_2263_v51_: _dafny.Array
                def lambda123_(d_2264_i7_):
                    return (d_2264_i7_) * (-822)

                init72_ = lambda123_
                nw444_ = _dafny.Array(None, 16)
                for i0_72_ in range(nw444_.length(0)):
                    nw444_[i0_72_] = init72_(i0_72_)
                d_2263_v51_ = nw444_
                index383_ = default__.safeIndex(216, (d_2262_v50_).length(0))
                (d_2262_v50_)[index383_] = d_2263_v51_
                index384_ = default__.safeIndex(216, (d_2262_v50_).length(0))
                nw445_ = _dafny.Array(int(0), 22)
                (d_2262_v50_)[index384_] = nw445_
                d_2265_v52_: _dafny.MultiSet
                d_2265_v52_ = _dafny.MultiSet([p3])
                d_2266_v53_: _dafny.Map
                d_2266_v53_ = _dafny.Map({p2: d_2257_cf8_})
                d_2267_v54_: _dafny.MultiSet
                d_2267_v54_ = _dafny.MultiSet([d_2266_v53_, d_2266_v53_])
                arr0_ = (d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))]
                index385_ = default__.safeIndex(240, ((d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))]).length(0))
                arr0_[index385_] = ((d_2265_v52_)[_dafny.CodePoint('u')] if (_dafny.CodePoint('u')) in (d_2265_v52_) else (d_2267_v54_).cardinality)
                d_2268_v55_: _dafny.Map
                d_2268_v55_ = _dafny.Map({p3: (d_2247_v45_).f22})
                arr1_ = (d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))]
                index386_ = default__.safeIndex(240, ((d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))]).length(0))
                arr1_[index386_] = default__.fm0((_dafny.CodePoint('k')) in (d_2268_v55_), globalState)
                d_2269_v56_: _dafny.Seq
                d_2269_v56_ = _dafny.SeqWithoutIsStrInference([(d_2247_v45_).f22, ((d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))])[default__.safeIndex(240, ((d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))]).length(0))], (d_2247_v45_).f22])
                d_2270_v57_: _dafny.Array
                nw446_ = _dafny.Array(None, 24)
                nw446_[int(0)] = default__.fm21(globalState)
                nw446_[int(1)] = True
                nw446_[int(2)] = (self).f17
                nw446_[int(3)] = (d_2245_v43_.f25) and (p1)
                nw446_[int(4)] = default__.fm27(len(d_2266_v53_), len((d_2247_v45_).f23), d_2244_i6_, d_2248_v46_, globalState)
                nw446_[int(5)] = (((d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))])[default__.safeIndex(240, ((d_2262_v50_)[default__.safeIndex(216, (d_2262_v50_).length(0))]).length(0))]) not in (d_2269_v56_)
                nw446_[int(6)] = (p1) not in (d_2191_v0_)
                nw446_[int(7)] = d_2258_cf7_
                nw446_[int(8)] = default__.fm43(((d_2192_v1_)[(self).fm14(globalState)] if ((self).fm14(globalState)) in (d_2192_v1_) else d_2261_cf4_), globalState)
                nw446_[int(9)] = self.f16
                nw446_[int(10)] = (self).f17
                nw446_[int(11)] = (self).f17
                nw446_[int(12)] = p1
                nw446_[int(13)] = d_2257_cf8_
                nw446_[int(14)] = (d_2246_v44_)[default__.safeIndex((d_2247_v45_).f22, len(d_2246_v44_))]
                nw446_[int(15)] = d_2257_cf8_
                nw446_[int(16)] = d_2245_v43_.f25
                nw446_[int(17)] = p0
                nw446_[int(18)] = (d_2269_v56_) != (d_2269_v56_)
                nw446_[int(19)] = (self).f17
                nw446_[int(20)] = d_2245_v43_.f25
                nw446_[int(21)] = d_2258_cf7_
                nw446_[int(22)] = True
                nw446_[int(23)] = ((d_2247_v45_).f23) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awdr")))
                d_2270_v57_ = nw446_
                index387_ = default__.safeIndex(602, (d_2270_v57_).length(0))
                (d_2270_v57_)[index387_] = self.f16
                index388_ = default__.safeIndex(602, (d_2270_v57_).length(0))
                rhs435_ = d_2258_cf7_
                rhs436_ = (len((d_2247_v45_).f23)) <= ((d_2247_v45_).f22)
                lhs340_ = d_2270_v57_
                lhs341_ = default__.safeIndex(602, (d_2270_v57_).length(0))
                lhs340_[lhs341_] = rhs435_
                d_2258_cf7_ = rhs436_
            elif source34_.is_DC3:
                d_2271___mcc_h5_ = source34_.cf3
                d_2272_cf3_ = d_2271___mcc_h5_
                d_2273_v58_: _dafny.Array
                def lambda124_(d_2274_cf3_, d_2275_p1_, d_2276_p0_):
                    def lambda125_(d_2277_i8_):
                        return (d_2274_cf3_) in (_dafny.Map({d_2275_p1_: d_2276_p0_}))

                    return lambda125_

                init73_ = lambda124_(d_2272_cf3_, p1, p0)
                nw447_ = _dafny.Array(None, 23)
                for i0_73_ in range(nw447_.length(0)):
                    nw447_[i0_73_] = init73_(i0_73_)
                d_2273_v58_ = nw447_
                index389_ = default__.safeIndex(979, (d_2273_v58_).length(0))
                (d_2273_v58_)[index389_] = p1
                d_2278_v59_: _dafny.Array
                nw448_ = _dafny.Array(None, 17)
                nw448_[int(0)] = p3
                nw448_[int(1)] = p3
                nw448_[int(2)] = p3
                nw448_[int(3)] = p3
                nw448_[int(4)] = p3
                nw448_[int(5)] = p3
                nw448_[int(6)] = p3
                nw448_[int(7)] = _dafny.CodePoint('g')
                nw448_[int(8)] = p3
                nw448_[int(9)] = p3
                nw448_[int(10)] = p3
                nw448_[int(11)] = p3
                nw448_[int(12)] = p3
                nw448_[int(13)] = _dafny.CodePoint('m')
                nw448_[int(14)] = p3
                nw448_[int(15)] = p3
                nw448_[int(16)] = p3
                d_2278_v59_ = nw448_
                d_2279_v60_: _dafny.Map
                d_2279_v60_ = _dafny.Map({(d_2247_v45_).f22: d_2278_v59_})
                d_2280_v61_: _dafny.Seq
                d_2280_v61_ = _dafny.SeqWithoutIsStrInference([(len(d_2279_v60_)) * (p2), d_2244_i6_, default__.fm0(p0, globalState), (d_2247_v45_).f22, default__.fm0(self.f16, globalState)])
                index390_ = default__.safeIndex(979, (d_2273_v58_).length(0))
                rhs437_ = ((len((d_2247_v45_).f23)) - (-659)) * (len((self.f18) + ((d_2247_v45_).f23)))
                rhs438_ = not(d_2245_v43_.f25)
                rhs439_ = not((self).f17)
                rhs440_ = (d_2280_v61_)[default__.safeIndex(d_2244_i6_, len(d_2280_v61_))]
                lhs342_ = d_2245_v43_
                lhs343_ = d_2273_v58_
                lhs344_ = default__.safeIndex(979, (d_2273_v58_).length(0))
                lhs345_ = self
                r0 = rhs437_
                lhs342_.f25 = rhs438_
                lhs343_[lhs344_] = rhs439_
                lhs345_.f14 = rhs440_
                (d_2245_v43_).f25 = default__.fm43(False, globalState)
                d_2281_v62_: _dafny.Map
                d_2281_v62_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p3, p3]): (0) - (self.f14)})
                d_2281_v62_ = (d_2281_v62_).set((d_2247_v45_).f23, 271)
                d_2282_v63_: _dafny.Array
                nw449_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_2282_v63_ = nw449_
                d_2283_v64_: _dafny.Array
                nw450_ = _dafny.Array(None, 22)
                nw450_[int(0)] = d_2282_v63_
                nw450_[int(1)] = d_2282_v63_
                nw450_[int(2)] = d_2282_v63_
                nw450_[int(3)] = d_2282_v63_
                nw450_[int(4)] = d_2282_v63_
                nw450_[int(5)] = d_2282_v63_
                nw450_[int(6)] = d_2282_v63_
                nw450_[int(7)] = d_2282_v63_
                nw450_[int(8)] = d_2282_v63_
                nw450_[int(9)] = d_2282_v63_
                nw450_[int(10)] = d_2282_v63_
                nw450_[int(11)] = d_2282_v63_
                nw450_[int(12)] = (d_2282_v63_ if True else d_2282_v63_)
                nw450_[int(13)] = d_2282_v63_
                nw450_[int(14)] = d_2282_v63_
                nw450_[int(15)] = d_2282_v63_
                nw450_[int(16)] = d_2282_v63_
                nw450_[int(17)] = d_2282_v63_
                nw450_[int(18)] = d_2282_v63_
                nw450_[int(19)] = d_2282_v63_
                nw450_[int(20)] = d_2282_v63_
                nw450_[int(21)] = d_2282_v63_
                d_2283_v64_ = nw450_
                index391_ = default__.safeIndex(40, (d_2283_v64_).length(0))
                (d_2283_v64_)[index391_] = d_2282_v63_
                index392_ = default__.safeIndex(40, (d_2283_v64_).length(0))
                (d_2283_v64_)[index392_] = d_2282_v63_
            elif True:
                d_2284___mcc_h6_ = source34_.cf9
                d_2285_cf9_ = d_2284___mcc_h6_
                rhs441_ = self.f14
                rhs442_ = d_2245_v43_.f25
                lhs346_ = globalState
                lhs346_.f2 = rhs441_
                r3 = rhs442_
                d_2286_v65_: _dafny.Seq
                d_2286_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2247_v45_).f22])])
                d_2287_v66_: _dafny.Map
                d_2287_v66_ = _dafny.Map({p0: (d_2286_v65_) + (d_2286_v65_)})
                (globalState).f7 = (0) - (len(((d_2287_v66_)[d_2245_v43_.f25] if (d_2245_v43_.f25) in (d_2287_v66_) else d_2286_v65_)))
                (d_2245_v43_).f25 = ((self).f17 if p0 else (False) == (not(self.f16)))
                d_2288_v67_: _dafny.Array
                nw451_ = _dafny.Array(None, 7)
                nw451_[int(0)] = self.f14
                nw451_[int(1)] = 613
                nw451_[int(2)] = self.f14
                nw451_[int(3)] = (d_2247_v45_).f22
                nw451_[int(4)] = (d_2247_v45_).f22
                nw451_[int(5)] = p2
                nw451_[int(6)] = (d_2247_v45_).f22
                d_2288_v67_ = nw451_
                d_2289_v68_: _dafny.Seq
                d_2289_v68_ = _dafny.SeqWithoutIsStrInference([d_2288_v67_])
                d_2290_v69_: _dafny.Set
                d_2290_v69_ = _dafny.Set({p0})
                d_2291_v70_: D15
                d_2291_v70_ = D15_DC29(d_2290_v69_)
                d_2292_v71_: _dafny.Map
                d_2292_v71_ = _dafny.Map({d_2289_v68_: d_2291_v70_})
                d_2292_v71_ = (d_2292_v71_).set(d_2289_v68_, d_2291_v70_)
        r0 = p2
        d_2293_v72_: _dafny.Set
        d_2293_v72_ = _dafny.Set({self.f14})
        d_2294_v73_: _dafny.Map
        d_2294_v73_ = _dafny.Map({d_2293_v72_: self.f14})
        d_2295_v74_: _dafny.Map
        d_2295_v74_ = _dafny.Map({self.f14: p2})
        d_2296_v76_: _dafny.Set
        def iife181_():
            coll81_ = _dafny.Map()
            compr_81_: int
            for compr_81_ in _dafny.IntegerRange(117, -228):
                d_2297_v75_: int = compr_81_
                if ((117) <= (d_2297_v75_)) and ((d_2297_v75_) < (-228)):
                    coll81_[(d_2297_v75_) - (self.f14)] = self.f14
            return _dafny.Map(coll81_)
        d_2296_v76_ = _dafny.Set({d_2295_v74_, iife181_()
        , d_2295_v74_, d_2295_v74_, d_2295_v74_})
        r1 = not(not((d_2296_v76_).ispropersubset(default__.fm66(p2, d_2294_v73_, p2, p0, globalState))))
        d_2298_v77_: D13
        d_2298_v77_ = D13_DC25(_dafny.Map({self.f18: p2}))
        d_2299_v78_: _dafny.Map
        d_2299_v78_ = _dafny.Map({self.f18: p2})
        pat_let_tv39_ = d_2299_v78_
        def lambda126_(source35_):
            if source35_.is_DC26:
                return (self.f18) + (self.f18)
            elif True:
                d_2300___mcc_h7_ = source35_.cf38
                d_2301_cf38_ = d_2300___mcc_h7_
                return self.f18

        def iife182_(_pat_let50_0):
            def iife183_(d_2302_dt__update__tmp_h1_):
                def iife184_(_pat_let51_0):
                    def iife185_(d_2303_dt__update_hcf38_h0_):
                        return D13_DC25(d_2303_dt__update_hcf38_h0_)
                    return iife185_(_pat_let51_0)
                return iife184_(pat_let_tv39_)
            return iife183_(_pat_let50_0)
        r2 = lambda126_(D13_DC25((iife182_(d_2298_v77_)).cf38))
        r3 = p1
        return r0, r1, r2, r3

    def m7(self, p0, p1, globalState):
        r0: bool = False
        d_2304_v0_: _dafny.Seq
        d_2304_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2305_v1_: _dafny.Map
        d_2305_v1_ = _dafny.Map({(self).f17: _dafny.MultiSet(d_2304_v0_)})
        d_2306_v2_: _dafny.MultiSet
        d_2306_v2_ = _dafny.MultiSet([self.f16])
        d_2307_v3_: _dafny.Map
        d_2307_v3_ = _dafny.Map({565: p0})
        d_2308_v4_: D16
        d_2308_v4_ = D16_DC33(self.f16, (self).f17, ((d_2306_v2_)[self.f16] if (self.f16) in (d_2306_v2_) else len(d_2307_v3_)), (self).f17)
        d_2309_v5_: D11
        d_2309_v5_ = D11_DC19(d_2307_v3_)
        d_2310_v6_: _dafny.Seq
        d_2310_v6_ = _dafny.SeqWithoutIsStrInference([d_2309_v5_, d_2309_v5_])
        d_2311_v7_: _dafny.MultiSet
        d_2311_v7_ = _dafny.MultiSet([self.f14, len((d_2310_v6_).set(default__.safeIndex(self.f14, len(d_2310_v6_)), d_2309_v5_))])
        d_2312_i0_: int
        d_2312_i0_ = 0
        with _dafny.label("14"):
            while default__.fm27((((d_2305_v1_)[(d_2308_v4_).cf50] if ((d_2308_v4_).cf50) in (d_2305_v1_) else d_2311_v7_)).cardinality, (0) - (len(self.f18)), self.f14, d_2307_v3_, globalState):
                with _dafny.c_label("14"):
                    if (d_2312_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2312_i0_ = (d_2312_i0_) + (1)
                    d_2313_v8_: _dafny.Map
                    d_2313_v8_ = _dafny.Map({len(self.f18): (self).f17})
                    (globalState).f7 = (0) - ((len(d_2313_v8_)) * (len(self.f18)))
                    (globalState).f7 = self.f14
                    d_2314_v9_: _dafny.Array
                    def lambda127_(d_2315_i1_):
                        return (d_2315_i1_) - ((0) - ((0) - (self.f14)))

                    init74_ = lambda127_
                    nw452_ = _dafny.Array(None, 21)
                    for i0_74_ in range(nw452_.length(0)):
                        nw452_[i0_74_] = init74_(i0_74_)
                    d_2314_v9_ = nw452_
                    d_2316_v10_: _dafny.MultiSet
                    d_2316_v10_ = _dafny.MultiSet([d_2314_v9_, d_2314_v9_])
                    (globalState).f7 = len(_dafny.Map({((d_2316_v10_)[d_2314_v9_] if (d_2314_v9_) in (d_2316_v10_) else self.f14): (self.f14 if self.f16 else len(self.f18))}))
                    d_2317_v11_: C4
                    nw453_ = C4()
                    nw453_.ctor__(not (self.f16) or (self.f16), (self).f17, (d_2316_v10_) == (d_2316_v10_))
                    d_2317_v11_ = nw453_
                    pass
            pass
        if False:
            (self).f14 = self.f14
            d_2318_v12_: D21
            d_2318_v12_ = D21_DC44(self.f14, (self).f17)
            d_2319_v13_: _dafny.Seq
            d_2319_v13_ = _dafny.SeqWithoutIsStrInference([(self).f17, (d_2318_v12_).cf70, self.f16])
            d_2320_v14_: C6
            nw454_ = C6()
            nw454_.ctor__(p0, (d_2319_v13_) <= (d_2319_v13_), self.f16, (self).f17, (p0) * (self.f14))
            d_2320_v14_ = nw454_
            d_2321_v15_: C3
            nw455_ = C3()
            nw455_.ctor__(p0, self.f18, self.f16, self.f16)
            d_2321_v15_ = nw455_
            d_2322_v16_: _dafny.Map
            d_2322_v16_ = _dafny.Map({d_2321_v15_: (self).f17})
            (globalState).f7 = len((d_2322_v16_) | (d_2322_v16_))
            d_2323_v17_: _dafny.Array
            nw456_ = _dafny.Array(int(0), 11)
            d_2323_v17_ = nw456_
            index393_ = default__.safeIndex(131, (d_2323_v17_).length(0))
            (d_2323_v17_)[index393_] = (12 if (d_2320_v14_).f20 else p0)
            index394_ = default__.safeIndex(131, (d_2323_v17_).length(0))
            (d_2323_v17_)[index394_] = (p0) + (640)
            d_2324_v18_: D12
            d_2324_v18_ = D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tw")))
            source36_ = d_2324_v18_
            if source36_.is_DC23:
                d_2325___mcc_h0_ = source36_.cf30
                d_2326___mcc_h1_ = source36_.cf31
                d_2327___mcc_h2_ = source36_.cf32
                d_2328_cf32_ = d_2327___mcc_h2_
                d_2329_cf31_ = d_2326___mcc_h1_
                d_2330_cf30_ = d_2325___mcc_h0_
                d_2331_v19_: _dafny.Array
                nw457_ = _dafny.Array(False, 18)
                d_2331_v19_ = nw457_
                d_2332_v20_: _dafny.Array
                d_2332_v20_ = d_2331_v19_
                d_2331_v19_ = (d_2332_v20_)
                d_2333_v21_: _dafny.Array
                nw458_ = _dafny.Array(_dafny.Map({}), 27)
                d_2333_v21_ = nw458_
                d_2333_v21_ = d_2333_v21_
                index395_ = default__.safeIndex(277, (d_2331_v19_).length(0))
                (d_2331_v19_)[index395_] = (self).f17
                index396_ = default__.safeIndex(277, (d_2331_v19_).length(0))
                (d_2331_v19_)[index396_] = d_2330_cf30_
                d_2334_v22_: str
                d_2334_v22_ = _dafny.CodePoint('k')
                d_2334_v22_ = d_2334_v22_
            elif source36_.is_DC24:
                d_2335___mcc_h3_ = source36_.cf33
                d_2336___mcc_h4_ = source36_.cf34
                d_2337___mcc_h5_ = source36_.cf35
                d_2338___mcc_h6_ = source36_.cf36
                d_2339___mcc_h7_ = source36_.cf37
                d_2340_cf37_ = d_2339___mcc_h7_
                d_2341_cf36_ = d_2338___mcc_h6_
                d_2342_cf35_ = d_2337___mcc_h5_
                d_2343_cf34_ = d_2336___mcc_h4_
                d_2344_cf33_ = d_2335___mcc_h3_
                d_2345_v23_: _dafny.Map
                d_2345_v23_ = _dafny.Map({d_2344_cf33_: (d_2320_v14_).f20})
                d_2345_v23_ = d_2345_v23_
                d_2346_v24_: D7
                d_2346_v24_ = D7_DC12()
                d_2347_v25_: _dafny.Map
                d_2347_v25_ = _dafny.Map({(d_2320_v14_).f20: (d_2321_v15_).f22})
                d_2348_v26_: _dafny.Map
                d_2348_v26_ = _dafny.Map({self.f14: d_2347_v25_})
                d_2346_v24_ = default__.fm67((d_2348_v26_) | (d_2348_v26_), globalState)
                d_2323_v17_ = d_2323_v17_
                d_2349_v27_: _dafny.Array
                def lambda128_(d_2350_v15_):
                    def lambda129_(d_2351_i2_):
                        return (d_2350_v15_).f23

                    return lambda129_

                init75_ = lambda128_(d_2321_v15_)
                nw459_ = _dafny.Array(None, 17)
                for i0_75_ in range(nw459_.length(0)):
                    nw459_[i0_75_] = init75_(i0_75_)
                d_2349_v27_ = nw459_
                d_2352_v28_: _dafny.Array
                nw460_ = _dafny.Array(None, 18)
                nw460_[int(0)] = d_2349_v27_
                nw460_[int(1)] = d_2349_v27_
                nw460_[int(2)] = d_2349_v27_
                nw460_[int(3)] = d_2349_v27_
                nw460_[int(4)] = d_2349_v27_
                nw460_[int(5)] = d_2349_v27_
                nw460_[int(6)] = d_2349_v27_
                nw460_[int(7)] = d_2349_v27_
                nw460_[int(8)] = d_2349_v27_
                nw460_[int(9)] = d_2349_v27_
                nw460_[int(10)] = d_2349_v27_
                nw460_[int(11)] = d_2349_v27_
                nw460_[int(12)] = d_2349_v27_
                nw460_[int(13)] = d_2349_v27_
                nw460_[int(14)] = d_2349_v27_
                nw460_[int(15)] = (d_2349_v27_ if d_2341_cf36_ else d_2349_v27_)
                nw460_[int(16)] = d_2349_v27_
                nw460_[int(17)] = d_2349_v27_
                d_2352_v28_ = nw460_
                index397_ = default__.safeIndex(99, (d_2352_v28_).length(0))
                (d_2352_v28_)[index397_] = d_2349_v27_
                d_2353_v29_: _dafny.Map
                d_2353_v29_ = _dafny.Map({default__.safeModuloInt((d_2320_v14_).f19, len(d_2304_v0_)): d_2349_v27_})
                index398_ = default__.safeIndex(99, (d_2352_v28_).length(0))
                (d_2352_v28_)[index398_] = ((d_2353_v29_)[(d_2323_v17_)[default__.safeIndex(131, (d_2323_v17_).length(0))]] if ((d_2323_v17_)[default__.safeIndex(131, (d_2323_v17_).length(0))]) in (d_2353_v29_) else d_2349_v27_)
            elif True:
                d_2354___mcc_h8_ = source36_.cf29
                d_2355_cf29_ = d_2354___mcc_h8_
                index399_ = default__.safeIndex(131, (d_2323_v17_).length(0))
                (d_2323_v17_)[index399_] = (0) - (p0)
                d_2356_v30_: _dafny.Array
                nw461_ = _dafny.Array(D9.default()(), 13)
                d_2356_v30_ = nw461_
                d_2357_v31_: _dafny.Seq
                d_2357_v31_ = _dafny.SeqWithoutIsStrInference([d_2356_v30_])
                d_2358_v32_: D3
                d_2358_v32_ = D3_DC3((self).f17)
                d_2359_v33_: _dafny.Seq
                d_2359_v33_ = _dafny.SeqWithoutIsStrInference([D3_DC3((self).f17), d_2358_v32_])
                d_2360_v34_: D12
                d_2360_v34_ = D12_DC24((d_2320_v14_).f20, d_2319_v13_, d_2357_v31_, (d_2320_v14_).f20, d_2359_v33_)
                d_2361_v35_: _dafny.Array
                nw462_ = _dafny.Array(None, 18)
                nw462_[int(0)] = self.f16
                nw462_[int(1)] = (self).f17
                nw462_[int(2)] = (self).f17
                nw462_[int(3)] = (self).f17
                nw462_[int(4)] = (d_2320_v14_).f20
                nw462_[int(5)] = (d_2320_v14_).f20
                nw462_[int(6)] = (d_2320_v14_).f20
                nw462_[int(7)] = (d_2320_v14_).f20
                nw462_[int(8)] = (d_2320_v14_).f20
                nw462_[int(9)] = (self).f17
                nw462_[int(10)] = (self).f17
                nw462_[int(11)] = self.f16
                nw462_[int(12)] = (self).f17
                nw462_[int(13)] = (d_2360_v34_).cf36
                nw462_[int(14)] = (self).f17
                nw462_[int(15)] = (d_2320_v14_).f20
                nw462_[int(16)] = (d_2320_v14_).f20
                nw462_[int(17)] = self.f16
                d_2361_v35_ = nw462_
                d_2362_v36_: _dafny.Seq
                d_2362_v36_ = _dafny.SeqWithoutIsStrInference([d_2361_v35_, d_2361_v35_, d_2361_v35_, d_2361_v35_, d_2361_v35_])
                d_2363_v37_: _dafny.Map
                d_2363_v37_ = _dafny.Map({(904) - ((d_2321_v15_).f22): ((d_2362_v36_) + (d_2362_v36_)).set(default__.safeIndex((d_2321_v15_).f22, len((d_2362_v36_) + (d_2362_v36_))), d_2361_v35_)})
                d_2363_v37_ = (d_2363_v37_).set((d_2320_v14_).f19, (d_2362_v36_) + (d_2362_v36_))
                index400_ = default__.safeIndex(131, (d_2323_v17_).length(0))
                (d_2323_v17_)[index400_] = self.f14
                d_2364_v38_: _dafny.Map
                d_2364_v38_ = _dafny.Map({_dafny.CodePoint('b'): (d_2304_v0_) < (d_2304_v0_)})
                d_2365_v39_: str
                d_2365_v39_ = _dafny.CodePoint('l')
                d_2364_v38_ = (d_2364_v38_).set(d_2365_v39_, (d_2360_v34_).cf36)
        elif True:
            d_2366_v40_: str
            d_2366_v40_ = _dafny.CodePoint('y')
            d_2367_v41_: D21
            d_2367_v41_ = D21_DC43(_dafny.MultiSet([d_2366_v40_, d_2366_v40_]))
            d_2368_v42_: _dafny.Seq
            d_2368_v42_ = _dafny.SeqWithoutIsStrInference([d_2304_v0_])
            d_2369_v43_: _dafny.Seq
            d_2369_v43_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_2370_i3_ in range(default__.abs(877))])]), d_2368_v42_])
            d_2371_v44_: _dafny.Map
            d_2371_v44_ = _dafny.Map({self.f16: (d_2369_v43_)[default__.safeIndex(self.f14, len(d_2369_v43_))]})
            rhs443_ = len(self.f18)
            rhs444_ = d_2367_v41_
            rhs445_ = ((((d_2371_v44_)[True] if (True) in (d_2371_v44_) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, self.f14, self.f14])]))).set(default__.safeIndex(self.f14, len(((d_2371_v44_)[True] if (True) in (d_2371_v44_) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, self.f14, self.f14])])))), d_2304_v0_)) + (d_2368_v42_)
            rhs446_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbxjip"))) + (self.f18))
            lhs347_ = globalState
            lhs348_ = globalState
            lhs347_.f7 = rhs443_
            d_2367_v41_ = rhs444_
            d_2368_v42_ = rhs445_
            lhs348_.f2 = rhs446_
            d_2372_v45_: _dafny.Array
            def lambda130_(d_2373_v40_):
                def lambda131_(d_2374_i4_):
                    return d_2373_v40_

                return lambda131_

            init76_ = lambda130_(d_2366_v40_)
            nw463_ = _dafny.Array(None, 25)
            for i0_76_ in range(nw463_.length(0)):
                nw463_[i0_76_] = init76_(i0_76_)
            d_2372_v45_ = nw463_
            d_2375_v46_: _dafny.Map
            d_2375_v46_ = _dafny.Map({d_2372_v45_: (self).f17})
            (globalState).f2 = len(d_2375_v46_)
            d_2376_v47_: _dafny.Map
            d_2376_v47_ = _dafny.Map({self.f14: (self).f17})
            d_2377_v48_: _dafny.Map
            d_2377_v48_ = _dafny.Map({d_2376_v47_: default__.safeModuloInt(p0, self.f14)})
            d_2377_v48_ = (d_2377_v48_).set(d_2376_v47_, len(d_2307_v3_))
            if True:
                d_2378_v49_: _dafny.Seq
                d_2378_v49_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                d_2379_v50_: _dafny.Map
                d_2379_v50_ = _dafny.Map({self.f14: d_2378_v49_})
                d_2380_v51_: _dafny.Map
                d_2380_v51_ = _dafny.Map({self.f16: (self).f17})
                d_2381_v52_: _dafny.Map
                d_2381_v52_ = _dafny.Map({d_2380_v51_: d_2378_v49_})
                d_2378_v49_ = (d_2378_v49_) + ((d_2378_v49_) + (((d_2379_v50_)[self.f14] if (self.f14) in (d_2379_v50_) else ((d_2381_v52_)[_dafny.Map({self.f16: (self).f17})] if (_dafny.Map({self.f16: (self).f17})) in (d_2381_v52_) else _dafny.SeqWithoutIsStrInference([self.f16, True])))))
                d_2382_v53_: D18
                d_2382_v53_ = D18_DC38(self.f14, self.f16, -384)
                d_2383_v54_: _dafny.Seq
                d_2383_v54_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f17: (self).f17})])
                d_2384_v55_: _dafny.Array
                nw464_ = _dafny.Array(int(0), 15)
                d_2384_v55_ = nw464_
                d_2385_v56_: _dafny.Map
                d_2385_v56_ = _dafny.Map({(0) - (-583): d_2384_v55_})
                d_2386_v57_: _dafny.Array
                nw465_ = _dafny.Array(None, 20)
                nw465_[int(0)] = self.f14
                nw465_[int(1)] = self.f14
                nw465_[int(2)] = self.f14
                nw465_[int(3)] = self.f14
                nw465_[int(4)] = len(d_2380_v51_)
                nw465_[int(5)] = p0
                nw465_[int(6)] = 946
                nw465_[int(7)] = -631
                nw465_[int(8)] = ((d_2311_v7_).cardinality) * ((0) - (self.f14))
                nw465_[int(9)] = 672
                nw465_[int(10)] = (d_2382_v53_).cf57
                nw465_[int(11)] = p0
                nw465_[int(12)] = len((self.f18) + (self.f18))
                nw465_[int(13)] = default__.fm0((self).f17, globalState)
                nw465_[int(14)] = 711
                nw465_[int(15)] = (len(d_2383_v54_)) * (len(d_2385_v56_))
                nw465_[int(16)] = (p0) + (self.f14)
                nw465_[int(17)] = p0
                nw465_[int(18)] = default__.safeDivisionInt(self.f14, len(d_2304_v0_))
                nw465_[int(19)] = default__.safeDivisionInt(self.f14, p0)
                d_2386_v57_ = nw465_
                index401_ = default__.safeIndex(690, (d_2386_v57_).length(0))
                (d_2386_v57_)[index401_] = p0
                d_2387_v58_: T0
                nw466_ = C6()
                nw466_.ctor__(self.f14, not(self.f16), self.f16, not(self.f16), (0) - (p0))
                d_2387_v58_ = nw466_
                d_2388_v59_: D20
                d_2388_v59_ = D20_DC42(p0, self.f16, d_2387_v58_)
                index402_ = default__.safeIndex(690, (d_2386_v57_).length(0))
                (d_2386_v57_)[index402_] = (p0) * (((d_2388_v59_).cf65) - (self.f14))
                d_2389_v60_: _dafny.Array
                nw467_ = _dafny.Array(_dafny.Map({}), 1)
                d_2389_v60_ = nw467_
                d_2390_v61_: _dafny.Map
                d_2390_v61_ = _dafny.Map({(self).f17: (d_2386_v57_)[default__.safeIndex(690, (d_2386_v57_).length(0))]})
                rhs447_ = d_2389_v60_
                rhs448_ = default__.fm31((len(d_2390_v61_)) in (d_2304_v0_), True, globalState)
                rhs449_ = 696
                rhs450_ = (d_2378_v49_) + (d_2378_v49_)
                lhs349_ = self
                lhs350_ = globalState
                d_2389_v60_ = rhs447_
                lhs349_.f16 = rhs448_
                lhs350_.f2 = rhs449_
                d_2378_v49_ = rhs450_
                d_2380_v51_ = (d_2380_v51_).set(not((d_2366_v40_) in (self.f18)), (self).f17)
                d_2391_v62_: D14
                d_2391_v62_ = D14_DC28(d_2387_v58_.f14, (self).f17, d_2387_v58_)
                d_2392_v63_: _dafny.Array
                nw468_ = _dafny.Array(None, 26)
                nw468_[int(0)] = self.f16
                nw468_[int(1)] = self.f16
                nw468_[int(2)] = self.f16
                nw468_[int(3)] = self.f16
                nw468_[int(4)] = self.f16
                nw468_[int(5)] = (self).f17
                nw468_[int(6)] = True
                nw468_[int(7)] = (self).f17
                nw468_[int(8)] = self.f16
                nw468_[int(9)] = self.f16
                nw468_[int(10)] = self.f16
                nw468_[int(11)] = self.f16
                nw468_[int(12)] = (self).f17
                nw468_[int(13)] = (self).f17
                nw468_[int(14)] = True
                nw468_[int(15)] = self.f16
                nw468_[int(16)] = (self).f17
                nw468_[int(17)] = self.f16
                nw468_[int(18)] = True
                nw468_[int(19)] = self.f16
                nw468_[int(20)] = (self).f17
                nw468_[int(21)] = (self).f17
                nw468_[int(22)] = (d_2391_v62_).cf41
                nw468_[int(23)] = self.f16
                nw468_[int(24)] = (self).f17
                nw468_[int(25)] = self.f16
                d_2392_v63_ = nw468_
                d_2393_v64_: _dafny.MultiSet
                d_2393_v64_ = _dafny.MultiSet([d_2392_v63_])
                nw469_ = _dafny.Array(None, 2)
                nw469_[int(0)] = self.f14
                nw469_[int(1)] = ((d_2393_v64_)[d_2392_v63_] if (d_2392_v63_) in (d_2393_v64_) else self.f14)
                d_2384_v55_ = nw469_
            elif True:
                d_2394_v65_: _dafny.Array
                nw470_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_2394_v65_ = nw470_
                d_2395_v67_: _dafny.Array
                def lambda132_(d_2396_i5_):
                    def iife186_():
                        coll82_ = _dafny.Map()
                        compr_82_: int
                        for compr_82_ in _dafny.IntegerRange(785, -560):
                            d_2397_v66_: int = compr_82_
                            if ((785) <= (d_2397_v66_)) and ((d_2397_v66_) < (-560)):
                                coll82_[(d_2397_v66_) - (self.f14)] = (self).f17
                        return _dafny.Map(coll82_)
                    return iife186_()
                    

                init77_ = lambda132_
                nw471_ = _dafny.Array(None, 10)
                for i0_77_ in range(nw471_.length(0)):
                    nw471_[i0_77_] = init77_(i0_77_)
                d_2395_v67_ = nw471_
                index403_ = default__.safeIndex(670, (d_2394_v65_).length(0))
                (d_2394_v65_)[index403_] = d_2395_v67_
                index404_ = default__.safeIndex(670, (d_2394_v65_).length(0))
                nw472_ = _dafny.Array(_dafny.Map({}), 2)
                (d_2394_v65_)[index404_] = nw472_
                d_2398_v68_: _dafny.Array
                def lambda133_(d_2399_i6_):
                    return self.f16

                init78_ = lambda133_
                nw473_ = _dafny.Array(None, 18)
                for i0_78_ in range(nw473_.length(0)):
                    nw473_[i0_78_] = init78_(i0_78_)
                d_2398_v68_ = nw473_
                index405_ = default__.safeIndex(611, (d_2398_v68_).length(0))
                (d_2398_v68_)[index405_] = (not(self.f16)) and (not(default__.fm31(self.f16, False, globalState)))
                d_2400_v69_: _dafny.Set
                d_2400_v69_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2401_i7_ in range(default__.abs(-978))])})
                index406_ = default__.safeIndex(611, (d_2398_v68_).length(0))
                (d_2398_v68_)[index406_] = ((d_2400_v69_) - (d_2400_v69_)).isdisjoint((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jm"))})).intersection(d_2400_v69_))
                d_2402_v70_: _dafny.Seq
                d_2402_v70_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                d_2403_v71_: _dafny.Map
                d_2403_v71_ = _dafny.Map({d_2398_v68_: (d_2402_v70_)[default__.safeIndex(self.f14, len(d_2402_v70_))]})
                (self).f16 = ((d_2403_v71_)[d_2398_v68_] if (d_2398_v68_) in (d_2403_v71_) else ((d_2398_v68_)[default__.safeIndex(611, (d_2398_v68_).length(0))] if (d_2398_v68_)[default__.safeIndex(611, (d_2398_v68_).length(0))] else self.f16))
                d_2404_v72_: _dafny.Seq
                d_2404_v72_ = _dafny.SeqWithoutIsStrInference([d_2311_v7_])
                d_2405_v73_: C1
                nw474_ = C1()
                nw474_.ctor__((len(d_2404_v72_)) == (self.f14), (p0) * (default__.fm0(False, globalState)))
                d_2405_v73_ = nw474_
                d_2406_v74_: _dafny.Map
                d_2406_v74_ = _dafny.Map({d_2405_v73_.f25: self.f16})
                d_2407_v75_: _dafny.Seq
                d_2407_v75_ = _dafny.SeqWithoutIsStrInference([d_2406_v74_, d_2406_v74_])
                (d_2405_v73_).f25 = ((_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}) for d_2408_i8_ in range(default__.abs(659))])) + (d_2407_v75_)) < (d_2407_v75_)
            d_2409_v76_: D19
            d_2409_v76_ = D19_DC40(p0, self.f16, p0)
            d_2410_v77_: _dafny.Array
            nw475_ = _dafny.Array(None, 20)
            nw475_[int(0)] = d_2304_v0_
            nw475_[int(1)] = (d_2304_v0_) + (d_2304_v0_)
            nw475_[int(2)] = d_2304_v0_
            nw475_[int(3)] = (_dafny.SeqWithoutIsStrInference([self.f14, p0, (d_2311_v7_).cardinality, len(d_2304_v0_), (d_2409_v76_).cf61])) + (d_2304_v0_)
            nw475_[int(4)] = d_2304_v0_
            nw475_[int(5)] = d_2304_v0_
            nw475_[int(6)] = d_2304_v0_
            nw475_[int(7)] = d_2304_v0_
            nw475_[int(8)] = d_2304_v0_
            nw475_[int(9)] = d_2304_v0_
            nw475_[int(10)] = d_2304_v0_
            nw475_[int(11)] = _dafny.SeqWithoutIsStrInference([self.f14 for d_2411_i9_ in range(default__.abs(302))])
            nw475_[int(12)] = d_2304_v0_
            nw475_[int(13)] = d_2304_v0_
            nw475_[int(14)] = d_2304_v0_
            nw475_[int(15)] = (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhr"))))]) if (self).f17 else d_2304_v0_)
            nw475_[int(16)] = d_2304_v0_
            nw475_[int(17)] = _dafny.SeqWithoutIsStrInference([len(d_2376_v47_)])
            nw475_[int(18)] = d_2304_v0_
            nw475_[int(19)] = (_dafny.SeqWithoutIsStrInference([131])) + (_dafny.SeqWithoutIsStrInference([p0]))
            d_2410_v77_ = nw475_
            d_2410_v77_ = d_2410_v77_
        (globalState).f7 = ((0) - (self.f14) if self.f16 else self.f14)
        d_2412_v78_: _dafny.Seq
        d_2412_v78_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17, default__.fm17((self).f17, self.f14, p0, (d_2311_v7_).cardinality, globalState)])
        hi20_ = p0
        for d_2413_i10_ in range(len(d_2412_v78_), hi20_):
            d_2414_v79_: _dafny.Map
            d_2414_v79_ = _dafny.Map({(self).f17: (self).f17})
            d_2415_v80_: _dafny.Array
            def lambda134_(d_2416_i11_):
                return self.f16

            init79_ = lambda134_
            nw476_ = _dafny.Array(None, 13)
            for i0_79_ in range(nw476_.length(0)):
                nw476_[i0_79_] = init79_(i0_79_)
            d_2415_v80_ = nw476_
            d_2417_v81_: _dafny.Map
            d_2417_v81_ = _dafny.Map({((d_2414_v79_)[self.f16] if (self.f16) in (d_2414_v79_) else (self).f17): d_2415_v80_})
            (globalState).f2 = (p0) * (len(d_2417_v81_))
            d_2418_v82_: str
            d_2418_v82_ = _dafny.CodePoint('l')
            d_2419_v83_: _dafny.Array
            def lambda135_(d_2420_p0_):
                def lambda136_(d_2421_i12_):
                    return (_dafny.Map({self.f14: (self).f17})) | (_dafny.Map({d_2420_p0_: self.f16}))

                return lambda136_

            init80_ = lambda135_(p0)
            nw477_ = _dafny.Array(None, 8)
            for i0_80_ in range(nw477_.length(0)):
                nw477_[i0_80_] = init80_(i0_80_)
            d_2419_v83_ = nw477_
            d_2422_v84_: _dafny.Map
            d_2422_v84_ = _dafny.Map({d_2413_i10_: (self).f17})
            index407_ = default__.safeIndex(490, (d_2419_v83_).length(0))
            (d_2419_v83_)[index407_] = (_dafny.Map({244: False})) | (d_2422_v84_)
            index408_ = default__.safeIndex(490, (d_2419_v83_).length(0))
            rhs451_ = not(((d_2413_i10_) * ((0) - ((self).fm12(d_2413_i10_, globalState)))) <= (d_2413_i10_))
            rhs452_ = (d_2418_v82_ if (d_2418_v82_) in (self.f18) else d_2418_v82_)
            rhs453_ = d_2422_v84_
            rhs454_ = p0
            lhs351_ = self
            lhs352_ = d_2419_v83_
            lhs353_ = default__.safeIndex(490, (d_2419_v83_).length(0))
            lhs354_ = globalState
            lhs351_.f16 = rhs451_
            d_2418_v82_ = rhs452_
            lhs352_[lhs353_] = rhs453_
            lhs354_.f7 = rhs454_
            if (default__.safeDivisionInt(self.f14, p0)) > (default__.safeModuloInt(d_2413_i10_, d_2413_i10_)):
                d_2423_v85_: _dafny.Array
                nw478_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                d_2423_v85_ = nw478_
                index409_ = default__.safeIndex(924, (d_2423_v85_).length(0))
                (d_2423_v85_)[index409_] = self.f18
                d_2424_v86_: D12
                d_2424_v86_ = D12_DC22(self.f18)
                index410_ = default__.safeIndex(924, (d_2423_v85_).length(0))
                (d_2423_v85_)[index410_] = (d_2424_v86_).cf29
                d_2425_v87_: _dafny.Map
                d_2425_v87_ = _dafny.Map({p0: (d_2419_v83_)[default__.safeIndex(490, (d_2419_v83_).length(0))]})
                d_2426_v88_: T0
                nw479_ = C7()
                nw479_.ctor__(self.f16, self.f16, default__.fm0(self.f16, globalState))
                d_2426_v88_ = nw479_
                d_2427_v89_: D14
                d_2427_v89_ = D14_DC28(647, self.f16, d_2426_v88_)
                d_2425_v87_ = (d_2425_v87_).set(((d_2427_v89_).cf40) - (self.f14), ((d_2419_v83_)[default__.safeIndex(490, (d_2419_v83_).length(0))]) | (d_2422_v84_))
                d_2428_v90_: _dafny.Array
                def lambda137_(d_2429_v86_):
                    def lambda138_(d_2430_i13_):
                        return default__.safeModuloInt(d_2430_i13_, len((d_2429_v86_).cf29))

                    return lambda138_

                init81_ = lambda137_(d_2424_v86_)
                nw480_ = _dafny.Array(None, 8)
                for i0_81_ in range(nw480_.length(0)):
                    nw480_[i0_81_] = init81_(i0_81_)
                d_2428_v90_ = nw480_
                d_2431_v91_: _dafny.Map
                d_2431_v91_ = _dafny.Map({self.f14: d_2428_v90_})
                d_2432_v92_: _dafny.Set
                d_2432_v92_ = _dafny.Set({self.f14, self.f14})
                rhs455_ = ((d_2431_v91_)[(0) - (d_2426_v88_.f14)] if ((0) - (d_2426_v88_.f14)) in (d_2431_v91_) else d_2428_v90_)
                rhs456_ = (d_2432_v92_).ispropersubset(d_2432_v92_)
                rhs457_ = d_2426_v88_.f14
                lhs355_ = self
                d_2428_v90_ = rhs455_
                r0 = rhs456_
                lhs355_.f14 = rhs457_
                d_2433_v93_: D23
                d_2433_v93_ = D23_DC47(_dafny.Map({d_2418_v82_: D20_DC42(d_2413_i10_, True, d_2426_v88_)}))
                d_2433_v93_ = d_2433_v93_
                d_2434_v94_: D9
                d_2434_v94_ = D9_DC15(self.f16, self.f16, _dafny.Set({d_2413_i10_, len(d_2414_v79_), 281, d_2413_i10_}), d_2428_v90_, d_2415_v80_)
                d_2435_v96_: _dafny.Array
                nw481_ = _dafny.Array(None, 20)
                nw481_[int(0)] = d_2434_v94_
                nw481_[int(1)] = d_2434_v94_
                nw481_[int(2)] = d_2434_v94_
                nw481_[int(3)] = d_2434_v94_
                nw481_[int(4)] = D9_DC15((self).f17, self.f16, d_2432_v92_, d_2428_v90_, d_2415_v80_)
                nw481_[int(5)] = d_2434_v94_
                nw481_[int(6)] = d_2434_v94_
                nw481_[int(7)] = d_2434_v94_
                nw481_[int(8)] = d_2434_v94_
                nw481_[int(9)] = d_2434_v94_
                nw481_[int(10)] = d_2434_v94_
                nw481_[int(11)] = d_2434_v94_
                nw481_[int(12)] = d_2434_v94_
                nw481_[int(13)] = d_2434_v94_
                nw481_[int(14)] = D9_DC15((self).fm14(globalState), (self).f17, _dafny.Set({-27}), d_2428_v90_, d_2415_v80_)
                nw481_[int(15)] = d_2434_v94_
                nw481_[int(16)] = d_2434_v94_
                nw481_[int(17)] = d_2434_v94_
                def iife187_():
                    coll83_ = _dafny.Set()
                    compr_83_: int
                    for compr_83_ in _dafny.IntegerRange(129, 847):
                        d_2436_v95_: int = compr_83_
                        if ((129) <= (d_2436_v95_)) and ((d_2436_v95_) < (847)):
                            coll83_ = coll83_.union(_dafny.Set([default__.safeModuloInt(d_2436_v95_, p0)]))
                    return _dafny.Set(coll83_)
                nw481_[int(18)] = D9_DC15((self).f17, (self).f17, iife187_()
, d_2428_v90_, d_2415_v80_)
                nw481_[int(19)] = d_2434_v94_
                d_2435_v96_ = nw481_
                d_2437_v97_: _dafny.Seq
                d_2437_v97_ = _dafny.SeqWithoutIsStrInference([d_2435_v96_])
                d_2438_v98_: D3
                d_2438_v98_ = D3_DC3(self.f16)
                d_2439_v99_: _dafny.Seq
                d_2439_v99_ = _dafny.SeqWithoutIsStrInference([d_2438_v98_])
                d_2440_v100_: D12
                d_2440_v100_ = D12_DC24(self.f16, d_2412_v78_, d_2437_v97_, self.f16, d_2439_v99_)
                d_2441_v101_: T2
                nw482_ = C5()
                nw482_.ctor__((self).f17, (d_2440_v100_).cf36)
                d_2441_v101_ = nw482_
                d_2442_v102_: _dafny.Map
                d_2442_v102_ = _dafny.Map({d_2441_v101_: d_2441_v101_.f16})
                d_2443_v104_: _dafny.MultiSet
                d_2443_v104_ = _dafny.MultiSet([self.f18])
                d_2444_v105_: _dafny.Map
                d_2444_v105_ = _dafny.Map({(d_2423_v85_)[default__.safeIndex(924, (d_2423_v85_).length(0))]: self.f16})
                d_2445_v106_: _dafny.Map
                def iife188_():
                    coll84_ = _dafny.Map()
                    compr_84_: _dafny.Seq
                    for compr_84_ in (d_2443_v104_).Elements:
                        d_2446_v103_: _dafny.Seq = compr_84_
                        if (d_2446_v103_) in (d_2443_v104_):
                            coll84_[d_2446_v103_] = (self).f17
                    return _dafny.Map(coll84_)
                d_2445_v106_ = _dafny.Map({default__.safeModuloInt(769, d_2426_v88_.f14): (iife188_()
                 if ((d_2442_v102_)[d_2441_v101_] if (d_2441_v101_) in (d_2442_v102_) else d_2441_v101_.f16) else d_2444_v105_)})
                d_2445_v106_ = (d_2445_v106_).set(840, (_dafny.Map({(d_2423_v85_)[default__.safeIndex(924, (d_2423_v85_).length(0))]: False})) | (d_2444_v105_))
            elif True:
                rhs458_ = (p0) - (d_2413_i10_)
                rhs459_ = self.f14
                rhs460_ = p0
                rhs461_ = d_2418_v82_
                lhs356_ = globalState
                lhs357_ = self
                lhs358_ = self
                lhs356_.f2 = rhs458_
                lhs357_.f14 = rhs459_
                lhs358_.f14 = rhs460_
                d_2418_v82_ = rhs461_
                d_2447_v107_: _dafny.Set
                d_2447_v107_ = _dafny.Set({self.f14})
                d_2447_v107_ = (d_2447_v107_) - (d_2447_v107_)
                d_2447_v107_ = d_2447_v107_
                (globalState).f7 = p0
                d_2448_v108_: _dafny.Set
                d_2448_v108_ = _dafny.Set({(self).f17})
                (self).f16 = not(not(not((d_2448_v108_).issubset((d_2448_v108_) | (d_2448_v108_)))))
            d_2449_v109_: C1
            nw483_ = C1()
            nw483_.ctor__(self.f16, d_2413_i10_)
            d_2449_v109_ = nw483_
        d_2450_v110_: _dafny.Map
        d_2450_v110_ = _dafny.Map({(d_2306_v2_).cardinality: default__.fm44(False, (self).f17, globalState)})
        d_2451_v111_: D18
        d_2451_v111_ = D18_DC38(self.f14, self.f16, len((d_2450_v110_) | (d_2450_v110_)))
        d_2451_v111_ = d_2451_v111_
        d_2452_v112_: _dafny.Set
        d_2452_v112_ = _dafny.Set({618})
        d_2453_v113_: _dafny.Array
        def lambda139_(d_2454_i14_):
            return (d_2454_i14_) - (len(self.f18))

        init82_ = lambda139_
        nw484_ = _dafny.Array(None, 16)
        for i0_82_ in range(nw484_.length(0)):
            nw484_[i0_82_] = init82_(i0_82_)
        d_2453_v113_ = nw484_
        d_2455_v114_: _dafny.Array
        nw485_ = _dafny.Array(False, 28)
        d_2455_v114_ = nw485_
        d_2456_v115_: D9
        d_2456_v115_ = D9_DC15((self).f17, (self).f17, d_2452_v112_, d_2453_v113_, d_2455_v114_)
        d_2457_v116_: _dafny.Array
        nw486_ = _dafny.Array(None, 6)
        nw486_[int(0)] = (d_2456_v115_).cf22
        nw486_[int(1)] = d_2455_v114_
        nw486_[int(2)] = d_2455_v114_
        nw486_[int(3)] = d_2455_v114_
        nw486_[int(4)] = d_2455_v114_
        nw486_[int(5)] = d_2455_v114_
        d_2457_v116_ = nw486_
        d_2458_v117_: _dafny.Map
        d_2458_v117_ = _dafny.Map({d_2457_v116_: (_dafny.Map({len(d_2307_v3_): (self).f17})).set(self.f14, (self).f17)})
        d_2459_v118_: _dafny.Map
        d_2459_v118_ = _dafny.Map({p0: (self).f17})
        d_2458_v117_ = (d_2458_v117_).set(d_2457_v116_, d_2459_v118_)
        r0 = (self).f17
        return r0

    def m8(self, p0, p1, p2, p3, globalState):
        if (self.f16) and ((self).f17):
            d_2460_v0_: C6
            nw487_ = C6()
            nw487_.ctor__(self.f14, self.f16, self.f16, (default__.fm43((self).f17, globalState) if (self).f17 else not(False)), p2)
            d_2460_v0_ = nw487_
            d_2461_v1_: _dafny.Set
            d_2461_v1_ = _dafny.Set({p2})
            d_2462_v2_: _dafny.Map
            d_2462_v2_ = _dafny.Map({p3: (d_2460_v0_).f19})
            rhs462_ = not(not((d_2461_v1_).issubset(d_2461_v1_)))
            rhs463_ = (d_2460_v0_).fm12(default__.fm0(not((d_2460_v0_).f20), globalState), globalState)
            rhs464_ = default__.fm0(default__.fm27(self.f14, p3, self.f14, d_2462_v2_, globalState), globalState)
            lhs359_ = self
            lhs360_ = globalState
            lhs361_ = globalState
            lhs359_.f16 = rhs462_
            lhs360_.f2 = rhs463_
            lhs361_.f7 = rhs464_
            d_2463_v3_: C4
            nw488_ = C4()
            nw488_.ctor__(True, (d_2460_v0_).f20, not(False))
            d_2463_v3_ = nw488_
            d_2464_v4_: str
            d_2464_v4_ = _dafny.CodePoint('r')
            d_2465_v5_: _dafny.Array
            def lambda140_(d_2466_v3_):
                def lambda141_(d_2467_i0_):
                    return (d_2466_v3_).f21

                return lambda141_

            init83_ = lambda140_(d_2463_v3_)
            nw489_ = _dafny.Array(None, 4)
            for i0_83_ in range(nw489_.length(0)):
                nw489_[i0_83_] = init83_(i0_83_)
            d_2465_v5_ = nw489_
            d_2468_v6_: _dafny.MultiSet
            d_2468_v6_ = _dafny.MultiSet([d_2465_v5_, d_2465_v5_])
            d_2469_v7_: _dafny.Map
            d_2469_v7_ = _dafny.Map({(d_2460_v0_).f20: self.f16})
            d_2470_v8_: _dafny.Map
            d_2470_v8_ = _dafny.Map({(d_2468_v6_).cardinality: d_2469_v7_})
            d_2471_v9_: int
            out35_: int
            out35_ = (d_2460_v0_).m3(d_2464_v4_, len((d_2470_v8_) | (d_2470_v8_)), (d_2463_v3_).f21, (d_2460_v0_).f19, globalState)
            d_2471_v9_ = out35_
            d_2472_v10_: _dafny.MultiSet
            d_2472_v10_ = _dafny.MultiSet([self.f18])
            (self).f16 = ((self).f17) or ((d_2472_v10_).ispropersubset(d_2472_v10_))
        elif True:
            (self).f16 = False
            (self).f14 = p3
            (self).f16 = (self).f17
            d_2473_v11_: _dafny.Array
            def lambda142_(d_2474_i1_):
                return _dafny.Map({_dafny.CodePoint('r'): (self).f17})

            init84_ = lambda142_
            nw490_ = _dafny.Array(None, 12)
            for i0_84_ in range(nw490_.length(0)):
                nw490_[i0_84_] = init84_(i0_84_)
            d_2473_v11_ = nw490_
            d_2475_v12_: D11
            d_2475_v12_ = D11_DC20()
            d_2476_v13_: _dafny.Map
            d_2476_v13_ = _dafny.Map({(default__.fm68(D11_DC21(d_2475_v12_), (self).f17, globalState)).cf14: False})
            index411_ = default__.safeIndex(637, (d_2473_v11_).length(0))
            (d_2473_v11_)[index411_] = (d_2476_v13_) | (d_2476_v13_)
            index412_ = default__.safeIndex(637, (d_2473_v11_).length(0))
            (d_2473_v11_)[index412_] = d_2476_v13_
            d_2477_v14_: _dafny.Array
            def lambda143_(d_2478_i2_):
                return (_dafny.MultiSet([(self).f17, (self).f17, False])) - (_dafny.MultiSet([self.f16]))

            init85_ = lambda143_
            nw491_ = _dafny.Array(None, 11)
            for i0_85_ in range(nw491_.length(0)):
                nw491_[i0_85_] = init85_(i0_85_)
            d_2477_v14_ = nw491_
            d_2479_v15_: _dafny.MultiSet
            d_2479_v15_ = _dafny.MultiSet([(self).f17])
            index413_ = default__.safeIndex(284, (d_2477_v14_).length(0))
            (d_2477_v14_)[index413_] = d_2479_v15_
            index414_ = default__.safeIndex(284, (d_2477_v14_).length(0))
            (d_2477_v14_)[index414_] = ((d_2479_v15_) - (_dafny.MultiSet([self.f16, self.f16, (self).f17, False, self.f16]))) - (d_2479_v15_)
        d_2480_v16_: _dafny.Array
        nw492_ = _dafny.Array(int(0), 13)
        d_2480_v16_ = nw492_
        d_2481_v17_: D9
        d_2481_v17_ = D9_DC14(d_2480_v16_)
        d_2482_v18_: _dafny.Set
        d_2482_v18_ = _dafny.Set({p2, len(self.f18), p3})
        d_2483_v19_: _dafny.Array
        nw493_ = _dafny.Array(False, 8)
        d_2483_v19_ = nw493_
        d_2484_v20_: D9
        d_2484_v20_ = D9_DC15(self.f16, self.f16, d_2482_v18_, d_2480_v16_, d_2483_v19_)
        d_2485_v21_: _dafny.Array
        nw494_ = _dafny.Array(None, 7)
        d_2485_v21_ = nw494_
        d_2486_v22_: _dafny.Map
        d_2486_v22_ = _dafny.Map({d_2485_v21_: d_2480_v16_})
        d_2487_v23_: _dafny.Array
        nw495_ = _dafny.Array(None, 26)
        nw495_[int(0)] = d_2480_v16_
        nw495_[int(1)] = d_2480_v16_
        nw495_[int(2)] = d_2480_v16_
        nw495_[int(3)] = d_2480_v16_
        nw495_[int(4)] = d_2480_v16_
        nw495_[int(5)] = (d_2481_v17_).cf17
        nw495_[int(6)] = d_2480_v16_
        nw495_[int(7)] = (d_2480_v16_ if self.f16 else d_2480_v16_)
        nw495_[int(8)] = d_2480_v16_
        nw495_[int(9)] = d_2480_v16_
        nw495_[int(10)] = (d_2484_v20_).cf21
        nw495_[int(11)] = d_2480_v16_
        nw495_[int(12)] = d_2480_v16_
        nw495_[int(13)] = d_2480_v16_
        nw495_[int(14)] = d_2480_v16_
        nw495_[int(15)] = ((d_2486_v22_)[d_2485_v21_] if (d_2485_v21_) in (d_2486_v22_) else d_2480_v16_)
        nw495_[int(16)] = (d_2484_v20_).cf21
        nw495_[int(17)] = d_2480_v16_
        nw495_[int(18)] = d_2480_v16_
        nw495_[int(19)] = d_2480_v16_
        nw495_[int(20)] = d_2480_v16_
        nw495_[int(21)] = d_2480_v16_
        nw495_[int(22)] = d_2480_v16_
        nw495_[int(23)] = d_2480_v16_
        nw495_[int(24)] = d_2480_v16_
        nw495_[int(25)] = d_2480_v16_
        d_2487_v23_ = nw495_
        index415_ = default__.safeIndex(294, (d_2487_v23_).length(0))
        (d_2487_v23_)[index415_] = d_2480_v16_
        index416_ = default__.safeIndex(294, (d_2487_v23_).length(0))
        (d_2487_v23_)[index416_] = d_2480_v16_
        (self).f16 = (self).f17
        d_2488_v24_: _dafny.Map
        d_2488_v24_ = _dafny.Map({p2: 175})
        d_2489_v25_: D11
        d_2489_v25_ = D11_DC19(d_2488_v24_)
        d_2490_i3_: int
        d_2490_i3_ = 0
        with _dafny.label("15"):
            pat_let_tv40_ = p3
            pat_let_tv41_ = p3
            def lambda144_(source38_):
                if source38_.is_DC20:
                    return self.f16
                elif source38_.is_DC19:
                    d_2517___mcc_h7_ = source38_.cf27
                    d_2518_cf27_ = d_2517___mcc_h7_
                    return (pat_let_tv40_) > (pat_let_tv41_)
                elif True:
                    d_2519___mcc_h8_ = source38_.cf28
                    d_2520_cf28_ = d_2519___mcc_h8_
                    return (self).f17

            while not(lambda144_(d_2489_v25_)):
                with _dafny.c_label("15"):
                    if (d_2490_i3_) >= (100):
                        raise _dafny.Break("15")
                    d_2490_i3_ = (d_2490_i3_) + (1)
                    d_2491_v26_: D3
                    d_2491_v26_ = D3_DC4(True, p2, p3, (self).f17, not(self.f16))
                    source37_ = d_2491_v26_
                    if source37_.is_DC4:
                        d_2492___mcc_h0_ = source37_.cf4
                        d_2493___mcc_h1_ = source37_.cf5
                        d_2494___mcc_h2_ = source37_.cf6
                        d_2495___mcc_h3_ = source37_.cf7
                        d_2496___mcc_h4_ = source37_.cf8
                        d_2497_cf8_ = d_2496___mcc_h4_
                        d_2498_cf7_ = d_2495___mcc_h3_
                        d_2499_cf6_ = d_2494___mcc_h2_
                        d_2500_cf5_ = d_2493___mcc_h1_
                        d_2501_cf4_ = d_2492___mcc_h0_
                        (self).f16 = d_2498_cf7_
                        index417_ = default__.safeIndex(294, (d_2483_v19_).length(0))
                        (d_2483_v19_)[index417_] = d_2497_cf8_
                        index418_ = default__.safeIndex(294, (d_2483_v19_).length(0))
                        def iife189_():
                            coll85_ = _dafny.Map()
                            compr_85_: int
                            for compr_85_ in _dafny.IntegerRange(348, 586):
                                d_2502_v27_: int = compr_85_
                                if ((348) <= (d_2502_v27_)) and ((d_2502_v27_) < (586)):
                                    coll85_[default__.safeDivisionInt(d_2502_v27_, d_2500_cf5_)] = p3
                            return _dafny.Map(coll85_)
                        (d_2483_v19_)[index418_] = (len((iife189_()
                        ) | (d_2488_v24_))) == (257)
                        d_2499_cf6_ = (d_2500_cf5_ if default__.fm17(not(d_2497_cf8_), (0) - (self.f14), len(self.f18), 867, globalState) else default__.safeModuloInt(d_2499_cf6_, p3))
                        d_2503_v28_: _dafny.Array
                        nw496_ = _dafny.Array(_dafny.Set({}), 19)
                        d_2503_v28_ = nw496_
                        d_2504_v29_: C0
                        nw497_ = C0()
                        nw497_.ctor__(d_2503_v28_, d_2500_cf5_)
                        d_2504_v29_ = nw497_
                    elif source37_.is_DC3:
                        d_2505___mcc_h5_ = source37_.cf3
                        d_2506_cf3_ = d_2505___mcc_h5_
                        (globalState).f7 = p2
                        (self).f16 = self.f16
                        d_2507_v30_: str
                        d_2507_v30_ = _dafny.CodePoint('m')
                        d_2506_cf3_ = (d_2507_v30_) not in ((self.f18) + (self.f18))
                        (globalState).f2 = self.f14
                    elif True:
                        d_2508___mcc_h6_ = source37_.cf9
                        d_2509_cf9_ = d_2508___mcc_h6_
                        d_2510_v31_: _dafny.Map
                        d_2510_v31_ = _dafny.Map({self.f16: self.f16})
                        d_2511_v32_: C5
                        nw498_ = C5()
                        nw498_.ctor__(False, ((d_2510_v31_)[self.f16] if (self.f16) in (d_2510_v31_) else self.f16))
                        d_2511_v32_ = nw498_
                        d_2512_v33_: D3
                        d_2512_v33_ = D3_DC3((self).f17)
                        d_2513_v34_: C2
                        nw499_ = C2()
                        nw499_.ctor__(d_2512_v33_, (self).f17, (self).f17)
                        d_2513_v34_ = nw499_
                        (globalState).f7 = (0) - ((p2) * ((default__.fm0(self.f16, globalState)) * (p3)))
                        d_2514_v35_: _dafny.Array
                        nw500_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                        d_2514_v35_ = nw500_
                        d_2515_v36_: str
                        d_2515_v36_ = _dafny.CodePoint('a')
                        index419_ = default__.safeIndex(83, (d_2514_v35_).length(0))
                        (d_2514_v35_)[index419_] = d_2515_v36_
                        index420_ = default__.safeIndex(83, (d_2514_v35_).length(0))
                        (d_2514_v35_)[index420_] = d_2515_v36_
                    index421_ = default__.safeIndex(23, (d_2480_v16_).length(0))
                    (d_2480_v16_)[index421_] = p3
                    index422_ = default__.safeIndex(23, (d_2480_v16_).length(0))
                    (d_2480_v16_)[index422_] = (715) * (p3)
                    d_2516_v37_: _dafny.Set
                    d_2516_v37_ = _dafny.Set({(self).f17, (self).f17, (self).f17})
                    d_2516_v37_ = ((d_2516_v37_).intersection(d_2516_v37_)).intersection((d_2516_v37_).intersection(d_2516_v37_))
                    (globalState).f7 = default__.fm0(self.f16, globalState)
                    pass
            pass
        (self).f18 = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('u') if self.f16 else _dafny.CodePoint('n')) for d_2521_i4_ in range(default__.abs(543))])
        d_2522_v38_: _dafny.Map
        d_2522_v38_ = _dafny.Map({self.f14: d_2483_v19_})
        d_2523_v39_: _dafny.Set
        d_2523_v39_ = _dafny.Set({d_2480_v16_})
        d_2522_v38_ = (d_2522_v38_).set(len(d_2523_v39_), d_2483_v19_)


class C9:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self):
        pass
        pass

    def m4(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_2524_v0_: _dafny.Seq
        d_2524_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnopbue"))
        d_2525_v1_: _dafny.Set
        d_2525_v1_ = _dafny.Set({False, True, p0})
        d_2526_v2_: C8
        nw501_ = C8()
        nw501_.ctor__(d_2524_v0_, (p0) not in (d_2525_v1_), (p0 if False else p0), (0) - ((488) * (179)))
        d_2526_v2_ = nw501_
        hi21_ = p3
        for d_2527_i0_ in range((p1 if p0 else p1), hi21_):
            (globalState).f2 = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nno"))))
            d_2528_v3_: _dafny.Array
            def lambda145_(d_2529_p1_):
                def lambda146_(d_2530_i1_):
                    return (d_2530_i1_) * (d_2529_p1_)

                return lambda146_

            init86_ = lambda145_(p1)
            nw502_ = _dafny.Array(None, 22)
            for i0_86_ in range(nw502_.length(0)):
                nw502_[i0_86_] = init86_(i0_86_)
            d_2528_v3_ = nw502_
            index423_ = default__.safeIndex(628, (d_2528_v3_).length(0))
            (d_2528_v3_)[index423_] = len(d_2526_v2_.f18)
            index424_ = default__.safeIndex(628, (d_2528_v3_).length(0))
            (d_2528_v3_)[index424_] = p2
            d_2531_v4_: str
            d_2531_v4_ = _dafny.CodePoint('g')
            d_2532_v5_: _dafny.MultiSet
            d_2532_v5_ = _dafny.MultiSet([d_2531_v4_, d_2531_v4_, d_2531_v4_, d_2531_v4_, d_2531_v4_])
            d_2533_v6_: _dafny.Seq
            d_2533_v6_ = _dafny.SeqWithoutIsStrInference([d_2532_v5_])
            index425_ = default__.safeIndex(628, (d_2528_v3_).length(0))
            index426_ = default__.safeIndex(628, (d_2528_v3_).length(0))
            rhs465_ = p2
            rhs466_ = ((d_2533_v6_)[default__.safeIndex(len(d_2526_v2_.f18), len(d_2533_v6_))]).cardinality
            rhs467_ = p3
            rhs468_ = (d_2525_v1_) | (d_2525_v1_)
            rhs469_ = p2
            lhs362_ = d_2528_v3_
            lhs363_ = default__.safeIndex(628, (d_2528_v3_).length(0))
            lhs364_ = globalState
            lhs365_ = d_2528_v3_
            lhs366_ = default__.safeIndex(628, (d_2528_v3_).length(0))
            lhs367_ = globalState
            lhs362_[lhs363_] = rhs465_
            lhs364_.f2 = rhs466_
            lhs365_[lhs366_] = rhs467_
            d_2525_v1_ = rhs468_
            lhs367_.f7 = rhs469_
            d_2534_v7_: _dafny.Array
            nw503_ = _dafny.Array(None, 5)
            d_2534_v7_ = nw503_
            d_2535_v8_: C6
            nw504_ = C6()
            nw504_.ctor__(297, p0, p0, not(p0), p1)
            d_2535_v8_ = nw504_
            index427_ = default__.safeIndex(145, (d_2534_v7_).length(0))
            (d_2534_v7_)[index427_] = d_2535_v8_
            index428_ = default__.safeIndex(145, (d_2534_v7_).length(0))
            (d_2534_v7_)[index428_] = d_2535_v8_
        d_2536_i2_: int
        d_2536_i2_ = 0
        with _dafny.label("16"):
            while False:
                with _dafny.c_label("16"):
                    if (d_2536_i2_) >= (100):
                        raise _dafny.Break("16")
                    d_2536_i2_ = (d_2536_i2_) + (1)
                    (globalState).f2 = 820
                    d_2537_v9_: _dafny.Array
                    def lambda147_(d_2538_p0_):
                        def lambda148_(d_2539_i3_):
                            return d_2538_p0_

                        return lambda148_

                    init87_ = lambda147_(p0)
                    nw505_ = _dafny.Array(None, 19)
                    for i0_87_ in range(nw505_.length(0)):
                        nw505_[i0_87_] = init87_(i0_87_)
                    d_2537_v9_ = nw505_
                    index429_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                    (d_2537_v9_)[index429_] = p0
                    index430_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                    (d_2537_v9_)[index430_] = p0
                    if default__.fm31((d_2537_v9_)[default__.safeIndex(756, (d_2537_v9_).length(0))], p0, globalState):
                        index431_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                        (d_2537_v9_)[index431_] = default__.fm17((d_2537_v9_)[default__.safeIndex(756, (d_2537_v9_).length(0))], 843, (637) - ((0) - (p1)), default__.safeDivisionInt(p2, p3), globalState)
                        d_2540_v10_: _dafny.Array
                        nw506_ = _dafny.Array(None, 6)
                        nw506_[int(0)] = p2
                        nw506_[int(1)] = p3
                        nw506_[int(2)] = p1
                        nw506_[int(3)] = 352
                        nw506_[int(4)] = (0) - ((0) - (len(d_2524_v0_)))
                        nw506_[int(5)] = p1
                        d_2540_v10_ = nw506_
                        index432_ = default__.safeIndex(36, (d_2540_v10_).length(0))
                        (d_2540_v10_)[index432_] = (len(d_2526_v2_.f18)) * (p2)
                        index433_ = default__.safeIndex(36, (d_2540_v10_).length(0))
                        (d_2540_v10_)[index433_] = p3
                        d_2541_v11_: _dafny.Seq
                        d_2541_v11_ = _dafny.SeqWithoutIsStrInference([p3, (d_2540_v10_)[default__.safeIndex(36, (d_2540_v10_).length(0))], p3, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiimxauy"))) if p0 else (d_2540_v10_)[default__.safeIndex(36, (d_2540_v10_).length(0))])])
                        d_2542_v12_: _dafny.Seq
                        d_2542_v12_ = d_2541_v11_
                        d_2541_v11_ = (((d_2542_v12_)).set(default__.safeIndex(-484, len((d_2542_v12_))), p3)) + (d_2541_v11_)
                        d_2537_v9_ = d_2537_v9_
                        index434_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                        (d_2537_v9_)[index434_] = (d_2537_v9_)[default__.safeIndex(756, (d_2537_v9_).length(0))]
                    elif True:
                        d_2543_v13_: str
                        d_2543_v13_ = _dafny.CodePoint('b')
                        d_2544_v14_: D15
                        d_2544_v14_ = D15_DC30(d_2543_v13_)
                        d_2545_v15_: _dafny.Map
                        d_2545_v15_ = _dafny.Map({(d_2544_v14_).cf44: True})
                        d_2546_v16_: _dafny.Map
                        d_2546_v16_ = _dafny.Map({(d_2543_v13_) in (d_2545_v15_): p3})
                        d_2546_v16_ = (d_2546_v16_).set(p0, p3)
                        index435_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                        (d_2537_v9_)[index435_] = (d_2525_v1_) == (d_2525_v1_)
                        d_2547_v17_: C7
                        nw507_ = C7()
                        nw507_.ctor__(p0, p0, (p3 if p0 else p2))
                        d_2547_v17_ = nw507_
                        index436_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                        (d_2537_v9_)[index436_] = p0
                        (globalState).f7 = (0) - (len(_dafny.SeqWithoutIsStrInference([d_2543_v13_ for d_2548_i4_ in range(default__.abs(378))])))
                    d_2549_v18_: str
                    d_2549_v18_ = _dafny.CodePoint('o')
                    d_2550_v19_: _dafny.Map
                    d_2550_v19_ = _dafny.Map({d_2549_v18_: p0})
                    index437_ = default__.safeIndex(756, (d_2537_v9_).length(0))
                    (d_2537_v9_)[index437_] = not(((d_2550_v19_)[d_2549_v18_] if (d_2549_v18_) in (d_2550_v19_) else False))
                    pass
            pass
        hi22_ = p3
        for d_2551_i5_ in range(p1, hi22_):
            (globalState).f7 = (0) - (default__.safeDivisionInt((0) - (-213), p2))
            r0 = default__.safeDivisionInt(p1, p1)
            d_2552_v20_: str
            d_2552_v20_ = _dafny.CodePoint('g')
            d_2552_v20_ = _dafny.CodePoint('w')
            d_2553_v21_: _dafny.Array
            def lambda149_(d_2554_i6_):
                return not(True)

            init88_ = lambda149_
            nw508_ = _dafny.Array(None, 7)
            for i0_88_ in range(nw508_.length(0)):
                nw508_[i0_88_] = init88_(i0_88_)
            d_2553_v21_ = nw508_
            index438_ = default__.safeIndex(236, (d_2553_v21_).length(0))
            (d_2553_v21_)[index438_] = p0
            index439_ = default__.safeIndex(236, (d_2553_v21_).length(0))
            (d_2553_v21_)[index439_] = p0
        d_2555_v22_: _dafny.Seq
        d_2555_v22_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2556_v23_: _dafny.Seq
        d_2556_v23_ = d_2555_v22_
        d_2557_v24_: _dafny.MultiSet
        d_2557_v24_ = _dafny.MultiSet([(d_2556_v23_), d_2555_v22_])
        hi23_ = p3
        for d_2558_i7_ in range((d_2526_v2_).fm12(((d_2557_v24_).set(d_2555_v22_, default__.abs((0) - (p2)))).cardinality, globalState), hi23_):
            (globalState).f7 = ((p3) + (d_2558_i7_) if p0 else p1)
            d_2559_v25_: str
            d_2559_v25_ = _dafny.CodePoint('u')
            d_2560_v26_: T0
            nw509_ = C8()
            nw509_.ctor__((d_2526_v2_.f18).set(default__.safeIndex(p2, len(d_2526_v2_.f18)), d_2559_v25_), p0, p0, (0) - (d_2558_i7_))
            d_2560_v26_ = nw509_
            d_2561_v27_: D16
            d_2561_v27_ = D16_DC34((d_2555_v22_)[default__.safeIndex(p3, len(d_2555_v22_))], d_2560_v26_, p1)
            source39_ = d_2561_v27_
            if source39_.is_DC33:
                d_2562___mcc_h0_ = source39_.cf47
                d_2563___mcc_h1_ = source39_.cf48
                d_2564___mcc_h2_ = source39_.cf49
                d_2565___mcc_h3_ = source39_.cf50
                d_2566_cf50_ = d_2565___mcc_h3_
                d_2567_cf49_ = d_2564___mcc_h2_
                d_2568_cf48_ = d_2563___mcc_h1_
                d_2569_cf47_ = d_2562___mcc_h0_
                d_2570_v28_: _dafny.Array
                nw510_ = _dafny.Array(int(0), 22)
                d_2570_v28_ = nw510_
                index440_ = default__.safeIndex(247, (d_2570_v28_).length(0))
                (d_2570_v28_)[index440_] = (0) - (p3)
                d_2571_v29_: _dafny.MultiSet
                d_2571_v29_ = _dafny.MultiSet([p3])
                d_2572_v30_: D16
                d_2572_v30_ = D16_DC33(False, not(p0), d_2558_i7_, d_2566_cf50_)
                pat_let_tv42_ = d_2568_cf48_
                index441_ = default__.safeIndex(247, (d_2570_v28_).length(0))
                def iife190_(_pat_let52_0):
                    def iife191_(d_2573_dt__update__tmp_h0_):
                        def iife192_(_pat_let53_0):
                            def iife193_(d_2574_dt__update_hcf48_h0_):
                                def iife194_(_pat_let54_0):
                                    def iife195_(d_2575_dt__update_hcf47_h0_):
                                        return D16_DC33(d_2575_dt__update_hcf47_h0_, d_2574_dt__update_hcf48_h0_, (d_2573_dt__update__tmp_h0_).cf49, (d_2573_dt__update__tmp_h0_).cf50)
                                    return iife195_(_pat_let54_0)
                                return iife194_(pat_let_tv42_)
                            return iife193_(_pat_let53_0)
                        return iife192_(True)
                    return iife191_(_pat_let52_0)
                (d_2570_v28_)[index441_] = default__.safeModuloInt((d_2571_v29_).cardinality, (iife190_(d_2572_v30_)).cf49)
                index442_ = default__.safeIndex(247, (d_2570_v28_).length(0))
                (d_2570_v28_)[index442_] = (d_2570_v28_)[default__.safeIndex(247, (d_2570_v28_).length(0))]
                d_2576_v31_: _dafny.MultiSet
                d_2576_v31_ = _dafny.MultiSet([d_2559_v25_])
                d_2577_v32_: D21
                d_2577_v32_ = D21_DC43(d_2576_v31_)
                d_2578_v33_: D21
                d_2578_v33_ = D21_DC45(d_2577_v32_)
                d_2578_v33_ = D21_DC45(d_2577_v32_)
                d_2559_v25_ = d_2559_v25_
            elif source39_.is_DC34:
                d_2579___mcc_h4_ = source39_.cf51
                d_2580___mcc_h5_ = source39_.cf52
                d_2581___mcc_h6_ = source39_.cf53
                d_2582_cf53_ = d_2581___mcc_h6_
                d_2583_cf52_ = d_2580___mcc_h5_
                d_2584_cf51_ = d_2579___mcc_h4_
                d_2585_v34_: _dafny.Array
                nw511_ = _dafny.Array(_dafny.Seq({}), 4)
                d_2585_v34_ = nw511_
                d_2585_v34_ = d_2585_v34_
                d_2584_cf51_ = (d_2555_v22_) <= ((d_2555_v22_).set(default__.safeIndex(len(d_2525_v1_), len(d_2555_v22_)), d_2584_cf51_))
                d_2584_cf51_ = p0
                d_2586_v36_: _dafny.Array
                def lambda150_(d_2587_i7_, d_2588_v26_, d_2589_cf52_, d_2590_p1_, d_2591_p3_):
                    def lambda151_(d_2592_i8_):
                        def iife196_():
                            coll86_ = _dafny.Set()
                            compr_86_: int
                            for compr_86_ in _dafny.IntegerRange(980, 226):
                                d_2593_v35_: int = compr_86_
                                if ((980) <= (d_2593_v35_)) and ((d_2593_v35_) < (226)):
                                    coll86_ = coll86_.union(_dafny.Set([default__.safeModuloInt(d_2593_v35_, (0) - (d_2588_v26_.f14))]))
                            return _dafny.Set(coll86_)
                        return _dafny.SeqWithoutIsStrInference([iife196_()
                        , _dafny.Set({d_2587_i7_, 458, d_2589_cf52_.f14, len(_dafny.Set({d_2590_p1_, d_2591_p3_, d_2590_p1_})), d_2588_v26_.f14})])

                    return lambda151_

                init89_ = lambda150_(d_2558_i7_, d_2560_v26_, d_2583_cf52_, p1, p3)
                nw512_ = _dafny.Array(None, 28)
                for i0_89_ in range(nw512_.length(0)):
                    nw512_[i0_89_] = init89_(i0_89_)
                d_2586_v36_ = nw512_
                d_2594_v37_: _dafny.Set
                d_2594_v37_ = _dafny.Set({p1})
                d_2595_v38_: _dafny.Seq
                d_2595_v38_ = _dafny.SeqWithoutIsStrInference([d_2594_v37_, d_2594_v37_, d_2594_v37_])
                d_2596_v39_: _dafny.Seq
                d_2596_v39_ = d_2595_v38_
                index443_ = default__.safeIndex(800, (d_2586_v36_).length(0))
                (d_2586_v36_)[index443_] = d_2596_v39_
                index444_ = default__.safeIndex(800, (d_2586_v36_).length(0))
                (d_2586_v36_)[index444_] = (d_2596_v39_ if (p2) >= ((0) - (d_2583_cf52_.f14)) else d_2596_v39_)
            elif source39_.is_DC32:
                d_2597___mcc_h7_ = source39_.cf46
                d_2598_cf46_ = d_2597___mcc_h7_
                d_2599_v40_: _dafny.Array
                def lambda152_(d_2600_p0_):
                    def lambda153_(d_2601_i9_):
                        return (d_2601_i9_) - (len(_dafny.SeqWithoutIsStrInference([d_2600_p0_])))

                    return lambda153_

                init90_ = lambda152_(p0)
                nw513_ = _dafny.Array(None, 20)
                for i0_90_ in range(nw513_.length(0)):
                    nw513_[i0_90_] = init90_(i0_90_)
                d_2599_v40_ = nw513_
                d_2602_v41_: _dafny.Array
                nw514_ = _dafny.Array(False, 14)
                d_2602_v41_ = nw514_
                d_2603_v42_: D9
                d_2603_v42_ = D9_DC15(p0, p0, default__.fm49(p1, p2, p0, d_2560_v26_.f14, globalState), d_2599_v40_, d_2602_v41_)
                d_2604_v43_: _dafny.Seq
                d_2604_v43_ = _dafny.SeqWithoutIsStrInference([d_2599_v40_])
                d_2605_v44_: _dafny.Set
                d_2605_v44_ = _dafny.Set({d_2599_v40_, d_2599_v40_, (d_2604_v43_)[default__.safeIndex(len((d_2555_v22_).set(default__.safeIndex(p3, len(d_2555_v22_)), p0)), len(d_2604_v43_))], d_2599_v40_})
                d_2606_v45_: _dafny.Array
                nw515_ = _dafny.Array(None, 3)
                nw515_[int(0)] = (d_2603_v42_).cf18
                nw515_[int(1)] = not((_dafny.Set({d_2599_v40_, d_2599_v40_})).isdisjoint(d_2605_v44_))
                nw515_[int(2)] = p0
                d_2606_v45_ = nw515_
                index445_ = default__.safeIndex(519, (d_2602_v41_).length(0))
                (d_2602_v41_)[index445_] = (d_2558_i7_) > (d_2560_v26_.f14)
                index446_ = default__.safeIndex(519, (d_2602_v41_).length(0))
                (d_2602_v41_)[index446_] = p0
                (globalState).f2 = (0) - (d_2560_v26_.f14)
                d_2607_v46_: C3
                nw516_ = C3()
                nw516_.ctor__(p2, (d_2524_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bomh"))), (d_2602_v41_)[default__.safeIndex(519, (d_2602_v41_).length(0))], (d_2602_v41_)[default__.safeIndex(519, (d_2602_v41_).length(0))])
                d_2607_v46_ = nw516_
                index447_ = default__.safeIndex(519, (d_2602_v41_).length(0))
                (d_2602_v41_)[index447_] = (default__.fm31(p0, False, globalState) if False else p0)
            elif True:
                d_2608___mcc_h8_ = source39_.cf54
                d_2609_cf54_ = d_2608___mcc_h8_
                d_2610_v47_: int
                d_2611_v48_: bool
                out36_: int
                out37_: bool
                out36_, out37_ = (self).m5(globalState)
                d_2610_v47_ = out36_
                d_2611_v48_ = out37_
                d_2612_v49_: _dafny.Array
                def lambda154_(d_2613_p1_):
                    def lambda155_(d_2614_i10_):
                        return default__.safeModuloInt(d_2614_i10_, d_2613_p1_)

                    return lambda155_

                init91_ = lambda154_(p1)
                nw517_ = _dafny.Array(None, 22)
                for i0_91_ in range(nw517_.length(0)):
                    nw517_[i0_91_] = init91_(i0_91_)
                d_2612_v49_ = nw517_
                index448_ = default__.safeIndex(740, (d_2612_v49_).length(0))
                (d_2612_v49_)[index448_] = p2
                index449_ = default__.safeIndex(740, (d_2612_v49_).length(0))
                (d_2612_v49_)[index449_] = d_2558_i7_
                d_2615_v50_: _dafny.MultiSet
                d_2615_v50_ = _dafny.MultiSet([d_2560_v26_.f14, p1])
                d_2615_v50_ = ((d_2615_v50_).set((d_2612_v49_)[default__.safeIndex(740, (d_2612_v49_).length(0))], default__.abs(d_2560_v26_.f14))).intersection(d_2615_v50_)
                (globalState).f2 = ((d_2615_v50_)[d_2560_v26_.f14] if (d_2560_v26_.f14) in (d_2615_v50_) else d_2560_v26_.f14)
            d_2616_v51_: bool
            d_2616_v51_ = False
            d_2617_v52_: _dafny.MultiSet
            d_2617_v52_ = _dafny.MultiSet([d_2558_i7_, d_2558_i7_, d_2558_i7_, len((d_2526_v2_.f18).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2559_v25_ for d_2618_i11_ in range(default__.abs(46))])), len(d_2526_v2_.f18)), d_2559_v25_)), p2])
            rhs470_ = ((p1) * (d_2560_v26_.f14)) == ((p1) + (p2))
            rhs471_ = not(not((d_2617_v52_).isdisjoint(d_2617_v52_)))
            rhs472_ = p0
            d_2616_v51_ = rhs470_
            d_2616_v51_ = rhs471_
            d_2616_v51_ = rhs472_
            if d_2616_v51_:
                d_2619_v53_: D3
                d_2619_v53_ = D3_DC4(p0, p2, p2, p0, d_2616_v51_)
                d_2620_v54_: _dafny.Map
                d_2620_v54_ = _dafny.Map({p0: d_2616_v51_})
                d_2621_v55_: _dafny.Map
                d_2621_v55_ = _dafny.Map({(0) - ((d_2526_v2_).fm12(len(_dafny.Map({-952: d_2560_v26_.f14})), globalState)): d_2526_v2_.f18})
                d_2622_v56_: _dafny.Map
                d_2622_v56_ = _dafny.Map({d_2560_v26_.f14: d_2560_v26_.f14})
                d_2623_v57_: _dafny.Seq
                d_2623_v57_ = _dafny.SeqWithoutIsStrInference([d_2524_v0_])
                d_2624_v58_: _dafny.Array
                nw518_ = _dafny.Array(None, 19)
                nw518_[int(0)] = d_2526_v2_.f18
                nw518_[int(1)] = default__.fm28(p0, _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faf"))): p3}), d_2619_v53_, d_2620_v54_, globalState)
                nw518_[int(2)] = d_2526_v2_.f18
                nw518_[int(3)] = d_2524_v0_
                nw518_[int(4)] = d_2526_v2_.f18
                nw518_[int(5)] = ((d_2526_v2_.f18).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_2617_v52_).cardinality])), len(d_2526_v2_.f18)), _dafny.CodePoint('c'))) + (d_2526_v2_.f18)
                nw518_[int(6)] = d_2526_v2_.f18
                nw518_[int(7)] = (d_2526_v2_.f18) + (d_2524_v0_)
                nw518_[int(8)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecxp"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecxp")))), d_2559_v25_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qoltf")))
                nw518_[int(9)] = d_2524_v0_
                nw518_[int(10)] = (d_2526_v2_.f18) + (((d_2621_v55_)[p1] if (p1) in (d_2621_v55_) else (d_2526_v2_.f18).set(default__.safeIndex(p1, len(d_2526_v2_.f18)), d_2559_v25_)))
                nw518_[int(11)] = d_2524_v0_
                nw518_[int(12)] = d_2526_v2_.f18
                nw518_[int(13)] = d_2526_v2_.f18
                nw518_[int(14)] = default__.fm28(p0, d_2622_v56_, d_2619_v53_, d_2620_v54_, globalState)
                nw518_[int(15)] = d_2526_v2_.f18
                nw518_[int(16)] = _dafny.SeqWithoutIsStrInference([d_2559_v25_ for d_2625_i12_ in range(default__.abs(-673))])
                nw518_[int(17)] = d_2526_v2_.f18
                nw518_[int(18)] = (d_2623_v57_)[default__.safeIndex(d_2558_i7_, len(d_2623_v57_))]
                d_2624_v58_ = nw518_
                d_2626_v59_: _dafny.Map
                d_2626_v59_ = _dafny.Map({not(d_2616_v51_): d_2620_v54_})
                d_2627_v60_: _dafny.Map
                d_2627_v60_ = _dafny.Map({d_2616_v51_: d_2626_v59_})
                d_2628_v61_: _dafny.Seq
                d_2628_v61_ = _dafny.SeqWithoutIsStrInference([d_2627_v60_])
                d_2629_v62_: _dafny.Map
                d_2629_v62_ = _dafny.Map({(d_2628_v61_)[default__.safeIndex((0) - ((d_2526_v2_).fm12(len(_dafny.SeqWithoutIsStrInference([d_2559_v25_ for d_2630_i13_ in range(default__.abs(528))])), globalState)), len(d_2628_v61_))]: d_2558_i7_})
                d_2631_v63_: D14
                d_2631_v63_ = D14_DC27(d_2629_v62_)
                rhs473_ = d_2624_v58_
                rhs474_ = d_2631_v63_
                d_2624_v58_ = rhs473_
                d_2631_v63_ = rhs474_
                d_2632_v64_: _dafny.Array
                def lambda156_(d_2633_v1_):
                    def lambda157_(d_2634_i14_):
                        return d_2633_v1_

                    return lambda157_

                init92_ = lambda156_(d_2525_v1_)
                nw519_ = _dafny.Array(None, 2)
                for i0_92_ in range(nw519_.length(0)):
                    nw519_[i0_92_] = init92_(i0_92_)
                d_2632_v64_ = nw519_
                d_2635_v65_: C0
                nw520_ = C0()
                nw520_.ctor__(d_2632_v64_, -759)
                d_2635_v65_ = nw520_
                nw521_ = C8()
                nw521_.ctor__((d_2623_v57_)[default__.safeIndex(d_2560_v26_.f14, len(d_2623_v57_))], (d_2560_v26_.f14) == ((0) - ((0) - (p3))), p0, d_2558_i7_)
                d_2526_v2_ = nw521_
                (d_2560_v26_).f14 = p2
                d_2559_v25_ = _dafny.CodePoint('t')
            elif True:
                d_2636_v66_: _dafny.Map
                d_2636_v66_ = _dafny.Map({p0: p2})
                d_2637_v67_: D18
                d_2637_v67_ = D18_DC38(len(d_2636_v66_), d_2616_v51_, d_2558_i7_)
                d_2638_v68_: C5
                nw522_ = C5()
                nw522_.ctor__((d_2637_v67_).cf58, d_2616_v51_)
                d_2638_v68_ = nw522_
                d_2639_v69_: C8
                nw523_ = C8()
                nw523_.ctor__(d_2526_v2_.f18, not(d_2616_v51_), p0, p3)
                d_2639_v69_ = nw523_
                d_2640_v70_: _dafny.Map
                d_2640_v70_ = _dafny.Map({d_2561_v27_: d_2560_v26_.f14})
                d_2640_v70_ = d_2640_v70_
                d_2616_v51_ = not(False)
                d_2641_v71_: int
                out38_: int
                out38_ = (d_2526_v2_).m3(d_2559_v25_, len(d_2555_v22_), False, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iir"))) + (_dafny.SeqWithoutIsStrInference([d_2559_v25_ for d_2642_i15_ in range(default__.abs(422))]))), globalState)
                d_2641_v71_ = out38_
        d_2643_v72_: _dafny.Map
        d_2643_v72_ = _dafny.Map({-978: p2})
        d_2644_v73_: _dafny.Map
        d_2644_v73_ = _dafny.Map({p0: p0})
        d_2645_i16_: int
        d_2645_i16_ = 0
        with _dafny.label("17"):
            while (p3) != (len((d_2643_v72_).set(len(d_2644_v73_), p3))):
                with _dafny.c_label("17"):
                    if (d_2645_i16_) >= (100):
                        raise _dafny.Break("17")
                    d_2645_i16_ = (d_2645_i16_) + (1)
                    d_2646_v74_: D15
                    d_2646_v74_ = D15_DC29(d_2525_v1_)
                    d_2647_v75_: C6
                    nw524_ = C6()
                    nw524_.ctor__(len(d_2526_v2_.f18), p0, p0, p0, len((d_2646_v74_).cf43))
                    d_2647_v75_ = nw524_
                    (globalState).f2 = (d_2647_v75_).f19
                    d_2648_v76_: _dafny.Array
                    nw525_ = _dafny.Array(_dafny.Array(None, 0), 13)
                    d_2648_v76_ = nw525_
                    d_2648_v76_ = d_2648_v76_
                    d_2649_v77_: _dafny.MultiSet
                    d_2649_v77_ = _dafny.MultiSet([(d_2647_v75_).f20, (d_2647_v75_).f20, (d_2647_v75_).f20])
                    d_2650_v78_: _dafny.Array
                    nw526_ = _dafny.Array(None, 5)
                    nw526_[int(0)] = (len(d_2555_v22_)) > ((d_2647_v75_).f19)
                    nw526_[int(1)] = p0
                    nw526_[int(2)] = not ((d_2647_v75_).f20) or (default__.fm17(p0, (d_2647_v75_).f19, p3, (0) - (p1), globalState))
                    nw526_[int(3)] = p0
                    nw526_[int(4)] = (d_2649_v77_).issubset(d_2649_v77_)
                    d_2650_v78_ = nw526_
                    index450_ = default__.safeIndex(85, (d_2650_v78_).length(0))
                    (d_2650_v78_)[index450_] = p0
                    d_2651_v79_: _dafny.Seq
                    d_2651_v79_ = _dafny.SeqWithoutIsStrInference([p1, (d_2649_v77_).cardinality, p1])
                    index451_ = default__.safeIndex(85, (d_2650_v78_).length(0))
                    (d_2650_v78_)[index451_] = default__.fm27(len(d_2651_v79_), p3, p2, d_2643_v72_, globalState)
                    pass
            pass
        d_2652_v81_: _dafny.Map
        d_2652_v81_ = _dafny.Map({p3: p0})
        d_2653_v83_: _dafny.Seq
        def iife197_():
            coll87_ = _dafny.Set()
            compr_87_: int
            for compr_87_ in (d_2652_v81_).keys.Elements:
                d_2654_v82_: int = compr_87_
                if (d_2654_v82_) in (d_2652_v81_):
                    coll87_ = coll87_.union(_dafny.Set([(d_2654_v82_) + (len(_dafny.SeqWithoutIsStrInference([-594, 445])))]))
            return _dafny.Set(coll87_)
        d_2653_v83_ = _dafny.SeqWithoutIsStrInference([len(iife197_()
        ), p3])
        def iife198_():
            coll88_ = _dafny.Map()
            compr_88_: int
            for compr_88_ in (d_2653_v83_).Elements:
                d_2655_v80_: int = compr_88_
                if (d_2655_v80_) in (d_2653_v83_):
                    coll88_[(d_2655_v80_) - (len(_dafny.Set({p2})))] = (0) - (p1)
            return _dafny.Map(coll88_)
        r0 = (p1 if (iife198_()
        ) != (d_2643_v72_) else p3)
        r1 = default__.safeModuloInt(len(d_2653_v83_), p2)
        return r0, r1

    def m5(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2656_v0_: int
        d_2656_v0_ = 205
        r0 = d_2656_v0_
        hi24_ = -132
        for d_2657_i0_ in range(d_2656_v0_, hi24_):
            d_2658_v1_: _dafny.Array
            def lambda158_(d_2659_v0_):
                def lambda159_(d_2660_i1_):
                    return (_dafny.Set({454})).isdisjoint(_dafny.Set({d_2659_v0_}))

                return lambda159_

            init93_ = lambda158_(d_2656_v0_)
            nw527_ = _dafny.Array(None, 23)
            for i0_93_ in range(nw527_.length(0)):
                nw527_[i0_93_] = init93_(i0_93_)
            d_2658_v1_ = nw527_
            d_2661_v2_: bool
            d_2661_v2_ = False
            index452_ = default__.safeIndex(717, (d_2658_v1_).length(0))
            (d_2658_v1_)[index452_] = d_2661_v2_
            d_2662_v3_: _dafny.Seq
            d_2662_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wluvi"))
            d_2663_v4_: _dafny.Map
            d_2663_v4_ = _dafny.Map({d_2662_v3_: -752})
            d_2664_v6_: _dafny.Seq
            d_2664_v6_ = _dafny.SeqWithoutIsStrInference([d_2662_v3_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2665_i2_ in range(default__.abs(511))])])
            d_2666_v9_: _dafny.Map
            d_2666_v9_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugrshctcj")): False})
            d_2667_v10_: _dafny.Map
            d_2667_v10_ = _dafny.Map({d_2661_v2_: d_2661_v2_})
            d_2668_v11_: _dafny.Map
            d_2668_v11_ = _dafny.Map({d_2667_v10_: d_2663_v4_})
            d_2669_v12_: _dafny.Seq
            d_2669_v12_ = _dafny.SeqWithoutIsStrInference([d_2663_v4_])
            d_2670_v13_: _dafny.Array
            nw528_ = _dafny.Array(None, 28)
            nw528_[int(0)] = d_2663_v4_
            nw528_[int(1)] = d_2663_v4_
            nw528_[int(2)] = _dafny.Map({d_2662_v3_: d_2657_i0_})
            nw528_[int(3)] = (d_2663_v4_) | (d_2663_v4_)
            nw528_[int(4)] = d_2663_v4_
            nw528_[int(5)] = d_2663_v4_
            nw528_[int(6)] = d_2663_v4_
            nw528_[int(7)] = d_2663_v4_
            def iife199_():
                coll89_ = _dafny.Map()
                compr_89_: _dafny.Seq
                for compr_89_ in (d_2664_v6_).Elements:
                    d_2671_v5_: _dafny.Seq = compr_89_
                    if (d_2671_v5_) in (d_2664_v6_):
                        coll89_[d_2671_v5_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwkk")))
                return _dafny.Map(coll89_)
            nw528_[int(8)] = (d_2663_v4_) | (iife199_()
            )
            nw528_[int(9)] = (d_2663_v4_).set(d_2662_v3_, d_2657_i0_)
            nw528_[int(10)] = (d_2663_v4_) | (d_2663_v4_)
            nw528_[int(11)] = (d_2663_v4_) | ((d_2663_v4_).set(d_2662_v3_, d_2656_v0_))
            def iife200_():
                coll90_ = _dafny.Map()
                compr_90_: _dafny.Seq
                for compr_90_ in (d_2664_v6_).Elements:
                    d_2672_v7_: _dafny.Seq = compr_90_
                    if (d_2672_v7_) in (d_2664_v6_):
                        coll90_[d_2672_v7_] = (_dafny.MultiSet([d_2657_i0_, len(d_2662_v3_)])).cardinality
                return _dafny.Map(coll90_)
            def iife201_():
                coll91_ = _dafny.Map()
                compr_91_: _dafny.Seq
                for compr_91_ in (d_2666_v9_).keys.Elements:
                    d_2673_v8_: _dafny.Seq = compr_91_
                    if (d_2673_v8_) in (d_2666_v9_):
                        coll91_[d_2673_v8_] = d_2657_i0_
                return _dafny.Map(coll91_)
            nw528_[int(12)] = (iife200_()
            ) | (iife201_()
            )
            nw528_[int(13)] = default__.fm69(d_2656_v0_, globalState)
            nw528_[int(14)] = d_2663_v4_
            nw528_[int(15)] = _dafny.Map({d_2662_v3_: default__.fm0(d_2661_v2_, globalState)})
            nw528_[int(16)] = (d_2663_v4_) | (d_2663_v4_)
            nw528_[int(17)] = (d_2663_v4_) | (d_2663_v4_)
            nw528_[int(18)] = (d_2663_v4_ if d_2661_v2_ else d_2663_v4_)
            nw528_[int(19)] = (d_2663_v4_) | (d_2663_v4_)
            nw528_[int(20)] = d_2663_v4_
            nw528_[int(21)] = d_2663_v4_
            nw528_[int(22)] = d_2663_v4_
            nw528_[int(23)] = d_2663_v4_
            nw528_[int(24)] = ((d_2668_v11_)[_dafny.Map({d_2661_v2_: d_2661_v2_})] if (_dafny.Map({d_2661_v2_: d_2661_v2_})) in (d_2668_v11_) else d_2663_v4_)
            nw528_[int(25)] = d_2663_v4_
            nw528_[int(26)] = d_2663_v4_
            nw528_[int(27)] = (d_2669_v12_)[default__.safeIndex(-890, len(d_2669_v12_))]
            d_2670_v13_ = nw528_
            index453_ = default__.safeIndex(571, (d_2670_v13_).length(0))
            (d_2670_v13_)[index453_] = (d_2669_v12_)[default__.safeIndex(d_2657_i0_, len(d_2669_v12_))]
            d_2674_v14_: _dafny.MultiSet
            d_2674_v14_ = _dafny.MultiSet([d_2661_v2_])
            d_2675_v15_: _dafny.Set
            d_2675_v15_ = _dafny.Set({d_2657_i0_, 929})
            d_2676_v16_: _dafny.Seq
            d_2676_v16_ = _dafny.SeqWithoutIsStrInference([d_2656_v0_, len(d_2675_v15_)])
            d_2677_v17_: _dafny.Seq
            d_2677_v17_ = _dafny.SeqWithoutIsStrInference([False, True])
            d_2678_v18_: _dafny.Map
            d_2678_v18_ = _dafny.Map({default__.fm0(not(d_2661_v2_), globalState): d_2656_v0_})
            d_2679_v19_: _dafny.Array
            nw529_ = _dafny.Array(None, 24)
            nw529_[int(0)] = (0) - (d_2656_v0_)
            nw529_[int(1)] = ((d_2674_v14_)[d_2661_v2_] if (d_2661_v2_) in (d_2674_v14_) else d_2656_v0_)
            nw529_[int(2)] = d_2656_v0_
            nw529_[int(3)] = (_dafny.MultiSet([d_2657_i0_, d_2656_v0_])).cardinality
            nw529_[int(4)] = 181
            nw529_[int(5)] = (d_2674_v14_).cardinality
            nw529_[int(6)] = d_2657_i0_
            nw529_[int(7)] = d_2657_i0_
            nw529_[int(8)] = d_2656_v0_
            nw529_[int(9)] = (d_2676_v16_)[default__.safeIndex(d_2656_v0_, len(d_2676_v16_))]
            nw529_[int(10)] = len(d_2677_v17_)
            nw529_[int(11)] = d_2656_v0_
            nw529_[int(12)] = d_2656_v0_
            nw529_[int(13)] = len(d_2678_v18_)
            nw529_[int(14)] = len(d_2662_v3_)
            nw529_[int(15)] = (0) - (len(d_2676_v16_))
            nw529_[int(16)] = 568
            nw529_[int(17)] = 188
            nw529_[int(18)] = d_2657_i0_
            nw529_[int(19)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
            nw529_[int(20)] = len(_dafny.Map({d_2657_i0_: (0) - (d_2656_v0_)}))
            nw529_[int(21)] = d_2656_v0_
            nw529_[int(22)] = (d_2676_v16_)[default__.safeIndex(d_2656_v0_, len(d_2676_v16_))]
            nw529_[int(23)] = (0) - (d_2656_v0_)
            d_2679_v19_ = nw529_
            d_2680_v20_: _dafny.Map
            d_2680_v20_ = _dafny.Map({d_2679_v19_: (d_2677_v17_) + ((default__.fm52(True, d_2656_v0_, d_2656_v0_, 229, globalState)).set(default__.safeIndex(d_2657_i0_, len(default__.fm52(True, d_2656_v0_, d_2656_v0_, 229, globalState))), True))})
            index454_ = default__.safeIndex(717, (d_2658_v1_).length(0))
            index455_ = default__.safeIndex(571, (d_2670_v13_).length(0))
            rhs475_ = (_dafny.MultiSet([d_2661_v2_])).issubset(d_2674_v14_)
            rhs476_ = d_2657_i0_
            rhs477_ = (d_2669_v12_)[default__.safeIndex(d_2657_i0_, len(d_2669_v12_))]
            rhs478_ = (d_2657_i0_) + (d_2657_i0_)
            rhs479_ = d_2680_v20_
            lhs368_ = d_2658_v1_
            lhs369_ = default__.safeIndex(717, (d_2658_v1_).length(0))
            lhs370_ = globalState
            lhs371_ = d_2670_v13_
            lhs372_ = default__.safeIndex(571, (d_2670_v13_).length(0))
            lhs368_[lhs369_] = rhs475_
            lhs370_.f7 = rhs476_
            lhs371_[lhs372_] = rhs477_
            d_2656_v0_ = rhs478_
            d_2680_v20_ = rhs479_
            d_2681_v21_: _dafny.Array
            def lambda160_(d_2682_v3_):
                def lambda161_(d_2683_i3_):
                    return d_2682_v3_

                return lambda161_

            init94_ = lambda160_(d_2662_v3_)
            nw530_ = _dafny.Array(None, 3)
            for i0_94_ in range(nw530_.length(0)):
                nw530_[i0_94_] = init94_(i0_94_)
            d_2681_v21_ = nw530_
            d_2684_v22_: _dafny.Map
            d_2684_v22_ = _dafny.Map({d_2681_v21_: (len(d_2677_v17_)) - (d_2657_i0_)})
            d_2684_v22_ = (d_2684_v22_).set(d_2681_v21_, default__.safeModuloInt((0) - (d_2656_v0_), d_2656_v0_))
            index456_ = default__.safeIndex(717, (d_2658_v1_).length(0))
            (d_2658_v1_)[index456_] = not((d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))])
            if (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))]:
                d_2685_v23_: D19
                d_2685_v23_ = D19_DC39(_dafny.MultiSet([d_2656_v0_]))
                d_2686_v24_: _dafny.MultiSet
                d_2686_v24_ = _dafny.MultiSet([d_2657_i0_])
                d_2687_v25_: _dafny.Map
                d_2687_v25_ = _dafny.Map({d_2657_i0_: (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))]})
                d_2688_v26_: D7
                d_2688_v26_ = D7_DC11(d_2687_v25_)
                pat_let_tv43_ = d_2686_v24_
                d_2689_v27_: _dafny.Map
                def iife202_(_pat_let55_0):
                    def iife203_(d_2690_dt__update__tmp_h0_):
                        def iife204_(_pat_let56_0):
                            def iife205_(d_2691_dt__update_hcf60_h0_):
                                return D19_DC39(d_2691_dt__update_hcf60_h0_)
                            return iife205_(_pat_let56_0)
                        return iife204_(pat_let_tv43_)
                    return iife203_(_pat_let55_0)
                d_2689_v27_ = _dafny.Map({iife202_(d_2685_v23_): d_2688_v26_})
                pat_let_tv44_ = d_2686_v24_
                def iife206_(_pat_let57_0):
                    def iife207_(d_2692_dt__update__tmp_h1_):
                        def iife208_(_pat_let58_0):
                            def iife209_(d_2693_dt__update_hcf60_h1_):
                                return D19_DC39(d_2693_dt__update_hcf60_h1_)
                            return iife209_(_pat_let58_0)
                        return iife208_(pat_let_tv44_)
                    return iife207_(_pat_let57_0)
                d_2689_v27_ = (d_2689_v27_).set(iife206_(d_2685_v23_), d_2688_v26_)
                r0 = ((0) - ((d_2674_v14_).cardinality)) + (d_2656_v0_)
                (globalState).f7 = d_2656_v0_
                d_2694_v28_: D3
                d_2694_v28_ = D3_DC4((d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))], d_2656_v0_, d_2656_v0_, d_2661_v2_, (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))])
                d_2662_v3_ = default__.fm28(True, d_2678_v18_, d_2694_v28_, d_2667_v10_, globalState)
                (globalState).f7 = default__.safeModuloInt((d_2674_v14_).cardinality, default__.fm0(d_2661_v2_, globalState))
            elif True:
                index457_ = default__.safeIndex(971, (d_2679_v19_).length(0))
                (d_2679_v19_)[index457_] = d_2657_i0_
                index458_ = default__.safeIndex(971, (d_2679_v19_).length(0))
                (d_2679_v19_)[index458_] = len(d_2662_v3_)
                r1 = (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))]
                d_2695_v29_: _dafny.Array
                nw531_ = _dafny.Array(_dafny.Set({}), 5)
                d_2695_v29_ = nw531_
                d_2696_v30_: _dafny.Set
                d_2696_v30_ = _dafny.Set({(d_2677_v17_)[default__.safeIndex((d_2679_v19_)[default__.safeIndex(971, (d_2679_v19_).length(0))], len(d_2677_v17_))], (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))], d_2661_v2_, d_2661_v2_, not((d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))])})
                index459_ = default__.safeIndex(66, (d_2695_v29_).length(0))
                (d_2695_v29_)[index459_] = d_2696_v30_
                index460_ = default__.safeIndex(66, (d_2695_v29_).length(0))
                (d_2695_v29_)[index460_] = d_2696_v30_
                index461_ = default__.safeIndex(971, (d_2679_v19_).length(0))
                rhs480_ = (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))]
                rhs481_ = (d_2677_v17_)[default__.safeIndex(default__.safeDivisionInt(len(d_2675_v15_), (d_2679_v19_)[default__.safeIndex(971, (d_2679_v19_).length(0))]), len(d_2677_v17_))]
                rhs482_ = (d_2662_v3_) != (d_2662_v3_)
                rhs483_ = default__.fm0((d_2661_v2_ if True else (d_2658_v1_)[default__.safeIndex(717, (d_2658_v1_).length(0))]), globalState)
                lhs373_ = d_2679_v19_
                lhs374_ = default__.safeIndex(971, (d_2679_v19_).length(0))
                r1 = rhs480_
                d_2661_v2_ = rhs481_
                d_2661_v2_ = rhs482_
                lhs373_[lhs374_] = rhs483_
                (globalState).f7 = (0) - (d_2656_v0_)
        d_2697_v31_: str
        d_2697_v31_ = _dafny.CodePoint('t')
        d_2698_v32_: _dafny.Map
        d_2698_v32_ = _dafny.Map({d_2697_v31_: d_2656_v0_})
        d_2699_v33_: bool
        d_2699_v33_ = False
        d_2700_v34_: _dafny.Seq
        d_2700_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpujsfs"))
        d_2701_v35_: _dafny.Map
        d_2701_v35_ = _dafny.Map({d_2656_v0_: len(d_2700_v34_)})
        d_2702_v36_: T0
        nw532_ = C1()
        nw532_.ctor__(d_2699_v33_, d_2656_v0_)
        d_2702_v36_ = nw532_
        d_2703_v37_: D20
        d_2703_v37_ = D20_DC42(len(_dafny.SeqWithoutIsStrInference([False])), d_2699_v33_, d_2702_v36_)
        d_2704_v38_: T1
        nw533_ = C6()
        nw533_.ctor__(d_2656_v0_, d_2699_v33_, (d_2703_v37_).cf66, d_2699_v33_, d_2702_v36_.f14)
        d_2704_v38_ = nw533_
        d_2705_v39_: _dafny.Seq
        d_2705_v39_ = _dafny.SeqWithoutIsStrInference([d_2704_v38_])
        d_2706_v40_: _dafny.Map
        d_2706_v40_ = _dafny.Map({d_2704_v38_.f16: d_2699_v33_})
        d_2698_v32_ = (d_2698_v32_).set(d_2697_v31_, len(default__.fm28(d_2699_v33_, d_2701_v35_, D3_DC4(True, len(d_2705_v39_), d_2656_v0_, d_2699_v33_, d_2704_v38_.f16), d_2706_v40_, globalState)))
        d_2707_v42_: _dafny.Array
        def lambda162_(d_2708_v38_):
            def lambda163_(d_2709_i4_):
                def iife210_():
                    coll92_ = _dafny.Set()
                    compr_92_: _dafny.Seq
                    for compr_92_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_2708_v38_.f16, True]): (d_2708_v38_).f17})).keys.Elements:
                        d_2710_v41_: _dafny.Seq = compr_92_
                        if (d_2710_v41_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_2708_v38_.f16, True]): (d_2708_v38_).f17})):
                            coll92_ = coll92_.union(_dafny.Set([d_2710_v41_]))
                    return _dafny.Set(coll92_)
                return default__.safeModuloInt(d_2709_i4_, len(iife210_()
                ))

            return lambda163_

        init95_ = lambda162_(d_2704_v38_)
        nw534_ = _dafny.Array(None, 20)
        for i0_95_ in range(nw534_.length(0)):
            nw534_[i0_95_] = init95_(i0_95_)
        d_2707_v42_ = nw534_
        d_2707_v42_ = d_2707_v42_
        d_2711_v43_: _dafny.Array
        nw535_ = _dafny.Array(_dafny.Seq({}), 28)
        d_2711_v43_ = nw535_
        d_2712_v44_: _dafny.Seq
        d_2712_v44_ = _dafny.SeqWithoutIsStrInference([True])
        index462_ = default__.safeIndex(754, (d_2711_v43_).length(0))
        (d_2711_v43_)[index462_] = d_2712_v44_
        index463_ = default__.safeIndex(754, (d_2711_v43_).length(0))
        (d_2711_v43_)[index463_] = d_2712_v44_
        hi25_ = d_2702_v36_.f14
        for d_2713_i5_ in range(d_2702_v36_.f14, hi25_):
            d_2714_v45_: _dafny.Array
            nw536_ = _dafny.Array(None, 9)
            nw536_[int(0)] = not((d_2704_v38_).f17)
            nw536_[int(1)] = (d_2704_v38_).f17
            nw536_[int(2)] = d_2704_v38_.f16
            nw536_[int(3)] = False
            nw536_[int(4)] = not(not(d_2704_v38_.f16))
            nw536_[int(5)] = (d_2702_v36_.f14) < (d_2656_v0_)
            nw536_[int(6)] = (d_2697_v31_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgmbqxi")))
            nw536_[int(7)] = d_2704_v38_.f16
            nw536_[int(8)] = (d_2704_v38_.f16) and (d_2699_v33_)
            d_2714_v45_ = nw536_
            d_2714_v45_ = d_2714_v45_
            d_2715_v46_: _dafny.Set
            d_2715_v46_ = _dafny.Set({d_2704_v38_.f16})
            d_2716_v47_: _dafny.Map
            d_2716_v47_ = _dafny.Map({d_2713_i5_: d_2715_v46_})
            d_2717_v48_: _dafny.Array
            nw537_ = _dafny.Array(None, 20)
            nw537_[int(0)] = d_2715_v46_
            nw537_[int(1)] = ((d_2716_v47_)[d_2713_i5_] if (d_2713_i5_) in (d_2716_v47_) else d_2715_v46_)
            nw537_[int(2)] = d_2715_v46_
            nw537_[int(3)] = d_2715_v46_
            nw537_[int(4)] = default__.fm44((d_2704_v38_).f17, d_2699_v33_, globalState)
            nw537_[int(5)] = default__.fm44(not(d_2704_v38_.f16), d_2699_v33_, globalState)
            nw537_[int(6)] = d_2715_v46_
            nw537_[int(7)] = d_2715_v46_
            nw537_[int(8)] = d_2715_v46_
            nw537_[int(9)] = d_2715_v46_
            nw537_[int(10)] = _dafny.Set({d_2699_v33_, not((d_2704_v38_).f17), d_2699_v33_})
            nw537_[int(11)] = _dafny.Set({(d_2704_v38_).f17, True})
            nw537_[int(12)] = d_2715_v46_
            nw537_[int(13)] = d_2715_v46_
            nw537_[int(14)] = d_2715_v46_
            nw537_[int(15)] = _dafny.Set({d_2704_v38_.f16, d_2704_v38_.f16, (d_2704_v38_).f17})
            nw537_[int(16)] = d_2715_v46_
            nw537_[int(17)] = d_2715_v46_
            nw537_[int(18)] = d_2715_v46_
            nw537_[int(19)] = _dafny.Set({False, (d_2704_v38_).f17})
            d_2717_v48_ = nw537_
            d_2718_v49_: C0
            nw538_ = C0()
            nw538_.ctor__(d_2717_v48_, d_2713_i5_)
            d_2718_v49_ = nw538_
            (d_2702_v36_).f14 = 101
            r1 = not (d_2704_v38_.f16) or (not((d_2702_v36_.f14) <= (d_2713_i5_)))
        r0 = (d_2656_v0_) * (d_2702_v36_.f14)
        r1 = ((d_2711_v43_)[default__.safeIndex(754, (d_2711_v43_).length(0))])[default__.safeIndex(d_2702_v36_.f14, len((d_2711_v43_)[default__.safeIndex(754, (d_2711_v43_).length(0))]))]
        return r0, r1


class C10(T1, T0):
    def  __init__(self):
        self._f16: bool = False
        self._f14: int = int(0)
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f16, f17, f14):
        (self).f16 = f16
        (self)._f17 = f17
        (self).f14 = f14

    def fm3(self, globalState):
        return ((_dafny.Map({self.f14: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgtxp")))})) | (_dafny.Map({self.f14: self.f14}))) | ((_dafny.Map({(0) - ((0) - (self.f14)): self.f14}) if (self).f17 else _dafny.Map({self.f14: len(_dafny.SeqWithoutIsStrInference([(self).f17]))})))

    def fm9(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsvqcphrb")) if self.f16 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txojgxj")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2719_i0_ in range(default__.abs(-432))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_2720_i1_ in range(default__.abs(348))])))

    def fm10(self, p0, p1, p2, globalState):
        return self.f16

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        if (self).f17:
            d_2721_v0_: str
            d_2721_v0_ = _dafny.CodePoint('h')
            d_2722_v1_: _dafny.MultiSet
            d_2722_v1_ = _dafny.MultiSet([d_2721_v0_])
            d_2723_v2_: _dafny.Seq
            d_2723_v2_ = _dafny.SeqWithoutIsStrInference([d_2722_v1_])
            d_2723_v2_ = ((d_2723_v2_).set(default__.safeIndex(p0, len(d_2723_v2_)), d_2722_v1_)).set(default__.safeIndex(p0, len((d_2723_v2_).set(default__.safeIndex(p0, len(d_2723_v2_)), d_2722_v1_))), (d_2723_v2_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2721_v0_ for d_2724_i0_ in range(default__.abs(688))])), len(d_2723_v2_))])
            d_2725_v3_: _dafny.Seq
            d_2725_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdwbsv"))
            d_2726_v4_: _dafny.Map
            d_2726_v4_ = _dafny.Map({p0: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scqd"))) + (d_2725_v3_)})
            d_2726_v4_ = d_2726_v4_
            d_2727_v5_: int
            d_2727_v5_ = self.f14
            d_2728_v6_: _dafny.Map
            d_2728_v6_ = _dafny.Map({False: p0})
            d_2729_v7_: _dafny.MultiSet
            d_2729_v7_ = _dafny.MultiSet([len(d_2728_v6_)])
            d_2730_v8_: _dafny.Set
            d_2730_v8_ = _dafny.Set({p1, (self).f17})
            d_2731_v9_: _dafny.Seq
            d_2731_v9_ = _dafny.SeqWithoutIsStrInference([(d_2727_v5_), self.f14, p0, (0) - (((d_2729_v7_)[len(d_2730_v8_)] if (len(d_2730_v8_)) in (d_2729_v7_) else p0)), 446])
            (globalState).f7 = len(((d_2731_v9_) + (d_2731_v9_)) + (d_2731_v9_))
            d_2732_v10_: _dafny.Seq
            d_2732_v10_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_2733_v11_: _dafny.Seq
            d_2733_v11_ = d_2732_v10_
            d_2734_v12_: _dafny.Map
            d_2734_v12_ = _dafny.Map({self.f16: (self).f17})
            d_2735_v13_: _dafny.Array
            nw539_ = _dafny.Array(None, 1)
            nw539_[int(0)] = d_2728_v6_
            d_2735_v13_ = nw539_
            d_2736_v14_: _dafny.Map
            d_2736_v14_ = _dafny.Map({not((self).fm10(d_2733_v11_, (self).fm9(self.f14, default__.fm11(self.f14, ((d_2734_v12_)[self.f16] if (self.f16) in (d_2734_v12_) else (self).f17), len(_dafny.Set({self.f14, p0})), globalState), d_2721_v0_, p0, globalState), d_2721_v0_, globalState)): d_2735_v13_})
            d_2736_v14_ = (d_2736_v14_).set(p1, d_2735_v13_)
            r0 = self.f14
        elif True:
            d_2737_v15_: _dafny.Map
            d_2737_v15_ = _dafny.Map({p1: (self.f14) >= (self.f14)})
            d_2738_v16_: D3
            d_2738_v16_ = D3_DC3(p1)
            d_2737_v15_ = (d_2737_v15_).set((p1) == ((d_2738_v16_).cf3), p1)
            if self.f16:
                (self).f16 = ((self).f17) == (True)
                (self).f16 = not(self.f16)
                d_2739_v17_: _dafny.Array
                nw540_ = _dafny.Array(False, 6)
                d_2739_v17_ = nw540_
                d_2740_v18_: _dafny.Seq
                d_2740_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nolroa"))
                rhs484_ = self.f16
                rhs485_ = d_2739_v17_
                rhs486_ = ((d_2740_v18_) + (d_2740_v18_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkamp")))
                lhs375_ = self
                lhs375_.f16 = rhs484_
                d_2739_v17_ = rhs485_
                d_2740_v18_ = rhs486_
                d_2741_v19_: str
                d_2741_v19_ = _dafny.CodePoint('d')
                d_2741_v19_ = _dafny.CodePoint('h')
                d_2742_v20_: _dafny.MultiSet
                d_2742_v20_ = _dafny.MultiSet([p1, (self).f17, self.f16])
                (globalState).f7 = ((d_2742_v20_)[self.f16] if (self.f16) in (d_2742_v20_) else self.f14)
            elif True:
                d_2743_v21_: _dafny.Array
                def lambda164_(d_2744_v15_):
                    def lambda165_(d_2745_i1_):
                        return (D4_DC6(d_2744_v15_)).cf10

                    return lambda165_

                init96_ = lambda164_(d_2737_v15_)
                nw541_ = _dafny.Array(None, 4)
                for i0_96_ in range(nw541_.length(0)):
                    nw541_[i0_96_] = init96_(i0_96_)
                d_2743_v21_ = nw541_
                d_2746_v22_: _dafny.Seq
                d_2746_v22_ = _dafny.SeqWithoutIsStrInference([d_2737_v15_])
                index464_ = default__.safeIndex(143, (d_2743_v21_).length(0))
                (d_2743_v21_)[index464_] = ((d_2746_v22_)[default__.safeIndex(p0, len(d_2746_v22_))]).set(self.f16, p1)
                index465_ = default__.safeIndex(143, (d_2743_v21_).length(0))
                (d_2743_v21_)[index465_] = (d_2737_v15_) | (d_2737_v15_)
                d_2747_v23_: _dafny.Array
                nw542_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_2747_v23_ = nw542_
                d_2748_v24_: _dafny.Array
                def lambda166_(d_2749_i2_):
                    return _dafny.CodePoint('j')

                init97_ = lambda166_
                nw543_ = _dafny.Array(None, 9)
                for i0_97_ in range(nw543_.length(0)):
                    nw543_[i0_97_] = init97_(i0_97_)
                d_2748_v24_ = nw543_
                index466_ = default__.safeIndex(345, (d_2747_v23_).length(0))
                (d_2747_v23_)[index466_] = d_2748_v24_
                index467_ = default__.safeIndex(345, (d_2747_v23_).length(0))
                (d_2747_v23_)[index467_] = d_2748_v24_
                d_2750_v25_: _dafny.Array
                def lambda167_(d_2751_i3_):
                    return (self).f17

                init98_ = lambda167_
                nw544_ = _dafny.Array(None, 22)
                for i0_98_ in range(nw544_.length(0)):
                    nw544_[i0_98_] = init98_(i0_98_)
                d_2750_v25_ = nw544_
                index468_ = default__.safeIndex(626, (d_2750_v25_).length(0))
                (d_2750_v25_)[index468_] = (self).f17
                index469_ = default__.safeIndex(626, (d_2750_v25_).length(0))
                (d_2750_v25_)[index469_] = (self).f17
                (self).f16 = self.f16
                index470_ = default__.safeIndex(626, (d_2750_v25_).length(0))
                (d_2750_v25_)[index470_] = p1
            d_2752_v26_: _dafny.MultiSet
            d_2752_v26_ = _dafny.MultiSet([self.f14, (0) - (self.f14)])
            d_2753_v27_: _dafny.Seq
            d_2753_v27_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_2754_v28_: _dafny.Seq
            d_2754_v28_ = d_2753_v27_
            d_2755_v29_: _dafny.Seq
            d_2755_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxnnlkg"))
            d_2756_v30_: str
            d_2756_v30_ = _dafny.CodePoint('c')
            d_2757_v31_: int
            d_2757_v31_ = 593
            d_2758_v32_: _dafny.MultiSet
            d_2758_v32_ = _dafny.MultiSet([p0])
            d_2759_v33_: _dafny.Array
            nw545_ = _dafny.Array(None, 13)
            nw545_[int(0)] = p1
            nw545_[int(1)] = False
            nw545_[int(2)] = self.f16
            nw545_[int(3)] = (d_2752_v26_).isdisjoint(d_2752_v26_)
            nw545_[int(4)] = self.f16
            nw545_[int(5)] = not (p1) or (self.f16)
            nw545_[int(6)] = (self).fm10(d_2753_v27_, d_2755_v29_, d_2756_v30_, globalState)
            nw545_[int(7)] = not(False)
            nw545_[int(8)] = not(p1)
            nw545_[int(9)] = not((d_2752_v26_) != (d_2752_v26_))
            nw545_[int(10)] = (self).f17
            nw545_[int(11)] = ((self).f17 if (self).f17 else self.f16)
            nw545_[int(12)] = (d_2753_v27_)[default__.safeIndex((d_2758_v32_).cardinality, len(d_2753_v27_))]
            d_2759_v33_ = nw545_
            index471_ = default__.safeIndex(484, (d_2759_v33_).length(0))
            (d_2759_v33_)[index471_] = (self).f17
            index472_ = default__.safeIndex(484, (d_2759_v33_).length(0))
            (d_2759_v33_)[index472_] = p1
            d_2760_v34_: C9
            nw546_ = C9()
            nw546_.ctor__()
            d_2760_v34_ = nw546_
            d_2761_v35_: _dafny.Seq
            d_2761_v35_ = _dafny.SeqWithoutIsStrInference([d_2755_v29_])
            d_2762_v36_: _dafny.Seq
            d_2762_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0, (d_2752_v26_).cardinality}) for d_2763_i4_ in range(default__.abs(-751))])
            d_2761_v35_ = (d_2761_v35_).set(default__.safeIndex(self.f14, len(d_2761_v35_)), (self).fm9(p0, (d_2762_v36_), d_2756_v30_, self.f14, globalState))
        if not(self.f16):
            d_2764_v37_: _dafny.Seq
            d_2764_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujipqu"))
            d_2765_v38_: _dafny.Array
            nw547_ = _dafny.Array(False, 8)
            d_2765_v38_ = nw547_
            d_2766_v39_: _dafny.Seq
            d_2766_v39_ = _dafny.SeqWithoutIsStrInference([d_2765_v38_, d_2765_v38_])
            d_2767_v40_: _dafny.Seq
            d_2767_v40_ = _dafny.SeqWithoutIsStrInference([True])
            d_2768_v41_: _dafny.Map
            d_2768_v41_ = _dafny.Map({self.f16: self.f14})
            d_2769_v42_: _dafny.Seq
            d_2769_v42_ = _dafny.SeqWithoutIsStrInference([self.f14, len(d_2766_v39_), len(d_2767_v40_), len(d_2768_v41_)])
            d_2770_v43_: C8
            nw548_ = C8()
            nw548_.ctor__(d_2764_v37_, True, not((self).f17), (0) - ((0) - ((d_2769_v42_)[default__.safeIndex(-8, len(d_2769_v42_))])))
            d_2770_v43_ = nw548_
            d_2771_v44_: str
            d_2771_v44_ = _dafny.CodePoint('e')
            d_2771_v44_ = _dafny.CodePoint('w')
            d_2772_v45_: _dafny.Array
            nw549_ = _dafny.Array(_dafny.CodePoint('D'), 11)
            d_2772_v45_ = nw549_
            d_2773_v46_: D12
            d_2773_v46_ = D12_DC23(self.f16, self.f14, p0)
            pat_let_tv45_ = p1
            pat_let_tv46_ = p0
            def iife211_(_pat_let59_0):
                def iife212_(d_2774_dt__update__tmp_h2_):
                    def iife213_(_pat_let60_0):
                        def iife214_(d_2775_dt__update_hcf30_h0_):
                            def iife215_(_pat_let61_0):
                                def iife216_(d_2776_dt__update_hcf31_h0_):
                                    return D12_DC23(d_2775_dt__update_hcf30_h0_, d_2776_dt__update_hcf31_h0_, (d_2774_dt__update__tmp_h2_).cf32)
                                return iife216_(_pat_let61_0)
                            return iife215_(pat_let_tv46_)
                        return iife214_(_pat_let60_0)
                    return iife213_(pat_let_tv45_)
                return iife212_(_pat_let59_0)
            rhs487_ = default__.fm0((iife211_(d_2773_v46_)).cf30, globalState)
            rhs488_ = (self.f14) + (((0) - (self.f14)) - (self.f14))
            rhs489_ = d_2772_v45_
            rhs490_ = p1
            lhs376_ = globalState
            lhs377_ = globalState
            lhs378_ = self
            lhs376_.f7 = rhs487_
            lhs377_.f2 = rhs488_
            d_2772_v45_ = rhs489_
            lhs378_.f16 = rhs490_
            d_2777_v47_: _dafny.MultiSet
            d_2777_v47_ = _dafny.MultiSet([(d_2767_v40_)[default__.safeIndex(p0, len(d_2767_v40_))]])
            d_2778_v48_: C4
            nw550_ = C4()
            nw550_.ctor__(p1, (d_2777_v47_).ispropersubset(d_2777_v47_), p1)
            d_2778_v48_ = nw550_
            d_2779_v49_: _dafny.Seq
            d_2779_v49_ = _dafny.SeqWithoutIsStrInference([d_2767_v40_, d_2767_v40_, d_2767_v40_])
            d_2780_v50_: bool
            out39_: bool
            out39_ = (d_2770_v43_).m7((d_2777_v47_).cardinality, d_2779_v49_, globalState)
            d_2780_v50_ = out39_
        elif True:
            d_2781_v51_: _dafny.MultiSet
            d_2781_v51_ = _dafny.MultiSet([p1])
            (self).f16 = (_dafny.MultiSet([self.f16, p1, self.f16])).ispropersubset((_dafny.MultiSet([False, self.f16, p1]) if False else d_2781_v51_))
            d_2782_v52_: _dafny.Seq
            d_2782_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsbh"))
            d_2783_v53_: _dafny.Set
            d_2783_v53_ = _dafny.Set({p1})
            d_2784_v54_: _dafny.Seq
            d_2784_v54_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p0])])
            d_2785_v55_: D23
            d_2785_v55_ = D23_DC48(self.f16, d_2784_v54_)
            pat_let_tv47_ = p1
            d_2786_v56_: C3
            nw551_ = C3()
            def iife217_(_pat_let62_0):
                def iife218_(d_2787_dt__update__tmp_h3_):
                    def iife219_(_pat_let63_0):
                        def iife220_(d_2788_dt__update_hcf74_h0_):
                            return D23_DC48(d_2788_dt__update_hcf74_h0_, (d_2787_dt__update__tmp_h3_).cf75)
                        return iife220_(_pat_let63_0)
                    return iife219_(not(pat_let_tv47_))
                return iife218_(_pat_let62_0)
            nw551_.ctor__(-458, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaguuedgl"))) + (d_2782_v52_), (d_2783_v53_) != (d_2783_v53_), (iife217_(d_2785_v55_)).cf74)
            d_2786_v56_ = nw551_
            d_2789_v57_: _dafny.Map
            d_2789_v57_ = _dafny.Map({not ((self).f17) or (self.f16): ((self).f17) == (self.f16)})
            d_2789_v57_ = (d_2789_v57_).set((self).f17, ((self).f17) or (self.f16))
            d_2790_v58_: _dafny.Array
            def lambda168_(d_2791_p0_):
                def lambda169_(d_2792_i5_):
                    return default__.safeDivisionInt(d_2792_i5_, d_2791_p0_)

                return lambda169_

            init99_ = lambda168_(p0)
            nw552_ = _dafny.Array(None, 28)
            for i0_99_ in range(nw552_.length(0)):
                nw552_[i0_99_] = init99_(i0_99_)
            d_2790_v58_ = nw552_
            (self).f16 = (d_2790_v58_) in (_dafny.MultiSet([d_2790_v58_]))
            (self).f16 = (self).f17
        (self).f16 = not((self.f14) > (p0))
        d_2793_v59_: _dafny.Array
        def lambda170_(d_2794_p0_):
            def lambda171_(d_2795_i6_):
                return (d_2795_i6_) - (d_2794_p0_)

            return lambda171_

        init100_ = lambda170_(p0)
        nw553_ = _dafny.Array(None, 24)
        for i0_100_ in range(nw553_.length(0)):
            nw553_[i0_100_] = init100_(i0_100_)
        d_2793_v59_ = nw553_
        d_2796_v60_: _dafny.MultiSet
        d_2796_v60_ = _dafny.MultiSet([d_2793_v59_, d_2793_v59_])
        if (d_2796_v60_).ispropersubset(d_2796_v60_):
            if (self.f14) <= (default__.safeDivisionInt(p0, (0) - (p0))):
                (self).f16 = (self).f17
                (globalState).f2 = p0
                d_2797_v61_: _dafny.MultiSet
                d_2797_v61_ = _dafny.MultiSet([p0])
                index473_ = default__.safeIndex(433, (d_2793_v59_).length(0))
                (d_2793_v59_)[index473_] = default__.safeDivisionInt(((d_2797_v61_)[self.f14] if (self.f14) in (d_2797_v61_) else self.f14), p0)
                d_2798_v62_: _dafny.Map
                d_2798_v62_ = _dafny.Map({d_2793_v59_: p0})
                index474_ = default__.safeIndex(433, (d_2793_v59_).length(0))
                (d_2793_v59_)[index474_] = len(((d_2798_v62_) | (d_2798_v62_)) | (d_2798_v62_))
                (self).f16 = not (not(p1)) or (True)
                d_2799_v63_: _dafny.Seq
                d_2799_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdlxlyqfn"))
                d_2799_v63_ = (d_2799_v63_) + (d_2799_v63_)
            elif True:
                d_2800_v64_: _dafny.Seq
                d_2800_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjmveunxy"))
                d_2801_v65_: D12
                d_2801_v65_ = D12_DC22((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giwttqei"))) + (d_2800_v64_))
                d_2801_v65_ = d_2801_v65_
                (self).f16 = not(p1)
                (self).f16 = (self).f17
                (self).f16 = self.f16
                d_2802_v66_: _dafny.Array
                nw554_ = _dafny.Array(False, 15)
                d_2802_v66_ = nw554_
                index475_ = default__.safeIndex(143, (d_2802_v66_).length(0))
                (d_2802_v66_)[index475_] = (self.f16) == (self.f16)
                index476_ = default__.safeIndex(143, (d_2802_v66_).length(0))
                (d_2802_v66_)[index476_] = (_dafny.CodePoint('b')) in (d_2800_v64_)
            d_2803_v67_: C9
            nw555_ = C9()
            nw555_.ctor__()
            d_2803_v67_ = nw555_
            index477_ = default__.safeIndex(752, (d_2793_v59_).length(0))
            (d_2793_v59_)[index477_] = self.f14
            index478_ = default__.safeIndex(752, (d_2793_v59_).length(0))
            (d_2793_v59_)[index478_] = self.f14
            d_2804_v68_: _dafny.Seq
            d_2804_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
            (self).f16 = (d_2804_v68_) < (d_2804_v68_)
            (self).f16 = p1
        elif True:
            d_2805_v69_: _dafny.Seq
            d_2805_v69_ = _dafny.SeqWithoutIsStrInference([self.f14])
            d_2806_v70_: _dafny.Set
            d_2806_v70_ = _dafny.Set({d_2805_v69_, d_2805_v69_, d_2805_v69_, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))]))).cardinality for d_2807_i7_ in range(default__.abs(385))])})
            if (d_2806_v70_).ispropersubset(d_2806_v70_):
                d_2808_v71_: C3
                nw556_ = C3()
                nw556_.ctor__(472, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkqdyi")), (self.f14) < (p0), (p1) or (p1))
                d_2808_v71_ = nw556_
                index479_ = default__.safeIndex(125, (d_2793_v59_).length(0))
                (d_2793_v59_)[index479_] = (0) - ((self.f14) - (p0))
                d_2809_v72_: _dafny.Seq
                d_2809_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwnvcomc"))
                index480_ = default__.safeIndex(125, (d_2793_v59_).length(0))
                rhs491_ = self.f14
                rhs492_ = 433
                rhs493_ = default__.safeModuloInt(p0, self.f14)
                rhs494_ = ((d_2809_v72_ if not(True) else (d_2808_v71_).f23)) + (d_2809_v72_)
                lhs379_ = d_2793_v59_
                lhs380_ = default__.safeIndex(125, (d_2793_v59_).length(0))
                lhs381_ = globalState
                lhs382_ = globalState
                lhs379_[lhs380_] = rhs491_
                lhs381_.f7 = rhs492_
                lhs382_.f7 = rhs493_
                d_2809_v72_ = rhs494_
                d_2810_v73_: _dafny.Array
                nw557_ = _dafny.Array(False, 26)
                d_2810_v73_ = nw557_
                d_2811_v74_: _dafny.Map
                d_2811_v74_ = _dafny.Map({(d_2793_v59_)[default__.safeIndex(125, (d_2793_v59_).length(0))]: p1})
                index481_ = default__.safeIndex(908, (d_2810_v73_).length(0))
                (d_2810_v73_)[index481_] = ((d_2811_v74_)[p0] if (p0) in (d_2811_v74_) else self.f16)
                index482_ = default__.safeIndex(908, (d_2810_v73_).length(0))
                (d_2810_v73_)[index482_] = ((d_2805_v69_) + (d_2805_v69_)) < (d_2805_v69_)
                d_2812_v75_: C9
                nw558_ = C9()
                nw558_.ctor__()
                d_2812_v75_ = nw558_
                d_2813_v77_: _dafny.Map
                def iife221_():
                    coll93_ = _dafny.Map()
                    compr_93_: int
                    for compr_93_ in _dafny.IntegerRange(320, 498):
                        d_2814_v76_: int = compr_93_
                        if ((320) <= (d_2814_v76_)) and ((d_2814_v76_) < (498)):
                            coll93_[(d_2814_v76_) - (-33)] = (d_2793_v59_)[default__.safeIndex(125, (d_2793_v59_).length(0))]
                    return _dafny.Map(coll93_)
                d_2813_v77_ = _dafny.Map({iife221_()
                : (d_2810_v73_)[default__.safeIndex(908, (d_2810_v73_).length(0))]})
                d_2815_v78_: _dafny.MultiSet
                d_2815_v78_ = _dafny.MultiSet([self.f14, (d_2793_v59_)[default__.safeIndex(125, (d_2793_v59_).length(0))], self.f14])
                d_2816_v79_: _dafny.Map
                d_2816_v79_ = _dafny.Map({(d_2793_v59_)[default__.safeIndex(125, (d_2793_v59_).length(0))]: (d_2815_v78_).cardinality})
                (globalState).f2 = len(((d_2813_v77_).set(d_2816_v79_, self.f16)).set(d_2816_v79_, not (self.f16) or ((d_2810_v73_)[default__.safeIndex(908, (d_2810_v73_).length(0))])))
            elif True:
                d_2817_v80_: _dafny.Set
                d_2817_v80_ = _dafny.Set({len(_dafny.Set({p1, self.f16}))})
                d_2818_v81_: _dafny.Array
                nw559_ = _dafny.Array(False, 2)
                d_2818_v81_ = nw559_
                d_2819_v82_: D9
                d_2819_v82_ = D9_DC15(p1, p1, d_2817_v80_, d_2793_v59_, d_2818_v81_)
                d_2820_v83_: _dafny.Set
                d_2820_v83_ = _dafny.Set({False, (self).f17, (d_2819_v82_).cf19, (self).f17, p1})
                d_2821_v84_: D15
                d_2821_v84_ = D15_DC29(d_2820_v83_)
                d_2822_v85_: _dafny.Seq
                d_2822_v85_ = _dafny.SeqWithoutIsStrInference([p1, self.f16])
                d_2823_v86_: _dafny.Array
                nw560_ = _dafny.Array(None, 20)
                nw560_[int(0)] = _dafny.Set({False, (self).f17, self.f16})
                nw560_[int(1)] = d_2820_v83_
                nw560_[int(2)] = d_2820_v83_
                nw560_[int(3)] = d_2820_v83_
                nw560_[int(4)] = d_2820_v83_
                nw560_[int(5)] = d_2820_v83_
                nw560_[int(6)] = d_2820_v83_
                nw560_[int(7)] = d_2820_v83_
                nw560_[int(8)] = default__.fm44(self.f16, p1, globalState)
                nw560_[int(9)] = d_2820_v83_
                nw560_[int(10)] = d_2820_v83_
                nw560_[int(11)] = d_2820_v83_
                nw560_[int(12)] = (d_2821_v84_).cf43
                nw560_[int(13)] = d_2820_v83_
                nw560_[int(14)] = d_2820_v83_
                nw560_[int(15)] = d_2820_v83_
                nw560_[int(16)] = _dafny.Set({(d_2822_v85_)[default__.safeIndex(920, len(d_2822_v85_))], self.f16, (d_2822_v85_)[default__.safeIndex(self.f14, len(d_2822_v85_))], not((self).f17), not(True)})
                nw560_[int(17)] = default__.fm44(default__.fm17((self).f17, p0, 831, self.f14, globalState), p1, globalState)
                nw560_[int(18)] = default__.fm44((self).f17, (self).f17, globalState)
                nw560_[int(19)] = default__.fm44((self).f17, False, globalState)
                d_2823_v86_ = nw560_
                d_2824_v87_: _dafny.Seq
                d_2824_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hubavys"))
                d_2825_v88_: C0
                nw561_ = C0()
                nw561_.ctor__(d_2823_v86_, len(d_2824_v87_))
                d_2825_v88_ = nw561_
                (self).f16 = not(self.f16)
                (self).f16 = (d_2825_v88_.f27) != (p0)
                d_2826_v89_: str
                d_2826_v89_ = _dafny.CodePoint('x')
                d_2827_v90_: _dafny.Map
                d_2827_v90_ = _dafny.Map({d_2826_v89_: self.f14})
                d_2828_v91_: _dafny.Map
                d_2828_v91_ = _dafny.Map({p0: d_2827_v90_})
                d_2828_v91_ = _dafny.Map({762: _dafny.Map({d_2826_v89_: (0) - (p0)})})
                (d_2825_v88_).f27 = d_2825_v88_.f27
            d_2829_v92_: _dafny.Array
            nw562_ = _dafny.Array(_dafny.Set({}), 23)
            d_2829_v92_ = nw562_
            d_2830_v93_: _dafny.Map
            d_2830_v93_ = _dafny.Map({True: self.f14})
            d_2831_v94_: _dafny.Set
            d_2831_v94_ = _dafny.Set({d_2830_v93_})
            index483_ = default__.safeIndex(17, (d_2829_v92_).length(0))
            (d_2829_v92_)[index483_] = d_2831_v94_
            d_2832_v95_: _dafny.MultiSet
            d_2832_v95_ = _dafny.MultiSet([self.f16, self.f16, p1])
            d_2833_v96_: _dafny.Seq
            d_2833_v96_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_2834_v97_: _dafny.Seq
            d_2834_v97_ = default__.fm52((self).f17, self.f14, len(_dafny.SeqWithoutIsStrInference([d_2832_v95_, _dafny.MultiSet([self.f16]), _dafny.MultiSet(d_2833_v96_), d_2832_v95_])), ((d_2830_v93_)[p1] if (p1) in (d_2830_v93_) else self.f14), globalState)
            d_2835_v98_: str
            d_2835_v98_ = _dafny.CodePoint('v')
            index484_ = default__.safeIndex(17, (d_2829_v92_).length(0))
            rhs495_ = (self).fm10(d_2834_v97_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2836_i8_ in range(default__.abs(462))]), d_2835_v98_, globalState)
            rhs496_ = (d_2831_v94_) | (d_2831_v94_)
            lhs383_ = self
            lhs384_ = d_2829_v92_
            lhs385_ = default__.safeIndex(17, (d_2829_v92_).length(0))
            lhs383_.f16 = rhs495_
            lhs384_[lhs385_] = rhs496_
            d_2837_v99_: _dafny.Seq
            d_2837_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhplwtr"))
            rhs497_ = (0) - (len(d_2837_v99_))
            rhs498_ = default__.safeModuloInt(self.f14, (0) - (p0))
            lhs386_ = self
            lhs387_ = globalState
            lhs386_.f14 = rhs497_
            lhs387_.f7 = rhs498_
            r0 = (721) - (self.f14)
            (self).f16 = not((D9_DC16(self.f14, False)).cf24)
        d_2838_v100_: _dafny.Map
        d_2838_v100_ = _dafny.Map({-935: self.f14})
        d_2839_v101_: _dafny.Map
        d_2839_v101_ = _dafny.Map({not(p1): (_dafny.Map({p0: default__.fm0(self.f16, globalState)})) | (d_2838_v100_)})
        d_2839_v101_ = (d_2839_v101_).set((self).f17, (self).fm3(globalState))
        d_2840_v102_: D7
        d_2840_v102_ = D7_DC12()
        d_2841_v103_: _dafny.Seq
        d_2841_v103_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2842_v104_: _dafny.Map
        d_2842_v104_ = _dafny.Map({d_2840_v102_: d_2841_v103_})
        d_2842_v104_ = d_2842_v104_
        r0 = p0
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_2843_v0_: str
        d_2843_v0_ = _dafny.CodePoint('k')
        d_2843_v0_ = d_2843_v0_
        if not((self).f17):
            d_2844_v1_: _dafny.Set
            d_2844_v1_ = _dafny.Set({(self).f17})
            d_2845_v2_: _dafny.Seq
            d_2845_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqllnxkei"))
            d_2846_v3_: _dafny.Map
            d_2846_v3_ = _dafny.Map({d_2844_v1_: d_2845_v2_})
            d_2846_v3_ = (d_2846_v3_).set((d_2844_v1_) | (d_2844_v1_), d_2845_v2_)
            d_2847_v4_: _dafny.Array
            nw563_ = _dafny.Array(int(0), 18)
            d_2847_v4_ = nw563_
            index485_ = default__.safeIndex(408, (d_2847_v4_).length(0))
            (d_2847_v4_)[index485_] = (p1) * (self.f14)
            index486_ = default__.safeIndex(408, (d_2847_v4_).length(0))
            (d_2847_v4_)[index486_] = self.f14
            d_2848_v5_: _dafny.Set
            d_2848_v5_ = _dafny.Set({p0, _dafny.CodePoint('h'), d_2843_v0_})
            d_2849_v6_: _dafny.Set
            d_2849_v6_ = _dafny.Set({d_2848_v5_})
            d_2849_v6_ = d_2849_v6_
            d_2847_v4_ = d_2847_v4_
            d_2850_v7_: D9
            d_2850_v7_ = D9_DC16(self.f14, p2)
            d_2851_v8_: D9
            d_2851_v8_ = D9_DC17(d_2850_v7_)
            pat_let_tv48_ = p1
            d_2852_v9_: _dafny.Array
            nw564_ = _dafny.Array(None, 1)
            def iife222_(_pat_let64_0):
                def iife223_(d_2853_dt__update__tmp_h0_):
                    def iife224_(_pat_let65_0):
                        def iife225_(d_2854_dt__update_hcf25_h0_):
                            return D9_DC17(d_2854_dt__update_hcf25_h0_)
                        return iife225_(_pat_let65_0)
                    return iife224_(D9_DC16(pat_let_tv48_, (self).f17))
                return iife223_(_pat_let64_0)
            nw564_[int(0)] = iife222_(d_2851_v8_)
            d_2852_v9_ = nw564_
            index487_ = default__.safeIndex(412, (d_2852_v9_).length(0))
            (d_2852_v9_)[index487_] = d_2851_v8_
            index488_ = default__.safeIndex(412, (d_2852_v9_).length(0))
            (d_2852_v9_)[index488_] = d_2851_v8_
        elif True:
            d_2855_v10_: _dafny.MultiSet
            d_2855_v10_ = _dafny.MultiSet([self.f14])
            d_2856_v11_: _dafny.Array
            nw565_ = _dafny.Array(None, 8)
            nw565_[int(0)] = (_dafny.MultiSet([p1])).cardinality
            nw565_[int(1)] = (self.f14) * (self.f14)
            nw565_[int(2)] = default__.safeDivisionInt((d_2855_v10_).cardinality, -184)
            nw565_[int(3)] = self.f14
            nw565_[int(4)] = default__.safeDivisionInt(self.f14, p1)
            nw565_[int(5)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ripqy"))), self.f14)
            nw565_[int(6)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fes"))))
            nw565_[int(7)] = default__.safeModuloInt(p1, p1)
            d_2856_v11_ = nw565_
            index489_ = default__.safeIndex(625, (d_2856_v11_).length(0))
            (d_2856_v11_)[index489_] = p1
            index490_ = default__.safeIndex(625, (d_2856_v11_).length(0))
            (d_2856_v11_)[index490_] = (0) - (p1)
            d_2857_v12_: C7
            nw566_ = C7()
            nw566_.ctor__(self.f16, (self).f17, self.f14)
            d_2857_v12_ = nw566_
            index491_ = default__.safeIndex(625, (d_2856_v11_).length(0))
            (d_2856_v11_)[index491_] = (0) - (default__.safeModuloInt((0) - (p3), len(_dafny.SeqWithoutIsStrInference([d_2857_v12_]))))
            index492_ = default__.safeIndex(625, (d_2856_v11_).length(0))
            (d_2856_v11_)[index492_] = default__.safeDivisionInt(default__.safeDivisionInt(-631, self.f14), 128)
            d_2858_v13_: _dafny.Seq
            d_2858_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ip"))
            d_2859_v14_: _dafny.Seq
            d_2859_v14_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_2860_v15_: C8
            nw567_ = C8()
            nw567_.ctor__((d_2858_v13_) + (d_2858_v13_), not (p2) or ((self).fm10(d_2859_v14_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), d_2843_v0_, globalState)), not (self.f16) or (p2), (len(d_2858_v13_)) * (self.f14))
            d_2860_v15_ = nw567_
            d_2861_v16_: _dafny.Map
            d_2861_v16_ = _dafny.Map({d_2856_v11_: d_2856_v11_})
            d_2862_v17_: _dafny.Map
            d_2862_v17_ = _dafny.Map({((d_2861_v16_)[d_2856_v11_] if (d_2856_v11_) in (d_2861_v16_) else d_2856_v11_): d_2843_v0_})
            d_2862_v17_ = (d_2862_v17_).set(d_2856_v11_, p0)
        d_2863_v18_: D12
        d_2863_v18_ = D12_DC23((self).f17, p3, p1)
        d_2864_v19_: _dafny.Seq
        d_2864_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lneeohgn"))
        (self).f16 = (default__.safeDivisionInt((d_2863_v18_).cf32, len(d_2864_v19_))) <= (p3)
        d_2865_v20_: _dafny.Array
        nw568_ = _dafny.Array(int(0), 28)
        d_2865_v20_ = nw568_
        index493_ = default__.safeIndex(514, (d_2865_v20_).length(0))
        (d_2865_v20_)[index493_] = p1
        index494_ = default__.safeIndex(514, (d_2865_v20_).length(0))
        (d_2865_v20_)[index494_] = (0) - (self.f14)
        d_2866_v21_: _dafny.Array
        nw569_ = _dafny.Array(None, 13)
        nw569_[int(0)] = p0
        nw569_[int(1)] = d_2843_v0_
        nw569_[int(2)] = p0
        nw569_[int(3)] = p0
        nw569_[int(4)] = p0
        nw569_[int(5)] = d_2843_v0_
        nw569_[int(6)] = d_2843_v0_
        nw569_[int(7)] = default__.fm39(globalState)
        nw569_[int(8)] = d_2843_v0_
        nw569_[int(9)] = d_2843_v0_
        nw569_[int(10)] = p0
        nw569_[int(11)] = (d_2843_v0_ if p2 else p0)
        nw569_[int(12)] = p0
        d_2866_v21_ = nw569_
        index495_ = default__.safeIndex(357, (d_2866_v21_).length(0))
        (d_2866_v21_)[index495_] = (d_2864_v19_)[default__.safeIndex(782, len(d_2864_v19_))]
        d_2867_v22_: _dafny.Seq
        d_2867_v22_ = _dafny.SeqWithoutIsStrInference([p1, ((d_2865_v20_)[default__.safeIndex(514, (d_2865_v20_).length(0))] if self.f16 else (d_2865_v20_)[default__.safeIndex(514, (d_2865_v20_).length(0))]), (d_2865_v20_)[default__.safeIndex(514, (d_2865_v20_).length(0))]])
        d_2868_v23_: _dafny.Map
        d_2868_v23_ = _dafny.Map({default__.fm21(globalState): self.f14})
        d_2869_v24_: _dafny.Map
        d_2869_v24_ = d_2868_v23_
        index496_ = default__.safeIndex(357, (d_2866_v21_).length(0))
        index497_ = default__.safeIndex(514, (d_2865_v20_).length(0))
        rhs499_ = d_2843_v0_
        rhs500_ = p1
        rhs501_ = default__.fm34((d_2863_v18_).cf32, (d_2865_v20_)[default__.safeIndex(514, (d_2865_v20_).length(0))], self.f16, (p1 if p2 else p1), globalState)
        rhs502_ = d_2867_v22_
        rhs503_ = default__.safeModuloInt(len((d_2869_v24_)), (0) - ((p3) + (p3)))
        lhs388_ = d_2866_v21_
        lhs389_ = default__.safeIndex(357, (d_2866_v21_).length(0))
        lhs390_ = d_2865_v20_
        lhs391_ = default__.safeIndex(514, (d_2865_v20_).length(0))
        lhs388_[lhs389_] = rhs499_
        lhs390_[lhs391_] = rhs500_
        d_2843_v0_ = rhs501_
        d_2867_v22_ = rhs502_
        r0 = rhs503_
        index498_ = default__.safeIndex(357, (d_2866_v21_).length(0))
        (d_2866_v21_)[index498_] = _dafny.CodePoint('w')
        d_2870_v25_: _dafny.Set
        d_2870_v25_ = _dafny.Set({p3, (d_2867_v22_)[default__.safeIndex(p1, len(d_2867_v22_))], len(d_2864_v19_), (d_2865_v20_)[default__.safeIndex(514, (d_2865_v20_).length(0))]})
        d_2871_v26_: _dafny.Map
        d_2871_v26_ = _dafny.Map({d_2870_v25_: p1})
        r0 = len(default__.fm66(p1, d_2871_v26_, (len(d_2864_v19_)) * (p1), p2, globalState))
        return r0


class C11(T0, T1):
    def  __init__(self):
        self._f14: int = int(0)
        self._f16: bool = False
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f14, f16, f17):
        (self).f14 = f14
        (self).f16 = f16
        (self)._f17 = f17

    def fm3(self, globalState):
        return (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f16, not(self.f16)])]))): (0) - (len(_dafny.Set({False, self.f16})))})) | (_dafny.Map({self.f14: self.f14}))

    def fm4(self, p0, p1, p2, globalState):
        return self.f16

    def fm5(self, p0, p1, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owwe"))])

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_2872_i0_: int
        d_2872_i0_ = 0
        with _dafny.label("18"):
            while p1:
                with _dafny.c_label("18"):
                    if (d_2872_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_2872_i0_ = (d_2872_i0_) + (1)
                    d_2873_v0_: _dafny.Array
                    nw570_ = _dafny.Array(int(0), 13)
                    d_2873_v0_ = nw570_
                    index499_ = default__.safeIndex(621, (d_2873_v0_).length(0))
                    (d_2873_v0_)[index499_] = p0
                    d_2874_v1_: str
                    d_2874_v1_ = _dafny.CodePoint('f')
                    d_2875_v2_: _dafny.Map
                    d_2875_v2_ = _dafny.Map({p1: d_2874_v1_})
                    index500_ = default__.safeIndex(621, (d_2873_v0_).length(0))
                    (d_2873_v0_)[index500_] = len((d_2875_v2_).set(self.f16, d_2874_v1_))
                    d_2876_v3_: _dafny.Seq
                    d_2876_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcloiqko"))
                    d_2877_v4_: _dafny.Seq
                    d_2877_v4_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                    d_2878_v5_: _dafny.Map
                    d_2878_v5_ = _dafny.Map({d_2877_v4_: default__.fm6(len(default__.fm7(p1, (self).f17, p1, globalState)), self.f16, not(False), p1, globalState)})
                    d_2879_v6_: _dafny.Seq
                    d_2879_v6_ = _dafny.SeqWithoutIsStrInference([d_2876_v3_])
                    d_2876_v3_ = ((d_2878_v5_)[d_2877_v4_] if (d_2877_v4_) in (d_2878_v5_) else ((d_2879_v6_)[default__.safeIndex((d_2873_v0_)[default__.safeIndex(621, (d_2873_v0_).length(0))], len(d_2879_v6_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chchihtm"))))
                    if p1:
                        d_2880_v7_: _dafny.Seq
                        d_2880_v7_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_2877_v4_ = _dafny.SeqWithoutIsStrInference([self.f16])
                        d_2881_v8_: _dafny.Array
                        nw571_ = _dafny.Array(_dafny.Map({}), 19)
                        d_2881_v8_ = nw571_
                        d_2882_v10_: _dafny.Map
                        def iife226_():
                            coll94_ = _dafny.Set()
                            compr_94_: int
                            for compr_94_ in _dafny.IntegerRange(350, -32):
                                d_2883_v9_: int = compr_94_
                                if ((350) <= (d_2883_v9_)) and ((d_2883_v9_) < (-32)):
                                    coll94_ = coll94_.union(_dafny.Set([(d_2883_v9_) - (p0)]))
                            return _dafny.Set(coll94_)
                        d_2882_v10_ = _dafny.Map({iife226_()
                        : self.f14})
                        index501_ = default__.safeIndex(962, (d_2881_v8_).length(0))
                        (d_2881_v8_)[index501_] = d_2882_v10_
                        d_2884_v11_: _dafny.Array
                        nw572_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                        d_2884_v11_ = nw572_
                        index502_ = default__.safeIndex(598, (d_2884_v11_).length(0))
                        (d_2884_v11_)[index502_] = d_2876_v3_
                        d_2885_v12_: _dafny.Array
                        nw573_ = _dafny.Array(False, 23)
                        d_2885_v12_ = nw573_
                        d_2886_v13_: _dafny.Array
                        d_2886_v13_ = d_2885_v12_
                        d_2887_v14_: _dafny.Array
                        nw574_ = _dafny.Array(None, 25)
                        nw574_[int(0)] = d_2885_v12_
                        nw574_[int(1)] = d_2885_v12_
                        nw574_[int(2)] = d_2885_v12_
                        nw574_[int(3)] = d_2885_v12_
                        nw574_[int(4)] = d_2885_v12_
                        nw574_[int(5)] = d_2885_v12_
                        nw574_[int(6)] = (d_2886_v13_)
                        nw574_[int(7)] = d_2885_v12_
                        nw574_[int(8)] = d_2885_v12_
                        nw574_[int(9)] = (d_2885_v12_ if self.f16 else d_2885_v12_)
                        nw574_[int(10)] = d_2885_v12_
                        nw574_[int(11)] = d_2885_v12_
                        nw574_[int(12)] = (d_2885_v12_ if self.f16 else d_2885_v12_)
                        nw574_[int(13)] = d_2885_v12_
                        nw574_[int(14)] = d_2885_v12_
                        nw574_[int(15)] = d_2885_v12_
                        nw574_[int(16)] = (d_2885_v12_ if (self).f17 else d_2885_v12_)
                        nw574_[int(17)] = d_2885_v12_
                        nw574_[int(18)] = d_2885_v12_
                        nw574_[int(19)] = d_2885_v12_
                        nw574_[int(20)] = d_2885_v12_
                        nw574_[int(21)] = d_2885_v12_
                        nw574_[int(22)] = d_2885_v12_
                        nw574_[int(23)] = d_2885_v12_
                        nw574_[int(24)] = d_2885_v12_
                        d_2887_v14_ = nw574_
                        index503_ = default__.safeIndex(409, (d_2887_v14_).length(0))
                        (d_2887_v14_)[index503_] = d_2885_v12_
                        d_2888_v15_: _dafny.MultiSet
                        d_2888_v15_ = _dafny.MultiSet([(0) - (self.f14), self.f14])
                        index504_ = default__.safeIndex(962, (d_2881_v8_).length(0))
                        index505_ = default__.safeIndex(598, (d_2884_v11_).length(0))
                        index506_ = default__.safeIndex(409, (d_2887_v14_).length(0))
                        rhs504_ = self.f16
                        rhs505_ = d_2882_v10_
                        rhs506_ = d_2876_v3_
                        rhs507_ = d_2885_v12_
                        rhs508_ = len((default__.fm8(not(p1), (self).f17, globalState)).set(d_2888_v15_, (d_2876_v3_ if not(p1) else d_2876_v3_)))
                        lhs392_ = self
                        lhs393_ = d_2881_v8_
                        lhs394_ = default__.safeIndex(962, (d_2881_v8_).length(0))
                        lhs395_ = d_2884_v11_
                        lhs396_ = default__.safeIndex(598, (d_2884_v11_).length(0))
                        lhs397_ = d_2887_v14_
                        lhs398_ = default__.safeIndex(409, (d_2887_v14_).length(0))
                        lhs399_ = globalState
                        lhs392_.f16 = rhs504_
                        lhs393_[lhs394_] = rhs505_
                        lhs395_[lhs396_] = rhs506_
                        lhs397_[lhs398_] = rhs507_
                        lhs399_.f7 = rhs508_
                        d_2889_v16_: _dafny.Seq
                        d_2889_v16_ = _dafny.SeqWithoutIsStrInference([d_2885_v12_])
                        d_2890_v17_: _dafny.Array
                        nw575_ = _dafny.Array(None, 27)
                        nw575_[int(0)] = d_2886_v13_
                        nw575_[int(1)] = d_2886_v13_
                        nw575_[int(2)] = d_2886_v13_
                        nw575_[int(3)] = d_2886_v13_
                        nw575_[int(4)] = d_2886_v13_
                        nw575_[int(5)] = d_2886_v13_
                        nw575_[int(6)] = (d_2887_v14_)[default__.safeIndex(409, (d_2887_v14_).length(0))]
                        nw575_[int(7)] = d_2886_v13_
                        nw575_[int(8)] = d_2886_v13_
                        nw575_[int(9)] = (d_2886_v13_ if self.f16 else d_2886_v13_)
                        nw575_[int(10)] = (d_2886_v13_ if p1 else d_2886_v13_)
                        nw575_[int(11)] = d_2886_v13_
                        nw575_[int(12)] = d_2886_v13_
                        nw575_[int(13)] = d_2886_v13_
                        nw575_[int(14)] = d_2886_v13_
                        nw575_[int(15)] = (d_2887_v14_)[default__.safeIndex(409, (d_2887_v14_).length(0))]
                        nw575_[int(16)] = d_2886_v13_
                        nw575_[int(17)] = d_2886_v13_
                        nw575_[int(18)] = d_2886_v13_
                        nw575_[int(19)] = (d_2887_v14_)[default__.safeIndex(409, (d_2887_v14_).length(0))]
                        nw575_[int(20)] = d_2886_v13_
                        nw575_[int(21)] = (d_2887_v14_)[default__.safeIndex(409, (d_2887_v14_).length(0))]
                        nw575_[int(22)] = d_2885_v12_
                        nw575_[int(23)] = d_2886_v13_
                        nw575_[int(24)] = d_2886_v13_
                        nw575_[int(25)] = (d_2889_v16_)[default__.safeIndex((d_2873_v0_)[default__.safeIndex(621, (d_2873_v0_).length(0))], len(d_2889_v16_))]
                        nw575_[int(26)] = d_2886_v13_
                        d_2890_v17_ = nw575_
                        index507_ = default__.safeIndex(401, (d_2890_v17_).length(0))
                        (d_2890_v17_)[index507_] = d_2886_v13_
                        index508_ = default__.safeIndex(401, (d_2890_v17_).length(0))
                        (d_2890_v17_)[index508_] = d_2886_v13_
                        d_2891_v18_: _dafny.Map
                        d_2891_v18_ = _dafny.Map({(d_2873_v0_)[default__.safeIndex(621, (d_2873_v0_).length(0))]: self.f14})
                        d_2892_v19_: _dafny.Map
                        d_2892_v19_ = _dafny.Map({(self).f17: len(d_2891_v18_)})
                        d_2893_v20_: int
                        d_2893_v20_ = ((d_2892_v19_)[(self).f17] if ((self).f17) in (d_2892_v19_) else self.f14)
                        index509_ = default__.safeIndex(621, (d_2873_v0_).length(0))
                        (d_2873_v0_)[index509_] = ((d_2873_v0_)[default__.safeIndex(621, (d_2873_v0_).length(0))])
                        (self).f16 = (self).f17
                    elif True:
                        d_2894_v21_: _dafny.Array
                        nw576_ = _dafny.Array(False, 5)
                        d_2894_v21_ = nw576_
                        d_2895_v22_: _dafny.Array
                        d_2895_v22_ = d_2894_v21_
                        d_2896_v23_: _dafny.MultiSet
                        d_2896_v23_ = _dafny.MultiSet([d_2895_v22_])
                        (self).f16 = not((d_2896_v23_).ispropersubset(d_2896_v23_))
                        d_2897_v24_: _dafny.MultiSet
                        d_2897_v24_ = _dafny.MultiSet([self.f14])
                        d_2898_v25_: T0
                        nw577_ = C6()
                        nw577_.ctor__(self.f14, (self).f17, self.f16, self.f16, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afxfh"))))
                        d_2898_v25_ = nw577_
                        d_2899_v26_: D20
                        d_2899_v26_ = D20_DC42(len(d_2876_v3_), default__.fm27(759, self.f14, (d_2897_v24_).cardinality, _dafny.Map({p0: 613}), globalState), d_2898_v25_)
                        d_2900_v27_: C1
                        nw578_ = C1()
                        nw578_.ctor__((self).f17, default__.fm0((d_2899_v26_).cf66, globalState))
                        d_2900_v27_ = nw578_
                        d_2901_v28_: _dafny.Array
                        nw579_ = _dafny.Array(_dafny.Seq({}), 21)
                        d_2901_v28_ = nw579_
                        d_2902_v29_: C9
                        nw580_ = C9()
                        nw580_.ctor__()
                        d_2902_v29_ = nw580_
                        d_2903_v30_: _dafny.Seq
                        d_2903_v30_ = _dafny.SeqWithoutIsStrInference([d_2902_v29_])
                        index510_ = default__.safeIndex(931, (d_2901_v28_).length(0))
                        (d_2901_v28_)[index510_] = (d_2903_v30_) + (d_2903_v30_)
                        index511_ = default__.safeIndex(931, (d_2901_v28_).length(0))
                        (d_2901_v28_)[index511_] = d_2903_v30_
                        d_2904_v31_: _dafny.Set
                        d_2904_v31_ = _dafny.Set({d_2876_v3_})
                        d_2904_v31_ = d_2904_v31_
                        d_2905_v32_: D12
                        d_2905_v32_ = D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocj")))
                        d_2906_v33_: _dafny.Set
                        d_2906_v33_ = _dafny.Set({self.f14, p0})
                        d_2876_v3_ = (d_2876_v3_) + ((((d_2905_v32_).cf29).set(default__.safeIndex(len(d_2906_v33_), len((d_2905_v32_).cf29)), d_2874_v1_)) + (d_2876_v3_))
                    (self).f16 = self.f16
                    pass
            pass
        d_2907_v34_: _dafny.MultiSet
        d_2907_v34_ = _dafny.MultiSet([self.f14, self.f14])
        d_2908_v35_: _dafny.Map
        d_2908_v35_ = _dafny.Map({self.f14: p0})
        d_2909_v36_: D19
        d_2909_v36_ = D19_DC39(_dafny.MultiSet([p0, len(d_2908_v35_)]))
        (self).f16 = (d_2907_v34_).isdisjoint((d_2909_v36_).cf60)
        d_2910_v37_: D6
        d_2910_v37_ = D6_DC10()
        def lambda172_(source40_):
            if source40_.is_DC10:
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhsmv"))
            elif True:
                d_2911___mcc_h0_ = source40_.cf14
                d_2912_cf14_ = d_2911___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krbh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcvjhvxyw")))

        r0 = (0) - (len(lambda172_(d_2910_v37_)))
        d_2913_v38_: _dafny.Map
        d_2913_v38_ = _dafny.Map({default__.fm17((self).f17, self.f14, p0, self.f14, globalState): not(True)})
        if ((d_2913_v38_)[default__.fm21(globalState)] if (default__.fm21(globalState)) in (d_2913_v38_) else self.f16):
            d_2914_v39_: D3
            d_2914_v39_ = D3_DC4(p1, 344, p0, self.f16, p1)
            if (d_2914_v39_).cf7:
                d_2915_v40_: _dafny.Seq
                d_2915_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eea"))
                d_2916_v41_: _dafny.Map
                d_2916_v41_ = _dafny.Map({_dafny.Set({self.f14, p0}): d_2915_v40_})
                d_2917_v42_: _dafny.Array
                nw581_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                d_2917_v42_ = nw581_
                d_2918_v43_: str
                d_2918_v43_ = _dafny.CodePoint('f')
                d_2919_v44_: D15
                d_2919_v44_ = D15_DC30(d_2918_v43_)
                index512_ = default__.safeIndex(709, (d_2917_v42_).length(0))
                (d_2917_v42_)[index512_] = (d_2919_v44_).cf44
                index513_ = default__.safeIndex(709, (d_2917_v42_).length(0))
                rhs509_ = d_2916_v41_
                rhs510_ = d_2918_v43_
                lhs400_ = d_2917_v42_
                lhs401_ = default__.safeIndex(709, (d_2917_v42_).length(0))
                d_2916_v41_ = rhs509_
                lhs400_[lhs401_] = rhs510_
                d_2920_v45_: _dafny.Map
                d_2920_v45_ = _dafny.Map({len(d_2915_v40_): default__.fm28(default__.fm21(globalState), (self).fm3(globalState), D3_DC4(p1, len(d_2915_v40_), p0, self.f16, (self).f17), d_2913_v38_, globalState)})
                rhs511_ = d_2920_v45_
                rhs512_ = p0
                d_2920_v45_ = rhs511_
                r0 = rhs512_
                d_2921_v46_: _dafny.Map
                d_2921_v46_ = _dafny.Map({self.f16: (d_2917_v42_)[default__.safeIndex(709, (d_2917_v42_).length(0))]})
                d_2922_v47_: _dafny.Map
                d_2922_v47_ = _dafny.Map({((d_2921_v46_)[p1] if (p1) in (d_2921_v46_) else d_2918_v43_): d_2918_v43_})
                d_2923_v48_: D16
                d_2923_v48_ = D16_DC33((self).f17, p1, self.f14, (self).f17)
                d_2922_v47_ = (d_2922_v47_).set(d_2918_v43_, (d_2915_v40_)[default__.safeIndex((d_2923_v48_).cf49, len(d_2915_v40_))])
                d_2924_v49_: _dafny.Array
                nw582_ = _dafny.Array(_dafny.Set({}), 14)
                d_2924_v49_ = nw582_
                d_2925_v50_: C0
                nw583_ = C0()
                nw583_.ctor__(d_2924_v49_, (0) - ((0) - (p0)))
                d_2925_v50_ = nw583_
                d_2926_v51_: _dafny.Array
                nw584_ = _dafny.Array(int(0), 26)
                d_2926_v51_ = nw584_
                index514_ = default__.safeIndex(290, (d_2926_v51_).length(0))
                (d_2926_v51_)[index514_] = p0
                index515_ = default__.safeIndex(290, (d_2926_v51_).length(0))
                (d_2926_v51_)[index515_] = p0
            elif True:
                (self).f16 = (self).f17
                d_2927_v52_: C4
                nw585_ = C4()
                nw585_.ctor__(True, (self).f17, (self).f17)
                d_2927_v52_ = nw585_
                d_2928_v53_: _dafny.Seq
                d_2928_v53_ = _dafny.SeqWithoutIsStrInference([d_2927_v52_])
                d_2929_v54_: _dafny.MultiSet
                d_2929_v54_ = _dafny.MultiSet([d_2927_v52_, d_2927_v52_])
                d_2930_v55_: _dafny.Seq
                d_2930_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2928_v53_)[default__.safeIndex(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2931_i1_ in range(default__.abs(524))]))})), len(d_2928_v53_))]]), d_2929_v54_])
                d_2932_v56_: _dafny.Seq
                d_2932_v56_ = _dafny.SeqWithoutIsStrInference([(self).f17, False])
                d_2933_v57_: T0
                nw586_ = C7()
                nw586_.ctor__(p1, (self).f17, p0)
                d_2933_v57_ = nw586_
                d_2934_v58_: D14
                d_2934_v58_ = D14_DC28(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2935_i2_ in range(default__.abs(477))])), (self).f17, d_2933_v57_)
                d_2936_v59_: D20
                d_2936_v59_ = D20_DC42((d_2934_v58_).cf40, self.f16, d_2933_v57_)
                d_2937_v61_: _dafny.Set
                d_2937_v61_ = _dafny.Set({(d_2927_v52_).f21})
                d_2938_v62_: _dafny.Seq
                d_2938_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whx"))
                d_2939_v63_: _dafny.Map
                d_2939_v63_ = _dafny.Map({len(d_2938_v62_): p1})
                d_2940_v65_: _dafny.Seq
                def iife227_():
                    def iife229_():
                        coll97_ = _dafny.Map()
                        compr_97_: _dafny.Set
                        for compr_97_ in (_dafny.MultiSet([d_2937_v61_])).Elements:
                            d_2941_v60_: _dafny.Set = compr_97_
                            if (d_2941_v60_) in (_dafny.MultiSet([d_2937_v61_])):
                                coll97_[d_2941_v60_] = _dafny.Map({d_2933_v57_.f14: (self).f17})
                        return _dafny.Map(coll97_)
                    coll95_ = _dafny.Set()
                    def iife228_():
                        coll96_ = _dafny.Map()
                        compr_96_: _dafny.Set
                        for compr_96_ in (_dafny.MultiSet([d_2937_v61_])).Elements:
                            d_2941_v60_: _dafny.Set = compr_96_
                            if (d_2941_v60_) in (_dafny.MultiSet([d_2937_v61_])):
                                coll96_[d_2941_v60_] = _dafny.Map({d_2933_v57_.f14: (self).f17})
                        return _dafny.Map(coll96_)
                    compr_95_: _dafny.Set
                    for compr_95_ in ((iife228_()
                    ).set(d_2937_v61_, d_2939_v63_)).keys.Elements:
                        d_2942_v64_: _dafny.Set = compr_95_
                        if (d_2942_v64_) in ((iife229_()
                        ).set(d_2937_v61_, d_2939_v63_)):
                            coll95_ = coll95_.union(_dafny.Set([d_2942_v64_]))
                    return _dafny.Set(coll95_)
                d_2940_v65_ = _dafny.SeqWithoutIsStrInference([len(iife227_()
                )])
                d_2943_v66_: str
                d_2943_v66_ = _dafny.CodePoint('u')
                d_2944_v67_: D15
                d_2944_v67_ = D15_DC30(d_2943_v66_)
                d_2945_v68_: _dafny.Set
                d_2945_v68_ = _dafny.Set({d_2944_v67_})
                d_2946_v70_: _dafny.MultiSet
                d_2946_v70_ = _dafny.MultiSet([d_2943_v66_])
                d_2947_v72_: _dafny.Array
                nw587_ = _dafny.Array(None, 24)
                nw587_[int(0)] = self.f14
                nw587_[int(1)] = (0) - ((((d_2930_v55_)[default__.safeIndex(p0, len(d_2930_v55_))]).cardinality) - (self.f14))
                nw587_[int(2)] = p0
                nw587_[int(3)] = (self.f14) - (p0)
                nw587_[int(4)] = (self.f14) - (self.f14)
                nw587_[int(5)] = default__.safeDivisionInt(p0, len(d_2932_v56_))
                nw587_[int(6)] = (d_2927_v52_).fm12(p0, globalState)
                nw587_[int(7)] = p0
                nw587_[int(8)] = (d_2936_v59_).cf65
                nw587_[int(9)] = default__.safeModuloInt(d_2933_v57_.f14, p0)
                nw587_[int(10)] = (d_2914_v39_).cf6
                nw587_[int(11)] = (0) - ((default__.fm0((self).f17, globalState)) - ((0) - (self.f14)))
                nw587_[int(12)] = default__.safeDivisionInt(p0, d_2933_v57_.f14)
                nw587_[int(13)] = (d_2940_v65_)[default__.safeIndex(len(d_2945_v68_), len(d_2940_v65_))]
                nw587_[int(14)] = p0
                nw587_[int(15)] = (p0) * (self.f14)
                def iife230_():
                    coll98_ = _dafny.Set()
                    compr_98_: int
                    for compr_98_ in _dafny.IntegerRange(698, 631):
                        d_2948_v69_: int = compr_98_
                        if ((698) <= (d_2948_v69_)) and ((d_2948_v69_) < (631)):
                            coll98_ = coll98_.union(_dafny.Set([(d_2948_v69_) - ((_dafny.MultiSet(d_2932_v56_)).cardinality)]))
                    return _dafny.Set(coll98_)
                nw587_[int(16)] = (len(iife230_()
                )) - (self.f14)
                nw587_[int(17)] = p0
                nw587_[int(18)] = (509) - (p0)
                nw587_[int(19)] = -150
                def iife231_():
                    coll99_ = _dafny.Set()
                    compr_99_: str
                    for compr_99_ in (d_2946_v70_).Elements:
                        d_2949_v71_: str = compr_99_
                        if (d_2949_v71_) in (d_2946_v70_):
                            coll99_ = coll99_.union(_dafny.Set([d_2949_v71_]))
                    return _dafny.Set(coll99_)
                nw587_[int(20)] = (len(iife231_()
                )) - ((0) - (self.f14))
                nw587_[int(21)] = d_2933_v57_.f14
                nw587_[int(22)] = self.f14
                nw587_[int(23)] = p0
                d_2947_v72_ = nw587_
                d_2947_v72_ = d_2947_v72_
                (globalState).f7 = -954
                d_2950_v73_: _dafny.Array
                def lambda173_(d_2951_v57_, d_2952_v52_):
                    def lambda174_(d_2953_i3_):
                        return D21_DC44(d_2951_v57_.f14, (d_2952_v52_).f21)

                    return lambda174_

                init101_ = lambda173_(d_2933_v57_, d_2927_v52_)
                nw588_ = _dafny.Array(None, 7)
                for i0_101_ in range(nw588_.length(0)):
                    nw588_[i0_101_] = init101_(i0_101_)
                d_2950_v73_ = nw588_
                d_2954_v74_: D21
                d_2954_v74_ = D21_DC44(self.f14, p1)
                index516_ = default__.safeIndex(996, (d_2950_v73_).length(0))
                (d_2950_v73_)[index516_] = d_2954_v74_
                index517_ = default__.safeIndex(996, (d_2950_v73_).length(0))
                (d_2950_v73_)[index517_] = (d_2954_v74_ if True else d_2954_v74_)
                d_2955_v75_: D15
                d_2955_v75_ = D15_DC30(_dafny.CodePoint('j'))
                d_2956_v76_: D15
                d_2956_v76_ = D15_DC31(d_2955_v75_)
                d_2957_v77_: D19
                d_2957_v77_ = D19_DC40(793, not(p1), d_2933_v57_.f14)
                rhs513_ = d_2956_v76_
                rhs514_ = p0
                rhs515_ = not ((d_2957_v77_).cf62) or (p1)
                rhs516_ = (self.f14) != (len(d_2932_v56_))
                rhs517_ = d_2927_v52_
                lhs402_ = globalState
                lhs403_ = self
                lhs404_ = self
                d_2956_v76_ = rhs513_
                lhs402_.f7 = rhs514_
                lhs403_.f16 = rhs515_
                lhs404_.f16 = rhs516_
                d_2927_v52_ = rhs517_
            d_2958_v78_: _dafny.Array
            def lambda175_(d_2959_i4_):
                return self.f16

            init102_ = lambda175_
            nw589_ = _dafny.Array(None, 13)
            for i0_102_ in range(nw589_.length(0)):
                nw589_[i0_102_] = init102_(i0_102_)
            d_2958_v78_ = nw589_
            index518_ = default__.safeIndex(525, (d_2958_v78_).length(0))
            (d_2958_v78_)[index518_] = (p1) == (self.f16)
            d_2960_v79_: D3
            d_2960_v79_ = D3_DC3(p1)
            d_2961_v80_: D18
            d_2961_v80_ = D18_DC38(self.f14, (d_2960_v79_).cf3, self.f14)
            index519_ = default__.safeIndex(525, (d_2958_v78_).length(0))
            (d_2958_v78_)[index519_] = ((self).f17) and ((d_2961_v80_).cf58)
            index520_ = default__.safeIndex(525, (d_2958_v78_).length(0))
            (d_2958_v78_)[index520_] = (d_2958_v78_)[default__.safeIndex(525, (d_2958_v78_).length(0))]
            d_2962_v81_: D11
            d_2962_v81_ = D11_DC20()
            source41_ = D11_DC21(d_2962_v81_)
            if source41_.is_DC20:
                d_2963_v82_: _dafny.MultiSet
                d_2963_v82_ = _dafny.MultiSet([False])
                d_2964_v83_: _dafny.Map
                d_2964_v83_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2965_i5_ in range(default__.abs(333))]): d_2963_v82_})
                r0 = default__.safeDivisionInt(default__.fm0(p1, globalState), len(d_2964_v83_))
                index521_ = default__.safeIndex(525, (d_2958_v78_).length(0))
                (d_2958_v78_)[index521_] = not(False)
                d_2966_v84_: _dafny.Seq
                d_2966_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fycuspo"))
                d_2967_v85_: C3
                nw590_ = C3()
                nw590_.ctor__(452, d_2966_v84_, self.f16, (self).f17)
                d_2967_v85_ = nw590_
                d_2968_v86_: _dafny.Map
                d_2968_v86_ = _dafny.Map({self.f16: d_2967_v85_})
                d_2968_v86_ = (d_2968_v86_).set((self).f17, d_2967_v85_)
                d_2969_v87_: _dafny.MultiSet
                d_2969_v87_ = _dafny.MultiSet([(d_2967_v85_).f23])
                d_2970_v88_: D21
                d_2970_v88_ = D21_DC44((d_2969_v87_).cardinality, p1)
                (self).f16 = (not(self.f16)) and ((d_2970_v88_).cf70)
            elif source41_.is_DC19:
                d_2971___mcc_h1_ = source41_.cf27
                d_2972_cf27_ = d_2971___mcc_h1_
                d_2973_v89_: _dafny.Seq
                d_2973_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtco"))
                d_2974_v90_: C3
                nw591_ = C3()
                nw591_.ctor__(self.f14, d_2973_v89_, self.f16, True)
                d_2974_v90_ = nw591_
                d_2974_v90_ = d_2974_v90_
                d_2975_v91_: int
                d_2975_v91_ = (0) - (len(d_2972_cf27_))
                d_2975_v91_ = d_2975_v91_
                d_2976_v92_: _dafny.Seq
                d_2976_v92_ = _dafny.SeqWithoutIsStrInference([p1, (self).f17])
                d_2977_v93_: _dafny.Array
                nw592_ = _dafny.Array(D9.default()(), 13)
                d_2977_v93_ = nw592_
                d_2978_v94_: _dafny.Seq
                d_2978_v94_ = _dafny.SeqWithoutIsStrInference([d_2977_v93_])
                d_2979_v95_: _dafny.Seq
                d_2979_v95_ = _dafny.SeqWithoutIsStrInference([d_2960_v79_, d_2960_v79_])
                d_2980_v96_: C2
                nw593_ = C2()
                nw593_.ctor__(d_2960_v79_, (d_2960_v79_).cf3, not((D12_DC24((self).f17, d_2976_v92_, (d_2978_v94_).set(default__.safeIndex((d_2974_v90_).f22, len(d_2978_v94_)), d_2977_v93_), (self).f17, (d_2979_v95_).set(default__.safeIndex(self.f14, len(d_2979_v95_)), D3_DC3((d_2958_v78_)[default__.safeIndex(525, (d_2958_v78_).length(0))])))).cf36))
                d_2980_v96_ = nw593_
                d_2981_v97_: D12
                d_2981_v97_ = D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efq")))
                d_2982_v98_: str
                d_2982_v98_ = _dafny.CodePoint('l')
                d_2983_v99_: _dafny.MultiSet
                d_2983_v99_ = _dafny.MultiSet([d_2982_v98_])
                d_2984_v100_: _dafny.MultiSet
                d_2984_v100_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsup"))])
                d_2985_v101_: _dafny.Set
                d_2985_v101_ = _dafny.Set({(d_2974_v90_).f22, len(d_2973_v89_)})
                d_2986_v102_: _dafny.MultiSet
                d_2986_v102_ = _dafny.MultiSet([p1])
                d_2987_v103_: _dafny.Array
                nw594_ = _dafny.Array(None, 29)
                nw594_[int(0)] = (0) - (self.f14)
                nw594_[int(1)] = p0
                nw594_[int(2)] = p0
                nw594_[int(3)] = len((d_2981_v97_).cf29)
                nw594_[int(4)] = self.f14
                nw594_[int(5)] = (d_2974_v90_).f22
                nw594_[int(6)] = self.f14
                nw594_[int(7)] = (d_2961_v80_).cf57
                nw594_[int(8)] = self.f14
                nw594_[int(9)] = (d_2974_v90_).f22
                nw594_[int(10)] = ((d_2983_v99_)[_dafny.CodePoint('u')] if (_dafny.CodePoint('u')) in (d_2983_v99_) else (d_2984_v100_).cardinality)
                nw594_[int(11)] = (0) - (default__.fm0(p1, globalState))
                nw594_[int(12)] = p0
                nw594_[int(13)] = self.f14
                nw594_[int(14)] = (_dafny.MultiSet([d_2985_v101_, d_2985_v101_])).cardinality
                nw594_[int(15)] = ((d_2974_v90_).f22) * (p0)
                nw594_[int(16)] = 29
                nw594_[int(17)] = (((d_2907_v34_)[p0] if (p0) in (d_2907_v34_) else len(_dafny.SeqWithoutIsStrInference([(d_2986_v102_).cardinality])))) - (p0)
                nw594_[int(18)] = (d_2974_v90_).f22
                nw594_[int(19)] = (0) - ((self.f14) * (p0))
                nw594_[int(20)] = self.f14
                nw594_[int(21)] = (d_2974_v90_).f22
                nw594_[int(22)] = ((d_2907_v34_)[p0] if (p0) in (d_2907_v34_) else 915)
                nw594_[int(23)] = 872
                nw594_[int(24)] = (p0) + (self.f14)
                nw594_[int(25)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byv"))) + (_dafny.SeqWithoutIsStrInference([d_2982_v98_ for d_2988_i6_ in range(default__.abs(418))])))
                nw594_[int(26)] = self.f14
                nw594_[int(27)] = (self.f14) * (self.f14)
                nw594_[int(28)] = default__.safeDivisionInt(self.f14, len(_dafny.Map({True: d_2986_v102_})))
                d_2987_v103_ = nw594_
                index522_ = default__.safeIndex(635, (d_2987_v103_).length(0))
                (d_2987_v103_)[index522_] = ((d_2974_v90_).f22) * (len((d_2974_v90_).f23))
                index523_ = default__.safeIndex(635, (d_2987_v103_).length(0))
                (d_2987_v103_)[index523_] = (d_2974_v90_).f22
            elif True:
                d_2989___mcc_h2_ = source41_.cf28
                d_2990_cf28_ = d_2989___mcc_h2_
                d_2913_v38_ = (d_2913_v38_).set((d_2958_v78_)[default__.safeIndex(525, (d_2958_v78_).length(0))], (d_2958_v78_)[default__.safeIndex(525, (d_2958_v78_).length(0))])
                d_2991_v104_: _dafny.Seq
                d_2991_v104_ = _dafny.SeqWithoutIsStrInference([(self.f14) + (p0)])
                d_2992_v105_: _dafny.Array
                nw595_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_2992_v105_ = nw595_
                d_2993_v106_: _dafny.Seq
                d_2993_v106_ = _dafny.SeqWithoutIsStrInference([d_2992_v105_])
                rhs518_ = d_2991_v104_
                rhs519_ = (((d_2993_v106_).set(default__.safeIndex(self.f14, len(d_2993_v106_)), d_2992_v105_)) + (d_2993_v106_)) + (((d_2993_v106_) + (d_2993_v106_)).set(default__.safeIndex(p0, len((d_2993_v106_) + (d_2993_v106_))), d_2992_v105_))
                d_2991_v104_ = rhs518_
                d_2993_v106_ = rhs519_
                d_2994_v107_: _dafny.Array
                nw596_ = _dafny.Array(int(0), 28)
                d_2994_v107_ = nw596_
                index524_ = default__.safeIndex(505, (d_2994_v107_).length(0))
                (d_2994_v107_)[index524_] = default__.safeModuloInt(self.f14, default__.fm0((self).f17, globalState))
                index525_ = default__.safeIndex(505, (d_2994_v107_).length(0))
                (d_2994_v107_)[index525_] = self.f14
                d_2995_v108_: _dafny.Seq
                d_2995_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etgkflwva"))
                d_2996_v109_: str
                d_2996_v109_ = _dafny.CodePoint('i')
                d_2997_v110_: _dafny.MultiSet
                d_2997_v110_ = _dafny.MultiSet([(d_2995_v108_ if not((self).fm4(len(d_2991_v104_), (self).f17, (d_2994_v107_)[default__.safeIndex(505, (d_2994_v107_).length(0))], globalState)) else d_2995_v108_), (d_2995_v108_).set(default__.safeIndex(self.f14, len(d_2995_v108_)), d_2996_v109_), d_2995_v108_, _dafny.SeqWithoutIsStrInference([d_2996_v109_ for d_2998_i7_ in range(default__.abs(845))])])
                d_2997_v110_ = _dafny.MultiSet([(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_2999_i8_ in range(default__.abs(806))])).set(default__.safeIndex(self.f14, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3000_i8_ in range(default__.abs(806))]))), d_2996_v109_)).set(default__.safeIndex(self.f14, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3001_i8_ in range(default__.abs(806))])).set(default__.safeIndex(self.f14, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3002_i8_ in range(default__.abs(806))]))), d_2996_v109_))), d_2996_v109_)) + (d_2995_v108_), (d_2995_v108_) + (d_2995_v108_), _dafny.SeqWithoutIsStrInference([d_2996_v109_ for d_3003_i9_ in range(default__.abs(967))]), d_2995_v108_])
            d_3004_v111_: str
            d_3004_v111_ = _dafny.CodePoint('l')
            d_3005_v112_: _dafny.Seq
            d_3005_v112_ = _dafny.SeqWithoutIsStrInference([p0])
            d_3006_v113_: _dafny.Map
            d_3006_v113_ = _dafny.Map({len(d_3005_v112_): d_3004_v111_})
            d_3004_v111_ = (D6_DC9(((d_3006_v113_)[self.f14] if (self.f14) in (d_3006_v113_) else d_3004_v111_))).cf14
        elif True:
            (self).f16 = self.f16
            d_3007_v114_: _dafny.Set
            d_3007_v114_ = _dafny.Set({default__.fm1(globalState)})
            d_3007_v114_ = d_3007_v114_
            d_3008_v115_: _dafny.Array
            nw597_ = _dafny.Array(False, 26)
            d_3008_v115_ = nw597_
            d_3009_v116_: _dafny.Seq
            d_3009_v116_ = _dafny.SeqWithoutIsStrInference([d_3008_v115_, d_3008_v115_, d_3008_v115_, d_3008_v115_])
            d_3010_v117_: T0
            nw598_ = C1()
            nw598_.ctor__(self.f16, (0) - (p0))
            d_3010_v117_ = nw598_
            d_3011_v118_: D20
            d_3011_v118_ = D20_DC42(self.f14, self.f16, d_3010_v117_)
            d_3012_v119_: _dafny.Set
            d_3012_v119_ = _dafny.Set({self.f14})
            d_3013_v120_: C10
            nw599_ = C10()
            nw599_.ctor__(p1, (self).f17, len(d_3012_v119_))
            d_3013_v120_ = nw599_
            d_3014_v121_: _dafny.Set
            d_3014_v121_ = _dafny.Set({d_3013_v120_})
            d_3015_v122_: _dafny.Array
            nw600_ = _dafny.Array(None, 9)
            nw600_[int(0)] = self.f14
            nw600_[int(1)] = default__.fm0(self.f16, globalState)
            nw600_[int(2)] = len(default__.fm6((0) - ((0) - ((0) - (self.f14))), False, p1, (d_3011_v118_).cf66, globalState))
            nw600_[int(3)] = len(_dafny.SeqWithoutIsStrInference([len(d_3014_v121_)]))
            nw600_[int(4)] = (p0) + (d_3010_v117_.f14)
            nw600_[int(5)] = (self.f14) + (d_3010_v117_.f14)
            nw600_[int(6)] = p0
            nw600_[int(7)] = p0
            nw600_[int(8)] = p0
            d_3015_v122_ = nw600_
            index526_ = default__.safeIndex(396, (d_3015_v122_).length(0))
            (d_3015_v122_)[index526_] = 387
            d_3016_v123_: _dafny.Seq
            d_3016_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "piced"))
            d_3017_v124_: _dafny.Seq
            d_3017_v124_ = _dafny.SeqWithoutIsStrInference([d_3010_v117_.f14])
            d_3018_v126_: _dafny.Map
            d_3018_v126_ = _dafny.Map({False: d_2913_v38_})
            d_3019_v127_: _dafny.Map
            d_3019_v127_ = (d_3018_v126_).set(self.f16, d_2913_v38_)
            d_3020_v128_: _dafny.Map
            d_3020_v128_ = _dafny.Map({(self).f17: d_3019_v127_})
            d_3021_v129_: _dafny.Map
            d_3021_v129_ = _dafny.Map({D14_DC27(_dafny.Map({d_3020_v128_: self.f14})): -956})
            index527_ = default__.safeIndex(396, (d_3015_v122_).length(0))
            def iife232_():
                coll100_ = _dafny.Map()
                compr_100_: D14
                for compr_100_ in (d_3021_v129_).keys.Elements:
                    d_3022_v125_: D14 = compr_100_
                    if (d_3022_v125_) in (d_3021_v129_):
                        coll100_[d_3022_v125_] = (self).f17
                return _dafny.Map(coll100_)
            rhs520_ = d_3009_v116_
            rhs521_ = 837
            rhs522_ = default__.fm27(len(d_3012_v119_), len(d_3016_v123_), (d_3017_v124_)[default__.safeIndex(d_3010_v117_.f14, len(d_3017_v124_))], _dafny.Map({self.f14: self.f14}), globalState)
            rhs523_ = (self).f17
            rhs524_ = ((self.f14) * ((0) - (len(iife232_()
            )))) * (-76)
            lhs405_ = globalState
            lhs406_ = self
            lhs407_ = self
            lhs408_ = d_3015_v122_
            lhs409_ = default__.safeIndex(396, (d_3015_v122_).length(0))
            d_3009_v116_ = rhs520_
            lhs405_.f7 = rhs521_
            lhs406_.f16 = rhs522_
            lhs407_.f16 = rhs523_
            lhs408_[lhs409_] = rhs524_
            index528_ = default__.safeIndex(141, (d_3008_v115_).length(0))
            (d_3008_v115_)[index528_] = (self).f17
            index529_ = default__.safeIndex(141, (d_3008_v115_).length(0))
            (d_3008_v115_)[index529_] = self.f16
            d_3017_v124_ = (d_3017_v124_) + ((d_3017_v124_ if self.f16 else d_3017_v124_))
        d_3023_v130_: _dafny.Seq
        d_3023_v130_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttsybhds"))
        d_3024_v131_: C3
        nw601_ = C3()
        nw601_.ctor__(p0, d_3023_v130_, self.f16, p1)
        d_3024_v131_ = nw601_
        d_3025_v132_: _dafny.Seq
        d_3025_v132_ = _dafny.SeqWithoutIsStrInference([(d_3024_v131_).f22, (d_3024_v131_).f22])
        hi26_ = default__.safeDivisionInt(486, len(d_3025_v132_))
        for d_3026_i10_ in range((0) - (p0), hi26_):
            d_3027_v133_: _dafny.Array
            def lambda176_(d_3028_i11_):
                return _dafny.CodePoint('o')

            init103_ = lambda176_
            nw602_ = _dafny.Array(None, 1)
            for i0_103_ in range(nw602_.length(0)):
                nw602_[i0_103_] = init103_(i0_103_)
            d_3027_v133_ = nw602_
            def lambda177_(d_3029_i12_):
                return _dafny.CodePoint('n')

            init104_ = lambda177_
            nw603_ = _dafny.Array(None, 14)
            for i0_104_ in range(nw603_.length(0)):
                nw603_[i0_104_] = init104_(i0_104_)
            d_3027_v133_ = nw603_
            d_3030_v134_: _dafny.Set
            d_3030_v134_ = _dafny.Set({d_3026_i10_})
            d_3030_v134_ = d_3030_v134_
            (globalState).f2 = (self.f14) + (default__.safeDivisionInt((d_3024_v131_).f22, len((d_3024_v131_).f23)))
            d_2907_v34_ = (d_2907_v34_) - (d_2907_v34_)
        r0 = (p0) * (176)
        return r0

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_3031_v0_: _dafny.Map
        d_3031_v0_ = _dafny.Map({p3: (p1) - (p3)})
        d_3032_v1_: _dafny.Seq
        d_3032_v1_ = _dafny.SeqWithoutIsStrInference([p1])
        d_3031_v0_ = (d_3031_v0_).set((self.f14) - (len(d_3031_v0_)), (d_3032_v1_)[default__.safeIndex(self.f14, len(d_3032_v1_))])
        d_3033_v2_: _dafny.Seq
        d_3033_v2_ = _dafny.SeqWithoutIsStrInference([(self).f17, self.f16])
        d_3034_v3_: _dafny.Seq
        d_3034_v3_ = _dafny.SeqWithoutIsStrInference([d_3033_v2_, d_3033_v2_, d_3033_v2_, d_3033_v2_, d_3033_v2_])
        (self).f16 = ((0) - (len(d_3034_v3_))) != ((0) - (p1))
        d_3035_v4_: _dafny.Map
        d_3035_v4_ = _dafny.Map({self.f14: p0})
        d_3035_v4_ = (d_3035_v4_).set(p3, _dafny.CodePoint('a'))
        d_3036_v6_: _dafny.Array
        def lambda178_(d_3037_p0_, d_3038_p1_):
            def lambda179_(d_3039_i1_):
                def iife233_():
                    coll101_ = _dafny.Map()
                    compr_101_: int
                    for compr_101_ in _dafny.IntegerRange(939, 588):
                        d_3040_v5_: int = compr_101_
                        if ((939) <= (d_3040_v5_)) and ((d_3040_v5_) < (588)):
                            coll101_[(d_3040_v5_) - ((0) - (d_3038_p1_))] = (self).f17
                    return _dafny.Map(coll101_)
                return (iife233_()
                ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_3037_p0_ for d_3041_i2_ in range(default__.abs(333))])): (self).f17}))

            return lambda179_

        init105_ = lambda178_(p0, p1)
        nw604_ = _dafny.Array(None, 11)
        for i0_105_ in range(nw604_.length(0)):
            nw604_[i0_105_] = init105_(i0_105_)
        d_3036_v6_ = nw604_
        guard_loop_14_: int
        for guard_loop_14_ in _dafny.IntegerRange(0, (d_3036_v6_).length(0)):
            d_3042_i0_: int = guard_loop_14_
            if (True) and (((0) <= (d_3042_i0_)) and ((d_3042_i0_) < ((d_3036_v6_).length(0)))):
                (d_3036_v6_)[(d_3042_i0_)] = (_dafny.Map({p1: (self).f17})) | (_dafny.Map({self.f14: not(not((self).f17))}))
        (self).f16 = self.f16
        (self).f16 = ((default__.fm0(p2, globalState)) * (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({-949})])))) < (p1)
        r0 = self.f14
        return r0


class C12(T0):
    def  __init__(self):
        self._f14: int = int(0)
        self._f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    @property
    def f14(self):
        return self._f14
    @f14.setter
    def f14(self, value):
        self._f14 = value
    def ctor__(self, f15, f14):
        (self)._f15 = f15
        (self).f14 = f14

    def fm2(self, p0, p1, p2, p3, globalState):
        return self.f14

    def m1(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_3043_v0_: _dafny.Map
        d_3043_v0_ = _dafny.Map({(self).f15: (self).f15})
        hi27_ = len(d_3043_v0_)
        for d_3044_i0_ in range(self.f14, hi27_):
            d_3045_v1_: _dafny.Array
            def lambda180_(d_3046_v0_):
                def lambda181_(d_3047_i1_):
                    return ((self).f15 if True else ((d_3046_v0_)[(self).f15] if ((self).f15) in (d_3046_v0_) else (self).f15))

                return lambda181_

            init106_ = lambda180_(d_3043_v0_)
            nw605_ = _dafny.Array(None, 22)
            for i0_106_ in range(nw605_.length(0)):
                nw605_[i0_106_] = init106_(i0_106_)
            d_3045_v1_ = nw605_
            index530_ = default__.safeIndex(311, (d_3045_v1_).length(0))
            (d_3045_v1_)[index530_] = p0
            index531_ = default__.safeIndex(311, (d_3045_v1_).length(0))
            (d_3045_v1_)[index531_] = p0
            d_3048_v2_: _dafny.Seq
            d_3048_v2_ = _dafny.SeqWithoutIsStrInference([(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))], p0])
            d_3049_v3_: _dafny.Seq
            d_3049_v3_ = d_3048_v2_
            d_3050_v4_: _dafny.Map
            d_3050_v4_ = _dafny.Map({len((d_3049_v3_)): len(d_3048_v2_)})
            d_3051_v5_: _dafny.Map
            d_3051_v5_ = _dafny.Map({d_3050_v4_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbsvqm"))})
            r2 = (len(d_3051_v5_)) + (d_3044_i0_)
            if p0:
                d_3052_v6_: _dafny.Array
                def lambda182_(d_3053_i0_):
                    def lambda183_(d_3054_i2_):
                        return (d_3054_i2_) + (d_3053_i0_)

                    return lambda183_

                init107_ = lambda182_(d_3044_i0_)
                nw606_ = _dafny.Array(None, 6)
                for i0_107_ in range(nw606_.length(0)):
                    nw606_[i0_107_] = init107_(i0_107_)
                d_3052_v6_ = nw606_
                index532_ = default__.safeIndex(486, (d_3052_v6_).length(0))
                (d_3052_v6_)[index532_] = d_3044_i0_
                index533_ = default__.safeIndex(486, (d_3052_v6_).length(0))
                (d_3052_v6_)[index533_] = self.f14
                index534_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                (d_3045_v1_)[index534_] = p0
                d_3055_v7_: _dafny.Array
                nw607_ = _dafny.Array(None, 1)
                nw607_[int(0)] = d_3052_v6_
                d_3055_v7_ = nw607_
                index535_ = default__.safeIndex(981, (d_3055_v7_).length(0))
                (d_3055_v7_)[index535_] = d_3052_v6_
                index536_ = default__.safeIndex(981, (d_3055_v7_).length(0))
                (d_3055_v7_)[index536_] = d_3052_v6_
                d_3056_v8_: _dafny.Seq
                d_3056_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goxhiusc"))
                d_3057_v9_: _dafny.Set
                d_3057_v9_ = _dafny.Set({(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))], False, (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))], (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))], (self).f15})
                r2 = ((d_3052_v6_)[default__.safeIndex(486, (d_3052_v6_).length(0))]) + ((0) - (default__.safeModuloInt(len(d_3056_v8_), len(d_3057_v9_))))
                index537_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                (d_3045_v1_)[index537_] = not ((self).f15) or (not (p0) or ((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]))
            elif True:
                (globalState).f7 = d_3044_i0_
                d_3058_v10_: T0
                nw608_ = C10()
                nw608_.ctor__(p0, (self).f15, d_3044_i0_)
                d_3058_v10_ = nw608_
                d_3059_v11_: _dafny.Seq
                d_3059_v11_ = _dafny.SeqWithoutIsStrInference([d_3058_v10_])
                (globalState).f2 = default__.safeDivisionInt(d_3044_i0_, (self).fm2(len(d_3059_v11_), d_3058_v10_.f14, default__.fm21(globalState), d_3058_v10_.f14, globalState))
                index538_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                (d_3045_v1_)[index538_] = (d_3058_v10_.f14) == (self.f14)
                d_3060_v12_: C4
                nw609_ = C4()
                nw609_.ctor__(not((not((self).f15)) == ((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))])), False, (self).f15)
                d_3060_v12_ = nw609_
                index539_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                (d_3045_v1_)[index539_] = not(False)
            if (self).f15:
                d_3061_v13_: _dafny.Map
                d_3061_v13_ = _dafny.Map({default__.fm0((self).f15, globalState): p0})
                d_3061_v13_ = default__.fm22((self).f15, globalState)
                d_3062_v14_: _dafny.Array
                nw610_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_3062_v14_ = nw610_
                d_3063_v15_: _dafny.MultiSet
                d_3063_v15_ = _dafny.MultiSet([((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))] if (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))] else (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))])])
                index540_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                rhs525_ = ((d_3063_v15_)[False] if (False) in (d_3063_v15_) else d_3044_i0_)
                rhs526_ = (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]
                rhs527_ = d_3062_v14_
                rhs528_ = self.f14
                lhs410_ = globalState
                lhs411_ = d_3045_v1_
                lhs412_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                lhs413_ = globalState
                lhs410_.f2 = rhs525_
                lhs411_[lhs412_] = rhs526_
                d_3062_v14_ = rhs527_
                lhs413_.f2 = rhs528_
                d_3064_v16_: _dafny.Array
                def lambda184_(d_3065_i3_):
                    return _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15])

                init108_ = lambda184_
                nw611_ = _dafny.Array(None, 16)
                for i0_108_ in range(nw611_.length(0)):
                    nw611_[i0_108_] = init108_(i0_108_)
                d_3064_v16_ = nw611_
                index541_ = default__.safeIndex(348, (d_3064_v16_).length(0))
                (d_3064_v16_)[index541_] = (d_3048_v2_ if (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))] else d_3048_v2_)
                d_3066_v17_: _dafny.Map
                d_3066_v17_ = _dafny.Map({self.f14: d_3048_v2_})
                d_3067_v18_: _dafny.Map
                d_3067_v18_ = _dafny.Map({d_3048_v2_: _dafny.SeqWithoutIsStrInference([(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]])})
                d_3068_v19_: _dafny.Map
                d_3068_v19_ = _dafny.Map({(self).f15: d_3048_v2_})
                index542_ = default__.safeIndex(348, (d_3064_v16_).length(0))
                (d_3064_v16_)[index542_] = ((d_3066_v17_)[d_3044_i0_] if (d_3044_i0_) in (d_3066_v17_) else (((d_3067_v18_)[((d_3068_v19_)[(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]] if ((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]) in (d_3068_v19_) else d_3048_v2_)] if (((d_3068_v19_)[(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]] if ((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]) in (d_3068_v19_) else d_3048_v2_)) in (d_3067_v18_) else (d_3048_v2_).set(default__.safeIndex(767, len(d_3048_v2_)), (self).f15))) + (d_3048_v2_))
                d_3069_v20_: _dafny.Array
                nw612_ = _dafny.Array(int(0), 16)
                d_3069_v20_ = nw612_
                index543_ = default__.safeIndex(914, (d_3069_v20_).length(0))
                (d_3069_v20_)[index543_] = self.f14
                index544_ = default__.safeIndex(168, (d_3069_v20_).length(0))
                (d_3069_v20_)[index544_] = d_3044_i0_
                index545_ = default__.safeIndex(914, (d_3069_v20_).length(0))
                index546_ = default__.safeIndex(168, (d_3069_v20_).length(0))
                rhs529_ = self.f14
                rhs530_ = d_3044_i0_
                lhs414_ = d_3069_v20_
                lhs415_ = default__.safeIndex(914, (d_3069_v20_).length(0))
                lhs416_ = d_3069_v20_
                lhs417_ = default__.safeIndex(168, (d_3069_v20_).length(0))
                lhs414_[lhs415_] = rhs529_
                lhs416_[lhs417_] = rhs530_
                index547_ = default__.safeIndex(914, (d_3069_v20_).length(0))
                (d_3069_v20_)[index547_] = d_3044_i0_
            elif True:
                d_3070_v21_: _dafny.Set
                d_3070_v21_ = _dafny.Set({(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]})
                d_3071_v22_: _dafny.MultiSet
                d_3071_v22_ = _dafny.MultiSet([(self).f15, (self).f15, (self).f15, p0])
                d_3072_v23_: _dafny.Seq
                d_3072_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skoh"))
                d_3073_v24_: str
                d_3073_v24_ = _dafny.CodePoint('w')
                d_3074_v25_: _dafny.Array
                nw613_ = _dafny.Array(None, 11)
                nw613_[int(0)] = ((d_3071_v22_)[(d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]] if ((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]) in (d_3071_v22_) else d_3044_i0_)
                nw613_[int(1)] = d_3044_i0_
                nw613_[int(2)] = len(_dafny.Map({d_3044_i0_: (d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]}))
                nw613_[int(3)] = d_3044_i0_
                nw613_[int(4)] = len((d_3072_v23_).set(default__.safeIndex(self.f14, len(d_3072_v23_)), d_3073_v24_))
                nw613_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_3044_i0_ for d_3075_i4_ in range(default__.abs(15))]))
                nw613_[int(6)] = self.f14
                nw613_[int(7)] = self.f14
                nw613_[int(8)] = d_3044_i0_
                nw613_[int(9)] = d_3044_i0_
                nw613_[int(10)] = -349
                d_3074_v25_ = nw613_
                d_3076_v26_: _dafny.Map
                d_3076_v26_ = _dafny.Map({d_3070_v21_: d_3074_v25_})
                index548_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                (d_3045_v1_)[index548_] = default__.fm27(d_3044_i0_, d_3044_i0_, default__.fm0(not((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]), globalState), _dafny.Map({self.f14: len(d_3076_v26_)}), globalState)
                (globalState).f7 = d_3044_i0_
                d_3070_v21_ = (d_3070_v21_) | (d_3070_v21_)
                index549_ = default__.safeIndex(486, (d_3074_v25_).length(0))
                (d_3074_v25_)[index549_] = d_3044_i0_
                index550_ = default__.safeIndex(486, (d_3074_v25_).length(0))
                index551_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                index552_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                rhs531_ = 709
                rhs532_ = not(default__.fm43((d_3071_v22_).ispropersubset(default__.fm1(globalState)), globalState))
                rhs533_ = ((d_3045_v1_)[default__.safeIndex(311, (d_3045_v1_).length(0))]) and (not(default__.fm21(globalState)))
                lhs418_ = d_3074_v25_
                lhs419_ = default__.safeIndex(486, (d_3074_v25_).length(0))
                lhs420_ = d_3045_v1_
                lhs421_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                lhs422_ = d_3045_v1_
                lhs423_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                lhs418_[lhs419_] = rhs531_
                lhs420_[lhs421_] = rhs532_
                lhs422_[lhs423_] = rhs533_
                index553_ = default__.safeIndex(311, (d_3045_v1_).length(0))
                (d_3045_v1_)[index553_] = ((d_3074_v25_)[default__.safeIndex(486, (d_3074_v25_).length(0))]) >= ((d_3044_i0_ if default__.fm21(globalState) else (d_3074_v25_)[default__.safeIndex(486, (d_3074_v25_).length(0))]))
        (globalState).f7 = default__.safeModuloInt(self.f14, self.f14)
        d_3077_v27_: _dafny.Map
        d_3077_v27_ = _dafny.Map({default__.fm17(not(not((self).f15)), (0) - (self.f14), self.f14, self.f14, globalState): self.f14})
        d_3078_v28_: _dafny.Seq
        d_3078_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
        d_3079_v29_: _dafny.Seq
        d_3079_v29_ = _dafny.SeqWithoutIsStrInference([(self).f15])
        d_3080_v30_: _dafny.Map
        d_3080_v30_ = _dafny.Map({(0) - (len(d_3078_v28_)): d_3079_v29_})
        d_3077_v27_ = (d_3077_v27_).set((self).f15, (len(d_3080_v30_) if (self).f15 else self.f14))
        d_3081_v31_: _dafny.Seq
        d_3081_v31_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f14), self.f14, self.f14])
        nw614_ = _dafny.Array(None, 3)
        nw614_[int(0)] = (d_3081_v31_) != (d_3081_v31_)
        nw614_[int(1)] = (self).f15
        nw614_[int(2)] = ((self).f15 if (self).f15 else p0)
        r1 = nw614_
        (globalState).f7 = ((self.f14) - (926)) - (405)
        d_3082_v32_: bool
        d_3082_v32_ = True
        d_3083_v33_: str
        d_3083_v33_ = _dafny.CodePoint('c')
        d_3084_v34_: _dafny.Map
        d_3084_v34_ = _dafny.Map({self.f14: self.f14})
        d_3085_v35_: D11
        d_3085_v35_ = D11_DC19(d_3084_v34_)
        d_3086_v36_: D11
        d_3086_v36_ = D11_DC21(d_3085_v35_)
        pat_let_tv49_ = d_3082_v32_
        def lambda185_(source42_):
            if source42_.is_DC20:
                return self.f14
            elif source42_.is_DC19:
                d_3087___mcc_h0_ = source42_.cf27
                d_3088_cf27_ = d_3087___mcc_h0_
                return (D18_DC38(self.f14, pat_let_tv49_, len(d_3088_cf27_))).cf59
            elif True:
                d_3089___mcc_h1_ = source42_.cf28
                d_3090_cf28_ = d_3089___mcc_h1_
                return self.f14

        rhs534_ = (self).f15
        rhs535_ = lambda185_(d_3086_v36_)
        rhs536_ = (self.f14) + (self.f14)
        rhs537_ = d_3083_v33_
        rhs538_ = d_3078_v28_
        lhs424_ = self
        d_3082_v32_ = rhs534_
        r2 = rhs535_
        lhs424_.f14 = rhs536_
        d_3083_v33_ = rhs537_
        d_3078_v28_ = rhs538_
        r0 = (((d_3081_v31_) + (d_3081_v31_)).set(default__.safeIndex(self.f14, len((d_3081_v31_) + (d_3081_v31_))), len(d_3078_v28_))) + (d_3081_v31_)
        nw615_ = _dafny.Array(None, 1)
        nw615_[int(0)] = p0
        r1 = nw615_
        r2 = self.f14
        return r0, r1, r2

    @property
    def f15(self):
        return self._f15

class C13:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    def ctor__(self):
        pass
        pass

    def m0(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        d_3091_v0_: int
        d_3091_v0_ = -653
        (globalState).f7 = default__.safeDivisionInt(d_3091_v0_, d_3091_v0_)
        if p0:
            d_3092_v1_: _dafny.MultiSet
            d_3092_v1_ = _dafny.MultiSet([p0])
            d_3093_v2_: _dafny.Map
            d_3093_v2_ = _dafny.Map({p0: d_3091_v0_})
            d_3094_v3_: _dafny.Map
            d_3094_v3_ = _dafny.Map({d_3092_v1_: len((d_3093_v2_) | (_dafny.Map({False: d_3091_v0_})))})
            d_3094_v3_ = (d_3094_v3_).set(d_3092_v1_, (0) - (default__.fm0(p0, globalState)))
            if p0:
                d_3095_v4_: _dafny.Map
                d_3095_v4_ = _dafny.Map({d_3091_v0_: d_3091_v0_})
                d_3096_v5_: _dafny.Map
                d_3096_v5_ = _dafny.Map({p0: p0})
                d_3097_v6_: _dafny.Map
                d_3097_v6_ = _dafny.Map({len(d_3096_v5_): p0})
                d_3098_v7_: _dafny.Array
                nw616_ = _dafny.Array(None, 8)
                nw616_[int(0)] = (_dafny.MultiSet([_dafny.Map({len(d_3095_v4_): p0}), (d_3097_v6_).set((0) - (d_3091_v0_), p0), _dafny.Map({d_3091_v0_: False}), _dafny.Map({d_3091_v0_: p0}), d_3097_v6_])).cardinality
                nw616_[int(1)] = 761
                nw616_[int(2)] = default__.safeDivisionInt(102, d_3091_v0_)
                nw616_[int(3)] = (d_3091_v0_) - (d_3091_v0_)
                nw616_[int(4)] = d_3091_v0_
                nw616_[int(5)] = (d_3091_v0_) - (d_3091_v0_)
                nw616_[int(6)] = (0) - (d_3091_v0_)
                nw616_[int(7)] = d_3091_v0_
                d_3098_v7_ = nw616_
                d_3098_v7_ = (d_3098_v7_ if False else d_3098_v7_)
                d_3099_v8_: _dafny.Array
                def lambda186_(d_3100_p0_):
                    def lambda187_(d_3101_i0_):
                        return (_dafny.Set({d_3100_p0_})).ispropersubset(_dafny.Set({d_3100_p0_}))

                    return lambda187_

                init109_ = lambda186_(p0)
                nw617_ = _dafny.Array(None, 18)
                for i0_109_ in range(nw617_.length(0)):
                    nw617_[i0_109_] = init109_(i0_109_)
                d_3099_v8_ = nw617_
                index554_ = default__.safeIndex(753, (d_3099_v8_).length(0))
                (d_3099_v8_)[index554_] = p0
                d_3102_v9_: _dafny.Seq
                d_3102_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft"))
                index555_ = default__.safeIndex(753, (d_3099_v8_).length(0))
                rhs539_ = d_3091_v0_
                rhs540_ = p0
                rhs541_ = d_3098_v7_
                rhs542_ = (not(p0) if p0 else (d_3102_v9_) < (d_3102_v9_))
                lhs425_ = d_3099_v8_
                lhs426_ = default__.safeIndex(753, (d_3099_v8_).length(0))
                d_3091_v0_ = rhs539_
                lhs425_[lhs426_] = rhs540_
                d_3098_v7_ = rhs541_
                r1 = rhs542_
                r1 = (d_3099_v8_)[default__.safeIndex(753, (d_3099_v8_).length(0))]
                d_3103_v10_: C6
                nw618_ = C6()
                nw618_.ctor__(d_3091_v0_, p0, default__.fm17((d_3099_v8_)[default__.safeIndex(753, (d_3099_v8_).length(0))], d_3091_v0_, d_3091_v0_, d_3091_v0_, globalState), (d_3099_v8_)[default__.safeIndex(753, (d_3099_v8_).length(0))], d_3091_v0_)
                d_3103_v10_ = nw618_
                (globalState).f2 = d_3091_v0_
            elif True:
                d_3104_v11_: _dafny.Seq
                d_3104_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkeaspy"))
                d_3105_v12_: _dafny.MultiSet
                d_3105_v12_ = _dafny.MultiSet([(d_3104_v11_) + (d_3104_v11_), d_3104_v11_])
                d_3105_v12_ = (d_3105_v12_).set(d_3104_v11_, default__.abs(d_3091_v0_))
                d_3106_v13_: D12
                d_3106_v13_ = D12_DC22(d_3104_v11_)
                d_3107_v14_: _dafny.Seq
                d_3107_v14_ = _dafny.SeqWithoutIsStrInference([p0])
                (globalState).f2 = (-285) * ((0) - ((len((d_3106_v13_).cf29)) + (len(d_3107_v14_))))
                (globalState).f7 = d_3091_v0_
                d_3108_v15_: _dafny.Array
                nw619_ = _dafny.Array(int(0), 25)
                d_3108_v15_ = nw619_
                index556_ = default__.safeIndex(482, (d_3108_v15_).length(0))
                (d_3108_v15_)[index556_] = d_3091_v0_
                index557_ = default__.safeIndex(482, (d_3108_v15_).length(0))
                (d_3108_v15_)[index557_] = d_3091_v0_
                d_3107_v14_ = d_3107_v14_
            d_3109_v16_: _dafny.Seq
            d_3109_v16_ = _dafny.SeqWithoutIsStrInference([not(p0), p0])
            d_3109_v16_ = d_3109_v16_
            d_3110_v17_: str
            d_3110_v17_ = _dafny.CodePoint('s')
            d_3111_v18_: D6
            d_3111_v18_ = D6_DC9(d_3110_v17_)
            d_3111_v18_ = d_3111_v18_
            d_3112_v19_: _dafny.Seq
            d_3112_v19_ = _dafny.SeqWithoutIsStrInference([802])
            d_3112_v19_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p0, globalState)])
        elif True:
            r1 = p0
            d_3113_v20_: _dafny.Set
            d_3113_v20_ = _dafny.Set({False, p0})
            d_3114_v21_: D15
            d_3114_v21_ = D15_DC29(default__.fm44(False, not(p0), globalState))
            d_3115_v22_: T0
            nw620_ = C7()
            nw620_.ctor__(p0, p0, len(default__.fm70(d_3114_v21_, p0, d_3091_v0_, globalState)))
            d_3115_v22_ = nw620_
            d_3116_v23_: D14
            d_3116_v23_ = D14_DC28(len(d_3113_v20_), not(not(p0)), d_3115_v22_)
            (globalState).f2 = (0) - ((0) - ((d_3116_v23_).cf40))
            r2 = p0
            d_3117_v24_: _dafny.MultiSet
            d_3117_v24_ = _dafny.MultiSet([d_3091_v0_])
            d_3118_v25_: _dafny.Map
            d_3118_v25_ = _dafny.Map({not(p0): True})
            d_3119_v26_: C5
            nw621_ = C5()
            nw621_.ctor__((d_3117_v24_).issubset(d_3117_v24_), (p0) or (((d_3118_v25_)[p0] if (p0) in (d_3118_v25_) else p0)))
            d_3119_v26_ = nw621_
            d_3120_v27_: _dafny.Map
            d_3120_v27_ = _dafny.Map({((d_3118_v25_)[p0] if (p0) in (d_3118_v25_) else p0): _dafny.Map({p0: d_3118_v25_})})
            d_3121_v28_: _dafny.Map
            d_3121_v28_ = _dafny.Map({d_3120_v27_: d_3115_v22_.f14})
            d_3122_v29_: str
            d_3122_v29_ = _dafny.CodePoint('m')
            d_3123_v30_: _dafny.Map
            d_3123_v30_ = _dafny.Map({D14_DC27(d_3121_v28_): d_3122_v29_})
            d_3123_v30_ = d_3123_v30_
        d_3124_v31_: C4
        nw622_ = C4()
        nw622_.ctor__(p0, p0, not(p0))
        d_3124_v31_ = nw622_
        d_3125_v32_: _dafny.Set
        d_3125_v32_ = _dafny.Set({d_3124_v31_})
        d_3126_v33_: D13
        d_3126_v33_ = D13_DC26()
        d_3127_v34_: _dafny.Map
        d_3127_v34_ = _dafny.Map({d_3125_v32_: d_3126_v33_})
        d_3128_v35_: _dafny.Seq
        d_3128_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tm"))
        d_3129_v36_: _dafny.Map
        d_3129_v36_ = _dafny.Map({d_3091_v0_: d_3091_v0_})
        d_3130_v37_: _dafny.Seq
        d_3130_v37_ = _dafny.SeqWithoutIsStrInference([d_3091_v0_, d_3091_v0_])
        d_3131_v38_: _dafny.Set
        d_3131_v38_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_3091_v0_])})
        d_3132_v39_: D12
        d_3132_v39_ = D12_DC23(p0, len(d_3131_v38_), len(d_3130_v37_))
        d_3133_v40_: _dafny.Array
        nw623_ = _dafny.Array(None, 16)
        nw623_[int(0)] = default__.safeDivisionInt(d_3091_v0_, len(d_3127_v34_))
        nw623_[int(1)] = len(d_3128_v35_)
        nw623_[int(2)] = -201
        nw623_[int(3)] = d_3091_v0_
        nw623_[int(4)] = d_3091_v0_
        nw623_[int(5)] = default__.fm0(p0, globalState)
        nw623_[int(6)] = d_3091_v0_
        nw623_[int(7)] = (d_3091_v0_ if False else d_3091_v0_)
        nw623_[int(8)] = d_3091_v0_
        nw623_[int(9)] = default__.safeDivisionInt(908, d_3091_v0_)
        nw623_[int(10)] = d_3091_v0_
        nw623_[int(11)] = (-981) * (d_3091_v0_)
        nw623_[int(12)] = len((d_3129_v36_).set(d_3091_v0_, d_3091_v0_))
        nw623_[int(13)] = (0) - ((d_3091_v0_) * (len(d_3130_v37_)))
        nw623_[int(14)] = (d_3132_v39_).cf31
        nw623_[int(15)] = d_3091_v0_
        d_3133_v40_ = nw623_
        index558_ = default__.safeIndex(610, (d_3133_v40_).length(0))
        (d_3133_v40_)[index558_] = (d_3091_v0_) + (d_3091_v0_)
        index559_ = default__.safeIndex(610, (d_3133_v40_).length(0))
        (d_3133_v40_)[index559_] = default__.safeModuloInt((0) - (d_3091_v0_), d_3091_v0_)
        d_3134_i1_: int
        d_3134_i1_ = 0
        with _dafny.label("19"):
            while p0:
                with _dafny.c_label("19"):
                    if (d_3134_i1_) >= (100):
                        raise _dafny.Break("19")
                    d_3134_i1_ = (d_3134_i1_) + (1)
                    d_3135_v41_: _dafny.Seq
                    d_3135_v41_ = _dafny.SeqWithoutIsStrInference([(d_3124_v31_).f21])
                    d_3136_v42_: D16
                    d_3136_v42_ = D16_DC33((d_3124_v31_).f21, (d_3124_v31_).f21, 465, (d_3124_v31_).f21)
                    d_3137_v43_: _dafny.Map
                    d_3137_v43_ = _dafny.Map({(d_3135_v41_) < (_dafny.SeqWithoutIsStrInference([p0, p0, p0])): d_3136_v42_})
                    d_3137_v43_ = (d_3137_v43_).set((d_3091_v0_) < ((0) - ((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))])), D16_DC33(p0, (d_3124_v31_).f21, (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], (d_3124_v31_).f21))
                    d_3138_v44_: _dafny.Map
                    d_3138_v44_ = _dafny.Map({(d_3124_v31_).f21: len(d_3128_v35_)})
                    d_3139_v45_: str
                    d_3139_v45_ = _dafny.CodePoint('l')
                    d_3140_v46_: D15
                    d_3140_v46_ = D15_DC30(d_3139_v45_)
                    d_3141_v47_: _dafny.Map
                    d_3141_v47_ = _dafny.Map({default__.safeModuloInt((0) - (len(d_3138_v44_)), (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]): d_3140_v46_})
                    d_3141_v47_ = (d_3141_v47_).set(d_3091_v0_, D15_DC30(d_3139_v45_))
                    index560_ = default__.safeIndex(610, (d_3133_v40_).length(0))
                    rhs543_ = not((False if (d_3124_v31_).f21 else (d_3124_v31_).f21))
                    rhs544_ = (d_3124_v31_).f21
                    rhs545_ = (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]
                    rhs546_ = not(((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]) < (d_3091_v0_))
                    rhs547_ = default__.safeModuloInt(((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]) * ((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]), (d_3130_v37_)[default__.safeIndex(d_3091_v0_, len(d_3130_v37_))])
                    lhs427_ = d_3133_v40_
                    lhs428_ = default__.safeIndex(610, (d_3133_v40_).length(0))
                    r1 = rhs543_
                    r2 = rhs544_
                    r0 = rhs545_
                    r2 = rhs546_
                    lhs427_[lhs428_] = rhs547_
                    if (d_3135_v41_)[default__.safeIndex(d_3091_v0_, len(d_3135_v41_))]:
                        d_3133_v40_ = d_3133_v40_
                        d_3142_v48_: D3
                        d_3142_v48_ = D3_DC3((d_3135_v41_)[default__.safeIndex((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], len(d_3135_v41_))])
                        d_3143_v49_: C2
                        nw624_ = C2()
                        nw624_.ctor__(d_3142_v48_, (d_3124_v31_).f21, (d_3124_v31_).f21)
                        d_3143_v49_ = nw624_
                        rhs548_ = d_3139_v45_
                        rhs549_ = d_3143_v49_
                        d_3139_v45_ = rhs548_
                        d_3143_v49_ = rhs549_
                        d_3144_v50_: _dafny.Map
                        d_3144_v50_ = _dafny.Map({(d_3124_v31_).f21: d_3139_v45_})
                        d_3145_v51_: _dafny.Map
                        d_3145_v51_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ec"))) == (d_3128_v35_): ((d_3144_v50_)[(d_3124_v31_).f21] if ((d_3124_v31_).f21) in (d_3144_v50_) else (d_3140_v46_).cf44)})
                        d_3144_v50_ = (d_3144_v50_)
                        r1 = (d_3135_v41_)[default__.safeIndex((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], len(d_3135_v41_))]
                        (globalState).f7 = (d_3091_v0_ if (d_3124_v31_).f21 else d_3091_v0_)
                    elif True:
                        d_3091_v0_ = (-387) + ((d_3091_v0_) + ((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]))
                        d_3146_v52_: C3
                        nw625_ = C3()
                        nw625_.ctor__(default__.safeModuloInt(d_3091_v0_, len(d_3128_v35_)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtmbyeffl"))) + (d_3128_v35_), (d_3124_v31_).f21, p0)
                        d_3146_v52_ = nw625_
                        d_3147_v53_: _dafny.Map
                        d_3147_v53_ = _dafny.Map({(d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]: (d_3124_v31_).f21})
                        index561_ = default__.safeIndex(610, (d_3133_v40_).length(0))
                        (d_3133_v40_)[index561_] = (len(d_3147_v53_)) - (default__.safeDivisionInt(d_3091_v0_, (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]))
                        d_3148_v54_: _dafny.Array
                        def lambda188_(d_3149_v31_):
                            def lambda189_(d_3150_i2_):
                                return (d_3149_v31_).f21

                            return lambda189_

                        init110_ = lambda188_(d_3124_v31_)
                        nw626_ = _dafny.Array(None, 13)
                        for i0_110_ in range(nw626_.length(0)):
                            nw626_[i0_110_] = init110_(i0_110_)
                        d_3148_v54_ = nw626_
                        index562_ = default__.safeIndex(982, (d_3148_v54_).length(0))
                        (d_3148_v54_)[index562_] = (d_3124_v31_).f21
                        d_3151_v55_: D12
                        d_3151_v55_ = D12_DC22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roekfqycj")))
                        index563_ = default__.safeIndex(982, (d_3148_v54_).length(0))
                        index564_ = default__.safeIndex(610, (d_3133_v40_).length(0))
                        rhs550_ = (d_3124_v31_).f21
                        rhs551_ = (0) - ((default__.safeDivisionInt((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], len((d_3151_v55_).cf29))) * ((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]))
                        rhs552_ = (d_3146_v52_).f22
                        lhs429_ = d_3148_v54_
                        lhs430_ = default__.safeIndex(982, (d_3148_v54_).length(0))
                        lhs431_ = d_3133_v40_
                        lhs432_ = default__.safeIndex(610, (d_3133_v40_).length(0))
                        lhs429_[lhs430_] = rhs550_
                        d_3091_v0_ = rhs551_
                        lhs431_[lhs432_] = rhs552_
                        d_3152_v56_: _dafny.Map
                        d_3152_v56_ = _dafny.Map({d_3128_v35_: (d_3124_v31_).f21})
                        d_3152_v56_ = (d_3152_v56_).set(((d_3146_v52_).f23) + (d_3128_v35_), (d_3124_v31_).f21)
                    pass
            pass
        d_3153_v57_: D21
        d_3153_v57_ = D21_DC44((0) - ((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]), (d_3124_v31_).f21)
        d_3154_v58_: _dafny.MultiSet
        d_3154_v58_ = _dafny.MultiSet([(d_3153_v57_).cf69])
        d_3155_v59_: str
        d_3155_v59_ = _dafny.CodePoint('f')
        d_3156_v60_: _dafny.Set
        d_3156_v60_ = _dafny.Set({(d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], d_3091_v0_})
        d_3157_v61_: _dafny.Array
        nw627_ = _dafny.Array(None, 26)
        nw627_[int(0)] = d_3130_v37_
        nw627_[int(1)] = (d_3130_v37_) + (d_3130_v37_)
        nw627_[int(2)] = d_3130_v37_
        nw627_[int(3)] = (d_3130_v37_) + (d_3130_v37_)
        nw627_[int(4)] = (d_3130_v37_) + (d_3130_v37_)
        nw627_[int(5)] = _dafny.SeqWithoutIsStrInference([d_3091_v0_, (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]])
        nw627_[int(6)] = d_3130_v37_
        nw627_[int(7)] = d_3130_v37_
        nw627_[int(8)] = d_3130_v37_
        nw627_[int(9)] = d_3130_v37_
        nw627_[int(10)] = (d_3130_v37_) + (d_3130_v37_)
        nw627_[int(11)] = d_3130_v37_
        nw627_[int(12)] = (d_3130_v37_).set(default__.safeIndex((d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], len(d_3130_v37_)), (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))])
        nw627_[int(13)] = (d_3130_v37_) + (d_3130_v37_)
        nw627_[int(14)] = default__.fm35((d_3154_v58_).cardinality, not(p0), (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))], d_3155_v59_, globalState)
        nw627_[int(15)] = d_3130_v37_
        nw627_[int(16)] = d_3130_v37_
        nw627_[int(17)] = d_3130_v37_
        nw627_[int(18)] = d_3130_v37_
        nw627_[int(19)] = d_3130_v37_
        nw627_[int(20)] = d_3130_v37_
        nw627_[int(21)] = _dafny.SeqWithoutIsStrInference([(d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))] for d_3158_i3_ in range(default__.abs(440))])
        nw627_[int(22)] = (_dafny.SeqWithoutIsStrInference([len(d_3156_v60_), 925])) + (d_3130_v37_)
        nw627_[int(23)] = (d_3130_v37_) + (d_3130_v37_)
        nw627_[int(24)] = d_3130_v37_
        nw627_[int(25)] = _dafny.SeqWithoutIsStrInference([d_3091_v0_ for d_3159_i4_ in range(default__.abs(-444))])
        d_3157_v61_ = nw627_
        index565_ = default__.safeIndex(70, (d_3157_v61_).length(0))
        (d_3157_v61_)[index565_] = d_3130_v37_
        index566_ = default__.safeIndex(70, (d_3157_v61_).length(0))
        (d_3157_v61_)[index566_] = d_3130_v37_
        r1 = (d_3124_v31_).f21
        r0 = (d_3133_v40_)[default__.safeIndex(610, (d_3133_v40_).length(0))]
        r1 = not ((d_3124_v31_).f21) or (p0)
        r2 = not(default__.fm21(globalState))
        return r0, r1, r2

