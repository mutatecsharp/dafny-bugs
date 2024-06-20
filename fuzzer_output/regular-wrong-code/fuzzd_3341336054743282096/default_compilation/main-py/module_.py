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
    def fm1(p0, p1, p2, p3, globalState):
        return (_dafny.Map({False: True})) | ((_dafny.Map({(D12_DC29(True, 620, False, len(_dafny.SeqWithoutIsStrInference([False])), False)).cf52: not(True)})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm2(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_0_i0_ in range(default__.abs(192))]): -91})).keys.Elements:
                d_1_v0_: _dafny.Seq = compr_0_
                if (d_1_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_0_i0_ in range(default__.abs(192))]): -91})):
                    coll0_[d_1_v0_] = 865
            return _dafny.Map(coll0_)
        return ((D9_DC22(len(iife0_()
))).cf38) <= (103)

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Seq
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_2_i0_ in range(default__.abs(895))])).Elements:
                d_3_v0_: _dafny.Seq = compr_1_
                if (d_3_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_2_i0_ in range(default__.abs(895))])):
                    coll1_[d_3_v0_] = _dafny.CodePoint('e')
            return _dafny.Map(coll1_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): _dafny.CodePoint('y')})) | (iife1_()
        )) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): _dafny.CodePoint('m')}))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return _dafny.Map({271: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_4_i0_ in range(default__.abs(-254))]))})

    @staticmethod
    def fm10(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(192, 644):
                d_5_v0_: int = compr_2_
                if ((192) <= (d_5_v0_)) and ((d_5_v0_) < (644)):
                    coll2_ = coll2_.union(_dafny.Set([(d_5_v0_) - (-269)]))
            return _dafny.Set(coll2_)
        return (_dafny.Set({810})).intersection((iife2_()
         if True else _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False])), -526})))

    @staticmethod
    def fm13(p0, p1, globalState):
        return _dafny.MultiSet([default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True]))]): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_6_i0_ in range(default__.abs(716))]))}))])), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_7_i1_ in range(default__.abs(38))])): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})) for d_8_i2_ in range(default__.abs(904))]))})))])

    @staticmethod
    def fm16(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))

    @staticmethod
    def fm17(globalState):
        return not(not(not(((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]) for d_9_i0_ in range(default__.abs(-780))])) < (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False])])) if True else (_dafny.Set({False})).ispropersubset(_dafny.Set({False}))))))

    @staticmethod
    def fm18(globalState):
        return (_dafny.MultiSet([_dafny.Set({not(not(True))})])).intersection((_dafny.MultiSet([_dafny.Set({True}), _dafny.Set({False, True})])) - ((D26_DC67(_dafny.MultiSet([_dafny.Set({False, not(False), True}), _dafny.Set({True, False})]))).cf111))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, True, False, False, not(False)])) + ((_dafny.SeqWithoutIsStrInference([True, False, False, False, False])) + (_dafny.SeqWithoutIsStrInference([False, False])))

    @staticmethod
    def fm20(p0, globalState):
        return (len((_dafny.Map({not(True): 291})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), _dafny.MultiSet([True]), _dafny.MultiSet([not(True), True])]))})))) + ((0) - (((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qeutyx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpui"))])) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ts")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkwix")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_10_i0_ in range(default__.abs(-889))])])))).cardinality))

    @staticmethod
    def fm21(globalState):
        if (False) == (True):
            return (_dafny.Map({817: True})) | (_dafny.Map({746: False}))
        elif True:
            return (_dafny.Map({494: True})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lno"))): True}))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return _dafny.Map({default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_11_i0_ in range(default__.abs(910))])), 905): -703})

    @staticmethod
    def fm23(p0, globalState):
        if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlbbvk")))) != ((0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False})))):
            return (_dafny.SeqWithoutIsStrInference([False, not(True), False, not(not(True)), True])) + (_dafny.SeqWithoutIsStrInference([not(True)]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([True, not(False), False])
        elif True:
            return _dafny.SeqWithoutIsStrInference([True])

    @staticmethod
    def fm24(p0, globalState):
        return _dafny.MultiSet([-72, -329])

    @staticmethod
    def fm25(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vy"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbgtm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pacqqiw"))))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return (D11_DC24(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le")): _dafny.CodePoint('s')}))).cf40

    @staticmethod
    def fm29(p0, globalState):
        if False:
            return _dafny.CodePoint('m')
        elif True:
            return _dafny.CodePoint('t')

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm31(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roi"))

    @staticmethod
    def fm32(p0, globalState):
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(89, 359):
                    d_12_v1_: int = compr_5_
                    if ((89) <= (d_12_v1_)) and ((d_12_v1_) < (359)):
                        coll5_[(d_12_v1_) - (952)] = False
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(89, 359):
                    d_12_v1_: int = compr_4_
                    if ((89) <= (d_12_v1_)) and ((d_12_v1_) < (359)):
                        coll4_[(d_12_v1_) - (952)] = False
                return _dafny.Map(coll4_)
            compr_3_: int
            for compr_3_ in (_dafny.Map({len(iife4_()
            ): len(_dafny.SeqWithoutIsStrInference([False]))})).keys.Elements:
                d_13_v0_: int = compr_3_
                if (d_13_v0_) in (_dafny.Map({len(iife5_()
                ): len(_dafny.SeqWithoutIsStrInference([False]))})):
                    coll3_[(d_13_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqy")))))] = 326
            return _dafny.Map(coll3_)
        return len(iife3_()
        )

    @staticmethod
    def fm33(p0, p1, globalState):
        return D6_DC14((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csnarsqlx"))))

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        if not((len(_dafny.Set({len(_dafny.Map({True: _dafny.CodePoint('i')}))}))) > (len(_dafny.SeqWithoutIsStrInference([550, 706])))):
            return D12_DC28(True, -395)
        elif True:
            return D12_DC28(not(False), len(_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm35(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([325])

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({_dafny.CodePoint('q'): -65})) | (_dafny.Map({_dafny.CodePoint('i'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xaxw")))}))) | (_dafny.Map({_dafny.CodePoint('d'): len(_dafny.Map({_dafny.Set({not(False)}): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utuis")))}))}))

    @staticmethod
    def fm37(p0, globalState):
        return D7_DC19(236)

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([True, True]))) + (_dafny.SeqWithoutIsStrInference([True, False]))

    @staticmethod
    def fm39(p0, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(888, 779):
                d_14_v0_: int = compr_6_
                if ((888) <= (d_14_v0_)) and ((d_14_v0_) < (779)):
                    coll6_[default__.safeModuloInt(d_14_v0_, 778)] = not(not(False))
            return _dafny.Map(coll6_)
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([376])).Elements:
                d_15_v1_: int = compr_7_
                if (d_15_v1_) in (_dafny.SeqWithoutIsStrInference([376])):
                    coll7_[(d_15_v1_) - (len(_dafny.SeqWithoutIsStrInference([True])))] = False
            return _dafny.Map(coll7_)
        return ((iife6_()
        ) | ((D28_DC71(iife7_()
)).cf116)) | ((_dafny.Map({884: False})) | (_dafny.Map({(0) - ((_dafny.MultiSet([True, False])).cardinality): False})))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        if True:
            return D12_DC27(_dafny.Set({True}))
        elif True:
            return D12_DC27(_dafny.Set({False, True}))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return _dafny.Map({(len(_dafny.Map({997: not(False)}))) + ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwotwnwen"))))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amno"))})

    @staticmethod
    def fm44(p0, globalState):
        return 862

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        return D14_DC35((656) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True, True]))]))), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_16_i0_ in range(default__.abs(65))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_17_i1_ in range(default__.abs(584))])), (_dafny.MultiSet([True, not(True)])).cardinality, (_dafny.MultiSet([not(True)])).issubset(_dafny.MultiSet([not(True)])), (len(_dafny.SeqWithoutIsStrInference([False, True]))) >= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edklwnr")))))

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_18_i0_ in range(default__.abs(493))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdfbv")))

    @staticmethod
    def fm47(p0, p1, globalState):
        if (-789) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('i')])).cardinality])):
            return (_dafny.Map({False: not(True)})) | (_dafny.Map({True: False}))
        elif True:
            return _dafny.Map({not(False): False})

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(958, 817):
                d_20_v0_: int = compr_8_
                if ((958) <= (d_20_v0_)) and ((d_20_v0_) < (817)):
                    coll8_[default__.safeDivisionInt(d_20_v0_, 713)] = 855
            return _dafny.Map(coll8_)
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([not(False), False, False]): -383})) for d_19_i0_ in range(default__.abs(395))]))) | ((_dafny.MultiSet([817])).intersection(_dafny.MultiSet([len(iife8_()
        ), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjajll")))])))

    @staticmethod
    def fm49(p0, globalState):
        return D0_DC1((814 if False else len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False]), _dafny.MultiSet([not(False)])]))))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(970, 345):
                d_21_v0_: int = compr_9_
                if ((970) <= (d_21_v0_)) and ((d_21_v0_) < (345)):
                    coll9_[default__.safeModuloInt(d_21_v0_, 169)] = not(False)
            return _dafny.Map(coll9_)
        return iife9_()
        

    @staticmethod
    def fm51(p0, p1, p2, globalState):
        return _dafny.CodePoint('k')

    @staticmethod
    def fm52(p0, p1, p2, globalState):
        if (True if False else True):
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: str
                for compr_10_ in (_dafny.Map({_dafny.CodePoint('j'): (D20_DC53(False)).cf94})).keys.Elements:
                    d_22_v0_: str = compr_10_
                    if (d_22_v0_) in (_dafny.Map({_dafny.CodePoint('j'): (D20_DC53(False)).cf94})):
                        coll10_[d_22_v0_] = 935
                return _dafny.Map(coll10_)
            return _dafny.Map({_dafny.CodePoint('d'): (0) - (len(iife10_()
            ))})
        elif True:
            return _dafny.Map({_dafny.CodePoint('s'): len(_dafny.SeqWithoutIsStrInference([910]))})

    @staticmethod
    def fm53(p0, globalState):
        return _dafny.Set({(len(_dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "souhprl"))), len(_dafny.Map({len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywax")))})): False})), 426, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "liy"))), (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])), -525])))}): False}))) >= (981), (not(not(False))) == (True), not(not (False) or (True)), True})

    @staticmethod
    def fm54(p0, p1, globalState):
        source0_ = D11_DC25(_dafny.CodePoint('t'), 75, True)
        if source0_.is_DC25:
            d_23___mcc_h0_ = source0_.cf41
            d_24___mcc_h1_ = source0_.cf42
            d_25___mcc_h2_ = source0_.cf43
            d_26_cf43_ = d_25___mcc_h2_
            d_27_cf42_ = d_24___mcc_h1_
            d_28_cf41_ = d_23___mcc_h0_
            if d_26_cf43_:
                return _dafny.Map({d_28_cf41_: d_27_cf42_})
            elif True:
                def iife11_():
                    coll11_ = _dafny.Map()
                    compr_11_: str
                    for compr_11_ in (_dafny.SeqWithoutIsStrInference([d_28_cf41_])).Elements:
                        d_29_v0_: str = compr_11_
                        if (d_29_v0_) in (_dafny.SeqWithoutIsStrInference([d_28_cf41_])):
                            coll11_[d_29_v0_] = d_27_cf42_
                    return _dafny.Map(coll11_)
                return iife11_()
                
        elif source0_.is_DC24:
            d_30___mcc_h3_ = source0_.cf40
            d_31_cf40_ = d_30___mcc_h3_
            if False:
                return _dafny.Map({_dafny.CodePoint('t'): (_dafny.MultiSet([False, True])).cardinality})
            elif True:
                return _dafny.Map({_dafny.CodePoint('q'): 997})
        elif True:
            d_32___mcc_h4_ = source0_.cf44
            d_33_cf44_ = d_32___mcc_h4_
            return _dafny.Map({_dafny.CodePoint('n'): len(_dafny.Map({True: not(True)}))})

    @staticmethod
    def fm55(p0, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, False, True, True]), _dafny.SeqWithoutIsStrInference([True, False, True]), _dafny.SeqWithoutIsStrInference([not(True), True, False, False, True])})) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([not(False), False])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, False, False])})))

    @staticmethod
    def fm56(p0, globalState):
        def iife12_():
            def iife14_():
                coll14_ = _dafny.Set()
                compr_14_: str
                for compr_14_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d')])).Elements:
                    d_36_v1_: str = compr_14_
                    if (d_36_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d')])):
                        coll14_ = coll14_.union(_dafny.Set([d_36_v1_]))
                return _dafny.Set(coll14_)
            coll12_ = _dafny.Map()
            def iife13_():
                coll13_ = _dafny.Set()
                compr_13_: str
                for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d')])).Elements:
                    d_34_v1_: str = compr_13_
                    if (d_34_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('f'), _dafny.CodePoint('o'), _dafny.CodePoint('y'), _dafny.CodePoint('d')])):
                        coll13_ = coll13_.union(_dafny.Set([d_34_v1_]))
                return _dafny.Set(coll13_)
            compr_12_: _dafny.Seq
            for compr_12_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(False)]): iife13_()
            })).keys.Elements:
                d_35_v0_: _dafny.Seq = compr_12_
                if (d_35_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(False)]): iife14_()
                })):
                    coll12_[d_35_v0_] = 421
            return _dafny.Map(coll12_)
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: _dafny.Seq
            for compr_15_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([not(False)])})).Elements:
                d_37_v2_: _dafny.Seq = compr_15_
                if (d_37_v2_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([not(False)])})):
                    coll15_[d_37_v2_] = 753
            return _dafny.Map(coll15_)
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: _dafny.Seq
            for compr_16_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, False])})).Elements:
                d_38_v3_: _dafny.Seq = compr_16_
                if (d_38_v3_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, False])})):
                    coll16_[d_38_v3_] = -829
            return _dafny.Map(coll16_)
        return ((iife12_()
        ) | (iife15_()
        )) | (iife16_()
        )

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        return (_dafny.Map({not(False): _dafny.CodePoint('g')})) | (_dafny.Map({True: _dafny.CodePoint('u')}))

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        return _dafny.MultiSet([_dafny.CodePoint('o')])

    @staticmethod
    def fm59(globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsdwwpuv"))})

    @staticmethod
    def fm60(p0, p1, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: D11
            for compr_17_ in ((_dafny.Set({D11_DC25(_dafny.CodePoint('o'), -847, True), D11_DC25(_dafny.CodePoint('d'), -798, True)}) if not(True) else _dafny.Set({D11_DC25(_dafny.CodePoint('h'), (0) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_39_i0_ in range(default__.abs(500))]))}))), True)}))).Elements:
                d_40_v0_: D11 = compr_17_
                if (d_40_v0_) in ((_dafny.Set({D11_DC25(_dafny.CodePoint('o'), -847, True), D11_DC25(_dafny.CodePoint('d'), -798, True)}) if not(True) else _dafny.Set({D11_DC25(_dafny.CodePoint('h'), (0) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_39_i0_ in range(default__.abs(500))]))}))), True)}))):
                    coll17_[d_40_v0_] = -489
            return _dafny.Map(coll17_)
        return iife17_()
        

    @staticmethod
    def fm61(p0, p1, globalState):
        if not(True):
            return D9_DC21(_dafny.MultiSet([False]))
        elif True:
            return D9_DC21(_dafny.MultiSet([True, True]))

    @staticmethod
    def fm62(p0, p1, p2, p3, globalState):
        return D15_DC39(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "liockxaf")), D9_DC21(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm63(p0, p1, globalState):
        return ((_dafny.MultiSet([True])).intersection(_dafny.MultiSet([not(True), True]))).intersection((_dafny.MultiSet([True, not(not(False))])) | (_dafny.MultiSet([True])))

    @staticmethod
    def fm64(globalState):
        return D4_DC10(default__.safeModuloInt(499, 664), ((0) - (len(_dafny.SeqWithoutIsStrInference([845, (_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-856])): True}))])).cardinality])))) - (575), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ga"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bywhwnw")))))

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        return (_dafny.Map({True: 704})) | ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True, True, True]))})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdjiex")))})))

    @staticmethod
    def fm66(p0, p1, globalState):
        return D16_DC42(default__.safeModuloInt(149, len(_dafny.Set({-16}))), (0) - ((0) - (len((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([True, False]))))), default__.safeModuloInt(318, -834))

    @staticmethod
    def fm67(p0, globalState):
        return D11_DC25(_dafny.CodePoint('a'), (0) - (len((_dafny.SeqWithoutIsStrInference([127 for d_41_i0_ in range(default__.abs(691))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})), len(_dafny.Set({True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhdvmehq")))])))), True)

    @staticmethod
    def fm68(p0, p1, globalState):
        return D14_DC34(_dafny.Map({-61: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vj"))}))

    @staticmethod
    def fm69(p0, globalState):
        return _dafny.Map({(0) - (((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfcecnt")))) if not((D20_DC54(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irhb")), len(_dafny.SeqWithoutIsStrInference([not(True)])), True)).cf97) else len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cicqjbm")))])))): D15_DC37(_dafny.Set({_dafny.SeqWithoutIsStrInference([313 for d_42_i0_ in range(default__.abs(484))]), _dafny.SeqWithoutIsStrInference([-802])}))})

    @staticmethod
    def fm70(p0, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(979, 837):
                d_43_v0_: int = compr_18_
                if ((979) <= (d_43_v0_)) and ((d_43_v0_) < (837)):
                    coll18_[(d_43_v0_) - (-243)] = len(_dafny.Map({True: len(_dafny.Map({True: _dafny.CodePoint('g')}))}))
            return _dafny.Map(coll18_)
        source1_ = D16_DC43(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqqftomtx")), _dafny.CodePoint('v'), (0) - (len(iife18_()
)))
        if source1_.is_DC42:
            d_44___mcc_h0_ = source1_.cf76
            d_45___mcc_h1_ = source1_.cf77
            d_46___mcc_h2_ = source1_.cf78
            d_47_cf78_ = d_46___mcc_h2_
            d_48_cf77_ = d_45___mcc_h1_
            d_49_cf76_ = d_44___mcc_h0_
            def iife19_():
                def iife21_():
                    coll21_ = _dafny.Map()
                    compr_21_: str
                    for compr_21_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_50_i0_ in range(default__.abs(782))])).Elements:
                        d_51_v1_: str = compr_21_
                        if (d_51_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_50_i0_ in range(default__.abs(782))])):
                            coll21_[d_51_v1_] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttkmplyt")) for d_52_i1_ in range(default__.abs(-369))]))
                    return _dafny.Map(coll21_)
                coll19_ = _dafny.Set()
                def iife20_():
                    coll20_ = _dafny.Map()
                    compr_20_: str
                    for compr_20_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_50_i0_ in range(default__.abs(782))])).Elements:
                        d_51_v1_: str = compr_20_
                        if (d_51_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_50_i0_ in range(default__.abs(782))])):
                            coll20_[d_51_v1_] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttkmplyt")) for d_52_i1_ in range(default__.abs(-369))]))
                    return _dafny.Map(coll20_)
                compr_19_: _dafny.Map
                for compr_19_ in (_dafny.SeqWithoutIsStrInference([iife20_()
                ])).Elements:
                    d_53_v2_: _dafny.Map = compr_19_
                    if (d_53_v2_) in (_dafny.SeqWithoutIsStrInference([iife21_()
                    ])):
                        coll19_ = coll19_.union(_dafny.Set([d_53_v2_]))
                return _dafny.Set(coll19_)
            return iife19_()
            
        elif source1_.is_DC43:
            d_54___mcc_h3_ = source1_.cf79
            d_55___mcc_h4_ = source1_.cf80
            d_56___mcc_h5_ = source1_.cf81
            d_57_cf81_ = d_56___mcc_h5_
            d_58_cf80_ = d_55___mcc_h4_
            d_59_cf79_ = d_54___mcc_h3_
            return (_dafny.Set({_dafny.Map({d_58_cf80_: len(d_59_cf79_)})})).intersection(_dafny.Set({_dafny.Map({d_58_cf80_: d_57_cf81_})}))
        elif source1_.is_DC44:
            d_60___mcc_h6_ = source1_.cf82
            d_61_cf82_ = d_60___mcc_h6_
            def iife22_():
                def iife24_():
                    coll24_ = _dafny.Map()
                    compr_24_: str
                    for compr_24_ in (_dafny.MultiSet([_dafny.CodePoint('u')])).Elements:
                        d_62_v3_: str = compr_24_
                        if (d_62_v3_) in (_dafny.MultiSet([_dafny.CodePoint('u')])):
                            coll24_[d_62_v3_] = d_61_cf82_
                    return _dafny.Map(coll24_)
                coll22_ = _dafny.Set()
                def iife23_():
                    coll23_ = _dafny.Map()
                    compr_23_: str
                    for compr_23_ in (_dafny.MultiSet([_dafny.CodePoint('u')])).Elements:
                        d_62_v3_: str = compr_23_
                        if (d_62_v3_) in (_dafny.MultiSet([_dafny.CodePoint('u')])):
                            coll23_[d_62_v3_] = d_61_cf82_
                    return _dafny.Map(coll23_)
                compr_22_: _dafny.Map
                for compr_22_ in (_dafny.Set({_dafny.Map({_dafny.CodePoint('d'): d_61_cf82_}), iife23_()
                })).Elements:
                    d_63_v4_: _dafny.Map = compr_22_
                    if (d_63_v4_) in (_dafny.Set({_dafny.Map({_dafny.CodePoint('d'): d_61_cf82_}), iife24_()
                    })):
                        coll22_ = coll22_.union(_dafny.Set([d_63_v4_]))
                return _dafny.Set(coll22_)
            def iife25_():
                coll25_ = _dafny.Set()
                compr_25_: _dafny.Map
                for compr_25_ in (_dafny.Set({_dafny.Map({_dafny.CodePoint('a'): 815})})).Elements:
                    d_64_v5_: _dafny.Map = compr_25_
                    if (d_64_v5_) in (_dafny.Set({_dafny.Map({_dafny.CodePoint('a'): 815})})):
                        coll25_ = coll25_.union(_dafny.Set([d_64_v5_]))
                return _dafny.Set(coll25_)
            return (iife22_()
            ).intersection(iife25_()
            )
        elif True:
            d_65___mcc_h7_ = source1_.cf75
            d_66_cf75_ = d_65___mcc_h7_
            def iife26_():
                coll26_ = _dafny.Map()
                compr_26_: str
                for compr_26_ in (_dafny.Map({_dafny.CodePoint('p'): False})).keys.Elements:
                    d_67_v6_: str = compr_26_
                    if (d_67_v6_) in (_dafny.Map({_dafny.CodePoint('p'): False})):
                        coll26_[d_67_v6_] = (0) - (-297)
                return _dafny.Map(coll26_)
            def iife27_():
                coll27_ = _dafny.Map()
                compr_27_: str
                for compr_27_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])).Elements:
                    d_68_v7_: str = compr_27_
                    if (d_68_v7_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])):
                        coll27_[d_68_v7_] = len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_69_i2_ in range(default__.abs(-771))]))}))
                return _dafny.Map(coll27_)
            return (_dafny.Set({_dafny.Map({_dafny.CodePoint('k'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imekcx")))}), _dafny.Map({_dafny.CodePoint('h'): -246}), iife26_()
            })) | (_dafny.Set({_dafny.Map({_dafny.CodePoint('d'): 85}), iife27_()
            }))

    @staticmethod
    def fm71(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, not(False)]), _dafny.SeqWithoutIsStrInference([not(True)]), (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True])), (_dafny.SeqWithoutIsStrInference([False, True, True])) + (_dafny.SeqWithoutIsStrInference([not(True)]))])

    @staticmethod
    def fm72(globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_70_i0_ in range(default__.abs(395))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usta")))])

    @staticmethod
    def Main(noArgsParameter__):
        d_71_globalState_: GlobalState
        nw0_ = GlobalState()
        nw0_.ctor__()
        d_71_globalState_ = nw0_
        d_72_v0_: int
        d_72_v0_ = 936
        d_73_v1_: C4
        nw1_ = C4()
        nw1_.ctor__()
        d_73_v1_ = nw1_
        d_74_v3_: _dafny.Seq
        d_74_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtn"))
        d_75_v4_: _dafny.MultiSet
        d_75_v4_ = _dafny.MultiSet([d_74_v3_, d_74_v3_])
        d_76_v5_: _dafny.Array
        def lambda0_(d_77_i0_):
            return _dafny.Set({True, True})

        init0_ = lambda0_
        nw2_ = _dafny.Array(None, 5)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_76_v5_ = nw2_
        d_78_v6_: D0
        d_78_v6_ = D0_DC0(d_76_v5_, d_72_v0_)
        d_79_v7_: _dafny.Seq
        d_79_v7_ = _dafny.SeqWithoutIsStrInference([d_72_v0_])
        d_80_v8_: D0
        d_80_v8_ = D0_DC1(d_72_v0_)
        d_81_v9_: C8
        nw3_ = C8()
        def iife28_():
            coll28_ = _dafny.Map()
            compr_28_: _dafny.Seq
            for compr_28_ in (d_75_v4_).Elements:
                d_82_v2_: _dafny.Seq = compr_28_
                if (d_82_v2_) in (d_75_v4_):
                    coll28_[d_82_v2_] = -676
            return _dafny.Map(coll28_)
        nw3_.ctor__(default__.safeDivisionInt(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_72_v0_])): d_73_v1_})), len(iife28_()
        )), d_78_v6_, (d_79_v7_) + (d_79_v7_), d_80_v8_)
        d_81_v9_ = nw3_
        d_83_v10_: bool
        d_83_v10_ = False
        pat_let_tv0_ = d_72_v0_
        d_84_v11_: int
        d_85_v12_: _dafny.Map
        d_86_v13_: int
        out0_: int
        out1_: _dafny.Map
        out2_: int
        def iife29_(_pat_let0_0):
            def iife30_(d_87_dt__update__tmp_h0_):
                def iife31_(_pat_let1_0):
                    def iife32_(d_88_dt__update_hcf1_h0_):
                        return D0_DC0((d_87_dt__update__tmp_h0_).cf0, d_88_dt__update_hcf1_h0_)
                    return iife32_(_pat_let1_0)
                return iife31_(pat_let_tv0_)
            return iife30_(_pat_let0_0)
        out0_, out1_, out2_ = (d_81_v9_).m3(490, iife29_(D0_DC0(d_76_v5_, len(d_74_v3_))), d_83_v10_, d_71_globalState_)
        d_84_v11_ = out0_
        d_85_v12_ = out1_
        d_86_v13_ = out2_
        d_89_v14_: _dafny.Map
        d_89_v14_ = _dafny.Map({d_83_v10_: 548})
        d_90_v15_: _dafny.Map
        d_90_v15_ = _dafny.Map({d_89_v14_: d_74_v3_})
        d_91_v16_: D18
        d_91_v16_ = D18_DC47(d_90_v15_)
        pat_let_tv1_ = d_91_v16_
        def iife33_(_pat_let2_0):
            def iife34_(d_92_dt__update__tmp_h1_):
                def iife35_(_pat_let3_0):
                    def iife36_(d_93_dt__update_hcf85_h0_):
                        return D18_DC47(d_93_dt__update_hcf85_h0_)
                    return iife36_(_pat_let3_0)
                return iife35_((pat_let_tv1_).cf85)
            return iife34_(_pat_let2_0)
        source2_ = (iife33_(d_91_v16_) if d_83_v10_ else d_91_v16_)
        if source2_.is_DC48:
            if d_83_v10_:
                d_94_v17_: _dafny.Array
                nw4_ = _dafny.Array(_dafny.Map({}), 8)
                d_94_v17_ = nw4_
                d_94_v17_ = d_94_v17_
                d_95_v18_: _dafny.Array
                nw5_ = _dafny.Array(False, 14)
                d_95_v18_ = nw5_
                index0_ = default__.safeIndex(213, (d_95_v18_).length(0))
                (d_95_v18_)[index0_] = d_83_v10_
                d_96_v19_: _dafny.MultiSet
                d_96_v19_ = _dafny.MultiSet([d_84_v11_])
                pat_let_tv2_ = d_81_v9_
                d_97_v20_: T1
                nw6_ = C7()
                def iife37_(_pat_let4_0):
                    def iife38_(d_98_dt__update__tmp_h2_):
                        def iife39_(_pat_let5_0):
                            def iife40_(d_99_dt__update_hcf2_h0_):
                                return D0_DC1(d_99_dt__update_hcf2_h0_)
                            return iife40_(_pat_let5_0)
                        return iife39_((pat_let_tv2_).f2)
                    return iife38_(_pat_let4_0)
                nw6_.ctor__(d_96_v19_, iife37_(d_80_v8_), (d_79_v7_) + (d_79_v7_), D0_DC1(d_72_v0_))
                d_97_v20_ = nw6_
                d_100_v21_: _dafny.Array
                nw7_ = _dafny.Array(int(0), 17)
                d_100_v21_ = nw7_
                index1_ = default__.safeIndex(721, (d_100_v21_).length(0))
                (d_100_v21_)[index1_] = len(default__.fm70(d_83_v10_, d_71_globalState_))
                d_101_v22_: str
                d_101_v22_ = _dafny.CodePoint('e')
                d_102_v23_: D16
                d_102_v23_ = D16_DC43(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_103_i1_ in range(default__.abs(403))]), d_101_v22_, d_84_v11_)
                d_104_v24_: C0
                nw8_ = C0()
                nw8_.ctor__((d_102_v23_).cf79)
                d_104_v24_ = nw8_
                d_105_v25_: _dafny.Map
                d_105_v25_ = _dafny.Map({d_104_v24_: d_83_v10_})
                index2_ = default__.safeIndex(213, (d_95_v18_).length(0))
                index3_ = default__.safeIndex(721, (d_100_v21_).length(0))
                rhs0_ = d_83_v10_
                rhs1_ = (d_97_v20_ if ((d_105_v25_)[d_104_v24_] if (d_104_v24_) in (d_105_v25_) else d_83_v10_) else d_97_v20_)
                rhs2_ = d_95_v18_
                rhs3_ = (d_81_v9_).f2
                rhs4_ = ((default__.fm44(True, d_71_globalState_)) < (d_86_v13_)) and (d_83_v10_)
                lhs0_ = d_95_v18_
                lhs1_ = default__.safeIndex(213, (d_95_v18_).length(0))
                lhs2_ = d_100_v21_
                lhs3_ = default__.safeIndex(721, (d_100_v21_).length(0))
                lhs0_[lhs1_] = rhs0_
                d_97_v20_ = rhs1_
                d_95_v18_ = rhs2_
                lhs2_[lhs3_] = rhs3_
                d_83_v10_ = rhs4_
                index4_ = default__.safeIndex(213, (d_95_v18_).length(0))
                (d_95_v18_)[index4_] = not ((d_83_v10_) not in (d_89_v14_)) or (True)
                d_106_v26_: _dafny.Map
                d_106_v26_ = _dafny.Map({(d_95_v18_)[default__.safeIndex(213, (d_95_v18_).length(0))]: default__.fm17(d_71_globalState_)})
                d_107_v27_: _dafny.Seq
                d_107_v27_ = _dafny.SeqWithoutIsStrInference([(d_95_v18_)[default__.safeIndex(213, (d_95_v18_).length(0))], False, d_83_v10_, d_83_v10_, (d_95_v18_)[default__.safeIndex(213, (d_95_v18_).length(0))]])
                d_108_v28_: _dafny.MultiSet
                d_108_v28_ = _dafny.MultiSet([(d_107_v27_)[default__.safeIndex(d_86_v13_, len(d_107_v27_))]])
                d_109_v29_: _dafny.Map
                d_109_v29_ = _dafny.Map({(d_81_v9_).f2: ((_dafny.MultiSet([not(((d_106_v26_)[d_83_v10_] if (d_83_v10_) in (d_106_v26_) else (d_95_v18_)[default__.safeIndex(213, (d_95_v18_).length(0))]))])) - (d_108_v28_)).cardinality})
                d_109_v29_ = (d_109_v29_).set(d_86_v13_, (d_81_v9_).f2)
                d_110_v30_: bool
                d_111_v31_: bool
                out3_: bool
                out4_: bool
                out3_, out4_ = (d_81_v9_).m1(((d_109_v29_)[(d_100_v21_)[default__.safeIndex(721, (d_100_v21_).length(0))]] if ((d_100_v21_)[default__.safeIndex(721, (d_100_v21_).length(0))]) in (d_109_v29_) else (d_102_v23_).cf81), d_71_globalState_)
                d_110_v30_ = out3_
                d_111_v31_ = out4_
            elif True:
                d_112_v32_: str
                d_112_v32_ = _dafny.CodePoint('c')
                d_113_v33_: _dafny.Map
                d_113_v33_ = _dafny.Map({d_112_v32_: (0) - (default__.fm32((d_81_v9_).f2, d_71_globalState_))})
                d_114_v34_: _dafny.Array
                nw9_ = _dafny.Array(int(0), 27)
                d_114_v34_ = nw9_
                d_115_v35_: T1
                nw10_ = C5()
                nw10_.ctor__(d_114_v34_, d_79_v7_, d_80_v8_)
                d_115_v35_ = nw10_
                d_116_v36_: C0
                nw11_ = C0()
                nw11_.ctor__(d_74_v3_)
                d_116_v36_ = nw11_
                d_117_v37_: _dafny.Map
                d_117_v37_ = _dafny.Map({d_115_v35_: d_116_v36_})
                d_118_v38_: _dafny.Array
                nw12_ = _dafny.Array(None, 29)
                nw12_[int(0)] = False
                nw12_[int(1)] = d_83_v10_
                nw12_[int(2)] = d_83_v10_
                nw12_[int(3)] = d_83_v10_
                nw12_[int(4)] = d_83_v10_
                nw12_[int(5)] = d_83_v10_
                nw12_[int(6)] = d_83_v10_
                nw12_[int(7)] = d_83_v10_
                nw12_[int(8)] = d_83_v10_
                nw12_[int(9)] = d_83_v10_
                nw12_[int(10)] = d_83_v10_
                nw12_[int(11)] = d_83_v10_
                nw12_[int(12)] = d_83_v10_
                nw12_[int(13)] = True
                nw12_[int(14)] = d_83_v10_
                nw12_[int(15)] = not(d_83_v10_)
                nw12_[int(16)] = default__.fm17(d_71_globalState_)
                nw12_[int(17)] = True
                nw12_[int(18)] = True
                nw12_[int(19)] = d_83_v10_
                nw12_[int(20)] = d_83_v10_
                nw12_[int(21)] = False
                nw12_[int(22)] = d_83_v10_
                nw12_[int(23)] = d_83_v10_
                nw12_[int(24)] = d_83_v10_
                nw12_[int(25)] = True
                nw12_[int(26)] = d_83_v10_
                nw12_[int(27)] = d_83_v10_
                nw12_[int(28)] = d_83_v10_
                d_118_v38_ = nw12_
                d_119_v39_: bool
                d_120_v40_: int
                out5_: bool
                out6_: int
                out5_, out6_ = (d_73_v1_).m10(((d_81_v9_).f2) + ((d_81_v9_).f2), len((d_113_v33_).set(default__.fm51(d_83_v10_, (d_81_v9_).f2, len(d_117_v37_), d_71_globalState_), d_72_v0_)), d_83_v10_, d_118_v38_, d_71_globalState_)
                d_119_v39_ = out5_
                d_120_v40_ = out6_
                d_121_v41_: _dafny.Seq
                d_121_v41_ = _dafny.SeqWithoutIsStrInference([d_83_v10_])
                d_121_v41_ = _dafny.SeqWithoutIsStrInference([(d_119_v39_) == (d_83_v10_), d_119_v39_, False, d_83_v10_])
                d_84_v11_ = -509
                d_122_v42_: _dafny.Seq
                d_122_v42_ = _dafny.SeqWithoutIsStrInference([d_118_v38_, d_118_v38_, d_118_v38_])
                d_83_v10_ = (_dafny.SeqWithoutIsStrInference([d_118_v38_, d_118_v38_])) <= (d_122_v42_)
                d_123_v43_: _dafny.Map
                d_123_v43_ = _dafny.Map({d_83_v10_: False})
                d_124_v44_: int
                d_125_v45_: _dafny.Map
                d_126_v46_: int
                out7_: int
                out8_: _dafny.Map
                out9_: int
                out7_, out8_, out9_ = (d_81_v9_).m3((0) - ((0) - (d_86_v13_)), d_78_v6_, ((d_123_v43_)[False] if (False) in (d_123_v43_) else not(d_83_v10_)), d_71_globalState_)
                d_124_v44_ = out7_
                d_125_v45_ = out8_
                d_126_v46_ = out9_
            d_127_v47_: _dafny.Seq
            d_127_v47_ = _dafny.SeqWithoutIsStrInference([not(d_83_v10_), d_83_v10_])
            d_128_v48_: _dafny.Set
            d_128_v48_ = _dafny.Set({_dafny.CodePoint('f')})
            d_129_v49_: _dafny.Map
            d_129_v49_ = _dafny.Map({d_127_v47_: (d_127_v47_).set(default__.safeIndex(len(d_128_v48_), len(d_127_v47_)), d_83_v10_)})
            d_129_v49_ = (_dafny.Map({d_127_v47_: d_127_v47_})) | ((d_129_v49_) | (d_129_v49_))
            d_83_v10_ = d_83_v10_
            d_130_v50_: T0
            nw13_ = C4()
            nw13_.ctor__()
            d_130_v50_ = nw13_
            d_131_v51_: D20
            d_131_v51_ = D20_DC53(d_83_v10_)
            d_132_v52_: _dafny.Map
            d_132_v52_ = _dafny.Map({d_130_v50_: (d_131_v51_).cf94})
            d_133_v53_: T0
            d_133_v53_ = d_130_v50_
            d_132_v52_ = (d_132_v52_).set((d_133_v53_), True)
        elif source2_.is_DC49:
            d_134___mcc_h0_ = source2_.cf86
            d_135_cf86_ = d_134___mcc_h0_
            if d_83_v10_:
                d_136_v54_: str
                d_136_v54_ = _dafny.CodePoint('n')
                d_137_v55_: _dafny.Map
                d_137_v55_ = _dafny.Map({d_136_v54_: d_80_v8_})
                d_138_v56_: _dafny.Map
                d_138_v56_ = _dafny.Map({True: default__.fm71((d_81_v9_).fm5((d_81_v9_).f2, d_83_v10_, d_137_v55_, d_83_v10_, d_71_globalState_), d_71_globalState_)})
                d_139_v57_: _dafny.Seq
                d_139_v57_ = _dafny.SeqWithoutIsStrInference([d_83_v10_, d_83_v10_, d_83_v10_])
                d_140_v58_: _dafny.MultiSet
                d_140_v58_ = _dafny.MultiSet([d_83_v10_, d_83_v10_])
                d_141_v59_: _dafny.Seq
                d_141_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_83_v10_, not(d_83_v10_)])])
                d_84_v11_ = len(((d_138_v56_)[(d_140_v58_).ispropersubset(_dafny.MultiSet(d_139_v57_))] if ((d_140_v58_).ispropersubset(_dafny.MultiSet(d_139_v57_))) in (d_138_v56_) else (d_141_v59_) + (d_141_v59_)))
                d_142_v60_: _dafny.Set
                d_142_v60_ = _dafny.Set({d_83_v10_})
                d_89_v14_ = (d_89_v14_).set((d_142_v60_).ispropersubset(d_142_v60_), (d_81_v9_).f2)
                d_143_v62_: _dafny.MultiSet
                def iife41_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in (d_79_v7_).Elements:
                        d_144_v61_: int = compr_29_
                        if (d_144_v61_) in (d_79_v7_):
                            coll29_ = coll29_.union(_dafny.Set([(d_144_v61_) - (185)]))
                    return _dafny.Set(coll29_)
                d_143_v62_ = _dafny.MultiSet([len(iife41_()
                ), 12])
                d_145_v63_: C7
                nw14_ = C7()
                nw14_.ctor__((((d_143_v62_).set((d_81_v9_).f2, default__.abs(d_72_v0_))).set(d_72_v0_, default__.abs(d_86_v13_))).set(d_72_v0_, default__.abs(d_72_v0_)), (default__.fm49(d_136_v54_, d_71_globalState_) if d_83_v10_ else d_80_v8_), d_79_v7_, d_80_v8_)
                d_145_v63_ = nw14_
                d_146_v64_: _dafny.Array
                nw15_ = _dafny.Array(_dafny.Map({}), 18)
                d_146_v64_ = nw15_
                d_147_v65_: C1
                nw16_ = C1()
                nw16_.ctor__(d_84_v11_, d_146_v64_, (d_79_v7_).set(default__.safeIndex((d_81_v9_).f2, len(d_79_v7_)), d_72_v0_), D0_DC1(d_72_v0_))
                d_147_v65_ = nw16_
                d_148_v66_: _dafny.Map
                d_148_v66_ = _dafny.Map({d_147_v65_: d_136_v54_})
                d_148_v66_ = (d_148_v66_).set(d_147_v65_, d_136_v54_)
                d_149_v67_: T0
                nw17_ = C8()
                nw17_.ctor__(d_147_v65_.f9, d_78_v6_, _dafny.SeqWithoutIsStrInference([(d_81_v9_).f2 for d_150_i2_ in range(default__.abs(169))]), D0_DC1(len(d_139_v57_)))
                d_149_v67_ = nw17_
                d_151_v68_: _dafny.Map
                d_151_v68_ = _dafny.Map({d_149_v67_: (d_81_v9_).f2})
                d_152_v69_: _dafny.Array
                nw18_ = _dafny.Array(False, 24)
                d_152_v69_ = nw18_
                index5_ = default__.safeIndex(414, (d_152_v69_).length(0))
                (d_152_v69_)[index5_] = d_83_v10_
                index6_ = default__.safeIndex(414, (d_152_v69_).length(0))
                rhs5_ = (d_151_v68_) | (d_151_v68_)
                rhs6_ = d_83_v10_
                lhs4_ = d_152_v69_
                lhs5_ = default__.safeIndex(414, (d_152_v69_).length(0))
                d_151_v68_ = rhs5_
                lhs4_[lhs5_] = rhs6_
            elif True:
                d_153_v70_: _dafny.Array
                nw19_ = _dafny.Array(D17.default()(), 1)
                d_153_v70_ = nw19_
                d_154_v71_: _dafny.MultiSet
                d_154_v71_ = _dafny.MultiSet([d_86_v13_, (d_81_v9_).f2])
                d_155_v72_: D17
                d_155_v72_ = D17_DC45(_dafny.MultiSet([d_154_v71_, d_154_v71_]))
                index7_ = default__.safeIndex(395, (d_153_v70_).length(0))
                (d_153_v70_)[index7_] = d_155_v72_
                index8_ = default__.safeIndex(395, (d_153_v70_).length(0))
                (d_153_v70_)[index8_] = d_155_v72_
                d_156_v73_: C4
                d_156_v73_ = d_73_v1_
                d_157_v74_: _dafny.Array
                nw20_ = _dafny.Array(None, 28)
                nw20_[int(0)] = d_73_v1_
                nw20_[int(1)] = d_73_v1_
                nw20_[int(2)] = d_73_v1_
                nw20_[int(3)] = d_73_v1_
                nw20_[int(4)] = d_73_v1_
                nw20_[int(5)] = d_73_v1_
                nw20_[int(6)] = d_73_v1_
                nw20_[int(7)] = d_73_v1_
                nw20_[int(8)] = d_73_v1_
                nw20_[int(9)] = d_73_v1_
                nw20_[int(10)] = d_73_v1_
                nw20_[int(11)] = d_73_v1_
                nw20_[int(12)] = d_73_v1_
                nw20_[int(13)] = d_73_v1_
                nw20_[int(14)] = d_73_v1_
                nw20_[int(15)] = d_73_v1_
                nw20_[int(16)] = d_73_v1_
                nw20_[int(17)] = d_73_v1_
                nw20_[int(18)] = d_73_v1_
                nw20_[int(19)] = d_73_v1_
                nw20_[int(20)] = d_73_v1_
                nw20_[int(21)] = (d_156_v73_)
                nw20_[int(22)] = d_73_v1_
                nw20_[int(23)] = d_73_v1_
                nw20_[int(24)] = d_73_v1_
                nw20_[int(25)] = d_73_v1_
                nw20_[int(26)] = d_73_v1_
                nw20_[int(27)] = d_73_v1_
                d_157_v74_ = nw20_
                d_158_v75_: _dafny.Seq
                d_158_v75_ = _dafny.SeqWithoutIsStrInference([d_157_v74_, d_157_v74_, d_157_v74_, d_157_v74_])
                rhs7_ = (d_81_v9_).f2
                rhs8_ = (d_81_v9_).fm7(d_86_v13_, d_71_globalState_)
                rhs9_ = ((_dafny.SeqWithoutIsStrInference([d_157_v74_, d_157_v74_])) + (_dafny.SeqWithoutIsStrInference([d_157_v74_])) if d_83_v10_ else d_158_v75_)
                d_84_v11_ = rhs7_
                d_74_v3_ = rhs8_
                d_158_v75_ = rhs9_
                d_159_v76_: _dafny.Array
                nw21_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_159_v76_ = nw21_
                d_160_v77_: D4
                d_160_v77_ = D4_DC10((0) - ((d_84_v11_) + ((d_81_v9_).f2)), d_86_v13_, len(d_74_v3_))
                d_161_v78_: _dafny.Seq
                d_161_v78_ = _dafny.SeqWithoutIsStrInference([d_159_v76_, d_159_v76_, d_159_v76_, d_159_v76_])
                d_162_v79_: _dafny.Map
                d_162_v79_ = _dafny.Map({d_83_v10_: (d_161_v78_)[default__.safeIndex(default__.fm44(d_83_v10_, d_71_globalState_), len(d_161_v78_))]})
                rhs10_ = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_163_i3_ in range(default__.abs(335))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_164_i4_ in range(default__.abs(235))])))
                rhs11_ = ((d_162_v79_)[True] if (True) in (d_162_v79_) else d_159_v76_)
                rhs12_ = d_160_v77_
                rhs13_ = ((d_81_v9_).f2 if d_83_v10_ else default__.fm32((0) - (d_72_v0_), d_71_globalState_))
                d_84_v11_ = rhs10_
                d_159_v76_ = rhs11_
                d_160_v77_ = rhs12_
                d_84_v11_ = rhs13_
                d_84_v11_ = ((d_72_v0_) + (-618)) - ((0) - (len(d_74_v3_)))
                d_84_v11_ = ((d_75_v4_)[d_74_v3_] if (d_74_v3_) in (d_75_v4_) else d_84_v11_)
            d_83_v10_ = d_83_v10_
            d_165_v80_: D7
            d_165_v80_ = D7_DC18()
            d_165_v80_ = d_165_v80_
            d_83_v10_ = (d_83_v10_) == ((d_84_v11_) == ((d_81_v9_).f2))
        elif source2_.is_DC50:
            d_166___mcc_h1_ = source2_.cf87
            d_167___mcc_h2_ = source2_.cf88
            d_168___mcc_h3_ = source2_.cf89
            d_169___mcc_h4_ = source2_.cf90
            d_170___mcc_h5_ = source2_.cf91
            d_171_cf91_ = d_170___mcc_h5_
            d_172_cf90_ = d_169___mcc_h4_
            d_173_cf89_ = d_168___mcc_h3_
            d_174_cf88_ = d_167___mcc_h2_
            d_175_cf87_ = d_166___mcc_h1_
            d_176_v81_: _dafny.Array
            nw22_ = _dafny.Array(_dafny.Map({}), 7)
            d_176_v81_ = nw22_
            d_177_v82_: _dafny.Seq
            d_177_v82_ = _dafny.SeqWithoutIsStrInference([d_79_v7_])
            d_178_v83_: T1
            nw23_ = C1()
            nw23_.ctor__((-276) + ((d_81_v9_).f2), (d_176_v81_ if d_173_cf89_ else d_176_v81_), _dafny.SeqWithoutIsStrInference([d_72_v0_, (_dafny.MultiSet([len(_dafny.Set({(d_177_v82_)[default__.safeIndex((d_81_v9_).f2, len(d_177_v82_))]}))])).cardinality]), d_80_v8_)
            d_178_v83_ = nw23_
            d_178_v83_ = d_178_v83_
            d_173_cf89_ = not(not(d_83_v10_))
            d_179_v84_: bool
            d_180_v85_: bool
            out10_: bool
            out11_: bool
            out10_, out11_ = (d_81_v9_).m1(d_86_v13_, d_71_globalState_)
            d_179_v84_ = out10_
            d_180_v85_ = out11_
            d_84_v11_ = (d_81_v9_).f2
        elif True:
            d_181___mcc_h6_ = source2_.cf85
            d_182_cf85_ = d_181___mcc_h6_
            d_83_v10_ = d_83_v10_
            d_83_v10_ = not((780) != ((d_84_v11_ if d_83_v10_ else 770)))
            d_84_v11_ = (d_86_v13_) * (default__.safeDivisionInt(d_84_v11_, 707))
            d_183_v86_: _dafny.Set
            d_183_v86_ = _dafny.Set({d_72_v0_})
            d_184_v87_: _dafny.Set
            d_184_v87_ = _dafny.Set({d_183_v86_})
            d_185_v88_: _dafny.Map
            d_185_v88_ = _dafny.Map({d_83_v10_: (D23_DC57(d_184_v87_)).cf100})
            d_185_v88_ = (d_185_v88_).set(d_83_v10_, d_184_v87_)
        d_186_v89_: D20
        d_186_v89_ = D20_DC53(not(d_83_v10_))
        if (d_186_v89_).cf94:
            d_84_v11_ = (d_81_v9_).f2
            d_187_v90_: _dafny.Array
            def lambda1_(d_188_v9_):
                def lambda2_(d_189_i5_):
                    return default__.safeModuloInt(d_189_i5_, (d_188_v9_).f2)

                return lambda2_

            init1_ = lambda1_(d_81_v9_)
            nw24_ = _dafny.Array(None, 23)
            for i0_1_ in range(nw24_.length(0)):
                nw24_[i0_1_] = init1_(i0_1_)
            d_187_v90_ = nw24_
            index9_ = default__.safeIndex(898, (d_187_v90_).length(0))
            (d_187_v90_)[index9_] = len((d_74_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnkl"))))
            d_190_v91_: _dafny.Set
            d_190_v91_ = _dafny.Set({d_83_v10_})
            index10_ = default__.safeIndex(898, (d_187_v90_).length(0))
            (d_187_v90_)[index10_] = len((_dafny.Set({d_83_v10_})) - (d_190_v91_))
            d_191_v92_: T0
            nw25_ = C2()
            nw25_.ctor__()
            d_191_v92_ = nw25_
            d_192_v93_: _dafny.MultiSet
            d_192_v93_ = _dafny.MultiSet([394])
            d_193_v94_: T0
            d_193_v94_ = d_191_v92_
            d_194_v95_: _dafny.Map
            d_194_v95_ = _dafny.Map({not((d_192_v93_) == (d_192_v93_)): (d_193_v94_)})
            d_195_v96_: _dafny.Map
            d_195_v96_ = _dafny.Map({_dafny.CodePoint('m'): d_80_v8_})
            d_196_v97_: _dafny.Map
            d_196_v97_ = _dafny.Map({not(not(d_83_v10_)): (d_191_v92_).fm5(d_72_v0_, d_83_v10_, d_195_v96_, d_83_v10_, d_71_globalState_)})
            d_191_v92_ = ((d_194_v95_)[((d_196_v97_)[d_83_v10_] if (d_83_v10_) in (d_196_v97_) else d_83_v10_)] if (((d_196_v97_)[d_83_v10_] if (d_83_v10_) in (d_196_v97_) else d_83_v10_)) in (d_194_v95_) else d_191_v92_)
            d_197_v98_: _dafny.Array
            nw26_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_197_v98_ = nw26_
            d_198_v99_: _dafny.Array
            def lambda3_(d_199_v9_, d_200_v14_, d_201_v0_):
                def lambda4_(d_202_i6_):
                    return D16_DC42((d_199_v9_).f2, len(d_200_v14_), d_201_v0_)

                return lambda4_

            init2_ = lambda3_(d_81_v9_, d_89_v14_, d_72_v0_)
            nw27_ = _dafny.Array(None, 15)
            for i0_2_ in range(nw27_.length(0)):
                nw27_[i0_2_] = init2_(i0_2_)
            d_198_v99_ = nw27_
            index11_ = default__.safeIndex(548, (d_197_v98_).length(0))
            (d_197_v98_)[index11_] = d_198_v99_
            d_203_v100_: D24
            d_203_v100_ = D24_DC59(d_198_v99_)
            index12_ = default__.safeIndex(548, (d_197_v98_).length(0))
            (d_197_v98_)[index12_] = (d_203_v100_).cf101
            d_204_v101_: str
            d_204_v101_ = _dafny.CodePoint('a')
            d_204_v101_ = d_204_v101_
        elif True:
            d_205_v102_: _dafny.Array
            nw28_ = _dafny.Array(False, 19)
            d_205_v102_ = nw28_
            index13_ = default__.safeIndex(189, (d_205_v102_).length(0))
            (d_205_v102_)[index13_] = d_83_v10_
            d_206_v103_: D12
            d_206_v103_ = D12_DC28(d_83_v10_, len(d_74_v3_))
            index14_ = default__.safeIndex(189, (d_205_v102_).length(0))
            (d_205_v102_)[index14_] = (default__.safeModuloInt((d_206_v103_).cf47, 860)) == (765)
            d_83_v10_ = not (d_83_v10_) or (d_83_v10_)
            d_207_v105_: _dafny.Set
            def iife42_():
                coll30_ = _dafny.Map()
                compr_30_: _dafny.Map
                for compr_30_ in (_dafny.SeqWithoutIsStrInference([d_85_v12_ for d_208_i7_ in range(default__.abs(163))])).Elements:
                    d_209_v104_: _dafny.Map = compr_30_
                    if (d_209_v104_) in (_dafny.SeqWithoutIsStrInference([d_85_v12_ for d_208_i7_ in range(default__.abs(163))])):
                        coll30_[d_209_v104_] = d_86_v13_
                return _dafny.Map(coll30_)
            d_207_v105_ = _dafny.Set({(0) - (d_72_v0_), d_84_v11_, len(iife42_()
            ), d_72_v0_})
            d_210_v106_: _dafny.Seq
            d_210_v106_ = _dafny.SeqWithoutIsStrInference([(d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]])
            d_211_v107_: str
            d_211_v107_ = _dafny.CodePoint('m')
            d_212_v108_: _dafny.Map
            d_212_v108_ = _dafny.Map({(d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]: d_211_v107_})
            d_213_v109_: _dafny.MultiSet
            d_213_v109_ = _dafny.MultiSet([(d_81_v9_).f2, len(d_212_v108_), d_84_v11_])
            nw29_ = _dafny.Array(None, 29)
            nw29_[int(0)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(1)] = d_83_v10_
            nw29_[int(2)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(3)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(4)] = not((d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))])
            nw29_[int(5)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(6)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(7)] = d_83_v10_
            nw29_[int(8)] = True
            nw29_[int(9)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(10)] = ((d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))] if d_83_v10_ else False)
            nw29_[int(11)] = d_83_v10_
            nw29_[int(12)] = d_83_v10_
            nw29_[int(13)] = (d_83_v10_) == (d_83_v10_)
            nw29_[int(14)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(15)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(16)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(17)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(18)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(19)] = (d_86_v13_) >= (d_72_v0_)
            nw29_[int(20)] = (d_207_v105_).issubset(_dafny.Set({d_72_v0_}))
            nw29_[int(21)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(22)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(23)] = (d_74_v3_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_214_i8_ in range(default__.abs(-976))]))
            nw29_[int(24)] = (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]
            nw29_[int(25)] = not((d_210_v106_)[default__.safeIndex((d_213_v109_).cardinality, len(d_210_v106_))])
            nw29_[int(26)] = ((d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]) == (True)
            nw29_[int(27)] = d_83_v10_
            nw29_[int(28)] = ((d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))]) and (d_83_v10_)
            d_205_v102_ = nw29_
            d_215_v110_: _dafny.Map
            d_215_v110_ = _dafny.Map({d_83_v10_: ((d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))] if d_83_v10_ else (d_205_v102_)[default__.safeIndex(189, (d_205_v102_).length(0))])})
            d_215_v110_ = (d_215_v110_).set((d_72_v0_) == (d_86_v13_), d_83_v10_)
            d_216_v111_: _dafny.Map
            d_216_v111_ = _dafny.Map({d_86_v13_: (d_81_v9_).fm6(d_72_v0_, len(d_79_v7_), d_71_globalState_)})
            d_72_v0_ = ((d_216_v111_)[d_84_v11_] if (d_84_v11_) in (d_216_v111_) else d_72_v0_)
        d_217_v112_: _dafny.Map
        d_217_v112_ = _dafny.Map({d_72_v0_: (d_81_v9_).f2})
        d_72_v0_ = ((d_81_v9_).f2) * ((0) - (len(d_217_v112_)))
        d_218_i9_: int
        d_218_i9_ = 0
        with _dafny.label("0"):
            pat_let_tv3_ = d_83_v10_
            pat_let_tv4_ = d_83_v10_
            pat_let_tv5_ = d_83_v10_
            def lambda5_(source3_):
                if source3_.is_DC48:
                    return pat_let_tv3_
                elif source3_.is_DC49:
                    d_224___mcc_h7_ = source3_.cf86
                    d_225_cf86_ = d_224___mcc_h7_
                    return pat_let_tv4_
                elif source3_.is_DC50:
                    d_226___mcc_h8_ = source3_.cf87
                    d_227___mcc_h9_ = source3_.cf88
                    d_228___mcc_h10_ = source3_.cf89
                    d_229___mcc_h11_ = source3_.cf90
                    d_230___mcc_h12_ = source3_.cf91
                    d_231_cf91_ = d_230___mcc_h12_
                    d_232_cf90_ = d_229___mcc_h11_
                    d_233_cf89_ = d_228___mcc_h10_
                    d_234_cf88_ = d_227___mcc_h9_
                    d_235_cf87_ = d_226___mcc_h8_
                    return d_235_cf87_
                elif True:
                    d_236___mcc_h13_ = source3_.cf85
                    d_237_cf85_ = d_236___mcc_h13_
                    return pat_let_tv5_

            while not(lambda5_(d_91_v16_)):
                with _dafny.c_label("0"):
                    if (d_218_i9_) >= (100):
                        raise _dafny.Break("0")
                    d_218_i9_ = (d_218_i9_) + (1)
                    d_219_v113_: _dafny.Map
                    d_219_v113_ = _dafny.Map({d_83_v10_: d_83_v10_})
                    rhs14_ = default__.fm47(((d_81_v9_).f2) * (d_72_v0_), ((d_81_v9_).f2) <= (d_72_v0_), d_71_globalState_)
                    rhs15_ = (d_84_v11_) > (d_72_v0_)
                    d_219_v113_ = rhs14_
                    d_83_v10_ = rhs15_
                    d_220_v114_: _dafny.Array
                    nw30_ = _dafny.Array(False, 5)
                    d_220_v114_ = nw30_
                    index15_ = default__.safeIndex(792, (d_220_v114_).length(0))
                    (d_220_v114_)[index15_] = d_83_v10_
                    d_221_v115_: T0
                    nw31_ = C2()
                    nw31_.ctor__()
                    d_221_v115_ = nw31_
                    index16_ = default__.safeIndex(792, (d_220_v114_).length(0))
                    rhs16_ = d_73_v1_
                    rhs17_ = not(not(d_83_v10_))
                    rhs18_ = d_220_v114_
                    rhs19_ = (d_86_v13_) + (len(_dafny.SeqWithoutIsStrInference([d_221_v115_])))
                    lhs6_ = d_220_v114_
                    lhs7_ = default__.safeIndex(792, (d_220_v114_).length(0))
                    d_73_v1_ = rhs16_
                    lhs6_[lhs7_] = rhs17_
                    d_220_v114_ = rhs18_
                    d_72_v0_ = rhs19_
                    d_222_v116_: C3
                    nw32_ = C3()
                    nw32_.ctor__()
                    d_222_v116_ = nw32_
                    d_223_v117_: _dafny.Seq
                    d_223_v117_ = _dafny.SeqWithoutIsStrInference([d_222_v116_])
                    d_222_v116_ = (d_222_v116_ if d_83_v10_ else (d_223_v117_)[default__.safeIndex((d_81_v9_).f2, len(d_223_v117_))])
                    d_83_v10_ = (d_220_v114_)[default__.safeIndex(792, (d_220_v114_).length(0))]
                    pass
            pass
        d_238_v118_: C4
        nw33_ = C4()
        nw33_.ctor__()
        d_238_v118_ = nw33_
        d_239_v119_: _dafny.Array
        nw34_ = _dafny.Array(int(0), 15)
        d_239_v119_ = nw34_
        d_240_v120_: _dafny.Map
        d_240_v120_ = _dafny.Map({d_239_v119_: d_239_v119_})
        d_241_v121_: _dafny.Map
        d_241_v121_ = _dafny.Map({d_83_v10_: _dafny.Map({d_239_v119_: d_239_v119_})})
        d_242_v122_: _dafny.Array
        nw35_ = _dafny.Array(None, 16)
        nw35_[int(0)] = (d_240_v120_ if d_83_v10_ else d_240_v120_)
        nw35_[int(1)] = (d_240_v120_) | (d_240_v120_)
        nw35_[int(2)] = (_dafny.Map({d_239_v119_: d_239_v119_})) | (_dafny.Map({d_239_v119_: d_239_v119_}))
        nw35_[int(3)] = (d_240_v120_) | (d_240_v120_)
        nw35_[int(4)] = (d_240_v120_) | (_dafny.Map({d_239_v119_: d_239_v119_}))
        nw35_[int(5)] = (d_240_v120_).set(d_239_v119_, d_239_v119_)
        nw35_[int(6)] = (_dafny.Map({d_239_v119_: d_239_v119_})) | (d_240_v120_)
        nw35_[int(7)] = ((d_241_v121_)[d_83_v10_] if (d_83_v10_) in (d_241_v121_) else d_240_v120_)
        nw35_[int(8)] = d_240_v120_
        nw35_[int(9)] = d_240_v120_
        nw35_[int(10)] = (d_240_v120_).set(d_239_v119_, d_239_v119_)
        nw35_[int(11)] = d_240_v120_
        nw35_[int(12)] = d_240_v120_
        nw35_[int(13)] = d_240_v120_
        nw35_[int(14)] = d_240_v120_
        nw35_[int(15)] = d_240_v120_
        d_242_v122_ = nw35_
        index17_ = default__.safeIndex(343, (d_242_v122_).length(0))
        (d_242_v122_)[index17_] = d_240_v120_
        index18_ = default__.safeIndex(343, (d_242_v122_).length(0))
        (d_242_v122_)[index18_] = d_240_v120_
        d_83_v10_ = False
        d_243_v123_: _dafny.Set
        d_243_v123_ = _dafny.Set({False, d_83_v10_, True, d_83_v10_, True})
        d_244_v124_: _dafny.Map
        d_244_v124_ = _dafny.Map({d_83_v10_: d_243_v123_})
        d_245_v125_: _dafny.Map
        d_245_v125_ = _dafny.Map({d_72_v0_: d_244_v124_})
        d_246_v126_: _dafny.Seq
        d_246_v126_ = _dafny.SeqWithoutIsStrInference([False, d_83_v10_, False])
        d_247_v127_: _dafny.Seq
        d_247_v127_ = _dafny.SeqWithoutIsStrInference([d_244_v124_, d_244_v124_, d_244_v124_])
        d_248_v128_: _dafny.Array
        nw36_ = _dafny.Array(None, 28)
        nw36_[int(0)] = ((d_245_v125_)[len(d_246_v126_)] if (len(d_246_v126_)) in (d_245_v125_) else d_244_v124_)
        nw36_[int(1)] = d_244_v124_
        nw36_[int(2)] = d_244_v124_
        nw36_[int(3)] = d_244_v124_
        nw36_[int(4)] = d_244_v124_
        nw36_[int(5)] = d_244_v124_
        nw36_[int(6)] = (d_244_v124_) | (_dafny.Map({False: d_243_v123_}))
        nw36_[int(7)] = (d_247_v127_)[default__.safeIndex(d_72_v0_, len(d_247_v127_))]
        nw36_[int(8)] = d_244_v124_
        nw36_[int(9)] = d_244_v124_
        nw36_[int(10)] = (_dafny.Map({True: d_243_v123_})) | (d_244_v124_)
        nw36_[int(11)] = d_244_v124_
        nw36_[int(12)] = d_244_v124_
        nw36_[int(13)] = d_244_v124_
        nw36_[int(14)] = d_244_v124_
        nw36_[int(15)] = d_244_v124_
        nw36_[int(16)] = d_244_v124_
        nw36_[int(17)] = d_244_v124_
        nw36_[int(18)] = (d_244_v124_ if d_83_v10_ else d_244_v124_)
        nw36_[int(19)] = (d_244_v124_) | (d_244_v124_)
        nw36_[int(20)] = d_244_v124_
        nw36_[int(21)] = ((d_245_v125_)[(d_81_v9_).f2] if ((d_81_v9_).f2) in (d_245_v125_) else d_244_v124_)
        nw36_[int(22)] = (d_244_v124_).set(True, d_243_v123_)
        nw36_[int(23)] = d_244_v124_
        nw36_[int(24)] = d_244_v124_
        nw36_[int(25)] = d_244_v124_
        nw36_[int(26)] = d_244_v124_
        nw36_[int(27)] = (_dafny.Map({False: d_243_v123_})) | (d_244_v124_)
        d_248_v128_ = nw36_
        index19_ = default__.safeIndex(312, (d_248_v128_).length(0))
        (d_248_v128_)[index19_] = d_244_v124_
        index20_ = default__.safeIndex(312, (d_248_v128_).length(0))
        (d_248_v128_)[index20_] = d_244_v124_
        d_83_v10_ = d_83_v10_
        d_249_v129_: str
        d_249_v129_ = _dafny.CodePoint('a')
        d_250_v130_: _dafny.Array
        nw37_ = _dafny.Array(None, 3)
        nw37_[int(0)] = d_249_v129_
        nw37_[int(1)] = d_249_v129_
        nw37_[int(2)] = d_249_v129_
        d_250_v130_ = nw37_
        d_251_v131_: _dafny.Seq
        d_251_v131_ = _dafny.SeqWithoutIsStrInference([d_250_v130_])
        d_252_v132_: _dafny.Map
        d_252_v132_ = _dafny.Map({(d_186_v89_).cf94: (d_251_v131_)[default__.safeIndex(d_84_v11_, len(d_251_v131_))]})
        d_252_v132_ = (d_252_v132_).set(d_83_v10_, d_250_v130_)
        if (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdgmtsaje")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhgie"))])) == (d_75_v4_):
            d_253_v133_: C0
            nw38_ = C0()
            nw38_.ctor__(d_74_v3_)
            d_253_v133_ = nw38_
            d_254_v134_: _dafny.Map
            d_254_v134_ = _dafny.Map({d_253_v133_: d_83_v10_})
            d_83_v10_ = ((d_83_v10_ if (d_246_v126_)[default__.safeIndex(d_86_v13_, len(d_246_v126_))] else d_83_v10_) if ((d_254_v134_)[d_253_v133_] if (d_253_v133_) in (d_254_v134_) else d_83_v10_) else d_83_v10_)
            d_255_v135_: D24
            d_255_v135_ = D24_DC62((d_81_v9_).f2)
            pat_let_tv6_ = d_84_v11_
            d_256_v136_: _dafny.Map
            def iife43_(_pat_let6_0):
                def iife44_(d_257_dt__update__tmp_h3_):
                    def iife45_(_pat_let7_0):
                        def iife46_(d_258_dt__update_hcf104_h0_):
                            return D24_DC62(d_258_dt__update_hcf104_h0_)
                        return iife46_(_pat_let7_0)
                    return iife45_(pat_let_tv6_)
                return iife44_(_pat_let6_0)
            d_256_v136_ = _dafny.Map({iife43_(d_255_v135_): d_239_v119_})
            d_259_v137_: _dafny.Array
            nw39_ = _dafny.Array(None, 27)
            nw39_[int(0)] = d_239_v119_
            nw39_[int(1)] = d_239_v119_
            nw39_[int(2)] = d_239_v119_
            nw39_[int(3)] = d_239_v119_
            nw39_[int(4)] = d_239_v119_
            nw39_[int(5)] = d_239_v119_
            nw39_[int(6)] = d_239_v119_
            nw39_[int(7)] = d_239_v119_
            nw39_[int(8)] = d_239_v119_
            nw39_[int(9)] = d_239_v119_
            nw39_[int(10)] = d_239_v119_
            nw39_[int(11)] = d_239_v119_
            nw39_[int(12)] = d_239_v119_
            nw39_[int(13)] = ((d_256_v136_)[d_255_v135_] if (d_255_v135_) in (d_256_v136_) else d_239_v119_)
            nw39_[int(14)] = d_239_v119_
            nw39_[int(15)] = d_239_v119_
            nw39_[int(16)] = d_239_v119_
            nw39_[int(17)] = (d_239_v119_ if True else d_239_v119_)
            nw39_[int(18)] = d_239_v119_
            nw39_[int(19)] = d_239_v119_
            nw39_[int(20)] = d_239_v119_
            nw39_[int(21)] = (d_239_v119_ if d_83_v10_ else d_239_v119_)
            nw39_[int(22)] = d_239_v119_
            nw39_[int(23)] = d_239_v119_
            nw39_[int(24)] = d_239_v119_
            nw39_[int(25)] = d_239_v119_
            nw39_[int(26)] = d_239_v119_
            d_259_v137_ = nw39_
            index21_ = default__.safeIndex(439, (d_259_v137_).length(0))
            (d_259_v137_)[index21_] = d_239_v119_
            index22_ = default__.safeIndex(439, (d_259_v137_).length(0))
            (d_259_v137_)[index22_] = d_239_v119_
            d_260_v138_: D15
            d_260_v138_ = D15_DC38(d_249_v129_)
            source4_ = d_260_v138_
            if source4_.is_DC38:
                d_261___mcc_h14_ = source4_.cf71
                d_262_cf71_ = d_261___mcc_h14_
                d_263_v140_: _dafny.Set
                d_263_v140_ = _dafny.Set({len(d_74_v3_), d_84_v11_})
                d_264_v141_: T0
                nw40_ = C8()
                def iife47_():
                    coll31_ = _dafny.Set()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(109, 874):
                        d_265_v139_: int = compr_31_
                        if ((109) <= (d_265_v139_)) and ((d_265_v139_) < (874)):
                            coll31_ = coll31_.union(_dafny.Set([(d_265_v139_) + ((d_81_v9_).f2)]))
                    return _dafny.Set(coll31_)
                nw40_.ctor__(len((iife47_()
                 if d_83_v10_ else d_263_v140_)), D0_DC0(d_76_v5_, (d_81_v9_).f2), (d_79_v7_) + (d_79_v7_), d_80_v8_)
                d_264_v141_ = nw40_
                nw41_ = C2()
                nw41_.ctor__()
                d_264_v141_ = nw41_
                d_266_v142_: _dafny.Map
                out12_: _dafny.Map
                out12_ = (d_81_v9_).m2(d_84_v11_, d_71_globalState_)
                d_266_v142_ = out12_
                d_267_v143_: D16
                d_267_v143_ = D16_DC42((d_81_v9_).f2, (d_81_v9_).f2, d_86_v13_)
                d_268_v145_: _dafny.MultiSet
                d_268_v145_ = _dafny.MultiSet([d_83_v10_])
                d_269_v146_: _dafny.Seq
                d_269_v146_ = _dafny.SeqWithoutIsStrInference([d_268_v145_, d_268_v145_, d_268_v145_])
                d_270_v147_: _dafny.Seq
                def iife48_():
                    coll32_ = _dafny.Map()
                    compr_32_: _dafny.MultiSet
                    for compr_32_ in (_dafny.MultiSet(d_269_v146_)).Elements:
                        d_271_v144_: _dafny.MultiSet = compr_32_
                        if (d_271_v144_) in (_dafny.MultiSet(d_269_v146_)):
                            coll32_[d_271_v144_] = d_262_cf71_
                    return _dafny.Map(coll32_)
                d_270_v147_ = _dafny.SeqWithoutIsStrInference([d_253_v133_.f8, d_253_v133_.f8, (d_253_v133_).fm15(_dafny.SeqWithoutIsStrInference([d_86_v13_, len(iife48_()
                ), (d_81_v9_).f2, d_84_v11_]), d_83_v10_, d_71_globalState_), d_253_v133_.f8, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))])
                d_272_v148_: _dafny.Array
                nw42_ = _dafny.Array(None, 7)
                nw42_[int(0)] = (_dafny.SeqWithoutIsStrInference([(D15_DC39(d_74_v3_, D9_DC21(_dafny.MultiSet([False, d_83_v10_])))).cf72 for d_273_i10_ in range(default__.abs(224))])) + ((_dafny.SeqWithoutIsStrInference([d_253_v133_.f8 for d_274_i11_ in range(default__.abs(-827))])).set(default__.safeIndex((d_267_v143_).cf77, len(_dafny.SeqWithoutIsStrInference([d_253_v133_.f8 for d_275_i11_ in range(default__.abs(-827))]))), d_74_v3_))
                nw42_[int(1)] = d_270_v147_
                nw42_[int(2)] = d_270_v147_
                nw42_[int(3)] = default__.fm72(d_71_globalState_)
                nw42_[int(4)] = d_270_v147_
                nw42_[int(5)] = _dafny.SeqWithoutIsStrInference([d_74_v3_, d_253_v133_.f8, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ea")), d_253_v133_.f8, d_253_v133_.f8])
                nw42_[int(6)] = (d_270_v147_).set(default__.safeIndex(549, len(d_270_v147_)), d_253_v133_.f8)
                d_272_v148_ = nw42_
                index23_ = default__.safeIndex(350, (d_272_v148_).length(0))
                (d_272_v148_)[index23_] = (d_270_v147_).set(default__.safeIndex(d_86_v13_, len(d_270_v147_)), d_74_v3_)
                index24_ = default__.safeIndex(350, (d_272_v148_).length(0))
                (d_272_v148_)[index24_] = (d_270_v147_) + ((d_270_v147_) + (d_270_v147_))
                d_262_cf71_ = d_262_cf71_
            elif source4_.is_DC39:
                d_276___mcc_h15_ = source4_.cf72
                d_277___mcc_h16_ = source4_.cf73
                d_278_cf73_ = d_277___mcc_h16_
                d_279_cf72_ = d_276___mcc_h15_
                d_280_v149_: _dafny.Map
                d_280_v149_ = _dafny.Map({d_75_v4_: (len(d_217_v112_)) >= ((d_81_v9_).f2)})
                d_280_v149_ = (d_280_v149_).set(_dafny.MultiSet([d_253_v133_.f8, d_253_v133_.f8, d_253_v133_.f8]), d_83_v10_)
                d_281_v150_: T0
                nw43_ = C8()
                nw43_.ctor__((d_81_v9_).f2, d_78_v6_, d_79_v7_, d_80_v8_)
                d_281_v150_ = nw43_
                d_282_v151_: _dafny.MultiSet
                d_282_v151_ = _dafny.MultiSet([d_281_v150_])
                d_283_v152_: _dafny.Map
                d_283_v152_ = _dafny.Map({default__.fm32((0) - (((d_282_v151_)[d_281_v150_] if (d_281_v150_) in (d_282_v151_) else (d_81_v9_).f2)), d_71_globalState_): d_253_v133_.f8})
                d_283_v152_ = (d_283_v152_).set(d_72_v0_, (d_253_v133_.f8) + (_dafny.SeqWithoutIsStrInference([d_249_v129_ for d_284_i12_ in range(default__.abs(915))])))
                d_285_v153_: C9
                nw44_ = C9()
                nw44_.ctor__()
                d_285_v153_ = nw44_
                d_286_v154_: _dafny.Array
                d_286_v154_ = (d_259_v137_)[default__.safeIndex(439, (d_259_v137_).length(0))]
                d_287_v155_: _dafny.MultiSet
                d_287_v155_ = _dafny.MultiSet([d_239_v119_, d_286_v154_, d_286_v154_])
                d_287_v155_ = d_287_v155_
            elif source4_.is_DC37:
                d_288___mcc_h17_ = source4_.cf70
                d_289_cf70_ = d_288___mcc_h17_
                d_290_v156_: _dafny.Array
                d_291_v157_: _dafny.Array
                d_292_v158_: bool
                out13_: _dafny.Array
                out14_: _dafny.Array
                out15_: bool
                out13_, out14_, out15_ = (d_73_v1_).m9(d_71_globalState_)
                d_290_v156_ = out13_
                d_291_v157_ = out14_
                d_292_v158_ = out15_
                d_84_v11_ = ((d_81_v9_).f2) * (default__.fm20(d_292_v158_, d_71_globalState_))
                d_293_v159_: _dafny.Set
                d_293_v159_ = _dafny.Set({(d_81_v9_).f2})
                d_294_v160_: _dafny.MultiSet
                d_294_v160_ = _dafny.MultiSet([d_292_v158_])
                d_295_v161_: _dafny.Array
                nw45_ = _dafny.Array(None, 22)
                nw45_[int(0)] = d_292_v158_
                nw45_[int(1)] = (d_253_v133_.f8) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtsgpyod")))
                nw45_[int(2)] = d_292_v158_
                nw45_[int(3)] = (d_293_v159_) == (d_293_v159_)
                nw45_[int(4)] = d_83_v10_
                nw45_[int(5)] = d_292_v158_
                nw45_[int(6)] = not((not(d_292_v158_)) not in (d_294_v160_))
                nw45_[int(7)] = d_83_v10_
                nw45_[int(8)] = d_292_v158_
                nw45_[int(9)] = d_292_v158_
                nw45_[int(10)] = d_83_v10_
                nw45_[int(11)] = d_83_v10_
                nw45_[int(12)] = not (d_292_v158_) or (d_292_v158_)
                nw45_[int(13)] = d_83_v10_
                nw45_[int(14)] = d_83_v10_
                nw45_[int(15)] = (d_246_v126_)[default__.safeIndex((_dafny.MultiSet([d_84_v11_])).cardinality, len(d_246_v126_))]
                nw45_[int(16)] = not(d_292_v158_)
                nw45_[int(17)] = d_292_v158_
                nw45_[int(18)] = d_83_v10_
                nw45_[int(19)] = (d_249_v129_) in (d_253_v133_.f8)
                nw45_[int(20)] = d_292_v158_
                nw45_[int(21)] = not (d_83_v10_) or (d_83_v10_)
                d_295_v161_ = nw45_
                index25_ = default__.safeIndex(491, (d_295_v161_).length(0))
                (d_295_v161_)[index25_] = not (d_83_v10_) or (d_83_v10_)
                index26_ = default__.safeIndex(491, (d_295_v161_).length(0))
                rhs20_ = (d_86_v13_) <= ((d_84_v11_) + ((d_81_v9_).f2))
                rhs21_ = (d_81_v9_).f2
                rhs22_ = not(d_83_v10_)
                lhs8_ = d_295_v161_
                lhs9_ = default__.safeIndex(491, (d_295_v161_).length(0))
                lhs8_[lhs9_] = rhs20_
                d_84_v11_ = rhs21_
                d_292_v158_ = rhs22_
                d_296_v162_: _dafny.MultiSet
                d_296_v162_ = _dafny.MultiSet([d_86_v13_])
                index27_ = default__.safeIndex(491, (d_295_v161_).length(0))
                index28_ = default__.safeIndex(491, (d_295_v161_).length(0))
                index29_ = default__.safeIndex(491, (d_295_v161_).length(0))
                rhs23_ = d_246_v126_
                rhs24_ = True
                rhs25_ = (d_296_v162_).issubset((d_296_v162_) | (default__.fm24(d_83_v10_, d_71_globalState_)))
                rhs26_ = (d_295_v161_)[default__.safeIndex(491, (d_295_v161_).length(0))]
                lhs10_ = d_295_v161_
                lhs11_ = default__.safeIndex(491, (d_295_v161_).length(0))
                lhs12_ = d_295_v161_
                lhs13_ = default__.safeIndex(491, (d_295_v161_).length(0))
                lhs14_ = d_295_v161_
                lhs15_ = default__.safeIndex(491, (d_295_v161_).length(0))
                d_246_v126_ = rhs23_
                lhs10_[lhs11_] = rhs24_
                lhs12_[lhs13_] = rhs25_
                lhs14_[lhs15_] = rhs26_
            elif True:
                d_297___mcc_h18_ = source4_.cf74
                d_298_cf74_ = d_297___mcc_h18_
                d_86_v13_ = default__.safeModuloInt(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_238_v118_])), (d_81_v9_).f2), (d_81_v9_).f2)
                d_299_v163_: C4
                nw46_ = C4()
                nw46_.ctor__()
                d_299_v163_ = nw46_
                d_300_v164_: _dafny.Map
                out16_: _dafny.Map
                out16_ = (d_238_v118_).m2((d_81_v9_).f2, d_71_globalState_)
                d_300_v164_ = out16_
                d_301_v165_: _dafny.Map
                d_301_v165_ = _dafny.Map({(d_81_v9_).f2: d_83_v10_})
                pat_let_tv7_ = d_81_v9_
                d_302_v166_: _dafny.Map
                def iife49_(_pat_let8_0):
                    def iife50_(d_303_dt__update__tmp_h4_):
                        def iife51_(_pat_let9_0):
                            def iife52_(d_304_dt__update_hcf2_h1_):
                                return D0_DC1(d_304_dt__update_hcf2_h1_)
                            return iife52_(_pat_let9_0)
                        return iife51_((pat_let_tv7_).f2)
                    return iife50_(_pat_let8_0)
                d_302_v166_ = _dafny.Map({d_249_v129_: iife49_(d_80_v8_)})
                d_83_v10_ = (d_81_v9_).fm5(default__.safeDivisionInt(200, d_84_v11_), (993) < (len(d_301_v165_)), d_302_v166_, d_83_v10_, d_71_globalState_)
            d_83_v10_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdjic"))) == (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) + (d_253_v133_.f8)).set(default__.safeIndex(len(d_246_v126_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) + (d_253_v133_.f8))), d_249_v129_))
            d_305_v167_: _dafny.Map
            d_305_v167_ = _dafny.Map({default__.safeModuloInt((d_81_v9_).f2, (d_81_v9_).f2): d_239_v119_})
            d_305_v167_ = d_305_v167_
        elif True:
            d_306_v168_: _dafny.Array
            nw47_ = _dafny.Array(None, 14)
            nw47_[int(0)] = d_239_v119_
            nw47_[int(1)] = d_239_v119_
            nw47_[int(2)] = d_239_v119_
            nw47_[int(3)] = d_239_v119_
            nw47_[int(4)] = d_239_v119_
            nw47_[int(5)] = d_239_v119_
            nw47_[int(6)] = d_239_v119_
            nw47_[int(7)] = d_239_v119_
            nw47_[int(8)] = d_239_v119_
            nw47_[int(9)] = d_239_v119_
            nw47_[int(10)] = d_239_v119_
            nw47_[int(11)] = d_239_v119_
            nw47_[int(12)] = d_239_v119_
            nw47_[int(13)] = d_239_v119_
            d_306_v168_ = nw47_
            d_307_v169_: _dafny.MultiSet
            d_307_v169_ = _dafny.MultiSet([d_83_v10_, d_83_v10_, d_83_v10_])
            d_308_v170_: _dafny.Map
            d_308_v170_ = _dafny.Map({((d_307_v169_)[d_83_v10_] if (d_83_v10_) in (d_307_v169_) else 202): _dafny.MultiSet(d_246_v126_)})
            d_309_v171_: D25
            d_309_v171_ = D25_DC63(d_308_v170_)
            d_310_v172_: _dafny.Map
            d_310_v172_ = _dafny.Map({d_306_v168_: ((d_309_v171_).cf105).set(879, ((d_308_v170_)[d_72_v0_] if (d_72_v0_) in (d_308_v170_) else d_307_v169_))})
            d_310_v172_ = (d_310_v172_).set(d_306_v168_, d_308_v170_)
            d_311_v173_: C4
            nw48_ = C4()
            nw48_.ctor__()
            d_311_v173_ = nw48_
            d_312_v174_: D24
            d_312_v174_ = D24_DC61(d_84_v11_, _dafny.MultiSet([d_239_v119_]))
            d_312_v174_ = d_312_v174_
            d_313_v175_: C5
            nw49_ = C5()
            nw49_.ctor__(d_239_v119_, _dafny.SeqWithoutIsStrInference([d_86_v13_]), d_80_v8_)
            d_313_v175_ = nw49_
            d_83_v10_ = ((d_74_v3_)[default__.safeIndex(len(d_74_v3_), len(d_74_v3_))]) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xg")))
        d_314_i13_: int
        d_314_i13_ = 0
        with _dafny.label("1"):
            while (d_83_v10_ if d_83_v10_ else d_83_v10_):
                with _dafny.c_label("1"):
                    if (d_314_i13_) >= (100):
                        raise _dafny.Break("1")
                    d_314_i13_ = (d_314_i13_) + (1)
                    d_315_v176_: D16
                    d_315_v176_ = D16_DC44(d_84_v11_)
                    d_316_v177_: _dafny.Set
                    d_316_v177_ = _dafny.Set({D16_DC44((d_81_v9_).f2), d_315_v176_, d_315_v176_, d_315_v176_})
                    d_317_v178_: D11
                    d_317_v178_ = D11_DC25(d_249_v129_, (d_81_v9_).f2, d_83_v10_)
                    rhs27_ = ((d_316_v177_).intersection(d_316_v177_)).isdisjoint(d_316_v177_)
                    rhs28_ = (d_317_v178_).cf43
                    rhs29_ = 559
                    rhs30_ = d_83_v10_
                    d_83_v10_ = rhs27_
                    d_83_v10_ = rhs28_
                    d_72_v0_ = rhs29_
                    d_83_v10_ = rhs30_
                    d_318_v179_: _dafny.Map
                    d_318_v179_ = _dafny.Map({d_246_v126_: (d_81_v9_).f2})
                    d_318_v179_ = (d_318_v179_).set(d_246_v126_, d_86_v13_)
                    d_84_v11_ = (d_84_v11_ if d_83_v10_ else (0) - ((d_81_v9_).f2))
                    d_319_v180_: bool
                    d_320_v181_: bool
                    out17_: bool
                    out18_: bool
                    out17_, out18_ = (d_238_v118_).m1((d_72_v0_) + (d_84_v11_), d_71_globalState_)
                    d_319_v180_ = out17_
                    d_320_v181_ = out18_
                    pass
            pass
        d_86_v13_ = 670
        d_74_v3_ = d_74_v3_
        _dafny.print(_dafny.string_of(d_72_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_74_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v4_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtn"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v5_)[0]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v5_)[1]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v5_)[2]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v5_)[3]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v5_)[4]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_78_v6_).cf0)[0]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_78_v6_).cf0)[1]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_78_v6_).cf0)[2]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_78_v6_).cf0)[3]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_78_v6_).cf0)[4]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v6_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_v7_) == (_dafny.SeqWithoutIsStrInference([936]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_80_v8_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v9_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_81_v9_).f3).cf0)[0]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_81_v9_).f3).cf0)[1]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_81_v9_).f3).cf0)[2]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_81_v9_).f3).cf0)[3]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_81_v9_).f3).cf0)[4]) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_81_v9_).f3).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_83_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_v12_) == (_dafny.Map({936: _dafny.CodePoint('i'), 3: _dafny.CodePoint('i')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_86_v13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v14_) == (_dafny.Map({False: 548}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_v15_) == (_dafny.Map({_dafny.Map({False: 548}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtn"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v16_).cf85) == (_dafny.Map({_dafny.Map({False: 548}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtn"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v89_).cf94))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v112_) == (_dafny.Map({936: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_240_v120_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_241_v121_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[0])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[1])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[2])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[3])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[4])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[5])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[6])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[7])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[8])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[9])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[10])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[11])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[12])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[13])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[14])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_242_v122_)[15])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v123_) == (_dafny.Set({False, True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v124_) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v125_) == (_dafny.Map({341: _dafny.Map({False: _dafny.Set({False, True})})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v126_) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v127_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: _dafny.Set({False, True})}), _dafny.Map({False: _dafny.Set({False, True})}), _dafny.Map({False: _dafny.Set({False, True})})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[0]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[1]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[2]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[3]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[4]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[5]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[6]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[7]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[8]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[9]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[10]) == (_dafny.Map({True: _dafny.Set({False, True}), False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[11]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[12]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[13]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[14]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[15]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[16]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[17]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[18]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[19]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[20]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[21]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[22]) == (_dafny.Map({False: _dafny.Set({False, True}), True: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[23]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[24]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[25]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[26]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_248_v128_)[27]) == (_dafny.Map({False: _dafny.Set({False, True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_249_v129_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v130_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v130_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v130_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_251_v131_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_252_v132_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_314_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(_dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC2(D1, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(int(0), False, _dafny.Array(None, 0))
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
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC7(D3, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC9(_dafny.MultiSet({}), _dafny.CodePoint('D'), int(0), _dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC9(D4, NamedTuple('DC9', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf22', Any), ('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC13(D5, NamedTuple('DC13', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC15(D6, NamedTuple('DC15', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({self.cf29.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC20(D8, NamedTuple('DC20', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC22(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC22(D9, NamedTuple('DC22', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC23(D10, NamedTuple('DC23', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25(_dafny.CodePoint('D'), int(0), False)
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

class D11_DC25(D11, NamedTuple('DC25', [('cf41', Any), ('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC28(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)

class D12_DC28(D12, NamedTuple('DC28', [('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC27(D12, NamedTuple('DC27', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC31(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)

class D13_DC31(D13, NamedTuple('DC31', [('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC32(D13, NamedTuple('DC32', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({self.cf57.VerbatimString(True)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC30(D13, NamedTuple('DC30', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC35(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D14_DC34)

class D14_DC35(D14, NamedTuple('DC35', [('cf63', Any), ('cf64', Any), ('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf63)}, {self.cf64.VerbatimString(True)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC34(D14, NamedTuple('DC34', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC34({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC34) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC38(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)

class D15_DC38(D15, NamedTuple('DC38', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC39(D15, NamedTuple('DC39', [('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({self.cf72.VerbatimString(True)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC42(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)

class D16_DC42(D16, NamedTuple('DC42', [('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({self.cf79.VerbatimString(True)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf82 == __o.cf82
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
        return lambda: D17_DC46(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)

class D17_DC46(D17, NamedTuple('DC46', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC45(D17, NamedTuple('DC45', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf83 == __o.cf83
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
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)
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

class D18_DC49(D18, NamedTuple('DC49', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {self.cf91.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC47(D18, NamedTuple('DC47', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)

class D19_DC51(D19, NamedTuple('DC51', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC53(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)

class D20_DC53(D20, NamedTuple('DC53', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf95', Any), ('cf96', Any), ('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({self.cf95.VerbatimString(True)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC52(D20, NamedTuple('DC52', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)

class D21_DC55(D21, NamedTuple('DC55', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)

class D22_DC56(D22, NamedTuple('DC56', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC58()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D23_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)

class D23_DC58(D23, NamedTuple('DC58', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC58'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC58)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC57(D23, NamedTuple('DC57', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC60()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D24_DC61)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D24_DC62)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D24_DC59)

class D24_DC60(D24, NamedTuple('DC60', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC61(D24, NamedTuple('DC61', [('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC61({_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC61) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC62(D24, NamedTuple('DC62', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC62({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC62) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC59(D24, NamedTuple('DC59', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC59({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC59) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC64(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D25_DC64)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D25_DC65)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D25_DC63)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D25_DC66)

class D25_DC64(D25, NamedTuple('DC64', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC64({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC64) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC65(D25, NamedTuple('DC65', [('cf107', Any), ('cf108', Any), ('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC65({_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC65) and self.cf107 == __o.cf107 and self.cf108 == __o.cf108 and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC63(D25, NamedTuple('DC63', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC63({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC63) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC66(D25, NamedTuple('DC66', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC66({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC66) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC68(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D26_DC68)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D26_DC69)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D26_DC67)

class D26_DC68(D26, NamedTuple('DC68', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC68({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC68) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC69(D26, NamedTuple('DC69', [('cf113', Any), ('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC69({self.cf113.VerbatimString(True)}, {_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC69) and self.cf113 == __o.cf113 and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC67(D26, NamedTuple('DC67', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC67({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC67) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D27_DC70)

class D27_DC70(D27, NamedTuple('DC70', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC70({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC70) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC72(int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D28_DC72)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D28_DC71)

class D28_DC72(D28, NamedTuple('DC72', [('cf117', Any), ('cf118', Any), ('cf119', Any), ('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC72({_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)}, {_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC72) and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119 and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC71(D28, NamedTuple('DC71', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC71({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC71) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, globalState):
        pass

    def m2(self, p0, globalState):
        pass


class T1:
    pass
    @property
    def f1(self):
        return self._f1
    @f1.setter
    def f1(self, value):
        self._f1 = value

class GlobalState:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self):
        pass
        pass


class C0:
    def  __init__(self):
        self.f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f8):
        (self).f8 = f8

    def fm14(self, globalState):
        return _dafny.CodePoint('w')

    def fm15(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_321_i0_ in range(default__.abs(5))])) + (self.f8)


class C1(T1):
    def  __init__(self):
        self._f1: D0 = D0.default()()
        self._f0: _dafny.Seq = _dafny.Seq({})
        self.f9: int = int(0)
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f1(self):
        return self._f1
    @f1.setter
    def f1(self, value):
        self._f1 = value
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f9, f10, f0, f1):
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f0 = f0
        (self).f1 = f1

    def fm6(self, p0, p1, globalState):
        def iife53_():
            coll33_ = _dafny.Set()
            compr_33_: int
            for compr_33_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drwqyvqsw")))])).Elements:
                d_322_v0_: int = compr_33_
                if (d_322_v0_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drwqyvqsw")))])):
                    coll33_ = coll33_.union(_dafny.Set([(d_322_v0_) + (613)]))
            return _dafny.Set(coll33_)
        return (_dafny.MultiSet([default__.safeDivisionInt(603, self.f9), self.f9, self.f9, len((iife53_()
        ) - (_dafny.Set({(0) - (self.f9), self.f9, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality}))), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_323_i0_ in range(default__.abs(797))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('c')])))])).cardinality

    def fm27(self, p0, p1, p2, p3, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([not(not(False))])) + (_dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))))

    def fm28(self, p0, p1, globalState):
        return self.f9

    def m13(self, p0, p1, globalState):
        r0: bool = False
        (self).f9 = self.f9
        r0 = p1
        if p1:
            d_324_v0_: _dafny.Seq
            d_324_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmxvlol"))
            d_325_v1_: C0
            nw50_ = C0()
            nw50_.ctor__(d_324_v0_)
            d_325_v1_ = nw50_
            d_326_v2_: C0
            nw51_ = C0()
            nw51_.ctor__((d_324_v0_) + (d_325_v1_.f8))
            d_326_v2_ = nw51_
            rhs31_ = True
            rhs32_ = p1
            rhs33_ = default__.safeDivisionInt((self.f9) - (self.f9), (self.f9) - (self.f9))
            lhs16_ = self
            r0 = rhs31_
            r0 = rhs32_
            lhs16_.f9 = rhs33_
            d_327_v3_: _dafny.MultiSet
            d_327_v3_ = _dafny.MultiSet([p1, p1])
            d_328_v4_: _dafny.Seq
            d_328_v4_ = _dafny.SeqWithoutIsStrInference([p1])
            d_329_v5_: D7
            d_329_v5_ = D7_DC17(d_328_v4_)
            d_330_v6_: _dafny.MultiSet
            d_330_v6_ = _dafny.MultiSet([(d_324_v0_) == (d_326_v2_.f8), p1, p1, not((_dafny.MultiSet((d_329_v5_).cf34)).ispropersubset(d_327_v3_))])
            d_327_v3_ = ((_dafny.MultiSet([p1])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1, p1])))) - (d_327_v3_)
            d_331_v7_: C0
            nw52_ = C0()
            nw52_.ctor__((d_324_v0_) + (d_325_v1_.f8))
            d_331_v7_ = nw52_
        elif True:
            d_332_v8_: _dafny.Array
            nw53_ = _dafny.Array(False, 29)
            d_332_v8_ = nw53_
            index30_ = default__.safeIndex(517, (d_332_v8_).length(0))
            (d_332_v8_)[index30_] = p1
            index31_ = default__.safeIndex(517, (d_332_v8_).length(0))
            (d_332_v8_)[index31_] = p1
            d_333_v9_: _dafny.Array
            nw54_ = _dafny.Array(int(0), 8)
            d_333_v9_ = nw54_
            index32_ = default__.safeIndex(609, (d_333_v9_).length(0))
            (d_333_v9_)[index32_] = self.f9
            index33_ = default__.safeIndex(609, (d_333_v9_).length(0))
            (d_333_v9_)[index33_] = default__.safeModuloInt(self.f9, self.f9)
            d_332_v8_ = d_332_v8_
            d_334_v10_: _dafny.Array
            def lambda6_(d_335_p0_, d_336_p1_):
                def lambda7_(d_337_i0_):
                    return (D11_DC25(d_335_p0_, (_dafny.MultiSet([D6_DC14(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqjlcv"))), D6_DC14(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcnl")))])).cardinality, d_336_p1_)).cf41

                return lambda7_

            init3_ = lambda6_(p0, p1)
            nw55_ = _dafny.Array(None, 2)
            for i0_3_ in range(nw55_.length(0)):
                nw55_[i0_3_] = init3_(i0_3_)
            d_334_v10_ = nw55_
            index34_ = default__.safeIndex(801, (d_334_v10_).length(0))
            (d_334_v10_)[index34_] = default__.fm29((d_332_v8_)[default__.safeIndex(517, (d_332_v8_).length(0))], globalState)
            index35_ = default__.safeIndex(801, (d_334_v10_).length(0))
            rhs34_ = p0
            rhs35_ = (d_332_v8_)[default__.safeIndex(517, (d_332_v8_).length(0))]
            lhs17_ = d_334_v10_
            lhs18_ = default__.safeIndex(801, (d_334_v10_).length(0))
            lhs17_[lhs18_] = rhs34_
            r0 = rhs35_
            index36_ = default__.safeIndex(609, (d_333_v9_).length(0))
            (d_333_v9_)[index36_] = (0) - (default__.safeDivisionInt((-65) + (self.f9), ((d_333_v9_)[default__.safeIndex(609, (d_333_v9_).length(0))]) * ((d_333_v9_)[default__.safeIndex(609, (d_333_v9_).length(0))])))
        d_338_v11_: _dafny.Array
        def lambda8_(d_339_p0_):
            def lambda9_(d_340_i2_):
                return d_339_p0_

            return lambda9_

        init4_ = lambda8_(p0)
        nw56_ = _dafny.Array(None, 23)
        for i0_4_ in range(nw56_.length(0)):
            nw56_[i0_4_] = init4_(i0_4_)
        d_338_v11_ = nw56_
        d_341_v12_: _dafny.Seq
        d_341_v12_ = _dafny.SeqWithoutIsStrInference([d_338_v11_, d_338_v11_, d_338_v11_])
        hi0_ = len(d_341_v12_)
        for d_342_i1_ in range(self.f9, hi0_):
            d_343_v13_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_343_v13_ = nw57_
            d_344_v14_: D2
            d_344_v14_ = D2_DC4(self.f9, p1, d_343_v13_)
            d_345_v15_: _dafny.MultiSet
            d_345_v15_ = _dafny.MultiSet([(self.f9) * ((d_344_v14_).cf5)])
            d_345_v15_ = d_345_v15_
            d_346_v16_: _dafny.Array
            def lambda10_(d_347_p1_):
                def lambda11_(d_348_i3_):
                    return (d_347_p1_) and (d_347_p1_)

                return lambda11_

            init5_ = lambda10_(p1)
            nw58_ = _dafny.Array(None, 14)
            for i0_5_ in range(nw58_.length(0)):
                nw58_[i0_5_] = init5_(i0_5_)
            d_346_v16_ = nw58_
            index37_ = default__.safeIndex(887, (d_346_v16_).length(0))
            (d_346_v16_)[index37_] = (p1) and (p1)
            index38_ = default__.safeIndex(887, (d_346_v16_).length(0))
            (d_346_v16_)[index38_] = p1
            d_349_v17_: _dafny.Array
            nw59_ = _dafny.Array(D6.default()(), 17)
            d_349_v17_ = nw59_
            d_350_v18_: _dafny.Seq
            d_350_v18_ = _dafny.SeqWithoutIsStrInference([p1])
            d_351_v19_: D6
            d_351_v19_ = D6_DC15(d_342_i1_, self.f9, default__.fm30(len(d_350_v18_), (d_346_v16_)[default__.safeIndex(887, (d_346_v16_).length(0))], (d_346_v16_)[default__.safeIndex(887, (d_346_v16_).length(0))], globalState))
            index39_ = default__.safeIndex(194, (d_349_v17_).length(0))
            (d_349_v17_)[index39_] = d_351_v19_
            d_352_v20_: _dafny.Set
            d_352_v20_ = _dafny.Set({self.f9})
            d_353_v21_: _dafny.Array
            nw60_ = _dafny.Array(None, 4)
            nw60_[int(0)] = _dafny.Set({self.f9})
            nw60_[int(1)] = (d_352_v20_) - (d_352_v20_)
            nw60_[int(2)] = d_352_v20_
            nw60_[int(3)] = d_352_v20_
            d_353_v21_ = nw60_
            d_354_v22_: _dafny.Array
            def lambda12_(d_355_i1_):
                def lambda13_(d_356_i4_):
                    return default__.safeDivisionInt(d_356_i4_, d_355_i1_)

                return lambda13_

            init6_ = lambda12_(d_342_i1_)
            nw61_ = _dafny.Array(None, 6)
            for i0_6_ in range(nw61_.length(0)):
                nw61_[i0_6_] = init6_(i0_6_)
            d_354_v22_ = nw61_
            index40_ = default__.safeIndex(256, (d_354_v22_).length(0))
            (d_354_v22_)[index40_] = default__.safeModuloInt(self.f9, self.f9)
            index41_ = default__.safeIndex(194, (d_349_v17_).length(0))
            index42_ = default__.safeIndex(887, (d_346_v16_).length(0))
            index43_ = default__.safeIndex(256, (d_354_v22_).length(0))
            rhs36_ = d_351_v19_
            rhs37_ = d_353_v21_
            rhs38_ = (d_346_v16_)[default__.safeIndex(887, (d_346_v16_).length(0))]
            rhs39_ = ((0) - (d_342_i1_) if p1 else (self.f9) - (self.f9))
            lhs19_ = d_349_v17_
            lhs20_ = default__.safeIndex(194, (d_349_v17_).length(0))
            lhs21_ = d_346_v16_
            lhs22_ = default__.safeIndex(887, (d_346_v16_).length(0))
            lhs23_ = d_354_v22_
            lhs24_ = default__.safeIndex(256, (d_354_v22_).length(0))
            lhs19_[lhs20_] = rhs36_
            d_353_v21_ = rhs37_
            lhs21_[lhs22_] = rhs38_
            lhs23_[lhs24_] = rhs39_
            d_357_v23_: _dafny.Seq
            d_357_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rarjd"))
            d_358_v24_: D6
            d_358_v24_ = D6_DC14((d_357_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmypydeh"))))
            pat_let_tv8_ = d_357_v23_
            index44_ = default__.safeIndex(256, (d_354_v22_).length(0))
            def iife54_(_pat_let10_0):
                def iife55_(d_359_dt__update__tmp_h0_):
                    def iife56_(_pat_let11_0):
                        def iife57_(d_360_dt__update_hcf29_h0_):
                            return D6_DC14(d_360_dt__update_hcf29_h0_)
                        return iife57_(_pat_let11_0)
                    return iife56_(pat_let_tv8_)
                return iife55_(_pat_let10_0)
            rhs40_ = iife54_(D6_DC14(d_357_v23_))
            rhs41_ = _dafny.SeqWithoutIsStrInference([(d_346_v16_)[default__.safeIndex(887, (d_346_v16_).length(0))], p1, p1])
            rhs42_ = self.f9
            lhs25_ = d_354_v22_
            lhs26_ = default__.safeIndex(256, (d_354_v22_).length(0))
            d_358_v24_ = rhs40_
            d_350_v18_ = rhs41_
            lhs25_[lhs26_] = rhs42_
        hi1_ = self.f9
        for d_361_i5_ in range(self.f9, hi1_):
            d_362_v25_: _dafny.Array
            nw62_ = _dafny.Array(_dafny.Seq({}), 3)
            d_362_v25_ = nw62_
            d_363_v26_: _dafny.Seq
            d_363_v26_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            index45_ = default__.safeIndex(908, (d_362_v25_).length(0))
            (d_362_v25_)[index45_] = (d_363_v26_) + ((d_363_v26_).set(default__.safeIndex(self.f9, len(d_363_v26_)), p1))
            d_364_v27_: D7
            d_364_v27_ = D7_DC17(d_363_v26_)
            d_365_v28_: _dafny.Map
            d_365_v28_ = _dafny.Map({p1: d_363_v26_})
            index46_ = default__.safeIndex(908, (d_362_v25_).length(0))
            (d_362_v25_)[index46_] = (((d_364_v27_ if p1 else D7_DC17(((d_365_v28_)[p1] if (p1) in (d_365_v28_) else d_363_v26_)))).cf34).set(default__.safeIndex(self.f9, len(((d_364_v27_ if p1 else D7_DC17(((d_365_v28_)[p1] if (p1) in (d_365_v28_) else d_363_v26_)))).cf34)), (D12_DC29(p1, d_361_i5_, p1, 897, p1)).cf50)
            d_366_v29_: _dafny.Seq
            d_366_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyg"))
            d_366_v29_ = (d_366_v29_) + (d_366_v29_)
            (self).f9 = default__.safeDivisionInt(self.f9, self.f9)
            d_367_v30_: _dafny.Map
            d_367_v30_ = _dafny.Map({len(d_366_v29_): p1})
            d_368_v31_: _dafny.Array
            nw63_ = _dafny.Array(False, 12)
            d_368_v31_ = nw63_
            d_369_v32_: _dafny.MultiSet
            d_369_v32_ = _dafny.MultiSet([d_368_v31_, d_368_v31_])
            d_370_v33_: _dafny.MultiSet
            d_370_v33_ = _dafny.MultiSet([p0])
            d_371_v34_: _dafny.MultiSet
            d_371_v34_ = _dafny.MultiSet([d_361_i5_, d_361_i5_, (d_370_v33_).cardinality])
            d_372_v35_: _dafny.Set
            d_372_v35_ = _dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_361_i5_])), d_371_v34_, _dafny.MultiSet((self).f0), _dafny.MultiSet([len(default__.fm31(self.f9, p1, self.f9, globalState))])})
            d_373_v36_: _dafny.Array
            nw64_ = _dafny.Array(None, 22)
            nw64_[int(0)] = ((0) - (d_361_i5_)) - (d_361_i5_)
            nw64_[int(1)] = len(d_367_v30_)
            nw64_[int(2)] = (self.f9) + (d_361_i5_)
            nw64_[int(3)] = d_361_i5_
            nw64_[int(4)] = default__.safeDivisionInt((0) - (d_361_i5_), d_361_i5_)
            nw64_[int(5)] = default__.safeDivisionInt(self.f9, d_361_i5_)
            nw64_[int(6)] = d_361_i5_
            nw64_[int(7)] = (self.f9) * (self.f9)
            nw64_[int(8)] = len((self).f0)
            nw64_[int(9)] = (0) - ((self).fm28(-80, p1, globalState))
            nw64_[int(10)] = ((d_369_v32_)[d_368_v31_] if (d_368_v31_) in (d_369_v32_) else len(d_372_v35_))
            nw64_[int(11)] = self.f9
            nw64_[int(12)] = d_361_i5_
            nw64_[int(13)] = d_361_i5_
            nw64_[int(14)] = self.f9
            nw64_[int(15)] = len((self).f0)
            nw64_[int(16)] = d_361_i5_
            nw64_[int(17)] = (d_361_i5_) * ((0) - (len(default__.fm31(self.f9, p1, (0) - (len(d_366_v29_)), globalState))))
            nw64_[int(18)] = (self.f9) + (345)
            nw64_[int(19)] = d_361_i5_
            nw64_[int(20)] = self.f9
            nw64_[int(21)] = self.f9
            d_373_v36_ = nw64_
            index47_ = default__.safeIndex(48, (d_373_v36_).length(0))
            (d_373_v36_)[index47_] = 882
            index48_ = default__.safeIndex(48, (d_373_v36_).length(0))
            (d_373_v36_)[index48_] = ((d_369_v32_)[d_368_v31_] if (d_368_v31_) in (d_369_v32_) else self.f9)
        d_374_v37_: D12
        d_374_v37_ = D12_DC29(p1, (self).fm28(self.f9, not(not(p1)), globalState), p1, self.f9, not(p1))
        d_375_v38_: _dafny.Map
        d_375_v38_ = _dafny.Map({(d_374_v37_).cf49: self.f9})
        d_376_v39_: _dafny.Map
        d_376_v39_ = _dafny.Map({self.f9: p1})
        d_377_v40_: D11
        d_377_v40_ = D11_DC25(p0, 390, ((d_376_v39_)[self.f9] if (self.f9) in (d_376_v39_) else p1))
        d_378_v41_: _dafny.Seq
        d_378_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhh"))
        hi2_ = len(d_378_v41_)
        for d_379_i6_ in range((0) - (((d_375_v38_)[(d_377_v40_).cf42] if ((d_377_v40_).cf42) in (d_375_v38_) else (_dafny.MultiSet([p0])).cardinality)), hi2_):
            r0 = (not(p1) if p1 else p1)
            d_380_v42_: _dafny.MultiSet
            d_380_v42_ = _dafny.MultiSet([d_379_i6_, d_379_i6_, (self).fm6(self.f9, self.f9, globalState), d_379_i6_, self.f9])
            (self).f9 = default__.safeModuloInt(((d_380_v42_).cardinality) + (d_379_i6_), ((self).fm6(self.f9, self.f9, globalState)) + (self.f9))
            d_381_v43_: str
            d_381_v43_ = _dafny.CodePoint('q')
            d_381_v43_ = d_381_v43_
            d_382_v44_: _dafny.Array
            def lambda14_(d_383_i7_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htrgdatxk"))

            init7_ = lambda14_
            nw65_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw65_.length(0)):
                nw65_[i0_7_] = init7_(i0_7_)
            d_382_v44_ = nw65_
            index49_ = default__.safeIndex(115, (d_382_v44_).length(0))
            (d_382_v44_)[index49_] = d_378_v41_
            index50_ = default__.safeIndex(115, (d_382_v44_).length(0))
            (d_382_v44_)[index50_] = (d_378_v41_ if p1 else d_378_v41_)
        r0 = ((p0) not in (default__.fm31(self.f9, p1, (_dafny.MultiSet([p0, p0, p0])).cardinality, globalState)) if (d_378_v41_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ly"))) else p1)
        return r0

    @property
    def f10(self):
        return self._f10

class C2(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, p2, p3, globalState):
        return False

    def fm41(self, p0, p1, globalState):
        return _dafny.MultiSet([False])

    def fm42(self, p0, globalState):
        def iife58_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in _dafny.IntegerRange(-558, 560):
                d_384_v0_: int = compr_34_
                if ((-558) <= (d_384_v0_)) and ((d_384_v0_) < (560)):
                    coll34_[default__.safeDivisionInt(d_384_v0_, len(_dafny.SeqWithoutIsStrInference([False])))] = _dafny.SeqWithoutIsStrInference([D2_DC5(-832, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): _dafny.CodePoint('p')})) for d_385_i0_ in range(default__.abs(-101))]), False), D2_DC5(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tam"))): _dafny.CodePoint('k')})), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, not(False), not(True)])), len(_dafny.Set({True}))]), True)])
            return _dafny.Map(coll34_)
        return (len(iife58_()
        )) < (281)

    def m1(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        d_386_v0_: _dafny.MultiSet
        d_386_v0_ = _dafny.MultiSet([p0, p0])
        d_387_v1_: _dafny.Map
        d_387_v1_ = _dafny.Map({(d_386_v0_).cardinality: p0})
        d_388_v2_: _dafny.Array
        nw66_ = _dafny.Array(_dafny.Map({}), 1)
        d_388_v2_ = nw66_
        d_389_v3_: _dafny.Seq
        d_389_v3_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpmro")))])
        d_390_v4_: T1
        nw67_ = C1()
        nw67_.ctor__(p0, d_388_v2_, d_389_v3_, D0_DC1(p0))
        d_390_v4_ = nw67_
        d_391_v5_: _dafny.Map
        d_391_v5_ = _dafny.Map({d_390_v4_: p0})
        d_392_v6_: C0
        nw68_ = C0()
        nw68_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "us")))
        d_392_v6_ = nw68_
        d_393_v7_: _dafny.Map
        d_393_v7_ = _dafny.Map({p0: d_392_v6_})
        d_394_v8_: _dafny.Seq
        d_394_v8_ = _dafny.SeqWithoutIsStrInference([d_387_v1_, _dafny.Map({p0: p0}), _dafny.Map({len(d_391_v5_): len(d_393_v7_)})])
        d_395_v9_: int
        d_395_v9_ = 156
        rhs43_ = (d_394_v8_) + (d_394_v8_)
        rhs44_ = len(_dafny.SeqWithoutIsStrInference([d_395_v9_ for d_396_i0_ in range(default__.abs(-425))]))
        rhs45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvuiwig"))
        lhs27_ = d_392_v6_
        d_394_v8_ = rhs43_
        d_395_v9_ = rhs44_
        lhs27_.f8 = rhs45_
        d_397_v10_: _dafny.Array
        def lambda15_(d_398_i2_):
            return (False) in (_dafny.Map({not(False): _dafny.MultiSet([False])}))

        init8_ = lambda15_
        nw69_ = _dafny.Array(None, 17)
        for i0_8_ in range(nw69_.length(0)):
            nw69_[i0_8_] = init8_(i0_8_)
        d_397_v10_ = nw69_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_397_v10_).length(0)):
            d_399_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_399_i1_)) and ((d_399_i1_) < ((d_397_v10_).length(0)))):
                (d_397_v10_)[(d_399_i1_)] = (d_395_v9_) == (((d_390_v4_).f0)[default__.safeIndex(d_395_v9_, len((d_390_v4_).f0))])
        d_400_v11_: bool
        d_400_v11_ = True
        d_401_v12_: _dafny.Map
        d_401_v12_ = _dafny.Map({d_400_v11_: d_395_v9_})
        d_402_v13_: _dafny.Seq
        d_402_v13_ = _dafny.SeqWithoutIsStrInference([d_392_v6_.f8, d_392_v6_.f8, d_392_v6_.f8])
        d_403_v14_: _dafny.Map
        d_403_v14_ = _dafny.Map({d_401_v12_: (d_402_v13_)[default__.safeIndex(359, len(d_402_v13_))]})
        d_404_v16_: _dafny.Seq
        d_404_v16_ = _dafny.SeqWithoutIsStrInference([d_401_v12_, d_401_v12_])
        d_405_v17_: int
        out19_: int
        def iife59_():
            coll35_ = _dafny.Map()
            compr_35_: _dafny.Map
            for compr_35_ in (d_404_v16_).Elements:
                d_406_v15_: _dafny.Map = compr_35_
                if (d_406_v15_) in (d_404_v16_):
                    coll35_[d_406_v15_] = d_392_v6_.f8
            return _dafny.Map(coll35_)
        out19_ = (self).m14(d_400_v11_, d_400_v11_, (d_403_v14_) | (iife59_()
        ), p0, globalState)
        d_405_v17_ = out19_
        r0 = d_400_v11_
        d_407_v18_: str
        d_407_v18_ = _dafny.CodePoint('x')
        d_408_i3_: int
        d_408_i3_ = 0
        with _dafny.label("2"):
            def iife61_():
                coll37_ = _dafny.Set()
                compr_37_: int
                for compr_37_ in (_dafny.Map({d_395_v9_: d_407_v18_})).keys.Elements:
                    d_418_v19_: int = compr_37_
                    if (d_418_v19_) in (_dafny.Map({d_395_v9_: d_407_v18_})):
                        coll37_ = coll37_.union(_dafny.Set([default__.safeModuloInt(d_418_v19_, (0) - (-404))]))
                return _dafny.Set(coll37_)
            while (iife61_()
            ).issubset(_dafny.Set({p0})):
                with _dafny.c_label("2"):
                    if (d_408_i3_) >= (100):
                        raise _dafny.Break("2")
                    d_408_i3_ = (d_408_i3_) + (1)
                    d_409_v20_: D6
                    d_409_v20_ = D6_DC15((d_390_v4_).fm6(d_405_v17_, d_405_v17_, globalState), p0, d_400_v11_)
                    d_410_v21_: D6
                    d_410_v21_ = D6_DC16(d_409_v20_)
                    d_411_v22_: D6
                    d_411_v22_ = D6_DC16(d_409_v20_)
                    d_411_v22_ = d_411_v22_
                    r1 = d_400_v11_
                    d_412_v24_: _dafny.Map
                    d_412_v24_ = _dafny.Map({d_392_v6_.f8: d_407_v18_})
                    d_413_v25_: D11
                    d_413_v25_ = D11_DC24(d_412_v24_)
                    d_414_v26_: D11
                    d_414_v26_ = D11_DC26(d_413_v25_)
                    d_415_v27_: _dafny.Seq
                    d_415_v27_ = _dafny.SeqWithoutIsStrInference([d_400_v11_, d_400_v11_])
                    d_416_v28_: D14
                    d_416_v28_ = D14_DC34(default__.fm43(d_414_v26_, d_400_v11_, len(d_415_v27_), globalState))
                    def iife60_():
                        coll36_ = _dafny.Map()
                        compr_36_: int
                        for compr_36_ in _dafny.IntegerRange(442, -928):
                            d_417_v23_: int = compr_36_
                            if ((442) <= (d_417_v23_)) and ((d_417_v23_) < (-928)):
                                coll36_[(d_417_v23_) + (d_405_v17_)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkbtfupqy"))
                        return _dafny.Map(coll36_)
                    d_400_v11_ = ((iife60_()
                    ) | (_dafny.Map({p0: d_392_v6_.f8}))) != ((d_416_v28_).cf62)
                    index51_ = default__.safeIndex(653, (d_397_v10_).length(0))
                    (d_397_v10_)[index51_] = d_400_v11_
                    index52_ = default__.safeIndex(653, (d_397_v10_).length(0))
                    (d_397_v10_)[index52_] = d_400_v11_
                    pass
            pass
        r1 = d_400_v11_
        r0 = d_400_v11_
        r1 = d_400_v11_
        return r0, r1

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_419_v0_: _dafny.Seq
        d_419_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgwejo"))
        d_419_v0_ = (d_419_v0_) + (d_419_v0_)
        d_420_i0_: int
        d_420_i0_ = 0
        with _dafny.label("3"):
            while True:
                with _dafny.c_label("3"):
                    if (d_420_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_420_i0_ = (d_420_i0_) + (1)
                    d_421_v1_: _dafny.Array
                    nw70_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
                    d_421_v1_ = nw70_
                    index53_ = default__.safeIndex(932, (d_421_v1_).length(0))
                    (d_421_v1_)[index53_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_422_i1_ in range(default__.abs(54))])
                    index54_ = default__.safeIndex(932, (d_421_v1_).length(0))
                    (d_421_v1_)[index54_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pekucbt"))
                    d_423_v2_: bool
                    d_423_v2_ = True
                    d_424_v3_: D0
                    d_424_v3_ = D0_DC1(p0)
                    d_425_v4_: _dafny.Seq
                    d_425_v4_ = _dafny.SeqWithoutIsStrInference([d_423_v2_])
                    d_426_v5_: _dafny.Map
                    d_426_v5_ = _dafny.Map({d_424_v3_: len(d_425_v4_)})
                    d_427_v6_: _dafny.Seq
                    d_427_v6_ = _dafny.SeqWithoutIsStrInference([d_426_v5_])
                    d_428_v7_: _dafny.Seq
                    d_428_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
                    d_429_v8_: _dafny.Map
                    d_429_v8_ = _dafny.Map({len(d_428_v7_): d_423_v2_})
                    d_430_v9_: _dafny.Map
                    d_430_v9_ = _dafny.Map({(0) - ((_dafny.MultiSet(d_427_v6_)).cardinality): (((d_429_v8_)[432] if (432) in (d_429_v8_) else d_423_v2_)) and (d_423_v2_)})
                    d_431_v10_: _dafny.Set
                    d_431_v10_ = _dafny.Set({p0})
                    d_423_v2_ = ((d_429_v8_)[len((_dafny.Map({len(d_431_v10_): False})) | (d_429_v8_))] if (len((_dafny.Map({len(d_431_v10_): False})) | (d_429_v8_))) in (d_429_v8_) else (d_423_v2_) == (d_423_v2_))
                    d_432_v11_: _dafny.Array
                    nw71_ = _dafny.Array(False, 2)
                    d_432_v11_ = nw71_
                    d_433_v12_: _dafny.Set
                    d_433_v12_ = _dafny.Set({d_432_v11_})
                    d_434_v13_: _dafny.Set
                    d_434_v13_ = d_433_v12_
                    d_435_v14_: _dafny.Map
                    d_435_v14_ = _dafny.Map({d_423_v2_: d_434_v13_})
                    d_435_v14_ = (d_435_v14_).set(d_423_v2_, d_433_v12_)
                    d_436_v15_: _dafny.Array
                    def lambda16_(d_437_p0_):
                        def lambda17_(d_438_i2_):
                            return (d_438_i2_) + (d_437_p0_)

                        return lambda17_

                    init9_ = lambda16_(p0)
                    nw72_ = _dafny.Array(None, 24)
                    for i0_9_ in range(nw72_.length(0)):
                        nw72_[i0_9_] = init9_(i0_9_)
                    d_436_v15_ = nw72_
                    index55_ = default__.safeIndex(249, (d_436_v15_).length(0))
                    (d_436_v15_)[index55_] = p0
                    d_439_v16_: _dafny.Array
                    nw73_ = _dafny.Array(_dafny.Map({}), 10)
                    d_439_v16_ = nw73_
                    index56_ = default__.safeIndex(948, (d_439_v16_).length(0))
                    (d_439_v16_)[index56_] = _dafny.Map({d_423_v2_: d_423_v2_})
                    d_440_v17_: _dafny.Map
                    d_440_v17_ = _dafny.Map({p0: (d_421_v1_)[default__.safeIndex(932, (d_421_v1_).length(0))]})
                    d_441_v18_: str
                    d_441_v18_ = _dafny.CodePoint('i')
                    d_442_v19_: _dafny.Map
                    d_442_v19_ = _dafny.Map({not((d_425_v4_)[default__.safeIndex(p0, len(d_425_v4_))]): d_423_v2_})
                    index57_ = default__.safeIndex(249, (d_436_v15_).length(0))
                    index58_ = default__.safeIndex(948, (d_439_v16_).length(0))
                    rhs46_ = default__.fm44(not((p0) < (len(_dafny.Set({d_423_v2_})))), globalState)
                    rhs47_ = (((((d_440_v17_)[p0] if (p0) in (d_440_v17_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_443_i3_ in range(default__.abs(120))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_444_i3_ in range(default__.abs(120))]))), d_441_v18_)) if d_423_v2_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhkvhgy")))).set(default__.safeIndex(p0, len((((d_440_v17_)[p0] if (p0) in (d_440_v17_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_445_i3_ in range(default__.abs(120))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_446_i3_ in range(default__.abs(120))]))), d_441_v18_)) if d_423_v2_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhkvhgy"))))), d_441_v18_)) + (d_419_v0_)
                    rhs48_ = d_423_v2_
                    rhs49_ = d_442_v19_
                    rhs50_ = d_428_v7_
                    lhs28_ = d_436_v15_
                    lhs29_ = default__.safeIndex(249, (d_436_v15_).length(0))
                    lhs30_ = d_439_v16_
                    lhs31_ = default__.safeIndex(948, (d_439_v16_).length(0))
                    lhs28_[lhs29_] = rhs46_
                    d_419_v0_ = rhs47_
                    d_423_v2_ = rhs48_
                    lhs30_[lhs31_] = rhs49_
                    d_428_v7_ = rhs50_
                    pass
            pass
        d_447_v20_: bool
        d_447_v20_ = True
        d_448_v21_: _dafny.Seq
        d_448_v21_ = _dafny.SeqWithoutIsStrInference([p0])
        source5_ = default__.fm45(d_447_v20_, d_447_v20_, default__.safeModuloInt(len(d_448_v21_), p0), globalState)
        if source5_.is_DC35:
            d_449___mcc_h0_ = source5_.cf63
            d_450___mcc_h1_ = source5_.cf64
            d_451___mcc_h2_ = source5_.cf65
            d_452___mcc_h3_ = source5_.cf66
            d_453___mcc_h4_ = source5_.cf67
            d_454_cf67_ = d_453___mcc_h4_
            d_455_cf66_ = d_452___mcc_h3_
            d_456_cf65_ = d_451___mcc_h2_
            d_457_cf64_ = d_450___mcc_h1_
            d_458_cf63_ = d_449___mcc_h0_
            d_459_v22_: _dafny.Map
            d_459_v22_ = _dafny.Map({d_458_cf63_: d_456_cf65_})
            d_460_v23_: D12
            d_460_v23_ = D12_DC29(d_447_v20_, ((d_459_v22_)[p0] if (p0) in (d_459_v22_) else d_456_cf65_), d_447_v20_, (0) - ((d_456_cf65_) + (d_458_cf63_)), d_455_cf66_)
            source6_ = d_460_v23_
            if source6_.is_DC28:
                d_461___mcc_h8_ = source6_.cf46
                d_462___mcc_h9_ = source6_.cf47
                d_463_cf47_ = d_462___mcc_h9_
                d_464_cf46_ = d_461___mcc_h8_
                d_463_cf47_ = (p0) + (d_463_cf47_)
                d_465_v24_: _dafny.Seq
                d_465_v24_ = _dafny.SeqWithoutIsStrInference([((d_448_v21_)[default__.safeIndex(d_458_cf63_, len(d_448_v21_))]) <= (d_463_cf47_)])
                d_465_v24_ = d_465_v24_
                d_458_cf63_ = d_456_cf65_
                d_456_cf65_ = d_463_cf47_
            elif source6_.is_DC29:
                d_466___mcc_h10_ = source6_.cf48
                d_467___mcc_h11_ = source6_.cf49
                d_468___mcc_h12_ = source6_.cf50
                d_469___mcc_h13_ = source6_.cf51
                d_470___mcc_h14_ = source6_.cf52
                d_471_cf52_ = d_470___mcc_h14_
                d_472_cf51_ = d_469___mcc_h13_
                d_473_cf50_ = d_468___mcc_h12_
                d_474_cf49_ = d_467___mcc_h11_
                d_475_cf48_ = d_466___mcc_h10_
                d_476_v25_: _dafny.Array
                nw74_ = _dafny.Array(int(0), 7)
                d_476_v25_ = nw74_
                index59_ = default__.safeIndex(982, (d_476_v25_).length(0))
                (d_476_v25_)[index59_] = (d_456_cf65_) - (p0)
                index60_ = default__.safeIndex(982, (d_476_v25_).length(0))
                (d_476_v25_)[index60_] = d_474_cf49_
                index61_ = default__.safeIndex(982, (d_476_v25_).length(0))
                (d_476_v25_)[index61_] = (len(d_419_v0_)) - (d_458_cf63_)
                d_477_v26_: D6
                d_477_v26_ = D6_DC14(d_419_v0_)
                d_478_v27_: str
                d_478_v27_ = _dafny.CodePoint('d')
                d_479_v28_: _dafny.Seq
                d_479_v28_ = _dafny.SeqWithoutIsStrInference([d_457_cf64_])
                d_480_v29_: _dafny.Array
                nw75_ = _dafny.Array(None, 28)
                nw75_[int(0)] = (((d_477_v26_).cf29).set(default__.safeIndex(default__.fm44(d_475_cf48_, globalState), len((d_477_v26_).cf29)), d_478_v27_)) + (d_419_v0_)
                nw75_[int(1)] = d_419_v0_
                nw75_[int(2)] = (d_457_cf64_) + (default__.fm46(True, len(_dafny.SeqWithoutIsStrInference([d_473_cf50_])), d_475_cf48_, globalState))
                nw75_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                nw75_[int(4)] = d_419_v0_
                nw75_[int(5)] = d_457_cf64_
                nw75_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rc"))
                nw75_[int(7)] = (d_457_cf64_) + (d_419_v0_)
                nw75_[int(8)] = d_419_v0_
                nw75_[int(9)] = d_457_cf64_
                nw75_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovqriexd"))
                nw75_[int(11)] = default__.fm46(True, default__.fm44(d_475_cf48_, globalState), d_473_cf50_, globalState)
                nw75_[int(12)] = _dafny.SeqWithoutIsStrInference([d_478_v27_ for d_481_i4_ in range(default__.abs(134))])
                nw75_[int(13)] = d_419_v0_
                nw75_[int(14)] = (d_477_v26_).cf29
                nw75_[int(15)] = d_419_v0_
                nw75_[int(16)] = d_457_cf64_
                nw75_[int(17)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_482_i5_ in range(default__.abs(552))])
                nw75_[int(18)] = d_419_v0_
                nw75_[int(19)] = d_419_v0_
                nw75_[int(20)] = _dafny.SeqWithoutIsStrInference([d_478_v27_ for d_483_i6_ in range(default__.abs(-733))])
                nw75_[int(21)] = (d_457_cf64_) + ((d_479_v28_)[default__.safeIndex(d_456_cf65_, len(d_479_v28_))])
                nw75_[int(22)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crcuft"))) + (_dafny.SeqWithoutIsStrInference([d_478_v27_ for d_484_i7_ in range(default__.abs(599))]))
                nw75_[int(23)] = (D6_DC14(d_457_cf64_)).cf29
                nw75_[int(24)] = d_419_v0_
                nw75_[int(25)] = d_419_v0_
                nw75_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nys"))
                nw75_[int(27)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jijhsve"))
                d_480_v29_ = nw75_
                index62_ = default__.safeIndex(595, (d_480_v29_).length(0))
                (d_480_v29_)[index62_] = d_457_cf64_
                d_485_v30_: _dafny.MultiSet
                d_485_v30_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfisdg"))).set(default__.safeIndex(d_458_cf63_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfisdg")))), d_478_v27_), d_457_cf64_])
                d_486_v31_: _dafny.Map
                d_486_v31_ = _dafny.Map({d_455_cf66_: d_485_v30_})
                d_487_v32_: _dafny.Map
                d_487_v32_ = _dafny.Map({p0: d_475_cf48_})
                index63_ = default__.safeIndex(595, (d_480_v29_).length(0))
                rhs51_ = d_476_v25_
                rhs52_ = d_419_v0_
                rhs53_ = (0) - (d_456_cf65_)
                rhs54_ = default__.safeModuloInt(-480, (((d_486_v31_)[((d_487_v32_)[(d_476_v25_)[default__.safeIndex(982, (d_476_v25_).length(0))]] if ((d_476_v25_)[default__.safeIndex(982, (d_476_v25_).length(0))]) in (d_487_v32_) else d_475_cf48_)] if (((d_487_v32_)[(d_476_v25_)[default__.safeIndex(982, (d_476_v25_).length(0))]] if ((d_476_v25_)[default__.safeIndex(982, (d_476_v25_).length(0))]) in (d_487_v32_) else d_475_cf48_)) in (d_486_v31_) else d_485_v30_)).cardinality)
                lhs32_ = d_480_v29_
                lhs33_ = default__.safeIndex(595, (d_480_v29_).length(0))
                d_476_v25_ = rhs51_
                lhs32_[lhs33_] = rhs52_
                d_458_cf63_ = rhs53_
                d_458_cf63_ = rhs54_
                d_454_cf67_ = d_475_cf48_
            elif True:
                d_488___mcc_h15_ = source6_.cf45
                d_489_cf45_ = d_488___mcc_h15_
                d_490_v33_: _dafny.Array
                def lambda18_(d_491_i8_):
                    return (d_491_i8_) - (len(_dafny.SeqWithoutIsStrInference([True])))

                init10_ = lambda18_
                nw76_ = _dafny.Array(None, 29)
                for i0_10_ in range(nw76_.length(0)):
                    nw76_[i0_10_] = init10_(i0_10_)
                d_490_v33_ = nw76_
                d_492_v34_: _dafny.Array
                d_492_v34_ = d_490_v33_
                d_492_v34_ = d_492_v34_
                index64_ = default__.safeIndex(921, (d_490_v33_).length(0))
                (d_490_v33_)[index64_] = d_458_cf63_
                index65_ = default__.safeIndex(921, (d_490_v33_).length(0))
                (d_490_v33_)[index65_] = (0) - (d_456_cf65_)
                d_493_v35_: str
                d_493_v35_ = _dafny.CodePoint('h')
                d_494_v36_: D0
                d_494_v36_ = D0_DC1(len(d_457_cf64_))
                d_495_v37_: _dafny.Map
                d_495_v37_ = _dafny.Map({d_493_v35_: d_494_v36_})
                d_496_v38_: _dafny.Map
                d_496_v38_ = _dafny.Map({not(d_455_cf66_): d_454_cf67_})
                d_497_v39_: _dafny.Array
                nw77_ = _dafny.Array(None, 28)
                nw77_[int(0)] = d_447_v20_
                nw77_[int(1)] = d_455_cf66_
                nw77_[int(2)] = False
                nw77_[int(3)] = d_454_cf67_
                nw77_[int(4)] = d_455_cf66_
                nw77_[int(5)] = d_454_cf67_
                nw77_[int(6)] = d_454_cf67_
                nw77_[int(7)] = d_454_cf67_
                nw77_[int(8)] = d_455_cf66_
                nw77_[int(9)] = d_455_cf66_
                nw77_[int(10)] = d_454_cf67_
                nw77_[int(11)] = (d_456_cf65_) < ((0) - (d_458_cf63_))
                nw77_[int(12)] = not(not(not(d_447_v20_)))
                nw77_[int(13)] = (self).fm42(d_455_cf66_, globalState)
                nw77_[int(14)] = d_447_v20_
                nw77_[int(15)] = d_455_cf66_
                nw77_[int(16)] = (self).fm42(d_455_cf66_, globalState)
                nw77_[int(17)] = d_454_cf67_
                nw77_[int(18)] = (d_456_cf65_) <= (d_456_cf65_)
                nw77_[int(19)] = d_454_cf67_
                nw77_[int(20)] = d_454_cf67_
                nw77_[int(21)] = not(d_455_cf66_)
                nw77_[int(22)] = (self).fm5((d_490_v33_)[default__.safeIndex(921, (d_490_v33_).length(0))], d_455_cf66_, d_495_v37_, (self).fm5(p0, True, _dafny.Map({d_493_v35_: d_494_v36_}), d_455_cf66_, globalState), globalState)
                nw77_[int(23)] = False
                nw77_[int(24)] = (d_455_cf66_) == (d_454_cf67_)
                nw77_[int(25)] = ((d_496_v38_)[(self).fm42(d_447_v20_, globalState)] if ((self).fm42(d_447_v20_, globalState)) in (d_496_v38_) else d_447_v20_)
                nw77_[int(26)] = d_447_v20_
                nw77_[int(27)] = d_455_cf66_
                d_497_v39_ = nw77_
                d_498_v40_: D2
                d_498_v40_ = D2_DC5(d_456_cf65_, d_448_v21_, d_447_v20_)
                index66_ = default__.safeIndex(7, (d_497_v39_).length(0))
                (d_497_v39_)[index66_] = ((d_498_v40_).cf10 if d_455_cf66_ else d_454_cf67_)
                index67_ = default__.safeIndex(7, (d_497_v39_).length(0))
                (d_497_v39_)[index67_] = (d_454_cf67_) or (d_455_cf66_)
                d_490_v33_ = d_490_v33_
            d_499_v41_: _dafny.Array
            def lambda19_(d_500_v20_):
                def lambda20_(d_501_i9_):
                    return d_500_v20_

                return lambda20_

            init11_ = lambda19_(d_447_v20_)
            nw78_ = _dafny.Array(None, 27)
            for i0_11_ in range(nw78_.length(0)):
                nw78_[i0_11_] = init11_(i0_11_)
            d_499_v41_ = nw78_
            d_502_v42_: _dafny.Seq
            d_502_v42_ = _dafny.SeqWithoutIsStrInference([True, False])
            index68_ = default__.safeIndex(737, (d_499_v41_).length(0))
            (d_499_v41_)[index68_] = (d_502_v42_)[default__.safeIndex(p0, len(d_502_v42_))]
            d_503_v43_: _dafny.Array
            def lambda21_(d_504_p0_):
                def lambda22_(d_505_i10_):
                    return (d_505_i10_) * (d_504_p0_)

                return lambda22_

            init12_ = lambda21_(p0)
            nw79_ = _dafny.Array(None, 5)
            for i0_12_ in range(nw79_.length(0)):
                nw79_[i0_12_] = init12_(i0_12_)
            d_503_v43_ = nw79_
            index69_ = default__.safeIndex(919, (d_503_v43_).length(0))
            (d_503_v43_)[index69_] = d_456_cf65_
            d_506_v44_: str
            d_506_v44_ = _dafny.CodePoint('t')
            d_507_v45_: D0
            d_507_v45_ = D0_DC1(p0)
            d_508_v46_: _dafny.Map
            d_508_v46_ = _dafny.Map({d_506_v44_: d_507_v45_})
            index70_ = default__.safeIndex(737, (d_499_v41_).length(0))
            index71_ = default__.safeIndex(919, (d_503_v43_).length(0))
            rhs55_ = (d_419_v0_) == (d_419_v0_)
            rhs56_ = default__.fm44((self).fm5(d_456_cf65_, d_455_cf66_, d_508_v46_, d_455_cf66_, globalState), globalState)
            lhs34_ = d_499_v41_
            lhs35_ = default__.safeIndex(737, (d_499_v41_).length(0))
            lhs36_ = d_503_v43_
            lhs37_ = default__.safeIndex(919, (d_503_v43_).length(0))
            lhs34_[lhs35_] = rhs55_
            lhs36_[lhs37_] = rhs56_
            if d_447_v20_:
                d_509_v47_: _dafny.Map
                d_509_v47_ = _dafny.Map({p0: d_455_cf66_})
                d_510_v48_: _dafny.Set
                d_510_v48_ = _dafny.Set({(d_509_v47_).set((d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))], (d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))]), _dafny.Map({len(default__.fm47(len(d_419_v0_), (d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))], globalState)): (d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))]})})
                d_511_v49_: D13
                d_511_v49_ = D13_DC31(False, len(_dafny.Map({d_447_v20_: d_456_cf65_})), (d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))])
                d_512_v50_: _dafny.MultiSet
                d_512_v50_ = _dafny.MultiSet([d_458_cf63_, len(d_510_v48_), default__.safeModuloInt(len(d_502_v42_), (d_511_v49_).cf56)])
                d_513_v51_: D7
                d_513_v51_ = D7_DC18()
                d_512_v50_ = (_dafny.MultiSet([(0) - (default__.fm44(False, globalState)), d_456_cf65_])) | (default__.fm48(d_513_v51_, (d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))], default__.fm44(True, globalState), -50, globalState))
                d_514_v52_: _dafny.Map
                d_514_v52_ = _dafny.Map({_dafny.Map({d_506_v44_: not((d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))])}): d_456_cf65_})
                d_514_v52_ = ((d_514_v52_ if d_447_v20_ else d_514_v52_)) | (d_514_v52_)
                d_515_v53_: _dafny.Array
                nw80_ = _dafny.Array(_dafny.Map({}), 24)
                d_515_v53_ = nw80_
                d_516_v54_: C1
                nw81_ = C1()
                nw81_.ctor__(d_458_cf63_, d_515_v53_, (d_448_v21_) + (d_448_v21_), d_507_v45_)
                d_516_v54_ = nw81_
                d_517_v55_: _dafny.Map
                d_517_v55_ = _dafny.Map({default__.fm46((d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))], (d_512_v50_).cardinality, False, globalState): (d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))]})
                d_518_v56_: _dafny.Map
                d_518_v56_ = _dafny.Map({d_448_v21_: d_457_cf64_})
                d_519_v57_: _dafny.Seq
                d_519_v57_ = _dafny.SeqWithoutIsStrInference([d_419_v0_, ((d_518_v56_)[d_448_v21_] if (d_448_v21_) in (d_518_v56_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmjnwlr"))), d_457_cf64_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))), d_506_v44_)])
                d_520_v58_: _dafny.Set
                d_520_v58_ = _dafny.Set({(d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))], d_447_v20_, True})
                index72_ = default__.safeIndex(737, (d_499_v41_).length(0))
                (d_499_v41_)[index72_] = ((d_517_v55_)[(d_519_v57_)[default__.safeIndex(558, len(d_519_v57_))]] if ((d_519_v57_)[default__.safeIndex(558, len(d_519_v57_))]) in (d_517_v55_) else (d_520_v58_).ispropersubset(_dafny.Set({(d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))]})))
                d_521_v59_: _dafny.Map
                d_521_v59_ = _dafny.Map({d_506_v44_: d_516_v54_.f9})
                d_522_v60_: _dafny.Map
                d_522_v60_ = _dafny.Map({(d_521_v59_) | (d_521_v59_): d_506_v44_})
                d_522_v60_ = d_522_v60_
            elif True:
                d_523_v61_: C0
                nw82_ = C0()
                nw82_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acttpcc")))
                d_523_v61_ = nw82_
                d_524_v62_: _dafny.Array
                nw83_ = _dafny.Array(None, 1)
                nw83_[int(0)] = d_502_v42_
                d_524_v62_ = nw83_
                index73_ = default__.safeIndex(392, (d_524_v62_).length(0))
                (d_524_v62_)[index73_] = d_502_v42_
                index74_ = default__.safeIndex(392, (d_524_v62_).length(0))
                (d_524_v62_)[index74_] = d_502_v42_
                d_447_v20_ = d_454_cf67_
                d_456_cf65_ = p0
                d_525_v63_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.Map({}), 14)
                d_525_v63_ = nw84_
                d_526_v64_: C1
                nw85_ = C1()
                nw85_.ctor__((d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))], d_525_v63_, (d_448_v21_).set(default__.safeIndex(d_456_cf65_, len(d_448_v21_)), d_456_cf65_), d_507_v45_)
                d_526_v64_ = nw85_
            if d_454_cf67_:
                d_527_v65_: _dafny.Map
                d_527_v65_ = _dafny.Map({p0: (d_419_v0_) <= (_dafny.SeqWithoutIsStrInference([d_506_v44_]))})
                d_527_v65_ = (d_527_v65_).set((d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))], True)
                d_447_v20_ = not((d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))])
                d_528_v66_: _dafny.Set
                d_528_v66_ = _dafny.Set({d_448_v21_})
                d_529_v67_: D15
                d_529_v67_ = D15_DC37(d_528_v66_)
                index75_ = default__.safeIndex(737, (d_499_v41_).length(0))
                (d_499_v41_)[index75_] = (d_528_v66_).issubset((d_529_v67_).cf70)
                index76_ = default__.safeIndex(737, (d_499_v41_).length(0))
                (d_499_v41_)[index76_] = d_455_cf66_
                index77_ = default__.safeIndex(919, (d_503_v43_).length(0))
                (d_503_v43_)[index77_] = default__.safeDivisionInt(p0, (d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))])
            elif True:
                index78_ = default__.safeIndex(919, (d_503_v43_).length(0))
                (d_503_v43_)[index78_] = d_456_cf65_
                d_530_v68_: _dafny.Set
                d_530_v68_ = _dafny.Set({(d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))], (self).fm5(p0, True, _dafny.Map({(d_419_v0_)[default__.safeIndex(len(_dafny.Map({d_458_cf63_: (d_499_v41_)[default__.safeIndex(737, (d_499_v41_).length(0))]})), len(d_419_v0_))]: d_507_v45_}), False, globalState)})
                d_531_v69_: _dafny.Map
                d_531_v69_ = _dafny.Map({d_530_v68_: d_447_v20_})
                d_456_cf65_ = ((d_503_v43_)[default__.safeIndex(919, (d_503_v43_).length(0))]) - (len(d_531_v69_))
                d_454_cf67_ = d_454_cf67_
                d_532_v70_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.Map({}), 8)
                d_532_v70_ = nw86_
                pat_let_tv9_ = d_458_cf63_
                d_533_v71_: C1
                nw87_ = C1()
                def iife62_(_pat_let12_0):
                    def iife63_(d_535_dt__update__tmp_h1_):
                        def iife64_(_pat_let13_0):
                            def iife65_(d_536_dt__update_hcf2_h0_):
                                return D0_DC1(d_536_dt__update_hcf2_h0_)
                            return iife65_(_pat_let13_0)
                        return iife64_(pat_let_tv9_)
                    return iife63_(_pat_let12_0)
                nw87_.ctor__(-834, d_532_v70_, _dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([d_447_v20_])).cardinality) for d_534_i11_ in range(default__.abs(-92))]), iife62_(d_507_v45_))
                d_533_v71_ = nw87_
                index79_ = default__.safeIndex(737, (d_499_v41_).length(0))
                (d_499_v41_)[index79_] = False
        elif source5_.is_DC36:
            d_537___mcc_h5_ = source5_.cf68
            d_538___mcc_h6_ = source5_.cf69
            d_539_cf69_ = d_538___mcc_h6_
            d_540_cf68_ = d_537___mcc_h5_
            d_541_v72_: _dafny.Array
            def lambda23_(d_542_p0_):
                def lambda24_(d_543_i12_):
                    return (d_543_i12_) - (d_542_p0_)

                return lambda24_

            init13_ = lambda23_(p0)
            nw88_ = _dafny.Array(None, 3)
            for i0_13_ in range(nw88_.length(0)):
                nw88_[i0_13_] = init13_(i0_13_)
            d_541_v72_ = nw88_
            d_544_v73_: _dafny.Map
            d_544_v73_ = _dafny.Map({True: d_541_v72_})
            d_545_v74_: D12
            d_545_v74_ = D12_DC28((d_544_v73_) == ((d_544_v73_).set(d_447_v20_, d_541_v72_)), p0)
            d_545_v74_ = d_545_v74_
            d_546_v75_: _dafny.Map
            d_546_v75_ = _dafny.Map({not(False): not (d_447_v20_) or (d_447_v20_)})
            d_546_v75_ = (d_546_v75_).set((d_447_v20_) not in (d_539_cf69_), d_447_v20_)
            d_547_v76_: _dafny.Map
            d_547_v76_ = _dafny.Map({p0: d_447_v20_})
            d_548_v77_: _dafny.Map
            d_548_v77_ = _dafny.Map({d_447_v20_: d_547_v76_})
            d_539_cf69_ = (d_539_cf69_).set((d_447_v20_) not in (d_548_v77_), p0)
            if d_447_v20_:
                d_549_v78_: int
                d_549_v78_ = 633
                d_549_v78_ = p0
                d_550_v79_: _dafny.Set
                d_550_v79_ = _dafny.Set({True, d_447_v20_})
                d_551_v80_: _dafny.MultiSet
                d_551_v80_ = _dafny.MultiSet([d_550_v79_])
                d_552_v81_: _dafny.Map
                d_552_v81_ = _dafny.Map({d_447_v20_: d_551_v80_})
                d_551_v80_ = ((d_551_v80_).set(d_550_v79_, default__.abs(d_549_v78_))) | (((d_552_v81_)[d_447_v20_] if (d_447_v20_) in (d_552_v81_) else _dafny.MultiSet([d_550_v79_])))
                d_544_v73_ = (d_544_v73_).set(((d_546_v75_)[d_447_v20_] if (d_447_v20_) in (d_546_v75_) else d_447_v20_), d_541_v72_)
                d_553_v82_: _dafny.Seq
                d_553_v82_ = _dafny.SeqWithoutIsStrInference([not(d_447_v20_), d_447_v20_, d_447_v20_])
                d_554_v83_: _dafny.Seq
                d_554_v83_ = _dafny.SeqWithoutIsStrInference([d_553_v82_, d_553_v82_])
                d_555_v84_: _dafny.Map
                d_555_v84_ = _dafny.Map({(d_553_v82_)[default__.safeIndex(d_549_v78_, len(d_553_v82_))]: (d_554_v83_)[default__.safeIndex((d_448_v21_)[default__.safeIndex(d_549_v78_, len(d_448_v21_))], len(d_554_v83_))]})
                d_555_v84_ = d_555_v84_
                d_556_v85_: C0
                nw89_ = C0()
                nw89_.ctor__(d_419_v0_)
                d_556_v85_ = nw89_
            elif True:
                d_557_v86_: _dafny.Array
                nw90_ = _dafny.Array(False, 19)
                d_557_v86_ = nw90_
                index80_ = default__.safeIndex(133, (d_557_v86_).length(0))
                (d_557_v86_)[index80_] = d_447_v20_
                index81_ = default__.safeIndex(133, (d_557_v86_).length(0))
                rhs57_ = d_448_v21_
                rhs58_ = not (d_447_v20_) or (d_447_v20_)
                lhs38_ = d_557_v86_
                lhs39_ = default__.safeIndex(133, (d_557_v86_).length(0))
                d_448_v21_ = rhs57_
                lhs38_[lhs39_] = rhs58_
                d_558_v87_: D6
                d_558_v87_ = D6_DC15(p0, default__.fm44((d_557_v86_)[default__.safeIndex(133, (d_557_v86_).length(0))], globalState), not(d_447_v20_))
                d_558_v87_ = d_558_v87_
                index82_ = default__.safeIndex(133, (d_557_v86_).length(0))
                (d_557_v86_)[index82_] = (d_419_v0_) <= (d_419_v0_)
                d_559_v88_: _dafny.Map
                d_559_v88_ = _dafny.Map({d_539_cf69_: d_419_v0_})
                d_560_v89_: int
                out20_: int
                out20_ = (self).m14(False, ((d_557_v86_)[default__.safeIndex(133, (d_557_v86_).length(0))]) and ((d_557_v86_)[default__.safeIndex(133, (d_557_v86_).length(0))]), (d_559_v88_) | (_dafny.Map({d_539_cf69_: d_419_v0_})), (0) - (len(d_419_v0_)), globalState)
                d_560_v89_ = out20_
                d_561_v90_: str
                d_561_v90_ = _dafny.CodePoint('s')
                d_562_v91_: _dafny.Map
                d_562_v91_ = _dafny.Map({False: d_561_v90_})
                d_562_v91_ = (d_562_v91_).set((d_557_v86_)[default__.safeIndex(133, (d_557_v86_).length(0))], d_561_v90_)
        elif True:
            d_563___mcc_h7_ = source5_.cf62
            d_564_cf62_ = d_563___mcc_h7_
            d_565_v92_: D12
            d_565_v92_ = D12_DC28(not (d_447_v20_) or (d_447_v20_), p0)
            d_565_v92_ = d_565_v92_
            d_447_v20_ = d_447_v20_
            d_566_v93_: _dafny.Array
            def lambda25_(d_567_v20_):
                def lambda26_(d_568_i13_):
                    return (_dafny.Set({d_567_v20_, d_567_v20_})) - (_dafny.Set({True}))

                return lambda26_

            init14_ = lambda25_(d_447_v20_)
            nw91_ = _dafny.Array(None, 9)
            for i0_14_ in range(nw91_.length(0)):
                nw91_[i0_14_] = init14_(i0_14_)
            d_566_v93_ = nw91_
            d_569_v94_: _dafny.Set
            d_569_v94_ = _dafny.Set({d_447_v20_})
            index83_ = default__.safeIndex(60, (d_566_v93_).length(0))
            (d_566_v93_)[index83_] = d_569_v94_
            index84_ = default__.safeIndex(60, (d_566_v93_).length(0))
            (d_566_v93_)[index84_] = d_569_v94_
            d_570_v95_: bool
            d_570_v95_ = d_447_v20_
            if not((d_570_v95_)):
                d_571_v96_: str
                d_571_v96_ = _dafny.CodePoint('c')
                d_572_v97_: _dafny.Map
                d_572_v97_ = _dafny.Map({(default__.fm44(d_447_v20_, globalState)) * (p0): p0})
                rhs59_ = (d_571_v96_ if d_447_v20_ else d_571_v96_)
                rhs60_ = (d_447_v20_) and (d_447_v20_)
                rhs61_ = d_572_v97_
                rhs62_ = (_dafny.Map({p0: (0) - (p0)})) == ((d_572_v97_) | (d_572_v97_))
                d_571_v96_ = rhs59_
                d_447_v20_ = rhs60_
                d_572_v97_ = rhs61_
                d_447_v20_ = rhs62_
                d_573_v98_: _dafny.Map
                d_573_v98_ = _dafny.Map({(p0) + (p0): (d_447_v20_) in ((d_566_v93_)[default__.safeIndex(60, (d_566_v93_).length(0))])})
                d_573_v98_ = (d_573_v98_).set(p0, True)
                d_447_v20_ = d_447_v20_
                d_574_v99_: _dafny.MultiSet
                d_574_v99_ = _dafny.MultiSet([True, d_447_v20_, False])
                d_575_v100_: _dafny.Map
                d_575_v100_ = _dafny.Map({p0: d_574_v99_})
                d_576_v101_: _dafny.Array
                nw92_ = _dafny.Array(None, 19)
                nw92_[int(0)] = d_574_v99_
                nw92_[int(1)] = d_574_v99_
                nw92_[int(2)] = (d_574_v99_) | (d_574_v99_)
                nw92_[int(3)] = _dafny.MultiSet([d_447_v20_])
                nw92_[int(4)] = d_574_v99_
                nw92_[int(5)] = (d_574_v99_) | (d_574_v99_)
                nw92_[int(6)] = _dafny.MultiSet([d_447_v20_, d_447_v20_])
                nw92_[int(7)] = d_574_v99_
                nw92_[int(8)] = d_574_v99_
                nw92_[int(9)] = d_574_v99_
                nw92_[int(10)] = d_574_v99_
                nw92_[int(11)] = d_574_v99_
                nw92_[int(12)] = ((d_575_v100_)[p0] if (p0) in (d_575_v100_) else d_574_v99_)
                nw92_[int(13)] = d_574_v99_
                nw92_[int(14)] = d_574_v99_
                nw92_[int(15)] = d_574_v99_
                nw92_[int(16)] = (_dafny.MultiSet([d_447_v20_, d_447_v20_])).set(d_447_v20_, default__.abs(p0))
                nw92_[int(17)] = d_574_v99_
                nw92_[int(18)] = d_574_v99_
                d_576_v101_ = nw92_
                d_577_v102_: _dafny.Seq
                d_577_v102_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_565_v92_).cf46])])
                index85_ = default__.safeIndex(741, (d_576_v101_).length(0))
                (d_576_v101_)[index85_] = (d_577_v102_)[default__.safeIndex(p0, len(d_577_v102_))]
                index86_ = default__.safeIndex(741, (d_576_v101_).length(0))
                (d_576_v101_)[index86_] = d_574_v99_
                d_447_v20_ = (p0) >= (p0)
            elif True:
                d_578_v103_: _dafny.Array
                nw93_ = _dafny.Array(False, 24)
                d_578_v103_ = nw93_
                d_579_v104_: D16
                d_579_v104_ = D16_DC41(d_578_v103_)
                d_578_v103_ = (d_579_v104_).cf75
                d_580_v105_: int
                d_580_v105_ = 795
                d_581_v106_: _dafny.Map
                d_581_v106_ = _dafny.Map({d_580_v105_: d_447_v20_})
                d_582_v107_: D0
                d_582_v107_ = D0_DC1(d_580_v105_)
                d_580_v105_ = default__.safeModuloInt((p0) * (default__.fm44(((d_581_v106_)[p0] if (p0) in (d_581_v106_) else d_447_v20_), globalState)), ((d_582_v107_).cf2) - (len(_dafny.SeqWithoutIsStrInference([d_447_v20_]))))
                d_447_v20_ = not(d_447_v20_)
                d_447_v20_ = (p0) != ((0) - (default__.safeDivisionInt(len(d_419_v0_), p0)))
                d_583_v108_: _dafny.Array
                nw94_ = _dafny.Array(D0.default()(), 13)
                d_583_v108_ = nw94_
                index87_ = default__.safeIndex(631, (d_583_v108_).length(0))
                (d_583_v108_)[index87_] = d_582_v107_
                d_584_v109_: str
                d_584_v109_ = _dafny.CodePoint('l')
                index88_ = default__.safeIndex(631, (d_583_v108_).length(0))
                (d_583_v108_)[index88_] = default__.fm49(d_584_v109_, globalState)
        if not(d_447_v20_):
            d_585_v110_: _dafny.MultiSet
            d_585_v110_ = _dafny.MultiSet([d_447_v20_])
            d_586_v111_: _dafny.Map
            d_586_v111_ = _dafny.Map({(d_585_v110_).cardinality: d_447_v20_})
            d_587_v112_: _dafny.Seq
            d_587_v112_ = _dafny.SeqWithoutIsStrInference([d_586_v111_])
            d_588_v113_: _dafny.Array
            nw95_ = _dafny.Array(None, 5)
            nw95_[int(0)] = d_586_v111_
            nw95_[int(1)] = d_586_v111_
            nw95_[int(2)] = default__.fm50(d_447_v20_, d_447_v20_, p0, globalState)
            nw95_[int(3)] = (d_587_v112_)[default__.safeIndex(p0, len(d_587_v112_))]
            nw95_[int(4)] = (_dafny.Map({p0: d_447_v20_})).set(-789, d_447_v20_)
            d_588_v113_ = nw95_
            d_589_v114_: _dafny.Map
            d_589_v114_ = _dafny.Map({d_447_v20_: d_588_v113_})
            d_590_v115_: C0
            nw96_ = C0()
            nw96_.ctor__(d_419_v0_)
            d_590_v115_ = nw96_
            d_591_v116_: _dafny.Seq
            d_591_v116_ = _dafny.SeqWithoutIsStrInference([d_590_v115_])
            d_589_v114_ = (d_589_v114_).set((len(d_591_v116_)) >= (p0), d_588_v113_)
            d_592_v117_: _dafny.Array
            nw97_ = _dafny.Array(int(0), 12)
            d_592_v117_ = nw97_
            index89_ = default__.safeIndex(99, (d_592_v117_).length(0))
            (d_592_v117_)[index89_] = default__.fm44(not(d_447_v20_), globalState)
            index90_ = default__.safeIndex(99, (d_592_v117_).length(0))
            (d_592_v117_)[index90_] = p0
            d_448_v21_ = (d_448_v21_) + (_dafny.SeqWithoutIsStrInference([p0 for d_593_i14_ in range(default__.abs(-305))]))
            d_447_v20_ = d_447_v20_
            d_594_v118_: _dafny.Set
            d_594_v118_ = _dafny.Set({p0, 151, (d_592_v117_)[default__.safeIndex(99, (d_592_v117_).length(0))]})
            d_594_v118_ = (d_594_v118_).intersection(d_594_v118_)
        elif True:
            d_447_v20_ = d_447_v20_
            d_595_v119_: _dafny.Array
            nw98_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_595_v119_ = nw98_
            d_596_v120_: _dafny.Set
            d_596_v120_ = _dafny.Set({d_447_v20_})
            d_597_v121_: _dafny.Seq
            d_597_v121_ = _dafny.SeqWithoutIsStrInference([d_596_v120_, d_596_v120_])
            d_598_v122_: D12
            d_598_v122_ = D12_DC27(d_596_v120_)
            d_599_v123_: _dafny.Array
            nw99_ = _dafny.Array(None, 18)
            nw99_[int(0)] = d_596_v120_
            nw99_[int(1)] = d_596_v120_
            nw99_[int(2)] = (d_597_v121_)[default__.safeIndex(p0, len(d_597_v121_))]
            nw99_[int(3)] = d_596_v120_
            nw99_[int(4)] = d_596_v120_
            nw99_[int(5)] = d_596_v120_
            nw99_[int(6)] = d_596_v120_
            nw99_[int(7)] = d_596_v120_
            nw99_[int(8)] = d_596_v120_
            nw99_[int(9)] = d_596_v120_
            nw99_[int(10)] = d_596_v120_
            nw99_[int(11)] = (d_598_v122_).cf45
            nw99_[int(12)] = d_596_v120_
            nw99_[int(13)] = d_596_v120_
            nw99_[int(14)] = d_596_v120_
            nw99_[int(15)] = d_596_v120_
            nw99_[int(16)] = d_596_v120_
            nw99_[int(17)] = d_596_v120_
            d_599_v123_ = nw99_
            d_600_v124_: D0
            d_600_v124_ = D0_DC0(d_599_v123_, 596)
            d_601_v125_: D15
            d_601_v125_ = D15_DC38(default__.fm51(d_447_v20_, p0, p0, globalState))
            d_602_v126_: _dafny.Map
            d_602_v126_ = _dafny.Map({d_600_v124_: (d_601_v125_).cf71})
            d_603_v127_: _dafny.Seq
            d_603_v127_ = _dafny.SeqWithoutIsStrInference([d_447_v20_])
            d_604_v128_: _dafny.Seq
            d_604_v128_ = _dafny.SeqWithoutIsStrInference([d_448_v21_])
            d_605_v129_: _dafny.Array
            nw100_ = _dafny.Array(None, 23)
            nw100_[int(0)] = p0
            nw100_[int(1)] = p0
            nw100_[int(2)] = 550
            nw100_[int(3)] = len(d_419_v0_)
            nw100_[int(4)] = default__.fm44(d_447_v20_, globalState)
            nw100_[int(5)] = p0
            nw100_[int(6)] = p0
            nw100_[int(7)] = (0) - ((0) - (p0))
            nw100_[int(8)] = p0
            nw100_[int(9)] = p0
            nw100_[int(10)] = p0
            nw100_[int(11)] = len(d_419_v0_)
            nw100_[int(12)] = p0
            nw100_[int(13)] = len(d_602_v126_)
            nw100_[int(14)] = len(d_603_v127_)
            nw100_[int(15)] = len(d_604_v128_)
            nw100_[int(16)] = p0
            nw100_[int(17)] = p0
            nw100_[int(18)] = (0) - (p0)
            nw100_[int(19)] = 259
            nw100_[int(20)] = p0
            nw100_[int(21)] = p0
            nw100_[int(22)] = -31
            d_605_v129_ = nw100_
            index91_ = default__.safeIndex(170, (d_595_v119_).length(0))
            (d_595_v119_)[index91_] = d_605_v129_
            d_606_v130_: _dafny.Array
            nw101_ = _dafny.Array(False, 14)
            d_606_v130_ = nw101_
            index92_ = default__.safeIndex(703, (d_606_v130_).length(0))
            (d_606_v130_)[index92_] = d_447_v20_
            index93_ = default__.safeIndex(483, (d_605_v129_).length(0))
            (d_605_v129_)[index93_] = default__.fm44(d_447_v20_, globalState)
            index94_ = default__.safeIndex(844, (d_605_v129_).length(0))
            (d_605_v129_)[index94_] = p0
            index95_ = default__.safeIndex(170, (d_595_v119_).length(0))
            index96_ = default__.safeIndex(703, (d_606_v130_).length(0))
            index97_ = default__.safeIndex(483, (d_605_v129_).length(0))
            index98_ = default__.safeIndex(844, (d_605_v129_).length(0))
            rhs63_ = d_605_v129_
            rhs64_ = d_447_v20_
            rhs65_ = p0
            rhs66_ = default__.safeModuloInt(p0, (d_448_v21_)[default__.safeIndex(p0, len(d_448_v21_))])
            lhs40_ = d_595_v119_
            lhs41_ = default__.safeIndex(170, (d_595_v119_).length(0))
            lhs42_ = d_606_v130_
            lhs43_ = default__.safeIndex(703, (d_606_v130_).length(0))
            lhs44_ = d_605_v129_
            lhs45_ = default__.safeIndex(483, (d_605_v129_).length(0))
            lhs46_ = d_605_v129_
            lhs47_ = default__.safeIndex(844, (d_605_v129_).length(0))
            lhs40_[lhs41_] = rhs63_
            lhs42_[lhs43_] = rhs64_
            lhs44_[lhs45_] = rhs65_
            lhs46_[lhs47_] = rhs66_
            d_447_v20_ = (self).fm42(d_447_v20_, globalState)
            d_607_v131_: _dafny.Set
            d_607_v131_ = _dafny.Set({d_606_v130_})
            d_608_v132_: _dafny.Set
            d_608_v132_ = d_607_v131_
            d_609_v133_: _dafny.Set
            d_609_v133_ = (d_608_v132_)
            source7_ = d_609_v133_
            d_610___mcc_h16_ = source7_
            d_611_cf36_ = d_610___mcc_h16_
            d_612_v134_: _dafny.Map
            d_612_v134_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_613_i15_ in range(default__.abs(518))]): (0) - (p0)})
            d_614_v135_: _dafny.Map
            d_614_v135_ = _dafny.Map({len(d_612_v134_): (d_606_v130_)[default__.safeIndex(703, (d_606_v130_).length(0))]})
            d_615_v137_: str
            d_615_v137_ = _dafny.CodePoint('f')
            def iife66_():
                coll38_ = _dafny.Map()
                compr_38_: str
                for compr_38_ in (_dafny.MultiSet([d_615_v137_, d_615_v137_, d_615_v137_, d_615_v137_, d_615_v137_])).Elements:
                    d_616_v136_: str = compr_38_
                    if (d_616_v136_) in (_dafny.MultiSet([d_615_v137_, d_615_v137_, d_615_v137_, d_615_v137_, d_615_v137_])):
                        coll38_[d_616_v136_] = D0_DC1(p0)
                return _dafny.Map(coll38_)
            d_447_v20_ = (self).fm5(len(d_614_v135_), not (d_447_v20_) or (d_447_v20_), iife66_()
            , d_447_v20_, globalState)
            index99_ = default__.safeIndex(607, (d_599_v123_).length(0))
            (d_599_v123_)[index99_] = d_596_v120_
            index100_ = default__.safeIndex(607, (d_599_v123_).length(0))
            (d_599_v123_)[index100_] = d_596_v120_
            d_617_v139_: _dafny.Array
            def lambda27_(d_618_v130_, d_619_v129_, d_620_p0_):
                def lambda28_(d_621_i16_):
                    def iife67_():
                        coll39_ = _dafny.Set()
                        compr_39_: _dafny.MultiSet
                        for compr_39_ in (_dafny.Set({_dafny.MultiSet([len(_dafny.Map({(d_618_v130_)[default__.safeIndex(703, (d_618_v130_).length(0))]: (d_618_v130_)[default__.safeIndex(703, (d_618_v130_).length(0))]})), d_620_p0_])})).Elements:
                            d_622_v138_: _dafny.MultiSet = compr_39_
                            if (d_622_v138_) in (_dafny.Set({_dafny.MultiSet([len(_dafny.Map({(d_618_v130_)[default__.safeIndex(703, (d_618_v130_).length(0))]: (d_618_v130_)[default__.safeIndex(703, (d_618_v130_).length(0))]})), d_620_p0_])})):
                                coll39_ = coll39_.union(_dafny.Set([d_622_v138_]))
                        return _dafny.Set(coll39_)
                    return (_dafny.MultiSet([_dafny.Map({(d_618_v130_)[default__.safeIndex(703, (d_618_v130_).length(0))]: (d_619_v129_)[default__.safeIndex(483, (d_619_v129_).length(0))]})])) | (_dafny.MultiSet([_dafny.Map({(d_618_v130_)[default__.safeIndex(703, (d_618_v130_).length(0))]: len(iife67_()
                    )})]))

                return lambda28_

            init15_ = lambda27_(d_606_v130_, d_605_v129_, p0)
            nw102_ = _dafny.Array(None, 6)
            for i0_15_ in range(nw102_.length(0)):
                nw102_[i0_15_] = init15_(i0_15_)
            d_617_v139_ = nw102_
            d_623_v140_: _dafny.Map
            d_623_v140_ = _dafny.Map({p0: (d_605_v129_)[default__.safeIndex(483, (d_605_v129_).length(0))]})
            d_624_v141_: _dafny.MultiSet
            d_624_v141_ = _dafny.MultiSet([_dafny.Map({not((d_606_v130_)[default__.safeIndex(703, (d_606_v130_).length(0))]): len(d_623_v140_)})])
            index101_ = default__.safeIndex(64, (d_617_v139_).length(0))
            (d_617_v139_)[index101_] = d_624_v141_
            index102_ = default__.safeIndex(64, (d_617_v139_).length(0))
            (d_617_v139_)[index102_] = d_624_v141_
            d_625_v142_: D13
            d_625_v142_ = D13_DC32(d_419_v0_, (d_606_v130_)[default__.safeIndex(703, (d_606_v130_).length(0))], (d_605_v129_)[default__.safeIndex(483, (d_605_v129_).length(0))], (d_595_v119_)[default__.safeIndex(170, (d_595_v119_).length(0))])
            d_626_v143_: _dafny.Map
            d_626_v143_ = _dafny.Map({not((d_625_v142_).cf58): p0})
            d_627_v144_: _dafny.MultiSet
            d_627_v144_ = _dafny.MultiSet([d_609_v133_])
            d_628_v145_: _dafny.Map
            d_628_v145_ = _dafny.Map({d_626_v143_: ((d_627_v144_)[_dafny.Set({d_606_v130_, d_606_v130_})] if (_dafny.Set({d_606_v130_, d_606_v130_})) in (d_627_v144_) else len(d_419_v0_))})
            d_629_v146_: D0
            d_629_v146_ = D0_DC1((d_605_v129_)[default__.safeIndex(483, (d_605_v129_).length(0))])
            d_630_v147_: _dafny.Map
            d_630_v147_ = _dafny.Map({d_615_v137_: d_629_v146_})
            d_628_v145_ = (_dafny.Map({_dafny.Map({d_447_v20_: p0}): ((d_626_v143_)[True] if (True) in (d_626_v143_) else p0)}) if (d_447_v20_) and ((self).fm5(630, d_447_v20_, d_630_v147_, (d_606_v130_)[default__.safeIndex(703, (d_606_v130_).length(0))], globalState)) else _dafny.Map({_dafny.Map({(d_606_v130_)[default__.safeIndex(703, (d_606_v130_).length(0))]: -451}): p0}))
            d_631_v148_: _dafny.Array
            def lambda29_(d_632_i17_):
                return _dafny.CodePoint('s')

            init16_ = lambda29_
            nw103_ = _dafny.Array(None, 19)
            for i0_16_ in range(nw103_.length(0)):
                nw103_[i0_16_] = init16_(i0_16_)
            d_631_v148_ = nw103_
            d_633_v149_: _dafny.Seq
            d_633_v149_ = _dafny.SeqWithoutIsStrInference([d_631_v148_])
            d_633_v149_ = d_633_v149_
        d_634_v150_: _dafny.Map
        d_634_v150_ = _dafny.Map({d_419_v0_: d_447_v20_})
        d_634_v150_ = (d_634_v150_).set((d_419_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_635_i18_ in range(default__.abs(680))])), d_447_v20_)
        if not(False):
            d_636_v151_: _dafny.Array
            nw104_ = _dafny.Array(int(0), 15)
            d_636_v151_ = nw104_
            d_637_v152_: _dafny.Array
            d_637_v152_ = d_636_v151_
            d_637_v152_ = d_637_v152_
            d_638_v153_: int
            d_638_v153_ = 404
            d_638_v153_ = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_639_i19_ in range(default__.abs(-90))])) + (d_419_v0_))
            d_640_v154_: str
            d_640_v154_ = _dafny.CodePoint('g')
            d_641_v155_: _dafny.Map
            d_641_v155_ = _dafny.Map({d_640_v154_: p0})
            d_641_v155_ = (d_641_v155_).set(d_640_v154_, d_638_v153_)
            d_640_v154_ = d_640_v154_
            d_638_v153_ = (0) - (p0)
        elif True:
            d_642_v156_: D4
            d_642_v156_ = D4_DC10(len(d_419_v0_), (0) - (p0), len(_dafny.Map({d_447_v20_: False})))
            d_643_v157_: str
            d_643_v157_ = _dafny.CodePoint('a')
            d_644_v158_: _dafny.Map
            d_644_v158_ = _dafny.Map({d_447_v20_: d_643_v157_})
            d_645_v159_: _dafny.Set
            d_645_v159_ = _dafny.Set({len(d_644_v158_)})
            d_646_v160_: _dafny.Array
            nw105_ = _dafny.Array(None, 24)
            nw105_[int(0)] = default__.fm44(d_447_v20_, globalState)
            nw105_[int(1)] = p0
            nw105_[int(2)] = 714
            nw105_[int(3)] = p0
            nw105_[int(4)] = p0
            nw105_[int(5)] = p0
            nw105_[int(6)] = (d_642_v156_).cf20
            nw105_[int(7)] = p0
            nw105_[int(8)] = p0
            nw105_[int(9)] = p0
            nw105_[int(10)] = p0
            nw105_[int(11)] = p0
            nw105_[int(12)] = p0
            nw105_[int(13)] = len(d_645_v159_)
            nw105_[int(14)] = p0
            nw105_[int(15)] = p0
            nw105_[int(16)] = p0
            nw105_[int(17)] = p0
            nw105_[int(18)] = len(d_419_v0_)
            nw105_[int(19)] = p0
            nw105_[int(20)] = p0
            nw105_[int(21)] = p0
            nw105_[int(22)] = p0
            nw105_[int(23)] = p0
            d_646_v160_ = nw105_
            d_647_v161_: _dafny.Array
            nw106_ = _dafny.Array(None, 12)
            nw106_[int(0)] = d_646_v160_
            nw106_[int(1)] = d_646_v160_
            nw106_[int(2)] = d_646_v160_
            nw106_[int(3)] = d_646_v160_
            nw106_[int(4)] = d_646_v160_
            nw106_[int(5)] = d_646_v160_
            nw106_[int(6)] = d_646_v160_
            nw106_[int(7)] = d_646_v160_
            nw106_[int(8)] = d_646_v160_
            nw106_[int(9)] = d_646_v160_
            nw106_[int(10)] = d_646_v160_
            nw106_[int(11)] = d_646_v160_
            d_647_v161_ = nw106_
            d_648_v162_: D2
            d_648_v162_ = D2_DC4(p0, d_447_v20_, d_647_v161_)
            source8_ = d_648_v162_
            if source8_.is_DC4:
                d_649___mcc_h17_ = source8_.cf5
                d_650___mcc_h18_ = source8_.cf6
                d_651___mcc_h19_ = source8_.cf7
                d_652_cf7_ = d_651___mcc_h19_
                d_653_cf6_ = d_650___mcc_h18_
                d_654_cf5_ = d_649___mcc_h17_
                d_447_v20_ = False
                d_655_v163_: _dafny.Array
                nw107_ = _dafny.Array(_dafny.Map({}), 16)
                d_655_v163_ = nw107_
                d_656_v164_: _dafny.MultiSet
                d_656_v164_ = _dafny.MultiSet([d_447_v20_])
                d_657_v165_: D9
                d_657_v165_ = D9_DC21(d_656_v164_)
                d_658_v166_: D15
                d_658_v166_ = D15_DC39(d_419_v0_, d_657_v165_)
                d_659_v167_: _dafny.Seq
                d_659_v167_ = _dafny.SeqWithoutIsStrInference([d_419_v0_, (d_658_v166_).cf72])
                d_660_v168_: C1
                nw108_ = C1()
                nw108_.ctor__(p0, d_655_v163_, (_dafny.SeqWithoutIsStrInference([len(d_448_v21_) for d_661_i20_ in range(default__.abs(378))])) + (d_448_v21_), D0_DC1(len(d_659_v167_)))
                d_660_v168_ = nw108_
                d_646_v160_ = d_646_v160_
                d_419_v0_ = d_419_v0_
            elif source8_.is_DC5:
                d_662___mcc_h20_ = source8_.cf8
                d_663___mcc_h21_ = source8_.cf9
                d_664___mcc_h22_ = source8_.cf10
                d_665_cf10_ = d_664___mcc_h22_
                d_666_cf9_ = d_663___mcc_h21_
                d_667_cf8_ = d_662___mcc_h20_
                index103_ = default__.safeIndex(240, (d_646_v160_).length(0))
                (d_646_v160_)[index103_] = 711
                d_668_v169_: _dafny.MultiSet
                d_668_v169_ = _dafny.MultiSet([True, d_447_v20_])
                index104_ = default__.safeIndex(240, (d_646_v160_).length(0))
                (d_646_v160_)[index104_] = default__.safeDivisionInt(default__.safeModuloInt(d_667_cf8_, len(d_419_v0_)), ((d_668_v169_)[d_447_v20_] if (d_447_v20_) in (d_668_v169_) else p0))
                d_669_v170_: _dafny.Seq
                d_669_v170_ = _dafny.SeqWithoutIsStrInference([d_419_v0_, d_419_v0_])
                d_670_v171_: _dafny.Map
                d_670_v171_ = _dafny.Map({d_669_v170_: p0})
                index105_ = default__.safeIndex(240, (d_646_v160_).length(0))
                (d_646_v160_)[index105_] = ((d_670_v171_)[_dafny.SeqWithoutIsStrInference([d_419_v0_ for d_671_i21_ in range(default__.abs(-291))])] if (_dafny.SeqWithoutIsStrInference([d_419_v0_ for d_672_i21_ in range(default__.abs(-291))])) in (d_670_v171_) else ((0) - (d_667_cf8_)) - (d_667_cf8_))
                d_673_v172_: _dafny.Array
                nw109_ = _dafny.Array(False, 10)
                d_673_v172_ = nw109_
                index106_ = default__.safeIndex(535, (d_673_v172_).length(0))
                (d_673_v172_)[index106_] = (d_665_cf10_) == (True)
                index107_ = default__.safeIndex(535, (d_673_v172_).length(0))
                (d_673_v172_)[index107_] = d_447_v20_
                index108_ = default__.safeIndex(240, (d_646_v160_).length(0))
                (d_646_v160_)[index108_] = default__.safeDivisionInt(default__.safeModuloInt((d_646_v160_)[default__.safeIndex(240, (d_646_v160_).length(0))], -912), (len(d_419_v0_)) * (p0))
            elif source8_.is_DC3:
                d_674___mcc_h23_ = source8_.cf4
                d_675_cf4_ = d_674___mcc_h23_
                d_676_v173_: int
                d_676_v173_ = 602
                d_677_v174_: D15
                d_677_v174_ = D15_DC38((default__.fm51(d_447_v20_, d_676_v173_, d_676_v173_, globalState) if d_447_v20_ else d_643_v157_))
                d_678_v175_: _dafny.Map
                d_678_v175_ = _dafny.Map({d_447_v20_: d_448_v21_})
                rhs67_ = len((((d_678_v175_)[d_447_v20_] if (d_447_v20_) in (d_678_v175_) else _dafny.SeqWithoutIsStrInference([p0 for d_679_i22_ in range(default__.abs(10))])) if d_447_v20_ else _dafny.SeqWithoutIsStrInference([d_676_v173_ for d_680_i23_ in range(default__.abs(-820))])))
                rhs68_ = (d_675_cf4_).intersection(d_675_cf4_)
                rhs69_ = d_677_v174_
                rhs70_ = d_447_v20_
                d_676_v173_ = rhs67_
                d_675_cf4_ = rhs68_
                d_677_v174_ = rhs69_
                d_447_v20_ = rhs70_
                d_676_v173_ = ((p0) + ((0) - (p0))) + (d_676_v173_)
                d_681_v176_: _dafny.Array
                nw110_ = _dafny.Array(False, 18)
                d_681_v176_ = nw110_
                d_682_v177_: _dafny.MultiSet
                d_682_v177_ = _dafny.MultiSet([d_447_v20_])
                index109_ = default__.safeIndex(612, (d_681_v176_).length(0))
                (d_681_v176_)[index109_] = (d_682_v177_).ispropersubset(_dafny.MultiSet([d_447_v20_, d_447_v20_]))
                index110_ = default__.safeIndex(612, (d_681_v176_).length(0))
                (d_681_v176_)[index110_] = d_447_v20_
                d_683_v178_: _dafny.Map
                d_683_v178_ = _dafny.Map({p0: d_447_v20_})
                d_683_v178_ = d_683_v178_
            elif True:
                d_684___mcc_h24_ = source8_.cf11
                d_685_cf11_ = d_684___mcc_h24_
                d_686_v179_: _dafny.Seq
                d_686_v179_ = _dafny.SeqWithoutIsStrInference([d_447_v20_, d_447_v20_])
                d_687_v180_: _dafny.Map
                d_687_v180_ = _dafny.Map({d_643_v157_: 68})
                d_688_v181_: D15
                d_688_v181_ = D15_DC38(d_643_v157_)
                d_689_v182_: _dafny.MultiSet
                d_689_v182_ = _dafny.MultiSet([p0, 317])
                d_690_v183_: _dafny.Array
                nw111_ = _dafny.Array(None, 9)
                nw111_[int(0)] = d_687_v180_
                nw111_[int(1)] = default__.fm52(True, d_688_v181_, d_447_v20_, globalState)
                nw111_[int(2)] = d_687_v180_
                nw111_[int(3)] = d_687_v180_
                nw111_[int(4)] = d_687_v180_
                nw111_[int(5)] = _dafny.Map({d_643_v157_: default__.fm44(d_447_v20_, globalState)})
                nw111_[int(6)] = _dafny.Map({d_643_v157_: ((d_689_v182_)[p0] if (p0) in (d_689_v182_) else p0)})
                nw111_[int(7)] = d_687_v180_
                nw111_[int(8)] = d_687_v180_
                d_690_v183_ = nw111_
                d_691_v184_: T1
                nw112_ = C1()
                nw112_.ctor__(len(d_686_v179_), d_690_v183_, d_448_v21_, D0_DC1(p0))
                d_691_v184_ = nw112_
                d_692_v185_: _dafny.Map
                d_692_v185_ = _dafny.Map({d_691_v184_: _dafny.CodePoint('b')})
                d_692_v185_ = (d_692_v185_).set(d_691_v184_, _dafny.CodePoint('e'))
                index111_ = default__.safeIndex(353, (d_646_v160_).length(0))
                (d_646_v160_)[index111_] = default__.safeModuloInt(p0, len(d_686_v179_))
                index112_ = default__.safeIndex(353, (d_646_v160_).length(0))
                (d_646_v160_)[index112_] = (602) - (98)
                d_693_v186_: C0
                nw113_ = C0()
                nw113_.ctor__(d_419_v0_)
                d_693_v186_ = nw113_
                d_694_v187_: _dafny.Array
                nw114_ = _dafny.Array(False, 20)
                d_694_v187_ = nw114_
                index113_ = default__.safeIndex(825, (d_694_v187_).length(0))
                (d_694_v187_)[index113_] = (d_447_v20_ if d_447_v20_ else d_447_v20_)
                d_695_v188_: _dafny.MultiSet
                d_695_v188_ = _dafny.MultiSet([d_686_v179_])
                index114_ = default__.safeIndex(353, (d_646_v160_).length(0))
                index115_ = default__.safeIndex(825, (d_694_v187_).length(0))
                rhs71_ = (d_689_v182_).cardinality
                rhs72_ = d_447_v20_
                rhs73_ = d_695_v188_
                lhs48_ = d_646_v160_
                lhs49_ = default__.safeIndex(353, (d_646_v160_).length(0))
                lhs50_ = d_694_v187_
                lhs51_ = default__.safeIndex(825, (d_694_v187_).length(0))
                lhs48_[lhs49_] = rhs71_
                lhs50_[lhs51_] = rhs72_
                d_695_v188_ = rhs73_
            index116_ = default__.safeIndex(653, (d_646_v160_).length(0))
            (d_646_v160_)[index116_] = p0
            d_696_v189_: _dafny.Map
            d_696_v189_ = _dafny.Map({False: p0})
            d_697_v190_: _dafny.Seq
            d_697_v190_ = _dafny.SeqWithoutIsStrInference([d_696_v189_, d_696_v189_, d_696_v189_, d_696_v189_])
            index117_ = default__.safeIndex(653, (d_646_v160_).length(0))
            (d_646_v160_)[index117_] = (p0) + (len(default__.fm53((d_697_v190_)[default__.safeIndex(p0, len(d_697_v190_))], globalState)))
            d_698_v191_: _dafny.Map
            d_698_v191_ = _dafny.Map({p0: d_648_v162_})
            d_699_v192_: _dafny.Map
            d_699_v192_ = _dafny.Map({d_447_v20_: d_698_v191_})
            d_700_v193_: _dafny.Map
            d_700_v193_ = _dafny.Map({p0: (d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]})
            index118_ = default__.safeIndex(653, (d_646_v160_).length(0))
            index119_ = default__.safeIndex(653, (d_646_v160_).length(0))
            index120_ = default__.safeIndex(653, (d_646_v160_).length(0))
            rhs74_ = (d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]
            rhs75_ = ((52) + (499) if (d_696_v189_) != (d_696_v189_) else p0)
            rhs76_ = ((d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]) > ((d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))])
            rhs77_ = (d_698_v191_) != (((d_699_v192_)[d_447_v20_] if (d_447_v20_) in (d_699_v192_) else _dafny.Map({(d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]: D2_DC4(438, d_447_v20_, d_647_v161_)})))
            rhs78_ = (p0) * (default__.safeDivisionInt(((d_700_v193_)[(_dafny.MultiSet(d_448_v21_)).cardinality] if ((_dafny.MultiSet(d_448_v21_)).cardinality) in (d_700_v193_) else p0), (d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]))
            lhs52_ = d_646_v160_
            lhs53_ = default__.safeIndex(653, (d_646_v160_).length(0))
            lhs54_ = d_646_v160_
            lhs55_ = default__.safeIndex(653, (d_646_v160_).length(0))
            lhs56_ = d_646_v160_
            lhs57_ = default__.safeIndex(653, (d_646_v160_).length(0))
            lhs52_[lhs53_] = rhs74_
            lhs54_[lhs55_] = rhs75_
            d_447_v20_ = rhs76_
            d_447_v20_ = rhs77_
            lhs56_[lhs57_] = rhs78_
            d_701_v194_: _dafny.Map
            d_701_v194_ = _dafny.Map({default__.fm44(not(d_447_v20_), globalState): d_447_v20_})
            d_702_v195_: _dafny.Set
            d_702_v195_ = _dafny.Set({d_643_v157_})
            d_447_v20_ = ((d_701_v194_)[p0] if (p0) in (d_701_v194_) else not((_dafny.Set({d_643_v157_, d_643_v157_, d_643_v157_, d_643_v157_})).ispropersubset(d_702_v195_)))
            d_703_v196_: _dafny.Seq
            d_703_v196_ = _dafny.SeqWithoutIsStrInference([d_447_v20_])
            d_704_v197_: _dafny.Set
            d_704_v197_ = _dafny.Set({d_447_v20_})
            d_705_v198_: _dafny.MultiSet
            d_705_v198_ = _dafny.MultiSet([default__.fm44(d_447_v20_, globalState)])
            d_706_v199_: _dafny.Seq
            d_706_v199_ = _dafny.SeqWithoutIsStrInference([D13_DC32(d_419_v0_, d_447_v20_, (d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))], d_646_v160_)])
            d_707_v200_: _dafny.Array
            nw115_ = _dafny.Array(None, 20)
            nw115_[int(0)] = (d_447_v20_ if d_447_v20_ else d_447_v20_)
            nw115_[int(1)] = True
            nw115_[int(2)] = (d_703_v196_)[default__.safeIndex((0) - ((d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]), len(d_703_v196_))]
            nw115_[int(3)] = not(d_447_v20_)
            nw115_[int(4)] = d_447_v20_
            nw115_[int(5)] = False
            nw115_[int(6)] = (d_704_v197_).isdisjoint(_dafny.Set({d_447_v20_, d_447_v20_}))
            nw115_[int(7)] = False
            nw115_[int(8)] = d_447_v20_
            nw115_[int(9)] = d_447_v20_
            nw115_[int(10)] = (_dafny.MultiSet([((d_696_v189_)[d_447_v20_] if (d_447_v20_) in (d_696_v189_) else p0), len(d_706_v199_)])).ispropersubset(d_705_v198_)
            nw115_[int(11)] = d_447_v20_
            nw115_[int(12)] = (p0) >= ((d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))])
            nw115_[int(13)] = ((d_701_v194_)[(d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]] if ((d_646_v160_)[default__.safeIndex(653, (d_646_v160_).length(0))]) in (d_701_v194_) else d_447_v20_)
            nw115_[int(14)] = (self).fm42(d_447_v20_, globalState)
            nw115_[int(15)] = (d_447_v20_ if d_447_v20_ else d_447_v20_)
            nw115_[int(16)] = (d_447_v20_ if d_447_v20_ else d_447_v20_)
            nw115_[int(17)] = not(d_447_v20_)
            nw115_[int(18)] = d_447_v20_
            nw115_[int(19)] = (d_643_v157_) in (d_419_v0_)
            d_707_v200_ = nw115_
            index121_ = default__.safeIndex(166, (d_707_v200_).length(0))
            (d_707_v200_)[index121_] = d_447_v20_
            index122_ = default__.safeIndex(166, (d_707_v200_).length(0))
            (d_707_v200_)[index122_] = (p0) == (p0)
        d_708_v201_: D4
        d_708_v201_ = D4_DC10(p0, p0, p0)
        pat_let_tv10_ = d_447_v20_
        pat_let_tv11_ = d_448_v21_
        pat_let_tv12_ = d_447_v20_
        pat_let_tv13_ = d_448_v21_
        pat_let_tv14_ = d_447_v20_
        pat_let_tv15_ = d_447_v20_
        pat_let_tv16_ = d_447_v20_
        pat_let_tv17_ = d_447_v20_
        pat_let_tv18_ = d_447_v20_
        pat_let_tv19_ = p0
        pat_let_tv20_ = d_447_v20_
        pat_let_tv21_ = d_447_v20_
        pat_let_tv22_ = p0
        pat_let_tv23_ = d_447_v20_
        pat_let_tv24_ = p0
        pat_let_tv25_ = p0
        def lambda30_(source9_):
            if source9_.is_DC9:
                d_709___mcc_h25_ = source9_.cf14
                d_710___mcc_h26_ = source9_.cf15
                d_711___mcc_h27_ = source9_.cf16
                d_712___mcc_h28_ = source9_.cf17
                d_713___mcc_h29_ = source9_.cf18
                d_714_cf18_ = d_713___mcc_h29_
                d_715_cf17_ = d_712___mcc_h28_
                d_716_cf16_ = d_711___mcc_h27_
                d_717_cf15_ = d_710___mcc_h26_
                d_718_cf14_ = d_709___mcc_h25_
                return _dafny.Map({pat_let_tv10_: _dafny.MultiSet(pat_let_tv11_)})
            elif source9_.is_DC10:
                d_719___mcc_h30_ = source9_.cf19
                d_720___mcc_h31_ = source9_.cf20
                d_721___mcc_h32_ = source9_.cf21
                d_722_cf21_ = d_721___mcc_h32_
                d_723_cf20_ = d_720___mcc_h31_
                d_724_cf19_ = d_719___mcc_h30_
                return _dafny.Map({pat_let_tv12_: _dafny.MultiSet((D2_DC5(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_725_i24_ in range(default__.abs(842))])), pat_let_tv13_, pat_let_tv14_)).cf9)})
            elif source9_.is_DC11:
                d_726___mcc_h33_ = source9_.cf22
                d_727___mcc_h34_ = source9_.cf23
                d_728___mcc_h35_ = source9_.cf24
                d_729___mcc_h36_ = source9_.cf25
                d_730___mcc_h37_ = source9_.cf26
                d_731_cf26_ = d_730___mcc_h37_
                d_732_cf25_ = d_729___mcc_h36_
                d_733_cf24_ = d_728___mcc_h35_
                d_734_cf23_ = d_727___mcc_h34_
                d_735_cf22_ = d_726___mcc_h33_
                return _dafny.Map({d_731_cf26_: _dafny.MultiSet([d_735_cf22_])})
            elif source9_.is_DC8:
                d_736___mcc_h38_ = source9_.cf13
                d_737_cf13_ = d_736___mcc_h38_
                return (_dafny.Map({pat_let_tv15_: _dafny.MultiSet([len(_dafny.Map({_dafny.Map({len(_dafny.Set({True})): pat_let_tv16_}): pat_let_tv17_}))])})) | (_dafny.Map({pat_let_tv18_: _dafny.MultiSet([pat_let_tv19_, len(_dafny.SeqWithoutIsStrInference([pat_let_tv20_, pat_let_tv21_]))])}))
            elif True:
                d_738___mcc_h39_ = source9_.cf27
                d_739_cf27_ = d_738___mcc_h39_
                return (_dafny.Map({False: _dafny.MultiSet([pat_let_tv22_])})) | (_dafny.Map({pat_let_tv23_: _dafny.MultiSet([pat_let_tv24_])}))

        def iife68_(_pat_let14_0):
            def iife69_(d_740_dt__update__tmp_h4_):
                def iife70_(_pat_let15_0):
                    def iife71_(d_741_dt__update_hcf21_h0_):
                        return D4_DC10((d_740_dt__update__tmp_h4_).cf19, (d_740_dt__update__tmp_h4_).cf20, d_741_dt__update_hcf21_h0_)
                    return iife71_(_pat_let15_0)
                return iife70_(pat_let_tv25_)
            return iife69_(_pat_let14_0)
        r0 = lambda30_((iife68_(d_708_v201_) if d_447_v20_ else d_708_v201_))
        return r0

    def m14(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r0 = default__.fm44(p1, globalState)
        d_742_v0_: _dafny.Seq
        d_742_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnm"))
        d_743_v1_: _dafny.Map
        d_743_v1_ = _dafny.Map({p3: True})
        hi3_ = (293) + (len(d_743_v1_))
        for d_744_i0_ in range((0) - (len(d_742_v0_)), hi3_):
            d_745_v2_: _dafny.Array
            nw116_ = _dafny.Array(False, 8)
            d_745_v2_ = nw116_
            index123_ = default__.safeIndex(332, (d_745_v2_).length(0))
            (d_745_v2_)[index123_] = p1
            index124_ = default__.safeIndex(332, (d_745_v2_).length(0))
            (d_745_v2_)[index124_] = (p1) or ((False if p0 else p1))
            d_746_v3_: _dafny.Set
            d_746_v3_ = _dafny.Set({p0})
            d_747_v4_: _dafny.Array
            nw117_ = _dafny.Array(None, 18)
            nw117_[int(0)] = d_746_v3_
            nw117_[int(1)] = d_746_v3_
            nw117_[int(2)] = d_746_v3_
            nw117_[int(3)] = d_746_v3_
            nw117_[int(4)] = d_746_v3_
            nw117_[int(5)] = d_746_v3_
            nw117_[int(6)] = d_746_v3_
            nw117_[int(7)] = d_746_v3_
            nw117_[int(8)] = _dafny.Set({(d_745_v2_)[default__.safeIndex(332, (d_745_v2_).length(0))], p0})
            nw117_[int(9)] = _dafny.Set({p0})
            nw117_[int(10)] = d_746_v3_
            nw117_[int(11)] = d_746_v3_
            nw117_[int(12)] = d_746_v3_
            nw117_[int(13)] = _dafny.Set({(d_745_v2_)[default__.safeIndex(332, (d_745_v2_).length(0))]})
            nw117_[int(14)] = d_746_v3_
            nw117_[int(15)] = d_746_v3_
            nw117_[int(16)] = d_746_v3_
            nw117_[int(17)] = d_746_v3_
            d_747_v4_ = nw117_
            d_748_v5_: D0
            d_748_v5_ = D0_DC0(d_747_v4_, 2)
            d_749_v6_: _dafny.Seq
            d_749_v6_ = _dafny.SeqWithoutIsStrInference([d_748_v5_, d_748_v5_])
            d_750_v7_: D4
            d_750_v7_ = D4_DC8(d_749_v6_)
            d_750_v7_ = d_750_v7_
            d_751_v8_: str
            d_751_v8_ = _dafny.CodePoint('n')
            d_752_v9_: _dafny.Map
            d_752_v9_ = _dafny.Map({d_751_v8_: p3})
            d_753_v11_: _dafny.Map
            def iife72_():
                coll40_ = _dafny.Map()
                compr_40_: str
                for compr_40_ in (d_742_v0_).Elements:
                    d_754_v10_: str = compr_40_
                    if (d_754_v10_) in (d_742_v0_):
                        coll40_[d_754_v10_] = p3
                return _dafny.Map(coll40_)
            d_753_v11_ = (d_752_v9_ if p1 else iife72_()
            )
            d_753_v11_ = default__.fm54((d_745_v2_)[default__.safeIndex(332, (d_745_v2_).length(0))], _dafny.SeqWithoutIsStrInference([p1, (d_745_v2_)[default__.safeIndex(332, (d_745_v2_).length(0))]]), globalState)
            r0 = (0) - (p3)
        d_755_v12_: bool
        d_755_v12_ = False
        d_755_v12_ = p0
        hi4_ = (p3 if p1 else p3)
        for d_756_i1_ in range(p3, hi4_):
            d_757_v13_: str
            d_757_v13_ = _dafny.CodePoint('g')
            d_757_v13_ = (d_757_v13_ if d_755_v12_ else d_757_v13_)
            d_755_v12_ = (d_742_v0_) <= (d_742_v0_)
            d_742_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")) if p1 else d_742_v0_)
            d_758_v14_: _dafny.Array
            nw118_ = _dafny.Array(int(0), 17)
            d_758_v14_ = nw118_
            index125_ = default__.safeIndex(783, (d_758_v14_).length(0))
            (d_758_v14_)[index125_] = p3
            d_759_v15_: _dafny.Seq
            d_759_v15_ = _dafny.SeqWithoutIsStrInference([False])
            index126_ = default__.safeIndex(783, (d_758_v14_).length(0))
            rhs79_ = (len((d_743_v1_) | (d_743_v1_))) - (d_756_i1_)
            rhs80_ = d_756_i1_
            rhs81_ = (d_759_v15_) + (_dafny.SeqWithoutIsStrInference([(D14_DC35(p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sia")), d_756_i1_, d_755_v12_, p0)).cf67, p0, p1, d_755_v12_]))
            lhs58_ = d_758_v14_
            lhs59_ = default__.safeIndex(783, (d_758_v14_).length(0))
            lhs58_[lhs59_] = rhs79_
            r0 = rhs80_
            d_759_v15_ = rhs81_
        d_760_v16_: _dafny.Set
        d_760_v16_ = _dafny.Set({p3, p3})
        d_761_v17_: D6
        d_761_v17_ = D6_DC15(len(d_760_v16_), p3, p0)
        d_762_v18_: D6
        d_762_v18_ = D6_DC16(d_761_v17_)
        d_763_v19_: _dafny.MultiSet
        d_763_v19_ = _dafny.MultiSet([d_762_v18_, D6_DC16((D6_DC16(d_761_v17_)).cf33)])
        d_755_v12_ = not(((p3 if p0 else p3)) not in ((_dafny.SeqWithoutIsStrInference([p3 for d_764_i2_ in range(default__.abs(204))]) if not(p1) else _dafny.SeqWithoutIsStrInference([(d_763_v19_).cardinality, p3, p3, p3]))))
        d_742_v0_ = d_742_v0_
        r0 = p3
        return r0


class C3(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, p2, p3, globalState):
        return not(not(((622) - ((0) - ((0) - (-779)))) in ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hi"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('f')]))]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqgtkbdg"))), len(_dafny.Set({-285, len(_dafny.Map({_dafny.CodePoint('w'): _dafny.CodePoint('k')})), len(_dafny.Map({-394: 772}))}))])))))

    def m1(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        d_765_v0_: bool
        d_765_v0_ = False
        if d_765_v0_:
            d_766_v1_: bool
            d_767_v2_: int
            d_768_v3_: int
            d_769_v4_: bool
            out21_: bool
            out22_: int
            out23_: int
            out24_: bool
            out21_, out22_, out23_, out24_ = (self).m12(not (d_765_v0_) or (d_765_v0_), globalState)
            d_766_v1_ = out21_
            d_767_v2_ = out22_
            d_768_v3_ = out23_
            d_769_v4_ = out24_
            if d_765_v0_:
                d_770_v5_: _dafny.Array
                nw119_ = _dafny.Array(_dafny.Map({}), 26)
                d_770_v5_ = nw119_
                d_771_v6_: str
                d_771_v6_ = _dafny.CodePoint('g')
                d_772_v7_: D0
                d_772_v7_ = D0_DC1(p0)
                d_773_v8_: _dafny.Seq
                d_773_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnnc"))
                d_774_v9_: _dafny.Map
                d_774_v9_ = _dafny.Map({(self).fm5(p0, False, _dafny.Map({d_771_v6_: d_772_v7_}), d_769_v4_, globalState): d_773_v8_})
                d_775_v10_: _dafny.Set
                d_775_v10_ = _dafny.Set({p0, d_768_v3_})
                d_776_v11_: _dafny.MultiSet
                d_776_v11_ = _dafny.MultiSet([p0])
                d_777_v12_: _dafny.Seq
                d_777_v12_ = _dafny.SeqWithoutIsStrInference([len(d_774_v9_), len(d_775_v10_), (d_776_v11_).cardinality, default__.fm32(p0, globalState), d_768_v3_])
                d_778_v13_: T1
                nw120_ = C1()
                nw120_.ctor__((0) - (-771), d_770_v5_, d_777_v12_, d_772_v7_)
                d_778_v13_ = nw120_
                d_779_v14_: _dafny.Seq
                d_779_v14_ = _dafny.SeqWithoutIsStrInference([d_778_v13_])
                d_779_v14_ = d_779_v14_
                d_780_v15_: _dafny.Array
                def lambda31_(d_781_v0_):
                    def lambda32_(d_782_i0_):
                        return default__.safeModuloInt(d_782_i0_, len(_dafny.SeqWithoutIsStrInference([not(d_781_v0_)])))

                    return lambda32_

                init17_ = lambda31_(d_765_v0_)
                nw121_ = _dafny.Array(None, 9)
                for i0_17_ in range(nw121_.length(0)):
                    nw121_[i0_17_] = init17_(i0_17_)
                d_780_v15_ = nw121_
                index127_ = default__.safeIndex(189, (d_780_v15_).length(0))
                (d_780_v15_)[index127_] = d_767_v2_
                index128_ = default__.safeIndex(189, (d_780_v15_).length(0))
                (d_780_v15_)[index128_] = d_767_v2_
                d_783_v16_: _dafny.Set
                d_783_v16_ = _dafny.Set({d_769_v4_})
                d_784_v17_: D12
                d_784_v17_ = D12_DC27(d_783_v16_)
                d_785_v18_: _dafny.Map
                d_785_v18_ = _dafny.Map({d_769_v4_: d_783_v16_})
                d_784_v17_ = D12_DC27(((d_785_v18_)[d_765_v0_] if (d_765_v0_) in (d_785_v18_) else _dafny.Set({d_765_v0_})))
                d_780_v15_ = d_780_v15_
                d_786_v19_: D6
                d_786_v19_ = D6_DC14(d_773_v8_)
                d_787_v20_: _dafny.MultiSet
                d_787_v20_ = _dafny.MultiSet([d_766_v1_])
                d_788_v21_: _dafny.Map
                d_788_v21_ = _dafny.Map({d_768_v3_: ((d_787_v20_)[d_769_v4_] if (d_769_v4_) in (d_787_v20_) else 262)})
                d_789_v22_: _dafny.Array
                nw122_ = _dafny.Array(None, 12)
                nw122_[int(0)] = d_786_v19_
                nw122_[int(1)] = d_786_v19_
                def iife73_(_pat_let16_0):
                    def iife74_(d_791_dt__update__tmp_h0_):
                        def iife75_(_pat_let17_0):
                            def iife76_(d_793_dt__update_hcf29_h0_):
                                return D6_DC14(d_793_dt__update_hcf29_h0_)
                            return iife76_(_pat_let17_0)
                        return iife75_(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_792_i2_ in range(default__.abs(287))]))
                    return iife74_(_pat_let16_0)
                nw122_[int(2)] = iife73_(default__.fm33(((d_788_v21_)[d_768_v3_] if (d_768_v3_) in (d_788_v21_) else p0), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_790_i1_ in range(default__.abs(911))]), globalState))
                nw122_[int(3)] = d_786_v19_
                nw122_[int(4)] = d_786_v19_
                nw122_[int(5)] = d_786_v19_
                nw122_[int(6)] = d_786_v19_
                nw122_[int(7)] = d_786_v19_
                nw122_[int(8)] = D6_DC14(d_773_v8_)
                nw122_[int(9)] = d_786_v19_
                nw122_[int(10)] = D6_DC14(d_773_v8_)
                nw122_[int(11)] = default__.fm33((d_780_v15_)[default__.safeIndex(189, (d_780_v15_).length(0))], _dafny.SeqWithoutIsStrInference([d_771_v6_ for d_794_i3_ in range(default__.abs(190))]), globalState)
                d_789_v22_ = nw122_
                pat_let_tv26_ = d_773_v8_
                index129_ = default__.safeIndex(402, (d_789_v22_).length(0))
                def iife77_(_pat_let18_0):
                    def iife78_(d_795_dt__update__tmp_h1_):
                        def iife79_(_pat_let19_0):
                            def iife80_(d_796_dt__update_hcf29_h1_):
                                return D6_DC14(d_796_dt__update_hcf29_h1_)
                            return iife80_(_pat_let19_0)
                        return iife79_(pat_let_tv26_)
                    return iife78_(_pat_let18_0)
                (d_789_v22_)[index129_] = iife77_(d_786_v19_)
                index130_ = default__.safeIndex(402, (d_789_v22_).length(0))
                (d_789_v22_)[index130_] = d_786_v19_
            elif True:
                d_797_v23_: _dafny.Array
                nw123_ = _dafny.Array(_dafny.Map({}), 4)
                d_797_v23_ = nw123_
                d_797_v23_ = d_797_v23_
                d_768_v3_ = (0) - (p0)
                r1 = d_769_v4_
                d_798_v24_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_798_v24_ = nw124_
                d_799_v25_: _dafny.Seq
                d_799_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvm"))
                index131_ = default__.safeIndex(93, (d_798_v24_).length(0))
                (d_798_v24_)[index131_] = d_799_v25_
                index132_ = default__.safeIndex(93, (d_798_v24_).length(0))
                (d_798_v24_)[index132_] = d_799_v25_
                d_800_v26_: _dafny.Map
                d_800_v26_ = _dafny.Map({(d_769_v4_) == (d_765_v0_): d_765_v0_})
                d_800_v26_ = (d_800_v26_) | (d_800_v26_)
            d_801_v27_: _dafny.Seq
            d_801_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opafnutak"))
            d_802_v28_: _dafny.Seq
            d_802_v28_ = _dafny.SeqWithoutIsStrInference([d_801_v27_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_803_i4_ in range(default__.abs(620))]), d_801_v27_, d_801_v27_, d_801_v27_])
            d_804_v29_: C0
            nw125_ = C0()
            nw125_.ctor__((d_802_v28_)[default__.safeIndex(171, len(d_802_v28_))])
            d_804_v29_ = nw125_
            if True:
                d_805_v30_: _dafny.Seq
                d_805_v30_ = _dafny.SeqWithoutIsStrInference([d_769_v4_])
                d_806_v31_: _dafny.Seq
                d_806_v31_ = _dafny.SeqWithoutIsStrInference([d_805_v30_, d_805_v30_])
                d_807_v32_: _dafny.Map
                d_807_v32_ = _dafny.Map({p0: d_768_v3_})
                d_808_v33_: _dafny.Seq
                d_808_v33_ = _dafny.SeqWithoutIsStrInference([len(d_801_v27_), 553, ((d_807_v32_)[-361] if (-361) in (d_807_v32_) else p0), (0) - (d_768_v3_), d_767_v2_])
                d_809_v34_: _dafny.Set
                d_809_v34_ = _dafny.Set({p0})
                d_810_v35_: _dafny.Array
                nw126_ = _dafny.Array(None, 28)
                nw126_[int(0)] = d_768_v3_
                nw126_[int(1)] = d_767_v2_
                nw126_[int(2)] = d_768_v3_
                nw126_[int(3)] = d_767_v2_
                nw126_[int(4)] = d_767_v2_
                nw126_[int(5)] = p0
                nw126_[int(6)] = p0
                nw126_[int(7)] = p0
                nw126_[int(8)] = d_767_v2_
                nw126_[int(9)] = p0
                nw126_[int(10)] = p0
                nw126_[int(11)] = d_767_v2_
                nw126_[int(12)] = d_767_v2_
                nw126_[int(13)] = 356
                nw126_[int(14)] = d_767_v2_
                nw126_[int(15)] = d_768_v3_
                nw126_[int(16)] = -723
                nw126_[int(17)] = (_dafny.MultiSet((d_806_v31_)[default__.safeIndex(d_768_v3_, len(d_806_v31_))])).cardinality
                nw126_[int(18)] = (d_808_v33_)[default__.safeIndex(d_768_v3_, len(d_808_v33_))]
                nw126_[int(19)] = d_768_v3_
                nw126_[int(20)] = p0
                nw126_[int(21)] = d_768_v3_
                nw126_[int(22)] = d_767_v2_
                nw126_[int(23)] = len(d_809_v34_)
                nw126_[int(24)] = d_768_v3_
                nw126_[int(25)] = p0
                nw126_[int(26)] = d_767_v2_
                nw126_[int(27)] = p0
                d_810_v35_ = nw126_
                d_811_v36_: _dafny.Map
                d_811_v36_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): d_810_v35_})
                d_812_v37_: D12
                d_812_v37_ = D12_DC28(d_765_v0_, d_767_v2_)
                d_813_v39_: _dafny.Array
                def lambda33_(d_814_p0_):
                    def lambda34_(d_815_i5_):
                        def iife81_():
                            coll41_ = _dafny.Map()
                            compr_41_: str
                            for compr_41_ in (_dafny.Set({_dafny.CodePoint('s')})).Elements:
                                d_816_v38_: str = compr_41_
                                if (d_816_v38_) in (_dafny.Set({_dafny.CodePoint('s')})):
                                    coll41_[d_816_v38_] = d_814_p0_
                            return _dafny.Map(coll41_)
                        return iife81_()
                        

                    return lambda34_

                init18_ = lambda33_(p0)
                nw127_ = _dafny.Array(None, 28)
                for i0_18_ in range(nw127_.length(0)):
                    nw127_[i0_18_] = init18_(i0_18_)
                d_813_v39_ = nw127_
                d_817_v40_: D0
                d_817_v40_ = D0_DC1(p0)
                d_818_v41_: C1
                nw128_ = C1()
                nw128_.ctor__(p0, d_813_v39_, d_808_v33_, d_817_v40_)
                d_818_v41_ = nw128_
                d_819_v42_: _dafny.Map
                d_819_v42_ = _dafny.Map({d_818_v41_: False})
                d_820_v43_: _dafny.Set
                d_820_v43_ = _dafny.Set({not(d_766_v1_)})
                d_821_v44_: _dafny.Array
                nw129_ = _dafny.Array(None, 20)
                nw129_[int(0)] = d_769_v4_
                nw129_[int(1)] = (d_768_v3_) < (d_768_v3_)
                nw129_[int(2)] = d_766_v1_
                nw129_[int(3)] = (d_767_v2_) < (len(d_811_v36_))
                nw129_[int(4)] = False
                nw129_[int(5)] = not(d_765_v0_)
                nw129_[int(6)] = not(d_766_v1_)
                nw129_[int(7)] = d_766_v1_
                nw129_[int(8)] = ((d_812_v37_).cf46 if d_769_v4_ else d_766_v1_)
                nw129_[int(9)] = d_766_v1_
                nw129_[int(10)] = (d_819_v42_) not in (_dafny.Map({d_819_v42_: d_820_v43_}))
                nw129_[int(11)] = d_766_v1_
                nw129_[int(12)] = d_769_v4_
                nw129_[int(13)] = not (False) or (d_769_v4_)
                nw129_[int(14)] = d_766_v1_
                nw129_[int(15)] = (d_768_v3_) >= (d_818_v41_.f9)
                nw129_[int(16)] = False
                nw129_[int(17)] = (d_809_v34_).issubset(d_809_v34_)
                nw129_[int(18)] = not((d_766_v1_ if d_769_v4_ else d_765_v0_))
                nw129_[int(19)] = d_766_v1_
                d_821_v44_ = nw129_
                index133_ = default__.safeIndex(637, (d_821_v44_).length(0))
                (d_821_v44_)[index133_] = d_769_v4_
                index134_ = default__.safeIndex(637, (d_821_v44_).length(0))
                (d_821_v44_)[index134_] = d_765_v0_
                d_768_v3_ = d_767_v2_
                d_822_v45_: _dafny.Map
                d_822_v45_ = _dafny.Map({d_765_v0_: (d_818_v41_.f9) + ((default__.fm34(d_766_v1_, d_765_v0_, default__.fm30(d_818_v41_.f9, (d_821_v44_)[default__.safeIndex(637, (d_821_v44_).length(0))], (d_821_v44_)[default__.safeIndex(637, (d_821_v44_).length(0))], globalState), 369, globalState)).cf47)})
                d_768_v3_ = len(d_822_v45_)
                (d_818_v41_).f9 = len(d_804_v29_.f8)
                d_823_v46_: C0
                nw130_ = C0()
                nw130_.ctor__(_dafny.SeqWithoutIsStrInference([(D11_DC25(_dafny.CodePoint('w'), d_767_v2_, (d_805_v30_)[default__.safeIndex(593, len(d_805_v30_))])).cf41 for d_824_i6_ in range(default__.abs(-8))]))
                d_823_v46_ = nw130_
            elif True:
                d_825_v47_: _dafny.MultiSet
                d_825_v47_ = _dafny.MultiSet([d_766_v1_])
                d_826_v48_: _dafny.MultiSet
                d_826_v48_ = _dafny.MultiSet([(d_825_v47_).cardinality])
                d_827_v49_: _dafny.Array
                def lambda35_(d_828_v1_):
                    def lambda36_(d_829_i7_):
                        return d_828_v1_

                    return lambda36_

                init19_ = lambda35_(d_766_v1_)
                nw131_ = _dafny.Array(None, 16)
                for i0_19_ in range(nw131_.length(0)):
                    nw131_[i0_19_] = init19_(i0_19_)
                d_827_v49_ = nw131_
                d_830_v50_: D12
                d_830_v50_ = D12_DC28(d_766_v1_, 524)
                d_831_v51_: str
                d_831_v51_ = _dafny.CodePoint('y')
                d_832_v52_: _dafny.Map
                d_832_v52_ = _dafny.Map({d_765_v0_: _dafny.CodePoint('c')})
                d_833_v53_: D0
                d_833_v53_ = D0_DC1(len(d_832_v52_))
                d_834_v54_: _dafny.Map
                d_834_v54_ = _dafny.Map({d_831_v51_: d_833_v53_})
                d_835_v55_: _dafny.Map
                d_835_v55_ = _dafny.Map({d_766_v1_: default__.fm30((0) - ((0) - (p0)), (self).fm5(d_767_v2_, d_769_v4_, d_834_v54_, d_765_v0_, globalState), d_769_v4_, globalState)})
                d_836_v56_: _dafny.Map
                d_836_v56_ = _dafny.Map({(0) - (d_767_v2_): d_766_v1_})
                d_837_v57_: _dafny.Map
                d_837_v57_ = _dafny.Map({d_765_v0_: len(_dafny.Map({d_768_v3_: d_766_v1_}))})
                d_838_v58_: _dafny.Array
                nw132_ = _dafny.Array(None, 22)
                nw132_[int(0)] = d_767_v2_
                nw132_[int(1)] = d_767_v2_
                nw132_[int(2)] = d_768_v3_
                nw132_[int(3)] = d_767_v2_
                nw132_[int(4)] = (default__.fm32(p0, globalState)) + (p0)
                nw132_[int(5)] = default__.safeDivisionInt(d_767_v2_, (0) - (d_767_v2_))
                nw132_[int(6)] = d_767_v2_
                nw132_[int(7)] = (0) - (p0)
                nw132_[int(8)] = d_768_v3_
                nw132_[int(9)] = (d_826_v48_).cardinality
                nw132_[int(10)] = p0
                nw132_[int(11)] = (_dafny.MultiSet([d_827_v49_, d_827_v49_, d_827_v49_])).cardinality
                nw132_[int(12)] = p0
                nw132_[int(13)] = (d_830_v50_).cf47
                nw132_[int(14)] = len(d_835_v55_)
                nw132_[int(15)] = (d_767_v2_) + (len(d_836_v56_))
                nw132_[int(16)] = p0
                nw132_[int(17)] = d_767_v2_
                nw132_[int(18)] = len(d_837_v57_)
                nw132_[int(19)] = p0
                nw132_[int(20)] = (0) - (p0)
                nw132_[int(21)] = (d_767_v2_) - (d_767_v2_)
                d_838_v58_ = nw132_
                index135_ = default__.safeIndex(289, (d_838_v58_).length(0))
                (d_838_v58_)[index135_] = (d_768_v3_ if not(d_766_v1_) else d_767_v2_)
                index136_ = default__.safeIndex(289, (d_838_v58_).length(0))
                (d_838_v58_)[index136_] = (default__.fm32(174, globalState)) + (p0)
                d_839_v59_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.MultiSet({}), 8)
                d_839_v59_ = nw133_
                d_840_v60_: _dafny.MultiSet
                d_840_v60_ = _dafny.MultiSet([d_831_v51_])
                rhs82_ = p0
                rhs83_ = d_839_v59_
                rhs84_ = d_840_v60_
                d_768_v3_ = rhs82_
                d_839_v59_ = rhs83_
                d_840_v60_ = rhs84_
                d_841_v61_: _dafny.Seq
                d_841_v61_ = _dafny.SeqWithoutIsStrInference([d_766_v1_, d_769_v4_, d_765_v0_])
                d_841_v61_ = ((d_841_v61_ if d_766_v1_ else d_841_v61_)).set(default__.safeIndex(p0, len((d_841_v61_ if d_766_v1_ else d_841_v61_))), (len(d_836_v56_)) >= ((d_838_v58_)[default__.safeIndex(289, (d_838_v58_).length(0))]))
                d_842_v62_: _dafny.Map
                d_842_v62_ = _dafny.Map({689: d_831_v51_})
                d_843_v63_: _dafny.Map
                d_843_v63_ = _dafny.Map({d_842_v62_: d_767_v2_})
                d_843_v63_ = (d_843_v63_).set(d_842_v62_, (0) - (d_767_v2_))
                d_801_v27_ = d_801_v27_
            d_844_v64_: C0
            nw134_ = C0()
            nw134_.ctor__(d_801_v27_)
            d_844_v64_ = nw134_
        elif True:
            d_845_v65_: _dafny.Seq
            d_845_v65_ = _dafny.SeqWithoutIsStrInference([d_765_v0_])
            d_846_v66_: _dafny.MultiSet
            d_846_v66_ = _dafny.MultiSet([d_765_v0_, True, d_765_v0_, d_765_v0_, (d_845_v65_)[default__.safeIndex(p0, len(d_845_v65_))]])
            d_847_v67_: _dafny.Map
            d_847_v67_ = _dafny.Map({len(_dafny.Map({False: p0})): default__.fm35(d_765_v0_, False, globalState)})
            d_848_v69_: _dafny.Seq
            d_848_v69_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            def iife82_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in (d_848_v69_).Elements:
                    d_849_v68_: int = compr_42_
                    if (d_849_v68_) in (d_848_v69_):
                        coll42_[(d_849_v68_) - ((0) - (p0))] = d_848_v69_
                return _dafny.Map(coll42_)
            d_765_v0_ = (((d_846_v66_)[not(d_765_v0_)] if (not(d_765_v0_)) in (d_846_v66_) else p0)) <= (len((d_847_v67_) | (iife82_()
            )))
            d_850_v70_: _dafny.Seq
            d_850_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "is"))
            d_851_v71_: str
            d_851_v71_ = _dafny.CodePoint('v')
            d_852_v72_: _dafny.Seq
            d_852_v72_ = _dafny.SeqWithoutIsStrInference([(d_850_v70_)[default__.safeIndex(p0, len(d_850_v70_))], (d_850_v70_)[default__.safeIndex(p0, len(d_850_v70_))], d_851_v71_, d_851_v71_])
            d_853_v73_: _dafny.Set
            d_853_v73_ = _dafny.Set({d_765_v0_})
            d_854_v74_: D11
            d_854_v74_ = D11_DC25((d_850_v70_)[default__.safeIndex(len(d_853_v73_), len(d_850_v70_))], p0, d_765_v0_)
            d_855_v75_: _dafny.Map
            d_855_v75_ = _dafny.Map({d_765_v0_: p0})
            d_856_v76_: D12
            d_856_v76_ = D12_DC29(d_765_v0_, p0, False, (0) - ((0) - (len(d_855_v75_))), d_765_v0_)
            source10_ = (d_856_v76_ if (d_765_v0_) == ((d_854_v74_).cf43) else d_856_v76_)
            if source10_.is_DC28:
                d_857___mcc_h0_ = source10_.cf46
                d_858___mcc_h1_ = source10_.cf47
                d_859_cf47_ = d_858___mcc_h1_
                d_860_cf46_ = d_857___mcc_h0_
                d_861_v77_: _dafny.Map
                d_861_v77_ = _dafny.Map({(0) - (p0): d_851_v71_})
                d_862_v78_: _dafny.Array
                nw135_ = _dafny.Array(None, 22)
                nw135_[int(0)] = _dafny.CodePoint('c')
                nw135_[int(1)] = d_851_v71_
                nw135_[int(2)] = d_851_v71_
                nw135_[int(3)] = d_851_v71_
                nw135_[int(4)] = d_851_v71_
                nw135_[int(5)] = d_851_v71_
                nw135_[int(6)] = d_851_v71_
                nw135_[int(7)] = d_851_v71_
                nw135_[int(8)] = _dafny.CodePoint('i')
                nw135_[int(9)] = ((d_861_v77_)[509] if (509) in (d_861_v77_) else d_851_v71_)
                nw135_[int(10)] = d_851_v71_
                nw135_[int(11)] = d_851_v71_
                nw135_[int(12)] = d_851_v71_
                nw135_[int(13)] = d_851_v71_
                nw135_[int(14)] = d_851_v71_
                nw135_[int(15)] = d_851_v71_
                nw135_[int(16)] = _dafny.CodePoint('h')
                nw135_[int(17)] = _dafny.CodePoint('b')
                nw135_[int(18)] = default__.fm29(d_860_cf46_, globalState)
                nw135_[int(19)] = d_851_v71_
                nw135_[int(20)] = d_851_v71_
                nw135_[int(21)] = d_851_v71_
                d_862_v78_ = nw135_
                d_863_v79_: _dafny.Map
                d_863_v79_ = _dafny.Map({d_862_v78_: d_853_v73_})
                d_863_v79_ = d_863_v79_
                d_864_v80_: _dafny.Array
                nw136_ = _dafny.Array(_dafny.Map({}), 2)
                d_864_v80_ = nw136_
                d_865_v81_: _dafny.MultiSet
                d_865_v81_ = _dafny.MultiSet([-804])
                d_866_v82_: C0
                nw137_ = C0()
                nw137_.ctor__(d_852_v72_)
                d_866_v82_ = nw137_
                d_867_v83_: _dafny.Set
                d_867_v83_ = _dafny.Set({d_866_v82_, d_866_v82_, d_866_v82_, d_866_v82_})
                d_868_v84_: D4
                d_868_v84_ = D4_DC9(d_865_v81_, d_851_v71_, d_859_cf47_, d_867_v83_, 712)
                d_869_v85_: _dafny.Map
                d_869_v85_ = _dafny.Map({(_dafny.Map({True: D4_DC9(d_865_v81_, d_851_v71_, 207, d_867_v83_, p0)})).set(d_765_v0_, d_868_v84_): _dafny.MultiSet(d_848_v69_)})
                index137_ = default__.safeIndex(682, (d_864_v80_).length(0))
                (d_864_v80_)[index137_] = (d_869_v85_) | (d_869_v85_)
                index138_ = default__.safeIndex(682, (d_864_v80_).length(0))
                (d_864_v80_)[index138_] = d_869_v85_
                d_870_v86_: _dafny.Array
                nw138_ = _dafny.Array(_dafny.Set({}), 25)
                d_870_v86_ = nw138_
                d_871_v87_: D0
                d_871_v87_ = D0_DC0(d_870_v86_, d_859_cf47_)
                d_872_v88_: _dafny.Map
                d_872_v88_ = _dafny.Map({default__.fm32(p0, globalState): d_765_v0_})
                d_873_v89_: _dafny.Map
                d_873_v89_ = _dafny.Map({d_851_v71_: len(d_872_v88_)})
                d_874_v91_: _dafny.Array
                nw139_ = _dafny.Array(None, 17)
                nw139_[int(0)] = d_873_v89_
                nw139_[int(1)] = (d_873_v89_).set(d_851_v71_, len(d_852_v72_))
                nw139_[int(2)] = default__.fm36(d_851_v71_, (0) - (len(_dafny.Map({d_855_v75_: d_851_v71_}))), default__.fm32(len(_dafny.Map({d_765_v0_: d_859_cf47_})), globalState), (D12_DC27(d_853_v73_)).cf45, globalState)
                nw139_[int(3)] = d_873_v89_
                nw139_[int(4)] = d_873_v89_
                nw139_[int(5)] = d_873_v89_
                nw139_[int(6)] = d_873_v89_
                nw139_[int(7)] = d_873_v89_
                nw139_[int(8)] = d_873_v89_
                def iife83_():
                    coll43_ = _dafny.Map()
                    compr_43_: str
                    for compr_43_ in (d_850_v70_).Elements:
                        d_875_v90_: str = compr_43_
                        if (d_875_v90_) in (d_850_v70_):
                            coll43_[d_875_v90_] = d_859_cf47_
                    return _dafny.Map(coll43_)
                nw139_[int(9)] = iife83_()
                
                nw139_[int(10)] = d_873_v89_
                nw139_[int(11)] = _dafny.Map({d_851_v71_: 774})
                nw139_[int(12)] = d_873_v89_
                nw139_[int(13)] = d_873_v89_
                nw139_[int(14)] = _dafny.Map({d_851_v71_: p0})
                nw139_[int(15)] = d_873_v89_
                nw139_[int(16)] = (d_873_v89_).set(d_851_v71_, p0)
                d_874_v91_ = nw139_
                d_876_v92_: D0
                d_876_v92_ = D0_DC1(len(d_872_v88_))
                d_877_v93_: C1
                nw140_ = C1()
                nw140_.ctor__((d_871_v87_).cf1, d_874_v91_, _dafny.SeqWithoutIsStrInference([p0, (0) - (p0)]), d_876_v92_)
                d_877_v93_ = nw140_
                d_878_v94_: _dafny.Array
                nw141_ = _dafny.Array(None, 22)
                nw141_[int(0)] = d_765_v0_
                nw141_[int(1)] = d_765_v0_
                nw141_[int(2)] = d_860_cf46_
                nw141_[int(3)] = d_765_v0_
                nw141_[int(4)] = True
                nw141_[int(5)] = d_765_v0_
                nw141_[int(6)] = d_765_v0_
                nw141_[int(7)] = False
                nw141_[int(8)] = d_765_v0_
                nw141_[int(9)] = d_860_cf46_
                nw141_[int(10)] = d_860_cf46_
                nw141_[int(11)] = not(d_860_cf46_)
                nw141_[int(12)] = d_860_cf46_
                nw141_[int(13)] = d_860_cf46_
                nw141_[int(14)] = d_765_v0_
                nw141_[int(15)] = not(d_860_cf46_)
                nw141_[int(16)] = d_765_v0_
                nw141_[int(17)] = d_860_cf46_
                nw141_[int(18)] = d_765_v0_
                nw141_[int(19)] = d_765_v0_
                nw141_[int(20)] = False
                nw141_[int(21)] = d_765_v0_
                d_878_v94_ = nw141_
                d_879_v95_: _dafny.Map
                d_880_v96_: bool
                d_881_v97_: bool
                d_882_v98_: _dafny.Set
                out25_: _dafny.Map
                out26_: bool
                out27_: bool
                out28_: _dafny.Set
                out25_, out26_, out27_, out28_ = (self).m11(d_851_v71_, d_878_v94_, d_846_v66_, default__.safeDivisionInt(d_877_v93_.f9, d_877_v93_.f9), globalState)
                d_879_v95_ = out25_
                d_880_v96_ = out26_
                d_881_v97_ = out27_
                d_882_v98_ = out28_
            elif source10_.is_DC29:
                d_883___mcc_h2_ = source10_.cf48
                d_884___mcc_h3_ = source10_.cf49
                d_885___mcc_h4_ = source10_.cf50
                d_886___mcc_h5_ = source10_.cf51
                d_887___mcc_h6_ = source10_.cf52
                d_888_cf52_ = d_887___mcc_h6_
                d_889_cf51_ = d_886___mcc_h5_
                d_890_cf50_ = d_885___mcc_h4_
                d_891_cf49_ = d_884___mcc_h3_
                d_892_cf48_ = d_883___mcc_h2_
                d_893_v99_: _dafny.Array
                nw142_ = _dafny.Array(_dafny.Map({}), 26)
                d_893_v99_ = nw142_
                d_894_v100_: D0
                d_894_v100_ = D0_DC1(d_891_cf49_)
                d_895_v101_: C1
                nw143_ = C1()
                nw143_.ctor__(d_891_cf49_, d_893_v99_, d_848_v69_, d_894_v100_)
                d_895_v101_ = nw143_
                d_889_cf51_ = 852
                d_891_cf49_ = (501 if d_890_cf50_ else default__.safeDivisionInt(d_889_cf51_, len(_dafny.SeqWithoutIsStrInference([d_851_v71_ for d_896_i8_ in range(default__.abs(436))]))))
                d_897_v102_: C0
                nw144_ = C0()
                nw144_.ctor__(((d_850_v70_) + (d_852_v72_)).set(default__.safeIndex(p0, len((d_850_v70_) + (d_852_v72_))), _dafny.CodePoint('p')))
                d_897_v102_ = nw144_
            elif True:
                d_898___mcc_h7_ = source10_.cf45
                d_899_cf45_ = d_898___mcc_h7_
                r1 = (p0) < (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oggyllotw"))).set(default__.safeIndex(len(d_855_v75_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oggyllotw")))), d_851_v71_)))
                d_900_v103_: int
                d_900_v103_ = 345
                d_900_v103_ = p0
                d_851_v71_ = d_851_v71_
                d_900_v103_ = default__.safeModuloInt(len(_dafny.Set({d_765_v0_, d_765_v0_, d_765_v0_, d_765_v0_})), 469)
            if False:
                d_901_v104_: _dafny.Array
                nw145_ = _dafny.Array(int(0), 2)
                d_901_v104_ = nw145_
                d_902_v105_: _dafny.Map
                d_902_v105_ = _dafny.Map({p0: d_765_v0_})
                index139_ = default__.safeIndex(318, (d_901_v104_).length(0))
                (d_901_v104_)[index139_] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxuolybl"))), len(d_902_v105_))
                d_903_v106_: _dafny.MultiSet
                d_903_v106_ = _dafny.MultiSet([d_848_v69_])
                d_904_v107_: D7
                d_904_v107_ = D7_DC19((d_903_v106_).cardinality)
                index140_ = default__.safeIndex(318, (d_901_v104_).length(0))
                def iife84_():
                    coll44_ = _dafny.Set()
                    compr_44_: D7
                    for compr_44_ in (_dafny.SeqWithoutIsStrInference([D7_DC19((0) - (p0)), default__.fm37(True, globalState), d_904_v107_, d_904_v107_, d_904_v107_])).Elements:
                        d_905_v108_: D7 = compr_44_
                        if (d_905_v108_) in (_dafny.SeqWithoutIsStrInference([D7_DC19((0) - (p0)), default__.fm37(True, globalState), d_904_v107_, d_904_v107_, d_904_v107_])):
                            coll44_ = coll44_.union(_dafny.Set([d_905_v108_]))
                    return _dafny.Set(coll44_)
                (d_901_v104_)[index140_] = default__.safeDivisionInt((p0) + (len(iife84_()
                )), len((_dafny.Map({d_765_v0_: p0})) | (d_855_v75_)))
                d_906_v109_: _dafny.MultiSet
                d_906_v109_ = _dafny.MultiSet([p0])
                d_907_v110_: C0
                nw146_ = C0()
                nw146_.ctor__(d_852_v72_)
                d_907_v110_ = nw146_
                d_908_v111_: _dafny.Set
                d_908_v111_ = _dafny.Set({d_907_v110_})
                d_909_v112_: D4
                d_909_v112_ = D4_DC9(_dafny.MultiSet([p0]), (d_852_v72_)[default__.safeIndex(p0, len(d_852_v72_))], 748, d_908_v111_, 83)
                d_910_v113_: _dafny.Array
                nw147_ = _dafny.Array(None, 17)
                nw147_[int(0)] = ((d_902_v105_)[-785] if (-785) in (d_902_v105_) else d_765_v0_)
                nw147_[int(1)] = d_765_v0_
                nw147_[int(2)] = d_765_v0_
                nw147_[int(3)] = (d_851_v71_) not in (d_850_v70_)
                nw147_[int(4)] = (d_906_v109_).isdisjoint((d_909_v112_).cf14)
                nw147_[int(5)] = d_765_v0_
                nw147_[int(6)] = d_765_v0_
                nw147_[int(7)] = (d_765_v0_) == (not(False))
                nw147_[int(8)] = d_765_v0_
                nw147_[int(9)] = (d_765_v0_) and (d_765_v0_)
                nw147_[int(10)] = d_765_v0_
                nw147_[int(11)] = d_765_v0_
                nw147_[int(12)] = d_765_v0_
                nw147_[int(13)] = False
                nw147_[int(14)] = d_765_v0_
                nw147_[int(15)] = d_765_v0_
                nw147_[int(16)] = d_765_v0_
                d_910_v113_ = nw147_
                nw148_ = _dafny.Array(False, 11)
                d_910_v113_ = nw148_
                d_911_v114_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Seq({}), 8)
                d_911_v114_ = nw149_
                d_912_v115_: _dafny.Seq
                d_912_v115_ = _dafny.SeqWithoutIsStrInference([d_852_v72_, d_852_v72_, d_852_v72_])
                index141_ = default__.safeIndex(171, (d_911_v114_).length(0))
                (d_911_v114_)[index141_] = d_912_v115_
                index142_ = default__.safeIndex(171, (d_911_v114_).length(0))
                (d_911_v114_)[index142_] = d_912_v115_
                (d_907_v110_).f8 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmpvievok"))) + (d_852_v72_)) + (d_907_v110_.f8)
                d_913_v116_: _dafny.Map
                d_913_v116_ = _dafny.Map({len((d_907_v110_).fm15(d_848_v69_, (d_845_v65_)[default__.safeIndex(980, len(d_845_v65_))], globalState)): _dafny.Map({((d_855_v75_)[not(d_765_v0_)] if (not(d_765_v0_)) in (d_855_v75_) else p0): d_852_v72_})})
                d_914_v117_: _dafny.Map
                d_914_v117_ = _dafny.Map({(d_901_v104_)[default__.safeIndex(318, (d_901_v104_).length(0))]: d_907_v110_.f8})
                d_913_v116_ = (d_913_v116_).set(p0, d_914_v117_)
            elif True:
                d_850_v70_ = d_850_v70_
                d_765_v0_ = d_765_v0_
                d_846_v66_ = d_846_v66_
                r1 = (default__.fm35(d_765_v0_, d_765_v0_, globalState)) < (d_848_v69_)
                r0 = d_765_v0_
            d_915_v118_: int
            d_915_v118_ = 128
            d_915_v118_ = p0
            r1 = d_765_v0_
        d_916_v119_: int
        d_916_v119_ = -72
        d_916_v119_ = default__.safeModuloInt(d_916_v119_, p0)
        d_917_v120_: _dafny.Seq
        d_917_v120_ = _dafny.SeqWithoutIsStrInference([p0, d_916_v119_, d_916_v119_, p0, p0])
        d_918_v121_: _dafny.MultiSet
        d_918_v121_ = _dafny.MultiSet([p0, 348])
        d_919_v122_: C0
        nw150_ = C0()
        nw150_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jctntdft")))
        d_919_v122_ = nw150_
        d_920_v123_: _dafny.Set
        d_920_v123_ = _dafny.Set({d_919_v122_})
        hi5_ = d_916_v119_
        for d_921_i9_ in range((d_917_v120_)[default__.safeIndex((D4_DC9(d_918_v121_, _dafny.CodePoint('n'), p0, d_920_v123_, p0)).cf16, len(d_917_v120_))], hi5_):
            d_922_v124_: _dafny.Array
            nw151_ = _dafny.Array(int(0), 4)
            d_922_v124_ = nw151_
            index143_ = default__.safeIndex(200, (d_922_v124_).length(0))
            (d_922_v124_)[index143_] = d_916_v119_
            d_923_v125_: _dafny.Array
            nw152_ = _dafny.Array(_dafny.Seq({}), 17)
            d_923_v125_ = nw152_
            d_924_v126_: _dafny.Map
            d_924_v126_ = _dafny.Map({d_921_i9_: p0})
            d_925_v127_: D13
            d_925_v127_ = D13_DC30(d_923_v125_)
            pat_let_tv27_ = d_923_v125_
            index144_ = default__.safeIndex(200, (d_922_v124_).length(0))
            def iife85_(_pat_let20_0):
                def iife86_(d_926_dt__update__tmp_h2_):
                    def iife87_(_pat_let21_0):
                        def iife88_(d_927_dt__update_hcf53_h0_):
                            return D13_DC30(d_927_dt__update_hcf53_h0_)
                        return iife88_(_pat_let21_0)
                    return iife87_(pat_let_tv27_)
                return iife86_(_pat_let20_0)
            rhs85_ = ((d_924_v126_)[(0) - (d_921_i9_)] if ((0) - (d_921_i9_)) in (d_924_v126_) else ((d_918_v121_)[d_921_i9_] if (d_921_i9_) in (d_918_v121_) else (0) - (d_916_v119_)))
            rhs86_ = p0
            rhs87_ = (iife85_(d_925_v127_)).cf53
            lhs60_ = d_922_v124_
            lhs61_ = default__.safeIndex(200, (d_922_v124_).length(0))
            d_916_v119_ = rhs85_
            lhs60_[lhs61_] = rhs86_
            d_923_v125_ = rhs87_
            r1 = not((d_765_v0_ if not (d_765_v0_) or (d_765_v0_) else d_765_v0_))
            (d_919_v122_).f8 = d_919_v122_.f8
            d_928_v128_: _dafny.Seq
            d_928_v128_ = _dafny.SeqWithoutIsStrInference([d_917_v120_, d_917_v120_])
            if ((-986) != (792) if (len((d_928_v128_)[default__.safeIndex(d_916_v119_, len(d_928_v128_))])) <= (len(d_917_v120_)) else d_765_v0_):
                index145_ = default__.safeIndex(200, (d_922_v124_).length(0))
                (d_922_v124_)[index145_] = (default__.safeDivisionInt(d_916_v119_, (d_922_v124_)[default__.safeIndex(200, (d_922_v124_).length(0))])) * (len(_dafny.SeqWithoutIsStrInference([d_765_v0_, d_765_v0_])))
                d_929_v129_: _dafny.Map
                d_929_v129_ = _dafny.Map({d_765_v0_: d_922_v124_})
                d_929_v129_ = (d_929_v129_).set(d_765_v0_, d_922_v124_)
                (d_919_v122_).f8 = ((d_919_v122_).fm15(d_917_v120_, not(d_765_v0_), globalState)) + (d_919_v122_.f8)
                d_925_v127_ = d_925_v127_
                index146_ = default__.safeIndex(200, (d_922_v124_).length(0))
                rhs88_ = d_918_v121_
                rhs89_ = d_919_v122_.f8
                rhs90_ = d_921_i9_
                rhs91_ = ((d_917_v120_) + (d_917_v120_) if d_765_v0_ else d_917_v120_)
                lhs62_ = d_919_v122_
                lhs63_ = d_922_v124_
                lhs64_ = default__.safeIndex(200, (d_922_v124_).length(0))
                d_918_v121_ = rhs88_
                lhs62_.f8 = rhs89_
                lhs63_[lhs64_] = rhs90_
                d_917_v120_ = rhs91_
            elif True:
                (d_919_v122_).f8 = d_919_v122_.f8
                d_930_v130_: _dafny.Seq
                d_930_v130_ = _dafny.SeqWithoutIsStrInference([d_919_v122_.f8])
                d_931_v131_: _dafny.Seq
                d_931_v131_ = _dafny.SeqWithoutIsStrInference([not(d_765_v0_)])
                rhs92_ = d_765_v0_
                rhs93_ = d_765_v0_
                rhs94_ = (d_930_v130_)[default__.safeIndex(len((d_931_v131_).set(default__.safeIndex(d_921_i9_, len(d_931_v131_)), d_765_v0_)), len(d_930_v130_))]
                lhs65_ = d_919_v122_
                d_765_v0_ = rhs92_
                d_765_v0_ = rhs93_
                lhs65_.f8 = rhs94_
                d_932_v132_: _dafny.Array
                nw153_ = _dafny.Array(False, 9)
                d_932_v132_ = nw153_
                d_933_v133_: _dafny.Array
                nw154_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_933_v133_ = nw154_
                d_934_v134_: D2
                d_934_v134_ = D2_DC4(289, d_765_v0_, d_933_v133_)
                d_935_v135_: str
                d_935_v135_ = _dafny.CodePoint('c')
                d_936_v136_: _dafny.MultiSet
                d_936_v136_ = _dafny.MultiSet([d_765_v0_])
                d_937_v137_: D0
                d_937_v137_ = D0_DC1((d_936_v136_).cardinality)
                index147_ = default__.safeIndex(581, (d_932_v132_).length(0))
                (d_932_v132_)[index147_] = (self).fm5((d_934_v134_).cf5, d_765_v0_, _dafny.Map({d_935_v135_: d_937_v137_}), d_765_v0_, globalState)
                index148_ = default__.safeIndex(581, (d_932_v132_).length(0))
                (d_932_v132_)[index148_] = (((d_936_v136_).set(d_765_v0_, default__.abs((d_922_v124_)[default__.safeIndex(200, (d_922_v124_).length(0))]))).ispropersubset(d_936_v136_) if d_765_v0_ else (not(not(d_765_v0_)) if d_765_v0_ else d_765_v0_))
                index149_ = default__.safeIndex(200, (d_922_v124_).length(0))
                (d_922_v124_)[index149_] = (d_922_v124_)[default__.safeIndex(200, (d_922_v124_).length(0))]
                d_938_v138_: C0
                nw155_ = C0()
                nw155_.ctor__(d_919_v122_.f8)
                d_938_v138_ = nw155_
        d_939_v139_: C0
        nw156_ = C0()
        nw156_.ctor__(d_919_v122_.f8)
        d_939_v139_ = nw156_
        d_940_v140_: _dafny.Map
        d_940_v140_ = _dafny.Map({d_765_v0_: d_916_v119_})
        hi6_ = (d_916_v119_ if d_765_v0_ else len(d_940_v140_))
        for d_941_i10_ in range(len(d_940_v140_), hi6_):
            r1 = d_765_v0_
            r1 = True
            (d_939_v139_).f8 = d_939_v139_.f8
            d_942_v141_: _dafny.Array
            nw157_ = _dafny.Array(False, 25)
            d_942_v141_ = nw157_
            index150_ = default__.safeIndex(704, (d_942_v141_).length(0))
            (d_942_v141_)[index150_] = not(d_765_v0_)
            index151_ = default__.safeIndex(704, (d_942_v141_).length(0))
            (d_942_v141_)[index151_] = d_765_v0_
        d_943_v142_: _dafny.Array
        def lambda37_(d_944_v139_, d_945_p0_):
            def lambda38_(d_946_i11_):
                return (d_946_i11_) - (len(_dafny.Map({d_944_v139_.f8: d_945_p0_})))

            return lambda38_

        init20_ = lambda37_(d_939_v139_, p0)
        nw158_ = _dafny.Array(None, 19)
        for i0_20_ in range(nw158_.length(0)):
            nw158_[i0_20_] = init20_(i0_20_)
        d_943_v142_ = nw158_
        index152_ = default__.safeIndex(605, (d_943_v142_).length(0))
        (d_943_v142_)[index152_] = d_916_v119_
        index153_ = default__.safeIndex(605, (d_943_v142_).length(0))
        (d_943_v142_)[index153_] = d_916_v119_
        r0 = d_765_v0_
        r1 = d_765_v0_
        return r0, r1

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_947_v0_: bool
        d_947_v0_ = True
        d_948_v1_: _dafny.MultiSet
        d_948_v1_ = _dafny.MultiSet([d_947_v0_, d_947_v0_])
        d_949_v2_: _dafny.Map
        d_949_v2_ = _dafny.Map({p0: (D9_DC21(d_948_v1_)).cf37})
        d_949_v2_ = (d_949_v2_).set(p0, d_948_v1_)
        d_950_v3_: D7
        d_950_v3_ = D7_DC18()
        d_951_v4_: _dafny.Array
        nw159_ = _dafny.Array(_dafny.Set({}), 4)
        d_951_v4_ = nw159_
        d_952_v5_: _dafny.Array
        nw160_ = _dafny.Array(None, 19)
        nw160_[int(0)] = d_947_v0_
        nw160_[int(1)] = d_947_v0_
        nw160_[int(2)] = d_947_v0_
        nw160_[int(3)] = d_947_v0_
        nw160_[int(4)] = d_947_v0_
        nw160_[int(5)] = d_947_v0_
        nw160_[int(6)] = d_947_v0_
        nw160_[int(7)] = False
        nw160_[int(8)] = d_947_v0_
        nw160_[int(9)] = d_947_v0_
        nw160_[int(10)] = d_947_v0_
        nw160_[int(11)] = d_947_v0_
        nw160_[int(12)] = default__.fm30(p0, d_947_v0_, d_947_v0_, globalState)
        nw160_[int(13)] = True
        nw160_[int(14)] = d_947_v0_
        nw160_[int(15)] = d_947_v0_
        nw160_[int(16)] = d_947_v0_
        nw160_[int(17)] = d_947_v0_
        nw160_[int(18)] = d_947_v0_
        d_952_v5_ = nw160_
        d_953_v6_: _dafny.Array
        nw161_ = _dafny.Array(None, 10)
        nw161_[int(0)] = d_952_v5_
        nw161_[int(1)] = d_952_v5_
        nw161_[int(2)] = d_952_v5_
        nw161_[int(3)] = d_952_v5_
        nw161_[int(4)] = d_952_v5_
        nw161_[int(5)] = d_952_v5_
        nw161_[int(6)] = d_952_v5_
        nw161_[int(7)] = d_952_v5_
        nw161_[int(8)] = d_952_v5_
        nw161_[int(9)] = d_952_v5_
        d_953_v6_ = nw161_
        rhs95_ = d_950_v3_
        rhs96_ = d_951_v4_
        rhs97_ = d_953_v6_
        d_950_v3_ = rhs95_
        d_951_v4_ = rhs96_
        d_953_v6_ = rhs97_
        hi7_ = p0
        for d_954_i0_ in range(default__.safeDivisionInt(p0, -664), hi7_):
            d_955_v7_: int
            d_955_v7_ = 373
            d_956_v8_: _dafny.Seq
            d_956_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbkmmu"))
            d_957_v9_: str
            d_957_v9_ = _dafny.CodePoint('a')
            d_955_v7_ = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_958_i1_ in range(default__.abs(-424))])) + ((d_956_v8_).set(default__.safeIndex(len(d_956_v8_), len(d_956_v8_)), d_957_v9_)))
            d_957_v9_ = d_957_v9_
            d_959_v10_: C0
            nw162_ = C0()
            nw162_.ctor__((d_956_v8_) + (d_956_v8_))
            d_959_v10_ = nw162_
            d_960_v11_: _dafny.Array
            def lambda39_(d_961_v7_):
                def lambda40_(d_962_i2_):
                    return (d_962_i2_) + (d_961_v7_)

                return lambda40_

            init21_ = lambda39_(d_955_v7_)
            nw163_ = _dafny.Array(None, 16)
            for i0_21_ in range(nw163_.length(0)):
                nw163_[i0_21_] = init21_(i0_21_)
            d_960_v11_ = nw163_
            index154_ = default__.safeIndex(616, (d_960_v11_).length(0))
            (d_960_v11_)[index154_] = (len(_dafny.SeqWithoutIsStrInference([p0 for d_963_i3_ in range(default__.abs(625))]))) + (d_954_i0_)
            index155_ = default__.safeIndex(616, (d_960_v11_).length(0))
            (d_960_v11_)[index155_] = (d_955_v7_ if d_947_v0_ else default__.safeDivisionInt(d_954_i0_, d_954_i0_))
        d_964_v12_: _dafny.Array
        nw164_ = _dafny.Array(D2.default()(), 14)
        d_964_v12_ = nw164_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_964_v12_).length(0)):
            d_965_i4_: int = guard_loop_1_
            if (True) and (((0) <= (d_965_i4_)) and ((d_965_i4_) < ((d_964_v12_).length(0)))):
                (d_964_v12_)[(d_965_i4_)] = D2_DC5(p0, _dafny.SeqWithoutIsStrInference([p0, p0]), d_947_v0_)
        d_966_v13_: bool
        d_967_v14_: int
        d_968_v15_: int
        d_969_v16_: bool
        out29_: bool
        out30_: int
        out31_: int
        out32_: bool
        out29_, out30_, out31_, out32_ = (self).m12(d_947_v0_, globalState)
        d_966_v13_ = out29_
        d_967_v14_ = out30_
        d_968_v15_ = out31_
        d_969_v16_ = out32_
        d_970_v17_: _dafny.Map
        d_970_v17_ = _dafny.Map({d_969_v16_: False})
        hi8_ = default__.safeDivisionInt(p0, d_968_v15_)
        for d_971_i5_ in range(len((d_970_v17_) | (d_970_v17_)), hi8_):
            d_972_v18_: D9
            d_972_v18_ = D9_DC22((d_968_v15_) - (d_967_v14_))
            source11_ = d_972_v18_
            if source11_.is_DC22:
                d_973___mcc_h0_ = source11_.cf38
                d_974_cf38_ = d_973___mcc_h0_
                d_975_v19_: _dafny.Seq
                d_975_v19_ = _dafny.SeqWithoutIsStrInference([d_969_v16_])
                d_976_v20_: _dafny.MultiSet
                d_976_v20_ = _dafny.MultiSet([d_975_v19_, (_dafny.SeqWithoutIsStrInference([d_947_v0_])) + (d_975_v19_), d_975_v19_])
                d_976_v20_ = ((d_976_v20_) | (d_976_v20_)) | (d_976_v20_)
                d_974_cf38_ = (p0) + ((d_967_v14_) + (d_968_v15_))
                d_977_v21_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.Seq({}), 23)
                d_977_v21_ = nw165_
                index156_ = default__.safeIndex(675, (d_977_v21_).length(0))
                (d_977_v21_)[index156_] = (default__.fm38(d_969_v16_, d_967_v14_, d_969_v16_, globalState)) + (d_975_v19_)
                index157_ = default__.safeIndex(675, (d_977_v21_).length(0))
                (d_977_v21_)[index157_] = d_975_v19_
                d_978_v22_: D2
                d_978_v22_ = D2_DC5(d_974_cf38_, _dafny.SeqWithoutIsStrInference([(0) - (d_967_v14_) for d_979_i6_ in range(default__.abs(24))]), d_969_v16_)
                d_980_v23_: D2
                d_980_v23_ = D2_DC6(d_978_v22_)
                d_981_v24_: _dafny.Map
                d_981_v24_ = _dafny.Map({d_966_v13_: p0})
                d_982_v25_: _dafny.Seq
                d_982_v25_ = _dafny.SeqWithoutIsStrInference([d_981_v24_])
                d_983_v26_: _dafny.MultiSet
                d_983_v26_ = _dafny.MultiSet([d_970_v17_, _dafny.Map({True: d_966_v13_}), (d_970_v17_).set(False, d_947_v0_), d_970_v17_, d_970_v17_])
                d_984_v27_: _dafny.Seq
                d_984_v27_ = _dafny.SeqWithoutIsStrInference([d_971_i5_, len((default__.fm38(d_966_v13_, len((d_982_v25_)[default__.safeIndex(((d_983_v26_)[d_970_v17_] if (d_970_v17_) in (d_983_v26_) else (0) - (default__.fm32(d_968_v15_, globalState))), len(d_982_v25_))]), d_969_v16_, globalState)).set(default__.safeIndex(d_974_cf38_, len(default__.fm38(d_966_v13_, len((d_982_v25_)[default__.safeIndex(((d_983_v26_)[d_970_v17_] if (d_970_v17_) in (d_983_v26_) else (0) - (default__.fm32(d_968_v15_, globalState))), len(d_982_v25_))]), d_969_v16_, globalState))), d_947_v0_)), d_968_v15_, default__.fm32(len((d_977_v21_)[default__.safeIndex(675, (d_977_v21_).length(0))]), globalState), d_974_cf38_])
                d_985_v28_: _dafny.Map
                d_985_v28_ = _dafny.Map({d_980_v23_: (d_984_v27_) != (d_984_v27_)})
                d_947_v0_ = ((d_985_v28_)[d_980_v23_] if (d_980_v23_) in (d_985_v28_) else d_969_v16_)
            elif True:
                d_986___mcc_h1_ = source11_.cf37
                d_987_cf37_ = d_986___mcc_h1_
                d_988_v29_: _dafny.Map
                d_988_v29_ = _dafny.Map({d_967_v14_: d_947_v0_})
                d_989_v30_: _dafny.Seq
                d_989_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmpc"))
                d_990_v31_: _dafny.Set
                d_990_v31_ = _dafny.Set({d_969_v16_})
                d_991_v32_: _dafny.Map
                d_991_v32_ = _dafny.Map({len(d_989_v30_): d_990_v31_})
                d_992_v33_: D12
                d_992_v33_ = D12_DC27(((d_991_v32_)[d_968_v15_] if (d_968_v15_) in (d_991_v32_) else d_990_v31_))
                d_993_v34_: _dafny.Map
                d_993_v34_ = _dafny.Map({(d_988_v29_).set(d_971_i5_, d_947_v0_): d_992_v33_})
                pat_let_tv28_ = d_990_v31_
                def iife89_(_pat_let22_0):
                    def iife90_(d_994_dt__update__tmp_h0_):
                        def iife91_(_pat_let23_0):
                            def iife92_(d_995_dt__update_hcf45_h0_):
                                return D12_DC27(d_995_dt__update_hcf45_h0_)
                            return iife92_(_pat_let23_0)
                        return iife91_(pat_let_tv28_)
                    return iife90_(_pat_let22_0)
                d_993_v34_ = (d_993_v34_).set(d_988_v29_, iife89_(d_992_v33_))
                d_996_v35_: _dafny.Set
                d_996_v35_ = _dafny.Set({d_952_v5_})
                d_997_v36_: _dafny.Set
                d_997_v36_ = d_996_v35_
                d_998_v37_: _dafny.Array
                nw166_ = _dafny.Array(None, 22)
                nw166_[int(0)] = d_997_v36_
                nw166_[int(1)] = d_996_v35_
                nw166_[int(2)] = _dafny.Set({d_952_v5_})
                nw166_[int(3)] = d_997_v36_
                nw166_[int(4)] = d_996_v35_
                nw166_[int(5)] = d_997_v36_
                nw166_[int(6)] = d_997_v36_
                nw166_[int(7)] = d_997_v36_
                nw166_[int(8)] = d_997_v36_
                nw166_[int(9)] = d_996_v35_
                nw166_[int(10)] = d_997_v36_
                nw166_[int(11)] = d_997_v36_
                nw166_[int(12)] = d_997_v36_
                nw166_[int(13)] = (d_997_v36_ if not(d_947_v0_) else d_996_v35_)
                nw166_[int(14)] = d_997_v36_
                nw166_[int(15)] = d_997_v36_
                nw166_[int(16)] = d_997_v36_
                nw166_[int(17)] = d_996_v35_
                nw166_[int(18)] = d_997_v36_
                nw166_[int(19)] = d_997_v36_
                nw166_[int(20)] = d_997_v36_
                nw166_[int(21)] = d_997_v36_
                d_998_v37_ = nw166_
                index158_ = default__.safeIndex(397, (d_998_v37_).length(0))
                (d_998_v37_)[index158_] = d_997_v36_
                index159_ = default__.safeIndex(397, (d_998_v37_).length(0))
                (d_998_v37_)[index159_] = d_997_v36_
                d_999_v38_: str
                d_999_v38_ = _dafny.CodePoint('f')
                d_1000_v39_: _dafny.Seq
                d_1000_v39_ = _dafny.SeqWithoutIsStrInference([d_969_v16_])
                d_1001_v40_: _dafny.Map
                d_1001_v40_ = _dafny.Map({d_999_v38_: d_968_v15_})
                d_1002_v42_: _dafny.Set
                d_1002_v42_ = _dafny.Set({d_999_v38_, d_999_v38_, d_999_v38_})
                d_1003_v43_: _dafny.Seq
                d_1003_v43_ = _dafny.SeqWithoutIsStrInference([d_1001_v40_, d_1001_v40_])
                d_1004_v44_: _dafny.Array
                nw167_ = _dafny.Array(None, 25)
                nw167_[int(0)] = _dafny.Map({d_999_v38_: len(d_1000_v39_)})
                nw167_[int(1)] = d_1001_v40_
                nw167_[int(2)] = d_1001_v40_
                nw167_[int(3)] = d_1001_v40_
                def iife93_():
                    coll45_ = _dafny.Map()
                    compr_45_: str
                    for compr_45_ in (d_1002_v42_).Elements:
                        d_1005_v41_: str = compr_45_
                        if (d_1005_v41_) in (d_1002_v42_):
                            coll45_[d_1005_v41_] = len(_dafny.SeqWithoutIsStrInference([d_999_v38_ for d_1006_i7_ in range(default__.abs(-584))]))
                    return _dafny.Map(coll45_)
                nw167_[int(4)] = iife93_()
                
                nw167_[int(5)] = _dafny.Map({default__.fm29(d_947_v0_, globalState): d_968_v15_})
                nw167_[int(6)] = _dafny.Map({d_999_v38_: default__.fm32(d_968_v15_, globalState)})
                nw167_[int(7)] = d_1001_v40_
                nw167_[int(8)] = d_1001_v40_
                nw167_[int(9)] = d_1001_v40_
                nw167_[int(10)] = (d_1003_v43_)[default__.safeIndex(p0, len(d_1003_v43_))]
                nw167_[int(11)] = (d_1003_v43_)[default__.safeIndex(d_967_v14_, len(d_1003_v43_))]
                nw167_[int(12)] = (d_1001_v40_)
                nw167_[int(13)] = d_1001_v40_
                nw167_[int(14)] = d_1001_v40_
                nw167_[int(15)] = d_1001_v40_
                nw167_[int(16)] = d_1001_v40_
                nw167_[int(17)] = _dafny.Map({d_999_v38_: d_971_i5_})
                nw167_[int(18)] = d_1001_v40_
                nw167_[int(19)] = d_1001_v40_
                nw167_[int(20)] = d_1001_v40_
                nw167_[int(21)] = d_1001_v40_
                nw167_[int(22)] = d_1001_v40_
                nw167_[int(23)] = d_1001_v40_
                nw167_[int(24)] = d_1001_v40_
                d_1004_v44_ = nw167_
                d_1007_v45_: _dafny.Seq
                d_1007_v45_ = _dafny.SeqWithoutIsStrInference([d_967_v14_])
                d_1008_v46_: D0
                d_1008_v46_ = D0_DC1(p0)
                d_1009_v47_: C1
                nw168_ = C1()
                nw168_.ctor__((d_971_i5_) + (d_971_i5_), d_1004_v44_, d_1007_v45_, d_1008_v46_)
                d_1009_v47_ = nw168_
                index160_ = default__.safeIndex(582, (d_952_v5_).length(0))
                (d_952_v5_)[index160_] = d_966_v13_
                index161_ = default__.safeIndex(582, (d_952_v5_).length(0))
                (d_952_v5_)[index161_] = (d_968_v15_) in (default__.fm39(d_966_v13_, globalState))
            d_968_v15_ = d_971_i5_
            d_1010_v48_: str
            d_1010_v48_ = _dafny.CodePoint('t')
            d_1011_v50_: D7
            def iife94_():
                coll46_ = _dafny.Set()
                compr_46_: int
                for compr_46_ in _dafny.IntegerRange(195, 383):
                    d_1012_v49_: int = compr_46_
                    if ((195) <= (d_1012_v49_)) and ((d_1012_v49_) < (383)):
                        coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1012_v49_, d_971_i5_)]))
                return _dafny.Set(coll46_)
            d_1011_v50_ = D7_DC19(len(iife94_()
))
            source12_ = default__.fm40(d_966_v13_, d_1010_v48_, d_969_v16_, d_1011_v50_, globalState)
            if source12_.is_DC28:
                d_1013___mcc_h2_ = source12_.cf46
                d_1014___mcc_h3_ = source12_.cf47
                d_1015_cf47_ = d_1014___mcc_h3_
                d_1016_cf46_ = d_1013___mcc_h2_
                d_968_v15_ = (0) - (d_971_i5_)
                d_1017_v51_: _dafny.Array
                nw169_ = _dafny.Array(int(0), 17)
                d_1017_v51_ = nw169_
                index162_ = default__.safeIndex(625, (d_1017_v51_).length(0))
                (d_1017_v51_)[index162_] = p0
                d_1018_v52_: _dafny.Map
                d_1018_v52_ = _dafny.Map({d_971_i5_: d_1017_v51_})
                index163_ = default__.safeIndex(625, (d_1017_v51_).length(0))
                (d_1017_v51_)[index163_] = len(_dafny.SeqWithoutIsStrInference([len(d_1018_v52_)]))
                d_1019_v53_: _dafny.Seq
                d_1019_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruavycy"))
                d_1020_v54_: D13
                d_1020_v54_ = D13_DC31(d_1016_cf46_, len(d_1019_v53_), len(_dafny.SeqWithoutIsStrInference([d_1010_v48_ for d_1021_i8_ in range(default__.abs(-504))])))
                d_1022_v55_: _dafny.Array
                nw170_ = _dafny.Array(None, 1)
                def iife95_(_pat_let24_0):
                    def iife96_(d_1023_dt__update__tmp_h4_):
                        def iife97_(_pat_let25_0):
                            def iife98_(d_1024_dt__update_hcf56_h0_):
                                return D13_DC31((d_1023_dt__update__tmp_h4_).cf54, (d_1023_dt__update__tmp_h4_).cf55, d_1024_dt__update_hcf56_h0_)
                            return iife98_(_pat_let25_0)
                        return iife97_(d_971_i5_)
                    return iife96_(_pat_let24_0)
                nw170_[int(0)] = iife95_(d_1020_v54_)
                d_1022_v55_ = nw170_
                d_1025_v56_: _dafny.Map
                d_1025_v56_ = _dafny.Map({default__.fm30(d_971_i5_, d_966_v13_, True, globalState): d_1022_v55_})
                d_1026_v57_: _dafny.Array
                nw171_ = _dafny.Array(None, 3)
                nw171_[int(0)] = d_1022_v55_
                nw171_[int(1)] = ((d_1025_v56_)[not(d_1016_cf46_)] if (not(d_1016_cf46_)) in (d_1025_v56_) else d_1022_v55_)
                nw171_[int(2)] = d_1022_v55_
                d_1026_v57_ = nw171_
                index164_ = default__.safeIndex(332, (d_1026_v57_).length(0))
                (d_1026_v57_)[index164_] = d_1022_v55_
                index165_ = default__.safeIndex(332, (d_1026_v57_).length(0))
                (d_1026_v57_)[index165_] = d_1022_v55_
                d_1027_v58_: _dafny.Seq
                d_1027_v58_ = _dafny.SeqWithoutIsStrInference([d_1015_cf47_])
                d_1028_v59_: _dafny.Map
                d_1028_v59_ = _dafny.Map({d_947_v0_: (0) - (len(_dafny.SeqWithoutIsStrInference([(0) - (d_968_v15_) for d_1029_i9_ in range(default__.abs(977))])))})
                rhs98_ = d_969_v16_
                rhs99_ = d_1019_v53_
                rhs100_ = (_dafny.SeqWithoutIsStrInference([(d_1020_v54_).cf55]) if (226) == (d_967_v14_) else _dafny.SeqWithoutIsStrInference([d_971_i5_, d_1015_cf47_, default__.fm32(d_971_i5_, globalState), len(d_1028_v59_)]))
                rhs101_ = d_969_v16_
                d_1016_cf46_ = rhs98_
                d_1019_v53_ = rhs99_
                d_1027_v58_ = rhs100_
                d_966_v13_ = rhs101_
            elif source12_.is_DC29:
                d_1030___mcc_h4_ = source12_.cf48
                d_1031___mcc_h5_ = source12_.cf49
                d_1032___mcc_h6_ = source12_.cf50
                d_1033___mcc_h7_ = source12_.cf51
                d_1034___mcc_h8_ = source12_.cf52
                d_1035_cf52_ = d_1034___mcc_h8_
                d_1036_cf51_ = d_1033___mcc_h7_
                d_1037_cf50_ = d_1032___mcc_h6_
                d_1038_cf49_ = d_1031___mcc_h5_
                d_1039_cf48_ = d_1030___mcc_h4_
                d_1040_v60_: _dafny.Seq
                d_1040_v60_ = _dafny.SeqWithoutIsStrInference([d_1039_cf48_, d_1035_cf52_])
                index166_ = default__.safeIndex(626, (d_952_v5_).length(0))
                (d_952_v5_)[index166_] = (d_1040_v60_)[default__.safeIndex(d_971_i5_, len(d_1040_v60_))]
                index167_ = default__.safeIndex(626, (d_952_v5_).length(0))
                (d_952_v5_)[index167_] = (d_969_v16_) or (not (True) or (d_966_v13_))
                d_1041_v61_: _dafny.Seq
                d_1041_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnk"))
                d_1042_v62_: D6
                d_1042_v62_ = D6_DC14(d_1041_v61_)
                d_1042_v62_ = d_1042_v62_
                d_1043_v63_: _dafny.Map
                d_1043_v63_ = _dafny.Map({d_1010_v48_: d_1036_cf51_})
                d_1044_v65_: _dafny.Seq
                d_1044_v65_ = _dafny.SeqWithoutIsStrInference([(d_1043_v63_).set(d_1010_v48_, p0), _dafny.Map({d_1010_v48_: d_967_v14_})])
                d_1045_v66_: _dafny.Array
                nw172_ = _dafny.Array(None, 18)
                nw172_[int(0)] = d_1043_v63_
                nw172_[int(1)] = _dafny.Map({d_1010_v48_: (0) - ((d_948_v1_).cardinality)})
                nw172_[int(2)] = d_1043_v63_
                nw172_[int(3)] = d_1043_v63_
                nw172_[int(4)] = d_1043_v63_
                def iife99_():
                    coll47_ = _dafny.Map()
                    compr_47_: str
                    for compr_47_ in ((d_1043_v63_).set(d_1010_v48_, d_1036_cf51_)).keys.Elements:
                        d_1046_v64_: str = compr_47_
                        if (d_1046_v64_) in ((d_1043_v63_).set(d_1010_v48_, d_1036_cf51_)):
                            coll47_[d_1046_v64_] = 813
                    return _dafny.Map(coll47_)
                nw172_[int(5)] = iife99_()
                
                nw172_[int(6)] = d_1043_v63_
                nw172_[int(7)] = d_1043_v63_
                nw172_[int(8)] = d_1043_v63_
                nw172_[int(9)] = d_1043_v63_
                nw172_[int(10)] = d_1043_v63_
                nw172_[int(11)] = d_1043_v63_
                nw172_[int(12)] = d_1043_v63_
                nw172_[int(13)] = d_1043_v63_
                nw172_[int(14)] = (d_1043_v63_).set(_dafny.CodePoint('k'), d_971_i5_)
                nw172_[int(15)] = d_1043_v63_
                nw172_[int(16)] = (d_1044_v65_)[default__.safeIndex(d_968_v15_, len(d_1044_v65_))]
                nw172_[int(17)] = _dafny.Map({d_1010_v48_: ((d_948_v1_)[d_1037_cf50_] if (d_1037_cf50_) in (d_948_v1_) else d_967_v14_)})
                d_1045_v66_ = nw172_
                d_1047_v67_: _dafny.Seq
                d_1047_v67_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1048_v68_: C1
                nw173_ = C1()
                nw173_.ctor__((0) - (d_1036_cf51_), (d_1045_v66_ if d_969_v16_ else d_1045_v66_), ((d_1047_v67_).set(default__.safeIndex(551, len(d_1047_v67_)), d_971_i5_)) + (d_1047_v67_), D0_DC1(default__.fm32(-714, globalState)))
                d_1048_v68_ = nw173_
                d_1049_v69_: _dafny.Map
                d_1049_v69_ = _dafny.Map({((d_948_v1_) - (d_948_v1_)).cardinality: 111})
                d_1049_v69_ = (d_1049_v69_).set(p0, (0) - ((d_971_i5_) - (d_1036_cf51_)))
            elif True:
                d_1050___mcc_h9_ = source12_.cf45
                d_1051_cf45_ = d_1050___mcc_h9_
                d_1052_v70_: _dafny.Array
                nw174_ = _dafny.Array(_dafny.Map({}), 21)
                d_1052_v70_ = nw174_
                d_1053_v71_: T0
                nw175_ = C2()
                nw175_.ctor__()
                d_1053_v71_ = nw175_
                d_1054_v72_: _dafny.Map
                d_1054_v72_ = _dafny.Map({d_1053_v71_: True})
                d_1055_v73_: _dafny.Seq
                d_1055_v73_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1010_v48_ for d_1056_i10_ in range(default__.abs(474))])), len(d_1054_v72_)])
                d_1057_v74_: D0
                d_1057_v74_ = D0_DC1(d_968_v15_)
                d_1058_v75_: C1
                nw176_ = C1()
                nw176_.ctor__(d_968_v15_, d_1052_v70_, d_1055_v73_, d_1057_v74_)
                d_1058_v75_ = nw176_
                d_1059_v76_: _dafny.Seq
                d_1059_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkjicxbn"))
                d_1060_v77_: C0
                nw177_ = C0()
                nw177_.ctor__((d_1059_v76_) + (d_1059_v76_))
                d_1060_v77_ = nw177_
                d_1061_v78_: _dafny.Array
                def lambda41_(d_1062_v75_):
                    def lambda42_(d_1063_i11_):
                        return default__.safeDivisionInt(d_1063_i11_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1062_v75_.f9])])))

                    return lambda42_

                init22_ = lambda41_(d_1058_v75_)
                nw178_ = _dafny.Array(None, 18)
                for i0_22_ in range(nw178_.length(0)):
                    nw178_[i0_22_] = init22_(i0_22_)
                d_1061_v78_ = nw178_
                index168_ = default__.safeIndex(84, (d_1061_v78_).length(0))
                (d_1061_v78_)[index168_] = len(d_1060_v77_.f8)
                index169_ = default__.safeIndex(84, (d_1061_v78_).length(0))
                (d_1061_v78_)[index169_] = d_971_i5_
                d_1064_v79_: _dafny.Set
                d_1064_v79_ = _dafny.Set({d_971_i5_})
                d_1064_v79_ = d_1064_v79_
            d_1010_v48_ = d_1010_v48_
        d_1065_v80_: _dafny.Seq
        d_1065_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohddunlw"))
        d_1066_v81_: _dafny.Set
        d_1066_v81_ = _dafny.Set({d_969_v16_, d_947_v0_})
        d_1067_v82_: _dafny.Seq
        d_1067_v82_ = _dafny.SeqWithoutIsStrInference([d_968_v15_, d_968_v15_, len(d_1066_v81_)])
        d_1068_v83_: _dafny.Map
        d_1068_v83_ = _dafny.Map({(d_1065_v80_) < (d_1065_v80_): _dafny.MultiSet((d_1067_v82_) + (d_1067_v82_))})
        r0 = d_1068_v83_
        return r0

    def m11(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Set = _dafny.Set({})
        d_1069_v0_: bool
        d_1069_v0_ = True
        index170_ = default__.safeIndex(659, (p1).length(0))
        (p1)[index170_] = d_1069_v0_
        d_1070_v1_: _dafny.Seq
        d_1070_v1_ = _dafny.SeqWithoutIsStrInference([p3])
        d_1071_v2_: _dafny.Seq
        d_1071_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
        d_1072_v3_: _dafny.Set
        d_1072_v3_ = _dafny.Set({_dafny.MultiSet([p3])})
        d_1073_v4_: _dafny.Map
        d_1073_v4_ = _dafny.Map({len(d_1071_v2_): d_1072_v3_})
        d_1074_v5_: _dafny.Array
        nw179_ = _dafny.Array(None, 3)
        nw179_[int(0)] = p3
        nw179_[int(1)] = p3
        nw179_[int(2)] = p3
        d_1074_v5_ = nw179_
        d_1075_v6_: _dafny.Map
        d_1075_v6_ = _dafny.Map({d_1074_v5_: p3})
        index171_ = default__.safeIndex(659, (p1).length(0))
        rhs102_ = True
        rhs103_ = not((((d_1073_v4_)[-931] if (-931) in (d_1073_v4_) else d_1072_v3_)).isdisjoint(d_1072_v3_))
        rhs104_ = _dafny.SeqWithoutIsStrInference([len(d_1075_v6_)])
        rhs105_ = d_1069_v0_
        lhs66_ = p1
        lhs67_ = default__.safeIndex(659, (p1).length(0))
        lhs66_[lhs67_] = rhs102_
        d_1069_v0_ = rhs103_
        d_1070_v1_ = rhs104_
        r1 = rhs105_
        if d_1069_v0_:
            d_1069_v0_ = d_1069_v0_
            d_1071_v2_ = d_1071_v2_
            index172_ = default__.safeIndex(659, (p1).length(0))
            (p1)[index172_] = (p1)[default__.safeIndex(659, (p1).length(0))]
            d_1076_v7_: _dafny.Map
            d_1076_v7_ = _dafny.Map({d_1069_v0_: (_dafny.SeqWithoutIsStrInference([p0, p0]) if d_1069_v0_ else d_1071_v2_)})
            d_1077_v8_: int
            d_1077_v8_ = -951
            d_1078_v9_: _dafny.Array
            def lambda43_(d_1079_p0_):
                def lambda44_(d_1080_i0_):
                    return d_1079_p0_

                return lambda44_

            init23_ = lambda43_(p0)
            nw180_ = _dafny.Array(None, 11)
            for i0_23_ in range(nw180_.length(0)):
                nw180_[i0_23_] = init23_(i0_23_)
            d_1078_v9_ = nw180_
            index173_ = default__.safeIndex(623, (d_1078_v9_).length(0))
            (d_1078_v9_)[index173_] = _dafny.CodePoint('v')
            index174_ = default__.safeIndex(623, (d_1078_v9_).length(0))
            rhs106_ = (d_1076_v7_) | (d_1076_v7_)
            rhs107_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([p0 for d_1081_i1_ in range(default__.abs(803))])), (491) - (d_1077_v8_))
            rhs108_ = p0
            rhs109_ = d_1077_v8_
            lhs68_ = d_1078_v9_
            lhs69_ = default__.safeIndex(623, (d_1078_v9_).length(0))
            d_1076_v7_ = rhs106_
            d_1077_v8_ = rhs107_
            lhs68_[lhs69_] = rhs108_
            d_1077_v8_ = rhs109_
            r1 = (d_1077_v8_) == (d_1077_v8_)
        elif True:
            d_1082_v10_: str
            d_1082_v10_ = _dafny.CodePoint('c')
            d_1083_v11_: int
            d_1083_v11_ = 388
            index175_ = default__.safeIndex(98, (d_1074_v5_).length(0))
            (d_1074_v5_)[index175_] = (p3) * (default__.fm32(d_1083_v11_, globalState))
            d_1084_v12_: D7
            d_1084_v12_ = D7_DC18()
            d_1085_v13_: _dafny.Map
            d_1085_v13_ = _dafny.Map({(p1)[default__.safeIndex(659, (p1).length(0))]: d_1084_v12_})
            d_1086_v14_: _dafny.Seq
            d_1086_v14_ = _dafny.SeqWithoutIsStrInference([d_1085_v13_, d_1085_v13_])
            d_1087_v15_: D13
            d_1087_v15_ = D13_DC31((p1)[default__.safeIndex(659, (p1).length(0))], d_1083_v11_, p3)
            d_1088_v16_: _dafny.MultiSet
            d_1088_v16_ = _dafny.MultiSet([default__.fm44(True, globalState), d_1083_v11_, d_1083_v11_, default__.fm44(d_1069_v0_, globalState)])
            d_1089_v17_: _dafny.MultiSet
            d_1089_v17_ = _dafny.MultiSet([d_1088_v16_, d_1088_v16_, (d_1088_v16_).intersection(_dafny.MultiSet(d_1070_v1_))])
            d_1090_v18_: _dafny.Seq
            d_1090_v18_ = _dafny.SeqWithoutIsStrInference([d_1069_v0_, d_1069_v0_])
            index176_ = default__.safeIndex(98, (d_1074_v5_).length(0))
            rhs110_ = default__.fm51(d_1069_v0_, p3, default__.safeDivisionInt(len(d_1071_v2_), (d_1070_v1_)[default__.safeIndex(len(d_1086_v14_), len(d_1070_v1_))]), globalState)
            rhs111_ = ((d_1087_v15_ if d_1069_v0_ else D13_DC31(d_1069_v0_, p3, p3))).cf54
            rhs112_ = ((d_1089_v17_)[((d_1088_v16_).set(d_1083_v11_, default__.abs(len(d_1090_v18_)))).intersection(_dafny.MultiSet([816]))] if (((d_1088_v16_).set(d_1083_v11_, default__.abs(len(d_1090_v18_)))).intersection(_dafny.MultiSet([816]))) in (d_1089_v17_) else d_1083_v11_)
            rhs113_ = not((p1)[default__.safeIndex(659, (p1).length(0))])
            rhs114_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghsv")))
            lhs70_ = d_1074_v5_
            lhs71_ = default__.safeIndex(98, (d_1074_v5_).length(0))
            d_1082_v10_ = rhs110_
            r2 = rhs111_
            d_1083_v11_ = rhs112_
            r2 = rhs113_
            lhs70_[lhs71_] = rhs114_
            if ((0) - ((d_1074_v5_)[default__.safeIndex(98, (d_1074_v5_).length(0))])) <= ((0) - (-737)):
                d_1091_v19_: _dafny.Set
                d_1091_v19_ = _dafny.Set({d_1082_v10_, p0, p0})
                d_1092_v20_: _dafny.Set
                d_1092_v20_ = _dafny.Set({d_1090_v18_, _dafny.SeqWithoutIsStrInference([d_1069_v0_])})
                d_1093_v21_: _dafny.Map
                d_1093_v21_ = _dafny.Map({len(d_1091_v19_): (d_1092_v20_).issubset(default__.fm55(p3, globalState))})
                d_1094_v22_: _dafny.Set
                d_1094_v22_ = _dafny.Set({830})
                d_1093_v21_ = (d_1093_v21_).set(len((_dafny.Set({(d_1074_v5_)[default__.safeIndex(98, (d_1074_v5_).length(0))], (0) - (d_1083_v11_), (d_1074_v5_)[default__.safeIndex(98, (d_1074_v5_).length(0))]})) - (d_1094_v22_)), (D11_DC25(d_1082_v10_, d_1083_v11_, False)).cf43)
                d_1095_v23_: _dafny.Array
                nw181_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_1095_v23_ = nw181_
                d_1095_v23_ = d_1095_v23_
                d_1096_v24_: _dafny.Array
                nw182_ = _dafny.Array(_dafny.Seq({}), 11)
                d_1096_v24_ = nw182_
                index177_ = default__.safeIndex(29, (d_1096_v24_).length(0))
                (d_1096_v24_)[index177_] = d_1070_v1_
                index178_ = default__.safeIndex(29, (d_1096_v24_).length(0))
                (d_1096_v24_)[index178_] = (d_1070_v1_) + (d_1070_v1_)
                d_1097_v25_: _dafny.MultiSet
                d_1097_v25_ = _dafny.MultiSet([True])
                d_1097_v25_ = (d_1097_v25_).set((p1)[default__.safeIndex(659, (p1).length(0))], default__.abs(d_1083_v11_))
                d_1098_v26_: _dafny.Seq
                d_1098_v26_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, (p1 if (p1)[default__.safeIndex(659, (p1).length(0))] else p1), p1])
                d_1098_v26_ = d_1098_v26_
            elif True:
                index179_ = default__.safeIndex(659, (p1).length(0))
                (p1)[index179_] = (d_1071_v2_) < (d_1071_v2_)
                d_1099_v27_: _dafny.Map
                d_1099_v27_ = _dafny.Map({(p1)[default__.safeIndex(659, (p1).length(0))]: (p1)[default__.safeIndex(659, (p1).length(0))]})
                d_1099_v27_ = (d_1099_v27_).set(True, (_dafny.SeqWithoutIsStrInference([d_1069_v0_])) != (d_1090_v18_))
                r1 = d_1069_v0_
                d_1100_v28_: _dafny.Set
                d_1100_v28_ = _dafny.Set({d_1069_v0_})
                d_1101_v29_: _dafny.Map
                d_1101_v29_ = _dafny.Map({d_1069_v0_: p3})
                index180_ = default__.safeIndex(659, (p1).length(0))
                (p1)[index180_] = (d_1100_v28_).isdisjoint(default__.fm53(d_1101_v29_, globalState))
                d_1102_v30_: _dafny.Map
                d_1102_v30_ = _dafny.Map({p3: d_1069_v0_})
                rhs115_ = ((0) - (d_1083_v11_)) >= ((len(_dafny.Map({d_1083_v11_: d_1069_v0_}))) * (len(d_1102_v30_)))
                rhs116_ = d_1071_v2_
                rhs117_ = (default__.fm44((p1)[default__.safeIndex(659, (p1).length(0))], globalState)) * (p3)
                r1 = rhs115_
                d_1071_v2_ = rhs116_
                d_1083_v11_ = rhs117_
            index181_ = default__.safeIndex(98, (d_1074_v5_).length(0))
            (d_1074_v5_)[index181_] = (d_1074_v5_)[default__.safeIndex(98, (d_1074_v5_).length(0))]
            d_1103_v31_: D12
            d_1103_v31_ = D12_DC29(False, p3, (p1)[default__.safeIndex(659, (p1).length(0))], d_1083_v11_, (p1)[default__.safeIndex(659, (p1).length(0))])
            index182_ = default__.safeIndex(98, (d_1074_v5_).length(0))
            (d_1074_v5_)[index182_] = (d_1103_v31_).cf49
            d_1104_v32_: D12
            d_1104_v32_ = D12_DC28((p1)[default__.safeIndex(659, (p1).length(0))], (0) - (d_1083_v11_))
            d_1105_v33_: _dafny.Set
            d_1105_v33_ = _dafny.Set({(0) - (p3), d_1083_v11_})
            pat_let_tv29_ = d_1105_v33_
            d_1106_v34_: _dafny.MultiSet
            def iife100_(_pat_let26_0):
                def iife101_(d_1107_dt__update__tmp_h0_):
                    def iife102_(_pat_let27_0):
                        def iife103_(d_1108_dt__update_hcf47_h0_):
                            return D12_DC28((d_1107_dt__update__tmp_h0_).cf46, d_1108_dt__update_hcf47_h0_)
                        return iife103_(_pat_let27_0)
                    return iife102_(len(pat_let_tv29_))
                return iife101_(_pat_let26_0)
            d_1106_v34_ = _dafny.MultiSet([D12_DC28((p1)[default__.safeIndex(659, (p1).length(0))], d_1083_v11_), d_1104_v32_, d_1104_v32_, iife100_(d_1104_v32_), d_1104_v32_])
            index183_ = default__.safeIndex(98, (d_1074_v5_).length(0))
            (d_1074_v5_)[index183_] = (d_1106_v34_).cardinality
        index184_ = default__.safeIndex(913, (d_1074_v5_).length(0))
        (d_1074_v5_)[index184_] = len(d_1071_v2_)
        index185_ = default__.safeIndex(913, (d_1074_v5_).length(0))
        (d_1074_v5_)[index185_] = p3
        d_1069_v0_ = not (not (not(d_1069_v0_)) or ((p1)[default__.safeIndex(659, (p1).length(0))])) or (d_1069_v0_)
        d_1109_v35_: bool
        d_1110_v36_: int
        d_1111_v37_: int
        d_1112_v38_: bool
        out33_: bool
        out34_: int
        out35_: int
        out36_: bool
        out33_, out34_, out35_, out36_ = (self).m12((p1)[default__.safeIndex(659, (p1).length(0))], globalState)
        d_1109_v35_ = out33_
        d_1110_v36_ = out34_
        d_1111_v37_ = out35_
        d_1112_v38_ = out36_
        index186_ = default__.safeIndex(913, (d_1074_v5_).length(0))
        (d_1074_v5_)[index186_] = d_1110_v36_
        d_1113_v39_: _dafny.Map
        d_1113_v39_ = _dafny.Map({d_1112_v38_: (p1)[default__.safeIndex(659, (p1).length(0))]})
        d_1114_v40_: D16
        d_1114_v40_ = D16_DC42(173, d_1110_v36_, d_1110_v36_)
        d_1115_v41_: _dafny.Map
        d_1115_v41_ = _dafny.Map({default__.fm30(d_1111_v37_, (p1)[default__.safeIndex(659, (p1).length(0))], d_1109_v35_, globalState): (d_1074_v5_)[default__.safeIndex(913, (d_1074_v5_).length(0))]})
        pat_let_tv30_ = d_1115_v41_
        def iife104_(_pat_let28_0):
            def iife105_(d_1116_dt__update__tmp_h1_):
                def iife106_(_pat_let29_0):
                    def iife107_(d_1117_dt__update_hcf76_h0_):
                        return D16_DC42(d_1117_dt__update_hcf76_h0_, (d_1116_dt__update__tmp_h1_).cf77, (d_1116_dt__update__tmp_h1_).cf78)
                    return iife107_(_pat_let29_0)
                return iife106_(len(pat_let_tv30_))
            return iife105_(_pat_let28_0)
        r0 = (_dafny.Map({(d_1074_v5_)[default__.safeIndex(913, (d_1074_v5_).length(0))]: d_1113_v39_})).set((iife104_(d_1114_v40_)).cf78, d_1113_v39_)
        r1 = (p1)[default__.safeIndex(659, (p1).length(0))]
        d_1118_v42_: D0
        d_1118_v42_ = D0_DC1(d_1110_v36_)
        d_1119_v43_: _dafny.Seq
        d_1119_v43_ = _dafny.SeqWithoutIsStrInference([d_1109_v35_, (p1)[default__.safeIndex(659, (p1).length(0))]])
        pat_let_tv31_ = d_1074_v5_
        pat_let_tv32_ = d_1074_v5_
        def iife108_(_pat_let30_0):
            def iife109_(d_1120_dt__update__tmp_h2_):
                def iife110_(_pat_let31_0):
                    def iife111_(d_1121_dt__update_hcf2_h0_):
                        return D0_DC1(d_1121_dt__update_hcf2_h0_)
                    return iife111_(_pat_let31_0)
                return iife110_((pat_let_tv32_)[default__.safeIndex(913, (pat_let_tv31_).length(0))])
            return iife109_(_pat_let30_0)
        r2 = (self).fm5((_dafny.MultiSet([(p1)[default__.safeIndex(659, (p1).length(0))]])).cardinality, d_1109_v35_, (_dafny.Map({p0: iife108_(d_1118_v42_)})).set(p0, D0_DC1(len(d_1119_v43_))), True, globalState)
        d_1122_v44_: _dafny.Set
        d_1122_v44_ = _dafny.Set({d_1115_v41_, d_1115_v41_, d_1115_v41_, d_1115_v41_})
        r3 = (_dafny.Set({_dafny.Map({(p1)[default__.safeIndex(659, (p1).length(0))]: (d_1074_v5_)[default__.safeIndex(913, (d_1074_v5_).length(0))]})})).intersection(d_1122_v44_)
        return r0, r1, r2, r3

    def m12(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_1123_v0_: int
        d_1123_v0_ = 886
        d_1124_v1_: _dafny.Seq
        d_1124_v1_ = _dafny.SeqWithoutIsStrInference([d_1123_v0_])
        rhs118_ = False
        rhs119_ = d_1123_v0_
        rhs120_ = (_dafny.SeqWithoutIsStrInference([d_1123_v0_])).set(default__.safeIndex(d_1123_v0_, len(_dafny.SeqWithoutIsStrInference([d_1123_v0_]))), 667)
        r3 = rhs118_
        r1 = rhs119_
        d_1124_v1_ = rhs120_
        hi9_ = d_1123_v0_
        for d_1125_i0_ in range(d_1123_v0_, hi9_):
            d_1126_v2_: _dafny.Seq
            d_1126_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxccix"))
            if (d_1123_v0_) != (len(d_1126_v2_)):
                rhs121_ = not(p0)
                rhs122_ = p0
                r3 = rhs121_
                r3 = rhs122_
                d_1127_v3_: _dafny.Array
                nw183_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1127_v3_ = nw183_
                d_1128_v4_: _dafny.Map
                d_1128_v4_ = _dafny.Map({not(p0): p0})
                d_1129_v5_: _dafny.Set
                d_1129_v5_ = _dafny.Set({len(d_1128_v4_), 415})
                d_1130_v7_: _dafny.Map
                d_1130_v7_ = _dafny.Map({p0: d_1125_i0_})
                d_1131_v8_: D14
                def iife112_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in _dafny.IntegerRange(381, 529):
                        d_1132_v6_: int = compr_48_
                        if ((381) <= (d_1132_v6_)) and ((d_1132_v6_) < (529)):
                            coll48_[(d_1132_v6_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = d_1125_i0_
                    return _dafny.Map(coll48_)
                d_1131_v8_ = D14_DC36(iife112_()
, d_1130_v7_)
                d_1133_v9_: _dafny.Map
                d_1133_v9_ = _dafny.Map({len(d_1126_v2_): d_1125_i0_})
                d_1134_v10_: _dafny.Array
                nw184_ = _dafny.Array(None, 6)
                nw184_[int(0)] = d_1131_v8_
                nw184_[int(1)] = D14_DC36(d_1133_v9_, d_1130_v7_)
                nw184_[int(2)] = d_1131_v8_
                nw184_[int(3)] = d_1131_v8_
                nw184_[int(4)] = d_1131_v8_
                nw184_[int(5)] = d_1131_v8_
                d_1134_v10_ = nw184_
                d_1135_v11_: _dafny.Seq
                d_1135_v11_ = _dafny.SeqWithoutIsStrInference([d_1134_v10_, d_1134_v10_])
                d_1136_v14_: _dafny.Array
                nw185_ = _dafny.Array(None, 25)
                nw185_[int(0)] = d_1129_v5_
                nw185_[int(1)] = d_1129_v5_
                nw185_[int(2)] = d_1129_v5_
                nw185_[int(3)] = d_1129_v5_
                nw185_[int(4)] = d_1129_v5_
                nw185_[int(5)] = _dafny.Set({len(d_1135_v11_), d_1125_i0_})
                nw185_[int(6)] = _dafny.Set({(0) - (d_1125_i0_), d_1125_i0_})
                nw185_[int(7)] = d_1129_v5_
                nw185_[int(8)] = d_1129_v5_
                nw185_[int(9)] = d_1129_v5_
                nw185_[int(10)] = d_1129_v5_
                nw185_[int(11)] = _dafny.Set({default__.fm44(p0, globalState)})
                nw185_[int(12)] = d_1129_v5_
                nw185_[int(13)] = d_1129_v5_
                def iife113_():
                    coll49_ = _dafny.Set()
                    compr_49_: int
                    for compr_49_ in _dafny.IntegerRange(0, 289):
                        d_1137_v12_: int = compr_49_
                        if ((0) <= (d_1137_v12_)) and ((d_1137_v12_) < (289)):
                            coll49_ = coll49_.union(_dafny.Set([(d_1137_v12_) + (len(d_1124_v1_))]))
                    return _dafny.Set(coll49_)
                nw185_[int(14)] = iife113_()
                
                nw185_[int(15)] = d_1129_v5_
                nw185_[int(16)] = d_1129_v5_
                nw185_[int(17)] = d_1129_v5_
                nw185_[int(18)] = d_1129_v5_
                nw185_[int(19)] = d_1129_v5_
                nw185_[int(20)] = d_1129_v5_
                nw185_[int(21)] = _dafny.Set({d_1123_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1138_i1_ in range(default__.abs(-591))])), d_1125_i0_, d_1125_i0_})
                nw185_[int(22)] = d_1129_v5_
                nw185_[int(23)] = d_1129_v5_
                def iife114_():
                    coll50_ = _dafny.Set()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(578, 846):
                        d_1139_v13_: int = compr_50_
                        if ((578) <= (d_1139_v13_)) and ((d_1139_v13_) < (846)):
                            coll50_ = coll50_.union(_dafny.Set([(d_1139_v13_) - (len(d_1128_v4_))]))
                    return _dafny.Set(coll50_)
                nw185_[int(24)] = iife114_()
                
                d_1136_v14_ = nw185_
                index187_ = default__.safeIndex(733, (d_1127_v3_).length(0))
                (d_1127_v3_)[index187_] = d_1136_v14_
                index188_ = default__.safeIndex(733, (d_1127_v3_).length(0))
                (d_1127_v3_)[index188_] = d_1136_v14_
                def lambda45_(d_1140_v5_):
                    def lambda46_(d_1141_i2_):
                        return d_1140_v5_

                    return lambda46_

                init24_ = lambda45_(d_1129_v5_)
                nw186_ = _dafny.Array(None, 17)
                for i0_24_ in range(nw186_.length(0)):
                    nw186_[i0_24_] = init24_(i0_24_)
                d_1136_v14_ = nw186_
                d_1142_v15_: str
                d_1142_v15_ = _dafny.CodePoint('f')
                d_1143_v16_: C0
                nw187_ = C0()
                nw187_.ctor__((d_1126_v2_) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1144_i3_ in range(default__.abs(790))])).set(default__.safeIndex(len(d_1126_v2_), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1145_i3_ in range(default__.abs(790))]))), d_1142_v15_)))
                d_1143_v16_ = nw187_
                d_1146_v17_: D0
                d_1146_v17_ = D0_DC1(len(d_1130_v7_))
                d_1147_v18_: _dafny.Map
                d_1147_v18_ = _dafny.Map({d_1142_v15_: d_1146_v17_})
                d_1128_v4_ = (d_1128_v4_).set((self).fm5(d_1125_i0_, p0, d_1147_v18_, p0, globalState), (d_1126_v2_) != (d_1126_v2_))
            elif True:
                d_1148_v19_: D7
                d_1148_v19_ = D7_DC18()
                d_1148_v19_ = d_1148_v19_
                r3 = p0
                d_1149_v20_: _dafny.MultiSet
                d_1149_v20_ = _dafny.MultiSet([True, p0])
                d_1150_v21_: _dafny.Map
                d_1150_v21_ = _dafny.Map({d_1125_i0_: p0})
                d_1151_v22_: _dafny.Seq
                d_1151_v22_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1152_v23_: _dafny.Map
                d_1152_v23_ = _dafny.Map({_dafny.MultiSet(d_1151_v22_): p0})
                d_1153_v24_: _dafny.Seq
                d_1153_v24_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm30(d_1125_i0_, ((d_1150_v21_)[len(d_1152_v23_)] if (len(d_1152_v23_)) in (d_1150_v21_) else p0), True, globalState), p0, p0])
                d_1154_v25_: D4
                d_1154_v25_ = D4_DC11(d_1123_v0_, p0, p0, p0, p0)
                r3 = not(((d_1149_v20_).isdisjoint(_dafny.MultiSet(d_1151_v22_)) if p0 else (d_1154_v25_).cf26))
                d_1155_v26_: str
                d_1155_v26_ = _dafny.CodePoint('k')
                d_1155_v26_ = (d_1155_v26_ if p0 else (d_1126_v2_)[default__.safeIndex(d_1123_v0_, len(d_1126_v2_))])
                d_1156_v27_: _dafny.Array
                nw188_ = _dafny.Array(None, 2)
                nw188_[int(0)] = d_1123_v0_
                nw188_[int(1)] = d_1123_v0_
                d_1156_v27_ = nw188_
                d_1157_v28_: _dafny.Set
                d_1157_v28_ = _dafny.Set({d_1156_v27_})
                d_1158_v29_: _dafny.MultiSet
                d_1158_v29_ = _dafny.MultiSet([d_1125_i0_, len(d_1157_v28_)])
                d_1159_v30_: _dafny.Map
                d_1159_v30_ = _dafny.Map({d_1158_v29_: p0})
                r3 = (((d_1159_v30_)[d_1158_v29_] if (d_1158_v29_) in (d_1159_v30_) else True)) == (p0)
            r0 = p0
            d_1160_v31_: str
            d_1160_v31_ = _dafny.CodePoint('a')
            d_1161_v32_: _dafny.Map
            d_1161_v32_ = _dafny.Map({p0: d_1125_i0_})
            d_1162_v33_: _dafny.Map
            d_1162_v33_ = _dafny.Map({d_1160_v31_: len(d_1161_v32_)})
            d_1163_v34_: _dafny.Map
            d_1163_v34_ = d_1162_v33_
            d_1164_v35_: D11
            d_1164_v35_ = D11_DC25(d_1160_v31_, len((_dafny.Map({d_1163_v34_: d_1123_v0_})).set(d_1163_v34_, d_1125_i0_)), p0)
            source13_ = d_1164_v35_
            if source13_.is_DC25:
                d_1165___mcc_h0_ = source13_.cf41
                d_1166___mcc_h1_ = source13_.cf42
                d_1167___mcc_h2_ = source13_.cf43
                d_1168_cf43_ = d_1167___mcc_h2_
                d_1169_cf42_ = d_1166___mcc_h1_
                d_1170_cf41_ = d_1165___mcc_h0_
                d_1171_v36_: _dafny.Map
                d_1171_v36_ = _dafny.Map({d_1123_v0_: d_1125_i0_})
                d_1172_v37_: _dafny.Map
                d_1172_v37_ = _dafny.Map({d_1125_i0_: p0})
                d_1173_v38_: _dafny.Seq
                d_1173_v38_ = _dafny.SeqWithoutIsStrInference([d_1168_cf43_])
                d_1174_v39_: _dafny.Set
                d_1174_v39_ = _dafny.Set({d_1173_v38_})
                d_1175_v40_: _dafny.Array
                nw189_ = _dafny.Array(None, 28)
                nw189_[int(0)] = d_1125_i0_
                nw189_[int(1)] = d_1123_v0_
                nw189_[int(2)] = d_1125_i0_
                nw189_[int(3)] = (d_1123_v0_) * (((d_1171_v36_)[d_1125_i0_] if (d_1125_i0_) in (d_1171_v36_) else len(_dafny.SeqWithoutIsStrInference([d_1160_v31_ for d_1176_i4_ in range(default__.abs(413))]))))
                nw189_[int(4)] = d_1125_i0_
                nw189_[int(5)] = default__.fm32(d_1125_i0_, globalState)
                nw189_[int(6)] = (len(d_1126_v2_)) + (d_1123_v0_)
                nw189_[int(7)] = d_1125_i0_
                nw189_[int(8)] = d_1125_i0_
                nw189_[int(9)] = ((d_1124_v1_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))), len(d_1124_v1_))]) - (len(d_1172_v37_))
                nw189_[int(10)] = len((d_1174_v39_).intersection(d_1174_v39_))
                nw189_[int(11)] = d_1169_cf42_
                nw189_[int(12)] = d_1125_i0_
                nw189_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_1170_cf41_ for d_1177_i5_ in range(default__.abs(793))]))
                nw189_[int(14)] = d_1125_i0_
                nw189_[int(15)] = default__.safeModuloInt(487, d_1169_cf42_)
                nw189_[int(16)] = (d_1125_i0_) * (len(d_1126_v2_))
                nw189_[int(17)] = d_1169_cf42_
                nw189_[int(18)] = (0) - ((d_1169_cf42_) + (d_1125_i0_))
                nw189_[int(19)] = d_1169_cf42_
                nw189_[int(20)] = d_1169_cf42_
                nw189_[int(21)] = (0) - (d_1123_v0_)
                nw189_[int(22)] = (d_1169_cf42_) * (len(d_1126_v2_))
                nw189_[int(23)] = d_1125_i0_
                nw189_[int(24)] = (d_1125_i0_ if p0 else 940)
                nw189_[int(25)] = default__.safeDivisionInt(d_1123_v0_, d_1125_i0_)
                nw189_[int(26)] = (727) - ((0) - (len(_dafny.Set({d_1170_cf41_}))))
                nw189_[int(27)] = d_1125_i0_
                d_1175_v40_ = nw189_
                index189_ = default__.safeIndex(518, (d_1175_v40_).length(0))
                (d_1175_v40_)[index189_] = default__.safeModuloInt(default__.fm32(d_1123_v0_, globalState), 824)
                d_1178_v41_: _dafny.Map
                d_1178_v41_ = _dafny.Map({p0: d_1175_v40_})
                d_1179_v42_: _dafny.Map
                d_1179_v42_ = _dafny.Map({d_1178_v41_: d_1169_cf42_})
                index190_ = default__.safeIndex(518, (d_1175_v40_).length(0))
                (d_1175_v40_)[index190_] = ((d_1179_v42_)[_dafny.Map({p0: d_1175_v40_})] if (_dafny.Map({p0: d_1175_v40_})) in (d_1179_v42_) else d_1123_v0_)
                d_1180_v43_: _dafny.Array
                def lambda47_(d_1181_cf43_):
                    def lambda48_(d_1182_i6_):
                        return d_1181_cf43_

                    return lambda48_

                init25_ = lambda47_(d_1168_cf43_)
                nw190_ = _dafny.Array(None, 6)
                for i0_25_ in range(nw190_.length(0)):
                    nw190_[i0_25_] = init25_(i0_25_)
                d_1180_v43_ = nw190_
                d_1183_v44_: _dafny.Map
                d_1183_v44_ = _dafny.Map({len(d_1171_v36_): d_1180_v43_})
                d_1184_v45_: _dafny.Seq
                d_1184_v45_ = _dafny.SeqWithoutIsStrInference([((d_1183_v44_)[(d_1175_v40_)[default__.safeIndex(518, (d_1175_v40_).length(0))]] if ((d_1175_v40_)[default__.safeIndex(518, (d_1175_v40_).length(0))]) in (d_1183_v44_) else d_1180_v43_)])
                d_1185_v46_: _dafny.MultiSet
                d_1185_v46_ = _dafny.MultiSet([d_1180_v43_])
                d_1186_v47_: D16
                d_1186_v47_ = D16_DC41((d_1184_v45_)[default__.safeIndex(((d_1185_v46_)[d_1180_v43_] if (d_1180_v43_) in (d_1185_v46_) else len(d_1173_v38_)), len(d_1184_v45_))])
                d_1186_v47_ = d_1186_v47_
                d_1187_v48_: _dafny.Set
                d_1187_v48_ = _dafny.Set({d_1169_cf42_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1188_i7_ in range(default__.abs(956))]))})
                rhs123_ = d_1187_v48_
                rhs124_ = default__.safeModuloInt(411, d_1123_v0_)
                rhs125_ = (d_1123_v0_) * (d_1169_cf42_)
                rhs126_ = (False if p0 else d_1168_cf43_)
                rhs127_ = ((d_1175_v40_)[default__.safeIndex(518, (d_1175_v40_).length(0))] if not(p0) else default__.safeDivisionInt(d_1169_cf42_, d_1125_i0_))
                d_1187_v48_ = rhs123_
                r1 = rhs124_
                r1 = rhs125_
                r3 = rhs126_
                d_1123_v0_ = rhs127_
                d_1189_v49_: _dafny.Seq
                d_1189_v49_ = _dafny.SeqWithoutIsStrInference([d_1173_v38_])
                d_1190_v50_: D0
                d_1190_v50_ = D0_DC1((d_1175_v40_)[default__.safeIndex(518, (d_1175_v40_).length(0))])
                d_1191_v51_: _dafny.Map
                d_1191_v51_ = _dafny.Map({d_1170_cf41_: d_1190_v50_})
                r3 = (self).fm5((0) - ((d_1125_i0_) + (len(d_1189_v49_))), p0, d_1191_v51_, p0, globalState)
            elif source13_.is_DC24:
                d_1192___mcc_h3_ = source13_.cf40
                d_1193_cf40_ = d_1192___mcc_h3_
                r3 = p0
                d_1194_v52_: _dafny.Seq
                d_1194_v52_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1195_v53_: D0
                d_1195_v53_ = D0_DC1(len(_dafny.Map({d_1160_v31_: d_1125_i0_})))
                d_1196_v54_: _dafny.Map
                d_1196_v54_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, (d_1194_v52_)[default__.safeIndex(d_1125_i0_, len(d_1194_v52_))], (self).fm5(d_1125_i0_, p0, _dafny.Map({d_1160_v31_: d_1195_v53_}), p0, globalState), p0, p0]): d_1123_v0_})
                d_1197_v55_: _dafny.Set
                d_1197_v55_ = _dafny.Set({p0})
                d_1198_v56_: _dafny.Seq
                d_1198_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0, p0})])
                d_1196_v54_ = (default__.fm56(False, globalState) if ((d_1198_v56_)[default__.safeIndex(d_1125_i0_, len(d_1198_v56_))]).ispropersubset(d_1197_v55_) else (d_1196_v54_ if p0 else d_1196_v54_))
                d_1199_v57_: _dafny.Array
                def lambda49_(d_1200_p0_):
                    def lambda50_(d_1201_i8_):
                        return d_1200_p0_

                    return lambda50_

                init26_ = lambda49_(p0)
                nw191_ = _dafny.Array(None, 15)
                for i0_26_ in range(nw191_.length(0)):
                    nw191_[i0_26_] = init26_(i0_26_)
                d_1199_v57_ = nw191_
                d_1202_v58_: _dafny.MultiSet
                d_1202_v58_ = _dafny.MultiSet([p0])
                index191_ = default__.safeIndex(820, (d_1199_v57_).length(0))
                (d_1199_v57_)[index191_] = not((d_1202_v58_).ispropersubset(d_1202_v58_))
                index192_ = default__.safeIndex(820, (d_1199_v57_).length(0))
                (d_1199_v57_)[index192_] = not((d_1194_v52_)[default__.safeIndex((d_1123_v0_) + (len(default__.fm57(p0, p0, p0, globalState))), len(d_1194_v52_))])
                d_1203_v59_: _dafny.Array
                nw192_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_1203_v59_ = nw192_
                d_1203_v59_ = d_1203_v59_
            elif True:
                d_1204___mcc_h4_ = source13_.cf44
                d_1205_cf44_ = d_1204___mcc_h4_
                d_1206_v60_: _dafny.Map
                d_1206_v60_ = _dafny.Map({d_1125_i0_: (d_1123_v0_) == (d_1125_i0_)})
                d_1206_v60_ = d_1206_v60_
                d_1207_v61_: C2
                nw193_ = C2()
                nw193_.ctor__()
                d_1207_v61_ = nw193_
                d_1208_v62_: _dafny.MultiSet
                d_1208_v62_ = _dafny.MultiSet([p0])
                r0 = ((d_1208_v62_).isdisjoint(d_1208_v62_)) or (p0)
                d_1209_v63_: _dafny.Array
                def lambda51_(d_1210_p0_):
                    def lambda52_(d_1211_i9_):
                        return not(d_1210_p0_)

                    return lambda52_

                init27_ = lambda51_(p0)
                nw194_ = _dafny.Array(None, 14)
                for i0_27_ in range(nw194_.length(0)):
                    nw194_[i0_27_] = init27_(i0_27_)
                d_1209_v63_ = nw194_
                d_1212_v64_: _dafny.Map
                d_1212_v64_ = _dafny.Map({d_1123_v0_: d_1209_v63_})
                d_1213_v65_: _dafny.Map
                d_1213_v65_ = _dafny.Map({(d_1126_v2_) + (d_1126_v2_): d_1212_v64_})
                d_1213_v65_ = (d_1213_v65_).set(d_1126_v2_, d_1212_v64_)
            r3 = p0
        d_1214_v66_: _dafny.Seq
        d_1214_v66_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1215_v67_: _dafny.Set
        d_1215_v67_ = _dafny.Set({980})
        d_1216_v69_: _dafny.Array
        nw195_ = _dafny.Array(None, 8)
        nw195_[int(0)] = p0
        nw195_[int(1)] = p0
        nw195_[int(2)] = p0
        nw195_[int(3)] = True
        nw195_[int(4)] = False
        nw195_[int(5)] = p0
        nw195_[int(6)] = p0
        nw195_[int(7)] = p0
        d_1216_v69_ = nw195_
        d_1217_v70_: _dafny.MultiSet
        d_1217_v70_ = _dafny.MultiSet([d_1216_v69_])
        d_1218_v71_: D4
        d_1218_v71_ = D4_DC11(d_1123_v0_, p0, p0, p0, not(p0))
        def iife115_():
            coll51_ = _dafny.Set()
            compr_51_: int
            for compr_51_ in _dafny.IntegerRange(445, 117):
                d_1219_v68_: int = compr_51_
                if ((445) <= (d_1219_v68_)) and ((d_1219_v68_) < (117)):
                    coll51_ = coll51_.union(_dafny.Set([(d_1219_v68_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjfgxccy"))))]))
            return _dafny.Set(coll51_)
        rhs128_ = not(((d_1215_v67_).intersection(iife115_()
        )).issubset(d_1215_v67_))
        rhs129_ = p0
        rhs130_ = (d_1217_v70_) != ((d_1217_v70_).intersection(d_1217_v70_))
        rhs131_ = (not (p0) or ((d_1218_v71_).cf23)) and (p0)
        rhs132_ = default__.fm38(p0, d_1123_v0_, p0, globalState)
        r0 = rhs128_
        r0 = rhs129_
        r3 = rhs130_
        r3 = rhs131_
        d_1214_v66_ = rhs132_
        d_1220_v72_: bool
        d_1220_v72_ = p0
        d_1221_v73_: _dafny.MultiSet
        d_1221_v73_ = _dafny.MultiSet([p0])
        d_1222_v74_: _dafny.Map
        d_1222_v74_ = _dafny.Map({d_1220_v72_: d_1221_v73_})
        d_1222_v74_ = (d_1222_v74_).set(d_1220_v72_, (_dafny.MultiSet([p0, p0])).intersection(d_1221_v73_))
        d_1223_v75_: _dafny.Map
        d_1223_v75_ = _dafny.Map({(d_1123_v0_) == (d_1123_v0_): len(d_1214_v66_)})
        d_1223_v75_ = d_1223_v75_
        d_1224_v76_: _dafny.Array
        def lambda53_(d_1225_v0_):
            def lambda54_(d_1226_i10_):
                return default__.safeModuloInt(d_1226_i10_, d_1225_v0_)

            return lambda54_

        init28_ = lambda53_(d_1123_v0_)
        nw196_ = _dafny.Array(None, 13)
        for i0_28_ in range(nw196_.length(0)):
            nw196_[i0_28_] = init28_(i0_28_)
        d_1224_v76_ = nw196_
        d_1224_v76_ = d_1224_v76_
        d_1227_v77_: _dafny.Seq
        d_1227_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtblq"))
        d_1228_v78_: D9
        d_1228_v78_ = D9_DC21(d_1221_v73_)
        d_1229_v79_: D15
        d_1229_v79_ = D15_DC39(d_1227_v77_, d_1228_v78_)
        pat_let_tv33_ = p0
        pat_let_tv34_ = p0
        pat_let_tv35_ = p0
        pat_let_tv36_ = p0
        pat_let_tv37_ = d_1228_v78_
        def lambda55_(source14_):
            if source14_.is_DC38:
                d_1230___mcc_h5_ = source14_.cf71
                d_1231_cf71_ = d_1230___mcc_h5_
                return pat_let_tv33_
            elif source14_.is_DC39:
                d_1232___mcc_h6_ = source14_.cf72
                d_1233___mcc_h7_ = source14_.cf73
                d_1234_cf73_ = d_1233___mcc_h7_
                d_1235_cf72_ = d_1232___mcc_h6_
                return pat_let_tv34_
            elif source14_.is_DC37:
                d_1236___mcc_h8_ = source14_.cf70
                d_1237_cf70_ = d_1236___mcc_h8_
                return pat_let_tv35_
            elif True:
                d_1238___mcc_h9_ = source14_.cf74
                d_1239_cf74_ = d_1238___mcc_h9_
                return pat_let_tv36_

        def iife116_(_pat_let32_0):
            def iife117_(d_1240_dt__update__tmp_h0_):
                def iife118_(_pat_let33_0):
                    def iife119_(d_1241_dt__update_hcf73_h0_):
                        return D15_DC39((d_1240_dt__update__tmp_h0_).cf72, d_1241_dt__update_hcf73_h0_)
                    return iife119_(_pat_let33_0)
                return iife118_(pat_let_tv37_)
            return iife117_(_pat_let32_0)
        r0 = lambda55_(iife116_(d_1229_v79_))
        r1 = d_1123_v0_
        r2 = d_1123_v0_
        r3 = p0
        return r0, r1, r2, r3


class C4(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, p2, p3, globalState):
        return not(not(False))

    def m1(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        d_1242_v0_: _dafny.Array
        def lambda56_(d_1243_i0_):
            return True

        init29_ = lambda56_
        nw197_ = _dafny.Array(None, 24)
        for i0_29_ in range(nw197_.length(0)):
            nw197_[i0_29_] = init29_(i0_29_)
        d_1242_v0_ = nw197_
        d_1244_v1_: bool
        d_1244_v1_ = True
        index193_ = default__.safeIndex(663, (d_1242_v0_).length(0))
        (d_1242_v0_)[index193_] = d_1244_v1_
        index194_ = default__.safeIndex(663, (d_1242_v0_).length(0))
        (d_1242_v0_)[index194_] = True
        if d_1244_v1_:
            d_1245_v2_: _dafny.Set
            d_1245_v2_ = _dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
            d_1246_v3_: _dafny.Array
            nw198_ = _dafny.Array(None, 14)
            nw198_[int(0)] = d_1245_v2_
            nw198_[int(1)] = d_1245_v2_
            nw198_[int(2)] = _dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], True, d_1244_v1_})
            nw198_[int(3)] = d_1245_v2_
            nw198_[int(4)] = d_1245_v2_
            nw198_[int(5)] = d_1245_v2_
            nw198_[int(6)] = d_1245_v2_
            nw198_[int(7)] = d_1245_v2_
            nw198_[int(8)] = d_1245_v2_
            nw198_[int(9)] = _dafny.Set({d_1244_v1_})
            nw198_[int(10)] = d_1245_v2_
            nw198_[int(11)] = _dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
            nw198_[int(12)] = d_1245_v2_
            nw198_[int(13)] = d_1245_v2_
            d_1246_v3_ = nw198_
            d_1247_v4_: D0
            d_1247_v4_ = D0_DC0(d_1246_v3_, p0)
            d_1248_v5_: _dafny.Map
            d_1248_v5_ = _dafny.Map({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]: d_1247_v4_})
            d_1248_v5_ = (d_1248_v5_).set(d_1244_v1_, D0_DC0(d_1246_v3_, default__.fm20(not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]), globalState)))
            index195_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            (d_1242_v0_)[index195_] = (default__.fm20(d_1244_v1_, globalState)) != ((0) - (len(_dafny.Map({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]: d_1244_v1_}))))
            d_1249_v6_: _dafny.Map
            d_1249_v6_ = _dafny.Map({p0: (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
            d_1250_v7_: _dafny.Map
            d_1250_v7_ = _dafny.Map({False: d_1249_v6_})
            d_1250_v7_ = (d_1250_v7_).set((True) == (not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])), (d_1249_v6_) | (default__.fm21(globalState)))
            index196_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            (d_1242_v0_)[index196_] = d_1244_v1_
            d_1251_v8_: _dafny.Array
            nw199_ = _dafny.Array(int(0), 14)
            d_1251_v8_ = nw199_
            d_1252_v9_: _dafny.MultiSet
            d_1252_v9_ = _dafny.MultiSet([d_1251_v8_, d_1251_v8_])
            index197_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            (d_1242_v0_)[index197_] = ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] if not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) else not((p0) <= ((d_1252_v9_).cardinality)))
        elif True:
            d_1253_v10_: _dafny.Array
            nw200_ = _dafny.Array(_dafny.Seq({}), 23)
            d_1253_v10_ = nw200_
            def lambda57_(d_1254_p0_):
                def lambda58_(d_1255_i1_):
                    return _dafny.SeqWithoutIsStrInference([d_1254_p0_, d_1254_p0_])

                return lambda58_

            init30_ = lambda57_(p0)
            nw201_ = _dafny.Array(None, 8)
            for i0_30_ in range(nw201_.length(0)):
                nw201_[i0_30_] = init30_(i0_30_)
            d_1253_v10_ = nw201_
            d_1256_v11_: _dafny.Array
            nw202_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1256_v11_ = nw202_
            d_1257_v12_: D2
            d_1257_v12_ = D2_DC4(p0, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], d_1256_v11_)
            index198_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            rhs133_ = d_1244_v1_
            rhs134_ = (d_1257_v12_).cf6
            lhs72_ = d_1242_v0_
            lhs73_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            lhs72_[lhs73_] = rhs133_
            d_1244_v1_ = rhs134_
            d_1258_v13_: _dafny.MultiSet
            d_1258_v13_ = _dafny.MultiSet([_dafny.CodePoint('t')])
            d_1259_v14_: _dafny.Map
            d_1259_v14_ = _dafny.Map({d_1258_v13_: p0})
            d_1259_v14_ = (d_1259_v14_).set((d_1258_v13_) - (d_1258_v13_), 474)
            d_1260_v15_: C0
            nw203_ = C0()
            nw203_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
            d_1260_v15_ = nw203_
            d_1261_v16_: _dafny.Set
            d_1261_v16_ = _dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
            d_1262_v17_: _dafny.Map
            d_1262_v17_ = _dafny.Map({d_1242_v0_: d_1261_v16_})
            d_1263_v18_: _dafny.Array
            nw204_ = _dafny.Array(None, 16)
            nw204_[int(0)] = (d_1261_v16_) - (d_1261_v16_)
            nw204_[int(1)] = d_1261_v16_
            nw204_[int(2)] = (_dafny.Set({True})).intersection(d_1261_v16_)
            nw204_[int(3)] = d_1261_v16_
            nw204_[int(4)] = d_1261_v16_
            nw204_[int(5)] = d_1261_v16_
            nw204_[int(6)] = d_1261_v16_
            nw204_[int(7)] = d_1261_v16_
            nw204_[int(8)] = d_1261_v16_
            nw204_[int(9)] = _dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], True})
            nw204_[int(10)] = d_1261_v16_
            nw204_[int(11)] = (d_1261_v16_) - (d_1261_v16_)
            nw204_[int(12)] = d_1261_v16_
            nw204_[int(13)] = (_dafny.Set({True, d_1244_v1_, d_1244_v1_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]}) if True else ((d_1262_v17_)[d_1242_v0_] if (d_1242_v0_) in (d_1262_v17_) else d_1261_v16_))
            nw204_[int(14)] = (d_1261_v16_).intersection(d_1261_v16_)
            nw204_[int(15)] = d_1261_v16_
            d_1263_v18_ = nw204_
            d_1263_v18_ = d_1263_v18_
        d_1264_v19_: _dafny.Array
        def lambda59_(d_1265_v1_, d_1266_v0_):
            def lambda60_(d_1267_i2_):
                return _dafny.Set({d_1265_v1_, d_1265_v1_, (d_1266_v0_)[default__.safeIndex(663, (d_1266_v0_).length(0))], d_1265_v1_})

            return lambda60_

        init31_ = lambda59_(d_1244_v1_, d_1242_v0_)
        nw205_ = _dafny.Array(None, 26)
        for i0_31_ in range(nw205_.length(0)):
            nw205_[i0_31_] = init31_(i0_31_)
        d_1264_v19_ = nw205_
        d_1268_v20_: D0
        d_1268_v20_ = D0_DC0(d_1264_v19_, p0)
        d_1269_v21_: _dafny.Seq
        d_1269_v21_ = _dafny.SeqWithoutIsStrInference([d_1268_v20_])
        d_1270_v22_: D4
        d_1270_v22_ = D4_DC8(d_1269_v21_)
        if ((d_1270_v22_).cf13) <= (d_1269_v21_):
            d_1271_v23_: _dafny.Map
            d_1271_v23_ = _dafny.Map({689: d_1244_v1_})
            d_1272_v24_: _dafny.Seq
            d_1272_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuqe"))
            d_1273_v25_: _dafny.Seq
            d_1273_v25_ = _dafny.SeqWithoutIsStrInference([len(d_1272_v24_)])
            d_1274_v26_: _dafny.Map
            d_1274_v26_ = _dafny.Map({d_1244_v1_: len(_dafny.Map({len(d_1273_v25_): p0}))})
            d_1275_v27_: _dafny.MultiSet
            d_1275_v27_ = _dafny.MultiSet([d_1244_v1_])
            d_1276_v28_: _dafny.MultiSet
            d_1276_v28_ = _dafny.MultiSet([p0])
            d_1277_v29_: _dafny.Map
            d_1277_v29_ = _dafny.Map({(0) - (p0): 159})
            d_1278_v30_: _dafny.Seq
            d_1278_v30_ = _dafny.SeqWithoutIsStrInference([d_1277_v29_])
            d_1279_v31_: _dafny.MultiSet
            d_1279_v31_ = _dafny.MultiSet([d_1242_v0_])
            d_1280_v32_: _dafny.Array
            nw206_ = _dafny.Array(None, 22)
            nw206_[int(0)] = p0
            nw206_[int(1)] = len(d_1271_v23_)
            nw206_[int(2)] = p0
            nw206_[int(3)] = len(d_1274_v26_)
            nw206_[int(4)] = len(d_1272_v24_)
            nw206_[int(5)] = p0
            nw206_[int(6)] = p0
            nw206_[int(7)] = len(d_1272_v24_)
            nw206_[int(8)] = (d_1275_v27_).cardinality
            nw206_[int(9)] = p0
            nw206_[int(10)] = (d_1273_v25_)[default__.safeIndex(p0, len(d_1273_v25_))]
            nw206_[int(11)] = -17
            nw206_[int(12)] = (d_1273_v25_)[default__.safeIndex(p0, len(d_1273_v25_))]
            nw206_[int(13)] = p0
            nw206_[int(14)] = (d_1276_v28_).cardinality
            nw206_[int(15)] = p0
            nw206_[int(16)] = len((d_1278_v30_)[default__.safeIndex(p0, len(d_1278_v30_))])
            nw206_[int(17)] = ((d_1279_v31_)[d_1242_v0_] if (d_1242_v0_) in (d_1279_v31_) else p0)
            nw206_[int(18)] = p0
            nw206_[int(19)] = p0
            nw206_[int(20)] = len(default__.fm22(False, p0, (d_1275_v27_).cardinality, globalState))
            nw206_[int(21)] = p0
            d_1280_v32_ = nw206_
            d_1281_v33_: _dafny.Seq
            d_1281_v33_ = _dafny.SeqWithoutIsStrInference([d_1280_v32_, d_1280_v32_, d_1280_v32_, d_1280_v32_, d_1280_v32_])
            d_1282_v34_: _dafny.Map
            d_1282_v34_ = _dafny.Map({(d_1281_v33_)[default__.safeIndex(p0, len(d_1281_v33_))]: p0})
            d_1282_v34_ = (d_1282_v34_).set((d_1280_v32_ if d_1244_v1_ else d_1280_v32_), p0)
            d_1283_v35_: C0
            nw207_ = C0()
            nw207_.ctor__(d_1272_v24_)
            d_1283_v35_ = nw207_
            d_1284_v36_: _dafny.Array
            nw208_ = _dafny.Array(None, 3)
            nw208_[int(0)] = d_1281_v33_
            nw208_[int(1)] = (d_1281_v33_) + (d_1281_v33_)
            nw208_[int(2)] = d_1281_v33_
            d_1284_v36_ = nw208_
            index199_ = default__.safeIndex(233, (d_1284_v36_).length(0))
            (d_1284_v36_)[index199_] = _dafny.SeqWithoutIsStrInference([d_1280_v32_])
            index200_ = default__.safeIndex(233, (d_1284_v36_).length(0))
            (d_1284_v36_)[index200_] = _dafny.SeqWithoutIsStrInference([d_1280_v32_])
            d_1285_v37_: _dafny.Map
            d_1285_v37_ = _dafny.Map({d_1244_v1_: (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
            source15_ = D4_DC11((len(d_1273_v25_)) + (p0), (not(False)) or (((d_1285_v37_)[True] if (True) in (d_1285_v37_) else False)), d_1244_v1_, (p0) >= (-703), d_1244_v1_)
            if source15_.is_DC9:
                d_1286___mcc_h0_ = source15_.cf14
                d_1287___mcc_h1_ = source15_.cf15
                d_1288___mcc_h2_ = source15_.cf16
                d_1289___mcc_h3_ = source15_.cf17
                d_1290___mcc_h4_ = source15_.cf18
                d_1291_cf18_ = d_1290___mcc_h4_
                d_1292_cf17_ = d_1289___mcc_h3_
                d_1293_cf16_ = d_1288___mcc_h2_
                d_1294_cf15_ = d_1287___mcc_h1_
                d_1295_cf14_ = d_1286___mcc_h0_
                d_1291_cf18_ = ((d_1293_cf16_) + ((0) - (p0))) + ((d_1293_cf16_ if False else d_1291_cf18_))
                d_1296_v38_: _dafny.Array
                nw209_ = _dafny.Array(None, 9)
                d_1296_v38_ = nw209_
                index201_ = default__.safeIndex(476, (d_1296_v38_).length(0))
                (d_1296_v38_)[index201_] = d_1283_v35_
                index202_ = default__.safeIndex(476, (d_1296_v38_).length(0))
                (d_1296_v38_)[index202_] = d_1283_v35_
                d_1297_v39_: _dafny.Map
                d_1297_v39_ = _dafny.Map({_dafny.CodePoint('c'): D0_DC1(486)})
                r1 = (self).fm5(d_1293_cf16_, d_1244_v1_, (d_1297_v39_) | (d_1297_v39_), ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] if (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] else not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])), globalState)
                index203_ = default__.safeIndex(663, (d_1242_v0_).length(0))
                (d_1242_v0_)[index203_] = (not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) if (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] else (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])
            elif source15_.is_DC10:
                d_1298___mcc_h5_ = source15_.cf19
                d_1299___mcc_h6_ = source15_.cf20
                d_1300___mcc_h7_ = source15_.cf21
                d_1301_cf21_ = d_1300___mcc_h7_
                d_1302_cf20_ = d_1299___mcc_h6_
                d_1303_cf19_ = d_1298___mcc_h5_
                d_1304_v40_: str
                d_1304_v40_ = _dafny.CodePoint('k')
                d_1274_v26_ = (d_1274_v26_).set((d_1304_v40_) not in (d_1283_v35_.f8), (d_1303_cf19_) + (len(d_1273_v25_)))
                d_1305_v41_: _dafny.Set
                d_1305_v41_ = _dafny.Set({d_1272_v24_, _dafny.SeqWithoutIsStrInference([d_1304_v40_, d_1304_v40_])})
                d_1306_v42_: D2
                d_1306_v42_ = D2_DC3(d_1305_v41_)
                d_1307_v43_: D2
                d_1307_v43_ = D2_DC6(d_1306_v42_)
                d_1308_v44_: _dafny.Map
                d_1308_v44_ = _dafny.Map({d_1307_v43_: d_1283_v35_.f8})
                d_1308_v44_ = d_1308_v44_
                d_1309_v45_: D0
                d_1309_v45_ = D0_DC1((d_1275_v27_).cardinality)
                d_1310_v46_: _dafny.Map
                d_1310_v46_ = _dafny.Map({d_1304_v40_: d_1309_v45_})
                d_1301_cf21_ = default__.safeModuloInt(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_1244_v1_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], ((d_1285_v37_)[(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]] if ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) in (d_1285_v37_) else d_1244_v1_)])), d_1302_cf20_), ((d_1274_v26_)[(self).fm5(d_1301_cf21_, d_1244_v1_, d_1310_v46_, d_1244_v1_, globalState)] if ((self).fm5(d_1301_cf21_, d_1244_v1_, d_1310_v46_, d_1244_v1_, globalState)) in (d_1274_v26_) else len(_dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]}))))
                d_1274_v26_ = ((d_1274_v26_) | (d_1274_v26_)) | ((_dafny.Map({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]: -533})) | (d_1274_v26_))
            elif source15_.is_DC11:
                d_1311___mcc_h8_ = source15_.cf22
                d_1312___mcc_h9_ = source15_.cf23
                d_1313___mcc_h10_ = source15_.cf24
                d_1314___mcc_h11_ = source15_.cf25
                d_1315___mcc_h12_ = source15_.cf26
                d_1316_cf26_ = d_1315___mcc_h12_
                d_1317_cf25_ = d_1314___mcc_h11_
                d_1318_cf24_ = d_1313___mcc_h10_
                d_1319_cf23_ = d_1312___mcc_h9_
                d_1320_cf22_ = d_1311___mcc_h8_
                d_1321_v47_: _dafny.Seq
                d_1321_v47_ = _dafny.SeqWithoutIsStrInference([d_1318_cf24_])
                d_1322_v48_: _dafny.Seq
                d_1322_v48_ = _dafny.SeqWithoutIsStrInference([d_1321_v47_, default__.fm23(d_1244_v1_, globalState), (d_1321_v47_).set(default__.safeIndex(d_1320_cf22_, len(d_1321_v47_)), d_1317_cf25_)])
                d_1322_v48_ = d_1322_v48_
                d_1323_v49_: _dafny.Seq
                d_1323_v49_ = _dafny.SeqWithoutIsStrInference([d_1242_v0_])
                d_1324_v50_: bool
                d_1325_v51_: int
                out37_: bool
                out38_: int
                out37_, out38_ = (self).m10(d_1320_cf22_, d_1320_cf22_, not(not (d_1244_v1_) or ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])), (d_1323_v49_)[default__.safeIndex(p0, len(d_1323_v49_))], globalState)
                d_1324_v50_ = out37_
                d_1325_v51_ = out38_
                d_1320_cf22_ = d_1320_cf22_
                d_1326_v52_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1326_v52_ = nw210_
                d_1326_v52_ = d_1326_v52_
            elif source15_.is_DC8:
                d_1327___mcc_h13_ = source15_.cf13
                d_1328_cf13_ = d_1327___mcc_h13_
                d_1329_v53_: int
                d_1329_v53_ = -892
                d_1329_v53_ = (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1330_i3_ in range(default__.abs(266))]))))
                d_1331_v54_: str
                d_1331_v54_ = _dafny.CodePoint('f')
                d_1331_v54_ = (d_1331_v54_ if d_1244_v1_ else d_1331_v54_)
                d_1332_v55_: D2
                d_1332_v55_ = D2_DC6(D2_DC5((0) - (d_1329_v53_), d_1273_v25_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]))
                d_1333_v56_: _dafny.Seq
                d_1333_v56_ = _dafny.SeqWithoutIsStrInference([d_1332_v55_])
                d_1334_v57_: bool
                d_1334_v57_ = (d_1332_v55_) in (d_1333_v56_)
                d_1335_v58_: _dafny.Map
                d_1335_v58_ = _dafny.Map({d_1331_v54_: d_1242_v0_})
                d_1336_v59_: _dafny.Set
                d_1336_v59_ = _dafny.Set({d_1283_v35_, d_1283_v35_})
                d_1337_v60_: D4
                d_1337_v60_ = D4_DC9((default__.fm24((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], globalState)).set((0) - ((0) - (d_1329_v53_)), default__.abs(d_1329_v53_)), _dafny.CodePoint('d'), (len(_dafny.Set({d_1244_v1_}))) + ((0) - (len(d_1271_v23_))), d_1336_v59_, (d_1329_v53_) - (d_1329_v53_))
                d_1338_v61_: D0
                d_1338_v61_ = D0_DC1(p0)
                d_1339_v62_: _dafny.Map
                d_1339_v62_ = _dafny.Map({d_1331_v54_: d_1338_v61_})
                index204_ = default__.safeIndex(663, (d_1242_v0_).length(0))
                rhs135_ = d_1334_v57_
                rhs136_ = (True) or ((self).fm5(d_1329_v53_, True, d_1339_v62_, not(not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])), globalState))
                rhs137_ = d_1335_v58_
                rhs138_ = d_1337_v60_
                lhs74_ = d_1242_v0_
                lhs75_ = default__.safeIndex(663, (d_1242_v0_).length(0))
                d_1334_v57_ = rhs135_
                lhs74_[lhs75_] = rhs136_
                d_1335_v58_ = rhs137_
                d_1337_v60_ = rhs138_
                d_1340_v63_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1340_v63_ = nw211_
                d_1341_v64_: _dafny.Array
                nw212_ = _dafny.Array(None, 17)
                nw212_[int(0)] = d_1280_v32_
                nw212_[int(1)] = d_1280_v32_
                nw212_[int(2)] = d_1280_v32_
                nw212_[int(3)] = d_1280_v32_
                nw212_[int(4)] = d_1280_v32_
                nw212_[int(5)] = d_1280_v32_
                nw212_[int(6)] = d_1280_v32_
                nw212_[int(7)] = d_1280_v32_
                nw212_[int(8)] = d_1280_v32_
                nw212_[int(9)] = d_1280_v32_
                nw212_[int(10)] = d_1280_v32_
                nw212_[int(11)] = d_1280_v32_
                nw212_[int(12)] = d_1280_v32_
                nw212_[int(13)] = d_1280_v32_
                nw212_[int(14)] = d_1280_v32_
                nw212_[int(15)] = d_1280_v32_
                nw212_[int(16)] = d_1280_v32_
                d_1341_v64_ = nw212_
                index205_ = default__.safeIndex(432, (d_1340_v63_).length(0))
                (d_1340_v63_)[index205_] = d_1341_v64_
                index206_ = default__.safeIndex(432, (d_1340_v63_).length(0))
                (d_1340_v63_)[index206_] = d_1341_v64_
            elif True:
                d_1342___mcc_h14_ = source15_.cf27
                d_1343_cf27_ = d_1342___mcc_h14_
                d_1244_v1_ = (((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) or ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) if d_1244_v1_ else (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])
                d_1280_v32_ = d_1280_v32_
                index207_ = default__.safeIndex(663, (d_1242_v0_).length(0))
                (d_1242_v0_)[index207_] = d_1244_v1_
                d_1344_v65_: int
                d_1344_v65_ = 226
                d_1345_v66_: _dafny.Array
                nw213_ = _dafny.Array(_dafny.Map({}), 28)
                d_1345_v66_ = nw213_
                d_1346_v67_: _dafny.Map
                d_1346_v67_ = _dafny.Map({d_1244_v1_: d_1242_v0_})
                index208_ = default__.safeIndex(390, (d_1345_v66_).length(0))
                (d_1345_v66_)[index208_] = d_1346_v67_
                d_1347_v68_: _dafny.Array
                def lambda61_(d_1348_p0_):
                    def lambda62_(d_1349_i4_):
                        return _dafny.Map({_dafny.CodePoint('f'): d_1348_p0_})

                    return lambda62_

                init32_ = lambda61_(p0)
                nw214_ = _dafny.Array(None, 2)
                for i0_32_ in range(nw214_.length(0)):
                    nw214_[i0_32_] = init32_(i0_32_)
                d_1347_v68_ = nw214_
                d_1350_v69_: _dafny.Map
                d_1350_v69_ = _dafny.Map({d_1344_v65_: d_1347_v68_})
                d_1351_v70_: _dafny.Array
                d_1351_v70_ = d_1347_v68_
                d_1352_v71_: _dafny.Array
                nw215_ = _dafny.Array(None, 26)
                nw215_[int(0)] = d_1347_v68_
                nw215_[int(1)] = d_1347_v68_
                nw215_[int(2)] = d_1347_v68_
                nw215_[int(3)] = d_1347_v68_
                nw215_[int(4)] = d_1347_v68_
                nw215_[int(5)] = d_1347_v68_
                nw215_[int(6)] = d_1347_v68_
                nw215_[int(7)] = d_1347_v68_
                nw215_[int(8)] = ((d_1350_v69_)[len(d_1273_v25_)] if (len(d_1273_v25_)) in (d_1350_v69_) else d_1347_v68_)
                nw215_[int(9)] = (d_1351_v70_)
                nw215_[int(10)] = d_1347_v68_
                nw215_[int(11)] = d_1347_v68_
                nw215_[int(12)] = d_1347_v68_
                nw215_[int(13)] = d_1347_v68_
                nw215_[int(14)] = d_1347_v68_
                nw215_[int(15)] = d_1347_v68_
                nw215_[int(16)] = d_1347_v68_
                nw215_[int(17)] = d_1347_v68_
                nw215_[int(18)] = d_1347_v68_
                nw215_[int(19)] = d_1347_v68_
                nw215_[int(20)] = d_1347_v68_
                nw215_[int(21)] = d_1347_v68_
                nw215_[int(22)] = d_1347_v68_
                nw215_[int(23)] = d_1347_v68_
                nw215_[int(24)] = d_1347_v68_
                nw215_[int(25)] = d_1347_v68_
                d_1352_v71_ = nw215_
                index209_ = default__.safeIndex(0, (d_1352_v71_).length(0))
                (d_1352_v71_)[index209_] = d_1347_v68_
                d_1353_v72_: str
                d_1353_v72_ = _dafny.CodePoint('p')
                d_1354_v73_: _dafny.Set
                d_1354_v73_ = _dafny.Set({d_1244_v1_, d_1244_v1_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
                d_1355_v74_: _dafny.Seq
                d_1355_v74_ = _dafny.SeqWithoutIsStrInference([d_1354_v73_])
                d_1356_v75_: _dafny.Seq
                d_1356_v75_ = _dafny.SeqWithoutIsStrInference([d_1346_v67_])
                index210_ = default__.safeIndex(390, (d_1345_v66_).length(0))
                index211_ = default__.safeIndex(0, (d_1352_v71_).length(0))
                rhs139_ = not(False)
                rhs140_ = len(d_1355_v74_)
                rhs141_ = (d_1346_v67_) | ((d_1356_v75_)[default__.safeIndex(d_1344_v65_, len(d_1356_v75_))])
                rhs142_ = d_1347_v68_
                rhs143_ = _dafny.CodePoint('d')
                lhs76_ = d_1345_v66_
                lhs77_ = default__.safeIndex(390, (d_1345_v66_).length(0))
                lhs78_ = d_1352_v71_
                lhs79_ = default__.safeIndex(0, (d_1352_v71_).length(0))
                r0 = rhs139_
                d_1344_v65_ = rhs140_
                lhs76_[lhs77_] = rhs141_
                lhs78_[lhs79_] = rhs142_
                d_1353_v72_ = rhs143_
            d_1357_v76_: _dafny.Array
            d_1358_v77_: _dafny.Array
            d_1359_v78_: bool
            out39_: _dafny.Array
            out40_: _dafny.Array
            out41_: bool
            out39_, out40_, out41_ = (self).m9(globalState)
            d_1357_v76_ = out39_
            d_1358_v77_ = out40_
            d_1359_v78_ = out41_
        elif True:
            d_1360_v79_: _dafny.Array
            d_1361_v80_: _dafny.Array
            d_1362_v81_: bool
            out42_: _dafny.Array
            out43_: _dafny.Array
            out44_: bool
            out42_, out43_, out44_ = (self).m9(globalState)
            d_1360_v79_ = out42_
            d_1361_v80_ = out43_
            d_1362_v81_ = out44_
            d_1363_v83_: _dafny.Set
            d_1363_v83_ = _dafny.Set({-990, p0})
            d_1364_v84_: _dafny.Set
            d_1364_v84_ = _dafny.Set({len(d_1363_v83_), p0})
            d_1365_v85_: _dafny.MultiSet
            d_1365_v85_ = _dafny.MultiSet([p0, len(d_1363_v83_)])
            d_1366_v86_: str
            d_1366_v86_ = _dafny.CodePoint('a')
            d_1367_v87_: _dafny.Map
            d_1367_v87_ = _dafny.Map({(d_1365_v85_).cardinality: d_1366_v86_})
            d_1368_v88_: _dafny.Seq
            def iife120_():
                coll52_ = _dafny.Map()
                compr_52_: int
                for compr_52_ in _dafny.IntegerRange(-592, 863):
                    d_1369_v82_: int = compr_52_
                    if ((-592) <= (d_1369_v82_)) and ((d_1369_v82_) < (863)):
                        coll52_[(d_1369_v82_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmetdhseu"))))] = _dafny.CodePoint('y')
                return _dafny.Map(coll52_)
            d_1368_v88_ = _dafny.SeqWithoutIsStrInference([(iife120_()
            ) | (d_1367_v87_)])
            d_1368_v88_ = d_1368_v88_
            d_1370_v89_: int
            d_1370_v89_ = 768
            d_1371_v90_: _dafny.Seq
            d_1371_v90_ = _dafny.SeqWithoutIsStrInference([not(d_1244_v1_)])
            d_1372_v91_: _dafny.Seq
            d_1372_v91_ = _dafny.SeqWithoutIsStrInference([d_1371_v90_])
            d_1373_v92_: _dafny.Seq
            d_1373_v92_ = _dafny.SeqWithoutIsStrInference([732, len((d_1371_v90_) + ((d_1372_v91_)[default__.safeIndex(p0, len(d_1372_v91_))])), p0, p0, d_1370_v89_])
            d_1374_v93_: _dafny.Map
            d_1374_v93_ = _dafny.Map({d_1362_v81_: p0})
            d_1375_v94_: _dafny.Seq
            d_1375_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvfkivcdy"))
            rhs144_ = 665
            rhs145_ = ((d_1374_v93_)[d_1244_v1_] if (d_1244_v1_) in (d_1374_v93_) else d_1370_v89_)
            rhs146_ = not ((d_1375_v94_) <= (d_1375_v94_)) or ((890) == (((d_1365_v85_)[p0] if (p0) in (d_1365_v85_) else len(_dafny.Map({p0: p0})))))
            rhs147_ = d_1373_v92_
            d_1370_v89_ = rhs144_
            d_1370_v89_ = rhs145_
            d_1362_v81_ = rhs146_
            d_1373_v92_ = rhs147_
            d_1376_v95_: _dafny.Seq
            d_1376_v95_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1366_v86_ for d_1377_i5_ in range(default__.abs(967))]), ((D6_DC14(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erond")))).cf29).set(default__.safeIndex(p0, len((D6_DC14(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erond")))).cf29)), _dafny.CodePoint('f')), _dafny.SeqWithoutIsStrInference([d_1366_v86_ for d_1378_i6_ in range(default__.abs(523))])])
            d_1379_v96_: C0
            nw216_ = C0()
            nw216_.ctor__((d_1375_v94_) + ((d_1376_v95_)[default__.safeIndex(d_1370_v89_, len(d_1376_v95_))]))
            d_1379_v96_ = nw216_
            d_1380_v97_: _dafny.Map
            d_1380_v97_ = _dafny.Map({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]: d_1362_v81_})
            d_1381_v98_: D2
            d_1381_v98_ = D2_DC5(len(d_1380_v97_), d_1373_v92_, d_1362_v81_)
            d_1244_v1_ = ((d_1381_v98_).cf10 if d_1362_v81_ else not((d_1371_v90_)[default__.safeIndex(d_1370_v89_, len(d_1371_v90_))]))
        d_1382_v99_: _dafny.Set
        d_1382_v99_ = _dafny.Set({p0, p0, p0})
        d_1383_v100_: _dafny.Map
        d_1383_v100_ = _dafny.Map({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]: len(d_1382_v99_)})
        d_1384_v101_: str
        d_1384_v101_ = _dafny.CodePoint('h')
        d_1385_v102_: D0
        d_1385_v102_ = D0_DC1(p0)
        d_1386_v103_: _dafny.Map
        d_1386_v103_ = _dafny.Map({d_1384_v101_: d_1385_v102_})
        if not ((self).fm5(len(d_1383_v100_), d_1244_v1_, d_1386_v103_, d_1244_v1_, globalState)) or ((p0) != (p0)):
            d_1387_v104_: _dafny.Array
            nw217_ = _dafny.Array(int(0), 22)
            d_1387_v104_ = nw217_
            d_1388_v105_: _dafny.Seq
            d_1388_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awvnycoec"))
            index212_ = default__.safeIndex(758, (d_1387_v104_).length(0))
            (d_1387_v104_)[index212_] = len(d_1388_v105_)
            d_1389_v106_: int
            d_1389_v106_ = 125
            index213_ = default__.safeIndex(155, (d_1387_v104_).length(0))
            (d_1387_v104_)[index213_] = p0
            index214_ = default__.safeIndex(758, (d_1387_v104_).length(0))
            index215_ = default__.safeIndex(155, (d_1387_v104_).length(0))
            rhs148_ = d_1389_v106_
            rhs149_ = (d_1389_v106_) - (d_1389_v106_)
            rhs150_ = p0
            lhs80_ = d_1387_v104_
            lhs81_ = default__.safeIndex(758, (d_1387_v104_).length(0))
            lhs82_ = d_1387_v104_
            lhs83_ = default__.safeIndex(155, (d_1387_v104_).length(0))
            lhs80_[lhs81_] = rhs148_
            d_1389_v106_ = rhs149_
            lhs82_[lhs83_] = rhs150_
            d_1390_v107_: _dafny.MultiSet
            d_1390_v107_ = _dafny.MultiSet([d_1242_v0_, d_1242_v0_])
            d_1391_v108_: D4
            d_1391_v108_ = D4_DC10(p0, d_1389_v106_, d_1389_v106_)
            index216_ = default__.safeIndex(758, (d_1387_v104_).length(0))
            rhs151_ = default__.safeModuloInt((((_dafny.MultiSet([d_1242_v0_])).set(d_1242_v0_, default__.abs(p0))) | (d_1390_v107_)).cardinality, d_1389_v106_)
            rhs152_ = default__.fm25(d_1391_v108_, d_1388_v105_, globalState)
            lhs84_ = d_1387_v104_
            lhs85_ = default__.safeIndex(758, (d_1387_v104_).length(0))
            lhs84_[lhs85_] = rhs151_
            d_1388_v105_ = rhs152_
            d_1392_v109_: _dafny.Map
            d_1392_v109_ = _dafny.Map({p0: (d_1387_v104_)[default__.safeIndex(758, (d_1387_v104_).length(0))]})
            d_1393_v110_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_1393_v110_ = nw218_
            d_1394_v111_: D2
            d_1394_v111_ = D2_DC4(((d_1392_v109_)[p0] if (p0) in (d_1392_v109_) else -269), False, d_1393_v110_)
            if (default__.fm24((d_1394_v111_).cf6, globalState)).issubset(_dafny.MultiSet([p0, p0])):
                d_1395_v112_: _dafny.Seq
                d_1395_v112_ = _dafny.SeqWithoutIsStrInference([(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])
                d_1396_v113_: D7
                d_1396_v113_ = D7_DC17(d_1395_v112_)
                d_1397_v114_: _dafny.Array
                nw219_ = _dafny.Array(None, 20)
                nw219_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_1394_v111_).cf6, False])
                nw219_[int(1)] = d_1395_v112_
                nw219_[int(2)] = d_1395_v112_
                nw219_[int(3)] = d_1395_v112_
                nw219_[int(4)] = d_1395_v112_
                nw219_[int(5)] = _dafny.SeqWithoutIsStrInference([(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])
                nw219_[int(6)] = d_1395_v112_
                nw219_[int(7)] = (_dafny.SeqWithoutIsStrInference([(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], d_1244_v1_, d_1244_v1_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])) + ((d_1396_v113_).cf34)
                nw219_[int(8)] = default__.fm23((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], globalState)
                nw219_[int(9)] = (d_1395_v112_) + (_dafny.SeqWithoutIsStrInference([d_1244_v1_]))
                nw219_[int(10)] = d_1395_v112_
                nw219_[int(11)] = d_1395_v112_
                nw219_[int(12)] = d_1395_v112_
                nw219_[int(13)] = d_1395_v112_
                nw219_[int(14)] = (d_1395_v112_) + (d_1395_v112_)
                nw219_[int(15)] = _dafny.SeqWithoutIsStrInference([True, d_1244_v1_, False])
                nw219_[int(16)] = (d_1395_v112_).set(default__.safeIndex(82, len(d_1395_v112_)), (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])
                nw219_[int(17)] = _dafny.SeqWithoutIsStrInference([not(not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])), (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])
                nw219_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_1244_v1_, d_1244_v1_, d_1244_v1_])) + (d_1395_v112_)
                nw219_[int(19)] = (d_1395_v112_) + (d_1395_v112_)
                d_1397_v114_ = nw219_
                index217_ = default__.safeIndex(198, (d_1397_v114_).length(0))
                (d_1397_v114_)[index217_] = (d_1395_v112_) + (d_1395_v112_)
                d_1398_v115_: _dafny.Seq
                d_1398_v115_ = _dafny.SeqWithoutIsStrInference([(d_1387_v104_)[default__.safeIndex(758, (d_1387_v104_).length(0))], len(d_1395_v112_), d_1389_v106_])
                d_1399_v116_: _dafny.Map
                d_1399_v116_ = _dafny.Map({d_1398_v115_: (d_1387_v104_)[default__.safeIndex(758, (d_1387_v104_).length(0))]})
                index218_ = default__.safeIndex(663, (d_1242_v0_).length(0))
                index219_ = default__.safeIndex(758, (d_1387_v104_).length(0))
                index220_ = default__.safeIndex(198, (d_1397_v114_).length(0))
                rhs153_ = not((d_1389_v106_) == (default__.fm20(d_1244_v1_, globalState)))
                rhs154_ = ((d_1399_v116_)[d_1398_v115_] if (d_1398_v115_) in (d_1399_v116_) else (d_1387_v104_)[default__.safeIndex(758, (d_1387_v104_).length(0))])
                rhs155_ = (_dafny.SeqWithoutIsStrInference([d_1244_v1_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([False]))
                lhs86_ = d_1242_v0_
                lhs87_ = default__.safeIndex(663, (d_1242_v0_).length(0))
                lhs88_ = d_1387_v104_
                lhs89_ = default__.safeIndex(758, (d_1387_v104_).length(0))
                lhs90_ = d_1397_v114_
                lhs91_ = default__.safeIndex(198, (d_1397_v114_).length(0))
                lhs86_[lhs87_] = rhs153_
                lhs88_[lhs89_] = rhs154_
                lhs90_[lhs91_] = rhs155_
                d_1400_v117_: C0
                nw220_ = C0()
                nw220_.ctor__(d_1388_v105_)
                d_1400_v117_ = nw220_
                d_1401_v118_: _dafny.Map
                d_1401_v118_ = _dafny.Map({d_1400_v117_: not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])})
                d_1401_v118_ = d_1401_v118_
                d_1402_v119_: C0
                nw221_ = C0()
                nw221_.ctor__((default__.fm25(d_1391_v108_, d_1388_v105_, globalState)) + (d_1388_v105_))
                d_1402_v119_ = nw221_
                d_1403_v120_: _dafny.Set
                d_1403_v120_ = _dafny.Set({d_1244_v1_, not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])})
                d_1404_v121_: _dafny.Seq
                d_1404_v121_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1403_v120_), 767])])
                d_1404_v121_ = ((d_1404_v121_) + (d_1404_v121_)).set(default__.safeIndex(d_1389_v106_, len((d_1404_v121_) + (d_1404_v121_))), _dafny.SeqWithoutIsStrInference([len(d_1392_v109_) for d_1405_i7_ in range(default__.abs(790))]))
                d_1406_v122_: _dafny.Map
                d_1406_v122_ = _dafny.Map({d_1384_v101_: default__.fm20((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], globalState)})
                d_1406_v122_ = (d_1406_v122_).set(d_1384_v101_, p0)
            elif True:
                d_1407_v123_: _dafny.Map
                d_1407_v123_ = _dafny.Map({d_1244_v1_: d_1242_v0_})
                d_1408_v124_: _dafny.Set
                d_1408_v124_ = _dafny.Set({d_1242_v0_, ((d_1407_v123_)[not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])] if (not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])) in (d_1407_v123_) else d_1242_v0_)})
                d_1409_v125_: _dafny.Array
                nw222_ = _dafny.Array(None, 4)
                nw222_[int(0)] = d_1408_v124_
                nw222_[int(1)] = d_1408_v124_
                nw222_[int(2)] = _dafny.Set({d_1242_v0_, d_1242_v0_})
                nw222_[int(3)] = (d_1408_v124_) | (_dafny.Set({d_1242_v0_}))
                d_1409_v125_ = nw222_
                d_1410_v126_: _dafny.Set
                d_1410_v126_ = d_1408_v124_
                index221_ = default__.safeIndex(434, (d_1409_v125_).length(0))
                (d_1409_v125_)[index221_] = ((d_1410_v126_) if d_1244_v1_ else d_1408_v124_)
                d_1411_v127_: _dafny.Map
                d_1411_v127_ = _dafny.Map({False: (d_1408_v124_) | (d_1408_v124_)})
                d_1412_v128_: _dafny.Map
                d_1412_v128_ = _dafny.Map({d_1244_v1_: d_1244_v1_})
                index222_ = default__.safeIndex(434, (d_1409_v125_).length(0))
                (d_1409_v125_)[index222_] = ((d_1411_v127_)[(((d_1412_v128_)[d_1244_v1_] if (d_1244_v1_) in (d_1412_v128_) else False) if d_1244_v1_ else not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]))] if ((((d_1412_v128_)[d_1244_v1_] if (d_1244_v1_) in (d_1412_v128_) else False) if d_1244_v1_ else not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]))) in (d_1411_v127_) else (d_1408_v124_) - (d_1408_v124_))
                d_1413_v129_: C0
                nw223_ = C0()
                nw223_.ctor__(_dafny.SeqWithoutIsStrInference([d_1384_v101_ for d_1414_i8_ in range(default__.abs(445))]))
                d_1413_v129_ = nw223_
                index223_ = default__.safeIndex(758, (d_1387_v104_).length(0))
                (d_1387_v104_)[index223_] = (_dafny.MultiSet([((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) and ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]), (self).fm5(d_1389_v106_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], d_1386_v103_, False, globalState)])).cardinality
                d_1415_v130_: _dafny.MultiSet
                d_1415_v130_ = _dafny.MultiSet([d_1389_v106_, d_1389_v106_])
                d_1416_v131_: D4
                d_1416_v131_ = D4_DC9(d_1415_v130_, d_1384_v101_, d_1389_v106_, _dafny.Set({d_1413_v129_}), d_1389_v106_)
                d_1417_v132_: C0
                nw224_ = C0()
                nw224_.ctor__((d_1388_v105_) + ((d_1413_v129_.f8).set(default__.safeIndex((d_1387_v104_)[default__.safeIndex(758, (d_1387_v104_).length(0))], len(d_1413_v129_.f8)), (d_1416_v131_).cf15)))
                d_1417_v132_ = nw224_
                index224_ = default__.safeIndex(758, (d_1387_v104_).length(0))
                (d_1387_v104_)[index224_] = len(_dafny.Map({not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]): p0}))
            index225_ = default__.safeIndex(758, (d_1387_v104_).length(0))
            (d_1387_v104_)[index225_] = d_1389_v106_
            d_1389_v106_ = p0
        elif True:
            r0 = not(d_1244_v1_)
            d_1244_v1_ = ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] if (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] else not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]))
            d_1418_v133_: int
            d_1418_v133_ = 58
            d_1418_v133_ = p0
            d_1419_v134_: _dafny.Seq
            d_1419_v134_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gv"))
            d_1420_v135_: C0
            nw225_ = C0()
            nw225_.ctor__((d_1419_v134_) + (_dafny.SeqWithoutIsStrInference([d_1384_v101_ for d_1421_i9_ in range(default__.abs(-868))])))
            d_1420_v135_ = nw225_
            index226_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            (d_1242_v0_)[index226_] = d_1244_v1_
        d_1422_v136_: _dafny.Seq
        d_1422_v136_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ng"))
        d_1423_v137_: _dafny.Seq
        d_1423_v137_ = _dafny.SeqWithoutIsStrInference([d_1244_v1_, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])
        d_1424_v138_: _dafny.Seq
        d_1424_v138_ = _dafny.SeqWithoutIsStrInference([d_1423_v137_, d_1423_v137_])
        d_1425_v139_: _dafny.Seq
        d_1425_v139_ = _dafny.SeqWithoutIsStrInference([d_1423_v137_, (d_1424_v138_)[default__.safeIndex(p0, len(d_1424_v138_))]])
        d_1426_v140_: D6
        d_1426_v140_ = D6_DC15(p0, len(d_1422_v136_), (_dafny.SeqWithoutIsStrInference([default__.fm23(d_1244_v1_, globalState), d_1423_v137_])) == (d_1424_v138_))
        source16_ = d_1426_v140_
        if source16_.is_DC15:
            d_1427___mcc_h15_ = source16_.cf30
            d_1428___mcc_h16_ = source16_.cf31
            d_1429___mcc_h17_ = source16_.cf32
            d_1430_cf32_ = d_1429___mcc_h17_
            d_1431_cf31_ = d_1428___mcc_h16_
            d_1432_cf30_ = d_1427___mcc_h15_
            d_1385_v102_ = D0_DC1((d_1431_cf31_) - (894))
            d_1423_v137_ = d_1423_v137_
            index227_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            (d_1242_v0_)[index227_] = d_1430_cf32_
            d_1433_v141_: _dafny.Array
            nw226_ = _dafny.Array(int(0), 2)
            d_1433_v141_ = nw226_
            index228_ = default__.safeIndex(541, (d_1433_v141_).length(0))
            (d_1433_v141_)[index228_] = (0) - (d_1432_cf30_)
            index229_ = default__.safeIndex(541, (d_1433_v141_).length(0))
            (d_1433_v141_)[index229_] = 335
        elif source16_.is_DC14:
            d_1434___mcc_h18_ = source16_.cf29
            d_1435_cf29_ = d_1434___mcc_h18_
            d_1436_v142_: int
            d_1436_v142_ = 308
            d_1436_v142_ = (0) - (p0)
            d_1437_v143_: C0
            nw227_ = C0()
            nw227_.ctor__(d_1435_cf29_)
            d_1437_v143_ = nw227_
            index230_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            rhs156_ = d_1437_v143_
            rhs157_ = d_1436_v142_
            rhs158_ = (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]
            lhs92_ = d_1242_v0_
            lhs93_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            d_1437_v143_ = rhs156_
            d_1436_v142_ = rhs157_
            lhs92_[lhs93_] = rhs158_
            rhs159_ = d_1426_v140_
            rhs160_ = (d_1436_v142_) + (len(d_1383_v100_))
            rhs161_ = d_1437_v143_
            d_1426_v140_ = rhs159_
            d_1436_v142_ = rhs160_
            d_1437_v143_ = rhs161_
            d_1438_v144_: _dafny.Array
            nw228_ = _dafny.Array(None, 8)
            d_1438_v144_ = nw228_
            d_1439_v145_: _dafny.MultiSet
            d_1439_v145_ = _dafny.MultiSet([d_1244_v1_])
            d_1440_v146_: D9
            d_1440_v146_ = D9_DC21(d_1439_v145_)
            pat_let_tv38_ = d_1423_v137_
            d_1441_v147_: _dafny.Array
            nw229_ = _dafny.Array(None, 23)
            nw229_[int(0)] = len(_dafny.Map({d_1244_v1_: p0}))
            nw229_[int(1)] = (d_1385_v102_).cf2
            nw229_[int(2)] = p0
            nw229_[int(3)] = (0) - (p0)
            nw229_[int(4)] = p0
            nw229_[int(5)] = d_1436_v142_
            nw229_[int(6)] = p0
            nw229_[int(7)] = p0
            nw229_[int(8)] = (0) - (p0)
            nw229_[int(9)] = p0
            nw229_[int(10)] = d_1436_v142_
            def iife121_(_pat_let34_0):
                def iife122_(d_1442_dt__update__tmp_h0_):
                    def iife123_(_pat_let35_0):
                        def iife124_(d_1443_dt__update_hcf37_h0_):
                            return D9_DC21(d_1443_dt__update_hcf37_h0_)
                        return iife124_(_pat_let35_0)
                    return iife123_(_dafny.MultiSet(pat_let_tv38_))
                return iife122_(_pat_let34_0)
            nw229_[int(11)] = (0) - (((iife121_(d_1440_v146_)).cf37).cardinality)
            nw229_[int(12)] = d_1436_v142_
            nw229_[int(13)] = p0
            nw229_[int(14)] = d_1436_v142_
            nw229_[int(15)] = d_1436_v142_
            nw229_[int(16)] = p0
            nw229_[int(17)] = d_1436_v142_
            nw229_[int(18)] = p0
            nw229_[int(19)] = 70
            nw229_[int(20)] = p0
            nw229_[int(21)] = p0
            nw229_[int(22)] = d_1436_v142_
            d_1441_v147_ = nw229_
            d_1444_v148_: _dafny.Array
            nw230_ = _dafny.Array(None, 8)
            nw230_[int(0)] = d_1441_v147_
            nw230_[int(1)] = d_1441_v147_
            nw230_[int(2)] = d_1441_v147_
            nw230_[int(3)] = d_1441_v147_
            nw230_[int(4)] = d_1441_v147_
            nw230_[int(5)] = d_1441_v147_
            nw230_[int(6)] = d_1441_v147_
            nw230_[int(7)] = d_1441_v147_
            d_1444_v148_ = nw230_
            d_1445_v149_: _dafny.Map
            d_1445_v149_ = _dafny.Map({d_1438_v144_: d_1444_v148_})
            d_1446_v150_: _dafny.Seq
            d_1446_v150_ = _dafny.SeqWithoutIsStrInference([d_1438_v144_])
            d_1447_v151_: _dafny.Seq
            d_1447_v151_ = _dafny.SeqWithoutIsStrInference([((d_1445_v149_)[(d_1446_v150_)[default__.safeIndex(d_1436_v142_, len(d_1446_v150_))]] if ((d_1446_v150_)[default__.safeIndex(d_1436_v142_, len(d_1446_v150_))]) in (d_1445_v149_) else d_1444_v148_)])
            d_1448_v152_: _dafny.Array
            nw231_ = _dafny.Array(None, 28)
            nw231_[int(0)] = (d_1447_v151_)[default__.safeIndex(p0, len(d_1447_v151_))]
            nw231_[int(1)] = d_1444_v148_
            nw231_[int(2)] = d_1444_v148_
            nw231_[int(3)] = d_1444_v148_
            nw231_[int(4)] = (d_1444_v148_ if (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] else d_1444_v148_)
            nw231_[int(5)] = d_1444_v148_
            nw231_[int(6)] = d_1444_v148_
            nw231_[int(7)] = d_1444_v148_
            nw231_[int(8)] = d_1444_v148_
            nw231_[int(9)] = d_1444_v148_
            nw231_[int(10)] = d_1444_v148_
            nw231_[int(11)] = d_1444_v148_
            nw231_[int(12)] = d_1444_v148_
            nw231_[int(13)] = d_1444_v148_
            nw231_[int(14)] = d_1444_v148_
            nw231_[int(15)] = d_1444_v148_
            nw231_[int(16)] = d_1444_v148_
            nw231_[int(17)] = d_1444_v148_
            nw231_[int(18)] = d_1444_v148_
            nw231_[int(19)] = d_1444_v148_
            nw231_[int(20)] = d_1444_v148_
            nw231_[int(21)] = d_1444_v148_
            nw231_[int(22)] = d_1444_v148_
            nw231_[int(23)] = (d_1444_v148_ if True else d_1444_v148_)
            nw231_[int(24)] = (d_1444_v148_ if d_1244_v1_ else d_1444_v148_)
            nw231_[int(25)] = d_1444_v148_
            nw231_[int(26)] = d_1444_v148_
            nw231_[int(27)] = d_1444_v148_
            d_1448_v152_ = nw231_
            index231_ = default__.safeIndex(766, (d_1448_v152_).length(0))
            (d_1448_v152_)[index231_] = d_1444_v148_
            index232_ = default__.safeIndex(334, (d_1441_v147_).length(0))
            (d_1441_v147_)[index232_] = len((d_1383_v100_) | (d_1383_v100_))
            d_1449_v153_: _dafny.MultiSet
            d_1449_v153_ = _dafny.MultiSet([default__.fm20(True, globalState)])
            d_1450_v154_: _dafny.MultiSet
            d_1450_v154_ = _dafny.MultiSet([d_1449_v153_, d_1449_v153_])
            index233_ = default__.safeIndex(766, (d_1448_v152_).length(0))
            index234_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            index235_ = default__.safeIndex(334, (d_1441_v147_).length(0))
            rhs162_ = d_1444_v148_
            rhs163_ = (_dafny.MultiSet([d_1449_v153_, d_1449_v153_, default__.fm24((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], globalState)])).issubset(d_1450_v154_)
            rhs164_ = (502 if (self).fm5(p0, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], d_1386_v103_, d_1244_v1_, globalState) else p0)
            rhs165_ = d_1436_v142_
            rhs166_ = d_1244_v1_
            lhs94_ = d_1448_v152_
            lhs95_ = default__.safeIndex(766, (d_1448_v152_).length(0))
            lhs96_ = d_1242_v0_
            lhs97_ = default__.safeIndex(663, (d_1242_v0_).length(0))
            lhs98_ = d_1441_v147_
            lhs99_ = default__.safeIndex(334, (d_1441_v147_).length(0))
            lhs94_[lhs95_] = rhs162_
            lhs96_[lhs97_] = rhs163_
            d_1436_v142_ = rhs164_
            lhs98_[lhs99_] = rhs165_
            r0 = rhs166_
        elif True:
            d_1451___mcc_h19_ = source16_.cf33
            d_1452_cf33_ = d_1451___mcc_h19_
            d_1453_v155_: _dafny.Map
            d_1453_v155_ = _dafny.Map({p0: not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))])})
            d_1454_v156_: _dafny.Seq
            d_1454_v156_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, len(d_1453_v155_)])
            d_1455_v157_: _dafny.Map
            d_1455_v157_ = _dafny.Map({_dafny.MultiSet(d_1454_v156_): default__.fm20((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], globalState)})
            d_1456_v158_: _dafny.Set
            d_1456_v158_ = _dafny.Set({(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], (d_1423_v137_)[default__.safeIndex(994, len(d_1423_v137_))], d_1244_v1_})
            d_1457_v159_: _dafny.Array
            nw232_ = _dafny.Array(None, 29)
            nw232_[int(0)] = p0
            nw232_[int(1)] = (p0) * (p0)
            nw232_[int(2)] = (0) - (p0)
            nw232_[int(3)] = p0
            nw232_[int(4)] = (0) - ((0) - (p0))
            nw232_[int(5)] = p0
            nw232_[int(6)] = (0) - (p0)
            nw232_[int(7)] = default__.fm20(d_1244_v1_, globalState)
            nw232_[int(8)] = p0
            nw232_[int(9)] = len(_dafny.SeqWithoutIsStrInference([d_1383_v100_, d_1383_v100_]))
            nw232_[int(10)] = default__.fm20(not(False), globalState)
            nw232_[int(11)] = p0
            nw232_[int(12)] = p0
            nw232_[int(13)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_1244_v1_, False, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]])), len(d_1455_v157_))
            nw232_[int(14)] = p0
            nw232_[int(15)] = len((d_1456_v158_) | (d_1456_v158_))
            nw232_[int(16)] = default__.safeDivisionInt(len(d_1423_v137_), p0)
            nw232_[int(17)] = p0
            nw232_[int(18)] = p0
            nw232_[int(19)] = (199) * (p0)
            nw232_[int(20)] = ((0) - (p0)) + (p0)
            nw232_[int(21)] = p0
            nw232_[int(22)] = p0
            nw232_[int(23)] = (len(d_1423_v137_) if (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))] else p0)
            nw232_[int(24)] = -208
            nw232_[int(25)] = p0
            nw232_[int(26)] = default__.fm20((self).fm5(p0, d_1244_v1_, d_1386_v103_, d_1244_v1_, globalState), globalState)
            nw232_[int(27)] = -406
            nw232_[int(28)] = p0
            d_1457_v159_ = nw232_
            index236_ = default__.safeIndex(691, (d_1457_v159_).length(0))
            (d_1457_v159_)[index236_] = (0) - (p0)
            index237_ = default__.safeIndex(112, (d_1457_v159_).length(0))
            (d_1457_v159_)[index237_] = p0
            d_1458_v160_: _dafny.Seq
            d_1458_v160_ = _dafny.SeqWithoutIsStrInference([d_1422_v136_])
            index238_ = default__.safeIndex(691, (d_1457_v159_).length(0))
            index239_ = default__.safeIndex(112, (d_1457_v159_).length(0))
            rhs167_ = (len(((d_1458_v160_)[default__.safeIndex(p0, len(d_1458_v160_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqfk"))))) + (p0)
            rhs168_ = p0
            lhs100_ = d_1457_v159_
            lhs101_ = default__.safeIndex(691, (d_1457_v159_).length(0))
            lhs102_ = d_1457_v159_
            lhs103_ = default__.safeIndex(112, (d_1457_v159_).length(0))
            lhs100_[lhs101_] = rhs167_
            lhs102_[lhs103_] = rhs168_
            d_1459_v161_: _dafny.MultiSet
            d_1459_v161_ = _dafny.MultiSet([(d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], d_1244_v1_])
            d_1460_v162_: _dafny.Map
            d_1460_v162_ = _dafny.Map({d_1422_v136_: d_1459_v161_})
            d_1461_v163_: _dafny.MultiSet
            d_1461_v163_ = _dafny.MultiSet([d_1460_v162_, d_1460_v162_])
            index240_ = default__.safeIndex(691, (d_1457_v159_).length(0))
            (d_1457_v159_)[index240_] = len(_dafny.SeqWithoutIsStrInference([(d_1459_v161_).cardinality, (((d_1461_v163_)[_dafny.Map({d_1422_v136_: d_1459_v161_})] if (_dafny.Map({d_1422_v136_: d_1459_v161_})) in (d_1461_v163_) else p0)) + ((d_1457_v159_)[default__.safeIndex(691, (d_1457_v159_).length(0))]), p0]))
            index241_ = default__.safeIndex(691, (d_1457_v159_).length(0))
            (d_1457_v159_)[index241_] = 468
            d_1462_v164_: _dafny.Map
            d_1462_v164_ = _dafny.Map({d_1244_v1_: d_1422_v136_})
            d_1463_v165_: bool
            d_1464_v166_: int
            out45_: bool
            out46_: int
            out45_, out46_ = (self).m10((len(d_1462_v164_)) + (p0), (d_1457_v159_)[default__.safeIndex(691, (d_1457_v159_).length(0))], not((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]), d_1242_v0_, globalState)
            d_1463_v165_ = out45_
            d_1464_v166_ = out46_
        d_1465_v167_: _dafny.Array
        def lambda63_(d_1466_i10_):
            return _dafny.MultiSet([True])

        init33_ = lambda63_
        nw233_ = _dafny.Array(None, 20)
        for i0_33_ in range(nw233_.length(0)):
            nw233_[i0_33_] = init33_(i0_33_)
        d_1465_v167_ = nw233_
        nw234_ = _dafny.Array(_dafny.MultiSet({}), 20)
        d_1465_v167_ = nw234_
        r0 = not ((d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]) or ((self).fm5(p0, (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))], d_1386_v103_, False, globalState))
        d_1467_v168_: _dafny.Map
        d_1467_v168_ = _dafny.Map({p0: (d_1242_v0_)[default__.safeIndex(663, (d_1242_v0_).length(0))]})
        r1 = (d_1467_v168_) == (d_1467_v168_)
        return r0, r1

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_1468_v0_: bool
        d_1468_v0_ = False
        d_1469_v1_: D6
        d_1469_v1_ = D6_DC14(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1470_i0_ in range(default__.abs(491))]))
        d_1471_v2_: C0
        nw235_ = C0()
        nw235_.ctor__(((d_1469_v1_).cf29 if d_1468_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpmds"))))
        d_1471_v2_ = nw235_
        d_1472_v3_: _dafny.Map
        d_1472_v3_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0 for d_1473_i1_ in range(default__.abs(751))])): d_1468_v0_})
        d_1474_v4_: _dafny.MultiSet
        d_1474_v4_ = _dafny.MultiSet([d_1468_v0_])
        d_1475_v5_: _dafny.MultiSet
        d_1475_v5_ = _dafny.MultiSet([p0, p0, p0])
        d_1476_v6_: _dafny.Set
        d_1476_v6_ = _dafny.Set({p0, p0})
        d_1477_v7_: _dafny.Map
        d_1477_v7_ = _dafny.Map({d_1475_v5_: len(d_1476_v6_)})
        d_1478_v8_: _dafny.Map
        d_1478_v8_ = _dafny.Map({len(d_1477_v7_): 828})
        d_1479_v9_: _dafny.Map
        d_1479_v9_ = _dafny.Map({False: len(d_1478_v8_)})
        d_1480_v10_: _dafny.Array
        nw236_ = _dafny.Array(None, 20)
        nw236_[int(0)] = p0
        nw236_[int(1)] = len(d_1471_v2_.f8)
        nw236_[int(2)] = 928
        nw236_[int(3)] = p0
        nw236_[int(4)] = p0
        nw236_[int(5)] = p0
        nw236_[int(6)] = p0
        nw236_[int(7)] = p0
        nw236_[int(8)] = p0
        nw236_[int(9)] = (d_1474_v4_).cardinality
        nw236_[int(10)] = p0
        nw236_[int(11)] = p0
        nw236_[int(12)] = len(d_1479_v9_)
        nw236_[int(13)] = p0
        nw236_[int(14)] = p0
        nw236_[int(15)] = p0
        nw236_[int(16)] = p0
        nw236_[int(17)] = (0) - (p0)
        nw236_[int(18)] = p0
        nw236_[int(19)] = p0
        d_1480_v10_ = nw236_
        d_1481_v11_: _dafny.Map
        d_1481_v11_ = _dafny.Map({d_1480_v10_: p0})
        d_1482_v12_: _dafny.Map
        d_1482_v12_ = _dafny.Map({d_1481_v11_: d_1480_v10_})
        d_1483_v13_: _dafny.Array
        nw237_ = _dafny.Array(None, 15)
        nw237_[int(0)] = d_1480_v10_
        nw237_[int(1)] = d_1480_v10_
        nw237_[int(2)] = ((d_1482_v12_)[d_1481_v11_] if (d_1481_v11_) in (d_1482_v12_) else d_1480_v10_)
        nw237_[int(3)] = d_1480_v10_
        nw237_[int(4)] = d_1480_v10_
        nw237_[int(5)] = d_1480_v10_
        nw237_[int(6)] = d_1480_v10_
        nw237_[int(7)] = d_1480_v10_
        nw237_[int(8)] = d_1480_v10_
        nw237_[int(9)] = d_1480_v10_
        nw237_[int(10)] = d_1480_v10_
        nw237_[int(11)] = d_1480_v10_
        nw237_[int(12)] = d_1480_v10_
        nw237_[int(13)] = d_1480_v10_
        nw237_[int(14)] = d_1480_v10_
        d_1483_v13_ = nw237_
        d_1484_v14_: D2
        d_1484_v14_ = D2_DC4(p0, d_1468_v0_, d_1483_v13_)
        d_1485_v15_: _dafny.Seq
        d_1485_v15_ = _dafny.SeqWithoutIsStrInference([((d_1479_v9_)[d_1468_v0_] if (d_1468_v0_) in (d_1479_v9_) else p0), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uotirlywr")))])
        d_1486_v16_: _dafny.Array
        d_1486_v16_ = d_1480_v10_
        d_1487_v17_: _dafny.Map
        d_1487_v17_ = _dafny.Map({p0: (d_1486_v16_)})
        d_1488_v18_: _dafny.Seq
        d_1488_v18_ = _dafny.SeqWithoutIsStrInference([d_1471_v2_])
        d_1489_v19_: _dafny.Array
        nw238_ = _dafny.Array(None, 19)
        nw238_[int(0)] = len(d_1471_v2_.f8)
        nw238_[int(1)] = p0
        nw238_[int(2)] = len(d_1472_v3_)
        nw238_[int(3)] = (p0) + (p0)
        nw238_[int(4)] = (p0) + (p0)
        nw238_[int(5)] = (0) - (p0)
        nw238_[int(6)] = p0
        nw238_[int(7)] = (p0) * (p0)
        nw238_[int(8)] = len(_dafny.SeqWithoutIsStrInference([default__.fm20(d_1468_v0_, globalState)]))
        nw238_[int(9)] = (0) - ((85) + (len(d_1471_v2_.f8)))
        nw238_[int(10)] = p0
        nw238_[int(11)] = p0
        nw238_[int(12)] = (d_1484_v14_).cf5
        nw238_[int(13)] = p0
        nw238_[int(14)] = (d_1485_v15_)[default__.safeIndex(p0, len(d_1485_v15_))]
        nw238_[int(15)] = p0
        nw238_[int(16)] = (0) - (len(d_1487_v17_))
        nw238_[int(17)] = default__.safeDivisionInt(len(d_1488_v18_), p0)
        nw238_[int(18)] = p0
        d_1489_v19_ = nw238_
        d_1490_v20_: _dafny.MultiSet
        d_1490_v20_ = _dafny.MultiSet([(d_1469_v1_).cf29, d_1471_v2_.f8])
        nw239_ = _dafny.Array(None, 12)
        nw239_[int(0)] = (d_1484_v14_).cf5
        nw239_[int(1)] = (d_1485_v15_)[default__.safeIndex(p0, len(d_1485_v15_))]
        nw239_[int(2)] = p0
        nw239_[int(3)] = default__.safeModuloInt((0) - (p0), p0)
        nw239_[int(4)] = p0
        nw239_[int(5)] = (p0) + ((d_1490_v20_).cardinality)
        nw239_[int(6)] = p0
        nw239_[int(7)] = (p0) + (75)
        nw239_[int(8)] = (0) - ((p0) + (p0))
        nw239_[int(9)] = (p0) + (p0)
        nw239_[int(10)] = p0
        nw239_[int(11)] = (p0) - (p0)
        d_1489_v19_ = nw239_
        index242_ = default__.safeIndex(655, (d_1480_v10_).length(0))
        (d_1480_v10_)[index242_] = default__.safeModuloInt(len(d_1471_v2_.f8), p0)
        index243_ = default__.safeIndex(655, (d_1480_v10_).length(0))
        (d_1480_v10_)[index243_] = p0
        d_1491_i2_: int
        d_1491_i2_ = 0
        with _dafny.label("4"):
            while not(False):
                with _dafny.c_label("4"):
                    if (d_1491_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_1491_i2_ = (d_1491_i2_) + (1)
                    d_1492_v21_: D4
                    d_1492_v21_ = D4_DC10(p0, p0, (d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))])
                    d_1493_v22_: str
                    d_1493_v22_ = _dafny.CodePoint('t')
                    d_1494_v23_: _dafny.Map
                    d_1494_v23_ = _dafny.Map({default__.fm25(d_1492_v21_, d_1471_v2_.f8, globalState): d_1493_v22_})
                    d_1495_v24_: _dafny.Map
                    d_1495_v24_ = _dafny.Map({d_1493_v22_: d_1494_v23_})
                    d_1496_v25_: _dafny.Seq
                    d_1496_v25_ = _dafny.SeqWithoutIsStrInference([d_1474_v4_])
                    d_1497_v27_: _dafny.Set
                    d_1497_v27_ = _dafny.Set({d_1471_v2_.f8})
                    d_1498_v28_: _dafny.Map
                    d_1498_v28_ = _dafny.Map({(0) - (p0): D2_DC3(d_1497_v27_)})
                    d_1499_v29_: _dafny.Map
                    d_1499_v29_ = _dafny.Map({d_1471_v2_.f8: d_1498_v28_})
                    d_1500_v30_: _dafny.Set
                    d_1500_v30_ = _dafny.Set({d_1468_v0_})
                    d_1501_v31_: _dafny.Map
                    d_1501_v31_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_1468_v0_])})
                    d_1502_v32_: _dafny.Map
                    d_1502_v32_ = _dafny.Map({d_1471_v2_: p0})
                    d_1503_v33_: _dafny.Seq
                    d_1503_v33_ = _dafny.SeqWithoutIsStrInference([d_1468_v0_, d_1468_v0_])
                    d_1504_v34_: D0
                    d_1504_v34_ = D0_DC1((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))])
                    d_1505_v37_: _dafny.Map
                    d_1505_v37_ = _dafny.Map({d_1471_v2_.f8: d_1468_v0_})
                    d_1506_v40_: _dafny.Map
                    d_1506_v40_ = _dafny.Map({default__.fm25(d_1492_v21_, d_1471_v2_.f8, globalState): (d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))]})
                    d_1507_v41_: D12
                    d_1507_v41_ = D12_DC27(d_1500_v30_)
                    d_1508_v42_: D11
                    d_1508_v42_ = D11_DC24(default__.fm26((d_1507_v41_).cf45, d_1503_v33_, d_1504_v34_, d_1493_v22_, globalState))
                    d_1509_v43_: _dafny.Array
                    nw240_ = _dafny.Array(None, 29)
                    nw240_[int(0)] = d_1494_v23_
                    nw240_[int(1)] = _dafny.Map({d_1471_v2_.f8: d_1493_v22_})
                    nw240_[int(2)] = d_1494_v23_
                    nw240_[int(3)] = d_1494_v23_
                    nw240_[int(4)] = d_1494_v23_
                    nw240_[int(5)] = (d_1494_v23_).set(d_1471_v2_.f8, d_1493_v22_)
                    nw240_[int(6)] = ((d_1495_v24_)[d_1493_v22_] if (d_1493_v22_) in (d_1495_v24_) else d_1494_v23_)
                    nw240_[int(7)] = _dafny.Map({(d_1471_v2_).fm15(_dafny.SeqWithoutIsStrInference([len(d_1496_v25_)]), d_1468_v0_, globalState): _dafny.CodePoint('s')})
                    nw240_[int(8)] = d_1494_v23_
                    def iife125_():
                        coll53_ = _dafny.Map()
                        compr_53_: _dafny.Seq
                        for compr_53_ in (d_1499_v29_).keys.Elements:
                            d_1510_v26_: _dafny.Seq = compr_53_
                            if (d_1510_v26_) in (d_1499_v29_):
                                coll53_[d_1510_v26_] = _dafny.CodePoint('t')
                        return _dafny.Map(coll53_)
                    nw240_[int(9)] = iife125_()
                    
                    nw240_[int(10)] = d_1494_v23_
                    nw240_[int(11)] = d_1494_v23_
                    nw240_[int(12)] = d_1494_v23_
                    nw240_[int(13)] = d_1494_v23_
                    nw240_[int(14)] = default__.fm26(d_1500_v30_, ((d_1501_v31_)[((d_1502_v32_)[d_1471_v2_] if (d_1471_v2_) in (d_1502_v32_) else (d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))])] if (((d_1502_v32_)[d_1471_v2_] if (d_1471_v2_) in (d_1502_v32_) else (d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))])) in (d_1501_v31_) else d_1503_v33_), d_1504_v34_, d_1493_v22_, globalState)
                    def iife126_():
                        coll54_ = _dafny.Map()
                        compr_54_: _dafny.Seq
                        for compr_54_ in (d_1490_v20_).Elements:
                            d_1511_v35_: _dafny.Seq = compr_54_
                            if (d_1511_v35_) in (d_1490_v20_):
                                coll54_[d_1511_v35_] = d_1493_v22_
                        return _dafny.Map(coll54_)
                    nw240_[int(15)] = (_dafny.Map({d_1471_v2_.f8: (d_1471_v2_).fm14(globalState)})) | ((iife126_()
                    ).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")), d_1493_v22_))
                    nw240_[int(16)] = d_1494_v23_
                    nw240_[int(17)] = (d_1494_v23_) | (_dafny.Map({d_1471_v2_.f8: (d_1471_v2_.f8)[default__.safeIndex((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))], len(d_1471_v2_.f8))]}))
                    def iife127_():
                        coll55_ = _dafny.Map()
                        compr_55_: _dafny.Seq
                        for compr_55_ in (d_1505_v37_).keys.Elements:
                            d_1512_v36_: _dafny.Seq = compr_55_
                            if (d_1512_v36_) in (d_1505_v37_):
                                coll55_[d_1512_v36_] = d_1493_v22_
                        return _dafny.Map(coll55_)
                    nw240_[int(18)] = iife127_()
                    
                    nw240_[int(19)] = d_1494_v23_
                    nw240_[int(20)] = d_1494_v23_
                    nw240_[int(21)] = d_1494_v23_
                    def iife128_():
                        def iife130_():
                            coll58_ = _dafny.Map()
                            compr_58_: _dafny.Seq
                            for compr_58_ in (d_1506_v40_).keys.Elements:
                                d_1513_v39_: _dafny.Seq = compr_58_
                                if (d_1513_v39_) in (d_1506_v40_):
                                    coll58_[d_1513_v39_] = d_1468_v0_
                            return _dafny.Map(coll58_)
                        coll56_ = _dafny.Map()
                        def iife129_():
                            coll57_ = _dafny.Map()
                            compr_57_: _dafny.Seq
                            for compr_57_ in (d_1506_v40_).keys.Elements:
                                d_1513_v39_: _dafny.Seq = compr_57_
                                if (d_1513_v39_) in (d_1506_v40_):
                                    coll57_[d_1513_v39_] = d_1468_v0_
                            return _dafny.Map(coll57_)
                        compr_56_: _dafny.Seq
                        for compr_56_ in (iife129_()
                        ).keys.Elements:
                            d_1514_v38_: _dafny.Seq = compr_56_
                            if (d_1514_v38_) in (iife130_()
                            ):
                                coll56_[d_1514_v38_] = d_1493_v22_
                        return _dafny.Map(coll56_)
                    nw240_[int(22)] = iife128_()
                    
                    nw240_[int(23)] = d_1494_v23_
                    nw240_[int(24)] = d_1494_v23_
                    nw240_[int(25)] = (d_1508_v42_).cf40
                    nw240_[int(26)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1493_v22_ for d_1515_i3_ in range(default__.abs(890))]): d_1493_v22_})
                    nw240_[int(27)] = (d_1494_v23_) | (d_1494_v23_)
                    nw240_[int(28)] = d_1494_v23_
                    d_1509_v43_ = nw240_
                    index244_ = default__.safeIndex(622, (d_1509_v43_).length(0))
                    (d_1509_v43_)[index244_] = d_1494_v23_
                    index245_ = default__.safeIndex(622, (d_1509_v43_).length(0))
                    (d_1509_v43_)[index245_] = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1493_v22_ for d_1516_i4_ in range(default__.abs(136))]): d_1493_v22_})
                    d_1517_v44_: _dafny.Array
                    nw241_ = _dafny.Array(_dafny.Map({}), 21)
                    d_1517_v44_ = nw241_
                    d_1518_v45_: _dafny.Map
                    d_1518_v45_ = _dafny.Map({d_1493_v22_: len(d_1476_v6_)})
                    d_1519_v46_: _dafny.MultiSet
                    d_1519_v46_ = _dafny.MultiSet([d_1518_v45_, d_1518_v45_])
                    d_1520_v47_: T0
                    nw242_ = C2()
                    nw242_.ctor__()
                    d_1520_v47_ = nw242_
                    d_1521_v48_: _dafny.Map
                    d_1521_v48_ = _dafny.Map({True: d_1520_v47_})
                    d_1522_v49_: _dafny.Map
                    d_1522_v49_ = _dafny.Map({d_1519_v46_: d_1521_v48_})
                    index246_ = default__.safeIndex(735, (d_1517_v44_).length(0))
                    (d_1517_v44_)[index246_] = (d_1522_v49_ if d_1468_v0_ else d_1522_v49_)
                    d_1523_v50_: _dafny.Seq
                    d_1523_v50_ = _dafny.SeqWithoutIsStrInference([d_1471_v2_.f8, d_1471_v2_.f8])
                    d_1524_v51_: _dafny.Map
                    d_1524_v51_ = _dafny.Map({d_1468_v0_: d_1521_v48_})
                    index247_ = default__.safeIndex(735, (d_1517_v44_).length(0))
                    index248_ = default__.safeIndex(655, (d_1480_v10_).length(0))
                    rhs169_ = d_1468_v0_
                    rhs170_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xone"))) < ((d_1523_v50_)[default__.safeIndex(p0, len(d_1523_v50_))])
                    rhs171_ = ((d_1522_v49_ if d_1468_v0_ else _dafny.Map({_dafny.MultiSet([d_1518_v45_]): d_1521_v48_}))) | (_dafny.Map({d_1519_v46_: ((d_1524_v51_)[d_1468_v0_] if (d_1468_v0_) in (d_1524_v51_) else d_1521_v48_)}))
                    rhs172_ = (0) - (((p0) + (default__.fm20(d_1468_v0_, globalState)) if (d_1472_v3_) != (d_1472_v3_) else ((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))]) + (p0)))
                    lhs104_ = d_1517_v44_
                    lhs105_ = default__.safeIndex(735, (d_1517_v44_).length(0))
                    lhs106_ = d_1480_v10_
                    lhs107_ = default__.safeIndex(655, (d_1480_v10_).length(0))
                    d_1468_v0_ = rhs169_
                    d_1468_v0_ = rhs170_
                    lhs104_[lhs105_] = rhs171_
                    lhs106_[lhs107_] = rhs172_
                    d_1525_v52_: _dafny.Map
                    d_1525_v52_ = _dafny.Map({_dafny.CodePoint('c'): d_1504_v34_})
                    d_1526_v53_: D2
                    d_1526_v53_ = D2_DC5((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))], d_1485_v15_, False)
                    if (self).fm5((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))], not(((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))]) <= (p0)), d_1525_v52_, (d_1526_v53_).cf10, globalState):
                        d_1527_v54_: _dafny.Array
                        def lambda64_(d_1528_i5_):
                            return True

                        init34_ = lambda64_
                        nw243_ = _dafny.Array(None, 12)
                        for i0_34_ in range(nw243_.length(0)):
                            nw243_[i0_34_] = init34_(i0_34_)
                        d_1527_v54_ = nw243_
                        d_1527_v54_ = d_1527_v54_
                        index249_ = default__.safeIndex(975, (d_1483_v13_).length(0))
                        (d_1483_v13_)[index249_] = d_1489_v19_
                        index250_ = default__.safeIndex(975, (d_1483_v13_).length(0))
                        (d_1483_v13_)[index250_] = d_1489_v19_
                        d_1529_v55_: bool
                        d_1530_v56_: bool
                        out47_: bool
                        out48_: bool
                        out47_, out48_ = (d_1520_v47_).m1((0) - (default__.safeModuloInt(p0, 5)), globalState)
                        d_1529_v55_ = out47_
                        d_1530_v56_ = out48_
                        d_1531_v57_: _dafny.Map
                        d_1531_v57_ = _dafny.Map({d_1527_v54_: True})
                        d_1530_v56_ = (d_1531_v57_) == (d_1531_v57_)
                        index251_ = default__.safeIndex(655, (d_1480_v10_).length(0))
                        (d_1480_v10_)[index251_] = (d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))]
                    elif True:
                        d_1532_v58_: D7
                        d_1532_v58_ = D7_DC17(d_1503_v33_)
                        index252_ = default__.safeIndex(655, (d_1480_v10_).length(0))
                        (d_1480_v10_)[index252_] = len((d_1532_v58_).cf34)
                        (d_1471_v2_).f8 = _dafny.SeqWithoutIsStrInference([d_1493_v22_ for d_1533_i6_ in range(default__.abs(324))])
                        d_1534_v59_: D16
                        d_1534_v59_ = D16_DC44((d_1480_v10_)[default__.safeIndex(655, (d_1480_v10_).length(0))])
                        pat_let_tv39_ = d_1485_v15_
                        pat_let_tv40_ = d_1479_v9_
                        pat_let_tv41_ = d_1485_v15_
                        def iife131_(_pat_let36_0):
                            def iife132_(d_1535_dt__update__tmp_h0_):
                                def iife133_(_pat_let37_0):
                                    def iife134_(d_1536_dt__update_hcf82_h0_):
                                        return D16_DC44(d_1536_dt__update_hcf82_h0_)
                                    return iife134_(_pat_let37_0)
                                return iife133_((pat_let_tv39_)[default__.safeIndex(len(pat_let_tv40_), len(pat_let_tv41_))])
                            return iife132_(_pat_let36_0)
                        d_1534_v59_ = iife131_(d_1534_v59_)
                        d_1505_v37_ = (d_1505_v37_).set(d_1471_v2_.f8, d_1468_v0_)
                        d_1475_v5_ = (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([-76 for d_1537_i7_ in range(default__.abs(311))])) + (d_1485_v15_))).intersection((d_1475_v5_) | (d_1475_v5_))
                    d_1538_v60_: D9
                    d_1538_v60_ = D9_DC21((D9_DC21(d_1474_v4_)).cf37)
                    d_1539_v61_: _dafny.Seq
                    d_1539_v61_ = _dafny.SeqWithoutIsStrInference([d_1538_v60_, d_1538_v60_, d_1538_v60_])
                    d_1539_v61_ = (d_1539_v61_) + (d_1539_v61_)
                    pass
            pass
        d_1540_v62_: C2
        nw244_ = C2()
        nw244_.ctor__()
        d_1540_v62_ = nw244_
        index253_ = default__.safeIndex(655, (d_1480_v10_).length(0))
        (d_1480_v10_)[index253_] = p0
        d_1541_v63_: _dafny.Map
        d_1541_v63_ = _dafny.Map({not(d_1468_v0_): d_1475_v5_})
        r0 = d_1541_v63_
        return r0

    def m9(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        d_1542_v0_: str
        d_1542_v0_ = _dafny.CodePoint('x')
        d_1543_v1_: _dafny.Seq
        d_1543_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvypfkdr"))
        d_1544_v2_: bool
        d_1544_v2_ = False
        d_1545_v3_: int
        d_1545_v3_ = -189
        d_1546_v4_: _dafny.Seq
        d_1546_v4_ = _dafny.SeqWithoutIsStrInference([d_1544_v2_])
        d_1547_v5_: _dafny.MultiSet
        d_1547_v5_ = _dafny.MultiSet([True, d_1544_v2_])
        d_1548_v6_: D0
        d_1548_v6_ = D0_DC1(d_1545_v3_)
        d_1549_v7_: _dafny.Map
        d_1549_v7_ = _dafny.Map({d_1542_v0_: d_1548_v6_})
        d_1550_v8_: _dafny.Set
        d_1550_v8_ = _dafny.Set({d_1544_v2_})
        d_1551_v9_: _dafny.Array
        nw245_ = _dafny.Array(None, 28)
        nw245_[int(0)] = (d_1542_v0_) in (d_1543_v1_)
        nw245_[int(1)] = d_1544_v2_
        nw245_[int(2)] = (d_1542_v0_) not in (_dafny.SeqWithoutIsStrInference([d_1542_v0_ for d_1552_i0_ in range(default__.abs(343))]))
        nw245_[int(3)] = (d_1543_v1_) == (default__.fm31(d_1545_v3_, d_1544_v2_, len(d_1546_v4_), globalState))
        nw245_[int(4)] = d_1544_v2_
        nw245_[int(5)] = (d_1547_v5_).isdisjoint(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).fm5(len(d_1543_v1_), d_1544_v2_, d_1549_v7_, d_1544_v2_, globalState)])))
        nw245_[int(6)] = d_1544_v2_
        nw245_[int(7)] = d_1544_v2_
        nw245_[int(8)] = d_1544_v2_
        nw245_[int(9)] = (d_1545_v3_) >= (d_1545_v3_)
        nw245_[int(10)] = (d_1544_v2_) and (d_1544_v2_)
        nw245_[int(11)] = d_1544_v2_
        nw245_[int(12)] = (d_1544_v2_) and (True)
        nw245_[int(13)] = d_1544_v2_
        nw245_[int(14)] = d_1544_v2_
        nw245_[int(15)] = d_1544_v2_
        nw245_[int(16)] = d_1544_v2_
        nw245_[int(17)] = (d_1550_v8_).isdisjoint(_dafny.Set({d_1544_v2_}))
        nw245_[int(18)] = d_1544_v2_
        nw245_[int(19)] = (not(d_1544_v2_) if d_1544_v2_ else True)
        nw245_[int(20)] = d_1544_v2_
        nw245_[int(21)] = not((d_1544_v2_ if d_1544_v2_ else True))
        nw245_[int(22)] = False
        nw245_[int(23)] = d_1544_v2_
        nw245_[int(24)] = (d_1545_v3_) < (d_1545_v3_)
        nw245_[int(25)] = (d_1547_v5_).issubset(d_1547_v5_)
        nw245_[int(26)] = d_1544_v2_
        nw245_[int(27)] = d_1544_v2_
        d_1551_v9_ = nw245_
        index254_ = default__.safeIndex(117, (d_1551_v9_).length(0))
        (d_1551_v9_)[index254_] = d_1544_v2_
        index255_ = default__.safeIndex(117, (d_1551_v9_).length(0))
        (d_1551_v9_)[index255_] = d_1544_v2_
        d_1553_v10_: _dafny.Seq
        d_1553_v10_ = _dafny.SeqWithoutIsStrInference([d_1545_v3_])
        d_1554_v11_: _dafny.Seq
        d_1554_v11_ = _dafny.SeqWithoutIsStrInference([d_1553_v10_, _dafny.SeqWithoutIsStrInference([len(d_1553_v10_) for d_1555_i1_ in range(default__.abs(185))])])
        d_1556_v12_: _dafny.Set
        d_1556_v12_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_1545_v3_]), (d_1554_v11_)[default__.safeIndex(d_1545_v3_, len(d_1554_v11_))]})
        d_1557_v13_: D15
        d_1557_v13_ = D15_DC37(d_1556_v12_)
        d_1557_v13_ = d_1557_v13_
        d_1558_v14_: C2
        nw246_ = C2()
        nw246_.ctor__()
        d_1558_v14_ = nw246_
        index256_ = default__.safeIndex(117, (d_1551_v9_).length(0))
        rhs173_ = d_1558_v14_
        rhs174_ = 689
        rhs175_ = d_1545_v3_
        rhs176_ = not ((d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))]) or ((d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))])
        lhs108_ = d_1551_v9_
        lhs109_ = default__.safeIndex(117, (d_1551_v9_).length(0))
        d_1558_v14_ = rhs173_
        d_1545_v3_ = rhs174_
        d_1545_v3_ = rhs175_
        lhs108_[lhs109_] = rhs176_
        d_1559_v15_: _dafny.Map
        d_1559_v15_ = _dafny.Map({d_1545_v3_: (d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))]})
        d_1559_v15_ = (d_1559_v15_).set(d_1545_v3_, (_dafny.Set({d_1544_v2_, d_1544_v2_, (d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))]})) == (d_1550_v8_))
        hi10_ = d_1545_v3_
        for d_1560_i2_ in range((d_1545_v3_) - (d_1545_v3_), hi10_):
            d_1543_v1_ = d_1543_v1_
            rhs177_ = True
            rhs178_ = (default__.safeModuloInt(d_1545_v3_, d_1545_v3_)) in (d_1553_v10_)
            r2 = rhs177_
            d_1544_v2_ = rhs178_
            d_1561_v16_: _dafny.Array
            nw247_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_1561_v16_ = nw247_
            d_1562_v17_: _dafny.Array
            nw248_ = _dafny.Array(None, 20)
            nw248_[int(0)] = d_1543_v1_
            nw248_[int(1)] = d_1543_v1_
            nw248_[int(2)] = d_1543_v1_
            nw248_[int(3)] = d_1543_v1_
            nw248_[int(4)] = (d_1543_v1_).set(default__.safeIndex(d_1560_i2_, len(d_1543_v1_)), d_1542_v0_)
            nw248_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oe"))
            nw248_[int(6)] = d_1543_v1_
            nw248_[int(7)] = d_1543_v1_
            nw248_[int(8)] = d_1543_v1_
            nw248_[int(9)] = d_1543_v1_
            nw248_[int(10)] = d_1543_v1_
            nw248_[int(11)] = d_1543_v1_
            nw248_[int(12)] = d_1543_v1_
            nw248_[int(13)] = d_1543_v1_
            nw248_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "diut"))
            nw248_[int(15)] = d_1543_v1_
            nw248_[int(16)] = d_1543_v1_
            nw248_[int(17)] = d_1543_v1_
            nw248_[int(18)] = d_1543_v1_
            nw248_[int(19)] = d_1543_v1_
            d_1562_v17_ = nw248_
            index257_ = default__.safeIndex(356, (d_1561_v16_).length(0))
            (d_1561_v16_)[index257_] = d_1562_v17_
            d_1563_v18_: _dafny.Seq
            d_1563_v18_ = _dafny.SeqWithoutIsStrInference([d_1562_v17_])
            index258_ = default__.safeIndex(356, (d_1561_v16_).length(0))
            (d_1561_v16_)[index258_] = (d_1563_v18_)[default__.safeIndex(d_1545_v3_, len(d_1563_v18_))]
            d_1564_v19_: _dafny.Set
            d_1564_v19_ = _dafny.Set({d_1551_v9_})
            index259_ = default__.safeIndex(117, (d_1551_v9_).length(0))
            (d_1551_v9_)[index259_] = (d_1564_v19_).ispropersubset(d_1564_v19_)
        d_1565_v20_: _dafny.Set
        d_1565_v20_ = _dafny.Set({d_1551_v9_, d_1551_v9_, d_1551_v9_})
        d_1566_v21_: _dafny.Set
        d_1566_v21_ = d_1565_v20_
        d_1567_v22_: _dafny.Set
        d_1567_v22_ = ((d_1566_v21_)) - (d_1565_v20_)
        source17_ = d_1566_v21_
        d_1568___mcc_h0_ = source17_
        d_1569_cf36_ = d_1568___mcc_h0_
        d_1545_v3_ = (0) - (default__.safeModuloInt(default__.safeDivisionInt(d_1545_v3_, d_1545_v3_), default__.safeDivisionInt(d_1545_v3_, 317)))
        d_1570_v23_: _dafny.Array
        nw249_ = _dafny.Array(int(0), 3)
        d_1570_v23_ = nw249_
        d_1570_v23_ = d_1570_v23_
        d_1571_v25_: _dafny.Map
        d_1571_v25_ = _dafny.Map({d_1542_v0_: d_1545_v3_})
        d_1572_v27_: _dafny.Set
        d_1572_v27_ = _dafny.Set({d_1542_v0_})
        d_1573_v28_: _dafny.Map
        d_1573_v28_ = _dafny.Map({d_1542_v0_: _dafny.SeqWithoutIsStrInference([(d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))], (d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))]])})
        d_1574_v30_: D6
        d_1574_v30_ = D6_DC14(d_1543_v1_)
        d_1575_v31_: _dafny.Map
        d_1575_v31_ = _dafny.Map({d_1542_v0_: d_1574_v30_})
        d_1576_v32_: _dafny.Map
        d_1576_v32_ = _dafny.Map({d_1545_v3_: d_1573_v28_})
        d_1577_v33_: _dafny.Map
        d_1577_v33_ = _dafny.Map({d_1545_v3_: d_1545_v3_})
        d_1578_v35_: _dafny.Array
        nw250_ = _dafny.Array(None, 23)
        def iife135_():
            coll59_ = _dafny.Map()
            compr_59_: str
            for compr_59_ in (d_1571_v25_).keys.Elements:
                d_1579_v24_: str = compr_59_
                if (d_1579_v24_) in (d_1571_v25_):
                    coll59_[d_1579_v24_] = d_1546_v4_
            return _dafny.Map(coll59_)
        nw250_[int(0)] = iife135_()
        
        def iife136_():
            coll60_ = _dafny.Map()
            compr_60_: str
            for compr_60_ in (d_1572_v27_).Elements:
                d_1580_v26_: str = compr_60_
                if (d_1580_v26_) in (d_1572_v27_):
                    coll60_[d_1580_v26_] = d_1546_v4_
            return _dafny.Map(coll60_)
        nw250_[int(1)] = iife136_()
        
        nw250_[int(2)] = (_dafny.Map({d_1542_v0_: d_1546_v4_}) if (d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))] else d_1573_v28_)
        nw250_[int(3)] = d_1573_v28_
        def iife137_():
            coll61_ = _dafny.Map()
            compr_61_: str
            for compr_61_ in (d_1575_v31_).keys.Elements:
                d_1581_v29_: str = compr_61_
                if (d_1581_v29_) in (d_1575_v31_):
                    coll61_[d_1581_v29_] = d_1546_v4_
            return _dafny.Map(coll61_)
        nw250_[int(4)] = iife137_()
        
        nw250_[int(5)] = ((d_1576_v32_)[((d_1577_v33_)[d_1545_v3_] if (d_1545_v3_) in (d_1577_v33_) else d_1545_v3_)] if (((d_1577_v33_)[d_1545_v3_] if (d_1545_v3_) in (d_1577_v33_) else d_1545_v3_)) in (d_1576_v32_) else d_1573_v28_)
        nw250_[int(6)] = d_1573_v28_
        nw250_[int(7)] = d_1573_v28_
        nw250_[int(8)] = d_1573_v28_
        nw250_[int(9)] = _dafny.Map({d_1542_v0_: d_1546_v4_})
        nw250_[int(10)] = d_1573_v28_
        nw250_[int(11)] = d_1573_v28_
        nw250_[int(12)] = d_1573_v28_
        nw250_[int(13)] = d_1573_v28_
        nw250_[int(14)] = d_1573_v28_
        nw250_[int(15)] = _dafny.Map({default__.fm51(False, d_1545_v3_, d_1545_v3_, globalState): _dafny.SeqWithoutIsStrInference([not((d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))])])})
        nw250_[int(16)] = (d_1573_v28_) | (d_1573_v28_)
        nw250_[int(17)] = (d_1573_v28_) | (d_1573_v28_)
        nw250_[int(18)] = _dafny.Map({d_1542_v0_: d_1546_v4_})
        nw250_[int(19)] = d_1573_v28_
        nw250_[int(20)] = d_1573_v28_
        nw250_[int(21)] = _dafny.Map({d_1542_v0_: d_1546_v4_})
        def iife138_():
            coll62_ = _dafny.Map()
            compr_62_: str
            for compr_62_ in (_dafny.SeqWithoutIsStrInference([d_1542_v0_])).Elements:
                d_1582_v34_: str = compr_62_
                if (d_1582_v34_) in (_dafny.SeqWithoutIsStrInference([d_1542_v0_])):
                    coll62_[d_1582_v34_] = d_1546_v4_
            return _dafny.Map(coll62_)
        nw250_[int(22)] = (iife138_()
        ) | ((d_1573_v28_).set(default__.fm51((d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))], d_1545_v3_, len(_dafny.SeqWithoutIsStrInference([d_1545_v3_])), globalState), d_1546_v4_))
        d_1578_v35_ = nw250_
        index260_ = default__.safeIndex(60, (d_1578_v35_).length(0))
        def iife139_():
            coll63_ = _dafny.Map()
            compr_63_: str
            for compr_63_ in (d_1573_v28_).keys.Elements:
                d_1583_v36_: str = compr_63_
                if (d_1583_v36_) in (d_1573_v28_):
                    coll63_[d_1583_v36_] = _dafny.SeqWithoutIsStrInference([d_1544_v2_, (d_1551_v9_)[default__.safeIndex(117, (d_1551_v9_).length(0))]])
            return _dafny.Map(coll63_)
        (d_1578_v35_)[index260_] = iife139_()
        
        index261_ = default__.safeIndex(60, (d_1578_v35_).length(0))
        def iife140_():
            coll64_ = _dafny.Map()
            compr_64_: str
            for compr_64_ in (d_1543_v1_).Elements:
                d_1584_v37_: str = compr_64_
                if (d_1584_v37_) in (d_1543_v1_):
                    coll64_[d_1584_v37_] = d_1546_v4_
            return _dafny.Map(coll64_)
        (d_1578_v35_)[index261_] = ((d_1573_v28_) | (iife140_()
        )) | (d_1573_v28_)
        d_1585_v38_: _dafny.Array
        nw251_ = _dafny.Array(_dafny.Map({}), 29)
        d_1585_v38_ = nw251_
        d_1586_v39_: C1
        nw252_ = C1()
        nw252_.ctor__(len((d_1543_v1_) + (d_1543_v1_)), d_1585_v38_, d_1553_v10_, d_1548_v6_)
        d_1586_v39_ = nw252_
        d_1587_v40_: _dafny.Array
        def lambda65_(d_1588_v6_):
            def lambda66_(d_1589_i3_):
                return d_1588_v6_

            return lambda66_

        init35_ = lambda65_(d_1548_v6_)
        nw253_ = _dafny.Array(None, 3)
        for i0_35_ in range(nw253_.length(0)):
            nw253_[i0_35_] = init35_(i0_35_)
        d_1587_v40_ = nw253_
        r0 = d_1587_v40_
        d_1590_v41_: _dafny.Array
        def lambda67_(d_1591_v3_):
            def lambda68_(d_1592_i4_):
                return default__.safeDivisionInt(d_1592_i4_, (0) - (d_1591_v3_))

            return lambda68_

        init36_ = lambda67_(d_1545_v3_)
        nw254_ = _dafny.Array(None, 3)
        for i0_36_ in range(nw254_.length(0)):
            nw254_[i0_36_] = init36_(i0_36_)
        d_1590_v41_ = nw254_
        d_1593_v42_: D13
        d_1593_v42_ = D13_DC32(d_1543_v1_, d_1544_v2_, (_dafny.MultiSet([d_1544_v2_, False])).cardinality, d_1590_v41_)
        d_1594_v43_: _dafny.Seq
        d_1594_v43_ = _dafny.SeqWithoutIsStrInference([d_1543_v1_])
        d_1595_v44_: _dafny.Map
        d_1595_v44_ = _dafny.Map({len((d_1594_v43_)[default__.safeIndex(d_1545_v3_, len(d_1594_v43_))]): d_1590_v41_})
        d_1596_v45_: _dafny.Array
        nw255_ = _dafny.Array(None, 27)
        nw255_[int(0)] = d_1590_v41_
        nw255_[int(1)] = (d_1590_v41_ if d_1544_v2_ else d_1590_v41_)
        nw255_[int(2)] = d_1590_v41_
        nw255_[int(3)] = d_1590_v41_
        nw255_[int(4)] = d_1590_v41_
        nw255_[int(5)] = d_1590_v41_
        nw255_[int(6)] = (d_1593_v42_).cf60
        nw255_[int(7)] = d_1590_v41_
        nw255_[int(8)] = d_1590_v41_
        nw255_[int(9)] = d_1590_v41_
        nw255_[int(10)] = d_1590_v41_
        nw255_[int(11)] = d_1590_v41_
        nw255_[int(12)] = d_1590_v41_
        nw255_[int(13)] = d_1590_v41_
        nw255_[int(14)] = d_1590_v41_
        nw255_[int(15)] = d_1590_v41_
        nw255_[int(16)] = d_1590_v41_
        nw255_[int(17)] = d_1590_v41_
        nw255_[int(18)] = d_1590_v41_
        nw255_[int(19)] = d_1590_v41_
        nw255_[int(20)] = d_1590_v41_
        nw255_[int(21)] = ((d_1595_v44_)[d_1545_v3_] if (d_1545_v3_) in (d_1595_v44_) else d_1590_v41_)
        nw255_[int(22)] = d_1590_v41_
        nw255_[int(23)] = d_1590_v41_
        nw255_[int(24)] = d_1590_v41_
        nw255_[int(25)] = d_1590_v41_
        nw255_[int(26)] = d_1590_v41_
        d_1596_v45_ = nw255_
        r1 = d_1596_v45_
        r2 = (len(d_1543_v1_)) > (396)
        return r0, r1, r2

    def m10(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        d_1597_v0_: str
        d_1597_v0_ = _dafny.CodePoint('u')
        d_1598_v1_: _dafny.Map
        d_1598_v1_ = _dafny.Map({d_1597_v0_: p1})
        d_1599_v2_: _dafny.Map
        d_1599_v2_ = d_1598_v1_
        d_1599_v2_ = d_1598_v1_
        hi11_ = p0
        for d_1600_i0_ in range(-403, hi11_):
            d_1601_v3_: _dafny.Map
            d_1601_v3_ = _dafny.Map({p2: p1})
            d_1602_v4_: _dafny.MultiSet
            d_1602_v4_ = _dafny.MultiSet([d_1600_i0_])
            d_1603_v5_: D0
            d_1603_v5_ = D0_DC1(p1)
            d_1604_v6_: _dafny.Map
            d_1604_v6_ = _dafny.Map({d_1597_v0_: d_1603_v5_})
            d_1605_v7_: _dafny.Seq
            d_1605_v7_ = _dafny.SeqWithoutIsStrInference([not((self).fm5(d_1600_i0_, p2, d_1604_v6_, p2, globalState))])
            d_1606_v8_: _dafny.Seq
            d_1606_v8_ = _dafny.SeqWithoutIsStrInference([len(d_1601_v3_)])
            d_1607_v9_: _dafny.Array
            nw256_ = _dafny.Array(int(0), 27)
            d_1607_v9_ = nw256_
            d_1608_v10_: _dafny.Map
            d_1608_v10_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1607_v9_])): p1})
            d_1609_v11_: _dafny.Seq
            d_1609_v11_ = _dafny.SeqWithoutIsStrInference([d_1602_v4_])
            d_1610_v12_: _dafny.Array
            nw257_ = _dafny.Array(None, 23)
            nw257_[int(0)] = (0) - (len((d_1601_v3_ if p2 else d_1601_v3_)))
            nw257_[int(1)] = (d_1602_v4_).cardinality
            nw257_[int(2)] = (p0) * (len(d_1605_v7_))
            nw257_[int(3)] = p0
            nw257_[int(4)] = len(d_1606_v8_)
            nw257_[int(5)] = 205
            nw257_[int(6)] = d_1600_i0_
            nw257_[int(7)] = (0) - (p0)
            nw257_[int(8)] = p0
            nw257_[int(9)] = p1
            nw257_[int(10)] = p1
            nw257_[int(11)] = p1
            nw257_[int(12)] = p1
            nw257_[int(13)] = -828
            nw257_[int(14)] = ((d_1608_v10_)[(0) - (d_1600_i0_)] if ((0) - (d_1600_i0_)) in (d_1608_v10_) else default__.fm20(False, globalState))
            nw257_[int(15)] = default__.safeDivisionInt(p0, d_1600_i0_)
            nw257_[int(16)] = 892
            nw257_[int(17)] = d_1600_i0_
            nw257_[int(18)] = (p1) * (d_1600_i0_)
            nw257_[int(19)] = p1
            nw257_[int(20)] = p1
            nw257_[int(21)] = d_1600_i0_
            nw257_[int(22)] = (0) - (len(_dafny.Map({779: len(d_1609_v11_)})))
            d_1610_v12_ = nw257_
            index262_ = default__.safeIndex(740, (d_1610_v12_).length(0))
            (d_1610_v12_)[index262_] = p1
            index263_ = default__.safeIndex(740, (d_1610_v12_).length(0))
            (d_1610_v12_)[index263_] = (0) - (p1)
            d_1597_v0_ = d_1597_v0_
            index264_ = default__.safeIndex(740, (d_1610_v12_).length(0))
            (d_1610_v12_)[index264_] = (0) - (p0)
            d_1611_v13_: _dafny.MultiSet
            d_1611_v13_ = _dafny.MultiSet([p2])
            r0 = (d_1611_v13_).issubset((d_1611_v13_).intersection(d_1611_v13_))
        d_1612_v14_: _dafny.Array
        def lambda69_(d_1613_p0_):
            def lambda70_(d_1614_i1_):
                return _dafny.MultiSet([(0) - (d_1613_p0_)])

            return lambda70_

        init37_ = lambda69_(p0)
        nw258_ = _dafny.Array(None, 27)
        for i0_37_ in range(nw258_.length(0)):
            nw258_[i0_37_] = init37_(i0_37_)
        d_1612_v14_ = nw258_
        d_1615_v15_: _dafny.MultiSet
        d_1615_v15_ = _dafny.MultiSet([p0])
        index265_ = default__.safeIndex(667, (d_1612_v14_).length(0))
        (d_1612_v14_)[index265_] = d_1615_v15_
        index266_ = default__.safeIndex(667, (d_1612_v14_).length(0))
        (d_1612_v14_)[index266_] = (d_1615_v15_ if p2 else (d_1615_v15_) - (d_1615_v15_))
        d_1616_v16_: _dafny.Seq
        d_1616_v16_ = _dafny.SeqWithoutIsStrInference([d_1597_v0_])
        d_1617_v17_: _dafny.Array
        def lambda71_(d_1618_v1_):
            def lambda72_(d_1619_i2_):
                return d_1618_v1_

            return lambda72_

        init38_ = lambda71_(d_1598_v1_)
        nw259_ = _dafny.Array(None, 3)
        for i0_38_ in range(nw259_.length(0)):
            nw259_[i0_38_] = init38_(i0_38_)
        d_1617_v17_ = nw259_
        d_1620_v18_: _dafny.Set
        d_1620_v18_ = _dafny.Set({p3, p3, p3})
        d_1621_v19_: _dafny.Seq
        d_1621_v19_ = _dafny.SeqWithoutIsStrInference([p0, len(d_1620_v18_), p0, p1, p1])
        d_1622_v20_: C1
        nw260_ = C1()
        nw260_.ctor__((p1) + (len(d_1616_v16_)), d_1617_v17_, (d_1621_v19_) + ((d_1621_v19_).set(default__.safeIndex(p1, len(d_1621_v19_)), 661)), D0_DC1(p1))
        d_1622_v20_ = nw260_
        d_1623_v21_: C2
        nw261_ = C2()
        nw261_.ctor__()
        d_1623_v21_ = nw261_
        d_1624_v22_: D0
        d_1624_v22_ = D0_DC1(d_1622_v20_.f9)
        d_1625_v23_: _dafny.Seq
        d_1625_v23_ = _dafny.SeqWithoutIsStrInference([d_1624_v22_])
        d_1626_v24_: _dafny.Map
        d_1626_v24_ = _dafny.Map({320: ((d_1625_v23_).set(default__.safeIndex(d_1622_v20_.f9, len(d_1625_v23_)), d_1624_v22_)) < (d_1625_v23_)})
        d_1626_v24_ = d_1626_v24_
        r0 = p2
        r1 = d_1622_v20_.f9
        return r0, r1


class C5(T1):
    def  __init__(self):
        self._f1: D0 = D0.default()()
        self._f0: _dafny.Seq = _dafny.Seq({})
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f1(self):
        return self._f1
    @f1.setter
    def f1(self, value):
        self._f1 = value
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f7, f0, f1):
        (self)._f7 = f7
        (self)._f0 = f0
        (self).f1 = f1

    def fm6(self, p0, p1, globalState):
        return 487

    def m7(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        d_1627_v0_: _dafny.Seq
        d_1627_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibpiei"))
        d_1628_v1_: str
        d_1628_v1_ = _dafny.CodePoint('c')
        d_1629_v2_: C0
        nw262_ = C0()
        nw262_.ctor__((default__.fm16(d_1627_v0_, p0, globalState)).set(default__.safeIndex(p0, len(default__.fm16(d_1627_v0_, p0, globalState))), d_1628_v1_))
        d_1629_v2_ = nw262_
        if p3:
            d_1630_v3_: _dafny.Set
            d_1630_v3_ = _dafny.Set({p3})
            index267_ = default__.safeIndex(80, ((self).f7).length(0))
            ((self).f7)[index267_] = ((self).f0)[default__.safeIndex(len(d_1630_v3_), len((self).f0))]
            d_1631_v4_: _dafny.Map
            d_1631_v4_ = _dafny.Map({p0: p0})
            index268_ = default__.safeIndex(80, ((self).f7).length(0))
            ((self).f7)[index268_] = (((d_1631_v4_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsieh")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsieh")))) in (d_1631_v4_) else p0)) - (333)
            if default__.fm17(globalState):
                r2 = p3
                d_1632_v5_: _dafny.Map
                d_1632_v5_ = _dafny.Map({((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]: p3})
                d_1633_v6_: _dafny.Seq
                d_1633_v6_ = _dafny.SeqWithoutIsStrInference([p3, ((d_1632_v5_)[((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]] if (((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]) in (d_1632_v5_) else p3)])
                d_1634_v7_: _dafny.Array
                def lambda73_(d_1635_p3_):
                    def lambda74_(d_1636_i0_):
                        return d_1635_p3_

                    return lambda74_

                init39_ = lambda73_(p3)
                nw263_ = _dafny.Array(None, 26)
                for i0_39_ in range(nw263_.length(0)):
                    nw263_[i0_39_] = init39_(i0_39_)
                d_1634_v7_ = nw263_
                d_1637_v8_: _dafny.Map
                d_1637_v8_ = _dafny.Map({d_1634_v7_: d_1633_v6_})
                d_1633_v6_ = (d_1633_v6_) + (((d_1637_v8_)[d_1634_v7_] if (d_1634_v7_) in (d_1637_v8_) else d_1633_v6_))
                r0 = default__.safeDivisionInt((-93) + (-786), ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))])
                index269_ = default__.safeIndex(80, ((self).f7).length(0))
                rhs179_ = p0
                rhs180_ = (((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]) - (p0)
                lhs110_ = (self).f7
                lhs111_ = default__.safeIndex(80, ((self).f7).length(0))
                lhs110_[lhs111_] = rhs179_
                r0 = rhs180_
                r0 = ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]
            elif True:
                index270_ = default__.safeIndex(80, ((self).f7).length(0))
                ((self).f7)[index270_] = (((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]) + (p0)
                d_1638_v9_: _dafny.Set
                d_1638_v9_ = _dafny.Set({_dafny.Set({p3}), d_1630_v3_})
                d_1639_v10_: _dafny.Array
                nw264_ = _dafny.Array(int(0), 11)
                d_1639_v10_ = nw264_
                def iife141_():
                    coll65_ = _dafny.Set()
                    compr_65_: _dafny.Set
                    for compr_65_ in (default__.fm18(globalState)).Elements:
                        d_1640_v11_: _dafny.Set = compr_65_
                        if (d_1640_v11_) in (default__.fm18(globalState)):
                            coll65_ = coll65_.union(_dafny.Set([d_1640_v11_]))
                    return _dafny.Set(coll65_)
                rhs181_ = default__.fm17(globalState)
                rhs182_ = iife141_()

                rhs183_ = d_1639_v10_
                rhs184_ = p3
                rhs185_ = not(p3)
                r2 = rhs181_
                d_1638_v9_ = rhs182_
                d_1639_v10_ = rhs183_
                r2 = rhs184_
                r2 = rhs185_
                r2 = p3
                d_1641_v12_: _dafny.Map
                d_1641_v12_ = _dafny.Map({p0: p3})
                d_1641_v12_ = (d_1641_v12_).set((self).fm6(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], p0, globalState), p3)
                d_1627_v0_ = (d_1629_v2_.f8) + ((d_1627_v0_) + (d_1629_v2_.f8))
            index271_ = default__.safeIndex(80, ((self).f7).length(0))
            ((self).f7)[index271_] = p0
            d_1642_v13_: _dafny.Seq
            d_1642_v13_ = _dafny.SeqWithoutIsStrInference([(self).f0])
            r2 = ((self).f0) not in (d_1642_v13_)
            if not(p3):
                d_1627_v0_ = d_1629_v2_.f8
                (self).f1 = self.f1
                r0 = default__.safeDivisionInt(202, p0)
                d_1643_v14_: _dafny.Map
                d_1643_v14_ = _dafny.Map({d_1628_v1_: p0})
                d_1644_v15_: _dafny.Map
                d_1644_v15_ = d_1643_v14_
                d_1645_v16_: _dafny.Map
                d_1645_v16_ = _dafny.Map({False: p3})
                d_1646_v17_: _dafny.Map
                d_1646_v17_ = _dafny.Map({default__.fm19(default__.fm17(globalState), p3, (d_1644_v15_), p3, globalState): ((d_1645_v16_)[p3] if (p3) in (d_1645_v16_) else not(False))})
                d_1646_v17_ = ((d_1646_v17_) | (d_1646_v17_)) | (d_1646_v17_)
                d_1647_v18_: _dafny.MultiSet
                d_1647_v18_ = _dafny.MultiSet([((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]])
                d_1648_v19_: _dafny.MultiSet
                d_1648_v19_ = _dafny.MultiSet([p3, p3, p3, not(p3), p3])
                d_1649_v20_: _dafny.Seq
                d_1649_v20_ = _dafny.SeqWithoutIsStrInference([(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]) >= (((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]), p3, (d_1648_v19_).issubset(d_1648_v19_)])
                d_1650_v21_: T0
                nw265_ = C4()
                nw265_.ctor__()
                d_1650_v21_ = nw265_
                d_1651_v22_: _dafny.Seq
                d_1651_v22_ = _dafny.SeqWithoutIsStrInference([d_1649_v20_])
                d_1652_v23_: _dafny.Seq
                d_1652_v23_ = _dafny.SeqWithoutIsStrInference([(d_1651_v22_)[default__.safeIndex(p0, len(d_1651_v22_))]])
                index272_ = default__.safeIndex(80, ((self).f7).length(0))
                rhs186_ = ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]
                rhs187_ = (d_1647_v18_) - (d_1647_v18_)
                rhs188_ = (d_1651_v22_)[default__.safeIndex(default__.fm32(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], globalState), len(d_1651_v22_))]
                rhs189_ = d_1650_v21_
                rhs190_ = ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]
                lhs112_ = (self).f7
                lhs113_ = default__.safeIndex(80, ((self).f7).length(0))
                r0 = rhs186_
                d_1647_v18_ = rhs187_
                d_1649_v20_ = rhs188_
                d_1650_v21_ = rhs189_
                lhs112_[lhs113_] = rhs190_
            elif True:
                d_1653_v24_: C2
                nw266_ = C2()
                nw266_.ctor__()
                d_1653_v24_ = nw266_
                d_1654_v25_: _dafny.Map
                d_1654_v25_ = _dafny.Map({((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]: p3})
                d_1655_v26_: _dafny.Seq
                d_1655_v26_ = _dafny.SeqWithoutIsStrInference([d_1654_v25_])
                d_1656_v27_: _dafny.Map
                d_1656_v27_ = _dafny.Map({(d_1655_v26_)[default__.safeIndex(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], len(d_1655_v26_))]: p0})
                d_1657_v28_: _dafny.Map
                d_1657_v28_ = _dafny.Map({d_1628_v1_: self.f1})
                d_1658_v29_: _dafny.Map
                d_1658_v29_ = _dafny.Map({p3: (d_1653_v24_).fm5(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], False, d_1657_v28_, p3, globalState)})
                d_1659_v30_: _dafny.Set
                d_1659_v30_ = _dafny.Set({(self).f0, _dafny.SeqWithoutIsStrInference([((self).f7)[default__.safeIndex(80, ((self).f7).length(0))] for d_1660_i1_ in range(default__.abs(-806))])})
                d_1661_v31_: _dafny.MultiSet
                d_1661_v31_ = _dafny.MultiSet([p3, True, p3, p3])
                index273_ = default__.safeIndex(80, ((self).f7).length(0))
                index274_ = default__.safeIndex(80, ((self).f7).length(0))
                rhs191_ = ((p0) + (174)) - (((d_1656_v27_)[_dafny.Map({((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]: p3})] if (_dafny.Map({((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]: p3})) in (d_1656_v27_) else len(d_1658_v29_)))
                rhs192_ = d_1629_v2_.f8
                rhs193_ = not(((default__.fm58(D16_DC42(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], p0, p0), True, d_1659_v30_, globalState)).issubset(_dafny.MultiSet([d_1628_v1_, d_1628_v1_])) if p3 else p3))
                rhs194_ = ((d_1661_v31_).cardinality) * (p0)
                rhs195_ = p3
                lhs114_ = (self).f7
                lhs115_ = default__.safeIndex(80, ((self).f7).length(0))
                lhs116_ = (self).f7
                lhs117_ = default__.safeIndex(80, ((self).f7).length(0))
                lhs114_[lhs115_] = rhs191_
                d_1627_v0_ = rhs192_
                r2 = rhs193_
                lhs116_[lhs117_] = rhs194_
                r2 = rhs195_
                index275_ = default__.safeIndex(80, ((self).f7).length(0))
                ((self).f7)[index275_] = ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]
                r2 = (d_1627_v0_) != (d_1629_v2_.f8)
                d_1662_v32_: _dafny.Array
                def lambda75_(d_1663_p3_):
                    def lambda76_(d_1664_i2_):
                        return (d_1663_p3_) or (False)

                    return lambda76_

                init40_ = lambda75_(p3)
                nw267_ = _dafny.Array(None, 15)
                for i0_40_ in range(nw267_.length(0)):
                    nw267_[i0_40_] = init40_(i0_40_)
                d_1662_v32_ = nw267_
                d_1665_v33_: _dafny.Map
                d_1665_v33_ = _dafny.Map({d_1628_v1_: 583})
                d_1666_v34_: _dafny.Set
                d_1666_v34_ = _dafny.Set({d_1665_v33_, d_1665_v33_})
                d_1667_v36_: _dafny.Map
                d_1667_v36_ = _dafny.Map({d_1628_v1_: True})
                index276_ = default__.safeIndex(536, (d_1662_v32_).length(0))
                def iife142_():
                    coll66_ = _dafny.Map()
                    compr_66_: str
                    for compr_66_ in (d_1667_v36_).keys.Elements:
                        d_1668_v35_: str = compr_66_
                        if (d_1668_v35_) in (d_1667_v36_):
                            coll66_[d_1668_v35_] = ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))]
                    return _dafny.Map(coll66_)
                (d_1662_v32_)[index276_] = (d_1666_v34_) == (_dafny.Set({d_1665_v33_, (iife142_()
                ).set(d_1628_v1_, (self).fm6(((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], ((self).f7)[default__.safeIndex(80, ((self).f7).length(0))], globalState))}))
                index277_ = default__.safeIndex(536, (d_1662_v32_).length(0))
                (d_1662_v32_)[index277_] = p3
        elif True:
            r2 = (p0) >= (p0)
            d_1669_v37_: _dafny.Seq
            d_1669_v37_ = _dafny.SeqWithoutIsStrInference([d_1629_v2_.f8])
            d_1670_v38_: _dafny.Seq
            d_1670_v38_ = _dafny.SeqWithoutIsStrInference([(d_1669_v37_)[default__.safeIndex(p0, len(d_1669_v37_))]])
            d_1671_v39_: _dafny.Map
            d_1671_v39_ = _dafny.Map({((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0, p0]))).cardinality) - (len(d_1669_v37_)): (_dafny.SeqWithoutIsStrInference([((self).f0)[default__.safeIndex(len((self).f0), len((self).f0))], p0, p0, p0, len(p1)])) + ((self).f0)})
            d_1671_v39_ = (d_1671_v39_).set(p0, (self).f0)
            r2 = p3
            r2 = not((p0) != ((0) - (p0)))
            d_1672_v40_: _dafny.Set
            d_1672_v40_ = _dafny.Set({p0})
            d_1672_v40_ = d_1672_v40_
        d_1673_v41_: _dafny.Seq
        d_1673_v41_ = _dafny.SeqWithoutIsStrInference([p3])
        rhs196_ = p3
        rhs197_ = (p3 if p3 else not((d_1673_v41_) < (default__.fm19(p3, not(False), _dafny.Map({d_1628_v1_: len(d_1629_v2_.f8)}), p3, globalState))))
        rhs198_ = p0
        rhs199_ = (0) - ((p0) * (p0))
        r2 = rhs196_
        r2 = rhs197_
        r0 = rhs198_
        r0 = rhs199_
        d_1674_v42_: _dafny.Map
        d_1674_v42_ = _dafny.Map({(0) - (p0): p3})
        d_1674_v42_ = (d_1674_v42_) | (d_1674_v42_)
        d_1675_v43_: D12
        d_1675_v43_ = D12_DC28(p3, p0)
        d_1675_v43_ = d_1675_v43_
        hi12_ = p0
        for d_1676_i3_ in range(p0, hi12_):
            (d_1629_v2_).f8 = d_1627_v0_
            r0 = p0
            d_1677_v44_: _dafny.Map
            d_1677_v44_ = _dafny.Map({p3: p3})
            d_1677_v44_ = (d_1677_v44_).set((p0) < (p0), p3)
            d_1678_v45_: D11
            d_1678_v45_ = D11_DC25(d_1628_v1_, len(d_1629_v2_.f8), True)
            d_1679_v46_: _dafny.Array
            nw268_ = _dafny.Array(None, 21)
            nw268_[int(0)] = p3
            nw268_[int(1)] = p3
            nw268_[int(2)] = p3
            nw268_[int(3)] = p3
            nw268_[int(4)] = p3
            nw268_[int(5)] = p3
            nw268_[int(6)] = p3
            nw268_[int(7)] = p3
            nw268_[int(8)] = True
            nw268_[int(9)] = p3
            nw268_[int(10)] = p3
            nw268_[int(11)] = (d_1678_v45_).cf43
            nw268_[int(12)] = p3
            nw268_[int(13)] = p3
            nw268_[int(14)] = p3
            nw268_[int(15)] = not(p3)
            nw268_[int(16)] = p3
            nw268_[int(17)] = p3
            nw268_[int(18)] = p3
            nw268_[int(19)] = True
            nw268_[int(20)] = p3
            d_1679_v46_ = nw268_
            d_1680_v47_: _dafny.Set
            d_1680_v47_ = _dafny.Set({d_1679_v46_})
            d_1681_v48_: _dafny.MultiSet
            d_1681_v48_ = _dafny.MultiSet([len(d_1680_v47_), d_1676_i3_, 239, p0, d_1676_i3_])
            d_1682_v49_: _dafny.Set
            d_1682_v49_ = _dafny.Set({p3, (d_1681_v48_) == (d_1681_v48_)})
            d_1682_v49_ = d_1682_v49_
        d_1683_v50_: _dafny.MultiSet
        d_1683_v50_ = _dafny.MultiSet([len(d_1627_v0_), (0) - (p0)])
        r0 = ((d_1683_v50_) | ((d_1683_v50_) | (d_1683_v50_))).cardinality
        r1 = default__.fm59(globalState)
        r2 = True
        return r0, r1, r2

    def m8(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: _dafny.Seq = _dafny.Seq({})
        d_1684_v0_: _dafny.Seq
        d_1684_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fi"))
        d_1685_v1_: C0
        nw269_ = C0()
        nw269_.ctor__(d_1684_v0_)
        d_1685_v1_ = nw269_
        d_1686_v2_: str
        d_1686_v2_ = _dafny.CodePoint('j')
        d_1687_v3_: int
        d_1687_v3_ = -151
        d_1688_v4_: D16
        d_1688_v4_ = D16_DC43(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekkqo")), d_1686_v2_, d_1687_v3_)
        pat_let_tv42_ = p1
        pat_let_tv43_ = p1
        pat_let_tv44_ = p1
        pat_let_tv45_ = p3
        def lambda77_(source18_):
            if source18_.is_DC42:
                d_1689___mcc_h0_ = source18_.cf76
                d_1690___mcc_h1_ = source18_.cf77
                d_1691___mcc_h2_ = source18_.cf78
                d_1692_cf78_ = d_1691___mcc_h2_
                d_1693_cf77_ = d_1690___mcc_h1_
                d_1694_cf76_ = d_1689___mcc_h0_
                return pat_let_tv42_
            elif source18_.is_DC43:
                d_1695___mcc_h3_ = source18_.cf79
                d_1696___mcc_h4_ = source18_.cf80
                d_1697___mcc_h5_ = source18_.cf81
                d_1698_cf81_ = d_1697___mcc_h5_
                d_1699_cf80_ = d_1696___mcc_h4_
                d_1700_cf79_ = d_1695___mcc_h3_
                return pat_let_tv43_
            elif source18_.is_DC44:
                d_1701___mcc_h6_ = source18_.cf82
                d_1702_cf82_ = d_1701___mcc_h6_
                return pat_let_tv44_
            elif True:
                d_1703___mcc_h7_ = source18_.cf75
                d_1704_cf75_ = d_1703___mcc_h7_
                return pat_let_tv45_

        r0 = not(lambda77_(d_1688_v4_))
        hi13_ = default__.fm44(p0, globalState)
        for d_1705_i0_ in range((0) - (len(d_1684_v0_)), hi13_):
            d_1706_v5_: C0
            nw270_ = C0()
            nw270_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))
            d_1706_v5_ = nw270_
            d_1687_v3_ = len((d_1706_v5_.f8 if p1 else (d_1685_v1_.f8) + (_dafny.SeqWithoutIsStrInference([d_1686_v2_ for d_1707_i1_ in range(default__.abs(386))]))))
            r0 = False
            d_1708_v6_: _dafny.Seq
            d_1708_v6_ = _dafny.SeqWithoutIsStrInference([True])
            d_1709_v7_: D6
            d_1709_v7_ = D6_DC15(len(_dafny.SeqWithoutIsStrInference([d_1686_v2_ for d_1710_i2_ in range(default__.abs(-836))])), len(d_1708_v6_), p3)
            d_1711_v8_: _dafny.Map
            d_1711_v8_ = _dafny.Map({p1: d_1709_v7_})
            source19_ = D6_DC16(((d_1711_v8_)[p3] if (p3) in (d_1711_v8_) else d_1709_v7_))
            if source19_.is_DC15:
                d_1712___mcc_h8_ = source19_.cf30
                d_1713___mcc_h9_ = source19_.cf31
                d_1714___mcc_h10_ = source19_.cf32
                d_1715_cf32_ = d_1714___mcc_h10_
                d_1716_cf31_ = d_1713___mcc_h9_
                d_1717_cf30_ = d_1712___mcc_h8_
                rhs200_ = ((d_1717_cf30_) * (d_1717_cf30_)) > ((d_1716_cf31_) * (263))
                rhs201_ = (True) or (p0)
                rhs202_ = (191) - (d_1687_v3_)
                rhs203_ = (0) - (d_1717_cf30_)
                r0 = rhs200_
                r0 = rhs201_
                d_1687_v3_ = rhs202_
                d_1717_cf30_ = rhs203_
                d_1718_v9_: _dafny.Array
                nw271_ = _dafny.Array(int(0), 6)
                d_1718_v9_ = nw271_
                d_1719_v13_: _dafny.Array
                def lambda78_(d_1720_i0_, d_1721_cf31_):
                    def lambda79_(d_1722_i3_):
                        def iife143_():
                            coll67_ = _dafny.Set()
                            compr_67_: int
                            for compr_67_ in _dafny.IntegerRange(892, 820):
                                d_1723_v10_: int = compr_67_
                                if ((892) <= (d_1723_v10_)) and ((d_1723_v10_) < (820)):
                                    coll67_ = coll67_.union(_dafny.Set([(d_1723_v10_) - (d_1720_i0_)]))
                            return _dafny.Set(coll67_)
                        def iife144_():
                            def iife146_():
                                coll70_ = _dafny.Map()
                                compr_70_: int
                                for compr_70_ in (_dafny.Set({d_1720_i0_})).Elements:
                                    d_1724_v11_: int = compr_70_
                                    if (d_1724_v11_) in (_dafny.Set({d_1720_i0_})):
                                        coll70_[default__.safeDivisionInt(d_1724_v11_, d_1721_cf31_)] = 49
                                return _dafny.Map(coll70_)
                            coll68_ = _dafny.Set()
                            def iife145_():
                                coll69_ = _dafny.Map()
                                compr_69_: int
                                for compr_69_ in (_dafny.Set({d_1720_i0_})).Elements:
                                    d_1724_v11_: int = compr_69_
                                    if (d_1724_v11_) in (_dafny.Set({d_1720_i0_})):
                                        coll69_[default__.safeDivisionInt(d_1724_v11_, d_1721_cf31_)] = 49
                                return _dafny.Map(coll69_)
                            compr_68_: int
                            for compr_68_ in (iife145_()
                            ).keys.Elements:
                                d_1725_v12_: int = compr_68_
                                if (d_1725_v12_) in (iife146_()
                                ):
                                    coll68_ = coll68_.union(_dafny.Set([default__.safeDivisionInt(d_1725_v12_, 793)]))
                            return _dafny.Set(coll68_)
                        return (iife143_()
                        ) != (iife144_()
                        )

                    return lambda79_

                init41_ = lambda78_(d_1705_i0_, d_1716_cf31_)
                nw272_ = _dafny.Array(None, 16)
                for i0_41_ in range(nw272_.length(0)):
                    nw272_[i0_41_] = init41_(i0_41_)
                d_1719_v13_ = nw272_
                index278_ = default__.safeIndex(311, (d_1719_v13_).length(0))
                (d_1719_v13_)[index278_] = default__.fm17(globalState)
                d_1726_v14_: _dafny.Set
                d_1726_v14_ = _dafny.Set({d_1687_v3_, -459})
                index279_ = default__.safeIndex(311, (d_1719_v13_).length(0))
                rhs204_ = (459) == (len((_dafny.SeqWithoutIsStrInference([d_1715_cf32_, not(p0), d_1715_cf32_])).set(default__.safeIndex((0) - (len(d_1726_v14_)), len(_dafny.SeqWithoutIsStrInference([d_1715_cf32_, not(p0), d_1715_cf32_]))), p3)))
                rhs205_ = d_1687_v3_
                rhs206_ = d_1718_v9_
                rhs207_ = p3
                lhs118_ = d_1719_v13_
                lhs119_ = default__.safeIndex(311, (d_1719_v13_).length(0))
                d_1715_cf32_ = rhs204_
                d_1687_v3_ = rhs205_
                d_1718_v9_ = rhs206_
                lhs118_[lhs119_] = rhs207_
                d_1715_cf32_ = (((self).f0) + ((self).f0)) not in ((((_dafny.SeqWithoutIsStrInference([(self).f0])).set(default__.safeIndex(default__.fm20((d_1719_v13_)[default__.safeIndex(311, (d_1719_v13_).length(0))], globalState), len(_dafny.SeqWithoutIsStrInference([(self).f0]))), (self).f0)).set(default__.safeIndex(d_1716_cf31_, len((_dafny.SeqWithoutIsStrInference([(self).f0])).set(default__.safeIndex(default__.fm20((d_1719_v13_)[default__.safeIndex(311, (d_1719_v13_).length(0))], globalState), len(_dafny.SeqWithoutIsStrInference([(self).f0]))), (self).f0))), _dafny.SeqWithoutIsStrInference([d_1705_i0_]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1687_v3_]), ((self).f0).set(default__.safeIndex(d_1687_v3_, len((self).f0)), len(_dafny.SeqWithoutIsStrInference([d_1686_v2_ for d_1727_i4_ in range(default__.abs(906))])))])))
                r0 = (d_1686_v2_) in (_dafny.SeqWithoutIsStrInference([d_1686_v2_ for d_1728_i5_ in range(default__.abs(-884))]))
            elif source19_.is_DC14:
                d_1729___mcc_h11_ = source19_.cf29
                d_1730_cf29_ = d_1729___mcc_h11_
                r2 = ((self).f0) + ((((self).f0) + ((self).f0)).set(default__.safeIndex(((self).f0)[default__.safeIndex(default__.fm44(default__.fm17(globalState), globalState), len((self).f0))], len(((self).f0) + ((self).f0))), d_1687_v3_))
                d_1731_v15_: C4
                nw273_ = C4()
                nw273_.ctor__()
                d_1731_v15_ = nw273_
                d_1686_v2_ = d_1686_v2_
                d_1732_v16_: C0
                nw274_ = C0()
                nw274_.ctor__(d_1684_v0_)
                d_1732_v16_ = nw274_
            elif True:
                d_1733___mcc_h12_ = source19_.cf33
                d_1734_cf33_ = d_1733___mcc_h12_
                d_1735_v17_: _dafny.Array
                nw275_ = _dafny.Array(False, 15)
                d_1735_v17_ = nw275_
                index280_ = default__.safeIndex(971, (d_1735_v17_).length(0))
                (d_1735_v17_)[index280_] = p1
                index281_ = default__.safeIndex(971, (d_1735_v17_).length(0))
                (d_1735_v17_)[index281_] = p3
                d_1687_v3_ = (-536) + (d_1687_v3_)
                index282_ = default__.safeIndex(971, (d_1735_v17_).length(0))
                (d_1735_v17_)[index282_] = ((0) - (default__.fm44(p1, globalState))) in (_dafny.Set({-289}))
                d_1735_v17_ = d_1735_v17_
        d_1736_i6_: int
        d_1736_i6_ = 0
        with _dafny.label("5"):
            while False:
                with _dafny.c_label("5"):
                    if (d_1736_i6_) >= (100):
                        raise _dafny.Break("5")
                    d_1736_i6_ = (d_1736_i6_) + (1)
                    r0 = p0
                    if not(p0):
                        d_1737_v18_: _dafny.Array
                        nw276_ = _dafny.Array(int(0), 25)
                        d_1737_v18_ = nw276_
                        d_1737_v18_ = (self).f7
                        d_1738_v19_: _dafny.Set
                        d_1738_v19_ = _dafny.Set({d_1687_v3_})
                        r0 = (d_1738_v19_).ispropersubset(d_1738_v19_)
                        d_1687_v3_ = (0) - ((default__.fm44(p1, globalState)) + (d_1687_v3_))
                        index283_ = default__.safeIndex(12, (d_1737_v18_).length(0))
                        (d_1737_v18_)[index283_] = d_1687_v3_
                        index284_ = default__.safeIndex(12, (d_1737_v18_).length(0))
                        (d_1737_v18_)[index284_] = default__.fm20(p1, globalState)
                        d_1739_v20_: _dafny.Array
                        def lambda80_(d_1740_v2_, d_1741_v3_):
                            def lambda81_(d_1742_i8_):
                                return _dafny.Map({d_1740_v2_: d_1741_v3_})

                            return lambda81_

                        init42_ = lambda80_(d_1686_v2_, d_1687_v3_)
                        nw277_ = _dafny.Array(None, 16)
                        for i0_42_ in range(nw277_.length(0)):
                            nw277_[i0_42_] = init42_(i0_42_)
                        d_1739_v20_ = nw277_
                        d_1743_v21_: C1
                        nw278_ = C1()
                        nw278_.ctor__(len((_dafny.SeqWithoutIsStrInference([d_1686_v2_ for d_1744_i7_ in range(default__.abs(966))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skvawdoq")))), d_1739_v20_, (self).f0, self.f1)
                        d_1743_v21_ = nw278_
                    elif True:
                        d_1745_v22_: _dafny.Array
                        nw279_ = _dafny.Array(_dafny.Array(None, 0), 5)
                        d_1745_v22_ = nw279_
                        d_1746_v23_: _dafny.Array
                        nw280_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                        d_1746_v23_ = nw280_
                        index285_ = default__.safeIndex(170, (d_1745_v22_).length(0))
                        (d_1745_v22_)[index285_] = d_1746_v23_
                        index286_ = default__.safeIndex(170, (d_1745_v22_).length(0))
                        (d_1745_v22_)[index286_] = d_1746_v23_
                        d_1747_v24_: _dafny.Set
                        d_1747_v24_ = _dafny.Set({d_1687_v3_})
                        d_1748_v25_: _dafny.Map
                        d_1748_v25_ = _dafny.Map({d_1747_v24_: (_dafny.SeqWithoutIsStrInference([d_1687_v3_ for d_1749_i9_ in range(default__.abs(408))])) + ((self).f0)})
                        d_1750_v26_: _dafny.Map
                        d_1750_v26_ = _dafny.Map({d_1687_v3_: d_1747_v24_})
                        d_1751_v27_: _dafny.Map
                        d_1751_v27_ = _dafny.Map({d_1687_v3_: p1})
                        def iife147_():
                            coll71_ = _dafny.Set()
                            compr_71_: int
                            for compr_71_ in _dafny.IntegerRange(-23, -358):
                                d_1752_v28_: int = compr_71_
                                if ((-23) <= (d_1752_v28_)) and ((d_1752_v28_) < (-358)):
                                    coll71_ = coll71_.union(_dafny.Set([default__.safeModuloInt(d_1752_v28_, d_1687_v3_)]))
                            return _dafny.Set(coll71_)
                        d_1748_v25_ = (d_1748_v25_).set(((d_1750_v26_)[len(d_1751_v27_)] if (len(d_1751_v27_)) in (d_1750_v26_) else iife147_()
                        ), (self).f0)
                        d_1753_v29_: _dafny.Array
                        nw281_ = _dafny.Array(_dafny.Set({}), 22)
                        d_1753_v29_ = nw281_
                        d_1754_v30_: _dafny.Set
                        d_1754_v30_ = _dafny.Set({p1, p0})
                        index287_ = default__.safeIndex(877, (d_1753_v29_).length(0))
                        (d_1753_v29_)[index287_] = d_1754_v30_
                        index288_ = default__.safeIndex(877, (d_1753_v29_).length(0))
                        (d_1753_v29_)[index288_] = d_1754_v30_
                        d_1755_v31_: _dafny.Map
                        d_1755_v31_ = _dafny.Map({False: d_1687_v3_})
                        d_1756_v32_: _dafny.Set
                        d_1756_v32_ = _dafny.Set({d_1755_v31_, d_1755_v31_})
                        d_1756_v32_ = (d_1756_v32_ if p3 else d_1756_v32_)
                        index289_ = default__.safeIndex(125, ((self).f7).length(0))
                        ((self).f7)[index289_] = 136
                        d_1757_v33_: D11
                        d_1757_v33_ = D11_DC25(d_1686_v2_, d_1687_v3_, not(p1))
                        d_1758_v34_: _dafny.Map
                        d_1758_v34_ = _dafny.Map({d_1757_v33_: -795})
                        d_1759_v36_: _dafny.Map
                        d_1759_v36_ = _dafny.Map({d_1757_v33_: p1})
                        d_1760_v37_: _dafny.MultiSet
                        def iife148_():
                            coll72_ = _dafny.Map()
                            compr_72_: D11
                            for compr_72_ in (d_1759_v36_).keys.Elements:
                                d_1761_v35_: D11 = compr_72_
                                if (d_1761_v35_) in (d_1759_v36_):
                                    coll72_[d_1761_v35_] = d_1687_v3_
                            return _dafny.Map(coll72_)
                        d_1760_v37_ = _dafny.MultiSet([(d_1758_v34_) | (d_1758_v34_), (d_1758_v34_) | (iife148_()
                        )])
                        index290_ = default__.safeIndex(125, ((self).f7).length(0))
                        ((self).f7)[index290_] = ((d_1760_v37_)[(default__.fm60(d_1687_v3_, p1, globalState)) | (_dafny.Map({d_1757_v33_: d_1687_v3_}))] if ((default__.fm60(d_1687_v3_, p1, globalState)) | (_dafny.Map({d_1757_v33_: d_1687_v3_}))) in (d_1760_v37_) else len(d_1684_v0_))
                    d_1687_v3_ = (d_1687_v3_) - (default__.safeDivisionInt(d_1687_v3_, d_1687_v3_))
                    d_1762_v38_: _dafny.Array
                    nw282_ = _dafny.Array(False, 8)
                    d_1762_v38_ = nw282_
                    index291_ = default__.safeIndex(153, (d_1762_v38_).length(0))
                    (d_1762_v38_)[index291_] = p0
                    index292_ = default__.safeIndex(153, (d_1762_v38_).length(0))
                    (d_1762_v38_)[index292_] = (p1 if True else p3)
                    pass
            pass
        d_1763_v39_: _dafny.Seq
        d_1763_v39_ = _dafny.SeqWithoutIsStrInference([default__.fm30(d_1687_v3_, p1, p1, globalState)])
        d_1764_v40_: _dafny.Map
        d_1764_v40_ = _dafny.Map({len(d_1763_v39_): d_1686_v2_})
        d_1764_v40_ = (d_1764_v40_).set((d_1687_v3_) + ((0) - (d_1687_v3_)), (d_1686_v2_ if True else d_1686_v2_))
        index293_ = default__.safeIndex(934, ((self).f7).length(0))
        ((self).f7)[index293_] = (d_1687_v3_) * (446)
        index294_ = default__.safeIndex(934, ((self).f7).length(0))
        ((self).f7)[index294_] = d_1687_v3_
        d_1765_v41_: D2
        d_1765_v41_ = D2_DC5(600, (self).f0, False)
        r0 = (d_1765_v41_).cf10
        r1 = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhwdeveiq")): p1})
        r2 = (self).f0
        return r0, r1, r2

    @property
    def f7(self):
        return self._f7

class C6:
    def  __init__(self):
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm11(self, p0, globalState):
        return 261

    def fm12(self, p0, p1, globalState):
        return 409

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        pat_let_tv46_ = p1
        def iife149_(_pat_let38_0):
            def iife150_(d_1766_dt__update__tmp_h0_):
                def iife151_(_pat_let39_0):
                    def iife152_(d_1767_dt__update_hcf2_h0_):
                        return D0_DC1(d_1767_dt__update_hcf2_h0_)
                    return iife152_(_pat_let39_0)
                return iife151_(len(_dafny.Set({(self).f6, len(pat_let_tv46_)})))
            return iife150_(_pat_let38_0)
        source20_ = iife149_(D0_DC1(p3))
        if source20_.is_DC0:
            d_1768___mcc_h0_ = source20_.cf0
            d_1769___mcc_h1_ = source20_.cf1
            d_1770_cf1_ = d_1769___mcc_h1_
            d_1771_cf0_ = d_1768___mcc_h0_
            d_1772_v0_: bool
            d_1772_v0_ = False
            if d_1772_v0_:
                d_1773_v1_: _dafny.Seq
                d_1773_v1_ = _dafny.SeqWithoutIsStrInference([(d_1772_v0_ if d_1772_v0_ else d_1772_v0_), d_1772_v0_, (p1) <= (p1), d_1772_v0_])
                d_1773_v1_ = d_1773_v1_
                d_1774_v2_: _dafny.Seq
                d_1774_v2_ = _dafny.SeqWithoutIsStrInference([d_1770_cf1_])
                r3 = (d_1774_v2_)[default__.safeIndex((self).fm12(d_1772_v0_, d_1772_v0_, globalState), len(d_1774_v2_))]
                d_1770_cf1_ = p3
                d_1775_v3_: _dafny.MultiSet
                d_1775_v3_ = _dafny.MultiSet([((self).f6) * (len(d_1773_v1_))])
                d_1776_v4_: str
                d_1776_v4_ = _dafny.CodePoint('r')
                d_1777_v5_: _dafny.Map
                d_1777_v5_ = _dafny.Map({(0) - ((self).f6): d_1772_v0_})
                d_1778_v6_: _dafny.MultiSet
                d_1778_v6_ = _dafny.MultiSet([False])
                d_1779_v7_: _dafny.Map
                d_1779_v7_ = _dafny.Map({(d_1778_v6_).cardinality: (self).f6})
                d_1780_v8_: _dafny.Seq
                d_1780_v8_ = _dafny.SeqWithoutIsStrInference([d_1779_v7_, d_1779_v7_])
                d_1775_v3_ = (default__.fm13(d_1776_v4_, not(((d_1777_v5_)[p3] if (p3) in (d_1777_v5_) else d_1772_v0_)), globalState)).set(111, default__.abs((p3) - (len(d_1780_v8_))))
                d_1781_v9_: C2
                nw283_ = C2()
                nw283_.ctor__()
                d_1781_v9_ = nw283_
            elif True:
                d_1782_v10_: _dafny.Array
                nw284_ = _dafny.Array(_dafny.Map({}), 28)
                d_1782_v10_ = nw284_
                d_1783_v11_: _dafny.Map
                d_1783_v11_ = _dafny.Map({not(d_1772_v0_): d_1772_v0_})
                index295_ = default__.safeIndex(55, (d_1782_v10_).length(0))
                (d_1782_v10_)[index295_] = (d_1783_v11_) | (d_1783_v11_)
                d_1784_v12_: _dafny.Seq
                d_1784_v12_ = _dafny.SeqWithoutIsStrInference([d_1772_v0_, d_1772_v0_])
                index296_ = default__.safeIndex(55, (d_1782_v10_).length(0))
                (d_1782_v10_)[index296_] = (d_1783_v11_) | (_dafny.Map({(d_1784_v12_)[default__.safeIndex((self).f6, len(d_1784_v12_))]: d_1772_v0_}))
                d_1785_v13_: _dafny.Map
                d_1785_v13_ = _dafny.Map({d_1772_v0_: (self).f6})
                d_1785_v13_ = (d_1785_v13_).set((d_1784_v12_)[default__.safeIndex(918, len(d_1784_v12_))], d_1770_cf1_)
                d_1786_v14_: _dafny.Array
                nw285_ = _dafny.Array(_dafny.Map({}), 5)
                d_1786_v14_ = nw285_
                d_1787_v15_: str
                d_1787_v15_ = _dafny.CodePoint('s')
                d_1788_v16_: _dafny.MultiSet
                d_1788_v16_ = _dafny.MultiSet([_dafny.CodePoint('t'), d_1787_v15_])
                d_1789_v17_: _dafny.Seq
                d_1789_v17_ = _dafny.SeqWithoutIsStrInference([p3, (0) - ((d_1788_v16_).cardinality), (self).f6])
                d_1790_v18_: C1
                nw286_ = C1()
                nw286_.ctor__(len(d_1785_v13_), d_1786_v14_, d_1789_v17_, p0)
                d_1790_v18_ = nw286_
                d_1791_v19_: _dafny.Map
                d_1791_v19_ = _dafny.Map({d_1790_v18_: d_1790_v18_.f9})
                d_1770_cf1_ = ((self).fm12(d_1772_v0_, not(d_1772_v0_), globalState)) + (((d_1791_v19_)[d_1790_v18_] if (d_1790_v18_) in (d_1791_v19_) else (self).f6))
                (d_1790_v18_).f9 = (d_1790_v18_).fm6(p3, d_1770_cf1_, globalState)
                d_1792_v20_: _dafny.Array
                def lambda82_(d_1793_v18_, d_1794_p1_, d_1795_v0_):
                    def lambda83_(d_1796_i0_):
                        return ((_dafny.MultiSet([d_1793_v18_.f9, (self).f6])).cardinality) < ((D4_DC11(len(d_1794_p1_), False, d_1795_v0_, d_1795_v0_, not(d_1795_v0_))).cf22)

                    return lambda83_

                init43_ = lambda82_(d_1790_v18_, p1, d_1772_v0_)
                nw287_ = _dafny.Array(None, 3)
                for i0_43_ in range(nw287_.length(0)):
                    nw287_[i0_43_] = init43_(i0_43_)
                d_1792_v20_ = nw287_
                index297_ = default__.safeIndex(276, (d_1792_v20_).length(0))
                (d_1792_v20_)[index297_] = d_1772_v0_
                index298_ = default__.safeIndex(276, (d_1792_v20_).length(0))
                (d_1792_v20_)[index298_] = d_1772_v0_
            d_1797_v21_: _dafny.Map
            d_1797_v21_ = _dafny.Map({p0: (self).f6})
            d_1797_v21_ = (d_1797_v21_).set(p0, 509)
            d_1798_v22_: str
            d_1798_v22_ = _dafny.CodePoint('i')
            d_1799_v23_: D16
            d_1799_v23_ = D16_DC43(p1, _dafny.CodePoint('u'), -656)
            d_1800_v24_: _dafny.Seq
            d_1800_v24_ = _dafny.SeqWithoutIsStrInference([p1, p1, (d_1799_v23_).cf79])
            d_1801_v25_: _dafny.Seq
            d_1801_v25_ = _dafny.SeqWithoutIsStrInference([(d_1800_v24_)[default__.safeIndex(p3, len(d_1800_v24_))]])
            d_1802_v26_: _dafny.Map
            d_1802_v26_ = _dafny.Map({p1: d_1772_v0_})
            d_1803_v27_: _dafny.Map
            d_1803_v27_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([p1, p1, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nydykbd"))).set(default__.safeIndex(-926, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nydykbd")))), d_1798_v22_), p1, p1])) + (d_1801_v25_): d_1802_v26_})
            d_1803_v27_ = (d_1803_v27_).set(d_1800_v24_, d_1802_v26_)
            d_1804_v28_: C3
            nw288_ = C3()
            nw288_.ctor__()
            d_1804_v28_ = nw288_
        elif True:
            d_1805___mcc_h2_ = source20_.cf2
            d_1806_cf2_ = d_1805___mcc_h2_
            d_1807_v29_: C0
            nw289_ = C0()
            nw289_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvlu")))
            d_1807_v29_ = nw289_
            d_1808_v30_: _dafny.Map
            d_1808_v30_ = _dafny.Map({461: True})
            d_1809_v31_: _dafny.Map
            d_1809_v31_ = _dafny.Map({(self).f6: not(((d_1808_v30_)[207] if (207) in (d_1808_v30_) else True))})
            d_1810_v32_: bool
            d_1810_v32_ = True
            r1 = ((d_1808_v30_) == (d_1809_v31_)) == (d_1810_v32_)
            r1 = not ((d_1810_v32_) and (d_1810_v32_)) or (d_1810_v32_)
            r2 = d_1810_v32_
        d_1811_v33_: bool
        d_1811_v33_ = True
        d_1812_v34_: _dafny.MultiSet
        d_1812_v34_ = _dafny.MultiSet([d_1811_v33_])
        d_1813_v35_: _dafny.Seq
        d_1813_v35_ = _dafny.SeqWithoutIsStrInference([d_1812_v34_, d_1812_v34_])
        d_1814_v36_: _dafny.Map
        d_1814_v36_ = _dafny.Map({p3: d_1811_v33_})
        d_1815_v37_: _dafny.MultiSet
        d_1815_v37_ = _dafny.MultiSet([(self).f6])
        d_1816_v38_: _dafny.Set
        d_1816_v38_ = _dafny.Set({d_1811_v33_})
        d_1817_v39_: _dafny.Array
        nw290_ = _dafny.Array(None, 29)
        nw290_[int(0)] = len(d_1813_v35_)
        nw290_[int(1)] = p3
        nw290_[int(2)] = (self).f6
        nw290_[int(3)] = -178
        nw290_[int(4)] = (p3) - (p3)
        nw290_[int(5)] = (0) - (len(d_1814_v36_))
        nw290_[int(6)] = ((0) - ((self).f6)) - ((0) - ((self).f6))
        nw290_[int(7)] = default__.safeModuloInt((self).f6, 762)
        nw290_[int(8)] = (p3) - ((self).f6)
        nw290_[int(9)] = (self).f6
        nw290_[int(10)] = 747
        nw290_[int(11)] = -715
        nw290_[int(12)] = (self).f6
        nw290_[int(13)] = ((self).f6) + ((((d_1812_v34_).set(d_1811_v33_, default__.abs((self).f6))).set(d_1811_v33_, default__.abs((d_1815_v37_).cardinality))).cardinality)
        nw290_[int(14)] = (0) - ((0) - ((self).f6))
        nw290_[int(15)] = p3
        nw290_[int(16)] = p3
        nw290_[int(17)] = (self).f6
        nw290_[int(18)] = (d_1815_v37_).cardinality
        nw290_[int(19)] = (self).f6
        nw290_[int(20)] = (p3) + ((self).f6)
        nw290_[int(21)] = len(d_1816_v38_)
        nw290_[int(22)] = 97
        nw290_[int(23)] = (0) - ((p3) * (p3))
        nw290_[int(24)] = (self).f6
        nw290_[int(25)] = (self).f6
        nw290_[int(26)] = 296
        nw290_[int(27)] = (self).f6
        nw290_[int(28)] = (self).f6
        d_1817_v39_ = nw290_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1817_v39_).length(0)):
            d_1818_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_1818_i1_)) and ((d_1818_i1_) < ((d_1817_v39_).length(0)))):
                (d_1817_v39_)[(d_1818_i1_)] = default__.safeDivisionInt(d_1818_i1_, (self).f6)
        index299_ = default__.safeIndex(574, (d_1817_v39_).length(0))
        (d_1817_v39_)[index299_] = (self).f6
        index300_ = default__.safeIndex(574, (d_1817_v39_).length(0))
        (d_1817_v39_)[index300_] = (0) - (default__.safeDivisionInt((p3) + (p3), ((self).f6) + ((0) - ((self).f6))))
        d_1819_v40_: str
        d_1819_v40_ = _dafny.CodePoint('l')
        d_1820_v41_: D15
        d_1820_v41_ = D15_DC38(d_1819_v40_)
        d_1821_v42_: _dafny.Set
        d_1821_v42_ = _dafny.Set({(d_1820_v41_).cf71})
        d_1822_v43_: _dafny.Seq
        d_1822_v43_ = _dafny.SeqWithoutIsStrInference([(self).f6, len(d_1821_v42_)])
        d_1823_i2_: int
        d_1823_i2_ = 0
        with _dafny.label("6"):
            while (d_1822_v43_) <= (d_1822_v43_):
                with _dafny.c_label("6"):
                    if (d_1823_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_1823_i2_ = (d_1823_i2_) + (1)
                    d_1811_v33_ = (d_1811_v33_) and ((d_1811_v33_ if d_1811_v33_ else d_1811_v33_))
                    if d_1811_v33_:
                        d_1824_v44_: D4
                        d_1824_v44_ = D4_DC10(p3, (self).f6, p3)
                        d_1825_v45_: D15
                        d_1825_v45_ = D15_DC39((default__.fm25(d_1824_v44_, p1, globalState)) + (p1), default__.fm61(-243, (self).f6, globalState))
                        d_1826_v46_: _dafny.Map
                        d_1826_v46_ = _dafny.Map({d_1819_v40_: (self).f6})
                        d_1827_v47_: C0
                        nw291_ = C0()
                        nw291_.ctor__(default__.fm16(default__.fm46(d_1811_v33_, p3, d_1811_v33_, globalState), (d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], globalState))
                        d_1827_v47_ = nw291_
                        d_1828_v48_: _dafny.Set
                        d_1828_v48_ = _dafny.Set({d_1827_v47_})
                        d_1829_v49_: D4
                        d_1829_v49_ = D4_DC9(d_1815_v37_, d_1819_v40_, p3, d_1828_v48_, (d_1822_v43_)[default__.safeIndex((self).f6, len(d_1822_v43_))])
                        d_1830_v50_: _dafny.Seq
                        d_1830_v50_ = _dafny.SeqWithoutIsStrInference([(d_1826_v46_).set((d_1829_v49_).cf15, p3)])
                        d_1831_v51_: D7
                        d_1831_v51_ = D7_DC19(len(d_1827_v47_.f8))
                        d_1832_v52_: _dafny.Map
                        d_1832_v52_ = _dafny.Map({d_1831_v51_: p3})
                        d_1833_v53_: _dafny.Map
                        d_1833_v53_ = _dafny.Map({d_1832_v52_: d_1811_v33_})
                        d_1825_v45_ = default__.fm62(d_1830_v50_, ((d_1833_v53_)[d_1832_v52_] if (d_1832_v52_) in (d_1833_v53_) else d_1811_v33_), (self).f6, (d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], globalState)
                        d_1834_v54_: D13
                        d_1834_v54_ = D13_DC32(p1, True, (self).f6, d_1817_v39_)
                        d_1817_v39_ = (d_1834_v54_).cf60
                        d_1835_v55_: _dafny.Seq
                        d_1835_v55_ = _dafny.SeqWithoutIsStrInference([True])
                        d_1811_v33_ = (d_1835_v55_)[default__.safeIndex(len(d_1835_v55_), len(d_1835_v55_))]
                        r2 = d_1811_v33_
                        r3 = ((len(default__.fm50(d_1811_v33_, d_1811_v33_, len(p1), globalState))) * (len((d_1827_v47_.f8).set(default__.safeIndex((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], len(d_1827_v47_.f8)), d_1819_v40_)))) + ((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))])
                    elif True:
                        d_1836_v56_: _dafny.Array
                        nw292_ = _dafny.Array(_dafny.Array(None, 0), 6)
                        d_1836_v56_ = nw292_
                        d_1837_v57_: _dafny.Seq
                        d_1837_v57_ = _dafny.SeqWithoutIsStrInference([d_1811_v33_, d_1811_v33_])
                        d_1838_v58_: _dafny.Array
                        nw293_ = _dafny.Array(None, 18)
                        nw293_[int(0)] = d_1811_v33_
                        nw293_[int(1)] = d_1811_v33_
                        nw293_[int(2)] = d_1811_v33_
                        nw293_[int(3)] = True
                        nw293_[int(4)] = default__.fm17(globalState)
                        nw293_[int(5)] = not(not(d_1811_v33_))
                        nw293_[int(6)] = default__.fm17(globalState)
                        nw293_[int(7)] = True
                        nw293_[int(8)] = not(False)
                        nw293_[int(9)] = True
                        nw293_[int(10)] = d_1811_v33_
                        nw293_[int(11)] = d_1811_v33_
                        nw293_[int(12)] = d_1811_v33_
                        nw293_[int(13)] = False
                        nw293_[int(14)] = not((d_1837_v57_)[default__.safeIndex(p3, len(d_1837_v57_))])
                        nw293_[int(15)] = d_1811_v33_
                        nw293_[int(16)] = d_1811_v33_
                        nw293_[int(17)] = not(d_1811_v33_)
                        d_1838_v58_ = nw293_
                        index301_ = default__.safeIndex(692, (d_1836_v56_).length(0))
                        (d_1836_v56_)[index301_] = d_1838_v58_
                        index302_ = default__.safeIndex(692, (d_1836_v56_).length(0))
                        (d_1836_v56_)[index302_] = d_1838_v58_
                        d_1814_v36_ = (d_1814_v36_).set((0) - (p3), (d_1811_v33_) and (default__.fm30((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], (d_1837_v57_)[default__.safeIndex((0) - ((self).f6), len(d_1837_v57_))], not(d_1811_v33_), globalState)))
                        r0 = d_1811_v33_
                        d_1839_v59_: _dafny.Map
                        d_1839_v59_ = _dafny.Map({d_1819_v40_: len(d_1814_v36_)})
                        d_1839_v59_ = (d_1839_v59_).set((d_1819_v40_ if d_1811_v33_ else default__.fm29(not(d_1811_v33_), globalState)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyicitvk"))))
                        index303_ = default__.safeIndex(574, (d_1817_v39_).length(0))
                        (d_1817_v39_)[index303_] = (d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))]
                    d_1840_v60_: _dafny.Seq
                    d_1840_v60_ = _dafny.SeqWithoutIsStrInference([d_1811_v33_, d_1811_v33_, default__.fm17(globalState), d_1811_v33_, d_1811_v33_])
                    d_1841_v61_: D12
                    d_1841_v61_ = D12_DC28((d_1840_v60_)[default__.safeIndex(p3, len(d_1840_v60_))], p3)
                    d_1842_v62_: _dafny.Map
                    d_1842_v62_ = _dafny.Map({(self).f6: p3})
                    index304_ = default__.safeIndex(574, (d_1817_v39_).length(0))
                    (d_1817_v39_)[index304_] = len((_dafny.SeqWithoutIsStrInference([p3, p3])) + (_dafny.SeqWithoutIsStrInference([len(p1), len(d_1840_v60_), (d_1841_v61_).cf47, len(d_1842_v62_)])))
                    d_1814_v36_ = (d_1814_v36_).set(p3, d_1811_v33_)
                    pass
            pass
        if ((len(p1)) - ((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))])) < ((self).f6):
            d_1843_v63_: D12
            d_1843_v63_ = D12_DC29(not(True), (self).f6, d_1811_v33_, default__.safeDivisionInt((self).f6, (d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))]), not(d_1811_v33_))
            d_1843_v63_ = D12_DC29(d_1811_v33_, p3, d_1811_v33_, ((self).f6) * ((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))]), not(d_1811_v33_))
            d_1844_v64_: _dafny.Seq
            d_1844_v64_ = _dafny.SeqWithoutIsStrInference([(d_1816_v38_) != (d_1816_v38_)])
            d_1844_v64_ = d_1844_v64_
            d_1845_v65_: D4
            d_1845_v65_ = D4_DC8(_dafny.SeqWithoutIsStrInference([p2, p2]))
            d_1846_v66_: _dafny.Map
            d_1846_v66_ = _dafny.Map({d_1811_v33_: d_1845_v65_})
            d_1847_v67_: _dafny.Seq
            d_1847_v67_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_1846_v66_ = (d_1846_v66_).set(not(d_1811_v33_), D4_DC8(d_1847_v67_))
            d_1848_v68_: C3
            nw294_ = C3()
            nw294_.ctor__()
            d_1848_v68_ = nw294_
            d_1849_v69_: _dafny.Seq
            d_1849_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knsogq"))
            d_1849_v69_ = p1
        elif True:
            d_1850_v70_: _dafny.Array
            nw295_ = _dafny.Array(None, 12)
            nw295_[int(0)] = d_1811_v33_
            nw295_[int(1)] = default__.fm17(globalState)
            nw295_[int(2)] = default__.fm30((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], True, not(d_1811_v33_), globalState)
            nw295_[int(3)] = d_1811_v33_
            nw295_[int(4)] = True
            nw295_[int(5)] = False
            nw295_[int(6)] = ((self).f6) > ((0) - (((d_1815_v37_)[272] if (272) in (d_1815_v37_) else p3)))
            nw295_[int(7)] = d_1811_v33_
            nw295_[int(8)] = d_1811_v33_
            nw295_[int(9)] = d_1811_v33_
            nw295_[int(10)] = d_1811_v33_
            nw295_[int(11)] = d_1811_v33_
            d_1850_v70_ = nw295_
            index305_ = default__.safeIndex(61, (d_1850_v70_).length(0))
            (d_1850_v70_)[index305_] = d_1811_v33_
            index306_ = default__.safeIndex(61, (d_1850_v70_).length(0))
            (d_1850_v70_)[index306_] = d_1811_v33_
            d_1851_v71_: _dafny.Seq
            d_1851_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydabfju"))
            d_1852_v72_: _dafny.MultiSet
            d_1852_v72_ = _dafny.MultiSet([(p1).set(default__.safeIndex(len(d_1822_v43_), len(p1)), d_1819_v40_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayhpyhe"))])
            d_1853_v73_: _dafny.Set
            d_1853_v73_ = _dafny.Set({(self).f6, (self).f6})
            index307_ = default__.safeIndex(61, (d_1850_v70_).length(0))
            rhs208_ = ((d_1852_v72_ if d_1811_v33_ else d_1852_v72_)).issubset(d_1852_v72_)
            rhs209_ = (self).f6
            rhs210_ = ((self).f6 if d_1811_v33_ else default__.safeModuloInt((d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], (d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))]))
            rhs211_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "su"))) + (d_1851_v71_)
            rhs212_ = (d_1853_v73_).ispropersubset(d_1853_v73_)
            lhs120_ = d_1850_v70_
            lhs121_ = default__.safeIndex(61, (d_1850_v70_).length(0))
            r2 = rhs208_
            r3 = rhs209_
            r3 = rhs210_
            d_1851_v71_ = rhs211_
            lhs120_[lhs121_] = rhs212_
            d_1854_v74_: D12
            d_1854_v74_ = D12_DC28(d_1811_v33_, p3)
            index308_ = default__.safeIndex(574, (d_1817_v39_).length(0))
            (d_1817_v39_)[index308_] = (d_1854_v74_).cf47
            r3 = default__.safeModuloInt((0) - ((self).f6), (self).fm12((d_1850_v70_)[default__.safeIndex(61, (d_1850_v70_).length(0))], (d_1850_v70_)[default__.safeIndex(61, (d_1850_v70_).length(0))], globalState))
            d_1811_v33_ = (d_1850_v70_)[default__.safeIndex(61, (d_1850_v70_).length(0))]
        d_1855_v75_: _dafny.Seq
        d_1855_v75_ = _dafny.SeqWithoutIsStrInference([((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fm"))) + (p1)).set(default__.safeIndex((0) - (len(_dafny.Map({p3: d_1811_v33_}))), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fm"))) + (p1))), default__.fm51(d_1811_v33_, (d_1817_v39_)[default__.safeIndex(574, (d_1817_v39_).length(0))], p3, globalState)), p1, p1, p1, p1])
        d_1855_v75_ = d_1855_v75_
        r0 = (d_1819_v40_) in (_dafny.SeqWithoutIsStrInference([d_1819_v40_ for d_1856_i3_ in range(default__.abs(-746))]))
        r1 = not(default__.fm17(globalState))
        r2 = d_1811_v33_
        d_1857_v76_: _dafny.Array
        def lambda84_(d_1858_v33_, d_1859_v36_, d_1860_v39_):
            def lambda85_(d_1861_i4_):
                return ((d_1859_v36_)[len(_dafny.SeqWithoutIsStrInference([d_1858_v33_]))] if (len(_dafny.SeqWithoutIsStrInference([d_1858_v33_]))) in (d_1859_v36_) else ((d_1859_v36_)[(d_1860_v39_)[default__.safeIndex(574, (d_1860_v39_).length(0))]] if ((d_1860_v39_)[default__.safeIndex(574, (d_1860_v39_).length(0))]) in (d_1859_v36_) else d_1858_v33_))

            return lambda85_

        init44_ = lambda84_(d_1811_v33_, d_1814_v36_, d_1817_v39_)
        nw296_ = _dafny.Array(None, 24)
        for i0_44_ in range(nw296_.length(0)):
            nw296_[i0_44_] = init44_(i0_44_)
        d_1857_v76_ = nw296_
        d_1862_v77_: _dafny.Set
        d_1862_v77_ = _dafny.Set({d_1857_v76_})
        d_1863_v78_: _dafny.Set
        d_1863_v78_ = d_1862_v77_
        d_1864_v79_: _dafny.Map
        d_1864_v79_ = _dafny.Map({d_1811_v33_: d_1863_v78_})
        d_1865_v80_: bool
        d_1865_v80_ = d_1811_v33_
        r3 = default__.safeDivisionInt(len((d_1864_v79_) | ((d_1864_v79_).set((d_1865_v80_), d_1862_v77_))), 93)
        return r0, r1, r2, r3

    @property
    def f6(self):
        return self._f6

class C7(T1):
    def  __init__(self):
        self._f1: D0 = D0.default()()
        self._f0: _dafny.Seq = _dafny.Seq({})
        self._f4: _dafny.MultiSet = _dafny.MultiSet({})
        self._f5: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f1(self):
        return self._f1
    @f1.setter
    def f1(self, value):
        self._f1 = value
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f4, f5, f0, f1):
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f0 = f0
        (self).f1 = f1

    def fm6(self, p0, p1, globalState):
        return (-76) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "af")))) - (-575))

    def fm8(self, p0, globalState):
        if (not(False)) and (False):
            return not((len(_dafny.SeqWithoutIsStrInference([False]))) != (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1866_i0_ in range(default__.abs(440))]))))
        elif True:
            return (not(not(not(not(not(not(False))))))) == (False)

    def fm9(self, p0, p1, p2, p3, globalState):
        def iife153_():
            coll73_ = _dafny.Map()
            compr_73_: int
            for compr_73_ in _dafny.IntegerRange(834, 708):
                d_1867_v0_: int = compr_73_
                if ((834) <= (d_1867_v0_)) and ((d_1867_v0_) < (708)):
                    coll73_[(d_1867_v0_) * (597)] = 52
            return _dafny.Map(coll73_)
        def iife154_():
            coll74_ = _dafny.Map()
            compr_74_: int
            for compr_74_ in (_dafny.Map({329: True})).keys.Elements:
                d_1868_v1_: int = compr_74_
                if (d_1868_v1_) in (_dafny.Map({329: True})):
                    coll74_[default__.safeModuloInt(d_1868_v1_, 809)] = 743
            return _dafny.Map(coll74_)
        return ((_dafny.Map({606: -633})) | (iife153_()
        )) | ((iife154_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1869_i0_ in range(default__.abs(312))])): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dx"))))})))

    def m5(self, p0, p1, globalState):
        d_1870_v0_: bool
        d_1870_v0_ = True
        d_1871_i0_: int
        d_1871_i0_ = 0
        with _dafny.label("7"):
            while d_1870_v0_:
                with _dafny.c_label("7"):
                    if (d_1871_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_1871_i0_ = (d_1871_i0_) + (1)
                    d_1870_v0_ = d_1870_v0_
                    d_1872_v1_: _dafny.Array
                    nw297_ = _dafny.Array(_dafny.CodePoint('D'), 4)
                    d_1872_v1_ = nw297_
                    def lambda86_(d_1873_i1_):
                        return _dafny.CodePoint('c')

                    init45_ = lambda86_
                    nw298_ = _dafny.Array(None, 12)
                    for i0_45_ in range(nw298_.length(0)):
                        nw298_[i0_45_] = init45_(i0_45_)
                    d_1872_v1_ = nw298_
                    d_1874_v2_: _dafny.Array
                    nw299_ = _dafny.Array(_dafny.Seq({}), 8)
                    d_1874_v2_ = nw299_
                    index309_ = default__.safeIndex(910, (d_1874_v2_).length(0))
                    (d_1874_v2_)[index309_] = (self).f0
                    index310_ = default__.safeIndex(910, (d_1874_v2_).length(0))
                    (d_1874_v2_)[index310_] = ((self).f0) + ((self).f0)
                    d_1875_v3_: _dafny.Array
                    def lambda87_(d_1876_i2_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1877_i3_ in range(default__.abs(737))])

                    init46_ = lambda87_
                    nw300_ = _dafny.Array(None, 25)
                    for i0_46_ in range(nw300_.length(0)):
                        nw300_[i0_46_] = init46_(i0_46_)
                    d_1875_v3_ = nw300_
                    d_1875_v3_ = d_1875_v3_
                    pass
            pass
        d_1878_v4_: _dafny.Array
        nw301_ = _dafny.Array(_dafny.Set({}), 28)
        d_1878_v4_ = nw301_
        d_1879_v5_: D0
        d_1879_v5_ = D0_DC0(d_1878_v4_, 600)
        d_1880_v6_: _dafny.MultiSet
        d_1880_v6_ = _dafny.MultiSet([d_1879_v5_])
        d_1881_v7_: _dafny.Set
        d_1881_v7_ = _dafny.Set({(self).fm6(((d_1880_v6_).set(d_1879_v5_, default__.abs(p0))).cardinality, p0, globalState)})
        d_1882_v8_: _dafny.Seq
        d_1882_v8_ = _dafny.SeqWithoutIsStrInference([d_1881_v7_, d_1881_v7_, d_1881_v7_, d_1881_v7_, d_1881_v7_])
        if not((default__.fm10(True, globalState)).issubset((d_1882_v8_)[default__.safeIndex(p0, len(d_1882_v8_))])):
            d_1881_v7_ = _dafny.Set({(p0 if d_1870_v0_ else p0), p0})
            d_1883_v9_: _dafny.Array
            nw302_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1883_v9_ = nw302_
            d_1883_v9_ = d_1883_v9_
            d_1884_v10_: str
            d_1884_v10_ = _dafny.CodePoint('c')
            d_1885_v11_: _dafny.Seq
            d_1885_v11_ = _dafny.SeqWithoutIsStrInference([d_1884_v10_, d_1884_v10_, d_1884_v10_, d_1884_v10_])
            d_1886_v12_: _dafny.Set
            d_1886_v12_ = _dafny.Set({(d_1885_v11_) + (d_1885_v11_)})
            d_1887_v13_: D2
            d_1887_v13_ = D2_DC3(d_1886_v12_)
            d_1888_v14_: _dafny.Seq
            d_1888_v14_ = _dafny.SeqWithoutIsStrInference([d_1885_v11_])
            pat_let_tv47_ = d_1888_v14_
            pat_let_tv48_ = d_1888_v14_
            def iife155_(_pat_let40_0):
                def iife156_(d_1889_dt__update__tmp_h0_):
                    def iife158_():
                        coll75_ = _dafny.Set()
                        compr_75_: _dafny.Seq
                        for compr_75_ in (pat_let_tv47_).Elements:
                            d_1890_v15_: _dafny.Seq = compr_75_
                            if (d_1890_v15_) in (pat_let_tv48_):
                                coll75_ = coll75_.union(_dafny.Set([d_1890_v15_]))
                        return _dafny.Set(coll75_)
                    def iife157_(_pat_let41_0):
                        def iife159_(d_1891_dt__update_hcf4_h0_):
                            return D2_DC3(d_1891_dt__update_hcf4_h0_)
                        return iife159_(_pat_let41_0)
                    return iife157_(iife158_()
                    )
                return iife156_(_pat_let40_0)
            d_1886_v12_ = (iife155_(d_1887_v13_)).cf4
            d_1892_v16_: _dafny.Array
            nw303_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_1892_v16_ = nw303_
            d_1893_v17_: D2
            d_1893_v17_ = D2_DC4(p0, d_1870_v0_, d_1892_v16_)
            if (d_1893_v17_).cf6:
                d_1894_v18_: _dafny.Map
                d_1894_v18_ = _dafny.Map({d_1884_v10_: d_1885_v11_})
                d_1895_v19_: _dafny.MultiSet
                d_1895_v19_ = _dafny.MultiSet([((d_1894_v18_)[d_1884_v10_] if (d_1884_v10_) in (d_1894_v18_) else d_1885_v11_), d_1885_v11_, (d_1885_v11_) + (d_1885_v11_), d_1885_v11_])
                d_1895_v19_ = ((d_1895_v19_) | (d_1895_v19_) if d_1870_v0_ else d_1895_v19_)
                d_1896_v20_: int
                d_1896_v20_ = 221
                d_1896_v20_ = d_1896_v20_
                d_1897_v21_: _dafny.Array
                nw304_ = _dafny.Array(_dafny.Map({}), 19)
                d_1897_v21_ = nw304_
                d_1898_v22_: C1
                nw305_ = C1()
                nw305_.ctor__(default__.fm32(d_1896_v20_, globalState), d_1897_v21_, (self).f0, default__.fm49(d_1884_v10_, globalState))
                d_1898_v22_ = nw305_
                d_1899_v23_: C3
                nw306_ = C3()
                nw306_.ctor__()
                d_1899_v23_ = nw306_
                d_1900_v24_: _dafny.Map
                d_1900_v24_ = _dafny.Map({d_1896_v20_: d_1896_v20_})
                d_1901_v25_: _dafny.Map
                d_1901_v25_ = _dafny.Map({((d_1900_v24_)[d_1896_v20_] if (d_1896_v20_) in (d_1900_v24_) else (0) - (d_1896_v20_)): (0) - (d_1896_v20_)})
                d_1902_v26_: _dafny.Seq
                d_1902_v26_ = _dafny.SeqWithoutIsStrInference([d_1870_v0_, d_1870_v0_])
                d_1901_v25_ = (d_1901_v25_).set(p0, len(d_1902_v26_))
            elif True:
                d_1903_v27_: _dafny.Map
                d_1903_v27_ = _dafny.Map({d_1870_v0_: p0})
                d_1904_v28_: _dafny.Map
                d_1904_v28_ = _dafny.Map({p0: d_1903_v27_})
                d_1905_v29_: _dafny.MultiSet
                d_1905_v29_ = _dafny.MultiSet([d_1870_v0_])
                d_1870_v0_ = (default__.fm63(((d_1904_v28_)[p0] if (p0) in (d_1904_v28_) else _dafny.Map({d_1870_v0_: len(_dafny.Map({not(d_1870_v0_): p0}))})), p0, globalState)) != (d_1905_v29_)
                d_1906_v30_: int
                d_1906_v30_ = 838
                d_1906_v30_ = d_1906_v30_
                d_1906_v30_ = d_1906_v30_
                d_1907_v31_: D12
                d_1907_v31_ = D12_DC28(False, d_1906_v30_)
                d_1908_v32_: _dafny.Map
                d_1908_v32_ = _dafny.Map({(d_1907_v31_).cf47: (self).f0})
                d_1909_v33_: _dafny.Map
                d_1909_v33_ = _dafny.Map({d_1870_v0_: False})
                d_1885_v11_ = (((d_1885_v11_).set(default__.safeIndex(len(d_1908_v32_), len(d_1885_v11_)), default__.fm51(d_1870_v0_, len(d_1909_v33_), d_1906_v30_, globalState)) if d_1870_v0_ else (d_1885_v11_) + (_dafny.SeqWithoutIsStrInference([d_1884_v10_, d_1884_v10_, d_1884_v10_])))).set(default__.safeIndex(default__.safeModuloInt(d_1906_v30_, p0), len(((d_1885_v11_).set(default__.safeIndex(len(d_1908_v32_), len(d_1885_v11_)), default__.fm51(d_1870_v0_, len(d_1909_v33_), d_1906_v30_, globalState)) if d_1870_v0_ else (d_1885_v11_) + (_dafny.SeqWithoutIsStrInference([d_1884_v10_, d_1884_v10_, d_1884_v10_]))))), d_1884_v10_)
                d_1910_v34_: _dafny.Seq
                d_1910_v34_ = _dafny.SeqWithoutIsStrInference([False, d_1870_v0_, (p0) != (281)])
                d_1910_v34_ = (d_1910_v34_) + (d_1910_v34_)
            d_1911_v35_: _dafny.Array
            def lambda88_(d_1912_v10_, d_1913_p0_):
                def lambda89_(d_1914_i4_):
                    return _dafny.Map({d_1912_v10_: d_1913_p0_})

                return lambda89_

            init47_ = lambda88_(d_1884_v10_, p0)
            nw307_ = _dafny.Array(None, 1)
            for i0_47_ in range(nw307_.length(0)):
                nw307_[i0_47_] = init47_(i0_47_)
            d_1911_v35_ = nw307_
            d_1915_v36_: C1
            nw308_ = C1()
            nw308_.ctor__(p0, d_1911_v35_, (self).f0, (self).f5)
            d_1915_v36_ = nw308_
            d_1916_v37_: _dafny.Set
            d_1916_v37_ = _dafny.Set({d_1915_v36_, d_1915_v36_, d_1915_v36_})
            d_1917_v38_: _dafny.Map
            d_1917_v38_ = _dafny.Map({(0) - (p0): len(d_1916_v37_)})
            d_1918_v39_: _dafny.Map
            d_1918_v39_ = _dafny.Map({d_1870_v0_: len(d_1881_v7_)})
            d_1919_v40_: _dafny.Map
            d_1919_v40_ = _dafny.Map({d_1884_v10_: d_1915_v36_.f9})
            d_1920_v41_: _dafny.Map
            d_1920_v41_ = _dafny.Map({default__.fm19(d_1870_v0_, d_1870_v0_, d_1919_v40_, False, globalState): p0})
            d_1921_v42_: _dafny.Seq
            d_1921_v42_ = _dafny.SeqWithoutIsStrInference([d_1917_v38_, (d_1917_v38_).set(len(d_1918_v39_), len(d_1920_v41_)), d_1917_v38_, (d_1917_v38_) | (d_1917_v38_)])
            d_1922_v43_: _dafny.Map
            d_1922_v43_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0 for d_1923_i5_ in range(default__.abs(-919))]): d_1921_v42_})
            d_1924_v44_: _dafny.MultiSet
            d_1924_v44_ = _dafny.MultiSet([d_1870_v0_])
            d_1925_v45_: _dafny.Seq
            d_1925_v45_ = _dafny.SeqWithoutIsStrInference([(self).f5, D0_DC1(p0), self.f1, D0_DC1(d_1915_v36_.f9), (self).f5])
            d_1921_v42_ = (((((d_1922_v43_)[((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)] if (((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)) in (d_1922_v43_) else d_1921_v42_) if (p0) == (d_1915_v36_.f9) else (d_1921_v42_) + (d_1921_v42_))).set(default__.safeIndex((d_1915_v36_).fm6(p0, d_1915_v36_.f9, globalState), len((((d_1922_v43_)[((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)] if (((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)) in (d_1922_v43_) else d_1921_v42_) if (p0) == (d_1915_v36_.f9) else (d_1921_v42_) + (d_1921_v42_)))), d_1917_v38_)).set(default__.safeIndex(581, len(((((d_1922_v43_)[((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)] if (((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)) in (d_1922_v43_) else d_1921_v42_) if (p0) == (d_1915_v36_.f9) else (d_1921_v42_) + (d_1921_v42_))).set(default__.safeIndex((d_1915_v36_).fm6(p0, d_1915_v36_.f9, globalState), len((((d_1922_v43_)[((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)] if (((self).f0).set(default__.safeIndex((d_1924_v44_).cardinality, len((self).f0)), d_1915_v36_.f9)) in (d_1922_v43_) else d_1921_v42_) if (p0) == (d_1915_v36_.f9) else (d_1921_v42_) + (d_1921_v42_)))), d_1917_v38_))), (self).fm9(d_1925_v45_, p0, _dafny.Map({p0: _dafny.Set({True})}), _dafny.SeqWithoutIsStrInference([True]), globalState))
        elif True:
            d_1926_v46_: int
            d_1926_v46_ = 842
            d_1926_v46_ = ((0) - (p0)) - (d_1926_v46_)
            d_1927_v47_: D4
            d_1927_v47_ = D4_DC10((d_1926_v46_) + (p0), d_1926_v46_, d_1926_v46_)
            d_1927_v47_ = d_1927_v47_
            d_1928_v48_: C2
            nw309_ = C2()
            nw309_.ctor__()
            d_1928_v48_ = nw309_
            d_1929_v49_: D16
            d_1929_v49_ = D16_DC44((0) - (p0))
            d_1930_v50_: T0
            nw310_ = C2()
            nw310_.ctor__()
            d_1930_v50_ = nw310_
            rhs213_ = d_1929_v49_
            rhs214_ = (d_1930_v50_ if d_1870_v0_ else d_1930_v50_)
            d_1929_v49_ = rhs213_
            d_1930_v50_ = rhs214_
            d_1931_v51_: _dafny.Seq
            d_1931_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ru"))
            d_1932_v52_: D14
            d_1932_v52_ = D14_DC35(default__.safeDivisionInt(p0, d_1926_v46_), (d_1931_v51_ if True else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1933_i6_ in range(default__.abs(-598))])), (((self).f4)[d_1926_v46_] if (d_1926_v46_) in ((self).f4) else d_1926_v46_), True, (p0) >= (p0))
            source21_ = d_1932_v52_
            if source21_.is_DC35:
                d_1934___mcc_h0_ = source21_.cf63
                d_1935___mcc_h1_ = source21_.cf64
                d_1936___mcc_h2_ = source21_.cf65
                d_1937___mcc_h3_ = source21_.cf66
                d_1938___mcc_h4_ = source21_.cf67
                d_1939_cf67_ = d_1938___mcc_h4_
                d_1940_cf66_ = d_1937___mcc_h3_
                d_1941_cf65_ = d_1936___mcc_h2_
                d_1942_cf64_ = d_1935___mcc_h1_
                d_1943_cf63_ = d_1934___mcc_h0_
                d_1870_v0_ = not(not (d_1939_cf67_) or (not (d_1940_cf66_) or (True)))
                d_1944_v53_: C0
                nw311_ = C0()
                nw311_.ctor__(d_1942_cf64_)
                d_1944_v53_ = nw311_
                d_1943_cf63_ = ((self).f0)[default__.safeIndex(d_1943_cf63_, len((self).f0))]
                d_1945_v54_: str
                d_1945_v54_ = _dafny.CodePoint('m')
                d_1945_v54_ = d_1945_v54_
            elif source21_.is_DC36:
                d_1946___mcc_h5_ = source21_.cf68
                d_1947___mcc_h6_ = source21_.cf69
                d_1948_cf69_ = d_1947___mcc_h6_
                d_1949_cf68_ = d_1946___mcc_h5_
                d_1870_v0_ = d_1870_v0_
                d_1950_v55_: _dafny.Set
                d_1950_v55_ = _dafny.Set({d_1870_v0_})
                d_1870_v0_ = (d_1950_v55_).isdisjoint(_dafny.Set({d_1870_v0_}))
                d_1951_v56_: C4
                nw312_ = C4()
                nw312_.ctor__()
                d_1951_v56_ = nw312_
                d_1952_v57_: _dafny.Seq
                d_1952_v57_ = _dafny.SeqWithoutIsStrInference([True, d_1870_v0_])
                d_1953_v58_: _dafny.Map
                d_1953_v58_ = _dafny.Map({(d_1881_v7_ if (d_1952_v57_)[default__.safeIndex(((d_1949_cf68_)[p0] if (p0) in (d_1949_cf68_) else p0), len(d_1952_v57_))] else _dafny.Set({d_1926_v46_})): d_1870_v0_})
                d_1954_v59_: _dafny.Map
                d_1954_v59_ = _dafny.Map({d_1870_v0_: False})
                d_1953_v58_ = (d_1953_v58_).set(_dafny.Set({602, default__.fm20(not(d_1870_v0_), globalState), d_1926_v46_, len(d_1954_v59_), p0}), d_1870_v0_)
            elif True:
                d_1955___mcc_h7_ = source21_.cf62
                d_1956_cf62_ = d_1955___mcc_h7_
                index311_ = default__.safeIndex(885, (p1).length(0))
                (p1)[index311_] = d_1926_v46_
                d_1957_v60_: str
                d_1957_v60_ = _dafny.CodePoint('y')
                d_1958_v61_: _dafny.Array
                def lambda90_(d_1959_v0_):
                    def lambda91_(d_1960_i7_):
                        return d_1959_v0_

                    return lambda91_

                init48_ = lambda90_(d_1870_v0_)
                nw313_ = _dafny.Array(None, 15)
                for i0_48_ in range(nw313_.length(0)):
                    nw313_[i0_48_] = init48_(i0_48_)
                d_1958_v61_ = nw313_
                d_1961_v62_: D7
                d_1961_v62_ = D7_DC18()
                d_1962_v63_: _dafny.Map
                d_1962_v63_ = _dafny.Map({True: d_1870_v0_})
                index312_ = default__.safeIndex(885, (p1).length(0))
                rhs215_ = not(not(not((default__.fm48(d_1961_v62_, len(d_1962_v63_), p0, p0, globalState)).issubset((self).f4))))
                rhs216_ = d_1926_v46_
                rhs217_ = d_1870_v0_
                rhs218_ = d_1957_v60_
                rhs219_ = d_1958_v61_
                lhs122_ = p1
                lhs123_ = default__.safeIndex(885, (p1).length(0))
                d_1870_v0_ = rhs215_
                lhs122_[lhs123_] = rhs216_
                d_1870_v0_ = rhs217_
                d_1957_v60_ = rhs218_
                d_1958_v61_ = rhs219_
                d_1963_v64_: C6
                nw314_ = C6()
                nw314_.ctor__(d_1926_v46_)
                d_1963_v64_ = nw314_
                index313_ = default__.safeIndex(747, (d_1958_v61_).length(0))
                (d_1958_v61_)[index313_] = d_1870_v0_
                d_1964_v65_: _dafny.MultiSet
                d_1964_v65_ = _dafny.MultiSet([not(d_1870_v0_)])
                d_1965_v66_: _dafny.Seq
                d_1965_v66_ = _dafny.SeqWithoutIsStrInference([d_1931_v51_])
                index314_ = default__.safeIndex(747, (d_1958_v61_).length(0))
                index315_ = default__.safeIndex(885, (p1).length(0))
                rhs220_ = ((d_1964_v65_) | (d_1964_v65_)) == (d_1964_v65_)
                rhs221_ = (d_1963_v64_).f6
                rhs222_ = d_1958_v61_
                rhs223_ = (d_1965_v66_)[default__.safeIndex(p0, len(d_1965_v66_))]
                lhs124_ = d_1958_v61_
                lhs125_ = default__.safeIndex(747, (d_1958_v61_).length(0))
                lhs126_ = p1
                lhs127_ = default__.safeIndex(885, (p1).length(0))
                lhs124_[lhs125_] = rhs220_
                lhs126_[lhs127_] = rhs221_
                d_1958_v61_ = rhs222_
                d_1931_v51_ = rhs223_
                index316_ = default__.safeIndex(885, (p1).length(0))
                (p1)[index316_] = -723
        hi14_ = p0
        for d_1966_i8_ in range((p0) + (p0), hi14_):
            d_1967_v67_: _dafny.Map
            d_1967_v67_ = _dafny.Map({d_1966_i8_: d_1870_v0_})
            d_1968_v68_: _dafny.Map
            d_1968_v68_ = _dafny.Map({p0: d_1967_v67_})
            d_1969_v69_: int
            d_1969_v69_ = 611
            rhs224_ = d_1968_v68_
            rhs225_ = (p0) - ((d_1966_i8_) + ((0) - (p0)))
            d_1968_v68_ = rhs224_
            d_1969_v69_ = rhs225_
            if d_1870_v0_:
                d_1970_v70_: _dafny.Seq
                d_1970_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ashcuj"))
                d_1971_v71_: C0
                nw315_ = C0()
                nw315_.ctor__(d_1970_v70_)
                d_1971_v71_ = nw315_
                d_1969_v69_ = d_1966_i8_
                d_1972_v72_: _dafny.Array
                nw316_ = _dafny.Array(False, 8)
                d_1972_v72_ = nw316_
                index317_ = default__.safeIndex(891, (d_1972_v72_).length(0))
                (d_1972_v72_)[index317_] = d_1870_v0_
                index318_ = default__.safeIndex(891, (d_1972_v72_).length(0))
                (d_1972_v72_)[index318_] = d_1870_v0_
                d_1969_v69_ = default__.safeModuloInt(p0, d_1969_v69_)
                d_1973_v73_: _dafny.Map
                d_1973_v73_ = _dafny.Map({False: p0})
                d_1974_v74_: _dafny.Seq
                d_1974_v74_ = _dafny.SeqWithoutIsStrInference([len((d_1973_v73_) | (_dafny.Map({(d_1972_v72_)[default__.safeIndex(891, (d_1972_v72_).length(0))]: d_1966_i8_}))), d_1969_v69_])
                d_1974_v74_ = (_dafny.SeqWithoutIsStrInference([d_1966_i8_ for d_1975_i9_ in range(default__.abs(11))])).set(default__.safeIndex(d_1969_v69_, len(_dafny.SeqWithoutIsStrInference([d_1966_i8_ for d_1976_i9_ in range(default__.abs(11))]))), d_1969_v69_)
            elif True:
                d_1977_v75_: str
                d_1977_v75_ = _dafny.CodePoint('m')
                d_1978_v76_: _dafny.Seq
                d_1978_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unqvelq"))
                d_1979_v77_: C0
                nw317_ = C0()
                nw317_.ctor__(d_1978_v76_)
                d_1979_v77_ = nw317_
                d_1980_v78_: _dafny.Set
                d_1980_v78_ = _dafny.Set({d_1979_v77_, d_1979_v77_, d_1979_v77_})
                d_1981_v79_: _dafny.Array
                nw318_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1981_v79_ = nw318_
                d_1982_v80_: D2
                d_1982_v80_ = D2_DC4(len(d_1979_v77_.f8), d_1870_v0_, d_1981_v79_)
                d_1983_v81_: D2
                d_1983_v81_ = D2_DC6(d_1982_v80_)
                d_1984_v82_: _dafny.Map
                d_1984_v82_ = _dafny.Map({len(d_1978_v76_): d_1983_v81_})
                d_1985_v83_: D4
                d_1985_v83_ = D4_DC9((self).f4, d_1977_v75_, d_1969_v69_, d_1980_v78_, len(d_1984_v82_))
                d_1986_v84_: _dafny.Map
                d_1986_v84_ = _dafny.Map({d_1870_v0_: d_1979_v77_.f8})
                d_1987_v85_: D12
                d_1987_v85_ = D12_DC29(d_1870_v0_, p0, d_1870_v0_, len(((d_1986_v84_)[d_1870_v0_] if (d_1870_v0_) in (d_1986_v84_) else d_1978_v76_)), d_1870_v0_)
                pat_let_tv49_ = d_1979_v77_
                d_1988_v86_: _dafny.MultiSet
                def iife160_(_pat_let42_0):
                    def iife161_(d_1989_dt__update__tmp_h1_):
                        def iife162_(_pat_let43_0):
                            def iife163_(d_1990_dt__update_hcf49_h0_):
                                return D12_DC29((d_1989_dt__update__tmp_h1_).cf48, d_1990_dt__update_hcf49_h0_, (d_1989_dt__update__tmp_h1_).cf50, (d_1989_dt__update__tmp_h1_).cf51, (d_1989_dt__update__tmp_h1_).cf52)
                            return iife163_(_pat_let43_0)
                        return iife162_(len(pat_let_tv49_.f8))
                    return iife161_(_pat_let42_0)
                d_1988_v86_ = _dafny.MultiSet([d_1987_v85_, d_1987_v85_, iife160_(d_1987_v85_)])
                rhs226_ = d_1870_v0_
                rhs227_ = p0
                rhs228_ = D4_DC9((self).f4, d_1977_v75_, ((d_1988_v86_)[d_1987_v85_] if (d_1987_v85_) in (d_1988_v86_) else p0), d_1980_v78_, default__.safeDivisionInt(d_1969_v69_, (0) - (d_1966_i8_)))
                rhs229_ = d_1870_v0_
                d_1870_v0_ = rhs226_
                d_1969_v69_ = rhs227_
                d_1985_v83_ = rhs228_
                d_1870_v0_ = rhs229_
                d_1969_v69_ = d_1969_v69_
                d_1991_v90_: D17
                d_1991_v90_ = D17_DC45(_dafny.MultiSet([(self).f4]))
                def iife164_():
                    coll76_ = _dafny.Map()
                    compr_76_: _dafny.MultiSet
                    for compr_76_ in (_dafny.SeqWithoutIsStrInference([(self).f4 for d_1992_i10_ in range(default__.abs(-905))])).Elements:
                        d_1993_v87_: _dafny.MultiSet = compr_76_
                        if (d_1993_v87_) in (_dafny.SeqWithoutIsStrInference([(self).f4 for d_1992_i10_ in range(default__.abs(-905))])):
                            coll76_[d_1993_v87_] = d_1870_v0_
                    return _dafny.Map(coll76_)
                def iife165_():
                    def iife167_():
                        coll79_ = _dafny.Map()
                        compr_79_: _dafny.MultiSet
                        for compr_79_ in ((d_1991_v90_).cf83).Elements:
                            d_1994_v89_: _dafny.MultiSet = compr_79_
                            if (d_1994_v89_) in ((d_1991_v90_).cf83):
                                coll79_[d_1994_v89_] = _dafny.Set({d_1870_v0_, d_1870_v0_})
                        return _dafny.Map(coll79_)
                    coll77_ = _dafny.Map()
                    def iife166_():
                        coll78_ = _dafny.Map()
                        compr_78_: _dafny.MultiSet
                        for compr_78_ in ((d_1991_v90_).cf83).Elements:
                            d_1994_v89_: _dafny.MultiSet = compr_78_
                            if (d_1994_v89_) in ((d_1991_v90_).cf83):
                                coll78_[d_1994_v89_] = _dafny.Set({d_1870_v0_, d_1870_v0_})
                        return _dafny.Map(coll78_)
                    compr_77_: _dafny.MultiSet
                    for compr_77_ in (iife166_()
                    ).keys.Elements:
                        d_1995_v88_: _dafny.MultiSet = compr_77_
                        if (d_1995_v88_) in (iife167_()
                        ):
                            coll77_[d_1995_v88_] = d_1870_v0_
                    return _dafny.Map(coll77_)
                rhs230_ = d_1978_v76_
                rhs231_ = default__.safeDivisionInt(len((iife164_()
                ) | (iife165_()
                )), p0)
                rhs232_ = p0
                d_1978_v76_ = rhs230_
                d_1969_v69_ = rhs231_
                d_1969_v69_ = rhs232_
                d_1996_v91_: _dafny.Seq
                d_1996_v91_ = _dafny.SeqWithoutIsStrInference([((d_1967_v67_)[len((self).f0)] if (len((self).f0)) in (d_1967_v67_) else d_1870_v0_), d_1870_v0_, d_1870_v0_])
                d_1870_v0_ = (((self).fm6(d_1969_v69_, (default__.fm24(d_1870_v0_, globalState)).cardinality, globalState) if d_1870_v0_ else len(d_1996_v91_))) != (p0)
                d_1997_v92_: _dafny.Array
                def lambda92_(d_1998_v77_, d_1999_v76_):
                    def lambda93_(d_2000_i11_):
                        return ((d_1998_v77_.f8)[default__.safeIndex(len(d_1999_v76_), len(d_1998_v77_.f8))]) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvmhkbqyi")))

                    return lambda93_

                init49_ = lambda92_(d_1979_v77_, d_1978_v76_)
                nw319_ = _dafny.Array(None, 5)
                for i0_49_ in range(nw319_.length(0)):
                    nw319_[i0_49_] = init49_(i0_49_)
                d_1997_v92_ = nw319_
                index319_ = default__.safeIndex(953, (d_1997_v92_).length(0))
                (d_1997_v92_)[index319_] = d_1870_v0_
                d_2001_v93_: _dafny.Array
                nw320_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_2001_v93_ = nw320_
                index320_ = default__.safeIndex(953, (d_1997_v92_).length(0))
                rhs233_ = False
                rhs234_ = d_1870_v0_
                rhs235_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyco"))
                rhs236_ = default__.fm30(default__.safeModuloInt(p0, d_1966_i8_), d_1870_v0_, (True) and (d_1870_v0_), globalState)
                rhs237_ = d_2001_v93_
                lhs128_ = d_1997_v92_
                lhs129_ = default__.safeIndex(953, (d_1997_v92_).length(0))
                lhs130_ = d_1979_v77_
                lhs128_[lhs129_] = rhs233_
                d_1870_v0_ = rhs234_
                lhs130_.f8 = rhs235_
                d_1870_v0_ = rhs236_
                d_2001_v93_ = rhs237_
            d_2002_v94_: D4
            d_2002_v94_ = D4_DC8(_dafny.SeqWithoutIsStrInference([d_1879_v5_]))
            d_2002_v94_ = d_2002_v94_
            d_1870_v0_ = d_1870_v0_
        d_2003_v95_: _dafny.Map
        d_2003_v95_ = _dafny.Map({d_1870_v0_: p0})
        d_2004_v96_: _dafny.Map
        d_2004_v96_ = _dafny.Map({d_2003_v95_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fakxfloyk"))})
        d_2005_v97_: _dafny.Array
        nw321_ = _dafny.Array(None, 3)
        nw321_[int(0)] = d_2004_v96_
        nw321_[int(1)] = (d_2004_v96_) | (d_2004_v96_)
        nw321_[int(2)] = (d_2004_v96_) | (d_2004_v96_)
        d_2005_v97_ = nw321_
        d_2006_v98_: D18
        d_2006_v98_ = D18_DC47(d_2004_v96_)
        index321_ = default__.safeIndex(894, (d_2005_v97_).length(0))
        (d_2005_v97_)[index321_] = (d_2006_v98_).cf85
        index322_ = default__.safeIndex(894, (d_2005_v97_).length(0))
        (d_2005_v97_)[index322_] = d_2004_v96_
        d_2007_v99_: _dafny.Map
        d_2007_v99_ = _dafny.Map({False: d_1870_v0_})
        d_2008_v100_: _dafny.Map
        d_2008_v100_ = _dafny.Map({982: p0})
        d_2009_v101_: _dafny.Set
        d_2009_v101_ = _dafny.Set({_dafny.Map({p0: len(d_2007_v99_)}), d_2008_v100_, d_2008_v100_})
        d_2010_v102_: _dafny.Map
        d_2010_v102_ = _dafny.Map({816: d_1870_v0_})
        d_2011_v103_: _dafny.Map
        d_2011_v103_ = _dafny.Map({len(_dafny.Set({((d_2010_v102_)[p0] if (p0) in (d_2010_v102_) else d_1870_v0_)})): d_1870_v0_})
        rhs238_ = (d_2009_v101_).ispropersubset(d_2009_v101_)
        rhs239_ = ((p0) - ((0) - (p0))) in (d_2011_v103_)
        rhs240_ = (d_1870_v0_ if (d_1870_v0_) and (d_1870_v0_) else True)
        d_1870_v0_ = rhs238_
        d_1870_v0_ = rhs239_
        d_1870_v0_ = rhs240_
        d_2012_i12_: int
        d_2012_i12_ = 0
        with _dafny.label("8"):
            while d_1870_v0_:
                with _dafny.c_label("8"):
                    if (d_2012_i12_) >= (100):
                        raise _dafny.Break("8")
                    d_2012_i12_ = (d_2012_i12_) + (1)
                    d_2013_v104_: _dafny.Seq
                    d_2013_v104_ = _dafny.SeqWithoutIsStrInference([d_1870_v0_])
                    d_2014_v105_: D7
                    d_2014_v105_ = D7_DC17(d_2013_v104_)
                    d_2015_v106_: D4
                    d_2015_v106_ = D4_DC11(p0, ((d_2014_v105_).cf34) != (d_2013_v104_), True, not(d_1870_v0_), d_1870_v0_)
                    d_2015_v106_ = d_2015_v106_
                    d_2016_v107_: int
                    d_2016_v107_ = 399
                    d_2016_v107_ = d_2016_v107_
                    d_2017_v108_: D2
                    d_2017_v108_ = D2_DC5(d_2016_v107_, (self).f0, d_1870_v0_)
                    d_2018_v109_: T1
                    nw322_ = C5()
                    nw322_.ctor__(p1, (d_2017_v108_).cf9, self.f1)
                    d_2018_v109_ = nw322_
                    d_2019_v110_: D12
                    d_2019_v110_ = D12_DC29(d_1870_v0_, d_2016_v107_, d_1870_v0_, 331, False)
                    d_2020_v111_: _dafny.Seq
                    d_2020_v111_ = _dafny.SeqWithoutIsStrInference([d_2018_v109_, d_2018_v109_, d_2018_v109_, (d_2018_v109_ if (d_2019_v110_).cf52 else d_2018_v109_)])
                    d_2018_v109_ = (d_2020_v111_)[default__.safeIndex(((self).fm6(d_2016_v107_, d_2016_v107_, globalState)) * (d_2016_v107_), len(d_2020_v111_))]
                    d_2021_v112_: C2
                    nw323_ = C2()
                    nw323_.ctor__()
                    d_2021_v112_ = nw323_
                    pass
            pass

    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5

class C8(T0, T1):
    def  __init__(self):
        self._f1: D0 = D0.default()()
        self._f0: _dafny.Seq = _dafny.Seq({})
        self._f2: int = int(0)
        self._f3: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f1(self):
        return self._f1
    @f1.setter
    def f1(self, value):
        self._f1 = value
    @property
    def f0(self):
        return self._f0
    def ctor__(self, f2, f3, f0, f1):
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f0 = f0
        (self).f1 = f1

    def fm5(self, p0, p1, p2, p3, globalState):
        return not((((self).f0)[default__.safeIndex((self).f2, len((self).f0))]) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f2 for d_2022_i0_ in range(default__.abs(191))]))))

    def fm6(self, p0, p1, globalState):
        if (len((self).f0)) >= ((self).f2):
            return (self).f2
        elif True:
            return len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: (self).f2}), _dafny.Map({not(False): (self).f2}), _dafny.Map({True: (self).f2}), _dafny.Map({False: (self).f2})]))

    def fm7(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyousyj"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))

    def m1(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        d_2023_v0_: str
        d_2023_v0_ = _dafny.CodePoint('d')
        d_2024_v1_: _dafny.Map
        d_2024_v1_ = _dafny.Map({d_2023_v0_: p0})
        d_2025_v2_: bool
        d_2025_v2_ = True
        d_2026_v3_: _dafny.Seq
        d_2026_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smiohn"))
        d_2027_v5_: _dafny.Map
        d_2027_v5_ = _dafny.Map({d_2023_v0_: d_2025_v2_})
        d_2028_v6_: _dafny.Array
        nw324_ = _dafny.Array(None, 25)
        nw324_[int(0)] = (d_2024_v1_).set(d_2023_v0_, (self).f2)
        nw324_[int(1)] = d_2024_v1_
        nw324_[int(2)] = d_2024_v1_
        nw324_[int(3)] = _dafny.Map({d_2023_v0_: 409})
        nw324_[int(4)] = default__.fm52(d_2025_v2_, D15_DC38(default__.fm51(True, len(d_2026_v3_), 535, globalState)), d_2025_v2_, globalState)
        nw324_[int(5)] = d_2024_v1_
        nw324_[int(6)] = d_2024_v1_
        nw324_[int(7)] = _dafny.Map({d_2023_v0_: 237})
        nw324_[int(8)] = d_2024_v1_
        nw324_[int(9)] = d_2024_v1_
        nw324_[int(10)] = _dafny.Map({d_2023_v0_: (self).f2})
        nw324_[int(11)] = d_2024_v1_
        nw324_[int(12)] = (_dafny.Map({_dafny.CodePoint('g'): (self).f2})).set(d_2023_v0_, (0) - (p0))
        nw324_[int(13)] = d_2024_v1_
        nw324_[int(14)] = d_2024_v1_
        nw324_[int(15)] = d_2024_v1_
        nw324_[int(16)] = d_2024_v1_
        nw324_[int(17)] = _dafny.Map({d_2023_v0_: p0})
        nw324_[int(18)] = d_2024_v1_
        def iife168_():
            coll80_ = _dafny.Map()
            compr_80_: str
            for compr_80_ in (d_2027_v5_).keys.Elements:
                d_2029_v4_: str = compr_80_
                if (d_2029_v4_) in (d_2027_v5_):
                    coll80_[d_2029_v4_] = p0
            return _dafny.Map(coll80_)
        nw324_[int(19)] = iife168_()
        
        nw324_[int(20)] = d_2024_v1_
        nw324_[int(21)] = _dafny.Map({default__.fm29(d_2025_v2_, globalState): len((_dafny.SeqWithoutIsStrInference([d_2023_v0_ for d_2030_i0_ in range(default__.abs(565))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_2023_v0_ for d_2031_i0_ in range(default__.abs(565))]))), d_2023_v0_))})
        nw324_[int(22)] = (d_2024_v1_).set(d_2023_v0_, (self).f2)
        nw324_[int(23)] = d_2024_v1_
        nw324_[int(24)] = d_2024_v1_
        d_2028_v6_ = nw324_
        d_2032_v7_: C1
        nw325_ = C1()
        nw325_.ctor__((self).f2, d_2028_v6_, ((self).f0 if d_2025_v2_ else (self).f0), self.f1)
        d_2032_v7_ = nw325_
        d_2033_v8_: _dafny.Array
        def lambda94_(d_2034_p0_):
            def lambda95_(d_2035_i1_):
                return default__.safeDivisionInt(d_2035_i1_, d_2034_p0_)

            return lambda95_

        init50_ = lambda94_(p0)
        nw326_ = _dafny.Array(None, 7)
        for i0_50_ in range(nw326_.length(0)):
            nw326_[i0_50_] = init50_(i0_50_)
        d_2033_v8_ = nw326_
        d_2036_v9_: _dafny.Map
        d_2036_v9_ = _dafny.Map({d_2026_v3_: d_2033_v8_})
        d_2037_v10_: _dafny.Map
        d_2037_v10_ = _dafny.Map({((d_2036_v9_)[d_2026_v3_] if (d_2026_v3_) in (d_2036_v9_) else d_2033_v8_): (p0) + (p0)})
        d_2037_v10_ = (d_2037_v10_).set(d_2033_v8_, len(d_2027_v5_))
        d_2038_v11_: _dafny.Seq
        d_2038_v11_ = _dafny.SeqWithoutIsStrInference([d_2025_v2_, False])
        d_2039_v12_: _dafny.Seq
        d_2039_v12_ = _dafny.SeqWithoutIsStrInference([d_2038_v11_])
        d_2038_v11_ = ((d_2038_v11_) + ((d_2039_v12_)[default__.safeIndex(p0, len(d_2039_v12_))])) + (d_2038_v11_)
        d_2040_v13_: _dafny.Map
        d_2040_v13_ = _dafny.Map({p0: d_2025_v2_})
        d_2041_v14_: _dafny.Map
        d_2041_v14_ = _dafny.Map({len(d_2040_v13_): d_2032_v7_.f9})
        d_2042_v15_: _dafny.Set
        d_2042_v15_ = _dafny.Set({(self).f2})
        d_2043_v16_: _dafny.Map
        d_2043_v16_ = _dafny.Map({d_2025_v2_: _dafny.MultiSet([(self).f2, len(d_2041_v14_), (0) - (len(d_2042_v15_))])})
        d_2044_v17_: _dafny.MultiSet
        d_2044_v17_ = _dafny.MultiSet([default__.fm44(d_2025_v2_, globalState)])
        hi15_ = (default__.fm44(d_2025_v2_, globalState)) + (p0)
        for d_2045_i2_ in range((((d_2043_v16_)[d_2025_v2_] if (d_2025_v2_) in (d_2043_v16_) else d_2044_v17_)).cardinality, hi15_):
            d_2025_v2_ = False
            if d_2025_v2_:
                d_2046_v18_: _dafny.Array
                nw327_ = _dafny.Array(None, 3)
                nw327_[int(0)] = d_2026_v3_
                nw327_[int(1)] = d_2026_v3_
                nw327_[int(2)] = d_2026_v3_
                d_2046_v18_ = nw327_
                index323_ = default__.safeIndex(389, (d_2046_v18_).length(0))
                (d_2046_v18_)[index323_] = d_2026_v3_
                index324_ = default__.safeIndex(389, (d_2046_v18_).length(0))
                (d_2046_v18_)[index324_] = (d_2026_v3_) + (d_2026_v3_)
                d_2047_v19_: _dafny.Map
                d_2047_v19_ = _dafny.Map({d_2025_v2_: not(d_2025_v2_)})
                (d_2032_v7_).f9 = (0) - ((_dafny.MultiSet([_dafny.Map({False: True}), _dafny.Map({d_2025_v2_: True}), ((_dafny.Map({not(d_2025_v2_): d_2025_v2_})).set(True, d_2025_v2_)) | (d_2047_v19_), d_2047_v19_, d_2047_v19_])).cardinality)
                d_2048_v20_: _dafny.Array
                nw328_ = _dafny.Array(None, 6)
                nw328_[int(0)] = d_2040_v13_
                nw328_[int(1)] = (d_2040_v13_).set((d_2032_v7_).fm28(d_2032_v7_.f9, d_2025_v2_, globalState), d_2025_v2_)
                nw328_[int(2)] = d_2040_v13_
                nw328_[int(3)] = d_2040_v13_
                nw328_[int(4)] = d_2040_v13_
                nw328_[int(5)] = default__.fm50(False, d_2025_v2_, d_2045_i2_, globalState)
                d_2048_v20_ = nw328_
                d_2049_v21_: _dafny.Array
                nw329_ = _dafny.Array(None, 8)
                nw329_[int(0)] = d_2048_v20_
                nw329_[int(1)] = d_2048_v20_
                nw329_[int(2)] = d_2048_v20_
                nw329_[int(3)] = d_2048_v20_
                nw329_[int(4)] = d_2048_v20_
                nw329_[int(5)] = d_2048_v20_
                nw329_[int(6)] = d_2048_v20_
                nw329_[int(7)] = d_2048_v20_
                d_2049_v21_ = nw329_
                index325_ = default__.safeIndex(775, (d_2049_v21_).length(0))
                (d_2049_v21_)[index325_] = d_2048_v20_
                index326_ = default__.safeIndex(775, (d_2049_v21_).length(0))
                (d_2049_v21_)[index326_] = d_2048_v20_
                r0 = d_2025_v2_
                d_2050_v22_: _dafny.Set
                d_2050_v22_ = _dafny.Set({d_2025_v2_})
                d_2051_v23_: _dafny.Map
                d_2051_v23_ = _dafny.Map({d_2025_v2_: len(d_2050_v22_)})
                r0 = (d_2025_v2_) not in (d_2051_v23_)
            elif True:
                d_2052_v24_: _dafny.Array
                nw330_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_2052_v24_ = nw330_
                d_2053_v25_: _dafny.Array
                def lambda96_(d_2054_v2_):
                    def lambda97_(d_2055_i3_):
                        return d_2054_v2_

                    return lambda97_

                init51_ = lambda96_(d_2025_v2_)
                nw331_ = _dafny.Array(None, 23)
                for i0_51_ in range(nw331_.length(0)):
                    nw331_[i0_51_] = init51_(i0_51_)
                d_2053_v25_ = nw331_
                index327_ = default__.safeIndex(748, (d_2052_v24_).length(0))
                (d_2052_v24_)[index327_] = d_2053_v25_
                d_2056_v26_: _dafny.Array
                nw332_ = _dafny.Array(D14.default()(), 2)
                d_2056_v26_ = nw332_
                d_2057_v27_: D14
                d_2057_v27_ = D14_DC35(len(_dafny.SeqWithoutIsStrInference([(0) - (d_2032_v7_.f9) for d_2058_i4_ in range(default__.abs(630))])), d_2026_v3_, d_2045_i2_, True, d_2025_v2_)
                pat_let_tv50_ = d_2023_v0_
                pat_let_tv51_ = d_2040_v13_
                pat_let_tv52_ = d_2032_v7_
                pat_let_tv53_ = d_2032_v7_
                pat_let_tv54_ = d_2040_v13_
                pat_let_tv55_ = d_2025_v2_
                index328_ = default__.safeIndex(606, (d_2056_v26_).length(0))
                def iife169_(_pat_let44_0):
                    def iife170_(d_2059_dt__update__tmp_h0_):
                        def iife171_(_pat_let45_0):
                            def iife172_(d_2061_dt__update_hcf64_h0_):
                                def iife173_(_pat_let46_0):
                                    def iife174_(d_2062_dt__update_hcf67_h0_):
                                        return D14_DC35((d_2059_dt__update__tmp_h0_).cf63, d_2061_dt__update_hcf64_h0_, (d_2059_dt__update__tmp_h0_).cf65, (d_2059_dt__update__tmp_h0_).cf66, d_2062_dt__update_hcf67_h0_)
                                    return iife174_(_pat_let46_0)
                                return iife173_(((pat_let_tv51_)[pat_let_tv52_.f9] if (pat_let_tv53_.f9) in (pat_let_tv54_) else pat_let_tv55_))
                            return iife172_(_pat_let45_0)
                        return iife171_(_dafny.SeqWithoutIsStrInference([pat_let_tv50_ for d_2060_i5_ in range(default__.abs(892))]))
                    return iife170_(_pat_let44_0)
                (d_2056_v26_)[index328_] = iife169_(d_2057_v27_)
                index329_ = default__.safeIndex(581, (d_2053_v25_).length(0))
                (d_2053_v25_)[index329_] = d_2025_v2_
                index330_ = default__.safeIndex(748, (d_2052_v24_).length(0))
                index331_ = default__.safeIndex(606, (d_2056_v26_).length(0))
                index332_ = default__.safeIndex(581, (d_2053_v25_).length(0))
                rhs241_ = d_2053_v25_
                rhs242_ = d_2057_v27_
                rhs243_ = d_2025_v2_
                rhs244_ = (self).f2
                rhs245_ = default__.fm17(globalState)
                lhs131_ = d_2052_v24_
                lhs132_ = default__.safeIndex(748, (d_2052_v24_).length(0))
                lhs133_ = d_2056_v26_
                lhs134_ = default__.safeIndex(606, (d_2056_v26_).length(0))
                lhs135_ = d_2053_v25_
                lhs136_ = default__.safeIndex(581, (d_2053_v25_).length(0))
                lhs137_ = d_2032_v7_
                lhs131_[lhs132_] = rhs241_
                lhs133_[lhs134_] = rhs242_
                lhs135_[lhs136_] = rhs243_
                lhs137_.f9 = rhs244_
                r1 = rhs245_
                (d_2032_v7_).f9 = (((default__.fm13(d_2023_v0_, True, globalState)).intersection(d_2044_v17_)).cardinality if d_2025_v2_ else d_2032_v7_.f9)
                d_2063_v28_: int
                d_2064_v29_: _dafny.Map
                d_2065_v30_: int
                out49_: int
                out50_: _dafny.Map
                out51_: int
                out49_, out50_, out51_ = (self).m3(d_2045_i2_, (self).f3, default__.fm30(p0, (d_2053_v25_)[default__.safeIndex(581, (d_2053_v25_).length(0))], (D18_DC50((d_2053_v25_)[default__.safeIndex(581, (d_2053_v25_).length(0))], d_2053_v25_, d_2025_v2_, (d_2053_v25_)[default__.safeIndex(581, (d_2053_v25_).length(0))], d_2026_v3_)).cf90, globalState), globalState)
                d_2063_v28_ = out49_
                d_2064_v29_ = out50_
                d_2065_v30_ = out51_
                (d_2032_v7_).f9 = default__.safeModuloInt((811) - (((self).f0)[default__.safeIndex(d_2045_i2_, len((self).f0))]), ((d_2044_v17_)[d_2063_v28_] if (d_2063_v28_) in (d_2044_v17_) else d_2032_v7_.f9))
                (d_2032_v7_).f9 = (self).fm6(d_2045_i2_, (0) - ((p0 if d_2025_v2_ else 877)), globalState)
            d_2066_v31_: _dafny.Array
            def lambda98_(d_2067_v0_):
                def lambda99_(d_2068_i6_):
                    return d_2067_v0_

                return lambda99_

            init52_ = lambda98_(d_2023_v0_)
            nw333_ = _dafny.Array(None, 26)
            for i0_52_ in range(nw333_.length(0)):
                nw333_[i0_52_] = init52_(i0_52_)
            d_2066_v31_ = nw333_
            index333_ = default__.safeIndex(149, (d_2066_v31_).length(0))
            (d_2066_v31_)[index333_] = d_2023_v0_
            index334_ = default__.safeIndex(149, (d_2066_v31_).length(0))
            (d_2066_v31_)[index334_] = d_2023_v0_
            d_2069_v32_: _dafny.Map
            d_2069_v32_ = _dafny.Map({d_2025_v2_: d_2026_v3_})
            d_2070_v33_: _dafny.Seq
            d_2070_v33_ = _dafny.SeqWithoutIsStrInference([(((d_2069_v32_)[d_2025_v2_] if (d_2025_v2_) in (d_2069_v32_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2071_i7_ in range(default__.abs(526))]))).set(default__.safeIndex(d_2032_v7_.f9, len(((d_2069_v32_)[d_2025_v2_] if (d_2025_v2_) in (d_2069_v32_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2072_i7_ in range(default__.abs(526))])))), d_2023_v0_)])
            d_2070_v33_ = _dafny.SeqWithoutIsStrInference([(d_2026_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chsubpa")))])
        d_2073_v34_: _dafny.Map
        d_2073_v34_ = _dafny.Map({d_2025_v2_: d_2023_v0_})
        d_2073_v34_ = (d_2073_v34_).set(False, d_2023_v0_)
        d_2074_v36_: _dafny.Map
        d_2074_v36_ = _dafny.Map({d_2025_v2_: p0})
        d_2075_v37_: D18
        def iife175_():
            coll81_ = _dafny.Map()
            compr_81_: _dafny.Map
            for compr_81_ in (_dafny.SeqWithoutIsStrInference([d_2074_v36_, d_2074_v36_])).Elements:
                d_2076_v35_: _dafny.Map = compr_81_
                if (d_2076_v35_) in (_dafny.SeqWithoutIsStrInference([d_2074_v36_, d_2074_v36_])):
                    coll81_[d_2076_v35_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            return _dafny.Map(coll81_)
        d_2075_v37_ = D18_DC47(iife175_()
)
        pat_let_tv56_ = d_2025_v2_
        pat_let_tv57_ = p0
        pat_let_tv58_ = d_2042_v15_
        pat_let_tv59_ = d_2042_v15_
        def lambda100_(source22_):
            if source22_.is_DC48:
                return pat_let_tv56_
            elif source22_.is_DC49:
                d_2077___mcc_h0_ = source22_.cf86
                d_2078_cf86_ = d_2077___mcc_h0_
                return (pat_let_tv57_) > ((self).f2)
            elif source22_.is_DC50:
                d_2079___mcc_h1_ = source22_.cf87
                d_2080___mcc_h2_ = source22_.cf88
                d_2081___mcc_h3_ = source22_.cf89
                d_2082___mcc_h4_ = source22_.cf90
                d_2083___mcc_h5_ = source22_.cf91
                d_2084_cf91_ = d_2083___mcc_h5_
                d_2085_cf90_ = d_2082___mcc_h4_
                d_2086_cf89_ = d_2081___mcc_h3_
                d_2087_cf88_ = d_2080___mcc_h2_
                d_2088_cf87_ = d_2079___mcc_h1_
                return d_2085_cf90_
            elif True:
                d_2089___mcc_h6_ = source22_.cf85
                d_2090_cf85_ = d_2089___mcc_h6_
                return (pat_let_tv58_).ispropersubset(pat_let_tv59_)

        rhs246_ = lambda100_(d_2075_v37_)
        rhs247_ = not((p0) > ((self).f2))
        r1 = rhs246_
        r1 = rhs247_
        r0 = (False if d_2025_v2_ else d_2025_v2_)
        r1 = default__.fm17(globalState)
        return r0, r1

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_2091_v0_: bool
        d_2091_v0_ = True
        d_2091_v0_ = d_2091_v0_
        d_2092_v1_: _dafny.Array
        nw334_ = _dafny.Array(False, 20)
        d_2092_v1_ = nw334_
        index335_ = default__.safeIndex(807, (d_2092_v1_).length(0))
        (d_2092_v1_)[index335_] = d_2091_v0_
        d_2093_v2_: bool
        d_2093_v2_ = d_2091_v0_
        index336_ = default__.safeIndex(807, (d_2092_v1_).length(0))
        (d_2092_v1_)[index336_] = d_2093_v2_
        if (p0) <= (p0):
            d_2094_v3_: int
            d_2094_v3_ = 14
            d_2095_v4_: _dafny.Array
            nw335_ = _dafny.Array(False, 10)
            d_2095_v4_ = nw335_
            d_2096_v5_: _dafny.MultiSet
            d_2096_v5_ = _dafny.MultiSet([d_2095_v4_])
            index337_ = default__.safeIndex(23, (d_2095_v4_).length(0))
            (d_2095_v4_)[index337_] = not(not(d_2091_v0_))
            index338_ = default__.safeIndex(23, (d_2095_v4_).length(0))
            rhs248_ = (self).f2
            rhs249_ = d_2096_v5_
            rhs250_ = not((d_2094_v3_) == (d_2094_v3_))
            rhs251_ = not(not(d_2091_v0_))
            lhs138_ = d_2095_v4_
            lhs139_ = default__.safeIndex(23, (d_2095_v4_).length(0))
            d_2094_v3_ = rhs248_
            d_2096_v5_ = rhs249_
            lhs138_[lhs139_] = rhs250_
            d_2091_v0_ = rhs251_
            d_2097_v6_: _dafny.Map
            d_2097_v6_ = _dafny.Map({p0: (d_2095_v4_)[default__.safeIndex(23, (d_2095_v4_).length(0))]})
            d_2098_v9_: _dafny.Map
            d_2098_v9_ = _dafny.Map({d_2091_v0_: (d_2095_v4_)[default__.safeIndex(23, (d_2095_v4_).length(0))]})
            d_2099_v10_: _dafny.Set
            d_2099_v10_ = _dafny.Set({d_2097_v6_, d_2097_v6_, _dafny.Map({len(d_2098_v9_): d_2091_v0_})})
            index339_ = default__.safeIndex(23, (d_2095_v4_).length(0))
            def iife176_():
                coll82_ = _dafny.Map()
                compr_82_: int
                for compr_82_ in _dafny.IntegerRange(567, -187):
                    d_2100_v7_: int = compr_82_
                    if ((567) <= (d_2100_v7_)) and ((d_2100_v7_) < (-187)):
                        coll82_[(d_2100_v7_) - ((self).f2)] = d_2091_v0_
                return _dafny.Map(coll82_)
            def iife177_():
                coll83_ = _dafny.Map()
                compr_83_: int
                for compr_83_ in _dafny.IntegerRange(-911, 126):
                    d_2101_v8_: int = compr_83_
                    if ((-911) <= (d_2101_v8_)) and ((d_2101_v8_) < (126)):
                        coll83_[(d_2101_v8_) - ((self).f2)] = not((d_2095_v4_)[default__.safeIndex(23, (d_2095_v4_).length(0))])
                return _dafny.Map(coll83_)
            (d_2095_v4_)[index339_] = (_dafny.Set({d_2097_v6_, iife176_()
            , iife177_()
            })).isdisjoint(d_2099_v10_)
            d_2102_v11_: _dafny.Map
            d_2102_v11_ = _dafny.Map({(0) - (((d_2096_v5_) | (d_2096_v5_)).cardinality): p0})
            d_2102_v11_ = (d_2102_v11_).set((-793) + (len(_dafny.Set({(self).f2}))), (self).f2)
            d_2103_v12_: _dafny.Seq
            d_2103_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            d_2104_v13_: _dafny.Map
            d_2104_v13_ = _dafny.Map({(d_2094_v3_) >= (d_2094_v3_): d_2103_v12_})
            d_2105_v14_: str
            d_2105_v14_ = _dafny.CodePoint('u')
            d_2104_v13_ = (d_2104_v13_).set(d_2091_v0_, ((d_2103_v12_).set(default__.safeIndex((self).f2, len(d_2103_v12_)), d_2105_v14_)) + (d_2103_v12_))
            d_2106_v15_: _dafny.Seq
            d_2106_v15_ = _dafny.SeqWithoutIsStrInference([d_2091_v0_])
            d_2107_v16_: _dafny.MultiSet
            d_2107_v16_ = _dafny.MultiSet([d_2094_v3_])
            d_2108_v17_: _dafny.Set
            d_2108_v17_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_2105_v14_ for d_2109_i1_ in range(default__.abs(492))])) for d_2110_i0_ in range(default__.abs(-438))])), d_2094_v3_, (self).f2})
            d_2111_v18_: _dafny.Set
            d_2111_v18_ = _dafny.Set({len(d_2108_v17_), (self).f2})
            d_2112_v19_: _dafny.Array
            nw336_ = _dafny.Array(None, 27)
            nw336_[int(0)] = p0
            nw336_[int(1)] = d_2094_v3_
            nw336_[int(2)] = p0
            nw336_[int(3)] = (self).f2
            nw336_[int(4)] = len(_dafny.SeqWithoutIsStrInference([d_2094_v3_, p0]))
            nw336_[int(5)] = len(d_2106_v15_)
            nw336_[int(6)] = d_2094_v3_
            nw336_[int(7)] = 292
            nw336_[int(8)] = (self).f2
            nw336_[int(9)] = (0) - (p0)
            nw336_[int(10)] = p0
            nw336_[int(11)] = (self).f2
            nw336_[int(12)] = p0
            nw336_[int(13)] = p0
            nw336_[int(14)] = d_2094_v3_
            nw336_[int(15)] = (self).f2
            nw336_[int(16)] = (self).f2
            nw336_[int(17)] = p0
            nw336_[int(18)] = len(d_2103_v12_)
            nw336_[int(19)] = (_dafny.MultiSet(d_2106_v15_)).cardinality
            nw336_[int(20)] = d_2094_v3_
            nw336_[int(21)] = d_2094_v3_
            nw336_[int(22)] = len(d_2102_v11_)
            nw336_[int(23)] = (d_2107_v16_).cardinality
            nw336_[int(24)] = len(d_2108_v17_)
            nw336_[int(25)] = d_2094_v3_
            nw336_[int(26)] = d_2094_v3_
            d_2112_v19_ = nw336_
            pat_let_tv60_ = d_2094_v3_
            d_2113_v20_: C5
            nw337_ = C5()
            def iife178_(_pat_let47_0):
                def iife179_(d_2114_dt__update__tmp_h0_):
                    def iife180_(_pat_let48_0):
                        def iife181_(d_2115_dt__update_hcf2_h0_):
                            return D0_DC1(d_2115_dt__update_hcf2_h0_)
                        return iife181_(_pat_let48_0)
                    return iife180_((0) - ((0) - (pat_let_tv60_)))
                return iife179_(_pat_let47_0)
            nw337_.ctor__(d_2112_v19_, (D2_DC5(p0, (self).f0, (d_2095_v4_)[default__.safeIndex(23, (d_2095_v4_).length(0))])).cf9, iife178_(self.f1))
            d_2113_v20_ = nw337_
            d_2116_v21_: _dafny.MultiSet
            d_2116_v21_ = _dafny.MultiSet([d_2113_v20_, d_2113_v20_])
            d_2116_v21_ = d_2116_v21_
        elif True:
            d_2117_v22_: int
            d_2117_v22_ = 150
            d_2117_v22_ = (d_2117_v22_ if not(True) else (self).f2)
            d_2118_v23_: _dafny.Seq
            d_2118_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uefbomho"))
            d_2119_v24_: D6
            d_2119_v24_ = D6_DC14(d_2118_v23_)
            source23_ = d_2119_v24_
            if source23_.is_DC15:
                d_2120___mcc_h0_ = source23_.cf30
                d_2121___mcc_h1_ = source23_.cf31
                d_2122___mcc_h2_ = source23_.cf32
                d_2123_cf32_ = d_2122___mcc_h2_
                d_2124_cf31_ = d_2121___mcc_h1_
                d_2125_cf30_ = d_2120___mcc_h0_
                d_2126_v25_: _dafny.Array
                def lambda101_(d_2127_cf32_):
                    def lambda102_(d_2128_i2_):
                        return _dafny.Set({not(d_2127_cf32_), d_2127_cf32_})

                    return lambda102_

                init53_ = lambda101_(d_2123_cf32_)
                nw338_ = _dafny.Array(None, 20)
                for i0_53_ in range(nw338_.length(0)):
                    nw338_[i0_53_] = init53_(i0_53_)
                d_2126_v25_ = nw338_
                d_2129_v26_: _dafny.Map
                d_2129_v26_ = _dafny.Map({(d_2117_v22_ if False else d_2117_v22_): d_2126_v25_})
                d_2129_v26_ = (d_2129_v26_).set(d_2117_v22_, d_2126_v25_)
                d_2123_cf32_ = True
                d_2130_v27_: _dafny.Array
                nw339_ = _dafny.Array(False, 6)
                d_2130_v27_ = nw339_
                index340_ = default__.safeIndex(669, (d_2130_v27_).length(0))
                (d_2130_v27_)[index340_] = d_2091_v0_
                index341_ = default__.safeIndex(9, (d_2130_v27_).length(0))
                (d_2130_v27_)[index341_] = d_2123_cf32_
                d_2131_v28_: _dafny.Seq
                d_2131_v28_ = _dafny.SeqWithoutIsStrInference([d_2123_cf32_, d_2091_v0_, d_2123_cf32_, d_2123_cf32_])
                index342_ = default__.safeIndex(669, (d_2130_v27_).length(0))
                index343_ = default__.safeIndex(9, (d_2130_v27_).length(0))
                rhs252_ = (d_2091_v0_ if (((self).f0)[default__.safeIndex(len(d_2118_v23_), len((self).f0))]) >= (p0) else (d_2131_v28_)[default__.safeIndex((self).f2, len(d_2131_v28_))])
                rhs253_ = d_2091_v0_
                rhs254_ = d_2123_cf32_
                rhs255_ = True
                rhs256_ = d_2091_v0_
                lhs140_ = d_2130_v27_
                lhs141_ = default__.safeIndex(669, (d_2130_v27_).length(0))
                lhs142_ = d_2130_v27_
                lhs143_ = default__.safeIndex(9, (d_2130_v27_).length(0))
                lhs140_[lhs141_] = rhs252_
                lhs142_[lhs143_] = rhs253_
                d_2091_v0_ = rhs254_
                d_2091_v0_ = rhs255_
                d_2091_v0_ = rhs256_
                d_2132_v29_: _dafny.MultiSet
                d_2132_v29_ = _dafny.MultiSet([d_2091_v0_])
                (self).m4(d_2091_v0_, 105, d_2132_v29_, p0, globalState)
            elif source23_.is_DC14:
                d_2133___mcc_h3_ = source23_.cf29
                d_2134_cf29_ = d_2133___mcc_h3_
                d_2135_v30_: _dafny.Array
                nw340_ = _dafny.Array(False, 19)
                d_2135_v30_ = nw340_
                d_2136_v31_: D18
                d_2136_v31_ = D18_DC50(not(d_2091_v0_), d_2135_v30_, False, False, d_2134_cf29_)
                index344_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                (d_2135_v30_)[index344_] = (d_2136_v31_).cf87
                index345_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                (d_2135_v30_)[index345_] = d_2091_v0_
                d_2137_v32_: C3
                nw341_ = C3()
                nw341_.ctor__()
                d_2137_v32_ = nw341_
                d_2138_v33_: str
                d_2138_v33_ = _dafny.CodePoint('m')
                d_2139_v34_: _dafny.MultiSet
                d_2139_v34_ = _dafny.MultiSet([d_2138_v33_, d_2138_v33_])
                d_2140_v35_: _dafny.Seq
                d_2140_v35_ = _dafny.SeqWithoutIsStrInference([d_2139_v34_])
                index346_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                index347_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                rhs257_ = ((d_2140_v35_).set(default__.safeIndex(287, len(d_2140_v35_)), d_2139_v34_)) < (_dafny.SeqWithoutIsStrInference([d_2139_v34_ for d_2141_i3_ in range(default__.abs(127))]))
                rhs258_ = (d_2135_v30_)[default__.safeIndex(206, (d_2135_v30_).length(0))]
                rhs259_ = not((((self).f2) > (len(d_2134_cf29_))) and (not((d_2135_v30_)[default__.safeIndex(206, (d_2135_v30_).length(0))])))
                rhs260_ = (0) - ((0) - ((default__.safeDivisionInt((0) - ((self).fm6((self).f2, (self).f2, globalState)), 833)) - ((self).f2)))
                lhs144_ = d_2135_v30_
                lhs145_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                lhs146_ = d_2135_v30_
                lhs147_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                lhs144_[lhs145_] = rhs257_
                lhs146_[lhs147_] = rhs258_
                d_2091_v0_ = rhs259_
                d_2117_v22_ = rhs260_
                d_2142_v36_: _dafny.Set
                d_2142_v36_ = _dafny.Set({_dafny.CodePoint('s')})
                d_2143_v37_: _dafny.Set
                d_2143_v37_ = _dafny.Set({d_2117_v22_, (self).f2, 25})
                d_2144_v38_: _dafny.Set
                d_2144_v38_ = _dafny.Set({d_2091_v0_})
                d_2145_v39_: _dafny.Map
                d_2145_v39_ = _dafny.Map({(0) - (d_2117_v22_): (0) - ((self).f2)})
                d_2146_v40_: _dafny.Map
                d_2146_v40_ = _dafny.Map({d_2138_v33_: d_2138_v33_})
                d_2147_v41_: _dafny.Seq
                d_2147_v41_ = _dafny.SeqWithoutIsStrInference([d_2146_v40_, d_2146_v40_])
                d_2148_v42_: _dafny.MultiSet
                d_2148_v42_ = _dafny.MultiSet([d_2091_v0_])
                d_2149_v43_: _dafny.Seq
                d_2149_v43_ = _dafny.SeqWithoutIsStrInference([False])
                d_2150_v44_: _dafny.Array
                nw342_ = _dafny.Array(None, 25)
                nw342_[int(0)] = 994
                nw342_[int(1)] = d_2117_v22_
                nw342_[int(2)] = (self).f2
                nw342_[int(3)] = d_2117_v22_
                nw342_[int(4)] = 314
                nw342_[int(5)] = len(d_2144_v38_)
                nw342_[int(6)] = (self).f2
                nw342_[int(7)] = d_2117_v22_
                nw342_[int(8)] = ((d_2145_v39_)[p0] if (p0) in (d_2145_v39_) else default__.fm44(d_2091_v0_, globalState))
                nw342_[int(9)] = p0
                nw342_[int(10)] = len((self).f0)
                nw342_[int(11)] = (0) - (len((self).f0))
                nw342_[int(12)] = (self).f2
                nw342_[int(13)] = len((d_2147_v41_)[default__.safeIndex((0) - ((d_2148_v42_).cardinality), len(d_2147_v41_))])
                nw342_[int(14)] = len(d_2134_cf29_)
                nw342_[int(15)] = d_2117_v22_
                nw342_[int(16)] = (self).f2
                nw342_[int(17)] = len((self).f0)
                nw342_[int(18)] = p0
                nw342_[int(19)] = d_2117_v22_
                nw342_[int(20)] = 266
                nw342_[int(21)] = d_2117_v22_
                nw342_[int(22)] = (_dafny.MultiSet(d_2149_v43_)).cardinality
                nw342_[int(23)] = (self).f2
                nw342_[int(24)] = p0
                d_2150_v44_ = nw342_
                d_2151_v45_: _dafny.Seq
                d_2151_v45_ = _dafny.SeqWithoutIsStrInference([d_2150_v44_, d_2150_v44_, d_2150_v44_])
                d_2152_v46_: _dafny.Map
                d_2152_v46_ = _dafny.Map({(d_2135_v30_)[default__.safeIndex(206, (d_2135_v30_).length(0))]: p0})
                d_2153_v47_: _dafny.Map
                d_2153_v47_ = _dafny.Map({(d_2151_v45_)[default__.safeIndex(len(d_2152_v46_), len(d_2151_v45_))]: (self).f2})
                d_2154_v48_: _dafny.MultiSet
                d_2154_v48_ = _dafny.MultiSet([len(default__.fm25(default__.fm64(globalState), (default__.fm16(d_2118_v23_, (_dafny.MultiSet([(self).f2, len(d_2149_v43_), d_2117_v22_, (self).f2, len(d_2143_v37_)])).cardinality, globalState)).set(default__.safeIndex((self).f2, len(default__.fm16(d_2118_v23_, (_dafny.MultiSet([(self).f2, len(d_2149_v43_), d_2117_v22_, (self).f2, len(d_2143_v37_)])).cardinality, globalState))), d_2138_v33_), globalState)), d_2117_v22_])
                d_2155_v49_: _dafny.Map
                d_2155_v49_ = _dafny.Map({d_2138_v33_: p0})
                d_2156_v50_: _dafny.Map
                d_2156_v50_ = _dafny.Map({p0: (d_2135_v30_)[default__.safeIndex(206, (d_2135_v30_).length(0))]})
                index348_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                def iife182_():
                    coll84_ = _dafny.Set()
                    compr_84_: int
                    for compr_84_ in (d_2156_v50_).keys.Elements:
                        d_2157_v51_: int = compr_84_
                        if (d_2157_v51_) in (d_2156_v50_):
                            coll84_ = coll84_.union(_dafny.Set([default__.safeDivisionInt(d_2157_v51_, len(_dafny.SeqWithoutIsStrInference([len((D2_DC5(len(_dafny.SeqWithoutIsStrInference([True, False, True, True])), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({671: -9}))]), True)).cf9)])))]))
                    return _dafny.Set(coll84_)
                rhs261_ = _dafny.Set({d_2138_v33_, d_2138_v33_, d_2138_v33_, d_2138_v33_, d_2138_v33_})
                rhs262_ = d_2091_v0_
                rhs263_ = (_dafny.Set({(d_2154_v48_).cardinality, 12, 535, len(d_2155_v49_)})) - ((iife182_()
                ).intersection(d_2143_v37_))
                rhs264_ = (d_2153_v47_) | (d_2153_v47_)
                rhs265_ = (0) - (default__.safeDivisionInt(p0, (d_2117_v22_ if (d_2135_v30_)[default__.safeIndex(206, (d_2135_v30_).length(0))] else (self).f2)))
                lhs148_ = d_2135_v30_
                lhs149_ = default__.safeIndex(206, (d_2135_v30_).length(0))
                d_2142_v36_ = rhs261_
                lhs148_[lhs149_] = rhs262_
                d_2143_v37_ = rhs263_
                d_2153_v47_ = rhs264_
                d_2117_v22_ = rhs265_
            elif True:
                d_2158___mcc_h4_ = source23_.cf33
                d_2159_cf33_ = d_2158___mcc_h4_
                d_2160_v52_: C4
                nw343_ = C4()
                nw343_.ctor__()
                d_2160_v52_ = nw343_
                d_2091_v0_ = d_2091_v0_
                d_2161_v53_: _dafny.Array
                nw344_ = _dafny.Array(_dafny.Seq({}), 20)
                d_2161_v53_ = nw344_
                d_2162_v54_: _dafny.Seq
                d_2162_v54_ = _dafny.SeqWithoutIsStrInference([True])
                index349_ = default__.safeIndex(956, (d_2161_v53_).length(0))
                (d_2161_v53_)[index349_] = d_2162_v54_
                d_2163_v55_: str
                d_2163_v55_ = _dafny.CodePoint('g')
                d_2164_v56_: _dafny.Map
                d_2164_v56_ = _dafny.Map({d_2163_v55_: d_2162_v54_})
                d_2165_v57_: D15
                d_2165_v57_ = D15_DC38(d_2163_v55_)
                index350_ = default__.safeIndex(956, (d_2161_v53_).length(0))
                (d_2161_v53_)[index350_] = ((d_2164_v56_)[(d_2165_v57_).cf71] if ((d_2165_v57_).cf71) in (d_2164_v56_) else ((d_2162_v54_) + (d_2162_v54_)).set(default__.safeIndex(d_2117_v22_, len((d_2162_v54_) + (d_2162_v54_))), d_2091_v0_))
                d_2166_v58_: _dafny.Array
                nw345_ = _dafny.Array(int(0), 4)
                d_2166_v58_ = nw345_
                index351_ = default__.safeIndex(442, (d_2166_v58_).length(0))
                (d_2166_v58_)[index351_] = (self).f2
                d_2167_v59_: _dafny.MultiSet
                d_2167_v59_ = _dafny.MultiSet([_dafny.CodePoint('i')])
                index352_ = default__.safeIndex(708, (d_2166_v58_).length(0))
                (d_2166_v58_)[index352_] = (d_2167_v59_).cardinality
                d_2168_v60_: _dafny.Seq
                d_2168_v60_ = _dafny.SeqWithoutIsStrInference([d_2166_v58_, d_2166_v58_, d_2166_v58_, d_2166_v58_])
                d_2169_v61_: _dafny.MultiSet
                d_2169_v61_ = _dafny.MultiSet([False])
                index353_ = default__.safeIndex(442, (d_2166_v58_).length(0))
                index354_ = default__.safeIndex(708, (d_2166_v58_).length(0))
                rhs266_ = len(d_2168_v60_)
                rhs267_ = (d_2169_v61_).cardinality
                lhs150_ = d_2166_v58_
                lhs151_ = default__.safeIndex(442, (d_2166_v58_).length(0))
                lhs152_ = d_2166_v58_
                lhs153_ = default__.safeIndex(708, (d_2166_v58_).length(0))
                lhs150_[lhs151_] = rhs266_
                lhs152_[lhs153_] = rhs267_
            pat_let_tv61_ = d_2117_v22_
            d_2170_v62_: int
            d_2171_v63_: _dafny.Map
            d_2172_v64_: int
            out52_: int
            out53_: _dafny.Map
            out54_: int
            def iife183_(_pat_let49_0):
                def iife184_(d_2173_dt__update__tmp_h1_):
                    def iife185_(_pat_let50_0):
                        def iife186_(d_2174_dt__update_hcf1_h0_):
                            return D0_DC0((d_2173_dt__update__tmp_h1_).cf0, d_2174_dt__update_hcf1_h0_)
                        return iife186_(_pat_let50_0)
                    return iife185_(pat_let_tv61_)
                return iife184_(_pat_let49_0)
            out52_, out53_, out54_ = (self).m3(len(_dafny.Map({len(d_2118_v23_): d_2091_v0_})), iife183_((self).f3), True, globalState)
            d_2170_v62_ = out52_
            d_2171_v63_ = out53_
            d_2172_v64_ = out54_
            d_2175_v65_: D7
            d_2175_v65_ = D7_DC19((self).f2)
            source24_ = d_2175_v65_
            if source24_.is_DC18:
                d_2176_v66_: _dafny.Map
                d_2176_v66_ = _dafny.Map({_dafny.CodePoint('s'): self.f1})
                d_2091_v0_ = (default__.fm44(not((self).fm5((self).f2, not(d_2091_v0_), d_2176_v66_, d_2091_v0_, globalState)), globalState)) == (-859)
                d_2177_v67_: _dafny.Seq
                d_2177_v67_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2170_v62_: d_2091_v0_})])
                d_2178_v68_: _dafny.Seq
                d_2178_v68_ = _dafny.SeqWithoutIsStrInference([len((d_2177_v67_)[default__.safeIndex(d_2170_v62_, len(d_2177_v67_))]), -681])
                d_2179_v69_: _dafny.Seq
                d_2179_v69_ = _dafny.SeqWithoutIsStrInference([d_2091_v0_])
                d_2180_v70_: _dafny.Array
                nw346_ = _dafny.Array(int(0), 20)
                d_2180_v70_ = nw346_
                d_2181_v71_: D13
                d_2181_v71_ = D13_DC32(d_2118_v23_, (d_2179_v69_)[default__.safeIndex(10, len(d_2179_v69_))], d_2117_v22_, d_2180_v70_)
                rhs268_ = (d_2181_v71_).cf58
                rhs269_ = ((d_2178_v68_) + (d_2178_v68_)) + (d_2178_v68_)
                d_2091_v0_ = rhs268_
                d_2178_v68_ = rhs269_
                d_2182_v72_: _dafny.Array
                nw347_ = _dafny.Array(False, 13)
                d_2182_v72_ = nw347_
                d_2182_v72_ = d_2182_v72_
                d_2091_v0_ = d_2091_v0_
            elif source24_.is_DC19:
                d_2183___mcc_h5_ = source24_.cf35
                d_2184_cf35_ = d_2183___mcc_h5_
                d_2091_v0_ = ((d_2118_v23_) + (d_2118_v23_)) <= ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2185_i4_ in range(default__.abs(588))])) + (d_2118_v23_))
                d_2172_v64_ = default__.fm44(True, globalState)
                d_2091_v0_ = d_2091_v0_
                d_2172_v64_ = (self).f2
            elif True:
                d_2186___mcc_h6_ = source24_.cf34
                d_2187_cf34_ = d_2186___mcc_h6_
                d_2188_v73_: _dafny.MultiSet
                d_2188_v73_ = _dafny.MultiSet([d_2170_v62_, d_2117_v22_, (self).fm6(default__.fm32(990, globalState), p0, globalState), p0, d_2172_v64_])
                d_2091_v0_ = (d_2188_v73_).issubset((d_2188_v73_) - (_dafny.MultiSet([(self).f2])))
                d_2189_v74_: _dafny.Array
                nw348_ = _dafny.Array(int(0), 13)
                d_2189_v74_ = nw348_
                index355_ = default__.safeIndex(230, (d_2189_v74_).length(0))
                (d_2189_v74_)[index355_] = 341
                index356_ = default__.safeIndex(230, (d_2189_v74_).length(0))
                (d_2189_v74_)[index356_] = ((0) - (len((self).f0))) - (-205)
                d_2190_v75_: str
                d_2190_v75_ = _dafny.CodePoint('x')
                d_2191_v76_: _dafny.MultiSet
                d_2191_v76_ = _dafny.MultiSet([d_2190_v75_, d_2190_v75_])
                d_2192_v77_: int
                d_2193_v78_: _dafny.Map
                d_2194_v79_: int
                out55_: int
                out56_: _dafny.Map
                out57_: int
                out55_, out56_, out57_ = (self).m3(147, (self).f3, (d_2191_v76_).issubset(d_2191_v76_), globalState)
                d_2192_v77_ = out55_
                d_2193_v78_ = out56_
                d_2194_v79_ = out57_
                d_2195_v80_: _dafny.Map
                d_2195_v80_ = _dafny.Map({d_2091_v0_: (d_2187_cf34_)[default__.safeIndex(len(d_2118_v23_), len(d_2187_cf34_))]})
                d_2091_v0_ = ((d_2195_v80_)[d_2091_v0_] if (d_2091_v0_) in (d_2195_v80_) else d_2091_v0_)
            d_2196_v81_: _dafny.Seq
            d_2196_v81_ = _dafny.SeqWithoutIsStrInference([d_2091_v0_, d_2091_v0_, default__.fm17(globalState)])
            d_2197_v82_: _dafny.Map
            d_2197_v82_ = _dafny.Map({(self).fm7(d_2172_v64_, globalState): d_2091_v0_})
            d_2198_v83_: _dafny.Map
            d_2198_v83_ = _dafny.Map({d_2117_v22_: d_2091_v0_})
            rhs270_ = _dafny.SeqWithoutIsStrInference([((d_2197_v82_)[d_2118_v23_] if (d_2118_v23_) in (d_2197_v82_) else ((d_2198_v83_)[d_2172_v64_] if (d_2172_v64_) in (d_2198_v83_) else d_2091_v0_))])
            rhs271_ = d_2172_v64_
            d_2196_v81_ = rhs270_
            d_2170_v62_ = rhs271_
        d_2199_v84_: _dafny.Array
        def lambda103_(d_2200_v0_):
            def lambda104_(d_2201_i5_):
                return d_2200_v0_

            return lambda104_

        init54_ = lambda103_(d_2091_v0_)
        nw349_ = _dafny.Array(None, 5)
        for i0_54_ in range(nw349_.length(0)):
            nw349_[i0_54_] = init54_(i0_54_)
        d_2199_v84_ = nw349_
        d_2199_v84_ = (d_2199_v84_ if d_2091_v0_ else d_2199_v84_)
        d_2202_v85_: int
        d_2202_v85_ = 844
        d_2202_v85_ = (0) - (p0)
        d_2203_v86_: _dafny.Seq
        d_2203_v86_ = _dafny.SeqWithoutIsStrInference([d_2091_v0_, not (d_2091_v0_) or (d_2091_v0_)])
        d_2204_i6_: int
        d_2204_i6_ = 0
        with _dafny.label("9"):
            while (d_2203_v86_)[default__.safeIndex((len(d_2203_v86_) if d_2091_v0_ else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2232_i7_ in range(default__.abs(73))]))), len(d_2203_v86_))]:
                with _dafny.c_label("9"):
                    if (d_2204_i6_) >= (100):
                        raise _dafny.Break("9")
                    d_2204_i6_ = (d_2204_i6_) + (1)
                    d_2202_v85_ = 724
                    if not((D6_DC15(p0, default__.fm20(d_2091_v0_, globalState), not(d_2091_v0_))).cf32):
                        d_2205_v87_: _dafny.Seq
                        d_2205_v87_ = _dafny.SeqWithoutIsStrInference([(d_2199_v84_ if d_2091_v0_ else d_2199_v84_), d_2199_v84_, d_2199_v84_])
                        d_2205_v87_ = d_2205_v87_
                        d_2206_v88_: _dafny.Array
                        def lambda105_(d_2207_v86_):
                            def lambda106_(d_2208_i8_):
                                return d_2207_v86_

                            return lambda106_

                        init55_ = lambda105_(d_2203_v86_)
                        nw350_ = _dafny.Array(None, 18)
                        for i0_55_ in range(nw350_.length(0)):
                            nw350_[i0_55_] = init55_(i0_55_)
                        d_2206_v88_ = nw350_
                        index357_ = default__.safeIndex(89, (d_2206_v88_).length(0))
                        (d_2206_v88_)[index357_] = d_2203_v86_
                        d_2209_v89_: D16
                        d_2209_v89_ = D16_DC44((self).f2)
                        d_2210_v90_: _dafny.Map
                        d_2210_v90_ = _dafny.Map({(d_2209_v89_).cf82: d_2203_v86_})
                        d_2211_v91_: _dafny.MultiSet
                        d_2211_v91_ = _dafny.MultiSet([d_2203_v86_])
                        index358_ = default__.safeIndex(89, (d_2206_v88_).length(0))
                        (d_2206_v88_)[index358_] = ((d_2210_v90_)[((d_2211_v91_).set(_dafny.SeqWithoutIsStrInference([d_2091_v0_, d_2091_v0_, False, not(d_2091_v0_), default__.fm17(globalState)]), default__.abs(d_2202_v85_))).cardinality] if (((d_2211_v91_).set(_dafny.SeqWithoutIsStrInference([d_2091_v0_, d_2091_v0_, False, not(d_2091_v0_), default__.fm17(globalState)]), default__.abs(d_2202_v85_))).cardinality) in (d_2210_v90_) else d_2203_v86_)
                        d_2202_v85_ = p0
                        d_2212_v92_: _dafny.Array
                        def lambda107_(d_2213_p0_):
                            def lambda108_(d_2214_i9_):
                                return default__.safeDivisionInt(d_2214_i9_, d_2213_p0_)

                            return lambda108_

                        init56_ = lambda107_(p0)
                        nw351_ = _dafny.Array(None, 2)
                        for i0_56_ in range(nw351_.length(0)):
                            nw351_[i0_56_] = init56_(i0_56_)
                        d_2212_v92_ = nw351_
                        d_2215_v93_: T1
                        nw352_ = C5()
                        nw352_.ctor__(d_2212_v92_, (self).f0, self.f1)
                        d_2215_v93_ = nw352_
                        d_2216_v94_: _dafny.Seq
                        d_2216_v94_ = _dafny.SeqWithoutIsStrInference([d_2215_v93_, d_2215_v93_])
                        d_2217_v95_: _dafny.Array
                        nw353_ = _dafny.Array(None, 13)
                        nw353_[int(0)] = d_2215_v93_
                        nw353_[int(1)] = (d_2215_v93_ if d_2091_v0_ else d_2215_v93_)
                        nw353_[int(2)] = d_2215_v93_
                        nw353_[int(3)] = d_2215_v93_
                        nw353_[int(4)] = d_2215_v93_
                        nw353_[int(5)] = d_2215_v93_
                        nw353_[int(6)] = (d_2216_v94_)[default__.safeIndex(d_2202_v85_, len(d_2216_v94_))]
                        nw353_[int(7)] = d_2215_v93_
                        nw353_[int(8)] = d_2215_v93_
                        nw353_[int(9)] = d_2215_v93_
                        nw353_[int(10)] = d_2215_v93_
                        nw353_[int(11)] = d_2215_v93_
                        nw353_[int(12)] = d_2215_v93_
                        d_2217_v95_ = nw353_
                        index359_ = default__.safeIndex(13, (d_2217_v95_).length(0))
                        (d_2217_v95_)[index359_] = d_2215_v93_
                        index360_ = default__.safeIndex(13, (d_2217_v95_).length(0))
                        (d_2217_v95_)[index360_] = d_2215_v93_
                        d_2218_v96_: str
                        d_2218_v96_ = _dafny.CodePoint('r')
                        d_2219_v97_: _dafny.Map
                        d_2219_v97_ = _dafny.Map({d_2218_v96_: d_2202_v85_})
                        d_2220_v98_: _dafny.Seq
                        d_2220_v98_ = _dafny.SeqWithoutIsStrInference([d_2203_v86_])
                        d_2221_v99_: _dafny.Set
                        d_2221_v99_ = _dafny.Set({default__.fm19(d_2091_v0_, d_2091_v0_, d_2219_v97_, not(d_2091_v0_), globalState), (d_2220_v98_)[default__.safeIndex((0) - ((self).f2), len(d_2220_v98_))], default__.fm19(d_2091_v0_, d_2091_v0_, d_2219_v97_, d_2091_v0_, globalState)})
                        d_2202_v85_ = len(d_2221_v99_)
                    elif True:
                        d_2202_v85_ = p0
                        d_2222_v100_: _dafny.Array
                        nw354_ = _dafny.Array(int(0), 17)
                        d_2222_v100_ = nw354_
                        d_2223_v101_: C5
                        nw355_ = C5()
                        nw355_.ctor__(d_2222_v100_, ((self).f0) + ((self).f0), D0_DC1(d_2202_v85_))
                        d_2223_v101_ = nw355_
                        d_2224_v102_: _dafny.Seq
                        d_2224_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                        d_2224_v102_ = (default__.fm16(d_2224_v102_, (0) - (len(d_2224_v102_)), globalState)) + (d_2224_v102_)
                        d_2225_v103_: _dafny.Map
                        d_2225_v103_ = _dafny.Map({d_2091_v0_: d_2091_v0_})
                        d_2226_v104_: str
                        d_2226_v104_ = _dafny.CodePoint('y')
                        d_2227_v105_: _dafny.Map
                        d_2227_v105_ = _dafny.Map({(d_2223_v101_).f7: _dafny.Map({d_2091_v0_: default__.fm65(len(d_2225_v103_), d_2226_v104_, (self).f2, globalState)})})
                        d_2228_v106_: _dafny.Map
                        d_2228_v106_ = _dafny.Map({d_2091_v0_: p0})
                        d_2229_v107_: _dafny.Map
                        d_2229_v107_ = _dafny.Map({True: d_2228_v106_})
                        d_2227_v105_ = (d_2227_v105_).set((d_2223_v101_).f7, d_2229_v107_)
                        d_2230_v108_: _dafny.Set
                        d_2230_v108_ = _dafny.Set({(self).f2})
                        d_2231_v109_: _dafny.Map
                        d_2231_v109_ = _dafny.Map({d_2199_v84_: d_2202_v85_})
                        rhs272_ = (d_2230_v108_) | (d_2230_v108_)
                        rhs273_ = default__.fm17(globalState)
                        rhs274_ = 363
                        rhs275_ = (d_2231_v109_) | (d_2231_v109_)
                        rhs276_ = (0) - ((0) - (((self).f2) - ((d_2223_v101_).fm6((self).f2, (0) - (-447), globalState))))
                        d_2230_v108_ = rhs272_
                        d_2091_v0_ = rhs273_
                        d_2202_v85_ = rhs274_
                        d_2231_v109_ = rhs275_
                        d_2202_v85_ = rhs276_
                    d_2202_v85_ = p0
                    d_2091_v0_ = not(not(d_2091_v0_))
                    pass
            pass
        d_2233_v110_: _dafny.Array
        nw356_ = _dafny.Array(None, 7)
        d_2233_v110_ = nw356_
        d_2234_v111_: _dafny.Map
        d_2234_v111_ = _dafny.Map({d_2091_v0_: _dafny.MultiSet([(self).f2])})
        d_2235_v112_: _dafny.Map
        d_2235_v112_ = _dafny.Map({d_2233_v110_: d_2234_v111_})
        d_2236_v113_: _dafny.MultiSet
        d_2236_v113_ = _dafny.MultiSet([d_2202_v85_])
        r0 = (((d_2235_v112_)[d_2233_v110_] if (d_2233_v110_) in (d_2235_v112_) else _dafny.Map({d_2091_v0_: d_2236_v113_}))) | (d_2234_v111_)
        return r0

    def m3(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        d_2237_i0_: int
        d_2237_i0_ = 0
        with _dafny.label("10"):
            while not(p2):
                with _dafny.c_label("10"):
                    if (d_2237_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_2237_i0_ = (d_2237_i0_) + (1)
                    d_2238_v0_: _dafny.MultiSet
                    d_2238_v0_ = _dafny.MultiSet([p2])
                    d_2239_v1_: _dafny.MultiSet
                    d_2239_v1_ = _dafny.MultiSet([(self).f2, 949, (self).f2, 303, (self).f2])
                    d_2240_v2_: _dafny.Map
                    d_2240_v2_ = _dafny.Map({not(p2): (self).f2})
                    d_2241_v3_: _dafny.Map
                    d_2241_v3_ = _dafny.Map({p2: (0) - (len(d_2240_v2_))})
                    d_2242_v4_: str
                    d_2242_v4_ = _dafny.CodePoint('m')
                    d_2243_v5_: _dafny.Map
                    d_2243_v5_ = _dafny.Map({-939: p0})
                    d_2244_v6_: _dafny.Map
                    d_2244_v6_ = _dafny.Map({981: p2})
                    d_2245_v7_: _dafny.Map
                    d_2245_v7_ = _dafny.Map({p2: _dafny.CodePoint('m')})
                    d_2246_v8_: _dafny.Set
                    d_2246_v8_ = _dafny.Set({d_2245_v7_})
                    d_2247_v9_: _dafny.Array
                    nw357_ = _dafny.Array(None, 2)
                    nw357_[int(0)] = False
                    nw357_[int(1)] = ((d_2244_v6_)[532] if (532) in (d_2244_v6_) else p2)
                    d_2247_v9_ = nw357_
                    d_2248_v10_: _dafny.Map
                    d_2248_v10_ = _dafny.Map({d_2247_v9_: p0})
                    d_2249_v11_: _dafny.Array
                    nw358_ = _dafny.Array(None, 22)
                    nw358_[int(0)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2250_i1_ in range(default__.abs(810))]))
                    nw358_[int(1)] = (d_2238_v0_).cardinality
                    nw358_[int(2)] = (self).f2
                    nw358_[int(3)] = ((0) - ((self).f2)) * ((d_2239_v1_).cardinality)
                    nw358_[int(4)] = (len(default__.fm31(-301, p2, p0, globalState)) if p2 else p0)
                    nw358_[int(5)] = p0
                    nw358_[int(6)] = len(d_2240_v2_)
                    nw358_[int(7)] = ((self).f2) + ((self).f2)
                    nw358_[int(8)] = 23
                    nw358_[int(9)] = len(_dafny.SeqWithoutIsStrInference([p2, (self).fm5((self).f2, p2, _dafny.Map({d_2242_v4_: self.f1}), p2, globalState)]))
                    nw358_[int(10)] = (self).f2
                    nw358_[int(11)] = (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f2, (0) - (((d_2243_v5_)[(0) - ((self).f2)] if ((0) - ((self).f2)) in (d_2243_v5_) else (self).f2)), (self).f2, len(d_2244_v6_)])))
                    nw358_[int(12)] = (self).f2
                    nw358_[int(13)] = (p0) + (len(_dafny.Map({len(d_2246_v8_): (self).f2})))
                    nw358_[int(14)] = 502
                    nw358_[int(15)] = (self).fm6((0) - ((self).f2), (self).f2, globalState)
                    nw358_[int(16)] = (self).f2
                    nw358_[int(17)] = (len(d_2248_v10_)) - (default__.fm44(p2, globalState))
                    nw358_[int(18)] = p0
                    nw358_[int(19)] = ((d_2240_v2_)[not(p2)] if (not(p2)) in (d_2240_v2_) else 800)
                    nw358_[int(20)] = 4
                    nw358_[int(21)] = p0
                    d_2249_v11_ = nw358_
                    index361_ = default__.safeIndex(854, (d_2249_v11_).length(0))
                    (d_2249_v11_)[index361_] = p0
                    index362_ = default__.safeIndex(854, (d_2249_v11_).length(0))
                    (d_2249_v11_)[index362_] = (p0) + (p0)
                    d_2251_v12_: _dafny.Seq
                    d_2251_v12_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                    d_2252_v13_: _dafny.Map
                    d_2252_v13_ = _dafny.Map({(D7_DC17(d_2251_v12_)).cf34: p2})
                    r0 = ((self).f0)[default__.safeIndex(len(d_2252_v13_), len((self).f0))]
                    d_2253_v14_: bool
                    d_2253_v14_ = True
                    d_2254_v15_: _dafny.Seq
                    d_2254_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hivvd"))
                    d_2253_v14_ = (self).fm5(len((d_2254_v15_) + (d_2254_v15_)), (d_2253_v14_ if True else p2), _dafny.Map({d_2242_v4_: self.f1}), (len(d_2254_v15_)) >= ((d_2249_v11_)[default__.safeIndex(854, (d_2249_v11_).length(0))]), globalState)
                    d_2255_v16_: C6
                    nw359_ = C6()
                    nw359_.ctor__(default__.safeModuloInt(len((self).f0), p0))
                    d_2255_v16_ = nw359_
                    pass
            pass
        d_2256_i2_: int
        d_2256_i2_ = 0
        with _dafny.label("11"):
            while (642) <= (default__.safeModuloInt((self).f2, p0)):
                with _dafny.c_label("11"):
                    if (d_2256_i2_) >= (100):
                        raise _dafny.Break("11")
                    d_2256_i2_ = (d_2256_i2_) + (1)
                    if p2:
                        d_2257_v17_: _dafny.Array
                        nw360_ = _dafny.Array(_dafny.Array(None, 0), 23)
                        d_2257_v17_ = nw360_
                        d_2257_v17_ = d_2257_v17_
                        d_2258_v18_: _dafny.Seq
                        d_2258_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkbca"))
                        d_2259_v19_: str
                        d_2259_v19_ = _dafny.CodePoint('h')
                        d_2260_v20_: _dafny.Map
                        d_2260_v20_ = _dafny.Map({(d_2258_v18_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snuobpymy"))).set(default__.safeIndex((self).f2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snuobpymy")))), d_2259_v19_)): (self).f2})
                        d_2260_v20_ = (d_2260_v20_).set(d_2258_v18_, default__.safeDivisionInt(459, (0) - (p0)))
                        d_2261_v21_: _dafny.Map
                        d_2261_v21_ = _dafny.Map({(self).f2: (self).f2})
                        d_2262_v22_: D13
                        d_2262_v22_ = D13_DC31(True, p0, (self).f2)
                        d_2263_v23_: _dafny.Map
                        d_2263_v23_ = _dafny.Map({False: True})
                        d_2264_v24_: _dafny.Array
                        nw361_ = _dafny.Array(None, 15)
                        nw361_[int(0)] = (self).f2
                        nw361_[int(1)] = len(((d_2258_v18_).set(default__.safeIndex(p0, len(d_2258_v18_)), d_2259_v19_) if p2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))))
                        nw361_[int(2)] = len(d_2258_v18_)
                        nw361_[int(3)] = default__.safeDivisionInt(len(_dafny.Set({p0, p0})), p0)
                        nw361_[int(4)] = default__.safeModuloInt((self).f2, (default__.fm66(len(d_2261_v21_), p2, globalState)).cf78)
                        nw361_[int(5)] = (p0 if p2 else p0)
                        nw361_[int(6)] = (d_2262_v22_).cf55
                        nw361_[int(7)] = (self).f2
                        nw361_[int(8)] = p0
                        nw361_[int(9)] = len(d_2263_v23_)
                        nw361_[int(10)] = p0
                        nw361_[int(11)] = p0
                        nw361_[int(12)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gk")))
                        nw361_[int(13)] = ((self).f2) - (default__.fm32(p0, globalState))
                        nw361_[int(14)] = p0
                        d_2264_v24_ = nw361_
                        d_2265_v25_: _dafny.MultiSet
                        d_2265_v25_ = _dafny.MultiSet([(self).f2])
                        index363_ = default__.safeIndex(102, (d_2264_v24_).length(0))
                        (d_2264_v24_)[index363_] = (d_2265_v25_).cardinality
                        d_2266_v26_: _dafny.Array
                        def lambda109_(d_2267_p2_):
                            def lambda110_(d_2268_i3_):
                                return d_2267_p2_

                            return lambda110_

                        init57_ = lambda109_(p2)
                        nw362_ = _dafny.Array(None, 25)
                        for i0_57_ in range(nw362_.length(0)):
                            nw362_[i0_57_] = init57_(i0_57_)
                        d_2266_v26_ = nw362_
                        d_2269_v27_: _dafny.Set
                        d_2269_v27_ = _dafny.Set({(d_2266_v26_ if p2 else d_2266_v26_), d_2266_v26_, d_2266_v26_})
                        d_2270_v28_: bool
                        d_2270_v28_ = False
                        index364_ = default__.safeIndex(102, (d_2264_v24_).length(0))
                        rhs277_ = default__.safeModuloInt((self).f2, p0)
                        rhs278_ = p0
                        rhs279_ = d_2269_v27_
                        rhs280_ = d_2270_v28_
                        lhs154_ = d_2264_v24_
                        lhs155_ = default__.safeIndex(102, (d_2264_v24_).length(0))
                        r2 = rhs277_
                        lhs154_[lhs155_] = rhs278_
                        d_2269_v27_ = rhs279_
                        d_2270_v28_ = rhs280_
                        index365_ = default__.safeIndex(228, (d_2266_v26_).length(0))
                        (d_2266_v26_)[index365_] = not (d_2270_v28_) or (p2)
                        index366_ = default__.safeIndex(228, (d_2266_v26_).length(0))
                        index367_ = default__.safeIndex(102, (d_2264_v24_).length(0))
                        rhs281_ = not(d_2270_v28_)
                        rhs282_ = d_2270_v28_
                        rhs283_ = (0) - ((d_2264_v24_)[default__.safeIndex(102, (d_2264_v24_).length(0))])
                        lhs156_ = d_2266_v26_
                        lhs157_ = default__.safeIndex(228, (d_2266_v26_).length(0))
                        lhs158_ = d_2264_v24_
                        lhs159_ = default__.safeIndex(102, (d_2264_v24_).length(0))
                        d_2270_v28_ = rhs281_
                        lhs156_[lhs157_] = rhs282_
                        lhs158_[lhs159_] = rhs283_
                        index368_ = default__.safeIndex(102, (d_2264_v24_).length(0))
                        (d_2264_v24_)[index368_] = ((d_2264_v24_)[default__.safeIndex(102, (d_2264_v24_).length(0))]) + (len(_dafny.SeqWithoutIsStrInference([d_2259_v19_ for d_2271_i4_ in range(default__.abs(141))])))
                    elif True:
                        d_2272_v29_: _dafny.MultiSet
                        d_2272_v29_ = _dafny.MultiSet([p2, p2, p2, p2])
                        d_2273_v30_: _dafny.MultiSet
                        d_2273_v30_ = _dafny.MultiSet([p0])
                        d_2274_v31_: _dafny.Set
                        d_2274_v31_ = _dafny.Set({(self).f2, ((d_2272_v29_)[p2] if (p2) in (d_2272_v29_) else ((d_2273_v30_)[(self).f2] if ((self).f2) in (d_2273_v30_) else (self).f2))})
                        d_2275_v32_: D9
                        d_2275_v32_ = D9_DC22(p0)
                        d_2276_v33_: _dafny.Map
                        d_2276_v33_ = _dafny.Map({d_2274_v31_: d_2275_v32_})
                        d_2276_v33_ = d_2276_v33_
                        d_2277_v34_: _dafny.Array
                        nw363_ = _dafny.Array(int(0), 6)
                        d_2277_v34_ = nw363_
                        index369_ = default__.safeIndex(39, (d_2277_v34_).length(0))
                        (d_2277_v34_)[index369_] = p0
                        index370_ = default__.safeIndex(39, (d_2277_v34_).length(0))
                        (d_2277_v34_)[index370_] = (0) - ((p0) - ((self).f2))
                        d_2278_v35_: _dafny.Seq
                        d_2278_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovve"))
                        d_2278_v35_ = d_2278_v35_
                        d_2279_v36_: bool
                        d_2279_v36_ = False
                        d_2279_v36_ = d_2279_v36_
                        d_2279_v36_ = (default__.safeDivisionInt(p0, (d_2277_v34_)[default__.safeIndex(39, (d_2277_v34_).length(0))])) != (len(((self).f0 if d_2279_v36_ else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({-225: len((self).f0)})) for d_2280_i5_ in range(default__.abs(388))]))))
                    d_2281_v37_: str
                    d_2281_v37_ = _dafny.CodePoint('l')
                    d_2282_v38_: D11
                    d_2282_v38_ = D11_DC25(d_2281_v37_, p0, (p0) != (800))
                    d_2283_v39_: _dafny.Set
                    d_2283_v39_ = _dafny.Set({(self).f2})
                    d_2284_v40_: _dafny.Map
                    d_2284_v40_ = _dafny.Map({(self).f2: p2})
                    d_2285_v41_: _dafny.Map
                    d_2285_v41_ = _dafny.Map({len(d_2283_v39_): ((d_2284_v40_)[p0] if (p0) in (d_2284_v40_) else p2)})
                    d_2286_v42_: D16
                    d_2286_v42_ = D16_DC44(p0)
                    d_2287_v43_: _dafny.Set
                    d_2287_v43_ = _dafny.Set({len((d_2284_v40_) | (_dafny.Map({p0: not(p2)}))), ((self).f2) - ((self).f2), (d_2286_v42_).cf82, p0})
                    d_2288_v44_: _dafny.Seq
                    d_2288_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntyajnp"))
                    d_2289_v45_: bool
                    d_2289_v45_ = True
                    d_2290_v46_: _dafny.Map
                    d_2290_v46_ = _dafny.Map({not(d_2289_v45_): d_2289_v45_})
                    d_2291_v47_: _dafny.Seq
                    d_2291_v47_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, p2, p2])
                    d_2292_v48_: _dafny.Set
                    d_2292_v48_ = _dafny.Set({_dafny.Map({p2: p2}), _dafny.Map({p2: d_2289_v45_}), _dafny.Map({(d_2291_v47_)[default__.safeIndex((self).f2, len(d_2291_v47_))]: p2})})
                    rhs284_ = default__.fm67((_dafny.Set({d_2290_v46_})).intersection(d_2292_v48_), globalState)
                    rhs285_ = (_dafny.Set({p0})) | (_dafny.Set({(_dafny.MultiSet([(self).f2])).cardinality}))
                    rhs286_ = (self).f2
                    rhs287_ = d_2288_v44_
                    rhs288_ = p2
                    d_2282_v38_ = rhs284_
                    d_2283_v39_ = rhs285_
                    r0 = rhs286_
                    d_2288_v44_ = rhs287_
                    d_2289_v45_ = rhs288_
                    d_2293_v49_: _dafny.Set
                    d_2293_v49_ = _dafny.Set({d_2288_v44_, _dafny.SeqWithoutIsStrInference([d_2281_v37_ for d_2294_i6_ in range(default__.abs(-745))])})
                    d_2295_v50_: _dafny.Map
                    d_2295_v50_ = _dafny.Map({d_2291_v47_: (d_2293_v49_) - (d_2293_v49_)})
                    d_2295_v50_ = (d_2295_v50_).set(d_2291_v47_, _dafny.Set({d_2288_v44_}))
                    d_2296_v51_: C2
                    nw364_ = C2()
                    nw364_.ctor__()
                    d_2296_v51_ = nw364_
                    pass
            pass
        d_2297_v52_: D16
        d_2297_v52_ = D16_DC44((self).f2)
        pat_let_tv62_ = p2
        pat_let_tv63_ = p2
        pat_let_tv64_ = p2
        pat_let_tv65_ = p2
        pat_let_tv66_ = p2
        pat_let_tv67_ = p2
        pat_let_tv68_ = p2
        pat_let_tv69_ = p2
        pat_let_tv70_ = p2
        def lambda111_(source25_):
            if source25_.is_DC42:
                d_2298___mcc_h0_ = source25_.cf76
                d_2299___mcc_h1_ = source25_.cf77
                d_2300___mcc_h2_ = source25_.cf78
                d_2301_cf78_ = d_2300___mcc_h2_
                d_2302_cf77_ = d_2299___mcc_h1_
                d_2303_cf76_ = d_2298___mcc_h0_
                return (0) - (len(_dafny.SeqWithoutIsStrInference([pat_let_tv62_])))
            elif source25_.is_DC43:
                d_2304___mcc_h3_ = source25_.cf79
                d_2305___mcc_h4_ = source25_.cf80
                d_2306___mcc_h5_ = source25_.cf81
                d_2307_cf81_ = d_2306___mcc_h5_
                d_2308_cf80_ = d_2305___mcc_h4_
                d_2309_cf79_ = d_2304___mcc_h3_
                return len(_dafny.Set({_dafny.Map({True: pat_let_tv63_}), _dafny.Map({not(pat_let_tv64_): pat_let_tv65_}), _dafny.Map({pat_let_tv66_: pat_let_tv67_}), _dafny.Map({pat_let_tv68_: True}), _dafny.Map({pat_let_tv69_: pat_let_tv70_})}))
            elif source25_.is_DC44:
                d_2310___mcc_h6_ = source25_.cf82
                d_2311_cf82_ = d_2310___mcc_h6_
                return default__.safeDivisionInt((self).f2, (self).f2)
            elif True:
                d_2312___mcc_h7_ = source25_.cf75
                d_2313_cf75_ = d_2312___mcc_h7_
                return (self).f2

        r2 = lambda111_(d_2297_v52_)
        d_2314_v53_: _dafny.Array
        nw365_ = _dafny.Array(False, 7)
        d_2314_v53_ = nw365_
        index371_ = default__.safeIndex(674, (d_2314_v53_).length(0))
        (d_2314_v53_)[index371_] = p2
        d_2315_v54_: bool
        d_2315_v54_ = True
        index372_ = default__.safeIndex(674, (d_2314_v53_).length(0))
        rhs289_ = d_2315_v54_
        rhs290_ = not(p2)
        lhs160_ = d_2314_v53_
        lhs161_ = default__.safeIndex(674, (d_2314_v53_).length(0))
        lhs160_[lhs161_] = rhs289_
        d_2315_v54_ = rhs290_
        hi16_ = ((0) - (-773) if (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))] else p0)
        for d_2316_i7_ in range((945) - (p0), hi16_):
            d_2315_v54_ = ((p0) >= ((self).f2) if p2 else (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))])
            d_2317_v55_: D12
            d_2317_v55_ = D12_DC29(True, p0, p2, d_2316_i7_, (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))])
            index373_ = default__.safeIndex(674, (d_2314_v53_).length(0))
            (d_2314_v53_)[index373_] = (p2 if ((d_2317_v55_).cf50 if p2 else (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]) else (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))])
            d_2318_v56_: _dafny.Array
            nw366_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
            d_2318_v56_ = nw366_
            d_2318_v56_ = d_2318_v56_
            d_2319_v57_: C2
            nw367_ = C2()
            nw367_.ctor__()
            d_2319_v57_ = nw367_
        d_2320_v58_: D12
        d_2320_v58_ = D12_DC28(True, (self).f2)
        d_2321_v59_: _dafny.Map
        d_2321_v59_ = _dafny.Map({not((d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]): (d_2320_v58_).cf46})
        d_2322_v61_: _dafny.Seq
        d_2322_v61_ = _dafny.SeqWithoutIsStrInference([d_2315_v54_])
        d_2323_v62_: _dafny.Map
        d_2323_v62_ = _dafny.Map({d_2322_v61_: p2})
        d_2324_v63_: _dafny.Map
        def iife187_():
            coll85_ = _dafny.Map()
            compr_85_: _dafny.Seq
            for compr_85_ in (d_2323_v62_).keys.Elements:
                d_2325_v60_: _dafny.Seq = compr_85_
                if (d_2325_v60_) in (d_2323_v62_):
                    coll85_[d_2325_v60_] = (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]
            return _dafny.Map(coll85_)
        d_2324_v63_ = _dafny.Map({(d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]: iife187_()
        })
        d_2326_v64_: _dafny.Set
        d_2326_v64_ = _dafny.Set({d_2315_v54_})
        d_2327_v65_: _dafny.Map
        d_2327_v65_ = _dafny.Map({p0: p0})
        d_2328_v66_: _dafny.Array
        def lambda112_(d_2329_p0_):
            def lambda113_(d_2330_i8_):
                return default__.safeDivisionInt(d_2330_i8_, d_2329_p0_)

            return lambda113_

        init58_ = lambda112_(p0)
        nw368_ = _dafny.Array(None, 18)
        for i0_58_ in range(nw368_.length(0)):
            nw368_[i0_58_] = init58_(i0_58_)
        d_2328_v66_ = nw368_
        d_2331_v67_: str
        d_2331_v67_ = _dafny.CodePoint('i')
        d_2332_v68_: _dafny.Map
        d_2332_v68_ = _dafny.Map({d_2328_v66_: default__.fm65((0) - ((self).f2), d_2331_v67_, p0, globalState)})
        d_2333_v69_: _dafny.Seq
        d_2333_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "haw"))
        d_2334_v70_: C0
        nw369_ = C0()
        nw369_.ctor__(d_2333_v69_)
        d_2334_v70_ = nw369_
        d_2335_v71_: _dafny.Seq
        d_2335_v71_ = _dafny.SeqWithoutIsStrInference([d_2334_v70_])
        d_2336_v72_: _dafny.Array
        nw370_ = _dafny.Array(None, 27)
        nw370_[int(0)] = ((self).f2) - (p0)
        nw370_[int(1)] = (941) * ((self).f2)
        nw370_[int(2)] = p0
        nw370_[int(3)] = (self).f2
        nw370_[int(4)] = len((self).f0)
        nw370_[int(5)] = ((self).f2) + ((self).f2)
        nw370_[int(6)] = p0
        nw370_[int(7)] = (self).f2
        nw370_[int(8)] = (len(_dafny.SeqWithoutIsStrInference([((d_2321_v59_)[(d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]] if ((d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]) in (d_2321_v59_) else p2)]))) + (323)
        nw370_[int(9)] = (self).f2
        nw370_[int(10)] = (len(((d_2324_v63_)[not(p2)] if (not(p2)) in (d_2324_v63_) else _dafny.Map({d_2322_v61_: d_2315_v54_})))) + ((self).f2)
        nw370_[int(11)] = len((d_2326_v64_) - (d_2326_v64_))
        nw370_[int(12)] = (self).f2
        nw370_[int(13)] = default__.safeDivisionInt(len(d_2327_v65_), len(((d_2332_v68_)[d_2328_v66_] if (d_2328_v66_) in (d_2332_v68_) else _dafny.Map({p2: (self).f2}))))
        nw370_[int(14)] = default__.safeDivisionInt((self).f2, 355)
        nw370_[int(15)] = (p0) * (len(d_2333_v69_))
        nw370_[int(16)] = (p0) - (len(_dafny.Set({(d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))], not((d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))]), p2, p2})))
        nw370_[int(17)] = (0) - ((self).f2)
        nw370_[int(18)] = p0
        nw370_[int(19)] = (len(d_2335_v71_) if (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))] else (self).f2)
        nw370_[int(20)] = (p0) * (default__.fm32(502, globalState))
        nw370_[int(21)] = (0) - ((len(d_2327_v65_)) + (p0))
        nw370_[int(22)] = p0
        nw370_[int(23)] = (self).f2
        nw370_[int(24)] = len((self).f0)
        nw370_[int(25)] = (len(_dafny.SeqWithoutIsStrInference([(self).f2]))) - (-469)
        nw370_[int(26)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([p0])), (self).f2)
        d_2336_v72_ = nw370_
        d_2337_v73_: D12
        d_2337_v73_ = D12_DC29(d_2315_v54_, len(d_2322_v61_), (d_2314_v53_)[default__.safeIndex(674, (d_2314_v53_).length(0))], 196, p2)
        index374_ = default__.safeIndex(644, (d_2336_v72_).length(0))
        (d_2336_v72_)[index374_] = (d_2337_v73_).cf49
        index375_ = default__.safeIndex(644, (d_2336_v72_).length(0))
        (d_2336_v72_)[index375_] = (((0) - (p0)) + (447)) + (383)
        d_2338_v74_: _dafny.Map
        d_2338_v74_ = _dafny.Map({default__.safeDivisionInt((0) - ((d_2336_v72_)[default__.safeIndex(644, (d_2336_v72_).length(0))]), (self).f2): d_2314_v53_})
        r0 = len(d_2338_v74_)
        d_2339_v75_: _dafny.Map
        d_2339_v75_ = _dafny.Map({((self).f0)[default__.safeIndex((d_2336_v72_)[default__.safeIndex(644, (d_2336_v72_).length(0))], len((self).f0))]: d_2331_v67_})
        r1 = (d_2339_v75_) | (_dafny.Map({len(d_2333_v69_): d_2331_v67_}))
        r2 = default__.safeModuloInt((d_2336_v72_)[default__.safeIndex(644, (d_2336_v72_).length(0))], p0)
        return r0, r1, r2

    def m4(self, p0, p1, p2, p3, globalState):
        d_2340_v0_: _dafny.Array
        def lambda114_(d_2341_i0_):
            return False

        init59_ = lambda114_
        nw371_ = _dafny.Array(None, 23)
        for i0_59_ in range(nw371_.length(0)):
            nw371_[i0_59_] = init59_(i0_59_)
        d_2340_v0_ = nw371_
        d_2342_v1_: _dafny.Array
        nw372_ = _dafny.Array(None, 7)
        nw372_[int(0)] = d_2340_v0_
        nw372_[int(1)] = d_2340_v0_
        nw372_[int(2)] = d_2340_v0_
        nw372_[int(3)] = d_2340_v0_
        nw372_[int(4)] = d_2340_v0_
        nw372_[int(5)] = d_2340_v0_
        nw372_[int(6)] = d_2340_v0_
        d_2342_v1_ = nw372_
        d_2343_v2_: str
        d_2343_v2_ = _dafny.CodePoint('a')
        d_2344_v3_: _dafny.Map
        d_2344_v3_ = _dafny.Map({d_2343_v2_: self.f1})
        d_2345_v4_: _dafny.Map
        d_2345_v4_ = _dafny.Map({(self).fm5(p3, not(p0), d_2344_v3_, p0, globalState): d_2342_v1_})
        d_2342_v1_ = ((d_2345_v4_)[(len((self).f0)) == (p3)] if ((len((self).f0)) == (p3)) in (d_2345_v4_) else d_2342_v1_)
        d_2346_v5_: _dafny.Seq
        d_2346_v5_ = _dafny.SeqWithoutIsStrInference([not(p0), p0, p0])
        d_2347_v6_: _dafny.Seq
        d_2347_v6_ = _dafny.SeqWithoutIsStrInference([d_2346_v5_])
        if (_dafny.SeqWithoutIsStrInference([p0])) not in ((_dafny.SeqWithoutIsStrInference([d_2346_v5_])) + (d_2347_v6_)):
            d_2348_v7_: _dafny.Set
            d_2348_v7_ = _dafny.Set({d_2340_v0_})
            d_2349_v8_: _dafny.Set
            d_2349_v8_ = d_2348_v7_
            d_2349_v8_ = d_2349_v8_
            d_2350_v9_: bool
            d_2350_v9_ = False
            d_2350_v9_ = d_2350_v9_
            if p0:
                d_2351_v10_: int
                d_2351_v10_ = 249
                d_2351_v10_ = p1
                d_2352_v11_: _dafny.Map
                d_2352_v11_ = _dafny.Map({868: not(p0)})
                d_2353_v12_: _dafny.Map
                d_2353_v12_ = _dafny.Map({d_2340_v0_: not(d_2350_v9_)})
                d_2354_v13_: _dafny.Array
                nw373_ = _dafny.Array(None, 26)
                nw373_[int(0)] = _dafny.Map({d_2340_v0_: ((d_2352_v11_)[601] if (601) in (d_2352_v11_) else d_2350_v9_)})
                nw373_[int(1)] = (_dafny.Map({d_2340_v0_: p0})) | (d_2353_v12_)
                nw373_[int(2)] = d_2353_v12_
                nw373_[int(3)] = (_dafny.Map({d_2340_v0_: p0})) | (d_2353_v12_)
                nw373_[int(4)] = (d_2353_v12_ if p0 else _dafny.Map({d_2340_v0_: p0}))
                nw373_[int(5)] = d_2353_v12_
                nw373_[int(6)] = (d_2353_v12_) | (_dafny.Map({d_2340_v0_: d_2350_v9_}))
                nw373_[int(7)] = (d_2353_v12_) | (d_2353_v12_)
                nw373_[int(8)] = (_dafny.Map({d_2340_v0_: p0})) | (d_2353_v12_)
                nw373_[int(9)] = d_2353_v12_
                nw373_[int(10)] = (d_2353_v12_) | (d_2353_v12_)
                nw373_[int(11)] = d_2353_v12_
                nw373_[int(12)] = d_2353_v12_
                nw373_[int(13)] = d_2353_v12_
                nw373_[int(14)] = (d_2353_v12_) | (d_2353_v12_)
                nw373_[int(15)] = d_2353_v12_
                nw373_[int(16)] = d_2353_v12_
                nw373_[int(17)] = d_2353_v12_
                nw373_[int(18)] = d_2353_v12_
                nw373_[int(19)] = d_2353_v12_
                nw373_[int(20)] = d_2353_v12_
                nw373_[int(21)] = (_dafny.Map({d_2340_v0_: p0})) | (d_2353_v12_)
                nw373_[int(22)] = d_2353_v12_
                nw373_[int(23)] = d_2353_v12_
                nw373_[int(24)] = d_2353_v12_
                nw373_[int(25)] = d_2353_v12_
                d_2354_v13_ = nw373_
                index376_ = default__.safeIndex(351, (d_2354_v13_).length(0))
                (d_2354_v13_)[index376_] = d_2353_v12_
                d_2355_v14_: _dafny.Map
                d_2355_v14_ = _dafny.Map({p1: (0) - ((self).f2)})
                d_2356_v15_: _dafny.Array
                def lambda115_(d_2357_v10_):
                    def lambda116_(d_2358_i1_):
                        return (d_2358_i1_) * (d_2357_v10_)

                    return lambda116_

                init60_ = lambda115_(d_2351_v10_)
                nw374_ = _dafny.Array(None, 28)
                for i0_60_ in range(nw374_.length(0)):
                    nw374_[i0_60_] = init60_(i0_60_)
                d_2356_v15_ = nw374_
                d_2359_v16_: C5
                nw375_ = C5()
                nw375_.ctor__(d_2356_v15_, _dafny.SeqWithoutIsStrInference([(self).f2, p1]), self.f1)
                d_2359_v16_ = nw375_
                d_2360_v17_: _dafny.Array
                nw376_ = _dafny.Array(None, 21)
                nw376_[int(0)] = d_2359_v16_
                nw376_[int(1)] = d_2359_v16_
                nw376_[int(2)] = d_2359_v16_
                nw376_[int(3)] = d_2359_v16_
                nw376_[int(4)] = d_2359_v16_
                nw376_[int(5)] = d_2359_v16_
                nw376_[int(6)] = d_2359_v16_
                nw376_[int(7)] = d_2359_v16_
                nw376_[int(8)] = d_2359_v16_
                nw376_[int(9)] = d_2359_v16_
                nw376_[int(10)] = d_2359_v16_
                nw376_[int(11)] = d_2359_v16_
                nw376_[int(12)] = d_2359_v16_
                nw376_[int(13)] = d_2359_v16_
                nw376_[int(14)] = d_2359_v16_
                nw376_[int(15)] = d_2359_v16_
                nw376_[int(16)] = d_2359_v16_
                nw376_[int(17)] = d_2359_v16_
                nw376_[int(18)] = d_2359_v16_
                nw376_[int(19)] = d_2359_v16_
                nw376_[int(20)] = d_2359_v16_
                d_2360_v17_ = nw376_
                d_2361_v18_: _dafny.Map
                d_2361_v18_ = _dafny.Map({p1: d_2360_v17_})
                d_2362_v19_: _dafny.Seq
                d_2362_v19_ = _dafny.SeqWithoutIsStrInference([((d_2361_v18_)[p1] if (p1) in (d_2361_v18_) else d_2360_v17_)])
                d_2363_v20_: D12
                d_2363_v20_ = D12_DC29(d_2350_v9_, p3, not(p0), len(d_2362_v19_), default__.fm30(d_2351_v10_, False, p0, globalState))
                index377_ = default__.safeIndex(351, (d_2354_v13_).length(0))
                rhs291_ = (default__.safeModuloInt(len(d_2355_v14_), 832)) - ((d_2363_v20_).cf49)
                rhs292_ = p0
                rhs293_ = d_2343_v2_
                rhs294_ = p0
                rhs295_ = ((d_2353_v12_) | (d_2353_v12_)) | (_dafny.Map({d_2340_v0_: p0}))
                lhs162_ = d_2354_v13_
                lhs163_ = default__.safeIndex(351, (d_2354_v13_).length(0))
                d_2351_v10_ = rhs291_
                d_2350_v9_ = rhs292_
                d_2343_v2_ = rhs293_
                d_2350_v9_ = rhs294_
                lhs162_[lhs163_] = rhs295_
                d_2364_v21_: _dafny.Seq
                d_2364_v21_ = _dafny.SeqWithoutIsStrInference([((self).f2) - (p3), p1, p1])
                d_2364_v21_ = ((self).f0) + (d_2364_v21_)
                d_2350_v9_ = d_2350_v9_
                d_2351_v10_ = len(((d_2364_v21_) + (d_2364_v21_)) + ((self).f0))
            elif True:
                d_2365_v22_: _dafny.Seq
                d_2365_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqhfxuca"))
                d_2366_v23_: _dafny.Array
                nw377_ = _dafny.Array(None, 4)
                nw377_[int(0)] = p1
                nw377_[int(1)] = (self).f2
                nw377_[int(2)] = (self).f2
                nw377_[int(3)] = len(d_2365_v22_)
                d_2366_v23_ = nw377_
                d_2367_v24_: C5
                nw378_ = C5()
                nw378_.ctor__(d_2366_v23_, _dafny.SeqWithoutIsStrInference([p3 for d_2368_i2_ in range(default__.abs(910))]), self.f1)
                d_2367_v24_ = nw378_
                d_2369_v25_: C5
                d_2369_v25_ = d_2367_v24_
                d_2370_v26_: _dafny.Set
                d_2370_v26_ = _dafny.Set({(d_2369_v25_), d_2367_v24_})
                d_2371_v27_: _dafny.Map
                d_2371_v27_ = _dafny.Map({p0: (d_2346_v5_)[default__.safeIndex((self).f2, len(d_2346_v5_))]})
                d_2372_v28_: _dafny.Map
                d_2372_v28_ = _dafny.Map({(d_2370_v26_) | (d_2370_v26_): len((d_2371_v27_) | (d_2371_v27_))})
                d_2372_v28_ = (d_2372_v28_).set(d_2370_v26_, (self).f2)
                rhs296_ = (p0 if d_2350_v9_ else d_2350_v9_)
                rhs297_ = p0
                d_2350_v9_ = rhs296_
                d_2350_v9_ = rhs297_
                d_2373_v29_: _dafny.Map
                d_2373_v29_ = _dafny.Map({(0) - (((self).f0)[default__.safeIndex(p3, len((self).f0))]): (d_2365_v22_).set(default__.safeIndex((self).f2, len(d_2365_v22_)), d_2343_v2_)})
                d_2374_v30_: D14
                d_2374_v30_ = D14_DC34(d_2373_v29_)
                d_2375_v32_: _dafny.Map
                d_2375_v32_ = _dafny.Map({(self).f2: p0})
                d_2376_v33_: _dafny.Map
                d_2376_v33_ = _dafny.Map({p1: d_2343_v2_})
                d_2377_v35_: _dafny.Set
                d_2377_v35_ = _dafny.Set({p3, p1})
                pat_let_tv71_ = p1
                pat_let_tv72_ = d_2365_v22_
                pat_let_tv73_ = d_2373_v29_
                pat_let_tv74_ = d_2373_v29_
                d_2378_v36_: _dafny.Array
                nw379_ = _dafny.Array(None, 23)
                nw379_[int(0)] = d_2374_v30_
                def iife188_():
                    coll86_ = _dafny.Map()
                    compr_86_: int
                    for compr_86_ in (d_2375_v32_).keys.Elements:
                        d_2379_v31_: int = compr_86_
                        if (d_2379_v31_) in (d_2375_v32_):
                            coll86_[(d_2379_v31_) * (p3)] = d_2365_v22_
                    return _dafny.Map(coll86_)
                nw379_[int(1)] = D14_DC34(iife188_()
)
                def iife189_(_pat_let51_0):
                    def iife190_(d_2380_dt__update__tmp_h0_):
                        def iife191_(_pat_let52_0):
                            def iife192_(d_2381_dt__update_hcf62_h0_):
                                return D14_DC34(d_2381_dt__update_hcf62_h0_)
                            return iife192_(_pat_let52_0)
                        return iife191_(_dafny.Map({pat_let_tv71_: pat_let_tv72_}))
                    return iife190_(_pat_let51_0)
                nw379_[int(2)] = iife189_(D14_DC34(d_2373_v29_))
                nw379_[int(3)] = d_2374_v30_
                nw379_[int(4)] = d_2374_v30_
                nw379_[int(5)] = d_2374_v30_
                nw379_[int(6)] = default__.fm68(p0, d_2365_v22_, globalState)
                nw379_[int(7)] = default__.fm68(not(p0), d_2365_v22_, globalState)
                nw379_[int(8)] = d_2374_v30_
                nw379_[int(9)] = D14_DC34(_dafny.Map({(0) - (len(d_2376_v33_)): d_2365_v22_}))
                def iife193_():
                    coll87_ = _dafny.Map()
                    compr_87_: int
                    for compr_87_ in (d_2377_v35_).Elements:
                        d_2382_v34_: int = compr_87_
                        if (d_2382_v34_) in (d_2377_v35_):
                            coll87_[default__.safeModuloInt(d_2382_v34_, p3)] = _dafny.SeqWithoutIsStrInference([d_2343_v2_ for d_2383_i3_ in range(default__.abs(-583))])
                    return _dafny.Map(coll87_)
                nw379_[int(10)] = D14_DC34(iife193_()
)
                nw379_[int(11)] = d_2374_v30_
                nw379_[int(12)] = d_2374_v30_
                nw379_[int(13)] = default__.fm68(p0, d_2365_v22_, globalState)
                nw379_[int(14)] = d_2374_v30_
                nw379_[int(15)] = d_2374_v30_
                nw379_[int(16)] = d_2374_v30_
                def iife194_(_pat_let53_0):
                    def iife195_(d_2384_dt__update__tmp_h1_):
                        def iife196_(_pat_let54_0):
                            def iife197_(d_2385_dt__update_hcf62_h1_):
                                return D14_DC34(d_2385_dt__update_hcf62_h1_)
                            return iife197_(_pat_let54_0)
                        return iife196_(pat_let_tv73_)
                    return iife195_(_pat_let53_0)
                nw379_[int(17)] = iife194_(d_2374_v30_)
                nw379_[int(18)] = d_2374_v30_
                nw379_[int(19)] = d_2374_v30_
                nw379_[int(20)] = d_2374_v30_
                nw379_[int(21)] = d_2374_v30_
                def iife198_(_pat_let55_0):
                    def iife199_(d_2386_dt__update__tmp_h2_):
                        def iife200_(_pat_let56_0):
                            def iife201_(d_2387_dt__update_hcf62_h2_):
                                return D14_DC34(d_2387_dt__update_hcf62_h2_)
                            return iife201_(_pat_let56_0)
                        return iife200_(pat_let_tv74_)
                    return iife199_(_pat_let55_0)
                nw379_[int(22)] = iife198_(d_2374_v30_)
                d_2378_v36_ = nw379_
                index378_ = default__.safeIndex(751, (d_2378_v36_).length(0))
                (d_2378_v36_)[index378_] = d_2374_v30_
                index379_ = default__.safeIndex(751, (d_2378_v36_).length(0))
                (d_2378_v36_)[index379_] = d_2374_v30_
                d_2375_v32_ = (default__.fm21(globalState)).set(p1, d_2350_v9_)
                d_2350_v9_ = p0
            d_2388_v37_: C3
            nw380_ = C3()
            nw380_.ctor__()
            d_2388_v37_ = nw380_
            d_2389_v38_: _dafny.Array
            nw381_ = _dafny.Array(D2.default()(), 26)
            d_2389_v38_ = nw381_
            d_2390_v39_: _dafny.Seq
            d_2390_v39_ = _dafny.SeqWithoutIsStrInference([d_2343_v2_])
            d_2391_v40_: _dafny.MultiSet
            d_2391_v40_ = _dafny.MultiSet([d_2390_v39_])
            d_2392_v42_: D2
            def iife202_():
                coll88_ = _dafny.Set()
                compr_88_: _dafny.Seq
                for compr_88_ in (d_2391_v40_).Elements:
                    d_2393_v41_: _dafny.Seq = compr_88_
                    if (d_2393_v41_) in (d_2391_v40_):
                        coll88_ = coll88_.union(_dafny.Set([d_2393_v41_]))
                return _dafny.Set(coll88_)
            d_2392_v42_ = D2_DC3(iife202_()
)
            pat_let_tv75_ = d_2390_v39_
            index380_ = default__.safeIndex(580, (d_2389_v38_).length(0))
            def iife203_(_pat_let57_0):
                def iife204_(d_2394_dt__update__tmp_h3_):
                    def iife205_(_pat_let58_0):
                        def iife206_(d_2395_dt__update_hcf4_h0_):
                            return D2_DC3(d_2395_dt__update_hcf4_h0_)
                        return iife206_(_pat_let58_0)
                    return iife205_(_dafny.Set({pat_let_tv75_}))
                return iife204_(_pat_let57_0)
            (d_2389_v38_)[index380_] = iife203_(d_2392_v42_)
            index381_ = default__.safeIndex(580, (d_2389_v38_).length(0))
            (d_2389_v38_)[index381_] = d_2392_v42_
        elif True:
            index382_ = default__.safeIndex(116, (d_2340_v0_).length(0))
            (d_2340_v0_)[index382_] = p0
            d_2396_v43_: int
            d_2396_v43_ = 950
            d_2397_v44_: _dafny.Seq
            d_2397_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvwvfkb"))
            d_2398_v45_: _dafny.Array
            nw382_ = _dafny.Array(int(0), 1)
            d_2398_v45_ = nw382_
            index383_ = default__.safeIndex(315, (d_2398_v45_).length(0))
            (d_2398_v45_)[index383_] = default__.safeModuloInt(len(_dafny.Map({(self).f0: (self).f2})), 232)
            d_2399_v46_: _dafny.Set
            d_2399_v46_ = _dafny.Set({False, p0})
            d_2400_v47_: _dafny.Map
            d_2400_v47_ = _dafny.Map({len(d_2399_v46_): d_2398_v45_})
            index384_ = default__.safeIndex(116, (d_2340_v0_).length(0))
            index385_ = default__.safeIndex(315, (d_2398_v45_).length(0))
            rhs298_ = p0
            rhs299_ = default__.fm20(not(p0), globalState)
            rhs300_ = (d_2397_v44_).set(default__.safeIndex(p3, len(d_2397_v44_)), (_dafny.CodePoint('u') if p0 else _dafny.CodePoint('q')))
            rhs301_ = (default__.safeDivisionInt(default__.fm44(p0, globalState), d_2396_v43_) if ((self).f2) in (d_2400_v47_) else p3)
            lhs164_ = d_2340_v0_
            lhs165_ = default__.safeIndex(116, (d_2340_v0_).length(0))
            lhs166_ = d_2398_v45_
            lhs167_ = default__.safeIndex(315, (d_2398_v45_).length(0))
            lhs164_[lhs165_] = rhs298_
            d_2396_v43_ = rhs299_
            d_2397_v44_ = rhs300_
            lhs166_[lhs167_] = rhs301_
            if (d_2340_v0_)[default__.safeIndex(116, (d_2340_v0_).length(0))]:
                index386_ = default__.safeIndex(315, (d_2398_v45_).length(0))
                (d_2398_v45_)[index386_] = (self).f2
                index387_ = default__.safeIndex(116, (d_2340_v0_).length(0))
                (d_2340_v0_)[index387_] = not((d_2340_v0_)[default__.safeIndex(116, (d_2340_v0_).length(0))])
                index388_ = default__.safeIndex(315, (d_2398_v45_).length(0))
                (d_2398_v45_)[index388_] = (0) - (((self).f0)[default__.safeIndex((self).f2, len((self).f0))])
                d_2401_v48_: T0
                nw383_ = C3()
                nw383_.ctor__()
                d_2401_v48_ = nw383_
                d_2401_v48_ = d_2401_v48_
                d_2397_v44_ = (d_2397_v44_) + ((d_2397_v44_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oo"))))
            elif True:
                index389_ = default__.safeIndex(116, (d_2340_v0_).length(0))
                (d_2340_v0_)[index389_] = ((0) - (default__.safeDivisionInt(len(d_2397_v44_), d_2396_v43_))) == (((self).f2 if True else default__.fm20(default__.fm17(globalState), globalState)))
                d_2402_v49_: _dafny.Map
                d_2402_v49_ = _dafny.Map({d_2396_v43_: _dafny.SeqWithoutIsStrInference([538 for d_2403_i4_ in range(default__.abs(-869))])})
                index390_ = default__.safeIndex(315, (d_2398_v45_).length(0))
                (d_2398_v45_)[index390_] = ((self).f2 if (d_2340_v0_)[default__.safeIndex(116, (d_2340_v0_).length(0))] else default__.safeDivisionInt(len(((d_2402_v49_)[p1] if (p1) in (d_2402_v49_) else (self).f0)), len(_dafny.Set({(self).f2}))))
                d_2404_v50_: _dafny.Map
                d_2404_v50_ = _dafny.Map({len(d_2346_v5_): 930})
                d_2405_v51_: _dafny.MultiSet
                d_2405_v51_ = _dafny.MultiSet([len(d_2404_v50_)])
                d_2406_v52_: _dafny.MultiSet
                d_2406_v52_ = _dafny.MultiSet([(d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))], ((d_2404_v50_)[-461] if (-461) in (d_2404_v50_) else (self).f2), p3, (self).f2, ((d_2405_v51_).cardinality) + (p3)])
                index391_ = default__.safeIndex(315, (d_2398_v45_).length(0))
                (d_2398_v45_)[index391_] = ((d_2405_v51_)[(936) * (p3)] if ((936) * (p3)) in (d_2405_v51_) else default__.fm44(p0, globalState))
                index392_ = default__.safeIndex(315, (d_2398_v45_).length(0))
                (d_2398_v45_)[index392_] = ((self).f0)[default__.safeIndex(404, len((self).f0))]
                d_2407_v53_: C4
                nw384_ = C4()
                nw384_.ctor__()
                d_2407_v53_ = nw384_
                d_2408_v54_: _dafny.Map
                d_2408_v54_ = _dafny.Map({p0: d_2407_v53_})
                d_2408_v54_ = (d_2408_v54_).set((d_2340_v0_)[default__.safeIndex(116, (d_2340_v0_).length(0))], d_2407_v53_)
            d_2409_v55_: _dafny.Map
            d_2409_v55_ = _dafny.Map({p3: d_2343_v2_})
            d_2343_v2_ = ((d_2409_v55_)[p1] if (p1) in (d_2409_v55_) else d_2343_v2_)
            d_2410_v56_: _dafny.Map
            d_2410_v56_ = _dafny.Map({p0: (d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))]})
            d_2411_v57_: _dafny.Map
            d_2411_v57_ = _dafny.Map({303: d_2410_v56_})
            if (len((d_2411_v57_).set(p3, d_2410_v56_))) != (p3):
                d_2412_v58_: C2
                nw385_ = C2()
                nw385_.ctor__()
                d_2412_v58_ = nw385_
                d_2413_v59_: _dafny.Set
                d_2413_v59_ = _dafny.Set({d_2412_v58_})
                d_2414_v60_: _dafny.Array
                nw386_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_2414_v60_ = nw386_
                d_2415_v61_: _dafny.Map
                d_2415_v61_ = _dafny.Map({(d_2413_v59_) | (d_2413_v59_): d_2414_v60_})
                d_2415_v61_ = (d_2415_v61_).set(_dafny.Set({d_2412_v58_}), d_2414_v60_)
                d_2396_v43_ = d_2396_v43_
                d_2416_v62_: _dafny.Seq
                d_2416_v62_ = _dafny.SeqWithoutIsStrInference([(self).f2])
                d_2416_v62_ = (D2_DC5(d_2396_v43_, d_2416_v62_, False)).cf9
                d_2417_v63_: _dafny.Array
                def lambda117_(d_2418_v2_):
                    def lambda118_(d_2419_i5_):
                        return d_2418_v2_

                    return lambda118_

                init61_ = lambda117_(d_2343_v2_)
                nw387_ = _dafny.Array(None, 5)
                for i0_61_ in range(nw387_.length(0)):
                    nw387_[i0_61_] = init61_(i0_61_)
                d_2417_v63_ = nw387_
                d_2420_v64_: _dafny.Seq
                d_2420_v64_ = _dafny.SeqWithoutIsStrInference([d_2417_v63_])
                d_2420_v64_ = d_2420_v64_
            elif True:
                index393_ = default__.safeIndex(315, (d_2398_v45_).length(0))
                (d_2398_v45_)[index393_] = (d_2396_v43_ if p0 else (d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))])
                d_2397_v44_ = d_2397_v44_
                d_2421_v65_: _dafny.Seq
                d_2421_v65_ = _dafny.SeqWithoutIsStrInference([(d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))]])
                d_2421_v65_ = d_2421_v65_
                d_2343_v2_ = (d_2397_v44_)[default__.safeIndex(default__.safeModuloInt(p3, (_dafny.MultiSet([d_2397_v44_])).cardinality), len(d_2397_v44_))]
                d_2422_v66_: _dafny.Array
                nw388_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_2422_v66_ = nw388_
                d_2423_v67_: _dafny.Array
                nw389_ = _dafny.Array(_dafny.Map({}), 29)
                d_2423_v67_ = nw389_
                index394_ = default__.safeIndex(29, (d_2422_v66_).length(0))
                (d_2422_v66_)[index394_] = (d_2423_v67_ if (d_2346_v5_)[default__.safeIndex(p3, len(d_2346_v5_))] else d_2423_v67_)
                d_2424_v68_: _dafny.Set
                d_2424_v68_ = _dafny.Set({len(d_2397_v44_)})
                d_2425_v69_: _dafny.Map
                d_2425_v69_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))]])): (d_2340_v0_)[default__.safeIndex(116, (d_2340_v0_).length(0))]})
                d_2426_v70_: _dafny.Map
                d_2426_v70_ = _dafny.Map({_dafny.Set({len(d_2397_v44_), (d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))], (d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))], (d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))], p1}): d_2425_v69_})
                index395_ = default__.safeIndex(29, (d_2422_v66_).length(0))
                index396_ = default__.safeIndex(116, (d_2340_v0_).length(0))
                rhs302_ = d_2423_v67_
                rhs303_ = (self).f2
                rhs304_ = (d_2424_v68_) not in ((d_2426_v70_) | (d_2426_v70_))
                lhs168_ = d_2422_v66_
                lhs169_ = default__.safeIndex(29, (d_2422_v66_).length(0))
                lhs170_ = d_2340_v0_
                lhs171_ = default__.safeIndex(116, (d_2340_v0_).length(0))
                lhs168_[lhs169_] = rhs302_
                d_2396_v43_ = rhs303_
                lhs170_[lhs171_] = rhs304_
            d_2427_v71_: D13
            d_2427_v71_ = D13_DC32(d_2397_v44_, p0, d_2396_v43_, d_2398_v45_)
            d_2428_v72_: _dafny.MultiSet
            d_2428_v72_ = _dafny.MultiSet([d_2397_v44_, (d_2427_v71_).cf57])
            index397_ = default__.safeIndex(116, (d_2340_v0_).length(0))
            (d_2340_v0_)[index397_] = (self).fm5(p3, p0, d_2344_v3_, ((d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))]) < (((d_2428_v72_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cu"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cu"))) in (d_2428_v72_) else (d_2398_v45_)[default__.safeIndex(315, (d_2398_v45_).length(0))])), globalState)
        d_2429_v73_: _dafny.Array
        nw390_ = _dafny.Array(int(0), 28)
        d_2429_v73_ = nw390_
        index398_ = default__.safeIndex(772, (d_2429_v73_).length(0))
        (d_2429_v73_)[index398_] = p3
        index399_ = default__.safeIndex(772, (d_2429_v73_).length(0))
        (d_2429_v73_)[index399_] = (p1) * ((self).f2)
        index400_ = default__.safeIndex(846, (d_2340_v0_).length(0))
        (d_2340_v0_)[index400_] = p0
        d_2430_v74_: _dafny.Set
        d_2430_v74_ = _dafny.Set({(self).f0, _dafny.SeqWithoutIsStrInference([p3])})
        d_2431_v75_: D15
        d_2431_v75_ = D15_DC37(d_2430_v74_)
        d_2432_v76_: _dafny.Map
        d_2432_v76_ = _dafny.Map({(d_2429_v73_)[default__.safeIndex(772, (d_2429_v73_).length(0))]: d_2431_v75_})
        d_2433_v77_: _dafny.Array
        nw391_ = _dafny.Array(_dafny.CodePoint('D'), 12)
        d_2433_v77_ = nw391_
        d_2434_v78_: D20
        d_2434_v78_ = D20_DC52(d_2433_v77_)
        d_2435_v79_: _dafny.Array
        nw392_ = _dafny.Array(None, 22)
        nw392_[int(0)] = d_2433_v77_
        nw392_[int(1)] = d_2433_v77_
        nw392_[int(2)] = d_2433_v77_
        nw392_[int(3)] = d_2433_v77_
        nw392_[int(4)] = d_2433_v77_
        nw392_[int(5)] = d_2433_v77_
        nw392_[int(6)] = d_2433_v77_
        nw392_[int(7)] = d_2433_v77_
        nw392_[int(8)] = d_2433_v77_
        nw392_[int(9)] = d_2433_v77_
        nw392_[int(10)] = d_2433_v77_
        nw392_[int(11)] = (d_2434_v78_).cf93
        nw392_[int(12)] = d_2433_v77_
        nw392_[int(13)] = d_2433_v77_
        nw392_[int(14)] = d_2433_v77_
        nw392_[int(15)] = d_2433_v77_
        nw392_[int(16)] = d_2433_v77_
        nw392_[int(17)] = (d_2434_v78_).cf93
        nw392_[int(18)] = d_2433_v77_
        nw392_[int(19)] = d_2433_v77_
        nw392_[int(20)] = d_2433_v77_
        nw392_[int(21)] = d_2433_v77_
        d_2435_v79_ = nw392_
        index401_ = default__.safeIndex(419, (d_2435_v79_).length(0))
        (d_2435_v79_)[index401_] = d_2433_v77_
        d_2436_v81_: _dafny.Map
        d_2436_v81_ = _dafny.Map({_dafny.Set({(self).f2, (d_2429_v73_)[default__.safeIndex(772, (d_2429_v73_).length(0))]}): default__.fm69(False, globalState)})
        d_2437_v82_: _dafny.Set
        d_2437_v82_ = _dafny.Set({(0) - ((0) - (p3))})
        d_2438_v83_: _dafny.Map
        d_2438_v83_ = _dafny.Map({p0: (self).f2})
        index402_ = default__.safeIndex(846, (d_2340_v0_).length(0))
        index403_ = default__.safeIndex(419, (d_2435_v79_).length(0))
        index404_ = default__.safeIndex(772, (d_2429_v73_).length(0))
        def iife207_():
            coll89_ = _dafny.Map()
            compr_89_: int
            for compr_89_ in _dafny.IntegerRange(311, 702):
                d_2439_v80_: int = compr_89_
                if ((311) <= (d_2439_v80_)) and ((d_2439_v80_) < (702)):
                    coll89_[(d_2439_v80_) - (-843)] = d_2431_v75_
            return _dafny.Map(coll89_)
        rhs305_ = p0
        rhs306_ = (iife207_()
        ) | (((d_2436_v81_)[d_2437_v82_] if (d_2437_v82_) in (d_2436_v81_) else d_2432_v76_))
        rhs307_ = d_2433_v77_
        rhs308_ = ((self).f0)[default__.safeIndex(((d_2438_v83_)[False] if (False) in (d_2438_v83_) else (self).f2), len((self).f0))]
        lhs172_ = d_2340_v0_
        lhs173_ = default__.safeIndex(846, (d_2340_v0_).length(0))
        lhs174_ = d_2435_v79_
        lhs175_ = default__.safeIndex(419, (d_2435_v79_).length(0))
        lhs176_ = d_2429_v73_
        lhs177_ = default__.safeIndex(772, (d_2429_v73_).length(0))
        lhs172_[lhs173_] = rhs305_
        d_2432_v76_ = rhs306_
        lhs174_[lhs175_] = rhs307_
        lhs176_[lhs177_] = rhs308_
        d_2343_v2_ = _dafny.CodePoint('x')
        d_2440_v84_: C4
        nw393_ = C4()
        nw393_.ctor__()
        d_2440_v84_ = nw393_

    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3

class C9:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        return 63

    def m0(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.Seq({})
        d_2441_v0_: bool
        d_2441_v0_ = False
        d_2442_v1_: _dafny.MultiSet
        d_2442_v1_ = _dafny.MultiSet([d_2441_v0_])
        r0 = (self).fm0(d_2441_v0_, (d_2442_v1_).cardinality, globalState)
        d_2443_v2_: str
        d_2443_v2_ = _dafny.CodePoint('q')
        d_2444_v3_: _dafny.Seq
        d_2444_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), d_2443_v2_])
        d_2445_v4_: _dafny.Seq
        d_2445_v4_ = _dafny.SeqWithoutIsStrInference([d_2444_v3_, d_2444_v3_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2446_i0_ in range(default__.abs(-249))]), (_dafny.SeqWithoutIsStrInference([d_2443_v2_ for d_2447_i1_ in range(default__.abs(398))])) + (d_2444_v3_), _dafny.SeqWithoutIsStrInference([d_2443_v2_])])
        d_2445_v4_ = d_2445_v4_
        d_2448_v5_: _dafny.MultiSet
        d_2448_v5_ = _dafny.MultiSet([d_2443_v2_])
        d_2449_v6_: _dafny.Array
        nw394_ = _dafny.Array(None, 1)
        nw394_[int(0)] = d_2448_v5_
        d_2449_v6_ = nw394_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2449_v6_).length(0)):
            d_2450_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_2450_i2_)) and ((d_2450_i2_) < ((d_2449_v6_).length(0)))):
                (d_2449_v6_)[(d_2450_i2_)] = d_2448_v5_
        d_2451_v7_: _dafny.Set
        d_2451_v7_ = _dafny.Set({d_2441_v0_})
        d_2452_i3_: int
        d_2452_i3_ = 0
        with _dafny.label("12"):
            while (d_2451_v7_).ispropersubset(d_2451_v7_):
                with _dafny.c_label("12"):
                    if (d_2452_i3_) >= (100):
                        raise _dafny.Break("12")
                    d_2452_i3_ = (d_2452_i3_) + (1)
                    d_2453_v8_: _dafny.Map
                    d_2453_v8_ = _dafny.Map({(d_2451_v7_).isdisjoint(d_2451_v7_): p0})
                    d_2454_v9_: _dafny.Seq
                    d_2454_v9_ = _dafny.SeqWithoutIsStrInference([True])
                    d_2455_v10_: _dafny.Array
                    nw395_ = _dafny.Array(None, 13)
                    nw395_[int(0)] = d_2441_v0_
                    nw395_[int(1)] = d_2441_v0_
                    nw395_[int(2)] = False
                    nw395_[int(3)] = d_2441_v0_
                    nw395_[int(4)] = d_2441_v0_
                    nw395_[int(5)] = not(not(d_2441_v0_))
                    nw395_[int(6)] = d_2441_v0_
                    nw395_[int(7)] = d_2441_v0_
                    nw395_[int(8)] = d_2441_v0_
                    nw395_[int(9)] = d_2441_v0_
                    nw395_[int(10)] = (d_2454_v9_)[default__.safeIndex(len(d_2444_v3_), len(d_2454_v9_))]
                    nw395_[int(11)] = d_2441_v0_
                    nw395_[int(12)] = d_2441_v0_
                    d_2455_v10_ = nw395_
                    d_2456_v11_: _dafny.Set
                    d_2456_v11_ = _dafny.Set({d_2455_v10_})
                    d_2453_v8_ = (d_2453_v8_).set(d_2441_v0_, (0) - (len((d_2456_v11_) - (d_2456_v11_))))
                    r1 = (0) - (p0)
                    d_2457_v12_: _dafny.Seq
                    d_2457_v12_ = _dafny.SeqWithoutIsStrInference([(p0) + (p0)])
                    d_2457_v12_ = (_dafny.SeqWithoutIsStrInference([p0, len(d_2444_v3_), p0])).set(default__.safeIndex((len(default__.fm1(p0, d_2441_v0_, p0, d_2441_v0_, globalState))) - ((self).fm0(False, p0, globalState)), len(_dafny.SeqWithoutIsStrInference([p0, len(d_2444_v3_), p0]))), p0)
                    d_2441_v0_ = d_2441_v0_
                    pass
            pass
        d_2458_v13_: _dafny.Map
        d_2458_v13_ = _dafny.Map({p0: d_2441_v0_})
        d_2458_v13_ = d_2458_v13_
        d_2459_v14_: _dafny.Array
        def lambda119_(d_2460_v7_):
            def lambda120_(d_2461_i4_):
                return d_2460_v7_

            return lambda120_

        init62_ = lambda119_(d_2451_v7_)
        nw396_ = _dafny.Array(None, 3)
        for i0_62_ in range(nw396_.length(0)):
            nw396_[i0_62_] = init62_(i0_62_)
        d_2459_v14_ = nw396_
        d_2462_v15_: D0
        d_2462_v15_ = D0_DC0((d_2459_v14_ if d_2441_v0_ else d_2459_v14_), p0)
        source26_ = d_2462_v15_
        if source26_.is_DC0:
            d_2463___mcc_h0_ = source26_.cf0
            d_2464___mcc_h1_ = source26_.cf1
            d_2465_cf1_ = d_2464___mcc_h1_
            d_2466_cf0_ = d_2463___mcc_h0_
            r0 = (0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "moityagb"))), p0))
            d_2467_v16_: _dafny.Seq
            d_2467_v16_ = _dafny.SeqWithoutIsStrInference([d_2441_v0_, d_2441_v0_, True, d_2441_v0_])
            d_2468_v17_: _dafny.Array
            nw397_ = _dafny.Array(None, 9)
            nw397_[int(0)] = not (d_2441_v0_) or (not(not(d_2441_v0_)))
            nw397_[int(1)] = (d_2467_v16_)[default__.safeIndex((self).fm0(d_2441_v0_, p0, globalState), len(d_2467_v16_))]
            nw397_[int(2)] = (d_2467_v16_)[default__.safeIndex(d_2465_cf1_, len(d_2467_v16_))]
            nw397_[int(3)] = (d_2467_v16_) != (d_2467_v16_)
            nw397_[int(4)] = (d_2441_v0_ if d_2441_v0_ else d_2441_v0_)
            nw397_[int(5)] = (p0) >= ((0) - (d_2465_cf1_))
            nw397_[int(6)] = (not(d_2441_v0_) if d_2441_v0_ else d_2441_v0_)
            nw397_[int(7)] = d_2441_v0_
            nw397_[int(8)] = not(d_2441_v0_)
            d_2468_v17_ = nw397_
            index405_ = default__.safeIndex(505, (d_2468_v17_).length(0))
            (d_2468_v17_)[index405_] = d_2441_v0_
            d_2469_v18_: _dafny.Map
            d_2469_v18_ = _dafny.Map({(d_2441_v0_) and (d_2441_v0_): d_2441_v0_})
            index406_ = default__.safeIndex(505, (d_2468_v17_).length(0))
            (d_2468_v17_)[index406_] = not(((d_2469_v18_)[not(d_2441_v0_)] if (not(d_2441_v0_)) in (d_2469_v18_) else d_2441_v0_))
            d_2470_v19_: _dafny.Seq
            d_2470_v19_ = _dafny.SeqWithoutIsStrInference([d_2451_v7_, _dafny.Set({d_2441_v0_}), d_2451_v7_, d_2451_v7_, _dafny.Set({(d_2468_v17_)[default__.safeIndex(505, (d_2468_v17_).length(0))]})])
            d_2471_v20_: _dafny.Set
            d_2471_v20_ = _dafny.Set({p0, d_2465_cf1_, len(d_2470_v19_)})
            index407_ = default__.safeIndex(505, (d_2468_v17_).length(0))
            rhs309_ = (568) + (p0)
            rhs310_ = d_2444_v3_
            rhs311_ = not(d_2441_v0_)
            rhs312_ = (d_2471_v20_).ispropersubset(d_2471_v20_)
            lhs178_ = d_2468_v17_
            lhs179_ = default__.safeIndex(505, (d_2468_v17_).length(0))
            r0 = rhs309_
            d_2444_v3_ = rhs310_
            lhs178_[lhs179_] = rhs311_
            d_2441_v0_ = rhs312_
            d_2472_v21_: _dafny.Array
            nw398_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_2472_v21_ = nw398_
            d_2473_v22_: _dafny.Array
            nw399_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
            d_2473_v22_ = nw399_
            index408_ = default__.safeIndex(401, (d_2472_v21_).length(0))
            (d_2472_v21_)[index408_] = d_2473_v22_
            index409_ = default__.safeIndex(401, (d_2472_v21_).length(0))
            index410_ = default__.safeIndex(505, (d_2468_v17_).length(0))
            rhs313_ = d_2473_v22_
            rhs314_ = (d_2467_v16_)[default__.safeIndex(d_2465_cf1_, len(d_2467_v16_))]
            rhs315_ = (d_2445_v4_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2443_v2_ for d_2474_i5_ in range(default__.abs(363))]), d_2444_v3_]))
            lhs180_ = d_2472_v21_
            lhs181_ = default__.safeIndex(401, (d_2472_v21_).length(0))
            lhs182_ = d_2468_v17_
            lhs183_ = default__.safeIndex(505, (d_2468_v17_).length(0))
            lhs180_[lhs181_] = rhs313_
            lhs182_[lhs183_] = rhs314_
            d_2445_v4_ = rhs315_
        elif True:
            d_2475___mcc_h2_ = source26_.cf2
            d_2476_cf2_ = d_2475___mcc_h2_
            if not(not((default__.fm2((d_2441_v0_), d_2476_cf2_, globalState) if d_2441_v0_ else d_2441_v0_))):
                d_2476_cf2_ = ((0) - (d_2476_cf2_)) + (p0)
                d_2441_v0_ = d_2441_v0_
                d_2477_v23_: _dafny.Seq
                d_2477_v23_ = _dafny.SeqWithoutIsStrInference([d_2441_v0_])
                d_2441_v0_ = default__.fm2((d_2477_v23_)[default__.safeIndex(p0, len(d_2477_v23_))], p0, globalState)
                d_2441_v0_ = default__.fm2(d_2441_v0_, default__.safeModuloInt(p0, 829), globalState)
                d_2478_v24_: _dafny.Map
                d_2478_v24_ = _dafny.Map({d_2477_v23_: d_2443_v2_})
                d_2478_v24_ = default__.fm3(d_2443_v2_, (d_2476_cf2_) * ((0) - (d_2476_cf2_)), d_2441_v0_, d_2441_v0_, globalState)
            elif True:
                d_2479_v25_: _dafny.Array
                nw400_ = _dafny.Array(int(0), 8)
                d_2479_v25_ = nw400_
                d_2480_v26_: _dafny.Map
                d_2480_v26_ = _dafny.Map({d_2441_v0_: d_2476_cf2_})
                index411_ = default__.safeIndex(785, (d_2479_v25_).length(0))
                (d_2479_v25_)[index411_] = ((d_2480_v26_)[True] if (True) in (d_2480_v26_) else len(d_2444_v3_))
                index412_ = default__.safeIndex(785, (d_2479_v25_).length(0))
                (d_2479_v25_)[index412_] = (self).fm0((d_2441_v0_) and (d_2441_v0_), (len(default__.fm4(d_2444_v3_, d_2443_v2_, d_2441_v0_, globalState))) - (p0), globalState)
                d_2444_v3_ = (d_2444_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbslwhc")))
                d_2481_v27_: _dafny.Map
                d_2481_v27_ = _dafny.Map({d_2441_v0_: d_2443_v2_})
                d_2482_v28_: bool
                d_2482_v28_ = False
                d_2481_v27_ = (d_2481_v27_).set(d_2482_v28_, _dafny.CodePoint('o'))
                d_2483_v29_: _dafny.Array
                nw401_ = _dafny.Array(None, 5)
                nw401_[int(0)] = not (d_2441_v0_) or (d_2441_v0_)
                nw401_[int(1)] = True
                nw401_[int(2)] = d_2441_v0_
                nw401_[int(3)] = d_2441_v0_
                nw401_[int(4)] = True
                d_2483_v29_ = nw401_
                index413_ = default__.safeIndex(550, (d_2483_v29_).length(0))
                (d_2483_v29_)[index413_] = False
                d_2484_v30_: D0
                d_2484_v30_ = D0_DC1((0) - ((d_2479_v25_)[default__.safeIndex(785, (d_2479_v25_).length(0))]))
                d_2485_v31_: _dafny.Map
                d_2485_v31_ = _dafny.Map({d_2484_v30_: d_2451_v7_})
                index414_ = default__.safeIndex(958, (d_2483_v29_).length(0))
                (d_2483_v29_)[index414_] = d_2441_v0_
                d_2486_v32_: _dafny.Map
                d_2486_v32_ = _dafny.Map({False: d_2441_v0_})
                index415_ = default__.safeIndex(550, (d_2483_v29_).length(0))
                index416_ = default__.safeIndex(958, (d_2483_v29_).length(0))
                rhs316_ = default__.safeModuloInt((d_2479_v25_)[default__.safeIndex(785, (d_2479_v25_).length(0))], p0)
                rhs317_ = not(d_2441_v0_)
                rhs318_ = default__.safeModuloInt(d_2476_cf2_, (d_2479_v25_)[default__.safeIndex(785, (d_2479_v25_).length(0))])
                rhs319_ = (d_2485_v31_ if (d_2441_v0_) in (d_2486_v32_) else (d_2485_v31_).set(d_2484_v30_, d_2451_v7_))
                rhs320_ = True
                lhs184_ = d_2483_v29_
                lhs185_ = default__.safeIndex(550, (d_2483_v29_).length(0))
                lhs186_ = d_2483_v29_
                lhs187_ = default__.safeIndex(958, (d_2483_v29_).length(0))
                r0 = rhs316_
                lhs184_[lhs185_] = rhs317_
                r0 = rhs318_
                d_2485_v31_ = rhs319_
                lhs186_[lhs187_] = rhs320_
                index417_ = default__.safeIndex(550, (d_2483_v29_).length(0))
                (d_2483_v29_)[index417_] = ((d_2479_v25_)[default__.safeIndex(785, (d_2479_v25_).length(0))]) <= (470)
            d_2487_v33_: C3
            nw402_ = C3()
            nw402_.ctor__()
            d_2487_v33_ = nw402_
            d_2476_cf2_ = (p0) * (p0)
            d_2441_v0_ = not((d_2441_v0_) == (not((_dafny.Set({True})).issubset(d_2451_v7_))))
        r0 = p0
        d_2488_v34_: T0
        nw403_ = C3()
        nw403_.ctor__()
        d_2488_v34_ = nw403_
        d_2489_v35_: _dafny.Seq
        d_2489_v35_ = _dafny.SeqWithoutIsStrInference([d_2441_v0_, d_2441_v0_, d_2441_v0_, d_2441_v0_])
        d_2490_v36_: _dafny.Map
        d_2490_v36_ = _dafny.Map({d_2488_v34_: d_2489_v35_})
        r1 = len(((d_2490_v36_)[d_2488_v34_] if (d_2488_v34_) in (d_2490_v36_) else d_2489_v35_))
        r2 = (d_2489_v35_) + (d_2489_v35_)
        return r0, r1, r2

