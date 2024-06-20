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
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([357 for d_0_i0_ in range(default__.abs(253))])): True}), (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-341])): True})) | (_dafny.Map({len(_dafny.Set({_dafny.CodePoint('q')})): False}))])

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        if True:
            return (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gggckn")))) - (-433))
        elif True:
            return 356
        elif True:
            return len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): -565}))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvqorjhh")))) < (-411)

    @staticmethod
    def fm7(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(692, 418):
                d_1_v0_: int = compr_0_
                if ((692) <= (d_1_v0_)) and ((d_1_v0_) < (418)):
                    coll0_[default__.safeDivisionInt(d_1_v0_, len(_dafny.SeqWithoutIsStrInference([not(not(True))])))] = False
            return _dafny.Map(coll0_)
        return ((_dafny.Map({True: 76})) | (_dafny.Map({False: -449}))) | ((_dafny.Map({False: len(iife0_()
        )})) | (_dafny.Map({True: 713})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return D2_DC6(_dafny.Map({False: 937}))

    @staticmethod
    def fm9(p0, globalState):
        return _dafny.Set({(True) == (True), False, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qylrshxjb"))])) < (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp")) for d_2_i0_ in range(default__.abs(340))])), not (True) or (True), False})

    @staticmethod
    def fm12(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsqi"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3_i0_ in range(default__.abs(946))]))

    @staticmethod
    def fm13(p0, p1, globalState):
        if (_dafny.CodePoint('m')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ip"))
        elif True:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lucm")))

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_4_i0_ in range(default__.abs(477))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_5_i1_ in range(default__.abs(866))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_6_i2_ in range(default__.abs(731))])))

    @staticmethod
    def fm23(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(861, 243):
                d_7_v0_: int = compr_1_
                if ((861) <= (d_7_v0_)) and ((d_7_v0_) < (243)):
                    coll1_[(d_7_v0_) + (959)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_8_i0_ in range(default__.abs(867))]))
            return _dafny.Map(coll1_)
        return not((228) >= ((0) - ((D29_DC78(236, len(iife1_()
), _dafny.Map({True: 926}))).cf135)))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        if (892) != (-451):
            return D1_DC4(False)
        elif True:
            return D1_DC4(True)

    @staticmethod
    def fm25(p0, globalState):
        return _dafny.MultiSet([(-381) * (-192)])

    @staticmethod
    def fm26(p0, p1, globalState):
        return D2_DC8(not((D1_DC4(False)).cf12), not((779) >= ((D25_DC67(not(True), 81, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa")))]), True)).cf117)), not(False))

    @staticmethod
    def fm27(globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: str
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('e')])).Elements:
                d_9_v0_: str = compr_2_
                if (d_9_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('e')])):
                    coll2_ = coll2_.union(_dafny.Set([d_9_v0_]))
            return _dafny.Set(coll2_)
        return ((_dafny.Map({not(False): _dafny.Set({_dafny.CodePoint('p'), _dafny.CodePoint('n'), _dafny.CodePoint('b')})})) | (_dafny.Map({False: _dafny.Set({_dafny.CodePoint('p')})}))) | ((_dafny.Map({True: iife2_()
        })) | (_dafny.Map({False: _dafny.Set({_dafny.CodePoint('j'), _dafny.CodePoint('c')})})))

    @staticmethod
    def fm28(p0, globalState):
        return D1_DC3(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_10_i0_ in range(default__.abs(-560))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvskreag")))), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(0) - (-92), -899, len(_dafny.Set({83})), len(_dafny.Map({-512: _dafny.CodePoint('w')})), 377])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nf")))), default__.safeDivisionInt(1, -868), 636, (0) - (((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])))) - (-203)))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return ((_dafny.Map({True: len(_dafny.Map({len(_dafny.Set({_dafny.Map({False: 308}), _dafny.Map({False: -361})})): True}))})) | (_dafny.Map({True: 312}))) | ((_dafny.Map({True: (0) - ((D1_DC3(112, len(_dafny.Map({not(True): True})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wasaeyknb"))): -405})), len(_dafny.Map({_dafny.CodePoint('r'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_11_i0_ in range(default__.abs(39))])})), -633)).cf8)})) | (_dafny.Map({False: 473})))

    @staticmethod
    def fm30(p0, p1, globalState):
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: int
                for compr_5_ in (_dafny.SeqWithoutIsStrInference([451, len(_dafny.Map({True: 829}))])).Elements:
                    d_14_v1_: int = compr_5_
                    if (d_14_v1_) in (_dafny.SeqWithoutIsStrInference([451, len(_dafny.Map({True: 829}))])):
                        coll5_ = coll5_.union(_dafny.Set([(d_14_v1_) + (len(_dafny.Set({False, not(False)})))]))
                return _dafny.Set(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: int
                for compr_4_ in (_dafny.SeqWithoutIsStrInference([451, len(_dafny.Map({True: 829}))])).Elements:
                    d_12_v1_: int = compr_4_
                    if (d_12_v1_) in (_dafny.SeqWithoutIsStrInference([451, len(_dafny.Map({True: 829}))])):
                        coll4_ = coll4_.union(_dafny.Set([(d_12_v1_) + (len(_dafny.Set({False, not(False)})))]))
                return _dafny.Set(coll4_)
            compr_3_: int
            for compr_3_ in ((_dafny.SeqWithoutIsStrInference([76])) + (_dafny.SeqWithoutIsStrInference([len(iife4_()
            )]))).Elements:
                d_13_v0_: int = compr_3_
                if (d_13_v0_) in ((_dafny.SeqWithoutIsStrInference([76])) + (_dafny.SeqWithoutIsStrInference([len(iife5_()
                )]))):
                    coll3_[(d_13_v0_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rg")))]))))] = (530) + (len(_dafny.SeqWithoutIsStrInference([459, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_15_i0_ in range(default__.abs(-871))])), (0) - (len(_dafny.Map({True: False}))), (0) - (len(_dafny.SeqWithoutIsStrInference([False]))), 63]))])))
            return _dafny.Map(coll3_)
        return iife3_()
        

    @staticmethod
    def fm31(p0, p1, p2, globalState):
        def iife6_():
            def iife7_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcf"))) for d_17_i0_ in range(default__.abs(-833))])).Elements:
                    d_18_v1_: int = compr_7_
                    if (d_18_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcf"))) for d_17_i0_ in range(default__.abs(-833))])):
                        coll7_ = coll7_.union(_dafny.Set([(d_18_v1_) * (-187)]))
                return _dafny.Set(coll7_)
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(471, 379):
                d_16_v0_: int = compr_6_
                if ((471) <= (d_16_v0_)) and ((d_16_v0_) < (379)):
                    coll6_[default__.safeDivisionInt(d_16_v0_, len(iife7_()
                    ))] = False
            return _dafny.Map(coll6_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(-609, 922):
                d_20_v2_: int = compr_8_
                if ((-609) <= (d_20_v2_)) and ((d_20_v2_) < (922)):
                    coll8_[default__.safeDivisionInt(d_20_v2_, len(_dafny.Map({True: -510})))] = True
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(438, 597):
                d_21_v3_: int = compr_9_
                if ((438) <= (d_21_v3_)) and ((d_21_v3_) < (597)):
                    coll9_[(d_21_v3_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_22_i2_ in range(default__.abs(963))])))] = True
            return _dafny.Map(coll9_)
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D3_DC13(len(iife6_()
), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fahmrgw")), len(_dafny.Set({354, 29})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnvpfxa"))), _dafny.CodePoint('s'))])) + (_dafny.SeqWithoutIsStrInference([D3_DC13(364, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_19_i1_ in range(default__.abs(-922))]), -517, len(iife8_()
), _dafny.CodePoint('q')), D3_DC13(len(_dafny.Map({True: True})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tosl")), len(_dafny.SeqWithoutIsStrInference([False])), 972, _dafny.CodePoint('q'))])))) | ((_dafny.MultiSet([D3_DC13(830, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xg")), len(_dafny.Set({True, True})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({iife9_()
}))])), _dafny.CodePoint('r'))])) | (_dafny.MultiSet([D3_DC13(len(_dafny.Map({True: len(_dafny.Set({202, 457}))})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiq")), 82, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aynd"))), _dafny.CodePoint('p')), D3_DC13(493, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uw")), 462, len(_dafny.SeqWithoutIsStrInference([True, True])), _dafny.CodePoint('f'))])))

    @staticmethod
    def fm32(globalState):
        if (False if True else False):
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True])))
        elif True:
            return (_dafny.MultiSet([False])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm33(p0, globalState):
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(911, 616):
                d_23_v0_: int = compr_10_
                if ((911) <= (d_23_v0_)) and ((d_23_v0_) < (616)):
                    coll10_ = coll10_.union(_dafny.Set([(d_23_v0_) - (-58)]))
            return _dafny.Set(coll10_)
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(iife10_()
        ): 527})), len(_dafny.SeqWithoutIsStrInference([94])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbmngnpin"))), -658, len(_dafny.SeqWithoutIsStrInference([D2_DC8(not(True), True, False)]))])) + (_dafny.SeqWithoutIsStrInference([783]))

    @staticmethod
    def fm34(p0, p1, globalState):
        return (_dafny.Map({424: D2_DC8(True, True, False)})) | (_dafny.Map({-165: D2_DC8(True, True, True)}))

    @staticmethod
    def fm35(p0, p1, globalState):
        return D0_DC1(False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mphupn"))), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('v')]), not((368) in (_dafny.SeqWithoutIsStrInference([(D16_DC40(len(_dafny.Set({False, False})), 450, 211, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, True])): False})))).cf72 for d_24_i0_ in range(default__.abs(491))]))), (0) - (default__.safeModuloInt(710, 14)))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return ((_dafny.Set({False, False, True})).intersection(_dafny.Set({False}))).intersection((_dafny.Set({False, (D10_DC25(False)).cf48, False, False}) if not(True) else _dafny.Set({False, True, not(True)})))

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_25_i1_ in range(default__.abs(494))]) for d_26_i0_ in range(default__.abs(829))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cykrhjkg")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fe"))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwmhocws")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boaslvno"))]))

    @staticmethod
    def fm40(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmsi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vh")))), -127, (0) - ((-289) + (-497))])

    @staticmethod
    def fm41(globalState):
        return D0_DC0((_dafny.SeqWithoutIsStrInference([624, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfvja")))), -280])) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([47, len(_dafny.Set({False}))])])))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return D9_DC22(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([-498, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tr")))])).cardinality), len(_dafny.SeqWithoutIsStrInference([True, False]))]))

    @staticmethod
    def fm43(p0, p1, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in (_dafny.Set({(_dafny.MultiSet([True, False, False])).cardinality, 71})).Elements:
                d_27_v0_: int = compr_11_
                if (d_27_v0_) in (_dafny.Set({(_dafny.MultiSet([True, False, False])).cardinality, 71})):
                    coll11_[default__.safeModuloInt(d_27_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgpxdfu"))))] = True
            return _dafny.Map(coll11_)
        return ((_dafny.Map({True: 715}) if True else _dafny.Map({not(not(not(not(True)))): 163}))) | ((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([423, 527, (0) - (len(_dafny.SeqWithoutIsStrInference([True, False, True, False, False])))]))})) | (_dafny.Map({True: len(_dafny.Map({len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([23]))).cardinality: len(iife11_()
        )})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmkulofbe")))}))})))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(not(False))])) + (_dafny.SeqWithoutIsStrInference([not(not(not(True))), True]))

    @staticmethod
    def fm45(globalState):
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: _dafny.Seq
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_31_i3_ in range(default__.abs(242))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkyiedaj"))])).Elements:
                d_32_v0_: _dafny.Seq = compr_12_
                if (d_32_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_31_i3_ in range(default__.abs(242))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkyiedaj"))])):
                    coll12_ = coll12_.union(_dafny.Set([d_32_v0_]))
            return _dafny.Set(coll12_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frqcl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njtvu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lb")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_28_i0_ in range(default__.abs(13))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_29_i1_ in range(default__.abs(588))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_30_i2_ in range(default__.abs(933))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anl"))}))) | ((D30_DC79(iife12_()
)).cf137)

    @staticmethod
    def fm46(globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        return (_dafny.Map({False: False})) | ((_dafny.Map({True: False})) | (_dafny.Map({True: False})))

    @staticmethod
    def fm48(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_33_i0_ in range(default__.abs(-505))])

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        return _dafny.Set({(len(_dafny.Map({len(_dafny.Set({_dafny.Set({(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opxh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgmmnul"))])).cardinality, 201, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), 666})})): 26}))) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "albnktp"))))})

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        return _dafny.Map({_dafny.Map({857: True}): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yog"))) not in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmbiql")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_34_i0_ in range(default__.abs(-948))])]))})

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(594, 461):
                d_35_v0_: int = compr_13_
                if ((594) <= (d_35_v0_)) and ((d_35_v0_) < (461)):
                    coll13_[default__.safeDivisionInt(d_35_v0_, 331)] = -447
            return _dafny.Map(coll13_)
        return _dafny.SeqWithoutIsStrInference([(iife13_()
        ) | (_dafny.Map({350: len(_dafny.Map({-505: False}))}))])

    @staticmethod
    def fm52(p0, p1, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsxh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yagibnn")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_36_i0_ in range(default__.abs(947))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_37_i1_ in range(default__.abs(671))])]))).intersection((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "seqmtp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fi")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_38_i2_ in range(default__.abs(335))])])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aahqhps")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ie"))]))))

    @staticmethod
    def fm53(globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(-241, 903):
                d_43_v0_: int = compr_14_
                if ((-241) <= (d_43_v0_)) and ((d_43_v0_) < (903)):
                    def iife15_():
                        coll15_ = _dafny.Map()
                        compr_15_: str
                        for compr_15_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])).Elements:
                            d_44_v1_: str = compr_15_
                            if (d_44_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])):
                                coll15_[d_44_v1_] = -854
                        return _dafny.Map(coll15_)
                    coll14_[(d_43_v0_) * ((D30_DC80(False, _dafny.Set({(_dafny.MultiSet([True])).cardinality}), True, -938)).cf141)] = len(_dafny.SeqWithoutIsStrInference([len(iife15_()
                    )]))
            return _dafny.Map(coll14_)
        return ((_dafny.SeqWithoutIsStrInference([D0_DC1(False, 952, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_39_i0_ in range(default__.abs(572))]), True, len(_dafny.SeqWithoutIsStrInference([430]))), D0_DC1(False, 482, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b')]), False, 280)])) + (_dafny.SeqWithoutIsStrInference([D0_DC1(False, 943, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n'), _dafny.CodePoint('a'), _dafny.CodePoint('n'), _dafny.CodePoint('f'), _dafny.CodePoint('w')]), True, 27) for d_40_i1_ in range(default__.abs(946))]))) + ((_dafny.SeqWithoutIsStrInference([D0_DC1(True, 822, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]), True, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_41_i2_ in range(default__.abs(983))]))), D0_DC1(True, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({-896: True})): True}))])), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_42_i3_ in range(default__.abs(400))]), True, len(iife14_()
))])) + (_dafny.SeqWithoutIsStrInference([D0_DC1(not(not(not(False))), 430, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s'), _dafny.CodePoint('r')]), not(True), 36), D0_DC1(not(not(not(False))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_45_i4_ in range(default__.abs(-918))]))])), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('m')]), False, 435), D0_DC1(False, 167, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')]), False, -988), D0_DC1(True, 334, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')]), True, len(_dafny.SeqWithoutIsStrInference([True, True])))])))

    @staticmethod
    def fm54(p0, p1, globalState):
        def iife16_():
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: str
                for compr_18_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('a')])).Elements:
                    d_46_v1_: str = compr_18_
                    if (d_46_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('a')])):
                        coll18_[d_46_v1_] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, True])): True}))
                return _dafny.Map(coll18_)
            coll16_ = _dafny.Map()
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: str
                for compr_17_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('a')])).Elements:
                    d_46_v1_: str = compr_17_
                    if (d_46_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('a')])):
                        coll17_[d_46_v1_] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, True])): True}))
                return _dafny.Map(coll17_)
            compr_16_: _dafny.Map
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([iife17_()
            , _dafny.Map({_dafny.CodePoint('q'): 499}), _dafny.Map({_dafny.CodePoint('j'): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_47_i0_ in range(default__.abs(771))]))})])).Elements:
                d_48_v0_: _dafny.Map = compr_16_
                if (d_48_v0_) in (_dafny.SeqWithoutIsStrInference([iife18_()
                , _dafny.Map({_dafny.CodePoint('q'): 499}), _dafny.Map({_dafny.CodePoint('j'): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_47_i0_ in range(default__.abs(771))]))})])):
                    coll16_[d_48_v0_] = (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))])).cardinality for d_49_i1_ in range(default__.abs(-647))]))) + ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([322, len(_dafny.SeqWithoutIsStrInference([False]))]))).cardinality)
            return _dafny.Map(coll16_)
        return iife16_()
        

    @staticmethod
    def fm55(p0, p1, p2, p3, globalState):
        if not((_dafny.MultiSet([890])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptwyd"))), (0) - (-377), len(_dafny.Map({False: not(False)})), 263, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnmsblqx"))))}))])))):
            return D9_DC23(False, not(False), (0) - (len(_dafny.Map({not(False): (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_50_i0_ in range(default__.abs(946))])))}))))
        elif True:
            return D9_DC23(True, True, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: True})])))

    @staticmethod
    def fm56(p0, p1, p2, p3, globalState):
        source0_ = D17_DC42(True)
        if source0_.is_DC42:
            d_51___mcc_h0_ = source0_.cf74
            d_52_cf74_ = d_51___mcc_h0_
            return D13_DC33(d_52_cf74_, len(_dafny.SeqWithoutIsStrInference([d_52_cf74_])), d_52_cf74_, 117)
        elif source0_.is_DC43:
            d_53___mcc_h1_ = source0_.cf75
            d_54___mcc_h2_ = source0_.cf76
            d_55_cf76_ = d_54___mcc_h2_
            d_56_cf75_ = d_53___mcc_h1_
            return D13_DC33(False, 399, False, 520)
        elif source0_.is_DC44:
            d_57___mcc_h3_ = source0_.cf77
            d_58___mcc_h4_ = source0_.cf78
            d_59_cf78_ = d_58___mcc_h4_
            d_60_cf77_ = d_57___mcc_h3_
            return D13_DC33(False, len(_dafny.Map({False: False})), not(True), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fn"))))
        elif True:
            d_61___mcc_h5_ = source0_.cf73
            d_62_cf73_ = d_61___mcc_h5_
            if not(False):
                return D13_DC33(False, len(_dafny.SeqWithoutIsStrInference([not(False)])), False, 203)
            elif True:
                return D13_DC33(False, -622, True, (0) - (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_63_i0_ in range(default__.abs(313))]))): _dafny.MultiSet([(0) - ((0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))}))))])}))))

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        source1_ = D3_DC13((_dafny.MultiSet([489])).cardinality, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybukhe")), len(_dafny.Set({True, (D3_DC11(not(True), True)).cf24})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujk"))): -931})), (D8_DC21(_dafny.CodePoint('n'), -482)).cf41)
        if source1_.is_DC11:
            d_64___mcc_h0_ = source1_.cf23
            d_65___mcc_h1_ = source1_.cf24
            d_66_cf24_ = d_65___mcc_h1_
            d_67_cf23_ = d_64___mcc_h0_
            return (_dafny.Map({len(_dafny.Map({391: _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_66_cf24_: 903}))])})): d_66_cf24_})) | (_dafny.Map({424: d_67_cf23_}))
        elif source1_.is_DC12:
            d_68___mcc_h2_ = source1_.cf25
            d_69___mcc_h3_ = source1_.cf26
            d_70___mcc_h4_ = source1_.cf27
            d_71___mcc_h5_ = source1_.cf28
            d_72___mcc_h6_ = source1_.cf29
            d_73_cf29_ = d_72___mcc_h6_
            d_74_cf28_ = d_71___mcc_h5_
            d_75_cf27_ = d_70___mcc_h4_
            d_76_cf26_ = d_69___mcc_h3_
            d_77_cf25_ = d_68___mcc_h2_
            return (D17_DC41(_dafny.Map({d_75_cf27_: d_76_cf26_}))).cf73
        elif source1_.is_DC13:
            d_78___mcc_h7_ = source1_.cf30
            d_79___mcc_h8_ = source1_.cf31
            d_80___mcc_h9_ = source1_.cf32
            d_81___mcc_h10_ = source1_.cf33
            d_82___mcc_h11_ = source1_.cf34
            d_83_cf34_ = d_82___mcc_h11_
            d_84_cf33_ = d_81___mcc_h10_
            d_85_cf32_ = d_80___mcc_h9_
            d_86_cf31_ = d_79___mcc_h8_
            d_87_cf30_ = d_78___mcc_h7_
            if False:
                return _dafny.Map({d_84_cf33_: False})
            elif True:
                return _dafny.Map({(0) - (d_84_cf33_): False})
        elif source1_.is_DC10:
            d_88___mcc_h12_ = source1_.cf22
            d_89_cf22_ = d_88___mcc_h12_
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))): not(not(False))})
        elif True:
            d_90___mcc_h13_ = source1_.cf35
            d_91_cf35_ = d_90___mcc_h13_
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvhocxjgf"))): True})), 502, -996, (0) - (-394), len(_dafny.Map({False: False}))])): not(False)})

    @staticmethod
    def fm58(p0, p1, globalState):
        return _dafny.Map({D19_DC50(False): True})

    @staticmethod
    def fm59(globalState):
        return D3_DC13((len(_dafny.SeqWithoutIsStrInference([False, False]))) - (len(_dafny.SeqWithoutIsStrInference([False, False]))), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_92_i0_ in range(default__.abs(149))]), (654 if False else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wblaoerrr")))), default__.safeModuloInt(77, (0) - (len(_dafny.SeqWithoutIsStrInference([len((D25_DC67(True, 399, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})), -158, len(_dafny.Map({True: not(False)}))]), True)).cf118)])))), _dafny.CodePoint('u'))

    @staticmethod
    def fm60(p0, p1, p2, globalState):
        return _dafny.Map({(D29_DC78(165, len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcb"))})), _dafny.Map({True: 931}))).cf135: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ww"))): -340})})

    @staticmethod
    def fm61(p0, globalState):
        source2_ = D20_DC53(True, (D19_DC50(True)).cf84, _dafny.CodePoint('l'))
        if source2_.is_DC53:
            d_93___mcc_h0_ = source2_.cf87
            d_94___mcc_h1_ = source2_.cf88
            d_95___mcc_h2_ = source2_.cf89
            d_96_cf89_ = d_95___mcc_h2_
            d_97_cf88_ = d_94___mcc_h1_
            d_98_cf87_ = d_93___mcc_h0_
            if (D19_DC49(False)).cf83:
                return D21_DC55(_dafny.Set({_dafny.Map({len(_dafny.Map({d_98_cf87_: d_96_cf89_})): len(_dafny.Map({d_97_cf88_: -677}))}), _dafny.Map({213: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoqthk")))]))})}))
            elif True:
                return D21_DC55(_dafny.Set({_dafny.Map({-230: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({(0) - ((_dafny.MultiSet([d_98_cf87_, d_98_cf87_])).cardinality): d_98_cf87_}): (D13_DC33(d_97_cf88_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), d_98_cf87_, -347)).cf56}))]))})}))
        elif source2_.is_DC52:
            d_99___mcc_h3_ = source2_.cf86
            d_100_cf86_ = d_99___mcc_h3_
            return D21_DC55(_dafny.Set({d_100_cf86_, d_100_cf86_, d_100_cf86_, d_100_cf86_}))
        elif True:
            d_101___mcc_h4_ = source2_.cf90
            d_102_cf90_ = d_101___mcc_h4_
            return D21_DC55(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_103_i0_ in range(default__.abs(-390))])): len(_dafny.Set({True}))})}))

    @staticmethod
    def fm62(p0, p1, globalState):
        return (D31_DC83(_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))).cf143

    @staticmethod
    def fm63(globalState):
        return _dafny.Map({((0) - (len(_dafny.Set({False, False})))) != (len(_dafny.Set({not(True), True, True}))): _dafny.CodePoint('c')})

    @staticmethod
    def fm64(p0, p1, p2, p3, globalState):
        if (False) and (not(True)):
            return (_dafny.Map({_dafny.CodePoint('q'): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([391]))])), 622, (_dafny.MultiSet([672, len(_dafny.Map({(D31_DC84(False, False)).cf144: _dafny.CodePoint('m')})), 718, len(_dafny.SeqWithoutIsStrInference([not(False), True]))])).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsuf")))), (0) - (-435)]))})) | (_dafny.Map({_dafny.CodePoint('h'): len(_dafny.Map({_dafny.Map({True: 224}): len(_dafny.Map({False: -9}))}))}))
        elif True:
            return _dafny.Map({_dafny.CodePoint('h'): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({250}))]))})

    @staticmethod
    def fm65(globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: _dafny.Seq
            for compr_19_ in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hw")) for d_104_i0_ in range(default__.abs(-409))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_105_i1_ in range(default__.abs(-614))])]))).Elements:
                d_106_v0_: _dafny.Seq = compr_19_
                if (d_106_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hw")) for d_104_i0_ in range(default__.abs(-409))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_105_i1_ in range(default__.abs(-614))])]))):
                    coll19_[d_106_v0_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isiwmcovd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkdiy"))))
            return _dafny.Map(coll19_)
        return iife19_()
        

    @staticmethod
    def m0(globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: _dafny.Map = _dafny.Map({})
        d_107_v0_: bool
        d_107_v0_ = False
        d_108_v1_: _dafny.Array
        nw0_ = _dafny.Array(None, 12)
        nw0_[int(0)] = d_107_v0_
        nw0_[int(1)] = d_107_v0_
        nw0_[int(2)] = d_107_v0_
        nw0_[int(3)] = d_107_v0_
        nw0_[int(4)] = d_107_v0_
        nw0_[int(5)] = d_107_v0_
        nw0_[int(6)] = d_107_v0_
        nw0_[int(7)] = d_107_v0_
        nw0_[int(8)] = False
        nw0_[int(9)] = d_107_v0_
        nw0_[int(10)] = d_107_v0_
        nw0_[int(11)] = d_107_v0_
        d_108_v1_ = nw0_
        d_109_v2_: _dafny.Seq
        d_109_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvivu"))
        d_110_v3_: D13
        d_110_v3_ = D13_DC34(d_108_v1_, d_109_v2_, d_107_v0_, d_107_v0_)
        if (d_110_v3_).cf61:
            d_111_v4_: _dafny.Map
            d_111_v4_ = _dafny.Map({d_108_v1_: d_107_v0_})
            d_112_v5_: D13
            d_112_v5_ = D13_DC32(d_111_v4_)
            d_113_v6_: int
            d_113_v6_ = -108
            d_114_v7_: _dafny.Map
            d_114_v7_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jm")): (0) - (d_113_v6_)})
            d_115_v8_: C5
            nw1_ = C5()
            nw1_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_116_i0_ in range(default__.abs(980))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lf")), (default__.fm65(globalState)) != (d_114_v7_))
            d_115_v8_ = nw1_
            d_117_v9_: _dafny.MultiSet
            d_117_v9_ = _dafny.MultiSet([(d_113_v6_) * (d_113_v6_)])
            d_118_v10_: _dafny.Array
            nw2_ = _dafny.Array(int(0), 15)
            d_118_v10_ = nw2_
            d_119_v11_: _dafny.Seq
            d_119_v11_ = _dafny.SeqWithoutIsStrInference([d_108_v1_])
            d_120_v12_: _dafny.MultiSet
            d_120_v12_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_108_v1_]), d_119_v11_, (d_119_v11_).set(default__.safeIndex(default__.fm1(default__.fm1(d_113_v6_, d_107_v0_, d_107_v0_, globalState), d_107_v0_, d_107_v0_, globalState), len(d_119_v11_)), d_108_v1_)])
            rhs0_ = d_112_v5_
            rhs1_ = d_115_v8_
            rhs2_ = d_118_v10_
            rhs3_ = ((d_120_v12_)[(d_119_v11_) + (d_119_v11_)] if ((d_119_v11_) + (d_119_v11_)) in (d_120_v12_) else (d_113_v6_) + (d_113_v6_))
            rhs4_ = d_117_v9_
            d_112_v5_ = rhs0_
            d_115_v8_ = rhs1_
            r2 = rhs2_
            d_113_v6_ = rhs3_
            d_117_v9_ = rhs4_
            d_121_v13_: str
            d_121_v13_ = _dafny.CodePoint('n')
            d_121_v13_ = default__.fm46(globalState)
            d_107_v0_ = ((d_113_v6_) - (d_113_v6_)) == (d_113_v6_)
            d_122_v14_: _dafny.Seq
            d_122_v14_ = _dafny.SeqWithoutIsStrInference([d_107_v0_])
            d_123_v15_: _dafny.Map
            d_123_v15_ = _dafny.Map({d_113_v6_: len(_dafny.SeqWithoutIsStrInference([d_121_v13_ for d_124_i1_ in range(default__.abs(100))]))})
            d_125_v16_: _dafny.Seq
            d_125_v16_ = _dafny.SeqWithoutIsStrInference([d_123_v15_, _dafny.Map({d_113_v6_: d_113_v6_})])
            d_126_v17_: T1
            nw3_ = C6()
            nw3_.ctor__(d_122_v14_, d_113_v6_, not(True), (d_115_v8_).f28, d_125_v16_)
            d_126_v17_ = nw3_
            d_127_v18_: _dafny.Map
            d_127_v18_ = _dafny.Map({d_126_v17_: d_121_v13_})
            rhs5_ = ((d_127_v18_)[d_126_v17_] if (d_126_v17_) in (d_127_v18_) else d_121_v13_)
            rhs6_ = (d_126_v17_).f19
            lhs0_ = globalState
            d_121_v13_ = rhs5_
            lhs0_.f9 = rhs6_
            d_128_v19_: _dafny.Array
            nw4_ = _dafny.Array(None, 3)
            d_128_v19_ = nw4_
            nw5_ = _dafny.Array(None, 4)
            d_128_v19_ = nw5_
        elif True:
            (globalState).f12 = (d_109_v2_) + (d_109_v2_)
            d_129_v20_: int
            d_129_v20_ = 825
            d_130_v21_: C0
            nw6_ = C0()
            nw6_.ctor__(len(_dafny.Map({default__.fm1(d_129_v20_, d_107_v0_, d_107_v0_, globalState): d_107_v0_})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_131_i2_ in range(default__.abs(692))])))
            d_130_v21_ = nw6_
            d_132_v22_: _dafny.MultiSet
            d_132_v22_ = _dafny.MultiSet([d_130_v21_, d_130_v21_])
            (globalState).f15 = ((d_132_v22_)[d_130_v21_] if (d_130_v21_) in (d_132_v22_) else 664)
            d_133_v23_: _dafny.Seq
            d_133_v23_ = _dafny.SeqWithoutIsStrInference([d_129_v20_])
            d_134_v24_: _dafny.Map
            d_134_v24_ = _dafny.Map({len(_dafny.Set({d_107_v0_})): (d_130_v21_).f32})
            d_135_v25_: _dafny.Seq
            d_135_v25_ = _dafny.SeqWithoutIsStrInference([d_134_v24_, d_134_v24_, d_134_v24_, d_134_v24_, d_134_v24_])
            d_136_v26_: T1
            nw7_ = C3()
            nw7_.ctor__(d_133_v23_, d_109_v2_, (d_135_v25_) + (d_135_v25_), (d_107_v0_ if d_107_v0_ else False))
            d_136_v26_ = nw7_
            d_136_v26_ = d_136_v26_
            (globalState).f15 = len((_dafny.SeqWithoutIsStrInference([(d_130_v21_).f32])) + (d_133_v23_))
            d_137_v27_: D21
            d_137_v27_ = D21_DC56((d_136_v26_).f18, d_129_v20_, 229)
            if (d_137_v27_).cf92:
                d_138_v28_: str
                d_138_v28_ = _dafny.CodePoint('k')
                d_138_v28_ = d_138_v28_
                index0_ = default__.safeIndex(373, (d_108_v1_).length(0))
                (d_108_v1_)[index0_] = d_107_v0_
                d_139_v29_: D16
                d_139_v29_ = D16_DC39(d_107_v0_, 306)
                index1_ = default__.safeIndex(373, (d_108_v1_).length(0))
                (d_108_v1_)[index1_] = (d_139_v29_).cf67
                d_140_v30_: C3
                nw8_ = C3()
                nw8_.ctor__(default__.fm33((d_108_v1_)[default__.safeIndex(373, (d_108_v1_).length(0))], globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xj")), (_dafny.SeqWithoutIsStrInference([d_134_v24_ for d_141_i3_ in range(default__.abs(807))])) + (d_135_v25_), (d_136_v26_).f18)
                d_140_v30_ = nw8_
                d_142_v31_: _dafny.Set
                d_142_v31_ = _dafny.Set({not((d_136_v26_).f18), (d_136_v26_).f18})
                index2_ = default__.safeIndex(373, (d_108_v1_).length(0))
                (d_108_v1_)[index2_] = (_dafny.Set({(d_108_v1_)[default__.safeIndex(373, (d_108_v1_).length(0))], d_107_v0_})).issubset(d_142_v31_)
                d_143_v32_: _dafny.MultiSet
                d_143_v32_ = _dafny.MultiSet([(d_136_v26_).f18])
                (globalState).f2 = (d_143_v32_).ispropersubset(d_143_v32_)
            elif True:
                d_107_v0_ = (d_136_v26_).f18
                d_144_v33_: D9
                d_144_v33_ = D9_DC23(d_107_v0_, (d_136_v26_).f18, (d_130_v21_).f31)
                d_145_v34_: C8
                nw9_ = C8()
                nw9_.ctor__((d_130_v21_).f32, (d_136_v26_).f18)
                d_145_v34_ = nw9_
                d_146_v35_: _dafny.MultiSet
                d_146_v35_ = _dafny.MultiSet([d_145_v34_])
                d_147_v36_: C3
                nw10_ = C3()
                nw10_.ctor__(_dafny.SeqWithoutIsStrInference([807, (d_130_v21_).f32, -416, (d_146_v35_).cardinality]), (d_136_v26_).f19, (d_136_v26_).f20, (d_136_v26_).f18)
                d_147_v36_ = nw10_
                d_148_v37_: C7
                nw11_ = C7()
                nw11_.ctor__(d_129_v20_, (d_136_v26_).f18)
                d_148_v37_ = nw11_
                d_149_v38_: _dafny.Set
                d_149_v38_ = _dafny.Set({(d_136_v26_).f18, d_107_v0_})
                rhs7_ = d_144_v33_
                rhs8_ = ((d_149_v38_).issubset(d_149_v38_)) == ((d_136_v26_).f18)
                rhs9_ = (d_130_v21_).f31
                rhs10_ = d_147_v36_
                rhs11_ = d_148_v37_
                lhs1_ = globalState
                lhs2_ = globalState
                d_144_v33_ = rhs7_
                lhs1_.f2 = rhs8_
                lhs2_.f3 = rhs9_
                d_147_v36_ = rhs10_
                d_148_v37_ = rhs11_
                index3_ = default__.safeIndex(858, (d_108_v1_).length(0))
                (d_108_v1_)[index3_] = False
                index4_ = default__.safeIndex(858, (d_108_v1_).length(0))
                (d_108_v1_)[index4_] = (d_136_v26_).f18
                (globalState).f0 = d_145_v34_.f23
                d_150_v39_: C6
                nw12_ = C6()
                nw12_.ctor__(_dafny.SeqWithoutIsStrInference([(d_136_v26_).f18]), len((d_136_v26_).f19), (d_136_v26_).f18, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_151_i4_ in range(default__.abs(-34))]), (d_136_v26_).f20)
                d_150_v39_ = nw12_
        d_152_i5_: int
        d_152_i5_ = 0
        with _dafny.label("0"):
            while not(d_107_v0_):
                with _dafny.c_label("0"):
                    if (d_152_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_152_i5_ = (d_152_i5_) + (1)
                    d_153_v40_: T0
                    nw13_ = C9()
                    nw13_.ctor__(d_108_v1_, d_107_v0_)
                    d_153_v40_ = nw13_
                    d_154_v41_: _dafny.Array
                    nw14_ = _dafny.Array(None, 7)
                    nw14_[int(0)] = d_153_v40_
                    nw14_[int(1)] = d_153_v40_
                    nw14_[int(2)] = d_153_v40_
                    nw14_[int(3)] = d_153_v40_
                    nw14_[int(4)] = d_153_v40_
                    nw14_[int(5)] = d_153_v40_
                    nw14_[int(6)] = d_153_v40_
                    d_154_v41_ = nw14_
                    index5_ = default__.safeIndex(954, (d_154_v41_).length(0))
                    (d_154_v41_)[index5_] = d_153_v40_
                    d_155_v42_: int
                    d_155_v42_ = 725
                    index6_ = default__.safeIndex(954, (d_154_v41_).length(0))
                    rhs12_ = d_155_v42_
                    rhs13_ = d_153_v40_
                    lhs3_ = d_154_v41_
                    lhs4_ = default__.safeIndex(954, (d_154_v41_).length(0))
                    r1 = rhs12_
                    lhs3_[lhs4_] = rhs13_
                    d_156_v43_: _dafny.Set
                    d_156_v43_ = _dafny.Set({(0) - ((0) - (d_155_v42_)), (0) - (d_155_v42_)})
                    d_156_v43_ = d_156_v43_
                    d_157_v44_: _dafny.Seq
                    d_157_v44_ = _dafny.SeqWithoutIsStrInference([d_107_v0_, d_107_v0_, d_107_v0_, d_107_v0_])
                    if (not(not(d_107_v0_))) or ((d_157_v44_)[default__.safeIndex(d_155_v42_, len(d_157_v44_))]):
                        d_158_v45_: C0
                        nw15_ = C0()
                        nw15_.ctor__((d_155_v42_) * (d_155_v42_), len(d_109_v2_))
                        d_158_v45_ = nw15_
                        d_159_v46_: _dafny.Map
                        d_159_v46_ = _dafny.Map({(d_158_v45_).f32: d_155_v42_})
                        d_160_v47_: _dafny.Map
                        d_160_v47_ = _dafny.Map({d_107_v0_: 815})
                        d_161_v48_: D3
                        d_161_v48_ = D3_DC13(d_155_v42_, d_109_v2_, d_155_v42_, ((d_160_v47_)[(d_153_v40_).f18] if ((d_153_v40_).f18) in (d_160_v47_) else 507), _dafny.CodePoint('c'))
                        d_162_v49_: _dafny.Seq
                        d_162_v49_ = _dafny.SeqWithoutIsStrInference([d_159_v46_, d_159_v46_, d_159_v46_])
                        d_163_v50_: T1
                        nw16_ = C6()
                        nw16_.ctor__(default__.fm44(len(d_159_v46_), (d_153_v40_).f18, (d_161_v48_).cf34, ((d_160_v47_)[(d_153_v40_).f18] if ((d_153_v40_).f18) in (d_160_v47_) else 550), globalState), ((d_158_v45_).f31) * (d_155_v42_), (d_153_v40_).f18, d_109_v2_, d_162_v49_)
                        d_163_v50_ = nw16_
                        d_163_v50_ = d_163_v50_
                        d_109_v2_ = (d_110_v3_).cf60
                        (globalState).f2 = True
                        (globalState).f2 = ((d_158_v45_).f32) != ((0) - ((d_158_v45_).f31))
                    elif True:
                        (globalState).f2 = False
                        d_107_v0_ = True
                        (globalState).f2 = d_107_v0_
                        d_107_v0_ = d_107_v0_
                        r0 = d_155_v42_
                    d_164_v51_: C1
                    nw17_ = C1()
                    nw17_.ctor__(d_107_v0_)
                    d_164_v51_ = nw17_
                    d_165_v52_: _dafny.Map
                    d_165_v52_ = _dafny.Map({d_164_v51_: (d_153_v40_).f18})
                    d_166_v53_: D11
                    d_166_v53_ = D11_DC27((d_165_v52_).set(d_164_v51_, (d_153_v40_).f18))
                    d_167_v54_: D11
                    d_167_v54_ = D11_DC29(d_166_v53_)
                    d_167_v54_ = d_167_v54_
                    pass
            pass
        if d_107_v0_:
            d_168_v55_: int
            d_168_v55_ = -816
            d_169_v56_: _dafny.Seq
            d_169_v56_ = _dafny.SeqWithoutIsStrInference([(0) - (d_168_v55_)])
            d_170_v57_: _dafny.Set
            d_170_v57_ = _dafny.Set({d_107_v0_, d_107_v0_})
            d_171_v58_: _dafny.Map
            d_171_v58_ = _dafny.Map({d_168_v55_: len(d_170_v57_)})
            d_172_v59_: _dafny.Seq
            d_172_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_168_v55_: d_168_v55_}), d_171_v58_])
            d_173_v60_: C3
            nw18_ = C3()
            nw18_.ctor__(d_169_v56_, d_109_v2_, d_172_v59_, not((d_107_v0_ if d_107_v0_ else d_107_v0_)))
            d_173_v60_ = nw18_
            d_173_v60_ = d_173_v60_
            (globalState).f2 = (d_107_v0_) or (d_107_v0_)
            d_174_v61_: str
            d_174_v61_ = _dafny.CodePoint('b')
            d_175_v62_: _dafny.Map
            d_175_v62_ = _dafny.Map({False: False})
            d_176_v63_: _dafny.MultiSet
            d_176_v63_ = _dafny.MultiSet([len(d_175_v62_)])
            d_177_v64_: _dafny.MultiSet
            d_177_v64_ = _dafny.MultiSet([d_107_v0_])
            d_178_v65_: _dafny.Seq
            d_178_v65_ = _dafny.SeqWithoutIsStrInference([d_177_v64_])
            d_179_v66_: _dafny.Map
            d_179_v66_ = _dafny.Map({(d_178_v65_)[default__.safeIndex(791, len(d_178_v65_))]: d_168_v55_})
            d_180_v67_: _dafny.Array
            nw19_ = _dafny.Array(None, 25)
            nw19_[int(0)] = d_109_v2_
            nw19_[int(1)] = d_109_v2_
            nw19_[int(2)] = ((d_109_v2_) + (d_109_v2_)).set(default__.safeIndex(((d_173_v60_).f35)[default__.safeIndex((0) - (d_168_v55_), len((d_173_v60_).f35))], len((d_109_v2_) + (d_109_v2_))), d_174_v61_)
            nw19_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgsgthyqi"))
            nw19_[int(4)] = _dafny.SeqWithoutIsStrInference([d_174_v61_ for d_181_i6_ in range(default__.abs(952))])
            nw19_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfhtk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkjwvty")))
            nw19_[int(6)] = d_109_v2_
            nw19_[int(7)] = d_109_v2_
            nw19_[int(8)] = (d_109_v2_).set(default__.safeIndex(d_168_v55_, len(d_109_v2_)), d_174_v61_)
            nw19_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwuxuka"))
            nw19_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yftjej"))
            nw19_[int(11)] = default__.fm16(d_168_v55_, True, d_174_v61_, d_168_v55_, globalState)
            nw19_[int(12)] = d_109_v2_
            nw19_[int(13)] = (d_109_v2_) + (d_109_v2_)
            nw19_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvtv"))
            nw19_[int(15)] = ((_dafny.SeqWithoutIsStrInference([d_174_v61_ for d_182_i7_ in range(default__.abs(625))]) if d_107_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnpdry")))).set(default__.safeIndex(((d_176_v63_)[(0) - (d_168_v55_)] if ((0) - (d_168_v55_)) in (d_176_v63_) else d_168_v55_), len((_dafny.SeqWithoutIsStrInference([d_174_v61_ for d_183_i7_ in range(default__.abs(625))]) if d_107_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnpdry"))))), d_174_v61_)
            nw19_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbke"))
            nw19_[int(17)] = (d_109_v2_ if default__.fm23(_dafny.Set({d_168_v55_}), len((d_179_v66_).set(d_177_v64_, len(_dafny.SeqWithoutIsStrInference([954 for d_184_i8_ in range(default__.abs(459))])))), globalState) else d_109_v2_)
            nw19_[int(18)] = (d_173_v60_).fm37(d_168_v55_, globalState)
            nw19_[int(19)] = _dafny.SeqWithoutIsStrInference([d_174_v61_ for d_185_i9_ in range(default__.abs(510))])
            nw19_[int(20)] = (d_109_v2_) + ((d_109_v2_).set(default__.safeIndex(d_168_v55_, len(d_109_v2_)), d_174_v61_))
            nw19_[int(21)] = _dafny.SeqWithoutIsStrInference([d_174_v61_ for d_186_i10_ in range(default__.abs(609))])
            nw19_[int(22)] = (d_109_v2_) + (_dafny.SeqWithoutIsStrInference([d_174_v61_ for d_187_i11_ in range(default__.abs(219))]))
            nw19_[int(23)] = d_109_v2_
            nw19_[int(24)] = d_109_v2_
            d_180_v67_ = nw19_
            d_180_v67_ = d_180_v67_
            d_188_v68_: C0
            nw20_ = C0()
            nw20_.ctor__(d_168_v55_, d_168_v55_)
            d_188_v68_ = nw20_
            rhs14_ = d_107_v0_
            rhs15_ = d_168_v55_
            d_107_v0_ = rhs14_
            d_168_v55_ = rhs15_
        elif True:
            d_189_v69_: _dafny.Set
            d_189_v69_ = _dafny.Set({d_107_v0_})
            d_190_v70_: T0
            nw21_ = C5()
            nw21_.ctor__(d_109_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_191_i12_ in range(default__.abs(853))]), (d_189_v69_).issubset(d_189_v69_))
            d_190_v70_ = nw21_
            d_190_v70_ = d_190_v70_
            d_192_v71_: C5
            nw22_ = C5()
            nw22_.ctor__(d_109_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_193_i13_ in range(default__.abs(821))]), ((d_190_v70_).f18 if (d_190_v70_).f18 else (d_190_v70_).f18))
            d_192_v71_ = nw22_
            d_192_v71_ = d_192_v71_
            d_194_v72_: int
            d_194_v72_ = 999
            d_195_v73_: _dafny.Map
            d_195_v73_ = _dafny.Map({(d_190_v70_).f18: d_194_v72_})
            (globalState).f2 = ((d_107_v0_ if (d_190_v70_).f18 else (d_190_v70_).f18)) in (d_195_v73_)
            index7_ = default__.safeIndex(930, (d_108_v1_).length(0))
            (d_108_v1_)[index7_] = (d_190_v70_).f18
            d_196_v74_: _dafny.Seq
            d_196_v74_ = _dafny.SeqWithoutIsStrInference([True])
            index8_ = default__.safeIndex(930, (d_108_v1_).length(0))
            (d_108_v1_)[index8_] = (not ((d_196_v74_)[default__.safeIndex(d_194_v72_, len(d_196_v74_))]) or (True)) or ((d_190_v70_).f18)
            (globalState).f12 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvvi"))
        d_197_v75_: int
        d_197_v75_ = 690
        d_198_v76_: D17
        d_198_v76_ = D17_DC44(d_108_v1_, d_197_v75_)
        d_198_v76_ = d_198_v76_
        r1 = d_197_v75_
        d_199_v77_: _dafny.Set
        d_199_v77_ = _dafny.Set({d_197_v75_, d_197_v75_})
        d_200_v78_: _dafny.Map
        d_200_v78_ = _dafny.Map({d_107_v0_: d_199_v77_})
        d_201_i14_: int
        d_201_i14_ = 0
        with _dafny.label("1"):
            while default__.fm23(((d_200_v78_)[d_107_v0_] if (d_107_v0_) in (d_200_v78_) else d_199_v77_), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_247_i15_ in range(default__.abs(-424))])) + (d_109_v2_)), globalState):
                with _dafny.c_label("1"):
                    if (d_201_i14_) >= (100):
                        raise _dafny.Break("1")
                    d_201_i14_ = (d_201_i14_) + (1)
                    d_202_v79_: _dafny.MultiSet
                    d_202_v79_ = _dafny.MultiSet([d_197_v75_, default__.fm1(-202, d_107_v0_, not(not(d_107_v0_)), globalState)])
                    d_202_v79_ = d_202_v79_
                    d_203_v80_: _dafny.MultiSet
                    d_203_v80_ = _dafny.MultiSet([d_107_v0_])
                    d_204_v81_: _dafny.Array
                    def lambda0_(d_205_i16_):
                        return (d_205_i16_) * (392)

                    init0_ = lambda0_
                    nw23_ = _dafny.Array(None, 10)
                    for i0_0_ in range(nw23_.length(0)):
                        nw23_[i0_0_] = init0_(i0_0_)
                    d_204_v81_ = nw23_
                    d_206_v82_: _dafny.Set
                    d_206_v82_ = _dafny.Set({d_204_v81_})
                    d_207_v83_: _dafny.Array
                    nw24_ = _dafny.Array(None, 3)
                    nw24_[int(0)] = d_197_v75_
                    nw24_[int(1)] = (d_203_v80_).cardinality
                    nw24_[int(2)] = (0) - (default__.safeDivisionInt(len(d_109_v2_), len(d_206_v82_)))
                    d_207_v83_ = nw24_
                    index9_ = default__.safeIndex(385, (d_204_v81_).length(0))
                    (d_204_v81_)[index9_] = d_197_v75_
                    d_208_v84_: C1
                    nw25_ = C1()
                    nw25_.ctor__(d_107_v0_)
                    d_208_v84_ = nw25_
                    d_209_v85_: _dafny.Set
                    d_209_v85_ = _dafny.Set({d_208_v84_})
                    d_210_v86_: D1
                    d_210_v86_ = D1_DC3(len(d_209_v85_), d_197_v75_, d_197_v75_, d_197_v75_, len(d_109_v2_))
                    d_211_v87_: _dafny.Seq
                    d_211_v87_ = _dafny.SeqWithoutIsStrInference([d_210_v86_, D1_DC3(d_197_v75_, len(d_109_v2_), 494, default__.fm1(d_197_v75_, d_107_v0_, d_107_v0_, globalState), d_197_v75_), d_210_v86_, d_210_v86_, d_210_v86_])
                    index10_ = default__.safeIndex(703, (d_204_v81_).length(0))
                    (d_204_v81_)[index10_] = default__.fm1((_dafny.MultiSet(d_211_v87_)).cardinality, False, d_107_v0_, globalState)
                    d_212_v88_: str
                    d_212_v88_ = _dafny.CodePoint('n')
                    d_213_v89_: _dafny.Map
                    d_213_v89_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([d_107_v0_, d_107_v0_])).set(default__.safeIndex(398, len(_dafny.SeqWithoutIsStrInference([d_107_v0_, d_107_v0_]))), d_107_v0_)): d_197_v75_})
                    d_214_v90_: _dafny.Seq
                    d_214_v90_ = _dafny.SeqWithoutIsStrInference([d_213_v89_, d_213_v89_])
                    d_215_v91_: T1
                    nw26_ = C4()
                    nw26_.ctor__(d_212_v88_, d_197_v75_, d_107_v0_, d_109_v2_, d_214_v90_)
                    d_215_v91_ = nw26_
                    d_216_v92_: T0
                    nw27_ = C4()
                    nw27_.ctor__(d_212_v88_, d_197_v75_, d_107_v0_, d_109_v2_, d_214_v90_)
                    d_216_v92_ = nw27_
                    d_217_v93_: _dafny.Map
                    d_217_v93_ = _dafny.Map({d_216_v92_: d_215_v91_})
                    d_218_v94_: _dafny.Seq
                    d_218_v94_ = _dafny.SeqWithoutIsStrInference([d_215_v91_, ((d_217_v93_)[d_216_v92_] if (d_216_v92_) in (d_217_v93_) else d_215_v91_), d_215_v91_])
                    d_219_v95_: _dafny.MultiSet
                    d_219_v95_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_215_v91_, d_215_v91_]), d_218_v94_, _dafny.SeqWithoutIsStrInference([d_215_v91_, d_215_v91_])])
                    d_220_v96_: D28
                    d_220_v96_ = D28_DC74((d_219_v95_ if d_107_v0_ else d_219_v95_))
                    index11_ = default__.safeIndex(825, (d_108_v1_).length(0))
                    (d_108_v1_)[index11_] = not(False)
                    d_221_v97_: D23
                    d_221_v97_ = D23_DC61(d_108_v1_, d_107_v0_, True, (d_215_v91_).f18)
                    d_222_v98_: _dafny.Seq
                    d_222_v98_ = _dafny.SeqWithoutIsStrInference([d_221_v97_, d_221_v97_])
                    d_223_v99_: _dafny.Map
                    d_223_v99_ = _dafny.Map({103: _dafny.SeqWithoutIsStrInference([d_221_v97_, d_221_v97_, d_221_v97_, D23_DC61(d_108_v1_, d_107_v0_, d_107_v0_, d_107_v0_)])})
                    d_224_v100_: _dafny.MultiSet
                    d_224_v100_ = _dafny.MultiSet([(_dafny.Map({341: d_222_v98_})) | (d_223_v99_)])
                    d_225_v101_: _dafny.Set
                    d_225_v101_ = _dafny.Set({(d_216_v92_).f18, (d_215_v91_).f18, d_107_v0_, d_107_v0_})
                    d_226_v102_: _dafny.Seq
                    d_226_v102_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_197_v75_, d_197_v75_, d_197_v75_, d_197_v75_, 255}), d_199_v77_, d_199_v77_, d_199_v77_, default__.fm49(len(d_225_v101_), 810, d_197_v75_, globalState)])
                    d_227_v103_: _dafny.Seq
                    d_227_v103_ = _dafny.SeqWithoutIsStrInference([d_197_v75_])
                    d_228_v104_: D1
                    d_228_v104_ = D1_DC5(D1_DC3(d_197_v75_, d_197_v75_, len(d_213_v89_), (_dafny.MultiSet([(d_216_v92_).f18])).cardinality, (d_227_v103_)[default__.safeIndex(d_197_v75_, len(d_227_v103_))]))
                    d_229_v105_: D1
                    d_229_v105_ = D1_DC5(d_228_v104_)
                    d_230_v106_: D1
                    d_230_v106_ = D1_DC5(d_229_v105_)
                    d_231_v107_: D1
                    d_231_v107_ = D1_DC5(D1_DC5(d_230_v106_))
                    d_232_v108_: _dafny.Map
                    d_232_v108_ = _dafny.Map({d_231_v107_: d_197_v75_})
                    index12_ = default__.safeIndex(385, (d_204_v81_).length(0))
                    index13_ = default__.safeIndex(703, (d_204_v81_).length(0))
                    index14_ = default__.safeIndex(825, (d_108_v1_).length(0))
                    rhs16_ = ((d_224_v100_)[(d_223_v99_) | (d_223_v99_)] if ((d_223_v99_) | (d_223_v99_)) in (d_224_v100_) else len(d_226_v102_))
                    rhs17_ = ((d_232_v108_)[D1_DC5(d_229_v105_)] if (D1_DC5(d_229_v105_)) in (d_232_v108_) else d_197_v75_)
                    rhs18_ = d_220_v96_
                    rhs19_ = not(False)
                    lhs5_ = d_204_v81_
                    lhs6_ = default__.safeIndex(385, (d_204_v81_).length(0))
                    lhs7_ = d_204_v81_
                    lhs8_ = default__.safeIndex(703, (d_204_v81_).length(0))
                    lhs9_ = d_108_v1_
                    lhs10_ = default__.safeIndex(825, (d_108_v1_).length(0))
                    lhs5_[lhs6_] = rhs16_
                    lhs7_[lhs8_] = rhs17_
                    d_220_v96_ = rhs18_
                    lhs9_[lhs10_] = rhs19_
                    d_233_v109_: D16
                    d_233_v109_ = D16_DC40((d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))], d_197_v75_, (d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))], (d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))])
                    d_234_v110_: _dafny.Map
                    d_234_v110_ = _dafny.Map({(d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))]: d_233_v109_})
                    d_235_v113_: _dafny.Array
                    nw28_ = _dafny.Array(None, 9)
                    nw28_[int(0)] = d_234_v110_
                    nw28_[int(1)] = d_234_v110_
                    nw28_[int(2)] = d_234_v110_
                    nw28_[int(3)] = d_234_v110_
                    nw28_[int(4)] = d_234_v110_
                    def iife20_():
                        coll20_ = _dafny.Map()
                        compr_20_: int
                        for compr_20_ in _dafny.IntegerRange(992, -148):
                            d_236_v111_: int = compr_20_
                            if ((992) <= (d_236_v111_)) and ((d_236_v111_) < (-148)):
                                coll20_[default__.safeDivisionInt(d_236_v111_, -711)] = d_233_v109_
                        return _dafny.Map(coll20_)
                    nw28_[int(5)] = iife20_()
                    
                    nw28_[int(6)] = d_234_v110_
                    def iife21_():
                        coll21_ = _dafny.Map()
                        compr_21_: int
                        for compr_21_ in (d_213_v89_).keys.Elements:
                            d_237_v112_: int = compr_21_
                            if (d_237_v112_) in (d_213_v89_):
                                coll21_[(d_237_v112_) * (d_197_v75_)] = d_233_v109_
                        return _dafny.Map(coll21_)
                    nw28_[int(7)] = iife21_()
                    
                    nw28_[int(8)] = d_234_v110_
                    d_235_v113_ = nw28_
                    d_238_v114_: D26
                    d_238_v114_ = D26_DC70(d_197_v75_, (d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))], d_107_v0_, d_235_v113_)
                    d_239_v115_: C5
                    nw29_ = C5()
                    nw29_.ctor__(d_109_v2_, d_109_v2_, (d_215_v91_).f18)
                    d_239_v115_ = nw29_
                    d_240_v116_: _dafny.Seq
                    d_240_v116_ = _dafny.SeqWithoutIsStrInference([d_239_v115_, d_239_v115_, d_239_v115_])
                    d_241_v117_: _dafny.Map
                    d_241_v117_ = _dafny.Map({(d_215_v91_).f18: d_107_v0_})
                    d_242_v118_: _dafny.Array
                    nw30_ = _dafny.Array(None, 21)
                    nw30_[int(0)] = d_238_v114_
                    nw30_[int(1)] = D26_DC70(d_197_v75_, 617, not((d_215_v91_).f18), d_235_v113_)
                    nw30_[int(2)] = d_238_v114_
                    def iife22_(_pat_let0_0):
                        def iife23_(d_243_dt__update__tmp_h0_):
                            def iife24_(_pat_let1_0):
                                def iife25_(d_244_dt__update_hcf123_h0_):
                                    return D26_DC70((d_243_dt__update__tmp_h0_).cf122, d_244_dt__update_hcf123_h0_, (d_243_dt__update__tmp_h0_).cf124, (d_243_dt__update__tmp_h0_).cf125)
                                return iife25_(_pat_let1_0)
                            return iife24_(871)
                        return iife23_(_pat_let0_0)
                    nw30_[int(3)] = iife22_(d_238_v114_)
                    nw30_[int(4)] = d_238_v114_
                    nw30_[int(5)] = d_238_v114_
                    nw30_[int(6)] = D26_DC70((0) - (d_197_v75_), (d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))], (d_215_v91_).f18, d_235_v113_)
                    nw30_[int(7)] = D26_DC70(d_197_v75_, len(d_240_v116_), False, d_235_v113_)
                    nw30_[int(8)] = d_238_v114_
                    nw30_[int(9)] = d_238_v114_
                    nw30_[int(10)] = d_238_v114_
                    nw30_[int(11)] = D26_DC70(len(_dafny.Set({(d_108_v1_)[default__.safeIndex(825, (d_108_v1_).length(0))]})), (d_203_v80_).cardinality, (d_215_v91_).f18, d_235_v113_)
                    nw30_[int(12)] = D26_DC70(len((d_213_v89_).set((d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))], (d_204_v81_)[default__.safeIndex(385, (d_204_v81_).length(0))])), d_197_v75_, (d_215_v91_).f18, d_235_v113_)
                    nw30_[int(13)] = D26_DC70(len(_dafny.SeqWithoutIsStrInference([(d_108_v1_)[default__.safeIndex(825, (d_108_v1_).length(0))]])), len(d_241_v117_), d_107_v0_, d_235_v113_)
                    nw30_[int(14)] = d_238_v114_
                    nw30_[int(15)] = d_238_v114_
                    nw30_[int(16)] = d_238_v114_
                    nw30_[int(17)] = d_238_v114_
                    nw30_[int(18)] = d_238_v114_
                    nw30_[int(19)] = d_238_v114_
                    nw30_[int(20)] = d_238_v114_
                    d_242_v118_ = nw30_
                    index15_ = default__.safeIndex(212, (d_242_v118_).length(0))
                    (d_242_v118_)[index15_] = (d_238_v114_ if (d_239_v115_).fm15(len(d_199_v77_), d_197_v75_, globalState) else d_238_v114_)
                    d_245_v119_: _dafny.Array
                    nw31_ = _dafny.Array(False, 19)
                    d_245_v119_ = nw31_
                    d_246_v120_: _dafny.Map
                    d_246_v120_ = _dafny.Map({d_245_v119_: d_212_v88_})
                    index16_ = default__.safeIndex(212, (d_242_v118_).length(0))
                    (d_242_v118_)[index16_] = D26_DC70(d_197_v75_, len(d_246_v120_), d_107_v0_, d_235_v113_)
                    d_212_v88_ = d_212_v88_
                    pass
            pass
        r0 = (0) - (d_197_v75_)
        r1 = (0) - ((0) - (d_197_v75_))
        d_248_v121_: _dafny.Seq
        d_248_v121_ = _dafny.SeqWithoutIsStrInference([d_107_v0_, d_107_v0_])
        d_249_v122_: _dafny.Seq
        d_249_v122_ = _dafny.SeqWithoutIsStrInference([d_197_v75_, d_197_v75_, len(d_248_v121_)])
        d_250_v123_: _dafny.Map
        d_250_v123_ = _dafny.Map({d_107_v0_: d_249_v122_})
        d_251_v124_: _dafny.Set
        d_251_v124_ = _dafny.Set({d_107_v0_, d_107_v0_})
        d_252_v125_: _dafny.Map
        d_252_v125_ = _dafny.Map({d_251_v124_: d_197_v75_})
        d_253_v126_: str
        d_253_v126_ = _dafny.CodePoint('v')
        d_254_v127_: _dafny.Map
        d_254_v127_ = _dafny.Map({d_197_v75_: len(d_109_v2_)})
        d_255_v128_: T0
        nw32_ = C9()
        nw32_.ctor__(d_108_v1_, True)
        d_255_v128_ = nw32_
        d_256_v129_: _dafny.Map
        d_256_v129_ = _dafny.Map({d_255_v128_: len(_dafny.SeqWithoutIsStrInference([d_253_v126_ for d_257_i18_ in range(default__.abs(105))]))})
        d_258_v130_: D8
        d_258_v130_ = D8_DC21(d_253_v126_, default__.fm1(d_197_v75_, d_107_v0_, (d_255_v128_).f18, globalState))
        d_259_v131_: _dafny.MultiSet
        d_259_v131_ = _dafny.MultiSet([default__.fm1(d_197_v75_, d_107_v0_, d_107_v0_, globalState)])
        d_260_v132_: _dafny.MultiSet
        d_260_v132_ = _dafny.MultiSet([d_253_v126_])
        d_261_v133_: _dafny.Array
        nw33_ = _dafny.Array(None, 26)
        nw33_[int(0)] = 87
        nw33_[int(1)] = len((d_250_v123_) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([d_197_v75_ for d_262_i17_ in range(default__.abs(-21))])})))
        nw33_[int(2)] = d_197_v75_
        nw33_[int(3)] = (0) - (((d_252_v125_)[d_251_v124_] if (d_251_v124_) in (d_252_v125_) else d_197_v75_))
        nw33_[int(4)] = d_197_v75_
        nw33_[int(5)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sihya"))).set(default__.safeIndex(d_197_v75_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sihya")))), d_253_v126_))
        nw33_[int(6)] = (0) - ((0) - (default__.safeModuloInt(d_197_v75_, len(d_249_v122_))))
        nw33_[int(7)] = d_197_v75_
        nw33_[int(8)] = 735
        nw33_[int(9)] = d_197_v75_
        nw33_[int(10)] = (len(d_109_v2_)) * (d_197_v75_)
        nw33_[int(11)] = -405
        nw33_[int(12)] = 232
        nw33_[int(13)] = ((0) - (len(d_254_v127_))) + (((d_256_v129_)[d_255_v128_] if (d_255_v128_) in (d_256_v129_) else (0) - (d_197_v75_)))
        nw33_[int(14)] = len(d_109_v2_)
        nw33_[int(15)] = d_197_v75_
        nw33_[int(16)] = len((d_248_v121_) + (d_248_v121_))
        nw33_[int(17)] = (d_258_v130_).cf42
        nw33_[int(18)] = (d_197_v75_) * (d_197_v75_)
        nw33_[int(19)] = d_197_v75_
        nw33_[int(20)] = d_197_v75_
        nw33_[int(21)] = (d_259_v131_).cardinality
        nw33_[int(22)] = default__.safeDivisionInt(702, (d_260_v132_).cardinality)
        nw33_[int(23)] = ((d_259_v131_)[len(d_109_v2_)] if (len(d_109_v2_)) in (d_259_v131_) else len(default__.fm49(d_197_v75_, d_197_v75_, d_197_v75_, globalState)))
        nw33_[int(24)] = default__.safeDivisionInt(d_197_v75_, (0) - (len(d_109_v2_)))
        nw33_[int(25)] = (d_197_v75_) + (d_197_v75_)
        d_261_v133_ = nw33_
        r2 = d_261_v133_
        def iife26_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in (d_199_v77_).Elements:
                d_263_v134_: int = compr_22_
                if (d_263_v134_) in (d_199_v77_):
                    coll22_[default__.safeModuloInt(d_263_v134_, d_197_v75_)] = d_107_v0_
            return _dafny.Map(coll22_)
        r3 = iife26_()
        
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_264_v0_: _dafny.Seq
        d_264_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewld"))
        d_265_v1_: _dafny.Map
        d_265_v1_ = _dafny.Map({True: True})
        d_266_v2_: int
        d_266_v2_ = -284
        d_267_v3_: bool
        d_267_v3_ = False
        d_268_v4_: _dafny.Array
        nw34_ = _dafny.Array(_dafny.Set({}), 5)
        d_268_v4_ = nw34_
        d_269_v5_: _dafny.Seq
        d_269_v5_ = _dafny.SeqWithoutIsStrInference([d_266_v2_])
        d_270_v6_: _dafny.Map
        d_270_v6_ = _dafny.Map({d_267_v3_: d_266_v2_})
        d_271_v7_: _dafny.Seq
        d_271_v7_ = _dafny.SeqWithoutIsStrInference([d_267_v3_, d_267_v3_])
        d_272_v8_: _dafny.Array
        nw35_ = _dafny.Array(None, 27)
        nw35_[int(0)] = len(d_269_v5_)
        nw35_[int(1)] = 51
        nw35_[int(2)] = d_266_v2_
        nw35_[int(3)] = d_266_v2_
        nw35_[int(4)] = d_266_v2_
        nw35_[int(5)] = d_266_v2_
        nw35_[int(6)] = ((d_270_v6_)[not(d_267_v3_)] if (not(d_267_v3_)) in (d_270_v6_) else d_266_v2_)
        nw35_[int(7)] = d_266_v2_
        nw35_[int(8)] = d_266_v2_
        nw35_[int(9)] = len(d_271_v7_)
        nw35_[int(10)] = -987
        nw35_[int(11)] = 466
        nw35_[int(12)] = len(d_264_v0_)
        nw35_[int(13)] = d_266_v2_
        nw35_[int(14)] = d_266_v2_
        nw35_[int(15)] = d_266_v2_
        nw35_[int(16)] = 666
        nw35_[int(17)] = d_266_v2_
        nw35_[int(18)] = d_266_v2_
        nw35_[int(19)] = d_266_v2_
        nw35_[int(20)] = d_266_v2_
        nw35_[int(21)] = d_266_v2_
        nw35_[int(22)] = len(d_264_v0_)
        nw35_[int(23)] = len(_dafny.SeqWithoutIsStrInference([d_267_v3_]))
        nw35_[int(24)] = d_266_v2_
        nw35_[int(25)] = -606
        nw35_[int(26)] = d_266_v2_
        d_272_v8_ = nw35_
        d_273_v9_: _dafny.Seq
        d_273_v9_ = _dafny.SeqWithoutIsStrInference([d_272_v8_])
        d_274_v10_: _dafny.MultiSet
        d_274_v10_ = _dafny.MultiSet([d_272_v8_])
        d_275_v11_: _dafny.Array
        nw36_ = _dafny.Array(None, 5)
        nw36_[int(0)] = False
        nw36_[int(1)] = d_267_v3_
        nw36_[int(2)] = not(not(True))
        nw36_[int(3)] = d_267_v3_
        nw36_[int(4)] = d_267_v3_
        d_275_v11_ = nw36_
        d_276_v12_: str
        d_276_v12_ = _dafny.CodePoint('l')
        d_277_v13_: _dafny.Map
        d_277_v13_ = _dafny.Map({d_275_v11_: d_276_v12_})
        d_278_globalState_: GlobalState
        nw37_ = GlobalState()
        nw37_.ctor__(644, False, True, 888, 99, False, 83, False, -470, d_264_v0_, (_dafny.SeqWithoutIsStrInference([len(d_265_v1_)])) + (_dafny.SeqWithoutIsStrInference([d_266_v2_, d_266_v2_, (0) - (d_266_v2_), (_dafny.MultiSet([d_267_v3_, d_267_v3_, d_267_v3_])).cardinality])), True, d_264_v0_, (d_268_v4_ if True else d_268_v4_), (d_273_v9_) + (d_273_v9_), 523, d_274_v10_, d_277_v13_)
        d_278_globalState_ = nw37_
        d_279_v14_: _dafny.Set
        d_279_v14_ = _dafny.Set({d_267_v3_})
        hi0_ = len(d_279_v14_)
        for d_280_i0_ in range(d_266_v2_, hi0_):
            if d_267_v3_:
                d_281_v15_: _dafny.Map
                d_281_v15_ = _dafny.Map({d_266_v2_: not(d_267_v3_)})
                d_282_v16_: _dafny.Set
                d_282_v16_ = _dafny.Set({len(d_270_v6_)})
                d_281_v15_ = _dafny.Map({(len(d_269_v5_)) + (len(d_282_v16_)): d_267_v3_})
                d_283_v17_: _dafny.Map
                d_283_v17_ = _dafny.Map({default__.fm0(d_278_globalState_): (0) - (len(d_282_v16_))})
                d_284_v18_: _dafny.Seq
                d_284_v18_ = _dafny.SeqWithoutIsStrInference([d_281_v15_])
                d_283_v17_ = (d_283_v17_).set(d_284_v18_, d_266_v2_)
                d_285_v19_: _dafny.Array
                nw38_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_285_v19_ = nw38_
                d_286_v20_: _dafny.Map
                d_286_v20_ = _dafny.Map({d_266_v2_: d_275_v11_})
                index17_ = default__.safeIndex(212, (d_285_v19_).length(0))
                (d_285_v19_)[index17_] = ((d_286_v20_)[152] if (152) in (d_286_v20_) else d_275_v11_)
                index18_ = default__.safeIndex(212, (d_285_v19_).length(0))
                (d_285_v19_)[index18_] = d_275_v11_
                (d_278_globalState_).f0 = d_280_i0_
                d_287_v21_: int
                d_288_v22_: int
                d_289_v23_: _dafny.Array
                d_290_v24_: _dafny.Map
                out0_: int
                out1_: int
                out2_: _dafny.Array
                out3_: _dafny.Map
                out0_, out1_, out2_, out3_ = default__.m0(d_278_globalState_)
                d_287_v21_ = out0_
                d_288_v22_ = out1_
                d_289_v23_ = out2_
                d_290_v24_ = out3_
            elif True:
                index19_ = default__.safeIndex(331, (d_275_v11_).length(0))
                (d_275_v11_)[index19_] = d_267_v3_
                index20_ = default__.safeIndex(331, (d_275_v11_).length(0))
                (d_275_v11_)[index20_] = (d_267_v3_ if d_267_v3_ else not(True))
                d_291_v25_: int
                d_292_v26_: int
                d_293_v27_: _dafny.Array
                d_294_v28_: _dafny.Map
                out4_: int
                out5_: int
                out6_: _dafny.Array
                out7_: _dafny.Map
                out4_, out5_, out6_, out7_ = default__.m0(d_278_globalState_)
                d_291_v25_ = out4_
                d_292_v26_ = out5_
                d_293_v27_ = out6_
                d_294_v28_ = out7_
                index21_ = default__.safeIndex(698, (d_293_v27_).length(0))
                (d_293_v27_)[index21_] = d_291_v25_
                index22_ = default__.safeIndex(698, (d_293_v27_).length(0))
                (d_293_v27_)[index22_] = default__.fm1((default__.fm1(-45, (d_275_v11_)[default__.safeIndex(331, (d_275_v11_).length(0))], False, d_278_globalState_) if d_267_v3_ else d_266_v2_), d_267_v3_, d_267_v3_, d_278_globalState_)
                index23_ = default__.safeIndex(331, (d_275_v11_).length(0))
                (d_275_v11_)[index23_] = (d_267_v3_) == (True)
                d_295_v29_: _dafny.Map
                d_295_v29_ = _dafny.Map({d_280_i0_: (d_293_v27_)[default__.safeIndex(698, (d_293_v27_).length(0))]})
                index24_ = default__.safeIndex(331, (d_275_v11_).length(0))
                rhs20_ = ((d_295_v29_) | (d_295_v29_)) | (d_295_v29_)
                rhs21_ = (d_291_v25_) <= ((748) * (d_291_v25_))
                rhs22_ = (d_271_v7_).set(default__.safeIndex((0) - ((210) - (len(d_294_v28_))), len(d_271_v7_)), (d_275_v11_)[default__.safeIndex(331, (d_275_v11_).length(0))])
                lhs11_ = d_275_v11_
                lhs12_ = default__.safeIndex(331, (d_275_v11_).length(0))
                d_295_v29_ = rhs20_
                lhs11_[lhs12_] = rhs21_
                d_271_v7_ = rhs22_
            d_275_v11_ = d_275_v11_
            (d_278_globalState_).f3 = d_266_v2_
            d_296_v30_: _dafny.Map
            d_296_v30_ = _dafny.Map({d_280_i0_: d_280_i0_})
            d_297_v31_: _dafny.Seq
            d_297_v31_ = _dafny.SeqWithoutIsStrInference([d_296_v30_])
            d_298_v32_: C4
            nw39_ = C4()
            nw39_.ctor__(_dafny.CodePoint('p'), d_266_v2_, (d_264_v0_) == (d_264_v0_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yennni")), (d_297_v31_) + (d_297_v31_))
            d_298_v32_ = nw39_
        if not (not((d_266_v2_) == (d_266_v2_))) or (d_267_v3_):
            d_299_v33_: int
            d_300_v34_: int
            d_301_v35_: _dafny.Array
            d_302_v36_: _dafny.Map
            out8_: int
            out9_: int
            out10_: _dafny.Array
            out11_: _dafny.Map
            out8_, out9_, out10_, out11_ = default__.m0(d_278_globalState_)
            d_299_v33_ = out8_
            d_300_v34_ = out9_
            d_301_v35_ = out10_
            d_302_v36_ = out11_
            d_303_v37_: int
            d_304_v38_: int
            d_305_v39_: _dafny.Array
            d_306_v40_: _dafny.Map
            out12_: int
            out13_: int
            out14_: _dafny.Array
            out15_: _dafny.Map
            out12_, out13_, out14_, out15_ = default__.m0(d_278_globalState_)
            d_303_v37_ = out12_
            d_304_v38_ = out13_
            d_305_v39_ = out14_
            d_306_v40_ = out15_
            d_265_v1_ = (d_265_v1_).set(d_267_v3_, d_267_v3_)
            d_275_v11_ = d_275_v11_
            rhs23_ = len(d_264_v0_)
            rhs24_ = d_300_v34_
            lhs13_ = d_278_globalState_
            d_300_v34_ = rhs23_
            lhs13_.f0 = rhs24_
        elif True:
            d_307_v41_: D9
            d_307_v41_ = D9_DC23(d_267_v3_, d_267_v3_, (0) - (d_266_v2_))
            d_308_v42_: _dafny.Set
            d_308_v42_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_267_v3_, d_267_v3_]))})
            d_309_v43_: _dafny.MultiSet
            d_309_v43_ = _dafny.MultiSet([(d_307_v41_).cf44, d_267_v3_, default__.fm23(d_308_v42_, d_266_v2_, d_278_globalState_), d_267_v3_])
            d_310_v44_: D0
            d_310_v44_ = D0_DC0(d_267_v3_)
            (d_278_globalState_).f0 = (len(default__.fm47(d_309_v43_, (0) - (d_266_v2_), d_266_v2_, (d_310_v44_).cf0, d_278_globalState_))) * (d_266_v2_)
            (d_278_globalState_).f3 = len(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_276_v12_ for d_311_i1_ in range(default__.abs(404))]), d_264_v0_}))
            d_312_v45_: _dafny.Array
            nw40_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_312_v45_ = nw40_
            d_313_v46_: _dafny.Array
            nw41_ = _dafny.Array(None, 1)
            nw41_[int(0)] = d_275_v11_
            d_313_v46_ = nw41_
            d_314_v47_: D26
            d_314_v47_ = D26_DC71(len(d_279_v14_), d_313_v46_, d_266_v2_)
            d_315_v48_: _dafny.MultiSet
            d_315_v48_ = _dafny.MultiSet([d_266_v2_])
            d_316_v49_: _dafny.Seq
            d_316_v49_ = _dafny.SeqWithoutIsStrInference([d_312_v45_, d_312_v45_])
            rhs25_ = d_267_v3_
            rhs26_ = len(_dafny.SeqWithoutIsStrInference([d_267_v3_, (d_266_v2_) >= ((d_314_v47_).cf128), d_267_v3_]))
            rhs27_ = (d_312_v45_ if (d_315_v48_).issubset(_dafny.MultiSet([d_266_v2_])) else (d_316_v49_)[default__.safeIndex(d_266_v2_, len(d_316_v49_))])
            lhs14_ = d_278_globalState_
            lhs15_ = d_278_globalState_
            lhs14_.f2 = rhs25_
            lhs15_.f0 = rhs26_
            d_312_v45_ = rhs27_
            d_317_v50_: _dafny.Map
            d_317_v50_ = _dafny.Map({d_266_v2_: default__.safeModuloInt(d_266_v2_, d_266_v2_)})
            d_317_v50_ = _dafny.Map({d_266_v2_: 30})
            (d_278_globalState_).f9 = (d_264_v0_) + (d_264_v0_)
        if ((d_267_v3_) == (d_267_v3_)) in (_dafny.MultiSet([d_267_v3_])):
            d_267_v3_ = d_267_v3_
            d_318_v51_: _dafny.Map
            d_318_v51_ = _dafny.Map({d_267_v3_: d_276_v12_})
            d_319_v52_: _dafny.Map
            d_319_v52_ = _dafny.Map({d_318_v51_: d_267_v3_})
            (d_278_globalState_).f2 = ((d_319_v52_)[_dafny.Map({d_267_v3_: d_276_v12_})] if (_dafny.Map({d_267_v3_: d_276_v12_})) in (d_319_v52_) else d_267_v3_)
            (d_278_globalState_).f15 = default__.safeModuloInt(d_266_v2_, d_266_v2_)
            index25_ = default__.safeIndex(163, (d_272_v8_).length(0))
            (d_272_v8_)[index25_] = len((d_269_v5_) + (d_269_v5_))
            index26_ = default__.safeIndex(163, (d_272_v8_).length(0))
            (d_272_v8_)[index26_] = d_266_v2_
            d_320_v53_: int
            d_321_v54_: int
            d_322_v55_: _dafny.Array
            d_323_v56_: _dafny.Map
            out16_: int
            out17_: int
            out18_: _dafny.Array
            out19_: _dafny.Map
            out16_, out17_, out18_, out19_ = default__.m0(d_278_globalState_)
            d_320_v53_ = out16_
            d_321_v54_ = out17_
            d_322_v55_ = out18_
            d_323_v56_ = out19_
        elif True:
            d_324_v57_: D1
            d_324_v57_ = D1_DC4(d_267_v3_)
            d_325_v58_: D2
            d_325_v58_ = D2_DC7(default__.fm1(135, (d_324_v57_).cf12, d_267_v3_, d_278_globalState_), d_267_v3_, ((0) - (-40)) + (len(d_264_v0_)))
            d_325_v58_ = d_325_v58_
            (d_278_globalState_).f2 = d_267_v3_
            (d_278_globalState_).f12 = (d_264_v0_) + (d_264_v0_)
            d_326_v59_: D1
            d_326_v59_ = D1_DC3((0) - (d_266_v2_), len(d_264_v0_), d_266_v2_, d_266_v2_, d_266_v2_)
            d_327_v60_: T0
            nw42_ = C5()
            nw42_.ctor__(d_264_v0_, d_264_v0_, d_267_v3_)
            d_327_v60_ = nw42_
            d_328_v61_: D1
            d_328_v61_ = D1_DC5((d_326_v59_ if d_267_v3_ else D1_DC2(d_327_v60_)))
            d_328_v61_ = d_328_v61_
            d_265_v1_ = (d_265_v1_).set((d_327_v60_).f18, (d_327_v60_).f18)
        d_276_v12_ = d_276_v12_
        hi1_ = d_266_v2_
        for d_329_i2_ in range((d_266_v2_) - (d_266_v2_), hi1_):
            d_330_v62_: _dafny.Array
            nw43_ = _dafny.Array(None, 21)
            d_330_v62_ = nw43_
            d_331_v63_: _dafny.Map
            d_331_v63_ = _dafny.Map({_dafny.CodePoint('i'): d_330_v62_})
            d_332_v64_: _dafny.Set
            d_332_v64_ = _dafny.Set({(0) - (d_266_v2_), len(d_331_v63_)})
            if default__.fm23(d_332_v64_, d_266_v2_, d_278_globalState_):
                (d_278_globalState_).f2 = not(d_267_v3_)
                d_267_v3_ = (D0_DC1(d_267_v3_, len(d_264_v0_), d_264_v0_, d_267_v3_, d_329_i2_)).cf1
                (d_278_globalState_).f0 = d_329_i2_
                (d_278_globalState_).f3 = (0) - (d_329_i2_)
                d_333_v65_: _dafny.MultiSet
                d_333_v65_ = _dafny.MultiSet([d_329_i2_, d_329_i2_])
                d_334_v66_: C8
                nw44_ = C8()
                nw44_.ctor__((d_333_v65_).cardinality, d_267_v3_)
                d_334_v66_ = nw44_
                d_335_v67_: T0
                nw45_ = C9()
                nw45_.ctor__(d_275_v11_, d_267_v3_)
                d_335_v67_ = nw45_
                rhs28_ = (d_335_v67_).f18
                rhs29_ = d_334_v66_
                rhs30_ = d_335_v67_
                rhs31_ = d_334_v66_.f23
                lhs16_ = d_278_globalState_
                d_267_v3_ = rhs28_
                d_334_v66_ = rhs29_
                d_335_v67_ = rhs30_
                lhs16_.f15 = rhs31_
            elif True:
                d_336_v68_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Set({}), 9)
                d_336_v68_ = nw46_
                index27_ = default__.safeIndex(992, (d_336_v68_).length(0))
                (d_336_v68_)[index27_] = d_332_v64_
                d_337_v69_: _dafny.Seq
                d_337_v69_ = _dafny.SeqWithoutIsStrInference([d_269_v5_, d_269_v5_, _dafny.SeqWithoutIsStrInference([d_329_i2_ for d_338_i3_ in range(default__.abs(-267))]), d_269_v5_])
                index28_ = default__.safeIndex(992, (d_336_v68_).length(0))
                rhs32_ = d_332_v64_
                rhs33_ = default__.fm1(d_329_i2_, d_267_v3_, False, d_278_globalState_)
                rhs34_ = d_337_v69_
                rhs35_ = d_272_v8_
                lhs17_ = d_336_v68_
                lhs18_ = default__.safeIndex(992, (d_336_v68_).length(0))
                lhs19_ = d_278_globalState_
                lhs17_[lhs18_] = rhs32_
                lhs19_.f15 = rhs33_
                d_337_v69_ = rhs34_
                d_272_v8_ = rhs35_
                index29_ = default__.safeIndex(12, (d_272_v8_).length(0))
                (d_272_v8_)[index29_] = (d_329_i2_) + (d_266_v2_)
                d_339_v70_: _dafny.MultiSet
                def iife27_(_pat_let2_0):
                    def iife28_(d_340_dt__update__tmp_h0_):
                        def iife29_(_pat_let3_0):
                            def iife30_(d_341_dt__update_hcf33_h0_):
                                return D3_DC13((d_340_dt__update__tmp_h0_).cf30, (d_340_dt__update__tmp_h0_).cf31, (d_340_dt__update__tmp_h0_).cf32, d_341_dt__update_hcf33_h0_, (d_340_dt__update__tmp_h0_).cf34)
                            return iife30_(_pat_let3_0)
                        return iife29_(828)
                    return iife28_(_pat_let2_0)
                d_339_v70_ = _dafny.MultiSet([(iife27_(D3_DC13(d_266_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppao")), d_329_i2_, 281, d_276_v12_))).cf31])
                index30_ = default__.safeIndex(12, (d_272_v8_).length(0))
                (d_272_v8_)[index30_] = ((d_266_v2_ if d_267_v3_ else d_329_i2_) if not(True) else ((d_339_v70_).cardinality) - (d_266_v2_))
                d_342_v71_: _dafny.Array
                nw47_ = _dafny.Array(int(0), 8)
                d_342_v71_ = nw47_
                d_343_v72_: _dafny.Set
                d_343_v72_ = _dafny.Set({d_342_v71_, d_342_v71_, d_342_v71_, (d_342_v71_ if d_267_v3_ else d_342_v71_)})
                d_344_v73_: _dafny.Array
                def lambda1_(d_345_v0_):
                    def lambda2_(d_346_i4_):
                        return d_345_v0_

                    return lambda2_

                init1_ = lambda1_(d_264_v0_)
                nw48_ = _dafny.Array(None, 23)
                for i0_1_ in range(nw48_.length(0)):
                    nw48_[i0_1_] = init1_(i0_1_)
                d_344_v73_ = nw48_
                index31_ = default__.safeIndex(120, (d_344_v73_).length(0))
                (d_344_v73_)[index31_] = (d_264_v0_) + ((_dafny.SeqWithoutIsStrInference([d_276_v12_ for d_347_i5_ in range(default__.abs(519))])).set(default__.safeIndex(d_266_v2_, len(_dafny.SeqWithoutIsStrInference([d_276_v12_ for d_348_i5_ in range(default__.abs(519))]))), d_276_v12_))
                index32_ = default__.safeIndex(120, (d_344_v73_).length(0))
                rhs36_ = default__.safeModuloInt((d_272_v8_)[default__.safeIndex(12, (d_272_v8_).length(0))], d_266_v2_)
                rhs37_ = d_343_v72_
                rhs38_ = (d_269_v5_)[default__.safeIndex((d_272_v8_)[default__.safeIndex(12, (d_272_v8_).length(0))], len(d_269_v5_))]
                rhs39_ = default__.fm16((len(d_264_v0_)) + (661), (d_279_v14_).isdisjoint(d_279_v14_), d_276_v12_, (d_272_v8_)[default__.safeIndex(12, (d_272_v8_).length(0))], d_278_globalState_)
                lhs20_ = d_278_globalState_
                lhs21_ = d_278_globalState_
                lhs22_ = d_344_v73_
                lhs23_ = default__.safeIndex(120, (d_344_v73_).length(0))
                lhs20_.f15 = rhs36_
                d_343_v72_ = rhs37_
                lhs21_.f0 = rhs38_
                lhs22_[lhs23_] = rhs39_
                d_349_v74_: _dafny.Set
                d_349_v74_ = _dafny.Set({d_269_v5_})
                d_350_v75_: C6
                nw49_ = C6()
                nw49_.ctor__(d_271_v7_, d_329_i2_, not(d_267_v3_), _dafny.SeqWithoutIsStrInference([d_276_v12_ for d_351_i6_ in range(default__.abs(652))]), _dafny.SeqWithoutIsStrInference([_dafny.Map({d_329_i2_: d_329_i2_}) for d_352_i7_ in range(default__.abs(999))]))
                d_350_v75_ = nw49_
                d_353_v76_: _dafny.Map
                d_353_v76_ = _dafny.Map({(d_349_v74_).ispropersubset(d_349_v74_): _dafny.Map({d_350_v75_: d_267_v3_})})
                d_354_v77_: _dafny.Map
                d_354_v77_ = _dafny.Map({d_267_v3_: d_353_v76_})
                d_355_v78_: D29
                d_355_v78_ = D29_DC76(((d_354_v77_)[d_267_v3_] if (d_267_v3_) in (d_354_v77_) else d_353_v76_))
                pat_let_tv0_ = d_353_v76_
                def iife31_(_pat_let4_0):
                    def iife32_(d_356_dt__update__tmp_h1_):
                        def iife33_(_pat_let5_0):
                            def iife34_(d_357_dt__update_hcf132_h0_):
                                return D29_DC76(d_357_dt__update_hcf132_h0_)
                            return iife34_(_pat_let5_0)
                        return iife33_(pat_let_tv0_)
                    return iife32_(_pat_let4_0)
                d_353_v76_ = (d_353_v76_) | ((iife31_(d_355_v78_)).cf132)
                d_358_v79_: _dafny.MultiSet
                d_358_v79_ = _dafny.MultiSet([661])
                d_359_v80_: T0
                nw50_ = C7()
                nw50_.ctor__(((d_358_v79_)[d_329_i2_] if (d_329_i2_) in (d_358_v79_) else len(_dafny.SeqWithoutIsStrInference([d_329_i2_]))), (d_276_v12_) not in (d_264_v0_))
                d_359_v80_ = nw50_
                d_359_v80_ = d_359_v80_
            d_360_v81_: _dafny.Map
            d_360_v81_ = _dafny.Map({d_267_v3_: d_276_v12_})
            d_361_v82_: _dafny.Map
            d_361_v82_ = _dafny.Map({d_267_v3_: d_360_v81_})
            d_362_v83_: D25
            d_362_v83_ = D25_DC66(d_361_v82_)
            d_363_v84_: D25
            d_363_v84_ = D25_DC68(d_362_v83_)
            source3_ = d_363_v84_
            if source3_.is_DC67:
                d_364___mcc_h0_ = source3_.cf116
                d_365___mcc_h1_ = source3_.cf117
                d_366___mcc_h2_ = source3_.cf118
                d_367___mcc_h3_ = source3_.cf119
                d_368_cf119_ = d_367___mcc_h3_
                d_369_cf118_ = d_366___mcc_h2_
                d_370_cf117_ = d_365___mcc_h1_
                d_371_cf116_ = d_364___mcc_h0_
                d_372_v85_: _dafny.Map
                d_372_v85_ = _dafny.Map({580: not(d_368_cf119_)})
                d_373_v86_: _dafny.Seq
                d_373_v86_ = _dafny.SeqWithoutIsStrInference([d_372_v85_])
                d_372_v85_ = (d_373_v86_)[default__.safeIndex(d_266_v2_, len(d_373_v86_))]
                d_374_v87_: int
                d_375_v88_: int
                d_376_v89_: _dafny.Array
                d_377_v90_: _dafny.Map
                out20_: int
                out21_: int
                out22_: _dafny.Array
                out23_: _dafny.Map
                out20_, out21_, out22_, out23_ = default__.m0(d_278_globalState_)
                d_374_v87_ = out20_
                d_375_v88_ = out21_
                d_376_v89_ = out22_
                d_377_v90_ = out23_
                (d_278_globalState_).f2 = d_368_cf119_
                index33_ = default__.safeIndex(958, (d_275_v11_).length(0))
                (d_275_v11_)[index33_] = (_dafny.CodePoint('b')) not in (d_264_v0_)
                index34_ = default__.safeIndex(958, (d_275_v11_).length(0))
                (d_275_v11_)[index34_] = (242) >= (d_375_v88_)
            elif source3_.is_DC66:
                d_378___mcc_h4_ = source3_.cf115
                d_379_cf115_ = d_378___mcc_h4_
                d_279_v14_ = d_279_v14_
                (d_278_globalState_).f2 = not(d_267_v3_)
                d_380_v91_: D17
                d_380_v91_ = D17_DC44(d_275_v11_, d_266_v2_)
                d_381_v92_: _dafny.Set
                d_381_v92_ = _dafny.Set({d_275_v11_, (d_380_v91_).cf77})
                (d_278_globalState_).f2 = (d_381_v92_).issubset((d_381_v92_) - (_dafny.Set({d_275_v11_})))
                d_382_v93_: _dafny.Map
                d_382_v93_ = _dafny.Map({d_329_i2_: d_275_v11_})
                d_382_v93_ = d_382_v93_
            elif True:
                d_383___mcc_h5_ = source3_.cf120
                d_384_cf120_ = d_383___mcc_h5_
                d_385_v94_: _dafny.MultiSet
                d_385_v94_ = _dafny.MultiSet([d_276_v12_, _dafny.CodePoint('r')])
                d_386_v95_: _dafny.Array
                def lambda3_(d_387_v2_, d_388_v0_):
                    def lambda4_(d_389_i8_):
                        return _dafny.Map({d_387_v2_: D16_DC40(d_387_v2_, -738, d_387_v2_, len(d_388_v0_))})

                    return lambda4_

                init2_ = lambda3_(d_266_v2_, d_264_v0_)
                nw51_ = _dafny.Array(None, 14)
                for i0_2_ in range(nw51_.length(0)):
                    nw51_[i0_2_] = init2_(i0_2_)
                d_386_v95_ = nw51_
                d_390_v96_: D26
                d_390_v96_ = D26_DC70(d_329_i2_, d_329_i2_, d_267_v3_, d_386_v95_)
                (d_278_globalState_).f9 = ((d_264_v0_) + ((d_264_v0_) + (d_264_v0_))).set(default__.safeIndex((((d_385_v94_)[d_276_v12_] if (d_276_v12_) in (d_385_v94_) else d_266_v2_)) * ((0) - ((d_390_v96_).cf123)), len((d_264_v0_) + ((d_264_v0_) + (d_264_v0_)))), d_276_v12_)
                d_391_v97_: _dafny.Array
                d_391_v97_ = d_272_v8_
                d_392_v98_: _dafny.Array
                nw52_ = _dafny.Array(None, 20)
                nw52_[int(0)] = d_272_v8_
                nw52_[int(1)] = d_272_v8_
                nw52_[int(2)] = d_272_v8_
                nw52_[int(3)] = d_272_v8_
                nw52_[int(4)] = d_272_v8_
                nw52_[int(5)] = d_272_v8_
                nw52_[int(6)] = d_272_v8_
                nw52_[int(7)] = d_272_v8_
                nw52_[int(8)] = d_272_v8_
                nw52_[int(9)] = d_272_v8_
                nw52_[int(10)] = d_272_v8_
                nw52_[int(11)] = d_272_v8_
                nw52_[int(12)] = d_272_v8_
                nw52_[int(13)] = (d_391_v97_)
                nw52_[int(14)] = d_272_v8_
                nw52_[int(15)] = d_272_v8_
                nw52_[int(16)] = d_272_v8_
                nw52_[int(17)] = d_272_v8_
                nw52_[int(18)] = d_272_v8_
                nw52_[int(19)] = d_272_v8_
                d_392_v98_ = nw52_
                d_392_v98_ = d_392_v98_
                d_393_v99_: _dafny.Map
                d_393_v99_ = _dafny.Map({d_329_i2_: d_267_v3_})
                d_393_v99_ = (d_393_v99_).set(len(d_269_v5_), True)
                d_394_v100_: D2
                d_394_v100_ = D2_DC7(d_266_v2_, d_267_v3_, d_329_i2_)
                d_395_v101_: T0
                nw53_ = C2()
                nw53_.ctor__(len(d_332_v64_), default__.fm40(d_394_v100_, d_278_globalState_), d_267_v3_)
                d_395_v101_ = nw53_
                d_396_v102_: _dafny.Seq
                d_396_v102_ = _dafny.SeqWithoutIsStrInference([d_395_v101_, d_395_v101_, d_395_v101_])
                d_397_v103_: _dafny.Map
                d_397_v103_ = _dafny.Map({(d_396_v102_)[default__.safeIndex((0) - (d_329_i2_), len(d_396_v102_))]: 155})
                d_397_v103_ = (d_397_v103_).set(d_395_v101_, d_329_i2_)
            d_398_v104_: int
            d_399_v105_: int
            d_400_v106_: _dafny.Array
            d_401_v107_: _dafny.Map
            out24_: int
            out25_: int
            out26_: _dafny.Array
            out27_: _dafny.Map
            out24_, out25_, out26_, out27_ = default__.m0(d_278_globalState_)
            d_398_v104_ = out24_
            d_399_v105_ = out25_
            d_400_v106_ = out26_
            d_401_v107_ = out27_
            (d_278_globalState_).f2 = not(d_267_v3_)
        index35_ = default__.safeIndex(173, (d_272_v8_).length(0))
        (d_272_v8_)[index35_] = d_266_v2_
        index36_ = default__.safeIndex(173, (d_272_v8_).length(0))
        (d_272_v8_)[index36_] = d_266_v2_
        d_402_v108_: C7
        nw54_ = C7()
        nw54_.ctor__((d_269_v5_)[default__.safeIndex(d_266_v2_, len(d_269_v5_))], d_267_v3_)
        d_402_v108_ = nw54_
        d_402_v108_ = d_402_v108_
        (d_278_globalState_).f0 = (d_272_v8_)[default__.safeIndex(173, (d_272_v8_).length(0))]
        (d_278_globalState_).f15 = d_266_v2_
        d_403_v109_: _dafny.Map
        d_403_v109_ = _dafny.Map({d_267_v3_: d_264_v0_})
        d_404_v110_: _dafny.Array
        nw55_ = _dafny.Array(None, 11)
        nw55_[int(0)] = d_264_v0_
        nw55_[int(1)] = d_264_v0_
        nw55_[int(2)] = _dafny.SeqWithoutIsStrInference([d_276_v12_ for d_405_i9_ in range(default__.abs(-354))])
        nw55_[int(3)] = d_264_v0_
        nw55_[int(4)] = d_264_v0_
        nw55_[int(5)] = d_264_v0_
        nw55_[int(6)] = (((d_403_v109_)[d_267_v3_] if (d_267_v3_) in (d_403_v109_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbdttde")))) + (d_264_v0_)
        nw55_[int(7)] = (d_264_v0_) + (d_264_v0_)
        nw55_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mp"))
        nw55_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sstrrcv")) if d_267_v3_ else d_264_v0_)
        nw55_[int(10)] = d_264_v0_
        d_404_v110_ = nw55_
        index37_ = default__.safeIndex(283, (d_404_v110_).length(0))
        (d_404_v110_)[index37_] = d_264_v0_
        index38_ = default__.safeIndex(283, (d_404_v110_).length(0))
        (d_404_v110_)[index38_] = d_264_v0_
        d_406_v111_: D17
        d_406_v111_ = D17_DC42(d_267_v3_)
        d_407_v112_: _dafny.Map
        d_407_v112_ = _dafny.Map({d_266_v2_: (d_406_v111_).cf74})
        hi2_ = (d_402_v108_.f24 if d_267_v3_ else 130)
        for d_408_i10_ in range(default__.safeDivisionInt(len(_dafny.Set({d_267_v3_})), len(d_407_v112_)), hi2_):
            (d_402_v108_).m1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scoeuba")), d_408_i10_, ((d_404_v110_)[default__.safeIndex(283, (d_404_v110_).length(0))]) + (d_264_v0_), d_278_globalState_)
            d_409_v113_: C5
            nw56_ = C5()
            nw56_.ctor__((d_404_v110_)[default__.safeIndex(283, (d_404_v110_).length(0))], (d_404_v110_)[default__.safeIndex(283, (d_404_v110_).length(0))], d_267_v3_)
            d_409_v113_ = nw56_
            d_410_v114_: _dafny.Map
            d_410_v114_ = _dafny.Map({(d_272_v8_)[default__.safeIndex(173, (d_272_v8_).length(0))]: d_266_v2_})
            d_411_v115_: _dafny.Map
            d_411_v115_ = _dafny.Map({d_409_v113_: len(d_410_v114_)})
            d_412_v116_: _dafny.Seq
            d_412_v116_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_267_v3_: len(d_411_v115_)})])
            (d_278_globalState_).f2 = not((_dafny.SeqWithoutIsStrInference([d_270_v6_, d_270_v6_])) <= (d_412_v116_))
            (d_409_v113_).m7((d_408_i10_) <= (len(d_269_v5_)), d_402_v108_.f24, d_278_globalState_)
            d_413_v117_: _dafny.Array
            nw57_ = _dafny.Array(None, 4)
            nw57_[int(0)] = (d_407_v112_) | (d_407_v112_)
            nw57_[int(1)] = d_407_v112_
            nw57_[int(2)] = d_407_v112_
            nw57_[int(3)] = d_407_v112_
            d_413_v117_ = nw57_
            def lambda5_(d_414_v112_):
                def lambda6_(d_415_i11_):
                    return d_414_v112_

                return lambda6_

            init3_ = lambda5_(d_407_v112_)
            nw58_ = _dafny.Array(None, 14)
            for i0_3_ in range(nw58_.length(0)):
                nw58_[i0_3_] = init3_(i0_3_)
            d_413_v117_ = nw58_
        d_416_v118_: _dafny.Map
        d_416_v118_ = _dafny.Map({d_276_v12_: (d_404_v110_)[default__.safeIndex(283, (d_404_v110_).length(0))]})
        d_417_v119_: _dafny.Array
        nw59_ = _dafny.Array(None, 24)
        d_417_v119_ = nw59_
        d_418_v120_: _dafny.Map
        d_418_v120_ = _dafny.Map({d_417_v119_: (d_404_v110_)[default__.safeIndex(283, (d_404_v110_).length(0))]})
        d_416_v118_ = (d_416_v118_).set(d_276_v12_, ((d_418_v120_)[d_417_v119_] if (d_417_v119_) in (d_418_v120_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "malcvx"))))
        index39_ = default__.safeIndex(635, (d_275_v11_).length(0))
        (d_275_v11_)[index39_] = d_267_v3_
        d_419_v121_: _dafny.Array
        def lambda7_(d_420_v12_):
            def lambda8_(d_421_i12_):
                return d_420_v12_

            return lambda8_

        init4_ = lambda7_(d_276_v12_)
        nw60_ = _dafny.Array(None, 8)
        for i0_4_ in range(nw60_.length(0)):
            nw60_[i0_4_] = init4_(i0_4_)
        d_419_v121_ = nw60_
        index40_ = default__.safeIndex(325, (d_419_v121_).length(0))
        (d_419_v121_)[index40_] = d_276_v12_
        d_422_v122_: _dafny.MultiSet
        d_422_v122_ = _dafny.MultiSet([(d_267_v3_) == (not(True))])
        d_423_v123_: _dafny.Set
        d_423_v123_ = _dafny.Set({d_402_v108_.f24, d_266_v2_, 469, d_266_v2_})
        index41_ = default__.safeIndex(635, (d_275_v11_).length(0))
        index42_ = default__.safeIndex(325, (d_419_v121_).length(0))
        rhs40_ = ((d_422_v122_ if d_267_v3_ else d_422_v122_)).isdisjoint(_dafny.MultiSet([False]))
        rhs41_ = d_276_v12_
        rhs42_ = (_dafny.MultiSet((d_271_v7_ if False else d_271_v7_))) | (d_422_v122_)
        rhs43_ = default__.fm23(d_423_v123_, -471, d_278_globalState_)
        rhs44_ = (d_264_v0_) < (d_264_v0_)
        lhs24_ = d_275_v11_
        lhs25_ = default__.safeIndex(635, (d_275_v11_).length(0))
        lhs26_ = d_419_v121_
        lhs27_ = default__.safeIndex(325, (d_419_v121_).length(0))
        lhs28_ = d_278_globalState_
        lhs24_[lhs25_] = rhs40_
        lhs26_[lhs27_] = rhs41_
        d_422_v122_ = rhs42_
        d_267_v3_ = rhs43_
        lhs28_.f2 = rhs44_
        (d_278_globalState_).f2 = (d_406_v111_).cf74
        d_424_v124_: _dafny.MultiSet
        d_424_v124_ = _dafny.MultiSet([(_dafny.Map({(d_275_v11_)[default__.safeIndex(635, (d_275_v11_).length(0))]: 871})) | (_dafny.Map({(d_275_v11_)[default__.safeIndex(635, (d_275_v11_).length(0))]: (d_422_v122_).cardinality}))])
        d_424_v124_ = (d_424_v124_).set(_dafny.Map({False: (d_272_v8_)[default__.safeIndex(173, (d_272_v8_).length(0))]}), default__.abs(d_266_v2_))
        index43_ = default__.safeIndex(173, (d_272_v8_).length(0))
        (d_272_v8_)[index43_] = (d_402_v108_.f24) + (default__.fm1(d_266_v2_, d_267_v3_, not(d_267_v3_), d_278_globalState_))
        _dafny.print((d_264_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v1_) == (_dafny.Map({True: True, False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_266_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_267_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v5_) == (_dafny.SeqWithoutIsStrInference([-284]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v6_) == (_dafny.Map({False: -284}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_v7_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v8_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_273_v9_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v10_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v11_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v11_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_276_v12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_277_v13_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_278_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_278_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_278_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_278_globalState_.f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_.f10) == (_dafny.SeqWithoutIsStrInference([1, -284, -284, 284, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_278_globalState_.f12).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_278_globalState_).f14)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_278_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_globalState_.f16).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_278_globalState_.f17)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_279_v14_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_402_v108_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_403_v109_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewld"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_404_v110_)[2]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_404_v110_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_406_v111_).cf74))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v112_) == (_dafny.Map({-284: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_416_v118_) == (_dafny.Map({_dafny.CodePoint('l'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewld"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_418_v120_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v121_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_422_v122_) == (_dafny.MultiSet([False, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_423_v123_) == (_dafny.Set({-284, 469}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_424_v124_) == (_dafny.MultiSet([_dafny.Map({True: 3}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284}), _dafny.Map({False: -284})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {self.cf3.VerbatimString(True)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
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
        return lambda: D1_DC3(int(0), int(0), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(int(0), False, int(0))
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
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)

class D2_DC7(D2, NamedTuple('DC7', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D3_DC13)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D3_DC14)

class D3_DC11(D3, NamedTuple('DC11', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC13(D3, NamedTuple('DC13', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC13({_dafny.string_of(self.cf30)}, {self.cf31.VerbatimString(True)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC13) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC14(D3, NamedTuple('DC14', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC14({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC14) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC15(D4, NamedTuple('DC15', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC16(D5, NamedTuple('DC16', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC18(D6, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC19(D7, NamedTuple('DC19', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(_dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC23(D9, NamedTuple('DC23', [('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC25(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)

class D10_DC25(D10, NamedTuple('DC25', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC28(D11, NamedTuple('DC28', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC31()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC31(D12, NamedTuple('DC31', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC33(False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC33(D13, NamedTuple('DC33', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf59', Any), ('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf59)}, {self.cf60.VerbatimString(True)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC32(D13, NamedTuple('DC32', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC36(D14, NamedTuple('DC36', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)

class D15_DC37(D15, NamedTuple('DC37', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC39(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)

class D16_DC39(D16, NamedTuple('DC39', [('cf67', Any), ('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC40(D16, NamedTuple('DC40', [('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC38(D16, NamedTuple('DC38', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC42(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)

class D17_DC42(D17, NamedTuple('DC42', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC43(D17, NamedTuple('DC43', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC44(D17, NamedTuple('DC44', [('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC41(D17, NamedTuple('DC41', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC46(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D18_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D18_DC45)

class D18_DC46(D18, NamedTuple('DC46', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC46({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC46) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D19_DC48(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D19_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D19_DC50)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)

class D19_DC48(D19, NamedTuple('DC48', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC49(D19, NamedTuple('DC49', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC49({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC49) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC50(D19, NamedTuple('DC50', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC50({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC50) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC51(D19, NamedTuple('DC51', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC53(False, False, _dafny.CodePoint('D'))
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

class D20_DC53(D20, NamedTuple('DC53', [('cf87', Any), ('cf88', Any), ('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC52(D20, NamedTuple('DC52', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC56(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)

class D21_DC56(D21, NamedTuple('DC56', [('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC58(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)

class D22_DC58(D22, NamedTuple('DC58', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC57(D22, NamedTuple('DC57', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC59(D22, NamedTuple('DC59', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC61(_dafny.Array(None, 0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D23_DC61)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D23_DC62)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D23_DC60)

class D23_DC61(D23, NamedTuple('DC61', [('cf99', Any), ('cf100', Any), ('cf101', Any), ('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC61({_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC61) and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC62(D23, NamedTuple('DC62', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC62({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC62) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC60(D23, NamedTuple('DC60', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC60({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC60) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC64(False, False, None, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D24_DC64)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D24_DC65)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D24_DC63)

class D24_DC64(D24, NamedTuple('DC64', [('cf105', Any), ('cf106', Any), ('cf107', Any), ('cf108', Any), ('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC64({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC64) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108 and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC65(D24, NamedTuple('DC65', [('cf110', Any), ('cf111', Any), ('cf112', Any), ('cf113', Any), ('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC65({_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC65) and self.cf110 == __o.cf110 and self.cf111 == __o.cf111 and self.cf112 == __o.cf112 and self.cf113 == __o.cf113 and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC63(D24, NamedTuple('DC63', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC63({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC63) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC67(False, int(0), _dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D25_DC67)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D25_DC66)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D25_DC68)

class D25_DC67(D25, NamedTuple('DC67', [('cf116', Any), ('cf117', Any), ('cf118', Any), ('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC67({_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC67) and self.cf116 == __o.cf116 and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC66(D25, NamedTuple('DC66', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC66({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC66) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC68(D25, NamedTuple('DC68', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC68({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC68) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC70(int(0), int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D26_DC70)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D26_DC71)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D26_DC69)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D26_DC72)

class D26_DC70(D26, NamedTuple('DC70', [('cf122', Any), ('cf123', Any), ('cf124', Any), ('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC70({_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)}, {_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC70) and self.cf122 == __o.cf122 and self.cf123 == __o.cf123 and self.cf124 == __o.cf124 and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC71(D26, NamedTuple('DC71', [('cf126', Any), ('cf127', Any), ('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC71({_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC71) and self.cf126 == __o.cf126 and self.cf127 == __o.cf127 and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC69(D26, NamedTuple('DC69', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC69({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC69) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC72(D26, NamedTuple('DC72', [('cf129', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC72({_dafny.string_of(self.cf129)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC72) and self.cf129 == __o.cf129
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D27_DC73)

class D27_DC73(D27, NamedTuple('DC73', [('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC73({_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC73) and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC75()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D28_DC75)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D28_DC74)

class D28_DC75(D28, NamedTuple('DC75', [])):
    def __dafnystr__(self) -> str:
        return f'D28.DC75'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC75)
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC74(D28, NamedTuple('DC74', [('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC74({_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC74) and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC77(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D29_DC77)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D29_DC78)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D29_DC76)

class D29_DC77(D29, NamedTuple('DC77', [('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC77({_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC77) and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC78(D29, NamedTuple('DC78', [('cf134', Any), ('cf135', Any), ('cf136', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC78({_dafny.string_of(self.cf134)}, {_dafny.string_of(self.cf135)}, {_dafny.string_of(self.cf136)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC78) and self.cf134 == __o.cf134 and self.cf135 == __o.cf135 and self.cf136 == __o.cf136
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC76(D29, NamedTuple('DC76', [('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC76({_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC76) and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC80(False, _dafny.Set({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D30_DC80)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D30_DC81)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D30_DC79)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D30_DC82)

class D30_DC80(D30, NamedTuple('DC80', [('cf138', Any), ('cf139', Any), ('cf140', Any), ('cf141', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC80({_dafny.string_of(self.cf138)}, {_dafny.string_of(self.cf139)}, {_dafny.string_of(self.cf140)}, {_dafny.string_of(self.cf141)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC80) and self.cf138 == __o.cf138 and self.cf139 == __o.cf139 and self.cf140 == __o.cf140 and self.cf141 == __o.cf141
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC81(D30, NamedTuple('DC81', [])):
    def __dafnystr__(self) -> str:
        return f'D30.DC81'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC81)
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC79(D30, NamedTuple('DC79', [('cf137', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC79({_dafny.string_of(self.cf137)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC79) and self.cf137 == __o.cf137
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC82(D30, NamedTuple('DC82', [('cf142', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC82({_dafny.string_of(self.cf142)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC82) and self.cf142 == __o.cf142
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: D31_DC84(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D31_DC84)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D31_DC83)

class D31_DC84(D31, NamedTuple('DC84', [('cf144', Any), ('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC84({_dafny.string_of(self.cf144)}, {_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC84) and self.cf144 == __o.cf144 and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()

class D31_DC83(D31, NamedTuple('DC83', [('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC83({_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC83) and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T1:
    pass
    def m2(self, p0, p1, p2, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f2: bool = False
        self.f3: int = int(0)
        self.f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f10: _dafny.Seq = _dafny.Seq({})
        self.f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f13: _dafny.Array = _dafny.Array(None, 0)
        self.f15: int = int(0)
        self.f16: _dafny.MultiSet = _dafny.MultiSet({})
        self.f17: _dafny.Map = _dafny.Map({})
        self._f1: bool = False
        self._f4: int = int(0)
        self._f5: bool = False
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: int = int(0)
        self._f11: bool = False
        self._f14: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self).f17 = f17

    @property
    def f1(self):
        return self._f1
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
    def f11(self):
        return self._f11
    @property
    def f14(self):
        return self._f14

class C0:
    def  __init__(self):
        self._f31: int = int(0)
        self._f32: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f31, f32):
        (self)._f31 = f31
        (self)._f32 = f32

    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32

class C1(T0):
    def  __init__(self):
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f18):
        (self)._f18 = f18

    def fm21(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsarcis")) if (self).f18 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_425_i0_ in range(default__.abs(-683))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjfcrawcl"))))

    def fm22(self, p0, p1, globalState):
        return default__.safeModuloInt(918, (0) - (default__.safeDivisionInt(-951, len(_dafny.Map({(self).f18: 618})))))

    def m1(self, p0, p1, p2, globalState):
        d_426_i0_: int
        d_426_i0_ = 0
        with _dafny.label("2"):
            while not ((self).f18) or (not((self).f18)):
                with _dafny.c_label("2"):
                    if (d_426_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_426_i0_ = (d_426_i0_) + (1)
                    (globalState).f2 = (p1) >= (default__.safeModuloInt(916, p1))
                    (globalState).f2 = (164) != (p1)
                    d_427_v0_: _dafny.Map
                    d_427_v0_ = _dafny.Map({(self).f18: (self).f18})
                    d_428_v1_: _dafny.Map
                    d_428_v1_ = _dafny.Map({len(d_427_v0_): (self).f18})
                    d_429_v2_: _dafny.Seq
                    d_429_v2_ = _dafny.SeqWithoutIsStrInference([(d_428_v1_).set(p1, (self).f18)])
                    d_430_v3_: D0
                    d_430_v3_ = D0_DC1((self).f18, p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_431_i1_ in range(default__.abs(856))]), (self).f18, p1)
                    d_432_v4_: _dafny.Map
                    d_432_v4_ = _dafny.Map({(p1) >= (len(d_429_v2_)): (p1) < ((self).fm22(p2, d_430_v3_, globalState))})
                    d_427_v0_ = (d_427_v0_).set(True, (self).f18)
                    d_433_v5_: _dafny.Array
                    def lambda9_(d_434_i2_):
                        return default__.safeDivisionInt(d_434_i2_, len(_dafny.Set({(self).f18, (self).f18})))

                    init5_ = lambda9_
                    nw61_ = _dafny.Array(None, 18)
                    for i0_5_ in range(nw61_.length(0)):
                        nw61_[i0_5_] = init5_(i0_5_)
                    d_433_v5_ = nw61_
                    d_435_v6_: _dafny.Seq
                    d_435_v6_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                    d_436_v7_: _dafny.Map
                    d_436_v7_ = _dafny.Map({(d_435_v6_)[default__.safeIndex(p1, len(d_435_v6_))]: d_433_v5_})
                    d_437_v8_: _dafny.Array
                    nw62_ = _dafny.Array(None, 9)
                    nw62_[int(0)] = d_433_v5_
                    nw62_[int(1)] = ((d_436_v7_)[(self).f18] if ((self).f18) in (d_436_v7_) else d_433_v5_)
                    nw62_[int(2)] = d_433_v5_
                    nw62_[int(3)] = d_433_v5_
                    nw62_[int(4)] = d_433_v5_
                    nw62_[int(5)] = d_433_v5_
                    nw62_[int(6)] = d_433_v5_
                    nw62_[int(7)] = d_433_v5_
                    nw62_[int(8)] = d_433_v5_
                    d_437_v8_ = nw62_
                    index44_ = default__.safeIndex(893, (d_437_v8_).length(0))
                    (d_437_v8_)[index44_] = d_433_v5_
                    d_438_v9_: _dafny.Array
                    nw63_ = _dafny.Array(False, 11)
                    d_438_v9_ = nw63_
                    index45_ = default__.safeIndex(840, (d_438_v9_).length(0))
                    (d_438_v9_)[index45_] = (self).f18
                    d_439_v10_: _dafny.Set
                    d_439_v10_ = _dafny.Set({p1})
                    index46_ = default__.safeIndex(893, (d_437_v8_).length(0))
                    index47_ = default__.safeIndex(840, (d_438_v9_).length(0))
                    rhs45_ = p1
                    rhs46_ = d_433_v5_
                    rhs47_ = default__.fm23(d_439_v10_, (default__.fm1(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_440_i3_ in range(default__.abs(188))])), (self).f18, (self).f18, globalState)) * (p1), globalState)
                    lhs29_ = globalState
                    lhs30_ = d_437_v8_
                    lhs31_ = default__.safeIndex(893, (d_437_v8_).length(0))
                    lhs32_ = d_438_v9_
                    lhs33_ = default__.safeIndex(840, (d_438_v9_).length(0))
                    lhs29_.f0 = rhs45_
                    lhs30_[lhs31_] = rhs46_
                    lhs32_[lhs33_] = rhs47_
                    pass
            pass
        d_441_v11_: _dafny.Map
        d_441_v11_ = _dafny.Map({(self).f18: p1})
        d_441_v11_ = (d_441_v11_) | (d_441_v11_)
        d_442_v12_: D1
        d_442_v12_ = D1_DC4(((self).f18) == ((self).f18))
        d_443_v13_: _dafny.Seq
        d_443_v13_ = _dafny.SeqWithoutIsStrInference([p1])
        d_444_v14_: str
        d_444_v14_ = _dafny.CodePoint('o')
        d_445_v15_: _dafny.Set
        d_445_v15_ = _dafny.Set({d_444_v14_})
        d_446_v16_: D2
        d_446_v16_ = D2_DC7((0) - (len(p2)), (self).f18, p1)
        d_442_v12_ = default__.fm24((len(d_443_v13_)) in (default__.fm25(p1, globalState)), (d_445_v15_) | (d_445_v15_), ((self).f18 if not((d_446_v16_).cf16) else (self).f18), globalState)
        d_447_v17_: D2
        d_447_v17_ = D2_DC8((self).f18, (self).f18, (self).f18)
        d_448_v18_: _dafny.Array
        nw64_ = _dafny.Array(_dafny.Seq({}), 4)
        d_448_v18_ = nw64_
        d_449_v19_: _dafny.Seq
        d_449_v19_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, not((self).f18), (self).f18, (self).f18])
        index48_ = default__.safeIndex(568, (d_448_v18_).length(0))
        (d_448_v18_)[index48_] = (d_449_v19_).set(default__.safeIndex(p1, len(d_449_v19_)), (self).f18)
        d_450_v20_: _dafny.Array
        def lambda10_(d_451_v17_):
            def lambda11_(d_452_i4_):
                return (d_451_v17_).cf19

            return lambda11_

        init6_ = lambda10_(d_447_v17_)
        nw65_ = _dafny.Array(None, 8)
        for i0_6_ in range(nw65_.length(0)):
            nw65_[i0_6_] = init6_(i0_6_)
        d_450_v20_ = nw65_
        index49_ = default__.safeIndex(910, (d_450_v20_).length(0))
        (d_450_v20_)[index49_] = ((self).f18) in (_dafny.Map({(self).f18: d_449_v19_}))
        index50_ = default__.safeIndex(568, (d_448_v18_).length(0))
        index51_ = default__.safeIndex(910, (d_450_v20_).length(0))
        rhs48_ = default__.fm26(d_443_v13_, (p1) * (p1), globalState)
        rhs49_ = d_449_v19_
        rhs50_ = (self).f18
        rhs51_ = (p1) == ((default__.fm1(p1, (self).f18, (self).f18, globalState)) - ((0) - (default__.fm1(len(p2), (self).f18, (self).f18, globalState))))
        lhs34_ = d_448_v18_
        lhs35_ = default__.safeIndex(568, (d_448_v18_).length(0))
        lhs36_ = d_450_v20_
        lhs37_ = default__.safeIndex(910, (d_450_v20_).length(0))
        lhs38_ = globalState
        d_447_v17_ = rhs48_
        lhs34_[lhs35_] = rhs49_
        lhs36_[lhs37_] = rhs50_
        lhs38_.f2 = rhs51_
        index52_ = default__.safeIndex(910, (d_450_v20_).length(0))
        (d_450_v20_)[index52_] = (self).f18
        if (d_446_v16_).cf16:
            d_453_v21_: D0
            d_453_v21_ = D0_DC1((self).f18, p1, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')]) if (d_450_v20_)[default__.safeIndex(910, (d_450_v20_).length(0))] else _dafny.SeqWithoutIsStrInference([d_444_v14_, d_444_v14_, d_444_v14_, d_444_v14_])), (self).f18, p1)
            pat_let_tv1_ = d_450_v20_
            pat_let_tv2_ = d_450_v20_
            pat_let_tv3_ = p1
            pat_let_tv4_ = p1
            index53_ = default__.safeIndex(910, (d_450_v20_).length(0))
            def iife35_(_pat_let6_0):
                def iife36_(d_454_dt__update__tmp_h0_):
                    def iife37_(_pat_let7_0):
                        def iife38_(d_455_dt__update_hcf16_h0_):
                            def iife39_(_pat_let8_0):
                                def iife40_(d_456_dt__update_hcf15_h0_):
                                    return D2_DC7(d_456_dt__update_hcf15_h0_, d_455_dt__update_hcf16_h0_, (d_454_dt__update__tmp_h0_).cf17)
                                return iife40_(_pat_let8_0)
                            return iife39_(default__.safeDivisionInt(pat_let_tv3_, pat_let_tv4_))
                        return iife38_(_pat_let7_0)
                    return iife37_((pat_let_tv2_)[default__.safeIndex(910, (pat_let_tv1_).length(0))])
                return iife36_(_pat_let6_0)
            rhs52_ = d_453_v21_
            rhs53_ = iife35_(D2_DC7(p1, (d_450_v20_)[default__.safeIndex(910, (d_450_v20_).length(0))], p1))
            rhs54_ = True
            lhs39_ = d_450_v20_
            lhs40_ = default__.safeIndex(910, (d_450_v20_).length(0))
            d_453_v21_ = rhs52_
            d_446_v16_ = rhs53_
            lhs39_[lhs40_] = rhs54_
            d_457_v22_: C0
            nw66_ = C0()
            nw66_.ctor__(p1, p1)
            d_457_v22_ = nw66_
            index54_ = default__.safeIndex(910, (d_450_v20_).length(0))
            (d_450_v20_)[index54_] = (50) >= ((p1) - (p1))
            d_458_v23_: C0
            nw67_ = C0()
            nw67_.ctor__((d_457_v22_).f31, (d_457_v22_).f31)
            d_458_v23_ = nw67_
            (globalState).f0 = (d_453_v21_).cf5
        elif True:
            d_459_v24_: _dafny.MultiSet
            d_459_v24_ = _dafny.MultiSet([p1])
            (globalState).f2 = (d_459_v24_).isdisjoint(d_459_v24_)
            d_460_v25_: int
            d_461_v26_: int
            d_462_v27_: _dafny.Array
            d_463_v28_: _dafny.Map
            out28_: int
            out29_: int
            out30_: _dafny.Array
            out31_: _dafny.Map
            out28_, out29_, out30_, out31_ = default__.m0(globalState)
            d_460_v25_ = out28_
            d_461_v26_ = out29_
            d_462_v27_ = out30_
            d_463_v28_ = out31_
            d_464_v29_: _dafny.Map
            d_464_v29_ = _dafny.Map({(self).f18: d_445_v15_})
            (globalState).f12 = (self).fm21(d_461_v26_, p2, d_464_v29_, d_449_v19_, globalState)
            (globalState).f0 = d_460_v25_
            index55_ = default__.safeIndex(301, (d_462_v27_).length(0))
            (d_462_v27_)[index55_] = len((p0).set(default__.safeIndex(len(p2), len(p0)), d_444_v14_))
            d_465_v30_: _dafny.MultiSet
            d_465_v30_ = _dafny.MultiSet([(d_450_v20_)[default__.safeIndex(910, (d_450_v20_).length(0))]])
            d_466_v31_: D1
            d_466_v31_ = D1_DC3(d_460_v25_, d_461_v26_, d_460_v25_, ((d_465_v30_)[(self).f18] if ((self).f18) in (d_465_v30_) else d_461_v26_), d_461_v26_)
            d_467_v32_: D1
            d_467_v32_ = D1_DC5(D1_DC5(d_466_v31_))
            d_468_v33_: _dafny.Map
            d_468_v33_ = _dafny.Map({(d_450_v20_)[default__.safeIndex(910, (d_450_v20_).length(0))]: (d_441_v11_).set((self).f18, d_460_v25_)})
            d_469_v34_: _dafny.Seq
            d_469_v34_ = _dafny.SeqWithoutIsStrInference([d_441_v11_, d_441_v11_])
            d_470_v35_: D0
            d_470_v35_ = D0_DC1(not((d_450_v20_)[default__.safeIndex(910, (d_450_v20_).length(0))]), len(((d_468_v33_)[True] if (True) in (d_468_v33_) else (d_469_v34_)[default__.safeIndex(((d_465_v30_)[(self).f18] if ((self).f18) in (d_465_v30_) else d_460_v25_), len(d_469_v34_))])), (self).fm21(len(p2), p2, default__.fm27(globalState), d_449_v19_, globalState), False, len(d_443_v13_))
            index56_ = default__.safeIndex(910, (d_450_v20_).length(0))
            index57_ = default__.safeIndex(301, (d_462_v27_).length(0))
            rhs55_ = d_461_v26_
            rhs56_ = d_460_v25_
            rhs57_ = False
            rhs58_ = (self).fm22((p2) + (p0), d_470_v35_, globalState)
            rhs59_ = d_467_v32_
            lhs41_ = globalState
            lhs42_ = d_450_v20_
            lhs43_ = default__.safeIndex(910, (d_450_v20_).length(0))
            lhs44_ = d_462_v27_
            lhs45_ = default__.safeIndex(301, (d_462_v27_).length(0))
            d_461_v26_ = rhs55_
            lhs41_.f0 = rhs56_
            lhs42_[lhs43_] = rhs57_
            lhs44_[lhs45_] = rhs58_
            d_467_v32_ = rhs59_


class C2(T0):
    def  __init__(self):
        self._f18: bool = False
        self.f33: int = int(0)
        self._f34: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f33, f34, f18):
        (self).f33 = f33
        (self)._f34 = f34
        (self)._f18 = f18

    def fm19(self, p0, globalState):
        return (self).f18

    def fm20(self, p0, p1, p2, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qa")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx"))})).intersection((_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_471_i0_ in range(default__.abs(401))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpnvrbl"))})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lftn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogsxkngw"))})))

    def m1(self, p0, p1, p2, globalState):
        hi3_ = default__.fm1(self.f33, (self).f18, (self).f18, globalState)
        for d_472_i0_ in range(p1, hi3_):
            d_473_v0_: _dafny.MultiSet
            d_473_v0_ = _dafny.MultiSet([(self).f18])
            (globalState).f2 = not(((self).f18) in ((d_473_v0_ if (self).f18 else d_473_v0_)))
            d_474_v1_: _dafny.Map
            d_474_v1_ = _dafny.Map({False: p1})
            d_475_v2_: str
            d_475_v2_ = _dafny.CodePoint('d')
            d_474_v1_ = _dafny.Map({(self).f18: len((p2) + ((p2).set(default__.safeIndex(d_472_i0_, len(p2)), d_475_v2_)))})
            d_476_v3_: D1
            d_476_v3_ = D1_DC4((self).f18)
            def iife41_(_pat_let9_0):
                def iife42_(d_477_dt__update__tmp_h0_):
                    def iife43_(_pat_let10_0):
                        def iife44_(d_478_dt__update_hcf12_h0_):
                            return D1_DC4(d_478_dt__update_hcf12_h0_)
                        return iife44_(_pat_let10_0)
                    return iife43_((self).f18)
                return iife42_(_pat_let9_0)
            source4_ = iife41_(d_476_v3_)
            if source4_.is_DC3:
                d_479___mcc_h0_ = source4_.cf7
                d_480___mcc_h1_ = source4_.cf8
                d_481___mcc_h2_ = source4_.cf9
                d_482___mcc_h3_ = source4_.cf10
                d_483___mcc_h4_ = source4_.cf11
                d_484_cf11_ = d_483___mcc_h4_
                d_485_cf10_ = d_482___mcc_h3_
                d_486_cf9_ = d_481___mcc_h2_
                d_487_cf8_ = d_480___mcc_h1_
                d_488_cf7_ = d_479___mcc_h0_
                d_489_v4_: _dafny.Seq
                d_489_v4_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
                d_490_v5_: D2
                d_490_v5_ = D2_DC7((d_473_v0_).cardinality, (self).f18, 601)
                d_491_v6_: _dafny.Array
                def lambda12_(d_492_cf9_):
                    def lambda13_(d_493_i1_):
                        return (d_493_i1_) - (d_492_cf9_)

                    return lambda13_

                init7_ = lambda12_(d_486_cf9_)
                nw68_ = _dafny.Array(None, 8)
                for i0_7_ in range(nw68_.length(0)):
                    nw68_[i0_7_] = init7_(i0_7_)
                d_491_v6_ = nw68_
                d_494_v7_: _dafny.Set
                d_494_v7_ = _dafny.Set({d_491_v6_, d_491_v6_})
                d_495_v8_: _dafny.Array
                nw69_ = _dafny.Array(None, 10)
                nw69_[int(0)] = (self).f18
                nw69_[int(1)] = (d_489_v4_)[default__.safeIndex(-876, len(d_489_v4_))]
                nw69_[int(2)] = (self).f18
                nw69_[int(3)] = (self).fm19(d_490_v5_, globalState)
                nw69_[int(4)] = (d_486_cf9_) > (len(p2))
                nw69_[int(5)] = (self).f18
                nw69_[int(6)] = (self).f18
                nw69_[int(7)] = not((d_494_v7_).isdisjoint(d_494_v7_))
                nw69_[int(8)] = not((self).f18)
                nw69_[int(9)] = (self).f18
                d_495_v8_ = nw69_
                index58_ = default__.safeIndex(104, (d_495_v8_).length(0))
                (d_495_v8_)[index58_] = (self).f18
                d_496_v9_: _dafny.Map
                d_496_v9_ = _dafny.Map({(self).f18: d_474_v1_})
                d_497_v10_: _dafny.Seq
                d_497_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f18: ((d_473_v0_)[(self).f18] if ((self).f18) in (d_473_v0_) else d_487_cf8_)}), d_474_v1_, d_474_v1_, d_474_v1_])
                d_498_v11_: D0
                d_498_v11_ = D0_DC1((self).f18, d_485_cf10_, p0, not((self).f18), p1)
                d_499_v12_: _dafny.MultiSet
                d_499_v12_ = _dafny.MultiSet([p0, (d_498_v11_).cf3])
                index59_ = default__.safeIndex(104, (d_495_v8_).length(0))
                rhs60_ = (p0) <= ((p2) + (p0))
                rhs61_ = _dafny.Map({(self).f18: (d_497_v10_)[default__.safeIndex((d_499_v12_).cardinality, len(d_497_v10_))]})
                lhs46_ = d_495_v8_
                lhs47_ = default__.safeIndex(104, (d_495_v8_).length(0))
                lhs46_[lhs47_] = rhs60_
                d_496_v9_ = rhs61_
                d_488_cf7_ = d_487_cf8_
                d_500_v13_: _dafny.Set
                d_500_v13_ = _dafny.Set({((p0)[default__.safeIndex(732, len(p0))] if (self).f18 else d_475_v2_), d_475_v2_, d_475_v2_, d_475_v2_, d_475_v2_})
                d_500_v13_ = (_dafny.Set({d_475_v2_})) | ((d_500_v13_).intersection(d_500_v13_))
                index60_ = default__.safeIndex(451, (d_491_v6_).length(0))
                (d_491_v6_)[index60_] = p1
                index61_ = default__.safeIndex(451, (d_491_v6_).length(0))
                (d_491_v6_)[index61_] = default__.fm1((d_473_v0_).cardinality, (d_495_v8_)[default__.safeIndex(104, (d_495_v8_).length(0))], (self).f18, globalState)
            elif source4_.is_DC4:
                d_501___mcc_h5_ = source4_.cf12
                d_502_cf12_ = d_501___mcc_h5_
                d_503_v14_: _dafny.Seq
                d_503_v14_ = _dafny.SeqWithoutIsStrInference([(self).f18, d_502_cf12_, (self).f18, (self).f18])
                d_503_v14_ = d_503_v14_
                (globalState).f3 = 272
                d_504_v15_: _dafny.Array
                nw70_ = _dafny.Array(False, 6)
                d_504_v15_ = nw70_
                index62_ = default__.safeIndex(492, (d_504_v15_).length(0))
                (d_504_v15_)[index62_] = not(d_502_cf12_)
                d_505_v16_: _dafny.Map
                d_505_v16_ = _dafny.Map({(self).f18: d_502_cf12_})
                index63_ = default__.safeIndex(492, (d_504_v15_).length(0))
                (d_504_v15_)[index63_] = ((d_505_v16_)[True] if (True) in (d_505_v16_) else d_502_cf12_)
                d_506_v17_: T0
                nw71_ = C1()
                nw71_.ctor__((d_504_v15_)[default__.safeIndex(492, (d_504_v15_).length(0))])
                d_506_v17_ = nw71_
                nw72_ = C1()
                nw72_.ctor__((d_506_v17_).f18)
                d_506_v17_ = nw72_
            elif source4_.is_DC2:
                d_507___mcc_h6_ = source4_.cf6
                d_508_cf6_ = d_507___mcc_h6_
                d_509_v19_: _dafny.Map
                d_509_v19_ = _dafny.Map({d_472_i0_: -697})
                d_510_v20_: _dafny.Map
                d_510_v20_ = _dafny.Map({len(d_509_v19_): len(p2)})
                d_511_v21_: _dafny.Seq
                d_511_v21_ = _dafny.SeqWithoutIsStrInference([d_510_v20_])
                d_512_v22_: _dafny.MultiSet
                d_512_v22_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p1 for d_513_i2_ in range(default__.abs(779))])])
                def iife45_():
                    coll23_ = _dafny.Map()
                    compr_23_: int
                    for compr_23_ in ((d_510_v20_) | ((d_511_v21_)[default__.safeIndex(((d_512_v22_)[_dafny.SeqWithoutIsStrInference([d_472_i0_, d_472_i0_])] if (_dafny.SeqWithoutIsStrInference([d_472_i0_, d_472_i0_])) in (d_512_v22_) else d_472_i0_), len(d_511_v21_))])).keys.Elements:
                        d_514_v18_: int = compr_23_
                        if (d_514_v18_) in ((d_510_v20_) | ((d_511_v21_)[default__.safeIndex(((d_512_v22_)[_dafny.SeqWithoutIsStrInference([d_472_i0_, d_472_i0_])] if (_dafny.SeqWithoutIsStrInference([d_472_i0_, d_472_i0_])) in (d_512_v22_) else d_472_i0_), len(d_511_v21_))])):
                            coll23_[default__.safeDivisionInt(d_514_v18_, -808)] = (self).f18
                    return _dafny.Map(coll23_)
                (globalState).f3 = len(iife45_()
                )
                (globalState).f12 = p0
                d_515_v23_: _dafny.Array
                nw73_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_515_v23_ = nw73_
                d_516_v24_: _dafny.Array
                nw74_ = _dafny.Array(False, 18)
                d_516_v24_ = nw74_
                index64_ = default__.safeIndex(927, (d_515_v23_).length(0))
                (d_515_v23_)[index64_] = d_516_v24_
                d_517_v25_: _dafny.Array
                nw75_ = _dafny.Array(_dafny.Seq({}), 11)
                d_517_v25_ = nw75_
                d_518_v26_: _dafny.Set
                d_518_v26_ = _dafny.Set({(self).f18})
                d_519_v27_: _dafny.MultiSet
                d_519_v27_ = _dafny.MultiSet([len((self).f34)])
                index65_ = default__.safeIndex(927, (d_515_v23_).length(0))
                rhs62_ = default__.safeModuloInt(len(d_518_v26_), (((d_474_v1_)[(self).f18] if ((self).f18) in (d_474_v1_) else p1)) - ((d_519_v27_).cardinality))
                rhs63_ = d_516_v24_
                rhs64_ = d_517_v25_
                lhs48_ = globalState
                lhs49_ = d_515_v23_
                lhs50_ = default__.safeIndex(927, (d_515_v23_).length(0))
                lhs48_.f0 = rhs62_
                lhs49_[lhs50_] = rhs63_
                d_517_v25_ = rhs64_
                d_520_v28_: _dafny.Map
                d_520_v28_ = _dafny.Map({d_472_i0_: (_dafny.MultiSet([p1, 731])) - (_dafny.MultiSet([d_472_i0_]))})
                d_520_v28_ = (d_520_v28_).set((0) - (d_472_i0_), d_519_v27_)
            elif True:
                d_521___mcc_h7_ = source4_.cf13
                d_522_cf13_ = d_521___mcc_h7_
                d_523_v29_: D1
                d_523_v29_ = D1_DC3(p1, p1, self.f33, 87, d_472_i0_)
                d_524_v30_: _dafny.Map
                d_524_v30_ = _dafny.Map({((d_473_v0_)[(self).f18] if ((self).f18) in (d_473_v0_) else d_472_i0_): 841})
                d_525_v31_: _dafny.Array
                nw76_ = _dafny.Array(None, 3)
                nw76_[int(0)] = self.f33
                nw76_[int(1)] = 783
                nw76_[int(2)] = p1
                d_525_v31_ = nw76_
                d_526_v32_: _dafny.Set
                d_526_v32_ = _dafny.Set({d_525_v31_})
                pat_let_tv5_ = p1
                pat_let_tv6_ = d_526_v32_
                d_527_v33_: _dafny.Array
                nw77_ = _dafny.Array(None, 26)
                nw77_[int(0)] = d_523_v29_
                nw77_[int(1)] = d_523_v29_
                nw77_[int(2)] = d_523_v29_
                nw77_[int(3)] = d_523_v29_
                nw77_[int(4)] = d_523_v29_
                def iife46_(_pat_let11_0):
                    def iife47_(d_528_dt__update__tmp_h1_):
                        def iife48_(_pat_let12_0):
                            def iife49_(d_529_dt__update_hcf9_h0_):
                                return D1_DC3((d_528_dt__update__tmp_h1_).cf7, (d_528_dt__update__tmp_h1_).cf8, d_529_dt__update_hcf9_h0_, (d_528_dt__update__tmp_h1_).cf10, (d_528_dt__update__tmp_h1_).cf11)
                            return iife49_(_pat_let12_0)
                        return iife48_(939)
                    return iife47_(_pat_let11_0)
                nw77_[int(5)] = iife46_(d_523_v29_)
                nw77_[int(6)] = d_523_v29_
                def iife50_(_pat_let13_0):
                    def iife51_(d_530_dt__update__tmp_h2_):
                        def iife52_(_pat_let14_0):
                            def iife53_(d_531_dt__update_hcf10_h0_):
                                def iife54_(_pat_let15_0):
                                    def iife55_(d_532_dt__update_hcf7_h0_):
                                        def iife56_(_pat_let16_0):
                                            def iife57_(d_533_dt__update_hcf8_h0_):
                                                return D1_DC3(d_532_dt__update_hcf7_h0_, d_533_dt__update_hcf8_h0_, (d_530_dt__update__tmp_h2_).cf9, d_531_dt__update_hcf10_h0_, (d_530_dt__update__tmp_h2_).cf11)
                                            return iife57_(_pat_let16_0)
                                        return iife56_(self.f33)
                                    return iife55_(_pat_let15_0)
                                return iife54_(pat_let_tv5_)
                            return iife53_(_pat_let14_0)
                        return iife52_(self.f33)
                    return iife51_(_pat_let13_0)
                nw77_[int(7)] = iife50_(D1_DC3(d_472_i0_, p1, self.f33, self.f33, len((self).f34)))
                nw77_[int(8)] = D1_DC3((0) - (p1), len(d_524_v30_), p1, self.f33, self.f33)
                nw77_[int(9)] = d_523_v29_
                nw77_[int(10)] = d_523_v29_
                nw77_[int(11)] = d_523_v29_
                nw77_[int(12)] = d_523_v29_
                nw77_[int(13)] = d_523_v29_
                nw77_[int(14)] = d_523_v29_
                nw77_[int(15)] = d_523_v29_
                nw77_[int(16)] = d_523_v29_
                nw77_[int(17)] = D1_DC3((0) - (self.f33), len(_dafny.SeqWithoutIsStrInference([(self).f18])), p1, 484, 572)
                nw77_[int(18)] = d_523_v29_
                def iife58_(_pat_let17_0):
                    def iife59_(d_534_dt__update__tmp_h3_):
                        def iife60_(_pat_let18_0):
                            def iife61_(d_535_dt__update_hcf10_h1_):
                                return D1_DC3((d_534_dt__update__tmp_h3_).cf7, (d_534_dt__update__tmp_h3_).cf8, (d_534_dt__update__tmp_h3_).cf9, d_535_dt__update_hcf10_h1_, (d_534_dt__update__tmp_h3_).cf11)
                            return iife61_(_pat_let18_0)
                        return iife60_(len(pat_let_tv6_))
                    return iife59_(_pat_let17_0)
                nw77_[int(19)] = iife58_(d_523_v29_)
                nw77_[int(20)] = d_523_v29_
                nw77_[int(21)] = d_523_v29_
                nw77_[int(22)] = d_523_v29_
                nw77_[int(23)] = d_523_v29_
                nw77_[int(24)] = d_523_v29_
                nw77_[int(25)] = D1_DC3(727, self.f33, -359, default__.fm1(self.f33, (self).f18, (self).f18, globalState), self.f33)
                d_527_v33_ = nw77_
                index66_ = default__.safeIndex(430, (d_527_v33_).length(0))
                (d_527_v33_)[index66_] = d_523_v29_
                d_536_v34_: D0
                d_536_v34_ = D0_DC0((self).f18)
                index67_ = default__.safeIndex(430, (d_527_v33_).length(0))
                def iife62_(_pat_let19_0):
                    def iife63_(d_537_dt__update__tmp_h4_):
                        def iife64_(_pat_let20_0):
                            def iife65_(d_538_dt__update_hcf0_h0_):
                                return D0_DC0(d_538_dt__update_hcf0_h0_)
                            return iife65_(_pat_let20_0)
                        return iife64_((self).f18)
                    return iife63_(_pat_let19_0)
                (d_527_v33_)[index67_] = default__.fm28(iife62_(d_536_v34_), globalState)
                (globalState).f0 = (0) - (self.f33)
                (globalState).f3 = p1
                d_539_v35_: _dafny.Array
                nw78_ = _dafny.Array(False, 25)
                d_539_v35_ = nw78_
                d_540_v36_: _dafny.Map
                d_540_v36_ = _dafny.Map({(self).f18: d_539_v35_})
                d_540_v36_ = (d_540_v36_).set((self).f18, d_539_v35_)
            d_541_v37_: _dafny.Array
            def lambda14_(d_542_i3_):
                return (self).f18

            init8_ = lambda14_
            nw79_ = _dafny.Array(None, 19)
            for i0_8_ in range(nw79_.length(0)):
                nw79_[i0_8_] = init8_(i0_8_)
            d_541_v37_ = nw79_
            d_543_v38_: _dafny.Set
            d_543_v38_ = _dafny.Set({(self).f18})
            d_544_v39_: _dafny.Seq
            d_544_v39_ = _dafny.SeqWithoutIsStrInference([d_543_v38_])
            d_545_v40_: _dafny.MultiSet
            d_545_v40_ = _dafny.MultiSet([d_475_v2_])
            d_546_v41_: _dafny.Array
            nw80_ = _dafny.Array(None, 10)
            nw80_[int(0)] = p1
            nw80_[int(1)] = d_472_i0_
            nw80_[int(2)] = d_472_i0_
            nw80_[int(3)] = default__.safeDivisionInt((0) - (d_472_i0_), default__.fm1(len(p2), (self).f18, (self).f18, globalState))
            nw80_[int(4)] = (0) - (d_472_i0_)
            nw80_[int(5)] = (d_472_i0_) * (len(d_544_v39_))
            nw80_[int(6)] = self.f33
            nw80_[int(7)] = self.f33
            nw80_[int(8)] = d_472_i0_
            nw80_[int(9)] = default__.safeDivisionInt(d_472_i0_, (d_545_v40_).cardinality)
            d_546_v41_ = nw80_
            index68_ = default__.safeIndex(765, (d_546_v41_).length(0))
            (d_546_v41_)[index68_] = 79
            d_547_v42_: _dafny.Map
            d_547_v42_ = _dafny.Map({len((self).f34): (self).f18})
            d_548_v43_: _dafny.Seq
            d_548_v43_ = _dafny.SeqWithoutIsStrInference([d_547_v42_])
            d_549_v44_: _dafny.Map
            d_549_v44_ = _dafny.Map({d_548_v43_: p1})
            d_550_v45_: D2
            d_550_v45_ = D2_DC7(p1, (self).f18, d_472_i0_)
            index69_ = default__.safeIndex(765, (d_546_v41_).length(0))
            rhs65_ = d_541_v37_
            rhs66_ = (d_472_i0_) * (self.f33)
            rhs67_ = (len((((self).f34).set(default__.safeIndex(d_472_i0_, len((self).f34)), (d_473_v0_).cardinality)) + ((self).f34))) - (((d_549_v44_)[d_548_v43_] if (d_548_v43_) in (d_549_v44_) else d_472_i0_))
            rhs68_ = not ((self).fm19(d_550_v45_, globalState)) or ((self).f18)
            lhs51_ = d_546_v41_
            lhs52_ = default__.safeIndex(765, (d_546_v41_).length(0))
            lhs53_ = globalState
            lhs54_ = globalState
            d_541_v37_ = rhs65_
            lhs51_[lhs52_] = rhs66_
            lhs53_.f0 = rhs67_
            lhs54_.f2 = rhs68_
        d_551_v46_: _dafny.MultiSet
        d_551_v46_ = _dafny.MultiSet([not((self).f18), (self).f18])
        d_552_v47_: D2
        d_552_v47_ = D2_DC7(p1, (D0_DC1((self).f18, (0) - (self.f33), p2, (self).f18, self.f33)).cf1, (0) - (self.f33))
        d_553_v48_: _dafny.Map
        d_553_v48_ = _dafny.Map({self.f33: (d_552_v47_).cf15})
        d_554_v49_: _dafny.Map
        d_554_v49_ = _dafny.Map({_dafny.CodePoint('o'): (self).f18})
        d_555_v50_: _dafny.Seq
        d_555_v50_ = _dafny.SeqWithoutIsStrInference([True, True, (self).f18, (self).f18, False])
        d_556_v51_: _dafny.Map
        d_556_v51_ = _dafny.Map({d_555_v50_: True})
        d_557_v52_: _dafny.Map
        d_557_v52_ = _dafny.Map({len(d_556_v51_): (self).f18})
        d_558_v53_: _dafny.MultiSet
        d_558_v53_ = _dafny.MultiSet([p1])
        d_559_v54_: _dafny.Map
        d_559_v54_ = _dafny.Map({d_558_v53_: (self).f18})
        d_560_v55_: _dafny.Seq
        d_560_v55_ = _dafny.SeqWithoutIsStrInference([(self).f18, (d_555_v50_)[default__.safeIndex(p1, len(d_555_v50_))], ((d_557_v52_)[(0) - (len(d_559_v54_))] if ((0) - (len(d_559_v54_))) in (d_557_v52_) else not(True))])
        d_561_v56_: D1
        d_561_v56_ = D1_DC3((0) - (p1), self.f33, p1, self.f33, (0) - (len(d_555_v50_)))
        d_562_v57_: _dafny.Map
        d_562_v57_ = _dafny.Map({(self).f18: (d_558_v53_).cardinality})
        d_563_v58_: _dafny.Set
        d_563_v58_ = _dafny.Set({(d_551_v46_).cardinality})
        d_564_v59_: _dafny.Array
        nw81_ = _dafny.Array(None, 23)
        nw81_[int(0)] = self.f33
        nw81_[int(1)] = (len(default__.fm29((self).f18, self.f33, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "deqxsk")), globalState))) - (self.f33)
        nw81_[int(2)] = (0) - (((d_551_v46_)[(self).f18] if ((self).f18) in (d_551_v46_) else len(d_553_v48_)))
        nw81_[int(3)] = default__.fm1(self.f33, (self).f18, (d_552_v47_).cf16, globalState)
        nw81_[int(4)] = len(d_554_v49_)
        nw81_[int(5)] = len(d_555_v50_)
        nw81_[int(6)] = p1
        nw81_[int(7)] = default__.fm1(self.f33, (self).f18, (self).f18, globalState)
        def iife66_(_pat_let21_0):
            def iife67_(d_565_dt__update__tmp_h5_):
                def iife68_(_pat_let22_0):
                    def iife69_(d_566_dt__update_hcf10_h2_):
                        def iife70_(_pat_let23_0):
                            def iife71_(d_567_dt__update_hcf11_h0_):
                                return D1_DC3((d_565_dt__update__tmp_h5_).cf7, (d_565_dt__update__tmp_h5_).cf8, (d_565_dt__update__tmp_h5_).cf9, d_566_dt__update_hcf10_h2_, d_567_dt__update_hcf11_h0_)
                            return iife71_(_pat_let23_0)
                        return iife70_(-280)
                    return iife69_(_pat_let22_0)
                return iife68_(self.f33)
            return iife67_(_pat_let21_0)
        nw81_[int(8)] = (iife66_(d_561_v56_)).cf10
        nw81_[int(9)] = self.f33
        nw81_[int(10)] = self.f33
        nw81_[int(11)] = p1
        nw81_[int(12)] = self.f33
        nw81_[int(13)] = ((self).f34)[default__.safeIndex(len(d_562_v57_), len((self).f34))]
        nw81_[int(14)] = len((p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_568_i4_ in range(default__.abs(593))])))
        nw81_[int(15)] = self.f33
        nw81_[int(16)] = (self.f33) + (self.f33)
        nw81_[int(17)] = (0) - ((len(d_563_v58_)) - ((0) - (((d_551_v46_)[(self).f18] if ((self).f18) in (d_551_v46_) else len(p2)))))
        nw81_[int(18)] = self.f33
        nw81_[int(19)] = p1
        nw81_[int(20)] = p1
        nw81_[int(21)] = len((self).f34)
        nw81_[int(22)] = (self.f33 if (self).f18 else p1)
        d_564_v59_ = nw81_
        d_564_v59_ = d_564_v59_
        d_569_v60_: _dafny.Array
        nw82_ = _dafny.Array(_dafny.Map({}), 25)
        d_569_v60_ = nw82_
        d_570_v61_: _dafny.Seq
        d_570_v61_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pb"))) + (p0), p2])
        d_571_v62_: _dafny.Set
        d_571_v62_ = _dafny.Set({not((self).f18)})
        d_572_v63_: _dafny.Array
        nw83_ = _dafny.Array(None, 8)
        nw83_[int(0)] = (len(d_571_v62_)) > (self.f33)
        nw83_[int(1)] = (self).f18
        nw83_[int(2)] = (self).f18
        nw83_[int(3)] = (d_558_v53_).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_562_v57_)])))
        nw83_[int(4)] = (self).f18
        nw83_[int(5)] = (self).f18
        nw83_[int(6)] = (self).f18
        nw83_[int(7)] = (self).f18
        d_572_v63_ = nw83_
        index70_ = default__.safeIndex(408, (d_572_v63_).length(0))
        (d_572_v63_)[index70_] = ((d_557_v52_)[self.f33] if (self.f33) in (d_557_v52_) else True)
        d_573_v64_: D2
        d_573_v64_ = D2_DC8((self).f18, False, (self).f18)
        d_574_v65_: _dafny.Map
        d_574_v65_ = _dafny.Map({d_573_v64_: d_572_v63_})
        d_575_v66_: str
        d_575_v66_ = _dafny.CodePoint('n')
        d_576_v67_: _dafny.Map
        d_576_v67_ = _dafny.Map({d_575_v66_: 790})
        index71_ = default__.safeIndex(408, (d_572_v63_).length(0))
        def iife72_():
            coll24_ = _dafny.Set()
            compr_24_: int
            for compr_24_ in (d_553_v48_).keys.Elements:
                d_577_v68_: int = compr_24_
                if (d_577_v68_) in (d_553_v48_):
                    coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_577_v68_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))))]))
            return _dafny.Set(coll24_)
        rhs69_ = (D3_DC10(d_569_v60_)).cf22
        rhs70_ = (d_570_v61_) + (d_570_v61_)
        rhs71_ = ((self).f34) + ((_dafny.SeqWithoutIsStrInference([len((p0).set(default__.safeIndex(((d_576_v67_)[d_575_v66_] if (d_575_v66_) in (d_576_v67_) else p1), len(p0)), d_575_v66_))])) + ((self).f34))
        rhs72_ = not ((self).fm19(d_552_v47_, globalState)) or ((d_563_v58_).isdisjoint(iife72_()
        ))
        rhs73_ = d_574_v65_
        lhs55_ = globalState
        lhs56_ = d_572_v63_
        lhs57_ = default__.safeIndex(408, (d_572_v63_).length(0))
        d_569_v60_ = rhs69_
        d_570_v61_ = rhs70_
        lhs55_.f10 = rhs71_
        lhs56_[lhs57_] = rhs72_
        d_574_v65_ = rhs73_
        d_578_i5_: int
        d_578_i5_ = 0
        with _dafny.label("3"):
            while True:
                with _dafny.c_label("3"):
                    if (d_578_i5_) >= (100):
                        raise _dafny.Break("3")
                    d_578_i5_ = (d_578_i5_) + (1)
                    rhs74_ = len(p0)
                    rhs75_ = (p1) + (self.f33)
                    lhs58_ = globalState
                    lhs59_ = globalState
                    lhs58_.f15 = rhs74_
                    lhs59_.f3 = rhs75_
                    index72_ = default__.safeIndex(408, (d_572_v63_).length(0))
                    (d_572_v63_)[index72_] = (self.f33) < (self.f33)
                    if (self.f33) > ((len(p2)) * (-440)):
                        rhs76_ = d_575_v66_
                        rhs77_ = d_571_v62_
                        rhs78_ = 815
                        lhs60_ = globalState
                        d_575_v66_ = rhs76_
                        d_571_v62_ = rhs77_
                        lhs60_.f15 = rhs78_
                        d_579_v69_: _dafny.Set
                        d_579_v69_ = _dafny.Set({d_551_v46_, (d_551_v46_).set((self).f18, default__.abs(p1)), d_551_v46_})
                        d_580_v70_: C0
                        nw84_ = C0()
                        nw84_.ctor__(self.f33, len(d_579_v69_))
                        d_580_v70_ = nw84_
                        d_581_v71_: C0
                        nw85_ = C0()
                        nw85_.ctor__(len(d_555_v50_), self.f33)
                        d_581_v71_ = nw85_
                        d_582_v72_: _dafny.Map
                        d_582_v72_ = _dafny.Map({d_572_v63_: (d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))]})
                        d_582_v72_ = (d_582_v72_).set(d_572_v63_, (self).f18)
                        d_583_v73_: _dafny.Map
                        d_583_v73_ = _dafny.Map({(self).f18: p2})
                        d_584_v74_: _dafny.Map
                        d_584_v74_ = _dafny.Map({(self).f18: ((d_583_v73_)[False] if (False) in (d_583_v73_) else p2)})
                        d_583_v73_ = (d_583_v73_).set((self).f18, p0)
                    elif True:
                        d_585_v75_: C0
                        nw86_ = C0()
                        nw86_.ctor__(self.f33, (0) - (len(p0)))
                        d_585_v75_ = nw86_
                        d_586_v76_: _dafny.Map
                        d_586_v76_ = _dafny.Map({(d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))]: d_562_v57_})
                        d_587_v77_: _dafny.Map
                        d_587_v77_ = _dafny.Map({(d_586_v76_) | (d_586_v76_): (d_585_v75_).f32})
                        d_587_v77_ = (d_587_v77_).set(d_586_v76_, 870)
                        d_588_v78_: _dafny.Map
                        d_588_v78_ = _dafny.Map({p0: (d_585_v75_).f32})
                        (globalState).f3 = ((d_558_v53_)[(len(d_588_v78_) if (self).f18 else (d_585_v75_).f31)] if ((len(d_588_v78_) if (self).f18 else (d_585_v75_).f31)) in (d_558_v53_) else (d_585_v75_).f32)
                        d_589_v79_: _dafny.Map
                        d_589_v79_ = _dafny.Map({(self).f18: True})
                        d_589_v79_ = ((d_589_v79_) | (_dafny.Map({False: (d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))]}))) | ((d_589_v79_).set(True, True))
                        d_590_v80_: _dafny.Array
                        nw87_ = _dafny.Array(None, 11)
                        nw87_[int(0)] = d_561_v56_
                        nw87_[int(1)] = d_561_v56_
                        def iife73_(_pat_let24_0):
                            def iife74_(d_591_dt__update__tmp_h6_):
                                def iife75_(_pat_let25_0):
                                    def iife76_(d_592_dt__update_hcf10_h3_):
                                        return D1_DC3((d_591_dt__update__tmp_h6_).cf7, (d_591_dt__update__tmp_h6_).cf8, (d_591_dt__update__tmp_h6_).cf9, d_592_dt__update_hcf10_h3_, (d_591_dt__update__tmp_h6_).cf11)
                                    return iife76_(_pat_let25_0)
                                return iife75_(1)
                            return iife74_(_pat_let24_0)
                        nw87_[int(2)] = iife73_(d_561_v56_)
                        nw87_[int(3)] = d_561_v56_
                        nw87_[int(4)] = d_561_v56_
                        nw87_[int(5)] = d_561_v56_
                        nw87_[int(6)] = D1_DC3(len(d_557_v52_), (d_585_v75_).f31, p1, (0) - (len(d_553_v48_)), (d_585_v75_).f32)
                        nw87_[int(7)] = d_561_v56_
                        nw87_[int(8)] = d_561_v56_
                        nw87_[int(9)] = d_561_v56_
                        nw87_[int(10)] = d_561_v56_
                        d_590_v80_ = nw87_
                        d_593_v81_: _dafny.Map
                        d_593_v81_ = _dafny.Map({(self).f34: p1})
                        index73_ = default__.safeIndex(879, (d_590_v80_).length(0))
                        (d_590_v80_)[index73_] = (d_561_v56_ if (d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))] else D1_DC3((d_585_v75_).f32, len(d_593_v81_), self.f33, p1, p1))
                        index74_ = default__.safeIndex(879, (d_590_v80_).length(0))
                        (d_590_v80_)[index74_] = (d_561_v56_ if (self).f18 else d_561_v56_)
                    (globalState).f2 = (self).f18
                    pass
            pass
        hi4_ = self.f33
        for d_594_i6_ in range(default__.fm1(self.f33, (self).f18, (d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))], globalState), hi4_):
            (globalState).f3 = self.f33
            d_595_v82_: _dafny.Map
            d_595_v82_ = _dafny.Map({True: d_572_v63_})
            d_596_v83_: _dafny.Seq
            d_596_v83_ = _dafny.SeqWithoutIsStrInference([d_595_v82_, d_595_v82_, d_595_v82_])
            d_597_v84_: _dafny.Map
            d_597_v84_ = _dafny.Map({883: (d_595_v82_) | ((d_596_v83_)[default__.safeIndex(p1, len(d_596_v83_))])})
            d_597_v84_ = (d_597_v84_).set(d_594_i6_, _dafny.Map({not((d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))]): d_572_v63_}))
            d_598_v85_: C0
            nw88_ = C0()
            nw88_.ctor__((len(_dafny.SeqWithoutIsStrInference([d_575_v66_ for d_599_i7_ in range(default__.abs(-68))]))) * (p1), self.f33)
            d_598_v85_ = nw88_
            index75_ = default__.safeIndex(408, (d_572_v63_).length(0))
            (d_572_v63_)[index75_] = (d_572_v63_)[default__.safeIndex(408, (d_572_v63_).length(0))]
        (globalState).f2 = (self).f18

    @property
    def f34(self):
        return self._f34

class C3(T1, T0):
    def  __init__(self):
        self._f18: bool = False
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f20: _dafny.Seq = _dafny.Seq({})
        self._f35: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f35, f19, f20, f18):
        (self)._f35 = f35
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f18 = f18

    def fm2(self, p0, p1, p2, globalState):
        return not ((self).f18) or (((D3_DC11((self).f18, not((self).f18))).cf23) not in (_dafny.SeqWithoutIsStrInference([(self).f18, True])))

    def fm3(self, p0, globalState):
        if (_dafny.MultiSet([_dafny.CodePoint('g')])).issubset(_dafny.MultiSet([_dafny.CodePoint('o'), _dafny.CodePoint('i')])):
            return (self).f18
        elif True:
            return (self).f18

    def fm37(self, p0, globalState):
        return (self).f19

    def fm38(self, p0, globalState):
        return 814

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        (globalState).f2 = p2
        d_600_v0_: _dafny.MultiSet
        d_600_v0_ = _dafny.MultiSet([p1])
        if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1, p1, p2, p1]))) == (d_600_v0_):
            (globalState).f2 = False
            d_601_v1_: _dafny.Array
            nw89_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_601_v1_ = nw89_
            d_602_v2_: _dafny.Map
            d_602_v2_ = _dafny.Map({p2: p2})
            d_603_v3_: _dafny.Array
            nw90_ = _dafny.Array(None, 16)
            nw90_[int(0)] = ((d_602_v2_)[p2] if (p2) in (d_602_v2_) else (self).f18)
            nw90_[int(1)] = p1
            nw90_[int(2)] = p1
            nw90_[int(3)] = p2
            nw90_[int(4)] = p2
            nw90_[int(5)] = p1
            nw90_[int(6)] = (self).f18
            nw90_[int(7)] = p2
            nw90_[int(8)] = p2
            nw90_[int(9)] = (self).f18
            nw90_[int(10)] = ((d_602_v2_)[p1] if (p1) in (d_602_v2_) else p2)
            nw90_[int(11)] = p2
            nw90_[int(12)] = p1
            nw90_[int(13)] = p2
            nw90_[int(14)] = (self).f18
            nw90_[int(15)] = p1
            d_603_v3_ = nw90_
            index76_ = default__.safeIndex(126, (d_601_v1_).length(0))
            (d_601_v1_)[index76_] = (d_603_v3_ if False else d_603_v3_)
            index77_ = default__.safeIndex(126, (d_601_v1_).length(0))
            (d_601_v1_)[index77_] = d_603_v3_
            (globalState).f2 = (self).f18
            d_604_v4_: int
            d_604_v4_ = 640
            (globalState).f0 = (d_604_v4_) - (len((self).fm37(len((self).f19), globalState)))
            d_605_v5_: _dafny.Seq
            d_605_v5_ = _dafny.SeqWithoutIsStrInference([p1])
            d_606_v6_: _dafny.Map
            d_606_v6_ = _dafny.Map({d_604_v4_: (self).f18})
            d_607_v7_: _dafny.Array
            nw91_ = _dafny.Array(None, 19)
            nw91_[int(0)] = d_600_v0_
            nw91_[int(1)] = d_600_v0_
            nw91_[int(2)] = (_dafny.MultiSet([(self).f18])) - (d_600_v0_)
            nw91_[int(3)] = (d_600_v0_).set(p1, default__.abs(d_604_v4_))
            nw91_[int(4)] = d_600_v0_
            nw91_[int(5)] = d_600_v0_
            nw91_[int(6)] = d_600_v0_
            nw91_[int(7)] = d_600_v0_
            nw91_[int(8)] = d_600_v0_
            nw91_[int(9)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f18]))).intersection(d_600_v0_)
            nw91_[int(10)] = (d_600_v0_).intersection(d_600_v0_)
            nw91_[int(11)] = (d_600_v0_) - (d_600_v0_)
            nw91_[int(12)] = (d_600_v0_).intersection(d_600_v0_)
            nw91_[int(13)] = d_600_v0_
            nw91_[int(14)] = d_600_v0_
            nw91_[int(15)] = _dafny.MultiSet(d_605_v5_)
            nw91_[int(16)] = (_dafny.MultiSet(d_605_v5_)).set(((d_606_v6_)[d_604_v4_] if (d_604_v4_) in (d_606_v6_) else p1), default__.abs(d_604_v4_))
            nw91_[int(17)] = d_600_v0_
            nw91_[int(18)] = d_600_v0_
            d_607_v7_ = nw91_
            d_608_v8_: D6
            d_608_v8_ = D6_DC17(d_607_v7_)
            d_607_v7_ = (d_608_v8_).cf38
        elif True:
            d_609_v9_: _dafny.Array
            nw92_ = _dafny.Array(_dafny.Map({}), 29)
            d_609_v9_ = nw92_
            d_610_v10_: int
            d_610_v10_ = 164
            d_611_v11_: _dafny.Map
            d_611_v11_ = _dafny.Map({p1: d_610_v10_})
            index78_ = default__.safeIndex(679, (d_609_v9_).length(0))
            (d_609_v9_)[index78_] = d_611_v11_
            index79_ = default__.safeIndex(679, (d_609_v9_).length(0))
            (d_609_v9_)[index79_] = d_611_v11_
            d_612_v12_: _dafny.Array
            def lambda15_(d_613_v10_):
                def lambda16_(d_614_i0_):
                    return (d_613_v10_) < (d_613_v10_)

                return lambda16_

            init9_ = lambda15_(d_610_v10_)
            nw93_ = _dafny.Array(None, 6)
            for i0_9_ in range(nw93_.length(0)):
                nw93_[i0_9_] = init9_(i0_9_)
            d_612_v12_ = nw93_
            d_612_v12_ = d_612_v12_
            (globalState).f0 = d_610_v10_
            if p2:
                d_615_v13_: _dafny.Array
                nw94_ = _dafny.Array(None, 2)
                nw94_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))
                nw94_[int(1)] = (len(default__.fm39(p1, 288, d_610_v10_, globalState))) * (d_610_v10_)
                d_615_v13_ = nw94_
                index80_ = default__.safeIndex(172, (d_615_v13_).length(0))
                (d_615_v13_)[index80_] = d_610_v10_
                d_616_v15_: _dafny.Map
                d_616_v15_ = _dafny.Map({d_610_v10_: d_610_v10_})
                d_617_v16_: _dafny.Set
                def iife77_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in _dafny.IntegerRange(309, 139):
                        d_618_v14_: int = compr_25_
                        if ((309) <= (d_618_v14_)) and ((d_618_v14_) < (139)):
                            coll25_[default__.safeDivisionInt(d_618_v14_, d_610_v10_)] = d_610_v10_
                    return _dafny.Map(coll25_)
                d_617_v16_ = _dafny.Set({(iife77_()
                ) != (d_616_v15_), (self).f18})
                d_619_v17_: _dafny.Map
                d_619_v17_ = _dafny.Map({(d_611_v11_) | ((d_609_v9_)[default__.safeIndex(679, (d_609_v9_).length(0))]): d_617_v16_})
                index81_ = default__.safeIndex(172, (d_615_v13_).length(0))
                rhs79_ = d_610_v10_
                rhs80_ = ((d_619_v17_)[d_611_v11_] if (d_611_v11_) in (d_619_v17_) else d_617_v16_)
                rhs81_ = d_610_v10_
                lhs61_ = d_615_v13_
                lhs62_ = default__.safeIndex(172, (d_615_v13_).length(0))
                lhs63_ = globalState
                lhs61_[lhs62_] = rhs79_
                d_617_v16_ = rhs80_
                lhs63_.f15 = rhs81_
                d_620_v18_: _dafny.Map
                d_620_v18_ = _dafny.Map({True: (p1) == (p1)})
                d_621_v19_: _dafny.Array
                d_621_v19_ = d_612_v12_
                d_622_v20_: _dafny.Map
                d_622_v20_ = _dafny.Map({d_612_v12_: d_610_v10_})
                d_620_v18_ = (d_620_v18_).set(((d_621_v19_)) in (d_622_v20_), p1)
                d_623_v21_: D2
                d_623_v21_ = D2_DC7(d_610_v10_, not(p1), d_610_v10_)
                (globalState).f10 = (default__.fm40(d_623_v21_, globalState)).set(default__.safeIndex(len((self).f19), len(default__.fm40(d_623_v21_, globalState))), d_610_v10_)
                d_624_v22_: _dafny.Seq
                d_624_v22_ = _dafny.SeqWithoutIsStrInference([p1])
                d_625_v23_: _dafny.Map
                d_625_v23_ = _dafny.Map({d_624_v22_: p1})
                d_610_v10_ = ((d_615_v13_)[default__.safeIndex(172, (d_615_v13_).length(0))] if ((d_625_v23_)[d_624_v22_] if (d_624_v22_) in (d_625_v23_) else p2) else d_610_v10_)
                d_626_v24_: _dafny.Seq
                d_626_v24_ = _dafny.SeqWithoutIsStrInference([d_612_v12_])
                d_627_v25_: _dafny.Array
                nw95_ = _dafny.Array(None, 22)
                nw95_[int(0)] = d_612_v12_
                nw95_[int(1)] = d_612_v12_
                nw95_[int(2)] = d_612_v12_
                nw95_[int(3)] = d_612_v12_
                nw95_[int(4)] = d_612_v12_
                nw95_[int(5)] = d_612_v12_
                nw95_[int(6)] = d_612_v12_
                nw95_[int(7)] = d_612_v12_
                nw95_[int(8)] = d_612_v12_
                nw95_[int(9)] = d_612_v12_
                nw95_[int(10)] = d_612_v12_
                nw95_[int(11)] = d_612_v12_
                nw95_[int(12)] = d_612_v12_
                nw95_[int(13)] = d_612_v12_
                nw95_[int(14)] = d_612_v12_
                nw95_[int(15)] = d_612_v12_
                nw95_[int(16)] = d_612_v12_
                nw95_[int(17)] = (d_626_v24_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_610_v10_ for d_628_i1_ in range(default__.abs(839))])), len(d_626_v24_))]
                nw95_[int(18)] = d_612_v12_
                nw95_[int(19)] = d_612_v12_
                nw95_[int(20)] = d_612_v12_
                nw95_[int(21)] = d_612_v12_
                d_627_v25_ = nw95_
                index82_ = default__.safeIndex(212, (d_627_v25_).length(0))
                (d_627_v25_)[index82_] = d_612_v12_
                d_629_v26_: _dafny.Map
                d_629_v26_ = _dafny.Map({p1: d_624_v22_})
                d_630_v27_: _dafny.MultiSet
                d_630_v27_ = _dafny.MultiSet([d_611_v11_])
                d_631_v28_: _dafny.Map
                d_631_v28_ = _dafny.Map({d_610_v10_: p2})
                d_632_v29_: _dafny.MultiSet
                d_632_v29_ = _dafny.MultiSet([d_610_v10_, len((self).f19), len(d_631_v28_)])
                index83_ = default__.safeIndex(212, (d_627_v25_).length(0))
                rhs82_ = (d_612_v12_ if p2 else d_612_v12_)
                rhs83_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_633_i2_ in range(default__.abs(596))])).set(default__.safeIndex(len((d_629_v26_).set((self).f18, d_624_v22_)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_634_i2_ in range(default__.abs(596))]))), _dafny.CodePoint('j'))
                rhs84_ = ((d_630_v27_)[(d_609_v9_)[default__.safeIndex(679, (d_609_v9_).length(0))]] if ((d_609_v9_)[default__.safeIndex(679, (d_609_v9_).length(0))]) in (d_630_v27_) else ((d_632_v29_)[d_610_v10_] if (d_610_v10_) in (d_632_v29_) else len((self).f19)))
                rhs85_ = ((d_600_v0_)[not (not(not(p2))) or (p2)] if (not (not(not(p2))) or (p2)) in (d_600_v0_) else d_610_v10_)
                lhs64_ = d_627_v25_
                lhs65_ = default__.safeIndex(212, (d_627_v25_).length(0))
                lhs66_ = globalState
                lhs67_ = globalState
                lhs68_ = globalState
                lhs64_[lhs65_] = rhs82_
                lhs66_.f9 = rhs83_
                lhs67_.f15 = rhs84_
                lhs68_.f15 = rhs85_
            elif True:
                d_635_v30_: C1
                nw96_ = C1()
                nw96_.ctor__(p2)
                d_635_v30_ = nw96_
                d_636_v31_: C1
                nw97_ = C1()
                nw97_.ctor__(not(p1))
                d_636_v31_ = nw97_
                (globalState).f3 = default__.safeDivisionInt((-898) * (d_610_v10_), d_610_v10_)
                d_637_v32_: _dafny.Array
                nw98_ = _dafny.Array(int(0), 9)
                d_637_v32_ = nw98_
                d_638_v33_: _dafny.Array
                d_638_v33_ = d_637_v32_
                d_638_v33_ = d_638_v33_
                d_639_v34_: _dafny.Seq
                d_639_v34_ = _dafny.SeqWithoutIsStrInference([p2])
                index84_ = default__.safeIndex(613, (d_637_v32_).length(0))
                (d_637_v32_)[index84_] = len(d_639_v34_)
                index85_ = default__.safeIndex(613, (d_637_v32_).length(0))
                (d_637_v32_)[index85_] = (0) - (d_610_v10_)
            (globalState).f2 = (p1 if True else (self).f18)
        (globalState).f2 = (self).f18
        d_640_v35_: _dafny.Array
        def lambda17_(d_641_p1_, d_642_p2_):
            def lambda18_(d_643_i3_):
                return (_dafny.Map({not(d_641_p1_): (self).f18})) | (_dafny.Map({(self).f18: d_642_p2_}))

            return lambda18_

        init10_ = lambda17_(p1, p2)
        nw99_ = _dafny.Array(None, 4)
        for i0_10_ in range(nw99_.length(0)):
            nw99_[i0_10_] = init10_(i0_10_)
        d_640_v35_ = nw99_
        d_644_v36_: D3
        d_644_v36_ = D3_DC11(p1, (self).f18)
        pat_let_tv7_ = p1
        pat_let_tv8_ = p2
        pat_let_tv9_ = p2
        pat_let_tv10_ = p1
        pat_let_tv11_ = d_600_v0_
        pat_let_tv12_ = d_600_v0_
        pat_let_tv13_ = d_600_v0_
        pat_let_tv14_ = p1
        def lambda19_(source5_):
            if source5_.is_DC11:
                d_645___mcc_h0_ = source5_.cf23
                d_646___mcc_h1_ = source5_.cf24
                d_647_cf24_ = d_646___mcc_h1_
                d_648_cf23_ = d_645___mcc_h0_
                return len((_dafny.SeqWithoutIsStrInference([d_647_cf24_]) if not(d_648_cf23_) else _dafny.SeqWithoutIsStrInference([pat_let_tv7_])))
            elif source5_.is_DC12:
                d_649___mcc_h2_ = source5_.cf25
                d_650___mcc_h3_ = source5_.cf26
                d_651___mcc_h4_ = source5_.cf27
                d_652___mcc_h5_ = source5_.cf28
                d_653___mcc_h6_ = source5_.cf29
                d_654_cf29_ = d_653___mcc_h6_
                d_655_cf28_ = d_652___mcc_h5_
                d_656_cf27_ = d_651___mcc_h4_
                d_657_cf26_ = d_650___mcc_h3_
                d_658_cf25_ = d_649___mcc_h2_
                return default__.safeDivisionInt(d_656_cf27_, d_656_cf27_)
            elif source5_.is_DC13:
                d_659___mcc_h7_ = source5_.cf30
                d_660___mcc_h8_ = source5_.cf31
                d_661___mcc_h9_ = source5_.cf32
                d_662___mcc_h10_ = source5_.cf33
                d_663___mcc_h11_ = source5_.cf34
                d_664_cf34_ = d_663___mcc_h11_
                d_665_cf33_ = d_662___mcc_h10_
                d_666_cf32_ = d_661___mcc_h9_
                d_667_cf31_ = d_660___mcc_h8_
                d_668_cf30_ = d_659___mcc_h7_
                return default__.safeDivisionInt(d_665_cf33_, len(_dafny.Set({pat_let_tv8_})))
            elif source5_.is_DC10:
                d_669___mcc_h12_ = source5_.cf22
                d_670_cf22_ = d_669___mcc_h12_
                return (D0_DC1((D0_DC0((self).f18)).cf0, 971, (self).f19, pat_let_tv9_, -299)).cf5
            elif True:
                d_671___mcc_h13_ = source5_.cf35
                d_672_cf35_ = d_671___mcc_h13_
                return ((_dafny.MultiSet([len((self).f19)])) | (_dafny.MultiSet([len((D3_DC13(len(_dafny.Map({pat_let_tv10_: len((self).f19)})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukjh")), (pat_let_tv11_).cardinality, ((pat_let_tv12_)[True] if (True) in (pat_let_tv13_) else len((self).f19)), _dafny.CodePoint('i'))).cf31)]))).cardinality

        def iife78_(_pat_let26_0):
            def iife79_(d_673_dt__update__tmp_h0_):
                def iife80_(_pat_let27_0):
                    def iife81_(d_674_dt__update_hcf23_h0_):
                        return D3_DC11(d_674_dt__update_hcf23_h0_, (d_673_dt__update__tmp_h0_).cf24)
                    return iife81_(_pat_let27_0)
                return iife80_(pat_let_tv14_)
            return iife79_(_pat_let26_0)
        rhs86_ = (((self).f19) + ((self).f19)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnb"))) + ((self).f19))
        rhs87_ = d_640_v35_
        rhs88_ = lambda19_(iife78_(d_644_v36_))
        lhs69_ = globalState
        lhs70_ = globalState
        lhs69_.f12 = rhs86_
        d_640_v35_ = rhs87_
        lhs70_.f3 = rhs88_
        d_675_v37_: int
        d_675_v37_ = 129
        d_676_i4_: int
        d_676_i4_ = 0
        with _dafny.label("4"):
            while (self).fm3(d_675_v37_, globalState):
                with _dafny.c_label("4"):
                    if (d_676_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_676_i4_ = (d_676_i4_) + (1)
                    d_677_v38_: _dafny.Array
                    nw100_ = _dafny.Array(None, 9)
                    nw100_[int(0)] = d_675_v37_
                    nw100_[int(1)] = (0) - ((0) - (d_675_v37_))
                    nw100_[int(2)] = (0) - (d_675_v37_)
                    nw100_[int(3)] = -119
                    nw100_[int(4)] = d_675_v37_
                    nw100_[int(5)] = d_675_v37_
                    nw100_[int(6)] = d_675_v37_
                    nw100_[int(7)] = d_675_v37_
                    nw100_[int(8)] = d_675_v37_
                    d_677_v38_ = nw100_
                    d_678_v39_: _dafny.Map
                    d_678_v39_ = _dafny.Map({(self).f19: d_677_v38_})
                    d_678_v39_ = (d_678_v39_).set((self).f19, d_677_v38_)
                    d_679_v40_: str
                    d_679_v40_ = _dafny.CodePoint('u')
                    index86_ = default__.safeIndex(414, (d_677_v38_).length(0))
                    (d_677_v38_)[index86_] = d_675_v37_
                    index87_ = default__.safeIndex(414, (d_677_v38_).length(0))
                    rhs89_ = d_679_v40_
                    rhs90_ = d_675_v37_
                    lhs71_ = d_677_v38_
                    lhs72_ = default__.safeIndex(414, (d_677_v38_).length(0))
                    d_679_v40_ = rhs89_
                    lhs71_[lhs72_] = rhs90_
                    (globalState).f0 = d_675_v37_
                    (globalState).f3 = 117
                    pass
            pass
        d_680_i5_: int
        d_680_i5_ = 0
        with _dafny.label("5"):
            while p2:
                with _dafny.c_label("5"):
                    if (d_680_i5_) >= (100):
                        raise _dafny.Break("5")
                    d_680_i5_ = (d_680_i5_) + (1)
                    d_681_v41_: _dafny.Array
                    def lambda20_(d_682_v37_):
                        def lambda21_(d_683_i6_):
                            return (d_683_i6_) * (d_682_v37_)

                        return lambda21_

                    init11_ = lambda20_(d_675_v37_)
                    nw101_ = _dafny.Array(None, 4)
                    for i0_11_ in range(nw101_.length(0)):
                        nw101_[i0_11_] = init11_(i0_11_)
                    d_681_v41_ = nw101_
                    d_684_v42_: _dafny.Seq
                    d_684_v42_ = _dafny.SeqWithoutIsStrInference([(self).f18, p1, (self).f18, (self).f18, p2])
                    index88_ = default__.safeIndex(306, (d_681_v41_).length(0))
                    (d_681_v41_)[index88_] = len(d_684_v42_)
                    d_685_v43_: _dafny.Array
                    nw102_ = _dafny.Array(False, 13)
                    d_685_v43_ = nw102_
                    d_686_v44_: _dafny.Seq
                    d_686_v44_ = _dafny.SeqWithoutIsStrInference([d_685_v43_, d_685_v43_, d_685_v43_])
                    d_687_v45_: _dafny.Seq
                    d_687_v45_ = _dafny.SeqWithoutIsStrInference([d_686_v44_, (d_686_v44_) + (d_686_v44_)])
                    d_688_v46_: _dafny.MultiSet
                    d_688_v46_ = _dafny.MultiSet([d_675_v37_])
                    index89_ = default__.safeIndex(306, (d_681_v41_).length(0))
                    rhs91_ = d_675_v37_
                    rhs92_ = (d_687_v45_)[default__.safeIndex(((d_688_v46_)[d_675_v37_] if (d_675_v37_) in (d_688_v46_) else d_675_v37_), len(d_687_v45_))]
                    rhs93_ = len((d_684_v42_) + ((d_684_v42_).set(default__.safeIndex(368, len(d_684_v42_)), p2)))
                    lhs73_ = d_681_v41_
                    lhs74_ = default__.safeIndex(306, (d_681_v41_).length(0))
                    lhs73_[lhs74_] = rhs91_
                    d_686_v44_ = rhs92_
                    d_675_v37_ = rhs93_
                    index90_ = default__.safeIndex(306, (d_681_v41_).length(0))
                    (d_681_v41_)[index90_] = default__.safeModuloInt(((self).f35)[default__.safeIndex(771, len((self).f35))], (287) - (583))
                    (globalState).f2 = False
                    d_689_v47_: _dafny.Map
                    d_689_v47_ = _dafny.Map({True: d_685_v43_})
                    d_690_v48_: _dafny.Map
                    d_690_v48_ = _dafny.Map({227: (self).f35})
                    d_691_v49_: C2
                    nw103_ = C2()
                    nw103_.ctor__(len(d_689_v47_), (_dafny.SeqWithoutIsStrInference([((self).f35)[default__.safeIndex(len((self).f35), len((self).f35))] for d_692_i7_ in range(default__.abs(774))])) + (((d_690_v48_)[(d_681_v41_)[default__.safeIndex(306, (d_681_v41_).length(0))]] if ((d_681_v41_)[default__.safeIndex(306, (d_681_v41_).length(0))]) in (d_690_v48_) else (self).f35)), (self).f18)
                    d_691_v49_ = nw103_
                    pass
            pass
        r0 = (self).fm38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhgblgpc")), globalState)
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_693_v0_: _dafny.Array
        def lambda22_(d_694_i1_):
            return (self).f18

        init12_ = lambda22_
        nw104_ = _dafny.Array(None, 24)
        for i0_12_ in range(nw104_.length(0)):
            nw104_[i0_12_] = init12_(i0_12_)
        d_693_v0_ = nw104_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_693_v0_).length(0)):
            d_695_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_695_i0_)) and ((d_695_i0_) < ((d_693_v0_).length(0)))):
                (d_693_v0_)[(d_695_i0_)] = ((self).f18 if False else ((0) - (p1)) <= (p1))
        d_696_v1_: _dafny.Array
        def lambda23_(d_697_i2_):
            return (_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])) + (_dafny.SeqWithoutIsStrInference([(self).f18, True, True, (self).f18]))

        init13_ = lambda23_
        nw105_ = _dafny.Array(None, 10)
        for i0_13_ in range(nw105_.length(0)):
            nw105_[i0_13_] = init13_(i0_13_)
        d_696_v1_ = nw105_
        d_698_v2_: _dafny.Seq
        d_698_v2_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
        index91_ = default__.safeIndex(264, (d_696_v1_).length(0))
        (d_696_v1_)[index91_] = d_698_v2_
        index92_ = default__.safeIndex(264, (d_696_v1_).length(0))
        (d_696_v1_)[index92_] = d_698_v2_
        (globalState).f0 = p1
        (globalState).f2 = False
        d_699_v3_: C1
        nw106_ = C1()
        nw106_.ctor__((self).f18)
        d_699_v3_ = nw106_
        (globalState).f2 = (self).f18

    def m9(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        r2: bool = False
        d_700_v0_: int
        d_700_v0_ = 154
        d_701_v1_: _dafny.Seq
        d_701_v1_ = _dafny.SeqWithoutIsStrInference([not((self).fm3(d_700_v0_, globalState))])
        d_702_i0_: int
        d_702_i0_ = 0
        with _dafny.label("6"):
            while ((self).f18) in (d_701_v1_):
                with _dafny.c_label("6"):
                    if (d_702_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_702_i0_ = (d_702_i0_) + (1)
                    rhs94_ = (0) - (-718)
                    rhs95_ = d_700_v0_
                    rhs96_ = (self).f18
                    lhs75_ = globalState
                    lhs76_ = globalState
                    lhs77_ = globalState
                    lhs75_.f0 = rhs94_
                    lhs76_.f3 = rhs95_
                    lhs77_.f2 = rhs96_
                    (globalState).f0 = len((self).f19)
                    d_703_v2_: _dafny.Array
                    def lambda24_(d_704_v0_, d_705_v1_):
                        def lambda25_(d_706_i1_):
                            return ((D8_DC20(_dafny.Set({d_704_v0_}))).cf40) - (_dafny.Set({len((self).f19), (D0_DC1((self).f18, d_704_v0_, (self).f19, True, (_dafny.MultiSet((self).f35)).cardinality)).cf5, (D2_DC7(len(d_705_v1_), (self).f18, len(_dafny.Set({(self).f18, (self).f18})))).cf15}))

                        return lambda25_

                    init14_ = lambda24_(d_700_v0_, d_701_v1_)
                    nw107_ = _dafny.Array(None, 20)
                    for i0_14_ in range(nw107_.length(0)):
                        nw107_[i0_14_] = init14_(i0_14_)
                    d_703_v2_ = nw107_
                    d_703_v2_ = d_703_v2_
                    d_700_v0_ = ((self).f35)[default__.safeIndex(d_700_v0_, len((self).f35))]
                    pass
            pass
        d_707_v3_: D1
        d_707_v3_ = D1_DC4((self).f18)
        d_708_v4_: _dafny.Array
        nw108_ = _dafny.Array(None, 16)
        nw108_[int(0)] = (self).f18
        nw108_[int(1)] = (self).f18
        nw108_[int(2)] = (d_707_v3_).cf12
        nw108_[int(3)] = (self).f18
        nw108_[int(4)] = (self).f18
        nw108_[int(5)] = True
        nw108_[int(6)] = (default__.fm41(globalState)).cf0
        nw108_[int(7)] = (self).f18
        nw108_[int(8)] = (self).f18
        nw108_[int(9)] = (self).f18
        nw108_[int(10)] = (d_701_v1_)[default__.safeIndex(d_700_v0_, len(d_701_v1_))]
        nw108_[int(11)] = (self).f18
        nw108_[int(12)] = not((self).f18)
        nw108_[int(13)] = (self).f18
        nw108_[int(14)] = (self).f18
        nw108_[int(15)] = (self).f18
        d_708_v4_ = nw108_
        d_709_v5_: _dafny.Map
        d_709_v5_ = _dafny.Map({d_708_v4_: d_708_v4_})
        d_710_v6_: _dafny.Array
        nw109_ = _dafny.Array(None, 14)
        nw109_[int(0)] = (d_709_v5_) | (_dafny.Map({d_708_v4_: d_708_v4_}))
        nw109_[int(1)] = d_709_v5_
        nw109_[int(2)] = d_709_v5_
        nw109_[int(3)] = d_709_v5_
        nw109_[int(4)] = d_709_v5_
        nw109_[int(5)] = d_709_v5_
        nw109_[int(6)] = d_709_v5_
        nw109_[int(7)] = d_709_v5_
        nw109_[int(8)] = d_709_v5_
        nw109_[int(9)] = (d_709_v5_) | (d_709_v5_)
        nw109_[int(10)] = d_709_v5_
        nw109_[int(11)] = d_709_v5_
        nw109_[int(12)] = d_709_v5_
        nw109_[int(13)] = (d_709_v5_).set(d_708_v4_, d_708_v4_)
        d_710_v6_ = nw109_
        index93_ = default__.safeIndex(854, (d_710_v6_).length(0))
        (d_710_v6_)[index93_] = (d_709_v5_ if (self).f18 else (d_709_v5_).set(d_708_v4_, d_708_v4_))
        index94_ = default__.safeIndex(854, (d_710_v6_).length(0))
        (d_710_v6_)[index94_] = d_709_v5_
        d_711_v7_: _dafny.Array
        nw110_ = _dafny.Array(int(0), 23)
        d_711_v7_ = nw110_
        index95_ = default__.safeIndex(306, (d_711_v7_).length(0))
        (d_711_v7_)[index95_] = d_700_v0_
        index96_ = default__.safeIndex(306, (d_711_v7_).length(0))
        (d_711_v7_)[index96_] = 257
        d_712_v8_: _dafny.Set
        d_712_v8_ = _dafny.Set({d_700_v0_})
        d_713_i2_: int
        d_713_i2_ = 0
        with _dafny.label("7"):
            while (d_712_v8_) == ((d_712_v8_) | (d_712_v8_)):
                with _dafny.c_label("7"):
                    if (d_713_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_713_i2_ = (d_713_i2_) + (1)
                    d_714_v9_: _dafny.Map
                    d_714_v9_ = _dafny.Map({(d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]: (d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]})
                    d_715_v10_: _dafny.Set
                    d_715_v10_ = _dafny.Set({(self).f18, False})
                    d_716_v11_: str
                    d_716_v11_ = _dafny.CodePoint('b')
                    d_717_v12_: D3
                    d_717_v12_ = D3_DC13((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))], (self).f19, default__.safeModuloInt(-519, (d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]), ((d_714_v9_)[d_700_v0_] if (d_700_v0_) in (d_714_v9_) else len(d_715_v10_)), d_716_v11_)
                    source6_ = d_717_v12_
                    if source6_.is_DC11:
                        d_718___mcc_h0_ = source6_.cf23
                        d_719___mcc_h1_ = source6_.cf24
                        d_720_cf24_ = d_719___mcc_h1_
                        d_721_cf23_ = d_718___mcc_h0_
                        (globalState).f15 = d_700_v0_
                        (globalState).f2 = not((d_720_cf24_ if d_720_cf24_ else (self).fm2((self).f18, (self).f18, d_720_cf24_, globalState)))
                        d_722_v13_: _dafny.Array
                        nw111_ = _dafny.Array(_dafny.MultiSet({}), 9)
                        d_722_v13_ = nw111_
                        d_723_v14_: _dafny.MultiSet
                        d_723_v14_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvysx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfamwn"))])
                        index97_ = default__.safeIndex(620, (d_722_v13_).length(0))
                        (d_722_v13_)[index97_] = (d_723_v14_) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_716_v11_ for d_724_i3_ in range(default__.abs(495))])]))
                        index98_ = default__.safeIndex(620, (d_722_v13_).length(0))
                        (d_722_v13_)[index98_] = (d_723_v14_) - (_dafny.MultiSet([(self).f19, (self).f19, (self).f19, (self).fm37(d_700_v0_, globalState)]))
                        (globalState).f9 = ((self).f19) + ((self).f19)
                    elif source6_.is_DC12:
                        d_725___mcc_h2_ = source6_.cf25
                        d_726___mcc_h3_ = source6_.cf26
                        d_727___mcc_h4_ = source6_.cf27
                        d_728___mcc_h5_ = source6_.cf28
                        d_729___mcc_h6_ = source6_.cf29
                        d_730_cf29_ = d_729___mcc_h6_
                        d_731_cf28_ = d_728___mcc_h5_
                        d_732_cf27_ = d_727___mcc_h4_
                        d_733_cf26_ = d_726___mcc_h3_
                        d_734_cf25_ = d_725___mcc_h2_
                        d_715_v10_ = (d_715_v10_) | (d_715_v10_)
                        index99_ = default__.safeIndex(306, (d_711_v7_).length(0))
                        (d_711_v7_)[index99_] = (d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]
                        (globalState).f2 = d_733_cf26_
                        (globalState).f3 = -860
                    elif source6_.is_DC13:
                        d_735___mcc_h7_ = source6_.cf30
                        d_736___mcc_h8_ = source6_.cf31
                        d_737___mcc_h9_ = source6_.cf32
                        d_738___mcc_h10_ = source6_.cf33
                        d_739___mcc_h11_ = source6_.cf34
                        d_740_cf34_ = d_739___mcc_h11_
                        d_741_cf33_ = d_738___mcc_h10_
                        d_742_cf32_ = d_737___mcc_h9_
                        d_743_cf31_ = d_736___mcc_h8_
                        d_744_cf30_ = d_735___mcc_h7_
                        (globalState).f3 = (0) - ((0) - (d_744_cf30_))
                        d_745_v15_: _dafny.Map
                        d_745_v15_ = _dafny.Map({default__.safeDivisionInt((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))], d_744_cf30_): _dafny.SeqWithoutIsStrInference([d_700_v0_ for d_746_i4_ in range(default__.abs(816))])})
                        d_745_v15_ = (d_745_v15_).set((0) - (d_744_cf30_), (self).f35)
                        d_747_v16_: _dafny.Array
                        nw112_ = _dafny.Array(_dafny.Map({}), 2)
                        d_747_v16_ = nw112_
                        index100_ = default__.safeIndex(193, (d_747_v16_).length(0))
                        (d_747_v16_)[index100_] = _dafny.Map({d_742_cf32_: len(_dafny.SeqWithoutIsStrInference([d_716_v11_, d_716_v11_]))})
                        index101_ = default__.safeIndex(193, (d_747_v16_).length(0))
                        (d_747_v16_)[index101_] = d_714_v9_
                        (globalState).f10 = (default__.fm42((self).f18, (self).f18, (self).f18, globalState)).cf43
                    elif source6_.is_DC10:
                        d_748___mcc_h12_ = source6_.cf22
                        d_749_cf22_ = d_748___mcc_h12_
                        index102_ = default__.safeIndex(202, (d_708_v4_).length(0))
                        (d_708_v4_)[index102_] = (self).fm2(False, True, (self).f18, globalState)
                        d_750_v17_: D0
                        d_750_v17_ = D0_DC1(False, len(_dafny.Map({(0) - ((self).fm38((self).f19, globalState)): (0) - (d_700_v0_)})), (self).f19, (self).f18, (self).fm38((self).f19, globalState))
                        index103_ = default__.safeIndex(306, (d_711_v7_).length(0))
                        index104_ = default__.safeIndex(202, (d_708_v4_).length(0))
                        rhs97_ = (d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]
                        rhs98_ = (d_750_v17_).cf2
                        rhs99_ = True
                        lhs78_ = d_711_v7_
                        lhs79_ = default__.safeIndex(306, (d_711_v7_).length(0))
                        lhs80_ = d_708_v4_
                        lhs81_ = default__.safeIndex(202, (d_708_v4_).length(0))
                        d_700_v0_ = rhs97_
                        lhs78_[lhs79_] = rhs98_
                        lhs80_[lhs81_] = rhs99_
                        d_751_v18_: _dafny.Seq
                        d_751_v18_ = _dafny.SeqWithoutIsStrInference([d_711_v7_])
                        d_752_v19_: _dafny.Array
                        def lambda26_(d_753_v11_):
                            def lambda27_(d_754_i5_):
                                return d_753_v11_

                            return lambda27_

                        init15_ = lambda26_(d_716_v11_)
                        nw113_ = _dafny.Array(None, 1)
                        for i0_15_ in range(nw113_.length(0)):
                            nw113_[i0_15_] = init15_(i0_15_)
                        d_752_v19_ = nw113_
                        index105_ = default__.safeIndex(10, (d_752_v19_).length(0))
                        (d_752_v19_)[index105_] = d_716_v11_
                        d_755_v20_: _dafny.Map
                        d_755_v20_ = _dafny.Map({(self).f18: d_751_v18_})
                        index106_ = default__.safeIndex(306, (d_711_v7_).length(0))
                        index107_ = default__.safeIndex(10, (d_752_v19_).length(0))
                        index108_ = default__.safeIndex(202, (d_708_v4_).length(0))
                        rhs100_ = ((d_751_v18_) + (d_751_v18_)) + ((((d_755_v20_)[(d_708_v4_)[default__.safeIndex(202, (d_708_v4_).length(0))]] if ((d_708_v4_)[default__.safeIndex(202, (d_708_v4_).length(0))]) in (d_755_v20_) else d_751_v18_)).set(default__.safeIndex((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))], len(((d_755_v20_)[(d_708_v4_)[default__.safeIndex(202, (d_708_v4_).length(0))]] if ((d_708_v4_)[default__.safeIndex(202, (d_708_v4_).length(0))]) in (d_755_v20_) else d_751_v18_))), d_711_v7_))
                        rhs101_ = (d_700_v0_) + ((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))])
                        rhs102_ = d_716_v11_
                        rhs103_ = (d_708_v4_)[default__.safeIndex(202, (d_708_v4_).length(0))]
                        lhs82_ = d_711_v7_
                        lhs83_ = default__.safeIndex(306, (d_711_v7_).length(0))
                        lhs84_ = d_752_v19_
                        lhs85_ = default__.safeIndex(10, (d_752_v19_).length(0))
                        lhs86_ = d_708_v4_
                        lhs87_ = default__.safeIndex(202, (d_708_v4_).length(0))
                        d_751_v18_ = rhs100_
                        lhs82_[lhs83_] = rhs101_
                        lhs84_[lhs85_] = rhs102_
                        lhs86_[lhs87_] = rhs103_
                        index109_ = default__.safeIndex(306, (d_711_v7_).length(0))
                        (d_711_v7_)[index109_] = len((self).fm37((d_700_v0_) - (d_700_v0_), globalState))
                        d_756_v21_: C0
                        nw114_ = C0()
                        nw114_.ctor__(len((self).f19), ((d_714_v9_)[(d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]] if ((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]) in (d_714_v9_) else (0) - (d_700_v0_)))
                        d_756_v21_ = nw114_
                    elif True:
                        d_757___mcc_h13_ = source6_.cf35
                        d_758_cf35_ = d_757___mcc_h13_
                        d_759_v22_: _dafny.MultiSet
                        d_759_v22_ = _dafny.MultiSet([d_716_v11_, _dafny.CodePoint('f'), d_716_v11_])
                        d_700_v0_ = ((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]) + (((d_759_v22_)[_dafny.CodePoint('u')] if (_dafny.CodePoint('u')) in (d_759_v22_) else ((d_714_v9_)[d_700_v0_] if (d_700_v0_) in (d_714_v9_) else d_700_v0_)))
                        d_760_v23_: _dafny.MultiSet
                        d_760_v23_ = _dafny.MultiSet([(d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))], d_700_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuiukyo"))), len(d_701_v1_)])
                        d_761_v24_: _dafny.Seq
                        d_761_v24_ = _dafny.SeqWithoutIsStrInference([d_701_v1_, d_701_v1_])
                        r1 = (((d_760_v23_)[d_700_v0_] if (d_700_v0_) in (d_760_v23_) else (_dafny.MultiSet(d_761_v24_)).cardinality)) <= (default__.safeDivisionInt(d_700_v0_, (d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]))
                        d_762_v25_: _dafny.Array
                        nw115_ = _dafny.Array(_dafny.MultiSet({}), 18)
                        d_762_v25_ = nw115_
                        index110_ = default__.safeIndex(718, (d_762_v25_).length(0))
                        (d_762_v25_)[index110_] = _dafny.MultiSet([(self).f18])
                        index111_ = default__.safeIndex(718, (d_762_v25_).length(0))
                        rhs104_ = (_dafny.SeqWithoutIsStrInference([(d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))], d_700_v0_, -405])) != ((self).f35)
                        rhs105_ = _dafny.SeqWithoutIsStrInference([746])
                        rhs106_ = not((self).f18)
                        rhs107_ = d_700_v0_
                        rhs108_ = _dafny.MultiSet([(self).f18])
                        lhs88_ = globalState
                        lhs89_ = globalState
                        lhs90_ = globalState
                        lhs91_ = d_762_v25_
                        lhs92_ = default__.safeIndex(718, (d_762_v25_).length(0))
                        lhs88_.f2 = rhs104_
                        lhs89_.f10 = rhs105_
                        r1 = rhs106_
                        lhs90_.f15 = rhs107_
                        lhs91_[lhs92_] = rhs108_
                        d_763_v26_: _dafny.Array
                        nw116_ = _dafny.Array(_dafny.Seq({}), 20)
                        d_763_v26_ = nw116_
                        index112_ = default__.safeIndex(247, (d_763_v26_).length(0))
                        (d_763_v26_)[index112_] = (self).f35
                        d_764_v27_: _dafny.Map
                        d_764_v27_ = _dafny.Map({len((self).f35): (self).f18})
                        index113_ = default__.safeIndex(247, (d_763_v26_).length(0))
                        rhs109_ = (self).f18
                        rhs110_ = d_700_v0_
                        rhs111_ = not(((self).f18) or ((self).f18))
                        rhs112_ = ((self).f35) + ((_dafny.SeqWithoutIsStrInference([len(d_764_v27_)])) + ((self).f35))
                        rhs113_ = d_715_v10_
                        lhs93_ = globalState
                        lhs94_ = globalState
                        lhs95_ = d_763_v26_
                        lhs96_ = default__.safeIndex(247, (d_763_v26_).length(0))
                        r2 = rhs109_
                        lhs93_.f0 = rhs110_
                        lhs94_.f2 = rhs111_
                        lhs95_[lhs96_] = rhs112_
                        d_715_v10_ = rhs113_
                    d_715_v10_ = _dafny.Set({True, (not((self).f18)) or (not((self).f18)), (self).f18, (self).f18})
                    index114_ = default__.safeIndex(306, (d_711_v7_).length(0))
                    (d_711_v7_)[index114_] = (d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]
                    index115_ = default__.safeIndex(306, (d_711_v7_).length(0))
                    (d_711_v7_)[index115_] = (-145) * ((d_700_v0_) - ((d_711_v7_)[default__.safeIndex(306, (d_711_v7_).length(0))]))
                    pass
            pass
        (globalState).f10 = (self).f35
        (globalState).f0 = (d_700_v0_) * (d_700_v0_)
        r0 = (self).f19
        r1 = (self).f18
        r2 = (self).f18
        return r0, r1, r2

    def m10(self, p0, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: int = int(0)
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_765_v0_: D10
        d_765_v0_ = D10_DC24(_dafny.Set({(self).f18, p0}))
        d_766_v1_: _dafny.Set
        d_766_v1_ = _dafny.Set({(self).f18})
        if not((((d_765_v0_).cf47).intersection(d_766_v1_)) == (d_766_v1_)):
            d_767_v2_: int
            d_767_v2_ = 910
            d_768_v3_: C2
            nw117_ = C2()
            nw117_.ctor__(d_767_v2_, (self).f35, (self).fm2((self).f18, True, p0, globalState))
            d_768_v3_ = nw117_
            (globalState).f2 = (self).fm3((self).fm38((self).f19, globalState), globalState)
            d_769_v4_: _dafny.Array
            nw118_ = _dafny.Array(False, 4)
            d_769_v4_ = nw118_
            index116_ = default__.safeIndex(596, (d_769_v4_).length(0))
            (d_769_v4_)[index116_] = (self).f18
            index117_ = default__.safeIndex(596, (d_769_v4_).length(0))
            (d_769_v4_)[index117_] = p0
            d_770_v5_: D9
            d_770_v5_ = D9_DC23((self).f18, p0, d_768_v3_.f33)
            d_771_v6_: _dafny.Array
            nw119_ = _dafny.Array(int(0), 28)
            d_771_v6_ = nw119_
            d_772_v7_: _dafny.Array
            d_772_v7_ = d_771_v6_
            d_773_v8_: _dafny.Array
            nw120_ = _dafny.Array(None, 13)
            nw120_[int(0)] = d_771_v6_
            nw120_[int(1)] = d_771_v6_
            nw120_[int(2)] = d_771_v6_
            nw120_[int(3)] = d_772_v7_
            nw120_[int(4)] = d_772_v7_
            nw120_[int(5)] = d_771_v6_
            nw120_[int(6)] = d_772_v7_
            nw120_[int(7)] = d_772_v7_
            nw120_[int(8)] = d_771_v6_
            nw120_[int(9)] = d_772_v7_
            nw120_[int(10)] = d_772_v7_
            nw120_[int(11)] = d_771_v6_
            nw120_[int(12)] = d_772_v7_
            d_773_v8_ = nw120_
            d_774_v9_: _dafny.Seq
            d_774_v9_ = _dafny.SeqWithoutIsStrInference([d_773_v8_])
            d_775_v10_: _dafny.Seq
            d_775_v10_ = _dafny.SeqWithoutIsStrInference([d_771_v6_])
            d_776_v11_: _dafny.Seq
            d_776_v11_ = _dafny.SeqWithoutIsStrInference([p0])
            d_777_v12_: _dafny.Map
            d_777_v12_ = _dafny.Map({d_768_v3_.f33: len(d_776_v11_)})
            d_778_v13_: D0
            d_778_v13_ = D0_DC1(p0, len(d_777_v12_), (self).f19, (self).f18, d_767_v2_)
            d_779_v14_: D1
            d_779_v14_ = D1_DC3((d_770_v5_).cf46, 589, default__.fm1((0) - (d_768_v3_.f33), True, (d_769_v4_)[default__.safeIndex(596, (d_769_v4_).length(0))], globalState), (len(d_774_v9_)) + (d_767_v2_), default__.safeModuloInt(len(d_775_v10_), (d_778_v13_).cf2))
            d_779_v14_ = d_779_v14_
            (globalState).f3 = default__.safeDivisionInt(d_768_v3_.f33, default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference([(0) - (d_767_v2_) for d_780_i0_ in range(default__.abs(53))]))), len((self).f19)))
        elif True:
            (globalState).f2 = (self).f18
            (globalState).f2 = (self).f18
            d_781_v15_: _dafny.MultiSet
            d_781_v15_ = _dafny.MultiSet([p0])
            d_782_v16_: C2
            nw121_ = C2()
            nw121_.ctor__((d_781_v15_).cardinality, (self).f35, (self).f18)
            d_782_v16_ = nw121_
            nw122_ = C2()
            nw122_.ctor__(default__.safeDivisionInt(d_782_v16_.f33, d_782_v16_.f33), (d_782_v16_).f34, p0)
            d_782_v16_ = nw122_
            d_783_v17_: _dafny.Map
            d_783_v17_ = _dafny.Map({d_782_v16_.f33: (505) > (d_782_v16_.f33)})
            d_784_v18_: _dafny.MultiSet
            d_784_v18_ = _dafny.MultiSet([_dafny.Map({d_782_v16_.f33: (self).fm3(d_782_v16_.f33, globalState)})])
            d_783_v17_ = (d_783_v17_).set(d_782_v16_.f33, (d_784_v18_).isdisjoint(d_784_v18_))
            d_785_v19_: _dafny.Array
            def lambda28_(d_786_i1_):
                return (self).f18

            init16_ = lambda28_
            nw123_ = _dafny.Array(None, 3)
            for i0_16_ in range(nw123_.length(0)):
                nw123_[i0_16_] = init16_(i0_16_)
            d_785_v19_ = nw123_
            d_787_v20_: _dafny.Seq
            d_787_v20_ = _dafny.SeqWithoutIsStrInference([d_785_v19_, d_785_v19_, d_785_v19_])
            d_788_v21_: _dafny.Array
            d_788_v21_ = d_785_v19_
            d_789_v22_: str
            d_789_v22_ = _dafny.CodePoint('a')
            d_790_v23_: _dafny.Map
            d_790_v23_ = _dafny.Map({d_789_v22_: d_785_v19_})
            d_791_v24_: _dafny.Array
            nw124_ = _dafny.Array(None, 25)
            nw124_[int(0)] = d_785_v19_
            nw124_[int(1)] = d_785_v19_
            nw124_[int(2)] = d_785_v19_
            nw124_[int(3)] = d_785_v19_
            nw124_[int(4)] = (d_785_v19_ if p0 else d_785_v19_)
            nw124_[int(5)] = d_785_v19_
            nw124_[int(6)] = d_785_v19_
            nw124_[int(7)] = d_785_v19_
            nw124_[int(8)] = d_785_v19_
            nw124_[int(9)] = d_785_v19_
            nw124_[int(10)] = d_785_v19_
            nw124_[int(11)] = (d_787_v20_)[default__.safeIndex(d_782_v16_.f33, len(d_787_v20_))]
            nw124_[int(12)] = (d_788_v21_)
            nw124_[int(13)] = d_785_v19_
            nw124_[int(14)] = d_785_v19_
            nw124_[int(15)] = d_785_v19_
            nw124_[int(16)] = d_785_v19_
            nw124_[int(17)] = d_785_v19_
            nw124_[int(18)] = ((d_790_v23_)[d_789_v22_] if (d_789_v22_) in (d_790_v23_) else d_785_v19_)
            nw124_[int(19)] = d_785_v19_
            nw124_[int(20)] = d_785_v19_
            nw124_[int(21)] = d_785_v19_
            nw124_[int(22)] = d_785_v19_
            nw124_[int(23)] = d_785_v19_
            nw124_[int(24)] = d_785_v19_
            d_791_v24_ = nw124_
            d_791_v24_ = d_791_v24_
        d_792_v25_: _dafny.Array
        nw125_ = _dafny.Array(None, 1)
        nw125_[int(0)] = (self).f18
        d_792_v25_ = nw125_
        index118_ = default__.safeIndex(341, (d_792_v25_).length(0))
        (d_792_v25_)[index118_] = True
        index119_ = default__.safeIndex(341, (d_792_v25_).length(0))
        (d_792_v25_)[index119_] = p0
        d_793_v26_: int
        d_793_v26_ = 89
        d_794_v27_: C0
        nw126_ = C0()
        nw126_.ctor__(default__.safeDivisionInt(d_793_v26_, d_793_v26_), d_793_v26_)
        d_794_v27_ = nw126_
        d_794_v27_ = d_794_v27_
        d_795_v28_: _dafny.Seq
        d_795_v28_ = _dafny.SeqWithoutIsStrInference([p0])
        d_796_v29_: _dafny.Map
        d_796_v29_ = _dafny.Map({len(d_795_v28_): (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]})
        if (((0) - ((d_794_v27_).f32)) not in (d_796_v29_)) or ((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]):
            (globalState).f0 = 832
            d_797_v30_: _dafny.Map
            d_797_v30_ = _dafny.Map({(d_792_v25_ if (self).f18 else d_792_v25_): ((self).f35)[default__.safeIndex(571, len((self).f35))]})
            d_797_v30_ = (d_797_v30_).set(d_792_v25_, (0) - (d_793_v26_))
            if (p0 if (self).f18 else ((d_794_v27_).f32) == (d_793_v26_)):
                d_796_v29_ = (d_796_v29_).set(d_793_v26_, (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))])
                d_795_v28_ = d_795_v28_
                (globalState).f3 = (d_794_v27_).f32
                (globalState).f3 = (d_794_v27_).f31
                index120_ = default__.safeIndex(341, (d_792_v25_).length(0))
                (d_792_v25_)[index120_] = (self).f18
            elif True:
                (globalState).f15 = (d_794_v27_).f31
                index121_ = default__.safeIndex(341, (d_792_v25_).length(0))
                (d_792_v25_)[index121_] = False
                d_798_v31_: str
                d_798_v31_ = _dafny.CodePoint('m')
                d_799_v32_: _dafny.Map
                d_799_v32_ = _dafny.Map({((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]) in (d_795_v28_): d_798_v31_})
                d_799_v32_ = (d_799_v32_).set((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], d_798_v31_)
                index122_ = default__.safeIndex(341, (d_792_v25_).length(0))
                rhs114_ = ((d_795_v28_)[default__.safeIndex(len(default__.fm43((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], p0, globalState)), len(d_795_v28_))] if (self).f18 else p0)
                rhs115_ = (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]
                lhs97_ = globalState
                lhs98_ = d_792_v25_
                lhs99_ = default__.safeIndex(341, (d_792_v25_).length(0))
                lhs97_.f2 = rhs114_
                lhs98_[lhs99_] = rhs115_
                d_800_v33_: _dafny.Array
                def lambda29_(d_801_v27_):
                    def lambda30_(d_802_i2_):
                        return (d_802_i2_) - ((d_801_v27_).f32)

                    return lambda30_

                init17_ = lambda29_(d_794_v27_)
                nw127_ = _dafny.Array(None, 19)
                for i0_17_ in range(nw127_.length(0)):
                    nw127_[i0_17_] = init17_(i0_17_)
                d_800_v33_ = nw127_
                d_803_v34_: _dafny.Map
                d_803_v34_ = _dafny.Map({d_800_v33_: d_800_v33_})
                d_803_v34_ = (d_803_v34_).set((d_800_v33_ if (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))] else d_800_v33_), d_800_v33_)
            d_804_v35_: _dafny.MultiSet
            d_804_v35_ = _dafny.MultiSet([(d_794_v27_).f31, d_793_v26_, len(d_795_v28_), (d_794_v27_).f32, d_793_v26_])
            d_805_v36_: str
            d_805_v36_ = _dafny.CodePoint('j')
            d_806_v37_: _dafny.Map
            d_806_v37_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwvvdp")): 364})
            d_807_v38_: D8
            d_807_v38_ = D8_DC21(d_805_v36_, len(d_806_v37_))
            d_808_v39_: _dafny.Array
            nw128_ = _dafny.Array(None, 24)
            nw128_[int(0)] = default__.fm1(((d_804_v35_)[(d_794_v27_).f32] if ((d_794_v27_).f32) in (d_804_v35_) else (d_794_v27_).f31), (self).fm2((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], globalState), (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], globalState)
            nw128_[int(1)] = len((d_795_v28_) + (d_795_v28_))
            nw128_[int(2)] = ((d_794_v27_).f31) - (len((self).f19))
            nw128_[int(3)] = default__.safeDivisionInt((d_794_v27_).f31, (0) - (len((self).f35)))
            nw128_[int(4)] = (d_794_v27_).f31
            nw128_[int(5)] = (d_794_v27_).f31
            nw128_[int(6)] = (d_794_v27_).f31
            nw128_[int(7)] = (d_794_v27_).f31
            nw128_[int(8)] = (d_794_v27_).f32
            nw128_[int(9)] = (d_807_v38_).cf42
            nw128_[int(10)] = (d_794_v27_).f31
            nw128_[int(11)] = (d_794_v27_).f31
            nw128_[int(12)] = (d_794_v27_).f32
            nw128_[int(13)] = (d_794_v27_).f31
            nw128_[int(14)] = (d_793_v26_) + ((d_794_v27_).f31)
            nw128_[int(15)] = (d_794_v27_).f31
            nw128_[int(16)] = default__.safeDivisionInt(d_793_v26_, 986)
            nw128_[int(17)] = ((d_794_v27_).f31) + ((0) - ((d_794_v27_).f32))
            nw128_[int(18)] = (d_794_v27_).f31
            nw128_[int(19)] = (d_794_v27_).f32
            nw128_[int(20)] = (d_794_v27_).f31
            nw128_[int(21)] = 816
            nw128_[int(22)] = len((self).f19)
            nw128_[int(23)] = (d_794_v27_).f31
            d_808_v39_ = nw128_
            index123_ = default__.safeIndex(892, (d_808_v39_).length(0))
            (d_808_v39_)[index123_] = len(d_796_v29_)
            index124_ = default__.safeIndex(892, (d_808_v39_).length(0))
            (d_808_v39_)[index124_] = (d_794_v27_).f32
            d_809_v40_: D1
            d_809_v40_ = D1_DC4((self).f18)
            source7_ = d_809_v40_
            if source7_.is_DC3:
                d_810___mcc_h0_ = source7_.cf7
                d_811___mcc_h1_ = source7_.cf8
                d_812___mcc_h2_ = source7_.cf9
                d_813___mcc_h3_ = source7_.cf10
                d_814___mcc_h4_ = source7_.cf11
                d_815_cf11_ = d_814___mcc_h4_
                d_816_cf10_ = d_813___mcc_h3_
                d_817_cf9_ = d_812___mcc_h2_
                d_818_cf8_ = d_811___mcc_h1_
                d_819_cf7_ = d_810___mcc_h0_
                d_820_v41_: T0
                nw129_ = C1()
                nw129_.ctor__(not((self).f18))
                d_820_v41_ = nw129_
                d_821_v42_: D1
                d_821_v42_ = D1_DC2(d_820_v41_)
                d_822_v43_: _dafny.Map
                d_822_v43_ = _dafny.Map({d_821_v42_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nr"))})
                d_823_v44_: D3
                d_823_v44_ = D3_DC13((d_794_v27_).f32, (self).f19, d_819_cf7_, (d_794_v27_).f31, d_805_v36_)
                d_824_v45_: _dafny.Map
                d_824_v45_ = _dafny.Map({(self).f18: ((d_822_v43_)[D1_DC2(d_820_v41_)] if (D1_DC2(d_820_v41_)) in (d_822_v43_) else (d_823_v44_).cf31)})
                d_824_v45_ = (d_824_v45_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imyjlnfkg"))) <= ((self).f19), (self).f19)
                d_766_v1_ = d_766_v1_
                d_825_v46_: _dafny.Set
                d_825_v46_ = _dafny.Set({d_805_v36_, d_805_v36_})
                d_826_v47_: D2
                d_826_v47_ = D2_DC7((0) - (d_819_cf7_), (d_820_v41_).f18, len(d_825_v46_))
                d_827_v48_: D2
                d_827_v48_ = D2_DC9(d_826_v47_)
                d_828_v49_: D2
                d_828_v49_ = D2_DC9(d_827_v48_)
                d_829_v50_: _dafny.Map
                d_829_v50_ = _dafny.Map({d_807_v38_: d_828_v49_})
                d_830_v52_: _dafny.Seq
                d_830_v52_ = _dafny.SeqWithoutIsStrInference([d_807_v38_])
                d_831_v54_: C1
                nw130_ = C1()
                nw130_.ctor__((self).f18)
                d_831_v54_ = nw130_
                d_832_v55_: _dafny.Map
                d_832_v55_ = _dafny.Map({d_831_v54_: False})
                d_833_v56_: D11
                d_833_v56_ = D11_DC27(d_832_v55_)
                d_834_v57_: _dafny.Array
                def lambda31_(d_835_cf9_, d_836_cf11_):
                    def lambda32_(d_837_i3_):
                        return _dafny.Set({d_835_cf9_, d_836_cf11_, len((self).f19)})

                    return lambda32_

                init18_ = lambda31_(d_817_cf9_, d_815_cf11_)
                nw131_ = _dafny.Array(None, 16)
                for i0_18_ in range(nw131_.length(0)):
                    nw131_[i0_18_] = init18_(i0_18_)
                d_834_v57_ = nw131_
                d_838_v58_: _dafny.Map
                d_838_v58_ = _dafny.Map({d_834_v57_: (d_829_v50_).set(d_807_v38_, d_828_v49_)})
                d_839_v59_: _dafny.Array
                nw132_ = _dafny.Array(None, 13)
                nw132_[int(0)] = d_829_v50_
                def iife82_():
                    coll26_ = _dafny.Map()
                    compr_26_: D8
                    for compr_26_ in (d_830_v52_).Elements:
                        d_840_v51_: D8 = compr_26_
                        if (d_840_v51_) in (d_830_v52_):
                            coll26_[d_840_v51_] = d_828_v49_
                    return _dafny.Map(coll26_)
                nw132_[int(1)] = iife82_()
                
                nw132_[int(2)] = (_dafny.Map({d_807_v38_: d_828_v49_})) | (d_829_v50_)
                nw132_[int(3)] = d_829_v50_
                nw132_[int(4)] = d_829_v50_
                nw132_[int(5)] = d_829_v50_
                nw132_[int(6)] = d_829_v50_
                def iife83_():
                    coll27_ = _dafny.Map()
                    compr_27_: D8
                    for compr_27_ in ((_dafny.SeqWithoutIsStrInference([d_807_v38_])).set(default__.safeIndex(len((d_833_v56_).cf50), len(_dafny.SeqWithoutIsStrInference([d_807_v38_]))), d_807_v38_)).Elements:
                        d_841_v53_: D8 = compr_27_
                        if (d_841_v53_) in ((_dafny.SeqWithoutIsStrInference([d_807_v38_])).set(default__.safeIndex(len((d_833_v56_).cf50), len(_dafny.SeqWithoutIsStrInference([d_807_v38_]))), d_807_v38_)):
                            coll27_[d_841_v53_] = d_828_v49_
                    return _dafny.Map(coll27_)
                nw132_[int(7)] = iife83_()
                
                nw132_[int(8)] = d_829_v50_
                nw132_[int(9)] = d_829_v50_
                nw132_[int(10)] = (((d_838_v58_)[d_834_v57_] if (d_834_v57_) in (d_838_v58_) else _dafny.Map({d_807_v38_: d_828_v49_}))).set(d_807_v38_, d_828_v49_)
                nw132_[int(11)] = d_829_v50_
                nw132_[int(12)] = _dafny.Map({d_807_v38_: d_828_v49_})
                d_839_v59_ = nw132_
                index125_ = default__.safeIndex(882, (d_839_v59_).length(0))
                (d_839_v59_)[index125_] = (d_829_v50_) | (_dafny.Map({d_807_v38_: d_828_v49_}))
                index126_ = default__.safeIndex(882, (d_839_v59_).length(0))
                (d_839_v59_)[index126_] = (d_829_v50_) | (d_829_v50_)
                d_842_v60_: _dafny.Map
                d_842_v60_ = _dafny.Map({d_792_v25_: (self).f19})
                d_842_v60_ = (d_842_v60_).set(d_792_v25_, (self).f19)
            elif source7_.is_DC4:
                d_843___mcc_h5_ = source7_.cf12
                d_844_cf12_ = d_843___mcc_h5_
                d_805_v36_ = d_805_v36_
                (globalState).f2 = (self).f18
                d_845_v61_: _dafny.Map
                d_845_v61_ = _dafny.Map({(self).f35: d_808_v39_})
                d_846_v62_: _dafny.Map
                d_846_v62_ = _dafny.Map({d_793_v26_: d_845_v61_})
                d_847_v63_: _dafny.Map
                d_847_v63_ = _dafny.Map({(len((_dafny.Map({d_793_v26_: (d_794_v27_).f32})).set((d_794_v27_).f32, (d_808_v39_)[default__.safeIndex(892, (d_808_v39_).length(0))]))) < ((d_794_v27_).f32): (((d_846_v62_)[len((self).f19)] if (len((self).f19)) in (d_846_v62_) else d_845_v61_)) | (_dafny.Map({(self).f35: d_808_v39_}))})
                d_847_v63_ = (d_847_v63_).set((self).f18, (d_845_v61_) | (d_845_v61_))
                d_848_v64_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_848_v64_ = nw133_
                index127_ = default__.safeIndex(370, (d_848_v64_).length(0))
                (d_848_v64_)[index127_] = d_808_v39_
                index128_ = default__.safeIndex(370, (d_848_v64_).length(0))
                (d_848_v64_)[index128_] = d_808_v39_
            elif source7_.is_DC2:
                d_849___mcc_h6_ = source7_.cf6
                d_850_cf6_ = d_849___mcc_h6_
                d_851_v65_: _dafny.Array
                nw134_ = _dafny.Array(_dafny.CodePoint('D'), 14)
                d_851_v65_ = nw134_
                index129_ = default__.safeIndex(533, (d_851_v65_).length(0))
                (d_851_v65_)[index129_] = d_805_v36_
                d_852_v66_: _dafny.MultiSet
                d_852_v66_ = _dafny.MultiSet([not(not((self).f18)), p0, (d_793_v26_) > ((d_794_v27_).f31)])
                index130_ = default__.safeIndex(533, (d_851_v65_).length(0))
                index131_ = default__.safeIndex(892, (d_808_v39_).length(0))
                rhs116_ = d_805_v36_
                rhs117_ = (_dafny.MultiSet(d_795_v28_)) | (d_852_v66_)
                rhs118_ = (d_794_v27_).f32
                lhs100_ = d_851_v65_
                lhs101_ = default__.safeIndex(533, (d_851_v65_).length(0))
                lhs102_ = d_808_v39_
                lhs103_ = default__.safeIndex(892, (d_808_v39_).length(0))
                lhs100_[lhs101_] = rhs116_
                d_852_v66_ = rhs117_
                lhs102_[lhs103_] = rhs118_
                (globalState).f3 = ((d_852_v66_)[True] if (True) in (d_852_v66_) else default__.safeDivisionInt(-414, (d_794_v27_).f32))
                index132_ = default__.safeIndex(341, (d_792_v25_).length(0))
                (d_792_v25_)[index132_] = (self).f18
                r3 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyagbxxr"))) + ((self).f19)) + ((self).f19)
            elif True:
                d_853___mcc_h7_ = source7_.cf13
                d_854_cf13_ = d_853___mcc_h7_
                d_855_v67_: _dafny.Map
                d_855_v67_ = _dafny.Map({(d_794_v27_).f31: d_795_v28_})
                d_856_v68_: _dafny.Map
                d_856_v68_ = _dafny.Map({(d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]: d_795_v28_})
                d_857_v69_: D12
                d_857_v69_ = D12_DC30(d_795_v28_)
                d_858_v70_: _dafny.Array
                nw135_ = _dafny.Array(None, 25)
                nw135_[int(0)] = d_795_v28_
                nw135_[int(1)] = d_795_v28_
                nw135_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f18])
                nw135_[int(3)] = d_795_v28_
                nw135_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]])
                nw135_[int(5)] = ((d_855_v67_)[(d_794_v27_).f32] if ((d_794_v27_).f32) in (d_855_v67_) else d_795_v28_)
                nw135_[int(6)] = d_795_v28_
                nw135_[int(7)] = default__.fm44((d_794_v27_).f31, (self).f18, d_805_v36_, (_dafny.MultiSet([default__.fm1((d_794_v27_).f32, (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], p0, globalState), (d_794_v27_).f31, (0) - ((d_794_v27_).f31)])).cardinality, globalState)
                nw135_[int(8)] = (d_795_v28_) + (d_795_v28_)
                nw135_[int(9)] = d_795_v28_
                nw135_[int(10)] = d_795_v28_
                nw135_[int(11)] = default__.fm44((d_794_v27_).f31, (self).f18, d_805_v36_, d_793_v26_, globalState)
                nw135_[int(12)] = (d_795_v28_).set(default__.safeIndex((d_808_v39_)[default__.safeIndex(892, (d_808_v39_).length(0))], len(d_795_v28_)), (d_795_v28_)[default__.safeIndex(d_793_v26_, len(d_795_v28_))])
                nw135_[int(13)] = (d_795_v28_) + (d_795_v28_)
                nw135_[int(14)] = (d_795_v28_) + (d_795_v28_)
                nw135_[int(15)] = (d_795_v28_).set(default__.safeIndex(-627, len(d_795_v28_)), not((self).f18))
                nw135_[int(16)] = (d_795_v28_) + (d_795_v28_)
                nw135_[int(17)] = (((d_856_v68_)[(self).f18] if ((self).f18) in (d_856_v68_) else d_795_v28_)) + (d_795_v28_)
                nw135_[int(18)] = d_795_v28_
                nw135_[int(19)] = (d_795_v28_) + (d_795_v28_)
                nw135_[int(20)] = (d_857_v69_).cf53
                nw135_[int(21)] = d_795_v28_
                nw135_[int(22)] = d_795_v28_
                nw135_[int(23)] = (d_795_v28_) + (d_795_v28_)
                nw135_[int(24)] = d_795_v28_
                d_858_v70_ = nw135_
                d_859_v71_: _dafny.Set
                d_859_v71_ = _dafny.Set({(self).f19})
                d_860_v73_: _dafny.Array
                nw136_ = _dafny.Array(None, 21)
                nw136_[int(0)] = _dafny.Set({(self).f19})
                nw136_[int(1)] = d_859_v71_
                nw136_[int(2)] = d_859_v71_
                nw136_[int(3)] = (d_859_v71_) - (d_859_v71_)
                nw136_[int(4)] = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dr"))})
                nw136_[int(5)] = d_859_v71_
                nw136_[int(6)] = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_805_v36_ for d_861_i4_ in range(default__.abs(235))]), (self).f19})
                nw136_[int(7)] = default__.fm45(globalState)
                nw136_[int(8)] = d_859_v71_
                nw136_[int(9)] = d_859_v71_
                def iife84_():
                    coll28_ = _dafny.Set()
                    compr_28_: _dafny.Seq
                    for compr_28_ in (_dafny.Map({(self).f19: (d_794_v27_).f32})).keys.Elements:
                        d_862_v72_: _dafny.Seq = compr_28_
                        if (d_862_v72_) in (_dafny.Map({(self).f19: (d_794_v27_).f32})):
                            coll28_ = coll28_.union(_dafny.Set([d_862_v72_]))
                    return _dafny.Set(coll28_)
                nw136_[int(10)] = (default__.fm45(globalState)) - (iife84_()
                )
                nw136_[int(11)] = _dafny.Set({(self).f19})
                nw136_[int(12)] = d_859_v71_
                nw136_[int(13)] = d_859_v71_
                nw136_[int(14)] = d_859_v71_
                nw136_[int(15)] = (d_859_v71_).intersection(_dafny.Set({(self).f19}))
                nw136_[int(16)] = d_859_v71_
                nw136_[int(17)] = d_859_v71_
                nw136_[int(18)] = d_859_v71_
                nw136_[int(19)] = d_859_v71_
                nw136_[int(20)] = d_859_v71_
                d_860_v73_ = nw136_
                index133_ = default__.safeIndex(902, (d_860_v73_).length(0))
                (d_860_v73_)[index133_] = d_859_v71_
                d_863_v74_: _dafny.Map
                d_863_v74_ = _dafny.Map({d_792_v25_: False})
                d_864_v75_: D13
                d_864_v75_ = D13_DC32(d_863_v74_)
                d_865_v76_: _dafny.Map
                d_865_v76_ = _dafny.Map({p0: len((d_863_v74_) | ((d_864_v75_).cf54))})
                index134_ = default__.safeIndex(902, (d_860_v73_).length(0))
                rhs119_ = d_765_v0_
                rhs120_ = d_858_v70_
                rhs121_ = d_859_v71_
                rhs122_ = (d_794_v27_).f31
                rhs123_ = (d_865_v76_).set((self).fm2((self).fm2(False, (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], p0, globalState), (self).fm2((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], not(p0), p0, globalState), True, globalState), (d_794_v27_).f31)
                lhs104_ = d_860_v73_
                lhs105_ = default__.safeIndex(902, (d_860_v73_).length(0))
                lhs106_ = globalState
                d_765_v0_ = rhs119_
                d_858_v70_ = rhs120_
                lhs104_[lhs105_] = rhs121_
                lhs106_.f3 = rhs122_
                d_865_v76_ = rhs123_
                d_805_v36_ = d_805_v36_
                d_866_v77_: _dafny.Array
                nw137_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_866_v77_ = nw137_
                index135_ = default__.safeIndex(938, (d_866_v77_).length(0))
                (d_866_v77_)[index135_] = d_792_v25_
                d_867_v78_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Map({}), 27)
                d_867_v78_ = nw138_
                d_868_v79_: _dafny.Map
                d_868_v79_ = _dafny.Map({not((self).f18): d_867_v78_})
                index136_ = default__.safeIndex(938, (d_866_v77_).length(0))
                index137_ = default__.safeIndex(892, (d_808_v39_).length(0))
                rhs124_ = (d_792_v25_ if (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))] else d_792_v25_)
                rhs125_ = 335
                rhs126_ = (self).f18
                rhs127_ = d_868_v79_
                lhs107_ = d_866_v77_
                lhs108_ = default__.safeIndex(938, (d_866_v77_).length(0))
                lhs109_ = d_808_v39_
                lhs110_ = default__.safeIndex(892, (d_808_v39_).length(0))
                lhs111_ = globalState
                lhs107_[lhs108_] = rhs124_
                lhs109_[lhs110_] = rhs125_
                lhs111_.f2 = rhs126_
                d_868_v79_ = rhs127_
                d_869_v80_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.Seq({}), 23)
                d_869_v80_ = nw139_
                index138_ = default__.safeIndex(751, (d_869_v80_).length(0))
                (d_869_v80_)[index138_] = (self).f35
                index139_ = default__.safeIndex(751, (d_869_v80_).length(0))
                rhs128_ = not (not((d_795_v28_)[default__.safeIndex(d_793_v26_, len(d_795_v28_))])) or ((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))])
                rhs129_ = p0
                rhs130_ = ((self).f35) + ((self).f35)
                rhs131_ = _dafny.SeqWithoutIsStrInference([((self).f19)[default__.safeIndex((d_794_v27_).f31, len((self).f19))] for d_870_i5_ in range(default__.abs(593))])
                lhs112_ = globalState
                lhs113_ = globalState
                lhs114_ = d_869_v80_
                lhs115_ = default__.safeIndex(751, (d_869_v80_).length(0))
                lhs116_ = globalState
                lhs112_.f2 = rhs128_
                lhs113_.f2 = rhs129_
                lhs114_[lhs115_] = rhs130_
                lhs116_.f9 = rhs131_
        elif True:
            (globalState).f15 = (0) - (d_793_v26_)
            if (self).fm2(p0, (self).f18, (self).f18, globalState):
                (globalState).f15 = (d_794_v27_).f32
                d_871_v81_: _dafny.Map
                d_871_v81_ = _dafny.Map({d_792_v25_: (0) - ((d_794_v27_).f31)})
                d_872_v82_: _dafny.Map
                d_872_v82_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_873_i6_ in range(default__.abs(178))]): (d_794_v27_).f31})
                d_871_v81_ = (d_871_v81_).set(d_792_v25_, len((d_872_v82_) | (d_872_v82_)))
                d_874_v83_: C1
                nw140_ = C1()
                nw140_.ctor__(p0)
                d_874_v83_ = nw140_
                d_875_v84_: _dafny.Array
                nw141_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_875_v84_ = nw141_
                d_875_v84_ = d_875_v84_
                index140_ = default__.safeIndex(341, (d_792_v25_).length(0))
                (d_792_v25_)[index140_] = p0
            elif True:
                nw142_ = _dafny.Array(False, 29)
                d_792_v25_ = nw142_
                index141_ = default__.safeIndex(341, (d_792_v25_).length(0))
                (d_792_v25_)[index141_] = ((self).f18 if False else (self).f18)
                rhs132_ = ((d_794_v27_).f31) == ((d_794_v27_).f32)
                rhs133_ = (self).fm38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsjphe")), globalState)
                lhs117_ = globalState
                lhs118_ = globalState
                lhs117_.f2 = rhs132_
                lhs118_.f15 = rhs133_
                d_876_v85_: str
                d_876_v85_ = _dafny.CodePoint('w')
                d_877_v86_: _dafny.MultiSet
                d_877_v86_ = _dafny.MultiSet([d_876_v85_, default__.fm46(globalState), ((self).f19)[default__.safeIndex((d_794_v27_).f32, len((self).f19))], d_876_v85_])
                d_877_v86_ = ((d_877_v86_) - (_dafny.MultiSet((self).f19))) | (d_877_v86_)
                (globalState).f2 = p0
            (globalState).f2 = (d_795_v28_)[default__.safeIndex((d_794_v27_).f32, len(d_795_v28_))]
            if not((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]):
                d_878_v87_: C1
                nw143_ = C1()
                nw143_.ctor__((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))])
                d_878_v87_ = nw143_
                (globalState).f2 = (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))]
                (globalState).f3 = (d_794_v27_).f32
                (d_878_v87_).m1((self).f19, default__.safeDivisionInt(default__.fm1((d_794_v27_).f32, (self).f18, False, globalState), (d_794_v27_).f31), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkpnyc")), globalState)
                d_879_v88_: _dafny.Set
                d_879_v88_ = _dafny.Set({(0) - ((d_794_v27_).f31)})
                d_880_v89_: _dafny.Map
                d_880_v89_ = _dafny.Map({d_879_v88_: p0})
                d_880_v89_ = (d_880_v89_).set(d_879_v88_, False)
            elif True:
                (globalState).f2 = (self).f18
                r1 = default__.safeDivisionInt(default__.fm1((d_794_v27_).f31, (self).fm3(len(d_795_v28_), globalState), (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], globalState), (d_794_v27_).f31)
                d_881_v90_: _dafny.Array
                nw144_ = _dafny.Array(int(0), 3)
                d_881_v90_ = nw144_
                index142_ = default__.safeIndex(22, (d_881_v90_).length(0))
                (d_881_v90_)[index142_] = 237
                d_882_v91_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.Set({}), 8)
                d_882_v91_ = nw145_
                d_883_v92_: C1
                nw146_ = C1()
                nw146_.ctor__((self).f18)
                d_883_v92_ = nw146_
                d_884_v93_: _dafny.Map
                d_884_v93_ = _dafny.Map({d_883_v92_: not(p0)})
                d_885_v94_: D11
                d_885_v94_ = D11_DC27(d_884_v93_)
                d_886_v95_: _dafny.Set
                d_886_v95_ = _dafny.Set({d_885_v94_})
                index143_ = default__.safeIndex(566, (d_882_v91_).length(0))
                (d_882_v91_)[index143_] = d_886_v95_
                index144_ = default__.safeIndex(22, (d_881_v90_).length(0))
                index145_ = default__.safeIndex(566, (d_882_v91_).length(0))
                rhs134_ = default__.safeDivisionInt(default__.safeDivisionInt((d_794_v27_).f32, d_793_v26_), len((self).f19))
                rhs135_ = default__.fm1(d_793_v26_, (self).fm2(False, (self).f18, p0, globalState), (self).f18, globalState)
                rhs136_ = (d_886_v95_) | (d_886_v95_)
                rhs137_ = (self).f18
                lhs119_ = globalState
                lhs120_ = d_881_v90_
                lhs121_ = default__.safeIndex(22, (d_881_v90_).length(0))
                lhs122_ = d_882_v91_
                lhs123_ = default__.safeIndex(566, (d_882_v91_).length(0))
                lhs124_ = globalState
                lhs119_.f15 = rhs134_
                lhs120_[lhs121_] = rhs135_
                lhs122_[lhs123_] = rhs136_
                lhs124_.f2 = rhs137_
                (globalState).f12 = (self).f19
                d_887_v96_: str
                d_887_v96_ = _dafny.CodePoint('v')
                d_887_v96_ = d_887_v96_
            d_888_v97_: C2
            nw147_ = C2()
            nw147_.ctor__((d_794_v27_).f32, (self).f35, p0)
            d_888_v97_ = nw147_
        d_889_v98_: C2
        nw148_ = C2()
        nw148_.ctor__(len((self).f19), (self).f35, (self).f18)
        d_889_v98_ = nw148_
        d_890_v99_: D13
        d_890_v99_ = D13_DC34(d_792_v25_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgpak")), (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], not((self).f18))
        d_891_v100_: _dafny.Seq
        d_891_v100_ = _dafny.SeqWithoutIsStrInference([d_890_v99_, D13_DC34((d_792_v25_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eivjeqg")), p0, p0), d_890_v99_])
        d_891_v100_ = (d_891_v100_) + (d_891_v100_)
        r0 = (d_766_v1_).intersection(d_766_v1_)
        d_892_v101_: D2
        d_892_v101_ = D2_DC8((d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))], (d_792_v25_)[default__.safeIndex(341, (d_792_v25_).length(0))])
        d_893_v102_: D2
        d_893_v102_ = D2_DC9(d_892_v101_)
        pat_let_tv15_ = d_889_v98_
        pat_let_tv16_ = d_794_v27_
        pat_let_tv17_ = d_889_v98_
        pat_let_tv18_ = d_889_v98_
        pat_let_tv19_ = d_794_v27_
        def lambda33_(source8_):
            if source8_.is_DC7:
                d_894___mcc_h8_ = source8_.cf15
                d_895___mcc_h9_ = source8_.cf16
                d_896___mcc_h10_ = source8_.cf17
                d_897_cf17_ = d_896___mcc_h10_
                d_898_cf16_ = d_895___mcc_h9_
                d_899_cf15_ = d_894___mcc_h8_
                return default__.safeDivisionInt((0) - (pat_let_tv15_.f33), (pat_let_tv16_).f32)
            elif source8_.is_DC8:
                d_900___mcc_h11_ = source8_.cf18
                d_901___mcc_h12_ = source8_.cf19
                d_902___mcc_h13_ = source8_.cf20
                d_903_cf20_ = d_902___mcc_h13_
                d_904_cf19_ = d_901___mcc_h12_
                d_905_cf18_ = d_900___mcc_h11_
                return pat_let_tv17_.f33
            elif source8_.is_DC6:
                d_906___mcc_h14_ = source8_.cf14
                d_907_cf14_ = d_906___mcc_h14_
                return pat_let_tv18_.f33
            elif True:
                d_908___mcc_h15_ = source8_.cf21
                d_909_cf21_ = d_908___mcc_h15_
                return (pat_let_tv19_).f31

        r1 = lambda33_((d_893_v102_ if not(True) else d_893_v102_))
        r2 = (0) - (d_793_v26_)
        r3 = ((self).f19) + ((((self).f19) + ((self).f19)).set(default__.safeIndex((d_794_v27_).f32, len(((self).f19) + ((self).f19))), _dafny.CodePoint('j')))
        return r0, r1, r2, r3

    @property
    def f35(self):
        return self._f35

class C4(T0, T1):
    def  __init__(self):
        self._f18: bool = False
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f20: _dafny.Seq = _dafny.Seq({})
        self._f29: str = _dafny.CodePoint('D')
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f29, f30, f18, f19, f20):
        (self)._f29 = f29
        (self)._f30 = f30
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20

    def fm2(self, p0, p1, p2, globalState):
        return not(((self).f30) == (((0) - ((self).f30)) - ((self).f30)))

    def fm3(self, p0, globalState):
        return (self).f18

    def fm17(self, p0, p1, globalState):
        return (-617) - ((0) - ((self).f30))

    def fm18(self, globalState):
        return (self).f30

    def m1(self, p0, p1, p2, globalState):
        if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcybxl")))) < ((self).f30):
            d_910_v0_: _dafny.Map
            d_910_v0_ = _dafny.Map({(self).f30: (self).f18})
            d_911_v1_: _dafny.Set
            d_911_v1_ = _dafny.Set({True})
            d_912_v2_: _dafny.Set
            d_912_v2_ = _dafny.Set({d_911_v1_, d_911_v1_, d_911_v1_, d_911_v1_})
            d_913_v3_: _dafny.Map
            d_913_v3_ = _dafny.Map({p1: (p1) * (len(d_912_v2_))})
            rhs138_ = (self).f18
            rhs139_ = (self).f18
            rhs140_ = d_910_v0_
            rhs141_ = ((_dafny.Map({p1: (self).f30})) | ((d_913_v3_).set(p1, len(d_910_v0_)))) | (d_913_v3_)
            lhs125_ = globalState
            lhs126_ = globalState
            lhs125_.f2 = rhs138_
            lhs126_.f2 = rhs139_
            d_910_v0_ = rhs140_
            d_913_v3_ = rhs141_
            (globalState).f2 = not((self).f18)
            d_914_v4_: _dafny.Array
            nw149_ = _dafny.Array(int(0), 4)
            d_914_v4_ = nw149_
            index146_ = default__.safeIndex(441, (d_914_v4_).length(0))
            (d_914_v4_)[index146_] = p1
            index147_ = default__.safeIndex(441, (d_914_v4_).length(0))
            (d_914_v4_)[index147_] = p1
            d_915_v5_: _dafny.Seq
            d_915_v5_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, (self).f18])
            d_916_v6_: C0
            nw150_ = C0()
            nw150_.ctor__(len(d_915_v5_), (self).f30)
            d_916_v6_ = nw150_
            index148_ = default__.safeIndex(441, (d_914_v4_).length(0))
            (d_914_v4_)[index148_] = (d_914_v4_)[default__.safeIndex(441, (d_914_v4_).length(0))]
        elif True:
            d_917_v7_: _dafny.Seq
            d_917_v7_ = _dafny.SeqWithoutIsStrInference([p1])
            d_918_v8_: _dafny.Map
            d_918_v8_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f30])) + (_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30])): ((self).f30) in (d_917_v7_)})
            d_919_v9_: _dafny.Map
            d_919_v9_ = _dafny.Map({(self).f30: (self).f18})
            d_918_v8_ = (d_918_v8_).set(_dafny.SeqWithoutIsStrInference([p1, (self).f30]), not(((d_919_v9_)[(self).f30] if ((self).f30) in (d_919_v9_) else (self).f18)))
            d_920_v10_: D0
            d_920_v10_ = D0_DC1((self).f18, p1, _dafny.SeqWithoutIsStrInference([(self).f29, _dafny.CodePoint('x')]), not((self).f18), (self).f30)
            d_921_v11_: _dafny.Map
            d_921_v11_ = _dafny.Map({d_920_v10_: (self).f18})
            d_922_v12_: _dafny.MultiSet
            d_922_v12_ = _dafny.MultiSet([(self).f18, True, (self).f18, (self).f18])
            d_921_v11_ = (d_921_v11_).set(d_920_v10_, (d_922_v12_).issubset(d_922_v12_))
            d_923_v13_: D2
            d_923_v13_ = D2_DC7(p1, (self).f18, (self).f30)
            if (d_923_v13_).cf16:
                (globalState).f3 = default__.safeModuloInt(((self).f30) + (p1), ((self).f30 if (self).f18 else p1))
                d_924_v14_: _dafny.Set
                d_924_v14_ = _dafny.Set({(self).f30})
                d_925_v15_: D2
                d_925_v15_ = D2_DC8((p1) not in (_dafny.SeqWithoutIsStrInference([p1, len(d_924_v14_), 810])), True, (p0) <= (p0))
                d_925_v15_ = d_925_v15_
                d_926_v16_: _dafny.Array
                nw151_ = _dafny.Array(None, 28)
                nw151_[int(0)] = (d_925_v15_).cf19
                nw151_[int(1)] = (self).f18
                nw151_[int(2)] = False
                nw151_[int(3)] = (self).f18
                nw151_[int(4)] = ((0) - ((0) - (p1))) == ((self).f30)
                nw151_[int(5)] = (self).f18
                nw151_[int(6)] = (self).f18
                nw151_[int(7)] = (self).f18
                nw151_[int(8)] = (self).f18
                nw151_[int(9)] = (self).f18
                nw151_[int(10)] = True
                nw151_[int(11)] = True
                nw151_[int(12)] = (self).f18
                nw151_[int(13)] = (p1) not in (d_917_v7_)
                nw151_[int(14)] = not(not((self).f18))
                nw151_[int(15)] = not ((self).f18) or ((self).f18)
                nw151_[int(16)] = (self).f18
                nw151_[int(17)] = (self).f18
                nw151_[int(18)] = (self).f18
                nw151_[int(19)] = False
                nw151_[int(20)] = not((self).f18)
                nw151_[int(21)] = (self).f18
                nw151_[int(22)] = True
                nw151_[int(23)] = (self).f18
                nw151_[int(24)] = (self).f18
                nw151_[int(25)] = (self).f18
                nw151_[int(26)] = (self).f18
                nw151_[int(27)] = False
                d_926_v16_ = nw151_
                index149_ = default__.safeIndex(98, (d_926_v16_).length(0))
                (d_926_v16_)[index149_] = not((self).f18)
                index150_ = default__.safeIndex(98, (d_926_v16_).length(0))
                (d_926_v16_)[index150_] = True
                d_927_v17_: _dafny.Array
                def lambda34_(d_928_v14_):
                    def lambda35_(d_929_i0_):
                        return (d_929_i0_) + (len(d_928_v14_))

                    return lambda35_

                init19_ = lambda34_(d_924_v14_)
                nw152_ = _dafny.Array(None, 10)
                for i0_19_ in range(nw152_.length(0)):
                    nw152_[i0_19_] = init19_(i0_19_)
                d_927_v17_ = nw152_
                index151_ = default__.safeIndex(726, (d_927_v17_).length(0))
                (d_927_v17_)[index151_] = len(p2)
                index152_ = default__.safeIndex(726, (d_927_v17_).length(0))
                (d_927_v17_)[index152_] = (self).f30
                d_930_v18_: _dafny.Map
                d_930_v18_ = _dafny.Map({d_927_v17_: (d_926_v16_)[default__.safeIndex(98, (d_926_v16_).length(0))]})
                d_931_v19_: _dafny.Seq
                d_931_v19_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                d_932_v20_: D1
                d_932_v20_ = D1_DC3(309, len(d_931_v19_), -878, p1, (d_927_v17_)[default__.safeIndex(726, (d_927_v17_).length(0))])
                index153_ = default__.safeIndex(98, (d_926_v16_).length(0))
                rhs142_ = ((d_930_v18_)[d_927_v17_] if (d_927_v17_) in (d_930_v18_) else (d_926_v16_)[default__.safeIndex(98, (d_926_v16_).length(0))])
                rhs143_ = (self).f18
                rhs144_ = True
                rhs145_ = (0) - ((0) - ((d_932_v20_).cf8))
                lhs127_ = d_926_v16_
                lhs128_ = default__.safeIndex(98, (d_926_v16_).length(0))
                lhs129_ = globalState
                lhs130_ = globalState
                lhs131_ = globalState
                lhs127_[lhs128_] = rhs142_
                lhs129_.f2 = rhs143_
                lhs130_.f2 = rhs144_
                lhs131_.f0 = rhs145_
            elif True:
                d_933_v21_: _dafny.Map
                d_933_v21_ = _dafny.Map({25: (d_922_v12_).cardinality})
                d_934_v22_: _dafny.Set
                d_934_v22_ = _dafny.Set({_dafny.Map({len(d_917_v7_): default__.fm1(len(d_917_v7_), (self).f18, (self).f18, globalState)}), d_933_v21_, d_933_v21_})
                d_935_v23_: _dafny.Map
                d_935_v23_ = _dafny.Map({True: d_934_v22_})
                d_935_v23_ = (d_935_v23_).set((self).f18, d_934_v22_)
                (globalState).f15 = (self).f30
                d_936_v24_: T0
                nw153_ = C1()
                nw153_.ctor__(True)
                d_936_v24_ = nw153_
                d_937_v25_: _dafny.Seq
                d_937_v25_ = _dafny.SeqWithoutIsStrInference([d_922_v12_])
                d_938_v27_: _dafny.Seq
                d_938_v27_ = _dafny.SeqWithoutIsStrInference([p2, (self).f19])
                d_939_v28_: _dafny.Map
                d_939_v28_ = _dafny.Map({(self).f19: (d_936_v24_).f18})
                def iife85_():
                    coll29_ = _dafny.Map()
                    compr_29_: _dafny.Seq
                    for compr_29_ in (d_938_v27_).Elements:
                        d_940_v26_: _dafny.Seq = compr_29_
                        if (d_940_v26_) in (d_938_v27_):
                            coll29_[d_940_v26_] = (d_936_v24_).f18
                    return _dafny.Map(coll29_)
                rhs146_ = default__.fm30(d_937_v25_, (iife85_()
                ) | (d_939_v28_), globalState)
                rhs147_ = d_936_v24_
                d_933_v21_ = rhs146_
                d_936_v24_ = rhs147_
                (globalState).f2 = (self).f18
                (globalState).f15 = (self).f30
            if (self).f18:
                d_941_v29_: _dafny.Map
                d_941_v29_ = _dafny.Map({(self).f18: (self).f18})
                d_942_v30_: _dafny.Array
                nw154_ = _dafny.Array(False, 27)
                d_942_v30_ = nw154_
                d_943_v31_: D2
                d_943_v31_ = D2_DC9(D2_DC8((self).f18, (self).f18, ((d_919_v9_)[(self).f30] if ((self).f30) in (d_919_v9_) else True)))
                d_944_v32_: _dafny.Map
                d_944_v32_ = _dafny.Map({(self).f30: (self).f29})
                d_945_v33_: _dafny.Array
                nw155_ = _dafny.Array(None, 22)
                nw155_[int(0)] = (self).f29
                nw155_[int(1)] = (self).f29
                nw155_[int(2)] = (self).f29
                nw155_[int(3)] = (self).f29
                nw155_[int(4)] = (self).f29
                nw155_[int(5)] = (self).f29
                nw155_[int(6)] = (self).f29
                nw155_[int(7)] = ((d_944_v32_)[len(p2)] if (len(p2)) in (d_944_v32_) else _dafny.CodePoint('q'))
                nw155_[int(8)] = _dafny.CodePoint('q')
                nw155_[int(9)] = (self).f29
                nw155_[int(10)] = (self).f29
                nw155_[int(11)] = (self).f29
                nw155_[int(12)] = (self).f29
                nw155_[int(13)] = (self).f29
                nw155_[int(14)] = (self).f29
                nw155_[int(15)] = (self).f29
                nw155_[int(16)] = ((self).f19)[default__.safeIndex(len(d_917_v7_), len((self).f19))]
                nw155_[int(17)] = (self).f29
                nw155_[int(18)] = (self).f29
                nw155_[int(19)] = (self).f29
                nw155_[int(20)] = (self).f29
                nw155_[int(21)] = _dafny.CodePoint('x')
                d_945_v33_ = nw155_
                d_946_v34_: _dafny.Map
                d_946_v34_ = _dafny.Map({(self).f18: (self).f29})
                d_947_v35_: _dafny.Map
                d_947_v35_ = _dafny.Map({d_945_v33_: d_946_v34_})
                d_948_v36_: _dafny.Map
                d_948_v36_ = _dafny.Map({(self).f18: (0) - (p1)})
                d_949_v37_: _dafny.MultiSet
                d_949_v37_ = _dafny.MultiSet([(self).f30, (self).f30])
                d_950_v38_: _dafny.Array
                nw156_ = _dafny.Array(None, 24)
                nw156_[int(0)] = (self).f30
                nw156_[int(1)] = len(d_947_v35_)
                nw156_[int(2)] = (self).f30
                nw156_[int(3)] = default__.fm1(p1, True, (self).f18, globalState)
                nw156_[int(4)] = p1
                nw156_[int(5)] = (self).f30
                nw156_[int(6)] = default__.safeModuloInt(len(d_919_v9_), (self).f30)
                nw156_[int(7)] = (0) - ((self).f30)
                nw156_[int(8)] = (self).f30
                nw156_[int(9)] = (755) * (len(p0))
                nw156_[int(10)] = ((self).f30) - (21)
                nw156_[int(11)] = len((self).f19)
                nw156_[int(12)] = (0) - ((self).f30)
                nw156_[int(13)] = 922
                nw156_[int(14)] = p1
                nw156_[int(15)] = p1
                nw156_[int(16)] = (p1) - ((self).f30)
                nw156_[int(17)] = ((d_948_v36_)[(self).f18] if ((self).f18) in (d_948_v36_) else 614)
                nw156_[int(18)] = (self).f30
                nw156_[int(19)] = (self).f30
                nw156_[int(20)] = (d_949_v37_).cardinality
                nw156_[int(21)] = default__.safeDivisionInt(p1, p1)
                nw156_[int(22)] = (self).f30
                nw156_[int(23)] = (self).f30
                d_950_v38_ = nw156_
                d_951_v39_: _dafny.Seq
                d_951_v39_ = _dafny.SeqWithoutIsStrInference([d_941_v29_, (d_941_v29_).set(((d_919_v9_)[len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))] if (len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))) in (d_919_v9_) else (self).f18), (self).f18), _dafny.Map({(self).f18: not((self).f18)})])
                d_952_v40_: _dafny.Seq
                d_952_v40_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                d_953_v41_: D2
                d_953_v41_ = D2_DC8((self).fm3(p1, globalState), (self).f18, (self).f18)
                pat_let_tv20_ = d_953_v41_
                def iife86_(_pat_let28_0):
                    def iife87_(d_954_dt__update__tmp_h0_):
                        def iife88_(_pat_let29_0):
                            def iife89_(d_955_dt__update_hcf21_h0_):
                                return D2_DC9(d_955_dt__update_hcf21_h0_)
                            return iife89_(_pat_let29_0)
                        return iife88_(pat_let_tv20_)
                    return iife87_(_pat_let28_0)
                rhs148_ = (d_951_v39_)[default__.safeIndex(p1, len(d_951_v39_))]
                rhs149_ = d_942_v30_
                rhs150_ = (len(p0)) - ((self).f30)
                rhs151_ = (iife86_(d_943_v31_) if (d_952_v40_)[default__.safeIndex((self).f30, len(d_952_v40_))] else d_943_v31_)
                rhs152_ = d_950_v38_
                lhs132_ = globalState
                d_941_v29_ = rhs148_
                d_942_v30_ = rhs149_
                lhs132_.f0 = rhs150_
                d_943_v31_ = rhs151_
                d_950_v38_ = rhs152_
                index154_ = default__.safeIndex(55, (d_942_v30_).length(0))
                (d_942_v30_)[index154_] = (254) <= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vd"))))
                index155_ = default__.safeIndex(55, (d_942_v30_).length(0))
                (d_942_v30_)[index155_] = not (default__.fm23(_dafny.Set({p1}), 274, globalState)) or ((d_922_v12_).ispropersubset(d_922_v12_))
                d_956_v42_: str
                d_956_v42_ = _dafny.CodePoint('v')
                d_956_v42_ = (self).f29
                index156_ = default__.safeIndex(55, (d_942_v30_).length(0))
                (d_942_v30_)[index156_] = (self).f18
                d_957_v43_: D0
                d_957_v43_ = D0_DC0((self).f18)
                index157_ = default__.safeIndex(55, (d_942_v30_).length(0))
                (d_942_v30_)[index157_] = (d_957_v43_).cf0
            elif True:
                d_958_v44_: _dafny.MultiSet
                d_958_v44_ = _dafny.MultiSet([(_dafny.MultiSet([114])).cardinality])
                d_959_v45_: _dafny.Map
                d_959_v45_ = _dafny.Map({d_917_v7_: (d_958_v44_).cardinality})
                d_960_v46_: D1
                d_960_v46_ = D1_DC4((self).f18)
                d_961_v47_: _dafny.Array
                nw157_ = _dafny.Array(None, 23)
                nw157_[int(0)] = default__.safeModuloInt((self).f30, p1)
                nw157_[int(1)] = (p1) - (p1)
                nw157_[int(2)] = (self).f30
                nw157_[int(3)] = (default__.fm31(len((d_959_v45_)), (self).f18, d_960_v46_, globalState)).cardinality
                nw157_[int(4)] = (self).f30
                nw157_[int(5)] = default__.safeDivisionInt((self).f30, (self).f30)
                nw157_[int(6)] = 197
                nw157_[int(7)] = (self).f30
                nw157_[int(8)] = (self).f30
                nw157_[int(9)] = (0) - (default__.safeModuloInt(p1, p1))
                nw157_[int(10)] = (self).f30
                nw157_[int(11)] = (self).f30
                nw157_[int(12)] = (self).f30
                nw157_[int(13)] = len(p2)
                nw157_[int(14)] = (self).f30
                nw157_[int(15)] = len(p0)
                nw157_[int(16)] = 672
                nw157_[int(17)] = ((self).f30) * (len(d_917_v7_))
                nw157_[int(18)] = p1
                nw157_[int(19)] = len(d_917_v7_)
                nw157_[int(20)] = len((p2).set(default__.safeIndex((d_922_v12_).cardinality, len(p2)), (self).f29))
                nw157_[int(21)] = p1
                nw157_[int(22)] = (self).f30
                d_961_v47_ = nw157_
                index158_ = default__.safeIndex(343, (d_961_v47_).length(0))
                (d_961_v47_)[index158_] = p1
                d_962_v48_: _dafny.MultiSet
                d_962_v48_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nla"))) + ((self).f19), _dafny.SeqWithoutIsStrInference([(self).f29 for d_963_i1_ in range(default__.abs(621))]), _dafny.SeqWithoutIsStrInference([(self).f29 for d_964_i2_ in range(default__.abs(228))])])
                index159_ = default__.safeIndex(343, (d_961_v47_).length(0))
                rhs153_ = ((self).f18) or ((self).f18)
                rhs154_ = (((self).f30) - (p1)) + ((0) - (p1))
                rhs155_ = (d_922_v12_).cardinality
                rhs156_ = ((d_962_v48_ if (self).f18 else (d_962_v48_).set(_dafny.SeqWithoutIsStrInference([(self).f29 for d_965_i3_ in range(default__.abs(77))]), default__.abs(p1)))) - ((d_962_v48_).set(p2, default__.abs((self).f30)))
                lhs133_ = globalState
                lhs134_ = globalState
                lhs135_ = d_961_v47_
                lhs136_ = default__.safeIndex(343, (d_961_v47_).length(0))
                lhs133_.f2 = rhs153_
                lhs134_.f3 = rhs154_
                lhs135_[lhs136_] = rhs155_
                d_962_v48_ = rhs156_
                d_966_v49_: _dafny.Array
                def lambda36_(d_967_v47_):
                    def lambda37_(d_968_i4_):
                        return _dafny.Map({(self).f30: (d_967_v47_)[default__.safeIndex(343, (d_967_v47_).length(0))]})

                    return lambda37_

                init20_ = lambda36_(d_961_v47_)
                nw158_ = _dafny.Array(None, 26)
                for i0_20_ in range(nw158_.length(0)):
                    nw158_[i0_20_] = init20_(i0_20_)
                d_966_v49_ = nw158_
                nw159_ = _dafny.Array(_dafny.Map({}), 22)
                d_966_v49_ = nw159_
                d_969_v50_: _dafny.Map
                d_969_v50_ = _dafny.Map({(self).f18: (self).f18})
                d_969_v50_ = (d_969_v50_).set((self).f18, not((self).f18))
                d_970_v51_: _dafny.Array
                nw160_ = _dafny.Array(False, 6)
                d_970_v51_ = nw160_
                d_971_v52_: _dafny.Array
                d_971_v52_ = d_961_v47_
                d_972_v53_: _dafny.Map
                d_972_v53_ = _dafny.Map({(self).f18: (d_971_v52_)})
                index160_ = default__.safeIndex(726, (d_970_v51_).length(0))
                (d_970_v51_)[index160_] = not((self).fm3(len((_dafny.Map({745: d_961_v47_})).set((d_961_v47_)[default__.safeIndex(343, (d_961_v47_).length(0))], ((d_972_v53_)[(self).f18] if ((self).f18) in (d_972_v53_) else d_961_v47_))), globalState))
                index161_ = default__.safeIndex(726, (d_970_v51_).length(0))
                (d_970_v51_)[index161_] = (self).f18
                index162_ = default__.safeIndex(726, (d_970_v51_).length(0))
                (d_970_v51_)[index162_] = (d_970_v51_)[default__.safeIndex(726, (d_970_v51_).length(0))]
            d_973_v54_: _dafny.Array
            nw161_ = _dafny.Array(False, 1)
            d_973_v54_ = nw161_
            index163_ = default__.safeIndex(843, (d_973_v54_).length(0))
            (d_973_v54_)[index163_] = (self).f18
            index164_ = default__.safeIndex(843, (d_973_v54_).length(0))
            (d_973_v54_)[index164_] = (self).fm3(len((self).f19), globalState)
        d_974_v55_: _dafny.MultiSet
        d_974_v55_ = _dafny.MultiSet([(self).f30])
        hi5_ = 821
        for d_975_i5_ in range((d_974_v55_).cardinality, hi5_):
            d_976_v56_: _dafny.Array
            nw162_ = _dafny.Array(int(0), 12)
            d_976_v56_ = nw162_
            d_977_v57_: _dafny.Seq
            d_977_v57_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_978_v58_: _dafny.Map
            d_978_v58_ = _dafny.Map({(self).f18: d_977_v57_})
            index165_ = default__.safeIndex(910, (d_976_v56_).length(0))
            (d_976_v56_)[index165_] = default__.safeDivisionInt(len(((d_978_v58_)[(self).f18] if ((self).f18) in (d_978_v58_) else d_977_v57_)), d_975_i5_)
            d_979_v59_: _dafny.Array
            def lambda38_(d_980_p1_):
                def lambda39_(d_981_i6_):
                    return (_dafny.Map({(self).f18: (self).f30})) | (_dafny.Map({True: d_980_p1_}))

                return lambda39_

            init21_ = lambda38_(p1)
            nw163_ = _dafny.Array(None, 28)
            for i0_21_ in range(nw163_.length(0)):
                nw163_[i0_21_] = init21_(i0_21_)
            d_979_v59_ = nw163_
            d_982_v60_: _dafny.Seq
            d_982_v60_ = _dafny.SeqWithoutIsStrInference([p1])
            d_983_v61_: _dafny.Map
            d_983_v61_ = _dafny.Map({not((self).f18): len((d_982_v60_).set(default__.safeIndex((self).f30, len(d_982_v60_)), d_975_i5_))})
            index166_ = default__.safeIndex(137, (d_979_v59_).length(0))
            (d_979_v59_)[index166_] = d_983_v61_
            d_984_v62_: _dafny.Seq
            d_984_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f18]), default__.fm32(globalState)])
            d_985_v64_: _dafny.Seq
            d_985_v64_ = _dafny.SeqWithoutIsStrInference([p2])
            d_986_v65_: _dafny.MultiSet
            d_986_v65_ = _dafny.MultiSet([True])
            d_987_v66_: _dafny.Map
            d_987_v66_ = _dafny.Map({d_975_i5_: (d_982_v60_)[default__.safeIndex((d_986_v65_).cardinality, len(d_982_v60_))]})
            index167_ = default__.safeIndex(910, (d_976_v56_).length(0))
            index168_ = default__.safeIndex(137, (d_979_v59_).length(0))
            def iife90_():
                coll30_ = _dafny.Map()
                compr_30_: _dafny.Seq
                for compr_30_ in (d_985_v64_).Elements:
                    d_988_v63_: _dafny.Seq = compr_30_
                    if (d_988_v63_) in (d_985_v64_):
                        coll30_[d_988_v63_] = (self).f18
                return _dafny.Map(coll30_)
            rhs157_ = len((default__.fm30(d_984_v62_, iife90_()
            , globalState)) | ((d_987_v66_ if (self).f18 else d_987_v66_)))
            rhs158_ = (0) - ((d_975_i5_) - (d_975_i5_))
            rhs159_ = d_975_i5_
            rhs160_ = d_983_v61_
            lhs137_ = globalState
            lhs138_ = d_976_v56_
            lhs139_ = default__.safeIndex(910, (d_976_v56_).length(0))
            lhs140_ = globalState
            lhs141_ = d_979_v59_
            lhs142_ = default__.safeIndex(137, (d_979_v59_).length(0))
            lhs137_.f0 = rhs157_
            lhs138_[lhs139_] = rhs158_
            lhs140_.f15 = rhs159_
            lhs141_[lhs142_] = rhs160_
            (globalState).f2 = (self).fm2((self).f18, False, (self).f18, globalState)
            if (self).f18:
                d_989_v67_: T0
                nw164_ = C2()
                nw164_.ctor__(default__.fm1((d_976_v56_)[default__.safeIndex(910, (d_976_v56_).length(0))], (self).f18, (self).f18, globalState), _dafny.SeqWithoutIsStrInference([len(p2)]), not((d_977_v57_)[default__.safeIndex((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30])), (self).f30])).cardinality, len(d_977_v57_))]))
                d_989_v67_ = nw164_
                d_990_v68_: D1
                d_990_v68_ = D1_DC2(d_989_v67_)
                d_991_v69_: _dafny.Map
                d_991_v69_ = _dafny.Map({D1_DC5(d_990_v68_): p1})
                d_992_v70_: D1
                d_992_v70_ = D1_DC5(d_990_v68_)
                d_991_v69_ = (d_991_v69_).set(d_992_v70_, p1)
                d_993_v71_: _dafny.Array
                def lambda40_(d_994_v61_):
                    def lambda41_(d_995_i7_):
                        return D2_DC6(d_994_v61_)

                    return lambda41_

                init22_ = lambda40_(d_983_v61_)
                nw165_ = _dafny.Array(None, 24)
                for i0_22_ in range(nw165_.length(0)):
                    nw165_[i0_22_] = init22_(i0_22_)
                d_993_v71_ = nw165_
                d_996_v72_: _dafny.Array
                nw166_ = _dafny.Array(None, 2)
                nw166_[int(0)] = d_993_v71_
                nw166_[int(1)] = d_993_v71_
                d_996_v72_ = nw166_
                rhs161_ = p1
                rhs162_ = d_996_v72_
                rhs163_ = p1
                rhs164_ = (self).f18
                lhs143_ = globalState
                lhs144_ = globalState
                lhs145_ = globalState
                lhs143_.f3 = rhs161_
                d_996_v72_ = rhs162_
                lhs144_.f3 = rhs163_
                lhs145_.f2 = rhs164_
                d_997_v73_: _dafny.MultiSet
                d_997_v73_ = _dafny.MultiSet([(d_987_v66_).set((d_976_v56_)[default__.safeIndex(910, (d_976_v56_).length(0))], 695), (d_987_v66_).set(d_975_i5_, (self).f30)])
                d_987_v66_ = (d_987_v66_).set((self).f30, (d_997_v73_).cardinality)
                index169_ = default__.safeIndex(910, (d_976_v56_).length(0))
                (d_976_v56_)[index169_] = default__.safeDivisionInt((d_976_v56_)[default__.safeIndex(910, (d_976_v56_).length(0))], p1)
                d_976_v56_ = d_976_v56_
            elif True:
                d_998_v74_: _dafny.Map
                d_998_v74_ = _dafny.Map({827: (d_974_v55_).issubset(d_974_v55_)})
                d_998_v74_ = (d_998_v74_).set((d_976_v56_)[default__.safeIndex(910, (d_976_v56_).length(0))], (self).f18)
                (globalState).f2 = (self).f18
                index170_ = default__.safeIndex(910, (d_976_v56_).length(0))
                (d_976_v56_)[index170_] = ((d_975_i5_) - (p1)) - (len(p2))
                (globalState).f3 = ((self).f30) + (d_975_i5_)
                (globalState).f2 = (self).f18
            d_999_v75_: _dafny.Map
            d_999_v75_ = _dafny.Map({p1: (self).f18})
            d_1000_v76_: C0
            nw167_ = C0()
            nw167_.ctor__(p1, default__.safeModuloInt(len((D0_DC1((self).fm3((d_976_v56_)[default__.safeIndex(910, (d_976_v56_).length(0))], globalState), d_975_i5_, p2, (self).f18, len(d_999_v75_))).cf3), (d_976_v56_)[default__.safeIndex(910, (d_976_v56_).length(0))]))
            d_1000_v76_ = nw167_
        (globalState).f2 = (self).f18
        d_1001_v77_: _dafny.Map
        d_1001_v77_ = _dafny.Map({(self).f30: (self).fm3((self).f30, globalState)})
        d_1002_i8_: int
        d_1002_i8_ = 0
        with _dafny.label("8"):
            while ((d_1001_v77_)[(self).f30] if ((self).f30) in (d_1001_v77_) else (self).f18):
                with _dafny.c_label("8"):
                    if (d_1002_i8_) >= (100):
                        raise _dafny.Break("8")
                    d_1002_i8_ = (d_1002_i8_) + (1)
                    (globalState).f2 = (self).f18
                    if not((self).f18):
                        d_1003_v78_: _dafny.Seq
                        d_1003_v78_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeDivisionInt(161, p1))])
                        d_1004_v79_: _dafny.Seq
                        d_1004_v79_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, (self).f18, (self).f18])
                        rhs165_ = (d_1003_v78_).set(default__.safeIndex(len(((self).f19) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdxtowp")))), len(d_1003_v78_)), p1)
                        rhs166_ = (default__.fm1((self).f30, (self).f18, True, globalState)) >= ((0) - (p1))
                        rhs167_ = (0) - ((972) * ((self).f30))
                        rhs168_ = ((self).f18) == ((self).f18)
                        rhs169_ = (len((d_1004_v79_).set(default__.safeIndex(p1, len(d_1004_v79_)), (self).f18))) * (p1)
                        lhs146_ = globalState
                        lhs147_ = globalState
                        lhs148_ = globalState
                        lhs149_ = globalState
                        lhs150_ = globalState
                        lhs146_.f10 = rhs165_
                        lhs147_.f2 = rhs166_
                        lhs148_.f15 = rhs167_
                        lhs149_.f2 = rhs168_
                        lhs150_.f0 = rhs169_
                        d_1005_v80_: C2
                        nw168_ = C2()
                        nw168_.ctor__((self).f30, default__.fm33((self).f18, globalState), (67) != (246))
                        d_1005_v80_ = nw168_
                        d_1006_v81_: _dafny.Array
                        def lambda42_(d_1007_i9_):
                            return (self).f18

                        init23_ = lambda42_
                        nw169_ = _dafny.Array(None, 9)
                        for i0_23_ in range(nw169_.length(0)):
                            nw169_[i0_23_] = init23_(i0_23_)
                        d_1006_v81_ = nw169_
                        d_1008_v82_: _dafny.MultiSet
                        d_1008_v82_ = _dafny.MultiSet([d_1006_v81_, d_1006_v81_, d_1006_v81_, d_1006_v81_])
                        d_1008_v82_ = (d_1008_v82_ if (self).f18 else d_1008_v82_)
                        d_1009_v84_: _dafny.Map
                        d_1009_v84_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f18]): (self).f30})
                        d_1010_v85_: _dafny.Map
                        d_1010_v85_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f18]): (self).fm3(p1, globalState)})
                        d_1011_v86_: C1
                        nw170_ = C1()
                        def iife91_():
                            coll31_ = _dafny.Map()
                            compr_31_: _dafny.Seq
                            for compr_31_ in (d_1009_v84_).keys.Elements:
                                d_1012_v83_: _dafny.Seq = compr_31_
                                if (d_1012_v83_) in (d_1009_v84_):
                                    coll31_[d_1012_v83_] = (self).f18
                            return _dafny.Map(coll31_)
                        nw170_.ctor__((iife91_()
                        ) != (d_1010_v85_))
                        d_1011_v86_ = nw170_
                        d_1013_v87_: D2
                        d_1013_v87_ = D2_DC8((self).f18, (self).f18, not((self).f18))
                        d_1014_v88_: C1
                        nw171_ = C1()
                        nw171_.ctor__(((d_1013_v87_).cf20) or (False))
                        d_1014_v88_ = nw171_
                    elif True:
                        d_1015_v89_: _dafny.Seq
                        d_1015_v89_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))])
                        (globalState).f0 = default__.safeModuloInt((self).f30, ((0) - (len(d_1015_v89_))) * ((self).f30))
                        d_1016_v90_: _dafny.Array
                        def lambda43_(d_1017_i10_):
                            return (self).f18

                        init24_ = lambda43_
                        nw172_ = _dafny.Array(None, 15)
                        for i0_24_ in range(nw172_.length(0)):
                            nw172_[i0_24_] = init24_(i0_24_)
                        d_1016_v90_ = nw172_
                        d_1018_v91_: _dafny.Seq
                        d_1018_v91_ = _dafny.SeqWithoutIsStrInference([True, (self).f18])
                        index171_ = default__.safeIndex(31, (d_1016_v90_).length(0))
                        (d_1016_v90_)[index171_] = (d_1018_v91_)[default__.safeIndex(p1, len(d_1018_v91_))]
                        index172_ = default__.safeIndex(31, (d_1016_v90_).length(0))
                        rhs170_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqgd"))
                        rhs171_ = (self).f18
                        rhs172_ = p1
                        rhs173_ = not (True) or (not((self).f18))
                        rhs174_ = (self).fm2((self).f18, (self).f18, (self).f18, globalState)
                        lhs151_ = globalState
                        lhs152_ = globalState
                        lhs153_ = globalState
                        lhs154_ = globalState
                        lhs155_ = d_1016_v90_
                        lhs156_ = default__.safeIndex(31, (d_1016_v90_).length(0))
                        lhs151_.f12 = rhs170_
                        lhs152_.f2 = rhs171_
                        lhs153_.f15 = rhs172_
                        lhs154_.f2 = rhs173_
                        lhs155_[lhs156_] = rhs174_
                        (globalState).f2 = (d_1016_v90_)[default__.safeIndex(31, (d_1016_v90_).length(0))]
                        d_1019_v92_: int
                        d_1020_v93_: int
                        d_1021_v94_: _dafny.Array
                        d_1022_v95_: _dafny.Map
                        out32_: int
                        out33_: int
                        out34_: _dafny.Array
                        out35_: _dafny.Map
                        out32_, out33_, out34_, out35_ = default__.m0(globalState)
                        d_1019_v92_ = out32_
                        d_1020_v93_ = out33_
                        d_1021_v94_ = out34_
                        d_1022_v95_ = out35_
                        (globalState).f2 = (self).f18
                    (globalState).f15 = p1
                    d_1001_v77_ = (d_1001_v77_).set(((self).f30) - (403), (self).f18)
                    pass
            pass
        d_1023_v96_: _dafny.Set
        d_1023_v96_ = _dafny.Set({p1})
        d_1024_v97_: _dafny.Seq
        d_1024_v97_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
        d_1025_v98_: _dafny.Seq
        d_1025_v98_ = _dafny.SeqWithoutIsStrInference([len(d_1023_v96_), len(d_1024_v97_), p1, (self).f30, p1])
        hi6_ = len(d_1025_v98_)
        for d_1026_i11_ in range(528, hi6_):
            d_1027_v99_: _dafny.Array
            nw173_ = _dafny.Array(None, 4)
            nw173_[int(0)] = (self).f18
            nw173_[int(1)] = (self).f18
            nw173_[int(2)] = (self).f18
            nw173_[int(3)] = (d_1024_v97_)[default__.safeIndex(d_1026_i11_, len(d_1024_v97_))]
            d_1027_v99_ = nw173_
            d_1028_v100_: _dafny.Map
            d_1028_v100_ = _dafny.Map({len((d_1024_v97_) + (d_1024_v97_)): d_1027_v99_})
            d_1028_v100_ = (d_1028_v100_).set(20, d_1027_v99_)
            d_1029_v101_: _dafny.MultiSet
            d_1029_v101_ = _dafny.MultiSet([(self).f18])
            d_1030_v102_: _dafny.Map
            d_1030_v102_ = _dafny.Map({(self).f18: p1})
            d_1031_v104_: _dafny.Seq
            d_1031_v104_ = _dafny.SeqWithoutIsStrInference([d_1023_v96_])
            d_1032_v105_: _dafny.Array
            nw174_ = _dafny.Array(None, 29)
            nw174_[int(0)] = (self).f30
            nw174_[int(1)] = p1
            nw174_[int(2)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btuvfidmd")))
            nw174_[int(3)] = -309
            nw174_[int(4)] = (self).fm17((0) - ((self).f30), d_1029_v101_, globalState)
            nw174_[int(5)] = p1
            nw174_[int(6)] = 725
            nw174_[int(7)] = (0) - (p1)
            nw174_[int(8)] = p1
            nw174_[int(9)] = d_1026_i11_
            nw174_[int(10)] = p1
            nw174_[int(11)] = len((p2).set(default__.safeIndex(d_1026_i11_, len(p2)), (self).f29))
            nw174_[int(12)] = p1
            nw174_[int(13)] = p1
            nw174_[int(14)] = len(d_1030_v102_)
            nw174_[int(15)] = (self).f30
            nw174_[int(16)] = (self).f30
            nw174_[int(17)] = d_1026_i11_
            nw174_[int(18)] = (self).f30
            nw174_[int(19)] = (self).fm18(globalState)
            nw174_[int(20)] = d_1026_i11_
            def iife92_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(711, 606):
                    d_1033_v103_: int = compr_32_
                    if ((711) <= (d_1033_v103_)) and ((d_1033_v103_) < (606)):
                        coll32_ = coll32_.union(_dafny.Set([(d_1033_v103_) - (d_1026_i11_)]))
                return _dafny.Set(coll32_)
            nw174_[int(21)] = len(iife92_()
            )
            nw174_[int(22)] = p1
            nw174_[int(23)] = (0) - ((self).f30)
            nw174_[int(24)] = (self).f30
            nw174_[int(25)] = p1
            nw174_[int(26)] = len(d_1031_v104_)
            nw174_[int(27)] = p1
            nw174_[int(28)] = p1
            d_1032_v105_ = nw174_
            d_1034_v106_: _dafny.Map
            d_1034_v106_ = _dafny.Map({(d_1024_v97_)[default__.safeIndex(932, len(d_1024_v97_))]: d_1032_v105_})
            d_1034_v106_ = (d_1034_v106_).set((self).f18, d_1032_v105_)
            (globalState).f15 = d_1026_i11_
            d_1035_v107_: _dafny.Map
            d_1035_v107_ = _dafny.Map({(self).f18: (d_1024_v97_)[default__.safeIndex(d_1026_i11_, len(d_1024_v97_))]})
            d_1035_v107_ = (d_1035_v107_).set(False, (self).f18)
        d_1036_v108_: D2
        d_1036_v108_ = D2_DC8((self).f18, (self).f18, (self).f18)
        d_1037_v109_: _dafny.Map
        d_1037_v109_ = _dafny.Map({len(d_1024_v97_): d_1036_v108_})
        d_1037_v109_ = (default__.fm34((self).f30, (self).f30, globalState)) | (d_1037_v109_)

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_1038_v0_: D2
        d_1038_v0_ = D2_DC6(_dafny.Map({p1: (self).f30}))
        source9_ = d_1038_v0_
        if source9_.is_DC7:
            d_1039___mcc_h0_ = source9_.cf15
            d_1040___mcc_h1_ = source9_.cf16
            d_1041___mcc_h2_ = source9_.cf17
            d_1042_cf17_ = d_1041___mcc_h2_
            d_1043_cf16_ = d_1040___mcc_h1_
            d_1044_cf15_ = d_1039___mcc_h0_
            d_1045_v1_: D3
            d_1045_v1_ = D3_DC14(D3_DC11(p1, p2))
            source10_ = d_1045_v1_
            if source10_.is_DC11:
                d_1046___mcc_h8_ = source10_.cf23
                d_1047___mcc_h9_ = source10_.cf24
                d_1048_cf24_ = d_1047___mcc_h9_
                d_1049_cf23_ = d_1046___mcc_h8_
                r0 = d_1044_cf15_
                d_1050_v2_: _dafny.Array
                nw175_ = _dafny.Array(False, 25)
                d_1050_v2_ = nw175_
                index173_ = default__.safeIndex(668, (d_1050_v2_).length(0))
                (d_1050_v2_)[index173_] = (self).f18
                index174_ = default__.safeIndex(668, (d_1050_v2_).length(0))
                (d_1050_v2_)[index174_] = (self).f18
                d_1051_v3_: _dafny.Set
                d_1051_v3_ = _dafny.Set({(self).f30})
                (globalState).f15 = default__.fm1((len(d_1051_v3_)) - (d_1042_cf17_), (default__.fm35(d_1042_cf17_, d_1048_cf24_, globalState)).cf4, d_1048_cf24_, globalState)
                (globalState).f0 = default__.safeDivisionInt(d_1044_cf15_, (default__.fm32(globalState)).cardinality)
            elif source10_.is_DC12:
                d_1052___mcc_h10_ = source10_.cf25
                d_1053___mcc_h11_ = source10_.cf26
                d_1054___mcc_h12_ = source10_.cf27
                d_1055___mcc_h13_ = source10_.cf28
                d_1056___mcc_h14_ = source10_.cf29
                d_1057_cf29_ = d_1056___mcc_h14_
                d_1058_cf28_ = d_1055___mcc_h13_
                d_1059_cf27_ = d_1054___mcc_h12_
                d_1060_cf26_ = d_1053___mcc_h11_
                d_1061_cf25_ = d_1052___mcc_h10_
                d_1062_v4_: _dafny.Seq
                d_1062_v4_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                d_1063_v5_: _dafny.Map
                d_1063_v5_ = _dafny.Map({d_1042_cf17_: len(d_1062_v4_)})
                d_1064_v6_: _dafny.Map
                d_1064_v6_ = _dafny.Map({d_1063_v5_: 505})
                d_1065_v7_: _dafny.Map
                d_1065_v7_ = _dafny.Map({d_1042_cf17_: d_1063_v5_})
                d_1066_v8_: _dafny.Set
                d_1066_v8_ = _dafny.Set({(self).f30, (self).f30})
                d_1067_v9_: _dafny.Seq
                d_1067_v9_ = _dafny.SeqWithoutIsStrInference([72, len(d_1066_v8_)])
                d_1068_v10_: _dafny.Set
                d_1068_v10_ = _dafny.Set({p1})
                d_1069_v11_: _dafny.Seq
                d_1069_v11_ = _dafny.SeqWithoutIsStrInference([d_1068_v10_, default__.fm36((self).f18, d_1059_cf27_, (self).f18, globalState)])
                d_1064_v6_ = (d_1064_v6_).set(((d_1065_v7_)[d_1042_cf17_] if (d_1042_cf17_) in (d_1065_v7_) else _dafny.Map({d_1044_cf15_: len(d_1067_v9_)})), len(d_1069_v11_))
                d_1070_v12_: _dafny.Seq
                d_1070_v12_ = _dafny.SeqWithoutIsStrInference([(d_1067_v9_) + (d_1067_v9_)])
                d_1070_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([590]) for d_1071_i0_ in range(default__.abs(88))])
                d_1072_v13_: str
                d_1072_v13_ = _dafny.CodePoint('v')
                d_1072_v13_ = ((self).f19)[default__.safeIndex(d_1044_cf15_, len((self).f19))]
                d_1073_v14_: D2
                d_1073_v14_ = D2_DC7(485, d_1060_cf26_, len((self).f19))
                d_1043_cf16_ = (d_1059_cf27_) >= (default__.fm1(d_1042_cf17_, d_1057_cf29_, (d_1073_v14_).cf16, globalState))
            elif source10_.is_DC13:
                d_1074___mcc_h15_ = source10_.cf30
                d_1075___mcc_h16_ = source10_.cf31
                d_1076___mcc_h17_ = source10_.cf32
                d_1077___mcc_h18_ = source10_.cf33
                d_1078___mcc_h19_ = source10_.cf34
                d_1079_cf34_ = d_1078___mcc_h19_
                d_1080_cf33_ = d_1077___mcc_h18_
                d_1081_cf32_ = d_1076___mcc_h17_
                d_1082_cf31_ = d_1075___mcc_h16_
                d_1083_cf30_ = d_1074___mcc_h15_
                d_1084_v15_: _dafny.Map
                d_1084_v15_ = _dafny.Map({p1: p1})
                d_1085_v16_: _dafny.Array
                nw176_ = _dafny.Array(None, 3)
                nw176_[int(0)] = d_1084_v15_
                nw176_[int(1)] = d_1084_v15_
                nw176_[int(2)] = _dafny.Map({p1: True})
                d_1085_v16_ = nw176_
                index175_ = default__.safeIndex(833, (d_1085_v16_).length(0))
                (d_1085_v16_)[index175_] = d_1084_v15_
                index176_ = default__.safeIndex(833, (d_1085_v16_).length(0))
                (d_1085_v16_)[index176_] = ((_dafny.Map({(self).f18: p2})) | (d_1084_v15_)) | (d_1084_v15_)
                d_1086_v17_: _dafny.MultiSet
                d_1086_v17_ = _dafny.MultiSet([p2])
                d_1087_v18_: C0
                nw177_ = C0()
                nw177_.ctor__((len((self).f19)) * (d_1083_cf30_), ((d_1086_v17_)[p2] if (p2) in (d_1086_v17_) else d_1080_cf33_))
                d_1087_v18_ = nw177_
                d_1088_v19_: _dafny.Array
                nw178_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_1088_v19_ = nw178_
                d_1089_v20_: _dafny.Array
                nw179_ = _dafny.Array(False, 9)
                d_1089_v20_ = nw179_
                index177_ = default__.safeIndex(870, (d_1088_v19_).length(0))
                (d_1088_v19_)[index177_] = d_1089_v20_
                index178_ = default__.safeIndex(870, (d_1088_v19_).length(0))
                rhs175_ = not((p2) and ((self).f18))
                rhs176_ = d_1089_v20_
                rhs177_ = p2
                lhs157_ = d_1088_v19_
                lhs158_ = default__.safeIndex(870, (d_1088_v19_).length(0))
                lhs159_ = globalState
                d_1043_cf16_ = rhs175_
                lhs157_[lhs158_] = rhs176_
                lhs159_.f2 = rhs177_
                d_1090_v21_: _dafny.Seq
                d_1090_v21_ = _dafny.SeqWithoutIsStrInference([d_1083_cf30_])
                (globalState).f3 = len(d_1090_v21_)
            elif source10_.is_DC10:
                d_1091___mcc_h20_ = source10_.cf22
                d_1092_cf22_ = d_1091___mcc_h20_
                d_1093_v22_: _dafny.Array
                def lambda44_(d_1094_cf15_):
                    def lambda45_(d_1095_i1_):
                        return ((self).f30) <= (d_1094_cf15_)

                    return lambda45_

                init25_ = lambda44_(d_1044_cf15_)
                nw180_ = _dafny.Array(None, 22)
                for i0_25_ in range(nw180_.length(0)):
                    nw180_[i0_25_] = init25_(i0_25_)
                d_1093_v22_ = nw180_
                index179_ = default__.safeIndex(415, (d_1093_v22_).length(0))
                (d_1093_v22_)[index179_] = p1
                d_1096_v23_: _dafny.Seq
                d_1096_v23_ = _dafny.SeqWithoutIsStrInference([d_1042_cf17_])
                d_1097_v24_: T1
                nw181_ = C3()
                nw181_.ctor__((d_1096_v23_) + (d_1096_v23_), (self).f19, (self).f20, d_1043_cf16_)
                d_1097_v24_ = nw181_
                index180_ = default__.safeIndex(415, (d_1093_v22_).length(0))
                rhs178_ = p2
                rhs179_ = (self).f19
                rhs180_ = ((self).f30) not in (d_1096_v23_)
                rhs181_ = d_1097_v24_
                lhs160_ = d_1093_v22_
                lhs161_ = default__.safeIndex(415, (d_1093_v22_).length(0))
                lhs162_ = globalState
                lhs163_ = globalState
                lhs160_[lhs161_] = rhs178_
                lhs162_.f12 = rhs179_
                lhs163_.f2 = rhs180_
                d_1097_v24_ = rhs181_
                d_1098_v25_: _dafny.Map
                d_1098_v25_ = _dafny.Map({((self).f18) == (d_1043_cf16_): (d_1044_cf15_) == (d_1042_cf17_)})
                d_1099_v26_: _dafny.MultiSet
                d_1099_v26_ = _dafny.MultiSet([(d_1096_v23_)[default__.safeIndex(d_1044_cf15_, len(d_1096_v23_))]])
                d_1100_v27_: _dafny.Set
                d_1100_v27_ = _dafny.Set({d_1043_cf16_, ((d_1099_v26_).set(d_1044_cf15_, default__.abs(default__.fm1(444, (self).f18, d_1043_cf16_, globalState)))).ispropersubset(d_1099_v26_), p2, False})
                d_1101_v28_: _dafny.MultiSet
                d_1101_v28_ = _dafny.MultiSet([not((self).fm3((self).f30, globalState))])
                d_1102_v29_: _dafny.Seq
                d_1102_v29_ = _dafny.SeqWithoutIsStrInference([(self).f18, p1])
                d_1103_v30_: _dafny.Map
                d_1103_v30_ = _dafny.Map({(self).f18: d_1102_v29_})
                d_1104_v31_: _dafny.Map
                d_1104_v31_ = d_1103_v30_
                rhs182_ = True
                rhs183_ = (default__.fm47(d_1101_v28_, d_1042_cf17_, len((d_1104_v31_)), (d_1093_v22_)[default__.safeIndex(415, (d_1093_v22_).length(0))], globalState)).set((d_1093_v22_)[default__.safeIndex(415, (d_1093_v22_).length(0))], (self).f18)
                rhs184_ = d_1100_v27_
                lhs164_ = globalState
                lhs164_.f2 = rhs182_
                d_1098_v25_ = rhs183_
                d_1100_v27_ = rhs184_
                d_1105_v32_: _dafny.Array
                nw182_ = _dafny.Array(int(0), 10)
                d_1105_v32_ = nw182_
                index181_ = default__.safeIndex(687, (d_1105_v32_).length(0))
                (d_1105_v32_)[index181_] = (self).f30
                index182_ = default__.safeIndex(687, (d_1105_v32_).length(0))
                (d_1105_v32_)[index182_] = (d_1044_cf15_) + (d_1042_cf17_)
                d_1106_v33_: _dafny.Map
                d_1106_v33_ = _dafny.Map({(self).f29: d_1105_v32_})
                d_1107_v34_: D2
                d_1107_v34_ = D2_DC7(d_1042_cf17_, False, len(d_1106_v33_))
                d_1108_v35_: C2
                nw183_ = C2()
                nw183_.ctor__((d_1044_cf15_) * ((d_1105_v32_)[default__.safeIndex(687, (d_1105_v32_).length(0))]), (d_1096_v23_) + (d_1096_v23_), (d_1107_v34_).cf16)
                d_1108_v35_ = nw183_
            elif True:
                d_1109___mcc_h21_ = source10_.cf35
                d_1110_cf35_ = d_1109___mcc_h21_
                d_1111_v36_: _dafny.Seq
                d_1111_v36_ = _dafny.SeqWithoutIsStrInference([False, d_1043_cf16_, (self).f18])
                d_1112_v37_: _dafny.Map
                d_1112_v37_ = _dafny.Map({d_1043_cf16_: (d_1111_v36_).set(default__.safeIndex(d_1044_cf15_, len(d_1111_v36_)), True)})
                d_1112_v37_ = (d_1112_v37_).set(p1, (d_1111_v36_) + (_dafny.SeqWithoutIsStrInference([p1])))
                d_1113_v38_: _dafny.Array
                nw184_ = _dafny.Array(_dafny.Set({}), 16)
                d_1113_v38_ = nw184_
                d_1114_v39_: _dafny.Set
                d_1114_v39_ = _dafny.Set({(self).f30, d_1044_cf15_})
                index183_ = default__.safeIndex(365, (d_1113_v38_).length(0))
                (d_1113_v38_)[index183_] = d_1114_v39_
                index184_ = default__.safeIndex(365, (d_1113_v38_).length(0))
                (d_1113_v38_)[index184_] = d_1114_v39_
                d_1115_v40_: _dafny.Seq
                d_1115_v40_ = _dafny.SeqWithoutIsStrInference([d_1042_cf17_, (self).f30])
                d_1116_v41_: _dafny.Seq
                d_1116_v41_ = _dafny.SeqWithoutIsStrInference([len(d_1115_v40_), d_1042_cf17_])
                d_1117_v42_: _dafny.Map
                d_1117_v42_ = _dafny.Map({_dafny.CodePoint('n'): (self).fm3(d_1042_cf17_, globalState)})
                (globalState).f3 = default__.fm1(((self).f30) + (default__.fm1(len(d_1116_v41_), d_1043_cf16_, p1, globalState)), d_1043_cf16_, ((d_1117_v42_)[((self).f19)[default__.safeIndex((0) - (d_1042_cf17_), len((self).f19))]] if (((self).f19)[default__.safeIndex((0) - (d_1042_cf17_), len((self).f19))]) in (d_1117_v42_) else p1), globalState)
                d_1118_v43_: _dafny.Array
                nw185_ = _dafny.Array(False, 22)
                d_1118_v43_ = nw185_
                index185_ = default__.safeIndex(325, (d_1118_v43_).length(0))
                (d_1118_v43_)[index185_] = ((self).f19) == (default__.fm48(globalState))
                index186_ = default__.safeIndex(325, (d_1118_v43_).length(0))
                (d_1118_v43_)[index186_] = (self).fm3(len((default__.fm49(708, (self).f30, d_1044_cf15_, globalState)).intersection(_dafny.Set({414, d_1042_cf17_, d_1042_cf17_}))), globalState)
            d_1119_v44_: _dafny.Array
            nw186_ = _dafny.Array(_dafny.Map({}), 16)
            d_1119_v44_ = nw186_
            d_1120_v45_: D3
            d_1120_v45_ = D3_DC10(d_1119_v44_)
            pat_let_tv21_ = d_1119_v44_
            def iife93_(_pat_let30_0):
                def iife94_(d_1121_dt__update__tmp_h0_):
                    def iife95_(_pat_let31_0):
                        def iife96_(d_1122_dt__update_hcf22_h0_):
                            return D3_DC10(d_1122_dt__update_hcf22_h0_)
                        return iife96_(_pat_let31_0)
                    return iife95_(pat_let_tv21_)
                return iife94_(_pat_let30_0)
            source11_ = iife93_(d_1120_v45_)
            if source11_.is_DC11:
                d_1123___mcc_h22_ = source11_.cf23
                d_1124___mcc_h23_ = source11_.cf24
                d_1125_cf24_ = d_1124___mcc_h23_
                d_1126_cf23_ = d_1123___mcc_h22_
                d_1127_v46_: _dafny.Map
                d_1127_v46_ = _dafny.Map({p2: d_1125_cf24_})
                d_1043_cf16_ = (d_1125_cf24_) and (((d_1127_v46_)[d_1126_cf23_] if (d_1126_cf23_) in (d_1127_v46_) else (self).f18))
                (globalState).f2 = p2
                d_1128_v47_: _dafny.Array
                nw187_ = _dafny.Array(int(0), 8)
                d_1128_v47_ = nw187_
                d_1129_v48_: _dafny.Set
                d_1129_v48_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([(self).f29]))})
                d_1130_v50_: _dafny.Map
                def iife97_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in (d_1129_v48_).Elements:
                        d_1131_v49_: int = compr_33_
                        if (d_1131_v49_) in (d_1129_v48_):
                            coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_1131_v49_, len(_dafny.SeqWithoutIsStrInference([False])))]))
                    return _dafny.Set(coll33_)
                d_1130_v50_ = _dafny.Map({d_1128_v47_: iife97_()
                })
                d_1130_v50_ = (d_1130_v50_) | (d_1130_v50_)
                d_1132_v51_: C1
                nw188_ = C1()
                nw188_.ctor__(((0) - (d_1044_cf15_)) < (len((self).f19)))
                d_1132_v51_ = nw188_
            elif source11_.is_DC12:
                d_1133___mcc_h24_ = source11_.cf25
                d_1134___mcc_h25_ = source11_.cf26
                d_1135___mcc_h26_ = source11_.cf27
                d_1136___mcc_h27_ = source11_.cf28
                d_1137___mcc_h28_ = source11_.cf29
                d_1138_cf29_ = d_1137___mcc_h28_
                d_1139_cf28_ = d_1136___mcc_h27_
                d_1140_cf27_ = d_1135___mcc_h26_
                d_1141_cf26_ = d_1134___mcc_h25_
                d_1142_cf25_ = d_1133___mcc_h24_
                d_1143_v52_: D2
                d_1143_v52_ = D2_DC8((self).f18, False, not(True))
                d_1144_v53_: D2
                d_1144_v53_ = D2_DC9(d_1143_v52_)
                d_1145_v54_: _dafny.MultiSet
                d_1145_v54_ = _dafny.MultiSet([d_1138_cf29_])
                d_1146_v55_: _dafny.MultiSet
                d_1146_v55_ = _dafny.MultiSet([d_1044_cf15_, (d_1145_v54_).cardinality])
                d_1147_v56_: _dafny.Map
                d_1147_v56_ = _dafny.Map({(self).f30: (d_1145_v54_).cardinality})
                d_1148_v57_: _dafny.Seq
                d_1148_v57_ = _dafny.SeqWithoutIsStrInference([d_1146_v55_])
                d_1149_v58_: _dafny.Map
                d_1149_v58_ = _dafny.Map({(_dafny.Map({d_1042_cf17_: (self).fm18(globalState)})) != (d_1147_v56_): (d_1148_v57_)[default__.safeIndex(d_1042_cf17_, len(d_1148_v57_))]})
                rhs185_ = d_1144_v53_
                rhs186_ = d_1138_cf29_
                rhs187_ = ((d_1149_v58_)[d_1141_cf26_] if (d_1141_cf26_) in (d_1149_v58_) else d_1146_v55_)
                lhs165_ = globalState
                d_1144_v53_ = rhs185_
                lhs165_.f2 = rhs186_
                d_1146_v55_ = rhs187_
                d_1150_v59_: _dafny.Set
                d_1150_v59_ = _dafny.Set({True})
                d_1151_v60_: _dafny.Map
                d_1151_v60_ = _dafny.Map({D2_DC8(d_1043_cf16_, d_1043_cf16_, False): p2})
                d_1152_v61_: str
                d_1152_v61_ = _dafny.CodePoint('c')
                rhs188_ = d_1150_v59_
                rhs189_ = d_1151_v60_
                rhs190_ = (d_1140_cf27_) == (d_1044_cf15_)
                rhs191_ = ((self).f29 if p1 else (d_1152_v61_ if (self).f18 else (self).f29))
                rhs192_ = (0) - (d_1044_cf15_)
                lhs166_ = globalState
                d_1150_v59_ = rhs188_
                d_1151_v60_ = rhs189_
                d_1138_cf29_ = rhs190_
                d_1152_v61_ = rhs191_
                lhs166_.f3 = rhs192_
                d_1153_v62_: C0
                nw189_ = C0()
                nw189_.ctor__((self).f30, (self).f30)
                d_1153_v62_ = nw189_
                (globalState).f3 = (0) - (d_1042_cf17_)
            elif source11_.is_DC13:
                d_1154___mcc_h29_ = source11_.cf30
                d_1155___mcc_h30_ = source11_.cf31
                d_1156___mcc_h31_ = source11_.cf32
                d_1157___mcc_h32_ = source11_.cf33
                d_1158___mcc_h33_ = source11_.cf34
                d_1159_cf34_ = d_1158___mcc_h33_
                d_1160_cf33_ = d_1157___mcc_h32_
                d_1161_cf32_ = d_1156___mcc_h31_
                d_1162_cf31_ = d_1155___mcc_h30_
                d_1163_cf30_ = d_1154___mcc_h29_
                d_1164_v63_: _dafny.Array
                def lambda46_(d_1165_cf33_):
                    def lambda47_(d_1166_i2_):
                        return (d_1166_i2_) - (len(_dafny.Map({d_1165_cf33_: -538})))

                    return lambda47_

                init26_ = lambda46_(d_1160_cf33_)
                nw190_ = _dafny.Array(None, 21)
                for i0_26_ in range(nw190_.length(0)):
                    nw190_[i0_26_] = init26_(i0_26_)
                d_1164_v63_ = nw190_
                d_1164_v63_ = d_1164_v63_
                d_1167_v64_: _dafny.MultiSet
                d_1167_v64_ = _dafny.MultiSet([(self).f30])
                d_1168_v65_: _dafny.Seq
                d_1168_v65_ = _dafny.SeqWithoutIsStrInference([(d_1167_v64_).ispropersubset(d_1167_v64_)])
                d_1168_v65_ = ((d_1168_v65_ if d_1043_cf16_ else d_1168_v65_)).set(default__.safeIndex(243, len((d_1168_v65_ if d_1043_cf16_ else d_1168_v65_))), p1)
                d_1043_cf16_ = p2
                d_1169_v66_: _dafny.Seq
                d_1169_v66_ = _dafny.SeqWithoutIsStrInference([547, d_1161_cf32_, d_1163_cf30_])
                d_1170_v67_: C3
                nw191_ = C3()
                nw191_.ctor__(d_1169_v66_, d_1162_cf31_, _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1163_cf30_: (_dafny.MultiSet([(self).f29])).cardinality}) for d_1171_i3_ in range(default__.abs(728))]), p2)
                d_1170_v67_ = nw191_
            elif source11_.is_DC10:
                d_1172___mcc_h34_ = source11_.cf22
                d_1173_cf22_ = d_1172___mcc_h34_
                d_1174_v68_: str
                d_1174_v68_ = _dafny.CodePoint('v')
                d_1174_v68_ = (d_1174_v68_ if p2 else d_1174_v68_)
                d_1175_v69_: C0
                nw192_ = C0()
                nw192_.ctor__((self).f30, d_1044_cf15_)
                d_1175_v69_ = nw192_
                d_1176_v70_: _dafny.Map
                d_1176_v70_ = _dafny.Map({not((self).f18): False})
                d_1177_v71_: _dafny.Set
                d_1177_v71_ = _dafny.Set({(d_1175_v69_).f32, len(d_1176_v70_)})
                d_1178_v72_: C0
                nw193_ = C0()
                nw193_.ctor__((d_1175_v69_).f31, len(default__.fm50(((d_1176_v70_)[default__.fm23(d_1177_v71_, (d_1175_v69_).f31, globalState)] if (default__.fm23(d_1177_v71_, (d_1175_v69_).f31, globalState)) in (d_1176_v70_) else p2), (self).f19, d_1176_v70_, d_1174_v68_, globalState)))
                d_1178_v72_ = nw193_
                d_1179_v73_: _dafny.Array
                def lambda48_(d_1180_p2_, d_1181_v72_):
                    def lambda49_(d_1182_i4_):
                        return default__.safeModuloInt(d_1182_i4_, len(_dafny.Map({d_1180_p2_: (d_1181_v72_).f31})))

                    return lambda49_

                init27_ = lambda48_(p2, d_1178_v72_)
                nw194_ = _dafny.Array(None, 5)
                for i0_27_ in range(nw194_.length(0)):
                    nw194_[i0_27_] = init27_(i0_27_)
                d_1179_v73_ = nw194_
                d_1183_v74_: _dafny.Seq
                d_1183_v74_ = _dafny.SeqWithoutIsStrInference([not(d_1043_cf16_)])
                d_1184_v75_: _dafny.MultiSet
                d_1184_v75_ = _dafny.MultiSet([True])
                d_1185_v76_: _dafny.Array
                nw195_ = _dafny.Array(None, 27)
                nw195_[int(0)] = (self).f29
                nw195_[int(1)] = d_1174_v68_
                nw195_[int(2)] = (self).f29
                nw195_[int(3)] = (self).f29
                nw195_[int(4)] = d_1174_v68_
                nw195_[int(5)] = d_1174_v68_
                nw195_[int(6)] = d_1174_v68_
                nw195_[int(7)] = (self).f29
                nw195_[int(8)] = (self).f29
                nw195_[int(9)] = d_1174_v68_
                nw195_[int(10)] = d_1174_v68_
                nw195_[int(11)] = (self).f29
                nw195_[int(12)] = (self).f29
                nw195_[int(13)] = (self).f29
                nw195_[int(14)] = (self).f29
                nw195_[int(15)] = d_1174_v68_
                nw195_[int(16)] = (self).f29
                nw195_[int(17)] = _dafny.CodePoint('y')
                nw195_[int(18)] = d_1174_v68_
                nw195_[int(19)] = (self).f29
                nw195_[int(20)] = d_1174_v68_
                nw195_[int(21)] = d_1174_v68_
                nw195_[int(22)] = (self).f29
                nw195_[int(23)] = (self).f29
                nw195_[int(24)] = (self).f29
                nw195_[int(25)] = d_1174_v68_
                nw195_[int(26)] = d_1174_v68_
                d_1185_v76_ = nw195_
                d_1186_v77_: D3
                d_1186_v77_ = D3_DC12(d_1185_v76_, p1, default__.fm1((d_1184_v75_).cardinality, p2, True, globalState), _dafny.Map({d_1184_v75_: (self).f18}), p2)
                d_1187_v78_: _dafny.Set
                d_1187_v78_ = _dafny.Set({(self).f29})
                d_1188_v79_: _dafny.Set
                d_1188_v79_ = _dafny.Set({False})
                d_1189_v80_: _dafny.Seq
                d_1189_v80_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                d_1190_v81_: _dafny.Seq
                d_1190_v81_ = _dafny.SeqWithoutIsStrInference([(d_1189_v80_)[default__.safeIndex((d_1175_v69_).f31, len(d_1189_v80_))]])
                d_1191_v82_: _dafny.Array
                nw196_ = _dafny.Array(None, 24)
                nw196_[int(0)] = p2
                nw196_[int(1)] = p1
                nw196_[int(2)] = p2
                nw196_[int(3)] = (d_1183_v74_) <= (default__.fm44((d_1184_v75_).cardinality, p1, d_1174_v68_, (d_1178_v72_).f32, globalState))
                nw196_[int(4)] = d_1043_cf16_
                nw196_[int(5)] = ((d_1184_v75_)).ispropersubset(_dafny.MultiSet(d_1183_v74_))
                nw196_[int(6)] = p1
                nw196_[int(7)] = p1
                nw196_[int(8)] = not(not ((self).f18) or ((self).f18))
                nw196_[int(9)] = p2
                nw196_[int(10)] = True
                nw196_[int(11)] = (self).f18
                nw196_[int(12)] = ((d_1175_v69_).f31) > (len(_dafny.Map({(d_1186_v77_).cf27: d_1187_v78_})))
                nw196_[int(13)] = (d_1188_v79_).issubset(d_1188_v79_)
                nw196_[int(14)] = not((d_1183_v74_)[default__.safeIndex(d_1042_cf17_, len(d_1183_v74_))])
                nw196_[int(15)] = not(p2)
                nw196_[int(16)] = d_1043_cf16_
                nw196_[int(17)] = ((0) - ((d_1178_v72_).f31)) < ((0) - ((d_1178_v72_).f31))
                nw196_[int(18)] = ((self).f19) < ((self).f19)
                nw196_[int(19)] = (self).f18
                nw196_[int(20)] = ((self).f19) not in (d_1190_v81_)
                nw196_[int(21)] = True
                nw196_[int(22)] = p1
                nw196_[int(23)] = p1
                d_1191_v82_ = nw196_
                index187_ = default__.safeIndex(995, (d_1191_v82_).length(0))
                (d_1191_v82_)[index187_] = p1
                index188_ = default__.safeIndex(995, (d_1191_v82_).length(0))
                rhs193_ = _dafny.CodePoint('v')
                rhs194_ = default__.fm1(d_1044_cf15_, ((0) - ((d_1178_v72_).f32)) == ((d_1178_v72_).f32), d_1043_cf16_, globalState)
                rhs195_ = ((d_1178_v72_).f31 if p2 else d_1042_cf17_)
                rhs196_ = d_1179_v73_
                rhs197_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tv"))) != ((self).f19)
                lhs167_ = globalState
                lhs168_ = globalState
                lhs169_ = d_1191_v82_
                lhs170_ = default__.safeIndex(995, (d_1191_v82_).length(0))
                d_1174_v68_ = rhs193_
                lhs167_.f0 = rhs194_
                lhs168_.f15 = rhs195_
                d_1179_v73_ = rhs196_
                lhs169_[lhs170_] = rhs197_
            elif True:
                d_1192___mcc_h35_ = source11_.cf35
                d_1193_cf35_ = d_1192___mcc_h35_
                (globalState).f0 = d_1042_cf17_
                d_1043_cf16_ = not ((self).f18) or (False)
                (globalState).f2 = (self).f18
                d_1194_v83_: _dafny.Seq
                d_1194_v83_ = _dafny.SeqWithoutIsStrInference([False, p2, False])
                d_1195_v84_: _dafny.Seq
                d_1195_v84_ = _dafny.SeqWithoutIsStrInference([d_1194_v83_])
                (globalState).f2 = (default__.safeDivisionInt(d_1044_cf15_, d_1042_cf17_)) <= (len((d_1195_v84_)[default__.safeIndex(len((self).f19), len(d_1195_v84_))]))
            (globalState).f9 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqwb"))
            d_1196_v85_: _dafny.Seq
            d_1196_v85_ = _dafny.SeqWithoutIsStrInference([d_1044_cf15_, d_1044_cf15_, 148, 505, d_1044_cf15_])
            d_1197_v86_: C3
            nw197_ = C3()
            nw197_.ctor__(d_1196_v85_, (self).f19, _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1042_cf17_: (self).f30}) for d_1198_i5_ in range(default__.abs(935))]), d_1043_cf16_)
            d_1197_v86_ = nw197_
        elif source9_.is_DC8:
            d_1199___mcc_h3_ = source9_.cf18
            d_1200___mcc_h4_ = source9_.cf19
            d_1201___mcc_h5_ = source9_.cf20
            d_1202_cf20_ = d_1201___mcc_h5_
            d_1203_cf19_ = d_1200___mcc_h4_
            d_1204_cf18_ = d_1199___mcc_h3_
            d_1205_v87_: _dafny.Map
            d_1205_v87_ = _dafny.Map({(self).f30: (self).f30})
            d_1206_v88_: _dafny.Set
            d_1206_v88_ = _dafny.Set({((d_1205_v87_)[(self).f30] if ((self).f30) in (d_1205_v87_) else (0) - ((self).f30)), 391, (self).f30, (self).f30})
            d_1207_v89_: _dafny.Set
            d_1207_v89_ = _dafny.Set({(_dafny.Set({(self).f30, (self).f30})).isdisjoint(d_1206_v88_), (self).f18})
            d_1208_v90_: str
            d_1208_v90_ = _dafny.CodePoint('l')
            d_1209_v91_: _dafny.Seq
            d_1209_v91_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1210_v92_: _dafny.Map
            d_1210_v92_ = _dafny.Map({(d_1206_v88_ if (self).fm2((self).f18, True, False, globalState) else _dafny.Set({len(d_1209_v91_)})): default__.fm23(d_1206_v88_, (self).f30, globalState)})
            rhs198_ = default__.fm23(d_1206_v88_, (0) - ((self).f30), globalState)
            rhs199_ = d_1207_v89_
            rhs200_ = (self).f18
            rhs201_ = (self).f29
            rhs202_ = d_1210_v92_
            lhs171_ = globalState
            lhs172_ = globalState
            lhs171_.f2 = rhs198_
            d_1207_v89_ = rhs199_
            lhs172_.f2 = rhs200_
            d_1208_v90_ = rhs201_
            d_1210_v92_ = rhs202_
            d_1211_v93_: _dafny.Map
            d_1211_v93_ = _dafny.Map({d_1208_v90_: len(d_1205_v87_)})
            d_1212_v94_: _dafny.Map
            d_1212_v94_ = _dafny.Map({len(d_1211_v93_): (self).f18})
            d_1213_v95_: _dafny.Map
            d_1213_v95_ = _dafny.Map({p2: d_1202_cf20_})
            d_1214_v96_: _dafny.Array
            nw198_ = _dafny.Array(None, 20)
            nw198_[int(0)] = d_1202_cf20_
            nw198_[int(1)] = d_1202_cf20_
            nw198_[int(2)] = d_1204_cf18_
            nw198_[int(3)] = d_1203_cf19_
            nw198_[int(4)] = d_1203_cf19_
            nw198_[int(5)] = p2
            nw198_[int(6)] = (self).f18
            nw198_[int(7)] = p1
            nw198_[int(8)] = not((self).f18)
            nw198_[int(9)] = ((d_1212_v94_)[(0) - (default__.fm1((self).f30, d_1202_cf20_, d_1204_cf18_, globalState))] if ((0) - (default__.fm1((self).f30, d_1202_cf20_, d_1204_cf18_, globalState))) in (d_1212_v94_) else d_1202_cf20_)
            nw198_[int(10)] = p2
            nw198_[int(11)] = d_1204_cf18_
            nw198_[int(12)] = (d_1209_v91_)[default__.safeIndex(len((self).f19), len(d_1209_v91_))]
            nw198_[int(13)] = d_1204_cf18_
            nw198_[int(14)] = ((d_1213_v95_)[p2] if (p2) in (d_1213_v95_) else d_1203_cf19_)
            nw198_[int(15)] = p2
            nw198_[int(16)] = d_1202_cf20_
            nw198_[int(17)] = d_1203_cf19_
            nw198_[int(18)] = (self).f18
            nw198_[int(19)] = not(d_1202_cf20_)
            d_1214_v96_ = nw198_
            d_1215_v97_: _dafny.Set
            d_1215_v97_ = _dafny.Set({d_1214_v96_})
            d_1216_v98_: _dafny.Seq
            d_1216_v98_ = _dafny.SeqWithoutIsStrInference([416, (self).f30])
            d_1217_v99_: C2
            nw199_ = C2()
            nw199_.ctor__(len((d_1215_v97_).intersection(d_1215_v97_)), d_1216_v98_, (d_1207_v89_).ispropersubset(_dafny.Set({p2})))
            d_1217_v99_ = nw199_
            d_1218_v100_: C2
            nw200_ = C2()
            nw200_.ctor__(len((self).f19), (d_1217_v99_).f34, d_1202_cf20_)
            d_1218_v100_ = nw200_
            d_1219_v101_: T0
            nw201_ = C2()
            nw201_.ctor__((d_1218_v100_.f33) * (len(_dafny.SeqWithoutIsStrInference([p2, p1]))), (d_1218_v100_).f34, p1)
            d_1219_v101_ = nw201_
            d_1220_v102_: D11
            d_1220_v102_ = D11_DC28(((self).f18 if True else d_1204_cf18_))
            rhs203_ = d_1209_v91_
            rhs204_ = (0) - (d_1218_v100_.f33)
            rhs205_ = d_1219_v101_
            rhs206_ = d_1220_v102_
            lhs173_ = d_1217_v99_
            d_1209_v91_ = rhs203_
            lhs173_.f33 = rhs204_
            d_1219_v101_ = rhs205_
            d_1220_v102_ = rhs206_
        elif source9_.is_DC6:
            d_1221___mcc_h6_ = source9_.cf14
            d_1222_cf14_ = d_1221___mcc_h6_
            d_1223_v103_: _dafny.Array
            def lambda50_(d_1224_p2_):
                def lambda51_(d_1225_i6_):
                    return d_1224_p2_

                return lambda51_

            init28_ = lambda50_(p2)
            nw202_ = _dafny.Array(None, 26)
            for i0_28_ in range(nw202_.length(0)):
                nw202_[i0_28_] = init28_(i0_28_)
            d_1223_v103_ = nw202_
            d_1226_v104_: _dafny.Array
            def lambda52_(d_1227_p1_):
                def lambda53_(d_1228_i7_):
                    return D9_DC23(d_1227_p1_, False, (self).f30)

                return lambda53_

            init29_ = lambda52_(p1)
            nw203_ = _dafny.Array(None, 27)
            for i0_29_ in range(nw203_.length(0)):
                nw203_[i0_29_] = init29_(i0_29_)
            d_1226_v104_ = nw203_
            d_1229_v105_: _dafny.Map
            d_1229_v105_ = _dafny.Map({d_1226_v104_: d_1223_v103_})
            d_1223_v103_ = ((d_1229_v105_)[d_1226_v104_] if (d_1226_v104_) in (d_1229_v105_) else d_1223_v103_)
            rhs207_ = (self).f18
            rhs208_ = ((self).f30) * (default__.safeModuloInt((self).f30, (self).f30))
            lhs174_ = globalState
            lhs175_ = globalState
            lhs174_.f2 = rhs207_
            lhs175_.f3 = rhs208_
            d_1230_v106_: _dafny.MultiSet
            d_1230_v106_ = _dafny.MultiSet([(self).f30])
            d_1231_v107_: _dafny.Set
            d_1231_v107_ = _dafny.Set({d_1230_v106_, d_1230_v106_})
            d_1232_v108_: _dafny.Map
            d_1232_v108_ = _dafny.Map({(self).f30: d_1230_v106_})
            d_1231_v107_ = (d_1231_v107_) | (_dafny.Set({d_1230_v106_, ((d_1232_v108_)[(self).f30] if ((self).f30) in (d_1232_v108_) else d_1230_v106_)}))
            d_1233_v109_: str
            d_1233_v109_ = _dafny.CodePoint('l')
            d_1233_v109_ = ((self).f19)[default__.safeIndex((len(default__.fm48(globalState))) + ((self).f30), len((self).f19))]
        elif True:
            d_1234___mcc_h7_ = source9_.cf21
            d_1235_cf21_ = d_1234___mcc_h7_
            d_1236_v110_: _dafny.MultiSet
            d_1236_v110_ = _dafny.MultiSet([(self).f30])
            d_1237_v111_: _dafny.Seq
            d_1237_v111_ = _dafny.SeqWithoutIsStrInference([(self).f30, len((self).f19), (self).f30, (d_1236_v110_).cardinality])
            d_1238_v112_: C0
            nw204_ = C0()
            nw204_.ctor__((self).f30, (d_1237_v111_)[default__.safeIndex((self).f30, len(d_1237_v111_))])
            d_1238_v112_ = nw204_
            d_1239_v113_: C1
            nw205_ = C1()
            nw205_.ctor__(p1)
            d_1239_v113_ = nw205_
            d_1240_v114_: _dafny.Array
            nw206_ = _dafny.Array(False, 19)
            d_1240_v114_ = nw206_
            d_1241_v115_: D0
            d_1241_v115_ = D0_DC1(p1, (d_1238_v112_).f32, _dafny.SeqWithoutIsStrInference([(self).f29 for d_1242_i8_ in range(default__.abs(78))]), True, (d_1238_v112_).f32)
            index189_ = default__.safeIndex(673, (d_1240_v114_).length(0))
            (d_1240_v114_)[index189_] = (d_1241_v115_).cf4
            index190_ = default__.safeIndex(673, (d_1240_v114_).length(0))
            (d_1240_v114_)[index190_] = (p2) or (p1)
            d_1243_v116_: _dafny.Seq
            d_1243_v116_ = _dafny.SeqWithoutIsStrInference([(d_1240_v114_)[default__.safeIndex(673, (d_1240_v114_).length(0))]])
            index191_ = default__.safeIndex(673, (d_1240_v114_).length(0))
            rhs209_ = p1
            rhs210_ = (_dafny.SeqWithoutIsStrInference([p1])) + (d_1243_v116_)
            lhs176_ = d_1240_v114_
            lhs177_ = default__.safeIndex(673, (d_1240_v114_).length(0))
            lhs176_[lhs177_] = rhs209_
            d_1243_v116_ = rhs210_
        d_1244_i9_: int
        d_1244_i9_ = 0
        with _dafny.label("9"):
            while (self).f18:
                with _dafny.c_label("9"):
                    if (d_1244_i9_) >= (100):
                        raise _dafny.Break("9")
                    d_1244_i9_ = (d_1244_i9_) + (1)
                    d_1245_v117_: C1
                    nw207_ = C1()
                    nw207_.ctor__(not(False))
                    d_1245_v117_ = nw207_
                    d_1246_v119_: _dafny.Map
                    def iife98_():
                        coll34_ = _dafny.Map()
                        compr_34_: str
                        for compr_34_ in (((self).f19).set(default__.safeIndex((self).f30, len((self).f19)), _dafny.CodePoint('r'))).Elements:
                            d_1247_v118_: str = compr_34_
                            if (d_1247_v118_) in (((self).f19).set(default__.safeIndex((self).f30, len((self).f19)), _dafny.CodePoint('r'))):
                                coll34_[d_1247_v118_] = (self).f19
                        return _dafny.Map(coll34_)
                    d_1246_v119_ = _dafny.Map({iife98_()
                    : d_1245_v117_})
                    d_1248_v120_: _dafny.Map
                    d_1248_v120_ = _dafny.Map({(self).f29: (self).f19})
                    d_1249_v121_: _dafny.Array
                    nw208_ = _dafny.Array(None, 28)
                    nw208_[int(0)] = d_1245_v117_
                    nw208_[int(1)] = d_1245_v117_
                    nw208_[int(2)] = ((d_1246_v119_)[d_1248_v120_] if (d_1248_v120_) in (d_1246_v119_) else d_1245_v117_)
                    nw208_[int(3)] = d_1245_v117_
                    nw208_[int(4)] = d_1245_v117_
                    nw208_[int(5)] = d_1245_v117_
                    nw208_[int(6)] = d_1245_v117_
                    nw208_[int(7)] = d_1245_v117_
                    nw208_[int(8)] = d_1245_v117_
                    nw208_[int(9)] = d_1245_v117_
                    nw208_[int(10)] = d_1245_v117_
                    nw208_[int(11)] = d_1245_v117_
                    nw208_[int(12)] = d_1245_v117_
                    nw208_[int(13)] = d_1245_v117_
                    nw208_[int(14)] = d_1245_v117_
                    nw208_[int(15)] = d_1245_v117_
                    nw208_[int(16)] = d_1245_v117_
                    nw208_[int(17)] = d_1245_v117_
                    nw208_[int(18)] = d_1245_v117_
                    nw208_[int(19)] = d_1245_v117_
                    nw208_[int(20)] = d_1245_v117_
                    nw208_[int(21)] = d_1245_v117_
                    nw208_[int(22)] = d_1245_v117_
                    nw208_[int(23)] = (d_1245_v117_ if p2 else d_1245_v117_)
                    nw208_[int(24)] = d_1245_v117_
                    nw208_[int(25)] = d_1245_v117_
                    nw208_[int(26)] = d_1245_v117_
                    nw208_[int(27)] = d_1245_v117_
                    d_1249_v121_ = nw208_
                    index192_ = default__.safeIndex(465, (d_1249_v121_).length(0))
                    (d_1249_v121_)[index192_] = d_1245_v117_
                    index193_ = default__.safeIndex(465, (d_1249_v121_).length(0))
                    (d_1249_v121_)[index193_] = d_1245_v117_
                    d_1250_v122_: _dafny.Array
                    nw209_ = _dafny.Array(int(0), 16)
                    d_1250_v122_ = nw209_
                    d_1250_v122_ = d_1250_v122_
                    d_1251_v123_: _dafny.Map
                    d_1251_v123_ = _dafny.Map({p1: p2})
                    (globalState).f2 = not (((d_1251_v123_)[True] if (True) in (d_1251_v123_) else p1)) or (False)
                    d_1252_v124_: _dafny.MultiSet
                    d_1252_v124_ = _dafny.MultiSet([(self).f30])
                    d_1252_v124_ = (d_1252_v124_) | ((d_1252_v124_ if (self).f18 else d_1252_v124_))
                    pass
            pass
        d_1253_v125_: _dafny.Set
        d_1253_v125_ = _dafny.Set({p2, not(True), p1, p1, (self).f18})
        d_1254_v126_: _dafny.Seq
        d_1254_v126_ = _dafny.SeqWithoutIsStrInference([len(d_1253_v125_), (self).f30, 603, (0) - ((self).f30)])
        d_1255_v127_: _dafny.Seq
        d_1255_v127_ = _dafny.SeqWithoutIsStrInference([p2, p1])
        d_1256_v128_: D10
        d_1256_v128_ = D10_DC24(d_1253_v125_)
        d_1257_v129_: C3
        nw210_ = C3()
        nw210_.ctor__(d_1254_v126_, (self).f19, default__.fm51(p2, (self).f30, (d_1255_v127_)[default__.safeIndex((0) - ((self).f30), len(d_1255_v127_))], d_1256_v128_, globalState), not (p1) or ((self).f18))
        d_1257_v129_ = nw210_
        d_1258_v130_: _dafny.Array
        nw211_ = _dafny.Array(_dafny.Map({}), 22)
        d_1258_v130_ = nw211_
        d_1259_v131_: _dafny.Map
        d_1259_v131_ = _dafny.Map({(self).f30: (self).f19})
        index194_ = default__.safeIndex(724, (d_1258_v130_).length(0))
        (d_1258_v130_)[index194_] = (d_1259_v131_) | (d_1259_v131_)
        index195_ = default__.safeIndex(724, (d_1258_v130_).length(0))
        (d_1258_v130_)[index195_] = _dafny.Map({default__.safeDivisionInt((self).f30, (self).f30): (self).f19})
        d_1260_v132_: C1
        nw212_ = C1()
        nw212_.ctor__(p1)
        d_1260_v132_ = nw212_
        (globalState).f2 = p2
        r0 = ((self).f30) + (len(d_1255_v127_))
        return r0

    def m8(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1261_v0_: _dafny.Map
        d_1261_v0_ = _dafny.Map({(self).f18: (self).f18})
        d_1262_v1_: _dafny.MultiSet
        d_1262_v1_ = _dafny.MultiSet([((d_1261_v0_)[(self).f18] if ((self).f18) in (d_1261_v0_) else (self).f18)])
        d_1263_v2_: _dafny.Seq
        d_1263_v2_ = _dafny.SeqWithoutIsStrInference([d_1262_v1_])
        d_1264_v3_: _dafny.Map
        d_1264_v3_ = _dafny.Map({p0: ((d_1262_v1_) | ((d_1263_v2_)[default__.safeIndex(p0, len(d_1263_v2_))])).cardinality})
        def iife99_():
            coll35_ = _dafny.Map()
            compr_35_: int
            for compr_35_ in _dafny.IntegerRange(584, 278):
                d_1265_v4_: int = compr_35_
                if ((584) <= (d_1265_v4_)) and ((d_1265_v4_) < (278)):
                    coll35_[default__.safeDivisionInt(d_1265_v4_, (0) - (p0))] = len(_dafny.SeqWithoutIsStrInference([895, p0]))
            return _dafny.Map(coll35_)
        d_1264_v3_ = iife99_()
        
        rhs211_ = (self).f18
        rhs212_ = ((self).f30 if (p0) < ((self).f30) else p0)
        lhs178_ = globalState
        lhs179_ = globalState
        lhs178_.f2 = rhs211_
        lhs179_.f3 = rhs212_
        (globalState).f2 = (self).f18
        d_1266_v5_: _dafny.Map
        d_1266_v5_ = _dafny.Map({(self).f29: (self).f18})
        d_1267_v7_: _dafny.Seq
        def iife100_():
            coll36_ = _dafny.Map()
            compr_36_: str
            for compr_36_ in (_dafny.SeqWithoutIsStrInference([(self).f29])).Elements:
                d_1268_v6_: str = compr_36_
                if (d_1268_v6_) in (_dafny.SeqWithoutIsStrInference([(self).f29])):
                    coll36_[d_1268_v6_] = (self).f18
            return _dafny.Map(coll36_)
        d_1267_v7_ = _dafny.SeqWithoutIsStrInference([d_1266_v5_, iife100_()
        , _dafny.Map({(self).f29: (self).f18}), d_1266_v5_])
        d_1269_v8_: _dafny.Map
        d_1269_v8_ = _dafny.Map({222: d_1264_v3_})
        d_1270_v10_: _dafny.Seq
        d_1270_v10_ = _dafny.SeqWithoutIsStrInference([default__.fm48(globalState)])
        d_1271_v11_: _dafny.MultiSet
        def iife101_():
            coll37_ = _dafny.Map()
            compr_37_: _dafny.Seq
            for compr_37_ in (d_1270_v10_).Elements:
                d_1272_v9_: _dafny.Seq = compr_37_
                if (d_1272_v9_) in (d_1270_v10_):
                    coll37_[d_1272_v9_] = not((self).f18)
            return _dafny.Map(coll37_)
        d_1271_v11_ = _dafny.MultiSet([d_1264_v3_, d_1264_v3_, (((d_1269_v8_)[len(d_1261_v0_)] if (len(d_1261_v0_)) in (d_1269_v8_) else d_1264_v3_)) | (default__.fm30(d_1263_v2_, iife101_()
        , globalState))])
        rhs213_ = ((d_1267_v7_) + (d_1267_v7_)) + ((d_1267_v7_) + (d_1267_v7_))
        rhs214_ = d_1271_v11_
        d_1267_v7_ = rhs213_
        d_1271_v11_ = rhs214_
        d_1273_v12_: _dafny.Seq
        d_1273_v12_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18, (self).f18])
        d_1274_v13_: _dafny.Map
        d_1274_v13_ = _dafny.Map({(d_1273_v12_)[default__.safeIndex(p0, len(d_1273_v12_))]: (self).f29})
        d_1275_v14_: _dafny.Map
        d_1275_v14_ = _dafny.Map({len(d_1274_v13_): (self).f18})
        d_1276_i0_: int
        d_1276_i0_ = 0
        with _dafny.label("10"):
            while (((d_1275_v14_)[p0] if (p0) in (d_1275_v14_) else (self).f18) if (self).f18 else (self).f18):
                with _dafny.c_label("10"):
                    if (d_1276_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1276_i0_ = (d_1276_i0_) + (1)
                    d_1277_v15_: _dafny.Array
                    nw213_ = _dafny.Array(int(0), 26)
                    d_1277_v15_ = nw213_
                    index196_ = default__.safeIndex(460, (d_1277_v15_).length(0))
                    (d_1277_v15_)[index196_] = (self).f30
                    index197_ = default__.safeIndex(460, (d_1277_v15_).length(0))
                    (d_1277_v15_)[index197_] = (self).f30
                    (globalState).f9 = (self).f19
                    d_1278_v16_: str
                    d_1278_v16_ = _dafny.CodePoint('t')
                    d_1278_v16_ = (self).f29
                    (globalState).f2 = not((self).f18)
                    pass
            pass
        d_1279_v17_: _dafny.Seq
        d_1279_v17_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f30), len((self).f19)])
        d_1280_v18_: _dafny.Map
        d_1280_v18_ = _dafny.Map({(self).f18: (d_1279_v17_).set(default__.safeIndex((0) - ((0) - ((0) - ((self).f30))), len(d_1279_v17_)), (self).f30)})
        d_1281_v19_: _dafny.MultiSet
        d_1281_v19_ = _dafny.MultiSet([(0) - (p0)])
        d_1282_v20_: _dafny.Map
        d_1282_v20_ = _dafny.Map({(self).f18: _dafny.Set({(self).f18})})
        d_1283_v21_: D9
        d_1283_v21_ = D9_DC23((self).f18, (self).f18, (self).f30)
        d_1284_v22_: _dafny.Array
        def lambda54_(d_1285_v12_):
            def lambda55_(d_1286_i4_):
                return default__.safeModuloInt(d_1286_i4_, len(d_1285_v12_))

            return lambda55_

        init30_ = lambda54_(d_1273_v12_)
        nw214_ = _dafny.Array(None, 7)
        for i0_30_ in range(nw214_.length(0)):
            nw214_[i0_30_] = init30_(i0_30_)
        d_1284_v22_ = nw214_
        d_1287_v23_: _dafny.Map
        d_1287_v23_ = _dafny.Map({d_1284_v22_: (self).f18})
        d_1288_v24_: _dafny.Map
        d_1288_v24_ = _dafny.Map({(self).f30: _dafny.SeqWithoutIsStrInference([235])})
        d_1289_v25_: _dafny.Map
        d_1289_v25_ = _dafny.Map({(self).f18: len(d_1288_v24_)})
        d_1290_v26_: _dafny.Array
        nw215_ = _dafny.Array(None, 24)
        nw215_[int(0)] = d_1279_v17_
        nw215_[int(1)] = _dafny.SeqWithoutIsStrInference([p0, (self).f30])
        nw215_[int(2)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D2_DC9(D2_DC9(D2_DC7((self).f30, (self).f18, ((d_1262_v1_)[(self).f18] if ((self).f18) in (d_1262_v1_) else 62)))), D2_DC9(D2_DC7((self).f30, (self).f18, len(d_1264_v3_))), D2_DC9(D2_DC9(D2_DC9(D2_DC7(len((self).f19), (self).f18, p0)))), D2_DC9(D2_DC9(D2_DC8((self).f18, (self).f18, (self).f18))), D2_DC9(D2_DC6(_dafny.Map({(self).f18: p0})))])) for d_1291_i2_ in range(default__.abs(-488))])
        nw215_[int(3)] = d_1279_v17_
        nw215_[int(4)] = _dafny.SeqWithoutIsStrInference([(self).f30])
        nw215_[int(5)] = d_1279_v17_
        nw215_[int(6)] = (d_1279_v17_) + (d_1279_v17_)
        nw215_[int(7)] = d_1279_v17_
        nw215_[int(8)] = d_1279_v17_
        nw215_[int(9)] = (d_1279_v17_) + (d_1279_v17_)
        nw215_[int(10)] = (d_1279_v17_).set(default__.safeIndex(758, len(d_1279_v17_)), p0)
        nw215_[int(11)] = _dafny.SeqWithoutIsStrInference([(self).f30 for d_1292_i3_ in range(default__.abs(656))])
        nw215_[int(12)] = d_1279_v17_
        nw215_[int(13)] = _dafny.SeqWithoutIsStrInference([p0])
        nw215_[int(14)] = (d_1279_v17_).set(default__.safeIndex((self).f30, len(d_1279_v17_)), p0)
        nw215_[int(15)] = ((d_1280_v18_)[True] if (True) in (d_1280_v18_) else _dafny.SeqWithoutIsStrInference([p0]))
        nw215_[int(16)] = d_1279_v17_
        nw215_[int(17)] = d_1279_v17_
        nw215_[int(18)] = d_1279_v17_
        nw215_[int(19)] = d_1279_v17_
        nw215_[int(20)] = d_1279_v17_
        nw215_[int(21)] = ((_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30, (d_1281_v19_).cardinality])).set(default__.safeIndex(len(d_1282_v20_), len(_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30, (d_1281_v19_).cardinality]))), (d_1283_v21_).cf46)).set(default__.safeIndex(len(d_1287_v23_), len((_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30, (d_1281_v19_).cardinality])).set(default__.safeIndex(len(d_1282_v20_), len(_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30, (d_1281_v19_).cardinality]))), (d_1283_v21_).cf46))), ((d_1289_v25_)[(self).f18] if ((self).f18) in (d_1289_v25_) else (d_1279_v17_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f18])), len(d_1279_v17_))]))
        nw215_[int(22)] = d_1279_v17_
        nw215_[int(23)] = (d_1279_v17_) + (_dafny.SeqWithoutIsStrInference([p0]))
        d_1290_v26_ = nw215_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1290_v26_).length(0)):
            d_1293_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_1293_i1_)) and ((d_1293_i1_) < ((d_1290_v26_).length(0)))):
                (d_1290_v26_)[(d_1293_i1_)] = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f30, p0), p0])
        r0 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1294_i5_ in range(default__.abs(431))])
        return r0

    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30

class C5(T0):
    def  __init__(self):
        self._f18: bool = False
        self.f27: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f28: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f27, f28, f18):
        (self).f27 = f27
        (self)._f28 = f28
        (self)._f18 = f18

    def fm14(self, p0, p1, p2, globalState):
        return (self).f18

    def fm15(self, p0, p1, globalState):
        return not ((461) != (len(_dafny.Map({(self).f18: (self).f18})))) or (False)

    def m1(self, p0, p1, p2, globalState):
        (globalState).f2 = (self).f18
        d_1295_v0_: D0
        d_1295_v0_ = D0_DC0(True)
        d_1296_v1_: _dafny.Map
        d_1296_v1_ = _dafny.Map({(self).f18: (0) - (p1)})
        d_1297_v2_: _dafny.Map
        d_1297_v2_ = _dafny.Map({(d_1295_v0_).cf0: d_1296_v1_})
        d_1298_v3_: _dafny.Map
        d_1298_v3_ = _dafny.Map({p1: len(((d_1297_v2_)[(self).f18] if ((self).f18) in (d_1297_v2_) else d_1296_v1_))})
        d_1299_v4_: _dafny.Seq
        d_1299_v4_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
        rhs215_ = (len(d_1298_v3_)) + (p1)
        rhs216_ = default__.safeDivisionInt(p1, len(d_1299_v4_))
        rhs217_ = (self).f18
        rhs218_ = (self).f18
        rhs219_ = (self).f18
        lhs180_ = globalState
        lhs181_ = globalState
        lhs182_ = globalState
        lhs183_ = globalState
        lhs184_ = globalState
        lhs180_.f0 = rhs215_
        lhs181_.f3 = rhs216_
        lhs182_.f2 = rhs217_
        lhs183_.f2 = rhs218_
        lhs184_.f2 = rhs219_
        if (p1) <= (582):
            (globalState).f15 = p1
            d_1300_v5_: _dafny.Map
            d_1300_v5_ = _dafny.Map({(self).f18: (self).f18})
            (globalState).f2 = ((len(d_1300_v5_)) * ((0) - (default__.fm1(p1, (self).f18, (self).f18, globalState)))) != (p1)
            (globalState).f2 = (self).f18
            d_1301_v6_: _dafny.Array
            nw216_ = _dafny.Array(False, 18)
            d_1301_v6_ = nw216_
            index198_ = default__.safeIndex(52, (d_1301_v6_).length(0))
            (d_1301_v6_)[index198_] = (self).f18
            index199_ = default__.safeIndex(52, (d_1301_v6_).length(0))
            (d_1301_v6_)[index199_] = not((d_1299_v4_) <= ((d_1299_v4_) + (d_1299_v4_)))
            if (d_1299_v4_)[default__.safeIndex(p1, len(d_1299_v4_))]:
                d_1302_v7_: str
                d_1302_v7_ = _dafny.CodePoint('w')
                d_1303_v8_: _dafny.Set
                d_1303_v8_ = _dafny.Set({p1})
                d_1304_v9_: _dafny.Array
                nw217_ = _dafny.Array(None, 24)
                nw217_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))).set(default__.safeIndex(224, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))), d_1302_v7_)
                nw217_[int(1)] = default__.fm16(p1, (self).f18, (self.f27)[default__.safeIndex(p1, len(self.f27))], p1, globalState)
                nw217_[int(2)] = p2
                nw217_[int(3)] = p2
                nw217_[int(4)] = self.f27
                nw217_[int(5)] = p0
                nw217_[int(6)] = self.f27
                nw217_[int(7)] = p0
                nw217_[int(8)] = (self).f28
                nw217_[int(9)] = (self).f28
                nw217_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))) + ((self).f28)
                nw217_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdjcvu"))
                nw217_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                nw217_[int(13)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxv"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxv")))), d_1302_v7_)) + (p0)
                nw217_[int(14)] = (p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tclvbcm")))
                nw217_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                nw217_[int(16)] = self.f27
                nw217_[int(17)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsde"))) + ((self).f28)
                nw217_[int(18)] = (p2) + ((self).f28)
                nw217_[int(19)] = (self).f28
                nw217_[int(20)] = ((self).f28 if False else p2)
                nw217_[int(21)] = _dafny.SeqWithoutIsStrInference([d_1302_v7_ for d_1305_i0_ in range(default__.abs(57))])
                nw217_[int(22)] = (self).f28
                nw217_[int(23)] = (self.f27).set(default__.safeIndex(len(d_1303_v8_), len(self.f27)), d_1302_v7_)
                d_1304_v9_ = nw217_
                index200_ = default__.safeIndex(94, (d_1304_v9_).length(0))
                (d_1304_v9_)[index200_] = p2
                d_1306_v11_: _dafny.Map
                d_1306_v11_ = _dafny.Map({p1: True})
                d_1307_v13_: _dafny.Seq
                d_1307_v13_ = _dafny.SeqWithoutIsStrInference([p1, (0) - (p1), p1, p1, p1])
                d_1308_v14_: _dafny.Map
                d_1308_v14_ = _dafny.Map({p1: _dafny.Map({p1: True})})
                d_1309_v15_: _dafny.MultiSet
                def iife102_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(120, 247):
                        d_1310_v10_: int = compr_38_
                        if ((120) <= (d_1310_v10_)) and ((d_1310_v10_) < (247)):
                            coll38_[(d_1310_v10_) + (23)] = (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))]
                    return _dafny.Map(coll38_)
                def iife103_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in (d_1307_v13_).Elements:
                        d_1311_v12_: int = compr_39_
                        if (d_1311_v12_) in (d_1307_v13_):
                            coll39_[default__.safeDivisionInt(d_1311_v12_, 163)] = (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))]
                    return _dafny.Map(coll39_)
                d_1309_v15_ = _dafny.MultiSet([(iife102_()
                ) | (d_1306_v11_), (iife103_()
                 if (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))] else d_1306_v11_), (((d_1308_v14_)[p1] if (p1) in (d_1308_v14_) else d_1306_v11_)) | (d_1306_v11_)])
                d_1312_v16_: _dafny.Seq
                d_1312_v16_ = _dafny.SeqWithoutIsStrInference([p0])
                index201_ = default__.safeIndex(94, (d_1304_v9_).length(0))
                rhs220_ = p0
                rhs221_ = (_dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(len(_dafny.Map({d_1302_v7_: (self).f18})), len(p2))] for d_1313_i1_ in range(default__.abs(475))])) + ((d_1312_v16_)[default__.safeIndex(len(self.f27), len(d_1312_v16_))])
                rhs222_ = d_1309_v15_
                lhs185_ = d_1304_v9_
                lhs186_ = default__.safeIndex(94, (d_1304_v9_).length(0))
                lhs187_ = globalState
                lhs185_[lhs186_] = rhs220_
                lhs187_.f12 = rhs221_
                d_1309_v15_ = rhs222_
                (globalState).f0 = (D2_DC7(len(d_1303_v8_), (self).f18, p1)).cf15
                (globalState).f15 = ((p1) * (p1)) + (p1)
                d_1314_v17_: _dafny.Array
                nw218_ = _dafny.Array(None, 9)
                nw218_[int(0)] = p1
                nw218_[int(1)] = p1
                nw218_[int(2)] = (p1 if True else p1)
                nw218_[int(3)] = p1
                nw218_[int(4)] = p1
                nw218_[int(5)] = (p1) * (len(_dafny.Map({(d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))]: d_1302_v7_})))
                nw218_[int(6)] = p1
                nw218_[int(7)] = default__.fm1(p1, (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))], (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))], globalState)
                nw218_[int(8)] = len((p2) + (self.f27))
                d_1314_v17_ = nw218_
                index202_ = default__.safeIndex(608, (d_1314_v17_).length(0))
                (d_1314_v17_)[index202_] = (p1) - (p1)
                index203_ = default__.safeIndex(608, (d_1314_v17_).length(0))
                (d_1314_v17_)[index203_] = default__.safeModuloInt(default__.safeDivisionInt(p1, p1), -138)
                rhs223_ = p1
                rhs224_ = (_dafny.CodePoint('s') if not((self).f18) else d_1302_v7_)
                lhs188_ = globalState
                lhs188_.f0 = rhs223_
                d_1302_v7_ = rhs224_
            elif True:
                index204_ = default__.safeIndex(52, (d_1301_v6_).length(0))
                (d_1301_v6_)[index204_] = (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))]
                d_1315_v18_: C1
                nw219_ = C1()
                nw219_.ctor__((d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))])
                d_1315_v18_ = nw219_
                d_1316_v19_: _dafny.Map
                d_1316_v19_ = _dafny.Map({d_1299_v4_: (p1) * (p1)})
                d_1317_v20_: _dafny.Array
                def lambda56_(d_1318_i2_):
                    return _dafny.CodePoint('n')

                init31_ = lambda56_
                nw220_ = _dafny.Array(None, 17)
                for i0_31_ in range(nw220_.length(0)):
                    nw220_[i0_31_] = init31_(i0_31_)
                d_1317_v20_ = nw220_
                d_1319_v22_: _dafny.MultiSet
                d_1319_v22_ = _dafny.MultiSet([True])
                d_1320_v23_: _dafny.Set
                def iife104_():
                    coll40_ = _dafny.Map()
                    compr_40_: _dafny.MultiSet
                    for compr_40_ in (_dafny.MultiSet([d_1319_v22_])).Elements:
                        d_1321_v21_: _dafny.MultiSet = compr_40_
                        if (d_1321_v21_) in (_dafny.MultiSet([d_1319_v22_])):
                            coll40_[d_1321_v21_] = (d_1301_v6_)[default__.safeIndex(52, (d_1301_v6_).length(0))]
                    return _dafny.Map(coll40_)
                d_1320_v23_ = _dafny.Set({(D3_DC12(d_1317_v20_, (self).f18, (0) - (p1), iife104_()
, (self).f18)).cf27, -198, 35})
                (globalState).f0 = ((d_1316_v19_)[d_1299_v4_] if (d_1299_v4_) in (d_1316_v19_) else default__.safeDivisionInt(len((self).f28), len(d_1320_v23_)))
                def iife105_():
                    coll41_ = _dafny.Set()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(290, 598):
                        d_1322_v24_: int = compr_41_
                        if ((290) <= (d_1322_v24_)) and ((d_1322_v24_) < (598)):
                            coll41_ = coll41_.union(_dafny.Set([(d_1322_v24_) * (p1)]))
                    return _dafny.Set(coll41_)
                d_1320_v23_ = (iife105_()
                ) | (d_1320_v23_)
                d_1323_v25_: _dafny.Seq
                d_1323_v25_ = _dafny.SeqWithoutIsStrInference([p1])
                (globalState).f2 = ((p1) * (p1)) <= (((d_1298_v3_)[p1] if (p1) in (d_1298_v3_) else len(d_1323_v25_)))
        elif True:
            d_1324_v26_: D9
            d_1324_v26_ = D9_DC23((d_1299_v4_)[default__.safeIndex(p1, len(d_1299_v4_))], (self).f18, p1)
            source12_ = D0_DC1((d_1299_v4_)[default__.safeIndex(p1, len(d_1299_v4_))], p1, p0, (d_1324_v26_).cf44, len(default__.fm43((self).f18, False, globalState)))
            if source12_.is_DC1:
                d_1325___mcc_h0_ = source12_.cf1
                d_1326___mcc_h1_ = source12_.cf2
                d_1327___mcc_h2_ = source12_.cf3
                d_1328___mcc_h3_ = source12_.cf4
                d_1329___mcc_h4_ = source12_.cf5
                d_1330_cf5_ = d_1329___mcc_h4_
                d_1331_cf4_ = d_1328___mcc_h3_
                d_1332_cf3_ = d_1327___mcc_h2_
                d_1333_cf2_ = d_1326___mcc_h1_
                d_1334_cf1_ = d_1325___mcc_h0_
                d_1335_v27_: _dafny.Map
                d_1335_v27_ = _dafny.Map({p2: (len(p2)) - ((0) - (d_1333_cf2_))})
                d_1335_v27_ = (d_1335_v27_).set(d_1332_cf3_, p1)
                d_1336_v28_: _dafny.Set
                d_1336_v28_ = _dafny.Set({d_1331_cf4_, (self).f18})
                rhs225_ = d_1333_cf2_
                rhs226_ = (self).f18
                rhs227_ = (((d_1296_v1_)[True] if (True) in (d_1296_v1_) else d_1333_cf2_)) - (len(self.f27))
                rhs228_ = (d_1336_v28_).ispropersubset(d_1336_v28_)
                rhs229_ = p1
                lhs189_ = globalState
                lhs190_ = globalState
                lhs191_ = globalState
                lhs189_.f15 = rhs225_
                d_1334_cf1_ = rhs226_
                d_1333_cf2_ = rhs227_
                lhs190_.f2 = rhs228_
                lhs191_.f3 = rhs229_
                d_1337_v29_: _dafny.Map
                d_1337_v29_ = _dafny.Map({d_1331_cf4_: d_1334_cf1_})
                d_1338_v30_: _dafny.Seq
                d_1338_v30_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({(self).f18: d_1334_cf1_})).set((self).f18, d_1334_cf1_), d_1337_v29_])
                d_1339_v31_: _dafny.Map
                d_1339_v31_ = _dafny.Map({p1: d_1338_v30_})
                d_1339_v31_ = (d_1339_v31_).set((0) - (((d_1298_v3_)[d_1333_cf2_] if (d_1333_cf2_) in (d_1298_v3_) else d_1330_cf5_)), (d_1338_v30_) + (d_1338_v30_))
                d_1340_v32_: _dafny.Seq
                d_1340_v32_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1341_i3_ in range(default__.abs(499))])), d_1333_cf2_])
                d_1342_v33_: _dafny.Set
                d_1342_v33_ = _dafny.Set({(d_1340_v32_)[default__.safeIndex(734, len(d_1340_v32_))]})
                d_1343_v34_: _dafny.Array
                def lambda57_(d_1344_i4_):
                    return not((D11_DC28((self).f18)).cf51)

                init32_ = lambda57_
                nw221_ = _dafny.Array(None, 21)
                for i0_32_ in range(nw221_.length(0)):
                    nw221_[i0_32_] = init32_(i0_32_)
                d_1343_v34_ = nw221_
                d_1345_v35_: _dafny.Map
                d_1345_v35_ = _dafny.Map({d_1343_v34_: False})
                d_1346_v36_: _dafny.Array
                nw222_ = _dafny.Array(None, 26)
                nw222_[int(0)] = d_1333_cf2_
                nw222_[int(1)] = p1
                nw222_[int(2)] = d_1330_cf5_
                nw222_[int(3)] = default__.safeModuloInt(d_1333_cf2_, len(d_1342_v33_))
                nw222_[int(4)] = (d_1333_cf2_) - (p1)
                nw222_[int(5)] = default__.safeModuloInt(d_1333_cf2_, len(_dafny.Map({d_1334_cf1_: False})))
                nw222_[int(6)] = (default__.fm1(len(p0), not(d_1334_cf1_), (self).f18, globalState)) + (p1)
                nw222_[int(7)] = d_1333_cf2_
                nw222_[int(8)] = d_1333_cf2_
                nw222_[int(9)] = p1
                nw222_[int(10)] = d_1333_cf2_
                nw222_[int(11)] = ((d_1340_v32_)[default__.safeIndex(default__.fm1(p1, True, (self).f18, globalState), len(d_1340_v32_))]) + (p1)
                nw222_[int(12)] = len(d_1340_v32_)
                nw222_[int(13)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kseym")))) - (d_1330_cf5_)
                nw222_[int(14)] = default__.safeDivisionInt(d_1333_cf2_, d_1330_cf5_)
                nw222_[int(15)] = d_1333_cf2_
                nw222_[int(16)] = len(d_1337_v29_)
                nw222_[int(17)] = ((d_1298_v3_)[69] if (69) in (d_1298_v3_) else p1)
                nw222_[int(18)] = len((d_1340_v32_) + (d_1340_v32_))
                nw222_[int(19)] = d_1330_cf5_
                nw222_[int(20)] = d_1333_cf2_
                nw222_[int(21)] = 680
                nw222_[int(22)] = (p1) + (p1)
                nw222_[int(23)] = 264
                nw222_[int(24)] = len(d_1345_v35_)
                nw222_[int(25)] = d_1333_cf2_
                d_1346_v36_ = nw222_
                index205_ = default__.safeIndex(607, (d_1346_v36_).length(0))
                (d_1346_v36_)[index205_] = d_1330_cf5_
                index206_ = default__.safeIndex(607, (d_1346_v36_).length(0))
                (d_1346_v36_)[index206_] = d_1333_cf2_
            elif True:
                d_1347___mcc_h5_ = source12_.cf0
                d_1348_cf0_ = d_1347___mcc_h5_
                (globalState).f3 = default__.safeDivisionInt(p1, p1)
                d_1349_v37_: _dafny.MultiSet
                d_1349_v37_ = _dafny.MultiSet([d_1348_cf0_])
                d_1350_v38_: _dafny.Seq
                d_1350_v38_ = _dafny.SeqWithoutIsStrInference([(0) - (p1), (d_1349_v37_).cardinality])
                d_1351_v39_: _dafny.Seq
                d_1351_v39_ = _dafny.SeqWithoutIsStrInference([d_1298_v3_, d_1298_v3_])
                d_1352_v40_: C3
                nw223_ = C3()
                nw223_.ctor__((d_1350_v38_) + (d_1350_v38_), self.f27, d_1351_v39_, (d_1348_cf0_) == (d_1348_cf0_))
                d_1352_v40_ = nw223_
                d_1353_v41_: C2
                nw224_ = C2()
                nw224_.ctor__(998, (d_1352_v40_).f35, ((self).f18 if (self).f18 else (d_1299_v4_)[default__.safeIndex(p1, len(d_1299_v4_))]))
                d_1353_v41_ = nw224_
                d_1354_v42_: _dafny.Set
                d_1354_v42_ = _dafny.Set({d_1348_cf0_})
                d_1355_v43_: str
                d_1355_v43_ = _dafny.CodePoint('k')
                pat_let_tv22_ = d_1348_cf0_
                d_1356_v44_: _dafny.MultiSet
                def iife106_(_pat_let32_0):
                    def iife107_(d_1357_dt__update__tmp_h0_):
                        def iife108_(_pat_let33_0):
                            def iife109_(d_1358_dt__update_hcf18_h0_):
                                def iife110_(_pat_let34_0):
                                    def iife111_(d_1359_dt__update_hcf20_h0_):
                                        return D2_DC8(d_1358_dt__update_hcf18_h0_, (d_1357_dt__update__tmp_h0_).cf19, d_1359_dt__update_hcf20_h0_)
                                    return iife111_(_pat_let34_0)
                                return iife110_(pat_let_tv22_)
                            return iife109_(_pat_let33_0)
                        return iife108_((self).f18)
                    return iife107_(_pat_let32_0)
                d_1356_v44_ = _dafny.MultiSet([iife106_(D2_DC8(d_1348_cf0_, d_1348_cf0_, not(d_1348_cf0_)))])
                d_1360_v45_: D0
                d_1360_v45_ = D0_DC1(d_1348_cf0_, len(d_1354_v42_), _dafny.SeqWithoutIsStrInference([d_1355_v43_]), (self).f18, (d_1356_v44_).cardinality)
                d_1361_v46_: _dafny.Set
                d_1361_v46_ = _dafny.Set({d_1353_v41_.f33})
                d_1362_v47_: _dafny.MultiSet
                d_1362_v47_ = _dafny.MultiSet([364, d_1353_v41_.f33, -990])
                d_1363_v48_: _dafny.Array
                nw225_ = _dafny.Array(None, 18)
                nw225_[int(0)] = d_1348_cf0_
                nw225_[int(1)] = d_1348_cf0_
                nw225_[int(2)] = d_1348_cf0_
                nw225_[int(3)] = (d_1348_cf0_) == (not((d_1360_v45_).cf4))
                nw225_[int(4)] = not(d_1348_cf0_)
                nw225_[int(5)] = d_1348_cf0_
                nw225_[int(6)] = d_1348_cf0_
                nw225_[int(7)] = (p1) >= (default__.fm1(d_1353_v41_.f33, (self).f18, (self).f18, globalState))
                nw225_[int(8)] = (d_1361_v46_) == (d_1361_v46_)
                nw225_[int(9)] = (self).fm14(d_1353_v41_.f33, d_1299_v4_, d_1362_v47_, globalState)
                nw225_[int(10)] = d_1348_cf0_
                nw225_[int(11)] = not (not(d_1348_cf0_)) or ((d_1352_v40_).fm2(d_1348_cf0_, False, (self).f18, globalState))
                nw225_[int(12)] = False
                nw225_[int(13)] = (self).f18
                nw225_[int(14)] = False
                nw225_[int(15)] = (d_1361_v46_).isdisjoint(d_1361_v46_)
                nw225_[int(16)] = (self).f18
                nw225_[int(17)] = d_1348_cf0_
                d_1363_v48_ = nw225_
                d_1363_v48_ = d_1363_v48_
            d_1364_v50_: _dafny.MultiSet
            d_1364_v50_ = _dafny.MultiSet([p1])
            d_1365_v52_: _dafny.MultiSet
            d_1365_v52_ = _dafny.MultiSet([(self).f18])
            d_1366_v54_: _dafny.Map
            def iife112_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(359, 764):
                    d_1367_v53_: int = compr_42_
                    if ((359) <= (d_1367_v53_)) and ((d_1367_v53_) < (764)):
                        coll42_[(d_1367_v53_) * (len(d_1296_v1_))] = (0) - (p1)
                return _dafny.Map(coll42_)
            d_1366_v54_ = _dafny.Map({p1: iife112_()
            })
            d_1368_v55_: _dafny.Seq
            d_1368_v55_ = _dafny.SeqWithoutIsStrInference([p1, (0) - (p1), p1, p1])
            d_1369_v56_: str
            d_1369_v56_ = _dafny.CodePoint('w')
            d_1370_v57_: _dafny.Array
            nw226_ = _dafny.Array(None, 25)
            nw226_[int(0)] = d_1298_v3_
            nw226_[int(1)] = d_1298_v3_
            nw226_[int(2)] = d_1298_v3_
            nw226_[int(3)] = d_1298_v3_
            def iife113_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(-938, 151):
                    d_1371_v49_: int = compr_43_
                    if ((-938) <= (d_1371_v49_)) and ((d_1371_v49_) < (151)):
                        coll43_[default__.safeDivisionInt(d_1371_v49_, p1)] = p1
                return _dafny.Map(coll43_)
            nw226_[int(4)] = iife113_()
            
            nw226_[int(5)] = (d_1298_v3_).set((d_1364_v50_).cardinality, p1)
            nw226_[int(6)] = (d_1298_v3_).set(p1, default__.fm1(p1, not((self).f18), (self).f18, globalState))
            nw226_[int(7)] = _dafny.Map({(0) - (p1): p1})
            nw226_[int(8)] = d_1298_v3_
            nw226_[int(9)] = d_1298_v3_
            def iife114_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in _dafny.IntegerRange(504, -597):
                    d_1372_v51_: int = compr_44_
                    if ((504) <= (d_1372_v51_)) and ((d_1372_v51_) < (-597)):
                        coll44_[default__.safeDivisionInt(d_1372_v51_, p1)] = p1
                return _dafny.Map(coll44_)
            nw226_[int(10)] = iife114_()
            
            nw226_[int(11)] = (d_1298_v3_).set(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxqm"))))
            nw226_[int(12)] = _dafny.Map({p1: ((d_1365_v52_)[False] if (False) in (d_1365_v52_) else 699)})
            nw226_[int(13)] = d_1298_v3_
            nw226_[int(14)] = d_1298_v3_
            nw226_[int(15)] = d_1298_v3_
            nw226_[int(16)] = ((d_1366_v54_)[len(default__.fm16(len(d_1368_v55_), (self).f18, d_1369_v56_, 858, globalState))] if (len(default__.fm16(len(d_1368_v55_), (self).f18, d_1369_v56_, 858, globalState))) in (d_1366_v54_) else d_1298_v3_)
            nw226_[int(17)] = d_1298_v3_
            nw226_[int(18)] = _dafny.Map({p1: (0) - (p1)})
            nw226_[int(19)] = (d_1298_v3_) | (d_1298_v3_)
            nw226_[int(20)] = d_1298_v3_
            nw226_[int(21)] = d_1298_v3_
            nw226_[int(22)] = d_1298_v3_
            nw226_[int(23)] = d_1298_v3_
            nw226_[int(24)] = d_1298_v3_
            d_1370_v57_ = nw226_
            index207_ = default__.safeIndex(393, (d_1370_v57_).length(0))
            (d_1370_v57_)[index207_] = d_1298_v3_
            index208_ = default__.safeIndex(393, (d_1370_v57_).length(0))
            (d_1370_v57_)[index208_] = d_1298_v3_
            if (self).f18:
                (globalState).f15 = default__.fm1(p1, (d_1295_v0_).cf0, False, globalState)
                d_1365_v52_ = (d_1365_v52_).intersection(d_1365_v52_)
                (globalState).f2 = (self).f18
                d_1373_v58_: _dafny.Map
                d_1373_v58_ = _dafny.Map({p1: not(True)})
                d_1374_v59_: _dafny.Set
                d_1374_v59_ = _dafny.Set({-892, p1})
                d_1375_v60_: _dafny.Seq
                d_1375_v60_ = _dafny.SeqWithoutIsStrInference([d_1374_v59_])
                d_1376_v61_: _dafny.Map
                d_1376_v61_ = _dafny.Map({True: d_1364_v50_})
                d_1373_v58_ = (d_1373_v58_).set((p1) + (len((d_1375_v60_)[default__.safeIndex(len(d_1376_v61_), len(d_1375_v60_))])), False)
                d_1377_v62_: _dafny.Seq
                d_1377_v62_ = _dafny.SeqWithoutIsStrInference([default__.fm43((self).f18, not((self).f18), globalState), d_1296_v1_, d_1296_v1_, d_1296_v1_])
                d_1378_v63_: C0
                nw227_ = C0()
                nw227_.ctor__(p1, p1)
                d_1378_v63_ = nw227_
                d_1379_v64_: _dafny.Map
                d_1379_v64_ = _dafny.Map({d_1378_v63_: (self).f18})
                d_1380_v65_: _dafny.Array
                nw228_ = _dafny.Array(None, 11)
                nw228_[int(0)] = d_1296_v1_
                nw228_[int(1)] = d_1296_v1_
                nw228_[int(2)] = d_1296_v1_
                nw228_[int(3)] = d_1296_v1_
                nw228_[int(4)] = d_1296_v1_
                nw228_[int(5)] = (d_1296_v1_).set(False, p1)
                nw228_[int(6)] = (d_1377_v62_)[default__.safeIndex(len(((d_1366_v54_)[p1] if (p1) in (d_1366_v54_) else _dafny.Map({p1: p1}))), len(d_1377_v62_))]
                nw228_[int(7)] = default__.fm29(True, p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsbldutix")), globalState)
                nw228_[int(8)] = (d_1296_v1_) | (d_1296_v1_)
                nw228_[int(9)] = _dafny.Map({(self).f18: len(d_1379_v64_)})
                nw228_[int(10)] = (d_1296_v1_).set(not((self).f18), 988)
                d_1380_v65_ = nw228_
                d_1380_v65_ = d_1380_v65_
            elif True:
                (globalState).f15 = p1
                (globalState).f2 = (self).f18
                d_1381_v66_: _dafny.MultiSet
                d_1381_v66_ = _dafny.MultiSet([d_1368_v55_])
                d_1382_v67_: _dafny.Array
                nw229_ = _dafny.Array(None, 19)
                nw229_[int(0)] = (488) + (-771)
                nw229_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmvtc")))
                nw229_[int(2)] = (0) - (-186)
                nw229_[int(3)] = -800
                nw229_[int(4)] = p1
                nw229_[int(5)] = p1
                nw229_[int(6)] = p1
                nw229_[int(7)] = p1
                nw229_[int(8)] = default__.safeDivisionInt(p1, ((d_1381_v66_)[d_1368_v55_] if (d_1368_v55_) in (d_1381_v66_) else p1))
                nw229_[int(9)] = -40
                nw229_[int(10)] = -801
                nw229_[int(11)] = (0) - (len(d_1368_v55_))
                nw229_[int(12)] = 577
                nw229_[int(13)] = (0) - (default__.safeModuloInt(((d_1296_v1_)[(self).f18] if ((self).f18) in (d_1296_v1_) else p1), (0) - (len(self.f27))))
                nw229_[int(14)] = (0) - ((0) - (p1))
                nw229_[int(15)] = (0) - (p1)
                nw229_[int(16)] = p1
                nw229_[int(17)] = p1
                nw229_[int(18)] = 426
                d_1382_v67_ = nw229_
                d_1383_v68_: _dafny.Array
                nw230_ = _dafny.Array(None, 9)
                nw230_[int(0)] = d_1369_v56_
                nw230_[int(1)] = (p0)[default__.safeIndex(p1, len(p0))]
                nw230_[int(2)] = d_1369_v56_
                nw230_[int(3)] = d_1369_v56_
                nw230_[int(4)] = d_1369_v56_
                nw230_[int(5)] = d_1369_v56_
                nw230_[int(6)] = d_1369_v56_
                nw230_[int(7)] = d_1369_v56_
                nw230_[int(8)] = default__.fm46(globalState)
                d_1383_v68_ = nw230_
                d_1384_v69_: _dafny.Seq
                d_1384_v69_ = _dafny.SeqWithoutIsStrInference([d_1365_v52_])
                d_1385_v70_: _dafny.Map
                d_1385_v70_ = _dafny.Map({(d_1384_v69_)[default__.safeIndex(100, len(d_1384_v69_))]: True})
                d_1386_v71_: D3
                d_1386_v71_ = D3_DC12(d_1383_v68_, not(not((self).f18)), p1, d_1385_v70_, (self).f18)
                d_1387_v72_: _dafny.Map
                d_1387_v72_ = _dafny.Map({p1: (self).f18})
                d_1388_v73_: _dafny.Map
                d_1388_v73_ = _dafny.Map({D2_DC6(d_1296_v1_): d_1369_v56_})
                d_1389_v74_: T1
                nw231_ = C4()
                nw231_.ctor__(d_1369_v56_, p1, (self).f18, p2, _dafny.SeqWithoutIsStrInference([(d_1370_v57_)[default__.safeIndex(393, (d_1370_v57_).length(0))] for d_1390_i5_ in range(default__.abs(574))]))
                d_1389_v74_ = nw231_
                d_1391_v75_: _dafny.Map
                d_1391_v75_ = _dafny.Map({((d_1296_v1_)[True] if (True) in (d_1296_v1_) else len((_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])).set(default__.safeIndex((0) - (p1), len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))), (self).f18))): d_1389_v74_})
                d_1392_v76_: _dafny.Seq
                d_1392_v76_ = _dafny.SeqWithoutIsStrInference([d_1391_v75_])
                d_1393_v77_: D13
                d_1393_v77_ = D13_DC33((d_1389_v74_).f18, p1, (self).f18, (0) - (len(d_1368_v55_)))
                nw232_ = _dafny.Array(None, 28)
                nw232_[int(0)] = p1
                nw232_[int(1)] = default__.safeModuloInt(p1, 836)
                nw232_[int(2)] = p1
                nw232_[int(3)] = p1
                nw232_[int(4)] = (p1 if (self).f18 else p1)
                nw232_[int(5)] = (len(d_1296_v1_)) - (p1)
                nw232_[int(6)] = default__.safeDivisionInt(486, p1)
                nw232_[int(7)] = (p1 if (self).f18 else p1)
                nw232_[int(8)] = (d_1386_v71_).cf27
                nw232_[int(9)] = p1
                nw232_[int(10)] = len(d_1387_v72_)
                nw232_[int(11)] = len(d_1388_v73_)
                nw232_[int(12)] = (-963) + (p1)
                nw232_[int(13)] = p1
                nw232_[int(14)] = (d_1364_v50_).cardinality
                nw232_[int(15)] = (p1) - (p1)
                nw232_[int(16)] = 942
                nw232_[int(17)] = (0) - ((379) + (len(d_1392_v76_)))
                nw232_[int(18)] = ((d_1364_v50_)[p1] if (p1) in (d_1364_v50_) else len(d_1368_v55_))
                nw232_[int(19)] = (0) - (p1)
                nw232_[int(20)] = ((0) - (p1)) - ((0) - (p1))
                nw232_[int(21)] = (d_1393_v77_).cf56
                nw232_[int(22)] = default__.safeModuloInt(p1, p1)
                nw232_[int(23)] = p1
                nw232_[int(24)] = p1
                nw232_[int(25)] = p1
                nw232_[int(26)] = default__.safeDivisionInt(p1, p1)
                nw232_[int(27)] = (p1) + (len(d_1368_v55_))
                d_1382_v67_ = nw232_
                d_1394_v78_: int
                d_1395_v79_: int
                d_1396_v80_: _dafny.Array
                d_1397_v81_: _dafny.Map
                out36_: int
                out37_: int
                out38_: _dafny.Array
                out39_: _dafny.Map
                out36_, out37_, out38_, out39_ = default__.m0(globalState)
                d_1394_v78_ = out36_
                d_1395_v79_ = out37_
                d_1396_v80_ = out38_
                d_1397_v81_ = out39_
                d_1398_v82_: _dafny.Array
                nw233_ = _dafny.Array(False, 15)
                d_1398_v82_ = nw233_
                d_1398_v82_ = d_1398_v82_
            d_1399_v83_: _dafny.MultiSet
            d_1399_v83_ = d_1365_v52_
            d_1399_v83_ = d_1399_v83_
            d_1400_v84_: _dafny.Array
            nw234_ = _dafny.Array(int(0), 3)
            d_1400_v84_ = nw234_
            index209_ = default__.safeIndex(173, (d_1400_v84_).length(0))
            (d_1400_v84_)[index209_] = len(p0)
            index210_ = default__.safeIndex(173, (d_1400_v84_).length(0))
            rhs230_ = _dafny.SeqWithoutIsStrInference([d_1369_v56_ for d_1401_i6_ in range(default__.abs(556))])
            rhs231_ = p1
            lhs192_ = globalState
            lhs193_ = d_1400_v84_
            lhs194_ = default__.safeIndex(173, (d_1400_v84_).length(0))
            lhs192_.f12 = rhs230_
            lhs193_[lhs194_] = rhs231_
        d_1402_v85_: _dafny.Map
        d_1402_v85_ = _dafny.Map({(self).f18: default__.fm48(globalState)})
        d_1402_v85_ = (d_1402_v85_).set((self).f18, self.f27)
        (globalState).f0 = default__.fm1(p1, (self).f18, (self).f18, globalState)
        (globalState).f3 = p1

    def m7(self, p0, p1, globalState):
        d_1403_v0_: _dafny.Set
        d_1403_v0_ = _dafny.Set({p1, p1, p1, p1, p1})
        d_1404_v1_: _dafny.MultiSet
        d_1404_v1_ = _dafny.MultiSet([default__.fm23(d_1403_v0_, p1, globalState), (self).f18])
        d_1405_v2_: _dafny.Seq
        d_1405_v2_ = _dafny.SeqWithoutIsStrInference([(d_1404_v1_).isdisjoint(d_1404_v1_)])
        (globalState).f2 = (d_1405_v2_)[default__.safeIndex(p1, len(d_1405_v2_))]
        if ((self).f18) or (p0):
            d_1406_v3_: _dafny.Array
            nw235_ = _dafny.Array(False, 5)
            d_1406_v3_ = nw235_
            index211_ = default__.safeIndex(291, (d_1406_v3_).length(0))
            (d_1406_v3_)[index211_] = p0
            index212_ = default__.safeIndex(291, (d_1406_v3_).length(0))
            (d_1406_v3_)[index212_] = (p0) and (True)
            d_1407_v4_: _dafny.Array
            def lambda58_(d_1408_v2_):
                def lambda59_(d_1409_i0_):
                    return _dafny.Map({False: d_1408_v2_})

                return lambda59_

            init33_ = lambda58_(d_1405_v2_)
            nw236_ = _dafny.Array(None, 24)
            for i0_33_ in range(nw236_.length(0)):
                nw236_[i0_33_] = init33_(i0_33_)
            d_1407_v4_ = nw236_
            d_1410_v5_: _dafny.Map
            d_1410_v5_ = _dafny.Map({p0: d_1405_v2_})
            index213_ = default__.safeIndex(588, (d_1407_v4_).length(0))
            (d_1407_v4_)[index213_] = d_1410_v5_
            index214_ = default__.safeIndex(588, (d_1407_v4_).length(0))
            (d_1407_v4_)[index214_] = d_1410_v5_
            index215_ = default__.safeIndex(291, (d_1406_v3_).length(0))
            index216_ = default__.safeIndex(291, (d_1406_v3_).length(0))
            rhs232_ = default__.safeModuloInt(p1, p1)
            rhs233_ = (self).f18
            rhs234_ = (p0) not in ((d_1405_v2_) + (d_1405_v2_))
            rhs235_ = (self.f27) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnynibgy")))
            rhs236_ = d_1406_v3_
            lhs195_ = globalState
            lhs196_ = d_1406_v3_
            lhs197_ = default__.safeIndex(291, (d_1406_v3_).length(0))
            lhs198_ = globalState
            lhs199_ = d_1406_v3_
            lhs200_ = default__.safeIndex(291, (d_1406_v3_).length(0))
            lhs195_.f3 = rhs232_
            lhs196_[lhs197_] = rhs233_
            lhs198_.f2 = rhs234_
            lhs199_[lhs200_] = rhs235_
            d_1406_v3_ = rhs236_
            if False:
                d_1411_v6_: _dafny.Map
                d_1411_v6_ = _dafny.Map({d_1406_v3_: ((d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))]) in (d_1405_v2_)})
                d_1411_v6_ = (d_1411_v6_).set(d_1406_v3_, p0)
                d_1412_v7_: _dafny.Set
                d_1412_v7_ = _dafny.Set({(self).f18, (self).f18, (self).f18})
                d_1413_v8_: _dafny.Map
                d_1413_v8_ = _dafny.Map({(d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))]: d_1412_v7_})
                d_1413_v8_ = (d_1413_v8_).set((d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))], (_dafny.Set({(d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))], not(not((self).f18))})) - (d_1412_v7_))
                index217_ = default__.safeIndex(291, (d_1406_v3_).length(0))
                (d_1406_v3_)[index217_] = (d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))]
                index218_ = default__.safeIndex(291, (d_1406_v3_).length(0))
                (d_1406_v3_)[index218_] = (d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))]
                (globalState).f0 = (p1 if (self).f18 else len((d_1403_v0_).intersection(d_1403_v0_)))
            elif True:
                d_1414_v9_: _dafny.Seq
                d_1414_v9_ = _dafny.SeqWithoutIsStrInference([(p1) - (p1)])
                (globalState).f0 = (d_1414_v9_)[default__.safeIndex(p1, len(d_1414_v9_))]
                (globalState).f15 = p1
                d_1415_v10_: _dafny.Array
                def lambda60_(d_1416_v9_):
                    def lambda61_(d_1417_i1_):
                        return d_1416_v9_

                    return lambda61_

                init34_ = lambda60_(d_1414_v9_)
                nw237_ = _dafny.Array(None, 16)
                for i0_34_ in range(nw237_.length(0)):
                    nw237_[i0_34_] = init34_(i0_34_)
                d_1415_v10_ = nw237_
                index219_ = default__.safeIndex(389, (d_1415_v10_).length(0))
                (d_1415_v10_)[index219_] = d_1414_v9_
                index220_ = default__.safeIndex(389, (d_1415_v10_).length(0))
                (d_1415_v10_)[index220_] = (d_1414_v9_) + ((_dafny.SeqWithoutIsStrInference([p1])) + (d_1414_v9_))
                d_1418_v11_: _dafny.Seq
                d_1418_v11_ = _dafny.SeqWithoutIsStrInference([d_1403_v0_])
                (globalState).f2 = default__.fm23((d_1418_v11_)[default__.safeIndex(p1, len(d_1418_v11_))], (0) - (p1), globalState)
                d_1419_v12_: str
                d_1419_v12_ = _dafny.CodePoint('w')
                d_1420_v13_: _dafny.MultiSet
                d_1420_v13_ = _dafny.MultiSet([(self).f28, (self.f27).set(default__.safeIndex(p1, len(self.f27)), d_1419_v12_), (self.f27).set(default__.safeIndex(98, len(self.f27)), d_1419_v12_)])
                d_1421_v14_: _dafny.Array
                nw238_ = _dafny.Array(None, 20)
                nw238_[int(0)] = d_1420_v13_
                nw238_[int(1)] = d_1420_v13_
                nw238_[int(2)] = d_1420_v13_
                nw238_[int(3)] = _dafny.MultiSet([self.f27, self.f27, self.f27])
                nw238_[int(4)] = (d_1420_v13_).intersection(_dafny.MultiSet([self.f27]))
                nw238_[int(5)] = d_1420_v13_
                nw238_[int(6)] = d_1420_v13_
                nw238_[int(7)] = d_1420_v13_
                nw238_[int(8)] = d_1420_v13_
                nw238_[int(9)] = default__.fm52(p1, p1, globalState)
                nw238_[int(10)] = d_1420_v13_
                nw238_[int(11)] = default__.fm52(p1, p1, globalState)
                nw238_[int(12)] = d_1420_v13_
                nw238_[int(13)] = (d_1420_v13_).intersection(d_1420_v13_)
                nw238_[int(14)] = _dafny.MultiSet([self.f27, self.f27, (self).f28])
                nw238_[int(15)] = (d_1420_v13_) | (d_1420_v13_)
                nw238_[int(16)] = (d_1420_v13_) - (d_1420_v13_)
                nw238_[int(17)] = d_1420_v13_
                nw238_[int(18)] = d_1420_v13_
                nw238_[int(19)] = default__.fm52((d_1404_v1_).cardinality, 798, globalState)
                d_1421_v14_ = nw238_
                index221_ = default__.safeIndex(919, (d_1421_v14_).length(0))
                (d_1421_v14_)[index221_] = d_1420_v13_
                d_1422_v15_: D12
                d_1422_v15_ = D12_DC31()
                index222_ = default__.safeIndex(919, (d_1421_v14_).length(0))
                rhs237_ = not((self).f18)
                rhs238_ = (_dafny.MultiSet([(self).f28, (self).f28, (self).f28])) - (d_1420_v13_)
                rhs239_ = (d_1419_v12_) in (self.f27)
                rhs240_ = d_1422_v15_
                lhs201_ = globalState
                lhs202_ = d_1421_v14_
                lhs203_ = default__.safeIndex(919, (d_1421_v14_).length(0))
                lhs204_ = globalState
                lhs201_.f2 = rhs237_
                lhs202_[lhs203_] = rhs238_
                lhs204_.f2 = rhs239_
                d_1422_v15_ = rhs240_
            d_1405_v2_ = _dafny.SeqWithoutIsStrInference([(d_1406_v3_)[default__.safeIndex(291, (d_1406_v3_).length(0))]])
        elif True:
            d_1423_v16_: _dafny.Array
            nw239_ = _dafny.Array(_dafny.Seq({}), 27)
            d_1423_v16_ = nw239_
            index223_ = default__.safeIndex(131, (d_1423_v16_).length(0))
            (d_1423_v16_)[index223_] = d_1405_v2_
            d_1424_v17_: _dafny.Seq
            d_1424_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f18]), (d_1405_v2_) + (_dafny.SeqWithoutIsStrInference([p0])), (d_1405_v2_).set(default__.safeIndex(p1, len(d_1405_v2_)), p0), (d_1405_v2_) + (d_1405_v2_)])
            d_1425_v18_: _dafny.Seq
            d_1425_v18_ = _dafny.SeqWithoutIsStrInference([p1])
            index224_ = default__.safeIndex(131, (d_1423_v16_).length(0))
            rhs241_ = p0
            rhs242_ = ((self).f18) and ((self).f18)
            rhs243_ = (d_1424_v17_)[default__.safeIndex(default__.fm1(len(d_1425_v18_), p0, p0, globalState), len(d_1424_v17_))]
            rhs244_ = ((d_1403_v0_) - (d_1403_v0_)) - (d_1403_v0_)
            lhs205_ = globalState
            lhs206_ = globalState
            lhs207_ = d_1423_v16_
            lhs208_ = default__.safeIndex(131, (d_1423_v16_).length(0))
            lhs205_.f2 = rhs241_
            lhs206_.f2 = rhs242_
            lhs207_[lhs208_] = rhs243_
            d_1403_v0_ = rhs244_
            d_1426_v19_: _dafny.Array
            nw240_ = _dafny.Array(False, 24)
            d_1426_v19_ = nw240_
            index225_ = default__.safeIndex(294, (d_1426_v19_).length(0))
            (d_1426_v19_)[index225_] = p0
            d_1427_v20_: D0
            d_1427_v20_ = D0_DC1(p0, -422, self.f27, p0, p1)
            d_1428_v21_: _dafny.Seq
            def iife115_(_pat_let35_0):
                def iife116_(d_1429_dt__update__tmp_h0_):
                    def iife117_(_pat_let36_0):
                        def iife118_(d_1430_dt__update_hcf4_h0_):
                            def iife119_(_pat_let37_0):
                                def iife120_(d_1431_dt__update_hcf1_h0_):
                                    return D0_DC1(d_1431_dt__update_hcf1_h0_, (d_1429_dt__update__tmp_h0_).cf2, (d_1429_dt__update__tmp_h0_).cf3, d_1430_dt__update_hcf4_h0_, (d_1429_dt__update__tmp_h0_).cf5)
                                return iife120_(_pat_let37_0)
                            return iife119_((self).f18)
                        return iife118_(_pat_let36_0)
                    return iife117_(False)
                return iife116_(_pat_let35_0)
            d_1428_v21_ = _dafny.SeqWithoutIsStrInference([default__.fm35(p1, (self).f18, globalState), iife115_(d_1427_v20_)])
            d_1432_v22_: _dafny.Seq
            d_1432_v22_ = _dafny.SeqWithoutIsStrInference([d_1428_v21_])
            d_1433_v23_: _dafny.Array
            nw241_ = _dafny.Array(None, 7)
            nw241_[int(0)] = (d_1428_v21_).set(default__.safeIndex(p1, len(d_1428_v21_)), d_1427_v20_)
            nw241_[int(1)] = d_1428_v21_
            nw241_[int(2)] = default__.fm53(globalState)
            nw241_[int(3)] = d_1428_v21_
            nw241_[int(4)] = (d_1432_v22_)[default__.safeIndex(p1, len(d_1432_v22_))]
            nw241_[int(5)] = (d_1428_v21_) + (d_1428_v21_)
            nw241_[int(6)] = (d_1428_v21_).set(default__.safeIndex(p1, len(d_1428_v21_)), d_1427_v20_)
            d_1433_v23_ = nw241_
            index226_ = default__.safeIndex(789, (d_1433_v23_).length(0))
            (d_1433_v23_)[index226_] = default__.fm53(globalState)
            d_1434_v24_: str
            d_1434_v24_ = _dafny.CodePoint('j')
            d_1435_v25_: _dafny.Seq
            d_1435_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbmoton"))])
            d_1436_v26_: _dafny.Map
            d_1436_v26_ = _dafny.Map({True: p1})
            d_1437_v27_: _dafny.MultiSet
            d_1437_v27_ = _dafny.MultiSet([d_1434_v24_])
            d_1438_v28_: _dafny.Map
            d_1438_v28_ = _dafny.Map({(self).f18: d_1437_v27_})
            d_1439_v29_: _dafny.Map
            d_1439_v29_ = _dafny.Map({p1: (((d_1438_v28_)[p0] if (p0) in (d_1438_v28_) else d_1437_v27_)).cardinality})
            d_1440_v30_: _dafny.Seq
            d_1440_v30_ = _dafny.SeqWithoutIsStrInference([d_1439_v29_])
            d_1441_v31_: C4
            nw242_ = C4()
            nw242_.ctor__(d_1434_v24_, (_dafny.MultiSet(d_1425_v18_)).cardinality, (self).f18, ((d_1435_v25_)[default__.safeIndex(((d_1436_v26_)[(self).f18] if ((self).f18) in (d_1436_v26_) else p1), len(d_1435_v25_))]).set(default__.safeIndex(len(_dafny.Map({len(d_1439_v29_): p0})), len((d_1435_v25_)[default__.safeIndex(((d_1436_v26_)[(self).f18] if ((self).f18) in (d_1436_v26_) else p1), len(d_1435_v25_))])), d_1434_v24_), d_1440_v30_)
            d_1441_v31_ = nw242_
            d_1442_v32_: _dafny.Map
            d_1442_v32_ = _dafny.Map({d_1441_v31_: p0})
            pat_let_tv23_ = p0
            pat_let_tv24_ = d_1441_v31_
            index227_ = default__.safeIndex(294, (d_1426_v19_).length(0))
            index228_ = default__.safeIndex(789, (d_1433_v23_).length(0))
            def iife121_(_pat_let38_0):
                def iife122_(d_1443_dt__update__tmp_h1_):
                    def iife123_(_pat_let39_0):
                        def iife124_(d_1444_dt__update_hcf4_h1_):
                            def iife125_(_pat_let40_0):
                                def iife126_(d_1445_dt__update_hcf2_h0_):
                                    return D0_DC1((d_1443_dt__update__tmp_h1_).cf1, d_1445_dt__update_hcf2_h0_, (d_1443_dt__update__tmp_h1_).cf3, d_1444_dt__update_hcf4_h1_, (d_1443_dt__update__tmp_h1_).cf5)
                                return iife126_(_pat_let40_0)
                            return iife125_((0) - ((pat_let_tv24_).f30))
                        return iife124_(_pat_let39_0)
                    return iife123_(pat_let_tv23_)
                return iife122_(_pat_let38_0)
            rhs245_ = (len(d_1442_v32_)) > (len(default__.fm16((d_1441_v31_).f30, (self).f18, (d_1441_v31_).f29, (d_1441_v31_).f30, globalState)))
            rhs246_ = (_dafny.SeqWithoutIsStrInference([iife121_(D0_DC1((self).f18, 484, (self).f28, p0, p1)) for d_1446_i2_ in range(default__.abs(975))])) + (d_1428_v21_)
            lhs209_ = d_1426_v19_
            lhs210_ = default__.safeIndex(294, (d_1426_v19_).length(0))
            lhs211_ = d_1433_v23_
            lhs212_ = default__.safeIndex(789, (d_1433_v23_).length(0))
            lhs209_[lhs210_] = rhs245_
            lhs211_[lhs212_] = rhs246_
            d_1447_v33_: C0
            nw243_ = C0()
            nw243_.ctor__(len((d_1423_v16_)[default__.safeIndex(131, (d_1423_v16_).length(0))]), (d_1441_v31_).f30)
            d_1447_v33_ = nw243_
            d_1448_v34_: _dafny.Map
            d_1448_v34_ = _dafny.Map({d_1426_v19_: p1})
            (globalState).f15 = default__.safeModuloInt((d_1447_v33_).f32, len(d_1448_v34_))
            (globalState).f3 = default__.safeModuloInt((d_1447_v33_).f32, p1)
        d_1449_v35_: _dafny.Array
        nw244_ = _dafny.Array(int(0), 25)
        d_1449_v35_ = nw244_
        d_1450_v36_: _dafny.MultiSet
        d_1450_v36_ = _dafny.MultiSet([d_1449_v35_])
        d_1451_v37_: _dafny.Seq
        d_1451_v37_ = _dafny.SeqWithoutIsStrInference([d_1449_v35_])
        hi7_ = default__.safeModuloInt(p1, p1)
        for d_1452_i3_ in range(((d_1450_v36_)[(d_1451_v37_)[default__.safeIndex((0) - (p1), len(d_1451_v37_))]] if ((d_1451_v37_)[default__.safeIndex((0) - (p1), len(d_1451_v37_))]) in (d_1450_v36_) else p1), hi7_):
            d_1453_v38_: _dafny.Array
            nw245_ = _dafny.Array(False, 23)
            d_1453_v38_ = nw245_
            index229_ = default__.safeIndex(177, (d_1453_v38_).length(0))
            (d_1453_v38_)[index229_] = p0
            index230_ = default__.safeIndex(177, (d_1453_v38_).length(0))
            rhs247_ = ((self).f28) + (((self).f28) + ((self).f28))
            rhs248_ = not ((d_1404_v1_) == (d_1404_v1_)) or ((self.f27) < ((self).f28))
            lhs213_ = globalState
            lhs214_ = d_1453_v38_
            lhs215_ = default__.safeIndex(177, (d_1453_v38_).length(0))
            lhs213_.f12 = rhs247_
            lhs214_[lhs215_] = rhs248_
            d_1454_v39_: _dafny.Array
            nw246_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_1454_v39_ = nw246_
            d_1455_v40_: _dafny.Map
            d_1455_v40_ = _dafny.Map({d_1454_v39_: default__.fm48(globalState)})
            d_1455_v40_ = (d_1455_v40_).set(d_1454_v39_, (self).f28)
            d_1456_v41_: _dafny.Seq
            d_1456_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "inthwlmb")), (self).f28, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgf"))) + (self.f27)])
            d_1456_v41_ = _dafny.SeqWithoutIsStrInference([(self).f28 for d_1457_i4_ in range(default__.abs(769))])
            d_1458_v42_: str
            d_1458_v42_ = _dafny.CodePoint('w')
            d_1458_v42_ = default__.fm46(globalState)
        d_1459_v43_: _dafny.MultiSet
        d_1459_v43_ = d_1404_v1_
        d_1460_i5_: int
        d_1460_i5_ = 0
        with _dafny.label("11"):
            while (((d_1459_v43_)) - (d_1404_v1_)).isdisjoint(d_1404_v1_):
                with _dafny.c_label("11"):
                    if (d_1460_i5_) >= (100):
                        raise _dafny.Break("11")
                    d_1460_i5_ = (d_1460_i5_) + (1)
                    (globalState).f3 = 188
                    if (self).f18:
                        (globalState).f0 = (0) - ((p1) * (p1))
                        d_1461_v44_: _dafny.Array
                        nw247_ = _dafny.Array(_dafny.MultiSet({}), 2)
                        d_1461_v44_ = nw247_
                        d_1461_v44_ = d_1461_v44_
                        (self).f27 = (self).f28
                        (globalState).f3 = p1
                        d_1462_v45_: int
                        d_1463_v46_: int
                        d_1464_v47_: _dafny.Array
                        d_1465_v48_: _dafny.Map
                        out40_: int
                        out41_: int
                        out42_: _dafny.Array
                        out43_: _dafny.Map
                        out40_, out41_, out42_, out43_ = default__.m0(globalState)
                        d_1462_v45_ = out40_
                        d_1463_v46_ = out41_
                        d_1464_v47_ = out42_
                        d_1465_v48_ = out43_
                    elif True:
                        d_1466_v49_: C0
                        nw248_ = C0()
                        nw248_.ctor__(p1, default__.safeDivisionInt(p1, p1))
                        d_1466_v49_ = nw248_
                        (globalState).f3 = (d_1466_v49_).f31
                        (globalState).f0 = p1
                        d_1467_v50_: _dafny.Array
                        nw249_ = _dafny.Array(D2.default()(), 3)
                        d_1467_v50_ = nw249_
                        d_1467_v50_ = (d_1467_v50_ if (self).f18 else d_1467_v50_)
                        (globalState).f3 = 511
                    d_1468_v51_: _dafny.Array
                    nw250_ = _dafny.Array(_dafny.MultiSet({}), 18)
                    d_1468_v51_ = nw250_
                    index231_ = default__.safeIndex(248, (d_1468_v51_).length(0))
                    (d_1468_v51_)[index231_] = d_1404_v1_
                    d_1469_v52_: _dafny.Array
                    nw251_ = _dafny.Array(_dafny.Array(None, 0), 26)
                    d_1469_v52_ = nw251_
                    d_1470_v53_: _dafny.Array
                    nw252_ = _dafny.Array(False, 20)
                    d_1470_v53_ = nw252_
                    d_1471_v54_: _dafny.Map
                    d_1471_v54_ = _dafny.Map({not(False): d_1470_v53_})
                    d_1472_v55_: _dafny.Map
                    d_1472_v55_ = _dafny.Map({(self).f18: (self).f18})
                    index232_ = default__.safeIndex(248, (d_1468_v51_).length(0))
                    rhs249_ = len(default__.fm36(not(not ((self).f18) or ((self).f18)), p1, ((0) - (p1)) < (447), globalState))
                    rhs250_ = (d_1404_v1_) - (_dafny.MultiSet([((d_1472_v55_)[p0] if (p0) in (d_1472_v55_) else (self).f18)]))
                    rhs251_ = d_1469_v52_
                    rhs252_ = d_1405_v2_
                    rhs253_ = (d_1471_v54_).set((self).f18, d_1470_v53_)
                    lhs216_ = globalState
                    lhs217_ = d_1468_v51_
                    lhs218_ = default__.safeIndex(248, (d_1468_v51_).length(0))
                    lhs216_.f3 = rhs249_
                    lhs217_[lhs218_] = rhs250_
                    d_1469_v52_ = rhs251_
                    d_1405_v2_ = rhs252_
                    d_1471_v54_ = rhs253_
                    d_1473_v56_: _dafny.Seq
                    d_1473_v56_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_1474_v57_: _dafny.Map
                    d_1474_v57_ = _dafny.Map({len(d_1473_v56_): d_1403_v0_})
                    (globalState).f0 = len((d_1474_v57_) | (d_1474_v57_))
                    pass
            pass
        hi8_ = p1
        for d_1475_i6_ in range(p1, hi8_):
            d_1476_v58_: C1
            nw253_ = C1()
            nw253_.ctor__((self).f18)
            d_1476_v58_ = nw253_
            d_1477_v59_: _dafny.Map
            d_1477_v59_ = _dafny.Map({d_1475_i6_: p1})
            d_1478_v60_: D13
            d_1478_v60_ = D13_DC33(p0, ((d_1477_v59_)[p1] if (p1) in (d_1477_v59_) else -363), (self).f18, p1)
            source13_ = d_1478_v60_
            if source13_.is_DC33:
                d_1479___mcc_h0_ = source13_.cf55
                d_1480___mcc_h1_ = source13_.cf56
                d_1481___mcc_h2_ = source13_.cf57
                d_1482___mcc_h3_ = source13_.cf58
                d_1483_cf58_ = d_1482___mcc_h3_
                d_1484_cf57_ = d_1481___mcc_h2_
                d_1485_cf56_ = d_1480___mcc_h1_
                d_1486_cf55_ = d_1479___mcc_h0_
                d_1487_v61_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_1487_v61_ = nw254_
                d_1487_v61_ = d_1487_v61_
                index233_ = default__.safeIndex(781, (d_1449_v35_).length(0))
                (d_1449_v35_)[index233_] = d_1485_cf56_
                d_1488_v62_: _dafny.Array
                nw255_ = _dafny.Array(False, 1)
                d_1488_v62_ = nw255_
                d_1489_v63_: _dafny.Map
                d_1489_v63_ = _dafny.Map({611: d_1488_v62_})
                d_1490_v64_: _dafny.Seq
                d_1490_v64_ = _dafny.SeqWithoutIsStrInference([d_1485_cf56_])
                index234_ = default__.safeIndex(781, (d_1449_v35_).length(0))
                rhs254_ = len(default__.fm16(d_1483_cf58_, (self).f18, _dafny.CodePoint('o'), (0) - (d_1485_cf56_), globalState))
                rhs255_ = d_1485_cf56_
                rhs256_ = p1
                rhs257_ = (_dafny.Map({(d_1490_v64_)[default__.safeIndex(len(d_1403_v0_), len(d_1490_v64_))]: d_1488_v62_})).set(d_1485_cf56_, d_1488_v62_)
                lhs219_ = d_1449_v35_
                lhs220_ = default__.safeIndex(781, (d_1449_v35_).length(0))
                lhs219_[lhs220_] = rhs254_
                d_1483_cf58_ = rhs255_
                d_1483_cf58_ = rhs256_
                d_1489_v63_ = rhs257_
                d_1491_v65_: _dafny.Set
                d_1491_v65_ = _dafny.Set({p0, (self).f18})
                d_1492_v66_: D10
                d_1492_v66_ = D10_DC24((d_1491_v65_ if (self).f18 else d_1491_v65_))
                pat_let_tv25_ = d_1491_v65_
                pat_let_tv26_ = d_1491_v65_
                def iife128_(_pat_let42_0):
                    def iife129_(d_1493_dt__update__tmp_h2_):
                        def iife130_(_pat_let43_0):
                            def iife131_(d_1494_dt__update_hcf47_h0_):
                                return D10_DC24(d_1494_dt__update_hcf47_h0_)
                            return iife131_(_pat_let43_0)
                        return iife130_(pat_let_tv25_)
                    return iife129_(_pat_let42_0)
                def iife127_(_pat_let41_0):
                    def iife132_(d_1495_dt__update__tmp_h3_):
                        def iife133_(_pat_let44_0):
                            def iife134_(d_1496_dt__update_hcf47_h1_):
                                return D10_DC24(d_1496_dt__update_hcf47_h1_)
                            return iife134_(_pat_let44_0)
                        return iife133_(pat_let_tv26_)
                    return iife132_(_pat_let41_0)
                d_1492_v66_ = iife127_(iife128_(d_1492_v66_))
                (globalState).f12 = self.f27
            elif source13_.is_DC34:
                d_1497___mcc_h4_ = source13_.cf59
                d_1498___mcc_h5_ = source13_.cf60
                d_1499___mcc_h6_ = source13_.cf61
                d_1500___mcc_h7_ = source13_.cf62
                d_1501_cf62_ = d_1500___mcc_h7_
                d_1502_cf61_ = d_1499___mcc_h6_
                d_1503_cf60_ = d_1498___mcc_h5_
                d_1504_cf59_ = d_1497___mcc_h4_
                (globalState).f15 = default__.safeDivisionInt((d_1475_i6_) + (p1), (0) - (d_1475_i6_))
                d_1505_v67_: int
                d_1506_v68_: int
                d_1507_v69_: _dafny.Array
                d_1508_v70_: _dafny.Map
                out44_: int
                out45_: int
                out46_: _dafny.Array
                out47_: _dafny.Map
                out44_, out45_, out46_, out47_ = default__.m0(globalState)
                d_1505_v67_ = out44_
                d_1506_v68_ = out45_
                d_1507_v69_ = out46_
                d_1508_v70_ = out47_
                d_1509_v71_: D0
                d_1509_v71_ = D0_DC1((self).f18, len(_dafny.SeqWithoutIsStrInference([d_1501_cf62_, (self).f18])), self.f27, (self).f18, len(d_1477_v59_))
                index235_ = default__.safeIndex(757, (d_1507_v69_).length(0))
                (d_1507_v69_)[index235_] = (d_1505_v67_) + ((d_1509_v71_).cf5)
                index236_ = default__.safeIndex(757, (d_1507_v69_).length(0))
                (d_1507_v69_)[index236_] = d_1505_v67_
                d_1405_v2_ = ((d_1405_v2_).set(default__.safeIndex(d_1505_v67_, len(d_1405_v2_)), not(d_1501_cf62_))) + (d_1405_v2_)
            elif source13_.is_DC32:
                d_1510___mcc_h8_ = source13_.cf54
                d_1511_cf54_ = d_1510___mcc_h8_
                d_1512_v72_: C1
                nw256_ = C1()
                nw256_.ctor__(not(False))
                d_1512_v72_ = nw256_
                d_1513_v73_: _dafny.Seq
                d_1513_v73_ = _dafny.SeqWithoutIsStrInference([d_1475_i6_])
                d_1514_v75_: _dafny.Set
                d_1514_v75_ = _dafny.Set({p0, (self).f18})
                d_1515_v78_: D0
                d_1515_v78_ = D0_DC1(p0, d_1475_i6_, self.f27, False, d_1475_i6_)
                d_1516_v79_: _dafny.MultiSet
                d_1516_v79_ = _dafny.MultiSet([p1])
                d_1517_v80_: _dafny.Array
                nw257_ = _dafny.Array(None, 27)
                nw257_[int(0)] = (self).f18
                nw257_[int(1)] = p0
                nw257_[int(2)] = p0
                nw257_[int(3)] = ((self).f28) == (self.f27)
                def iife135_():
                    coll45_ = _dafny.Set()
                    compr_45_: int
                    for compr_45_ in (d_1513_v73_).Elements:
                        d_1518_v74_: int = compr_45_
                        if (d_1518_v74_) in (d_1513_v73_):
                            coll45_ = coll45_.union(_dafny.Set([default__.safeModuloInt(d_1518_v74_, 742)]))
                    return _dafny.Set(coll45_)
                nw257_[int(4)] = (iife135_()
                ).isdisjoint(d_1403_v0_)
                nw257_[int(5)] = (self).f18
                nw257_[int(6)] = (self).f18
                nw257_[int(7)] = (d_1514_v75_) == (d_1514_v75_)
                nw257_[int(8)] = p0
                def iife136_():
                    coll46_ = _dafny.Set()
                    compr_46_: int
                    for compr_46_ in (d_1403_v0_).Elements:
                        d_1519_v76_: int = compr_46_
                        if (d_1519_v76_) in (d_1403_v0_):
                            coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1519_v76_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gm"))))]))
                    return _dafny.Set(coll46_)
                nw257_[int(9)] = default__.fm23(iife136_()
                , (d_1513_v73_)[default__.safeIndex(d_1475_i6_, len(d_1513_v73_))], globalState)
                nw257_[int(10)] = p0
                nw257_[int(11)] = (self).f18
                nw257_[int(12)] = (self).f18
                nw257_[int(13)] = False
                def iife137_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(868, 822):
                        d_1520_v77_: int = compr_47_
                        if ((868) <= (d_1520_v77_)) and ((d_1520_v77_) < (822)):
                            coll47_ = coll47_.union(_dafny.Set([default__.safeModuloInt(d_1520_v77_, d_1475_i6_)]))
                    return _dafny.Set(coll47_)
                nw257_[int(14)] = default__.fm23(iife137_()
                , p1, globalState)
                nw257_[int(15)] = (self).fm15(179, d_1475_i6_, globalState)
                nw257_[int(16)] = p0
                nw257_[int(17)] = default__.fm23(d_1403_v0_, len(self.f27), globalState)
                nw257_[int(18)] = p0
                nw257_[int(19)] = ((d_1512_v72_).fm22(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bm")), d_1515_v78_, globalState)) < (len(d_1513_v73_))
                nw257_[int(20)] = p0
                nw257_[int(21)] = (self).f18
                nw257_[int(22)] = (self).f18
                nw257_[int(23)] = (p0 if (self).f18 else not((self).f18))
                nw257_[int(24)] = default__.fm23(d_1403_v0_, (0) - (p1), globalState)
                nw257_[int(25)] = (default__.fm25(p1, globalState)).isdisjoint(d_1516_v79_)
                nw257_[int(26)] = True
                d_1517_v80_ = nw257_
                index237_ = default__.safeIndex(576, (d_1517_v80_).length(0))
                (d_1517_v80_)[index237_] = p0
                d_1521_v81_: D9
                d_1521_v81_ = D9_DC23(p0, (self).f18, len(_dafny.SeqWithoutIsStrInference([p1 for d_1522_i7_ in range(default__.abs(-102))])))
                d_1523_v82_: _dafny.Map
                d_1523_v82_ = _dafny.Map({False: len(self.f27)})
                index238_ = default__.safeIndex(576, (d_1517_v80_).length(0))
                rhs258_ = (self).f18
                rhs259_ = ((d_1521_v81_ if p0 else d_1521_v81_)).cf44
                rhs260_ = d_1404_v1_
                rhs261_ = (p1) * (p1)
                rhs262_ = ((((d_1523_v82_)[p0] if (p0) in (d_1523_v82_) else p1)) >= (d_1475_i6_) if (self).f18 else (self).f18)
                lhs221_ = globalState
                lhs222_ = globalState
                lhs223_ = globalState
                lhs224_ = d_1517_v80_
                lhs225_ = default__.safeIndex(576, (d_1517_v80_).length(0))
                lhs221_.f2 = rhs258_
                lhs222_.f2 = rhs259_
                d_1404_v1_ = rhs260_
                lhs223_.f0 = rhs261_
                lhs224_[lhs225_] = rhs262_
                (globalState).f2 = p0
                (globalState).f3 = (d_1475_i6_ if (self).fm15(d_1475_i6_, d_1475_i6_, globalState) else p1)
            elif True:
                d_1524___mcc_h9_ = source13_.cf63
                d_1525_cf63_ = d_1524___mcc_h9_
                d_1526_v83_: _dafny.Array
                nw258_ = _dafny.Array(D0.default()(), 17)
                d_1526_v83_ = nw258_
                d_1527_v84_: D0
                d_1527_v84_ = D0_DC0(p0)
                index239_ = default__.safeIndex(494, (d_1526_v83_).length(0))
                (d_1526_v83_)[index239_] = (d_1527_v84_ if False else d_1527_v84_)
                d_1528_v85_: str
                d_1528_v85_ = _dafny.CodePoint('j')
                d_1529_v86_: _dafny.Seq
                d_1529_v86_ = _dafny.SeqWithoutIsStrInference([p1])
                index240_ = default__.safeIndex(494, (d_1526_v83_).length(0))
                def iife138_(_pat_let45_0):
                    def iife139_(d_1530_dt__update__tmp_h4_):
                        def iife140_(_pat_let46_0):
                            def iife141_(d_1531_dt__update_hcf0_h0_):
                                return D0_DC0(d_1531_dt__update_hcf0_h0_)
                            return iife141_(_pat_let46_0)
                        return iife140_(True)
                    return iife139_(_pat_let45_0)
                rhs263_ = not (p0) or ((self).f18)
                rhs264_ = (len(default__.fm44(p1, p0, d_1528_v85_, d_1475_i6_, globalState))) <= (len((d_1529_v86_) + (d_1529_v86_)))
                rhs265_ = iife138_(d_1527_v84_)
                rhs266_ = p0
                lhs226_ = globalState
                lhs227_ = globalState
                lhs228_ = d_1526_v83_
                lhs229_ = default__.safeIndex(494, (d_1526_v83_).length(0))
                lhs230_ = globalState
                lhs226_.f2 = rhs263_
                lhs227_.f2 = rhs264_
                lhs228_[lhs229_] = rhs265_
                lhs230_.f2 = rhs266_
                d_1532_v87_: _dafny.Map
                d_1532_v87_ = _dafny.Map({d_1528_v85_: p0})
                (globalState).f15 = (450) + (len((d_1532_v87_) | (d_1532_v87_)))
                d_1533_v88_: _dafny.Seq
                d_1533_v88_ = _dafny.SeqWithoutIsStrInference([d_1477_v59_, d_1477_v59_])
                d_1534_v89_: C3
                nw259_ = C3()
                nw259_.ctor__(d_1529_v86_, (self.f27) + (self.f27), d_1533_v88_, (p1) == (d_1475_i6_))
                d_1534_v89_ = nw259_
                d_1535_v90_: _dafny.Seq
                d_1535_v90_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1536_i8_ in range(default__.abs(926))])])
                d_1537_v91_: C4
                nw260_ = C4()
                nw260_.ctor__(d_1528_v85_, p1, False, (self.f27) + ((d_1535_v90_)[default__.safeIndex(d_1475_i6_, len(d_1535_v90_))]), _dafny.SeqWithoutIsStrInference([d_1477_v59_ for d_1538_i9_ in range(default__.abs(792))]))
                d_1537_v91_ = nw260_
            (globalState).f3 = p1
            d_1539_v92_: _dafny.Seq
            d_1539_v92_ = _dafny.SeqWithoutIsStrInference([p1, 571, ((d_1477_v59_)[p1] if (p1) in (d_1477_v59_) else d_1475_i6_), len(default__.fm54(d_1475_i6_, p1, globalState))])
            (globalState).f2 = not((_dafny.SeqWithoutIsStrInference([len(self.f27) for d_1540_i10_ in range(default__.abs(-129))])) < (d_1539_v92_))
        source14_ = d_1459_v43_
        d_1541___mcc_h10_ = source14_
        d_1542_cf65_ = d_1541___mcc_h10_
        d_1543_v93_: str
        d_1543_v93_ = _dafny.CodePoint('l')
        d_1544_v94_: _dafny.MultiSet
        d_1544_v94_ = _dafny.MultiSet([((self).f28).set(default__.safeIndex(len(self.f27), len((self).f28)), d_1543_v93_), self.f27, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mnrphd")), (self).f28, (self).f28])
        (globalState).f0 = default__.safeModuloInt(p1, ((d_1544_v94_)[(self).f28] if ((self).f28) in (d_1544_v94_) else p1))
        index241_ = default__.safeIndex(49, (d_1449_v35_).length(0))
        (d_1449_v35_)[index241_] = (p1) + (p1)
        index242_ = default__.safeIndex(49, (d_1449_v35_).length(0))
        (d_1449_v35_)[index242_] = p1
        d_1545_v95_: _dafny.Array
        nw261_ = _dafny.Array(_dafny.Seq({}), 14)
        d_1545_v95_ = nw261_
        d_1546_v96_: _dafny.Seq
        d_1546_v96_ = _dafny.SeqWithoutIsStrInference([(d_1449_v35_)[default__.safeIndex(49, (d_1449_v35_).length(0))], (_dafny.MultiSet([p0, (self).f18])).cardinality, (d_1449_v35_)[default__.safeIndex(49, (d_1449_v35_).length(0))]])
        index243_ = default__.safeIndex(95, (d_1545_v95_).length(0))
        (d_1545_v95_)[index243_] = d_1546_v96_
        index244_ = default__.safeIndex(95, (d_1545_v95_).length(0))
        rhs267_ = not(p0)
        rhs268_ = (d_1546_v96_) + (d_1546_v96_)
        rhs269_ = p0
        lhs231_ = globalState
        lhs232_ = d_1545_v95_
        lhs233_ = default__.safeIndex(95, (d_1545_v95_).length(0))
        lhs234_ = globalState
        lhs231_.f2 = rhs267_
        lhs232_[lhs233_] = rhs268_
        lhs234_.f2 = rhs269_
        d_1547_v97_: _dafny.Array
        nw262_ = _dafny.Array(D13.default()(), 8)
        d_1547_v97_ = nw262_
        d_1548_v98_: _dafny.Array
        nw263_ = _dafny.Array(False, 25)
        d_1548_v98_ = nw263_
        index245_ = default__.safeIndex(836, (d_1548_v98_).length(0))
        (d_1548_v98_)[index245_] = p0
        index246_ = default__.safeIndex(836, (d_1548_v98_).length(0))
        rhs270_ = d_1547_v97_
        rhs271_ = p0
        rhs272_ = p0
        rhs273_ = p1
        lhs235_ = d_1548_v98_
        lhs236_ = default__.safeIndex(836, (d_1548_v98_).length(0))
        lhs237_ = globalState
        lhs238_ = globalState
        d_1547_v97_ = rhs270_
        lhs235_[lhs236_] = rhs271_
        lhs237_.f2 = rhs272_
        lhs238_.f15 = rhs273_

    @property
    def f28(self):
        return self._f28

class C6(T0, T1):
    def  __init__(self):
        self._f18: bool = False
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f20: _dafny.Seq = _dafny.Seq({})
        self.f26: int = int(0)
        self._f25: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f25, f26, f18, f19, f20):
        (self)._f25 = f25
        (self).f26 = f26
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20

    def fm2(self, p0, p1, p2, globalState):
        return (self).f18

    def fm3(self, p0, globalState):
        return ((self).f18) and (((self).f18 if (self).f18 else (self).f18))

    def m1(self, p0, p1, p2, globalState):
        d_1549_v0_: _dafny.Set
        d_1549_v0_ = _dafny.Set({p1})
        (globalState).f2 = (d_1549_v0_).isdisjoint(_dafny.Set({30, -321}))
        d_1550_i0_: int
        d_1550_i0_ = 0
        with _dafny.label("12"):
            while ((431 if (self).f18 else p1)) < (p1):
                with _dafny.c_label("12"):
                    if (d_1550_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_1550_i0_ = (d_1550_i0_) + (1)
                    d_1551_v1_: _dafny.MultiSet
                    d_1551_v1_ = _dafny.MultiSet([(self).f18])
                    d_1552_v2_: D0
                    d_1552_v2_ = D0_DC1((self).f18, len(p0), ((self).f19) + (p0), ((self).fm2((self).f18, (self).f18, (D2_DC8((self).f18, (self).f18, (self).f18)).cf18, globalState)) == ((self).f18), (0) - ((d_1551_v1_).cardinality))
                    source15_ = d_1552_v2_
                    if source15_.is_DC1:
                        d_1553___mcc_h0_ = source15_.cf1
                        d_1554___mcc_h1_ = source15_.cf2
                        d_1555___mcc_h2_ = source15_.cf3
                        d_1556___mcc_h3_ = source15_.cf4
                        d_1557___mcc_h4_ = source15_.cf5
                        d_1558_cf5_ = d_1557___mcc_h4_
                        d_1559_cf4_ = d_1556___mcc_h3_
                        d_1560_cf3_ = d_1555___mcc_h2_
                        d_1561_cf2_ = d_1554___mcc_h1_
                        d_1562_cf1_ = d_1553___mcc_h0_
                        d_1563_v3_: _dafny.MultiSet
                        d_1563_v3_ = _dafny.MultiSet([d_1561_cf2_])
                        d_1564_v4_: _dafny.Seq
                        d_1564_v4_ = _dafny.SeqWithoutIsStrInference([d_1558_cf5_])
                        d_1565_v5_: str
                        d_1565_v5_ = _dafny.CodePoint('f')
                        d_1566_v6_: _dafny.Array
                        nw264_ = _dafny.Array(None, 14)
                        nw264_[int(0)] = (self.f26) * (default__.fm1(p1, d_1562_cf1_, d_1562_cf1_, globalState))
                        nw264_[int(1)] = default__.safeDivisionInt((d_1563_v3_).cardinality, default__.fm1((0) - (self.f26), (self).f18, (self).f18, globalState))
                        nw264_[int(2)] = default__.safeModuloInt((0) - (p1), (0) - ((_dafny.MultiSet([(d_1551_v1_).cardinality])).cardinality))
                        nw264_[int(3)] = d_1561_cf2_
                        nw264_[int(4)] = len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1567_i1_ in range(default__.abs(846))])) + (default__.fm13(d_1559_cf4_, p1, globalState))).set(default__.safeIndex((d_1564_v4_)[default__.safeIndex(self.f26, len(d_1564_v4_))], len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1568_i1_ in range(default__.abs(846))])) + (default__.fm13(d_1559_cf4_, p1, globalState)))), d_1565_v5_))
                        nw264_[int(5)] = p1
                        nw264_[int(6)] = ((d_1551_v1_)[d_1559_cf4_] if (d_1559_cf4_) in (d_1551_v1_) else p1)
                        nw264_[int(7)] = d_1561_cf2_
                        nw264_[int(8)] = (d_1561_cf2_ if d_1562_cf1_ else (0) - (d_1558_cf5_))
                        nw264_[int(9)] = d_1558_cf5_
                        nw264_[int(10)] = self.f26
                        nw264_[int(11)] = len((self).f25)
                        nw264_[int(12)] = d_1558_cf5_
                        nw264_[int(13)] = self.f26
                        d_1566_v6_ = nw264_
                        index247_ = default__.safeIndex(171, (d_1566_v6_).length(0))
                        (d_1566_v6_)[index247_] = (0) - ((d_1558_cf5_) + (self.f26))
                        index248_ = default__.safeIndex(171, (d_1566_v6_).length(0))
                        (d_1566_v6_)[index248_] = self.f26
                        d_1563_v3_ = d_1563_v3_
                        d_1569_v7_: T0
                        nw265_ = C2()
                        nw265_.ctor__(170, d_1564_v4_, (self).f18)
                        d_1569_v7_ = nw265_
                        d_1570_v8_: _dafny.Array
                        nw266_ = _dafny.Array(False, 29)
                        d_1570_v8_ = nw266_
                        d_1571_v9_: _dafny.Map
                        d_1571_v9_ = _dafny.Map({d_1569_v7_: d_1570_v8_})
                        d_1572_v10_: _dafny.Seq
                        d_1572_v10_ = _dafny.SeqWithoutIsStrInference([((d_1571_v9_)[d_1569_v7_] if (d_1569_v7_) in (d_1571_v9_) else d_1570_v8_), d_1570_v8_])
                        d_1573_v11_: _dafny.Map
                        d_1573_v11_ = _dafny.Map({d_1559_cf4_: d_1572_v10_})
                        d_1574_v12_: _dafny.Array
                        nw267_ = _dafny.Array(None, 20)
                        nw267_[int(0)] = (d_1572_v10_) + (d_1572_v10_)
                        nw267_[int(1)] = d_1572_v10_
                        nw267_[int(2)] = (d_1572_v10_) + (d_1572_v10_)
                        nw267_[int(3)] = (d_1572_v10_) + ((d_1572_v10_).set(default__.safeIndex((d_1566_v6_)[default__.safeIndex(171, (d_1566_v6_).length(0))], len(d_1572_v10_)), d_1570_v8_))
                        nw267_[int(4)] = d_1572_v10_
                        nw267_[int(5)] = d_1572_v10_
                        nw267_[int(6)] = d_1572_v10_
                        nw267_[int(7)] = d_1572_v10_
                        nw267_[int(8)] = d_1572_v10_
                        nw267_[int(9)] = d_1572_v10_
                        nw267_[int(10)] = d_1572_v10_
                        nw267_[int(11)] = (d_1572_v10_) + (d_1572_v10_)
                        nw267_[int(12)] = (((d_1573_v11_)[d_1562_cf1_] if (d_1562_cf1_) in (d_1573_v11_) else d_1572_v10_)) + ((d_1572_v10_).set(default__.safeIndex(self.f26, len(d_1572_v10_)), d_1570_v8_))
                        nw267_[int(13)] = d_1572_v10_
                        nw267_[int(14)] = (d_1572_v10_) + (d_1572_v10_)
                        nw267_[int(15)] = d_1572_v10_
                        nw267_[int(16)] = d_1572_v10_
                        nw267_[int(17)] = (d_1572_v10_).set(default__.safeIndex(627, len(d_1572_v10_)), d_1570_v8_)
                        nw267_[int(18)] = d_1572_v10_
                        nw267_[int(19)] = _dafny.SeqWithoutIsStrInference([d_1570_v8_])
                        d_1574_v12_ = nw267_
                        index249_ = default__.safeIndex(160, (d_1574_v12_).length(0))
                        (d_1574_v12_)[index249_] = d_1572_v10_
                        index250_ = default__.safeIndex(160, (d_1574_v12_).length(0))
                        (d_1574_v12_)[index250_] = d_1572_v10_
                        d_1575_v13_: _dafny.Array
                        nw268_ = _dafny.Array(None, 7)
                        nw268_[int(0)] = (self).f25
                        nw268_[int(1)] = default__.fm44((d_1566_v6_)[default__.safeIndex(171, (d_1566_v6_).length(0))], d_1562_cf1_, d_1565_v5_, len(_dafny.SeqWithoutIsStrInference([(self).f18, d_1559_cf4_])), globalState)
                        nw268_[int(2)] = (self).f25
                        nw268_[int(3)] = (self).f25
                        nw268_[int(4)] = (self).f25
                        nw268_[int(5)] = (self).f25
                        nw268_[int(6)] = (self).f25
                        d_1575_v13_ = nw268_
                        d_1576_v14_: _dafny.Map
                        d_1576_v14_ = _dafny.Map({d_1570_v8_: d_1575_v13_})
                        d_1576_v14_ = (d_1576_v14_).set(d_1570_v8_, d_1575_v13_)
                    elif True:
                        d_1577___mcc_h5_ = source15_.cf0
                        d_1578_cf0_ = d_1577___mcc_h5_
                        (globalState).f15 = (0) - (p1)
                        d_1579_v15_: _dafny.Array
                        nw269_ = _dafny.Array(False, 25)
                        d_1579_v15_ = nw269_
                        d_1580_v16_: T0
                        nw270_ = C1()
                        nw270_.ctor__(d_1578_cf0_)
                        d_1580_v16_ = nw270_
                        d_1581_v17_: _dafny.Map
                        d_1581_v17_ = _dafny.Map({d_1579_v15_: d_1580_v16_})
                        d_1581_v17_ = (d_1581_v17_).set(d_1579_v15_, d_1580_v16_)
                        d_1582_v18_: _dafny.Array
                        nw271_ = _dafny.Array(_dafny.Seq({}), 12)
                        d_1582_v18_ = nw271_
                        d_1583_v19_: _dafny.Set
                        d_1583_v19_ = _dafny.Set({d_1578_cf0_, d_1578_cf0_})
                        rhs274_ = p1
                        rhs275_ = default__.fm1(len(d_1583_v19_), not (d_1578_cf0_) or ((self).f18), (self).f18, globalState)
                        rhs276_ = (self).f18
                        rhs277_ = d_1582_v18_
                        lhs239_ = globalState
                        lhs240_ = globalState
                        lhs241_ = globalState
                        lhs239_.f15 = rhs274_
                        lhs240_.f0 = rhs275_
                        lhs241_.f2 = rhs276_
                        d_1582_v18_ = rhs277_
                        d_1584_v20_: bool
                        d_1585_v21_: int
                        out48_: bool
                        out49_: int
                        out48_, out49_ = (self).m6(globalState)
                        d_1584_v20_ = out48_
                        d_1585_v21_ = out49_
                    d_1586_v22_: _dafny.Set
                    d_1586_v22_ = _dafny.Set({p2, (self).f19})
                    d_1587_v23_: D11
                    d_1587_v23_ = D11_DC28((self).f18)
                    d_1588_v24_: _dafny.Seq
                    d_1588_v24_ = _dafny.SeqWithoutIsStrInference([D11_DC29(d_1587_v23_), d_1587_v23_])
                    d_1589_v25_: _dafny.Map
                    d_1589_v25_ = _dafny.Map({(d_1586_v22_ if True else d_1586_v22_): D11_DC29((d_1588_v24_)[default__.safeIndex(p1, len(d_1588_v24_))])})
                    d_1589_v25_ = (d_1589_v25_).set((d_1586_v22_).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfwjy"))})), D11_DC29(d_1587_v23_))
                    d_1590_v26_: _dafny.Seq
                    d_1590_v26_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_1591_v27_: C3
                    nw272_ = C3()
                    nw272_.ctor__(d_1590_v26_, (p2) + ((self).f19), ((self).f20) + ((self).f20), (self).f18)
                    d_1591_v27_ = nw272_
                    d_1592_v28_: _dafny.Map
                    d_1592_v28_ = _dafny.Map({(self).f18: (self).f18})
                    d_1593_v29_: _dafny.Set
                    d_1593_v29_ = _dafny.Set({(d_1592_v28_) | (d_1592_v28_)})
                    d_1593_v29_ = (d_1593_v29_) | (d_1593_v29_)
                    pass
            pass
        d_1594_v30_: C1
        nw273_ = C1()
        nw273_.ctor__(False)
        d_1594_v30_ = nw273_
        d_1595_v31_: D11
        d_1595_v31_ = D11_DC27(_dafny.Map({d_1594_v30_: (self).f18}))
        d_1596_v32_: D11
        d_1596_v32_ = D11_DC29(d_1595_v31_)
        d_1597_v33_: D11
        d_1597_v33_ = D11_DC29(d_1595_v31_)
        source16_ = d_1597_v33_
        if source16_.is_DC28:
            d_1598___mcc_h6_ = source16_.cf51
            d_1599_cf51_ = d_1598___mcc_h6_
            (globalState).f0 = 287
            d_1600_v34_: _dafny.MultiSet
            d_1600_v34_ = _dafny.MultiSet([self.f26])
            d_1601_v35_: _dafny.Map
            d_1601_v35_ = _dafny.Map({d_1599_cf51_: (d_1600_v34_).cardinality})
            d_1602_v36_: _dafny.Set
            d_1602_v36_ = _dafny.Set({d_1601_v35_})
            d_1603_v37_: _dafny.Seq
            d_1603_v37_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f26), p1, (0) - (self.f26)])
            rhs278_ = default__.safeDivisionInt(self.f26, ((d_1603_v37_)[default__.safeIndex((d_1600_v34_).cardinality, len(d_1603_v37_))]) - (p1))
            rhs279_ = p1
            rhs280_ = d_1602_v36_
            rhs281_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1604_i2_ in range(default__.abs(424))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1605_i3_ in range(default__.abs(839))]))
            lhs242_ = globalState
            lhs243_ = self
            lhs244_ = globalState
            lhs242_.f0 = rhs278_
            lhs243_.f26 = rhs279_
            d_1602_v36_ = rhs280_
            lhs244_.f12 = rhs281_
            if not ((self).f18) or (not(False)):
                d_1606_v38_: _dafny.Map
                d_1606_v38_ = _dafny.Map({len((self).f19): len(d_1603_v37_)})
                d_1607_v39_: C2
                nw274_ = C2()
                nw274_.ctor__(119, d_1603_v37_, d_1599_cf51_)
                d_1607_v39_ = nw274_
                d_1608_v40_: _dafny.Map
                d_1608_v40_ = _dafny.Map({d_1607_v39_: False})
                d_1606_v38_ = (d_1606_v38_).set((0) - (self.f26), len(d_1608_v40_))
                d_1609_v41_: _dafny.Array
                nw275_ = _dafny.Array(False, 17)
                d_1609_v41_ = nw275_
                index251_ = default__.safeIndex(957, (d_1609_v41_).length(0))
                (d_1609_v41_)[index251_] = (self).f18
                d_1610_v42_: T1
                nw276_ = C3()
                nw276_.ctor__(_dafny.SeqWithoutIsStrInference([len((d_1607_v39_).f34)]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cutdr")), ((self).f20) + (_dafny.SeqWithoutIsStrInference([(d_1606_v38_).set(p1, self.f26) for d_1611_i4_ in range(default__.abs(639))])), ((self).f25)[default__.safeIndex(327, len((self).f25))])
                d_1610_v42_ = nw276_
                d_1612_v43_: _dafny.Array
                nw277_ = _dafny.Array(_dafny.Set({}), 8)
                d_1612_v43_ = nw277_
                index252_ = default__.safeIndex(898, (d_1612_v43_).length(0))
                (d_1612_v43_)[index252_] = default__.fm49(len((d_1610_v42_).f19), d_1607_v39_.f33, p1, globalState)
                index253_ = default__.safeIndex(957, (d_1609_v41_).length(0))
                index254_ = default__.safeIndex(898, (d_1612_v43_).length(0))
                rhs282_ = (d_1610_v42_).f18
                rhs283_ = (d_1610_v42_).f18
                rhs284_ = d_1610_v42_
                rhs285_ = self.f26
                rhs286_ = d_1549_v0_
                lhs245_ = d_1609_v41_
                lhs246_ = default__.safeIndex(957, (d_1609_v41_).length(0))
                lhs247_ = globalState
                lhs248_ = d_1612_v43_
                lhs249_ = default__.safeIndex(898, (d_1612_v43_).length(0))
                d_1599_cf51_ = rhs282_
                lhs245_[lhs246_] = rhs283_
                d_1610_v42_ = rhs284_
                lhs247_.f3 = rhs285_
                lhs248_[lhs249_] = rhs286_
                (globalState).f0 = ((self.f26) + (d_1607_v39_.f33)) - (default__.fm1(len((d_1612_v43_)[default__.safeIndex(898, (d_1612_v43_).length(0))]), not(not((d_1609_v41_)[default__.safeIndex(957, (d_1609_v41_).length(0))])), (d_1610_v42_).f18, globalState))
                d_1613_v44_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_1613_v44_ = nw278_
                d_1614_v45_: str
                d_1614_v45_ = _dafny.CodePoint('p')
                index255_ = default__.safeIndex(116, (d_1613_v44_).length(0))
                (d_1613_v44_)[index255_] = ((self).f19).set(default__.safeIndex(d_1607_v39_.f33, len((self).f19)), d_1614_v45_)
                index256_ = default__.safeIndex(116, (d_1613_v44_).length(0))
                (d_1613_v44_)[index256_] = ((self).f19 if (d_1610_v42_).f18 else p0)
                d_1601_v35_ = (d_1601_v35_).set(d_1599_cf51_, (self.f26) - (self.f26))
            elif True:
                d_1615_v46_: _dafny.Array
                def lambda62_(d_1616_cf51_):
                    def lambda63_(d_1617_i5_):
                        return _dafny.Map({d_1616_cf51_: 908})

                    return lambda63_

                init35_ = lambda62_(d_1599_cf51_)
                nw279_ = _dafny.Array(None, 11)
                for i0_35_ in range(nw279_.length(0)):
                    nw279_[i0_35_] = init35_(i0_35_)
                d_1615_v46_ = nw279_
                index257_ = default__.safeIndex(653, (d_1615_v46_).length(0))
                (d_1615_v46_)[index257_] = d_1601_v35_
                d_1618_v47_: _dafny.Array
                nw280_ = _dafny.Array(None, 9)
                nw280_[int(0)] = (self).f18
                nw280_[int(1)] = d_1599_cf51_
                nw280_[int(2)] = d_1599_cf51_
                nw280_[int(3)] = (self).f18
                nw280_[int(4)] = d_1599_cf51_
                nw280_[int(5)] = not(d_1599_cf51_)
                nw280_[int(6)] = d_1599_cf51_
                nw280_[int(7)] = d_1599_cf51_
                nw280_[int(8)] = (self).f18
                d_1618_v47_ = nw280_
                d_1619_v48_: _dafny.Map
                d_1619_v48_ = _dafny.Map({False: d_1618_v47_})
                d_1620_v49_: _dafny.Map
                d_1620_v49_ = _dafny.Map({(d_1619_v48_) | (d_1619_v48_): True})
                index258_ = default__.safeIndex(653, (d_1615_v46_).length(0))
                rhs287_ = ((self).f18) or (False)
                rhs288_ = ((d_1620_v49_)[_dafny.Map({(self).f18: d_1618_v47_})] if (_dafny.Map({(self).f18: d_1618_v47_})) in (d_1620_v49_) else d_1599_cf51_)
                rhs289_ = d_1599_cf51_
                rhs290_ = (d_1601_v35_) | (d_1601_v35_)
                lhs250_ = globalState
                lhs251_ = globalState
                lhs252_ = d_1615_v46_
                lhs253_ = default__.safeIndex(653, (d_1615_v46_).length(0))
                lhs250_.f2 = rhs287_
                lhs251_.f2 = rhs288_
                d_1599_cf51_ = rhs289_
                lhs252_[lhs253_] = rhs290_
                (globalState).f15 = len(d_1601_v35_)
                (globalState).f2 = (self).f18
                d_1600_v34_ = (d_1600_v34_).intersection(d_1600_v34_)
                rhs291_ = (self).f18
                rhs292_ = not((self).f18)
                rhs293_ = (0) - ((0) - (default__.safeDivisionInt((0) - ((0) - (-768)), 453)))
                lhs254_ = globalState
                lhs255_ = globalState
                lhs256_ = globalState
                lhs254_.f2 = rhs291_
                lhs255_.f2 = rhs292_
                lhs256_.f3 = rhs293_
            if d_1599_cf51_:
                d_1621_v50_: _dafny.Array
                def lambda64_(d_1622_i6_):
                    return False

                init36_ = lambda64_
                nw281_ = _dafny.Array(None, 9)
                for i0_36_ in range(nw281_.length(0)):
                    nw281_[i0_36_] = init36_(i0_36_)
                d_1621_v50_ = nw281_
                d_1623_v51_: _dafny.Seq
                d_1623_v51_ = _dafny.SeqWithoutIsStrInference([d_1621_v50_])
                d_1624_v52_: _dafny.Map
                d_1624_v52_ = _dafny.Map({d_1599_cf51_: d_1599_cf51_})
                d_1625_v53_: _dafny.Array
                nw282_ = _dafny.Array(None, 26)
                nw282_[int(0)] = d_1621_v50_
                nw282_[int(1)] = d_1621_v50_
                nw282_[int(2)] = d_1621_v50_
                nw282_[int(3)] = d_1621_v50_
                nw282_[int(4)] = d_1621_v50_
                nw282_[int(5)] = d_1621_v50_
                nw282_[int(6)] = (d_1623_v51_)[default__.safeIndex(len((self).f25), len(d_1623_v51_))]
                nw282_[int(7)] = d_1621_v50_
                nw282_[int(8)] = (d_1623_v51_)[default__.safeIndex(len(d_1624_v52_), len(d_1623_v51_))]
                nw282_[int(9)] = d_1621_v50_
                nw282_[int(10)] = d_1621_v50_
                nw282_[int(11)] = d_1621_v50_
                nw282_[int(12)] = d_1621_v50_
                nw282_[int(13)] = d_1621_v50_
                nw282_[int(14)] = d_1621_v50_
                nw282_[int(15)] = d_1621_v50_
                nw282_[int(16)] = d_1621_v50_
                nw282_[int(17)] = d_1621_v50_
                nw282_[int(18)] = d_1621_v50_
                nw282_[int(19)] = d_1621_v50_
                nw282_[int(20)] = d_1621_v50_
                nw282_[int(21)] = d_1621_v50_
                nw282_[int(22)] = d_1621_v50_
                nw282_[int(23)] = d_1621_v50_
                nw282_[int(24)] = d_1621_v50_
                nw282_[int(25)] = d_1621_v50_
                d_1625_v53_ = nw282_
                index259_ = default__.safeIndex(29, (d_1625_v53_).length(0))
                (d_1625_v53_)[index259_] = d_1621_v50_
                index260_ = default__.safeIndex(29, (d_1625_v53_).length(0))
                (d_1625_v53_)[index260_] = d_1621_v50_
                d_1626_v54_: str
                d_1626_v54_ = _dafny.CodePoint('f')
                d_1627_v55_: C4
                nw283_ = C4()
                nw283_.ctor__(d_1626_v54_, self.f26, (self).f18, (p2) + ((self).f19), (self).f20)
                d_1627_v55_ = nw283_
                (globalState).f3 = self.f26
                arr0_ = (d_1625_v53_)[default__.safeIndex(29, (d_1625_v53_).length(0))]
                index261_ = default__.safeIndex(93, ((d_1625_v53_)[default__.safeIndex(29, (d_1625_v53_).length(0))]).length(0))
                arr0_[index261_] = True
                arr1_ = (d_1625_v53_)[default__.safeIndex(29, (d_1625_v53_).length(0))]
                index262_ = default__.safeIndex(93, ((d_1625_v53_)[default__.safeIndex(29, (d_1625_v53_).length(0))]).length(0))
                arr1_[index262_] = (self).f18
                d_1628_v56_: _dafny.Map
                d_1628_v56_ = _dafny.Map({len(d_1624_v52_): True})
                d_1628_v56_ = (d_1628_v56_).set(p1, True)
            elif True:
                (globalState).f3 = default__.safeDivisionInt(default__.safeModuloInt(p1, self.f26), p1)
                d_1629_v57_: _dafny.MultiSet
                d_1629_v57_ = _dafny.MultiSet([p0])
                d_1629_v57_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukw"))) + (p0), (p2) + ((self).f19), (self).f19, (p2) + ((self).f19), p2])
                (globalState).f0 = p1
                d_1630_v58_: _dafny.Map
                d_1630_v58_ = _dafny.Map({(self).f18: not(d_1599_cf51_)})
                (globalState).f2 = (p1) == (default__.safeDivisionInt(664, len(d_1630_v58_)))
                d_1631_v59_: str
                d_1631_v59_ = _dafny.CodePoint('m')
                d_1632_v60_: _dafny.Map
                d_1632_v60_ = _dafny.Map({(0) - ((d_1603_v37_)[default__.safeIndex(p1, len(d_1603_v37_))]): d_1631_v59_})
                (globalState).f0 = default__.fm1(len((d_1632_v60_) | (d_1632_v60_)), d_1599_cf51_, True, globalState)
        elif source16_.is_DC27:
            d_1633___mcc_h7_ = source16_.cf50
            d_1634_cf50_ = d_1633___mcc_h7_
            (globalState).f0 = p1
            d_1635_v61_: _dafny.Array
            nw284_ = _dafny.Array(_dafny.Seq({}), 17)
            d_1635_v61_ = nw284_
            index263_ = default__.safeIndex(451, (d_1635_v61_).length(0))
            (d_1635_v61_)[index263_] = (self).f25
            index264_ = default__.safeIndex(451, (d_1635_v61_).length(0))
            (d_1635_v61_)[index264_] = ((self).f25) + (((self).f25) + ((self).f25))
            if (p1) <= (len(p0)):
                (globalState).f9 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1636_i7_ in range(default__.abs(326))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))
                (globalState).f2 = (self).f18
                (globalState).f2 = False
                (globalState).f2 = ((self).f25)[default__.safeIndex(p1, len((self).f25))]
                d_1637_v62_: _dafny.Array
                def lambda65_(d_1638_p1_):
                    def lambda66_(d_1639_i8_):
                        return (d_1639_i8_) - (len(_dafny.SeqWithoutIsStrInference([(self).f18, (D0_DC1((self).f18, d_1638_p1_, (self).f19, (self).f18, -3)).cf1, (self).f18])))

                    return lambda66_

                init37_ = lambda65_(p1)
                nw285_ = _dafny.Array(None, 12)
                for i0_37_ in range(nw285_.length(0)):
                    nw285_[i0_37_] = init37_(i0_37_)
                d_1637_v62_ = nw285_
                index265_ = default__.safeIndex(135, (d_1637_v62_).length(0))
                (d_1637_v62_)[index265_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")) if not((self).f18) else (self).f19))
                index266_ = default__.safeIndex(135, (d_1637_v62_).length(0))
                (d_1637_v62_)[index266_] = 411
            elif True:
                d_1640_v63_: str
                d_1640_v63_ = _dafny.CodePoint('t')
                d_1641_v64_: _dafny.Map
                d_1641_v64_ = _dafny.Map({d_1640_v63_: (self).f18})
                d_1642_v65_: _dafny.Set
                d_1642_v65_ = _dafny.Set({d_1641_v64_, d_1641_v64_})
                d_1643_v66_: _dafny.Map
                d_1643_v66_ = _dafny.Map({self.f26: (_dafny.Set({d_1641_v64_, d_1641_v64_})).isdisjoint(d_1642_v65_)})
                d_1644_v67_: D8
                d_1644_v67_ = D8_DC20(d_1549_v0_)
                (globalState).f2 = ((d_1643_v66_)[default__.safeModuloInt(self.f26, p1)] if (default__.safeModuloInt(self.f26, p1)) in (d_1643_v66_) else default__.fm23((d_1644_v67_).cf40, self.f26, globalState))
                index267_ = default__.safeIndex(451, (d_1635_v61_).length(0))
                (d_1635_v61_)[index267_] = ((d_1635_v61_)[default__.safeIndex(451, (d_1635_v61_).length(0))]) + ((self).f25)
                d_1645_v68_: _dafny.Array
                nw286_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_1645_v68_ = nw286_
                index268_ = default__.safeIndex(792, (d_1645_v68_).length(0))
                (d_1645_v68_)[index268_] = d_1640_v63_
                index269_ = default__.safeIndex(792, (d_1645_v68_).length(0))
                (d_1645_v68_)[index269_] = d_1640_v63_
                d_1646_v69_: _dafny.Map
                d_1646_v69_ = _dafny.Map({((d_1635_v61_)[default__.safeIndex(451, (d_1635_v61_).length(0))])[default__.safeIndex(self.f26, len((d_1635_v61_)[default__.safeIndex(451, (d_1635_v61_).length(0))]))]: (self).f18})
                d_1647_v70_: _dafny.MultiSet
                d_1647_v70_ = _dafny.MultiSet([d_1646_v69_, d_1646_v69_])
                d_1648_v71_: _dafny.MultiSet
                d_1648_v71_ = _dafny.MultiSet([(self).f18])
                d_1649_v72_: _dafny.Array
                nw287_ = _dafny.Array(None, 19)
                nw287_[int(0)] = ((self).f19) < (p0)
                nw287_[int(1)] = (self).f18
                nw287_[int(2)] = not (((d_1635_v61_)[default__.safeIndex(451, (d_1635_v61_).length(0))])[default__.safeIndex(403, len((d_1635_v61_)[default__.safeIndex(451, (d_1635_v61_).length(0))]))]) or ((self).f18)
                nw287_[int(3)] = ((self).f25)[default__.safeIndex(p1, len((self).f25))]
                nw287_[int(4)] = (self).f18
                nw287_[int(5)] = default__.fm23(d_1549_v0_, p1, globalState)
                nw287_[int(6)] = (self.f26) != (self.f26)
                nw287_[int(7)] = ((d_1646_v69_)[(self).f18] if ((self).f18) in (d_1646_v69_) else (self).f18)
                nw287_[int(8)] = (self).f18
                nw287_[int(9)] = (self).f18
                nw287_[int(10)] = (self).f18
                nw287_[int(11)] = (self).f18
                nw287_[int(12)] = (p1) >= ((d_1647_v70_).cardinality)
                nw287_[int(13)] = (self).f18
                nw287_[int(14)] = True
                nw287_[int(15)] = (d_1648_v71_).ispropersubset(d_1648_v71_)
                nw287_[int(16)] = (self).f18
                nw287_[int(17)] = (self).f18
                nw287_[int(18)] = (self).f18
                d_1649_v72_ = nw287_
                index270_ = default__.safeIndex(488, (d_1649_v72_).length(0))
                (d_1649_v72_)[index270_] = ((self).f18 if (self).f18 else (self).fm2((self).f18, (self).f18, not((self).f18), globalState))
                index271_ = default__.safeIndex(488, (d_1649_v72_).length(0))
                (d_1649_v72_)[index271_] = not((self).f18)
                d_1650_v73_: _dafny.Map
                d_1650_v73_ = _dafny.Map({890: d_1643_v66_})
                (globalState).f2 = ((d_1643_v66_) | (d_1643_v66_)) != ((((d_1650_v73_)[self.f26] if (self.f26) in (d_1650_v73_) else d_1643_v66_)).set(self.f26, (d_1649_v72_)[default__.safeIndex(488, (d_1649_v72_).length(0))]))
            (globalState).f0 = self.f26
        elif True:
            d_1651___mcc_h8_ = source16_.cf52
            d_1652_cf52_ = d_1651___mcc_h8_
            (globalState).f2 = (self).f18
            (globalState).f2 = (self).f18
            if not((len(p2)) == ((0) - ((0) - (self.f26)))):
                d_1653_v74_: _dafny.Array
                nw288_ = _dafny.Array(int(0), 4)
                d_1653_v74_ = nw288_
                def lambda67_(d_1654_i9_):
                    return default__.safeDivisionInt(d_1654_i9_, self.f26)

                init38_ = lambda67_
                nw289_ = _dafny.Array(None, 17)
                for i0_38_ in range(nw289_.length(0)):
                    nw289_[i0_38_] = init38_(i0_38_)
                d_1653_v74_ = nw289_
                (globalState).f3 = p1
                d_1655_v75_: _dafny.MultiSet
                d_1655_v75_ = _dafny.MultiSet([len(d_1549_v0_)])
                d_1656_v76_: _dafny.MultiSet
                d_1656_v76_ = _dafny.MultiSet([((D16_DC38(d_1655_v75_)).cf66).cardinality])
                d_1657_v77_: _dafny.Seq
                d_1657_v77_ = _dafny.SeqWithoutIsStrInference([p1, self.f26])
                d_1658_v78_: str
                d_1658_v78_ = _dafny.CodePoint('p')
                d_1659_v79_: T1
                nw290_ = C4()
                nw290_.ctor__(d_1658_v78_, 563, (self).f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stxe")), (self).f20)
                d_1659_v79_ = nw290_
                d_1660_v80_: _dafny.Map
                d_1660_v80_ = _dafny.Map({436: d_1659_v79_})
                d_1661_v81_: _dafny.Seq
                d_1661_v81_ = _dafny.SeqWithoutIsStrInference([(self).f19])
                d_1662_v82_: _dafny.Map
                d_1662_v82_ = _dafny.Map({(d_1657_v77_)[default__.safeIndex(len(d_1660_v80_), len(d_1657_v77_))]: len(d_1661_v81_)})
                d_1663_v83_: _dafny.Set
                d_1663_v83_ = _dafny.Set({d_1657_v77_})
                d_1664_v84_: _dafny.Map
                d_1664_v84_ = _dafny.Map({len(d_1663_v83_): (d_1659_v79_).f19})
                rhs294_ = False
                rhs295_ = not((self.f26) <= (self.f26))
                rhs296_ = ((_dafny.SeqWithoutIsStrInference([p1 for d_1665_i10_ in range(default__.abs(700))])).set(default__.safeIndex((d_1655_v75_).cardinality, len(_dafny.SeqWithoutIsStrInference([p1 for d_1666_i10_ in range(default__.abs(700))]))), (d_1657_v77_)[default__.safeIndex(self.f26, len(d_1657_v77_))])).set(default__.safeIndex(((d_1662_v82_)[self.f26] if (self.f26) in (d_1662_v82_) else p1), len((_dafny.SeqWithoutIsStrInference([p1 for d_1667_i10_ in range(default__.abs(700))])).set(default__.safeIndex((d_1655_v75_).cardinality, len(_dafny.SeqWithoutIsStrInference([p1 for d_1668_i10_ in range(default__.abs(700))]))), (d_1657_v77_)[default__.safeIndex(self.f26, len(d_1657_v77_))]))), len(d_1664_v84_))
                rhs297_ = (self.f26) - (-826)
                lhs257_ = globalState
                lhs258_ = globalState
                lhs259_ = globalState
                lhs260_ = globalState
                lhs257_.f2 = rhs294_
                lhs258_.f2 = rhs295_
                lhs259_.f10 = rhs296_
                lhs260_.f0 = rhs297_
                d_1669_v85_: _dafny.Array
                nw291_ = _dafny.Array(None, 9)
                nw291_[int(0)] = d_1594_v30_
                nw291_[int(1)] = d_1594_v30_
                nw291_[int(2)] = d_1594_v30_
                nw291_[int(3)] = d_1594_v30_
                nw291_[int(4)] = d_1594_v30_
                nw291_[int(5)] = d_1594_v30_
                nw291_[int(6)] = d_1594_v30_
                nw291_[int(7)] = d_1594_v30_
                nw291_[int(8)] = d_1594_v30_
                d_1669_v85_ = nw291_
                index272_ = default__.safeIndex(611, (d_1669_v85_).length(0))
                (d_1669_v85_)[index272_] = d_1594_v30_
                index273_ = default__.safeIndex(611, (d_1669_v85_).length(0))
                (d_1669_v85_)[index273_] = d_1594_v30_
                d_1670_v86_: _dafny.Array
                def lambda68_(d_1671_v79_):
                    def lambda69_(d_1672_i11_):
                        return not ((d_1671_v79_).f18) or ((self).f18)

                    return lambda69_

                init39_ = lambda68_(d_1659_v79_)
                nw292_ = _dafny.Array(None, 28)
                for i0_39_ in range(nw292_.length(0)):
                    nw292_[i0_39_] = init39_(i0_39_)
                d_1670_v86_ = nw292_
                index274_ = default__.safeIndex(408, (d_1670_v86_).length(0))
                (d_1670_v86_)[index274_] = (d_1659_v79_).f18
                index275_ = default__.safeIndex(408, (d_1670_v86_).length(0))
                (d_1670_v86_)[index275_] = ((self).f18) not in (_dafny.SeqWithoutIsStrInference([not((self).f18)]))
            elif True:
                d_1673_v87_: bool
                d_1674_v88_: int
                out50_: bool
                out51_: int
                out50_, out51_ = (self).m6(globalState)
                d_1673_v87_ = out50_
                d_1674_v88_ = out51_
                d_1675_v89_: str
                d_1675_v89_ = _dafny.CodePoint('x')
                d_1675_v89_ = d_1675_v89_
                d_1676_v90_: _dafny.Array
                nw293_ = _dafny.Array(None, 4)
                nw293_[int(0)] = d_1675_v89_
                nw293_[int(1)] = d_1675_v89_
                nw293_[int(2)] = d_1675_v89_
                nw293_[int(3)] = default__.fm46(globalState)
                d_1676_v90_ = nw293_
                d_1676_v90_ = d_1676_v90_
                d_1677_v91_: D2
                d_1677_v91_ = D2_DC8(d_1673_v87_, False, (self).f18)
                d_1678_v92_: _dafny.Array
                nw294_ = _dafny.Array(None, 15)
                nw294_[int(0)] = d_1673_v87_
                nw294_[int(1)] = False
                nw294_[int(2)] = d_1673_v87_
                nw294_[int(3)] = False
                nw294_[int(4)] = (self).f18
                nw294_[int(5)] = False
                nw294_[int(6)] = (self).f18
                nw294_[int(7)] = (self).f18
                nw294_[int(8)] = (self).f18
                nw294_[int(9)] = d_1673_v87_
                nw294_[int(10)] = (self).f18
                nw294_[int(11)] = (self).f18
                nw294_[int(12)] = d_1673_v87_
                nw294_[int(13)] = d_1673_v87_
                nw294_[int(14)] = d_1673_v87_
                d_1678_v92_ = nw294_
                d_1679_v93_: D13
                d_1679_v93_ = D13_DC34(d_1678_v92_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")), not((self).f18), False)
                d_1680_v94_: _dafny.Map
                d_1680_v94_ = _dafny.Map({d_1677_v91_: D13_DC35(d_1679_v93_)})
                def iife142_(_pat_let47_0):
                    def iife143_(d_1681_dt__update__tmp_h0_):
                        def iife144_(_pat_let48_0):
                            def iife145_(d_1682_dt__update_hcf19_h0_):
                                return D2_DC8((d_1681_dt__update__tmp_h0_).cf18, d_1682_dt__update_hcf19_h0_, (d_1681_dt__update__tmp_h0_).cf20)
                            return iife145_(_pat_let48_0)
                        return iife144_(not(True))
                    return iife143_(_pat_let47_0)
                d_1680_v94_ = (d_1680_v94_).set(iife142_(D2_DC8((self).f18, d_1673_v87_, d_1673_v87_)), D13_DC35(d_1679_v93_))
                (globalState).f15 = d_1674_v88_
            (globalState).f0 = (0) - ((0) - (p1))
        d_1683_v95_: str
        d_1683_v95_ = _dafny.CodePoint('c')
        d_1683_v95_ = d_1683_v95_
        hi9_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsjsavsim")))
        for d_1684_i12_ in range(2, hi9_):
            (globalState).f15 = d_1684_i12_
            (globalState).f12 = p0
            d_1685_v96_: D0
            d_1685_v96_ = D0_DC1((self).f18, self.f26, (self).f19, not((self).f18), p1)
            d_1686_v97_: _dafny.Seq
            d_1686_v97_ = _dafny.SeqWithoutIsStrInference([self.f26, p1])
            d_1687_v98_: _dafny.Map
            d_1687_v98_ = _dafny.Map({(self).f18: p2})
            d_1688_v99_: _dafny.Array
            nw295_ = _dafny.Array(None, 21)
            nw295_[int(0)] = (0) - (default__.safeModuloInt(len((self).f25), self.f26))
            nw295_[int(1)] = self.f26
            nw295_[int(2)] = (d_1684_i12_) - (p1)
            nw295_[int(3)] = d_1684_i12_
            nw295_[int(4)] = default__.fm1(self.f26, (self).f18, (self).f18, globalState)
            nw295_[int(5)] = d_1684_i12_
            nw295_[int(6)] = p1
            nw295_[int(7)] = d_1684_i12_
            nw295_[int(8)] = self.f26
            nw295_[int(9)] = p1
            nw295_[int(10)] = self.f26
            nw295_[int(11)] = self.f26
            nw295_[int(12)] = self.f26
            nw295_[int(13)] = p1
            nw295_[int(14)] = 487
            nw295_[int(15)] = self.f26
            nw295_[int(16)] = (d_1594_v30_).fm22((self).f19, d_1685_v96_, globalState)
            nw295_[int(17)] = d_1684_i12_
            nw295_[int(18)] = d_1684_i12_
            nw295_[int(19)] = (len((D0_DC1((self).f18, len(d_1686_v97_), default__.fm48(globalState), (self).f18, d_1684_i12_)).cf3)) * (len(((d_1687_v98_)[(self).f18] if ((self).f18) in (d_1687_v98_) else (p0).set(default__.safeIndex(p1, len(p0)), _dafny.CodePoint('a')))))
            nw295_[int(20)] = p1
            d_1688_v99_ = nw295_
            d_1688_v99_ = d_1688_v99_
            d_1689_v100_: D0
            d_1689_v100_ = D0_DC0((self).f18)
            d_1690_v101_: _dafny.Map
            d_1690_v101_ = _dafny.Map({self.f26: d_1689_v100_})
            d_1691_v102_: _dafny.Map
            d_1691_v102_ = _dafny.Map({d_1690_v101_: (self).f18})
            d_1692_v103_: _dafny.Map
            d_1692_v103_ = _dafny.Map({(self).f18: ((d_1691_v102_)[d_1690_v101_] if (d_1690_v101_) in (d_1691_v102_) else (self).f18)})
            (globalState).f0 = len(d_1692_v103_)
        d_1693_v104_: _dafny.Set
        d_1693_v104_ = _dafny.Set({d_1683_v95_, d_1683_v95_})
        d_1694_v105_: _dafny.Map
        d_1694_v105_ = _dafny.Map({True: d_1693_v104_})
        d_1695_v106_: _dafny.Array
        nw296_ = _dafny.Array(None, 26)
        nw296_[int(0)] = (p2) + (p0)
        nw296_[int(1)] = p2
        nw296_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mamkhfj")) if (self).f18 else p2)
        nw296_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_1683_v95_ for d_1696_i13_ in range(default__.abs(387))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_1683_v95_ for d_1697_i13_ in range(default__.abs(387))]))), d_1683_v95_)
        nw296_[int(4)] = p0
        nw296_[int(5)] = (self).f19
        nw296_[int(6)] = p2
        nw296_[int(7)] = p2
        nw296_[int(8)] = (d_1594_v30_).fm21(len(p0), (self).f19, d_1694_v105_, (self).f25, globalState)
        nw296_[int(9)] = (((self).f19).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(0) - (self.f26) for d_1698_i14_ in range(default__.abs(120))])), len((self).f19)), d_1683_v95_) if (self).f18 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
        nw296_[int(10)] = (self).f19
        nw296_[int(11)] = (self).f19
        nw296_[int(12)] = (p0 if (self).f18 else default__.fm48(globalState))
        nw296_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yntu"))
        nw296_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_1683_v95_ for d_1699_i15_ in range(default__.abs(-992))])) + (p2)
        nw296_[int(15)] = (self).f19
        nw296_[int(16)] = (self).f19
        nw296_[int(17)] = ((p0).set(default__.safeIndex(len(_dafny.Set({179})), len(p0)), default__.fm46(globalState))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohly")))
        nw296_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mssdb"))
        nw296_[int(19)] = (p0) + (p2)
        nw296_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leqkgaus"))
        nw296_[int(21)] = (_dafny.SeqWithoutIsStrInference([d_1683_v95_ for d_1700_i16_ in range(default__.abs(-677))])) + ((self).f19)
        nw296_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkvl"))
        nw296_[int(23)] = p2
        nw296_[int(24)] = (p2) + (p2)
        nw296_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocvvqga"))
        d_1695_v106_ = nw296_
        index276_ = default__.safeIndex(165, (d_1695_v106_).length(0))
        (d_1695_v106_)[index276_] = p0
        index277_ = default__.safeIndex(165, (d_1695_v106_).length(0))
        (d_1695_v106_)[index277_] = ((self).f19 if not ((self).f18) or ((self).f18) else p2)

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        if False:
            if (self).f18:
                d_1701_v0_: str
                d_1701_v0_ = _dafny.CodePoint('w')
                d_1702_v1_: D3
                d_1702_v1_ = D3_DC13(len((self).f19), (self).f19, 972, self.f26, d_1701_v0_)
                d_1703_v2_: _dafny.Map
                d_1703_v2_ = _dafny.Map({p2: (_dafny.SeqWithoutIsStrInference([D3_DC13((0) - (self.f26), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpkoolb")), self.f26, 809, _dafny.CodePoint('f')) for d_1704_i0_ in range(default__.abs(238))])) + (_dafny.SeqWithoutIsStrInference([d_1702_v1_, d_1702_v1_, d_1702_v1_, d_1702_v1_]))})
                d_1705_v3_: _dafny.Map
                d_1705_v3_ = _dafny.Map({self.f26: (self).f18})
                d_1706_v4_: _dafny.Map
                d_1706_v4_ = _dafny.Map({d_1705_v3_: len((self).f19)})
                d_1707_v5_: _dafny.Seq
                d_1707_v5_ = _dafny.SeqWithoutIsStrInference([D3_DC13(self.f26, _dafny.SeqWithoutIsStrInference([d_1701_v0_ for d_1708_i1_ in range(default__.abs(984))]), self.f26, ((d_1706_v4_)[d_1705_v3_] if (d_1705_v3_) in (d_1706_v4_) else 505), d_1701_v0_), d_1702_v1_])
                d_1703_v2_ = (d_1703_v2_).set(True, (d_1707_v5_) + (d_1707_v5_))
                d_1709_v6_: _dafny.MultiSet
                d_1709_v6_ = _dafny.MultiSet([self.f26, self.f26])
                d_1710_v7_: _dafny.Set
                d_1710_v7_ = _dafny.Set({(self).f18, p1})
                d_1711_v8_: _dafny.Seq
                d_1711_v8_ = _dafny.SeqWithoutIsStrInference([-53, len(d_1710_v7_), default__.fm1(self.f26, p2, p2, globalState), self.f26])
                d_1712_v9_: _dafny.Set
                d_1712_v9_ = _dafny.Set({self.f26})
                rhs298_ = self.f26
                rhs299_ = p1
                rhs300_ = (_dafny.MultiSet(((d_1711_v8_).set(default__.safeIndex(self.f26, len(d_1711_v8_)), len(d_1712_v9_))) + (d_1711_v8_))).issubset((d_1709_v6_).intersection(d_1709_v6_))
                lhs261_ = globalState
                lhs262_ = globalState
                lhs263_ = globalState
                lhs261_.f0 = rhs298_
                lhs262_.f2 = rhs299_
                lhs263_.f2 = rhs300_
                d_1713_v10_: _dafny.Map
                d_1713_v10_ = _dafny.Map({(self.f26) + ((0) - (self.f26)): default__.fm52(self.f26, (d_1711_v8_)[default__.safeIndex(self.f26, len(d_1711_v8_))], globalState)})
                d_1714_v11_: _dafny.MultiSet
                d_1714_v11_ = _dafny.MultiSet([(self).f19, (self).f19, (self).f19, (self).f19, _dafny.SeqWithoutIsStrInference([((self).f19)[default__.safeIndex(self.f26, len((self).f19))] for d_1715_i2_ in range(default__.abs(987))])])
                d_1713_v10_ = (d_1713_v10_).set(self.f26, (d_1714_v11_) | (d_1714_v11_))
                d_1716_v12_: int
                d_1717_v13_: int
                d_1718_v14_: _dafny.Array
                d_1719_v15_: _dafny.Map
                out52_: int
                out53_: int
                out54_: _dafny.Array
                out55_: _dafny.Map
                out52_, out53_, out54_, out55_ = default__.m0(globalState)
                d_1716_v12_ = out52_
                d_1717_v13_ = out53_
                d_1718_v14_ = out54_
                d_1719_v15_ = out55_
                index278_ = default__.safeIndex(423, (d_1718_v14_).length(0))
                (d_1718_v14_)[index278_] = len((self).f19)
                d_1720_v16_: _dafny.Array
                nw297_ = _dafny.Array(None, 25)
                nw297_[int(0)] = _dafny.CodePoint('m')
                nw297_[int(1)] = d_1701_v0_
                nw297_[int(2)] = d_1701_v0_
                nw297_[int(3)] = d_1701_v0_
                nw297_[int(4)] = d_1701_v0_
                nw297_[int(5)] = d_1701_v0_
                nw297_[int(6)] = d_1701_v0_
                nw297_[int(7)] = d_1701_v0_
                nw297_[int(8)] = d_1701_v0_
                nw297_[int(9)] = d_1701_v0_
                nw297_[int(10)] = d_1701_v0_
                nw297_[int(11)] = d_1701_v0_
                nw297_[int(12)] = d_1701_v0_
                nw297_[int(13)] = d_1701_v0_
                nw297_[int(14)] = d_1701_v0_
                nw297_[int(15)] = default__.fm46(globalState)
                nw297_[int(16)] = d_1701_v0_
                nw297_[int(17)] = d_1701_v0_
                nw297_[int(18)] = default__.fm46(globalState)
                nw297_[int(19)] = d_1701_v0_
                nw297_[int(20)] = d_1701_v0_
                nw297_[int(21)] = d_1701_v0_
                nw297_[int(22)] = _dafny.CodePoint('i')
                nw297_[int(23)] = _dafny.CodePoint('g')
                nw297_[int(24)] = d_1701_v0_
                d_1720_v16_ = nw297_
                d_1721_v17_: _dafny.Set
                d_1721_v17_ = _dafny.Set({d_1720_v16_, d_1720_v16_, (d_1720_v16_ if (self).fm2(False, p1, p2, globalState) else d_1720_v16_), d_1720_v16_, d_1720_v16_})
                index279_ = default__.safeIndex(423, (d_1718_v14_).length(0))
                rhs301_ = self.f26
                rhs302_ = (0) - (default__.safeModuloInt(d_1717_v13_, d_1717_v13_))
                rhs303_ = _dafny.Set({d_1720_v16_})
                lhs264_ = d_1718_v14_
                lhs265_ = default__.safeIndex(423, (d_1718_v14_).length(0))
                lhs266_ = globalState
                lhs264_[lhs265_] = rhs301_
                lhs266_.f0 = rhs302_
                d_1721_v17_ = rhs303_
            elif True:
                d_1722_v18_: _dafny.Array
                nw298_ = _dafny.Array(False, 9)
                d_1722_v18_ = nw298_
                d_1723_v19_: D13
                d_1723_v19_ = D13_DC34(d_1722_v18_, default__.fm48(globalState), p1, p1)
                d_1724_v20_: _dafny.Map
                d_1724_v20_ = _dafny.Map({p2: d_1722_v18_})
                d_1725_v21_: _dafny.Set
                d_1725_v21_ = _dafny.Set({len(d_1724_v20_), self.f26})
                d_1726_v22_: _dafny.Map
                d_1726_v22_ = _dafny.Map({(d_1723_v19_ if True else d_1723_v19_): (self.f26) not in (d_1725_v21_)})
                d_1726_v22_ = (d_1726_v22_).set(d_1723_v19_, ((self).f18) or (not((self).f18)))
                d_1727_v23_: _dafny.Array
                nw299_ = _dafny.Array(int(0), 15)
                d_1727_v23_ = nw299_
                index280_ = default__.safeIndex(467, (d_1727_v23_).length(0))
                (d_1727_v23_)[index280_] = 813
                d_1728_v24_: _dafny.Map
                d_1728_v24_ = _dafny.Map({self.f26: True})
                index281_ = default__.safeIndex(467, (d_1727_v23_).length(0))
                rhs304_ = (self).f18
                rhs305_ = (default__.safeModuloInt((0) - (len(d_1728_v24_)), self.f26)) - (self.f26)
                rhs306_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fphpatnfw"))
                rhs307_ = (self.f26) < (len(default__.fm36(not((self).f18), self.f26, p1, globalState)))
                rhs308_ = self.f26
                lhs267_ = globalState
                lhs268_ = d_1727_v23_
                lhs269_ = default__.safeIndex(467, (d_1727_v23_).length(0))
                lhs270_ = globalState
                lhs271_ = globalState
                lhs272_ = globalState
                lhs267_.f2 = rhs304_
                lhs268_[lhs269_] = rhs305_
                lhs270_.f12 = rhs306_
                lhs271_.f2 = rhs307_
                lhs272_.f3 = rhs308_
                d_1729_v25_: str
                d_1729_v25_ = _dafny.CodePoint('f')
                d_1730_v26_: T1
                nw300_ = C4()
                nw300_.ctor__(d_1729_v25_, (0) - ((0) - (self.f26)), p1, (self).f19, (self).f20)
                d_1730_v26_ = nw300_
                d_1731_v27_: _dafny.Seq
                d_1731_v27_ = _dafny.SeqWithoutIsStrInference([d_1730_v26_])
                d_1732_v28_: T1
                nw301_ = C3()
                nw301_.ctor__(_dafny.SeqWithoutIsStrInference([self.f26, len(d_1731_v27_)]), (d_1730_v26_).f19, (self).f20, p1)
                d_1732_v28_ = nw301_
                d_1733_v29_: _dafny.Map
                d_1733_v29_ = _dafny.Map({(default__.fm1(self.f26, p1, p1, globalState)) < (len(_dafny.SeqWithoutIsStrInference([self.f26 for d_1734_i3_ in range(default__.abs(841))]))): d_1730_v26_})
                d_1735_v30_: _dafny.Map
                d_1735_v30_ = _dafny.Map({(d_1732_v28_).f18: (d_1732_v28_).f18})
                d_1733_v29_ = (d_1733_v29_).set((d_1730_v26_).fm2((D0_DC0((d_1730_v26_).f18)).cf0, (d_1732_v28_).f18, ((self).f25)[default__.safeIndex(len(_dafny.Map({len(d_1735_v30_): d_1729_v25_})), len((self).f25))], globalState), d_1730_v26_)
                d_1736_v31_: _dafny.Array
                nw302_ = _dafny.Array(None, 2)
                nw302_[int(0)] = d_1727_v23_
                nw302_[int(1)] = (d_1727_v23_ if (d_1730_v26_).f18 else d_1727_v23_)
                d_1736_v31_ = nw302_
                d_1737_v32_: _dafny.MultiSet
                d_1737_v32_ = _dafny.MultiSet([self.f26, default__.fm1(self.f26, p2, p2, globalState)])
                d_1738_v33_: D16
                d_1738_v33_ = D16_DC39(not((d_1737_v32_).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f26, (0) - ((d_1727_v23_)[default__.safeIndex(467, (d_1727_v23_).length(0))])])))), 951)
                d_1739_v34_: D2
                d_1739_v34_ = D2_DC8(p2, (self).f18, True)
                d_1740_v35_: _dafny.MultiSet
                d_1740_v35_ = _dafny.MultiSet([p1, (d_1739_v34_).cf19, p1, (d_1730_v26_).f18])
                d_1741_v36_: _dafny.Map
                d_1741_v36_ = _dafny.Map({d_1730_v26_: d_1729_v25_})
                pat_let_tv27_ = d_1727_v23_
                pat_let_tv28_ = d_1727_v23_
                def iife146_(_pat_let49_0):
                    def iife147_(d_1742_dt__update__tmp_h0_):
                        def iife148_(_pat_let50_0):
                            def iife149_(d_1743_dt__update_hcf67_h0_):
                                return D16_DC39(d_1743_dt__update_hcf67_h0_, (d_1742_dt__update__tmp_h0_).cf68)
                            return iife149_(_pat_let50_0)
                        return iife148_(((pat_let_tv28_)[default__.safeIndex(467, (pat_let_tv27_).length(0))]) > (self.f26))
                    return iife147_(_pat_let49_0)
                rhs309_ = (d_1727_v23_)[default__.safeIndex(467, (d_1727_v23_).length(0))]
                rhs310_ = d_1736_v31_
                rhs311_ = iife146_(d_1738_v33_)
                rhs312_ = (default__.fm32(globalState)) - ((_dafny.MultiSet([(d_1730_v26_).fm2(False, (d_1730_v26_).f18, (d_1732_v28_).f18, globalState)])) | ((d_1740_v35_).set(p2, default__.abs(len(d_1741_v36_)))))
                rhs313_ = (default__.fm55(self.f26, p2, len(d_1735_v30_), (d_1732_v28_).f18, globalState)).cf46
                lhs273_ = globalState
                lhs274_ = globalState
                lhs273_.f15 = rhs309_
                d_1736_v31_ = rhs310_
                d_1738_v33_ = rhs311_
                d_1740_v35_ = rhs312_
                lhs274_.f15 = rhs313_
                rhs314_ = (d_1732_v28_).f19
                rhs315_ = ((0) - ((d_1727_v23_)[default__.safeIndex(467, (d_1727_v23_).length(0))])) * (-479)
                lhs275_ = globalState
                lhs276_ = globalState
                lhs275_.f9 = rhs314_
                lhs276_.f3 = rhs315_
            d_1744_v37_: _dafny.Set
            d_1744_v37_ = _dafny.Set({self.f26})
            d_1745_v38_: _dafny.Map
            d_1745_v38_ = _dafny.Map({(self).f25: p1})
            d_1746_v39_: _dafny.Seq
            d_1746_v39_ = _dafny.SeqWithoutIsStrInference([len(d_1744_v37_), default__.safeModuloInt(self.f26, self.f26), self.f26, default__.fm1(len((self).f25), ((d_1745_v38_)[(self).f25] if ((self).f25) in (d_1745_v38_) else p1), (self).f18, globalState)])
            (globalState).f0 = (d_1746_v39_)[default__.safeIndex((self.f26) * (self.f26), len(d_1746_v39_))]
            d_1747_v40_: C1
            nw303_ = C1()
            nw303_.ctor__(p1)
            d_1747_v40_ = nw303_
            (globalState).f9 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1748_i4_ in range(default__.abs(838))])) + ((self).f19)
            d_1749_v41_: _dafny.Array
            nw304_ = _dafny.Array(None, 3)
            nw304_[int(0)] = (self).f18
            nw304_[int(1)] = p1
            nw304_[int(2)] = p1
            d_1749_v41_ = nw304_
            d_1750_v42_: _dafny.Map
            d_1750_v42_ = _dafny.Map({d_1749_v41_: True})
            (globalState).f0 = len(d_1750_v42_)
        elif True:
            d_1751_v43_: str
            d_1751_v43_ = _dafny.CodePoint('r')
            d_1752_v44_: C4
            nw305_ = C4()
            nw305_.ctor__(d_1751_v43_, self.f26, p1, ((self).f19) + ((self).f19), (self).f20)
            d_1752_v44_ = nw305_
            d_1753_v45_: _dafny.Map
            d_1753_v45_ = _dafny.Map({(self).f18: (self).f18})
            (globalState).f2 = (d_1753_v45_) == ((_dafny.Map({p1: p1})) | ((d_1753_v45_).set(p2, True)))
            if p1:
                (self).f26 = (d_1752_v44_).f30
                (globalState).f15 = default__.safeDivisionInt(self.f26, 941)
                d_1754_v46_: _dafny.Array
                nw306_ = _dafny.Array(int(0), 5)
                d_1754_v46_ = nw306_
                index282_ = default__.safeIndex(593, (d_1754_v46_).length(0))
                (d_1754_v46_)[index282_] = (len((self).f19)) - (self.f26)
                d_1755_v47_: _dafny.Map
                d_1755_v47_ = _dafny.Map({(self).f18: (d_1752_v44_).f30})
                index283_ = default__.safeIndex(593, (d_1754_v46_).length(0))
                (d_1754_v46_)[index283_] = (183) - (len((_dafny.Map({p2: len((self).f25)})) | (d_1755_v47_)))
                (globalState).f2 = p1
                (globalState).f12 = (self).f19
            elif True:
                d_1756_v48_: C1
                nw307_ = C1()
                nw307_.ctor__(False)
                d_1756_v48_ = nw307_
                d_1757_v49_: _dafny.MultiSet
                d_1757_v49_ = _dafny.MultiSet([d_1756_v48_, d_1756_v48_, d_1756_v48_])
                (globalState).f2 = ((_dafny.MultiSet([d_1756_v48_, d_1756_v48_])).intersection(d_1757_v49_)).ispropersubset(d_1757_v49_)
                d_1758_v51_: _dafny.Array
                def lambda70_(d_1759_i5_):
                    def iife150_():
                        coll48_ = _dafny.Set()
                        compr_48_: int
                        for compr_48_ in _dafny.IntegerRange(186, 24):
                            d_1760_v50_: int = compr_48_
                            if ((186) <= (d_1760_v50_)) and ((d_1760_v50_) < (24)):
                                coll48_ = coll48_.union(_dafny.Set([(d_1760_v50_) - ((0) - (self.f26))]))
                        return _dafny.Set(coll48_)
                    return iife150_()
                    

                init40_ = lambda70_
                nw308_ = _dafny.Array(None, 22)
                for i0_40_ in range(nw308_.length(0)):
                    nw308_[i0_40_] = init40_(i0_40_)
                d_1758_v51_ = nw308_
                d_1761_v52_: _dafny.Map
                d_1761_v52_ = _dafny.Map({(d_1752_v44_).f30: p2})
                d_1762_v53_: _dafny.Set
                d_1762_v53_ = _dafny.Set({len(d_1761_v52_), len(_dafny.SeqWithoutIsStrInference([(d_1752_v44_).f29 for d_1763_i6_ in range(default__.abs(-716))])), (d_1752_v44_).f30})
                index284_ = default__.safeIndex(58, (d_1758_v51_).length(0))
                (d_1758_v51_)[index284_] = d_1762_v53_
                index285_ = default__.safeIndex(58, (d_1758_v51_).length(0))
                (d_1758_v51_)[index285_] = d_1762_v53_
                d_1764_v54_: _dafny.Set
                d_1764_v54_ = _dafny.Set({p2, p2})
                d_1765_v55_: _dafny.MultiSet
                d_1765_v55_ = _dafny.MultiSet([d_1764_v54_, d_1764_v54_])
                d_1766_v56_: _dafny.Map
                d_1766_v56_ = _dafny.Map({not(p1): ((d_1765_v55_) | (d_1765_v55_)).cardinality})
                d_1766_v56_ = (d_1766_v56_).set((self).f18, (d_1752_v44_).f30)
                d_1767_v57_: C1
                nw309_ = C1()
                nw309_.ctor__((self).f18)
                d_1767_v57_ = nw309_
                d_1768_v58_: D17
                d_1768_v58_ = D17_DC41(d_1761_v52_)
                d_1769_v59_: _dafny.Seq
                d_1769_v59_ = _dafny.SeqWithoutIsStrInference([(d_1761_v52_) | ((d_1768_v58_).cf73)])
                d_1769_v59_ = (d_1769_v59_) + (d_1769_v59_)
            d_1770_v60_: _dafny.Array
            nw310_ = _dafny.Array(False, 14)
            d_1770_v60_ = nw310_
            d_1771_v61_: _dafny.Map
            d_1771_v61_ = _dafny.Map({d_1770_v60_: d_1770_v60_})
            d_1772_v62_: _dafny.Array
            def lambda71_(d_1773_i7_):
                return _dafny.CodePoint('s')

            init41_ = lambda71_
            nw311_ = _dafny.Array(None, 17)
            for i0_41_ in range(nw311_.length(0)):
                nw311_[i0_41_] = init41_(i0_41_)
            d_1772_v62_ = nw311_
            index286_ = default__.safeIndex(724, (d_1772_v62_).length(0))
            (d_1772_v62_)[index286_] = _dafny.CodePoint('h')
            d_1774_v63_: D13
            d_1774_v63_ = D13_DC33((d_1752_v44_).fm3(self.f26, globalState), default__.safeModuloInt((d_1752_v44_).f30, 788), p1, (self.f26) + ((d_1752_v44_).f30))
            index287_ = default__.safeIndex(724, (d_1772_v62_).length(0))
            rhs316_ = d_1771_v61_
            rhs317_ = (d_1752_v44_).f29
            rhs318_ = (d_1751_v43_ if (self).f18 else (d_1752_v44_).f29)
            rhs319_ = default__.fm56(p2, (self).f18, self.f26, _dafny.SeqWithoutIsStrInference([(d_1752_v44_).f29 for d_1775_i8_ in range(default__.abs(798))]), globalState)
            lhs277_ = d_1772_v62_
            lhs278_ = default__.safeIndex(724, (d_1772_v62_).length(0))
            d_1771_v61_ = rhs316_
            d_1751_v43_ = rhs317_
            lhs277_[lhs278_] = rhs318_
            d_1774_v63_ = rhs319_
            d_1776_v64_: _dafny.Array
            nw312_ = _dafny.Array(int(0), 10)
            d_1776_v64_ = nw312_
            d_1777_v65_: _dafny.MultiSet
            d_1777_v65_ = _dafny.MultiSet([d_1776_v64_])
            d_1778_v66_: C4
            nw313_ = C4()
            nw313_.ctor__(_dafny.CodePoint('m'), (d_1752_v44_).f30, p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swvwp")), _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1777_v65_).cardinality: self.f26})]))
            d_1778_v66_ = nw313_
        (self).f26 = self.f26
        if (p2) or ((not((self).f18)) or (not((self).f18))):
            d_1779_v67_: _dafny.Array
            def lambda72_(d_1780_i9_):
                return (self).f18

            init42_ = lambda72_
            nw314_ = _dafny.Array(None, 14)
            for i0_42_ in range(nw314_.length(0)):
                nw314_[i0_42_] = init42_(i0_42_)
            d_1779_v67_ = nw314_
            index288_ = default__.safeIndex(501, (d_1779_v67_).length(0))
            (d_1779_v67_)[index288_] = (self).f18
            index289_ = default__.safeIndex(501, (d_1779_v67_).length(0))
            (d_1779_v67_)[index289_] = (self).f18
            d_1781_v68_: _dafny.Seq
            d_1781_v68_ = _dafny.SeqWithoutIsStrInference([471, len(_dafny.Map({(d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]: False}))])
            d_1782_v69_: C2
            nw315_ = C2()
            nw315_.ctor__(self.f26, (d_1781_v68_ if p2 else d_1781_v68_), p1)
            d_1782_v69_ = nw315_
            if ((0) - ((0) - (d_1782_v69_.f33))) == (len((self).f19)):
                d_1783_v70_: _dafny.MultiSet
                d_1783_v70_ = _dafny.MultiSet([not((self).f18), p2, True])
                d_1784_v71_: str
                d_1784_v71_ = _dafny.CodePoint('d')
                d_1785_v72_: D9
                d_1785_v72_ = D9_DC23(p2, p2, d_1782_v69_.f33)
                d_1786_v73_: _dafny.Set
                d_1786_v73_ = _dafny.Set({((self).f25)[default__.safeIndex(d_1782_v69_.f33, len((self).f25))]})
                d_1787_v74_: _dafny.Map
                d_1787_v74_ = _dafny.Map({d_1782_v69_.f33: d_1782_v69_.f33})
                d_1788_v75_: _dafny.Map
                d_1788_v75_ = _dafny.Map({len(d_1787_v74_): self.f26})
                d_1789_v76_: _dafny.MultiSet
                d_1789_v76_ = _dafny.MultiSet([d_1787_v74_])
                d_1790_v77_: D17
                d_1790_v77_ = D17_DC43(d_1789_v76_, (self).f18)
                d_1791_v78_: _dafny.Map
                d_1791_v78_ = _dafny.Map({(self).f18: d_1782_v69_.f33})
                d_1792_v79_: D2
                d_1792_v79_ = D2_DC7(d_1782_v69_.f33, (self).f18, self.f26)
                d_1793_v80_: _dafny.Array
                nw316_ = _dafny.Array(None, 29)
                nw316_[int(0)] = self.f26
                nw316_[int(1)] = 432
                nw316_[int(2)] = d_1782_v69_.f33
                nw316_[int(3)] = ((d_1783_v70_)[(d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]] if ((d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]) in (d_1783_v70_) else self.f26)
                nw316_[int(4)] = d_1782_v69_.f33
                nw316_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_1784_v71_, d_1784_v71_]))
                nw316_[int(6)] = d_1782_v69_.f33
                nw316_[int(7)] = (default__.fm1(((d_1782_v69_).f34)[default__.safeIndex(self.f26, len((d_1782_v69_).f34))], p2, (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))], globalState)) + ((_dafny.MultiSet([d_1781_v68_, default__.fm33(True, globalState)])).cardinality)
                nw316_[int(8)] = 884
                nw316_[int(9)] = default__.safeModuloInt(self.f26, d_1782_v69_.f33)
                nw316_[int(10)] = d_1782_v69_.f33
                nw316_[int(11)] = default__.safeDivisionInt((d_1785_v72_).cf46, d_1782_v69_.f33)
                nw316_[int(12)] = d_1782_v69_.f33
                nw316_[int(13)] = (D17_DC44(d_1779_v67_, d_1782_v69_.f33)).cf78
                nw316_[int(14)] = default__.safeModuloInt(len(d_1786_v73_), d_1782_v69_.f33)
                nw316_[int(15)] = self.f26
                nw316_[int(16)] = len(_dafny.Map({True: (d_1790_v77_).cf76}))
                nw316_[int(17)] = self.f26
                nw316_[int(18)] = self.f26
                nw316_[int(19)] = (d_1782_v69_.f33) - (len(_dafny.SeqWithoutIsStrInference([d_1784_v71_ for d_1794_i10_ in range(default__.abs(323))])))
                nw316_[int(20)] = self.f26
                nw316_[int(21)] = ((d_1783_v70_)[(self).f18] if ((self).f18) in (d_1783_v70_) else default__.fm1((0) - (self.f26), (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))], p2, globalState))
                nw316_[int(22)] = d_1782_v69_.f33
                nw316_[int(23)] = ((d_1782_v69_).f34)[default__.safeIndex(d_1782_v69_.f33, len((d_1782_v69_).f34))]
                nw316_[int(24)] = (((d_1791_v78_)[(self).f18] if ((self).f18) in (d_1791_v78_) else 56) if (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))] else default__.fm1(903, (self).f18, (d_1782_v69_).fm19(d_1792_v79_, globalState), globalState))
                nw316_[int(25)] = d_1782_v69_.f33
                nw316_[int(26)] = d_1782_v69_.f33
                nw316_[int(27)] = (845) - (d_1782_v69_.f33)
                nw316_[int(28)] = self.f26
                d_1793_v80_ = nw316_
                index290_ = default__.safeIndex(145, (d_1793_v80_).length(0))
                (d_1793_v80_)[index290_] = d_1782_v69_.f33
                index291_ = default__.safeIndex(145, (d_1793_v80_).length(0))
                (d_1793_v80_)[index291_] = d_1782_v69_.f33
                d_1795_v81_: _dafny.Map
                d_1795_v81_ = _dafny.Map({236: d_1793_v80_})
                d_1795_v81_ = ((d_1795_v81_) | (d_1795_v81_)) | (d_1795_v81_)
                d_1796_v82_: _dafny.Map
                d_1796_v82_ = _dafny.Map({p2: default__.fm57(d_1782_v69_.f33, len((self).f19), True, globalState)})
                d_1796_v82_ = d_1796_v82_
                d_1797_v83_: D8
                d_1797_v83_ = D8_DC21(d_1784_v71_, len(_dafny.SeqWithoutIsStrInference([p2])))
                d_1784_v71_ = (d_1797_v83_).cf41
                d_1788_v75_ = (d_1788_v75_).set((len((self).f25)) - ((d_1793_v80_)[default__.safeIndex(145, (d_1793_v80_).length(0))]), d_1782_v69_.f33)
            elif True:
                (globalState).f0 = self.f26
                (globalState).f2 = (self).f18
                index292_ = default__.safeIndex(501, (d_1779_v67_).length(0))
                (d_1779_v67_)[index292_] = (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]
                index293_ = default__.safeIndex(501, (d_1779_v67_).length(0))
                (d_1779_v67_)[index293_] = ((self).f25)[default__.safeIndex(len(((d_1781_v68_).set(default__.safeIndex((0) - (-208), len(d_1781_v68_)), len(d_1781_v68_)) if p2 else d_1781_v68_)), len((self).f25))]
                d_1798_v84_: str
                d_1798_v84_ = _dafny.CodePoint('d')
                d_1799_v85_: _dafny.Set
                d_1799_v85_ = _dafny.Set({len(default__.fm43(p1, p2, globalState))})
                rhs320_ = default__.fm46(globalState)
                rhs321_ = (d_1782_v69_.f33 if (d_1782_v69_.f33) not in (d_1799_v85_) else d_1782_v69_.f33)
                rhs322_ = (0) - (d_1782_v69_.f33)
                lhs279_ = d_1782_v69_
                lhs280_ = d_1782_v69_
                d_1798_v84_ = rhs320_
                lhs279_.f33 = rhs321_
                lhs280_.f33 = rhs322_
            d_1800_v86_: _dafny.Set
            d_1800_v86_ = _dafny.Set({True})
            d_1801_v87_: str
            d_1801_v87_ = _dafny.CodePoint('s')
            d_1802_v88_: _dafny.Set
            d_1802_v88_ = _dafny.Set({d_1801_v87_, d_1801_v87_})
            source17_ = default__.fm24((d_1800_v86_).issubset(d_1800_v86_), d_1802_v88_, (self.f26) == (d_1782_v69_.f33), globalState)
            if source17_.is_DC3:
                d_1803___mcc_h0_ = source17_.cf7
                d_1804___mcc_h1_ = source17_.cf8
                d_1805___mcc_h2_ = source17_.cf9
                d_1806___mcc_h3_ = source17_.cf10
                d_1807___mcc_h4_ = source17_.cf11
                d_1808_cf11_ = d_1807___mcc_h4_
                d_1809_cf10_ = d_1806___mcc_h3_
                d_1810_cf9_ = d_1805___mcc_h2_
                d_1811_cf8_ = d_1804___mcc_h1_
                d_1812_cf7_ = d_1803___mcc_h0_
                d_1813_v89_: _dafny.Array
                def lambda73_(d_1814_i11_):
                    return _dafny.MultiSet([(self).f18])

                init43_ = lambda73_
                nw317_ = _dafny.Array(None, 4)
                for i0_43_ in range(nw317_.length(0)):
                    nw317_[i0_43_] = init43_(i0_43_)
                d_1813_v89_ = nw317_
                d_1815_v90_: D6
                d_1815_v90_ = D6_DC17(d_1813_v89_)
                d_1816_v91_: _dafny.Map
                d_1816_v91_ = _dafny.Map({d_1815_v90_: (self).f18})
                d_1817_v92_: _dafny.Seq
                d_1817_v92_ = _dafny.SeqWithoutIsStrInference([d_1816_v91_])
                d_1818_v93_: _dafny.Array
                nw318_ = _dafny.Array(None, 22)
                nw318_[int(0)] = d_1816_v91_
                nw318_[int(1)] = (d_1817_v92_)[default__.safeIndex(d_1810_cf9_, len(d_1817_v92_))]
                nw318_[int(2)] = d_1816_v91_
                nw318_[int(3)] = d_1816_v91_
                nw318_[int(4)] = (d_1816_v91_) | (d_1816_v91_)
                nw318_[int(5)] = (d_1816_v91_) | (d_1816_v91_)
                nw318_[int(6)] = (_dafny.Map({d_1815_v90_: p2})).set(d_1815_v90_, p1)
                nw318_[int(7)] = (d_1816_v91_) | (d_1816_v91_)
                nw318_[int(8)] = _dafny.Map({d_1815_v90_: (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]})
                nw318_[int(9)] = (d_1816_v91_).set(d_1815_v90_, (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))])
                nw318_[int(10)] = (d_1817_v92_)[default__.safeIndex(d_1811_cf8_, len(d_1817_v92_))]
                nw318_[int(11)] = (d_1816_v91_).set(d_1815_v90_, (self).f18)
                nw318_[int(12)] = (d_1816_v91_) | (d_1816_v91_)
                nw318_[int(13)] = d_1816_v91_
                nw318_[int(14)] = d_1816_v91_
                nw318_[int(15)] = (d_1816_v91_ if (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))] else d_1816_v91_)
                nw318_[int(16)] = (d_1816_v91_ if (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))] else _dafny.Map({d_1815_v90_: (self).f18}))
                nw318_[int(17)] = d_1816_v91_
                nw318_[int(18)] = d_1816_v91_
                nw318_[int(19)] = (d_1816_v91_) | (d_1816_v91_)
                nw318_[int(20)] = (d_1816_v91_) | (_dafny.Map({D6_DC17(d_1813_v89_): (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]}))
                nw318_[int(21)] = (d_1816_v91_).set(D6_DC17(d_1813_v89_), (d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))])
                d_1818_v93_ = nw318_
                d_1819_v94_: C1
                nw319_ = C1()
                nw319_.ctor__((self).f18)
                d_1819_v94_ = nw319_
                rhs323_ = d_1818_v93_
                rhs324_ = d_1801_v87_
                rhs325_ = p2
                rhs326_ = self.f26
                rhs327_ = d_1819_v94_
                lhs281_ = globalState
                lhs282_ = globalState
                d_1818_v93_ = rhs323_
                d_1801_v87_ = rhs324_
                lhs281_.f2 = rhs325_
                lhs282_.f15 = rhs326_
                d_1819_v94_ = rhs327_
                d_1820_v95_: _dafny.Map
                d_1820_v95_ = _dafny.Map({p2: d_1811_cf8_})
                rhs328_ = (d_1820_v95_) | (d_1820_v95_)
                rhs329_ = d_1810_cf9_
                lhs283_ = d_1782_v69_
                d_1820_v95_ = rhs328_
                lhs283_.f33 = rhs329_
                (globalState).f2 = p2
                (globalState).f2 = p2
            elif source17_.is_DC4:
                d_1821___mcc_h5_ = source17_.cf12
                d_1822_cf12_ = d_1821___mcc_h5_
                d_1823_v96_: _dafny.MultiSet
                d_1823_v96_ = _dafny.MultiSet([d_1782_v69_.f33, 364, -216, self.f26])
                d_1823_v96_ = default__.fm25(((0) - (d_1782_v69_.f33)) * (d_1782_v69_.f33), globalState)
                d_1824_v97_: _dafny.Array
                def lambda74_(d_1825_v69_):
                    def lambda75_(d_1826_i12_):
                        return ((_dafny.SeqWithoutIsStrInference([-914])).set(default__.safeIndex(d_1825_v69_.f33, len(_dafny.SeqWithoutIsStrInference([-914]))), d_1825_v69_.f33)) + ((d_1825_v69_).f34)

                    return lambda75_

                init44_ = lambda74_(d_1782_v69_)
                nw320_ = _dafny.Array(None, 18)
                for i0_44_ in range(nw320_.length(0)):
                    nw320_[i0_44_] = init44_(i0_44_)
                d_1824_v97_ = nw320_
                d_1827_v98_: _dafny.Map
                d_1827_v98_ = _dafny.Map({self.f26: p1})
                d_1828_v99_: _dafny.Map
                d_1828_v99_ = _dafny.Map({self.f26: ((d_1827_v98_)[564] if (564) in (d_1827_v98_) else d_1822_cf12_)})
                index294_ = default__.safeIndex(594, (d_1824_v97_).length(0))
                (d_1824_v97_)[index294_] = default__.fm33(((d_1827_v98_)[len((self).f19)] if (len((self).f19)) in (d_1827_v98_) else p2), globalState)
                index295_ = default__.safeIndex(594, (d_1824_v97_).length(0))
                (d_1824_v97_)[index295_] = (((d_1782_v69_).f34) + ((d_1782_v69_).f34) if p1 else (d_1781_v68_) + (d_1781_v68_))
                d_1828_v99_ = (d_1828_v99_).set(self.f26, (d_1782_v69_.f33) != (len((self).f19)))
                d_1829_v100_: D16
                d_1829_v100_ = D16_DC38(d_1823_v96_)
                d_1829_v100_ = d_1829_v100_
            elif source17_.is_DC2:
                d_1830___mcc_h6_ = source17_.cf6
                d_1831_cf6_ = d_1830___mcc_h6_
                (globalState).f15 = len((self).f19)
                index296_ = default__.safeIndex(501, (d_1779_v67_).length(0))
                (d_1779_v67_)[index296_] = ((0) - (d_1782_v69_.f33)) == ((0) - ((0) - (default__.fm1(d_1782_v69_.f33, p2, (self).f18, globalState))))
                index297_ = default__.safeIndex(501, (d_1779_v67_).length(0))
                (d_1779_v67_)[index297_] = not((D9_DC23((d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))], p1, -803)).cf44)
                d_1832_v101_: _dafny.Array
                nw321_ = _dafny.Array(int(0), 7)
                d_1832_v101_ = nw321_
                index298_ = default__.safeIndex(901, (d_1832_v101_).length(0))
                (d_1832_v101_)[index298_] = self.f26
                index299_ = default__.safeIndex(901, (d_1832_v101_).length(0))
                (d_1832_v101_)[index299_] = default__.safeModuloInt((self.f26 if not(p2) else self.f26), d_1782_v69_.f33)
            elif True:
                d_1833___mcc_h7_ = source17_.cf13
                d_1834_cf13_ = d_1833___mcc_h7_
                d_1835_v102_: _dafny.Map
                d_1835_v102_ = _dafny.Map({(d_1779_v67_)[default__.safeIndex(501, (d_1779_v67_).length(0))]: (d_1779_v67_ if True else d_1779_v67_)})
                d_1835_v102_ = (_dafny.Map({(self).f18: d_1779_v67_})) | (d_1835_v102_)
                d_1836_v103_: _dafny.Map
                d_1836_v103_ = _dafny.Map({d_1801_v87_: d_1782_v69_.f33})
                d_1836_v103_ = (d_1836_v103_).set(_dafny.CodePoint('s'), d_1782_v69_.f33)
                (d_1782_v69_).m1(((self).f19) + ((self).f19), self.f26, (self).f19, globalState)
                d_1837_v104_: _dafny.Array
                nw322_ = _dafny.Array(None, 10)
                d_1837_v104_ = nw322_
                index300_ = default__.safeIndex(501, (d_1779_v67_).length(0))
                (d_1779_v67_)[index300_] = (len(_dafny.Map({self.f26: d_1837_v104_}))) >= (696)
            r0 = d_1782_v69_.f33
        elif True:
            d_1838_v105_: str
            d_1838_v105_ = _dafny.CodePoint('g')
            d_1839_v106_: D8
            d_1839_v106_ = D8_DC21(d_1838_v105_, self.f26)
            d_1840_v107_: _dafny.MultiSet
            d_1840_v107_ = _dafny.MultiSet([d_1839_v106_])
            d_1841_v108_: _dafny.Seq
            d_1841_v108_ = _dafny.SeqWithoutIsStrInference([d_1840_v107_])
            d_1842_v109_: _dafny.Set
            d_1842_v109_ = _dafny.Set({-316})
            d_1843_v110_: D18
            d_1843_v110_ = D18_DC45(d_1841_v108_)
            d_1844_v111_: _dafny.Seq
            d_1844_v111_ = _dafny.SeqWithoutIsStrInference([d_1841_v108_, (d_1841_v108_).set(default__.safeIndex(len(d_1842_v109_), len(d_1841_v108_)), d_1840_v107_), _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([D8_DC21(d_1838_v105_, self.f26), d_1839_v106_]) for d_1845_i13_ in range(default__.abs(50))]), (d_1843_v110_).cf79])
            d_1841_v108_ = (d_1844_v111_)[default__.safeIndex((self.f26) - (self.f26), len(d_1844_v111_))]
            d_1846_v112_: _dafny.Array
            nw323_ = _dafny.Array(_dafny.CodePoint('D'), 13)
            d_1846_v112_ = nw323_
            index301_ = default__.safeIndex(521, (d_1846_v112_).length(0))
            (d_1846_v112_)[index301_] = d_1838_v105_
            index302_ = default__.safeIndex(521, (d_1846_v112_).length(0))
            (d_1846_v112_)[index302_] = d_1838_v105_
            if ((self).f25)[default__.safeIndex(self.f26, len((self).f25))]:
                d_1847_v113_: bool
                d_1848_v114_: int
                out56_: bool
                out57_: int
                out56_, out57_ = (self).m6(globalState)
                d_1847_v113_ = out56_
                d_1848_v114_ = out57_
                d_1849_v115_: _dafny.Array
                nw324_ = _dafny.Array(False, 27)
                d_1849_v115_ = nw324_
                index303_ = default__.safeIndex(745, (d_1849_v115_).length(0))
                (d_1849_v115_)[index303_] = (998) < (d_1848_v114_)
                index304_ = default__.safeIndex(745, (d_1849_v115_).length(0))
                (d_1849_v115_)[index304_] = p2
                d_1850_v116_: _dafny.Seq
                d_1850_v116_ = _dafny.SeqWithoutIsStrInference([690, 320, len((self).f25), self.f26])
                index305_ = default__.safeIndex(745, (d_1849_v115_).length(0))
                (d_1849_v115_)[index305_] = (p2 if (_dafny.SeqWithoutIsStrInference([d_1848_v114_, 912, self.f26, d_1848_v114_])) <= (d_1850_v116_) else False)
                d_1851_v117_: _dafny.MultiSet
                d_1851_v117_ = _dafny.MultiSet([p1])
                d_1852_v118_: _dafny.Seq
                d_1852_v118_ = _dafny.SeqWithoutIsStrInference([d_1851_v117_, d_1851_v117_])
                d_1853_v119_: _dafny.Map
                d_1853_v119_ = _dafny.Map({self.f26: (d_1851_v117_).ispropersubset((d_1852_v118_)[default__.safeIndex(d_1848_v114_, len(d_1852_v118_))])})
                index306_ = default__.safeIndex(745, (d_1849_v115_).length(0))
                rhs330_ = (d_1849_v115_)[default__.safeIndex(745, (d_1849_v115_).length(0))]
                rhs331_ = len(d_1853_v119_)
                lhs284_ = d_1849_v115_
                lhs285_ = default__.safeIndex(745, (d_1849_v115_).length(0))
                lhs284_[lhs285_] = rhs330_
                d_1848_v114_ = rhs331_
                index307_ = default__.safeIndex(745, (d_1849_v115_).length(0))
                (d_1849_v115_)[index307_] = (not((self).f18)) == (default__.fm23(d_1842_v109_, d_1848_v114_, globalState))
            elif True:
                d_1854_v120_: _dafny.MultiSet
                d_1854_v120_ = _dafny.MultiSet([p2, False])
                (globalState).f0 = ((d_1854_v120_)[(self).f18] if ((self).f18) in (d_1854_v120_) else 577)
                d_1855_v121_: _dafny.Map
                d_1855_v121_ = _dafny.Map({(self).fm2((self).f18, p2, p1, globalState): p2})
                d_1856_v122_: _dafny.Array
                nw325_ = _dafny.Array(None, 15)
                nw325_[int(0)] = (self).f18
                nw325_[int(1)] = ((self).f18 if p2 else p2)
                nw325_[int(2)] = default__.fm23(d_1842_v109_, self.f26, globalState)
                nw325_[int(3)] = (self).f18
                nw325_[int(4)] = p1
                nw325_[int(5)] = p2
                nw325_[int(6)] = p1
                nw325_[int(7)] = (self).f18
                nw325_[int(8)] = (self).f18
                nw325_[int(9)] = p2
                nw325_[int(10)] = not((self.f26) >= (64))
                nw325_[int(11)] = not(True)
                nw325_[int(12)] = ((d_1855_v121_)[p1] if (p1) in (d_1855_v121_) else not(p2))
                nw325_[int(13)] = p1
                nw325_[int(14)] = (p2) in (d_1854_v120_)
                d_1856_v122_ = nw325_
                index308_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                (d_1856_v122_)[index308_] = p1
                index309_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                (d_1856_v122_)[index309_] = (self).f18
                index310_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                (d_1856_v122_)[index310_] = p1
                d_1857_v123_: D17
                d_1857_v123_ = D17_DC42((self).f18)
                index311_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                index312_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                rhs332_ = (d_1856_v122_)[default__.safeIndex(664, (d_1856_v122_).length(0))]
                rhs333_ = (d_1856_v122_)[default__.safeIndex(664, (d_1856_v122_).length(0))]
                rhs334_ = (_dafny.SeqWithoutIsStrInference([(d_1846_v112_)[default__.safeIndex(521, (d_1846_v112_).length(0))] for d_1858_i14_ in range(default__.abs(-542))])) + (((self).f19) + ((self).f19))
                rhs335_ = d_1857_v123_
                rhs336_ = (d_1856_v122_)[default__.safeIndex(664, (d_1856_v122_).length(0))]
                lhs286_ = d_1856_v122_
                lhs287_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                lhs288_ = globalState
                lhs289_ = globalState
                lhs290_ = d_1856_v122_
                lhs291_ = default__.safeIndex(664, (d_1856_v122_).length(0))
                lhs286_[lhs287_] = rhs332_
                lhs288_.f2 = rhs333_
                lhs289_.f12 = rhs334_
                d_1857_v123_ = rhs335_
                lhs290_[lhs291_] = rhs336_
                d_1859_v124_: _dafny.Map
                d_1859_v124_ = _dafny.Map({self.f26: (self).f19})
                (globalState).f0 = default__.safeModuloInt((0) - (self.f26), len(d_1859_v124_))
            d_1860_v125_: _dafny.Map
            d_1860_v125_ = _dafny.Map({self.f26: p2})
            d_1861_v126_: _dafny.Seq
            d_1861_v126_ = _dafny.SeqWithoutIsStrInference([171, len(d_1860_v125_)])
            d_1862_v127_: _dafny.Map
            d_1862_v127_ = _dafny.Map({d_1861_v126_: not(p1)})
            d_1862_v127_ = d_1862_v127_
            d_1863_v128_: C1
            nw326_ = C1()
            nw326_.ctor__(p2)
            d_1863_v128_ = nw326_
            d_1864_v129_: _dafny.Map
            d_1864_v129_ = _dafny.Map({d_1863_v128_: p1})
            d_1865_v130_: D11
            d_1865_v130_ = D11_DC27(((d_1864_v129_).set(d_1863_v128_, p2)) | (d_1864_v129_))
            d_1865_v130_ = D11_DC27(d_1864_v129_)
        (globalState).f9 = ((self).f19) + ((self).f19)
        if (self).f18:
            d_1866_v131_: _dafny.Array
            nw327_ = _dafny.Array(False, 20)
            d_1866_v131_ = nw327_
            index313_ = default__.safeIndex(611, (d_1866_v131_).length(0))
            (d_1866_v131_)[index313_] = p1
            index314_ = default__.safeIndex(611, (d_1866_v131_).length(0))
            (d_1866_v131_)[index314_] = p2
            d_1867_v132_: _dafny.Array
            def lambda76_(d_1868_i15_):
                return (d_1868_i15_) - (len(_dafny.SeqWithoutIsStrInference([self.f26])))

            init45_ = lambda76_
            nw328_ = _dafny.Array(None, 26)
            for i0_45_ in range(nw328_.length(0)):
                nw328_[i0_45_] = init45_(i0_45_)
            d_1867_v132_ = nw328_
            index315_ = default__.safeIndex(915, (d_1867_v132_).length(0))
            (d_1867_v132_)[index315_] = default__.safeModuloInt(self.f26, 536)
            d_1869_v133_: _dafny.Map
            d_1869_v133_ = _dafny.Map({not((self).f18): (0) - (self.f26)})
            index316_ = default__.safeIndex(915, (d_1867_v132_).length(0))
            index317_ = default__.safeIndex(611, (d_1866_v131_).length(0))
            rhs337_ = (self.f26) < (self.f26)
            rhs338_ = self.f26
            rhs339_ = (self).fm2(p2, (-148) < (len(d_1869_v133_)), ((0) - (default__.fm1(len(_dafny.SeqWithoutIsStrInference([self.f26])), (d_1866_v131_)[default__.safeIndex(611, (d_1866_v131_).length(0))], (d_1866_v131_)[default__.safeIndex(611, (d_1866_v131_).length(0))], globalState))) < (self.f26), globalState)
            lhs292_ = globalState
            lhs293_ = d_1867_v132_
            lhs294_ = default__.safeIndex(915, (d_1867_v132_).length(0))
            lhs295_ = d_1866_v131_
            lhs296_ = default__.safeIndex(611, (d_1866_v131_).length(0))
            lhs292_.f2 = rhs337_
            lhs293_[lhs294_] = rhs338_
            lhs295_[lhs296_] = rhs339_
            d_1870_v134_: _dafny.Set
            d_1870_v134_ = _dafny.Set({(d_1866_v131_)[default__.safeIndex(611, (d_1866_v131_).length(0))]})
            d_1871_v135_: _dafny.Map
            d_1871_v135_ = _dafny.Map({False: (d_1870_v134_).isdisjoint(d_1870_v134_)})
            d_1871_v135_ = (d_1871_v135_).set(p2, (False) and (p2))
            d_1872_v136_: _dafny.Array
            nw329_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_1872_v136_ = nw329_
            d_1873_v137_: _dafny.Array
            d_1873_v137_ = d_1867_v132_
            index318_ = default__.safeIndex(461, (d_1872_v136_).length(0))
            (d_1872_v136_)[index318_] = ((d_1873_v137_) if False else d_1867_v132_)
            index319_ = default__.safeIndex(461, (d_1872_v136_).length(0))
            (d_1872_v136_)[index319_] = d_1867_v132_
            d_1874_v138_: str
            d_1874_v138_ = _dafny.CodePoint('d')
            d_1875_v139_: _dafny.Set
            d_1875_v139_ = _dafny.Set({d_1874_v138_})
            r0 = (default__.safeDivisionInt((d_1867_v132_)[default__.safeIndex(915, (d_1867_v132_).length(0))], len(d_1875_v139_)) if p1 else self.f26)
        elif True:
            (globalState).f0 = self.f26
            d_1876_v140_: _dafny.MultiSet
            d_1876_v140_ = _dafny.MultiSet([(self).f18])
            d_1877_v141_: _dafny.MultiSet
            d_1877_v141_ = _dafny.MultiSet([d_1876_v140_, d_1876_v140_, d_1876_v140_, d_1876_v140_])
            d_1878_v142_: _dafny.Map
            d_1878_v142_ = _dafny.Map({self.f26: p1})
            d_1879_v143_: _dafny.MultiSet
            d_1879_v143_ = _dafny.MultiSet([len(default__.fm49(((D19_DC47(d_1877_v141_)).cf81).cardinality, len(d_1878_v142_), self.f26, globalState)), (_dafny.MultiSet([(self).f18, p1])).cardinality, self.f26])
            d_1880_v144_: _dafny.Seq
            d_1880_v144_ = _dafny.SeqWithoutIsStrInference([self.f26])
            d_1881_v145_: _dafny.Set
            d_1881_v145_ = _dafny.Set({p1, p2})
            d_1882_v146_: str
            d_1882_v146_ = _dafny.CodePoint('b')
            d_1883_v147_: _dafny.Map
            d_1883_v147_ = _dafny.Map({d_1882_v146_: d_1882_v146_})
            d_1884_v148_: D13
            d_1884_v148_ = D13_DC33(False, self.f26, (self).f18, len((self).f19))
            d_1885_v149_: _dafny.Map
            d_1885_v149_ = _dafny.Map({(d_1884_v148_).cf55: self.f26})
            d_1886_v150_: _dafny.Map
            d_1886_v150_ = _dafny.Map({d_1883_v147_: len(d_1885_v149_)})
            d_1887_v151_: _dafny.Array
            nw330_ = _dafny.Array(None, 27)
            nw330_[int(0)] = not(p2)
            nw330_[int(1)] = not(not((self).f18))
            nw330_[int(2)] = p1
            nw330_[int(3)] = (self).f18
            nw330_[int(4)] = (d_1879_v143_) == (_dafny.MultiSet([self.f26]))
            nw330_[int(5)] = p2
            nw330_[int(6)] = not((self).f18)
            nw330_[int(7)] = p2
            nw330_[int(8)] = (d_1880_v144_) < (d_1880_v144_)
            nw330_[int(9)] = p2
            nw330_[int(10)] = (self).f18
            nw330_[int(11)] = (d_1881_v145_).ispropersubset(d_1881_v145_)
            nw330_[int(12)] = (d_1881_v145_).isdisjoint(default__.fm36(p1, len(d_1886_v150_), False, globalState))
            nw330_[int(13)] = (self).f18
            nw330_[int(14)] = p2
            nw330_[int(15)] = p1
            nw330_[int(16)] = False
            nw330_[int(17)] = p1
            nw330_[int(18)] = p2
            nw330_[int(19)] = p1
            nw330_[int(20)] = p1
            nw330_[int(21)] = p1
            nw330_[int(22)] = p1
            nw330_[int(23)] = p2
            nw330_[int(24)] = (p2) not in (d_1876_v140_)
            nw330_[int(25)] = (self.f26) > (self.f26)
            nw330_[int(26)] = (self).f18
            d_1887_v151_ = nw330_
            d_1888_v152_: _dafny.Set
            d_1888_v152_ = _dafny.Set({self.f26})
            index320_ = default__.safeIndex(61, (d_1887_v151_).length(0))
            (d_1887_v151_)[index320_] = not(default__.fm23(d_1888_v152_, self.f26, globalState))
            index321_ = default__.safeIndex(61, (d_1887_v151_).length(0))
            (d_1887_v151_)[index321_] = not(p1)
            d_1889_v153_: _dafny.Seq
            d_1889_v153_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            d_1890_v154_: _dafny.Set
            d_1890_v154_ = _dafny.Set({(d_1889_v153_)[default__.safeIndex(self.f26, len(d_1889_v153_))]})
            d_1891_v155_: _dafny.Map
            d_1891_v155_ = _dafny.Map({(self.f26) != (self.f26): d_1890_v154_})
            d_1891_v155_ = (d_1891_v155_).set(not((d_1887_v151_)[default__.safeIndex(61, (d_1887_v151_).length(0))]), d_1890_v154_)
            d_1892_v156_: _dafny.Array
            nw331_ = _dafny.Array(int(0), 23)
            d_1892_v156_ = nw331_
            d_1893_v157_: _dafny.Map
            d_1893_v157_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfnihgwpf")): (d_1884_v148_).cf58})
            index322_ = default__.safeIndex(85, (d_1892_v156_).length(0))
            (d_1892_v156_)[index322_] = len(d_1893_v157_)
            index323_ = default__.safeIndex(85, (d_1892_v156_).length(0))
            (d_1892_v156_)[index323_] = len(default__.fm58((self.f26) >= (self.f26), p1, globalState))
            (globalState).f2 = ((True) == ((self).f18) if (d_1879_v143_).issubset(d_1879_v143_) else (d_1887_v151_)[default__.safeIndex(61, (d_1887_v151_).length(0))])
        if (self).fm3(self.f26, globalState):
            (globalState).f0 = self.f26
            d_1894_v158_: str
            d_1894_v158_ = _dafny.CodePoint('g')
            d_1894_v158_ = d_1894_v158_
            d_1895_v159_: _dafny.Set
            d_1895_v159_ = _dafny.Set({self.f26})
            def iife151_():
                coll49_ = _dafny.Set()
                compr_49_: int
                for compr_49_ in _dafny.IntegerRange(981, 989):
                    d_1896_v160_: int = compr_49_
                    if ((981) <= (d_1896_v160_)) and ((d_1896_v160_) < (989)):
                        coll49_ = coll49_.union(_dafny.Set([(d_1896_v160_) * (self.f26)]))
                return _dafny.Set(coll49_)
            d_1895_v159_ = iife151_()
            
            d_1897_v161_: _dafny.Seq
            d_1897_v161_ = _dafny.SeqWithoutIsStrInference([self.f26])
            d_1898_v162_: _dafny.Seq
            d_1898_v162_ = _dafny.SeqWithoutIsStrInference([(self).f19])
            d_1899_v163_: C3
            nw332_ = C3()
            nw332_.ctor__(d_1897_v161_, ((d_1898_v162_)[default__.safeIndex(len(d_1895_v159_), len(d_1898_v162_))]).set(default__.safeIndex(4, len((d_1898_v162_)[default__.safeIndex(len(d_1895_v159_), len(d_1898_v162_))])), d_1894_v158_), (self).f20, (290) <= (self.f26))
            d_1899_v163_ = nw332_
            d_1900_v164_: _dafny.Array
            def lambda77_(d_1901_i16_):
                return (d_1901_i16_) - (self.f26)

            init46_ = lambda77_
            nw333_ = _dafny.Array(None, 26)
            for i0_46_ in range(nw333_.length(0)):
                nw333_[i0_46_] = init46_(i0_46_)
            d_1900_v164_ = nw333_
            d_1902_v165_: _dafny.MultiSet
            d_1902_v165_ = _dafny.MultiSet([(self).f18])
            rhs340_ = d_1900_v164_
            rhs341_ = d_1902_v165_
            d_1900_v164_ = rhs340_
            d_1902_v165_ = rhs341_
        elif True:
            d_1903_v166_: D13
            d_1903_v166_ = D13_DC33(p1, self.f26, p2, self.f26)
            d_1904_v167_: _dafny.MultiSet
            d_1904_v167_ = _dafny.MultiSet([p1, (self).f18])
            d_1905_v168_: _dafny.Array
            nw334_ = _dafny.Array(None, 20)
            nw334_[int(0)] = (d_1903_v166_).cf57
            nw334_[int(1)] = (False if p1 else (self).f18)
            nw334_[int(2)] = p1
            nw334_[int(3)] = p2
            nw334_[int(4)] = not((self).f18)
            nw334_[int(5)] = p2
            nw334_[int(6)] = p1
            nw334_[int(7)] = (self).f18
            nw334_[int(8)] = p1
            nw334_[int(9)] = (self).f18
            nw334_[int(10)] = not (False) or ((self).f18)
            nw334_[int(11)] = not((self).f18)
            nw334_[int(12)] = p1
            nw334_[int(13)] = p1
            nw334_[int(14)] = p1
            nw334_[int(15)] = p1
            nw334_[int(16)] = not((d_1904_v167_) != (d_1904_v167_))
            nw334_[int(17)] = (self).f18
            nw334_[int(18)] = (self).fm3(204, globalState)
            nw334_[int(19)] = p1
            d_1905_v168_ = nw334_
            index324_ = default__.safeIndex(74, (d_1905_v168_).length(0))
            (d_1905_v168_)[index324_] = p2
            index325_ = default__.safeIndex(74, (d_1905_v168_).length(0))
            (d_1905_v168_)[index325_] = (self).fm2(True, p2, p1, globalState)
            r0 = 487
            d_1906_v169_: str
            d_1906_v169_ = _dafny.CodePoint('y')
            d_1907_v170_: _dafny.Seq
            d_1907_v170_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f19])
            d_1908_v171_: _dafny.Array
            nw335_ = _dafny.Array(None, 26)
            nw335_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlqtclqmw"))
            nw335_[int(1)] = (self).f19
            nw335_[int(2)] = (self).f19
            nw335_[int(3)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1909_i17_ in range(default__.abs(610))])).set(default__.safeIndex(self.f26, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1910_i17_ in range(default__.abs(610))]))), d_1906_v169_)
            nw335_[int(4)] = (((self).f19).set(default__.safeIndex(self.f26, len((self).f19)), d_1906_v169_)) + ((self).f19)
            nw335_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rskscm"))
            nw335_[int(6)] = default__.fm48(globalState)
            nw335_[int(7)] = (self).f19
            nw335_[int(8)] = _dafny.SeqWithoutIsStrInference([d_1906_v169_ for d_1911_i18_ in range(default__.abs(618))])
            nw335_[int(9)] = (_dafny.SeqWithoutIsStrInference([d_1906_v169_ for d_1912_i19_ in range(default__.abs(728))])) + ((self).f19)
            nw335_[int(10)] = (d_1907_v170_)[default__.safeIndex(237, len(d_1907_v170_))]
            nw335_[int(11)] = ((self).f19) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfpoexew")))
            nw335_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkvw"))
            nw335_[int(13)] = (self).f19
            nw335_[int(14)] = default__.fm16(self.f26, False, d_1906_v169_, self.f26, globalState)
            nw335_[int(15)] = (self).f19
            nw335_[int(16)] = (self).f19
            nw335_[int(17)] = (self).f19
            nw335_[int(18)] = (self).f19
            nw335_[int(19)] = ((self).f19) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1913_i20_ in range(default__.abs(855))]))
            nw335_[int(20)] = _dafny.SeqWithoutIsStrInference([d_1906_v169_ for d_1914_i21_ in range(default__.abs(40))])
            nw335_[int(21)] = (self).f19
            nw335_[int(22)] = (((self).f19).set(default__.safeIndex(len((self).f19), len((self).f19)), d_1906_v169_)) + ((self).f19)
            nw335_[int(23)] = (self).f19
            nw335_[int(24)] = (self).f19
            nw335_[int(25)] = (self).f19
            d_1908_v171_ = nw335_
            index326_ = default__.safeIndex(195, (d_1908_v171_).length(0))
            (d_1908_v171_)[index326_] = (self).f19
            index327_ = default__.safeIndex(195, (d_1908_v171_).length(0))
            rhs342_ = False
            rhs343_ = (((self).f19) + ((self).f19)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrmigu")))
            rhs344_ = ((0) - (self.f26)) > (self.f26)
            rhs345_ = (_dafny.SeqWithoutIsStrInference([d_1906_v169_ for d_1915_i22_ in range(default__.abs(162))])) + ((self).f19)
            rhs346_ = d_1906_v169_
            lhs297_ = globalState
            lhs298_ = globalState
            lhs299_ = globalState
            lhs300_ = d_1908_v171_
            lhs301_ = default__.safeIndex(195, (d_1908_v171_).length(0))
            lhs297_.f2 = rhs342_
            lhs298_.f9 = rhs343_
            lhs299_.f2 = rhs344_
            lhs300_[lhs301_] = rhs345_
            d_1906_v169_ = rhs346_
            d_1906_v169_ = d_1906_v169_
            (globalState).f15 = (default__.fm59(globalState)).cf33
        r0 = self.f26
        return r0

    def m6(self, globalState):
        r0: bool = False
        r1: int = int(0)
        d_1916_v0_: T0
        nw336_ = C1()
        nw336_.ctor__((self).f18)
        d_1916_v0_ = nw336_
        d_1917_v1_: _dafny.Set
        d_1917_v1_ = _dafny.Set({d_1916_v0_})
        d_1918_v2_: _dafny.Seq
        d_1918_v2_ = _dafny.SeqWithoutIsStrInference([len((self).f19), len(d_1917_v1_)])
        d_1919_v3_: _dafny.Map
        d_1919_v3_ = _dafny.Map({default__.fm33((d_1916_v0_).f18, globalState): d_1918_v2_})
        d_1920_v4_: _dafny.Array
        nw337_ = _dafny.Array(None, 12)
        nw337_[int(0)] = d_1918_v2_
        nw337_[int(1)] = (d_1918_v2_) + (d_1918_v2_)
        nw337_[int(2)] = (_dafny.SeqWithoutIsStrInference([self.f26, self.f26, self.f26, len((self).f19)])) + (d_1918_v2_)
        nw337_[int(3)] = (((d_1919_v3_)[d_1918_v2_] if (d_1918_v2_) in (d_1919_v3_) else d_1918_v2_)) + (d_1918_v2_)
        nw337_[int(4)] = (_dafny.SeqWithoutIsStrInference([self.f26 for d_1921_i0_ in range(default__.abs(964))])) + (d_1918_v2_)
        nw337_[int(5)] = d_1918_v2_
        nw337_[int(6)] = d_1918_v2_
        nw337_[int(7)] = _dafny.SeqWithoutIsStrInference([len((self).f19), self.f26, self.f26])
        nw337_[int(8)] = d_1918_v2_
        nw337_[int(9)] = (_dafny.SeqWithoutIsStrInference([self.f26 for d_1922_i1_ in range(default__.abs(696))]) if (d_1916_v0_).f18 else (d_1918_v2_).set(default__.safeIndex(82, len(d_1918_v2_)), self.f26))
        nw337_[int(10)] = d_1918_v2_
        nw337_[int(11)] = d_1918_v2_
        d_1920_v4_ = nw337_
        d_1920_v4_ = d_1920_v4_
        hi10_ = self.f26
        for d_1923_i2_ in range(len(((self).f19 if (d_1916_v0_).f18 else (self).f19)), hi10_):
            r0 = ((self).f18 if (self).f18 else ((d_1916_v0_).f18) == ((d_1916_v0_).f18))
            (globalState).f0 = self.f26
            d_1924_v5_: str
            d_1924_v5_ = _dafny.CodePoint('a')
            (globalState).f12 = ((self).f19) + (((self).f19).set(default__.safeIndex(538, len((self).f19)), d_1924_v5_))
            r0 = (self).fm2((d_1916_v0_).f18, (self).f18, (self.f26) >= (d_1923_i2_), globalState)
        if (self).f18:
            d_1925_v6_: _dafny.Array
            def lambda78_(d_1926_i3_):
                return ((self).f25) < ((self).f25)

            init47_ = lambda78_
            nw338_ = _dafny.Array(None, 22)
            for i0_47_ in range(nw338_.length(0)):
                nw338_[i0_47_] = init47_(i0_47_)
            d_1925_v6_ = nw338_
            d_1927_v7_: _dafny.Set
            d_1927_v7_ = _dafny.Set({self.f26})
            index328_ = default__.safeIndex(110, (d_1925_v6_).length(0))
            (d_1925_v6_)[index328_] = default__.fm23(d_1927_v7_, len((self).f19), globalState)
            index329_ = default__.safeIndex(110, (d_1925_v6_).length(0))
            (d_1925_v6_)[index329_] = (self).f18
            index330_ = default__.safeIndex(110, (d_1925_v6_).length(0))
            (d_1925_v6_)[index330_] = not(not(((self).f19) < ((self).f19)))
            (globalState).f9 = ((self).f19) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1928_i4_ in range(default__.abs(820))]))
            d_1929_v8_: _dafny.Array
            nw339_ = _dafny.Array(None, 26)
            d_1929_v8_ = nw339_
            d_1929_v8_ = d_1929_v8_
            if (d_1916_v0_).f18:
                d_1930_v9_: _dafny.Map
                d_1930_v9_ = _dafny.Map({532: self.f26})
                d_1931_v10_: _dafny.Set
                d_1931_v10_ = _dafny.Set({(D20_DC52(d_1930_v9_)).cf86})
                d_1932_v12_: D21
                def iife152_():
                    coll50_ = _dafny.Set()
                    compr_50_: _dafny.Map
                    for compr_50_ in ((self).f20).Elements:
                        d_1933_v11_: _dafny.Map = compr_50_
                        if (d_1933_v11_) in ((self).f20):
                            coll50_ = coll50_.union(_dafny.Set([d_1933_v11_]))
                    return _dafny.Set(coll50_)
                d_1932_v12_ = D21_DC55(iife152_()
)
                d_1931_v10_ = (d_1932_v12_).cf91
                d_1934_v13_: _dafny.MultiSet
                d_1934_v13_ = _dafny.MultiSet([(self).f18])
                (globalState).f3 = ((_dafny.MultiSet([(d_1916_v0_).f18]) if (d_1916_v0_).f18 else (d_1934_v13_).intersection(_dafny.MultiSet((self).f25)))).cardinality
                index331_ = default__.safeIndex(110, (d_1925_v6_).length(0))
                (d_1925_v6_)[index331_] = not(not(not((self).f18)))
                d_1935_v14_: _dafny.Array
                nw340_ = _dafny.Array(int(0), 7)
                d_1935_v14_ = nw340_
                d_1936_v15_: _dafny.Map
                d_1936_v15_ = _dafny.Map({len((self).f19): d_1930_v9_})
                index332_ = default__.safeIndex(209, (d_1935_v14_).length(0))
                (d_1935_v14_)[index332_] = len(((d_1936_v15_)[self.f26] if (self.f26) in (d_1936_v15_) else _dafny.Map({self.f26: 888})))
                index333_ = default__.safeIndex(209, (d_1935_v14_).length(0))
                (d_1935_v14_)[index333_] = (self.f26) + (self.f26)
                index334_ = default__.safeIndex(209, (d_1935_v14_).length(0))
                (d_1935_v14_)[index334_] = (d_1935_v14_)[default__.safeIndex(209, (d_1935_v14_).length(0))]
            elif True:
                d_1937_v16_: _dafny.Seq
                d_1937_v16_ = _dafny.SeqWithoutIsStrInference([((self).f25).set(default__.safeIndex(self.f26, len((self).f25)), (d_1916_v0_).f18)])
                d_1938_v17_: _dafny.Map
                d_1938_v17_ = _dafny.Map({d_1937_v16_: self.f26})
                d_1938_v17_ = (d_1938_v17_).set(d_1937_v16_, self.f26)
                d_1939_v18_: _dafny.Map
                d_1939_v18_ = _dafny.Map({(d_1925_v6_)[default__.safeIndex(110, (d_1925_v6_).length(0))]: (self).f18})
                d_1940_v19_: _dafny.Map
                d_1940_v19_ = _dafny.Map({len(d_1939_v18_): self.f26})
                d_1941_v20_: _dafny.Seq
                d_1941_v20_ = _dafny.SeqWithoutIsStrInference([((d_1918_v2_).set(default__.safeIndex(len(d_1940_v19_), len(d_1918_v2_)), self.f26)).set(default__.safeIndex(self.f26, len((d_1918_v2_).set(default__.safeIndex(len(d_1940_v19_), len(d_1918_v2_)), self.f26))), self.f26), d_1918_v2_])
                d_1942_v21_: C3
                nw341_ = C3()
                nw341_.ctor__((d_1941_v20_)[default__.safeIndex(self.f26, len(d_1941_v20_))], ((self).f19) + ((self).f19), (self).f20, (self).f18)
                d_1942_v21_ = nw341_
                d_1943_v22_: _dafny.Array
                nw342_ = _dafny.Array(int(0), 19)
                d_1943_v22_ = nw342_
                index335_ = default__.safeIndex(624, (d_1943_v22_).length(0))
                (d_1943_v22_)[index335_] = self.f26
                index336_ = default__.safeIndex(624, (d_1943_v22_).length(0))
                (d_1943_v22_)[index336_] = len((self).f19)
                (globalState).f2 = (d_1916_v0_).f18
                (globalState).f3 = self.f26
        elif True:
            d_1944_v23_: _dafny.MultiSet
            d_1944_v23_ = _dafny.MultiSet([(self).f18, (self).f18])
            d_1945_v24_: _dafny.Seq
            d_1945_v24_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_1946_v25_: _dafny.Array
            nw343_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_1946_v25_ = nw343_
            d_1947_v26_: _dafny.MultiSet
            d_1947_v26_ = d_1944_v23_
            d_1948_v27_: str
            d_1948_v27_ = _dafny.CodePoint('b')
            rhs347_ = (d_1916_v0_).f18
            rhs348_ = (d_1947_v26_)
            rhs349_ = (self.f26) > (self.f26)
            rhs350_ = default__.fm44((self.f26) - (self.f26), (d_1916_v0_).f18, d_1948_v27_, -143, globalState)
            rhs351_ = d_1946_v25_
            lhs302_ = globalState
            lhs303_ = globalState
            lhs302_.f2 = rhs347_
            d_1944_v23_ = rhs348_
            lhs303_.f2 = rhs349_
            d_1945_v24_ = rhs350_
            d_1946_v25_ = rhs351_
            (globalState).f2 = (d_1916_v0_).f18
            d_1949_v28_: _dafny.Map
            d_1949_v28_ = _dafny.Map({self.f26: self.f26})
            (globalState).f15 = (self.f26) + (default__.safeDivisionInt(len(d_1918_v2_), len(d_1949_v28_)))
            if False:
                d_1950_v29_: _dafny.Array
                def lambda79_(d_1951_i5_):
                    return (d_1951_i5_) * (self.f26)

                init48_ = lambda79_
                nw344_ = _dafny.Array(None, 10)
                for i0_48_ in range(nw344_.length(0)):
                    nw344_[i0_48_] = init48_(i0_48_)
                d_1950_v29_ = nw344_
                index337_ = default__.safeIndex(31, (d_1950_v29_).length(0))
                (d_1950_v29_)[index337_] = self.f26
                d_1952_v30_: _dafny.Array
                nw345_ = _dafny.Array(False, 21)
                d_1952_v30_ = nw345_
                d_1953_v31_: _dafny.Set
                d_1953_v31_ = _dafny.Set({d_1952_v30_, d_1952_v30_})
                d_1954_v32_: _dafny.MultiSet
                d_1954_v32_ = _dafny.MultiSet([self.f26, self.f26])
                index338_ = default__.safeIndex(31, (d_1950_v29_).length(0))
                rhs352_ = (d_1916_v0_).f18
                rhs353_ = self.f26
                rhs354_ = (d_1953_v31_).isdisjoint(d_1953_v31_)
                rhs355_ = not(((d_1954_v32_).cardinality) == (self.f26))
                rhs356_ = default__.safeModuloInt((self.f26) + ((d_1918_v2_)[default__.safeIndex(len((self).f19), len(d_1918_v2_))]), self.f26)
                lhs304_ = globalState
                lhs305_ = d_1950_v29_
                lhs306_ = default__.safeIndex(31, (d_1950_v29_).length(0))
                lhs307_ = globalState
                lhs308_ = globalState
                lhs309_ = globalState
                lhs304_.f2 = rhs352_
                lhs305_[lhs306_] = rhs353_
                lhs307_.f2 = rhs354_
                lhs308_.f2 = rhs355_
                lhs309_.f3 = rhs356_
                d_1955_v33_: _dafny.Array
                d_1955_v33_ = d_1952_v30_
                d_1956_v34_: _dafny.Map
                d_1956_v34_ = _dafny.Map({d_1955_v33_: (d_1916_v0_).f18})
                r0 = not(((d_1956_v34_)[d_1952_v30_] if (d_1952_v30_) in (d_1956_v34_) else (d_1916_v0_).f18))
                d_1957_v35_: _dafny.Array
                nw346_ = _dafny.Array(None, 14)
                nw346_[int(0)] = d_1948_v27_
                nw346_[int(1)] = ((self).f19)[default__.safeIndex((d_1950_v29_)[default__.safeIndex(31, (d_1950_v29_).length(0))], len((self).f19))]
                nw346_[int(2)] = d_1948_v27_
                nw346_[int(3)] = (d_1948_v27_ if False else d_1948_v27_)
                nw346_[int(4)] = d_1948_v27_
                nw346_[int(5)] = d_1948_v27_
                nw346_[int(6)] = d_1948_v27_
                nw346_[int(7)] = d_1948_v27_
                nw346_[int(8)] = d_1948_v27_
                nw346_[int(9)] = d_1948_v27_
                nw346_[int(10)] = d_1948_v27_
                nw346_[int(11)] = d_1948_v27_
                nw346_[int(12)] = d_1948_v27_
                nw346_[int(13)] = _dafny.CodePoint('v')
                d_1957_v35_ = nw346_
                index339_ = default__.safeIndex(479, (d_1957_v35_).length(0))
                (d_1957_v35_)[index339_] = ((self).f19)[default__.safeIndex((d_1950_v29_)[default__.safeIndex(31, (d_1950_v29_).length(0))], len((self).f19))]
                index340_ = default__.safeIndex(479, (d_1957_v35_).length(0))
                (d_1957_v35_)[index340_] = default__.fm46(globalState)
                rhs357_ = d_1950_v29_
                rhs358_ = d_1948_v27_
                d_1950_v29_ = rhs357_
                d_1948_v27_ = rhs358_
                index341_ = default__.safeIndex(31, (d_1950_v29_).length(0))
                (d_1950_v29_)[index341_] = default__.fm1(((d_1950_v29_)[default__.safeIndex(31, (d_1950_v29_).length(0))]) * (623), False, (d_1916_v0_).f18, globalState)
            elif True:
                (globalState).f3 = self.f26
                d_1958_v36_: D12
                d_1958_v36_ = D12_DC30((self).f25)
                d_1959_v37_: _dafny.Seq
                d_1959_v37_ = _dafny.SeqWithoutIsStrInference([d_1958_v36_])
                d_1960_v38_: _dafny.Set
                d_1960_v38_ = _dafny.Set({self.f26})
                d_1961_v39_: _dafny.Map
                d_1961_v39_ = _dafny.Map({self.f26: d_1960_v38_})
                d_1962_v40_: _dafny.Map
                d_1962_v40_ = _dafny.Map({(d_1916_v0_).f18: self.f26})
                d_1963_v41_: _dafny.Set
                d_1963_v41_ = _dafny.Set({(D19_DC50((d_1916_v0_).f18)).cf84, (d_1916_v0_).f18})
                d_1964_v42_: _dafny.Map
                d_1964_v42_ = _dafny.Map({(d_1916_v0_).f18: d_1918_v2_})
                d_1965_v43_: _dafny.MultiSet
                d_1965_v43_ = _dafny.MultiSet([self.f26, len(d_1963_v41_), len(_dafny.SeqWithoutIsStrInference([-917 for d_1966_i6_ in range(default__.abs(-538))]))])
                d_1967_v44_: _dafny.Array
                nw347_ = _dafny.Array(None, 27)
                nw347_[int(0)] = self.f26
                nw347_[int(1)] = len(d_1961_v39_)
                nw347_[int(2)] = self.f26
                nw347_[int(3)] = 302
                nw347_[int(4)] = len(d_1962_v40_)
                nw347_[int(5)] = self.f26
                nw347_[int(6)] = default__.fm1(len(d_1963_v41_), (d_1916_v0_).f18, not((d_1916_v0_).f18), globalState)
                nw347_[int(7)] = self.f26
                nw347_[int(8)] = self.f26
                nw347_[int(9)] = default__.fm1(self.f26, (d_1916_v0_).f18, (d_1916_v0_).f18, globalState)
                nw347_[int(10)] = self.f26
                nw347_[int(11)] = (0) - (default__.safeModuloInt(len(d_1963_v41_), self.f26))
                nw347_[int(12)] = self.f26
                nw347_[int(13)] = (self.f26 if (d_1916_v0_).f18 else self.f26)
                nw347_[int(14)] = default__.safeModuloInt(len(d_1949_v28_), self.f26)
                nw347_[int(15)] = (self.f26) * (self.f26)
                nw347_[int(16)] = (self.f26) * (self.f26)
                nw347_[int(17)] = self.f26
                nw347_[int(18)] = self.f26
                nw347_[int(19)] = len(d_1945_v24_)
                nw347_[int(20)] = len(((d_1964_v42_)[(d_1916_v0_).f18] if ((d_1916_v0_).f18) in (d_1964_v42_) else _dafny.SeqWithoutIsStrInference([(d_1965_v43_).cardinality])))
                nw347_[int(21)] = (self.f26 if (self).f18 else self.f26)
                nw347_[int(22)] = (0) - (self.f26)
                nw347_[int(23)] = len((self).f25)
                nw347_[int(24)] = 822
                nw347_[int(25)] = (default__.fm32(globalState)).cardinality
                nw347_[int(26)] = self.f26
                d_1967_v44_ = nw347_
                d_1968_v45_: _dafny.Map
                d_1968_v45_ = _dafny.Map({(d_1916_v0_).f18: (d_1916_v0_).f18})
                d_1969_v46_: _dafny.MultiSet
                d_1969_v46_ = _dafny.MultiSet([d_1967_v44_, d_1967_v44_, d_1967_v44_])
                rhs359_ = d_1959_v37_
                rhs360_ = (self).f18
                rhs361_ = (d_1916_v0_).f18
                rhs362_ = ((d_1968_v45_)[(d_1969_v46_).isdisjoint(d_1969_v46_)] if ((d_1969_v46_).isdisjoint(d_1969_v46_)) in (d_1968_v45_) else (d_1916_v0_).f18)
                rhs363_ = d_1967_v44_
                lhs310_ = globalState
                lhs311_ = globalState
                lhs312_ = globalState
                d_1959_v37_ = rhs359_
                lhs310_.f2 = rhs360_
                lhs311_.f2 = rhs361_
                lhs312_.f2 = rhs362_
                d_1967_v44_ = rhs363_
                d_1970_v47_: _dafny.Map
                d_1970_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")): _dafny.Map({default__.fm35(self.f26, (self).f18, globalState): self.f26})})
                d_1971_v48_: D0
                d_1971_v48_ = D0_DC1((self).f18, self.f26, (self).f19, (d_1916_v0_).f18, self.f26)
                d_1972_v49_: _dafny.Map
                d_1972_v49_ = _dafny.Map({d_1971_v48_: len(d_1945_v24_)})
                d_1970_v47_ = (d_1970_v47_).set((self).f19, d_1972_v49_)
                d_1973_v50_: C0
                nw348_ = C0()
                nw348_.ctor__(self.f26, ((0) - (self.f26)) - (default__.fm1(self.f26, (d_1916_v0_).f18, not((d_1916_v0_).f18), globalState)))
                d_1973_v50_ = nw348_
                d_1974_v51_: _dafny.Array
                def lambda80_(d_1975_v24_):
                    def lambda81_(d_1976_i7_):
                        return d_1975_v24_

                    return lambda81_

                init49_ = lambda80_(d_1945_v24_)
                nw349_ = _dafny.Array(None, 28)
                for i0_49_ in range(nw349_.length(0)):
                    nw349_[i0_49_] = init49_(i0_49_)
                d_1974_v51_ = nw349_
                d_1977_v52_: _dafny.Map
                d_1977_v52_ = _dafny.Map({825: d_1974_v51_})
                d_1977_v52_ = (d_1977_v52_).set(((d_1973_v50_).f32) + (self.f26), d_1974_v51_)
            d_1978_v53_: _dafny.Set
            d_1978_v53_ = _dafny.Set({len(d_1918_v2_)})
            d_1979_v54_: _dafny.Seq
            d_1979_v54_ = _dafny.SeqWithoutIsStrInference([(self).f19])
            d_1980_v55_: _dafny.MultiSet
            d_1980_v55_ = _dafny.MultiSet([(d_1979_v54_)[default__.safeIndex((0) - (self.f26), len(d_1979_v54_))], (self).f19])
            d_1945_v24_ = default__.fm44(self.f26, default__.fm23(d_1978_v53_, self.f26, globalState), d_1948_v27_, (d_1918_v2_)[default__.safeIndex(((d_1980_v55_)[_dafny.SeqWithoutIsStrInference([d_1948_v27_])] if (_dafny.SeqWithoutIsStrInference([d_1948_v27_])) in (d_1980_v55_) else default__.fm1(self.f26, (d_1916_v0_).f18, (d_1916_v0_).f18, globalState)), len(d_1918_v2_))], globalState)
        d_1981_v56_: _dafny.Array
        def lambda82_(d_1982_i8_):
            return (self).f19

        init50_ = lambda82_
        nw350_ = _dafny.Array(None, 12)
        for i0_50_ in range(nw350_.length(0)):
            nw350_[i0_50_] = init50_(i0_50_)
        d_1981_v56_ = nw350_
        index342_ = default__.safeIndex(138, (d_1981_v56_).length(0))
        (d_1981_v56_)[index342_] = (self).f19
        index343_ = default__.safeIndex(138, (d_1981_v56_).length(0))
        (d_1981_v56_)[index343_] = ((self).f19) + ((self).f19)
        d_1983_i9_: int
        d_1983_i9_ = 0
        with _dafny.label("13"):
            while (d_1916_v0_).f18:
                with _dafny.c_label("13"):
                    if (d_1983_i9_) >= (100):
                        raise _dafny.Break("13")
                    d_1983_i9_ = (d_1983_i9_) + (1)
                    d_1984_v57_: _dafny.Array
                    def lambda83_(d_1985_i10_):
                        return False

                    init51_ = lambda83_
                    nw351_ = _dafny.Array(None, 24)
                    for i0_51_ in range(nw351_.length(0)):
                        nw351_[i0_51_] = init51_(i0_51_)
                    d_1984_v57_ = nw351_
                    d_1986_v58_: _dafny.Set
                    d_1986_v58_ = _dafny.Set({d_1984_v57_, d_1984_v57_})
                    d_1987_v59_: _dafny.Array
                    d_1987_v59_ = d_1984_v57_
                    d_1988_v60_: _dafny.Map
                    d_1988_v60_ = _dafny.Map({len((d_1986_v58_).intersection(d_1986_v58_)): (d_1987_v59_)})
                    d_1988_v60_ = (d_1988_v60_) | (d_1988_v60_)
                    d_1989_v61_: str
                    d_1989_v61_ = _dafny.CodePoint('i')
                    d_1990_v62_: C5
                    nw352_ = C5()
                    nw352_.ctor__(((d_1981_v56_)[default__.safeIndex(138, (d_1981_v56_).length(0))]) + ((self).f19), ((d_1981_v56_)[default__.safeIndex(138, (d_1981_v56_).length(0))]) + ((d_1981_v56_)[default__.safeIndex(138, (d_1981_v56_).length(0))]), (d_1989_v61_) not in ((self).f19))
                    d_1990_v62_ = nw352_
                    d_1991_v63_: _dafny.Map
                    d_1991_v63_ = _dafny.Map({d_1989_v61_: len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))})
                    d_1991_v63_ = (d_1991_v63_).set(default__.fm46(globalState), self.f26)
                    d_1992_v65_: C4
                    nw353_ = C4()
                    def iife153_():
                        coll51_ = _dafny.Map()
                        compr_51_: int
                        for compr_51_ in (_dafny.SeqWithoutIsStrInference([self.f26])).Elements:
                            d_1993_v64_: int = compr_51_
                            if (d_1993_v64_) in (_dafny.SeqWithoutIsStrInference([self.f26])):
                                coll51_[(d_1993_v64_) - (self.f26)] = len((self).f25)
                        return _dafny.Map(coll51_)
                    nw353_.ctor__(d_1989_v61_, self.f26, (self.f26) == ((d_1918_v2_)[default__.safeIndex(self.f26, len(d_1918_v2_))]), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgfvxeb"))) + ((d_1990_v62_).f28), ((self).f20).set(default__.safeIndex(self.f26, len((self).f20)), iife153_()
                    ))
                    d_1992_v65_ = nw353_
                    pass
            pass
        (globalState).f2 = not((self.f26) >= ((self.f26 if True else self.f26)))
        r0 = (self).f18
        d_1994_v66_: _dafny.Map
        d_1994_v66_ = _dafny.Map({not((d_1916_v0_).f18): len(_dafny.Map({self.f26: 876}))})
        r1 = len(d_1994_v66_)
        return r0, r1

    @property
    def f25(self):
        return self._f25

class C7(T0):
    def  __init__(self):
        self._f18: bool = False
        self.f24: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f24, f18):
        (self).f24 = f24
        (self)._f18 = f18

    def fm10(self, p0, p1, p2, globalState):
        return (self).f18

    def fm11(self, p0, p1, p2, globalState):
        return ((_dafny.Map({(self).f18: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifgrjv"))})) | (_dafny.Map({(self).f18: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpppgtxgx"))}))) | ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdc"))})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgxqth"))})))

    def m1(self, p0, p1, p2, globalState):
        (globalState).f12 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oukov"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdkrn")))) + (p2)
        d_1995_v0_: _dafny.Array
        nw354_ = _dafny.Array(int(0), 11)
        d_1995_v0_ = nw354_
        d_1996_v1_: _dafny.Map
        d_1996_v1_ = _dafny.Map({True: d_1995_v0_})
        d_1997_v2_: _dafny.MultiSet
        d_1997_v2_ = _dafny.MultiSet([d_1996_v1_])
        (globalState).f3 = ((d_1997_v2_)[(d_1996_v1_) | (d_1996_v1_)] if ((d_1996_v1_) | (d_1996_v1_)) in (d_1997_v2_) else self.f24)
        d_1998_v3_: str
        d_1998_v3_ = _dafny.CodePoint('a')
        d_1999_v4_: _dafny.Map
        d_1999_v4_ = _dafny.Map({(self).f18: d_1998_v3_})
        d_2000_v5_: _dafny.Map
        d_2000_v5_ = _dafny.Map({d_1999_v4_: (self).f18})
        d_2001_v6_: _dafny.Map
        d_2001_v6_ = _dafny.Map({p0: d_1998_v3_})
        d_2002_v7_: _dafny.Map
        d_2002_v7_ = _dafny.Map({d_2000_v5_: d_2001_v6_})
        d_2003_v9_: _dafny.Map
        d_2003_v9_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qv")): False})
        def iife154_():
            coll52_ = _dafny.Map()
            compr_52_: _dafny.Seq
            for compr_52_ in (d_2003_v9_).keys.Elements:
                d_2004_v8_: _dafny.Seq = compr_52_
                if (d_2004_v8_) in (d_2003_v9_):
                    coll52_[d_2004_v8_] = d_1998_v3_
            return _dafny.Map(coll52_)
        d_2002_v7_ = (d_2002_v7_).set(d_2000_v5_, iife154_()
        )
        d_2005_v10_: _dafny.Array
        def lambda84_(d_2006_i0_):
            return (self).f18

        init52_ = lambda84_
        nw355_ = _dafny.Array(None, 26)
        for i0_52_ in range(nw355_.length(0)):
            nw355_[i0_52_] = init52_(i0_52_)
        d_2005_v10_ = nw355_
        index344_ = default__.safeIndex(80, (d_2005_v10_).length(0))
        (d_2005_v10_)[index344_] = not(((self).f18) == (not(False)))
        index345_ = default__.safeIndex(80, (d_2005_v10_).length(0))
        (d_2005_v10_)[index345_] = not(((self).fm10((self).f18, (self).f18, (self).f18, globalState) if (self).f18 else (p1) <= (self.f24)))
        d_1998_v3_ = d_1998_v3_
        d_2007_v11_: _dafny.Map
        d_2007_v11_ = _dafny.Map({(self).f18: (self).f18})
        d_2008_v12_: _dafny.Map
        d_2008_v12_ = _dafny.Map({(self).f18: ((d_2007_v11_)[(self).f18] if ((self).f18) in (d_2007_v11_) else True)})
        d_2009_v13_: _dafny.Set
        d_2009_v13_ = _dafny.Set({len(d_2008_v12_)})
        d_2010_v14_: _dafny.Map
        d_2010_v14_ = _dafny.Map({(d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))]: p1})
        d_2011_v16_: _dafny.Seq
        d_2011_v16_ = _dafny.SeqWithoutIsStrInference([self.f24])
        d_2012_i1_: int
        d_2012_i1_ = 0
        with _dafny.label("14"):
            def iife155_():
                coll53_ = _dafny.Map()
                compr_53_: int
                for compr_53_ in (d_2011_v16_).Elements:
                    d_2029_v15_: int = compr_53_
                    if (d_2029_v15_) in (d_2011_v16_):
                        coll53_[(d_2029_v15_) - (p1)] = (self).f18
                return _dafny.Map(coll53_)
            while ((d_2009_v13_) == (d_2009_v13_)) == (((d_2010_v14_).set((d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))], len(iife155_()
            ))) != (d_2010_v14_)):
                with _dafny.c_label("14"):
                    if (d_2012_i1_) >= (100):
                        raise _dafny.Break("14")
                    d_2012_i1_ = (d_2012_i1_) + (1)
                    d_2013_v17_: _dafny.Seq
                    d_2013_v17_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_2014_v18_: _dafny.Array
                    nw356_ = _dafny.Array(None, 28)
                    nw356_[int(0)] = p2
                    nw356_[int(1)] = p2
                    nw356_[int(2)] = p0
                    nw356_[int(3)] = (p2) + (p0)
                    nw356_[int(4)] = p2
                    nw356_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1998_v3_ for d_2015_i2_ in range(default__.abs(470))])
                    nw356_[int(6)] = p0
                    nw356_[int(7)] = (p0) + (p2)
                    nw356_[int(8)] = p0
                    nw356_[int(9)] = p2
                    nw356_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmjimlo")))
                    nw356_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1998_v3_ for d_2016_i3_ in range(default__.abs(218))])
                    nw356_[int(12)] = p0
                    nw356_[int(13)] = p2
                    nw356_[int(14)] = p2
                    nw356_[int(15)] = (p2) + ((d_2013_v17_)[default__.safeIndex(len(p2), len(d_2013_v17_))])
                    nw356_[int(16)] = p0
                    nw356_[int(17)] = (default__.fm12(len(_dafny.SeqWithoutIsStrInference([d_1998_v3_ for d_2017_i4_ in range(default__.abs(853))])), (d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))], globalState)) + (p0)
                    nw356_[int(18)] = p0
                    nw356_[int(19)] = (p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))
                    nw356_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etpqxgf"))
                    nw356_[int(21)] = p0
                    nw356_[int(22)] = p2
                    nw356_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nuxgycs"))
                    nw356_[int(24)] = p2
                    nw356_[int(25)] = ((_dafny.SeqWithoutIsStrInference([d_1998_v3_ for d_2018_i5_ in range(default__.abs(-679))])) + (default__.fm12(self.f24, True, globalState))).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([d_1998_v3_ for d_2019_i5_ in range(default__.abs(-679))])) + (default__.fm12(self.f24, True, globalState)))), d_1998_v3_)
                    nw356_[int(26)] = p2
                    nw356_[int(27)] = _dafny.SeqWithoutIsStrInference([d_1998_v3_ for d_2020_i6_ in range(default__.abs(59))])
                    d_2014_v18_ = nw356_
                    d_2021_v19_: _dafny.Seq
                    d_2021_v19_ = _dafny.SeqWithoutIsStrInference([(self).f18, not((d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))])])
                    index346_ = default__.safeIndex(678, (d_2014_v18_).length(0))
                    (d_2014_v18_)[index346_] = ((p2).set(default__.safeIndex(p1, len(p2)), _dafny.CodePoint('s'))).set(default__.safeIndex(len(d_2021_v19_), len((p2).set(default__.safeIndex(p1, len(p2)), _dafny.CodePoint('s')))), d_1998_v3_)
                    index347_ = default__.safeIndex(678, (d_2014_v18_).length(0))
                    (d_2014_v18_)[index347_] = (p0).set(default__.safeIndex(p1, len(p0)), d_1998_v3_)
                    (globalState).f3 = (p1 if (self).f18 else len(_dafny.SeqWithoutIsStrInference([_dafny.Map({self.f24: 512}) for d_2022_i7_ in range(default__.abs(586))])))
                    d_2023_v20_: _dafny.MultiSet
                    d_2023_v20_ = _dafny.MultiSet([(p1) > (p1)])
                    d_2023_v20_ = d_2023_v20_
                    d_2024_v21_: _dafny.Map
                    d_2024_v21_ = _dafny.Map({341: default__.fm1(p1, (d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))], (d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))], globalState)})
                    d_2025_v22_: _dafny.Seq
                    d_2025_v22_ = _dafny.SeqWithoutIsStrInference([d_2024_v21_, d_2024_v21_])
                    d_2026_v23_: T1
                    nw357_ = C6()
                    nw357_.ctor__(_dafny.SeqWithoutIsStrInference([(d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))], (d_2005_v10_)[default__.safeIndex(80, (d_2005_v10_).length(0))]]), self.f24, not(((d_2014_v18_)[default__.safeIndex(678, (d_2014_v18_).length(0))]) == (p2)), p2, (d_2025_v22_) + (d_2025_v22_))
                    d_2026_v23_ = nw357_
                    d_2027_v24_: _dafny.Map
                    d_2027_v24_ = _dafny.Map({p2: d_2026_v23_})
                    d_2028_v25_: _dafny.Map
                    d_2028_v25_ = _dafny.Map({_dafny.CodePoint('y'): d_2026_v23_})
                    d_2026_v23_ = ((d_2027_v24_)[(d_2026_v23_).f19] if ((d_2026_v23_).f19) in (d_2027_v24_) else ((d_2028_v25_)[d_1998_v3_] if (d_1998_v3_) in (d_2028_v25_) else d_2026_v23_))
                    pass
            pass

    def m5(self, p0, p1, p2, p3, globalState):
        d_2030_v0_: D17
        d_2030_v0_ = D17_DC42((self).f18)
        d_2030_v0_ = d_2030_v0_
        (globalState).f2 = not((self).f18)
        hi11_ = p3
        for d_2031_i0_ in range(default__.safeDivisionInt(self.f24, 599), hi11_):
            d_2032_v1_: _dafny.Seq
            d_2032_v1_ = _dafny.SeqWithoutIsStrInference([len(default__.fm43((self).f18, not((self).f18), globalState)), default__.safeModuloInt(p1, p3)])
            (globalState).f10 = d_2032_v1_
            d_2033_v2_: _dafny.Seq
            d_2033_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqugt"))
            d_2034_v3_: _dafny.Map
            d_2034_v3_ = _dafny.Map({d_2033_v2_: (self).f18})
            d_2035_v4_: C0
            nw358_ = C0()
            nw358_.ctor__(((p2)[p1] if (p1) in (p2) else p1), (d_2032_v1_)[default__.safeIndex(len(d_2034_v3_), len(d_2032_v1_))])
            d_2035_v4_ = nw358_
            d_2036_v5_: D11
            d_2036_v5_ = D11_DC29(D11_DC28((self).f18))
            d_2036_v5_ = d_2036_v5_
            d_2037_v6_: _dafny.Array
            def lambda85_(d_2038_i1_):
                return (_dafny.Map({(self).f18: not(not((self).f18))})) | (_dafny.Map({(self).f18: (self).f18}))

            init53_ = lambda85_
            nw359_ = _dafny.Array(None, 26)
            for i0_53_ in range(nw359_.length(0)):
                nw359_[i0_53_] = init53_(i0_53_)
            d_2037_v6_ = nw359_
            d_2039_v7_: _dafny.Map
            d_2039_v7_ = _dafny.Map({(self).f18: not((self).f18)})
            index348_ = default__.safeIndex(987, (d_2037_v6_).length(0))
            (d_2037_v6_)[index348_] = d_2039_v7_
            index349_ = default__.safeIndex(987, (d_2037_v6_).length(0))
            (d_2037_v6_)[index349_] = ((d_2039_v7_) | (_dafny.Map({(self).f18: True})) if (self).f18 else d_2039_v7_)
        d_2040_v8_: _dafny.Array
        nw360_ = _dafny.Array(int(0), 11)
        d_2040_v8_ = nw360_
        index350_ = default__.safeIndex(914, (d_2040_v8_).length(0))
        (d_2040_v8_)[index350_] = (p3) * (self.f24)
        d_2041_v9_: _dafny.Array
        nw361_ = _dafny.Array(None, 6)
        nw361_[int(0)] = (p1) >= ((0) - (self.f24))
        nw361_[int(1)] = (self).f18
        nw361_[int(2)] = (self).f18
        nw361_[int(3)] = True
        nw361_[int(4)] = (self).f18
        nw361_[int(5)] = (self).f18
        d_2041_v9_ = nw361_
        index351_ = default__.safeIndex(914, (d_2040_v8_).length(0))
        rhs364_ = p3
        rhs365_ = d_2041_v9_
        rhs366_ = (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ic"))), self.f24)) + (p3)
        lhs313_ = d_2040_v8_
        lhs314_ = default__.safeIndex(914, (d_2040_v8_).length(0))
        lhs315_ = globalState
        lhs313_[lhs314_] = rhs364_
        d_2041_v9_ = rhs365_
        lhs315_.f3 = rhs366_
        d_2042_v10_: str
        d_2042_v10_ = _dafny.CodePoint('e')
        d_2043_v11_: _dafny.Seq
        d_2043_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uermrgdr"))
        d_2044_v12_: _dafny.Seq
        d_2044_v12_ = _dafny.SeqWithoutIsStrInference([(self).f18])
        d_2045_v13_: _dafny.Seq
        d_2045_v13_ = _dafny.SeqWithoutIsStrInference([len(d_2044_v12_), len(_dafny.Set({(self).f18, True}))])
        d_2046_v14_: C2
        nw362_ = C2()
        nw362_.ctor__(p3, d_2045_v13_, (self).f18)
        d_2046_v14_ = nw362_
        d_2047_v15_: D12
        d_2047_v15_ = D12_DC31()
        d_2048_v16_: _dafny.Map
        d_2048_v16_ = _dafny.Map({d_2046_v14_: d_2047_v15_})
        d_2049_v17_: _dafny.Map
        d_2049_v17_ = _dafny.Map({len(d_2048_v16_): self.f24})
        d_2050_v18_: _dafny.Seq
        d_2050_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p3: p1}), d_2049_v17_, d_2049_v17_, d_2049_v17_])
        d_2051_v19_: C4
        nw363_ = C4()
        nw363_.ctor__(d_2042_v10_, ((d_2040_v8_)[default__.safeIndex(914, (d_2040_v8_).length(0))]) + (p3), (self).f18, d_2043_v11_, d_2050_v18_)
        d_2051_v19_ = nw363_
        d_2052_v20_: _dafny.MultiSet
        d_2052_v20_ = _dafny.MultiSet([(self).f18, (self).f18])
        if (d_2044_v12_)[default__.safeIndex((self.f24) - ((d_2052_v20_).cardinality), len(d_2044_v12_))]:
            d_2053_v21_: C0
            nw364_ = C0()
            nw364_.ctor__(74, (d_2051_v19_).fm18(globalState))
            d_2053_v21_ = nw364_
            d_2042_v10_ = default__.fm46(globalState)
            d_2054_v22_: D17
            d_2054_v22_ = D17_DC44(d_2041_v9_, self.f24)
            d_2055_v23_: _dafny.Set
            d_2055_v23_ = _dafny.Set({(d_2040_v8_)[default__.safeIndex(914, (d_2040_v8_).length(0))]})
            pat_let_tv29_ = d_2041_v9_
            pat_let_tv30_ = p1
            pat_let_tv31_ = d_2055_v23_
            d_2056_v24_: _dafny.Array
            nw365_ = _dafny.Array(None, 20)
            def iife156_(_pat_let51_0):
                def iife157_(d_2057_dt__update__tmp_h0_):
                    def iife158_(_pat_let52_0):
                        def iife159_(d_2058_dt__update_hcf77_h0_):
                            return D17_DC44(d_2058_dt__update_hcf77_h0_, (d_2057_dt__update__tmp_h0_).cf78)
                        return iife159_(_pat_let52_0)
                    return iife158_(pat_let_tv29_)
                return iife157_(_pat_let51_0)
            nw365_[int(0)] = iife156_(d_2054_v22_)
            nw365_[int(1)] = d_2054_v22_
            nw365_[int(2)] = d_2054_v22_
            nw365_[int(3)] = d_2054_v22_
            nw365_[int(4)] = d_2054_v22_
            nw365_[int(5)] = d_2054_v22_
            def iife160_(_pat_let53_0):
                def iife161_(d_2059_dt__update__tmp_h1_):
                    def iife162_(_pat_let54_0):
                        def iife163_(d_2060_dt__update_hcf78_h0_):
                            return D17_DC44((d_2059_dt__update__tmp_h1_).cf77, d_2060_dt__update_hcf78_h0_)
                        return iife163_(_pat_let54_0)
                    return iife162_(pat_let_tv30_)
                return iife161_(_pat_let53_0)
            nw365_[int(6)] = (iife160_(d_2054_v22_) if (self).f18 else d_2054_v22_)
            nw365_[int(7)] = d_2054_v22_
            nw365_[int(8)] = d_2054_v22_
            nw365_[int(9)] = d_2054_v22_
            def iife164_(_pat_let55_0):
                def iife165_(d_2061_dt__update__tmp_h2_):
                    def iife166_(_pat_let56_0):
                        def iife167_(d_2062_dt__update_hcf78_h1_):
                            return D17_DC44((d_2061_dt__update__tmp_h2_).cf77, d_2062_dt__update_hcf78_h1_)
                        return iife167_(_pat_let56_0)
                    return iife166_(len(pat_let_tv31_))
                return iife165_(_pat_let55_0)
            nw365_[int(10)] = iife164_(D17_DC44(d_2041_v9_, default__.fm1(d_2046_v14_.f33, (self).f18, True, globalState)))
            nw365_[int(11)] = d_2054_v22_
            nw365_[int(12)] = d_2054_v22_
            nw365_[int(13)] = d_2054_v22_
            nw365_[int(14)] = d_2054_v22_
            nw365_[int(15)] = d_2054_v22_
            nw365_[int(16)] = D17_DC44(d_2041_v9_, d_2046_v14_.f33)
            nw365_[int(17)] = d_2054_v22_
            nw365_[int(18)] = d_2054_v22_
            nw365_[int(19)] = d_2054_v22_
            d_2056_v24_ = nw365_
            index352_ = default__.safeIndex(903, (d_2056_v24_).length(0))
            (d_2056_v24_)[index352_] = d_2054_v22_
            d_2063_v25_: C6
            nw366_ = C6()
            nw366_.ctor__((d_2044_v12_) + (_dafny.SeqWithoutIsStrInference([(self).f18, False, (self).f18])), p1, (self).f18, d_2043_v11_, (d_2050_v18_).set(default__.safeIndex(self.f24, len(d_2050_v18_)), d_2049_v17_))
            d_2063_v25_ = nw366_
            d_2064_v26_: D13
            d_2064_v26_ = D13_DC33((self).f18, (d_2051_v19_).f30, not((self).f18), d_2046_v14_.f33)
            d_2065_v27_: D3
            d_2065_v27_ = D3_DC11((self).f18, (self).f18)
            index353_ = default__.safeIndex(903, (d_2056_v24_).length(0))
            rhs367_ = ((self).f18) == ((d_2051_v19_).fm2((self).f18, (self).f18, (self).f18, globalState))
            rhs368_ = not((d_2064_v26_).cf57)
            rhs369_ = (d_2043_v11_).set(default__.safeIndex(len(d_2049_v17_), len(d_2043_v11_)), d_2042_v10_)
            rhs370_ = D17_DC44(d_2041_v9_, ((d_2053_v21_).f31 if not((d_2065_v27_).cf24) else self.f24))
            rhs371_ = d_2063_v25_
            lhs316_ = globalState
            lhs317_ = globalState
            lhs318_ = globalState
            lhs319_ = d_2056_v24_
            lhs320_ = default__.safeIndex(903, (d_2056_v24_).length(0))
            lhs316_.f2 = rhs367_
            lhs317_.f2 = rhs368_
            lhs318_.f12 = rhs369_
            lhs319_[lhs320_] = rhs370_
            d_2063_v25_ = rhs371_
            d_2066_v28_: _dafny.Set
            d_2066_v28_ = _dafny.Set({(d_2051_v19_).f29, (d_2051_v19_).f29})
            d_2067_v29_: _dafny.Map
            d_2067_v29_ = _dafny.Map({(d_2051_v19_).fm3(72, globalState): len(d_2066_v28_)})
            d_2068_v30_: _dafny.Map
            d_2068_v30_ = _dafny.Map({((d_2067_v29_)[False] if (False) in (d_2067_v29_) else (d_2053_v21_).f32): (self).f18})
            d_2068_v30_ = (d_2068_v30_).set(len(d_2043_v11_), not((self).f18))
            def iife168_():
                coll54_ = _dafny.Map()
                compr_54_: int
                for compr_54_ in (d_2045_v13_).Elements:
                    d_2069_v31_: int = compr_54_
                    if (d_2069_v31_) in (d_2045_v13_):
                        coll54_[(d_2069_v31_) - ((d_2053_v21_).f31)] = (self).f18
                return _dafny.Map(coll54_)
            d_2068_v30_ = (d_2068_v30_) | (iife168_()
            )
        elif True:
            (globalState).f2 = (self).f18
            (globalState).f2 = not((self).f18)
            d_2070_v32_: _dafny.Array
            nw367_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
            d_2070_v32_ = nw367_
            index354_ = default__.safeIndex(283, (d_2070_v32_).length(0))
            (d_2070_v32_)[index354_] = (D13_DC34(d_2041_v9_, d_2043_v11_, (self).f18, (self).f18)).cf60
            d_2071_v33_: _dafny.Seq
            d_2071_v33_ = _dafny.SeqWithoutIsStrInference([d_2043_v11_])
            d_2072_v34_: _dafny.Map
            d_2072_v34_ = _dafny.Map({(self).f18: (d_2071_v33_)[default__.safeIndex(d_2046_v14_.f33, len(d_2071_v33_))]})
            index355_ = default__.safeIndex(283, (d_2070_v32_).length(0))
            (d_2070_v32_)[index355_] = ((d_2072_v34_)[((d_2040_v8_)[default__.safeIndex(914, (d_2040_v8_).length(0))]) < (len((d_2046_v14_).f34))] if (((d_2040_v8_)[default__.safeIndex(914, (d_2040_v8_).length(0))]) < (len((d_2046_v14_).f34))) in (d_2072_v34_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnxxapdgd")))
            d_2073_v35_: _dafny.Array
            def lambda86_(d_2074_p3_, d_2075_v17_):
                def lambda87_(d_2076_i2_):
                    return (_dafny.Map({d_2074_p3_: 86})) | (d_2075_v17_)

                return lambda87_

            init54_ = lambda86_(p3, d_2049_v17_)
            nw368_ = _dafny.Array(None, 12)
            for i0_54_ in range(nw368_.length(0)):
                nw368_[i0_54_] = init54_(i0_54_)
            d_2073_v35_ = nw368_
            index356_ = default__.safeIndex(951, (d_2073_v35_).length(0))
            (d_2073_v35_)[index356_] = d_2049_v17_
            d_2077_v36_: _dafny.Seq
            d_2077_v36_ = _dafny.SeqWithoutIsStrInference([d_2052_v20_])
            index357_ = default__.safeIndex(951, (d_2073_v35_).length(0))
            def iife169_():
                coll55_ = _dafny.Map()
                compr_55_: _dafny.Seq
                for compr_55_ in (d_2071_v33_).Elements:
                    d_2078_v37_: _dafny.Seq = compr_55_
                    if (d_2078_v37_) in (d_2071_v33_):
                        coll55_[d_2078_v37_] = (self).f18
                return _dafny.Map(coll55_)
            (d_2073_v35_)[index357_] = default__.fm30(d_2077_v36_, iife169_()
            , globalState)
            d_2079_v38_: _dafny.Set
            d_2079_v38_ = _dafny.Set({(self).f18})
            d_2080_v39_: D10
            d_2080_v39_ = D10_DC24(d_2079_v38_)
            d_2081_v40_: D10
            d_2081_v40_ = D10_DC26(D10_DC26(d_2080_v39_))
            source18_ = d_2081_v40_
            if source18_.is_DC25:
                d_2082___mcc_h0_ = source18_.cf48
                d_2083_cf48_ = d_2082___mcc_h0_
                d_2084_v41_: _dafny.Seq
                d_2084_v41_ = _dafny.SeqWithoutIsStrInference([default__.fm44(d_2046_v14_.f33, d_2083_cf48_, default__.fm46(globalState), d_2046_v14_.f33, globalState)])
                d_2085_v42_: _dafny.Map
                d_2085_v42_ = _dafny.Map({(self).f18: (d_2084_v41_)[default__.safeIndex((d_2051_v19_).f30, len(d_2084_v41_))]})
                d_2086_v43_: _dafny.Map
                d_2086_v43_ = _dafny.Map({d_2041_v9_: not(not(True))})
                d_2044_v12_ = (((d_2085_v42_)[d_2083_cf48_] if (d_2083_cf48_) in (d_2085_v42_) else (d_2044_v12_) + (d_2044_v12_))).set(default__.safeIndex(len(d_2086_v43_), len(((d_2085_v42_)[d_2083_cf48_] if (d_2083_cf48_) in (d_2085_v42_) else (d_2044_v12_) + (d_2044_v12_)))), d_2083_cf48_)
                d_2087_v44_: _dafny.Map
                d_2087_v44_ = _dafny.Map({d_2046_v14_.f33: d_2083_cf48_})
                d_2087_v44_ = (d_2087_v44_).set((d_2051_v19_).f30, ((d_2044_v12_)[default__.safeIndex(len(d_2045_v13_), len(d_2044_v12_))]) or ((self).f18))
                (globalState).f2 = ((self).f18 if (self).f18 else (self).f18)
                (globalState).f3 = (d_2040_v8_)[default__.safeIndex(914, (d_2040_v8_).length(0))]
            elif source18_.is_DC24:
                d_2088___mcc_h1_ = source18_.cf47
                d_2089_cf47_ = d_2088___mcc_h1_
                d_2090_v45_: C4
                nw369_ = C4()
                nw369_.ctor__((d_2051_v19_).f29, (d_2051_v19_).f30, (self).f18, (d_2070_v32_)[default__.safeIndex(283, (d_2070_v32_).length(0))], d_2050_v18_)
                d_2090_v45_ = nw369_
                index358_ = default__.safeIndex(914, (d_2040_v8_).length(0))
                rhs372_ = (self).f18
                rhs373_ = len(d_2043_v11_)
                rhs374_ = (False) == ((self).f18)
                rhs375_ = ((d_2051_v19_).f30) + (len((d_2070_v32_)[default__.safeIndex(283, (d_2070_v32_).length(0))]))
                lhs321_ = globalState
                lhs322_ = d_2040_v8_
                lhs323_ = default__.safeIndex(914, (d_2040_v8_).length(0))
                lhs324_ = globalState
                lhs325_ = d_2046_v14_
                lhs321_.f2 = rhs372_
                lhs322_[lhs323_] = rhs373_
                lhs324_.f2 = rhs374_
                lhs325_.f33 = rhs375_
                (globalState).f0 = (265) + (d_2046_v14_.f33)
                index359_ = default__.safeIndex(951, (d_2073_v35_).length(0))
                (d_2073_v35_)[index359_] = ((d_2073_v35_)[default__.safeIndex(951, (d_2073_v35_).length(0))]).set(((d_2090_v45_).f30) + (p1), p1)
            elif True:
                d_2091___mcc_h2_ = source18_.cf49
                d_2092_cf49_ = d_2091___mcc_h2_
                index360_ = default__.safeIndex(914, (d_2040_v8_).length(0))
                (d_2040_v8_)[index360_] = default__.safeDivisionInt(default__.fm1((d_2051_v19_).f30, (self).f18, (self).f18, globalState), (d_2046_v14_.f33) - (42))
                d_2093_v46_: _dafny.Seq
                d_2093_v46_ = _dafny.SeqWithoutIsStrInference([(d_2046_v14_).f34, (d_2046_v14_).f34])
                (globalState).f2 = (((d_2046_v14_).f34) + (d_2045_v13_)) in (d_2093_v46_)
                (globalState).f0 = (((d_2046_v14_).f34)[default__.safeIndex(d_2046_v14_.f33, len((d_2046_v14_).f34))]) * (d_2046_v14_.f33)
                (globalState).f2 = (self).f18


class C8(T0):
    def  __init__(self):
        self._f18: bool = False
        self.f23: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f23, f18):
        (self).f23 = f23
        (self)._f18 = f18

    def m1(self, p0, p1, p2, globalState):
        d_2094_v0_: _dafny.Set
        d_2094_v0_ = _dafny.Set({default__.fm1(p1, (self).f18, (self).f18, globalState)})
        hi12_ = (486) + (len(d_2094_v0_))
        for d_2095_i0_ in range(p1, hi12_):
            d_2096_v1_: str
            d_2096_v1_ = _dafny.CodePoint('w')
            d_2097_v2_: _dafny.MultiSet
            d_2097_v2_ = _dafny.MultiSet([d_2096_v1_, d_2096_v1_, d_2096_v1_])
            d_2097_v2_ = (d_2097_v2_).intersection((d_2097_v2_).intersection((d_2097_v2_).set(d_2096_v1_, default__.abs(d_2095_i0_))))
            (globalState).f0 = p1
            d_2098_v3_: _dafny.Map
            d_2098_v3_ = _dafny.Map({(self).f18: (self).f18})
            d_2099_v4_: _dafny.Seq
            d_2099_v4_ = _dafny.SeqWithoutIsStrInference([self.f23, len(d_2098_v3_)])
            d_2100_v5_: _dafny.MultiSet
            d_2100_v5_ = _dafny.MultiSet([p1])
            if not(not(default__.fm6((d_2099_v4_) + (d_2099_v4_), False, (d_2100_v5_) - (d_2100_v5_), p1, globalState))):
                d_2101_v6_: _dafny.Array
                nw370_ = _dafny.Array(int(0), 17)
                d_2101_v6_ = nw370_
                d_2102_v7_: _dafny.Seq
                d_2102_v7_ = _dafny.SeqWithoutIsStrInference([d_2101_v6_])
                d_2103_v8_: _dafny.Seq
                d_2103_v8_ = _dafny.SeqWithoutIsStrInference([d_2102_v7_, d_2102_v7_, _dafny.SeqWithoutIsStrInference([d_2101_v6_]), d_2102_v7_, d_2102_v7_])
                (globalState).f2 = (((d_2103_v8_)[default__.safeIndex(len(p0), len(d_2103_v8_))]) + (d_2102_v7_)) < (d_2102_v7_)
                d_2104_v9_: _dafny.Array
                nw371_ = _dafny.Array(False, 17)
                d_2104_v9_ = nw371_
                d_2105_v10_: _dafny.Map
                d_2105_v10_ = _dafny.Map({d_2104_v9_: p1})
                (globalState).f2 = (((d_2105_v10_)[d_2104_v9_] if (d_2104_v9_) in (d_2105_v10_) else 9)) >= ((d_2095_i0_) + (d_2095_i0_))
                d_2106_v11_: _dafny.Map
                d_2106_v11_ = _dafny.Map({default__.fm1(d_2095_i0_, (self).f18, (self).f18, globalState): (self).f18})
                d_2107_v12_: _dafny.Map
                d_2107_v12_ = _dafny.Map({self.f23: (d_2106_v11_).set(670, (self).f18)})
                d_2108_v13_: D2
                d_2108_v13_ = D2_DC6(_dafny.Map({(self).f18: default__.fm1(default__.fm1(p1, (self).f18, (self).f18, globalState), default__.fm6(d_2099_v4_, (self).f18, d_2100_v5_, d_2095_i0_, globalState), (self).f18, globalState)}))
                d_2109_v14_: _dafny.Map
                d_2109_v14_ = _dafny.Map({(self).f18: (_dafny.MultiSet([(self).f18, (self).f18, False])).cardinality})
                d_2110_v15_: _dafny.Set
                d_2110_v15_ = _dafny.Set({not((self).f18)})
                d_2111_v16_: _dafny.Map
                d_2111_v16_ = _dafny.Map({p1: d_2098_v3_})
                pat_let_tv32_ = d_2109_v14_
                pat_let_tv33_ = d_2109_v14_
                pat_let_tv34_ = d_2109_v14_
                d_2112_v17_: _dafny.Array
                nw372_ = _dafny.Array(None, 25)
                nw372_[int(0)] = D2_DC6(default__.fm7(len(d_2107_v12_), (self).f18, globalState))
                nw372_[int(1)] = d_2108_v13_
                nw372_[int(2)] = d_2108_v13_
                nw372_[int(3)] = d_2108_v13_
                nw372_[int(4)] = d_2108_v13_
                nw372_[int(5)] = D2_DC6(d_2109_v14_)
                def iife170_(_pat_let57_0):
                    def iife171_(d_2113_dt__update__tmp_h0_):
                        def iife172_(_pat_let58_0):
                            def iife173_(d_2114_dt__update_hcf14_h0_):
                                return D2_DC6(d_2114_dt__update_hcf14_h0_)
                            return iife173_(_pat_let58_0)
                        return iife172_(pat_let_tv32_)
                    return iife171_(_pat_let57_0)
                nw372_[int(6)] = iife170_(d_2108_v13_)
                nw372_[int(7)] = d_2108_v13_
                nw372_[int(8)] = D2_DC6(d_2109_v14_)
                nw372_[int(9)] = default__.fm8(d_2110_v15_, (self).f18, (self).f18, d_2096_v1_, globalState)
                nw372_[int(10)] = d_2108_v13_
                def iife175_(_pat_let60_0):
                    def iife176_(d_2115_dt__update__tmp_h1_):
                        def iife177_(_pat_let61_0):
                            def iife178_(d_2116_dt__update_hcf14_h1_):
                                return D2_DC6(d_2116_dt__update_hcf14_h1_)
                            return iife178_(_pat_let61_0)
                        return iife177_(pat_let_tv33_)
                    return iife176_(_pat_let60_0)
                def iife174_(_pat_let59_0):
                    def iife179_(d_2117_dt__update__tmp_h2_):
                        def iife180_(_pat_let62_0):
                            def iife181_(d_2118_dt__update_hcf14_h2_):
                                return D2_DC6(d_2118_dt__update_hcf14_h2_)
                            return iife181_(_pat_let62_0)
                        return iife180_(pat_let_tv34_)
                    return iife179_(_pat_let59_0)
                nw372_[int(11)] = iife174_(iife175_(d_2108_v13_))
                nw372_[int(12)] = d_2108_v13_
                nw372_[int(13)] = d_2108_v13_
                nw372_[int(14)] = d_2108_v13_
                nw372_[int(15)] = d_2108_v13_
                nw372_[int(16)] = d_2108_v13_
                nw372_[int(17)] = (default__.fm8(d_2110_v15_, (self).f18, not((self).f18), d_2096_v1_, globalState) if default__.fm6(d_2099_v4_, not((self).f18), _dafny.MultiSet([len(d_2111_v16_), -262]), len(d_2110_v15_), globalState) else d_2108_v13_)
                nw372_[int(18)] = D2_DC6(default__.fm7(56, False, globalState))
                nw372_[int(19)] = d_2108_v13_
                nw372_[int(20)] = D2_DC6(_dafny.Map({(self).f18: len(d_2099_v4_)}))
                nw372_[int(21)] = default__.fm8(d_2110_v15_, (self).f18, (self).f18, d_2096_v1_, globalState)
                nw372_[int(22)] = d_2108_v13_
                nw372_[int(23)] = d_2108_v13_
                nw372_[int(24)] = D2_DC6(d_2109_v14_)
                d_2112_v17_ = nw372_
                index361_ = default__.safeIndex(459, (d_2112_v17_).length(0))
                (d_2112_v17_)[index361_] = d_2108_v13_
                index362_ = default__.safeIndex(459, (d_2112_v17_).length(0))
                (d_2112_v17_)[index362_] = d_2108_v13_
                (globalState).f2 = ((default__.fm9((self).f18, globalState)) | (_dafny.Set({(self).f18, (self).f18, False, (self).f18}))).isdisjoint(default__.fm9((self).f18, globalState))
                d_2119_v18_: _dafny.Array
                nw373_ = _dafny.Array(None, 11)
                nw373_[int(0)] = p0
                nw373_[int(1)] = p0
                nw373_[int(2)] = (p0) + (p2)
                nw373_[int(3)] = p0
                nw373_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                nw373_[int(5)] = _dafny.SeqWithoutIsStrInference([d_2096_v1_ for d_2120_i1_ in range(default__.abs(747))])
                nw373_[int(6)] = p0
                nw373_[int(7)] = p0
                nw373_[int(8)] = (D0_DC1((self).f18, len((p2).set(default__.safeIndex(self.f23, len(p2)), d_2096_v1_)), p2, (self).f18, self.f23)).cf3
                nw373_[int(9)] = p0
                nw373_[int(10)] = p2
                d_2119_v18_ = nw373_
                index363_ = default__.safeIndex(818, (d_2119_v18_).length(0))
                (d_2119_v18_)[index363_] = (p2) + (p2)
                index364_ = default__.safeIndex(818, (d_2119_v18_).length(0))
                (d_2119_v18_)[index364_] = (p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msaxakyd")))
            elif True:
                d_2121_v19_: _dafny.Set
                d_2121_v19_ = _dafny.Set({True, ((self).f18) and ((self).f18), (self).f18, (d_2096_v1_) not in (p0)})
                d_2121_v19_ = ((_dafny.Set({not((self).f18), default__.fm6(_dafny.SeqWithoutIsStrInference([self.f23, self.f23]), (self).f18, d_2100_v5_, 800, globalState), (self).f18}) if not((self).f18) else _dafny.Set({(self).f18}))) | (d_2121_v19_)
                (globalState).f2 = (self).f18
                (globalState).f3 = (0) - ((417) - ((self.f23) * (p1)))
                d_2122_v20_: _dafny.Map
                d_2122_v20_ = _dafny.Map({self.f23: self.f23})
                d_2123_v21_: T0
                nw374_ = C4()
                nw374_.ctor__(d_2096_v1_, 611, (self).f18, _dafny.SeqWithoutIsStrInference([d_2096_v1_ for d_2124_i2_ in range(default__.abs(842))]), _dafny.SeqWithoutIsStrInference([d_2122_v20_, d_2122_v20_]))
                d_2123_v21_ = nw374_
                d_2125_v22_: D1
                d_2125_v22_ = D1_DC2(d_2123_v21_)
                d_2126_v23_: _dafny.Map
                d_2126_v23_ = _dafny.Map({d_2125_v22_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdy")))})
                d_2127_v24_: _dafny.Map
                d_2127_v24_ = _dafny.Map({(self).f18: p2})
                d_2128_v25_: _dafny.Map
                d_2128_v25_ = _dafny.Map({((d_2126_v23_)[D1_DC2(d_2123_v21_)] if (D1_DC2(d_2123_v21_)) in (d_2126_v23_) else self.f23): len(d_2127_v24_)})
                d_2129_v26_: _dafny.Map
                d_2129_v26_ = _dafny.Map({(d_2123_v21_).f18: 442})
                d_2130_v27_: _dafny.Seq
                d_2130_v27_ = _dafny.SeqWithoutIsStrInference([(self).f18, (d_2123_v21_).f18, (d_2123_v21_).f18, (d_2123_v21_).f18, (d_2123_v21_).f18])
                d_2122_v20_ = ((d_2122_v20_) | (_dafny.Map({len(d_2129_v26_): len(d_2130_v27_)}))) | ((d_2122_v20_) | (d_2122_v20_))
                d_2128_v25_ = (d_2128_v25_).set(929, d_2095_i0_)
            d_2131_v28_: D19
            d_2131_v28_ = D19_DC50((self).f18)
            d_2132_v29_: D2
            d_2132_v29_ = D2_DC8((d_2131_v28_).cf84, (self).f18, True)
            d_2133_v30_: _dafny.Map
            d_2133_v30_ = _dafny.Map({d_2132_v29_: (self).f18})
            d_2134_v31_: _dafny.Array
            def lambda88_(d_2135_i3_):
                return (d_2135_i3_) * (self.f23)

            init55_ = lambda88_
            nw375_ = _dafny.Array(None, 8)
            for i0_55_ in range(nw375_.length(0)):
                nw375_[i0_55_] = init55_(i0_55_)
            d_2134_v31_ = nw375_
            d_2136_v32_: _dafny.Seq
            d_2136_v32_ = _dafny.SeqWithoutIsStrInference([d_2134_v31_, d_2134_v31_])
            d_2137_v33_: _dafny.MultiSet
            d_2137_v33_ = _dafny.MultiSet([default__.fm6(d_2099_v4_, (self).f18, d_2100_v5_, self.f23, globalState), (self).f18])
            rhs376_ = d_2133_v30_
            rhs377_ = (_dafny.SeqWithoutIsStrInference([d_2134_v31_, d_2134_v31_])) + (d_2136_v32_)
            rhs378_ = default__.fm1(d_2095_i0_, (d_2137_v33_).isdisjoint(d_2137_v33_), (self).f18, globalState)
            lhs326_ = globalState
            d_2133_v30_ = rhs376_
            d_2136_v32_ = rhs377_
            lhs326_.f3 = rhs378_
        d_2138_v34_: C7
        nw376_ = C7()
        nw376_.ctor__(self.f23, (self).f18)
        d_2138_v34_ = nw376_
        if (self).f18:
            d_2139_v35_: C1
            nw377_ = C1()
            nw377_.ctor__((self).f18)
            d_2139_v35_ = nw377_
            d_2140_v36_: _dafny.Seq
            d_2140_v36_ = _dafny.SeqWithoutIsStrInference([d_2138_v34_.f24, len(p2)])
            (globalState).f2 = (self.f23) != (default__.fm1(len(d_2140_v36_), (self).f18, (self).f18, globalState))
            if (self).f18:
                d_2141_v37_: _dafny.Seq
                d_2141_v37_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2142_v38_: _dafny.MultiSet
                d_2142_v38_ = _dafny.MultiSet([(self).f18])
                d_2143_v39_: _dafny.Map
                d_2143_v39_ = _dafny.Map({self.f23: (d_2141_v37_)[default__.safeIndex((0) - ((d_2142_v38_).cardinality), len(d_2141_v37_))]})
                d_2143_v39_ = (d_2143_v39_).set((p1) - (d_2138_v34_.f24), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bko")))
                d_2144_v40_: _dafny.Map
                d_2144_v40_ = _dafny.Map({d_2140_v36_: (self).f18})
                d_2144_v40_ = (d_2144_v40_).set((d_2140_v36_) + (d_2140_v36_), (self).f18)
                (globalState).f2 = ((self).f18) or (not (not((d_2138_v34_).fm10(True, (self).f18, (self).f18, globalState))) or (False))
                (globalState).f10 = d_2140_v36_
                (globalState).f2 = (self).f18
            elif True:
                d_2145_v41_: _dafny.MultiSet
                d_2145_v41_ = _dafny.MultiSet([p1])
                d_2145_v41_ = d_2145_v41_
                d_2146_v42_: _dafny.MultiSet
                d_2146_v42_ = _dafny.MultiSet([(self).f18])
                d_2147_v43_: _dafny.Seq
                d_2147_v43_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                d_2148_v44_: D12
                d_2148_v44_ = D12_DC30(d_2147_v43_)
                d_2149_v45_: _dafny.MultiSet
                d_2149_v45_ = d_2146_v42_
                d_2150_v46_: _dafny.Map
                d_2150_v46_ = _dafny.Map({d_2145_v41_: d_2146_v42_})
                d_2151_v47_: _dafny.Array
                nw378_ = _dafny.Array(None, 20)
                nw378_[int(0)] = d_2146_v42_
                nw378_[int(1)] = (_dafny.MultiSet([(self).f18, (self).f18])).set(not((self).f18), default__.abs(self.f23))
                nw378_[int(2)] = (_dafny.MultiSet((d_2148_v44_).cf53)).intersection(d_2146_v42_)
                nw378_[int(3)] = (d_2149_v45_)
                nw378_[int(4)] = d_2146_v42_
                nw378_[int(5)] = (d_2146_v42_).set(False, default__.abs(d_2138_v34_.f24))
                nw378_[int(6)] = d_2146_v42_
                nw378_[int(7)] = (d_2146_v42_).intersection(d_2146_v42_)
                nw378_[int(8)] = d_2146_v42_
                nw378_[int(9)] = ((d_2150_v46_)[_dafny.MultiSet(d_2140_v36_)] if (_dafny.MultiSet(d_2140_v36_)) in (d_2150_v46_) else d_2146_v42_)
                nw378_[int(10)] = (d_2146_v42_) | (_dafny.MultiSet([(self).f18, False, not((self).f18)]))
                nw378_[int(11)] = default__.fm32(globalState)
                nw378_[int(12)] = d_2146_v42_
                nw378_[int(13)] = d_2146_v42_
                nw378_[int(14)] = _dafny.MultiSet([(self).f18])
                nw378_[int(15)] = d_2146_v42_
                nw378_[int(16)] = d_2146_v42_
                nw378_[int(17)] = _dafny.MultiSet(d_2147_v43_)
                nw378_[int(18)] = (d_2146_v42_).set((self).f18, default__.abs(((D16_DC38(d_2145_v41_)).cf66).cardinality))
                nw378_[int(19)] = (_dafny.MultiSet(d_2147_v43_)) | (d_2146_v42_)
                d_2151_v47_ = nw378_
                index365_ = default__.safeIndex(667, (d_2151_v47_).length(0))
                (d_2151_v47_)[index365_] = _dafny.MultiSet([(self).f18, False, (self).f18])
                index366_ = default__.safeIndex(667, (d_2151_v47_).length(0))
                (d_2151_v47_)[index366_] = _dafny.MultiSet([(not((self).f18) if (self).f18 else (self).f18)])
                d_2152_v48_: str
                d_2152_v48_ = _dafny.CodePoint('s')
                d_2153_v49_: D3
                d_2153_v49_ = D3_DC13(d_2138_v34_.f24, p2, self.f23, default__.fm1(self.f23, (self).f18, (self).f18, globalState), d_2152_v48_)
                d_2154_v50_: _dafny.Map
                d_2154_v50_ = _dafny.Map({(d_2153_v49_).cf34: ((self).f18) not in ((d_2151_v47_)[default__.safeIndex(667, (d_2151_v47_).length(0))])})
                d_2154_v50_ = d_2154_v50_
                d_2155_v51_: _dafny.Array
                nw379_ = _dafny.Array(None, 3)
                nw379_[int(0)] = (self).f18
                nw379_[int(1)] = True
                nw379_[int(2)] = default__.fm23(d_2094_v0_, p1, globalState)
                d_2155_v51_ = nw379_
                index367_ = default__.safeIndex(18, (d_2155_v51_).length(0))
                (d_2155_v51_)[index367_] = False
                index368_ = default__.safeIndex(18, (d_2155_v51_).length(0))
                (d_2155_v51_)[index368_] = ((_dafny.SeqWithoutIsStrInference([d_2152_v48_ for d_2156_i4_ in range(default__.abs(472))])) + (p2)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afcymv")))
                (globalState).f2 = not((self).f18)
            d_2157_v52_: D1
            d_2157_v52_ = D1_DC4(True)
            d_2158_v53_: D1
            d_2158_v53_ = D1_DC5(d_2157_v52_)
            source19_ = d_2158_v53_
            if source19_.is_DC3:
                d_2159___mcc_h0_ = source19_.cf7
                d_2160___mcc_h1_ = source19_.cf8
                d_2161___mcc_h2_ = source19_.cf9
                d_2162___mcc_h3_ = source19_.cf10
                d_2163___mcc_h4_ = source19_.cf11
                d_2164_cf11_ = d_2163___mcc_h4_
                d_2165_cf10_ = d_2162___mcc_h3_
                d_2166_cf9_ = d_2161___mcc_h2_
                d_2167_cf8_ = d_2160___mcc_h1_
                d_2168_cf7_ = d_2159___mcc_h0_
                d_2169_v54_: _dafny.Map
                d_2169_v54_ = _dafny.Map({p1: (self).f18})
                d_2169_v54_ = (d_2169_v54_).set((d_2138_v34_.f24) + (d_2138_v34_.f24), (self).f18)
                (globalState).f2 = (p0) < ((p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sunv"))))
                (globalState).f3 = self.f23
                d_2170_v55_: _dafny.Array
                def lambda89_(d_2171_i5_):
                    return (self).f18

                init56_ = lambda89_
                nw380_ = _dafny.Array(None, 29)
                for i0_56_ in range(nw380_.length(0)):
                    nw380_[i0_56_] = init56_(i0_56_)
                d_2170_v55_ = nw380_
                index369_ = default__.safeIndex(354, (d_2170_v55_).length(0))
                (d_2170_v55_)[index369_] = (self).f18
                index370_ = default__.safeIndex(354, (d_2170_v55_).length(0))
                (d_2170_v55_)[index370_] = (p0) != (p2)
            elif source19_.is_DC4:
                d_2172___mcc_h5_ = source19_.cf12
                d_2173_cf12_ = d_2172___mcc_h5_
                d_2174_v56_: _dafny.MultiSet
                d_2174_v56_ = _dafny.MultiSet([p0])
                (d_2138_v34_).m5((d_2174_v56_).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbqo")), p0])), default__.safeModuloInt(d_2138_v34_.f24, p1), default__.fm25(-768, globalState), len((p0) + (p2)), globalState)
                (globalState).f2 = (d_2140_v36_) < (d_2140_v36_)
                d_2175_v57_: _dafny.MultiSet
                d_2175_v57_ = _dafny.MultiSet([d_2173_cf12_])
                d_2176_v58_: _dafny.Map
                d_2176_v58_ = _dafny.Map({(d_2138_v34_.f24 if d_2173_cf12_ else self.f23): len(d_2094_v0_)})
                d_2177_v59_: D0
                d_2177_v59_ = D0_DC1((self).f18, -99, p2, d_2173_cf12_, default__.fm1(d_2138_v34_.f24, False, d_2173_cf12_, globalState))
                rhs379_ = (p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2178_i6_ in range(default__.abs(-1))]))
                rhs380_ = d_2175_v57_
                rhs381_ = ((d_2176_v58_)[len(d_2140_v36_)] if (len(d_2140_v36_)) in (d_2176_v58_) else (0) - (d_2138_v34_.f24))
                rhs382_ = (d_2139_v35_).fm22(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2179_i7_ in range(default__.abs(721))]), d_2177_v59_, globalState)
                rhs383_ = d_2173_cf12_
                lhs327_ = globalState
                lhs328_ = d_2138_v34_
                lhs329_ = globalState
                lhs330_ = globalState
                lhs327_.f12 = rhs379_
                d_2175_v57_ = rhs380_
                lhs328_.f24 = rhs381_
                lhs329_.f3 = rhs382_
                lhs330_.f2 = rhs383_
                d_2180_v60_: _dafny.Array
                def lambda90_(d_2181_v58_):
                    def lambda91_(d_2182_i8_):
                        return D20_DC52(d_2181_v58_)

                    return lambda91_

                init57_ = lambda90_(d_2176_v58_)
                nw381_ = _dafny.Array(None, 14)
                for i0_57_ in range(nw381_.length(0)):
                    nw381_[i0_57_] = init57_(i0_57_)
                d_2180_v60_ = nw381_
                d_2183_v63_: D20
                def iife182_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in _dafny.IntegerRange(-108, 430):
                        d_2184_v61_: int = compr_56_
                        if ((-108) <= (d_2184_v61_)) and ((d_2184_v61_) < (430)):
                            def iife183_():
                                coll57_ = _dafny.Map()
                                compr_57_: int
                                for compr_57_ in _dafny.IntegerRange(424, 779):
                                    d_2185_v62_: int = compr_57_
                                    if ((424) <= (d_2185_v62_)) and ((d_2185_v62_) < (779)):
                                        coll57_[(d_2185_v62_) + (p1)] = True
                                return _dafny.Map(coll57_)
                            coll56_[default__.safeModuloInt(d_2184_v61_, self.f23)] = len(iife183_()
                            )
                    return _dafny.Map(coll56_)
                d_2183_v63_ = D20_DC52(iife182_()
)
                index371_ = default__.safeIndex(965, (d_2180_v60_).length(0))
                (d_2180_v60_)[index371_] = d_2183_v63_
                index372_ = default__.safeIndex(965, (d_2180_v60_).length(0))
                (d_2180_v60_)[index372_] = d_2183_v63_
            elif source19_.is_DC2:
                d_2186___mcc_h6_ = source19_.cf6
                d_2187_cf6_ = d_2186___mcc_h6_
                (globalState).f2 = (self).f18
                d_2188_v64_: str
                d_2188_v64_ = _dafny.CodePoint('f')
                d_2189_v65_: _dafny.MultiSet
                d_2189_v65_ = _dafny.MultiSet([p0, p0, ((p0 if False else p0)).set(default__.safeIndex(self.f23, len((p0 if False else p0))), d_2188_v64_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlsydos")), (p0) + (p2)])
                d_2189_v65_ = default__.fm52(d_2138_v34_.f24, d_2138_v34_.f24, globalState)
                d_2188_v64_ = d_2188_v64_
                d_2190_v66_: _dafny.Map
                d_2190_v66_ = _dafny.Map({d_2138_v34_.f24: (d_2187_cf6_).f18})
                (d_2138_v34_).f24 = (d_2138_v34_.f24) + (len((d_2190_v66_) | (d_2190_v66_)))
            elif True:
                d_2191___mcc_h7_ = source19_.cf13
                d_2192_cf13_ = d_2191___mcc_h7_
                d_2193_v67_: _dafny.Array
                nw382_ = _dafny.Array(False, 2)
                d_2193_v67_ = nw382_
                d_2194_v68_: _dafny.Set
                d_2194_v68_ = _dafny.Set({d_2193_v67_, d_2193_v67_})
                (globalState).f2 = ((d_2194_v68_) - (d_2194_v68_)).issubset(d_2194_v68_)
                d_2195_v69_: C0
                nw383_ = C0()
                nw383_.ctor__(d_2138_v34_.f24, (0) - (p1))
                d_2195_v69_ = nw383_
                d_2196_v70_: _dafny.Array
                def lambda92_(d_2197_i9_):
                    return (_dafny.SeqWithoutIsStrInference([(self).f18])) + (_dafny.SeqWithoutIsStrInference([not(True)]))

                init58_ = lambda92_
                nw384_ = _dafny.Array(None, 7)
                for i0_58_ in range(nw384_.length(0)):
                    nw384_[i0_58_] = init58_(i0_58_)
                d_2196_v70_ = nw384_
                index373_ = default__.safeIndex(919, (d_2196_v70_).length(0))
                (d_2196_v70_)[index373_] = _dafny.SeqWithoutIsStrInference([default__.fm23(d_2094_v0_, d_2138_v34_.f24, globalState), (self).f18, not((self).f18), not((self).f18), (self).f18])
                d_2198_v71_: _dafny.Seq
                d_2198_v71_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
                index374_ = default__.safeIndex(919, (d_2196_v70_).length(0))
                (d_2196_v70_)[index374_] = (_dafny.SeqWithoutIsStrInference([not((self).f18)])) + (d_2198_v71_)
                (d_2138_v34_).f24 = (0) - (d_2138_v34_.f24)
            d_2199_v72_: D6
            d_2199_v72_ = D6_DC18()
            source20_ = d_2199_v72_
            if source20_.is_DC18:
                d_2200_v73_: _dafny.Array
                def lambda93_(d_2201_i10_):
                    return _dafny.CodePoint('c')

                init59_ = lambda93_
                nw385_ = _dafny.Array(None, 4)
                for i0_59_ in range(nw385_.length(0)):
                    nw385_[i0_59_] = init59_(i0_59_)
                d_2200_v73_ = nw385_
                d_2202_v74_: _dafny.MultiSet
                d_2202_v74_ = _dafny.MultiSet([(self).f18, (self).f18])
                d_2203_v75_: _dafny.Map
                d_2203_v75_ = _dafny.Map({d_2202_v74_: (d_2138_v34_).fm10((self).f18, (self).f18, (self).f18, globalState)})
                d_2204_v76_: D3
                d_2204_v76_ = D3_DC12(d_2200_v73_, (self).f18, p1, d_2203_v75_, (self).f18)
                (globalState).f3 = ((0) - (((d_2140_v36_)[default__.safeIndex(p1, len(d_2140_v36_))]) - ((d_2204_v76_).cf27))) * (self.f23)
                d_2205_v77_: str
                d_2205_v77_ = _dafny.CodePoint('i')
                d_2206_v78_: D20
                d_2206_v78_ = D20_DC53((self).f18, (self).f18, d_2205_v77_)
                d_2207_v79_: _dafny.Seq
                d_2207_v79_ = _dafny.SeqWithoutIsStrInference([d_2139_v35_, d_2139_v35_, d_2139_v35_, d_2139_v35_, d_2139_v35_])
                rhs384_ = d_2206_v78_
                rhs385_ = (d_2207_v79_) + (d_2207_v79_)
                d_2206_v78_ = rhs384_
                d_2207_v79_ = rhs385_
                (d_2138_v34_).f24 = d_2138_v34_.f24
                d_2208_v80_: _dafny.Array
                nw386_ = _dafny.Array(False, 15)
                d_2208_v80_ = nw386_
                d_2209_v81_: _dafny.Array
                nw387_ = _dafny.Array(D1.default()(), 6)
                d_2209_v81_ = nw387_
                d_2210_v82_: D1
                d_2210_v82_ = D1_DC4((self).f18)
                index375_ = default__.safeIndex(718, (d_2209_v81_).length(0))
                (d_2209_v81_)[index375_] = d_2210_v82_
                d_2211_v83_: _dafny.Seq
                d_2211_v83_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2212_v84_: _dafny.Seq
                d_2212_v84_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
                d_2213_v85_: _dafny.Map
                d_2213_v85_ = _dafny.Map({self.f23: len((d_2211_v83_)[default__.safeIndex(len(d_2212_v84_), len(d_2211_v83_))])})
                d_2214_v86_: _dafny.Map
                d_2214_v86_ = _dafny.Map({p1: d_2213_v85_})
                d_2215_v87_: _dafny.Seq
                d_2215_v87_ = _dafny.SeqWithoutIsStrInference([d_2214_v86_])
                d_2216_v90_: _dafny.Array
                nw388_ = _dafny.Array(None, 25)
                nw388_[int(0)] = d_2214_v86_
                nw388_[int(1)] = (d_2214_v86_) | ((d_2214_v86_).set(len(d_2213_v85_), d_2213_v85_))
                nw388_[int(2)] = d_2214_v86_
                nw388_[int(3)] = d_2214_v86_
                nw388_[int(4)] = _dafny.Map({d_2138_v34_.f24: d_2213_v85_})
                nw388_[int(5)] = (d_2215_v87_)[default__.safeIndex(951, len(d_2215_v87_))]
                nw388_[int(6)] = d_2214_v86_
                nw388_[int(7)] = _dafny.Map({self.f23: d_2213_v85_})
                nw388_[int(8)] = d_2214_v86_
                nw388_[int(9)] = d_2214_v86_
                def iife184_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in _dafny.IntegerRange(-146, -195):
                        d_2217_v88_: int = compr_58_
                        if ((-146) <= (d_2217_v88_)) and ((d_2217_v88_) < (-195)):
                            coll58_[(d_2217_v88_) * ((0) - (len(d_2213_v85_)))] = (d_2140_v36_)[default__.safeIndex(len(d_2212_v84_), len(d_2140_v36_))]
                    return _dafny.Map(coll58_)
                nw388_[int(10)] = (_dafny.Map({default__.fm1(self.f23, (self).f18, (self).f18, globalState): iife184_()
                }) if (d_2212_v84_)[default__.safeIndex(d_2138_v34_.f24, len(d_2212_v84_))] else d_2214_v86_)
                nw388_[int(11)] = (d_2214_v86_) | (d_2214_v86_)
                nw388_[int(12)] = d_2214_v86_
                nw388_[int(13)] = d_2214_v86_
                nw388_[int(14)] = default__.fm60(d_2212_v84_, (self).f18, (self).f18, globalState)
                def iife185_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(884, 857):
                        d_2218_v89_: int = compr_59_
                        if ((884) <= (d_2218_v89_)) and ((d_2218_v89_) < (857)):
                            coll59_[default__.safeModuloInt(d_2218_v89_, d_2138_v34_.f24)] = d_2213_v85_
                    return _dafny.Map(coll59_)
                nw388_[int(15)] = iife185_()
                
                nw388_[int(16)] = (d_2214_v86_) | (d_2214_v86_)
                nw388_[int(17)] = d_2214_v86_
                nw388_[int(18)] = d_2214_v86_
                nw388_[int(19)] = d_2214_v86_
                nw388_[int(20)] = d_2214_v86_
                nw388_[int(21)] = (d_2214_v86_) | (d_2214_v86_)
                nw388_[int(22)] = d_2214_v86_
                nw388_[int(23)] = _dafny.Map({p1: _dafny.Map({d_2138_v34_.f24: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "babsckc")))})})
                nw388_[int(24)] = (d_2214_v86_).set(d_2138_v34_.f24, d_2213_v85_)
                d_2216_v90_ = nw388_
                index376_ = default__.safeIndex(164, (d_2216_v90_).length(0))
                (d_2216_v90_)[index376_] = d_2214_v86_
                d_2219_v91_: D2
                d_2219_v91_ = D2_DC7(d_2138_v34_.f24, (self).f18, 956)
                index377_ = default__.safeIndex(718, (d_2209_v81_).length(0))
                index378_ = default__.safeIndex(164, (d_2216_v90_).length(0))
                rhs386_ = d_2208_v80_
                rhs387_ = (-41) * (974)
                rhs388_ = D1_DC4(not((self).f18))
                rhs389_ = default__.fm40(d_2219_v91_, globalState)
                rhs390_ = d_2214_v86_
                lhs331_ = globalState
                lhs332_ = d_2209_v81_
                lhs333_ = default__.safeIndex(718, (d_2209_v81_).length(0))
                lhs334_ = globalState
                lhs335_ = d_2216_v90_
                lhs336_ = default__.safeIndex(164, (d_2216_v90_).length(0))
                d_2208_v80_ = rhs386_
                lhs331_.f0 = rhs387_
                lhs332_[lhs333_] = rhs388_
                lhs334_.f10 = rhs389_
                lhs335_[lhs336_] = rhs390_
            elif True:
                d_2220___mcc_h8_ = source20_.cf38
                d_2221_cf38_ = d_2220___mcc_h8_
                rhs391_ = (self).f18
                rhs392_ = d_2138_v34_.f24
                rhs393_ = (self).f18
                rhs394_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2222_i11_ in range(default__.abs(763))])
                lhs337_ = globalState
                lhs338_ = globalState
                lhs339_ = globalState
                lhs340_ = globalState
                lhs337_.f2 = rhs391_
                lhs338_.f3 = rhs392_
                lhs339_.f2 = rhs393_
                lhs340_.f9 = rhs394_
                (globalState).f2 = (self).f18
                (globalState).f2 = (self).f18
                d_2223_v92_: int
                d_2224_v93_: int
                d_2225_v94_: _dafny.Array
                d_2226_v95_: _dafny.Map
                out58_: int
                out59_: int
                out60_: _dafny.Array
                out61_: _dafny.Map
                out58_, out59_, out60_, out61_ = default__.m0(globalState)
                d_2223_v92_ = out58_
                d_2224_v93_ = out59_
                d_2225_v94_ = out60_
                d_2226_v95_ = out61_
        elif True:
            d_2227_v96_: _dafny.Seq
            d_2227_v96_ = _dafny.SeqWithoutIsStrInference([True, False])
            d_2228_v97_: _dafny.MultiSet
            d_2228_v97_ = _dafny.MultiSet([not((self).f18)])
            d_2229_v98_: _dafny.Map
            d_2229_v98_ = _dafny.Map({d_2138_v34_.f24: (self).f18})
            d_2230_v99_: _dafny.Map
            d_2230_v99_ = _dafny.Map({(self).f18: (self).f18})
            d_2231_v100_: _dafny.Map
            d_2231_v100_ = _dafny.Map({d_2230_v99_: p1})
            d_2232_v101_: _dafny.Seq
            d_2232_v101_ = _dafny.SeqWithoutIsStrInference([self.f23, d_2138_v34_.f24, len(d_2231_v100_), p1, d_2138_v34_.f24])
            d_2233_v102_: T1
            nw389_ = C3()
            nw389_.ctor__(d_2232_v101_, p0, _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2138_v34_.f24: len(_dafny.Set({(self).f18, (self).f18, (self).f18}))}) for d_2234_i12_ in range(default__.abs(203))]), (self).f18)
            d_2233_v102_ = nw389_
            d_2235_v103_: _dafny.Map
            d_2235_v103_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljsg")): d_2138_v34_.f24})
            d_2236_v104_: _dafny.Array
            nw390_ = _dafny.Array(None, 17)
            nw390_[int(0)] = len(_dafny.SeqWithoutIsStrInference([d_2227_v96_, d_2227_v96_]))
            nw390_[int(1)] = ((d_2228_v97_)[((d_2229_v98_)[len(p0)] if (len(p0)) in (d_2229_v98_) else (d_2227_v96_)[default__.safeIndex(973, len(d_2227_v96_))])] if (((d_2229_v98_)[len(p0)] if (len(p0)) in (d_2229_v98_) else (d_2227_v96_)[default__.safeIndex(973, len(d_2227_v96_))])) in (d_2228_v97_) else len(d_2232_v101_))
            nw390_[int(2)] = p1
            nw390_[int(3)] = d_2138_v34_.f24
            nw390_[int(4)] = len(_dafny.SeqWithoutIsStrInference([d_2233_v102_]))
            nw390_[int(5)] = self.f23
            nw390_[int(6)] = len(_dafny.Map({(d_2233_v102_).f18: (d_2233_v102_).f18}))
            nw390_[int(7)] = len(d_2235_v103_)
            nw390_[int(8)] = (0) - (d_2138_v34_.f24)
            nw390_[int(9)] = d_2138_v34_.f24
            nw390_[int(10)] = self.f23
            nw390_[int(11)] = len(_dafny.Map({d_2138_v34_.f24: d_2138_v34_.f24}))
            nw390_[int(12)] = len(d_2227_v96_)
            nw390_[int(13)] = p1
            nw390_[int(14)] = (d_2228_v97_).cardinality
            nw390_[int(15)] = 373
            nw390_[int(16)] = d_2138_v34_.f24
            d_2236_v104_ = nw390_
            d_2237_v105_: _dafny.Map
            d_2237_v105_ = _dafny.Map({d_2236_v104_: p2})
            d_2237_v105_ = (d_2237_v105_) | (_dafny.Map({d_2236_v104_: p0}))
            index379_ = default__.safeIndex(337, (d_2236_v104_).length(0))
            (d_2236_v104_)[index379_] = (d_2138_v34_.f24 if (d_2233_v102_).f18 else p1)
            index380_ = default__.safeIndex(337, (d_2236_v104_).length(0))
            (d_2236_v104_)[index380_] = (297) * (d_2138_v34_.f24)
            d_2238_v106_: str
            d_2238_v106_ = _dafny.CodePoint('m')
            d_2239_v107_: _dafny.Map
            d_2239_v107_ = _dafny.Map({self.f23: d_2238_v106_})
            d_2239_v107_ = (d_2239_v107_).set(self.f23, d_2238_v106_)
            (globalState).f2 = (self).f18
            d_2240_v108_: _dafny.Array
            nw391_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_2240_v108_ = nw391_
            d_2241_v109_: _dafny.Array
            def lambda94_(d_2242_i13_):
                return _dafny.MultiSet([self.f23])

            init60_ = lambda94_
            nw392_ = _dafny.Array(None, 9)
            for i0_60_ in range(nw392_.length(0)):
                nw392_[i0_60_] = init60_(i0_60_)
            d_2241_v109_ = nw392_
            d_2243_v110_: D22
            d_2243_v110_ = D22_DC57(d_2241_v109_)
            index381_ = default__.safeIndex(935, (d_2240_v108_).length(0))
            (d_2240_v108_)[index381_] = (d_2243_v110_).cf95
            d_2244_v111_: D17
            d_2244_v111_ = D17_DC41(d_2229_v98_)
            d_2245_v112_: _dafny.Array
            def lambda95_(d_2246_i14_):
                return True

            init61_ = lambda95_
            nw393_ = _dafny.Array(None, 24)
            for i0_61_ in range(nw393_.length(0)):
                nw393_[i0_61_] = init61_(i0_61_)
            d_2245_v112_ = nw393_
            d_2247_v113_: D10
            d_2247_v113_ = D10_DC25((d_2233_v102_).f18)
            index382_ = default__.safeIndex(170, (d_2245_v112_).length(0))
            (d_2245_v112_)[index382_] = (d_2247_v113_).cf48
            index383_ = default__.safeIndex(935, (d_2240_v108_).length(0))
            index384_ = default__.safeIndex(170, (d_2245_v112_).length(0))
            rhs395_ = d_2241_v109_
            rhs396_ = d_2244_v111_
            rhs397_ = (self).f18
            rhs398_ = (self).f18
            lhs341_ = d_2240_v108_
            lhs342_ = default__.safeIndex(935, (d_2240_v108_).length(0))
            lhs343_ = d_2245_v112_
            lhs344_ = default__.safeIndex(170, (d_2245_v112_).length(0))
            lhs345_ = globalState
            lhs341_[lhs342_] = rhs395_
            d_2244_v111_ = rhs396_
            lhs343_[lhs344_] = rhs397_
            lhs345_.f2 = rhs398_
        d_2248_v114_: D1
        d_2248_v114_ = D1_DC4((self).f18)
        def lambda96_(source21_):
            if source21_.is_DC3:
                d_2249___mcc_h9_ = source21_.cf7
                d_2250___mcc_h10_ = source21_.cf8
                d_2251___mcc_h11_ = source21_.cf9
                d_2252___mcc_h12_ = source21_.cf10
                d_2253___mcc_h13_ = source21_.cf11
                d_2254_cf11_ = d_2253___mcc_h13_
                d_2255_cf10_ = d_2252___mcc_h12_
                d_2256_cf9_ = d_2251___mcc_h11_
                d_2257_cf8_ = d_2250___mcc_h10_
                d_2258_cf7_ = d_2249___mcc_h9_
                return not ((self).f18) or (not((self).f18))
            elif source21_.is_DC4:
                d_2259___mcc_h14_ = source21_.cf12
                d_2260_cf12_ = d_2259___mcc_h14_
                return (self).f18
            elif source21_.is_DC2:
                d_2261___mcc_h15_ = source21_.cf6
                d_2262_cf6_ = d_2261___mcc_h15_
                return (_dafny.MultiSet([(d_2262_cf6_).f18])).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])))
            elif True:
                d_2263___mcc_h16_ = source21_.cf13
                d_2264_cf13_ = d_2263___mcc_h16_
                return (self).f18

        (globalState).f2 = lambda96_(d_2248_v114_)
        (globalState).f2 = default__.fm23(d_2094_v0_, d_2138_v34_.f24, globalState)
        if (self).f18:
            d_2265_v115_: _dafny.Map
            d_2265_v115_ = _dafny.Map({(p1) > ((0) - (p1)): default__.safeDivisionInt(d_2138_v34_.f24, self.f23)})
            d_2265_v115_ = (d_2265_v115_).set((self).f18, self.f23)
            d_2266_v116_: _dafny.Array
            def lambda97_(d_2267_p1_):
                def lambda98_(d_2268_i15_):
                    return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2267_p1_]))

                return lambda98_

            init62_ = lambda97_(p1)
            nw394_ = _dafny.Array(None, 28)
            for i0_62_ in range(nw394_.length(0)):
                nw394_[i0_62_] = init62_(i0_62_)
            d_2266_v116_ = nw394_
            d_2269_v117_: _dafny.MultiSet
            d_2269_v117_ = _dafny.MultiSet([p1])
            index385_ = default__.safeIndex(363, (d_2266_v116_).length(0))
            (d_2266_v116_)[index385_] = (d_2269_v117_) | (d_2269_v117_)
            index386_ = default__.safeIndex(363, (d_2266_v116_).length(0))
            (d_2266_v116_)[index386_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([40 for d_2270_i16_ in range(default__.abs(692))]))).set(p1, default__.abs(default__.safeModuloInt(472, self.f23)))
            (globalState).f9 = default__.fm13((self).f18, p1, globalState)
            (globalState).f9 = (p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
            d_2265_v115_ = default__.fm43((self).f18, False, globalState)
        elif True:
            d_2271_v118_: _dafny.Array
            nw395_ = _dafny.Array(None, 1)
            nw395_[int(0)] = d_2138_v34_.f24
            d_2271_v118_ = nw395_
            d_2272_v119_: D17
            d_2272_v119_ = D17_DC42((self).f18)
            d_2273_v120_: _dafny.Map
            d_2273_v120_ = _dafny.Map({(d_2272_v119_).cf74: d_2138_v34_.f24})
            d_2274_v121_: _dafny.Set
            d_2274_v121_ = _dafny.Set({d_2273_v120_})
            d_2275_v122_: _dafny.Map
            d_2275_v122_ = _dafny.Map({d_2271_v118_: (p1) + (len(d_2274_v121_))})
            d_2275_v122_ = (d_2275_v122_).set(d_2271_v118_, (self.f23) - (self.f23))
            (globalState).f2 = not (((self).f18) and ((self).f18)) or ((self).f18)
            d_2276_v123_: _dafny.Set
            d_2276_v123_ = _dafny.Set({(self).f18, (self).f18})
            d_2277_v124_: _dafny.Seq
            d_2277_v124_ = _dafny.SeqWithoutIsStrInference([d_2276_v123_, d_2276_v123_])
            d_2278_v125_: str
            d_2278_v125_ = _dafny.CodePoint('h')
            d_2279_v126_: D3
            d_2279_v126_ = D3_DC13(d_2138_v34_.f24, (p2) + (p0), (d_2138_v34_.f24) + (default__.fm1((_dafny.MultiSet(d_2277_v124_)).cardinality, (self).f18, (self).f18, globalState)), d_2138_v34_.f24, d_2278_v125_)
            source22_ = d_2279_v126_
            if source22_.is_DC11:
                d_2280___mcc_h17_ = source22_.cf23
                d_2281___mcc_h18_ = source22_.cf24
                d_2282_cf24_ = d_2281___mcc_h18_
                d_2283_cf23_ = d_2280___mcc_h17_
                d_2284_v127_: _dafny.Array
                def lambda99_(d_2285_cf23_):
                    def lambda100_(d_2286_i17_):
                        return d_2285_cf23_

                    return lambda100_

                init63_ = lambda99_(d_2283_cf23_)
                nw396_ = _dafny.Array(None, 9)
                for i0_63_ in range(nw396_.length(0)):
                    nw396_[i0_63_] = init63_(i0_63_)
                d_2284_v127_ = nw396_
                d_2287_v128_: _dafny.Map
                d_2287_v128_ = _dafny.Map({default__.fm12((0) - (d_2138_v34_.f24), d_2283_cf23_, globalState): d_2284_v127_})
                d_2288_v129_: _dafny.MultiSet
                d_2288_v129_ = _dafny.MultiSet([670, d_2138_v34_.f24])
                rhs399_ = (d_2287_v128_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocfdgltdd")): d_2284_v127_}))
                rhs400_ = (d_2288_v129_).issubset(d_2288_v129_)
                d_2287_v128_ = rhs399_
                d_2283_cf23_ = rhs400_
                d_2289_v130_: D13
                d_2289_v130_ = D13_DC33((self).f18, d_2138_v34_.f24, d_2283_cf23_, -352)
                d_2290_v131_: _dafny.Map
                d_2290_v131_ = _dafny.Map({self.f23: d_2283_cf23_})
                d_2291_v132_: D16
                d_2291_v132_ = D16_DC39(False, self.f23)
                d_2292_v133_: D2
                d_2292_v133_ = D2_DC7((d_2288_v129_).cardinality, (d_2291_v132_).cf67, p1)
                d_2293_v134_: _dafny.Map
                d_2293_v134_ = _dafny.Map({d_2283_cf23_: d_2282_cf24_})
                pat_let_tv35_ = d_2282_cf24_
                d_2294_v135_: _dafny.Array
                nw397_ = _dafny.Array(None, 14)
                nw397_[int(0)] = d_2289_v130_
                nw397_[int(1)] = d_2289_v130_
                nw397_[int(2)] = D13_DC33(((d_2290_v131_)[d_2138_v34_.f24] if (d_2138_v34_.f24) in (d_2290_v131_) else d_2283_cf23_), 553, True, d_2138_v34_.f24)
                nw397_[int(3)] = d_2289_v130_
                def iife186_(_pat_let63_0):
                    def iife187_(d_2295_dt__update__tmp_h3_):
                        def iife188_(_pat_let64_0):
                            def iife189_(d_2296_dt__update_hcf57_h0_):
                                def iife190_(_pat_let65_0):
                                    def iife191_(d_2297_dt__update_hcf55_h0_):
                                        return D13_DC33(d_2297_dt__update_hcf55_h0_, (d_2295_dt__update__tmp_h3_).cf56, d_2296_dt__update_hcf57_h0_, (d_2295_dt__update__tmp_h3_).cf58)
                                    return iife191_(_pat_let65_0)
                                return iife190_((self).f18)
                            return iife189_(_pat_let64_0)
                        return iife188_(pat_let_tv35_)
                    return iife187_(_pat_let63_0)
                nw397_[int(4)] = iife186_(d_2289_v130_)
                nw397_[int(5)] = d_2289_v130_
                nw397_[int(6)] = d_2289_v130_
                nw397_[int(7)] = d_2289_v130_
                nw397_[int(8)] = d_2289_v130_
                nw397_[int(9)] = d_2289_v130_
                nw397_[int(10)] = (D13_DC33(True, p1, d_2283_cf23_, d_2138_v34_.f24) if default__.fm23(d_2094_v0_, self.f23, globalState) else d_2289_v130_)
                nw397_[int(11)] = d_2289_v130_
                nw397_[int(12)] = D13_DC33((d_2292_v133_).cf16, p1, d_2282_cf24_, len(d_2293_v134_))
                nw397_[int(13)] = d_2289_v130_
                d_2294_v135_ = nw397_
                d_2294_v135_ = d_2294_v135_
                d_2298_v136_: _dafny.Seq
                d_2298_v136_ = _dafny.SeqWithoutIsStrInference([d_2283_cf23_, (self).f18])
                d_2299_v137_: _dafny.Map
                d_2299_v137_ = _dafny.Map({d_2298_v136_: p1})
                d_2300_v138_: _dafny.Seq
                d_2300_v138_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_2298_v136_: (0) - (p1)})).set(d_2298_v136_, 151)])
                rhs401_ = (self).f18
                rhs402_ = (d_2300_v138_)[default__.safeIndex(d_2138_v34_.f24, len(d_2300_v138_))]
                rhs403_ = (0) - ((len(p2)) * (p1))
                lhs346_ = globalState
                lhs347_ = globalState
                lhs346_.f2 = rhs401_
                d_2299_v137_ = rhs402_
                lhs347_.f0 = rhs403_
                index387_ = default__.safeIndex(649, (d_2271_v118_).length(0))
                (d_2271_v118_)[index387_] = (len(p0)) - (d_2138_v34_.f24)
                index388_ = default__.safeIndex(649, (d_2271_v118_).length(0))
                (d_2271_v118_)[index388_] = 813
            elif source22_.is_DC12:
                d_2301___mcc_h19_ = source22_.cf25
                d_2302___mcc_h20_ = source22_.cf26
                d_2303___mcc_h21_ = source22_.cf27
                d_2304___mcc_h22_ = source22_.cf28
                d_2305___mcc_h23_ = source22_.cf29
                d_2306_cf29_ = d_2305___mcc_h23_
                d_2307_cf28_ = d_2304___mcc_h22_
                d_2308_cf27_ = d_2303___mcc_h21_
                d_2309_cf26_ = d_2302___mcc_h20_
                d_2310_cf25_ = d_2301___mcc_h19_
                (globalState).f15 = (0) - ((d_2138_v34_.f24) - ((-724) * (p1)))
                (globalState).f2 = (self).f18
                d_2311_v139_: _dafny.Array
                nw398_ = _dafny.Array(False, 20)
                d_2311_v139_ = nw398_
                index389_ = default__.safeIndex(344, (d_2311_v139_).length(0))
                (d_2311_v139_)[index389_] = d_2306_cf29_
                index390_ = default__.safeIndex(344, (d_2311_v139_).length(0))
                (d_2311_v139_)[index390_] = (self).f18
                d_2308_cf27_ = 3
            elif source22_.is_DC13:
                d_2312___mcc_h24_ = source22_.cf30
                d_2313___mcc_h25_ = source22_.cf31
                d_2314___mcc_h26_ = source22_.cf32
                d_2315___mcc_h27_ = source22_.cf33
                d_2316___mcc_h28_ = source22_.cf34
                d_2317_cf34_ = d_2316___mcc_h28_
                d_2318_cf33_ = d_2315___mcc_h27_
                d_2319_cf32_ = d_2314___mcc_h26_
                d_2320_cf31_ = d_2313___mcc_h25_
                d_2321_cf30_ = d_2312___mcc_h24_
                d_2322_v140_: _dafny.Array
                nw399_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2322_v140_ = nw399_
                d_2323_v141_: _dafny.Array
                nw400_ = _dafny.Array(None, 23)
                nw400_[int(0)] = d_2317_cf34_
                nw400_[int(1)] = d_2317_cf34_
                nw400_[int(2)] = d_2278_v125_
                nw400_[int(3)] = d_2278_v125_
                nw400_[int(4)] = (p2)[default__.safeIndex(self.f23, len(p2))]
                nw400_[int(5)] = d_2317_cf34_
                nw400_[int(6)] = d_2317_cf34_
                nw400_[int(7)] = d_2317_cf34_
                nw400_[int(8)] = (d_2279_v126_).cf34
                nw400_[int(9)] = d_2278_v125_
                nw400_[int(10)] = d_2278_v125_
                nw400_[int(11)] = d_2317_cf34_
                nw400_[int(12)] = d_2317_cf34_
                nw400_[int(13)] = d_2317_cf34_
                nw400_[int(14)] = d_2317_cf34_
                nw400_[int(15)] = _dafny.CodePoint('n')
                nw400_[int(16)] = d_2278_v125_
                nw400_[int(17)] = (d_2320_cf31_)[default__.safeIndex(d_2319_cf32_, len(d_2320_cf31_))]
                nw400_[int(18)] = _dafny.CodePoint('b')
                nw400_[int(19)] = d_2317_cf34_
                nw400_[int(20)] = d_2278_v125_
                nw400_[int(21)] = d_2278_v125_
                nw400_[int(22)] = d_2278_v125_
                d_2323_v141_ = nw400_
                index391_ = default__.safeIndex(43, (d_2322_v140_).length(0))
                (d_2322_v140_)[index391_] = d_2323_v141_
                index392_ = default__.safeIndex(43, (d_2322_v140_).length(0))
                (d_2322_v140_)[index392_] = d_2323_v141_
                rhs404_ = _dafny.Set({d_2321_cf30_, default__.fm1((0) - (d_2321_cf30_), (self).f18, (self).f18, globalState), d_2318_cf33_, default__.safeModuloInt(self.f23, d_2138_v34_.f24), default__.safeDivisionInt(p1, d_2138_v34_.f24)})
                rhs405_ = (self).f18
                rhs406_ = d_2278_v125_
                lhs348_ = globalState
                d_2094_v0_ = rhs404_
                lhs348_.f2 = rhs405_
                d_2317_cf34_ = rhs406_
                (globalState).f9 = p2
                (self).f23 = 442
            elif source22_.is_DC10:
                d_2324___mcc_h29_ = source22_.cf22
                d_2325_cf22_ = d_2324___mcc_h29_
                (globalState).f2 = (((_dafny.SeqWithoutIsStrInference([d_2278_v125_ for d_2326_i18_ in range(default__.abs(312))])).set(default__.safeIndex(d_2138_v34_.f24, len(_dafny.SeqWithoutIsStrInference([d_2278_v125_ for d_2327_i18_ in range(default__.abs(312))]))), _dafny.CodePoint('l'))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbxr")))) or ((self).f18)
                d_2328_v142_: _dafny.MultiSet
                d_2328_v142_ = _dafny.MultiSet([(self).f18])
                d_2328_v142_ = d_2328_v142_
                d_2329_v143_: _dafny.Map
                d_2329_v143_ = _dafny.Map({d_2138_v34_.f24: (self).f18})
                d_2329_v143_ = (d_2329_v143_).set((self.f23) - (d_2138_v34_.f24), (self).f18)
                d_2330_v144_: _dafny.Array
                def lambda101_(d_2331_i19_):
                    return (self).f18

                init64_ = lambda101_
                nw401_ = _dafny.Array(None, 24)
                for i0_64_ in range(nw401_.length(0)):
                    nw401_[i0_64_] = init64_(i0_64_)
                d_2330_v144_ = nw401_
                d_2332_v145_: D3
                d_2332_v145_ = D3_DC11((self).f18, (self).f18)
                d_2333_v146_: _dafny.Seq
                d_2333_v146_ = _dafny.SeqWithoutIsStrInference([not((d_2332_v145_).cf23)])
                d_2334_v147_: _dafny.Map
                d_2334_v147_ = _dafny.Map({d_2138_v34_.f24: d_2138_v34_.f24})
                d_2335_v148_: _dafny.Set
                d_2335_v148_ = _dafny.Set({d_2334_v147_})
                d_2336_v149_: _dafny.Map
                d_2336_v149_ = _dafny.Map({len(d_2333_v146_): d_2335_v148_})
                d_2337_v150_: D21
                d_2337_v150_ = D21_DC55((((d_2336_v149_)[self.f23] if (self.f23) in (d_2336_v149_) else _dafny.Set({_dafny.Map({len(d_2094_v0_): d_2138_v34_.f24}), d_2334_v147_}))).intersection(d_2335_v148_))
                d_2338_v151_: _dafny.Seq
                d_2338_v151_ = _dafny.SeqWithoutIsStrInference([d_2138_v34_.f24])
                rhs407_ = d_2330_v144_
                rhs408_ = d_2278_v125_
                rhs409_ = (d_2338_v151_) < (d_2338_v151_)
                rhs410_ = default__.fm61((self).f18, globalState)
                rhs411_ = d_2271_v118_
                lhs349_ = globalState
                d_2330_v144_ = rhs407_
                d_2278_v125_ = rhs408_
                lhs349_.f2 = rhs409_
                d_2337_v150_ = rhs410_
                d_2271_v118_ = rhs411_
            elif True:
                d_2339___mcc_h30_ = source22_.cf35
                d_2340_cf35_ = d_2339___mcc_h30_
                d_2341_v152_: _dafny.Array
                nw402_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_2341_v152_ = nw402_
                d_2342_v153_: _dafny.Array
                nw403_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_2342_v153_ = nw403_
                index393_ = default__.safeIndex(396, (d_2341_v152_).length(0))
                (d_2341_v152_)[index393_] = d_2342_v153_
                index394_ = default__.safeIndex(396, (d_2341_v152_).length(0))
                (d_2341_v152_)[index394_] = d_2342_v153_
                (globalState).f2 = (self).f18
                d_2343_v154_: D1
                d_2343_v154_ = D1_DC3(d_2138_v34_.f24, -308, (0) - ((0) - (d_2138_v34_.f24)), p1, p1)
                d_2344_v155_: D1
                d_2344_v155_ = D1_DC5(d_2343_v154_)
                d_2344_v155_ = d_2344_v155_
                (globalState).f2 = (self).f18
            (globalState).f0 = default__.fm1(d_2138_v34_.f24, (self).f18, (self).f18, globalState)
            d_2345_v156_: _dafny.Array
            def lambda102_(d_2346_i20_):
                return False

            init65_ = lambda102_
            nw404_ = _dafny.Array(None, 19)
            for i0_65_ in range(nw404_.length(0)):
                nw404_[i0_65_] = init65_(i0_65_)
            d_2345_v156_ = nw404_
            d_2347_v157_: _dafny.Seq
            d_2347_v157_ = _dafny.SeqWithoutIsStrInference([d_2345_v156_, d_2345_v156_, d_2345_v156_, d_2345_v156_])
            d_2347_v157_ = d_2347_v157_


class C9(T0):
    def  __init__(self):
        self._f18: bool = False
        self.f22: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f22, f18):
        (self).f22 = f22
        (self)._f18 = f18

    def fm4(self, p0, p1, globalState):
        return (835) * ((127 if (self).f18 else len(_dafny.SeqWithoutIsStrInference([(self).f18, (self).f18]))))

    def fm5(self, p0, p1, globalState):
        return (((0) - ((_dafny.MultiSet([len(_dafny.Map({-689: (self).f18}))])).cardinality)) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maumun"))), -107})))) - (848)

    def m1(self, p0, p1, p2, globalState):
        d_2348_v0_: D0
        d_2348_v0_ = D0_DC0((self).f18)
        d_2349_i0_: int
        d_2349_i0_ = 0
        with _dafny.label("15"):
            while (d_2348_v0_).cf0:
                with _dafny.c_label("15"):
                    if (d_2349_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2349_i0_ = (d_2349_i0_) + (1)
                    rhs412_ = (self).f18
                    rhs413_ = ((p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srlae")))) == ((p2) + (p2))
                    lhs350_ = globalState
                    lhs351_ = globalState
                    lhs350_.f2 = rhs412_
                    lhs351_.f2 = rhs413_
                    d_2350_v1_: D1
                    d_2350_v1_ = D1_DC3(p1, p1, (0) - (p1), p1, p1)
                    d_2351_v2_: _dafny.Map
                    d_2351_v2_ = _dafny.Map({p1: d_2350_v1_})
                    d_2351_v2_ = (d_2351_v2_).set(default__.safeModuloInt(p1, p1), d_2350_v1_)
                    source23_ = d_2350_v1_
                    if source23_.is_DC3:
                        d_2352___mcc_h0_ = source23_.cf7
                        d_2353___mcc_h1_ = source23_.cf8
                        d_2354___mcc_h2_ = source23_.cf9
                        d_2355___mcc_h3_ = source23_.cf10
                        d_2356___mcc_h4_ = source23_.cf11
                        d_2357_cf11_ = d_2356___mcc_h4_
                        d_2358_cf10_ = d_2355___mcc_h3_
                        d_2359_cf9_ = d_2354___mcc_h2_
                        d_2360_cf8_ = d_2353___mcc_h1_
                        d_2361_cf7_ = d_2352___mcc_h0_
                        d_2362_v3_: _dafny.Set
                        d_2362_v3_ = _dafny.Set({(self).f18})
                        d_2363_v4_: _dafny.Seq
                        d_2363_v4_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(p1, d_2360_cf8_), (d_2358_cf10_ if True else len(d_2362_v3_)), d_2358_cf10_])
                        d_2364_v5_: _dafny.Map
                        d_2364_v5_ = _dafny.Map({(self).f18: p1})
                        d_2365_v6_: D2
                        d_2365_v6_ = D2_DC6(d_2364_v5_)
                        (globalState).f3 = (d_2363_v4_)[default__.safeIndex((d_2360_cf8_) * (default__.fm1(len((d_2365_v6_).cf14), (self).f18, (self).f18, globalState)), len(d_2363_v4_))]
                        (globalState).f2 = (len(p0)) != (d_2358_cf10_)
                        d_2366_v7_: str
                        d_2366_v7_ = _dafny.CodePoint('r')
                        d_2367_v8_: _dafny.Set
                        d_2367_v8_ = _dafny.Set({d_2366_v7_})
                        d_2368_v9_: D23
                        d_2368_v9_ = D23_DC60(d_2367_v8_)
                        d_2369_v10_: C7
                        nw405_ = C7()
                        nw405_.ctor__(len(((d_2368_v9_).cf98) | (d_2367_v8_)), (self).f18)
                        d_2369_v10_ = nw405_
                        (globalState).f15 = default__.safeDivisionInt(default__.safeDivisionInt(d_2361_cf7_, d_2357_cf11_), default__.safeDivisionInt((0) - (p1), p1))
                    elif source23_.is_DC4:
                        d_2370___mcc_h5_ = source23_.cf12
                        d_2371_cf12_ = d_2370___mcc_h5_
                        d_2372_v11_: _dafny.Array
                        nw406_ = _dafny.Array(None, 19)
                        nw406_[int(0)] = p2
                        nw406_[int(1)] = p0
                        nw406_[int(2)] = p0
                        nw406_[int(3)] = p0
                        nw406_[int(4)] = p2
                        nw406_[int(5)] = (p2) + (p0)
                        nw406_[int(6)] = p2
                        nw406_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2373_i1_ in range(default__.abs(-357))])
                        nw406_[int(8)] = p2
                        nw406_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bm"))
                        nw406_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvgq"))) + (p0)
                        nw406_[int(11)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2374_i2_ in range(default__.abs(-59))])
                        nw406_[int(12)] = p0
                        nw406_[int(13)] = p0
                        nw406_[int(14)] = (p0) + (p2)
                        nw406_[int(15)] = (p2) + (p2)
                        nw406_[int(16)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2375_i3_ in range(default__.abs(349))])
                        nw406_[int(17)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2376_i4_ in range(default__.abs(-196))])
                        nw406_[int(18)] = ((p2) + (p0)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2377_i5_ in range(default__.abs(-781))])), len((p2) + (p0))), _dafny.CodePoint('i'))
                        d_2372_v11_ = nw406_
                        index395_ = default__.safeIndex(946, (d_2372_v11_).length(0))
                        (d_2372_v11_)[index395_] = p2
                        d_2378_v12_: _dafny.Set
                        d_2378_v12_ = _dafny.Set({p1, p1})
                        d_2379_v13_: _dafny.Set
                        d_2379_v13_ = _dafny.Set({d_2378_v12_, d_2378_v12_})
                        d_2380_v14_: _dafny.Map
                        d_2380_v14_ = _dafny.Map({d_2371_cf12_: d_2379_v13_})
                        d_2381_v15_: _dafny.Seq
                        d_2381_v15_ = _dafny.SeqWithoutIsStrInference([d_2378_v12_])
                        d_2382_v17_: _dafny.Seq
                        def iife192_():
                            coll60_ = _dafny.Set()
                            compr_60_: _dafny.Set
                            for compr_60_ in (d_2381_v15_).Elements:
                                d_2383_v16_: _dafny.Set = compr_60_
                                if (d_2383_v16_) in (d_2381_v15_):
                                    coll60_ = coll60_.union(_dafny.Set([d_2383_v16_]))
                            return _dafny.Set(coll60_)
                        d_2382_v17_ = _dafny.SeqWithoutIsStrInference([iife192_()
                        ])
                        d_2384_v18_: D24
                        d_2384_v18_ = D24_DC63(d_2379_v13_)
                        d_2385_v19_: _dafny.Seq
                        d_2385_v19_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_2386_v20_: _dafny.Seq
                        d_2386_v20_ = _dafny.SeqWithoutIsStrInference([p1, (d_2385_v19_)[default__.safeIndex(len(d_2378_v12_), len(d_2385_v19_))]])
                        d_2387_v22_: _dafny.Array
                        nw407_ = _dafny.Array(None, 17)
                        nw407_[int(0)] = d_2379_v13_
                        nw407_[int(1)] = ((d_2380_v14_)[not((self).f18)] if (not((self).f18)) in (d_2380_v14_) else _dafny.Set({default__.fm49(p1, p1, p1, globalState)}))
                        nw407_[int(2)] = (d_2379_v13_ if (self).f18 else (d_2382_v17_)[default__.safeIndex(p1, len(d_2382_v17_))])
                        nw407_[int(3)] = d_2379_v13_
                        nw407_[int(4)] = (d_2384_v18_).cf104
                        nw407_[int(5)] = d_2379_v13_
                        nw407_[int(6)] = d_2379_v13_
                        nw407_[int(7)] = d_2379_v13_
                        def iife193_():
                            coll61_ = _dafny.Set()
                            compr_61_: int
                            for compr_61_ in (d_2386_v20_).Elements:
                                d_2388_v21_: int = compr_61_
                                if (d_2388_v21_) in (d_2386_v20_):
                                    coll61_ = coll61_.union(_dafny.Set([(d_2388_v21_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('q')])))]))
                            return _dafny.Set(coll61_)
                        nw407_[int(8)] = _dafny.Set({iife193_()
                        , d_2378_v12_})
                        nw407_[int(9)] = _dafny.Set({d_2378_v12_, _dafny.Set({p1})})
                        nw407_[int(10)] = d_2379_v13_
                        nw407_[int(11)] = d_2379_v13_
                        nw407_[int(12)] = (d_2379_v13_).intersection(d_2379_v13_)
                        nw407_[int(13)] = (d_2379_v13_) - (d_2379_v13_)
                        nw407_[int(14)] = (d_2379_v13_ if (self).f18 else d_2379_v13_)
                        nw407_[int(15)] = d_2379_v13_
                        nw407_[int(16)] = d_2379_v13_
                        d_2387_v22_ = nw407_
                        index396_ = default__.safeIndex(873, (d_2387_v22_).length(0))
                        (d_2387_v22_)[index396_] = d_2379_v13_
                        index397_ = default__.safeIndex(946, (d_2372_v11_).length(0))
                        index398_ = default__.safeIndex(873, (d_2387_v22_).length(0))
                        rhs414_ = p0
                        rhs415_ = p1
                        rhs416_ = ((d_2379_v13_) | (d_2379_v13_) if (self).f18 else d_2379_v13_)
                        lhs352_ = d_2372_v11_
                        lhs353_ = default__.safeIndex(946, (d_2372_v11_).length(0))
                        lhs354_ = globalState
                        lhs355_ = d_2387_v22_
                        lhs356_ = default__.safeIndex(873, (d_2387_v22_).length(0))
                        lhs352_[lhs353_] = rhs414_
                        lhs354_.f3 = rhs415_
                        lhs355_[lhs356_] = rhs416_
                        (globalState).f15 = default__.safeDivisionInt((0) - (p1), p1)
                        d_2389_v23_: _dafny.Array
                        nw408_ = _dafny.Array(int(0), 5)
                        d_2389_v23_ = nw408_
                        d_2389_v23_ = d_2389_v23_
                        (globalState).f15 = p1
                    elif source23_.is_DC2:
                        d_2390___mcc_h6_ = source23_.cf6
                        d_2391_cf6_ = d_2390___mcc_h6_
                        d_2392_v24_: _dafny.Set
                        d_2392_v24_ = _dafny.Set({(d_2391_cf6_).f18})
                        d_2393_v25_: C5
                        nw409_ = C5()
                        nw409_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) + (p0), p2, (d_2392_v24_).ispropersubset(d_2392_v24_))
                        d_2393_v25_ = nw409_
                        (globalState).f3 = p1
                        d_2394_v26_: _dafny.Seq
                        d_2394_v26_ = _dafny.SeqWithoutIsStrInference([default__.fm36((d_2391_cf6_).f18, default__.fm1(p1, (d_2391_cf6_).f18, (d_2391_cf6_).f18, globalState), False, globalState), d_2392_v24_, d_2392_v24_])
                        d_2395_v27_: str
                        d_2395_v27_ = _dafny.CodePoint('n')
                        d_2396_v28_: _dafny.Map
                        d_2396_v28_ = _dafny.Map({p1: p1})
                        d_2397_v29_: _dafny.Map
                        d_2397_v29_ = _dafny.Map({False: d_2396_v28_})
                        d_2398_v30_: _dafny.MultiSet
                        d_2398_v30_ = _dafny.MultiSet([p1, 561, len(d_2397_v29_), p1])
                        d_2399_v31_: _dafny.Map
                        d_2399_v31_ = _dafny.Map({p1: True})
                        d_2400_v32_: _dafny.Map
                        d_2400_v32_ = _dafny.Map({(0) - ((self).fm5(d_2398_v30_, len(d_2399_v31_), globalState)): d_2395_v27_})
                        d_2401_v33_: _dafny.Array
                        nw410_ = _dafny.Array(None, 12)
                        nw410_[int(0)] = d_2395_v27_
                        nw410_[int(1)] = d_2395_v27_
                        nw410_[int(2)] = d_2395_v27_
                        nw410_[int(3)] = d_2395_v27_
                        nw410_[int(4)] = d_2395_v27_
                        nw410_[int(5)] = _dafny.CodePoint('k')
                        nw410_[int(6)] = d_2395_v27_
                        nw410_[int(7)] = d_2395_v27_
                        nw410_[int(8)] = d_2395_v27_
                        nw410_[int(9)] = d_2395_v27_
                        nw410_[int(10)] = d_2395_v27_
                        nw410_[int(11)] = ((d_2400_v32_)[p1] if (p1) in (d_2400_v32_) else d_2395_v27_)
                        d_2401_v33_ = nw410_
                        d_2402_v34_: _dafny.Map
                        d_2402_v34_ = _dafny.Map({p1: d_2401_v33_})
                        d_2403_v35_: _dafny.Map
                        d_2403_v35_ = _dafny.Map({len(d_2402_v34_): p1})
                        d_2394_v26_ = default__.fm62(p1, d_2403_v35_, globalState)
                        d_2404_v36_: _dafny.Seq
                        d_2404_v36_ = _dafny.SeqWithoutIsStrInference([self.f22])
                        d_2404_v36_ = d_2404_v36_
                    elif True:
                        d_2405___mcc_h7_ = source23_.cf13
                        d_2406_cf13_ = d_2405___mcc_h7_
                        d_2407_v37_: _dafny.MultiSet
                        d_2407_v37_ = _dafny.MultiSet([True, (self).f18, (self).f18])
                        d_2408_v38_: _dafny.MultiSet
                        d_2408_v38_ = d_2407_v37_
                        d_2409_v39_: _dafny.Seq
                        d_2409_v39_ = _dafny.SeqWithoutIsStrInference([p1, 539])
                        d_2410_v40_: _dafny.Seq
                        d_2410_v40_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f18])
                        d_2411_v41_: _dafny.MultiSet
                        d_2411_v41_ = _dafny.MultiSet([p1, len(d_2410_v40_)])
                        d_2412_v42_: D16
                        d_2412_v42_ = D16_DC38(d_2411_v41_)
                        d_2413_v43_: _dafny.Array
                        nw411_ = _dafny.Array(None, 15)
                        nw411_[int(0)] = (d_2407_v37_).set((self).f18, default__.abs(p1))
                        nw411_[int(1)] = d_2407_v37_
                        nw411_[int(2)] = d_2407_v37_
                        nw411_[int(3)] = (d_2407_v37_) | (_dafny.MultiSet([(self).f18, True, (self).f18, (self).f18, (self).f18]))
                        nw411_[int(4)] = d_2407_v37_
                        nw411_[int(5)] = d_2407_v37_
                        nw411_[int(6)] = _dafny.MultiSet([True])
                        nw411_[int(7)] = (d_2408_v38_)
                        nw411_[int(8)] = (_dafny.MultiSet([(self).f18, (self).f18])) - (d_2407_v37_)
                        nw411_[int(9)] = _dafny.MultiSet([default__.fm6(d_2409_v39_, (self).f18, (d_2412_v42_).cf66, p1, globalState)])
                        nw411_[int(10)] = d_2407_v37_
                        nw411_[int(11)] = (d_2407_v37_) - ((default__.fm32(globalState)).set((self).f18, default__.abs(213)))
                        nw411_[int(12)] = (default__.fm32(globalState)) - (_dafny.MultiSet(d_2410_v40_))
                        nw411_[int(13)] = d_2407_v37_
                        nw411_[int(14)] = d_2407_v37_
                        d_2413_v43_ = nw411_
                        index399_ = default__.safeIndex(377, (d_2413_v43_).length(0))
                        (d_2413_v43_)[index399_] = d_2407_v37_
                        d_2414_v44_: _dafny.Array
                        def lambda103_(d_2415_p1_):
                            def lambda104_(d_2416_i6_):
                                return default__.safeDivisionInt(d_2416_i6_, d_2415_p1_)

                            return lambda104_

                        init66_ = lambda103_(p1)
                        nw412_ = _dafny.Array(None, 9)
                        for i0_66_ in range(nw412_.length(0)):
                            nw412_[i0_66_] = init66_(i0_66_)
                        d_2414_v44_ = nw412_
                        index400_ = default__.safeIndex(978, (d_2414_v44_).length(0))
                        (d_2414_v44_)[index400_] = (0) - (p1)
                        d_2417_v45_: C0
                        nw413_ = C0()
                        nw413_.ctor__(p1, len(d_2409_v39_))
                        d_2417_v45_ = nw413_
                        d_2418_v46_: _dafny.Seq
                        d_2418_v46_ = _dafny.SeqWithoutIsStrInference([d_2417_v45_])
                        index401_ = default__.safeIndex(48, (d_2414_v44_).length(0))
                        (d_2414_v44_)[index401_] = (d_2417_v45_).f32
                        index402_ = default__.safeIndex(377, (d_2413_v43_).length(0))
                        index403_ = default__.safeIndex(978, (d_2414_v44_).length(0))
                        index404_ = default__.safeIndex(48, (d_2414_v44_).length(0))
                        rhs417_ = d_2407_v37_
                        rhs418_ = (d_2417_v45_).f31
                        rhs419_ = d_2418_v46_
                        rhs420_ = (self).f18
                        rhs421_ = (0) - (p1)
                        lhs357_ = d_2413_v43_
                        lhs358_ = default__.safeIndex(377, (d_2413_v43_).length(0))
                        lhs359_ = d_2414_v44_
                        lhs360_ = default__.safeIndex(978, (d_2414_v44_).length(0))
                        lhs361_ = globalState
                        lhs362_ = d_2414_v44_
                        lhs363_ = default__.safeIndex(48, (d_2414_v44_).length(0))
                        lhs357_[lhs358_] = rhs417_
                        lhs359_[lhs360_] = rhs418_
                        d_2418_v46_ = rhs419_
                        lhs361_.f2 = rhs420_
                        lhs362_[lhs363_] = rhs421_
                        (globalState).f3 = (d_2414_v44_)[default__.safeIndex(978, (d_2414_v44_).length(0))]
                        d_2419_v47_: C1
                        nw414_ = C1()
                        nw414_.ctor__((self).f18)
                        d_2419_v47_ = nw414_
                        d_2420_v48_: str
                        d_2420_v48_ = _dafny.CodePoint('c')
                        d_2421_v49_: _dafny.Map
                        d_2421_v49_ = _dafny.Map({d_2420_v48_: (self).f18})
                        d_2422_v50_: _dafny.Set
                        d_2422_v50_ = _dafny.Set({d_2421_v49_})
                        d_2423_v51_: _dafny.Map
                        d_2423_v51_ = _dafny.Map({d_2409_v39_: d_2422_v50_})
                        d_2423_v51_ = (d_2423_v51_).set(d_2409_v39_, (d_2422_v50_) - (d_2422_v50_))
                    (globalState).f3 = default__.safeDivisionInt(p1, p1)
                    pass
            pass
        d_2424_v52_: D17
        d_2424_v52_ = D17_DC41(_dafny.Map({p1: (self).f18}))
        d_2425_v53_: _dafny.Map
        d_2425_v53_ = _dafny.Map({p1: (self).f18})
        d_2426_v54_: _dafny.MultiSet
        d_2426_v54_ = _dafny.MultiSet([(0) - (p1)])
        d_2427_v55_: _dafny.Seq
        d_2427_v55_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2428_v56_: _dafny.Seq
        d_2428_v56_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_2427_v55_)), (0) - (p1)])
        d_2429_v57_: _dafny.Map
        d_2429_v57_ = _dafny.Map({172: p1})
        d_2430_v58_: _dafny.Seq
        d_2430_v58_ = _dafny.SeqWithoutIsStrInference([d_2426_v54_])
        d_2431_v59_: _dafny.Map
        d_2431_v59_ = _dafny.Map({(self).f18: (d_2427_v55_).set(default__.safeIndex(p1, len(d_2427_v55_)), 302)})
        pat_let_tv36_ = p1
        pat_let_tv37_ = d_2425_v53_
        d_2432_v60_: _dafny.Array
        nw415_ = _dafny.Array(None, 16)
        def iife195_(_pat_let67_0):
            def iife196_(d_2433_dt__update__tmp_h0_):
                def iife197_(_pat_let68_0):
                    def iife198_(d_2434_dt__update_hcf73_h0_):
                        return D17_DC41(d_2434_dt__update_hcf73_h0_)
                    return iife198_(_pat_let68_0)
                return iife197_(_dafny.Map({pat_let_tv36_: (self).f18}))
            return iife196_(_pat_let67_0)
        def iife194_(_pat_let66_0):
            def iife199_(d_2435_dt__update__tmp_h1_):
                def iife200_(_pat_let69_0):
                    def iife201_(d_2436_dt__update_hcf73_h1_):
                        return D17_DC41(d_2436_dt__update_hcf73_h1_)
                    return iife201_(_pat_let69_0)
                return iife200_(pat_let_tv37_)
            return iife199_(_pat_let66_0)
        nw415_[int(0)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1, len((iife194_(iife195_(d_2424_v52_))).cf73), len(p2)]))
        nw415_[int(1)] = _dafny.MultiSet([p1])
        nw415_[int(2)] = d_2426_v54_
        nw415_[int(3)] = _dafny.MultiSet([p1])
        nw415_[int(4)] = d_2426_v54_
        nw415_[int(5)] = default__.fm25((0) - (p1), globalState)
        nw415_[int(6)] = d_2426_v54_
        nw415_[int(7)] = d_2426_v54_
        nw415_[int(8)] = d_2426_v54_
        nw415_[int(9)] = (_dafny.MultiSet([(d_2427_v55_)[default__.safeIndex(p1, len(d_2427_v55_))], p1, p1]) if default__.fm6(d_2428_v56_, (self).f18, d_2426_v54_, len(d_2429_v57_), globalState) else d_2426_v54_)
        nw415_[int(10)] = (d_2430_v58_)[default__.safeIndex(p1, len(d_2430_v58_))]
        nw415_[int(11)] = (_dafny.MultiSet((((d_2431_v59_)[(self).f18] if ((self).f18) in (d_2431_v59_) else d_2427_v55_)).set(default__.safeIndex(p1, len(((d_2431_v59_)[(self).f18] if ((self).f18) in (d_2431_v59_) else d_2427_v55_))), p1))).intersection(d_2426_v54_)
        nw415_[int(12)] = _dafny.MultiSet([p1])
        nw415_[int(13)] = d_2426_v54_
        nw415_[int(14)] = (d_2426_v54_ if True else d_2426_v54_)
        nw415_[int(15)] = _dafny.MultiSet([(0) - (p1)])
        d_2432_v60_ = nw415_
        d_2437_v61_: str
        d_2437_v61_ = _dafny.CodePoint('y')
        d_2438_v62_: _dafny.MultiSet
        d_2438_v62_ = _dafny.MultiSet([d_2437_v61_])
        index405_ = default__.safeIndex(968, (d_2432_v60_).length(0))
        (d_2432_v60_)[index405_] = _dafny.MultiSet([((d_2438_v62_)[d_2437_v61_] if (d_2437_v61_) in (d_2438_v62_) else len(_dafny.Set({d_2437_v61_, d_2437_v61_}))), 108, 273, p1])
        index406_ = default__.safeIndex(968, (d_2432_v60_).length(0))
        (d_2432_v60_)[index406_] = (_dafny.MultiSet(d_2427_v55_)) | ((d_2426_v54_) - (d_2426_v54_))
        hi13_ = p1
        for d_2439_i7_ in range(p1, hi13_):
            d_2440_v63_: _dafny.Seq
            d_2440_v63_ = _dafny.SeqWithoutIsStrInference([d_2429_v57_])
            d_2441_v64_: C4
            nw416_ = C4()
            nw416_.ctor__(d_2437_v61_, d_2439_i7_, (p1) >= (d_2439_i7_), p2, d_2440_v63_)
            d_2441_v64_ = nw416_
            arr2_ = self.f22
            index407_ = default__.safeIndex(585, (self.f22).length(0))
            arr2_[index407_] = ((0) - ((d_2441_v64_).f30)) < (p1)
            arr3_ = self.f22
            index408_ = default__.safeIndex(585, (self.f22).length(0))
            arr3_[index408_] = not((d_2439_i7_) > (((d_2432_v60_)[default__.safeIndex(968, (d_2432_v60_).length(0))]).cardinality))
            (globalState).f3 = (p1) * (905)
            arr4_ = self.f22
            index409_ = default__.safeIndex(585, (self.f22).length(0))
            arr4_[index409_] = True
        d_2442_v65_: int
        d_2443_v66_: int
        d_2444_v67_: _dafny.Array
        d_2445_v68_: _dafny.Map
        out62_: int
        out63_: int
        out64_: _dafny.Array
        out65_: _dafny.Map
        out62_, out63_, out64_, out65_ = default__.m0(globalState)
        d_2442_v65_ = out62_
        d_2443_v66_ = out63_
        d_2444_v67_ = out64_
        d_2445_v68_ = out65_
        d_2446_v69_: _dafny.Map
        d_2446_v69_ = _dafny.Map({(p1) >= (d_2443_v66_): d_2437_v61_})
        d_2446_v69_ = _dafny.Map({(self).f18: d_2437_v61_})
        d_2447_v70_: _dafny.Seq
        d_2447_v70_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f18])) - (default__.fm32(globalState))])
        d_2447_v70_ = d_2447_v70_


class C10(T1, T0):
    def  __init__(self):
        self._f18: bool = False
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f20: _dafny.Seq = _dafny.Seq({})
        self._f21: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f21, f19, f20, f18):
        (self)._f21 = f21
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f18 = f18

    def fm2(self, p0, p1, p2, globalState):
        return True

    def fm3(self, p0, globalState):
        return (D0_DC1((self).f21, len((self).f19), ((self).f19).set(default__.safeIndex(207, len((self).f19)), _dafny.CodePoint('t')), (self).f18, -625)).cf1

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_2448_i0_: int
        d_2448_i0_ = 0
        with _dafny.label("16"):
            while (self).f21:
                with _dafny.c_label("16"):
                    if (d_2448_i0_) >= (100):
                        raise _dafny.Break("16")
                    d_2448_i0_ = (d_2448_i0_) + (1)
                    if p1:
                        (globalState).f15 = -738
                        d_2449_v0_: int
                        d_2449_v0_ = -840
                        d_2450_v1_: _dafny.Map
                        d_2450_v1_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2451_i1_ in range(default__.abs(357))])) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2452_i2_ in range(default__.abs(136))])): d_2449_v0_})
                        d_2450_v1_ = (d_2450_v1_).set(p2, d_2449_v0_)
                        d_2453_v2_: _dafny.Array
                        nw417_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_2453_v2_ = nw417_
                        d_2454_v3_: _dafny.Array
                        nw418_ = _dafny.Array(False, 20)
                        d_2454_v3_ = nw418_
                        index410_ = default__.safeIndex(904, (d_2453_v2_).length(0))
                        (d_2453_v2_)[index410_] = d_2454_v3_
                        d_2455_v4_: T0
                        nw419_ = C1()
                        nw419_.ctor__((self).f21)
                        d_2455_v4_ = nw419_
                        d_2456_v5_: D1
                        d_2456_v5_ = D1_DC2(d_2455_v4_)
                        pat_let_tv38_ = d_2455_v4_
                        d_2457_v6_: _dafny.MultiSet
                        def iife202_(_pat_let70_0):
                            def iife203_(d_2458_dt__update__tmp_h0_):
                                def iife204_(_pat_let71_0):
                                    def iife205_(d_2459_dt__update_hcf6_h0_):
                                        return D1_DC2(d_2459_dt__update_hcf6_h0_)
                                    return iife205_(_pat_let71_0)
                                return iife204_(pat_let_tv38_)
                            return iife203_(_pat_let70_0)
                        d_2457_v6_ = _dafny.MultiSet([(D1_DC2((iife202_(d_2456_v5_)).cf6)).cf6])
                        d_2460_v7_: _dafny.Map
                        d_2460_v7_ = _dafny.Map({(d_2457_v6_).cardinality: (0) - (d_2449_v0_)})
                        d_2461_v8_: _dafny.MultiSet
                        d_2461_v8_ = _dafny.MultiSet([len(d_2460_v7_)])
                        index411_ = default__.safeIndex(904, (d_2453_v2_).length(0))
                        rhs422_ = (0) - ((161) + (d_2449_v0_))
                        rhs423_ = d_2454_v3_
                        rhs424_ = (self).f18
                        rhs425_ = ((d_2461_v8_)[d_2449_v0_] if (d_2449_v0_) in (d_2461_v8_) else (52) * (d_2449_v0_))
                        lhs364_ = globalState
                        lhs365_ = d_2453_v2_
                        lhs366_ = default__.safeIndex(904, (d_2453_v2_).length(0))
                        lhs367_ = globalState
                        lhs364_.f15 = rhs422_
                        lhs365_[lhs366_] = rhs423_
                        lhs367_.f2 = rhs424_
                        d_2449_v0_ = rhs425_
                        d_2462_v9_: _dafny.MultiSet
                        d_2462_v9_ = _dafny.MultiSet([p2])
                        d_2463_v10_: _dafny.Map
                        d_2463_v10_ = _dafny.Map({(self).f19: p2})
                        d_2464_v11_: _dafny.Set
                        d_2464_v11_ = _dafny.Set({d_2449_v0_, d_2449_v0_})
                        d_2465_v12_: _dafny.Seq
                        d_2465_v12_ = _dafny.SeqWithoutIsStrInference([d_2464_v11_, d_2464_v11_, d_2464_v11_])
                        d_2466_v13_: _dafny.Map
                        d_2466_v13_ = _dafny.Map({(d_2455_v4_).f18: (d_2465_v12_)[default__.safeIndex(d_2449_v0_, len(d_2465_v12_))]})
                        (self).m3((default__.fm1(d_2449_v0_, p2, (d_2455_v4_).f18, globalState)) + (len(default__.fm30(_dafny.SeqWithoutIsStrInference([default__.fm32(globalState), d_2462_v9_]), d_2463_v10_, globalState))), d_2466_v13_, (p2 if p2 else (self).f18), globalState)
                        d_2467_v14_: _dafny.Map
                        d_2467_v14_ = _dafny.Map({(d_2449_v0_) - (d_2449_v0_): (self).f18})
                        d_2468_v15_: _dafny.Seq
                        d_2468_v15_ = _dafny.SeqWithoutIsStrInference([True, (d_2455_v4_).f18])
                        d_2469_v16_: str
                        d_2469_v16_ = _dafny.CodePoint('t')
                        d_2470_v17_: _dafny.Seq
                        d_2470_v17_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_2449_v0_, p2, p1, globalState), d_2449_v0_, d_2449_v0_, d_2449_v0_, len(_dafny.Map({(d_2455_v4_).f18: d_2469_v16_}))])
                        d_2471_v18_: D16
                        d_2471_v18_ = D16_DC38((d_2461_v8_).set(len(d_2470_v17_), default__.abs(len((self).f19))))
                        d_2472_v19_: _dafny.Map
                        d_2472_v19_ = _dafny.Map({d_2471_v18_: ((self).fm3(d_2449_v0_, globalState)) and ((self).f18)})
                        rhs426_ = (not(p1) if not(p2) else (p2) in (d_2468_v15_))
                        rhs427_ = (_dafny.Map({d_2449_v0_: (self).f18})).set(d_2449_v0_, (d_2455_v4_).f18)
                        rhs428_ = (self).f21
                        rhs429_ = len(d_2472_v19_)
                        rhs430_ = (d_2455_v4_).f18
                        lhs368_ = globalState
                        lhs369_ = globalState
                        lhs370_ = globalState
                        lhs371_ = globalState
                        lhs368_.f2 = rhs426_
                        d_2467_v14_ = rhs427_
                        lhs369_.f2 = rhs428_
                        lhs370_.f15 = rhs429_
                        lhs371_.f2 = rhs430_
                    elif True:
                        d_2473_v20_: int
                        d_2473_v20_ = 320
                        d_2474_v21_: _dafny.Map
                        d_2474_v21_ = _dafny.Map({(self).f21: p1})
                        d_2475_v22_: _dafny.Seq
                        d_2475_v22_ = _dafny.SeqWithoutIsStrInference([(self).f18, True, (self).f18, (self).fm3(d_2473_v20_, globalState), ((d_2474_v21_)[p1] if (p1) in (d_2474_v21_) else (self).f18)])
                        d_2476_v24_: _dafny.Map
                        d_2476_v24_ = _dafny.Map({d_2473_v20_: (self).f21})
                        def iife206_():
                            coll62_ = _dafny.Map()
                            compr_62_: int
                            for compr_62_ in (d_2476_v24_).keys.Elements:
                                d_2477_v23_: int = compr_62_
                                if (d_2477_v23_) in (d_2476_v24_):
                                    coll62_[(d_2477_v23_) * (d_2473_v20_)] = d_2473_v20_
                            return _dafny.Map(coll62_)
                        (globalState).f2 = not((len(d_2475_v22_)) not in (iife206_()
                        ))
                        (globalState).f2 = p2
                        d_2478_v25_: str
                        d_2478_v25_ = _dafny.CodePoint('m')
                        d_2479_v26_: _dafny.Array
                        nw420_ = _dafny.Array(None, 13)
                        nw420_[int(0)] = ((self).f19)[default__.safeIndex(len(d_2475_v22_), len((self).f19))]
                        nw420_[int(1)] = d_2478_v25_
                        nw420_[int(2)] = _dafny.CodePoint('e')
                        nw420_[int(3)] = d_2478_v25_
                        nw420_[int(4)] = d_2478_v25_
                        nw420_[int(5)] = default__.fm46(globalState)
                        nw420_[int(6)] = d_2478_v25_
                        nw420_[int(7)] = d_2478_v25_
                        nw420_[int(8)] = d_2478_v25_
                        nw420_[int(9)] = d_2478_v25_
                        nw420_[int(10)] = d_2478_v25_
                        nw420_[int(11)] = d_2478_v25_
                        nw420_[int(12)] = d_2478_v25_
                        d_2479_v26_ = nw420_
                        index412_ = default__.safeIndex(350, (d_2479_v26_).length(0))
                        (d_2479_v26_)[index412_] = d_2478_v25_
                        index413_ = default__.safeIndex(350, (d_2479_v26_).length(0))
                        (d_2479_v26_)[index413_] = (d_2478_v25_ if p1 else d_2478_v25_)
                        d_2480_v27_: _dafny.Map
                        d_2480_v27_ = _dafny.Map({False: d_2478_v25_})
                        d_2473_v20_ = default__.safeDivisionInt(d_2473_v20_, len(_dafny.Map({d_2473_v20_: ((d_2480_v27_)[not((self).fm3(d_2473_v20_, globalState))] if (not((self).fm3(d_2473_v20_, globalState))) in (d_2480_v27_) else d_2478_v25_)})))
                        d_2481_v28_: _dafny.Array
                        def lambda105_(d_2482_i3_):
                            return default__.safeModuloInt(d_2482_i3_, 408)

                        init67_ = lambda105_
                        nw421_ = _dafny.Array(None, 9)
                        for i0_67_ in range(nw421_.length(0)):
                            nw421_[i0_67_] = init67_(i0_67_)
                        d_2481_v28_ = nw421_
                        index414_ = default__.safeIndex(446, (d_2481_v28_).length(0))
                        (d_2481_v28_)[index414_] = d_2473_v20_
                        d_2483_v29_: _dafny.Seq
                        d_2483_v29_ = _dafny.SeqWithoutIsStrInference([d_2481_v28_])
                        d_2484_v30_: _dafny.Seq
                        d_2484_v30_ = _dafny.SeqWithoutIsStrInference([d_2473_v20_])
                        d_2485_v31_: C3
                        nw422_ = C3()
                        nw422_.ctor__(d_2484_v30_, (self).f19, (self).f20, (self).f18)
                        d_2485_v31_ = nw422_
                        d_2486_v32_: _dafny.Set
                        d_2486_v32_ = _dafny.Set({d_2485_v31_, d_2485_v31_, d_2485_v31_})
                        d_2487_v33_: _dafny.Set
                        d_2487_v33_ = _dafny.Set({d_2486_v32_})
                        index415_ = default__.safeIndex(446, (d_2481_v28_).length(0))
                        (d_2481_v28_)[index415_] = ((len(d_2483_v29_) if p2 else d_2473_v20_)) + (len(d_2487_v33_))
                    d_2488_v34_: _dafny.Array
                    def lambda106_(d_2489_p2_):
                        def lambda107_(d_2490_i4_):
                            return _dafny.Map({_dafny.CodePoint('o'): d_2489_p2_})

                        return lambda107_

                    init68_ = lambda106_(p2)
                    nw423_ = _dafny.Array(None, 24)
                    for i0_68_ in range(nw423_.length(0)):
                        nw423_[i0_68_] = init68_(i0_68_)
                    d_2488_v34_ = nw423_
                    d_2491_v35_: D3
                    d_2491_v35_ = D3_DC10(d_2488_v34_)
                    source24_ = d_2491_v35_
                    if source24_.is_DC11:
                        d_2492___mcc_h0_ = source24_.cf23
                        d_2493___mcc_h1_ = source24_.cf24
                        d_2494_cf24_ = d_2493___mcc_h1_
                        d_2495_cf23_ = d_2492___mcc_h0_
                        d_2496_v36_: _dafny.Seq
                        d_2496_v36_ = _dafny.SeqWithoutIsStrInference([not(d_2495_cf23_)])
                        d_2497_v37_: int
                        d_2497_v37_ = 884
                        d_2498_v38_: _dafny.Set
                        d_2498_v38_ = _dafny.Set({331, d_2497_v37_})
                        d_2499_v39_: C6
                        nw424_ = C6()
                        nw424_.ctor__((d_2496_v36_) + (_dafny.SeqWithoutIsStrInference([p2, p2])), d_2497_v37_, default__.fm23(d_2498_v38_, d_2497_v37_, globalState), (self).f19, (self).f20)
                        d_2499_v39_ = nw424_
                        d_2495_cf23_ = d_2495_cf23_
                        d_2500_v40_: T0
                        nw425_ = C8()
                        nw425_.ctor__(d_2499_v39_.f26, not(d_2494_cf24_))
                        d_2500_v40_ = nw425_
                        d_2501_v41_: _dafny.Seq
                        d_2501_v41_ = _dafny.SeqWithoutIsStrInference([d_2500_v40_])
                        d_2502_v42_: _dafny.Seq
                        d_2502_v42_ = _dafny.SeqWithoutIsStrInference([d_2499_v39_.f26])
                        d_2503_v43_: _dafny.MultiSet
                        d_2503_v43_ = _dafny.MultiSet([len(_dafny.Set({(d_2501_v41_)[default__.safeIndex(d_2499_v39_.f26, len(d_2501_v41_))]})), len(d_2502_v42_)])
                        d_2504_v44_: str
                        d_2504_v44_ = _dafny.CodePoint('p')
                        d_2505_v45_: _dafny.Set
                        d_2505_v45_ = _dafny.Set({(d_2500_v40_).f18, p1})
                        (globalState).f12 = (((self).f19).set(default__.safeIndex((807) + (((d_2503_v43_)[d_2497_v37_] if (d_2497_v37_) in (d_2503_v43_) else d_2499_v39_.f26)), len((self).f19)), d_2504_v44_)).set(default__.safeIndex((d_2497_v37_) - (len(d_2505_v45_)), len(((self).f19).set(default__.safeIndex((807) + (((d_2503_v43_)[d_2497_v37_] if (d_2497_v37_) in (d_2503_v43_) else d_2499_v39_.f26)), len((self).f19)), d_2504_v44_))), d_2504_v44_)
                        d_2506_v46_: _dafny.Array
                        def lambda108_(d_2507_p2_):
                            def lambda109_(d_2508_i5_):
                                return d_2507_p2_

                            return lambda109_

                        init69_ = lambda108_(p2)
                        nw426_ = _dafny.Array(None, 28)
                        for i0_69_ in range(nw426_.length(0)):
                            nw426_[i0_69_] = init69_(i0_69_)
                        d_2506_v46_ = nw426_
                        d_2506_v46_ = (d_2506_v46_ if d_2494_cf24_ else d_2506_v46_)
                    elif source24_.is_DC12:
                        d_2509___mcc_h2_ = source24_.cf25
                        d_2510___mcc_h3_ = source24_.cf26
                        d_2511___mcc_h4_ = source24_.cf27
                        d_2512___mcc_h5_ = source24_.cf28
                        d_2513___mcc_h6_ = source24_.cf29
                        d_2514_cf29_ = d_2513___mcc_h6_
                        d_2515_cf28_ = d_2512___mcc_h5_
                        d_2516_cf27_ = d_2511___mcc_h4_
                        d_2517_cf26_ = d_2510___mcc_h3_
                        d_2518_cf25_ = d_2509___mcc_h2_
                        d_2519_v47_: _dafny.Seq
                        d_2519_v47_ = _dafny.SeqWithoutIsStrInference([d_2516_cf27_])
                        d_2520_v48_: D9
                        d_2520_v48_ = D9_DC22(d_2519_v47_)
                        d_2521_v49_: _dafny.Map
                        d_2521_v49_ = _dafny.Map({((self).f19) + ((self).f19): d_2520_v48_})
                        d_2521_v49_ = (d_2521_v49_).set((self).f19, d_2520_v48_)
                        (globalState).f2 = (self).f18
                        d_2522_v50_: _dafny.Seq
                        d_2522_v50_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2491_v35_, d_2491_v35_, d_2491_v35_, d_2491_v35_})])
                        d_2523_v51_: T0
                        nw427_ = C8()
                        nw427_.ctor__(d_2516_cf27_, d_2517_cf26_)
                        d_2523_v51_ = nw427_
                        d_2524_v52_: _dafny.Map
                        d_2524_v52_ = _dafny.Map({_dafny.Map({d_2523_v51_: (0) - (d_2516_cf27_)}): d_2522_v50_})
                        d_2525_v53_: _dafny.Map
                        d_2525_v53_ = _dafny.Map({d_2523_v51_: d_2516_cf27_})
                        d_2526_v54_: str
                        d_2526_v54_ = _dafny.CodePoint('d')
                        d_2527_v55_: C4
                        nw428_ = C4()
                        nw428_.ctor__(d_2526_v54_, d_2516_cf27_, True, (self).f19, (self).f20)
                        d_2527_v55_ = nw428_
                        d_2528_v56_: _dafny.Set
                        d_2528_v56_ = _dafny.Set({d_2491_v35_})
                        d_2522_v50_ = ((d_2524_v52_)[d_2525_v53_] if (d_2525_v53_) in (d_2524_v52_) else (d_2522_v50_).set(default__.safeIndex(len(_dafny.Map({(_dafny.SeqWithoutIsStrInference([d_2527_v55_])).set(default__.safeIndex((d_2527_v55_).f30, len(_dafny.SeqWithoutIsStrInference([d_2527_v55_]))), d_2527_v55_): (d_2527_v55_).f30})), len(d_2522_v50_)), d_2528_v56_))
                        d_2529_v57_: _dafny.Map
                        d_2529_v57_ = _dafny.Map({(self).f21: d_2516_cf27_})
                        d_2530_v58_: D3
                        d_2530_v58_ = D3_DC13((d_2527_v55_).f30, (self).f19, (d_2527_v55_).f30, -495, ((self).f19)[default__.safeIndex(((d_2529_v57_)[True] if (True) in (d_2529_v57_) else (d_2527_v55_).f30), len((self).f19))])
                        d_2531_v59_: _dafny.Map
                        d_2531_v59_ = _dafny.Map({p1: (d_2530_v58_).cf30})
                        d_2531_v59_ = (d_2531_v59_).set((self).f18, (d_2527_v55_).f30)
                    elif source24_.is_DC13:
                        d_2532___mcc_h7_ = source24_.cf30
                        d_2533___mcc_h8_ = source24_.cf31
                        d_2534___mcc_h9_ = source24_.cf32
                        d_2535___mcc_h10_ = source24_.cf33
                        d_2536___mcc_h11_ = source24_.cf34
                        d_2537_cf34_ = d_2536___mcc_h11_
                        d_2538_cf33_ = d_2535___mcc_h10_
                        d_2539_cf32_ = d_2534___mcc_h9_
                        d_2540_cf31_ = d_2533___mcc_h8_
                        d_2541_cf30_ = d_2532___mcc_h7_
                        d_2542_v60_: _dafny.Array
                        nw429_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                        d_2542_v60_ = nw429_
                        d_2543_v61_: D13
                        d_2543_v61_ = D13_DC33((self).f21, 195, p1, len(_dafny.Map({p1: p1})))
                        d_2544_v62_: _dafny.Seq
                        d_2544_v62_ = _dafny.SeqWithoutIsStrInference([(d_2543_v61_).cf55])
                        d_2545_v63_: _dafny.Map
                        d_2545_v63_ = _dafny.Map({d_2542_v60_: d_2544_v62_})
                        d_2545_v63_ = ((d_2545_v63_).set(d_2542_v60_, d_2544_v62_)) | (d_2545_v63_)
                        d_2546_v64_: _dafny.Set
                        d_2546_v64_ = _dafny.Set({(self).f18})
                        d_2547_v65_: _dafny.Map
                        d_2547_v65_ = _dafny.Map({(self).f21: d_2538_cf33_})
                        d_2548_v66_: _dafny.Map
                        d_2548_v66_ = _dafny.Map({d_2539_cf32_: len(d_2547_v65_)})
                        d_2549_v67_: _dafny.Array
                        nw430_ = _dafny.Array(None, 21)
                        nw430_[int(0)] = 482
                        nw430_[int(1)] = d_2541_cf30_
                        nw430_[int(2)] = d_2538_cf33_
                        nw430_[int(3)] = d_2539_cf32_
                        nw430_[int(4)] = d_2539_cf32_
                        nw430_[int(5)] = d_2541_cf30_
                        nw430_[int(6)] = 746
                        nw430_[int(7)] = len(d_2546_v64_)
                        nw430_[int(8)] = -278
                        nw430_[int(9)] = d_2539_cf32_
                        nw430_[int(10)] = d_2539_cf32_
                        nw430_[int(11)] = d_2541_cf30_
                        nw430_[int(12)] = d_2539_cf32_
                        nw430_[int(13)] = len(d_2548_v66_)
                        nw430_[int(14)] = d_2538_cf33_
                        nw430_[int(15)] = len((self).f19)
                        nw430_[int(16)] = d_2539_cf32_
                        nw430_[int(17)] = d_2539_cf32_
                        nw430_[int(18)] = d_2538_cf33_
                        nw430_[int(19)] = len(d_2544_v62_)
                        nw430_[int(20)] = d_2538_cf33_
                        d_2549_v67_ = nw430_
                        d_2550_v68_: _dafny.Array
                        nw431_ = _dafny.Array(None, 1)
                        nw431_[int(0)] = d_2549_v67_
                        d_2550_v68_ = nw431_
                        d_2551_v69_: _dafny.Map
                        d_2551_v69_ = _dafny.Map({d_2539_cf32_: p2})
                        d_2552_v70_: _dafny.Map
                        d_2552_v70_ = _dafny.Map({(self).f21: d_2551_v69_})
                        d_2553_v71_: _dafny.Map
                        d_2553_v71_ = _dafny.Map({len((d_2551_v69_) | (((d_2552_v70_)[p2] if (p2) in (d_2552_v70_) else d_2551_v69_))): d_2550_v68_})
                        rhs431_ = ((d_2553_v71_)[(default__.fm1(d_2539_cf32_, (self).f18, p1, globalState)) - (d_2541_cf30_)] if ((default__.fm1(d_2539_cf32_, (self).f18, p1, globalState)) - (d_2541_cf30_)) in (d_2553_v71_) else d_2550_v68_)
                        rhs432_ = ((d_2539_cf32_) + ((0) - (d_2539_cf32_))) - (d_2538_cf33_)
                        lhs372_ = globalState
                        d_2550_v68_ = rhs431_
                        lhs372_.f3 = rhs432_
                        d_2554_v72_: _dafny.MultiSet
                        d_2554_v72_ = _dafny.MultiSet([p2])
                        d_2555_v73_: _dafny.MultiSet
                        d_2555_v73_ = _dafny.MultiSet([d_2554_v72_, d_2554_v72_, _dafny.MultiSet((d_2544_v62_) + (d_2544_v62_))])
                        d_2555_v73_ = d_2555_v73_
                        d_2556_v74_: _dafny.Set
                        d_2556_v74_ = _dafny.Set({d_2547_v65_})
                        d_2539_cf32_ = len(d_2556_v74_)
                    elif source24_.is_DC10:
                        d_2557___mcc_h12_ = source24_.cf22
                        d_2558_cf22_ = d_2557___mcc_h12_
                        d_2559_v75_: int
                        d_2559_v75_ = 14
                        d_2560_v76_: str
                        d_2560_v76_ = _dafny.CodePoint('e')
                        d_2561_v77_: _dafny.Array
                        nw432_ = _dafny.Array(_dafny.Map({}), 3)
                        d_2561_v77_ = nw432_
                        d_2562_v78_: _dafny.Map
                        d_2562_v78_ = _dafny.Map({default__.safeDivisionInt(d_2559_v75_, len(((self).f19).set(default__.safeIndex(d_2559_v75_, len((self).f19)), d_2560_v76_))): (d_2561_v77_ if (self).f21 else d_2561_v77_)})
                        d_2562_v78_ = (d_2562_v78_).set(d_2559_v75_, d_2561_v77_)
                        d_2563_v79_: _dafny.MultiSet
                        d_2563_v79_ = _dafny.MultiSet([d_2559_v75_])
                        d_2564_v80_: _dafny.Map
                        d_2564_v80_ = _dafny.Map({(d_2563_v79_).cardinality: d_2559_v75_})
                        d_2565_v81_: C2
                        nw433_ = C2()
                        nw433_.ctor__(d_2559_v75_, _dafny.SeqWithoutIsStrInference([len(d_2564_v80_)]), p1)
                        d_2565_v81_ = nw433_
                        d_2565_v81_ = d_2565_v81_
                        d_2566_v82_: _dafny.Seq
                        d_2566_v82_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f19])
                        d_2567_v83_: _dafny.Set
                        d_2567_v83_ = _dafny.Set({p2})
                        d_2568_v84_: C4
                        nw434_ = C4()
                        nw434_.ctor__(d_2560_v76_, d_2559_v75_, p2, (d_2566_v82_)[default__.safeIndex(len(d_2567_v83_), len(d_2566_v82_))], (self).f20)
                        d_2568_v84_ = nw434_
                        d_2569_v85_: _dafny.Map
                        d_2569_v85_ = _dafny.Map({False: d_2568_v84_})
                        d_2570_v86_: _dafny.Seq
                        d_2570_v86_ = _dafny.SeqWithoutIsStrInference([d_2568_v84_])
                        d_2571_v87_: _dafny.Seq
                        d_2571_v87_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                        d_2572_v88_: _dafny.Array
                        nw435_ = _dafny.Array(None, 24)
                        nw435_[int(0)] = d_2568_v84_
                        nw435_[int(1)] = d_2568_v84_
                        nw435_[int(2)] = d_2568_v84_
                        nw435_[int(3)] = d_2568_v84_
                        nw435_[int(4)] = ((d_2569_v85_)[False] if (False) in (d_2569_v85_) else d_2568_v84_)
                        nw435_[int(5)] = d_2568_v84_
                        nw435_[int(6)] = d_2568_v84_
                        nw435_[int(7)] = d_2568_v84_
                        nw435_[int(8)] = d_2568_v84_
                        nw435_[int(9)] = d_2568_v84_
                        nw435_[int(10)] = d_2568_v84_
                        nw435_[int(11)] = d_2568_v84_
                        nw435_[int(12)] = d_2568_v84_
                        nw435_[int(13)] = (d_2570_v86_)[default__.safeIndex(d_2559_v75_, len(d_2570_v86_))]
                        nw435_[int(14)] = d_2568_v84_
                        nw435_[int(15)] = d_2568_v84_
                        nw435_[int(16)] = d_2568_v84_
                        nw435_[int(17)] = (d_2568_v84_ if (self).f18 else d_2568_v84_)
                        nw435_[int(18)] = d_2568_v84_
                        nw435_[int(19)] = d_2568_v84_
                        nw435_[int(20)] = (d_2568_v84_ if (self).fm3(len(d_2571_v87_), globalState) else d_2568_v84_)
                        nw435_[int(21)] = d_2568_v84_
                        nw435_[int(22)] = d_2568_v84_
                        nw435_[int(23)] = d_2568_v84_
                        d_2572_v88_ = nw435_
                        index416_ = default__.safeIndex(933, (d_2572_v88_).length(0))
                        (d_2572_v88_)[index416_] = d_2568_v84_
                        index417_ = default__.safeIndex(933, (d_2572_v88_).length(0))
                        (d_2572_v88_)[index417_] = d_2568_v84_
                        (d_2568_v84_).m1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atpvbkx")), (0) - (-305), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ib")), globalState)
                    elif True:
                        d_2573___mcc_h13_ = source24_.cf35
                        d_2574_cf35_ = d_2573___mcc_h13_
                        (globalState).f2 = p1
                        (globalState).f2 = (self).f21
                        d_2575_v89_: int
                        d_2575_v89_ = 638
                        (globalState).f0 = ((0) - (d_2575_v89_) if not(p1) else (0) - ((d_2575_v89_) - (d_2575_v89_)))
                        d_2576_v90_: _dafny.Set
                        d_2576_v90_ = _dafny.Set({len((self).f19), d_2575_v89_, d_2575_v89_})
                        d_2577_v91_: _dafny.Seq
                        d_2577_v91_ = _dafny.SeqWithoutIsStrInference([len(d_2576_v90_)])
                        d_2578_v92_: _dafny.Set
                        d_2578_v92_ = _dafny.Set({(self).f21})
                        d_2579_v93_: D10
                        d_2579_v93_ = D10_DC24(d_2578_v92_)
                        d_2580_v94_: D13
                        d_2580_v94_ = D13_DC33(False, d_2575_v89_, not((self).f21), d_2575_v89_)
                        d_2581_v95_: C3
                        nw436_ = C3()
                        nw436_.ctor__(d_2577_v91_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akrqthbvy")), default__.fm51(p2, d_2575_v89_, p1, d_2579_v93_, globalState), (d_2580_v94_).cf57)
                        d_2581_v95_ = nw436_
                    d_2582_v96_: _dafny.MultiSet
                    d_2582_v96_ = _dafny.MultiSet([-876])
                    d_2583_v97_: int
                    d_2583_v97_ = 259
                    (globalState).f0 = (((d_2582_v96_)[d_2583_v97_] if (d_2583_v97_) in (d_2582_v96_) else d_2583_v97_)) * ((d_2583_v97_) + (d_2583_v97_))
                    d_2584_v98_: _dafny.Set
                    d_2584_v98_ = _dafny.Set({not(p2), (self).f18})
                    d_2585_v99_: str
                    d_2585_v99_ = _dafny.CodePoint('b')
                    d_2586_v100_: _dafny.Seq
                    d_2586_v100_ = _dafny.SeqWithoutIsStrInference([(self).f18, p2])
                    d_2587_v101_: _dafny.Map
                    d_2587_v101_ = _dafny.Map({d_2586_v100_: -758})
                    d_2588_v102_: _dafny.Map
                    d_2588_v102_ = _dafny.Map({(self).f21: (self).f19})
                    d_2589_v103_: _dafny.Map
                    d_2589_v103_ = _dafny.Map({d_2583_v97_: len(((d_2588_v102_)[(self).f18] if ((self).f18) in (d_2588_v102_) else (self).f19))})
                    d_2590_v104_: _dafny.Set
                    d_2590_v104_ = _dafny.Set({-682})
                    d_2591_v105_: _dafny.Map
                    d_2591_v105_ = _dafny.Map({(self).f18: d_2583_v97_})
                    d_2592_v106_: _dafny.MultiSet
                    d_2592_v106_ = _dafny.MultiSet([default__.fm23(d_2590_v104_, len(d_2591_v105_), globalState), True, p2, p2, True])
                    d_2593_v108_: _dafny.MultiSet
                    d_2593_v108_ = _dafny.MultiSet([d_2585_v99_])
                    d_2594_v109_: D0
                    d_2594_v109_ = D0_DC0(False)
                    d_2595_v110_: _dafny.Seq
                    d_2595_v110_ = _dafny.SeqWithoutIsStrInference([-670, d_2583_v97_])
                    d_2596_v111_: _dafny.Seq
                    d_2596_v111_ = _dafny.SeqWithoutIsStrInference([d_2583_v97_, (d_2595_v110_)[default__.safeIndex(d_2583_v97_, len(d_2595_v110_))], len((self).f19), (0) - (d_2583_v97_)])
                    d_2597_v112_: C2
                    nw437_ = C2()
                    nw437_.ctor__(d_2583_v97_, d_2595_v110_, p2)
                    d_2597_v112_ = nw437_
                    d_2598_v113_: _dafny.Map
                    d_2598_v113_ = _dafny.Map({(d_2594_v109_).cf0: d_2597_v112_})
                    d_2599_v114_: _dafny.Array
                    nw438_ = _dafny.Array(None, 23)
                    nw438_[int(0)] = d_2583_v97_
                    nw438_[int(1)] = d_2583_v97_
                    nw438_[int(2)] = (len(d_2584_v98_)) + ((_dafny.MultiSet([d_2585_v99_])).cardinality)
                    nw438_[int(3)] = ((d_2587_v101_)[_dafny.SeqWithoutIsStrInference([p1])] if (_dafny.SeqWithoutIsStrInference([p1])) in (d_2587_v101_) else d_2583_v97_)
                    nw438_[int(4)] = default__.safeDivisionInt(d_2583_v97_, d_2583_v97_)
                    nw438_[int(5)] = d_2583_v97_
                    nw438_[int(6)] = d_2583_v97_
                    nw438_[int(7)] = d_2583_v97_
                    nw438_[int(8)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_2583_v97_, 111]): p1})) for d_2600_i6_ in range(default__.abs(888))]))
                    nw438_[int(9)] = default__.safeModuloInt(119, d_2583_v97_)
                    nw438_[int(10)] = len((self).f19)
                    nw438_[int(11)] = ((d_2589_v103_)[(0) - (d_2583_v97_)] if ((0) - (d_2583_v97_)) in (d_2589_v103_) else d_2583_v97_)
                    nw438_[int(12)] = (0) - ((d_2592_v106_).cardinality)
                    def iife207_():
                        coll63_ = _dafny.Map()
                        compr_63_: str
                        for compr_63_ in (d_2593_v108_).Elements:
                            d_2601_v107_: str = compr_63_
                            if (d_2601_v107_) in (d_2593_v108_):
                                coll63_[d_2601_v107_] = (self).f21
                        return _dafny.Map(coll63_)
                    nw438_[int(13)] = len((_dafny.Set({(0) - (len(iife207_()
                    )), len(d_2598_v113_), len((self).f19)})) | (d_2590_v104_))
                    nw438_[int(14)] = (0) - (default__.safeModuloInt(d_2597_v112_.f33, d_2597_v112_.f33))
                    nw438_[int(15)] = d_2597_v112_.f33
                    nw438_[int(16)] = 366
                    nw438_[int(17)] = d_2597_v112_.f33
                    nw438_[int(18)] = default__.safeDivisionInt(d_2597_v112_.f33, d_2583_v97_)
                    nw438_[int(19)] = d_2597_v112_.f33
                    nw438_[int(20)] = d_2583_v97_
                    nw438_[int(21)] = (len(d_2586_v100_)) + (d_2583_v97_)
                    nw438_[int(22)] = 291
                    d_2599_v114_ = nw438_
                    index418_ = default__.safeIndex(0, (d_2599_v114_).length(0))
                    (d_2599_v114_)[index418_] = (d_2583_v97_) + (d_2583_v97_)
                    d_2602_v115_: _dafny.Map
                    d_2602_v115_ = _dafny.Map({(self).f21: p1})
                    d_2603_v116_: _dafny.Map
                    d_2603_v116_ = _dafny.Map({d_2602_v115_: d_2583_v97_})
                    index419_ = default__.safeIndex(0, (d_2599_v114_).length(0))
                    (d_2599_v114_)[index419_] = ((d_2603_v116_)[(d_2602_v115_) | ((d_2602_v115_).set(default__.fm23(d_2590_v104_, d_2583_v97_, globalState), (self).f18))] if ((d_2602_v115_) | ((d_2602_v115_).set(default__.fm23(d_2590_v104_, d_2583_v97_, globalState), (self).f18))) in (d_2603_v116_) else d_2597_v112_.f33)
                    pass
            pass
        d_2604_v117_: D19
        d_2604_v117_ = D19_DC49(not((self).f21))
        d_2605_v118_: T0
        nw439_ = C1()
        nw439_.ctor__((d_2604_v117_).cf83)
        d_2605_v118_ = nw439_
        d_2606_v119_: _dafny.Set
        d_2606_v119_ = _dafny.Set({d_2605_v118_})
        d_2607_v120_: int
        d_2607_v120_ = -915
        (globalState).f0 = (0) - (default__.safeDivisionInt(len(d_2606_v119_), d_2607_v120_))
        if (d_2605_v118_).f18:
            d_2608_v122_: _dafny.Array
            def lambda110_(d_2609_v120_):
                def lambda111_(d_2610_i7_):
                    def iife208_():
                        coll64_ = _dafny.Set()
                        compr_64_: int
                        for compr_64_ in _dafny.IntegerRange(980, -303):
                            d_2611_v121_: int = compr_64_
                            if ((980) <= (d_2611_v121_)) and ((d_2611_v121_) < (-303)):
                                coll64_ = coll64_.union(_dafny.Set([default__.safeDivisionInt(d_2611_v121_, d_2609_v120_)]))
                        return _dafny.Set(coll64_)
                    return (d_2610_i7_) - (len(iife208_()
                    ))

                return lambda111_

            init70_ = lambda110_(d_2607_v120_)
            nw440_ = _dafny.Array(None, 18)
            for i0_70_ in range(nw440_.length(0)):
                nw440_[i0_70_] = init70_(i0_70_)
            d_2608_v122_ = nw440_
            index420_ = default__.safeIndex(897, (d_2608_v122_).length(0))
            (d_2608_v122_)[index420_] = d_2607_v120_
            d_2612_v123_: str
            d_2612_v123_ = _dafny.CodePoint('x')
            d_2613_v124_: _dafny.Map
            d_2613_v124_ = _dafny.Map({d_2612_v123_: (default__.fm63(globalState)) | (_dafny.Map({(d_2605_v118_).f18: d_2612_v123_}))})
            d_2614_v125_: D16
            d_2614_v125_ = D16_DC39((d_2605_v118_).f18, 553)
            index421_ = default__.safeIndex(897, (d_2608_v122_).length(0))
            def iife209_():
                coll65_ = _dafny.Map()
                compr_65_: str
                for compr_65_ in (_dafny.Map({d_2612_v123_: d_2612_v123_})).keys.Elements:
                    d_2615_v126_: str = compr_65_
                    if (d_2615_v126_) in (_dafny.Map({d_2612_v123_: d_2612_v123_})):
                        coll65_[d_2615_v126_] = _dafny.Map({(self).f18: d_2612_v123_})
                return _dafny.Map(coll65_)
            rhs433_ = d_2607_v120_
            rhs434_ = d_2607_v120_
            rhs435_ = (d_2614_v125_).cf67
            rhs436_ = iife209_()

            lhs373_ = d_2608_v122_
            lhs374_ = default__.safeIndex(897, (d_2608_v122_).length(0))
            lhs375_ = globalState
            lhs376_ = globalState
            lhs373_[lhs374_] = rhs433_
            lhs375_.f15 = rhs434_
            lhs376_.f2 = rhs435_
            d_2613_v124_ = rhs436_
            index422_ = default__.safeIndex(897, (d_2608_v122_).length(0))
            (d_2608_v122_)[index422_] = 703
            d_2616_v127_: D6
            d_2616_v127_ = D6_DC18()
            source25_ = d_2616_v127_
            if source25_.is_DC18:
                d_2617_v128_: _dafny.Array
                nw441_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_2617_v128_ = nw441_
                index423_ = default__.safeIndex(427, (d_2617_v128_).length(0))
                (d_2617_v128_)[index423_] = d_2608_v122_
                d_2618_v129_: _dafny.Seq
                d_2618_v129_ = _dafny.SeqWithoutIsStrInference([(d_2605_v118_).f18, (self).f18])
                d_2619_v130_: _dafny.Seq
                d_2619_v130_ = _dafny.SeqWithoutIsStrInference([d_2618_v129_])
                d_2620_v131_: _dafny.Map
                d_2620_v131_ = _dafny.Map({d_2619_v130_: d_2608_v122_})
                d_2621_v132_: _dafny.Array
                d_2621_v132_ = d_2608_v122_
                index424_ = default__.safeIndex(427, (d_2617_v128_).length(0))
                (d_2617_v128_)[index424_] = ((d_2620_v131_)[d_2619_v130_] if (d_2619_v130_) in (d_2620_v131_) else (d_2621_v132_))
                d_2622_v133_: _dafny.Array
                nw442_ = _dafny.Array(None, 12)
                d_2622_v133_ = nw442_
                d_2623_v134_: _dafny.Seq
                d_2623_v134_ = _dafny.SeqWithoutIsStrInference([d_2622_v133_, d_2622_v133_])
                d_2624_v135_: _dafny.Set
                d_2624_v135_ = _dafny.Set({len(d_2623_v134_)})
                d_2625_v136_: _dafny.Map
                d_2625_v136_ = _dafny.Map({False: (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))]})
                index425_ = default__.safeIndex(897, (d_2608_v122_).length(0))
                rhs437_ = False
                rhs438_ = (d_2624_v135_).intersection((_dafny.Set({d_2607_v120_})) | (_dafny.Set({len(d_2625_v136_)})))
                rhs439_ = p2
                rhs440_ = (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))]
                lhs377_ = globalState
                lhs378_ = globalState
                lhs379_ = d_2608_v122_
                lhs380_ = default__.safeIndex(897, (d_2608_v122_).length(0))
                lhs377_.f2 = rhs437_
                d_2624_v135_ = rhs438_
                lhs378_.f2 = rhs439_
                lhs379_[lhs380_] = rhs440_
                d_2626_v137_: _dafny.Map
                d_2626_v137_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([p2])).set(default__.safeIndex(d_2607_v120_, len(_dafny.SeqWithoutIsStrInference([p2]))), p1)): (self).f21})
                d_2626_v137_ = d_2626_v137_
                d_2627_v138_: _dafny.Seq
                d_2627_v138_ = _dafny.SeqWithoutIsStrInference([d_2607_v120_, 496])
                d_2628_v139_: _dafny.Map
                d_2628_v139_ = _dafny.Map({(d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))]: d_2607_v120_})
                d_2629_v140_: _dafny.MultiSet
                d_2629_v140_ = _dafny.MultiSet([True])
                d_2630_v141_: _dafny.Map
                d_2630_v141_ = _dafny.Map({d_2629_v140_: (d_2605_v118_).f18})
                d_2631_v142_: _dafny.Map
                d_2631_v142_ = _dafny.Map({len(d_2630_v141_): _dafny.Set({len(_dafny.Map({_dafny.Map({not(p1): (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))]}): d_2612_v123_}))})})
                d_2632_v143_: C5
                nw443_ = C5()
                nw443_.ctor__((self).f19, (self).f19, (d_2605_v118_).f18)
                d_2632_v143_ = nw443_
                d_2633_v144_: _dafny.MultiSet
                d_2633_v144_ = _dafny.MultiSet([(d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))]])
                d_2634_v145_: D24
                d_2634_v145_ = D24_DC64(default__.fm23(((d_2631_v142_)[d_2607_v120_] if (d_2607_v120_) in (d_2631_v142_) else d_2624_v135_), len(d_2618_v129_), globalState), (self).f21, d_2632_v143_, d_2607_v120_, (d_2633_v144_).cardinality)
                d_2635_v146_: C3
                nw444_ = C3()
                nw444_.ctor__(d_2627_v138_, default__.fm48(globalState), ((self).f20) + (_dafny.SeqWithoutIsStrInference([d_2628_v139_])), (d_2634_v145_).cf106)
                d_2635_v146_ = nw444_
            elif True:
                d_2636___mcc_h14_ = source25_.cf38
                d_2637_cf38_ = d_2636___mcc_h14_
                d_2638_v147_: D22
                d_2638_v147_ = D22_DC58((d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))])
                d_2638_v147_ = d_2638_v147_
                d_2639_v148_: _dafny.Array
                nw445_ = _dafny.Array(None, 24)
                d_2639_v148_ = nw445_
                d_2640_v149_: C5
                nw446_ = C5()
                nw446_.ctor__(((self).f19).set(default__.safeIndex((d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))], len((self).f19)), d_2612_v123_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2641_i8_ in range(default__.abs(957))]), (d_2605_v118_).f18)
                d_2640_v149_ = nw446_
                index426_ = default__.safeIndex(293, (d_2639_v148_).length(0))
                (d_2639_v148_)[index426_] = (D24_DC64(not((self).f18), (d_2605_v118_).f18, d_2640_v149_, (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))], (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))])).cf107
                index427_ = default__.safeIndex(293, (d_2639_v148_).length(0))
                (d_2639_v148_)[index427_] = d_2640_v149_
                d_2642_v150_: _dafny.Array
                nw447_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_2642_v150_ = nw447_
                d_2643_v151_: _dafny.Set
                d_2643_v151_ = _dafny.Set({(d_2639_v148_)[default__.safeIndex(293, (d_2639_v148_).length(0))]})
                d_2644_v152_: _dafny.Seq
                d_2644_v152_ = _dafny.SeqWithoutIsStrInference([(d_2639_v148_)[default__.safeIndex(293, (d_2639_v148_).length(0))]])
                d_2645_v153_: _dafny.MultiSet
                d_2645_v153_ = _dafny.MultiSet([d_2643_v151_, _dafny.Set({(d_2644_v152_)[default__.safeIndex((d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))], len(d_2644_v152_))]}), d_2643_v151_])
                index428_ = default__.safeIndex(166, (d_2642_v150_).length(0))
                (d_2642_v150_)[index428_] = d_2645_v153_
                d_2646_v154_: _dafny.MultiSet
                d_2646_v154_ = _dafny.MultiSet([p1])
                d_2647_v155_: _dafny.Map
                d_2647_v155_ = _dafny.Map({((d_2646_v154_)[not(False)] if (not(False)) in (d_2646_v154_) else (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))]): d_2643_v151_})
                d_2648_v156_: _dafny.Seq
                d_2648_v156_ = _dafny.SeqWithoutIsStrInference([d_2643_v151_, d_2643_v151_, d_2643_v151_, _dafny.Set({(d_2639_v148_)[default__.safeIndex(293, (d_2639_v148_).length(0))], d_2640_v149_})])
                d_2649_v157_: D24
                d_2649_v157_ = D24_DC64(p2, not((self).f18), d_2640_v149_, (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))], (d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))])
                d_2650_v158_: _dafny.Seq
                d_2650_v158_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([((d_2647_v155_)[d_2607_v120_] if (d_2607_v120_) in (d_2647_v155_) else d_2643_v151_)]), (d_2648_v156_) + (d_2648_v156_), _dafny.SeqWithoutIsStrInference([d_2643_v151_, _dafny.Set({(d_2649_v157_).cf107, d_2640_v149_})])])
                index429_ = default__.safeIndex(166, (d_2642_v150_).length(0))
                (d_2642_v150_)[index429_] = _dafny.MultiSet((d_2650_v158_)[default__.safeIndex(d_2607_v120_, len(d_2650_v158_))])
                d_2607_v120_ = (0) - (len(((self).f19) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csklyy")))))
            (globalState).f2 = (self).fm3(((d_2608_v122_)[default__.safeIndex(897, (d_2608_v122_).length(0))] if not((self).f18) else d_2607_v120_), globalState)
            d_2651_v159_: _dafny.MultiSet
            d_2651_v159_ = _dafny.MultiSet([not((self).f18)])
            rhs441_ = not((self).f21)
            rhs442_ = d_2651_v159_
            lhs381_ = globalState
            lhs381_.f2 = rhs441_
            d_2651_v159_ = rhs442_
        elif True:
            d_2652_v160_: _dafny.Array
            nw448_ = _dafny.Array(int(0), 13)
            d_2652_v160_ = nw448_
            index430_ = default__.safeIndex(663, (d_2652_v160_).length(0))
            (d_2652_v160_)[index430_] = d_2607_v120_
            index431_ = default__.safeIndex(663, (d_2652_v160_).length(0))
            (d_2652_v160_)[index431_] = d_2607_v120_
            d_2653_v161_: D10
            d_2653_v161_ = D10_DC25(p1)
            (globalState).f2 = (d_2653_v161_).cf48
            d_2654_v162_: _dafny.Seq
            d_2654_v162_ = _dafny.SeqWithoutIsStrInference([(self).f21])
            d_2655_v163_: D21
            d_2655_v163_ = D21_DC56((d_2654_v162_) <= (d_2654_v162_), d_2607_v120_, (778) * ((d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]))
            source26_ = d_2655_v163_
            if source26_.is_DC56:
                d_2656___mcc_h15_ = source26_.cf92
                d_2657___mcc_h16_ = source26_.cf93
                d_2658___mcc_h17_ = source26_.cf94
                d_2659_cf94_ = d_2658___mcc_h17_
                d_2660_cf93_ = d_2657___mcc_h16_
                d_2661_cf92_ = d_2656___mcc_h15_
                d_2662_v164_: D2
                d_2662_v164_ = D2_DC7((0) - (d_2607_v120_), (self).f21, d_2660_cf93_)
                def iife210_(_pat_let72_0):
                    def iife211_(d_2663_dt__update__tmp_h1_):
                        def iife212_(_pat_let73_0):
                            def iife213_(d_2664_dt__update_hcf17_h0_):
                                return D2_DC7((d_2663_dt__update__tmp_h1_).cf15, (d_2663_dt__update__tmp_h1_).cf16, d_2664_dt__update_hcf17_h0_)
                            return iife213_(_pat_let73_0)
                        return iife212_(555)
                    return iife211_(_pat_let72_0)
                (globalState).f10 = default__.fm40(iife210_(d_2662_v164_), globalState)
                d_2665_v165_: _dafny.MultiSet
                d_2665_v165_ = _dafny.MultiSet([(d_2605_v118_).f18])
                d_2666_v166_: _dafny.MultiSet
                d_2666_v166_ = _dafny.MultiSet([(d_2665_v165_).set((d_2605_v118_).f18, default__.abs((d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]))])
                d_2667_v167_: D19
                d_2667_v167_ = D19_DC47(d_2666_v166_)
                d_2668_v168_: C1
                nw449_ = C1()
                nw449_.ctor__((self).f18)
                d_2668_v168_ = nw449_
                d_2669_v169_: _dafny.Map
                d_2669_v169_ = _dafny.Map({d_2667_v167_: d_2668_v168_})
                d_2670_v170_: _dafny.Seq
                d_2670_v170_ = _dafny.SeqWithoutIsStrInference([d_2660_cf93_, d_2607_v120_, d_2660_cf93_, (len(d_2669_v169_)) * (d_2659_cf94_)])
                (globalState).f3 = (d_2670_v170_)[default__.safeIndex(default__.fm1(d_2659_cf94_, p2, (d_2605_v118_).f18, globalState), len(d_2670_v170_))]
                d_2671_v171_: C7
                nw450_ = C7()
                nw450_.ctor__(d_2659_cf94_, (d_2605_v118_).f18)
                d_2671_v171_ = nw450_
                d_2671_v171_ = d_2671_v171_
                d_2672_v172_: C6
                nw451_ = C6()
                nw451_.ctor__(d_2654_v162_, d_2659_cf94_, (d_2605_v118_).f18, (self).f19, (self).f20)
                d_2672_v172_ = nw451_
                d_2673_v173_: _dafny.Seq
                d_2673_v173_ = _dafny.SeqWithoutIsStrInference([d_2672_v172_])
                d_2674_v174_: C6
                nw452_ = C6()
                nw452_.ctor__(d_2654_v162_, len(d_2673_v173_), False, (self).f19, (self).f20)
                d_2674_v174_ = nw452_
                d_2675_v175_: _dafny.Map
                d_2675_v175_ = _dafny.Map({d_2672_v172_: d_2660_cf93_})
                d_2676_v176_: str
                d_2676_v176_ = _dafny.CodePoint('u')
                d_2677_v177_: _dafny.Map
                d_2677_v177_ = _dafny.Map({len(d_2675_v175_): d_2676_v176_})
                d_2678_v178_: _dafny.Array
                nw453_ = _dafny.Array(None, 23)
                nw453_[int(0)] = True
                nw453_[int(1)] = (d_2605_v118_).f18
                nw453_[int(2)] = p2
                nw453_[int(3)] = (d_2605_v118_).f18
                nw453_[int(4)] = (d_2605_v118_).f18
                nw453_[int(5)] = (d_2605_v118_).f18
                nw453_[int(6)] = (d_2605_v118_).f18
                nw453_[int(7)] = False
                nw453_[int(8)] = (d_2659_cf94_) <= (d_2607_v120_)
                nw453_[int(9)] = ((self).f21) == ((self).f18)
                nw453_[int(10)] = d_2661_cf92_
                nw453_[int(11)] = d_2661_cf92_
                nw453_[int(12)] = d_2661_cf92_
                nw453_[int(13)] = (d_2605_v118_).f18
                nw453_[int(14)] = (d_2605_v118_).f18
                nw453_[int(15)] = (d_2605_v118_).f18
                nw453_[int(16)] = (d_2654_v162_)[default__.safeIndex(d_2607_v120_, len(d_2654_v162_))]
                nw453_[int(17)] = d_2661_cf92_
                nw453_[int(18)] = (d_2605_v118_).f18
                nw453_[int(19)] = (d_2605_v118_).f18
                nw453_[int(20)] = (self).f18
                nw453_[int(21)] = not((self).f18)
                nw453_[int(22)] = (d_2677_v177_) != (d_2677_v177_)
                d_2678_v178_ = nw453_
                index432_ = default__.safeIndex(673, (d_2678_v178_).length(0))
                (d_2678_v178_)[index432_] = (True) or (p1)
                index433_ = default__.safeIndex(673, (d_2678_v178_).length(0))
                (d_2678_v178_)[index433_] = (self).f21
            elif True:
                d_2679___mcc_h18_ = source26_.cf91
                d_2680_cf91_ = d_2679___mcc_h18_
                d_2681_v179_: _dafny.Array
                def lambda112_(d_2682_i9_):
                    return (self).f21

                init71_ = lambda112_
                nw454_ = _dafny.Array(None, 28)
                for i0_71_ in range(nw454_.length(0)):
                    nw454_[i0_71_] = init71_(i0_71_)
                d_2681_v179_ = nw454_
                d_2683_v180_: _dafny.Map
                d_2683_v180_ = _dafny.Map({(self).f18: d_2681_v179_})
                d_2684_v181_: _dafny.Map
                d_2684_v181_ = _dafny.Map({(d_2683_v180_).set((self).f18, d_2681_v179_): (d_2605_v118_).f18})
                (globalState).f2 = not(((d_2684_v181_)[_dafny.Map({(d_2605_v118_).f18: d_2681_v179_})] if (_dafny.Map({(d_2605_v118_).f18: d_2681_v179_})) in (d_2684_v181_) else p2))
                (globalState).f2 = (d_2605_v118_).f18
                index434_ = default__.safeIndex(663, (d_2652_v160_).length(0))
                (d_2652_v160_)[index434_] = (d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]
                d_2685_v182_: str
                d_2685_v182_ = _dafny.CodePoint('s')
                d_2685_v182_ = d_2685_v182_
            (globalState).f2 = (self).f21
            if True:
                d_2686_v183_: str
                d_2686_v183_ = _dafny.CodePoint('r')
                d_2687_v184_: _dafny.MultiSet
                d_2687_v184_ = _dafny.MultiSet([722])
                d_2688_v185_: C4
                nw455_ = C4()
                nw455_.ctor__(d_2686_v183_, ((d_2687_v184_)[(d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]] if ((d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]) in (d_2687_v184_) else (d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]), (self).f18, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgqnwe")), (((self).f20) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]: -459}) for d_2689_i10_ in range(default__.abs(44))]))).set(default__.safeIndex((d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))], len(((self).f20) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]: -459}) for d_2690_i10_ in range(default__.abs(44))])))), _dafny.Map({(d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]: len(((self).f19).set(default__.safeIndex(d_2607_v120_, len((self).f19)), _dafny.CodePoint('c')))})))
                d_2688_v185_ = nw455_
                d_2688_v185_ = d_2688_v185_
                d_2691_v186_: _dafny.Array
                nw456_ = _dafny.Array(False, 26)
                d_2691_v186_ = nw456_
                index435_ = default__.safeIndex(451, (d_2691_v186_).length(0))
                (d_2691_v186_)[index435_] = (len(_dafny.SeqWithoutIsStrInference([(d_2605_v118_).f18]))) <= ((d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))])
                index436_ = default__.safeIndex(451, (d_2691_v186_).length(0))
                rhs443_ = not ((self).f21) or (p2)
                rhs444_ = (d_2605_v118_).f18
                lhs382_ = globalState
                lhs383_ = d_2691_v186_
                lhs384_ = default__.safeIndex(451, (d_2691_v186_).length(0))
                lhs382_.f2 = rhs443_
                lhs383_[lhs384_] = rhs444_
                (globalState).f2 = (self).f21
                d_2692_v187_: _dafny.Map
                d_2692_v187_ = _dafny.Map({p1: (d_2605_v118_).f18})
                d_2693_v188_: _dafny.Map
                d_2693_v188_ = _dafny.Map({d_2691_v186_: d_2692_v187_})
                (globalState).f0 = len((d_2693_v188_ if p1 else ((d_2693_v188_).set(d_2691_v186_, d_2692_v187_)) | (d_2693_v188_)))
                d_2694_v189_: D9
                d_2694_v189_ = D9_DC23(p1, p1, (d_2688_v185_).f30)
                d_2695_v190_: _dafny.MultiSet
                d_2695_v190_ = _dafny.MultiSet([d_2694_v189_, d_2694_v189_])
                d_2695_v190_ = d_2695_v190_
            elif True:
                d_2696_v191_: str
                d_2696_v191_ = _dafny.CodePoint('k')
                d_2697_v192_: _dafny.Map
                d_2697_v192_ = _dafny.Map({(d_2605_v118_).f18: p2})
                d_2698_v193_: _dafny.Array
                def lambda113_(d_2699_i11_):
                    return (self).f21

                init72_ = lambda113_
                nw457_ = _dafny.Array(None, 23)
                for i0_72_ in range(nw457_.length(0)):
                    nw457_[i0_72_] = init72_(i0_72_)
                d_2698_v193_ = nw457_
                d_2700_v194_: _dafny.Map
                d_2700_v194_ = _dafny.Map({d_2698_v193_: _dafny.SeqWithoutIsStrInference([d_2696_v191_ for d_2701_i12_ in range(default__.abs(823))])})
                d_2702_v195_: D0
                d_2702_v195_ = D0_DC1(((d_2697_v192_)[p2] if (p2) in (d_2697_v192_) else (self).f18), (d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))], ((d_2700_v194_)[d_2698_v193_] if (d_2698_v193_) in (d_2700_v194_) else (self).f19), (d_2605_v118_).f18, d_2607_v120_)
                d_2703_v196_: _dafny.Map
                d_2703_v196_ = _dafny.Map({(d_2702_v195_).cf1: (self).f21})
                d_2704_v197_: _dafny.Map
                d_2704_v197_ = _dafny.Map({d_2607_v120_: (d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]})
                d_2705_v198_: C4
                nw458_ = C4()
                nw458_.ctor__(d_2696_v191_, d_2607_v120_, ((d_2697_v192_)[(d_2605_v118_).f18] if ((d_2605_v118_).f18) in (d_2697_v192_) else (d_2605_v118_).f18), (self).f19, (_dafny.SeqWithoutIsStrInference([d_2704_v197_])).set(default__.safeIndex(len(d_2703_v196_), len(_dafny.SeqWithoutIsStrInference([d_2704_v197_]))), _dafny.Map({d_2607_v120_: (d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]})))
                d_2705_v198_ = nw458_
                d_2706_v199_: _dafny.Map
                d_2706_v199_ = _dafny.Map({((d_2697_v192_)[p1] if (p1) in (d_2697_v192_) else (d_2605_v118_).f18): len(d_2654_v162_)})
                d_2707_v200_: _dafny.Map
                d_2707_v200_ = _dafny.Map({d_2696_v191_: d_2706_v199_})
                index437_ = default__.safeIndex(663, (d_2652_v160_).length(0))
                (d_2652_v160_)[index437_] = default__.fm1(749, (d_2605_v118_).f18, ((_dafny.Map({d_2696_v191_: d_2706_v199_})).set(_dafny.CodePoint('b'), d_2706_v199_)) != (d_2707_v200_), globalState)
                d_2708_v201_: _dafny.Set
                d_2708_v201_ = _dafny.Set({(d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))]})
                d_2709_v202_: _dafny.MultiSet
                d_2709_v202_ = _dafny.MultiSet([d_2708_v201_])
                d_2710_v203_: _dafny.MultiSet
                d_2710_v203_ = _dafny.MultiSet([p2])
                d_2711_v204_: _dafny.Map
                d_2711_v204_ = _dafny.Map({(d_2605_v118_).f18: d_2696_v191_})
                d_2712_v205_: _dafny.Map
                d_2712_v205_ = _dafny.Map({p2: d_2711_v204_})
                d_2713_v206_: D25
                d_2713_v206_ = D25_DC66(d_2712_v205_)
                d_2714_v207_: _dafny.Map
                d_2714_v207_ = _dafny.Map({(d_2709_v202_).set(_dafny.Set({d_2607_v120_}), default__.abs(((d_2710_v203_)[(d_2605_v118_).f18] if ((d_2605_v118_).f18) in (d_2710_v203_) else d_2607_v120_))): ((d_2713_v206_).cf115 if p2 else d_2712_v205_)})
                d_2714_v207_ = (d_2714_v207_).set((_dafny.MultiSet([d_2708_v201_]) if (self).f18 else d_2709_v202_), d_2712_v205_)
                d_2715_v209_: _dafny.Seq
                def iife214_():
                    coll66_ = _dafny.Map()
                    compr_66_: int
                    for compr_66_ in _dafny.IntegerRange(448, 295):
                        d_2716_v208_: int = compr_66_
                        if ((448) <= (d_2716_v208_)) and ((d_2716_v208_) < (295)):
                            coll66_[(d_2716_v208_) * (168)] = -990
                    return _dafny.Map(coll66_)
                d_2715_v209_ = _dafny.SeqWithoutIsStrInference([d_2607_v120_, (0) - (len(iife214_()
                ))])
                d_2717_v210_: C3
                nw459_ = C3()
                nw459_.ctor__(d_2715_v209_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cedehmcug")), (self).f20, (self).f21)
                d_2717_v210_ = nw459_
                d_2718_v211_: C1
                nw460_ = C1()
                nw460_.ctor__(p2)
                d_2718_v211_ = nw460_
                d_2719_v212_: _dafny.Map
                d_2719_v212_ = _dafny.Map({d_2717_v210_: d_2718_v211_})
                d_2719_v212_ = (d_2719_v212_).set(d_2717_v210_, d_2718_v211_)
                d_2696_v191_ = ((self).f19)[default__.safeIndex((d_2652_v160_)[default__.safeIndex(663, (d_2652_v160_).length(0))], len((self).f19))]
        (globalState).f0 = d_2607_v120_
        (globalState).f2 = (self).f21
        (globalState).f2 = not((d_2607_v120_) != (default__.fm1(d_2607_v120_, (self).f18, p1, globalState)))
        r0 = d_2607_v120_
        return r0

    def m1(self, p0, p1, p2, globalState):
        (globalState).f2 = (self).f21
        d_2720_v0_: _dafny.Array
        nw461_ = _dafny.Array(_dafny.CodePoint('D'), 5)
        d_2720_v0_ = nw461_
        d_2721_v1_: str
        d_2721_v1_ = _dafny.CodePoint('q')
        index438_ = default__.safeIndex(188, (d_2720_v0_).length(0))
        (d_2720_v0_)[index438_] = d_2721_v1_
        index439_ = default__.safeIndex(188, (d_2720_v0_).length(0))
        (d_2720_v0_)[index439_] = d_2721_v1_
        d_2722_v2_: C1
        nw462_ = C1()
        nw462_.ctor__((self).f21)
        d_2722_v2_ = nw462_
        d_2723_v3_: _dafny.Map
        d_2723_v3_ = _dafny.Map({d_2722_v2_: (self).f18})
        d_2723_v3_ = (d_2723_v3_).set(d_2722_v2_, (self).fm3(p1, globalState))
        d_2724_v4_: _dafny.Seq
        d_2724_v4_ = _dafny.SeqWithoutIsStrInference([d_2720_v0_])
        d_2725_v5_: _dafny.Map
        d_2725_v5_ = _dafny.Map({260: (d_2724_v4_)[default__.safeIndex((0) - (-354), len(d_2724_v4_))]})
        d_2725_v5_ = ((d_2725_v5_) | (d_2725_v5_) if (self).f21 else (d_2725_v5_) | (d_2725_v5_))
        d_2726_v6_: _dafny.Set
        d_2726_v6_ = _dafny.Set({(self).f21})
        d_2727_v7_: _dafny.Set
        d_2727_v7_ = _dafny.Set({d_2726_v6_})
        d_2728_v8_: D0
        d_2728_v8_ = D0_DC1(False, p1, p2, (self).f18, 655)
        rhs445_ = ((0) - (-851)) + (p1)
        rhs446_ = (d_2727_v7_) != (d_2727_v7_)
        rhs447_ = ((d_2722_v2_).fm22((self).f19, d_2728_v8_, globalState)) * (p1)
        lhs385_ = globalState
        lhs386_ = globalState
        lhs387_ = globalState
        lhs385_.f3 = rhs445_
        lhs386_.f2 = rhs446_
        lhs387_.f3 = rhs447_
        (globalState).f15 = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f19 for d_2729_i0_ in range(default__.abs(495))])), p1)

    def m3(self, p0, p1, p2, globalState):
        d_2730_v0_: str
        d_2730_v0_ = _dafny.CodePoint('h')
        d_2731_v1_: _dafny.Map
        d_2731_v1_ = _dafny.Map({p0: d_2730_v0_})
        d_2732_v2_: _dafny.Seq
        d_2732_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2733_v3_: T0
        nw463_ = C4()
        nw463_.ctor__(((d_2731_v1_)[len(d_2732_v2_)] if (len(d_2732_v2_)) in (d_2731_v1_) else d_2730_v0_), p0, (self).f21, ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjdlrtgp"))) + ((self).f19)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjdlrtgp"))) + ((self).f19))), d_2730_v0_), (self).f20)
        d_2733_v3_ = nw463_
        d_2734_v4_: _dafny.Map
        d_2734_v4_ = _dafny.Map({181: d_2733_v3_})
        d_2733_v3_ = ((d_2734_v4_)[p0] if (p0) in (d_2734_v4_) else d_2733_v3_)
        d_2735_v5_: _dafny.Array
        nw464_ = _dafny.Array(int(0), 21)
        d_2735_v5_ = nw464_
        d_2736_v6_: C5
        nw465_ = C5()
        nw465_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ued")), (self).f19, (self).f18)
        d_2736_v6_ = nw465_
        d_2737_v7_: _dafny.Array
        nw466_ = _dafny.Array(None, 6)
        nw466_[int(0)] = d_2736_v6_
        nw466_[int(1)] = d_2736_v6_
        nw466_[int(2)] = d_2736_v6_
        nw466_[int(3)] = d_2736_v6_
        nw466_[int(4)] = d_2736_v6_
        nw466_[int(5)] = d_2736_v6_
        d_2737_v7_ = nw466_
        d_2738_v8_: _dafny.Seq
        d_2738_v8_ = _dafny.SeqWithoutIsStrInference([d_2737_v7_])
        d_2739_v9_: _dafny.Seq
        d_2739_v9_ = _dafny.SeqWithoutIsStrInference([d_2738_v8_])
        d_2740_v10_: _dafny.Map
        d_2740_v10_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_2737_v7_])) != ((d_2739_v9_)[default__.safeIndex(default__.fm1(p0, False, True, globalState), len(d_2739_v9_))]): d_2735_v5_})
        d_2735_v5_ = ((d_2740_v10_)[False] if (False) in (d_2740_v10_) else d_2735_v5_)
        (globalState).f2 = (d_2736_v6_).fm15(p0, p0, globalState)
        hi14_ = len((d_2736_v6_.f27) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwis"))))
        for d_2741_i0_ in range((p0) + ((0) - (default__.fm1(p0, True, (d_2733_v3_).f18, globalState))), hi14_):
            d_2742_v11_: _dafny.Map
            d_2742_v11_ = _dafny.Map({default__.safeModuloInt(p0, (0) - (d_2741_i0_)): (((d_2736_v6_).f28).set(default__.safeIndex(p0, len((d_2736_v6_).f28)), _dafny.CodePoint('w'))) != (d_2736_v6_.f27)})
            (globalState).f2 = ((d_2742_v11_)[p0] if (p0) in (d_2742_v11_) else p2)
            d_2743_v12_: _dafny.Array
            def lambda114_(d_2744_v0_):
                def lambda115_(d_2745_i1_):
                    return d_2744_v0_

                return lambda115_

            init73_ = lambda114_(d_2730_v0_)
            nw467_ = _dafny.Array(None, 27)
            for i0_73_ in range(nw467_.length(0)):
                nw467_[i0_73_] = init73_(i0_73_)
            d_2743_v12_ = nw467_
            index440_ = default__.safeIndex(958, (d_2743_v12_).length(0))
            (d_2743_v12_)[index440_] = _dafny.CodePoint('w')
            index441_ = default__.safeIndex(958, (d_2743_v12_).length(0))
            (d_2743_v12_)[index441_] = d_2730_v0_
            d_2735_v5_ = d_2735_v5_
            d_2746_v13_: _dafny.Seq
            d_2746_v13_ = _dafny.SeqWithoutIsStrInference([False, True, p2])
            d_2747_v14_: _dafny.Array
            nw468_ = _dafny.Array(None, 29)
            nw468_[int(0)] = not((self).f21)
            nw468_[int(1)] = not((d_2746_v13_)[default__.safeIndex(p0, len(d_2746_v13_))])
            nw468_[int(2)] = (self).f18
            nw468_[int(3)] = (self).f21
            nw468_[int(4)] = (self).f18
            nw468_[int(5)] = (self).f18
            nw468_[int(6)] = (d_2733_v3_).f18
            nw468_[int(7)] = (self).f18
            nw468_[int(8)] = (self).f18
            nw468_[int(9)] = False
            nw468_[int(10)] = (d_2733_v3_).f18
            nw468_[int(11)] = (d_2733_v3_).f18
            nw468_[int(12)] = p2
            nw468_[int(13)] = (d_2733_v3_).f18
            nw468_[int(14)] = True
            nw468_[int(15)] = False
            nw468_[int(16)] = p2
            nw468_[int(17)] = not((self).f21)
            nw468_[int(18)] = (self).f21
            nw468_[int(19)] = (d_2733_v3_).f18
            nw468_[int(20)] = not((d_2733_v3_).f18)
            nw468_[int(21)] = p2
            nw468_[int(22)] = (self).f21
            nw468_[int(23)] = (d_2733_v3_).f18
            nw468_[int(24)] = (d_2733_v3_).f18
            nw468_[int(25)] = (d_2733_v3_).f18
            nw468_[int(26)] = p2
            nw468_[int(27)] = p2
            nw468_[int(28)] = (d_2733_v3_).f18
            d_2747_v14_ = nw468_
            d_2748_v15_: C9
            nw469_ = C9()
            nw469_.ctor__(d_2747_v14_, (d_2733_v3_).f18)
            d_2748_v15_ = nw469_
        index442_ = default__.safeIndex(674, (d_2735_v5_).length(0))
        (d_2735_v5_)[index442_] = (0) - (default__.safeDivisionInt(p0, p0))
        index443_ = default__.safeIndex(674, (d_2735_v5_).length(0))
        (d_2735_v5_)[index443_] = p0
        (globalState).f2 = (self).f18

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        d_2749_v0_: int
        d_2749_v0_ = 534
        d_2750_v1_: C7
        nw470_ = C7()
        nw470_.ctor__(d_2749_v0_, (self).f18)
        d_2750_v1_ = nw470_
        d_2751_i0_: int
        d_2751_i0_ = 0
        with _dafny.label("17"):
            while (self).f18:
                with _dafny.c_label("17"):
                    if (d_2751_i0_) >= (100):
                        raise _dafny.Break("17")
                    d_2751_i0_ = (d_2751_i0_) + (1)
                    (globalState).f9 = default__.fm12(d_2749_v0_, (self).f21, globalState)
                    d_2752_v2_: _dafny.Set
                    d_2752_v2_ = _dafny.Set({len((self).f19), d_2749_v0_})
                    (d_2750_v1_).f24 = default__.fm1(d_2750_v1_.f24, default__.fm23(d_2752_v2_, len(_dafny.Map({p0: p0})), globalState), (self).f21, globalState)
                    d_2753_v3_: C1
                    nw471_ = C1()
                    nw471_.ctor__(p0)
                    d_2753_v3_ = nw471_
                    d_2754_v4_: C7
                    nw472_ = C7()
                    nw472_.ctor__((451) - (d_2750_v1_.f24), (d_2749_v0_) == (79))
                    d_2754_v4_ = nw472_
                    pass
            pass
        d_2755_v5_: C3
        nw473_ = C3()
        nw473_.ctor__(_dafny.SeqWithoutIsStrInference([d_2750_v1_.f24 for d_2756_i1_ in range(default__.abs(866))]), (self).f19, (self).f20, p0)
        d_2755_v5_ = nw473_
        d_2757_v6_: _dafny.Seq
        d_2757_v6_ = _dafny.SeqWithoutIsStrInference([d_2755_v5_, d_2755_v5_])
        d_2757_v6_ = (d_2757_v6_) + (_dafny.SeqWithoutIsStrInference([d_2755_v5_, d_2755_v5_]))
        d_2758_v7_: _dafny.MultiSet
        d_2758_v7_ = _dafny.MultiSet([(d_2749_v0_) - (-987)])
        d_2758_v7_ = d_2758_v7_
        d_2759_v8_: _dafny.Set
        d_2759_v8_ = _dafny.Set({(self).f19, (self).f19, (self).f19, (self).f19, (self).f19})
        d_2760_v9_: _dafny.Array
        def lambda116_(d_2761_v1_):
            def lambda117_(d_2762_i2_):
                return (_dafny.Map({_dafny.CodePoint('x'): d_2761_v1_.f24})) | (_dafny.Map({_dafny.CodePoint('h'): d_2761_v1_.f24}))

            return lambda117_

        init74_ = lambda116_(d_2750_v1_)
        nw474_ = _dafny.Array(None, 26)
        for i0_74_ in range(nw474_.length(0)):
            nw474_[i0_74_] = init74_(i0_74_)
        d_2760_v9_ = nw474_
        d_2763_v10_: str
        d_2763_v10_ = _dafny.CodePoint('i')
        d_2764_v11_: _dafny.Map
        d_2764_v11_ = _dafny.Map({d_2763_v10_: 165})
        index444_ = default__.safeIndex(513, (d_2760_v9_).length(0))
        (d_2760_v9_)[index444_] = (d_2764_v11_) | (d_2764_v11_)
        d_2765_v12_: _dafny.Set
        d_2765_v12_ = _dafny.Set({d_2749_v0_})
        d_2766_v13_: _dafny.Map
        d_2766_v13_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2763_v10_ for d_2767_i3_ in range(default__.abs(-385))])): d_2749_v0_})
        d_2768_v15_: _dafny.Map
        d_2768_v15_ = _dafny.Map({d_2763_v10_: True})
        index445_ = default__.safeIndex(513, (d_2760_v9_).length(0))
        def iife215_():
            coll67_ = _dafny.Map()
            compr_67_: str
            for compr_67_ in (d_2768_v15_).keys.Elements:
                d_2769_v14_: str = compr_67_
                if (d_2769_v14_) in (d_2768_v15_):
                    coll67_[d_2769_v14_] = (0) - (d_2749_v0_)
            return _dafny.Map(coll67_)
        rhs448_ = default__.fm23(d_2765_v12_, len(d_2766_v13_), globalState)
        rhs449_ = d_2749_v0_
        rhs450_ = (_dafny.Set({(self).f19, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pewnmqk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cswc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erclvwh")), (self).f19})) | (d_2759_v8_)
        rhs451_ = d_2749_v0_
        rhs452_ = ((default__.fm64(d_2763_v10_, d_2750_v1_.f24, p0, p0, globalState)) | (d_2764_v11_)) | (iife215_()
        )
        lhs388_ = globalState
        lhs389_ = globalState
        lhs390_ = d_2760_v9_
        lhs391_ = default__.safeIndex(513, (d_2760_v9_).length(0))
        r2 = rhs448_
        lhs388_.f15 = rhs449_
        d_2759_v8_ = rhs450_
        lhs389_.f0 = rhs451_
        lhs390_[lhs391_] = rhs452_
        d_2770_v16_: D17
        d_2770_v16_ = D17_DC42((self).f18)
        source27_ = d_2770_v16_
        if source27_.is_DC42:
            d_2771___mcc_h0_ = source27_.cf74
            d_2772_cf74_ = d_2771___mcc_h0_
            d_2773_v17_: _dafny.Array
            def lambda118_(d_2774_p0_):
                def lambda119_(d_2775_i4_):
                    return (True) or (not(d_2774_p0_))

                return lambda119_

            init75_ = lambda118_(p0)
            nw475_ = _dafny.Array(None, 16)
            for i0_75_ in range(nw475_.length(0)):
                nw475_[i0_75_] = init75_(i0_75_)
            d_2773_v17_ = nw475_
            index446_ = default__.safeIndex(145, (d_2773_v17_).length(0))
            (d_2773_v17_)[index446_] = (self).f18
            index447_ = default__.safeIndex(145, (d_2773_v17_).length(0))
            (d_2773_v17_)[index447_] = (d_2772_cf74_) and ((self).f21)
            d_2776_v18_: _dafny.Map
            d_2776_v18_ = _dafny.Map({d_2750_v1_.f24: (self).f18})
            d_2777_v19_: D17
            d_2777_v19_ = D17_DC41(d_2776_v18_)
            d_2778_v20_: _dafny.Set
            d_2778_v20_ = _dafny.Set({d_2777_v19_, d_2777_v19_})
            d_2779_v21_: _dafny.Array
            def lambda120_(d_2780_v1_):
                def lambda121_(d_2781_i5_):
                    return (d_2781_i5_) + (d_2780_v1_.f24)

                return lambda121_

            init76_ = lambda120_(d_2750_v1_)
            nw476_ = _dafny.Array(None, 10)
            for i0_76_ in range(nw476_.length(0)):
                nw476_[i0_76_] = init76_(i0_76_)
            d_2779_v21_ = nw476_
            d_2782_v22_: D24
            d_2782_v22_ = D24_DC65(p0, d_2778_v20_, (self).f21, d_2749_v0_, d_2779_v21_)
            if (d_2782_v22_).cf112:
                d_2783_v23_: _dafny.Set
                d_2783_v23_ = _dafny.Set({p0, d_2772_cf74_})
                d_2784_v24_: _dafny.Seq
                d_2784_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not((d_2773_v17_)[default__.safeIndex(145, (d_2773_v17_).length(0))])})) for d_2785_i6_ in range(default__.abs(41))]), ((d_2755_v5_).f35).set(default__.safeIndex(d_2749_v0_, len((d_2755_v5_).f35)), len(d_2783_v23_))])
                d_2786_v25_: _dafny.Map
                d_2786_v25_ = _dafny.Map({(d_2784_v24_)[default__.safeIndex(len((self).f19), len(d_2784_v24_))]: d_2772_cf74_})
                d_2786_v25_ = ((d_2786_v25_) | ((d_2786_v25_).set((d_2755_v5_).f35, p0))) | ((d_2786_v25_ if (self).f18 else d_2786_v25_))
                d_2787_v26_: _dafny.Array
                nw477_ = _dafny.Array(None, 6)
                d_2787_v26_ = nw477_
                d_2788_v27_: C2
                nw478_ = C2()
                nw478_.ctor__(d_2750_v1_.f24, _dafny.SeqWithoutIsStrInference([d_2750_v1_.f24, 904]), False)
                d_2788_v27_ = nw478_
                index448_ = default__.safeIndex(468, (d_2787_v26_).length(0))
                (d_2787_v26_)[index448_] = d_2788_v27_
                index449_ = default__.safeIndex(468, (d_2787_v26_).length(0))
                (d_2787_v26_)[index449_] = d_2788_v27_
                (globalState).f2 = (default__.safeModuloInt(d_2788_v27_.f33, (0) - (d_2750_v1_.f24))) != ((0) - ((d_2755_v5_).fm38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqcnaoqeg")), globalState)))
                index450_ = default__.safeIndex(460, (d_2779_v21_).length(0))
                (d_2779_v21_)[index450_] = d_2750_v1_.f24
                d_2789_v28_: _dafny.Map
                d_2789_v28_ = _dafny.Map({d_2773_v17_: (d_2773_v17_)[default__.safeIndex(145, (d_2773_v17_).length(0))]})
                d_2790_v29_: D13
                d_2790_v29_ = D13_DC32(d_2789_v28_)
                pat_let_tv39_ = d_2789_v28_
                pat_let_tv40_ = d_2789_v28_
                pat_let_tv41_ = d_2773_v17_
                d_2791_v30_: _dafny.Array
                nw479_ = _dafny.Array(None, 12)
                nw479_[int(0)] = d_2790_v29_
                def iife216_(_pat_let74_0):
                    def iife217_(d_2792_dt__update__tmp_h0_):
                        def iife218_(_pat_let75_0):
                            def iife219_(d_2793_dt__update_hcf54_h0_):
                                return D13_DC32(d_2793_dt__update_hcf54_h0_)
                            return iife219_(_pat_let75_0)
                        return iife218_(pat_let_tv39_)
                    return iife217_(_pat_let74_0)
                nw479_[int(1)] = iife216_(d_2790_v29_)
                nw479_[int(2)] = d_2790_v29_
                nw479_[int(3)] = d_2790_v29_
                nw479_[int(4)] = d_2790_v29_
                def iife220_(_pat_let76_0):
                    def iife221_(d_2794_dt__update__tmp_h1_):
                        def iife222_(_pat_let77_0):
                            def iife223_(d_2795_dt__update_hcf54_h1_):
                                return D13_DC32(d_2795_dt__update_hcf54_h1_)
                            return iife223_(_pat_let77_0)
                        return iife222_((pat_let_tv40_).set(pat_let_tv41_, (self).f21))
                    return iife221_(_pat_let76_0)
                nw479_[int(5)] = iife220_(d_2790_v29_)
                nw479_[int(6)] = D13_DC32(d_2789_v28_)
                nw479_[int(7)] = d_2790_v29_
                nw479_[int(8)] = d_2790_v29_
                nw479_[int(9)] = (d_2790_v29_ if p0 else d_2790_v29_)
                nw479_[int(10)] = D13_DC32(d_2789_v28_)
                nw479_[int(11)] = d_2790_v29_
                d_2791_v30_ = nw479_
                index451_ = default__.safeIndex(271, (d_2791_v30_).length(0))
                (d_2791_v30_)[index451_] = d_2790_v29_
                pat_let_tv42_ = d_2789_v28_
                index452_ = default__.safeIndex(460, (d_2779_v21_).length(0))
                index453_ = default__.safeIndex(271, (d_2791_v30_).length(0))
                index454_ = default__.safeIndex(145, (d_2773_v17_).length(0))
                def iife224_(_pat_let78_0):
                    def iife225_(d_2796_dt__update__tmp_h2_):
                        def iife226_(_pat_let79_0):
                            def iife227_(d_2797_dt__update_hcf54_h2_):
                                return D13_DC32(d_2797_dt__update_hcf54_h2_)
                            return iife227_(_pat_let79_0)
                        return iife226_(pat_let_tv42_)
                    return iife225_(_pat_let78_0)
                rhs453_ = (658) + (236)
                rhs454_ = (d_2790_v29_ if not (d_2772_cf74_) or ((self).f21) else iife224_(d_2790_v29_))
                rhs455_ = (self).f18
                rhs456_ = (d_2763_v10_) not in ((self).f19)
                rhs457_ = (d_2755_v5_).f35
                lhs392_ = d_2779_v21_
                lhs393_ = default__.safeIndex(460, (d_2779_v21_).length(0))
                lhs394_ = d_2791_v30_
                lhs395_ = default__.safeIndex(271, (d_2791_v30_).length(0))
                lhs396_ = d_2773_v17_
                lhs397_ = default__.safeIndex(145, (d_2773_v17_).length(0))
                lhs398_ = globalState
                lhs392_[lhs393_] = rhs453_
                lhs394_[lhs395_] = rhs454_
                lhs396_[lhs397_] = rhs455_
                d_2772_cf74_ = rhs456_
                lhs398_.f10 = rhs457_
                index455_ = default__.safeIndex(460, (d_2779_v21_).length(0))
                (d_2779_v21_)[index455_] = d_2750_v1_.f24
            elif True:
                r2 = (self).f21
                d_2798_v31_: _dafny.Seq
                d_2798_v31_ = _dafny.SeqWithoutIsStrInference([(self).f18, False])
                d_2799_v32_: C5
                nw480_ = C5()
                nw480_.ctor__((self).f19, (self).f19, True)
                d_2799_v32_ = nw480_
                d_2800_v33_: D24
                d_2800_v33_ = D24_DC64((d_2798_v31_)[default__.safeIndex(314, len(d_2798_v31_))], d_2772_cf74_, d_2799_v32_, len((self).f19), default__.safeModuloInt(d_2750_v1_.f24, (0) - (d_2749_v0_)))
                def iife228_(_pat_let80_0):
                    def iife229_(d_2801_dt__update__tmp_h3_):
                        def iife230_(_pat_let81_0):
                            def iife231_(d_2802_dt__update_hcf106_h0_):
                                return D24_DC64((d_2801_dt__update__tmp_h3_).cf105, d_2802_dt__update_hcf106_h0_, (d_2801_dt__update__tmp_h3_).cf107, (d_2801_dt__update__tmp_h3_).cf108, (d_2801_dt__update__tmp_h3_).cf109)
                            return iife231_(_pat_let81_0)
                        return iife230_((self).f18)
                    return iife229_(_pat_let80_0)
                d_2800_v33_ = iife228_(D24_DC64(p0, d_2772_cf74_, d_2799_v32_, d_2750_v1_.f24, d_2750_v1_.f24))
                d_2803_v34_: C1
                nw481_ = C1()
                nw481_.ctor__(d_2772_cf74_)
                d_2803_v34_ = nw481_
                d_2804_v35_: _dafny.Array
                def lambda122_(d_2805_v10_):
                    def lambda123_(d_2806_i7_):
                        return d_2805_v10_

                    return lambda123_

                init77_ = lambda122_(d_2763_v10_)
                nw482_ = _dafny.Array(None, 3)
                for i0_77_ in range(nw482_.length(0)):
                    nw482_[i0_77_] = init77_(i0_77_)
                d_2804_v35_ = nw482_
                d_2807_v36_: _dafny.Map
                d_2807_v36_ = _dafny.Map({p0: _dafny.Map({d_2804_v35_: d_2750_v1_.f24})})
                d_2808_v37_: _dafny.Map
                d_2808_v37_ = _dafny.Map({d_2804_v35_: d_2750_v1_.f24})
                d_2807_v36_ = (d_2807_v36_).set(p0, (d_2808_v37_) | (_dafny.Map({d_2804_v35_: len(d_2776_v18_)})))
                d_2809_v38_: _dafny.Set
                d_2810_v39_: int
                d_2811_v40_: int
                d_2812_v41_: _dafny.Seq
                out66_: _dafny.Set
                out67_: int
                out68_: int
                out69_: _dafny.Seq
                out66_, out67_, out68_, out69_ = (d_2755_v5_).m10(p0, globalState)
                d_2809_v38_ = out66_
                d_2810_v39_ = out67_
                d_2811_v40_ = out68_
                d_2812_v41_ = out69_
            (globalState).f0 = 950
            d_2813_v42_: _dafny.Seq
            d_2813_v42_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f19])
            d_2814_v43_: D26
            d_2814_v43_ = D26_DC69(d_2813_v42_)
            d_2813_v42_ = (d_2814_v43_).cf121
        elif source27_.is_DC43:
            d_2815___mcc_h1_ = source27_.cf75
            d_2816___mcc_h2_ = source27_.cf76
            d_2817_cf76_ = d_2816___mcc_h2_
            d_2818_cf75_ = d_2815___mcc_h1_
            (d_2750_v1_).f24 = (0) - ((d_2749_v0_ if d_2817_cf76_ else (d_2755_v5_).fm38((self).f19, globalState)))
            d_2819_v44_: _dafny.Seq
            d_2819_v44_ = _dafny.SeqWithoutIsStrInference([(self).f21, False])
            d_2819_v44_ = d_2819_v44_
            d_2820_v45_: _dafny.Set
            d_2820_v45_ = _dafny.Set({(self).f18, d_2817_cf76_, d_2817_cf76_, p0, not((self).f18)})
            d_2821_v46_: _dafny.Seq
            d_2821_v46_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({p0})).intersection(d_2820_v45_), d_2820_v45_])
            d_2822_v47_: _dafny.Map
            d_2822_v47_ = _dafny.Map({d_2749_v0_: d_2750_v1_.f24})
            d_2821_v46_ = default__.fm62((_dafny.MultiSet([d_2750_v1_.f24])).cardinality, d_2822_v47_, globalState)
            d_2819_v44_ = ((d_2819_v44_) + (d_2819_v44_)) + ((d_2819_v44_ if (self).f21 else (d_2819_v44_).set(default__.safeIndex(d_2749_v0_, len(d_2819_v44_)), (self).f18)))
        elif source27_.is_DC44:
            d_2823___mcc_h3_ = source27_.cf77
            d_2824___mcc_h4_ = source27_.cf78
            d_2825_cf78_ = d_2824___mcc_h4_
            d_2826_cf77_ = d_2823___mcc_h3_
            d_2763_v10_ = d_2763_v10_
            r1 = (self).f18
            d_2827_v48_: D19
            d_2827_v48_ = D19_DC50(not((self).f18))
            source28_ = D19_DC51(d_2827_v48_)
            if source28_.is_DC48:
                d_2828___mcc_h6_ = source28_.cf82
                d_2829_cf82_ = d_2828___mcc_h6_
                d_2830_v49_: _dafny.Map
                d_2830_v49_ = _dafny.Map({(self).f18: True})
                d_2830_v49_ = (d_2830_v49_).set(d_2829_cf82_, (self).f21)
                d_2831_v50_: C2
                nw483_ = C2()
                nw483_.ctor__(((d_2758_v7_).cardinality if p0 else d_2750_v1_.f24), (d_2755_v5_).f35, (self).f21)
                d_2831_v50_ = nw483_
                (globalState).f15 = (d_2750_v1_.f24) - (d_2750_v1_.f24)
                d_2832_v51_: _dafny.Seq
                d_2832_v51_ = _dafny.SeqWithoutIsStrInference([(self).f19, ((self).f19).set(default__.safeIndex(d_2750_v1_.f24, len((self).f19)), d_2763_v10_), (self).f19])
                (globalState).f3 = len((((self).f19) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eofv"))) if (self).f21 else (d_2832_v51_)[default__.safeIndex(d_2750_v1_.f24, len(d_2832_v51_))]))
            elif source28_.is_DC49:
                d_2833___mcc_h7_ = source28_.cf83
                d_2834_cf83_ = d_2833___mcc_h7_
                (globalState).f15 = len((d_2755_v5_).f35)
                (d_2750_v1_).f24 = (len(default__.fm48(globalState))) * (d_2750_v1_.f24)
                index456_ = default__.safeIndex(760, (d_2826_cf77_).length(0))
                (d_2826_cf77_)[index456_] = p0
                d_2835_v52_: D10
                d_2835_v52_ = D10_DC25(d_2834_cf83_)
                index457_ = default__.safeIndex(760, (d_2826_cf77_).length(0))
                (d_2826_cf77_)[index457_] = (d_2835_v52_).cf48
                (globalState).f15 = d_2825_cf78_
            elif source28_.is_DC50:
                d_2836___mcc_h8_ = source28_.cf84
                d_2837_cf84_ = d_2836___mcc_h8_
                d_2838_v53_: _dafny.Map
                d_2838_v53_ = _dafny.Map({len((self).f19): d_2837_cf84_})
                d_2839_v54_: _dafny.MultiSet
                d_2839_v54_ = _dafny.MultiSet([(self).f19, (self).f19, (self).f19, ((self).f19).set(default__.safeIndex(len(d_2838_v53_), len((self).f19)), d_2763_v10_), _dafny.SeqWithoutIsStrInference([d_2763_v10_ for d_2840_i8_ in range(default__.abs(186))])])
                (d_2750_v1_).m5(d_2839_v54_, d_2750_v1_.f24, d_2758_v7_, default__.safeDivisionInt(d_2750_v1_.f24, 782), globalState)
                d_2841_v55_: _dafny.Array
                nw484_ = _dafny.Array(int(0), 23)
                d_2841_v55_ = nw484_
                d_2842_v56_: _dafny.Seq
                d_2842_v56_ = _dafny.SeqWithoutIsStrInference([d_2838_v53_])
                index458_ = default__.safeIndex(14, (d_2841_v55_).length(0))
                (d_2841_v55_)[index458_] = len(d_2842_v56_)
                index459_ = default__.safeIndex(14, (d_2841_v55_).length(0))
                (d_2841_v55_)[index459_] = default__.safeModuloInt(d_2750_v1_.f24, d_2749_v0_)
                d_2843_v57_: T1
                nw485_ = C3()
                nw485_.ctor__(_dafny.SeqWithoutIsStrInference([d_2750_v1_.f24, 754]), (d_2755_v5_).fm37(d_2750_v1_.f24, globalState), (self).f20, (d_2755_v5_).fm2(not(d_2837_cf84_), d_2837_cf84_, d_2837_cf84_, globalState))
                d_2843_v57_ = nw485_
                d_2844_v58_: _dafny.MultiSet
                d_2844_v58_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2843_v57_])])
                d_2845_v59_: _dafny.Seq
                d_2845_v59_ = _dafny.SeqWithoutIsStrInference([d_2843_v57_])
                d_2846_v60_: _dafny.Seq
                d_2846_v60_ = _dafny.SeqWithoutIsStrInference([d_2845_v59_])
                d_2847_v61_: _dafny.Seq
                d_2847_v61_ = _dafny.SeqWithoutIsStrInference([d_2844_v58_, _dafny.MultiSet(d_2846_v60_), d_2844_v58_, d_2844_v58_])
                d_2848_v62_: _dafny.Map
                d_2848_v62_ = _dafny.Map({p0: d_2845_v59_})
                d_2849_v63_: _dafny.Seq
                d_2849_v63_ = d_2846_v60_
                d_2850_v64_: _dafny.Array
                nw486_ = _dafny.Array(None, 20)
                nw486_[int(0)] = d_2844_v58_
                nw486_[int(1)] = (d_2844_v58_) - (d_2844_v58_)
                nw486_[int(2)] = d_2844_v58_
                nw486_[int(3)] = d_2844_v58_
                nw486_[int(4)] = (d_2844_v58_) - (d_2844_v58_)
                nw486_[int(5)] = d_2844_v58_
                nw486_[int(6)] = d_2844_v58_
                nw486_[int(7)] = d_2844_v58_
                nw486_[int(8)] = d_2844_v58_
                nw486_[int(9)] = (d_2844_v58_) | ((d_2847_v61_)[default__.safeIndex(d_2750_v1_.f24, len(d_2847_v61_))])
                nw486_[int(10)] = (d_2844_v58_) - (_dafny.MultiSet([((d_2848_v62_)[d_2837_cf84_] if (d_2837_cf84_) in (d_2848_v62_) else d_2845_v59_)]))
                nw486_[int(11)] = d_2844_v58_
                nw486_[int(12)] = _dafny.MultiSet([d_2845_v59_])
                nw486_[int(13)] = d_2844_v58_
                nw486_[int(14)] = d_2844_v58_
                nw486_[int(15)] = (d_2844_v58_).set(d_2845_v59_, default__.abs(len((d_2843_v57_).f19)))
                nw486_[int(16)] = d_2844_v58_
                nw486_[int(17)] = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2843_v57_, d_2843_v57_]), (d_2845_v59_).set(default__.safeIndex(d_2749_v0_, len(d_2845_v59_)), d_2843_v57_)])
                nw486_[int(18)] = (d_2844_v58_) - (_dafny.MultiSet((d_2849_v63_)))
                nw486_[int(19)] = (d_2844_v58_).intersection(d_2844_v58_)
                d_2850_v64_ = nw486_
                d_2851_v65_: _dafny.Map
                d_2851_v65_ = _dafny.Map({d_2837_cf84_: d_2844_v58_})
                index460_ = default__.safeIndex(88, (d_2850_v64_).length(0))
                (d_2850_v64_)[index460_] = ((d_2851_v65_)[(self).f21] if ((self).f21) in (d_2851_v65_) else d_2844_v58_)
                d_2852_v66_: D28
                d_2852_v66_ = D28_DC74(d_2844_v58_)
                index461_ = default__.safeIndex(88, (d_2850_v64_).length(0))
                rhs458_ = default__.fm12(len((d_2755_v5_).f35), (d_2843_v57_).f18, globalState)
                rhs459_ = (d_2852_v66_).cf131
                lhs399_ = globalState
                lhs400_ = d_2850_v64_
                lhs401_ = default__.safeIndex(88, (d_2850_v64_).length(0))
                lhs399_.f9 = rhs458_
                lhs400_[lhs401_] = rhs459_
                d_2853_v67_: _dafny.Array
                nw487_ = _dafny.Array(_dafny.Map({}), 10)
                d_2853_v67_ = nw487_
                d_2854_v68_: D20
                d_2854_v68_ = D20_DC53((self).f18, (d_2843_v57_).f18, d_2763_v10_)
                index462_ = default__.safeIndex(291, (d_2853_v67_).length(0))
                (d_2853_v67_)[index462_] = (_dafny.Map({d_2837_cf84_: p0})).set((d_2854_v68_).cf87, d_2837_cf84_)
                d_2855_v69_: _dafny.Seq
                d_2855_v69_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_2856_v70_: _dafny.Seq
                d_2856_v70_ = _dafny.SeqWithoutIsStrInference([(d_2855_v69_)[default__.safeIndex(len((self).f19), len(d_2855_v69_))]])
                d_2857_v71_: _dafny.Map
                d_2857_v71_ = _dafny.Map({(d_2855_v69_)[default__.safeIndex((0) - (d_2750_v1_.f24), len(d_2855_v69_))]: d_2837_cf84_})
                d_2858_v72_: _dafny.Seq
                d_2858_v72_ = _dafny.SeqWithoutIsStrInference([d_2857_v71_])
                d_2859_v73_: _dafny.Seq
                d_2859_v73_ = _dafny.SeqWithoutIsStrInference([(d_2857_v71_) | (d_2857_v71_), _dafny.Map({(self).f18: (self).f21}), d_2857_v71_, (d_2858_v72_)[default__.safeIndex(d_2750_v1_.f24, len(d_2858_v72_))], d_2857_v71_])
                index463_ = default__.safeIndex(291, (d_2853_v67_).length(0))
                (d_2853_v67_)[index463_] = (d_2858_v72_)[default__.safeIndex(((_dafny.MultiSet([(d_2843_v57_).f18])).cardinality) - (len((self).f19)), len(d_2858_v72_))]
            elif source28_.is_DC47:
                d_2860___mcc_h9_ = source28_.cf81
                d_2861_cf81_ = d_2860___mcc_h9_
                d_2862_v74_: _dafny.Array
                nw488_ = _dafny.Array(int(0), 24)
                d_2862_v74_ = nw488_
                d_2863_v75_: _dafny.Seq
                d_2863_v75_ = _dafny.SeqWithoutIsStrInference([d_2862_v74_, d_2862_v74_])
                d_2864_v76_: _dafny.Map
                d_2864_v76_ = _dafny.Map({p0: d_2862_v74_})
                d_2865_v77_: _dafny.Seq
                d_2865_v77_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_2866_v78_: _dafny.Map
                d_2866_v78_ = _dafny.Map({d_2862_v74_: len(d_2865_v77_)})
                d_2867_v79_: _dafny.Set
                d_2867_v79_ = _dafny.Set({True})
                d_2868_v80_: D16
                d_2868_v80_ = D16_DC40((default__.fm25(632, globalState)).cardinality, (d_2755_v5_).fm38((self).f19, globalState), d_2750_v1_.f24, len(d_2867_v79_))
                d_2869_v81_: _dafny.Map
                d_2869_v81_ = _dafny.Map({d_2750_v1_.f24: d_2868_v80_})
                d_2870_v83_: _dafny.Array
                nw489_ = _dafny.Array(None, 17)
                nw489_[int(0)] = d_2869_v81_
                nw489_[int(1)] = d_2869_v81_
                nw489_[int(2)] = d_2869_v81_
                nw489_[int(3)] = d_2869_v81_
                nw489_[int(4)] = d_2869_v81_
                nw489_[int(5)] = d_2869_v81_
                nw489_[int(6)] = d_2869_v81_
                nw489_[int(7)] = d_2869_v81_
                def iife232_():
                    coll68_ = _dafny.Map()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(579, 785):
                        d_2871_v82_: int = compr_68_
                        if ((579) <= (d_2871_v82_)) and ((d_2871_v82_) < (785)):
                            coll68_[(d_2871_v82_) * (d_2750_v1_.f24)] = d_2868_v80_
                    return _dafny.Map(coll68_)
                nw489_[int(8)] = iife232_()
                
                nw489_[int(9)] = d_2869_v81_
                nw489_[int(10)] = (d_2869_v81_).set(d_2750_v1_.f24, d_2868_v80_)
                nw489_[int(11)] = _dafny.Map({d_2749_v0_: d_2868_v80_})
                nw489_[int(12)] = d_2869_v81_
                nw489_[int(13)] = d_2869_v81_
                nw489_[int(14)] = d_2869_v81_
                nw489_[int(15)] = d_2869_v81_
                nw489_[int(16)] = d_2869_v81_
                d_2870_v83_ = nw489_
                d_2872_v84_: _dafny.Map
                d_2872_v84_ = _dafny.Map({len(d_2866_v78_): d_2870_v83_})
                d_2873_v85_: D26
                d_2873_v85_ = D26_DC70(d_2750_v1_.f24, d_2825_cf78_, p0, ((d_2872_v84_)[d_2750_v1_.f24] if (d_2750_v1_.f24) in (d_2872_v84_) else d_2870_v83_))
                d_2874_v86_: _dafny.Map
                d_2874_v86_ = _dafny.Map({d_2873_v85_: (self).f18})
                d_2875_v87_: _dafny.Map
                d_2875_v87_ = _dafny.Map({d_2825_cf78_: d_2862_v74_})
                d_2876_v88_: _dafny.Array
                nw490_ = _dafny.Array(None, 27)
                nw490_[int(0)] = d_2862_v74_
                nw490_[int(1)] = (d_2863_v75_)[default__.safeIndex(d_2825_cf78_, len(d_2863_v75_))]
                nw490_[int(2)] = ((d_2864_v76_)[((d_2874_v86_)[d_2873_v85_] if (d_2873_v85_) in (d_2874_v86_) else (self).f18)] if (((d_2874_v86_)[d_2873_v85_] if (d_2873_v85_) in (d_2874_v86_) else (self).f18)) in (d_2864_v76_) else d_2862_v74_)
                nw490_[int(3)] = d_2862_v74_
                nw490_[int(4)] = d_2862_v74_
                nw490_[int(5)] = d_2862_v74_
                nw490_[int(6)] = d_2862_v74_
                nw490_[int(7)] = d_2862_v74_
                nw490_[int(8)] = d_2862_v74_
                nw490_[int(9)] = (d_2862_v74_ if False else d_2862_v74_)
                nw490_[int(10)] = d_2862_v74_
                nw490_[int(11)] = d_2862_v74_
                nw490_[int(12)] = d_2862_v74_
                nw490_[int(13)] = d_2862_v74_
                nw490_[int(14)] = d_2862_v74_
                nw490_[int(15)] = d_2862_v74_
                nw490_[int(16)] = d_2862_v74_
                nw490_[int(17)] = (d_2862_v74_ if p0 else d_2862_v74_)
                nw490_[int(18)] = d_2862_v74_
                nw490_[int(19)] = d_2862_v74_
                nw490_[int(20)] = d_2862_v74_
                nw490_[int(21)] = d_2862_v74_
                nw490_[int(22)] = ((d_2875_v87_)[d_2749_v0_] if (d_2749_v0_) in (d_2875_v87_) else d_2862_v74_)
                nw490_[int(23)] = d_2862_v74_
                nw490_[int(24)] = d_2862_v74_
                nw490_[int(25)] = d_2862_v74_
                nw490_[int(26)] = d_2862_v74_
                d_2876_v88_ = nw490_
                index464_ = default__.safeIndex(179, (d_2876_v88_).length(0))
                (d_2876_v88_)[index464_] = d_2862_v74_
                d_2877_v89_: _dafny.Map
                d_2877_v89_ = _dafny.Map({d_2750_v1_.f24: p0})
                d_2878_v90_: D17
                d_2878_v90_ = D17_DC41(d_2877_v89_)
                d_2879_v92_: D24
                def iife233_():
                    coll69_ = _dafny.Set()
                    compr_69_: D17
                    for compr_69_ in (_dafny.MultiSet([d_2878_v90_])).Elements:
                        d_2880_v91_: D17 = compr_69_
                        if (d_2880_v91_) in (_dafny.MultiSet([d_2878_v90_])):
                            coll69_ = coll69_.union(_dafny.Set([d_2880_v91_]))
                    return _dafny.Set(coll69_)
                d_2879_v92_ = D24_DC65(False, iife233_()
, (self).f18, d_2750_v1_.f24, d_2862_v74_)
                index465_ = default__.safeIndex(179, (d_2876_v88_).length(0))
                (d_2876_v88_)[index465_] = (d_2879_v92_).cf114
                d_2881_v93_: _dafny.Array
                nw491_ = _dafny.Array(None, 3)
                nw491_[int(0)] = d_2758_v7_
                nw491_[int(1)] = d_2758_v7_
                nw491_[int(2)] = d_2758_v7_
                d_2881_v93_ = nw491_
                d_2882_v94_: D22
                d_2882_v94_ = D22_DC59(D22_DC57(d_2881_v93_))
                d_2883_v95_: D22
                d_2883_v95_ = D22_DC59(d_2882_v94_)
                arr5_ = (d_2876_v88_)[default__.safeIndex(179, (d_2876_v88_).length(0))]
                index466_ = default__.safeIndex(885, ((d_2876_v88_)[default__.safeIndex(179, (d_2876_v88_).length(0))]).length(0))
                arr5_[index466_] = d_2750_v1_.f24
                arr6_ = (d_2876_v88_)[default__.safeIndex(179, (d_2876_v88_).length(0))]
                index467_ = default__.safeIndex(885, ((d_2876_v88_)[default__.safeIndex(179, (d_2876_v88_).length(0))]).length(0))
                rhs460_ = (self).f21
                rhs461_ = d_2883_v95_
                rhs462_ = d_2826_cf77_
                rhs463_ = (0) - (default__.safeModuloInt(d_2750_v1_.f24, (default__.fm1(d_2750_v1_.f24, p0, p0, globalState) if (self).f21 else d_2825_cf78_)))
                lhs402_ = (d_2876_v88_)[default__.safeIndex(179, (d_2876_v88_).length(0))]
                lhs403_ = default__.safeIndex(885, ((d_2876_v88_)[default__.safeIndex(179, (d_2876_v88_).length(0))]).length(0))
                r1 = rhs460_
                d_2883_v95_ = rhs461_
                d_2826_cf77_ = rhs462_
                lhs402_[lhs403_] = rhs463_
                d_2884_v96_: _dafny.Map
                d_2884_v96_ = _dafny.Map({449: 551})
                (globalState).f3 = len(d_2884_v96_)
                d_2885_v97_: _dafny.Map
                d_2885_v97_ = _dafny.Map({(self).f18: p0})
                d_2886_v98_: _dafny.Map
                d_2886_v98_ = _dafny.Map({d_2749_v0_: d_2885_v97_})
                d_2886_v98_ = (d_2886_v98_).set(d_2750_v1_.f24, d_2885_v97_)
            elif True:
                d_2887___mcc_h10_ = source28_.cf85
                d_2888_cf85_ = d_2887___mcc_h10_
                def iife234_():
                    coll70_ = _dafny.Set()
                    compr_70_: int
                    for compr_70_ in (_dafny.MultiSet([d_2749_v0_])).Elements:
                        d_2889_v99_: int = compr_70_
                        if (d_2889_v99_) in (_dafny.MultiSet([d_2749_v0_])):
                            coll70_ = coll70_.union(_dafny.Set([(d_2889_v99_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D17_DC43(_dafny.MultiSet([_dafny.Map({len(_dafny.Map({False: True})): -961})]), False)).cf76, False, True, False]), _dafny.SeqWithoutIsStrInference([False, False])]))))]))
                    return _dafny.Set(coll70_)
                (globalState).f3 = default__.safeModuloInt(d_2750_v1_.f24, (d_2825_cf78_ if p0 else len(iife234_()
                )))
                (globalState).f15 = d_2825_cf78_
                r1 = False
                d_2890_v100_: D25
                d_2890_v100_ = D25_DC67((self).f21, len((self).f19), (d_2755_v5_).f35, (self).f21)
                d_2891_v101_: _dafny.Array
                nw492_ = _dafny.Array(None, 6)
                nw492_[int(0)] = (d_2755_v5_).f35
                nw492_[int(1)] = (d_2755_v5_).f35
                nw492_[int(2)] = (d_2755_v5_).f35
                nw492_[int(3)] = ((d_2755_v5_).f35) + (_dafny.SeqWithoutIsStrInference([d_2750_v1_.f24, d_2825_cf78_]))
                nw492_[int(4)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2750_v1_.f24: True})) for d_2892_i9_ in range(default__.abs(606))])) + (_dafny.SeqWithoutIsStrInference([d_2750_v1_.f24]))
                nw492_[int(5)] = (d_2890_v100_).cf118
                d_2891_v101_ = nw492_
                index468_ = default__.safeIndex(452, (d_2891_v101_).length(0))
                (d_2891_v101_)[index468_] = ((d_2755_v5_).f35) + ((d_2890_v100_).cf118)
                index469_ = default__.safeIndex(452, (d_2891_v101_).length(0))
                def iife235_():
                    coll71_ = _dafny.Set()
                    compr_71_: int
                    for compr_71_ in (_dafny.SeqWithoutIsStrInference([len((d_2755_v5_).f35) for d_2893_i10_ in range(default__.abs(698))])).Elements:
                        d_2894_v102_: int = compr_71_
                        if (d_2894_v102_) in (_dafny.SeqWithoutIsStrInference([len((d_2755_v5_).f35) for d_2893_i10_ in range(default__.abs(698))])):
                            coll71_ = coll71_.union(_dafny.Set([(d_2894_v102_) - (341)]))
                    return _dafny.Set(coll71_)
                (d_2891_v101_)[index469_] = (_dafny.SeqWithoutIsStrInference([len((self).f19), len(_dafny.Map({(self).f21: d_2750_v1_.f24}))])).set(default__.safeIndex((d_2749_v0_) + (d_2750_v1_.f24), len(_dafny.SeqWithoutIsStrInference([len((self).f19), len(_dafny.Map({(self).f21: d_2750_v1_.f24}))]))), (d_2825_cf78_) + (len(iife235_()
                )))
            index470_ = default__.safeIndex(193, (d_2826_cf77_).length(0))
            (d_2826_cf77_)[index470_] = not(p0)
            index471_ = default__.safeIndex(193, (d_2826_cf77_).length(0))
            (d_2826_cf77_)[index471_] = p0
        elif True:
            d_2895___mcc_h5_ = source27_.cf73
            d_2896_cf73_ = d_2895___mcc_h5_
            d_2897_v103_: _dafny.Seq
            d_2897_v103_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_2898_v104_: _dafny.Map
            d_2898_v104_ = _dafny.Map({d_2763_v10_: d_2763_v10_})
            d_2899_v105_: _dafny.MultiSet
            d_2899_v105_ = _dafny.MultiSet([p0, (self).f18, not((self).f18)])
            d_2900_v106_: _dafny.Map
            d_2900_v106_ = _dafny.Map({(self).f18: d_2899_v105_})
            if (_dafny.MultiSet((d_2897_v103_ if (self).f18 else default__.fm44(d_2749_v0_, (self).f18, ((d_2898_v104_)[d_2763_v10_] if (d_2763_v10_) in (d_2898_v104_) else d_2763_v10_), 173, globalState)))) != (((d_2900_v106_)[p0] if (p0) in (d_2900_v106_) else d_2899_v105_)):
                d_2749_v0_ = (d_2750_v1_.f24) * ((d_2758_v7_).cardinality)
                d_2901_v107_: C5
                nw493_ = C5()
                nw493_.ctor__(((self).f19) + ((self).f19), (self).f19, False)
                d_2901_v107_ = nw493_
                r2 = (self).f18
                (globalState).f12 = (d_2901_v107_).f28
                (globalState).f0 = default__.safeModuloInt(675, 773)
            elif True:
                (d_2750_v1_).f24 = d_2750_v1_.f24
                (d_2750_v1_).m1((self).f19, d_2750_v1_.f24, (self).f19, globalState)
                rhs464_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnpnbbeug")) if (self).f18 else (self).f19)) + (default__.fm12(d_2750_v1_.f24, p0, globalState))
                rhs465_ = default__.safeDivisionInt(default__.safeModuloInt(d_2750_v1_.f24, d_2749_v0_), d_2750_v1_.f24)
                lhs404_ = globalState
                lhs405_ = d_2750_v1_
                lhs404_.f12 = rhs464_
                lhs405_.f24 = rhs465_
                r1 = (self).f18
                d_2902_v108_: _dafny.Array
                nw494_ = _dafny.Array(_dafny.MultiSet({}), 16)
                d_2902_v108_ = nw494_
                d_2902_v108_ = d_2902_v108_
            d_2903_v109_: D20
            d_2903_v109_ = D20_DC53((self).f18, (self).f18, d_2763_v10_)
            d_2904_v110_: _dafny.Array
            nw495_ = _dafny.Array(None, 24)
            nw495_[int(0)] = ((self).f19)[default__.safeIndex(d_2750_v1_.f24, len((self).f19))]
            nw495_[int(1)] = d_2763_v10_
            nw495_[int(2)] = _dafny.CodePoint('m')
            nw495_[int(3)] = d_2763_v10_
            nw495_[int(4)] = d_2763_v10_
            nw495_[int(5)] = d_2763_v10_
            nw495_[int(6)] = d_2763_v10_
            nw495_[int(7)] = d_2763_v10_
            nw495_[int(8)] = _dafny.CodePoint('a')
            nw495_[int(9)] = d_2763_v10_
            nw495_[int(10)] = d_2763_v10_
            nw495_[int(11)] = d_2763_v10_
            nw495_[int(12)] = ((self).f19)[default__.safeIndex((0) - (d_2749_v0_), len((self).f19))]
            nw495_[int(13)] = d_2763_v10_
            nw495_[int(14)] = d_2763_v10_
            nw495_[int(15)] = d_2763_v10_
            nw495_[int(16)] = (d_2763_v10_ if (self).f18 else d_2763_v10_)
            nw495_[int(17)] = d_2763_v10_
            nw495_[int(18)] = d_2763_v10_
            nw495_[int(19)] = (d_2903_v109_).cf89
            nw495_[int(20)] = d_2763_v10_
            nw495_[int(21)] = d_2763_v10_
            nw495_[int(22)] = d_2763_v10_
            nw495_[int(23)] = _dafny.CodePoint('k')
            d_2904_v110_ = nw495_
            d_2904_v110_ = d_2904_v110_
            d_2905_v111_: _dafny.Array
            nw496_ = _dafny.Array(False, 4)
            d_2905_v111_ = nw496_
            d_2906_v112_: _dafny.Map
            d_2906_v112_ = _dafny.Map({p0: d_2750_v1_.f24})
            d_2907_v113_: _dafny.Map
            d_2907_v113_ = _dafny.Map({d_2906_v112_: d_2749_v0_})
            d_2908_v114_: T1
            nw497_ = C3()
            nw497_.ctor__(_dafny.SeqWithoutIsStrInference([917]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), (self).f20, p0)
            d_2908_v114_ = nw497_
            index472_ = default__.safeIndex(238, (d_2905_v111_).length(0))
            (d_2905_v111_)[index472_] = (len(_dafny.Map({len(d_2907_v113_): d_2908_v114_}))) != (d_2750_v1_.f24)
            index473_ = default__.safeIndex(238, (d_2905_v111_).length(0))
            (d_2905_v111_)[index473_] = (d_2897_v103_)[default__.safeIndex(len((self).f19), len(d_2897_v103_))]
            d_2905_v111_ = d_2905_v111_
        r0 = True
        r1 = (self).f21
        d_2909_v115_: _dafny.Map
        d_2909_v115_ = _dafny.Map({d_2750_v1_.f24: (d_2755_v5_).fm2(p0, (self).f21, p0, globalState)})
        r2 = ((d_2909_v115_)[d_2750_v1_.f24] if (d_2750_v1_.f24) in (d_2909_v115_) else True)
        return r0, r1, r2

    @property
    def f21(self):
        return self._f21
