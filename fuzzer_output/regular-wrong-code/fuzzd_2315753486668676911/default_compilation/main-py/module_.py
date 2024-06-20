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
    def fm5(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_0_i0_ in range(default__.abs(643))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svlade")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ls")))

    @staticmethod
    def fm11(p0, globalState):
        return ((_dafny.Map({True: -106})) | (_dafny.Map({True: 767}))) | ((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eoqmvc")))})))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return D3_DC9((_dafny.CodePoint('v') if True else _dafny.CodePoint('v')))

    @staticmethod
    def fm17(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(492, -555):
                d_1_v0_: int = compr_0_
                if ((492) <= (d_1_v0_)) and ((d_1_v0_) < (-555)):
                    def iife1_():
                        coll1_ = _dafny.Map()
                        compr_1_: int
                        for compr_1_ in _dafny.IntegerRange(581, 375):
                            d_2_v1_: int = compr_1_
                            if ((581) <= (d_2_v1_)) and ((d_2_v1_) < (375)):
                                coll1_[default__.safeModuloInt(d_2_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfftcgcr"))))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_3_i0_ in range(default__.abs(540))]))
                        return _dafny.Map(coll1_)
                    coll0_ = coll0_.union(_dafny.Set([default__.safeDivisionInt(d_1_v0_, len(iife1_()
))]))
            return _dafny.Set(coll0_)
        return ((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False])))).cardinality) + ((len(iife0_()
        )) * (431))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in ((_dafny.SeqWithoutIsStrInference([316 for d_4_i0_ in range(default__.abs(-246))])) + (_dafny.SeqWithoutIsStrInference([-88 for d_5_i1_ in range(default__.abs(-167))]))).Elements:
                d_6_v0_: int = compr_2_
                if (d_6_v0_) in ((_dafny.SeqWithoutIsStrInference([316 for d_4_i0_ in range(default__.abs(-246))])) + (_dafny.SeqWithoutIsStrInference([-88 for d_5_i1_ in range(default__.abs(-167))]))):
                    coll2_[default__.safeModuloInt(d_6_v0_, len(_dafny.Set({453, 243})))] = not(True)
            return _dafny.Map(coll2_)
        return iife2_()
        

    @staticmethod
    def fm19(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hafht"))

    @staticmethod
    def fm21(p0, p1, globalState):
        return D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scdrbxig")), False)

    @staticmethod
    def fm22(p0, p1, globalState):
        return _dafny.Map({not((True if not(not(not(False))) else not(True))): not(not(False))})

    @staticmethod
    def fm26(globalState):
        return True

    @staticmethod
    def fm27(globalState):
        if True:
            return _dafny.SeqWithoutIsStrInference([False])
        elif False:
            return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            return (D14_DC36((D14_DC36(_dafny.SeqWithoutIsStrInference([True, False]))).cf53)).cf53

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-786, 405):
                    d_7_v1_: int = compr_5_
                    if ((-786) <= (d_7_v1_)) and ((d_7_v1_) < (405)):
                        coll5_[(d_7_v1_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_8_i0_ in range(default__.abs(628))])))] = False
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(-786, 405):
                    d_7_v1_: int = compr_4_
                    if ((-786) <= (d_7_v1_)) and ((d_7_v1_) < (405)):
                        coll4_[(d_7_v1_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_8_i0_ in range(default__.abs(628))])))] = False
                return _dafny.Map(coll4_)
            compr_3_: int
            for compr_3_ in (iife4_()
            ).keys.Elements:
                d_9_v0_: int = compr_3_
                if (d_9_v0_) in (iife5_()
                ):
                    coll3_[(d_9_v0_) - (696)] = False
            return _dafny.Map(coll3_)
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(227, 175):
                d_10_v2_: int = compr_6_
                if ((227) <= (d_10_v2_)) and ((d_10_v2_) < (175)):
                    coll6_[default__.safeModuloInt(d_10_v2_, len(_dafny.SeqWithoutIsStrInference([-493])))] = False
            return _dafny.Map(coll6_)
        source0_ = D14_DC35(_dafny.SeqWithoutIsStrInference([iife3_()
, iife6_()
]))
        if source0_.is_DC36:
            d_11___mcc_h0_ = source0_.cf53
            d_12_cf53_ = d_11___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsqaxwqte"))
        elif source0_.is_DC35:
            d_13___mcc_h1_ = source0_.cf52
            d_14_cf52_ = d_13___mcc_h1_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxhjglaic"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_15_i1_ in range(default__.abs(-301))]))
        elif True:
            d_16___mcc_h2_ = source0_.cf54
            d_17_cf54_ = d_16___mcc_h2_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mplnkdwif"))

    @staticmethod
    def fm31(p0, p1, globalState):
        return (_dafny.MultiSet([72])) | (_dafny.MultiSet([77, (_dafny.MultiSet([not((D20_DC55(len(_dafny.SeqWithoutIsStrInference([186, len(_dafny.Set({False, False}))])), len((D15_DC38(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_18_i0_ in range(default__.abs(116))])): True}))).cf55), False, True)).cf87), True, True])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmniuvfes"))), 728, 310]))

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(23, 808):
                d_19_v0_: int = compr_7_
                if ((23) <= (d_19_v0_)) and ((d_19_v0_) < (808)):
                    coll7_ = coll7_.union(_dafny.Set([(d_19_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwgcuel"))))]))
            return _dafny.Set(coll7_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(-804, 56):
                d_20_v1_: int = compr_8_
                if ((-804) <= (d_20_v1_)) and ((d_20_v1_) < (56)):
                    coll8_[(d_20_v1_) * (867)] = False
            return _dafny.Map(coll8_)
        return (iife7_()
        ) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(iife8_()
        ), -849, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tm"))), 84, 249])), 843, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')]))}))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return _dafny.Map({(False) or (not(not(False))): False})

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([True, False])).intersection(_dafny.MultiSet([False, True]))).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True]))))

    @staticmethod
    def fm35(globalState):
        return _dafny.Map({(113) + (-13): (False) in (_dafny.Map({True: -78}))})

    @staticmethod
    def fm36(p0, globalState):
        return ((_dafny.Map({False: _dafny.Set({False})})) | (_dafny.Map({True: _dafny.Set({False})}))) | ((_dafny.Map({True: _dafny.Set({True, False})})) | (_dafny.Map({True: _dafny.Set({False, True})})))

    @staticmethod
    def fm37(globalState):
        return ((_dafny.Set({False, False})) - (_dafny.Set({False}))) | ((_dafny.Set({True})) | (_dafny.Set({False, True})))

    @staticmethod
    def fm38(p0, globalState):
        source1_ = D12_DC33(True)
        if source1_.is_DC33:
            d_21___mcc_h0_ = source1_.cf50
            d_22_cf50_ = d_21___mcc_h0_
            return _dafny.CodePoint('s')
        elif True:
            d_23___mcc_h1_ = source1_.cf49
            d_24_cf49_ = d_23___mcc_h1_
            if True:
                return _dafny.CodePoint('w')
            elif True:
                return _dafny.CodePoint('t')

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([669, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([930]), _dafny.MultiSet([897, 342, (0) - ((_dafny.MultiSet([False])).cardinality), len(_dafny.SeqWithoutIsStrInference([162 for d_25_i0_ in range(default__.abs(503))]))])]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({True: 732})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnqkq")))}))]))) + ((_dafny.SeqWithoutIsStrInference([157, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rysmqpp")))]) if False else _dafny.SeqWithoutIsStrInference([(0) - (-412), len(_dafny.Map({True: False})), 193, len(_dafny.Map({True: True})), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ng"))): False})): _dafny.CodePoint('x')}))])), -685]))])))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return D1_DC2((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clskkt"))) if True else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_26_i0_ in range(default__.abs(681))]))))

    @staticmethod
    def fm41(p0, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxpjsdcn"))))])).Elements:
                d_28_v0_: int = compr_9_
                if (d_28_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxpjsdcn"))))])):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_28_v0_, (_dafny.MultiSet([854])).cardinality)]))
            return _dafny.Set(coll9_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjghubcnr"))): len(_dafny.SeqWithoutIsStrInference([False, not((D19_DC52(272, True, True)).cf82)]))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_27_i0_ in range(default__.abs(-738))])): len(iife9_()
        )}))) | ((_dafny.Map({448: len(_dafny.Map({False: False}))})) | (_dafny.Map({-717: 806})))

    @staticmethod
    def fm42(p0, p1, globalState):
        if (not(True)) == (not(True)):
            return D1_DC4(_dafny.MultiSet([-88]))
        elif True:
            return D1_DC4(_dafny.MultiSet([len(_dafny.Map({not(True): 819}))]))

    @staticmethod
    def fm43(p0, globalState):
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(858, 165):
                d_29_v0_: int = compr_10_
                if ((858) <= (d_29_v0_)) and ((d_29_v0_) < (165)):
                    coll10_ = coll10_.union(_dafny.Set([(d_29_v0_) * (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality])))]))
            return _dafny.Set(coll10_)
        return (_dafny.MultiSet([iife10_()
        , _dafny.Set({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ade")))]))}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "euwroheny")))})])) - ((_dafny.MultiSet([_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icwgdatuh"))), 632})])).intersection(_dafny.MultiSet([_dafny.Set({len(_dafny.Map({-518: False}))}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfnisofy")))}), _dafny.Set({-276})])))

    @staticmethod
    def fm44(p0, p1, globalState):
        return _dafny.Map({((D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frlkbglny")), not(False))).cf3) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_30_i0_ in range(default__.abs(-612))])): 698})

    @staticmethod
    def fm45(globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(-567, 334):
                d_31_v0_: int = compr_11_
                if ((-567) <= (d_31_v0_)) and ((d_31_v0_) < (334)):
                    coll11_[(d_31_v0_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = (D12_DC33(True)).cf50
            return _dafny.Map(coll11_)
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False])]))) for d_32_i0_ in range(default__.abs(393))])).Elements:
                d_33_v1_: int = compr_12_
                if (d_33_v1_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False])]))) for d_32_i0_ in range(default__.abs(393))])):
                    coll12_[default__.safeModuloInt(d_33_v1_, (0) - (len(_dafny.Set({False}))))] = not(not(True))
            return _dafny.Map(coll12_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klqhbvu"))): False})])) + (_dafny.SeqWithoutIsStrInference([iife11_()
        , iife12_()
        ]))

    @staticmethod
    def fm46(p0, globalState):
        return D16_DC42(not(False), True, (True) == (True))

    @staticmethod
    def fm47(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_34_i0_ in range(default__.abs(29))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_35_i1_ in range(default__.abs(168))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pabhldsjv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cksrrod")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwrut"))])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddnqnu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nb"))]))).intersection((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itm")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_36_i2_ in range(default__.abs(177))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juvss"))]))).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_37_i3_ in range(default__.abs(669))])])))

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        return D14_DC35(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False}), _dafny.Map({len(_dafny.Set({582, len(_dafny.SeqWithoutIsStrInference([True, True, not(False)]))})): not(False)}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([343])): False})]))

    @staticmethod
    def fm49(p0, p1, globalState):
        return D14_DC36((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([False, False])))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(-474, -262):
                d_40_v0_: int = compr_13_
                if ((-474) <= (d_40_v0_)) and ((d_40_v0_) < (-262)):
                    coll13_[(d_40_v0_) * (357)] = True
            return _dafny.Map(coll13_)
        return ((_dafny.MultiSet([_dafny.MultiSet([D2_DC6(-174)])])) | (_dafny.MultiSet([_dafny.MultiSet([D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udq")))), D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivkmw"))))])]))).intersection(_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D2_DC6(214), D2_DC6(533)])), _dafny.MultiSet([D2_DC6(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_38_i0_ in range(default__.abs(654))])), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxvtbcyv"))): True})), len(_dafny.SeqWithoutIsStrInference([True]))]))), D2_DC6(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_39_i1_ in range(default__.abs(286))])))]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D2_DC6((_dafny.MultiSet([False])).cardinality)])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D2_DC6((_dafny.MultiSet([D16_DC42(False, False, True), D16_DC42(not(True), not(False), False), D16_DC42(not(False), False, False)])).cardinality), D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "resvq")))), D2_DC6((_dafny.MultiSet([True])).cardinality), D2_DC6(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhcaqip")))]))), D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "og"))))])), _dafny.MultiSet([D2_DC6(473), D2_DC6(len(iife13_()
))])]))

    @staticmethod
    def fm51(p0, p1, globalState):
        return D11_DC31(D11_DC31(D11_DC30(True, 684)))

    @staticmethod
    def fm52(globalState):
        def iife14_():
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(334, 762):
                    d_41_v0_: int = compr_16_
                    if ((334) <= (d_41_v0_)) and ((d_41_v0_) < (762)):
                        coll16_[default__.safeDivisionInt(d_41_v0_, 610)] = len(_dafny.Set({not(True)}))
                return _dafny.Map(coll16_)
            coll14_ = _dafny.Set()
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(334, 762):
                    d_41_v0_: int = compr_15_
                    if ((334) <= (d_41_v0_)) and ((d_41_v0_) < (762)):
                        coll15_[default__.safeDivisionInt(d_41_v0_, 610)] = len(_dafny.Set({not(True)}))
                return _dafny.Map(coll15_)
            compr_14_: int
            for compr_14_ in (_dafny.SeqWithoutIsStrInference([len(iife15_()
            )])).Elements:
                d_42_v1_: int = compr_14_
                if (d_42_v1_) in (_dafny.SeqWithoutIsStrInference([len(iife16_()
                )])):
                    coll14_ = coll14_.union(_dafny.Set([(d_42_v1_) * (len(_dafny.SeqWithoutIsStrInference([True, False, False])))]))
            return _dafny.Set(coll14_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in (_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fviyvd"))): _dafny.CodePoint('h')})), -349])).cardinality: False})).keys.Elements:
                d_44_v2_: int = compr_17_
                if (d_44_v2_) in (_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fviyvd"))): _dafny.CodePoint('h')})), -349])).cardinality: False})):
                    coll17_[(d_44_v2_) + (461)] = _dafny.MultiSet([True])
            return _dafny.Map(coll17_)
        return ((_dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC28(iife14_()
)), D11_DC31(D11_DC29())]) if True else _dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC29()), D11_DC31(D11_DC29()), D11_DC31(D11_DC28(_dafny.Set({-644})))]))) + ((_dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC28(_dafny.Set({695}))) for d_43_i0_ in range(default__.abs(543))])) + (_dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC29()), D11_DC31(D11_DC31(D11_DC30(True, len(iife17_()
))))])))

    @staticmethod
    def fm53(p0, globalState):
        return D3_DC10(-282, default__.safeDivisionInt(-116, len(_dafny.Map({not(True): False}))))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        source2_ = D12_DC32(_dafny.Set({False, False, False}))
        if source2_.is_DC33:
            d_45___mcc_h0_ = source2_.cf50
            d_46_cf50_ = d_45___mcc_h0_
            return D19_DC53()
        elif True:
            d_47___mcc_h1_ = source2_.cf49
            d_48_cf49_ = d_47___mcc_h1_
            return D19_DC53()

    @staticmethod
    def fm55(p0, p1, p2, p3, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: str
            for compr_18_ in (_dafny.Map({_dafny.CodePoint('w'): _dafny.CodePoint('t')})).keys.Elements:
                d_49_v0_: str = compr_18_
                if (d_49_v0_) in (_dafny.Map({_dafny.CodePoint('w'): _dafny.CodePoint('t')})):
                    coll18_[d_49_v0_] = False
            return _dafny.Map(coll18_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.Map({_dafny.CodePoint('w'): not(True)})) | (_dafny.Map({_dafny.CodePoint('c'): True})), (iife18_()
        ) | (_dafny.Map({_dafny.CodePoint('t'): True}))])

    @staticmethod
    def fm56(p0, p1, globalState):
        return D7_DC22((_dafny.Set({not(False), False})).issubset(_dafny.Set({not(False)})), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwvgtdwty"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))), (_dafny.MultiSet([False, True, True, True, True])).ispropersubset(_dafny.MultiSet([False, True, True])), (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): True})) for d_50_i0_ in range(default__.abs(498))]))) * (-797))

    @staticmethod
    def fm57(p0, p1, globalState):
        return ((_dafny.MultiSet([D7_DC22(False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, False, 483), D7_DC22(True, -171, False, 747)])) - (_dafny.MultiSet([D7_DC22(False, 640, False, len(_dafny.Map({219: -399}))), D7_DC22(not(not(True)), (0) - ((0) - (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([True]))): _dafny.CodePoint('t')})))), False, len(_dafny.Set({False, True, True})))]))).intersection(_dafny.MultiSet([D7_DC22(not(True), 898, False, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_51_i0_ in range(default__.abs(-133))])))]))

    @staticmethod
    def fm58(p0, p1, globalState):
        return (_dafny.Map({-157: D21_DC57(428, 475, D5_DC14(_dafny.CodePoint('p'), True, len(_dafny.Set({True})), (D18_DC49(True, not(False), _dafny.CodePoint('o'), False)).cf76))})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irgs"))): D21_DC57((_dafny.MultiSet([(_dafny.MultiSet([315])).cardinality, 639, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_52_i0_ in range(default__.abs(202))]))), 680, 30])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yptolnu"))), D5_DC14(_dafny.CodePoint('o'), False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tusd"))), False))})) | (_dafny.Map({490: D21_DC57(814, -888, D5_DC14(_dafny.CodePoint('i'), not(True), 374, True))})))

    @staticmethod
    def fm59(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([(0) - ((0) - (-493)), 469])).cardinality: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xamnu")))})])

    @staticmethod
    def fm60(p0, p1, globalState):
        return D19_DC52(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))), 994), (_dafny.MultiSet([False])).issubset(_dafny.MultiSet([False])), not(not(False)))

    @staticmethod
    def Main(noArgsParameter__):
        d_53_v0_: bool
        d_53_v0_ = False
        d_54_v1_: int
        d_54_v1_ = 215
        d_55_v2_: _dafny.Map
        d_55_v2_ = _dafny.Map({not(d_53_v0_): d_54_v1_})
        d_56_v3_: _dafny.Set
        d_56_v3_ = _dafny.Set({d_54_v1_})
        d_57_v4_: _dafny.Map
        d_57_v4_ = _dafny.Map({not(True): d_53_v0_})
        d_58_v5_: _dafny.Seq
        d_58_v5_ = _dafny.SeqWithoutIsStrInference([d_53_v0_])
        d_59_v6_: _dafny.MultiSet
        d_59_v6_ = _dafny.MultiSet([d_58_v5_])
        d_60_globalState_: GlobalState
        nw0_ = GlobalState()
        nw0_.ctor__((_dafny.SeqWithoutIsStrInference([d_53_v0_, d_53_v0_])) + (_dafny.SeqWithoutIsStrInference([d_53_v0_])), True, _dafny.CodePoint('t'), 51, 503, True, d_55_v2_, 231, d_56_v3_, True, True, (d_57_v4_) | (_dafny.Map({not(d_53_v0_): d_53_v0_})), 38, 113, True, 964, d_59_v6_, True, 454, True)
        d_60_globalState_ = nw0_
        d_61_v7_: C4
        nw1_ = C4()
        nw1_.ctor__()
        d_61_v7_ = nw1_
        hi0_ = d_54_v1_
        for d_62_i0_ in range(d_54_v1_, hi0_):
            d_63_v8_: _dafny.Array
            def lambda0_(d_64_v0_):
                def lambda1_(d_65_i1_):
                    return d_64_v0_

                return lambda1_

            init0_ = lambda0_(d_53_v0_)
            nw2_ = _dafny.Array(None, 23)
            for i0_0_ in range(nw2_.length(0)):
                nw2_[i0_0_] = init0_(i0_0_)
            d_63_v8_ = nw2_
            index0_ = default__.safeIndex(805, (d_63_v8_).length(0))
            (d_63_v8_)[index0_] = not((True if d_53_v0_ else d_53_v0_))
            d_66_v9_: _dafny.Seq
            d_66_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eabrtymy"))
            index1_ = default__.safeIndex(805, (d_63_v8_).length(0))
            rhs0_ = not (True) or (d_53_v0_)
            rhs1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_67_i2_ in range(default__.abs(588))])
            rhs2_ = d_54_v1_
            rhs3_ = ((((_dafny.MultiSet(d_58_v5_)).set(d_53_v0_, default__.abs(d_62_i0_))).cardinality) == (default__.fm17(d_54_v1_, d_60_globalState_)) if d_53_v0_ else d_53_v0_)
            rhs4_ = d_62_i0_
            lhs0_ = d_63_v8_
            lhs1_ = default__.safeIndex(805, (d_63_v8_).length(0))
            lhs0_[lhs1_] = rhs0_
            d_66_v9_ = rhs1_
            d_54_v1_ = rhs2_
            d_53_v0_ = rhs3_
            d_54_v1_ = rhs4_
            d_58_v5_ = d_58_v5_
            d_68_v10_: _dafny.Seq
            d_68_v10_ = _dafny.SeqWithoutIsStrInference([d_66_v9_, d_66_v9_, d_66_v9_, d_66_v9_])
            d_68_v10_ = d_68_v10_
            if True:
                d_69_v11_: _dafny.Map
                d_69_v11_ = _dafny.Map({d_53_v0_: D4_DC11(d_68_v10_)})
                d_70_v12_: _dafny.Map
                d_70_v12_ = _dafny.Map({len(d_57_v4_): d_63_v8_})
                d_71_v13_: _dafny.Set
                d_71_v13_ = _dafny.Set({((d_70_v12_)[898] if (898) in (d_70_v12_) else d_63_v8_)})
                d_72_v14_: D1
                d_72_v14_ = D1_DC2(d_62_i0_)
                index2_ = default__.safeIndex(805, (d_63_v8_).length(0))
                rhs5_ = ((0) - (default__.safeDivisionInt(918, len(_dafny.Set({d_54_v1_, (0) - (d_62_i0_), d_62_i0_}))))) + (d_54_v1_)
                rhs6_ = (((d_61_v7_).fm23(d_69_v11_, len(d_66_v9_), len(d_71_v13_), d_60_globalState_)) * (len(d_66_v9_))) <= (len((d_56_v3_).intersection(_dafny.Set({(d_72_v14_).cf2}))))
                lhs2_ = d_60_globalState_
                lhs3_ = d_63_v8_
                lhs4_ = default__.safeIndex(805, (d_63_v8_).length(0))
                lhs2_.f15 = rhs5_
                lhs3_[lhs4_] = rhs6_
                (d_60_globalState_).f4 = 642
                d_73_v15_: D0
                d_73_v15_ = D0_DC1(d_63_v8_)
                d_74_v16_: _dafny.Map
                d_74_v16_ = _dafny.Map({(d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))]: d_63_v8_})
                d_73_v15_ = D0_DC1(((d_74_v16_)[(d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))]] if ((d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))]) in (d_74_v16_) else d_63_v8_))
                index3_ = default__.safeIndex(805, (d_63_v8_).length(0))
                rhs7_ = (d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))]
                rhs8_ = ((0) - (d_54_v1_) if d_53_v0_ else d_54_v1_)
                lhs5_ = d_63_v8_
                lhs6_ = default__.safeIndex(805, (d_63_v8_).length(0))
                lhs5_[lhs6_] = rhs7_
                d_54_v1_ = rhs8_
                d_75_v17_: _dafny.MultiSet
                d_75_v17_ = _dafny.MultiSet([d_53_v0_, (d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))]])
                d_76_v18_: _dafny.Array
                nw3_ = _dafny.Array(int(0), 11)
                d_76_v18_ = nw3_
                d_77_v19_: D2
                d_77_v19_ = D2_DC7(False, d_75_v17_, d_76_v18_)
                d_78_v20_: D2
                d_78_v20_ = D2_DC8(d_77_v19_)
                d_79_v21_: C3
                nw4_ = C3()
                nw4_.ctor__(d_78_v20_, (d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))])
                d_79_v21_ = nw4_
                d_80_v22_: _dafny.Map
                d_80_v22_ = _dafny.Map({d_79_v21_: (d_62_i0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojorbrkgu"))))})
                d_80_v22_ = ((d_80_v22_).set(d_79_v21_, d_62_i0_) if d_79_v21_.f25 else (d_80_v22_) | (d_80_v22_))
            elif True:
                d_81_v23_: _dafny.MultiSet
                d_81_v23_ = _dafny.MultiSet([d_62_i0_, d_62_i0_, (86) * (d_54_v1_)])
                d_81_v23_ = d_81_v23_
                d_82_v24_: str
                d_82_v24_ = _dafny.CodePoint('w')
                d_83_v25_: _dafny.Map
                d_83_v25_ = _dafny.Map({-880: d_82_v24_})
                d_83_v25_ = (d_83_v25_).set(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_54_v1_ for d_84_i3_ in range(default__.abs(953))])), d_62_i0_), (d_82_v24_ if not((d_63_v8_)[default__.safeIndex(805, (d_63_v8_).length(0))]) else d_82_v24_))
                (d_60_globalState_).f15 = d_54_v1_
                d_85_v26_: C8
                nw5_ = C8()
                nw5_.ctor__()
                d_85_v26_ = nw5_
                d_86_v27_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_86_v27_ = nw6_
                d_86_v27_ = d_86_v27_
        d_87_v28_: C6
        nw7_ = C6()
        nw7_.ctor__()
        d_87_v28_ = nw7_
        d_88_v29_: _dafny.MultiSet
        d_88_v29_ = _dafny.MultiSet([d_87_v28_])
        d_89_v30_: D20
        d_89_v30_ = D20_DC55(d_54_v1_, d_54_v1_, d_53_v0_, (_dafny.MultiSet([d_87_v28_])) != (d_88_v29_))
        source3_ = d_89_v30_
        if source3_.is_DC55:
            d_90___mcc_h0_ = source3_.cf85
            d_91___mcc_h1_ = source3_.cf86
            d_92___mcc_h2_ = source3_.cf87
            d_93___mcc_h3_ = source3_.cf88
            d_94_cf88_ = d_93___mcc_h3_
            d_95_cf87_ = d_92___mcc_h2_
            d_96_cf86_ = d_91___mcc_h1_
            d_97_cf85_ = d_90___mcc_h0_
            d_98_v31_: _dafny.Seq
            d_98_v31_ = _dafny.SeqWithoutIsStrInference([d_96_cf86_])
            d_99_v32_: _dafny.Seq
            d_99_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnextlf"))
            (d_60_globalState_).f15 = (d_54_v1_) - ((d_98_v31_)[default__.safeIndex(len(d_99_v32_), len(d_98_v31_))])
            d_100_v33_: D16
            d_100_v33_ = D16_DC42(False, d_95_cf87_, d_94_cf88_)
            d_101_v34_: _dafny.Seq
            d_101_v34_ = _dafny.SeqWithoutIsStrInference([d_100_v33_, d_100_v33_, d_100_v33_])
            rhs9_ = (_dafny.SeqWithoutIsStrInference([d_100_v33_ for d_102_i4_ in range(default__.abs(740))])) + (d_101_v34_)
            rhs10_ = (d_98_v31_)[default__.safeIndex(d_96_cf86_, len(d_98_v31_))]
            rhs11_ = (d_97_cf85_) + (d_96_cf86_)
            rhs12_ = not (d_94_cf88_) or (d_94_cf88_)
            d_101_v34_ = rhs9_
            d_97_cf85_ = rhs10_
            d_96_cf86_ = rhs11_
            d_53_v0_ = rhs12_
            d_103_v35_: C8
            nw8_ = C8()
            nw8_.ctor__()
            d_103_v35_ = nw8_
            (d_60_globalState_).f7 = d_97_cf85_
        elif True:
            d_104___mcc_h4_ = source3_.cf84
            d_105_cf84_ = d_104___mcc_h4_
            d_106_v36_: int
            d_107_v37_: _dafny.MultiSet
            d_108_v38_: bool
            out0_: int
            out1_: _dafny.MultiSet
            out2_: bool
            out0_, out1_, out2_ = (d_87_v28_).m9(d_54_v1_, d_53_v0_, not(not(d_53_v0_)), d_60_globalState_)
            d_106_v36_ = out0_
            d_107_v37_ = out1_
            d_108_v38_ = out2_
            d_108_v38_ = d_53_v0_
            d_53_v0_ = d_53_v0_
            d_54_v1_ = (0) - ((len(d_58_v5_)) - ((d_54_v1_ if False else d_106_v36_)))
        d_109_v39_: D11
        d_109_v39_ = D11_DC29()
        d_110_v40_: D11
        d_110_v40_ = D11_DC31(d_109_v39_)
        d_111_v41_: D11
        d_111_v41_ = D11_DC31(d_109_v39_)
        d_112_v42_: _dafny.Seq
        d_112_v42_ = _dafny.SeqWithoutIsStrInference([d_111_v41_, d_111_v41_])
        d_113_v43_: D19
        d_113_v43_ = D19_DC51(d_112_v42_)
        d_113_v43_ = d_113_v43_
        d_114_v44_: _dafny.Seq
        d_114_v44_ = _dafny.SeqWithoutIsStrInference([((d_55_v2_)[d_53_v0_] if (d_53_v0_) in (d_55_v2_) else d_54_v1_)])
        (d_60_globalState_).f15 = len(d_114_v44_)
        (d_60_globalState_).f15 = (253) + (d_54_v1_)
        d_115_v45_: _dafny.Array
        nw9_ = _dafny.Array(_dafny.Map({}), 6)
        d_115_v45_ = nw9_
        d_115_v45_ = d_115_v45_
        d_53_v0_ = not(d_53_v0_)
        d_116_v46_: _dafny.MultiSet
        d_116_v46_ = _dafny.MultiSet([d_56_v3_, d_56_v3_])
        d_117_v47_: D5
        d_117_v47_ = D5_DC15(D5_DC13(d_116_v46_))
        d_118_v48_: _dafny.Seq
        d_118_v48_ = _dafny.SeqWithoutIsStrInference([d_117_v47_, d_117_v47_, d_117_v47_])
        d_118_v48_ = d_118_v48_
        d_53_v0_ = d_53_v0_
        d_119_v49_: _dafny.Array
        nw10_ = _dafny.Array(None, 10)
        nw10_[int(0)] = d_58_v5_
        nw10_[int(1)] = _dafny.SeqWithoutIsStrInference([d_53_v0_])
        nw10_[int(2)] = d_58_v5_
        nw10_[int(3)] = _dafny.SeqWithoutIsStrInference([d_53_v0_])
        nw10_[int(4)] = d_58_v5_
        nw10_[int(5)] = d_58_v5_
        nw10_[int(6)] = d_58_v5_
        nw10_[int(7)] = d_58_v5_
        nw10_[int(8)] = (d_58_v5_) + (d_58_v5_)
        nw10_[int(9)] = (d_58_v5_).set(default__.safeIndex(d_54_v1_, len(d_58_v5_)), d_53_v0_)
        d_119_v49_ = nw10_
        d_120_v50_: _dafny.Seq
        d_120_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipyktu"))
        d_121_v51_: C12
        nw11_ = C12()
        nw11_.ctor__(len(d_120_v50_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xofwgsc")))
        d_121_v51_ = nw11_
        d_122_v52_: _dafny.Array
        nw12_ = _dafny.Array(None, 4)
        nw12_[int(0)] = d_121_v51_
        nw12_[int(1)] = d_121_v51_
        nw12_[int(2)] = d_121_v51_
        nw12_[int(3)] = d_121_v51_
        d_122_v52_ = nw12_
        d_123_v53_: _dafny.Seq
        d_123_v53_ = _dafny.SeqWithoutIsStrInference([d_122_v52_, d_122_v52_, d_122_v52_, d_122_v52_])
        index4_ = default__.safeIndex(986, (d_119_v49_).length(0))
        (d_119_v49_)[index4_] = _dafny.SeqWithoutIsStrInference([False, (d_61_v7_).fm8(len(d_123_v53_), d_53_v0_, (d_121_v51_).f22, d_60_globalState_)])
        index5_ = default__.safeIndex(986, (d_119_v49_).length(0))
        (d_119_v49_)[index5_] = (d_58_v5_) + (default__.fm27(d_60_globalState_))
        d_124_v54_: T0
        nw13_ = C11()
        nw13_.ctor__()
        d_124_v54_ = nw13_
        d_125_v55_: _dafny.Seq
        d_125_v55_ = _dafny.SeqWithoutIsStrInference([d_124_v54_])
        d_126_v56_: _dafny.Map
        d_126_v56_ = _dafny.Map({d_53_v0_: d_125_v55_})
        if (d_58_v5_)[default__.safeIndex(len(((d_126_v56_)[d_53_v0_] if (d_53_v0_) in (d_126_v56_) else d_125_v55_)), len(d_58_v5_))]:
            (d_60_globalState_).f15 = (d_121_v51_).f22
            d_127_v57_: D1
            d_127_v57_ = D1_DC3(d_120_v50_, d_53_v0_)
            source4_ = d_127_v57_
            if source4_.is_DC3:
                d_128___mcc_h5_ = source4_.cf3
                d_129___mcc_h6_ = source4_.cf4
                d_130_cf4_ = d_129___mcc_h6_
                d_131_cf3_ = d_128___mcc_h5_
                (d_60_globalState_).f13 = ((d_121_v51_).f22) + (604)
                d_132_v58_: _dafny.Array
                def lambda2_(d_133_v51_):
                    def lambda3_(d_134_i5_):
                        return (d_134_i5_) - ((d_133_v51_).f22)

                    return lambda3_

                init1_ = lambda2_(d_121_v51_)
                nw14_ = _dafny.Array(None, 6)
                for i0_1_ in range(nw14_.length(0)):
                    nw14_[i0_1_] = init1_(i0_1_)
                d_132_v58_ = nw14_
                d_135_v59_: _dafny.Map
                d_135_v59_ = _dafny.Map({d_53_v0_: d_132_v58_})
                (d_124_v54_).m1(default__.fm19(d_130_cf4_, d_53_v0_, d_60_globalState_), (d_135_v59_) | (d_135_v59_), d_60_globalState_)
                (d_60_globalState_).f13 = (d_121_v51_).f22
                d_136_v60_: str
                d_136_v60_ = _dafny.CodePoint('w')
                d_137_v61_: _dafny.MultiSet
                d_137_v61_ = _dafny.MultiSet([d_136_v60_])
                d_137_v61_ = d_137_v61_
            elif source4_.is_DC4:
                d_138___mcc_h7_ = source4_.cf5
                d_139_cf5_ = d_138___mcc_h7_
                d_140_v62_: _dafny.Map
                d_140_v62_ = _dafny.Map({d_124_v54_: (not(True)) and (d_53_v0_)})
                d_53_v0_ = ((d_140_v62_)[(d_124_v54_ if d_53_v0_ else d_124_v54_)] if ((d_124_v54_ if d_53_v0_ else d_124_v54_)) in (d_140_v62_) else d_53_v0_)
                d_54_v1_ = default__.safeModuloInt(d_54_v1_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbcyxwik")) if d_53_v0_ else d_120_v50_)))
                d_141_v63_: str
                d_141_v63_ = _dafny.CodePoint('r')
                d_142_v64_: _dafny.Array
                def lambda4_(d_143_v0_):
                    def lambda5_(d_144_i6_):
                        return d_143_v0_

                    return lambda5_

                init2_ = lambda4_(d_53_v0_)
                nw15_ = _dafny.Array(None, 18)
                for i0_2_ in range(nw15_.length(0)):
                    nw15_[i0_2_] = init2_(i0_2_)
                d_142_v64_ = nw15_
                index6_ = default__.safeIndex(581, (d_142_v64_).length(0))
                (d_142_v64_)[index6_] = d_53_v0_
                d_145_v65_: _dafny.Array
                def lambda6_(d_146_v1_):
                    def lambda7_(d_147_i7_):
                        return default__.safeModuloInt(d_147_i7_, d_146_v1_)

                    return lambda7_

                init3_ = lambda6_(d_54_v1_)
                nw16_ = _dafny.Array(None, 18)
                for i0_3_ in range(nw16_.length(0)):
                    nw16_[i0_3_] = init3_(i0_3_)
                d_145_v65_ = nw16_
                d_148_v66_: D2
                d_148_v66_ = D2_DC8(D2_DC5(d_145_v65_))
                d_149_v67_: D2
                d_149_v67_ = D2_DC5(d_145_v65_)
                d_150_v68_: D2
                d_150_v68_ = D2_DC8(d_149_v67_)
                pat_let_tv0_ = d_150_v68_
                d_151_v69_: C3
                nw17_ = C3()
                def iife19_(_pat_let0_0):
                    def iife20_(d_152_dt__update__tmp_h0_):
                        def iife21_(_pat_let1_0):
                            def iife22_(d_153_dt__update_hcf11_h0_):
                                return D2_DC8(d_153_dt__update_hcf11_h0_)
                            return iife22_(_pat_let1_0)
                        return iife21_(pat_let_tv0_)
                    return iife20_(_pat_let0_0)
                nw17_.ctor__(iife19_(d_148_v66_), d_53_v0_)
                d_151_v69_ = nw17_
                d_154_v70_: _dafny.Map
                d_154_v70_ = _dafny.Map({d_151_v69_: (d_121_v51_).f22})
                d_155_v71_: _dafny.Set
                d_155_v71_ = _dafny.Set({d_53_v0_, d_53_v0_})
                d_156_v72_: _dafny.Seq
                d_156_v72_ = _dafny.SeqWithoutIsStrInference([d_155_v71_])
                d_157_v73_: _dafny.Array
                nw18_ = _dafny.Array(None, 22)
                nw18_[int(0)] = (len(d_154_v70_)) + ((0) - ((d_151_v69_).fm2((d_121_v51_).fm3(d_151_v69_.f25, d_60_globalState_), d_60_globalState_)))
                nw18_[int(1)] = (d_121_v51_).f22
                nw18_[int(2)] = 801
                nw18_[int(3)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usr"))))
                nw18_[int(4)] = (d_121_v51_).f22
                nw18_[int(5)] = (d_121_v51_).f22
                nw18_[int(6)] = 152
                nw18_[int(7)] = (0) - ((len(d_156_v72_)) * (d_54_v1_))
                nw18_[int(8)] = (d_121_v51_).f22
                nw18_[int(9)] = (len(d_120_v50_)) * (-562)
                nw18_[int(10)] = (len((d_121_v51_).f23)) * (-945)
                nw18_[int(11)] = ((d_121_v51_).f22 if d_53_v0_ else d_54_v1_)
                nw18_[int(12)] = (d_114_v44_)[default__.safeIndex(len(d_55_v2_), len(d_114_v44_))]
                nw18_[int(13)] = (d_121_v51_).f22
                nw18_[int(14)] = 798
                nw18_[int(15)] = -850
                nw18_[int(16)] = (d_121_v51_).f22
                nw18_[int(17)] = ((d_121_v51_).f22) * ((d_121_v51_).f22)
                nw18_[int(18)] = 495
                nw18_[int(19)] = (d_114_v44_)[default__.safeIndex((d_121_v51_).f22, len(d_114_v44_))]
                nw18_[int(20)] = d_54_v1_
                nw18_[int(21)] = -503
                d_157_v73_ = nw18_
                index7_ = default__.safeIndex(303, (d_145_v65_).length(0))
                (d_145_v65_)[index7_] = (d_121_v51_).f22
                index8_ = default__.safeIndex(581, (d_142_v64_).length(0))
                index9_ = default__.safeIndex(303, (d_145_v65_).length(0))
                rhs13_ = d_120_v50_
                rhs14_ = d_141_v63_
                rhs15_ = True
                rhs16_ = d_151_v69_.f25
                rhs17_ = d_54_v1_
                lhs7_ = d_142_v64_
                lhs8_ = default__.safeIndex(581, (d_142_v64_).length(0))
                lhs9_ = d_145_v65_
                lhs10_ = default__.safeIndex(303, (d_145_v65_).length(0))
                d_120_v50_ = rhs13_
                d_141_v63_ = rhs14_
                d_53_v0_ = rhs15_
                lhs7_[lhs8_] = rhs16_
                lhs9_[lhs10_] = rhs17_
                d_57_v4_ = (d_57_v4_).set(True, (d_142_v64_)[default__.safeIndex(581, (d_142_v64_).length(0))])
            elif True:
                d_158___mcc_h8_ = source4_.cf2
                d_159_cf2_ = d_158___mcc_h8_
                (d_60_globalState_).f15 = d_159_cf2_
                d_160_v74_: C13
                nw19_ = C13()
                nw19_.ctor__((d_121_v51_).f22, not((d_87_v28_).fm16(366, d_60_globalState_)))
                d_160_v74_ = nw19_
                d_161_v75_: _dafny.MultiSet
                d_161_v75_ = _dafny.MultiSet([len((d_121_v51_).f23), d_54_v1_])
                rhs18_ = d_160_v74_
                rhs19_ = False
                rhs20_ = default__.safeModuloInt((d_161_v75_).cardinality, (d_160_v74_).f20)
                lhs11_ = d_60_globalState_
                d_160_v74_ = rhs18_
                d_53_v0_ = rhs19_
                lhs11_.f13 = rhs20_
                (d_121_v51_).m2((0) - (len((d_114_v44_) + (d_114_v44_))), d_53_v0_, -853, d_60_globalState_)
                (d_160_v74_).f21 = False
            if d_53_v0_:
                d_162_v76_: _dafny.Array
                nw20_ = _dafny.Array(None, 3)
                nw20_[int(0)] = 259
                nw20_[int(1)] = len((d_121_v51_).f23)
                nw20_[int(2)] = (d_121_v51_).f22
                d_162_v76_ = nw20_
                d_162_v76_ = d_162_v76_
                d_55_v2_ = (d_55_v2_) | (d_55_v2_)
                d_163_v77_: _dafny.Map
                d_163_v77_ = _dafny.Map({d_53_v0_: d_162_v76_})
                d_164_v78_: D23
                d_164_v78_ = D23_DC63(d_163_v77_)
                (d_124_v54_).m1((d_121_v51_).f23, (d_163_v77_) | ((d_164_v78_).cf106), d_60_globalState_)
                d_165_v79_: D15
                d_165_v79_ = D15_DC39(d_57_v4_, not(d_53_v0_), (_dafny.MultiSet([d_54_v1_])).cardinality)
                d_166_v80_: _dafny.Map
                d_166_v80_ = _dafny.Map({985: d_165_v79_})
                d_166_v80_ = d_166_v80_
                d_162_v76_ = d_162_v76_
            elif True:
                d_167_v81_: str
                d_167_v81_ = _dafny.CodePoint('p')
                d_167_v81_ = d_167_v81_
                d_168_v82_: _dafny.Set
                d_168_v82_ = _dafny.Set({d_53_v0_})
                d_169_v83_: D12
                d_169_v83_ = D12_DC32(d_168_v82_)
                d_169_v83_ = D12_DC32((d_168_v82_) | (d_168_v82_))
                d_170_v84_: _dafny.Array
                nw21_ = _dafny.Array(False, 14)
                d_170_v84_ = nw21_
                d_171_v85_: D0
                d_171_v85_ = D0_DC0(d_170_v84_)
                d_172_v86_: _dafny.Seq
                d_172_v86_ = _dafny.SeqWithoutIsStrInference([d_171_v85_])
                (d_60_globalState_).f4 = (len((d_172_v86_ if d_53_v0_ else _dafny.SeqWithoutIsStrInference([d_171_v85_])))) * (d_54_v1_)
                (d_60_globalState_).f4 = 857
                d_173_v87_: T1
                nw22_ = C2()
                nw22_.ctor__()
                d_173_v87_ = nw22_
                d_174_v88_: _dafny.Seq
                d_174_v88_ = _dafny.SeqWithoutIsStrInference([d_173_v87_, d_173_v87_])
                (d_121_v51_).m2(len((d_174_v88_) + (d_174_v88_)), d_53_v0_, (d_121_v51_).f22, d_60_globalState_)
            (d_60_globalState_).f13 = (0) - (d_54_v1_)
            d_175_v89_: _dafny.Array
            nw23_ = _dafny.Array(int(0), 17)
            d_175_v89_ = nw23_
            index10_ = default__.safeIndex(341, (d_175_v89_).length(0))
            (d_175_v89_)[index10_] = ((d_55_v2_)[not(d_53_v0_)] if (not(d_53_v0_)) in (d_55_v2_) else (d_121_v51_).f22)
            index11_ = default__.safeIndex(341, (d_175_v89_).length(0))
            (d_175_v89_)[index11_] = len((d_121_v51_).f23)
        elif True:
            d_176_v90_: _dafny.Set
            d_176_v90_ = _dafny.Set({d_53_v0_})
            if (default__.fm37(d_60_globalState_)).issubset(d_176_v90_):
                d_177_v91_: T1
                nw24_ = C9()
                nw24_.ctor__()
                d_177_v91_ = nw24_
                d_178_v93_: _dafny.Map
                def iife23_():
                    coll19_ = _dafny.Set()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(-884, 510):
                        d_179_v92_: int = compr_19_
                        if ((-884) <= (d_179_v92_)) and ((d_179_v92_) < (510)):
                            coll19_ = coll19_.union(_dafny.Set([(d_179_v92_) + (-656)]))
                    return _dafny.Set(coll19_)
                d_178_v93_ = _dafny.Map({d_177_v91_: _dafny.MultiSet([d_56_v3_, iife23_()
                ])})
                d_178_v93_ = (d_178_v93_).set(d_177_v91_, _dafny.MultiSet([d_56_v3_]))
                d_180_v94_: C7
                nw25_ = C7()
                nw25_.ctor__()
                d_180_v94_ = nw25_
                d_181_v95_: _dafny.Map
                d_181_v95_ = _dafny.Map({d_180_v94_: d_54_v1_})
                (d_60_globalState_).f15 = ((d_181_v95_)[d_180_v94_] if (d_180_v94_) in (d_181_v95_) else (d_121_v51_).f22)
                d_182_v96_: str
                d_182_v96_ = _dafny.CodePoint('c')
                d_183_v97_: _dafny.Map
                d_183_v97_ = _dafny.Map({(d_121_v51_).f22: d_182_v96_})
                d_184_v98_: D19
                d_184_v98_ = D19_DC52(len((d_183_v97_) | (d_183_v97_)), d_53_v0_, d_53_v0_)
                d_185_v99_: _dafny.Map
                d_185_v99_ = _dafny.Map({d_54_v1_: d_53_v0_})
                d_184_v98_ = default__.fm60(d_185_v99_, (0) - (len((d_121_v51_).f23)), d_60_globalState_)
                (d_60_globalState_).f15 = ((d_121_v51_).f22) + ((d_121_v51_).f22)
                d_186_v100_: _dafny.Seq
                d_186_v100_ = _dafny.SeqWithoutIsStrInference([(d_57_v4_) | (d_57_v4_), d_57_v4_, d_57_v4_])
                d_187_v101_: _dafny.Map
                d_187_v101_ = _dafny.Map({896: _dafny.Map({d_53_v0_: d_53_v0_})})
                rhs21_ = (d_121_v51_).f22
                rhs22_ = (d_186_v100_) + ((d_186_v100_) + (_dafny.SeqWithoutIsStrInference([((d_187_v101_)[(d_121_v51_).f22] if ((d_121_v51_).f22) in (d_187_v101_) else d_57_v4_), d_57_v4_])))
                rhs23_ = default__.safeModuloInt(default__.safeDivisionInt(466, d_54_v1_), -194)
                lhs12_ = d_60_globalState_
                lhs13_ = d_60_globalState_
                lhs12_.f4 = rhs21_
                d_186_v100_ = rhs22_
                lhs13_.f7 = rhs23_
            elif True:
                d_188_v102_: _dafny.MultiSet
                d_188_v102_ = _dafny.MultiSet([d_53_v0_, d_53_v0_, d_53_v0_])
                d_189_v103_: _dafny.Map
                d_189_v103_ = _dafny.Map({False: d_188_v102_})
                d_190_v104_: _dafny.Array
                nw26_ = _dafny.Array(int(0), 6)
                d_190_v104_ = nw26_
                d_191_v105_: D2
                d_191_v105_ = D2_DC5(d_190_v104_)
                (d_61_v7_).m6(d_54_v1_, (((d_189_v103_)[True] if (True) in (d_189_v103_) else d_188_v102_)).set(not(d_53_v0_), default__.abs(len(default__.fm32((d_121_v51_).f22, (d_121_v51_).f22, _dafny.Set({d_53_v0_, ((d_57_v4_)[True] if (True) in (d_57_v4_) else d_53_v0_), d_53_v0_}), d_60_globalState_)))), d_53_v0_, d_191_v105_, d_60_globalState_)
                d_192_v106_: _dafny.MultiSet
                d_192_v106_ = _dafny.MultiSet([(d_121_v51_).f22])
                d_193_v107_: _dafny.Map
                d_193_v107_ = _dafny.Map({len(d_114_v44_): d_192_v106_})
                (d_121_v51_).m2(len(d_193_v107_), d_53_v0_, d_54_v1_, d_60_globalState_)
                d_194_v108_: D7
                d_194_v108_ = D7_DC21(len((d_121_v51_).f23), -681, d_53_v0_, d_53_v0_, d_53_v0_)
                (d_121_v51_).m2((d_121_v51_).f22, d_53_v0_, (d_121_v51_).fm2(len(_dafny.Map({not((d_194_v108_).cf32): (d_121_v51_).f22})), d_60_globalState_), d_60_globalState_)
                d_195_v109_: _dafny.Map
                d_195_v109_ = _dafny.Map({d_53_v0_: d_57_v4_})
                d_196_v110_: D14
                d_196_v110_ = D14_DC36(_dafny.SeqWithoutIsStrInference([d_53_v0_, d_53_v0_, d_53_v0_, (d_121_v51_).fm1(d_195_v109_, 79, d_60_globalState_)]))
                d_197_v111_: _dafny.Set
                d_197_v111_ = _dafny.Set({d_114_v44_, d_114_v44_, d_114_v44_, d_114_v44_})
                rhs24_ = default__.safeModuloInt(d_54_v1_, (d_121_v51_).f22)
                rhs25_ = d_196_v110_
                rhs26_ = (d_197_v111_).ispropersubset(d_197_v111_)
                rhs27_ = ((d_55_v2_) == (d_55_v2_)) == ((d_59_v6_).ispropersubset(d_59_v6_))
                lhs14_ = d_60_globalState_
                lhs14_.f7 = rhs24_
                d_196_v110_ = rhs25_
                d_53_v0_ = rhs26_
                d_53_v0_ = rhs27_
                (d_60_globalState_).f4 = (0) - ((d_121_v51_).f22)
            (d_60_globalState_).f15 = ((d_121_v51_).f22) + (len(d_176_v90_))
            d_53_v0_ = (d_176_v90_).ispropersubset(d_176_v90_)
            d_198_v112_: int
            d_199_v113_: _dafny.MultiSet
            d_200_v114_: bool
            out3_: int
            out4_: _dafny.MultiSet
            out5_: bool
            out3_, out4_, out5_ = (d_87_v28_).m9(default__.safeDivisionInt((d_121_v51_).f22, (d_121_v51_).f22), False, d_53_v0_, d_60_globalState_)
            d_198_v112_ = out3_
            d_199_v113_ = out4_
            d_200_v114_ = out5_
            d_200_v114_ = (d_114_v44_) < ((d_114_v44_) + (d_114_v44_))
        d_201_v115_: C5
        nw27_ = C5()
        nw27_.ctor__()
        d_201_v115_ = nw27_
        d_202_v116_: _dafny.Set
        d_202_v116_ = _dafny.Set({d_201_v115_, d_201_v115_})
        d_203_v117_: C1
        nw28_ = C1()
        nw28_.ctor__()
        d_203_v117_ = nw28_
        d_204_v118_: C8
        nw29_ = C8()
        nw29_.ctor__()
        d_204_v118_ = nw29_
        d_205_v119_: _dafny.Map
        d_205_v119_ = _dafny.Map({d_203_v117_: d_204_v118_})
        d_206_v120_: _dafny.Array
        nw30_ = _dafny.Array(None, 13)
        nw30_[int(0)] = (0) - ((d_121_v51_).f22)
        nw30_[int(1)] = (d_121_v51_).f22
        nw30_[int(2)] = (d_54_v1_) - (len(default__.fm5(d_60_globalState_)))
        nw30_[int(3)] = len(((d_119_v49_)[default__.safeIndex(986, (d_119_v49_).length(0))]) + ((d_119_v49_)[default__.safeIndex(986, (d_119_v49_).length(0))]))
        nw30_[int(4)] = d_54_v1_
        nw30_[int(5)] = d_54_v1_
        nw30_[int(6)] = ((d_121_v51_).f22) - ((d_121_v51_).f22)
        nw30_[int(7)] = d_54_v1_
        nw30_[int(8)] = default__.safeDivisionInt(d_54_v1_, d_54_v1_)
        nw30_[int(9)] = d_54_v1_
        nw30_[int(10)] = len(d_202_v116_)
        nw30_[int(11)] = d_54_v1_
        nw30_[int(12)] = default__.safeDivisionInt(len((d_205_v119_).set(d_203_v117_, d_204_v118_)), 382)
        d_206_v120_ = nw30_
        index12_ = default__.safeIndex(52, (d_206_v120_).length(0))
        (d_206_v120_)[index12_] = d_54_v1_
        index13_ = default__.safeIndex(52, (d_206_v120_).length(0))
        (d_206_v120_)[index13_] = default__.safeModuloInt((d_121_v51_).f22, (d_121_v51_).f22)
        d_207_v121_: str
        d_207_v121_ = _dafny.CodePoint('i')
        d_208_v122_: D6
        d_208_v122_ = D6_DC16(d_114_v44_)
        d_209_v123_: _dafny.Map
        d_209_v123_ = _dafny.Map({default__.fm39((d_114_v44_)[default__.safeIndex(d_54_v1_, len(d_114_v44_))], d_207_v121_, (d_208_v122_).cf23, d_53_v0_, d_60_globalState_): ((d_121_v51_).f23).set(default__.safeIndex(d_54_v1_, len((d_121_v51_).f23)), d_207_v121_)})
        d_209_v123_ = (d_209_v123_).set(_dafny.SeqWithoutIsStrInference([len((d_121_v51_).f23) for d_210_i8_ in range(default__.abs(278))]), ((d_121_v51_).f23) + ((d_121_v51_).f23))
        d_211_v124_: _dafny.Map
        d_211_v124_ = _dafny.Map({d_53_v0_: d_120_v50_})
        rhs28_ = (((d_211_v124_)[d_53_v0_] if (d_53_v0_) in (d_211_v124_) else d_120_v50_)) + ((d_121_v51_).f23)
        rhs29_ = (d_121_v51_).f22
        lhs15_ = d_60_globalState_
        d_120_v50_ = rhs28_
        lhs15_.f15 = rhs29_
        d_55_v2_ = ((d_55_v2_) | (d_55_v2_)) | (d_55_v2_)
        _dafny.print(_dafny.string_of(d_53_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_54_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v2_) == (_dafny.Map({True: 215}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v3_) == (_dafny.Set({215}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_57_v4_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_58_v5_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_v6_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_globalState_).f0) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_60_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_globalState_).f6) == (_dafny.Map({True: 215}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_60_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_globalState_).f8) == (_dafny.Set({215}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_globalState_).f11) == (_dafny.Map({False: False, True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_60_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_60_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_globalState_).f16) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_88_v29_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v30_).cf85))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v30_).cf86))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v30_).cf87))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v30_).cf88))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v42_) == (_dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC29()), D11_DC31(D11_DC29())]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_113_v43_).cf80) == (_dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC29()), D11_DC31(D11_DC29())]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v44_) == (_dafny.SeqWithoutIsStrInference([215]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v46_) == (_dafny.MultiSet([_dafny.Set({215}), _dafny.Set({215})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_117_v47_).cf22).cf17) == (_dafny.MultiSet([_dafny.Set({215}), _dafny.Set({215})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v48_) == (_dafny.SeqWithoutIsStrInference([D5_DC15(D5_DC13(_dafny.MultiSet([_dafny.Set({215}), _dafny.Set({215})]))), D5_DC15(D5_DC13(_dafny.MultiSet([_dafny.Set({215}), _dafny.Set({215})]))), D5_DC15(D5_DC13(_dafny.MultiSet([_dafny.Set({215}), _dafny.Set({215})])))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[0]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[1]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[2]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[3]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[4]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[5]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[6]) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[7]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[8]) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_119_v49_)[9]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_120_v50_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v51_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_121_v51_).f23).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v52_)[0]).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_v52_)[0]).f23).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v52_)[1]).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_v52_)[1]).f23).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v52_)[2]).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_v52_)[2]).f23).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_122_v52_)[3]).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_122_v52_)[3]).f23).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_123_v53_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_125_v55_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_126_v56_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_202_v116_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_205_v119_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v120_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_207_v121_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_208_v122_).cf23) == (_dafny.SeqWithoutIsStrInference([215]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v123_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([669, 2, 1, 412, 1, 193, 1, 2]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xofwgic")), _dafny.SeqWithoutIsStrInference([7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xofwgscxofwgsc"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_211_v124_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipyktu"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0))
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
        return lambda: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
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

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf3.VerbatimString(True)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC6(D2, NamedTuple('DC6', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(_dafny.CodePoint('D'), False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC14(D5, NamedTuple('DC14', [('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC17(D6, NamedTuple('DC17', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC20(D7, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC24(D8, NamedTuple('DC24', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC26(D9, NamedTuple('DC26', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC27(D10, NamedTuple('DC27', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC29()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)

class D11_DC29(D11, NamedTuple('DC29', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC28(D11, NamedTuple('DC28', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC33(D12, NamedTuple('DC33', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC34(D13, NamedTuple('DC34', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC36(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)

class D14_DC36(D14, NamedTuple('DC36', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC35(D14, NamedTuple('DC35', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC39(_dafny.Map({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)

class D15_DC39(D15, NamedTuple('DC39', [('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC38(D15, NamedTuple('DC38', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC42(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)

class D16_DC42(D16, NamedTuple('DC42', [('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC41(D16, NamedTuple('DC41', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf66 == __o.cf66
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
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)

class D17_DC46(D17, NamedTuple('DC46', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC47(D17, NamedTuple('DC47', [('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC45(D17, NamedTuple('DC45', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC49(False, False, _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D18_DC48)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)

class D18_DC49(D18, NamedTuple('DC49', [('cf75', Any), ('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC48(D18, NamedTuple('DC48', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC48({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC48) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC50(D18, NamedTuple('DC50', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC52(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D19_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)

class D19_DC52(D19, NamedTuple('DC52', [('cf81', Any), ('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC53(D19, NamedTuple('DC53', [])):
    def __dafnystr__(self) -> str:
        return f'D19.DC53'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC53)
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC51(D19, NamedTuple('DC51', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC55(int(0), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)

class D20_DC55(D20, NamedTuple('DC55', [('cf85', Any), ('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC57(int(0), int(0), D5.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D21_DC58)

class D21_DC57(D21, NamedTuple('DC57', [('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC56(D21, NamedTuple('DC56', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC58(D21, NamedTuple('DC58', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC58({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC58) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC60(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D22_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D22_DC61)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D22_DC62)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)

class D22_DC60(D22, NamedTuple('DC60', [('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC60({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC60) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC61(D22, NamedTuple('DC61', [('cf97', Any), ('cf98', Any), ('cf99', Any), ('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC61({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {self.cf100.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC61) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC62(D22, NamedTuple('DC62', [('cf101', Any), ('cf102', Any), ('cf103', Any), ('cf104', Any), ('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC62({self.cf101.VerbatimString(True)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC62) and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104 and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC59(D22, NamedTuple('DC59', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC64(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D23_DC64)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D23_DC63)

class D23_DC64(D23, NamedTuple('DC64', [('cf107', Any), ('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC64({_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC64) and self.cf107 == __o.cf107 and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC63(D23, NamedTuple('DC63', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC63({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC63) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, globalState):
        pass


class T1:
    pass
    def m6(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f4: int = int(0)
        self.f7: int = int(0)
        self.f13: int = int(0)
        self.f15: int = int(0)
        self._f0: _dafny.Seq = _dafny.Seq({})
        self._f1: bool = False
        self._f2: str = _dafny.CodePoint('D')
        self._f3: int = int(0)
        self._f5: bool = False
        self._f6: _dafny.Map = _dafny.Map({})
        self._f8: _dafny.Set = _dafny.Set({})
        self._f9: bool = False
        self._f10: bool = False
        self._f11: _dafny.Map = _dafny.Map({})
        self._f12: int = int(0)
        self._f14: bool = False
        self._f16: _dafny.MultiSet = _dafny.MultiSet({})
        self._f17: bool = False
        self._f18: int = int(0)
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19

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
    def f14(self):
        return self._f14
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

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm28(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_212_i0_ in range(default__.abs(852))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdqgbyeq")))

    def fm29(self, p0, p1, p2, p3, globalState):
        def iife24_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sosvg"))): 968})).keys.Elements:
                d_213_v0_: int = compr_20_
                if (d_213_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sosvg"))): 968})):
                    coll20_[default__.safeModuloInt(d_213_v0_, (_dafny.MultiSet([False, False])).cardinality)] = -322
            return _dafny.Map(coll20_)
        return (D5_DC14(_dafny.CodePoint('n'), False, len(_dafny.Map({len(_dafny.Map({True: D7_DC21(739, len(iife24_()
), True, not(True), False)})): 206})), not(True))).cf21


class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm25(self, p0, p1, globalState):
        return default__.safeModuloInt(3, default__.safeDivisionInt(-692, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_214_i0_ in range(default__.abs(997))]))))

    def m13(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: D7 = D7.default()()
        r3: D1 = D1.default()()
        d_215_v0_: bool
        d_215_v0_ = True
        d_216_v1_: _dafny.Array
        def lambda8_(d_217_p1_):
            def lambda9_(d_218_i0_):
                return (d_218_i0_) * (d_217_p1_)

            return lambda9_

        init4_ = lambda8_(p1)
        nw31_ = _dafny.Array(None, 21)
        for i0_4_ in range(nw31_.length(0)):
            nw31_[i0_4_] = init4_(i0_4_)
        d_216_v1_ = nw31_
        index14_ = default__.safeIndex(958, (d_216_v1_).length(0))
        (d_216_v1_)[index14_] = default__.safeDivisionInt(p1, p1)
        d_219_v2_: _dafny.Set
        d_219_v2_ = _dafny.Set({p1})
        index15_ = default__.safeIndex(958, (d_216_v1_).length(0))
        rhs30_ = (d_219_v2_).issubset(d_219_v2_)
        rhs31_ = 38
        lhs16_ = d_216_v1_
        lhs17_ = default__.safeIndex(958, (d_216_v1_).length(0))
        d_215_v0_ = rhs30_
        lhs16_[lhs17_] = rhs31_
        d_220_v3_: _dafny.Seq
        d_220_v3_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])])
        d_221_v4_: D6
        d_221_v4_ = D6_DC16(d_220_v3_)
        d_222_v5_: _dafny.Map
        d_222_v5_ = _dafny.Map({d_215_v0_: d_215_v0_})
        index16_ = default__.safeIndex(958, (d_216_v1_).length(0))
        (d_216_v1_)[index16_] = default__.safeDivisionInt((0) - ((_dafny.MultiSet((d_221_v4_).cf23)).cardinality), ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]) * (len(d_222_v5_)))
        d_223_i1_: int
        d_223_i1_ = 0
        with _dafny.label("0"):
            while default__.fm26(globalState):
                with _dafny.c_label("0"):
                    if (d_223_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_223_i1_ = (d_223_i1_) + (1)
                    d_224_v6_: _dafny.Array
                    def lambda10_(d_225_v0_):
                        def lambda11_(d_226_i2_):
                            return d_225_v0_

                        return lambda11_

                    init5_ = lambda10_(d_215_v0_)
                    nw32_ = _dafny.Array(None, 23)
                    for i0_5_ in range(nw32_.length(0)):
                        nw32_[i0_5_] = init5_(i0_5_)
                    d_224_v6_ = nw32_
                    d_227_v7_: _dafny.Seq
                    d_227_v7_ = _dafny.SeqWithoutIsStrInference([d_224_v6_, d_224_v6_])
                    source5_ = D7_DC19((d_227_v7_) + ((d_227_v7_).set(default__.safeIndex(p1, len(d_227_v7_)), d_224_v6_)))
                    if source5_.is_DC20:
                        d_228___mcc_h0_ = source5_.cf29
                        d_229_cf29_ = d_228___mcc_h0_
                        d_230_v8_: _dafny.Seq
                        d_230_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slrsv"))
                        d_231_v9_: str
                        d_231_v9_ = _dafny.CodePoint('c')
                        d_232_v10_: _dafny.Map
                        d_232_v10_ = _dafny.Map({not(not((d_230_v8_) != ((d_230_v8_).set(default__.safeIndex((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))], len(d_230_v8_)), d_231_v9_)))): p1})
                        d_232_v10_ = (d_232_v10_).set(d_215_v0_, p1)
                        d_233_v11_: _dafny.MultiSet
                        d_233_v11_ = _dafny.MultiSet([d_215_v0_])
                        d_234_v12_: _dafny.Seq
                        d_234_v12_ = _dafny.SeqWithoutIsStrInference([d_215_v0_])
                        d_215_v0_ = not((((d_233_v11_) - (_dafny.MultiSet(d_234_v12_))).cardinality) > (d_229_cf29_))
                        d_215_v0_ = default__.fm26(globalState)
                        index17_ = default__.safeIndex(485, (d_224_v6_).length(0))
                        (d_224_v6_)[index17_] = d_215_v0_
                        index18_ = default__.safeIndex(485, (d_224_v6_).length(0))
                        (d_224_v6_)[index18_] = d_215_v0_
                    elif source5_.is_DC21:
                        d_235___mcc_h1_ = source5_.cf30
                        d_236___mcc_h2_ = source5_.cf31
                        d_237___mcc_h3_ = source5_.cf32
                        d_238___mcc_h4_ = source5_.cf33
                        d_239___mcc_h5_ = source5_.cf34
                        d_240_cf34_ = d_239___mcc_h5_
                        d_241_cf33_ = d_238___mcc_h4_
                        d_242_cf32_ = d_237___mcc_h3_
                        d_243_cf31_ = d_236___mcc_h2_
                        d_244_cf30_ = d_235___mcc_h1_
                        d_243_cf31_ = d_243_cf31_
                        d_245_v13_: _dafny.Seq
                        d_245_v13_ = _dafny.SeqWithoutIsStrInference([d_242_cf32_])
                        d_246_v14_: _dafny.Array
                        nw33_ = _dafny.Array(None, 28)
                        nw33_[int(0)] = d_245_v13_
                        nw33_[int(1)] = (d_245_v13_) + ((d_245_v13_).set(default__.safeIndex(d_243_cf31_, len(d_245_v13_)), d_240_cf34_))
                        nw33_[int(2)] = d_245_v13_
                        nw33_[int(3)] = (d_245_v13_).set(default__.safeIndex(p1, len(d_245_v13_)), not(d_215_v0_))
                        nw33_[int(4)] = _dafny.SeqWithoutIsStrInference([d_215_v0_, not(not(d_242_cf32_))])
                        nw33_[int(5)] = d_245_v13_
                        nw33_[int(6)] = d_245_v13_
                        nw33_[int(7)] = default__.fm27(globalState)
                        nw33_[int(8)] = d_245_v13_
                        nw33_[int(9)] = d_245_v13_
                        nw33_[int(10)] = default__.fm27(globalState)
                        nw33_[int(11)] = (d_245_v13_).set(default__.safeIndex(d_243_cf31_, len(d_245_v13_)), default__.fm26(globalState))
                        nw33_[int(12)] = (_dafny.SeqWithoutIsStrInference([True, d_242_cf32_, d_215_v0_, d_240_cf34_])) + (d_245_v13_)
                        nw33_[int(13)] = _dafny.SeqWithoutIsStrInference([d_242_cf32_])
                        nw33_[int(14)] = d_245_v13_
                        nw33_[int(15)] = d_245_v13_
                        nw33_[int(16)] = d_245_v13_
                        nw33_[int(17)] = d_245_v13_
                        nw33_[int(18)] = d_245_v13_
                        nw33_[int(19)] = d_245_v13_
                        nw33_[int(20)] = _dafny.SeqWithoutIsStrInference([d_240_cf34_, d_215_v0_])
                        nw33_[int(21)] = d_245_v13_
                        nw33_[int(22)] = _dafny.SeqWithoutIsStrInference([d_241_cf33_, d_215_v0_])
                        nw33_[int(23)] = _dafny.SeqWithoutIsStrInference([d_241_cf33_, d_240_cf34_, False, True, d_242_cf32_])
                        nw33_[int(24)] = d_245_v13_
                        nw33_[int(25)] = d_245_v13_
                        nw33_[int(26)] = d_245_v13_
                        nw33_[int(27)] = d_245_v13_
                        d_246_v14_ = nw33_
                        d_247_v15_: _dafny.Array
                        d_247_v15_ = d_246_v14_
                        d_248_v16_: _dafny.Seq
                        d_248_v16_ = _dafny.SeqWithoutIsStrInference([(d_220_v3_).set(default__.safeIndex(959, len(d_220_v3_)), p1)])
                        index19_ = default__.safeIndex(958, (d_216_v1_).length(0))
                        rhs32_ = (p1) <= (d_244_cf30_)
                        rhs33_ = (d_247_v15_)
                        rhs34_ = len(d_248_v16_)
                        rhs35_ = (d_244_cf30_) + (p1)
                        lhs18_ = globalState
                        lhs19_ = d_216_v1_
                        lhs20_ = default__.safeIndex(958, (d_216_v1_).length(0))
                        d_241_cf33_ = rhs32_
                        d_246_v14_ = rhs33_
                        lhs18_.f7 = rhs34_
                        lhs19_[lhs20_] = rhs35_
                        d_249_v17_: _dafny.MultiSet
                        d_249_v17_ = _dafny.MultiSet([d_244_cf30_, (0) - ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])])
                        d_250_v18_: str
                        d_250_v18_ = _dafny.CodePoint('r')
                        d_251_v19_: _dafny.Seq
                        d_251_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fk"))
                        d_252_v20_: _dafny.Seq
                        d_252_v20_ = _dafny.SeqWithoutIsStrInference([(d_251_v19_)[default__.safeIndex(p1, len(d_251_v19_))]])
                        d_253_v21_: _dafny.Map
                        d_253_v21_ = _dafny.Map({d_249_v17_: D5_DC14(d_250_v18_, (d_245_v13_)[default__.safeIndex(p1, len(d_245_v13_))], (0) - ((0) - (len(d_251_v19_))), d_215_v0_)})
                        d_253_v21_ = d_253_v21_
                        d_247_v15_ = d_246_v14_
                    elif source5_.is_DC22:
                        d_254___mcc_h6_ = source5_.cf35
                        d_255___mcc_h7_ = source5_.cf36
                        d_256___mcc_h8_ = source5_.cf37
                        d_257___mcc_h9_ = source5_.cf38
                        d_258_cf38_ = d_257___mcc_h9_
                        d_259_cf37_ = d_256___mcc_h8_
                        d_260_cf36_ = d_255___mcc_h7_
                        d_261_cf35_ = d_254___mcc_h6_
                        d_261_cf35_ = not(default__.fm26(globalState))
                        d_262_v22_: _dafny.Set
                        d_262_v22_ = _dafny.Set({d_216_v1_, d_216_v1_})
                        index20_ = default__.safeIndex(958, (d_216_v1_).length(0))
                        (d_216_v1_)[index20_] = len((d_262_v22_).intersection((d_262_v22_) - (d_262_v22_)))
                        d_261_cf35_ = d_261_cf35_
                        d_258_cf38_ = (0) - (-589)
                    elif True:
                        d_263___mcc_h10_ = source5_.cf28
                        d_264_cf28_ = d_263___mcc_h10_
                        d_215_v0_ = d_215_v0_
                        d_215_v0_ = d_215_v0_
                        d_265_v23_: C0
                        nw34_ = C0()
                        nw34_.ctor__()
                        d_265_v23_ = nw34_
                        d_266_v24_: _dafny.Array
                        nw35_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                        d_266_v24_ = nw35_
                        d_267_v25_: _dafny.Seq
                        d_267_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yaw"))
                        index21_ = default__.safeIndex(591, (d_266_v24_).length(0))
                        (d_266_v24_)[index21_] = d_267_v25_
                        d_268_v26_: _dafny.MultiSet
                        d_268_v26_ = _dafny.MultiSet([d_216_v1_])
                        index22_ = default__.safeIndex(591, (d_266_v24_).length(0))
                        rhs36_ = (d_267_v25_) + (d_267_v25_)
                        rhs37_ = (p1) - ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])
                        rhs38_ = not((d_268_v26_).issubset(d_268_v26_))
                        rhs39_ = p1
                        rhs40_ = not(d_215_v0_)
                        lhs21_ = d_266_v24_
                        lhs22_ = default__.safeIndex(591, (d_266_v24_).length(0))
                        lhs23_ = globalState
                        lhs24_ = globalState
                        lhs21_[lhs22_] = rhs36_
                        lhs23_.f7 = rhs37_
                        d_215_v0_ = rhs38_
                        lhs24_.f4 = rhs39_
                        d_215_v0_ = rhs40_
                    d_224_v6_ = d_224_v6_
                    index23_ = default__.safeIndex(424, (d_224_v6_).length(0))
                    (d_224_v6_)[index23_] = d_215_v0_
                    d_269_v27_: _dafny.Array
                    def lambda12_(d_270_i3_):
                        return _dafny.SeqWithoutIsStrInference([False])

                    init6_ = lambda12_
                    nw36_ = _dafny.Array(None, 13)
                    for i0_6_ in range(nw36_.length(0)):
                        nw36_[i0_6_] = init6_(i0_6_)
                    d_269_v27_ = nw36_
                    d_271_v28_: _dafny.Seq
                    d_271_v28_ = _dafny.SeqWithoutIsStrInference([d_269_v27_])
                    d_272_v29_: _dafny.Array
                    d_272_v29_ = (d_271_v28_)[default__.safeIndex(p1, len(d_271_v28_))]
                    d_273_v30_: _dafny.Set
                    d_273_v30_ = _dafny.Set({d_272_v29_})
                    d_274_v31_: _dafny.Map
                    d_274_v31_ = _dafny.Map({(0) - (p1): d_273_v30_})
                    index24_ = default__.safeIndex(424, (d_224_v6_).length(0))
                    (d_224_v6_)[index24_] = (p1) in (d_274_v31_)
                    d_275_v32_: D8
                    d_275_v32_ = D8_DC25((0) - (len(d_220_v3_)))
                    d_276_v33_: _dafny.Seq
                    d_276_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cd"))
                    d_277_v34_: _dafny.Seq
                    d_277_v34_ = _dafny.SeqWithoutIsStrInference([d_215_v0_, d_215_v0_, d_215_v0_])
                    rhs41_ = d_275_v32_
                    rhs42_ = (0) - (-416)
                    rhs43_ = d_216_v1_
                    rhs44_ = (0) - (len(d_276_v33_))
                    rhs45_ = (len(d_277_v34_)) * ((p1) * ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]))
                    lhs25_ = globalState
                    lhs26_ = globalState
                    lhs27_ = globalState
                    d_275_v32_ = rhs41_
                    lhs25_.f4 = rhs42_
                    d_216_v1_ = rhs43_
                    lhs26_.f13 = rhs44_
                    lhs27_.f4 = rhs45_
                    pass
            pass
        d_278_v35_: D2
        d_278_v35_ = D2_DC5(d_216_v1_)
        source6_ = d_278_v35_
        if source6_.is_DC6:
            d_279___mcc_h11_ = source6_.cf7
            d_280_cf7_ = d_279___mcc_h11_
            d_281_v36_: _dafny.Map
            d_281_v36_ = _dafny.Map({(d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]: d_215_v0_})
            d_282_v37_: _dafny.Seq
            d_282_v37_ = _dafny.SeqWithoutIsStrInference([not(d_215_v0_)])
            d_215_v0_ = not(((_dafny.Map({p1: d_215_v0_})) == (d_281_v36_)) == ((d_282_v37_)[default__.safeIndex(d_280_cf7_, len(d_282_v37_))]))
            d_283_v38_: _dafny.Map
            d_283_v38_ = _dafny.Map({(self).fm25(d_215_v0_, d_215_v0_, globalState): _dafny.CodePoint('b')})
            d_284_v39_: _dafny.Array
            nw37_ = _dafny.Array(None, 20)
            nw37_[int(0)] = d_282_v37_
            nw37_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_215_v0_])).set(default__.safeIndex(len(d_283_v38_), len(_dafny.SeqWithoutIsStrInference([d_215_v0_]))), False)
            nw37_[int(2)] = d_282_v37_
            nw37_[int(3)] = d_282_v37_
            nw37_[int(4)] = _dafny.SeqWithoutIsStrInference([d_215_v0_])
            nw37_[int(5)] = d_282_v37_
            nw37_[int(6)] = d_282_v37_
            nw37_[int(7)] = d_282_v37_
            nw37_[int(8)] = d_282_v37_
            nw37_[int(9)] = default__.fm27(globalState)
            nw37_[int(10)] = _dafny.SeqWithoutIsStrInference([d_215_v0_, d_215_v0_, d_215_v0_])
            nw37_[int(11)] = d_282_v37_
            nw37_[int(12)] = d_282_v37_
            nw37_[int(13)] = d_282_v37_
            nw37_[int(14)] = _dafny.SeqWithoutIsStrInference([d_215_v0_, default__.fm26(globalState)])
            nw37_[int(15)] = d_282_v37_
            nw37_[int(16)] = _dafny.SeqWithoutIsStrInference([not(False), True])
            nw37_[int(17)] = d_282_v37_
            nw37_[int(18)] = d_282_v37_
            nw37_[int(19)] = (d_282_v37_).set(default__.safeIndex((self).fm25(d_215_v0_, not(d_215_v0_), globalState), len(d_282_v37_)), d_215_v0_)
            d_284_v39_ = nw37_
            d_285_v40_: _dafny.Array
            d_285_v40_ = d_284_v39_
            source7_ = d_285_v40_
            d_286___mcc_h17_ = source7_
            d_287_cf43_ = d_286___mcc_h17_
            d_288_v41_: _dafny.Seq
            d_288_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owcdswdik"))
            d_289_v42_: _dafny.Seq
            d_289_v42_ = _dafny.SeqWithoutIsStrInference([(d_288_v41_)[default__.safeIndex((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))], len(d_288_v41_))]])
            (globalState).f15 = (0) - (len(d_288_v41_))
            d_290_v43_: D8
            d_290_v43_ = D8_DC23(d_282_v37_)
            d_291_v44_: _dafny.Map
            d_291_v44_ = _dafny.Map({d_215_v0_: d_288_v41_})
            d_292_v45_: _dafny.Map
            d_292_v45_ = _dafny.Map({(d_291_v44_).set(d_215_v0_, d_288_v41_): d_215_v0_})
            d_293_v46_: _dafny.MultiSet
            d_293_v46_ = _dafny.MultiSet([False, not(d_215_v0_), d_215_v0_, ((d_292_v45_)[d_291_v44_] if (d_291_v44_) in (d_292_v45_) else d_215_v0_)])
            rhs46_ = d_290_v43_
            rhs47_ = d_293_v46_
            rhs48_ = d_215_v0_
            d_290_v43_ = rhs46_
            d_293_v46_ = rhs47_
            d_215_v0_ = rhs48_
            (globalState).f7 = (d_280_cf7_) * ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])
            d_294_v47_: _dafny.Map
            d_294_v47_ = _dafny.Map({default__.safeModuloInt(d_280_cf7_, len(d_288_v41_)): d_280_cf7_})
            d_295_v48_: _dafny.MultiSet
            d_295_v48_ = _dafny.MultiSet([(d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]])
            d_296_v49_: _dafny.Set
            d_296_v49_ = _dafny.Set({d_215_v0_})
            d_297_v50_: str
            d_297_v50_ = _dafny.CodePoint('t')
            d_298_v51_: _dafny.Array
            nw38_ = _dafny.Array(None, 11)
            nw38_[int(0)] = (d_296_v49_).isdisjoint(d_296_v49_)
            nw38_[int(1)] = d_215_v0_
            nw38_[int(2)] = (_dafny.Set({len(d_220_v3_)})).isdisjoint(d_219_v2_)
            nw38_[int(3)] = True
            nw38_[int(4)] = ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]) == ((self).fm25(False, d_215_v0_, globalState))
            nw38_[int(5)] = d_215_v0_
            nw38_[int(6)] = d_215_v0_
            nw38_[int(7)] = d_215_v0_
            nw38_[int(8)] = (d_297_v50_) in (d_289_v42_)
            nw38_[int(9)] = (not(d_215_v0_)) and (d_215_v0_)
            nw38_[int(10)] = (_dafny.MultiSet([d_215_v0_])).ispropersubset(_dafny.MultiSet([d_215_v0_, d_215_v0_]))
            d_298_v51_ = nw38_
            index25_ = default__.safeIndex(236, (d_298_v51_).length(0))
            (d_298_v51_)[index25_] = ((d_281_v36_)[(d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]] if ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]) in (d_281_v36_) else True)
            index26_ = default__.safeIndex(236, (d_298_v51_).length(0))
            rhs49_ = d_294_v47_
            rhs50_ = d_295_v48_
            rhs51_ = not (d_215_v0_) or (d_215_v0_)
            lhs28_ = d_298_v51_
            lhs29_ = default__.safeIndex(236, (d_298_v51_).length(0))
            d_294_v47_ = rhs49_
            d_295_v48_ = rhs50_
            lhs28_[lhs29_] = rhs51_
            d_299_v52_: str
            d_299_v52_ = _dafny.CodePoint('a')
            d_300_v53_: _dafny.Seq
            d_300_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uljkewecy"))
            if (d_299_v52_) in ((d_300_v53_) + (d_300_v53_)):
                d_301_v54_: _dafny.Array
                nw39_ = _dafny.Array(None, 5)
                nw39_[int(0)] = ((d_281_v36_)[p1] if (p1) in (d_281_v36_) else d_215_v0_)
                nw39_[int(1)] = d_215_v0_
                nw39_[int(2)] = d_215_v0_
                nw39_[int(3)] = d_215_v0_
                nw39_[int(4)] = d_215_v0_
                d_301_v54_ = nw39_
                index27_ = default__.safeIndex(416, (d_301_v54_).length(0))
                (d_301_v54_)[index27_] = False
                d_302_v55_: _dafny.Array
                def lambda13_(d_303_v0_, d_304_p1_):
                    def lambda14_(d_305_i4_):
                        return (_dafny.Map({not(d_303_v0_): d_304_p1_})) | (_dafny.Map({d_303_v0_: d_304_p1_}))

                    return lambda14_

                init7_ = lambda13_(d_215_v0_, p1)
                nw40_ = _dafny.Array(None, 3)
                for i0_7_ in range(nw40_.length(0)):
                    nw40_[i0_7_] = init7_(i0_7_)
                d_302_v55_ = nw40_
                index28_ = default__.safeIndex(416, (d_301_v54_).length(0))
                rhs52_ = d_215_v0_
                rhs53_ = d_302_v55_
                lhs30_ = d_301_v54_
                lhs31_ = default__.safeIndex(416, (d_301_v54_).length(0))
                lhs30_[lhs31_] = rhs52_
                d_302_v55_ = rhs53_
                index29_ = default__.safeIndex(416, (d_301_v54_).length(0))
                rhs54_ = d_299_v52_
                rhs55_ = not(not(d_215_v0_))
                lhs32_ = d_301_v54_
                lhs33_ = default__.safeIndex(416, (d_301_v54_).length(0))
                d_299_v52_ = rhs54_
                lhs32_[lhs33_] = rhs55_
                index30_ = default__.safeIndex(416, (d_301_v54_).length(0))
                index31_ = default__.safeIndex(416, (d_301_v54_).length(0))
                index32_ = default__.safeIndex(416, (d_301_v54_).length(0))
                rhs56_ = (d_301_v54_)[default__.safeIndex(416, (d_301_v54_).length(0))]
                rhs57_ = d_280_cf7_
                rhs58_ = not (not (True) or (d_215_v0_)) or (d_215_v0_)
                rhs59_ = default__.fm26(globalState)
                rhs60_ = (d_301_v54_)[default__.safeIndex(416, (d_301_v54_).length(0))]
                lhs34_ = d_301_v54_
                lhs35_ = default__.safeIndex(416, (d_301_v54_).length(0))
                lhs36_ = globalState
                lhs37_ = d_301_v54_
                lhs38_ = default__.safeIndex(416, (d_301_v54_).length(0))
                lhs39_ = d_301_v54_
                lhs40_ = default__.safeIndex(416, (d_301_v54_).length(0))
                lhs34_[lhs35_] = rhs56_
                lhs36_.f7 = rhs57_
                lhs37_[lhs38_] = rhs58_
                d_215_v0_ = rhs59_
                lhs39_[lhs40_] = rhs60_
                d_306_v56_: _dafny.Map
                d_306_v56_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([(d_301_v54_)[default__.safeIndex(416, (d_301_v54_).length(0))]])})
                r0 = ((d_283_v38_)[(879) - (len(((d_306_v56_)[(d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]] if ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]) in (d_306_v56_) else d_282_v37_)))] if ((879) - (len(((d_306_v56_)[(d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]] if ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]) in (d_306_v56_) else d_282_v37_)))) in (d_283_v38_) else d_299_v52_)
                d_307_v57_: _dafny.Set
                d_307_v57_ = _dafny.Set({d_215_v0_})
                d_308_v58_: D7
                d_308_v58_ = D7_DC22(d_215_v0_, d_280_cf7_, (d_282_v37_)[default__.safeIndex(p1, len(d_282_v37_))], (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])
                d_309_v59_: _dafny.Seq
                d_309_v59_ = _dafny.SeqWithoutIsStrInference([d_308_v58_, d_308_v58_])
                def iife25_():
                    coll21_ = _dafny.Set()
                    compr_21_: int
                    for compr_21_ in (d_220_v3_).Elements:
                        d_310_v60_: int = compr_21_
                        if (d_310_v60_) in (d_220_v3_):
                            coll21_ = coll21_.union(_dafny.Set([default__.safeDivisionInt(d_310_v60_, 22)]))
                    return _dafny.Set(coll21_)
                d_300_v53_ = (((d_300_v53_).set(default__.safeIndex(len(d_307_v57_), len(d_300_v53_)), d_299_v52_)) + (default__.fm30(len(d_309_v59_), default__.fm26(globalState), iife25_()
                , (d_301_v54_)[default__.safeIndex(416, (d_301_v54_).length(0))], globalState))) + (d_300_v53_)
            elif True:
                d_311_v61_: _dafny.Array
                nw41_ = _dafny.Array(None, 12)
                nw41_[int(0)] = d_215_v0_
                nw41_[int(1)] = True
                nw41_[int(2)] = d_215_v0_
                nw41_[int(3)] = d_215_v0_
                nw41_[int(4)] = d_215_v0_
                nw41_[int(5)] = d_215_v0_
                nw41_[int(6)] = d_215_v0_
                nw41_[int(7)] = d_215_v0_
                nw41_[int(8)] = d_215_v0_
                nw41_[int(9)] = d_215_v0_
                nw41_[int(10)] = True
                nw41_[int(11)] = d_215_v0_
                d_311_v61_ = nw41_
                d_215_v0_ = (default__.safeDivisionInt((0) - (len(_dafny.Map({d_311_v61_: 292}))), d_280_cf7_)) <= ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])
                (globalState).f7 = (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]
                d_312_v62_: _dafny.Array
                def lambda15_(d_313_cf7_):
                    def lambda16_(d_314_i5_):
                        return D4_DC12(d_313_cf7_)

                    return lambda16_

                init8_ = lambda15_(d_280_cf7_)
                nw42_ = _dafny.Array(None, 29)
                for i0_8_ in range(nw42_.length(0)):
                    nw42_[i0_8_] = init8_(i0_8_)
                d_312_v62_ = nw42_
                d_315_v63_: D4
                d_315_v63_ = D4_DC12(d_280_cf7_)
                index33_ = default__.safeIndex(59, (d_312_v62_).length(0))
                (d_312_v62_)[index33_] = d_315_v63_
                index34_ = default__.safeIndex(59, (d_312_v62_).length(0))
                (d_312_v62_)[index34_] = d_315_v63_
                d_316_v64_: _dafny.Map
                d_316_v64_ = _dafny.Map({((d_222_v5_)[d_215_v0_] if (d_215_v0_) in (d_222_v5_) else d_215_v0_): p1})
                d_317_v65_: _dafny.MultiSet
                d_317_v65_ = _dafny.MultiSet([d_215_v0_, d_215_v0_])
                index35_ = default__.safeIndex(419, (d_311_v61_).length(0))
                (d_311_v61_)[index35_] = ((d_317_v65_).cardinality) == (-438)
                d_318_v66_: D7
                d_318_v66_ = D7_DC21(376, d_280_cf7_, d_215_v0_, not(d_215_v0_), d_215_v0_)
                d_319_v67_: D6
                d_319_v67_ = D6_DC17((d_318_v66_).cf33, d_215_v0_, d_215_v0_)
                d_320_v68_: _dafny.MultiSet
                d_320_v68_ = _dafny.MultiSet([d_319_v67_])
                index36_ = default__.safeIndex(419, (d_311_v61_).length(0))
                rhs61_ = (d_316_v64_) | (d_316_v64_)
                rhs62_ = d_215_v0_
                rhs63_ = ((105) - ((d_320_v68_).cardinality)) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_215_v0_: d_215_v0_})) for d_321_i6_ in range(default__.abs(-919))]))
                rhs64_ = -271
                rhs65_ = (d_220_v3_) != (d_220_v3_)
                lhs41_ = d_311_v61_
                lhs42_ = default__.safeIndex(419, (d_311_v61_).length(0))
                lhs43_ = globalState
                d_316_v64_ = rhs61_
                lhs41_[lhs42_] = rhs62_
                d_215_v0_ = rhs63_
                lhs43_.f15 = rhs64_
                d_215_v0_ = rhs65_
                d_322_v69_: D6
                d_322_v69_ = D6_DC16(d_220_v3_)
                d_323_v70_: D6
                d_323_v70_ = D6_DC18(d_322_v69_)
                d_324_v71_: _dafny.MultiSet
                d_324_v71_ = _dafny.MultiSet([d_323_v70_])
                d_325_v72_: _dafny.Seq
                d_325_v72_ = _dafny.SeqWithoutIsStrInference([D6_DC18(d_322_v69_)])
                d_324_v71_ = ((d_324_v71_) - (d_324_v71_)).intersection(_dafny.MultiSet(d_325_v72_))
            r1 = d_280_cf7_
        elif source6_.is_DC7:
            d_326___mcc_h12_ = source6_.cf8
            d_327___mcc_h13_ = source6_.cf9
            d_328___mcc_h14_ = source6_.cf10
            d_329_cf10_ = d_328___mcc_h14_
            d_330_cf9_ = d_327___mcc_h13_
            d_331_cf8_ = d_326___mcc_h12_
            d_332_v75_: _dafny.Array
            def lambda17_(d_333_p1_, d_334_cf8_):
                def lambda18_(d_335_i7_):
                    def iife26_():
                        def iife28_():
                            coll24_ = _dafny.Set()
                            compr_24_: _dafny.MultiSet
                            for compr_24_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([130])])).Elements:
                                d_338_v73_: _dafny.MultiSet = compr_24_
                                if (d_338_v73_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([130])])):
                                    coll24_ = coll24_.union(_dafny.Set([d_338_v73_]))
                            return _dafny.Set(coll24_)
                        coll22_ = _dafny.Set()
                        def iife27_():
                            coll23_ = _dafny.Set()
                            compr_23_: _dafny.MultiSet
                            for compr_23_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([130])])).Elements:
                                d_336_v73_: _dafny.MultiSet = compr_23_
                                if (d_336_v73_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([130])])):
                                    coll23_ = coll23_.union(_dafny.Set([d_336_v73_]))
                            return _dafny.Set(coll23_)
                        compr_22_: _dafny.MultiSet
                        for compr_22_ in (iife27_()
                        ).Elements:
                            d_337_v74_: _dafny.MultiSet = compr_22_
                            if (d_337_v74_) in (iife28_()
                            ):
                                coll22_ = coll22_.union(_dafny.Set([d_337_v74_]))
                        return _dafny.Set(coll22_)
                    return (_dafny.Set({_dafny.MultiSet([d_333_p1_, d_333_p1_]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_334_cf8_, d_334_cf8_]))])})).intersection(iife26_()
                    )

                return lambda18_

            init9_ = lambda17_(p1, d_331_cf8_)
            nw43_ = _dafny.Array(None, 2)
            for i0_9_ in range(nw43_.length(0)):
                nw43_[i0_9_] = init9_(i0_9_)
            d_332_v75_ = nw43_
            d_339_v76_: _dafny.MultiSet
            d_339_v76_ = _dafny.MultiSet([p1, (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]])
            d_340_v77_: _dafny.Set
            d_340_v77_ = _dafny.Set({d_339_v76_, d_339_v76_, default__.fm31((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))], (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))], globalState)})
            index37_ = default__.safeIndex(578, (d_332_v75_).length(0))
            (d_332_v75_)[index37_] = d_340_v77_
            index38_ = default__.safeIndex(578, (d_332_v75_).length(0))
            (d_332_v75_)[index38_] = (d_340_v77_).intersection(d_340_v77_)
            (globalState).f13 = p1
            d_341_v78_: _dafny.Seq
            d_341_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vs"))
            index39_ = default__.safeIndex(958, (d_216_v1_).length(0))
            (d_216_v1_)[index39_] = len((d_341_v78_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pq"))) + (d_341_v78_)))
            if d_215_v0_:
                (globalState).f7 = (0) - (p1)
                d_342_v79_: C0
                nw44_ = C0()
                nw44_.ctor__()
                d_342_v79_ = nw44_
                (globalState).f15 = len((d_341_v78_) + ((d_341_v78_) + (d_341_v78_)))
                (globalState).f13 = (0) - (((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]) + ((self).fm25(d_331_cf8_, d_331_cf8_, globalState)))
                d_343_v80_: _dafny.Map
                d_343_v80_ = _dafny.Map({d_278_v35_: d_215_v0_})
                d_344_v81_: _dafny.Map
                d_344_v81_ = d_343_v80_
                d_343_v80_ = (d_344_v81_)
            elif True:
                d_216_v1_ = d_329_cf10_
                d_345_v82_: C0
                nw45_ = C0()
                nw45_.ctor__()
                d_345_v82_ = nw45_
                d_346_v83_: str
                d_346_v83_ = _dafny.CodePoint('w')
                r1 = len((d_341_v78_) + ((d_341_v78_).set(default__.safeIndex(643, len(d_341_v78_)), d_346_v83_)))
                d_331_cf8_ = d_331_cf8_
                d_347_v84_: C0
                nw46_ = C0()
                nw46_.ctor__()
                d_347_v84_ = nw46_
        elif source6_.is_DC5:
            d_348___mcc_h15_ = source6_.cf6
            d_349_cf6_ = d_348___mcc_h15_
            d_350_v85_: _dafny.Seq
            d_350_v85_ = _dafny.SeqWithoutIsStrInference([(d_215_v0_) == (d_215_v0_), default__.fm26(globalState), not (d_215_v0_) or (d_215_v0_), True, d_215_v0_])
            d_350_v85_ = ((d_350_v85_) + (d_350_v85_)).set(default__.safeIndex(782, len((d_350_v85_) + (d_350_v85_))), d_215_v0_)
            d_215_v0_ = (p1) > ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))])
            d_278_v35_ = d_278_v35_
            d_215_v0_ = (default__.safeDivisionInt(p1, p1)) != ((114) * (p1))
        elif True:
            d_351___mcc_h16_ = source6_.cf11
            d_352_cf11_ = d_351___mcc_h16_
            d_215_v0_ = default__.fm26(globalState)
            index40_ = default__.safeIndex(958, (d_216_v1_).length(0))
            (d_216_v1_)[index40_] = (0) - (default__.safeDivisionInt((0) - ((d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]), (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]))
            d_353_v86_: _dafny.Seq
            d_353_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxgofwv"))
            d_354_v87_: str
            d_354_v87_ = _dafny.CodePoint('e')
            d_355_v88_: _dafny.Map
            d_355_v88_ = _dafny.Map({d_215_v0_: d_354_v87_})
            d_356_v89_: _dafny.Seq
            d_356_v89_ = _dafny.SeqWithoutIsStrInference([d_215_v0_])
            d_357_v90_: _dafny.Map
            d_357_v90_ = _dafny.Map({len(d_220_v3_): d_353_v86_})
            d_358_v91_: _dafny.Array
            nw47_ = _dafny.Array(None, 12)
            nw47_[int(0)] = d_353_v86_
            nw47_[int(1)] = (d_353_v86_) + (d_353_v86_)
            nw47_[int(2)] = d_353_v86_
            nw47_[int(3)] = d_353_v86_
            nw47_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsd"))) + (default__.fm30(len(d_219_v2_), d_215_v0_, _dafny.Set({len(d_355_v88_), (_dafny.MultiSet(d_356_v89_)).cardinality}), d_215_v0_, globalState))
            nw47_[int(5)] = d_353_v86_
            nw47_[int(6)] = ((d_357_v90_)[len(d_219_v2_)] if (len(d_219_v2_)) in (d_357_v90_) else d_353_v86_)
            nw47_[int(7)] = d_353_v86_
            nw47_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otlcts"))
            nw47_[int(9)] = d_353_v86_
            nw47_[int(10)] = d_353_v86_
            nw47_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ke"))
            d_358_v91_ = nw47_
            index41_ = default__.safeIndex(954, (d_358_v91_).length(0))
            (d_358_v91_)[index41_] = d_353_v86_
            index42_ = default__.safeIndex(954, (d_358_v91_).length(0))
            rhs66_ = d_353_v86_
            rhs67_ = p1
            lhs44_ = d_358_v91_
            lhs45_ = default__.safeIndex(954, (d_358_v91_).length(0))
            lhs46_ = globalState
            lhs44_[lhs45_] = rhs66_
            lhs46_.f15 = rhs67_
            d_359_v92_: D7
            d_359_v92_ = D7_DC20(p1)
            d_359_v92_ = d_359_v92_
        d_360_v93_: _dafny.Array
        nw48_ = _dafny.Array(False, 20)
        d_360_v93_ = nw48_
        d_361_v94_: _dafny.Seq
        d_361_v94_ = _dafny.SeqWithoutIsStrInference([d_215_v0_])
        index43_ = default__.safeIndex(687, (d_360_v93_).length(0))
        (d_360_v93_)[index43_] = (d_361_v94_)[default__.safeIndex(p1, len(d_361_v94_))]
        d_362_v95_: _dafny.MultiSet
        d_362_v95_ = _dafny.MultiSet([d_215_v0_])
        d_363_v96_: _dafny.Set
        d_363_v96_ = _dafny.Set({d_362_v95_})
        d_364_v98_: _dafny.Map
        d_364_v98_ = _dafny.Map({p1: ((d_362_v95_)[d_215_v0_] if (d_215_v0_) in (d_362_v95_) else 909)})
        index44_ = default__.safeIndex(687, (d_360_v93_).length(0))
        def iife29_():
            coll25_ = _dafny.Set()
            compr_25_: _dafny.MultiSet
            for compr_25_ in (_dafny.SeqWithoutIsStrInference([d_362_v95_])).Elements:
                d_365_v97_: _dafny.MultiSet = compr_25_
                if (d_365_v97_) in (_dafny.SeqWithoutIsStrInference([d_362_v95_])):
                    coll25_ = coll25_.union(_dafny.Set([d_365_v97_]))
            return _dafny.Set(coll25_)
        rhs68_ = d_215_v0_
        rhs69_ = (d_363_v96_).isdisjoint((iife29_()
        ) - (d_363_v96_))
        rhs70_ = (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]
        rhs71_ = (p1) - ((len(d_364_v98_)) + (188))
        rhs72_ = d_215_v0_
        lhs47_ = globalState
        lhs48_ = globalState
        lhs49_ = d_360_v93_
        lhs50_ = default__.safeIndex(687, (d_360_v93_).length(0))
        d_215_v0_ = rhs68_
        d_215_v0_ = rhs69_
        lhs47_.f15 = rhs70_
        lhs48_.f15 = rhs71_
        lhs49_[lhs50_] = rhs72_
        hi1_ = (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))]
        for d_366_i8_ in range((self).fm25(d_215_v0_, d_215_v0_, globalState), hi1_):
            d_220_v3_ = (d_220_v3_) + (_dafny.SeqWithoutIsStrInference([574 for d_367_i9_ in range(default__.abs(612))]))
            (globalState).f15 = ((d_364_v98_)[len((d_219_v2_).intersection(d_219_v2_))] if (len((d_219_v2_).intersection(d_219_v2_))) in (d_364_v98_) else ((d_220_v3_)[default__.safeIndex(p1, len(d_220_v3_))]) - (d_366_i8_))
            d_368_v99_: _dafny.Set
            d_368_v99_ = _dafny.Set({d_215_v0_, d_215_v0_, d_215_v0_, not((d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))]), True})
            index45_ = default__.safeIndex(687, (d_360_v93_).length(0))
            rhs73_ = ((d_361_v94_).set(default__.safeIndex(p1, len(d_361_v94_)), (d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))])) + ((d_361_v94_) + (d_361_v94_))
            rhs74_ = (self).fm25(d_215_v0_, (d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))], globalState)
            rhs75_ = (default__.fm32(len(d_361_v94_), (d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))], d_368_v99_, globalState)) == ((default__.fm32(len(_dafny.SeqWithoutIsStrInference([(d_216_v1_)[default__.safeIndex(958, (d_216_v1_).length(0))] for d_369_i10_ in range(default__.abs(188))])), 466, d_368_v99_, globalState)) - (d_219_v2_))
            lhs51_ = globalState
            lhs52_ = d_360_v93_
            lhs53_ = default__.safeIndex(687, (d_360_v93_).length(0))
            d_361_v94_ = rhs73_
            lhs51_.f7 = rhs74_
            lhs52_[lhs53_] = rhs75_
            (globalState).f15 = default__.safeDivisionInt((self).fm25(not(not(True)), (d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))], globalState), d_366_i8_)
        d_370_v100_: D6
        d_370_v100_ = D6_DC17(True, (d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))], not((d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))]))
        pat_let_tv1_ = d_360_v93_
        pat_let_tv2_ = d_360_v93_
        def lambda19_(source8_):
            if source8_.is_DC17:
                d_371___mcc_h18_ = source8_.cf24
                d_372___mcc_h19_ = source8_.cf25
                d_373___mcc_h20_ = source8_.cf26
                d_374_cf26_ = d_373___mcc_h20_
                d_375_cf25_ = d_372___mcc_h19_
                d_376_cf24_ = d_371___mcc_h18_
                return _dafny.CodePoint('c')
            elif source8_.is_DC16:
                d_377___mcc_h21_ = source8_.cf23
                d_378_cf23_ = d_377___mcc_h21_
                if (pat_let_tv2_)[default__.safeIndex(687, (pat_let_tv1_).length(0))]:
                    return _dafny.CodePoint('b')
                elif True:
                    return _dafny.CodePoint('y')
            elif True:
                d_379___mcc_h22_ = source8_.cf27
                d_380_cf27_ = d_379___mcc_h22_
                return _dafny.CodePoint('x')

        r0 = lambda19_(d_370_v100_)
        r1 = p1
        d_381_v101_: _dafny.Seq
        d_381_v101_ = _dafny.SeqWithoutIsStrInference([d_360_v93_])
        r2 = (D7_DC19(d_381_v101_) if not((d_360_v93_)[default__.safeIndex(687, (d_360_v93_).length(0))]) else p0)
        d_382_v102_: _dafny.Seq
        d_382_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpsnk"))
        d_383_v103_: D1
        d_383_v103_ = D1_DC3(d_382_v102_, d_215_v0_)
        r3 = (D1_DC3((d_383_v103_).cf3, d_215_v0_) if (d_215_v0_) and (False) else d_383_v103_)
        return r0, r1, r2, r3


class C2(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def fm7(self, p0, globalState):
        def iife30_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in (_dafny.MultiSet([258, len(_dafny.Set({False, False, True, True}))])).Elements:
                d_384_v0_: int = compr_26_
                if (d_384_v0_) in (_dafny.MultiSet([258, len(_dafny.Set({False, False, True, True}))])):
                    def iife31_():
                        coll27_ = _dafny.Set()
                        compr_27_: int
                        for compr_27_ in _dafny.IntegerRange(568, 301):
                            d_385_v1_: int = compr_27_
                            if ((568) <= (d_385_v1_)) and ((d_385_v1_) < (301)):
                                coll27_ = coll27_.union(_dafny.Set([(d_385_v1_) * (len(_dafny.Map({_dafny.CodePoint('y'): not(True)})))]))
                        return _dafny.Set(coll27_)
                    coll26_ = coll26_.union(_dafny.Set([default__.safeModuloInt(d_384_v0_, len(iife31_()
))]))
            return _dafny.Set(coll26_)
        return _dafny.MultiSet([(_dafny.Set({117})) - (_dafny.Set({(0) - (len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality: 679})))})), _dafny.Set({(0) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjblw")): (D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywhli")), False)).cf4}))), (0) - (len(_dafny.Set({(D7_DC20(898)).cf29, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (-596): 483}))])), 46})))}), iife30_()
        , (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, not(False)]))})) - (_dafny.Set({-746}))])

    def fm8(self, p0, p1, p2, globalState):
        if not (True) or (not(False)):
            def iife32_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in (_dafny.MultiSet([813])).Elements:
                    d_386_v0_: int = compr_28_
                    if (d_386_v0_) in (_dafny.MultiSet([813])):
                        coll28_[(d_386_v0_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([595, len(_dafny.Map({True: len(_dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('a')})): True}))}))]))])))] = True
                return _dafny.Map(coll28_)
            return not((D3_DC10(len(iife32_()
), len(_dafny.SeqWithoutIsStrInference([True, True])))) in (_dafny.Map({D3_DC10(818, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fve")))): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aclnkmm"))))})))
        elif True:
            return (653) < (len(_dafny.SeqWithoutIsStrInference([782 for d_387_i0_ in range(default__.abs(488))])))

    def fm1(self, p0, p1, globalState):
        return True

    def fm2(self, p0, globalState):
        return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_388_i0_ in range(default__.abs(690))])), -61)

    def fm24(self, p0, p1, p2, p3, globalState):
        return not(True)

    def m6(self, p0, p1, p2, p3, globalState):
        hi2_ = default__.safeModuloInt((0) - (p0), 693)
        for d_389_i0_ in range(p0, hi2_):
            d_390_v0_: bool
            d_390_v0_ = False
            d_391_v1_: _dafny.Array
            nw49_ = _dafny.Array(False, 25)
            d_391_v1_ = nw49_
            d_392_v2_: _dafny.MultiSet
            d_392_v2_ = _dafny.MultiSet([d_391_v1_, d_391_v1_])
            d_390_v0_ = (d_391_v1_) not in ((d_392_v2_).intersection(d_392_v2_))
            (globalState).f4 = d_389_i0_
            d_393_v3_: _dafny.Array
            nw50_ = _dafny.Array(int(0), 14)
            d_393_v3_ = nw50_
            index46_ = default__.safeIndex(591, (d_393_v3_).length(0))
            (d_393_v3_)[index46_] = p0
            d_394_v4_: _dafny.Map
            d_394_v4_ = _dafny.Map({p2: d_390_v0_})
            d_395_v5_: _dafny.Map
            d_395_v5_ = _dafny.Map({d_390_v0_: d_394_v4_})
            index47_ = default__.safeIndex(591, (d_393_v3_).length(0))
            rhs76_ = default__.safeModuloInt(default__.safeDivisionInt(p0, d_389_i0_), len(_dafny.Set({d_390_v0_})))
            rhs77_ = (self).fm1((d_395_v5_) | ((d_395_v5_).set(p2, d_394_v4_)), d_389_i0_, globalState)
            lhs54_ = d_393_v3_
            lhs55_ = default__.safeIndex(591, (d_393_v3_).length(0))
            lhs54_[lhs55_] = rhs76_
            d_390_v0_ = rhs77_
            d_390_v0_ = not(d_390_v0_)
        d_396_v6_: bool
        d_396_v6_ = False
        d_397_v7_: _dafny.Seq
        d_397_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtrupg"))
        d_396_v6_ = (default__.safeModuloInt(p0, p0)) >= (default__.safeDivisionInt(len(d_397_v7_), -925))
        d_398_i1_: int
        d_398_i1_ = 0
        with _dafny.label("1"):
            while d_396_v6_:
                with _dafny.c_label("1"):
                    if (d_398_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_398_i1_ = (d_398_i1_) + (1)
                    d_396_v6_ = d_396_v6_
                    d_399_v8_: _dafny.Set
                    d_399_v8_ = _dafny.Set({d_396_v6_})
                    d_400_v9_: _dafny.Map
                    d_400_v9_ = _dafny.Map({p0: _dafny.MultiSet([d_399_v8_, d_399_v8_])})
                    d_401_v10_: _dafny.MultiSet
                    d_401_v10_ = _dafny.MultiSet([d_399_v8_])
                    d_400_v9_ = (d_400_v9_).set(p0, d_401_v10_)
                    d_402_v11_: C0
                    nw51_ = C0()
                    nw51_.ctor__()
                    d_402_v11_ = nw51_
                    d_403_v12_: _dafny.Array
                    nw52_ = _dafny.Array(_dafny.Set({}), 9)
                    d_403_v12_ = nw52_
                    index48_ = default__.safeIndex(592, (d_403_v12_).length(0))
                    (d_403_v12_)[index48_] = d_399_v8_
                    index49_ = default__.safeIndex(592, (d_403_v12_).length(0))
                    (d_403_v12_)[index49_] = (d_399_v8_).intersection((d_399_v8_).intersection(d_399_v8_))
                    pass
            pass
        d_404_v13_: _dafny.Array
        nw53_ = _dafny.Array(_dafny.Set({}), 11)
        d_404_v13_ = nw53_
        d_405_v14_: _dafny.Set
        d_405_v14_ = _dafny.Set({True})
        d_406_v15_: _dafny.Set
        d_406_v15_ = _dafny.Set({d_405_v14_, d_405_v14_, d_405_v14_})
        index50_ = default__.safeIndex(801, (d_404_v13_).length(0))
        (d_404_v13_)[index50_] = (d_406_v15_ if d_396_v6_ else d_406_v15_)
        index51_ = default__.safeIndex(801, (d_404_v13_).length(0))
        (d_404_v13_)[index51_] = d_406_v15_
        d_407_v16_: _dafny.Array
        nw54_ = _dafny.Array(_dafny.Seq({}), 28)
        d_407_v16_ = nw54_
        d_408_v17_: _dafny.Set
        d_408_v17_ = _dafny.Set({p0, (self).fm2(len(_dafny.Map({p0: not(p2)})), globalState)})
        d_409_v18_: _dafny.Map
        d_409_v18_ = _dafny.Map({p2: False})
        d_410_v19_: _dafny.Map
        d_410_v19_ = _dafny.Map({p0: False})
        d_411_v20_: _dafny.Map
        d_411_v20_ = _dafny.Map({756: d_408_v17_})
        d_412_v21_: _dafny.Map
        d_412_v21_ = _dafny.Map({(d_397_v7_) + (default__.fm30(p0, d_396_v6_, d_408_v17_, d_396_v6_, globalState)): (d_409_v18_) | (default__.fm33(not(p2), ((d_410_v19_)[p0] if (p0) in (d_410_v19_) else d_396_v6_), len(((d_411_v20_)[p0] if (p0) in (d_411_v20_) else d_408_v17_)), globalState))})
        d_413_v22_: _dafny.Array
        nw55_ = _dafny.Array(_dafny.MultiSet({}), 2)
        d_413_v22_ = nw55_
        d_414_v23_: C0
        nw56_ = C0()
        nw56_.ctor__()
        d_414_v23_ = nw56_
        d_415_v24_: _dafny.Map
        d_415_v24_ = _dafny.Map({d_414_v23_: p0})
        d_416_v25_: _dafny.MultiSet
        d_416_v25_ = _dafny.MultiSet([d_415_v24_])
        index52_ = default__.safeIndex(977, (d_413_v22_).length(0))
        (d_413_v22_)[index52_] = (d_416_v25_) | (d_416_v25_)
        d_417_v27_: D11
        d_417_v27_ = D11_DC28(d_408_v17_)
        d_418_v28_: _dafny.MultiSet
        d_418_v28_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_419_i2_ in range(default__.abs(954))]), default__.fm30(p0, p2, (d_417_v27_).cf45, p2, globalState), d_397_v7_])
        index53_ = default__.safeIndex(977, (d_413_v22_).length(0))
        def iife33_():
            coll29_ = _dafny.Map()
            compr_29_: _dafny.Seq
            for compr_29_ in (d_418_v28_).Elements:
                d_420_v26_: _dafny.Seq = compr_29_
                if (d_420_v26_) in (d_418_v28_):
                    coll29_[d_420_v26_] = d_409_v18_
            return _dafny.Map(coll29_)
        rhs78_ = d_407_v16_
        rhs79_ = (iife33_()
        ).set(d_397_v7_, (d_409_v18_) | (d_409_v18_))
        rhs80_ = d_416_v25_
        rhs81_ = p0
        rhs82_ = d_396_v6_
        lhs56_ = d_413_v22_
        lhs57_ = default__.safeIndex(977, (d_413_v22_).length(0))
        lhs58_ = globalState
        d_407_v16_ = rhs78_
        d_412_v21_ = rhs79_
        lhs56_[lhs57_] = rhs80_
        lhs58_.f4 = rhs81_
        d_396_v6_ = rhs82_
        d_421_v29_: _dafny.Array
        def lambda20_(d_422_v6_):
            def lambda21_(d_423_i3_):
                return (_dafny.CodePoint('g') if d_422_v6_ else _dafny.CodePoint('e'))

            return lambda21_

        init10_ = lambda20_(d_396_v6_)
        nw57_ = _dafny.Array(None, 25)
        for i0_10_ in range(nw57_.length(0)):
            nw57_[i0_10_] = init10_(i0_10_)
        d_421_v29_ = nw57_
        d_424_v30_: str
        d_424_v30_ = _dafny.CodePoint('l')
        index54_ = default__.safeIndex(341, (d_421_v29_).length(0))
        (d_421_v29_)[index54_] = d_424_v30_
        d_425_v31_: _dafny.Array
        nw58_ = _dafny.Array(None, 1)
        nw58_[int(0)] = p2
        d_425_v31_ = nw58_
        index55_ = default__.safeIndex(89, (d_425_v31_).length(0))
        (d_425_v31_)[index55_] = d_396_v6_
        index56_ = default__.safeIndex(341, (d_421_v29_).length(0))
        index57_ = default__.safeIndex(89, (d_425_v31_).length(0))
        rhs83_ = d_424_v30_
        rhs84_ = D11_DC28(d_408_v17_)
        rhs85_ = p2
        rhs86_ = p0
        rhs87_ = p0
        lhs59_ = d_421_v29_
        lhs60_ = default__.safeIndex(341, (d_421_v29_).length(0))
        lhs61_ = d_425_v31_
        lhs62_ = default__.safeIndex(89, (d_425_v31_).length(0))
        lhs63_ = globalState
        lhs64_ = globalState
        lhs59_[lhs60_] = rhs83_
        d_417_v27_ = rhs84_
        lhs61_[lhs62_] = rhs85_
        lhs63_.f7 = rhs86_
        lhs64_.f13 = rhs87_

    def m1(self, p0, p1, globalState):
        d_426_v0_: _dafny.Seq
        d_426_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_427_i0_ in range(default__.abs(-77))])])
        d_428_v1_: D4
        d_428_v1_ = D4_DC11((d_426_v0_) + (_dafny.SeqWithoutIsStrInference([p0])))
        source9_ = d_428_v1_
        if source9_.is_DC12:
            d_429___mcc_h0_ = source9_.cf16
            d_430_cf16_ = d_429___mcc_h0_
            d_431_v2_: _dafny.Array
            def lambda22_(d_432_i1_):
                return default__.safeDivisionInt(d_432_i1_, 344)

            init11_ = lambda22_
            nw59_ = _dafny.Array(None, 28)
            for i0_11_ in range(nw59_.length(0)):
                nw59_[i0_11_] = init11_(i0_11_)
            d_431_v2_ = nw59_
            index58_ = default__.safeIndex(354, (d_431_v2_).length(0))
            (d_431_v2_)[index58_] = d_430_cf16_
            index59_ = default__.safeIndex(354, (d_431_v2_).length(0))
            (d_431_v2_)[index59_] = d_430_cf16_
            d_433_v3_: bool
            d_433_v3_ = False
            if d_433_v3_:
                (globalState).f4 = (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))]
                d_434_v4_: _dafny.Map
                d_434_v4_ = _dafny.Map({(d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))]: (d_430_cf16_) - (d_430_cf16_)})
                d_434_v4_ = d_434_v4_
                d_435_v5_: _dafny.Array
                def lambda23_(d_436_i2_):
                    return D3_DC9(_dafny.CodePoint('v'))

                init12_ = lambda23_
                nw60_ = _dafny.Array(None, 28)
                for i0_12_ in range(nw60_.length(0)):
                    nw60_[i0_12_] = init12_(i0_12_)
                d_435_v5_ = nw60_
                d_437_v6_: str
                d_437_v6_ = _dafny.CodePoint('m')
                d_438_v7_: D3
                d_438_v7_ = D3_DC9(d_437_v6_)
                index60_ = default__.safeIndex(894, (d_435_v5_).length(0))
                (d_435_v5_)[index60_] = (d_438_v7_ if d_433_v3_ else d_438_v7_)
                index61_ = default__.safeIndex(894, (d_435_v5_).length(0))
                (d_435_v5_)[index61_] = d_438_v7_
                rhs88_ = d_431_v2_
                rhs89_ = default__.safeDivisionInt((d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))])
                lhs65_ = globalState
                d_431_v2_ = rhs88_
                lhs65_.f7 = rhs89_
                d_433_v3_ = not(not(not(d_433_v3_)))
            elif True:
                (globalState).f13 = (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))]
                index62_ = default__.safeIndex(354, (d_431_v2_).length(0))
                (d_431_v2_)[index62_] = (0) - ((d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))])
                (globalState).f15 = (0) - ((d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))])
                d_439_v8_: D3
                d_439_v8_ = D3_DC10((d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], (0) - (d_430_cf16_))
                d_440_v9_: _dafny.Map
                d_440_v9_ = _dafny.Map({d_439_v8_: d_433_v3_})
                d_441_v10_: _dafny.MultiSet
                d_441_v10_ = _dafny.MultiSet([len(d_440_v9_)])
                d_442_v11_: _dafny.Seq
                d_442_v11_ = _dafny.SeqWithoutIsStrInference([720])
                d_443_v12_: D1
                d_443_v12_ = D1_DC4(d_441_v10_)
                d_444_v13_: _dafny.Seq
                d_444_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], d_430_cf16_])])
                d_445_v14_: _dafny.Map
                d_445_v14_ = _dafny.Map({(d_441_v10_).cardinality: d_433_v3_})
                d_446_v15_: _dafny.Array
                nw61_ = _dafny.Array(None, 23)
                nw61_[int(0)] = (_dafny.MultiSet([(d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))]])) - (d_441_v10_)
                nw61_[int(1)] = d_441_v10_
                nw61_[int(2)] = _dafny.MultiSet((d_442_v11_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_430_cf16_})) for d_447_i3_ in range(default__.abs(222))])))
                nw61_[int(3)] = d_441_v10_
                nw61_[int(4)] = d_441_v10_
                nw61_[int(5)] = d_441_v10_
                nw61_[int(6)] = d_441_v10_
                nw61_[int(7)] = (d_443_v12_).cf5
                nw61_[int(8)] = d_441_v10_
                nw61_[int(9)] = default__.fm31(d_430_cf16_, (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], globalState)
                nw61_[int(10)] = d_441_v10_
                nw61_[int(11)] = (d_444_v13_)[default__.safeIndex((d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], len(d_444_v13_))]
                nw61_[int(12)] = d_441_v10_
                nw61_[int(13)] = d_441_v10_
                nw61_[int(14)] = _dafny.MultiSet([len(d_445_v14_), (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))]])
                nw61_[int(15)] = (d_441_v10_) - (d_441_v10_)
                nw61_[int(16)] = d_441_v10_
                nw61_[int(17)] = _dafny.MultiSet([(d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))]])
                nw61_[int(18)] = _dafny.MultiSet(d_442_v11_)
                nw61_[int(19)] = d_441_v10_
                nw61_[int(20)] = d_441_v10_
                nw61_[int(21)] = d_441_v10_
                nw61_[int(22)] = _dafny.MultiSet([(d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], (0) - (len(p0)), len(d_442_v11_), 526])
                d_446_v15_ = nw61_
                index63_ = default__.safeIndex(564, (d_446_v15_).length(0))
                (d_446_v15_)[index63_] = (d_441_v10_) - (d_441_v10_)
                index64_ = default__.safeIndex(564, (d_446_v15_).length(0))
                rhs90_ = d_439_v8_
                rhs91_ = d_433_v3_
                rhs92_ = d_441_v10_
                lhs66_ = d_446_v15_
                lhs67_ = default__.safeIndex(564, (d_446_v15_).length(0))
                d_439_v8_ = rhs90_
                d_433_v3_ = rhs91_
                lhs66_[lhs67_] = rhs92_
                (self).m12(d_433_v3_, d_433_v3_, d_431_v2_, (D7_DC22(d_433_v3_, (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))], d_433_v3_, (d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))])).cf35, globalState)
            d_448_v16_: D8
            d_448_v16_ = D8_DC25((d_431_v2_)[default__.safeIndex(354, (d_431_v2_).length(0))])
            d_448_v16_ = d_448_v16_
            d_449_v17_: _dafny.Seq
            d_449_v17_ = _dafny.SeqWithoutIsStrInference([(len(p0)) <= (len(p0))])
            d_433_v3_ = (d_449_v17_)[default__.safeIndex(d_430_cf16_, len(d_449_v17_))]
        elif True:
            d_450___mcc_h1_ = source9_.cf15
            d_451_cf15_ = d_450___mcc_h1_
            d_452_v18_: int
            d_452_v18_ = 757
            (globalState).f7 = (self).fm2(d_452_v18_, globalState)
            d_453_v19_: bool
            d_453_v19_ = False
            d_453_v19_ = (936) == ((0) - ((self).fm2(d_452_v18_, globalState)))
            d_454_v20_: _dafny.Set
            d_454_v20_ = _dafny.Set({d_453_v19_})
            d_455_v21_: _dafny.Seq
            d_455_v21_ = _dafny.SeqWithoutIsStrInference([d_453_v19_])
            d_456_v22_: D1
            d_456_v22_ = D1_DC3(p0, not(False))
            d_457_v23_: _dafny.Array
            nw62_ = _dafny.Array(None, 10)
            nw62_[int(0)] = d_453_v19_
            nw62_[int(1)] = d_453_v19_
            nw62_[int(2)] = d_453_v19_
            nw62_[int(3)] = (d_453_v19_) == (d_453_v19_)
            nw62_[int(4)] = (d_453_v19_ if d_453_v19_ else d_453_v19_)
            nw62_[int(5)] = d_453_v19_
            nw62_[int(6)] = (_dafny.Set({d_453_v19_})).issubset(d_454_v20_)
            nw62_[int(7)] = d_453_v19_
            nw62_[int(8)] = not(((d_455_v21_)[default__.safeIndex(d_452_v18_, len(d_455_v21_))]) == (d_453_v19_))
            nw62_[int(9)] = (d_456_v22_).cf4
            d_457_v23_ = nw62_
            d_457_v23_ = d_457_v23_
            if d_453_v19_:
                d_458_v24_: _dafny.Array
                def lambda24_(d_459_p0_):
                    def lambda25_(d_460_i4_):
                        return (d_459_p0_) + (d_459_p0_)

                    return lambda25_

                init13_ = lambda24_(p0)
                nw63_ = _dafny.Array(None, 10)
                for i0_13_ in range(nw63_.length(0)):
                    nw63_[i0_13_] = init13_(i0_13_)
                d_458_v24_ = nw63_
                d_458_v24_ = d_458_v24_
                d_458_v24_ = d_458_v24_
                d_461_v25_: C0
                nw64_ = C0()
                nw64_.ctor__()
                d_461_v25_ = nw64_
                d_462_v26_: _dafny.Map
                d_462_v26_ = _dafny.Map({default__.safeDivisionInt(d_452_v18_, d_452_v18_): d_452_v18_})
                d_462_v26_ = (d_462_v26_).set((self).fm2(d_452_v18_, globalState), d_452_v18_)
                d_453_v19_ = ((d_455_v21_) + (d_455_v21_)) <= ((default__.fm27(globalState)) + (d_455_v21_))
            elif True:
                (globalState).f13 = (972) * (d_452_v18_)
                d_463_v27_: _dafny.Seq
                d_463_v27_ = _dafny.SeqWithoutIsStrInference([d_452_v18_, len(_dafny.Map({d_453_v19_: 377})), 735])
                d_464_v28_: _dafny.Map
                d_464_v28_ = _dafny.Map({d_463_v27_: d_457_v23_})
                rhs93_ = (d_452_v18_) + (d_452_v18_)
                rhs94_ = d_457_v23_
                rhs95_ = (d_463_v27_) in (d_464_v28_)
                rhs96_ = len((p0) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmislcqxm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtsjc")))))
                lhs68_ = globalState
                lhs69_ = globalState
                lhs68_.f15 = rhs93_
                d_457_v23_ = rhs94_
                d_453_v19_ = rhs95_
                lhs69_.f4 = rhs96_
                (globalState).f13 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_465_i5_ in range(default__.abs(815))]))
                d_466_v29_: C0
                nw65_ = C0()
                nw65_.ctor__()
                d_466_v29_ = nw65_
                d_467_v30_: C1
                nw66_ = C1()
                nw66_.ctor__()
                d_467_v30_ = nw66_
        d_468_v31_: _dafny.Array
        def lambda26_(d_469_i6_):
            return (D8_DC24(len(_dafny.Map({len(_dafny.Map({959: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: _dafny.CodePoint('d')}), _dafny.Map({True: _dafny.CodePoint('d')})]))).cardinality})): (_dafny.MultiSet([True, True])).cardinality})), True)).cf41

        init14_ = lambda26_
        nw67_ = _dafny.Array(None, 13)
        for i0_14_ in range(nw67_.length(0)):
            nw67_[i0_14_] = init14_(i0_14_)
        d_468_v31_ = nw67_
        d_470_v32_: bool
        d_470_v32_ = False
        d_471_v33_: D6
        d_471_v33_ = D6_DC17(d_470_v32_, False, True)
        d_472_v34_: _dafny.Seq
        d_472_v34_ = _dafny.SeqWithoutIsStrInference([d_470_v32_, d_470_v32_, (d_471_v33_).cf26])
        d_473_v35_: int
        d_473_v35_ = 348
        d_474_v36_: _dafny.Map
        d_474_v36_ = _dafny.Map({True: d_473_v35_})
        d_475_v37_: _dafny.Map
        d_475_v37_ = _dafny.Map({(d_472_v34_)[default__.safeIndex(len(d_474_v36_), len(d_472_v34_))]: d_468_v31_})
        d_476_v38_: _dafny.Map
        d_476_v38_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_468_v31_, ((d_475_v37_)[d_470_v32_] if (d_470_v32_) in (d_475_v37_) else d_468_v31_)]): d_473_v35_})
        d_477_v39_: _dafny.MultiSet
        d_477_v39_ = _dafny.MultiSet([(self).fm8((_dafny.MultiSet(default__.fm27(globalState))).cardinality, d_470_v32_, d_473_v35_, globalState), not(False)])
        (globalState).f15 = (0) - ((0) - (((d_476_v38_)[_dafny.SeqWithoutIsStrInference([d_468_v31_])] if (_dafny.SeqWithoutIsStrInference([d_468_v31_])) in (d_476_v38_) else default__.safeDivisionInt((d_477_v39_).cardinality, d_473_v35_))))
        (globalState).f7 = ((_dafny.MultiSet([not(d_470_v32_)])).cardinality) + ((_dafny.MultiSet(p0)).cardinality)
        d_478_v40_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Array(None, 0), 1)
        d_478_v40_ = nw68_
        d_479_v41_: _dafny.Map
        d_479_v41_ = _dafny.Map({d_473_v35_: d_468_v31_})
        d_480_v42_: _dafny.Array
        nw69_ = _dafny.Array(None, 29)
        nw69_[int(0)] = d_468_v31_
        nw69_[int(1)] = d_468_v31_
        nw69_[int(2)] = d_468_v31_
        nw69_[int(3)] = d_468_v31_
        nw69_[int(4)] = d_468_v31_
        nw69_[int(5)] = d_468_v31_
        nw69_[int(6)] = d_468_v31_
        nw69_[int(7)] = d_468_v31_
        nw69_[int(8)] = d_468_v31_
        nw69_[int(9)] = d_468_v31_
        nw69_[int(10)] = d_468_v31_
        nw69_[int(11)] = d_468_v31_
        nw69_[int(12)] = d_468_v31_
        nw69_[int(13)] = d_468_v31_
        nw69_[int(14)] = d_468_v31_
        nw69_[int(15)] = d_468_v31_
        nw69_[int(16)] = d_468_v31_
        nw69_[int(17)] = d_468_v31_
        nw69_[int(18)] = d_468_v31_
        nw69_[int(19)] = d_468_v31_
        nw69_[int(20)] = d_468_v31_
        nw69_[int(21)] = d_468_v31_
        nw69_[int(22)] = d_468_v31_
        nw69_[int(23)] = d_468_v31_
        nw69_[int(24)] = d_468_v31_
        nw69_[int(25)] = d_468_v31_
        nw69_[int(26)] = d_468_v31_
        nw69_[int(27)] = d_468_v31_
        nw69_[int(28)] = ((d_479_v41_)[-38] if (-38) in (d_479_v41_) else d_468_v31_)
        d_480_v42_ = nw69_
        d_481_v43_: D0
        d_481_v43_ = D0_DC1(d_468_v31_)
        index65_ = default__.safeIndex(85, (d_480_v42_).length(0))
        (d_480_v42_)[index65_] = (d_481_v43_).cf1
        d_482_v44_: _dafny.Array
        def lambda27_(d_483_v39_):
            def lambda28_(d_484_i7_):
                return d_483_v39_

            return lambda28_

        init15_ = lambda27_(d_477_v39_)
        nw70_ = _dafny.Array(None, 29)
        for i0_15_ in range(nw70_.length(0)):
            nw70_[i0_15_] = init15_(i0_15_)
        d_482_v44_ = nw70_
        index66_ = default__.safeIndex(279, (d_482_v44_).length(0))
        (d_482_v44_)[index66_] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_470_v32_])) + (d_472_v34_))
        d_485_v45_: _dafny.Set
        d_485_v45_ = _dafny.Set({d_470_v32_, d_470_v32_, False, d_470_v32_, d_470_v32_})
        d_486_v46_: _dafny.Seq
        d_486_v46_ = _dafny.SeqWithoutIsStrInference([d_478_v40_, d_478_v40_, d_478_v40_, d_478_v40_, d_478_v40_])
        index67_ = default__.safeIndex(85, (d_480_v42_).length(0))
        index68_ = default__.safeIndex(279, (d_482_v44_).length(0))
        rhs97_ = (d_486_v46_)[default__.safeIndex(d_473_v35_, len(d_486_v46_))]
        rhs98_ = d_468_v31_
        rhs99_ = _dafny.MultiSet([d_470_v32_, d_470_v32_])
        rhs100_ = d_485_v45_
        lhs70_ = d_480_v42_
        lhs71_ = default__.safeIndex(85, (d_480_v42_).length(0))
        lhs72_ = d_482_v44_
        lhs73_ = default__.safeIndex(279, (d_482_v44_).length(0))
        d_478_v40_ = rhs97_
        lhs70_[lhs71_] = rhs98_
        lhs72_[lhs73_] = rhs99_
        d_485_v45_ = rhs100_
        pat_let_tv3_ = p0
        pat_let_tv4_ = p0
        pat_let_tv5_ = d_472_v34_
        pat_let_tv6_ = d_473_v35_
        def lambda29_(source10_):
            if source10_.is_DC12:
                d_487___mcc_h2_ = source10_.cf16
                d_488_cf16_ = d_487___mcc_h2_
                return (_dafny.MultiSet([len(_dafny.Set({_dafny.MultiSet([_dafny.CodePoint('y'), _dafny.CodePoint('f')]), _dafny.MultiSet(pat_let_tv3_), _dafny.MultiSet(pat_let_tv4_)})), len(pat_let_tv5_), d_488_cf16_, 561, (0) - (-846)])).cardinality
            elif True:
                d_489___mcc_h3_ = source10_.cf15
                d_490_cf15_ = d_489___mcc_h3_
                return pat_let_tv6_

        (globalState).f15 = lambda29_(d_428_v1_)
        d_491_v47_: _dafny.Array
        nw71_ = _dafny.Array(None, 20)
        nw71_[int(0)] = len(d_472_v34_)
        nw71_[int(1)] = d_473_v35_
        nw71_[int(2)] = d_473_v35_
        nw71_[int(3)] = (self).fm2(d_473_v35_, globalState)
        nw71_[int(4)] = d_473_v35_
        nw71_[int(5)] = len(p0)
        nw71_[int(6)] = d_473_v35_
        nw71_[int(7)] = d_473_v35_
        nw71_[int(8)] = d_473_v35_
        nw71_[int(9)] = d_473_v35_
        nw71_[int(10)] = d_473_v35_
        nw71_[int(11)] = len(d_472_v34_)
        nw71_[int(12)] = d_473_v35_
        nw71_[int(13)] = d_473_v35_
        nw71_[int(14)] = d_473_v35_
        nw71_[int(15)] = d_473_v35_
        nw71_[int(16)] = 0
        nw71_[int(17)] = d_473_v35_
        nw71_[int(18)] = -971
        nw71_[int(19)] = (0) - (((d_477_v39_)[d_470_v32_] if (d_470_v32_) in (d_477_v39_) else -672))
        d_491_v47_ = nw71_
        d_492_v48_: _dafny.Map
        d_492_v48_ = _dafny.Map({d_491_v47_: True})
        d_492_v48_ = (d_492_v48_).set(d_491_v47_, d_470_v32_)

    def m11(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_493_v0_: int
        d_493_v0_ = 978
        d_494_v1_: _dafny.Set
        d_494_v1_ = _dafny.Set({d_493_v0_, d_493_v0_})
        if (d_494_v1_) == (_dafny.Set({d_493_v0_})):
            d_495_v2_: _dafny.Array
            def lambda30_(d_496_v0_):
                def lambda31_(d_497_i0_):
                    return (d_497_i0_) * (d_496_v0_)

                return lambda31_

            init16_ = lambda30_(d_493_v0_)
            nw72_ = _dafny.Array(None, 2)
            for i0_16_ in range(nw72_.length(0)):
                nw72_[i0_16_] = init16_(i0_16_)
            d_495_v2_ = nw72_
            index69_ = default__.safeIndex(436, (d_495_v2_).length(0))
            (d_495_v2_)[index69_] = len(_dafny.SeqWithoutIsStrInference([d_495_v2_, d_495_v2_, d_495_v2_]))
            index70_ = default__.safeIndex(436, (d_495_v2_).length(0))
            (d_495_v2_)[index70_] = (d_493_v0_) * (d_493_v0_)
            (globalState).f7 = (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]
            d_498_v3_: D2
            d_498_v3_ = D2_DC7(p3, default__.fm34(d_493_v0_, d_493_v0_, p3, globalState), d_495_v2_)
            source11_ = d_498_v3_
            if source11_.is_DC6:
                d_499___mcc_h0_ = source11_.cf7
                d_500_cf7_ = d_499___mcc_h0_
                d_501_v4_: bool
                d_501_v4_ = False
                d_501_v4_ = not((self).fm8(d_500_cf7_, p3, (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))], globalState))
                d_502_v5_: C1
                nw73_ = C1()
                nw73_.ctor__()
                d_502_v5_ = nw73_
                d_503_v6_: _dafny.Map
                d_503_v6_ = _dafny.Map({d_494_v1_: d_493_v0_})
                d_503_v6_ = d_503_v6_
                d_501_v4_ = (d_501_v4_ if d_501_v4_ else d_501_v4_)
            elif source11_.is_DC7:
                d_504___mcc_h1_ = source11_.cf8
                d_505___mcc_h2_ = source11_.cf9
                d_506___mcc_h3_ = source11_.cf10
                d_507_cf10_ = d_506___mcc_h3_
                d_508_cf9_ = d_505___mcc_h2_
                d_509_cf8_ = d_504___mcc_h1_
                d_510_v7_: str
                d_510_v7_ = _dafny.CodePoint('q')
                d_511_v8_: _dafny.MultiSet
                d_511_v8_ = _dafny.MultiSet([d_493_v0_])
                d_512_v9_: _dafny.Map
                d_512_v9_ = _dafny.Map({d_493_v0_: (_dafny.Set({((d_511_v8_)[d_493_v0_] if (d_493_v0_) in (d_511_v8_) else 286), len((default__.fm35(globalState)).set((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))], d_509_cf8_))})) | (d_494_v1_)})
                d_513_v10_: _dafny.Array
                nw74_ = _dafny.Array(False, 22)
                d_513_v10_ = nw74_
                d_514_v11_: _dafny.Map
                d_514_v11_ = _dafny.Map({d_493_v0_: p0})
                d_515_v12_: _dafny.Seq
                d_515_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jf"))
                d_516_v13_: _dafny.Map
                d_516_v13_ = _dafny.Map({(d_511_v8_).set(len(d_515_v12_), default__.abs((0) - ((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]))): d_510_v7_})
                d_517_v14_: _dafny.Map
                d_517_v14_ = _dafny.Map({len(d_515_v12_): p3})
                rhs101_ = (((d_516_v13_)[d_511_v8_] if (d_511_v8_) in (d_516_v13_) else _dafny.CodePoint('j')) if p3 else d_510_v7_)
                rhs102_ = d_512_v9_
                rhs103_ = p0
                rhs104_ = ((_dafny.Map({len(d_517_v14_): p0})) | (d_514_v11_)) | (d_514_v11_)
                rhs105_ = (p3 if d_509_cf8_ else not(p3))
                d_510_v7_ = rhs101_
                d_512_v9_ = rhs102_
                d_513_v10_ = rhs103_
                d_514_v11_ = rhs104_
                d_509_cf8_ = rhs105_
                d_517_v14_ = (d_517_v14_).set(d_493_v0_, (p1)[default__.safeIndex(519, len(p1))])
                d_509_cf8_ = True
                d_518_v15_: _dafny.Set
                d_518_v15_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfqn"))})
                rhs106_ = (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]
                rhs107_ = d_518_v15_
                lhs74_ = globalState
                lhs74_.f4 = rhs106_
                d_518_v15_ = rhs107_
            elif source11_.is_DC5:
                d_519___mcc_h4_ = source11_.cf6
                d_520_cf6_ = d_519___mcc_h4_
                d_521_v16_: _dafny.Array
                nw75_ = _dafny.Array(_dafny.Map({}), 5)
                d_521_v16_ = nw75_
                d_522_v17_: _dafny.Map
                d_522_v17_ = _dafny.Map({p3: p2})
                index71_ = default__.safeIndex(778, (d_521_v16_).length(0))
                (d_521_v16_)[index71_] = (d_522_v17_) | (_dafny.Map({p3: p2}))
                index72_ = default__.safeIndex(778, (d_521_v16_).length(0))
                (d_521_v16_)[index72_] = (d_522_v17_) | ((d_522_v17_) | (default__.fm36((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))], globalState)))
                d_520_cf6_ = ((d_498_v3_).cf10 if p3 else (d_520_cf6_ if p3 else d_520_cf6_))
                d_523_v18_: _dafny.Set
                d_523_v18_ = _dafny.Set({d_495_v2_, d_520_cf6_, d_520_cf6_, d_520_cf6_})
                index73_ = default__.safeIndex(300, (p0).length(0))
                (p0)[index73_] = (d_523_v18_).ispropersubset(d_523_v18_)
                index74_ = default__.safeIndex(300, (p0).length(0))
                rhs108_ = ((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]) - (default__.safeModuloInt((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))], -246))
                rhs109_ = p3
                rhs110_ = (0) - (d_493_v0_)
                lhs75_ = p0
                lhs76_ = default__.safeIndex(300, (p0).length(0))
                d_493_v0_ = rhs108_
                lhs75_[lhs76_] = rhs109_
                d_493_v0_ = rhs110_
                d_524_v19_: D2
                d_524_v19_ = D2_DC5(d_495_v2_)
                d_525_v20_: _dafny.Map
                d_525_v20_ = _dafny.Map({d_524_v19_: p3})
                d_526_v21_: _dafny.Array
                nw76_ = _dafny.Array(None, 2)
                nw76_[int(0)] = d_525_v20_
                nw76_[int(1)] = _dafny.Map({d_524_v19_: p3})
                d_526_v21_ = nw76_
                d_527_v22_: _dafny.Map
                d_527_v22_ = _dafny.Map({d_524_v19_: True})
                index75_ = default__.safeIndex(91, (d_526_v21_).length(0))
                (d_526_v21_)[index75_] = d_527_v22_
                d_528_v23_: _dafny.Seq
                d_528_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "racjbb"))
                d_529_v24_: D1
                d_529_v24_ = D1_DC3(d_528_v23_, p3)
                index76_ = default__.safeIndex(91, (d_526_v21_).length(0))
                rhs111_ = d_527_v22_
                rhs112_ = (d_529_v24_ if (p0)[default__.safeIndex(300, (p0).length(0))] else d_529_v24_)
                rhs113_ = (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]
                lhs77_ = d_526_v21_
                lhs78_ = default__.safeIndex(91, (d_526_v21_).length(0))
                lhs79_ = globalState
                lhs77_[lhs78_] = rhs111_
                d_529_v24_ = rhs112_
                lhs79_.f13 = rhs113_
            elif True:
                d_530___mcc_h5_ = source11_.cf11
                d_531_cf11_ = d_530___mcc_h5_
                d_532_v25_: _dafny.Map
                d_532_v25_ = _dafny.Map({p3: (0) - (d_493_v0_)})
                d_533_v26_: _dafny.Seq
                d_533_v26_ = _dafny.SeqWithoutIsStrInference([331, (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]])
                d_534_v27_: _dafny.MultiSet
                d_534_v27_ = _dafny.MultiSet([(d_533_v26_)[default__.safeIndex((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))], len(d_533_v26_))], 86])
                d_535_v28_: _dafny.Seq
                d_535_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([((d_532_v25_)[p3] if (p3) in (d_532_v25_) else (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))])]), d_534_v27_, d_534_v27_])
                d_535_v28_ = d_535_v28_
                d_536_v29_: _dafny.Map
                d_536_v29_ = _dafny.Map({d_493_v0_: p3})
                index77_ = default__.safeIndex(436, (d_495_v2_).length(0))
                (d_495_v2_)[index77_] = (0) - ((0) - ((default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([len(d_536_v29_)])), d_493_v0_)) * (((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))]) * (d_493_v0_))))
                (globalState).f13 = default__.safeDivisionInt(d_493_v0_, (d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))])
                index78_ = default__.safeIndex(129, (p0).length(0))
                (p0)[index78_] = default__.fm26(globalState)
                index79_ = default__.safeIndex(129, (p0).length(0))
                (p0)[index79_] = False
            index80_ = default__.safeIndex(436, (d_495_v2_).length(0))
            (d_495_v2_)[index80_] = default__.safeDivisionInt(d_493_v0_, d_493_v0_)
            d_537_v30_: _dafny.Seq
            d_537_v30_ = _dafny.SeqWithoutIsStrInference([d_495_v2_])
            d_538_v31_: _dafny.Seq
            d_538_v31_ = _dafny.SeqWithoutIsStrInference([(d_537_v30_)[default__.safeIndex((d_495_v2_)[default__.safeIndex(436, (d_495_v2_).length(0))], len(d_537_v30_))], d_495_v2_, d_495_v2_])
            d_538_v31_ = d_538_v31_
        elif True:
            d_539_v32_: _dafny.Array
            def lambda32_(d_540_p1_, d_541_v0_):
                def lambda33_(d_542_i1_):
                    return (d_540_p1_)[default__.safeIndex(d_541_v0_, len(d_540_p1_))]

                return lambda33_

            init17_ = lambda32_(p1, d_493_v0_)
            nw77_ = _dafny.Array(None, 24)
            for i0_17_ in range(nw77_.length(0)):
                nw77_[i0_17_] = init17_(i0_17_)
            d_539_v32_ = nw77_
            d_543_v33_: _dafny.Array
            nw78_ = _dafny.Array(_dafny.Seq({}), 24)
            d_543_v33_ = nw78_
            rhs114_ = -932
            rhs115_ = d_493_v0_
            rhs116_ = d_539_v32_
            rhs117_ = d_543_v33_
            lhs80_ = globalState
            lhs81_ = globalState
            lhs80_.f15 = rhs114_
            lhs81_.f7 = rhs115_
            d_539_v32_ = rhs116_
            d_543_v33_ = rhs117_
            d_544_v34_: bool
            d_544_v34_ = False
            d_544_v34_ = p3
            d_545_v35_: C1
            nw79_ = C1()
            nw79_.ctor__()
            d_545_v35_ = nw79_
            d_546_v36_: _dafny.Seq
            d_546_v36_ = _dafny.SeqWithoutIsStrInference([d_493_v0_, d_493_v0_])
            d_544_v34_ = (d_493_v0_) >= (((_dafny.MultiSet((d_546_v36_).set(default__.safeIndex(d_493_v0_, len(d_546_v36_)), d_493_v0_))).cardinality if p3 else d_493_v0_))
            d_544_v34_ = p3
        d_547_v37_: _dafny.Array
        nw80_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_547_v37_ = nw80_
        d_548_v38_: bool
        d_548_v38_ = False
        rhs118_ = (784) + (46)
        rhs119_ = d_547_v37_
        rhs120_ = not(p3)
        rhs121_ = d_548_v38_
        rhs122_ = d_548_v38_
        lhs82_ = globalState
        lhs82_.f13 = rhs118_
        d_547_v37_ = rhs119_
        d_548_v38_ = rhs120_
        d_548_v38_ = rhs121_
        d_548_v38_ = rhs122_
        hi3_ = -680
        for d_549_i2_ in range(d_493_v0_, hi3_):
            d_550_v39_: D0
            d_550_v39_ = D0_DC1(p0)
            d_551_v40_: _dafny.MultiSet
            d_551_v40_ = _dafny.MultiSet([d_550_v39_, d_550_v39_, d_550_v39_, d_550_v39_])
            d_552_v41_: _dafny.Map
            d_552_v41_ = _dafny.Map({d_549_i2_: d_551_v40_})
            d_553_v42_: _dafny.Seq
            d_553_v42_ = _dafny.SeqWithoutIsStrInference([d_550_v39_])
            d_552_v41_ = (d_552_v41_).set(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayhpbe"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_554_i3_ in range(default__.abs(613))]))), _dafny.MultiSet(d_553_v42_))
            d_555_v43_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 2)
            d_555_v43_ = nw81_
            index81_ = default__.safeIndex(370, (d_555_v43_).length(0))
            (d_555_v43_)[index81_] = d_549_i2_
            index82_ = default__.safeIndex(370, (d_555_v43_).length(0))
            (d_555_v43_)[index82_] = d_549_i2_
            d_548_v38_ = (d_549_i2_) == (d_549_i2_)
            if p3:
                d_556_v44_: _dafny.Seq
                d_556_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bypilo"))
                rhs123_ = (d_556_v44_) + (d_556_v44_)
                rhs124_ = (0) - ((d_555_v43_)[default__.safeIndex(370, (d_555_v43_).length(0))])
                lhs83_ = globalState
                d_556_v44_ = rhs123_
                lhs83_.f15 = rhs124_
                d_557_v45_: C0
                nw82_ = C0()
                nw82_.ctor__()
                d_557_v45_ = nw82_
                d_548_v38_ = ((d_555_v43_)[default__.safeIndex(370, (d_555_v43_).length(0))]) != (d_493_v0_)
                d_558_v46_: _dafny.Map
                d_558_v46_ = _dafny.Map({not(p3): p3})
                d_559_v47_: _dafny.MultiSet
                d_559_v47_ = _dafny.MultiSet([p3, not(d_548_v38_)])
                d_560_v48_: _dafny.MultiSet
                d_560_v48_ = _dafny.MultiSet([(d_559_v47_).cardinality])
                d_558_v46_ = (d_558_v46_).set(d_548_v38_, (d_560_v48_).isdisjoint(d_560_v48_))
                d_561_v49_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.Seq({}), 7)
                d_561_v49_ = nw83_
                d_561_v49_ = d_561_v49_
            elif True:
                d_562_v50_: str
                d_562_v50_ = _dafny.CodePoint('s')
                d_563_v51_: _dafny.Map
                d_563_v51_ = _dafny.Map({d_562_v50_: p0})
                (globalState).f4 = len(d_563_v51_)
                d_564_v52_: _dafny.Seq
                d_564_v52_ = _dafny.SeqWithoutIsStrInference([d_548_v38_, d_548_v38_, not(d_548_v38_)])
                d_564_v52_ = default__.fm27(globalState)
                d_565_v53_: _dafny.Array
                def lambda34_(d_566_p3_):
                    def lambda35_(d_567_i4_):
                        return _dafny.SeqWithoutIsStrInference([d_566_p3_, d_566_p3_, not(True)])

                    return lambda35_

                init18_ = lambda34_(p3)
                nw84_ = _dafny.Array(None, 15)
                for i0_18_ in range(nw84_.length(0)):
                    nw84_[i0_18_] = init18_(i0_18_)
                d_565_v53_ = nw84_
                d_568_v54_: _dafny.Map
                d_568_v54_ = _dafny.Map({d_565_v53_: d_548_v38_})
                d_569_v55_: _dafny.MultiSet
                d_569_v55_ = _dafny.MultiSet([d_568_v54_, d_568_v54_, d_568_v54_])
                d_548_v38_ = not((d_569_v55_).ispropersubset(d_569_v55_))
                d_570_v56_: _dafny.Seq
                d_570_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcbscyt"))
                d_571_v57_: _dafny.Array
                nw85_ = _dafny.Array(None, 3)
                nw85_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfaydljs"))
                nw85_[int(1)] = d_570_v56_
                nw85_[int(2)] = (d_570_v56_) + (d_570_v56_)
                d_571_v57_ = nw85_
                index83_ = default__.safeIndex(994, (d_571_v57_).length(0))
                (d_571_v57_)[index83_] = (d_570_v56_) + (d_570_v56_)
                d_572_v58_: _dafny.Seq
                d_572_v58_ = _dafny.SeqWithoutIsStrInference([d_570_v56_])
                index84_ = default__.safeIndex(994, (d_571_v57_).length(0))
                rhs125_ = d_548_v38_
                rhs126_ = (d_570_v56_ if ((0) - ((d_555_v43_)[default__.safeIndex(370, (d_555_v43_).length(0))])) >= ((d_555_v43_)[default__.safeIndex(370, (d_555_v43_).length(0))]) else _dafny.SeqWithoutIsStrInference([d_562_v50_ for d_573_i5_ in range(default__.abs(352))]))
                rhs127_ = (((d_572_v58_).set(default__.safeIndex((d_555_v43_)[default__.safeIndex(370, (d_555_v43_).length(0))], len(d_572_v58_)), d_570_v56_)) + (d_572_v58_)) + (d_572_v58_)
                rhs128_ = not (p3) or (not(d_548_v38_))
                lhs84_ = d_571_v57_
                lhs85_ = default__.safeIndex(994, (d_571_v57_).length(0))
                d_548_v38_ = rhs125_
                lhs84_[lhs85_] = rhs126_
                d_572_v58_ = rhs127_
                d_548_v38_ = rhs128_
                d_574_v59_: C1
                nw86_ = C1()
                nw86_.ctor__()
                d_574_v59_ = nw86_
        d_548_v38_ = not (not(d_548_v38_)) or (p3)
        d_575_v60_: D6
        d_575_v60_ = D6_DC17(d_548_v38_, not(p3), d_548_v38_)
        d_576_i6_: int
        d_576_i6_ = 0
        with _dafny.label("2"):
            while not (d_548_v38_) or ((d_575_v60_).cf25):
                with _dafny.c_label("2"):
                    if (d_576_i6_) >= (100):
                        raise _dafny.Break("2")
                    d_576_i6_ = (d_576_i6_) + (1)
                    d_577_v61_: _dafny.Seq
                    d_577_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfjnh"))
                    d_577_v61_ = d_577_v61_
                    d_548_v38_ = (default__.safeModuloInt(d_493_v0_, d_493_v0_)) >= (d_493_v0_)
                    d_578_v62_: _dafny.Array
                    def lambda36_(d_579_v0_):
                        def lambda37_(d_580_i7_):
                            return (d_580_i7_) * (d_579_v0_)

                        return lambda37_

                    init19_ = lambda36_(d_493_v0_)
                    nw87_ = _dafny.Array(None, 8)
                    for i0_19_ in range(nw87_.length(0)):
                        nw87_[i0_19_] = init19_(i0_19_)
                    d_578_v62_ = nw87_
                    index85_ = default__.safeIndex(420, (d_578_v62_).length(0))
                    (d_578_v62_)[index85_] = (0) - ((len(_dafny.Set({d_493_v0_}))) + (d_493_v0_))
                    index86_ = default__.safeIndex(420, (d_578_v62_).length(0))
                    (d_578_v62_)[index86_] = (0) - (d_493_v0_)
                    d_581_v63_: _dafny.Map
                    d_581_v63_ = _dafny.Map({True: _dafny.Set({(p1)[default__.safeIndex((d_578_v62_)[default__.safeIndex(420, (d_578_v62_).length(0))], len(p1))]})})
                    d_582_v64_: _dafny.Map
                    d_582_v64_ = _dafny.Map({(d_578_v62_)[default__.safeIndex(420, (d_578_v62_).length(0))]: _dafny.Set({not(p3)})})
                    d_583_v65_: _dafny.Array
                    nw88_ = _dafny.Array(None, 21)
                    nw88_[int(0)] = (p2) | (p2)
                    nw88_[int(1)] = (p2) - (p2)
                    nw88_[int(2)] = p2
                    nw88_[int(3)] = p2
                    nw88_[int(4)] = p2
                    nw88_[int(5)] = (p2) - (default__.fm37(globalState))
                    nw88_[int(6)] = p2
                    nw88_[int(7)] = (_dafny.Set({not(p3)})) - (p2)
                    nw88_[int(8)] = (p2) - (default__.fm37(globalState))
                    nw88_[int(9)] = p2
                    nw88_[int(10)] = p2
                    nw88_[int(11)] = (p2) - (p2)
                    nw88_[int(12)] = p2
                    nw88_[int(13)] = (p2) - (((d_581_v63_)[d_548_v38_] if (d_548_v38_) in (d_581_v63_) else _dafny.Set({True})))
                    nw88_[int(14)] = p2
                    nw88_[int(15)] = p2
                    nw88_[int(16)] = p2
                    nw88_[int(17)] = (p2 if d_548_v38_ else p2)
                    nw88_[int(18)] = (p2).intersection(p2)
                    nw88_[int(19)] = (p2) | (p2)
                    nw88_[int(20)] = ((d_582_v64_)[d_493_v0_] if (d_493_v0_) in (d_582_v64_) else p2)
                    d_583_v65_ = nw88_
                    index87_ = default__.safeIndex(438, (d_583_v65_).length(0))
                    (d_583_v65_)[index87_] = p2
                    index88_ = default__.safeIndex(438, (d_583_v65_).length(0))
                    (d_583_v65_)[index88_] = _dafny.Set({d_548_v38_})
                    pass
            pass
        (globalState).f15 = d_493_v0_
        r0 = (786 if d_548_v38_ else 111)
        d_584_v66_: D4
        d_584_v66_ = D4_DC11(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpuhuc"))]))
        pat_let_tv7_ = d_493_v0_
        def lambda38_(source12_):
            if source12_.is_DC12:
                d_585___mcc_h6_ = source12_.cf16
                d_586_cf16_ = d_585___mcc_h6_
                return d_586_cf16_
            elif True:
                d_587___mcc_h7_ = source12_.cf15
                d_588_cf15_ = d_587___mcc_h7_
                return (0) - (default__.safeDivisionInt(204, pat_let_tv7_))

        r1 = lambda38_(d_584_v66_)
        r2 = (0) - ((d_493_v0_) * (798))
        return r0, r1, r2

    def m12(self, p0, p1, p2, p3, globalState):
        d_589_v0_: _dafny.Set
        d_589_v0_ = _dafny.Set({default__.fm38(p1, globalState)})
        d_590_v1_: str
        d_590_v1_ = _dafny.CodePoint('q')
        d_591_v2_: _dafny.Seq
        d_591_v2_ = _dafny.SeqWithoutIsStrInference([d_590_v1_])
        d_592_v4_: _dafny.Seq
        def iife34_():
            coll30_ = _dafny.Set()
            compr_30_: str
            for compr_30_ in (d_591_v2_).Elements:
                d_593_v3_: str = compr_30_
                if (d_593_v3_) in (d_591_v2_):
                    coll30_ = coll30_.union(_dafny.Set([d_593_v3_]))
            return _dafny.Set(coll30_)
        d_592_v4_ = _dafny.SeqWithoutIsStrInference([iife34_()
        , d_589_v0_])
        d_594_i0_: int
        d_594_i0_ = 0
        with _dafny.label("3"):
            while (_dafny.SeqWithoutIsStrInference([d_589_v0_])) == ((_dafny.SeqWithoutIsStrInference([d_589_v0_ for d_651_i1_ in range(default__.abs(221))])) + (d_592_v4_)):
                with _dafny.c_label("3"):
                    if (d_594_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_594_i0_ = (d_594_i0_) + (1)
                    d_595_v5_: int
                    d_595_v5_ = 813
                    (globalState).f13 = ((72) + (d_595_v5_) if p3 else d_595_v5_)
                    index89_ = default__.safeIndex(561, (p2).length(0))
                    (p2)[index89_] = d_595_v5_
                    index90_ = default__.safeIndex(561, (p2).length(0))
                    def iife35_():
                        coll31_ = _dafny.Set()
                        compr_31_: int
                        for compr_31_ in _dafny.IntegerRange(218, 157):
                            d_596_v6_: int = compr_31_
                            if ((218) <= (d_596_v6_)) and ((d_596_v6_) < (157)):
                                coll31_ = coll31_.union(_dafny.Set([(d_596_v6_) + ((_dafny.MultiSet([p1, p0])).cardinality)]))
                        return _dafny.Set(coll31_)
                    (p2)[index90_] = len(iife35_()
                    )
                    d_597_v7_: D2
                    d_597_v7_ = D2_DC6(d_595_v5_)
                    d_598_v8_: D2
                    d_598_v8_ = D2_DC8(d_597_v7_)
                    source13_ = d_598_v8_
                    if source13_.is_DC6:
                        d_599___mcc_h0_ = source13_.cf7
                        d_600_cf7_ = d_599___mcc_h0_
                        d_601_v9_: bool
                        d_601_v9_ = True
                        d_602_v10_: _dafny.Map
                        d_602_v10_ = _dafny.Map({d_595_v5_: (0) - (-213)})
                        d_601_v9_ = (self).fm24(d_595_v5_, p0, d_602_v10_, d_600_cf7_, globalState)
                        d_601_v9_ = p3
                        d_603_v11_: _dafny.Array
                        nw89_ = _dafny.Array(_dafny.Seq({}), 24)
                        d_603_v11_ = nw89_
                        d_604_v12_: _dafny.Seq
                        d_604_v12_ = _dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex(561, (p2).length(0))]])
                        index91_ = default__.safeIndex(348, (d_603_v11_).length(0))
                        (d_603_v11_)[index91_] = d_604_v12_
                        d_605_v13_: _dafny.Map
                        d_605_v13_ = _dafny.Map({(p2)[default__.safeIndex(561, (p2).length(0))]: p0})
                        index92_ = default__.safeIndex(348, (d_603_v11_).length(0))
                        (d_603_v11_)[index92_] = ((_dafny.SeqWithoutIsStrInference([d_600_cf7_ for d_606_i2_ in range(default__.abs(428))])) + ((d_604_v12_ if p0 else d_604_v12_))).set(default__.safeIndex((p2)[default__.safeIndex(561, (p2).length(0))], len((_dafny.SeqWithoutIsStrInference([d_600_cf7_ for d_607_i2_ in range(default__.abs(428))])) + ((d_604_v12_ if p0 else d_604_v12_)))), len(d_605_v13_))
                        d_608_v14_: _dafny.Array
                        nw90_ = _dafny.Array(False, 26)
                        d_608_v14_ = nw90_
                        d_609_v15_: D0
                        d_609_v15_ = D0_DC1(d_608_v14_)
                        d_610_v16_: _dafny.Set
                        d_610_v16_ = _dafny.Set({(d_609_v15_).cf1, d_608_v14_})
                        rhs129_ = (d_610_v16_) | (d_610_v16_)
                        rhs130_ = ((p2)[default__.safeIndex(561, (p2).length(0))]) - (185)
                        rhs131_ = ((0) - (d_600_cf7_)) >= (default__.safeModuloInt(d_600_cf7_, d_600_cf7_))
                        rhs132_ = (d_591_v2_) < ((d_591_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqk"))))
                        d_610_v16_ = rhs129_
                        d_595_v5_ = rhs130_
                        d_601_v9_ = rhs131_
                        d_601_v9_ = rhs132_
                    elif source13_.is_DC7:
                        d_611___mcc_h1_ = source13_.cf8
                        d_612___mcc_h2_ = source13_.cf9
                        d_613___mcc_h3_ = source13_.cf10
                        d_614_cf10_ = d_613___mcc_h3_
                        d_615_cf9_ = d_612___mcc_h2_
                        d_616_cf8_ = d_611___mcc_h1_
                        (globalState).f13 = (0) - ((p2)[default__.safeIndex(561, (p2).length(0))])
                        d_617_v17_: _dafny.Array
                        def lambda39_(d_618_p1_):
                            def lambda40_(d_619_i3_):
                                return d_618_p1_

                            return lambda40_

                        init20_ = lambda39_(p1)
                        nw91_ = _dafny.Array(None, 13)
                        for i0_20_ in range(nw91_.length(0)):
                            nw91_[i0_20_] = init20_(i0_20_)
                        d_617_v17_ = nw91_
                        d_620_v18_: _dafny.MultiSet
                        d_620_v18_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_590_v1_ for d_621_i4_ in range(default__.abs(114))]), d_591_v2_])
                        d_622_v19_: _dafny.Set
                        d_622_v19_ = _dafny.Set({p3})
                        d_623_v20_: _dafny.Seq
                        d_623_v20_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_624_v21_: D1
                        d_624_v21_ = D1_DC2(len(d_623_v20_))
                        d_625_v22_: _dafny.Map
                        d_625_v22_ = _dafny.Map({(p2)[default__.safeIndex(561, (p2).length(0))]: d_624_v21_})
                        index93_ = default__.safeIndex(121, (d_617_v17_).length(0))
                        (d_617_v17_)[index93_] = not((d_620_v18_) != (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqrdii")), default__.fm30(len(d_622_v19_), p1, _dafny.Set({d_595_v5_, len(d_625_v22_), d_595_v5_}), p3, globalState), _dafny.SeqWithoutIsStrInference([d_590_v1_ for d_626_i5_ in range(default__.abs(-842))]), d_591_v2_])))
                        d_627_v23_: _dafny.Map
                        d_627_v23_ = _dafny.Map({p3: p0})
                        index94_ = default__.safeIndex(121, (d_617_v17_).length(0))
                        (d_617_v17_)[index94_] = ((d_627_v23_)[p0] if (p0) in (d_627_v23_) else p1)
                        d_595_v5_ = ((p2)[default__.safeIndex(561, (p2).length(0))]) - (d_595_v5_)
                        d_591_v2_ = (D1_DC3(d_591_v2_, True)).cf3
                    elif source13_.is_DC5:
                        d_628___mcc_h4_ = source13_.cf6
                        d_629_cf6_ = d_628___mcc_h4_
                        d_630_v24_: _dafny.Array
                        nw92_ = _dafny.Array(_dafny.Array(None, 0), 11)
                        d_630_v24_ = nw92_
                        d_631_v25_: _dafny.Array
                        nw93_ = _dafny.Array(False, 9)
                        d_631_v25_ = nw93_
                        index95_ = default__.safeIndex(964, (d_630_v24_).length(0))
                        (d_630_v24_)[index95_] = d_631_v25_
                        d_632_v26_: _dafny.Map
                        d_632_v26_ = _dafny.Map({not(p0): d_595_v5_})
                        index96_ = default__.safeIndex(964, (d_630_v24_).length(0))
                        rhs133_ = (p2)[default__.safeIndex(561, (p2).length(0))]
                        rhs134_ = (len(d_632_v26_)) + ((p2)[default__.safeIndex(561, (p2).length(0))])
                        rhs135_ = d_631_v25_
                        rhs136_ = (d_591_v2_)[default__.safeIndex(((p2)[default__.safeIndex(561, (p2).length(0))]) - (len(d_591_v2_)), len(d_591_v2_))]
                        rhs137_ = d_595_v5_
                        lhs86_ = globalState
                        lhs87_ = globalState
                        lhs88_ = d_630_v24_
                        lhs89_ = default__.safeIndex(964, (d_630_v24_).length(0))
                        lhs90_ = globalState
                        lhs86_.f4 = rhs133_
                        lhs87_.f4 = rhs134_
                        lhs88_[lhs89_] = rhs135_
                        d_590_v1_ = rhs136_
                        lhs90_.f7 = rhs137_
                        d_633_v27_: _dafny.Seq
                        d_633_v27_ = _dafny.SeqWithoutIsStrInference([(self).fm2(d_595_v5_, globalState)])
                        index97_ = default__.safeIndex(561, (p2).length(0))
                        (p2)[index97_] = (default__.safeModuloInt((self).fm2((p2)[default__.safeIndex(561, (p2).length(0))], globalState), (p2)[default__.safeIndex(561, (p2).length(0))])) + ((d_633_v27_)[default__.safeIndex(d_595_v5_, len(d_633_v27_))])
                        d_634_v28_: bool
                        d_634_v28_ = False
                        d_635_v29_: _dafny.MultiSet
                        d_635_v29_ = _dafny.MultiSet([len(d_591_v2_), (0) - (d_595_v5_), d_595_v5_])
                        rhs138_ = d_595_v5_
                        rhs139_ = (0) - (d_595_v5_)
                        rhs140_ = True
                        rhs141_ = ((624) * (d_595_v5_)) - ((d_635_v29_).cardinality)
                        rhs142_ = d_634_v28_
                        lhs91_ = globalState
                        lhs92_ = globalState
                        lhs91_.f15 = rhs138_
                        d_595_v5_ = rhs139_
                        d_634_v28_ = rhs140_
                        lhs92_.f7 = rhs141_
                        d_634_v28_ = rhs142_
                        d_636_v30_: _dafny.Seq
                        d_636_v30_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_637_v31_: _dafny.Set
                        d_637_v31_ = _dafny.Set({p1, d_634_v28_})
                        d_638_v32_: int
                        d_639_v33_: int
                        d_640_v34_: int
                        out6_: int
                        out7_: int
                        out8_: int
                        out6_, out7_, out8_ = (self).m11((d_630_v24_)[default__.safeIndex(964, (d_630_v24_).length(0))], d_636_v30_, d_637_v31_, (d_595_v5_) <= (-823), globalState)
                        d_638_v32_ = out6_
                        d_639_v33_ = out7_
                        d_640_v34_ = out8_
                    elif True:
                        d_641___mcc_h5_ = source13_.cf11
                        d_642_cf11_ = d_641___mcc_h5_
                        (globalState).f4 = d_595_v5_
                        d_643_v35_: _dafny.Set
                        d_643_v35_ = _dafny.Set({(p2)[default__.safeIndex(561, (p2).length(0))], d_595_v5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))})
                        d_644_v36_: _dafny.Seq
                        d_644_v36_ = _dafny.SeqWithoutIsStrInference([len(d_589_v0_)])
                        d_645_v38_: _dafny.Seq
                        def iife36_():
                            coll32_ = _dafny.Set()
                            compr_32_: int
                            for compr_32_ in (d_644_v36_).Elements:
                                d_646_v37_: int = compr_32_
                                if (d_646_v37_) in (d_644_v36_):
                                    coll32_ = coll32_.union(_dafny.Set([(d_646_v37_) - (len(_dafny.Map({D7_DC20(315): True})))]))
                            return _dafny.Set(coll32_)
                        d_645_v38_ = _dafny.SeqWithoutIsStrInference([d_643_v35_, _dafny.Set({(p2)[default__.safeIndex(561, (p2).length(0))]}), iife36_()
                        , _dafny.Set({(self).fm2((p2)[default__.safeIndex(561, (p2).length(0))], globalState)})])
                        d_647_v39_: _dafny.Seq
                        d_647_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(p2)[default__.safeIndex(561, (p2).length(0))], (p2)[default__.safeIndex(561, (p2).length(0))]}), d_643_v35_, (d_645_v38_)[default__.safeIndex((p2)[default__.safeIndex(561, (p2).length(0))], len(d_645_v38_))], d_643_v35_])
                        (globalState).f13 = len((d_647_v39_)[default__.safeIndex(((p2)[default__.safeIndex(561, (p2).length(0))]) + ((p2)[default__.safeIndex(561, (p2).length(0))]), len(d_647_v39_))])
                        d_648_v40_: _dafny.Map
                        d_648_v40_ = _dafny.Map({(p2)[default__.safeIndex(561, (p2).length(0))]: d_591_v2_})
                        d_649_v41_: D5
                        d_649_v41_ = D5_DC14(d_590_v1_, p3, (0) - (d_595_v5_), p0)
                        d_591_v2_ = (((d_648_v40_)[d_595_v5_] if (d_595_v5_) in (d_648_v40_) else d_591_v2_)) + ((d_591_v2_).set(default__.safeIndex(d_595_v5_, len(d_591_v2_)), (d_649_v41_).cf18))
                        index98_ = default__.safeIndex(561, (p2).length(0))
                        (p2)[index98_] = default__.safeModuloInt(d_595_v5_, d_595_v5_)
                    d_650_v42_: bool
                    d_650_v42_ = True
                    d_650_v42_ = p0
                    pass
            pass
        d_652_v43_: int
        d_652_v43_ = -298
        d_653_v44_: _dafny.Seq
        d_653_v44_ = _dafny.SeqWithoutIsStrInference([(0) - (d_652_v43_), 580, d_652_v43_])
        index99_ = default__.safeIndex(911, (p2).length(0))
        (p2)[index99_] = len(d_653_v44_)
        d_654_v45_: _dafny.Array
        nw94_ = _dafny.Array(_dafny.MultiSet({}), 4)
        d_654_v45_ = nw94_
        d_655_v46_: _dafny.Map
        d_655_v46_ = _dafny.Map({632: p1})
        d_656_v47_: _dafny.Seq
        d_656_v47_ = _dafny.SeqWithoutIsStrInference([not(((d_655_v46_)[264] if (264) in (d_655_v46_) else p1)), p3, p3])
        d_657_v48_: D7
        d_657_v48_ = D7_DC22(p3, len(d_656_v47_), p1, len(d_653_v44_))
        d_658_v49_: _dafny.Seq
        d_658_v49_ = _dafny.SeqWithoutIsStrInference([d_657_v48_, d_657_v48_])
        d_659_v50_: _dafny.MultiSet
        d_659_v50_ = _dafny.MultiSet([d_652_v43_, len(d_658_v49_)])
        index100_ = default__.safeIndex(203, (d_654_v45_).length(0))
        (d_654_v45_)[index100_] = d_659_v50_
        index101_ = default__.safeIndex(911, (p2).length(0))
        index102_ = default__.safeIndex(203, (d_654_v45_).length(0))
        rhs143_ = d_652_v43_
        rhs144_ = d_659_v50_
        lhs93_ = p2
        lhs94_ = default__.safeIndex(911, (p2).length(0))
        lhs95_ = d_654_v45_
        lhs96_ = default__.safeIndex(203, (d_654_v45_).length(0))
        lhs93_[lhs94_] = rhs143_
        lhs95_[lhs96_] = rhs144_
        d_660_v51_: _dafny.Array
        nw95_ = _dafny.Array(int(0), 28)
        d_660_v51_ = nw95_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_660_v51_).length(0)):
            d_661_i6_: int = guard_loop_0_
            if (True) and (((0) <= (d_661_i6_)) and ((d_661_i6_) < ((d_660_v51_).length(0)))):
                (d_660_v51_)[(d_661_i6_)] = (d_661_i6_) * (d_652_v43_)
        (globalState).f15 = (self).fm2(d_652_v43_, globalState)
        index103_ = default__.safeIndex(911, (p2).length(0))
        (p2)[index103_] = (d_652_v43_ if p1 else (p2)[default__.safeIndex(911, (p2).length(0))])
        d_662_v52_: bool
        d_662_v52_ = True
        d_662_v52_ = p0


class C3(T0):
    def  __init__(self):
        self.f24: D2 = D2.default()()
        self.f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f24, f25):
        (self).f24 = f24
        (self).f25 = f25

    def fm1(self, p0, p1, globalState):
        return self.f25

    def fm2(self, p0, globalState):
        return default__.safeModuloInt(((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([328]))).cardinality if not(self.f25) else (0) - (-540)), 714)

    def m1(self, p0, p1, globalState):
        d_663_v0_: int
        d_663_v0_ = -140
        d_664_v1_: _dafny.Seq
        d_664_v1_ = _dafny.SeqWithoutIsStrInference([d_663_v0_, d_663_v0_])
        d_665_v2_: _dafny.Map
        d_665_v2_ = _dafny.Map({d_663_v0_: (d_664_v1_)[default__.safeIndex(d_663_v0_, len(d_664_v1_))]})
        d_666_v3_: _dafny.Map
        d_666_v3_ = _dafny.Map({not(self.f25): d_663_v0_})
        d_667_v4_: str
        d_667_v4_ = _dafny.CodePoint('k')
        d_668_v5_: _dafny.MultiSet
        d_668_v5_ = _dafny.MultiSet([_dafny.CodePoint('y'), d_667_v4_])
        d_669_v6_: _dafny.Array
        def lambda41_(d_670_v0_):
            def lambda42_(d_671_i0_):
                return D1_DC4(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_670_v0_])))

            return lambda42_

        init21_ = lambda41_(d_663_v0_)
        nw96_ = _dafny.Array(None, 18)
        for i0_21_ in range(nw96_.length(0)):
            nw96_[i0_21_] = init21_(i0_21_)
        d_669_v6_ = nw96_
        d_672_v7_: _dafny.Map
        d_672_v7_ = _dafny.Map({d_669_v6_: self.f25})
        d_673_v8_: _dafny.Map
        d_673_v8_ = _dafny.Map({d_663_v0_: self.f25})
        d_674_v9_: _dafny.Seq
        d_674_v9_ = _dafny.SeqWithoutIsStrInference([self.f25, self.f25, ((d_673_v8_)[d_663_v0_] if (d_663_v0_) in (d_673_v8_) else self.f25)])
        d_675_v10_: _dafny.MultiSet
        d_675_v10_ = _dafny.MultiSet([True, self.f25, self.f25, self.f25, self.f25])
        d_676_v11_: _dafny.Array
        nw97_ = _dafny.Array(None, 28)
        nw97_[int(0)] = default__.safeDivisionInt(-941, d_663_v0_)
        nw97_[int(1)] = d_663_v0_
        nw97_[int(2)] = (d_663_v0_ if self.f25 else 422)
        nw97_[int(3)] = d_663_v0_
        nw97_[int(4)] = (len(d_665_v2_)) * (d_663_v0_)
        nw97_[int(5)] = d_663_v0_
        nw97_[int(6)] = d_663_v0_
        nw97_[int(7)] = ((d_666_v3_)[not(self.f25)] if (not(self.f25)) in (d_666_v3_) else d_663_v0_)
        nw97_[int(8)] = default__.safeModuloInt(len(p0), ((d_668_v5_)[d_667_v4_] if (d_667_v4_) in (d_668_v5_) else d_663_v0_))
        nw97_[int(9)] = 25
        nw97_[int(10)] = default__.safeDivisionInt(d_663_v0_, d_663_v0_)
        nw97_[int(11)] = -127
        nw97_[int(12)] = d_663_v0_
        nw97_[int(13)] = d_663_v0_
        nw97_[int(14)] = 19
        nw97_[int(15)] = len(d_672_v7_)
        nw97_[int(16)] = (0) - (d_663_v0_)
        nw97_[int(17)] = d_663_v0_
        nw97_[int(18)] = d_663_v0_
        nw97_[int(19)] = len(d_674_v9_)
        nw97_[int(20)] = d_663_v0_
        nw97_[int(21)] = (d_663_v0_) + (652)
        nw97_[int(22)] = d_663_v0_
        nw97_[int(23)] = len(d_674_v9_)
        nw97_[int(24)] = ((d_675_v10_)[self.f25] if (self.f25) in (d_675_v10_) else d_663_v0_)
        nw97_[int(25)] = default__.safeModuloInt(d_663_v0_, d_663_v0_)
        nw97_[int(26)] = d_663_v0_
        nw97_[int(27)] = default__.safeModuloInt((0) - (d_663_v0_), d_663_v0_)
        d_676_v11_ = nw97_
        index104_ = default__.safeIndex(7, (d_676_v11_).length(0))
        (d_676_v11_)[index104_] = d_663_v0_
        d_677_v12_: _dafny.Array
        nw98_ = _dafny.Array(None, 17)
        nw98_[int(0)] = _dafny.SeqWithoutIsStrInference([d_663_v0_, 572, d_663_v0_])
        nw98_[int(1)] = d_664_v1_
        nw98_[int(2)] = _dafny.SeqWithoutIsStrInference([d_663_v0_ for d_678_i1_ in range(default__.abs(440))])
        nw98_[int(3)] = _dafny.SeqWithoutIsStrInference([969])
        nw98_[int(4)] = (d_664_v1_).set(default__.safeIndex(d_663_v0_, len(d_664_v1_)), d_663_v0_)
        nw98_[int(5)] = d_664_v1_
        nw98_[int(6)] = _dafny.SeqWithoutIsStrInference([d_663_v0_, len(_dafny.Map({self.f25: (self).fm2(len(d_666_v3_), globalState)}))])
        nw98_[int(7)] = d_664_v1_
        nw98_[int(8)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nht")))])
        nw98_[int(9)] = _dafny.SeqWithoutIsStrInference([d_663_v0_, d_663_v0_])
        nw98_[int(10)] = d_664_v1_
        nw98_[int(11)] = (D6_DC16(d_664_v1_)).cf23
        nw98_[int(12)] = d_664_v1_
        nw98_[int(13)] = d_664_v1_
        nw98_[int(14)] = (d_664_v1_).set(default__.safeIndex(len(p0), len(d_664_v1_)), d_663_v0_)
        nw98_[int(15)] = d_664_v1_
        nw98_[int(16)] = d_664_v1_
        d_677_v12_ = nw98_
        d_679_v13_: _dafny.Seq
        d_679_v13_ = _dafny.SeqWithoutIsStrInference([d_677_v12_, d_677_v12_, d_677_v12_])
        d_680_v14_: _dafny.Array
        nw99_ = _dafny.Array(None, 15)
        nw99_[int(0)] = (d_677_v12_ if self.f25 else d_677_v12_)
        nw99_[int(1)] = (d_679_v13_)[default__.safeIndex(d_663_v0_, len(d_679_v13_))]
        nw99_[int(2)] = (d_677_v12_ if self.f25 else d_677_v12_)
        nw99_[int(3)] = d_677_v12_
        nw99_[int(4)] = d_677_v12_
        nw99_[int(5)] = d_677_v12_
        nw99_[int(6)] = d_677_v12_
        nw99_[int(7)] = d_677_v12_
        nw99_[int(8)] = d_677_v12_
        nw99_[int(9)] = d_677_v12_
        nw99_[int(10)] = d_677_v12_
        nw99_[int(11)] = d_677_v12_
        nw99_[int(12)] = d_677_v12_
        nw99_[int(13)] = d_677_v12_
        nw99_[int(14)] = d_677_v12_
        d_680_v14_ = nw99_
        index105_ = default__.safeIndex(229, (d_680_v14_).length(0))
        (d_680_v14_)[index105_] = d_677_v12_
        index106_ = default__.safeIndex(7, (d_676_v11_).length(0))
        index107_ = default__.safeIndex(229, (d_680_v14_).length(0))
        rhs145_ = 250
        rhs146_ = d_663_v0_
        rhs147_ = d_677_v12_
        rhs148_ = default__.safeDivisionInt((0) - ((D1_DC2(d_663_v0_)).cf2), d_663_v0_)
        lhs97_ = d_676_v11_
        lhs98_ = default__.safeIndex(7, (d_676_v11_).length(0))
        lhs99_ = globalState
        lhs100_ = d_680_v14_
        lhs101_ = default__.safeIndex(229, (d_680_v14_).length(0))
        lhs97_[lhs98_] = rhs145_
        lhs99_.f15 = rhs146_
        lhs100_[lhs101_] = rhs147_
        d_663_v0_ = rhs148_
        d_663_v0_ = d_663_v0_
        d_681_v15_: _dafny.Array
        nw100_ = _dafny.Array(False, 9)
        d_681_v15_ = nw100_
        d_682_v16_: _dafny.Seq
        d_682_v16_ = _dafny.SeqWithoutIsStrInference([d_681_v15_])
        d_683_v17_: D7
        d_683_v17_ = D7_DC19(d_682_v16_)
        rhs149_ = self.f25
        rhs150_ = d_667_v4_
        rhs151_ = (d_683_v17_).cf28
        lhs102_ = self
        lhs102_.f25 = rhs149_
        d_667_v4_ = rhs150_
        d_682_v16_ = rhs151_
        hi4_ = (d_663_v0_ if self.f25 else (d_676_v11_)[default__.safeIndex(7, (d_676_v11_).length(0))])
        for d_684_i2_ in range((len(d_665_v2_)) - ((0) - (d_663_v0_)), hi4_):
            d_685_v18_: _dafny.Array
            nw101_ = _dafny.Array(_dafny.Map({}), 24)
            d_685_v18_ = nw101_
            d_686_v19_: _dafny.Seq
            d_686_v19_ = _dafny.SeqWithoutIsStrInference([d_685_v18_, d_685_v18_, d_685_v18_])
            d_687_v20_: _dafny.Array
            nw102_ = _dafny.Array(None, 29)
            nw102_[int(0)] = (d_686_v19_)[default__.safeIndex(len(p0), len(d_686_v19_))]
            nw102_[int(1)] = d_685_v18_
            nw102_[int(2)] = d_685_v18_
            nw102_[int(3)] = d_685_v18_
            nw102_[int(4)] = d_685_v18_
            nw102_[int(5)] = d_685_v18_
            nw102_[int(6)] = d_685_v18_
            nw102_[int(7)] = d_685_v18_
            nw102_[int(8)] = d_685_v18_
            nw102_[int(9)] = (d_686_v19_)[default__.safeIndex((d_676_v11_)[default__.safeIndex(7, (d_676_v11_).length(0))], len(d_686_v19_))]
            nw102_[int(10)] = d_685_v18_
            nw102_[int(11)] = d_685_v18_
            nw102_[int(12)] = d_685_v18_
            nw102_[int(13)] = d_685_v18_
            nw102_[int(14)] = d_685_v18_
            nw102_[int(15)] = (d_685_v18_ if True else d_685_v18_)
            nw102_[int(16)] = d_685_v18_
            nw102_[int(17)] = d_685_v18_
            nw102_[int(18)] = d_685_v18_
            nw102_[int(19)] = d_685_v18_
            nw102_[int(20)] = d_685_v18_
            nw102_[int(21)] = d_685_v18_
            nw102_[int(22)] = d_685_v18_
            nw102_[int(23)] = d_685_v18_
            nw102_[int(24)] = d_685_v18_
            nw102_[int(25)] = d_685_v18_
            nw102_[int(26)] = d_685_v18_
            nw102_[int(27)] = d_685_v18_
            nw102_[int(28)] = d_685_v18_
            d_687_v20_ = nw102_
            index108_ = default__.safeIndex(447, (d_687_v20_).length(0))
            (d_687_v20_)[index108_] = d_685_v18_
            d_688_v21_: D2
            d_688_v21_ = D2_DC6((d_676_v11_)[default__.safeIndex(7, (d_676_v11_).length(0))])
            index109_ = default__.safeIndex(447, (d_687_v20_).length(0))
            rhs152_ = d_685_v18_
            rhs153_ = self.f25
            rhs154_ = True
            rhs155_ = d_688_v21_
            lhs103_ = d_687_v20_
            lhs104_ = default__.safeIndex(447, (d_687_v20_).length(0))
            lhs105_ = self
            lhs106_ = self
            lhs103_[lhs104_] = rhs152_
            lhs105_.f25 = rhs153_
            lhs106_.f25 = rhs154_
            d_688_v21_ = rhs155_
            (globalState).f4 = d_663_v0_
            (self).f25 = self.f25
            (self).f25 = not ((d_663_v0_) == (len(p0))) or (self.f25)
        d_663_v0_ = (d_676_v11_)[default__.safeIndex(7, (d_676_v11_).length(0))]
        (self).f25 = self.f25

    def m10(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Seq = _dafny.Seq({})
        d_689_v0_: _dafny.Array
        def lambda43_(d_690_i0_):
            return True

        init22_ = lambda43_
        nw103_ = _dafny.Array(None, 1)
        for i0_22_ in range(nw103_.length(0)):
            nw103_[i0_22_] = init22_(i0_22_)
        d_689_v0_ = nw103_
        index110_ = default__.safeIndex(361, (d_689_v0_).length(0))
        (d_689_v0_)[index110_] = not((self.f25 if self.f25 else self.f25))
        d_691_v1_: int
        d_691_v1_ = 199
        d_692_v2_: _dafny.Set
        d_692_v2_ = _dafny.Set({d_691_v1_})
        d_693_v3_: _dafny.Map
        d_693_v3_ = _dafny.Map({(_dafny.Set({d_691_v1_, d_691_v1_, len(d_692_v2_), d_691_v1_, d_691_v1_})) | (d_692_v2_): d_691_v1_})
        d_694_v4_: str
        d_694_v4_ = _dafny.CodePoint('v')
        d_695_v5_: _dafny.Seq
        d_695_v5_ = _dafny.SeqWithoutIsStrInference([self.f25])
        d_696_v6_: D8
        d_696_v6_ = D8_DC23(_dafny.SeqWithoutIsStrInference([self.f25, self.f25]))
        index111_ = default__.safeIndex(361, (d_689_v0_).length(0))
        rhs156_ = ((d_695_v5_) + ((d_696_v6_).cf39)) < (_dafny.SeqWithoutIsStrInference([self.f25, self.f25]))
        rhs157_ = ((0) - (d_691_v1_)) + (d_691_v1_)
        rhs158_ = (default__.safeModuloInt(d_691_v1_, d_691_v1_)) != (-740)
        rhs159_ = d_693_v3_
        rhs160_ = d_694_v4_
        lhs107_ = d_689_v0_
        lhs108_ = default__.safeIndex(361, (d_689_v0_).length(0))
        lhs109_ = globalState
        lhs110_ = self
        lhs107_[lhs108_] = rhs156_
        lhs109_.f4 = rhs157_
        lhs110_.f25 = rhs158_
        d_693_v3_ = rhs159_
        d_694_v4_ = rhs160_
        d_697_v7_: _dafny.Array
        nw104_ = _dafny.Array(int(0), 23)
        d_697_v7_ = nw104_
        d_698_v8_: _dafny.Array
        nw105_ = _dafny.Array(None, 18)
        nw105_[int(0)] = d_697_v7_
        nw105_[int(1)] = d_697_v7_
        nw105_[int(2)] = d_697_v7_
        nw105_[int(3)] = d_697_v7_
        nw105_[int(4)] = d_697_v7_
        nw105_[int(5)] = d_697_v7_
        nw105_[int(6)] = d_697_v7_
        nw105_[int(7)] = d_697_v7_
        nw105_[int(8)] = d_697_v7_
        nw105_[int(9)] = d_697_v7_
        nw105_[int(10)] = d_697_v7_
        nw105_[int(11)] = d_697_v7_
        nw105_[int(12)] = d_697_v7_
        nw105_[int(13)] = d_697_v7_
        nw105_[int(14)] = d_697_v7_
        nw105_[int(15)] = d_697_v7_
        nw105_[int(16)] = d_697_v7_
        nw105_[int(17)] = d_697_v7_
        d_698_v8_ = nw105_
        d_698_v8_ = d_698_v8_
        d_699_v9_: _dafny.Map
        d_699_v9_ = _dafny.Map({d_689_v0_: d_689_v0_})
        d_700_v10_: _dafny.Seq
        d_700_v10_ = _dafny.SeqWithoutIsStrInference([d_691_v1_, len(d_699_v9_)])
        (self).f25 = (len((d_700_v10_ if self.f25 else _dafny.SeqWithoutIsStrInference([d_691_v1_ for d_701_i1_ in range(default__.abs(159))])))) not in (d_700_v10_)
        d_702_v11_: _dafny.MultiSet
        d_702_v11_ = _dafny.MultiSet([self.f25])
        d_703_v12_: _dafny.Seq
        d_703_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpoxflb"))
        d_704_v13_: D7
        d_704_v13_ = D7_DC21(((d_702_v11_)[self.f25] if (self.f25) in (d_702_v11_) else len(d_703_v12_)), d_691_v1_, False, not ((d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]) or (self.f25), False)
        source14_ = d_704_v13_
        if source14_.is_DC20:
            d_705___mcc_h0_ = source14_.cf29
            d_706_cf29_ = d_705___mcc_h0_
            if (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]:
                d_707_v14_: C1
                nw106_ = C1()
                nw106_.ctor__()
                d_707_v14_ = nw106_
                index112_ = default__.safeIndex(857, (d_697_v7_).length(0))
                (d_697_v7_)[index112_] = default__.safeModuloInt(d_691_v1_, len(d_695_v5_))
                index113_ = default__.safeIndex(361, (d_689_v0_).length(0))
                index114_ = default__.safeIndex(857, (d_697_v7_).length(0))
                index115_ = default__.safeIndex(361, (d_689_v0_).length(0))
                rhs161_ = self.f25
                rhs162_ = (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]
                rhs163_ = len((_dafny.SeqWithoutIsStrInference([d_694_v4_ for d_708_i2_ in range(default__.abs(788))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tswns"))))
                rhs164_ = d_697_v7_
                rhs165_ = (default__.fm39(d_691_v1_, d_694_v4_, d_700_v10_, not(self.f25), globalState)) < ((_dafny.SeqWithoutIsStrInference([d_706_cf29_, (0) - (len(_dafny.SeqWithoutIsStrInference([D11_DC28(d_692_v2_) for d_709_i3_ in range(default__.abs(318))]))), len(d_703_v12_)])) + (d_700_v10_))
                lhs111_ = d_689_v0_
                lhs112_ = default__.safeIndex(361, (d_689_v0_).length(0))
                lhs113_ = self
                lhs114_ = d_697_v7_
                lhs115_ = default__.safeIndex(857, (d_697_v7_).length(0))
                lhs116_ = d_689_v0_
                lhs117_ = default__.safeIndex(361, (d_689_v0_).length(0))
                lhs111_[lhs112_] = rhs161_
                lhs113_.f25 = rhs162_
                lhs114_[lhs115_] = rhs163_
                r1 = rhs164_
                lhs116_[lhs117_] = rhs165_
                (self).f25 = False
                d_710_v15_: D1
                d_710_v15_ = D1_DC2(len(d_700_v10_))
                pat_let_tv8_ = d_691_v1_
                d_711_v16_: _dafny.Array
                nw107_ = _dafny.Array(None, 20)
                nw107_[int(0)] = d_710_v15_
                nw107_[int(1)] = d_710_v15_
                nw107_[int(2)] = d_710_v15_
                nw107_[int(3)] = d_710_v15_
                nw107_[int(4)] = D1_DC2(d_691_v1_)
                nw107_[int(5)] = d_710_v15_
                nw107_[int(6)] = D1_DC2(d_706_cf29_)
                nw107_[int(7)] = d_710_v15_
                nw107_[int(8)] = D1_DC2((d_700_v10_)[default__.safeIndex(len((d_703_v12_).set(default__.safeIndex(len(d_692_v2_), len(d_703_v12_)), d_694_v4_)), len(d_700_v10_))])
                nw107_[int(9)] = D1_DC2(d_706_cf29_)
                nw107_[int(10)] = d_710_v15_
                nw107_[int(11)] = d_710_v15_
                nw107_[int(12)] = d_710_v15_
                nw107_[int(13)] = d_710_v15_
                nw107_[int(14)] = d_710_v15_
                def iife37_(_pat_let2_0):
                    def iife38_(d_712_dt__update__tmp_h0_):
                        def iife39_(_pat_let3_0):
                            def iife40_(d_713_dt__update_hcf2_h0_):
                                return D1_DC2(d_713_dt__update_hcf2_h0_)
                            return iife40_(_pat_let3_0)
                        return iife39_(pat_let_tv8_)
                    return iife38_(_pat_let2_0)
                nw107_[int(15)] = iife37_(D1_DC2((d_697_v7_)[default__.safeIndex(857, (d_697_v7_).length(0))]))
                nw107_[int(16)] = d_710_v15_
                nw107_[int(17)] = d_710_v15_
                nw107_[int(18)] = d_710_v15_
                nw107_[int(19)] = default__.fm40((d_697_v7_)[default__.safeIndex(857, (d_697_v7_).length(0))], (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))], self.f25, d_696_v6_, globalState)
                d_711_v16_ = nw107_
                d_711_v16_ = d_711_v16_
                index116_ = default__.safeIndex(857, (d_697_v7_).length(0))
                (d_697_v7_)[index116_] = 603
            elif True:
                d_714_v17_: _dafny.Map
                d_714_v17_ = _dafny.Map({d_692_v2_: self.f25})
                d_700_v10_ = ((d_700_v10_ if ((d_714_v17_)[d_692_v2_] if (d_692_v2_) in (d_714_v17_) else default__.fm26(globalState)) else d_700_v10_)) + (_dafny.SeqWithoutIsStrInference([d_691_v1_]))
                d_715_v18_: _dafny.Map
                d_715_v18_ = _dafny.Map({d_691_v1_: (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]})
                d_716_v19_: _dafny.Map
                d_716_v19_ = _dafny.Map({d_703_v12_: d_706_cf29_})
                d_715_v18_ = (d_715_v18_).set(default__.safeModuloInt(len(d_716_v19_), d_691_v1_), True)
                d_695_v5_ = default__.fm27(globalState)
                index117_ = default__.safeIndex(927, (d_697_v7_).length(0))
                (d_697_v7_)[index117_] = d_691_v1_
                index118_ = default__.safeIndex(927, (d_697_v7_).length(0))
                (d_697_v7_)[index118_] = d_706_cf29_
                index119_ = default__.safeIndex(927, (d_697_v7_).length(0))
                (d_697_v7_)[index119_] = default__.safeModuloInt(d_706_cf29_, d_691_v1_)
            d_717_v20_: _dafny.Set
            d_717_v20_ = _dafny.Set({default__.fm41(d_694_v4_, globalState)})
            d_718_v21_: _dafny.Map
            d_718_v21_ = _dafny.Map({d_717_v20_: d_694_v4_})
            d_718_v21_ = (d_718_v21_).set(d_717_v20_, d_694_v4_)
            d_719_v22_: _dafny.Array
            nw108_ = _dafny.Array(D7.default()(), 12)
            d_719_v22_ = nw108_
            d_720_v23_: D3
            d_720_v23_ = D3_DC10(808, d_706_cf29_)
            d_721_v24_: _dafny.Map
            d_721_v24_ = _dafny.Map({d_719_v22_: d_720_v23_})
            (self).f25 = (len(d_721_v24_)) == (d_691_v1_)
            (globalState).f4 = ((d_702_v11_).cardinality) - (d_691_v1_)
        elif source14_.is_DC21:
            d_722___mcc_h1_ = source14_.cf30
            d_723___mcc_h2_ = source14_.cf31
            d_724___mcc_h3_ = source14_.cf32
            d_725___mcc_h4_ = source14_.cf33
            d_726___mcc_h5_ = source14_.cf34
            d_727_cf34_ = d_726___mcc_h5_
            d_728_cf33_ = d_725___mcc_h4_
            d_729_cf32_ = d_724___mcc_h3_
            d_730_cf31_ = d_723___mcc_h2_
            d_731_cf30_ = d_722___mcc_h1_
            pat_let_tv9_ = d_692_v2_
            def iife41_(_pat_let4_0):
                def iife42_(d_732_dt__update__tmp_h1_):
                    def iife43_(_pat_let5_0):
                        def iife44_(d_733_dt__update_hcf45_h0_):
                            return D11_DC28(d_733_dt__update_hcf45_h0_)
                        return iife44_(_pat_let5_0)
                    return iife43_(pat_let_tv9_)
                return iife42_(_pat_let4_0)
            source15_ = iife41_(D11_DC28(d_692_v2_))
            if source15_.is_DC29:
                (globalState).f4 = (d_731_cf30_) * (d_731_cf30_)
                index120_ = default__.safeIndex(361, (d_689_v0_).length(0))
                (d_689_v0_)[index120_] = (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]
                d_734_v25_: C2
                nw109_ = C2()
                nw109_.ctor__()
                d_734_v25_ = nw109_
                d_735_v26_: _dafny.Map
                d_735_v26_ = _dafny.Map({d_691_v1_: d_729_cf32_})
                d_736_v27_: D7
                d_736_v27_ = D7_DC20(316)
                rhs166_ = d_730_cf31_
                rhs167_ = (_dafny.SeqWithoutIsStrInference([True, not(self.f25)])) != ((d_695_v5_).set(default__.safeIndex((d_736_v27_).cf29, len(d_695_v5_)), (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]))
                rhs168_ = d_735_v26_
                rhs169_ = not (self.f25) or (d_727_cf34_)
                rhs170_ = d_729_cf32_
                lhs118_ = globalState
                lhs118_.f4 = rhs166_
                d_729_cf32_ = rhs167_
                d_735_v26_ = rhs168_
                d_729_cf32_ = rhs169_
                d_729_cf32_ = rhs170_
            elif source15_.is_DC30:
                d_737___mcc_h11_ = source15_.cf46
                d_738___mcc_h12_ = source15_.cf47
                d_739_cf47_ = d_738___mcc_h12_
                d_740_cf46_ = d_737___mcc_h11_
                d_741_v28_: _dafny.Map
                d_741_v28_ = _dafny.Map({-980: len(_dafny.Map({d_727_cf34_: d_691_v1_}))})
                d_742_v29_: _dafny.Map
                d_742_v29_ = _dafny.Map({(0) - (len(_dafny.Map({d_729_cf32_: ((d_741_v28_)[d_730_cf31_] if (d_730_cf31_) in (d_741_v28_) else d_691_v1_)}))): _dafny.SeqWithoutIsStrInference([d_694_v4_ for d_743_i4_ in range(default__.abs(59))])})
                d_744_v30_: D1
                d_744_v30_ = D1_DC2(d_731_cf30_)
                d_745_v31_: _dafny.Set
                d_745_v31_ = _dafny.Set({d_744_v30_})
                d_746_v32_: _dafny.Map
                d_746_v32_ = _dafny.Map({(d_742_v29_) | (d_742_v29_): d_745_v31_})
                d_747_v33_: _dafny.Map
                d_747_v33_ = _dafny.Map({d_739_cf47_: d_745_v31_})
                d_748_v34_: _dafny.MultiSet
                d_748_v34_ = _dafny.MultiSet([d_744_v30_])
                def iife45_():
                    coll33_ = _dafny.Set()
                    compr_33_: D1
                    for compr_33_ in (d_748_v34_).Elements:
                        d_749_v35_: D1 = compr_33_
                        if (d_749_v35_) in (d_748_v34_):
                            coll33_ = coll33_.union(_dafny.Set([d_749_v35_]))
                    return _dafny.Set(coll33_)
                d_746_v32_ = (d_746_v32_).set(d_742_v29_, ((d_747_v33_)[d_731_cf30_] if (d_731_cf30_) in (d_747_v33_) else iife45_()
                ))
                (globalState).f4 = 897
                d_750_v36_: C2
                nw110_ = C2()
                nw110_.ctor__()
                d_750_v36_ = nw110_
                d_751_v37_: _dafny.Map
                d_751_v37_ = _dafny.Map({(d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]: True})
                d_752_v38_: _dafny.Map
                d_752_v38_ = _dafny.Map({d_751_v37_: d_689_v0_})
                d_753_v39_: _dafny.Array
                def lambda44_(d_754_v5_):
                    def lambda45_(d_755_i5_):
                        return d_754_v5_

                    return lambda45_

                init23_ = lambda44_(d_695_v5_)
                nw111_ = _dafny.Array(None, 7)
                for i0_23_ in range(nw111_.length(0)):
                    nw111_[i0_23_] = init23_(i0_23_)
                d_753_v39_ = nw111_
                index121_ = default__.safeIndex(76, (d_753_v39_).length(0))
                (d_753_v39_)[index121_] = (_dafny.SeqWithoutIsStrInference([self.f25, True])) + (d_695_v5_)
                index122_ = default__.safeIndex(76, (d_753_v39_).length(0))
                rhs171_ = d_752_v38_
                rhs172_ = default__.fm27(globalState)
                lhs119_ = d_753_v39_
                lhs120_ = default__.safeIndex(76, (d_753_v39_).length(0))
                d_752_v38_ = rhs171_
                lhs119_[lhs120_] = rhs172_
            elif source15_.is_DC28:
                d_756___mcc_h13_ = source15_.cf45
                d_757_cf45_ = d_756___mcc_h13_
                (globalState).f15 = (0) - (d_691_v1_)
                (globalState).f13 = d_691_v1_
                d_729_cf32_ = self.f25
                (globalState).f7 = d_730_cf31_
            elif True:
                d_758___mcc_h14_ = source15_.cf48
                d_759_cf48_ = d_758___mcc_h14_
                d_760_v40_: _dafny.Map
                d_760_v40_ = _dafny.Map({self.f25: _dafny.Map({d_728_cf33_: d_728_cf33_})})
                d_761_v41_: _dafny.Map
                d_761_v41_ = _dafny.Map({(self).fm1(d_760_v40_, len(d_703_v12_), globalState): d_727_cf34_})
                d_762_v42_: _dafny.Map
                d_762_v42_ = _dafny.Map({d_728_cf33_: d_761_v41_})
                d_762_v42_ = (d_762_v42_).set(not((d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]), d_761_v41_)
                (self).f25 = (d_695_v5_)[default__.safeIndex(d_730_cf31_, len(d_695_v5_))]
                d_763_v43_: _dafny.MultiSet
                d_763_v43_ = _dafny.MultiSet([len(d_695_v5_)])
                d_764_v44_: D1
                d_764_v44_ = D1_DC4(d_763_v43_)
                def iife46_():
                    coll34_ = _dafny.Set()
                    compr_34_: int
                    for compr_34_ in _dafny.IntegerRange(376, 392):
                        d_765_v45_: int = compr_34_
                        if ((376) <= (d_765_v45_)) and ((d_765_v45_) < (392)):
                            coll34_ = coll34_.union(_dafny.Set([(d_765_v45_) + (len(_dafny.Map({d_731_cf30_: d_731_cf30_})))]))
                    return _dafny.Set(coll34_)
                d_764_v44_ = default__.fm42(d_731_cf30_, len(iife46_()
                ), globalState)
                d_766_v46_: D0
                d_766_v46_ = D0_DC0(d_689_v0_)
                d_767_v47_: _dafny.Seq
                d_767_v47_ = _dafny.SeqWithoutIsStrInference([d_766_v46_])
                (self).f25 = ((d_767_v47_) + (d_767_v47_)) == (d_767_v47_)
            d_691_v1_ = 698
            d_768_v48_: _dafny.MultiSet
            d_768_v48_ = _dafny.MultiSet([d_692_v2_, d_692_v2_, d_692_v2_])
            source16_ = D5_DC13((d_768_v48_) | (_dafny.MultiSet([d_692_v2_, d_692_v2_, d_692_v2_])))
            if source16_.is_DC14:
                d_769___mcc_h15_ = source16_.cf18
                d_770___mcc_h16_ = source16_.cf19
                d_771___mcc_h17_ = source16_.cf20
                d_772___mcc_h18_ = source16_.cf21
                d_773_cf21_ = d_772___mcc_h18_
                d_774_cf20_ = d_771___mcc_h17_
                d_775_cf19_ = d_770___mcc_h16_
                d_776_cf18_ = d_769___mcc_h15_
                d_773_cf21_ = d_729_cf32_
                d_777_v49_: C0
                nw112_ = C0()
                nw112_.ctor__()
                d_777_v49_ = nw112_
                d_778_v50_: C2
                nw113_ = C2()
                nw113_.ctor__()
                d_778_v50_ = nw113_
                (globalState).f4 = ((d_702_v11_)[d_775_cf19_] if (d_775_cf19_) in (d_702_v11_) else (len(d_692_v2_) if d_775_cf19_ else d_730_cf31_))
            elif source16_.is_DC13:
                d_779___mcc_h19_ = source16_.cf17
                d_780_cf17_ = d_779___mcc_h19_
                d_781_v51_: _dafny.MultiSet
                d_781_v51_ = _dafny.MultiSet([d_730_cf31_])
                index123_ = default__.safeIndex(563, (d_697_v7_).length(0))
                (d_697_v7_)[index123_] = ((d_781_v51_).cardinality) + (len(d_703_v12_))
                d_782_v52_: D5
                d_782_v52_ = D5_DC13(default__.fm43((0) - (d_730_cf31_), globalState))
                d_783_v53_: _dafny.Map
                d_783_v53_ = _dafny.Map({(d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]: d_703_v12_})
                d_784_v54_: _dafny.Seq
                d_784_v54_ = _dafny.SeqWithoutIsStrInference([d_703_v12_, d_703_v12_, d_703_v12_, d_703_v12_, d_703_v12_])
                d_785_v55_: _dafny.Seq
                d_785_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_703_v12_, d_703_v12_, ((d_783_v53_)[not((d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))])] if (not((d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))])) in (d_783_v53_) else d_703_v12_), d_703_v12_]), (d_784_v54_).set(default__.safeIndex(d_691_v1_, len(d_784_v54_)), d_703_v12_), d_784_v54_])
                index124_ = default__.safeIndex(563, (d_697_v7_).length(0))
                rhs173_ = (d_731_cf30_) * (d_691_v1_)
                rhs174_ = len((d_785_v55_)[default__.safeIndex((0) - ((self).fm2(d_730_cf31_, globalState)), len(d_785_v55_))])
                rhs175_ = d_782_v52_
                rhs176_ = True
                rhs177_ = d_702_v11_
                lhs121_ = d_697_v7_
                lhs122_ = default__.safeIndex(563, (d_697_v7_).length(0))
                lhs121_[lhs122_] = rhs173_
                d_691_v1_ = rhs174_
                d_782_v52_ = rhs175_
                d_728_cf33_ = rhs176_
                d_702_v11_ = rhs177_
                d_786_v56_: _dafny.Array
                nw114_ = _dafny.Array(_dafny.MultiSet({}), 6)
                d_786_v56_ = nw114_
                d_787_v57_: _dafny.Set
                d_787_v57_ = _dafny.Set({self.f25, d_729_cf32_, (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]})
                d_788_v58_: D12
                d_788_v58_ = D12_DC32(d_787_v57_)
                index125_ = default__.safeIndex(650, (d_786_v56_).length(0))
                (d_786_v56_)[index125_] = (_dafny.MultiSet([d_787_v57_, (d_788_v58_).cf49, d_787_v57_, d_787_v57_, d_787_v57_])).set(d_787_v57_, default__.abs(d_730_cf31_))
                d_789_v59_: _dafny.MultiSet
                d_789_v59_ = _dafny.MultiSet([d_787_v57_])
                d_790_v60_: _dafny.Map
                d_790_v60_ = _dafny.Map({d_729_cf32_: d_731_cf30_})
                index126_ = default__.safeIndex(650, (d_786_v56_).length(0))
                index127_ = default__.safeIndex(361, (d_689_v0_).length(0))
                rhs178_ = (d_789_v59_) | ((d_789_v59_).intersection(d_789_v59_))
                rhs179_ = default__.fm26(globalState)
                rhs180_ = (d_700_v10_)[default__.safeIndex(((d_790_v60_)[not(not(d_727_cf34_))] if (not(not(d_727_cf34_))) in (d_790_v60_) else d_731_cf30_), len(d_700_v10_))]
                lhs123_ = d_786_v56_
                lhs124_ = default__.safeIndex(650, (d_786_v56_).length(0))
                lhs125_ = d_689_v0_
                lhs126_ = default__.safeIndex(361, (d_689_v0_).length(0))
                lhs127_ = globalState
                lhs123_[lhs124_] = rhs178_
                lhs125_[lhs126_] = rhs179_
                lhs127_.f7 = rhs180_
                (globalState).f15 = d_730_cf31_
                d_791_v61_: _dafny.Map
                d_791_v61_ = _dafny.Map({d_691_v1_: d_728_cf33_})
                d_792_v62_: _dafny.Map
                d_792_v62_ = _dafny.Map({default__.fm35(globalState): d_730_cf31_})
                d_793_v63_: _dafny.Array
                nw115_ = _dafny.Array(None, 10)
                nw115_[int(0)] = default__.safeDivisionInt(len(d_700_v10_), d_691_v1_)
                nw115_[int(1)] = default__.safeModuloInt(len(d_791_v61_), len(d_792_v62_))
                nw115_[int(2)] = d_731_cf30_
                nw115_[int(3)] = (0) - ((d_697_v7_)[default__.safeIndex(563, (d_697_v7_).length(0))])
                nw115_[int(4)] = default__.safeDivisionInt(len(d_791_v61_), d_731_cf30_)
                nw115_[int(5)] = d_731_cf30_
                nw115_[int(6)] = (d_697_v7_)[default__.safeIndex(563, (d_697_v7_).length(0))]
                nw115_[int(7)] = (d_697_v7_)[default__.safeIndex(563, (d_697_v7_).length(0))]
                nw115_[int(8)] = d_730_cf31_
                nw115_[int(9)] = (self).fm2((d_697_v7_)[default__.safeIndex(563, (d_697_v7_).length(0))], globalState)
                d_793_v63_ = nw115_
                r1 = d_793_v63_
            elif True:
                d_794___mcc_h20_ = source16_.cf22
                d_795_cf22_ = d_794___mcc_h20_
                d_796_v64_: _dafny.Map
                d_796_v64_ = _dafny.Map({False: len(d_703_v12_)})
                d_797_v65_: _dafny.Map
                d_797_v65_ = d_796_v64_
                d_798_v66_: _dafny.MultiSet
                d_798_v66_ = _dafny.MultiSet([d_796_v64_, (d_797_v65_), (_dafny.Map({(d_695_v5_)[default__.safeIndex(d_731_cf30_, len(d_695_v5_))]: d_691_v1_})) | (default__.fm44(d_731_cf30_, True, globalState)), (d_796_v64_) | (_dafny.Map({default__.fm26(globalState): d_731_cf30_}))])
                (globalState).f4 = ((d_798_v66_)[d_796_v64_] if (d_796_v64_) in (d_798_v66_) else default__.safeModuloInt(d_691_v1_, d_691_v1_))
                (globalState).f7 = d_691_v1_
                d_697_v7_ = (D2_DC7((d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))], d_702_v11_, d_697_v7_)).cf10
                d_727_cf34_ = (d_730_cf31_) in (d_700_v10_)
            d_702_v11_ = (d_702_v11_).set(True, default__.abs(643))
        elif source14_.is_DC22:
            d_799___mcc_h6_ = source14_.cf35
            d_800___mcc_h7_ = source14_.cf36
            d_801___mcc_h8_ = source14_.cf37
            d_802___mcc_h9_ = source14_.cf38
            d_803_cf38_ = d_802___mcc_h9_
            d_804_cf37_ = d_801___mcc_h8_
            d_805_cf36_ = d_800___mcc_h7_
            d_806_cf35_ = d_799___mcc_h6_
            (self).f25 = (d_702_v11_).issubset(d_702_v11_)
            d_807_v67_: _dafny.Set
            d_807_v67_ = _dafny.Set({(d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]})
            d_808_v68_: _dafny.Map
            d_808_v68_ = _dafny.Map({(0) - (len(d_807_v67_)): d_805_cf36_})
            d_809_v69_: D3
            d_809_v69_ = D3_DC10(d_691_v1_, len(d_808_v68_))
            d_810_v70_: _dafny.Map
            d_810_v70_ = _dafny.Map({d_703_v12_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "de"))})
            pat_let_tv10_ = d_691_v1_
            d_811_v71_: _dafny.Seq
            def iife47_(_pat_let6_0):
                def iife48_(d_812_dt__update__tmp_h2_):
                    def iife49_(_pat_let7_0):
                        def iife50_(d_813_dt__update_hcf13_h0_):
                            return D3_DC10(d_813_dt__update_hcf13_h0_, (d_812_dt__update__tmp_h2_).cf14)
                        return iife50_(_pat_let7_0)
                    return iife49_(pat_let_tv10_)
                return iife48_(_pat_let6_0)
            d_811_v71_ = _dafny.SeqWithoutIsStrInference([d_809_v69_, iife47_(d_809_v69_), d_809_v69_, d_809_v69_, D3_DC10(len(d_810_v70_), -373)])
            d_814_v72_: _dafny.Map
            d_814_v72_ = _dafny.Map({(d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]: d_811_v71_})
            d_815_v73_: _dafny.Map
            d_815_v73_ = _dafny.Map({d_806_cf35_: self.f25})
            d_814_v72_ = _dafny.Map({((d_815_v73_)[(d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]] if ((d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]) in (d_815_v73_) else d_806_cf35_): (_dafny.SeqWithoutIsStrInference([d_809_v69_, d_809_v69_])).set(default__.safeIndex(d_805_cf36_, len(_dafny.SeqWithoutIsStrInference([d_809_v69_, d_809_v69_]))), d_809_v69_)})
            d_804_cf37_ = (d_692_v2_).issubset(d_692_v2_)
            index128_ = default__.safeIndex(361, (d_689_v0_).length(0))
            (d_689_v0_)[index128_] = default__.fm26(globalState)
        elif True:
            d_816___mcc_h10_ = source14_.cf28
            d_817_cf28_ = d_816___mcc_h10_
            d_818_v74_: _dafny.Map
            d_818_v74_ = _dafny.Map({655: _dafny.CodePoint('g')})
            index129_ = default__.safeIndex(361, (d_689_v0_).length(0))
            index130_ = default__.safeIndex(361, (d_689_v0_).length(0))
            rhs181_ = self.f25
            rhs182_ = d_818_v74_
            rhs183_ = (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]
            rhs184_ = (d_691_v1_) - ((0) - (d_691_v1_))
            rhs185_ = d_697_v7_
            lhs128_ = d_689_v0_
            lhs129_ = default__.safeIndex(361, (d_689_v0_).length(0))
            lhs130_ = d_689_v0_
            lhs131_ = default__.safeIndex(361, (d_689_v0_).length(0))
            lhs128_[lhs129_] = rhs181_
            d_818_v74_ = rhs182_
            lhs130_[lhs131_] = rhs183_
            d_691_v1_ = rhs184_
            d_697_v7_ = rhs185_
            d_695_v5_ = d_695_v5_
            (self).f25 = (self.f25 if True else (self.f25) == (self.f25))
            d_819_v75_: _dafny.Array
            nw116_ = _dafny.Array(_dafny.Map({}), 11)
            d_819_v75_ = nw116_
            index131_ = default__.safeIndex(530, (d_819_v75_).length(0))
            (d_819_v75_)[index131_] = (d_818_v74_) | (d_818_v74_)
            index132_ = default__.safeIndex(530, (d_819_v75_).length(0))
            (d_819_v75_)[index132_] = d_818_v74_
        d_820_v76_: _dafny.Set
        d_820_v76_ = _dafny.Set({d_697_v7_})
        d_821_v77_: _dafny.Map
        d_821_v77_ = _dafny.Map({d_703_v12_: d_691_v1_})
        hi5_ = ((d_821_v77_)[d_703_v12_] if (d_703_v12_) in (d_821_v77_) else d_691_v1_)
        for d_822_i6_ in range(len(d_820_v76_), hi5_):
            index133_ = default__.safeIndex(361, (d_689_v0_).length(0))
            (d_689_v0_)[index133_] = (d_704_v13_).cf34
            (self).f25 = ((d_822_i6_) + (211)) != (d_822_i6_)
            d_823_v78_: C2
            nw117_ = C2()
            nw117_.ctor__()
            d_823_v78_ = nw117_
            d_824_v79_: _dafny.Map
            d_824_v79_ = _dafny.Map({self.f25: d_823_v78_})
            d_825_v80_: _dafny.Map
            d_825_v80_ = _dafny.Map({self.f25: self.f25})
            d_826_v81_: _dafny.Map
            d_826_v81_ = _dafny.Map({self.f25: d_825_v80_})
            index134_ = default__.safeIndex(361, (d_689_v0_).length(0))
            rhs186_ = (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]
            rhs187_ = True
            rhs188_ = ((d_824_v79_)[(d_823_v78_).fm1(d_826_v81_, d_822_i6_, globalState)] if ((d_823_v78_).fm1(d_826_v81_, d_822_i6_, globalState)) in (d_824_v79_) else d_823_v78_)
            lhs132_ = d_689_v0_
            lhs133_ = default__.safeIndex(361, (d_689_v0_).length(0))
            lhs134_ = self
            lhs132_[lhs133_] = rhs186_
            lhs134_.f25 = rhs187_
            d_823_v78_ = rhs188_
            index135_ = default__.safeIndex(361, (d_689_v0_).length(0))
            (d_689_v0_)[index135_] = (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]
        d_827_v82_: C1
        nw118_ = C1()
        nw118_.ctor__()
        d_827_v82_ = nw118_
        r0 = _dafny.Set({d_691_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvydf")))})
        d_828_v83_: _dafny.MultiSet
        d_828_v83_ = _dafny.MultiSet([d_691_v1_])
        r1 = (d_697_v7_ if ((default__.fm42(len(_dafny.SeqWithoutIsStrInference([23 for d_829_i7_ in range(default__.abs(-269))])), d_691_v1_, globalState)).cf5).isdisjoint(d_828_v83_) else d_697_v7_)
        d_830_v84_: _dafny.Map
        d_830_v84_ = _dafny.Map({d_691_v1_: (d_689_v0_)[default__.safeIndex(361, (d_689_v0_).length(0))]})
        d_831_v86_: _dafny.Seq
        def iife51_():
            coll35_ = _dafny.Map()
            compr_35_: int
            for compr_35_ in _dafny.IntegerRange(11, 121):
                d_832_v85_: int = compr_35_
                if ((11) <= (d_832_v85_)) and ((d_832_v85_) < (121)):
                    coll35_[(d_832_v85_) - (d_691_v1_)] = not(True)
            return _dafny.Map(coll35_)
        d_831_v86_ = _dafny.SeqWithoutIsStrInference([d_830_v84_, iife51_()
        , _dafny.Map({d_691_v1_: self.f25}), d_830_v84_, d_830_v84_])
        d_833_v87_: D14
        d_833_v87_ = D14_DC35(d_831_v86_)
        r2 = (d_833_v87_).cf52
        return r0, r1, r2


class C4(T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm7(self, p0, globalState):
        def iife52_():
            coll36_ = _dafny.Set()
            compr_36_: int
            for compr_36_ in _dafny.IntegerRange(470, 446):
                d_834_v0_: int = compr_36_
                if ((470) <= (d_834_v0_)) and ((d_834_v0_) < (446)):
                    coll36_ = coll36_.union(_dafny.Set([(d_834_v0_) * (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True, False])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_835_i0_ in range(default__.abs(412))]))])))]))
            return _dafny.Set(coll36_)
        return ((D5_DC13(_dafny.MultiSet([_dafny.Set({-472, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])), len(_dafny.SeqWithoutIsStrInference([False, True]))})]))).cf17).intersection((_dafny.MultiSet([_dafny.Set({-140}), iife52_()
        ])) | (_dafny.MultiSet([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, False])), 710})])))

    def fm8(self, p0, p1, p2, globalState):
        return (_dafny.Set({604, (0) - (-912), 783})).ispropersubset(_dafny.Set({(_dafny.MultiSet([241, len(_dafny.Map({not(not(False)): False})), -235])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecngi")))}))

    def fm23(self, p0, p1, p2, globalState):
        return default__.safeModuloInt((909 if not(True) else 731), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xueolns"))) if True else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apewrwfc")))))

    def m6(self, p0, p1, p2, p3, globalState):
        d_836_v0_: _dafny.Array
        nw119_ = _dafny.Array(None, 5)
        nw119_[int(0)] = p2
        nw119_[int(1)] = p2
        nw119_[int(2)] = p2
        nw119_[int(3)] = p2
        nw119_[int(4)] = p2
        d_836_v0_ = nw119_
        d_837_v1_: D0
        d_837_v1_ = D0_DC1(d_836_v0_)
        source17_ = d_837_v1_
        if source17_.is_DC1:
            d_838___mcc_h0_ = source17_.cf1
            d_839_cf1_ = d_838___mcc_h0_
            d_839_cf1_ = d_839_cf1_
            d_840_v2_: _dafny.Set
            d_840_v2_ = _dafny.Set({p0})
            d_841_v3_: _dafny.Seq
            d_841_v3_ = _dafny.SeqWithoutIsStrInference([d_840_v2_])
            d_842_v4_: _dafny.Map
            d_842_v4_ = _dafny.Map({p2: d_841_v3_})
            d_843_v5_: _dafny.Seq
            d_843_v5_ = _dafny.SeqWithoutIsStrInference([p2])
            d_844_v6_: _dafny.Seq
            d_844_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnhu"))
            d_842_v4_ = (d_842_v4_).set((d_843_v5_)[default__.safeIndex(len(d_844_v6_), len(d_843_v5_))], (d_841_v3_ if p2 else d_841_v3_))
            (globalState).f15 = ((p0) + (len(d_844_v6_))) - (len(d_844_v6_))
            d_845_v7_: bool
            d_845_v7_ = True
            d_845_v7_ = p2
        elif True:
            d_846___mcc_h1_ = source17_.cf0
            d_847_cf0_ = d_846___mcc_h1_
            d_848_v8_: _dafny.Array
            def lambda46_(d_849_p2_):
                def lambda47_(d_850_i0_):
                    return (_dafny.Set({d_849_p2_})) | (_dafny.Set({d_849_p2_}))

                return lambda47_

            init24_ = lambda46_(p2)
            nw120_ = _dafny.Array(None, 15)
            for i0_24_ in range(nw120_.length(0)):
                nw120_[i0_24_] = init24_(i0_24_)
            d_848_v8_ = nw120_
            d_851_v9_: _dafny.Set
            d_851_v9_ = _dafny.Set({p2, p2})
            index136_ = default__.safeIndex(914, (d_848_v8_).length(0))
            (d_848_v8_)[index136_] = d_851_v9_
            d_852_v10_: _dafny.Map
            d_852_v10_ = _dafny.Map({True: p2})
            index137_ = default__.safeIndex(914, (d_848_v8_).length(0))
            (d_848_v8_)[index137_] = _dafny.Set({((d_852_v10_)[p2] if (p2) in (d_852_v10_) else p2)})
            d_853_v11_: _dafny.Map
            d_853_v11_ = _dafny.Map({p0: p2})
            d_854_v12_: bool
            d_854_v12_ = True
            d_855_v13_: _dafny.Seq
            d_855_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wihxevenh"))
            d_856_v14_: D1
            d_856_v14_ = D1_DC3(d_855_v13_, p2)
            d_857_v15_: T0
            nw121_ = C2()
            nw121_.ctor__()
            d_857_v15_ = nw121_
            d_858_v16_: _dafny.Set
            d_858_v16_ = _dafny.Set({d_857_v15_})
            rhs189_ = _dafny.Map({p0: (d_856_v14_).cf4})
            rhs190_ = False
            rhs191_ = (len(d_852_v10_) if not ((self).fm8(len(d_858_v16_), False, p0, globalState)) or (d_854_v12_) else p0)
            lhs135_ = globalState
            d_853_v11_ = rhs189_
            d_854_v12_ = rhs190_
            lhs135_.f4 = rhs191_
            d_859_v17_: str
            d_859_v17_ = _dafny.CodePoint('g')
            d_860_v18_: _dafny.Set
            d_860_v18_ = _dafny.Set({_dafny.CodePoint('s'), d_859_v17_})
            d_861_v19_: _dafny.Map
            d_861_v19_ = _dafny.Map({p0: len(d_860_v18_)})
            d_861_v19_ = d_861_v19_
            d_852_v10_ = ((d_852_v10_) | (d_852_v10_)) | (d_852_v10_)
        d_862_v20_: _dafny.Map
        d_862_v20_ = _dafny.Map({p0: p2})
        d_862_v20_ = d_862_v20_
        d_863_v21_: _dafny.Set
        d_863_v21_ = _dafny.Set({default__.fm27(globalState)})
        if (d_863_v21_).issubset(d_863_v21_):
            d_864_v22_: _dafny.Array
            nw122_ = _dafny.Array(_dafny.CodePoint('D'), 19)
            d_864_v22_ = nw122_
            index138_ = default__.safeIndex(586, (d_864_v22_).length(0))
            (d_864_v22_)[index138_] = default__.fm38(p2, globalState)
            d_865_v23_: _dafny.Seq
            d_865_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnltn"))
            d_866_v24_: _dafny.Seq
            d_866_v24_ = _dafny.SeqWithoutIsStrInference([p2])
            index139_ = default__.safeIndex(405, (d_836_v0_).length(0))
            (d_836_v0_)[index139_] = (d_866_v24_)[default__.safeIndex((p1).cardinality, len(d_866_v24_))]
            d_867_v25_: bool
            d_867_v25_ = True
            d_868_v26_: str
            d_868_v26_ = _dafny.CodePoint('f')
            d_869_v27_: _dafny.Seq
            d_869_v27_ = _dafny.SeqWithoutIsStrInference([p0])
            index140_ = default__.safeIndex(586, (d_864_v22_).length(0))
            index141_ = default__.safeIndex(405, (d_836_v0_).length(0))
            rhs192_ = d_868_v26_
            rhs193_ = d_865_v23_
            rhs194_ = not(not((self).fm8((d_869_v27_)[default__.safeIndex(p0, len(d_869_v27_))], (d_867_v25_ if p2 else d_867_v25_), p0, globalState)))
            rhs195_ = d_867_v25_
            lhs136_ = d_864_v22_
            lhs137_ = default__.safeIndex(586, (d_864_v22_).length(0))
            lhs138_ = d_836_v0_
            lhs139_ = default__.safeIndex(405, (d_836_v0_).length(0))
            lhs136_[lhs137_] = rhs192_
            d_865_v23_ = rhs193_
            lhs138_[lhs139_] = rhs194_
            d_867_v25_ = rhs195_
            d_870_v28_: _dafny.Array
            nw123_ = _dafny.Array(D1.default()(), 5)
            d_870_v28_ = nw123_
            d_871_v29_: D1
            d_871_v29_ = D1_DC2(p0)
            index142_ = default__.safeIndex(272, (d_870_v28_).length(0))
            (d_870_v28_)[index142_] = d_871_v29_
            index143_ = default__.safeIndex(272, (d_870_v28_).length(0))
            (d_870_v28_)[index143_] = d_871_v29_
            if p2:
                d_865_v23_ = d_865_v23_
                d_872_v30_: _dafny.Set
                d_872_v30_ = _dafny.Set({(D11_DC30(p2, p0)).cf47})
                d_873_v31_: _dafny.Seq
                d_873_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0, (_dafny.MultiSet([494, p0])).cardinality, p0}), d_872_v30_, d_872_v30_, d_872_v30_])
                d_874_v32_: _dafny.Set
                d_874_v32_ = _dafny.Set({default__.fm26(globalState), not(p2)})
                rhs196_ = False
                rhs197_ = default__.fm30(p0, p2, (d_873_v31_)[default__.safeIndex(p0, len(d_873_v31_))], not(not(not((d_874_v32_).ispropersubset(d_874_v32_)))), globalState)
                d_867_v25_ = rhs196_
                d_865_v23_ = rhs197_
                index144_ = default__.safeIndex(405, (d_836_v0_).length(0))
                (d_836_v0_)[index144_] = not(p2)
                d_875_v33_: _dafny.Seq
                d_875_v33_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfny"))])
                d_876_v34_: D4
                d_876_v34_ = D4_DC11(d_875_v33_)
                pat_let_tv11_ = d_875_v33_
                d_877_v35_: _dafny.Map
                def iife53_(_pat_let8_0):
                    def iife54_(d_878_dt__update__tmp_h0_):
                        def iife55_(_pat_let9_0):
                            def iife56_(d_879_dt__update_hcf15_h0_):
                                return D4_DC11(d_879_dt__update_hcf15_h0_)
                            return iife56_(_pat_let9_0)
                        return iife55_(pat_let_tv11_)
                    return iife54_(_pat_let8_0)
                d_877_v35_ = _dafny.Map({(d_836_v0_)[default__.safeIndex(405, (d_836_v0_).length(0))]: iife53_(d_876_v34_)})
                d_880_v36_: T1
                nw124_ = C2()
                nw124_.ctor__()
                d_880_v36_ = nw124_
                d_881_v37_: _dafny.Seq
                d_881_v37_ = _dafny.SeqWithoutIsStrInference([d_880_v36_, d_880_v36_])
                (globalState).f13 = (self).fm23(d_877_v35_, default__.safeDivisionInt(len((d_875_v33_)[default__.safeIndex(p0, len(d_875_v33_))]), len(d_881_v37_)), p0, globalState)
                d_867_v25_ = (((p1).cardinality) - (p0)) > (p0)
            elif True:
                d_882_v38_: _dafny.Seq
                d_882_v38_ = _dafny.SeqWithoutIsStrInference([d_865_v23_])
                d_883_v39_: _dafny.Map
                d_883_v39_ = _dafny.Map({p2: D4_DC11(d_882_v38_)})
                d_884_v40_: _dafny.Set
                d_884_v40_ = _dafny.Set({p0, (self).fm23(d_883_v39_, len(d_865_v23_), -193, globalState), len(_dafny.SeqWithoutIsStrInference([p0])), p0, p0})
                d_882_v38_ = ((_dafny.SeqWithoutIsStrInference([d_865_v23_ for d_885_i1_ in range(default__.abs(950))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_865_v23_ for d_886_i1_ in range(default__.abs(950))]))), default__.fm30(len(_dafny.SeqWithoutIsStrInference([278])), not(p2), d_884_v40_, d_867_v25_, globalState))) + (d_882_v38_)
                (globalState).f4 = 310
                d_887_v41_: C2
                nw125_ = C2()
                nw125_.ctor__()
                d_887_v41_ = nw125_
                d_865_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsu"))
                d_888_v42_: _dafny.Array
                nw126_ = _dafny.Array(None, 5)
                nw126_[int(0)] = (d_836_v0_)[default__.safeIndex(405, (d_836_v0_).length(0))]
                nw126_[int(1)] = p2
                nw126_[int(2)] = d_867_v25_
                nw126_[int(3)] = p2
                nw126_[int(4)] = not(not (p2) or (d_867_v25_))
                d_888_v42_ = nw126_
                rhs198_ = (d_888_v42_ if d_867_v25_ else d_888_v42_)
                rhs199_ = len(d_866_v24_)
                lhs140_ = globalState
                d_888_v42_ = rhs198_
                lhs140_.f15 = rhs199_
            (globalState).f13 = p0
            d_889_v43_: _dafny.Array
            nw127_ = _dafny.Array(D4.default()(), 27)
            d_889_v43_ = nw127_
            d_890_v44_: D4
            d_890_v44_ = D4_DC12(p0)
            index145_ = default__.safeIndex(581, (d_889_v43_).length(0))
            (d_889_v43_)[index145_] = d_890_v44_
            index146_ = default__.safeIndex(581, (d_889_v43_).length(0))
            (d_889_v43_)[index146_] = d_890_v44_
        elif True:
            d_891_v45_: str
            d_891_v45_ = _dafny.CodePoint('y')
            d_892_v46_: _dafny.Map
            d_892_v46_ = _dafny.Map({d_891_v45_: False})
            d_893_v47_: _dafny.Seq
            d_893_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbe"))
            d_892_v46_ = (d_892_v46_).set(d_891_v45_, (len(d_893_v47_)) >= (p0))
            (globalState).f13 = (-853) * (p0)
            d_894_v48_: _dafny.Map
            d_894_v48_ = _dafny.Map({((0) - (p0)) > ((0) - (p0)): p2})
            d_894_v48_ = (d_894_v48_).set(p2, p2)
            d_895_v49_: _dafny.Seq
            d_895_v49_ = _dafny.SeqWithoutIsStrInference([(True if p2 else p2)])
            d_896_v50_: _dafny.Set
            d_896_v50_ = _dafny.Set({p0})
            d_897_v51_: _dafny.Array
            nw128_ = _dafny.Array(None, 3)
            nw128_[int(0)] = 784
            nw128_[int(1)] = (0) - (default__.safeModuloInt(p0, len(d_896_v50_)))
            nw128_[int(2)] = p0
            d_897_v51_ = nw128_
            index147_ = default__.safeIndex(405, (d_897_v51_).length(0))
            (d_897_v51_)[index147_] = (-129) + (p0)
            index148_ = default__.safeIndex(405, (d_897_v51_).length(0))
            rhs200_ = d_895_v49_
            rhs201_ = 886
            lhs141_ = d_897_v51_
            lhs142_ = default__.safeIndex(405, (d_897_v51_).length(0))
            d_895_v49_ = rhs200_
            lhs141_[lhs142_] = rhs201_
            d_898_v52_: bool
            d_898_v52_ = False
            def iife57_():
                coll37_ = _dafny.Map()
                compr_37_: int
                for compr_37_ in _dafny.IntegerRange(745, -795):
                    d_899_v53_: int = compr_37_
                    if ((745) <= (d_899_v53_)) and ((d_899_v53_) < (-795)):
                        coll37_[default__.safeModuloInt(d_899_v53_, (d_897_v51_)[default__.safeIndex(405, (d_897_v51_).length(0))])] = d_893_v47_
                return _dafny.Map(coll37_)
            d_898_v52_ = not(((0) - ((0) - (len(iife57_()
            )))) >= ((d_897_v51_)[default__.safeIndex(405, (d_897_v51_).length(0))]))
        d_900_v54_: _dafny.Map
        d_900_v54_ = _dafny.Map({p0: -152})
        d_901_v55_: D4
        d_901_v55_ = D4_DC12(len(d_900_v54_))
        d_902_v56_: _dafny.Seq
        d_902_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpfda"))
        d_903_v57_: _dafny.Set
        d_903_v57_ = _dafny.Set({len(d_902_v56_)})
        d_904_v58_: _dafny.MultiSet
        d_904_v58_ = _dafny.MultiSet([d_903_v57_])
        d_905_v59_: _dafny.Set
        d_905_v59_ = _dafny.Set({p2, p2})
        d_906_v60_: D7
        d_906_v60_ = D7_DC22(True, p0, p2, p0)
        d_907_v61_: _dafny.MultiSet
        d_907_v61_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0])])
        d_908_v62_: _dafny.Array
        nw129_ = _dafny.Array(None, 23)
        nw129_[int(0)] = p0
        nw129_[int(1)] = 805
        nw129_[int(2)] = p0
        nw129_[int(3)] = (0) - ((d_901_v55_).cf16)
        nw129_[int(4)] = len(_dafny.Map({(p1).cardinality: p2}))
        nw129_[int(5)] = p0
        nw129_[int(6)] = p0
        nw129_[int(7)] = (d_904_v58_).cardinality
        nw129_[int(8)] = p0
        nw129_[int(9)] = len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([((d_862_v20_)[p0] if (p0) in (d_862_v20_) else p2), (D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_909_i4_ in range(default__.abs(-685))]), not(p2))).cf4, p2]) for d_910_i3_ in range(default__.abs(658))])).set(default__.safeIndex(len(d_905_v59_), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([((d_862_v20_)[p0] if (p0) in (d_862_v20_) else p2), (D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_911_i4_ in range(default__.abs(-685))]), not(p2))).cf4, p2]) for d_912_i3_ in range(default__.abs(658))]))), _dafny.SeqWithoutIsStrInference([(d_906_v60_).cf37, p2])))
        nw129_[int(10)] = p0
        nw129_[int(11)] = p0
        nw129_[int(12)] = p0
        nw129_[int(13)] = p0
        nw129_[int(14)] = (_dafny.MultiSet([p2])).cardinality
        nw129_[int(15)] = p0
        nw129_[int(16)] = p0
        nw129_[int(17)] = p0
        nw129_[int(18)] = p0
        nw129_[int(19)] = (d_907_v61_).cardinality
        nw129_[int(20)] = p0
        nw129_[int(21)] = p0
        nw129_[int(22)] = p0
        d_908_v62_ = nw129_
        d_913_v63_: _dafny.Seq
        d_913_v63_ = _dafny.SeqWithoutIsStrInference([d_908_v62_, d_908_v62_])
        d_914_i2_: int
        d_914_i2_ = 0
        with _dafny.label("4"):
            while not((d_913_v63_) == (d_913_v63_)):
                with _dafny.c_label("4"):
                    if (d_914_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_914_i2_ = (d_914_i2_) + (1)
                    index149_ = default__.safeIndex(512, (d_908_v62_).length(0))
                    (d_908_v62_)[index149_] = p0
                    d_915_v64_: D14
                    d_915_v64_ = D14_DC35(default__.fm45(globalState))
                    d_916_v65_: _dafny.Seq
                    d_916_v65_ = _dafny.SeqWithoutIsStrInference([d_915_v64_, d_915_v64_])
                    index150_ = default__.safeIndex(512, (d_908_v62_).length(0))
                    (d_908_v62_)[index150_] = len((d_916_v65_ if False else d_916_v65_))
                    if p2:
                        (globalState).f4 = p0
                        d_917_v66_: bool
                        d_917_v66_ = True
                        d_917_v66_ = d_917_v66_
                        d_918_v67_: _dafny.Set
                        d_918_v67_ = _dafny.Set({d_862_v20_, d_862_v20_})
                        d_919_v68_: _dafny.Seq
                        d_919_v68_ = _dafny.SeqWithoutIsStrInference([p2])
                        rhs202_ = not(p2)
                        rhs203_ = d_917_v66_
                        rhs204_ = not(not (((d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))]) <= (len(d_918_v67_))) or ((d_917_v66_) in (d_919_v68_)))
                        d_917_v66_ = rhs202_
                        d_917_v66_ = rhs203_
                        d_917_v66_ = rhs204_
                        d_920_v69_: _dafny.Seq
                        d_920_v69_ = _dafny.SeqWithoutIsStrInference([(d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], 431, (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))]])
                        (globalState).f15 = len((((d_920_v69_).set(default__.safeIndex(p0, len(d_920_v69_)), len(d_900_v54_))) + (d_920_v69_)) + (d_920_v69_))
                        index151_ = default__.safeIndex(346, (d_836_v0_).length(0))
                        (d_836_v0_)[index151_] = not(p2)
                        d_921_v70_: _dafny.Array
                        nw130_ = _dafny.Array(D6.default()(), 11)
                        d_921_v70_ = nw130_
                        d_922_v71_: D6
                        d_922_v71_ = D6_DC16(_dafny.SeqWithoutIsStrInference([len(d_902_v56_), p0]))
                        d_923_v72_: D6
                        d_923_v72_ = D6_DC18(d_922_v71_)
                        index152_ = default__.safeIndex(615, (d_921_v70_).length(0))
                        (d_921_v70_)[index152_] = d_923_v72_
                        index153_ = default__.safeIndex(346, (d_836_v0_).length(0))
                        index154_ = default__.safeIndex(615, (d_921_v70_).length(0))
                        rhs205_ = (d_919_v68_)[default__.safeIndex((d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], len(d_919_v68_))]
                        rhs206_ = d_923_v72_
                        rhs207_ = p2
                        lhs143_ = d_836_v0_
                        lhs144_ = default__.safeIndex(346, (d_836_v0_).length(0))
                        lhs145_ = d_921_v70_
                        lhs146_ = default__.safeIndex(615, (d_921_v70_).length(0))
                        lhs143_[lhs144_] = rhs205_
                        lhs145_[lhs146_] = rhs206_
                        d_917_v66_ = rhs207_
                    elif True:
                        d_924_v73_: bool
                        d_924_v73_ = True
                        d_924_v73_ = p2
                        d_925_v74_: _dafny.Array
                        def lambda48_(d_926_v62_):
                            def lambda49_(d_927_i5_):
                                return default__.safeDivisionInt(d_927_i5_, (d_926_v62_)[default__.safeIndex(512, (d_926_v62_).length(0))])

                            return lambda49_

                        init25_ = lambda48_(d_908_v62_)
                        nw131_ = _dafny.Array(None, 23)
                        for i0_25_ in range(nw131_.length(0)):
                            nw131_[i0_25_] = init25_(i0_25_)
                        d_925_v74_ = nw131_
                        d_928_v75_: D2
                        d_928_v75_ = D2_DC5(d_925_v74_)
                        d_929_v76_: D2
                        d_929_v76_ = D2_DC8(d_928_v75_)
                        d_930_v77_: C3
                        nw132_ = C3()
                        nw132_.ctor__(d_929_v76_, (p0) < (len(d_905_v59_)))
                        d_930_v77_ = nw132_
                        d_931_v78_: _dafny.Array
                        def lambda50_(d_932_v56_):
                            def lambda51_(d_933_i6_):
                                return d_932_v56_

                            return lambda51_

                        init26_ = lambda50_(d_902_v56_)
                        nw133_ = _dafny.Array(None, 23)
                        for i0_26_ in range(nw133_.length(0)):
                            nw133_[i0_26_] = init26_(i0_26_)
                        d_931_v78_ = nw133_
                        index155_ = default__.safeIndex(708, (d_931_v78_).length(0))
                        (d_931_v78_)[index155_] = d_902_v56_
                        d_934_v79_: _dafny.MultiSet
                        d_934_v79_ = _dafny.MultiSet([(d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], 729])
                        d_935_v80_: D8
                        d_935_v80_ = D8_DC24(p0, not (p2) or (((d_862_v20_)[(d_934_v79_).cardinality] if ((d_934_v79_).cardinality) in (d_862_v20_) else d_930_v77_.f25)))
                        d_936_v81_: _dafny.Map
                        d_936_v81_ = _dafny.Map({(False if d_930_v77_.f25 else True): d_836_v0_})
                        index156_ = default__.safeIndex(708, (d_931_v78_).length(0))
                        rhs208_ = d_902_v56_
                        rhs209_ = d_935_v80_
                        rhs210_ = (d_936_v81_) | ((d_936_v81_) | (_dafny.Map({p2: d_836_v0_})))
                        lhs147_ = d_931_v78_
                        lhs148_ = default__.safeIndex(708, (d_931_v78_).length(0))
                        lhs147_[lhs148_] = rhs208_
                        d_935_v80_ = rhs209_
                        d_936_v81_ = rhs210_
                        d_937_v82_: _dafny.Map
                        d_937_v82_ = _dafny.Map({p1: d_925_v74_})
                        d_937_v82_ = (d_937_v82_).set(default__.fm34(p0, p0, d_924_v73_, globalState), d_925_v74_)
                        (globalState).f13 = (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))]
                    d_938_v83_: _dafny.Seq
                    d_938_v83_ = _dafny.SeqWithoutIsStrInference([default__.fm30((d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], p2, d_903_v57_, True, globalState)])
                    d_939_v84_: D4
                    d_939_v84_ = D4_DC11(d_938_v83_)
                    d_940_v85_: _dafny.Map
                    d_940_v85_ = _dafny.Map({True: d_939_v84_})
                    d_941_v86_: _dafny.Seq
                    d_941_v86_ = _dafny.SeqWithoutIsStrInference([(self).fm23(d_940_v85_, (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], p0, globalState)])
                    d_941_v86_ = (d_941_v86_).set(default__.safeIndex(len(default__.fm30(p0, p2, d_903_v57_, p2, globalState)), len(d_941_v86_)), (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))])
                    d_942_v87_: D11
                    d_942_v87_ = D11_DC29()
                    source18_ = d_942_v87_
                    if source18_.is_DC29:
                        d_943_v88_: bool
                        d_943_v88_ = False
                        d_943_v88_ = True
                        (globalState).f7 = ((0) - (p0)) + ((0) - (p0))
                        d_943_v88_ = True
                        d_944_v89_: _dafny.Array
                        nw134_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                        d_944_v89_ = nw134_
                        nw135_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                        d_944_v89_ = nw135_
                    elif source18_.is_DC30:
                        d_945___mcc_h2_ = source18_.cf46
                        d_946___mcc_h3_ = source18_.cf47
                        d_947_cf47_ = d_946___mcc_h3_
                        d_948_cf46_ = d_945___mcc_h2_
                        d_949_v90_: str
                        d_949_v90_ = _dafny.CodePoint('b')
                        d_950_v91_: _dafny.Array
                        nw136_ = _dafny.Array(None, 2)
                        nw136_[int(0)] = (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))]
                        nw136_[int(1)] = len(default__.fm41(d_949_v90_, globalState))
                        d_950_v91_ = nw136_
                        d_951_v92_: D2
                        d_951_v92_ = D2_DC7(d_948_cf46_, p1, d_950_v91_)
                        d_952_v93_: D2
                        d_952_v93_ = D2_DC8(d_951_v92_)
                        pat_let_tv12_ = d_951_v92_
                        d_953_v94_: C3
                        nw137_ = C3()
                        def iife58_(_pat_let10_0):
                            def iife59_(d_954_dt__update__tmp_h1_):
                                def iife60_(_pat_let11_0):
                                    def iife61_(d_955_dt__update_hcf11_h0_):
                                        return D2_DC8(d_955_dt__update_hcf11_h0_)
                                    return iife61_(_pat_let11_0)
                                return iife60_(pat_let_tv12_)
                            return iife59_(_pat_let10_0)
                        nw137_.ctor__(iife58_(d_952_v93_), d_948_cf46_)
                        d_953_v94_ = nw137_
                        d_948_cf46_ = p2
                        d_949_v90_ = d_949_v90_
                        d_956_v95_: _dafny.Array
                        nw138_ = _dafny.Array(_dafny.Seq({}), 11)
                        d_956_v95_ = nw138_
                        d_957_v96_: D7
                        d_957_v96_ = D7_DC21((d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], 508, d_948_cf46_, (self).fm8(p0, d_953_v94_.f25, (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], globalState), d_953_v94_.f25)
                        index157_ = default__.safeIndex(637, (d_956_v95_).length(0))
                        (d_956_v95_)[index157_] = ((d_941_v86_).set(default__.safeIndex((self).fm23(d_940_v85_, (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], d_947_cf47_, globalState), len(d_941_v86_)), d_947_cf47_) if (d_957_v96_).cf32 else d_941_v86_)
                        d_958_v97_: _dafny.Array
                        nw139_ = _dafny.Array(_dafny.Array(None, 0), 17)
                        d_958_v97_ = nw139_
                        index158_ = default__.safeIndex(970, (d_958_v97_).length(0))
                        (d_958_v97_)[index158_] = d_836_v0_
                        d_959_v98_: _dafny.Array
                        nw140_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                        d_959_v98_ = nw140_
                        index159_ = default__.safeIndex(129, (d_959_v98_).length(0))
                        (d_959_v98_)[index159_] = d_902_v56_
                        d_960_v99_: _dafny.Map
                        d_960_v99_ = _dafny.Map({d_836_v0_: d_902_v56_})
                        d_961_v100_: _dafny.Seq
                        d_961_v100_ = _dafny.SeqWithoutIsStrInference([d_941_v86_, d_941_v86_])
                        d_962_v101_: _dafny.Seq
                        d_962_v101_ = _dafny.SeqWithoutIsStrInference([d_836_v0_])
                        index160_ = default__.safeIndex(637, (d_956_v95_).length(0))
                        index161_ = default__.safeIndex(970, (d_958_v97_).length(0))
                        index162_ = default__.safeIndex(129, (d_959_v98_).length(0))
                        rhs211_ = ((((d_941_v86_ if d_953_v94_.f25 else d_941_v86_)).set(default__.safeIndex(d_947_cf47_, len((d_941_v86_ if d_953_v94_.f25 else d_941_v86_))), d_947_cf47_)).set(default__.safeIndex((p1).cardinality, len(((d_941_v86_ if d_953_v94_.f25 else d_941_v86_)).set(default__.safeIndex(d_947_cf47_, len((d_941_v86_ if d_953_v94_.f25 else d_941_v86_))), d_947_cf47_))), len(d_961_v100_))) + (d_941_v86_)
                        rhs212_ = ((d_836_v0_ if p2 else d_836_v0_) if True else (d_836_v0_ if d_948_cf46_ else (d_962_v101_)[default__.safeIndex((d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))], len(d_962_v101_))]))
                        rhs213_ = not (d_948_cf46_) or (d_948_cf46_)
                        rhs214_ = ((d_902_v56_).set(default__.safeIndex(d_947_cf47_, len(d_902_v56_)), default__.fm38(True, globalState)) if d_948_cf46_ else d_902_v56_)
                        rhs215_ = (_dafny.Map({d_836_v0_: d_902_v56_})) | (d_960_v99_)
                        lhs149_ = d_956_v95_
                        lhs150_ = default__.safeIndex(637, (d_956_v95_).length(0))
                        lhs151_ = d_958_v97_
                        lhs152_ = default__.safeIndex(970, (d_958_v97_).length(0))
                        lhs153_ = d_959_v98_
                        lhs154_ = default__.safeIndex(129, (d_959_v98_).length(0))
                        lhs149_[lhs150_] = rhs211_
                        lhs151_[lhs152_] = rhs212_
                        d_948_cf46_ = rhs213_
                        lhs153_[lhs154_] = rhs214_
                        d_960_v99_ = rhs215_
                    elif source18_.is_DC28:
                        d_963___mcc_h4_ = source18_.cf45
                        d_964_cf45_ = d_963___mcc_h4_
                        d_965_v102_: _dafny.Array
                        nw141_ = _dafny.Array(D11.default()(), 20)
                        d_965_v102_ = nw141_
                        index163_ = default__.safeIndex(623, (d_965_v102_).length(0))
                        (d_965_v102_)[index163_] = d_942_v87_
                        index164_ = default__.safeIndex(623, (d_965_v102_).length(0))
                        (d_965_v102_)[index164_] = d_942_v87_
                        d_966_v103_: bool
                        d_966_v103_ = False
                        index165_ = default__.safeIndex(512, (d_908_v62_).length(0))
                        rhs216_ = (d_905_v59_) - ((d_905_v59_) | (d_905_v59_))
                        rhs217_ = d_966_v103_
                        rhs218_ = (p0) + ((p0 if p2 else p0))
                        rhs219_ = 153
                        lhs155_ = globalState
                        lhs156_ = d_908_v62_
                        lhs157_ = default__.safeIndex(512, (d_908_v62_).length(0))
                        d_905_v59_ = rhs216_
                        d_966_v103_ = rhs217_
                        lhs155_.f15 = rhs218_
                        lhs156_[lhs157_] = rhs219_
                        d_966_v103_ = p2
                        d_967_v104_: C1
                        nw142_ = C1()
                        nw142_.ctor__()
                        d_967_v104_ = nw142_
                    elif True:
                        d_968___mcc_h5_ = source18_.cf48
                        d_969_cf48_ = d_968___mcc_h5_
                        d_900_v54_ = d_900_v54_
                        (globalState).f4 = p0
                        d_970_v105_: bool
                        d_970_v105_ = False
                        d_971_v106_: _dafny.Array
                        def lambda52_(d_972_i7_):
                            return (d_972_i7_) * (264)

                        init27_ = lambda52_
                        nw143_ = _dafny.Array(None, 28)
                        for i0_27_ in range(nw143_.length(0)):
                            nw143_[i0_27_] = init27_(i0_27_)
                        d_971_v106_ = nw143_
                        d_973_v107_: D2
                        d_973_v107_ = D2_DC5(d_971_v106_)
                        index166_ = default__.safeIndex(512, (d_908_v62_).length(0))
                        rhs220_ = default__.fm26(globalState)
                        rhs221_ = p0
                        rhs222_ = (d_908_v62_)[default__.safeIndex(512, (d_908_v62_).length(0))]
                        rhs223_ = p3
                        rhs224_ = p0
                        lhs158_ = globalState
                        lhs159_ = globalState
                        lhs160_ = d_908_v62_
                        lhs161_ = default__.safeIndex(512, (d_908_v62_).length(0))
                        d_970_v105_ = rhs220_
                        lhs158_.f13 = rhs221_
                        lhs159_.f13 = rhs222_
                        d_973_v107_ = rhs223_
                        lhs160_[lhs161_] = rhs224_
                        d_974_v108_: str
                        d_974_v108_ = _dafny.CodePoint('p')
                        d_975_v109_: _dafny.Map
                        d_975_v109_ = _dafny.Map({d_974_v108_: d_836_v0_})
                        d_975_v109_ = (d_975_v109_).set(d_974_v108_, d_836_v0_)
                    pass
            pass
        d_976_v110_: bool
        d_976_v110_ = False
        d_976_v110_ = d_976_v110_
        hi6_ = p0
        for d_977_i8_ in range((p0 if False else p0), hi6_):
            d_978_v111_: _dafny.Array
            nw144_ = _dafny.Array(D1.default()(), 19)
            d_978_v111_ = nw144_
            index167_ = default__.safeIndex(312, (d_836_v0_).length(0))
            (d_836_v0_)[index167_] = p2
            d_979_v112_: D2
            d_979_v112_ = D2_DC7(d_976_v110_, p1, d_908_v62_)
            d_980_v113_: _dafny.Set
            d_980_v113_ = _dafny.Set({d_902_v56_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_981_i9_ in range(default__.abs(204))]), d_902_v56_, d_902_v56_, d_902_v56_})
            index168_ = default__.safeIndex(312, (d_836_v0_).length(0))
            rhs225_ = (p2 if (d_902_v56_) == (d_902_v56_) else (d_979_v112_).cf8)
            rhs226_ = d_978_v111_
            rhs227_ = (d_980_v113_).issubset(d_980_v113_)
            rhs228_ = d_976_v110_
            lhs162_ = d_836_v0_
            lhs163_ = default__.safeIndex(312, (d_836_v0_).length(0))
            d_976_v110_ = rhs225_
            d_978_v111_ = rhs226_
            lhs162_[lhs163_] = rhs227_
            d_976_v110_ = rhs228_
            d_982_v114_: _dafny.Seq
            d_982_v114_ = _dafny.SeqWithoutIsStrInference([d_902_v56_, d_902_v56_])
            d_983_v115_: D4
            d_983_v115_ = D4_DC11(d_982_v114_)
            d_984_v116_: _dafny.Map
            d_984_v116_ = _dafny.Map({d_976_v110_: d_983_v115_})
            d_985_v117_: _dafny.MultiSet
            d_985_v117_ = _dafny.MultiSet([(self).fm23(d_984_v116_, d_977_i8_, 330, globalState), (0) - (p0)])
            d_976_v110_ = (d_985_v117_).issubset(d_985_v117_)
            d_986_v118_: _dafny.Seq
            d_986_v118_ = _dafny.SeqWithoutIsStrInference([d_976_v110_, default__.fm26(globalState), d_976_v110_])
            d_986_v118_ = d_986_v118_
            index169_ = default__.safeIndex(312, (d_836_v0_).length(0))
            index170_ = default__.safeIndex(312, (d_836_v0_).length(0))
            rhs229_ = d_976_v110_
            rhs230_ = (d_986_v118_)[default__.safeIndex(p0, len(d_986_v118_))]
            lhs164_ = d_836_v0_
            lhs165_ = default__.safeIndex(312, (d_836_v0_).length(0))
            lhs166_ = d_836_v0_
            lhs167_ = default__.safeIndex(312, (d_836_v0_).length(0))
            lhs164_[lhs165_] = rhs229_
            lhs166_[lhs167_] = rhs230_


class C5(T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def fm7(self, p0, globalState):
        def iife62_():
            def iife66_():
                def iife68_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in _dafny.IntegerRange(741, 602):
                        d_987_v0_: int = compr_44_
                        if ((741) <= (d_987_v0_)) and ((d_987_v0_) < (602)):
                            coll44_[(d_987_v0_) - (len(_dafny.SeqWithoutIsStrInference([-193])))] = 587
                    return _dafny.Map(coll44_)
                coll42_ = _dafny.Set()
                def iife67_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(741, 602):
                        d_987_v0_: int = compr_43_
                        if ((741) <= (d_987_v0_)) and ((d_987_v0_) < (602)):
                            coll43_[(d_987_v0_) - (len(_dafny.SeqWithoutIsStrInference([-193])))] = 587
                    return _dafny.Map(coll43_)
                compr_42_: int
                for compr_42_ in (iife67_()
                ).keys.Elements:
                    d_990_v1_: int = compr_42_
                    if (d_990_v1_) in (iife68_()
                    ):
                        coll42_ = coll42_.union(_dafny.Set([default__.safeDivisionInt(d_990_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "we"))))]))
                return _dafny.Set(coll42_)
            coll38_ = _dafny.Set()
            def iife63_():
                def iife65_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(741, 602):
                        d_987_v0_: int = compr_41_
                        if ((741) <= (d_987_v0_)) and ((d_987_v0_) < (602)):
                            coll41_[(d_987_v0_) - (len(_dafny.SeqWithoutIsStrInference([-193])))] = 587
                    return _dafny.Map(coll41_)
                coll39_ = _dafny.Set()
                def iife64_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(741, 602):
                        d_987_v0_: int = compr_40_
                        if ((741) <= (d_987_v0_)) and ((d_987_v0_) < (602)):
                            coll40_[(d_987_v0_) - (len(_dafny.SeqWithoutIsStrInference([-193])))] = 587
                    return _dafny.Map(coll40_)
                compr_39_: int
                for compr_39_ in (iife64_()
                ).keys.Elements:
                    d_988_v1_: int = compr_39_
                    if (d_988_v1_) in (iife65_()
                    ):
                        coll39_ = coll39_.union(_dafny.Set([default__.safeDivisionInt(d_988_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "we"))))]))
                return _dafny.Set(coll39_)
            compr_38_: int
            for compr_38_ in (iife63_()
            ).Elements:
                d_989_v2_: int = compr_38_
                if (d_989_v2_) in (iife66_()
                ):
                    coll38_ = coll38_.union(_dafny.Set([default__.safeDivisionInt(d_989_v2_, 769)]))
            return _dafny.Set(coll38_)
        def iife69_():
            coll45_ = _dafny.Set()
            compr_45_: int
            for compr_45_ in _dafny.IntegerRange(374, 894):
                d_991_v3_: int = compr_45_
                if ((374) <= (d_991_v3_)) and ((d_991_v3_) < (894)):
                    coll45_ = coll45_.union(_dafny.Set([(d_991_v3_) * (245)]))
            return _dafny.Set(coll45_)
        def iife70_():
            coll46_ = _dafny.Set()
            compr_46_: int
            for compr_46_ in _dafny.IntegerRange(666, 713):
                d_992_v4_: int = compr_46_
                if ((666) <= (d_992_v4_)) and ((d_992_v4_) < (713)):
                    coll46_ = coll46_.union(_dafny.Set([(d_992_v4_) * (144)]))
            return _dafny.Set(coll46_)
        def iife71_():
            coll47_ = _dafny.Map()
            compr_47_: int
            for compr_47_ in _dafny.IntegerRange(337, -99):
                d_993_v5_: int = compr_47_
                if ((337) <= (d_993_v5_)) and ((d_993_v5_) < (-99)):
                    coll47_[(d_993_v5_) + (len(_dafny.Map({False: 500})))] = 853
            return _dafny.Map(coll47_)
        def iife72_():
            coll48_ = _dafny.Set()
            compr_48_: int
            for compr_48_ in (_dafny.MultiSet([215, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))])).Elements:
                d_994_v6_: int = compr_48_
                if (d_994_v6_) in (_dafny.MultiSet([215, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))])):
                    coll48_ = coll48_.union(_dafny.Set([default__.safeDivisionInt(d_994_v6_, 772)]))
            return _dafny.Set(coll48_)
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({180}), _dafny.Set({-596})]))).intersection(_dafny.MultiSet([iife62_()
        , iife69_()
        , iife70_()
        ]))) - ((_dafny.MultiSet([_dafny.Set({412, 183, len(_dafny.Map({len(iife71_()
        ): -161}))}), iife72_()
        , _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jghlcq"))), len(_dafny.Set({False}))})])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')]))).cardinality}), _dafny.Set({138})]))))

    def fm8(self, p0, p1, p2, globalState):
        return ((99) * (952)) > (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False}))])))

    def fm20(self, globalState):
        return not(((not(not(False))) and (False) if (_dafny.SeqWithoutIsStrInference([True, True])) <= (_dafny.SeqWithoutIsStrInference([False])) else (_dafny.MultiSet([(0) - (-758)])).issubset(_dafny.MultiSet([785, 488, 825]))))

    def m6(self, p0, p1, p2, p3, globalState):
        (globalState).f15 = p0
        if p2:
            d_995_v0_: bool
            d_995_v0_ = False
            d_996_v1_: str
            d_996_v1_ = _dafny.CodePoint('g')
            d_997_v2_: _dafny.Set
            d_997_v2_ = _dafny.Set({d_995_v0_})
            d_998_v3_: _dafny.Array
            def lambda53_(d_999_p0_):
                def lambda54_(d_1000_i0_):
                    return default__.safeDivisionInt(d_1000_i0_, d_999_p0_)

                return lambda54_

            init28_ = lambda53_(p0)
            nw145_ = _dafny.Array(None, 28)
            for i0_28_ in range(nw145_.length(0)):
                nw145_[i0_28_] = init28_(i0_28_)
            d_998_v3_ = nw145_
            d_1001_v4_: _dafny.Array
            nw146_ = _dafny.Array(None, 23)
            nw146_[int(0)] = True
            nw146_[int(1)] = (True) == (d_995_v0_)
            nw146_[int(2)] = d_995_v0_
            nw146_[int(3)] = (default__.fm21(_dafny.CodePoint('m'), d_996_v1_, globalState)).cf4
            nw146_[int(4)] = p2
            nw146_[int(5)] = p2
            nw146_[int(6)] = (d_995_v0_ if True else True)
            nw146_[int(7)] = d_995_v0_
            nw146_[int(8)] = (self).fm20(globalState)
            nw146_[int(9)] = (p1).isdisjoint(p1)
            nw146_[int(10)] = p2
            nw146_[int(11)] = p2
            nw146_[int(12)] = p2
            nw146_[int(13)] = (d_997_v2_).issubset(d_997_v2_)
            nw146_[int(14)] = p2
            nw146_[int(15)] = d_995_v0_
            nw146_[int(16)] = (_dafny.MultiSet([d_998_v3_, d_998_v3_])).ispropersubset((_dafny.MultiSet([d_998_v3_])).set(d_998_v3_, default__.abs((0) - (p0))))
            nw146_[int(17)] = p2
            nw146_[int(18)] = (self).fm20(globalState)
            nw146_[int(19)] = False
            nw146_[int(20)] = not(p2)
            nw146_[int(21)] = p2
            nw146_[int(22)] = (self).fm8(p0, p2, p0, globalState)
            d_1001_v4_ = nw146_
            d_1002_v5_: _dafny.Set
            d_1002_v5_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))})
            index171_ = default__.safeIndex(670, (d_1001_v4_).length(0))
            (d_1001_v4_)[index171_] = (d_1002_v5_).issubset(d_1002_v5_)
            d_1003_v6_: _dafny.Seq
            d_1003_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slclxhejc"))
            d_1004_v7_: _dafny.Seq
            d_1004_v7_ = _dafny.SeqWithoutIsStrInference([d_1003_v6_, (d_1003_v6_).set(default__.safeIndex(p0, len(d_1003_v6_)), _dafny.CodePoint('x'))])
            d_1005_v8_: D4
            d_1005_v8_ = D4_DC11(d_1004_v7_)
            index172_ = default__.safeIndex(670, (d_1001_v4_).length(0))
            rhs231_ = p2
            rhs232_ = p2
            rhs233_ = (d_1005_v8_).cf15
            lhs168_ = d_1001_v4_
            lhs169_ = default__.safeIndex(670, (d_1001_v4_).length(0))
            d_995_v0_ = rhs231_
            lhs168_[lhs169_] = rhs232_
            d_1004_v7_ = rhs233_
            index173_ = default__.safeIndex(876, (d_998_v3_).length(0))
            (d_998_v3_)[index173_] = p0
            index174_ = default__.safeIndex(876, (d_998_v3_).length(0))
            (d_998_v3_)[index174_] = (p0 if (d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))] else p0)
            index175_ = default__.safeIndex(670, (d_1001_v4_).length(0))
            rhs234_ = ((d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))]) < (default__.safeDivisionInt((d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))], (d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))]))
            rhs235_ = d_996_v1_
            lhs170_ = d_1001_v4_
            lhs171_ = default__.safeIndex(670, (d_1001_v4_).length(0))
            lhs170_[lhs171_] = rhs234_
            d_996_v1_ = rhs235_
            d_1006_v9_: _dafny.Map
            d_1006_v9_ = _dafny.Map({p2: d_995_v0_})
            d_1007_v10_: _dafny.Map
            d_1007_v10_ = _dafny.Map({(d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))]: d_998_v3_})
            d_1008_v11_: _dafny.Map
            d_1008_v11_ = _dafny.Map({(0) - (len(d_1007_v10_)): d_1006_v9_})
            d_1009_v12_: _dafny.Array
            nw147_ = _dafny.Array(None, 17)
            nw147_[int(0)] = d_1006_v9_
            nw147_[int(1)] = (d_1006_v9_) | (d_1006_v9_)
            nw147_[int(2)] = (d_1006_v9_) | (_dafny.Map({p2: p2}))
            nw147_[int(3)] = _dafny.Map({(d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))]: d_995_v0_})
            nw147_[int(4)] = d_1006_v9_
            nw147_[int(5)] = (d_1006_v9_) | (d_1006_v9_)
            nw147_[int(6)] = _dafny.Map({(d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))]: p2})
            nw147_[int(7)] = d_1006_v9_
            nw147_[int(8)] = (_dafny.Map({p2: (d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))]})).set(p2, (d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))])
            nw147_[int(9)] = (d_1006_v9_) | (d_1006_v9_)
            nw147_[int(10)] = d_1006_v9_
            nw147_[int(11)] = (d_1006_v9_) | (d_1006_v9_)
            nw147_[int(12)] = d_1006_v9_
            nw147_[int(13)] = (default__.fm22(d_1004_v7_, (d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))], globalState) if d_995_v0_ else ((d_1008_v11_)[p0] if (p0) in (d_1008_v11_) else d_1006_v9_))
            nw147_[int(14)] = d_1006_v9_
            nw147_[int(15)] = d_1006_v9_
            nw147_[int(16)] = d_1006_v9_
            d_1009_v12_ = nw147_
            index176_ = default__.safeIndex(90, (d_1009_v12_).length(0))
            (d_1009_v12_)[index176_] = (_dafny.Map({d_995_v0_: (d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))]}) if False else default__.fm22(d_1004_v7_, p0, globalState))
            index177_ = default__.safeIndex(90, (d_1009_v12_).length(0))
            rhs236_ = not(p2)
            rhs237_ = (d_1006_v9_) | (d_1006_v9_)
            rhs238_ = (d_1001_v4_)[default__.safeIndex(670, (d_1001_v4_).length(0))]
            lhs172_ = d_1009_v12_
            lhs173_ = default__.safeIndex(90, (d_1009_v12_).length(0))
            d_995_v0_ = rhs236_
            lhs172_[lhs173_] = rhs237_
            d_995_v0_ = rhs238_
            d_1010_v13_: D2
            d_1010_v13_ = D2_DC6(((d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))]) + ((d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))]))
            pat_let_tv13_ = p0
            def iife73_(_pat_let12_0):
                def iife74_(d_1011_dt__update__tmp_h0_):
                    def iife75_(_pat_let13_0):
                        def iife76_(d_1012_dt__update_hcf7_h0_):
                            return D2_DC6(d_1012_dt__update_hcf7_h0_)
                        return iife76_(_pat_let13_0)
                    return iife75_(pat_let_tv13_)
                return iife74_(_pat_let12_0)
            d_1010_v13_ = iife73_(D2_DC6((d_998_v3_)[default__.safeIndex(876, (d_998_v3_).length(0))]))
        elif True:
            d_1013_v14_: _dafny.Array
            nw148_ = _dafny.Array(int(0), 14)
            d_1013_v14_ = nw148_
            d_1014_v15_: D2
            d_1014_v15_ = D2_DC8(D2_DC5(d_1013_v14_))
            d_1015_v16_: D2
            d_1015_v16_ = D2_DC8(d_1014_v15_)
            pat_let_tv14_ = d_1014_v15_
            d_1016_v17_: C3
            nw149_ = C3()
            def iife77_(_pat_let14_0):
                def iife78_(d_1017_dt__update__tmp_h1_):
                    def iife79_(_pat_let15_0):
                        def iife80_(d_1018_dt__update_hcf11_h0_):
                            return D2_DC8(d_1018_dt__update_hcf11_h0_)
                        return iife80_(_pat_let15_0)
                    return iife79_(pat_let_tv14_)
                return iife78_(_pat_let14_0)
            nw149_.ctor__(iife77_(d_1015_v16_), p2)
            d_1016_v17_ = nw149_
            d_1019_v18_: _dafny.Seq
            d_1019_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            d_1020_v19_: _dafny.Set
            d_1020_v19_ = _dafny.Set({p0})
            d_1021_v20_: _dafny.MultiSet
            d_1021_v20_ = _dafny.MultiSet([len(d_1020_v19_), p0])
            d_1022_v21_: _dafny.Seq
            d_1022_v21_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1023_v22_: _dafny.Array
            nw150_ = _dafny.Array(None, 8)
            nw150_[int(0)] = d_1021_v20_
            nw150_[int(1)] = d_1021_v20_
            nw150_[int(2)] = (d_1021_v20_) - (_dafny.MultiSet(d_1022_v21_))
            nw150_[int(3)] = d_1021_v20_
            nw150_[int(4)] = d_1021_v20_
            nw150_[int(5)] = (d_1021_v20_) | (d_1021_v20_)
            nw150_[int(6)] = (d_1021_v20_).intersection((d_1021_v20_).set(p0, default__.abs(p0)))
            nw150_[int(7)] = d_1021_v20_
            d_1023_v22_ = nw150_
            d_1024_v23_: _dafny.Seq
            d_1024_v23_ = _dafny.SeqWithoutIsStrInference([d_1019_v18_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1025_i1_ in range(default__.abs(790))])])
            d_1026_v24_: _dafny.Map
            d_1026_v24_ = _dafny.Map({p0: d_1016_v17_.f25})
            d_1027_v25_: _dafny.Map
            d_1027_v25_ = _dafny.Map({len(d_1026_v24_): d_1016_v17_.f25})
            d_1028_v26_: D11
            d_1028_v26_ = D11_DC30(d_1016_v17_.f25, len(d_1026_v24_))
            d_1029_v27_: _dafny.Map
            d_1029_v27_ = _dafny.Map({(d_1028_v26_).cf46: p0})
            d_1030_v28_: _dafny.Seq
            d_1030_v28_ = _dafny.SeqWithoutIsStrInference([d_1016_v17_.f25, p2])
            d_1031_v29_: str
            d_1031_v29_ = _dafny.CodePoint('o')
            rhs239_ = (d_1019_v18_) + ((d_1024_v23_)[default__.safeIndex(p0, len(d_1024_v23_))])
            rhs240_ = d_1023_v22_
            rhs241_ = len(default__.fm39((len((d_1029_v27_).set(default__.fm26(globalState), -592))) - (len(d_1030_v28_)), d_1031_v29_, (_dafny.SeqWithoutIsStrInference([p0 for d_1032_i2_ in range(default__.abs(524))])) + (d_1022_v21_), p2, globalState))
            rhs242_ = p0
            lhs174_ = globalState
            lhs175_ = globalState
            d_1019_v18_ = rhs239_
            d_1023_v22_ = rhs240_
            lhs174_.f4 = rhs241_
            lhs175_.f13 = rhs242_
            d_1033_v30_: _dafny.Map
            d_1033_v30_ = _dafny.Map({p2: p2})
            d_1034_v31_: _dafny.MultiSet
            d_1034_v31_ = _dafny.MultiSet([d_1033_v30_])
            d_1013_v14_ = (d_1013_v14_ if (d_1034_v31_).ispropersubset(d_1034_v31_) else d_1013_v14_)
            if ((_dafny.SeqWithoutIsStrInference([d_1031_v29_ for d_1035_i3_ in range(default__.abs(374))])) + (d_1019_v18_)) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1036_i4_ in range(default__.abs(111))])):
                index178_ = default__.safeIndex(385, (d_1013_v14_).length(0))
                (d_1013_v14_)[index178_] = ((d_1016_v17_).fm2(len(_dafny.Map({d_1016_v17_.f25: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xs"))})), globalState)) + (len(d_1030_v28_))
                index179_ = default__.safeIndex(385, (d_1013_v14_).length(0))
                (d_1013_v14_)[index179_] = p0
                (d_1016_v17_).f25 = (d_1016_v17_.f25) == ((d_1030_v28_) == (_dafny.SeqWithoutIsStrInference([d_1016_v17_.f25])))
                (d_1016_v17_).f25 = ((d_1013_v14_)[default__.safeIndex(385, (d_1013_v14_).length(0))]) < ((d_1013_v14_)[default__.safeIndex(385, (d_1013_v14_).length(0))])
                (d_1016_v17_).f25 = True
                d_1033_v30_ = (d_1033_v30_).set(False, (d_1016_v17_.f25) or (p2))
            elif True:
                d_1031_v29_ = default__.fm38((d_1016_v17_.f25) == (p2), globalState)
                d_1031_v29_ = d_1031_v29_
                d_1037_v32_: _dafny.Set
                d_1037_v32_ = _dafny.Set({(d_1030_v28_)[default__.safeIndex(p0, len(d_1030_v28_))]})
                (globalState).f4 = (0) - ((len(d_1033_v30_)) + ((len(d_1019_v18_)) + ((_dafny.MultiSet([len(d_1037_v32_)])).cardinality)))
                (globalState).f13 = p0
                (d_1016_v17_).f25 = (p2 if p2 else d_1016_v17_.f25)
            (d_1016_v17_).f25 = p2
        d_1038_v33_: _dafny.Seq
        d_1038_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phf"))
        d_1039_v34_: _dafny.Map
        d_1039_v34_ = _dafny.Map({d_1038_v33_: p2})
        d_1040_i5_: int
        d_1040_i5_ = 0
        with _dafny.label("5"):
            while not(((d_1039_v34_)[d_1038_v33_] if (d_1038_v33_) in (d_1039_v34_) else p2)):
                with _dafny.c_label("5"):
                    if (d_1040_i5_) >= (100):
                        raise _dafny.Break("5")
                    d_1040_i5_ = (d_1040_i5_) + (1)
                    d_1041_v35_: _dafny.Map
                    d_1041_v35_ = _dafny.Map({175: p0})
                    d_1041_v35_ = (d_1041_v35_).set(default__.safeModuloInt(p0, p0), p0)
                    d_1041_v35_ = d_1041_v35_
                    d_1042_v36_: str
                    d_1042_v36_ = _dafny.CodePoint('r')
                    d_1042_v36_ = d_1042_v36_
                    d_1043_v37_: _dafny.Set
                    d_1043_v37_ = _dafny.Set({p2})
                    d_1044_v38_: _dafny.Array
                    nw151_ = _dafny.Array(None, 17)
                    nw151_[int(0)] = d_1043_v37_
                    nw151_[int(1)] = d_1043_v37_
                    nw151_[int(2)] = d_1043_v37_
                    nw151_[int(3)] = default__.fm37(globalState)
                    nw151_[int(4)] = d_1043_v37_
                    nw151_[int(5)] = (d_1043_v37_) | (default__.fm37(globalState))
                    nw151_[int(6)] = (d_1043_v37_ if p2 else d_1043_v37_)
                    nw151_[int(7)] = d_1043_v37_
                    nw151_[int(8)] = d_1043_v37_
                    nw151_[int(9)] = _dafny.Set({p2})
                    nw151_[int(10)] = _dafny.Set({p2})
                    nw151_[int(11)] = d_1043_v37_
                    nw151_[int(12)] = d_1043_v37_
                    nw151_[int(13)] = _dafny.Set({not(p2)})
                    nw151_[int(14)] = d_1043_v37_
                    nw151_[int(15)] = d_1043_v37_
                    nw151_[int(16)] = _dafny.Set({not(p2), p2, False, p2})
                    d_1044_v38_ = nw151_
                    index180_ = default__.safeIndex(208, (d_1044_v38_).length(0))
                    (d_1044_v38_)[index180_] = d_1043_v37_
                    index181_ = default__.safeIndex(208, (d_1044_v38_).length(0))
                    (d_1044_v38_)[index181_] = d_1043_v37_
                    pass
            pass
        d_1045_v39_: _dafny.Array
        nw152_ = _dafny.Array(False, 12)
        d_1045_v39_ = nw152_
        index182_ = default__.safeIndex(817, (d_1045_v39_).length(0))
        (d_1045_v39_)[index182_] = True
        d_1046_v40_: bool
        d_1046_v40_ = False
        d_1047_v41_: _dafny.Map
        d_1047_v41_ = _dafny.Map({True: p0})
        d_1048_v42_: _dafny.Map
        d_1048_v42_ = d_1047_v41_
        index183_ = default__.safeIndex(817, (d_1045_v39_).length(0))
        rhs243_ = (p0) + (p0)
        rhs244_ = d_1046_v40_
        rhs245_ = not(p2)
        rhs246_ = d_1048_v42_
        lhs176_ = globalState
        lhs177_ = d_1045_v39_
        lhs178_ = default__.safeIndex(817, (d_1045_v39_).length(0))
        lhs176_.f13 = rhs243_
        lhs177_[lhs178_] = rhs244_
        d_1046_v40_ = rhs245_
        d_1048_v42_ = rhs246_
        d_1049_v43_: _dafny.Array
        nw153_ = _dafny.Array(int(0), 20)
        d_1049_v43_ = nw153_
        d_1049_v43_ = d_1049_v43_
        index184_ = default__.safeIndex(817, (d_1045_v39_).length(0))
        (d_1045_v39_)[index184_] = p2


class C6:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm15(self, p0, p1, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1050_i0_ in range(default__.abs(353))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmtigl"))})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggpkkfhx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mo")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhtp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhv"))}))) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1051_i1_ in range(default__.abs(-201))])}))

    def fm16(self, p0, globalState):
        return False

    def m9(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: bool = False
        d_1052_v0_: _dafny.MultiSet
        d_1052_v0_ = _dafny.MultiSet([p1, (self).fm16(p0, globalState)])
        d_1053_v1_: _dafny.Array
        nw154_ = _dafny.Array(int(0), 8)
        d_1053_v1_ = nw154_
        d_1054_v2_: D2
        d_1054_v2_ = D2_DC7(True, d_1052_v0_, d_1053_v1_)
        d_1055_v3_: _dafny.MultiSet
        d_1055_v3_ = _dafny.MultiSet([d_1054_v2_, d_1054_v2_, d_1054_v2_, d_1054_v2_, d_1054_v2_])
        d_1056_v4_: _dafny.Seq
        d_1056_v4_ = _dafny.SeqWithoutIsStrInference([d_1055_v3_])
        d_1057_i0_: int
        d_1057_i0_ = 0
        with _dafny.label("6"):
            while ((d_1056_v4_)[default__.safeIndex(p0, len(d_1056_v4_))]).issubset(d_1055_v3_):
                with _dafny.c_label("6"):
                    if (d_1057_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1057_i0_ = (d_1057_i0_) + (1)
                    d_1058_v5_: str
                    d_1058_v5_ = _dafny.CodePoint('b')
                    d_1058_v5_ = _dafny.CodePoint('x')
                    (globalState).f4 = p0
                    r2 = (p0) != (p0)
                    d_1059_v6_: _dafny.Seq
                    d_1059_v6_ = _dafny.SeqWithoutIsStrInference([default__.fm17(p0, globalState), p0])
                    d_1060_v7_: _dafny.Map
                    d_1060_v7_ = _dafny.Map({p1: p0})
                    d_1061_v8_: _dafny.Set
                    d_1061_v8_ = _dafny.Set({(d_1059_v6_)[default__.safeIndex(len(d_1060_v7_), len(d_1059_v6_))]})
                    r2 = (d_1061_v8_).isdisjoint(d_1061_v8_)
                    pass
            pass
        d_1062_v9_: _dafny.Array
        nw155_ = _dafny.Array(_dafny.Map({}), 17)
        d_1062_v9_ = nw155_
        d_1063_v10_: _dafny.Map
        d_1063_v10_ = _dafny.Map({p0: not(p1)})
        index185_ = default__.safeIndex(998, (d_1062_v9_).length(0))
        def iife81_():
            coll49_ = _dafny.Map()
            compr_49_: int
            for compr_49_ in _dafny.IntegerRange(693, 711):
                d_1064_v11_: int = compr_49_
                if ((693) <= (d_1064_v11_)) and ((d_1064_v11_) < (711)):
                    coll49_[default__.safeModuloInt(d_1064_v11_, (0) - (p0))] = p1
            return _dafny.Map(coll49_)
        (d_1062_v9_)[index185_] = (d_1063_v10_) | (iife81_()
        )
        d_1065_v12_: str
        d_1065_v12_ = _dafny.CodePoint('h')
        d_1066_v13_: _dafny.MultiSet
        d_1066_v13_ = _dafny.MultiSet([d_1065_v12_, d_1065_v12_])
        d_1067_v14_: _dafny.Map
        d_1067_v14_ = _dafny.Map({p0: d_1066_v13_})
        d_1068_v15_: _dafny.Seq
        d_1068_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvqumpwp"))
        index186_ = default__.safeIndex(998, (d_1062_v9_).length(0))
        (d_1062_v9_)[index186_] = (default__.fm18(p0, 643, ((d_1067_v14_)[len(d_1068_v15_)] if (len(d_1068_v15_)) in (d_1067_v14_) else d_1066_v13_), globalState)).set(p0, (p2) == (p1))
        d_1069_v16_: _dafny.Array
        nw156_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
        d_1069_v16_ = nw156_
        index187_ = default__.safeIndex(736, (d_1069_v16_).length(0))
        (d_1069_v16_)[index187_] = (default__.fm19(p1, p1, globalState)) + (d_1068_v15_)
        d_1070_v17_: _dafny.Array
        nw157_ = _dafny.Array(None, 14)
        nw157_[int(0)] = d_1053_v1_
        nw157_[int(1)] = d_1053_v1_
        nw157_[int(2)] = d_1053_v1_
        nw157_[int(3)] = d_1053_v1_
        nw157_[int(4)] = d_1053_v1_
        nw157_[int(5)] = d_1053_v1_
        nw157_[int(6)] = d_1053_v1_
        nw157_[int(7)] = d_1053_v1_
        nw157_[int(8)] = d_1053_v1_
        nw157_[int(9)] = d_1053_v1_
        nw157_[int(10)] = d_1053_v1_
        nw157_[int(11)] = d_1053_v1_
        nw157_[int(12)] = d_1053_v1_
        nw157_[int(13)] = d_1053_v1_
        d_1070_v17_ = nw157_
        index188_ = default__.safeIndex(306, (d_1070_v17_).length(0))
        (d_1070_v17_)[index188_] = d_1053_v1_
        d_1071_v18_: _dafny.Map
        d_1071_v18_ = _dafny.Map({D1_DC2(p0): d_1068_v15_})
        d_1072_v19_: D1
        d_1072_v19_ = D1_DC2(p0)
        d_1073_v20_: _dafny.Seq
        d_1073_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1065_v12_ for d_1074_i1_ in range(default__.abs(-938))])])
        d_1075_v21_: _dafny.Map
        d_1075_v21_ = _dafny.Map({False: p1})
        d_1076_v22_: _dafny.Map
        d_1076_v22_ = _dafny.Map({(p2) and (p1): d_1053_v1_})
        index189_ = default__.safeIndex(736, (d_1069_v16_).length(0))
        index190_ = default__.safeIndex(306, (d_1070_v17_).length(0))
        rhs247_ = (p0) + (p0)
        rhs248_ = (((d_1071_v18_)[d_1072_v19_] if (d_1072_v19_) in (d_1071_v18_) else d_1068_v15_)).set(default__.safeIndex(p0, len(((d_1071_v18_)[d_1072_v19_] if (d_1072_v19_) in (d_1071_v18_) else d_1068_v15_))), d_1065_v12_)
        rhs249_ = ((d_1068_v15_) + (d_1068_v15_)) + ((d_1073_v20_)[default__.safeIndex(len(d_1075_v21_), len(d_1073_v20_))])
        rhs250_ = ((d_1076_v22_)[(p0) > (p0)] if ((p0) > (p0)) in (d_1076_v22_) else d_1053_v1_)
        lhs179_ = globalState
        lhs180_ = d_1069_v16_
        lhs181_ = default__.safeIndex(736, (d_1069_v16_).length(0))
        lhs182_ = d_1070_v17_
        lhs183_ = default__.safeIndex(306, (d_1070_v17_).length(0))
        lhs179_.f15 = rhs247_
        lhs180_[lhs181_] = rhs248_
        d_1068_v15_ = rhs249_
        lhs182_[lhs183_] = rhs250_
        if False:
            d_1077_v23_: _dafny.Array
            nw158_ = _dafny.Array(False, 5)
            d_1077_v23_ = nw158_
            index191_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            (d_1077_v23_)[index191_] = (d_1065_v12_) in (d_1068_v15_)
            index192_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            (d_1077_v23_)[index192_] = p1
            d_1078_v24_: D0
            d_1078_v24_ = D0_DC1(d_1077_v23_)
            d_1078_v24_ = d_1078_v24_
            d_1079_v25_: _dafny.Seq
            d_1079_v25_ = _dafny.SeqWithoutIsStrInference([(d_1077_v23_)[default__.safeIndex(0, (d_1077_v23_).length(0))], False])
            d_1080_v26_: _dafny.Map
            d_1080_v26_ = _dafny.Map({d_1079_v25_: (False if (d_1077_v23_)[default__.safeIndex(0, (d_1077_v23_).length(0))] else p1)})
            d_1081_v27_: _dafny.Seq
            d_1081_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1065_v12_})])
            index193_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            (d_1077_v23_)[index193_] = ((d_1080_v26_)[_dafny.SeqWithoutIsStrInference([p1])] if (_dafny.SeqWithoutIsStrInference([p1])) in (d_1080_v26_) else (((d_1062_v9_)[default__.safeIndex(998, (d_1062_v9_).length(0))])[(0) - (len(d_1081_v27_))] if ((0) - (len(d_1081_v27_))) in ((d_1062_v9_)[default__.safeIndex(998, (d_1062_v9_).length(0))]) else p1))
            d_1082_v28_: D0
            d_1082_v28_ = D0_DC0(d_1077_v23_)
            index194_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            index195_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            rhs251_ = (p0) * (p0)
            rhs252_ = ((d_1063_v10_)[p0] if (p0) in (d_1063_v10_) else p1)
            rhs253_ = d_1082_v28_
            rhs254_ = not(p2)
            rhs255_ = (d_1079_v25_) + (d_1079_v25_)
            lhs184_ = globalState
            lhs185_ = d_1077_v23_
            lhs186_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            lhs187_ = d_1077_v23_
            lhs188_ = default__.safeIndex(0, (d_1077_v23_).length(0))
            lhs184_.f15 = rhs251_
            lhs185_[lhs186_] = rhs252_
            d_1082_v28_ = rhs253_
            lhs187_[lhs188_] = rhs254_
            d_1079_v25_ = rhs255_
            d_1083_v29_: D1
            d_1083_v29_ = D1_DC3(d_1068_v15_, p1)
            if not ((d_1083_v29_).cf4) or ((d_1077_v23_)[default__.safeIndex(0, (d_1077_v23_).length(0))]):
                r2 = (d_1077_v23_)[default__.safeIndex(0, (d_1077_v23_).length(0))]
                d_1084_v30_: _dafny.Seq
                d_1084_v30_ = _dafny.SeqWithoutIsStrInference([d_1077_v23_])
                d_1085_v31_: _dafny.Set
                d_1085_v31_ = _dafny.Set({len(d_1084_v30_)})
                d_1086_v32_: _dafny.Map
                d_1086_v32_ = _dafny.Map({d_1085_v31_: p1})
                d_1087_v33_: _dafny.Map
                d_1087_v33_ = _dafny.Map({p1: p0})
                r2 = ((d_1086_v32_)[(_dafny.Set({p0, ((d_1052_v0_)[p1] if (p1) in (d_1052_v0_) else p0)})) - (_dafny.Set({len(d_1087_v33_), 721}))] if ((_dafny.Set({p0, ((d_1052_v0_)[p1] if (p1) in (d_1052_v0_) else p0)})) - (_dafny.Set({len(d_1087_v33_), 721}))) in (d_1086_v32_) else not((d_1077_v23_)[default__.safeIndex(0, (d_1077_v23_).length(0))]))
                index196_ = default__.safeIndex(609, (d_1053_v1_).length(0))
                (d_1053_v1_)[index196_] = default__.safeModuloInt(p0, p0)
                index197_ = default__.safeIndex(609, (d_1053_v1_).length(0))
                (d_1053_v1_)[index197_] = p0
                d_1068_v15_ = ((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))]) + (d_1068_v15_)
                d_1073_v20_ = d_1073_v20_
            elif True:
                d_1088_v34_: D3
                d_1088_v34_ = D3_DC9(d_1065_v12_)
                d_1065_v12_ = (d_1088_v34_).cf12
                index198_ = default__.safeIndex(0, (d_1077_v23_).length(0))
                rhs256_ = p0
                rhs257_ = p0
                rhs258_ = p1
                lhs189_ = globalState
                lhs190_ = globalState
                lhs191_ = d_1077_v23_
                lhs192_ = default__.safeIndex(0, (d_1077_v23_).length(0))
                lhs189_.f4 = rhs256_
                lhs190_.f4 = rhs257_
                lhs191_[lhs192_] = rhs258_
                d_1089_v35_: C0
                nw159_ = C0()
                nw159_.ctor__()
                d_1089_v35_ = nw159_
                index199_ = default__.safeIndex(0, (d_1077_v23_).length(0))
                (d_1077_v23_)[index199_] = (default__.safeModuloInt(p0, p0)) == (p0)
                d_1090_v36_: D15
                d_1090_v36_ = D15_DC38(d_1063_v10_)
                d_1091_v37_: _dafny.MultiSet
                d_1091_v37_ = _dafny.MultiSet([(d_1090_v36_).cf55, (_dafny.Map({(0) - (p0): p2}) if p1 else default__.fm35(globalState)), (d_1090_v36_).cf55])
                rhs259_ = d_1091_v37_
                rhs260_ = (d_1077_v23_)[default__.safeIndex(0, (d_1077_v23_).length(0))]
                d_1091_v37_ = rhs259_
                r2 = rhs260_
        elif True:
            if ((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))]) <= (((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))]).set(default__.safeIndex(p0, len((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))])), d_1065_v12_)):
                (globalState).f15 = p0
                d_1092_v38_: C4
                nw160_ = C4()
                nw160_.ctor__()
                d_1092_v38_ = nw160_
                d_1093_v39_: _dafny.Seq
                d_1093_v39_ = _dafny.SeqWithoutIsStrInference([p2, (d_1092_v38_).fm8((_dafny.MultiSet([p0])).cardinality, p1, p0, globalState), p2, p2])
                d_1052_v0_ = (default__.fm34(603, len(d_1093_v39_), p2, globalState)) | (d_1052_v0_)
                r2 = p2
                d_1094_v40_: _dafny.Seq
                d_1094_v40_ = _dafny.SeqWithoutIsStrInference([d_1062_v9_])
                d_1095_v41_: _dafny.Map
                d_1095_v41_ = _dafny.Map({p0: d_1062_v9_})
                d_1096_v42_: _dafny.Set
                d_1096_v42_ = _dafny.Set({False})
                d_1097_v43_: _dafny.Array
                nw161_ = _dafny.Array(None, 28)
                nw161_[int(0)] = d_1062_v9_
                nw161_[int(1)] = d_1062_v9_
                nw161_[int(2)] = d_1062_v9_
                nw161_[int(3)] = d_1062_v9_
                nw161_[int(4)] = d_1062_v9_
                nw161_[int(5)] = d_1062_v9_
                nw161_[int(6)] = d_1062_v9_
                nw161_[int(7)] = d_1062_v9_
                nw161_[int(8)] = d_1062_v9_
                nw161_[int(9)] = d_1062_v9_
                nw161_[int(10)] = d_1062_v9_
                nw161_[int(11)] = (d_1094_v40_)[default__.safeIndex(p0, len(d_1094_v40_))]
                nw161_[int(12)] = d_1062_v9_
                nw161_[int(13)] = d_1062_v9_
                nw161_[int(14)] = d_1062_v9_
                nw161_[int(15)] = d_1062_v9_
                nw161_[int(16)] = d_1062_v9_
                nw161_[int(17)] = d_1062_v9_
                nw161_[int(18)] = (d_1062_v9_ if p2 else ((d_1095_v41_)[len(d_1096_v42_)] if (len(d_1096_v42_)) in (d_1095_v41_) else (d_1094_v40_)[default__.safeIndex(p0, len(d_1094_v40_))]))
                nw161_[int(19)] = d_1062_v9_
                nw161_[int(20)] = d_1062_v9_
                nw161_[int(21)] = (d_1062_v9_ if False else d_1062_v9_)
                nw161_[int(22)] = d_1062_v9_
                nw161_[int(23)] = d_1062_v9_
                nw161_[int(24)] = d_1062_v9_
                nw161_[int(25)] = d_1062_v9_
                nw161_[int(26)] = d_1062_v9_
                nw161_[int(27)] = d_1062_v9_
                d_1097_v43_ = nw161_
                index200_ = default__.safeIndex(188, (d_1097_v43_).length(0))
                (d_1097_v43_)[index200_] = d_1062_v9_
                index201_ = default__.safeIndex(188, (d_1097_v43_).length(0))
                (d_1097_v43_)[index201_] = d_1062_v9_
            elif True:
                (globalState).f4 = default__.safeModuloInt(p0, 887)
                r2 = not(((p0) * (-924)) != (default__.safeDivisionInt(p0, p0)))
                d_1098_v44_: _dafny.Seq
                d_1098_v44_ = _dafny.SeqWithoutIsStrInference([len(default__.fm33(p1, p1, p0, globalState)), p0])
                r2 = ((d_1098_v44_) + (_dafny.SeqWithoutIsStrInference([535 for d_1099_i2_ in range(default__.abs(-375))]))) <= (d_1098_v44_)
                d_1100_v45_: _dafny.Seq
                d_1100_v45_ = _dafny.SeqWithoutIsStrInference([not (not(p2)) or (p2)])
                d_1101_v46_: _dafny.Seq
                d_1101_v46_ = _dafny.SeqWithoutIsStrInference([d_1075_v21_])
                d_1102_v47_: _dafny.Seq
                d_1102_v47_ = _dafny.SeqWithoutIsStrInference([d_1101_v46_])
                d_1103_v48_: _dafny.MultiSet
                d_1103_v48_ = _dafny.MultiSet([(_dafny.MultiSet((d_1102_v47_)[default__.safeIndex(len((d_1062_v9_)[default__.safeIndex(998, (d_1062_v9_).length(0))]), len(d_1102_v47_))])).cardinality])
                rhs261_ = p0
                rhs262_ = _dafny.SeqWithoutIsStrInference([p2, p1, not((d_1103_v48_).issubset(d_1103_v48_))])
                lhs193_ = globalState
                lhs193_.f15 = rhs261_
                d_1100_v45_ = rhs262_
                r2 = False
            d_1104_v49_: _dafny.MultiSet
            d_1104_v49_ = _dafny.MultiSet([p0, p0])
            d_1105_v50_: _dafny.Map
            d_1105_v50_ = _dafny.Map({((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))])[default__.safeIndex(p0, len((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))]))]: d_1104_v49_})
            d_1105_v50_ = (d_1105_v50_).set(d_1065_v12_, (d_1104_v49_) | (d_1104_v49_))
            if default__.fm26(globalState):
                index202_ = default__.safeIndex(862, (d_1053_v1_).length(0))
                (d_1053_v1_)[index202_] = p0
                index203_ = default__.safeIndex(862, (d_1053_v1_).length(0))
                (d_1053_v1_)[index203_] = (default__.safeDivisionInt(len(_dafny.Map({d_1053_v1_: default__.fm26(globalState)})), p0)) - (default__.fm17(p0, globalState))
                r2 = p2
                r2 = (p1) == ((p1 if p1 else p1))
                index204_ = default__.safeIndex(862, (d_1053_v1_).length(0))
                (d_1053_v1_)[index204_] = (p0) * ((d_1053_v1_)[default__.safeIndex(862, (d_1053_v1_).length(0))])
                r2 = True
            elif True:
                d_1106_v51_: C2
                nw162_ = C2()
                nw162_.ctor__()
                d_1106_v51_ = nw162_
                d_1107_v52_: _dafny.Array
                def lambda55_(d_1108_p1_):
                    def lambda56_(d_1109_i3_):
                        return d_1108_p1_

                    return lambda56_

                init29_ = lambda55_(p1)
                nw163_ = _dafny.Array(None, 19)
                for i0_29_ in range(nw163_.length(0)):
                    nw163_[i0_29_] = init29_(i0_29_)
                d_1107_v52_ = nw163_
                d_1110_v53_: _dafny.Map
                d_1110_v53_ = _dafny.Map({p1: len(_dafny.SeqWithoutIsStrInference([d_1107_v52_]))})
                (globalState).f13 = (0) - (default__.safeDivisionInt(len(d_1110_v53_), len(_dafny.SeqWithoutIsStrInference([(D3_DC9(d_1065_v12_)).cf12 for d_1111_i4_ in range(default__.abs(-592))]))))
                r2 = p2
                index205_ = default__.safeIndex(842, (d_1107_v52_).length(0))
                (d_1107_v52_)[index205_] = p1
                index206_ = default__.safeIndex(842, (d_1107_v52_).length(0))
                (d_1107_v52_)[index206_] = not (p2) or (p2)
                arr0_ = (d_1070_v17_)[default__.safeIndex(306, (d_1070_v17_).length(0))]
                index207_ = default__.safeIndex(200, ((d_1070_v17_)[default__.safeIndex(306, (d_1070_v17_).length(0))]).length(0))
                arr0_[index207_] = len(d_1068_v15_)
                arr1_ = (d_1070_v17_)[default__.safeIndex(306, (d_1070_v17_).length(0))]
                index208_ = default__.safeIndex(200, ((d_1070_v17_)[default__.safeIndex(306, (d_1070_v17_).length(0))]).length(0))
                arr1_[index208_] = p0
            if True:
                r2 = ((d_1075_v21_)[not(p2)] if (not(p2)) in (d_1075_v21_) else (d_1052_v0_).ispropersubset(d_1052_v0_))
                d_1112_v54_: _dafny.Map
                d_1112_v54_ = _dafny.Map({p0: p0})
                d_1113_v57_: _dafny.Array
                nw164_ = _dafny.Array(None, 14)
                nw164_[int(0)] = _dafny.Map({387: p0})
                nw164_[int(1)] = d_1112_v54_
                nw164_[int(2)] = d_1112_v54_
                nw164_[int(3)] = d_1112_v54_
                nw164_[int(4)] = (d_1112_v54_).set((0) - (p0), p0)
                nw164_[int(5)] = d_1112_v54_
                nw164_[int(6)] = d_1112_v54_
                nw164_[int(7)] = d_1112_v54_
                nw164_[int(8)] = ((d_1112_v54_).set(p0, 733)) | (d_1112_v54_)
                nw164_[int(9)] = (d_1112_v54_ if p2 else d_1112_v54_)
                def iife82_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(337, 833):
                        d_1114_v55_: int = compr_50_
                        if ((337) <= (d_1114_v55_)) and ((d_1114_v55_) < (833)):
                            coll50_[(d_1114_v55_) - (p0)] = p0
                    return _dafny.Map(coll50_)
                nw164_[int(10)] = iife82_()
                
                def iife83_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(132, 774):
                        d_1115_v56_: int = compr_51_
                        if ((132) <= (d_1115_v56_)) and ((d_1115_v56_) < (774)):
                            coll51_[default__.safeDivisionInt(d_1115_v56_, p0)] = p0
                    return _dafny.Map(coll51_)
                nw164_[int(11)] = iife83_()
                
                nw164_[int(12)] = d_1112_v54_
                nw164_[int(13)] = d_1112_v54_
                d_1113_v57_ = nw164_
                d_1113_v57_ = d_1113_v57_
                d_1116_v58_: _dafny.Array
                nw165_ = _dafny.Array(False, 18)
                d_1116_v58_ = nw165_
                index209_ = default__.safeIndex(353, (d_1116_v58_).length(0))
                (d_1116_v58_)[index209_] = (self).fm16(p0, globalState)
                d_1117_v59_: _dafny.Seq
                d_1117_v59_ = _dafny.SeqWithoutIsStrInference([False, p1])
                d_1118_v60_: _dafny.Seq
                d_1118_v60_ = _dafny.SeqWithoutIsStrInference([d_1104_v49_, _dafny.MultiSet([p0, 209])])
                index210_ = default__.safeIndex(353, (d_1116_v58_).length(0))
                (d_1116_v58_)[index210_] = (((d_1104_v49_).set(len(d_1117_v59_), default__.abs(p0))).intersection(_dafny.MultiSet([p0]))) in (_dafny.MultiSet(d_1118_v60_))
                index211_ = default__.safeIndex(353, (d_1116_v58_).length(0))
                (d_1116_v58_)[index211_] = p2
                index212_ = default__.safeIndex(353, (d_1116_v58_).length(0))
                (d_1116_v58_)[index212_] = p1
            elif True:
                index213_ = default__.safeIndex(736, (d_1069_v16_).length(0))
                (d_1069_v16_)[index213_] = (_dafny.SeqWithoutIsStrInference([d_1065_v12_ for d_1119_i5_ in range(default__.abs(-551))])) + ((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))])
                d_1120_v61_: D4
                d_1120_v61_ = D4_DC12(p0)
                d_1121_v62_: _dafny.MultiSet
                d_1121_v62_ = _dafny.MultiSet([(d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))]])
                d_1122_v63_: _dafny.Map
                d_1122_v63_ = _dafny.Map({d_1120_v61_: (d_1121_v62_).cardinality})
                d_1122_v63_ = (d_1122_v63_).set(D4_DC12(p0), p0)
                r2 = p2
                (globalState).f15 = (0) - (p0)
                index214_ = default__.safeIndex(87, (d_1053_v1_).length(0))
                (d_1053_v1_)[index214_] = p0
                d_1123_v64_: _dafny.Array
                def lambda57_(d_1124_p1_):
                    def lambda58_(d_1125_i6_):
                        return (d_1124_p1_ if d_1124_p1_ else d_1124_p1_)

                    return lambda58_

                init30_ = lambda57_(p1)
                nw166_ = _dafny.Array(None, 9)
                for i0_30_ in range(nw166_.length(0)):
                    nw166_[i0_30_] = init30_(i0_30_)
                d_1123_v64_ = nw166_
                index215_ = default__.safeIndex(97, (d_1123_v64_).length(0))
                (d_1123_v64_)[index215_] = (p2) == (p1)
                d_1126_v65_: _dafny.Map
                d_1126_v65_ = _dafny.Map({p1: p0})
                d_1127_v66_: _dafny.Seq
                d_1127_v66_ = _dafny.SeqWithoutIsStrInference([885, p0])
                index216_ = default__.safeIndex(87, (d_1053_v1_).length(0))
                index217_ = default__.safeIndex(97, (d_1123_v64_).length(0))
                rhs263_ = len((d_1126_v65_) | (d_1126_v65_))
                rhs264_ = (_dafny.MultiSet([not(p1)])).set(default__.fm26(globalState), default__.abs((p0) * ((d_1127_v66_)[default__.safeIndex(p0, len(d_1127_v66_))])))
                rhs265_ = (((d_1052_v0_).set(p1, default__.abs(p0))).intersection(d_1052_v0_)).issubset(_dafny.MultiSet([p1]))
                rhs266_ = p0
                lhs194_ = d_1053_v1_
                lhs195_ = default__.safeIndex(87, (d_1053_v1_).length(0))
                lhs196_ = d_1123_v64_
                lhs197_ = default__.safeIndex(97, (d_1123_v64_).length(0))
                lhs198_ = globalState
                lhs194_[lhs195_] = rhs263_
                d_1052_v0_ = rhs264_
                lhs196_[lhs197_] = rhs265_
                lhs198_.f4 = rhs266_
            (globalState).f15 = p0
        d_1128_v67_: D2
        d_1128_v67_ = D2_DC7(p1, d_1052_v0_, (d_1070_v17_)[default__.safeIndex(306, (d_1070_v17_).length(0))])
        d_1129_v68_: D2
        d_1129_v68_ = D2_DC8(d_1128_v67_)
        d_1130_v69_: C3
        nw167_ = C3()
        nw167_.ctor__(d_1129_v68_, False)
        d_1130_v69_ = nw167_
        hi7_ = p0
        for d_1131_i7_ in range(p0, hi7_):
            index218_ = default__.safeIndex(457, (d_1053_v1_).length(0))
            (d_1053_v1_)[index218_] = -834
            d_1132_v70_: _dafny.Set
            d_1132_v70_ = _dafny.Set({d_1131_i7_})
            d_1133_v71_: _dafny.Seq
            d_1133_v71_ = _dafny.SeqWithoutIsStrInference([d_1132_v70_])
            index219_ = default__.safeIndex(457, (d_1053_v1_).length(0))
            rhs267_ = p0
            rhs268_ = (0) - (len(d_1133_v71_))
            lhs199_ = d_1053_v1_
            lhs200_ = default__.safeIndex(457, (d_1053_v1_).length(0))
            lhs201_ = globalState
            lhs199_[lhs200_] = rhs267_
            lhs201_.f15 = rhs268_
            r2 = d_1130_v69_.f25
            r2 = not(d_1130_v69_.f25)
            d_1134_v72_: _dafny.MultiSet
            d_1134_v72_ = _dafny.MultiSet([p0, (d_1053_v1_)[default__.safeIndex(457, (d_1053_v1_).length(0))]])
            (globalState).f4 = ((d_1134_v72_)[p0] if (p0) in (d_1134_v72_) else p0)
        r0 = (d_1130_v69_).fm2((0) - (default__.safeModuloInt(p0, len((d_1062_v9_)[default__.safeIndex(998, (d_1062_v9_).length(0))]))), globalState)
        d_1135_v73_: _dafny.Map
        d_1135_v73_ = _dafny.Map({p0: p0})
        d_1136_v74_: _dafny.Seq
        d_1136_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([((d_1135_v73_)[p0] if (p0) in (d_1135_v73_) else p0)])])
        d_1137_v75_: _dafny.Set
        d_1137_v75_ = _dafny.Set({_dafny.Map({p0: len((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))])})})
        d_1138_v76_: _dafny.MultiSet
        d_1138_v76_ = _dafny.MultiSet([len(d_1137_v75_), len((d_1069_v16_)[default__.safeIndex(736, (d_1069_v16_).length(0))])])
        r1 = ((d_1136_v74_)[default__.safeIndex(761, len(d_1136_v74_))]) - ((d_1138_v76_).set(p0, default__.abs(p0)))
        d_1139_v77_: D15
        d_1139_v77_ = D15_DC38((d_1062_v9_)[default__.safeIndex(998, (d_1062_v9_).length(0))])
        d_1140_v78_: D15
        d_1140_v78_ = D15_DC40(d_1139_v77_)
        d_1141_v79_: D15
        d_1141_v79_ = D15_DC40(d_1139_v77_)
        d_1142_v80_: D15
        d_1142_v80_ = D15_DC40(d_1139_v77_)
        pat_let_tv15_ = d_1130_v69_
        pat_let_tv16_ = d_1130_v69_
        pat_let_tv17_ = d_1130_v69_
        def lambda59_(source19_):
            if source19_.is_DC39:
                d_1143___mcc_h0_ = source19_.cf56
                d_1144___mcc_h1_ = source19_.cf57
                d_1145___mcc_h2_ = source19_.cf58
                d_1146_cf58_ = d_1145___mcc_h2_
                d_1147_cf57_ = d_1144___mcc_h1_
                d_1148_cf56_ = d_1143___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference([d_1147_cf57_])) == (_dafny.SeqWithoutIsStrInference([pat_let_tv15_.f25]))
            elif source19_.is_DC38:
                d_1149___mcc_h3_ = source19_.cf55
                d_1150_cf55_ = d_1149___mcc_h3_
                return pat_let_tv16_.f25
            elif True:
                d_1151___mcc_h4_ = source19_.cf59
                d_1152_cf59_ = d_1151___mcc_h4_
                return pat_let_tv17_.f25

        r2 = lambda59_(d_1142_v80_)
        return r0, r1, r2


class C7(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, globalState):
        return (_dafny.CodePoint('w')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bppvmww")))

    def fm2(self, p0, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, not(False), True}))]) if True else _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "po")))) for d_1153_i1_ in range(default__.abs(888))]))) for d_1154_i0_ in range(default__.abs(293))]))).cardinality

    def fm12(self, globalState):
        return (True) == (not((D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtaill")), True)).cf4))

    def fm13(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkonhsjjq"))) + (_dafny.SeqWithoutIsStrInference([(D3_DC9(_dafny.CodePoint('l'))).cf12 for d_1155_i0_ in range(default__.abs(542))]))

    def m1(self, p0, p1, globalState):
        d_1156_v0_: bool
        d_1156_v0_ = True
        d_1157_v1_: D1
        d_1157_v1_ = D1_DC3(p0, d_1156_v0_)
        d_1156_v0_ = (d_1156_v0_ if (d_1157_v1_).cf4 else d_1156_v0_)
        d_1158_v2_: int
        d_1158_v2_ = 792
        d_1159_v3_: _dafny.Map
        d_1159_v3_ = _dafny.Map({d_1158_v2_: False})
        d_1160_v4_: _dafny.Array
        nw168_ = _dafny.Array(None, 5)
        nw168_[int(0)] = (0) - (d_1158_v2_)
        nw168_[int(1)] = (0) - ((0) - (len(d_1159_v3_)))
        nw168_[int(2)] = d_1158_v2_
        nw168_[int(3)] = d_1158_v2_
        nw168_[int(4)] = d_1158_v2_
        d_1160_v4_ = nw168_
        d_1161_v5_: _dafny.Map
        d_1161_v5_ = _dafny.Map({d_1160_v4_: len(_dafny.SeqWithoutIsStrInference([d_1158_v2_]))})
        d_1161_v5_ = d_1161_v5_
        d_1162_v6_: _dafny.Map
        d_1162_v6_ = _dafny.Map({d_1156_v0_: d_1158_v2_})
        d_1162_v6_ = d_1162_v6_
        d_1163_v7_: _dafny.MultiSet
        d_1163_v7_ = _dafny.MultiSet([d_1156_v0_])
        d_1164_v8_: _dafny.Seq
        d_1164_v8_ = _dafny.SeqWithoutIsStrInference([78])
        d_1165_i0_: int
        d_1165_i0_ = 0
        with _dafny.label("7"):
            while not(not((((d_1163_v7_)[not(d_1156_v0_)] if (not(d_1156_v0_)) in (d_1163_v7_) else (d_1164_v8_)[default__.safeIndex(d_1158_v2_, len(d_1164_v8_))])) >= ((d_1158_v2_ if d_1156_v0_ else d_1158_v2_)))):
                with _dafny.c_label("7"):
                    if (d_1165_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_1165_i0_ = (d_1165_i0_) + (1)
                    d_1166_v9_: _dafny.Seq
                    d_1166_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkph"))
                    d_1167_v10_: str
                    d_1167_v10_ = _dafny.CodePoint('a')
                    d_1168_v11_: _dafny.Seq
                    d_1168_v11_ = _dafny.SeqWithoutIsStrInference([d_1163_v7_, d_1163_v7_, d_1163_v7_])
                    rhs269_ = (d_1168_v11_)[default__.safeIndex(d_1158_v2_, len(d_1168_v11_))]
                    rhs270_ = (d_1166_v9_ if d_1156_v0_ else d_1166_v9_)
                    rhs271_ = d_1167_v10_
                    rhs272_ = len(d_1166_v9_)
                    d_1163_v7_ = rhs269_
                    d_1166_v9_ = rhs270_
                    d_1167_v10_ = rhs271_
                    d_1158_v2_ = rhs272_
                    d_1169_v12_: _dafny.Map
                    d_1169_v12_ = _dafny.Map({d_1156_v0_: _dafny.SeqWithoutIsStrInference([d_1158_v2_, d_1158_v2_])})
                    d_1170_v13_: _dafny.Map
                    d_1170_v13_ = _dafny.Map({(((d_1169_v12_)[d_1156_v0_] if (d_1156_v0_) in (d_1169_v12_) else d_1164_v8_)) + (d_1164_v8_): d_1156_v0_})
                    d_1171_v14_: _dafny.MultiSet
                    d_1171_v14_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1167_v10_ for d_1172_i1_ in range(default__.abs(-69))])), d_1158_v2_])
                    if ((d_1170_v13_)[d_1164_v8_] if (d_1164_v8_) in (d_1170_v13_) else (d_1171_v14_).issubset(d_1171_v14_)):
                        rhs273_ = default__.safeModuloInt((d_1158_v2_) * (d_1158_v2_), d_1158_v2_)
                        rhs274_ = (d_1164_v8_)[default__.safeIndex(d_1158_v2_, len(d_1164_v8_))]
                        lhs202_ = globalState
                        lhs202_.f7 = rhs273_
                        d_1158_v2_ = rhs274_
                        d_1160_v4_ = (d_1160_v4_ if d_1156_v0_ else d_1160_v4_)
                        (globalState).f4 = (0) - (d_1158_v2_)
                        (globalState).f15 = (0) - (d_1158_v2_)
                        (globalState).f4 = (self).fm2(d_1158_v2_, globalState)
                    elif True:
                        d_1173_v15_: _dafny.Map
                        d_1173_v15_ = _dafny.Map({d_1156_v0_: d_1156_v0_})
                        d_1156_v0_ = ((d_1173_v15_)[False] if (False) in (d_1173_v15_) else d_1156_v0_)
                        d_1174_v16_: _dafny.Seq
                        d_1174_v16_ = _dafny.SeqWithoutIsStrInference([d_1156_v0_, d_1156_v0_])
                        d_1175_v17_: _dafny.Map
                        d_1175_v17_ = _dafny.Map({d_1156_v0_: d_1173_v15_})
                        d_1176_v18_: _dafny.Array
                        nw169_ = _dafny.Array(None, 27)
                        nw169_[int(0)] = d_1156_v0_
                        nw169_[int(1)] = d_1156_v0_
                        nw169_[int(2)] = False
                        nw169_[int(3)] = d_1156_v0_
                        nw169_[int(4)] = d_1156_v0_
                        nw169_[int(5)] = d_1156_v0_
                        nw169_[int(6)] = not(False)
                        nw169_[int(7)] = d_1156_v0_
                        nw169_[int(8)] = d_1156_v0_
                        nw169_[int(9)] = d_1156_v0_
                        nw169_[int(10)] = not(d_1156_v0_)
                        nw169_[int(11)] = d_1156_v0_
                        nw169_[int(12)] = (d_1174_v16_)[default__.safeIndex(d_1158_v2_, len(d_1174_v16_))]
                        nw169_[int(13)] = d_1156_v0_
                        nw169_[int(14)] = True
                        nw169_[int(15)] = d_1156_v0_
                        nw169_[int(16)] = d_1156_v0_
                        nw169_[int(17)] = d_1156_v0_
                        nw169_[int(18)] = (self).fm1(d_1175_v17_, len(_dafny.Map({d_1158_v2_: d_1158_v2_})), globalState)
                        nw169_[int(19)] = d_1156_v0_
                        nw169_[int(20)] = d_1156_v0_
                        nw169_[int(21)] = (d_1174_v16_)[default__.safeIndex(d_1158_v2_, len(d_1174_v16_))]
                        nw169_[int(22)] = d_1156_v0_
                        nw169_[int(23)] = d_1156_v0_
                        nw169_[int(24)] = d_1156_v0_
                        nw169_[int(25)] = True
                        nw169_[int(26)] = ((d_1159_v3_)[len(_dafny.Set({d_1158_v2_, d_1158_v2_, d_1158_v2_}))] if (len(_dafny.Set({d_1158_v2_, d_1158_v2_, d_1158_v2_}))) in (d_1159_v3_) else d_1156_v0_)
                        d_1176_v18_ = nw169_
                        d_1177_v19_: _dafny.Map
                        d_1177_v19_ = _dafny.Map({d_1176_v18_: (not(not(False)) if d_1156_v0_ else d_1156_v0_)})
                        d_1156_v0_ = ((d_1177_v19_)[d_1176_v18_] if (d_1176_v18_) in (d_1177_v19_) else False)
                        index220_ = default__.safeIndex(706, (d_1176_v18_).length(0))
                        (d_1176_v18_)[index220_] = ((self).fm1(d_1175_v17_, d_1158_v2_, globalState)) in (_dafny.Map({False: d_1158_v2_}))
                        d_1178_v20_: _dafny.Set
                        d_1178_v20_ = _dafny.Set({d_1156_v0_})
                        index221_ = default__.safeIndex(706, (d_1176_v18_).length(0))
                        (d_1176_v18_)[index221_] = (d_1178_v20_).isdisjoint(d_1178_v20_)
                        d_1179_v21_: _dafny.Seq
                        d_1179_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1158_v2_]), d_1164_v8_])
                        rhs275_ = d_1156_v0_
                        rhs276_ = d_1179_v21_
                        rhs277_ = d_1158_v2_
                        rhs278_ = (d_1175_v17_) | ((d_1175_v17_) | (d_1175_v17_))
                        d_1156_v0_ = rhs275_
                        d_1179_v21_ = rhs276_
                        d_1158_v2_ = rhs277_
                        d_1175_v17_ = rhs278_
                        d_1180_v22_: D3
                        d_1180_v22_ = D3_DC9(d_1167_v10_)
                        d_1181_v23_: _dafny.Seq
                        d_1181_v23_ = _dafny.SeqWithoutIsStrInference([d_1176_v18_])
                        d_1182_v24_: _dafny.Array
                        nw170_ = _dafny.Array(None, 7)
                        nw170_[int(0)] = d_1176_v18_
                        nw170_[int(1)] = d_1176_v18_
                        nw170_[int(2)] = d_1176_v18_
                        nw170_[int(3)] = (d_1181_v23_)[default__.safeIndex(613, len(d_1181_v23_))]
                        nw170_[int(4)] = d_1176_v18_
                        nw170_[int(5)] = d_1176_v18_
                        nw170_[int(6)] = d_1176_v18_
                        d_1182_v24_ = nw170_
                        index222_ = default__.safeIndex(49, (d_1182_v24_).length(0))
                        (d_1182_v24_)[index222_] = d_1176_v18_
                        d_1183_v25_: _dafny.Map
                        d_1183_v25_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1156_v0_, (d_1176_v18_)[default__.safeIndex(706, (d_1176_v18_).length(0))], (d_1176_v18_)[default__.safeIndex(706, (d_1176_v18_).length(0))]])): d_1160_v4_})
                        d_1184_v26_: _dafny.Set
                        d_1184_v26_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejjgk"))})
                        d_1185_v28_: _dafny.Map
                        def iife84_():
                            coll52_ = _dafny.Map()
                            compr_52_: int
                            for compr_52_ in _dafny.IntegerRange(803, -816):
                                d_1186_v27_: int = compr_52_
                                if ((803) <= (d_1186_v27_)) and ((d_1186_v27_) < (-816)):
                                    coll52_[default__.safeModuloInt(d_1186_v27_, d_1158_v2_)] = d_1156_v0_
                            return _dafny.Map(coll52_)
                        d_1185_v28_ = _dafny.Map({iife84_()
                        : not(not((d_1176_v18_)[default__.safeIndex(706, (d_1176_v18_).length(0))]))})
                        index223_ = default__.safeIndex(706, (d_1176_v18_).length(0))
                        index224_ = default__.safeIndex(49, (d_1182_v24_).length(0))
                        rhs279_ = (((_dafny.MultiSet([d_1156_v0_])).cardinality) - (len(d_1183_v25_))) == ((301) - (d_1158_v2_))
                        rhs280_ = default__.fm14((d_1171_v14_).cardinality, d_1184_v26_, d_1158_v2_, (d_1185_v28_) | ((_dafny.Map({d_1159_v3_: (d_1176_v18_)[default__.safeIndex(706, (d_1176_v18_).length(0))]})).set((_dafny.Map({len(d_1174_v16_): (d_1176_v18_)[default__.safeIndex(706, (d_1176_v18_).length(0))]})).set(len(_dafny.Set({946, d_1158_v2_})), d_1156_v0_), (d_1176_v18_)[default__.safeIndex(706, (d_1176_v18_).length(0))])), globalState)
                        rhs281_ = d_1176_v18_
                        lhs203_ = d_1176_v18_
                        lhs204_ = default__.safeIndex(706, (d_1176_v18_).length(0))
                        lhs205_ = d_1182_v24_
                        lhs206_ = default__.safeIndex(49, (d_1182_v24_).length(0))
                        lhs203_[lhs204_] = rhs279_
                        d_1180_v22_ = rhs280_
                        lhs205_[lhs206_] = rhs281_
                    index225_ = default__.safeIndex(706, (d_1160_v4_).length(0))
                    (d_1160_v4_)[index225_] = len((d_1159_v3_).set(d_1158_v2_, not(d_1156_v0_)))
                    index226_ = default__.safeIndex(706, (d_1160_v4_).length(0))
                    (d_1160_v4_)[index226_] = ((0) - (d_1158_v2_) if d_1156_v0_ else d_1158_v2_)
                    d_1187_v29_: D1
                    d_1187_v29_ = D1_DC2(d_1158_v2_)
                    if ((d_1160_v4_)[default__.safeIndex(706, (d_1160_v4_).length(0))]) > ((d_1187_v29_).cf2):
                        (globalState).f13 = (d_1160_v4_)[default__.safeIndex(706, (d_1160_v4_).length(0))]
                        d_1188_v30_: C0
                        nw171_ = C0()
                        nw171_.ctor__()
                        d_1188_v30_ = nw171_
                        d_1189_v31_: C2
                        nw172_ = C2()
                        nw172_.ctor__()
                        d_1189_v31_ = nw172_
                        d_1190_v32_: _dafny.Seq
                        d_1190_v32_ = _dafny.SeqWithoutIsStrInference([d_1156_v0_])
                        d_1191_v33_: D11
                        d_1191_v33_ = D11_DC30(d_1156_v0_, (d_1160_v4_)[default__.safeIndex(706, (d_1160_v4_).length(0))])
                        d_1192_v34_: _dafny.Map
                        d_1192_v34_ = _dafny.Map({d_1190_v32_: d_1191_v33_})
                        d_1193_v35_: _dafny.Map
                        d_1193_v35_ = _dafny.Map({(_dafny.Map({_dafny.SeqWithoutIsStrInference([False, d_1156_v0_]): D11_DC30(d_1156_v0_, 364)})) | (d_1192_v34_): d_1156_v0_})
                        d_1193_v35_ = (d_1193_v35_).set((d_1192_v34_) | (d_1192_v34_), d_1156_v0_)
                        (globalState).f13 = (d_1160_v4_)[default__.safeIndex(706, (d_1160_v4_).length(0))]
                    elif True:
                        d_1194_v36_: D4
                        d_1194_v36_ = D4_DC12((d_1160_v4_)[default__.safeIndex(706, (d_1160_v4_).length(0))])
                        pat_let_tv18_ = d_1160_v4_
                        pat_let_tv19_ = d_1160_v4_
                        d_1195_v37_: _dafny.Map
                        def iife85_(_pat_let16_0):
                            def iife86_(d_1196_dt__update__tmp_h0_):
                                def iife87_(_pat_let17_0):
                                    def iife88_(d_1197_dt__update_hcf16_h0_):
                                        return D4_DC12(d_1197_dt__update_hcf16_h0_)
                                    return iife88_(_pat_let17_0)
                                return iife87_((pat_let_tv19_)[default__.safeIndex(706, (pat_let_tv18_).length(0))])
                            return iife86_(_pat_let16_0)
                        d_1195_v37_ = _dafny.Map({iife85_(d_1194_v36_): (d_1160_v4_)[default__.safeIndex(706, (d_1160_v4_).length(0))]})
                        rhs282_ = (d_1195_v37_) | (d_1195_v37_)
                        rhs283_ = d_1158_v2_
                        rhs284_ = 660
                        lhs207_ = globalState
                        lhs208_ = globalState
                        d_1195_v37_ = rhs282_
                        lhs207_.f13 = rhs283_
                        lhs208_.f4 = rhs284_
                        d_1198_v38_: _dafny.Set
                        d_1198_v38_ = _dafny.Set({not(d_1156_v0_), not(False)})
                        d_1199_v39_: D12
                        d_1199_v39_ = D12_DC32((d_1198_v38_ if True else d_1198_v38_))
                        d_1199_v39_ = D12_DC32(d_1198_v38_)
                        (globalState).f7 = 641
                        d_1200_v40_: _dafny.Array
                        def lambda60_(d_1201_v0_):
                            def lambda61_(d_1202_i2_):
                                return _dafny.SeqWithoutIsStrInference([d_1201_v0_, d_1201_v0_])

                            return lambda61_

                        init31_ = lambda60_(d_1156_v0_)
                        nw173_ = _dafny.Array(None, 15)
                        for i0_31_ in range(nw173_.length(0)):
                            nw173_[i0_31_] = init31_(i0_31_)
                        d_1200_v40_ = nw173_
                        d_1203_v41_: _dafny.Array
                        d_1203_v41_ = d_1200_v40_
                        d_1204_v42_: _dafny.Map
                        d_1204_v42_ = _dafny.Map({(d_1171_v14_) | (_dafny.MultiSet([len(d_1166_v9_)])): d_1203_v41_})
                        d_1204_v42_ = (d_1204_v42_).set(d_1171_v14_, d_1203_v41_)
                        d_1205_v43_: _dafny.Array
                        nw174_ = _dafny.Array(False, 7)
                        d_1205_v43_ = nw174_
                        d_1206_v44_: _dafny.Seq
                        d_1206_v44_ = _dafny.SeqWithoutIsStrInference([d_1205_v43_, d_1205_v43_, d_1205_v43_, d_1205_v43_])
                        d_1207_v45_: D7
                        d_1207_v45_ = D7_DC19(d_1206_v44_)
                        d_1207_v45_ = d_1207_v45_
                    pass
            pass
        d_1208_v46_: str
        d_1208_v46_ = _dafny.CodePoint('k')
        d_1209_v47_: _dafny.Array
        nw175_ = _dafny.Array(None, 24)
        nw175_[int(0)] = True
        nw175_[int(1)] = (d_1156_v0_) and (True)
        nw175_[int(2)] = d_1156_v0_
        nw175_[int(3)] = d_1156_v0_
        nw175_[int(4)] = (d_1158_v2_) != (d_1158_v2_)
        nw175_[int(5)] = d_1156_v0_
        nw175_[int(6)] = d_1156_v0_
        nw175_[int(7)] = not(False)
        nw175_[int(8)] = d_1156_v0_
        nw175_[int(9)] = d_1156_v0_
        nw175_[int(10)] = (d_1208_v46_) not in (p0)
        nw175_[int(11)] = d_1156_v0_
        nw175_[int(12)] = d_1156_v0_
        nw175_[int(13)] = (d_1156_v0_) == (d_1156_v0_)
        nw175_[int(14)] = d_1156_v0_
        nw175_[int(15)] = d_1156_v0_
        nw175_[int(16)] = d_1156_v0_
        nw175_[int(17)] = d_1156_v0_
        nw175_[int(18)] = ((d_1159_v3_)[-201] if (-201) in (d_1159_v3_) else not(d_1156_v0_))
        nw175_[int(19)] = (55) == (d_1158_v2_)
        nw175_[int(20)] = (d_1164_v8_) < (d_1164_v8_)
        nw175_[int(21)] = default__.fm26(globalState)
        nw175_[int(22)] = (self).fm12(globalState)
        nw175_[int(23)] = d_1156_v0_
        d_1209_v47_ = nw175_
        index227_ = default__.safeIndex(428, (d_1209_v47_).length(0))
        (d_1209_v47_)[index227_] = not(not(False))
        index228_ = default__.safeIndex(428, (d_1209_v47_).length(0))
        (d_1209_v47_)[index228_] = d_1156_v0_
        d_1210_v48_: _dafny.Seq
        d_1210_v48_ = _dafny.SeqWithoutIsStrInference([d_1156_v0_])
        index229_ = default__.safeIndex(428, (d_1209_v47_).length(0))
        (d_1209_v47_)[index229_] = (d_1158_v2_) == (len(d_1210_v48_))


class C8(T0, T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_1211_i0_ in range(default__.abs(326))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.Set({404}))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwtfiavx"))), (_dafny.MultiSet([True])).cardinality]))) < ((_dafny.SeqWithoutIsStrInference([766, 187, 520])) + (_dafny.SeqWithoutIsStrInference([474])))

    def fm2(self, p0, globalState):
        if not((True) or (True)):
            return (0) - (len(((D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmw")), not(False))).cf3) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvx")))))
        elif True:
            return 231

    def fm7(self, p0, globalState):
        def iife89_():
            coll53_ = _dafny.Set()
            compr_53_: int
            for compr_53_ in _dafny.IntegerRange(-228, 300):
                d_1212_v0_: int = compr_53_
                if ((-228) <= (d_1212_v0_)) and ((d_1212_v0_) < (300)):
                    coll53_ = coll53_.union(_dafny.Set([(d_1212_v0_) - ((_dafny.MultiSet([(_dafny.MultiSet([255])).cardinality, 129])).cardinality)]))
            return _dafny.Set(coll53_)
        def iife90_():
            coll54_ = _dafny.Set()
            compr_54_: int
            for compr_54_ in _dafny.IntegerRange(-367, 626):
                d_1213_v1_: int = compr_54_
                if ((-367) <= (d_1213_v1_)) and ((d_1213_v1_) < (626)):
                    coll54_ = coll54_.union(_dafny.Set([(d_1213_v1_) * (122)]))
            return _dafny.Set(coll54_)
        def iife91_():
            coll55_ = _dafny.Set()
            compr_55_: int
            for compr_55_ in (_dafny.Map({9: len(_dafny.Set({len(_dafny.Set({_dafny.CodePoint('w')}))}))})).keys.Elements:
                d_1214_v2_: int = compr_55_
                if (d_1214_v2_) in (_dafny.Map({9: len(_dafny.Set({len(_dafny.Set({_dafny.CodePoint('w')}))}))})):
                    coll55_ = coll55_.union(_dafny.Set([default__.safeModuloInt(d_1214_v2_, 477)]))
            return _dafny.Set(coll55_)
        def iife92_():
            def iife95_():
                coll59_ = _dafny.Map()
                compr_59_: int
                for compr_59_ in _dafny.IntegerRange(723, -208):
                    d_1215_v3_: int = compr_59_
                    if ((723) <= (d_1215_v3_)) and ((d_1215_v3_) < (-208)):
                        def iife96_():
                            coll60_ = _dafny.Map()
                            compr_60_: _dafny.Seq
                            for compr_60_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xypuautvo"))])).Elements:
                                d_1216_v4_: _dafny.Seq = compr_60_
                                if (d_1216_v4_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xypuautvo"))])):
                                    coll60_[d_1216_v4_] = False
                            return _dafny.Map(coll60_)
                        coll59_[(d_1215_v3_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1217_i0_ in range(default__.abs(273))])))] = len(iife96_()
                        )
                return _dafny.Map(coll59_)
            coll56_ = _dafny.Set()
            def iife93_():
                coll57_ = _dafny.Map()
                compr_57_: int
                for compr_57_ in _dafny.IntegerRange(723, -208):
                    d_1215_v3_: int = compr_57_
                    if ((723) <= (d_1215_v3_)) and ((d_1215_v3_) < (-208)):
                        def iife94_():
                            coll58_ = _dafny.Map()
                            compr_58_: _dafny.Seq
                            for compr_58_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xypuautvo"))])).Elements:
                                d_1216_v4_: _dafny.Seq = compr_58_
                                if (d_1216_v4_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xypuautvo"))])):
                                    coll58_[d_1216_v4_] = False
                            return _dafny.Map(coll58_)
                        coll57_[(d_1215_v3_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1217_i0_ in range(default__.abs(273))])))] = len(iife94_()
                        )
                return _dafny.Map(coll57_)
            compr_56_: int
            for compr_56_ in (iife93_()
            ).keys.Elements:
                d_1218_v5_: int = compr_56_
                if (d_1218_v5_) in (iife95_()
                ):
                    coll56_ = coll56_.union(_dafny.Set([(d_1218_v5_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubinwfp"))))]))
            return _dafny.Set(coll56_)
        return (_dafny.MultiSet([iife89_()
        ])) | ((_dafny.MultiSet([iife90_()
        , iife91_()
        ])) - (_dafny.MultiSet([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, True]))}), _dafny.Set({(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(False), True])])).cardinality, 622}), _dafny.Set({18}), iife92_()
        , _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhq")))})])))

    def fm8(self, p0, p1, p2, globalState):
        if (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False])])).issubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])):
            return True
        elif True:
            return (909) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydocfi"))))

    def m1(self, p0, p1, globalState):
        d_1219_v0_: bool
        d_1219_v0_ = True
        d_1219_v0_ = d_1219_v0_
        d_1220_v1_: _dafny.Array
        def lambda62_(d_1221_v0_):
            def lambda63_(d_1222_i1_):
                return d_1221_v0_

            return lambda63_

        init32_ = lambda62_(d_1219_v0_)
        nw176_ = _dafny.Array(None, 6)
        for i0_32_ in range(nw176_.length(0)):
            nw176_[i0_32_] = init32_(i0_32_)
        d_1220_v1_ = nw176_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1220_v1_).length(0)):
            d_1223_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_1223_i0_)) and ((d_1223_i0_) < ((d_1220_v1_).length(0)))):
                (d_1220_v1_)[(d_1223_i0_)] = not (d_1219_v0_) or (not(not(d_1219_v0_)))
        d_1224_v2_: int
        d_1224_v2_ = 551
        d_1225_v4_: _dafny.Map
        d_1225_v4_ = _dafny.Map({d_1224_v2_: not(d_1219_v0_)})
        def iife97_():
            coll61_ = _dafny.Map()
            compr_61_: int
            for compr_61_ in (d_1225_v4_).keys.Elements:
                d_1226_v3_: int = compr_61_
                if (d_1226_v3_) in (d_1225_v4_):
                    coll61_[default__.safeModuloInt(d_1226_v3_, d_1224_v2_)] = d_1224_v2_
            return _dafny.Map(coll61_)
        hi8_ = (d_1224_v2_ if d_1219_v0_ else len(iife97_()
        ))
        for d_1227_i2_ in range(len(default__.fm11(d_1219_v0_, globalState)), hi8_):
            d_1228_v5_: int
            d_1229_v6_: bool
            out9_: int
            out10_: bool
            out9_, out10_ = (self).m8(globalState)
            d_1228_v5_ = out9_
            d_1229_v6_ = out10_
            d_1230_v7_: _dafny.Seq
            d_1230_v7_ = _dafny.SeqWithoutIsStrInference([d_1219_v0_, d_1219_v0_])
            d_1230_v7_ = (d_1230_v7_) + (d_1230_v7_)
            d_1219_v0_ = (d_1219_v0_) and (d_1219_v0_)
            (globalState).f4 = d_1228_v5_
        d_1231_v8_: C5
        nw177_ = C5()
        nw177_.ctor__()
        d_1231_v8_ = nw177_
        d_1220_v1_ = d_1220_v1_
        d_1232_v9_: _dafny.Map
        d_1232_v9_ = _dafny.Map({d_1220_v1_: d_1219_v0_})
        d_1233_i3_: int
        d_1233_i3_ = 0
        with _dafny.label("8"):
            while ((d_1232_v9_)[d_1220_v1_] if (d_1220_v1_) in (d_1232_v9_) else d_1219_v0_):
                with _dafny.c_label("8"):
                    if (d_1233_i3_) >= (100):
                        raise _dafny.Break("8")
                    d_1233_i3_ = (d_1233_i3_) + (1)
                    d_1234_v10_: _dafny.Array
                    nw178_ = _dafny.Array(int(0), 13)
                    d_1234_v10_ = nw178_
                    d_1235_v11_: _dafny.MultiSet
                    d_1235_v11_ = _dafny.MultiSet([d_1234_v10_])
                    d_1236_v12_: _dafny.Map
                    d_1236_v12_ = _dafny.Map({(d_1235_v11_).issubset((d_1235_v11_).set(d_1234_v10_, default__.abs(d_1224_v2_))): (d_1219_v0_) == (d_1219_v0_)})
                    d_1236_v12_ = (d_1236_v12_).set(d_1219_v0_, not(d_1219_v0_))
                    d_1237_v13_: D7
                    d_1237_v13_ = D7_DC21(637, -530, d_1219_v0_, d_1219_v0_, d_1219_v0_)
                    (globalState).f4 = (d_1237_v13_).cf31
                    d_1238_v14_: _dafny.Array
                    def lambda64_(d_1239_v0_):
                        def lambda65_(d_1240_i4_):
                            return _dafny.MultiSet([d_1239_v0_])

                        return lambda65_

                    init33_ = lambda64_(d_1219_v0_)
                    nw179_ = _dafny.Array(None, 28)
                    for i0_33_ in range(nw179_.length(0)):
                        nw179_[i0_33_] = init33_(i0_33_)
                    d_1238_v14_ = nw179_
                    d_1241_v15_: _dafny.Map
                    d_1241_v15_ = _dafny.Map({d_1238_v14_: d_1219_v0_})
                    if ((d_1241_v15_)[d_1238_v14_] if (d_1238_v14_) in (d_1241_v15_) else (d_1219_v0_ if d_1219_v0_ else not(d_1219_v0_))):
                        d_1242_v16_: D6
                        d_1242_v16_ = D6_DC16(_dafny.SeqWithoutIsStrInference([d_1224_v2_, d_1224_v2_, d_1224_v2_]))
                        d_1242_v16_ = d_1242_v16_
                        d_1243_v17_: _dafny.MultiSet
                        d_1243_v17_ = _dafny.MultiSet([(0) - (d_1224_v2_), d_1224_v2_])
                        index230_ = default__.safeIndex(841, (d_1220_v1_).length(0))
                        (d_1220_v1_)[index230_] = (((d_1243_v17_)[d_1224_v2_] if (d_1224_v2_) in (d_1243_v17_) else d_1224_v2_)) > (d_1224_v2_)
                        index231_ = default__.safeIndex(841, (d_1220_v1_).length(0))
                        (d_1220_v1_)[index231_] = ((d_1225_v4_)[default__.safeModuloInt(d_1224_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmlde"))))] if (default__.safeModuloInt(d_1224_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmlde"))))) in (d_1225_v4_) else d_1219_v0_)
                        d_1219_v0_ = (d_1220_v1_)[default__.safeIndex(841, (d_1220_v1_).length(0))]
                        d_1244_v18_: C1
                        nw180_ = C1()
                        nw180_.ctor__()
                        d_1244_v18_ = nw180_
                        d_1245_v19_: D15
                        d_1245_v19_ = D15_DC38(d_1225_v4_)
                        index232_ = default__.safeIndex(794, (d_1234_v10_).length(0))
                        (d_1234_v10_)[index232_] = (0) - (len((d_1245_v19_).cf55))
                        index233_ = default__.safeIndex(841, (d_1220_v1_).length(0))
                        index234_ = default__.safeIndex(794, (d_1234_v10_).length(0))
                        rhs285_ = (d_1220_v1_)[default__.safeIndex(841, (d_1220_v1_).length(0))]
                        rhs286_ = d_1224_v2_
                        rhs287_ = default__.safeDivisionInt(d_1224_v2_, d_1224_v2_)
                        rhs288_ = (d_1220_v1_)[default__.safeIndex(841, (d_1220_v1_).length(0))]
                        lhs209_ = d_1220_v1_
                        lhs210_ = default__.safeIndex(841, (d_1220_v1_).length(0))
                        lhs211_ = globalState
                        lhs212_ = d_1234_v10_
                        lhs213_ = default__.safeIndex(794, (d_1234_v10_).length(0))
                        lhs209_[lhs210_] = rhs285_
                        lhs211_.f7 = rhs286_
                        lhs212_[lhs213_] = rhs287_
                        d_1219_v0_ = rhs288_
                    elif True:
                        d_1219_v0_ = d_1219_v0_
                        d_1246_v20_: C1
                        nw181_ = C1()
                        nw181_.ctor__()
                        d_1246_v20_ = nw181_
                        d_1247_v21_: _dafny.Seq
                        d_1247_v21_ = _dafny.SeqWithoutIsStrInference([d_1246_v20_, d_1246_v20_, d_1246_v20_, d_1246_v20_, d_1246_v20_])
                        d_1246_v20_ = (d_1247_v21_)[default__.safeIndex(d_1224_v2_, len(d_1247_v21_))]
                        d_1224_v2_ = 328
                        d_1248_v22_: _dafny.Seq
                        d_1248_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqpqi"))
                        d_1249_v23_: _dafny.Seq
                        d_1249_v23_ = _dafny.SeqWithoutIsStrInference([(d_1224_v2_) + (d_1224_v2_)])
                        d_1250_v24_: _dafny.Array
                        nw182_ = _dafny.Array(_dafny.Array(None, 0), 12)
                        d_1250_v24_ = nw182_
                        index235_ = default__.safeIndex(347, (d_1250_v24_).length(0))
                        (d_1250_v24_)[index235_] = d_1220_v1_
                        d_1251_v25_: _dafny.Seq
                        d_1251_v25_ = _dafny.SeqWithoutIsStrInference([d_1219_v0_])
                        index236_ = default__.safeIndex(347, (d_1250_v24_).length(0))
                        rhs289_ = (d_1246_v20_).fm25(d_1219_v0_, d_1219_v0_, globalState)
                        rhs290_ = _dafny.SeqWithoutIsStrInference([(D3_DC9(_dafny.CodePoint('d'))).cf12 for d_1252_i5_ in range(default__.abs(947))])
                        rhs291_ = (d_1249_v23_) + (d_1249_v23_)
                        rhs292_ = default__.safeDivisionInt(d_1224_v2_, default__.safeModuloInt(d_1224_v2_, len(d_1251_v25_)))
                        rhs293_ = d_1220_v1_
                        lhs214_ = globalState
                        lhs215_ = d_1250_v24_
                        lhs216_ = default__.safeIndex(347, (d_1250_v24_).length(0))
                        lhs214_.f4 = rhs289_
                        d_1248_v22_ = rhs290_
                        d_1249_v23_ = rhs291_
                        d_1224_v2_ = rhs292_
                        lhs215_[lhs216_] = rhs293_
                        d_1253_v26_: _dafny.Map
                        d_1253_v26_ = _dafny.Map({d_1219_v0_: d_1224_v2_})
                        d_1219_v0_ = (False if (d_1253_v26_) != (_dafny.Map({d_1219_v0_: d_1224_v2_})) else (d_1224_v2_) != ((0) - (d_1224_v2_)))
                    d_1254_v27_: _dafny.MultiSet
                    d_1254_v27_ = _dafny.MultiSet([d_1219_v0_])
                    d_1255_v28_: D2
                    d_1255_v28_ = D2_DC7(d_1219_v0_, d_1254_v27_, d_1234_v10_)
                    d_1256_v29_: D2
                    d_1256_v29_ = D2_DC8(d_1255_v28_)
                    d_1257_v30_: C3
                    nw183_ = C3()
                    nw183_.ctor__(d_1256_v29_, d_1219_v0_)
                    d_1257_v30_ = nw183_
                    pass
            pass

    def m6(self, p0, p1, p2, p3, globalState):
        d_1258_v0_: _dafny.Map
        d_1258_v0_ = _dafny.Map({p2: (self).fm2(p0, globalState)})
        d_1259_v1_: _dafny.Seq
        d_1259_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqid"))
        d_1260_v2_: D4
        d_1260_v2_ = D4_DC11(_dafny.SeqWithoutIsStrInference([d_1259_v1_, d_1259_v1_, d_1259_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1261_i0_ in range(default__.abs(321))]), d_1259_v1_]))
        d_1262_v3_: _dafny.Map
        d_1262_v3_ = _dafny.Map({d_1258_v0_: (d_1260_v2_ if p2 else d_1260_v2_)})
        d_1263_v4_: _dafny.Seq
        d_1263_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1264_i1_ in range(default__.abs(144))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1265_i2_ in range(default__.abs(-780))]), d_1259_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfjnmala"))])
        d_1266_v5_: _dafny.Map
        d_1266_v5_ = _dafny.Map({(0) - (p0): p2})
        d_1267_v6_: _dafny.Seq
        d_1267_v6_ = _dafny.SeqWithoutIsStrInference([(d_1263_v4_)[default__.safeIndex(len(d_1266_v5_), len(d_1263_v4_))], d_1259_v1_])
        d_1262_v3_ = (d_1262_v3_).set(d_1258_v0_, D4_DC11(d_1263_v4_))
        d_1268_i3_: int
        d_1268_i3_ = 0
        with _dafny.label("9"):
            while p2:
                with _dafny.c_label("9"):
                    if (d_1268_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_1268_i3_ = (d_1268_i3_) + (1)
                    d_1269_v7_: _dafny.Array
                    nw184_ = _dafny.Array(False, 1)
                    d_1269_v7_ = nw184_
                    d_1270_v8_: _dafny.Seq
                    d_1270_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: p2})])
                    d_1271_v9_: _dafny.Seq
                    d_1271_v9_ = _dafny.SeqWithoutIsStrInference([(D16_DC41(d_1270_v8_)).cf60])
                    d_1272_v10_: _dafny.Map
                    d_1272_v10_ = _dafny.Map({d_1269_v7_: ((d_1271_v9_)[default__.safeIndex(p0, len(d_1271_v9_))]).set(default__.safeIndex(p0, len((d_1271_v9_)[default__.safeIndex(p0, len(d_1271_v9_))])), _dafny.Map({p2: p2}))})
                    d_1273_v11_: _dafny.Map
                    d_1273_v11_ = _dafny.Map({not(p2): p2})
                    d_1272_v10_ = (d_1272_v10_).set(d_1269_v7_, _dafny.SeqWithoutIsStrInference([_dafny.Map({p2: p2}), d_1273_v11_, d_1273_v11_]))
                    d_1274_v12_: bool
                    d_1274_v12_ = True
                    d_1274_v12_ = (d_1274_v12_) == (d_1274_v12_)
                    d_1275_v13_: _dafny.Seq
                    d_1275_v13_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_1276_v14_: C0
                    nw185_ = C0()
                    nw185_.ctor__()
                    d_1276_v14_ = nw185_
                    d_1277_v15_: _dafny.Set
                    d_1277_v15_ = _dafny.Set({d_1276_v14_})
                    d_1278_v16_: _dafny.Map
                    d_1278_v16_ = _dafny.Map({p2: d_1259_v1_})
                    d_1279_v17_: _dafny.Array
                    nw186_ = _dafny.Array(None, 24)
                    nw186_[int(0)] = default__.safeModuloInt(len(_dafny.Map({p0: (d_1267_v6_)[default__.safeIndex((0) - (p0), len(d_1267_v6_))]})), len(d_1275_v13_))
                    nw186_[int(1)] = 961
                    nw186_[int(2)] = p0
                    nw186_[int(3)] = (-500) * (p0)
                    nw186_[int(4)] = (self).fm2(-428, globalState)
                    nw186_[int(5)] = p0
                    nw186_[int(6)] = (self).fm2(p0, globalState)
                    nw186_[int(7)] = (p0) - (p0)
                    nw186_[int(8)] = len(d_1277_v15_)
                    nw186_[int(9)] = p0
                    nw186_[int(10)] = -819
                    nw186_[int(11)] = p0
                    nw186_[int(12)] = (p1).cardinality
                    nw186_[int(13)] = p0
                    nw186_[int(14)] = len((d_1259_v1_ if d_1274_v12_ else d_1259_v1_))
                    nw186_[int(15)] = len(((d_1278_v16_)[d_1274_v12_] if (d_1274_v12_) in (d_1278_v16_) else d_1259_v1_))
                    nw186_[int(16)] = (0) - (((p1).intersection(p1)).cardinality)
                    nw186_[int(17)] = len((d_1275_v13_) + (default__.fm27(globalState)))
                    nw186_[int(18)] = p0
                    nw186_[int(19)] = (p0) * (p0)
                    nw186_[int(20)] = p0
                    nw186_[int(21)] = (0) - (p0)
                    nw186_[int(22)] = default__.safeModuloInt(p0, p0)
                    nw186_[int(23)] = p0
                    d_1279_v17_ = nw186_
                    d_1280_v18_: D8
                    d_1280_v18_ = D8_DC25(p0)
                    pat_let_tv20_ = p0
                    index237_ = default__.safeIndex(885, (d_1279_v17_).length(0))
                    def iife98_(_pat_let18_0):
                        def iife99_(d_1281_dt__update__tmp_h0_):
                            def iife100_(_pat_let19_0):
                                def iife101_(d_1282_dt__update_hcf42_h0_):
                                    return D8_DC25(d_1282_dt__update_hcf42_h0_)
                                return iife101_(_pat_let19_0)
                            return iife100_(pat_let_tv20_)
                        return iife99_(_pat_let18_0)
                    (d_1279_v17_)[index237_] = (iife98_(d_1280_v18_)).cf42
                    d_1283_v19_: C5
                    nw187_ = C5()
                    nw187_.ctor__()
                    d_1283_v19_ = nw187_
                    index238_ = default__.safeIndex(885, (d_1279_v17_).length(0))
                    rhs294_ = 215
                    rhs295_ = d_1283_v19_
                    lhs217_ = d_1279_v17_
                    lhs218_ = default__.safeIndex(885, (d_1279_v17_).length(0))
                    lhs217_[lhs218_] = rhs294_
                    d_1283_v19_ = rhs295_
                    (globalState).f15 = (d_1279_v17_)[default__.safeIndex(885, (d_1279_v17_).length(0))]
                    pass
            pass
        hi9_ = p0
        for d_1284_i4_ in range(p0, hi9_):
            d_1285_v20_: bool
            d_1285_v20_ = False
            d_1286_v21_: _dafny.Array
            nw188_ = _dafny.Array(None, 15)
            nw188_[int(0)] = p2
            nw188_[int(1)] = p2
            nw188_[int(2)] = p2
            nw188_[int(3)] = d_1285_v20_
            nw188_[int(4)] = d_1285_v20_
            nw188_[int(5)] = not(d_1285_v20_)
            nw188_[int(6)] = default__.fm26(globalState)
            nw188_[int(7)] = d_1285_v20_
            nw188_[int(8)] = d_1285_v20_
            nw188_[int(9)] = d_1285_v20_
            nw188_[int(10)] = d_1285_v20_
            nw188_[int(11)] = d_1285_v20_
            nw188_[int(12)] = not(d_1285_v20_)
            nw188_[int(13)] = d_1285_v20_
            nw188_[int(14)] = False
            d_1286_v21_ = nw188_
            d_1287_v22_: _dafny.Map
            d_1287_v22_ = _dafny.Map({d_1285_v20_: d_1286_v21_})
            d_1288_v23_: _dafny.Set
            d_1288_v23_ = _dafny.Set({p2})
            d_1289_v24_: D12
            d_1289_v24_ = D12_DC32(d_1288_v23_)
            d_1290_v25_: _dafny.Map
            d_1290_v25_ = _dafny.Map({((d_1287_v22_)[d_1285_v20_] if (d_1285_v20_) in (d_1287_v22_) else d_1286_v21_): d_1289_v24_})
            d_1285_v20_ = (d_1290_v25_) != (d_1290_v25_)
            d_1291_v26_: D2
            d_1291_v26_ = D2_DC6(p0)
            d_1292_v27_: C3
            nw189_ = C3()
            nw189_.ctor__(D2_DC8(d_1291_v26_), p2)
            d_1292_v27_ = nw189_
            d_1293_v28_: _dafny.Map
            d_1293_v28_ = _dafny.Map({len(d_1259_v1_): p0})
            d_1294_v29_: _dafny.Map
            d_1294_v29_ = _dafny.Map({len(d_1293_v28_): d_1284_i4_})
            d_1295_v30_: _dafny.Seq
            d_1295_v30_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1284_i4_), default__.fm17(len(d_1293_v28_), globalState)])
            index239_ = default__.safeIndex(464, (d_1286_v21_).length(0))
            (d_1286_v21_)[index239_] = (self).fm8(len(d_1295_v30_), not(d_1292_v27_.f25), d_1284_i4_, globalState)
            index240_ = default__.safeIndex(464, (d_1286_v21_).length(0))
            (d_1286_v21_)[index240_] = not(False)
            d_1296_v31_: _dafny.Array
            def lambda66_(d_1297_v27_):
                def lambda67_(d_1298_i5_):
                    return _dafny.Map({not(d_1297_v27_.f25): 851})

                return lambda67_

            init34_ = lambda66_(d_1292_v27_)
            nw190_ = _dafny.Array(None, 2)
            for i0_34_ in range(nw190_.length(0)):
                nw190_[i0_34_] = init34_(i0_34_)
            d_1296_v31_ = nw190_
            index241_ = default__.safeIndex(984, (d_1296_v31_).length(0))
            (d_1296_v31_)[index241_] = (d_1258_v0_)
            d_1299_v32_: str
            d_1299_v32_ = _dafny.CodePoint('x')
            d_1300_v33_: _dafny.Map
            d_1300_v33_ = _dafny.Map({p0: d_1299_v32_})
            d_1301_v34_: _dafny.Array
            nw191_ = _dafny.Array(None, 14)
            nw191_[int(0)] = _dafny.CodePoint('p')
            nw191_[int(1)] = d_1299_v32_
            nw191_[int(2)] = d_1299_v32_
            nw191_[int(3)] = _dafny.CodePoint('m')
            nw191_[int(4)] = d_1299_v32_
            nw191_[int(5)] = d_1299_v32_
            nw191_[int(6)] = ((d_1300_v33_)[d_1284_i4_] if (d_1284_i4_) in (d_1300_v33_) else d_1299_v32_)
            nw191_[int(7)] = _dafny.CodePoint('x')
            nw191_[int(8)] = _dafny.CodePoint('h')
            nw191_[int(9)] = d_1299_v32_
            nw191_[int(10)] = d_1299_v32_
            nw191_[int(11)] = (d_1299_v32_ if not(p2) else d_1299_v32_)
            nw191_[int(12)] = d_1299_v32_
            nw191_[int(13)] = d_1299_v32_
            d_1301_v34_ = nw191_
            index242_ = default__.safeIndex(938, (d_1301_v34_).length(0))
            (d_1301_v34_)[index242_] = d_1299_v32_
            d_1302_v35_: C5
            nw192_ = C5()
            nw192_.ctor__()
            d_1302_v35_ = nw192_
            d_1303_v36_: _dafny.Map
            d_1303_v36_ = _dafny.Map({p2: d_1284_i4_})
            d_1304_v37_: _dafny.Map
            d_1304_v37_ = _dafny.Map({p2: len(d_1303_v36_)})
            d_1305_v38_: _dafny.Seq
            d_1305_v38_ = _dafny.SeqWithoutIsStrInference([p1, p1, (p1).set(p2, default__.abs(p0))])
            index243_ = default__.safeIndex(984, (d_1296_v31_).length(0))
            index244_ = default__.safeIndex(938, (d_1301_v34_).length(0))
            rhs296_ = (d_1299_v32_) in (d_1259_v1_)
            rhs297_ = (d_1303_v36_).set(not ((d_1286_v21_)[default__.safeIndex(464, (d_1286_v21_).length(0))]) or (True), p0)
            rhs298_ = (d_1259_v1_)[default__.safeIndex(((d_1305_v38_)[default__.safeIndex(len(d_1288_v23_), len(d_1305_v38_))]).cardinality, len(d_1259_v1_))]
            rhs299_ = (d_1286_v21_)[default__.safeIndex(464, (d_1286_v21_).length(0))]
            rhs300_ = d_1302_v35_
            lhs219_ = d_1296_v31_
            lhs220_ = default__.safeIndex(984, (d_1296_v31_).length(0))
            lhs221_ = d_1301_v34_
            lhs222_ = default__.safeIndex(938, (d_1301_v34_).length(0))
            d_1285_v20_ = rhs296_
            lhs219_[lhs220_] = rhs297_
            lhs221_[lhs222_] = rhs298_
            d_1285_v20_ = rhs299_
            d_1302_v35_ = rhs300_
        d_1306_v39_: _dafny.Seq
        d_1306_v39_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1307_v40_: _dafny.Map
        d_1307_v40_ = _dafny.Map({p2: (d_1306_v39_)[default__.safeIndex(-696, len(d_1306_v39_))]})
        d_1308_v41_: _dafny.MultiSet
        d_1308_v41_ = _dafny.MultiSet([default__.fm38(False, globalState)])
        d_1309_v42_: _dafny.Set
        d_1309_v42_ = _dafny.Set({(d_1308_v41_).cardinality})
        d_1307_v40_ = (d_1307_v40_).set(p2, (d_1309_v42_).isdisjoint(d_1309_v42_))
        d_1310_v43_: int
        d_1311_v44_: bool
        out11_: int
        out12_: bool
        out11_, out12_ = (self).m8(globalState)
        d_1310_v43_ = out11_
        d_1311_v44_ = out12_
        d_1312_v45_: str
        d_1312_v45_ = _dafny.CodePoint('k')
        d_1313_v46_: D15
        d_1313_v46_ = D15_DC38(d_1266_v5_)
        pat_let_tv21_ = p0
        pat_let_tv22_ = p0
        def lambda68_(source20_):
            if source20_.is_DC39:
                d_1314___mcc_h0_ = source20_.cf56
                d_1315___mcc_h1_ = source20_.cf57
                d_1316___mcc_h2_ = source20_.cf58
                d_1317_cf58_ = d_1316___mcc_h2_
                d_1318_cf57_ = d_1315___mcc_h1_
                d_1319_cf56_ = d_1314___mcc_h0_
                return pat_let_tv21_
            elif source20_.is_DC38:
                d_1320___mcc_h3_ = source20_.cf55
                d_1321_cf55_ = d_1320___mcc_h3_
                return pat_let_tv22_
            elif True:
                d_1322___mcc_h4_ = source20_.cf59
                d_1323_cf59_ = d_1322___mcc_h4_
                return (626) * (818)

        rhs301_ = (d_1259_v1_) + ((d_1259_v1_).set(default__.safeIndex((_dafny.MultiSet([d_1311_v44_, not(d_1311_v44_)])).cardinality, len(d_1259_v1_)), d_1312_v45_))
        rhs302_ = (0) - (p0)
        rhs303_ = not (not(d_1311_v44_)) or ((p0) == (p0))
        rhs304_ = lambda68_(d_1313_v46_)
        rhs305_ = p0
        lhs223_ = globalState
        lhs224_ = globalState
        d_1259_v1_ = rhs301_
        lhs223_.f15 = rhs302_
        d_1311_v44_ = rhs303_
        lhs224_.f4 = rhs304_
        d_1310_v43_ = rhs305_

    def m8(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1324_v0_: int
        d_1324_v0_ = 919
        (globalState).f15 = d_1324_v0_
        d_1325_v1_: bool
        d_1325_v1_ = False
        d_1326_v2_: _dafny.Map
        d_1326_v2_ = _dafny.Map({d_1325_v1_: d_1324_v0_})
        d_1327_v3_: _dafny.Array
        nw193_ = _dafny.Array(None, 11)
        nw193_[int(0)] = -660
        nw193_[int(1)] = ((0) - (d_1324_v0_)) + (d_1324_v0_)
        nw193_[int(2)] = d_1324_v0_
        nw193_[int(3)] = d_1324_v0_
        nw193_[int(4)] = 541
        nw193_[int(5)] = d_1324_v0_
        nw193_[int(6)] = d_1324_v0_
        nw193_[int(7)] = len(_dafny.Set({d_1326_v2_}))
        nw193_[int(8)] = d_1324_v0_
        nw193_[int(9)] = d_1324_v0_
        nw193_[int(10)] = (d_1324_v0_) + (d_1324_v0_)
        d_1327_v3_ = nw193_
        d_1328_v4_: _dafny.Map
        d_1328_v4_ = _dafny.Map({d_1325_v1_: d_1325_v1_})
        d_1327_v3_ = (d_1327_v3_ if (d_1328_v4_) != (_dafny.Map({False: d_1325_v1_})) else d_1327_v3_)
        d_1329_v5_: D6
        d_1329_v5_ = D6_DC17(d_1325_v1_, d_1325_v1_, d_1325_v1_)
        pat_let_tv23_ = d_1324_v0_
        pat_let_tv24_ = d_1325_v1_
        pat_let_tv25_ = d_1325_v1_
        pat_let_tv26_ = d_1324_v0_
        def lambda69_(source21_):
            if source21_.is_DC17:
                d_1330___mcc_h0_ = source21_.cf24
                d_1331___mcc_h1_ = source21_.cf25
                d_1332___mcc_h2_ = source21_.cf26
                d_1333_cf26_ = d_1332___mcc_h2_
                d_1334_cf25_ = d_1331___mcc_h1_
                d_1335_cf24_ = d_1330___mcc_h0_
                return pat_let_tv23_
            elif source21_.is_DC16:
                d_1336___mcc_h3_ = source21_.cf23
                d_1337_cf23_ = d_1336___mcc_h3_
                return len((_dafny.SeqWithoutIsStrInference([pat_let_tv24_])) + (_dafny.SeqWithoutIsStrInference([pat_let_tv25_])))
            elif True:
                d_1338___mcc_h4_ = source21_.cf27
                d_1339_cf27_ = d_1338___mcc_h4_
                return pat_let_tv26_

        (globalState).f13 = lambda69_(d_1329_v5_)
        rhs306_ = d_1325_v1_
        rhs307_ = default__.fm17(d_1324_v0_, globalState)
        lhs225_ = globalState
        r1 = rhs306_
        lhs225_.f7 = rhs307_
        d_1340_v6_: _dafny.Array
        nw194_ = _dafny.Array(_dafny.CodePoint('D'), 13)
        d_1340_v6_ = nw194_
        d_1341_v7_: str
        d_1341_v7_ = _dafny.CodePoint('k')
        index245_ = default__.safeIndex(972, (d_1340_v6_).length(0))
        (d_1340_v6_)[index245_] = d_1341_v7_
        index246_ = default__.safeIndex(972, (d_1340_v6_).length(0))
        (d_1340_v6_)[index246_] = d_1341_v7_
        d_1342_v8_: _dafny.MultiSet
        d_1342_v8_ = _dafny.MultiSet([d_1325_v1_])
        d_1343_v9_: _dafny.Array
        nw195_ = _dafny.Array(None, 8)
        nw195_[int(0)] = d_1342_v8_
        nw195_[int(1)] = (d_1342_v8_).intersection(_dafny.MultiSet([d_1325_v1_]))
        nw195_[int(2)] = d_1342_v8_
        nw195_[int(3)] = (_dafny.MultiSet([d_1325_v1_])) | (d_1342_v8_)
        nw195_[int(4)] = _dafny.MultiSet([d_1325_v1_])
        nw195_[int(5)] = (_dafny.MultiSet([d_1325_v1_])) | (d_1342_v8_)
        nw195_[int(6)] = d_1342_v8_
        nw195_[int(7)] = (d_1342_v8_) - (d_1342_v8_)
        d_1343_v9_ = nw195_
        d_1344_v10_: _dafny.Seq
        d_1344_v10_ = _dafny.SeqWithoutIsStrInference([d_1325_v1_])
        d_1345_v11_: _dafny.Seq
        d_1345_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1344_v10_), _dafny.MultiSet(d_1344_v10_)])
        d_1346_v12_: _dafny.Set
        d_1346_v12_ = _dafny.Set({d_1325_v1_, d_1325_v1_})
        d_1347_v13_: _dafny.Seq
        d_1347_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apmhaajn"))
        d_1348_v14_: _dafny.Map
        d_1348_v14_ = _dafny.Map({((d_1326_v2_)[d_1325_v1_] if (d_1325_v1_) in (d_1326_v2_) else len(d_1347_v13_)): d_1342_v8_})
        nw196_ = _dafny.Array(None, 12)
        nw196_[int(0)] = _dafny.MultiSet(d_1344_v10_)
        nw196_[int(1)] = d_1342_v8_
        nw196_[int(2)] = d_1342_v8_
        nw196_[int(3)] = d_1342_v8_
        nw196_[int(4)] = ((d_1345_v11_)[default__.safeIndex(d_1324_v0_, len(d_1345_v11_))]) - (d_1342_v8_)
        nw196_[int(5)] = d_1342_v8_
        nw196_[int(6)] = (d_1342_v8_).set(d_1325_v1_, default__.abs(len(d_1346_v12_)))
        nw196_[int(7)] = _dafny.MultiSet(d_1344_v10_)
        nw196_[int(8)] = d_1342_v8_
        nw196_[int(9)] = d_1342_v8_
        nw196_[int(10)] = _dafny.MultiSet(default__.fm27(globalState))
        nw196_[int(11)] = (d_1342_v8_) - (((d_1348_v14_)[d_1324_v0_] if (d_1324_v0_) in (d_1348_v14_) else d_1342_v8_))
        d_1343_v9_ = nw196_
        r0 = d_1324_v0_
        d_1349_v15_: D3
        d_1349_v15_ = D3_DC10(d_1324_v0_, d_1324_v0_)
        pat_let_tv27_ = d_1325_v1_
        pat_let_tv28_ = d_1325_v1_
        pat_let_tv29_ = d_1325_v1_
        def lambda70_(source22_):
            if source22_.is_DC10:
                d_1350___mcc_h5_ = source22_.cf13
                d_1351___mcc_h6_ = source22_.cf14
                d_1352_cf14_ = d_1351___mcc_h6_
                d_1353_cf13_ = d_1350___mcc_h5_
                return not (pat_let_tv27_) or (pat_let_tv28_)
            elif True:
                d_1354___mcc_h7_ = source22_.cf12
                d_1355_cf12_ = d_1354___mcc_h7_
                return pat_let_tv29_

        r1 = lambda70_(d_1349_v15_)
        return r0, r1


class C9(T0, T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, globalState):
        return ((D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1356_i0_ in range(default__.abs(-16))]), False) if not(not(True)) else D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")), True))).cf4

    def fm2(self, p0, globalState):
        return default__.safeModuloInt(len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: False})), 330)

    def fm7(self, p0, globalState):
        def iife102_():
            coll62_ = _dafny.Set()
            compr_62_: int
            for compr_62_ in (_dafny.Set({829})).Elements:
                d_1359_v0_: int = compr_62_
                if (d_1359_v0_) in (_dafny.Set({829})):
                    coll62_ = coll62_.union(_dafny.Set([default__.safeDivisionInt(d_1359_v0_, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality for d_1360_i2_ in range(default__.abs(98))])), len(_dafny.Set({not(False)}))])).cardinality)]))
            return _dafny.Set(coll62_)
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({-569}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(_dafny.Map({False: -381}))): True})), len(_dafny.Set({True})), len(_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1357_i0_ in range(default__.abs(6))]))})): 850})), 49, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1358_i1_ in range(default__.abs(24))]))])])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otolbkm")))})]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet([True])).cardinality})])))) - ((_dafny.MultiSet([_dafny.Set({833}), _dafny.Set({705}), _dafny.Set({909})])) - (_dafny.MultiSet([iife102_()
        ])))

    def fm8(self, p0, p1, p2, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): 390})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): 20}))) != (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqyrsao"))): 481}))

    def fm9(self, p0, globalState):
        source23_ = D2_DC6(-528)
        if source23_.is_DC6:
            d_1361___mcc_h0_ = source23_.cf7
            d_1362_cf7_ = d_1361___mcc_h0_
            def iife103_():
                coll63_ = _dafny.Map()
                compr_63_: int
                for compr_63_ in _dafny.IntegerRange(484, 492):
                    d_1363_v0_: int = compr_63_
                    if ((484) <= (d_1363_v0_)) and ((d_1363_v0_) < (492)):
                        coll63_[default__.safeDivisionInt(d_1363_v0_, d_1362_cf7_)] = not(False)
                return _dafny.Map(coll63_)
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1362_cf7_])), d_1362_cf7_, d_1362_cf7_, 338]))).issubset(_dafny.MultiSet([len(iife103_()
            )]))
        elif source23_.is_DC7:
            d_1364___mcc_h1_ = source23_.cf8
            d_1365___mcc_h2_ = source23_.cf9
            d_1366___mcc_h3_ = source23_.cf10
            d_1367_cf10_ = d_1366___mcc_h3_
            d_1368_cf9_ = d_1365___mcc_h2_
            d_1369_cf8_ = d_1364___mcc_h1_
            return d_1369_cf8_
        elif source23_.is_DC5:
            d_1370___mcc_h4_ = source23_.cf6
            d_1371_cf6_ = d_1370___mcc_h4_
            return ((0) - (-84)) >= (len(_dafny.Map({True: False})))
        elif True:
            d_1372___mcc_h5_ = source23_.cf11
            d_1373_cf11_ = d_1372___mcc_h5_
            return not(not(not(not(True))))

    def fm10(self, p0, p1, p2, globalState):
        return False

    def m1(self, p0, p1, globalState):
        d_1374_v0_: int
        d_1374_v0_ = -428
        d_1375_v1_: _dafny.Map
        d_1375_v1_ = _dafny.Map({d_1374_v0_: d_1374_v0_})
        d_1376_v3_: _dafny.Map
        d_1376_v3_ = _dafny.Map({True: 78})
        d_1377_v4_: _dafny.Seq
        d_1377_v4_ = _dafny.SeqWithoutIsStrInference([len(d_1376_v3_), d_1374_v0_])
        d_1378_v10_: bool
        d_1378_v10_ = False
        d_1379_v11_: _dafny.Array
        nw197_ = _dafny.Array(None, 23)
        nw197_[int(0)] = _dafny.Map({(self).fm2(d_1374_v0_, globalState): d_1374_v0_})
        nw197_[int(1)] = d_1375_v1_
        def iife104_():
            coll64_ = _dafny.Map()
            compr_64_: int
            for compr_64_ in (d_1377_v4_).Elements:
                d_1380_v2_: int = compr_64_
                if (d_1380_v2_) in (d_1377_v4_):
                    coll64_[(d_1380_v2_) + (d_1374_v0_)] = d_1374_v0_
            return _dafny.Map(coll64_)
        nw197_[int(2)] = iife104_()
        
        def iife105_():
            coll65_ = _dafny.Map()
            compr_65_: int
            for compr_65_ in _dafny.IntegerRange(42, 698):
                d_1381_v5_: int = compr_65_
                if ((42) <= (d_1381_v5_)) and ((d_1381_v5_) < (698)):
                    coll65_[default__.safeDivisionInt(d_1381_v5_, d_1374_v0_)] = d_1374_v0_
            return _dafny.Map(coll65_)
        nw197_[int(3)] = iife105_()
        
        nw197_[int(4)] = d_1375_v1_
        nw197_[int(5)] = d_1375_v1_
        nw197_[int(6)] = d_1375_v1_
        nw197_[int(7)] = (d_1375_v1_).set(d_1374_v0_, d_1374_v0_)
        nw197_[int(8)] = _dafny.Map({(d_1377_v4_)[default__.safeIndex(len(p0), len(d_1377_v4_))]: -607})
        def iife106_():
            coll66_ = _dafny.Map()
            compr_66_: int
            for compr_66_ in _dafny.IntegerRange(326, -776):
                d_1382_v6_: int = compr_66_
                if ((326) <= (d_1382_v6_)) and ((d_1382_v6_) < (-776)):
                    coll66_[default__.safeDivisionInt(d_1382_v6_, d_1374_v0_)] = len(d_1376_v3_)
            return _dafny.Map(coll66_)
        nw197_[int(9)] = iife106_()
        
        def iife107_():
            coll67_ = _dafny.Map()
            compr_67_: int
            for compr_67_ in (_dafny.SeqWithoutIsStrInference([d_1374_v0_, d_1374_v0_])).Elements:
                d_1383_v7_: int = compr_67_
                if (d_1383_v7_) in (_dafny.SeqWithoutIsStrInference([d_1374_v0_, d_1374_v0_])):
                    coll67_[default__.safeDivisionInt(d_1383_v7_, d_1374_v0_)] = d_1374_v0_
            return _dafny.Map(coll67_)
        nw197_[int(10)] = (d_1375_v1_) | (iife107_()
        )
        nw197_[int(11)] = (d_1375_v1_) | (d_1375_v1_)
        nw197_[int(12)] = (d_1375_v1_) | (d_1375_v1_)
        def iife108_():
            coll68_ = _dafny.Map()
            compr_68_: int
            for compr_68_ in _dafny.IntegerRange(-666, 400):
                d_1384_v8_: int = compr_68_
                if ((-666) <= (d_1384_v8_)) and ((d_1384_v8_) < (400)):
                    coll68_[default__.safeModuloInt(d_1384_v8_, 365)] = (0) - ((0) - ((_dafny.MultiSet([True, True])).cardinality))
            return _dafny.Map(coll68_)
        nw197_[int(13)] = iife108_()
        
        nw197_[int(14)] = (d_1375_v1_) | (d_1375_v1_)
        nw197_[int(15)] = d_1375_v1_
        def iife109_():
            coll69_ = _dafny.Map()
            compr_69_: int
            for compr_69_ in (_dafny.Map({d_1374_v0_: d_1378_v10_})).keys.Elements:
                d_1385_v9_: int = compr_69_
                if (d_1385_v9_) in (_dafny.Map({d_1374_v0_: d_1378_v10_})):
                    coll69_[(d_1385_v9_) * (215)] = d_1374_v0_
            return _dafny.Map(coll69_)
        nw197_[int(16)] = (iife109_()
        ).set(d_1374_v0_, d_1374_v0_)
        nw197_[int(17)] = (d_1375_v1_) | (d_1375_v1_)
        nw197_[int(18)] = d_1375_v1_
        nw197_[int(19)] = (d_1375_v1_ if d_1378_v10_ else d_1375_v1_)
        nw197_[int(20)] = d_1375_v1_
        nw197_[int(21)] = d_1375_v1_
        nw197_[int(22)] = _dafny.Map({-308: d_1374_v0_})
        d_1379_v11_ = nw197_
        index247_ = default__.safeIndex(808, (d_1379_v11_).length(0))
        def iife110_():
            coll70_ = _dafny.Map()
            compr_70_: int
            for compr_70_ in _dafny.IntegerRange(537, 936):
                d_1386_v12_: int = compr_70_
                if ((537) <= (d_1386_v12_)) and ((d_1386_v12_) < (936)):
                    coll70_[(d_1386_v12_) * (len(d_1377_v4_))] = d_1374_v0_
            return _dafny.Map(coll70_)
        (d_1379_v11_)[index247_] = (iife110_()
        ) | (_dafny.Map({d_1374_v0_: len(p0)}))
        index248_ = default__.safeIndex(808, (d_1379_v11_).length(0))
        (d_1379_v11_)[index248_] = (d_1375_v1_) | ((d_1375_v1_ if not(d_1378_v10_) else d_1375_v1_))
        d_1375_v1_ = (d_1375_v1_).set((d_1377_v4_)[default__.safeIndex(d_1374_v0_, len(d_1377_v4_))], d_1374_v0_)
        d_1387_v13_: _dafny.Set
        d_1387_v13_ = _dafny.Set({d_1374_v0_, d_1374_v0_})
        hi10_ = len(d_1387_v13_)
        for d_1388_i0_ in range(d_1374_v0_, hi10_):
            d_1389_v14_: _dafny.Map
            d_1389_v14_ = _dafny.Map({len(p0): d_1378_v10_})
            if ((d_1389_v14_)[d_1374_v0_] if (d_1374_v0_) in (d_1389_v14_) else d_1378_v10_):
                d_1378_v10_ = d_1378_v10_
                d_1390_v15_: _dafny.MultiSet
                d_1390_v15_ = _dafny.MultiSet([d_1378_v10_])
                d_1391_v16_: T1
                nw198_ = C8()
                nw198_.ctor__()
                d_1391_v16_ = nw198_
                d_1392_v17_: _dafny.Map
                d_1392_v17_ = _dafny.Map({d_1391_v16_: d_1378_v10_})
                d_1393_v18_: _dafny.Map
                d_1393_v18_ = _dafny.Map({d_1378_v10_: d_1392_v17_})
                d_1394_v19_: _dafny.Set
                d_1394_v19_ = _dafny.Set({d_1378_v10_})
                d_1395_v20_: _dafny.Array
                nw199_ = _dafny.Array(None, 10)
                nw199_[int(0)] = len(d_1389_v14_)
                nw199_[int(1)] = d_1388_i0_
                nw199_[int(2)] = len((_dafny.Map({d_1378_v10_: not(d_1378_v10_)})).set(d_1378_v10_, d_1378_v10_))
                nw199_[int(3)] = d_1374_v0_
                nw199_[int(4)] = len(d_1393_v18_)
                nw199_[int(5)] = 829
                nw199_[int(6)] = len(default__.fm32(d_1388_i0_, d_1388_i0_, d_1394_v19_, globalState))
                nw199_[int(7)] = d_1374_v0_
                nw199_[int(8)] = d_1374_v0_
                nw199_[int(9)] = d_1388_i0_
                d_1395_v20_ = nw199_
                d_1396_v21_: D2
                d_1396_v21_ = D2_DC7(False, d_1390_v15_, d_1395_v20_)
                d_1397_v22_: _dafny.Seq
                d_1397_v22_ = _dafny.SeqWithoutIsStrInference([d_1378_v10_, d_1378_v10_, d_1378_v10_, True, d_1378_v10_])
                d_1398_v23_: _dafny.Array
                nw200_ = _dafny.Array(None, 22)
                nw200_[int(0)] = d_1396_v21_
                nw200_[int(1)] = d_1396_v21_
                nw200_[int(2)] = d_1396_v21_
                nw200_[int(3)] = d_1396_v21_
                nw200_[int(4)] = d_1396_v21_
                nw200_[int(5)] = d_1396_v21_
                nw200_[int(6)] = (d_1396_v21_ if d_1378_v10_ else d_1396_v21_)
                nw200_[int(7)] = d_1396_v21_
                nw200_[int(8)] = d_1396_v21_
                nw200_[int(9)] = d_1396_v21_
                nw200_[int(10)] = d_1396_v21_
                nw200_[int(11)] = D2_DC7(d_1378_v10_, _dafny.MultiSet((d_1397_v22_).set(default__.safeIndex(d_1374_v0_, len(d_1397_v22_)), d_1378_v10_)), d_1395_v20_)
                nw200_[int(12)] = d_1396_v21_
                nw200_[int(13)] = d_1396_v21_
                nw200_[int(14)] = d_1396_v21_
                nw200_[int(15)] = D2_DC7(d_1378_v10_, d_1390_v15_, d_1395_v20_)
                nw200_[int(16)] = d_1396_v21_
                nw200_[int(17)] = d_1396_v21_
                nw200_[int(18)] = D2_DC7(True, _dafny.MultiSet([d_1378_v10_]), d_1395_v20_)
                nw200_[int(19)] = D2_DC7(d_1378_v10_, d_1390_v15_, d_1395_v20_)
                nw200_[int(20)] = d_1396_v21_
                nw200_[int(21)] = D2_DC7(d_1378_v10_, default__.fm34(len(default__.fm19(d_1378_v10_, not(d_1378_v10_), globalState)), d_1374_v0_, ((d_1389_v14_)[d_1388_i0_] if (d_1388_i0_) in (d_1389_v14_) else d_1378_v10_), globalState), d_1395_v20_)
                d_1398_v23_ = nw200_
                nw201_ = _dafny.Array(D2.default()(), 25)
                d_1398_v23_ = nw201_
                d_1378_v10_ = (d_1394_v19_).issubset(d_1394_v19_)
                d_1378_v10_ = ((d_1377_v4_) + (d_1377_v4_)) != (_dafny.SeqWithoutIsStrInference([((d_1376_v3_)[d_1378_v10_] if (d_1378_v10_) in (d_1376_v3_) else d_1388_i0_) for d_1399_i1_ in range(default__.abs(305))]))
                d_1400_v24_: _dafny.Map
                d_1400_v24_ = _dafny.Map({_dafny.CodePoint('k'): d_1378_v10_})
                d_1401_v25_: str
                d_1401_v25_ = _dafny.CodePoint('f')
                d_1402_v26_: _dafny.Array
                nw202_ = _dafny.Array(None, 17)
                nw202_[int(0)] = d_1378_v10_
                nw202_[int(1)] = d_1378_v10_
                nw202_[int(2)] = d_1378_v10_
                nw202_[int(3)] = False
                nw202_[int(4)] = d_1378_v10_
                nw202_[int(5)] = d_1378_v10_
                nw202_[int(6)] = d_1378_v10_
                nw202_[int(7)] = d_1378_v10_
                nw202_[int(8)] = d_1378_v10_
                nw202_[int(9)] = d_1378_v10_
                nw202_[int(10)] = d_1378_v10_
                nw202_[int(11)] = d_1378_v10_
                nw202_[int(12)] = d_1378_v10_
                nw202_[int(13)] = not(d_1378_v10_)
                nw202_[int(14)] = d_1378_v10_
                nw202_[int(15)] = ((d_1400_v24_)[d_1401_v25_] if (d_1401_v25_) in (d_1400_v24_) else False)
                nw202_[int(16)] = d_1378_v10_
                d_1402_v26_ = nw202_
                d_1403_v27_: _dafny.Map
                d_1403_v27_ = _dafny.Map({d_1378_v10_: d_1402_v26_})
                d_1403_v27_ = (d_1403_v27_).set(d_1378_v10_, (d_1402_v26_ if True else d_1402_v26_))
            elif True:
                d_1404_v28_: C2
                nw203_ = C2()
                nw203_.ctor__()
                d_1404_v28_ = nw203_
                d_1404_v28_ = d_1404_v28_
                (globalState).f4 = d_1374_v0_
                d_1405_v29_: _dafny.Map
                d_1405_v29_ = _dafny.Map({d_1378_v10_: p0})
                d_1405_v29_ = d_1405_v29_
                d_1406_v30_: C5
                nw204_ = C5()
                nw204_.ctor__()
                d_1406_v30_ = nw204_
                d_1378_v10_ = d_1378_v10_
            d_1407_v31_: str
            d_1407_v31_ = _dafny.CodePoint('h')
            d_1408_v32_: _dafny.Array
            nw205_ = _dafny.Array(False, 13)
            d_1408_v32_ = nw205_
            d_1409_v33_: _dafny.Seq
            d_1409_v33_ = _dafny.SeqWithoutIsStrInference([d_1378_v10_, not (False) or (False), d_1378_v10_])
            d_1410_v34_: D5
            d_1410_v34_ = D5_DC14(d_1407_v31_, d_1378_v10_, d_1374_v0_, d_1378_v10_)
            rhs308_ = (d_1378_v10_) == (d_1378_v10_)
            rhs309_ = not((d_1409_v33_)[default__.safeIndex(d_1388_i0_, len(d_1409_v33_))])
            rhs310_ = (d_1410_v34_).cf18
            rhs311_ = d_1408_v32_
            rhs312_ = (self).fm2(default__.safeDivisionInt(d_1374_v0_, d_1388_i0_), globalState)
            lhs226_ = globalState
            d_1378_v10_ = rhs308_
            d_1378_v10_ = rhs309_
            d_1407_v31_ = rhs310_
            d_1408_v32_ = rhs311_
            lhs226_.f13 = rhs312_
            d_1411_v35_: _dafny.Map
            d_1411_v35_ = _dafny.Map({_dafny.MultiSet([d_1407_v31_]): d_1408_v32_})
            d_1412_v36_: _dafny.MultiSet
            d_1412_v36_ = _dafny.MultiSet([d_1407_v31_, d_1407_v31_, d_1407_v31_, d_1407_v31_])
            d_1411_v35_ = (d_1411_v35_).set(d_1412_v36_, d_1408_v32_)
            d_1413_v37_: _dafny.MultiSet
            d_1413_v37_ = _dafny.MultiSet([d_1374_v0_, len(d_1376_v3_), d_1388_i0_])
            d_1414_v38_: C6
            nw206_ = C6()
            nw206_.ctor__()
            d_1414_v38_ = nw206_
            d_1415_v39_: _dafny.MultiSet
            d_1415_v39_ = _dafny.MultiSet([d_1414_v38_, d_1414_v38_, d_1414_v38_])
            d_1416_v40_: _dafny.Array
            nw207_ = _dafny.Array(None, 17)
            nw207_[int(0)] = len(p0)
            nw207_[int(1)] = 264
            nw207_[int(2)] = d_1388_i0_
            nw207_[int(3)] = (d_1413_v37_).cardinality
            nw207_[int(4)] = (d_1388_i0_) - ((0) - (d_1374_v0_))
            nw207_[int(5)] = 496
            nw207_[int(6)] = d_1388_i0_
            nw207_[int(7)] = (self).fm2(d_1374_v0_, globalState)
            nw207_[int(8)] = d_1388_i0_
            nw207_[int(9)] = d_1374_v0_
            nw207_[int(10)] = d_1374_v0_
            nw207_[int(11)] = d_1374_v0_
            nw207_[int(12)] = d_1374_v0_
            nw207_[int(13)] = (210) + ((d_1413_v37_).cardinality)
            nw207_[int(14)] = d_1374_v0_
            nw207_[int(15)] = ((d_1415_v39_).cardinality if d_1378_v10_ else d_1374_v0_)
            nw207_[int(16)] = default__.safeModuloInt(-570, d_1374_v0_)
            d_1416_v40_ = nw207_
            index249_ = default__.safeIndex(757, (d_1416_v40_).length(0))
            (d_1416_v40_)[index249_] = d_1374_v0_
            index250_ = default__.safeIndex(757, (d_1416_v40_).length(0))
            (d_1416_v40_)[index250_] = d_1374_v0_
        d_1417_v41_: _dafny.Array
        nw208_ = _dafny.Array(False, 11)
        d_1417_v41_ = nw208_
        index251_ = default__.safeIndex(442, (d_1417_v41_).length(0))
        (d_1417_v41_)[index251_] = False
        index252_ = default__.safeIndex(442, (d_1417_v41_).length(0))
        (d_1417_v41_)[index252_] = d_1378_v10_
        d_1418_v42_: D7
        d_1418_v42_ = D7_DC21(d_1374_v0_, d_1374_v0_, d_1378_v10_, (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))], d_1378_v10_)
        if not(((d_1418_v42_).cf32) == ((d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))])):
            d_1419_v43_: _dafny.Array
            def lambda71_(d_1420_v41_):
                def lambda72_(d_1421_i2_):
                    return _dafny.MultiSet([not(not(True)), (d_1420_v41_)[default__.safeIndex(442, (d_1420_v41_).length(0))]])

                return lambda72_

            init35_ = lambda71_(d_1417_v41_)
            nw209_ = _dafny.Array(None, 7)
            for i0_35_ in range(nw209_.length(0)):
                nw209_[i0_35_] = init35_(i0_35_)
            d_1419_v43_ = nw209_
            index253_ = default__.safeIndex(320, (d_1419_v43_).length(0))
            (d_1419_v43_)[index253_] = _dafny.MultiSet([True])
            d_1422_v44_: _dafny.Set
            d_1422_v44_ = _dafny.Set({(d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))]})
            d_1423_v45_: _dafny.Map
            d_1423_v45_ = _dafny.Map({d_1422_v44_: (self).fm9(d_1374_v0_, globalState)})
            d_1424_v46_: _dafny.MultiSet
            d_1424_v46_ = _dafny.MultiSet([((d_1423_v45_)[d_1422_v44_] if (d_1422_v44_) in (d_1423_v45_) else (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))])])
            index254_ = default__.safeIndex(320, (d_1419_v43_).length(0))
            (d_1419_v43_)[index254_] = d_1424_v46_
            d_1425_v47_: _dafny.Array
            nw210_ = _dafny.Array(int(0), 19)
            d_1425_v47_ = nw210_
            index255_ = default__.safeIndex(117, (d_1425_v47_).length(0))
            (d_1425_v47_)[index255_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1426_i3_ in range(default__.abs(-664))]))
            index256_ = default__.safeIndex(117, (d_1425_v47_).length(0))
            (d_1425_v47_)[index256_] = d_1374_v0_
            d_1425_v47_ = d_1425_v47_
            d_1427_v48_: _dafny.Seq
            d_1427_v48_ = _dafny.SeqWithoutIsStrInference([d_1378_v10_])
            index257_ = default__.safeIndex(442, (d_1417_v41_).length(0))
            rhs313_ = True
            rhs314_ = d_1427_v48_
            rhs315_ = ((d_1422_v44_) - (d_1422_v44_)) - (d_1422_v44_)
            rhs316_ = (d_1425_v47_)[default__.safeIndex(117, (d_1425_v47_).length(0))]
            lhs227_ = d_1417_v41_
            lhs228_ = default__.safeIndex(442, (d_1417_v41_).length(0))
            lhs229_ = globalState
            lhs227_[lhs228_] = rhs313_
            d_1427_v48_ = rhs314_
            d_1422_v44_ = rhs315_
            lhs229_.f13 = rhs316_
            d_1428_v49_: str
            d_1428_v49_ = _dafny.CodePoint('h')
            d_1429_v50_: _dafny.Seq
            d_1429_v50_ = _dafny.SeqWithoutIsStrInference([((p0).set(default__.safeIndex(d_1374_v0_, len(p0)), d_1428_v49_)) + (p0), default__.fm19(d_1378_v10_, d_1378_v10_, globalState), (p0) + (p0), (p0).set(default__.safeIndex((d_1425_v47_)[default__.safeIndex(117, (d_1425_v47_).length(0))], len(p0)), d_1428_v49_)])
            d_1429_v50_ = d_1429_v50_
        elif True:
            (globalState).f7 = d_1374_v0_
            index258_ = default__.safeIndex(442, (d_1417_v41_).length(0))
            (d_1417_v41_)[index258_] = ((d_1374_v0_) > (d_1374_v0_) if d_1378_v10_ else d_1378_v10_)
            d_1430_v51_: C8
            nw211_ = C8()
            nw211_.ctor__()
            d_1430_v51_ = nw211_
            d_1431_v52_: _dafny.Array
            def lambda73_(d_1432_v0_):
                def lambda74_(d_1433_i4_):
                    return (d_1433_i4_) - (d_1432_v0_)

                return lambda74_

            init36_ = lambda73_(d_1374_v0_)
            nw212_ = _dafny.Array(None, 28)
            for i0_36_ in range(nw212_.length(0)):
                nw212_[i0_36_] = init36_(i0_36_)
            d_1431_v52_ = nw212_
            index259_ = default__.safeIndex(790, (d_1431_v52_).length(0))
            (d_1431_v52_)[index259_] = d_1374_v0_
            index260_ = default__.safeIndex(790, (d_1431_v52_).length(0))
            (d_1431_v52_)[index260_] = (0) - ((0) - ((default__.fm17((self).fm2(d_1374_v0_, globalState), globalState)) * (default__.safeModuloInt(d_1374_v0_, d_1374_v0_))))
            if ((d_1431_v52_)[default__.safeIndex(790, (d_1431_v52_).length(0))]) < (default__.safeModuloInt((d_1431_v52_)[default__.safeIndex(790, (d_1431_v52_).length(0))], len(p0))):
                d_1434_v53_: D6
                d_1434_v53_ = D6_DC17((d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))], (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))], (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))])
                d_1378_v10_ = (self).fm8((d_1431_v52_)[default__.safeIndex(790, (d_1431_v52_).length(0))], ((d_1434_v53_).cf26 if d_1378_v10_ else d_1378_v10_), (0) - (d_1374_v0_), globalState)
                d_1435_v54_: C2
                nw213_ = C2()
                nw213_.ctor__()
                d_1435_v54_ = nw213_
                d_1436_v55_: _dafny.MultiSet
                d_1436_v55_ = _dafny.MultiSet([(d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))]])
                d_1437_v56_: _dafny.Map
                d_1437_v56_ = _dafny.Map({(d_1431_v52_)[default__.safeIndex(790, (d_1431_v52_).length(0))]: d_1378_v10_})
                d_1438_v57_: D15
                d_1438_v57_ = D15_DC38(d_1437_v56_)
                d_1439_v58_: _dafny.Map
                d_1439_v58_ = _dafny.Map({(d_1436_v55_) - (d_1436_v55_): d_1438_v57_})
                d_1440_v59_: str
                d_1440_v59_ = _dafny.CodePoint('u')
                d_1441_v60_: _dafny.Map
                d_1441_v60_ = _dafny.Map({d_1378_v10_: d_1440_v59_})
                d_1439_v58_ = (d_1439_v58_).set(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).fm10(d_1441_v60_, d_1374_v0_, (d_1379_v11_)[default__.safeIndex(808, (d_1379_v11_).length(0))], globalState), d_1378_v10_, (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))], (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))]])), d_1438_v57_)
                (globalState).f4 = d_1374_v0_
                (globalState).f4 = d_1374_v0_
            elif True:
                d_1377_v4_ = d_1377_v4_
                d_1442_v61_: C4
                nw214_ = C4()
                nw214_.ctor__()
                d_1442_v61_ = nw214_
                (globalState).f4 = (d_1431_v52_)[default__.safeIndex(790, (d_1431_v52_).length(0))]
                d_1376_v3_ = (d_1376_v3_).set(d_1378_v10_, (d_1431_v52_)[default__.safeIndex(790, (d_1431_v52_).length(0))])
                d_1443_v62_: _dafny.MultiSet
                d_1443_v62_ = _dafny.MultiSet([not((d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))])])
                d_1444_v63_: D2
                d_1444_v63_ = D2_DC5(d_1431_v52_)
                (d_1430_v51_).m6(-905, (d_1443_v62_) | (d_1443_v62_), True, d_1444_v63_, globalState)
        d_1445_v64_: _dafny.Array
        nw215_ = _dafny.Array(int(0), 24)
        d_1445_v64_ = nw215_
        d_1446_v65_: D2
        d_1446_v65_ = D2_DC7((d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))], _dafny.MultiSet([(d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))]]), d_1445_v64_)
        d_1447_v66_: D2
        d_1447_v66_ = D2_DC8(d_1446_v65_)
        d_1448_v67_: _dafny.Map
        d_1448_v67_ = _dafny.Map({d_1374_v0_: d_1378_v10_})
        d_1449_v68_: C3
        nw216_ = C3()
        nw216_.ctor__(d_1447_v66_, ((d_1448_v67_)[d_1374_v0_] if (d_1374_v0_) in (d_1448_v67_) else (d_1417_v41_)[default__.safeIndex(442, (d_1417_v41_).length(0))]))
        d_1449_v68_ = nw216_

    def m6(self, p0, p1, p2, p3, globalState):
        d_1450_v0_: _dafny.Set
        d_1450_v0_ = _dafny.Set({p0})
        d_1451_v1_: _dafny.Seq
        d_1451_v1_ = _dafny.SeqWithoutIsStrInference([d_1450_v0_])
        d_1452_v2_: _dafny.Seq
        d_1452_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwnxp"))
        d_1453_v3_: D5
        d_1453_v3_ = D5_DC13(_dafny.MultiSet([(d_1451_v1_)[default__.safeIndex(len(d_1452_v2_), len(d_1451_v1_))]]))
        source24_ = d_1453_v3_
        if source24_.is_DC14:
            d_1454___mcc_h0_ = source24_.cf18
            d_1455___mcc_h1_ = source24_.cf19
            d_1456___mcc_h2_ = source24_.cf20
            d_1457___mcc_h3_ = source24_.cf21
            d_1458_cf21_ = d_1457___mcc_h3_
            d_1459_cf20_ = d_1456___mcc_h2_
            d_1460_cf19_ = d_1455___mcc_h1_
            d_1461_cf18_ = d_1454___mcc_h0_
            d_1462_v4_: _dafny.MultiSet
            d_1462_v4_ = _dafny.MultiSet([435])
            d_1463_v5_: _dafny.Map
            d_1463_v5_ = _dafny.Map({p0: d_1462_v4_})
            d_1464_v6_: _dafny.Seq
            d_1464_v6_ = _dafny.SeqWithoutIsStrInference([D1_DC4(((d_1463_v5_)[p0] if (p0) in (d_1463_v5_) else d_1462_v4_))])
            d_1464_v6_ = _dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.MultiSet([len(_dafny.Set({p0, d_1459_cf20_, d_1459_cf20_, -24, len(_dafny.Set({d_1460_cf19_, d_1458_cf21_, not(d_1458_cf21_)}))})), p0, p0])) for d_1465_i0_ in range(default__.abs(-20))])
            d_1466_v7_: D11
            d_1466_v7_ = D11_DC30(p2, default__.fm17((0) - (p0), globalState))
            (globalState).f7 = (d_1466_v7_).cf47
            d_1467_v8_: _dafny.Array
            nw217_ = _dafny.Array(False, 5)
            d_1467_v8_ = nw217_
            d_1468_v9_: _dafny.Seq
            d_1468_v9_ = _dafny.SeqWithoutIsStrInference([d_1467_v8_, d_1467_v8_, d_1467_v8_])
            d_1469_v10_: _dafny.Map
            d_1469_v10_ = _dafny.Map({False: d_1467_v8_})
            d_1468_v9_ = _dafny.SeqWithoutIsStrInference([d_1467_v8_, d_1467_v8_, d_1467_v8_, ((d_1469_v10_)[d_1460_cf19_] if (d_1460_cf19_) in (d_1469_v10_) else d_1467_v8_), d_1467_v8_])
            if not(not (d_1460_cf19_) or ((p1).ispropersubset(p1))):
                d_1467_v8_ = d_1467_v8_
                d_1470_v12_: D11
                d_1470_v12_ = D11_DC29()
                d_1471_v13_: _dafny.Set
                d_1471_v13_ = _dafny.Set({d_1470_v12_, d_1470_v12_})
                d_1472_v14_: _dafny.Map
                d_1472_v14_ = _dafny.Map({p0: d_1460_cf19_})
                d_1473_v15_: D12
                d_1473_v15_ = D12_DC32(_dafny.Set({p2}))
                d_1474_v16_: _dafny.Seq
                d_1474_v16_ = _dafny.SeqWithoutIsStrInference([d_1473_v15_, d_1473_v15_])
                d_1475_v17_: _dafny.Map
                d_1475_v17_ = _dafny.Map({d_1471_v13_: (d_1472_v14_).set(len(d_1474_v16_), d_1458_cf21_)})
                def iife111_():
                    coll71_ = _dafny.Map()
                    compr_71_: _dafny.Set
                    for compr_71_ in (d_1475_v17_).keys.Elements:
                        d_1476_v11_: _dafny.Set = compr_71_
                        if (d_1476_v11_) in (d_1475_v17_):
                            coll71_[d_1476_v11_] = False
                    return _dafny.Map(coll71_)
                (globalState).f13 = len(iife111_()
                )
                d_1477_v18_: _dafny.Map
                d_1477_v18_ = _dafny.Map({p2: d_1460_cf19_})
                d_1478_v19_: _dafny.Map
                d_1478_v19_ = _dafny.Map({p2: d_1477_v18_})
                d_1479_v20_: _dafny.Seq
                d_1479_v20_ = _dafny.SeqWithoutIsStrInference([p2, d_1458_cf21_, d_1458_cf21_, (d_1458_cf21_ if d_1458_cf21_ else (self).fm1(d_1478_v19_, len(d_1452_v2_), globalState))])
                d_1480_v21_: D16
                d_1480_v21_ = D16_DC42(d_1460_cf19_, d_1458_cf21_, d_1460_cf19_)
                d_1481_v22_: _dafny.Array
                nw218_ = _dafny.Array(None, 12)
                nw218_[int(0)] = d_1480_v21_
                nw218_[int(1)] = d_1480_v21_
                nw218_[int(2)] = d_1480_v21_
                nw218_[int(3)] = d_1480_v21_
                nw218_[int(4)] = D16_DC42(p2, (d_1480_v21_).cf61, d_1460_cf19_)
                nw218_[int(5)] = d_1480_v21_
                nw218_[int(6)] = default__.fm46(p2, globalState)
                nw218_[int(7)] = d_1480_v21_
                nw218_[int(8)] = d_1480_v21_
                nw218_[int(9)] = d_1480_v21_
                nw218_[int(10)] = d_1480_v21_
                nw218_[int(11)] = d_1480_v21_
                d_1481_v22_ = nw218_
                index261_ = default__.safeIndex(482, (d_1481_v22_).length(0))
                (d_1481_v22_)[index261_] = d_1480_v21_
                index262_ = default__.safeIndex(482, (d_1481_v22_).length(0))
                rhs317_ = (d_1479_v20_) + ((d_1479_v20_).set(default__.safeIndex(p0, len(d_1479_v20_)), not(d_1458_cf21_)))
                rhs318_ = p0
                rhs319_ = d_1480_v21_
                lhs230_ = globalState
                lhs231_ = d_1481_v22_
                lhs232_ = default__.safeIndex(482, (d_1481_v22_).length(0))
                d_1479_v20_ = rhs317_
                lhs230_.f13 = rhs318_
                lhs231_[lhs232_] = rhs319_
                d_1482_v23_: int
                d_1483_v24_: _dafny.MultiSet
                out13_: int
                out14_: _dafny.MultiSet
                out13_, out14_ = (self).m7(default__.safeModuloInt(d_1459_cf20_, 821), d_1460_cf19_, globalState)
                d_1482_v23_ = out13_
                d_1483_v24_ = out14_
                index263_ = default__.safeIndex(900, (d_1467_v8_).length(0))
                (d_1467_v8_)[index263_] = (self).fm8(d_1459_cf20_, d_1458_cf21_, p0, globalState)
                index264_ = default__.safeIndex(900, (d_1467_v8_).length(0))
                (d_1467_v8_)[index264_] = not(False)
            elif True:
                index265_ = default__.safeIndex(818, (d_1467_v8_).length(0))
                (d_1467_v8_)[index265_] = d_1458_cf21_
                d_1484_v25_: _dafny.Seq
                d_1484_v25_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1459_cf20_)])
                d_1485_v26_: _dafny.MultiSet
                d_1485_v26_ = _dafny.MultiSet([d_1452_v2_, (d_1452_v2_).set(default__.safeIndex(len(d_1484_v25_), len(d_1452_v2_)), d_1461_cf18_), d_1452_v2_])
                index266_ = default__.safeIndex(818, (d_1467_v8_).length(0))
                (d_1467_v8_)[index266_] = (default__.fm47(p2, (0) - (d_1459_cf20_), globalState)).ispropersubset(d_1485_v26_)
                d_1452_v2_ = (d_1452_v2_) + ((d_1452_v2_) + (d_1452_v2_))
                d_1486_v27_: C1
                nw219_ = C1()
                nw219_.ctor__()
                d_1486_v27_ = nw219_
                d_1487_v28_: _dafny.Seq
                d_1487_v28_ = _dafny.SeqWithoutIsStrInference([d_1486_v27_, d_1486_v27_])
                d_1488_v29_: _dafny.Map
                d_1488_v29_ = _dafny.Map({133: d_1487_v28_})
                d_1489_v30_: _dafny.Map
                d_1489_v30_ = _dafny.Map({771: d_1488_v29_})
                d_1490_v31_: _dafny.Map
                d_1490_v31_ = _dafny.Map({not(d_1460_cf19_): d_1459_cf20_})
                d_1491_v32_: _dafny.Array
                nw220_ = _dafny.Array(None, 23)
                nw220_[int(0)] = (d_1488_v29_) | ((d_1488_v29_).set(d_1459_cf20_, d_1487_v28_))
                nw220_[int(1)] = d_1488_v29_
                nw220_[int(2)] = d_1488_v29_
                nw220_[int(3)] = d_1488_v29_
                nw220_[int(4)] = (d_1488_v29_).set(d_1459_cf20_, d_1487_v28_)
                nw220_[int(5)] = d_1488_v29_
                nw220_[int(6)] = (_dafny.Map({len(d_1452_v2_): d_1487_v28_}) if d_1458_cf21_ else d_1488_v29_)
                nw220_[int(7)] = d_1488_v29_
                nw220_[int(8)] = d_1488_v29_
                nw220_[int(9)] = (d_1488_v29_) | (d_1488_v29_)
                nw220_[int(10)] = (d_1488_v29_) | (((d_1489_v30_)[(d_1484_v25_)[default__.safeIndex(d_1459_cf20_, len(d_1484_v25_))]] if ((d_1484_v25_)[default__.safeIndex(d_1459_cf20_, len(d_1484_v25_))]) in (d_1489_v30_) else _dafny.Map({d_1459_cf20_: _dafny.SeqWithoutIsStrInference([d_1486_v27_])})))
                nw220_[int(11)] = d_1488_v29_
                nw220_[int(12)] = (d_1488_v29_) | (_dafny.Map({p0: d_1487_v28_}))
                nw220_[int(13)] = (d_1488_v29_) | (d_1488_v29_)
                nw220_[int(14)] = d_1488_v29_
                nw220_[int(15)] = (_dafny.Map({p0: d_1487_v28_})).set((0) - (d_1459_cf20_), d_1487_v28_)
                nw220_[int(16)] = d_1488_v29_
                nw220_[int(17)] = d_1488_v29_
                nw220_[int(18)] = d_1488_v29_
                nw220_[int(19)] = (_dafny.Map({p0: d_1487_v28_})).set(len((d_1490_v31_).set(d_1460_cf19_, p0)), d_1487_v28_)
                nw220_[int(20)] = d_1488_v29_
                nw220_[int(21)] = (d_1488_v29_) | (d_1488_v29_)
                nw220_[int(22)] = (d_1488_v29_) | (d_1488_v29_)
                d_1491_v32_ = nw220_
                index267_ = default__.safeIndex(319, (d_1491_v32_).length(0))
                (d_1491_v32_)[index267_] = d_1488_v29_
                index268_ = default__.safeIndex(319, (d_1491_v32_).length(0))
                (d_1491_v32_)[index268_] = d_1488_v29_
                (globalState).f4 = p0
                d_1492_v33_: _dafny.Map
                d_1492_v33_ = _dafny.Map({d_1458_cf21_: d_1460_cf19_})
                d_1493_v34_: _dafny.Map
                d_1493_v34_ = _dafny.Map({len(d_1492_v33_): d_1459_cf20_})
                d_1493_v34_ = (d_1493_v34_).set((d_1459_cf20_) + ((0) - (p0)), (241) - (d_1459_cf20_))
        elif source24_.is_DC13:
            d_1494___mcc_h4_ = source24_.cf17
            d_1495_cf17_ = d_1494___mcc_h4_
            d_1496_v35_: bool
            d_1496_v35_ = False
            d_1496_v35_ = p2
            d_1497_v36_: _dafny.Seq
            d_1497_v36_ = _dafny.SeqWithoutIsStrInference([d_1452_v2_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1498_i1_ in range(default__.abs(279))]) if p2 else d_1452_v2_), (d_1452_v2_) + (d_1452_v2_), d_1452_v2_, d_1452_v2_])
            d_1499_v37_: T0
            nw221_ = C2()
            nw221_.ctor__()
            d_1499_v37_ = nw221_
            d_1500_v38_: _dafny.Set
            d_1500_v38_ = _dafny.Set({d_1499_v37_, d_1499_v37_, d_1499_v37_, d_1499_v37_, d_1499_v37_})
            d_1501_v39_: _dafny.Map
            d_1501_v39_ = _dafny.Map({len(d_1500_v38_): p0})
            d_1502_v40_: _dafny.Seq
            d_1502_v40_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_1501_v39_), 269])
            d_1503_v41_: _dafny.Seq
            d_1503_v41_ = _dafny.SeqWithoutIsStrInference([len(d_1450_v0_), (0) - ((d_1502_v40_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "is"))), len(d_1502_v40_))]), (0) - (p0)])
            d_1504_v42_: _dafny.Seq
            d_1504_v42_ = _dafny.SeqWithoutIsStrInference([d_1502_v40_])
            d_1452_v2_ = (d_1497_v36_)[default__.safeIndex(len(d_1504_v42_), len(d_1497_v36_))]
            d_1505_v43_: _dafny.Array
            def lambda75_(d_1506_v35_):
                def lambda76_(d_1507_i2_):
                    return d_1506_v35_

                return lambda76_

            init37_ = lambda75_(d_1496_v35_)
            nw222_ = _dafny.Array(None, 2)
            for i0_37_ in range(nw222_.length(0)):
                nw222_[i0_37_] = init37_(i0_37_)
            d_1505_v43_ = nw222_
            index269_ = default__.safeIndex(747, (d_1505_v43_).length(0))
            (d_1505_v43_)[index269_] = p2
            index270_ = default__.safeIndex(747, (d_1505_v43_).length(0))
            (d_1505_v43_)[index270_] = not(p2)
            d_1508_v44_: _dafny.Array
            nw223_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1508_v44_ = nw223_
            d_1509_v45_: _dafny.Array
            def lambda77_(d_1510_i3_):
                return (d_1510_i3_) + (-411)

            init38_ = lambda77_
            nw224_ = _dafny.Array(None, 9)
            for i0_38_ in range(nw224_.length(0)):
                nw224_[i0_38_] = init38_(i0_38_)
            d_1509_v45_ = nw224_
            index271_ = default__.safeIndex(169, (d_1508_v44_).length(0))
            (d_1508_v44_)[index271_] = d_1509_v45_
            index272_ = default__.safeIndex(169, (d_1508_v44_).length(0))
            rhs320_ = d_1509_v45_
            rhs321_ = (d_1505_v43_)[default__.safeIndex(747, (d_1505_v43_).length(0))]
            rhs322_ = d_1501_v39_
            rhs323_ = default__.safeModuloInt(len(d_1452_v2_), p0)
            lhs233_ = d_1508_v44_
            lhs234_ = default__.safeIndex(169, (d_1508_v44_).length(0))
            lhs235_ = globalState
            lhs233_[lhs234_] = rhs320_
            d_1496_v35_ = rhs321_
            d_1501_v39_ = rhs322_
            lhs235_.f7 = rhs323_
        elif True:
            d_1511___mcc_h5_ = source24_.cf22
            d_1512_cf22_ = d_1511___mcc_h5_
            if (p0) > (((p1)[True] if (True) in (p1) else p0)):
                d_1513_v46_: _dafny.MultiSet
                d_1513_v46_ = _dafny.MultiSet([793, p0, p0])
                d_1514_v47_: _dafny.Seq
                d_1514_v47_ = _dafny.SeqWithoutIsStrInference([(p0) in (d_1513_v46_), p2, (self).fm9(p0, globalState)])
                d_1514_v47_ = d_1514_v47_
                (globalState).f13 = p0
                d_1515_v48_: C6
                nw225_ = C6()
                nw225_.ctor__()
                d_1515_v48_ = nw225_
                d_1516_v49_: _dafny.Seq
                d_1516_v49_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((0) - (p0), p0), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkgfi"))) + (d_1452_v2_)), p0])
                d_1516_v49_ = _dafny.SeqWithoutIsStrInference([p0 for d_1517_i4_ in range(default__.abs(237))])
                d_1518_v50_: _dafny.Map
                d_1518_v50_ = _dafny.Map({(0) - (p0): p2})
                d_1519_v51_: _dafny.Map
                d_1519_v51_ = _dafny.Map({898: ((d_1518_v50_)[((d_1513_v46_)[p0] if (p0) in (d_1513_v46_) else p0)] if (((d_1513_v46_)[p0] if (p0) in (d_1513_v46_) else p0)) in (d_1518_v50_) else (d_1514_v47_)[default__.safeIndex(p0, len(d_1514_v47_))])})
                d_1520_v52_: _dafny.Seq
                d_1520_v52_ = _dafny.SeqWithoutIsStrInference([d_1519_v51_])
                d_1521_v53_: D14
                d_1521_v53_ = D14_DC35(d_1520_v52_)
                d_1522_v54_: _dafny.Seq
                d_1522_v54_ = _dafny.SeqWithoutIsStrInference([d_1521_v53_])
                d_1523_v55_: D1
                d_1523_v55_ = D1_DC3(d_1452_v2_, p2)
                d_1522_v54_ = (d_1522_v54_) + (_dafny.SeqWithoutIsStrInference([D14_DC35(d_1520_v52_), default__.fm48(not(p2), p2, (d_1523_v55_).cf4, globalState)]))
            elif True:
                d_1524_v56_: bool
                d_1524_v56_ = False
                d_1525_v57_: _dafny.Map
                d_1525_v57_ = _dafny.Map({p0: d_1450_v0_})
                d_1524_v56_ = (d_1450_v0_) == (((d_1525_v57_)[p0] if (p0) in (d_1525_v57_) else d_1450_v0_))
                d_1524_v56_ = False
                (globalState).f7 = (len(d_1450_v0_)) * (p0)
                d_1526_v58_: C5
                nw226_ = C5()
                nw226_.ctor__()
                d_1526_v58_ = nw226_
                d_1527_v59_: _dafny.Seq
                d_1527_v59_ = _dafny.SeqWithoutIsStrInference([d_1524_v56_, p2])
                d_1528_v60_: _dafny.Seq
                d_1528_v60_ = _dafny.SeqWithoutIsStrInference([p2, (d_1527_v59_)[default__.safeIndex(len(d_1527_v59_), len(d_1527_v59_))]])
                d_1529_v61_: _dafny.Set
                d_1529_v61_ = _dafny.Set({p2})
                d_1530_v62_: D2
                d_1530_v62_ = D2_DC6(default__.safeDivisionInt(p0, (0) - (p0)))
                rhs324_ = d_1452_v2_
                rhs325_ = (d_1527_v59_) + ((_dafny.SeqWithoutIsStrInference([d_1524_v56_, p2])).set(default__.safeIndex(584, len(_dafny.SeqWithoutIsStrInference([d_1524_v56_, p2]))), p2))
                rhs326_ = d_1529_v61_
                rhs327_ = d_1530_v62_
                d_1452_v2_ = rhs324_
                d_1527_v59_ = rhs325_
                d_1529_v61_ = rhs326_
                d_1530_v62_ = rhs327_
            d_1531_v63_: _dafny.Map
            d_1531_v63_ = _dafny.Map({p0: 246})
            d_1532_v64_: _dafny.Map
            d_1532_v64_ = _dafny.Map({d_1531_v63_: p2})
            d_1533_v65_: D17
            d_1533_v65_ = D17_DC45(d_1532_v64_)
            d_1532_v64_ = (d_1532_v64_ if not(p2) else (d_1532_v64_) | ((d_1533_v65_).cf67))
            d_1534_v66_: bool
            d_1534_v66_ = True
            d_1535_v67_: str
            d_1535_v67_ = _dafny.CodePoint('j')
            d_1536_v68_: _dafny.Map
            d_1536_v68_ = _dafny.Map({d_1535_v67_: d_1534_v66_})
            d_1537_v69_: _dafny.Array
            nw227_ = _dafny.Array(None, 8)
            nw227_[int(0)] = p2
            nw227_[int(1)] = False
            nw227_[int(2)] = (self).fm9(p0, globalState)
            nw227_[int(3)] = True
            nw227_[int(4)] = d_1534_v66_
            nw227_[int(5)] = ((d_1536_v68_)[d_1535_v67_] if (d_1535_v67_) in (d_1536_v68_) else p2)
            nw227_[int(6)] = p2
            nw227_[int(7)] = p2
            d_1537_v69_ = nw227_
            d_1538_v70_: _dafny.Set
            d_1538_v70_ = _dafny.Set({d_1537_v69_, d_1537_v69_, d_1537_v69_})
            d_1539_v71_: _dafny.Seq
            d_1539_v71_ = _dafny.SeqWithoutIsStrInference([d_1534_v66_, p2, not(d_1534_v66_)])
            rhs328_ = p2
            rhs329_ = (d_1539_v71_)[default__.safeIndex((p0) - (len(d_1531_v63_)), len(d_1539_v71_))]
            rhs330_ = d_1538_v70_
            d_1534_v66_ = rhs328_
            d_1534_v66_ = rhs329_
            d_1538_v70_ = rhs330_
            d_1540_v72_: D14
            d_1540_v72_ = D14_DC37(D14_DC36(d_1539_v71_))
            d_1541_v73_: D14
            d_1541_v73_ = D14_DC36(d_1539_v71_)
            d_1542_v74_: _dafny.Map
            d_1542_v74_ = _dafny.Map({p0: d_1534_v66_})
            pat_let_tv30_ = d_1541_v73_
            pat_let_tv31_ = d_1541_v73_
            d_1543_v75_: _dafny.Array
            nw228_ = _dafny.Array(None, 25)
            nw228_[int(0)] = d_1540_v72_
            nw228_[int(1)] = d_1540_v72_
            nw228_[int(2)] = d_1540_v72_
            def iife112_(_pat_let20_0):
                def iife113_(d_1544_dt__update__tmp_h0_):
                    def iife114_(_pat_let21_0):
                        def iife115_(d_1545_dt__update_hcf54_h0_):
                            return D14_DC37(d_1545_dt__update_hcf54_h0_)
                        return iife115_(_pat_let21_0)
                    return iife114_(pat_let_tv30_)
                return iife113_(_pat_let20_0)
            nw228_[int(3)] = iife112_(D14_DC37(d_1541_v73_))
            nw228_[int(4)] = d_1540_v72_
            def iife116_(_pat_let22_0):
                def iife117_(d_1546_dt__update__tmp_h1_):
                    def iife118_(_pat_let23_0):
                        def iife119_(d_1547_dt__update_hcf54_h1_):
                            return D14_DC37(d_1547_dt__update_hcf54_h1_)
                        return iife119_(_pat_let23_0)
                    return iife118_(pat_let_tv31_)
                return iife117_(_pat_let22_0)
            nw228_[int(5)] = iife116_(d_1540_v72_)
            nw228_[int(6)] = d_1540_v72_
            nw228_[int(7)] = d_1540_v72_
            nw228_[int(8)] = d_1540_v72_
            nw228_[int(9)] = d_1540_v72_
            nw228_[int(10)] = d_1540_v72_
            nw228_[int(11)] = D14_DC37(d_1541_v73_)
            nw228_[int(12)] = d_1540_v72_
            nw228_[int(13)] = D14_DC37(default__.fm48(d_1534_v66_, p2, ((d_1542_v74_)[p0] if (p0) in (d_1542_v74_) else p2), globalState))
            nw228_[int(14)] = d_1540_v72_
            nw228_[int(15)] = d_1540_v72_
            nw228_[int(16)] = d_1540_v72_
            nw228_[int(17)] = d_1540_v72_
            nw228_[int(18)] = d_1540_v72_
            nw228_[int(19)] = d_1540_v72_
            nw228_[int(20)] = (d_1540_v72_ if p2 else d_1540_v72_)
            nw228_[int(21)] = D14_DC37(d_1541_v73_)
            nw228_[int(22)] = d_1540_v72_
            nw228_[int(23)] = d_1540_v72_
            nw228_[int(24)] = D14_DC37(D14_DC35(_dafny.SeqWithoutIsStrInference([(d_1542_v74_).set(p0, p2) for d_1548_i5_ in range(default__.abs(969))])))
            d_1543_v75_ = nw228_
            index273_ = default__.safeIndex(983, (d_1543_v75_).length(0))
            (d_1543_v75_)[index273_] = d_1540_v72_
            index274_ = default__.safeIndex(983, (d_1543_v75_).length(0))
            (d_1543_v75_)[index274_] = d_1540_v72_
        if not(((p0) <= (p0)) and (not((p0) >= (807)))):
            d_1452_v2_ = (d_1452_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiseggug")))
            d_1549_v76_: C0
            nw229_ = C0()
            nw229_.ctor__()
            d_1549_v76_ = nw229_
            d_1550_v77_: _dafny.Array
            def lambda78_(d_1551_i6_):
                return _dafny.CodePoint('e')

            init39_ = lambda78_
            nw230_ = _dafny.Array(None, 18)
            for i0_39_ in range(nw230_.length(0)):
                nw230_[i0_39_] = init39_(i0_39_)
            d_1550_v77_ = nw230_
            d_1550_v77_ = d_1550_v77_
            if not((len(d_1450_v0_)) == (p0)):
                (globalState).f13 = default__.safeModuloInt((0) - (p0), p0)
                d_1552_v78_: _dafny.Seq
                d_1552_v78_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, p2, p2])
                d_1553_v79_: _dafny.Seq
                d_1553_v79_ = _dafny.SeqWithoutIsStrInference([p0, len(d_1552_v78_)])
                d_1554_v80_: _dafny.Seq
                d_1554_v80_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_1553_v79_)])
                d_1555_v81_: _dafny.Seq
                d_1555_v81_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0]), d_1554_v80_])
                d_1556_v82_: _dafny.Map
                d_1556_v82_ = _dafny.Map({p2: d_1452_v2_})
                d_1557_v83_: str
                d_1557_v83_ = _dafny.CodePoint('p')
                d_1558_v84_: _dafny.Set
                d_1558_v84_ = _dafny.Set({p2})
                d_1559_v85_: _dafny.Seq
                d_1559_v85_ = _dafny.SeqWithoutIsStrInference([d_1558_v84_])
                d_1560_v86_: _dafny.Map
                d_1560_v86_ = _dafny.Map({d_1559_v85_: d_1452_v2_})
                d_1561_v87_: _dafny.Map
                d_1561_v87_ = _dafny.Map({p2: p0})
                d_1562_v88_: T0
                nw231_ = C2()
                nw231_.ctor__()
                d_1562_v88_ = nw231_
                d_1563_v89_: _dafny.Seq
                d_1563_v89_ = _dafny.SeqWithoutIsStrInference([default__.fm49(p0, len(_dafny.Set({d_1562_v88_, d_1562_v88_})), globalState)])
                d_1564_v90_: _dafny.Map
                d_1564_v90_ = _dafny.Map({len((_dafny.Map({p2: (p1).cardinality})).set(p2, -649)): _dafny.Set({len(d_1561_v87_), len(d_1563_v89_)})})
                d_1565_v91_: _dafny.Map
                d_1565_v91_ = _dafny.Map({p0: (d_1555_v81_)[default__.safeIndex(p0, len(d_1555_v81_))]})
                d_1566_v92_: _dafny.Array
                nw232_ = _dafny.Array(None, 27)
                nw232_[int(0)] = (d_1549_v76_).fm28(d_1555_v81_, globalState)
                nw232_[int(1)] = (d_1452_v2_) + (d_1452_v2_)
                nw232_[int(2)] = ((((d_1556_v82_)[p2] if (p2) in (d_1556_v82_) else d_1452_v2_)).set(default__.safeIndex((0) - (p0), len(((d_1556_v82_)[p2] if (p2) in (d_1556_v82_) else d_1452_v2_))), d_1557_v83_) if p2 else d_1452_v2_)
                nw232_[int(3)] = d_1452_v2_
                nw232_[int(4)] = d_1452_v2_
                nw232_[int(5)] = (d_1452_v2_) + (d_1452_v2_)
                nw232_[int(6)] = _dafny.SeqWithoutIsStrInference([d_1557_v83_ for d_1567_i7_ in range(default__.abs(145))])
                nw232_[int(7)] = ((d_1560_v86_)[_dafny.SeqWithoutIsStrInference([_dafny.Set({p2, p2, p2}), _dafny.Set({p2, False}), d_1558_v84_, d_1558_v84_])] if (_dafny.SeqWithoutIsStrInference([_dafny.Set({p2, p2, p2}), _dafny.Set({p2, False}), d_1558_v84_, d_1558_v84_])) in (d_1560_v86_) else d_1452_v2_)
                nw232_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuauuuwp"))
                nw232_[int(9)] = default__.fm30(535, p2, _dafny.Set({len(d_1564_v90_), 769}), p2, globalState)
                nw232_[int(10)] = (d_1452_v2_ if p2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiqjnpdrf")))
                nw232_[int(11)] = (d_1452_v2_) + (d_1452_v2_)
                nw232_[int(12)] = d_1452_v2_
                nw232_[int(13)] = (d_1452_v2_) + ((d_1549_v76_).fm28(d_1555_v81_, globalState))
                nw232_[int(14)] = d_1452_v2_
                nw232_[int(15)] = (d_1452_v2_) + (d_1452_v2_)
                nw232_[int(16)] = d_1452_v2_
                nw232_[int(17)] = (d_1452_v2_) + (d_1452_v2_)
                nw232_[int(18)] = (d_1549_v76_).fm28(_dafny.SeqWithoutIsStrInference([((d_1565_v91_)[(_dafny.MultiSet([p2])).cardinality] if ((_dafny.MultiSet([p2])).cardinality) in (d_1565_v91_) else d_1553_v79_), d_1554_v80_]), globalState)
                nw232_[int(19)] = d_1452_v2_
                nw232_[int(20)] = ((d_1556_v82_)[p2] if (p2) in (d_1556_v82_) else d_1452_v2_)
                nw232_[int(21)] = d_1452_v2_
                nw232_[int(22)] = d_1452_v2_
                nw232_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxbmntpf"))
                nw232_[int(24)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggfo"))) + (d_1452_v2_)
                nw232_[int(25)] = d_1452_v2_
                nw232_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkqchmef"))
                d_1566_v92_ = nw232_
                index275_ = default__.safeIndex(61, (d_1566_v92_).length(0))
                (d_1566_v92_)[index275_] = (d_1452_v2_) + (d_1452_v2_)
                index276_ = default__.safeIndex(61, (d_1566_v92_).length(0))
                (d_1566_v92_)[index276_] = d_1452_v2_
                (globalState).f15 = p0
                d_1568_v93_: _dafny.Array
                nw233_ = _dafny.Array(int(0), 22)
                d_1568_v93_ = nw233_
                d_1569_v94_: D2
                d_1569_v94_ = D2_DC5(d_1568_v93_)
                d_1569_v94_ = D2_DC5(d_1568_v93_)
                (globalState).f4 = 493
            elif True:
                d_1570_v95_: bool
                d_1570_v95_ = True
                d_1571_v96_: _dafny.Array
                nw234_ = _dafny.Array(int(0), 1)
                d_1571_v96_ = nw234_
                index277_ = default__.safeIndex(529, (d_1571_v96_).length(0))
                (d_1571_v96_)[index277_] = p0
                d_1572_v97_: _dafny.Seq
                d_1572_v97_ = _dafny.SeqWithoutIsStrInference([p2])
                d_1573_v98_: _dafny.Map
                d_1573_v98_ = _dafny.Map({p2: d_1570_v95_})
                d_1574_v99_: _dafny.Seq
                d_1574_v99_ = _dafny.SeqWithoutIsStrInference([15, -836])
                d_1575_v100_: D15
                d_1575_v100_ = D15_DC39((d_1573_v98_).set(d_1570_v95_, d_1570_v95_), d_1570_v95_, len(d_1574_v99_))
                index278_ = default__.safeIndex(529, (d_1571_v96_).length(0))
                rhs331_ = (p2) in (p1)
                rhs332_ = (p0) + ((len(d_1573_v98_)) - (p0))
                rhs333_ = d_1572_v97_
                rhs334_ = (d_1575_v100_).cf57
                rhs335_ = d_1570_v95_
                lhs236_ = d_1571_v96_
                lhs237_ = default__.safeIndex(529, (d_1571_v96_).length(0))
                d_1570_v95_ = rhs331_
                lhs236_[lhs237_] = rhs332_
                d_1572_v97_ = rhs333_
                d_1570_v95_ = rhs334_
                d_1570_v95_ = rhs335_
                d_1576_v101_: D14
                d_1576_v101_ = D14_DC35(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1571_v96_)[default__.safeIndex(529, (d_1571_v96_).length(0))]: p2}) for d_1577_i8_ in range(default__.abs(7))]))
                d_1578_v102_: D14
                d_1578_v102_ = D14_DC35((d_1576_v101_).cf52)
                d_1579_v103_: D14
                d_1579_v103_ = D14_DC37(d_1578_v102_)
                d_1579_v103_ = d_1579_v103_
                d_1570_v95_ = d_1570_v95_
                d_1580_v104_: _dafny.Array
                def lambda79_(d_1581_p1_):
                    def lambda80_(d_1582_i9_):
                        return (d_1581_p1_).ispropersubset(d_1581_p1_)

                    return lambda80_

                init40_ = lambda79_(p1)
                nw235_ = _dafny.Array(None, 29)
                for i0_40_ in range(nw235_.length(0)):
                    nw235_[i0_40_] = init40_(i0_40_)
                d_1580_v104_ = nw235_
                index279_ = default__.safeIndex(47, (d_1580_v104_).length(0))
                (d_1580_v104_)[index279_] = p2
                d_1583_v105_: _dafny.MultiSet
                d_1583_v105_ = _dafny.MultiSet([(d_1571_v96_)[default__.safeIndex(529, (d_1571_v96_).length(0))], (d_1574_v99_)[default__.safeIndex((d_1571_v96_)[default__.safeIndex(529, (d_1571_v96_).length(0))], len(d_1574_v99_))]])
                index280_ = default__.safeIndex(47, (d_1580_v104_).length(0))
                (d_1580_v104_)[index280_] = (d_1583_v105_).issubset((_dafny.MultiSet([(d_1571_v96_)[default__.safeIndex(529, (d_1571_v96_).length(0))], p0])) | (default__.fm31((d_1571_v96_)[default__.safeIndex(529, (d_1571_v96_).length(0))], p0, globalState)))
                d_1584_v106_: _dafny.Map
                d_1584_v106_ = _dafny.Map({d_1574_v99_: not((d_1580_v104_)[default__.safeIndex(47, (d_1580_v104_).length(0))])})
                d_1584_v106_ = d_1584_v106_
            (globalState).f13 = ((self).fm2(p0, globalState)) - (p0)
        elif True:
            d_1585_v107_: _dafny.Set
            d_1585_v107_ = _dafny.Set({p2, p2})
            d_1586_v108_: _dafny.Map
            d_1586_v108_ = _dafny.Map({p2: False})
            d_1587_v109_: _dafny.Map
            d_1587_v109_ = _dafny.Map({p2: p0})
            d_1588_v110_: _dafny.Seq
            d_1588_v110_ = _dafny.SeqWithoutIsStrInference([not((self).fm8(len(d_1585_v107_), ((d_1586_v108_)[True] if (True) in (d_1586_v108_) else not(not(p2))), ((d_1587_v109_)[p2] if (p2) in (d_1587_v109_) else (0) - ((self).fm2(p0, globalState))), globalState)), p2])
            d_1589_v111_: _dafny.Array
            nw236_ = _dafny.Array(None, 24)
            nw236_[int(0)] = (0) - (default__.safeDivisionInt((0) - (p0), p0))
            nw236_[int(1)] = (p0) - (len(d_1588_v110_))
            nw236_[int(2)] = p0
            nw236_[int(3)] = (0) - (p0)
            nw236_[int(4)] = p0
            nw236_[int(5)] = p0
            nw236_[int(6)] = p0
            nw236_[int(7)] = 75
            nw236_[int(8)] = p0
            nw236_[int(9)] = p0
            nw236_[int(10)] = p0
            nw236_[int(11)] = 917
            nw236_[int(12)] = p0
            nw236_[int(13)] = default__.safeDivisionInt(len(d_1585_v107_), p0)
            nw236_[int(14)] = 265
            nw236_[int(15)] = p0
            nw236_[int(16)] = p0
            nw236_[int(17)] = p0
            nw236_[int(18)] = p0
            nw236_[int(19)] = p0
            nw236_[int(20)] = p0
            nw236_[int(21)] = (p1).cardinality
            nw236_[int(22)] = (self).fm2(default__.fm17(len(d_1452_v2_), globalState), globalState)
            nw236_[int(23)] = p0
            d_1589_v111_ = nw236_
            d_1589_v111_ = d_1589_v111_
            (globalState).f4 = p0
            d_1590_v112_: _dafny.Array
            def lambda81_(d_1591_v110_, d_1592_p0_):
                def lambda82_(d_1593_i10_):
                    return D12_DC32(_dafny.Set({(d_1591_v110_)[default__.safeIndex(d_1592_p0_, len(d_1591_v110_))]}))

                return lambda82_

            init41_ = lambda81_(d_1588_v110_, p0)
            nw237_ = _dafny.Array(None, 28)
            for i0_41_ in range(nw237_.length(0)):
                nw237_[i0_41_] = init41_(i0_41_)
            d_1590_v112_ = nw237_
            d_1594_v113_: D12
            d_1594_v113_ = D12_DC32(_dafny.Set({p2}))
            index281_ = default__.safeIndex(615, (d_1590_v112_).length(0))
            (d_1590_v112_)[index281_] = d_1594_v113_
            d_1595_v114_: _dafny.Array
            nw238_ = _dafny.Array(D4.default()(), 15)
            d_1595_v114_ = nw238_
            d_1596_v115_: D4
            d_1596_v115_ = D4_DC11(_dafny.SeqWithoutIsStrInference([d_1452_v2_ for d_1597_i11_ in range(default__.abs(608))]))
            index282_ = default__.safeIndex(46, (d_1595_v114_).length(0))
            (d_1595_v114_)[index282_] = d_1596_v115_
            index283_ = default__.safeIndex(615, (d_1590_v112_).length(0))
            index284_ = default__.safeIndex(46, (d_1595_v114_).length(0))
            rhs336_ = (p0 if p2 else p0)
            rhs337_ = d_1585_v107_
            rhs338_ = D12_DC32(d_1585_v107_)
            rhs339_ = p0
            rhs340_ = d_1596_v115_
            lhs238_ = globalState
            lhs239_ = d_1590_v112_
            lhs240_ = default__.safeIndex(615, (d_1590_v112_).length(0))
            lhs241_ = globalState
            lhs242_ = d_1595_v114_
            lhs243_ = default__.safeIndex(46, (d_1595_v114_).length(0))
            lhs238_.f4 = rhs336_
            d_1585_v107_ = rhs337_
            lhs239_[lhs240_] = rhs338_
            lhs241_.f4 = rhs339_
            lhs242_[lhs243_] = rhs340_
            d_1598_v116_: _dafny.Seq
            d_1598_v116_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1599_v117_: _dafny.MultiSet
            d_1599_v117_ = _dafny.MultiSet([p0, len(d_1598_v116_), p0, len(d_1452_v2_)])
            d_1600_v119_: _dafny.MultiSet
            def iife120_():
                coll72_ = _dafny.Map()
                compr_72_: int
                for compr_72_ in _dafny.IntegerRange(436, 908):
                    d_1601_v118_: int = compr_72_
                    if ((436) <= (d_1601_v118_)) and ((d_1601_v118_) < (908)):
                        coll72_[default__.safeDivisionInt(d_1601_v118_, p0)] = p2
                return _dafny.Map(coll72_)
            d_1600_v119_ = _dafny.MultiSet([(p0) * (p0), -763, (0) - (((d_1599_v117_)[p0] if (p0) in (d_1599_v117_) else (0) - (p0))), (len(_dafny.Map({not(p2): len(iife120_()
            )})) if False else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1602_i12_ in range(default__.abs(815))]))), p0])
            d_1599_v117_ = d_1599_v117_
            if p2:
                d_1603_v120_: _dafny.MultiSet
                d_1603_v120_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0 for d_1604_i13_ in range(default__.abs(279))])])
                d_1605_v121_: D1
                d_1605_v121_ = D1_DC3(d_1452_v2_, not(not(True)))
                d_1606_v122_: _dafny.Map
                d_1606_v122_ = _dafny.Map({(0) - ((((d_1603_v120_).set(d_1598_v116_, default__.abs(len(d_1585_v107_)))).set(d_1598_v116_, default__.abs(-51))).cardinality): (d_1605_v121_).cf3})
                d_1606_v122_ = (_dafny.Map({p0: d_1452_v2_}) if (d_1585_v107_).issubset(_dafny.Set({p2})) else (d_1606_v122_) | (d_1606_v122_))
                d_1607_v123_: bool
                d_1607_v123_ = True
                d_1607_v123_ = (default__.safeDivisionInt(p0, p0)) > (p0)
                d_1608_v124_: _dafny.Array
                nw239_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1608_v124_ = nw239_
                d_1609_v125_: _dafny.Array
                nw240_ = _dafny.Array(False, 3)
                d_1609_v125_ = nw240_
                d_1610_v126_: _dafny.Map
                d_1610_v126_ = _dafny.Map({False: d_1609_v125_})
                index285_ = default__.safeIndex(305, (d_1608_v124_).length(0))
                (d_1608_v124_)[index285_] = ((d_1610_v126_)[True] if (True) in (d_1610_v126_) else d_1609_v125_)
                index286_ = default__.safeIndex(305, (d_1608_v124_).length(0))
                (d_1608_v124_)[index286_] = d_1609_v125_
                index287_ = default__.safeIndex(305, (d_1608_v124_).length(0))
                (d_1608_v124_)[index287_] = d_1609_v125_
                d_1611_v127_: D8
                d_1611_v127_ = D8_DC25((p0) - (p0))
                d_1611_v127_ = d_1611_v127_
            elif True:
                d_1612_v128_: str
                d_1612_v128_ = _dafny.CodePoint('x')
                d_1613_v129_: _dafny.Array
                nw241_ = _dafny.Array(None, 17)
                nw241_[int(0)] = d_1452_v2_
                nw241_[int(1)] = d_1452_v2_
                nw241_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1614_i14_ in range(default__.abs(-789))])
                nw241_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxum"))
                nw241_[int(4)] = d_1452_v2_
                nw241_[int(5)] = default__.fm19(p2, not(p2), globalState)
                nw241_[int(6)] = d_1452_v2_
                nw241_[int(7)] = d_1452_v2_
                nw241_[int(8)] = d_1452_v2_
                nw241_[int(9)] = d_1452_v2_
                nw241_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhvdynhug"))
                nw241_[int(11)] = (d_1452_v2_).set(default__.safeIndex(default__.fm17((0) - (p0), globalState), len(d_1452_v2_)), d_1612_v128_)
                nw241_[int(12)] = d_1452_v2_
                nw241_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1615_i15_ in range(default__.abs(886))])
                nw241_[int(14)] = d_1452_v2_
                nw241_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                nw241_[int(16)] = d_1452_v2_
                d_1613_v129_ = nw241_
                d_1616_v130_: _dafny.Map
                d_1616_v130_ = _dafny.Map({d_1613_v129_: (p0) == (p0)})
                d_1616_v130_ = (d_1616_v130_).set(d_1613_v129_, p2)
                d_1617_v131_: bool
                d_1617_v131_ = True
                rhs341_ = p0
                rhs342_ = (d_1617_v131_) and ((p1).ispropersubset(_dafny.MultiSet([d_1617_v131_, p2])))
                rhs343_ = p2
                lhs244_ = globalState
                lhs244_.f4 = rhs341_
                d_1617_v131_ = rhs342_
                d_1617_v131_ = rhs343_
                d_1618_v132_: _dafny.Array
                def lambda83_(d_1619_v131_):
                    def lambda84_(d_1620_i16_):
                        return d_1619_v131_

                    return lambda84_

                init42_ = lambda83_(d_1617_v131_)
                nw242_ = _dafny.Array(None, 6)
                for i0_42_ in range(nw242_.length(0)):
                    nw242_[i0_42_] = init42_(i0_42_)
                d_1618_v132_ = nw242_
                d_1621_v133_: _dafny.Map
                d_1621_v133_ = _dafny.Map({not(d_1617_v131_): d_1586_v108_})
                index288_ = default__.safeIndex(469, (d_1618_v132_).length(0))
                (d_1618_v132_)[index288_] = (self).fm1(d_1621_v133_, len(d_1588_v110_), globalState)
                index289_ = default__.safeIndex(469, (d_1618_v132_).length(0))
                (d_1618_v132_)[index289_] = False
                d_1622_v134_: C7
                nw243_ = C7()
                nw243_.ctor__()
                d_1622_v134_ = nw243_
                d_1623_v135_: _dafny.Map
                d_1623_v135_ = _dafny.Map({d_1618_v132_: d_1618_v132_})
                d_1624_v136_: _dafny.Map
                d_1624_v136_ = _dafny.Map({d_1618_v132_: d_1623_v135_})
                d_1623_v135_ = ((d_1624_v136_)[d_1618_v132_] if (d_1618_v132_) in (d_1624_v136_) else d_1623_v135_)
        d_1625_v137_: _dafny.Array
        nw244_ = _dafny.Array(False, 20)
        d_1625_v137_ = nw244_
        d_1626_v138_: _dafny.Map
        d_1626_v138_ = _dafny.Map({d_1625_v137_: p2})
        d_1626_v138_ = (d_1626_v138_).set(d_1625_v137_, p2)
        hi11_ = p0
        for d_1627_i17_ in range(default__.safeDivisionInt(p0, p0), hi11_):
            d_1625_v137_ = d_1625_v137_
            d_1628_v139_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1628_v139_ = nw245_
            d_1629_v140_: _dafny.Array
            def lambda85_(d_1630_p0_):
                def lambda86_(d_1631_i18_):
                    return default__.safeModuloInt(d_1631_i18_, d_1630_p0_)

                return lambda86_

            init43_ = lambda85_(p0)
            nw246_ = _dafny.Array(None, 20)
            for i0_43_ in range(nw246_.length(0)):
                nw246_[i0_43_] = init43_(i0_43_)
            d_1629_v140_ = nw246_
            index290_ = default__.safeIndex(342, (d_1628_v139_).length(0))
            (d_1628_v139_)[index290_] = d_1629_v140_
            d_1632_v141_: _dafny.MultiSet
            d_1632_v141_ = _dafny.MultiSet([(d_1627_i17_ if p2 else -605), (d_1627_i17_) * (d_1627_i17_)])
            d_1633_v142_: _dafny.Seq
            d_1633_v142_ = _dafny.SeqWithoutIsStrInference([d_1632_v141_, d_1632_v141_])
            d_1634_v143_: T0
            nw247_ = C7()
            nw247_.ctor__()
            d_1634_v143_ = nw247_
            d_1635_v144_: _dafny.Seq
            d_1635_v144_ = _dafny.SeqWithoutIsStrInference([((d_1633_v142_)[default__.safeIndex(p0, len(d_1633_v142_))]).set(d_1627_i17_, default__.abs((D17_DC47(d_1634_v143_, True, d_1627_i17_, d_1627_i17_, default__.fm17(d_1627_i17_, globalState))).cf72))])
            index291_ = default__.safeIndex(342, (d_1628_v139_).length(0))
            rhs344_ = d_1629_v140_
            rhs345_ = (p1).cardinality
            rhs346_ = (d_1635_v144_)[default__.safeIndex(d_1627_i17_, len(d_1635_v144_))]
            lhs245_ = d_1628_v139_
            lhs246_ = default__.safeIndex(342, (d_1628_v139_).length(0))
            lhs247_ = globalState
            lhs245_[lhs246_] = rhs344_
            lhs247_.f4 = rhs345_
            d_1632_v141_ = rhs346_
            if p2:
                d_1636_v145_: bool
                d_1636_v145_ = True
                d_1637_v146_: _dafny.Map
                d_1637_v146_ = _dafny.Map({p0: (d_1634_v143_).fm2(p0, globalState)})
                d_1638_v147_: _dafny.Set
                d_1638_v147_ = _dafny.Set({d_1637_v146_, d_1637_v146_})
                d_1636_v145_ = (d_1638_v147_).issubset(d_1638_v147_)
                (globalState).f15 = ((_dafny.MultiSet([-81])).cardinality) - (default__.safeModuloInt((d_1632_v141_).cardinality, d_1627_i17_))
                d_1636_v145_ = d_1636_v145_
                d_1639_v148_: _dafny.MultiSet
                d_1639_v148_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1627_i17_ for d_1640_i19_ in range(default__.abs(910))]), _dafny.SeqWithoutIsStrInference([p0 for d_1641_i20_ in range(default__.abs(285))])])
                d_1642_v149_: _dafny.Seq
                d_1642_v149_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1643_v150_: str
                d_1643_v150_ = _dafny.CodePoint('y')
                d_1644_v151_: _dafny.Set
                d_1644_v151_ = _dafny.Set({(self).fm9(d_1627_i17_, globalState)})
                rhs347_ = p2
                rhs348_ = _dafny.MultiSet([d_1642_v149_, default__.fm39(d_1627_i17_, d_1643_v150_, d_1642_v149_, d_1636_v145_, globalState)])
                rhs349_ = (751) != (len(d_1644_v151_))
                d_1636_v145_ = rhs347_
                d_1639_v148_ = rhs348_
                d_1636_v145_ = rhs349_
                d_1645_v152_: _dafny.Map
                d_1645_v152_ = _dafny.Map({d_1636_v145_: p0})
                d_1646_v153_: _dafny.Array
                nw248_ = _dafny.Array(None, 13)
                nw248_[int(0)] = d_1645_v152_
                nw248_[int(1)] = d_1645_v152_
                nw248_[int(2)] = d_1645_v152_
                nw248_[int(3)] = (d_1645_v152_).set(d_1636_v145_, d_1627_i17_)
                nw248_[int(4)] = d_1645_v152_
                nw248_[int(5)] = d_1645_v152_
                nw248_[int(6)] = d_1645_v152_
                nw248_[int(7)] = default__.fm44(d_1627_i17_, d_1636_v145_, globalState)
                nw248_[int(8)] = _dafny.Map({d_1636_v145_: p0})
                nw248_[int(9)] = d_1645_v152_
                nw248_[int(10)] = d_1645_v152_
                nw248_[int(11)] = d_1645_v152_
                nw248_[int(12)] = _dafny.Map({p2: p0})
                d_1646_v153_ = nw248_
                d_1647_v154_: _dafny.Map
                d_1647_v154_ = _dafny.Map({d_1646_v153_: p2})
                d_1636_v145_ = ((d_1647_v154_)[d_1646_v153_] if (d_1646_v153_) in (d_1647_v154_) else (False if p2 else False))
            elif True:
                d_1648_v155_: D1
                d_1648_v155_ = D1_DC4(d_1632_v141_)
                d_1632_v141_ = (d_1632_v141_).intersection((d_1648_v155_).cf5)
                d_1649_v156_: C8
                nw249_ = C8()
                nw249_.ctor__()
                d_1649_v156_ = nw249_
                d_1650_v157_: _dafny.Set
                d_1650_v157_ = _dafny.Set({p2, True, p2})
                d_1651_v158_: _dafny.Map
                d_1651_v158_ = _dafny.Map({p2: d_1650_v157_})
                d_1652_v159_: _dafny.MultiSet
                d_1652_v159_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1653_i21_ in range(default__.abs(768))])])
                d_1651_v158_ = (d_1651_v158_).set((d_1652_v159_) != (d_1652_v159_), d_1650_v157_)
                d_1654_v160_: D1
                d_1654_v160_ = D1_DC2(p0)
                (globalState).f13 = (default__.safeModuloInt(p0, p0)) - ((d_1654_v160_).cf2)
                d_1650_v157_ = _dafny.Set({p2, p2, p2, p2, p2})
            d_1655_v161_: str
            d_1655_v161_ = _dafny.CodePoint('u')
            d_1656_v162_: _dafny.Map
            d_1656_v162_ = _dafny.Map({d_1625_v137_: d_1655_v161_})
            d_1657_v163_: _dafny.Map
            d_1657_v163_ = _dafny.Map({(p0) - (p0): ((d_1656_v162_)[d_1625_v137_] if (d_1625_v137_) in (d_1656_v162_) else d_1655_v161_)})
            d_1657_v163_ = (d_1657_v163_).set(115, d_1655_v161_)
        d_1658_v164_: _dafny.Array
        def lambda87_(d_1659_i22_):
            return (d_1659_i22_) + (871)

        init44_ = lambda87_
        nw250_ = _dafny.Array(None, 27)
        for i0_44_ in range(nw250_.length(0)):
            nw250_[i0_44_] = init44_(i0_44_)
        d_1658_v164_ = nw250_
        d_1658_v164_ = d_1658_v164_
        hi12_ = p0
        for d_1660_i23_ in range(p0, hi12_):
            d_1661_v165_: _dafny.Array
            nw251_ = _dafny.Array(_dafny.CodePoint('D'), 14)
            d_1661_v165_ = nw251_
            nw252_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_1661_v165_ = nw252_
            d_1662_v166_: T0
            nw253_ = C7()
            nw253_.ctor__()
            d_1662_v166_ = nw253_
            d_1662_v166_ = d_1662_v166_
            d_1663_v167_: bool
            d_1663_v167_ = False
            d_1663_v167_ = d_1663_v167_
            d_1663_v167_ = d_1663_v167_

    def m7(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_1664_v0_: _dafny.Array
        nw254_ = _dafny.Array(False, 25)
        d_1664_v0_ = nw254_
        d_1665_v1_: _dafny.MultiSet
        d_1665_v1_ = _dafny.MultiSet([d_1664_v0_, d_1664_v0_])
        r0 = ((d_1665_v1_)[d_1664_v0_] if (d_1664_v0_) in (d_1665_v1_) else p0)
        d_1666_i0_: int
        d_1666_i0_ = 0
        with _dafny.label("10"):
            while p1:
                with _dafny.c_label("10"):
                    if (d_1666_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1666_i0_ = (d_1666_i0_) + (1)
                    if True:
                        d_1667_v2_: _dafny.Array
                        def lambda88_(d_1668_i1_):
                            return D3_DC10(65, 626)

                        init45_ = lambda88_
                        nw255_ = _dafny.Array(None, 7)
                        for i0_45_ in range(nw255_.length(0)):
                            nw255_[i0_45_] = init45_(i0_45_)
                        d_1667_v2_ = nw255_
                        d_1669_v3_: _dafny.MultiSet
                        d_1669_v3_ = _dafny.MultiSet([p1])
                        d_1670_v4_: _dafny.Seq
                        d_1670_v4_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                        d_1671_v5_: _dafny.Map
                        d_1671_v5_ = _dafny.Map({len(d_1670_v4_): p0})
                        d_1672_v6_: _dafny.Seq
                        d_1672_v6_ = _dafny.SeqWithoutIsStrInference([d_1671_v5_])
                        d_1673_v7_: _dafny.Seq
                        d_1673_v7_ = _dafny.SeqWithoutIsStrInference([((d_1669_v3_).set(p1, default__.abs(p0))).cardinality, (self).fm2(p0, globalState), len(d_1672_v6_)])
                        d_1674_v8_: D3
                        d_1674_v8_ = D3_DC10(len(d_1673_v7_), p0)
                        index292_ = default__.safeIndex(36, (d_1667_v2_).length(0))
                        (d_1667_v2_)[index292_] = d_1674_v8_
                        d_1675_v9_: _dafny.Set
                        d_1675_v9_ = _dafny.Set({not(p1)})
                        d_1676_v10_: _dafny.Set
                        d_1676_v10_ = _dafny.Set({d_1675_v9_, d_1675_v9_})
                        pat_let_tv32_ = d_1671_v5_
                        pat_let_tv33_ = d_1676_v10_
                        pat_let_tv34_ = d_1676_v10_
                        pat_let_tv35_ = d_1671_v5_
                        pat_let_tv36_ = p0
                        pat_let_tv37_ = p0
                        index293_ = default__.safeIndex(36, (d_1667_v2_).length(0))
                        def iife121_(_pat_let24_0):
                            def iife122_(d_1677_dt__update__tmp_h0_):
                                def iife123_(_pat_let25_0):
                                    def iife124_(d_1678_dt__update_hcf14_h0_):
                                        return D3_DC10((d_1677_dt__update__tmp_h0_).cf13, d_1678_dt__update_hcf14_h0_)
                                    return iife124_(_pat_let25_0)
                                return iife123_(default__.safeModuloInt(((pat_let_tv32_)[len(pat_let_tv33_)] if (len(pat_let_tv34_)) in (pat_let_tv35_) else pat_let_tv36_), pat_let_tv37_))
                            return iife122_(_pat_let24_0)
                        (d_1667_v2_)[index293_] = iife121_(d_1674_v8_)
                        d_1679_v11_: C5
                        nw256_ = C5()
                        nw256_.ctor__()
                        d_1679_v11_ = nw256_
                        (globalState).f15 = default__.safeModuloInt(p0, (0) - ((len(_dafny.Map({p0: d_1679_v11_})) if p1 else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmqug"))))))
                        d_1680_v12_: D18
                        d_1680_v12_ = D18_DC48(d_1671_v5_)
                        d_1681_v13_: _dafny.Map
                        d_1681_v13_ = _dafny.Map({p1: p1})
                        d_1682_v14_: _dafny.Map
                        d_1682_v14_ = _dafny.Map({p1: d_1681_v13_})
                        pat_let_tv38_ = d_1671_v5_
                        d_1683_v15_: _dafny.Map
                        def iife125_(_pat_let26_0):
                            def iife126_(d_1684_dt__update__tmp_h1_):
                                def iife127_(_pat_let27_0):
                                    def iife128_(d_1685_dt__update_hcf74_h0_):
                                        return D18_DC48(d_1685_dt__update_hcf74_h0_)
                                    return iife128_(_pat_let27_0)
                                return iife127_(pat_let_tv38_)
                            return iife126_(_pat_let26_0)
                        d_1683_v15_ = _dafny.Map({len(_dafny.Set({p0, p0, (0) - (len((iife125_(d_1680_v12_)).cf74))})): (self).fm1(d_1682_v14_, p0, globalState)})
                        d_1686_v17_: _dafny.Set
                        d_1686_v17_ = _dafny.Set({(0) - (p0), p0})
                        d_1687_v18_: _dafny.Seq
                        def iife129_():
                            coll73_ = _dafny.Set()
                            compr_73_: int
                            for compr_73_ in (d_1683_v15_).keys.Elements:
                                d_1688_v16_: int = compr_73_
                                if (d_1688_v16_) in (d_1683_v15_):
                                    coll73_ = coll73_.union(_dafny.Set([(d_1688_v16_) - (-68)]))
                            return _dafny.Set(coll73_)
                        d_1687_v18_ = _dafny.SeqWithoutIsStrInference([iife129_()
                        , d_1686_v17_])
                        d_1689_v19_: _dafny.Map
                        d_1689_v19_ = _dafny.Map({d_1686_v17_: _dafny.Set({p0, len(d_1673_v7_)})})
                        d_1690_v21_: _dafny.Seq
                        def iife130_():
                            coll74_ = _dafny.Set()
                            compr_74_: int
                            for compr_74_ in _dafny.IntegerRange(594, 392):
                                d_1691_v20_: int = compr_74_
                                if ((594) <= (d_1691_v20_)) and ((d_1691_v20_) < (392)):
                                    coll74_ = coll74_.union(_dafny.Set([(d_1691_v20_) + (p0)]))
                            return _dafny.Set(coll74_)
                        def iife131_():
                            coll75_ = _dafny.Set()
                            compr_75_: int
                            for compr_75_ in _dafny.IntegerRange(594, 392):
                                d_1692_v20_: int = compr_75_
                                if ((594) <= (d_1692_v20_)) and ((d_1692_v20_) < (392)):
                                    coll75_ = coll75_.union(_dafny.Set([(d_1692_v20_) + (p0)]))
                            return _dafny.Set(coll75_)
                        d_1690_v21_ = _dafny.SeqWithoutIsStrInference([(d_1687_v18_)[default__.safeIndex(398, len(d_1687_v18_))], ((d_1689_v19_)[iife130_()
                        ] if (iife131_()
                        ) in (d_1689_v19_) else d_1686_v17_), _dafny.Set({(0) - (p0), p0, p0}), d_1686_v17_, d_1686_v17_])
                        (globalState).f13 = len(d_1690_v21_)
                        d_1693_v22_: _dafny.Array
                        nw257_ = _dafny.Array(int(0), 22)
                        d_1693_v22_ = nw257_
                        d_1694_v23_: _dafny.Map
                        d_1694_v23_ = _dafny.Map({d_1693_v22_: p0})
                        d_1695_v24_: _dafny.Seq
                        d_1695_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxs"))
                        d_1696_v25_: _dafny.Map
                        d_1696_v25_ = _dafny.Map({d_1695_v24_: d_1694_v23_})
                        index294_ = default__.safeIndex(636, (d_1664_v0_).length(0))
                        (d_1664_v0_)[index294_] = (d_1694_v23_) != (((d_1696_v25_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idknhjqof"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idknhjqof"))) in (d_1696_v25_) else d_1694_v23_))
                        index295_ = default__.safeIndex(636, (d_1664_v0_).length(0))
                        (d_1664_v0_)[index295_] = (d_1679_v11_).fm20(globalState)
                        d_1697_v26_: D2
                        d_1697_v26_ = D2_DC6(p0)
                        d_1698_v27_: _dafny.MultiSet
                        d_1698_v27_ = _dafny.MultiSet([d_1697_v26_, d_1697_v26_, d_1697_v26_])
                        d_1699_v28_: _dafny.MultiSet
                        d_1699_v28_ = _dafny.MultiSet([(d_1698_v27_).set(d_1697_v26_, default__.abs(p0)), d_1698_v27_, d_1698_v27_, d_1698_v27_, d_1698_v27_])
                        d_1700_v29_: str
                        d_1700_v29_ = _dafny.CodePoint('f')
                        d_1701_v30_: _dafny.Map
                        d_1701_v30_ = _dafny.Map({(d_1664_v0_)[default__.safeIndex(636, (d_1664_v0_).length(0))]: d_1700_v29_})
                        d_1702_v31_: _dafny.Seq
                        d_1702_v31_ = _dafny.SeqWithoutIsStrInference([d_1698_v27_, (d_1698_v27_).set(d_1697_v26_, default__.abs(p0))])
                        d_1703_v32_: _dafny.Map
                        d_1703_v32_ = _dafny.Map({p0: _dafny.MultiSet(d_1702_v31_)})
                        d_1704_v33_: _dafny.Array
                        nw258_ = _dafny.Array(None, 9)
                        nw258_[int(0)] = d_1699_v28_
                        nw258_[int(1)] = (d_1699_v28_ if p1 else d_1699_v28_)
                        nw258_[int(2)] = (d_1699_v28_).intersection(d_1699_v28_)
                        nw258_[int(3)] = default__.fm50((d_1664_v0_)[default__.safeIndex(636, (d_1664_v0_).length(0))], (self).fm10(d_1701_v30_, 790, d_1671_v5_, globalState), p0, ((d_1683_v15_)[p0] if (p0) in (d_1683_v15_) else p1), globalState)
                        nw258_[int(4)] = ((d_1703_v32_)[p0] if (p0) in (d_1703_v32_) else d_1699_v28_)
                        nw258_[int(5)] = d_1699_v28_
                        nw258_[int(6)] = d_1699_v28_
                        nw258_[int(7)] = d_1699_v28_
                        nw258_[int(8)] = (default__.fm50((d_1664_v0_)[default__.safeIndex(636, (d_1664_v0_).length(0))], (self).fm9(p0, globalState), p0, p1, globalState)) | (d_1699_v28_)
                        d_1704_v33_ = nw258_
                        index296_ = default__.safeIndex(732, (d_1704_v33_).length(0))
                        (d_1704_v33_)[index296_] = d_1699_v28_
                        index297_ = default__.safeIndex(732, (d_1704_v33_).length(0))
                        (d_1704_v33_)[index297_] = d_1699_v28_
                    elif True:
                        d_1705_v34_: bool
                        d_1705_v34_ = True
                        d_1706_v35_: _dafny.MultiSet
                        d_1706_v35_ = _dafny.MultiSet([p0, 842, (0) - (p0)])
                        d_1707_v36_: _dafny.Seq
                        d_1707_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                        rhs350_ = (0) - (p0)
                        rhs351_ = ((d_1707_v36_ if p1 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1708_i2_ in range(default__.abs(506))]))) <= (d_1707_v36_)
                        rhs352_ = ((d_1706_v35_).intersection(d_1706_v35_)) | (d_1706_v35_)
                        lhs248_ = globalState
                        lhs248_.f13 = rhs350_
                        d_1705_v34_ = rhs351_
                        d_1706_v35_ = rhs352_
                        d_1709_v37_: _dafny.Set
                        d_1709_v37_ = _dafny.Set({p0, p0, p0, p0, p0})
                        d_1710_v38_: _dafny.MultiSet
                        d_1710_v38_ = _dafny.MultiSet([d_1709_v37_])
                        d_1711_v39_: D5
                        d_1711_v39_ = D5_DC13(d_1710_v38_)
                        d_1712_v40_: _dafny.MultiSet
                        d_1712_v40_ = _dafny.MultiSet([d_1711_v39_])
                        (globalState).f4 = (d_1712_v40_).cardinality
                        index298_ = default__.safeIndex(84, (d_1664_v0_).length(0))
                        (d_1664_v0_)[index298_] = p1
                        d_1713_v41_: _dafny.Map
                        d_1713_v41_ = _dafny.Map({not(p1): p1})
                        d_1714_v42_: _dafny.Map
                        d_1714_v42_ = _dafny.Map({d_1705_v34_: d_1713_v41_})
                        index299_ = default__.safeIndex(84, (d_1664_v0_).length(0))
                        rhs353_ = d_1664_v0_
                        rhs354_ = (0) - (default__.safeModuloInt((p0) * ((0) - (p0)), default__.fm17(p0, globalState)))
                        rhs355_ = (self).fm1(d_1714_v42_, p0, globalState)
                        lhs249_ = globalState
                        lhs250_ = d_1664_v0_
                        lhs251_ = default__.safeIndex(84, (d_1664_v0_).length(0))
                        d_1664_v0_ = rhs353_
                        lhs249_.f13 = rhs354_
                        lhs250_[lhs251_] = rhs355_
                        d_1705_v34_ = (d_1709_v37_).isdisjoint((d_1709_v37_ if False else _dafny.Set({p0})))
                        d_1715_v43_: _dafny.Map
                        d_1715_v43_ = _dafny.Map({670: d_1707_v36_})
                        d_1716_v44_: _dafny.Map
                        d_1716_v44_ = _dafny.Map({False: (d_1715_v43_) | (_dafny.Map({p0: d_1707_v36_}))})
                        def iife132_():
                            coll76_ = _dafny.Map()
                            compr_76_: int
                            for compr_76_ in (d_1706_v35_).Elements:
                                d_1717_v45_: int = compr_76_
                                if (d_1717_v45_) in (d_1706_v35_):
                                    coll76_[(d_1717_v45_) - (p0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1718_i3_ in range(default__.abs(124))])
                            return _dafny.Map(coll76_)
                        d_1716_v44_ = (d_1716_v44_).set((d_1664_v0_)[default__.safeIndex(84, (d_1664_v0_).length(0))], iife132_()
                        )
                    (globalState).f15 = default__.safeModuloInt(p0, p0)
                    (globalState).f15 = default__.fm17(p0, globalState)
                    d_1719_v46_: bool
                    d_1719_v46_ = True
                    d_1720_v47_: _dafny.Seq
                    d_1720_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfrlx"))
                    d_1721_v48_: _dafny.MultiSet
                    d_1721_v48_ = _dafny.MultiSet([p1])
                    rhs356_ = d_1719_v46_
                    rhs357_ = p1
                    rhs358_ = d_1720_v47_
                    rhs359_ = d_1721_v48_
                    d_1719_v46_ = rhs356_
                    d_1719_v46_ = rhs357_
                    d_1720_v47_ = rhs358_
                    d_1721_v48_ = rhs359_
                    pass
            pass
        r0 = (0) - ((0) - (p0))
        d_1722_v49_: _dafny.Set
        d_1722_v49_ = _dafny.Set({p1, p1})
        d_1723_v50_: _dafny.Set
        d_1723_v50_ = _dafny.Set({len(d_1722_v49_)})
        d_1724_v51_: D11
        d_1724_v51_ = D11_DC28(d_1723_v50_)
        d_1725_v52_: _dafny.Seq
        d_1725_v52_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "holgkicqv"))), p0])
        d_1726_v53_: _dafny.Seq
        d_1726_v53_ = _dafny.SeqWithoutIsStrInference([default__.fm51(d_1725_v52_, p0, globalState)])
        d_1727_v54_: D11
        d_1727_v54_ = D11_DC31(D11_DC29())
        d_1728_v55_: D19
        d_1728_v55_ = D19_DC51(d_1726_v53_)
        d_1729_v56_: _dafny.Seq
        d_1729_v56_ = _dafny.SeqWithoutIsStrInference([d_1726_v53_, d_1726_v53_])
        pat_let_tv39_ = d_1724_v51_
        d_1730_v57_: _dafny.Array
        nw259_ = _dafny.Array(None, 21)
        nw259_[int(0)] = _dafny.SeqWithoutIsStrInference([D11_DC31(d_1724_v51_)])
        nw259_[int(1)] = d_1726_v53_
        nw259_[int(2)] = (_dafny.SeqWithoutIsStrInference([D11_DC31(D11_DC31(d_1724_v51_)), d_1727_v54_, d_1727_v54_])) + (d_1726_v53_)
        nw259_[int(3)] = d_1726_v53_
        nw259_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1727_v54_])
        def iife133_(_pat_let28_0):
            def iife134_(d_1731_dt__update__tmp_h2_):
                def iife135_(_pat_let29_0):
                    def iife136_(d_1732_dt__update_hcf48_h0_):
                        return D11_DC31(d_1732_dt__update_hcf48_h0_)
                    return iife136_(_pat_let29_0)
                return iife135_(pat_let_tv39_)
            return iife134_(_pat_let28_0)
        nw259_[int(5)] = _dafny.SeqWithoutIsStrInference([iife133_(d_1727_v54_)])
        nw259_[int(6)] = d_1726_v53_
        nw259_[int(7)] = (d_1726_v53_ if p1 else (d_1728_v55_).cf80)
        nw259_[int(8)] = (_dafny.SeqWithoutIsStrInference([d_1727_v54_])) + (d_1726_v53_)
        nw259_[int(9)] = (d_1729_v56_)[default__.safeIndex(p0, len(d_1729_v56_))]
        nw259_[int(10)] = d_1726_v53_
        nw259_[int(11)] = (default__.fm52(globalState)) + (d_1726_v53_)
        nw259_[int(12)] = d_1726_v53_
        nw259_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1727_v54_ for d_1733_i4_ in range(default__.abs(768))])
        nw259_[int(14)] = d_1726_v53_
        nw259_[int(15)] = d_1726_v53_
        nw259_[int(16)] = d_1726_v53_
        nw259_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1727_v54_ for d_1734_i5_ in range(default__.abs(-243))])
        nw259_[int(18)] = d_1726_v53_
        nw259_[int(19)] = d_1726_v53_
        nw259_[int(20)] = (d_1726_v53_).set(default__.safeIndex(p0, len(d_1726_v53_)), default__.fm51(d_1725_v52_, p0, globalState))
        d_1730_v57_ = nw259_
        index300_ = default__.safeIndex(537, (d_1730_v57_).length(0))
        (d_1730_v57_)[index300_] = _dafny.SeqWithoutIsStrInference([D11_DC31(d_1724_v51_), d_1727_v54_])
        d_1735_v58_: bool
        d_1735_v58_ = True
        d_1736_v59_: D0
        d_1736_v59_ = D0_DC0(d_1664_v0_)
        d_1737_v60_: _dafny.Map
        d_1737_v60_ = _dafny.Map({p0: False})
        d_1738_v61_: _dafny.Seq
        d_1738_v61_ = _dafny.SeqWithoutIsStrInference([d_1735_v58_, ((d_1737_v60_)[len((d_1725_v52_).set(default__.safeIndex(667, len(d_1725_v52_)), p0))] if (len((d_1725_v52_).set(default__.safeIndex(667, len(d_1725_v52_)), p0))) in (d_1737_v60_) else d_1735_v58_)])
        d_1739_v62_: _dafny.Seq
        d_1739_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksdbdq"))
        d_1740_v63_: _dafny.Seq
        d_1740_v63_ = _dafny.SeqWithoutIsStrInference([d_1739_v62_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhqnf")), default__.fm30(p0, d_1735_v58_, d_1723_v50_, True, globalState)])
        d_1741_v64_: D7
        d_1741_v64_ = D7_DC21(p0, p0, False, d_1735_v58_, p1)
        index301_ = default__.safeIndex(537, (d_1730_v57_).length(0))
        rhs360_ = ((d_1736_v59_ if d_1735_v58_ else d_1736_v59_)).cf0
        rhs361_ = (default__.fm52(globalState)).set(default__.safeIndex((len(d_1738_v61_)) + (-182), len(default__.fm52(globalState))), d_1727_v54_)
        rhs362_ = len((d_1740_v63_)[default__.safeIndex((p0) + (p0), len(d_1740_v63_))])
        rhs363_ = (((0) - ((d_1741_v64_).cf30)) + (-837)) > ((p0) - (p0))
        lhs252_ = d_1730_v57_
        lhs253_ = default__.safeIndex(537, (d_1730_v57_).length(0))
        lhs254_ = globalState
        d_1664_v0_ = rhs360_
        lhs252_[lhs253_] = rhs361_
        lhs254_.f4 = rhs362_
        d_1735_v58_ = rhs363_
        d_1742_i6_: int
        d_1742_i6_ = 0
        with _dafny.label("11"):
            while p1:
                with _dafny.c_label("11"):
                    if (d_1742_i6_) >= (100):
                        raise _dafny.Break("11")
                    d_1742_i6_ = (d_1742_i6_) + (1)
                    index302_ = default__.safeIndex(332, (d_1664_v0_).length(0))
                    (d_1664_v0_)[index302_] = False
                    index303_ = default__.safeIndex(825, (d_1664_v0_).length(0))
                    (d_1664_v0_)[index303_] = not((p0) <= (p0))
                    d_1743_v65_: str
                    d_1743_v65_ = _dafny.CodePoint('l')
                    d_1744_v66_: _dafny.Map
                    d_1744_v66_ = _dafny.Map({False: p1})
                    d_1745_v67_: _dafny.Map
                    d_1745_v67_ = _dafny.Map({d_1735_v58_: d_1743_v65_})
                    d_1746_v68_: _dafny.Map
                    d_1746_v68_ = _dafny.Map({p0: len(d_1739_v62_)})
                    d_1747_v69_: _dafny.Map
                    d_1747_v69_ = _dafny.Map({((d_1744_v66_)[p1] if (p1) in (d_1744_v66_) else p1): (self).fm10(d_1745_v67_, p0, d_1746_v68_, globalState)})
                    d_1748_v70_: _dafny.Map
                    d_1748_v70_ = _dafny.Map({p1: d_1747_v69_})
                    index304_ = default__.safeIndex(332, (d_1664_v0_).length(0))
                    index305_ = default__.safeIndex(825, (d_1664_v0_).length(0))
                    rhs364_ = (p0) == (p0)
                    rhs365_ = ((_dafny.SeqWithoutIsStrInference([d_1735_v58_, p1])) + (d_1738_v61_)) + (d_1738_v61_)
                    rhs366_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqqe"))) == (((d_1739_v62_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "behwcy")))).set(default__.safeIndex((0) - (p0), len((d_1739_v62_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "behwcy"))))), d_1743_v65_))
                    rhs367_ = (self).fm1(d_1748_v70_, p0, globalState)
                    lhs255_ = d_1664_v0_
                    lhs256_ = default__.safeIndex(332, (d_1664_v0_).length(0))
                    lhs257_ = d_1664_v0_
                    lhs258_ = default__.safeIndex(825, (d_1664_v0_).length(0))
                    d_1735_v58_ = rhs364_
                    d_1738_v61_ = rhs365_
                    lhs255_[lhs256_] = rhs366_
                    lhs257_[lhs258_] = rhs367_
                    (globalState).f4 = (0) - (default__.safeModuloInt(p0, 488))
                    d_1749_v71_: D4
                    d_1749_v71_ = D4_DC12(((0) - (len(d_1739_v62_)) if p1 else p0))
                    source25_ = d_1749_v71_
                    if source25_.is_DC12:
                        d_1750___mcc_h0_ = source25_.cf16
                        d_1751_cf16_ = d_1750___mcc_h0_
                        d_1752_v72_: C2
                        nw260_ = C2()
                        nw260_.ctor__()
                        d_1752_v72_ = nw260_
                        d_1753_v73_: _dafny.Array
                        def lambda89_(d_1754_cf16_):
                            def lambda90_(d_1755_i7_):
                                return default__.safeModuloInt(d_1755_i7_, d_1754_cf16_)

                            return lambda90_

                        init46_ = lambda89_(d_1751_cf16_)
                        nw261_ = _dafny.Array(None, 19)
                        for i0_46_ in range(nw261_.length(0)):
                            nw261_[i0_46_] = init46_(i0_46_)
                        d_1753_v73_ = nw261_
                        index306_ = default__.safeIndex(409, (d_1753_v73_).length(0))
                        (d_1753_v73_)[index306_] = p0
                        d_1756_v74_: _dafny.Array
                        nw262_ = _dafny.Array(False, 27)
                        d_1756_v74_ = nw262_
                        d_1757_v75_: _dafny.Map
                        d_1757_v75_ = _dafny.Map({d_1756_v74_: (d_1735_v58_ if p1 else (self).fm8(634, p1, p0, globalState))})
                        index307_ = default__.safeIndex(409, (d_1753_v73_).length(0))
                        (d_1753_v73_)[index307_] = len(d_1757_v75_)
                        d_1758_v76_: _dafny.Array
                        nw263_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_1758_v76_ = nw263_
                        d_1759_v77_: _dafny.Map
                        d_1759_v77_ = _dafny.Map({d_1735_v58_: d_1758_v76_})
                        d_1760_v78_: _dafny.MultiSet
                        d_1760_v78_ = _dafny.MultiSet([d_1751_cf16_])
                        d_1759_v77_ = (d_1759_v77_).set((d_1760_v78_).ispropersubset((d_1760_v78_).set(p0, default__.abs(d_1751_cf16_))), d_1758_v76_)
                        d_1761_v80_: D15
                        d_1761_v80_ = D15_DC38(_dafny.Map({d_1751_cf16_: d_1735_v58_}))
                        d_1762_v81_: _dafny.MultiSet
                        def iife137_():
                            coll77_ = _dafny.Map()
                            compr_77_: int
                            for compr_77_ in (_dafny.SeqWithoutIsStrInference([(d_1753_v73_)[default__.safeIndex(409, (d_1753_v73_).length(0))] for d_1763_i8_ in range(default__.abs(-584))])).Elements:
                                d_1764_v79_: int = compr_77_
                                if (d_1764_v79_) in (_dafny.SeqWithoutIsStrInference([(d_1753_v73_)[default__.safeIndex(409, (d_1753_v73_).length(0))] for d_1763_i8_ in range(default__.abs(-584))])):
                                    coll77_[(d_1764_v79_) + (len(_dafny.Map({p1: False})))] = p1
                            return _dafny.Map(coll77_)
                        d_1762_v81_ = _dafny.MultiSet([D15_DC38(d_1737_v60_), D15_DC38(iife137_()
), d_1761_v80_])
                        index308_ = default__.safeIndex(409, (d_1753_v73_).length(0))
                        (d_1753_v73_)[index308_] = ((d_1762_v81_)[d_1761_v80_] if (d_1761_v80_) in (d_1762_v81_) else len((_dafny.Set({(d_1664_v0_)[default__.safeIndex(332, (d_1664_v0_).length(0))], (d_1664_v0_)[default__.safeIndex(332, (d_1664_v0_).length(0))]})) | (d_1722_v49_)))
                    elif True:
                        d_1765___mcc_h1_ = source25_.cf15
                        d_1766_cf15_ = d_1765___mcc_h1_
                        d_1735_v58_ = not(not(True))
                        d_1767_v82_: D3
                        d_1767_v82_ = D3_DC10(p0, p0)
                        d_1767_v82_ = default__.fm53((self).fm2(p0, globalState), globalState)
                        d_1735_v58_ = d_1735_v58_
                        d_1768_v83_: _dafny.Map
                        d_1768_v83_ = _dafny.Map({(self).fm9(p0, globalState): (p0) + (p0)})
                        d_1769_v84_: _dafny.Array
                        nw264_ = _dafny.Array(None, 21)
                        nw264_[int(0)] = (0) - (p0)
                        nw264_[int(1)] = p0
                        nw264_[int(2)] = (self).fm2(len(d_1737_v60_), globalState)
                        nw264_[int(3)] = p0
                        nw264_[int(4)] = p0
                        nw264_[int(5)] = (_dafny.MultiSet(d_1725_v52_)).cardinality
                        nw264_[int(6)] = p0
                        nw264_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcuryoktk")))
                        nw264_[int(8)] = len(_dafny.Map({len(d_1739_v62_): d_1735_v58_}))
                        nw264_[int(9)] = p0
                        nw264_[int(10)] = p0
                        nw264_[int(11)] = p0
                        nw264_[int(12)] = p0
                        nw264_[int(13)] = p0
                        nw264_[int(14)] = p0
                        nw264_[int(15)] = len((d_1739_v62_).set(default__.safeIndex(len(d_1747_v69_), len(d_1739_v62_)), d_1743_v65_))
                        nw264_[int(16)] = len(d_1739_v62_)
                        nw264_[int(17)] = p0
                        nw264_[int(18)] = p0
                        nw264_[int(19)] = p0
                        nw264_[int(20)] = (d_1725_v52_)[default__.safeIndex(364, len(d_1725_v52_))]
                        d_1769_v84_ = nw264_
                        d_1770_v85_: _dafny.Map
                        d_1770_v85_ = _dafny.Map({d_1769_v84_: d_1725_v52_})
                        d_1771_v86_: _dafny.Seq
                        d_1771_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_1772_i9_ in range(default__.abs(277))])])
                        d_1773_v87_: _dafny.Map
                        d_1773_v87_ = _dafny.Map({len(d_1725_v52_): d_1738_v61_})
                        d_1774_v88_: D1
                        d_1774_v88_ = D1_DC3(d_1739_v62_, p1)
                        d_1768_v83_ = _dafny.Map({(((d_1770_v85_)[d_1769_v84_] if (d_1769_v84_) in (d_1770_v85_) else (d_1771_v86_)[default__.safeIndex((0) - (len(((d_1773_v87_)[p0] if (p0) in (d_1773_v87_) else d_1738_v61_))), len(d_1771_v86_))])) < (d_1725_v52_): len(((d_1774_v88_).cf3) + (d_1739_v62_))})
                    (globalState).f7 = ((0) - (default__.safeModuloInt(p0, 766))) * (p0)
                    pass
            pass
        d_1775_v89_: C0
        nw265_ = C0()
        nw265_.ctor__()
        d_1775_v89_ = nw265_
        r0 = (self).fm2((0) - (p0), globalState)
        d_1776_v90_: _dafny.Set
        d_1776_v90_ = _dafny.Set({D0_DC0(d_1664_v0_)})
        d_1777_v91_: _dafny.MultiSet
        d_1777_v91_ = _dafny.MultiSet([d_1776_v90_])
        r1 = ((d_1777_v91_) | (d_1777_v91_)) | (d_1777_v91_)
        return r0, r1


class C10(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, globalState):
        return True

    def fm2(self, p0, globalState):
        return (0) - (-601)

    def fm6(self, p0, p1, p2, p3, globalState):
        def lambda91_(source26_):
            if source26_.is_DC3:
                d_1778___mcc_h0_ = source26_.cf3
                d_1779___mcc_h1_ = source26_.cf4
                d_1780_cf4_ = d_1779___mcc_h1_
                d_1781_cf3_ = d_1778___mcc_h0_
                return (D1_DC2(len(_dafny.Map({d_1780_cf4_: (0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1780_cf4_, d_1780_cf4_])).cardinality, (0) - ((_dafny.MultiSet([621, len(_dafny.Map({d_1780_cf4_: d_1780_cf4_})), (_dafny.MultiSet([60])).cardinality, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([107 for d_1782_i0_ in range(default__.abs(919))])): d_1780_cf4_})), 262])).cardinality)])))})))).cf2
            elif source26_.is_DC4:
                d_1783___mcc_h2_ = source26_.cf5
                d_1784_cf5_ = d_1783___mcc_h2_
                return (0) - (len((_dafny.SeqWithoutIsStrInference([False, not(True), True, not(not(False))])) + (_dafny.SeqWithoutIsStrInference([True, not(False)]))))
            elif True:
                d_1785___mcc_h3_ = source26_.cf2
                d_1786_cf2_ = d_1785___mcc_h3_
                return d_1786_cf2_

        return (0) - (lambda91_(D1_DC2(647)))

    def m1(self, p0, p1, globalState):
        d_1787_v0_: _dafny.Array
        nw266_ = _dafny.Array(int(0), 6)
        d_1787_v0_ = nw266_
        d_1788_v1_: _dafny.Seq
        d_1788_v1_ = _dafny.SeqWithoutIsStrInference([d_1787_v0_, d_1787_v0_])
        d_1789_v2_: int
        d_1789_v2_ = 216
        d_1787_v0_ = (d_1788_v1_)[default__.safeIndex(d_1789_v2_, len(d_1788_v1_))]
        (globalState).f4 = (d_1789_v2_) - (d_1789_v2_)
        d_1790_v3_: bool
        d_1790_v3_ = True
        d_1791_v4_: _dafny.Map
        d_1791_v4_ = _dafny.Map({d_1790_v3_: d_1790_v3_})
        d_1790_v3_ = (self).fm1(_dafny.Map({d_1790_v3_: d_1791_v4_}), len(d_1791_v4_), globalState)
        d_1792_v5_: _dafny.MultiSet
        d_1792_v5_ = _dafny.MultiSet([d_1790_v3_])
        d_1793_v6_: D2
        d_1793_v6_ = D2_DC7(True, d_1792_v5_, d_1787_v0_)
        pat_let_tv40_ = d_1787_v0_
        d_1794_v7_: _dafny.Array
        nw267_ = _dafny.Array(None, 9)
        nw267_[int(0)] = d_1787_v0_
        nw267_[int(1)] = d_1787_v0_
        nw267_[int(2)] = d_1787_v0_
        nw267_[int(3)] = d_1787_v0_
        nw267_[int(4)] = d_1787_v0_
        nw267_[int(5)] = d_1787_v0_
        nw267_[int(6)] = d_1787_v0_
        def iife138_(_pat_let30_0):
            def iife139_(d_1795_dt__update__tmp_h0_):
                def iife140_(_pat_let31_0):
                    def iife141_(d_1796_dt__update_hcf10_h0_):
                        def iife142_(_pat_let32_0):
                            def iife143_(d_1797_dt__update_hcf8_h0_):
                                return D2_DC7(d_1797_dt__update_hcf8_h0_, (d_1795_dt__update__tmp_h0_).cf9, d_1796_dt__update_hcf10_h0_)
                            return iife143_(_pat_let32_0)
                        return iife142_(False)
                    return iife141_(_pat_let31_0)
                return iife140_(pat_let_tv40_)
            return iife139_(_pat_let30_0)
        nw267_[int(7)] = (iife138_(d_1793_v6_)).cf10
        nw267_[int(8)] = d_1787_v0_
        d_1794_v7_ = nw267_
        nw268_ = _dafny.Array(_dafny.Array(None, 0), 21)
        d_1794_v7_ = nw268_
        hi13_ = default__.safeDivisionInt((d_1792_v5_).cardinality, (0) - (d_1789_v2_))
        for d_1798_i0_ in range(d_1789_v2_, hi13_):
            d_1799_v8_: _dafny.Array
            nw269_ = _dafny.Array(False, 17)
            d_1799_v8_ = nw269_
            d_1799_v8_ = d_1799_v8_
            d_1800_v9_: str
            d_1800_v9_ = _dafny.CodePoint('j')
            d_1800_v9_ = d_1800_v9_
            d_1801_v10_: _dafny.Map
            d_1801_v10_ = _dafny.Map({d_1790_v3_: d_1791_v4_})
            d_1802_v11_: _dafny.Map
            d_1802_v11_ = _dafny.Map({d_1789_v2_: (0) - ((d_1792_v5_).cardinality)})
            d_1803_v12_: _dafny.Map
            d_1803_v12_ = _dafny.Map({((self).fm6(False, (self).fm1(d_1801_v10_, d_1789_v2_, globalState), d_1802_v11_, _dafny.CodePoint('w'), globalState)) - (446): (d_1792_v5_).ispropersubset(_dafny.MultiSet([d_1790_v3_]))})
            d_1803_v12_ = (d_1803_v12_).set((0) - (d_1798_i0_), d_1790_v3_)
            d_1804_v13_: C4
            nw270_ = C4()
            nw270_.ctor__()
            d_1804_v13_ = nw270_
        d_1805_v14_: str
        d_1805_v14_ = _dafny.CodePoint('u')
        d_1805_v14_ = d_1805_v14_

    def m5(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: str = _dafny.CodePoint('D')
        d_1806_v0_: _dafny.Set
        d_1806_v0_ = _dafny.Set({p0, p0, p0})
        hi14_ = p0
        for d_1807_i0_ in range(len(d_1806_v0_), hi14_):
            r1 = not (True) or (p2)
            r1 = (p0) > (d_1807_i0_)
            d_1808_v1_: _dafny.MultiSet
            d_1808_v1_ = _dafny.MultiSet([d_1807_i0_, d_1807_i0_])
            d_1809_v2_: _dafny.Seq
            d_1809_v2_ = _dafny.SeqWithoutIsStrInference([p0])
            r1 = (d_1808_v1_).issubset((d_1808_v1_) | (_dafny.MultiSet(d_1809_v2_)))
            d_1810_v3_: _dafny.Array
            def lambda92_(d_1811_i0_):
                def lambda93_(d_1812_i1_):
                    return (d_1812_i1_) + (d_1811_i0_)

                return lambda93_

            init47_ = lambda92_(d_1807_i0_)
            nw271_ = _dafny.Array(None, 26)
            for i0_47_ in range(nw271_.length(0)):
                nw271_[i0_47_] = init47_(i0_47_)
            d_1810_v3_ = nw271_
            d_1813_v4_: _dafny.MultiSet
            d_1813_v4_ = _dafny.MultiSet([False])
            index309_ = default__.safeIndex(302, (d_1810_v3_).length(0))
            (d_1810_v3_)[index309_] = ((d_1808_v1_)[(d_1813_v4_).cardinality] if ((d_1813_v4_).cardinality) in (d_1808_v1_) else p0)
            d_1814_v5_: _dafny.Map
            d_1814_v5_ = _dafny.Map({p1: p0})
            index310_ = default__.safeIndex(302, (d_1810_v3_).length(0))
            (d_1810_v3_)[index310_] = (0) - (((d_1807_i0_) + (p0)) * (((d_1814_v5_)[False] if (False) in (d_1814_v5_) else (0) - (-617))))
        d_1815_v6_: _dafny.Map
        d_1815_v6_ = _dafny.Map({p2: False})
        d_1816_v7_: _dafny.Map
        d_1816_v7_ = _dafny.Map({p2: d_1815_v6_})
        d_1817_i2_: int
        d_1817_i2_ = 0
        with _dafny.label("12"):
            while (self).fm1(d_1816_v7_, p0, globalState):
                with _dafny.c_label("12"):
                    if (d_1817_i2_) >= (100):
                        raise _dafny.Break("12")
                    d_1817_i2_ = (d_1817_i2_) + (1)
                    d_1818_v8_: _dafny.Seq
                    d_1818_v8_ = _dafny.SeqWithoutIsStrInference([p1])
                    r1 = (((d_1818_v8_).set(default__.safeIndex(p0, len(d_1818_v8_)), p2)) + (d_1818_v8_)) < (d_1818_v8_)
                    d_1819_v9_: str
                    d_1819_v9_ = _dafny.CodePoint('k')
                    r2 = d_1819_v9_
                    r1 = (p2 if False else p1)
                    (globalState).f4 = p0
                    pass
            pass
        d_1820_v10_: _dafny.Map
        d_1820_v10_ = _dafny.Map({p0: p0})
        d_1821_v11_: _dafny.Seq
        d_1821_v11_ = _dafny.SeqWithoutIsStrInference([not(not(p2)), p1])
        d_1822_v12_: D12
        d_1822_v12_ = D12_DC33(p1)
        d_1823_v13_: str
        d_1823_v13_ = _dafny.CodePoint('k')
        d_1824_v14_: D11
        d_1824_v14_ = D11_DC30(p1, p0)
        d_1825_v15_: _dafny.Seq
        d_1825_v15_ = _dafny.SeqWithoutIsStrInference([620, p0])
        d_1826_v16_: _dafny.MultiSet
        d_1826_v16_ = _dafny.MultiSet([not(p2), p1])
        d_1827_v17_: C5
        nw272_ = C5()
        nw272_.ctor__()
        d_1827_v17_ = nw272_
        d_1828_v18_: _dafny.Array
        nw273_ = _dafny.Array(None, 27)
        nw273_[int(0)] = p1
        nw273_[int(1)] = p2
        nw273_[int(2)] = (self).fm1(d_1816_v7_, len(d_1820_v10_), globalState)
        nw273_[int(3)] = p1
        nw273_[int(4)] = True
        nw273_[int(5)] = not(p1)
        nw273_[int(6)] = (p0) < (p0)
        nw273_[int(7)] = p1
        nw273_[int(8)] = (self).fm1(d_1816_v7_, p0, globalState)
        nw273_[int(9)] = p1
        nw273_[int(10)] = (d_1821_v11_) < (d_1821_v11_)
        nw273_[int(11)] = not(p2)
        nw273_[int(12)] = (p0) <= (p0)
        nw273_[int(13)] = (p1 if p2 else (self).fm1(d_1816_v7_, p0, globalState))
        nw273_[int(14)] = (d_1822_v12_).cf50
        nw273_[int(15)] = ((self).fm6(p2, p1, d_1820_v10_, d_1823_v13_, globalState)) == ((self).fm6(p2, p1, _dafny.Map({(d_1824_v14_).cf47: p0}), d_1823_v13_, globalState))
        nw273_[int(16)] = p2
        nw273_[int(17)] = p2
        nw273_[int(18)] = p1
        nw273_[int(19)] = (d_1825_v15_) < (d_1825_v15_)
        nw273_[int(20)] = p1
        nw273_[int(21)] = p1
        nw273_[int(22)] = (d_1826_v16_).ispropersubset(d_1826_v16_)
        nw273_[int(23)] = p2
        nw273_[int(24)] = False
        nw273_[int(25)] = (not(p1)) not in (_dafny.Map({p2: d_1827_v17_}))
        nw273_[int(26)] = p1
        d_1828_v18_ = nw273_
        d_1828_v18_ = d_1828_v18_
        d_1829_v19_: _dafny.Seq
        d_1829_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpee"))
        d_1829_v19_ = (default__.fm19(p2, p1, globalState)).set(default__.safeIndex(p0, len(default__.fm19(p2, p1, globalState))), d_1823_v13_)
        if False:
            d_1830_v20_: _dafny.Array
            nw274_ = _dafny.Array(D19.default()(), 23)
            d_1830_v20_ = nw274_
            d_1830_v20_ = d_1830_v20_
            (globalState).f15 = p0
            d_1829_v19_ = default__.fm30((p0) - (p0), (p0) != (p0), d_1806_v0_, (p1 if True else p1), globalState)
            if p2:
                r1 = p1
                d_1831_v21_: _dafny.Array
                nw275_ = _dafny.Array(int(0), 4)
                d_1831_v21_ = nw275_
                index311_ = default__.safeIndex(469, (d_1831_v21_).length(0))
                (d_1831_v21_)[index311_] = (0) - (p0)
                index312_ = default__.safeIndex(469, (d_1831_v21_).length(0))
                (d_1831_v21_)[index312_] = p0
                d_1832_v22_: _dafny.Map
                d_1832_v22_ = _dafny.Map({(d_1828_v18_ if p2 else d_1828_v18_): default__.fm17(p0, globalState)})
                (globalState).f15 = ((d_1832_v22_)[d_1828_v18_] if (d_1828_v18_) in (d_1832_v22_) else (0) - ((d_1831_v21_)[default__.safeIndex(469, (d_1831_v21_).length(0))]))
                d_1833_v23_: _dafny.Array
                nw276_ = _dafny.Array(None, 17)
                nw276_[int(0)] = d_1831_v21_
                nw276_[int(1)] = d_1831_v21_
                nw276_[int(2)] = d_1831_v21_
                nw276_[int(3)] = d_1831_v21_
                nw276_[int(4)] = d_1831_v21_
                nw276_[int(5)] = d_1831_v21_
                nw276_[int(6)] = d_1831_v21_
                nw276_[int(7)] = d_1831_v21_
                nw276_[int(8)] = d_1831_v21_
                nw276_[int(9)] = d_1831_v21_
                nw276_[int(10)] = d_1831_v21_
                nw276_[int(11)] = d_1831_v21_
                nw276_[int(12)] = d_1831_v21_
                nw276_[int(13)] = d_1831_v21_
                nw276_[int(14)] = d_1831_v21_
                nw276_[int(15)] = d_1831_v21_
                nw276_[int(16)] = d_1831_v21_
                d_1833_v23_ = nw276_
                d_1834_v24_: D20
                d_1834_v24_ = D20_DC54(d_1833_v23_)
                d_1835_v25_: _dafny.Array
                nw277_ = _dafny.Array(None, 4)
                nw277_[int(0)] = d_1833_v23_
                nw277_[int(1)] = d_1833_v23_
                nw277_[int(2)] = d_1833_v23_
                nw277_[int(3)] = (d_1834_v24_).cf84
                d_1835_v25_ = nw277_
                index313_ = default__.safeIndex(673, (d_1835_v25_).length(0))
                (d_1835_v25_)[index313_] = d_1833_v23_
                index314_ = default__.safeIndex(673, (d_1835_v25_).length(0))
                (d_1835_v25_)[index314_] = d_1833_v23_
                d_1836_v26_: _dafny.Set
                d_1836_v26_ = _dafny.Set({p1})
                rhs368_ = d_1831_v21_
                rhs369_ = (d_1831_v21_)[default__.safeIndex(469, (d_1831_v21_).length(0))]
                rhs370_ = (d_1836_v26_).ispropersubset(d_1836_v26_)
                rhs371_ = (d_1829_v19_) == (default__.fm19(False, p1, globalState))
                lhs259_ = globalState
                d_1831_v21_ = rhs368_
                lhs259_.f13 = rhs369_
                r1 = rhs370_
                r1 = rhs371_
            elif True:
                d_1837_v27_: _dafny.Array
                def lambda94_(d_1838_p1_):
                    def lambda95_(d_1839_i3_):
                        return _dafny.Map({d_1838_p1_: 789})

                    return lambda95_

                init48_ = lambda94_(p1)
                nw278_ = _dafny.Array(None, 18)
                for i0_48_ in range(nw278_.length(0)):
                    nw278_[i0_48_] = init48_(i0_48_)
                d_1837_v27_ = nw278_
                index315_ = default__.safeIndex(593, (d_1837_v27_).length(0))
                (d_1837_v27_)[index315_] = (_dafny.Map({p1: p0})) | (_dafny.Map({p2: len(d_1825_v15_)}))
                d_1840_v28_: _dafny.Map
                d_1840_v28_ = _dafny.Map({p2: p0})
                index316_ = default__.safeIndex(593, (d_1837_v27_).length(0))
                rhs372_ = p1
                rhs373_ = ((d_1840_v28_ if p1 else _dafny.Map({p1: p0}))) | (d_1840_v28_)
                rhs374_ = p0
                rhs375_ = d_1823_v13_
                lhs260_ = d_1837_v27_
                lhs261_ = default__.safeIndex(593, (d_1837_v27_).length(0))
                lhs262_ = globalState
                r1 = rhs372_
                lhs260_[lhs261_] = rhs373_
                lhs262_.f7 = rhs374_
                d_1823_v13_ = rhs375_
                d_1841_v29_: _dafny.Array
                nw279_ = _dafny.Array(int(0), 18)
                d_1841_v29_ = nw279_
                index317_ = default__.safeIndex(591, (d_1841_v29_).length(0))
                (d_1841_v29_)[index317_] = p0
                index318_ = default__.safeIndex(591, (d_1841_v29_).length(0))
                (d_1841_v29_)[index318_] = (d_1825_v15_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "booqtq"))), len(d_1825_v15_))]
                d_1842_v30_: C8
                nw280_ = C8()
                nw280_.ctor__()
                d_1842_v30_ = nw280_
                r1 = not((p2) or (p2))
                d_1843_v31_: _dafny.Seq
                d_1843_v31_ = _dafny.SeqWithoutIsStrInference([d_1828_v18_])
                d_1844_v32_: _dafny.Array
                nw281_ = _dafny.Array(None, 23)
                nw281_[int(0)] = d_1828_v18_
                nw281_[int(1)] = d_1828_v18_
                nw281_[int(2)] = d_1828_v18_
                nw281_[int(3)] = d_1828_v18_
                nw281_[int(4)] = d_1828_v18_
                nw281_[int(5)] = ((d_1843_v31_)[default__.safeIndex(len(d_1820_v10_), len(d_1843_v31_))] if p2 else d_1828_v18_)
                nw281_[int(6)] = d_1828_v18_
                nw281_[int(7)] = d_1828_v18_
                nw281_[int(8)] = d_1828_v18_
                nw281_[int(9)] = d_1828_v18_
                nw281_[int(10)] = d_1828_v18_
                nw281_[int(11)] = d_1828_v18_
                nw281_[int(12)] = d_1828_v18_
                nw281_[int(13)] = d_1828_v18_
                nw281_[int(14)] = d_1828_v18_
                nw281_[int(15)] = d_1828_v18_
                nw281_[int(16)] = d_1828_v18_
                nw281_[int(17)] = d_1828_v18_
                nw281_[int(18)] = d_1828_v18_
                nw281_[int(19)] = d_1828_v18_
                nw281_[int(20)] = (d_1828_v18_ if p2 else d_1828_v18_)
                nw281_[int(21)] = d_1828_v18_
                nw281_[int(22)] = d_1828_v18_
                d_1844_v32_ = nw281_
                index319_ = default__.safeIndex(99, (d_1844_v32_).length(0))
                (d_1844_v32_)[index319_] = d_1828_v18_
                index320_ = default__.safeIndex(99, (d_1844_v32_).length(0))
                (d_1844_v32_)[index320_] = d_1828_v18_
            if p1:
                r1 = p2
                d_1845_v33_: _dafny.Array
                def lambda96_(d_1846_v0_):
                    def lambda97_(d_1847_i4_):
                        return d_1846_v0_

                    return lambda97_

                init49_ = lambda96_(d_1806_v0_)
                nw282_ = _dafny.Array(None, 9)
                for i0_49_ in range(nw282_.length(0)):
                    nw282_[i0_49_] = init49_(i0_49_)
                d_1845_v33_ = nw282_
                index321_ = default__.safeIndex(50, (d_1845_v33_).length(0))
                (d_1845_v33_)[index321_] = _dafny.Set({p0})
                index322_ = default__.safeIndex(50, (d_1845_v33_).length(0))
                (d_1845_v33_)[index322_] = d_1806_v0_
                d_1848_v34_: _dafny.MultiSet
                d_1848_v34_ = _dafny.MultiSet([p0])
                r1 = ((d_1848_v34_) - (d_1848_v34_)).issubset(d_1848_v34_)
                d_1849_v35_: _dafny.Map
                d_1849_v35_ = _dafny.Map({p2: p0})
                d_1850_v36_: _dafny.Map
                d_1850_v36_ = d_1849_v35_
                d_1851_v37_: _dafny.Map
                d_1851_v37_ = _dafny.Map({d_1849_v35_: (d_1829_v19_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1852_i5_ in range(default__.abs(976))]))})
                d_1851_v37_ = (d_1851_v37_).set(d_1850_v36_, d_1829_v19_)
                d_1853_v38_: _dafny.Set
                d_1853_v38_ = _dafny.Set({p1})
                d_1854_v39_: _dafny.Map
                d_1854_v39_ = _dafny.Map({(p0) * (p0): d_1853_v38_})
                d_1854_v39_ = _dafny.Map({p0: d_1853_v38_})
            elif True:
                d_1855_v40_: C2
                nw283_ = C2()
                nw283_.ctor__()
                d_1855_v40_ = nw283_
                d_1856_v41_: _dafny.Seq
                d_1856_v41_ = _dafny.SeqWithoutIsStrInference([d_1855_v40_])
                d_1857_v42_: _dafny.Map
                d_1857_v42_ = _dafny.Map({True: p0})
                d_1855_v40_ = (d_1856_v41_)[default__.safeIndex(((d_1857_v42_)[not(p1)] if (not(p1)) in (d_1857_v42_) else p0), len(d_1856_v41_))]
                d_1858_v43_: _dafny.Array
                nw284_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                d_1858_v43_ = nw284_
                d_1859_v44_: D21
                d_1859_v44_ = D21_DC56(d_1858_v43_)
                d_1860_v45_: _dafny.Array
                nw285_ = _dafny.Array(None, 23)
                nw285_[int(0)] = d_1858_v43_
                nw285_[int(1)] = d_1858_v43_
                nw285_[int(2)] = d_1858_v43_
                nw285_[int(3)] = d_1858_v43_
                nw285_[int(4)] = d_1858_v43_
                nw285_[int(5)] = d_1858_v43_
                nw285_[int(6)] = d_1858_v43_
                nw285_[int(7)] = (d_1858_v43_ if p2 else d_1858_v43_)
                nw285_[int(8)] = d_1858_v43_
                nw285_[int(9)] = d_1858_v43_
                nw285_[int(10)] = d_1858_v43_
                nw285_[int(11)] = d_1858_v43_
                nw285_[int(12)] = d_1858_v43_
                nw285_[int(13)] = d_1858_v43_
                nw285_[int(14)] = d_1858_v43_
                nw285_[int(15)] = d_1858_v43_
                nw285_[int(16)] = d_1858_v43_
                nw285_[int(17)] = d_1858_v43_
                nw285_[int(18)] = (d_1859_v44_).cf89
                nw285_[int(19)] = d_1858_v43_
                nw285_[int(20)] = d_1858_v43_
                nw285_[int(21)] = d_1858_v43_
                nw285_[int(22)] = d_1858_v43_
                d_1860_v45_ = nw285_
                index323_ = default__.safeIndex(182, (d_1860_v45_).length(0))
                (d_1860_v45_)[index323_] = d_1858_v43_
                index324_ = default__.safeIndex(182, (d_1860_v45_).length(0))
                (d_1860_v45_)[index324_] = d_1858_v43_
                d_1861_v46_: D5
                d_1861_v46_ = D5_DC14(d_1823_v13_, not(p2), p0, p1)
                d_1862_v47_: D5
                d_1862_v47_ = D5_DC15(d_1861_v46_)
                d_1862_v47_ = d_1862_v47_
                d_1863_v48_: C0
                nw286_ = C0()
                nw286_.ctor__()
                d_1863_v48_ = nw286_
                (globalState).f7 = default__.safeModuloInt(p0, len(d_1825_v15_))
        elif True:
            d_1864_v49_: _dafny.Array
            nw287_ = _dafny.Array(None, 2)
            d_1864_v49_ = nw287_
            d_1865_v50_: _dafny.MultiSet
            d_1865_v50_ = _dafny.MultiSet([d_1864_v49_])
            d_1866_v51_: _dafny.Map
            d_1866_v51_ = _dafny.Map({d_1806_v0_: (d_1865_v50_).cardinality})
            (globalState).f7 = ((d_1866_v51_)[(d_1806_v0_) | (d_1806_v0_)] if ((d_1806_v0_) | (d_1806_v0_)) in (d_1866_v51_) else 781)
            r1 = ((d_1825_v15_) + (d_1825_v15_)) < (d_1825_v15_)
            d_1867_v52_: T0
            nw288_ = C7()
            nw288_.ctor__()
            d_1867_v52_ = nw288_
            d_1868_v53_: D17
            d_1868_v53_ = D17_DC47(d_1867_v52_, p1, p0, p0, p0)
            d_1869_v54_: _dafny.Map
            d_1869_v54_ = _dafny.Map({d_1868_v53_: (_dafny.CodePoint('h')) not in (d_1829_v19_)})
            rhs376_ = d_1828_v18_
            rhs377_ = d_1869_v54_
            d_1828_v18_ = rhs376_
            d_1869_v54_ = rhs377_
            rhs378_ = (p0) * (((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1823_v13_ for d_1870_i6_ in range(default__.abs(913))])])))) * ((d_1825_v15_)[default__.safeIndex(p0, len(d_1825_v15_))]))
            rhs379_ = p0
            rhs380_ = p1
            lhs263_ = globalState
            lhs264_ = globalState
            lhs263_.f7 = rhs378_
            lhs264_.f15 = rhs379_
            r1 = rhs380_
            d_1871_v55_: C4
            nw289_ = C4()
            nw289_.ctor__()
            d_1871_v55_ = nw289_
        if p2:
            (globalState).f7 = (p0 if not(True) else p0)
            if p2:
                d_1872_v56_: T0
                nw290_ = C8()
                nw290_.ctor__()
                d_1872_v56_ = nw290_
                d_1873_v57_: _dafny.Map
                d_1873_v57_ = _dafny.Map({p2: d_1872_v56_})
                d_1873_v57_ = (d_1873_v57_).set(p2, (d_1872_v56_ if p2 else d_1872_v56_))
                def lambda98_(d_1874_p2_):
                    def lambda99_(d_1875_i7_):
                        return (d_1874_p2_) == (d_1874_p2_)

                    return lambda99_

                init50_ = lambda98_(p2)
                nw291_ = _dafny.Array(None, 10)
                for i0_50_ in range(nw291_.length(0)):
                    nw291_[i0_50_] = init50_(i0_50_)
                d_1828_v18_ = nw291_
                r1 = p1
                d_1876_v58_: _dafny.Seq
                d_1876_v58_ = _dafny.SeqWithoutIsStrInference([d_1815_v6_])
                d_1877_v59_: D16
                d_1877_v59_ = D16_DC41(d_1876_v58_)
                d_1878_v60_: D16
                d_1878_v60_ = D16_DC44(d_1877_v59_)
                d_1878_v60_ = d_1878_v60_
                index325_ = default__.safeIndex(692, (d_1828_v18_).length(0))
                (d_1828_v18_)[index325_] = (d_1821_v11_) < (d_1821_v11_)
                index326_ = default__.safeIndex(692, (d_1828_v18_).length(0))
                (d_1828_v18_)[index326_] = p1
            elif True:
                d_1879_v61_: _dafny.Array
                nw292_ = _dafny.Array(int(0), 15)
                d_1879_v61_ = nw292_
                index327_ = default__.safeIndex(662, (d_1879_v61_).length(0))
                (d_1879_v61_)[index327_] = (986) - (p0)
                d_1880_v62_: D3
                d_1880_v62_ = D3_DC10(p0, p0)
                d_1881_v63_: D2
                d_1881_v63_ = D2_DC7(p2, d_1826_v16_, d_1879_v61_)
                d_1882_v64_: _dafny.Seq
                d_1882_v64_ = _dafny.SeqWithoutIsStrInference([d_1881_v63_, d_1881_v63_, d_1881_v63_])
                index328_ = default__.safeIndex(662, (d_1879_v61_).length(0))
                rhs381_ = d_1821_v11_
                rhs382_ = p1
                rhs383_ = (d_1880_v62_).cf14
                rhs384_ = (d_1882_v64_) == ((d_1882_v64_) + (d_1882_v64_))
                rhs385_ = ((p0) * (p0)) - (-95)
                lhs265_ = globalState
                lhs266_ = d_1879_v61_
                lhs267_ = default__.safeIndex(662, (d_1879_v61_).length(0))
                d_1821_v11_ = rhs381_
                r1 = rhs382_
                lhs265_.f13 = rhs383_
                r1 = rhs384_
                lhs266_[lhs267_] = rhs385_
                r1 = p2
                d_1879_v61_ = d_1879_v61_
                index329_ = default__.safeIndex(662, (d_1879_v61_).length(0))
                (d_1879_v61_)[index329_] = p0
                d_1883_v65_: D2
                d_1883_v65_ = D2_DC5(d_1879_v61_)
                (d_1827_v17_).m6((len((d_1821_v11_).set(default__.safeIndex(len(_dafny.Map({128: p1})), len(d_1821_v11_)), p1)) if p2 else (0) - ((0) - ((0) - (len(default__.fm19(p2, (D19_DC52(p0, p1, p2)).cf82, globalState)))))), (d_1826_v16_) | (d_1826_v16_), (p2 if p2 else p1), d_1883_v65_, globalState)
            d_1884_v66_: C0
            nw293_ = C0()
            nw293_.ctor__()
            d_1884_v66_ = nw293_
            d_1885_v67_: _dafny.Set
            d_1885_v67_ = _dafny.Set({d_1806_v0_})
            d_1886_v68_: _dafny.Array
            nw294_ = _dafny.Array(int(0), 25)
            d_1886_v68_ = nw294_
            d_1887_v69_: _dafny.Map
            d_1887_v69_ = _dafny.Map({p0: d_1886_v68_})
            d_1888_v70_: _dafny.Seq
            d_1888_v70_ = _dafny.SeqWithoutIsStrInference([d_1828_v18_, d_1828_v18_, d_1828_v18_])
            d_1889_v71_: _dafny.MultiSet
            d_1889_v71_ = _dafny.MultiSet([p0])
            d_1890_v72_: D19
            d_1890_v72_ = D19_DC52(297, False, p2)
            d_1891_v73_: _dafny.Array
            nw295_ = _dafny.Array(None, 27)
            nw295_[int(0)] = p0
            nw295_[int(1)] = p0
            nw295_[int(2)] = len(_dafny.Set({p0, len(d_1887_v69_), p0}))
            nw295_[int(3)] = p0
            nw295_[int(4)] = len(d_1888_v70_)
            nw295_[int(5)] = p0
            nw295_[int(6)] = p0
            nw295_[int(7)] = len(d_1829_v19_)
            nw295_[int(8)] = (d_1889_v71_).cardinality
            nw295_[int(9)] = p0
            nw295_[int(10)] = p0
            nw295_[int(11)] = ((d_1820_v10_)[len(d_1829_v19_)] if (len(d_1829_v19_)) in (d_1820_v10_) else p0)
            nw295_[int(12)] = p0
            nw295_[int(13)] = 827
            nw295_[int(14)] = p0
            nw295_[int(15)] = p0
            nw295_[int(16)] = p0
            nw295_[int(17)] = (d_1890_v72_).cf81
            nw295_[int(18)] = (0) - (p0)
            nw295_[int(19)] = 607
            nw295_[int(20)] = p0
            nw295_[int(21)] = (default__.fm31(p0, p0, globalState)).cardinality
            nw295_[int(22)] = len(d_1829_v19_)
            nw295_[int(23)] = p0
            nw295_[int(24)] = p0
            nw295_[int(25)] = p0
            nw295_[int(26)] = p0
            d_1891_v73_ = nw295_
            d_1892_v74_: _dafny.MultiSet
            d_1892_v74_ = _dafny.MultiSet([d_1891_v73_])
            d_1893_v75_: _dafny.Array
            nw296_ = _dafny.Array(None, 4)
            nw296_[int(0)] = d_1892_v74_
            nw296_[int(1)] = ((d_1892_v74_).set(d_1891_v73_, default__.abs(p0))) - (d_1892_v74_)
            nw296_[int(2)] = d_1892_v74_
            nw296_[int(3)] = d_1892_v74_
            d_1893_v75_ = nw296_
            index330_ = default__.safeIndex(80, (d_1893_v75_).length(0))
            (d_1893_v75_)[index330_] = d_1892_v74_
            index331_ = default__.safeIndex(962, (d_1891_v73_).length(0))
            (d_1891_v73_)[index331_] = ((0) - (p0)) * ((0) - (len(d_1821_v11_)))
            index332_ = default__.safeIndex(80, (d_1893_v75_).length(0))
            index333_ = default__.safeIndex(962, (d_1891_v73_).length(0))
            rhs386_ = ((d_1885_v67_) | (d_1885_v67_)) - ((d_1885_v67_) - (_dafny.Set({d_1806_v0_, d_1806_v0_})))
            rhs387_ = (p0) + (len(d_1821_v11_))
            rhs388_ = ((d_1892_v74_) - (d_1892_v74_) if not(not(p2)) else d_1892_v74_)
            rhs389_ = (p0) != (p0)
            rhs390_ = default__.safeDivisionInt((d_1825_v15_)[default__.safeIndex(p0, len(d_1825_v15_))], p0)
            lhs268_ = globalState
            lhs269_ = d_1893_v75_
            lhs270_ = default__.safeIndex(80, (d_1893_v75_).length(0))
            lhs271_ = d_1891_v73_
            lhs272_ = default__.safeIndex(962, (d_1891_v73_).length(0))
            d_1885_v67_ = rhs386_
            lhs268_.f13 = rhs387_
            lhs269_[lhs270_] = rhs388_
            r1 = rhs389_
            lhs271_[lhs272_] = rhs390_
            if not(((d_1821_v11_) + (_dafny.SeqWithoutIsStrInference([p2, p2]))) <= ((d_1821_v11_) + (d_1821_v11_))):
                d_1894_v76_: C5
                nw297_ = C5()
                nw297_.ctor__()
                d_1894_v76_ = nw297_
                (d_1894_v76_).m6((d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))], default__.fm34((0) - (-881), (d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))], not(p2), globalState), not(True), D2_DC5(d_1886_v68_), globalState)
                (globalState).f4 = ((d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]) * ((0) - ((d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]))
                d_1895_v77_: _dafny.Array
                nw298_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1895_v77_ = nw298_
                rhs391_ = (d_1828_v18_ if p2 else d_1828_v18_)
                rhs392_ = d_1895_v77_
                rhs393_ = (d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]
                lhs273_ = globalState
                d_1828_v18_ = rhs391_
                d_1895_v77_ = rhs392_
                lhs273_.f13 = rhs393_
                d_1896_v78_: _dafny.Seq
                d_1896_v78_ = _dafny.SeqWithoutIsStrInference([d_1886_v68_, d_1891_v73_])
                d_1897_v79_: _dafny.Map
                d_1897_v79_ = _dafny.Map({(d_1896_v78_)[default__.safeIndex(p0, len(d_1896_v78_))]: (d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]})
                d_1897_v79_ = d_1897_v79_
            elif True:
                d_1898_v80_: _dafny.Array
                def lambda100_(d_1899_p1_, d_1900_p2_, d_1901_v73_, d_1902_p0_):
                    def lambda101_(d_1903_i8_):
                        return (_dafny.Map({d_1900_p2_: (d_1901_v73_)[default__.safeIndex(962, (d_1901_v73_).length(0))]}) if d_1899_p1_ else _dafny.Map({d_1899_p1_: d_1902_p0_}))

                    return lambda101_

                init51_ = lambda100_(p1, p2, d_1891_v73_, p0)
                nw299_ = _dafny.Array(None, 9)
                for i0_51_ in range(nw299_.length(0)):
                    nw299_[i0_51_] = init51_(i0_51_)
                d_1898_v80_ = nw299_
                d_1904_v81_: _dafny.Map
                d_1904_v81_ = _dafny.Map({p2: p0})
                index334_ = default__.safeIndex(149, (d_1898_v80_).length(0))
                (d_1898_v80_)[index334_] = (d_1904_v81_) | (_dafny.Map({(d_1890_v72_).cf83: p0}))
                index335_ = default__.safeIndex(149, (d_1898_v80_).length(0))
                (d_1898_v80_)[index335_] = (d_1904_v81_) | (d_1904_v81_)
                d_1905_v82_: _dafny.Set
                d_1905_v82_ = _dafny.Set({p2})
                d_1906_v83_: _dafny.Seq
                d_1906_v83_ = _dafny.SeqWithoutIsStrInference([d_1905_v82_, d_1905_v82_, d_1905_v82_])
                d_1907_v84_: D5
                d_1907_v84_ = D5_DC14(d_1823_v13_, False, p0, p2)
                rhs394_ = False
                rhs395_ = (D18_DC49(not(p1), p2, (d_1907_v84_).cf18, p1)).cf78
                rhs396_ = d_1906_v83_
                rhs397_ = d_1825_v15_
                r1 = rhs394_
                r1 = rhs395_
                d_1906_v83_ = rhs396_
                d_1825_v15_ = rhs397_
                d_1908_v85_: D19
                d_1908_v85_ = D19_DC53()
                d_1908_v85_ = default__.fm54((d_1884_v66_).fm29(d_1821_v11_, d_1829_v19_, p2, p0, globalState), d_1815_v6_, d_1889_v71_, globalState)
                d_1909_v86_: C0
                nw300_ = C0()
                nw300_.ctor__()
                d_1909_v86_ = nw300_
                d_1910_v88_: _dafny.Map
                d_1910_v88_ = _dafny.Map({d_1823_v13_: (self).fm2((d_1826_v16_).cardinality, globalState)})
                d_1911_v89_: _dafny.Seq
                def iife144_():
                    coll78_ = _dafny.Map()
                    compr_78_: str
                    for compr_78_ in (d_1910_v88_).keys.Elements:
                        d_1912_v87_: str = compr_78_
                        if (d_1912_v87_) in (d_1910_v88_):
                            coll78_[d_1912_v87_] = p2
                    return _dafny.Map(coll78_)
                d_1911_v89_ = _dafny.SeqWithoutIsStrInference([iife144_()
                ])
                d_1913_v90_: T0
                nw301_ = C2()
                nw301_.ctor__()
                d_1913_v90_ = nw301_
                d_1914_v91_: D17
                d_1914_v91_ = D17_DC47(d_1913_v90_, p2, (d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))], len(d_1829_v19_), (d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))])
                d_1915_v92_: _dafny.Seq
                d_1915_v92_ = _dafny.SeqWithoutIsStrInference([(d_1898_v80_)[default__.safeIndex(149, (d_1898_v80_).length(0))]])
                d_1916_v93_: _dafny.Map
                d_1916_v93_ = _dafny.Map({(d_1829_v19_)[default__.safeIndex(len((d_1915_v92_)[default__.safeIndex(p0, len(d_1915_v92_))]), len(d_1829_v19_))]: (d_1827_v17_).fm8(p0, False, (d_1913_v90_).fm2(p0, globalState), globalState)})
                d_1917_v94_: D14
                d_1917_v94_ = D14_DC36(d_1821_v11_)
                d_1918_v95_: _dafny.Map
                d_1918_v95_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p1])})
                d_1919_v96_: _dafny.Array
                nw302_ = _dafny.Array(None, 12)
                nw302_[int(0)] = d_1911_v89_
                nw302_[int(1)] = d_1911_v89_
                nw302_[int(2)] = d_1911_v89_
                nw302_[int(3)] = d_1911_v89_
                nw302_[int(4)] = (d_1911_v89_).set(default__.safeIndex((d_1914_v91_).cf73, len(d_1911_v89_)), d_1916_v93_)
                nw302_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1916_v93_ for d_1920_i9_ in range(default__.abs(881))])
                nw302_[int(6)] = _dafny.SeqWithoutIsStrInference([d_1916_v93_ for d_1921_i10_ in range(default__.abs(858))])
                nw302_[int(7)] = d_1911_v89_
                nw302_[int(8)] = (d_1911_v89_) + (_dafny.SeqWithoutIsStrInference([d_1916_v93_, d_1916_v93_, d_1916_v93_, d_1916_v93_]))
                nw302_[int(9)] = default__.fm55(p1, p0, p2, d_1917_v94_, globalState)
                nw302_[int(10)] = (d_1911_v89_) + (_dafny.SeqWithoutIsStrInference([d_1916_v93_ for d_1922_i11_ in range(default__.abs(-814))]))
                nw302_[int(11)] = default__.fm55(p1, len(((d_1918_v95_)[(d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]] if ((d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]) in (d_1918_v95_) else d_1821_v11_)), p1, d_1917_v94_, globalState)
                d_1919_v96_ = nw302_
                index336_ = default__.safeIndex(272, (d_1919_v96_).length(0))
                (d_1919_v96_)[index336_] = d_1911_v89_
                d_1923_v97_: _dafny.Map
                d_1923_v97_ = _dafny.Map({(d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]: d_1823_v13_})
                index337_ = default__.safeIndex(272, (d_1919_v96_).length(0))
                rhs398_ = (p0) + ((d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))])
                rhs399_ = p1
                rhs400_ = ((d_1911_v89_) + (d_1911_v89_)) + (((d_1911_v89_).set(default__.safeIndex((d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))], len(d_1911_v89_)), _dafny.Map({d_1823_v13_: False}))) + ((d_1911_v89_).set(default__.safeIndex(((d_1904_v81_)[p2] if (p2) in (d_1904_v81_) else (d_1891_v73_)[default__.safeIndex(962, (d_1891_v73_).length(0))]), len(d_1911_v89_)), _dafny.Map({((d_1923_v97_)[(0) - ((d_1826_v16_).cardinality)] if ((0) - ((d_1826_v16_).cardinality)) in (d_1923_v97_) else d_1823_v13_): p1}))))
                lhs274_ = globalState
                lhs275_ = d_1919_v96_
                lhs276_ = default__.safeIndex(272, (d_1919_v96_).length(0))
                lhs274_.f15 = rhs398_
                r1 = rhs399_
                lhs275_[lhs276_] = rhs400_
        elif True:
            d_1924_v98_: D1
            d_1924_v98_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_1823_v13_ for d_1925_i12_ in range(default__.abs(627))]), p1)
            pat_let_tv41_ = d_1829_v19_
            def iife145_(_pat_let33_0):
                def iife146_(d_1926_dt__update__tmp_h1_):
                    def iife147_(_pat_let34_0):
                        def iife148_(d_1927_dt__update_hcf3_h0_):
                            return D1_DC3(d_1927_dt__update_hcf3_h0_, (d_1926_dt__update__tmp_h1_).cf4)
                        return iife148_(_pat_let34_0)
                    return iife147_(pat_let_tv41_)
                return iife146_(_pat_let33_0)
            source27_ = iife145_(d_1924_v98_)
            if source27_.is_DC3:
                d_1928___mcc_h0_ = source27_.cf3
                d_1929___mcc_h1_ = source27_.cf4
                d_1930_cf4_ = d_1929___mcc_h1_
                d_1931_cf3_ = d_1928___mcc_h0_
                d_1932_v99_: D3
                d_1932_v99_ = D3_DC10(p0, (d_1826_v16_).cardinality)
                (globalState).f4 = (d_1932_v99_).cf13
                d_1933_v100_: _dafny.Array
                nw303_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                d_1933_v100_ = nw303_
                index338_ = default__.safeIndex(306, (d_1933_v100_).length(0))
                (d_1933_v100_)[index338_] = d_1931_cf3_
                index339_ = default__.safeIndex(306, (d_1933_v100_).length(0))
                rhs401_ = not((p1) or (p1))
                rhs402_ = (d_1931_cf3_) + (d_1931_cf3_)
                lhs277_ = d_1933_v100_
                lhs278_ = default__.safeIndex(306, (d_1933_v100_).length(0))
                d_1930_cf4_ = rhs401_
                lhs277_[lhs278_] = rhs402_
                (globalState).f7 = default__.safeDivisionInt(((d_1820_v10_)[982] if (982) in (d_1820_v10_) else p0), (p0) - (p0))
                r1 = ((D16_DC42(d_1930_cf4_, False, True)).cf61) == (p2)
            elif source27_.is_DC4:
                d_1934___mcc_h2_ = source27_.cf5
                d_1935_cf5_ = d_1934___mcc_h2_
                d_1936_v101_: _dafny.Map
                d_1936_v101_ = _dafny.Map({len(_dafny.Set({p2, p1, p1})): (D8_DC23(d_1821_v11_)).cf39})
                r1 = (120) == (default__.fm17(len(d_1936_v101_), globalState))
                d_1937_v102_: _dafny.Array
                nw304_ = _dafny.Array(int(0), 27)
                d_1937_v102_ = nw304_
                index340_ = default__.safeIndex(566, (d_1937_v102_).length(0))
                (d_1937_v102_)[index340_] = len(d_1829_v19_)
                index341_ = default__.safeIndex(566, (d_1937_v102_).length(0))
                (d_1937_v102_)[index341_] = p0
                index342_ = default__.safeIndex(140, (d_1828_v18_).length(0))
                (d_1828_v18_)[index342_] = p2
                index343_ = default__.safeIndex(140, (d_1828_v18_).length(0))
                (d_1828_v18_)[index343_] = (d_1821_v11_)[default__.safeIndex(default__.safeModuloInt(p0, (d_1937_v102_)[default__.safeIndex(566, (d_1937_v102_).length(0))]), len(d_1821_v11_))]
                index344_ = default__.safeIndex(140, (d_1828_v18_).length(0))
                (d_1828_v18_)[index344_] = p2
            elif True:
                d_1938___mcc_h3_ = source27_.cf2
                d_1939_cf2_ = d_1938___mcc_h3_
                d_1940_v103_: C2
                nw305_ = C2()
                nw305_.ctor__()
                d_1940_v103_ = nw305_
                r1 = p2
                index345_ = default__.safeIndex(864, (d_1828_v18_).length(0))
                (d_1828_v18_)[index345_] = (d_1829_v19_) <= (d_1829_v19_)
                index346_ = default__.safeIndex(864, (d_1828_v18_).length(0))
                rhs403_ = not(False)
                rhs404_ = p0
                rhs405_ = p1
                rhs406_ = (d_1939_cf2_) != (len(d_1825_v15_))
                lhs279_ = d_1828_v18_
                lhs280_ = default__.safeIndex(864, (d_1828_v18_).length(0))
                r1 = rhs403_
                d_1939_cf2_ = rhs404_
                lhs279_[lhs280_] = rhs405_
                r1 = rhs406_
                index347_ = default__.safeIndex(864, (d_1828_v18_).length(0))
                (d_1828_v18_)[index347_] = False
            r1 = p1
            d_1941_v104_: _dafny.Seq
            d_1941_v104_ = _dafny.SeqWithoutIsStrInference([d_1826_v16_])
            d_1942_v105_: _dafny.Array
            nw306_ = _dafny.Array(None, 18)
            nw306_[int(0)] = d_1826_v16_
            nw306_[int(1)] = (d_1826_v16_) | (d_1826_v16_)
            nw306_[int(2)] = default__.fm34(p0, (d_1826_v16_).cardinality, p2, globalState)
            nw306_[int(3)] = (d_1826_v16_) | (d_1826_v16_)
            nw306_[int(4)] = d_1826_v16_
            nw306_[int(5)] = d_1826_v16_
            nw306_[int(6)] = d_1826_v16_
            nw306_[int(7)] = d_1826_v16_
            nw306_[int(8)] = d_1826_v16_
            nw306_[int(9)] = _dafny.MultiSet([not(p2), not(p1)])
            nw306_[int(10)] = (d_1826_v16_).intersection(d_1826_v16_)
            nw306_[int(11)] = (d_1826_v16_).intersection(d_1826_v16_)
            nw306_[int(12)] = d_1826_v16_
            nw306_[int(13)] = (d_1826_v16_).intersection(d_1826_v16_)
            nw306_[int(14)] = d_1826_v16_
            nw306_[int(15)] = (d_1941_v104_)[default__.safeIndex(p0, len(d_1941_v104_))]
            nw306_[int(16)] = d_1826_v16_
            nw306_[int(17)] = d_1826_v16_
            d_1942_v105_ = nw306_
            rhs407_ = (d_1829_v19_) + ((d_1829_v19_) + (d_1829_v19_))
            rhs408_ = d_1942_v105_
            rhs409_ = p0
            lhs281_ = globalState
            d_1829_v19_ = rhs407_
            r0 = rhs408_
            lhs281_.f13 = rhs409_
            d_1943_v106_: C4
            nw307_ = C4()
            nw307_.ctor__()
            d_1943_v106_ = nw307_
            d_1944_v107_: _dafny.Map
            d_1944_v107_ = _dafny.Map({_dafny.Map({p0: d_1943_v106_}): p2})
            d_1944_v107_ = d_1944_v107_
            d_1945_v108_: _dafny.Array
            nw308_ = _dafny.Array(D21.default()(), 13)
            d_1945_v108_ = nw308_
            d_1946_v109_: _dafny.Array
            nw309_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_1946_v109_ = nw309_
            d_1947_v110_: D21
            d_1947_v110_ = D21_DC56(d_1946_v109_)
            d_1948_v111_: D21
            d_1948_v111_ = D21_DC56((d_1947_v110_).cf89)
            index348_ = default__.safeIndex(57, (d_1945_v108_).length(0))
            (d_1945_v108_)[index348_] = d_1948_v111_
            index349_ = default__.safeIndex(57, (d_1945_v108_).length(0))
            (d_1945_v108_)[index349_] = d_1947_v110_
        d_1949_v112_: _dafny.Array
        nw310_ = _dafny.Array(_dafny.MultiSet({}), 4)
        d_1949_v112_ = nw310_
        r0 = d_1949_v112_
        r1 = p1
        r2 = d_1823_v13_
        return r0, r1, r2


class C11(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, p1, globalState):
        return False

    def fm2(self, p0, globalState):
        def lambda102_(source28_):
            if source28_.is_DC3:
                d_1950___mcc_h0_ = source28_.cf3
                d_1951___mcc_h1_ = source28_.cf4
                d_1952_cf4_ = d_1951___mcc_h1_
                d_1953_cf3_ = d_1950___mcc_h0_
                return d_1953_cf3_
            elif source28_.is_DC4:
                d_1954___mcc_h2_ = source28_.cf5
                d_1955_cf5_ = d_1954___mcc_h2_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbfn"))
            elif True:
                d_1956___mcc_h3_ = source28_.cf2
                d_1957_cf2_ = d_1956___mcc_h3_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aredcecm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lopn")))

        return len(lambda102_((D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjiblb")), False) if False else D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahql")), True))))

    def m1(self, p0, p1, globalState):
        d_1958_v0_: _dafny.Seq
        d_1958_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_1958_v0_ = d_1958_v0_
        d_1959_i0_: int
        d_1959_i0_ = 0
        with _dafny.label("13"):
            while False:
                with _dafny.c_label("13"):
                    if (d_1959_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_1959_i0_ = (d_1959_i0_) + (1)
                    d_1960_v1_: bool
                    d_1960_v1_ = False
                    d_1960_v1_ = d_1960_v1_
                    if True:
                        d_1960_v1_ = True
                        d_1961_v2_: str
                        d_1961_v2_ = _dafny.CodePoint('g')
                        d_1962_v3_: int
                        d_1962_v3_ = -854
                        rhs410_ = (d_1962_v3_ if d_1960_v1_ else d_1962_v3_)
                        rhs411_ = d_1961_v2_
                        lhs282_ = globalState
                        lhs282_.f7 = rhs410_
                        d_1961_v2_ = rhs411_
                        d_1963_v4_: _dafny.Map
                        d_1963_v4_ = _dafny.Map({d_1962_v3_: len(_dafny.Set({-798}))})
                        d_1964_v5_: _dafny.Seq
                        d_1964_v5_ = _dafny.SeqWithoutIsStrInference([len(d_1963_v4_), d_1962_v3_])
                        d_1965_v6_: _dafny.Set
                        d_1965_v6_ = _dafny.Set({d_1962_v3_, -516, len(_dafny.SeqWithoutIsStrInference([d_1961_v2_ for d_1966_i1_ in range(default__.abs(-427))]))})
                        d_1967_v7_: _dafny.Array
                        nw311_ = _dafny.Array(None, 5)
                        nw311_[int(0)] = d_1960_v1_
                        nw311_[int(1)] = d_1960_v1_
                        nw311_[int(2)] = d_1960_v1_
                        nw311_[int(3)] = d_1960_v1_
                        nw311_[int(4)] = d_1960_v1_
                        d_1967_v7_ = nw311_
                        d_1968_v8_: _dafny.Map
                        d_1968_v8_ = _dafny.Map({d_1967_v7_: 950})
                        d_1969_v9_: _dafny.Seq
                        d_1969_v9_ = _dafny.SeqWithoutIsStrInference([d_1960_v1_])
                        d_1970_v10_: _dafny.Array
                        nw312_ = _dafny.Array(None, 25)
                        nw312_[int(0)] = d_1962_v3_
                        nw312_[int(1)] = d_1962_v3_
                        nw312_[int(2)] = 591
                        nw312_[int(3)] = d_1962_v3_
                        nw312_[int(4)] = len(d_1964_v5_)
                        nw312_[int(5)] = len(d_1965_v6_)
                        nw312_[int(6)] = (0) - (d_1962_v3_)
                        nw312_[int(7)] = len(d_1968_v8_)
                        nw312_[int(8)] = d_1962_v3_
                        nw312_[int(9)] = 321
                        nw312_[int(10)] = d_1962_v3_
                        nw312_[int(11)] = d_1962_v3_
                        nw312_[int(12)] = d_1962_v3_
                        nw312_[int(13)] = d_1962_v3_
                        nw312_[int(14)] = len(d_1969_v9_)
                        nw312_[int(15)] = d_1962_v3_
                        nw312_[int(16)] = d_1962_v3_
                        nw312_[int(17)] = len(p0)
                        nw312_[int(18)] = (0) - ((0) - ((0) - (len(d_1964_v5_))))
                        nw312_[int(19)] = d_1962_v3_
                        nw312_[int(20)] = d_1962_v3_
                        nw312_[int(21)] = d_1962_v3_
                        nw312_[int(22)] = d_1962_v3_
                        nw312_[int(23)] = (0) - (d_1962_v3_)
                        nw312_[int(24)] = d_1962_v3_
                        d_1970_v10_ = nw312_
                        d_1971_v11_: _dafny.Seq
                        d_1971_v11_ = _dafny.SeqWithoutIsStrInference([d_1970_v10_])
                        rhs412_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))) + (d_1958_v0_)) != ((default__.fm5(globalState)).set(default__.safeIndex(len(d_1964_v5_), len(default__.fm5(globalState))), d_1961_v2_))
                        rhs413_ = d_1962_v3_
                        rhs414_ = d_1971_v11_
                        lhs283_ = globalState
                        d_1960_v1_ = rhs412_
                        lhs283_.f15 = rhs413_
                        d_1971_v11_ = rhs414_
                        d_1958_v0_ = p0
                        (globalState).f13 = -597
                    elif True:
                        d_1972_v12_: _dafny.Array
                        nw313_ = _dafny.Array(D1.default()(), 21)
                        d_1972_v12_ = nw313_
                        d_1973_v13_: D1
                        d_1973_v13_ = D1_DC3(p0, True)
                        index350_ = default__.safeIndex(71, (d_1972_v12_).length(0))
                        (d_1972_v12_)[index350_] = d_1973_v13_
                        index351_ = default__.safeIndex(71, (d_1972_v12_).length(0))
                        (d_1972_v12_)[index351_] = d_1973_v13_
                        d_1960_v1_ = d_1960_v1_
                        d_1974_v14_: int
                        d_1974_v14_ = 586
                        d_1975_v15_: _dafny.Map
                        d_1975_v15_ = _dafny.Map({d_1974_v14_: d_1974_v14_})
                        (globalState).f13 = (self).fm2((d_1974_v14_) + (len(d_1975_v15_)), globalState)
                        d_1976_v16_: str
                        d_1976_v16_ = _dafny.CodePoint('w')
                        d_1977_v17_: _dafny.Map
                        d_1977_v17_ = _dafny.Map({d_1976_v16_: p0})
                        d_1977_v17_ = ((d_1977_v17_) | (d_1977_v17_)) | (_dafny.Map({d_1976_v16_: d_1958_v0_}))
                        (globalState).f4 = d_1974_v14_
                    d_1978_v18_: _dafny.Array
                    nw314_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                    d_1978_v18_ = nw314_
                    index352_ = default__.safeIndex(295, (d_1978_v18_).length(0))
                    (d_1978_v18_)[index352_] = _dafny.CodePoint('l')
                    d_1979_v19_: str
                    d_1979_v19_ = _dafny.CodePoint('m')
                    index353_ = default__.safeIndex(295, (d_1978_v18_).length(0))
                    (d_1978_v18_)[index353_] = d_1979_v19_
                    d_1980_v20_: int
                    d_1980_v20_ = 534
                    d_1981_v21_: _dafny.Seq
                    d_1981_v21_ = _dafny.SeqWithoutIsStrInference([-358, d_1980_v20_])
                    d_1982_v22_: _dafny.MultiSet
                    d_1982_v22_ = _dafny.MultiSet([d_1980_v20_, d_1980_v20_, (0) - (d_1980_v20_), len(d_1981_v21_), len(_dafny.SeqWithoutIsStrInference([d_1980_v20_, d_1980_v20_]))])
                    d_1983_v23_: _dafny.MultiSet
                    d_1983_v23_ = _dafny.MultiSet([d_1960_v1_, d_1960_v1_])
                    index354_ = default__.safeIndex(295, (d_1978_v18_).length(0))
                    rhs415_ = (0) - (((d_1982_v22_).set(d_1980_v20_, default__.abs((d_1983_v23_).cardinality))).cardinality)
                    rhs416_ = d_1979_v19_
                    rhs417_ = (d_1980_v20_) + ((0) - (d_1980_v20_))
                    rhs418_ = (self).fm2((d_1980_v20_) * (d_1980_v20_), globalState)
                    rhs419_ = (d_1980_v20_) <= (len(d_1981_v21_))
                    lhs284_ = globalState
                    lhs285_ = d_1978_v18_
                    lhs286_ = default__.safeIndex(295, (d_1978_v18_).length(0))
                    lhs287_ = globalState
                    lhs288_ = globalState
                    lhs284_.f13 = rhs415_
                    lhs285_[lhs286_] = rhs416_
                    lhs287_.f7 = rhs417_
                    lhs288_.f13 = rhs418_
                    d_1960_v1_ = rhs419_
                    pass
            pass
        d_1984_v24_: bool
        d_1984_v24_ = True
        d_1985_v25_: int
        d_1985_v25_ = 261
        d_1986_v26_: _dafny.MultiSet
        d_1986_v26_ = _dafny.MultiSet([d_1985_v25_])
        d_1987_v27_: D1
        d_1987_v27_ = D1_DC4(d_1986_v26_)
        source29_ = (d_1987_v27_ if d_1984_v24_ else d_1987_v27_)
        if source29_.is_DC3:
            d_1988___mcc_h0_ = source29_.cf3
            d_1989___mcc_h1_ = source29_.cf4
            d_1990_cf4_ = d_1989___mcc_h1_
            d_1991_cf3_ = d_1988___mcc_h0_
            d_1992_v28_: D1
            d_1992_v28_ = D1_DC2((d_1986_v26_).cardinality)
            (globalState).f7 = ((D1_DC2(495) if d_1990_cf4_ else d_1992_v28_)).cf2
            rhs420_ = d_1985_v25_
            rhs421_ = d_1990_cf4_
            lhs289_ = globalState
            lhs289_.f4 = rhs420_
            d_1984_v24_ = rhs421_
            d_1990_cf4_ = d_1990_cf4_
            d_1993_v29_: T0
            nw315_ = C10()
            nw315_.ctor__()
            d_1993_v29_ = nw315_
            d_1993_v29_ = d_1993_v29_
        elif source29_.is_DC4:
            d_1994___mcc_h2_ = source29_.cf5
            d_1995_cf5_ = d_1994___mcc_h2_
            d_1996_v30_: _dafny.Array
            nw316_ = _dafny.Array(None, 7)
            nw316_[int(0)] = d_1984_v24_
            nw316_[int(1)] = False
            nw316_[int(2)] = d_1984_v24_
            nw316_[int(3)] = d_1984_v24_
            nw316_[int(4)] = d_1984_v24_
            nw316_[int(5)] = d_1984_v24_
            nw316_[int(6)] = d_1984_v24_
            d_1996_v30_ = nw316_
            d_1997_v31_: _dafny.Seq
            d_1997_v31_ = _dafny.SeqWithoutIsStrInference([d_1996_v30_, d_1996_v30_])
            d_1998_v32_: D7
            d_1998_v32_ = D7_DC19(d_1997_v31_)
            d_1999_v33_: _dafny.Map
            d_1999_v33_ = _dafny.Map({_dafny.CodePoint('r'): d_1998_v32_})
            d_2000_v34_: _dafny.Set
            d_2000_v34_ = _dafny.Set({d_1985_v25_})
            d_2001_v35_: _dafny.Array
            nw317_ = _dafny.Array(None, 19)
            nw317_[int(0)] = 223
            nw317_[int(1)] = d_1985_v25_
            nw317_[int(2)] = d_1985_v25_
            nw317_[int(3)] = d_1985_v25_
            nw317_[int(4)] = (self).fm2(704, globalState)
            nw317_[int(5)] = len(d_1999_v33_)
            nw317_[int(6)] = d_1985_v25_
            nw317_[int(7)] = d_1985_v25_
            nw317_[int(8)] = d_1985_v25_
            nw317_[int(9)] = 447
            nw317_[int(10)] = d_1985_v25_
            nw317_[int(11)] = len(d_2000_v34_)
            nw317_[int(12)] = d_1985_v25_
            nw317_[int(13)] = d_1985_v25_
            nw317_[int(14)] = d_1985_v25_
            nw317_[int(15)] = d_1985_v25_
            nw317_[int(16)] = d_1985_v25_
            nw317_[int(17)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvirhay")))
            nw317_[int(18)] = d_1985_v25_
            d_2001_v35_ = nw317_
            d_2002_v36_: _dafny.Map
            d_2002_v36_ = _dafny.Map({(0) - ((918) * (d_1985_v25_)): d_2001_v35_})
            d_2002_v36_ = (d_2002_v36_).set((0) - (d_1985_v25_), d_2001_v35_)
            d_2003_v37_: C9
            nw318_ = C9()
            nw318_.ctor__()
            d_2003_v37_ = nw318_
            d_2004_v38_: str
            d_2004_v38_ = _dafny.CodePoint('x')
            d_2005_v39_: _dafny.Set
            d_2005_v39_ = _dafny.Set({d_2004_v38_, d_2004_v38_, d_2004_v38_, default__.fm38(d_1984_v24_, globalState), _dafny.CodePoint('a')})
            d_2006_v41_: _dafny.Seq
            def iife149_():
                coll79_ = _dafny.Set()
                compr_79_: str
                for compr_79_ in (d_2005_v39_).Elements:
                    d_2007_v40_: str = compr_79_
                    if (d_2007_v40_) in (d_2005_v39_):
                        coll79_ = coll79_.union(_dafny.Set([d_2007_v40_]))
                return _dafny.Set(coll79_)
            d_2006_v41_ = _dafny.SeqWithoutIsStrInference([(d_2005_v39_) | (iife149_()
            )])
            d_2005_v39_ = (d_2006_v41_)[default__.safeIndex(default__.safeDivisionInt(d_1985_v25_, d_1985_v25_), len(d_2006_v41_))]
            d_1985_v25_ = d_1985_v25_
        elif True:
            d_2008___mcc_h3_ = source29_.cf2
            d_2009_cf2_ = d_2008___mcc_h3_
            d_2010_v42_: _dafny.Array
            nw319_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_2010_v42_ = nw319_
            d_2011_v43_: _dafny.Array
            nw320_ = _dafny.Array(False, 5)
            d_2011_v43_ = nw320_
            index355_ = default__.safeIndex(545, (d_2010_v42_).length(0))
            (d_2010_v42_)[index355_] = d_2011_v43_
            index356_ = default__.safeIndex(545, (d_2010_v42_).length(0))
            rhs422_ = d_2009_cf2_
            rhs423_ = d_2011_v43_
            rhs424_ = True
            lhs290_ = d_2010_v42_
            lhs291_ = default__.safeIndex(545, (d_2010_v42_).length(0))
            d_1985_v25_ = rhs422_
            lhs290_[lhs291_] = rhs423_
            d_1984_v24_ = rhs424_
            if d_1984_v24_:
                d_2012_v44_: _dafny.Map
                d_2012_v44_ = _dafny.Map({not (d_1984_v24_) or (d_1984_v24_): default__.fm5(globalState)})
                d_2013_v45_: D1
                d_2013_v45_ = D1_DC3(d_1958_v0_, d_1984_v24_)
                d_2012_v44_ = (d_2012_v44_).set(d_1984_v24_, ((d_2013_v45_).cf3) + (d_1958_v0_))
                d_2014_v46_: _dafny.Seq
                d_2014_v46_ = _dafny.SeqWithoutIsStrInference([d_1984_v24_])
                d_2015_v47_: _dafny.Map
                d_2015_v47_ = _dafny.Map({d_2014_v46_: d_1984_v24_})
                d_2015_v47_ = (d_2015_v47_) | (d_2015_v47_)
                d_2016_v48_: _dafny.Seq
                d_2016_v48_ = _dafny.SeqWithoutIsStrInference([d_1986_v26_, d_1986_v26_])
                d_2017_v49_: _dafny.Seq
                d_2017_v49_ = _dafny.SeqWithoutIsStrInference([(d_2016_v48_)[default__.safeIndex(d_1985_v25_, len(d_2016_v48_))]])
                (globalState).f4 = len(((d_2016_v48_) + (d_2016_v48_)) + (d_2016_v48_))
                d_2018_v50_: C10
                nw321_ = C10()
                nw321_.ctor__()
                d_2018_v50_ = nw321_
                d_2018_v50_ = d_2018_v50_
                (globalState).f15 = d_2009_cf2_
            elif True:
                d_2019_v51_: _dafny.Array
                def lambda103_(d_2020_v24_):
                    def lambda104_(d_2021_i2_):
                        return _dafny.Set({d_2020_v24_})

                    return lambda104_

                init52_ = lambda103_(d_1984_v24_)
                nw322_ = _dafny.Array(None, 15)
                for i0_52_ in range(nw322_.length(0)):
                    nw322_[i0_52_] = init52_(i0_52_)
                d_2019_v51_ = nw322_
                d_2019_v51_ = d_2019_v51_
                arr2_ = (d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]
                index357_ = default__.safeIndex(335, ((d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]).length(0))
                arr2_[index357_] = False
                d_2022_v52_: _dafny.Array
                nw323_ = _dafny.Array(int(0), 19)
                d_2022_v52_ = nw323_
                d_2023_v53_: _dafny.Set
                d_2023_v53_ = _dafny.Set({d_2022_v52_, d_2022_v52_})
                d_2024_v54_: _dafny.Seq
                d_2024_v54_ = _dafny.SeqWithoutIsStrInference([len(d_2023_v53_)])
                arr3_ = (d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]
                index358_ = default__.safeIndex(335, ((d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]).length(0))
                arr3_[index358_] = ((d_2009_cf2_) * (d_2009_cf2_)) > ((d_2024_v54_)[default__.safeIndex(984, len(d_2024_v54_))])
                d_2025_v55_: str
                d_2025_v55_ = _dafny.CodePoint('i')
                d_2026_v56_: _dafny.MultiSet
                d_2026_v56_ = _dafny.MultiSet([d_2025_v55_])
                arr4_ = (d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]
                index359_ = default__.safeIndex(335, ((d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]).length(0))
                arr4_[index359_] = (d_2026_v56_).isdisjoint(d_2026_v56_)
                (globalState).f13 = default__.fm17((0) - (d_1985_v25_), globalState)
                arr5_ = (d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]
                index360_ = default__.safeIndex(335, ((d_2010_v42_)[default__.safeIndex(545, (d_2010_v42_).length(0))]).length(0))
                arr5_[index360_] = d_1984_v24_
            d_2027_v57_: _dafny.Map
            d_2027_v57_ = _dafny.Map({d_1984_v24_: d_1984_v24_})
            d_2028_v58_: _dafny.Map
            d_2028_v58_ = _dafny.Map({((d_2027_v57_)[not(d_1984_v24_)] if (not(d_1984_v24_)) in (d_2027_v57_) else d_1984_v24_): d_2027_v57_})
            d_2029_v59_: _dafny.Seq
            d_2029_v59_ = _dafny.SeqWithoutIsStrInference([(self).fm1(d_2028_v58_, d_2009_cf2_, globalState)])
            d_2030_v60_: _dafny.Map
            d_2030_v60_ = _dafny.Map({d_2029_v59_: d_2009_cf2_})
            d_2031_v61_: str
            d_2031_v61_ = _dafny.CodePoint('o')
            d_2032_v62_: _dafny.Map
            d_2032_v62_ = _dafny.Map({d_2031_v61_: d_1986_v26_})
            d_2033_v64_: _dafny.Seq
            def iife150_():
                coll80_ = _dafny.Map()
                compr_80_: str
                for compr_80_ in (d_1958_v0_).Elements:
                    d_2034_v63_: str = compr_80_
                    if (d_2034_v63_) in (d_1958_v0_):
                        coll80_[d_2034_v63_] = d_1986_v26_
                return _dafny.Map(coll80_)
            d_2033_v64_ = _dafny.SeqWithoutIsStrInference([d_2032_v62_, iife150_()
            , d_2032_v62_])
            d_2035_v65_: _dafny.Seq
            d_2035_v65_ = _dafny.SeqWithoutIsStrInference([d_2009_cf2_, (len(d_2030_v60_)) - (default__.fm17(d_1985_v25_, globalState)), len((d_2033_v64_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1984_v24_, d_1984_v24_])), len(d_2033_v64_))])])
            d_2035_v65_ = d_2035_v65_
            d_1984_v24_ = ((0) - (d_2009_cf2_)) >= (678)
        d_2036_v66_: _dafny.Array
        nw324_ = _dafny.Array(int(0), 12)
        d_2036_v66_ = nw324_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_2036_v66_).length(0)):
            d_2037_i3_: int = guard_loop_2_
            if (True) and (((0) <= (d_2037_i3_)) and ((d_2037_i3_) < ((d_2036_v66_).length(0)))):
                (d_2036_v66_)[(d_2037_i3_)] = (d_2037_i3_) - ((d_1985_v25_) + (d_1985_v25_))
        d_2038_v67_: _dafny.Array
        nw325_ = _dafny.Array(False, 4)
        d_2038_v67_ = nw325_
        d_2039_v68_: D0
        d_2039_v68_ = D0_DC0(d_2038_v67_)
        pat_let_tv42_ = d_2038_v67_
        pat_let_tv43_ = d_2038_v67_
        d_2040_v69_: _dafny.Array
        nw326_ = _dafny.Array(None, 5)
        def iife151_(_pat_let35_0):
            def iife152_(d_2041_dt__update__tmp_h0_):
                def iife153_(_pat_let36_0):
                    def iife154_(d_2042_dt__update_hcf0_h0_):
                        return D0_DC0(d_2042_dt__update_hcf0_h0_)
                    return iife154_(_pat_let36_0)
                return iife153_(pat_let_tv42_)
            return iife152_(_pat_let35_0)
        nw326_[int(0)] = iife151_(d_2039_v68_)
        def iife155_(_pat_let37_0):
            def iife156_(d_2043_dt__update__tmp_h1_):
                def iife157_(_pat_let38_0):
                    def iife158_(d_2044_dt__update_hcf0_h1_):
                        return D0_DC0(d_2044_dt__update_hcf0_h1_)
                    return iife158_(_pat_let38_0)
                return iife157_(pat_let_tv43_)
            return iife156_(_pat_let37_0)
        nw326_[int(1)] = iife155_(d_2039_v68_)
        nw326_[int(2)] = d_2039_v68_
        nw326_[int(3)] = d_2039_v68_
        nw326_[int(4)] = d_2039_v68_
        d_2040_v69_ = nw326_
        index361_ = default__.safeIndex(835, (d_2040_v69_).length(0))
        (d_2040_v69_)[index361_] = D0_DC0(d_2038_v67_)
        d_2045_v70_: _dafny.Array
        nw327_ = _dafny.Array(_dafny.Array(None, 0), 7)
        d_2045_v70_ = nw327_
        index362_ = default__.safeIndex(703, (d_2045_v70_).length(0))
        (d_2045_v70_)[index362_] = d_2036_v66_
        d_2046_v71_: D6
        d_2046_v71_ = D6_DC17(d_1984_v24_, not(d_1984_v24_), d_1984_v24_)
        d_2047_v72_: _dafny.Seq
        d_2047_v72_ = _dafny.SeqWithoutIsStrInference([d_2046_v71_, d_2046_v71_])
        index363_ = default__.safeIndex(592, (d_2036_v66_).length(0))
        (d_2036_v66_)[index363_] = len(d_2047_v72_)
        d_2048_v73_: _dafny.Map
        d_2048_v73_ = _dafny.Map({(d_2036_v66_ if d_1984_v24_ else d_2036_v66_): d_1985_v25_})
        d_2049_v74_: _dafny.MultiSet
        d_2049_v74_ = _dafny.MultiSet([d_1984_v24_, d_1984_v24_])
        index364_ = default__.safeIndex(835, (d_2040_v69_).length(0))
        index365_ = default__.safeIndex(703, (d_2045_v70_).length(0))
        index366_ = default__.safeIndex(592, (d_2036_v66_).length(0))
        rhs425_ = d_2039_v68_
        rhs426_ = d_2036_v66_
        rhs427_ = default__.safeDivisionInt((d_1985_v25_ if d_1984_v24_ else (0) - (d_1985_v25_)), default__.safeModuloInt((d_2049_v74_).cardinality, d_1985_v25_))
        rhs428_ = d_2048_v73_
        lhs292_ = d_2040_v69_
        lhs293_ = default__.safeIndex(835, (d_2040_v69_).length(0))
        lhs294_ = d_2045_v70_
        lhs295_ = default__.safeIndex(703, (d_2045_v70_).length(0))
        lhs296_ = d_2036_v66_
        lhs297_ = default__.safeIndex(592, (d_2036_v66_).length(0))
        lhs292_[lhs293_] = rhs425_
        lhs294_[lhs295_] = rhs426_
        lhs296_[lhs297_] = rhs427_
        d_2048_v73_ = rhs428_
        d_2050_v75_: int
        d_2051_v76_: _dafny.Map
        d_2052_v77_: int
        d_2053_v78_: _dafny.Array
        out15_: int
        out16_: _dafny.Map
        out17_: int
        out18_: _dafny.Array
        out15_, out16_, out17_, out18_ = (self).m3(globalState)
        d_2050_v75_ = out15_
        d_2051_v76_ = out16_
        d_2052_v77_ = out17_
        d_2053_v78_ = out18_

    def m3(self, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_2054_v0_: C9
        nw328_ = C9()
        nw328_.ctor__()
        d_2054_v0_ = nw328_
        d_2055_v1_: int
        d_2055_v1_ = 196
        d_2056_v2_: D4
        d_2056_v2_ = D4_DC12((0) - (d_2055_v1_))
        d_2057_v3_: _dafny.Map
        d_2057_v3_ = _dafny.Map({d_2055_v1_: (default__.fm56((0) - (d_2055_v1_), (d_2056_v2_).cf16, globalState)).cf35})
        d_2057_v3_ = (d_2057_v3_).set(d_2055_v1_, True)
        d_2058_v4_: bool
        d_2058_v4_ = False
        d_2059_v5_: D20
        d_2059_v5_ = D20_DC55((0) - (d_2055_v1_), d_2055_v1_, d_2058_v4_, not(d_2058_v4_))
        d_2058_v4_ = (d_2059_v5_).cf87
        d_2060_v6_: _dafny.Seq
        d_2060_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxqw"))
        d_2061_v7_: _dafny.Map
        d_2061_v7_ = _dafny.Map({d_2058_v4_: d_2055_v1_})
        d_2062_v8_: _dafny.Seq
        d_2062_v8_ = _dafny.SeqWithoutIsStrInference([d_2058_v4_, (d_2054_v0_).fm8(len(d_2061_v7_), d_2058_v4_, 887, globalState)])
        d_2063_v9_: _dafny.Set
        d_2063_v9_ = _dafny.Set({d_2055_v1_})
        d_2064_v10_: D1
        d_2064_v10_ = D1_DC3(default__.fm30(557, not((d_2062_v8_)[default__.safeIndex(len(default__.fm5(globalState)), len(d_2062_v8_))]), d_2063_v9_, False, globalState), False)
        d_2060_v6_ = (d_2064_v10_).cf3
        d_2065_v11_: _dafny.Map
        d_2065_v11_ = _dafny.Map({d_2055_v1_: len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2066_i0_ in range(default__.abs(-67))])) + (d_2060_v6_))})
        d_2065_v11_ = (d_2065_v11_).set(d_2055_v1_, d_2055_v1_)
        d_2067_v12_: _dafny.Array
        nw329_ = _dafny.Array(int(0), 12)
        d_2067_v12_ = nw329_
        d_2068_v13_: _dafny.MultiSet
        d_2068_v13_ = _dafny.MultiSet([d_2058_v4_, d_2058_v4_])
        d_2069_v14_: _dafny.Map
        d_2069_v14_ = _dafny.Map({(d_2068_v13_).cardinality: d_2067_v12_})
        d_2070_v15_: _dafny.Array
        nw330_ = _dafny.Array(None, 12)
        nw330_[int(0)] = d_2067_v12_
        nw330_[int(1)] = d_2067_v12_
        nw330_[int(2)] = d_2067_v12_
        nw330_[int(3)] = d_2067_v12_
        nw330_[int(4)] = d_2067_v12_
        nw330_[int(5)] = d_2067_v12_
        nw330_[int(6)] = d_2067_v12_
        nw330_[int(7)] = d_2067_v12_
        nw330_[int(8)] = d_2067_v12_
        nw330_[int(9)] = ((d_2069_v14_)[456] if (456) in (d_2069_v14_) else d_2067_v12_)
        nw330_[int(10)] = d_2067_v12_
        nw330_[int(11)] = d_2067_v12_
        d_2070_v15_ = nw330_
        d_2071_v16_: _dafny.Seq
        d_2071_v16_ = _dafny.SeqWithoutIsStrInference([d_2070_v15_, d_2070_v15_, d_2070_v15_, d_2070_v15_])
        d_2072_v17_: _dafny.Array
        nw331_ = _dafny.Array(None, 25)
        nw331_[int(0)] = d_2070_v15_
        nw331_[int(1)] = d_2070_v15_
        nw331_[int(2)] = d_2070_v15_
        nw331_[int(3)] = d_2070_v15_
        nw331_[int(4)] = d_2070_v15_
        nw331_[int(5)] = d_2070_v15_
        nw331_[int(6)] = d_2070_v15_
        nw331_[int(7)] = d_2070_v15_
        nw331_[int(8)] = d_2070_v15_
        nw331_[int(9)] = d_2070_v15_
        nw331_[int(10)] = d_2070_v15_
        nw331_[int(11)] = d_2070_v15_
        nw331_[int(12)] = d_2070_v15_
        nw331_[int(13)] = d_2070_v15_
        nw331_[int(14)] = d_2070_v15_
        nw331_[int(15)] = d_2070_v15_
        nw331_[int(16)] = d_2070_v15_
        nw331_[int(17)] = d_2070_v15_
        nw331_[int(18)] = d_2070_v15_
        nw331_[int(19)] = d_2070_v15_
        nw331_[int(20)] = d_2070_v15_
        nw331_[int(21)] = d_2070_v15_
        nw331_[int(22)] = d_2070_v15_
        nw331_[int(23)] = d_2070_v15_
        nw331_[int(24)] = (d_2071_v16_)[default__.safeIndex((0) - (d_2055_v1_), len(d_2071_v16_))]
        d_2072_v17_ = nw331_
        index367_ = default__.safeIndex(948, (d_2072_v17_).length(0))
        (d_2072_v17_)[index367_] = d_2070_v15_
        index368_ = default__.safeIndex(948, (d_2072_v17_).length(0))
        (d_2072_v17_)[index368_] = d_2070_v15_
        d_2073_v18_: _dafny.Array
        def lambda105_(d_2074_i1_):
            return _dafny.CodePoint('b')

        init53_ = lambda105_
        nw332_ = _dafny.Array(None, 28)
        for i0_53_ in range(nw332_.length(0)):
            nw332_[i0_53_] = init53_(i0_53_)
        d_2073_v18_ = nw332_
        d_2075_v19_: _dafny.Map
        d_2075_v19_ = _dafny.Map({default__.fm30(d_2055_v1_, d_2058_v4_, d_2063_v9_, not(d_2058_v4_), globalState): d_2073_v18_})
        r0 = len((d_2075_v19_) | ((d_2075_v19_) | (d_2075_v19_)))
        d_2076_v20_: _dafny.Map
        d_2076_v20_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2055_v1_, d_2055_v1_])): d_2058_v4_})
        d_2077_v21_: _dafny.Seq
        d_2077_v21_ = _dafny.SeqWithoutIsStrInference([d_2076_v20_, d_2076_v20_])
        r1 = (d_2077_v21_)[default__.safeIndex((166) - (d_2055_v1_), len(d_2077_v21_))]
        r2 = d_2055_v1_
        d_2078_v22_: _dafny.Array
        def lambda106_(d_2079_v7_):
            def lambda107_(d_2080_i2_):
                return d_2079_v7_

            return lambda107_

        init54_ = lambda106_(d_2061_v7_)
        nw333_ = _dafny.Array(None, 23)
        for i0_54_ in range(nw333_.length(0)):
            nw333_[i0_54_] = init54_(i0_54_)
        d_2078_v22_ = nw333_
        r3 = d_2078_v22_
        return r0, r1, r2, r3

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_2081_i0_: int
        d_2081_i0_ = 0
        with _dafny.label("14"):
            while p3:
                with _dafny.c_label("14"):
                    if (d_2081_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2081_i0_ = (d_2081_i0_) + (1)
                    d_2082_v0_: bool
                    d_2082_v0_ = True
                    d_2083_v1_: str
                    d_2083_v1_ = _dafny.CodePoint('w')
                    d_2084_v2_: D5
                    d_2084_v2_ = D5_DC14(d_2083_v1_, p1, p0, p3)
                    d_2082_v0_ = (p1) and ((d_2084_v2_).cf21)
                    d_2085_v3_: int
                    d_2086_v4_: _dafny.Map
                    d_2087_v5_: int
                    d_2088_v6_: _dafny.Array
                    out19_: int
                    out20_: _dafny.Map
                    out21_: int
                    out22_: _dafny.Array
                    out19_, out20_, out21_, out22_ = (self).m3(globalState)
                    d_2085_v3_ = out19_
                    d_2086_v4_ = out20_
                    d_2087_v5_ = out21_
                    d_2088_v6_ = out22_
                    if True:
                        (globalState).f13 = d_2087_v5_
                        (globalState).f4 = (d_2087_v5_) - ((self).fm2(d_2087_v5_, globalState))
                        d_2089_v7_: _dafny.Seq
                        d_2089_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxs"))
                        d_2090_v8_: _dafny.Map
                        d_2090_v8_ = _dafny.Map({d_2087_v5_: 327})
                        rhs429_ = default__.fm17(p0, globalState)
                        rhs430_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wse"))
                        rhs431_ = (d_2085_v3_) + ((723) * (default__.fm17(d_2085_v3_, globalState)))
                        rhs432_ = ((0) - (p0)) * (p0)
                        rhs433_ = d_2090_v8_
                        lhs298_ = globalState
                        lhs299_ = globalState
                        lhs300_ = globalState
                        lhs298_.f4 = rhs429_
                        d_2089_v7_ = rhs430_
                        lhs299_.f13 = rhs431_
                        lhs300_.f13 = rhs432_
                        d_2090_v8_ = rhs433_
                        d_2091_v9_: C0
                        nw334_ = C0()
                        nw334_.ctor__()
                        d_2091_v9_ = nw334_
                        d_2092_v10_: _dafny.Seq
                        d_2092_v10_ = _dafny.SeqWithoutIsStrInference([p1, d_2082_v0_, p1])
                        d_2093_v11_: _dafny.Seq
                        d_2093_v11_ = _dafny.SeqWithoutIsStrInference([len(d_2092_v10_)])
                        d_2094_v12_: _dafny.Seq
                        d_2094_v12_ = _dafny.SeqWithoutIsStrInference([d_2093_v11_])
                        d_2095_v13_: _dafny.Seq
                        d_2095_v13_ = _dafny.SeqWithoutIsStrInference([d_2089_v7_, d_2089_v7_])
                        d_2096_v14_: _dafny.Array
                        nw335_ = _dafny.Array(None, 21)
                        nw335_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gif"))
                        nw335_[int(1)] = d_2089_v7_
                        nw335_[int(2)] = d_2089_v7_
                        nw335_[int(3)] = (default__.fm5(globalState)).set(default__.safeIndex(d_2087_v5_, len(default__.fm5(globalState))), d_2083_v1_)
                        nw335_[int(4)] = _dafny.SeqWithoutIsStrInference([d_2083_v1_ for d_2097_i1_ in range(default__.abs(-717))])
                        nw335_[int(5)] = d_2089_v7_
                        nw335_[int(6)] = _dafny.SeqWithoutIsStrInference([d_2083_v1_ for d_2098_i2_ in range(default__.abs(343))])
                        nw335_[int(7)] = (d_2091_v9_).fm28(d_2094_v12_, globalState)
                        nw335_[int(8)] = d_2089_v7_
                        nw335_[int(9)] = ((d_2089_v7_).set(default__.safeIndex(d_2085_v3_, len(d_2089_v7_)), _dafny.CodePoint('n'))).set(default__.safeIndex(p0, len((d_2089_v7_).set(default__.safeIndex(d_2085_v3_, len(d_2089_v7_)), _dafny.CodePoint('n')))), _dafny.CodePoint('x'))
                        nw335_[int(10)] = d_2089_v7_
                        nw335_[int(11)] = d_2089_v7_
                        nw335_[int(12)] = d_2089_v7_
                        nw335_[int(13)] = d_2089_v7_
                        nw335_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jc"))
                        nw335_[int(15)] = d_2089_v7_
                        nw335_[int(16)] = d_2089_v7_
                        nw335_[int(17)] = d_2089_v7_
                        nw335_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "befkyx"))
                        nw335_[int(19)] = (d_2095_v13_)[default__.safeIndex(d_2087_v5_, len(d_2095_v13_))]
                        nw335_[int(20)] = d_2089_v7_
                        d_2096_v14_ = nw335_
                        d_2099_v15_: _dafny.Array
                        nw336_ = _dafny.Array(False, 10)
                        d_2099_v15_ = nw336_
                        d_2100_v16_: _dafny.Seq
                        d_2100_v16_ = _dafny.SeqWithoutIsStrInference([d_2099_v15_, d_2099_v15_])
                        d_2101_v17_: _dafny.Map
                        d_2101_v17_ = _dafny.Map({d_2096_v14_: d_2100_v16_})
                        d_2101_v17_ = (d_2101_v17_).set(d_2096_v14_, (d_2100_v16_) + (d_2100_v16_))
                    elif True:
                        d_2102_v18_: _dafny.Map
                        d_2102_v18_ = _dafny.Map({(0) - ((0) - (d_2087_v5_)): d_2087_v5_})
                        (globalState).f4 = ((d_2102_v18_)[(d_2087_v5_) + (d_2085_v3_)] if ((d_2087_v5_) + (d_2085_v3_)) in (d_2102_v18_) else p0)
                        d_2103_v19_: _dafny.Set
                        d_2103_v19_ = _dafny.Set({not(p1)})
                        d_2104_v20_: _dafny.Array
                        nw337_ = _dafny.Array(None, 8)
                        nw337_[int(0)] = d_2082_v0_
                        nw337_[int(1)] = d_2082_v0_
                        nw337_[int(2)] = p1
                        nw337_[int(3)] = (d_2103_v19_).isdisjoint(d_2103_v19_)
                        nw337_[int(4)] = d_2082_v0_
                        nw337_[int(5)] = p3
                        nw337_[int(6)] = p3
                        nw337_[int(7)] = p1
                        d_2104_v20_ = nw337_
                        index369_ = default__.safeIndex(320, (d_2104_v20_).length(0))
                        (d_2104_v20_)[index369_] = d_2082_v0_
                        index370_ = default__.safeIndex(320, (d_2104_v20_).length(0))
                        (d_2104_v20_)[index370_] = p1
                        d_2105_v21_: _dafny.Map
                        d_2105_v21_ = _dafny.Map({d_2082_v0_: not(p3)})
                        d_2082_v0_ = not(not((self).fm1(_dafny.Map({p1: d_2105_v21_}), p0, globalState)))
                        d_2085_v3_ = p0
                        d_2106_v23_: _dafny.Seq
                        d_2106_v23_ = _dafny.SeqWithoutIsStrInference([d_2087_v5_])
                        def iife159_():
                            coll81_ = _dafny.Map()
                            compr_81_: int
                            for compr_81_ in (d_2106_v23_).Elements:
                                d_2107_v22_: int = compr_81_
                                if (d_2107_v22_) in (d_2106_v23_):
                                    coll81_[(d_2107_v22_) + (d_2087_v5_)] = d_2085_v3_
                            return _dafny.Map(coll81_)
                        d_2102_v18_ = ((d_2102_v18_) | (iife159_()
                        )) | (d_2102_v18_)
                    d_2108_v24_: _dafny.Array
                    nw338_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                    d_2108_v24_ = nw338_
                    d_2109_v25_: _dafny.Seq
                    d_2109_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihvfprud"))
                    index371_ = default__.safeIndex(198, (d_2108_v24_).length(0))
                    (d_2108_v24_)[index371_] = d_2109_v25_
                    index372_ = default__.safeIndex(198, (d_2108_v24_).length(0))
                    (d_2108_v24_)[index372_] = (d_2109_v25_) + (d_2109_v25_)
                    pass
            pass
        d_2110_v26_: _dafny.Map
        d_2110_v26_ = _dafny.Map({p3: p3})
        d_2111_v27_: _dafny.Map
        d_2111_v27_ = _dafny.Map({p1: d_2110_v26_})
        d_2112_v28_: _dafny.Seq
        d_2112_v28_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
        d_2113_v29_: _dafny.Map
        d_2113_v29_ = _dafny.Map({p3: (self).fm1(d_2111_v27_, len(d_2112_v28_), globalState)})
        d_2114_v30_: _dafny.Map
        d_2114_v30_ = _dafny.Map({p3: d_2110_v26_})
        d_2115_v31_: _dafny.Seq
        d_2115_v31_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2116_v32_: _dafny.MultiSet
        d_2116_v32_ = _dafny.MultiSet([d_2115_v31_])
        d_2117_v33_: _dafny.Set
        d_2117_v33_ = _dafny.Set({(self).fm1(d_2111_v27_, (self).fm2((d_2116_v32_).cardinality, globalState), globalState)})
        d_2118_v34_: str
        d_2118_v34_ = _dafny.CodePoint('l')
        d_2119_v35_: _dafny.Map
        d_2119_v35_ = _dafny.Map({d_2118_v34_: _dafny.Set({p1})})
        pat_let_tv44_ = d_2119_v35_
        pat_let_tv45_ = d_2118_v34_
        pat_let_tv46_ = d_2118_v34_
        pat_let_tv47_ = d_2119_v35_
        pat_let_tv48_ = d_2117_v33_
        def iife160_(_pat_let39_0):
            def iife161_(d_2120_dt__update__tmp_h0_):
                def iife162_(_pat_let40_0):
                    def iife163_(d_2121_dt__update_hcf49_h0_):
                        return D12_DC32(d_2121_dt__update_hcf49_h0_)
                    return iife163_(_pat_let40_0)
                return iife162_(((pat_let_tv44_)[pat_let_tv45_] if (pat_let_tv46_) in (pat_let_tv47_) else pat_let_tv48_))
            return iife161_(_pat_let39_0)
        source30_ = iife160_(D12_DC32(d_2117_v33_))
        if source30_.is_DC33:
            d_2122___mcc_h0_ = source30_.cf50
            d_2123_cf50_ = d_2122___mcc_h0_
            (globalState).f13 = 461
            d_2124_v36_: D3
            d_2124_v36_ = D3_DC10(468, p0)
            source31_ = d_2124_v36_
            if source31_.is_DC10:
                d_2125___mcc_h2_ = source31_.cf13
                d_2126___mcc_h3_ = source31_.cf14
                d_2127_cf14_ = d_2126___mcc_h3_
                d_2128_cf13_ = d_2125___mcc_h2_
                (globalState).f4 = p0
                d_2129_v37_: _dafny.Seq
                d_2129_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmyrhykkq"))
                d_2130_v38_: _dafny.Map
                d_2130_v38_ = _dafny.Map({d_2127_cf14_: False})
                d_2131_v39_: _dafny.Array
                def lambda108_(d_2132_cf14_):
                    def lambda109_(d_2133_i3_):
                        return (d_2133_i3_) - (len(_dafny.SeqWithoutIsStrInference([d_2132_cf14_ for d_2134_i4_ in range(default__.abs(78))])))

                    return lambda109_

                init55_ = lambda108_(d_2127_cf14_)
                nw339_ = _dafny.Array(None, 7)
                for i0_55_ in range(nw339_.length(0)):
                    nw339_[i0_55_] = init55_(i0_55_)
                d_2131_v39_ = nw339_
                d_2135_v40_: D2
                d_2135_v40_ = D2_DC5(d_2131_v39_)
                d_2136_v41_: _dafny.Set
                d_2136_v41_ = _dafny.Set({d_2135_v40_})
                rhs434_ = default__.fm26(globalState)
                rhs435_ = not ((d_2123_cf50_ if p1 else not(p3))) or (((d_2130_v38_)[len(d_2117_v33_)] if (len(d_2117_v33_)) in (d_2130_v38_) else p3))
                rhs436_ = ((d_2118_v34_ if (d_2115_v31_)[default__.safeIndex(29, len(d_2115_v31_))] else d_2118_v34_)) not in (default__.fm19((d_2115_v31_)[default__.safeIndex(len(d_2136_v41_), len(d_2115_v31_))], d_2123_cf50_, globalState))
                rhs437_ = d_2110_v26_
                rhs438_ = (d_2129_v37_) + (d_2129_v37_)
                d_2123_cf50_ = rhs434_
                d_2123_cf50_ = rhs435_
                d_2123_cf50_ = rhs436_
                d_2110_v26_ = rhs437_
                d_2129_v37_ = rhs438_
                d_2137_v42_: _dafny.Array
                nw340_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_2137_v42_ = nw340_
                rhs439_ = d_2137_v42_
                rhs440_ = (not (p1) or (p3) if p3 else d_2123_cf50_)
                rhs441_ = _dafny.SeqWithoutIsStrInference([(d_2128_cf13_) - (len(d_2130_v38_)) for d_2138_i5_ in range(default__.abs(-409))])
                rhs442_ = d_2118_v34_
                rhs443_ = p3
                d_2137_v42_ = rhs439_
                d_2123_cf50_ = rhs440_
                d_2112_v28_ = rhs441_
                d_2118_v34_ = rhs442_
                d_2123_cf50_ = rhs443_
                d_2129_v37_ = d_2129_v37_
            elif True:
                d_2139___mcc_h4_ = source31_.cf12
                d_2140_cf12_ = d_2139___mcc_h4_
                d_2123_cf50_ = default__.fm26(globalState)
                d_2141_v43_: C1
                nw341_ = C1()
                nw341_.ctor__()
                d_2141_v43_ = nw341_
                d_2142_v44_: _dafny.Seq
                d_2142_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qu"))
                d_2142_v44_ = d_2142_v44_
                d_2123_cf50_ = (d_2117_v33_).ispropersubset(_dafny.Set({False}))
            d_2143_v45_: _dafny.Map
            d_2143_v45_ = _dafny.Map({(self).fm2(p0, globalState): p0})
            d_2143_v45_ = (d_2143_v45_).set(default__.safeModuloInt((d_2112_v28_)[default__.safeIndex(p0, len(d_2112_v28_))], p0), (0) - ((self).fm2(p0, globalState)))
            if False:
                d_2123_cf50_ = not(True)
                d_2144_v46_: _dafny.Seq
                d_2144_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auf"))
                d_2144_v46_ = d_2144_v46_
                (globalState).f13 = p0
                d_2145_v47_: C10
                nw342_ = C10()
                nw342_.ctor__()
                d_2145_v47_ = nw342_
                d_2146_v48_: _dafny.MultiSet
                d_2146_v48_ = _dafny.MultiSet([388])
                (globalState).f7 = default__.safeModuloInt(p0, (p0) + (((d_2146_v48_)[p0] if (p0) in (d_2146_v48_) else p0)))
            elif True:
                d_2147_v49_: _dafny.MultiSet
                d_2147_v49_ = _dafny.MultiSet([not(p3)])
                d_2148_v50_: _dafny.MultiSet
                d_2148_v50_ = _dafny.MultiSet([(d_2147_v49_).set(d_2123_cf50_, default__.abs(p0)), _dafny.MultiSet([p3, p3, p3, d_2123_cf50_])])
                d_2149_v51_: _dafny.Array
                nw343_ = _dafny.Array(None, 4)
                nw343_[int(0)] = d_2123_cf50_
                nw343_[int(1)] = (d_2148_v50_).isdisjoint(d_2148_v50_)
                nw343_[int(2)] = p1
                nw343_[int(3)] = p1
                d_2149_v51_ = nw343_
                d_2150_v52_: D6
                d_2150_v52_ = D6_DC16(_dafny.SeqWithoutIsStrInference([p0]))
                index373_ = default__.safeIndex(85, (d_2149_v51_).length(0))
                (d_2149_v51_)[index373_] = ((d_2150_v52_).cf23) == (d_2112_v28_)
                index374_ = default__.safeIndex(85, (d_2149_v51_).length(0))
                (d_2149_v51_)[index374_] = p3
                d_2151_v53_: _dafny.Set
                d_2151_v53_ = _dafny.Set({(len(d_2112_v28_)) + (p0), 655})
                d_2151_v53_ = ((d_2151_v53_) | (d_2151_v53_)).intersection(d_2151_v53_)
                d_2152_v54_: D7
                d_2152_v54_ = D7_DC22(p3, p0, (d_2149_v51_)[default__.safeIndex(85, (d_2149_v51_).length(0))], p0)
                d_2153_v55_: _dafny.MultiSet
                d_2153_v55_ = _dafny.MultiSet([d_2152_v54_, D7_DC22(p3, p0, p1, p0), d_2152_v54_, D7_DC22(p3, p0, True, p0), d_2152_v54_])
                d_2154_v56_: D11
                d_2154_v56_ = D11_DC30(d_2123_cf50_, p0)
                d_2155_v57_: _dafny.Set
                d_2155_v57_ = _dafny.Set({default__.fm38(d_2123_cf50_, globalState)})
                d_2156_v58_: _dafny.Map
                d_2156_v58_ = _dafny.Map({d_2118_v34_: d_2118_v34_})
                d_2157_v60_: D5
                d_2157_v60_ = D5_DC13(default__.fm43(p0, globalState))
                d_2158_v61_: _dafny.Map
                d_2158_v61_ = _dafny.Map({p0: d_2157_v60_})
                index375_ = default__.safeIndex(85, (d_2149_v51_).length(0))
                def iife164_():
                    coll82_ = _dafny.Set()
                    compr_82_: str
                    for compr_82_ in (d_2156_v58_).keys.Elements:
                        d_2159_v59_: str = compr_82_
                        if (d_2159_v59_) in (d_2156_v58_):
                            coll82_ = coll82_.union(_dafny.Set([d_2159_v59_]))
                    return _dafny.Set(coll82_)
                rhs444_ = ((d_2155_v57_) - (iife164_()
                )).ispropersubset(_dafny.Set({d_2118_v34_}))
                rhs445_ = default__.fm57(p0, d_2158_v61_, globalState)
                rhs446_ = d_2154_v56_
                lhs301_ = d_2149_v51_
                lhs302_ = default__.safeIndex(85, (d_2149_v51_).length(0))
                lhs301_[lhs302_] = rhs444_
                d_2153_v55_ = rhs445_
                d_2154_v56_ = rhs446_
                d_2160_v62_: _dafny.Array
                nw344_ = _dafny.Array(_dafny.Map({}), 1)
                d_2160_v62_ = nw344_
                d_2161_v63_: _dafny.Map
                d_2161_v63_ = _dafny.Map({d_2118_v34_: p0})
                d_2162_v64_: _dafny.Seq
                d_2162_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhuvugk"))
                d_2163_v65_: C6
                nw345_ = C6()
                nw345_.ctor__()
                d_2163_v65_ = nw345_
                d_2164_v66_: _dafny.Map
                d_2164_v66_ = _dafny.Map({788: d_2163_v65_})
                d_2165_v67_: _dafny.Array
                nw346_ = _dafny.Array(None, 24)
                nw346_[int(0)] = ((d_2143_v45_)[len(d_2161_v63_)] if (len(d_2161_v63_)) in (d_2143_v45_) else p0)
                nw346_[int(1)] = p0
                nw346_[int(2)] = p0
                nw346_[int(3)] = p0
                nw346_[int(4)] = p0
                nw346_[int(5)] = len(d_2162_v64_)
                nw346_[int(6)] = p0
                nw346_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le")))
                nw346_[int(8)] = p0
                nw346_[int(9)] = p0
                nw346_[int(10)] = p0
                nw346_[int(11)] = len(d_2164_v66_)
                nw346_[int(12)] = len(d_2162_v64_)
                nw346_[int(13)] = p0
                nw346_[int(14)] = 720
                nw346_[int(15)] = p0
                nw346_[int(16)] = p0
                nw346_[int(17)] = p0
                nw346_[int(18)] = p0
                nw346_[int(19)] = -383
                nw346_[int(20)] = p0
                nw346_[int(21)] = p0
                nw346_[int(22)] = p0
                nw346_[int(23)] = p0
                d_2165_v67_ = nw346_
                d_2166_v68_: _dafny.Map
                d_2166_v68_ = _dafny.Map({D2_DC5(d_2165_v67_): d_2123_cf50_})
                index376_ = default__.safeIndex(942, (d_2160_v62_).length(0))
                (d_2160_v62_)[index376_] = d_2166_v68_
                index377_ = default__.safeIndex(942, (d_2160_v62_).length(0))
                (d_2160_v62_)[index377_] = d_2166_v68_
                d_2167_v69_: _dafny.Array
                nw347_ = _dafny.Array(None, 25)
                nw347_[int(0)] = d_2115_v31_
                nw347_[int(1)] = d_2115_v31_
                nw347_[int(2)] = d_2115_v31_
                nw347_[int(3)] = d_2115_v31_
                nw347_[int(4)] = d_2115_v31_
                nw347_[int(5)] = d_2115_v31_
                nw347_[int(6)] = d_2115_v31_
                nw347_[int(7)] = _dafny.SeqWithoutIsStrInference([(d_2149_v51_)[default__.safeIndex(85, (d_2149_v51_).length(0))]])
                nw347_[int(8)] = d_2115_v31_
                nw347_[int(9)] = d_2115_v31_
                nw347_[int(10)] = d_2115_v31_
                nw347_[int(11)] = (d_2115_v31_).set(default__.safeIndex(p0, len(d_2115_v31_)), (d_2149_v51_)[default__.safeIndex(85, (d_2149_v51_).length(0))])
                nw347_[int(12)] = _dafny.SeqWithoutIsStrInference([p3])
                nw347_[int(13)] = default__.fm27(globalState)
                nw347_[int(14)] = d_2115_v31_
                nw347_[int(15)] = d_2115_v31_
                nw347_[int(16)] = d_2115_v31_
                nw347_[int(17)] = d_2115_v31_
                nw347_[int(18)] = d_2115_v31_
                nw347_[int(19)] = d_2115_v31_
                nw347_[int(20)] = d_2115_v31_
                nw347_[int(21)] = d_2115_v31_
                nw347_[int(22)] = _dafny.SeqWithoutIsStrInference([d_2123_cf50_, True])
                nw347_[int(23)] = d_2115_v31_
                nw347_[int(24)] = d_2115_v31_
                d_2167_v69_ = nw347_
                d_2168_v70_: _dafny.Array
                d_2168_v70_ = d_2167_v69_
                d_2169_v71_: _dafny.Seq
                d_2169_v71_ = _dafny.SeqWithoutIsStrInference([d_2165_v67_, d_2165_v67_, d_2165_v67_, d_2165_v67_, d_2165_v67_])
                d_2170_v72_: _dafny.Seq
                d_2170_v72_ = _dafny.SeqWithoutIsStrInference([(d_2169_v71_)[default__.safeIndex(-155, len(d_2169_v71_))]])
                d_2171_v73_: D2
                d_2171_v73_ = D2_DC5(d_2165_v67_)
                index378_ = default__.safeIndex(85, (d_2149_v51_).length(0))
                rhs447_ = d_2168_v70_
                rhs448_ = (_dafny.SeqWithoutIsStrInference([d_2165_v67_, (d_2171_v73_).cf6])) + (d_2169_v71_)
                rhs449_ = (868) < (p0)
                rhs450_ = (d_2163_v65_).fm16(p0, globalState)
                rhs451_ = p1
                lhs303_ = d_2149_v51_
                lhs304_ = default__.safeIndex(85, (d_2149_v51_).length(0))
                d_2168_v70_ = rhs447_
                d_2170_v72_ = rhs448_
                lhs303_[lhs304_] = rhs449_
                d_2123_cf50_ = rhs450_
                d_2123_cf50_ = rhs451_
        elif True:
            d_2172___mcc_h1_ = source30_.cf49
            d_2173_cf49_ = d_2172___mcc_h1_
            d_2174_v74_: _dafny.Array
            nw348_ = _dafny.Array(None, 11)
            nw348_[int(0)] = d_2118_v34_
            nw348_[int(1)] = default__.fm38(p1, globalState)
            nw348_[int(2)] = d_2118_v34_
            nw348_[int(3)] = d_2118_v34_
            nw348_[int(4)] = _dafny.CodePoint('i')
            nw348_[int(5)] = d_2118_v34_
            nw348_[int(6)] = d_2118_v34_
            nw348_[int(7)] = d_2118_v34_
            nw348_[int(8)] = _dafny.CodePoint('p')
            nw348_[int(9)] = d_2118_v34_
            nw348_[int(10)] = d_2118_v34_
            d_2174_v74_ = nw348_
            d_2174_v74_ = d_2174_v74_
            d_2175_v75_: _dafny.Map
            d_2175_v75_ = _dafny.Map({p0: p1})
            d_2176_v76_: T0
            nw349_ = C9()
            nw349_.ctor__()
            d_2176_v76_ = nw349_
            d_2177_v77_: D17
            d_2177_v77_ = D17_DC47(d_2176_v76_, p3, p0, p0, p0)
            d_2178_v78_: D18
            d_2178_v78_ = D18_DC49(p1, True, default__.fm38(p3, globalState), (d_2177_v77_).cf70)
            d_2179_v79_: _dafny.Array
            nw350_ = _dafny.Array(None, 18)
            nw350_[int(0)] = p3
            nw350_[int(1)] = not(p3)
            nw350_[int(2)] = False
            nw350_[int(3)] = p1
            nw350_[int(4)] = p3
            nw350_[int(5)] = p1
            nw350_[int(6)] = p1
            nw350_[int(7)] = p1
            nw350_[int(8)] = p3
            nw350_[int(9)] = ((d_2175_v75_)[627] if (627) in (d_2175_v75_) else not(False))
            nw350_[int(10)] = p1
            nw350_[int(11)] = not((d_2178_v78_).cf78)
            nw350_[int(12)] = False
            nw350_[int(13)] = p1
            nw350_[int(14)] = p3
            nw350_[int(15)] = p3
            nw350_[int(16)] = p1
            nw350_[int(17)] = p3
            d_2179_v79_ = nw350_
            d_2180_v80_: _dafny.Map
            d_2180_v80_ = _dafny.Map({595: d_2179_v79_})
            d_2180_v80_ = (d_2180_v80_).set(p0, d_2179_v79_)
            d_2181_v81_: C8
            nw351_ = C8()
            nw351_.ctor__()
            d_2181_v81_ = nw351_
            d_2182_v82_: _dafny.Array
            nw352_ = _dafny.Array(None, 15)
            nw352_[int(0)] = d_2179_v79_
            nw352_[int(1)] = d_2179_v79_
            nw352_[int(2)] = d_2179_v79_
            nw352_[int(3)] = d_2179_v79_
            nw352_[int(4)] = d_2179_v79_
            nw352_[int(5)] = d_2179_v79_
            nw352_[int(6)] = d_2179_v79_
            nw352_[int(7)] = d_2179_v79_
            nw352_[int(8)] = d_2179_v79_
            nw352_[int(9)] = d_2179_v79_
            nw352_[int(10)] = d_2179_v79_
            nw352_[int(11)] = d_2179_v79_
            nw352_[int(12)] = d_2179_v79_
            nw352_[int(13)] = d_2179_v79_
            nw352_[int(14)] = d_2179_v79_
            d_2182_v82_ = nw352_
            d_2182_v82_ = d_2182_v82_
        d_2183_v83_: C10
        nw353_ = C10()
        nw353_.ctor__()
        d_2183_v83_ = nw353_
        d_2184_v84_: bool
        d_2184_v84_ = False
        d_2184_v84_ = d_2184_v84_
        d_2185_v85_: D5
        d_2185_v85_ = D5_DC14(d_2118_v34_, p1, p0, p1)
        d_2186_v86_: _dafny.Seq
        d_2186_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juct"))
        d_2187_v87_: _dafny.Array
        nw354_ = _dafny.Array(None, 5)
        nw354_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_2185_v85_).cf18, (d_2186_v86_)[default__.safeIndex(p0, len(d_2186_v86_))]])
        nw354_[int(1)] = d_2186_v86_
        nw354_[int(2)] = default__.fm5(globalState)
        nw354_[int(3)] = d_2186_v86_
        nw354_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_2118_v34_, (d_2186_v86_)[default__.safeIndex(p0, len(d_2186_v86_))]])).set(default__.safeIndex(len(d_2115_v31_), len(_dafny.SeqWithoutIsStrInference([d_2118_v34_, (d_2186_v86_)[default__.safeIndex(p0, len(d_2186_v86_))]]))), d_2118_v34_)
        d_2187_v87_ = nw354_
        d_2188_v88_: D21
        d_2188_v88_ = D21_DC56(d_2187_v87_)
        source32_ = d_2188_v88_
        if source32_.is_DC57:
            d_2189___mcc_h5_ = source32_.cf90
            d_2190___mcc_h6_ = source32_.cf91
            d_2191___mcc_h7_ = source32_.cf92
            d_2192_cf92_ = d_2191___mcc_h7_
            d_2193_cf91_ = d_2190___mcc_h6_
            d_2194_cf90_ = d_2189___mcc_h5_
            d_2195_v89_: _dafny.Map
            d_2195_v89_ = _dafny.Map({d_2184_v84_: 892})
            d_2196_v90_: _dafny.Map
            d_2196_v90_ = _dafny.Map({((d_2195_v89_)[p3] if (p3) in (d_2195_v89_) else d_2194_cf90_): p1})
            d_2196_v90_ = (default__.fm35(globalState)) | (_dafny.Map({d_2194_cf90_: p1}))
            d_2118_v34_ = _dafny.CodePoint('r')
            if d_2184_v84_:
                d_2197_v91_: T1
                nw355_ = C8()
                nw355_.ctor__()
                d_2197_v91_ = nw355_
                d_2198_v92_: _dafny.MultiSet
                d_2198_v92_ = _dafny.MultiSet([p3, p1])
                rhs452_ = d_2197_v91_
                rhs453_ = p1
                rhs454_ = d_2198_v92_
                d_2197_v91_ = rhs452_
                d_2184_v84_ = rhs453_
                d_2198_v92_ = rhs454_
                d_2184_v84_ = d_2184_v84_
                d_2199_v93_: _dafny.MultiSet
                d_2199_v93_ = _dafny.MultiSet([(d_2183_v83_).fm2(len(d_2196_v90_), globalState)])
                d_2200_v94_: _dafny.Set
                d_2200_v94_ = _dafny.Set({(d_2199_v93_).cardinality})
                d_2201_v95_: _dafny.Array
                nw356_ = _dafny.Array(None, 2)
                nw356_[int(0)] = d_2184_v84_
                nw356_[int(1)] = (d_2183_v83_).fm1(d_2114_v30_, len(d_2200_v94_), globalState)
                d_2201_v95_ = nw356_
                d_2202_v96_: _dafny.Seq
                d_2202_v96_ = _dafny.SeqWithoutIsStrInference([d_2201_v95_])
                d_2203_v97_: _dafny.Map
                d_2203_v97_ = _dafny.Map({d_2201_v95_: d_2194_cf90_})
                rhs455_ = d_2194_cf90_
                rhs456_ = d_2202_v96_
                rhs457_ = (d_2201_v95_) not in (d_2203_v97_)
                rhs458_ = ((0) - (p0)) + (d_2193_cf91_)
                lhs305_ = globalState
                lhs306_ = globalState
                lhs305_.f7 = rhs455_
                d_2202_v96_ = rhs456_
                d_2184_v84_ = rhs457_
                lhs306_.f15 = rhs458_
                d_2204_v98_: C9
                nw357_ = C9()
                nw357_.ctor__()
                d_2204_v98_ = nw357_
                d_2205_v99_: C6
                nw358_ = C6()
                nw358_.ctor__()
                d_2205_v99_ = nw358_
                d_2206_v100_: _dafny.Map
                d_2206_v100_ = _dafny.Map({(d_2194_cf90_ if p1 else len(d_2186_v86_)): d_2205_v99_})
                (globalState).f13 = len(d_2206_v100_)
            elif True:
                d_2207_v101_: _dafny.MultiSet
                d_2207_v101_ = _dafny.MultiSet([len(_dafny.Map({len(d_2117_v33_): False}))])
                d_2112_v28_ = (_dafny.SeqWithoutIsStrInference([d_2194_cf90_, (d_2207_v101_).cardinality, p0])) + (d_2112_v28_)
                d_2208_v102_: D6
                d_2208_v102_ = D6_DC16(_dafny.SeqWithoutIsStrInference([p0 for d_2209_i6_ in range(default__.abs(494))]))
                d_2210_v103_: D6
                d_2210_v103_ = D6_DC18(d_2208_v102_)
                d_2211_v104_: D6
                d_2211_v104_ = D6_DC18(d_2210_v103_)
                d_2212_v105_: D6
                d_2212_v105_ = D6_DC18(d_2210_v103_)
                d_2213_v106_: D6
                d_2213_v106_ = D6_DC18(d_2208_v102_)
                d_2214_v107_: D6
                d_2214_v107_ = D6_DC18(d_2210_v103_)
                d_2214_v107_ = D6_DC18(D6_DC18(d_2210_v103_))
                d_2215_v108_: _dafny.Map
                d_2215_v108_ = _dafny.Map({d_2184_v84_: _dafny.Map({(0) - (d_2194_cf90_): p1})})
                d_2216_v109_: _dafny.Map
                d_2216_v109_ = _dafny.Map({p0: len(d_2215_v108_)})
                d_2216_v109_ = (d_2216_v109_).set(p0, default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2217_i7_ in range(default__.abs(255))]))))
                d_2184_v84_ = (d_2184_v84_) and (p3)
                d_2117_v33_ = (d_2117_v33_) | (d_2117_v33_)
            if not (p3) or (not(d_2184_v84_)):
                d_2218_v111_: _dafny.Set
                d_2218_v111_ = _dafny.Set({d_2193_cf91_})
                d_2219_v112_: _dafny.MultiSet
                def iife165_():
                    coll83_ = _dafny.Map()
                    compr_83_: int
                    for compr_83_ in (d_2218_v111_).Elements:
                        d_2220_v110_: int = compr_83_
                        if (d_2220_v110_) in (d_2218_v111_):
                            coll83_[(d_2220_v110_) * (len(d_2186_v86_))] = D21_DC57(d_2194_cf90_, d_2193_cf91_, d_2192_cf92_)
                    return _dafny.Map(coll83_)
                d_2219_v112_ = _dafny.MultiSet([iife165_()
                ])
                d_2221_v113_: _dafny.MultiSet
                d_2221_v113_ = _dafny.MultiSet([p0])
                d_2222_v114_: D21
                d_2222_v114_ = D21_DC57((d_2221_v113_).cardinality, d_2193_cf91_, D5_DC14(d_2118_v34_, p1, p0, p1))
                d_2223_v115_: _dafny.Map
                d_2223_v115_ = _dafny.Map({d_2193_cf91_: d_2222_v114_})
                d_2224_v116_: _dafny.Seq
                d_2224_v116_ = _dafny.SeqWithoutIsStrInference([d_2223_v115_])
                d_2225_v117_: _dafny.Seq
                d_2225_v117_ = _dafny.SeqWithoutIsStrInference([d_2219_v112_, _dafny.MultiSet(d_2224_v116_)])
                d_2226_v118_: _dafny.Seq
                d_2226_v118_ = _dafny.SeqWithoutIsStrInference([d_2219_v112_, (d_2225_v117_)[default__.safeIndex(d_2193_cf91_, len(d_2225_v117_))]])
                d_2227_v119_: _dafny.Seq
                d_2227_v119_ = _dafny.SeqWithoutIsStrInference([(d_2225_v117_)[default__.safeIndex(p0, len(d_2225_v117_))]])
                d_2228_v120_: _dafny.Set
                d_2228_v120_ = _dafny.Set({d_2118_v34_, d_2118_v34_, d_2118_v34_})
                d_2229_v121_: D22
                d_2229_v121_ = D22_DC59(d_2219_v112_)
                pat_let_tv49_ = d_2219_v112_
                d_2230_v122_: _dafny.Array
                nw359_ = _dafny.Array(None, 29)
                nw359_[int(0)] = d_2219_v112_
                nw359_[int(1)] = d_2219_v112_
                nw359_[int(2)] = d_2219_v112_
                nw359_[int(3)] = d_2219_v112_
                nw359_[int(4)] = d_2219_v112_
                nw359_[int(5)] = d_2219_v112_
                nw359_[int(6)] = d_2219_v112_
                nw359_[int(7)] = (d_2219_v112_) - ((d_2226_v118_)[default__.safeIndex(p0, len(d_2226_v118_))])
                nw359_[int(8)] = d_2219_v112_
                nw359_[int(9)] = d_2219_v112_
                nw359_[int(10)] = (_dafny.MultiSet((d_2224_v116_).set(default__.safeIndex((d_2183_v83_).fm6(p1, p3, _dafny.Map({-245: p0}), default__.fm38(p3, globalState), globalState), len(d_2224_v116_)), d_2223_v115_))) | (d_2219_v112_)
                nw359_[int(11)] = (d_2225_v117_)[default__.safeIndex(780, len(d_2225_v117_))]
                nw359_[int(12)] = d_2219_v112_
                nw359_[int(13)] = (d_2219_v112_).intersection(d_2219_v112_)
                nw359_[int(14)] = d_2219_v112_
                nw359_[int(15)] = d_2219_v112_
                nw359_[int(16)] = d_2219_v112_
                nw359_[int(17)] = d_2219_v112_
                nw359_[int(18)] = d_2219_v112_
                nw359_[int(19)] = d_2219_v112_
                nw359_[int(20)] = (d_2227_v119_)[default__.safeIndex(231, len(d_2227_v119_))]
                nw359_[int(21)] = d_2219_v112_
                nw359_[int(22)] = _dafny.MultiSet([(d_2224_v116_)[default__.safeIndex(len(d_2228_v120_), len(d_2224_v116_))], default__.fm58(d_2184_v84_, d_2118_v34_, globalState)])
                nw359_[int(23)] = d_2219_v112_
                def iife166_(_pat_let41_0):
                    def iife167_(d_2231_dt__update__tmp_h1_):
                        def iife168_(_pat_let42_0):
                            def iife169_(d_2232_dt__update_hcf94_h0_):
                                return D22_DC59(d_2232_dt__update_hcf94_h0_)
                            return iife169_(_pat_let42_0)
                        return iife168_(pat_let_tv49_)
                    return iife167_(_pat_let41_0)
                nw359_[int(24)] = (iife166_(d_2229_v121_)).cf94
                nw359_[int(25)] = d_2219_v112_
                nw359_[int(26)] = d_2219_v112_
                nw359_[int(27)] = d_2219_v112_
                nw359_[int(28)] = (d_2219_v112_).intersection(d_2219_v112_)
                d_2230_v122_ = nw359_
                index379_ = default__.safeIndex(816, (d_2230_v122_).length(0))
                (d_2230_v122_)[index379_] = d_2219_v112_
                index380_ = default__.safeIndex(816, (d_2230_v122_).length(0))
                (d_2230_v122_)[index380_] = d_2219_v112_
                d_2184_v84_ = d_2184_v84_
                d_2184_v84_ = d_2184_v84_
                d_2233_v123_: _dafny.Array
                def lambda110_(d_2234_cf91_):
                    def lambda111_(d_2235_i8_):
                        return (d_2235_i8_) + ((0) - (d_2234_cf91_))

                    return lambda111_

                init56_ = lambda110_(d_2193_cf91_)
                nw360_ = _dafny.Array(None, 28)
                for i0_56_ in range(nw360_.length(0)):
                    nw360_[i0_56_] = init56_(i0_56_)
                d_2233_v123_ = nw360_
                index381_ = default__.safeIndex(237, (d_2233_v123_).length(0))
                (d_2233_v123_)[index381_] = p0
                d_2236_v124_: _dafny.MultiSet
                d_2236_v124_ = _dafny.MultiSet([p1, not(d_2184_v84_), False])
                d_2237_v125_: _dafny.Map
                d_2237_v125_ = _dafny.Map({d_2186_v86_: d_2194_cf90_})
                index382_ = default__.safeIndex(237, (d_2233_v123_).length(0))
                rhs459_ = default__.safeDivisionInt(d_2193_cf91_, (d_2221_v113_).cardinality)
                rhs460_ = (d_2236_v124_).cardinality
                rhs461_ = ((d_2237_v125_)[d_2186_v86_] if (d_2186_v86_) in (d_2237_v125_) else default__.safeDivisionInt(p0, len(d_2186_v86_)))
                lhs307_ = globalState
                lhs308_ = d_2233_v123_
                lhs309_ = default__.safeIndex(237, (d_2233_v123_).length(0))
                lhs310_ = globalState
                lhs307_.f7 = rhs459_
                lhs308_[lhs309_] = rhs460_
                lhs310_.f13 = rhs461_
                d_2238_v126_: _dafny.Seq
                d_2238_v126_ = _dafny.SeqWithoutIsStrInference([d_2218_v111_, d_2218_v111_, (d_2218_v111_) - (_dafny.Set({-54, 672, d_2194_cf90_, d_2193_cf91_}))])
                d_2238_v126_ = ((d_2238_v126_) + (d_2238_v126_)) + (d_2238_v126_)
            elif True:
                d_2239_v127_: _dafny.Array
                def lambda112_(d_2240_v34_):
                    def lambda113_(d_2241_i9_):
                        return d_2240_v34_

                    return lambda113_

                init57_ = lambda112_(d_2118_v34_)
                nw361_ = _dafny.Array(None, 5)
                for i0_57_ in range(nw361_.length(0)):
                    nw361_[i0_57_] = init57_(i0_57_)
                d_2239_v127_ = nw361_
                d_2242_v128_: _dafny.Map
                d_2242_v128_ = _dafny.Map({p1: d_2239_v127_})
                d_2242_v128_ = (d_2242_v128_).set(p1, d_2239_v127_)
                d_2243_v129_: _dafny.Map
                d_2243_v129_ = _dafny.Map({default__.safeModuloInt(d_2194_cf90_, d_2194_cf90_): (d_2112_v28_) + (d_2112_v28_)})
                d_2243_v129_ = (d_2243_v129_).set((d_2194_cf90_) * (p0), d_2112_v28_)
                d_2184_v84_ = p1
                d_2244_v130_: C6
                nw362_ = C6()
                nw362_.ctor__()
                d_2244_v130_ = nw362_
                d_2112_v28_ = (d_2112_v28_) + (d_2112_v28_)
        elif source32_.is_DC56:
            d_2245___mcc_h8_ = source32_.cf89
            d_2246_cf89_ = d_2245___mcc_h8_
            (globalState).f7 = p0
            d_2184_v84_ = p3
            d_2247_v131_: C9
            nw363_ = C9()
            nw363_.ctor__()
            d_2247_v131_ = nw363_
            d_2248_v132_: _dafny.Array
            def lambda114_(d_2249_p1_):
                def lambda115_(d_2250_i10_):
                    return d_2249_p1_

                return lambda115_

            init58_ = lambda114_(p1)
            nw364_ = _dafny.Array(None, 23)
            for i0_58_ in range(nw364_.length(0)):
                nw364_[i0_58_] = init58_(i0_58_)
            d_2248_v132_ = nw364_
            d_2251_v133_: D0
            d_2251_v133_ = D0_DC0(d_2248_v132_)
            d_2252_v134_: _dafny.Map
            d_2252_v134_ = _dafny.Map({415: d_2248_v132_})
            d_2253_v135_: _dafny.Map
            d_2253_v135_ = _dafny.Map({d_2184_v84_: d_2248_v132_})
            d_2254_v136_: _dafny.Seq
            d_2254_v136_ = _dafny.SeqWithoutIsStrInference([d_2248_v132_])
            d_2255_v137_: _dafny.Array
            nw365_ = _dafny.Array(None, 27)
            nw365_[int(0)] = d_2248_v132_
            nw365_[int(1)] = (d_2251_v133_).cf0
            nw365_[int(2)] = d_2248_v132_
            nw365_[int(3)] = d_2248_v132_
            nw365_[int(4)] = d_2248_v132_
            nw365_[int(5)] = d_2248_v132_
            nw365_[int(6)] = d_2248_v132_
            nw365_[int(7)] = ((d_2252_v134_)[p0] if (p0) in (d_2252_v134_) else d_2248_v132_)
            nw365_[int(8)] = ((d_2253_v135_)[p3] if (p3) in (d_2253_v135_) else d_2248_v132_)
            nw365_[int(9)] = d_2248_v132_
            nw365_[int(10)] = (d_2248_v132_ if p1 else d_2248_v132_)
            nw365_[int(11)] = d_2248_v132_
            nw365_[int(12)] = d_2248_v132_
            nw365_[int(13)] = d_2248_v132_
            nw365_[int(14)] = d_2248_v132_
            nw365_[int(15)] = d_2248_v132_
            nw365_[int(16)] = (d_2254_v136_)[default__.safeIndex(p0, len(d_2254_v136_))]
            nw365_[int(17)] = d_2248_v132_
            nw365_[int(18)] = d_2248_v132_
            nw365_[int(19)] = d_2248_v132_
            nw365_[int(20)] = d_2248_v132_
            nw365_[int(21)] = d_2248_v132_
            nw365_[int(22)] = ((d_2253_v135_)[True] if (True) in (d_2253_v135_) else (d_2254_v136_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([-388 for d_2256_i11_ in range(default__.abs(397))])), len(d_2254_v136_))])
            nw365_[int(23)] = d_2248_v132_
            nw365_[int(24)] = d_2248_v132_
            nw365_[int(25)] = d_2248_v132_
            nw365_[int(26)] = d_2248_v132_
            d_2255_v137_ = nw365_
            d_2257_v138_: D0
            d_2257_v138_ = D0_DC1(d_2248_v132_)
            index383_ = default__.safeIndex(340, (d_2255_v137_).length(0))
            (d_2255_v137_)[index383_] = (d_2257_v138_).cf1
            index384_ = default__.safeIndex(340, (d_2255_v137_).length(0))
            def lambda116_(d_2258_p1_):
                def lambda117_(d_2259_i12_):
                    return d_2258_p1_

                return lambda117_

            init59_ = lambda116_(p1)
            nw366_ = _dafny.Array(None, 22)
            for i0_59_ in range(nw366_.length(0)):
                nw366_[i0_59_] = init59_(i0_59_)
            (d_2255_v137_)[index384_] = nw366_
        elif True:
            d_2260___mcc_h9_ = source32_.cf93
            d_2261_cf93_ = d_2260___mcc_h9_
            (globalState).f4 = p0
            d_2262_v139_: _dafny.Array
            def lambda118_(d_2263_p0_):
                def lambda119_(d_2264_i13_):
                    return (d_2264_i13_) * (d_2263_p0_)

                return lambda119_

            init60_ = lambda118_(p0)
            nw367_ = _dafny.Array(None, 8)
            for i0_60_ in range(nw367_.length(0)):
                nw367_[i0_60_] = init60_(i0_60_)
            d_2262_v139_ = nw367_
            index385_ = default__.safeIndex(63, (d_2262_v139_).length(0))
            (d_2262_v139_)[index385_] = p0
            d_2265_v140_: _dafny.Map
            d_2265_v140_ = _dafny.Map({d_2262_v139_: (0) - (p0)})
            index386_ = default__.safeIndex(63, (d_2262_v139_).length(0))
            rhs462_ = default__.safeDivisionInt(p0, len(((d_2265_v140_).set(d_2262_v139_, p0)).set(d_2262_v139_, len(d_2115_v31_))))
            rhs463_ = (p0) - (p0)
            lhs311_ = d_2262_v139_
            lhs312_ = default__.safeIndex(63, (d_2262_v139_).length(0))
            lhs313_ = globalState
            lhs311_[lhs312_] = rhs462_
            lhs313_.f7 = rhs463_
            d_2118_v34_ = _dafny.CodePoint('b')
            d_2266_v141_: _dafny.Set
            d_2266_v141_ = _dafny.Set({d_2118_v34_})
            d_2267_v142_: D22
            d_2267_v142_ = D22_DC61(p0, 128, (self).fm2((d_2262_v139_)[default__.safeIndex(63, (d_2262_v139_).length(0))], globalState), (d_2186_v86_).set(default__.safeIndex(len(d_2266_v141_), len(d_2186_v86_)), d_2118_v34_))
            d_2268_v143_: _dafny.MultiSet
            d_2268_v143_ = _dafny.MultiSet([d_2267_v142_])
            (globalState).f13 = (0) - (((d_2262_v139_)[default__.safeIndex(63, (d_2262_v139_).length(0))]) - ((d_2268_v143_).cardinality))
        (globalState).f7 = p0
        d_2269_v144_: _dafny.MultiSet
        d_2269_v144_ = _dafny.MultiSet([(0) - (p0)])
        r0 = d_2269_v144_
        return r0


class C12(T0):
    def  __init__(self):
        self._f22: int = int(0)
        self._f23: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    def ctor__(self, f22, f23):
        (self)._f22 = f22
        (self)._f23 = f23

    def fm1(self, p0, p1, globalState):
        return ((self).f22) >= ((self).f22)

    def fm2(self, p0, globalState):
        return ((self).f22) * ((_dafny.MultiSet([True])).cardinality)

    def fm3(self, p0, globalState):
        return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(self).f22])), default__.safeModuloInt((self).f22, 561))

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife170_():
            coll84_ = _dafny.Set()
            compr_84_: int
            for compr_84_ in (_dafny.Map({(self).f22: False})).keys.Elements:
                d_2270_v0_: int = compr_84_
                if (d_2270_v0_) in (_dafny.Map({(self).f22: False})):
                    coll84_ = coll84_.union(_dafny.Set([(d_2270_v0_) * (-991)]))
            return _dafny.Set(coll84_)
        return ((_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))).cardinality})) | (_dafny.Set({(self).f22, (_dafny.MultiSet([True])).cardinality, len(_dafny.Map({True: (self).f22}))}))).intersection((_dafny.Set({(_dafny.MultiSet([(self).f22, (self).f22])).cardinality, (self).f22})).intersection(iife170_()
        ))

    def m1(self, p0, p1, globalState):
        d_2271_v0_: bool
        d_2271_v0_ = True
        (globalState).f4 = (((self).f22 if d_2271_v0_ else (self).f22)) * ((self).f22)
        d_2272_i0_: int
        d_2272_i0_ = 0
        with _dafny.label("15"):
            while d_2271_v0_:
                with _dafny.c_label("15"):
                    if (d_2272_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2272_i0_ = (d_2272_i0_) + (1)
                    d_2273_v1_: str
                    d_2273_v1_ = _dafny.CodePoint('r')
                    d_2273_v1_ = (p0)[default__.safeIndex((0) - ((self).f22), len(p0))]
                    d_2274_v2_: _dafny.Array
                    nw368_ = _dafny.Array(int(0), 29)
                    d_2274_v2_ = nw368_
                    d_2274_v2_ = d_2274_v2_
                    d_2275_v3_: _dafny.Set
                    d_2275_v3_ = _dafny.Set({(self).f22, (self).fm2((self).f22, globalState)})
                    d_2276_v4_: _dafny.Map
                    d_2276_v4_ = _dafny.Map({(d_2275_v3_).ispropersubset(d_2275_v3_): d_2271_v0_})
                    if ((d_2276_v4_)[d_2271_v0_] if (d_2271_v0_) in (d_2276_v4_) else False):
                        d_2277_v5_: _dafny.Map
                        d_2277_v5_ = _dafny.Map({d_2271_v0_: d_2276_v4_})
                        d_2278_v6_: _dafny.Seq
                        d_2278_v6_ = _dafny.SeqWithoutIsStrInference([d_2271_v0_, d_2271_v0_, d_2271_v0_])
                        d_2279_v7_: _dafny.Array
                        nw369_ = _dafny.Array(None, 6)
                        nw369_[int(0)] = (self).fm1(d_2277_v5_, (self).f22, globalState)
                        nw369_[int(1)] = (d_2278_v6_)[default__.safeIndex(732, len(d_2278_v6_))]
                        nw369_[int(2)] = (873) <= ((self).f22)
                        nw369_[int(3)] = True
                        nw369_[int(4)] = (p0) == (p0)
                        nw369_[int(5)] = d_2271_v0_
                        d_2279_v7_ = nw369_
                        index387_ = default__.safeIndex(274, (d_2279_v7_).length(0))
                        (d_2279_v7_)[index387_] = d_2271_v0_
                        index388_ = default__.safeIndex(274, (d_2279_v7_).length(0))
                        (d_2279_v7_)[index388_] = d_2271_v0_
                        (globalState).f7 = (self).f22
                        (globalState).f13 = (self).fm2((self).f22, globalState)
                        index389_ = default__.safeIndex(274, (d_2279_v7_).length(0))
                        (d_2279_v7_)[index389_] = (d_2273_v1_) not in ((self).f23)
                        nw370_ = _dafny.Array(False, 10)
                        d_2279_v7_ = nw370_
                    elif True:
                        d_2280_v8_: _dafny.Seq
                        d_2280_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlrwshec"))
                        d_2280_v8_ = ((self).f23 if ((self).f22) < (-642) else ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umwm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npllbkx")))).set(default__.safeIndex((0) - ((self).f22), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umwm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npllbkx"))))), d_2273_v1_))
                        d_2281_v9_: _dafny.Array
                        nw371_ = _dafny.Array(_dafny.Array(None, 0), 11)
                        d_2281_v9_ = nw371_
                        d_2282_v10_: _dafny.Array
                        def lambda120_(d_2283_v0_):
                            def lambda121_(d_2284_i1_):
                                return _dafny.Map({d_2283_v0_: -494})

                            return lambda121_

                        init61_ = lambda120_(d_2271_v0_)
                        nw372_ = _dafny.Array(None, 20)
                        for i0_61_ in range(nw372_.length(0)):
                            nw372_[i0_61_] = init61_(i0_61_)
                        d_2282_v10_ = nw372_
                        index390_ = default__.safeIndex(133, (d_2281_v9_).length(0))
                        (d_2281_v9_)[index390_] = d_2282_v10_
                        d_2285_v11_: _dafny.Array
                        def lambda122_(d_2286_v0_):
                            def lambda123_(d_2287_i2_):
                                return d_2286_v0_

                            return lambda123_

                        init62_ = lambda122_(d_2271_v0_)
                        nw373_ = _dafny.Array(None, 6)
                        for i0_62_ in range(nw373_.length(0)):
                            nw373_[i0_62_] = init62_(i0_62_)
                        d_2285_v11_ = nw373_
                        d_2288_v12_: _dafny.Seq
                        d_2288_v12_ = _dafny.SeqWithoutIsStrInference([d_2285_v11_])
                        index391_ = default__.safeIndex(133, (d_2281_v9_).length(0))
                        rhs464_ = d_2282_v10_
                        rhs465_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oexfuyyh"))
                        rhs466_ = len(d_2288_v12_)
                        lhs314_ = d_2281_v9_
                        lhs315_ = default__.safeIndex(133, (d_2281_v9_).length(0))
                        lhs316_ = globalState
                        lhs314_[lhs315_] = rhs464_
                        d_2280_v8_ = rhs465_
                        lhs316_.f7 = rhs466_
                        d_2280_v8_ = p0
                        d_2289_v13_: C6
                        nw374_ = C6()
                        nw374_.ctor__()
                        d_2289_v13_ = nw374_
                        (globalState).f7 = ((self).f22) * ((self).f22)
                    d_2271_v0_ = not(d_2271_v0_)
                    pass
            pass
        d_2290_v14_: _dafny.Map
        d_2290_v14_ = _dafny.Map({(self).f22: d_2271_v0_})
        d_2291_v15_: D14
        d_2291_v15_ = D14_DC37(D14_DC35(_dafny.SeqWithoutIsStrInference([d_2290_v14_])))
        d_2292_v16_: _dafny.Seq
        d_2292_v16_ = _dafny.SeqWithoutIsStrInference([d_2271_v0_, d_2271_v0_, d_2271_v0_])
        d_2291_v15_ = D14_DC37(D14_DC36(d_2292_v16_))
        hi15_ = (self).f22
        for d_2293_i3_ in range((self).f22, hi15_):
            if ((self).f22) > (d_2293_i3_):
                d_2294_v17_: _dafny.Array
                def lambda124_(d_2295_i3_):
                    def lambda125_(d_2296_i4_):
                        return default__.safeDivisionInt(d_2296_i4_, d_2295_i3_)

                    return lambda125_

                init63_ = lambda124_(d_2293_i3_)
                nw375_ = _dafny.Array(None, 7)
                for i0_63_ in range(nw375_.length(0)):
                    nw375_[i0_63_] = init63_(i0_63_)
                d_2294_v17_ = nw375_
                d_2294_v17_ = d_2294_v17_
                d_2271_v0_ = (d_2271_v0_) == (d_2271_v0_)
                d_2297_v18_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_2297_v18_ = nw376_
                d_2298_v19_: str
                d_2298_v19_ = _dafny.CodePoint('a')
                index392_ = default__.safeIndex(279, (d_2297_v18_).length(0))
                (d_2297_v18_)[index392_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spf"))).set(default__.safeIndex((self).f22, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spf")))), d_2298_v19_)
                d_2299_v20_: _dafny.Seq
                d_2299_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                index393_ = default__.safeIndex(279, (d_2297_v18_).length(0))
                rhs467_ = d_2294_v17_
                rhs468_ = (self).f23
                rhs469_ = ((((p0) + (p0)).set(default__.safeIndex(((self).f22) + (len(_dafny.SeqWithoutIsStrInference([d_2298_v19_ for d_2300_i5_ in range(default__.abs(847))]))), len((p0) + (p0))), d_2298_v19_)).set(default__.safeIndex(default__.safeDivisionInt(d_2293_i3_, 534), len(((p0) + (p0)).set(default__.safeIndex(((self).f22) + (len(_dafny.SeqWithoutIsStrInference([d_2298_v19_ for d_2301_i5_ in range(default__.abs(847))]))), len((p0) + (p0))), d_2298_v19_))), d_2298_v19_)).set(default__.safeIndex(((0) - ((0) - (len(d_2292_v16_))) if d_2271_v0_ else d_2293_i3_), len((((p0) + (p0)).set(default__.safeIndex(((self).f22) + (len(_dafny.SeqWithoutIsStrInference([d_2298_v19_ for d_2302_i5_ in range(default__.abs(847))]))), len((p0) + (p0))), d_2298_v19_)).set(default__.safeIndex(default__.safeDivisionInt(d_2293_i3_, 534), len(((p0) + (p0)).set(default__.safeIndex(((self).f22) + (len(_dafny.SeqWithoutIsStrInference([d_2298_v19_ for d_2303_i5_ in range(default__.abs(847))]))), len((p0) + (p0))), d_2298_v19_))), d_2298_v19_))), d_2298_v19_)
                lhs317_ = d_2297_v18_
                lhs318_ = default__.safeIndex(279, (d_2297_v18_).length(0))
                d_2294_v17_ = rhs467_
                lhs317_[lhs318_] = rhs468_
                d_2299_v20_ = rhs469_
                d_2304_v21_: D8
                d_2304_v21_ = D8_DC25((self).f22)
                def iife172_(_pat_let44_0):
                    def iife173_(d_2305_dt__update__tmp_h0_):
                        def iife174_(_pat_let45_0):
                            def iife175_(d_2306_dt__update_hcf42_h0_):
                                return D8_DC25(d_2306_dt__update_hcf42_h0_)
                            return iife175_(_pat_let45_0)
                        return iife174_(d_2293_i3_)
                    return iife173_(_pat_let44_0)
                def iife171_(_pat_let43_0):
                    def iife176_(d_2307_dt__update__tmp_h1_):
                        def iife177_(_pat_let46_0):
                            def iife178_(d_2308_dt__update_hcf42_h1_):
                                return D8_DC25(d_2308_dt__update_hcf42_h1_)
                            return iife178_(_pat_let46_0)
                        return iife177_((d_2293_i3_) - ((self).f22))
                    return iife176_(_pat_let43_0)
                d_2304_v21_ = iife171_(iife172_(d_2304_v21_))
                d_2298_v19_ = d_2298_v19_
            elif True:
                d_2309_v22_: _dafny.MultiSet
                d_2309_v22_ = _dafny.MultiSet([d_2271_v0_])
                (globalState).f13 = (d_2309_v22_).cardinality
                d_2310_v23_: _dafny.Array
                nw377_ = _dafny.Array(int(0), 7)
                d_2310_v23_ = nw377_
                index394_ = default__.safeIndex(387, (d_2310_v23_).length(0))
                (d_2310_v23_)[index394_] = d_2293_i3_
                d_2311_v24_: _dafny.Map
                d_2311_v24_ = _dafny.Map({d_2271_v0_: d_2271_v0_})
                d_2312_v25_: D15
                d_2312_v25_ = D15_DC39(d_2311_v24_, d_2271_v0_, (d_2309_v22_).cardinality)
                pat_let_tv50_ = d_2271_v0_
                index395_ = default__.safeIndex(387, (d_2310_v23_).length(0))
                def iife179_(_pat_let47_0):
                    def iife180_(d_2313_dt__update__tmp_h2_):
                        def iife181_(_pat_let48_0):
                            def iife182_(d_2314_dt__update_hcf57_h0_):
                                return D15_DC39((d_2313_dt__update__tmp_h2_).cf56, d_2314_dt__update_hcf57_h0_, (d_2313_dt__update__tmp_h2_).cf58)
                            return iife182_(_pat_let48_0)
                        return iife181_(not(pat_let_tv50_))
                    return iife180_(_pat_let47_0)
                (d_2310_v23_)[index395_] = ((iife179_(d_2312_v25_)).cf58 if d_2271_v0_ else (0) - ((0) - (((self).f22) * (d_2293_i3_))))
                d_2315_v26_: C2
                nw378_ = C2()
                nw378_.ctor__()
                d_2315_v26_ = nw378_
                d_2271_v0_ = not (d_2271_v0_) or ((d_2315_v26_).fm8(24, d_2271_v0_, d_2293_i3_, globalState))
                d_2316_v27_: _dafny.Array
                nw379_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_2316_v27_ = nw379_
                d_2317_v28_: _dafny.Array
                nw380_ = _dafny.Array(None, 18)
                nw380_[int(0)] = d_2271_v0_
                nw380_[int(1)] = d_2271_v0_
                nw380_[int(2)] = d_2271_v0_
                nw380_[int(3)] = d_2271_v0_
                nw380_[int(4)] = not(d_2271_v0_)
                nw380_[int(5)] = d_2271_v0_
                nw380_[int(6)] = d_2271_v0_
                nw380_[int(7)] = d_2271_v0_
                nw380_[int(8)] = d_2271_v0_
                nw380_[int(9)] = d_2271_v0_
                nw380_[int(10)] = d_2271_v0_
                nw380_[int(11)] = default__.fm26(globalState)
                nw380_[int(12)] = d_2271_v0_
                nw380_[int(13)] = d_2271_v0_
                nw380_[int(14)] = (d_2315_v26_).fm8((d_2310_v23_)[default__.safeIndex(387, (d_2310_v23_).length(0))], d_2271_v0_, default__.fm17(57, globalState), globalState)
                nw380_[int(15)] = d_2271_v0_
                nw380_[int(16)] = d_2271_v0_
                nw380_[int(17)] = ((d_2290_v14_)[d_2293_i3_] if (d_2293_i3_) in (d_2290_v14_) else d_2271_v0_)
                d_2317_v28_ = nw380_
                index396_ = default__.safeIndex(814, (d_2316_v27_).length(0))
                (d_2316_v27_)[index396_] = d_2317_v28_
                d_2318_v29_: str
                d_2318_v29_ = _dafny.CodePoint('a')
                d_2319_v30_: _dafny.MultiSet
                d_2319_v30_ = _dafny.MultiSet([d_2318_v29_, d_2318_v29_])
                d_2320_v31_: D7
                d_2320_v31_ = D7_DC22(d_2271_v0_, (d_2310_v23_)[default__.safeIndex(387, (d_2310_v23_).length(0))], d_2271_v0_, (d_2310_v23_)[default__.safeIndex(387, (d_2310_v23_).length(0))])
                index397_ = default__.safeIndex(814, (d_2316_v27_).length(0))
                nw381_ = _dafny.Array(None, 22)
                nw381_[int(0)] = (self).fm1(_dafny.Map({True: d_2311_v24_}), (0) - (d_2293_i3_), globalState)
                nw381_[int(1)] = (d_2271_v0_) or (d_2271_v0_)
                nw381_[int(2)] = d_2271_v0_
                nw381_[int(3)] = d_2271_v0_
                nw381_[int(4)] = (d_2271_v0_ if d_2271_v0_ else d_2271_v0_)
                nw381_[int(5)] = d_2271_v0_
                nw381_[int(6)] = (d_2319_v30_).issubset(d_2319_v30_)
                nw381_[int(7)] = True
                nw381_[int(8)] = not (True) or (d_2271_v0_)
                nw381_[int(9)] = (d_2271_v0_ if d_2271_v0_ else d_2271_v0_)
                nw381_[int(10)] = d_2271_v0_
                nw381_[int(11)] = d_2271_v0_
                nw381_[int(12)] = ((0) - (d_2293_i3_)) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_2310_v23_)[default__.safeIndex(387, (d_2310_v23_).length(0))]])))
                nw381_[int(13)] = d_2271_v0_
                nw381_[int(14)] = not(d_2271_v0_)
                nw381_[int(15)] = ((d_2290_v14_)[len(default__.fm35(globalState))] if (len(default__.fm35(globalState))) in (d_2290_v14_) else (d_2320_v31_).cf37)
                nw381_[int(16)] = False
                nw381_[int(17)] = d_2271_v0_
                nw381_[int(18)] = d_2271_v0_
                nw381_[int(19)] = True
                nw381_[int(20)] = False
                nw381_[int(21)] = not(d_2271_v0_)
                (d_2316_v27_)[index397_] = nw381_
            d_2321_v32_: _dafny.MultiSet
            d_2321_v32_ = _dafny.MultiSet([False, d_2271_v0_])
            d_2322_v33_: _dafny.Array
            nw382_ = _dafny.Array(None, 2)
            nw382_[int(0)] = d_2271_v0_
            nw382_[int(1)] = (_dafny.MultiSet([d_2271_v0_])).ispropersubset(d_2321_v32_)
            d_2322_v33_ = nw382_
            index398_ = default__.safeIndex(298, (d_2322_v33_).length(0))
            (d_2322_v33_)[index398_] = not (d_2271_v0_) or (d_2271_v0_)
            index399_ = default__.safeIndex(298, (d_2322_v33_).length(0))
            (d_2322_v33_)[index399_] = d_2271_v0_
            d_2323_v34_: _dafny.Array
            nw383_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_2323_v34_ = nw383_
            d_2324_v35_: _dafny.Array
            nw384_ = _dafny.Array(None, 19)
            nw384_[int(0)] = 684
            nw384_[int(1)] = (self).f22
            nw384_[int(2)] = (self).f22
            nw384_[int(3)] = (self).f22
            nw384_[int(4)] = 871
            nw384_[int(5)] = len(default__.fm27(globalState))
            nw384_[int(6)] = d_2293_i3_
            nw384_[int(7)] = (self).f22
            nw384_[int(8)] = d_2293_i3_
            nw384_[int(9)] = (self).f22
            nw384_[int(10)] = (self).f22
            nw384_[int(11)] = len(d_2292_v16_)
            nw384_[int(12)] = (self).f22
            nw384_[int(13)] = d_2293_i3_
            nw384_[int(14)] = d_2293_i3_
            nw384_[int(15)] = (d_2321_v32_).cardinality
            nw384_[int(16)] = (self).f22
            nw384_[int(17)] = (self).f22
            nw384_[int(18)] = (self).f22
            d_2324_v35_ = nw384_
            d_2325_v36_: D2
            d_2325_v36_ = D2_DC5(d_2324_v35_)
            d_2326_v38_: _dafny.Map
            def iife183_():
                coll85_ = _dafny.Set()
                compr_85_: int
                for compr_85_ in _dafny.IntegerRange(303, -84):
                    d_2327_v37_: int = compr_85_
                    if ((303) <= (d_2327_v37_)) and ((d_2327_v37_) < (-84)):
                        coll85_ = coll85_.union(_dafny.Set([(d_2327_v37_) * (d_2293_i3_)]))
                return _dafny.Set(coll85_)
            d_2326_v38_ = _dafny.Map({(self).f22: default__.fm30(default__.fm17(d_2293_i3_, globalState), (d_2322_v33_)[default__.safeIndex(298, (d_2322_v33_).length(0))], iife183_()
            , d_2271_v0_, globalState)})
            d_2328_v39_: _dafny.Seq
            d_2328_v39_ = _dafny.SeqWithoutIsStrInference([d_2326_v38_, _dafny.Map({(self).f22: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpjnbt"))}), d_2326_v38_])
            d_2329_v40_: _dafny.Map
            d_2329_v40_ = _dafny.Map({len((d_2328_v39_)[default__.safeIndex((0) - ((self).f22), len(d_2328_v39_))]): d_2324_v35_})
            d_2330_v41_: _dafny.Array
            nw385_ = _dafny.Array(None, 12)
            nw385_[int(0)] = (d_2325_v36_).cf6
            nw385_[int(1)] = d_2324_v35_
            nw385_[int(2)] = d_2324_v35_
            nw385_[int(3)] = d_2324_v35_
            nw385_[int(4)] = d_2324_v35_
            nw385_[int(5)] = d_2324_v35_
            nw385_[int(6)] = d_2324_v35_
            nw385_[int(7)] = d_2324_v35_
            nw385_[int(8)] = d_2324_v35_
            nw385_[int(9)] = d_2324_v35_
            nw385_[int(10)] = d_2324_v35_
            nw385_[int(11)] = ((d_2329_v40_)[d_2293_i3_] if (d_2293_i3_) in (d_2329_v40_) else d_2324_v35_)
            d_2330_v41_ = nw385_
            index400_ = default__.safeIndex(647, (d_2323_v34_).length(0))
            (d_2323_v34_)[index400_] = d_2330_v41_
            index401_ = default__.safeIndex(647, (d_2323_v34_).length(0))
            (d_2323_v34_)[index401_] = d_2330_v41_
            if (p0) <= (p0):
                (globalState).f15 = (self).f22
                d_2331_v42_: _dafny.Seq
                d_2331_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks"))
                d_2331_v42_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2332_i6_ in range(default__.abs(-693))]) if (d_2322_v33_)[default__.safeIndex(298, (d_2322_v33_).length(0))] else (d_2331_v42_) + (p0))
                d_2333_v43_: _dafny.Array
                nw386_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_2333_v43_ = nw386_
                index402_ = default__.safeIndex(736, (d_2333_v43_).length(0))
                (d_2333_v43_)[index402_] = d_2322_v33_
                d_2334_v44_: C2
                nw387_ = C2()
                nw387_.ctor__()
                d_2334_v44_ = nw387_
                d_2335_v45_: _dafny.Map
                d_2335_v45_ = _dafny.Map({d_2334_v44_: d_2322_v33_})
                index403_ = default__.safeIndex(736, (d_2333_v43_).length(0))
                (d_2333_v43_)[index403_] = ((d_2335_v45_)[d_2334_v44_] if (d_2334_v44_) in (d_2335_v45_) else d_2322_v33_)
                d_2336_v46_: _dafny.Set
                d_2336_v46_ = _dafny.Set({(d_2322_v33_)[default__.safeIndex(298, (d_2322_v33_).length(0))]})
                d_2290_v14_ = (d_2290_v14_).set(default__.safeDivisionInt((0) - ((self).f22), d_2293_i3_), (d_2336_v46_).isdisjoint(d_2336_v46_))
                d_2336_v46_ = (d_2336_v46_) | (d_2336_v46_)
            elif True:
                d_2337_v47_: T0
                nw388_ = C11()
                nw388_.ctor__()
                d_2337_v47_ = nw388_
                d_2338_v48_: _dafny.Map
                d_2338_v48_ = _dafny.Map({d_2337_v47_: (d_2322_v33_)[default__.safeIndex(298, (d_2322_v33_).length(0))]})
                d_2339_v49_: _dafny.Seq
                d_2339_v49_ = _dafny.SeqWithoutIsStrInference([d_2338_v48_, _dafny.Map({d_2337_v47_: (d_2322_v33_)[default__.safeIndex(298, (d_2322_v33_).length(0))]})])
                d_2271_v0_ = not((d_2339_v49_) != (d_2339_v49_))
                index404_ = default__.safeIndex(298, (d_2322_v33_).length(0))
                (d_2322_v33_)[index404_] = not(d_2271_v0_)
                d_2340_v50_: C2
                nw389_ = C2()
                nw389_.ctor__()
                d_2340_v50_ = nw389_
                d_2341_v51_: _dafny.Seq
                d_2341_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhounbr"))
                d_2342_v52_: D0
                d_2342_v52_ = D0_DC1(d_2322_v33_)
                index405_ = default__.safeIndex(298, (d_2322_v33_).length(0))
                rhs470_ = (0) - (d_2293_i3_)
                rhs471_ = d_2271_v0_
                rhs472_ = ((d_2326_v38_)[len((self).f23)] if (len((self).f23)) in (d_2326_v38_) else (self).f23)
                rhs473_ = d_2342_v52_
                lhs319_ = globalState
                lhs320_ = d_2322_v33_
                lhs321_ = default__.safeIndex(298, (d_2322_v33_).length(0))
                lhs319_.f15 = rhs470_
                lhs320_[lhs321_] = rhs471_
                d_2341_v51_ = rhs472_
                d_2342_v52_ = rhs473_
                d_2343_v53_: _dafny.Seq
                d_2343_v53_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt((self).f22, d_2293_i3_)])
                d_2344_v54_: _dafny.Seq
                d_2344_v54_ = _dafny.SeqWithoutIsStrInference([d_2341_v51_])
                rhs474_ = len(d_2290_v14_)
                rhs475_ = d_2343_v53_
                rhs476_ = p0
                rhs477_ = d_2344_v54_
                lhs322_ = globalState
                lhs322_.f4 = rhs474_
                d_2343_v53_ = rhs475_
                d_2341_v51_ = rhs476_
                d_2344_v54_ = rhs477_
        d_2345_v55_: C0
        nw390_ = C0()
        nw390_.ctor__()
        d_2345_v55_ = nw390_
        d_2346_v56_: _dafny.Array
        def lambda126_(d_2347_i7_):
            return (d_2347_i7_) - ((self).f22)

        init64_ = lambda126_
        nw391_ = _dafny.Array(None, 13)
        for i0_64_ in range(nw391_.length(0)):
            nw391_[i0_64_] = init64_(i0_64_)
        d_2346_v56_ = nw391_
        index406_ = default__.safeIndex(366, (d_2346_v56_).length(0))
        (d_2346_v56_)[index406_] = (self).f22
        d_2348_v57_: T1
        nw392_ = C2()
        nw392_.ctor__()
        d_2348_v57_ = nw392_
        d_2349_v58_: str
        d_2349_v58_ = _dafny.CodePoint('n')
        d_2350_v59_: _dafny.Set
        d_2350_v59_ = _dafny.Set({d_2271_v0_})
        d_2351_v60_: D5
        d_2351_v60_ = D5_DC14(d_2349_v58_, d_2271_v0_, len(d_2350_v59_), d_2271_v0_)
        d_2352_v61_: _dafny.Map
        d_2352_v61_ = _dafny.Map({d_2351_v60_: (self).f22})
        d_2353_v62_: _dafny.Map
        d_2353_v62_ = _dafny.Map({_dafny.CodePoint('h'): len(d_2352_v61_)})
        index407_ = default__.safeIndex(366, (d_2346_v56_).length(0))
        (d_2346_v56_)[index407_] = (len(_dafny.SeqWithoutIsStrInference([d_2348_v57_, d_2348_v57_]))) - (((d_2353_v62_)[d_2349_v58_] if (d_2349_v58_) in (d_2353_v62_) else (self).f22))

    def m2(self, p0, p1, p2, globalState):
        d_2354_v0_: str
        d_2354_v0_ = _dafny.CodePoint('k')
        d_2355_v1_: _dafny.Set
        d_2355_v1_ = _dafny.Set({(self).f23, ((self).f23).set(default__.safeIndex((self).f22, len((self).f23)), d_2354_v0_)})
        if ((self).f23) in ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jieavydw")), default__.fm19(p1, p1, globalState)})) | (d_2355_v1_)):
            d_2356_v2_: _dafny.Seq
            d_2356_v2_ = _dafny.SeqWithoutIsStrInference([p1, not(not(not(p1)))])
            d_2357_v3_: _dafny.Seq
            d_2357_v3_ = _dafny.SeqWithoutIsStrInference([(d_2356_v2_)[default__.safeIndex(len(d_2356_v2_), len(d_2356_v2_))]])
            d_2358_v4_: D18
            d_2358_v4_ = D18_DC49(p1, p1, d_2354_v0_, (d_2357_v3_)[default__.safeIndex(p2, len(d_2357_v3_))])
            source33_ = d_2358_v4_
            if source33_.is_DC49:
                d_2359___mcc_h0_ = source33_.cf75
                d_2360___mcc_h1_ = source33_.cf76
                d_2361___mcc_h2_ = source33_.cf77
                d_2362___mcc_h3_ = source33_.cf78
                d_2363_cf78_ = d_2362___mcc_h3_
                d_2364_cf77_ = d_2361___mcc_h2_
                d_2365_cf76_ = d_2360___mcc_h1_
                d_2366_cf75_ = d_2359___mcc_h0_
                d_2367_v5_: _dafny.Array
                nw393_ = _dafny.Array(D5.default()(), 21)
                d_2367_v5_ = nw393_
                d_2368_v6_: _dafny.Map
                d_2368_v6_ = _dafny.Map({p0: d_2363_cf78_})
                d_2369_v7_: _dafny.Set
                d_2369_v7_ = _dafny.Set({len(d_2368_v6_)})
                d_2370_v9_: _dafny.MultiSet
                def iife184_():
                    coll86_ = _dafny.Set()
                    compr_86_: int
                    for compr_86_ in _dafny.IntegerRange(713, 165):
                        d_2371_v8_: int = compr_86_
                        if ((713) <= (d_2371_v8_)) and ((d_2371_v8_) < (165)):
                            coll86_ = coll86_.union(_dafny.Set([(d_2371_v8_) * (p0)]))
                    return _dafny.Set(coll86_)
                d_2370_v9_ = _dafny.MultiSet([d_2369_v7_, iife184_()
                ])
                d_2372_v10_: D5
                d_2372_v10_ = D5_DC13(d_2370_v9_)
                index408_ = default__.safeIndex(443, (d_2367_v5_).length(0))
                (d_2367_v5_)[index408_] = d_2372_v10_
                pat_let_tv51_ = d_2369_v7_
                pat_let_tv52_ = d_2369_v7_
                pat_let_tv53_ = d_2370_v9_
                index409_ = default__.safeIndex(443, (d_2367_v5_).length(0))
                def iife185_(_pat_let49_0):
                    def iife186_(d_2373_dt__update__tmp_h0_):
                        def iife187_(_pat_let50_0):
                            def iife188_(d_2374_dt__update_hcf17_h0_):
                                return D5_DC13(d_2374_dt__update_hcf17_h0_)
                            return iife188_(_pat_let50_0)
                        return iife187_((_dafny.MultiSet([pat_let_tv51_, pat_let_tv52_])) - (pat_let_tv53_))
                    return iife186_(_pat_let49_0)
                (d_2367_v5_)[index409_] = iife185_(d_2372_v10_)
                d_2375_v11_: _dafny.MultiSet
                d_2375_v11_ = _dafny.MultiSet([False])
                d_2376_v12_: _dafny.Array
                nw394_ = _dafny.Array(int(0), 27)
                d_2376_v12_ = nw394_
                d_2377_v13_: D2
                d_2377_v13_ = D2_DC8(D2_DC7(((d_2368_v6_)[(self).f22] if ((self).f22) in (d_2368_v6_) else d_2366_cf75_), d_2375_v11_, d_2376_v12_))
                d_2378_v14_: C3
                nw395_ = C3()
                nw395_.ctor__(d_2377_v13_, False)
                d_2378_v14_ = nw395_
                rhs478_ = (p2) - ((self).f22)
                rhs479_ = (len((d_2356_v2_) + (_dafny.SeqWithoutIsStrInference([d_2378_v14_.f25])))) < (p0)
                rhs480_ = d_2378_v14_
                lhs323_ = globalState
                lhs323_.f4 = rhs478_
                d_2363_cf78_ = rhs479_
                d_2378_v14_ = rhs480_
                d_2363_cf78_ = not(d_2366_cf75_)
                d_2379_v15_: _dafny.MultiSet
                d_2379_v15_ = _dafny.MultiSet([(_dafny.MultiSet(d_2357_v3_)).cardinality, (self).f22])
                d_2380_v16_: _dafny.Seq
                d_2380_v16_ = _dafny.SeqWithoutIsStrInference([(d_2379_v15_).cardinality, len((self).f23)])
                (globalState).f15 = len(d_2380_v16_)
            elif source33_.is_DC48:
                d_2381___mcc_h4_ = source33_.cf74
                d_2382_cf74_ = d_2381___mcc_h4_
                d_2383_v17_: _dafny.Array
                nw396_ = _dafny.Array(_dafny.Seq({}), 18)
                d_2383_v17_ = nw396_
                d_2384_v18_: _dafny.Array
                def lambda127_(d_2385_v0_):
                    def lambda128_(d_2386_i0_):
                        return d_2385_v0_

                    return lambda128_

                init65_ = lambda127_(d_2354_v0_)
                nw397_ = _dafny.Array(None, 22)
                for i0_65_ in range(nw397_.length(0)):
                    nw397_[i0_65_] = init65_(i0_65_)
                d_2384_v18_ = nw397_
                d_2387_v19_: _dafny.Seq
                d_2387_v19_ = _dafny.SeqWithoutIsStrInference([d_2384_v18_])
                index410_ = default__.safeIndex(764, (d_2383_v17_).length(0))
                (d_2383_v17_)[index410_] = d_2387_v19_
                d_2388_v20_: _dafny.Seq
                d_2388_v20_ = _dafny.SeqWithoutIsStrInference([d_2387_v19_])
                d_2389_v21_: _dafny.Seq
                d_2389_v21_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2384_v18_, d_2384_v18_, d_2384_v18_])) + (d_2387_v19_), ((d_2388_v20_)[default__.safeIndex(p0, len(d_2388_v20_))]).set(default__.safeIndex((0) - ((self).f22), len((d_2388_v20_)[default__.safeIndex(p0, len(d_2388_v20_))])), d_2384_v18_), d_2387_v19_])
                index411_ = default__.safeIndex(764, (d_2383_v17_).length(0))
                (d_2383_v17_)[index411_] = (d_2389_v21_)[default__.safeIndex((len((self).f23)) - ((0) - ((self).f22)), len(d_2389_v21_))]
                d_2390_v22_: _dafny.MultiSet
                d_2390_v22_ = _dafny.MultiSet([p1, p1])
                d_2391_v23_: D1
                d_2391_v23_ = D1_DC2(((d_2390_v22_)[p1] if (p1) in (d_2390_v22_) else 143))
                d_2392_v24_: D20
                d_2392_v24_ = D20_DC55(p2, (self).f22, p1, p1)
                d_2393_v25_: D8
                d_2393_v25_ = D8_DC24(p0, False)
                d_2394_v26_: _dafny.Map
                d_2394_v26_ = _dafny.Map({len(_dafny.Map({p1: p1})): _dafny.Map({p1: d_2393_v25_})})
                d_2395_v27_: _dafny.Map
                d_2395_v27_ = _dafny.Map({p1: d_2393_v25_})
                d_2396_v28_: _dafny.Set
                d_2396_v28_ = _dafny.Set({not(not(p1)), not(p1)})
                d_2397_v29_: _dafny.Array
                nw398_ = _dafny.Array(None, 11)
                nw398_[int(0)] = (self).f22
                nw398_[int(1)] = p0
                nw398_[int(2)] = 678
                nw398_[int(3)] = p2
                nw398_[int(4)] = (self).f22
                nw398_[int(5)] = (d_2392_v24_).cf85
                nw398_[int(6)] = len(((d_2394_v26_)[p2] if (p2) in (d_2394_v26_) else d_2395_v27_))
                nw398_[int(7)] = (d_2390_v22_).cardinality
                nw398_[int(8)] = len(d_2396_v28_)
                nw398_[int(9)] = p2
                nw398_[int(10)] = -907
                d_2397_v29_ = nw398_
                d_2398_v30_: _dafny.Map
                d_2398_v30_ = _dafny.Map({d_2397_v29_: _dafny.MultiSet([p1])})
                d_2399_v31_: _dafny.Array
                nw399_ = _dafny.Array(None, 21)
                nw399_[int(0)] = (_dafny.MultiSet((d_2356_v2_).set(default__.safeIndex(len(default__.fm19(p1, p1, globalState)), len(d_2356_v2_)), p1))) | (d_2390_v22_)
                nw399_[int(1)] = _dafny.MultiSet([not(p1), p1])
                nw399_[int(2)] = _dafny.MultiSet([p1])
                nw399_[int(3)] = (d_2390_v22_).set(False, default__.abs(p0))
                nw399_[int(4)] = (d_2390_v22_) | (default__.fm34((self).f22, p0, p1, globalState))
                nw399_[int(5)] = _dafny.MultiSet(d_2357_v3_)
                nw399_[int(6)] = d_2390_v22_
                nw399_[int(7)] = (d_2390_v22_).set(False, default__.abs(p2))
                nw399_[int(8)] = d_2390_v22_
                nw399_[int(9)] = d_2390_v22_
                nw399_[int(10)] = _dafny.MultiSet([not(p1)])
                nw399_[int(11)] = (d_2390_v22_) | (d_2390_v22_)
                nw399_[int(12)] = d_2390_v22_
                nw399_[int(13)] = d_2390_v22_
                nw399_[int(14)] = (d_2390_v22_) | (d_2390_v22_)
                nw399_[int(15)] = d_2390_v22_
                nw399_[int(16)] = (d_2390_v22_ if p1 else d_2390_v22_)
                nw399_[int(17)] = d_2390_v22_
                nw399_[int(18)] = _dafny.MultiSet([p1])
                nw399_[int(19)] = (d_2390_v22_ if p1 else d_2390_v22_)
                nw399_[int(20)] = ((d_2398_v30_)[d_2397_v29_] if (d_2397_v29_) in (d_2398_v30_) else d_2390_v22_)
                d_2399_v31_ = nw399_
                d_2400_v32_: _dafny.Map
                d_2400_v32_ = _dafny.Map({p1: (d_2357_v3_)[default__.safeIndex(p0, len(d_2357_v3_))]})
                index412_ = default__.safeIndex(7, (d_2384_v18_).length(0))
                (d_2384_v18_)[index412_] = ((self).f23)[default__.safeIndex((self).f22, len((self).f23))]
                d_2401_v33_: _dafny.Seq
                d_2401_v33_ = _dafny.SeqWithoutIsStrInference([d_2399_v31_, d_2399_v31_])
                index413_ = default__.safeIndex(7, (d_2384_v18_).length(0))
                rhs481_ = d_2391_v23_
                rhs482_ = (d_2401_v33_)[default__.safeIndex(509, len(d_2401_v33_))]
                rhs483_ = d_2400_v32_
                rhs484_ = d_2354_v0_
                rhs485_ = (0) - (p2)
                lhs324_ = d_2384_v18_
                lhs325_ = default__.safeIndex(7, (d_2384_v18_).length(0))
                lhs326_ = globalState
                d_2391_v23_ = rhs481_
                d_2399_v31_ = rhs482_
                d_2400_v32_ = rhs483_
                lhs324_[lhs325_] = rhs484_
                lhs326_.f13 = rhs485_
                (globalState).f15 = (0) - (p0)
                d_2402_v34_: _dafny.Seq
                d_2402_v34_ = _dafny.SeqWithoutIsStrInference([p0, p2])
                (globalState).f15 = (d_2402_v34_)[default__.safeIndex(219, len(d_2402_v34_))]
            elif True:
                d_2403___mcc_h5_ = source33_.cf79
                d_2404_cf79_ = d_2403___mcc_h5_
                d_2405_v35_: _dafny.Map
                d_2405_v35_ = _dafny.Map({p1: p1})
                d_2406_v36_: _dafny.Map
                d_2406_v36_ = _dafny.Map({p1: d_2405_v35_})
                d_2407_v37_: _dafny.Array
                nw400_ = _dafny.Array(None, 20)
                nw400_[int(0)] = d_2354_v0_
                nw400_[int(1)] = d_2354_v0_
                nw400_[int(2)] = d_2354_v0_
                nw400_[int(3)] = d_2354_v0_
                nw400_[int(4)] = d_2354_v0_
                nw400_[int(5)] = d_2354_v0_
                nw400_[int(6)] = (d_2354_v0_ if p1 else d_2354_v0_)
                nw400_[int(7)] = d_2354_v0_
                nw400_[int(8)] = default__.fm38(p1, globalState)
                nw400_[int(9)] = d_2354_v0_
                nw400_[int(10)] = (d_2354_v0_ if (self).fm1(d_2406_v36_, len(d_2356_v2_), globalState) else d_2354_v0_)
                nw400_[int(11)] = d_2354_v0_
                nw400_[int(12)] = d_2354_v0_
                nw400_[int(13)] = (d_2354_v0_ if p1 else d_2354_v0_)
                nw400_[int(14)] = d_2354_v0_
                nw400_[int(15)] = d_2354_v0_
                nw400_[int(16)] = d_2354_v0_
                nw400_[int(17)] = d_2354_v0_
                nw400_[int(18)] = d_2354_v0_
                nw400_[int(19)] = d_2354_v0_
                d_2407_v37_ = nw400_
                d_2407_v37_ = d_2407_v37_
                d_2408_v38_: _dafny.Map
                d_2408_v38_ = _dafny.Map({(0) - ((self).f22): p2})
                d_2409_v39_: _dafny.Map
                d_2409_v39_ = _dafny.Map({(self).f23: D18_DC48(d_2408_v38_)})
                d_2410_v41_: _dafny.MultiSet
                d_2410_v41_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2354_v0_ for d_2411_i1_ in range(default__.abs(-696))])])
                def iife189_():
                    coll87_ = _dafny.Map()
                    compr_87_: _dafny.Seq
                    for compr_87_ in (d_2410_v41_).Elements:
                        d_2412_v40_: _dafny.Seq = compr_87_
                        if (d_2412_v40_) in (d_2410_v41_):
                            coll87_[d_2412_v40_] = D18_DC48(d_2408_v38_)
                    return _dafny.Map(coll87_)
                (globalState).f15 = ((self).f22) * (len((d_2409_v39_) | (iife189_()
                )))
                d_2413_v42_: bool
                d_2413_v42_ = True
                d_2413_v42_ = d_2413_v42_
                d_2414_v43_: _dafny.Set
                d_2414_v43_ = _dafny.Set({p1})
                d_2415_v44_: _dafny.Seq
                d_2415_v44_ = _dafny.SeqWithoutIsStrInference([(self).f22, len(d_2414_v43_), p2])
                d_2416_v45_: D6
                d_2416_v45_ = D6_DC17((self).fm1(d_2406_v36_, len(d_2415_v44_), globalState), not(d_2413_v42_), p1)
                d_2413_v42_ = (d_2416_v45_).cf24
            d_2417_v46_: _dafny.Set
            d_2417_v46_ = _dafny.Set({not(p1), p1})
            d_2418_v47_: _dafny.Seq
            d_2418_v47_ = _dafny.SeqWithoutIsStrInference([p2, len(d_2417_v46_), (self).f22])
            d_2419_v48_: _dafny.Seq
            d_2419_v48_ = _dafny.SeqWithoutIsStrInference([d_2418_v47_, d_2418_v47_])
            d_2419_v48_ = (d_2419_v48_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(0) - (p0)])), len(d_2419_v48_)), d_2418_v47_)
            (globalState).f7 = default__.fm17(p0, globalState)
            d_2357_v3_ = d_2357_v3_
            d_2420_v49_: _dafny.Array
            def lambda129_(d_2421_p1_):
                def lambda130_(d_2422_i2_):
                    return d_2421_p1_

                return lambda130_

            init66_ = lambda129_(p1)
            nw401_ = _dafny.Array(None, 29)
            for i0_66_ in range(nw401_.length(0)):
                nw401_[i0_66_] = init66_(i0_66_)
            d_2420_v49_ = nw401_
            index414_ = default__.safeIndex(12, (d_2420_v49_).length(0))
            (d_2420_v49_)[index414_] = p1
            index415_ = default__.safeIndex(12, (d_2420_v49_).length(0))
            (d_2420_v49_)[index415_] = p1
        elif True:
            d_2423_v50_: _dafny.Array
            def lambda131_(d_2424_p1_):
                def lambda132_(d_2425_i3_):
                    return d_2424_p1_

                return lambda132_

            init67_ = lambda131_(p1)
            nw402_ = _dafny.Array(None, 5)
            for i0_67_ in range(nw402_.length(0)):
                nw402_[i0_67_] = init67_(i0_67_)
            d_2423_v50_ = nw402_
            d_2423_v50_ = d_2423_v50_
            d_2426_v51_: bool
            d_2426_v51_ = True
            d_2427_v52_: T0
            nw403_ = C11()
            nw403_.ctor__()
            d_2427_v52_ = nw403_
            d_2428_v53_: _dafny.Set
            d_2428_v53_ = _dafny.Set({d_2427_v52_, d_2427_v52_})
            d_2429_v54_: D7
            d_2429_v54_ = D7_DC22(p1, p0, p1, len(d_2428_v53_))
            d_2430_v55_: _dafny.Map
            d_2430_v55_ = _dafny.Map({p1: d_2426_v51_})
            d_2431_v56_: D7
            d_2431_v56_ = D7_DC22(not((d_2429_v54_).cf35), (0) - (p0), ((d_2430_v55_)[d_2426_v51_] if (d_2426_v51_) in (d_2430_v55_) else True), p0)
            d_2426_v51_ = (d_2431_v56_).cf37
            d_2432_v57_: _dafny.MultiSet
            d_2432_v57_ = _dafny.MultiSet([d_2426_v51_])
            d_2426_v51_ = not((p1) not in (d_2432_v57_))
            d_2433_v58_: C8
            nw404_ = C8()
            nw404_.ctor__()
            d_2433_v58_ = nw404_
            d_2423_v50_ = d_2423_v50_
        d_2434_v59_: C4
        nw405_ = C4()
        nw405_.ctor__()
        d_2434_v59_ = nw405_
        d_2435_v61_: _dafny.MultiSet
        d_2435_v61_ = _dafny.MultiSet([(self).f22, p0, (self).f22])
        d_2436_v62_: _dafny.Seq
        def iife190_():
            coll88_ = _dafny.Map()
            compr_88_: int
            for compr_88_ in (d_2435_v61_).Elements:
                d_2437_v60_: int = compr_88_
                if (d_2437_v60_) in (d_2435_v61_):
                    coll88_[default__.safeDivisionInt(d_2437_v60_, 13)] = p1
            return _dafny.Map(coll88_)
        d_2436_v62_ = _dafny.SeqWithoutIsStrInference([iife190_()
        ])
        d_2438_v63_: D14
        d_2438_v63_ = D14_DC35(d_2436_v62_)
        d_2439_i4_: int
        d_2439_i4_ = 0
        with _dafny.label("16"):
            pat_let_tv54_ = p1
            pat_let_tv55_ = p1
            pat_let_tv56_ = p1
            def lambda140_(source35_):
                if source35_.is_DC36:
                    d_2510___mcc_h18_ = source35_.cf53
                    d_2511_cf53_ = d_2510___mcc_h18_
                    return pat_let_tv54_
                elif source35_.is_DC35:
                    d_2512___mcc_h19_ = source35_.cf52
                    d_2513_cf52_ = d_2512___mcc_h19_
                    return pat_let_tv55_
                elif True:
                    d_2514___mcc_h20_ = source35_.cf54
                    d_2515_cf54_ = d_2514___mcc_h20_
                    return not(pat_let_tv56_)

            while lambda140_(d_2438_v63_):
                with _dafny.c_label("16"):
                    if (d_2439_i4_) >= (100):
                        raise _dafny.Break("16")
                    d_2439_i4_ = (d_2439_i4_) + (1)
                    d_2440_v64_: C8
                    nw406_ = C8()
                    nw406_.ctor__()
                    d_2440_v64_ = nw406_
                    d_2441_v65_: D22
                    d_2441_v65_ = D22_DC60((self).f22, (self).f22)
                    source34_ = d_2441_v65_
                    if source34_.is_DC60:
                        d_2442___mcc_h6_ = source34_.cf95
                        d_2443___mcc_h7_ = source34_.cf96
                        d_2444_cf96_ = d_2443___mcc_h7_
                        d_2445_cf95_ = d_2442___mcc_h6_
                        d_2446_v66_: C9
                        nw407_ = C9()
                        nw407_.ctor__()
                        d_2446_v66_ = nw407_
                        d_2447_v67_: bool
                        d_2447_v67_ = True
                        d_2447_v67_ = True
                        d_2448_v68_: _dafny.Array
                        def lambda133_(d_2449_p1_):
                            def lambda134_(d_2450_i5_):
                                return d_2449_p1_

                            return lambda134_

                        init68_ = lambda133_(p1)
                        nw408_ = _dafny.Array(None, 1)
                        for i0_68_ in range(nw408_.length(0)):
                            nw408_[i0_68_] = init68_(i0_68_)
                        d_2448_v68_ = nw408_
                        d_2448_v68_ = d_2448_v68_
                        (globalState).f7 = d_2445_cf95_
                    elif source34_.is_DC61:
                        d_2451___mcc_h8_ = source34_.cf97
                        d_2452___mcc_h9_ = source34_.cf98
                        d_2453___mcc_h10_ = source34_.cf99
                        d_2454___mcc_h11_ = source34_.cf100
                        d_2455_cf100_ = d_2454___mcc_h11_
                        d_2456_cf99_ = d_2453___mcc_h10_
                        d_2457_cf98_ = d_2452___mcc_h9_
                        d_2458_cf97_ = d_2451___mcc_h8_
                        d_2459_v69_: bool
                        d_2459_v69_ = True
                        d_2459_v69_ = p1
                        d_2460_v70_: _dafny.Array
                        def lambda135_(d_2461_v0_):
                            def lambda136_(d_2462_i6_):
                                return (_dafny.SeqWithoutIsStrInference([d_2461_v0_ for d_2463_i7_ in range(default__.abs(-459))])) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "claiu")))

                            return lambda136_

                        init69_ = lambda135_(d_2354_v0_)
                        nw409_ = _dafny.Array(None, 19)
                        for i0_69_ in range(nw409_.length(0)):
                            nw409_[i0_69_] = init69_(i0_69_)
                        d_2460_v70_ = nw409_
                        d_2460_v70_ = d_2460_v70_
                        d_2464_v71_: _dafny.MultiSet
                        d_2464_v71_ = _dafny.MultiSet([p1])
                        d_2465_v72_: _dafny.Map
                        d_2465_v72_ = _dafny.Map({d_2456_cf99_: p2})
                        d_2466_v73_: _dafny.Seq
                        d_2466_v73_ = _dafny.SeqWithoutIsStrInference([d_2459_v69_])
                        d_2467_v74_: _dafny.Set
                        d_2467_v74_ = _dafny.Set({(self).f22})
                        d_2468_v75_: D22
                        d_2468_v75_ = D22_DC62(d_2455_cf100_, p2, len(d_2467_v74_), d_2458_cf97_, p1)
                        d_2469_v76_: _dafny.Array
                        nw410_ = _dafny.Array(None, 23)
                        nw410_[int(0)] = ((d_2465_v72_)[p2] if (p2) in (d_2465_v72_) else d_2456_cf99_)
                        nw410_[int(1)] = p2
                        nw410_[int(2)] = len(d_2466_v73_)
                        nw410_[int(3)] = 343
                        nw410_[int(4)] = (self).f22
                        nw410_[int(5)] = d_2457_cf98_
                        nw410_[int(6)] = p0
                        nw410_[int(7)] = len(_dafny.Map({p0: d_2468_v75_}))
                        nw410_[int(8)] = (0) - (p2)
                        nw410_[int(9)] = (0) - (d_2458_cf97_)
                        nw410_[int(10)] = d_2457_cf98_
                        nw410_[int(11)] = p2
                        nw410_[int(12)] = -744
                        nw410_[int(13)] = d_2458_cf97_
                        nw410_[int(14)] = p0
                        nw410_[int(15)] = (self).f22
                        nw410_[int(16)] = (self).f22
                        nw410_[int(17)] = 661
                        nw410_[int(18)] = p0
                        nw410_[int(19)] = (self).fm3(p1, globalState)
                        nw410_[int(20)] = (self).f22
                        nw410_[int(21)] = d_2457_cf98_
                        nw410_[int(22)] = len(d_2455_cf100_)
                        d_2469_v76_ = nw410_
                        d_2470_v77_: D2
                        d_2470_v77_ = D2_DC5(d_2469_v76_)
                        (d_2440_v64_).m6(len((self).f23), d_2464_v71_, (p1 if d_2459_v69_ else p1), d_2470_v77_, globalState)
                        d_2354_v0_ = d_2354_v0_
                    elif source34_.is_DC62:
                        d_2471___mcc_h12_ = source34_.cf101
                        d_2472___mcc_h13_ = source34_.cf102
                        d_2473___mcc_h14_ = source34_.cf103
                        d_2474___mcc_h15_ = source34_.cf104
                        d_2475___mcc_h16_ = source34_.cf105
                        d_2476_cf105_ = d_2475___mcc_h16_
                        d_2477_cf104_ = d_2474___mcc_h15_
                        d_2478_cf103_ = d_2473___mcc_h14_
                        d_2479_cf102_ = d_2472___mcc_h13_
                        d_2480_cf101_ = d_2471___mcc_h12_
                        d_2481_v78_: _dafny.Map
                        d_2481_v78_ = _dafny.Map({(p1 if p1 else p1): (self).f22})
                        d_2481_v78_ = d_2481_v78_
                        d_2482_v79_: _dafny.Array
                        def lambda137_(d_2483_p0_):
                            def lambda138_(d_2484_i8_):
                                return default__.safeModuloInt(d_2484_i8_, d_2483_p0_)

                            return lambda138_

                        init70_ = lambda137_(p0)
                        nw411_ = _dafny.Array(None, 21)
                        for i0_70_ in range(nw411_.length(0)):
                            nw411_[i0_70_] = init70_(i0_70_)
                        d_2482_v79_ = nw411_
                        index416_ = default__.safeIndex(350, (d_2482_v79_).length(0))
                        (d_2482_v79_)[index416_] = d_2479_cf102_
                        index417_ = default__.safeIndex(350, (d_2482_v79_).length(0))
                        (d_2482_v79_)[index417_] = 205
                        d_2485_v80_: D2
                        d_2485_v80_ = D2_DC5(d_2482_v79_)
                        d_2486_v81_: D2
                        d_2486_v81_ = D2_DC8(d_2485_v80_)
                        d_2486_v81_ = d_2486_v81_
                        d_2487_v82_: D22
                        d_2487_v82_ = D22_DC61(d_2479_cf102_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwpuonys"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goqrmbm"))), d_2480_cf101_)
                        d_2488_v83_: _dafny.Map
                        d_2488_v83_ = _dafny.Map({p1: p1})
                        d_2489_v84_: _dafny.Map
                        d_2489_v84_ = _dafny.Map({d_2488_v83_: d_2480_cf101_})
                        d_2490_v85_: _dafny.Map
                        d_2490_v85_ = _dafny.Map({(d_2482_v79_)[default__.safeIndex(350, (d_2482_v79_).length(0))]: d_2480_cf101_})
                        d_2491_v86_: _dafny.Array
                        nw412_ = _dafny.Array(None, 18)
                        nw412_[int(0)] = (d_2480_cf101_) + ((self).f23)
                        nw412_[int(1)] = ((self).f23) + (d_2480_cf101_)
                        nw412_[int(2)] = (self).f23
                        nw412_[int(3)] = (d_2487_v82_).cf100
                        nw412_[int(4)] = (self).f23
                        nw412_[int(5)] = ((d_2487_v82_).cf100) + (d_2480_cf101_)
                        nw412_[int(6)] = d_2480_cf101_
                        nw412_[int(7)] = (d_2480_cf101_).set(default__.safeIndex((d_2482_v79_)[default__.safeIndex(350, (d_2482_v79_).length(0))], len(d_2480_cf101_)), _dafny.CodePoint('o'))
                        nw412_[int(8)] = (d_2480_cf101_).set(default__.safeIndex(-44, len(d_2480_cf101_)), d_2354_v0_)
                        nw412_[int(9)] = d_2480_cf101_
                        nw412_[int(10)] = (((d_2489_v84_)[d_2488_v83_] if (d_2488_v83_) in (d_2489_v84_) else (self).f23)) + ((self).f23)
                        nw412_[int(11)] = (self).f23
                        nw412_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                        nw412_[int(13)] = (self).f23
                        nw412_[int(14)] = d_2480_cf101_
                        nw412_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfd"))
                        nw412_[int(16)] = ((d_2490_v85_)[(d_2440_v64_).fm2((0) - (d_2478_cf103_), globalState)] if ((d_2440_v64_).fm2((0) - (d_2478_cf103_), globalState)) in (d_2490_v85_) else (self).f23)
                        nw412_[int(17)] = (self).f23
                        d_2491_v86_ = nw412_
                        index418_ = default__.safeIndex(470, (d_2491_v86_).length(0))
                        (d_2491_v86_)[index418_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnimhcmi"))) + ((self).f23)
                        d_2492_v87_: _dafny.Seq
                        d_2492_v87_ = _dafny.SeqWithoutIsStrInference([d_2476_cf105_])
                        d_2493_v88_: _dafny.Array
                        def lambda139_(d_2494_i9_):
                            return False

                        init71_ = lambda139_
                        nw413_ = _dafny.Array(None, 19)
                        for i0_71_ in range(nw413_.length(0)):
                            nw413_[i0_71_] = init71_(i0_71_)
                        d_2493_v88_ = nw413_
                        d_2495_v89_: D0
                        d_2495_v89_ = D0_DC1(d_2493_v88_)
                        d_2496_v90_: _dafny.Map
                        d_2496_v90_ = _dafny.Map({d_2495_v89_: d_2492_v87_})
                        d_2497_v91_: T1
                        nw414_ = C5()
                        nw414_.ctor__()
                        d_2497_v91_ = nw414_
                        d_2498_v92_: _dafny.Map
                        d_2498_v92_ = _dafny.Map({d_2497_v91_: d_2492_v87_})
                        d_2499_v93_: _dafny.Array
                        nw415_ = _dafny.Array(None, 26)
                        nw415_[int(0)] = d_2492_v87_
                        nw415_[int(1)] = d_2492_v87_
                        nw415_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2476_cf105_, d_2476_cf105_])
                        nw415_[int(3)] = _dafny.SeqWithoutIsStrInference([d_2476_cf105_])
                        nw415_[int(4)] = d_2492_v87_
                        nw415_[int(5)] = _dafny.SeqWithoutIsStrInference([(d_2434_v59_).fm8(d_2477_cf104_, d_2476_cf105_, 302, globalState)])
                        nw415_[int(6)] = (D14_DC36(_dafny.SeqWithoutIsStrInference([p1, d_2476_cf105_]))).cf53
                        nw415_[int(7)] = d_2492_v87_
                        nw415_[int(8)] = _dafny.SeqWithoutIsStrInference([p1])
                        nw415_[int(9)] = d_2492_v87_
                        nw415_[int(10)] = d_2492_v87_
                        nw415_[int(11)] = d_2492_v87_
                        nw415_[int(12)] = d_2492_v87_
                        nw415_[int(13)] = ((d_2496_v90_)[d_2495_v89_] if (d_2495_v89_) in (d_2496_v90_) else d_2492_v87_)
                        nw415_[int(14)] = d_2492_v87_
                        nw415_[int(15)] = (d_2492_v87_).set(default__.safeIndex(d_2477_cf104_, len(d_2492_v87_)), d_2476_cf105_)
                        nw415_[int(16)] = d_2492_v87_
                        nw415_[int(17)] = d_2492_v87_
                        nw415_[int(18)] = d_2492_v87_
                        nw415_[int(19)] = d_2492_v87_
                        nw415_[int(20)] = d_2492_v87_
                        nw415_[int(21)] = default__.fm27(globalState)
                        nw415_[int(22)] = d_2492_v87_
                        nw415_[int(23)] = d_2492_v87_
                        nw415_[int(24)] = d_2492_v87_
                        nw415_[int(25)] = ((d_2498_v92_)[d_2497_v91_] if (d_2497_v91_) in (d_2498_v92_) else d_2492_v87_)
                        d_2499_v93_ = nw415_
                        d_2500_v94_: _dafny.Seq
                        d_2500_v94_ = _dafny.SeqWithoutIsStrInference([d_2499_v93_, d_2499_v93_, d_2499_v93_])
                        d_2501_v95_: _dafny.Array
                        nw416_ = _dafny.Array(None, 17)
                        nw416_[int(0)] = d_2499_v93_
                        nw416_[int(1)] = d_2499_v93_
                        nw416_[int(2)] = d_2499_v93_
                        nw416_[int(3)] = d_2499_v93_
                        nw416_[int(4)] = d_2499_v93_
                        nw416_[int(5)] = d_2499_v93_
                        nw416_[int(6)] = d_2499_v93_
                        nw416_[int(7)] = d_2499_v93_
                        nw416_[int(8)] = d_2499_v93_
                        nw416_[int(9)] = d_2499_v93_
                        nw416_[int(10)] = d_2499_v93_
                        nw416_[int(11)] = d_2499_v93_
                        nw416_[int(12)] = (d_2500_v94_)[default__.safeIndex(d_2478_cf103_, len(d_2500_v94_))]
                        nw416_[int(13)] = d_2499_v93_
                        nw416_[int(14)] = d_2499_v93_
                        nw416_[int(15)] = d_2499_v93_
                        nw416_[int(16)] = d_2499_v93_
                        d_2501_v95_ = nw416_
                        index419_ = default__.safeIndex(721, (d_2501_v95_).length(0))
                        (d_2501_v95_)[index419_] = d_2499_v93_
                        index420_ = default__.safeIndex(470, (d_2491_v86_).length(0))
                        index421_ = default__.safeIndex(721, (d_2501_v95_).length(0))
                        rhs486_ = (self).f23
                        rhs487_ = (d_2499_v93_ if p1 else d_2499_v93_)
                        rhs488_ = (0) - ((d_2477_cf104_) - ((d_2482_v79_)[default__.safeIndex(350, (d_2482_v79_).length(0))]))
                        rhs489_ = d_2354_v0_
                        lhs327_ = d_2491_v86_
                        lhs328_ = default__.safeIndex(470, (d_2491_v86_).length(0))
                        lhs329_ = d_2501_v95_
                        lhs330_ = default__.safeIndex(721, (d_2501_v95_).length(0))
                        lhs331_ = globalState
                        lhs327_[lhs328_] = rhs486_
                        lhs329_[lhs330_] = rhs487_
                        lhs331_.f15 = rhs488_
                        d_2354_v0_ = rhs489_
                    elif True:
                        d_2502___mcc_h17_ = source34_.cf94
                        d_2503_cf94_ = d_2502___mcc_h17_
                        d_2504_v96_: D8
                        d_2504_v96_ = D8_DC25(((0) - (p2)) * ((self).f22))
                        d_2504_v96_ = d_2504_v96_
                        (globalState).f13 = (p2 if p1 else p0)
                        d_2505_v97_: bool
                        d_2505_v97_ = False
                        d_2505_v97_ = d_2505_v97_
                        d_2506_v98_: _dafny.Seq
                        d_2506_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdybm"))
                        rhs490_ = (self).f23
                        rhs491_ = ((d_2435_v61_).intersection(default__.fm31((self).f22, (self).f22, globalState))) - ((_dafny.MultiSet([(self).f22, p0, -8])) | (_dafny.MultiSet([len(d_2506_v98_)])))
                        d_2506_v98_ = rhs490_
                        d_2435_v61_ = rhs491_
                    d_2507_v99_: bool
                    d_2507_v99_ = True
                    d_2507_v99_ = d_2507_v99_
                    d_2508_v100_: _dafny.Seq
                    d_2508_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkiv"))
                    d_2508_v100_ = _dafny.SeqWithoutIsStrInference([d_2354_v0_ for d_2509_i10_ in range(default__.abs(-119))])
                    pass
            pass
        d_2516_v101_: _dafny.Array
        nw417_ = _dafny.Array(int(0), 17)
        d_2516_v101_ = nw417_
        d_2517_v102_: _dafny.Seq
        d_2517_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwc"))
        rhs492_ = d_2516_v101_
        rhs493_ = (self).f23
        d_2516_v101_ = rhs492_
        d_2517_v102_ = rhs493_
        d_2518_v103_: D6
        d_2518_v103_ = D6_DC17(p1, True, p1)
        d_2519_i11_: int
        d_2519_i11_ = 0
        with _dafny.label("17"):
            while ((d_2518_v103_ if p1 else d_2518_v103_)).cf26:
                with _dafny.c_label("17"):
                    if (d_2519_i11_) >= (100):
                        raise _dafny.Break("17")
                    d_2519_i11_ = (d_2519_i11_) + (1)
                    d_2520_v104_: _dafny.MultiSet
                    d_2520_v104_ = _dafny.MultiSet([p1, not(p1), (p1 if p1 else p1), False, p1])
                    d_2521_v105_: _dafny.Set
                    d_2521_v105_ = _dafny.Set({not(not(p1))})
                    rhs494_ = (_dafny.MultiSet([p1])).intersection(_dafny.MultiSet([p1, p1]))
                    rhs495_ = d_2516_v101_
                    rhs496_ = (d_2521_v105_) | (_dafny.Set({p1}))
                    d_2520_v104_ = rhs494_
                    d_2516_v101_ = rhs495_
                    d_2521_v105_ = rhs496_
                    if False:
                        (globalState).f13 = p0
                        d_2522_v106_: _dafny.Array
                        def lambda141_(d_2523_p1_):
                            def lambda142_(d_2524_i12_):
                                return d_2523_p1_

                            return lambda142_

                        init72_ = lambda141_(p1)
                        nw418_ = _dafny.Array(None, 11)
                        for i0_72_ in range(nw418_.length(0)):
                            nw418_[i0_72_] = init72_(i0_72_)
                        d_2522_v106_ = nw418_
                        d_2525_v107_: _dafny.Map
                        d_2525_v107_ = _dafny.Map({d_2521_v105_: p1})
                        index422_ = default__.safeIndex(970, (d_2522_v106_).length(0))
                        (d_2522_v106_)[index422_] = ((d_2434_v59_).fm8(len(d_2525_v107_), p1, (0) - (p2), globalState)) or (p1)
                        d_2526_v108_: bool
                        d_2526_v108_ = True
                        d_2527_v109_: _dafny.Seq
                        d_2527_v109_ = _dafny.SeqWithoutIsStrInference([p2, p0])
                        d_2528_v110_: D1
                        d_2528_v110_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohan")), False)
                        index423_ = default__.safeIndex(970, (d_2522_v106_).length(0))
                        rhs497_ = (default__.safeDivisionInt((self).f22, len(d_2517_v102_))) * (len(d_2527_v109_))
                        rhs498_ = True
                        rhs499_ = default__.fm30(p2, (p2) >= ((self).f22), (self).fm4(d_2517_v102_, p2, d_2526_v108_, d_2528_v110_, globalState), d_2526_v108_, globalState)
                        rhs500_ = (p0) >= (p0)
                        rhs501_ = p1
                        lhs332_ = globalState
                        lhs333_ = d_2522_v106_
                        lhs334_ = default__.safeIndex(970, (d_2522_v106_).length(0))
                        lhs332_.f13 = rhs497_
                        lhs333_[lhs334_] = rhs498_
                        d_2517_v102_ = rhs499_
                        d_2526_v108_ = rhs500_
                        d_2526_v108_ = rhs501_
                        d_2529_v111_: D14
                        d_2529_v111_ = D14_DC35(d_2436_v62_)
                        d_2530_v112_: D14
                        d_2530_v112_ = D14_DC37(d_2529_v111_)
                        d_2531_v113_: D14
                        d_2531_v113_ = D14_DC37(d_2529_v111_)
                        d_2532_v114_: D14
                        d_2532_v114_ = D14_DC37(d_2530_v112_)
                        d_2533_v115_: _dafny.Map
                        d_2533_v115_ = _dafny.Map({d_2532_v114_: p2})
                        d_2534_v116_: T0
                        nw419_ = C9()
                        nw419_.ctor__()
                        d_2534_v116_ = nw419_
                        d_2535_v117_: D17
                        d_2535_v117_ = D17_DC47(d_2534_v116_, d_2526_v108_, (d_2435_v61_).cardinality, (d_2534_v116_).fm2(p2, globalState), (self).f22)
                        d_2536_v118_: _dafny.Map
                        d_2536_v118_ = _dafny.Map({d_2526_v108_: 708})
                        (globalState).f15 = (default__.safeDivisionInt(((d_2533_v115_)[d_2532_v114_] if (d_2532_v114_) in (d_2533_v115_) else p0), (self).f22)) - (((d_2535_v117_).cf72) - (((d_2536_v118_)[d_2526_v108_] if (d_2526_v108_) in (d_2536_v118_) else p2)))
                        index424_ = default__.safeIndex(368, (d_2516_v101_).length(0))
                        (d_2516_v101_)[index424_] = p2
                        index425_ = default__.safeIndex(368, (d_2516_v101_).length(0))
                        (d_2516_v101_)[index425_] = p0
                        d_2537_v119_: C9
                        nw420_ = C9()
                        nw420_.ctor__()
                        d_2537_v119_ = nw420_
                        d_2538_v120_: _dafny.MultiSet
                        d_2538_v120_ = _dafny.MultiSet([d_2537_v119_, d_2537_v119_, d_2537_v119_, d_2537_v119_, d_2537_v119_])
                        d_2538_v120_ = _dafny.MultiSet([d_2537_v119_, d_2537_v119_])
                    elif True:
                        d_2539_v121_: _dafny.Seq
                        d_2539_v121_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_2539_v121_ = (d_2539_v121_ if p1 else d_2539_v121_)
                        d_2540_v122_: bool
                        d_2540_v122_ = False
                        rhs502_ = not (d_2540_v122_) or (d_2540_v122_)
                        rhs503_ = d_2520_v104_
                        d_2540_v122_ = rhs502_
                        d_2520_v104_ = rhs503_
                        d_2541_v123_: _dafny.Seq
                        d_2541_v123_ = _dafny.SeqWithoutIsStrInference([(self).f23])
                        d_2517_v102_ = (((((d_2541_v123_)[default__.safeIndex((self).f22, len(d_2541_v123_))]).set(default__.safeIndex(p2, len((d_2541_v123_)[default__.safeIndex((self).f22, len(d_2541_v123_))])), d_2354_v0_)).set(default__.safeIndex((self).f22, len(((d_2541_v123_)[default__.safeIndex((self).f22, len(d_2541_v123_))]).set(default__.safeIndex(p2, len((d_2541_v123_)[default__.safeIndex((self).f22, len(d_2541_v123_))])), d_2354_v0_))), d_2354_v0_)) + (d_2517_v102_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxdtrqap")))
                        d_2516_v101_ = d_2516_v101_
                        index426_ = default__.safeIndex(335, (d_2516_v101_).length(0))
                        (d_2516_v101_)[index426_] = (0) - (p0)
                        index427_ = default__.safeIndex(335, (d_2516_v101_).length(0))
                        (d_2516_v101_)[index427_] = p0
                    (globalState).f13 = len(default__.fm59(d_2354_v0_, p1, globalState))
                    d_2542_v124_: C1
                    nw421_ = C1()
                    nw421_.ctor__()
                    d_2542_v124_ = nw421_
                    pass
            pass
        d_2543_v125_: C1
        nw422_ = C1()
        nw422_.ctor__()
        d_2543_v125_ = nw422_

    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23

class C13:
    def  __init__(self):
        self.f21: bool = False
        self._f20: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    def ctor__(self, f20, f21):
        (self)._f20 = f20
        (self).f21 = f21

    def fm0(self, p0, p1, p2, p3, globalState):
        def iife191_():
            coll89_ = _dafny.Map()
            compr_89_: int
            for compr_89_ in _dafny.IntegerRange(180, 472):
                d_2544_v0_: int = compr_89_
                if ((180) <= (d_2544_v0_)) and ((d_2544_v0_) < (472)):
                    coll89_[default__.safeModuloInt(d_2544_v0_, (self).f20)] = (self).f20
            return _dafny.Map(coll89_)
        return (_dafny.Map({(self).f20: (self).f20})) | ((iife191_()
        ) | (_dafny.Map({(self).f20: len(_dafny.Map({self.f21: self.f21}))})))

    def m0(self, p0, p1, p2, globalState):
        r0: str = _dafny.CodePoint('D')
        if p2:
            d_2545_v0_: _dafny.Array
            nw423_ = _dafny.Array(False, 16)
            d_2545_v0_ = nw423_
            d_2546_v1_: D0
            d_2546_v1_ = D0_DC0(d_2545_v0_)
            d_2547_v2_: _dafny.Seq
            d_2547_v2_ = _dafny.SeqWithoutIsStrInference([d_2545_v0_, d_2545_v0_, d_2545_v0_, d_2545_v0_])
            pat_let_tv57_ = d_2545_v0_
            d_2548_v3_: _dafny.Seq
            def iife192_(_pat_let51_0):
                def iife193_(d_2549_dt__update__tmp_h0_):
                    def iife194_(_pat_let52_0):
                        def iife195_(d_2550_dt__update_hcf0_h0_):
                            return D0_DC0(d_2550_dt__update_hcf0_h0_)
                        return iife195_(_pat_let52_0)
                    return iife194_(pat_let_tv57_)
                return iife193_(_pat_let51_0)
            d_2548_v3_ = _dafny.SeqWithoutIsStrInference([d_2545_v0_, (iife192_(d_2546_v1_)).cf0, (d_2547_v2_)[default__.safeIndex((D1_DC2((_dafny.MultiSet([(self).f20])).cardinality)).cf2, len(d_2547_v2_))], d_2545_v0_])
            d_2548_v3_ = (d_2547_v2_) + (d_2548_v3_)
            d_2551_v4_: str
            d_2551_v4_ = _dafny.CodePoint('c')
            d_2552_v5_: _dafny.Map
            d_2552_v5_ = _dafny.Map({d_2551_v4_: d_2551_v4_})
            d_2552_v5_ = (d_2552_v5_).set(d_2551_v4_, d_2551_v4_)
            d_2553_v6_: C7
            nw424_ = C7()
            nw424_.ctor__()
            d_2553_v6_ = nw424_
            (self).f21 = self.f21
            d_2554_v7_: _dafny.Seq
            d_2554_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dj"))
            d_2555_v8_: C12
            nw425_ = C12()
            nw425_.ctor__(-287, d_2554_v7_)
            d_2555_v8_ = nw425_
        elif True:
            (globalState).f13 = (0) - ((self).f20)
            d_2556_v9_: C2
            nw426_ = C2()
            nw426_.ctor__()
            d_2556_v9_ = nw426_
            d_2557_v10_: _dafny.MultiSet
            d_2557_v10_ = _dafny.MultiSet([False, False, p1, p2])
            d_2558_v11_: _dafny.Set
            d_2558_v11_ = _dafny.Set({p2, p2})
            d_2559_v12_: _dafny.Seq
            d_2559_v12_ = _dafny.SeqWithoutIsStrInference([self.f21, p1])
            d_2557_v10_ = default__.fm34((len(d_2558_v11_)) * ((self).f20), (_dafny.MultiSet(d_2559_v12_)).cardinality, (self.f21) or ((d_2559_v12_)[default__.safeIndex((self).f20, len(d_2559_v12_))]), globalState)
            d_2560_v13_: _dafny.Array
            def lambda143_(d_2561_i0_):
                return (d_2561_i0_) + (439)

            init73_ = lambda143_
            nw427_ = _dafny.Array(None, 13)
            for i0_73_ in range(nw427_.length(0)):
                nw427_[i0_73_] = init73_(i0_73_)
            d_2560_v13_ = nw427_
            d_2562_v14_: _dafny.Set
            d_2562_v14_ = _dafny.Set({d_2560_v13_, d_2560_v13_, d_2560_v13_, d_2560_v13_})
            (self).f21 = ((d_2562_v14_).intersection(d_2562_v14_)).ispropersubset((_dafny.Set({d_2560_v13_, d_2560_v13_})).intersection(d_2562_v14_))
            d_2563_v15_: _dafny.Map
            d_2563_v15_ = _dafny.Map({(978) == (len(d_2559_v12_)): (d_2558_v11_).intersection(d_2558_v11_)})
            d_2563_v15_ = d_2563_v15_
        d_2564_v16_: _dafny.Seq
        d_2564_v16_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2565_v17_: _dafny.Map
        d_2565_v17_ = _dafny.Map({self.f21: _dafny.SeqWithoutIsStrInference([False, p1, p1, p0, p0])})
        d_2566_v18_: _dafny.Array
        nw428_ = _dafny.Array(None, 6)
        nw428_[int(0)] = d_2564_v16_
        nw428_[int(1)] = d_2564_v16_
        nw428_[int(2)] = ((d_2565_v17_)[self.f21] if (self.f21) in (d_2565_v17_) else d_2564_v16_)
        nw428_[int(3)] = default__.fm27(globalState)
        nw428_[int(4)] = (D8_DC23(d_2564_v16_)).cf39
        nw428_[int(5)] = d_2564_v16_
        d_2566_v18_ = nw428_
        d_2567_v19_: _dafny.MultiSet
        d_2567_v19_ = _dafny.MultiSet([(self).f20, (self).f20, (self).f20])
        d_2568_v20_: _dafny.MultiSet
        d_2568_v20_ = _dafny.MultiSet([((d_2567_v19_).set((self).f20, default__.abs((self).f20))).cardinality])
        d_2569_v21_: _dafny.Seq
        d_2569_v21_ = _dafny.SeqWithoutIsStrInference([d_2567_v19_, d_2568_v20_])
        d_2570_v22_: _dafny.Seq
        d_2570_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onoohtgc"))
        d_2571_v23_: _dafny.Map
        d_2571_v23_ = _dafny.Map({(self).f20: False})
        d_2572_v24_: _dafny.Array
        nw429_ = _dafny.Array(None, 22)
        nw429_[int(0)] = p2
        nw429_[int(1)] = p1
        nw429_[int(2)] = default__.fm26(globalState)
        nw429_[int(3)] = (d_2569_v21_) == (d_2569_v21_)
        nw429_[int(4)] = p2
        nw429_[int(5)] = p1
        nw429_[int(6)] = self.f21
        nw429_[int(7)] = not(not (p2) or (p0))
        nw429_[int(8)] = p2
        nw429_[int(9)] = p0
        nw429_[int(10)] = p0
        nw429_[int(11)] = (d_2564_v16_)[default__.safeIndex(len(d_2570_v22_), len(d_2564_v16_))]
        nw429_[int(12)] = default__.fm26(globalState)
        nw429_[int(13)] = (p2 if not(p2) else ((d_2571_v23_)[(self).f20] if ((self).f20) in (d_2571_v23_) else p1))
        nw429_[int(14)] = p2
        nw429_[int(15)] = p0
        nw429_[int(16)] = self.f21
        nw429_[int(17)] = True
        nw429_[int(18)] = (d_2570_v22_) < (d_2570_v22_)
        nw429_[int(19)] = p2
        nw429_[int(20)] = p2
        nw429_[int(21)] = (p2) and (self.f21)
        d_2572_v24_ = nw429_
        index428_ = default__.safeIndex(895, (d_2572_v24_).length(0))
        (d_2572_v24_)[index428_] = (p0 if self.f21 else self.f21)
        d_2573_v25_: str
        d_2573_v25_ = _dafny.CodePoint('l')
        pat_let_tv58_ = d_2570_v22_
        pat_let_tv59_ = d_2570_v22_
        index429_ = default__.safeIndex(895, (d_2572_v24_).length(0))
        def lambda144_(source36_):
            if source36_.is_DC20:
                d_2574___mcc_h0_ = source36_.cf29
                d_2575_cf29_ = d_2574___mcc_h0_
                return 881
            elif source36_.is_DC21:
                d_2576___mcc_h1_ = source36_.cf30
                d_2577___mcc_h2_ = source36_.cf31
                d_2578___mcc_h3_ = source36_.cf32
                d_2579___mcc_h4_ = source36_.cf33
                d_2580___mcc_h5_ = source36_.cf34
                d_2581_cf34_ = d_2580___mcc_h5_
                d_2582_cf33_ = d_2579___mcc_h4_
                d_2583_cf32_ = d_2578___mcc_h3_
                d_2584_cf31_ = d_2577___mcc_h2_
                d_2585_cf30_ = d_2576___mcc_h1_
                return -379
            elif source36_.is_DC22:
                d_2586___mcc_h6_ = source36_.cf35
                d_2587___mcc_h7_ = source36_.cf36
                d_2588___mcc_h8_ = source36_.cf37
                d_2589___mcc_h9_ = source36_.cf38
                d_2590_cf38_ = d_2589___mcc_h9_
                d_2591_cf37_ = d_2588___mcc_h8_
                d_2592_cf36_ = d_2587___mcc_h7_
                d_2593_cf35_ = d_2586___mcc_h6_
                return len(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2594_i1_ in range(default__.abs(170))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug")), pat_let_tv58_, (D22_DC62(pat_let_tv59_, (self).f20, d_2590_cf38_, d_2590_cf38_, False)).cf101}))
            elif True:
                d_2595___mcc_h10_ = source36_.cf28
                d_2596_cf28_ = d_2595___mcc_h10_
                return (self).f20

        rhs504_ = lambda144_(D7_DC21((self).f20, (self).f20, (d_2564_v16_)[default__.safeIndex((self).f20, len(d_2564_v16_))], self.f21, False))
        rhs505_ = d_2566_v18_
        rhs506_ = not((d_2573_v25_) in (d_2570_v22_))
        rhs507_ = (self).f20
        lhs335_ = globalState
        lhs336_ = d_2572_v24_
        lhs337_ = default__.safeIndex(895, (d_2572_v24_).length(0))
        lhs338_ = globalState
        lhs335_.f15 = rhs504_
        d_2566_v18_ = rhs505_
        lhs336_[lhs337_] = rhs506_
        lhs338_.f7 = rhs507_
        d_2597_v26_: _dafny.MultiSet
        d_2597_v26_ = _dafny.MultiSet([self.f21])
        d_2598_v27_: _dafny.Seq
        d_2598_v27_ = _dafny.SeqWithoutIsStrInference([(d_2597_v26_).cardinality])
        d_2599_v28_: _dafny.Map
        d_2599_v28_ = _dafny.Map({p2: self.f21})
        d_2600_v29_: _dafny.Map
        d_2600_v29_ = _dafny.Map({len(d_2598_v27_): len(d_2599_v28_)})
        d_2601_v30_: D18
        d_2601_v30_ = D18_DC48(d_2600_v29_)
        d_2602_v31_: _dafny.Seq
        d_2602_v31_ = _dafny.SeqWithoutIsStrInference([d_2601_v30_, d_2601_v30_])
        d_2603_v32_: _dafny.Array
        nw430_ = _dafny.Array(int(0), 19)
        d_2603_v32_ = nw430_
        d_2604_v33_: D2
        d_2604_v33_ = D2_DC7(p2, _dafny.MultiSet([True, not(p0)]), d_2603_v32_)
        d_2605_v34_: D2
        d_2605_v34_ = D2_DC8(d_2604_v33_)
        d_2606_v35_: C3
        nw431_ = C3()
        nw431_.ctor__(d_2605_v34_, False)
        d_2606_v35_ = nw431_
        d_2607_v36_: _dafny.Set
        d_2607_v36_ = _dafny.Set({d_2606_v35_})
        d_2608_v37_: _dafny.Set
        d_2608_v37_ = _dafny.Set({d_2606_v35_.f25})
        d_2609_v38_: _dafny.Map
        d_2609_v38_ = _dafny.Map({(self).f20: d_2564_v16_})
        d_2610_v39_: _dafny.Map
        d_2610_v39_ = _dafny.Map({d_2573_v25_: d_2609_v38_})
        d_2611_v40_: _dafny.Array
        nw432_ = _dafny.Array(None, 23)
        nw432_[int(0)] = len(_dafny.SeqWithoutIsStrInference([len(d_2570_v22_)]))
        nw432_[int(1)] = ((d_2597_v26_)[p1] if (p1) in (d_2597_v26_) else len(_dafny.Map({False: d_2602_v31_})))
        nw432_[int(2)] = (self).f20
        nw432_[int(3)] = len(d_2607_v36_)
        nw432_[int(4)] = (self).f20
        nw432_[int(5)] = (self).f20
        nw432_[int(6)] = (self).f20
        nw432_[int(7)] = (d_2568_v20_).cardinality
        nw432_[int(8)] = (self).f20
        nw432_[int(9)] = len(d_2598_v27_)
        nw432_[int(10)] = (self).f20
        nw432_[int(11)] = len(d_2608_v37_)
        nw432_[int(12)] = (self).f20
        nw432_[int(13)] = (self).f20
        nw432_[int(14)] = (self).f20
        nw432_[int(15)] = (self).f20
        nw432_[int(16)] = (self).f20
        nw432_[int(17)] = len(d_2599_v28_)
        nw432_[int(18)] = (self).f20
        nw432_[int(19)] = len(((d_2610_v39_)[default__.fm38(d_2606_v35_.f25, globalState)] if (default__.fm38(d_2606_v35_.f25, globalState)) in (d_2610_v39_) else d_2609_v38_))
        nw432_[int(20)] = 587
        nw432_[int(21)] = (self).f20
        nw432_[int(22)] = default__.fm17((self).f20, globalState)
        d_2611_v40_ = nw432_
        d_2612_v41_: D2
        d_2612_v41_ = D2_DC5(d_2603_v32_)
        d_2612_v41_ = d_2612_v41_
        d_2613_v42_: C10
        nw433_ = C10()
        nw433_.ctor__()
        d_2613_v42_ = nw433_
        d_2614_v43_: _dafny.Set
        d_2614_v43_ = _dafny.Set({190, (self).f20})
        d_2615_v44_: _dafny.Array
        nw434_ = _dafny.Array(None, 25)
        nw434_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxr"))
        nw434_[int(1)] = d_2570_v22_
        nw434_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlbyiwl"))
        nw434_[int(3)] = (d_2570_v22_) + (_dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2616_i2_ in range(default__.abs(214))]))
        nw434_[int(4)] = _dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2617_i3_ in range(default__.abs(488))])
        nw434_[int(5)] = d_2570_v22_
        nw434_[int(6)] = d_2570_v22_
        nw434_[int(7)] = (d_2570_v22_) + (d_2570_v22_)
        nw434_[int(8)] = (d_2570_v22_).set(default__.safeIndex((self).f20, len(d_2570_v22_)), d_2573_v25_)
        nw434_[int(9)] = (d_2570_v22_) + (d_2570_v22_)
        nw434_[int(10)] = d_2570_v22_
        nw434_[int(11)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2618_i4_ in range(default__.abs(858))])
        nw434_[int(12)] = _dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2619_i5_ in range(default__.abs(757))])
        nw434_[int(13)] = d_2570_v22_
        nw434_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2620_i6_ in range(default__.abs(-511))])) + (d_2570_v22_)
        nw434_[int(15)] = d_2570_v22_
        nw434_[int(16)] = (_dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2621_i7_ in range(default__.abs(30))])) + (d_2570_v22_)
        nw434_[int(17)] = default__.fm30((self).f20, False, d_2614_v43_, p0, globalState)
        nw434_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2622_i8_ in range(default__.abs(992))]) if (d_2572_v24_)[default__.safeIndex(895, (d_2572_v24_).length(0))] else d_2570_v22_)
        nw434_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihj"))
        nw434_[int(20)] = d_2570_v22_
        nw434_[int(21)] = d_2570_v22_
        nw434_[int(22)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfq"))) + (d_2570_v22_)
        nw434_[int(23)] = d_2570_v22_
        nw434_[int(24)] = (_dafny.SeqWithoutIsStrInference([d_2573_v25_ for d_2623_i9_ in range(default__.abs(74))])) + (d_2570_v22_)
        d_2615_v44_ = nw434_
        index430_ = default__.safeIndex(162, (d_2615_v44_).length(0))
        (d_2615_v44_)[index430_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idcvyt"))) + ((d_2570_v22_).set(default__.safeIndex((self).f20, len(d_2570_v22_)), default__.fm38(self.f21, globalState)))
        index431_ = default__.safeIndex(162, (d_2615_v44_).length(0))
        (d_2615_v44_)[index431_] = (d_2570_v22_) + (d_2570_v22_)
        d_2624_v45_: _dafny.Array
        def lambda145_(d_2625_v23_):
            def lambda146_(d_2626_i11_):
                return d_2625_v23_

            return lambda146_

        init74_ = lambda145_(d_2571_v23_)
        nw435_ = _dafny.Array(None, 8)
        for i0_74_ in range(nw435_.length(0)):
            nw435_[i0_74_] = init74_(i0_74_)
        d_2624_v45_ = nw435_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2624_v45_).length(0)):
            d_2627_i10_: int = guard_loop_3_
            if (True) and (((0) <= (d_2627_i10_)) and ((d_2627_i10_) < ((d_2624_v45_).length(0)))):
                (d_2624_v45_)[(d_2627_i10_)] = d_2571_v23_
        r0 = d_2573_v25_
        return r0

    @property
    def f20(self):
        return self._f20
