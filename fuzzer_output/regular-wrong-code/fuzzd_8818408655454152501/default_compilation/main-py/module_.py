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
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmyjs"))) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_0_i0_ in range(default__.abs(743))]))

    @staticmethod
    def fm1(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-470, 70):
                d_1_v0_: int = compr_0_
                if ((-470) <= (d_1_v0_)) and ((d_1_v0_) < (70)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_1_v0_, -241)]))
            return _dafny.Set(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.Map({-550: False})).keys.Elements:
                d_3_v1_: int = compr_1_
                if (d_3_v1_) in (_dafny.Map({-550: False})):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_3_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owyrbapmv"))))]))
            return _dafny.Set(coll1_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.Set({iife0_()
        , _dafny.Set({616}), _dafny.Set({44, (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([766, 916, (_dafny.MultiSet([not(not(True))])).cardinality, 948])])).cardinality, 194}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2_i0_ in range(default__.abs(-194))])), (0) - (len(_dafny.Map({len(_dafny.Map({not(True): True})): len(_dafny.SeqWithoutIsStrInference([False]))})))})})).issubset(_dafny.Set({iife1_()
        })), (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D1_DC5()]), _dafny.SeqWithoutIsStrInference([D1_DC5()]), _dafny.SeqWithoutIsStrInference([D1_DC5() for d_4_i1_ in range(default__.abs(643))]), _dafny.SeqWithoutIsStrInference([D1_DC5()])])).issubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D1_DC5(), D1_DC5()])])), (_dafny.MultiSet([False])).issubset(_dafny.MultiSet([False])), (len(_dafny.Set({(_dafny.MultiSet([True])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "magrhjlxm")))}))) != (-600), not (not(False)) or (True)])

    @staticmethod
    def fm2(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_5_i0_ in range(default__.abs(927))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_6_i1_ in range(default__.abs(20))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))

    @staticmethod
    def fm4(p0, globalState):
        if True:
            return (_dafny.MultiSet([False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True])))
        elif True:
            return _dafny.MultiSet([not(not(not(False))), not(not(True)), not(True), False])

    @staticmethod
    def fm5(p0, globalState):
        return 563

    @staticmethod
    def fm6(globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_7_i0_ in range(default__.abs(292))])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dutfa")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qorshhqg")), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_8_i1_ in range(default__.abs(111))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqdlj"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dveovbd"))])

    @staticmethod
    def fm15(globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([745, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmo")))]))

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({244})).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "virko")))}))) | (_dafny.Set({62}))

    @staticmethod
    def fm18(globalState):
        return default__.safeDivisionInt(len(_dafny.Map({191: (D4_DC13(-484, True, False, False)).cf19})), 890)

    @staticmethod
    def fm19(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})), 166, 716, len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkqc")))]) if True else _dafny.SeqWithoutIsStrInference([79 for d_9_i0_ in range(default__.abs(22))])))

    @staticmethod
    def fm20(p0, globalState):
        source0_ = D13_DC36(_dafny.Map({not(True): not(not(False))}), True, 963)
        if source0_.is_DC36:
            d_10___mcc_h0_ = source0_.cf54
            d_11___mcc_h1_ = source0_.cf55
            d_12___mcc_h2_ = source0_.cf56
            d_13_cf56_ = d_12___mcc_h2_
            d_14_cf55_ = d_11___mcc_h1_
            d_15_cf54_ = d_10___mcc_h0_
            if (D0_DC3(len(_dafny.Map({d_13_cf56_: False})), d_14_cf55_)).cf5:
                return _dafny.Set({not(d_14_cf55_), True, True, d_14_cf55_})
            elif True:
                return _dafny.Set({d_14_cf55_, d_14_cf55_})
        elif True:
            d_16___mcc_h3_ = source0_.cf53
            d_17_cf53_ = d_16___mcc_h3_
            return (_dafny.Set({False, True})) - (_dafny.Set({True, not(False)}))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return 330

    @staticmethod
    def fm22(p0, p1, globalState):
        return D4_DC13(len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_18_i0_ in range(default__.abs(954))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvhr")))), not(True), True, True)

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        source1_ = D14_DC37(_dafny.Map({2: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_19_i0_ in range(default__.abs(200))]))}))
        if source1_.is_DC38:
            d_20___mcc_h0_ = source1_.cf58
            d_21_cf58_ = d_20___mcc_h0_
            return D1_DC5()
        elif True:
            d_22___mcc_h1_ = source1_.cf57
            d_23_cf57_ = d_22___mcc_h1_
            return D1_DC5()

    @staticmethod
    def fm27(p0, p1, globalState):
        return _dafny.Set({(938) + ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([499 for d_24_i0_ in range(default__.abs(38))]))).cardinality), 542, ((0) - (len(_dafny.Set({False, True}))) if not(True) else (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True, False]))).cardinality)})

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({False: len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([True])))}))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfdlh")))}))) | ((_dafny.Map({not(not(True)): 402}) if True else _dafny.Map({False: 227})))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(993, 142):
                d_25_v0_: int = compr_2_
                if ((993) <= (d_25_v0_)) and ((d_25_v0_) < (142)):
                    def iife3_():
                        def iife5_():
                            coll5_ = _dafny.Set()
                            compr_5_: int
                            for compr_5_ in _dafny.IntegerRange(372, 544):
                                d_28_v2_: int = compr_5_
                                if ((372) <= (d_28_v2_)) and ((d_28_v2_) < (544)):
                                    coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_28_v2_, len(_dafny.Map({True: True})))]))
                            return _dafny.Set(coll5_)
                        coll3_ = _dafny.Map()
                        def iife4_():
                            coll4_ = _dafny.Set()
                            compr_4_: int
                            for compr_4_ in _dafny.IntegerRange(372, 544):
                                d_26_v2_: int = compr_4_
                                if ((372) <= (d_26_v2_)) and ((d_26_v2_) < (544)):
                                    coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_26_v2_, len(_dafny.Map({True: True})))]))
                            return _dafny.Set(coll4_)
                        compr_3_: int
                        for compr_3_ in (iife4_()
                        ).Elements:
                            d_27_v1_: int = compr_3_
                            if (d_27_v1_) in (iife5_()
                            ):
                                coll3_[(d_27_v1_) + (len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crkbo"))})))] = False
                        return _dafny.Map(coll3_)
                    coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_25_v0_, len(iife3_()
))]))
            return _dafny.Set(coll2_)
        return ((_dafny.SeqWithoutIsStrInference([484])) + (_dafny.SeqWithoutIsStrInference([len(iife2_()
        )]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([293 for d_29_i0_ in range(default__.abs(407))]))])) + (_dafny.SeqWithoutIsStrInference([-13, (0) - ((_dafny.MultiSet([True, True])).cardinality), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjkrgyb")) for d_30_i1_ in range(default__.abs(770))])))])))

    @staticmethod
    def fm30(p0, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in (_dafny.Map({-525: not(True)})).keys.Elements:
                d_31_v0_: int = compr_6_
                if (d_31_v0_) in (_dafny.Map({-525: not(True)})):
                    coll6_ = coll6_.union(_dafny.Set([(d_31_v0_) - ((0) - (len(_dafny.Map({True: _dafny.CodePoint('q')}))))]))
            return _dafny.Set(coll6_)
        if (False if (D18_DC44(len(_dafny.SeqWithoutIsStrInference([True])), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyswmpk")), True, _dafny.SeqWithoutIsStrInference([_dafny.Set({-836}), (D4_DC11(iife6_()
)).cf18, _dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))})]))).cf66 else False):
            return D1_DC6(_dafny.CodePoint('v'), True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsxoyrdy")))
        elif True:
            return D1_DC6(_dafny.CodePoint('m'), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nv")))

    @staticmethod
    def fm31(p0, globalState):
        return -258

    @staticmethod
    def fm32(p0, globalState):
        return (_dafny.MultiSet([-755])) | (_dafny.MultiSet([537]))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(-998, 189):
                d_33_v0_: int = compr_7_
                if ((-998) <= (d_33_v0_)) and ((d_33_v0_) < (189)):
                    coll7_ = coll7_.union(_dafny.Set([(d_33_v0_) + (-224)]))
            return _dafny.Set(coll7_)
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in (_dafny.MultiSet([428, len(_dafny.SeqWithoutIsStrInference([343]))])).Elements:
                d_34_v1_: int = compr_8_
                if (d_34_v1_) in (_dafny.MultiSet([428, len(_dafny.SeqWithoutIsStrInference([343]))])):
                    coll8_ = coll8_.union(_dafny.Set([(d_34_v1_) - (-983)]))
            return _dafny.Set(coll8_)
        if (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-843])) for d_32_i0_ in range(default__.abs(95))])) < (_dafny.SeqWithoutIsStrInference([(D18_DC44(404, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdtgfltt")), not(True), _dafny.SeqWithoutIsStrInference([iife7_()
, iife8_()
]))).cf63, 373, 252, -399, -973])):
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kft"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdsnhhe")))])).Elements:
                    d_35_v2_: int = compr_9_
                    if (d_35_v2_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kft"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdsnhhe")))])):
                        coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_35_v2_, (0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(False)])).cardinality for d_36_i1_ in range(default__.abs(268))]))))]))
                return _dafny.Set(coll9_)
            return D4_DC11(iife9_()
)
        elif True:
            return D4_DC11(_dafny.Set({(_dafny.MultiSet([False, not(True)])).cardinality}))

    @staticmethod
    def fm35(p0, globalState):
        return (_dafny.Map({False: _dafny.MultiSet([True, True])})) | (_dafny.Map({False: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, not(True)]))}))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return _dafny.Map({default__.safeDivisionInt((_dafny.MultiSet([D20_DC50(D5_DC17(len(_dafny.Map({True: True})), True), True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jydkg")))])).cardinality, 365): True})

    @staticmethod
    def fm37(p0, globalState):
        source2_ = D0_DC0(False)
        if source2_.is_DC1:
            d_37___mcc_h0_ = source2_.cf1
            d_38___mcc_h1_ = source2_.cf2
            d_39_cf2_ = d_38___mcc_h1_
            d_40_cf1_ = d_37___mcc_h0_
            return _dafny.Map({d_39_cf2_: d_39_cf2_})
        elif source2_.is_DC2:
            d_41___mcc_h2_ = source2_.cf3
            d_42_cf3_ = d_41___mcc_h2_
            return (_dafny.Map({d_42_cf3_: d_42_cf3_})) | (_dafny.Map({d_42_cf3_: d_42_cf3_}))
        elif source2_.is_DC3:
            d_43___mcc_h3_ = source2_.cf4
            d_44___mcc_h4_ = source2_.cf5
            d_45_cf5_ = d_44___mcc_h4_
            d_46_cf4_ = d_43___mcc_h3_
            return (_dafny.Map({d_45_cf5_: d_45_cf5_})) | (_dafny.Map({d_45_cf5_: d_45_cf5_}))
        elif True:
            d_47___mcc_h5_ = source2_.cf0
            d_48_cf0_ = d_47___mcc_h5_
            return _dafny.Map({d_48_cf0_: d_48_cf0_})

    @staticmethod
    def fm39(p0, globalState):
        if not(False):
            return _dafny.CodePoint('s')
        elif False:
            return _dafny.CodePoint('e')
        elif True:
            return _dafny.CodePoint('l')

    @staticmethod
    def fm40(globalState):
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: _dafny.Map
            for compr_10_ in (_dafny.Set({_dafny.Map({False: 518})})).Elements:
                d_49_v0_: _dafny.Map = compr_10_
                if (d_49_v0_) in (_dafny.Set({_dafny.Map({False: 518})})):
                    coll10_ = coll10_.union(_dafny.Set([d_49_v0_]))
            return _dafny.Set(coll10_)
        return (_dafny.Set({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifrnalu")))})})) | (iife10_()
        )

    @staticmethod
    def fm41(globalState):
        return (_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference([297, 125])))})) | (_dafny.Map({True: -699}))

    @staticmethod
    def fm42(p0, globalState):
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: _dafny.Map
            for compr_11_ in (_dafny.MultiSet([_dafny.Map({False: False}), _dafny.Map({True: True}), _dafny.Map({False: True})])).Elements:
                d_50_v0_: _dafny.Map = compr_11_
                if (d_50_v0_) in (_dafny.MultiSet([_dafny.Map({False: False}), _dafny.Map({True: True}), _dafny.Map({False: True})])):
                    coll11_ = coll11_.union(_dafny.Set([d_50_v0_]))
            return _dafny.Set(coll11_)
        return D0_DC3(default__.safeDivisionInt(len(iife11_()
), 439), (D18_DC44(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbu"))), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "newukpfn")), True, _dafny.SeqWithoutIsStrInference([_dafny.Set({508}), _dafny.Set({(0) - (-353), len(_dafny.Map({_dafny.CodePoint('j'): False}))})]))).cf64)

    @staticmethod
    def fm43(p0, p1, globalState):
        return ((_dafny.Map({True: False})) | (_dafny.Map({not(not(False)): False}))) | (_dafny.Map({False: not(True)}))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return (483) * (((_dafny.MultiSet([False])).intersection(_dafny.MultiSet([False]))).cardinality)

    @staticmethod
    def fm45(p0, p1, globalState):
        return D5_DC16(False)

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        return D10_DC29((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_51_i0_ in range(default__.abs(799))]))) >= ((_dafny.MultiSet([(_dafny.MultiSet([True, False, not(False)])).cardinality])).cardinality), (_dafny.CodePoint('e')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aif"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djacvwmb"))))

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        if (True if True else True):
            if False:
                def iife12_():
                    coll12_ = _dafny.Set()
                    compr_12_: int
                    for compr_12_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({319: True}): -97})), 216])).Elements:
                        d_52_v0_: int = compr_12_
                        if (d_52_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({319: True}): -97})), 216])):
                            coll12_ = coll12_.union(_dafny.Set([(d_52_v0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_53_i0_ in range(default__.abs(476))])))]))
                    return _dafny.Set(coll12_)
                return iife12_()
                
            elif True:
                return _dafny.Set({151, 219})
        elif True:
            return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgaketdgp")))})).intersection(_dafny.Set({len(_dafny.Map({_dafny.CodePoint('f'): -13}))}))

    @staticmethod
    def fm48(p0, p1, p2, p3, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({132: _dafny.Set({False, False})})) for d_54_i0_ in range(default__.abs(-109))]))})).Elements:
                d_55_v0_: int = compr_13_
                if (d_55_v0_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({132: _dafny.Set({False, False})})) for d_54_i0_ in range(default__.abs(-109))]))})):
                    coll13_[(d_55_v0_) * (len(_dafny.Map({len(_dafny.Set({False, False})): (0) - ((_dafny.MultiSet([_dafny.CodePoint('r')])).cardinality)})))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdm")))
            return _dafny.Map(coll13_)
        return _dafny.Map({(len(iife13_()
        )) < (-844): _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('o'), _dafny.CodePoint('k')]) if False else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('v')])))})

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        return D5_DC15((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(D18_DC44(478, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rilg")), False, _dafny.SeqWithoutIsStrInference([_dafny.Set({340})]))).cf66})])), 148])).intersection(_dafny.MultiSet([-832])))

    @staticmethod
    def fm50(p0, p1, p2, p3, globalState):
        return (D15_DC39(_dafny.MultiSet([True]))).cf59

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.MultiSet([not(not(True))]), _dafny.MultiSet([True])})

    @staticmethod
    def fm52(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference([208]): 869})

    @staticmethod
    def fm53(p0, p1, p2, globalState):
        return D0_DC0(not(not((not(False)) and (True))))

    @staticmethod
    def fm54(p0, globalState):
        return D0_DC2(False)

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('r'), _dafny.CodePoint('e'), _dafny.CodePoint('e'), _dafny.CodePoint('v')])) - ((D24_DC61(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('h')])))).cf94)) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('c'), _dafny.CodePoint('s')]))) | (_dafny.MultiSet([_dafny.CodePoint('k')])))

    @staticmethod
    def fm56(p0, globalState):
        if (_dafny.MultiSet([-290])).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qd")))])), len(_dafny.SeqWithoutIsStrInference([586, 394]))])):
            if not(False):
                return D0_DC1(not(True), False)
            elif True:
                return D0_DC1(not(False), not(not(True)))
        elif True:
            return D0_DC1(False, True)

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        return _dafny.Map({(874 if not(False) else len(_dafny.SeqWithoutIsStrInference([D4_DC14(D4_DC14(D4_DC11(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kffsxfa")))})))) for d_56_i0_ in range(default__.abs(-241))]))): (0) - (default__.safeModuloInt(598, 843))})

    @staticmethod
    def fm58(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')]) for d_57_i0_ in range(default__.abs(778))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('t'), _dafny.CodePoint('g'), _dafny.CodePoint('g')]) for d_58_i1_ in range(default__.abs(627))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('p')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('q')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')]) for d_59_i2_ in range(default__.abs(524))])))

    @staticmethod
    def fm59(p0, p1, globalState):
        def iife14_():
            def iife16_():
                coll16_ = _dafny.Set()
                compr_16_: _dafny.Seq
                for compr_16_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_60_i0_ in range(default__.abs(659))])})).Elements:
                    d_63_v1_: _dafny.Seq = compr_16_
                    if (d_63_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_60_i0_ in range(default__.abs(659))])})):
                        coll16_ = coll16_.union(_dafny.Set([d_63_v1_]))
                return _dafny.Set(coll16_)
            coll14_ = _dafny.Map()
            def iife15_():
                coll15_ = _dafny.Set()
                compr_15_: _dafny.Seq
                for compr_15_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_60_i0_ in range(default__.abs(659))])})).Elements:
                    d_61_v1_: _dafny.Seq = compr_15_
                    if (d_61_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_60_i0_ in range(default__.abs(659))])})):
                        coll15_ = coll15_.union(_dafny.Set([d_61_v1_]))
                return _dafny.Set(coll15_)
            compr_14_: _dafny.Seq
            for compr_14_ in ((iife15_()
            ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cplx"))}))).Elements:
                d_62_v0_: _dafny.Seq = compr_14_
                if (d_62_v0_) in ((iife16_()
                ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cplx"))}))):
                    coll14_[d_62_v0_] = len((_dafny.Map({True: not(True)})) | (_dafny.Map({not(True): not(True)})))
            return _dafny.Map(coll14_)
        return iife14_()
        

    @staticmethod
    def fm60(globalState):
        source3_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fij")))): not(not(True))})
        d_64___mcc_h0_ = source3_
        d_65_cf60_ = d_64___mcc_h0_
        return D6_DC19(D0_DC3(len(_dafny.Map({False: False})), True), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_66_i0_ in range(default__.abs(509))])))

    @staticmethod
    def fm61(p0, p1, p2, globalState):
        def iife17_():
            coll17_ = _dafny.Set()
            compr_17_: _dafny.Set
            for compr_17_ in ((_dafny.SeqWithoutIsStrInference([_dafny.Set({-25}), _dafny.Set({-605, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahqnoab"))), 839}), _dafny.Set({-105}), _dafny.Set({len(_dafny.Map({True: 476})), len(_dafny.SeqWithoutIsStrInference([-585])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')])), 870, -522})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Map({not(True): _dafny.CodePoint('c')}))}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlcnkyi")))}), _dafny.Set({-248})]))).Elements:
                d_67_v0_: _dafny.Set = compr_17_
                if (d_67_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.Set({-25}), _dafny.Set({-605, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahqnoab"))), 839}), _dafny.Set({-105}), _dafny.Set({len(_dafny.Map({True: 476})), len(_dafny.SeqWithoutIsStrInference([-585])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p')])), 870, -522})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Map({not(True): _dafny.CodePoint('c')}))}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlcnkyi")))}), _dafny.Set({-248})]))):
                    coll17_ = coll17_.union(_dafny.Set([d_67_v0_]))
            return _dafny.Set(coll17_)
        return iife17_()
        

    @staticmethod
    def fm62(globalState):
        return ((_dafny.Map({False: D10_DC28(_dafny.Map({_dafny.CodePoint('x'): True}))})) | (_dafny.Map({False: D10_DC28(_dafny.Map({_dafny.CodePoint('m'): True}))}))) | ((_dafny.Map({not(True): D10_DC28(_dafny.Map({_dafny.CodePoint('p'): False}))}) if True else _dafny.Map({True: D10_DC28(_dafny.Map({_dafny.CodePoint('a'): True}))})))

    @staticmethod
    def fm63(p0, p1, p2, p3, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_69_i1_ in range(default__.abs(772))])), len(_dafny.SeqWithoutIsStrInference([False]))])), len(_dafny.Map({False: -862}))])).Elements:
                d_70_v0_: int = compr_18_
                if (d_70_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_69_i1_ in range(default__.abs(772))])), len(_dafny.SeqWithoutIsStrInference([False]))])), len(_dafny.Map({False: -862}))])):
                    coll18_[(d_70_v0_) + (len(_dafny.Set({True, False})))] = not(True)
            return _dafny.Map(coll18_)
        return (_dafny.MultiSet([_dafny.Set({490, 278}), _dafny.Set({-498, 752}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))), (0) - ((0) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([956]): True}))))}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([-118 for d_68_i0_ in range(default__.abs(178))]))})])) | (_dafny.MultiSet([_dafny.Set({-822}), _dafny.Set({(D0_DC3(len(iife18_()
), True)).cf4}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyskdkmkh"))), len(_dafny.Map({len(_dafny.Map({-82: (0) - ((0) - (len(_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faso"))): True})}))))})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huu"))}))})]))

    @staticmethod
    def fm64(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.CodePoint('r'), _dafny.CodePoint('v'), _dafny.CodePoint('j'), _dafny.CodePoint('h'), _dafny.CodePoint('f')})

    @staticmethod
    def fm65(p0, globalState):
        if False:
            if False:
                return _dafny.Map({384: not(False)})
            elif True:
                return _dafny.Map({len(_dafny.Map({not(False): (_dafny.MultiSet([414, 657])).cardinality})): not(True)})
        elif True:
            return _dafny.Map({len(_dafny.Set({True})): True})

    @staticmethod
    def fm66(globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in (_dafny.SeqWithoutIsStrInference([526])).Elements:
                d_71_v0_: int = compr_19_
                if (d_71_v0_) in (_dafny.SeqWithoutIsStrInference([526])):
                    coll19_[(d_71_v0_) - (991)] = _dafny.CodePoint('f')
            return _dafny.Map(coll19_)
        source4_ = _dafny.Map({len(iife19_()
        ): False})
        d_72___mcc_h0_ = source4_
        d_73_cf60_ = d_72___mcc_h0_
        return D19_DC48((0) - (-238))

    @staticmethod
    def fm67(p0, p1, p2, globalState):
        source5_ = D24_DC63(562, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bikakvgqv")), 847, 174)
        if source5_.is_DC62:
            d_74___mcc_h0_ = source5_.cf95
            d_75_cf95_ = d_74___mcc_h0_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqrsg"))})
        elif source5_.is_DC63:
            d_76___mcc_h1_ = source5_.cf96
            d_77___mcc_h2_ = source5_.cf97
            d_78___mcc_h3_ = source5_.cf98
            d_79___mcc_h4_ = source5_.cf99
            d_80_cf99_ = d_79___mcc_h4_
            d_81_cf98_ = d_78___mcc_h3_
            d_82_cf97_ = d_77___mcc_h2_
            d_83_cf96_ = d_76___mcc_h1_
            def iife20_():
                coll20_ = _dafny.Set()
                compr_20_: _dafny.Seq
                for compr_20_ in (_dafny.SeqWithoutIsStrInference([d_82_cf97_])).Elements:
                    d_84_v0_: _dafny.Seq = compr_20_
                    if (d_84_v0_) in (_dafny.SeqWithoutIsStrInference([d_82_cf97_])):
                        coll20_ = coll20_.union(_dafny.Set([d_84_v0_]))
                return _dafny.Set(coll20_)
            return iife20_()
            
        elif source5_.is_DC64:
            d_85___mcc_h5_ = source5_.cf100
            d_86_cf100_ = d_85___mcc_h5_
            def iife21_():
                coll21_ = _dafny.Set()
                compr_21_: _dafny.Seq
                for compr_21_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_87_i0_ in range(default__.abs(825))])])).Elements:
                    d_88_v1_: _dafny.Seq = compr_21_
                    if (d_88_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_87_i0_ in range(default__.abs(825))])])):
                        coll21_ = coll21_.union(_dafny.Set([d_88_v1_]))
                return _dafny.Set(coll21_)
            return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgkdbmw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coc"))})) - (iife21_()
            )
        elif True:
            d_89___mcc_h6_ = source5_.cf94
            d_90_cf94_ = d_89___mcc_h6_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_91_i1_ in range(default__.abs(318))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isxpgybe"))})

    @staticmethod
    def fm68(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.SeqWithoutIsStrInference([True, False, True])) for d_92_i0_ in range(default__.abs(145))])) + (_dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.SeqWithoutIsStrInference([True, True]))]))) + ((_dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.SeqWithoutIsStrInference([True]))])) + (_dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.SeqWithoutIsStrInference([False, True]))])))

    @staticmethod
    def fm69(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, not(not(True))]) for d_93_i0_ in range(default__.abs(-685))]))

    @staticmethod
    def fm70(p0, p1, p2, globalState):
        source6_ = D4_DC14(D4_DC11(_dafny.Set({(_dafny.MultiSet([len(_dafny.Set({528})), 191])).cardinality})))
        if source6_.is_DC12:
            return D14_DC38(len(_dafny.Set({60})))
        elif source6_.is_DC13:
            d_94___mcc_h0_ = source6_.cf19
            d_95___mcc_h1_ = source6_.cf20
            d_96___mcc_h2_ = source6_.cf21
            d_97___mcc_h3_ = source6_.cf22
            d_98_cf22_ = d_97___mcc_h3_
            d_99_cf21_ = d_96___mcc_h2_
            d_100_cf20_ = d_95___mcc_h1_
            d_101_cf19_ = d_94___mcc_h0_
            return D14_DC38(d_101_cf19_)
        elif source6_.is_DC11:
            d_102___mcc_h4_ = source6_.cf18
            d_103_cf18_ = d_102___mcc_h4_
            return D14_DC38(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggt"))))
        elif True:
            d_104___mcc_h5_ = source6_.cf23
            d_105_cf23_ = d_104___mcc_h5_
            return D14_DC38(147)

    @staticmethod
    def fm71(p0, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(False), False]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), not(not(False)), True, True, False])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))]))) == ((D13_DC36(_dafny.Map({not(True): False}), not(False), 749)).cf56):
            return D1_DC4(_dafny.SeqWithoutIsStrInference([True, True, True, not(True), False]))
        elif True:
            return D1_DC4(_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm72(p0, p1, p2, p3, globalState):
        source7_ = D15_DC40()
        if source7_.is_DC40:
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([677]))]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})) for d_106_i0_ in range(default__.abs(860))]))), 175])])
        elif True:
            d_107___mcc_h0_ = source7_.cf59
            d_108_cf59_ = d_107___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D10_DC29(True, True, len(_dafny.Map({len(_dafny.Map({False: 276})): _dafny.CodePoint('h')})))])) for d_109_i1_ in range(default__.abs(606))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_110_i2_ in range(default__.abs(-534))]))])]))

    @staticmethod
    def fm73(p0, globalState):
        return D12_DC32((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (0) - (len(_dafny.Map({len(_dafny.Set({310, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), 131]))})): not(False)})))}))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vowu"))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, len(_dafny.SeqWithoutIsStrInference([False]))])) < ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))]))))

    @staticmethod
    def fm74(globalState):
        return _dafny.SeqWithoutIsStrInference([D22_DC56(_dafny.CodePoint('k'), True)])

    @staticmethod
    def m0(globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        d_111_v0_: bool
        d_111_v0_ = True
        d_112_v1_: int
        d_112_v1_ = 740
        d_113_v2_: _dafny.Map
        d_113_v2_ = _dafny.Map({not (d_111_v0_) or (d_111_v0_): (d_112_v1_ if d_111_v0_ else (0) - (d_112_v1_))})
        d_114_v3_: _dafny.Array
        nw0_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
        d_114_v3_ = nw0_
        d_115_v4_: _dafny.Seq
        d_115_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwvaps"))
        index0_ = default__.safeIndex(360, (d_114_v3_).length(0))
        (d_114_v3_)[index0_] = d_115_v4_
        d_116_v5_: str
        d_116_v5_ = _dafny.CodePoint('e')
        index1_ = default__.safeIndex(360, (d_114_v3_).length(0))
        rhs0_ = d_111_v0_
        rhs1_ = d_115_v4_
        rhs2_ = (d_113_v2_) | (d_113_v2_)
        rhs3_ = ((d_115_v4_) + (d_115_v4_)).set(default__.safeIndex((0) - (len(d_115_v4_)), len((d_115_v4_) + (d_115_v4_))), d_116_v5_)
        rhs4_ = d_112_v1_
        lhs0_ = globalState
        lhs1_ = globalState
        lhs2_ = d_114_v3_
        lhs3_ = default__.safeIndex(360, (d_114_v3_).length(0))
        lhs0_.f12 = rhs0_
        lhs1_.f10 = rhs1_
        d_113_v2_ = rhs2_
        lhs2_[lhs3_] = rhs3_
        d_112_v1_ = rhs4_
        d_117_i0_: int
        d_117_i0_ = 0
        with _dafny.label("0"):
            while default__.fm0(default__.fm5(d_111_v0_, globalState), globalState):
                with _dafny.c_label("0"):
                    if (d_117_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_117_i0_ = (d_117_i0_) + (1)
                    d_118_v6_: _dafny.Array
                    def lambda0_(d_119_v1_):
                        def lambda1_(d_120_i1_):
                            return (d_120_i1_) - (d_119_v1_)

                        return lambda1_

                    init0_ = lambda0_(d_112_v1_)
                    nw1_ = _dafny.Array(None, 20)
                    for i0_0_ in range(nw1_.length(0)):
                        nw1_[i0_0_] = init0_(i0_0_)
                    d_118_v6_ = nw1_
                    index2_ = default__.safeIndex(244, (d_118_v6_).length(0))
                    (d_118_v6_)[index2_] = (d_112_v1_) - (d_112_v1_)
                    index3_ = default__.safeIndex(244, (d_118_v6_).length(0))
                    (d_118_v6_)[index3_] = (d_112_v1_) * ((-485) + (d_112_v1_))
                    d_115_v4_ = (d_114_v3_)[default__.safeIndex(360, (d_114_v3_).length(0))]
                    d_121_v7_: _dafny.Array
                    nw2_ = _dafny.Array(None, 14)
                    d_121_v7_ = nw2_
                    d_121_v7_ = d_121_v7_
                    d_122_v8_: _dafny.MultiSet
                    d_122_v8_ = _dafny.MultiSet([458])
                    d_123_v9_: _dafny.Seq
                    d_123_v9_ = _dafny.SeqWithoutIsStrInference([d_122_v8_])
                    d_124_v10_: D23
                    d_124_v10_ = D23_DC58(d_123_v9_)
                    (globalState).f12 = ((d_123_v9_) + (d_123_v9_)) == ((d_124_v10_).cf90)
                    pass
            pass
        d_125_v11_: _dafny.MultiSet
        d_125_v11_ = _dafny.MultiSet([False, d_111_v0_])
        d_126_v12_: T1
        nw3_ = C12()
        nw3_.ctor__(d_112_v1_, not(d_111_v0_), ((d_125_v11_)[not(default__.fm0(318, globalState))] if (not(default__.fm0(318, globalState))) in (d_125_v11_) else d_112_v1_), d_114_v3_)
        d_126_v12_ = nw3_
        d_127_v13_: D22
        d_127_v13_ = D22_DC55((d_126_v12_).f26)
        nw4_ = C6()
        nw4_.ctor__(d_112_v1_, (d_127_v13_).cf86)
        d_126_v12_ = nw4_
        d_128_v14_: _dafny.Array
        nw5_ = _dafny.Array(False, 29)
        d_128_v14_ = nw5_
        r0 = d_128_v14_
        d_129_i2_: int
        d_129_i2_ = 0
        with _dafny.label("1"):
            while d_111_v0_:
                with _dafny.c_label("1"):
                    if (d_129_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_129_i2_ = (d_129_i2_) + (1)
                    d_112_v1_ = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_130_i3_ in range(default__.abs(628))])))
                    d_131_v15_: _dafny.MultiSet
                    d_131_v15_ = _dafny.MultiSet([(0) - (d_112_v1_), default__.safeDivisionInt(d_112_v1_, 321), d_112_v1_])
                    d_131_v15_ = d_131_v15_
                    index4_ = default__.safeIndex(149, (d_128_v14_).length(0))
                    (d_128_v14_)[index4_] = d_111_v0_
                    index5_ = default__.safeIndex(149, (d_128_v14_).length(0))
                    (d_128_v14_)[index5_] = d_111_v0_
                    d_132_v16_: C12
                    nw6_ = C12()
                    nw6_.ctor__((d_126_v12_).f25, not(d_111_v0_), 432, (d_126_v12_).f26)
                    d_132_v16_ = nw6_
                    d_132_v16_ = d_132_v16_
                    pass
            pass
        d_133_v17_: C0
        nw7_ = C0()
        nw7_.ctor__()
        d_133_v17_ = nw7_
        r0 = d_128_v14_
        def lambda2_(d_134_v12_):
            def lambda3_(d_135_i4_):
                return (d_135_i4_) * ((d_134_v12_).f25)

            return lambda3_

        init1_ = lambda2_(d_126_v12_)
        nw8_ = _dafny.Array(None, 13)
        for i0_1_ in range(nw8_.length(0)):
            nw8_[i0_1_] = init1_(i0_1_)
        r1 = nw8_
        r2 = ((d_114_v3_)[default__.safeIndex(360, (d_114_v3_).length(0))]) != (d_115_v4_)
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_136_v0_: _dafny.Array
        def lambda4_(d_137_i0_):
            return False

        init2_ = lambda4_
        nw9_ = _dafny.Array(None, 27)
        for i0_2_ in range(nw9_.length(0)):
            nw9_[i0_2_] = init2_(i0_2_)
        d_136_v0_ = nw9_
        d_138_v1_: int
        d_138_v1_ = 454
        d_139_v2_: _dafny.Seq
        d_139_v2_ = _dafny.SeqWithoutIsStrInference([d_138_v1_])
        d_140_v3_: _dafny.Seq
        d_140_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvkwml"))
        d_141_v4_: _dafny.Seq
        d_141_v4_ = _dafny.SeqWithoutIsStrInference([d_140_v3_])
        d_142_v5_: str
        d_142_v5_ = _dafny.CodePoint('k')
        d_143_v6_: _dafny.MultiSet
        d_143_v6_ = _dafny.MultiSet([d_140_v3_])
        d_144_v7_: _dafny.Array
        def lambda5_(d_145_v1_):
            def lambda6_(d_146_i1_):
                return _dafny.Map({d_145_v1_: (D0_DC2(True)).cf3})

            return lambda6_

        init3_ = lambda5_(d_138_v1_)
        nw10_ = _dafny.Array(None, 23)
        for i0_3_ in range(nw10_.length(0)):
            nw10_[i0_3_] = init3_(i0_3_)
        d_144_v7_ = nw10_
        d_147_v8_: bool
        d_147_v8_ = False
        d_148_v9_: _dafny.MultiSet
        d_148_v9_ = _dafny.MultiSet([d_147_v8_, d_147_v8_])
        d_149_v10_: _dafny.Map
        d_149_v10_ = _dafny.Map({249: d_148_v9_})
        d_150_v11_: _dafny.Set
        d_150_v11_ = _dafny.Set({d_147_v8_, d_147_v8_})
        d_151_globalState_: GlobalState
        nw11_ = GlobalState()
        nw11_.ctor__(286, 656, True, 606, d_136_v0_, 150, d_139_v2_, 55, 667, d_140_v3_, ((d_141_v4_)[default__.safeIndex(d_138_v1_, len(d_141_v4_))]) + ((d_140_v3_).set(default__.safeIndex(478, len(d_140_v3_)), d_142_v5_)), True, True, False, False, True, (d_143_v6_) | (d_143_v6_), d_144_v7_, d_149_v10_, True, d_150_v11_, 517, 61, _dafny.CodePoint('p'))
        d_151_globalState_ = nw11_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_136_v0_).length(0)):
            d_152_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_152_i2_)) and ((d_152_i2_) < ((d_136_v0_).length(0)))):
                (d_136_v0_)[(d_152_i2_)] = d_147_v8_
        d_153_v12_: _dafny.Map
        d_153_v12_ = _dafny.Map({False: default__.fm0(d_138_v1_, d_151_globalState_)})
        d_154_v13_: _dafny.Map
        d_154_v13_ = _dafny.Map({True: len(d_153_v12_)})
        d_155_v14_: _dafny.Map
        d_155_v14_ = _dafny.Map({d_138_v1_: (d_154_v13_) | (_dafny.Map({d_147_v8_: len(d_140_v3_)}))})
        d_156_v16_: D0
        d_156_v16_ = D0_DC3(d_138_v1_, d_147_v8_)
        d_157_v17_: _dafny.MultiSet
        d_157_v17_ = _dafny.MultiSet([(d_156_v16_).cf4, d_138_v1_])
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in (d_157_v17_).Elements:
                d_158_v15_: int = compr_22_
                if (d_158_v15_) in (d_157_v17_):
                    coll22_[default__.safeModuloInt(d_158_v15_, len(_dafny.SeqWithoutIsStrInference([d_147_v8_, d_147_v8_, d_147_v8_, d_147_v8_])))] = d_154_v13_
            return _dafny.Map(coll22_)
        d_155_v14_ = iife22_()
        
        d_159_v18_: _dafny.Seq
        d_159_v18_ = _dafny.SeqWithoutIsStrInference([d_147_v8_])
        d_160_v19_: _dafny.Seq
        d_160_v19_ = _dafny.SeqWithoutIsStrInference([d_159_v18_])
        d_161_i3_: int
        d_161_i3_ = 0
        with _dafny.label("2"):
            while not((_dafny.MultiSet(d_160_v19_)).issubset(_dafny.MultiSet([d_159_v18_, default__.fm1((d_159_v18_)[default__.safeIndex(-684, len(d_159_v18_))], d_151_globalState_)]))):
                with _dafny.c_label("2"):
                    if (d_161_i3_) >= (100):
                        raise _dafny.Break("2")
                    d_161_i3_ = (d_161_i3_) + (1)
                    (d_151_globalState_).f23 = d_142_v5_
                    index6_ = default__.safeIndex(176, (d_136_v0_).length(0))
                    (d_136_v0_)[index6_] = d_147_v8_
                    index7_ = default__.safeIndex(176, (d_136_v0_).length(0))
                    (d_136_v0_)[index7_] = (d_159_v18_) != (d_159_v18_)
                    d_162_v20_: _dafny.Set
                    d_162_v20_ = _dafny.Set({d_138_v1_})
                    d_138_v1_ = ((d_139_v2_)[default__.safeIndex(d_138_v1_, len(d_139_v2_))]) + ((len(d_162_v20_)) * (d_138_v1_))
                    d_147_v8_ = (d_136_v0_)[default__.safeIndex(176, (d_136_v0_).length(0))]
                    pass
            pass
        if (len((d_140_v3_) + (default__.fm2(d_138_v1_, d_138_v1_, d_151_globalState_)))) == (len(d_139_v2_)):
            d_163_v21_: D0
            d_163_v21_ = D0_DC1(d_147_v8_, d_147_v8_)
            d_164_v22_: _dafny.Set
            d_164_v22_ = _dafny.Set({d_163_v21_})
            (d_151_globalState_).f12 = ((d_164_v22_).intersection(d_164_v22_)).issubset(d_164_v22_)
            d_138_v1_ = default__.safeDivisionInt(d_138_v1_, 347)
            index8_ = default__.safeIndex(988, (d_136_v0_).length(0))
            (d_136_v0_)[index8_] = (d_147_v8_ if d_147_v8_ else d_147_v8_)
            d_165_v23_: D1
            d_165_v23_ = D1_DC4(d_159_v18_)
            index9_ = default__.safeIndex(988, (d_136_v0_).length(0))
            (d_136_v0_)[index9_] = ((d_140_v3_)[default__.safeIndex(len((d_165_v23_).cf6), len(d_140_v3_))]) in (d_140_v3_)
            d_138_v1_ = (((d_157_v17_)[d_138_v1_] if (d_138_v1_) in (d_157_v17_) else d_138_v1_) if d_147_v8_ else d_138_v1_)
            d_166_v24_: D1
            d_166_v24_ = D1_DC6(d_142_v5_, (d_136_v0_)[default__.safeIndex(988, (d_136_v0_).length(0))], d_140_v3_)
            d_167_v25_: _dafny.Map
            d_167_v25_ = _dafny.Map({default__.fm0(d_138_v1_, d_151_globalState_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qujrkn"))})
            d_168_v26_: _dafny.Array
            nw12_ = _dafny.Array(None, 20)
            nw12_[int(0)] = (_dafny.SeqWithoutIsStrInference([d_142_v5_ for d_169_i4_ in range(default__.abs(78))])) + (d_140_v3_)
            nw12_[int(1)] = d_140_v3_
            nw12_[int(2)] = d_140_v3_
            nw12_[int(3)] = (d_166_v24_).cf9
            nw12_[int(4)] = (d_140_v3_) + (d_140_v3_)
            nw12_[int(5)] = (d_140_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dakda")))
            nw12_[int(6)] = d_140_v3_
            nw12_[int(7)] = d_140_v3_
            nw12_[int(8)] = d_140_v3_
            nw12_[int(9)] = d_140_v3_
            nw12_[int(10)] = d_140_v3_
            nw12_[int(11)] = (default__.fm2(d_138_v1_, d_138_v1_, d_151_globalState_) if d_147_v8_ else default__.fm2(d_138_v1_, 159, d_151_globalState_))
            nw12_[int(12)] = d_140_v3_
            nw12_[int(13)] = d_140_v3_
            nw12_[int(14)] = d_140_v3_
            nw12_[int(15)] = (((d_167_v25_)[(d_136_v0_)[default__.safeIndex(988, (d_136_v0_).length(0))]] if ((d_136_v0_)[default__.safeIndex(988, (d_136_v0_).length(0))]) in (d_167_v25_) else d_140_v3_)) + (d_140_v3_)
            nw12_[int(16)] = d_140_v3_
            nw12_[int(17)] = _dafny.SeqWithoutIsStrInference([d_142_v5_ for d_170_i5_ in range(default__.abs(-947))])
            nw12_[int(18)] = _dafny.SeqWithoutIsStrInference([d_142_v5_ for d_171_i6_ in range(default__.abs(-516))])
            nw12_[int(19)] = d_140_v3_
            d_168_v26_ = nw12_
            index10_ = default__.safeIndex(647, (d_168_v26_).length(0))
            (d_168_v26_)[index10_] = d_140_v3_
            index11_ = default__.safeIndex(647, (d_168_v26_).length(0))
            (d_168_v26_)[index11_] = (d_140_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otnycwpnk")))
        elif True:
            d_172_v27_: _dafny.Array
            nw13_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_172_v27_ = nw13_
            d_172_v27_ = (d_172_v27_ if default__.fm0(len(d_159_v18_), d_151_globalState_) else d_172_v27_)
            d_173_v28_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.Map({}), 11)
            d_173_v28_ = nw14_
            d_174_v29_: _dafny.Array
            nw15_ = _dafny.Array(_dafny.CodePoint('D'), 5)
            d_174_v29_ = nw15_
            d_175_v30_: _dafny.Map
            d_175_v30_ = _dafny.Map({d_173_v28_: (d_174_v29_ if False else d_174_v29_)})
            d_175_v30_ = (d_175_v30_).set(d_173_v28_, d_174_v29_)
            d_176_v31_: _dafny.Array
            d_177_v32_: _dafny.Array
            d_178_v33_: bool
            out0_: _dafny.Array
            out1_: _dafny.Array
            out2_: bool
            out0_, out1_, out2_ = default__.m0(d_151_globalState_)
            d_176_v31_ = out0_
            d_177_v32_ = out1_
            d_178_v33_ = out2_
            if (d_147_v8_) == (False):
                d_138_v1_ = (len(d_139_v2_) if not((d_138_v1_) != (d_138_v1_)) else d_138_v1_)
                d_179_v34_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
                d_179_v34_ = nw16_
                d_180_v35_: C4
                nw17_ = C4()
                nw17_.ctor__(d_179_v34_, d_138_v1_, d_179_v34_)
                d_180_v35_ = nw17_
                d_181_v36_: _dafny.Map
                d_181_v36_ = _dafny.Map({d_142_v5_: d_147_v8_})
                d_181_v36_ = (d_181_v36_).set(d_142_v5_, (d_178_v33_) == (d_147_v8_))
                d_182_v37_: D19
                d_182_v37_ = D19_DC47(d_136_v0_, d_147_v8_)
                d_183_v38_: D3
                d_183_v38_ = D3_DC10(d_178_v33_, d_136_v0_, d_178_v33_, d_138_v1_, d_138_v1_)
                rhs5_ = False
                rhs6_ = default__.fm0((d_183_v38_).cf17, d_151_globalState_)
                rhs7_ = d_182_v37_
                lhs4_ = d_151_globalState_
                d_178_v33_ = rhs5_
                lhs4_.f12 = rhs6_
                d_182_v37_ = rhs7_
                index12_ = default__.safeIndex(525, (d_136_v0_).length(0))
                (d_136_v0_)[index12_] = d_178_v33_
                d_184_v39_: _dafny.Map
                d_184_v39_ = _dafny.Map({d_142_v5_: (0) - ((d_139_v2_)[default__.safeIndex(d_138_v1_, len(d_139_v2_))])})
                index13_ = default__.safeIndex(28, (d_177_v32_).length(0))
                (d_177_v32_)[index13_] = (0) - (len(d_184_v39_))
                index14_ = default__.safeIndex(525, (d_136_v0_).length(0))
                index15_ = default__.safeIndex(28, (d_177_v32_).length(0))
                rhs8_ = not (d_147_v8_) or ((d_138_v1_) == (d_138_v1_))
                rhs9_ = d_174_v29_
                rhs10_ = default__.safeModuloInt(562, d_138_v1_)
                lhs5_ = d_136_v0_
                lhs6_ = default__.safeIndex(525, (d_136_v0_).length(0))
                lhs7_ = d_177_v32_
                lhs8_ = default__.safeIndex(28, (d_177_v32_).length(0))
                lhs5_[lhs6_] = rhs8_
                d_174_v29_ = rhs9_
                lhs7_[lhs8_] = rhs10_
            elif True:
                index16_ = default__.safeIndex(194, (d_136_v0_).length(0))
                (d_136_v0_)[index16_] = True
                index17_ = default__.safeIndex(194, (d_136_v0_).length(0))
                rhs11_ = (d_138_v1_) != (default__.safeDivisionInt(d_138_v1_, len(d_140_v3_)))
                rhs12_ = default__.fm31(d_140_v3_, d_151_globalState_)
                lhs9_ = d_136_v0_
                lhs10_ = default__.safeIndex(194, (d_136_v0_).length(0))
                lhs9_[lhs10_] = rhs11_
                d_138_v1_ = rhs12_
                d_177_v32_ = d_177_v32_
                d_138_v1_ = (d_138_v1_) - ((d_138_v1_) - ((0) - (len(d_140_v3_))))
                d_185_v40_: _dafny.MultiSet
                d_185_v40_ = _dafny.MultiSet([d_142_v5_])
                index18_ = default__.safeIndex(194, (d_136_v0_).length(0))
                index19_ = default__.safeIndex(194, (d_136_v0_).length(0))
                rhs13_ = False
                rhs14_ = d_147_v8_
                rhs15_ = (0) - (len(d_159_v18_))
                rhs16_ = default__.fm39((d_140_v3_) < (d_140_v3_), d_151_globalState_)
                rhs17_ = default__.safeModuloInt(d_138_v1_, ((d_185_v40_)[d_142_v5_] if (d_142_v5_) in (d_185_v40_) else d_138_v1_))
                lhs11_ = d_136_v0_
                lhs12_ = default__.safeIndex(194, (d_136_v0_).length(0))
                lhs13_ = d_136_v0_
                lhs14_ = default__.safeIndex(194, (d_136_v0_).length(0))
                lhs15_ = d_151_globalState_
                lhs11_[lhs12_] = rhs13_
                lhs13_[lhs14_] = rhs14_
                d_138_v1_ = rhs15_
                lhs15_.f23 = rhs16_
                d_138_v1_ = rhs17_
                index20_ = default__.safeIndex(194, (d_136_v0_).length(0))
                (d_136_v0_)[index20_] = d_147_v8_
            d_186_v41_: D5
            d_186_v41_ = D5_DC15(d_157_v17_)
            d_187_v42_: _dafny.Map
            d_187_v42_ = _dafny.Map({796: d_186_v41_})
            d_187_v42_ = (d_187_v42_).set((d_139_v2_)[default__.safeIndex(len(d_150_v11_), len(d_139_v2_))], d_186_v41_)
        d_188_v43_: _dafny.Map
        d_188_v43_ = _dafny.Map({len(d_139_v2_): d_147_v8_})
        d_189_i7_: int
        d_189_i7_ = 0
        with _dafny.label("3"):
            while ((d_188_v43_)[d_138_v1_] if (d_138_v1_) in (d_188_v43_) else default__.fm0(-719, d_151_globalState_)):
                with _dafny.c_label("3"):
                    if (d_189_i7_) >= (100):
                        raise _dafny.Break("3")
                    d_189_i7_ = (d_189_i7_) + (1)
                    d_190_v44_: _dafny.Array
                    nw18_ = _dafny.Array(_dafny.Map({}), 16)
                    d_190_v44_ = nw18_
                    d_191_v45_: _dafny.Map
                    d_191_v45_ = _dafny.Map({d_140_v3_: d_138_v1_})
                    index21_ = default__.safeIndex(251, (d_190_v44_).length(0))
                    (d_190_v44_)[index21_] = (d_191_v45_ if False else d_191_v45_)
                    index22_ = default__.safeIndex(251, (d_190_v44_).length(0))
                    (d_190_v44_)[index22_] = (d_191_v45_).set((d_140_v3_) + (d_140_v3_), len((_dafny.SeqWithoutIsStrInference([d_138_v1_ for d_192_i8_ in range(default__.abs(-743))])) + (d_139_v2_)))
                    d_147_v8_ = (d_150_v11_).isdisjoint((default__.fm20(d_147_v8_, d_151_globalState_)) - (d_150_v11_))
                    d_193_v46_: D19
                    d_193_v46_ = D19_DC47(d_136_v0_, (d_159_v18_)[default__.safeIndex(d_138_v1_, len(d_159_v18_))])
                    pat_let_tv0_ = d_136_v0_
                    def iife23_(_pat_let0_0):
                        def iife24_(d_194_dt__update__tmp_h0_):
                            def iife25_(_pat_let1_0):
                                def iife26_(d_195_dt__update_hcf69_h0_):
                                    return D19_DC47(d_195_dt__update_hcf69_h0_, (d_194_dt__update__tmp_h0_).cf70)
                                return iife26_(_pat_let1_0)
                            return iife25_(pat_let_tv0_)
                        return iife24_(_pat_let0_0)
                    source8_ = iife23_(d_193_v46_)
                    if source8_.is_DC47:
                        d_196___mcc_h0_ = source8_.cf69
                        d_197___mcc_h1_ = source8_.cf70
                        d_198_cf70_ = d_197___mcc_h1_
                        d_199_cf69_ = d_196___mcc_h0_
                        d_200_v47_: D0
                        d_200_v47_ = D0_DC1(d_198_cf70_, d_198_cf70_)
                        d_201_v48_: _dafny.Map
                        d_201_v48_ = _dafny.Map({d_200_v47_: (d_140_v3_)[default__.safeIndex(default__.fm31((default__.fm2(d_138_v1_, d_138_v1_, d_151_globalState_)).set(default__.safeIndex(len(d_159_v18_), len(default__.fm2(d_138_v1_, d_138_v1_, d_151_globalState_))), d_142_v5_), d_151_globalState_), len(d_140_v3_))]})
                        d_201_v48_ = (d_201_v48_).set(d_200_v47_, d_142_v5_)
                        d_138_v1_ = d_138_v1_
                        d_202_v49_: D15
                        d_202_v49_ = D15_DC40()
                        d_202_v49_ = d_202_v49_
                        (d_151_globalState_).f10 = d_140_v3_
                    elif source8_.is_DC48:
                        d_203___mcc_h2_ = source8_.cf71
                        d_204_cf71_ = d_203___mcc_h2_
                        d_205_v50_: D0
                        d_205_v50_ = D0_DC0(not(d_147_v8_))
                        d_206_v51_: _dafny.Map
                        d_206_v51_ = _dafny.Map({d_205_v50_: d_147_v8_})
                        d_207_v52_: _dafny.Array
                        nw19_ = _dafny.Array(None, 9)
                        nw19_[int(0)] = (default__.fm21(d_206_v51_, d_138_v1_, d_147_v8_, d_140_v3_, d_151_globalState_)) - (d_204_cf71_)
                        nw19_[int(1)] = 488
                        nw19_[int(2)] = d_204_cf71_
                        nw19_[int(3)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgha"))).set(default__.safeIndex(d_204_cf71_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgha")))), _dafny.CodePoint('c')))
                        nw19_[int(4)] = d_204_cf71_
                        nw19_[int(5)] = 666
                        nw19_[int(6)] = 273
                        nw19_[int(7)] = d_138_v1_
                        nw19_[int(8)] = (d_138_v1_) - (d_204_cf71_)
                        d_207_v52_ = nw19_
                        d_207_v52_ = d_207_v52_
                        d_208_v53_: _dafny.Map
                        d_208_v53_ = d_188_v43_
                        d_209_v54_: _dafny.Array
                        nw20_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                        d_209_v54_ = nw20_
                        d_210_v55_: D22
                        d_210_v55_ = D22_DC55(d_209_v54_)
                        d_211_v56_: T0
                        nw21_ = C10()
                        nw21_.ctor__(d_142_v5_, 114, (d_210_v55_).cf86)
                        d_211_v56_ = nw21_
                        rhs18_ = (_dafny.Map({(0) - (d_138_v1_): d_147_v8_})) | (d_188_v43_)
                        rhs19_ = d_211_v56_
                        d_208_v53_ = rhs18_
                        d_211_v56_ = rhs19_
                        d_212_v57_: _dafny.Array
                        d_213_v58_: _dafny.Array
                        d_214_v59_: bool
                        out3_: _dafny.Array
                        out4_: _dafny.Array
                        out5_: bool
                        out3_, out4_, out5_ = default__.m0(d_151_globalState_)
                        d_212_v57_ = out3_
                        d_213_v58_ = out4_
                        d_214_v59_ = out5_
                        d_204_cf71_ = len((d_159_v18_ if (d_204_cf71_) == (d_204_cf71_) else d_159_v18_))
                    elif True:
                        d_215___mcc_h3_ = source8_.cf68
                        d_216_cf68_ = d_215___mcc_h3_
                        index23_ = default__.safeIndex(191, (d_136_v0_).length(0))
                        (d_136_v0_)[index23_] = not (d_147_v8_) or (d_147_v8_)
                        index24_ = default__.safeIndex(191, (d_136_v0_).length(0))
                        (d_136_v0_)[index24_] = d_147_v8_
                        d_217_v60_: _dafny.Array
                        nw22_ = _dafny.Array(int(0), 17)
                        d_217_v60_ = nw22_
                        rhs20_ = d_217_v60_
                        rhs21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkcekenj"))
                        rhs22_ = default__.safeDivisionInt((d_138_v1_) - ((0) - (default__.fm44((d_136_v0_)[default__.safeIndex(191, (d_136_v0_).length(0))], d_147_v8_, (d_136_v0_)[default__.safeIndex(191, (d_136_v0_).length(0))], d_138_v1_, d_151_globalState_))), d_138_v1_)
                        lhs16_ = d_151_globalState_
                        d_217_v60_ = rhs20_
                        lhs16_.f10 = rhs21_
                        d_138_v1_ = rhs22_
                        d_218_v61_: C5
                        nw23_ = C5()
                        nw23_.ctor__((d_136_v0_)[default__.safeIndex(191, (d_136_v0_).length(0))])
                        d_218_v61_ = nw23_
                        d_218_v61_ = (d_218_v61_ if False else d_218_v61_)
                        d_138_v1_ = 17
                    d_138_v1_ = (d_138_v1_) - (d_138_v1_)
                    pass
            pass
        if True:
            d_219_v62_: _dafny.Array
            def lambda7_(d_220_v1_):
                def lambda8_(d_221_i9_):
                    return default__.safeDivisionInt(d_221_i9_, d_220_v1_)

                return lambda8_

            init4_ = lambda7_(d_138_v1_)
            nw24_ = _dafny.Array(None, 15)
            for i0_4_ in range(nw24_.length(0)):
                nw24_[i0_4_] = init4_(i0_4_)
            d_219_v62_ = nw24_
            d_222_v63_: _dafny.Array
            nw25_ = _dafny.Array(None, 12)
            nw25_[int(0)] = d_140_v3_
            nw25_[int(1)] = d_140_v3_
            nw25_[int(2)] = d_140_v3_
            nw25_[int(3)] = d_140_v3_
            nw25_[int(4)] = default__.fm2((d_148_v9_).cardinality, d_138_v1_, d_151_globalState_)
            nw25_[int(5)] = d_140_v3_
            nw25_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ik"))
            nw25_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgljbt"))
            nw25_[int(8)] = d_140_v3_
            nw25_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aomrfrrje"))
            nw25_[int(10)] = d_140_v3_
            nw25_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oroy"))
            d_222_v63_ = nw25_
            d_223_v64_: C12
            nw26_ = C12()
            nw26_.ctor__(d_138_v1_, d_147_v8_, d_138_v1_, d_222_v63_)
            d_223_v64_ = nw26_
            d_224_v65_: C1
            nw27_ = C1()
            nw27_.ctor__(-356, d_222_v63_)
            d_224_v65_ = nw27_
            d_225_v66_: _dafny.Map
            d_225_v66_ = _dafny.Map({d_223_v64_: d_224_v65_})
            index25_ = default__.safeIndex(12, (d_219_v62_).length(0))
            (d_219_v62_)[index25_] = len(d_225_v66_)
            index26_ = default__.safeIndex(12, (d_219_v62_).length(0))
            (d_219_v62_)[index26_] = ((d_223_v64_).f27) * ((d_223_v64_).f27)
            d_138_v1_ = ((d_148_v9_)[d_147_v8_] if (d_147_v8_) in (d_148_v9_) else d_138_v1_)
            index27_ = default__.safeIndex(12, (d_219_v62_).length(0))
            (d_219_v62_)[index27_] = 308
            d_226_v67_: _dafny.Set
            d_226_v67_ = _dafny.Set({(d_219_v62_)[default__.safeIndex(12, (d_219_v62_).length(0))]})
            if ((d_226_v67_) | (d_226_v67_)).ispropersubset((d_226_v67_) | (d_226_v67_)):
                index28_ = default__.safeIndex(288, (d_136_v0_).length(0))
                (d_136_v0_)[index28_] = d_223_v64_.f28
                index29_ = default__.safeIndex(288, (d_136_v0_).length(0))
                index30_ = default__.safeIndex(12, (d_219_v62_).length(0))
                rhs23_ = (d_139_v2_).set(default__.safeIndex(default__.safeModuloInt((d_223_v64_).f27, (d_219_v62_)[default__.safeIndex(12, (d_219_v62_).length(0))]), len(d_139_v2_)), (d_219_v62_)[default__.safeIndex(12, (d_219_v62_).length(0))])
                rhs24_ = False
                rhs25_ = d_223_v64_.f28
                rhs26_ = d_147_v8_
                rhs27_ = 589
                lhs17_ = d_223_v64_
                lhs18_ = d_136_v0_
                lhs19_ = default__.safeIndex(288, (d_136_v0_).length(0))
                lhs20_ = d_151_globalState_
                lhs21_ = d_219_v62_
                lhs22_ = default__.safeIndex(12, (d_219_v62_).length(0))
                d_139_v2_ = rhs23_
                lhs17_.f28 = rhs24_
                lhs18_[lhs19_] = rhs25_
                lhs20_.f12 = rhs26_
                lhs21_[lhs22_] = rhs27_
                d_227_v68_: _dafny.Array
                nw28_ = _dafny.Array(None, 5)
                nw28_[int(0)] = False
                nw28_[int(1)] = d_223_v64_.f28
                nw28_[int(2)] = d_223_v64_.f28
                nw28_[int(3)] = d_147_v8_
                nw28_[int(4)] = d_223_v64_.f28
                d_227_v68_ = nw28_
                d_228_v69_: D19
                d_228_v69_ = D19_DC47(d_227_v68_, d_223_v64_.f28)
                d_229_v70_: _dafny.Array
                nw29_ = _dafny.Array(None, 11)
                nw29_[int(0)] = d_227_v68_
                nw29_[int(1)] = d_227_v68_
                nw29_[int(2)] = (d_228_v69_).cf69
                nw29_[int(3)] = d_227_v68_
                nw29_[int(4)] = d_227_v68_
                nw29_[int(5)] = d_227_v68_
                nw29_[int(6)] = d_227_v68_
                nw29_[int(7)] = d_227_v68_
                nw29_[int(8)] = d_227_v68_
                nw29_[int(9)] = d_227_v68_
                nw29_[int(10)] = d_227_v68_
                d_229_v70_ = nw29_
                d_230_v71_: int
                d_231_v72_: int
                out6_: int
                out7_: int
                out6_, out7_ = (d_224_v65_).m2(d_229_v70_, d_147_v8_, d_223_v64_.f28, d_219_v62_, d_151_globalState_)
                d_230_v71_ = out6_
                d_231_v72_ = out7_
                d_232_v73_: D1
                d_232_v73_ = D1_DC5()
                d_233_v74_: D8
                d_233_v74_ = D8_DC24(d_232_v73_)
                d_234_v75_: _dafny.Map
                d_234_v75_ = _dafny.Map({(d_136_v0_)[default__.safeIndex(288, (d_136_v0_).length(0))]: d_233_v74_})
                d_234_v75_ = d_234_v75_
                d_230_v71_ = (d_139_v2_)[default__.safeIndex(((d_157_v17_).cardinality) + (d_138_v1_), len(d_139_v2_))]
                d_138_v1_ = (d_223_v64_).f27
            elif True:
                d_235_v76_: _dafny.Seq
                d_235_v76_ = _dafny.SeqWithoutIsStrInference([(d_219_v62_)[default__.safeIndex(12, (d_219_v62_).length(0))], len(d_140_v3_)])
                d_236_v77_: _dafny.Map
                d_236_v77_ = _dafny.Map({701: (d_219_v62_)[default__.safeIndex(12, (d_219_v62_).length(0))]})
                d_237_v78_: bool
                d_238_v79_: bool
                d_239_v80_: _dafny.Seq
                out8_: bool
                out9_: bool
                out10_: _dafny.Seq
                out8_, out9_, out10_ = (d_223_v64_).m5(d_219_v62_, d_223_v64_.f28, (d_139_v2_).set(default__.safeIndex(len(d_236_v77_), len(d_139_v2_)), (d_223_v64_).f27), d_147_v8_, d_151_globalState_)
                d_237_v78_ = out8_
                d_238_v79_ = out9_
                d_239_v80_ = out10_
                index31_ = default__.safeIndex(373, (d_136_v0_).length(0))
                (d_136_v0_)[index31_] = d_223_v64_.f28
                d_240_v81_: _dafny.Set
                d_240_v81_ = _dafny.Set({d_236_v77_, (d_236_v77_).set(109, d_138_v1_)})
                d_241_v82_: D21
                d_241_v82_ = D21_DC53(d_240_v81_, True, d_219_v62_)
                index32_ = default__.safeIndex(373, (d_136_v0_).length(0))
                (d_136_v0_)[index32_] = (d_241_v82_).cf83
                d_242_v83_: D1
                d_242_v83_ = D1_DC4(default__.fm1(d_223_v64_.f28, d_151_globalState_))
                (d_223_v64_).m4(d_242_v83_, d_151_globalState_)
                d_243_v85_: _dafny.Seq
                def iife27_():
                    coll23_ = _dafny.Set()
                    compr_23_: int
                    for compr_23_ in (d_139_v2_).Elements:
                        d_244_v84_: int = compr_23_
                        if (d_244_v84_) in (d_139_v2_):
                            coll23_ = coll23_.union(_dafny.Set([(d_244_v84_) - (len(_dafny.SeqWithoutIsStrInference([True])))]))
                    return _dafny.Set(coll23_)
                d_243_v85_ = _dafny.SeqWithoutIsStrInference([iife27_()
                , d_226_v67_])
                d_245_v86_: D18
                d_245_v86_ = D18_DC44(len(d_236_v77_), d_238_v79_, d_140_v3_, (d_136_v0_)[default__.safeIndex(373, (d_136_v0_).length(0))], d_243_v85_)
                index33_ = default__.safeIndex(12, (d_219_v62_).length(0))
                (d_219_v62_)[index33_] = (d_245_v86_).cf63
                d_246_v87_: bool
                d_247_v88_: bool
                d_248_v89_: _dafny.Seq
                out11_: bool
                out12_: bool
                out13_: _dafny.Seq
                out11_, out12_, out13_ = (d_223_v64_).m5(d_219_v62_, ((d_223_v64_).f27) > (d_138_v1_), d_139_v2_, d_147_v8_, d_151_globalState_)
                d_246_v87_ = out11_
                d_247_v88_ = out12_
                d_248_v89_ = out13_
            d_249_v90_: _dafny.Array
            def lambda9_(d_250_v64_):
                def lambda10_(d_251_i10_):
                    return _dafny.SeqWithoutIsStrInference([not(d_250_v64_.f28)])

                return lambda10_

            init5_ = lambda9_(d_223_v64_)
            nw30_ = _dafny.Array(None, 23)
            for i0_5_ in range(nw30_.length(0)):
                nw30_[i0_5_] = init5_(i0_5_)
            d_249_v90_ = nw30_
            d_249_v90_ = d_249_v90_
        elif True:
            d_252_v91_: _dafny.Array
            d_253_v92_: _dafny.Array
            d_254_v93_: bool
            out14_: _dafny.Array
            out15_: _dafny.Array
            out16_: bool
            out14_, out15_, out16_ = default__.m0(d_151_globalState_)
            d_252_v91_ = out14_
            d_253_v92_ = out15_
            d_254_v93_ = out16_
            d_255_v94_: _dafny.Seq
            d_255_v94_ = _dafny.SeqWithoutIsStrInference([d_150_v11_, d_150_v11_])
            if (d_254_v93_ if (len((d_255_v94_)[default__.safeIndex(d_138_v1_, len(d_255_v94_))])) != (d_138_v1_) else d_254_v93_):
                index34_ = default__.safeIndex(79, (d_252_v91_).length(0))
                (d_252_v91_)[index34_] = default__.fm0(d_138_v1_, d_151_globalState_)
                index35_ = default__.safeIndex(79, (d_252_v91_).length(0))
                (d_252_v91_)[index35_] = (d_147_v8_ if d_254_v93_ else d_147_v8_)
                d_256_v95_: _dafny.Array
                nw31_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                d_256_v95_ = nw31_
                d_257_v96_: C6
                nw32_ = C6()
                nw32_.ctor__(d_138_v1_, d_256_v95_)
                d_257_v96_ = nw32_
                rhs28_ = d_136_v0_
                rhs29_ = (d_139_v2_)[default__.safeIndex(default__.fm31(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_258_i11_ in range(default__.abs(563))]), d_151_globalState_), len(d_139_v2_))]
                d_136_v0_ = rhs28_
                d_138_v1_ = rhs29_
                d_259_v97_: _dafny.Map
                d_259_v97_ = _dafny.Map({d_138_v1_: d_142_v5_})
                d_260_v98_: D0
                d_260_v98_ = D0_DC0(default__.fm0(d_138_v1_, d_151_globalState_))
                d_261_v99_: _dafny.Map
                d_261_v99_ = _dafny.Map({d_260_v98_: (d_252_v91_)[default__.safeIndex(79, (d_252_v91_).length(0))]})
                (d_151_globalState_).f23 = ((d_259_v97_)[default__.fm21(d_261_v99_, d_138_v1_, d_147_v8_, (d_140_v3_).set(default__.safeIndex(len(d_140_v3_), len(d_140_v3_)), d_142_v5_), d_151_globalState_)] if (default__.fm21(d_261_v99_, d_138_v1_, d_147_v8_, (d_140_v3_).set(default__.safeIndex(len(d_140_v3_), len(d_140_v3_)), d_142_v5_), d_151_globalState_)) in (d_259_v97_) else d_142_v5_)
                (d_151_globalState_).f23 = d_142_v5_
            elif True:
                d_262_v100_: _dafny.Array
                nw33_ = _dafny.Array(None, 7)
                d_262_v100_ = nw33_
                rhs30_ = -166
                rhs31_ = len(d_140_v3_)
                rhs32_ = d_262_v100_
                d_138_v1_ = rhs30_
                d_138_v1_ = rhs31_
                d_262_v100_ = rhs32_
                d_263_v101_: C13
                nw34_ = C13()
                nw34_.ctor__(not(d_147_v8_))
                d_263_v101_ = nw34_
                d_264_v102_: _dafny.Seq
                d_264_v102_ = _dafny.SeqWithoutIsStrInference([d_263_v101_])
                d_265_v103_: _dafny.Map
                d_265_v103_ = _dafny.Map({d_253_v92_: (d_264_v102_)[default__.safeIndex(d_138_v1_, len(d_264_v102_))]})
                d_265_v103_ = (d_265_v103_).set(d_253_v92_, d_263_v101_)
                rhs33_ = (((d_140_v3_) + (d_140_v3_)) + ((d_140_v3_).set(default__.safeIndex(d_138_v1_, len(d_140_v3_)), default__.fm33(-47, d_254_v93_, d_138_v1_, d_151_globalState_)))).set(default__.safeIndex(d_138_v1_, len(((d_140_v3_) + (d_140_v3_)) + ((d_140_v3_).set(default__.safeIndex(d_138_v1_, len(d_140_v3_)), default__.fm33(-47, d_254_v93_, d_138_v1_, d_151_globalState_))))), d_142_v5_)
                rhs34_ = d_138_v1_
                lhs23_ = d_151_globalState_
                lhs23_.f10 = rhs33_
                d_138_v1_ = rhs34_
                d_266_v104_: _dafny.Array
                nw35_ = _dafny.Array(None, 12)
                d_266_v104_ = nw35_
                d_267_v105_: C11
                nw36_ = C11()
                nw36_.ctor__()
                d_267_v105_ = nw36_
                index36_ = default__.safeIndex(218, (d_266_v104_).length(0))
                (d_266_v104_)[index36_] = d_267_v105_
                index37_ = default__.safeIndex(218, (d_266_v104_).length(0))
                rhs35_ = d_267_v105_
                rhs36_ = d_136_v0_
                rhs37_ = d_138_v1_
                lhs24_ = d_266_v104_
                lhs25_ = default__.safeIndex(218, (d_266_v104_).length(0))
                lhs24_[lhs25_] = rhs35_
                d_136_v0_ = rhs36_
                d_138_v1_ = rhs37_
                d_268_v106_: C5
                nw37_ = C5()
                nw37_.ctor__(d_254_v93_)
                d_268_v106_ = nw37_
            d_269_v107_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.CodePoint('D'), 8)
            d_269_v107_ = nw38_
            index38_ = default__.safeIndex(885, (d_269_v107_).length(0))
            (d_269_v107_)[index38_] = d_142_v5_
            index39_ = default__.safeIndex(885, (d_269_v107_).length(0))
            (d_269_v107_)[index39_] = default__.fm33(d_138_v1_, d_254_v93_, d_138_v1_, d_151_globalState_)
            (d_151_globalState_).f12 = default__.fm0((0) - (d_138_v1_), d_151_globalState_)
            if d_254_v93_:
                (d_151_globalState_).f23 = (d_269_v107_)[default__.safeIndex(885, (d_269_v107_).length(0))]
                d_270_v108_: _dafny.Map
                d_270_v108_ = _dafny.Map({d_138_v1_: d_159_v18_})
                d_271_v109_: _dafny.Array
                nw39_ = _dafny.Array(None, 17)
                nw39_[int(0)] = d_159_v18_
                nw39_[int(1)] = (((d_270_v108_)[d_138_v1_] if (d_138_v1_) in (d_270_v108_) else d_159_v18_)) + (d_159_v18_)
                nw39_[int(2)] = (d_159_v18_) + (d_159_v18_)
                nw39_[int(3)] = d_159_v18_
                nw39_[int(4)] = d_159_v18_
                nw39_[int(5)] = (d_159_v18_) + (_dafny.SeqWithoutIsStrInference([d_147_v8_]))
                nw39_[int(6)] = d_159_v18_
                nw39_[int(7)] = d_159_v18_
                nw39_[int(8)] = default__.fm1(d_147_v8_, d_151_globalState_)
                nw39_[int(9)] = d_159_v18_
                nw39_[int(10)] = default__.fm1(d_254_v93_, d_151_globalState_)
                nw39_[int(11)] = d_159_v18_
                nw39_[int(12)] = d_159_v18_
                nw39_[int(13)] = (default__.fm1(d_147_v8_, d_151_globalState_)) + (d_159_v18_)
                nw39_[int(14)] = d_159_v18_
                nw39_[int(15)] = d_159_v18_
                nw39_[int(16)] = (_dafny.SeqWithoutIsStrInference([(d_159_v18_)[default__.safeIndex(d_138_v1_, len(d_159_v18_))], d_147_v8_])) + (default__.fm1(d_254_v93_, d_151_globalState_))
                d_271_v109_ = nw39_
                index40_ = default__.safeIndex(512, (d_271_v109_).length(0))
                (d_271_v109_)[index40_] = d_159_v18_
                d_272_v110_: _dafny.Map
                d_272_v110_ = _dafny.Map({d_138_v1_: (0) - (d_138_v1_)})
                d_273_v111_: D1
                d_273_v111_ = D1_DC6(_dafny.CodePoint('n'), d_147_v8_, d_140_v3_)
                d_274_v112_: _dafny.Array
                nw40_ = _dafny.Array(None, 14)
                nw40_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sddesiht"))
                nw40_[int(1)] = ((d_140_v3_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_147_v8_])), len(d_140_v3_)), d_142_v5_)).set(default__.safeIndex(d_138_v1_, len((d_140_v3_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_147_v8_])), len(d_140_v3_)), d_142_v5_))), (d_269_v107_)[default__.safeIndex(885, (d_269_v107_).length(0))])
                nw40_[int(2)] = d_140_v3_
                nw40_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtvkkmo"))
                nw40_[int(4)] = default__.fm2(d_138_v1_, len(d_272_v110_), d_151_globalState_)
                nw40_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcb"))
                nw40_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_275_i12_ in range(default__.abs(166))])
                nw40_[int(7)] = d_140_v3_
                nw40_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                nw40_[int(9)] = d_140_v3_
                nw40_[int(10)] = d_140_v3_
                nw40_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybch"))
                nw40_[int(12)] = (d_140_v3_).set(default__.safeIndex(d_138_v1_, len(d_140_v3_)), (d_269_v107_)[default__.safeIndex(885, (d_269_v107_).length(0))])
                nw40_[int(13)] = (d_273_v111_).cf9
                d_274_v112_ = nw40_
                d_276_v113_: _dafny.Array
                nw41_ = _dafny.Array(None, 1)
                nw41_[int(0)] = d_274_v112_
                d_276_v113_ = nw41_
                index41_ = default__.safeIndex(429, (d_276_v113_).length(0))
                (d_276_v113_)[index41_] = d_274_v112_
                index42_ = default__.safeIndex(512, (d_271_v109_).length(0))
                index43_ = default__.safeIndex(429, (d_276_v113_).length(0))
                rhs38_ = d_159_v18_
                rhs39_ = d_274_v112_
                lhs26_ = d_271_v109_
                lhs27_ = default__.safeIndex(512, (d_271_v109_).length(0))
                lhs28_ = d_276_v113_
                lhs29_ = default__.safeIndex(429, (d_276_v113_).length(0))
                lhs26_[lhs27_] = rhs38_
                lhs28_[lhs29_] = rhs39_
                rhs40_ = _dafny.CodePoint('t')
                rhs41_ = d_138_v1_
                d_142_v5_ = rhs40_
                d_138_v1_ = rhs41_
                (d_151_globalState_).f12 = d_254_v93_
                (d_151_globalState_).f10 = d_140_v3_
            elif True:
                d_138_v1_ = d_138_v1_
                d_138_v1_ = (d_138_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "au"))))
                d_277_v114_: _dafny.Array
                nw42_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_277_v114_ = nw42_
                d_278_v115_: C10
                nw43_ = C10()
                nw43_.ctor__((d_269_v107_)[default__.safeIndex(885, (d_269_v107_).length(0))], d_138_v1_, d_277_v114_)
                d_278_v115_ = nw43_
                d_279_v116_: T1
                nw44_ = C1()
                nw44_.ctor__(d_138_v1_, d_277_v114_)
                d_279_v116_ = nw44_
                d_280_v117_: _dafny.MultiSet
                d_280_v117_ = _dafny.MultiSet([d_279_v116_, d_279_v116_])
                d_281_v118_: _dafny.Set
                d_281_v118_ = _dafny.Set({(d_280_v117_).cardinality})
                d_282_v119_: _dafny.Map
                d_282_v119_ = _dafny.Map({(d_281_v118_).ispropersubset(d_281_v118_): d_142_v5_})
                d_282_v119_ = (d_282_v119_).set(d_254_v93_, d_142_v5_)
                (d_151_globalState_).f12 = d_147_v8_
        d_283_v120_: _dafny.Array
        d_284_v121_: _dafny.Array
        d_285_v122_: bool
        out17_: _dafny.Array
        out18_: _dafny.Array
        out19_: bool
        out17_, out18_, out19_ = default__.m0(d_151_globalState_)
        d_283_v120_ = out17_
        d_284_v121_ = out18_
        d_285_v122_ = out19_
        d_147_v8_ = False
        hi0_ = d_138_v1_
        for d_286_i13_ in range(d_138_v1_, hi0_):
            d_136_v0_ = d_136_v0_
            index44_ = default__.safeIndex(327, (d_136_v0_).length(0))
            (d_136_v0_)[index44_] = d_285_v122_
            d_287_v123_: _dafny.Map
            d_287_v123_ = _dafny.Map({d_138_v1_: d_138_v1_})
            d_288_v124_: D14
            d_288_v124_ = D14_DC37(d_287_v123_)
            d_289_v125_: D22
            d_289_v125_ = D22_DC56(_dafny.CodePoint('m'), d_147_v8_)
            d_290_v126_: _dafny.Map
            d_290_v126_ = _dafny.Map({d_289_v125_: d_285_v122_})
            d_291_v129_: _dafny.Seq
            d_291_v129_ = _dafny.SeqWithoutIsStrInference([d_289_v125_])
            d_292_v130_: _dafny.Seq
            def iife28_():
                def iife30_():
                    coll26_ = _dafny.Map()
                    compr_26_: D22
                    for compr_26_ in (d_291_v129_).Elements:
                        d_293_v128_: D22 = compr_26_
                        if (d_293_v128_) in (d_291_v129_):
                            coll26_[d_293_v128_] = d_285_v122_
                    return _dafny.Map(coll26_)
                coll24_ = _dafny.Map()
                def iife29_():
                    coll25_ = _dafny.Map()
                    compr_25_: D22
                    for compr_25_ in (d_291_v129_).Elements:
                        d_293_v128_: D22 = compr_25_
                        if (d_293_v128_) in (d_291_v129_):
                            coll25_[d_293_v128_] = d_285_v122_
                    return _dafny.Map(coll25_)
                compr_24_: D22
                for compr_24_ in (iife29_()
                ).keys.Elements:
                    d_294_v127_: D22 = compr_24_
                    if (d_294_v127_) in (iife30_()
                    ):
                        coll24_[d_294_v127_] = d_285_v122_
                return _dafny.Map(coll24_)
            d_292_v130_ = _dafny.SeqWithoutIsStrInference([d_290_v126_, iife28_()
            ])
            d_295_v131_: _dafny.Map
            d_295_v131_ = _dafny.Map({d_288_v124_: d_292_v130_})
            d_296_v133_: _dafny.MultiSet
            def iife31_():
                coll27_ = _dafny.Map()
                compr_27_: D22
                for compr_27_ in (default__.fm74(d_151_globalState_)).Elements:
                    d_297_v132_: D22 = compr_27_
                    if (d_297_v132_) in (default__.fm74(d_151_globalState_)):
                        coll27_[d_297_v132_] = d_285_v122_
                return _dafny.Map(coll27_)
            d_296_v133_ = _dafny.MultiSet([d_290_v126_, d_290_v126_, iife31_()
            ])
            index45_ = default__.safeIndex(327, (d_136_v0_).length(0))
            (d_136_v0_)[index45_] = (d_296_v133_).ispropersubset(_dafny.MultiSet(((d_295_v131_)[d_288_v124_] if (d_288_v124_) in (d_295_v131_) else d_292_v130_)))
            d_298_v134_: _dafny.Array
            nw45_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
            d_298_v134_ = nw45_
            index46_ = default__.safeIndex(379, (d_298_v134_).length(0))
            (d_298_v134_)[index46_] = d_140_v3_
            index47_ = default__.safeIndex(379, (d_298_v134_).length(0))
            (d_298_v134_)[index47_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_299_i14_ in range(default__.abs(213))])) + ((d_140_v3_ if d_147_v8_ else d_140_v3_))
            d_284_v121_ = d_284_v121_
        rhs42_ = (d_140_v3_)[default__.safeIndex(605, len(d_140_v3_))]
        rhs43_ = (d_141_v4_)[default__.safeIndex((d_138_v1_) * (default__.fm18(d_151_globalState_)), len(d_141_v4_))]
        rhs44_ = d_285_v122_
        lhs30_ = d_151_globalState_
        lhs31_ = d_151_globalState_
        lhs30_.f23 = rhs42_
        lhs31_.f10 = rhs43_
        d_285_v122_ = rhs44_
        d_285_v122_ = not(not(d_147_v8_))
        d_188_v43_ = (d_188_v43_).set(d_138_v1_, d_147_v8_)
        d_300_v135_: D19
        d_300_v135_ = D19_DC47(d_136_v0_, d_147_v8_)
        pat_let_tv1_ = d_283_v120_
        d_301_v136_: _dafny.MultiSet
        def iife32_(_pat_let2_0):
            def iife33_(d_302_dt__update__tmp_h3_):
                def iife34_(_pat_let3_0):
                    def iife35_(d_303_dt__update_hcf69_h1_):
                        return D19_DC47(d_303_dt__update_hcf69_h1_, (d_302_dt__update__tmp_h3_).cf70)
                    return iife35_(_pat_let3_0)
                return iife34_(pat_let_tv1_)
            return iife33_(_pat_let2_0)
        d_301_v136_ = _dafny.MultiSet([iife32_(D19_DC47(d_136_v0_, d_147_v8_)), d_300_v135_, d_300_v135_, d_300_v135_])
        d_304_v137_: _dafny.MultiSet
        d_304_v137_ = _dafny.MultiSet([d_142_v5_])
        d_305_v138_: _dafny.Set
        d_305_v138_ = _dafny.Set({(d_304_v137_).cardinality})
        d_306_v139_: _dafny.Seq
        d_306_v139_ = (default__.fm19(not(d_285_v122_), d_151_globalState_)).set(default__.safeIndex(d_138_v1_, len(default__.fm19(not(d_285_v122_), d_151_globalState_))), len(d_305_v138_))
        pat_let_tv2_ = d_138_v1_
        def lambda11_(source9_):
            d_307___mcc_h4_ = source9_
            d_308_cf11_ = d_307___mcc_h4_
            return pat_let_tv2_

        rhs45_ = d_157_v17_
        rhs46_ = lambda11_(d_306_v139_)
        rhs47_ = ((d_301_v136_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_300_v135_, d_300_v135_]))) if (d_285_v122_) or (d_147_v8_) else d_301_v136_)
        d_157_v17_ = rhs45_
        d_138_v1_ = rhs46_
        d_301_v136_ = rhs47_
        d_142_v5_ = d_142_v5_
        d_309_v140_: C11
        nw46_ = C11()
        nw46_.ctor__()
        d_309_v140_ = nw46_
        d_147_v8_ = d_147_v8_
        _dafny.print(_dafny.string_of((d_136_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v0_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v2_) == (_dafny.SeqWithoutIsStrInference([454]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_140_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_141_v4_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvkwml"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v6_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvkwml"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[0]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[1]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[2]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[3]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[4]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[5]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[6]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[7]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[8]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[9]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[10]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[11]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[12]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[13]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[14]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[15]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[16]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[17]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[18]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[19]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[20]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[21]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_144_v7_)[22]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_147_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v9_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v10_) == (_dafny.Map({249: _dafny.MultiSet([False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v11_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f4)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_.f6) == (_dafny.SeqWithoutIsStrInference([454]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_151_globalState_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_151_globalState_.f10).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_151_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_.f16) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvkwml")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvkwml"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[0]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[1]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[2]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[3]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[4]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[5]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[6]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[7]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[8]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[9]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[10]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[11]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[12]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[13]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[14]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[15]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[16]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[17]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[18]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[19]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[20]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[21]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_globalState_).f17)[22]) == (_dafny.Map({454: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f18) == (_dafny.Map({249: _dafny.MultiSet([False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_globalState_).f20) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_151_globalState_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v12_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v13_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v14_) == (_dafny.Map({2: _dafny.Map({True: 1})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v16_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v16_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v17_) == (_dafny.MultiSet([454, 454]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v18_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v19_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_188_v43_) == (_dafny.Map({1: False, 2: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_189_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v120_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v121_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_285_v122_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_300_v135_).cf69)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v135_).cf70))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_301_v136_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v137_) == (_dafny.MultiSet([_dafny.CodePoint('k')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v138_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_306_v139_)) == (_dafny.SeqWithoutIsStrInference([1, 166, 1, 1, 4]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
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
        return lambda: D1_DC5()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)

class D1_DC5(D1, NamedTuple('DC5', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {self.cf9.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

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
        return lambda: D3_DC10(False, _dafny.Array(None, 0), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
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
        return lambda: D4_DC12()
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
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)

class D4_DC12(D4, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC16(D5, NamedTuple('DC16', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC19(D0.default()(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC19(D6, NamedTuple('DC19', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC22(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC22(D7, NamedTuple('DC22', [('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(D1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC24(D8, NamedTuple('DC24', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC27(D9, NamedTuple('DC27', [('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC30(D11, NamedTuple('DC30', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC32(D12, NamedTuple('DC32', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC36(_dafny.Map({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC36(D13, NamedTuple('DC36', [('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC38(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)

class D14_DC38(D14, NamedTuple('DC38', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC40()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)

class D15_DC40(D15, NamedTuple('DC40', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC39(D15, NamedTuple('DC39', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)

class D16_DC41(D16, NamedTuple('DC41', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)

class D17_DC42(D17, NamedTuple('DC42', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC44(int(0), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D18_DC45)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)

class D18_DC44(D18, NamedTuple('DC44', [('cf63', Any), ('cf64', Any), ('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)}, {self.cf65.VerbatimString(True)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC45(D18, NamedTuple('DC45', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC45'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC45)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC47(_dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)

class D19_DC47(D19, NamedTuple('DC47', [('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC46(D19, NamedTuple('DC46', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC50(D5.default()(), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)

class D20_DC50(D20, NamedTuple('DC50', [('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {self.cf75.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC49(D20, NamedTuple('DC49', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC52(int(0), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)

class D21_DC52(D21, NamedTuple('DC52', [('cf77', Any), ('cf78', Any), ('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {self.cf79.VerbatimString(True)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC53(D21, NamedTuple('DC53', [('cf82', Any), ('cf83', Any), ('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC51(D21, NamedTuple('DC51', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC56(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D22_DC55)

class D22_DC56(D22, NamedTuple('DC56', [('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC57(D22, NamedTuple('DC57', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC55(D22, NamedTuple('DC55', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC55({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC55) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC59(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D23_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D23_DC58)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D23_DC60)

class D23_DC59(D23, NamedTuple('DC59', [('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC58(D23, NamedTuple('DC58', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC58({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC58) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC60(D23, NamedTuple('DC60', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC60({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC60) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC62(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D24_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D24_DC63)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D24_DC64)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D24_DC61)

class D24_DC62(D24, NamedTuple('DC62', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC62({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC62) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC63(D24, NamedTuple('DC63', [('cf96', Any), ('cf97', Any), ('cf98', Any), ('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC63({_dafny.string_of(self.cf96)}, {self.cf97.VerbatimString(True)}, {_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC63) and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98 and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC64(D24, NamedTuple('DC64', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC64({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC64) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC61(D24, NamedTuple('DC61', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC61({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC61) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D25_DC65)

class D25_DC65(D25, NamedTuple('DC65', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC65({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC65) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m2(self, p0, p1, p2, p3, globalState):
        pass

    def m3(self, p0, p1, globalState):
        pass


class T1:
    pass
    def m4(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f6: _dafny.Seq = _dafny.Seq({})
        self.f10: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f12: bool = False
        self.f16: _dafny.MultiSet = _dafny.MultiSet({})
        self.f23: str = _dafny.CodePoint('D')
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f2: bool = False
        self._f3: int = int(0)
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f5: int = int(0)
        self._f7: int = int(0)
        self._f8: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f11: bool = False
        self._f13: bool = False
        self._f14: bool = False
        self._f15: bool = False
        self._f17: _dafny.Array = _dafny.Array(None, 0)
        self._f18: _dafny.Map = _dafny.Map({})
        self._f19: bool = False
        self._f20: _dafny.Set = _dafny.Set({})
        self._f21: int = int(0)
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self).f23 = f23

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
    def f5(self):
        return self._f5
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
    def f11(self):
        return self._f11
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
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
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm26(self, p0, p1, globalState):
        source10_ = D4_DC14(D4_DC12())
        if source10_.is_DC12:
            return len(_dafny.SeqWithoutIsStrInference([not(not(False))]))
        elif source10_.is_DC13:
            d_310___mcc_h0_ = source10_.cf19
            d_311___mcc_h1_ = source10_.cf20
            d_312___mcc_h2_ = source10_.cf21
            d_313___mcc_h3_ = source10_.cf22
            d_314_cf22_ = d_313___mcc_h3_
            d_315_cf21_ = d_312___mcc_h2_
            d_316_cf20_ = d_311___mcc_h1_
            d_317_cf19_ = d_310___mcc_h0_
            def iife36_():
                def iife38_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(345, 518):
                        d_320_v0_: int = compr_30_
                        if ((345) <= (d_320_v0_)) and ((d_320_v0_) < (518)):
                            coll30_ = coll30_.union(_dafny.Set([(d_320_v0_) + (d_317_cf19_)]))
                    return _dafny.Set(coll30_)
                coll28_ = _dafny.Set()
                def iife37_():
                    coll29_ = _dafny.Set()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(345, 518):
                        d_318_v0_: int = compr_29_
                        if ((345) <= (d_318_v0_)) and ((d_318_v0_) < (518)):
                            coll29_ = coll29_.union(_dafny.Set([(d_318_v0_) + (d_317_cf19_)]))
                    return _dafny.Set(coll29_)
                compr_28_: int
                for compr_28_ in (iife37_()
                ).Elements:
                    d_319_v1_: int = compr_28_
                    if (d_319_v1_) in (iife38_()
                    ):
                        coll28_ = coll28_.union(_dafny.Set([default__.safeDivisionInt(d_319_v1_, 274)]))
                return _dafny.Set(coll28_)
            return len((_dafny.SeqWithoutIsStrInference([_dafny.Set({356}), (D4_DC11(_dafny.Set({d_317_cf19_}))).cf18]) if d_316_cf20_ else _dafny.SeqWithoutIsStrInference([_dafny.Set({d_317_cf19_}), iife36_()
            , _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crfbyfmc"))))}), _dafny.Set({d_317_cf19_})])))
        elif source10_.is_DC11:
            d_321___mcc_h4_ = source10_.cf18
            d_322_cf18_ = d_321___mcc_h4_
            return (987) * (75)
        elif True:
            d_323___mcc_h5_ = source10_.cf23
            d_324_cf23_ = d_323___mcc_h5_
            return -90


class C1(T1, T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f25, f26):
        (self)._f25 = f25
        (self)._f26 = f26

    def m4(self, p0, globalState):
        d_325_v0_: int
        d_325_v0_ = -112
        d_326_v1_: bool
        d_326_v1_ = True
        d_327_v2_: _dafny.Seq
        d_327_v2_ = _dafny.SeqWithoutIsStrInference([d_326_v1_])
        d_328_v3_: str
        d_328_v3_ = _dafny.CodePoint('p')
        d_329_v4_: _dafny.Seq
        d_329_v4_ = _dafny.SeqWithoutIsStrInference([d_328_v3_, d_328_v3_])
        d_325_v0_ = default__.safeDivisionInt(default__.safeDivisionInt((0) - (len((d_327_v2_).set(default__.safeIndex((self).f25, len(d_327_v2_)), d_326_v1_))), (self).f25), len((_dafny.SeqWithoutIsStrInference([d_328_v3_])) + (d_329_v4_)))
        d_326_v1_ = (not(d_326_v1_) if d_326_v1_ else d_326_v1_)
        d_330_v5_: _dafny.Map
        d_330_v5_ = _dafny.Map({d_326_v1_: _dafny.CodePoint('x')})
        d_330_v5_ = (d_330_v5_).set(d_326_v1_, d_328_v3_)
        d_331_v6_: _dafny.Map
        d_331_v6_ = _dafny.Map({default__.safeDivisionInt(d_325_v0_, len(d_329_v4_)): 495})
        d_331_v6_ = (d_331_v6_).set(((0) - (d_325_v0_)) + (d_325_v0_), d_325_v0_)
        d_332_v7_: _dafny.Array
        def lambda12_(d_333_i0_):
            return False

        init6_ = lambda12_
        nw47_ = _dafny.Array(None, 17)
        for i0_6_ in range(nw47_.length(0)):
            nw47_[i0_6_] = init6_(i0_6_)
        d_332_v7_ = nw47_
        d_334_v8_: D3
        d_334_v8_ = D3_DC10(d_326_v1_, d_332_v7_, d_326_v1_, (self).f25, d_325_v0_)
        d_335_v9_: _dafny.Array
        nw48_ = _dafny.Array(None, 20)
        nw48_[int(0)] = d_332_v7_
        nw48_[int(1)] = d_332_v7_
        nw48_[int(2)] = d_332_v7_
        nw48_[int(3)] = d_332_v7_
        nw48_[int(4)] = d_332_v7_
        nw48_[int(5)] = d_332_v7_
        nw48_[int(6)] = d_332_v7_
        nw48_[int(7)] = (d_332_v7_ if d_326_v1_ else d_332_v7_)
        nw48_[int(8)] = d_332_v7_
        nw48_[int(9)] = (d_334_v8_).cf14
        nw48_[int(10)] = d_332_v7_
        nw48_[int(11)] = d_332_v7_
        nw48_[int(12)] = d_332_v7_
        nw48_[int(13)] = d_332_v7_
        nw48_[int(14)] = d_332_v7_
        nw48_[int(15)] = d_332_v7_
        nw48_[int(16)] = d_332_v7_
        nw48_[int(17)] = d_332_v7_
        nw48_[int(18)] = (d_334_v8_).cf14
        nw48_[int(19)] = d_332_v7_
        d_335_v9_ = nw48_
        index48_ = default__.safeIndex(724, (d_335_v9_).length(0))
        (d_335_v9_)[index48_] = d_332_v7_
        d_336_v10_: D1
        d_336_v10_ = D1_DC5()
        pat_let_tv3_ = d_327_v2_
        pat_let_tv4_ = d_325_v0_
        pat_let_tv5_ = d_327_v2_
        pat_let_tv6_ = d_326_v1_
        pat_let_tv7_ = d_326_v1_
        index49_ = default__.safeIndex(724, (d_335_v9_).length(0))
        def lambda13_(source11_):
            if source11_.is_DC5:
                return (pat_let_tv3_)[default__.safeIndex(pat_let_tv4_, len(pat_let_tv5_))]
            elif source11_.is_DC6:
                d_337___mcc_h0_ = source11_.cf7
                d_338___mcc_h1_ = source11_.cf8
                d_339___mcc_h2_ = source11_.cf9
                d_340_cf9_ = d_339___mcc_h2_
                d_341_cf8_ = d_338___mcc_h1_
                d_342_cf7_ = d_337___mcc_h0_
                return d_341_cf8_
            elif source11_.is_DC4:
                d_343___mcc_h3_ = source11_.cf6
                d_344_cf6_ = d_343___mcc_h3_
                return pat_let_tv6_
            elif True:
                d_345___mcc_h4_ = source11_.cf10
                d_346_cf10_ = d_345___mcc_h4_
                return pat_let_tv7_

        rhs48_ = d_332_v7_
        rhs49_ = lambda13_(d_336_v10_)
        lhs32_ = d_335_v9_
        lhs33_ = default__.safeIndex(724, (d_335_v9_).length(0))
        lhs34_ = globalState
        lhs32_[lhs33_] = rhs48_
        lhs34_.f12 = rhs49_
        d_347_v11_: _dafny.Array
        nw49_ = _dafny.Array(_dafny.Seq({}), 18)
        d_347_v11_ = nw49_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_347_v11_).length(0)):
            d_348_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_348_i1_)) and ((d_348_i1_) < ((d_347_v11_).length(0)))):
                (d_347_v11_)[(d_348_i1_)] = ((_dafny.SeqWithoutIsStrInference([(self).f25, 683]) if d_326_v1_ else _dafny.SeqWithoutIsStrInference([len(d_329_v4_)]))) + (_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25]))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_349_v0_: _dafny.MultiSet
        d_349_v0_ = _dafny.MultiSet([p2])
        d_350_v1_: _dafny.Seq
        d_350_v1_ = _dafny.SeqWithoutIsStrInference([p1, p2, ((self).f25) <= ((self).f25)])
        d_349_v0_ = _dafny.MultiSet(d_350_v1_)
        d_351_v2_: _dafny.Seq
        d_351_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urjv"))
        d_352_v3_: _dafny.Map
        d_352_v3_ = _dafny.Map({len(d_351_v2_): (d_351_v2_) + (d_351_v2_)})
        d_352_v3_ = d_352_v3_
        r1 = ((self).f25) + ((self).f25)
        r0 = (self).f25
        r0 = default__.safeModuloInt((0) - (default__.safeDivisionInt((self).f25, (self).f25)), (self).f25)
        d_350_v1_ = d_350_v1_
        r0 = (0) - (((self).f25) * ((self).f25))
        r1 = default__.fm31(d_351_v2_, globalState)
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        if p1:
            d_353_v0_: D0
            d_353_v0_ = D0_DC1(True, (not(p1)) and (p1))
            d_353_v0_ = d_353_v0_
            d_354_v1_: _dafny.Seq
            d_354_v1_ = _dafny.SeqWithoutIsStrInference([(self).f25, 995, (self).f25, (self).f25, (self).f25])
            d_355_v2_: _dafny.Map
            d_355_v2_ = _dafny.Map({not (p1) or (p1): default__.safeModuloInt((d_354_v1_)[default__.safeIndex((self).f25, len(d_354_v1_))], (self).f25)})
            d_356_v3_: _dafny.MultiSet
            d_356_v3_ = _dafny.MultiSet([default__.fm0((self).f25, globalState), p1, p1])
            d_355_v2_ = (d_355_v2_).set(p1, (d_356_v3_).cardinality)
            d_357_v4_: _dafny.Map
            d_357_v4_ = _dafny.Map({(0) - ((self).f25): (d_354_v1_)[default__.safeIndex((self).f25, len(d_354_v1_))]})
            d_358_v5_: D4
            d_358_v5_ = D4_DC13(len(d_357_v4_), not(p1), p1, p1)
            d_359_v6_: _dafny.Map
            d_359_v6_ = _dafny.Map({p1: d_358_v5_})
            d_359_v6_ = (d_359_v6_).set(p1, d_358_v5_)
            d_360_v7_: _dafny.Array
            nw50_ = _dafny.Array(False, 1)
            d_360_v7_ = nw50_
            d_361_v8_: _dafny.Map
            d_361_v8_ = _dafny.Map({d_360_v7_: p1})
            if ((d_361_v8_)[d_360_v7_] if (d_360_v7_) in (d_361_v8_) else ((self).f25) > ((self).f25)):
                d_362_v9_: _dafny.Seq
                d_362_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvs"))
                d_363_v10_: _dafny.Map
                d_363_v10_ = _dafny.Map({p1: ((self).f25) == (default__.fm31(d_362_v9_, globalState))})
                d_364_v11_: _dafny.Seq
                d_364_v11_ = _dafny.SeqWithoutIsStrInference([p1])
                d_365_v12_: str
                d_365_v12_ = _dafny.CodePoint('s')
                d_366_v13_: _dafny.Map
                d_366_v13_ = _dafny.Map({d_365_v12_: d_356_v3_})
                d_363_v10_ = (d_363_v10_).set(p1, (d_364_v11_)[default__.safeIndex(len(d_366_v13_), len(d_364_v11_))])
                d_367_v14_: D1
                d_367_v14_ = D1_DC4((_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex((self).f25, len(_dafny.SeqWithoutIsStrInference([p1]))), p1))
                d_364_v11_ = (d_367_v14_).cf6
                (globalState).f10 = ((_dafny.SeqWithoutIsStrInference([d_365_v12_ for d_368_i0_ in range(default__.abs(443))])) + ((_dafny.SeqWithoutIsStrInference([d_365_v12_ for d_369_i1_ in range(default__.abs(905))])) + (d_362_v9_))).set(default__.safeIndex((self).f25, len((_dafny.SeqWithoutIsStrInference([d_365_v12_ for d_370_i0_ in range(default__.abs(443))])) + ((_dafny.SeqWithoutIsStrInference([d_365_v12_ for d_371_i1_ in range(default__.abs(905))])) + (d_362_v9_)))), d_365_v12_)
                d_357_v4_ = (d_357_v4_).set(-679, (len(d_364_v11_)) * ((self).f25))
                d_355_v2_ = (d_355_v2_).set(False, (self).f25)
            elif True:
                (globalState).f12 = not((p1) or (p1))
                d_372_v15_: _dafny.Seq
                d_372_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pntpd"))
                d_373_v16_: _dafny.Map
                d_373_v16_ = _dafny.Map({(self).f25: (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_374_i2_ in range(default__.abs(396))])) != (d_372_v15_)})
                d_373_v16_ = (d_373_v16_).set(default__.safeModuloInt((self).f25, (self).f25), ((self).f25) >= ((d_356_v3_).cardinality))
                r0 = ((self).f25) - (((self).f25 if True else (self).f25))
                d_375_v17_: _dafny.MultiSet
                d_375_v17_ = _dafny.MultiSet([(self).f25, (self).f25])
                d_376_v19_: D5
                d_376_v19_ = D5_DC15((d_375_v17_).set((self).f25, default__.abs((self).f25)))
                d_377_v20_: _dafny.Array
                nw51_ = _dafny.Array(None, 29)
                nw51_[int(0)] = d_375_v17_
                nw51_[int(1)] = default__.fm32((self).f25, globalState)
                nw51_[int(2)] = d_375_v17_
                def iife39_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(293, 694):
                        d_378_v18_: int = compr_31_
                        if ((293) <= (d_378_v18_)) and ((d_378_v18_) < (694)):
                            coll31_[(d_378_v18_) - ((self).f25)] = _dafny.CodePoint('i')
                    return _dafny.Map(coll31_)
                nw51_[int(3)] = ((d_375_v17_).set((self).f25, default__.abs(len(iife39_()
                )))) - (_dafny.MultiSet([(self).f25]))
                nw51_[int(4)] = (d_375_v17_) | (d_375_v17_)
                nw51_[int(5)] = d_375_v17_
                nw51_[int(6)] = d_375_v17_
                nw51_[int(7)] = d_375_v17_
                nw51_[int(8)] = _dafny.MultiSet([(self).f25, default__.fm31(d_372_v15_, globalState), (self).f25, (self).f25])
                nw51_[int(9)] = d_375_v17_
                nw51_[int(10)] = d_375_v17_
                nw51_[int(11)] = d_375_v17_
                nw51_[int(12)] = (d_375_v17_) - (d_375_v17_)
                nw51_[int(13)] = d_375_v17_
                nw51_[int(14)] = _dafny.MultiSet(d_354_v1_)
                nw51_[int(15)] = d_375_v17_
                nw51_[int(16)] = d_375_v17_
                nw51_[int(17)] = (d_375_v17_).intersection(_dafny.MultiSet([(self).f25, (self).f25, (0) - ((d_375_v17_).cardinality), (self).f25, (self).f25]))
                nw51_[int(18)] = d_375_v17_
                nw51_[int(19)] = ((_dafny.MultiSet([len(d_372_v15_)])).set(default__.fm31(d_372_v15_, globalState), default__.abs((self).f25)) if default__.fm0((self).f25, globalState) else (d_375_v17_).set((self).f25, default__.abs(147)))
                nw51_[int(20)] = (d_375_v17_).intersection(d_375_v17_)
                nw51_[int(21)] = d_375_v17_
                nw51_[int(22)] = d_375_v17_
                nw51_[int(23)] = default__.fm32(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_379_i3_ in range(default__.abs(701))])), globalState)
                nw51_[int(24)] = (d_375_v17_).intersection((d_376_v19_).cf24)
                nw51_[int(25)] = d_375_v17_
                nw51_[int(26)] = d_375_v17_
                nw51_[int(27)] = d_375_v17_
                nw51_[int(28)] = d_375_v17_
                d_377_v20_ = nw51_
                d_380_v21_: _dafny.Array
                nw52_ = _dafny.Array(int(0), 26)
                d_380_v21_ = nw52_
                index50_ = default__.safeIndex(382, (d_377_v20_).length(0))
                (d_377_v20_)[index50_] = (_dafny.MultiSet([(self).f25])).intersection(default__.fm32(len(_dafny.SeqWithoutIsStrInference([d_380_v21_])), globalState))
                index51_ = default__.safeIndex(382, (d_377_v20_).length(0))
                (d_377_v20_)[index51_] = d_375_v17_
                d_381_v22_: _dafny.Seq
                d_381_v22_ = _dafny.SeqWithoutIsStrInference([p1])
                d_382_v23_: _dafny.Set
                d_382_v23_ = _dafny.Set({(self).f25, default__.fm31(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htgrnhtir")), globalState)})
                d_383_v24_: _dafny.Map
                d_383_v24_ = _dafny.Map({(d_381_v22_)[default__.safeIndex((self).f25, len(d_381_v22_))]: (_dafny.Set({697, (self).f25, (self).f25, (self).f25})).issubset(d_382_v23_)})
                d_384_v25_: D1
                d_384_v25_ = D1_DC6(default__.fm33((self).f25, not(p1), (self).f25, globalState), p1, d_372_v15_)
                d_383_v24_ = (d_383_v24_).set(False, (d_384_v25_).cf8)
            d_385_v26_: _dafny.Seq
            d_385_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqhshe"))
            r0 = default__.fm31((d_385_v26_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_386_i4_ in range(default__.abs(898))])), globalState)
        elif True:
            r0 = (self).f25
            d_387_v27_: _dafny.Array
            def lambda14_(d_388_i5_):
                return default__.safeModuloInt(d_388_i5_, (self).f25)

            init7_ = lambda14_
            nw53_ = _dafny.Array(None, 26)
            for i0_7_ in range(nw53_.length(0)):
                nw53_[i0_7_] = init7_(i0_7_)
            d_387_v27_ = nw53_
            index52_ = default__.safeIndex(69, (d_387_v27_).length(0))
            (d_387_v27_)[index52_] = (self).f25
            d_389_v28_: _dafny.Seq
            d_389_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hy"))
            d_390_v29_: _dafny.Seq
            d_390_v29_ = _dafny.SeqWithoutIsStrInference([len(d_389_v28_)])
            d_391_v30_: _dafny.Seq
            d_391_v30_ = _dafny.SeqWithoutIsStrInference([d_390_v29_])
            d_392_v31_: _dafny.Array
            nw54_ = _dafny.Array(None, 7)
            nw54_[int(0)] = d_391_v30_
            nw54_[int(1)] = d_391_v30_
            nw54_[int(2)] = d_391_v30_
            nw54_[int(3)] = d_391_v30_
            nw54_[int(4)] = (d_391_v30_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f25, (0) - ((self).f25)]) for d_393_i6_ in range(default__.abs(392))]))
            nw54_[int(5)] = d_391_v30_
            nw54_[int(6)] = (d_391_v30_).set(default__.safeIndex((self).f25, len(d_391_v30_)), d_390_v29_)
            d_392_v31_ = nw54_
            index53_ = default__.safeIndex(450, (d_392_v31_).length(0))
            (d_392_v31_)[index53_] = d_391_v30_
            d_394_v32_: _dafny.Array
            nw55_ = _dafny.Array(False, 29)
            d_394_v32_ = nw55_
            index54_ = default__.safeIndex(13, (d_394_v32_).length(0))
            (d_394_v32_)[index54_] = (default__.fm33((self).f25, not(True), (self).f25, globalState)) in (d_389_v28_)
            d_395_v33_: _dafny.Map
            d_395_v33_ = _dafny.Map({p1: d_389_v28_})
            d_396_v34_: _dafny.Array
            def lambda15_(d_397_p1_):
                def lambda16_(d_398_i7_):
                    return D5_DC16(d_397_p1_)

                return lambda16_

            init8_ = lambda15_(p1)
            nw56_ = _dafny.Array(None, 23)
            for i0_8_ in range(nw56_.length(0)):
                nw56_[i0_8_] = init8_(i0_8_)
            d_396_v34_ = nw56_
            d_399_v35_: _dafny.Map
            d_399_v35_ = _dafny.Map({(self).f25: d_396_v34_})
            d_400_v36_: _dafny.MultiSet
            d_400_v36_ = _dafny.MultiSet([(self).f25, len(d_399_v35_)])
            d_401_v37_: D5
            d_401_v37_ = D5_DC15(d_400_v36_)
            d_402_v38_: _dafny.Seq
            d_402_v38_ = _dafny.SeqWithoutIsStrInference([d_401_v37_, d_401_v37_, d_401_v37_])
            d_403_v39_: _dafny.Seq
            d_403_v39_ = _dafny.SeqWithoutIsStrInference([d_402_v38_, d_402_v38_])
            index55_ = default__.safeIndex(69, (d_387_v27_).length(0))
            index56_ = default__.safeIndex(450, (d_392_v31_).length(0))
            index57_ = default__.safeIndex(13, (d_394_v32_).length(0))
            rhs50_ = (len(_dafny.SeqWithoutIsStrInference([(self).f25]))) * ((self).f25)
            rhs51_ = default__.safeDivisionInt(((self).f25) + ((self).f25), (self).f25)
            rhs52_ = default__.safeDivisionInt(len(d_395_v33_), (self).f25)
            rhs53_ = d_391_v30_
            rhs54_ = (d_403_v39_) < (d_403_v39_)
            lhs35_ = d_387_v27_
            lhs36_ = default__.safeIndex(69, (d_387_v27_).length(0))
            lhs37_ = d_392_v31_
            lhs38_ = default__.safeIndex(450, (d_392_v31_).length(0))
            lhs39_ = d_394_v32_
            lhs40_ = default__.safeIndex(13, (d_394_v32_).length(0))
            r0 = rhs50_
            r0 = rhs51_
            lhs35_[lhs36_] = rhs52_
            lhs37_[lhs38_] = rhs53_
            lhs39_[lhs40_] = rhs54_
            index58_ = default__.safeIndex(13, (d_394_v32_).length(0))
            rhs55_ = p1
            rhs56_ = ((d_387_v27_)[default__.safeIndex(69, (d_387_v27_).length(0))]) != ((self).f25)
            lhs41_ = d_394_v32_
            lhs42_ = default__.safeIndex(13, (d_394_v32_).length(0))
            lhs43_ = globalState
            lhs41_[lhs42_] = rhs55_
            lhs43_.f12 = rhs56_
            d_400_v36_ = d_400_v36_
            (globalState).f12 = default__.fm0(((self).f25) * ((d_387_v27_)[default__.safeIndex(69, (d_387_v27_).length(0))]), globalState)
        d_404_v40_: _dafny.Seq
        d_404_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vp"))
        d_405_v41_: _dafny.Set
        d_405_v41_ = _dafny.Set({len(d_404_v40_)})
        d_406_v42_: D4
        d_406_v42_ = D4_DC11(d_405_v41_)
        d_407_v43_: _dafny.Seq
        d_407_v43_ = _dafny.SeqWithoutIsStrInference([d_406_v42_, default__.fm34(p1, p1, p1, (self).f25, globalState)])
        d_408_v44_: _dafny.Map
        d_408_v44_ = _dafny.Map({d_407_v43_: False})
        d_408_v44_ = (d_408_v44_).set(_dafny.SeqWithoutIsStrInference([d_406_v42_, d_406_v42_, d_406_v42_]), p1)
        d_409_v45_: _dafny.Array
        nw57_ = _dafny.Array(False, 23)
        d_409_v45_ = nw57_
        d_409_v45_ = d_409_v45_
        r0 = (self).f25
        hi1_ = (self).f25
        for d_410_i8_ in range((self).f25, hi1_):
            source12_ = d_406_v42_
            if source12_.is_DC12:
                d_411_v46_: _dafny.Map
                d_411_v46_ = _dafny.Map({p1: (d_405_v41_).issubset(_dafny.Set({(self).f25}))})
                d_412_v47_: str
                d_412_v47_ = _dafny.CodePoint('j')
                d_413_v48_: _dafny.MultiSet
                d_413_v48_ = _dafny.MultiSet([d_412_v47_, d_412_v47_, _dafny.CodePoint('t')])
                d_414_v49_: _dafny.Map
                d_414_v49_ = _dafny.Map({d_412_v47_: d_413_v48_})
                rhs57_ = d_410_i8_
                rhs58_ = d_411_v46_
                rhs59_ = p1
                rhs60_ = d_404_v40_
                rhs61_ = d_414_v49_
                lhs44_ = globalState
                lhs45_ = globalState
                r0 = rhs57_
                d_411_v46_ = rhs58_
                lhs44_.f12 = rhs59_
                lhs45_.f10 = rhs60_
                d_414_v49_ = rhs61_
                index59_ = default__.safeIndex(730, (d_409_v45_).length(0))
                (d_409_v45_)[index59_] = p1
                d_415_v50_: _dafny.Seq
                d_415_v50_ = _dafny.SeqWithoutIsStrInference([p1, p1, not (p1) or (p1), p1])
                index60_ = default__.safeIndex(730, (d_409_v45_).length(0))
                (d_409_v45_)[index60_] = (d_415_v50_)[default__.safeIndex(default__.fm31(d_404_v40_, globalState), len(d_415_v50_))]
                r0 = d_410_i8_
                (globalState).f12 = p1
            elif source12_.is_DC13:
                d_416___mcc_h0_ = source12_.cf19
                d_417___mcc_h1_ = source12_.cf20
                d_418___mcc_h2_ = source12_.cf21
                d_419___mcc_h3_ = source12_.cf22
                d_420_cf22_ = d_419___mcc_h3_
                d_421_cf21_ = d_418___mcc_h2_
                d_422_cf20_ = d_417___mcc_h1_
                d_423_cf19_ = d_416___mcc_h0_
                (globalState).f23 = _dafny.CodePoint('k')
                d_423_cf19_ = (d_423_cf19_) - ((-388) - (d_423_cf19_))
                d_424_v51_: C0
                nw58_ = C0()
                nw58_.ctor__()
                d_424_v51_ = nw58_
                d_425_v52_: _dafny.Array
                nw59_ = _dafny.Array(int(0), 29)
                d_425_v52_ = nw59_
                d_426_v53_: _dafny.MultiSet
                d_426_v53_ = _dafny.MultiSet([d_425_v52_, d_425_v52_, (d_425_v52_ if d_421_cf21_ else d_425_v52_), d_425_v52_])
                rhs62_ = (d_423_cf19_) - ((251) * (d_423_cf19_))
                rhs63_ = d_426_v53_
                r0 = rhs62_
                d_426_v53_ = rhs63_
            elif source12_.is_DC11:
                d_427___mcc_h4_ = source12_.cf18
                d_428_cf18_ = d_427___mcc_h4_
                d_429_v54_: _dafny.Map
                d_429_v54_ = _dafny.Map({D5_DC17(len(default__.fm35(d_410_i8_, globalState)), p1): p1})
                d_430_v55_: _dafny.MultiSet
                d_430_v55_ = _dafny.MultiSet([d_429_v54_, (d_429_v54_) | (d_429_v54_)])
                rhs64_ = d_430_v55_
                rhs65_ = ((self).f25 if p1 else len(_dafny.SeqWithoutIsStrInference([(0) - (d_410_i8_) for d_431_i9_ in range(default__.abs(551))])))
                rhs66_ = default__.fm33((self).f25, not(p1), (self).f25, globalState)
                lhs46_ = globalState
                d_430_v55_ = rhs64_
                r0 = rhs65_
                lhs46_.f23 = rhs66_
                (globalState).f12 = default__.fm0((0) - (d_410_i8_), globalState)
                rhs67_ = d_410_i8_
                rhs68_ = (0) - (default__.safeDivisionInt(default__.safeDivisionInt((self).f25, len(d_404_v40_)), 226))
                r0 = rhs67_
                r0 = rhs68_
                d_432_v56_: _dafny.MultiSet
                d_432_v56_ = _dafny.MultiSet([p1])
                d_433_v57_: D0
                d_433_v57_ = D0_DC2(True)
                d_434_v58_: _dafny.Set
                d_434_v58_ = _dafny.Set({p1})
                rhs69_ = (d_432_v56_).set(not((d_433_v57_).cf3), default__.abs(d_410_i8_))
                rhs70_ = not((d_434_v58_).ispropersubset((d_434_v58_) - (d_434_v58_)))
                lhs47_ = globalState
                d_432_v56_ = rhs69_
                lhs47_.f12 = rhs70_
            elif True:
                d_435___mcc_h5_ = source12_.cf23
                d_436_cf23_ = d_435___mcc_h5_
                d_437_v59_: _dafny.Array
                nw60_ = _dafny.Array(int(0), 16)
                d_437_v59_ = nw60_
                index61_ = default__.safeIndex(482, (d_437_v59_).length(0))
                (d_437_v59_)[index61_] = ((0) - ((self).f25)) * ((self).f25)
                d_438_v60_: _dafny.MultiSet
                d_438_v60_ = _dafny.MultiSet([((self).f25) - ((self).f25)])
                index62_ = default__.safeIndex(482, (d_437_v59_).length(0))
                rhs71_ = (d_438_v60_).cardinality
                rhs72_ = p1
                rhs73_ = d_404_v40_
                rhs74_ = ((((_dafny.MultiSet([p1])).set(p1, default__.abs((d_438_v60_).cardinality))).cardinality) + (d_410_i8_)) == (d_410_i8_)
                lhs48_ = d_437_v59_
                lhs49_ = default__.safeIndex(482, (d_437_v59_).length(0))
                lhs50_ = globalState
                lhs51_ = globalState
                lhs52_ = globalState
                lhs48_[lhs49_] = rhs71_
                lhs50_.f12 = rhs72_
                lhs51_.f10 = rhs73_
                lhs52_.f12 = rhs74_
                (globalState).f12 = ((0) - ((0) - ((self).f25))) != ((d_437_v59_)[default__.safeIndex(482, (d_437_v59_).length(0))])
                d_439_v61_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                d_439_v61_ = nw61_
                d_440_v62_: _dafny.Map
                d_440_v62_ = _dafny.Map({d_439_v61_: p1})
                d_440_v62_ = (d_440_v62_).set(d_439_v61_, p1)
                index63_ = default__.safeIndex(482, (d_437_v59_).length(0))
                (d_437_v59_)[index63_] = d_410_i8_
            d_441_v63_: _dafny.Array
            def lambda17_(d_442_v41_):
                def lambda18_(d_443_i10_):
                    return d_442_v41_

                return lambda18_

            init9_ = lambda17_(d_405_v41_)
            nw62_ = _dafny.Array(None, 9)
            for i0_9_ in range(nw62_.length(0)):
                nw62_[i0_9_] = init9_(i0_9_)
            d_441_v63_ = nw62_
            index64_ = default__.safeIndex(59, (d_441_v63_).length(0))
            (d_441_v63_)[index64_] = d_405_v41_
            index65_ = default__.safeIndex(59, (d_441_v63_).length(0))
            (d_441_v63_)[index65_] = (D4_DC11(_dafny.Set({d_410_i8_}))).cf18
            d_444_v64_: _dafny.Array
            nw63_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_444_v64_ = nw63_
            d_445_v65_: _dafny.MultiSet
            d_445_v65_ = _dafny.MultiSet([(self).f25])
            d_446_v66_: _dafny.Seq
            d_446_v66_ = _dafny.SeqWithoutIsStrInference([d_410_i8_, len(d_404_v40_), d_410_i8_, (d_445_v65_).cardinality])
            d_447_v67_: _dafny.Seq
            d_447_v67_ = d_446_v66_
            d_448_v68_: _dafny.Seq
            d_448_v68_ = _dafny.SeqWithoutIsStrInference([p1])
            d_449_v69_: _dafny.Array
            nw64_ = _dafny.Array(None, 11)
            nw64_[int(0)] = _dafny.SeqWithoutIsStrInference([(self).f25 for d_450_i11_ in range(default__.abs(433))])
            nw64_[int(1)] = _dafny.SeqWithoutIsStrInference([d_410_i8_ for d_451_i12_ in range(default__.abs(671))])
            nw64_[int(2)] = d_446_v66_
            nw64_[int(3)] = d_446_v66_
            nw64_[int(4)] = (d_447_v67_)
            nw64_[int(5)] = _dafny.SeqWithoutIsStrInference([(self).f25, len(d_448_v68_)])
            nw64_[int(6)] = d_446_v66_
            nw64_[int(7)] = d_446_v66_
            nw64_[int(8)] = d_446_v66_
            nw64_[int(9)] = d_446_v66_
            nw64_[int(10)] = d_446_v66_
            d_449_v69_ = nw64_
            index66_ = default__.safeIndex(108, (d_444_v64_).length(0))
            (d_444_v64_)[index66_] = d_449_v69_
            index67_ = default__.safeIndex(108, (d_444_v64_).length(0))
            (d_444_v64_)[index67_] = d_449_v69_
            if False:
                (globalState).f10 = ((d_404_v40_) + (d_404_v40_)) + (d_404_v40_)
                (globalState).f12 = False
                d_452_v70_: _dafny.Array
                def lambda19_(d_453_p1_):
                    def lambda20_(d_454_i13_):
                        return default__.safeDivisionInt(d_454_i13_, len(_dafny.Map({d_453_p1_: _dafny.CodePoint('j')})))

                    return lambda20_

                init10_ = lambda19_(p1)
                nw65_ = _dafny.Array(None, 18)
                for i0_10_ in range(nw65_.length(0)):
                    nw65_[i0_10_] = init10_(i0_10_)
                d_452_v70_ = nw65_
                d_452_v70_ = d_452_v70_
                d_455_v71_: str
                d_455_v71_ = _dafny.CodePoint('k')
                r0 = len((((d_404_v40_) + ((d_404_v40_) + ((d_404_v40_).set(default__.safeIndex(d_410_i8_, len(d_404_v40_)), d_455_v71_)))).set(default__.safeIndex((0) - ((self).f25), len((d_404_v40_) + ((d_404_v40_) + ((d_404_v40_).set(default__.safeIndex(d_410_i8_, len(d_404_v40_)), d_455_v71_))))), default__.fm33(len(d_448_v68_), p1, (self).f25, globalState))).set(default__.safeIndex((self).f25, len(((d_404_v40_) + ((d_404_v40_) + ((d_404_v40_).set(default__.safeIndex(d_410_i8_, len(d_404_v40_)), d_455_v71_)))).set(default__.safeIndex((0) - ((self).f25), len((d_404_v40_) + ((d_404_v40_) + ((d_404_v40_).set(default__.safeIndex(d_410_i8_, len(d_404_v40_)), d_455_v71_))))), default__.fm33(len(d_448_v68_), p1, (self).f25, globalState)))), (d_455_v71_ if False else d_455_v71_)))
                d_456_v72_: C0
                nw66_ = C0()
                nw66_.ctor__()
                d_456_v72_ = nw66_
            elif True:
                r0 = len(d_446_v66_)
                (globalState).f12 = p1
                d_457_v73_: _dafny.Map
                d_457_v73_ = _dafny.Map({d_410_i8_: d_404_v40_})
                d_457_v73_ = (d_457_v73_).set((self).f25, (d_404_v40_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_458_i14_ in range(default__.abs(688))])))
                d_459_v74_: _dafny.Array
                nw67_ = _dafny.Array(D5.default()(), 22)
                d_459_v74_ = nw67_
                d_459_v74_ = d_459_v74_
                rhs75_ = ((d_404_v40_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_460_i15_ in range(default__.abs(747))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gltbwdt")))
                rhs76_ = (self).f25
                rhs77_ = d_409_v45_
                lhs53_ = globalState
                lhs53_.f10 = rhs75_
                r0 = rhs76_
                d_409_v45_ = rhs77_
        d_461_v75_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_461_v75_ = nw68_
        d_461_v75_ = d_461_v75_
        r0 = 178
        return r0


class C2:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def fm23(self, p0, p1, p2, globalState):
        return len(_dafny.Map({default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_462_i0_ in range(default__.abs(-603))])), 139): 231}))

    def fm24(self, p0, globalState):
        return 110

    def m16(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.Seq({})
        r3: D1 = D1.default()()
        d_463_v0_: int
        d_463_v0_ = 840
        hi2_ = ((0) - (d_463_v0_)) + ((self).fm24(d_463_v0_, globalState))
        for d_464_i0_ in range(d_463_v0_, hi2_):
            d_465_v1_: _dafny.Array
            def lambda21_(d_466_i1_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgmmbr"))

            init11_ = lambda21_
            nw69_ = _dafny.Array(None, 1)
            for i0_11_ in range(nw69_.length(0)):
                nw69_[i0_11_] = init11_(i0_11_)
            d_465_v1_ = nw69_
            d_467_v2_: _dafny.Seq
            d_467_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmlqj"))
            index68_ = default__.safeIndex(178, (d_465_v1_).length(0))
            (d_465_v1_)[index68_] = d_467_v2_
            d_468_v3_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.Map({}), 11)
            d_468_v3_ = nw70_
            d_469_v4_: _dafny.Array
            nw71_ = _dafny.Array(D3.default()(), 12)
            d_469_v4_ = nw71_
            d_470_v5_: _dafny.Map
            d_470_v5_ = _dafny.Map({d_469_v4_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_471_i2_ in range(default__.abs(416))]))})
            index69_ = default__.safeIndex(66, (d_468_v3_).length(0))
            (d_468_v3_)[index69_] = d_470_v5_
            d_472_v6_: str
            d_472_v6_ = _dafny.CodePoint('f')
            d_473_v7_: bool
            d_473_v7_ = False
            d_474_v8_: D1
            d_474_v8_ = D1_DC6(d_472_v6_, d_473_v7_, d_467_v2_)
            index70_ = default__.safeIndex(178, (d_465_v1_).length(0))
            index71_ = default__.safeIndex(66, (d_468_v3_).length(0))
            rhs78_ = (d_474_v8_).cf9
            rhs79_ = d_470_v5_
            lhs54_ = d_465_v1_
            lhs55_ = default__.safeIndex(178, (d_465_v1_).length(0))
            lhs56_ = d_468_v3_
            lhs57_ = default__.safeIndex(66, (d_468_v3_).length(0))
            lhs54_[lhs55_] = rhs78_
            lhs56_[lhs57_] = rhs79_
            if d_473_v7_:
                (globalState).f12 = default__.fm0(d_464_i0_, globalState)
                d_475_v9_: _dafny.Array
                nw72_ = _dafny.Array(False, 16)
                d_475_v9_ = nw72_
                index72_ = default__.safeIndex(419, (d_475_v9_).length(0))
                (d_475_v9_)[index72_] = not(d_473_v7_)
                index73_ = default__.safeIndex(419, (d_475_v9_).length(0))
                (d_475_v9_)[index73_] = default__.fm0(default__.safeDivisionInt(d_463_v0_, d_464_i0_), globalState)
                d_476_v10_: _dafny.Array
                def lambda22_(d_477_i3_):
                    return (d_477_i3_) * (981)

                init12_ = lambda22_
                nw73_ = _dafny.Array(None, 16)
                for i0_12_ in range(nw73_.length(0)):
                    nw73_[i0_12_] = init12_(i0_12_)
                d_476_v10_ = nw73_
                d_476_v10_ = d_476_v10_
                d_478_v11_: _dafny.Map
                d_478_v11_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_479_i4_ in range(default__.abs(885))]))) * (d_463_v0_): (d_464_i0_) < (708)})
                d_480_v12_: _dafny.Seq
                d_480_v12_ = _dafny.SeqWithoutIsStrInference([d_473_v7_])
                d_478_v11_ = (d_478_v11_).set((0) - (d_463_v0_), (d_480_v12_)[default__.safeIndex(d_464_i0_, len(d_480_v12_))])
                d_481_v13_: D1
                d_481_v13_ = D1_DC5()
                d_482_v14_: _dafny.Array
                nw74_ = _dafny.Array(None, 20)
                nw74_[int(0)] = d_481_v13_
                nw74_[int(1)] = D1_DC5()
                nw74_[int(2)] = d_481_v13_
                nw74_[int(3)] = d_481_v13_
                nw74_[int(4)] = d_481_v13_
                nw74_[int(5)] = (d_481_v13_ if d_473_v7_ else d_481_v13_)
                nw74_[int(6)] = d_481_v13_
                nw74_[int(7)] = d_481_v13_
                nw74_[int(8)] = d_481_v13_
                nw74_[int(9)] = default__.fm25(d_473_v7_, (d_475_v9_)[default__.safeIndex(419, (d_475_v9_).length(0))], d_464_i0_, (d_475_v9_)[default__.safeIndex(419, (d_475_v9_).length(0))], globalState)
                nw74_[int(10)] = D1_DC5()
                nw74_[int(11)] = d_481_v13_
                nw74_[int(12)] = D1_DC5()
                nw74_[int(13)] = d_481_v13_
                nw74_[int(14)] = (d_481_v13_ if not((d_475_v9_)[default__.safeIndex(419, (d_475_v9_).length(0))]) else default__.fm25(d_473_v7_, d_473_v7_, d_464_i0_, d_473_v7_, globalState))
                nw74_[int(15)] = d_481_v13_
                nw74_[int(16)] = d_481_v13_
                nw74_[int(17)] = d_481_v13_
                nw74_[int(18)] = default__.fm25(d_473_v7_, d_473_v7_, 153, d_473_v7_, globalState)
                nw74_[int(19)] = d_481_v13_
                d_482_v14_ = nw74_
                index74_ = default__.safeIndex(186, (d_482_v14_).length(0))
                (d_482_v14_)[index74_] = d_481_v13_
                index75_ = default__.safeIndex(186, (d_482_v14_).length(0))
                (d_482_v14_)[index75_] = d_481_v13_
            elif True:
                rhs80_ = d_463_v0_
                rhs81_ = d_473_v7_
                rhs82_ = (d_463_v0_) + (d_464_i0_)
                lhs58_ = globalState
                r1 = rhs80_
                lhs58_.f12 = rhs81_
                r1 = rhs82_
                (globalState).f10 = (((d_465_v1_)[default__.safeIndex(178, (d_465_v1_).length(0))]) + ((d_465_v1_)[default__.safeIndex(178, (d_465_v1_).length(0))])) + ((d_465_v1_)[default__.safeIndex(178, (d_465_v1_).length(0))])
                (globalState).f12 = d_473_v7_
                d_483_v15_: C0
                nw75_ = C0()
                nw75_.ctor__()
                d_483_v15_ = nw75_
                d_484_v16_: _dafny.Set
                d_484_v16_ = _dafny.Set({d_464_i0_})
                d_485_v18_: _dafny.Set
                def iife40_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in _dafny.IntegerRange(470, -771):
                        d_486_v17_: int = compr_32_
                        if ((470) <= (d_486_v17_)) and ((d_486_v17_) < (-771)):
                            coll32_[default__.safeDivisionInt(d_486_v17_, d_463_v0_)] = (d_465_v1_)[default__.safeIndex(178, (d_465_v1_).length(0))]
                    return _dafny.Map(coll32_)
                d_485_v18_ = _dafny.Set({d_484_v16_, d_484_v16_, (d_484_v16_) - (default__.fm27(d_473_v7_, iife40_()
                , globalState))})
                d_485_v18_ = d_485_v18_
            d_487_v19_: _dafny.Map
            d_487_v19_ = _dafny.Map({d_473_v7_: d_473_v7_})
            d_488_v20_: D4
            d_488_v20_ = D4_DC13(len((d_487_v19_) | (d_487_v19_)), not (d_473_v7_) or (d_473_v7_), d_473_v7_, d_473_v7_)
            d_488_v20_ = d_488_v20_
            d_489_v21_: _dafny.Array
            def lambda23_(d_490_v7_):
                def lambda24_(d_491_i5_):
                    return (_dafny.SeqWithoutIsStrInference([not(d_490_v7_)])) + (_dafny.SeqWithoutIsStrInference([d_490_v7_]))

                return lambda24_

            init13_ = lambda23_(d_473_v7_)
            nw76_ = _dafny.Array(None, 11)
            for i0_13_ in range(nw76_.length(0)):
                nw76_[i0_13_] = init13_(i0_13_)
            d_489_v21_ = nw76_
            d_492_v22_: _dafny.Seq
            d_492_v22_ = _dafny.SeqWithoutIsStrInference([d_473_v7_, d_473_v7_, d_473_v7_, d_473_v7_])
            index76_ = default__.safeIndex(866, (d_489_v21_).length(0))
            (d_489_v21_)[index76_] = d_492_v22_
            index77_ = default__.safeIndex(866, (d_489_v21_).length(0))
            (d_489_v21_)[index77_] = _dafny.SeqWithoutIsStrInference([True, d_473_v7_])
        if default__.fm0(-959, globalState):
            d_493_v23_: _dafny.Seq
            d_493_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjlplkxx"))
            d_494_v24_: _dafny.Seq
            d_494_v24_ = _dafny.SeqWithoutIsStrInference([d_463_v0_, len(d_493_v23_), d_463_v0_])
            d_495_v25_: bool
            d_495_v25_ = False
            d_496_v26_: _dafny.Map
            d_496_v26_ = _dafny.Map({(0) - ((d_494_v24_)[default__.safeIndex(d_463_v0_, len(d_494_v24_))]): d_495_v25_})
            d_497_v27_: _dafny.MultiSet
            d_497_v27_ = _dafny.MultiSet([_dafny.Map({(d_494_v24_)[default__.safeIndex(d_463_v0_, len(d_494_v24_))]: d_495_v25_}), d_496_v26_])
            (globalState).f12 = not((d_497_v27_).issubset(d_497_v27_))
            d_498_v28_: C0
            nw77_ = C0()
            nw77_.ctor__()
            d_498_v28_ = nw77_
            d_463_v0_ = (0) - ((self).fm23(default__.fm28(d_495_v25_, d_463_v0_, d_493_v23_, d_463_v0_, globalState), d_495_v25_, not (default__.fm0(d_463_v0_, globalState)) or (d_495_v25_), globalState))
            d_499_v29_: _dafny.Array
            def lambda25_(d_500_v0_):
                def lambda26_(d_501_i6_):
                    return default__.safeDivisionInt(d_501_i6_, d_500_v0_)

                return lambda26_

            init14_ = lambda25_(d_463_v0_)
            nw78_ = _dafny.Array(None, 12)
            for i0_14_ in range(nw78_.length(0)):
                nw78_[i0_14_] = init14_(i0_14_)
            d_499_v29_ = nw78_
            index78_ = default__.safeIndex(359, (d_499_v29_).length(0))
            (d_499_v29_)[index78_] = d_463_v0_
            index79_ = default__.safeIndex(359, (d_499_v29_).length(0))
            (d_499_v29_)[index79_] = d_463_v0_
            (globalState).f12 = d_495_v25_
        elif True:
            d_502_v30_: bool
            d_502_v30_ = False
            d_503_v31_: _dafny.Seq
            d_503_v31_ = _dafny.SeqWithoutIsStrInference([d_502_v30_, True, d_502_v30_])
            if (len(d_503_v31_)) == (d_463_v0_):
                d_504_v32_: _dafny.Array
                nw79_ = _dafny.Array(None, 8)
                nw79_[int(0)] = d_503_v31_
                nw79_[int(1)] = d_503_v31_
                nw79_[int(2)] = _dafny.SeqWithoutIsStrInference([d_502_v30_])
                nw79_[int(3)] = (d_503_v31_) + (default__.fm1(d_502_v30_, globalState))
                nw79_[int(4)] = _dafny.SeqWithoutIsStrInference([d_502_v30_])
                nw79_[int(5)] = d_503_v31_
                nw79_[int(6)] = d_503_v31_
                nw79_[int(7)] = d_503_v31_
                d_504_v32_ = nw79_
                index80_ = default__.safeIndex(500, (d_504_v32_).length(0))
                (d_504_v32_)[index80_] = d_503_v31_
                index81_ = default__.safeIndex(500, (d_504_v32_).length(0))
                (d_504_v32_)[index81_] = ((_dafny.SeqWithoutIsStrInference([d_502_v30_])) + (d_503_v31_)) + (((d_503_v31_).set(default__.safeIndex(d_463_v0_, len(d_503_v31_)), d_502_v30_)) + (_dafny.SeqWithoutIsStrInference([d_502_v30_])))
                d_505_v33_: C0
                nw80_ = C0()
                nw80_.ctor__()
                d_505_v33_ = nw80_
                (globalState).f12 = not(not (d_502_v30_) or (d_502_v30_))
                d_506_v34_: _dafny.Seq
                d_506_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fycjace"))
                d_507_v35_: _dafny.Seq
                d_507_v35_ = _dafny.SeqWithoutIsStrInference([(d_506_v34_)[default__.safeIndex(d_463_v0_, len(d_506_v34_))]])
                (globalState).f6 = default__.fm29((d_507_v35_) + (default__.fm2(d_463_v0_, d_463_v0_, globalState)), d_502_v30_, d_463_v0_, d_507_v35_, globalState)
                r0 = d_463_v0_
            elif True:
                d_508_v36_: _dafny.Seq
                d_508_v36_ = _dafny.SeqWithoutIsStrInference([d_463_v0_])
                r1 = default__.safeDivisionInt(len(d_508_v36_), d_463_v0_)
                def iife41_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(666, 838):
                        d_509_v37_: int = compr_33_
                        if ((666) <= (d_509_v37_)) and ((d_509_v37_) < (838)):
                            coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_509_v37_, d_463_v0_)]))
                    return _dafny.Set(coll33_)
                r0 = len(iife41_()
                )
                d_510_v38_: _dafny.Map
                d_510_v38_ = _dafny.Map({(d_502_v30_) or (d_502_v30_): d_502_v30_})
                d_510_v38_ = d_510_v38_
                d_511_v39_: _dafny.Map
                d_511_v39_ = _dafny.Map({(0) - (d_463_v0_): d_502_v30_})
                d_512_v40_: _dafny.Map
                d_512_v40_ = _dafny.Map({d_502_v30_: d_463_v0_})
                d_513_v41_: _dafny.Map
                d_513_v41_ = _dafny.Map({(d_463_v0_) != ((self).fm23(_dafny.Map({d_502_v30_: d_463_v0_}), False, d_502_v30_, globalState)): default__.fm30((d_511_v39_).set((self).fm23(d_512_v40_, d_502_v30_, d_502_v30_, globalState), default__.fm0(d_463_v0_, globalState)), globalState)})
                d_514_v42_: str
                d_514_v42_ = _dafny.CodePoint('x')
                d_515_v43_: _dafny.Seq
                d_515_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hymwetbrg"))
                d_516_v44_: D1
                d_516_v44_ = D1_DC6(d_514_v42_, not(d_502_v30_), d_515_v43_)
                d_513_v41_ = (d_513_v41_).set(d_502_v30_, d_516_v44_)
                d_517_v45_: _dafny.Array
                nw81_ = _dafny.Array(int(0), 17)
                d_517_v45_ = nw81_
                d_518_v46_: _dafny.Array
                d_519_v47_: int
                out20_: _dafny.Array
                out21_: int
                out20_, out21_ = (self).m17((d_463_v0_) > (d_463_v0_), d_517_v45_, d_502_v30_, (d_502_v30_ if d_502_v30_ else d_502_v30_), globalState)
                d_518_v46_ = out20_
                d_519_v47_ = out21_
            d_520_v48_: C0
            nw82_ = C0()
            nw82_.ctor__()
            d_520_v48_ = nw82_
            (globalState).f12 = not(d_502_v30_)
            d_521_v49_: _dafny.Seq
            d_521_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "singysn"))
            r1 = (d_463_v0_) - (len(d_521_v49_))
            d_522_v50_: _dafny.Map
            d_522_v50_ = _dafny.Map({d_520_v48_: d_463_v0_})
            d_523_v51_: _dafny.Set
            d_523_v51_ = _dafny.Set({d_463_v0_})
            d_524_v52_: _dafny.Map
            d_524_v52_ = _dafny.Map({d_523_v51_: d_523_v51_})
            r1 = ((d_522_v50_)[d_520_v48_] if (d_520_v48_) in (d_522_v50_) else len(d_524_v52_))
        hi3_ = d_463_v0_
        for d_525_i7_ in range((-324) - (d_463_v0_), hi3_):
            d_526_v53_: _dafny.Array
            def lambda27_(d_527_i8_):
                return (d_527_i8_) * (len(_dafny.Set({False})))

            init15_ = lambda27_
            nw83_ = _dafny.Array(None, 19)
            for i0_15_ in range(nw83_.length(0)):
                nw83_[i0_15_] = init15_(i0_15_)
            d_526_v53_ = nw83_
            d_526_v53_ = d_526_v53_
            d_528_v54_: _dafny.Map
            d_528_v54_ = _dafny.Map({d_463_v0_: d_525_i7_})
            d_529_v55_: _dafny.Seq
            d_529_v55_ = _dafny.SeqWithoutIsStrInference([((d_528_v54_)[d_525_i7_] if (d_525_i7_) in (d_528_v54_) else 655), d_463_v0_, d_525_i7_])
            d_530_v56_: bool
            d_530_v56_ = False
            d_531_v57_: _dafny.Map
            d_531_v57_ = _dafny.Map({d_529_v55_: d_530_v56_})
            d_532_v58_: _dafny.Seq
            d_532_v58_ = _dafny.SeqWithoutIsStrInference([d_530_v56_])
            d_531_v57_ = (d_531_v57_).set(d_529_v55_, (d_530_v56_) == ((d_532_v58_)[default__.safeIndex(d_525_i7_, len(d_532_v58_))]))
            index82_ = default__.safeIndex(588, (d_526_v53_).length(0))
            (d_526_v53_)[index82_] = 182
            d_533_v59_: _dafny.Seq
            d_533_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bunpdm"))
            d_534_v60_: _dafny.MultiSet
            d_534_v60_ = _dafny.MultiSet([len(d_533_v59_)])
            index83_ = default__.safeIndex(588, (d_526_v53_).length(0))
            rhs83_ = not (d_530_v56_) or (d_530_v56_)
            rhs84_ = (d_533_v59_) + (d_533_v59_)
            rhs85_ = ((d_534_v60_)[(d_525_i7_ if d_530_v56_ else d_525_i7_)] if ((d_525_i7_ if d_530_v56_ else d_525_i7_)) in (d_534_v60_) else (0) - (d_525_i7_))
            rhs86_ = (0) - (d_463_v0_)
            rhs87_ = d_531_v57_
            lhs59_ = globalState
            lhs60_ = d_526_v53_
            lhs61_ = default__.safeIndex(588, (d_526_v53_).length(0))
            d_530_v56_ = rhs83_
            lhs59_.f10 = rhs84_
            lhs60_[lhs61_] = rhs85_
            r1 = rhs86_
            d_531_v57_ = rhs87_
            d_535_v61_: D0
            d_535_v61_ = D0_DC3(d_525_i7_, (True if d_530_v56_ else d_530_v56_))
            d_535_v61_ = d_535_v61_
        r1 = d_463_v0_
        d_536_v62_: _dafny.Array
        nw84_ = _dafny.Array(False, 26)
        d_536_v62_ = nw84_
        d_537_v63_: _dafny.Set
        d_537_v63_ = _dafny.Set({d_536_v62_, d_536_v62_})
        d_538_i9_: int
        d_538_i9_ = 0
        with _dafny.label("4"):
            while (_dafny.Set({d_536_v62_, d_536_v62_})).isdisjoint((d_537_v63_) | (_dafny.Set({d_536_v62_}))):
                with _dafny.c_label("4"):
                    if (d_538_i9_) >= (100):
                        raise _dafny.Break("4")
                    d_538_i9_ = (d_538_i9_) + (1)
                    (globalState).f12 = (d_463_v0_) < (default__.safeModuloInt(d_463_v0_, (0) - (d_463_v0_)))
                    d_539_v64_: bool
                    d_539_v64_ = True
                    if d_539_v64_:
                        r1 = ((0) - (d_463_v0_)) - (d_463_v0_)
                        d_540_v65_: _dafny.Array
                        def lambda28_(d_541_v0_):
                            def lambda29_(d_542_i10_):
                                return (d_542_i10_) - (d_541_v0_)

                            return lambda29_

                        init16_ = lambda28_(d_463_v0_)
                        nw85_ = _dafny.Array(None, 6)
                        for i0_16_ in range(nw85_.length(0)):
                            nw85_[i0_16_] = init16_(i0_16_)
                        d_540_v65_ = nw85_
                        d_543_v66_: _dafny.MultiSet
                        d_543_v66_ = _dafny.MultiSet([d_540_v65_])
                        d_544_v67_: _dafny.MultiSet
                        d_544_v67_ = _dafny.MultiSet([(d_543_v66_).ispropersubset(_dafny.MultiSet([d_540_v65_])), not(d_539_v64_)])
                        r1 = (d_544_v67_).cardinality
                        d_545_v68_: _dafny.Set
                        d_545_v68_ = _dafny.Set({d_539_v64_, d_539_v64_})
                        d_546_v69_: _dafny.Map
                        d_546_v69_ = _dafny.Map({d_539_v64_: d_545_v68_})
                        d_547_v70_: _dafny.Map
                        d_547_v70_ = _dafny.Map({d_546_v69_: d_463_v0_})
                        d_547_v70_ = (d_547_v70_).set((d_546_v69_).set(False, d_545_v68_), d_463_v0_)
                        d_463_v0_ = (d_463_v0_ if (d_463_v0_) > (d_463_v0_) else (0) - (d_463_v0_))
                        d_548_v71_: str
                        d_548_v71_ = _dafny.CodePoint('h')
                        (globalState).f23 = d_548_v71_
                    elif True:
                        d_463_v0_ = (0) - (d_463_v0_)
                        d_549_v72_: D4
                        d_549_v72_ = D4_DC12()
                        d_549_v72_ = (d_549_v72_ if d_539_v64_ else d_549_v72_)
                        d_550_v73_: _dafny.Seq
                        d_550_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uunsmmj"))
                        d_551_v74_: _dafny.MultiSet
                        d_551_v74_ = _dafny.MultiSet([d_539_v64_])
                        d_552_v75_: str
                        d_552_v75_ = _dafny.CodePoint('v')
                        d_553_v76_: D1
                        d_553_v76_ = D1_DC6(d_552_v75_, d_539_v64_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyxjlokv")))
                        d_554_v77_: _dafny.Array
                        nw86_ = _dafny.Array(None, 7)
                        nw86_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qu"))
                        nw86_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_555_i11_ in range(default__.abs(198))])
                        nw86_[int(2)] = d_550_v73_
                        nw86_[int(3)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_556_i12_ in range(default__.abs(-519))])).set(default__.safeIndex((d_551_v74_).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_557_i12_ in range(default__.abs(-519))]))), d_552_v75_)
                        nw86_[int(4)] = _dafny.SeqWithoutIsStrInference([d_552_v75_ for d_558_i13_ in range(default__.abs(219))])
                        nw86_[int(5)] = (d_553_v76_).cf9
                        nw86_[int(6)] = d_550_v73_
                        d_554_v77_ = nw86_
                        d_559_v78_: T1
                        nw87_ = C1()
                        nw87_.ctor__(d_463_v0_, d_554_v77_)
                        d_559_v78_ = nw87_
                        d_560_v79_: _dafny.Set
                        d_560_v79_ = _dafny.Set({d_559_v78_, d_559_v78_})
                        (globalState).f12 = ((d_560_v79_).intersection(d_560_v79_)).issubset((d_560_v79_ if d_539_v64_ else d_560_v79_))
                        r1 = (d_559_v78_).f25
                        d_536_v62_ = d_536_v62_
                    d_561_v80_: _dafny.Seq
                    d_561_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybi"))
                    d_463_v0_ = len(d_561_v80_)
                    (globalState).f12 = d_539_v64_
                    pass
            pass
        d_562_v81_: bool
        d_562_v81_ = False
        d_563_i14_: int
        d_563_i14_ = 0
        with _dafny.label("5"):
            while d_562_v81_:
                with _dafny.c_label("5"):
                    if (d_563_i14_) >= (100):
                        raise _dafny.Break("5")
                    d_563_i14_ = (d_563_i14_) + (1)
                    (globalState).f12 = d_562_v81_
                    d_463_v0_ = d_463_v0_
                    d_564_v82_: _dafny.Seq
                    d_564_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvsjnjdbt"))
                    d_565_v83_: str
                    d_565_v83_ = _dafny.CodePoint('j')
                    d_566_v84_: _dafny.Map
                    d_566_v84_ = _dafny.Map({(d_463_v0_) + (len(d_564_v82_)): (d_565_v83_) not in (d_564_v82_)})
                    d_567_v85_: _dafny.MultiSet
                    d_567_v85_ = _dafny.MultiSet([d_562_v81_, d_562_v81_, d_562_v81_])
                    d_568_v86_: _dafny.Map
                    d_568_v86_ = _dafny.Map({d_562_v81_: 868})
                    d_566_v84_ = (d_566_v84_).set((86) - (((d_567_v85_).set(d_562_v81_, default__.abs((self).fm23(d_568_v86_, d_562_v81_, d_562_v81_, globalState)))).cardinality), d_562_v81_)
                    d_569_v87_: _dafny.MultiSet
                    d_569_v87_ = _dafny.MultiSet([d_564_v82_, d_564_v82_, d_564_v82_])
                    d_570_v88_: _dafny.Seq
                    d_570_v88_ = _dafny.SeqWithoutIsStrInference([True, d_562_v81_, d_562_v81_, d_562_v81_])
                    d_571_v89_: _dafny.Map
                    d_571_v89_ = _dafny.Map({d_562_v81_: d_565_v83_})
                    d_572_v90_: _dafny.Array
                    nw88_ = _dafny.Array(None, 7)
                    nw88_[int(0)] = default__.safeModuloInt(d_463_v0_, d_463_v0_)
                    nw88_[int(1)] = d_463_v0_
                    nw88_[int(2)] = ((d_569_v87_)[d_564_v82_] if (d_564_v82_) in (d_569_v87_) else d_463_v0_)
                    nw88_[int(3)] = len((d_570_v88_) + (d_570_v88_))
                    nw88_[int(4)] = d_463_v0_
                    nw88_[int(5)] = default__.safeDivisionInt(len(d_571_v89_), d_463_v0_)
                    nw88_[int(6)] = d_463_v0_
                    d_572_v90_ = nw88_
                    index84_ = default__.safeIndex(183, (d_572_v90_).length(0))
                    (d_572_v90_)[index84_] = len(d_564_v82_)
                    index85_ = default__.safeIndex(183, (d_572_v90_).length(0))
                    rhs88_ = ((0) - (d_463_v0_)) + (3)
                    rhs89_ = d_565_v83_
                    rhs90_ = default__.safeModuloInt(default__.safeModuloInt((0) - (d_463_v0_), d_463_v0_), d_463_v0_)
                    rhs91_ = default__.safeModuloInt(d_463_v0_, d_463_v0_)
                    lhs62_ = d_572_v90_
                    lhs63_ = default__.safeIndex(183, (d_572_v90_).length(0))
                    r1 = rhs88_
                    d_565_v83_ = rhs89_
                    lhs62_[lhs63_] = rhs90_
                    d_463_v0_ = rhs91_
                    pass
            pass
        r0 = d_463_v0_
        r1 = d_463_v0_
        d_573_v91_: _dafny.Seq
        d_573_v91_ = _dafny.SeqWithoutIsStrInference([True])
        d_574_v92_: _dafny.MultiSet
        d_574_v92_ = _dafny.MultiSet([(d_573_v91_)[default__.safeIndex(len(d_573_v91_), len(d_573_v91_))]])
        d_575_v93_: _dafny.Map
        d_575_v93_ = _dafny.Map({630: d_562_v81_})
        d_576_v94_: _dafny.Map
        d_576_v94_ = _dafny.Map({d_463_v0_: d_575_v93_})
        r2 = ((_dafny.SeqWithoutIsStrInference([d_463_v0_, (0) - (d_463_v0_)])).set(default__.safeIndex(((d_574_v92_).cardinality) * (d_463_v0_), len(_dafny.SeqWithoutIsStrInference([d_463_v0_, (0) - (d_463_v0_)]))), default__.safeDivisionInt(d_463_v0_, len(d_576_v94_)))).set(default__.safeIndex(d_463_v0_, len((_dafny.SeqWithoutIsStrInference([d_463_v0_, (0) - (d_463_v0_)])).set(default__.safeIndex(((d_574_v92_).cardinality) * (d_463_v0_), len(_dafny.SeqWithoutIsStrInference([d_463_v0_, (0) - (d_463_v0_)]))), default__.safeDivisionInt(d_463_v0_, len(d_576_v94_))))), (d_463_v0_ if d_562_v81_ else d_463_v0_))
        d_577_v95_: D1
        d_577_v95_ = D1_DC4(_dafny.SeqWithoutIsStrInference([d_562_v81_]))
        r3 = (d_577_v95_ if d_562_v81_ else d_577_v95_)
        return r0, r1, r2, r3

    def m17(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_578_v0_: _dafny.Array
        nw89_ = _dafny.Array(False, 4)
        d_578_v0_ = nw89_
        d_579_v1_: _dafny.Seq
        d_579_v1_ = _dafny.SeqWithoutIsStrInference([p1])
        d_580_v2_: _dafny.Map
        d_580_v2_ = _dafny.Map({(d_579_v1_) + (d_579_v1_): d_578_v0_})
        d_578_v0_ = ((d_580_v2_)[d_579_v1_] if (d_579_v1_) in (d_580_v2_) else d_578_v0_)
        d_581_v3_: int
        d_581_v3_ = 136
        r1 = (0) - (d_581_v3_)
        d_582_v4_: _dafny.Map
        d_582_v4_ = _dafny.Map({(0) - (d_581_v3_): p2})
        d_583_i0_: int
        d_583_i0_ = 0
        with _dafny.label("6"):
            while ((d_582_v4_)[d_581_v3_] if (d_581_v3_) in (d_582_v4_) else p0):
                with _dafny.c_label("6"):
                    if (d_583_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_583_i0_ = (d_583_i0_) + (1)
                    d_584_v5_: _dafny.Set
                    d_584_v5_ = _dafny.Set({d_578_v0_, d_578_v0_, d_578_v0_})
                    d_585_v6_: D3
                    d_585_v6_ = D3_DC9(d_584_v5_)
                    d_586_v7_: str
                    d_586_v7_ = _dafny.CodePoint('u')
                    d_587_v8_: _dafny.MultiSet
                    d_587_v8_ = _dafny.MultiSet([d_581_v3_])
                    pat_let_tv8_ = d_578_v0_
                    pat_let_tv9_ = d_578_v0_
                    pat_let_tv10_ = d_578_v0_
                    pat_let_tv11_ = d_578_v0_
                    def iife42_(_pat_let4_0):
                        def iife43_(d_588_dt__update__tmp_h0_):
                            def iife44_(_pat_let5_0):
                                def iife45_(d_589_dt__update_hcf12_h0_):
                                    return D3_DC9(d_589_dt__update_hcf12_h0_)
                                return iife45_(_pat_let5_0)
                            return iife44_(_dafny.Set({pat_let_tv8_, pat_let_tv9_, pat_let_tv10_, pat_let_tv11_}))
                        return iife43_(_pat_let4_0)
                    d_585_v6_ = (d_585_v6_ if (len(default__.fm36(p2, d_582_v4_, d_586_v7_, globalState))) in (d_587_v8_) else iife42_(D3_DC9(d_584_v5_)))
                    d_590_v9_: C0
                    nw90_ = C0()
                    nw90_.ctor__()
                    d_590_v9_ = nw90_
                    d_591_v10_: _dafny.Array
                    nw91_ = _dafny.Array(None, 1)
                    nw91_[int(0)] = _dafny.SeqWithoutIsStrInference([d_586_v7_ for d_592_i1_ in range(default__.abs(281))])
                    d_591_v10_ = nw91_
                    d_593_v11_: T1
                    nw92_ = C1()
                    nw92_.ctor__(d_581_v3_, d_591_v10_)
                    d_593_v11_ = nw92_
                    d_593_v11_ = d_593_v11_
                    d_594_v12_: _dafny.Seq
                    d_594_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppb"))
                    d_581_v3_ = default__.fm31((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rudhwqxa"))) + (d_594_v12_), globalState)
                    pass
            pass
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_578_v0_).length(0)):
            d_595_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_595_i2_)) and ((d_595_i2_) < ((d_578_v0_).length(0)))):
                (d_578_v0_)[(d_595_i2_)] = p0
        d_596_v13_: _dafny.Set
        d_596_v13_ = _dafny.Set({d_581_v3_, d_581_v3_})
        d_597_v14_: _dafny.Array
        nw93_ = _dafny.Array(None, 6)
        nw93_[int(0)] = d_581_v3_
        nw93_[int(1)] = d_581_v3_
        nw93_[int(2)] = d_581_v3_
        nw93_[int(3)] = 796
        nw93_[int(4)] = len(d_596_v13_)
        nw93_[int(5)] = d_581_v3_
        d_597_v14_ = nw93_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_597_v14_).length(0)):
            d_598_i3_: int = guard_loop_3_
            if (True) and (((0) <= (d_598_i3_)) and ((d_598_i3_) < ((d_597_v14_).length(0)))):
                (d_597_v14_)[(d_598_i3_)] = default__.safeModuloInt(d_598_i3_, d_581_v3_)
        d_599_v15_: _dafny.Array
        nw94_ = _dafny.Array(D1.default()(), 16)
        d_599_v15_ = nw94_
        d_600_v16_: _dafny.Seq
        d_600_v16_ = _dafny.SeqWithoutIsStrInference([not(not(p3))])
        index86_ = default__.safeIndex(746, (d_599_v15_).length(0))
        (d_599_v15_)[index86_] = D1_DC4(d_600_v16_)
        d_601_v17_: D1
        d_601_v17_ = D1_DC4(_dafny.SeqWithoutIsStrInference([True]))
        index87_ = default__.safeIndex(746, (d_599_v15_).length(0))
        (d_599_v15_)[index87_] = d_601_v17_
        r0 = p1
        d_602_v18_: _dafny.Map
        d_602_v18_ = _dafny.Map({p0: not(p2)})
        d_603_v19_: _dafny.Seq
        d_603_v19_ = _dafny.SeqWithoutIsStrInference([len(d_602_v18_)])
        r1 = default__.safeDivisionInt(len(((d_603_v19_).set(default__.safeIndex(d_581_v3_, len(d_603_v19_)), d_581_v3_) if default__.fm0(d_581_v3_, globalState) else _dafny.SeqWithoutIsStrInference([-40]))), -753)
        return r0, r1


class C3(T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        self._f37: int = int(0)
        self._f38: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f37, f38, f25, f26):
        (self)._f37 = f37
        (self)._f38 = f38
        (self)._f25 = f25
        (self)._f26 = f26

    def fm38(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([(self).f37 for d_604_i0_ in range(default__.abs(650))])) + (_dafny.SeqWithoutIsStrInference([(self).f25, (self).f37]))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_605_v0_: _dafny.Array
        nw95_ = _dafny.Array(_dafny.MultiSet({}), 29)
        d_605_v0_ = nw95_
        d_606_v1_: _dafny.Array
        def lambda30_(d_607_p1_):
            def lambda31_(d_608_i0_):
                return not(d_607_p1_)

            return lambda31_

        init17_ = lambda30_(p1)
        nw96_ = _dafny.Array(None, 16)
        for i0_17_ in range(nw96_.length(0)):
            nw96_[i0_17_] = init17_(i0_17_)
        d_606_v1_ = nw96_
        index88_ = default__.safeIndex(742, (d_605_v0_).length(0))
        (d_605_v0_)[index88_] = _dafny.MultiSet([d_606_v1_, d_606_v1_, d_606_v1_])
        d_609_v2_: _dafny.MultiSet
        d_609_v2_ = _dafny.MultiSet([d_606_v1_])
        index89_ = default__.safeIndex(742, (d_605_v0_).length(0))
        (d_605_v0_)[index89_] = (d_609_v2_).intersection(_dafny.MultiSet([d_606_v1_]))
        d_610_v3_: C0
        nw97_ = C0()
        nw97_.ctor__()
        d_610_v3_ = nw97_
        r0 = (self).f37
        if p2:
            index90_ = default__.safeIndex(355, (d_606_v1_).length(0))
            (d_606_v1_)[index90_] = p2
            d_611_v4_: _dafny.Seq
            d_611_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gp"))
            index91_ = default__.safeIndex(355, (d_606_v1_).length(0))
            (d_606_v1_)[index91_] = (375) >= (len(d_611_v4_))
            d_612_v5_: _dafny.Array
            nw98_ = _dafny.Array(False, 18)
            d_612_v5_ = nw98_
            d_613_v6_: _dafny.Map
            d_613_v6_ = _dafny.Map({d_612_v5_: (self).f37})
            d_614_v7_: _dafny.Seq
            d_614_v7_ = _dafny.SeqWithoutIsStrInference([(d_606_v1_)[default__.safeIndex(355, (d_606_v1_).length(0))]])
            d_615_v8_: _dafny.Set
            d_615_v8_ = _dafny.Set({(self).f37, (self).f25})
            d_616_v9_: _dafny.Seq
            d_616_v9_ = _dafny.SeqWithoutIsStrInference([len(d_614_v7_), len(d_615_v8_)])
            d_617_v10_: D6
            d_617_v10_ = D6_DC18((_dafny.Map({d_612_v5_: (self).f37})) | ((d_613_v6_).set(d_612_v5_, len(d_616_v9_))))
            d_618_v11_: D0
            d_618_v11_ = D0_DC1(p2, p2)
            d_619_v12_: _dafny.Set
            d_619_v12_ = _dafny.Set({(d_618_v11_).cf2})
            d_620_v13_: _dafny.Map
            d_620_v13_ = _dafny.Map({_dafny.MultiSet([len(d_619_v12_), (self).f25]): (d_606_v1_)[default__.safeIndex(355, (d_606_v1_).length(0))]})
            d_621_v14_: D1
            d_621_v14_ = D1_DC5()
            d_622_v15_: _dafny.Array
            nw99_ = _dafny.Array(None, 24)
            nw99_[int(0)] = d_621_v14_
            nw99_[int(1)] = d_621_v14_
            nw99_[int(2)] = d_621_v14_
            nw99_[int(3)] = d_621_v14_
            nw99_[int(4)] = d_621_v14_
            nw99_[int(5)] = d_621_v14_
            nw99_[int(6)] = d_621_v14_
            nw99_[int(7)] = d_621_v14_
            nw99_[int(8)] = d_621_v14_
            nw99_[int(9)] = d_621_v14_
            nw99_[int(10)] = d_621_v14_
            nw99_[int(11)] = D1_DC5()
            nw99_[int(12)] = d_621_v14_
            nw99_[int(13)] = D1_DC5()
            nw99_[int(14)] = d_621_v14_
            nw99_[int(15)] = d_621_v14_
            nw99_[int(16)] = d_621_v14_
            nw99_[int(17)] = d_621_v14_
            nw99_[int(18)] = d_621_v14_
            nw99_[int(19)] = d_621_v14_
            nw99_[int(20)] = d_621_v14_
            nw99_[int(21)] = d_621_v14_
            nw99_[int(22)] = d_621_v14_
            nw99_[int(23)] = d_621_v14_
            d_622_v15_ = nw99_
            index92_ = default__.safeIndex(460, (d_622_v15_).length(0))
            (d_622_v15_)[index92_] = D1_DC5()
            d_623_v16_: _dafny.Array
            nw100_ = _dafny.Array(int(0), 8)
            d_623_v16_ = nw100_
            d_624_v17_: D7
            d_624_v17_ = D7_DC21(p3)
            index93_ = default__.safeIndex(460, (d_622_v15_).length(0))
            rhs92_ = d_617_v10_
            rhs93_ = d_620_v13_
            rhs94_ = d_621_v14_
            rhs95_ = (p1) == ((d_606_v1_)[default__.safeIndex(355, (d_606_v1_).length(0))])
            rhs96_ = (d_624_v17_).cf32
            lhs64_ = d_622_v15_
            lhs65_ = default__.safeIndex(460, (d_622_v15_).length(0))
            lhs66_ = globalState
            d_617_v10_ = rhs92_
            d_620_v13_ = rhs93_
            lhs64_[lhs65_] = rhs94_
            lhs66_.f12 = rhs95_
            d_623_v16_ = rhs96_
            d_625_v18_: _dafny.Map
            d_625_v18_ = _dafny.Map({((self).f25) > ((self).f37): (0) - (default__.safeModuloInt((self).f25, (self).f25))})
            d_625_v18_ = (d_625_v18_).set((self).f38, (self).f25)
            d_626_v19_: str
            d_626_v19_ = _dafny.CodePoint('p')
            d_627_v20_: D1
            d_627_v20_ = D1_DC6(d_626_v19_, p1, d_611_v4_)
            d_628_v22_: _dafny.Map
            def iife46_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in (d_615_v8_).Elements:
                    d_629_v21_: int = compr_34_
                    if (d_629_v21_) in (d_615_v8_):
                        coll34_ = coll34_.union(_dafny.Set([(d_629_v21_) * ((0) - ((_dafny.MultiSet([True, True, False])).cardinality))]))
                return _dafny.Set(coll34_)
            d_628_v22_ = _dafny.Map({d_627_v20_: iife46_()
            })
            d_615_v8_ = ((d_628_v22_)[d_627_v20_] if (d_627_v20_) in (d_628_v22_) else (d_615_v8_).intersection(d_615_v8_))
            d_630_v23_: _dafny.Map
            d_630_v23_ = _dafny.Map({(self).f25: _dafny.SeqWithoutIsStrInference([d_626_v19_ for d_631_i1_ in range(default__.abs(112))])})
            d_632_v24_: _dafny.Array
            nw101_ = _dafny.Array(None, 12)
            nw101_[int(0)] = (d_611_v4_) + (d_611_v4_)
            nw101_[int(1)] = (d_611_v4_) + (d_611_v4_)
            nw101_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            nw101_[int(3)] = d_611_v4_
            nw101_[int(4)] = d_611_v4_
            nw101_[int(5)] = d_611_v4_
            nw101_[int(6)] = ((d_630_v23_)[(self).f37] if ((self).f37) in (d_630_v23_) else d_611_v4_)
            nw101_[int(7)] = (d_611_v4_) + (d_611_v4_)
            nw101_[int(8)] = d_611_v4_
            nw101_[int(9)] = d_611_v4_
            nw101_[int(10)] = _dafny.SeqWithoutIsStrInference([(d_611_v4_)[default__.safeIndex((0) - ((self).f25), len(d_611_v4_))] for d_633_i2_ in range(default__.abs(603))])
            nw101_[int(11)] = d_611_v4_
            d_632_v24_ = nw101_
            d_634_v25_: _dafny.Map
            d_634_v25_ = _dafny.Map({(self).f25: len(d_616_v9_)})
            index94_ = default__.safeIndex(355, (d_606_v1_).length(0))
            rhs97_ = default__.safeDivisionInt((self).f37, (self).f25)
            rhs98_ = d_632_v24_
            rhs99_ = (d_606_v1_)[default__.safeIndex(355, (d_606_v1_).length(0))]
            rhs100_ = default__.safeDivisionInt((len(d_619_v12_) if default__.fm0((self).f37, globalState) else (d_610_v3_).fm26(1, (d_606_v1_)[default__.safeIndex(355, (d_606_v1_).length(0))], globalState)), ((d_634_v25_)[(self).f25] if ((self).f25) in (d_634_v25_) else (self).f37))
            lhs67_ = d_606_v1_
            lhs68_ = default__.safeIndex(355, (d_606_v1_).length(0))
            r1 = rhs97_
            d_632_v24_ = rhs98_
            lhs67_[lhs68_] = rhs99_
            r1 = rhs100_
        elif True:
            d_635_v26_: D3
            d_635_v26_ = D3_DC10((self).f38, d_606_v1_, p1, (self).f25, len(_dafny.Set({_dafny.CodePoint('o'), default__.fm39(p1, globalState)})))
            d_636_v27_: _dafny.Map
            d_636_v27_ = _dafny.Map({(self).f38: (d_635_v26_).cf14})
            d_636_v27_ = d_636_v27_
            d_637_v28_: _dafny.Seq
            d_637_v28_ = _dafny.SeqWithoutIsStrInference([True])
            r0 = len((d_637_v28_) + ((d_637_v28_).set(default__.safeIndex((self).f37, len(d_637_v28_)), p1)))
            (globalState).f10 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjn"))
            d_638_v29_: D7
            d_638_v29_ = D7_DC21(p3)
            d_639_v30_: _dafny.Map
            d_639_v30_ = _dafny.Map({d_606_v1_: d_638_v29_})
            d_639_v30_ = d_639_v30_
            r1 = (self).f25
        d_640_v31_: _dafny.Map
        d_640_v31_ = _dafny.Map({d_606_v1_: p1})
        d_641_v32_: D8
        d_641_v32_ = D8_DC23(d_640_v31_)
        if ((d_606_v1_ if True else d_606_v1_)) in ((d_641_v32_).cf35):
            (globalState).f12 = (p2) or (p1)
            d_642_v33_: _dafny.Seq
            d_642_v33_ = _dafny.SeqWithoutIsStrInference([p1])
            d_643_v34_: D1
            d_643_v34_ = D1_DC4(d_642_v33_)
            d_644_v35_: D1
            d_644_v35_ = D1_DC7(d_643_v34_)
            pat_let_tv12_ = d_643_v34_
            def iife47_(_pat_let6_0):
                def iife48_(d_645_dt__update__tmp_h0_):
                    def iife49_(_pat_let7_0):
                        def iife50_(d_646_dt__update_hcf10_h0_):
                            return D1_DC7(d_646_dt__update_hcf10_h0_)
                        return iife50_(_pat_let7_0)
                    return iife49_(pat_let_tv12_)
                return iife48_(_pat_let6_0)
            source13_ = iife47_(d_644_v35_)
            if source13_.is_DC5:
                d_647_v36_: _dafny.Seq
                d_647_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlmwcmkr"))
                index95_ = default__.safeIndex(802, ((self).f26).length(0))
                ((self).f26)[index95_] = d_647_v36_
                index96_ = default__.safeIndex(802, ((self).f26).length(0))
                ((self).f26)[index96_] = d_647_v36_
                d_648_v37_: _dafny.Map
                d_648_v37_ = _dafny.Map({p2: (self).f25})
                d_649_v38_: _dafny.Set
                d_649_v38_ = _dafny.Set({d_648_v37_, d_648_v37_})
                d_650_v39_: D9
                d_650_v39_ = D9_DC26(d_648_v37_)
                d_651_v40_: str
                d_651_v40_ = _dafny.CodePoint('l')
                d_652_v41_: _dafny.Map
                d_652_v41_ = _dafny.Map({len(_dafny.Set({(self).f37})): d_651_v40_})
                d_653_v42_: _dafny.Map
                d_653_v42_ = _dafny.Map({(d_650_v39_).cf38: d_652_v41_})
                d_654_v44_: _dafny.Map
                d_654_v44_ = _dafny.Map({default__.fm41(globalState): (self).f37})
                d_655_v46_: _dafny.Array
                nw102_ = _dafny.Array(None, 21)
                nw102_[int(0)] = d_649_v38_
                nw102_[int(1)] = default__.fm40(globalState)
                nw102_[int(2)] = (d_649_v38_) - (d_649_v38_)
                nw102_[int(3)] = d_649_v38_
                nw102_[int(4)] = d_649_v38_
                nw102_[int(5)] = (d_649_v38_ if (self).f38 else d_649_v38_)
                nw102_[int(6)] = d_649_v38_
                nw102_[int(7)] = _dafny.Set({d_648_v37_})
                nw102_[int(8)] = d_649_v38_
                nw102_[int(9)] = d_649_v38_
                nw102_[int(10)] = d_649_v38_
                nw102_[int(11)] = _dafny.Set({(d_648_v37_).set(p1, (self).f37), default__.fm41(globalState)})
                def iife51_():
                    coll35_ = _dafny.Set()
                    compr_35_: _dafny.Map
                    for compr_35_ in (d_653_v42_).keys.Elements:
                        d_656_v43_: _dafny.Map = compr_35_
                        if (d_656_v43_) in (d_653_v42_):
                            coll35_ = coll35_.union(_dafny.Set([d_656_v43_]))
                    return _dafny.Set(coll35_)
                nw102_[int(12)] = iife51_()
                
                nw102_[int(13)] = d_649_v38_
                nw102_[int(14)] = (d_649_v38_).intersection(d_649_v38_)
                nw102_[int(15)] = d_649_v38_
                nw102_[int(16)] = d_649_v38_
                def iife52_():
                    coll36_ = _dafny.Set()
                    compr_36_: _dafny.Map
                    for compr_36_ in (d_654_v44_).keys.Elements:
                        d_657_v45_: _dafny.Map = compr_36_
                        if (d_657_v45_) in (d_654_v44_):
                            coll36_ = coll36_.union(_dafny.Set([d_657_v45_]))
                    return _dafny.Set(coll36_)
                nw102_[int(17)] = (d_649_v38_) | (iife52_()
                )
                nw102_[int(18)] = _dafny.Set({default__.fm41(globalState), d_648_v37_})
                nw102_[int(19)] = (default__.fm40(globalState) if (self).f38 else d_649_v38_)
                nw102_[int(20)] = d_649_v38_
                d_655_v46_ = nw102_
                index97_ = default__.safeIndex(196, (d_655_v46_).length(0))
                (d_655_v46_)[index97_] = d_649_v38_
                d_658_v47_: _dafny.Set
                d_658_v47_ = _dafny.Set({default__.fm0(len(d_648_v37_), globalState)})
                d_659_v48_: D4
                d_659_v48_ = D4_DC13((0) - ((self).f37), (self).f38, (d_658_v47_).isdisjoint(d_658_v47_), (self).f38)
                d_660_v49_: _dafny.MultiSet
                d_660_v49_ = _dafny.MultiSet([(self).f25])
                index98_ = default__.safeIndex(196, (d_655_v46_).length(0))
                rhs101_ = (((d_660_v49_).cardinality) * ((0) - ((self).f37)) if (not((d_642_v33_)[default__.safeIndex((self).f37, len(d_642_v33_))])) and ((self).f38) else (self).f25)
                rhs102_ = p1
                rhs103_ = d_649_v38_
                rhs104_ = not((self).f38)
                rhs105_ = d_659_v48_
                lhs69_ = globalState
                lhs70_ = d_655_v46_
                lhs71_ = default__.safeIndex(196, (d_655_v46_).length(0))
                lhs72_ = globalState
                r1 = rhs101_
                lhs69_.f12 = rhs102_
                lhs70_[lhs71_] = rhs103_
                lhs72_.f12 = rhs104_
                d_659_v48_ = rhs105_
                (globalState).f10 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xehpc"))) + ((d_647_v36_) + (((self).f26)[default__.safeIndex(802, ((self).f26).length(0))]))
                index99_ = default__.safeIndex(587, (p3).length(0))
                (p3)[index99_] = (self).f25
                d_661_v50_: _dafny.Map
                d_661_v50_ = _dafny.Map({877: len(d_642_v33_)})
                index100_ = default__.safeIndex(587, (p3).length(0))
                (p3)[index100_] = (0) - (((d_610_v3_).fm26(len(d_647_v36_), not(True), globalState)) - ((0) - ((d_610_v3_).fm26((0) - ((self).f37), not((d_642_v33_)[default__.safeIndex(((d_661_v50_)[107] if (107) in (d_661_v50_) else (self).f25), len(d_642_v33_))]), globalState))))
            elif source13_.is_DC6:
                d_662___mcc_h0_ = source13_.cf7
                d_663___mcc_h1_ = source13_.cf8
                d_664___mcc_h2_ = source13_.cf9
                d_665_cf9_ = d_664___mcc_h2_
                d_666_cf8_ = d_663___mcc_h1_
                d_667_cf7_ = d_662___mcc_h0_
                d_668_v51_: D5
                d_668_v51_ = D5_DC16(not(p1))
                d_668_v51_ = d_668_v51_
                r0 = ((self).f25) + ((self).f37)
                (globalState).f10 = d_665_cf9_
                d_669_v52_: _dafny.Set
                d_669_v52_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_667_cf7_ for d_670_i3_ in range(default__.abs(-429))])), (self).f25})
                d_671_v53_: _dafny.Set
                d_671_v53_ = _dafny.Set({(self).f38})
                d_672_v54_: D3
                d_672_v54_ = D3_DC10((_dafny.Set({(self).f37, (0) - (len(d_671_v53_))})).issubset(d_669_v52_), d_606_v1_, default__.fm0(len(d_671_v53_), globalState), (795) - ((self).f25), (0) - ((0) - ((self).f37)))
                rhs106_ = d_672_v54_
                rhs107_ = p1
                d_672_v54_ = rhs106_
                d_666_cf8_ = rhs107_
            elif source13_.is_DC4:
                d_673___mcc_h3_ = source13_.cf6
                d_674_cf6_ = d_673___mcc_h3_
                (globalState).f12 = ((self).f37) < ((self).f37)
                d_675_v55_: C0
                nw103_ = C0()
                nw103_.ctor__()
                d_675_v55_ = nw103_
                (globalState).f12 = p2
                r0 = (0) - (((self).f37) + ((self).f37))
            elif True:
                d_676___mcc_h4_ = source13_.cf10
                d_677_cf10_ = d_676___mcc_h4_
                d_678_v56_: _dafny.Seq
                d_678_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egqfhrepj"))
                d_679_v58_: _dafny.Map
                d_679_v58_ = _dafny.Map({(self).f37: (0) - ((self).f37)})
                d_680_v59_: _dafny.MultiSet
                def iife53_():
                    coll37_ = _dafny.Map()
                    compr_37_: int
                    for compr_37_ in _dafny.IntegerRange(134, 936):
                        d_681_v57_: int = compr_37_
                        if ((134) <= (d_681_v57_)) and ((d_681_v57_) < (936)):
                            coll37_[(d_681_v57_) + ((self).f25)] = (self).f25
                    return _dafny.Map(coll37_)
                d_680_v59_ = _dafny.MultiSet([_dafny.Map({len(d_678_v56_): (self).f37}), (iife53_()
                ) | (d_679_v58_)])
                d_682_v60_: _dafny.Array
                def lambda32_(d_683_p2_):
                    def lambda33_(d_684_i4_):
                        return D0_DC3((0) - (len(_dafny.SeqWithoutIsStrInference([(self).f37]))), d_683_p2_)

                    return lambda33_

                init18_ = lambda32_(p2)
                nw104_ = _dafny.Array(None, 12)
                for i0_18_ in range(nw104_.length(0)):
                    nw104_[i0_18_] = init18_(i0_18_)
                d_682_v60_ = nw104_
                d_685_v61_: _dafny.Set
                d_685_v61_ = _dafny.Set({not(p1), (self).f38, True, p1})
                index101_ = default__.safeIndex(555, (d_682_v60_).length(0))
                (d_682_v60_)[index101_] = default__.fm42(len(d_685_v61_), globalState)
                index102_ = default__.safeIndex(43, (d_606_v1_).length(0))
                (d_606_v1_)[index102_] = False
                d_686_v62_: D0
                d_686_v62_ = D0_DC3((self).f25, p1)
                d_687_v63_: D6
                d_687_v63_ = D6_DC19(d_686_v62_, (self).f25)
                d_688_v64_: str
                d_688_v64_ = _dafny.CodePoint('m')
                pat_let_tv13_ = d_686_v62_
                index103_ = default__.safeIndex(555, (d_682_v60_).length(0))
                index104_ = default__.safeIndex(43, (d_606_v1_).length(0))
                def iife54_(_pat_let8_0):
                    def iife55_(d_689_dt__update__tmp_h1_):
                        def iife56_(_pat_let9_0):
                            def iife57_(d_690_dt__update_hcf29_h0_):
                                return D6_DC19(d_690_dt__update_hcf29_h0_, (d_689_dt__update__tmp_h1_).cf30)
                            return iife57_(_pat_let9_0)
                        return iife56_(pat_let_tv13_)
                    return iife55_(_pat_let8_0)
                rhs108_ = d_680_v59_
                rhs109_ = d_686_v62_
                rhs110_ = not ((self).f38) or ((not((self).f38)) in (d_642_v33_))
                rhs111_ = iife54_(d_687_v63_)
                rhs112_ = d_688_v64_
                lhs73_ = d_682_v60_
                lhs74_ = default__.safeIndex(555, (d_682_v60_).length(0))
                lhs75_ = d_606_v1_
                lhs76_ = default__.safeIndex(43, (d_606_v1_).length(0))
                lhs77_ = globalState
                d_680_v59_ = rhs108_
                lhs73_[lhs74_] = rhs109_
                lhs75_[lhs76_] = rhs110_
                d_687_v63_ = rhs111_
                lhs77_.f23 = rhs112_
                d_691_v65_: _dafny.Array
                def lambda34_(d_692_v33_):
                    def lambda35_(d_693_i5_):
                        return d_692_v33_

                    return lambda35_

                init19_ = lambda34_(d_642_v33_)
                nw105_ = _dafny.Array(None, 18)
                for i0_19_ in range(nw105_.length(0)):
                    nw105_[i0_19_] = init19_(i0_19_)
                d_691_v65_ = nw105_
                d_694_v66_: _dafny.Seq
                d_694_v66_ = _dafny.SeqWithoutIsStrInference([d_642_v33_, d_642_v33_])
                index105_ = default__.safeIndex(485, (d_691_v65_).length(0))
                (d_691_v65_)[index105_] = (_dafny.SeqWithoutIsStrInference([p2])) + ((d_694_v66_)[default__.safeIndex((self).f37, len(d_694_v66_))])
                d_695_v67_: _dafny.Map
                d_695_v67_ = _dafny.Map({d_688_v64_: (d_606_v1_)[default__.safeIndex(43, (d_606_v1_).length(0))]})
                d_696_v69_: _dafny.Array
                nw106_ = _dafny.Array(None, 13)
                nw106_[int(0)] = (d_695_v67_) | (d_695_v67_)
                nw106_[int(1)] = (d_695_v67_) | ((d_695_v67_).set(_dafny.CodePoint('u'), p1))
                nw106_[int(2)] = d_695_v67_
                nw106_[int(3)] = d_695_v67_
                nw106_[int(4)] = d_695_v67_
                nw106_[int(5)] = d_695_v67_
                nw106_[int(6)] = d_695_v67_
                nw106_[int(7)] = d_695_v67_
                nw106_[int(8)] = d_695_v67_
                def iife58_():
                    coll38_ = _dafny.Map()
                    compr_38_: str
                    for compr_38_ in (d_678_v56_).Elements:
                        d_697_v68_: str = compr_38_
                        if (d_697_v68_) in (d_678_v56_):
                            coll38_[d_697_v68_] = False
                    return _dafny.Map(coll38_)
                nw106_[int(9)] = iife58_()
                
                nw106_[int(10)] = _dafny.Map({d_688_v64_: p2})
                nw106_[int(11)] = d_695_v67_
                nw106_[int(12)] = d_695_v67_
                d_696_v69_ = nw106_
                d_698_v70_: D10
                d_698_v70_ = D10_DC28(d_695_v67_)
                d_699_v71_: _dafny.Seq
                d_699_v71_ = _dafny.SeqWithoutIsStrInference([d_695_v67_, d_695_v67_, (d_698_v70_).cf41])
                index106_ = default__.safeIndex(869, (d_696_v69_).length(0))
                (d_696_v69_)[index106_] = ((d_699_v71_)[default__.safeIndex((self).f37, len(d_699_v71_))]) | (_dafny.Map({d_688_v64_: (self).f38}))
                index107_ = default__.safeIndex(485, (d_691_v65_).length(0))
                index108_ = default__.safeIndex(43, (d_606_v1_).length(0))
                index109_ = default__.safeIndex(43, (d_606_v1_).length(0))
                index110_ = default__.safeIndex(869, (d_696_v69_).length(0))
                rhs113_ = d_642_v33_
                rhs114_ = ((self).f37) == (((self).f37 if p1 else (self).f25))
                rhs115_ = p1
                rhs116_ = (_dafny.Map({d_688_v64_: default__.fm0((self).f25, globalState)})) | (d_695_v67_)
                lhs78_ = d_691_v65_
                lhs79_ = default__.safeIndex(485, (d_691_v65_).length(0))
                lhs80_ = d_606_v1_
                lhs81_ = default__.safeIndex(43, (d_606_v1_).length(0))
                lhs82_ = d_606_v1_
                lhs83_ = default__.safeIndex(43, (d_606_v1_).length(0))
                lhs84_ = d_696_v69_
                lhs85_ = default__.safeIndex(869, (d_696_v69_).length(0))
                lhs78_[lhs79_] = rhs113_
                lhs80_[lhs81_] = rhs114_
                lhs82_[lhs83_] = rhs115_
                lhs84_[lhs85_] = rhs116_
                d_700_v72_: _dafny.Seq
                d_700_v72_ = _dafny.SeqWithoutIsStrInference([(self).f37])
                (globalState).f6 = (d_700_v72_) + (d_700_v72_)
                index111_ = default__.safeIndex(43, (d_606_v1_).length(0))
                (d_606_v1_)[index111_] = (d_642_v33_)[default__.safeIndex((0) - ((self).f25), len(d_642_v33_))]
            d_701_v73_: _dafny.Array
            nw107_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_701_v73_ = nw107_
            d_702_v74_: _dafny.Array
            nw108_ = _dafny.Array(D7.default()(), 5)
            d_702_v74_ = nw108_
            index112_ = default__.safeIndex(722, (d_701_v73_).length(0))
            (d_701_v73_)[index112_] = d_702_v74_
            index113_ = default__.safeIndex(722, (d_701_v73_).length(0))
            (d_701_v73_)[index113_] = d_702_v74_
            d_703_v75_: _dafny.Seq
            d_703_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkk"))
            d_704_v76_: _dafny.Seq
            d_704_v76_ = _dafny.SeqWithoutIsStrInference([len(d_703_v75_), -949, (self).f25, (self).f37])
            d_705_v77_: _dafny.Seq
            d_705_v77_ = d_704_v76_
            source14_ = d_705_v77_
            d_706___mcc_h5_ = source14_
            d_707_cf11_ = d_706___mcc_h5_
            (globalState).f12 = p2
            d_708_v78_: _dafny.Map
            d_708_v78_ = _dafny.Map({(self).f38: True})
            d_709_v79_: _dafny.Map
            d_709_v79_ = _dafny.Map({(self).f38: p2})
            d_710_v80_: _dafny.Seq
            d_710_v80_ = _dafny.SeqWithoutIsStrInference([(d_708_v78_), (d_709_v79_).set(p2, True), d_709_v79_])
            d_710_v80_ = ((d_710_v80_) + (d_710_v80_)) + (_dafny.SeqWithoutIsStrInference([d_709_v79_]))
            (globalState).f23 = default__.fm39((self).f38, globalState)
            (globalState).f12 = p2
            d_711_v82_: _dafny.Map
            d_711_v82_ = _dafny.Map({_dafny.CodePoint('a'): True})
            def iife59_():
                coll39_ = _dafny.Map()
                compr_39_: str
                for compr_39_ in (d_711_v82_).keys.Elements:
                    d_712_v81_: str = compr_39_
                    if (d_712_v81_) in (d_711_v82_):
                        coll39_[d_712_v81_] = p1
                return _dafny.Map(coll39_)
            (globalState).f12 = default__.fm0(len(iife59_()
            ), globalState)
        elif True:
            d_713_v83_: _dafny.Seq
            d_713_v83_ = _dafny.SeqWithoutIsStrInference([p3])
            (globalState).f12 = (d_713_v83_) <= (d_713_v83_)
            d_714_v84_: str
            d_714_v84_ = _dafny.CodePoint('g')
            d_715_v85_: _dafny.Map
            d_715_v85_ = _dafny.Map({p1: d_714_v84_})
            d_716_v86_: _dafny.Map
            d_716_v86_ = _dafny.Map({d_715_v85_: p2})
            d_717_v87_: _dafny.Seq
            d_717_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bh"))
            d_718_v88_: _dafny.Seq
            d_718_v88_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f37, (self).f37, (self).f25])), (self).f37])
            d_719_v89_: D7
            d_719_v89_ = D7_DC21(p3)
            d_720_v90_: _dafny.Map
            d_720_v90_ = _dafny.Map({d_719_v89_: (self).f37})
            d_721_v91_: D3
            d_721_v91_ = D3_DC10((self).f38, d_606_v1_, default__.fm0(len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f25)])), globalState), 939, -840)
            d_722_v92_: _dafny.Map
            d_722_v92_ = _dafny.Map({True: p2})
            d_723_v93_: _dafny.Map
            d_723_v93_ = _dafny.Map({p2: (d_718_v88_)[default__.safeIndex(len(d_717_v87_), len(d_718_v88_))]})
            d_724_v94_: _dafny.Seq
            d_724_v94_ = _dafny.SeqWithoutIsStrInference([p2, False])
            d_725_v95_: _dafny.Set
            d_725_v95_ = _dafny.Set({p1})
            d_726_v96_: _dafny.Array
            nw109_ = _dafny.Array(None, 27)
            nw109_[int(0)] = (self).f37
            nw109_[int(1)] = (self).f25
            nw109_[int(2)] = default__.safeDivisionInt((self).f25, -814)
            nw109_[int(3)] = -329
            nw109_[int(4)] = len((d_716_v86_) | (d_716_v86_))
            nw109_[int(5)] = (0) - ((len(d_717_v87_)) * ((self).f25))
            nw109_[int(6)] = ((self).f25) * ((self).f37)
            nw109_[int(7)] = default__.safeDivisionInt((self).f25, 110)
            nw109_[int(8)] = (self).f25
            nw109_[int(9)] = (d_610_v3_).fm26((self).f37, p1, globalState)
            nw109_[int(10)] = (self).f25
            nw109_[int(11)] = (d_718_v88_)[default__.safeIndex(((d_720_v90_)[d_719_v89_] if (d_719_v89_) in (d_720_v90_) else (self).f25), len(d_718_v88_))]
            nw109_[int(12)] = default__.safeDivisionInt((0) - ((0) - ((d_721_v91_).cf17)), (self).f25)
            nw109_[int(13)] = 662
            nw109_[int(14)] = len(d_717_v87_)
            nw109_[int(15)] = (self).f25
            nw109_[int(16)] = (self).f25
            nw109_[int(17)] = (self).f25
            nw109_[int(18)] = (self).f25
            nw109_[int(19)] = (0) - (default__.safeModuloInt((self).f25, (self).f25))
            nw109_[int(20)] = 40
            nw109_[int(21)] = len(d_722_v92_)
            nw109_[int(22)] = (d_610_v3_).fm26((self).f25, (self).f38, globalState)
            nw109_[int(23)] = (self).f37
            nw109_[int(24)] = ((d_610_v3_).fm26((self).f25, p1, globalState)) * ((self).f25)
            nw109_[int(25)] = ((d_723_v93_)[p2] if (p2) in (d_723_v93_) else len(d_724_v94_))
            nw109_[int(26)] = len(d_725_v95_)
            d_726_v96_ = nw109_
            nw110_ = _dafny.Array(None, 10)
            nw110_[int(0)] = (self).f37
            nw110_[int(1)] = -890
            nw110_[int(2)] = (self).f37
            nw110_[int(3)] = (self).f37
            nw110_[int(4)] = (d_610_v3_).fm26((self).f37, True, globalState)
            nw110_[int(5)] = (self).f25
            nw110_[int(6)] = (self).f37
            nw110_[int(7)] = (self).f25
            nw110_[int(8)] = (len(d_718_v88_)) - ((self).f25)
            nw110_[int(9)] = (self).f37
            d_726_v96_ = nw110_
            d_727_v97_: _dafny.Map
            d_727_v97_ = _dafny.Map({(self).f37: (self).f38})
            d_728_v98_: _dafny.Map
            d_728_v98_ = _dafny.Map({(self).f25: d_714_v84_})
            rhs117_ = ((d_727_v97_)[(self).f25] if ((self).f25) in (d_727_v97_) else (len(d_717_v87_)) < (len(d_728_v98_)))
            rhs118_ = (self).f38
            lhs86_ = globalState
            lhs87_ = globalState
            lhs86_.f12 = rhs117_
            lhs87_.f12 = rhs118_
            index114_ = default__.safeIndex(613, (p3).length(0))
            (p3)[index114_] = len(d_725_v95_)
            index115_ = default__.safeIndex(613, (p3).length(0))
            rhs119_ = len((d_717_v87_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxowgeyge"))))
            rhs120_ = p2
            rhs121_ = (self).f25
            lhs88_ = p3
            lhs89_ = default__.safeIndex(613, (p3).length(0))
            lhs90_ = globalState
            lhs88_[lhs89_] = rhs119_
            lhs90_.f12 = rhs120_
            r0 = rhs121_
            (globalState).f23 = d_714_v84_
        d_729_v99_: str
        d_729_v99_ = _dafny.CodePoint('l')
        d_730_v100_: _dafny.Seq
        d_730_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymxtrxics"))
        d_731_v101_: _dafny.Set
        d_731_v101_ = _dafny.Set({p2})
        (globalState).f12 = not ((d_729_v99_) in (d_730_v100_)) or ((d_731_v101_).ispropersubset(d_731_v101_))
        r0 = (0) - (len(default__.fm1(False, globalState)))
        r1 = -711
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        d_732_v0_: _dafny.Seq
        d_732_v0_ = _dafny.SeqWithoutIsStrInference([(self).f37])
        d_733_v1_: _dafny.Map
        d_733_v1_ = _dafny.Map({(self).f25: p1})
        d_734_v2_: _dafny.Seq
        d_734_v2_ = _dafny.SeqWithoutIsStrInference([len(d_732_v0_), (self).f25, len(d_733_v1_), (self).f37])
        d_735_v3_: _dafny.Seq
        d_735_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olbsqy"))
        d_736_v4_: str
        d_736_v4_ = _dafny.CodePoint('b')
        d_737_v5_: _dafny.Map
        d_737_v5_ = _dafny.Map({p1: (self).f25})
        d_738_v6_: _dafny.Seq
        d_738_v6_ = _dafny.SeqWithoutIsStrInference([p1])
        d_739_v7_: _dafny.Array
        nw111_ = _dafny.Array(None, 20)
        nw111_[int(0)] = len(d_735_v3_)
        nw111_[int(1)] = (self).f37
        nw111_[int(2)] = 801
        nw111_[int(3)] = (self).f25
        nw111_[int(4)] = (self).f37
        nw111_[int(5)] = 355
        nw111_[int(6)] = (self).f37
        nw111_[int(7)] = (self).f37
        nw111_[int(8)] = (self).f25
        nw111_[int(9)] = (self).f25
        nw111_[int(10)] = (self).f25
        nw111_[int(11)] = (self).f25
        nw111_[int(12)] = (self).f25
        nw111_[int(13)] = (self).f37
        nw111_[int(14)] = (self).f25
        nw111_[int(15)] = (self).f25
        nw111_[int(16)] = (0) - ((d_732_v0_)[default__.safeIndex(len(((d_735_v3_).set(default__.safeIndex((self).f25, len(d_735_v3_)), d_736_v4_)).set(default__.safeIndex(len(d_737_v5_), len((d_735_v3_).set(default__.safeIndex((self).f25, len(d_735_v3_)), d_736_v4_))), d_736_v4_)), len(d_732_v0_))])
        nw111_[int(17)] = len(d_738_v6_)
        nw111_[int(18)] = (self).f25
        nw111_[int(19)] = (self).f25
        d_739_v7_ = nw111_
        d_740_v8_: _dafny.Map
        d_740_v8_ = _dafny.Map({False: d_739_v7_})
        d_741_v9_: _dafny.Set
        d_741_v9_ = _dafny.Set({not((self).f38), p1, p1, (self).f38, (self).f38})
        d_742_v10_: _dafny.MultiSet
        d_742_v10_ = _dafny.MultiSet([default__.fm2((self).f25, (self).f25, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))])
        d_743_v11_: _dafny.MultiSet
        d_743_v11_ = _dafny.MultiSet([False])
        d_744_v12_: _dafny.Seq
        d_744_v12_ = _dafny.SeqWithoutIsStrInference([d_743_v11_, d_743_v11_, d_743_v11_, _dafny.MultiSet(d_738_v6_), d_743_v11_])
        d_745_v13_: _dafny.Map
        d_745_v13_ = _dafny.Map({(self).f38: d_744_v12_})
        d_746_v14_: _dafny.Array
        nw112_ = _dafny.Array(None, 20)
        nw112_[int(0)] = (0) - ((self).f37)
        nw112_[int(1)] = (self).f37
        nw112_[int(2)] = len((d_732_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f37, (self).f25])))
        nw112_[int(3)] = (self).f37
        nw112_[int(4)] = (666 if (self).f38 else (self).f37)
        nw112_[int(5)] = ((self).f25 if p1 else (self).f25)
        nw112_[int(6)] = ((self).f25) * (len(d_740_v8_))
        nw112_[int(7)] = len(d_741_v9_)
        nw112_[int(8)] = len(d_733_v1_)
        nw112_[int(9)] = (self).f37
        nw112_[int(10)] = (self).f25
        nw112_[int(11)] = (685) + ((self).f37)
        nw112_[int(12)] = (self).f25
        nw112_[int(13)] = default__.safeDivisionInt(616, -96)
        nw112_[int(14)] = ((self).f37) - ((0) - ((d_742_v10_).cardinality))
        nw112_[int(15)] = ((self).f37) - ((self).f25)
        nw112_[int(16)] = (979) + ((self).f25)
        nw112_[int(17)] = (self).f37
        nw112_[int(18)] = (self).f25
        nw112_[int(19)] = len((((d_745_v13_)[(self).f38] if ((self).f38) in (d_745_v13_) else d_744_v12_)) + (_dafny.SeqWithoutIsStrInference([d_743_v11_, d_743_v11_])))
        d_746_v14_ = nw112_
        index116_ = default__.safeIndex(314, (d_739_v7_).length(0))
        (d_739_v7_)[index116_] = 764
        index117_ = default__.safeIndex(314, (d_739_v7_).length(0))
        (d_739_v7_)[index117_] = 919
        index118_ = default__.safeIndex(314, (d_739_v7_).length(0))
        rhs122_ = len((d_735_v3_) + ((_dafny.SeqWithoutIsStrInference([d_736_v4_ for d_747_i0_ in range(default__.abs(42))])) + (d_735_v3_)))
        rhs123_ = (self).f37
        lhs91_ = d_739_v7_
        lhs92_ = default__.safeIndex(314, (d_739_v7_).length(0))
        r0 = rhs122_
        lhs91_[lhs92_] = rhs123_
        if default__.fm0(default__.safeModuloInt((self).f37, (self).f25), globalState):
            source15_ = d_734_v2_
            d_748___mcc_h0_ = source15_
            d_749_cf11_ = d_748___mcc_h0_
            d_750_v15_: _dafny.Map
            d_750_v15_ = _dafny.Map({(self).f38: p1})
            d_751_v16_: _dafny.Map
            d_751_v16_ = _dafny.Map({(self).f38: d_750_v15_})
            d_752_v17_: _dafny.Map
            d_752_v17_ = _dafny.Map({d_739_v7_: _dafny.Map({(self).f38: (self).f38})})
            d_753_v18_: _dafny.Map
            d_753_v18_ = _dafny.Map({default__.fm44(p1, (self).f38, p1, len(d_735_v3_), globalState): d_750_v15_})
            d_754_v19_: _dafny.Seq
            d_754_v19_ = _dafny.SeqWithoutIsStrInference([d_750_v15_])
            d_755_v20_: _dafny.Array
            nw113_ = _dafny.Array(None, 26)
            nw113_[int(0)] = d_750_v15_
            nw113_[int(1)] = ((d_751_v16_)[p1] if (p1) in (d_751_v16_) else d_750_v15_)
            nw113_[int(2)] = d_750_v15_
            nw113_[int(3)] = d_750_v15_
            nw113_[int(4)] = d_750_v15_
            nw113_[int(5)] = d_750_v15_
            nw113_[int(6)] = _dafny.Map({p1: p1})
            nw113_[int(7)] = ((d_752_v17_)[d_746_v14_] if (d_746_v14_) in (d_752_v17_) else (_dafny.Map({p1: p1})))
            nw113_[int(8)] = d_750_v15_
            nw113_[int(9)] = d_750_v15_
            nw113_[int(10)] = _dafny.Map({p1: (self).f38})
            nw113_[int(11)] = d_750_v15_
            nw113_[int(12)] = (d_750_v15_).set((d_738_v6_)[default__.safeIndex(-969, len(d_738_v6_))], p1)
            nw113_[int(13)] = _dafny.Map({p1: (self).f38})
            nw113_[int(14)] = (d_750_v15_) | (d_750_v15_)
            nw113_[int(15)] = (d_750_v15_) | (d_750_v15_)
            nw113_[int(16)] = (default__.fm43((self).f38, (self).f38, globalState)) | ((d_750_v15_).set((self).f38, (self).f38))
            nw113_[int(17)] = d_750_v15_
            nw113_[int(18)] = d_750_v15_
            nw113_[int(19)] = d_750_v15_
            nw113_[int(20)] = _dafny.Map({(self).f38: p1})
            nw113_[int(21)] = _dafny.Map({not((self).f38): p1})
            nw113_[int(22)] = (((d_753_v18_)[default__.fm44(p1, (self).f38, p1, (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))], globalState)] if (default__.fm44(p1, (self).f38, p1, (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))], globalState)) in (d_753_v18_) else _dafny.Map({((d_733_v1_)[(self).f37] if ((self).f37) in (d_733_v1_) else not(p1)): (self).f38}))) | ((d_750_v15_).set(p1, (self).f38))
            nw113_[int(23)] = _dafny.Map({default__.fm0((self).f37, globalState): (self).f38})
            nw113_[int(24)] = (d_754_v19_)[default__.safeIndex((self).f25, len(d_754_v19_))]
            nw113_[int(25)] = d_750_v15_
            d_755_v20_ = nw113_
            index119_ = default__.safeIndex(722, (d_755_v20_).length(0))
            (d_755_v20_)[index119_] = (d_750_v15_) | (d_750_v15_)
            d_756_v21_: _dafny.Array
            def lambda36_(d_757_p1_):
                def lambda37_(d_758_i1_):
                    return d_757_p1_

                return lambda37_

            init20_ = lambda36_(p1)
            nw114_ = _dafny.Array(None, 10)
            for i0_20_ in range(nw114_.length(0)):
                nw114_[i0_20_] = init20_(i0_20_)
            d_756_v21_ = nw114_
            d_759_v22_: D3
            d_759_v22_ = D3_DC10((self).f38, d_756_v21_, p1, 43, (self).f37)
            d_760_v23_: _dafny.Map
            d_760_v23_ = _dafny.Map({d_756_v21_: D3_DC10(p1, d_756_v21_, (d_759_v22_).cf13, len(default__.fm1((self).f38, globalState)), default__.fm44(not(False), p1, (self).f38, (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))], globalState))})
            index120_ = default__.safeIndex(722, (d_755_v20_).length(0))
            index121_ = default__.safeIndex(314, (d_739_v7_).length(0))
            rhs124_ = d_750_v15_
            rhs125_ = len(d_760_v23_)
            lhs93_ = d_755_v20_
            lhs94_ = default__.safeIndex(722, (d_755_v20_).length(0))
            lhs95_ = d_739_v7_
            lhs96_ = default__.safeIndex(314, (d_739_v7_).length(0))
            lhs93_[lhs94_] = rhs124_
            lhs95_[lhs96_] = rhs125_
            d_761_v24_: _dafny.MultiSet
            d_761_v24_ = _dafny.MultiSet([(d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))]])
            (globalState).f12 = (-576) in (d_761_v24_)
            d_733_v1_ = d_733_v1_
            d_762_v25_: C2
            nw115_ = C2()
            nw115_.ctor__()
            d_762_v25_ = nw115_
            (globalState).f12 = default__.fm0((self).f25, globalState)
            d_763_v26_: _dafny.Set
            d_763_v26_ = _dafny.Set({len(d_732_v0_), len(d_735_v3_)})
            d_764_v27_: _dafny.Map
            d_764_v27_ = _dafny.Map({len((d_763_v26_) | (d_763_v26_)): (self).f25})
            d_764_v27_ = (d_764_v27_).set(541, ((self).f37) * ((_dafny.MultiSet([(self).f25])).cardinality))
            d_764_v27_ = (d_764_v27_).set((self).f25, (self).f25)
            d_765_v28_: _dafny.Map
            d_765_v28_ = _dafny.Map({d_735_v3_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnufifda"))})
            d_765_v28_ = (d_765_v28_).set((d_735_v3_).set(default__.safeIndex((0) - ((0) - ((self).f37)), len(d_735_v3_)), d_736_v4_), d_735_v3_)
        elif True:
            d_766_v29_: _dafny.Seq
            d_766_v29_ = _dafny.SeqWithoutIsStrInference([d_737_v5_, d_737_v5_, d_737_v5_])
            d_766_v29_ = d_766_v29_
            d_767_v30_: _dafny.Array
            nw116_ = _dafny.Array(D5.default()(), 9)
            d_767_v30_ = nw116_
            d_768_v31_: D5
            d_768_v31_ = D5_DC16((self).f38)
            pat_let_tv14_ = p1
            index122_ = default__.safeIndex(288, (d_767_v30_).length(0))
            def iife60_(_pat_let10_0):
                def iife61_(d_769_dt__update__tmp_h0_):
                    def iife62_(_pat_let11_0):
                        def iife63_(d_770_dt__update_hcf25_h0_):
                            return D5_DC16(d_770_dt__update_hcf25_h0_)
                        return iife63_(_pat_let11_0)
                    return iife62_(pat_let_tv14_)
                return iife61_(_pat_let10_0)
            (d_767_v30_)[index122_] = iife60_(d_768_v31_)
            index123_ = default__.safeIndex(288, (d_767_v30_).length(0))
            (d_767_v30_)[index123_] = default__.fm45(637, default__.fm44((self).f38, p1, p1, (self).f37, globalState), globalState)
            d_771_v32_: D0
            d_771_v32_ = D0_DC3((self).f37, p1)
            pat_let_tv15_ = d_739_v7_
            pat_let_tv16_ = d_739_v7_
            d_772_v33_: D6
            def iife64_(_pat_let12_0):
                def iife65_(d_773_dt__update__tmp_h1_):
                    def iife66_(_pat_let13_0):
                        def iife67_(d_774_dt__update_hcf4_h0_):
                            return D0_DC3(d_774_dt__update_hcf4_h0_, (d_773_dt__update__tmp_h1_).cf5)
                        return iife67_(_pat_let13_0)
                    return iife66_((pat_let_tv16_)[default__.safeIndex(314, (pat_let_tv15_).length(0))])
                return iife65_(_pat_let12_0)
            d_772_v33_ = D6_DC19(iife64_(d_771_v32_), (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))])
            d_775_v34_: _dafny.Array
            nw117_ = _dafny.Array(D6.default()(), 25)
            d_775_v34_ = nw117_
            pat_let_tv17_ = d_771_v32_
            d_776_v35_: _dafny.Map
            def iife68_(_pat_let14_0):
                def iife69_(d_777_dt__update__tmp_h2_):
                    def iife70_(_pat_let15_0):
                        def iife71_(d_778_dt__update_hcf29_h0_):
                            return D6_DC19(d_778_dt__update_hcf29_h0_, (d_777_dt__update__tmp_h2_).cf30)
                        return iife71_(_pat_let15_0)
                    return iife70_(pat_let_tv17_)
                return iife69_(_pat_let14_0)
            d_776_v35_ = _dafny.Map({iife68_(d_772_v33_): d_775_v34_})
            d_779_v36_: D12
            d_779_v36_ = D12_DC31(d_776_v35_)
            index124_ = default__.safeIndex(314, (d_739_v7_).length(0))
            (d_739_v7_)[index124_] = len(((d_779_v36_).cf46).set(d_772_v33_, d_775_v34_))
            (globalState).f12 = default__.fm0((0) - (default__.safeModuloInt(len(_dafny.Map({True: p1})), default__.fm44(p1, True, (self).f38, (self).f37, globalState))), globalState)
            d_735_v3_ = (d_735_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_780_i2_ in range(default__.abs(255))]))
        d_781_v37_: D9
        d_781_v37_ = D9_DC27((self).f37, (0) - (-605))
        pat_let_tv18_ = d_739_v7_
        pat_let_tv19_ = d_739_v7_
        def iife72_(_pat_let16_0):
            def iife73_(d_782_dt__update__tmp_h3_):
                def iife74_(_pat_let17_0):
                    def iife75_(d_783_dt__update_hcf40_h0_):
                        return D9_DC27((d_782_dt__update__tmp_h3_).cf39, d_783_dt__update_hcf40_h0_)
                    return iife75_(_pat_let17_0)
                return iife74_(((0) - ((self).f37)) + ((pat_let_tv19_)[default__.safeIndex(314, (pat_let_tv18_).length(0))]))
            return iife73_(_pat_let16_0)
        source16_ = iife72_(d_781_v37_)
        if source16_.is_DC27:
            d_784___mcc_h1_ = source16_.cf39
            d_785___mcc_h2_ = source16_.cf40
            d_786_cf40_ = d_785___mcc_h2_
            d_787_cf39_ = d_784___mcc_h1_
            index125_ = default__.safeIndex(314, (d_739_v7_).length(0))
            rhs126_ = (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))]
            rhs127_ = (0) - (d_787_cf39_)
            rhs128_ = (self).f38
            rhs129_ = default__.fm2((d_786_cf40_) * ((0) - ((self).f25)), (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))], globalState)
            lhs97_ = d_739_v7_
            lhs98_ = default__.safeIndex(314, (d_739_v7_).length(0))
            lhs99_ = globalState
            lhs100_ = globalState
            d_786_cf40_ = rhs126_
            lhs97_[lhs98_] = rhs127_
            lhs99_.f12 = rhs128_
            lhs100_.f10 = rhs129_
            d_788_v38_: _dafny.Array
            def lambda38_(d_789_v3_):
                def lambda39_(d_790_i3_):
                    return _dafny.Map({len(d_789_v3_): (self).f25})

                return lambda39_

            init21_ = lambda38_(d_735_v3_)
            nw118_ = _dafny.Array(None, 29)
            for i0_21_ in range(nw118_.length(0)):
                nw118_[i0_21_] = init21_(i0_21_)
            d_788_v38_ = nw118_
            d_791_v39_: _dafny.Map
            d_791_v39_ = _dafny.Map({464: (self).f25})
            index126_ = default__.safeIndex(189, (d_788_v38_).length(0))
            (d_788_v38_)[index126_] = d_791_v39_
            d_792_v40_: _dafny.Map
            d_792_v40_ = _dafny.Map({p1: d_791_v39_})
            index127_ = default__.safeIndex(189, (d_788_v38_).length(0))
            (d_788_v38_)[index127_] = (_dafny.Map({d_787_cf39_: len(d_792_v40_)})) | (d_791_v39_)
            d_793_v41_: D1
            d_793_v41_ = D1_DC6(d_736_v4_, (self).f38, d_735_v3_)
            d_794_v42_: _dafny.Map
            d_794_v42_ = default__.fm43(((d_733_v1_)[(self).f25] if ((self).f25) in (d_733_v1_) else (d_793_v41_).cf8), (self).f38, globalState)
            d_795_v43_: _dafny.Array
            nw119_ = _dafny.Array(False, 7)
            d_795_v43_ = nw119_
            d_796_v44_: _dafny.Seq
            d_796_v44_ = _dafny.SeqWithoutIsStrInference([d_795_v43_, d_795_v43_, d_795_v43_])
            d_797_v45_: _dafny.Map
            d_797_v45_ = _dafny.Map({482: d_795_v43_})
            d_798_v46_: _dafny.Map
            d_798_v46_ = _dafny.Map({d_736_v4_: d_795_v43_})
            d_799_v47_: _dafny.Array
            nw120_ = _dafny.Array(None, 29)
            nw120_[int(0)] = d_795_v43_
            nw120_[int(1)] = (d_796_v44_)[default__.safeIndex(d_787_cf39_, len(d_796_v44_))]
            nw120_[int(2)] = d_795_v43_
            nw120_[int(3)] = d_795_v43_
            nw120_[int(4)] = d_795_v43_
            nw120_[int(5)] = (((d_797_v45_)[(self).f37] if ((self).f37) in (d_797_v45_) else d_795_v43_) if not(p1) else d_795_v43_)
            nw120_[int(6)] = d_795_v43_
            nw120_[int(7)] = d_795_v43_
            nw120_[int(8)] = ((d_798_v46_)[d_736_v4_] if (d_736_v4_) in (d_798_v46_) else d_795_v43_)
            nw120_[int(9)] = d_795_v43_
            nw120_[int(10)] = d_795_v43_
            nw120_[int(11)] = d_795_v43_
            nw120_[int(12)] = d_795_v43_
            nw120_[int(13)] = d_795_v43_
            nw120_[int(14)] = d_795_v43_
            nw120_[int(15)] = d_795_v43_
            nw120_[int(16)] = d_795_v43_
            nw120_[int(17)] = d_795_v43_
            nw120_[int(18)] = d_795_v43_
            nw120_[int(19)] = d_795_v43_
            nw120_[int(20)] = d_795_v43_
            nw120_[int(21)] = d_795_v43_
            nw120_[int(22)] = d_795_v43_
            nw120_[int(23)] = d_795_v43_
            nw120_[int(24)] = d_795_v43_
            nw120_[int(25)] = (d_795_v43_ if p1 else d_795_v43_)
            nw120_[int(26)] = (d_796_v44_)[default__.safeIndex(591, len(d_796_v44_))]
            nw120_[int(27)] = ((d_797_v45_)[d_787_cf39_] if (d_787_cf39_) in (d_797_v45_) else (d_796_v44_)[default__.safeIndex((self).f37, len(d_796_v44_))])
            nw120_[int(28)] = d_795_v43_
            d_799_v47_ = nw120_
            index128_ = default__.safeIndex(780, (d_799_v47_).length(0))
            (d_799_v47_)[index128_] = d_795_v43_
            d_800_v48_: _dafny.Array
            nw121_ = _dafny.Array(_dafny.Map({}), 9)
            d_800_v48_ = nw121_
            d_801_v49_: _dafny.Seq
            d_801_v49_ = _dafny.SeqWithoutIsStrInference([d_733_v1_, d_733_v1_, d_733_v1_, d_733_v1_])
            index129_ = default__.safeIndex(71, (d_800_v48_).length(0))
            (d_800_v48_)[index129_] = (d_801_v49_)[default__.safeIndex((self).f25, len(d_801_v49_))]
            d_802_v50_: _dafny.MultiSet
            d_802_v50_ = _dafny.MultiSet([((d_740_v8_)[(self).f38] if ((self).f38) in (d_740_v8_) else d_746_v14_), d_746_v14_])
            index130_ = default__.safeIndex(780, (d_799_v47_).length(0))
            index131_ = default__.safeIndex(71, (d_800_v48_).length(0))
            rhs130_ = d_794_v42_
            rhs131_ = ((d_802_v50_)[d_746_v14_] if (d_746_v14_) in (d_802_v50_) else (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))])
            rhs132_ = default__.safeDivisionInt(904, (self).f37)
            rhs133_ = (d_795_v43_ if (d_738_v6_) <= (_dafny.SeqWithoutIsStrInference([p1])) else d_795_v43_)
            rhs134_ = (d_801_v49_)[default__.safeIndex(default__.safeDivisionInt(d_786_cf40_, (self).f37), len(d_801_v49_))]
            lhs101_ = d_799_v47_
            lhs102_ = default__.safeIndex(780, (d_799_v47_).length(0))
            lhs103_ = d_800_v48_
            lhs104_ = default__.safeIndex(71, (d_800_v48_).length(0))
            d_794_v42_ = rhs130_
            d_786_cf40_ = rhs131_
            d_786_cf40_ = rhs132_
            lhs101_[lhs102_] = rhs133_
            lhs103_[lhs104_] = rhs134_
            d_786_cf40_ = (self).f25
        elif True:
            d_803___mcc_h3_ = source16_.cf38
            d_804_cf38_ = d_803___mcc_h3_
            (globalState).f12 = p1
            if p1:
                d_805_v51_: D10
                d_805_v51_ = D10_DC29(p1, p1, ((d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))] if (self).f38 else (self).f37))
                d_805_v51_ = default__.fm46(((0) - ((self).f37) if False else (self).f37), (self).f38, not(((self).f38 if (self).f38 else False)), globalState)
                d_806_v52_: C2
                nw122_ = C2()
                nw122_.ctor__()
                d_806_v52_ = nw122_
                d_807_v53_: D0
                d_807_v53_ = D0_DC3(411, p1)
                r0 = (d_807_v53_).cf4
                d_808_v54_: _dafny.Array
                def lambda40_(d_809_i4_):
                    return _dafny.CodePoint('f')

                init22_ = lambda40_
                nw123_ = _dafny.Array(None, 11)
                for i0_22_ in range(nw123_.length(0)):
                    nw123_[i0_22_] = init22_(i0_22_)
                d_808_v54_ = nw123_
                index132_ = default__.safeIndex(941, (d_808_v54_).length(0))
                (d_808_v54_)[index132_] = default__.fm39(p1, globalState)
                d_810_v55_: _dafny.Array
                def lambda41_(d_811_v2_):
                    def lambda42_(d_812_i5_):
                        return (d_811_v2_) + (d_811_v2_)

                    return lambda42_

                init23_ = lambda41_(d_734_v2_)
                nw124_ = _dafny.Array(None, 12)
                for i0_23_ in range(nw124_.length(0)):
                    nw124_[i0_23_] = init23_(i0_23_)
                d_810_v55_ = nw124_
                index133_ = default__.safeIndex(518, (d_810_v55_).length(0))
                (d_810_v55_)[index133_] = _dafny.SeqWithoutIsStrInference([(d_806_v52_).fm24((self).f25, globalState)])
                index134_ = default__.safeIndex(941, (d_808_v54_).length(0))
                index135_ = default__.safeIndex(518, (d_810_v55_).length(0))
                rhs135_ = d_736_v4_
                rhs136_ = (_dafny.SeqWithoutIsStrInference([(self).f37 for d_813_i6_ in range(default__.abs(74))])).set(default__.safeIndex((self).f25, len(_dafny.SeqWithoutIsStrInference([(self).f37 for d_814_i6_ in range(default__.abs(74))]))), ((self).f25) + ((self).f25))
                lhs105_ = d_808_v54_
                lhs106_ = default__.safeIndex(941, (d_808_v54_).length(0))
                lhs107_ = d_810_v55_
                lhs108_ = default__.safeIndex(518, (d_810_v55_).length(0))
                lhs105_[lhs106_] = rhs135_
                lhs107_[lhs108_] = rhs136_
                (globalState).f12 = p1
            elif True:
                d_738_v6_ = d_738_v6_
                index136_ = default__.safeIndex(314, (d_739_v7_).length(0))
                (d_739_v7_)[index136_] = -300
                d_815_v56_: _dafny.MultiSet
                d_815_v56_ = _dafny.MultiSet([(self).f37])
                d_816_v57_: _dafny.Set
                d_816_v57_ = _dafny.Set({(d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))]})
                d_817_v58_: _dafny.Seq
                d_817_v58_ = _dafny.SeqWithoutIsStrInference([(d_815_v56_).set((self).f25, default__.abs(len(d_737_v5_))), d_815_v56_, _dafny.MultiSet([len(d_816_v57_)]), d_815_v56_, d_815_v56_])
                d_818_v59_: _dafny.Map
                d_818_v59_ = _dafny.Map({613: d_736_v4_})
                d_819_v60_: _dafny.Set
                d_819_v60_ = _dafny.Set({len(d_817_v58_), len(d_818_v59_), (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))], (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))]})
                d_820_v61_: _dafny.Array
                def lambda43_(d_821_p1_):
                    def lambda44_(d_822_i7_):
                        return d_821_p1_

                    return lambda44_

                init24_ = lambda43_(p1)
                nw125_ = _dafny.Array(None, 4)
                for i0_24_ in range(nw125_.length(0)):
                    nw125_[i0_24_] = init24_(i0_24_)
                d_820_v61_ = nw125_
                d_823_v62_: _dafny.Map
                d_823_v62_ = _dafny.Map({d_820_v61_: d_819_v60_})
                d_824_v63_: _dafny.Seq
                d_824_v63_ = _dafny.SeqWithoutIsStrInference([d_820_v61_, d_820_v61_])
                rhs137_ = ((d_823_v62_)[((d_824_v63_)[default__.safeIndex((self).f37, len(d_824_v63_))] if default__.fm0((self).f25, globalState) else d_820_v61_)] if (((d_824_v63_)[default__.safeIndex((self).f37, len(d_824_v63_))] if default__.fm0((self).f25, globalState) else d_820_v61_)) in (d_823_v62_) else d_816_v57_)
                rhs138_ = (_dafny.Set({len(d_732_v0_), len(_dafny.Set({(self).f38})), (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))]})).intersection(d_819_v60_)
                d_819_v60_ = rhs137_
                d_816_v57_ = rhs138_
                d_825_v64_: _dafny.Seq
                d_825_v64_ = _dafny.SeqWithoutIsStrInference([d_746_v14_, d_739_v7_, d_746_v14_, d_739_v7_, d_739_v7_])
                d_825_v64_ = _dafny.SeqWithoutIsStrInference([d_739_v7_, d_746_v14_, d_746_v14_, d_746_v14_])
                (globalState).f12 = (self).f38
            d_826_v65_: _dafny.Map
            d_826_v65_ = _dafny.Map({p1: p1})
            r0 = default__.fm44(((d_826_v65_).set(p1, False)) == (d_826_v65_), p1, p1, (self).f37, globalState)
            d_733_v1_ = (d_733_v1_).set(default__.safeDivisionInt((0) - ((self).f25), (d_743_v11_).cardinality), (self).f38)
        d_827_v66_: _dafny.Seq
        d_827_v66_ = _dafny.SeqWithoutIsStrInference([d_735_v3_, d_735_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igt"))])
        d_828_v67_: D4
        d_828_v67_ = D4_DC13(len(d_827_v66_), p1, default__.fm0((d_734_v2_)[default__.safeIndex(-752, len(d_734_v2_))], globalState), p1)
        d_829_v68_: _dafny.Map
        d_829_v68_ = _dafny.Map({default__.safeDivisionInt((self).f25, (self).f25): (d_828_v67_).cf19})
        d_829_v68_ = (d_829_v68_).set(len(default__.fm1(True, globalState)), (d_739_v7_)[default__.safeIndex(314, (d_739_v7_).length(0))])
        (globalState).f12 = p1
        r0 = (-59) + ((self).f25)
        return r0

    def m18(self, p0, p1, globalState):
        d_830_i0_: int
        d_830_i0_ = 0
        with _dafny.label("7"):
            while (self).f38:
                with _dafny.c_label("7"):
                    if (d_830_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_830_i0_ = (d_830_i0_) + (1)
                    d_831_v0_: str
                    d_831_v0_ = _dafny.CodePoint('g')
                    (globalState).f23 = d_831_v0_
                    d_832_v1_: int
                    d_832_v1_ = 880
                    d_832_v1_ = (147) * (d_832_v1_)
                    d_833_v2_: D12
                    d_833_v2_ = D12_DC32(p0)
                    d_833_v2_ = d_833_v2_
                    if (self).f38:
                        d_834_v3_: _dafny.MultiSet
                        d_834_v3_ = _dafny.MultiSet([d_831_v0_])
                        d_835_v4_: _dafny.Seq
                        d_835_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcea"))
                        d_836_v5_: _dafny.Seq
                        d_836_v5_ = _dafny.SeqWithoutIsStrInference([d_835_v4_, d_835_v4_, d_835_v4_])
                        d_837_v6_: _dafny.Set
                        d_837_v6_ = _dafny.Set({(self).f37, d_832_v1_, (self).f37, 579})
                        d_838_v7_: _dafny.Seq
                        d_838_v7_ = _dafny.SeqWithoutIsStrInference([len(d_835_v4_)])
                        d_839_v8_: _dafny.Map
                        d_839_v8_ = _dafny.Map({len(d_838_v7_): d_832_v1_})
                        d_840_v9_: _dafny.Map
                        d_840_v9_ = _dafny.Map({d_835_v4_: default__.fm47(p1, default__.fm44(p0, p0, True, d_832_v1_, globalState), len(_dafny.Map({(self).f37: len(d_839_v8_)})), d_831_v0_, globalState)})
                        d_841_v10_: _dafny.MultiSet
                        d_841_v10_ = _dafny.MultiSet([(0) - ((self).f25)])
                        d_842_v11_: _dafny.Seq
                        d_842_v11_ = _dafny.SeqWithoutIsStrInference([(self).f38])
                        d_843_v12_: _dafny.Map
                        d_843_v12_ = _dafny.Map({p0: len(d_842_v11_)})
                        rhs139_ = (_dafny.MultiSet([d_831_v0_, _dafny.CodePoint('f')]) if not(p1) else d_834_v3_)
                        rhs140_ = (d_836_v5_) + (d_836_v5_)
                        rhs141_ = 829
                        rhs142_ = (d_837_v6_).intersection(((d_840_v9_)[d_835_v4_] if (d_835_v4_) in (d_840_v9_) else d_837_v6_))
                        rhs143_ = default__.fm44((d_841_v10_).issubset(d_841_v10_), default__.fm0(default__.fm44((self).f38, p0, p0, ((d_843_v12_)[p0] if (p0) in (d_843_v12_) else 993), globalState), globalState), p0, (self).f25, globalState)
                        d_834_v3_ = rhs139_
                        d_836_v5_ = rhs140_
                        d_832_v1_ = rhs141_
                        d_837_v6_ = rhs142_
                        d_832_v1_ = rhs143_
                        (globalState).f6 = _dafny.SeqWithoutIsStrInference([(self).f37 for d_844_i1_ in range(default__.abs(-146))])
                        d_845_v13_: _dafny.Map
                        d_845_v13_ = _dafny.Map({p0: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')]))})
                        d_846_v14_: _dafny.Map
                        d_846_v14_ = _dafny.Map({d_831_v0_: p0})
                        d_847_v15_: _dafny.Array
                        nw126_ = _dafny.Array(int(0), 25)
                        d_847_v15_ = nw126_
                        d_848_v16_: D0
                        d_848_v16_ = D0_DC3(len(_dafny.SeqWithoutIsStrInference([d_847_v15_])), (self).f38)
                        d_849_v17_: _dafny.Array
                        nw127_ = _dafny.Array(None, 23)
                        nw127_[int(0)] = _dafny.Map({p0: d_834_v3_})
                        nw127_[int(1)] = (d_845_v13_) | ((d_845_v13_).set((self).f38, d_834_v3_))
                        nw127_[int(2)] = d_845_v13_
                        nw127_[int(3)] = d_845_v13_
                        nw127_[int(4)] = d_845_v13_
                        nw127_[int(5)] = d_845_v13_
                        nw127_[int(6)] = (_dafny.Map({(self).f38: d_834_v3_})) | (d_845_v13_)
                        nw127_[int(7)] = _dafny.Map({p1: d_834_v3_})
                        nw127_[int(8)] = d_845_v13_
                        nw127_[int(9)] = (d_845_v13_ if (self).f38 else d_845_v13_)
                        nw127_[int(10)] = d_845_v13_
                        nw127_[int(11)] = default__.fm48((self).f38, d_832_v1_, d_832_v1_, len(((d_835_v4_).set(default__.safeIndex(len(d_846_v14_), len(d_835_v4_)), d_831_v0_)).set(default__.safeIndex((self).f25, len((d_835_v4_).set(default__.safeIndex(len(d_846_v14_), len(d_835_v4_)), d_831_v0_))), d_831_v0_)), globalState)
                        nw127_[int(12)] = _dafny.Map({(self).f38: d_834_v3_})
                        nw127_[int(13)] = _dafny.Map({p0: (d_834_v3_).set(_dafny.CodePoint('j'), default__.abs(d_832_v1_))})
                        nw127_[int(14)] = d_845_v13_
                        nw127_[int(15)] = d_845_v13_
                        nw127_[int(16)] = d_845_v13_
                        nw127_[int(17)] = d_845_v13_
                        nw127_[int(18)] = d_845_v13_
                        nw127_[int(19)] = (d_845_v13_).set((d_848_v16_).cf5, (d_834_v3_).set(d_831_v0_, default__.abs((self).f37)))
                        nw127_[int(20)] = d_845_v13_
                        nw127_[int(21)] = d_845_v13_
                        nw127_[int(22)] = d_845_v13_
                        d_849_v17_ = nw127_
                        index137_ = default__.safeIndex(289, (d_849_v17_).length(0))
                        (d_849_v17_)[index137_] = d_845_v13_
                        index138_ = default__.safeIndex(289, (d_849_v17_).length(0))
                        (d_849_v17_)[index138_] = (_dafny.Map({p0: d_834_v3_})) | (default__.fm48(default__.fm0((self).f37, globalState), (self).f37, (self).f37, d_832_v1_, globalState))
                        (globalState).f12 = (default__.fm0(d_832_v1_, globalState)) or (default__.fm0(len(d_842_v11_), globalState))
                        d_850_v18_: _dafny.Array
                        nw128_ = _dafny.Array(_dafny.Array(None, 0), 6)
                        d_850_v18_ = nw128_
                        index139_ = default__.safeIndex(777, (d_850_v18_).length(0))
                        (d_850_v18_)[index139_] = d_847_v15_
                        index140_ = default__.safeIndex(777, (d_850_v18_).length(0))
                        rhs144_ = d_832_v1_
                        rhs145_ = (_dafny.SeqWithoutIsStrInference([p0])) + (d_842_v11_)
                        rhs146_ = (self).f38
                        rhs147_ = default__.safeModuloInt(d_832_v1_, (self).f25)
                        rhs148_ = d_847_v15_
                        lhs109_ = globalState
                        lhs110_ = d_850_v18_
                        lhs111_ = default__.safeIndex(777, (d_850_v18_).length(0))
                        d_832_v1_ = rhs144_
                        d_842_v11_ = rhs145_
                        lhs109_.f12 = rhs146_
                        d_832_v1_ = rhs147_
                        lhs110_[lhs111_] = rhs148_
                    elif True:
                        d_851_v19_: _dafny.Seq
                        d_851_v19_ = _dafny.SeqWithoutIsStrInference([(self).f38, (self).f38])
                        d_831_v0_ = (d_831_v0_ if (d_851_v19_)[default__.safeIndex((self).f25, len(d_851_v19_))] else _dafny.CodePoint('l'))
                        d_852_v20_: _dafny.Array
                        d_853_v21_: _dafny.Array
                        d_854_v22_: bool
                        out22_: _dafny.Array
                        out23_: _dafny.Array
                        out24_: bool
                        out22_, out23_, out24_ = default__.m0(globalState)
                        d_852_v20_ = out22_
                        d_853_v21_ = out23_
                        d_854_v22_ = out24_
                        d_855_v23_: _dafny.Array
                        d_856_v24_: _dafny.Array
                        d_857_v25_: bool
                        out25_: _dafny.Array
                        out26_: _dafny.Array
                        out27_: bool
                        out25_, out26_, out27_ = default__.m0(globalState)
                        d_855_v23_ = out25_
                        d_856_v24_ = out26_
                        d_857_v25_ = out27_
                        d_858_v26_: _dafny.Seq
                        d_858_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppywdo"))
                        rhs149_ = d_858_v26_
                        rhs150_ = (self).f25
                        rhs151_ = (d_852_v20_ if d_857_v25_ else d_852_v20_)
                        rhs152_ = (((self).f38 if d_854_v22_ else p1)) == (p0)
                        rhs153_ = not(d_857_v25_)
                        lhs112_ = globalState
                        lhs113_ = globalState
                        lhs114_ = globalState
                        lhs112_.f10 = rhs149_
                        d_832_v1_ = rhs150_
                        d_852_v20_ = rhs151_
                        lhs113_.f12 = rhs152_
                        lhs114_.f12 = rhs153_
                        d_859_v27_: _dafny.Map
                        d_859_v27_ = _dafny.Map({(self).f25: d_851_v19_})
                        d_859_v27_ = (d_859_v27_).set((self).f37, d_851_v19_)
                    pass
            pass
        d_860_v28_: _dafny.Seq
        d_860_v28_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_861_i2_: int
        d_861_i2_ = 0
        with _dafny.label("8"):
            while default__.fm0(default__.safeDivisionInt(default__.fm44(p0, p0, p0, (d_860_v28_)[default__.safeIndex((self).f25, len(d_860_v28_))], globalState), (0) - (len(_dafny.Set({(self).f25})))), globalState):
                with _dafny.c_label("8"):
                    if (d_861_i2_) >= (100):
                        raise _dafny.Break("8")
                    d_861_i2_ = (d_861_i2_) + (1)
                    d_862_v29_: str
                    d_862_v29_ = _dafny.CodePoint('r')
                    (globalState).f23 = d_862_v29_
                    d_863_v30_: _dafny.Seq
                    d_863_v30_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_864_v31_: _dafny.Seq
                    d_864_v31_ = _dafny.SeqWithoutIsStrInference([p0, (d_863_v30_)[default__.safeIndex((self).f37, len(d_863_v30_))], (self).f38, default__.fm0((self).f25, globalState)])
                    if ((self).f38) not in (d_863_v30_):
                        d_862_v29_ = d_862_v29_
                        d_865_v32_: int
                        d_865_v32_ = 578
                        d_865_v32_ = ((self).f37) + ((d_865_v32_) + ((self).f25))
                        (globalState).f12 = (self).f38
                        d_866_v33_: _dafny.Array
                        def lambda45_(d_867_v32_):
                            def lambda46_(d_868_i3_):
                                return D9_DC27((self).f37, d_867_v32_)

                            return lambda46_

                        init25_ = lambda45_(d_865_v32_)
                        nw129_ = _dafny.Array(None, 22)
                        for i0_25_ in range(nw129_.length(0)):
                            nw129_[i0_25_] = init25_(i0_25_)
                        d_866_v33_ = nw129_
                        d_869_v34_: _dafny.Seq
                        d_869_v34_ = _dafny.SeqWithoutIsStrInference([d_866_v33_, d_866_v33_, d_866_v33_])
                        d_869_v34_ = d_869_v34_
                        d_870_v35_: _dafny.Seq
                        d_870_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                        d_871_v36_: _dafny.Array
                        nw130_ = _dafny.Array(None, 3)
                        nw130_[int(0)] = p1
                        nw130_[int(1)] = (self).f38
                        nw130_[int(2)] = p1
                        d_871_v36_ = nw130_
                        d_872_v37_: _dafny.Map
                        d_872_v37_ = _dafny.Map({d_870_v35_: d_871_v36_})
                        d_872_v37_ = (d_872_v37_).set(d_870_v35_, d_871_v36_)
                    elif True:
                        d_873_v38_: C1
                        nw131_ = C1()
                        nw131_.ctor__((self).f37, (self).f26)
                        d_873_v38_ = nw131_
                        d_874_v39_: _dafny.Array
                        def lambda47_(d_875_i4_):
                            return default__.safeDivisionInt(d_875_i4_, (0) - ((self).f25))

                        init26_ = lambda47_
                        nw132_ = _dafny.Array(None, 20)
                        for i0_26_ in range(nw132_.length(0)):
                            nw132_[i0_26_] = init26_(i0_26_)
                        d_874_v39_ = nw132_
                        index141_ = default__.safeIndex(960, (d_874_v39_).length(0))
                        (d_874_v39_)[index141_] = (self).f37
                        index142_ = default__.safeIndex(960, (d_874_v39_).length(0))
                        (d_874_v39_)[index142_] = (self).f37
                        d_876_v40_: D5
                        d_876_v40_ = D5_DC15(_dafny.MultiSet([(d_874_v39_)[default__.safeIndex(960, (d_874_v39_).length(0))]]))
                        d_877_v41_: D0
                        d_877_v41_ = D0_DC2((self).f38)
                        d_876_v40_ = default__.fm49((self).f37, p0, ((self).f37 if (d_877_v41_).cf3 else (self).f25), d_862_v29_, globalState)
                        d_878_v42_: D0
                        d_878_v42_ = D0_DC0(True)
                        d_879_v43_: _dafny.Array
                        nw133_ = _dafny.Array(None, 29)
                        nw133_[int(0)] = (self).f38
                        nw133_[int(1)] = False
                        nw133_[int(2)] = p1
                        nw133_[int(3)] = p1
                        nw133_[int(4)] = p1
                        nw133_[int(5)] = True
                        nw133_[int(6)] = (d_878_v42_).cf0
                        nw133_[int(7)] = (self).f38
                        nw133_[int(8)] = p1
                        nw133_[int(9)] = p1
                        nw133_[int(10)] = p1
                        nw133_[int(11)] = False
                        nw133_[int(12)] = p0
                        nw133_[int(13)] = p1
                        nw133_[int(14)] = p1
                        nw133_[int(15)] = True
                        nw133_[int(16)] = not((self).f38)
                        nw133_[int(17)] = p0
                        nw133_[int(18)] = p0
                        nw133_[int(19)] = p1
                        nw133_[int(20)] = p1
                        nw133_[int(21)] = p1
                        nw133_[int(22)] = False
                        nw133_[int(23)] = (self).f38
                        nw133_[int(24)] = p0
                        nw133_[int(25)] = not(p0)
                        nw133_[int(26)] = False
                        nw133_[int(27)] = not(default__.fm0((d_874_v39_)[default__.safeIndex(960, (d_874_v39_).length(0))], globalState))
                        nw133_[int(28)] = default__.fm0((self).f37, globalState)
                        d_879_v43_ = nw133_
                        d_880_v44_: _dafny.Map
                        d_880_v44_ = _dafny.Map({d_879_v43_: (self).f38})
                        d_881_v45_: _dafny.Set
                        d_881_v45_ = _dafny.Set({default__.fm0((self).f25, globalState)})
                        d_882_v46_: D12
                        d_882_v46_ = D12_DC32((self).f38)
                        d_883_v47_: _dafny.Map
                        d_883_v47_ = _dafny.Map({(self).f38: False})
                        d_884_v48_: _dafny.Map
                        d_884_v48_ = _dafny.Map({d_882_v46_: len(d_883_v47_)})
                        d_885_v49_: _dafny.Array
                        nw134_ = _dafny.Array(None, 28)
                        nw134_[int(0)] = p1
                        nw134_[int(1)] = p0
                        nw134_[int(2)] = (self).f38
                        nw134_[int(3)] = ((self).f38) == (p0)
                        nw134_[int(4)] = p0
                        nw134_[int(5)] = not((p1) and ((self).f38))
                        nw134_[int(6)] = (default__.fm0((self).f25, globalState)) and (((d_880_v44_)[d_879_v43_] if (d_879_v43_) in (d_880_v44_) else (self).f38))
                        nw134_[int(7)] = not (p1) or (p0)
                        nw134_[int(8)] = (self).f38
                        nw134_[int(9)] = (d_881_v45_).isdisjoint(d_881_v45_)
                        nw134_[int(10)] = p1
                        nw134_[int(11)] = ((d_874_v39_)[default__.safeIndex(960, (d_874_v39_).length(0))]) < (130)
                        nw134_[int(12)] = not ((d_863_v30_)[default__.safeIndex((d_874_v39_)[default__.safeIndex(960, (d_874_v39_).length(0))], len(d_863_v30_))]) or (not(not(p0)))
                        nw134_[int(13)] = p0
                        nw134_[int(14)] = (d_882_v46_) not in (d_884_v48_)
                        nw134_[int(15)] = True
                        nw134_[int(16)] = False
                        nw134_[int(17)] = p0
                        nw134_[int(18)] = not(not(default__.fm0((self).f37, globalState)))
                        nw134_[int(19)] = (self).f38
                        nw134_[int(20)] = p1
                        nw134_[int(21)] = p1
                        nw134_[int(22)] = default__.fm0(-100, globalState)
                        nw134_[int(23)] = p1
                        nw134_[int(24)] = (self).f38
                        nw134_[int(25)] = (self).f38
                        nw134_[int(26)] = (self).f38
                        nw134_[int(27)] = p1
                        d_885_v49_ = nw134_
                        index143_ = default__.safeIndex(600, (d_879_v43_).length(0))
                        (d_879_v43_)[index143_] = p0
                        index144_ = default__.safeIndex(600, (d_879_v43_).length(0))
                        (d_879_v43_)[index144_] = default__.fm0((self).f37, globalState)
                        d_886_v50_: _dafny.Seq
                        d_886_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wow"))
                        d_887_v51_: _dafny.Map
                        d_887_v51_ = _dafny.Map({not(p0): (self).f25})
                        d_888_v52_: _dafny.Map
                        d_888_v52_ = _dafny.Map({(len(d_886_v50_)) <= ((D0_DC3((self).f37, p0)).cf4): (d_887_v51_ if p0 else d_887_v51_)})
                        d_888_v52_ = ((d_888_v52_ if (d_879_v43_)[default__.safeIndex(600, (d_879_v43_).length(0))] else d_888_v52_)) | (d_888_v52_)
                    d_889_v53_: _dafny.Map
                    d_889_v53_ = _dafny.Map({p0: p1})
                    d_890_v54_: _dafny.Map
                    d_890_v54_ = _dafny.Map({d_860_v28_: _dafny.Map({d_889_v53_: len(d_864_v31_)})})
                    d_891_v55_: _dafny.Map
                    d_891_v55_ = _dafny.Map({d_889_v53_: len(_dafny.SeqWithoutIsStrInference([(self).f25 for d_892_i5_ in range(default__.abs(704))]))})
                    d_890_v54_ = (d_890_v54_).set(d_860_v28_, d_891_v55_)
                    d_893_v56_: D8
                    d_893_v56_ = D8_DC24(D1_DC5())
                    source17_ = d_893_v56_
                    if source17_.is_DC24:
                        d_894___mcc_h0_ = source17_.cf36
                        d_895_cf36_ = d_894___mcc_h0_
                        d_896_v57_: _dafny.Array
                        nw135_ = _dafny.Array(False, 7)
                        d_896_v57_ = nw135_
                        index145_ = default__.safeIndex(495, (d_896_v57_).length(0))
                        (d_896_v57_)[index145_] = p1
                        index146_ = default__.safeIndex(495, (d_896_v57_).length(0))
                        (d_896_v57_)[index146_] = p0
                        d_897_v58_: _dafny.Seq
                        d_897_v58_ = _dafny.SeqWithoutIsStrInference([d_860_v28_])
                        d_898_v59_: D0
                        d_898_v59_ = D0_DC0((self).f38)
                        d_889_v53_ = (d_889_v53_).set(p1, ((self).f37) < (len((d_897_v58_).set(default__.safeIndex((self).f25, len(d_897_v58_)), (self).fm38(d_898_v59_, globalState)))))
                        d_899_v60_: _dafny.Array
                        def lambda48_(d_900_i6_):
                            return (d_900_i6_) * ((self).f37)

                        init27_ = lambda48_
                        nw136_ = _dafny.Array(None, 3)
                        for i0_27_ in range(nw136_.length(0)):
                            nw136_[i0_27_] = init27_(i0_27_)
                        d_899_v60_ = nw136_
                        d_901_v61_: _dafny.Map
                        d_901_v61_ = _dafny.Map({d_899_v60_: p1})
                        d_902_v62_: _dafny.MultiSet
                        d_902_v62_ = _dafny.MultiSet([False, ((d_889_v53_)[(d_896_v57_)[default__.safeIndex(495, (d_896_v57_).length(0))]] if ((d_896_v57_)[default__.safeIndex(495, (d_896_v57_).length(0))]) in (d_889_v53_) else (d_896_v57_)[default__.safeIndex(495, (d_896_v57_).length(0))])])
                        d_903_v63_: int
                        d_903_v63_ = 497
                        d_904_v64_: _dafny.Seq
                        d_904_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hr"))
                        d_905_v65_: _dafny.Set
                        d_905_v65_ = _dafny.Set({d_904_v64_})
                        d_906_v66_: _dafny.Seq
                        d_906_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_904_v64_}), d_905_v65_])
                        rhs154_ = (((d_906_v66_)[default__.safeIndex(d_903_v63_, len(d_906_v66_))] if (self).f38 else d_905_v65_)).isdisjoint(_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_904_v64_)[default__.safeIndex((self).f37, len(d_904_v64_))] for d_907_i7_ in range(default__.abs(395))]), d_904_v64_}))
                        rhs155_ = d_901_v61_
                        rhs156_ = (_dafny.MultiSet([p1])) | (_dafny.MultiSet([(self).f38, True, (d_896_v57_)[default__.safeIndex(495, (d_896_v57_).length(0))]]))
                        rhs157_ = (d_903_v63_ if (d_896_v57_)[default__.safeIndex(495, (d_896_v57_).length(0))] else (self).f25)
                        lhs115_ = globalState
                        lhs115_.f12 = rhs154_
                        d_901_v61_ = rhs155_
                        d_902_v62_ = rhs156_
                        d_903_v63_ = rhs157_
                        d_903_v63_ = (d_860_v28_)[default__.safeIndex(default__.safeDivisionInt((self).f25, d_903_v63_), len(d_860_v28_))]
                    elif source17_.is_DC23:
                        d_908___mcc_h1_ = source17_.cf35
                        d_909_cf35_ = d_908___mcc_h1_
                        d_910_v67_: C1
                        nw137_ = C1()
                        nw137_.ctor__((0) - ((self).f37), (self).f26)
                        d_910_v67_ = nw137_
                        d_911_v68_: int
                        d_911_v68_ = 889
                        d_911_v68_ = (0) - (default__.safeDivisionInt((self).f37, (self).f25))
                        (globalState).f23 = d_862_v29_
                        rhs158_ = (106) * (398)
                        rhs159_ = d_864_v31_
                        rhs160_ = default__.safeModuloInt((0) - (((0) - ((self).f37) if p0 else d_911_v68_)), (self).f37)
                        d_911_v68_ = rhs158_
                        d_864_v31_ = rhs159_
                        d_911_v68_ = rhs160_
                    elif True:
                        d_912___mcc_h2_ = source17_.cf37
                        d_913_cf37_ = d_912___mcc_h2_
                        d_914_v69_: _dafny.Set
                        d_914_v69_ = _dafny.Set({p0, (self).f38})
                        d_915_v70_: _dafny.Map
                        d_915_v70_ = _dafny.Map({(self).f37: (self).f25})
                        d_916_v71_: _dafny.Map
                        d_916_v71_ = _dafny.Map({(D5_DC16((self).f38)).cf25: -328})
                        d_917_v72_: _dafny.MultiSet
                        d_917_v72_ = _dafny.MultiSet([p1, p0])
                        d_918_v73_: _dafny.Array
                        nw138_ = _dafny.Array(None, 22)
                        nw138_[int(0)] = len(d_914_v69_)
                        nw138_[int(1)] = 332
                        nw138_[int(2)] = len(d_915_v70_)
                        nw138_[int(3)] = ((d_916_v71_)[(self).f38] if ((self).f38) in (d_916_v71_) else (self).f25)
                        nw138_[int(4)] = len(d_860_v28_)
                        nw138_[int(5)] = default__.fm44(p0, (self).f38, p1, (self).f25, globalState)
                        nw138_[int(6)] = (self).f25
                        nw138_[int(7)] = (self).f37
                        nw138_[int(8)] = (self).f25
                        nw138_[int(9)] = (self).f37
                        nw138_[int(10)] = (self).f37
                        nw138_[int(11)] = (self).f37
                        nw138_[int(12)] = (self).f37
                        nw138_[int(13)] = (self).f25
                        nw138_[int(14)] = default__.fm44(not((self).f38), p0, p0, (self).f37, globalState)
                        nw138_[int(15)] = 267
                        nw138_[int(16)] = (d_917_v72_).cardinality
                        nw138_[int(17)] = (self).f25
                        nw138_[int(18)] = (self).f25
                        nw138_[int(19)] = (self).f25
                        nw138_[int(20)] = 562
                        nw138_[int(21)] = (self).f25
                        d_918_v73_ = nw138_
                        d_919_v74_: _dafny.Map
                        d_919_v74_ = _dafny.Map({d_918_v73_: (self).f37})
                        d_919_v74_ = (d_919_v74_).set(d_918_v73_, (self).f37)
                        d_920_v75_: int
                        d_920_v75_ = 547
                        def iife76_():
                            coll40_ = _dafny.Map()
                            compr_40_: int
                            for compr_40_ in (d_860_v28_).Elements:
                                d_921_v76_: int = compr_40_
                                if (d_921_v76_) in (d_860_v28_):
                                    coll40_[default__.safeDivisionInt(d_921_v76_, d_920_v75_)] = d_917_v72_
                            return _dafny.Map(coll40_)
                        d_920_v75_ = (len(iife76_()
                        )) * (d_920_v75_)
                        d_916_v71_ = (d_916_v71_).set(False, (291) + (832))
                        index147_ = default__.safeIndex(332, ((self).f26).length(0))
                        ((self).f26)[index147_] = _dafny.SeqWithoutIsStrInference([d_862_v29_ for d_922_i8_ in range(default__.abs(562))])
                        d_923_v77_: _dafny.Seq
                        d_923_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkvfdt"))
                        index148_ = default__.safeIndex(332, ((self).f26).length(0))
                        ((self).f26)[index148_] = d_923_v77_
                    pass
            pass
        d_924_v78_: _dafny.Array
        def lambda49_(d_925_i9_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmh"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlcfnt")))

        init28_ = lambda49_
        nw139_ = _dafny.Array(None, 26)
        for i0_28_ in range(nw139_.length(0)):
            nw139_[i0_28_] = init28_(i0_28_)
        d_924_v78_ = nw139_
        index149_ = default__.safeIndex(131, (d_924_v78_).length(0))
        (d_924_v78_)[index149_] = (self).f38
        index150_ = default__.safeIndex(131, (d_924_v78_).length(0))
        (d_924_v78_)[index150_] = (p0) == (default__.fm0(len(d_860_v28_), globalState))
        d_926_i10_: int
        d_926_i10_ = 0
        with _dafny.label("9"):
            while (self).f38:
                with _dafny.c_label("9"):
                    if (d_926_i10_) >= (100):
                        raise _dafny.Break("9")
                    d_926_i10_ = (d_926_i10_) + (1)
                    d_927_v79_: _dafny.Array
                    def lambda50_(d_928_p1_):
                        def lambda51_(d_929_i11_):
                            return _dafny.SeqWithoutIsStrInference([d_928_p1_, d_928_p1_])

                        return lambda51_

                    init29_ = lambda50_(p1)
                    nw140_ = _dafny.Array(None, 15)
                    for i0_29_ in range(nw140_.length(0)):
                        nw140_[i0_29_] = init29_(i0_29_)
                    d_927_v79_ = nw140_
                    d_930_v80_: _dafny.Array
                    def lambda52_(d_931_i12_):
                        return (d_931_i12_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f25, (self).f37})])))

                    init30_ = lambda52_
                    nw141_ = _dafny.Array(None, 26)
                    for i0_30_ in range(nw141_.length(0)):
                        nw141_[i0_30_] = init30_(i0_30_)
                    d_930_v80_ = nw141_
                    index151_ = default__.safeIndex(983, (d_930_v80_).length(0))
                    (d_930_v80_)[index151_] = 697
                    d_932_v81_: int
                    d_932_v81_ = 484
                    index152_ = default__.safeIndex(983, (d_930_v80_).length(0))
                    rhs161_ = d_927_v79_
                    rhs162_ = default__.safeModuloInt(519, d_932_v81_)
                    rhs163_ = True
                    rhs164_ = (self).f37
                    lhs116_ = d_930_v80_
                    lhs117_ = default__.safeIndex(983, (d_930_v80_).length(0))
                    lhs118_ = globalState
                    d_927_v79_ = rhs161_
                    lhs116_[lhs117_] = rhs162_
                    lhs118_.f12 = rhs163_
                    d_932_v81_ = rhs164_
                    d_933_v82_: _dafny.Map
                    d_933_v82_ = _dafny.Map({p1: p0})
                    d_934_v83_: _dafny.MultiSet
                    d_934_v83_ = _dafny.MultiSet([d_860_v28_, d_860_v28_])
                    d_935_v84_: _dafny.Set
                    d_935_v84_ = _dafny.Set({540, (self).f37, (d_934_v83_).cardinality})
                    d_933_v82_ = (d_933_v82_).set((d_935_v84_).issubset(d_935_v84_), (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))])
                    index153_ = default__.safeIndex(131, (d_924_v78_).length(0))
                    (d_924_v78_)[index153_] = False
                    d_935_v84_ = d_935_v84_
                    pass
            pass
        d_936_v85_: _dafny.Seq
        d_936_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iei"))
        d_937_v86_: str
        d_937_v86_ = _dafny.CodePoint('h')
        d_938_v87_: _dafny.MultiSet
        d_938_v87_ = _dafny.MultiSet([len((d_936_v85_).set(default__.safeIndex((self).f37, len(d_936_v85_)), d_937_v86_))])
        hi4_ = (d_938_v87_).cardinality
        for d_939_i13_ in range((self).f37, hi4_):
            d_940_v88_: _dafny.Set
            d_940_v88_ = _dafny.Set({(d_860_v28_)[default__.safeIndex((self).f25, len(d_860_v28_))]})
            d_941_v89_: _dafny.Map
            d_941_v89_ = _dafny.Map({d_939_i13_: p0})
            d_942_v90_: D5
            d_942_v90_ = D5_DC17(len(d_940_v88_), ((d_941_v89_)[d_939_i13_] if (d_939_i13_) in (d_941_v89_) else p0))
            d_943_v91_: _dafny.Seq
            d_943_v91_ = _dafny.SeqWithoutIsStrInference([p1, (self).f38])
            d_944_v92_: _dafny.MultiSet
            d_944_v92_ = _dafny.MultiSet([(d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))]])
            d_945_v93_: _dafny.Set
            d_945_v93_ = _dafny.Set({default__.fm50(d_939_i13_, d_936_v85_, (d_942_v90_).cf27, len(d_943_v91_), globalState), d_944_v92_, d_944_v92_, (d_944_v92_).set((d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))], default__.abs((0) - ((self).f37))), d_944_v92_})
            d_946_v94_: int
            d_946_v94_ = 821
            d_947_v95_: _dafny.Map
            d_947_v95_ = _dafny.Map({(self).f37: d_939_i13_})
            index154_ = default__.safeIndex(131, (d_924_v78_).length(0))
            rhs165_ = (d_938_v87_).ispropersubset((d_938_v87_) | (d_938_v87_))
            rhs166_ = default__.fm51((d_940_v88_) - (default__.fm47((self).f38, 170, len(d_947_v95_), d_937_v86_, globalState)), -486, d_937_v86_, d_946_v94_, globalState)
            rhs167_ = (d_860_v28_)[default__.safeIndex(default__.safeDivisionInt((d_860_v28_)[default__.safeIndex((0) - ((self).f37), len(d_860_v28_))], (self).f37), len(d_860_v28_))]
            lhs119_ = d_924_v78_
            lhs120_ = default__.safeIndex(131, (d_924_v78_).length(0))
            lhs119_[lhs120_] = rhs165_
            d_945_v93_ = rhs166_
            d_946_v94_ = rhs167_
            d_943_v91_ = (default__.fm1((self).f38, globalState)) + (_dafny.SeqWithoutIsStrInference([(self).f38]))
            d_948_v96_: _dafny.Set
            d_948_v96_ = _dafny.Set({p1, p0})
            index155_ = default__.safeIndex(131, (d_924_v78_).length(0))
            rhs168_ = False
            rhs169_ = d_946_v94_
            rhs170_ = default__.safeDivisionInt(default__.safeModuloInt(len(d_948_v96_), -109), -219)
            rhs171_ = (d_943_v91_)[default__.safeIndex((default__.fm44((self).f38, (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))], True, (self).f37, globalState)) * (d_939_i13_), len(d_943_v91_))]
            rhs172_ = d_946_v94_
            lhs121_ = d_924_v78_
            lhs122_ = default__.safeIndex(131, (d_924_v78_).length(0))
            lhs123_ = globalState
            lhs121_[lhs122_] = rhs168_
            d_946_v94_ = rhs169_
            d_946_v94_ = rhs170_
            lhs123_.f12 = rhs171_
            d_946_v94_ = rhs172_
            index156_ = default__.safeIndex(131, (d_924_v78_).length(0))
            (d_924_v78_)[index156_] = p0
        d_949_v97_: D12
        d_949_v97_ = D12_DC32(False)
        d_950_v98_: D12
        d_950_v98_ = D12_DC34(d_949_v97_)
        source18_ = d_950_v98_
        if source18_.is_DC32:
            d_951___mcc_h3_ = source18_.cf47
            d_952_cf47_ = d_951___mcc_h3_
            d_952_cf47_ = (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))]
            d_953_v99_: C1
            nw142_ = C1()
            nw142_.ctor__(802, (self).f26)
            d_953_v99_ = nw142_
            d_954_v100_: D0
            d_954_v100_ = D0_DC0(d_952_cf47_)
            d_955_v101_: _dafny.Set
            d_955_v101_ = _dafny.Set({len(d_936_v85_), len((self).fm38(d_954_v100_, globalState))})
            d_956_v102_: _dafny.Map
            d_956_v102_ = _dafny.Map({d_955_v101_: (self).f25})
            d_957_v104_: _dafny.Map
            d_957_v104_ = _dafny.Map({(0) - ((self).f37): (self).f25})
            def iife77_():
                coll41_ = _dafny.Set()
                compr_41_: int
                for compr_41_ in (_dafny.MultiSet(d_860_v28_)).Elements:
                    d_958_v103_: int = compr_41_
                    if (d_958_v103_) in (_dafny.MultiSet(d_860_v28_)):
                        coll41_ = coll41_.union(_dafny.Set([(d_958_v103_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False}))])))]))
                return _dafny.Set(coll41_)
            d_956_v102_ = (d_956_v102_).set((d_955_v101_ if (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))] else iife77_()
            ), ((d_957_v104_)[(self).f37] if ((self).f37) in (d_957_v104_) else (self).f37))
            d_959_v105_: _dafny.Seq
            d_959_v105_ = _dafny.SeqWithoutIsStrInference([d_955_v101_])
            d_960_v106_: _dafny.Map
            d_960_v106_ = _dafny.Map({(self).f37: d_937_v86_})
            d_961_v108_: _dafny.Map
            d_961_v108_ = _dafny.Map({(self).f37: d_952_cf47_})
            d_962_v111_: _dafny.Array
            nw143_ = _dafny.Array(None, 23)
            nw143_[int(0)] = default__.fm47(not(not(p1)), (self).f37, len(d_955_v101_), d_937_v86_, globalState)
            nw143_[int(1)] = d_955_v101_
            nw143_[int(2)] = d_955_v101_
            nw143_[int(3)] = (d_955_v101_) - (d_955_v101_)
            nw143_[int(4)] = (d_959_v105_)[default__.safeIndex(392, len(d_959_v105_))]
            nw143_[int(5)] = _dafny.Set({(self).f37, (d_938_v87_).cardinality, (self).f25})
            nw143_[int(6)] = d_955_v101_
            nw143_[int(7)] = (d_955_v101_).intersection(d_955_v101_)
            nw143_[int(8)] = (d_955_v101_).intersection(d_955_v101_)
            nw143_[int(9)] = _dafny.Set({(self).f25})
            nw143_[int(10)] = d_955_v101_
            nw143_[int(11)] = d_955_v101_
            nw143_[int(12)] = d_955_v101_
            nw143_[int(13)] = (default__.fm47((self).f38, (self).f25, (self).f37, d_937_v86_, globalState)) | (d_955_v101_)
            nw143_[int(14)] = d_955_v101_
            nw143_[int(15)] = d_955_v101_
            def iife78_():
                coll42_ = _dafny.Set()
                compr_42_: int
                for compr_42_ in (d_960_v106_).keys.Elements:
                    d_963_v107_: int = compr_42_
                    if (d_963_v107_) in (d_960_v106_):
                        coll42_ = coll42_.union(_dafny.Set([default__.safeModuloInt(d_963_v107_, -123)]))
                return _dafny.Set(coll42_)
            nw143_[int(16)] = (d_955_v101_) - (iife78_()
            )
            nw143_[int(17)] = (d_955_v101_) | (d_955_v101_)
            nw143_[int(18)] = _dafny.Set({(0) - (default__.fm44(True, (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))], ((d_961_v108_)[(self).f25] if ((self).f25) in (d_961_v108_) else True), default__.fm44(d_952_cf47_, p1, (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))], (self).f37, globalState), globalState))})
            nw143_[int(19)] = (d_955_v101_).intersection(d_955_v101_)
            nw143_[int(20)] = _dafny.Set({(self).f37, (self).f37})
            def iife79_():
                coll43_ = _dafny.Set()
                compr_43_: int
                for compr_43_ in (d_957_v104_).keys.Elements:
                    d_964_v109_: int = compr_43_
                    if (d_964_v109_) in (d_957_v104_):
                        coll43_ = coll43_.union(_dafny.Set([(d_964_v109_) - (172)]))
                return _dafny.Set(coll43_)
            def iife80_():
                coll44_ = _dafny.Set()
                compr_44_: int
                for compr_44_ in (d_955_v101_).Elements:
                    d_965_v110_: int = compr_44_
                    if (d_965_v110_) in (d_955_v101_):
                        coll44_ = coll44_.union(_dafny.Set([(d_965_v110_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))))]))
                return _dafny.Set(coll44_)
            nw143_[int(21)] = (iife79_()
             if (self).f38 else (D4_DC11(iife80_()
)).cf18)
            nw143_[int(22)] = d_955_v101_
            d_962_v111_ = nw143_
            index157_ = default__.safeIndex(346, (d_962_v111_).length(0))
            (d_962_v111_)[index157_] = d_955_v101_
            d_966_v112_: _dafny.Seq
            d_966_v112_ = _dafny.SeqWithoutIsStrInference([(self).f38])
            d_967_v113_: int
            d_967_v113_ = 533
            d_968_v114_: _dafny.Array
            nw144_ = _dafny.Array(None, 2)
            nw144_[int(0)] = d_967_v113_
            nw144_[int(1)] = ((self).f25) * ((self).f37)
            d_968_v114_ = nw144_
            index158_ = default__.safeIndex(289, (d_968_v114_).length(0))
            (d_968_v114_)[index158_] = (len(d_961_v108_)) - ((self).f25)
            index159_ = default__.safeIndex(346, (d_962_v111_).length(0))
            index160_ = default__.safeIndex(131, (d_924_v78_).length(0))
            index161_ = default__.safeIndex(289, (d_968_v114_).length(0))
            rhs173_ = d_955_v101_
            rhs174_ = (d_966_v112_) + (d_966_v112_)
            rhs175_ = 947
            rhs176_ = (default__.safeDivisionInt((self).f37, d_967_v113_)) == ((len(d_936_v85_)) + ((self).f25))
            rhs177_ = (self).f25
            lhs124_ = d_962_v111_
            lhs125_ = default__.safeIndex(346, (d_962_v111_).length(0))
            lhs126_ = d_924_v78_
            lhs127_ = default__.safeIndex(131, (d_924_v78_).length(0))
            lhs128_ = d_968_v114_
            lhs129_ = default__.safeIndex(289, (d_968_v114_).length(0))
            lhs124_[lhs125_] = rhs173_
            d_966_v112_ = rhs174_
            d_967_v113_ = rhs175_
            lhs126_[lhs127_] = rhs176_
            lhs128_[lhs129_] = rhs177_
        elif source18_.is_DC33:
            d_969___mcc_h4_ = source18_.cf48
            d_970___mcc_h5_ = source18_.cf49
            d_971___mcc_h6_ = source18_.cf50
            d_972___mcc_h7_ = source18_.cf51
            d_973_cf51_ = d_972___mcc_h7_
            d_974_cf50_ = d_971___mcc_h6_
            d_975_cf49_ = d_970___mcc_h5_
            d_976_cf48_ = d_969___mcc_h4_
            d_977_v116_: _dafny.Set
            d_977_v116_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hglnqof")), (d_936_v85_).set(default__.safeIndex(len(d_936_v85_), len(d_936_v85_)), d_937_v86_)})
            def iife81_():
                coll45_ = _dafny.Map()
                compr_45_: _dafny.Seq
                for compr_45_ in (d_977_v116_).Elements:
                    d_978_v115_: _dafny.Seq = compr_45_
                    if (d_978_v115_) in (d_977_v116_):
                        coll45_[d_978_v115_] = False
                return _dafny.Map(coll45_)
            rhs178_ = not(False)
            rhs179_ = (self).f37
            rhs180_ = (0) - ((self).f37)
            rhs181_ = d_973_cf51_
            rhs182_ = len(iife81_()
            )
            d_974_cf50_ = rhs178_
            d_973_cf51_ = rhs179_
            d_973_cf51_ = rhs180_
            d_973_cf51_ = rhs181_
            d_973_cf51_ = rhs182_
            d_979_v117_: _dafny.MultiSet
            d_979_v117_ = _dafny.MultiSet([(d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))]])
            if not((default__.fm0((self).f37, globalState) if (d_974_cf50_) and (p1) else (d_979_v117_).isdisjoint(d_979_v117_))):
                d_980_v118_: _dafny.Map
                d_980_v118_ = _dafny.Map({d_973_cf51_: d_936_v85_})
                d_981_v119_: _dafny.Map
                d_981_v119_ = _dafny.Map({d_936_v85_: d_980_v118_})
                d_981_v119_ = (d_981_v119_).set(d_936_v85_, d_980_v118_)
                (globalState).f12 = (self).f38
                (globalState).f12 = ((d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))] if (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))] else p1)
                d_982_v120_: _dafny.Map
                d_982_v120_ = _dafny.Map({(self).f37: p0})
                d_983_v121_: _dafny.Seq
                d_983_v121_ = _dafny.SeqWithoutIsStrInference([d_982_v120_])
                d_984_v122_: _dafny.Set
                d_984_v122_ = _dafny.Set({len(d_983_v121_), (self).f25, (self).f37})
                d_985_v123_: _dafny.MultiSet
                d_985_v123_ = _dafny.MultiSet([d_984_v122_, d_984_v122_])
                d_973_cf51_ = ((d_985_v123_)[d_984_v122_] if (d_984_v122_) in (d_985_v123_) else d_973_cf51_)
                d_986_v124_: _dafny.Array
                def lambda53_(d_987_v28_):
                    def lambda54_(d_988_i14_):
                        return (d_988_i14_) * (len(d_987_v28_))

                    return lambda54_

                init31_ = lambda53_(d_860_v28_)
                nw145_ = _dafny.Array(None, 8)
                for i0_31_ in range(nw145_.length(0)):
                    nw145_[i0_31_] = init31_(i0_31_)
                d_986_v124_ = nw145_
                d_986_v124_ = d_986_v124_
            elif True:
                d_989_v125_: _dafny.Set
                d_989_v125_ = _dafny.Set({738})
                d_990_v126_: _dafny.Map
                d_990_v126_ = _dafny.Map({False: (self).f38})
                d_991_v127_: D9
                d_991_v127_ = D9_DC27(len(d_989_v125_), len(d_990_v126_))
                d_973_cf51_ = default__.safeModuloInt((0) - ((d_991_v127_).cf40), d_973_cf51_)
                d_992_v128_: C1
                nw146_ = C1()
                nw146_.ctor__((self).f25, (self).f26)
                d_992_v128_ = nw146_
                d_993_v129_: _dafny.Map
                d_993_v129_ = _dafny.Map({d_992_v128_: default__.safeModuloInt(982, (self).f37)})
                d_993_v129_ = d_993_v129_
                d_860_v28_ = d_860_v28_
                (globalState).f12 = False
                d_994_v130_: _dafny.Array
                nw147_ = _dafny.Array(_dafny.Set({}), 7)
                d_994_v130_ = nw147_
                index162_ = default__.safeIndex(127, (d_994_v130_).length(0))
                (d_994_v130_)[index162_] = d_989_v125_
                d_995_v131_: D4
                d_995_v131_ = D4_DC11(default__.fm47(default__.fm0(d_973_cf51_, globalState), (self).f37, (self).f25, _dafny.CodePoint('v'), globalState))
                index163_ = default__.safeIndex(127, (d_994_v130_).length(0))
                rhs183_ = (d_995_v131_).cf18
                rhs184_ = d_936_v85_
                rhs185_ = (self).f37
                rhs186_ = (0) - (((self).f37 if (d_973_cf51_) != (747) else ((0) - (d_973_cf51_)) - ((self).f37)))
                lhs130_ = d_994_v130_
                lhs131_ = default__.safeIndex(127, (d_994_v130_).length(0))
                lhs130_[lhs131_] = rhs183_
                d_936_v85_ = rhs184_
                d_973_cf51_ = rhs185_
                d_973_cf51_ = rhs186_
            if d_974_cf50_:
                d_973_cf51_ = (self).f37
                d_996_v132_: C2
                nw148_ = C2()
                nw148_.ctor__()
                d_996_v132_ = nw148_
                d_997_v133_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_997_v133_ = nw149_
                d_998_v134_: _dafny.Array
                def lambda55_(d_999_v85_):
                    def lambda56_(d_1000_i15_):
                        return (d_999_v85_)[default__.safeIndex((self).f37, len(d_999_v85_))]

                    return lambda56_

                init32_ = lambda55_(d_936_v85_)
                nw150_ = _dafny.Array(None, 26)
                for i0_32_ in range(nw150_.length(0)):
                    nw150_[i0_32_] = init32_(i0_32_)
                d_998_v134_ = nw150_
                d_1001_v135_: D13
                d_1001_v135_ = D13_DC35(d_998_v134_)
                index164_ = default__.safeIndex(718, (d_997_v133_).length(0))
                (d_997_v133_)[index164_] = (d_1001_v135_).cf53
                index165_ = default__.safeIndex(718, (d_997_v133_).length(0))
                def lambda57_(d_1002_v86_):
                    def lambda58_(d_1003_i16_):
                        return d_1002_v86_

                    return lambda58_

                init33_ = lambda57_(d_937_v86_)
                nw151_ = _dafny.Array(None, 7)
                for i0_33_ in range(nw151_.length(0)):
                    nw151_[i0_33_] = init33_(i0_33_)
                (d_997_v133_)[index165_] = nw151_
                (globalState).f23 = default__.fm39(p1, globalState)
                (globalState).f12 = not ((d_973_cf51_) > ((self).f37)) or ((self).f38)
            elif True:
                d_1004_v136_: _dafny.Map
                d_1004_v136_ = _dafny.Map({not (p0) or (False): (self).f37})
                d_1005_v137_: _dafny.Map
                d_1005_v137_ = _dafny.Map({d_973_cf51_: (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))]})
                rhs187_ = d_936_v85_
                rhs188_ = ((d_1004_v136_)[not (d_974_cf50_) or ((d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))])] if (not (d_974_cf50_) or ((d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))])) in (d_1004_v136_) else default__.fm44(not(((d_1005_v137_)[d_973_cf51_] if (d_973_cf51_) in (d_1005_v137_) else (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))])), p0, p0, (0) - ((self).f25), globalState))
                lhs132_ = globalState
                lhs132_.f10 = rhs187_
                d_973_cf51_ = rhs188_
                (globalState).f12 = (d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))]
                d_1006_v138_: _dafny.Array
                nw152_ = _dafny.Array(int(0), 26)
                d_1006_v138_ = nw152_
                index166_ = default__.safeIndex(114, (d_1006_v138_).length(0))
                def iife82_():
                    coll46_ = _dafny.Map()
                    compr_46_: _dafny.Seq
                    for compr_46_ in (default__.fm52(d_974_cf50_, (self).f37, d_973_cf51_, globalState)).keys.Elements:
                        d_1007_v139_: _dafny.Seq = compr_46_
                        if (d_1007_v139_) in (default__.fm52(d_974_cf50_, (self).f37, d_973_cf51_, globalState)):
                            coll46_[d_1007_v139_] = 3
                    return _dafny.Map(coll46_)
                (d_1006_v138_)[index166_] = len(iife82_()
                )
                index167_ = default__.safeIndex(114, (d_1006_v138_).length(0))
                (d_1006_v138_)[index167_] = ((self).f25) - (len(default__.fm47(True, len(d_936_v85_), len(d_1005_v137_), d_937_v86_, globalState)))
                index168_ = default__.safeIndex(114, (d_1006_v138_).length(0))
                (d_1006_v138_)[index168_] = default__.safeModuloInt((self).f37, (self).f25)
                d_1008_v140_: _dafny.MultiSet
                d_1008_v140_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))])
                index169_ = default__.safeIndex(131, (d_924_v78_).length(0))
                rhs189_ = ((d_936_v85_) + (d_936_v85_)).set(default__.safeIndex(default__.fm44(default__.fm0((d_1006_v138_)[default__.safeIndex(114, (d_1006_v138_).length(0))], globalState), default__.fm0((self).f25, globalState), not(False), (d_1006_v138_)[default__.safeIndex(114, (d_1006_v138_).length(0))], globalState), len((d_936_v85_) + (d_936_v85_))), d_937_v86_)
                rhs190_ = not(((d_1008_v140_).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_937_v86_ for d_1009_i17_ in range(default__.abs(116))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1010_i18_ in range(default__.abs(-740))])]))).issubset(d_1008_v140_))
                lhs133_ = globalState
                lhs134_ = d_924_v78_
                lhs135_ = default__.safeIndex(131, (d_924_v78_).length(0))
                lhs133_.f10 = rhs189_
                lhs134_[lhs135_] = rhs190_
            d_973_cf51_ = (d_973_cf51_) - (d_973_cf51_)
        elif source18_.is_DC31:
            d_1011___mcc_h8_ = source18_.cf46
            d_1012_cf46_ = d_1011___mcc_h8_
            d_1013_v141_: _dafny.Set
            d_1013_v141_ = _dafny.Set({(d_924_v78_)[default__.safeIndex(131, (d_924_v78_).length(0))]})
            d_1014_v142_: _dafny.Set
            d_1014_v142_ = _dafny.Set({(self).f25, (self).f25})
            d_1015_v143_: D10
            d_1015_v143_ = D10_DC29(p1, (self).f38, len(d_1014_v142_))
            d_1016_v144_: _dafny.Map
            d_1016_v144_ = _dafny.Map({not((self).f38): (d_1013_v141_) - (_dafny.Set({(d_1015_v143_).cf43}))})
            d_1016_v144_ = (d_1016_v144_).set(p1, d_1013_v141_)
            d_1017_v145_: C1
            nw153_ = C1()
            nw153_.ctor__(((self).f25) - ((self).f37), (self).f26)
            d_1017_v145_ = nw153_
            d_1018_v146_: C0
            nw154_ = C0()
            nw154_.ctor__()
            d_1018_v146_ = nw154_
            d_1019_v147_: int
            d_1019_v147_ = 385
            d_1019_v147_ = (0) - (len((d_936_v85_) + (d_936_v85_)))
        elif True:
            d_1020___mcc_h9_ = source18_.cf52
            d_1021_cf52_ = d_1020___mcc_h9_
            d_1022_v148_: C2
            nw155_ = C2()
            nw155_.ctor__()
            d_1022_v148_ = nw155_
            (globalState).f23 = d_937_v86_
            d_1023_v149_: C1
            nw156_ = C1()
            nw156_.ctor__((self).f37, (self).f26)
            d_1023_v149_ = nw156_
            d_1024_v150_: int
            d_1024_v150_ = 519
            d_1024_v150_ = len((d_936_v85_) + (d_936_v85_))

    @property
    def f37(self):
        return self._f37
    @property
    def f38(self):
        return self._f38

class C4(T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        self._f36: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f36, f25, f26):
        (self)._f36 = f36
        (self)._f25 = f25
        (self)._f26 = f26

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r0 = (self).f25
        if not(p1):
            index170_ = default__.safeIndex(990, (p3).length(0))
            (p3)[index170_] = ((self).f25) - ((self).f25)
            index171_ = default__.safeIndex(990, (p3).length(0))
            (p3)[index171_] = (self).f25
            d_1025_v0_: _dafny.Seq
            d_1025_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            d_1026_v1_: _dafny.Set
            d_1026_v1_ = _dafny.Set({p1, p1})
            d_1027_v2_: _dafny.Map
            d_1027_v2_ = _dafny.Map({(self).f25: d_1026_v1_})
            d_1028_v3_: _dafny.Map
            d_1028_v3_ = _dafny.Map({(d_1026_v1_ if not(p2) else ((d_1027_v2_)[(self).f25] if ((self).f25) in (d_1027_v2_) else _dafny.Set({p1, default__.fm0(834, globalState), True, p2}))): len((_dafny.Map({(self).f25: p2})).set(len(d_1026_v1_), not(p1)))})
            index172_ = default__.safeIndex(990, (p3).length(0))
            rhs191_ = (self).f25
            rhs192_ = (d_1025_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctwwm")))
            rhs193_ = ((d_1028_v3_)[(d_1026_v1_) | (d_1026_v1_)] if ((d_1026_v1_) | (d_1026_v1_)) in (d_1028_v3_) else (self).f25)
            lhs136_ = globalState
            lhs137_ = p3
            lhs138_ = default__.safeIndex(990, (p3).length(0))
            r1 = rhs191_
            lhs136_.f10 = rhs192_
            lhs137_[lhs138_] = rhs193_
            (globalState).f12 = not(p1)
            d_1029_v4_: _dafny.Array
            def lambda59_(d_1030_i0_):
                return _dafny.CodePoint('e')

            init34_ = lambda59_
            nw157_ = _dafny.Array(None, 13)
            for i0_34_ in range(nw157_.length(0)):
                nw157_[i0_34_] = init34_(i0_34_)
            d_1029_v4_ = nw157_
            d_1031_v5_: str
            d_1031_v5_ = _dafny.CodePoint('a')
            index173_ = default__.safeIndex(652, (d_1029_v4_).length(0))
            (d_1029_v4_)[index173_] = d_1031_v5_
            index174_ = default__.safeIndex(652, (d_1029_v4_).length(0))
            (d_1029_v4_)[index174_] = d_1031_v5_
            d_1032_v6_: D0
            d_1032_v6_ = D0_DC0(p2)
            d_1033_v7_: _dafny.Map
            d_1033_v7_ = _dafny.Map({d_1032_v6_: p1})
            d_1034_v8_: _dafny.Seq
            d_1034_v8_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            d_1035_v9_: _dafny.Map
            d_1035_v9_ = _dafny.Map({((p3)[default__.safeIndex(990, (p3).length(0))]) + (default__.fm21(d_1033_v7_, -60, False, d_1025_v0_, globalState)): (d_1034_v8_)[default__.safeIndex((p3)[default__.safeIndex(990, (p3).length(0))], len(d_1034_v8_))]})
            d_1035_v9_ = _dafny.Map({(0) - ((p3)[default__.safeIndex(990, (p3).length(0))]): (p3)[default__.safeIndex(990, (p3).length(0))]})
        elif True:
            d_1036_v10_: _dafny.Seq
            d_1036_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bliyyoh"))
            (globalState).f10 = (default__.fm2((self).f25, (self).f25, globalState)) + ((d_1036_v10_) + (d_1036_v10_))
            d_1037_v11_: _dafny.Seq
            d_1037_v11_ = _dafny.SeqWithoutIsStrInference([default__.fm0((0) - ((self).f25), globalState), not (p2) or (p2)])
            (globalState).f12 = (d_1037_v11_)[default__.safeIndex((self).f25, len(d_1037_v11_))]
            if False:
                r0 = default__.safeModuloInt((self).f25, (self).f25)
                d_1038_v12_: _dafny.Set
                d_1038_v12_ = _dafny.Set({p2})
                d_1039_v13_: _dafny.Array
                nw158_ = _dafny.Array(None, 3)
                nw158_[int(0)] = default__.safeModuloInt((0) - (len(d_1038_v12_)), (self).f25)
                nw158_[int(1)] = (self).f25
                nw158_[int(2)] = (self).f25
                d_1039_v13_ = nw158_
                d_1040_v14_: D1
                d_1040_v14_ = D1_DC4(_dafny.SeqWithoutIsStrInference([p2]))
                d_1041_v15_: D1
                d_1041_v15_ = D1_DC7(d_1040_v14_)
                d_1042_v16_: _dafny.Seq
                d_1042_v16_ = _dafny.SeqWithoutIsStrInference([d_1041_v15_])
                rhs194_ = default__.safeModuloInt((0) - (len((_dafny.SeqWithoutIsStrInference([D1_DC7(D1_DC6(_dafny.CodePoint('o'), p2, d_1036_v10_)) for d_1043_i1_ in range(default__.abs(325))])) + (d_1042_v16_))), (self).f25)
                rhs195_ = p2
                rhs196_ = d_1039_v13_
                lhs139_ = globalState
                r1 = rhs194_
                lhs139_.f12 = rhs195_
                d_1039_v13_ = rhs196_
                d_1044_v17_: str
                d_1044_v17_ = _dafny.CodePoint('h')
                (globalState).f23 = d_1044_v17_
                (globalState).f12 = False
                (globalState).f12 = p1
            elif True:
                index175_ = default__.safeIndex(806, (p3).length(0))
                (p3)[index175_] = (self).f25
                index176_ = default__.safeIndex(806, (p3).length(0))
                (p3)[index176_] = (0) - ((self).f25)
                index177_ = default__.safeIndex(297, ((self).f26).length(0))
                ((self).f26)[index177_] = d_1036_v10_
                index178_ = default__.safeIndex(297, ((self).f26).length(0))
                ((self).f26)[index178_] = (d_1036_v10_) + (d_1036_v10_)
                (globalState).f23 = (d_1036_v10_)[default__.safeIndex(default__.safeDivisionInt((p3)[default__.safeIndex(806, (p3).length(0))], (self).f25), len(d_1036_v10_))]
                d_1045_v18_: _dafny.Map
                d_1045_v18_ = _dafny.Map({(p3)[default__.safeIndex(806, (p3).length(0))]: d_1037_v11_})
                d_1046_v19_: _dafny.Map
                d_1046_v19_ = _dafny.Map({d_1036_v10_: len((((d_1045_v18_)[(p3)[default__.safeIndex(806, (p3).length(0))]] if ((p3)[default__.safeIndex(806, (p3).length(0))]) in (d_1045_v18_) else (d_1037_v11_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kytjdmn"))), len(d_1037_v11_)), p2))) + (d_1037_v11_))})
                d_1046_v19_ = (d_1046_v19_).set(((d_1036_v10_) + (((self).f26)[default__.safeIndex(297, ((self).f26).length(0))])).set(default__.safeIndex((p3)[default__.safeIndex(806, (p3).length(0))], len((d_1036_v10_) + (((self).f26)[default__.safeIndex(297, ((self).f26).length(0))]))), _dafny.CodePoint('f')), default__.safeModuloInt((p3)[default__.safeIndex(806, (p3).length(0))], (p3)[default__.safeIndex(806, (p3).length(0))]))
                d_1047_v20_: _dafny.Array
                nw159_ = _dafny.Array(False, 29)
                d_1047_v20_ = nw159_
                d_1048_v21_: _dafny.Seq
                d_1048_v21_ = _dafny.SeqWithoutIsStrInference([d_1047_v20_])
                d_1049_v22_: str
                d_1049_v22_ = _dafny.CodePoint('v')
                index179_ = default__.safeIndex(297, ((self).f26).length(0))
                rhs197_ = (D1_DC6(d_1049_v22_, True, d_1036_v10_)).cf9
                rhs198_ = (d_1037_v11_)[default__.safeIndex(len(((self).f26)[default__.safeIndex(297, ((self).f26).length(0))]), len(d_1037_v11_))]
                rhs199_ = d_1049_v22_
                rhs200_ = d_1048_v21_
                lhs140_ = (self).f26
                lhs141_ = default__.safeIndex(297, ((self).f26).length(0))
                lhs142_ = globalState
                lhs143_ = globalState
                lhs140_[lhs141_] = rhs197_
                lhs142_.f12 = rhs198_
                lhs143_.f23 = rhs199_
                d_1048_v21_ = rhs200_
            if p1:
                d_1050_v23_: _dafny.Array
                nw160_ = _dafny.Array(_dafny.Map({}), 6)
                d_1050_v23_ = nw160_
                d_1051_v24_: _dafny.Map
                d_1051_v24_ = _dafny.Map({p1: 630})
                d_1052_v25_: _dafny.Map
                d_1052_v25_ = _dafny.Map({p2: d_1051_v24_})
                d_1053_v26_: _dafny.Seq
                d_1053_v26_ = _dafny.SeqWithoutIsStrInference([d_1052_v25_])
                index180_ = default__.safeIndex(375, (d_1050_v23_).length(0))
                (d_1050_v23_)[index180_] = (d_1053_v26_)[default__.safeIndex((self).f25, len(d_1053_v26_))]
                index181_ = default__.safeIndex(375, (d_1050_v23_).length(0))
                (d_1050_v23_)[index181_] = d_1052_v25_
                d_1054_v27_: _dafny.Array
                nw161_ = _dafny.Array(False, 8)
                d_1054_v27_ = nw161_
                index182_ = default__.safeIndex(848, (d_1054_v27_).length(0))
                (d_1054_v27_)[index182_] = not(p1)
                d_1055_v28_: _dafny.MultiSet
                d_1055_v28_ = _dafny.MultiSet([(self).f25])
                d_1056_v29_: _dafny.Map
                d_1056_v29_ = _dafny.Map({(self).f25: d_1055_v28_})
                d_1057_v30_: _dafny.Set
                d_1057_v30_ = _dafny.Set({p2, p1})
                d_1058_v31_: _dafny.Map
                d_1058_v31_ = _dafny.Map({(self).f25: (((d_1056_v29_)[len(d_1057_v30_)] if (len(d_1057_v30_)) in (d_1056_v29_) else d_1055_v28_)).cardinality})
                d_1059_v32_: _dafny.Map
                d_1059_v32_ = _dafny.Map({(self).f25: (len(d_1058_v31_)) + ((self).f25)})
                d_1060_v33_: _dafny.Seq
                d_1060_v33_ = _dafny.SeqWithoutIsStrInference([(self).f25, ((self).f25) + ((self).f25)])
                d_1061_v34_: _dafny.Map
                d_1061_v34_ = _dafny.Map({p1: p2})
                index183_ = default__.safeIndex(848, (d_1054_v27_).length(0))
                rhs201_ = not(p2)
                rhs202_ = d_1059_v32_
                rhs203_ = d_1060_v33_
                rhs204_ = ((len(d_1058_v31_)) - (len(d_1036_v10_))) * ((len(d_1051_v24_)) - (len(d_1061_v34_)))
                lhs144_ = d_1054_v27_
                lhs145_ = default__.safeIndex(848, (d_1054_v27_).length(0))
                lhs146_ = globalState
                lhs144_[lhs145_] = rhs201_
                d_1059_v32_ = rhs202_
                lhs146_.f6 = rhs203_
                r1 = rhs204_
                d_1062_v35_: _dafny.Map
                d_1062_v35_ = _dafny.Map({p1: d_1036_v10_})
                d_1062_v35_ = (d_1062_v35_) | (d_1062_v35_)
                (globalState).f12 = p2
                d_1063_v36_: D0
                d_1063_v36_ = D0_DC3((self).f25, False)
                (globalState).f12 = (d_1063_v36_).cf5
            elif True:
                d_1064_v37_: _dafny.Map
                d_1064_v37_ = _dafny.Map({p1: (p2) and (p2)})
                (globalState).f12 = ((d_1064_v37_)[((self).f25) != ((self).f25)] if (((self).f25) != ((self).f25)) in (d_1064_v37_) else p2)
                d_1065_v38_: _dafny.Set
                d_1065_v38_ = _dafny.Set({p3})
                d_1066_v39_: _dafny.Array
                nw162_ = _dafny.Array(None, 4)
                nw162_[int(0)] = (_dafny.Set({p3}) if False else _dafny.Set({p3, p3, p3}))
                nw162_[int(1)] = (d_1065_v38_) | (d_1065_v38_)
                nw162_[int(2)] = d_1065_v38_
                nw162_[int(3)] = d_1065_v38_
                d_1066_v39_ = nw162_
                index184_ = default__.safeIndex(281, (d_1066_v39_).length(0))
                (d_1066_v39_)[index184_] = (_dafny.Set({p3})).intersection(d_1065_v38_)
                index185_ = default__.safeIndex(281, (d_1066_v39_).length(0))
                (d_1066_v39_)[index185_] = d_1065_v38_
                d_1067_v40_: _dafny.Array
                nw163_ = _dafny.Array(_dafny.Map({}), 29)
                d_1067_v40_ = nw163_
                index186_ = default__.safeIndex(694, (d_1067_v40_).length(0))
                (d_1067_v40_)[index186_] = (_dafny.Map({p2: False})) | (d_1064_v37_)
                index187_ = default__.safeIndex(105, (p3).length(0))
                (p3)[index187_] = ((self).f25) + ((0) - ((self).f25))
                d_1068_v41_: _dafny.Seq
                d_1068_v41_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25])
                index188_ = default__.safeIndex(134, (p3).length(0))
                (p3)[index188_] = len(d_1068_v41_)
                d_1069_v42_: D0
                d_1069_v42_ = D0_DC0(p1)
                d_1070_v43_: _dafny.Map
                d_1070_v43_ = _dafny.Map({d_1069_v42_: p1})
                index189_ = default__.safeIndex(694, (d_1067_v40_).length(0))
                index190_ = default__.safeIndex(105, (p3).length(0))
                index191_ = default__.safeIndex(134, (p3).length(0))
                rhs205_ = d_1064_v37_
                rhs206_ = ((self).f25 if (d_1037_v11_)[default__.safeIndex(default__.fm21(d_1070_v43_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlkroi"))), p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aufslk")), globalState), len(d_1037_v11_))] else (self).f25)
                rhs207_ = (d_1036_v10_) + (d_1036_v10_)
                rhs208_ = len(d_1037_v11_)
                lhs147_ = d_1067_v40_
                lhs148_ = default__.safeIndex(694, (d_1067_v40_).length(0))
                lhs149_ = p3
                lhs150_ = default__.safeIndex(105, (p3).length(0))
                lhs151_ = globalState
                lhs152_ = p3
                lhs153_ = default__.safeIndex(134, (p3).length(0))
                lhs147_[lhs148_] = rhs205_
                lhs149_[lhs150_] = rhs206_
                lhs151_.f10 = rhs207_
                lhs152_[lhs153_] = rhs208_
                index192_ = default__.safeIndex(105, (p3).length(0))
                (p3)[index192_] = (default__.safeDivisionInt((p3)[default__.safeIndex(105, (p3).length(0))], (p3)[default__.safeIndex(105, (p3).length(0))])) * ((self).f25)
                d_1071_v44_: _dafny.Array
                nw164_ = _dafny.Array(_dafny.Map({}), 20)
                d_1071_v44_ = nw164_
                d_1072_v45_: _dafny.Map
                d_1072_v45_ = _dafny.Map({(self).f25: False})
                index193_ = default__.safeIndex(521, (d_1071_v44_).length(0))
                (d_1071_v44_)[index193_] = d_1072_v45_
                index194_ = default__.safeIndex(521, (d_1071_v44_).length(0))
                (d_1071_v44_)[index194_] = d_1072_v45_
            if not(not(((self).f25) < ((self).f25))):
                (globalState).f12 = (D0_DC2(p2)).cf3
                d_1073_v46_: _dafny.Array
                nw165_ = _dafny.Array(_dafny.MultiSet({}), 6)
                d_1073_v46_ = nw165_
                d_1074_v47_: _dafny.MultiSet
                d_1074_v47_ = _dafny.MultiSet([p2, p2])
                index195_ = default__.safeIndex(795, (d_1073_v46_).length(0))
                (d_1073_v46_)[index195_] = (_dafny.MultiSet([p2])) | (d_1074_v47_)
                index196_ = default__.safeIndex(795, (d_1073_v46_).length(0))
                rhs209_ = (((d_1074_v47_).set(p1, default__.abs(len(d_1036_v10_)))) | (d_1074_v47_)).set(default__.fm0((self).f25, globalState), default__.abs((self).f25))
                rhs210_ = not(p2)
                lhs154_ = d_1073_v46_
                lhs155_ = default__.safeIndex(795, (d_1073_v46_).length(0))
                lhs156_ = globalState
                lhs154_[lhs155_] = rhs209_
                lhs156_.f12 = rhs210_
                d_1075_v48_: D0
                d_1075_v48_ = D0_DC0(p2)
                d_1076_v49_: _dafny.Map
                d_1076_v49_ = _dafny.Map({d_1075_v48_: not(p1)})
                d_1077_v50_: _dafny.Map
                d_1077_v50_ = _dafny.Map({len(d_1036_v10_): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1078_i2_ in range(default__.abs(360))])})
                d_1079_v51_: str
                d_1079_v51_ = _dafny.CodePoint('w')
                d_1080_v52_: D1
                d_1080_v52_ = D1_DC6(d_1079_v51_, False, d_1036_v10_)
                d_1081_v53_: _dafny.Seq
                d_1081_v53_ = _dafny.SeqWithoutIsStrInference([(d_1080_v52_).cf9, d_1036_v10_])
                r0 = default__.fm21(d_1076_v49_, (self).f25, False, ((d_1077_v50_)[len((d_1081_v53_)[default__.safeIndex((self).f25, len(d_1081_v53_))])] if (len((d_1081_v53_)[default__.safeIndex((self).f25, len(d_1081_v53_))])) in (d_1077_v50_) else d_1036_v10_), globalState)
                (globalState).f12 = ((0) - (((self).f25) * ((self).f25))) != (-63)
                d_1082_v54_: _dafny.Map
                d_1082_v54_ = _dafny.Map({False: (self).f36})
                d_1083_v55_: _dafny.Set
                d_1083_v55_ = _dafny.Set({(self).f25, (self).f25})
                d_1082_v54_ = (d_1082_v54_).set((d_1083_v55_).issubset(d_1083_v55_), (self).f26)
            elif True:
                d_1084_v56_: _dafny.Array
                nw166_ = _dafny.Array(_dafny.MultiSet({}), 3)
                d_1084_v56_ = nw166_
                d_1085_v57_: _dafny.Array
                nw167_ = _dafny.Array(False, 25)
                d_1085_v57_ = nw167_
                d_1086_v58_: _dafny.MultiSet
                d_1086_v58_ = _dafny.MultiSet([d_1085_v57_, d_1085_v57_])
                index197_ = default__.safeIndex(551, (d_1084_v56_).length(0))
                (d_1084_v56_)[index197_] = (d_1086_v58_).set(d_1085_v57_, default__.abs(len(d_1037_v11_)))
                d_1087_v59_: _dafny.MultiSet
                d_1087_v59_ = _dafny.MultiSet([p2])
                index198_ = default__.safeIndex(551, (d_1084_v56_).length(0))
                (d_1084_v56_)[index198_] = ((d_1086_v58_ if p1 else d_1086_v58_)).set(d_1085_v57_, default__.abs(len(_dafny.Set({(d_1087_v59_).cardinality}))))
                d_1088_v60_: _dafny.Array
                def lambda60_(d_1089_p2_, d_1090_p1_):
                    def lambda61_(d_1091_i3_):
                        return D4_DC13((self).f25, d_1089_p2_, d_1090_p1_, True)

                    return lambda61_

                init35_ = lambda60_(p2, p1)
                nw168_ = _dafny.Array(None, 18)
                for i0_35_ in range(nw168_.length(0)):
                    nw168_[i0_35_] = init35_(i0_35_)
                d_1088_v60_ = nw168_
                d_1092_v61_: D4
                d_1092_v61_ = D4_DC13((self).f25, p1, p2, p2)
                index199_ = default__.safeIndex(861, (d_1088_v60_).length(0))
                (d_1088_v60_)[index199_] = d_1092_v61_
                d_1093_v62_: _dafny.MultiSet
                d_1093_v62_ = _dafny.MultiSet([(_dafny.MultiSet([p1])).cardinality, (self).f25])
                d_1094_v63_: _dafny.Set
                d_1094_v63_ = _dafny.Set({(_dafny.MultiSet(d_1037_v11_)).cardinality, (d_1093_v62_).cardinality, (self).f25, (self).f25})
                index200_ = default__.safeIndex(861, (d_1088_v60_).length(0))
                (d_1088_v60_)[index200_] = default__.fm22(d_1094_v63_, (self).f25, globalState)
                d_1095_v64_: _dafny.Array
                nw169_ = _dafny.Array(_dafny.Set({}), 11)
                d_1095_v64_ = nw169_
                d_1096_v65_: _dafny.Set
                d_1096_v65_ = _dafny.Set({p2})
                index201_ = default__.safeIndex(288, (d_1095_v64_).length(0))
                (d_1095_v64_)[index201_] = d_1096_v65_
                index202_ = default__.safeIndex(288, (d_1095_v64_).length(0))
                (d_1095_v64_)[index202_] = (d_1096_v65_).intersection((d_1096_v65_) - (d_1096_v65_))
                d_1097_v66_: _dafny.Seq
                d_1097_v66_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                d_1098_v67_: _dafny.Map
                d_1098_v67_ = _dafny.Map({(d_1097_v66_ if p2 else _dafny.SeqWithoutIsStrInference([(self).f25 for d_1099_i4_ in range(default__.abs(-301))])): p2})
                d_1098_v67_ = ((d_1098_v67_) | (d_1098_v67_)) | ((_dafny.Map({d_1097_v66_: False})) | (d_1098_v67_))
                d_1100_v68_: _dafny.Map
                d_1100_v68_ = _dafny.Map({p2: (self).f25})
                d_1101_v69_: D0
                d_1102_v70_: bool
                d_1103_v71_: _dafny.Map
                out28_: D0
                out29_: bool
                out30_: _dafny.Map
                out28_, out29_, out30_ = (self).m15(not (p1) or (not(p2)), d_1100_v68_, len(d_1036_v10_), globalState)
                d_1101_v69_ = out28_
                d_1102_v70_ = out29_
                d_1103_v71_ = out30_
        d_1104_v72_: D0
        d_1104_v72_ = D0_DC2(False)
        d_1105_v73_: _dafny.Map
        d_1105_v73_ = _dafny.Map({p2: d_1104_v72_})
        d_1106_v74_: _dafny.Map
        d_1106_v74_ = _dafny.Map({(self).f25: (self).f25})
        hi5_ = ((d_1106_v74_)[(0) - ((self).f25)] if ((0) - ((self).f25)) in (d_1106_v74_) else (self).f25)
        for d_1107_i5_ in range(len(d_1105_v73_), hi5_):
            d_1108_v75_: _dafny.Seq
            d_1108_v75_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1109_v76_: D1
            d_1109_v76_ = D1_DC4(d_1108_v75_)
            d_1110_v77_: _dafny.Set
            d_1110_v77_ = _dafny.Set({d_1109_v76_, D1_DC4(_dafny.SeqWithoutIsStrInference([p1])), d_1109_v76_})
            d_1110_v77_ = ((d_1110_v77_) - (d_1110_v77_)) | ((d_1110_v77_) - (d_1110_v77_))
            d_1111_v78_: D4
            d_1111_v78_ = D4_DC12()
            rhs211_ = (0) - ((self).f25)
            rhs212_ = d_1111_v78_
            r0 = rhs211_
            d_1111_v78_ = rhs212_
            if True:
                d_1112_v79_: _dafny.Seq
                d_1112_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opd"))
                r0 = (d_1107_i5_) + (len((d_1112_v79_) + (d_1112_v79_)))
                index203_ = default__.safeIndex(770, (p3).length(0))
                (p3)[index203_] = (self).f25
                index204_ = default__.safeIndex(770, (p3).length(0))
                (p3)[index204_] = ((self).f25) - (816)
                d_1113_v80_: _dafny.Set
                d_1113_v80_ = _dafny.Set({(self).f25})
                d_1114_v81_: _dafny.Map
                d_1114_v81_ = _dafny.Map({p1: not(p1)})
                d_1115_v82_: _dafny.Seq
                d_1115_v82_ = _dafny.SeqWithoutIsStrInference([(p3)[default__.safeIndex(770, (p3).length(0))], len(d_1114_v81_), len(d_1113_v80_), len(d_1108_v75_), d_1107_i5_])
                index205_ = default__.safeIndex(770, (p3).length(0))
                rhs213_ = not(p1)
                rhs214_ = default__.safeModuloInt(len(d_1112_v79_), ((p3)[default__.safeIndex(770, (p3).length(0))]) + ((0) - (len(d_1113_v80_))))
                rhs215_ = (-232 if p1 else (d_1115_v82_)[default__.safeIndex(175, len(d_1115_v82_))])
                lhs157_ = globalState
                lhs158_ = p3
                lhs159_ = default__.safeIndex(770, (p3).length(0))
                lhs157_.f12 = rhs213_
                r0 = rhs214_
                lhs158_[lhs159_] = rhs215_
                d_1116_v83_: D0
                d_1116_v83_ = D0_DC0(p1)
                d_1117_v84_: str
                d_1117_v84_ = _dafny.CodePoint('d')
                d_1118_v85_: _dafny.MultiSet
                d_1118_v85_ = _dafny.MultiSet([default__.fm0(166, globalState)])
                d_1119_v86_: _dafny.Map
                d_1119_v86_ = _dafny.Map({(p3)[default__.safeIndex(770, (p3).length(0))]: p2})
                r0 = default__.safeDivisionInt((default__.fm21((_dafny.Map({d_1116_v83_: default__.fm0((self).f25, globalState)})).set(d_1116_v83_, p1), d_1107_i5_, False, ((D1_DC6(_dafny.CodePoint('c'), p2, d_1112_v79_)).cf9).set(default__.safeIndex((self).f25, len((D1_DC6(_dafny.CodePoint('c'), p2, d_1112_v79_)).cf9)), d_1117_v84_), globalState)) * (-305), ((d_1118_v85_)[((d_1119_v86_)[(self).f25] if ((self).f25) in (d_1119_v86_) else p1)] if (((d_1119_v86_)[(self).f25] if ((self).f25) in (d_1119_v86_) else p1)) in (d_1118_v85_) else d_1107_i5_))
                (globalState).f12 = (d_1107_i5_) >= (934)
            elif True:
                d_1120_v87_: str
                d_1120_v87_ = _dafny.CodePoint('f')
                d_1121_v88_: _dafny.Map
                d_1121_v88_ = _dafny.Map({d_1120_v87_: (self).f25})
                d_1122_v89_: _dafny.Set
                d_1122_v89_ = _dafny.Set({(self).f25})
                (globalState).f12 = (default__.safeModuloInt(len(d_1121_v88_), len(d_1122_v89_))) >= (d_1107_i5_)
                d_1123_v90_: _dafny.Seq
                d_1123_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtm"))
                d_1124_v91_: _dafny.MultiSet
                d_1124_v91_ = _dafny.MultiSet([False, p1, (d_1123_v90_) <= (d_1123_v90_)])
                d_1125_v92_: _dafny.Set
                d_1125_v92_ = _dafny.Set({p1, p1})
                d_1126_v93_: _dafny.Set
                d_1126_v93_ = _dafny.Set({d_1120_v87_})
                r1 = ((d_1124_v91_)[(d_1125_v92_).issubset(d_1125_v92_)] if ((d_1125_v92_).issubset(d_1125_v92_)) in (d_1124_v91_) else len(d_1126_v93_))
                (globalState).f12 = (p2 if not (p1) or (not(default__.fm0((self).f25, globalState))) else p2)
                d_1127_v94_: D0
                d_1127_v94_ = D0_DC0(p1)
                pat_let_tv20_ = p1
                d_1128_v95_: _dafny.Map
                def iife83_(_pat_let18_0):
                    def iife84_(d_1129_dt__update__tmp_h0_):
                        def iife85_(_pat_let19_0):
                            def iife86_(d_1130_dt__update_hcf0_h0_):
                                return D0_DC0(d_1130_dt__update_hcf0_h0_)
                            return iife86_(_pat_let19_0)
                        return iife85_(pat_let_tv20_)
                    return iife84_(_pat_let18_0)
                d_1128_v95_ = _dafny.Map({iife83_(d_1127_v94_): p1})
                (globalState).f12 = ((self).f25) != (default__.safeDivisionInt((self).f25, default__.fm21(d_1128_v95_, 555, p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbb")), globalState)))
                d_1108_v75_ = d_1108_v75_
            d_1131_v96_: D0
            d_1131_v96_ = D0_DC0(p2)
            d_1132_v97_: _dafny.Map
            d_1132_v97_ = _dafny.Map({d_1131_v96_: p1})
            d_1133_v98_: _dafny.Seq
            d_1133_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjlaotpx"))
            d_1134_v99_: _dafny.Array
            def lambda62_(d_1135_p1_):
                def lambda63_(d_1136_i6_):
                    return d_1135_p1_

                return lambda63_

            init36_ = lambda62_(p1)
            nw170_ = _dafny.Array(None, 20)
            for i0_36_ in range(nw170_.length(0)):
                nw170_[i0_36_] = init36_(i0_36_)
            d_1134_v99_ = nw170_
            d_1137_v100_: _dafny.Seq
            d_1137_v100_ = _dafny.SeqWithoutIsStrInference([d_1134_v99_, d_1134_v99_, d_1134_v99_])
            d_1138_v101_: D4
            d_1138_v101_ = D4_DC13((default__.fm21(d_1132_v97_, len(d_1133_v98_), p1, d_1133_v98_, globalState)) * (len(d_1108_v75_)), False, default__.fm0(((d_1106_v74_)[len(d_1137_v100_)] if (len(d_1137_v100_)) in (d_1106_v74_) else 546), globalState), not(not(p2)))
            d_1138_v101_ = d_1138_v101_
        (globalState).f12 = (p2) or ((p1) and (p1))
        d_1139_v102_: _dafny.Array
        nw171_ = _dafny.Array(False, 8)
        d_1139_v102_ = nw171_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1139_v102_).length(0)):
            d_1140_i7_: int = guard_loop_4_
            if (True) and (((0) <= (d_1140_i7_)) and ((d_1140_i7_) < ((d_1139_v102_).length(0)))):
                def iife87_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(36, 152):
                        d_1141_v103_: int = compr_47_
                        if ((36) <= (d_1141_v103_)) and ((d_1141_v103_) < (152)):
                            coll47_ = coll47_.union(_dafny.Set([default__.safeModuloInt(d_1141_v103_, (0) - ((self).f25))]))
                    return _dafny.Set(coll47_)
                (d_1139_v102_)[(d_1140_i7_)] = ((_dafny.Set({(self).f25, (0) - (len(_dafny.Map({p2: (self).f25})))})) - (iife87_()
                )).issubset((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1142_i8_ in range(default__.abs(205))]))})) - (_dafny.Set({(D4_DC13(len(d_1106_v74_), p1, p1, not(p2))).cf19, (self).f25})))
        d_1143_v104_: D0
        d_1143_v104_ = D0_DC0(not(True))
        d_1144_v105_: _dafny.Map
        d_1144_v105_ = _dafny.Map({d_1143_v104_: default__.fm0((self).f25, globalState)})
        d_1145_v106_: _dafny.MultiSet
        d_1145_v106_ = _dafny.MultiSet([(self).f25])
        d_1146_v107_: _dafny.Set
        d_1146_v107_ = _dafny.Set({d_1145_v106_})
        d_1147_v108_: _dafny.Seq
        d_1147_v108_ = _dafny.SeqWithoutIsStrInference([default__.fm21(d_1144_v105_, len(d_1146_v107_), p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgfs")), globalState)])
        d_1148_v109_: _dafny.Seq
        d_1148_v109_ = d_1147_v108_
        pat_let_tv21_ = p2
        def lambda64_(source19_):
            d_1149___mcc_h0_ = source19_
            d_1150_cf11_ = d_1149___mcc_h0_
            return pat_let_tv21_

        (globalState).f12 = lambda64_(d_1148_v109_)
        r0 = (self).f25
        d_1151_v110_: D1
        d_1151_v110_ = D1_DC4(_dafny.SeqWithoutIsStrInference([not(p1)]))
        d_1152_v111_: _dafny.MultiSet
        d_1152_v111_ = _dafny.MultiSet([d_1151_v110_])
        r1 = default__.safeDivisionInt(((d_1152_v111_)[d_1151_v110_] if (d_1151_v110_) in (d_1152_v111_) else (self).f25), (self).f25)
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        d_1153_v0_: _dafny.Map
        d_1153_v0_ = _dafny.Map({(0) - (((self).f25) + ((self).f25)): p1})
        if ((d_1153_v0_)[(self).f25] if ((self).f25) in (d_1153_v0_) else True):
            d_1154_v1_: _dafny.Seq
            d_1154_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            if not((d_1154_v1_) <= (d_1154_v1_)):
                r0 = (self).f25
                d_1155_v2_: C2
                nw172_ = C2()
                nw172_.ctor__()
                d_1155_v2_ = nw172_
                d_1156_v3_: _dafny.Set
                d_1156_v3_ = _dafny.Set({(self).f25, (self).f25, (self).f25})
                (globalState).f12 = (d_1156_v3_).isdisjoint((d_1156_v3_).intersection(d_1156_v3_))
                d_1157_v4_: _dafny.Array
                def lambda65_(d_1158_p1_):
                    def lambda66_(d_1159_i0_):
                        return d_1158_p1_

                    return lambda66_

                init37_ = lambda65_(p1)
                nw173_ = _dafny.Array(None, 4)
                for i0_37_ in range(nw173_.length(0)):
                    nw173_[i0_37_] = init37_(i0_37_)
                d_1157_v4_ = nw173_
                d_1160_v5_: _dafny.Map
                d_1160_v5_ = _dafny.Map({d_1157_v4_: (self).f25})
                d_1161_v6_: D6
                d_1161_v6_ = D6_DC18(d_1160_v5_)
                rhs216_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lk"))
                rhs217_ = (((d_1161_v6_).cf28) | (d_1160_v5_)) != (_dafny.Map({d_1157_v4_: (self).f25}))
                lhs160_ = globalState
                d_1154_v1_ = rhs216_
                lhs160_.f12 = rhs217_
                (globalState).f12 = default__.fm0((self).f25, globalState)
            elif True:
                d_1162_v7_: C1
                nw174_ = C1()
                nw174_.ctor__((self).f25, (self).f26)
                d_1162_v7_ = nw174_
                d_1163_v8_: _dafny.Seq
                d_1163_v8_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1164_v9_: _dafny.Map
                d_1164_v9_ = _dafny.Map({(d_1163_v8_)[default__.safeIndex((self).f25, len(d_1163_v8_))]: not(True)})
                d_1165_v10_: _dafny.Seq
                d_1165_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: p1})])
                d_1166_v11_: _dafny.Map
                d_1166_v11_ = _dafny.Map({d_1164_v9_: len(d_1165_v10_)})
                r0 = ((d_1166_v11_)[d_1164_v9_] if (d_1164_v9_) in (d_1166_v11_) else (self).f25)
                r0 = (self).f25
                d_1167_v12_: D0
                d_1167_v12_ = D0_DC3(len(d_1164_v9_), not(p1))
                d_1168_v13_: D6
                d_1168_v13_ = D6_DC19(d_1167_v12_, (self).f25)
                d_1169_v14_: C1
                nw175_ = C1()
                nw175_.ctor__(((0) - ((self).f25)) + ((d_1168_v13_).cf30), (self).f36)
                d_1169_v14_ = nw175_
                r0 = ((self).f25) + ((self).f25)
            d_1170_v15_: _dafny.Set
            d_1170_v15_ = _dafny.Set({p1, p1, p1, p1})
            d_1171_v16_: _dafny.Map
            d_1171_v16_ = _dafny.Map({((self).f25) > ((self).f25): d_1170_v15_})
            d_1171_v16_ = _dafny.Map({p1: d_1170_v15_})
            d_1172_v17_: _dafny.Array
            def lambda67_(d_1173_p1_):
                def lambda68_(d_1174_i1_):
                    return not(d_1173_p1_)

                return lambda68_

            init38_ = lambda67_(p1)
            nw176_ = _dafny.Array(None, 14)
            for i0_38_ in range(nw176_.length(0)):
                nw176_[i0_38_] = init38_(i0_38_)
            d_1172_v17_ = nw176_
            d_1172_v17_ = d_1172_v17_
            rhs218_ = ((self).f25) >= ((0) - ((self).f25))
            rhs219_ = d_1172_v17_
            lhs161_ = globalState
            lhs161_.f12 = rhs218_
            d_1172_v17_ = rhs219_
            index206_ = default__.safeIndex(63, (d_1172_v17_).length(0))
            (d_1172_v17_)[index206_] = True
            index207_ = default__.safeIndex(63, (d_1172_v17_).length(0))
            (d_1172_v17_)[index207_] = True
        elif True:
            d_1175_v18_: _dafny.Set
            d_1175_v18_ = _dafny.Set({p1, p1})
            d_1176_v19_: str
            d_1176_v19_ = _dafny.CodePoint('f')
            d_1177_v20_: _dafny.Map
            d_1177_v20_ = _dafny.Map({(self).f25: d_1176_v19_})
            d_1178_v21_: _dafny.MultiSet
            d_1178_v21_ = _dafny.MultiSet([d_1177_v20_, d_1177_v20_, d_1177_v20_, d_1177_v20_, d_1177_v20_])
            d_1179_v22_: _dafny.Array
            nw177_ = _dafny.Array(None, 16)
            nw177_[int(0)] = not(default__.fm0((self).f25, globalState))
            nw177_[int(1)] = not(p1)
            nw177_[int(2)] = p1
            nw177_[int(3)] = default__.fm0((self).f25, globalState)
            nw177_[int(4)] = p1
            nw177_[int(5)] = (_dafny.Set({default__.fm0((self).f25, globalState)})).ispropersubset(d_1175_v18_)
            nw177_[int(6)] = (p1) and (((d_1153_v0_)[(self).f25] if ((self).f25) in (d_1153_v0_) else p1))
            nw177_[int(7)] = (d_1178_v21_) == (d_1178_v21_)
            nw177_[int(8)] = p1
            nw177_[int(9)] = p1
            nw177_[int(10)] = p1
            nw177_[int(11)] = p1
            nw177_[int(12)] = ((d_1153_v0_)[(self).f25] if ((self).f25) in (d_1153_v0_) else not(True))
            nw177_[int(13)] = not(False)
            nw177_[int(14)] = False
            nw177_[int(15)] = p1
            d_1179_v22_ = nw177_
            d_1180_v23_: _dafny.Seq
            d_1180_v23_ = _dafny.SeqWithoutIsStrInference([d_1179_v22_])
            index208_ = default__.safeIndex(538, (d_1179_v22_).length(0))
            (d_1179_v22_)[index208_] = not((d_1180_v23_) <= (d_1180_v23_))
            d_1181_v24_: _dafny.Seq
            d_1181_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqpkwm"))
            d_1182_v25_: _dafny.MultiSet
            d_1182_v25_ = _dafny.MultiSet([(self).f25, (0) - ((self).f25)])
            d_1183_v26_: _dafny.MultiSet
            d_1183_v26_ = _dafny.MultiSet([d_1182_v25_, _dafny.MultiSet([(self).f25])])
            index209_ = default__.safeIndex(538, (d_1179_v22_).length(0))
            rhs220_ = (d_1181_v24_) <= (d_1181_v24_)
            rhs221_ = len(default__.fm37(d_1183_v26_, globalState))
            lhs162_ = d_1179_v22_
            lhs163_ = default__.safeIndex(538, (d_1179_v22_).length(0))
            lhs162_[lhs163_] = rhs220_
            r0 = rhs221_
            d_1184_v27_: _dafny.Seq
            d_1184_v27_ = _dafny.SeqWithoutIsStrInference([d_1181_v24_])
            d_1181_v24_ = default__.fm2(len(((d_1184_v27_).set(default__.safeIndex((self).f25, len(d_1184_v27_)), d_1181_v24_)) + (d_1184_v27_)), (self).f25, globalState)
            index210_ = default__.safeIndex(538, (d_1179_v22_).length(0))
            rhs222_ = d_1181_v24_
            rhs223_ = p1
            lhs164_ = d_1179_v22_
            lhs165_ = default__.safeIndex(538, (d_1179_v22_).length(0))
            d_1181_v24_ = rhs222_
            lhs164_[lhs165_] = rhs223_
            r0 = default__.safeDivisionInt(default__.safeModuloInt((0) - ((self).f25), (0) - (-231)), (len(d_1153_v0_)) + ((self).f25))
            d_1185_v28_: _dafny.Array
            nw178_ = _dafny.Array(int(0), 2)
            d_1185_v28_ = nw178_
            index211_ = default__.safeIndex(760, (d_1185_v28_).length(0))
            (d_1185_v28_)[index211_] = (self).f25
            index212_ = default__.safeIndex(760, (d_1185_v28_).length(0))
            (d_1185_v28_)[index212_] = (855) - ((self).f25)
        r0 = (self).f25
        d_1186_v29_: _dafny.Array
        def lambda69_(d_1187_p1_):
            def lambda70_(d_1188_i2_):
                return (_dafny.Map({d_1187_p1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "embe"))})) | (_dafny.Map({not(d_1187_p1_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvjmfhrm"))}))

            return lambda70_

        init39_ = lambda69_(p1)
        nw179_ = _dafny.Array(None, 4)
        for i0_39_ in range(nw179_.length(0)):
            nw179_[i0_39_] = init39_(i0_39_)
        d_1186_v29_ = nw179_
        d_1189_v30_: _dafny.Seq
        d_1189_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awybtal"))
        d_1190_v31_: _dafny.Map
        d_1190_v31_ = _dafny.Map({p1: d_1189_v30_})
        index213_ = default__.safeIndex(351, (d_1186_v29_).length(0))
        (d_1186_v29_)[index213_] = d_1190_v31_
        index214_ = default__.safeIndex(351, (d_1186_v29_).length(0))
        (d_1186_v29_)[index214_] = ((d_1190_v31_) | (d_1190_v31_)) | ((d_1190_v31_) | (d_1190_v31_))
        hi6_ = ((self).f25) * ((self).f25)
        for d_1191_i3_ in range((self).f25, hi6_):
            d_1192_v32_: _dafny.MultiSet
            d_1192_v32_ = _dafny.MultiSet([len(d_1189_v30_)])
            d_1193_v33_: _dafny.Seq
            d_1193_v33_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1194_v34_: _dafny.Map
            d_1194_v34_ = _dafny.Map({((d_1153_v0_)[d_1191_i3_] if (d_1191_i3_) in (d_1153_v0_) else default__.fm0(len(d_1193_v33_), globalState)): p1})
            d_1195_v35_: D0
            d_1195_v35_ = D0_DC3(((d_1192_v32_)[d_1191_i3_] if (d_1191_i3_) in (d_1192_v32_) else len(d_1194_v34_)), p1)
            d_1195_v35_ = d_1195_v35_
            d_1196_v36_: _dafny.Array
            def lambda71_(d_1197_p1_, d_1198_i3_):
                def lambda72_(d_1199_i4_):
                    return (_dafny.Map({d_1197_p1_: (D5_DC17(d_1198_i3_, d_1197_p1_)).cf26}) if d_1197_p1_ else _dafny.Map({d_1197_p1_: (self).f25}))

                return lambda72_

            init40_ = lambda71_(p1, d_1191_i3_)
            nw180_ = _dafny.Array(None, 16)
            for i0_40_ in range(nw180_.length(0)):
                nw180_[i0_40_] = init40_(i0_40_)
            d_1196_v36_ = nw180_
            nw181_ = _dafny.Array(_dafny.Map({}), 6)
            d_1196_v36_ = nw181_
            (globalState).f12 = (p1 if p1 else not(p1))
            d_1200_v37_: D0
            d_1200_v37_ = D0_DC0(p1)
            d_1201_v38_: _dafny.Map
            d_1201_v38_ = _dafny.Map({d_1200_v37_: p1})
            d_1202_v39_: _dafny.Seq
            d_1202_v39_ = _dafny.SeqWithoutIsStrInference([default__.fm21(d_1201_v38_, d_1191_i3_, default__.fm0(len(d_1189_v30_), globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "noqaw")), globalState), (self).f25, d_1191_i3_])
            d_1203_v40_: _dafny.Map
            d_1203_v40_ = _dafny.Map({d_1202_v39_: d_1191_i3_})
            r0 = ((d_1203_v40_)[d_1202_v39_] if (d_1202_v39_) in (d_1203_v40_) else ((self).f25) * ((0) - (len(_dafny.Map({p1: d_1191_i3_})))))
        d_1204_v41_: T0
        nw182_ = C3()
        nw182_.ctor__((self).f25, p1, (self).f25, (self).f36)
        d_1204_v41_ = nw182_
        d_1204_v41_ = d_1204_v41_
        d_1205_v42_: str
        d_1205_v42_ = _dafny.CodePoint('y')
        d_1206_v43_: D1
        d_1206_v43_ = D1_DC6(d_1205_v42_, p1, d_1189_v30_)
        d_1207_i5_: int
        d_1207_i5_ = 0
        with _dafny.label("10"):
            while (D1_DC6(d_1205_v42_, p1, (d_1206_v43_).cf9)).cf8:
                with _dafny.c_label("10"):
                    if (d_1207_i5_) >= (100):
                        raise _dafny.Break("10")
                    d_1207_i5_ = (d_1207_i5_) + (1)
                    d_1208_v44_: C3
                    nw183_ = C3()
                    nw183_.ctor__((self).f25, p1, (0) - ((self).f25), (d_1204_v41_).f26)
                    d_1208_v44_ = nw183_
                    d_1209_v45_: _dafny.Array
                    def lambda73_(d_1210_p1_):
                        def lambda74_(d_1211_i6_):
                            return _dafny.MultiSet([d_1210_p1_])

                        return lambda74_

                    init41_ = lambda73_(p1)
                    nw184_ = _dafny.Array(None, 3)
                    for i0_41_ in range(nw184_.length(0)):
                        nw184_[i0_41_] = init41_(i0_41_)
                    d_1209_v45_ = nw184_
                    d_1209_v45_ = d_1209_v45_
                    d_1212_v46_: _dafny.MultiSet
                    d_1212_v46_ = _dafny.MultiSet([(d_1208_v44_).f38])
                    d_1213_v47_: _dafny.Map
                    d_1213_v47_ = _dafny.Map({d_1212_v46_: p1})
                    d_1214_v49_: _dafny.Map
                    def iife88_():
                        coll48_ = _dafny.Map()
                        compr_48_: _dafny.MultiSet
                        for compr_48_ in (d_1213_v47_).keys.Elements:
                            d_1215_v48_: _dafny.MultiSet = compr_48_
                            if (d_1215_v48_) in (d_1213_v47_):
                                coll48_[d_1215_v48_] = p1
                        return _dafny.Map(coll48_)
                    d_1214_v49_ = _dafny.Map({not(default__.fm0((d_1204_v41_).f25, globalState)): (d_1213_v47_) | (iife88_()
                    )})
                    d_1214_v49_ = (d_1214_v49_).set(((d_1208_v44_).f37) < (754), _dafny.Map({d_1212_v46_: (d_1208_v44_).f38}))
                    d_1216_v50_: _dafny.Array
                    nw185_ = _dafny.Array(False, 3)
                    d_1216_v50_ = nw185_
                    index215_ = default__.safeIndex(595, (d_1216_v50_).length(0))
                    (d_1216_v50_)[index215_] = (d_1208_v44_).f38
                    index216_ = default__.safeIndex(595, (d_1216_v50_).length(0))
                    (d_1216_v50_)[index216_] = not((d_1208_v44_).f38)
                    pass
            pass
        d_1217_v51_: _dafny.Array
        nw186_ = _dafny.Array(False, 28)
        d_1217_v51_ = nw186_
        d_1218_v52_: _dafny.MultiSet
        d_1218_v52_ = _dafny.MultiSet([p1, default__.fm0((_dafny.MultiSet([d_1217_v51_])).cardinality, globalState)])
        r0 = ((d_1218_v52_)[p1] if (p1) in (d_1218_v52_) else 38)
        return r0

    def m15(self, p0, p1, p2, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_1219_v0_: int
        d_1219_v0_ = 227
        d_1220_v1_: D0
        d_1220_v1_ = D0_DC0(p0)
        d_1221_v2_: _dafny.Seq
        d_1221_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
        d_1222_v3_: _dafny.Array
        nw187_ = _dafny.Array(None, 25)
        nw187_[int(0)] = D0_DC0(p0)
        nw187_[int(1)] = D0_DC0(p0)
        nw187_[int(2)] = d_1220_v1_
        nw187_[int(3)] = d_1220_v1_
        nw187_[int(4)] = d_1220_v1_
        nw187_[int(5)] = D0_DC0(not(p0))
        nw187_[int(6)] = d_1220_v1_
        nw187_[int(7)] = d_1220_v1_
        nw187_[int(8)] = d_1220_v1_
        nw187_[int(9)] = default__.fm53(p0, p2, p0, globalState)
        nw187_[int(10)] = d_1220_v1_
        nw187_[int(11)] = d_1220_v1_
        nw187_[int(12)] = (d_1220_v1_ if p0 else d_1220_v1_)
        nw187_[int(13)] = d_1220_v1_
        nw187_[int(14)] = d_1220_v1_
        nw187_[int(15)] = d_1220_v1_
        nw187_[int(16)] = d_1220_v1_
        nw187_[int(17)] = D0_DC0(p0)
        nw187_[int(18)] = D0_DC0(True)
        nw187_[int(19)] = d_1220_v1_
        nw187_[int(20)] = d_1220_v1_
        nw187_[int(21)] = d_1220_v1_
        nw187_[int(22)] = d_1220_v1_
        nw187_[int(23)] = (D0_DC0(p0) if not((default__.fm45(len(d_1221_v2_), d_1219_v0_, globalState)).cf25) else d_1220_v1_)
        nw187_[int(24)] = D0_DC0(p0)
        d_1222_v3_ = nw187_
        pat_let_tv22_ = p0
        pat_let_tv23_ = p0
        index217_ = default__.safeIndex(3, (d_1222_v3_).length(0))
        def iife90_(_pat_let21_0):
            def iife91_(d_1223_dt__update__tmp_h0_):
                def iife92_(_pat_let22_0):
                    def iife93_(d_1224_dt__update_hcf0_h0_):
                        return D0_DC0(d_1224_dt__update_hcf0_h0_)
                    return iife93_(_pat_let22_0)
                return iife92_(pat_let_tv22_)
            return iife91_(_pat_let21_0)
        def iife89_(_pat_let20_0):
            def iife94_(d_1225_dt__update__tmp_h1_):
                def iife95_(_pat_let23_0):
                    def iife96_(d_1226_dt__update_hcf0_h1_):
                        return D0_DC0(d_1226_dt__update_hcf0_h1_)
                    return iife96_(_pat_let23_0)
                return iife95_(pat_let_tv23_)
            return iife94_(_pat_let20_0)
        (d_1222_v3_)[index217_] = iife89_(iife90_(D0_DC0(p0)))
        d_1227_v4_: _dafny.Map
        d_1227_v4_ = _dafny.Map({d_1221_v2_: not(p0)})
        d_1228_v5_: _dafny.Seq
        d_1228_v5_ = _dafny.SeqWithoutIsStrInference([d_1221_v2_])
        index218_ = default__.safeIndex(3, (d_1222_v3_).length(0))
        rhs224_ = p0
        rhs225_ = p0
        rhs226_ = d_1219_v0_
        rhs227_ = d_1220_v1_
        rhs228_ = ((d_1227_v4_)[(d_1228_v5_)[default__.safeIndex(p2, len(d_1228_v5_))]] if ((d_1228_v5_)[default__.safeIndex(p2, len(d_1228_v5_))]) in (d_1227_v4_) else ((self).f25) == ((0) - (p2)))
        lhs166_ = globalState
        lhs167_ = d_1222_v3_
        lhs168_ = default__.safeIndex(3, (d_1222_v3_).length(0))
        lhs166_.f12 = rhs224_
        r1 = rhs225_
        d_1219_v0_ = rhs226_
        lhs167_[lhs168_] = rhs227_
        r1 = rhs228_
        d_1229_v6_: str
        d_1229_v6_ = _dafny.CodePoint('d')
        d_1221_v2_ = (d_1221_v2_).set(default__.safeIndex(-86, len(d_1221_v2_)), d_1229_v6_)
        d_1230_v7_: D1
        d_1230_v7_ = D1_DC6(d_1229_v6_, False, d_1221_v2_)
        d_1231_v8_: _dafny.Array
        def lambda75_(d_1232_p0_):
            def lambda76_(d_1233_i0_):
                return d_1232_p0_

            return lambda76_

        init42_ = lambda75_(p0)
        nw188_ = _dafny.Array(None, 24)
        for i0_42_ in range(nw188_.length(0)):
            nw188_[i0_42_] = init42_(i0_42_)
        d_1231_v8_ = nw188_
        d_1234_v9_: D12
        d_1234_v9_ = D12_DC33(d_1231_v8_, d_1231_v8_, p0, d_1219_v0_)
        d_1235_v10_: _dafny.MultiSet
        d_1235_v10_ = _dafny.MultiSet([d_1234_v9_])
        d_1236_v11_: _dafny.Array
        nw189_ = _dafny.Array(None, 17)
        nw189_[int(0)] = False
        nw189_[int(1)] = p0
        nw189_[int(2)] = p0
        nw189_[int(3)] = p0
        nw189_[int(4)] = p0
        nw189_[int(5)] = not(((d_1230_v7_).cf9) < (d_1221_v2_))
        nw189_[int(6)] = p0
        nw189_[int(7)] = p0
        nw189_[int(8)] = not(False)
        nw189_[int(9)] = (d_1235_v10_).ispropersubset(d_1235_v10_)
        nw189_[int(10)] = p0
        nw189_[int(11)] = (False) or (not(p0))
        nw189_[int(12)] = not(p0)
        nw189_[int(13)] = p0
        nw189_[int(14)] = (p0) or (not(p0))
        nw189_[int(15)] = p0
        nw189_[int(16)] = p0
        d_1236_v11_ = nw189_
        d_1237_v12_: _dafny.Map
        d_1237_v12_ = _dafny.Map({(self).f25: d_1236_v11_})
        d_1236_v11_ = ((d_1237_v12_)[p2] if (p2) in (d_1237_v12_) else d_1236_v11_)
        d_1238_v13_: _dafny.Seq
        d_1238_v13_ = _dafny.SeqWithoutIsStrInference([True])
        d_1239_v14_: _dafny.Map
        d_1239_v14_ = _dafny.Map({len(d_1238_v13_): p0})
        d_1240_v15_: _dafny.Seq
        d_1240_v15_ = _dafny.SeqWithoutIsStrInference([((d_1239_v14_)[((p1)[p0] if (p0) in (p1) else default__.fm31(d_1221_v2_, globalState))] if (((p1)[p0] if (p0) in (p1) else default__.fm31(d_1221_v2_, globalState))) in (d_1239_v14_) else p0)])
        d_1241_v16_: _dafny.Map
        d_1241_v16_ = _dafny.Map({998: len(d_1238_v13_)})
        d_1242_v19_: _dafny.Map
        d_1242_v19_ = _dafny.Map({p2: d_1241_v16_})
        d_1243_v21_: _dafny.Seq
        d_1243_v21_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1221_v2_])).cardinality, d_1219_v0_, (self).f25])
        d_1244_v22_: D14
        d_1244_v22_ = D14_DC37(_dafny.Map({d_1219_v0_: d_1219_v0_}))
        d_1245_v24_: _dafny.Map
        d_1245_v24_ = _dafny.Map({p1: d_1219_v0_})
        d_1246_v25_: _dafny.Array
        nw190_ = _dafny.Array(None, 18)
        nw190_[int(0)] = d_1241_v16_
        nw190_[int(1)] = d_1241_v16_
        def iife97_():
            def iife99_():
                coll51_ = _dafny.Set()
                compr_51_: int
                for compr_51_ in _dafny.IntegerRange(-449, 406):
                    d_1249_v18_: int = compr_51_
                    if ((-449) <= (d_1249_v18_)) and ((d_1249_v18_) < (406)):
                        coll51_ = coll51_.union(_dafny.Set([(d_1249_v18_) * (d_1219_v0_)]))
                return _dafny.Set(coll51_)
            coll49_ = _dafny.Map()
            def iife98_():
                coll50_ = _dafny.Set()
                compr_50_: int
                for compr_50_ in _dafny.IntegerRange(-449, 406):
                    d_1247_v18_: int = compr_50_
                    if ((-449) <= (d_1247_v18_)) and ((d_1247_v18_) < (406)):
                        coll50_ = coll50_.union(_dafny.Set([(d_1247_v18_) * (d_1219_v0_)]))
                return _dafny.Set(coll50_)
            compr_49_: int
            for compr_49_ in (iife98_()
            ).Elements:
                d_1248_v17_: int = compr_49_
                if (d_1248_v17_) in (iife99_()
                ):
                    coll49_[(d_1248_v17_) + (p2)] = p2
            return _dafny.Map(coll49_)
        nw190_[int(2)] = iife97_()
        
        nw190_[int(3)] = d_1241_v16_
        nw190_[int(4)] = _dafny.Map({len(d_1221_v2_): p2})
        nw190_[int(5)] = d_1241_v16_
        nw190_[int(6)] = d_1241_v16_
        nw190_[int(7)] = ((((d_1242_v19_)[(self).f25] if ((self).f25) in (d_1242_v19_) else d_1241_v16_)).set((self).f25, (self).f25)) | (d_1241_v16_)
        def iife100_():
            coll52_ = _dafny.Map()
            compr_52_: int
            for compr_52_ in (d_1243_v21_).Elements:
                d_1250_v20_: int = compr_52_
                if (d_1250_v20_) in (d_1243_v21_):
                    coll52_[default__.safeDivisionInt(d_1250_v20_, -531)] = -150
            return _dafny.Map(coll52_)
        nw190_[int(8)] = iife100_()
        
        nw190_[int(9)] = (d_1241_v16_).set(p2, d_1219_v0_)
        nw190_[int(10)] = d_1241_v16_
        nw190_[int(11)] = d_1241_v16_
        def iife101_():
            coll53_ = _dafny.Map()
            compr_53_: int
            for compr_53_ in _dafny.IntegerRange(-965, 817):
                d_1251_v23_: int = compr_53_
                if ((-965) <= (d_1251_v23_)) and ((d_1251_v23_) < (817)):
                    coll53_[(d_1251_v23_) - (p2)] = d_1219_v0_
            return _dafny.Map(coll53_)
        nw190_[int(12)] = ((d_1244_v22_).cf57) | (iife101_()
        )
        nw190_[int(13)] = _dafny.Map({d_1219_v0_: default__.fm31(d_1221_v2_, globalState)})
        nw190_[int(14)] = (d_1241_v16_) | (d_1241_v16_)
        nw190_[int(15)] = _dafny.Map({len(d_1245_v24_): (self).f25})
        nw190_[int(16)] = _dafny.Map({(self).f25: (self).f25})
        nw190_[int(17)] = d_1241_v16_
        d_1246_v25_ = nw190_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1246_v25_).length(0)):
            d_1252_i1_: int = guard_loop_5_
            if (True) and (((0) <= (d_1252_i1_)) and ((d_1252_i1_) < ((d_1246_v25_).length(0)))):
                (d_1246_v25_)[(d_1252_i1_)] = d_1241_v16_
        index219_ = default__.safeIndex(864, (d_1231_v8_).length(0))
        (d_1231_v8_)[index219_] = p0
        index220_ = default__.safeIndex(864, (d_1231_v8_).length(0))
        (d_1231_v8_)[index220_] = p0
        d_1253_v26_: _dafny.Map
        d_1253_v26_ = _dafny.Map({d_1231_v8_: ((self).f25) >= (p2)})
        if ((d_1253_v26_)[d_1231_v8_] if (d_1231_v8_) in (d_1253_v26_) else p0):
            if (not (not(p0)) or ((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))])) or (p0):
                d_1254_v27_: _dafny.Array
                nw191_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_1254_v27_ = nw191_
                d_1255_v28_: D13
                d_1255_v28_ = D13_DC35(d_1254_v27_)
                d_1255_v28_ = (D13_DC35(d_1254_v27_) if p0 else D13_DC35(d_1254_v27_))
                index221_ = default__.safeIndex(864, (d_1231_v8_).length(0))
                (d_1231_v8_)[index221_] = ((p0 if False else not((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))])) if (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))] else p0)
                (globalState).f23 = d_1229_v6_
                d_1256_v29_: _dafny.Map
                d_1256_v29_ = _dafny.Map({d_1241_v16_: len((p1).set((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))], d_1219_v0_))})
                def iife102_():
                    coll54_ = _dafny.Set()
                    compr_54_: int
                    for compr_54_ in (d_1241_v16_).keys.Elements:
                        d_1257_v30_: int = compr_54_
                        if (d_1257_v30_) in (d_1241_v16_):
                            coll54_ = coll54_.union(_dafny.Set([default__.safeModuloInt(d_1257_v30_, len(_dafny.Set({True, not(True), False, True, True})))]))
                    return _dafny.Set(coll54_)
                d_1256_v29_ = (d_1256_v29_).set(_dafny.Map({(self).f25: len(iife102_()
                )}), (len(p1)) - (len(d_1238_v13_)))
                (globalState).f10 = (_dafny.SeqWithoutIsStrInference([d_1229_v6_ for d_1258_i2_ in range(default__.abs(-11))])) + (d_1221_v2_)
            elif True:
                d_1259_v31_: C3
                nw192_ = C3()
                nw192_.ctor__((p2) * (d_1219_v0_), not(not(p0)), 953, (self).f26)
                d_1259_v31_ = nw192_
                d_1260_v32_: _dafny.Map
                d_1260_v32_ = _dafny.Map({(d_1259_v31_).f38: p0})
                d_1261_v33_: _dafny.Map
                d_1261_v33_ = _dafny.Map({(d_1259_v31_).f38: ((d_1260_v32_)[(d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]] if ((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]) in (d_1260_v32_) else (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))])})
                d_1262_v34_: D13
                d_1262_v34_ = D13_DC36((d_1261_v33_) | (_dafny.Map({(d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]: not(not((d_1259_v31_).f38))})), p0, ((self).f25) - ((0) - ((d_1259_v31_).f37)))
                d_1262_v34_ = d_1262_v34_
                d_1219_v0_ = (0) - ((default__.safeDivisionInt(d_1219_v0_, (d_1259_v31_).f37)) * ((d_1259_v31_).f37))
                (globalState).f12 = ((p2 if p0 else (0) - ((d_1259_v31_).f37))) < (((0) - ((self).f25)) - ((d_1259_v31_).f37))
                d_1263_v35_: C2
                nw193_ = C2()
                nw193_.ctor__()
                d_1263_v35_ = nw193_
            d_1264_v36_: _dafny.Array
            def lambda77_(d_1265_v0_):
                def lambda78_(d_1266_i3_):
                    return (d_1266_i3_) * (d_1265_v0_)

                return lambda78_

            init43_ = lambda77_(d_1219_v0_)
            nw194_ = _dafny.Array(None, 16)
            for i0_43_ in range(nw194_.length(0)):
                nw194_[i0_43_] = init43_(i0_43_)
            d_1264_v36_ = nw194_
            d_1267_v37_: D7
            d_1267_v37_ = D7_DC21(d_1264_v36_)
            d_1264_v36_ = (d_1267_v37_).cf32
            d_1268_v38_: _dafny.Map
            d_1268_v38_ = _dafny.Map({(d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]: (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]})
            d_1269_v39_: _dafny.MultiSet
            d_1269_v39_ = _dafny.MultiSet([True])
            d_1270_v40_: D15
            d_1270_v40_ = D15_DC39(d_1269_v39_)
            d_1268_v38_ = (d_1268_v38_).set(p0, not((_dafny.MultiSet([not(True)])).ispropersubset((d_1270_v40_).cf59)))
            d_1271_v41_: _dafny.Array
            nw195_ = _dafny.Array(_dafny.Seq({}), 18)
            d_1271_v41_ = nw195_
            d_1272_v42_: D9
            d_1272_v42_ = D9_DC26(p1)
            d_1273_v43_: _dafny.Seq
            d_1273_v43_ = _dafny.SeqWithoutIsStrInference([d_1272_v42_, d_1272_v42_, D9_DC26(p1), d_1272_v42_, d_1272_v42_])
            d_1274_v44_: _dafny.Seq
            d_1274_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1272_v42_]), d_1273_v43_])
            index222_ = default__.safeIndex(838, (d_1271_v41_).length(0))
            (d_1271_v41_)[index222_] = d_1274_v44_
            index223_ = default__.safeIndex(838, (d_1271_v41_).length(0))
            (d_1271_v41_)[index223_] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1272_v42_, d_1272_v42_])]) if (d_1240_v15_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwphtpp"))).set(default__.safeIndex(len(d_1221_v2_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwphtpp")))), d_1229_v6_)), len(d_1240_v15_))] else d_1274_v44_)
            d_1275_v45_: D12
            d_1275_v45_ = D12_DC32(p0)
            d_1276_v46_: _dafny.Map
            d_1276_v46_ = _dafny.Map({default__.fm33(p2, (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))], p2, globalState): (d_1275_v45_).cf47})
            d_1277_v47_: D10
            d_1277_v47_ = D10_DC28(d_1276_v46_)
            index224_ = default__.safeIndex(864, (d_1231_v8_).length(0))
            rhs229_ = d_1277_v47_
            rhs230_ = (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]
            rhs231_ = p0
            rhs232_ = (d_1268_v38_) == ((d_1268_v38_) | (d_1268_v38_))
            rhs233_ = d_1227_v4_
            lhs169_ = d_1231_v8_
            lhs170_ = default__.safeIndex(864, (d_1231_v8_).length(0))
            d_1277_v47_ = rhs229_
            r1 = rhs230_
            lhs169_[lhs170_] = rhs231_
            r1 = rhs232_
            d_1227_v4_ = rhs233_
        elif True:
            d_1278_v48_: _dafny.Array
            def lambda79_(d_1279_i4_):
                return (d_1279_i4_) + (502)

            init44_ = lambda79_
            nw196_ = _dafny.Array(None, 14)
            for i0_44_ in range(nw196_.length(0)):
                nw196_[i0_44_] = init44_(i0_44_)
            d_1278_v48_ = nw196_
            index225_ = default__.safeIndex(88, (d_1278_v48_).length(0))
            (d_1278_v48_)[index225_] = p2
            index226_ = default__.safeIndex(88, (d_1278_v48_).length(0))
            (d_1278_v48_)[index226_] = 339
            d_1280_v49_: C2
            nw197_ = C2()
            nw197_.ctor__()
            d_1280_v49_ = nw197_
            d_1281_v50_: _dafny.MultiSet
            d_1281_v50_ = _dafny.MultiSet([len(_dafny.Map({len(d_1238_v13_): d_1219_v0_}))])
            d_1282_v51_: _dafny.MultiSet
            d_1282_v51_ = _dafny.MultiSet([p0, p0])
            d_1283_v52_: C3
            nw198_ = C3()
            nw198_.ctor__((self).f25, (d_1281_v50_).isdisjoint(d_1281_v50_), ((d_1282_v51_)[True] if (True) in (d_1282_v51_) else 228), (self).f26)
            d_1283_v52_ = nw198_
            d_1283_v52_ = d_1283_v52_
            d_1278_v48_ = d_1278_v48_
            if (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]:
                index227_ = default__.safeIndex(864, (d_1231_v8_).length(0))
                (d_1231_v8_)[index227_] = not((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))])
                index228_ = default__.safeIndex(88, (d_1278_v48_).length(0))
                (d_1278_v48_)[index228_] = len((d_1240_v15_) + (d_1240_v15_))
                d_1219_v0_ = (D12_DC33(d_1236_v11_, d_1231_v8_, (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))], (d_1283_v52_).f37)).cf51
                d_1284_v53_: _dafny.Array
                nw199_ = _dafny.Array(_dafny.Seq({}), 6)
                d_1284_v53_ = nw199_
                index229_ = default__.safeIndex(378, (d_1284_v53_).length(0))
                (d_1284_v53_)[index229_] = d_1240_v15_
                index230_ = default__.safeIndex(378, (d_1284_v53_).length(0))
                index231_ = default__.safeIndex(88, (d_1278_v48_).length(0))
                rhs234_ = (d_1278_v48_)[default__.safeIndex(88, (d_1278_v48_).length(0))]
                rhs235_ = p2
                rhs236_ = (d_1240_v15_) + ((default__.fm1(((d_1239_v14_)[(d_1283_v52_).f37] if ((d_1283_v52_).f37) in (d_1239_v14_) else (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]), globalState)) + (d_1240_v15_))
                rhs237_ = (d_1278_v48_)[default__.safeIndex(88, (d_1278_v48_).length(0))]
                rhs238_ = d_1219_v0_
                lhs171_ = d_1284_v53_
                lhs172_ = default__.safeIndex(378, (d_1284_v53_).length(0))
                lhs173_ = d_1278_v48_
                lhs174_ = default__.safeIndex(88, (d_1278_v48_).length(0))
                d_1219_v0_ = rhs234_
                d_1219_v0_ = rhs235_
                lhs171_[lhs172_] = rhs236_
                d_1219_v0_ = rhs237_
                lhs173_[lhs174_] = rhs238_
                d_1285_v54_: _dafny.Seq
                d_1285_v54_ = _dafny.SeqWithoutIsStrInference([d_1231_v8_, d_1236_v11_])
                d_1286_v55_: _dafny.Map
                d_1286_v55_ = _dafny.Map({(d_1285_v54_) + (d_1285_v54_): not(not(p0))})
                d_1286_v55_ = (d_1286_v55_).set(d_1285_v54_, default__.fm0((d_1243_v21_)[default__.safeIndex((d_1283_v52_).f37, len(d_1243_v21_))], globalState))
            elif True:
                (globalState).f12 = (d_1283_v52_).f38
                d_1287_v56_: D0
                d_1287_v56_ = D0_DC2(p0)
                pat_let_tv24_ = d_1240_v15_
                pat_let_tv25_ = d_1221_v2_
                pat_let_tv26_ = d_1219_v0_
                pat_let_tv27_ = d_1221_v2_
                pat_let_tv28_ = d_1229_v6_
                pat_let_tv29_ = d_1240_v15_
                d_1288_v57_: _dafny.Array
                nw200_ = _dafny.Array(None, 8)
                nw200_[int(0)] = d_1287_v56_
                nw200_[int(1)] = d_1287_v56_
                nw200_[int(2)] = d_1287_v56_
                nw200_[int(3)] = (d_1287_v56_ if (d_1283_v52_).f38 else default__.fm54(_dafny.MultiSet(d_1240_v15_), globalState))
                nw200_[int(4)] = d_1287_v56_
                nw200_[int(5)] = D0_DC2((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))])
                nw200_[int(6)] = d_1287_v56_
                def iife103_(_pat_let24_0):
                    def iife104_(d_1289_dt__update__tmp_h2_):
                        def iife105_(_pat_let25_0):
                            def iife106_(d_1290_dt__update_hcf3_h0_):
                                return D0_DC2(d_1290_dt__update_hcf3_h0_)
                            return iife106_(_pat_let25_0)
                        return iife105_((pat_let_tv24_)[default__.safeIndex(len((pat_let_tv25_).set(default__.safeIndex(pat_let_tv26_, len(pat_let_tv27_)), pat_let_tv28_)), len(pat_let_tv29_))])
                    return iife104_(_pat_let24_0)
                nw200_[int(7)] = iife103_(d_1287_v56_)
                d_1288_v57_ = nw200_
                index232_ = default__.safeIndex(319, (d_1288_v57_).length(0))
                (d_1288_v57_)[index232_] = D0_DC2(p0)
                d_1291_v58_: _dafny.Array
                def lambda80_(d_1292_v50_):
                    def lambda81_(d_1293_i5_):
                        return d_1292_v50_

                    return lambda81_

                init45_ = lambda80_(d_1281_v50_)
                nw201_ = _dafny.Array(None, 5)
                for i0_45_ in range(nw201_.length(0)):
                    nw201_[i0_45_] = init45_(i0_45_)
                d_1291_v58_ = nw201_
                index233_ = default__.safeIndex(319, (d_1288_v57_).length(0))
                rhs239_ = d_1281_v50_
                rhs240_ = (d_1287_v56_ if ((d_1278_v48_)[default__.safeIndex(88, (d_1278_v48_).length(0))]) <= ((d_1283_v52_).f37) else d_1287_v56_)
                rhs241_ = d_1291_v58_
                lhs175_ = d_1288_v57_
                lhs176_ = default__.safeIndex(319, (d_1288_v57_).length(0))
                d_1281_v50_ = rhs239_
                lhs175_[lhs176_] = rhs240_
                d_1291_v58_ = rhs241_
                d_1294_v59_: _dafny.Map
                d_1294_v59_ = _dafny.Map({(d_1283_v52_).f38: (d_1283_v52_).f38})
                (globalState).f12 = not((((d_1278_v48_)[default__.safeIndex(88, (d_1278_v48_).length(0))]) * (len(d_1294_v59_))) <= (default__.fm31(d_1221_v2_, globalState)))
                (globalState).f12 = p0
                d_1238_v13_ = _dafny.SeqWithoutIsStrInference([True, not((d_1283_v52_).f38)])
        r0 = D0_DC2(p0)
        r1 = (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]
        d_1295_v60_: _dafny.Map
        d_1295_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, (d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))]]): ((d_1231_v8_)[default__.safeIndex(864, (d_1231_v8_).length(0))] if True else p0)})
        r2 = d_1295_v60_
        return r0, r1, r2

    @property
    def f36(self):
        return self._f36

class C5:
    def  __init__(self):
        self._f35: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f35):
        (self)._f35 = f35

    def m13(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        r3: _dafny.Array = _dafny.Array(None, 0)
        r2 = (p1) == (not((self).f35))
        d_1296_v0_: _dafny.Map
        d_1296_v0_ = _dafny.Map({p1: p0})
        d_1296_v0_ = (d_1296_v0_).set(p2, default__.fm18(globalState))
        r0 = p0
        if p2:
            (globalState).f12 = (default__.safeDivisionInt(p0, p0)) > (-283)
            r0 = p0
            (globalState).f6 = (default__.fm19(p1, globalState)) + ((_dafny.SeqWithoutIsStrInference([p0 for d_1297_i0_ in range(default__.abs(181))])) + (_dafny.SeqWithoutIsStrInference([p0])))
            d_1298_v1_: _dafny.Array
            nw202_ = _dafny.Array(int(0), 4)
            d_1298_v1_ = nw202_
            index234_ = default__.safeIndex(966, (d_1298_v1_).length(0))
            (d_1298_v1_)[index234_] = p0
            index235_ = default__.safeIndex(966, (d_1298_v1_).length(0))
            (d_1298_v1_)[index235_] = default__.safeDivisionInt(len((default__.fm20(default__.fm0(p0, globalState), globalState)) | (_dafny.Set({(self).f35, p2, p2, p2}))), default__.safeDivisionInt(p0, p0))
            d_1298_v1_ = d_1298_v1_
        elif True:
            (globalState).f12 = (self).f35
            if (p1 if p2 else not((self).f35)):
                d_1299_v2_: _dafny.Seq
                d_1299_v2_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1300_v3_: _dafny.Array
                nw203_ = _dafny.Array(None, 1)
                nw203_[int(0)] = d_1299_v2_
                d_1300_v3_ = nw203_
                index236_ = default__.safeIndex(579, (d_1300_v3_).length(0))
                (d_1300_v3_)[index236_] = (_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(default__.fm18(globalState), len(_dafny.SeqWithoutIsStrInference([p0]))), p0)
                d_1301_v4_: _dafny.Seq
                d_1301_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_1302_v5_: _dafny.Set
                d_1302_v5_ = _dafny.Set({778, len(d_1301_v4_)})
                d_1303_v6_: D4
                d_1303_v6_ = D4_DC11(d_1302_v5_)
                index237_ = default__.safeIndex(579, (d_1300_v3_).length(0))
                (d_1300_v3_)[index237_] = _dafny.SeqWithoutIsStrInference([p0, len((d_1303_v6_).cf18), default__.fm18(globalState)])
                r0 = p0
                r0 = p0
                r0 = p0
                d_1304_v7_: _dafny.Array
                def lambda82_(d_1305_p1_):
                    def lambda83_(d_1306_i1_):
                        return (D0_DC1(d_1305_p1_, not(d_1305_p1_)) if (self).f35 else D0_DC1((self).f35, d_1305_p1_))

                    return lambda83_

                init46_ = lambda82_(p1)
                nw204_ = _dafny.Array(None, 7)
                for i0_46_ in range(nw204_.length(0)):
                    nw204_[i0_46_] = init46_(i0_46_)
                d_1304_v7_ = nw204_
                d_1307_v8_: str
                d_1307_v8_ = _dafny.CodePoint('n')
                d_1308_v9_: D0
                d_1308_v9_ = D0_DC1((D1_DC6(d_1307_v8_, (self).f35, d_1301_v4_)).cf8, p1)
                index238_ = default__.safeIndex(387, (d_1304_v7_).length(0))
                (d_1304_v7_)[index238_] = d_1308_v9_
                index239_ = default__.safeIndex(387, (d_1304_v7_).length(0))
                (d_1304_v7_)[index239_] = D0_DC1((p0) >= (p0), p2)
            elif True:
                d_1309_v10_: D4
                d_1309_v10_ = D4_DC13(p0, (self).f35, default__.fm0(p0, globalState), p2)
                (globalState).f12 = default__.fm0(default__.safeDivisionInt(default__.fm18(globalState), (d_1309_v10_).cf19), globalState)
                d_1310_v11_: _dafny.Array
                nw205_ = _dafny.Array(int(0), 25)
                d_1310_v11_ = nw205_
                r3 = d_1310_v11_
                d_1311_v12_: D0
                d_1311_v12_ = D0_DC0((p2) and ((self).f35))
                d_1312_v13_: str
                d_1312_v13_ = _dafny.CodePoint('t')
                d_1313_v14_: _dafny.Seq
                d_1313_v14_ = _dafny.SeqWithoutIsStrInference([p0, p0, default__.fm18(globalState), (0) - (len(_dafny.Map({False: d_1312_v13_}))), p0])
                rhs242_ = (d_1313_v14_)[default__.safeIndex(565, len(d_1313_v14_))]
                rhs243_ = d_1311_v12_
                rhs244_ = p0
                rhs245_ = (default__.fm0(p0, globalState)) or (p1)
                r0 = rhs242_
                d_1311_v12_ = rhs243_
                r0 = rhs244_
                r2 = rhs245_
                (globalState).f12 = p1
                d_1314_v15_: _dafny.Seq
                d_1314_v15_ = _dafny.SeqWithoutIsStrInference([d_1310_v11_])
                d_1315_v16_: _dafny.Array
                nw206_ = _dafny.Array(None, 21)
                nw206_[int(0)] = d_1310_v11_
                nw206_[int(1)] = d_1310_v11_
                nw206_[int(2)] = (d_1314_v15_)[default__.safeIndex(p0, len(d_1314_v15_))]
                nw206_[int(3)] = d_1310_v11_
                nw206_[int(4)] = d_1310_v11_
                nw206_[int(5)] = d_1310_v11_
                nw206_[int(6)] = d_1310_v11_
                nw206_[int(7)] = d_1310_v11_
                nw206_[int(8)] = d_1310_v11_
                nw206_[int(9)] = d_1310_v11_
                nw206_[int(10)] = d_1310_v11_
                nw206_[int(11)] = d_1310_v11_
                nw206_[int(12)] = d_1310_v11_
                nw206_[int(13)] = (d_1310_v11_ if default__.fm0(p0, globalState) else d_1310_v11_)
                nw206_[int(14)] = d_1310_v11_
                nw206_[int(15)] = d_1310_v11_
                nw206_[int(16)] = (d_1310_v11_ if (self).f35 else d_1310_v11_)
                nw206_[int(17)] = d_1310_v11_
                nw206_[int(18)] = d_1310_v11_
                nw206_[int(19)] = d_1310_v11_
                nw206_[int(20)] = d_1310_v11_
                d_1315_v16_ = nw206_
                index240_ = default__.safeIndex(613, (d_1315_v16_).length(0))
                (d_1315_v16_)[index240_] = d_1310_v11_
                index241_ = default__.safeIndex(613, (d_1315_v16_).length(0))
                (d_1315_v16_)[index241_] = d_1310_v11_
            d_1316_v17_: _dafny.Array
            nw207_ = _dafny.Array(D3.default()(), 14)
            d_1316_v17_ = nw207_
            d_1317_v18_: _dafny.Array
            nw208_ = _dafny.Array(False, 13)
            d_1317_v18_ = nw208_
            d_1318_v19_: D3
            d_1318_v19_ = D3_DC9(_dafny.Set({d_1317_v18_}))
            index242_ = default__.safeIndex(492, (d_1316_v17_).length(0))
            (d_1316_v17_)[index242_] = d_1318_v19_
            d_1319_v20_: str
            d_1319_v20_ = _dafny.CodePoint('n')
            index243_ = default__.safeIndex(492, (d_1316_v17_).length(0))
            rhs246_ = d_1318_v19_
            rhs247_ = (p0) + (p0)
            rhs248_ = d_1319_v20_
            lhs177_ = d_1316_v17_
            lhs178_ = default__.safeIndex(492, (d_1316_v17_).length(0))
            lhs179_ = globalState
            lhs177_[lhs178_] = rhs246_
            r0 = rhs247_
            lhs179_.f23 = rhs248_
            (globalState).f23 = d_1319_v20_
            (globalState).f12 = (778) > (p0)
        d_1320_v21_: _dafny.MultiSet
        d_1320_v21_ = _dafny.MultiSet([p0, p0])
        hi7_ = p0
        for d_1321_i2_ in range((((d_1320_v21_).set(p0, default__.abs(p0))) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([p2, p2, True])), p0, p0, p0]))).cardinality, hi7_):
            (globalState).f12 = not(default__.fm0(len(_dafny.SeqWithoutIsStrInference([d_1321_i2_ for d_1322_i3_ in range(default__.abs(639))])), globalState))
            r2 = p1
            d_1323_v22_: _dafny.Seq
            d_1323_v22_ = _dafny.SeqWithoutIsStrInference([p2, p1])
            d_1324_v23_: _dafny.Map
            d_1324_v23_ = _dafny.Map({d_1323_v22_: p2})
            d_1325_v24_: _dafny.Array
            nw209_ = _dafny.Array(None, 7)
            nw209_[int(0)] = d_1324_v23_
            nw209_[int(1)] = d_1324_v23_
            nw209_[int(2)] = d_1324_v23_
            nw209_[int(3)] = (d_1324_v23_) | (d_1324_v23_)
            nw209_[int(4)] = d_1324_v23_
            nw209_[int(5)] = (d_1324_v23_) | (d_1324_v23_)
            nw209_[int(6)] = d_1324_v23_
            d_1325_v24_ = nw209_
            index244_ = default__.safeIndex(180, (d_1325_v24_).length(0))
            (d_1325_v24_)[index244_] = d_1324_v23_
            index245_ = default__.safeIndex(180, (d_1325_v24_).length(0))
            (d_1325_v24_)[index245_] = (d_1324_v23_) | (d_1324_v23_)
            d_1326_v25_: _dafny.Set
            d_1326_v25_ = _dafny.Set({True})
            if ((self).f35) == ((_dafny.Set({False})).ispropersubset(d_1326_v25_)):
                (globalState).f12 = p1
                d_1327_v26_: str
                d_1327_v26_ = _dafny.CodePoint('x')
                (globalState).f23 = d_1327_v26_
                d_1328_v27_: _dafny.Seq
                d_1328_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_1329_v28_: D5
                d_1329_v28_ = D5_DC17(p0, p2)
                d_1330_v29_: _dafny.Seq
                d_1330_v29_ = _dafny.SeqWithoutIsStrInference([d_1321_i2_, (d_1329_v28_).cf26, d_1321_i2_])
                d_1331_v30_: _dafny.Seq
                d_1331_v30_ = _dafny.SeqWithoutIsStrInference([len(d_1330_v29_)])
                d_1332_v31_: _dafny.Seq
                d_1332_v31_ = _dafny.SeqWithoutIsStrInference([d_1321_i2_, len(d_1328_v27_), len(d_1330_v29_), len(d_1330_v29_)])
                d_1333_v32_: _dafny.Array
                def lambda84_(d_1334_v27_):
                    def lambda85_(d_1335_i4_):
                        return d_1334_v27_

                    return lambda85_

                init47_ = lambda84_(d_1328_v27_)
                nw210_ = _dafny.Array(None, 16)
                for i0_47_ in range(nw210_.length(0)):
                    nw210_[i0_47_] = init47_(i0_47_)
                d_1333_v32_ = nw210_
                d_1336_v33_: C3
                nw211_ = C3()
                nw211_.ctor__((d_1331_v30_)[default__.safeIndex(d_1321_i2_, len(d_1331_v30_))], (d_1326_v25_).isdisjoint(d_1326_v25_), default__.safeModuloInt(d_1321_i2_, d_1321_i2_), d_1333_v32_)
                d_1336_v33_ = nw211_
                d_1337_v34_: C1
                nw212_ = C1()
                nw212_.ctor__(d_1321_i2_, d_1333_v32_)
                d_1337_v34_ = nw212_
                d_1338_v35_: _dafny.Array
                nw213_ = _dafny.Array(None, 1)
                nw213_[int(0)] = len(d_1323_v22_)
                d_1338_v35_ = nw213_
                d_1339_v36_: D7
                d_1339_v36_ = D7_DC21(d_1338_v35_)
                r3 = (d_1339_v36_).cf32
            elif True:
                d_1340_v37_: _dafny.Array
                nw214_ = _dafny.Array(False, 29)
                d_1340_v37_ = nw214_
                d_1341_v38_: D12
                d_1341_v38_ = D12_DC33(d_1340_v37_, d_1340_v37_, p2, p0)
                d_1342_v39_: _dafny.Map
                d_1342_v39_ = _dafny.Map({(d_1341_v38_).cf51: d_1340_v37_})
                d_1342_v39_ = (d_1342_v39_).set(d_1321_i2_, d_1340_v37_)
                d_1343_v40_: _dafny.Array
                def lambda86_(d_1344_i2_):
                    def lambda87_(d_1345_i5_):
                        return (d_1345_i5_) * (d_1344_i2_)

                    return lambda87_

                init48_ = lambda86_(d_1321_i2_)
                nw215_ = _dafny.Array(None, 28)
                for i0_48_ in range(nw215_.length(0)):
                    nw215_[i0_48_] = init48_(i0_48_)
                d_1343_v40_ = nw215_
                index246_ = default__.safeIndex(87, (d_1343_v40_).length(0))
                (d_1343_v40_)[index246_] = p0
                index247_ = default__.safeIndex(87, (d_1343_v40_).length(0))
                (d_1343_v40_)[index247_] = d_1321_i2_
                d_1346_v41_: D3
                d_1346_v41_ = D3_DC10(((self).f35) and (p1), d_1340_v37_, (self).f35, (d_1321_i2_) * (len(d_1296_v0_)), default__.fm18(globalState))
                d_1347_v42_: _dafny.Array
                nw216_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_1347_v42_ = nw216_
                d_1348_v43_: str
                d_1348_v43_ = _dafny.CodePoint('f')
                index248_ = default__.safeIndex(889, (d_1347_v42_).length(0))
                (d_1347_v42_)[index248_] = d_1348_v43_
                index249_ = default__.safeIndex(889, (d_1347_v42_).length(0))
                rhs249_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('f') if (self).f35 else _dafny.CodePoint('q')) for d_1349_i6_ in range(default__.abs(990))])
                rhs250_ = d_1346_v41_
                rhs251_ = d_1348_v43_
                lhs180_ = globalState
                lhs181_ = d_1347_v42_
                lhs182_ = default__.safeIndex(889, (d_1347_v42_).length(0))
                lhs180_.f10 = rhs249_
                d_1346_v41_ = rhs250_
                lhs181_[lhs182_] = rhs251_
                (globalState).f10 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdpspo"))
                index250_ = default__.safeIndex(87, (d_1343_v40_).length(0))
                (d_1343_v40_)[index250_] = d_1321_i2_
        d_1350_v44_: _dafny.Seq
        d_1350_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvtp"))
        d_1351_v45_: _dafny.Map
        d_1351_v45_ = _dafny.Map({p0: p1})
        rhs252_ = d_1350_v44_
        rhs253_ = ((p0) == (len(d_1351_v45_))) and (p1)
        rhs254_ = d_1350_v44_
        lhs183_ = globalState
        lhs184_ = globalState
        lhs185_ = globalState
        lhs183_.f10 = rhs252_
        lhs184_.f12 = rhs253_
        lhs185_.f10 = rhs254_
        r0 = p0
        d_1352_v46_: _dafny.Set
        d_1352_v46_ = _dafny.Set({p0})
        def iife107_():
            coll55_ = _dafny.Set()
            compr_55_: int
            for compr_55_ in _dafny.IntegerRange(896, 496):
                d_1353_v47_: int = compr_55_
                if ((896) <= (d_1353_v47_)) and ((d_1353_v47_) < (496)):
                    coll55_ = coll55_.union(_dafny.Set([default__.safeModuloInt(d_1353_v47_, p0)]))
            return _dafny.Set(coll55_)
        r1 = ((d_1352_v46_) - (iife107_()
        )).intersection(d_1352_v46_)
        r2 = default__.fm0(p0, globalState)
        d_1354_v48_: _dafny.Array
        nw217_ = _dafny.Array(int(0), 15)
        d_1354_v48_ = nw217_
        d_1355_v49_: _dafny.Map
        d_1355_v49_ = _dafny.Map({(d_1320_v21_) - (d_1320_v21_): d_1354_v48_})
        r3 = ((d_1355_v49_)[d_1320_v21_] if (d_1320_v21_) in (d_1355_v49_) else d_1354_v48_)
        return r0, r1, r2, r3

    def m14(self, p0, p1, p2, globalState):
        if (p0) >= (-963):
            d_1356_v0_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
            d_1356_v0_ = nw218_
            d_1357_v1_: C1
            nw219_ = C1()
            nw219_.ctor__(len(_dafny.SeqWithoutIsStrInference([False])), d_1356_v0_)
            d_1357_v1_ = nw219_
            d_1358_v2_: _dafny.Array
            nw220_ = _dafny.Array(False, 24)
            d_1358_v2_ = nw220_
            index251_ = default__.safeIndex(783, (d_1358_v2_).length(0))
            (d_1358_v2_)[index251_] = (self).f35
            index252_ = default__.safeIndex(783, (d_1358_v2_).length(0))
            (d_1358_v2_)[index252_] = p2
            d_1359_v3_: int
            d_1359_v3_ = -49
            d_1360_v4_: _dafny.Array
            def lambda88_(d_1361_p0_):
                def lambda89_(d_1362_i0_):
                    return (d_1362_i0_) - (d_1361_p0_)

                return lambda89_

            init49_ = lambda88_(p0)
            nw221_ = _dafny.Array(None, 3)
            for i0_49_ in range(nw221_.length(0)):
                nw221_[i0_49_] = init49_(i0_49_)
            d_1360_v4_ = nw221_
            d_1363_v5_: _dafny.Map
            d_1363_v5_ = _dafny.Map({d_1360_v4_: (d_1358_v2_)[default__.safeIndex(783, (d_1358_v2_).length(0))]})
            d_1359_v3_ = len(d_1363_v5_)
            d_1364_v6_: _dafny.Map
            d_1364_v6_ = _dafny.Map({(self).f35: (self).f35})
            d_1365_v7_: _dafny.Map
            d_1365_v7_ = _dafny.Map({d_1359_v3_: d_1364_v6_})
            d_1365_v7_ = (d_1365_v7_).set(-393, (d_1364_v6_) | (d_1364_v6_))
            d_1366_v8_: str
            d_1366_v8_ = _dafny.CodePoint('r')
            d_1367_v9_: _dafny.Map
            d_1367_v9_ = _dafny.Map({(self).f35: d_1366_v8_})
            d_1368_v10_: C4
            nw222_ = C4()
            nw222_.ctor__(d_1356_v0_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhtp"))).set(default__.safeIndex(d_1359_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhtp")))), ((d_1367_v9_)[(self).f35] if ((self).f35) in (d_1367_v9_) else d_1366_v8_))), d_1356_v0_)
            d_1368_v10_ = nw222_
        elif True:
            d_1369_v11_: int
            d_1369_v11_ = 108
            d_1370_v12_: D0
            d_1370_v12_ = D0_DC0(p2)
            d_1371_v13_: _dafny.Map
            d_1371_v13_ = _dafny.Map({d_1370_v12_: True})
            d_1372_v14_: _dafny.Seq
            d_1372_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsxb"))
            d_1373_v15_: _dafny.Map
            d_1373_v15_ = _dafny.Map({(self).f35: d_1372_v14_})
            d_1374_v16_: str
            d_1374_v16_ = _dafny.CodePoint('t')
            d_1369_v11_ = default__.safeDivisionInt(default__.fm21(d_1371_v13_, default__.fm31(d_1372_v14_, globalState), p2, (((d_1373_v15_)[(self).f35] if ((self).f35) in (d_1373_v15_) else d_1372_v14_)).set(default__.safeIndex(d_1369_v11_, len(((d_1373_v15_)[(self).f35] if ((self).f35) in (d_1373_v15_) else d_1372_v14_))), d_1374_v16_), globalState), d_1369_v11_)
            d_1375_v17_: _dafny.Array
            nw223_ = _dafny.Array(_dafny.Map({}), 19)
            d_1375_v17_ = nw223_
            d_1376_v18_: _dafny.Map
            d_1376_v18_ = _dafny.Map({d_1375_v17_: p2})
            d_1377_v19_: _dafny.Seq
            d_1377_v19_ = _dafny.SeqWithoutIsStrInference([p2, (self).f35])
            d_1376_v18_ = (d_1376_v18_).set(d_1375_v17_, (d_1377_v19_) == (_dafny.SeqWithoutIsStrInference([p2])))
            d_1378_v20_: _dafny.Map
            d_1378_v20_ = _dafny.Map({p0: p0})
            d_1378_v20_ = d_1378_v20_
            d_1379_v21_: D4
            d_1379_v21_ = D4_DC12()
            d_1380_v22_: D4
            d_1380_v22_ = D4_DC14(d_1379_v21_)
            d_1381_v23_: _dafny.MultiSet
            d_1381_v23_ = _dafny.MultiSet([_dafny.MultiSet([p2, not(p2)]), _dafny.MultiSet(d_1377_v19_)])
            d_1382_v24_: _dafny.Map
            d_1382_v24_ = _dafny.Map({d_1380_v22_: d_1381_v23_})
            d_1382_v24_ = (d_1382_v24_).set(d_1380_v22_, (d_1381_v23_) - (d_1381_v23_))
            d_1383_v25_: _dafny.Array
            nw224_ = _dafny.Array(None, 24)
            nw224_[int(0)] = d_1372_v14_
            nw224_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyqt"))
            nw224_[int(2)] = d_1372_v14_
            nw224_[int(3)] = d_1372_v14_
            nw224_[int(4)] = d_1372_v14_
            nw224_[int(5)] = d_1372_v14_
            nw224_[int(6)] = d_1372_v14_
            nw224_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1374_v16_ for d_1384_i1_ in range(default__.abs(108))])
            nw224_[int(8)] = d_1372_v14_
            nw224_[int(9)] = d_1372_v14_
            nw224_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "om"))
            nw224_[int(11)] = d_1372_v14_
            nw224_[int(12)] = d_1372_v14_
            nw224_[int(13)] = d_1372_v14_
            nw224_[int(14)] = d_1372_v14_
            nw224_[int(15)] = d_1372_v14_
            nw224_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cupth"))
            nw224_[int(17)] = d_1372_v14_
            nw224_[int(18)] = d_1372_v14_
            nw224_[int(19)] = d_1372_v14_
            nw224_[int(20)] = d_1372_v14_
            nw224_[int(21)] = d_1372_v14_
            nw224_[int(22)] = d_1372_v14_
            nw224_[int(23)] = d_1372_v14_
            d_1383_v25_ = nw224_
            d_1385_v26_: C1
            nw225_ = C1()
            nw225_.ctor__((0) - (d_1369_v11_), d_1383_v25_)
            d_1385_v26_ = nw225_
        d_1386_v27_: _dafny.Map
        d_1386_v27_ = _dafny.Map({p0: p2})
        d_1386_v27_ = (d_1386_v27_).set(p0, (p0) < (p0))
        d_1387_v28_: _dafny.Array
        def lambda90_(d_1388_v27_):
            def lambda91_(d_1389_i3_):
                return d_1388_v27_

            return lambda91_

        init50_ = lambda90_(d_1386_v27_)
        nw226_ = _dafny.Array(None, 1)
        for i0_50_ in range(nw226_.length(0)):
            nw226_[i0_50_] = init50_(i0_50_)
        d_1387_v28_ = nw226_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1387_v28_).length(0)):
            d_1390_i2_: int = guard_loop_6_
            if (True) and (((0) <= (d_1390_i2_)) and ((d_1390_i2_) < ((d_1387_v28_).length(0)))):
                (d_1387_v28_)[(d_1390_i2_)] = (d_1386_v27_) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-527])): (self).f35}))
        d_1391_v29_: _dafny.Array
        nw227_ = _dafny.Array(False, 13)
        d_1391_v29_ = nw227_
        d_1392_v30_: _dafny.Set
        d_1392_v30_ = _dafny.Set({True})
        index253_ = default__.safeIndex(813, (d_1391_v29_).length(0))
        (d_1391_v29_)[index253_] = (d_1392_v30_) == (d_1392_v30_)
        index254_ = default__.safeIndex(813, (d_1391_v29_).length(0))
        (d_1391_v29_)[index254_] = (self).f35
        d_1393_v31_: _dafny.Array
        def lambda92_(d_1394_p1_):
            def lambda93_(d_1395_i5_):
                return d_1394_p1_

            return lambda93_

        init51_ = lambda92_(p1)
        nw228_ = _dafny.Array(None, 11)
        for i0_51_ in range(nw228_.length(0)):
            nw228_[i0_51_] = init51_(i0_51_)
        d_1393_v31_ = nw228_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1393_v31_).length(0)):
            d_1396_i4_: int = guard_loop_7_
            if (True) and (((0) <= (d_1396_i4_)) and ((d_1396_i4_) < ((d_1393_v31_).length(0)))):
                (d_1393_v31_)[(d_1396_i4_)] = p1
        d_1386_v27_ = d_1386_v27_

    @property
    def f35(self):
        return self._f35

class C6(T1, T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f25, f26):
        (self)._f25 = f25
        (self)._f26 = f26

    def m4(self, p0, globalState):
        d_1397_v0_: _dafny.Seq
        d_1397_v0_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, (self).f25, (self).f25])
        d_1398_v1_: C4
        nw229_ = C4()
        nw229_.ctor__((self).f26, len(d_1397_v0_), (self).f26)
        d_1398_v1_ = nw229_
        d_1399_v2_: bool
        d_1399_v2_ = False
        if d_1399_v2_:
            d_1400_v3_: _dafny.Array
            nw230_ = _dafny.Array(int(0), 28)
            d_1400_v3_ = nw230_
            d_1401_v4_: D7
            d_1401_v4_ = D7_DC21(d_1400_v3_)
            source20_ = d_1401_v4_
            if source20_.is_DC22:
                d_1402___mcc_h0_ = source20_.cf33
                d_1403___mcc_h1_ = source20_.cf34
                d_1404_cf34_ = d_1403___mcc_h1_
                d_1405_cf33_ = d_1402___mcc_h0_
                d_1406_v5_: C2
                nw231_ = C2()
                nw231_.ctor__()
                d_1406_v5_ = nw231_
                d_1407_v6_: _dafny.Seq
                d_1407_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgkrb"))
                d_1408_v7_: str
                d_1408_v7_ = _dafny.CodePoint('i')
                d_1409_v8_: D1
                d_1409_v8_ = D1_DC6(d_1408_v7_, d_1399_v2_, d_1407_v6_)
                (globalState).f10 = (d_1407_v6_) + ((d_1409_v8_).cf9)
                d_1410_v9_: _dafny.Map
                d_1410_v9_ = _dafny.Map({d_1407_v6_: not(not (d_1399_v2_) or (d_1399_v2_))})
                d_1410_v9_ = (d_1410_v9_).set(d_1407_v6_, (-573) <= (535))
                (globalState).f12 = ((self).f25) > (((self).f25) * ((self).f25))
            elif True:
                d_1411___mcc_h2_ = source20_.cf32
                d_1412_cf32_ = d_1411___mcc_h2_
                d_1413_v10_: _dafny.Map
                d_1413_v10_ = _dafny.Map({d_1399_v2_: d_1412_cf32_})
                rhs255_ = ((d_1413_v10_)[d_1399_v2_] if (d_1399_v2_) in (d_1413_v10_) else (d_1400_v3_ if d_1399_v2_ else d_1400_v3_))
                rhs256_ = d_1399_v2_
                lhs186_ = globalState
                d_1412_cf32_ = rhs255_
                lhs186_.f12 = rhs256_
                index255_ = default__.safeIndex(507, (d_1412_cf32_).length(0))
                (d_1412_cf32_)[index255_] = (self).f25
                d_1414_v11_: _dafny.Seq
                d_1414_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opuevflja"))
                d_1415_v12_: _dafny.Map
                d_1415_v12_ = _dafny.Map({d_1399_v2_: (self).f25})
                index256_ = default__.safeIndex(507, (d_1412_cf32_).length(0))
                (d_1412_cf32_)[index256_] = (len(d_1414_v11_) if d_1399_v2_ else default__.safeDivisionInt((self).f25, len(d_1415_v12_)))
                d_1416_v13_: str
                d_1416_v13_ = _dafny.CodePoint('s')
                index257_ = default__.safeIndex(507, (d_1412_cf32_).length(0))
                (d_1412_cf32_)[index257_] = default__.fm31((d_1414_v11_) + ((d_1414_v11_).set(default__.safeIndex((self).f25, len(d_1414_v11_)), d_1416_v13_)), globalState)
                d_1417_v14_: _dafny.MultiSet
                d_1417_v14_ = _dafny.MultiSet([True, d_1399_v2_])
                d_1418_v15_: _dafny.Seq
                d_1418_v15_ = _dafny.SeqWithoutIsStrInference([d_1415_v12_])
                d_1419_v16_: D9
                d_1419_v16_ = D9_DC26(d_1415_v12_)
                d_1420_v17_: _dafny.Array
                nw232_ = _dafny.Array(None, 19)
                nw232_[int(0)] = d_1415_v12_
                nw232_[int(1)] = (_dafny.Map({d_1399_v2_: (0) - (-439)})).set(not(False), (0) - (((d_1417_v14_)[d_1399_v2_] if (d_1399_v2_) in (d_1417_v14_) else (self).f25)))
                nw232_[int(2)] = d_1415_v12_
                nw232_[int(3)] = (d_1415_v12_).set(d_1399_v2_, (d_1412_cf32_)[default__.safeIndex(507, (d_1412_cf32_).length(0))])
                nw232_[int(4)] = d_1415_v12_
                nw232_[int(5)] = d_1415_v12_
                nw232_[int(6)] = d_1415_v12_
                nw232_[int(7)] = d_1415_v12_
                nw232_[int(8)] = d_1415_v12_
                nw232_[int(9)] = d_1415_v12_
                nw232_[int(10)] = d_1415_v12_
                nw232_[int(11)] = (d_1418_v15_)[default__.safeIndex(603, len(d_1418_v15_))]
                nw232_[int(12)] = d_1415_v12_
                nw232_[int(13)] = d_1415_v12_
                nw232_[int(14)] = d_1415_v12_
                nw232_[int(15)] = d_1415_v12_
                nw232_[int(16)] = (d_1419_v16_).cf38
                nw232_[int(17)] = _dafny.Map({default__.fm0((self).f25, globalState): (self).f25})
                nw232_[int(18)] = d_1415_v12_
                d_1420_v17_ = nw232_
                d_1421_v18_: _dafny.Map
                d_1421_v18_ = _dafny.Map({(self).f25: d_1420_v17_})
                d_1421_v18_ = (d_1421_v18_).set((0) - (((d_1412_cf32_)[default__.safeIndex(507, (d_1412_cf32_).length(0))]) + ((d_1412_cf32_)[default__.safeIndex(507, (d_1412_cf32_).length(0))])), d_1420_v17_)
            d_1422_v19_: _dafny.Seq
            d_1422_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhdow"))
            d_1423_v20_: _dafny.Map
            d_1423_v20_ = _dafny.Map({d_1422_v19_: True})
            def iife108_():
                coll56_ = _dafny.Map()
                compr_56_: _dafny.Seq
                for compr_56_ in (d_1423_v20_).keys.Elements:
                    d_1424_v21_: _dafny.Seq = compr_56_
                    if (d_1424_v21_) in (d_1423_v20_):
                        coll56_[d_1424_v21_] = d_1399_v2_
                return _dafny.Map(coll56_)
            d_1423_v20_ = iife108_()
            
            source21_ = D9_DC27((self).f25, (self).f25)
            if source21_.is_DC27:
                d_1425___mcc_h3_ = source21_.cf39
                d_1426___mcc_h4_ = source21_.cf40
                d_1427_cf40_ = d_1426___mcc_h4_
                d_1428_cf39_ = d_1425___mcc_h3_
                d_1429_v22_: _dafny.MultiSet
                d_1429_v22_ = _dafny.MultiSet([d_1427_cf40_])
                d_1430_v23_: _dafny.MultiSet
                d_1430_v23_ = _dafny.MultiSet([d_1429_v22_, _dafny.MultiSet(d_1397_v0_), (_dafny.MultiSet([155, (self).f25, 974])) | (d_1429_v22_)])
                d_1430_v23_ = (d_1430_v23_) - ((d_1430_v23_) - (d_1430_v23_))
                d_1399_v2_ = not(True)
                d_1431_v24_: str
                d_1431_v24_ = _dafny.CodePoint('p')
                d_1432_v25_: D1
                d_1432_v25_ = D1_DC6(d_1431_v24_, not(d_1399_v2_), d_1422_v19_)
                d_1433_v26_: _dafny.Map
                d_1433_v26_ = _dafny.Map({d_1399_v2_: (_dafny.MultiSet([len((d_1432_v25_).cf9)])).intersection(d_1429_v22_)})
                rhs257_ = d_1431_v24_
                rhs258_ = d_1399_v2_
                rhs259_ = d_1433_v26_
                lhs187_ = globalState
                lhs188_ = globalState
                lhs187_.f23 = rhs257_
                lhs188_.f12 = rhs258_
                d_1433_v26_ = rhs259_
                (globalState).f12 = d_1399_v2_
            elif True:
                d_1434___mcc_h5_ = source21_.cf38
                d_1435_cf38_ = d_1434___mcc_h5_
                d_1436_v27_: C1
                nw233_ = C1()
                nw233_.ctor__(-131, (d_1398_v1_).f36)
                d_1436_v27_ = nw233_
                d_1437_v28_: _dafny.Array
                nw234_ = _dafny.Array(None, 4)
                nw234_[int(0)] = d_1436_v27_
                nw234_[int(1)] = d_1436_v27_
                nw234_[int(2)] = d_1436_v27_
                nw234_[int(3)] = d_1436_v27_
                d_1437_v28_ = nw234_
                index258_ = default__.safeIndex(52, (d_1437_v28_).length(0))
                (d_1437_v28_)[index258_] = d_1436_v27_
                index259_ = default__.safeIndex(52, (d_1437_v28_).length(0))
                (d_1437_v28_)[index259_] = (d_1436_v27_ if d_1399_v2_ else d_1436_v27_)
                d_1438_v29_: _dafny.Array
                nw235_ = _dafny.Array(_dafny.Seq({}), 6)
                d_1438_v29_ = nw235_
                d_1439_v30_: _dafny.Array
                nw236_ = _dafny.Array(None, 11)
                nw236_[int(0)] = d_1399_v2_
                nw236_[int(1)] = d_1399_v2_
                nw236_[int(2)] = d_1399_v2_
                nw236_[int(3)] = d_1399_v2_
                nw236_[int(4)] = d_1399_v2_
                nw236_[int(5)] = d_1399_v2_
                nw236_[int(6)] = d_1399_v2_
                nw236_[int(7)] = True
                nw236_[int(8)] = d_1399_v2_
                nw236_[int(9)] = d_1399_v2_
                nw236_[int(10)] = d_1399_v2_
                d_1439_v30_ = nw236_
                d_1440_v31_: _dafny.Seq
                d_1440_v31_ = _dafny.SeqWithoutIsStrInference([d_1439_v30_, d_1439_v30_])
                index260_ = default__.safeIndex(168, (d_1438_v29_).length(0))
                (d_1438_v29_)[index260_] = d_1440_v31_
                index261_ = default__.safeIndex(168, (d_1438_v29_).length(0))
                (d_1438_v29_)[index261_] = d_1440_v31_
                index262_ = default__.safeIndex(9, (d_1439_v30_).length(0))
                (d_1439_v30_)[index262_] = d_1399_v2_
                index263_ = default__.safeIndex(9, (d_1439_v30_).length(0))
                (d_1439_v30_)[index263_] = d_1399_v2_
                d_1441_v32_: str
                d_1441_v32_ = _dafny.CodePoint('k')
                (globalState).f10 = (default__.fm2((self).f25, (self).f25, globalState) if ((self).f25) in (d_1397_v0_) else (d_1422_v19_).set(default__.safeIndex((self).f25, len(d_1422_v19_)), d_1441_v32_))
            d_1442_v33_: C0
            nw237_ = C0()
            nw237_.ctor__()
            d_1442_v33_ = nw237_
            d_1443_v34_: int
            d_1443_v34_ = 731
            d_1443_v34_ = default__.safeModuloInt((0) - (len(d_1422_v19_)), (self).f25)
        elif True:
            d_1444_v35_: str
            d_1444_v35_ = _dafny.CodePoint('r')
            d_1445_v36_: _dafny.Map
            d_1445_v36_ = _dafny.Map({_dafny.CodePoint('n'): d_1444_v35_})
            d_1446_v37_: _dafny.Set
            d_1446_v37_ = _dafny.Set({d_1445_v36_})
            d_1447_v38_: _dafny.MultiSet
            d_1447_v38_ = _dafny.MultiSet([False])
            d_1448_v39_: _dafny.Seq
            d_1448_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrtlea"))
            rhs260_ = (_dafny.MultiSet([d_1399_v2_, d_1399_v2_, d_1399_v2_])) != ((d_1447_v38_) - (d_1447_v38_))
            rhs261_ = (d_1448_v39_)[default__.safeIndex(default__.safeModuloInt((0) - ((self).f25), (self).f25), len(d_1448_v39_))]
            rhs262_ = d_1446_v37_
            lhs189_ = globalState
            lhs190_ = globalState
            lhs189_.f12 = rhs260_
            lhs190_.f23 = rhs261_
            d_1446_v37_ = rhs262_
            d_1449_v40_: _dafny.Array
            def lambda94_(d_1450_v38_):
                def lambda95_(d_1451_i0_):
                    return default__.safeDivisionInt(d_1451_i0_, (d_1450_v38_).cardinality)

                return lambda95_

            init52_ = lambda94_(d_1447_v38_)
            nw238_ = _dafny.Array(None, 5)
            for i0_52_ in range(nw238_.length(0)):
                nw238_[i0_52_] = init52_(i0_52_)
            d_1449_v40_ = nw238_
            d_1452_v41_: _dafny.Map
            d_1452_v41_ = _dafny.Map({D0_DC0(d_1399_v2_): d_1399_v2_})
            index264_ = default__.safeIndex(361, (d_1449_v40_).length(0))
            (d_1449_v40_)[index264_] = default__.fm21(d_1452_v41_, (self).f25, d_1399_v2_, (d_1448_v39_).set(default__.safeIndex(169, len(d_1448_v39_)), d_1444_v35_), globalState)
            index265_ = default__.safeIndex(361, (d_1449_v40_).length(0))
            (d_1449_v40_)[index265_] = (self).f25
            d_1453_v42_: _dafny.Map
            d_1453_v42_ = _dafny.Map({d_1399_v2_: not(d_1399_v2_)})
            d_1454_v43_: D13
            d_1454_v43_ = D13_DC36(d_1453_v42_, d_1399_v2_, (self).f25)
            d_1455_v44_: _dafny.Map
            d_1455_v44_ = _dafny.Map({d_1454_v43_: 827})
            d_1456_v45_: D1
            d_1456_v45_ = D1_DC6(d_1444_v35_, d_1399_v2_, d_1448_v39_)
            d_1457_v46_: _dafny.Map
            d_1457_v46_ = _dafny.Map({d_1399_v2_: len(_dafny.SeqWithoutIsStrInference([len((d_1456_v45_).cf9), (self).f25]))})
            d_1399_v2_ = (((d_1447_v38_)[d_1399_v2_] if (d_1399_v2_) in (d_1447_v38_) else (d_1449_v40_)[default__.safeIndex(361, (d_1449_v40_).length(0))])) <= (((0) - (default__.fm18(globalState)) if d_1399_v2_ else ((d_1455_v44_)[d_1454_v43_] if (d_1454_v43_) in (d_1455_v44_) else len(d_1457_v46_))))
            index266_ = default__.safeIndex(361, (d_1449_v40_).length(0))
            (d_1449_v40_)[index266_] = 272
            d_1458_v47_: _dafny.MultiSet
            d_1458_v47_ = _dafny.MultiSet([len(d_1448_v39_)])
            if ((d_1458_v47_).intersection(_dafny.MultiSet([(self).f25]))).ispropersubset((d_1458_v47_ if d_1399_v2_ else _dafny.MultiSet([-675]))):
                d_1459_v48_: _dafny.Array
                nw239_ = _dafny.Array(False, 18)
                d_1459_v48_ = nw239_
                d_1460_v49_: _dafny.Set
                d_1460_v49_ = _dafny.Set({d_1459_v48_, d_1459_v48_})
                d_1461_v50_: _dafny.MultiSet
                d_1461_v50_ = _dafny.MultiSet([d_1460_v49_])
                d_1462_v51_: _dafny.MultiSet
                d_1462_v51_ = _dafny.MultiSet([d_1444_v35_])
                d_1463_v52_: _dafny.Map
                d_1463_v52_ = _dafny.Map({d_1462_v51_: d_1447_v38_})
                index267_ = default__.safeIndex(361, (d_1449_v40_).length(0))
                (d_1449_v40_)[index267_] = ((d_1461_v50_)[d_1460_v49_] if (d_1460_v49_) in (d_1461_v50_) else (((d_1463_v52_)[default__.fm55(d_1399_v2_, (d_1462_v51_).cardinality, len(d_1448_v39_), globalState)] if (default__.fm55(d_1399_v2_, (d_1462_v51_).cardinality, len(d_1448_v39_), globalState)) in (d_1463_v52_) else d_1447_v38_)).cardinality)
                (globalState).f23 = d_1444_v35_
                (globalState).f12 = default__.fm0(default__.safeModuloInt(len(d_1397_v0_), 571), globalState)
                d_1464_v53_: _dafny.Seq
                d_1464_v53_ = _dafny.SeqWithoutIsStrInference([d_1399_v2_, d_1399_v2_, d_1399_v2_, d_1399_v2_, d_1399_v2_])
                d_1465_v54_: _dafny.Map
                d_1465_v54_ = _dafny.Map({(0) - ((d_1449_v40_)[default__.safeIndex(361, (d_1449_v40_).length(0))]): (d_1449_v40_)[default__.safeIndex(361, (d_1449_v40_).length(0))]})
                d_1466_v55_: _dafny.Map
                d_1466_v55_ = _dafny.Map({(d_1449_v40_)[default__.safeIndex(361, (d_1449_v40_).length(0))]: _dafny.SeqWithoutIsStrInference([d_1399_v2_, False])})
                d_1467_v56_: D10
                d_1467_v56_ = D10_DC29(d_1399_v2_, d_1399_v2_, (self).f25)
                pat_let_tv30_ = d_1464_v53_
                d_1468_v57_: _dafny.Array
                nw240_ = _dafny.Array(None, 24)
                nw240_[int(0)] = d_1464_v53_
                nw240_[int(1)] = (d_1464_v53_).set(default__.safeIndex(len(d_1465_v54_), len(d_1464_v53_)), d_1399_v2_)
                nw240_[int(2)] = d_1464_v53_
                nw240_[int(3)] = d_1464_v53_
                def iife109_(_pat_let26_0):
                    def iife110_(d_1469_dt__update__tmp_h0_):
                        def iife111_(_pat_let27_0):
                            def iife112_(d_1470_dt__update_hcf6_h0_):
                                return D1_DC4(d_1470_dt__update_hcf6_h0_)
                            return iife112_(_pat_let27_0)
                        return iife111_(pat_let_tv30_)
                    return iife110_(_pat_let26_0)
                nw240_[int(4)] = (iife109_(p0)).cf6
                nw240_[int(5)] = (default__.fm1(d_1399_v2_, globalState)) + (_dafny.SeqWithoutIsStrInference([True, d_1399_v2_, d_1399_v2_]))
                nw240_[int(6)] = d_1464_v53_
                nw240_[int(7)] = d_1464_v53_
                nw240_[int(8)] = (D1_DC4(_dafny.SeqWithoutIsStrInference([d_1399_v2_]))).cf6
                nw240_[int(9)] = d_1464_v53_
                nw240_[int(10)] = d_1464_v53_
                nw240_[int(11)] = d_1464_v53_
                nw240_[int(12)] = d_1464_v53_
                nw240_[int(13)] = (d_1464_v53_) + (d_1464_v53_)
                nw240_[int(14)] = (_dafny.SeqWithoutIsStrInference([not(False)])) + (d_1464_v53_)
                nw240_[int(15)] = ((d_1466_v55_)[(d_1467_v56_).cf44] if ((d_1467_v56_).cf44) in (d_1466_v55_) else d_1464_v53_)
                nw240_[int(16)] = default__.fm1(d_1399_v2_, globalState)
                nw240_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1399_v2_, d_1399_v2_, (d_1464_v53_)[default__.safeIndex(537, len(d_1464_v53_))]])
                nw240_[int(18)] = d_1464_v53_
                nw240_[int(19)] = d_1464_v53_
                nw240_[int(20)] = d_1464_v53_
                nw240_[int(21)] = (d_1464_v53_) + (d_1464_v53_)
                nw240_[int(22)] = d_1464_v53_
                nw240_[int(23)] = d_1464_v53_
                d_1468_v57_ = nw240_
                index268_ = default__.safeIndex(148, (d_1468_v57_).length(0))
                (d_1468_v57_)[index268_] = (d_1464_v53_ if default__.fm0((self).f25, globalState) else d_1464_v53_)
                d_1471_v58_: D0
                d_1471_v58_ = D0_DC1(True, d_1399_v2_)
                index269_ = default__.safeIndex(148, (d_1468_v57_).length(0))
                rhs263_ = d_1399_v2_
                rhs264_ = default__.fm0((self).f25, globalState)
                rhs265_ = (d_1464_v53_) + (d_1464_v53_)
                rhs266_ = default__.fm56(d_1399_v2_, globalState)
                lhs191_ = globalState
                lhs192_ = d_1468_v57_
                lhs193_ = default__.safeIndex(148, (d_1468_v57_).length(0))
                lhs191_.f12 = rhs263_
                d_1399_v2_ = rhs264_
                lhs192_[lhs193_] = rhs265_
                d_1471_v58_ = rhs266_
                index270_ = default__.safeIndex(778, (d_1459_v48_).length(0))
                (d_1459_v48_)[index270_] = d_1399_v2_
                index271_ = default__.safeIndex(778, (d_1459_v48_).length(0))
                (d_1459_v48_)[index271_] = d_1399_v2_
            elif True:
                (globalState).f10 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrtsnphg"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcfgfvj"))) + (d_1448_v39_))
                (globalState).f10 = (d_1448_v39_) + (_dafny.SeqWithoutIsStrInference([d_1444_v35_ for d_1472_i1_ in range(default__.abs(438))]))
                index272_ = default__.safeIndex(361, (d_1449_v40_).length(0))
                (d_1449_v40_)[index272_] = (self).f25
                d_1449_v40_ = d_1449_v40_
                d_1473_v59_: _dafny.Seq
                d_1473_v59_ = _dafny.SeqWithoutIsStrInference([d_1448_v39_])
                index273_ = default__.safeIndex(680, ((d_1398_v1_).f36).length(0))
                ((d_1398_v1_).f36)[index273_] = ((d_1473_v59_)[default__.safeIndex((d_1449_v40_)[default__.safeIndex(361, (d_1449_v40_).length(0))], len(d_1473_v59_))]) + (_dafny.SeqWithoutIsStrInference([d_1444_v35_ for d_1474_i2_ in range(default__.abs(-695))]))
                index274_ = default__.safeIndex(680, ((d_1398_v1_).f36).length(0))
                ((d_1398_v1_).f36)[index274_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mipgqij"))).set(default__.safeIndex(((self).f25) + (len(d_1448_v39_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mipgqij")))), d_1444_v35_)
        d_1399_v2_ = d_1399_v2_
        d_1475_v60_: _dafny.Seq
        d_1475_v60_ = _dafny.SeqWithoutIsStrInference([d_1399_v2_, d_1399_v2_])
        d_1476_v61_: _dafny.Map
        d_1476_v61_ = _dafny.Map({(self).f25: not((d_1475_v60_)[default__.safeIndex((self).f25, len(d_1475_v60_))])})
        d_1476_v61_ = d_1476_v61_
        d_1477_v62_: _dafny.Array
        def lambda96_(d_1478_v2_):
            def lambda97_(d_1479_i3_):
                return (True if d_1478_v2_ else d_1478_v2_)

            return lambda97_

        init53_ = lambda96_(d_1399_v2_)
        nw241_ = _dafny.Array(None, 2)
        for i0_53_ in range(nw241_.length(0)):
            nw241_[i0_53_] = init53_(i0_53_)
        d_1477_v62_ = nw241_
        index275_ = default__.safeIndex(224, (d_1477_v62_).length(0))
        (d_1477_v62_)[index275_] = not(d_1399_v2_)
        index276_ = default__.safeIndex(224, (d_1477_v62_).length(0))
        (d_1477_v62_)[index276_] = d_1399_v2_
        d_1480_v63_: str
        d_1480_v63_ = _dafny.CodePoint('y')
        if (d_1480_v63_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fiqda"))):
            d_1481_v64_: _dafny.Map
            d_1481_v64_ = _dafny.Map({(self).f25: default__.fm18(globalState)})
            d_1481_v64_ = (d_1481_v64_).set((self).f25, (self).f25)
            d_1482_v65_: D9
            d_1482_v65_ = D9_DC27(len(_dafny.SeqWithoutIsStrInference([d_1480_v63_ for d_1483_i4_ in range(default__.abs(672))])), (self).f25)
            d_1484_v66_: _dafny.MultiSet
            d_1484_v66_ = _dafny.MultiSet([d_1482_v65_, d_1482_v65_, d_1482_v65_])
            d_1485_v67_: _dafny.Set
            d_1485_v67_ = _dafny.Set({d_1484_v66_})
            d_1485_v67_ = (d_1485_v67_).intersection(d_1485_v67_)
            d_1486_v68_: C2
            nw242_ = C2()
            nw242_.ctor__()
            d_1486_v68_ = nw242_
            d_1487_v69_: _dafny.Array
            nw243_ = _dafny.Array(_dafny.CodePoint('D'), 29)
            d_1487_v69_ = nw243_
            d_1488_v70_: _dafny.Seq
            d_1488_v70_ = _dafny.SeqWithoutIsStrInference([d_1487_v69_])
            d_1487_v69_ = (d_1488_v70_)[default__.safeIndex(((self).f25) - ((self).f25), len(d_1488_v70_))]
            d_1489_v71_: _dafny.Map
            d_1489_v71_ = _dafny.Map({_dafny.Map({(self).f25: (self).f25}): not(d_1399_v2_)})
            d_1490_v72_: _dafny.Seq
            d_1490_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_1489_v71_ = (d_1489_v71_).set((d_1481_v64_) | (default__.fm57((d_1477_v62_)[default__.safeIndex(224, (d_1477_v62_).length(0))], default__.fm0((self).f25, globalState), (self).f25, globalState)), (d_1480_v63_) in (d_1490_v72_))
        elif True:
            d_1491_v73_: C2
            nw244_ = C2()
            nw244_.ctor__()
            d_1491_v73_ = nw244_
            index277_ = default__.safeIndex(224, (d_1477_v62_).length(0))
            rhs267_ = not((d_1477_v62_)[default__.safeIndex(224, (d_1477_v62_).length(0))])
            rhs268_ = d_1477_v62_
            lhs194_ = d_1477_v62_
            lhs195_ = default__.safeIndex(224, (d_1477_v62_).length(0))
            lhs194_[lhs195_] = rhs267_
            d_1477_v62_ = rhs268_
            d_1492_v74_: _dafny.Map
            d_1492_v74_ = _dafny.Map({d_1399_v2_: d_1399_v2_})
            d_1493_v75_: _dafny.Map
            d_1493_v75_ = _dafny.Map({(self).f25: d_1492_v74_})
            d_1493_v75_ = (d_1493_v75_).set((d_1491_v73_).fm24((self).f25, globalState), d_1492_v74_)
            d_1494_v76_: int
            d_1494_v76_ = 413
            d_1494_v76_ = (self).f25
            d_1495_v77_: _dafny.Array
            def lambda98_(d_1496_v76_):
                def lambda99_(d_1497_i5_):
                    return default__.safeModuloInt(d_1497_i5_, d_1496_v76_)

                return lambda99_

            init54_ = lambda98_(d_1494_v76_)
            nw245_ = _dafny.Array(None, 8)
            for i0_54_ in range(nw245_.length(0)):
                nw245_[i0_54_] = init54_(i0_54_)
            d_1495_v77_ = nw245_
            d_1498_v78_: _dafny.Map
            d_1498_v78_ = _dafny.Map({d_1495_v77_: 452})
            index278_ = default__.safeIndex(956, (d_1495_v77_).length(0))
            (d_1495_v77_)[index278_] = ((self).f25) * (((d_1498_v78_)[d_1495_v77_] if (d_1495_v77_) in (d_1498_v78_) else 296))
            d_1499_v79_: _dafny.Set
            d_1499_v79_ = _dafny.Set({False, d_1399_v2_, d_1399_v2_, d_1399_v2_, d_1399_v2_})
            d_1500_v80_: _dafny.Seq
            d_1500_v80_ = _dafny.SeqWithoutIsStrInference([d_1499_v79_, d_1499_v79_])
            index279_ = default__.safeIndex(956, (d_1495_v77_).length(0))
            (d_1495_v77_)[index279_] = (len((d_1499_v79_).intersection((d_1500_v80_)[default__.safeIndex(d_1494_v76_, len(d_1500_v80_))]))) - (d_1494_v76_)

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        if p1:
            r0 = (0) - ((0) - ((self).f25))
            d_1501_v0_: _dafny.Seq
            d_1501_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olyddmha"))
            d_1502_v1_: C1
            nw246_ = C1()
            nw246_.ctor__(default__.fm18(globalState), (self).f26)
            d_1502_v1_ = nw246_
            d_1503_v2_: _dafny.MultiSet
            d_1503_v2_ = _dafny.MultiSet([d_1502_v1_, d_1502_v1_])
            d_1504_v3_: _dafny.Map
            d_1504_v3_ = _dafny.Map({(len(d_1501_v0_)) + ((self).f25): d_1503_v2_})
            d_1504_v3_ = (d_1504_v3_).set((self).f25, d_1503_v2_)
            r0 = (self).f25
            (globalState).f12 = (d_1501_v0_) < ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1505_i0_ in range(default__.abs(444))])) + (d_1501_v0_))
            index280_ = default__.safeIndex(46, (p3).length(0))
            (p3)[index280_] = (self).f25
            index281_ = default__.safeIndex(46, (p3).length(0))
            (p3)[index281_] = (self).f25
        elif True:
            d_1506_v4_: _dafny.MultiSet
            d_1506_v4_ = _dafny.MultiSet([False, p1])
            d_1507_v5_: _dafny.Map
            d_1507_v5_ = _dafny.Map({p1: p1})
            d_1508_v6_: _dafny.Set
            d_1508_v6_ = _dafny.Set({(d_1506_v4_).issubset(d_1506_v4_), ((d_1507_v5_)[p1] if (p1) in (d_1507_v5_) else p2)})
            d_1508_v6_ = ((d_1508_v6_).intersection(d_1508_v6_)).intersection(d_1508_v6_)
            d_1509_v7_: _dafny.Seq
            d_1509_v7_ = _dafny.SeqWithoutIsStrInference([p1])
            (globalState).f12 = (_dafny.Set({p2, (d_1509_v7_)[default__.safeIndex((self).f25, len(d_1509_v7_))], p1, p2})).issubset((_dafny.Set({p1, p2})) - (d_1508_v6_))
            d_1510_v8_: _dafny.Array
            nw247_ = _dafny.Array(D10.default()(), 16)
            d_1510_v8_ = nw247_
            d_1510_v8_ = d_1510_v8_
            d_1511_v9_: D0
            d_1511_v9_ = D0_DC1(p2, p1)
            d_1512_v10_: _dafny.Array
            nw248_ = _dafny.Array(None, 21)
            nw248_[int(0)] = p1
            nw248_[int(1)] = p2
            nw248_[int(2)] = p1
            nw248_[int(3)] = p1
            nw248_[int(4)] = p2
            nw248_[int(5)] = p1
            nw248_[int(6)] = p1
            nw248_[int(7)] = False
            nw248_[int(8)] = p2
            nw248_[int(9)] = p1
            nw248_[int(10)] = False
            nw248_[int(11)] = p1
            nw248_[int(12)] = (d_1511_v9_).cf1
            nw248_[int(13)] = p2
            nw248_[int(14)] = p2
            nw248_[int(15)] = p2
            nw248_[int(16)] = p1
            nw248_[int(17)] = not(p1)
            nw248_[int(18)] = p2
            nw248_[int(19)] = p1
            nw248_[int(20)] = p1
            d_1512_v10_ = nw248_
            d_1513_v11_: _dafny.Map
            d_1513_v11_ = _dafny.Map({p2: d_1512_v10_})
            d_1514_v12_: D4
            d_1514_v12_ = D4_DC13((self).f25, not(p1), p1, True)
            d_1513_v11_ = (d_1513_v11_).set(not (p2) or ((d_1514_v12_).cf21), d_1512_v10_)
            d_1515_v15_: _dafny.Array
            def lambda100_(d_1516_i1_):
                def iife113_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in (_dafny.Set({(_dafny.MultiSet([(self).f25, (self).f25, (self).f25])).cardinality})).Elements:
                        d_1517_v13_: int = compr_57_
                        if (d_1517_v13_) in (_dafny.Set({(_dafny.MultiSet([(self).f25, (self).f25, (self).f25])).cardinality})):
                            def iife114_():
                                coll58_ = _dafny.Map()
                                compr_58_: int
                                for compr_58_ in _dafny.IntegerRange(384, 189):
                                    d_1518_v14_: int = compr_58_
                                    if ((384) <= (d_1518_v14_)) and ((d_1518_v14_) < (189)):
                                        coll58_[default__.safeModuloInt(d_1518_v14_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('i'), _dafny.CodePoint('e')])))] = False
                                return _dafny.Map(coll58_)
                            coll57_ = coll57_.union(_dafny.Set([default__.safeModuloInt(d_1517_v13_, len(iife114_()
))]))
                    return _dafny.Set(coll57_)
                return (iife113_()
                ) - (_dafny.Set({(self).f25}))

            init55_ = lambda100_
            nw249_ = _dafny.Array(None, 19)
            for i0_55_ in range(nw249_.length(0)):
                nw249_[i0_55_] = init55_(i0_55_)
            d_1515_v15_ = nw249_
            d_1519_v16_: _dafny.Set
            d_1519_v16_ = _dafny.Set({(self).f25, (self).f25, (self).f25})
            index282_ = default__.safeIndex(829, (d_1515_v15_).length(0))
            (d_1515_v15_)[index282_] = d_1519_v16_
            d_1520_v17_: _dafny.Map
            d_1520_v17_ = _dafny.Map({(self).f25: d_1519_v16_})
            d_1521_v18_: _dafny.Seq
            d_1521_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbpat"))
            index283_ = default__.safeIndex(829, (d_1515_v15_).length(0))
            (d_1515_v15_)[index283_] = ((d_1520_v17_)[default__.fm31(d_1521_v18_, globalState)] if (default__.fm31(d_1521_v18_, globalState)) in (d_1520_v17_) else d_1519_v16_)
        if p1:
            d_1522_v19_: _dafny.Seq
            d_1522_v19_ = _dafny.SeqWithoutIsStrInference([p2, p2, p1, p2, True])
            d_1523_v20_: _dafny.Map
            d_1523_v20_ = _dafny.Map({True: p2})
            d_1524_v21_: _dafny.Seq
            d_1524_v21_ = _dafny.SeqWithoutIsStrInference([d_1522_v19_, d_1522_v19_, (d_1522_v19_ if p2 else _dafny.SeqWithoutIsStrInference([((d_1523_v20_)[True] if (True) in (d_1523_v20_) else p2), p2, p1, p2, ((d_1523_v20_)[p1] if (p1) in (d_1523_v20_) else p1)]))])
            d_1524_v21_ = (d_1524_v21_).set(default__.safeIndex((self).f25, len(d_1524_v21_)), d_1522_v19_)
            d_1525_v22_: T0
            nw250_ = C4()
            nw250_.ctor__((self).f26, (0) - ((0) - (default__.safeDivisionInt((self).f25, (self).f25))), (self).f26)
            d_1525_v22_ = nw250_
            d_1525_v22_ = d_1525_v22_
            d_1526_v23_: _dafny.Array
            nw251_ = _dafny.Array(False, 29)
            d_1526_v23_ = nw251_
            d_1527_v24_: D8
            d_1527_v24_ = D8_DC23(_dafny.Map({d_1526_v23_: p2}))
            d_1528_v25_: D8
            d_1528_v25_ = D8_DC25(d_1527_v24_)
            d_1529_v26_: D8
            d_1529_v26_ = D8_DC25(D8_DC25(D8_DC25(d_1528_v25_)))
            d_1530_v27_: _dafny.Map
            d_1530_v27_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esapojtap")): d_1529_v26_})
            d_1530_v27_ = (d_1530_v27_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxdasflhy")), d_1529_v26_)
            d_1531_v28_: _dafny.Seq
            d_1531_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
            (globalState).f10 = d_1531_v28_
            r0 = default__.fm21(_dafny.Map({D0_DC0(p2): True}), (self).f25, (p2) in (d_1522_v19_), d_1531_v28_, globalState)
        elif True:
            (globalState).f12 = p2
            (globalState).f12 = False
            if p2:
                d_1532_v29_: _dafny.Map
                d_1532_v29_ = _dafny.Map({((self).f25 if p2 else (self).f25): (self).f25})
                d_1533_v30_: _dafny.Array
                def lambda101_(d_1534_p1_):
                    def lambda102_(d_1535_i2_):
                        return d_1534_p1_

                    return lambda102_

                init56_ = lambda101_(p1)
                nw252_ = _dafny.Array(None, 2)
                for i0_56_ in range(nw252_.length(0)):
                    nw252_[i0_56_] = init56_(i0_56_)
                d_1533_v30_ = nw252_
                d_1536_v31_: _dafny.Map
                d_1536_v31_ = _dafny.Map({p2: d_1533_v30_})
                d_1537_v32_: _dafny.Set
                d_1537_v32_ = _dafny.Set({((d_1536_v31_)[False] if (False) in (d_1536_v31_) else d_1533_v30_)})
                d_1538_v33_: _dafny.Seq
                d_1538_v33_ = _dafny.SeqWithoutIsStrInference([d_1537_v32_])
                d_1539_v34_: D3
                d_1539_v34_ = D3_DC9((d_1538_v33_)[default__.safeIndex((self).f25, len(d_1538_v33_))])
                d_1540_v35_: _dafny.Seq
                d_1540_v35_ = _dafny.SeqWithoutIsStrInference([D3_DC9(d_1537_v32_), d_1539_v34_, d_1539_v34_])
                d_1532_v29_ = (d_1532_v29_).set(default__.safeModuloInt((self).f25, (self).f25), ((self).f25) * (len(d_1540_v35_)))
                d_1541_v36_: _dafny.Seq
                d_1541_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfim"))
                index284_ = default__.safeIndex(586, (p3).length(0))
                (p3)[index284_] = len(d_1541_v36_)
                index285_ = default__.safeIndex(586, (p3).length(0))
                (p3)[index285_] = 664
                d_1542_v37_: _dafny.Map
                d_1542_v37_ = _dafny.Map({d_1541_v36_: (self).f25})
                d_1543_v38_: D5
                d_1543_v38_ = D5_DC16(p1)
                d_1544_v39_: D6
                d_1544_v39_ = D6_DC19(D0_DC3(len(d_1542_v37_), (d_1543_v38_).cf25), -841)
                r0 = (d_1544_v39_).cf30
                index286_ = default__.safeIndex(288, ((self).f26).length(0))
                ((self).f26)[index286_] = d_1541_v36_
                d_1545_v40_: _dafny.Seq
                d_1545_v40_ = _dafny.SeqWithoutIsStrInference([d_1541_v36_, d_1541_v36_])
                d_1546_v41_: _dafny.Map
                d_1546_v41_ = _dafny.Map({(0) - ((p3)[default__.safeIndex(586, (p3).length(0))]): p1})
                index287_ = default__.safeIndex(288, ((self).f26).length(0))
                ((self).f26)[index287_] = (d_1545_v40_)[default__.safeIndex(default__.safeModuloInt(len(d_1546_v41_), (p3)[default__.safeIndex(586, (p3).length(0))]), len(d_1545_v40_))]
                (globalState).f12 = p1
            elif True:
                d_1547_v42_: C2
                nw253_ = C2()
                nw253_.ctor__()
                d_1547_v42_ = nw253_
                d_1547_v42_ = d_1547_v42_
                d_1548_v43_: _dafny.Array
                nw254_ = _dafny.Array(False, 23)
                d_1548_v43_ = nw254_
                d_1548_v43_ = d_1548_v43_
                index288_ = default__.safeIndex(948, (d_1548_v43_).length(0))
                (d_1548_v43_)[index288_] = p2
                index289_ = default__.safeIndex(948, (d_1548_v43_).length(0))
                rhs269_ = p1
                rhs270_ = ((self).f25) - ((self).f25)
                rhs271_ = (0) - ((self).f25)
                rhs272_ = p2
                lhs196_ = globalState
                lhs197_ = d_1548_v43_
                lhs198_ = default__.safeIndex(948, (d_1548_v43_).length(0))
                lhs196_.f12 = rhs269_
                r1 = rhs270_
                r1 = rhs271_
                lhs197_[lhs198_] = rhs272_
                d_1549_v44_: _dafny.Seq
                d_1549_v44_ = _dafny.SeqWithoutIsStrInference([928])
                d_1550_v45_: _dafny.Map
                d_1550_v45_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1551_i3_ in range(default__.abs(-80))]): p2})
                d_1552_v46_: _dafny.Seq
                d_1552_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                d_1553_v47_: _dafny.Set
                d_1553_v47_ = _dafny.Set({p1, default__.fm0((self).f25, globalState), ((d_1550_v45_)[d_1552_v46_] if (d_1552_v46_) in (d_1550_v45_) else (d_1548_v43_)[default__.safeIndex(948, (d_1548_v43_).length(0))]), not(p2), (d_1548_v43_)[default__.safeIndex(948, (d_1548_v43_).length(0))]})
                d_1554_v48_: _dafny.Map
                d_1554_v48_ = _dafny.Map({(_dafny.Set({(d_1548_v43_)[default__.safeIndex(948, (d_1548_v43_).length(0))], default__.fm0(default__.fm44(default__.fm0((d_1549_v44_)[default__.safeIndex(160, len(d_1549_v44_))], globalState), (d_1548_v43_)[default__.safeIndex(948, (d_1548_v43_).length(0))], p1, (self).f25, globalState), globalState)})) | (d_1553_v47_): d_1549_v44_})
                d_1554_v48_ = d_1554_v48_
                d_1555_v49_: _dafny.Array
                def lambda103_(d_1556_i4_):
                    return (d_1556_i4_) * ((self).f25)

                init57_ = lambda103_
                nw255_ = _dafny.Array(None, 10)
                for i0_57_ in range(nw255_.length(0)):
                    nw255_[i0_57_] = init57_(i0_57_)
                d_1555_v49_ = nw255_
                d_1555_v49_ = d_1555_v49_
            d_1557_v50_: _dafny.Array
            nw256_ = _dafny.Array(None, 12)
            nw256_[int(0)] = default__.fm0((self).f25, globalState)
            nw256_[int(1)] = p2
            nw256_[int(2)] = not(p2)
            nw256_[int(3)] = not(p2)
            nw256_[int(4)] = p1
            nw256_[int(5)] = p2
            nw256_[int(6)] = p2
            nw256_[int(7)] = p2
            nw256_[int(8)] = p1
            nw256_[int(9)] = p2
            nw256_[int(10)] = False
            nw256_[int(11)] = p2
            d_1557_v50_ = nw256_
            d_1558_v51_: _dafny.Seq
            d_1558_v51_ = _dafny.SeqWithoutIsStrInference([d_1557_v50_])
            d_1559_v52_: str
            d_1559_v52_ = _dafny.CodePoint('y')
            d_1560_v53_: _dafny.Seq
            d_1560_v53_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1559_v52_, d_1559_v52_, d_1559_v52_])])
            d_1561_v54_: D0
            d_1561_v54_ = D0_DC0(p1)
            d_1562_v55_: _dafny.Map
            d_1562_v55_ = _dafny.Map({d_1561_v54_: p1})
            d_1563_v56_: _dafny.Seq
            d_1563_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpr"))
            d_1564_v58_: _dafny.Seq
            d_1564_v58_ = _dafny.SeqWithoutIsStrInference([d_1561_v54_])
            def iife115_():
                coll59_ = _dafny.Map()
                compr_59_: D0
                for compr_59_ in (d_1564_v58_).Elements:
                    d_1565_v57_: D0 = compr_59_
                    if (d_1565_v57_) in (d_1564_v58_):
                        coll59_[d_1565_v57_] = p2
                return _dafny.Map(coll59_)
            rhs273_ = default__.fm21(d_1562_v55_, default__.safeModuloInt((self).f25, len(d_1560_v53_)), (p1) or (p2), d_1563_v56_, globalState)
            rhs274_ = (d_1558_v51_).set(default__.safeIndex((self).f25, len(d_1558_v51_)), d_1557_v50_)
            rhs275_ = (_dafny.SeqWithoutIsStrInference([d_1563_v56_, _dafny.SeqWithoutIsStrInference([d_1559_v52_, default__.fm39(p1, globalState), d_1559_v52_, d_1559_v52_, d_1559_v52_])])) + (default__.fm58(d_1563_v56_, default__.fm21(iife115_()
            , (self).f25, p1, (d_1563_v56_).set(default__.safeIndex((self).f25, len(d_1563_v56_)), d_1559_v52_), globalState), p1, (self).f25, globalState))
            rhs276_ = (self).f25
            r1 = rhs273_
            d_1558_v51_ = rhs274_
            d_1560_v53_ = rhs275_
            r1 = rhs276_
            d_1566_v59_: D12
            d_1566_v59_ = D12_DC32(p1)
            d_1567_v60_: C3
            nw257_ = C3()
            nw257_.ctor__((self).f25, (d_1566_v59_).cf47, (self).f25, (self).f26)
            d_1567_v60_ = nw257_
        index290_ = default__.safeIndex(853, (p3).length(0))
        (p3)[index290_] = default__.safeModuloInt((self).f25, (self).f25)
        d_1568_v61_: _dafny.Map
        d_1568_v61_ = _dafny.Map({not(p1): (self).f25})
        d_1569_v62_: D9
        d_1569_v62_ = D9_DC26(d_1568_v61_)
        index291_ = default__.safeIndex(853, (p3).length(0))
        def lambda104_(source22_):
            if source22_.is_DC27:
                d_1570___mcc_h0_ = source22_.cf39
                d_1571___mcc_h1_ = source22_.cf40
                d_1572_cf40_ = d_1571___mcc_h1_
                d_1573_cf39_ = d_1570___mcc_h0_
                return (_dafny.MultiSet([(self).f25, d_1572_cf40_])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1572_cf40_: d_1573_cf39_}))])))
            elif True:
                d_1574___mcc_h2_ = source22_.cf38
                d_1575_cf38_ = d_1574___mcc_h2_
                return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjvuwsmtm"))), 183, (0) - ((self).f25), (0) - ((self).f25), (self).f25])

        (p3)[index291_] = (0) - ((lambda104_(d_1569_v62_)).cardinality)
        d_1576_v63_: _dafny.Map
        d_1576_v63_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfbukxl")): 223})
        d_1577_v64_: _dafny.Seq
        d_1577_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
        d_1578_v65_: _dafny.Map
        d_1578_v65_ = _dafny.Map({(0) - ((self).f25): (self).f25})
        d_1579_v66_: _dafny.Set
        d_1579_v66_ = _dafny.Set({((d_1578_v65_)[467] if (467) in (d_1578_v65_) else (self).f25), len(d_1577_v64_)})
        d_1580_v67_: D4
        d_1580_v67_ = D4_DC11(d_1579_v66_)
        d_1576_v63_ = (d_1576_v63_) | (default__.fm59(d_1577_v64_, d_1580_v67_, globalState))
        d_1581_v68_: _dafny.Seq
        d_1581_v68_ = _dafny.SeqWithoutIsStrInference([(d_1577_v64_) + (d_1577_v64_), d_1577_v64_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1582_i5_ in range(default__.abs(373))])])
        d_1577_v64_ = (d_1581_v68_)[default__.safeIndex((0) - (default__.fm44(p2, p1, True, (self).f25, globalState)), len(d_1581_v68_))]
        d_1583_v69_: _dafny.MultiSet
        d_1583_v69_ = _dafny.MultiSet([(p3)[default__.safeIndex(853, (p3).length(0))]])
        d_1583_v69_ = _dafny.MultiSet([(p3)[default__.safeIndex(853, (p3).length(0))], (p3)[default__.safeIndex(853, (p3).length(0))], (p3)[default__.safeIndex(853, (p3).length(0))]])
        r0 = (p3)[default__.safeIndex(853, (p3).length(0))]
        r1 = (default__.safeModuloInt((p3)[default__.safeIndex(853, (p3).length(0))], (p3)[default__.safeIndex(853, (p3).length(0))])) * (default__.fm44(False, not(p2), p1, 7, globalState))
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        (globalState).f12 = ((self).f25) != (702)
        r0 = (0) - ((701) - ((self).f25))
        d_1584_v0_: C0
        nw258_ = C0()
        nw258_.ctor__()
        d_1584_v0_ = nw258_
        d_1585_i0_: int
        d_1585_i0_ = 0
        with _dafny.label("11"):
            while default__.fm0((self).f25, globalState):
                with _dafny.c_label("11"):
                    if (d_1585_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1585_i0_ = (d_1585_i0_) + (1)
                    r0 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwkghx")))
                    (globalState).f12 = p1
                    d_1586_v1_: _dafny.Array
                    nw259_ = _dafny.Array(int(0), 21)
                    d_1586_v1_ = nw259_
                    index292_ = default__.safeIndex(988, (d_1586_v1_).length(0))
                    (d_1586_v1_)[index292_] = default__.safeDivisionInt((self).f25, (0) - ((self).f25))
                    d_1587_v2_: _dafny.Seq
                    d_1587_v2_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25])
                    d_1588_v3_: _dafny.Map
                    d_1588_v3_ = _dafny.Map({(0) - ((self).f25): p1})
                    index293_ = default__.safeIndex(988, (d_1586_v1_).length(0))
                    (d_1586_v1_)[index293_] = (513) + (((d_1587_v2_)[default__.safeIndex(len(d_1588_v3_), len(d_1587_v2_))] if p1 else (self).f25))
                    (globalState).f12 = p1
                    pass
            pass
        d_1589_v4_: _dafny.Set
        d_1589_v4_ = _dafny.Set({(self).f25, (self).f25})
        r0 = len((d_1589_v4_) | (d_1589_v4_))
        d_1590_v5_: C0
        nw260_ = C0()
        nw260_.ctor__()
        d_1590_v5_ = nw260_
        r0 = default__.safeDivisionInt(-364, (self).f25)
        return r0

    def m11(self, p0, p1, p2, globalState):
        r0: D3 = D3.default()()
        r1: bool = False
        d_1591_v0_: str
        d_1591_v0_ = _dafny.CodePoint('q')
        d_1592_v1_: _dafny.Seq
        d_1592_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxtquho"))
        d_1593_v2_: D1
        d_1593_v2_ = D1_DC6(d_1591_v0_, p1, d_1592_v1_)
        source23_ = d_1593_v2_
        if source23_.is_DC5:
            r1 = p1
            (globalState).f12 = p1
            d_1594_v3_: _dafny.Seq
            d_1594_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            d_1594_v3_ = d_1594_v3_
            if p1:
                d_1595_v4_: int
                d_1595_v4_ = 834
                d_1595_v4_ = -982
                (globalState).f12 = (291) != (d_1595_v4_)
                d_1595_v4_ = d_1595_v4_
                d_1595_v4_ = d_1595_v4_
                d_1596_v5_: _dafny.MultiSet
                d_1596_v5_ = _dafny.MultiSet([p1])
                d_1596_v5_ = _dafny.MultiSet([(d_1592_v1_) < (d_1592_v1_)])
            elif True:
                d_1597_v6_: D0
                d_1597_v6_ = D0_DC3((self).f25, p1)
                d_1598_v7_: D6
                def iife116_(_pat_let28_0):
                    def iife117_(d_1599_dt__update__tmp_h0_):
                        def iife118_(_pat_let29_0):
                            def iife119_(d_1600_dt__update_hcf4_h0_):
                                return D0_DC3(d_1600_dt__update_hcf4_h0_, (d_1599_dt__update__tmp_h0_).cf5)
                            return iife119_(_pat_let29_0)
                        return iife118_((self).f25)
                    return iife117_(_pat_let28_0)
                d_1598_v7_ = D6_DC19(iife116_(d_1597_v6_), (default__.fm44(p1, p1, p1, (self).f25, globalState)) * (915))
                d_1598_v7_ = default__.fm60(globalState)
                d_1601_v8_: _dafny.Array
                def lambda105_(d_1602_p1_):
                    def lambda106_(d_1603_i0_):
                        return (d_1602_p1_) or (d_1602_p1_)

                    return lambda106_

                init58_ = lambda105_(p1)
                nw261_ = _dafny.Array(None, 18)
                for i0_58_ in range(nw261_.length(0)):
                    nw261_[i0_58_] = init58_(i0_58_)
                d_1601_v8_ = nw261_
                index294_ = default__.safeIndex(491, (d_1601_v8_).length(0))
                (d_1601_v8_)[index294_] = p1
                d_1604_v9_: T0
                nw262_ = C4()
                nw262_.ctor__((self).f26, (self).f25, (self).f26)
                d_1604_v9_ = nw262_
                index295_ = default__.safeIndex(339, (d_1601_v8_).length(0))
                (d_1601_v8_)[index295_] = False
                d_1605_v10_: _dafny.Set
                d_1605_v10_ = _dafny.Set({not(p1), p1})
                d_1606_v11_: _dafny.MultiSet
                d_1606_v11_ = _dafny.MultiSet([d_1591_v0_, d_1591_v0_, d_1591_v0_])
                index296_ = default__.safeIndex(491, (d_1601_v8_).length(0))
                index297_ = default__.safeIndex(339, (d_1601_v8_).length(0))
                rhs277_ = (len((d_1605_v10_) - (d_1605_v10_))) < ((d_1604_v9_).f25)
                rhs278_ = (d_1604_v9_ if ((d_1606_v11_).cardinality) < ((d_1604_v9_).f25) else d_1604_v9_)
                rhs279_ = p1
                lhs199_ = d_1601_v8_
                lhs200_ = default__.safeIndex(491, (d_1601_v8_).length(0))
                lhs201_ = d_1601_v8_
                lhs202_ = default__.safeIndex(339, (d_1601_v8_).length(0))
                lhs199_[lhs200_] = rhs277_
                d_1604_v9_ = rhs278_
                lhs201_[lhs202_] = rhs279_
                index298_ = default__.safeIndex(491, (d_1601_v8_).length(0))
                rhs280_ = (d_1601_v8_)[default__.safeIndex(491, (d_1601_v8_).length(0))]
                rhs281_ = p1
                rhs282_ = (d_1601_v8_)[default__.safeIndex(491, (d_1601_v8_).length(0))]
                rhs283_ = default__.fm0(-567, globalState)
                rhs284_ = (d_1591_v0_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muieyuiqy")))
                lhs203_ = globalState
                lhs204_ = globalState
                lhs205_ = d_1601_v8_
                lhs206_ = default__.safeIndex(491, (d_1601_v8_).length(0))
                lhs207_ = globalState
                lhs203_.f12 = rhs280_
                lhs204_.f12 = rhs281_
                lhs205_[lhs206_] = rhs282_
                lhs207_.f12 = rhs283_
                r1 = rhs284_
                (globalState).f10 = ((d_1592_v1_) + (d_1592_v1_)) + ((d_1592_v1_) + (d_1592_v1_))
                d_1607_v12_: _dafny.Map
                d_1607_v12_ = _dafny.Map({True: p0})
                d_1607_v12_ = (d_1607_v12_).set(p1, p0)
        elif source23_.is_DC6:
            d_1608___mcc_h0_ = source23_.cf7
            d_1609___mcc_h1_ = source23_.cf8
            d_1610___mcc_h2_ = source23_.cf9
            d_1611_cf9_ = d_1610___mcc_h2_
            d_1612_cf8_ = d_1609___mcc_h1_
            d_1613_cf7_ = d_1608___mcc_h0_
            d_1614_v13_: int
            d_1614_v13_ = 873
            d_1614_v13_ = default__.safeDivisionInt(default__.fm18(globalState), (self).f25)
            d_1614_v13_ = (self).f25
            (globalState).f12 = d_1612_cf8_
            d_1615_v14_: _dafny.MultiSet
            d_1615_v14_ = _dafny.MultiSet([len(_dafny.Map({d_1612_cf8_: p1}))])
            d_1615_v14_ = d_1615_v14_
        elif source23_.is_DC4:
            d_1616___mcc_h3_ = source23_.cf6
            d_1617_cf6_ = d_1616___mcc_h3_
            if p1:
                d_1618_v15_: int
                d_1618_v15_ = 903
                d_1619_v16_: _dafny.Array
                nw263_ = _dafny.Array(False, 19)
                d_1619_v16_ = nw263_
                d_1620_v17_: _dafny.MultiSet
                d_1620_v17_ = _dafny.MultiSet([d_1618_v15_, d_1618_v15_, (self).f25])
                index299_ = default__.safeIndex(842, (d_1619_v16_).length(0))
                (d_1619_v16_)[index299_] = (d_1617_cf6_)[default__.safeIndex(((d_1620_v17_)[(self).f25] if ((self).f25) in (d_1620_v17_) else d_1618_v15_), len(d_1617_cf6_))]
                index300_ = default__.safeIndex(708, (p0).length(0))
                (p0)[index300_] = d_1618_v15_
                d_1621_v18_: D10
                d_1621_v18_ = D10_DC29(p1, p1, (self).f25)
                index301_ = default__.safeIndex(842, (d_1619_v16_).length(0))
                index302_ = default__.safeIndex(708, (p0).length(0))
                rhs285_ = (0) - ((D14_DC38(d_1618_v15_)).cf58)
                rhs286_ = p1
                rhs287_ = (0) - (((self).f25) + ((d_1621_v18_).cf44))
                rhs288_ = (self).f25
                rhs289_ = default__.fm0(default__.safeDivisionInt(d_1618_v15_, (self).f25), globalState)
                lhs208_ = d_1619_v16_
                lhs209_ = default__.safeIndex(842, (d_1619_v16_).length(0))
                lhs210_ = p0
                lhs211_ = default__.safeIndex(708, (p0).length(0))
                lhs212_ = globalState
                d_1618_v15_ = rhs285_
                lhs208_[lhs209_] = rhs286_
                lhs210_[lhs211_] = rhs287_
                d_1618_v15_ = rhs288_
                lhs212_.f12 = rhs289_
                d_1622_v19_: _dafny.MultiSet
                d_1622_v19_ = _dafny.MultiSet([(d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))], p1, p1, p1, (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]])
                d_1623_v20_: _dafny.Map
                d_1623_v20_ = _dafny.Map({(d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]: ((d_1622_v19_)[p1] if (p1) in (d_1622_v19_) else 279)})
                d_1623_v20_ = (d_1623_v20_).set(p1, -990)
                d_1624_v22_: _dafny.Map
                def iife120_():
                    coll60_ = _dafny.Set()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(621, 454):
                        d_1625_v21_: int = compr_60_
                        if ((621) <= (d_1625_v21_)) and ((d_1625_v21_) < (454)):
                            coll60_ = coll60_.union(_dafny.Set([(d_1625_v21_) - (969)]))
                    return _dafny.Set(coll60_)
                d_1624_v22_ = _dafny.Map({True: _dafny.SeqWithoutIsStrInference([iife120_()
 for d_1626_i1_ in range(default__.abs(212))])})
                d_1627_v23_: _dafny.Set
                d_1627_v23_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyq")))})
                d_1628_v24_: _dafny.Seq
                d_1628_v24_ = _dafny.SeqWithoutIsStrInference([d_1627_v23_, d_1627_v23_, d_1627_v23_, d_1627_v23_])
                d_1629_v26_: _dafny.Set
                d_1629_v26_ = _dafny.Set({default__.fm47((d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))], (p0)[default__.safeIndex(708, (p0).length(0))], d_1618_v15_, d_1591_v0_, globalState)})
                d_1630_v28_: _dafny.Map
                d_1630_v28_ = _dafny.Map({(p0)[default__.safeIndex(708, (p0).length(0))]: d_1618_v15_})
                d_1631_v30_: _dafny.Seq
                d_1631_v30_ = _dafny.SeqWithoutIsStrInference([default__.fm44(True, p1, (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))], (self).f25, globalState)])
                d_1632_v34_: _dafny.Array
                nw264_ = _dafny.Array(None, 29)
                def iife121_():
                    coll61_ = _dafny.Set()
                    compr_61_: _dafny.Set
                    for compr_61_ in (((d_1624_v22_)[(d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]] if ((d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]) in (d_1624_v22_) else d_1628_v24_)).Elements:
                        d_1633_v25_: _dafny.Set = compr_61_
                        if (d_1633_v25_) in (((d_1624_v22_)[(d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]] if ((d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]) in (d_1624_v22_) else d_1628_v24_)):
                            coll61_ = coll61_.union(_dafny.Set([d_1633_v25_]))
                    return _dafny.Set(coll61_)
                nw264_[int(0)] = iife121_()
                
                nw264_[int(1)] = (d_1629_v26_ if (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))] else _dafny.Set({(d_1628_v24_)[default__.safeIndex(d_1618_v15_, len(d_1628_v24_))]}))
                nw264_[int(2)] = d_1629_v26_
                nw264_[int(3)] = (d_1629_v26_) - (d_1629_v26_)
                nw264_[int(4)] = d_1629_v26_
                nw264_[int(5)] = d_1629_v26_
                nw264_[int(6)] = d_1629_v26_
                nw264_[int(7)] = d_1629_v26_
                nw264_[int(8)] = d_1629_v26_
                nw264_[int(9)] = (d_1629_v26_) - (d_1629_v26_)
                nw264_[int(10)] = d_1629_v26_
                nw264_[int(11)] = d_1629_v26_
                def iife122_():
                    coll62_ = _dafny.Set()
                    compr_62_: int
                    for compr_62_ in (d_1627_v23_).Elements:
                        d_1634_v27_: int = compr_62_
                        if (d_1634_v27_) in (d_1627_v23_):
                            coll62_ = coll62_.union(_dafny.Set([default__.safeDivisionInt(d_1634_v27_, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohtqnsan"))): (_dafny.MultiSet([383, len(_dafny.Map({len(_dafny.Map({True: False})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvft")))})), -285])).cardinality})))]))
                    return _dafny.Set(coll62_)
                def iife123_():
                    coll63_ = _dafny.Set()
                    compr_63_: int
                    for compr_63_ in (d_1630_v28_).keys.Elements:
                        d_1635_v29_: int = compr_63_
                        if (d_1635_v29_) in (d_1630_v28_):
                            coll63_ = coll63_.union(_dafny.Set([(d_1635_v29_) - (873)]))
                    return _dafny.Set(coll63_)
                nw264_[int(12)] = _dafny.Set({d_1627_v23_, iife122_()
                , iife123_()
                , _dafny.Set({len(d_1631_v30_)})})
                nw264_[int(13)] = d_1629_v26_
                def iife124_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in (d_1627_v23_).Elements:
                        d_1636_v31_: int = compr_64_
                        if (d_1636_v31_) in (d_1627_v23_):
                            coll64_[(d_1636_v31_) * ((d_1622_v19_).cardinality)] = (d_1617_cf6_)[default__.safeIndex(len(_dafny.Map({(d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]: (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]})), len(d_1617_cf6_))]
                    return _dafny.Map(coll64_)
                nw264_[int(14)] = (_dafny.Set({d_1627_v23_, _dafny.Set({len(iife124_()
                ), d_1618_v15_})})) | (d_1629_v26_)
                nw264_[int(15)] = (_dafny.Set({_dafny.Set({(p0)[default__.safeIndex(708, (p0).length(0))]}), d_1627_v23_})) - (d_1629_v26_)
                def iife125_():
                    coll65_ = _dafny.Set()
                    compr_65_: int
                    for compr_65_ in _dafny.IntegerRange(-431, -120):
                        d_1637_v32_: int = compr_65_
                        if ((-431) <= (d_1637_v32_)) and ((d_1637_v32_) < (-120)):
                            coll65_ = coll65_.union(_dafny.Set([(d_1637_v32_) + (987)]))
                    return _dafny.Set(coll65_)
                nw264_[int(16)] = (_dafny.Set({d_1627_v23_, d_1627_v23_})).intersection(_dafny.Set({iife125_()
                , d_1627_v23_}))
                nw264_[int(17)] = _dafny.Set({d_1627_v23_, d_1627_v23_})
                nw264_[int(18)] = d_1629_v26_
                nw264_[int(19)] = d_1629_v26_
                nw264_[int(20)] = d_1629_v26_
                nw264_[int(21)] = d_1629_v26_
                nw264_[int(22)] = d_1629_v26_
                nw264_[int(23)] = default__.fm61(d_1618_v15_, (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))], p1, globalState)
                nw264_[int(24)] = d_1629_v26_
                nw264_[int(25)] = default__.fm61((self).f25, not((d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]), (d_1617_cf6_)[default__.safeIndex(d_1618_v15_, len(d_1617_cf6_))], globalState)
                nw264_[int(26)] = d_1629_v26_
                def iife126_():
                    coll66_ = _dafny.Set()
                    compr_66_: _dafny.Set
                    for compr_66_ in (d_1628_v24_).Elements:
                        d_1638_v33_: _dafny.Set = compr_66_
                        if (d_1638_v33_) in (d_1628_v24_):
                            coll66_ = coll66_.union(_dafny.Set([d_1638_v33_]))
                    return _dafny.Set(coll66_)
                nw264_[int(27)] = iife126_()
                
                nw264_[int(28)] = default__.fm61((p0)[default__.safeIndex(708, (p0).length(0))], p1, p1, globalState)
                d_1632_v34_ = nw264_
                index303_ = default__.safeIndex(851, (d_1632_v34_).length(0))
                (d_1632_v34_)[index303_] = default__.fm61((self).f25, p1, (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))], globalState)
                d_1639_v35_: D9
                d_1639_v35_ = D9_DC27((self).f25, (p0)[default__.safeIndex(708, (p0).length(0))])
                index304_ = default__.safeIndex(851, (d_1632_v34_).length(0))
                index305_ = default__.safeIndex(708, (p0).length(0))
                rhs290_ = ((d_1629_v26_) - (d_1629_v26_) if default__.fm0(d_1618_v15_, globalState) else d_1629_v26_)
                rhs291_ = (False) and ((d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))])
                rhs292_ = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gf"))).set(default__.safeIndex((p0)[default__.safeIndex(708, (p0).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gf")))), d_1591_v0_) if (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))] else (D1_DC6(d_1591_v0_, (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))], (d_1592_v1_).set(default__.safeIndex((self).f25, len(d_1592_v1_)), d_1591_v0_))).cf9))
                rhs293_ = len(d_1631_v30_)
                rhs294_ = D9_DC27(((self).f25) - ((p0)[default__.safeIndex(708, (p0).length(0))]), -817)
                lhs213_ = d_1632_v34_
                lhs214_ = default__.safeIndex(851, (d_1632_v34_).length(0))
                lhs215_ = globalState
                lhs216_ = p0
                lhs217_ = default__.safeIndex(708, (p0).length(0))
                lhs213_[lhs214_] = rhs290_
                lhs215_.f12 = rhs291_
                d_1618_v15_ = rhs292_
                lhs216_[lhs217_] = rhs293_
                d_1639_v35_ = rhs294_
                d_1640_v36_: _dafny.Array
                nw265_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_1640_v36_ = nw265_
                d_1641_v37_: _dafny.Map
                d_1641_v37_ = _dafny.Map({(d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]: d_1640_v36_})
                d_1641_v37_ = (d_1641_v37_).set(p1, d_1640_v36_)
                d_1642_v38_: _dafny.Set
                d_1642_v38_ = _dafny.Set({p1})
                d_1643_v39_: _dafny.Map
                d_1643_v39_ = _dafny.Map({d_1642_v38_: _dafny.Map({p1: (d_1619_v16_)[default__.safeIndex(842, (d_1619_v16_).length(0))]})})
                index306_ = default__.safeIndex(708, (p0).length(0))
                (p0)[index306_] = len((d_1643_v39_) | (d_1643_v39_))
            elif True:
                (globalState).f12 = not (p1) or (not(p1))
                d_1644_v40_: int
                d_1644_v40_ = 747
                d_1644_v40_ = (self).f25
                d_1645_v41_: D14
                d_1645_v41_ = D14_DC38(d_1644_v40_)
                d_1646_v42_: _dafny.Set
                d_1646_v42_ = _dafny.Set({d_1645_v41_})
                d_1647_v43_: _dafny.Map
                d_1647_v43_ = _dafny.Map({d_1646_v42_: d_1592_v1_})
                d_1647_v43_ = d_1647_v43_
                d_1644_v40_ = len(default__.fm20(not (p1) or (p1), globalState))
                d_1648_v44_: _dafny.MultiSet
                d_1648_v44_ = _dafny.MultiSet([p1, False, (d_1592_v1_) < (d_1592_v1_)])
                d_1648_v44_ = default__.fm50(168, (d_1592_v1_ if p1 else d_1592_v1_), p1, 18, globalState)
            if p1:
                d_1649_v45_: _dafny.Set
                d_1649_v45_ = _dafny.Set({p1, p1})
                d_1649_v45_ = ((d_1649_v45_).intersection(d_1649_v45_)) | (d_1649_v45_)
                d_1650_v46_: _dafny.Array
                nw266_ = _dafny.Array(False, 18)
                d_1650_v46_ = nw266_
                d_1651_v47_: _dafny.Seq
                d_1651_v47_ = _dafny.SeqWithoutIsStrInference([d_1650_v46_, d_1650_v46_, d_1650_v46_])
                d_1650_v46_ = (d_1651_v47_)[default__.safeIndex((self).f25, len(d_1651_v47_))]
                d_1652_v48_: int
                d_1652_v48_ = 309
                d_1653_v49_: _dafny.Map
                d_1653_v49_ = _dafny.Map({default__.fm0((self).f25, globalState): default__.fm44(p1, p1, p1, 661, globalState)})
                d_1654_v50_: _dafny.Set
                d_1654_v50_ = _dafny.Set({d_1652_v48_, ((d_1653_v49_)[p1] if (p1) in (d_1653_v49_) else d_1652_v48_)})
                d_1655_v51_: _dafny.Seq
                d_1655_v51_ = _dafny.SeqWithoutIsStrInference([d_1652_v48_])
                def iife127_():
                    coll67_ = _dafny.Set()
                    compr_67_: int
                    for compr_67_ in (d_1655_v51_).Elements:
                        d_1656_v52_: int = compr_67_
                        if (d_1656_v52_) in (d_1655_v51_):
                            coll67_ = coll67_.union(_dafny.Set([(d_1656_v52_) - ((_dafny.MultiSet([_dafny.CodePoint('m')])).cardinality)]))
                    return _dafny.Set(coll67_)
                d_1652_v48_ = len((d_1654_v50_) - (iife127_()
                ))
                d_1657_v53_: D1
                d_1657_v53_ = D1_DC4(d_1617_cf6_)
                d_1658_v54_: _dafny.Map
                d_1658_v54_ = _dafny.Map({p0: d_1657_v53_})
                pat_let_tv31_ = d_1617_cf6_
                def iife128_(_pat_let30_0):
                    def iife129_(d_1659_dt__update__tmp_h1_):
                        def iife130_(_pat_let31_0):
                            def iife131_(d_1660_dt__update_hcf6_h0_):
                                return D1_DC4(d_1660_dt__update_hcf6_h0_)
                            return iife131_(_pat_let31_0)
                        return iife130_(pat_let_tv31_)
                    return iife129_(_pat_let30_0)
                d_1658_v54_ = (d_1658_v54_).set(p0, iife128_(D1_DC4(d_1617_cf6_)))
                d_1661_v55_: _dafny.Map
                d_1661_v55_ = _dafny.Map({len(d_1592_v1_): p1})
                d_1662_v57_: _dafny.Map
                d_1662_v57_ = _dafny.Map({p1: p1})
                d_1663_v58_: _dafny.Seq
                d_1663_v58_ = _dafny.SeqWithoutIsStrInference([d_1661_v55_])
                d_1664_v59_: _dafny.Map
                d_1664_v59_ = d_1661_v55_
                d_1665_v60_: _dafny.Array
                nw267_ = _dafny.Array(None, 26)
                nw267_[int(0)] = (d_1661_v55_) | (d_1661_v55_)
                nw267_[int(1)] = d_1661_v55_
                nw267_[int(2)] = (d_1661_v55_) | (d_1661_v55_)
                def iife132_():
                    coll68_ = _dafny.Map()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(629, -124):
                        d_1666_v56_: int = compr_68_
                        if ((629) <= (d_1666_v56_)) and ((d_1666_v56_) < (-124)):
                            coll68_[default__.safeModuloInt(d_1666_v56_, len(d_1655_v51_))] = p1
                    return _dafny.Map(coll68_)
                nw267_[int(3)] = (iife132_()
                ) | (d_1661_v55_)
                nw267_[int(4)] = d_1661_v55_
                nw267_[int(5)] = (d_1661_v55_) | (d_1661_v55_)
                nw267_[int(6)] = d_1661_v55_
                nw267_[int(7)] = _dafny.Map({default__.fm44(p1, not(p1), p1, d_1652_v48_, globalState): ((d_1662_v57_)[p1] if (p1) in (d_1662_v57_) else p1)})
                nw267_[int(8)] = d_1661_v55_
                nw267_[int(9)] = (d_1661_v55_) | (_dafny.Map({d_1652_v48_: p1}))
                nw267_[int(10)] = default__.fm36(p1, d_1661_v55_, d_1591_v0_, globalState)
                nw267_[int(11)] = d_1661_v55_
                nw267_[int(12)] = _dafny.Map({(self).f25: p1})
                nw267_[int(13)] = d_1661_v55_
                nw267_[int(14)] = d_1661_v55_
                nw267_[int(15)] = d_1661_v55_
                nw267_[int(16)] = d_1661_v55_
                nw267_[int(17)] = d_1661_v55_
                nw267_[int(18)] = _dafny.Map({(D0_DC3((self).f25, default__.fm0(len(d_1654_v50_), globalState))).cf4: p1})
                nw267_[int(19)] = (d_1661_v55_).set(d_1652_v48_, p1)
                nw267_[int(20)] = (d_1661_v55_) | (d_1661_v55_)
                nw267_[int(21)] = d_1661_v55_
                nw267_[int(22)] = (d_1661_v55_) | (d_1661_v55_)
                nw267_[int(23)] = (d_1663_v58_)[default__.safeIndex(d_1652_v48_, len(d_1663_v58_))]
                nw267_[int(24)] = d_1661_v55_
                nw267_[int(25)] = ((d_1664_v59_)) | (d_1661_v55_)
                d_1665_v60_ = nw267_
                def lambda107_(d_1667_v55_):
                    def lambda108_(d_1668_i2_):
                        return d_1667_v55_

                    return lambda108_

                init59_ = lambda107_(d_1661_v55_)
                nw268_ = _dafny.Array(None, 14)
                for i0_59_ in range(nw268_.length(0)):
                    nw268_[i0_59_] = init59_(i0_59_)
                d_1665_v60_ = nw268_
            elif True:
                d_1669_v61_: int
                d_1669_v61_ = 750
                d_1669_v61_ = ((0) - ((self).f25)) + (len(d_1617_cf6_))
                d_1670_v62_: _dafny.Seq
                d_1670_v62_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26])
                d_1671_v63_: _dafny.MultiSet
                d_1671_v63_ = _dafny.MultiSet([d_1591_v0_])
                d_1672_v64_: _dafny.Seq
                d_1672_v64_ = _dafny.SeqWithoutIsStrInference([(d_1671_v63_).cardinality])
                d_1673_v65_: _dafny.Set
                d_1673_v65_ = _dafny.Set({d_1672_v64_})
                d_1674_v66_: C4
                nw269_ = C4()
                nw269_.ctor__((d_1670_v62_)[default__.safeIndex((self).f25, len(d_1670_v62_))], len((d_1673_v65_) - (_dafny.Set({d_1672_v64_}))), (self).f26)
                d_1674_v66_ = nw269_
                d_1675_v67_: _dafny.MultiSet
                d_1675_v67_ = _dafny.MultiSet([p1])
                d_1669_v61_ = default__.safeDivisionInt(102, default__.safeDivisionInt((d_1675_v67_).cardinality, d_1669_v61_))
                d_1676_v68_: _dafny.Array
                def lambda109_(d_1677_v67_):
                    def lambda110_(d_1678_i3_):
                        return (d_1677_v67_) | (d_1677_v67_)

                    return lambda110_

                init60_ = lambda109_(d_1675_v67_)
                nw270_ = _dafny.Array(None, 17)
                for i0_60_ in range(nw270_.length(0)):
                    nw270_[i0_60_] = init60_(i0_60_)
                d_1676_v68_ = nw270_
                index307_ = default__.safeIndex(403, (d_1676_v68_).length(0))
                (d_1676_v68_)[index307_] = default__.fm50(-87, d_1592_v1_, p1, (self).f25, globalState)
                index308_ = default__.safeIndex(403, (d_1676_v68_).length(0))
                (d_1676_v68_)[index308_] = (_dafny.MultiSet(d_1617_cf6_)) | ((d_1675_v67_) - (d_1675_v67_))
                d_1679_v69_: _dafny.MultiSet
                d_1679_v69_ = _dafny.MultiSet([d_1669_v61_, ((d_1676_v68_)[default__.safeIndex(403, (d_1676_v68_).length(0))]).cardinality, d_1669_v61_])
                d_1679_v69_ = (_dafny.MultiSet([(self).f25, (self).f25])) | (d_1679_v69_)
            d_1680_v70_: _dafny.Array
            nw271_ = _dafny.Array(None, 9)
            nw271_[int(0)] = not(p1)
            nw271_[int(1)] = p1
            nw271_[int(2)] = p1
            nw271_[int(3)] = p1
            nw271_[int(4)] = default__.fm0((self).f25, globalState)
            nw271_[int(5)] = False
            nw271_[int(6)] = p1
            nw271_[int(7)] = p1
            nw271_[int(8)] = p1
            d_1680_v70_ = nw271_
            d_1681_v71_: _dafny.Map
            d_1681_v71_ = _dafny.Map({d_1680_v70_: (self).f25})
            d_1681_v71_ = (d_1681_v71_).set(d_1680_v70_, (self).f25)
            (globalState).f12 = p1
        elif True:
            d_1682___mcc_h4_ = source23_.cf10
            d_1683_cf10_ = d_1682___mcc_h4_
            d_1684_v72_: D7
            d_1684_v72_ = D7_DC21(p0)
            d_1685_v73_: int
            d_1685_v73_ = 299
            d_1686_v74_: _dafny.Seq
            d_1686_v74_ = _dafny.SeqWithoutIsStrInference([d_1592_v1_])
            d_1687_v75_: D9
            d_1687_v75_ = D9_DC27((self).f25, d_1685_v73_)
            d_1688_v76_: _dafny.Seq
            d_1688_v76_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1])), len(d_1592_v1_)])
            rhs295_ = d_1684_v72_
            rhs296_ = p1
            rhs297_ = (((d_1686_v74_) + (default__.fm58(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), d_1591_v0_, d_1591_v0_]), (self).f25, p1, default__.fm44(p1, default__.fm0((self).f25, globalState), p1, len(_dafny.SeqWithoutIsStrInference([(self).f25, len(d_1592_v1_), (self).f25, d_1685_v73_])), globalState), globalState))).set(default__.safeIndex(len(d_1592_v1_), len((d_1686_v74_) + (default__.fm58(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), d_1591_v0_, d_1591_v0_]), (self).f25, p1, default__.fm44(p1, default__.fm0((self).f25, globalState), p1, len(_dafny.SeqWithoutIsStrInference([(self).f25, len(d_1592_v1_), (self).f25, d_1685_v73_])), globalState), globalState)))), d_1592_v1_)) <= ((_dafny.SeqWithoutIsStrInference([d_1592_v1_, d_1592_v1_, d_1592_v1_])) + (d_1686_v74_))
            rhs298_ = default__.safeModuloInt(len(_dafny.Set({(self).f25, (self).f25, (self).f25, (d_1687_v75_).cf39, d_1685_v73_})), (self).f25)
            rhs299_ = ((d_1592_v1_) + (d_1592_v1_)) + ((_dafny.SeqWithoutIsStrInference([d_1591_v0_, d_1591_v0_, d_1591_v0_, _dafny.CodePoint('j'), d_1591_v0_])) + ((d_1686_v74_)[default__.safeIndex(len(d_1688_v76_), len(d_1686_v74_))]))
            lhs218_ = globalState
            lhs219_ = globalState
            d_1684_v72_ = rhs295_
            lhs218_.f12 = rhs296_
            r1 = rhs297_
            d_1685_v73_ = rhs298_
            lhs219_.f10 = rhs299_
            d_1689_v77_: _dafny.Map
            d_1689_v77_ = _dafny.Map({d_1591_v0_: (self).f25})
            d_1690_v78_: _dafny.Seq
            d_1690_v78_ = _dafny.SeqWithoutIsStrInference([p1, p1, False])
            d_1691_v79_: _dafny.Map
            d_1691_v79_ = _dafny.Map({(d_1690_v78_)[default__.safeIndex(56, len(d_1690_v78_))]: (self).f25})
            d_1689_v77_ = _dafny.Map({d_1591_v0_: default__.safeModuloInt(len(d_1691_v79_), (self).f25)})
            d_1692_v80_: D5
            d_1692_v80_ = D5_DC17((self).f25, (d_1690_v78_)[default__.safeIndex(d_1685_v73_, len(d_1690_v78_))])
            d_1693_v81_: _dafny.Set
            d_1693_v81_ = _dafny.Set({p1})
            pat_let_tv32_ = d_1693_v81_
            pat_let_tv33_ = d_1693_v81_
            def iife134_(_pat_let33_0):
                def iife135_(d_1694_dt__update__tmp_h2_):
                    def iife136_(_pat_let34_0):
                        def iife137_(d_1695_dt__update_hcf26_h0_):
                            return D5_DC17(d_1695_dt__update_hcf26_h0_, (d_1694_dt__update__tmp_h2_).cf27)
                        return iife137_(_pat_let34_0)
                    return iife136_((self).f25)
                return iife135_(_pat_let33_0)
            def iife133_(_pat_let32_0):
                def iife138_(d_1696_dt__update__tmp_h3_):
                    def iife139_(_pat_let35_0):
                        def iife140_(d_1697_dt__update_hcf27_h0_):
                            return D5_DC17((d_1696_dt__update__tmp_h3_).cf26, d_1697_dt__update_hcf27_h0_)
                        return iife140_(_pat_let35_0)
                    return iife139_((pat_let_tv32_).isdisjoint(pat_let_tv33_))
                return iife138_(_pat_let32_0)
            source24_ = iife133_(iife134_(d_1692_v80_))
            if source24_.is_DC16:
                d_1698___mcc_h5_ = source24_.cf25
                d_1699_cf25_ = d_1698___mcc_h5_
                d_1700_v83_: _dafny.Map
                def iife141_():
                    coll69_ = _dafny.Map()
                    compr_69_: str
                    for compr_69_ in (d_1592_v1_).Elements:
                        d_1701_v82_: str = compr_69_
                        if (d_1701_v82_) in (d_1592_v1_):
                            coll69_[d_1701_v82_] = p1
                    return _dafny.Map(coll69_)
                d_1700_v83_ = _dafny.Map({not(d_1699_cf25_): D10_DC28(iife141_()
)})
                d_1702_v84_: C0
                nw272_ = C0()
                nw272_.ctor__()
                d_1702_v84_ = nw272_
                d_1703_v85_: _dafny.Map
                d_1703_v85_ = _dafny.Map({d_1591_v0_: d_1699_cf25_})
                d_1704_v86_: D10
                d_1704_v86_ = D10_DC28(d_1703_v85_)
                d_1705_v87_: _dafny.Map
                d_1705_v87_ = _dafny.Map({len(_dafny.Map({d_1684_v72_: d_1702_v84_})): _dafny.Map({p1: d_1704_v86_})})
                d_1706_v88_: _dafny.Array
                nw273_ = _dafny.Array(None, 14)
                nw273_[int(0)] = d_1700_v83_
                nw273_[int(1)] = d_1700_v83_
                nw273_[int(2)] = d_1700_v83_
                nw273_[int(3)] = ((d_1705_v87_)[d_1685_v73_] if (d_1685_v73_) in (d_1705_v87_) else d_1700_v83_)
                nw273_[int(4)] = d_1700_v83_
                nw273_[int(5)] = d_1700_v83_
                nw273_[int(6)] = _dafny.Map({d_1699_cf25_: d_1704_v86_})
                nw273_[int(7)] = _dafny.Map({p1: d_1704_v86_})
                nw273_[int(8)] = d_1700_v83_
                nw273_[int(9)] = _dafny.Map({p1: d_1704_v86_})
                nw273_[int(10)] = (d_1700_v83_) | (_dafny.Map({d_1699_cf25_: d_1704_v86_}))
                nw273_[int(11)] = (d_1700_v83_) | (d_1700_v83_)
                nw273_[int(12)] = default__.fm62(globalState)
                nw273_[int(13)] = _dafny.Map({p1: d_1704_v86_})
                d_1706_v88_ = nw273_
                index309_ = default__.safeIndex(82, (d_1706_v88_).length(0))
                (d_1706_v88_)[index309_] = d_1700_v83_
                index310_ = default__.safeIndex(82, (d_1706_v88_).length(0))
                (d_1706_v88_)[index310_] = ((d_1700_v83_) | (d_1700_v83_)) | (((d_1700_v83_).set(p1, d_1704_v86_)) | (d_1700_v83_))
                d_1685_v73_ = (d_1702_v84_).fm26(((d_1691_v79_)[p1] if (p1) in (d_1691_v79_) else d_1685_v73_), (d_1699_cf25_ if False else d_1699_cf25_), globalState)
                d_1707_v89_: _dafny.Map
                d_1707_v89_ = _dafny.Map({((self).f25) - (574): d_1699_cf25_})
                d_1707_v89_ = (d_1707_v89_).set(((self).f25) * ((self).f25), d_1699_cf25_)
                (globalState).f6 = d_1688_v76_
            elif source24_.is_DC17:
                d_1708___mcc_h6_ = source24_.cf26
                d_1709___mcc_h7_ = source24_.cf27
                d_1710_cf27_ = d_1709___mcc_h7_
                d_1711_cf26_ = d_1708___mcc_h6_
                d_1712_v90_: _dafny.Array
                nw274_ = _dafny.Array(False, 20)
                d_1712_v90_ = nw274_
                def lambda111_(d_1713_v78_):
                    def lambda112_(d_1714_i4_):
                        return (d_1713_v78_) != (d_1713_v78_)

                    return lambda112_

                init61_ = lambda111_(d_1690_v78_)
                nw275_ = _dafny.Array(None, 1)
                for i0_61_ in range(nw275_.length(0)):
                    nw275_[i0_61_] = init61_(i0_61_)
                d_1712_v90_ = nw275_
                d_1715_v91_: D0
                d_1715_v91_ = D0_DC3((self).f25, d_1710_cf27_)
                d_1691_v79_ = (d_1691_v79_).set((622) > ((d_1715_v91_).cf4), d_1711_cf26_)
                d_1716_v92_: _dafny.Map
                d_1716_v92_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): p1})
                d_1716_v92_ = (d_1716_v92_).set(_dafny.SeqWithoutIsStrInference([p1, p1, d_1710_cf27_]), ((self).f25) > ((self).f25))
                d_1717_v93_: _dafny.Array
                def lambda113_(d_1718_v0_):
                    def lambda114_(d_1719_i5_):
                        return d_1718_v0_

                    return lambda114_

                init62_ = lambda113_(d_1591_v0_)
                nw276_ = _dafny.Array(None, 26)
                for i0_62_ in range(nw276_.length(0)):
                    nw276_[i0_62_] = init62_(i0_62_)
                d_1717_v93_ = nw276_
                index311_ = default__.safeIndex(102, (d_1717_v93_).length(0))
                (d_1717_v93_)[index311_] = default__.fm39(d_1710_cf27_, globalState)
                index312_ = default__.safeIndex(692, (d_1717_v93_).length(0))
                (d_1717_v93_)[index312_] = d_1591_v0_
                index313_ = default__.safeIndex(102, (d_1717_v93_).length(0))
                index314_ = default__.safeIndex(692, (d_1717_v93_).length(0))
                rhs300_ = (d_1712_v90_ if p1 else d_1712_v90_)
                rhs301_ = d_1591_v0_
                rhs302_ = d_1591_v0_
                lhs220_ = d_1717_v93_
                lhs221_ = default__.safeIndex(102, (d_1717_v93_).length(0))
                lhs222_ = d_1717_v93_
                lhs223_ = default__.safeIndex(692, (d_1717_v93_).length(0))
                d_1712_v90_ = rhs300_
                lhs220_[lhs221_] = rhs301_
                lhs222_[lhs223_] = rhs302_
            elif True:
                d_1720___mcc_h8_ = source24_.cf24
                d_1721_cf24_ = d_1720___mcc_h8_
                (globalState).f12 = p1
                d_1722_v94_: _dafny.Map
                d_1722_v94_ = _dafny.Map({default__.fm19(p1, globalState): False})
                d_1723_v95_: D0
                d_1723_v95_ = D0_DC2(p1)
                d_1724_v96_: _dafny.Map
                d_1724_v96_ = _dafny.Map({d_1722_v94_: d_1723_v95_})
                d_1724_v96_ = (d_1724_v96_).set(d_1722_v94_, d_1723_v95_)
                d_1725_v97_: _dafny.Array
                def lambda115_(d_1726_i6_):
                    return False

                init63_ = lambda115_
                nw277_ = _dafny.Array(None, 10)
                for i0_63_ in range(nw277_.length(0)):
                    nw277_[i0_63_] = init63_(i0_63_)
                d_1725_v97_ = nw277_
                index315_ = default__.safeIndex(741, (d_1725_v97_).length(0))
                (d_1725_v97_)[index315_] = p1
                index316_ = default__.safeIndex(741, (d_1725_v97_).length(0))
                (d_1725_v97_)[index316_] = (d_1592_v1_) <= ((d_1592_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xum"))))
                d_1727_v98_: _dafny.Map
                d_1727_v98_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1591_v0_ for d_1728_i7_ in range(default__.abs(919))]): d_1725_v97_})
                d_1729_v99_: _dafny.Map
                d_1729_v99_ = _dafny.Map({((d_1727_v98_)[_dafny.SeqWithoutIsStrInference([d_1591_v0_ for d_1730_i8_ in range(default__.abs(814))])] if (_dafny.SeqWithoutIsStrInference([d_1591_v0_ for d_1731_i8_ in range(default__.abs(814))])) in (d_1727_v98_) else d_1725_v97_): (d_1725_v97_)[default__.safeIndex(741, (d_1725_v97_).length(0))]})
                d_1732_v100_: D8
                d_1732_v100_ = D8_DC23(d_1729_v99_)
                d_1733_v101_: D8
                d_1733_v101_ = D8_DC25(d_1732_v100_)
                d_1734_v102_: _dafny.Set
                d_1734_v102_ = _dafny.Set({d_1733_v101_})
                d_1685_v73_ = default__.safeModuloInt(569, len(((d_1592_v1_).set(default__.safeIndex(len(d_1734_v102_), len(d_1592_v1_)), d_1591_v0_)) + (d_1592_v1_)))
            rhs303_ = p1
            rhs304_ = d_1591_v0_
            rhs305_ = d_1685_v73_
            rhs306_ = default__.fm39(p1, globalState)
            rhs307_ = d_1592_v1_
            lhs224_ = globalState
            lhs225_ = globalState
            r1 = rhs303_
            lhs224_.f23 = rhs304_
            d_1685_v73_ = rhs305_
            lhs225_.f23 = rhs306_
            d_1592_v1_ = rhs307_
        d_1735_v103_: _dafny.Array
        nw278_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
        d_1735_v103_ = nw278_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1735_v103_).length(0)):
            d_1736_i9_: int = guard_loop_8_
            if (True) and (((0) <= (d_1736_i9_)) and ((d_1736_i9_) < ((d_1735_v103_).length(0)))):
                (d_1735_v103_)[(d_1736_i9_)] = _dafny.SeqWithoutIsStrInference([d_1591_v0_ for d_1737_i10_ in range(default__.abs(404))])
        d_1738_v104_: _dafny.Array
        nw279_ = _dafny.Array(_dafny.Array(None, 0), 11)
        d_1738_v104_ = nw279_
        d_1739_v105_: _dafny.Map
        d_1739_v105_ = _dafny.Map({p1: (self).f25})
        d_1740_v106_: _dafny.Map
        d_1740_v106_ = _dafny.Map({(self).f25: d_1739_v105_})
        d_1741_v107_: _dafny.Map
        d_1741_v107_ = _dafny.Map({((self).f25) in (d_1740_v106_): d_1738_v104_})
        d_1742_v108_: _dafny.Map
        d_1742_v108_ = _dafny.Map({p1: p1})
        d_1738_v104_ = ((d_1741_v107_)[((d_1742_v108_)[p1] if (p1) in (d_1742_v108_) else p1)] if (((d_1742_v108_)[p1] if (p1) in (d_1742_v108_) else p1)) in (d_1741_v107_) else d_1738_v104_)
        d_1743_v109_: int
        d_1743_v109_ = 573
        d_1743_v109_ = (d_1743_v109_) + (d_1743_v109_)
        hi8_ = default__.safeModuloInt(d_1743_v109_, (self).f25)
        for d_1744_i11_ in range((self).f25, hi8_):
            d_1745_v110_: _dafny.MultiSet
            d_1745_v110_ = _dafny.MultiSet([p1])
            d_1743_v109_ = (((d_1745_v110_)[p1] if (p1) in (d_1745_v110_) else (self).f25) if (d_1744_i11_) >= ((self).f25) else d_1744_i11_)
            index317_ = default__.safeIndex(396, (p0).length(0))
            (p0)[index317_] = (d_1744_i11_ if p1 else d_1744_i11_)
            d_1746_v111_: _dafny.Seq
            d_1746_v111_ = _dafny.SeqWithoutIsStrInference([d_1745_v110_, d_1745_v110_, d_1745_v110_])
            d_1747_v112_: T0
            nw280_ = C4()
            nw280_.ctor__((self).f26, d_1743_v109_, d_1735_v103_)
            d_1747_v112_ = nw280_
            d_1748_v113_: _dafny.Map
            d_1748_v113_ = _dafny.Map({d_1747_v112_: p1})
            d_1749_v114_: _dafny.Seq
            d_1749_v114_ = _dafny.SeqWithoutIsStrInference([d_1748_v113_])
            d_1750_v115_: _dafny.Map
            d_1750_v115_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtllwx"))): (d_1747_v112_).f25})
            index318_ = default__.safeIndex(396, (p0).length(0))
            rhs308_ = (d_1749_v114_) <= (d_1749_v114_)
            rhs309_ = default__.safeDivisionInt(d_1743_v109_, default__.safeDivisionInt((d_1747_v112_).f25, (0) - (len(d_1750_v115_))))
            rhs310_ = (0) - ((self).f25)
            rhs311_ = ((d_1747_v112_).f25) - (d_1744_i11_)
            rhs312_ = d_1746_v111_
            lhs226_ = globalState
            lhs227_ = p0
            lhs228_ = default__.safeIndex(396, (p0).length(0))
            lhs226_.f12 = rhs308_
            d_1743_v109_ = rhs309_
            lhs227_[lhs228_] = rhs310_
            d_1743_v109_ = rhs311_
            d_1746_v111_ = rhs312_
            d_1751_v116_: _dafny.Array
            nw281_ = _dafny.Array(False, 22)
            d_1751_v116_ = nw281_
            index319_ = default__.safeIndex(590, (d_1751_v116_).length(0))
            (d_1751_v116_)[index319_] = default__.fm0((0) - ((d_1747_v112_).f25), globalState)
            index320_ = default__.safeIndex(590, (d_1751_v116_).length(0))
            (d_1751_v116_)[index320_] = True
            d_1752_v117_: _dafny.Map
            d_1752_v117_ = _dafny.Map({(d_1751_v116_)[default__.safeIndex(590, (d_1751_v116_).length(0))]: d_1592_v1_})
            d_1753_v118_: _dafny.Seq
            d_1753_v118_ = _dafny.SeqWithoutIsStrInference([((d_1745_v110_)[p1] if (p1) in (d_1745_v110_) else (d_1747_v112_).f25)])
            d_1754_v119_: _dafny.MultiSet
            d_1754_v119_ = _dafny.MultiSet([len(d_1753_v118_), (self).f25])
            d_1755_v120_: _dafny.Seq
            d_1755_v120_ = _dafny.SeqWithoutIsStrInference([default__.fm18(globalState), default__.fm31(((d_1752_v117_)[p1] if (p1) in (d_1752_v117_) else d_1592_v1_), globalState), ((d_1754_v119_)[d_1744_i11_] if (d_1744_i11_) in (d_1754_v119_) else d_1744_i11_)])
            rhs313_ = ((0) - ((d_1745_v110_).cardinality)) * (d_1744_i11_)
            rhs314_ = (d_1753_v118_)[default__.safeIndex(((self).f25) - (len(d_1739_v105_)), len(d_1753_v118_))]
            rhs315_ = d_1592_v1_
            rhs316_ = not((p1 if (d_1745_v110_).issubset(_dafny.MultiSet([p1, p1])) else (d_1751_v116_)[default__.safeIndex(590, (d_1751_v116_).length(0))]))
            rhs317_ = _dafny.SeqWithoutIsStrInference([len((d_1592_v1_) + (d_1592_v1_))])
            lhs229_ = globalState
            lhs230_ = globalState
            d_1743_v109_ = rhs313_
            d_1743_v109_ = rhs314_
            lhs229_.f10 = rhs315_
            r1 = rhs316_
            lhs230_.f6 = rhs317_
        d_1756_i12_: int
        d_1756_i12_ = 0
        with _dafny.label("12"):
            while p1:
                with _dafny.c_label("12"):
                    if (d_1756_i12_) >= (100):
                        raise _dafny.Break("12")
                    d_1756_i12_ = (d_1756_i12_) + (1)
                    d_1742_v108_ = (d_1742_v108_).set(p1, (d_1743_v109_) <= ((default__.fm63(p1, (self).f25, p1, default__.fm0((self).f25, globalState), globalState)).cardinality))
                    if p1:
                        index321_ = default__.safeIndex(186, ((self).f26).length(0))
                        ((self).f26)[index321_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lubqn"))
                        index322_ = default__.safeIndex(186, ((self).f26).length(0))
                        ((self).f26)[index322_] = d_1592_v1_
                        d_1757_v121_: _dafny.MultiSet
                        d_1757_v121_ = _dafny.MultiSet([p1, p1, p1, p1])
                        d_1758_v122_: _dafny.MultiSet
                        d_1758_v122_ = _dafny.MultiSet([d_1757_v121_])
                        d_1758_v122_ = d_1758_v122_
                        d_1743_v109_ = d_1743_v109_
                        d_1759_v123_: _dafny.Seq
                        d_1759_v123_ = _dafny.SeqWithoutIsStrInference([(D13_DC36(d_1742_v108_, p1, (self).f25)).cf56, (self).f25])
                        d_1760_v124_: _dafny.Map
                        d_1760_v124_ = _dafny.Map({(d_1759_v123_).set(default__.safeIndex((self).f25, len(d_1759_v123_)), d_1743_v109_): (0) - ((self).f25)})
                        d_1761_v125_: _dafny.Seq
                        d_1761_v125_ = _dafny.SeqWithoutIsStrInference([d_1759_v123_])
                        d_1760_v124_ = (d_1760_v124_).set((_dafny.SeqWithoutIsStrInference([d_1743_v109_])) + ((d_1761_v125_)[default__.safeIndex((self).f25, len(d_1761_v125_))]), (((d_1739_v105_)[p1] if (p1) in (d_1739_v105_) else -199)) + (d_1743_v109_))
                        d_1762_v126_: _dafny.MultiSet
                        d_1762_v126_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1591_v0_ for d_1763_i13_ in range(default__.abs(784))]))])
                        d_1764_v127_: _dafny.Map
                        d_1764_v127_ = _dafny.Map({(d_1762_v126_).cardinality: p1})
                        d_1765_v128_: D13
                        d_1765_v128_ = D13_DC36(d_1742_v108_, p1, len(d_1764_v127_))
                        d_1766_v129_: _dafny.Array
                        nw282_ = _dafny.Array(None, 24)
                        nw282_[int(0)] = True
                        nw282_[int(1)] = p1
                        nw282_[int(2)] = p1
                        nw282_[int(3)] = p1
                        nw282_[int(4)] = True
                        nw282_[int(5)] = p1
                        nw282_[int(6)] = p1
                        nw282_[int(7)] = p1
                        nw282_[int(8)] = p1
                        nw282_[int(9)] = True
                        nw282_[int(10)] = p1
                        nw282_[int(11)] = p1
                        nw282_[int(12)] = p1
                        nw282_[int(13)] = False
                        nw282_[int(14)] = not(p1)
                        nw282_[int(15)] = p1
                        nw282_[int(16)] = p1
                        nw282_[int(17)] = p1
                        nw282_[int(18)] = p1
                        nw282_[int(19)] = p1
                        nw282_[int(20)] = p1
                        nw282_[int(21)] = (d_1765_v128_).cf55
                        nw282_[int(22)] = p1
                        nw282_[int(23)] = p1
                        d_1766_v129_ = nw282_
                        d_1767_v130_: _dafny.MultiSet
                        d_1767_v130_ = _dafny.MultiSet([d_1766_v129_])
                        (globalState).f12 = (d_1767_v130_).issubset(d_1767_v130_)
                    elif True:
                        d_1768_v131_: D0
                        d_1768_v131_ = D0_DC1(False, p1)
                        pat_let_tv34_ = globalState
                        d_1769_v132_: C3
                        nw283_ = C3()
                        def iife142_(_pat_let36_0):
                            def iife143_(d_1770_dt__update__tmp_h4_):
                                def iife144_(_pat_let37_0):
                                    def iife145_(d_1771_dt__update_hcf2_h0_):
                                        return D0_DC1((d_1770_dt__update__tmp_h4_).cf1, d_1771_dt__update_hcf2_h0_)
                                    return iife145_(_pat_let37_0)
                                return iife144_(default__.fm0((self).f25, pat_let_tv34_))
                            return iife143_(_pat_let36_0)
                        nw283_.ctor__(398, (p1 if (iife142_(d_1768_v131_)).cf2 else p1), ((self).f25 if p1 else (self).f25), (self).f26)
                        d_1769_v132_ = nw283_
                        d_1743_v109_ = (len(_dafny.Map({default__.fm0(d_1743_v109_, globalState): p1}))) * (932)
                        d_1772_v133_: _dafny.Array
                        nw284_ = _dafny.Array(False, 13)
                        d_1772_v133_ = nw284_
                        d_1773_v134_: _dafny.Set
                        d_1773_v134_ = _dafny.Set({not((d_1769_v132_).f38)})
                        d_1774_v135_: _dafny.Seq
                        d_1774_v135_ = _dafny.SeqWithoutIsStrInference([d_1773_v134_])
                        d_1775_v136_: _dafny.Map
                        d_1775_v136_ = _dafny.Map({d_1772_v133_: (d_1774_v135_)[default__.safeIndex((self).f25, len(d_1774_v135_))]})
                        d_1776_v137_: _dafny.Seq
                        d_1776_v137_ = _dafny.SeqWithoutIsStrInference([d_1743_v109_])
                        d_1775_v136_ = (_dafny.Map({d_1772_v133_: d_1773_v134_}) if (d_1776_v137_) <= (d_1776_v137_) else d_1775_v136_)
                        r1 = False
                        (globalState).f12 = False
                    d_1743_v109_ = (0) - ((self).f25)
                    d_1777_v138_: D5
                    d_1777_v138_ = D5_DC16(p1)
                    d_1778_v139_: _dafny.Set
                    d_1778_v139_ = _dafny.Set({d_1777_v138_})
                    d_1779_v140_: _dafny.Map
                    d_1779_v140_ = _dafny.Map({(d_1778_v139_) - (d_1778_v139_): (self).f25})
                    d_1780_v141_: _dafny.Seq
                    d_1780_v141_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_1779_v140_ = (d_1779_v140_).set((d_1778_v139_) | (d_1778_v139_), len(d_1780_v141_))
                    pass
            pass
        d_1781_v142_: _dafny.Array
        def lambda116_(d_1782_p1_):
            def lambda117_(d_1783_i14_):
                return d_1782_p1_

            return lambda117_

        init64_ = lambda116_(p1)
        nw285_ = _dafny.Array(None, 6)
        for i0_64_ in range(nw285_.length(0)):
            nw285_[i0_64_] = init64_(i0_64_)
        d_1781_v142_ = nw285_
        d_1784_v143_: D3
        d_1784_v143_ = D3_DC10(p1, d_1781_v142_, p1, -294, (self).f25)
        d_1785_v144_: _dafny.Seq
        d_1785_v144_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
        pat_let_tv35_ = d_1785_v144_
        pat_let_tv36_ = d_1785_v144_
        pat_let_tv37_ = d_1781_v142_
        def iife146_(_pat_let38_0):
            def iife147_(d_1786_dt__update__tmp_h5_):
                def iife148_(_pat_let39_0):
                    def iife149_(d_1787_dt__update_hcf15_h0_):
                        def iife150_(_pat_let40_0):
                            def iife151_(d_1788_dt__update_hcf14_h0_):
                                return D3_DC10((d_1786_dt__update__tmp_h5_).cf13, d_1788_dt__update_hcf14_h0_, d_1787_dt__update_hcf15_h0_, (d_1786_dt__update__tmp_h5_).cf16, (d_1786_dt__update__tmp_h5_).cf17)
                            return iife151_(_pat_let40_0)
                        return iife150_(pat_let_tv37_)
                    return iife149_(_pat_let39_0)
                return iife148_((pat_let_tv35_)[default__.safeIndex((self).f25, len(pat_let_tv36_))])
            return iife147_(_pat_let38_0)
        r0 = iife146_(d_1784_v143_)
        r1 = p1
        return r0, r1

    def m12(self, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: bool = False
        d_1789_v0_: _dafny.Array
        def lambda118_(d_1790_i0_):
            return ((self).f25) == ((self).f25)

        init65_ = lambda118_
        nw286_ = _dafny.Array(None, 26)
        for i0_65_ in range(nw286_.length(0)):
            nw286_[i0_65_] = init65_(i0_65_)
        d_1789_v0_ = nw286_
        d_1791_v1_: bool
        d_1791_v1_ = True
        index323_ = default__.safeIndex(562, (d_1789_v0_).length(0))
        (d_1789_v0_)[index323_] = d_1791_v1_
        index324_ = default__.safeIndex(562, (d_1789_v0_).length(0))
        (d_1789_v0_)[index324_] = not(not (True) or (d_1791_v1_))
        d_1792_i1_: int
        d_1792_i1_ = 0
        with _dafny.label("13"):
            while ((d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]) == ((d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]):
                with _dafny.c_label("13"):
                    if (d_1792_i1_) >= (100):
                        raise _dafny.Break("13")
                    d_1792_i1_ = (d_1792_i1_) + (1)
                    d_1793_v2_: _dafny.Array
                    def lambda119_(d_1794_i2_):
                        return default__.safeDivisionInt(d_1794_i2_, (self).f25)

                    init66_ = lambda119_
                    nw287_ = _dafny.Array(None, 2)
                    for i0_66_ in range(nw287_.length(0)):
                        nw287_[i0_66_] = init66_(i0_66_)
                    d_1793_v2_ = nw287_
                    index325_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    (d_1793_v2_)[index325_] = (self).f25
                    index326_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    (d_1793_v2_)[index326_] = (567) * ((self).f25)
                    d_1795_v3_: _dafny.Map
                    d_1795_v3_ = _dafny.Map({(d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]: (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]})
                    d_1796_v4_: _dafny.Seq
                    d_1796_v4_ = _dafny.SeqWithoutIsStrInference([not(((d_1795_v3_)[d_1791_v1_] if (d_1791_v1_) in (d_1795_v3_) else not(d_1791_v1_)))])
                    index327_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    index328_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    index329_ = default__.safeIndex(562, (d_1789_v0_).length(0))
                    index330_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    rhs318_ = (d_1793_v2_)[default__.safeIndex(916, (d_1793_v2_).length(0))]
                    rhs319_ = (self).f25
                    rhs320_ = d_1789_v0_
                    rhs321_ = ((d_1796_v4_)[default__.safeIndex((self).f25, len(d_1796_v4_))]) and (((d_1793_v2_)[default__.safeIndex(916, (d_1793_v2_).length(0))]) > ((d_1793_v2_)[default__.safeIndex(916, (d_1793_v2_).length(0))]))
                    rhs322_ = (0) - ((((0) - (-32)) * ((d_1793_v2_)[default__.safeIndex(916, (d_1793_v2_).length(0))])) * ((self).f25))
                    lhs231_ = d_1793_v2_
                    lhs232_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    lhs233_ = d_1793_v2_
                    lhs234_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    lhs235_ = d_1789_v0_
                    lhs236_ = default__.safeIndex(562, (d_1789_v0_).length(0))
                    lhs237_ = d_1793_v2_
                    lhs238_ = default__.safeIndex(916, (d_1793_v2_).length(0))
                    lhs231_[lhs232_] = rhs318_
                    lhs233_[lhs234_] = rhs319_
                    d_1789_v0_ = rhs320_
                    lhs235_[lhs236_] = rhs321_
                    lhs237_[lhs238_] = rhs322_
                    d_1797_v5_: str
                    d_1797_v5_ = _dafny.CodePoint('o')
                    d_1798_v6_: _dafny.Map
                    d_1798_v6_ = _dafny.Map({not(d_1791_v1_): d_1797_v5_})
                    (globalState).f23 = ((d_1798_v6_)[d_1791_v1_] if (d_1791_v1_) in (d_1798_v6_) else d_1797_v5_)
                    d_1799_v7_: _dafny.Seq
                    d_1799_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbnx"))
                    d_1800_v8_: _dafny.Map
                    d_1800_v8_ = _dafny.Map({d_1799_v7_: (self).f25})
                    rhs323_ = (d_1800_v8_) | (d_1800_v8_)
                    rhs324_ = (d_1799_v7_) <= (d_1799_v7_)
                    rhs325_ = (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]
                    lhs239_ = globalState
                    lhs240_ = globalState
                    d_1800_v8_ = rhs323_
                    lhs239_.f12 = rhs324_
                    lhs240_.f12 = rhs325_
                    pass
            pass
        d_1801_i3_: int
        d_1801_i3_ = 0
        with _dafny.label("14"):
            while (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]:
                with _dafny.c_label("14"):
                    if (d_1801_i3_) >= (100):
                        raise _dafny.Break("14")
                    d_1801_i3_ = (d_1801_i3_) + (1)
                    d_1802_v9_: int
                    d_1802_v9_ = 782
                    d_1802_v9_ = d_1802_v9_
                    d_1803_v10_: _dafny.Array
                    def lambda120_(d_1804_v9_):
                        def lambda121_(d_1805_i4_):
                            return default__.safeDivisionInt(d_1805_i4_, len(_dafny.SeqWithoutIsStrInference([(0) - (d_1804_v9_), d_1804_v9_])))

                        return lambda121_

                    init67_ = lambda120_(d_1802_v9_)
                    nw288_ = _dafny.Array(None, 28)
                    for i0_67_ in range(nw288_.length(0)):
                        nw288_[i0_67_] = init67_(i0_67_)
                    d_1803_v10_ = nw288_
                    index331_ = default__.safeIndex(414, (d_1803_v10_).length(0))
                    (d_1803_v10_)[index331_] = (self).f25
                    d_1806_v11_: _dafny.Map
                    d_1806_v11_ = _dafny.Map({((0) - ((self).f25)) * (default__.fm31(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdedfs")), globalState)): (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]})
                    index332_ = default__.safeIndex(414, (d_1803_v10_).length(0))
                    (d_1803_v10_)[index332_] = (0) - (len(d_1806_v11_))
                    d_1807_v12_: C0
                    nw289_ = C0()
                    nw289_.ctor__()
                    d_1807_v12_ = nw289_
                    d_1808_v13_: C1
                    nw290_ = C1()
                    nw290_.ctor__(((self).f25) + (len(d_1806_v11_)), (self).f26)
                    d_1808_v13_ = nw290_
                    pass
            pass
        d_1809_v14_: _dafny.Seq
        d_1809_v14_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_1810_v15_: _dafny.MultiSet
        d_1810_v15_ = _dafny.MultiSet([d_1791_v1_, default__.fm0(default__.fm31(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1811_i6_ in range(default__.abs(117))]), globalState), globalState), (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))], False, (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]])
        d_1812_v16_: _dafny.Array
        nw291_ = _dafny.Array(None, 8)
        nw291_[int(0)] = (self).f25
        nw291_[int(1)] = (len(d_1809_v14_)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edjrpfsui"))))
        nw291_[int(2)] = (self).f25
        nw291_[int(3)] = (len(d_1809_v14_)) * (len((d_1809_v14_).set(default__.safeIndex((self).f25, len(d_1809_v14_)), (self).f25)))
        nw291_[int(4)] = (self).f25
        nw291_[int(5)] = (self).f25
        nw291_[int(6)] = ((d_1810_v15_)[False] if (False) in (d_1810_v15_) else (self).f25)
        nw291_[int(7)] = (self).f25
        d_1812_v16_ = nw291_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_1812_v16_).length(0)):
            d_1813_i5_: int = guard_loop_9_
            if (True) and (((0) <= (d_1813_i5_)) and ((d_1813_i5_) < ((d_1812_v16_).length(0)))):
                (d_1812_v16_)[(d_1813_i5_)] = default__.safeDivisionInt(d_1813_i5_, default__.safeModuloInt((self).f25, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1791_v1_: (self).f25}) for d_1814_i7_ in range(default__.abs(475))]))))
        d_1815_v17_: str
        d_1815_v17_ = _dafny.CodePoint('d')
        d_1816_v18_: _dafny.Map
        d_1816_v18_ = _dafny.Map({d_1815_v17_: (self).f25})
        d_1817_v19_: _dafny.Map
        d_1817_v19_ = _dafny.Map({d_1816_v18_: d_1789_v0_})
        d_1817_v19_ = ((d_1817_v19_) | (_dafny.Map({d_1816_v18_: d_1789_v0_}))) | (d_1817_v19_)
        d_1818_v20_: _dafny.Map
        d_1818_v20_ = _dafny.Map({(d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]: d_1789_v0_})
        d_1819_v21_: D3
        d_1820_v22_: bool
        out31_: D3
        out32_: bool
        out31_, out32_ = (self).m11(d_1812_v16_, d_1791_v1_, d_1818_v20_, globalState)
        d_1819_v21_ = out31_
        d_1820_v22_ = out32_
        r0 = _dafny.SeqWithoutIsStrInference([(D1_DC6(d_1815_v17_, (d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wubujscf")))).cf7 for d_1821_i8_ in range(default__.abs(690))])
        d_1822_v23_: _dafny.Seq
        d_1822_v23_ = _dafny.SeqWithoutIsStrInference([(d_1789_v0_)[default__.safeIndex(562, (d_1789_v0_).length(0))]])
        d_1823_v24_: _dafny.MultiSet
        d_1823_v24_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1820_v22_]), d_1822_v23_])
        d_1824_v25_: _dafny.Map
        d_1824_v25_ = _dafny.Map({d_1820_v22_: d_1822_v23_})
        d_1825_v26_: _dafny.MultiSet
        d_1825_v26_ = _dafny.MultiSet([((d_1823_v24_)[((d_1824_v25_)[d_1820_v22_] if (d_1820_v22_) in (d_1824_v25_) else d_1822_v23_)] if (((d_1824_v25_)[d_1820_v22_] if (d_1820_v22_) in (d_1824_v25_) else d_1822_v23_)) in (d_1823_v24_) else (self).f25)])
        r1 = d_1825_v26_
        d_1826_v27_: _dafny.Seq
        d_1826_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
        d_1827_v28_: D1
        d_1827_v28_ = D1_DC6(d_1815_v17_, False, d_1826_v27_)
        r2 = (d_1827_v28_).cf8
        return r0, r1, r2


class C7:
    def  __init__(self):
        self._f34: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self, f34):
        (self)._f34 = f34

    def fm17(self, p0, p1, p2, p3, globalState):
        return ((self).f34) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdlw")))))

    def m10(self, p0, p1, globalState):
        r0: int = int(0)
        d_1828_v0_: C2
        nw292_ = C2()
        nw292_.ctor__()
        d_1828_v0_ = nw292_
        d_1829_v1_: _dafny.Array
        nw293_ = _dafny.Array(int(0), 22)
        d_1829_v1_ = nw293_
        d_1830_v2_: bool
        d_1830_v2_ = False
        d_1831_v3_: _dafny.Map
        d_1831_v3_ = _dafny.Map({(self).f34: not(not(d_1830_v2_))})
        index333_ = default__.safeIndex(22, (d_1829_v1_).length(0))
        (d_1829_v1_)[index333_] = ((self).f34 if ((d_1831_v3_)[(self).f34] if ((self).f34) in (d_1831_v3_) else True) else p1)
        index334_ = default__.safeIndex(22, (d_1829_v1_).length(0))
        rhs326_ = d_1830_v2_
        rhs327_ = (self).f34
        rhs328_ = d_1829_v1_
        rhs329_ = p1
        lhs241_ = globalState
        lhs242_ = d_1829_v1_
        lhs243_ = default__.safeIndex(22, (d_1829_v1_).length(0))
        lhs241_.f12 = rhs326_
        r0 = rhs327_
        d_1829_v1_ = rhs328_
        lhs242_[lhs243_] = rhs329_
        d_1832_v4_: _dafny.Map
        d_1832_v4_ = _dafny.Map({(d_1830_v2_) and (False): 522})
        d_1833_v5_: _dafny.Set
        d_1833_v5_ = _dafny.Set({d_1830_v2_})
        d_1834_v6_: _dafny.MultiSet
        d_1834_v6_ = _dafny.MultiSet([d_1833_v5_, d_1833_v5_])
        r0 = ((d_1832_v4_)[(d_1834_v6_).ispropersubset(_dafny.MultiSet([_dafny.Set({d_1830_v2_, d_1830_v2_}), d_1833_v5_, d_1833_v5_, d_1833_v5_]))] if ((d_1834_v6_).ispropersubset(_dafny.MultiSet([_dafny.Set({d_1830_v2_, d_1830_v2_}), d_1833_v5_, d_1833_v5_, d_1833_v5_]))) in (d_1832_v4_) else (d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))])
        (globalState).f12 = d_1830_v2_
        if d_1830_v2_:
            (globalState).f12 = default__.fm0((self).f34, globalState)
            index335_ = default__.safeIndex(22, (d_1829_v1_).length(0))
            (d_1829_v1_)[index335_] = p1
            d_1835_v7_: _dafny.Seq
            d_1835_v7_ = _dafny.SeqWithoutIsStrInference([d_1830_v2_])
            d_1836_v8_: _dafny.Seq
            d_1836_v8_ = _dafny.SeqWithoutIsStrInference([(d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))], (self).f34, ((d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))]) + (len(d_1835_v7_))])
            r0 = (0) - ((d_1836_v8_)[default__.safeIndex((self).f34, len(d_1836_v8_))])
            r0 = (self).f34
            d_1837_v9_: D0
            d_1837_v9_ = D0_DC2(d_1830_v2_)
            d_1838_v10_: _dafny.Array
            nw294_ = _dafny.Array(None, 14)
            nw294_[int(0)] = d_1830_v2_
            nw294_[int(1)] = not(d_1830_v2_)
            nw294_[int(2)] = d_1830_v2_
            nw294_[int(3)] = d_1830_v2_
            nw294_[int(4)] = d_1830_v2_
            nw294_[int(5)] = not(d_1830_v2_)
            nw294_[int(6)] = d_1830_v2_
            nw294_[int(7)] = (d_1837_v9_).cf3
            nw294_[int(8)] = True
            nw294_[int(9)] = d_1830_v2_
            nw294_[int(10)] = d_1830_v2_
            nw294_[int(11)] = d_1830_v2_
            nw294_[int(12)] = True
            nw294_[int(13)] = d_1830_v2_
            d_1838_v10_ = nw294_
            d_1839_v11_: _dafny.Map
            d_1839_v11_ = _dafny.Map({True: d_1838_v10_})
            d_1839_v11_ = (d_1839_v11_) | (d_1839_v11_)
        elif True:
            d_1840_v12_: _dafny.MultiSet
            d_1840_v12_ = _dafny.MultiSet([default__.fm0((d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))], globalState)])
            d_1831_v3_ = (d_1831_v3_).set(((d_1840_v12_)[d_1830_v2_] if (d_1830_v2_) in (d_1840_v12_) else p1), (d_1830_v2_) or (d_1830_v2_))
            index336_ = default__.safeIndex(22, (d_1829_v1_).length(0))
            (d_1829_v1_)[index336_] = (d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))]
            d_1841_v13_: int
            d_1842_v14_: int
            d_1843_v15_: _dafny.Seq
            d_1844_v16_: D1
            out33_: int
            out34_: int
            out35_: _dafny.Seq
            out36_: D1
            out33_, out34_, out35_, out36_ = (d_1828_v0_).m16(globalState)
            d_1841_v13_ = out33_
            d_1842_v14_ = out34_
            d_1843_v15_ = out35_
            d_1844_v16_ = out36_
            d_1845_v17_: C0
            nw295_ = C0()
            nw295_.ctor__()
            d_1845_v17_ = nw295_
            d_1841_v13_ = (self).f34
        d_1846_v18_: _dafny.Array
        nw296_ = _dafny.Array(None, 4)
        nw296_[int(0)] = (d_1829_v1_ if d_1830_v2_ else d_1829_v1_)
        nw296_[int(1)] = d_1829_v1_
        nw296_[int(2)] = d_1829_v1_
        nw296_[int(3)] = d_1829_v1_
        d_1846_v18_ = nw296_
        index337_ = default__.safeIndex(731, (d_1846_v18_).length(0))
        (d_1846_v18_)[index337_] = d_1829_v1_
        index338_ = default__.safeIndex(731, (d_1846_v18_).length(0))
        (d_1846_v18_)[index338_] = d_1829_v1_
        d_1847_v19_: _dafny.Seq
        d_1847_v19_ = _dafny.SeqWithoutIsStrInference([77, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1848_i0_ in range(default__.abs(903))])))])
        d_1849_v20_: _dafny.Set
        d_1849_v20_ = _dafny.Set({d_1847_v19_, _dafny.SeqWithoutIsStrInference([(d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))]])})
        d_1850_v21_: D0
        d_1850_v21_ = D0_DC1(d_1830_v2_, d_1830_v2_)
        r0 = ((d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))]) - ((self).fm17((d_1829_v1_)[default__.safeIndex(22, (d_1829_v1_).length(0))], d_1849_v20_, d_1850_v21_, p0, globalState))
        return r0

    @property
    def f34(self):
        return self._f34

class C8(T1, T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        self._f32: bool = False
        self._f33: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f32, f33, f25, f26):
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f25 = f25
        (self)._f26 = f26

    def fm13(self, globalState):
        source25_ = D0_DC3((self).f25, (self).f33)
        if source25_.is_DC1:
            d_1851___mcc_h0_ = source25_.cf1
            d_1852___mcc_h1_ = source25_.cf2
            d_1853_cf2_ = d_1852___mcc_h1_
            d_1854_cf1_ = d_1851___mcc_h0_
            return (self).f32
        elif source25_.is_DC2:
            d_1855___mcc_h2_ = source25_.cf3
            d_1856_cf3_ = d_1855___mcc_h2_
            return d_1856_cf3_
        elif source25_.is_DC3:
            d_1857___mcc_h3_ = source25_.cf4
            d_1858___mcc_h4_ = source25_.cf5
            d_1859_cf5_ = d_1858___mcc_h4_
            d_1860_cf4_ = d_1857___mcc_h3_
            return (self).f32
        elif True:
            d_1861___mcc_h5_ = source25_.cf0
            d_1862_cf0_ = d_1861___mcc_h5_
            return not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgkclqti"))) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1863_i0_ in range(default__.abs(652))])))

    def fm14(self, p0, globalState):
        return (D0_DC3((self).f25, False)).cf4

    def m4(self, p0, globalState):
        hi9_ = (self).f25
        for d_1864_i0_ in range((self).f25, hi9_):
            d_1865_v0_: _dafny.Seq
            d_1865_v0_ = _dafny.SeqWithoutIsStrInference([(self).f33])
            d_1866_v1_: _dafny.MultiSet
            d_1866_v1_ = _dafny.MultiSet([(self).f32])
            d_1867_v2_: _dafny.Map
            d_1867_v2_ = _dafny.Map({d_1864_i0_: (d_1866_v1_).cardinality})
            d_1868_v3_: _dafny.Array
            def lambda122_(d_1869_i1_):
                return D1_DC7(D1_DC6(_dafny.CodePoint('t'), (self).f32, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwpkpvo"))))

            init68_ = lambda122_
            nw297_ = _dafny.Array(None, 2)
            for i0_68_ in range(nw297_.length(0)):
                nw297_[i0_68_] = init68_(i0_68_)
            d_1868_v3_ = nw297_
            d_1870_v4_: _dafny.Array
            nw298_ = _dafny.Array(None, 27)
            nw298_[int(0)] = len(d_1865_v0_)
            nw298_[int(1)] = len(default__.fm1((self).f33, globalState))
            nw298_[int(2)] = (0) - ((self).f25)
            nw298_[int(3)] = 866
            nw298_[int(4)] = d_1864_i0_
            nw298_[int(5)] = d_1864_i0_
            nw298_[int(6)] = (0) - (d_1864_i0_)
            nw298_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcib")))
            nw298_[int(8)] = (self).f25
            nw298_[int(9)] = ((d_1867_v2_)[d_1864_i0_] if (d_1864_i0_) in (d_1867_v2_) else (self).f25)
            nw298_[int(10)] = (self).f25
            nw298_[int(11)] = -58
            nw298_[int(12)] = (882) - ((self).fm14((self).f33, globalState))
            nw298_[int(13)] = d_1864_i0_
            nw298_[int(14)] = (self).f25
            nw298_[int(15)] = (self).f25
            nw298_[int(16)] = (self).fm14(default__.fm0(len(d_1865_v0_), globalState), globalState)
            nw298_[int(17)] = (self).f25
            nw298_[int(18)] = d_1864_i0_
            nw298_[int(19)] = (self).f25
            nw298_[int(20)] = len(_dafny.Map({d_1868_v3_: (self).f25}))
            nw298_[int(21)] = (d_1864_i0_) * ((self).f25)
            nw298_[int(22)] = 575
            nw298_[int(23)] = d_1864_i0_
            nw298_[int(24)] = (self).f25
            nw298_[int(25)] = 276
            nw298_[int(26)] = len(_dafny.Set({(self).f32, True, (self).f32, (self).f32, (self).f33}))
            d_1870_v4_ = nw298_
            index339_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            (d_1870_v4_)[index339_] = 424
            d_1871_v5_: _dafny.Seq
            d_1871_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsr"))
            index340_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            (d_1870_v4_)[index340_] = len(d_1871_v5_)
            d_1872_v6_: _dafny.Seq
            d_1872_v6_ = _dafny.SeqWithoutIsStrInference([(d_1870_v4_)[default__.safeIndex(106, (d_1870_v4_).length(0))]])
            d_1873_v7_: _dafny.MultiSet
            d_1873_v7_ = _dafny.MultiSet([len(d_1872_v6_)])
            (globalState).f12 = (d_1873_v7_).isdisjoint(default__.fm15(globalState))
            index341_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            rhs330_ = (0) - ((391) - ((len(d_1872_v6_)) - ((d_1870_v4_)[default__.safeIndex(106, (d_1870_v4_).length(0))])))
            rhs331_ = (self).f32
            lhs244_ = d_1870_v4_
            lhs245_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            lhs246_ = globalState
            lhs244_[lhs245_] = rhs330_
            lhs246_.f12 = rhs331_
            d_1874_v8_: _dafny.Set
            d_1874_v8_ = _dafny.Set({(0) - (len(d_1871_v5_))})
            d_1875_v9_: D0
            d_1875_v9_ = D0_DC2(not((self).f33))
            d_1876_v10_: str
            d_1876_v10_ = _dafny.CodePoint('n')
            d_1877_v11_: D1
            d_1877_v11_ = D1_DC6(d_1876_v10_, (self).f33, d_1871_v5_)
            d_1878_v12_: _dafny.Map
            d_1878_v12_ = _dafny.Map({d_1876_v10_: (d_1870_v4_)[default__.safeIndex(106, (d_1870_v4_).length(0))]})
            d_1879_v13_: _dafny.Seq
            d_1879_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1864_i0_ for d_1880_i2_ in range(default__.abs(134))])])
            d_1881_v14_: _dafny.Seq
            d_1881_v14_ = _dafny.SeqWithoutIsStrInference([((d_1871_v5_).set(default__.safeIndex(982, len(d_1871_v5_)), d_1876_v10_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_1870_v4_)[default__.safeIndex(106, (d_1870_v4_).length(0))]])), len((d_1871_v5_).set(default__.safeIndex(982, len(d_1871_v5_)), d_1876_v10_))), d_1876_v10_)])
            pat_let_tv38_ = d_1871_v5_
            index342_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            index343_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            def iife152_(_pat_let41_0):
                def iife153_(d_1882_dt__update__tmp_h0_):
                    def iife154_(_pat_let42_0):
                        def iife155_(d_1883_dt__update_hcf9_h0_):
                            return D1_DC6((d_1882_dt__update__tmp_h0_).cf7, (d_1882_dt__update__tmp_h0_).cf8, d_1883_dt__update_hcf9_h0_)
                        return iife155_(_pat_let42_0)
                    return iife154_(pat_let_tv38_)
                return iife153_(_pat_let41_0)
            rhs332_ = default__.fm16(d_1875_v9_, (self).f32, (self).f32, (iife152_(d_1877_v11_)).cf7, globalState)
            rhs333_ = (not((d_1864_i0_) != (((d_1878_v12_)[d_1876_v10_] if (d_1876_v10_) in (d_1878_v12_) else len((d_1879_v13_)[default__.safeIndex(874, len(d_1879_v13_))]))))) in (d_1865_v0_)
            rhs334_ = _dafny.SeqWithoutIsStrInference([(self).f32, not(not ((self).f32) or ((self).f33)), (self).f33, (self).f32])
            rhs335_ = default__.safeDivisionInt(31, (0) - ((0) - ((d_1864_i0_) - ((self).f25))))
            rhs336_ = len((d_1881_v14_)[default__.safeIndex((len(d_1871_v5_)) - (len(d_1871_v5_)), len(d_1881_v14_))])
            lhs247_ = globalState
            lhs248_ = d_1870_v4_
            lhs249_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            lhs250_ = d_1870_v4_
            lhs251_ = default__.safeIndex(106, (d_1870_v4_).length(0))
            d_1874_v8_ = rhs332_
            lhs247_.f12 = rhs333_
            d_1865_v0_ = rhs334_
            lhs248_[lhs249_] = rhs335_
            lhs250_[lhs251_] = rhs336_
        d_1884_v15_: _dafny.Set
        d_1884_v15_ = _dafny.Set({(self).f25, (self).f25})
        hi10_ = (self).f25
        for d_1885_i3_ in range((len(d_1884_v15_)) - (609), hi10_):
            d_1886_v16_: C5
            nw299_ = C5()
            nw299_.ctor__((self).f32)
            d_1886_v16_ = nw299_
            d_1887_v17_: _dafny.Set
            d_1887_v17_ = _dafny.Set({(self).f32})
            d_1888_v18_: _dafny.Seq
            d_1888_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yh"))
            d_1889_v19_: _dafny.Map
            d_1889_v19_ = _dafny.Map({((self).f25) - (len(d_1887_v17_)): len(d_1888_v18_)})
            d_1889_v19_ = (d_1889_v19_).set(d_1885_i3_, (self).fm14((self).f32, globalState))
            d_1890_v20_: int
            d_1890_v20_ = 308
            d_1891_v21_: _dafny.MultiSet
            d_1891_v21_ = _dafny.MultiSet([not(True)])
            d_1890_v20_ = default__.fm44((d_1891_v21_).ispropersubset(d_1891_v21_), (self).f33, (self).fm13(globalState), -445, globalState)
            d_1892_v22_: _dafny.Array
            nw300_ = _dafny.Array(None, 2)
            nw300_[int(0)] = d_1885_i3_
            nw300_[int(1)] = (self).f25
            d_1892_v22_ = nw300_
            d_1892_v22_ = d_1892_v22_
        if (default__.safeModuloInt((0) - ((self).f25), (self).f25)) >= ((0) - (default__.safeDivisionInt(396, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1893_i4_ in range(default__.abs(157))]))))):
            d_1894_v23_: int
            d_1894_v23_ = 56
            d_1894_v23_ = (self).fm14((self).f33, globalState)
            d_1895_v24_: _dafny.Seq
            d_1895_v24_ = _dafny.SeqWithoutIsStrInference([False])
            d_1896_v25_: _dafny.Map
            d_1896_v25_ = _dafny.Map({(len(d_1895_v24_)) * (753): d_1894_v23_})
            d_1897_v26_: _dafny.MultiSet
            d_1897_v26_ = _dafny.MultiSet([(self).fm14(True, globalState)])
            d_1898_v27_: _dafny.MultiSet
            d_1898_v27_ = _dafny.MultiSet([d_1897_v26_])
            d_1896_v25_ = (d_1896_v25_).set(d_1894_v23_, len(default__.fm37(d_1898_v27_, globalState)))
            d_1899_v28_: _dafny.Seq
            d_1899_v28_ = _dafny.SeqWithoutIsStrInference([d_1894_v23_, d_1894_v23_, (self).f25, (self).f25, len(d_1895_v24_)])
            (globalState).f12 = ((0) - ((0) - ((self).f25))) in (d_1899_v28_)
            d_1900_v29_: _dafny.Array
            nw301_ = _dafny.Array(int(0), 19)
            d_1900_v29_ = nw301_
            d_1901_v30_: D4
            d_1901_v30_ = D4_DC13((self).f25, (self).f33, (self).f33, (self).f33)
            d_1902_v31_: _dafny.Map
            d_1902_v31_ = _dafny.Map({(d_1901_v30_).cf20: not((self).f33)})
            d_1903_v32_: _dafny.Map
            d_1903_v32_ = d_1902_v31_
            d_1904_v33_: _dafny.Seq
            d_1904_v33_ = _dafny.SeqWithoutIsStrInference([d_1903_v32_, d_1903_v32_])
            index344_ = default__.safeIndex(842, (d_1900_v29_).length(0))
            (d_1900_v29_)[index344_] = len(d_1904_v33_)
            index345_ = default__.safeIndex(842, (d_1900_v29_).length(0))
            (d_1900_v29_)[index345_] = d_1894_v23_
            def iife156_():
                coll70_ = _dafny.Set()
                compr_70_: int
                for compr_70_ in (_dafny.Set({((d_1896_v25_)[(0) - ((self).f25)] if ((0) - ((self).f25)) in (d_1896_v25_) else d_1894_v23_)})).Elements:
                    d_1905_v34_: int = compr_70_
                    if (d_1905_v34_) in (_dafny.Set({((d_1896_v25_)[(0) - ((self).f25)] if ((0) - ((self).f25)) in (d_1896_v25_) else d_1894_v23_)})):
                        coll70_ = coll70_.union(_dafny.Set([(d_1905_v34_) * (len(_dafny.SeqWithoutIsStrInference([391, 821])))]))
                return _dafny.Set(coll70_)
            if (iife156_()
            ).issubset(d_1884_v15_):
                d_1894_v23_ = (d_1900_v29_)[default__.safeIndex(842, (d_1900_v29_).length(0))]
                index346_ = default__.safeIndex(842, (d_1900_v29_).length(0))
                (d_1900_v29_)[index346_] = default__.safeModuloInt(302, (d_1900_v29_)[default__.safeIndex(842, (d_1900_v29_).length(0))])
                d_1906_v35_: D5
                d_1906_v35_ = D5_DC15(d_1897_v26_)
                d_1907_v36_: _dafny.Array
                nw302_ = _dafny.Array(False, 11)
                d_1907_v36_ = nw302_
                index347_ = default__.safeIndex(538, (d_1907_v36_).length(0))
                (d_1907_v36_)[index347_] = (self).f33
                d_1908_v37_: _dafny.Seq
                d_1908_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsjsoh"))
                index348_ = default__.safeIndex(538, (d_1907_v36_).length(0))
                rhs337_ = d_1906_v35_
                rhs338_ = d_1908_v37_
                rhs339_ = (d_1894_v23_) == ((d_1900_v29_)[default__.safeIndex(842, (d_1900_v29_).length(0))])
                rhs340_ = not(not((self).f33))
                lhs252_ = globalState
                lhs253_ = globalState
                lhs254_ = d_1907_v36_
                lhs255_ = default__.safeIndex(538, (d_1907_v36_).length(0))
                d_1906_v35_ = rhs337_
                lhs252_.f10 = rhs338_
                lhs253_.f12 = rhs339_
                lhs254_[lhs255_] = rhs340_
                (globalState).f12 = default__.fm0(((self).f25) - ((self).f25), globalState)
                d_1909_v38_: C2
                nw303_ = C2()
                nw303_.ctor__()
                d_1909_v38_ = nw303_
            elif True:
                d_1910_v39_: _dafny.MultiSet
                d_1910_v39_ = _dafny.MultiSet([(self).f32])
                d_1894_v23_ = ((0) - ((d_1910_v39_).cardinality)) - (d_1894_v23_)
                d_1894_v23_ = (d_1900_v29_)[default__.safeIndex(842, (d_1900_v29_).length(0))]
                d_1911_v40_: _dafny.Array
                nw304_ = _dafny.Array(_dafny.Set({}), 11)
                d_1911_v40_ = nw304_
                d_1912_v41_: _dafny.Map
                d_1912_v41_ = _dafny.Map({(self).f32: d_1911_v40_})
                d_1913_v42_: _dafny.Seq
                d_1913_v42_ = _dafny.SeqWithoutIsStrInference([((d_1912_v41_)[(self).f33] if ((self).f33) in (d_1912_v41_) else d_1911_v40_), d_1911_v40_, d_1911_v40_])
                d_1911_v40_ = (d_1913_v42_)[default__.safeIndex(((d_1900_v29_)[default__.safeIndex(842, (d_1900_v29_).length(0))]) + (len(d_1895_v24_)), len(d_1913_v42_))]
                index349_ = default__.safeIndex(842, (d_1900_v29_).length(0))
                (d_1900_v29_)[index349_] = 744
                d_1900_v29_ = d_1900_v29_
        elif True:
            d_1914_v43_: D9
            d_1914_v43_ = D9_DC27((self).f25, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1915_i5_ in range(default__.abs(704))])), (self).f25))
            source26_ = d_1914_v43_
            if source26_.is_DC27:
                d_1916___mcc_h0_ = source26_.cf39
                d_1917___mcc_h1_ = source26_.cf40
                d_1918_cf40_ = d_1917___mcc_h1_
                d_1919_cf39_ = d_1916___mcc_h0_
                d_1920_v44_: _dafny.Seq
                d_1920_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgikbaco"))
                d_1919_cf39_ = ((d_1918_cf40_ if True else (0) - (len(_dafny.Map({_dafny.MultiSet([(self).f33]): (self).f33}))))) + ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1921_i6_ in range(default__.abs(-569))]))) * ((0) - (len(d_1920_v44_))))
                d_1922_v45_: _dafny.Array
                def lambda123_(d_1923_i7_):
                    return (215) >= (492)

                init69_ = lambda123_
                nw305_ = _dafny.Array(None, 23)
                for i0_69_ in range(nw305_.length(0)):
                    nw305_[i0_69_] = init69_(i0_69_)
                d_1922_v45_ = nw305_
                index350_ = default__.safeIndex(262, (d_1922_v45_).length(0))
                (d_1922_v45_)[index350_] = (self).f33
                d_1924_v46_: _dafny.Array
                nw306_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_1924_v46_ = nw306_
                d_1925_v47_: _dafny.Array
                def lambda124_(d_1926_cf39_):
                    def lambda125_(d_1927_i8_):
                        return default__.safeDivisionInt(d_1927_i8_, len(_dafny.Map({d_1926_cf39_: _dafny.Map({(self).f33: (self).f33})})))

                    return lambda125_

                init70_ = lambda124_(d_1919_cf39_)
                nw307_ = _dafny.Array(None, 7)
                for i0_70_ in range(nw307_.length(0)):
                    nw307_[i0_70_] = init70_(i0_70_)
                d_1925_v47_ = nw307_
                index351_ = default__.safeIndex(146, (d_1924_v46_).length(0))
                (d_1924_v46_)[index351_] = d_1925_v47_
                d_1928_v48_: _dafny.Map
                d_1928_v48_ = _dafny.Map({d_1918_cf40_: d_1919_cf39_})
                d_1929_v49_: str
                d_1929_v49_ = _dafny.CodePoint('i')
                d_1930_v50_: _dafny.Map
                d_1930_v50_ = _dafny.Map({d_1929_v49_: d_1925_v47_})
                d_1931_v51_: _dafny.Map
                d_1931_v51_ = _dafny.Map({d_1919_cf39_: False})
                d_1932_v52_: _dafny.Map
                d_1932_v52_ = _dafny.Map({(self).f25: not (((d_1931_v51_)[(self).f25] if ((self).f25) in (d_1931_v51_) else (self).f32)) or ((self).f33)})
                d_1933_v53_: _dafny.Map
                d_1933_v53_ = _dafny.Map({False: (self).f32})
                d_1934_v54_: _dafny.Seq
                d_1934_v54_ = _dafny.SeqWithoutIsStrInference([(self).f33, (self).f32])
                d_1935_v55_: _dafny.Map
                d_1935_v55_ = _dafny.Map({len((d_1934_v54_) + (d_1934_v54_)): d_1925_v47_})
                index352_ = default__.safeIndex(262, (d_1922_v45_).length(0))
                index353_ = default__.safeIndex(146, (d_1924_v46_).length(0))
                def iife157_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(104, -568):
                        d_1936_v56_: int = compr_71_
                        if ((104) <= (d_1936_v56_)) and ((d_1936_v56_) < (-568)):
                            coll71_[default__.safeModuloInt(d_1936_v56_, (self).f25)] = -573
                    return _dafny.Map(coll71_)
                rhs341_ = (len(d_1930_v50_)) >= ((self).f25)
                rhs342_ = ((d_1931_v51_)[default__.safeModuloInt(len(d_1933_v53_), d_1918_cf40_)] if (default__.safeModuloInt(len(d_1933_v53_), d_1918_cf40_)) in (d_1931_v51_) else ((self).f32) or (not((self).f32)))
                rhs343_ = default__.safeModuloInt(len(d_1934_v54_), -823)
                rhs344_ = ((d_1935_v55_)[len(d_1928_v48_)] if (len(d_1928_v48_)) in (d_1935_v55_) else d_1925_v47_)
                rhs345_ = (d_1928_v48_) | (iife157_()
                )
                lhs256_ = globalState
                lhs257_ = d_1922_v45_
                lhs258_ = default__.safeIndex(262, (d_1922_v45_).length(0))
                lhs259_ = d_1924_v46_
                lhs260_ = default__.safeIndex(146, (d_1924_v46_).length(0))
                lhs256_.f12 = rhs341_
                lhs257_[lhs258_] = rhs342_
                d_1919_cf39_ = rhs343_
                lhs259_[lhs260_] = rhs344_
                d_1928_v48_ = rhs345_
                d_1928_v48_ = (d_1928_v48_).set(d_1918_cf40_, d_1918_cf40_)
                d_1918_cf40_ = 822
            elif True:
                d_1937___mcc_h2_ = source26_.cf38
                d_1938_cf38_ = d_1937___mcc_h2_
                d_1939_v57_: _dafny.Array
                nw308_ = _dafny.Array(False, 7)
                d_1939_v57_ = nw308_
                index354_ = default__.safeIndex(113, (d_1939_v57_).length(0))
                (d_1939_v57_)[index354_] = True
                index355_ = default__.safeIndex(113, (d_1939_v57_).length(0))
                (d_1939_v57_)[index355_] = not(not((self).f32))
                d_1940_v58_: _dafny.Array
                nw309_ = _dafny.Array(_dafny.Map({}), 25)
                d_1940_v58_ = nw309_
                index356_ = default__.safeIndex(587, (d_1940_v58_).length(0))
                (d_1940_v58_)[index356_] = _dafny.Map({(self).f25: (self).f33})
                d_1941_v59_: _dafny.Map
                d_1941_v59_ = _dafny.Map({(0) - ((self).f25): (d_1939_v57_)[default__.safeIndex(113, (d_1939_v57_).length(0))]})
                index357_ = default__.safeIndex(587, (d_1940_v58_).length(0))
                (d_1940_v58_)[index357_] = (d_1941_v59_) | (d_1941_v59_)
                d_1942_v60_: int
                d_1942_v60_ = 238
                d_1942_v60_ = (self).f25
                d_1938_cf38_ = (d_1938_cf38_).set((self).fm13(globalState), (self).f25)
            d_1943_v61_: int
            d_1943_v61_ = -292
            d_1944_v62_: _dafny.Seq
            d_1944_v62_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32, (self).f33])
            d_1943_v61_ = default__.safeModuloInt((self).f25, len((d_1944_v62_) + (d_1944_v62_)))
            d_1945_v63_: str
            d_1945_v63_ = _dafny.CodePoint('w')
            (globalState).f23 = d_1945_v63_
            (globalState).f12 = (self).f33
            d_1946_v64_: _dafny.Array
            nw310_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_1946_v64_ = nw310_
            d_1946_v64_ = d_1946_v64_
        d_1947_v65_: _dafny.Seq
        d_1947_v65_ = _dafny.SeqWithoutIsStrInference([(self).f33])
        d_1948_v66_: _dafny.MultiSet
        d_1948_v66_ = _dafny.MultiSet([(self).f33, (self).f33, not((self).f33)])
        d_1949_v69_: _dafny.Map
        def iife158_():
            coll72_ = _dafny.Set()
            compr_72_: int
            for compr_72_ in (_dafny.MultiSet([(self).f25, (self).f25])).Elements:
                d_1950_v67_: int = compr_72_
                if (d_1950_v67_) in (_dafny.MultiSet([(self).f25, (self).f25])):
                    def iife159_():
                        coll73_ = _dafny.Map()
                        compr_73_: int
                        for compr_73_ in _dafny.IntegerRange(64, -997):
                            d_1951_v68_: int = compr_73_
                            if ((64) <= (d_1951_v68_)) and ((d_1951_v68_) < (-997)):
                                coll73_[default__.safeModuloInt(d_1951_v68_, 258)] = 365
                        return _dafny.Map(coll73_)
                    coll72_ = coll72_.union(_dafny.Set([default__.safeModuloInt(d_1950_v67_, len(_dafny.Map({len(iife159_()
): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksgs")))})))]))
            return _dafny.Set(coll72_)
        d_1949_v69_ = _dafny.Map({len(default__.fm64((self).f25, (self).f32, _dafny.CodePoint('q'), (self).f32, globalState)): len(iife158_()
        )})
        d_1952_v70_: _dafny.Map
        d_1952_v70_ = _dafny.Map({(_dafny.MultiSet(d_1947_v65_)).isdisjoint(d_1948_v66_): ((0) - (((d_1949_v69_)[(self).f25] if ((self).f25) in (d_1949_v69_) else (self).f25))) + ((self).f25)})
        d_1952_v70_ = (d_1952_v70_).set((self).f33, (0) - ((-165) + ((self).f25)))
        if (self).f32:
            d_1953_v71_: _dafny.Set
            d_1953_v71_ = _dafny.Set({(self).f32})
            (globalState).f12 = (d_1953_v71_).issubset(d_1953_v71_)
            d_1954_v72_: int
            d_1954_v72_ = 75
            d_1954_v72_ = (default__.safeDivisionInt(d_1954_v72_, 877) if (self).f32 else d_1954_v72_)
            d_1955_v73_: C2
            nw311_ = C2()
            nw311_.ctor__()
            d_1955_v73_ = nw311_
            d_1953_v71_ = (_dafny.Set({default__.fm0(d_1954_v72_, globalState)})).intersection(d_1953_v71_)
            d_1956_v74_: _dafny.Array
            def lambda126_(d_1957_v72_):
                def lambda127_(d_1958_i9_):
                    return (d_1958_i9_) * (d_1957_v72_)

                return lambda127_

            init71_ = lambda126_(d_1954_v72_)
            nw312_ = _dafny.Array(None, 3)
            for i0_71_ in range(nw312_.length(0)):
                nw312_[i0_71_] = init71_(i0_71_)
            d_1956_v74_ = nw312_
            d_1959_v75_: _dafny.Array
            nw313_ = _dafny.Array(None, 16)
            nw313_[int(0)] = (self).f33
            nw313_[int(1)] = (self).f32
            nw313_[int(2)] = (self).f32
            nw313_[int(3)] = not((self).f33)
            nw313_[int(4)] = (self).f33
            nw313_[int(5)] = (self).f33
            nw313_[int(6)] = (self).f32
            nw313_[int(7)] = (self).f32
            nw313_[int(8)] = (self).f32
            nw313_[int(9)] = (self).f33
            nw313_[int(10)] = (self).f32
            nw313_[int(11)] = (self).f32
            nw313_[int(12)] = True
            nw313_[int(13)] = (self).f32
            nw313_[int(14)] = (self).f32
            nw313_[int(15)] = False
            d_1959_v75_ = nw313_
            d_1960_v76_: C5
            nw314_ = C5()
            nw314_.ctor__(False)
            d_1960_v76_ = nw314_
            d_1961_v77_: _dafny.Array
            nw315_ = _dafny.Array(None, 19)
            nw315_[int(0)] = d_1960_v76_
            nw315_[int(1)] = d_1960_v76_
            nw315_[int(2)] = d_1960_v76_
            nw315_[int(3)] = d_1960_v76_
            nw315_[int(4)] = d_1960_v76_
            nw315_[int(5)] = d_1960_v76_
            nw315_[int(6)] = d_1960_v76_
            nw315_[int(7)] = d_1960_v76_
            nw315_[int(8)] = d_1960_v76_
            nw315_[int(9)] = d_1960_v76_
            nw315_[int(10)] = d_1960_v76_
            nw315_[int(11)] = d_1960_v76_
            nw315_[int(12)] = d_1960_v76_
            nw315_[int(13)] = d_1960_v76_
            nw315_[int(14)] = d_1960_v76_
            nw315_[int(15)] = d_1960_v76_
            nw315_[int(16)] = d_1960_v76_
            nw315_[int(17)] = d_1960_v76_
            nw315_[int(18)] = d_1960_v76_
            d_1961_v77_ = nw315_
            d_1962_v78_: _dafny.Map
            d_1962_v78_ = _dafny.Map({d_1959_v75_: d_1961_v77_})
            index358_ = default__.safeIndex(832, (d_1956_v74_).length(0))
            (d_1956_v74_)[index358_] = len(d_1962_v78_)
            index359_ = default__.safeIndex(832, (d_1956_v74_).length(0))
            (d_1956_v74_)[index359_] = default__.safeModuloInt(d_1954_v72_, d_1954_v72_)
        elif True:
            d_1963_v79_: int
            d_1963_v79_ = -719
            d_1964_v80_: D12
            d_1964_v80_ = D12_DC32((self).f32)
            d_1965_v82_: _dafny.Seq
            def iife160_():
                coll74_ = _dafny.Set()
                compr_74_: int
                for compr_74_ in _dafny.IntegerRange(448, 957):
                    d_1966_v81_: int = compr_74_
                    if ((448) <= (d_1966_v81_)) and ((d_1966_v81_) < (957)):
                        coll74_ = coll74_.union(_dafny.Set([default__.safeDivisionInt(d_1966_v81_, (self).f25)]))
                return _dafny.Set(coll74_)
            d_1965_v82_ = _dafny.SeqWithoutIsStrInference([d_1963_v79_, len(iife160_()
            ), d_1963_v79_])
            d_1967_v83_: _dafny.MultiSet
            d_1967_v83_ = _dafny.MultiSet([(self).f25, len(d_1965_v82_)])
            rhs346_ = (d_1967_v83_) == (d_1967_v83_)
            rhs347_ = default__.safeDivisionInt((self).f25, d_1963_v79_)
            rhs348_ = d_1964_v80_
            lhs261_ = globalState
            lhs261_.f12 = rhs346_
            d_1963_v79_ = rhs347_
            d_1964_v80_ = rhs348_
            (globalState).f6 = (d_1965_v82_).set(default__.safeIndex(-234, len(d_1965_v82_)), d_1963_v79_)
            d_1968_v84_: _dafny.Set
            d_1968_v84_ = _dafny.Set({(self).f33, (self).f32})
            d_1949_v69_ = (d_1949_v69_).set(len(d_1965_v82_), (len(d_1968_v84_)) + (289))
            d_1969_v85_: D10
            d_1969_v85_ = D10_DC29((self).f32, (self).f33, len(d_1884_v15_))
            d_1970_v86_: _dafny.Array
            nw316_ = _dafny.Array(None, 1)
            def iife161_(_pat_let43_0):
                def iife162_(d_1971_dt__update__tmp_h1_):
                    def iife163_(_pat_let44_0):
                        def iife164_(d_1972_dt__update_hcf43_h0_):
                            return D10_DC29((d_1971_dt__update__tmp_h1_).cf42, d_1972_dt__update_hcf43_h0_, (d_1971_dt__update__tmp_h1_).cf44)
                        return iife164_(_pat_let44_0)
                    return iife163_((self).f33)
                return iife162_(_pat_let43_0)
            nw316_[int(0)] = iife161_(d_1969_v85_)
            d_1970_v86_ = nw316_
            index360_ = default__.safeIndex(559, (d_1970_v86_).length(0))
            (d_1970_v86_)[index360_] = d_1969_v85_
            d_1973_v87_: _dafny.Array
            nw317_ = _dafny.Array(_dafny.Seq({}), 16)
            d_1973_v87_ = nw317_
            index361_ = default__.safeIndex(783, (d_1973_v87_).length(0))
            (d_1973_v87_)[index361_] = (_dafny.SeqWithoutIsStrInference([(self).f25 for d_1974_i10_ in range(default__.abs(-828))]) if (self).f32 else d_1965_v82_)
            index362_ = default__.safeIndex(559, (d_1970_v86_).length(0))
            index363_ = default__.safeIndex(783, (d_1973_v87_).length(0))
            rhs349_ = default__.safeModuloInt(default__.safeDivisionInt(d_1963_v79_, (self).f25), (0) - (-478))
            rhs350_ = d_1969_v85_
            rhs351_ = default__.fm19(((self).f32) == (True), globalState)
            lhs262_ = d_1970_v86_
            lhs263_ = default__.safeIndex(559, (d_1970_v86_).length(0))
            lhs264_ = d_1973_v87_
            lhs265_ = default__.safeIndex(783, (d_1973_v87_).length(0))
            d_1963_v79_ = rhs349_
            lhs262_[lhs263_] = rhs350_
            lhs264_[lhs265_] = rhs351_
            d_1967_v83_ = (_dafny.MultiSet([(self).f25])).intersection(d_1967_v83_)
        d_1975_v88_: _dafny.Seq
        d_1975_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnyeciygy"))
        d_1976_v89_: _dafny.MultiSet
        d_1976_v89_ = _dafny.MultiSet([len(d_1975_v88_)])
        d_1977_v90_: _dafny.MultiSet
        d_1977_v90_ = _dafny.MultiSet([len(_dafny.Set({d_1976_v89_})), len(d_1975_v88_), (self).f25, 969])
        d_1978_v91_: _dafny.Seq
        d_1978_v91_ = _dafny.SeqWithoutIsStrInference([(self).f25, 578, (self).f25])
        d_1979_v92_: _dafny.Set
        d_1979_v92_ = _dafny.Set({(self).f33})
        d_1980_v93_: _dafny.Array
        nw318_ = _dafny.Array(None, 25)
        nw318_[int(0)] = (self).f25
        nw318_[int(1)] = (self).f25
        nw318_[int(2)] = (self).f25
        nw318_[int(3)] = (self).f25
        nw318_[int(4)] = (d_1976_v89_).cardinality
        nw318_[int(5)] = (d_1978_v91_)[default__.safeIndex((self).f25, len(d_1978_v91_))]
        nw318_[int(6)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1981_i11_ in range(default__.abs(994))]))
        nw318_[int(7)] = (self).f25
        nw318_[int(8)] = (self).f25
        nw318_[int(9)] = len(_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33]))
        nw318_[int(10)] = (self).f25
        nw318_[int(11)] = (self).f25
        nw318_[int(12)] = len(d_1979_v92_)
        nw318_[int(13)] = len(d_1975_v88_)
        nw318_[int(14)] = (self).f25
        nw318_[int(15)] = (self).f25
        nw318_[int(16)] = 317
        nw318_[int(17)] = (self).f25
        nw318_[int(18)] = (self).f25
        nw318_[int(19)] = 874
        nw318_[int(20)] = (self).f25
        nw318_[int(21)] = (self).f25
        nw318_[int(22)] = -466
        nw318_[int(23)] = (self).f25
        nw318_[int(24)] = 57
        d_1980_v93_ = nw318_
        d_1982_v94_: _dafny.Seq
        d_1982_v94_ = _dafny.SeqWithoutIsStrInference([d_1980_v93_, d_1980_v93_])
        (globalState).f23 = default__.fm33((self).f25, (self).f32, len(d_1982_v94_), globalState)

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_1983_v0_: _dafny.Array
        def lambda128_(d_1984_p1_, d_1985_p2_):
            def lambda129_(d_1986_i1_):
                return (d_1986_i1_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1984_p1_, d_1985_p2_, (self).f33]) for d_1987_i2_ in range(default__.abs(96))])))

            return lambda129_

        init72_ = lambda128_(p1, p2)
        nw319_ = _dafny.Array(None, 20)
        for i0_72_ in range(nw319_.length(0)):
            nw319_[i0_72_] = init72_(i0_72_)
        d_1983_v0_ = nw319_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_1983_v0_).length(0)):
            d_1988_i0_: int = guard_loop_10_
            if (True) and (((0) <= (d_1988_i0_)) and ((d_1988_i0_) < ((d_1983_v0_).length(0)))):
                (d_1983_v0_)[(d_1988_i0_)] = default__.safeDivisionInt(d_1988_i0_, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1989_i3_ in range(default__.abs(580))])), (self).f25))
        d_1990_v1_: C2
        nw320_ = C2()
        nw320_.ctor__()
        d_1990_v1_ = nw320_
        d_1991_v2_: _dafny.Set
        d_1991_v2_ = _dafny.Set({p1})
        rhs352_ = default__.safeDivisionInt((0) - (-506), ((self).f25) + (len(d_1991_v2_)))
        rhs353_ = (default__.fm20(p2, globalState)).ispropersubset(_dafny.Set({p1, (self).f32}))
        rhs354_ = d_1990_v1_
        lhs266_ = globalState
        r0 = rhs352_
        lhs266_.f12 = rhs353_
        d_1990_v1_ = rhs354_
        d_1992_v3_: D10
        d_1992_v3_ = D10_DC29(p1, True, (self).f25)
        hi11_ = (self).f25
        for d_1993_i4_ in range((0) - ((d_1992_v3_).cf44), hi11_):
            d_1994_v4_: _dafny.Array
            nw321_ = _dafny.Array(False, 29)
            d_1994_v4_ = nw321_
            rhs355_ = d_1994_v4_
            rhs356_ = p2
            lhs267_ = globalState
            d_1994_v4_ = rhs355_
            lhs267_.f12 = rhs356_
            index364_ = default__.safeIndex(659, (p3).length(0))
            (p3)[index364_] = d_1993_i4_
            d_1995_v5_: _dafny.Seq
            d_1995_v5_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p2: p2}))])
            index365_ = default__.safeIndex(659, (p3).length(0))
            (p3)[index365_] = (0) - (default__.safeDivisionInt(d_1993_i4_, (0) - (default__.safeModuloInt((self).f25, (d_1995_v5_)[default__.safeIndex(d_1993_i4_, len(d_1995_v5_))]))))
            index366_ = default__.safeIndex(186, (d_1994_v4_).length(0))
            (d_1994_v4_)[index366_] = (self).f33
            index367_ = default__.safeIndex(186, (d_1994_v4_).length(0))
            index368_ = default__.safeIndex(659, (p3).length(0))
            index369_ = default__.safeIndex(659, (p3).length(0))
            rhs357_ = True
            rhs358_ = (p3)[default__.safeIndex(659, (p3).length(0))]
            rhs359_ = (self).f32
            rhs360_ = (p3)[default__.safeIndex(659, (p3).length(0))]
            lhs268_ = d_1994_v4_
            lhs269_ = default__.safeIndex(186, (d_1994_v4_).length(0))
            lhs270_ = p3
            lhs271_ = default__.safeIndex(659, (p3).length(0))
            lhs272_ = globalState
            lhs273_ = p3
            lhs274_ = default__.safeIndex(659, (p3).length(0))
            lhs268_[lhs269_] = rhs357_
            lhs270_[lhs271_] = rhs358_
            lhs272_.f12 = rhs359_
            lhs273_[lhs274_] = rhs360_
            d_1996_v6_: _dafny.Seq
            d_1996_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1993_i4_]), (d_1995_v5_) + (d_1995_v5_), d_1995_v5_, d_1995_v5_, _dafny.SeqWithoutIsStrInference([d_1993_i4_])])
            d_1997_v7_: _dafny.Seq
            d_1997_v7_ = _dafny.SeqWithoutIsStrInference([(self).f32, p1])
            index370_ = default__.safeIndex(659, (p3).length(0))
            index371_ = default__.safeIndex(659, (p3).length(0))
            index372_ = default__.safeIndex(659, (p3).length(0))
            rhs361_ = (d_1996_v6_)[default__.safeIndex(len((_dafny.Map({d_1994_v4_: (self).f25}) if (self).f33 else _dafny.Map({d_1994_v4_: d_1993_i4_}))), len(d_1996_v6_))]
            rhs362_ = d_1994_v4_
            rhs363_ = len(d_1997_v7_)
            rhs364_ = (d_1993_i4_) + ((p3)[default__.safeIndex(659, (p3).length(0))])
            rhs365_ = d_1993_i4_
            lhs275_ = p3
            lhs276_ = default__.safeIndex(659, (p3).length(0))
            lhs277_ = p3
            lhs278_ = default__.safeIndex(659, (p3).length(0))
            lhs279_ = p3
            lhs280_ = default__.safeIndex(659, (p3).length(0))
            d_1995_v5_ = rhs361_
            d_1994_v4_ = rhs362_
            lhs275_[lhs276_] = rhs363_
            lhs277_[lhs278_] = rhs364_
            lhs279_[lhs280_] = rhs365_
        d_1998_v8_: D7
        d_1998_v8_ = D7_DC22((-783) + ((self).f25), (self).f25)
        d_1998_v8_ = d_1998_v8_
        if p2:
            d_1999_v9_: str
            d_1999_v9_ = _dafny.CodePoint('s')
            d_2000_v10_: _dafny.Seq
            d_2000_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxkxcmnk"))
            d_2001_v11_: D1
            d_2001_v11_ = D1_DC6(d_1999_v9_, (self).f32, d_2000_v10_)
            d_2002_v12_: _dafny.Map
            d_2002_v12_ = _dafny.Map({d_2000_v10_: p2})
            d_2003_v13_: _dafny.Map
            d_2003_v13_ = d_2002_v12_
            d_2004_v14_: _dafny.Map
            d_2004_v14_ = _dafny.Map({len((d_2001_v11_).cf9): (d_2003_v13_)})
            d_2005_v15_: _dafny.Map
            d_2005_v15_ = _dafny.Map({len((((d_2004_v14_)[(self).f25] if ((self).f25) in (d_2004_v14_) else d_2002_v12_)) | (_dafny.Map({(D1_DC6(_dafny.CodePoint('o'), (self).f32, d_2000_v10_)).cf9: (self).f32}))): d_2000_v10_})
            d_2006_v16_: _dafny.Seq
            d_2006_v16_ = _dafny.SeqWithoutIsStrInference([d_2005_v15_, _dafny.Map({(self).f25: d_2000_v10_})])
            d_2007_v17_: _dafny.MultiSet
            d_2007_v17_ = _dafny.MultiSet([True])
            d_2005_v15_ = (d_2006_v16_)[default__.safeIndex(((d_2007_v17_).intersection(d_2007_v17_)).cardinality, len(d_2006_v16_))]
            d_2008_v18_: _dafny.Array
            nw322_ = _dafny.Array(None, 19)
            nw322_[int(0)] = (d_2000_v10_)[default__.safeIndex(len(d_2000_v10_), len(d_2000_v10_))]
            nw322_[int(1)] = (default__.fm39(p2, globalState) if not((self).f32) else d_1999_v9_)
            nw322_[int(2)] = d_1999_v9_
            nw322_[int(3)] = d_1999_v9_
            nw322_[int(4)] = d_1999_v9_
            nw322_[int(5)] = d_1999_v9_
            nw322_[int(6)] = d_1999_v9_
            nw322_[int(7)] = (d_1999_v9_ if (self).f33 else _dafny.CodePoint('w'))
            nw322_[int(8)] = d_1999_v9_
            nw322_[int(9)] = default__.fm39(p1, globalState)
            nw322_[int(10)] = d_1999_v9_
            nw322_[int(11)] = default__.fm39((self).f33, globalState)
            nw322_[int(12)] = _dafny.CodePoint('k')
            nw322_[int(13)] = d_1999_v9_
            nw322_[int(14)] = (d_2000_v10_)[default__.safeIndex((self).f25, len(d_2000_v10_))]
            nw322_[int(15)] = d_1999_v9_
            nw322_[int(16)] = d_1999_v9_
            nw322_[int(17)] = d_1999_v9_
            nw322_[int(18)] = d_1999_v9_
            d_2008_v18_ = nw322_
            index373_ = default__.safeIndex(482, (d_2008_v18_).length(0))
            (d_2008_v18_)[index373_] = d_1999_v9_
            d_2009_v19_: D9
            d_2009_v19_ = D9_DC27((0) - ((self).f25), (self).f25)
            d_2010_v20_: _dafny.Seq
            d_2010_v20_ = _dafny.SeqWithoutIsStrInference([d_2009_v19_])
            d_2011_v21_: _dafny.Seq
            d_2011_v21_ = _dafny.SeqWithoutIsStrInference([(d_2010_v20_) + (d_2010_v20_), (d_2010_v20_) + (d_2010_v20_), _dafny.SeqWithoutIsStrInference([d_2009_v19_])])
            index374_ = default__.safeIndex(482, (d_2008_v18_).length(0))
            rhs366_ = d_1999_v9_
            rhs367_ = ((self).f25) - ((self).f25)
            rhs368_ = _dafny.CodePoint('w')
            rhs369_ = (default__.fm2((self).f25, (self).f25, globalState)) + (d_2000_v10_)
            rhs370_ = (d_2011_v21_)[default__.safeIndex((self).f25, len(d_2011_v21_))]
            lhs281_ = d_2008_v18_
            lhs282_ = default__.safeIndex(482, (d_2008_v18_).length(0))
            lhs283_ = globalState
            d_1999_v9_ = rhs366_
            r1 = rhs367_
            lhs281_[lhs282_] = rhs368_
            lhs283_.f10 = rhs369_
            d_2010_v20_ = rhs370_
            (globalState).f12 = (self).f32
            d_2012_v22_: _dafny.Map
            d_2012_v22_ = _dafny.Map({d_2000_v10_: (d_2008_v18_)[default__.safeIndex(482, (d_2008_v18_).length(0))]})
            d_2013_v23_: C1
            nw323_ = C1()
            nw323_.ctor__(len(d_2012_v22_), (self).f26)
            d_2013_v23_ = nw323_
            d_2014_v24_: _dafny.Seq
            d_2014_v24_ = _dafny.SeqWithoutIsStrInference([(self).f33])
            d_2014_v24_ = d_2014_v24_
        elif True:
            d_2015_v25_: D15
            d_2015_v25_ = D15_DC40()
            source27_ = d_2015_v25_
            if source27_.is_DC40:
                d_2016_v26_: T1
                out37_: T1
                out37_ = (self).m8(p1, p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maar")), globalState)
                d_2016_v26_ = out37_
                d_2017_v27_: str
                d_2017_v27_ = _dafny.CodePoint('d')
                d_2018_v28_: _dafny.Seq
                d_2018_v28_ = _dafny.SeqWithoutIsStrInference([d_2017_v27_])
                d_2019_v29_: _dafny.Map
                d_2019_v29_ = _dafny.Map({p1: (self).f25})
                index375_ = default__.safeIndex(607, ((d_2016_v26_).f26).length(0))
                ((d_2016_v26_).f26)[index375_] = (d_2018_v28_) + ((d_2018_v28_).set(default__.safeIndex(len(d_2019_v29_), len(d_2018_v28_)), d_2017_v27_))
                d_2020_v30_: _dafny.MultiSet
                d_2020_v30_ = _dafny.MultiSet([d_2018_v28_])
                d_2021_v31_: _dafny.Seq
                d_2021_v31_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2022_i5_ in range(default__.abs(141))]))])
                d_2023_v32_: D1
                d_2023_v32_ = D1_DC6((d_2018_v28_)[default__.safeIndex(len(d_2021_v31_), len(d_2018_v28_))], (self).f33, d_2018_v28_)
                d_2024_v33_: _dafny.Map
                d_2024_v33_ = _dafny.Map({len(d_2018_v28_): (self).f32})
                d_2025_v34_: _dafny.Seq
                d_2025_v34_ = _dafny.SeqWithoutIsStrInference([True])
                index376_ = default__.safeIndex(607, ((d_2016_v26_).f26).length(0))
                rhs371_ = not(not (((0) - (((d_2020_v30_)[(d_2023_v32_).cf9] if ((d_2023_v32_).cf9) in (d_2020_v30_) else len(d_2018_v28_)))) != ((self).f25)) or ((self).f32))
                rhs372_ = (_dafny.SeqWithoutIsStrInference([d_2017_v27_ for d_2026_i6_ in range(default__.abs(538))]) if (self).f33 else (d_2018_v28_).set(default__.safeIndex(len(d_2019_v29_), len(d_2018_v28_)), default__.fm39(False, globalState)))
                rhs373_ = (((d_2024_v33_)[(self).f25] if ((self).f25) in (d_2024_v33_) else (self).f32) if (self).f33 else (d_2025_v34_)[default__.safeIndex((self).f25, len(d_2025_v34_))])
                lhs284_ = globalState
                lhs285_ = (d_2016_v26_).f26
                lhs286_ = default__.safeIndex(607, ((d_2016_v26_).f26).length(0))
                lhs287_ = globalState
                lhs284_.f12 = rhs371_
                lhs285_[lhs286_] = rhs372_
                lhs287_.f12 = rhs373_
                d_2027_v35_: _dafny.Array
                def lambda130_(d_2028_v29_):
                    def lambda131_(d_2029_i7_):
                        return _dafny.MultiSet([len(d_2028_v29_), (_dafny.MultiSet([(self).f32])).cardinality])

                    return lambda131_

                init73_ = lambda130_(d_2019_v29_)
                nw324_ = _dafny.Array(None, 29)
                for i0_73_ in range(nw324_.length(0)):
                    nw324_[i0_73_] = init73_(i0_73_)
                d_2027_v35_ = nw324_
                d_2027_v35_ = d_2027_v35_
                d_2030_v36_: _dafny.Map
                d_2030_v36_ = _dafny.Map({len(d_2018_v28_): d_2017_v27_})
                (globalState).f12 = ((d_2017_v27_) not in (default__.fm2(len(d_2030_v36_), 425, globalState)) if p2 else not (not(p2)) or (p1))
            elif True:
                d_2031___mcc_h0_ = source27_.cf59
                d_2032_cf59_ = d_2031___mcc_h0_
                r1 = (0) - ((self).f25)
                r0 = (self).f25
                (globalState).f12 = p2
                r0 = ((0) - (default__.fm31(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkfqgj")), globalState)) if not(not((self).f32)) else (self).f25)
            d_2033_v37_: str
            d_2033_v37_ = _dafny.CodePoint('a')
            d_2034_v38_: _dafny.Array
            nw325_ = _dafny.Array(None, 22)
            nw325_[int(0)] = _dafny.CodePoint('f')
            nw325_[int(1)] = d_2033_v37_
            nw325_[int(2)] = d_2033_v37_
            nw325_[int(3)] = d_2033_v37_
            nw325_[int(4)] = d_2033_v37_
            nw325_[int(5)] = _dafny.CodePoint('b')
            nw325_[int(6)] = d_2033_v37_
            nw325_[int(7)] = d_2033_v37_
            nw325_[int(8)] = d_2033_v37_
            nw325_[int(9)] = d_2033_v37_
            nw325_[int(10)] = d_2033_v37_
            nw325_[int(11)] = d_2033_v37_
            nw325_[int(12)] = d_2033_v37_
            nw325_[int(13)] = d_2033_v37_
            nw325_[int(14)] = d_2033_v37_
            nw325_[int(15)] = d_2033_v37_
            nw325_[int(16)] = d_2033_v37_
            nw325_[int(17)] = d_2033_v37_
            nw325_[int(18)] = d_2033_v37_
            nw325_[int(19)] = d_2033_v37_
            nw325_[int(20)] = d_2033_v37_
            nw325_[int(21)] = _dafny.CodePoint('a')
            d_2034_v38_ = nw325_
            d_2035_v39_: _dafny.Map
            d_2035_v39_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")): d_2034_v38_})
            d_2036_v40_: _dafny.Map
            d_2036_v40_ = _dafny.Map({(d_2035_v39_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_2033_v37_ for d_2037_i8_ in range(default__.abs(223))]): d_2034_v38_})): ((self).f25) != ((self).f25)})
            d_2036_v40_ = (d_2036_v40_).set(d_2035_v39_, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wya")))) != ((self).f25))
            (globalState).f12 = (self).fm13(globalState)
            d_2038_v41_: D0
            d_2038_v41_ = D0_DC0(p1)
            d_2039_v42_: _dafny.Seq
            d_2039_v42_ = _dafny.SeqWithoutIsStrInference([d_2038_v41_])
            d_2040_v43_: _dafny.Set
            d_2040_v43_ = _dafny.Set({d_2039_v42_})
            if (d_2040_v43_).issubset((d_2040_v43_) | (d_2040_v43_)):
                d_2041_v44_: _dafny.Array
                nw326_ = _dafny.Array(_dafny.Map({}), 20)
                d_2041_v44_ = nw326_
                d_2042_v45_: _dafny.Map
                d_2042_v45_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f25 for d_2043_i9_ in range(default__.abs(650))])): (self).f25})
                index377_ = default__.safeIndex(906, (d_2041_v44_).length(0))
                (d_2041_v44_)[index377_] = _dafny.Map({len(d_2042_v45_): (self).f33})
                d_2044_v46_: _dafny.Seq
                d_2044_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpmh"))
                d_2045_v47_: _dafny.MultiSet
                d_2045_v47_ = _dafny.MultiSet([(0) - ((self).f25), (self).f25])
                d_2046_v48_: _dafny.Map
                d_2046_v48_ = _dafny.Map({(self).f25: p1})
                index378_ = default__.safeIndex(906, (d_2041_v44_).length(0))
                def iife165_():
                    coll75_ = _dafny.Map()
                    compr_75_: int
                    for compr_75_ in _dafny.IntegerRange(349, -119):
                        d_2047_v49_: int = compr_75_
                        if ((349) <= (d_2047_v49_)) and ((d_2047_v49_) < (-119)):
                            coll75_[(d_2047_v49_) * ((0) - ((self).f25))] = p1
                    return _dafny.Map(coll75_)
                rhs374_ = (self).f33
                rhs375_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdjfim"))) + (d_2044_v46_)
                rhs376_ = (((d_2045_v47_)[len(d_2044_v46_)] if (len(d_2044_v46_)) in (d_2045_v47_) else default__.fm44((self).f33, p1, (self).f32, (self).f25, globalState))) - ((self).f25)
                rhs377_ = (d_2046_v48_) | ((iife165_()
                ) | (d_2046_v48_))
                lhs288_ = globalState
                lhs289_ = globalState
                lhs290_ = d_2041_v44_
                lhs291_ = default__.safeIndex(906, (d_2041_v44_).length(0))
                lhs288_.f12 = rhs374_
                lhs289_.f10 = rhs375_
                r1 = rhs376_
                lhs290_[lhs291_] = rhs377_
                d_2048_v50_: C5
                nw327_ = C5()
                nw327_.ctor__(False)
                d_2048_v50_ = nw327_
                d_2049_v51_: _dafny.Array
                def lambda132_(d_2050_v47_):
                    def lambda133_(d_2051_i10_):
                        return d_2050_v47_

                    return lambda133_

                init74_ = lambda132_(d_2045_v47_)
                nw328_ = _dafny.Array(None, 9)
                for i0_74_ in range(nw328_.length(0)):
                    nw328_[i0_74_] = init74_(i0_74_)
                d_2049_v51_ = nw328_
                index379_ = default__.safeIndex(693, (d_2049_v51_).length(0))
                (d_2049_v51_)[index379_] = d_2045_v47_
                d_2052_v52_: D7
                d_2052_v52_ = D7_DC21(d_1983_v0_)
                d_2053_v53_: _dafny.Array
                nw329_ = _dafny.Array(None, 24)
                nw329_[int(0)] = p3
                nw329_[int(1)] = p3
                nw329_[int(2)] = p3
                nw329_[int(3)] = p3
                nw329_[int(4)] = d_1983_v0_
                nw329_[int(5)] = p3
                nw329_[int(6)] = d_1983_v0_
                nw329_[int(7)] = d_1983_v0_
                nw329_[int(8)] = d_1983_v0_
                nw329_[int(9)] = p3
                nw329_[int(10)] = (d_2052_v52_).cf32
                nw329_[int(11)] = p3
                nw329_[int(12)] = d_1983_v0_
                nw329_[int(13)] = d_1983_v0_
                nw329_[int(14)] = (d_2052_v52_).cf32
                nw329_[int(15)] = p3
                nw329_[int(16)] = p3
                nw329_[int(17)] = p3
                nw329_[int(18)] = p3
                nw329_[int(19)] = d_1983_v0_
                nw329_[int(20)] = d_1983_v0_
                nw329_[int(21)] = p3
                nw329_[int(22)] = p3
                nw329_[int(23)] = d_1983_v0_
                d_2053_v53_ = nw329_
                index380_ = default__.safeIndex(206, (d_2053_v53_).length(0))
                (d_2053_v53_)[index380_] = d_1983_v0_
                d_2054_v54_: _dafny.Seq
                d_2054_v54_ = _dafny.SeqWithoutIsStrInference([d_1991_v2_, (d_1991_v2_).intersection(_dafny.Set({p1, p1})), d_1991_v2_])
                d_2055_v55_: _dafny.Seq
                d_2055_v55_ = _dafny.SeqWithoutIsStrInference([d_1983_v0_, p3, d_1983_v0_, d_1983_v0_, d_1983_v0_])
                index381_ = default__.safeIndex(693, (d_2049_v51_).length(0))
                index382_ = default__.safeIndex(206, (d_2053_v53_).length(0))
                rhs378_ = len(d_2054_v54_)
                rhs379_ = d_2045_v47_
                rhs380_ = (d_2055_v55_)[default__.safeIndex(default__.fm44(p2, p1, p1, (0) - ((0) - ((self).f25)), globalState), len(d_2055_v55_))]
                lhs292_ = d_2049_v51_
                lhs293_ = default__.safeIndex(693, (d_2049_v51_).length(0))
                lhs294_ = d_2053_v53_
                lhs295_ = default__.safeIndex(206, (d_2053_v53_).length(0))
                r1 = rhs378_
                lhs292_[lhs293_] = rhs379_
                lhs294_[lhs295_] = rhs380_
                (globalState).f10 = _dafny.SeqWithoutIsStrInference([d_2033_v37_ for d_2056_i11_ in range(default__.abs(-673))])
                (globalState).f10 = d_2044_v46_
            elif True:
                d_2057_v56_: _dafny.Array
                nw330_ = _dafny.Array(False, 27)
                d_2057_v56_ = nw330_
                d_2058_v57_: _dafny.Map
                d_2058_v57_ = _dafny.Map({(self).f32: (self).f32})
                index383_ = default__.safeIndex(693, (d_2057_v56_).length(0))
                (d_2057_v56_)[index383_] = (D13_DC36(d_2058_v57_, (self).fm13(globalState), (self).f25)).cf55
                index384_ = default__.safeIndex(693, (d_2057_v56_).length(0))
                (d_2057_v56_)[index384_] = p1
                d_2059_v58_: _dafny.Map
                d_2059_v58_ = _dafny.Map({d_2033_v37_: (0) - (((self).f25) - ((self).f25))})
                d_2059_v58_ = (d_2059_v58_).set(d_2033_v37_, (self).f25)
                r0 = (22) + ((self).f25)
                d_2058_v57_ = (d_2058_v57_).set(p1, not((d_2057_v56_)[default__.safeIndex(693, (d_2057_v56_).length(0))]))
                (globalState).f23 = (d_2033_v37_ if (d_2057_v56_)[default__.safeIndex(693, (d_2057_v56_).length(0))] else d_2033_v37_)
            (globalState).f12 = not (p2) or ((self).f32)
        d_2060_v59_: _dafny.Array
        nw331_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_2060_v59_ = nw331_
        index385_ = default__.safeIndex(9, (d_2060_v59_).length(0))
        (d_2060_v59_)[index385_] = p3
        d_2061_v60_: str
        d_2061_v60_ = _dafny.CodePoint('b')
        d_2062_v61_: _dafny.Array
        nw332_ = _dafny.Array(None, 3)
        nw332_[int(0)] = (_dafny.CodePoint('f') if p1 else d_2061_v60_)
        nw332_[int(1)] = d_2061_v60_
        nw332_[int(2)] = _dafny.CodePoint('f')
        d_2062_v61_ = nw332_
        index386_ = default__.safeIndex(622, (d_2062_v61_).length(0))
        (d_2062_v61_)[index386_] = d_2061_v60_
        d_2063_v62_: _dafny.Seq
        d_2063_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpkq"))
        index387_ = default__.safeIndex(9, (d_2060_v59_).length(0))
        index388_ = default__.safeIndex(622, (d_2062_v61_).length(0))
        rhs381_ = p3
        rhs382_ = (d_2063_v62_)[default__.safeIndex(len(d_2063_v62_), len(d_2063_v62_))]
        lhs296_ = d_2060_v59_
        lhs297_ = default__.safeIndex(9, (d_2060_v59_).length(0))
        lhs298_ = d_2062_v61_
        lhs299_ = default__.safeIndex(622, (d_2062_v61_).length(0))
        lhs296_[lhs297_] = rhs381_
        lhs298_[lhs299_] = rhs382_
        r0 = 457
        d_2064_v63_: _dafny.Map
        d_2064_v63_ = _dafny.Map({(self).f32: (self).f25})
        r1 = (len(d_2064_v63_)) - ((len(d_2064_v63_)) + ((0) - ((0) - ((self).f25))))
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        (globalState).f12 = True
        d_2065_v1_: str
        d_2065_v1_ = _dafny.CodePoint('r')
        d_2066_v2_: _dafny.MultiSet
        d_2066_v2_ = _dafny.MultiSet([d_2065_v1_])
        d_2067_v3_: _dafny.Map
        d_2067_v3_ = _dafny.Map({(d_2066_v2_).cardinality: not(False)})
        d_2068_v4_: D7
        d_2068_v4_ = D7_DC22((self).f25, len(d_2067_v3_))
        d_2069_v5_: _dafny.Seq
        d_2069_v5_ = _dafny.SeqWithoutIsStrInference([(self).f32, not((self).f32), p1])
        d_2070_v6_: _dafny.Map
        d_2070_v6_ = _dafny.Map({d_2068_v4_: d_2069_v5_})
        d_2071_v7_: C4
        nw333_ = C4()
        def iife166_():
            coll76_ = _dafny.Map()
            compr_76_: D7
            for compr_76_ in (d_2070_v6_).keys.Elements:
                d_2072_v0_: D7 = compr_76_
                if (d_2072_v0_) in (d_2070_v6_):
                    coll76_[d_2072_v0_] = (self).f32
            return _dafny.Map(coll76_)
        nw333_.ctor__((self).f26, len((iife166_()
        ) | (_dafny.Map({d_2068_v4_: (self).f33}))), (self).f26)
        d_2071_v7_ = nw333_
        d_2073_v8_: _dafny.MultiSet
        d_2073_v8_ = _dafny.MultiSet([(self).f25])
        (globalState).f12 = (d_2073_v8_).ispropersubset(d_2073_v8_)
        d_2067_v3_ = (d_2067_v3_).set((self).f25, True)
        d_2074_v9_: _dafny.Seq
        d_2074_v9_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_2075_v10_: _dafny.Array
        nw334_ = _dafny.Array(None, 5)
        nw334_[int(0)] = (d_2069_v5_)[default__.safeIndex(515, len(d_2069_v5_))]
        nw334_[int(1)] = (self).f33
        nw334_[int(2)] = (self).f32
        nw334_[int(3)] = (d_2074_v9_) != (d_2074_v9_)
        nw334_[int(4)] = (self).f33
        d_2075_v10_ = nw334_
        index389_ = default__.safeIndex(180, (d_2075_v10_).length(0))
        (d_2075_v10_)[index389_] = not(default__.fm0((0) - (len((_dafny.Map({d_2065_v1_: (self).f25})).set(d_2065_v1_, (self).f25))), globalState))
        index390_ = default__.safeIndex(180, (d_2075_v10_).length(0))
        (d_2075_v10_)[index390_] = (self).f33
        (globalState).f6 = (d_2074_v9_) + (d_2074_v9_)
        d_2076_v11_: C2
        nw335_ = C2()
        nw335_.ctor__()
        d_2076_v11_ = nw335_
        d_2077_v12_: _dafny.MultiSet
        d_2077_v12_ = _dafny.MultiSet([d_2076_v11_])
        d_2078_v13_: _dafny.Map
        d_2078_v13_ = _dafny.Map({(d_2077_v12_).cardinality: (self).f25})
        d_2079_v14_: _dafny.Set
        d_2079_v14_ = _dafny.Set({(self).f25})
        r0 = ((self).f25) * (((d_2074_v9_)[default__.safeIndex((self).f25, len(d_2074_v9_))]) * (((d_2078_v13_)[(self).f25] if ((self).f25) in (d_2078_v13_) else len(d_2079_v14_))))
        return r0

    def m8(self, p0, p1, p2, globalState):
        r0: T1 = None
        d_2080_v0_: int
        d_2080_v0_ = 668
        d_2080_v0_ = (self).f25
        (globalState).f12 = (self).f33
        d_2081_v1_: _dafny.Seq
        d_2081_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        d_2082_v2_: str
        d_2082_v2_ = _dafny.CodePoint('l')
        d_2083_v3_: _dafny.MultiSet
        d_2083_v3_ = _dafny.MultiSet([(250 if p0 else len(p2)), d_2080_v0_, d_2080_v0_, default__.safeModuloInt(len((d_2081_v1_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gs"))).set(default__.safeIndex((self).f25, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gs")))), d_2082_v2_)), len(d_2081_v1_))]), d_2080_v0_), (self).f25])
        d_2083_v3_ = d_2083_v3_
        d_2084_v4_: _dafny.Array
        nw336_ = _dafny.Array(False, 8)
        d_2084_v4_ = nw336_
        d_2084_v4_ = d_2084_v4_
        d_2085_v5_: _dafny.Set
        d_2085_v5_ = _dafny.Set({d_2082_v2_, d_2082_v2_})
        d_2086_i0_: int
        d_2086_i0_ = 0
        with _dafny.label("15"):
            while (d_2085_v5_).isdisjoint(d_2085_v5_):
                with _dafny.c_label("15"):
                    if (d_2086_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2086_i0_ = (d_2086_i0_) + (1)
                    d_2087_v6_: _dafny.Seq
                    d_2087_v6_ = _dafny.SeqWithoutIsStrInference([not(p0), not((self).f32)])
                    d_2088_v7_: _dafny.Set
                    d_2088_v7_ = _dafny.Set({(self).f32})
                    rhs383_ = (self).f25
                    rhs384_ = p0
                    rhs385_ = default__.safeModuloInt(default__.safeDivisionInt((self).f25, (self).f25), default__.safeModuloInt(d_2080_v0_, len(d_2087_v6_)))
                    rhs386_ = default__.safeModuloInt(314, (0) - (len((_dafny.Map({d_2088_v7_: p2})) | (_dafny.Map({d_2088_v7_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwy"))})))))
                    lhs300_ = globalState
                    d_2080_v0_ = rhs383_
                    lhs300_.f12 = rhs384_
                    d_2080_v0_ = rhs385_
                    d_2080_v0_ = rhs386_
                    d_2089_v8_: _dafny.MultiSet
                    d_2089_v8_ = _dafny.MultiSet([(self).f32, p0])
                    (globalState).f12 = (d_2089_v8_).isdisjoint((d_2089_v8_).intersection(d_2089_v8_))
                    if (self).f33:
                        index391_ = default__.safeIndex(677, (p1).length(0))
                        (p1)[index391_] = (0) - ((self).f25)
                        index392_ = default__.safeIndex(677, (p1).length(0))
                        (p1)[index392_] = (self).f25
                        d_2090_v9_: _dafny.Seq
                        d_2090_v9_ = _dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(677, (p1).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfla")))])
                        d_2091_v10_: _dafny.Set
                        d_2091_v10_ = _dafny.Set({d_2083_v3_, (d_2083_v3_).intersection((d_2083_v3_).set((p1)[default__.safeIndex(677, (p1).length(0))], default__.abs(len(d_2090_v9_))))})
                        d_2092_v11_: D18
                        d_2092_v11_ = D18_DC43(d_2091_v10_)
                        d_2091_v10_ = (d_2092_v11_).cf62
                        d_2093_v12_: C3
                        nw337_ = C3()
                        nw337_.ctor__((p1)[default__.safeIndex(677, (p1).length(0))], p0, (0) - ((self).f25), (self).f26)
                        d_2093_v12_ = nw337_
                        d_2094_v13_: _dafny.Map
                        d_2094_v13_ = _dafny.Map({d_2084_v4_: (d_2093_v12_).f38})
                        d_2094_v13_ = (d_2094_v13_).set(d_2084_v4_, (self).f33)
                        d_2095_v14_: C2
                        nw338_ = C2()
                        nw338_.ctor__()
                        d_2095_v14_ = nw338_
                    elif True:
                        d_2096_v15_: _dafny.Set
                        d_2096_v15_ = _dafny.Set({d_2080_v0_, (self).f25})
                        d_2097_v16_: _dafny.Map
                        d_2097_v16_ = _dafny.Map({(_dafny.Set({433})) | (d_2096_v15_): d_2087_v6_})
                        d_2097_v16_ = _dafny.Map({d_2096_v15_: d_2087_v6_})
                        d_2098_v17_: _dafny.Array
                        nw339_ = _dafny.Array(None, 20)
                        nw339_[int(0)] = d_2084_v4_
                        nw339_[int(1)] = d_2084_v4_
                        nw339_[int(2)] = d_2084_v4_
                        nw339_[int(3)] = d_2084_v4_
                        nw339_[int(4)] = d_2084_v4_
                        nw339_[int(5)] = d_2084_v4_
                        nw339_[int(6)] = d_2084_v4_
                        nw339_[int(7)] = d_2084_v4_
                        nw339_[int(8)] = d_2084_v4_
                        nw339_[int(9)] = d_2084_v4_
                        nw339_[int(10)] = d_2084_v4_
                        nw339_[int(11)] = d_2084_v4_
                        nw339_[int(12)] = d_2084_v4_
                        nw339_[int(13)] = d_2084_v4_
                        nw339_[int(14)] = d_2084_v4_
                        nw339_[int(15)] = d_2084_v4_
                        nw339_[int(16)] = (d_2084_v4_ if (self).f33 else d_2084_v4_)
                        nw339_[int(17)] = d_2084_v4_
                        nw339_[int(18)] = d_2084_v4_
                        nw339_[int(19)] = d_2084_v4_
                        d_2098_v17_ = nw339_
                        d_2098_v17_ = d_2098_v17_
                        d_2099_v18_: _dafny.Seq
                        d_2099_v18_ = _dafny.SeqWithoutIsStrInference([d_2080_v0_])
                        d_2100_v19_: _dafny.Map
                        d_2100_v19_ = _dafny.Map({d_2082_v2_: default__.fm0(len(d_2099_v18_), globalState)})
                        d_2101_v20_: D10
                        d_2101_v20_ = D10_DC28(d_2100_v19_)
                        d_2101_v20_ = (d_2101_v20_ if False else D10_DC28(d_2100_v19_))
                        (globalState).f12 = (p2) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amuiefuy")))
                        index393_ = default__.safeIndex(323, (p1).length(0))
                        (p1)[index393_] = d_2080_v0_
                        index394_ = default__.safeIndex(323, (p1).length(0))
                        (p1)[index394_] = (0) - (default__.fm44(p0, (self).f32, default__.fm0((d_2083_v3_).cardinality, globalState), d_2080_v0_, globalState))
                    if (self).f32:
                        (globalState).f12 = (self).f32
                        index395_ = default__.safeIndex(411, (d_2084_v4_).length(0))
                        (d_2084_v4_)[index395_] = default__.fm0(d_2080_v0_, globalState)
                        index396_ = default__.safeIndex(411, (d_2084_v4_).length(0))
                        rhs387_ = (self).f33
                        rhs388_ = (self).f25
                        lhs301_ = d_2084_v4_
                        lhs302_ = default__.safeIndex(411, (d_2084_v4_).length(0))
                        lhs301_[lhs302_] = rhs387_
                        d_2080_v0_ = rhs388_
                        d_2102_v21_: _dafny.Map
                        d_2102_v21_ = _dafny.Map({((_dafny.SeqWithoutIsStrInference([d_2082_v2_ for d_2103_i1_ in range(default__.abs(30))])) + (p2)).set(default__.safeIndex(len(d_2081_v1_), len((_dafny.SeqWithoutIsStrInference([d_2082_v2_ for d_2104_i1_ in range(default__.abs(30))])) + (p2))), d_2082_v2_): (d_2087_v6_) + (_dafny.SeqWithoutIsStrInference([not(p0)]))})
                        d_2087_v6_ = ((d_2102_v21_)[p2] if (p2) in (d_2102_v21_) else ((d_2087_v6_) + (d_2087_v6_)).set(default__.safeIndex(len(d_2088_v7_), len((d_2087_v6_) + (d_2087_v6_))), True))
                        (globalState).f12 = (self).f32
                        d_2105_v22_: C6
                        nw340_ = C6()
                        nw340_.ctor__((d_2080_v0_) * ((self).f25), (self).f26)
                        d_2105_v22_ = nw340_
                    elif True:
                        d_2106_v23_: D0
                        d_2106_v23_ = D0_DC3((self).f25, (self).f33)
                        d_2107_v24_: _dafny.Array
                        def lambda134_(d_2108_v23_, d_2109_v8_, d_2110_v0_, d_2111_p2_):
                            def lambda135_(d_2112_i2_):
                                return D6_DC19(d_2108_v23_, ((d_2109_v8_)[(self).f33] if ((self).f33) in (d_2109_v8_) else (D18_DC44(d_2110_v0_, False, d_2111_p2_, (self).f33, _dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boihs"))), -626})]))).cf63))

                            return lambda135_

                        init75_ = lambda134_(d_2106_v23_, d_2089_v8_, d_2080_v0_, p2)
                        nw341_ = _dafny.Array(None, 20)
                        for i0_75_ in range(nw341_.length(0)):
                            nw341_[i0_75_] = init75_(i0_75_)
                        d_2107_v24_ = nw341_
                        d_2113_v25_: _dafny.Map
                        d_2113_v25_ = _dafny.Map({D6_DC19(d_2106_v23_, (self).f25): d_2107_v24_})
                        d_2114_v26_: D12
                        d_2114_v26_ = D12_DC31(d_2113_v25_)
                        pat_let_tv39_ = d_2113_v25_
                        def iife167_(_pat_let45_0):
                            def iife168_(d_2115_dt__update__tmp_h0_):
                                def iife169_(_pat_let46_0):
                                    def iife170_(d_2116_dt__update_hcf46_h0_):
                                        return D12_DC31(d_2116_dt__update_hcf46_h0_)
                                    return iife170_(_pat_let46_0)
                                return iife169_(pat_let_tv39_)
                            return iife168_(_pat_let45_0)
                        d_2114_v26_ = iife167_(d_2114_v26_)
                        (globalState).f12 = p0
                        (globalState).f10 = default__.fm2(d_2080_v0_, 99, globalState)
                        d_2117_v27_: _dafny.Map
                        d_2117_v27_ = _dafny.Map({(self).f25: default__.safeDivisionInt(d_2080_v0_, d_2080_v0_)})
                        d_2117_v27_ = d_2117_v27_
                        d_2080_v0_ = (((self).f25) + (((d_2089_v8_)[(self).f32] if ((self).f32) in (d_2089_v8_) else d_2080_v0_))) - (len(p2))
                    pass
            pass
        d_2118_v28_: _dafny.Map
        d_2118_v28_ = _dafny.Map({(self).f33: (d_2080_v0_) * ((self).f25)})
        d_2118_v28_ = _dafny.Map({(self).fm13(globalState): d_2080_v0_})
        d_2119_v29_: T1
        nw342_ = C1()
        nw342_.ctor__((self).f25, (self).f26)
        d_2119_v29_ = nw342_
        r0 = d_2119_v29_
        return r0

    def m9(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_2120_v0_: _dafny.Seq
        d_2120_v0_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f33])
        d_2121_v1_: _dafny.Array
        nw343_ = _dafny.Array(None, 4)
        nw343_[int(0)] = (d_2120_v0_ if not((self).f32) else default__.fm1(not((self).f32), globalState))
        nw343_[int(1)] = (d_2120_v0_) + (d_2120_v0_)
        nw343_[int(2)] = d_2120_v0_
        nw343_[int(3)] = d_2120_v0_
        d_2121_v1_ = nw343_
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, (d_2121_v1_).length(0)):
            d_2122_i0_: int = guard_loop_11_
            if (True) and (((0) <= (d_2122_i0_)) and ((d_2122_i0_) < ((d_2121_v1_).length(0)))):
                (d_2121_v1_)[(d_2122_i0_)] = (d_2120_v0_) + (d_2120_v0_)
        d_2123_v2_: C3
        nw344_ = C3()
        nw344_.ctor__((self).f25, ((self).f33) == (False), p0, (self).f26)
        d_2123_v2_ = nw344_
        if ((self).f25) == ((self).f25):
            if (p0) != ((0) - ((0) - ((d_2123_v2_).f37))):
                (globalState).f12 = (d_2123_v2_).f38
                d_2124_v3_: C0
                nw345_ = C0()
                nw345_.ctor__()
                d_2124_v3_ = nw345_
                d_2125_v4_: _dafny.Array
                def lambda136_(d_2126_i1_):
                    return (self).f32

                init76_ = lambda136_
                nw346_ = _dafny.Array(None, 7)
                for i0_76_ in range(nw346_.length(0)):
                    nw346_[i0_76_] = init76_(i0_76_)
                d_2125_v4_ = nw346_
                d_2127_v5_: _dafny.Set
                d_2127_v5_ = _dafny.Set({d_2125_v4_, d_2125_v4_})
                d_2128_v6_: D3
                d_2128_v6_ = D3_DC9(d_2127_v5_)
                d_2129_v7_: _dafny.Map
                d_2129_v7_ = _dafny.Map({(d_2123_v2_).f37: d_2128_v6_})
                d_2130_v8_: D19
                d_2130_v8_ = D19_DC46(d_2129_v7_)
                d_2131_v9_: _dafny.Array
                nw347_ = _dafny.Array(int(0), 9)
                d_2131_v9_ = nw347_
                d_2132_v10_: _dafny.Map
                d_2132_v10_ = _dafny.Map({(d_2130_v8_).cf68: d_2131_v9_})
                d_2133_v11_: _dafny.Seq
                d_2133_v11_ = _dafny.SeqWithoutIsStrInference([d_2132_v10_, d_2132_v10_, d_2132_v10_])
                rhs389_ = (d_2132_v10_) | ((d_2133_v11_)[default__.safeIndex((d_2123_v2_).f37, len(d_2133_v11_))])
                rhs390_ = (self).f33
                d_2132_v10_ = rhs389_
                r0 = rhs390_
                d_2134_v12_: _dafny.Map
                d_2134_v12_ = _dafny.Map({d_2125_v4_: (self).f33})
                (globalState).f12 = (d_2134_v12_) != ((d_2134_v12_) | (d_2134_v12_))
                d_2135_v13_: _dafny.Map
                d_2135_v13_ = _dafny.Map({(d_2123_v2_).f37: (self).f32})
                d_2136_v14_: _dafny.Map
                d_2136_v14_ = d_2135_v13_
                d_2136_v14_ = default__.fm65(((d_2123_v2_).f37) + ((d_2123_v2_).f37), globalState)
            elif True:
                r2 = (self).f32
                d_2137_v15_: _dafny.Array
                nw348_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_2137_v15_ = nw348_
                d_2138_v16_: _dafny.Seq
                d_2138_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')])
                d_2139_v17_: _dafny.Set
                d_2139_v17_ = _dafny.Set({len(d_2138_v16_)})
                d_2140_v18_: _dafny.Map
                d_2140_v18_ = _dafny.Map({(self).f33: (d_2123_v2_).f38})
                d_2141_v19_: D12
                d_2141_v19_ = D12_DC32(((d_2140_v18_)[(d_2123_v2_).f38] if ((d_2123_v2_).f38) in (d_2140_v18_) else (d_2123_v2_).f38))
                d_2142_v20_: _dafny.Array
                nw349_ = _dafny.Array(int(0), 25)
                d_2142_v20_ = nw349_
                pat_let_tv40_ = d_2139_v17_
                d_2143_v21_: int
                d_2144_v22_: int
                out38_: int
                out39_: int
                def iife171_(_pat_let47_0):
                    def iife172_(d_2145_dt__update__tmp_h0_):
                        def iife173_(_pat_let48_0):
                            def iife174_(d_2146_dt__update_hcf18_h0_):
                                return D4_DC11(d_2146_dt__update_hcf18_h0_)
                            return iife174_(_pat_let48_0)
                        return iife173_(pat_let_tv40_)
                    return iife172_(_pat_let47_0)
                out38_, out39_ = (d_2123_v2_).m2(d_2137_v15_, (d_2139_v17_).isdisjoint((iife171_(D4_DC11(d_2139_v17_))).cf18), (d_2141_v19_).cf47, d_2142_v20_, globalState)
                d_2143_v21_ = out38_
                d_2144_v22_ = out39_
                rhs391_ = (self).f33
                rhs392_ = (d_2123_v2_).f38
                rhs393_ = default__.safeDivisionInt((691) - ((d_2123_v2_).f37), (d_2123_v2_).f37)
                lhs303_ = globalState
                r0 = rhs391_
                lhs303_.f12 = rhs392_
                d_2143_v21_ = rhs393_
                d_2147_v23_: D19
                d_2147_v23_ = D19_DC48(default__.safeModuloInt((self).f25, 789))
                d_2147_v23_ = default__.fm66(globalState)
                d_2148_v24_: _dafny.Array
                def lambda137_(d_2149_v2_):
                    def lambda138_(d_2150_i2_):
                        return (d_2149_v2_).f38

                    return lambda138_

                init77_ = lambda137_(d_2123_v2_)
                nw350_ = _dafny.Array(None, 13)
                for i0_77_ in range(nw350_.length(0)):
                    nw350_[i0_77_] = init77_(i0_77_)
                d_2148_v24_ = nw350_
                d_2151_v25_: D12
                d_2151_v25_ = D12_DC33(d_2148_v24_, d_2148_v24_, (self).f33, (d_2123_v2_).f37)
                d_2152_v26_: D19
                d_2152_v26_ = D19_DC47(d_2148_v24_, (self).f32)
                d_2153_v27_: _dafny.MultiSet
                d_2153_v27_ = _dafny.MultiSet([(d_2123_v2_).f38])
                d_2154_v28_: _dafny.Array
                nw351_ = _dafny.Array(None, 16)
                nw351_[int(0)] = (d_2123_v2_).f38
                nw351_[int(1)] = True
                nw351_[int(2)] = (d_2123_v2_).f38
                nw351_[int(3)] = (d_2123_v2_).f38
                nw351_[int(4)] = not(not((d_2123_v2_).f38))
                nw351_[int(5)] = ((self).f33 if (d_2123_v2_).f38 else (d_2151_v25_).cf50)
                nw351_[int(6)] = ((d_2140_v18_)[(d_2152_v26_).cf70] if ((d_2152_v26_).cf70) in (d_2140_v18_) else (d_2123_v2_).f38)
                nw351_[int(7)] = (d_2123_v2_).f38
                nw351_[int(8)] = (self).f33
                nw351_[int(9)] = (d_2153_v27_).ispropersubset(d_2153_v27_)
                nw351_[int(10)] = (d_2123_v2_).f38
                nw351_[int(11)] = (self).f33
                nw351_[int(12)] = (d_2144_v22_) < (len(d_2120_v0_))
                nw351_[int(13)] = (d_2123_v2_).f38
                nw351_[int(14)] = (d_2123_v2_).f38
                nw351_[int(15)] = (d_2123_v2_).f38
                d_2154_v28_ = nw351_
                index397_ = default__.safeIndex(221, (d_2154_v28_).length(0))
                (d_2154_v28_)[index397_] = (d_2123_v2_).f38
                index398_ = default__.safeIndex(221, (d_2154_v28_).length(0))
                (d_2154_v28_)[index398_] = (self).f32
            d_2155_v29_: str
            d_2155_v29_ = _dafny.CodePoint('o')
            d_2156_v30_: _dafny.Seq
            d_2156_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gypjrxbv"))
            d_2157_v31_: C5
            nw352_ = C5()
            nw352_.ctor__((d_2155_v29_) in (d_2156_v30_))
            d_2157_v31_ = nw352_
            d_2158_v32_: C2
            nw353_ = C2()
            nw353_.ctor__()
            d_2158_v32_ = nw353_
            d_2159_v33_: _dafny.Map
            d_2159_v33_ = _dafny.Map({(d_2156_v30_) != (_dafny.SeqWithoutIsStrInference([d_2155_v29_ for d_2160_i3_ in range(default__.abs(259))])): d_2156_v30_})
            d_2159_v33_ = (d_2159_v33_).set((d_2157_v31_).f35, d_2156_v30_)
            d_2161_v34_: C2
            nw354_ = C2()
            nw354_.ctor__()
            d_2161_v34_ = nw354_
        elif True:
            d_2162_v35_: int
            d_2162_v35_ = 774
            d_2162_v35_ = (0) - (d_2162_v35_)
            d_2163_v36_: _dafny.Array
            nw355_ = _dafny.Array(int(0), 28)
            d_2163_v36_ = nw355_
            index399_ = default__.safeIndex(216, (d_2163_v36_).length(0))
            (d_2163_v36_)[index399_] = (d_2123_v2_).f37
            index400_ = default__.safeIndex(216, (d_2163_v36_).length(0))
            (d_2163_v36_)[index400_] = d_2162_v35_
            d_2164_v37_: _dafny.Map
            d_2164_v37_ = _dafny.Map({(self).f25: (self).f32})
            d_2165_v38_: _dafny.MultiSet
            d_2165_v38_ = _dafny.MultiSet([True, default__.fm0((d_2123_v2_).f37, globalState), (d_2123_v2_).f38])
            d_2166_v39_: _dafny.Set
            d_2166_v39_ = _dafny.Set({D14_DC38(p0)})
            d_2167_v40_: _dafny.Map
            d_2167_v40_ = _dafny.Map({D14_DC38(p0): (d_2123_v2_).f37})
            def iife175_():
                coll77_ = _dafny.Set()
                compr_77_: D14
                for compr_77_ in (d_2167_v40_).keys.Elements:
                    d_2168_v41_: D14 = compr_77_
                    if (d_2168_v41_) in (d_2167_v40_):
                        coll77_ = coll77_.union(_dafny.Set([d_2168_v41_]))
                return _dafny.Set(coll77_)
            d_2164_v37_ = (d_2164_v37_).set((d_2165_v38_).cardinality, (iife175_()
            ).ispropersubset(d_2166_v39_))
            d_2169_v42_: str
            d_2169_v42_ = _dafny.CodePoint('m')
            (globalState).f23 = d_2169_v42_
            d_2170_v43_: _dafny.Array
            nw356_ = _dafny.Array(None, 6)
            nw356_[int(0)] = (d_2123_v2_).f38
            nw356_[int(1)] = not((d_2123_v2_).f38)
            nw356_[int(2)] = (d_2123_v2_).f38
            nw356_[int(3)] = (d_2123_v2_).f38
            nw356_[int(4)] = (d_2123_v2_).f38
            nw356_[int(5)] = (d_2123_v2_).f38
            d_2170_v43_ = nw356_
            d_2171_v44_: _dafny.Seq
            d_2171_v44_ = _dafny.SeqWithoutIsStrInference([d_2170_v43_])
            d_2172_v45_: D19
            d_2172_v45_ = D19_DC47((d_2171_v44_)[default__.safeIndex(d_2162_v35_, len(d_2171_v44_))], (self).f32)
            pat_let_tv41_ = d_2170_v43_
            def iife176_(_pat_let49_0):
                def iife177_(d_2173_dt__update__tmp_h1_):
                    def iife178_(_pat_let50_0):
                        def iife179_(d_2174_dt__update_hcf69_h0_):
                            return D19_DC47(d_2174_dt__update_hcf69_h0_, (d_2173_dt__update__tmp_h1_).cf70)
                        return iife179_(_pat_let50_0)
                    return iife178_(pat_let_tv41_)
                return iife177_(_pat_let49_0)
            d_2172_v45_ = iife176_(d_2172_v45_)
        d_2175_v46_: int
        d_2175_v46_ = 323
        d_2176_v47_: D10
        d_2176_v47_ = D10_DC29((d_2123_v2_).f38, True, (d_2123_v2_).f37)
        d_2177_v48_: _dafny.MultiSet
        d_2177_v48_ = _dafny.MultiSet([(self).f25, d_2175_v46_])
        d_2178_v49_: _dafny.Map
        d_2178_v49_ = _dafny.Map({d_2177_v48_: _dafny.Set({(self).f32})})
        pat_let_tv42_ = d_2123_v2_
        d_2179_v50_: _dafny.Set
        def iife180_(_pat_let51_0):
            def iife181_(d_2180_dt__update__tmp_h2_):
                def iife182_(_pat_let52_0):
                    def iife183_(d_2181_dt__update_hcf42_h0_):
                        def iife184_(_pat_let53_0):
                            def iife185_(d_2182_dt__update_hcf44_h0_):
                                return D10_DC29(d_2181_dt__update_hcf42_h0_, (d_2180_dt__update__tmp_h2_).cf43, d_2182_dt__update_hcf44_h0_)
                            return iife185_(_pat_let53_0)
                        return iife184_((pat_let_tv42_).f37)
                    return iife183_(_pat_let52_0)
                return iife182_(True)
            return iife181_(_pat_let51_0)
        d_2179_v50_ = _dafny.Set({p0, (iife180_(d_2176_v47_)).cf44, len(default__.fm67((self).f32, d_2175_v46_, d_2178_v49_, globalState)), default__.safeModuloInt((self).f25, p0)})
        d_2175_v46_ = len(d_2179_v50_)
        d_2183_v51_: _dafny.Array
        def lambda139_(d_2184_i4_):
            return (self).f32

        init78_ = lambda139_
        nw357_ = _dafny.Array(None, 14)
        for i0_78_ in range(nw357_.length(0)):
            nw357_[i0_78_] = init78_(i0_78_)
        d_2183_v51_ = nw357_
        d_2185_v52_: _dafny.Map
        d_2185_v52_ = _dafny.Map({True: -578})
        d_2186_v53_: _dafny.Seq
        d_2186_v53_ = _dafny.SeqWithoutIsStrInference([len(d_2185_v52_), (d_2123_v2_).f37, p0])
        d_2187_v54_: D3
        d_2187_v54_ = D3_DC10((self).f32, d_2183_v51_, False, (d_2123_v2_).f37, len(d_2186_v53_))
        d_2188_v55_: D12
        d_2188_v55_ = D12_DC33(d_2183_v51_, (d_2187_v54_).cf14, (d_2123_v2_).f38, (d_2123_v2_).f37)
        source28_ = d_2188_v55_
        if source28_.is_DC32:
            d_2189___mcc_h0_ = source28_.cf47
            d_2190_cf47_ = d_2189___mcc_h0_
            d_2191_v57_: str
            d_2191_v57_ = _dafny.CodePoint('t')
            d_2192_v58_: _dafny.Seq
            d_2192_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kflkte"))
            d_2193_v59_: _dafny.Array
            nw358_ = _dafny.Array(None, 19)
            nw358_[int(0)] = (self).f25
            nw358_[int(1)] = (self).f25
            nw358_[int(2)] = p0
            def iife186_():
                coll78_ = _dafny.Map()
                compr_78_: str
                for compr_78_ in (_dafny.Map({d_2191_v57_: d_2175_v46_})).keys.Elements:
                    d_2194_v56_: str = compr_78_
                    if (d_2194_v56_) in (_dafny.Map({d_2191_v57_: d_2175_v46_})):
                        coll78_[d_2194_v56_] = len(_dafny.Map({len((d_2186_v53_)): (d_2123_v2_).f37}))
                return _dafny.Map(coll78_)
            nw358_[int(3)] = len(iife186_()
            )
            nw358_[int(4)] = (self).f25
            nw358_[int(5)] = (d_2123_v2_).f37
            nw358_[int(6)] = p0
            nw358_[int(7)] = (d_2123_v2_).f37
            nw358_[int(8)] = (d_2123_v2_).f37
            nw358_[int(9)] = (d_2123_v2_).f37
            nw358_[int(10)] = len(d_2192_v58_)
            nw358_[int(11)] = (D4_DC13((d_2123_v2_).f37, False, (self).f33, (self).f33)).cf19
            nw358_[int(12)] = (self).f25
            nw358_[int(13)] = (self).f25
            nw358_[int(14)] = (d_2123_v2_).f37
            nw358_[int(15)] = (d_2123_v2_).f37
            nw358_[int(16)] = (d_2188_v55_).cf51
            nw358_[int(17)] = (self).f25
            nw358_[int(18)] = (self).f25
            d_2193_v59_ = nw358_
            d_2195_v60_: _dafny.Map
            d_2195_v60_ = _dafny.Map({d_2193_v59_: (self).f25})
            d_2196_v61_: _dafny.Array
            nw359_ = _dafny.Array(None, 17)
            nw359_[int(0)] = (0) - (d_2175_v46_)
            nw359_[int(1)] = len(d_2120_v0_)
            nw359_[int(2)] = (0) - ((d_2123_v2_).f37)
            nw359_[int(3)] = (0) - ((d_2123_v2_).f37)
            nw359_[int(4)] = p0
            nw359_[int(5)] = len(d_2195_v60_)
            nw359_[int(6)] = (self).f25
            nw359_[int(7)] = 978
            nw359_[int(8)] = (self).f25
            nw359_[int(9)] = len((d_2192_v58_) + (d_2192_v58_))
            nw359_[int(10)] = (self).f25
            nw359_[int(11)] = (self).f25
            nw359_[int(12)] = d_2175_v46_
            nw359_[int(13)] = (d_2123_v2_).f37
            nw359_[int(14)] = (d_2175_v46_) - (len(d_2192_v58_))
            nw359_[int(15)] = 730
            nw359_[int(16)] = (d_2123_v2_).f37
            d_2196_v61_ = nw359_
            index401_ = default__.safeIndex(275, (d_2193_v59_).length(0))
            (d_2193_v59_)[index401_] = len(d_2192_v58_)
            index402_ = default__.safeIndex(275, (d_2193_v59_).length(0))
            (d_2193_v59_)[index402_] = (d_2123_v2_).f37
            index403_ = default__.safeIndex(275, (d_2193_v59_).length(0))
            (d_2193_v59_)[index403_] = (d_2123_v2_).f37
            d_2197_v62_: _dafny.MultiSet
            d_2197_v62_ = _dafny.MultiSet([not(True)])
            r2 = not((d_2197_v62_).isdisjoint(d_2197_v62_))
            index404_ = default__.safeIndex(275, (d_2193_v59_).length(0))
            (d_2193_v59_)[index404_] = len(d_2192_v58_)
        elif source28_.is_DC33:
            d_2198___mcc_h1_ = source28_.cf48
            d_2199___mcc_h2_ = source28_.cf49
            d_2200___mcc_h3_ = source28_.cf50
            d_2201___mcc_h4_ = source28_.cf51
            d_2202_cf51_ = d_2201___mcc_h4_
            d_2203_cf50_ = d_2200___mcc_h3_
            d_2204_cf49_ = d_2199___mcc_h2_
            d_2205_cf48_ = d_2198___mcc_h1_
            d_2202_cf51_ = (default__.safeDivisionInt((self).f25, len(d_2120_v0_))) + ((self).f25)
            d_2206_v63_: _dafny.Map
            d_2206_v63_ = _dafny.Map({(d_2123_v2_).f37: (self).f25})
            d_2207_v64_: _dafny.Seq
            d_2207_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrmgmmsn"))
            d_2208_v65_: _dafny.Map
            d_2208_v65_ = _dafny.Map({d_2207_v64_: True})
            d_2209_v66_: _dafny.Array
            nw360_ = _dafny.Array(None, 4)
            nw360_[int(0)] = len((_dafny.Map({len(d_2206_v63_): d_2207_v64_})).set(len(d_2208_v65_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtftgurm"))))
            nw360_[int(1)] = d_2202_cf51_
            nw360_[int(2)] = default__.fm44((self).f32, not((self).f33), (self).f33, (0) - ((0) - ((d_2123_v2_).f37)), globalState)
            nw360_[int(3)] = default__.safeModuloInt((self).f25, p0)
            d_2209_v66_ = nw360_
            index405_ = default__.safeIndex(438, (d_2209_v66_).length(0))
            (d_2209_v66_)[index405_] = 788
            index406_ = default__.safeIndex(438, (d_2209_v66_).length(0))
            (d_2209_v66_)[index406_] = (len((d_2120_v0_) + (d_2120_v0_))) - (default__.fm18(globalState))
            index407_ = default__.safeIndex(438, (d_2209_v66_).length(0))
            (d_2209_v66_)[index407_] = (d_2209_v66_)[default__.safeIndex(438, (d_2209_v66_).length(0))]
            if True:
                (globalState).f10 = d_2207_v64_
                index408_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                (d_2209_v66_)[index408_] = ((d_2123_v2_).f37) - ((d_2123_v2_).f37)
                index409_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                (d_2209_v66_)[index409_] = (self).f25
                d_2210_v67_: _dafny.Set
                d_2210_v67_ = _dafny.Set({d_2177_v48_, d_2177_v48_, _dafny.MultiSet([(d_2123_v2_).f37]), d_2177_v48_})
                index410_ = default__.safeIndex(210, (d_2183_v51_).length(0))
                (d_2183_v51_)[index410_] = (p0) <= (len(d_2210_v67_))
                index411_ = default__.safeIndex(210, (d_2183_v51_).length(0))
                (d_2183_v51_)[index411_] = not((d_2175_v46_) >= ((d_2123_v2_).f37))
                d_2211_v68_: C0
                nw361_ = C0()
                nw361_.ctor__()
                d_2211_v68_ = nw361_
            elif True:
                d_2212_v69_: D4
                d_2212_v69_ = D4_DC11(d_2179_v50_)
                d_2213_v70_: _dafny.MultiSet
                d_2213_v70_ = _dafny.MultiSet([d_2212_v69_])
                d_2214_v71_: _dafny.Seq
                d_2214_v71_ = _dafny.SeqWithoutIsStrInference([d_2213_v70_])
                d_2213_v70_ = (d_2214_v71_)[default__.safeIndex(d_2175_v46_, len(d_2214_v71_))]
                d_2203_cf50_ = (d_2123_v2_).f38
                d_2215_v72_: str
                d_2215_v72_ = _dafny.CodePoint('s')
                index412_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                index413_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                rhs394_ = ((d_2202_cf51_) - (d_2202_cf51_)) + ((self).f25)
                rhs395_ = d_2215_v72_
                rhs396_ = ((default__.fm15(globalState)).cardinality) + ((d_2186_v53_)[default__.safeIndex((d_2123_v2_).f37, len(d_2186_v53_))])
                lhs304_ = d_2209_v66_
                lhs305_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                lhs306_ = globalState
                lhs307_ = d_2209_v66_
                lhs308_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                lhs304_[lhs305_] = rhs394_
                lhs306_.f23 = rhs395_
                lhs307_[lhs308_] = rhs396_
                d_2216_v74_: _dafny.Map
                d_2216_v74_ = _dafny.Map({d_2175_v46_: (d_2123_v2_).f38})
                d_2217_v75_: _dafny.Set
                def iife187_():
                    coll79_ = _dafny.Map()
                    compr_79_: int
                    for compr_79_ in _dafny.IntegerRange(470, -568):
                        d_2218_v73_: int = compr_79_
                        if ((470) <= (d_2218_v73_)) and ((d_2218_v73_) < (-568)):
                            coll79_[default__.safeDivisionInt(d_2218_v73_, len(_dafny.SeqWithoutIsStrInference([d_2203_cf50_])))] = (d_2123_v2_).f38
                    return _dafny.Map(coll79_)
                d_2217_v75_ = _dafny.Set({iife187_()
                , _dafny.Map({(d_2209_v66_)[default__.safeIndex(438, (d_2209_v66_).length(0))]: (self).f32}), d_2216_v74_})
                index414_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                rhs397_ = (d_2217_v75_) - (d_2217_v75_)
                rhs398_ = (d_2123_v2_).f37
                lhs309_ = d_2209_v66_
                lhs310_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                d_2217_v75_ = rhs397_
                lhs309_[lhs310_] = rhs398_
                index415_ = default__.safeIndex(438, (d_2209_v66_).length(0))
                (d_2209_v66_)[index415_] = (d_2175_v46_) + (((self).f25) + (d_2175_v46_))
        elif source28_.is_DC31:
            d_2219___mcc_h5_ = source28_.cf46
            d_2220_cf46_ = d_2219___mcc_h5_
            d_2183_v51_ = d_2183_v51_
            d_2221_v77_: _dafny.Seq
            d_2221_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erjupre"))
            d_2222_v78_: _dafny.Map
            def iife188_():
                coll80_ = _dafny.Map()
                compr_80_: int
                for compr_80_ in _dafny.IntegerRange(329, 650):
                    d_2223_v76_: int = compr_80_
                    if ((329) <= (d_2223_v76_)) and ((d_2223_v76_) < (650)):
                        coll80_[(d_2223_v76_) - ((d_2123_v2_).f37)] = p0
                return _dafny.Map(coll80_)
            d_2222_v78_ = _dafny.Map({(0) - (default__.safeModuloInt(len(iife188_()
            ), (d_2123_v2_).f37)): (d_2221_v77_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2224_i5_ in range(default__.abs(177))]))})
            d_2222_v78_ = d_2222_v78_
            d_2175_v46_ = (0) - ((0) - (p0))
            d_2225_v79_: D5
            d_2225_v79_ = D5_DC15((d_2177_v48_) - (d_2177_v48_))
            d_2225_v79_ = d_2225_v79_
        elif True:
            d_2226___mcc_h6_ = source28_.cf52
            d_2227_cf52_ = d_2226___mcc_h6_
            d_2228_v80_: _dafny.Array
            def lambda140_(d_2229_i6_):
                return default__.safeDivisionInt(d_2229_i6_, 175)

            init79_ = lambda140_
            nw362_ = _dafny.Array(None, 17)
            for i0_79_ in range(nw362_.length(0)):
                nw362_[i0_79_] = init79_(i0_79_)
            d_2228_v80_ = nw362_
            index416_ = default__.safeIndex(723, (d_2183_v51_).length(0))
            (d_2183_v51_)[index416_] = (d_2123_v2_).f38
            index417_ = default__.safeIndex(354, (d_2228_v80_).length(0))
            (d_2228_v80_)[index417_] = (self).f25
            index418_ = default__.safeIndex(723, (d_2183_v51_).length(0))
            index419_ = default__.safeIndex(354, (d_2228_v80_).length(0))
            rhs399_ = (d_2228_v80_ if (d_2120_v0_)[default__.safeIndex((d_2123_v2_).f37, len(d_2120_v0_))] else d_2228_v80_)
            rhs400_ = (d_2123_v2_).f38
            rhs401_ = default__.safeDivisionInt((d_2175_v46_) * (d_2175_v46_), p0)
            lhs311_ = d_2183_v51_
            lhs312_ = default__.safeIndex(723, (d_2183_v51_).length(0))
            lhs313_ = d_2228_v80_
            lhs314_ = default__.safeIndex(354, (d_2228_v80_).length(0))
            d_2228_v80_ = rhs399_
            lhs311_[lhs312_] = rhs400_
            lhs313_[lhs314_] = rhs401_
            index420_ = default__.safeIndex(354, (d_2228_v80_).length(0))
            (d_2228_v80_)[index420_] = default__.safeModuloInt((d_2123_v2_).f37, p0)
            d_2230_v81_: _dafny.Array
            nw363_ = _dafny.Array(False, 4)
            d_2230_v81_ = nw363_
            d_2231_v82_: D6
            d_2231_v82_ = D6_DC18(_dafny.Map({d_2230_v81_: (d_2228_v80_)[default__.safeIndex(354, (d_2228_v80_).length(0))]}))
            source29_ = d_2231_v82_
            if source29_.is_DC19:
                d_2232___mcc_h7_ = source29_.cf29
                d_2233___mcc_h8_ = source29_.cf30
                d_2234_cf30_ = d_2233___mcc_h8_
                d_2235_cf29_ = d_2232___mcc_h7_
                rhs402_ = (d_2228_v80_)[default__.safeIndex(354, (d_2228_v80_).length(0))]
                rhs403_ = (0) - ((d_2123_v2_).f37)
                d_2175_v46_ = rhs402_
                d_2234_cf30_ = rhs403_
                index421_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                (d_2228_v80_)[index421_] = (0) - (((d_2123_v2_).f37) + ((96) - (-591)))
                d_2234_cf30_ = ((d_2186_v53_)[default__.safeIndex(d_2234_cf30_, len(d_2186_v53_))]) + ((236) + (len(d_2120_v0_)))
                d_2175_v46_ = d_2175_v46_
            elif source29_.is_DC18:
                d_2236___mcc_h9_ = source29_.cf28
                d_2237_cf28_ = d_2236___mcc_h9_
                d_2238_v83_: _dafny.Array
                nw364_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_2238_v83_ = nw364_
                index422_ = default__.safeIndex(641, (d_2238_v83_).length(0))
                (d_2238_v83_)[index422_] = d_2228_v80_
                d_2239_v84_: _dafny.Set
                d_2239_v84_ = _dafny.Set({(d_2123_v2_).f38})
                d_2240_v85_: D7
                d_2240_v85_ = D7_DC21(d_2228_v80_)
                d_2241_v86_: _dafny.MultiSet
                d_2241_v86_ = _dafny.MultiSet([(self).f32, (d_2123_v2_).f38])
                index423_ = default__.safeIndex(641, (d_2238_v83_).length(0))
                index424_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                index425_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                rhs404_ = (d_2240_v85_).cf32
                rhs405_ = d_2175_v46_
                rhs406_ = len((d_2120_v0_) + (d_2120_v0_))
                rhs407_ = d_2239_v84_
                rhs408_ = (default__.safeModuloInt((d_2241_v86_).cardinality, (d_2123_v2_).f37)) * (663)
                lhs315_ = d_2238_v83_
                lhs316_ = default__.safeIndex(641, (d_2238_v83_).length(0))
                lhs317_ = d_2228_v80_
                lhs318_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                lhs319_ = d_2228_v80_
                lhs320_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                lhs315_[lhs316_] = rhs404_
                d_2175_v46_ = rhs405_
                lhs317_[lhs318_] = rhs406_
                d_2239_v84_ = rhs407_
                lhs319_[lhs320_] = rhs408_
                d_2175_v46_ = (0) - (p0)
                index426_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                (d_2228_v80_)[index426_] = (d_2123_v2_).f37
                index427_ = default__.safeIndex(354, (d_2228_v80_).length(0))
                (d_2228_v80_)[index427_] = (self).f25
            elif True:
                d_2242___mcc_h10_ = source29_.cf31
                d_2243_cf31_ = d_2242___mcc_h10_
                d_2175_v46_ = (((d_2123_v2_).f37) + ((d_2123_v2_).f37)) * ((d_2123_v2_).f37)
                d_2175_v46_ = (_dafny.MultiSet([-635, (self).f25])).cardinality
                d_2244_v87_: _dafny.Set
                d_2244_v87_ = _dafny.Set({default__.fm39((self).f32, globalState)})
                d_2245_v88_: str
                d_2245_v88_ = _dafny.CodePoint('k')
                d_2246_v89_: _dafny.Map
                d_2246_v89_ = _dafny.Map({(d_2123_v2_).f38: d_2245_v88_})
                d_2244_v87_ = default__.fm64(default__.fm44(not((d_2123_v2_).f38), (d_2123_v2_).f38, (d_2123_v2_).f38, len(_dafny.Map({(d_2123_v2_).f38: len(d_2246_v89_)})), globalState), (d_2123_v2_).f38, d_2245_v88_, (d_2123_v2_).f38, globalState)
                d_2247_v90_: _dafny.Set
                d_2247_v90_ = _dafny.Set({(self).f32})
                d_2248_v91_: _dafny.Map
                d_2248_v91_ = _dafny.Map({(d_2247_v90_).intersection(_dafny.Set({(self).f32, (d_2123_v2_).f38, (d_2123_v2_).f38})): d_2230_v81_})
                d_2249_v92_: _dafny.Map
                d_2249_v92_ = _dafny.Map({(self).f33: d_2248_v91_})
                d_2250_v93_: _dafny.Seq
                d_2250_v93_ = _dafny.SeqWithoutIsStrInference([d_2230_v81_])
                d_2251_v94_: C6
                nw365_ = C6()
                nw365_.ctor__((d_2228_v80_)[default__.safeIndex(354, (d_2228_v80_).length(0))], (self).f26)
                d_2251_v94_ = nw365_
                d_2252_v95_: _dafny.Map
                d_2252_v95_ = _dafny.Map({d_2251_v94_: (self).f33})
                d_2248_v91_ = (d_2248_v91_) | (((d_2249_v92_)[(d_2123_v2_).f38] if ((d_2123_v2_).f38) in (d_2249_v92_) else _dafny.Map({d_2247_v90_: (d_2250_v93_)[default__.safeIndex(len(d_2252_v95_), len(d_2250_v93_))]})))
            d_2253_v96_: _dafny.Seq
            d_2253_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjlljdwo"))
            d_2254_v97_: _dafny.Set
            d_2254_v97_ = _dafny.Set({d_2253_v96_, d_2253_v96_})
            r2 = (d_2254_v97_).issubset(d_2254_v97_)
        d_2255_i7_: int
        d_2255_i7_ = 0
        with _dafny.label("16"):
            while (self).f32:
                with _dafny.c_label("16"):
                    if (d_2255_i7_) >= (100):
                        raise _dafny.Break("16")
                    d_2255_i7_ = (d_2255_i7_) + (1)
                    d_2175_v46_ = d_2175_v46_
                    d_2256_v98_: _dafny.Array
                    nw366_ = _dafny.Array(D7.default()(), 14)
                    d_2256_v98_ = nw366_
                    index428_ = default__.safeIndex(674, (d_2256_v98_).length(0))
                    (d_2256_v98_)[index428_] = D7_DC22((d_2123_v2_).f37, (self).f25)
                    d_2257_v99_: D7
                    d_2257_v99_ = D7_DC22((d_2123_v2_).f37, 649)
                    index429_ = default__.safeIndex(674, (d_2256_v98_).length(0))
                    (d_2256_v98_)[index429_] = (d_2257_v99_ if default__.fm0(392, globalState) else d_2257_v99_)
                    d_2175_v46_ = default__.safeModuloInt(len(((d_2185_v52_).set((d_2123_v2_).f38, p0)) | (d_2185_v52_)), len(_dafny.SeqWithoutIsStrInference([d_2186_v53_, _dafny.SeqWithoutIsStrInference([(self).f25 for d_2258_i8_ in range(default__.abs(-814))]), _dafny.SeqWithoutIsStrInference([(0) - ((self).f25)])])))
                    d_2259_v100_: _dafny.Array
                    def lambda141_(d_2260_v2_):
                        def lambda142_(d_2261_i9_):
                            return (d_2261_i9_) * (len(_dafny.SeqWithoutIsStrInference([(d_2260_v2_).f38])))

                        return lambda142_

                    init80_ = lambda141_(d_2123_v2_)
                    nw367_ = _dafny.Array(None, 28)
                    for i0_80_ in range(nw367_.length(0)):
                        nw367_[i0_80_] = init80_(i0_80_)
                    d_2259_v100_ = nw367_
                    index430_ = default__.safeIndex(559, (d_2259_v100_).length(0))
                    (d_2259_v100_)[index430_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))
                    index431_ = default__.safeIndex(559, (d_2259_v100_).length(0))
                    (d_2259_v100_)[index431_] = (d_2123_v2_).f37
                    pass
            pass
        d_2262_v101_: _dafny.Map
        d_2262_v101_ = _dafny.Map({d_2183_v51_: (self).f32})
        r0 = ((d_2262_v101_)[d_2183_v51_] if (d_2183_v51_) in (d_2262_v101_) else (self).f33)
        d_2263_v102_: _dafny.Map
        d_2263_v102_ = _dafny.Map({(d_2123_v2_).f38: (self).f33})
        r1 = d_2263_v102_
        d_2264_v103_: _dafny.Map
        d_2264_v103_ = _dafny.Map({(0) - ((d_2123_v2_).f37): (self).f32})
        d_2265_v104_: _dafny.Map
        d_2265_v104_ = d_2264_v103_
        d_2266_v105_: _dafny.Set
        d_2266_v105_ = _dafny.Set({d_2265_v104_})
        r2 = (d_2266_v105_).issubset(d_2266_v105_)
        r3 = (d_2120_v0_) + ((d_2120_v0_) + ((d_2120_v0_).set(default__.safeIndex((d_2123_v2_).f37, len(d_2120_v0_)), (self).f33)))
        return r0, r1, r2, r3

    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33

class C9:
    def  __init__(self):
        self.f30: _dafny.Array = _dafny.Array(None, 0)
        self._f31: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self, f30, f31):
        (self).f30 = f30
        (self)._f31 = f31

    def fm11(self, p0, p1, globalState):
        return default__.safeModuloInt(((self).f31) + (len(_dafny.Map({False: 326}))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ce"))))

    def fm12(self, p0, p1, p2, globalState):
        return (0) - ((self).f31)

    def m7(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        d_2267_v0_: bool
        d_2267_v0_ = True
        (globalState).f12 = d_2267_v0_
        d_2268_v1_: _dafny.Array
        nw368_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
        d_2268_v1_ = nw368_
        d_2269_v2_: C4
        nw369_ = C4()
        nw369_.ctor__(d_2268_v1_, p3, d_2268_v1_)
        d_2269_v2_ = nw369_
        d_2270_v3_: int
        d_2270_v3_ = 717
        d_2270_v3_ = -527
        d_2271_v4_: _dafny.Set
        d_2271_v4_ = _dafny.Set({d_2270_v3_})
        d_2272_i0_: int
        d_2272_i0_ = 0
        with _dafny.label("17"):
            while not((default__.safeDivisionInt(d_2270_v3_, len(_dafny.Map({d_2271_v4_: d_2270_v3_})))) < ((self).f31)):
                with _dafny.c_label("17"):
                    if (d_2272_i0_) >= (100):
                        raise _dafny.Break("17")
                    d_2272_i0_ = (d_2272_i0_) + (1)
                    (globalState).f12 = default__.fm0(d_2270_v3_, globalState)
                    d_2273_v5_: _dafny.Map
                    d_2273_v5_ = _dafny.Map({(p2 if d_2267_v0_ else p2): d_2267_v0_})
                    d_2274_v6_: D10
                    d_2274_v6_ = D10_DC28(d_2273_v5_)
                    d_2273_v5_ = (d_2274_v6_).cf41
                    d_2275_v7_: _dafny.Seq
                    d_2275_v7_ = _dafny.SeqWithoutIsStrInference([self.f30])
                    d_2276_v8_: int
                    out40_: int
                    out40_ = (d_2269_v2_).m3((d_2275_v7_) + (d_2275_v7_), d_2267_v0_, globalState)
                    d_2276_v8_ = out40_
                    d_2277_v9_: C1
                    nw370_ = C1()
                    nw370_.ctor__(p3, (d_2269_v2_).f36)
                    d_2277_v9_ = nw370_
                    d_2277_v9_ = d_2277_v9_
                    pass
            pass
        d_2278_v10_: _dafny.Seq
        d_2278_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwksjhjgk"))
        d_2279_v11_: _dafny.Seq
        d_2279_v11_ = _dafny.SeqWithoutIsStrInference([d_2278_v10_, d_2278_v10_, d_2278_v10_, d_2278_v10_])
        d_2280_v12_: _dafny.Map
        d_2280_v12_ = _dafny.Map({_dafny.MultiSet([self.f30]): _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lv")), (d_2279_v11_)[default__.safeIndex(d_2270_v3_, len(d_2279_v11_))]])})
        d_2281_v13_: _dafny.MultiSet
        d_2281_v13_ = _dafny.MultiSet([self.f30])
        d_2282_v14_: _dafny.Map
        d_2282_v14_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntjlswu"))})
        d_2283_v15_: _dafny.MultiSet
        d_2283_v15_ = _dafny.MultiSet([((d_2282_v14_)[p2] if (p2) in (d_2282_v14_) else d_2278_v10_), ((d_2278_v10_).set(default__.safeIndex(p3, len(d_2278_v10_)), p2)).set(default__.safeIndex(p3, len((d_2278_v10_).set(default__.safeIndex(p3, len(d_2278_v10_)), p2))), p2)])
        (globalState).f16 = ((d_2280_v12_)[d_2281_v13_] if (d_2281_v13_) in (d_2280_v12_) else d_2283_v15_)
        def iife189_():
            coll81_ = _dafny.Set()
            compr_81_: int
            for compr_81_ in (_dafny.Set({(0) - ((_dafny.MultiSet([d_2278_v10_, _dafny.SeqWithoutIsStrInference([p2 for d_2284_i2_ in range(default__.abs(-153))])])).cardinality), p3})).Elements:
                d_2285_v16_: int = compr_81_
                if (d_2285_v16_) in (_dafny.Set({(0) - ((_dafny.MultiSet([d_2278_v10_, _dafny.SeqWithoutIsStrInference([p2 for d_2284_i2_ in range(default__.abs(-153))])])).cardinality), p3})):
                    coll81_ = coll81_.union(_dafny.Set([(d_2285_v16_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")) for d_2286_i3_ in range(default__.abs(685))])))]))
            return _dafny.Set(coll81_)
        hi12_ = (0) - ((self).fm11(iife189_()
        , p2, globalState))
        for d_2287_i1_ in range(p3, hi12_):
            def lambda143_(d_2288_i1_):
                def lambda144_(d_2289_i4_):
                    return (d_2289_i4_) - (d_2288_i1_)

                return lambda144_

            init81_ = lambda143_(d_2287_i1_)
            nw371_ = _dafny.Array(None, 4)
            for i0_81_ in range(nw371_.length(0)):
                nw371_[i0_81_] = init81_(i0_81_)
            (self).f30 = nw371_
            d_2290_v17_: C2
            nw372_ = C2()
            nw372_.ctor__()
            d_2290_v17_ = nw372_
            d_2270_v3_ = d_2270_v3_
            (globalState).f12 = (p1) == (p1)
        r0 = d_2267_v0_
        r1 = d_2267_v0_
        return r0, r1

    @property
    def f31(self):
        return self._f31

class C10(T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        self._f29: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f29, f25, f26):
        (self)._f29 = f29
        (self)._f25 = f25
        (self)._f26 = f26

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_2291_i0_: int
        d_2291_i0_ = 0
        with _dafny.label("18"):
            while p2:
                with _dafny.c_label("18"):
                    if (d_2291_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_2291_i0_ = (d_2291_i0_) + (1)
                    (globalState).f23 = (self).f29
                    d_2292_v0_: _dafny.Array
                    nw373_ = _dafny.Array(_dafny.Seq({}), 5)
                    d_2292_v0_ = nw373_
                    d_2293_v1_: _dafny.Seq
                    d_2293_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnkgb"))
                    d_2294_v2_: _dafny.Seq
                    d_2294_v2_ = _dafny.SeqWithoutIsStrInference([d_2293_v1_])
                    index432_ = default__.safeIndex(133, (d_2292_v0_).length(0))
                    (d_2292_v0_)[index432_] = d_2294_v2_
                    index433_ = default__.safeIndex(676, (p3).length(0))
                    (p3)[index433_] = default__.safeModuloInt((self).f25, (self).f25)
                    index434_ = default__.safeIndex(133, (d_2292_v0_).length(0))
                    index435_ = default__.safeIndex(676, (p3).length(0))
                    rhs409_ = d_2294_v2_
                    rhs410_ = (((self).f25 if p2 else (self).f25)) + (-425)
                    lhs321_ = d_2292_v0_
                    lhs322_ = default__.safeIndex(133, (d_2292_v0_).length(0))
                    lhs323_ = p3
                    lhs324_ = default__.safeIndex(676, (p3).length(0))
                    lhs321_[lhs322_] = rhs409_
                    lhs323_[lhs324_] = rhs410_
                    if p1:
                        d_2295_v3_: C0
                        nw374_ = C0()
                        nw374_.ctor__()
                        d_2295_v3_ = nw374_
                        d_2296_v4_: _dafny.Array
                        def lambda145_(d_2297_p2_, d_2298_p3_):
                            def lambda146_(d_2299_i1_):
                                return (_dafny.Set({(self).f25})).ispropersubset(_dafny.Set({(_dafny.MultiSet([(D10_DC29(d_2297_p2_, d_2297_p2_, (d_2298_p3_)[default__.safeIndex(676, (d_2298_p3_).length(0))])).cf44, (d_2298_p3_)[default__.safeIndex(676, (d_2298_p3_).length(0))]])).cardinality}))

                            return lambda146_

                        init82_ = lambda145_(p2, p3)
                        nw375_ = _dafny.Array(None, 21)
                        for i0_82_ in range(nw375_.length(0)):
                            nw375_[i0_82_] = init82_(i0_82_)
                        d_2296_v4_ = nw375_
                        d_2296_v4_ = d_2296_v4_
                        index436_ = default__.safeIndex(676, (p3).length(0))
                        (p3)[index436_] = (self).f25
                        d_2300_v5_: C5
                        nw376_ = C5()
                        nw376_.ctor__(False)
                        d_2300_v5_ = nw376_
                        (globalState).f23 = (self).f29
                    elif True:
                        d_2301_v6_: _dafny.MultiSet
                        d_2301_v6_ = _dafny.MultiSet([p1, p2])
                        d_2302_v7_: _dafny.Map
                        d_2302_v7_ = _dafny.Map({(self).f25: p1})
                        d_2303_v8_: D0
                        d_2303_v8_ = D0_DC3(((d_2301_v6_).set(p1, default__.abs(len(d_2302_v7_)))).cardinality, p2)
                        index437_ = default__.safeIndex(676, (p3).length(0))
                        (p3)[index437_] = (0) - ((((p3)[default__.safeIndex(676, (p3).length(0))]) * ((d_2303_v8_).cf4)) - ((self).f25))
                        index438_ = default__.safeIndex(676, (p3).length(0))
                        (p3)[index438_] = ((0) - ((self).f25) if True else (self).f25)
                        d_2304_v9_: _dafny.Seq
                        d_2304_v9_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                        index439_ = default__.safeIndex(676, (p3).length(0))
                        rhs411_ = not(((self).f25) != ((p3)[default__.safeIndex(676, (p3).length(0))]))
                        rhs412_ = ((p3)[default__.safeIndex(676, (p3).length(0))]) - (len(_dafny.SeqWithoutIsStrInference([len(d_2293_v1_) for d_2305_i2_ in range(default__.abs(641))])))
                        rhs413_ = ((len(d_2304_v9_)) + ((p3)[default__.safeIndex(676, (p3).length(0))])) + ((self).f25)
                        rhs414_ = ((p3)[default__.safeIndex(676, (p3).length(0))]) >= ((p3)[default__.safeIndex(676, (p3).length(0))])
                        lhs325_ = globalState
                        lhs326_ = p3
                        lhs327_ = default__.safeIndex(676, (p3).length(0))
                        lhs328_ = globalState
                        lhs325_.f12 = rhs411_
                        lhs326_[lhs327_] = rhs412_
                        r0 = rhs413_
                        lhs328_.f12 = rhs414_
                        (globalState).f6 = ((d_2304_v9_) + (_dafny.SeqWithoutIsStrInference([(p3)[default__.safeIndex(676, (p3).length(0))] for d_2306_i3_ in range(default__.abs(81))]))) + (d_2304_v9_)
                        r1 = ((self).f25) * ((0) - ((p3)[default__.safeIndex(676, (p3).length(0))]))
                    (globalState).f10 = (_dafny.SeqWithoutIsStrInference([(self).f29 for d_2307_i4_ in range(default__.abs(650))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmuer")))
                    pass
            pass
        r1 = (self).f25
        d_2308_v10_: C6
        nw377_ = C6()
        nw377_.ctor__((self).f25, (self).f26)
        d_2308_v10_ = nw377_
        hi13_ = ((self).f25) - ((self).f25)
        for d_2309_i5_ in range((self).f25, hi13_):
            d_2310_v11_: _dafny.Seq
            d_2310_v11_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26, (self).f26, (self).f26])
            d_2311_v12_: C3
            nw378_ = C3()
            nw378_.ctor__(d_2309_i5_, p1, (self).f25, (d_2310_v11_)[default__.safeIndex((self).f25, len(d_2310_v11_))])
            d_2311_v12_ = nw378_
            r0 = (self).f25
            d_2312_v13_: _dafny.Seq
            d_2312_v13_ = _dafny.SeqWithoutIsStrInference([(d_2311_v12_).f38, not(True), default__.fm0((0) - ((self).f25), globalState)])
            d_2313_v14_: _dafny.Array
            nw379_ = _dafny.Array(False, 25)
            d_2313_v14_ = nw379_
            d_2314_v15_: _dafny.Map
            d_2314_v15_ = _dafny.Map({default__.fm0((self).f25, globalState): d_2313_v14_})
            rhs415_ = default__.fm1(p1, globalState)
            rhs416_ = (self).f25
            rhs417_ = default__.fm0((d_2311_v12_).f37, globalState)
            rhs418_ = (d_2313_v14_ if p2 else ((d_2314_v15_)[p2] if (p2) in (d_2314_v15_) else d_2313_v14_))
            rhs419_ = (d_2311_v12_).f38
            lhs329_ = globalState
            lhs330_ = globalState
            d_2312_v13_ = rhs415_
            r0 = rhs416_
            lhs329_.f12 = rhs417_
            d_2313_v14_ = rhs418_
            lhs330_.f12 = rhs419_
            d_2315_v16_: _dafny.Seq
            d_2315_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqn"))
            d_2316_v17_: D10
            d_2316_v17_ = D10_DC29((d_2311_v12_).f38, p2, (self).f25)
            d_2317_v18_: _dafny.MultiSet
            d_2317_v18_ = _dafny.MultiSet([(d_2316_v17_).cf42])
            if (d_2317_v18_).ispropersubset((_dafny.MultiSet(d_2312_v13_)).intersection(default__.fm50((d_2311_v12_).f37, d_2315_v16_, p1, -333, globalState))):
                d_2318_v19_: _dafny.Seq
                d_2318_v19_ = _dafny.SeqWithoutIsStrInference([(d_2311_v12_).f37, (d_2311_v12_).f37])
                d_2319_v20_: _dafny.Seq
                d_2319_v20_ = _dafny.SeqWithoutIsStrInference([d_2318_v19_, d_2318_v19_])
                d_2319_v20_ = d_2319_v20_
                d_2313_v14_ = d_2313_v14_
                d_2320_v21_: C0
                nw380_ = C0()
                nw380_.ctor__()
                d_2320_v21_ = nw380_
                (globalState).f12 = p2
                (globalState).f12 = p2
            elif True:
                d_2321_v22_: _dafny.Set
                d_2321_v22_ = _dafny.Set({p2, not(p2)})
                d_2322_v23_: _dafny.MultiSet
                d_2322_v23_ = _dafny.MultiSet([default__.fm20(p2, globalState), d_2321_v22_])
                d_2322_v23_ = d_2322_v23_
                d_2323_v24_: _dafny.Map
                d_2323_v24_ = _dafny.Map({p2: (p2) and (p2)})
                d_2323_v24_ = (d_2323_v24_).set(False, (d_2311_v12_).f38)
                d_2324_v25_: _dafny.Array
                nw381_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_2324_v25_ = nw381_
                d_2324_v25_ = d_2324_v25_
                (globalState).f12 = True
                d_2325_v26_: D3
                d_2326_v27_: bool
                out41_: D3
                out42_: bool
                out41_, out42_ = (d_2308_v10_).m11(p3, (d_2311_v12_).f38, (d_2314_v15_) | (d_2314_v15_), globalState)
                d_2325_v26_ = out41_
                d_2326_v27_ = out42_
        if p1:
            (globalState).f12 = p1
            r0 = ((0) - (default__.safeModuloInt(default__.fm44(p1, p2, False, (self).f25, globalState), (self).f25))) + ((self).f25)
            (globalState).f12 = not(p2)
            if p2:
                d_2327_v28_: D13
                d_2327_v28_ = D13_DC36(_dafny.Map({p2: p2}), p2, (self).f25)
                d_2328_v29_: _dafny.Map
                d_2328_v29_ = _dafny.Map({p1: p1})
                d_2329_v30_: _dafny.Seq
                d_2329_v30_ = _dafny.SeqWithoutIsStrInference([p2])
                pat_let_tv43_ = d_2328_v29_
                d_2330_v31_: _dafny.Array
                nw382_ = _dafny.Array(None, 13)
                nw382_[int(0)] = d_2327_v28_
                nw382_[int(1)] = D13_DC36(d_2328_v29_, False, len(d_2329_v30_))
                nw382_[int(2)] = d_2327_v28_
                nw382_[int(3)] = d_2327_v28_
                nw382_[int(4)] = d_2327_v28_
                def iife190_(_pat_let54_0):
                    def iife191_(d_2331_dt__update__tmp_h0_):
                        def iife192_(_pat_let55_0):
                            def iife193_(d_2332_dt__update_hcf54_h0_):
                                return D13_DC36(d_2332_dt__update_hcf54_h0_, (d_2331_dt__update__tmp_h0_).cf55, (d_2331_dt__update__tmp_h0_).cf56)
                            return iife193_(_pat_let55_0)
                        return iife192_(pat_let_tv43_)
                    return iife191_(_pat_let54_0)
                nw382_[int(5)] = iife190_(D13_DC36(d_2328_v29_, p1, (self).f25))
                nw382_[int(6)] = d_2327_v28_
                nw382_[int(7)] = d_2327_v28_
                nw382_[int(8)] = d_2327_v28_
                nw382_[int(9)] = d_2327_v28_
                nw382_[int(10)] = d_2327_v28_
                nw382_[int(11)] = d_2327_v28_
                nw382_[int(12)] = D13_DC36(d_2328_v29_, p2, (0) - ((self).f25))
                d_2330_v31_ = nw382_
                index440_ = default__.safeIndex(884, (d_2330_v31_).length(0))
                (d_2330_v31_)[index440_] = d_2327_v28_
                index441_ = default__.safeIndex(696, ((self).f26).length(0))
                ((self).f26)[index441_] = _dafny.SeqWithoutIsStrInference([default__.fm39(p2, globalState), (self).f29, (self).f29, (self).f29])
                d_2333_v32_: _dafny.Seq
                d_2333_v32_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                index442_ = default__.safeIndex(884, (d_2330_v31_).length(0))
                index443_ = default__.safeIndex(696, ((self).f26).length(0))
                rhs420_ = d_2327_v28_
                rhs421_ = d_2333_v32_
                rhs422_ = (d_2333_v32_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbgxd")))
                lhs331_ = d_2330_v31_
                lhs332_ = default__.safeIndex(884, (d_2330_v31_).length(0))
                lhs333_ = (self).f26
                lhs334_ = default__.safeIndex(696, ((self).f26).length(0))
                lhs335_ = globalState
                lhs331_[lhs332_] = rhs420_
                lhs333_[lhs334_] = rhs421_
                lhs335_.f12 = rhs422_
                d_2334_v33_: _dafny.Array
                nw383_ = _dafny.Array(False, 16)
                d_2334_v33_ = nw383_
                d_2334_v33_ = d_2334_v33_
                d_2335_v34_: _dafny.Set
                d_2335_v34_ = _dafny.Set({len((_dafny.SeqWithoutIsStrInference([(self).f29])).set(default__.safeIndex((self).f25, len(_dafny.SeqWithoutIsStrInference([(self).f29]))), _dafny.CodePoint('y')))})
                d_2336_v35_: _dafny.Map
                d_2336_v35_ = _dafny.Map({len(d_2335_v34_): d_2333_v32_})
                index444_ = default__.safeIndex(718, (p3).length(0))
                (p3)[index444_] = len((((d_2336_v35_)[(self).f25] if ((self).f25) in (d_2336_v35_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))) + (d_2333_v32_))
                index445_ = default__.safeIndex(718, (p3).length(0))
                (p3)[index445_] = (self).f25
                d_2337_v36_: _dafny.Array
                def lambda147_(d_2338_v30_, d_2339_p1_, d_2340_p2_):
                    def lambda148_(d_2341_i6_):
                        return (_dafny.MultiSet(d_2338_v30_)) | (_dafny.MultiSet([not(d_2339_p1_), d_2340_p2_]))

                    return lambda148_

                init83_ = lambda147_(d_2329_v30_, p1, p2)
                nw384_ = _dafny.Array(None, 25)
                for i0_83_ in range(nw384_.length(0)):
                    nw384_[i0_83_] = init83_(i0_83_)
                d_2337_v36_ = nw384_
                d_2342_v37_: D12
                d_2342_v37_ = D12_DC32(p2)
                index446_ = default__.safeIndex(997, (d_2337_v36_).length(0))
                (d_2337_v36_)[index446_] = _dafny.MultiSet([(d_2342_v37_).cf47])
                d_2343_v38_: _dafny.MultiSet
                d_2343_v38_ = _dafny.MultiSet([not(p1)])
                d_2344_v39_: _dafny.Map
                d_2344_v39_ = _dafny.Map({(self).f25: d_2343_v38_})
                d_2345_v40_: _dafny.Seq
                d_2345_v40_ = _dafny.SeqWithoutIsStrInference([len((((self).f26)[default__.safeIndex(696, ((self).f26).length(0))]).set(default__.safeIndex((p3)[default__.safeIndex(718, (p3).length(0))], len(((self).f26)[default__.safeIndex(696, ((self).f26).length(0))])), (self).f29))])
                index447_ = default__.safeIndex(997, (d_2337_v36_).length(0))
                (d_2337_v36_)[index447_] = ((d_2344_v39_)[(0) - (((p3)[default__.safeIndex(718, (p3).length(0))]) * ((d_2345_v40_)[default__.safeIndex((self).f25, len(d_2345_v40_))]))] if ((0) - (((p3)[default__.safeIndex(718, (p3).length(0))]) * ((d_2345_v40_)[default__.safeIndex((self).f25, len(d_2345_v40_))]))) in (d_2344_v39_) else d_2343_v38_)
                (globalState).f12 = ((d_2329_v30_) + (d_2329_v30_)) != ((_dafny.SeqWithoutIsStrInference([p1])) + (d_2329_v30_))
            elif True:
                d_2346_v41_: _dafny.Array
                nw385_ = _dafny.Array(False, 17)
                d_2346_v41_ = nw385_
                d_2347_v42_: _dafny.Set
                d_2347_v42_ = _dafny.Set({d_2346_v41_})
                d_2348_v43_: D3
                d_2348_v43_ = D3_DC9(d_2347_v42_)
                d_2349_v44_: _dafny.MultiSet
                d_2349_v44_ = _dafny.MultiSet([d_2348_v43_, d_2348_v43_, d_2348_v43_])
                d_2350_v45_: _dafny.Seq
                d_2350_v45_ = _dafny.SeqWithoutIsStrInference([d_2349_v44_, (d_2349_v44_).intersection(d_2349_v44_), d_2349_v44_, d_2349_v44_, (d_2349_v44_) - (_dafny.MultiSet([d_2348_v43_]))])
                d_2351_v46_: _dafny.Map
                d_2351_v46_ = _dafny.Map({(self).f29: p2})
                d_2352_v47_: _dafny.Map
                d_2352_v47_ = _dafny.Map({(self).f25: len(d_2351_v46_)})
                d_2353_v48_: _dafny.Seq
                d_2353_v48_ = _dafny.SeqWithoutIsStrInference([len(d_2352_v47_)])
                d_2354_v49_: _dafny.Map
                d_2354_v49_ = _dafny.Map({p1: d_2353_v48_})
                rhs423_ = d_2350_v45_
                rhs424_ = p1
                rhs425_ = ((self).f25) * (len(((d_2354_v49_)[not(not(p1))] if (not(not(p1))) in (d_2354_v49_) else d_2353_v48_)))
                lhs336_ = globalState
                d_2350_v45_ = rhs423_
                lhs336_.f12 = rhs424_
                r0 = rhs425_
                d_2355_v50_: _dafny.Array
                nw386_ = _dafny.Array(None, 20)
                nw386_[int(0)] = p3
                nw386_[int(1)] = p3
                nw386_[int(2)] = p3
                nw386_[int(3)] = p3
                nw386_[int(4)] = p3
                nw386_[int(5)] = p3
                nw386_[int(6)] = p3
                nw386_[int(7)] = p3
                nw386_[int(8)] = p3
                nw386_[int(9)] = p3
                nw386_[int(10)] = p3
                nw386_[int(11)] = p3
                nw386_[int(12)] = p3
                nw386_[int(13)] = (p3 if p2 else p3)
                nw386_[int(14)] = p3
                nw386_[int(15)] = p3
                nw386_[int(16)] = p3
                nw386_[int(17)] = p3
                nw386_[int(18)] = p3
                nw386_[int(19)] = p3
                d_2355_v50_ = nw386_
                d_2355_v50_ = d_2355_v50_
                r0 = 551
                (globalState).f12 = ((self).f25) != ((0) - ((self).f25))
                d_2356_v51_: _dafny.Map
                d_2356_v51_ = _dafny.Map({(self).f25: d_2353_v48_})
                d_2357_v52_: _dafny.MultiSet
                d_2357_v52_ = _dafny.MultiSet([(self).f25])
                (globalState).f12 = (not(p2) if p2 else (_dafny.MultiSet(((d_2356_v51_)[len(d_2353_v48_)] if (len(d_2353_v48_)) in (d_2356_v51_) else d_2353_v48_))).issubset(d_2357_v52_))
            index448_ = default__.safeIndex(76, (p3).length(0))
            (p3)[index448_] = default__.safeModuloInt((self).f25, (self).f25)
            d_2358_v53_: _dafny.MultiSet
            d_2358_v53_ = _dafny.MultiSet([p1])
            d_2359_v54_: _dafny.Seq
            d_2359_v54_ = _dafny.SeqWithoutIsStrInference([p1, not(p2)])
            d_2360_v55_: D0
            d_2360_v55_ = D0_DC1(not(p2), p1)
            d_2361_v56_: _dafny.Map
            d_2361_v56_ = _dafny.Map({d_2360_v55_: 918})
            d_2362_v57_: _dafny.Seq
            d_2362_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jveby"))
            d_2363_v58_: _dafny.Map
            d_2363_v58_ = _dafny.Map({p2: p1})
            d_2364_v59_: _dafny.Map
            d_2364_v59_ = _dafny.Map({p2: (0) - ((self).f25)})
            d_2365_v60_: _dafny.Array
            nw387_ = _dafny.Array(None, 21)
            nw387_[int(0)] = d_2358_v53_
            nw387_[int(1)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2, default__.fm0((self).f25, globalState)]))
            nw387_[int(2)] = _dafny.MultiSet(d_2359_v54_)
            nw387_[int(3)] = (d_2358_v53_).set(p1, default__.abs((self).f25))
            nw387_[int(4)] = d_2358_v53_
            nw387_[int(5)] = default__.fm50(len((d_2361_v56_).set(d_2360_v55_, (self).f25)), d_2362_v57_, p1, (self).f25, globalState)
            nw387_[int(6)] = d_2358_v53_
            nw387_[int(7)] = (d_2358_v53_) - (_dafny.MultiSet(d_2359_v54_))
            nw387_[int(8)] = d_2358_v53_
            nw387_[int(9)] = d_2358_v53_
            nw387_[int(10)] = d_2358_v53_
            nw387_[int(11)] = (_dafny.MultiSet([p2])).set(p1, default__.abs((self).f25))
            nw387_[int(12)] = (d_2358_v53_).set(p2, default__.abs(len(d_2363_v58_)))
            nw387_[int(13)] = d_2358_v53_
            nw387_[int(14)] = d_2358_v53_
            nw387_[int(15)] = default__.fm50(len(d_2364_v59_), d_2362_v57_, True, (self).f25, globalState)
            nw387_[int(16)] = d_2358_v53_
            nw387_[int(17)] = d_2358_v53_
            nw387_[int(18)] = d_2358_v53_
            nw387_[int(19)] = d_2358_v53_
            nw387_[int(20)] = (d_2358_v53_) | (d_2358_v53_)
            d_2365_v60_ = nw387_
            index449_ = default__.safeIndex(0, (d_2365_v60_).length(0))
            (d_2365_v60_)[index449_] = d_2358_v53_
            index450_ = default__.safeIndex(76, (p3).length(0))
            index451_ = default__.safeIndex(0, (d_2365_v60_).length(0))
            rhs426_ = (0) - (default__.safeDivisionInt(len(_dafny.Map({p1: p2})), (self).f25))
            rhs427_ = (default__.fm18(globalState)) * (-367)
            rhs428_ = d_2358_v53_
            lhs337_ = p3
            lhs338_ = default__.safeIndex(76, (p3).length(0))
            lhs339_ = d_2365_v60_
            lhs340_ = default__.safeIndex(0, (d_2365_v60_).length(0))
            r1 = rhs426_
            lhs337_[lhs338_] = rhs427_
            lhs339_[lhs340_] = rhs428_
        elif True:
            index452_ = default__.safeIndex(246, (p3).length(0))
            (p3)[index452_] = ((self).f25) * (len(_dafny.SeqWithoutIsStrInference([(self).f25, 243])))
            d_2366_v61_: _dafny.Array
            def lambda149_(d_2367_p1_, d_2368_p2_):
                def lambda150_(d_2369_i7_):
                    return (_dafny.Map({d_2367_p1_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2370_i8_ in range(default__.abs(163))])})) | (_dafny.Map({d_2368_p2_: _dafny.SeqWithoutIsStrInference([(self).f29 for d_2371_i9_ in range(default__.abs(337))])}))

                return lambda150_

            init84_ = lambda149_(p1, p2)
            nw388_ = _dafny.Array(None, 14)
            for i0_84_ in range(nw388_.length(0)):
                nw388_[i0_84_] = init84_(i0_84_)
            d_2366_v61_ = nw388_
            d_2372_v62_: _dafny.Seq
            d_2372_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "is"))
            d_2373_v63_: _dafny.Map
            d_2373_v63_ = _dafny.Map({p1: d_2372_v62_})
            index453_ = default__.safeIndex(948, (d_2366_v61_).length(0))
            (d_2366_v61_)[index453_] = d_2373_v63_
            d_2374_v64_: _dafny.MultiSet
            d_2374_v64_ = _dafny.MultiSet([p1])
            d_2375_v65_: _dafny.Set
            d_2375_v65_ = _dafny.Set({default__.fm44(p1, p1, p2, (self).f25, globalState), (d_2374_v64_).cardinality, default__.safeModuloInt((self).f25, len(_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25])))})
            index454_ = default__.safeIndex(246, (p3).length(0))
            index455_ = default__.safeIndex(948, (d_2366_v61_).length(0))
            rhs429_ = len(d_2375_v65_)
            rhs430_ = d_2373_v63_
            lhs341_ = p3
            lhs342_ = default__.safeIndex(246, (p3).length(0))
            lhs343_ = d_2366_v61_
            lhs344_ = default__.safeIndex(948, (d_2366_v61_).length(0))
            lhs341_[lhs342_] = rhs429_
            lhs343_[lhs344_] = rhs430_
            if (p1 if p2 else p2):
                d_2376_v66_: _dafny.Map
                d_2376_v66_ = _dafny.Map({77: d_2372_v62_})
                d_2377_v67_: _dafny.Map
                d_2377_v67_ = _dafny.Map({p2: d_2376_v66_})
                d_2377_v67_ = (d_2377_v67_).set(not(p2), (d_2376_v66_) | (d_2376_v66_))
                d_2378_v68_: _dafny.Array
                def lambda151_(d_2379_i10_):
                    return (self).f29

                init85_ = lambda151_
                nw389_ = _dafny.Array(None, 25)
                for i0_85_ in range(nw389_.length(0)):
                    nw389_[i0_85_] = init85_(i0_85_)
                d_2378_v68_ = nw389_
                d_2380_v69_: D13
                d_2380_v69_ = D13_DC35(d_2378_v68_)
                d_2381_v70_: _dafny.Set
                d_2381_v70_ = _dafny.Set({d_2380_v69_, d_2380_v69_})
                d_2381_v70_ = d_2381_v70_
                index456_ = default__.safeIndex(246, (p3).length(0))
                (p3)[index456_] = (0) - ((0) - ((self).f25))
                (globalState).f12 = not(p2)
                d_2382_v71_: _dafny.Seq
                d_2382_v71_ = _dafny.SeqWithoutIsStrInference([p1, p1, p2])
                d_2383_v72_: _dafny.Seq
                d_2383_v72_ = _dafny.SeqWithoutIsStrInference([d_2382_v71_])
                d_2383_v72_ = (d_2383_v72_) + ((_dafny.SeqWithoutIsStrInference([(D1_DC4(d_2382_v71_)).cf6 for d_2384_i11_ in range(default__.abs(598))])) + (d_2383_v72_))
            elif True:
                index457_ = default__.safeIndex(246, (p3).length(0))
                def iife194_():
                    coll82_ = _dafny.Map()
                    compr_82_: int
                    for compr_82_ in _dafny.IntegerRange(298, 891):
                        d_2385_v73_: int = compr_82_
                        if ((298) <= (d_2385_v73_)) and ((d_2385_v73_) < (891)):
                            coll82_[(d_2385_v73_) + ((self).f25)] = p1
                    return _dafny.Map(coll82_)
                (p3)[index457_] = default__.safeDivisionInt(len(iife194_()
                ), 374)
                (globalState).f12 = p2
                (globalState).f12 = ((_dafny.SeqWithoutIsStrInference([(self).f29 for d_2386_i12_ in range(default__.abs(905))])) + ((d_2372_v62_).set(default__.safeIndex((p3)[default__.safeIndex(246, (p3).length(0))], len(d_2372_v62_)), (self).f29))) < ((d_2372_v62_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsanvurs"))))
                d_2387_v74_: _dafny.Seq
                d_2387_v74_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                d_2388_v75_: _dafny.Map
                d_2388_v75_ = _dafny.Map({(self).f25: d_2372_v62_})
                d_2389_v77_: _dafny.Seq
                def iife195_():
                    coll83_ = _dafny.Set()
                    compr_83_: int
                    for compr_83_ in _dafny.IntegerRange(362, 861):
                        d_2390_v76_: int = compr_83_
                        if ((362) <= (d_2390_v76_)) and ((d_2390_v76_) < (861)):
                            coll83_ = coll83_.union(_dafny.Set([(d_2390_v76_) + (len(_dafny.Map({(0) - (len(_dafny.Map({p1: False}))): 43})))]))
                    return _dafny.Set(coll83_)
                d_2389_v77_ = _dafny.SeqWithoutIsStrInference([iife195_()
                , d_2375_v65_])
                d_2391_v78_: _dafny.Map
                d_2391_v78_ = _dafny.Map({D18_DC44((self).f25, p2, d_2372_v62_, p2, d_2389_v77_): (p3)[default__.safeIndex(246, (p3).length(0))]})
                d_2392_v80_: _dafny.Set
                d_2392_v80_ = _dafny.Set({D18_DC44((p3)[default__.safeIndex(246, (p3).length(0))], p1, d_2372_v62_, True, d_2389_v77_)})
                d_2393_v81_: _dafny.Map
                d_2393_v81_ = _dafny.Map({p2: (p3)[default__.safeIndex(246, (p3).length(0))]})
                d_2394_v82_: _dafny.Map
                d_2394_v82_ = _dafny.Map({489: d_2393_v81_})
                d_2395_v83_: D5
                d_2395_v83_ = D5_DC17(len(((d_2394_v82_)[(self).f25] if ((self).f25) in (d_2394_v82_) else d_2393_v81_)), p1)
                d_2396_v84_: _dafny.Array
                nw390_ = _dafny.Array(None, 17)
                nw390_[int(0)] = default__.fm0((p3)[default__.safeIndex(246, (p3).length(0))], globalState)
                nw390_[int(1)] = p1
                nw390_[int(2)] = (d_2387_v74_)[default__.safeIndex(len(((d_2388_v75_)[-237] if (-237) in (d_2388_v75_) else d_2372_v62_)), len(d_2387_v74_))]
                def iife196_():
                    coll84_ = _dafny.Set()
                    compr_84_: D18
                    for compr_84_ in (d_2391_v78_).keys.Elements:
                        d_2397_v79_: D18 = compr_84_
                        if (d_2397_v79_) in (d_2391_v78_):
                            coll84_ = coll84_.union(_dafny.Set([d_2397_v79_]))
                    return _dafny.Set(coll84_)
                nw390_[int(3)] = (d_2392_v80_).ispropersubset(iife196_()
                )
                nw390_[int(4)] = (282) >= (848)
                nw390_[int(5)] = (d_2395_v83_).cf27
                nw390_[int(6)] = (d_2387_v74_)[default__.safeIndex(22, len(d_2387_v74_))]
                nw390_[int(7)] = p1
                nw390_[int(8)] = p2
                nw390_[int(9)] = (p2 if not(p1) else not(False))
                nw390_[int(10)] = p2
                nw390_[int(11)] = p1
                nw390_[int(12)] = p1
                nw390_[int(13)] = False
                nw390_[int(14)] = not(p1)
                nw390_[int(15)] = (d_2387_v74_) < (d_2387_v74_)
                nw390_[int(16)] = p2
                d_2396_v84_ = nw390_
                index458_ = default__.safeIndex(187, (d_2396_v84_).length(0))
                (d_2396_v84_)[index458_] = (d_2372_v62_) == (d_2372_v62_)
                index459_ = default__.safeIndex(187, (d_2396_v84_).length(0))
                (d_2396_v84_)[index459_] = not(((d_2372_v62_)[default__.safeIndex(546, len(d_2372_v62_))]) not in ((d_2372_v62_) + (d_2372_v62_)))
                d_2396_v84_ = d_2396_v84_
            (globalState).f10 = d_2372_v62_
            r0 = (p3)[default__.safeIndex(246, (p3).length(0))]
            d_2398_v85_: _dafny.Array
            nw391_ = _dafny.Array(_dafny.MultiSet({}), 5)
            d_2398_v85_ = nw391_
            d_2399_v86_: _dafny.Map
            d_2399_v86_ = _dafny.Map({(0) - (len(d_2372_v62_)): False})
            index460_ = default__.safeIndex(180, (d_2398_v85_).length(0))
            (d_2398_v85_)[index460_] = _dafny.MultiSet([d_2399_v86_])
            d_2400_v88_: _dafny.MultiSet
            d_2400_v88_ = _dafny.MultiSet([len(d_2372_v62_), (p3)[default__.safeIndex(246, (p3).length(0))]])
            d_2401_v89_: _dafny.MultiSet
            def iife197_():
                coll85_ = _dafny.Map()
                compr_85_: int
                for compr_85_ in (d_2400_v88_).Elements:
                    d_2402_v87_: int = compr_85_
                    if (d_2402_v87_) in (d_2400_v88_):
                        coll85_[(d_2402_v87_) - ((self).f25)] = p2
                return _dafny.Map(coll85_)
            d_2401_v89_ = _dafny.MultiSet([iife197_()
            , d_2399_v86_, d_2399_v86_])
            index461_ = default__.safeIndex(180, (d_2398_v85_).length(0))
            (d_2398_v85_)[index461_] = d_2401_v89_
        index462_ = default__.safeIndex(251, (p3).length(0))
        (p3)[index462_] = ((self).f25) * ((self).f25)
        index463_ = default__.safeIndex(251, (p3).length(0))
        (p3)[index463_] = (self).f25
        d_2403_v90_: _dafny.Map
        d_2403_v90_ = _dafny.Map({p2: p1})
        d_2404_v91_: _dafny.Seq
        d_2404_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qinj"))
        d_2405_v92_: _dafny.Set
        d_2405_v92_ = _dafny.Set({636})
        d_2406_v93_: D18
        d_2406_v93_ = D18_DC44(len(_dafny.Map({p1: ((d_2403_v90_)[p1] if (p1) in (d_2403_v90_) else p2)})), True, d_2404_v91_, p2, _dafny.SeqWithoutIsStrInference([d_2405_v92_, d_2405_v92_]))
        pat_let_tv44_ = p3
        pat_let_tv45_ = p3
        pat_let_tv46_ = p2
        pat_let_tv47_ = p2
        def lambda152_(source30_):
            if source30_.is_DC44:
                d_2407___mcc_h0_ = source30_.cf63
                d_2408___mcc_h1_ = source30_.cf64
                d_2409___mcc_h2_ = source30_.cf65
                d_2410___mcc_h3_ = source30_.cf66
                d_2411___mcc_h4_ = source30_.cf67
                d_2412_cf67_ = d_2411___mcc_h4_
                d_2413_cf66_ = d_2410___mcc_h3_
                d_2414_cf65_ = d_2409___mcc_h2_
                d_2415_cf64_ = d_2408___mcc_h1_
                d_2416_cf63_ = d_2407___mcc_h0_
                return d_2416_cf63_
            elif source30_.is_DC45:
                return (pat_let_tv45_)[default__.safeIndex(251, (pat_let_tv44_).length(0))]
            elif True:
                d_2417___mcc_h5_ = source30_.cf62
                d_2418_cf62_ = d_2417___mcc_h5_
                return len(_dafny.Set({False, pat_let_tv46_, pat_let_tv47_}))

        r0 = lambda152_(d_2406_v93_)
        r1 = (p3)[default__.safeIndex(251, (p3).length(0))]
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        d_2419_i0_: int
        d_2419_i0_ = 0
        with _dafny.label("19"):
            while p1:
                with _dafny.c_label("19"):
                    if (d_2419_i0_) >= (100):
                        raise _dafny.Break("19")
                    d_2419_i0_ = (d_2419_i0_) + (1)
                    (globalState).f23 = (self).f29
                    d_2420_v0_: D0
                    d_2420_v0_ = D0_DC2(p1)
                    d_2421_v1_: _dafny.MultiSet
                    d_2421_v1_ = _dafny.MultiSet([d_2420_v0_])
                    d_2422_v2_: _dafny.Seq
                    d_2422_v2_ = _dafny.SeqWithoutIsStrInference([d_2420_v0_])
                    d_2423_v3_: _dafny.Seq
                    d_2423_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_2422_v2_), d_2421_v1_])
                    d_2424_v4_: _dafny.MultiSet
                    d_2424_v4_ = _dafny.MultiSet([p1])
                    pat_let_tv48_ = p1
                    d_2425_v5_: _dafny.Array
                    nw392_ = _dafny.Array(None, 22)
                    nw392_[int(0)] = d_2421_v1_
                    nw392_[int(1)] = (d_2421_v1_).set(d_2420_v0_, default__.abs((self).f25))
                    nw392_[int(2)] = d_2421_v1_
                    nw392_[int(3)] = d_2421_v1_
                    nw392_[int(4)] = d_2421_v1_
                    nw392_[int(5)] = d_2421_v1_
                    nw392_[int(6)] = d_2421_v1_
                    nw392_[int(7)] = d_2421_v1_
                    nw392_[int(8)] = d_2421_v1_
                    nw392_[int(9)] = d_2421_v1_
                    def iife198_(_pat_let56_0):
                        def iife199_(d_2426_dt__update__tmp_h0_):
                            def iife200_(_pat_let57_0):
                                def iife201_(d_2427_dt__update_hcf3_h0_):
                                    return D0_DC2(d_2427_dt__update_hcf3_h0_)
                                return iife201_(_pat_let57_0)
                            return iife200_(pat_let_tv48_)
                        return iife199_(_pat_let56_0)
                    nw392_[int(10)] = _dafny.MultiSet([iife198_(d_2420_v0_)])
                    nw392_[int(11)] = d_2421_v1_
                    nw392_[int(12)] = d_2421_v1_
                    nw392_[int(13)] = (d_2423_v3_)[default__.safeIndex(27, len(d_2423_v3_))]
                    nw392_[int(14)] = (d_2421_v1_ if p1 else d_2421_v1_)
                    nw392_[int(15)] = (d_2421_v1_) - ((d_2421_v1_).set(d_2420_v0_, default__.abs((self).f25)))
                    nw392_[int(16)] = d_2421_v1_
                    nw392_[int(17)] = d_2421_v1_
                    nw392_[int(18)] = d_2421_v1_
                    nw392_[int(19)] = _dafny.MultiSet([d_2420_v0_, d_2420_v0_])
                    nw392_[int(20)] = ((_dafny.MultiSet(d_2422_v2_)).set(d_2420_v0_, default__.abs(((d_2424_v4_)[p1] if (p1) in (d_2424_v4_) else (self).f25)))).intersection(d_2421_v1_)
                    nw392_[int(21)] = d_2421_v1_
                    d_2425_v5_ = nw392_
                    index464_ = default__.safeIndex(69, (d_2425_v5_).length(0))
                    (d_2425_v5_)[index464_] = d_2421_v1_
                    d_2428_v6_: _dafny.Seq
                    d_2428_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yupwkrstt"))
                    index465_ = default__.safeIndex(69, (d_2425_v5_).length(0))
                    rhs431_ = (d_2428_v6_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "in"))))
                    rhs432_ = ((self).f25) + (((self).f25) - (((d_2424_v4_)[p1] if (p1) in (d_2424_v4_) else (self).f25)))
                    rhs433_ = (d_2421_v1_).set(D0_DC2(p1), default__.abs((self).f25))
                    lhs345_ = globalState
                    lhs346_ = d_2425_v5_
                    lhs347_ = default__.safeIndex(69, (d_2425_v5_).length(0))
                    lhs345_.f10 = rhs431_
                    r0 = rhs432_
                    lhs346_[lhs347_] = rhs433_
                    if p1:
                        d_2429_v7_: C1
                        nw393_ = C1()
                        nw393_.ctor__((self).f25, (self).f26)
                        d_2429_v7_ = nw393_
                        d_2430_v8_: _dafny.Set
                        d_2430_v8_ = _dafny.Set({d_2429_v7_})
                        d_2431_v9_: _dafny.Map
                        d_2431_v9_ = _dafny.Map({d_2428_v6_: len(d_2430_v8_)})
                        d_2431_v9_ = d_2431_v9_
                        r0 = len(d_2428_v6_)
                        d_2432_v11_: _dafny.MultiSet
                        d_2432_v11_ = _dafny.MultiSet([(self).f29])
                        d_2433_v13_: _dafny.Map
                        def iife202_():
                            def iife204_():
                                coll88_ = _dafny.Set()
                                compr_88_: str
                                for compr_88_ in (d_2432_v11_).Elements:
                                    d_2436_v12_: str = compr_88_
                                    if (d_2436_v12_) in (d_2432_v11_):
                                        coll88_ = coll88_.union(_dafny.Set([d_2436_v12_]))
                                return _dafny.Set(coll88_)
                            coll86_ = _dafny.Map()
                            def iife203_():
                                coll87_ = _dafny.Set()
                                compr_87_: str
                                for compr_87_ in (d_2432_v11_).Elements:
                                    d_2434_v12_: str = compr_87_
                                    if (d_2434_v12_) in (d_2432_v11_):
                                        coll87_ = coll87_.union(_dafny.Set([d_2434_v12_]))
                                return _dafny.Set(coll87_)
                            compr_86_: str
                            for compr_86_ in (iife203_()
                            ).Elements:
                                d_2435_v10_: str = compr_86_
                                if (d_2435_v10_) in (iife204_()
                                ):
                                    coll86_[d_2435_v10_] = 165
                            return _dafny.Map(coll86_)
                        d_2433_v13_ = _dafny.Map({p1: default__.fm0(len(iife202_()
                        ), globalState)})
                        d_2437_v14_: _dafny.Array
                        nw394_ = _dafny.Array(None, 16)
                        nw394_[int(0)] = (-982) <= ((self).f25)
                        nw394_[int(1)] = p1
                        nw394_[int(2)] = True
                        nw394_[int(3)] = (p1) == (p1)
                        nw394_[int(4)] = not(((d_2433_v13_)[p1] if (p1) in (d_2433_v13_) else p1))
                        nw394_[int(5)] = not(p1)
                        nw394_[int(6)] = p1
                        nw394_[int(7)] = p1
                        nw394_[int(8)] = p1
                        nw394_[int(9)] = p1
                        nw394_[int(10)] = p1
                        nw394_[int(11)] = p1
                        nw394_[int(12)] = p1
                        nw394_[int(13)] = p1
                        nw394_[int(14)] = default__.fm0((self).f25, globalState)
                        nw394_[int(15)] = p1
                        d_2437_v14_ = nw394_
                        index466_ = default__.safeIndex(767, (d_2437_v14_).length(0))
                        (d_2437_v14_)[index466_] = p1
                        index467_ = default__.safeIndex(767, (d_2437_v14_).length(0))
                        (d_2437_v14_)[index467_] = p1
                        d_2438_v15_: _dafny.Set
                        d_2438_v15_ = _dafny.Set({(self).f25, (self).f25})
                        d_2439_v16_: C4
                        nw395_ = C4()
                        nw395_.ctor__((self).f26, len((d_2438_v15_).intersection(d_2438_v15_)), (self).f26)
                        d_2439_v16_ = nw395_
                        d_2437_v14_ = d_2437_v14_
                    elif True:
                        r0 = (self).f25
                        d_2440_v17_: _dafny.Array
                        nw396_ = _dafny.Array(False, 25)
                        d_2440_v17_ = nw396_
                        d_2441_v18_: _dafny.Seq
                        d_2441_v18_ = _dafny.SeqWithoutIsStrInference([-886, (0) - (len(d_2428_v6_)), (self).f25])
                        d_2442_v19_: _dafny.Map
                        d_2442_v19_ = _dafny.Map({(d_2441_v18_) + (d_2441_v18_): d_2428_v6_})
                        d_2443_v20_: D4
                        d_2443_v20_ = D4_DC13((self).f25, p1, default__.fm0((self).f25, globalState), p1)
                        d_2444_v21_: _dafny.Seq
                        d_2444_v21_ = _dafny.SeqWithoutIsStrInference([(d_2443_v20_).cf22])
                        rhs434_ = ((self).f25) + ((self).f25)
                        rhs435_ = len(d_2442_v19_)
                        rhs436_ = (p1) in ((default__.fm1(p1, globalState)) + (d_2444_v21_))
                        rhs437_ = d_2440_v17_
                        lhs348_ = globalState
                        r0 = rhs434_
                        r0 = rhs435_
                        lhs348_.f12 = rhs436_
                        d_2440_v17_ = rhs437_
                        d_2445_v22_: _dafny.Array
                        nw397_ = _dafny.Array(int(0), 13)
                        d_2445_v22_ = nw397_
                        index468_ = default__.safeIndex(532, (d_2445_v22_).length(0))
                        (d_2445_v22_)[index468_] = -11
                        d_2446_v23_: _dafny.Set
                        d_2446_v23_ = _dafny.Set({p1, p1, p1})
                        index469_ = default__.safeIndex(532, (d_2445_v22_).length(0))
                        (d_2445_v22_)[index469_] = len((d_2446_v23_ if p1 else d_2446_v23_))
                        index470_ = default__.safeIndex(553, (d_2440_v17_).length(0))
                        (d_2440_v17_)[index470_] = p1
                        index471_ = default__.safeIndex(553, (d_2440_v17_).length(0))
                        (d_2440_v17_)[index471_] = p1
                        d_2447_v24_: _dafny.Set
                        d_2447_v24_ = _dafny.Set({(0) - ((d_2445_v22_)[default__.safeIndex(532, (d_2445_v22_).length(0))])})
                        index472_ = default__.safeIndex(532, (d_2445_v22_).length(0))
                        rhs438_ = (d_2445_v22_)[default__.safeIndex(532, (d_2445_v22_).length(0))]
                        rhs439_ = (d_2440_v17_)[default__.safeIndex(553, (d_2440_v17_).length(0))]
                        rhs440_ = (len(_dafny.SeqWithoutIsStrInference([(self).f29 for d_2448_i1_ in range(default__.abs(801))]))) in (d_2447_v24_)
                        lhs349_ = d_2445_v22_
                        lhs350_ = default__.safeIndex(532, (d_2445_v22_).length(0))
                        lhs351_ = globalState
                        lhs352_ = globalState
                        lhs349_[lhs350_] = rhs438_
                        lhs351_.f12 = rhs439_
                        lhs352_.f12 = rhs440_
                    d_2449_v25_: _dafny.Array
                    nw398_ = _dafny.Array(int(0), 4)
                    d_2449_v25_ = nw398_
                    d_2450_v26_: _dafny.Array
                    nw399_ = _dafny.Array(None, 11)
                    nw399_[int(0)] = d_2449_v25_
                    nw399_[int(1)] = (d_2449_v25_ if p1 else d_2449_v25_)
                    nw399_[int(2)] = d_2449_v25_
                    nw399_[int(3)] = d_2449_v25_
                    nw399_[int(4)] = d_2449_v25_
                    nw399_[int(5)] = d_2449_v25_
                    nw399_[int(6)] = d_2449_v25_
                    nw399_[int(7)] = d_2449_v25_
                    nw399_[int(8)] = d_2449_v25_
                    nw399_[int(9)] = d_2449_v25_
                    nw399_[int(10)] = (p0)[default__.safeIndex((self).f25, len(p0))]
                    d_2450_v26_ = nw399_
                    d_2450_v26_ = d_2450_v26_
                    pass
            pass
        d_2451_v27_: C6
        nw400_ = C6()
        nw400_.ctor__((self).f25, (self).f26)
        d_2451_v27_ = nw400_
        d_2451_v27_ = d_2451_v27_
        d_2452_v28_: D0
        d_2452_v28_ = D0_DC1(p1, ((_dafny.MultiSet([(self).f25])).cardinality) <= ((self).f25))
        d_2453_v29_: _dafny.Seq
        d_2453_v29_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_2454_v30_: _dafny.Set
        d_2454_v30_ = _dafny.Set({len(d_2453_v29_)})
        d_2455_v31_: _dafny.Set
        d_2455_v31_ = _dafny.Set({(self).f25, len(d_2454_v30_)})
        d_2456_v32_: D4
        d_2456_v32_ = D4_DC14(D4_DC11(d_2455_v31_))
        d_2457_v33_: D4
        d_2457_v33_ = D4_DC12()
        pat_let_tv49_ = d_2452_v28_
        pat_let_tv50_ = d_2452_v28_
        pat_let_tv51_ = d_2452_v28_
        pat_let_tv52_ = d_2457_v33_
        def lambda153_(source31_):
            if source31_.is_DC12:
                return pat_let_tv49_
            elif source31_.is_DC13:
                d_2458___mcc_h0_ = source31_.cf19
                d_2459___mcc_h1_ = source31_.cf20
                d_2460___mcc_h2_ = source31_.cf21
                d_2461___mcc_h3_ = source31_.cf22
                d_2462_cf22_ = d_2461___mcc_h3_
                d_2463_cf21_ = d_2460___mcc_h2_
                d_2464_cf20_ = d_2459___mcc_h1_
                d_2465_cf19_ = d_2458___mcc_h0_
                return D0_DC1(d_2464_cf20_, d_2464_cf20_)
            elif source31_.is_DC11:
                d_2466___mcc_h4_ = source31_.cf18
                d_2467_cf18_ = d_2466___mcc_h4_
                return pat_let_tv50_
            elif True:
                d_2468___mcc_h5_ = source31_.cf23
                d_2469_cf23_ = d_2468___mcc_h5_
                return pat_let_tv51_

        def iife205_(_pat_let58_0):
            def iife206_(d_2470_dt__update__tmp_h1_):
                def iife207_(_pat_let59_0):
                    def iife208_(d_2471_dt__update_hcf23_h0_):
                        return D4_DC14(d_2471_dt__update_hcf23_h0_)
                    return iife208_(_pat_let59_0)
                return iife207_(pat_let_tv52_)
            return iife206_(_pat_let58_0)
        d_2452_v28_ = lambda153_(iife205_(d_2456_v32_))
        d_2472_v34_: _dafny.MultiSet
        d_2472_v34_ = _dafny.MultiSet([((self).f25) >= ((self).f25), p1, p1])
        d_2473_v35_: _dafny.Seq
        d_2473_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
        d_2474_v36_: _dafny.Map
        d_2474_v36_ = _dafny.Map({p1: p1})
        d_2472_v34_ = default__.fm50((self).f25, d_2473_v35_, ((d_2474_v36_)[True] if (True) in (d_2474_v36_) else p1), (self).f25, globalState)
        d_2475_v37_: _dafny.Array
        nw401_ = _dafny.Array(False, 2)
        d_2475_v37_ = nw401_
        d_2476_v38_: _dafny.MultiSet
        d_2476_v38_ = _dafny.MultiSet([(self).f25, (self).f25, (self).f25])
        index473_ = default__.safeIndex(526, (d_2475_v37_).length(0))
        (d_2475_v37_)[index473_] = (d_2476_v38_).issubset(_dafny.MultiSet([(self).f25]))
        index474_ = default__.safeIndex(526, (d_2475_v37_).length(0))
        (d_2475_v37_)[index474_] = p1
        hi14_ = (0) - ((self).f25)
        for d_2477_i2_ in range((self).f25, hi14_):
            d_2478_v39_: _dafny.Array
            def lambda154_(d_2479_v35_):
                def lambda155_(d_2480_i3_):
                    return d_2479_v35_

                return lambda155_

            init86_ = lambda154_(d_2473_v35_)
            nw402_ = _dafny.Array(None, 23)
            for i0_86_ in range(nw402_.length(0)):
                nw402_[i0_86_] = init86_(i0_86_)
            d_2478_v39_ = nw402_
            rhs441_ = default__.fm18(globalState)
            rhs442_ = d_2478_v39_
            r0 = rhs441_
            d_2478_v39_ = rhs442_
            d_2481_v40_: _dafny.Seq
            d_2481_v40_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isuhwhf"))), (self).f25])
            d_2482_v41_: _dafny.Set
            d_2482_v41_ = _dafny.Set({d_2481_v40_})
            r0 = default__.safeModuloInt(len(d_2482_v41_), (self).f25)
            d_2483_v42_: _dafny.Seq
            d_2484_v43_: _dafny.MultiSet
            d_2485_v44_: bool
            out43_: _dafny.Seq
            out44_: _dafny.MultiSet
            out45_: bool
            out43_, out44_, out45_ = (d_2451_v27_).m12(globalState)
            d_2483_v42_ = out43_
            d_2484_v43_ = out44_
            d_2485_v44_ = out45_
            d_2486_v45_: C5
            nw403_ = C5()
            nw403_.ctor__((d_2475_v37_)[default__.safeIndex(526, (d_2475_v37_).length(0))])
            d_2486_v45_ = nw403_
        r0 = default__.fm18(globalState)
        return r0

    @property
    def f29(self):
        return self._f29

class C11:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    def ctor__(self):
        pass
        pass

    def fm9(self, p0, p1, globalState):
        return _dafny.CodePoint('l')

    def fm10(self, globalState):
        return D1_DC5()

    def m6(self, p0, p1, globalState):
        r0: bool = False
        d_2487_i0_: int
        d_2487_i0_ = 0
        with _dafny.label("20"):
            while False:
                with _dafny.c_label("20"):
                    if (d_2487_i0_) >= (100):
                        raise _dafny.Break("20")
                    d_2487_i0_ = (d_2487_i0_) + (1)
                    d_2488_v0_: bool
                    d_2488_v0_ = False
                    d_2489_v1_: _dafny.Map
                    d_2489_v1_ = _dafny.Map({d_2488_v0_: d_2488_v0_})
                    d_2490_v2_: _dafny.Map
                    d_2490_v2_ = _dafny.Map({((d_2489_v1_)[d_2488_v0_] if (d_2488_v0_) in (d_2489_v1_) else d_2488_v0_): d_2488_v0_})
                    d_2490_v2_ = (d_2490_v2_).set(d_2488_v0_, d_2488_v0_)
                    d_2491_v3_: _dafny.Seq
                    d_2491_v3_ = _dafny.SeqWithoutIsStrInference([(not(d_2488_v0_) if d_2488_v0_ else d_2488_v0_)])
                    d_2492_v4_: int
                    d_2492_v4_ = -601
                    d_2491_v3_ = (p0).set(default__.safeIndex(d_2492_v4_, len(p0)), d_2488_v0_)
                    d_2493_v5_: _dafny.Array
                    nw404_ = _dafny.Array(_dafny.Seq({}), 10)
                    d_2493_v5_ = nw404_
                    d_2494_v6_: _dafny.Array
                    nw405_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                    d_2494_v6_ = nw405_
                    d_2495_v7_: T0
                    nw406_ = C4()
                    nw406_.ctor__(d_2494_v6_, d_2492_v4_, d_2494_v6_)
                    d_2495_v7_ = nw406_
                    d_2496_v8_: _dafny.Seq
                    d_2496_v8_ = _dafny.SeqWithoutIsStrInference([d_2495_v7_])
                    index475_ = default__.safeIndex(1, (d_2493_v5_).length(0))
                    (d_2493_v5_)[index475_] = (d_2496_v8_) + ((d_2496_v8_).set(default__.safeIndex(d_2492_v4_, len(d_2496_v8_)), d_2495_v7_))
                    index476_ = default__.safeIndex(1, (d_2493_v5_).length(0))
                    (d_2493_v5_)[index476_] = (_dafny.SeqWithoutIsStrInference([d_2495_v7_])).set(default__.safeIndex((d_2495_v7_).f25, len(_dafny.SeqWithoutIsStrInference([d_2495_v7_]))), d_2495_v7_)
                    d_2497_v9_: _dafny.Map
                    d_2497_v9_ = _dafny.Map({(d_2495_v7_).f25: d_2488_v0_})
                    d_2498_v10_: _dafny.Set
                    d_2498_v10_ = _dafny.Set({d_2488_v0_})
                    d_2499_v11_: _dafny.Map
                    d_2499_v11_ = _dafny.Map({d_2498_v10_: d_2488_v0_})
                    d_2500_v12_: _dafny.Array
                    nw407_ = _dafny.Array(None, 19)
                    nw407_[int(0)] = d_2492_v4_
                    nw407_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pb")))
                    nw407_[int(2)] = default__.fm44(d_2488_v0_, d_2488_v0_, d_2488_v0_, (d_2495_v7_).f25, globalState)
                    nw407_[int(3)] = default__.safeDivisionInt((d_2495_v7_).f25, d_2492_v4_)
                    nw407_[int(4)] = (d_2495_v7_).f25
                    nw407_[int(5)] = ((0) - ((d_2495_v7_).f25)) + ((d_2495_v7_).f25)
                    nw407_[int(6)] = (d_2495_v7_).f25
                    nw407_[int(7)] = len(d_2497_v9_)
                    nw407_[int(8)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xeoqcuqdk")))
                    nw407_[int(9)] = (d_2495_v7_).f25
                    nw407_[int(10)] = (d_2495_v7_).f25
                    nw407_[int(11)] = d_2492_v4_
                    nw407_[int(12)] = default__.fm18(globalState)
                    nw407_[int(13)] = (378) - (d_2492_v4_)
                    nw407_[int(14)] = d_2492_v4_
                    nw407_[int(15)] = len(d_2499_v11_)
                    nw407_[int(16)] = (0) - (d_2492_v4_)
                    nw407_[int(17)] = (d_2492_v4_) - (d_2492_v4_)
                    nw407_[int(18)] = (d_2495_v7_).f25
                    d_2500_v12_ = nw407_
                    d_2500_v12_ = d_2500_v12_
                    pass
            pass
        d_2501_v13_: int
        d_2501_v13_ = -572
        d_2502_v14_: _dafny.Seq
        d_2502_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkg"))
        d_2501_v13_ = len(d_2502_v14_)
        d_2503_v15_: bool
        d_2503_v15_ = True
        d_2504_v16_: _dafny.Array
        nw408_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_2504_v16_ = nw408_
        d_2505_v17_: C3
        nw409_ = C3()
        nw409_.ctor__(d_2501_v13_, False, 637, (d_2504_v16_ if d_2503_v15_ else d_2504_v16_))
        d_2505_v17_ = nw409_
        d_2506_v18_: C1
        nw410_ = C1()
        nw410_.ctor__(d_2501_v13_, d_2504_v16_)
        d_2506_v18_ = nw410_
        d_2507_v19_: _dafny.Map
        d_2507_v19_ = _dafny.Map({d_2506_v18_: default__.safeDivisionInt((d_2505_v17_).f37, (d_2505_v17_).f37)})
        d_2507_v19_ = (d_2507_v19_).set(d_2506_v18_, 610)
        if d_2503_v15_:
            if (d_2505_v17_).f38:
                d_2503_v15_ = (p1) in (d_2502_v14_)
                d_2508_v20_: _dafny.Array
                nw411_ = _dafny.Array(int(0), 17)
                d_2508_v20_ = nw411_
                d_2509_v21_: _dafny.Set
                d_2509_v21_ = _dafny.Set({d_2508_v20_, d_2508_v20_, d_2508_v20_, d_2508_v20_})
                d_2510_v22_: _dafny.Map
                d_2510_v22_ = _dafny.Map({(d_2509_v21_).intersection(d_2509_v21_): (d_2505_v17_).f37})
                d_2510_v22_ = (d_2510_v22_).set(d_2509_v21_, default__.safeModuloInt(d_2501_v13_, d_2501_v13_))
                d_2501_v13_ = (default__.safeModuloInt((d_2505_v17_).f37, (d_2505_v17_).f37)) - (len((d_2502_v14_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p1 for d_2511_i1_ in range(default__.abs(956))])), len(d_2502_v14_)), p1)))
                d_2512_v23_: _dafny.MultiSet
                d_2512_v23_ = _dafny.MultiSet([(0) - (d_2501_v13_), (d_2505_v17_).f37, 254])
                d_2501_v13_ = ((d_2505_v17_).f37) + ((default__.fm44(d_2503_v15_, False, (d_2505_v17_).f38, (d_2505_v17_).f37, globalState)) * ((d_2512_v23_).cardinality))
                d_2513_v24_: D0
                d_2513_v24_ = D0_DC0((d_2505_v17_).f38)
                d_2514_v25_: _dafny.Map
                d_2514_v25_ = _dafny.Map({(d_2505_v17_).fm38(d_2513_v24_, globalState): (d_2505_v17_).f38})
                index477_ = default__.safeIndex(544, (d_2508_v20_).length(0))
                (d_2508_v20_)[index477_] = len(d_2514_v25_)
                index478_ = default__.safeIndex(544, (d_2508_v20_).length(0))
                (d_2508_v20_)[index478_] = default__.safeDivisionInt(default__.fm18(globalState), len(_dafny.Set({default__.fm44((d_2505_v17_).f38, d_2503_v15_, (d_2505_v17_).f38, (d_2512_v23_).cardinality, globalState), (d_2505_v17_).f37, (d_2505_v17_).f37})))
            elif True:
                d_2515_v26_: _dafny.Array
                def lambda156_(d_2516_v17_):
                    def lambda157_(d_2517_i2_):
                        return (d_2516_v17_).f38

                    return lambda157_

                init87_ = lambda156_(d_2505_v17_)
                nw412_ = _dafny.Array(None, 8)
                for i0_87_ in range(nw412_.length(0)):
                    nw412_[i0_87_] = init87_(i0_87_)
                d_2515_v26_ = nw412_
                index479_ = default__.safeIndex(878, (d_2515_v26_).length(0))
                (d_2515_v26_)[index479_] = (d_2505_v17_).f38
                index480_ = default__.safeIndex(878, (d_2515_v26_).length(0))
                (d_2515_v26_)[index480_] = not((d_2505_v17_).f38)
                d_2501_v13_ = (0) - ((d_2505_v17_).f37)
                d_2518_v27_: _dafny.Set
                d_2518_v27_ = _dafny.Set({209})
                d_2519_v28_: _dafny.Array
                nw413_ = _dafny.Array(None, 9)
                nw413_[int(0)] = (d_2505_v17_).f37
                nw413_[int(1)] = (d_2505_v17_).f37
                nw413_[int(2)] = (d_2505_v17_).f37
                nw413_[int(3)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2520_i3_ in range(default__.abs(307))]))
                nw413_[int(4)] = (d_2505_v17_).f37
                nw413_[int(5)] = (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbievbw"))).set(default__.safeIndex((d_2505_v17_).f37, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbievbw")))), p1)))
                nw413_[int(6)] = d_2501_v13_
                nw413_[int(7)] = len(d_2518_v27_)
                nw413_[int(8)] = 429
                d_2519_v28_ = nw413_
                d_2521_v29_: _dafny.Map
                d_2521_v29_ = _dafny.Map({(d_2505_v17_).f38: (d_2505_v17_).f37})
                d_2522_v30_: C7
                nw414_ = C7()
                nw414_.ctor__(len(_dafny.Map({d_2519_v28_: len((d_2521_v29_).set((d_2515_v26_)[default__.safeIndex(878, (d_2515_v26_).length(0))], len(d_2502_v14_)))})))
                d_2522_v30_ = nw414_
                d_2523_v31_: _dafny.Set
                d_2523_v31_ = _dafny.Set({d_2515_v26_})
                d_2524_v32_: D3
                d_2524_v32_ = D3_DC9(d_2523_v31_)
                d_2525_v33_: _dafny.Seq
                d_2525_v33_ = _dafny.SeqWithoutIsStrInference([(d_2505_v17_).f37])
                d_2526_v34_: C4
                nw415_ = C4()
                nw415_.ctor__(d_2504_v16_, (d_2522_v30_).f34, d_2504_v16_)
                d_2526_v34_ = nw415_
                d_2527_v35_: _dafny.Map
                d_2527_v35_ = _dafny.Map({d_2525_v33_: d_2526_v34_})
                rhs443_ = d_2522_v30_
                rhs444_ = not ((d_2505_v17_).f38) or (d_2503_v15_)
                rhs445_ = (d_2505_v17_).f37
                rhs446_ = default__.safeDivisionInt((0) - ((len(d_2502_v14_)) * (len(d_2521_v29_))), len(d_2527_v35_))
                rhs447_ = d_2524_v32_
                d_2522_v30_ = rhs443_
                d_2503_v15_ = rhs444_
                d_2501_v13_ = rhs445_
                d_2501_v13_ = rhs446_
                d_2524_v32_ = rhs447_
                d_2501_v13_ = d_2501_v13_
                d_2528_v36_: _dafny.Map
                d_2528_v36_ = _dafny.Map({((d_2522_v30_).f34) - (len(d_2502_v14_)): d_2515_v26_})
                d_2529_v37_: _dafny.Array
                nw416_ = _dafny.Array(D6.default()(), 20)
                d_2529_v37_ = nw416_
                d_2530_v38_: _dafny.Map
                d_2530_v38_ = _dafny.Map({d_2515_v26_: 832})
                d_2531_v39_: D6
                d_2531_v39_ = D6_DC18(d_2530_v38_)
                index481_ = default__.safeIndex(432, (d_2529_v37_).length(0))
                (d_2529_v37_)[index481_] = d_2531_v39_
                index482_ = default__.safeIndex(878, (d_2515_v26_).length(0))
                index483_ = default__.safeIndex(432, (d_2529_v37_).length(0))
                rhs448_ = (d_2505_v17_).f38
                rhs449_ = ((0) - ((d_2505_v17_).f37)) <= (d_2501_v13_)
                rhs450_ = _dafny.Map({(d_2505_v17_).f37: d_2515_v26_})
                rhs451_ = d_2531_v39_
                lhs353_ = d_2515_v26_
                lhs354_ = default__.safeIndex(878, (d_2515_v26_).length(0))
                lhs355_ = d_2529_v37_
                lhs356_ = default__.safeIndex(432, (d_2529_v37_).length(0))
                r0 = rhs448_
                lhs353_[lhs354_] = rhs449_
                d_2528_v36_ = rhs450_
                lhs355_[lhs356_] = rhs451_
            d_2532_v40_: _dafny.Array
            def lambda158_(d_2533_v15_):
                def lambda159_(d_2534_i4_):
                    return default__.safeDivisionInt(d_2534_i4_, len(_dafny.Map({d_2533_v15_: 872})))

                return lambda159_

            init88_ = lambda158_(d_2503_v15_)
            nw417_ = _dafny.Array(None, 11)
            for i0_88_ in range(nw417_.length(0)):
                nw417_[i0_88_] = init88_(i0_88_)
            d_2532_v40_ = nw417_
            index484_ = default__.safeIndex(525, (d_2532_v40_).length(0))
            (d_2532_v40_)[index484_] = (0) - (((d_2505_v17_).f37) + ((d_2505_v17_).f37))
            index485_ = default__.safeIndex(525, (d_2532_v40_).length(0))
            (d_2532_v40_)[index485_] = (((0) - ((d_2505_v17_).f37)) * (d_2501_v13_)) * (default__.fm44((d_2505_v17_).f38, (d_2505_v17_).f38, (d_2505_v17_).f38, (d_2505_v17_).f37, globalState))
            (globalState).f12 = (d_2505_v17_).f38
            if (d_2505_v17_).f38:
                d_2535_v41_: _dafny.Map
                d_2535_v41_ = _dafny.Map({(d_2505_v17_).f38: (d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))]})
                d_2535_v41_ = (d_2535_v41_).set((d_2505_v17_).f38, (d_2505_v17_).f37)
                d_2536_v42_: _dafny.Seq
                d_2536_v42_ = _dafny.SeqWithoutIsStrInference([d_2532_v40_])
                d_2532_v40_ = (d_2536_v42_)[default__.safeIndex((d_2505_v17_).f37, len(d_2536_v42_))]
                d_2501_v13_ = (d_2505_v17_).f37
                d_2537_v43_: _dafny.Array
                nw418_ = _dafny.Array(False, 18)
                d_2537_v43_ = nw418_
                d_2538_v44_: D3
                d_2538_v44_ = D3_DC9(_dafny.Set({d_2537_v43_}))
                d_2539_v45_: _dafny.Map
                d_2539_v45_ = _dafny.Map({(d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))]: d_2538_v44_})
                d_2540_v46_: D19
                d_2540_v46_ = D19_DC46((d_2539_v45_) | (d_2539_v45_))
                d_2540_v46_ = d_2540_v46_
                d_2541_v47_: D1
                d_2541_v47_ = D1_DC4(p0)
                pat_let_tv53_ = p0
                d_2542_v48_: _dafny.Seq
                def iife209_(_pat_let60_0):
                    def iife210_(d_2543_dt__update__tmp_h0_):
                        def iife211_(_pat_let61_0):
                            def iife212_(d_2544_dt__update_hcf6_h0_):
                                return D1_DC4(d_2544_dt__update_hcf6_h0_)
                            return iife212_(_pat_let61_0)
                        return iife211_(pat_let_tv53_)
                    return iife210_(_pat_let60_0)
                d_2542_v48_ = _dafny.SeqWithoutIsStrInference([d_2541_v47_, iife209_(d_2541_v47_), d_2541_v47_])
                d_2545_v49_: _dafny.Array
                nw419_ = _dafny.Array(None, 26)
                nw419_[int(0)] = (d_2542_v48_).set(default__.safeIndex((d_2505_v17_).f37, len(d_2542_v48_)), D1_DC4(p0))
                nw419_[int(1)] = d_2542_v48_
                nw419_[int(2)] = d_2542_v48_
                nw419_[int(3)] = d_2542_v48_
                nw419_[int(4)] = d_2542_v48_
                nw419_[int(5)] = d_2542_v48_
                nw419_[int(6)] = _dafny.SeqWithoutIsStrInference([d_2541_v47_, D1_DC4(p0)])
                nw419_[int(7)] = d_2542_v48_
                nw419_[int(8)] = default__.fm68((d_2505_v17_).f38, (d_2505_v17_).f38, d_2502_v14_, (self).fm9(d_2503_v15_, len(_dafny.Map({(d_2505_v17_).f38: d_2503_v15_})), globalState), globalState)
                nw419_[int(9)] = d_2542_v48_
                nw419_[int(10)] = (d_2542_v48_) + ((d_2542_v48_).set(default__.safeIndex((d_2505_v17_).f37, len(d_2542_v48_)), d_2541_v47_))
                nw419_[int(11)] = _dafny.SeqWithoutIsStrInference([d_2541_v47_])
                nw419_[int(12)] = d_2542_v48_
                nw419_[int(13)] = d_2542_v48_
                nw419_[int(14)] = _dafny.SeqWithoutIsStrInference([d_2541_v47_ for d_2546_i5_ in range(default__.abs(869))])
                nw419_[int(15)] = (d_2542_v48_) + (d_2542_v48_)
                nw419_[int(16)] = d_2542_v48_
                nw419_[int(17)] = d_2542_v48_
                nw419_[int(18)] = (d_2542_v48_).set(default__.safeIndex((d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))], len(d_2542_v48_)), d_2541_v47_)
                nw419_[int(19)] = d_2542_v48_
                nw419_[int(20)] = d_2542_v48_
                nw419_[int(21)] = _dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.SeqWithoutIsStrInference([d_2503_v15_])), d_2541_v47_])
                nw419_[int(22)] = d_2542_v48_
                nw419_[int(23)] = ((d_2542_v48_).set(default__.safeIndex((d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))], len(d_2542_v48_)), d_2541_v47_)) + (d_2542_v48_)
                nw419_[int(24)] = ((_dafny.SeqWithoutIsStrInference([d_2541_v47_])).set(default__.safeIndex((d_2505_v17_).f37, len(_dafny.SeqWithoutIsStrInference([d_2541_v47_]))), d_2541_v47_)).set(default__.safeIndex(len(p0), len((_dafny.SeqWithoutIsStrInference([d_2541_v47_])).set(default__.safeIndex((d_2505_v17_).f37, len(_dafny.SeqWithoutIsStrInference([d_2541_v47_]))), d_2541_v47_))), d_2541_v47_)
                nw419_[int(25)] = _dafny.SeqWithoutIsStrInference([d_2541_v47_])
                d_2545_v49_ = nw419_
                d_2545_v49_ = d_2545_v49_
            elif True:
                d_2547_v50_: _dafny.Array
                def lambda160_(d_2548_v17_):
                    def lambda161_(d_2549_i6_):
                        return _dafny.Set({(d_2548_v17_).f38, (d_2548_v17_).f38, False})

                    return lambda161_

                init89_ = lambda160_(d_2505_v17_)
                nw420_ = _dafny.Array(None, 24)
                for i0_89_ in range(nw420_.length(0)):
                    nw420_[i0_89_] = init89_(i0_89_)
                d_2547_v50_ = nw420_
                d_2547_v50_ = d_2547_v50_
                d_2550_v51_: _dafny.Map
                d_2550_v51_ = _dafny.Map({d_2503_v15_: (d_2505_v17_).f38})
                index486_ = default__.safeIndex(525, (d_2532_v40_).length(0))
                (d_2532_v40_)[index486_] = ((d_2505_v17_).f37 if ((d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))]) > ((d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))]) else default__.safeDivisionInt(len(d_2550_v51_), (d_2505_v17_).f37))
                d_2551_v52_: _dafny.Array
                nw421_ = _dafny.Array(None, 26)
                nw421_[int(0)] = (d_2505_v17_).f38
                nw421_[int(1)] = (d_2505_v17_).f38
                nw421_[int(2)] = (d_2505_v17_).f38
                nw421_[int(3)] = (d_2505_v17_).f38
                nw421_[int(4)] = d_2503_v15_
                nw421_[int(5)] = (d_2505_v17_).f38
                nw421_[int(6)] = (d_2505_v17_).f38
                nw421_[int(7)] = (d_2505_v17_).f38
                nw421_[int(8)] = (d_2505_v17_).f38
                nw421_[int(9)] = (d_2505_v17_).f38
                nw421_[int(10)] = (d_2505_v17_).f38
                nw421_[int(11)] = (d_2505_v17_).f38
                nw421_[int(12)] = (p0)[default__.safeIndex((default__.fm66(globalState)).cf71, len(p0))]
                nw421_[int(13)] = (d_2505_v17_).f38
                nw421_[int(14)] = (d_2505_v17_).f38
                nw421_[int(15)] = (d_2505_v17_).f38
                nw421_[int(16)] = (d_2505_v17_).f38
                nw421_[int(17)] = (d_2505_v17_).f38
                nw421_[int(18)] = (d_2505_v17_).f38
                nw421_[int(19)] = (d_2505_v17_).f38
                nw421_[int(20)] = (d_2505_v17_).f38
                nw421_[int(21)] = (d_2505_v17_).f38
                nw421_[int(22)] = (d_2505_v17_).f38
                nw421_[int(23)] = not(not(not(True)))
                nw421_[int(24)] = d_2503_v15_
                nw421_[int(25)] = (d_2505_v17_).f38
                d_2551_v52_ = nw421_
                d_2552_v53_: _dafny.Seq
                d_2552_v53_ = _dafny.SeqWithoutIsStrInference([d_2551_v52_, d_2551_v52_, d_2551_v52_])
                d_2553_v54_: _dafny.Array
                nw422_ = _dafny.Array(None, 26)
                nw422_[int(0)] = d_2551_v52_
                nw422_[int(1)] = (d_2552_v53_)[default__.safeIndex(590, len(d_2552_v53_))]
                nw422_[int(2)] = d_2551_v52_
                nw422_[int(3)] = d_2551_v52_
                nw422_[int(4)] = d_2551_v52_
                nw422_[int(5)] = d_2551_v52_
                nw422_[int(6)] = d_2551_v52_
                nw422_[int(7)] = d_2551_v52_
                nw422_[int(8)] = d_2551_v52_
                nw422_[int(9)] = d_2551_v52_
                nw422_[int(10)] = (d_2551_v52_ if (d_2505_v17_).f38 else d_2551_v52_)
                nw422_[int(11)] = d_2551_v52_
                nw422_[int(12)] = d_2551_v52_
                nw422_[int(13)] = d_2551_v52_
                nw422_[int(14)] = d_2551_v52_
                nw422_[int(15)] = d_2551_v52_
                nw422_[int(16)] = (d_2551_v52_ if False else d_2551_v52_)
                nw422_[int(17)] = d_2551_v52_
                nw422_[int(18)] = d_2551_v52_
                nw422_[int(19)] = d_2551_v52_
                nw422_[int(20)] = d_2551_v52_
                nw422_[int(21)] = d_2551_v52_
                nw422_[int(22)] = d_2551_v52_
                nw422_[int(23)] = d_2551_v52_
                nw422_[int(24)] = d_2551_v52_
                nw422_[int(25)] = d_2551_v52_
                d_2553_v54_ = nw422_
                d_2554_v55_: _dafny.MultiSet
                d_2554_v55_ = _dafny.MultiSet([(d_2505_v17_).f38, False])
                d_2555_v56_: D20
                d_2555_v56_ = D20_DC49(d_2553_v54_)
                index487_ = default__.safeIndex(525, (d_2532_v40_).length(0))
                rhs452_ = (D14_DC38((d_2532_v40_)[default__.safeIndex(525, (d_2532_v40_).length(0))])).cf58
                rhs453_ = (d_2555_v56_).cf72
                rhs454_ = d_2501_v13_
                rhs455_ = d_2554_v55_
                lhs357_ = d_2532_v40_
                lhs358_ = default__.safeIndex(525, (d_2532_v40_).length(0))
                d_2501_v13_ = rhs452_
                d_2553_v54_ = rhs453_
                lhs357_[lhs358_] = rhs454_
                d_2554_v55_ = rhs455_
                d_2556_v57_: _dafny.Seq
                d_2556_v57_ = _dafny.SeqWithoutIsStrInference([d_2503_v15_])
                d_2556_v57_ = (d_2556_v57_) + ((d_2556_v57_).set(default__.safeIndex((d_2505_v17_).f37, len(d_2556_v57_)), d_2503_v15_))
                d_2557_v58_: C5
                nw423_ = C5()
                nw423_.ctor__((d_2505_v17_).f38)
                d_2557_v58_ = nw423_
            d_2532_v40_ = (d_2532_v40_ if (d_2505_v17_).f38 else d_2532_v40_)
        elif True:
            if ((0) - (default__.safeDivisionInt(d_2501_v13_, (d_2505_v17_).f37))) != ((d_2505_v17_).f37):
                d_2558_v59_: _dafny.Set
                d_2558_v59_ = _dafny.Set({(d_2505_v17_).f38, (d_2505_v17_).f38, d_2503_v15_, default__.fm0((d_2505_v17_).f37, globalState), (d_2505_v17_).f38})
                d_2558_v59_ = d_2558_v59_
                d_2559_v60_: _dafny.Map
                d_2559_v60_ = _dafny.Map({((0) - ((d_2505_v17_).f37)) + ((_dafny.MultiSet([len((d_2502_v14_).set(default__.safeIndex((d_2505_v17_).f37, len(d_2502_v14_)), p1)), (d_2505_v17_).f37])).cardinality): True})
                d_2559_v60_ = (d_2559_v60_).set((d_2505_v17_).f37, d_2503_v15_)
                d_2560_v61_: _dafny.Map
                d_2560_v61_ = _dafny.Map({d_2501_v13_: default__.safeDivisionInt((d_2505_v17_).f37, (d_2505_v17_).f37)})
                d_2560_v61_ = (d_2560_v61_).set(d_2501_v13_, 439)
                (globalState).f12 = (p1) not in (d_2502_v14_)
                d_2561_v62_: D0
                d_2561_v62_ = D0_DC0((d_2505_v17_).f38)
                d_2562_v63_: _dafny.Set
                d_2562_v63_ = _dafny.Set({d_2561_v62_, d_2561_v62_})
                d_2563_v64_: _dafny.Array
                nw424_ = _dafny.Array(None, 4)
                nw424_[int(0)] = (d_2505_v17_).f37
                nw424_[int(1)] = (d_2505_v17_).f37
                nw424_[int(2)] = len(d_2562_v63_)
                nw424_[int(3)] = (0) - ((d_2505_v17_).f37)
                d_2563_v64_ = nw424_
                d_2564_v65_: _dafny.Map
                d_2564_v65_ = _dafny.Map({(0) - (-64): d_2563_v64_})
                d_2565_v66_: _dafny.MultiSet
                d_2565_v66_ = _dafny.MultiSet([_dafny.Map({-182: d_2563_v64_}), d_2564_v65_, d_2564_v65_, _dafny.Map({(d_2505_v17_).f37: d_2563_v64_})])
                d_2501_v13_ = ((d_2565_v66_)[_dafny.Map({d_2501_v13_: d_2563_v64_})] if (_dafny.Map({d_2501_v13_: d_2563_v64_})) in (d_2565_v66_) else 195)
            elif True:
                d_2566_v67_: _dafny.Set
                d_2566_v67_ = _dafny.Set({d_2503_v15_, (d_2505_v17_).f38})
                d_2567_v68_: D5
                d_2567_v68_ = D5_DC16((d_2566_v67_).isdisjoint(d_2566_v67_))
                d_2567_v68_ = D5_DC16((d_2505_v17_).f38)
                d_2503_v15_ = (d_2505_v17_).f38
                d_2568_v69_: _dafny.Set
                d_2568_v69_ = _dafny.Set({(d_2505_v17_).f37})
                rhs456_ = _dafny.Set({(d_2505_v17_).f37, (d_2505_v17_).f37, (d_2505_v17_).f37})
                rhs457_ = (0) - (d_2501_v13_)
                rhs458_ = (d_2505_v17_).f38
                d_2568_v69_ = rhs456_
                d_2501_v13_ = rhs457_
                r0 = rhs458_
                d_2501_v13_ = (0) - ((0) - ((d_2505_v17_).f37))
                d_2569_v70_: _dafny.Array
                def lambda162_(d_2570_v15_):
                    def lambda163_(d_2571_i7_):
                        return d_2570_v15_

                    return lambda163_

                init90_ = lambda162_(d_2503_v15_)
                nw425_ = _dafny.Array(None, 25)
                for i0_90_ in range(nw425_.length(0)):
                    nw425_[i0_90_] = init90_(i0_90_)
                d_2569_v70_ = nw425_
                d_2572_v71_: _dafny.Seq
                d_2572_v71_ = _dafny.SeqWithoutIsStrInference([(d_2505_v17_).f37])
                d_2573_v72_: _dafny.MultiSet
                d_2573_v72_ = _dafny.MultiSet([(d_2505_v17_).f38])
                d_2574_v73_: _dafny.MultiSet
                d_2574_v73_ = _dafny.MultiSet([(D3_DC10(d_2503_v15_, d_2569_v70_, (d_2505_v17_).f38, (d_2572_v71_)[default__.safeIndex((d_2573_v72_).cardinality, len(d_2572_v71_))], len(d_2566_v67_))).cf13])
                d_2575_v74_: _dafny.Seq
                d_2575_v74_ = _dafny.SeqWithoutIsStrInference([d_2573_v72_, _dafny.MultiSet([(d_2505_v17_).f38])])
                d_2576_v75_: _dafny.Map
                d_2576_v75_ = _dafny.Map({d_2574_v73_: len((default__.fm69((d_2505_v17_).f38, globalState)) + (d_2575_v74_))})
                d_2577_v76_: D0
                d_2577_v76_ = D0_DC0(not((d_2505_v17_).f38))
                d_2578_v77_: _dafny.Map
                d_2578_v77_ = _dafny.Map({d_2577_v76_: d_2503_v15_})
                rhs459_ = (d_2576_v75_) | ((_dafny.Map({d_2573_v72_: (d_2505_v17_).f37})) | (d_2576_v75_))
                rhs460_ = d_2569_v70_
                rhs461_ = default__.fm21((_dafny.Map({d_2577_v76_: True})) | (d_2578_v77_), d_2501_v13_, not((d_2505_v17_).f38), d_2502_v14_, globalState)
                rhs462_ = len(d_2502_v14_)
                rhs463_ = d_2503_v15_
                lhs359_ = globalState
                d_2576_v75_ = rhs459_
                d_2569_v70_ = rhs460_
                d_2501_v13_ = rhs461_
                d_2501_v13_ = rhs462_
                lhs359_.f12 = rhs463_
            (globalState).f23 = p1
            (globalState).f12 = (d_2505_v17_).f38
            d_2501_v13_ = ((d_2505_v17_).f37) * (default__.fm18(globalState))
            d_2501_v13_ = (d_2505_v17_).f37
        d_2579_v78_: D1
        d_2579_v78_ = D1_DC4(p0)
        pat_let_tv54_ = d_2503_v15_
        pat_let_tv55_ = d_2503_v15_
        pat_let_tv56_ = d_2503_v15_
        def lambda164_(source32_):
            if source32_.is_DC5:
                return pat_let_tv54_
            elif source32_.is_DC6:
                d_2580___mcc_h0_ = source32_.cf7
                d_2581___mcc_h1_ = source32_.cf8
                d_2582___mcc_h2_ = source32_.cf9
                d_2583_cf9_ = d_2582___mcc_h2_
                d_2584_cf8_ = d_2581___mcc_h1_
                d_2585_cf7_ = d_2580___mcc_h0_
                return (False) and (d_2584_cf8_)
            elif source32_.is_DC4:
                d_2586___mcc_h3_ = source32_.cf6
                d_2587_cf6_ = d_2586___mcc_h3_
                return pat_let_tv55_
            elif True:
                d_2588___mcc_h4_ = source32_.cf10
                d_2589_cf10_ = d_2588___mcc_h4_
                return pat_let_tv56_

        if lambda164_(d_2579_v78_):
            d_2590_v79_: C0
            nw426_ = C0()
            nw426_.ctor__()
            d_2590_v79_ = nw426_
            d_2591_v80_: _dafny.Seq
            d_2591_v80_ = _dafny.SeqWithoutIsStrInference([d_2590_v79_])
            d_2592_v81_: _dafny.Array
            nw427_ = _dafny.Array(None, 26)
            nw427_[int(0)] = d_2590_v79_
            nw427_[int(1)] = d_2590_v79_
            nw427_[int(2)] = d_2590_v79_
            nw427_[int(3)] = d_2590_v79_
            nw427_[int(4)] = d_2590_v79_
            nw427_[int(5)] = d_2590_v79_
            nw427_[int(6)] = d_2590_v79_
            nw427_[int(7)] = (d_2591_v80_)[default__.safeIndex(d_2501_v13_, len(d_2591_v80_))]
            nw427_[int(8)] = d_2590_v79_
            nw427_[int(9)] = d_2590_v79_
            nw427_[int(10)] = d_2590_v79_
            nw427_[int(11)] = d_2590_v79_
            nw427_[int(12)] = d_2590_v79_
            nw427_[int(13)] = d_2590_v79_
            nw427_[int(14)] = d_2590_v79_
            nw427_[int(15)] = d_2590_v79_
            nw427_[int(16)] = d_2590_v79_
            nw427_[int(17)] = d_2590_v79_
            nw427_[int(18)] = d_2590_v79_
            nw427_[int(19)] = d_2590_v79_
            nw427_[int(20)] = d_2590_v79_
            nw427_[int(21)] = d_2590_v79_
            nw427_[int(22)] = d_2590_v79_
            nw427_[int(23)] = d_2590_v79_
            nw427_[int(24)] = d_2590_v79_
            nw427_[int(25)] = d_2590_v79_
            d_2592_v81_ = nw427_
            index488_ = default__.safeIndex(478, (d_2592_v81_).length(0))
            (d_2592_v81_)[index488_] = d_2590_v79_
            index489_ = default__.safeIndex(478, (d_2592_v81_).length(0))
            nw428_ = C0()
            nw428_.ctor__()
            (d_2592_v81_)[index489_] = nw428_
            d_2593_v82_: _dafny.MultiSet
            d_2593_v82_ = _dafny.MultiSet([(d_2505_v17_).f38])
            d_2594_v83_: _dafny.Set
            d_2594_v83_ = _dafny.Set({(d_2505_v17_).f38})
            d_2595_v84_: _dafny.Map
            d_2595_v84_ = _dafny.Map({((d_2505_v17_).f37) + (d_2501_v13_): len(d_2502_v14_)})
            d_2596_v85_: _dafny.Seq
            d_2596_v85_ = _dafny.SeqWithoutIsStrInference([(d_2505_v17_).f37, d_2501_v13_, (d_2505_v17_).f37])
            rhs464_ = _dafny.SeqWithoutIsStrInference([((d_2593_v82_).cardinality if (d_2505_v17_).f38 else len(d_2594_v83_))])
            rhs465_ = ((d_2595_v84_)[default__.safeDivisionInt(len(d_2596_v85_), len(d_2502_v14_))] if (default__.safeDivisionInt(len(d_2596_v85_), len(d_2502_v14_))) in (d_2595_v84_) else (d_2505_v17_).f37)
            rhs466_ = (d_2505_v17_).f37
            lhs360_ = globalState
            lhs360_.f6 = rhs464_
            d_2501_v13_ = rhs465_
            d_2501_v13_ = rhs466_
            d_2597_v86_: _dafny.Array
            nw429_ = _dafny.Array(False, 21)
            d_2597_v86_ = nw429_
            d_2597_v86_ = d_2597_v86_
            (globalState).f10 = (d_2502_v14_) + (d_2502_v14_)
            d_2598_v87_: C5
            nw430_ = C5()
            nw430_.ctor__((d_2505_v17_).f38)
            d_2598_v87_ = nw430_
        elif True:
            d_2599_v88_: _dafny.Seq
            d_2599_v88_ = _dafny.SeqWithoutIsStrInference([d_2501_v13_])
            d_2600_v89_: _dafny.Map
            d_2600_v89_ = _dafny.Map({d_2599_v88_: ((d_2505_v17_).f38 if (d_2505_v17_).f38 else not(d_2503_v15_))})
            d_2600_v89_ = d_2600_v89_
            d_2601_v90_: C4
            nw431_ = C4()
            nw431_.ctor__(d_2504_v16_, (d_2505_v17_).f37, d_2504_v16_)
            d_2601_v90_ = nw431_
            d_2602_v91_: _dafny.Seq
            d_2602_v91_ = _dafny.SeqWithoutIsStrInference([d_2601_v90_, d_2601_v90_, d_2601_v90_])
            d_2601_v90_ = (d_2602_v91_)[default__.safeIndex(d_2501_v13_, len(d_2602_v91_))]
            r0 = (d_2505_v17_).f38
            d_2603_v92_: T1
            nw432_ = C6()
            nw432_.ctor__((len(d_2502_v14_)) - ((d_2505_v17_).f37), d_2504_v16_)
            d_2603_v92_ = nw432_
            d_2603_v92_ = d_2603_v92_
            d_2501_v13_ = d_2501_v13_
        d_2604_v93_: D4
        d_2604_v93_ = D4_DC13(d_2501_v13_, d_2503_v15_, default__.fm0((d_2505_v17_).f37, globalState), (d_2505_v17_).f38)
        pat_let_tv57_ = d_2505_v17_
        pat_let_tv58_ = d_2503_v15_
        pat_let_tv59_ = d_2505_v17_
        pat_let_tv60_ = d_2505_v17_
        pat_let_tv61_ = d_2503_v15_
        pat_let_tv62_ = d_2505_v17_
        def lambda165_(source33_):
            if source33_.is_DC12:
                return (_dafny.MultiSet([(pat_let_tv57_).f38])).ispropersubset(_dafny.MultiSet([pat_let_tv58_, not((pat_let_tv59_).f38), (pat_let_tv60_).f38]))
            elif source33_.is_DC13:
                d_2605___mcc_h5_ = source33_.cf19
                d_2606___mcc_h6_ = source33_.cf20
                d_2607___mcc_h7_ = source33_.cf21
                d_2608___mcc_h8_ = source33_.cf22
                d_2609_cf22_ = d_2608___mcc_h8_
                d_2610_cf21_ = d_2607___mcc_h7_
                d_2611_cf20_ = d_2606___mcc_h6_
                d_2612_cf19_ = d_2605___mcc_h5_
                return d_2611_cf20_
            elif source33_.is_DC11:
                d_2613___mcc_h9_ = source33_.cf18
                d_2614_cf18_ = d_2613___mcc_h9_
                return pat_let_tv61_
            elif True:
                d_2615___mcc_h10_ = source33_.cf23
                d_2616_cf23_ = d_2615___mcc_h10_
                return (pat_let_tv62_).f38

        r0 = lambda165_(d_2604_v93_)
        return r0


class C12(T1, T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f26: _dafny.Array = _dafny.Array(None, 0)
        self.f28: bool = False
        self._f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f27, f28, f25, f26):
        (self)._f27 = f27
        (self).f28 = f28
        (self)._f25 = f25
        (self)._f26 = f26

    def fm7(self, globalState):
        return _dafny.CodePoint('q')

    def fm8(self, p0, p1, globalState):
        return (D0_DC3((0) - ((self).f27), self.f28)).cf4

    def m4(self, p0, globalState):
        d_2617_v0_: _dafny.Seq
        d_2617_v0_ = _dafny.SeqWithoutIsStrInference([(self).fm8(default__.fm0(len(_dafny.Set({(self).f27})), globalState), self.f28, globalState)])
        d_2618_v1_: _dafny.Seq
        d_2618_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntbeock"))
        d_2619_v2_: _dafny.Map
        d_2619_v2_ = _dafny.Map({(self).f25: d_2618_v1_})
        d_2620_v3_: _dafny.Array
        nw433_ = _dafny.Array(int(0), 23)
        d_2620_v3_ = nw433_
        d_2621_v4_: _dafny.Map
        d_2621_v4_ = _dafny.Map({(self).f27: d_2620_v3_})
        d_2622_v5_: _dafny.Map
        d_2622_v5_ = _dafny.Map({(self).f27: len(d_2621_v4_)})
        d_2623_i0_: int
        d_2623_i0_ = 0
        with _dafny.label("21"):
            while (default__.safeDivisionInt((self).f27, (self).f27)) >= (((d_2617_v0_)[default__.safeIndex((self).f25, len(d_2617_v0_))]) * (len(((d_2619_v2_)[len(d_2622_v5_)] if (len(d_2622_v5_)) in (d_2619_v2_) else d_2618_v1_)))):
                with _dafny.c_label("21"):
                    if (d_2623_i0_) >= (100):
                        raise _dafny.Break("21")
                    d_2623_i0_ = (d_2623_i0_) + (1)
                    d_2624_v6_: _dafny.Seq
                    d_2624_v6_ = _dafny.SeqWithoutIsStrInference([self.f28, self.f28])
                    d_2625_v7_: _dafny.Array
                    nw434_ = _dafny.Array(None, 24)
                    nw434_[int(0)] = self.f28
                    nw434_[int(1)] = self.f28
                    nw434_[int(2)] = self.f28
                    nw434_[int(3)] = self.f28
                    nw434_[int(4)] = False
                    nw434_[int(5)] = default__.fm0((self).f25, globalState)
                    nw434_[int(6)] = self.f28
                    nw434_[int(7)] = not(self.f28)
                    nw434_[int(8)] = self.f28
                    nw434_[int(9)] = (d_2624_v6_)[default__.safeIndex((self).f27, len(d_2624_v6_))]
                    nw434_[int(10)] = self.f28
                    nw434_[int(11)] = self.f28
                    nw434_[int(12)] = self.f28
                    nw434_[int(13)] = self.f28
                    nw434_[int(14)] = True
                    nw434_[int(15)] = False
                    nw434_[int(16)] = self.f28
                    nw434_[int(17)] = self.f28
                    nw434_[int(18)] = self.f28
                    nw434_[int(19)] = self.f28
                    nw434_[int(20)] = self.f28
                    nw434_[int(21)] = self.f28
                    nw434_[int(22)] = self.f28
                    nw434_[int(23)] = self.f28
                    d_2625_v7_ = nw434_
                    d_2626_v8_: _dafny.Map
                    d_2626_v8_ = _dafny.Map({self.f28: d_2625_v7_})
                    d_2627_v9_: _dafny.Seq
                    d_2627_v9_ = _dafny.SeqWithoutIsStrInference([d_2626_v8_])
                    d_2626_v8_ = (((d_2626_v8_).set(self.f28, d_2625_v7_)) | (d_2626_v8_)) | (((d_2627_v9_)[default__.safeIndex((self).f27, len(d_2627_v9_))]) | (d_2626_v8_))
                    d_2628_v10_: int
                    d_2628_v10_ = 40
                    d_2628_v10_ = (default__.safeDivisionInt((self).f27, d_2628_v10_)) + (d_2628_v10_)
                    d_2625_v7_ = d_2625_v7_
                    d_2629_v11_: _dafny.Set
                    d_2629_v11_ = _dafny.Set({d_2617_v0_, d_2617_v0_})
                    d_2629_v11_ = d_2629_v11_
                    pass
            pass
        d_2630_v12_: _dafny.MultiSet
        d_2630_v12_ = _dafny.MultiSet([self.f28, self.f28])
        d_2631_v13_: _dafny.Seq
        d_2631_v13_ = _dafny.SeqWithoutIsStrInference([self.f28, self.f28])
        d_2632_v14_: C8
        nw435_ = C8()
        nw435_.ctor__((d_2630_v12_).ispropersubset(_dafny.MultiSet(d_2631_v13_)), not(self.f28), (0) - (default__.safeDivisionInt((self).f25, (self).f25)), (self).f26)
        d_2632_v14_ = nw435_
        d_2633_v15_: _dafny.Array
        def lambda166_(d_2634_i2_):
            return not(self.f28)

        init91_ = lambda166_
        nw436_ = _dafny.Array(None, 24)
        for i0_91_ in range(nw436_.length(0)):
            nw436_[i0_91_] = init91_(i0_91_)
        d_2633_v15_ = nw436_
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, (d_2633_v15_).length(0)):
            d_2635_i1_: int = guard_loop_12_
            if (True) and (((0) <= (d_2635_i1_)) and ((d_2635_i1_) < ((d_2633_v15_).length(0)))):
                (d_2633_v15_)[(d_2635_i1_)] = (D5_DC17((self).f27, (d_2632_v14_).f32)).cf27
        d_2636_v16_: _dafny.MultiSet
        d_2636_v16_ = _dafny.MultiSet([(self).f27])
        d_2637_v17_: _dafny.Set
        d_2637_v17_ = _dafny.Set({(d_2631_v13_) + (d_2631_v13_), (d_2631_v13_) + (_dafny.SeqWithoutIsStrInference([(d_2632_v14_).f33, (d_2632_v14_).f33, (d_2632_v14_).f33])), d_2631_v13_})
        d_2638_v18_: _dafny.Array
        nw437_ = _dafny.Array(_dafny.MultiSet({}), 7)
        d_2638_v18_ = nw437_
        d_2639_v19_: D3
        d_2639_v19_ = D3_DC10(not((d_2632_v14_).f33), d_2633_v15_, (d_2632_v14_).f32, (d_2632_v14_).fm14(self.f28, globalState), (self).f27)
        d_2640_v20_: _dafny.MultiSet
        d_2640_v20_ = _dafny.MultiSet([d_2639_v19_, d_2639_v19_])
        index490_ = default__.safeIndex(922, (d_2638_v18_).length(0))
        (d_2638_v18_)[index490_] = (d_2640_v20_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2639_v19_, D3_DC10(self.f28, d_2633_v15_, False, (self).f27, (self).f27), d_2639_v19_, d_2639_v19_, d_2639_v19_])))
        index491_ = default__.safeIndex(74, (d_2620_v3_).length(0))
        (d_2620_v3_)[index491_] = (self).f27
        d_2641_v21_: str
        d_2641_v21_ = _dafny.CodePoint('y')
        index492_ = default__.safeIndex(922, (d_2638_v18_).length(0))
        index493_ = default__.safeIndex(74, (d_2620_v3_).length(0))
        rhs467_ = (default__.fm15(globalState) if default__.fm0(883, globalState) else ((d_2636_v16_).set((self).f27, default__.abs((self).f25)) if self.f28 else _dafny.MultiSet([(self).f25, (self).f27])))
        rhs468_ = d_2641_v21_
        rhs469_ = (d_2637_v17_) | (d_2637_v17_)
        rhs470_ = (_dafny.MultiSet([d_2639_v19_])) - (d_2640_v20_)
        rhs471_ = (self).f25
        lhs361_ = globalState
        lhs362_ = d_2638_v18_
        lhs363_ = default__.safeIndex(922, (d_2638_v18_).length(0))
        lhs364_ = d_2620_v3_
        lhs365_ = default__.safeIndex(74, (d_2620_v3_).length(0))
        d_2636_v16_ = rhs467_
        lhs361_.f23 = rhs468_
        d_2637_v17_ = rhs469_
        lhs362_[lhs363_] = rhs470_
        lhs364_[lhs365_] = rhs471_
        guard_loop_13_: int
        for guard_loop_13_ in _dafny.IntegerRange(0, (d_2620_v3_).length(0)):
            d_2642_i3_: int = guard_loop_13_
            if (True) and (((0) <= (d_2642_i3_)) and ((d_2642_i3_) < ((d_2620_v3_).length(0)))):
                (d_2620_v3_)[(d_2642_i3_)] = (d_2642_i3_) - (default__.safeDivisionInt(582, (0) - ((self).f27)))
        d_2643_v22_: _dafny.Array
        nw438_ = _dafny.Array(_dafny.MultiSet({}), 24)
        d_2643_v22_ = nw438_
        d_2644_v23_: D4
        d_2644_v23_ = D4_DC13((d_2620_v3_)[default__.safeIndex(74, (d_2620_v3_).length(0))], (d_2632_v14_).f33, (d_2632_v14_).f33, self.f28)
        d_2645_v24_: _dafny.MultiSet
        d_2645_v24_ = _dafny.MultiSet([D4_DC13((self).f25, self.f28, (d_2632_v14_).f32, (d_2632_v14_).f32), d_2644_v23_])
        index494_ = default__.safeIndex(908, (d_2643_v22_).length(0))
        (d_2643_v22_)[index494_] = (d_2645_v24_) | (d_2645_v24_)
        d_2646_v25_: _dafny.Seq
        d_2646_v25_ = d_2617_v0_
        d_2647_v26_: D14
        d_2647_v26_ = D14_DC38((d_2632_v14_).fm14(self.f28, globalState))
        pat_let_tv63_ = d_2620_v3_
        pat_let_tv64_ = d_2620_v3_
        index495_ = default__.safeIndex(74, (d_2620_v3_).length(0))
        index496_ = default__.safeIndex(908, (d_2643_v22_).length(0))
        def lambda167_(source34_):
            if source34_.is_DC38:
                d_2648___mcc_h0_ = source34_.cf58
                d_2649_cf58_ = d_2648___mcc_h0_
                if not(False):
                    return (pat_let_tv64_)[default__.safeIndex(74, (pat_let_tv63_).length(0))]
                elif True:
                    return (self).f25
            elif True:
                d_2650___mcc_h1_ = source34_.cf57
                d_2651_cf57_ = d_2650___mcc_h1_
                return (self).f27

        rhs472_ = lambda167_(d_2647_v26_)
        rhs473_ = (d_2645_v24_ if True else (d_2645_v24_) | (d_2645_v24_))
        rhs474_ = d_2620_v3_
        rhs475_ = d_2646_v25_
        lhs366_ = d_2620_v3_
        lhs367_ = default__.safeIndex(74, (d_2620_v3_).length(0))
        lhs368_ = d_2643_v22_
        lhs369_ = default__.safeIndex(908, (d_2643_v22_).length(0))
        lhs366_[lhs367_] = rhs472_
        lhs368_[lhs369_] = rhs473_
        d_2620_v3_ = rhs474_
        d_2646_v25_ = rhs475_

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_2652_v0_: str
        d_2652_v0_ = _dafny.CodePoint('x')
        d_2653_v1_: _dafny.Seq
        d_2653_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evoys"))
        r0 = ((0) - ((self).f27) if default__.fm0(len((D1_DC6(d_2652_v0_, p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvrcenaw")))).cf9), globalState) else len(d_2653_v1_))
        d_2652_v0_ = _dafny.CodePoint('b')
        d_2654_i0_: int
        d_2654_i0_ = 0
        with _dafny.label("22"):
            while self.f28:
                with _dafny.c_label("22"):
                    if (d_2654_i0_) >= (100):
                        raise _dafny.Break("22")
                    d_2654_i0_ = (d_2654_i0_) + (1)
                    d_2655_v2_: _dafny.Seq
                    d_2655_v2_ = _dafny.SeqWithoutIsStrInference([p2, p1])
                    d_2656_v3_: _dafny.MultiSet
                    d_2656_v3_ = _dafny.MultiSet([d_2655_v2_])
                    d_2657_v4_: C8
                    nw439_ = C8()
                    nw439_.ctor__((d_2656_v3_).issubset(d_2656_v3_), p1, default__.fm18(globalState), (self).f26)
                    d_2657_v4_ = nw439_
                    d_2657_v4_ = d_2657_v4_
                    (globalState).f10 = (d_2653_v1_) + (d_2653_v1_)
                    index497_ = default__.safeIndex(212, (p3).length(0))
                    (p3)[index497_] = default__.safeModuloInt(886, (self).f25)
                    index498_ = default__.safeIndex(212, (p3).length(0))
                    (p3)[index498_] = (self).fm8(p2, self.f28, globalState)
                    (self).f28 = (d_2657_v4_).fm13(globalState)
                    pass
            pass
        d_2658_v5_: _dafny.MultiSet
        d_2658_v5_ = _dafny.MultiSet([(self).f27, (self).f27])
        d_2659_v6_: _dafny.Map
        d_2659_v6_ = _dafny.Map({_dafny.Set({p2, p2}): p1})
        d_2660_v7_: C10
        nw440_ = C10()
        nw440_.ctor__(d_2652_v0_, ((d_2658_v5_)[default__.fm44(p2, p2, p2, len(d_2659_v6_), globalState)] if (default__.fm44(p2, p2, p2, len(d_2659_v6_), globalState)) in (d_2658_v5_) else len(_dafny.Map({p1: _dafny.CodePoint('a')}))), (self).f26)
        d_2660_v7_ = nw440_
        d_2661_v8_: _dafny.Set
        d_2661_v8_ = _dafny.Set({d_2660_v7_})
        d_2662_v9_: _dafny.Set
        d_2662_v9_ = _dafny.Set({d_2661_v8_, d_2661_v8_, d_2661_v8_})
        r1 = ((self).f27 if self.f28 else len((d_2662_v9_).intersection(d_2662_v9_)))
        d_2663_i1_: int
        d_2663_i1_ = 0
        with _dafny.label("23"):
            while p2:
                with _dafny.c_label("23"):
                    if (d_2663_i1_) >= (100):
                        raise _dafny.Break("23")
                    d_2663_i1_ = (d_2663_i1_) + (1)
                    r1 = (0) - ((self).f25)
                    d_2664_v10_: _dafny.Set
                    d_2664_v10_ = _dafny.Set({p2})
                    d_2665_v11_: _dafny.Map
                    d_2665_v11_ = _dafny.Map({d_2664_v10_: d_2658_v5_})
                    d_2666_v12_: _dafny.Seq
                    d_2666_v12_ = _dafny.SeqWithoutIsStrInference([604, (self).f27])
                    d_2665_v11_ = (d_2665_v11_).set((d_2664_v10_) | (d_2664_v10_), (d_2658_v5_).set(-477, default__.abs(len(d_2666_v12_))))
                    d_2667_v13_: D5
                    d_2667_v13_ = D5_DC17((self).f27, p2)
                    d_2668_v14_: _dafny.Map
                    d_2668_v14_ = _dafny.Map({(d_2667_v13_).cf27: (self).f27})
                    d_2669_v15_: D9
                    d_2669_v15_ = D9_DC26(d_2668_v14_)
                    source35_ = d_2669_v15_
                    if source35_.is_DC27:
                        d_2670___mcc_h0_ = source35_.cf39
                        d_2671___mcc_h1_ = source35_.cf40
                        d_2672_cf40_ = d_2671___mcc_h1_
                        d_2673_cf39_ = d_2670___mcc_h0_
                        d_2653_v1_ = ((d_2653_v1_).set(default__.safeIndex(d_2673_cf39_, len(d_2653_v1_)), d_2652_v0_)) + ((d_2653_v1_).set(default__.safeIndex((self).f25, len(d_2653_v1_)), _dafny.CodePoint('u')))
                        d_2674_v16_: _dafny.Seq
                        d_2674_v16_ = _dafny.SeqWithoutIsStrInference([False])
                        d_2675_v17_: C9
                        nw441_ = C9()
                        nw441_.ctor__((p3 if True else p3), ((d_2666_v12_)[default__.safeIndex(len(d_2674_v16_), len(d_2666_v12_))]) + (len(d_2666_v12_)))
                        d_2675_v17_ = nw441_
                        d_2676_v18_: _dafny.Seq
                        d_2676_v18_ = _dafny.SeqWithoutIsStrInference([d_2667_v13_])
                        d_2677_v19_: C4
                        nw442_ = C4()
                        nw442_.ctor__((self).f26, (len(d_2666_v12_)) * (len(d_2676_v18_)), (self).f26)
                        d_2677_v19_ = nw442_
                        (d_2675_v17_).f30 = d_2675_v17_.f30
                    elif True:
                        d_2678___mcc_h2_ = source35_.cf38
                        d_2679_cf38_ = d_2678___mcc_h2_
                        (globalState).f6 = (d_2666_v12_ if not(p2) else d_2666_v12_)
                        d_2680_v20_: _dafny.Map
                        d_2680_v20_ = _dafny.Map({(self).f25: _dafny.SeqWithoutIsStrInference([(d_2660_v7_).f29 for d_2681_i2_ in range(default__.abs(842))])})
                        (globalState).f6 = _dafny.SeqWithoutIsStrInference([len(d_2666_v12_), ((self).f27) + (len(default__.fm27(p2, d_2680_v20_, globalState))), (self).f25])
                        d_2682_v21_: _dafny.Array
                        nw443_ = _dafny.Array(False, 10)
                        d_2682_v21_ = nw443_
                        d_2683_v22_: _dafny.Seq
                        d_2683_v22_ = _dafny.SeqWithoutIsStrInference([d_2653_v1_, d_2653_v1_])
                        index499_ = default__.safeIndex(291, (d_2682_v21_).length(0))
                        (d_2682_v21_)[index499_] = ((d_2683_v22_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuaexm"))), len(d_2683_v22_))]) != (d_2653_v1_)
                        d_2684_v23_: _dafny.Map
                        d_2684_v23_ = _dafny.Map({d_2664_v10_: (self).f25})
                        index500_ = default__.safeIndex(291, (d_2682_v21_).length(0))
                        (d_2682_v21_)[index500_] = (d_2684_v23_) == (d_2684_v23_)
                        r1 = len(d_2653_v1_)
                    d_2685_v24_: _dafny.Map
                    d_2685_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2686_i3_ in range(default__.abs(397))]): d_2668_v14_})
                    def iife213_():
                        coll89_ = _dafny.Map()
                        compr_89_: _dafny.Seq
                        for compr_89_ in (_dafny.Map({d_2653_v1_: d_2653_v1_})).keys.Elements:
                            d_2687_v25_: _dafny.Seq = compr_89_
                            if (d_2687_v25_) in (_dafny.Map({d_2653_v1_: d_2653_v1_})):
                                coll89_[d_2687_v25_] = d_2668_v14_
                        return _dafny.Map(coll89_)
                    d_2685_v24_ = (iife213_()
                    ).set(d_2653_v1_, d_2668_v14_)
                    pass
            pass
        d_2688_v26_: _dafny.Seq
        d_2688_v26_ = _dafny.SeqWithoutIsStrInference([p2, default__.fm0((self).f25, globalState)])
        d_2689_v27_: _dafny.Map
        d_2689_v27_ = _dafny.Map({(d_2688_v26_)[default__.safeIndex((self).f27, len(d_2688_v26_))]: p3})
        d_2689_v27_ = (d_2689_v27_).set(p2, (p3 if not(p1) else p3))
        r0 = (0) - (default__.safeDivisionInt((self).f27, ((self).f27) - ((self).f27)))
        r1 = default__.safeDivisionInt((self).fm8(p1, not(self.f28), globalState), 405)
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        d_2690_v0_: _dafny.Seq
        d_2690_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2691_v1_: D0
        d_2691_v1_ = D0_DC0(not((d_2690_v0_)[default__.safeIndex((self).f27, len(d_2690_v0_))]))
        source36_ = d_2691_v1_
        if source36_.is_DC1:
            d_2692___mcc_h0_ = source36_.cf1
            d_2693___mcc_h1_ = source36_.cf2
            d_2694_cf2_ = d_2693___mcc_h1_
            d_2695_cf1_ = d_2692___mcc_h0_
            d_2696_v2_: _dafny.Set
            d_2696_v2_ = _dafny.Set({d_2694_cf2_, d_2694_cf2_, self.f28, p1, self.f28})
            if ((d_2696_v2_).isdisjoint(d_2696_v2_) if self.f28 else True):
                d_2697_v3_: _dafny.Array
                def lambda168_(d_2698_i0_):
                    return (d_2698_i0_) + ((self).f27)

                init92_ = lambda168_
                nw444_ = _dafny.Array(None, 4)
                for i0_92_ in range(nw444_.length(0)):
                    nw444_[i0_92_] = init92_(i0_92_)
                d_2697_v3_ = nw444_
                d_2699_v4_: _dafny.Map
                d_2699_v4_ = _dafny.Map({(self).f27: self.f28})
                d_2700_v5_: _dafny.Map
                d_2700_v5_ = _dafny.Map({self.f28: len(d_2699_v4_)})
                d_2701_v6_: _dafny.Map
                d_2701_v6_ = _dafny.Map({self.f28: d_2700_v5_})
                d_2702_v7_: C9
                nw445_ = C9()
                nw445_.ctor__(d_2697_v3_, len(d_2701_v6_))
                d_2702_v7_ = nw445_
                d_2703_v8_: _dafny.Seq
                d_2703_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                d_2704_v9_: _dafny.Seq
                d_2704_v9_ = _dafny.SeqWithoutIsStrInference([679, (0) - (len(d_2703_v8_)), (self).f25, (self).f25, (self).f25])
                d_2694_cf2_ = (d_2704_v9_) < ((d_2704_v9_).set(default__.safeIndex((d_2702_v7_).f31, len(d_2704_v9_)), (self).f27))
                d_2705_v10_: C0
                nw446_ = C0()
                nw446_.ctor__()
                d_2705_v10_ = nw446_
                def iife214_():
                    coll90_ = _dafny.Map()
                    compr_90_: int
                    for compr_90_ in (_dafny.Map({(self).f27: d_2695_cf1_})).keys.Elements:
                        d_2706_v11_: int = compr_90_
                        if (d_2706_v11_) in (_dafny.Map({(self).f27: d_2695_cf1_})):
                            coll90_[(d_2706_v11_) + ((d_2702_v7_).f31)] = _dafny.CodePoint('s')
                    return _dafny.Map(coll90_)
                (globalState).f12 = (default__.fm44(d_2695_cf1_, d_2695_cf1_, self.f28, len(iife214_()
                ), globalState)) not in (_dafny.SeqWithoutIsStrInference([(d_2702_v7_).f31, (d_2702_v7_).f31, (self).f27]))
                d_2707_v12_: _dafny.Map
                d_2707_v12_ = _dafny.Map({d_2695_cf1_: False})
                d_2707_v12_ = (d_2707_v12_).set(((default__.fm30(d_2699_v4_, globalState)).cf8) not in (d_2690_v0_), d_2694_cf2_)
            elif True:
                d_2708_v13_: _dafny.Map
                d_2708_v13_ = _dafny.Map({(self).f25: True})
                d_2708_v13_ = (d_2708_v13_).set((self).f27, d_2694_cf2_)
                d_2709_v14_: str
                d_2709_v14_ = _dafny.CodePoint('u')
                d_2710_v15_: _dafny.MultiSet
                d_2710_v15_ = _dafny.MultiSet([d_2709_v14_, d_2709_v14_])
                d_2711_v16_: _dafny.Map
                d_2711_v16_ = _dafny.Map({True: (d_2710_v15_).intersection(d_2710_v15_)})
                d_2712_v17_: _dafny.MultiSet
                d_2712_v17_ = _dafny.MultiSet([(self).f25])
                d_2713_v18_: _dafny.MultiSet
                d_2713_v18_ = _dafny.MultiSet([p1])
                d_2711_v16_ = default__.fm48(d_2695_cf1_, default__.safeDivisionInt((d_2712_v17_).cardinality, -146), (self).f25, ((d_2713_v18_).intersection(d_2713_v18_)).cardinality, globalState)
                r0 = ((self).f25 if p1 else (self).f25)
                (globalState).f23 = d_2709_v14_
                d_2694_cf2_ = (len(d_2696_v2_)) >= (442)
            d_2714_v19_: C11
            nw447_ = C11()
            nw447_.ctor__()
            d_2714_v19_ = nw447_
            d_2715_v20_: _dafny.Seq
            d_2715_v20_ = _dafny.SeqWithoutIsStrInference([d_2714_v19_, d_2714_v19_])
            d_2716_v22_: _dafny.Map
            def iife215_():
                coll91_ = _dafny.Map()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(-774, 769):
                    d_2717_v21_: int = compr_91_
                    if ((-774) <= (d_2717_v21_)) and ((d_2717_v21_) < (769)):
                        coll91_[(d_2717_v21_) - ((self).f25)] = (self).f27
                return _dafny.Map(coll91_)
            d_2716_v22_ = _dafny.Map({d_2695_cf1_: len(iife215_()
            )})
            d_2718_v23_: _dafny.Map
            d_2718_v23_ = _dafny.Map({p1: (d_2715_v20_)[default__.safeIndex(len(d_2716_v22_), len(d_2715_v20_))]})
            d_2719_v24_: _dafny.Map
            d_2719_v24_ = _dafny.Map({not(d_2694_cf2_): ((d_2718_v23_)[d_2695_cf1_] if (d_2695_cf1_) in (d_2718_v23_) else d_2714_v19_)})
            d_2719_v24_ = (d_2719_v24_).set(not(p1), d_2714_v19_)
            d_2720_v25_: _dafny.Seq
            d_2720_v25_ = _dafny.SeqWithoutIsStrInference([(self).f25, (0) - ((self).f27), 849])
            d_2721_v26_: D5
            d_2721_v26_ = D5_DC17((0) - (((0) - ((self).f25)) - ((self).f27)), ((d_2720_v25_)[default__.safeIndex((self).f25, len(d_2720_v25_))]) > ((self).f27))
            d_2722_v27_: _dafny.Seq
            d_2722_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfdfdgsi"))
            d_2723_v28_: _dafny.Seq
            d_2723_v28_ = _dafny.SeqWithoutIsStrInference([d_2722_v27_])
            d_2724_v29_: _dafny.Array
            nw448_ = _dafny.Array(None, 4)
            nw448_[int(0)] = ((d_2723_v28_)[default__.safeIndex(739, len(d_2723_v28_))]) != (d_2722_v27_)
            nw448_[int(1)] = True
            nw448_[int(2)] = not(not(p1))
            nw448_[int(3)] = ((self).f27) < ((self).f27)
            d_2724_v29_ = nw448_
            index501_ = default__.safeIndex(762, (d_2724_v29_).length(0))
            (d_2724_v29_)[index501_] = self.f28
            d_2725_v30_: _dafny.Map
            d_2725_v30_ = _dafny.Map({len(d_2720_v25_): d_2722_v27_})
            pat_let_tv65_ = d_2695_cf1_
            index502_ = default__.safeIndex(762, (d_2724_v29_).length(0))
            def iife216_(_pat_let62_0):
                def iife217_(d_2726_dt__update__tmp_h0_):
                    def iife218_(_pat_let63_0):
                        def iife219_(d_2727_dt__update_hcf27_h0_):
                            return D5_DC17((d_2726_dt__update__tmp_h0_).cf26, d_2727_dt__update_hcf27_h0_)
                        return iife219_(_pat_let63_0)
                    return iife218_(pat_let_tv65_)
                return iife217_(_pat_let62_0)
            rhs476_ = (self).f25
            rhs477_ = (d_2722_v27_) <= (((d_2725_v30_)[(self).f27] if ((self).f27) in (d_2725_v30_) else d_2722_v27_))
            rhs478_ = iife216_(d_2721_v26_)
            rhs479_ = ((self).f27) < ((self).f25)
            lhs370_ = d_2724_v29_
            lhs371_ = default__.safeIndex(762, (d_2724_v29_).length(0))
            r0 = rhs476_
            d_2695_cf1_ = rhs477_
            d_2721_v26_ = rhs478_
            lhs370_[lhs371_] = rhs479_
            index503_ = default__.safeIndex(762, (d_2724_v29_).length(0))
            (d_2724_v29_)[index503_] = d_2694_cf2_
        elif source36_.is_DC2:
            d_2728___mcc_h2_ = source36_.cf3
            d_2729_cf3_ = d_2728___mcc_h2_
            d_2730_v31_: _dafny.Set
            d_2730_v31_ = _dafny.Set({(self).f27})
            d_2729_cf3_ = ((d_2730_v31_).intersection(_dafny.Set({(self).f25}))).issubset((d_2730_v31_) - (d_2730_v31_))
            d_2731_v32_: _dafny.Map
            d_2731_v32_ = _dafny.Map({(self).f25: ((0) - ((self).f27)) * (716)})
            d_2731_v32_ = (d_2731_v32_).set((self).f27, (self).f25)
            r0 = (self).f25
            d_2732_v33_: _dafny.Set
            d_2732_v33_ = _dafny.Set({self.f28})
            d_2733_v34_: _dafny.MultiSet
            d_2733_v34_ = _dafny.MultiSet([d_2732_v33_, _dafny.Set({d_2729_cf3_})])
            d_2734_v35_: _dafny.Seq
            d_2734_v35_ = _dafny.SeqWithoutIsStrInference([d_2732_v33_])
            d_2735_v36_: _dafny.MultiSet
            d_2735_v36_ = _dafny.MultiSet([d_2733_v34_, _dafny.MultiSet(d_2734_v35_)])
            d_2736_v37_: _dafny.Seq
            d_2736_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjamaf"))
            d_2737_v38_: str
            d_2737_v38_ = _dafny.CodePoint('n')
            d_2738_v39_: _dafny.MultiSet
            d_2738_v39_ = _dafny.MultiSet([d_2737_v38_, d_2737_v38_, d_2737_v38_, _dafny.CodePoint('n'), d_2737_v38_])
            d_2739_v40_: _dafny.Map
            d_2739_v40_ = _dafny.Map({(self).f27: self.f28})
            d_2740_v42_: _dafny.Map
            d_2740_v42_ = _dafny.Map({d_2691_v1_: p1})
            d_2741_v43_: _dafny.Array
            nw449_ = _dafny.Array(None, 21)
            nw449_[int(0)] = (self).f25
            nw449_[int(1)] = (0) - ((self).f27)
            nw449_[int(2)] = ((d_2735_v36_)[d_2733_v34_] if (d_2733_v34_) in (d_2735_v36_) else len(_dafny.Map({False: d_2729_cf3_})))
            nw449_[int(3)] = (self).f25
            nw449_[int(4)] = default__.fm31(d_2736_v37_, globalState)
            nw449_[int(5)] = (self).f27
            nw449_[int(6)] = (self).f25
            nw449_[int(7)] = 89
            nw449_[int(8)] = (self).f27
            nw449_[int(9)] = (d_2738_v39_).cardinality
            nw449_[int(10)] = (self).f25
            nw449_[int(11)] = (self).f25
            def iife220_():
                coll92_ = _dafny.Set()
                compr_92_: int
                for compr_92_ in (d_2739_v40_).keys.Elements:
                    d_2742_v41_: int = compr_92_
                    if (d_2742_v41_) in (d_2739_v40_):
                        coll92_ = coll92_.union(_dafny.Set([default__.safeModuloInt(d_2742_v41_, (0) - (len(_dafny.SeqWithoutIsStrInference([True, not(True)]))))]))
                return _dafny.Set(coll92_)
            nw449_[int(12)] = len(iife220_()
            )
            nw449_[int(13)] = ((d_2731_v32_)[default__.fm21(d_2740_v42_, len(p0), d_2729_cf3_, d_2736_v37_, globalState)] if (default__.fm21(d_2740_v42_, len(p0), d_2729_cf3_, d_2736_v37_, globalState)) in (d_2731_v32_) else (self).f25)
            nw449_[int(14)] = len(d_2731_v32_)
            nw449_[int(15)] = (self).f27
            nw449_[int(16)] = (self).f27
            nw449_[int(17)] = len(d_2732_v33_)
            nw449_[int(18)] = len(d_2730_v31_)
            nw449_[int(19)] = (self).f27
            nw449_[int(20)] = (self).f27
            d_2741_v43_ = nw449_
            d_2743_v44_: _dafny.Seq
            d_2743_v44_ = _dafny.SeqWithoutIsStrInference([(self).f27, 551])
            d_2744_v45_: _dafny.MultiSet
            d_2744_v45_ = _dafny.MultiSet([(d_2743_v44_)[default__.safeIndex((self).f27, len(d_2743_v44_))], (self).f27, (self).f27])
            d_2745_v46_: D5
            d_2745_v46_ = D5_DC15(d_2744_v45_)
            d_2746_v47_: _dafny.Seq
            d_2746_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f25, (self).f25, len(d_2739_v40_)]), d_2744_v45_, (d_2745_v46_).cf24])
            d_2747_v48_: _dafny.Map
            d_2747_v48_ = _dafny.Map({d_2741_v43_: (d_2746_v47_)[default__.safeIndex((self).f27, len(d_2746_v47_))]})
            d_2747_v48_ = d_2747_v48_
        elif source36_.is_DC3:
            d_2748___mcc_h3_ = source36_.cf4
            d_2749___mcc_h4_ = source36_.cf5
            d_2750_cf5_ = d_2749___mcc_h4_
            d_2751_cf4_ = d_2748___mcc_h3_
            d_2752_v49_: _dafny.Array
            nw450_ = _dafny.Array(False, 23)
            d_2752_v49_ = nw450_
            d_2753_v50_: _dafny.Map
            d_2753_v50_ = _dafny.Map({default__.fm0(d_2751_cf4_, globalState): d_2752_v49_})
            d_2753_v50_ = (d_2753_v50_).set(d_2750_cf5_, d_2752_v49_)
            d_2754_v51_: _dafny.Map
            d_2754_v51_ = _dafny.Map({self.f28: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2755_i1_ in range(default__.abs(326))])})
            (globalState).f10 = ((d_2754_v51_)[p1] if (p1) in (d_2754_v51_) else default__.fm2(267, (self).f27, globalState))
            index504_ = default__.safeIndex(619, (d_2752_v49_).length(0))
            (d_2752_v49_)[index504_] = (d_2690_v0_)[default__.safeIndex(d_2751_cf4_, len(d_2690_v0_))]
            index505_ = default__.safeIndex(619, (d_2752_v49_).length(0))
            (d_2752_v49_)[index505_] = not(not(d_2750_cf5_))
            (globalState).f23 = default__.fm33(d_2751_cf4_, not(d_2750_cf5_), (self).fm8((d_2752_v49_)[default__.safeIndex(619, (d_2752_v49_).length(0))], (d_2752_v49_)[default__.safeIndex(619, (d_2752_v49_).length(0))], globalState), globalState)
        elif True:
            d_2756___mcc_h5_ = source36_.cf0
            d_2757_cf0_ = d_2756___mcc_h5_
            d_2758_v52_: _dafny.Set
            d_2758_v52_ = _dafny.Set({p1, not(self.f28), d_2757_cf0_})
            d_2759_v53_: _dafny.Seq
            d_2759_v53_ = _dafny.SeqWithoutIsStrInference([d_2758_v52_, d_2758_v52_, d_2758_v52_, d_2758_v52_, d_2758_v52_])
            d_2759_v53_ = d_2759_v53_
            d_2760_v54_: _dafny.Array
            nw451_ = _dafny.Array(False, 3)
            d_2760_v54_ = nw451_
            index506_ = default__.safeIndex(225, (d_2760_v54_).length(0))
            (d_2760_v54_)[index506_] = not(p1)
            index507_ = default__.safeIndex(225, (d_2760_v54_).length(0))
            (d_2760_v54_)[index507_] = d_2757_cf0_
            if ((self).f25) >= ((self).f27):
                index508_ = default__.safeIndex(225, (d_2760_v54_).length(0))
                (d_2760_v54_)[index508_] = not(((self).f25) >= ((302) * ((self).f27)))
                rhs480_ = (d_2760_v54_)[default__.safeIndex(225, (d_2760_v54_).length(0))]
                rhs481_ = not(p1)
                rhs482_ = (self).f27
                rhs483_ = ((self).f25) - ((self).f25)
                lhs372_ = globalState
                lhs373_ = globalState
                lhs372_.f12 = rhs480_
                lhs373_.f12 = rhs481_
                r0 = rhs482_
                r0 = rhs483_
                r0 = 454
                (self).f28 = (default__.safeModuloInt((self).f25, (self).f27)) > (506)
                d_2761_v55_: _dafny.Array
                def lambda169_(d_2762_v0_):
                    def lambda170_(d_2763_i2_):
                        return (d_2763_i2_) + (len(d_2762_v0_))

                    return lambda170_

                init93_ = lambda169_(d_2690_v0_)
                nw452_ = _dafny.Array(None, 2)
                for i0_93_ in range(nw452_.length(0)):
                    nw452_[i0_93_] = init93_(i0_93_)
                d_2761_v55_ = nw452_
                d_2764_v56_: _dafny.Map
                d_2764_v56_ = _dafny.Map({(self).f27: (self).f25})
                rhs484_ = self.f28
                rhs485_ = (d_2761_v55_ if (859) <= (len(d_2764_v56_)) else d_2761_v55_)
                rhs486_ = (0) - ((self).f25)
                rhs487_ = p1
                rhs488_ = (d_2690_v0_)[default__.safeIndex((self).f27, len(d_2690_v0_))]
                lhs374_ = self
                lhs375_ = self
                lhs376_ = self
                lhs374_.f28 = rhs484_
                d_2761_v55_ = rhs485_
                r0 = rhs486_
                lhs375_.f28 = rhs487_
                lhs376_.f28 = rhs488_
            elif True:
                r0 = (self).f25
                d_2765_v57_: _dafny.Map
                d_2765_v57_ = _dafny.Map({(d_2690_v0_)[default__.safeIndex((self).f27, len(d_2690_v0_))]: (self).f25})
                d_2765_v57_ = (d_2765_v57_) | ((d_2765_v57_) | (d_2765_v57_))
                d_2766_v58_: _dafny.Array
                nw453_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_2766_v58_ = nw453_
                d_2767_v59_: _dafny.Array
                nw454_ = _dafny.Array(None, 8)
                nw454_[int(0)] = d_2766_v58_
                nw454_[int(1)] = d_2766_v58_
                nw454_[int(2)] = d_2766_v58_
                nw454_[int(3)] = d_2766_v58_
                nw454_[int(4)] = d_2766_v58_
                nw454_[int(5)] = d_2766_v58_
                nw454_[int(6)] = d_2766_v58_
                nw454_[int(7)] = d_2766_v58_
                d_2767_v59_ = nw454_
                index509_ = default__.safeIndex(664, (d_2767_v59_).length(0))
                (d_2767_v59_)[index509_] = d_2766_v58_
                d_2768_v60_: str
                d_2768_v60_ = _dafny.CodePoint('o')
                d_2769_v61_: _dafny.Seq
                d_2769_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                d_2770_v62_: _dafny.Map
                d_2770_v62_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2768_v60_ for d_2771_i3_ in range(default__.abs(701))]): d_2766_v58_})
                index510_ = default__.safeIndex(664, (d_2767_v59_).length(0))
                (d_2767_v59_)[index510_] = (d_2766_v58_ if (d_2768_v60_) in (d_2769_v61_) else ((d_2770_v62_)[d_2769_v61_] if (d_2769_v61_) in (d_2770_v62_) else d_2766_v58_))
                d_2772_v63_: D14
                d_2772_v63_ = D14_DC38((self).f25)
                d_2773_v64_: _dafny.MultiSet
                d_2773_v64_ = _dafny.MultiSet([d_2757_cf0_])
                d_2772_v63_ = default__.fm70((d_2760_v54_)[default__.safeIndex(225, (d_2760_v54_).length(0))], ((self).f25) - ((d_2773_v64_).cardinality), (d_2760_v54_)[default__.safeIndex(225, (d_2760_v54_).length(0))], globalState)
                (globalState).f12 = (((self).f27 if p1 else -488)) <= ((self).f25)
            d_2774_v65_: C11
            nw455_ = C11()
            nw455_.ctor__()
            d_2774_v65_ = nw455_
            d_2774_v65_ = d_2774_v65_
        d_2775_v66_: _dafny.MultiSet
        d_2775_v66_ = _dafny.MultiSet([p1])
        (self).f28 = (d_2775_v66_) == ((d_2775_v66_) - (_dafny.MultiSet([not(self.f28)])))
        d_2776_v67_: _dafny.Seq
        d_2776_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igxf"))
        d_2777_v68_: _dafny.Set
        d_2777_v68_ = _dafny.Set({d_2776_v67_, d_2776_v67_})
        d_2778_v69_: _dafny.Seq
        d_2778_v69_ = _dafny.SeqWithoutIsStrInference([len(d_2776_v67_), (self).f27, len(d_2777_v68_)])
        d_2779_v70_: C6
        nw456_ = C6()
        nw456_.ctor__(len((d_2778_v69_) + (d_2778_v69_)), (self).f26)
        d_2779_v70_ = nw456_
        pat_let_tv66_ = d_2776_v67_
        pat_let_tv67_ = d_2776_v67_
        pat_let_tv68_ = d_2776_v67_
        pat_let_tv69_ = p1
        pat_let_tv70_ = d_2776_v67_
        pat_let_tv71_ = d_2776_v67_
        def lambda171_(source38_):
            if source38_.is_DC5:
                return _dafny.Map({pat_let_tv66_: self.f28})
            elif source38_.is_DC6:
                d_2780___mcc_h7_ = source38_.cf7
                d_2781___mcc_h8_ = source38_.cf8
                d_2782___mcc_h9_ = source38_.cf9
                d_2783_cf9_ = d_2782___mcc_h9_
                d_2784_cf8_ = d_2781___mcc_h8_
                d_2785_cf7_ = d_2780___mcc_h7_
                return _dafny.Map({pat_let_tv67_: self.f28})
            elif source38_.is_DC4:
                d_2786___mcc_h10_ = source38_.cf6
                d_2787_cf6_ = d_2786___mcc_h10_
                return _dafny.Map({pat_let_tv68_: False})
            elif True:
                d_2788___mcc_h11_ = source38_.cf10
                d_2789_cf10_ = d_2788___mcc_h11_
                if pat_let_tv69_:
                    return _dafny.Map({pat_let_tv70_: self.f28})
                elif True:
                    return _dafny.Map({pat_let_tv71_: False})

        source37_ = lambda171_(D1_DC4(d_2690_v0_))
        d_2790___mcc_h6_ = source37_
        d_2791_cf61_ = d_2790___mcc_h6_
        d_2792_v71_: _dafny.Map
        d_2792_v71_ = _dafny.Map({not (p1) or (p1): default__.safeDivisionInt((self).f27, (self).f27)})
        d_2792_v71_ = d_2792_v71_
        d_2793_v72_: T0
        nw457_ = C4()
        nw457_.ctor__((self).f26, (self).f25, (self).f26)
        d_2793_v72_ = nw457_
        d_2794_v73_: _dafny.Seq
        d_2794_v73_ = _dafny.SeqWithoutIsStrInference([d_2793_v72_])
        d_2795_v74_: _dafny.Set
        d_2795_v74_ = _dafny.Set({d_2794_v73_, d_2794_v73_, d_2794_v73_})
        d_2796_v75_: _dafny.Seq
        d_2796_v75_ = _dafny.SeqWithoutIsStrInference([d_2794_v73_])
        d_2797_v76_: _dafny.Array
        def lambda172_(d_2798_p1_):
            def lambda173_(d_2799_i4_):
                return d_2798_p1_

            return lambda173_

        init94_ = lambda172_(p1)
        nw458_ = _dafny.Array(None, 22)
        for i0_94_ in range(nw458_.length(0)):
            nw458_[i0_94_] = init94_(i0_94_)
        d_2797_v76_ = nw458_
        d_2800_v77_: C5
        nw459_ = C5()
        nw459_.ctor__(p1)
        d_2800_v77_ = nw459_
        d_2801_v78_: _dafny.Map
        d_2801_v78_ = _dafny.Map({d_2797_v76_: d_2800_v77_})
        d_2802_v79_: _dafny.Array
        nw460_ = _dafny.Array(None, 16)
        nw460_[int(0)] = _dafny.Set({d_2794_v73_})
        nw460_[int(1)] = d_2795_v74_
        nw460_[int(2)] = d_2795_v74_
        nw460_[int(3)] = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_2793_v72_])})
        nw460_[int(4)] = d_2795_v74_
        nw460_[int(5)] = (d_2795_v74_) | (d_2795_v74_)
        nw460_[int(6)] = (d_2795_v74_) - (d_2795_v74_)
        nw460_[int(7)] = d_2795_v74_
        nw460_[int(8)] = d_2795_v74_
        nw460_[int(9)] = (d_2795_v74_).intersection(d_2795_v74_)
        nw460_[int(10)] = d_2795_v74_
        nw460_[int(11)] = (_dafny.Set({d_2794_v73_, d_2794_v73_})) - (d_2795_v74_)
        nw460_[int(12)] = _dafny.Set({d_2794_v73_, (d_2796_v75_)[default__.safeIndex(819, len(d_2796_v75_))], d_2794_v73_, (_dafny.SeqWithoutIsStrInference([d_2793_v72_])).set(default__.safeIndex(len((d_2801_v78_).set(d_2797_v76_, d_2800_v77_)), len(_dafny.SeqWithoutIsStrInference([d_2793_v72_]))), d_2793_v72_), d_2794_v73_})
        nw460_[int(13)] = d_2795_v74_
        nw460_[int(14)] = (d_2795_v74_).intersection(d_2795_v74_)
        nw460_[int(15)] = _dafny.Set({d_2794_v73_, d_2794_v73_})
        d_2802_v79_ = nw460_
        index511_ = default__.safeIndex(818, (d_2802_v79_).length(0))
        (d_2802_v79_)[index511_] = (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_2793_v72_, d_2793_v72_, d_2793_v72_])}) if self.f28 else d_2795_v74_)
        d_2803_v80_: D4
        d_2803_v80_ = D4_DC13(777, False, (d_2800_v77_).f35, True)
        d_2804_v81_: _dafny.Map
        d_2804_v81_ = _dafny.Map({(d_2803_v80_).cf19: d_2797_v76_})
        d_2805_v82_: str
        d_2805_v82_ = _dafny.CodePoint('n')
        d_2806_v83_: _dafny.Map
        d_2806_v83_ = _dafny.Map({(self).f27: d_2805_v82_})
        d_2807_v84_: _dafny.Set
        d_2807_v84_ = _dafny.Set({d_2797_v76_})
        pat_let_tv72_ = d_2795_v74_
        index512_ = default__.safeIndex(818, (d_2802_v79_).length(0))
        def iife221_(_pat_let64_0):
            def iife222_(d_2808_dt__update__tmp_h1_):
                def iife223_(_pat_let65_0):
                    def iife224_(d_2809_dt__update_hcf76_h0_):
                        return D21_DC51(d_2809_dt__update_hcf76_h0_)
                    return iife224_(_pat_let65_0)
                return iife223_(pat_let_tv72_)
            return iife222_(_pat_let64_0)
        rhs489_ = len(((d_2804_v81_).set((d_2793_v72_).f25, d_2797_v76_)) | (_dafny.Map({len(d_2806_v83_): d_2797_v76_})))
        rhs490_ = default__.safeDivisionInt((self).f25, len(d_2807_v84_))
        rhs491_ = (iife221_(D21_DC51(d_2795_v74_))).cf76
        lhs377_ = d_2802_v79_
        lhs378_ = default__.safeIndex(818, (d_2802_v79_).length(0))
        r0 = rhs489_
        r0 = rhs490_
        lhs377_[lhs378_] = rhs491_
        (globalState).f6 = d_2778_v69_
        d_2810_v85_: D3
        d_2810_v85_ = D3_DC9(d_2807_v84_)
        d_2811_v86_: _dafny.Map
        d_2811_v86_ = _dafny.Map({(self).f27: d_2810_v85_})
        d_2812_v87_: D19
        d_2812_v87_ = D19_DC46(d_2811_v86_)
        source39_ = d_2812_v87_
        if source39_.is_DC47:
            d_2813___mcc_h12_ = source39_.cf69
            d_2814___mcc_h13_ = source39_.cf70
            d_2815_cf70_ = d_2814___mcc_h13_
            d_2816_cf69_ = d_2813___mcc_h12_
            d_2817_v88_: _dafny.Set
            d_2817_v88_ = _dafny.Set({(d_2778_v69_).set(default__.safeIndex((self).f25, len(d_2778_v69_)), (self).f25)})
            (globalState).f12 = (not(((d_2793_v72_).f25) < ((d_2793_v72_).f25)) if (d_2817_v88_).isdisjoint(_dafny.Set({d_2778_v69_, _dafny.SeqWithoutIsStrInference([len(d_2776_v67_) for d_2818_i5_ in range(default__.abs(504))]), d_2778_v69_, d_2778_v69_})) else self.f28)
            index513_ = default__.safeIndex(928, (d_2816_cf69_).length(0))
            (d_2816_cf69_)[index513_] = (d_2800_v77_).f35
            index514_ = default__.safeIndex(928, (d_2816_cf69_).length(0))
            (d_2816_cf69_)[index514_] = default__.fm0((d_2793_v72_).f25, globalState)
            (globalState).f12 = (d_2800_v77_).f35
            d_2815_cf70_ = default__.fm0(((d_2793_v72_).f25) - ((d_2793_v72_).f25), globalState)
        elif source39_.is_DC48:
            d_2819___mcc_h14_ = source39_.cf71
            d_2820_cf71_ = d_2819___mcc_h14_
            d_2690_v0_ = d_2690_v0_
            index515_ = default__.safeIndex(624, (d_2797_v76_).length(0))
            (d_2797_v76_)[index515_] = ((d_2800_v77_).f35 if self.f28 else p1)
            d_2821_v89_: _dafny.Set
            d_2821_v89_ = _dafny.Set({(self).f25, d_2820_cf71_})
            index516_ = default__.safeIndex(624, (d_2797_v76_).length(0))
            (d_2797_v76_)[index516_] = (d_2821_v89_).ispropersubset(d_2821_v89_)
            (globalState).f12 = (d_2797_v76_)[default__.safeIndex(624, (d_2797_v76_).length(0))]
            d_2822_v90_: _dafny.Array
            def lambda174_(d_2823_i6_):
                return D4_DC12()

            init95_ = lambda174_
            nw461_ = _dafny.Array(None, 28)
            for i0_95_ in range(nw461_.length(0)):
                nw461_[i0_95_] = init95_(i0_95_)
            d_2822_v90_ = nw461_
            d_2824_v91_: D4
            d_2824_v91_ = D4_DC12()
            index517_ = default__.safeIndex(31, (d_2822_v90_).length(0))
            (d_2822_v90_)[index517_] = d_2824_v91_
            d_2825_v92_: _dafny.Map
            d_2825_v92_ = _dafny.Map({d_2690_v0_: (d_2793_v72_).f25})
            d_2826_v93_: _dafny.Array
            nw462_ = _dafny.Array(None, 3)
            nw462_[int(0)] = (d_2800_v77_).f35
            nw462_[int(1)] = True
            nw462_[int(2)] = (d_2800_v77_).f35
            d_2826_v93_ = nw462_
            d_2827_v94_: _dafny.Array
            nw463_ = _dafny.Array(None, 16)
            nw463_[int(0)] = (self).f25
            nw463_[int(1)] = (self).f25
            nw463_[int(2)] = (d_2793_v72_).f25
            nw463_[int(3)] = default__.fm18(globalState)
            nw463_[int(4)] = d_2820_cf71_
            nw463_[int(5)] = (d_2793_v72_).f25
            nw463_[int(6)] = len(d_2776_v67_)
            nw463_[int(7)] = ((0) - ((self).f25)) + ((d_2793_v72_).f25)
            nw463_[int(8)] = d_2820_cf71_
            nw463_[int(9)] = (d_2793_v72_).f25
            nw463_[int(10)] = d_2820_cf71_
            nw463_[int(11)] = (d_2775_v66_).cardinality
            nw463_[int(12)] = ((d_2793_v72_).f25) * ((self).f27)
            nw463_[int(13)] = ((d_2825_v92_)[d_2690_v0_] if (d_2690_v0_) in (d_2825_v92_) else len(_dafny.Map({d_2826_v93_: (0) - (d_2820_cf71_)})))
            nw463_[int(14)] = 79
            nw463_[int(15)] = len((d_2821_v89_).intersection(d_2821_v89_))
            d_2827_v94_ = nw463_
            index518_ = default__.safeIndex(31, (d_2822_v90_).length(0))
            rhs492_ = d_2805_v82_
            rhs493_ = D4_DC12()
            rhs494_ = (p0)[default__.safeIndex((d_2793_v72_).f25, len(p0))]
            lhs379_ = globalState
            lhs380_ = d_2822_v90_
            lhs381_ = default__.safeIndex(31, (d_2822_v90_).length(0))
            lhs379_.f23 = rhs492_
            lhs380_[lhs381_] = rhs493_
            d_2827_v94_ = rhs494_
        elif True:
            d_2828___mcc_h15_ = source39_.cf68
            d_2829_cf68_ = d_2828___mcc_h15_
            d_2830_v95_: _dafny.Map
            d_2830_v95_ = _dafny.Map({((d_2792_v71_)[(d_2800_v77_).f35] if ((d_2800_v77_).f35) in (d_2792_v71_) else (self).f25): _dafny.SeqWithoutIsStrInference([(d_2800_v77_).f35])})
            index519_ = default__.safeIndex(438, (d_2797_v76_).length(0))
            (d_2797_v76_)[index519_] = (((d_2830_v95_)[(d_2793_v72_).f25] if ((d_2793_v72_).f25) in (d_2830_v95_) else _dafny.SeqWithoutIsStrInference([p1, self.f28]))) <= (d_2690_v0_)
            index520_ = default__.safeIndex(438, (d_2797_v76_).length(0))
            (d_2797_v76_)[index520_] = self.f28
            index521_ = default__.safeIndex(438, (d_2797_v76_).length(0))
            (d_2797_v76_)[index521_] = p1
            r0 = (default__.safeDivisionInt(-200, len(d_2776_v67_)) if (d_2800_v77_).f35 else (d_2793_v72_).f25)
            d_2831_v96_: C0
            nw464_ = C0()
            nw464_.ctor__()
            d_2831_v96_ = nw464_
        if not((d_2776_v67_) <= ((d_2776_v67_ if p1 else default__.fm2(551, (self).f27, globalState)))):
            d_2832_v97_: str
            d_2832_v97_ = _dafny.CodePoint('y')
            (globalState).f10 = (D1_DC6(d_2832_v97_, (d_2690_v0_)[default__.safeIndex((self).f27, len(d_2690_v0_))], d_2776_v67_)).cf9
            (globalState).f12 = not(p1)
            d_2833_v98_: _dafny.Map
            d_2833_v98_ = _dafny.Map({self.f28: self.f28})
            d_2834_v99_: D21
            d_2834_v99_ = D21_DC52(default__.safeDivisionInt(164, len(d_2833_v98_)), ((0) - ((self).f25)) * ((self).f25), d_2776_v67_, self.f28, (self).f27)
            pat_let_tv73_ = d_2776_v67_
            pat_let_tv74_ = d_2776_v67_
            def iife225_(_pat_let66_0):
                def iife226_(d_2835_dt__update__tmp_h2_):
                    def iife227_(_pat_let67_0):
                        def iife228_(d_2836_dt__update_hcf80_h0_):
                            def iife229_(_pat_let68_0):
                                def iife230_(d_2837_dt__update_hcf79_h0_):
                                    def iife231_(_pat_let69_0):
                                        def iife232_(d_2838_dt__update_hcf81_h0_):
                                            return D21_DC52((d_2835_dt__update__tmp_h2_).cf77, (d_2835_dt__update__tmp_h2_).cf78, d_2837_dt__update_hcf79_h0_, d_2836_dt__update_hcf80_h0_, d_2838_dt__update_hcf81_h0_)
                                        return iife232_(_pat_let69_0)
                                    return iife231_(default__.safeDivisionInt((self).f27, (self).f25))
                                return iife230_(_pat_let68_0)
                            return iife229_((pat_let_tv73_) + (pat_let_tv74_))
                        return iife228_(_pat_let67_0)
                    return iife227_(not(False))
                return iife226_(_pat_let66_0)
            d_2834_v99_ = iife225_(d_2834_v99_)
            d_2839_v100_: _dafny.Map
            d_2839_v100_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))): _dafny.Set({self.f28})})
            d_2840_v101_: _dafny.Set
            d_2840_v101_ = _dafny.Set({self.f28, self.f28, False})
            d_2841_v102_: _dafny.Map
            d_2841_v102_ = _dafny.Map({d_2832_v97_: p1})
            d_2842_v103_: _dafny.Set
            d_2842_v103_ = _dafny.Set({(self).f25, (self).f25, (0) - (len((d_2839_v100_).set((self).f25, d_2840_v101_))), len(d_2841_v102_), (self).f25})
            d_2843_v104_: D4
            d_2843_v104_ = D4_DC11(d_2842_v103_)
            d_2843_v104_ = d_2843_v104_
            (globalState).f12 = p1
        elif True:
            d_2844_v105_: C7
            nw465_ = C7()
            nw465_.ctor__((self).f27)
            d_2844_v105_ = nw465_
            d_2845_v106_: _dafny.Map
            d_2845_v106_ = _dafny.Map({(d_2844_v105_).f34: True})
            d_2846_v107_: _dafny.Set
            d_2846_v107_ = _dafny.Set({not(False), p1, ((d_2845_v106_)[(d_2844_v105_).f34] if ((d_2844_v105_).f34) in (d_2845_v106_) else p1), self.f28, False})
            d_2847_v108_: _dafny.Map
            d_2847_v108_ = _dafny.Map({default__.safeDivisionInt((self).f25, (d_2844_v105_).f34): (len(default__.fm2((self).f25, len(d_2690_v0_), globalState))) * (len(d_2846_v107_))})
            d_2847_v108_ = (d_2847_v108_).set((190) * ((self).f27), default__.safeDivisionInt((self).f25, (d_2844_v105_).f34))
            d_2848_v109_: _dafny.Array
            nw466_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_2848_v109_ = nw466_
            d_2849_v110_: _dafny.Array
            def lambda175_(d_2850_p1_, d_2851_v105_):
                def lambda176_(d_2852_i7_):
                    return (D13_DC36(_dafny.Map({False: d_2850_p1_}), self.f28, (d_2851_v105_).f34)).cf55

                return lambda176_

            init96_ = lambda175_(p1, d_2844_v105_)
            nw467_ = _dafny.Array(None, 23)
            for i0_96_ in range(nw467_.length(0)):
                nw467_[i0_96_] = init96_(i0_96_)
            d_2849_v110_ = nw467_
            index522_ = default__.safeIndex(986, (d_2848_v109_).length(0))
            (d_2848_v109_)[index522_] = d_2849_v110_
            d_2853_v111_: _dafny.Map
            d_2853_v111_ = _dafny.Map({p1: p1})
            index523_ = default__.safeIndex(986, (d_2848_v109_).length(0))
            (d_2848_v109_)[index523_] = (D3_DC10(p1, d_2849_v110_, ((d_2853_v111_)[p1] if (p1) in (d_2853_v111_) else self.f28), (d_2844_v105_).f34, (self).f27)).cf14
            d_2854_v112_: C4
            nw468_ = C4()
            nw468_.ctor__((self).f26, len(d_2846_v107_), (self).f26)
            d_2854_v112_ = nw468_
            d_2855_v113_: _dafny.Set
            d_2855_v113_ = _dafny.Set({d_2847_v108_, (d_2847_v108_).set(-244, (0) - (len(d_2776_v67_)))})
            d_2856_v114_: _dafny.Array
            nw469_ = _dafny.Array(int(0), 10)
            d_2856_v114_ = nw469_
            d_2857_v115_: D21
            d_2857_v115_ = D21_DC53(d_2855_v113_, p1, d_2856_v114_)
            d_2858_v116_: _dafny.Map
            d_2858_v116_ = _dafny.Map({False: d_2857_v115_})
            d_2859_v118_: str
            d_2859_v118_ = _dafny.CodePoint('u')
            d_2860_v119_: _dafny.Map
            d_2860_v119_ = _dafny.Map({p1: d_2859_v118_})
            d_2861_v120_: _dafny.Array
            nw470_ = _dafny.Array(None, 17)
            nw470_[int(0)] = d_2847_v108_
            nw470_[int(1)] = (d_2847_v108_) | (default__.fm57(p1, p1, (self).f25, globalState))
            nw470_[int(2)] = d_2847_v108_
            nw470_[int(3)] = d_2847_v108_
            nw470_[int(4)] = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqkjki")))): (self).f27})
            nw470_[int(5)] = d_2847_v108_
            nw470_[int(6)] = d_2847_v108_
            nw470_[int(7)] = _dafny.Map({len(d_2776_v67_): len(d_2858_v116_)})
            nw470_[int(8)] = default__.fm57(False, p1, (self).f25, globalState)
            nw470_[int(9)] = d_2847_v108_
            nw470_[int(10)] = (d_2847_v108_) | (d_2847_v108_)
            nw470_[int(11)] = d_2847_v108_
            nw470_[int(12)] = d_2847_v108_
            def iife233_():
                coll93_ = _dafny.Map()
                compr_93_: int
                for compr_93_ in _dafny.IntegerRange(231, 498):
                    d_2862_v117_: int = compr_93_
                    if ((231) <= (d_2862_v117_)) and ((d_2862_v117_) < (498)):
                        coll93_[(d_2862_v117_) - (-424)] = 398
                return _dafny.Map(coll93_)
            nw470_[int(13)] = iife233_()
            
            nw470_[int(14)] = d_2847_v108_
            nw470_[int(15)] = ((d_2847_v108_).set(len(_dafny.Map({len(d_2860_v119_): p1})), (d_2775_v66_).cardinality)) | (d_2847_v108_)
            nw470_[int(16)] = _dafny.Map({(self).f27: (d_2844_v105_).f34})
            d_2861_v120_ = nw470_
            index524_ = default__.safeIndex(278, (d_2861_v120_).length(0))
            (d_2861_v120_)[index524_] = _dafny.Map({(self).f25: (self).f27})
            d_2863_v121_: _dafny.Map
            d_2863_v121_ = _dafny.Map({(0) - ((d_2844_v105_).f34): d_2776_v67_})
            d_2864_v122_: _dafny.Map
            d_2864_v122_ = _dafny.Map({default__.fm44(self.f28, self.f28, False, (d_2844_v105_).f34, globalState): d_2849_v110_})
            d_2865_v123_: D8
            d_2865_v123_ = D8_DC23(_dafny.Map({((d_2864_v122_)[(self).f27] if ((self).f27) in (d_2864_v122_) else (d_2848_v109_)[default__.safeIndex(986, (d_2848_v109_).length(0))]): p1}))
            d_2866_v124_: D8
            d_2866_v124_ = D8_DC25(d_2865_v123_)
            d_2867_v125_: D8
            d_2867_v125_ = D8_DC25(d_2866_v124_)
            pat_let_tv75_ = d_2865_v123_
            d_2868_v126_: _dafny.Seq
            def iife234_(_pat_let70_0):
                def iife235_(d_2869_dt__update__tmp_h3_):
                    def iife236_(_pat_let71_0):
                        def iife237_(d_2870_dt__update_hcf37_h0_):
                            return D8_DC25(d_2870_dt__update_hcf37_h0_)
                        return iife237_(_pat_let71_0)
                    return iife236_(pat_let_tv75_)
                return iife235_(_pat_let70_0)
            d_2868_v126_ = _dafny.SeqWithoutIsStrInference([D8_DC25(d_2865_v123_), d_2867_v125_, d_2867_v125_, iife234_(d_2867_v125_)])
            d_2871_v128_: _dafny.MultiSet
            d_2871_v128_ = _dafny.MultiSet([len((d_2776_v67_).set(default__.safeIndex((d_2844_v105_).f34, len(d_2776_v67_)), d_2859_v118_)), (self).f27])
            d_2872_v129_: _dafny.Map
            d_2872_v129_ = _dafny.Map({942: d_2871_v128_})
            index525_ = default__.safeIndex(278, (d_2861_v120_).length(0))
            def iife238_():
                coll94_ = _dafny.Map()
                compr_94_: int
                for compr_94_ in _dafny.IntegerRange(937, 61):
                    d_2873_v127_: int = compr_94_
                    if ((937) <= (d_2873_v127_)) and ((d_2873_v127_) < (61)):
                        coll94_[(d_2873_v127_) - ((d_2844_v105_).f34)] = len(_dafny.Map({False: len(d_2690_v0_)}))
                return _dafny.Map(coll94_)
            rhs495_ = iife238_()

            rhs496_ = (self).f27
            rhs497_ = ((_dafny.Map({(0) - ((self).fm8(p1, True, globalState)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lf"))})) | (d_2863_v121_)).set((0) - ((self).f27), (d_2776_v67_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))))
            rhs498_ = ((d_2868_v126_) + (d_2868_v126_)) + (d_2868_v126_)
            rhs499_ = not (p1) or ((len((default__.fm71(((d_2872_v129_)[(self).f27] if ((self).f27) in (d_2872_v129_) else _dafny.MultiSet(d_2778_v69_)), globalState)).cf6)) <= ((self).f25))
            lhs382_ = d_2861_v120_
            lhs383_ = default__.safeIndex(278, (d_2861_v120_).length(0))
            lhs384_ = globalState
            lhs382_[lhs383_] = rhs495_
            r0 = rhs496_
            d_2863_v121_ = rhs497_
            d_2868_v126_ = rhs498_
            lhs384_.f12 = rhs499_
        d_2874_v130_: _dafny.Array
        nw471_ = _dafny.Array(_dafny.CodePoint('D'), 24)
        d_2874_v130_ = nw471_
        d_2875_v131_: _dafny.Map
        d_2875_v131_ = _dafny.Map({d_2874_v130_: ((self).f27) < ((self).f27)})
        d_2875_v131_ = (d_2875_v131_).set(d_2874_v130_, (True) and (False))
        r0 = -857
        return r0

    def m5(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2876_v0_: str
        d_2876_v0_ = _dafny.CodePoint('u')
        (globalState).f23 = d_2876_v0_
        d_2877_v1_: _dafny.Array
        def lambda177_(d_2878_i0_):
            return self.f28

        init97_ = lambda177_
        nw472_ = _dafny.Array(None, 14)
        for i0_97_ in range(nw472_.length(0)):
            nw472_[i0_97_] = init97_(i0_97_)
        d_2877_v1_ = nw472_
        d_2879_v2_: _dafny.Set
        d_2879_v2_ = _dafny.Set({d_2876_v0_})
        index526_ = default__.safeIndex(187, (d_2877_v1_).length(0))
        (d_2877_v1_)[index526_] = not((d_2879_v2_).issubset(d_2879_v2_))
        d_2880_v3_: int
        d_2880_v3_ = 17
        d_2881_v4_: _dafny.Map
        d_2881_v4_ = _dafny.Map({p3: self.f28})
        d_2882_v5_: D4
        d_2882_v5_ = D4_DC13(d_2880_v3_, p3, p3, self.f28)
        d_2883_v6_: _dafny.Map
        d_2883_v6_ = _dafny.Map({(self).f27: False})
        d_2884_v7_: _dafny.Seq
        d_2884_v7_ = _dafny.SeqWithoutIsStrInference([len(d_2883_v6_), d_2880_v3_])
        d_2885_v8_: _dafny.MultiSet
        d_2885_v8_ = _dafny.MultiSet([d_2876_v0_])
        d_2886_v9_: _dafny.Map
        d_2886_v9_ = _dafny.Map({len(d_2884_v7_): d_2885_v8_})
        d_2887_v10_: _dafny.MultiSet
        d_2887_v10_ = _dafny.MultiSet([(((d_2886_v9_)[d_2880_v3_] if (d_2880_v3_) in (d_2886_v9_) else d_2885_v8_)).cardinality, d_2880_v3_, (self).f25])
        d_2888_v11_: _dafny.Seq
        d_2888_v11_ = _dafny.SeqWithoutIsStrInference([((d_2881_v4_)[(d_2882_v5_).cf21] if ((d_2882_v5_).cf21) in (d_2881_v4_) else p1), (d_2876_v0_) not in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2889_i1_ in range(default__.abs(-932))])), (d_2887_v10_).isdisjoint(_dafny.MultiSet([-893]))])
        index527_ = default__.safeIndex(187, (d_2877_v1_).length(0))
        rhs500_ = (d_2880_v3_) > ((self).f27)
        rhs501_ = (self).f27
        rhs502_ = (d_2888_v11_)[default__.safeIndex(802, len(d_2888_v11_))]
        rhs503_ = (0) - (d_2880_v3_)
        lhs385_ = d_2877_v1_
        lhs386_ = default__.safeIndex(187, (d_2877_v1_).length(0))
        lhs385_[lhs386_] = rhs500_
        d_2880_v3_ = rhs501_
        r1 = rhs502_
        d_2880_v3_ = rhs503_
        if (p1) or (((0) - (d_2880_v3_)) > ((self).f25)):
            d_2890_v12_: _dafny.MultiSet
            d_2890_v12_ = _dafny.MultiSet([p3, p1, self.f28])
            d_2881_v4_ = (d_2881_v4_).set((d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))], (_dafny.MultiSet([(d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))], p1])).ispropersubset(d_2890_v12_))
            (globalState).f12 = not(p1)
            index528_ = default__.safeIndex(187, (d_2877_v1_).length(0))
            (d_2877_v1_)[index528_] = not((p3) or (((d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))] if p3 else self.f28)))
            index529_ = default__.safeIndex(362, (p0).length(0))
            (p0)[index529_] = d_2880_v3_
            d_2891_v13_: _dafny.Map
            d_2891_v13_ = _dafny.Map({(self).f25: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2892_i2_ in range(default__.abs(271))])})
            d_2893_v14_: _dafny.Set
            d_2893_v14_ = _dafny.Set({(self).f27})
            d_2894_v15_: _dafny.Map
            d_2894_v15_ = _dafny.Map({(self).f27: (0) - ((0) - (d_2880_v3_))})
            d_2895_v16_: _dafny.Map
            d_2895_v16_ = _dafny.Map({default__.fm44((d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))], not(p3), default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cojfkt"))), globalState), d_2880_v3_, globalState): (0) - (((d_2894_v15_)[482] if (482) in (d_2894_v15_) else 799))})
            d_2896_v17_: _dafny.Seq
            d_2896_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqj"))
            d_2897_v18_: _dafny.Set
            d_2897_v18_ = _dafny.Set({p3})
            d_2898_v19_: _dafny.Map
            d_2898_v19_ = _dafny.Map({d_2897_v18_: default__.fm44(True, (d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))], self.f28, d_2880_v3_, globalState)})
            index530_ = default__.safeIndex(362, (p0).length(0))
            rhs504_ = (self).f25
            rhs505_ = (d_2893_v14_).issubset(default__.fm27(p1, d_2891_v13_, globalState))
            rhs506_ = ((d_2894_v15_)[default__.safeModuloInt(len(d_2896_v17_), (self).f25)] if (default__.safeModuloInt(len(d_2896_v17_), (self).f25)) in (d_2894_v15_) else len((d_2898_v19_) | (d_2898_v19_)))
            rhs507_ = True
            lhs387_ = p0
            lhs388_ = default__.safeIndex(362, (p0).length(0))
            lhs389_ = globalState
            lhs390_ = self
            lhs387_[lhs388_] = rhs504_
            lhs389_.f12 = rhs505_
            d_2880_v3_ = rhs506_
            lhs390_.f28 = rhs507_
            d_2899_v20_: _dafny.Map
            d_2899_v20_ = _dafny.Map({d_2876_v0_: (self).f27})
            d_2899_v20_ = (d_2899_v20_).set(d_2876_v0_, (self).f25)
        elif True:
            d_2900_v21_: _dafny.Seq
            d_2900_v21_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2901_v22_: _dafny.MultiSet
            d_2901_v22_ = _dafny.MultiSet([p3, (len(_dafny.Set({(d_2900_v21_)[default__.safeIndex((self).f27, len(d_2900_v21_))]}))) == ((self).f27), True])
            d_2901_v22_ = (d_2901_v22_) - ((d_2901_v22_) - (d_2901_v22_))
            d_2902_v23_: _dafny.Array
            def lambda178_(d_2903_v0_):
                def lambda179_(d_2904_i3_):
                    return default__.safeModuloInt(d_2904_i3_, len(_dafny.SeqWithoutIsStrInference([d_2903_v0_ for d_2905_i4_ in range(default__.abs(499))])))

                return lambda179_

            init98_ = lambda178_(d_2876_v0_)
            nw473_ = _dafny.Array(None, 12)
            for i0_98_ in range(nw473_.length(0)):
                nw473_[i0_98_] = init98_(i0_98_)
            d_2902_v23_ = nw473_
            d_2902_v23_ = p0
            if (d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]:
                d_2906_v24_: _dafny.Map
                d_2906_v24_ = _dafny.Map({p1: d_2876_v0_})
                d_2907_v25_: C9
                nw474_ = C9()
                nw474_.ctor__(p0, d_2880_v3_)
                d_2907_v25_ = nw474_
                d_2908_v26_: _dafny.Map
                d_2908_v26_ = _dafny.Map({(d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]: d_2907_v25_})
                d_2909_v27_: _dafny.Map
                d_2909_v27_ = _dafny.Map({d_2906_v24_: d_2908_v26_})
                d_2910_v29_: _dafny.Map
                def iife239_():
                    coll95_ = _dafny.Set()
                    compr_95_: _dafny.Seq
                    for compr_95_ in (default__.fm72(p3, p3, (d_2907_v25_).f31, d_2880_v3_, globalState)).Elements:
                        d_2911_v28_: _dafny.Seq = compr_95_
                        if (d_2911_v28_) in (default__.fm72(p3, p3, (d_2907_v25_).f31, d_2880_v3_, globalState)):
                            coll95_ = coll95_.union(_dafny.Set([d_2911_v28_]))
                    return _dafny.Set(coll95_)
                d_2910_v29_ = _dafny.Map({(d_2909_v27_).set(d_2906_v24_, d_2908_v26_): iife239_()
                })
                d_2912_v30_: _dafny.Set
                d_2912_v30_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f27 for d_2913_i5_ in range(default__.abs(-868))])})
                d_2910_v29_ = (d_2910_v29_).set((_dafny.Map({d_2906_v24_: d_2908_v26_})) | (_dafny.Map({d_2906_v24_: d_2908_v26_})), d_2912_v30_)
                (globalState).f12 = p1
                (self).f28 = not(not((d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]))
                d_2914_v31_: _dafny.Seq
                d_2914_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjfg"))
                rhs508_ = (d_2914_v31_ if (d_2888_v11_)[default__.safeIndex(d_2880_v3_, len(d_2888_v11_))] else d_2914_v31_)
                rhs509_ = (d_2907_v25_).f31
                rhs510_ = len((d_2914_v31_) + (_dafny.SeqWithoutIsStrInference([d_2876_v0_ for d_2915_i6_ in range(default__.abs(-750))])))
                r2 = rhs508_
                d_2880_v3_ = rhs509_
                d_2880_v3_ = rhs510_
                index531_ = default__.safeIndex(18, (d_2902_v23_).length(0))
                (d_2902_v23_)[index531_] = -852
                d_2916_v32_: _dafny.Map
                d_2916_v32_ = _dafny.Map({(d_2907_v25_).f31: (self).f25})
                index532_ = default__.safeIndex(18, (d_2902_v23_).length(0))
                (d_2902_v23_)[index532_] = ((self).f27) * (((d_2907_v25_).f31) + (len(d_2916_v32_)))
            elif True:
                index533_ = default__.safeIndex(187, (d_2877_v1_).length(0))
                (d_2877_v1_)[index533_] = not(not(self.f28))
                d_2880_v3_ = len(default__.fm2((self).f27, (self).f25, globalState))
                d_2880_v3_ = d_2880_v3_
                r0 = not(not(self.f28))
                d_2917_v33_: D3
                d_2917_v33_ = D3_DC10((d_2888_v11_)[default__.safeIndex(d_2880_v3_, len(d_2888_v11_))], d_2877_v1_, self.f28, (self).f25, d_2880_v3_)
                (globalState).f12 = (d_2917_v33_).cf15
            if True:
                d_2918_v34_: _dafny.Array
                def lambda180_(d_2919_v22_):
                    def lambda181_(d_2920_i7_):
                        return D15_DC39(d_2919_v22_)

                    return lambda181_

                init99_ = lambda180_(d_2901_v22_)
                nw475_ = _dafny.Array(None, 24)
                for i0_99_ in range(nw475_.length(0)):
                    nw475_[i0_99_] = init99_(i0_99_)
                d_2918_v34_ = nw475_
                d_2921_v35_: _dafny.Seq
                d_2921_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                rhs511_ = d_2918_v34_
                rhs512_ = d_2880_v3_
                rhs513_ = d_2880_v3_
                rhs514_ = (d_2876_v0_) in (d_2921_v35_)
                rhs515_ = (self).f25
                d_2918_v34_ = rhs511_
                d_2880_v3_ = rhs512_
                d_2880_v3_ = rhs513_
                r0 = rhs514_
                d_2880_v3_ = rhs515_
                d_2922_v36_: D14
                d_2922_v36_ = D14_DC38(((self).f27) + (868))
                d_2923_v37_: _dafny.MultiSet
                d_2923_v37_ = _dafny.MultiSet([d_2887_v10_, _dafny.MultiSet(d_2884_v7_), d_2887_v10_])
                d_2924_v38_: _dafny.Map
                d_2924_v38_ = _dafny.Map({p1: d_2923_v37_})
                pat_let_tv76_ = d_2881_v4_
                def iife241_(_pat_let73_0):
                    def iife242_(d_2925_dt__update__tmp_h0_):
                        def iife243_(_pat_let74_0):
                            def iife244_(d_2926_dt__update_hcf58_h0_):
                                return D14_DC38(d_2926_dt__update_hcf58_h0_)
                            return iife244_(_pat_let74_0)
                        return iife243_(len(pat_let_tv76_))
                    return iife242_(_pat_let73_0)
                def iife240_(_pat_let72_0):
                    def iife245_(d_2927_dt__update__tmp_h1_):
                        def iife246_(_pat_let75_0):
                            def iife247_(d_2928_dt__update_hcf58_h1_):
                                return D14_DC38(d_2928_dt__update_hcf58_h1_)
                            return iife247_(_pat_let75_0)
                        return iife246_((self).f27)
                    return iife245_(_pat_let72_0)
                rhs516_ = iife240_(iife241_(d_2922_v36_))
                rhs517_ = (d_2880_v3_) + ((self).f25)
                rhs518_ = ((((d_2924_v38_)[(d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]] if ((d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]) in (d_2924_v38_) else _dafny.MultiSet([d_2887_v10_, _dafny.MultiSet(d_2884_v7_), d_2887_v10_]))).cardinality) == ((d_2880_v3_) - (810))
                rhs519_ = ((self).f25) - ((self).f27)
                rhs520_ = (d_2900_v21_)[default__.safeIndex(((d_2884_v7_)[default__.safeIndex(d_2880_v3_, len(d_2884_v7_))]) + (d_2880_v3_), len(d_2900_v21_))]
                d_2922_v36_ = rhs516_
                d_2880_v3_ = rhs517_
                r0 = rhs518_
                d_2880_v3_ = rhs519_
                d_2902_v23_ = rhs520_
                d_2880_v3_ = len((d_2888_v11_) + (d_2888_v11_))
                d_2929_v39_: _dafny.Map
                d_2929_v39_ = _dafny.Map({601: (self).f27})
                d_2930_v40_: D14
                d_2930_v40_ = D14_DC37(d_2929_v39_)
                d_2931_v41_: _dafny.Set
                d_2931_v41_ = _dafny.Set({p1})
                pat_let_tv77_ = d_2931_v41_
                def iife248_(_pat_let76_0):
                    def iife249_(d_2932_dt__update__tmp_h2_):
                        def iife250_(_pat_let77_0):
                            def iife251_(d_2933_dt__update_hcf57_h0_):
                                return D14_DC37(d_2933_dt__update_hcf57_h0_)
                            return iife251_(_pat_let77_0)
                        return iife250_(_dafny.Map({len(pat_let_tv77_): (self).f25}))
                    return iife249_(_pat_let76_0)
                d_2930_v40_ = iife248_(D14_DC37(d_2929_v39_))
                (globalState).f12 = (default__.fm73(p1, globalState)).cf47
            elif True:
                d_2934_v42_: _dafny.Map
                d_2934_v42_ = _dafny.Map({p1: d_2876_v0_})
                d_2935_v43_: _dafny.Map
                d_2935_v43_ = _dafny.Map({len(d_2934_v42_): (self).f25})
                d_2936_v44_: _dafny.Set
                d_2936_v44_ = _dafny.Set({d_2935_v43_})
                d_2937_v45_: D21
                d_2937_v45_ = D21_DC53(d_2936_v44_, default__.fm0((self).f25, globalState), p0)
                d_2938_v46_: D7
                d_2938_v46_ = D7_DC21(d_2902_v23_)
                d_2939_v47_: _dafny.Array
                nw476_ = _dafny.Array(None, 25)
                nw476_[int(0)] = d_2902_v23_
                nw476_[int(1)] = d_2902_v23_
                nw476_[int(2)] = (d_2937_v45_).cf84
                nw476_[int(3)] = p0
                nw476_[int(4)] = p0
                nw476_[int(5)] = p0
                nw476_[int(6)] = d_2902_v23_
                nw476_[int(7)] = (d_2938_v46_).cf32
                nw476_[int(8)] = p0
                nw476_[int(9)] = p0
                nw476_[int(10)] = d_2902_v23_
                nw476_[int(11)] = d_2902_v23_
                nw476_[int(12)] = d_2902_v23_
                nw476_[int(13)] = p0
                nw476_[int(14)] = d_2902_v23_
                nw476_[int(15)] = d_2902_v23_
                nw476_[int(16)] = p0
                nw476_[int(17)] = d_2902_v23_
                nw476_[int(18)] = d_2902_v23_
                nw476_[int(19)] = d_2902_v23_
                nw476_[int(20)] = d_2902_v23_
                nw476_[int(21)] = p0
                nw476_[int(22)] = d_2902_v23_
                nw476_[int(23)] = p0
                nw476_[int(24)] = d_2902_v23_
                d_2939_v47_ = nw476_
                d_2940_v48_: D15
                d_2940_v48_ = D15_DC39(_dafny.MultiSet([(d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))], (d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]]))
                d_2941_v49_: _dafny.Map
                d_2941_v49_ = _dafny.Map({d_2940_v48_: d_2939_v47_})
                d_2939_v47_ = ((d_2941_v49_)[D15_DC39(d_2901_v22_)] if (D15_DC39(d_2901_v22_)) in (d_2941_v49_) else d_2939_v47_)
                d_2942_v50_: _dafny.Array
                d_2943_v51_: _dafny.Array
                d_2944_v52_: bool
                out46_: _dafny.Array
                out47_: _dafny.Array
                out48_: bool
                out46_, out47_, out48_ = default__.m0(globalState)
                d_2942_v50_ = out46_
                d_2943_v51_ = out47_
                d_2944_v52_ = out48_
                index534_ = default__.safeIndex(988, (p0).length(0))
                (p0)[index534_] = len(d_2888_v11_)
                index535_ = default__.safeIndex(988, (p0).length(0))
                (p0)[index535_] = (self).f25
                index536_ = default__.safeIndex(988, (p0).length(0))
                rhs521_ = default__.fm44(False, not(self.f28), False, default__.fm18(globalState), globalState)
                rhs522_ = (self).f25
                rhs523_ = d_2888_v11_
                lhs391_ = p0
                lhs392_ = default__.safeIndex(988, (p0).length(0))
                lhs391_[lhs392_] = rhs521_
                d_2880_v3_ = rhs522_
                d_2888_v11_ = rhs523_
                (self).f28 = default__.fm0(((p0)[default__.safeIndex(988, (p0).length(0))]) + ((self).f27), globalState)
            if p1:
                d_2945_v53_: C7
                nw477_ = C7()
                nw477_.ctor__(default__.safeModuloInt(d_2880_v3_, 576))
                d_2945_v53_ = nw477_
                d_2946_v54_: _dafny.Seq
                d_2946_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmclbyo"))
                index537_ = default__.safeIndex(187, (d_2877_v1_).length(0))
                rhs524_ = (((d_2901_v22_)[not(p1)] if (not(p1)) in (d_2901_v22_) else (self).f27)) <= ((d_2945_v53_).f34)
                rhs525_ = p3
                rhs526_ = p0
                rhs527_ = d_2946_v54_
                rhs528_ = (0) - ((self).f27)
                lhs393_ = globalState
                lhs394_ = d_2877_v1_
                lhs395_ = default__.safeIndex(187, (d_2877_v1_).length(0))
                lhs393_.f12 = rhs524_
                lhs394_[lhs395_] = rhs525_
                d_2902_v23_ = rhs526_
                r2 = rhs527_
                d_2880_v3_ = rhs528_
                d_2880_v3_ = (d_2945_v53_).f34
                d_2880_v3_ = default__.safeModuloInt((0) - ((0) - (d_2880_v3_)), (self).f25)
                d_2884_v7_ = ((d_2884_v7_).set(default__.safeIndex((self).f27, len(d_2884_v7_)), 155)) + (d_2884_v7_)
            elif True:
                d_2947_v55_: _dafny.Array
                d_2948_v56_: _dafny.Array
                d_2949_v57_: bool
                out49_: _dafny.Array
                out50_: _dafny.Array
                out51_: bool
                out49_, out50_, out51_ = default__.m0(globalState)
                d_2947_v55_ = out49_
                d_2948_v56_ = out50_
                d_2949_v57_ = out51_
                d_2950_v58_: _dafny.Map
                d_2950_v58_ = _dafny.Map({self.f28: d_2901_v22_})
                d_2950_v58_ = d_2950_v58_
                d_2951_v59_: _dafny.Set
                d_2951_v59_ = _dafny.Set({d_2902_v23_, d_2948_v56_, d_2902_v23_, p0, d_2902_v23_})
                d_2952_v60_: D1
                d_2952_v60_ = D1_DC4(d_2888_v11_)
                d_2953_v61_: _dafny.Seq
                d_2953_v61_ = _dafny.SeqWithoutIsStrInference([D1_DC7(d_2952_v60_)])
                rhs529_ = ((self).f25) + (len(d_2953_v61_))
                rhs530_ = d_2951_v59_
                d_2880_v3_ = rhs529_
                d_2951_v59_ = rhs530_
                d_2954_v62_: _dafny.Seq
                d_2954_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjomdhuxy"))
                r2 = (default__.fm2((0) - (d_2880_v3_), (self).f27, globalState)) + ((d_2954_v62_) + (d_2954_v62_))
                d_2955_v63_: _dafny.Map
                d_2955_v63_ = _dafny.Map({self.f28: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2956_i8_ in range(default__.abs(229))])})
                d_2955_v63_ = (d_2955_v63_).set(self.f28, d_2954_v62_)
        _ingredientsd_0 = []
        guard_loop_14_: int
        for guard_loop_14_ in _dafny.IntegerRange(0, (d_2877_v1_).length(0)):
            d_2957_i9_: int = guard_loop_14_
            if (True) and (((0) <= (d_2957_i9_)) and ((d_2957_i9_) < ((d_2877_v1_).length(0)))):
                _ingredientsd_0.append((d_2877_v1_, int(d_2957_i9_), ((d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))] if not(self.f28) else True)))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        r1 = (len(_dafny.SeqWithoutIsStrInference([(d_2884_v7_)[default__.safeIndex((self).f27, len(d_2884_v7_))], (self).f27]))) <= ((self).f27)
        index538_ = default__.safeIndex(187, (d_2877_v1_).length(0))
        (d_2877_v1_)[index538_] = (d_2877_v1_)[default__.safeIndex(187, (d_2877_v1_).length(0))]
        r0 = (not (False) or (p1)) == (self.f28)
        r1 = ((self).f27) < ((self).f25)
        d_2958_v64_: _dafny.Seq
        d_2958_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bowcxhc"))
        r2 = d_2958_v64_
        return r0, r1, r2

    @property
    def f27(self):
        return self._f27

class C13:
    def  __init__(self):
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    def ctor__(self, f24):
        (self)._f24 = f24

    def fm3(self, p0, p1, p2, p3, globalState):
        return (self).f24

    def m1(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: _dafny.Seq = _dafny.Seq({})
        d_2959_v0_: _dafny.Seq
        d_2959_v0_ = _dafny.SeqWithoutIsStrInference([False])
        d_2960_v1_: D1
        d_2960_v1_ = D1_DC4(d_2959_v0_)
        d_2961_v2_: D1
        d_2961_v2_ = D1_DC7(d_2960_v1_)
        d_2962_v3_: _dafny.Set
        d_2962_v3_ = _dafny.Set({d_2961_v2_, d_2961_v2_, d_2961_v2_, D1_DC7(D1_DC7(d_2960_v1_))})
        d_2963_v4_: _dafny.Seq
        d_2963_v4_ = _dafny.SeqWithoutIsStrInference([d_2962_v3_])
        d_2964_v5_: _dafny.Seq
        d_2964_v5_ = _dafny.SeqWithoutIsStrInference([D1_DC4(d_2959_v0_), D1_DC5()])
        d_2965_v6_: _dafny.MultiSet
        d_2965_v6_ = _dafny.MultiSet([d_2961_v2_, d_2961_v2_, D1_DC7((d_2964_v5_)[default__.safeIndex(p0, len(d_2964_v5_))]), d_2961_v2_])
        def iife252_():
            coll96_ = _dafny.Set()
            compr_96_: D1
            for compr_96_ in (d_2965_v6_).Elements:
                d_2966_v7_: D1 = compr_96_
                if (d_2966_v7_) in (d_2965_v6_):
                    coll96_ = coll96_.union(_dafny.Set([d_2966_v7_]))
            return _dafny.Set(coll96_)
        r3 = _dafny.SeqWithoutIsStrInference([(iife252_()
        ).issubset((d_2963_v4_)[default__.safeIndex(p0, len(d_2963_v4_))]), False])
        d_2967_v8_: str
        d_2967_v8_ = _dafny.CodePoint('c')
        (globalState).f23 = d_2967_v8_
        d_2968_v9_: _dafny.MultiSet
        d_2968_v9_ = _dafny.MultiSet([(self).f24, (self).f24])
        hi15_ = (0) - (((default__.fm4((self).f24, globalState)) - (d_2968_v9_)).cardinality)
        for d_2969_i0_ in range(p0, hi15_):
            if ((self).f24) == (p1):
                d_2970_v10_: D0
                d_2970_v10_ = D0_DC2(p1)
                d_2971_v11_: _dafny.Map
                d_2971_v11_ = _dafny.Map({88: d_2970_v10_})
                d_2971_v11_ = (d_2971_v11_).set(len(_dafny.Map({p0: p0})), d_2970_v10_)
                d_2972_v12_: int
                d_2972_v12_ = 701
                d_2972_v12_ = (d_2969_i0_) - ((d_2972_v12_) + (d_2969_i0_))
                d_2973_v13_: D0
                d_2973_v13_ = D0_DC3(d_2972_v12_, (self).f24)
                d_2974_v14_: _dafny.Map
                d_2974_v14_ = _dafny.Map({(d_2973_v13_).cf4: ((self).f24) not in (_dafny.Set({(self).f24, (self).f24, (self).f24}))})
                (globalState).f12 = ((d_2974_v14_)[(0) - (p0)] if ((0) - (p0)) in (d_2974_v14_) else ((self).f24) or ((self).f24))
                d_2975_v15_: _dafny.Array
                nw478_ = _dafny.Array(int(0), 12)
                d_2975_v15_ = nw478_
                index539_ = default__.safeIndex(6, (d_2975_v15_).length(0))
                (d_2975_v15_)[index539_] = p0
                d_2976_v16_: _dafny.Array
                nw479_ = _dafny.Array(False, 27)
                d_2976_v16_ = nw479_
                index540_ = default__.safeIndex(589, (d_2976_v16_).length(0))
                (d_2976_v16_)[index540_] = p1
                index541_ = default__.safeIndex(224, (d_2976_v16_).length(0))
                (d_2976_v16_)[index541_] = p1
                d_2977_v17_: _dafny.Map
                d_2977_v17_ = _dafny.Map({d_2975_v15_: len(d_2974_v14_)})
                d_2978_v18_: _dafny.Seq
                d_2978_v18_ = _dafny.SeqWithoutIsStrInference([484])
                d_2979_v19_: _dafny.Map
                d_2979_v19_ = _dafny.Map({(self).f24: (d_2959_v0_)[default__.safeIndex((d_2973_v13_).cf4, len(d_2959_v0_))]})
                index542_ = default__.safeIndex(6, (d_2975_v15_).length(0))
                index543_ = default__.safeIndex(589, (d_2976_v16_).length(0))
                index544_ = default__.safeIndex(224, (d_2976_v16_).length(0))
                rhs531_ = (len(d_2977_v17_)) * (len(d_2978_v18_))
                rhs532_ = d_2967_v8_
                rhs533_ = (self).f24
                rhs534_ = not (((d_2979_v19_)[p1] if (p1) in (d_2979_v19_) else p1)) or (p1)
                lhs396_ = d_2975_v15_
                lhs397_ = default__.safeIndex(6, (d_2975_v15_).length(0))
                lhs398_ = globalState
                lhs399_ = d_2976_v16_
                lhs400_ = default__.safeIndex(589, (d_2976_v16_).length(0))
                lhs401_ = d_2976_v16_
                lhs402_ = default__.safeIndex(224, (d_2976_v16_).length(0))
                lhs396_[lhs397_] = rhs531_
                lhs398_.f23 = rhs532_
                lhs399_[lhs400_] = rhs533_
                lhs401_[lhs402_] = rhs534_
                index545_ = default__.safeIndex(6, (d_2975_v15_).length(0))
                rhs535_ = d_2974_v14_
                rhs536_ = p0
                rhs537_ = False
                rhs538_ = default__.fm5((self).f24, globalState)
                lhs403_ = d_2975_v15_
                lhs404_ = default__.safeIndex(6, (d_2975_v15_).length(0))
                lhs405_ = globalState
                d_2974_v14_ = rhs535_
                lhs403_[lhs404_] = rhs536_
                lhs405_.f12 = rhs537_
                d_2972_v12_ = rhs538_
            elif True:
                d_2980_v20_: int
                d_2980_v20_ = 472
                d_2981_v21_: _dafny.Seq
                d_2981_v21_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (p0)])
                d_2982_v22_: _dafny.MultiSet
                d_2982_v22_ = _dafny.MultiSet([p0, d_2969_i0_, default__.safeModuloInt(d_2980_v20_, p0), default__.fm5(not((self).f24), globalState), default__.safeModuloInt((d_2981_v21_)[default__.safeIndex(p0, len(d_2981_v21_))], p0)])
                d_2983_v23_: D0
                d_2983_v23_ = D0_DC2((self).f24)
                d_2984_v24_: D0
                d_2984_v24_ = D0_DC1(p1, p1)
                d_2985_v25_: _dafny.Seq
                d_2985_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdwocu"))
                d_2986_v26_: _dafny.Map
                d_2986_v26_ = _dafny.Map({True: p0})
                d_2987_v27_: _dafny.Set
                d_2987_v27_ = _dafny.Set({d_2969_i0_, p0})
                d_2988_v28_: _dafny.Array
                nw480_ = _dafny.Array(None, 16)
                nw480_[int(0)] = (d_2983_v23_).cf3
                nw480_[int(1)] = (self).f24
                nw480_[int(2)] = False
                nw480_[int(3)] = (p1) or ((self).f24)
                nw480_[int(4)] = not(p1)
                nw480_[int(5)] = (self).f24
                nw480_[int(6)] = ((self).fm3(p0, d_2981_v21_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1, (d_2984_v24_).cf2]), d_2959_v0_])), p1, globalState)) and ((self).f24)
                nw480_[int(7)] = p1
                nw480_[int(8)] = (self).f24
                nw480_[int(9)] = p1
                nw480_[int(10)] = (p1) and (not(p1))
                nw480_[int(11)] = (d_2980_v20_) == (d_2980_v20_)
                nw480_[int(12)] = p1
                nw480_[int(13)] = (d_2987_v27_).ispropersubset(_dafny.Set({((d_2968_v9_)[(self).f24] if ((self).f24) in (d_2968_v9_) else p0), len(d_2985_v25_), len(d_2986_v26_), d_2969_i0_}))
                nw480_[int(14)] = p1
                nw480_[int(15)] = p1
                d_2988_v28_ = nw480_
                index546_ = default__.safeIndex(361, (d_2988_v28_).length(0))
                (d_2988_v28_)[index546_] = p1
                d_2989_v29_: _dafny.Map
                d_2989_v29_ = _dafny.Map({p1: d_2988_v28_})
                d_2990_v30_: _dafny.Map
                d_2990_v30_ = _dafny.Map({len((d_2989_v29_).set(p1, d_2988_v28_)): d_2980_v20_})
                d_2991_v31_: _dafny.MultiSet
                d_2991_v31_ = _dafny.MultiSet([d_2990_v30_])
                d_2992_v32_: _dafny.Seq
                d_2992_v32_ = d_2981_v21_
                index547_ = default__.safeIndex(361, (d_2988_v28_).length(0))
                rhs539_ = (self).f24
                rhs540_ = p0
                rhs541_ = ((d_2982_v22_) | (d_2982_v22_)).intersection(_dafny.MultiSet([d_2969_i0_, d_2980_v20_, (d_2991_v31_).cardinality, d_2980_v20_, len((d_2992_v32_))]))
                rhs542_ = not((d_2959_v0_)[default__.safeIndex(d_2969_i0_, len(d_2959_v0_))])
                rhs543_ = default__.fm5(p1, globalState)
                lhs406_ = globalState
                lhs407_ = d_2988_v28_
                lhs408_ = default__.safeIndex(361, (d_2988_v28_).length(0))
                lhs406_.f12 = rhs539_
                d_2980_v20_ = rhs540_
                d_2982_v22_ = rhs541_
                lhs407_[lhs408_] = rhs542_
                d_2980_v20_ = rhs543_
                d_2993_v33_: _dafny.Array
                nw481_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_2993_v33_ = nw481_
                d_2994_v34_: _dafny.Map
                d_2994_v34_ = _dafny.Map({920: (self).f24})
                d_2995_v35_: _dafny.Map
                d_2995_v35_ = _dafny.Map({d_2993_v33_: d_2994_v34_})
                d_2995_v35_ = (d_2995_v35_).set(d_2993_v33_, d_2994_v34_)
                d_2996_v36_: _dafny.MultiSet
                d_2996_v36_ = _dafny.MultiSet([d_2983_v23_])
                d_2996_v36_ = (d_2996_v36_) | (d_2996_v36_)
                d_2997_v37_: D0
                d_2997_v37_ = D0_DC3(p0, not((self).f24))
                (globalState).f12 = ((d_2969_i0_) * ((d_2997_v37_).cf4)) < (len(((_dafny.SeqWithoutIsStrInference([d_2969_i0_])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_2969_i0_]))), d_2969_i0_)) + (d_2981_v21_)))
                r2 = (d_2988_v28_)[default__.safeIndex(361, (d_2988_v28_).length(0))]
            d_2998_v38_: _dafny.Array
            def lambda182_(d_2999_i0_):
                def lambda183_(d_3000_i1_):
                    return default__.safeModuloInt(d_3000_i1_, d_2999_i0_)

                return lambda183_

            init100_ = lambda182_(d_2969_i0_)
            nw482_ = _dafny.Array(None, 4)
            for i0_100_ in range(nw482_.length(0)):
                nw482_[i0_100_] = init100_(i0_100_)
            d_2998_v38_ = nw482_
            index548_ = default__.safeIndex(833, (d_2998_v38_).length(0))
            (d_2998_v38_)[index548_] = p0
            index549_ = default__.safeIndex(833, (d_2998_v38_).length(0))
            (d_2998_v38_)[index549_] = (default__.safeModuloInt(p0, d_2969_i0_)) - ((len(d_2959_v0_)) * (d_2969_i0_))
            d_3001_v39_: _dafny.Map
            d_3001_v39_ = _dafny.Map({len(d_2959_v0_): (self).f24})
            d_3002_v40_: _dafny.Seq
            d_3002_v40_ = _dafny.SeqWithoutIsStrInference([d_3001_v39_, d_3001_v39_])
            d_3003_v41_: _dafny.Map
            d_3003_v41_ = _dafny.Map({((d_2998_v38_)[default__.safeIndex(833, (d_2998_v38_).length(0))]) < (d_2969_i0_): _dafny.MultiSet(((d_2959_v0_).set(default__.safeIndex((d_2998_v38_)[default__.safeIndex(833, (d_2998_v38_).length(0))], len(d_2959_v0_)), (self).f24)).set(default__.safeIndex(len((d_3002_v40_)[default__.safeIndex(d_2969_i0_, len(d_3002_v40_))]), len((d_2959_v0_).set(default__.safeIndex((d_2998_v38_)[default__.safeIndex(833, (d_2998_v38_).length(0))], len(d_2959_v0_)), (self).f24))), (self).f24))})
            d_3003_v41_ = (d_3003_v41_).set((self).f24, d_2968_v9_)
            d_3004_v42_: _dafny.Seq
            d_3004_v42_ = _dafny.SeqWithoutIsStrInference([d_2969_i0_, (d_2968_v9_).cardinality])
            source40_ = d_3004_v42_
            d_3005___mcc_h0_ = source40_
            d_3006_cf11_ = d_3005___mcc_h0_
            d_3007_v43_: _dafny.Array
            nw483_ = _dafny.Array(False, 27)
            d_3007_v43_ = nw483_
            d_3008_v44_: _dafny.MultiSet
            d_3008_v44_ = _dafny.MultiSet([p0])
            index550_ = default__.safeIndex(739, (d_3007_v43_).length(0))
            (d_3007_v43_)[index550_] = not(not(((d_3001_v39_)[((d_3008_v44_)[p0] if (p0) in (d_3008_v44_) else (d_2998_v38_)[default__.safeIndex(833, (d_2998_v38_).length(0))])] if (((d_3008_v44_)[p0] if (p0) in (d_3008_v44_) else (d_2998_v38_)[default__.safeIndex(833, (d_2998_v38_).length(0))])) in (d_3001_v39_) else False)))
            index551_ = default__.safeIndex(739, (d_3007_v43_).length(0))
            (d_3007_v43_)[index551_] = p1
            index552_ = default__.safeIndex(739, (d_3007_v43_).length(0))
            (d_3007_v43_)[index552_] = ((self).f24) and ((self).f24)
            d_3009_v45_: _dafny.Map
            d_3009_v45_ = _dafny.Map({(d_3007_v43_)[default__.safeIndex(739, (d_3007_v43_).length(0))]: p1})
            d_3010_v46_: _dafny.Map
            d_3010_v46_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2969_i0_ for d_3011_i2_ in range(default__.abs(288))])): len(d_3009_v45_)})
            rhs544_ = d_3010_v46_
            rhs545_ = (((d_3009_v45_).set((d_3007_v43_)[default__.safeIndex(739, (d_3007_v43_).length(0))], p1)) | (d_3009_v45_)) | (d_3009_v45_)
            rhs546_ = d_3007_v43_
            d_3010_v46_ = rhs544_
            d_3009_v45_ = rhs545_
            d_3007_v43_ = rhs546_
            d_3012_v47_: D1
            d_3012_v47_ = D1_DC4(d_2959_v0_)
            d_3012_v47_ = d_3012_v47_
        r0 = (p0) != ((0) - (p0))
        d_3013_v48_: D0
        d_3013_v48_ = D0_DC0(p1)
        pat_let_tv78_ = d_2967_v8_
        pat_let_tv79_ = d_2967_v8_
        pat_let_tv80_ = d_2967_v8_
        pat_let_tv81_ = d_2967_v8_
        def lambda184_(source41_):
            if source41_.is_DC1:
                d_3014___mcc_h1_ = source41_.cf1
                d_3015___mcc_h2_ = source41_.cf2
                d_3016_cf2_ = d_3015___mcc_h2_
                d_3017_cf1_ = d_3014___mcc_h1_
                return pat_let_tv78_
            elif source41_.is_DC2:
                d_3018___mcc_h3_ = source41_.cf3
                d_3019_cf3_ = d_3018___mcc_h3_
                return pat_let_tv79_
            elif source41_.is_DC3:
                d_3020___mcc_h4_ = source41_.cf4
                d_3021___mcc_h5_ = source41_.cf5
                d_3022_cf5_ = d_3021___mcc_h5_
                d_3023_cf4_ = d_3020___mcc_h4_
                return pat_let_tv80_
            elif True:
                d_3024___mcc_h6_ = source41_.cf0
                d_3025_cf0_ = d_3024___mcc_h6_
                return pat_let_tv81_

        (globalState).f23 = lambda184_(d_3013_v48_)
        d_3026_v49_: D0
        d_3026_v49_ = D0_DC3((p0) * (p0), False)
        source42_ = d_3026_v49_
        if source42_.is_DC1:
            d_3027___mcc_h7_ = source42_.cf1
            d_3028___mcc_h8_ = source42_.cf2
            d_3029_cf2_ = d_3028___mcc_h8_
            d_3030_cf1_ = d_3027___mcc_h7_
            d_3031_v50_: _dafny.Array
            def lambda185_(d_3032_i3_):
                return (self).f24

            init101_ = lambda185_
            nw484_ = _dafny.Array(None, 5)
            for i0_101_ in range(nw484_.length(0)):
                nw484_[i0_101_] = init101_(i0_101_)
            d_3031_v50_ = nw484_
            d_3033_v51_: _dafny.Array
            def lambda186_(d_3034_p0_):
                def lambda187_(d_3035_i4_):
                    return default__.safeModuloInt(d_3035_i4_, d_3034_p0_)

                return lambda187_

            init102_ = lambda186_(p0)
            nw485_ = _dafny.Array(None, 8)
            for i0_102_ in range(nw485_.length(0)):
                nw485_[i0_102_] = init102_(i0_102_)
            d_3033_v51_ = nw485_
            index553_ = default__.safeIndex(202, (d_3033_v51_).length(0))
            (d_3033_v51_)[index553_] = p0
            index554_ = default__.safeIndex(202, (d_3033_v51_).length(0))
            rhs547_ = d_3030_cf1_
            rhs548_ = d_3031_v50_
            rhs549_ = 285
            lhs409_ = d_3033_v51_
            lhs410_ = default__.safeIndex(202, (d_3033_v51_).length(0))
            d_3029_cf2_ = rhs547_
            d_3031_v50_ = rhs548_
            lhs409_[lhs410_] = rhs549_
            d_3036_v52_: _dafny.Seq
            d_3036_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugn"))
            d_3037_v53_: _dafny.Map
            d_3037_v53_ = _dafny.Map({p0: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfodmffx")) if d_3029_cf2_ else d_3036_v52_)})
            d_3037_v53_ = (d_3037_v53_).set(463, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqq")))
            d_3038_v54_: _dafny.Set
            d_3038_v54_ = _dafny.Set({(p0) - (p0), (d_3033_v51_)[default__.safeIndex(202, (d_3033_v51_).length(0))]})
            d_3038_v54_ = d_3038_v54_
            d_3039_v55_: _dafny.Map
            d_3039_v55_ = _dafny.Map({(d_3033_v51_)[default__.safeIndex(202, (d_3033_v51_).length(0))]: (0) - ((p0 if d_3029_cf2_ else p0))})
            d_3039_v55_ = d_3039_v55_
        elif source42_.is_DC2:
            d_3040___mcc_h9_ = source42_.cf3
            d_3041_cf3_ = d_3040___mcc_h9_
            if (self).f24:
                d_3042_v56_: _dafny.Array
                d_3043_v57_: _dafny.Array
                d_3044_v58_: bool
                out52_: _dafny.Array
                out53_: _dafny.Array
                out54_: bool
                out52_, out53_, out54_ = default__.m0(globalState)
                d_3042_v56_ = out52_
                d_3043_v57_ = out53_
                d_3044_v58_ = out54_
                d_3045_v59_: int
                d_3045_v59_ = 337
                d_3045_v59_ = d_3045_v59_
                d_3046_v60_: _dafny.Set
                d_3046_v60_ = _dafny.Set({d_3042_v56_, d_3042_v56_, d_3042_v56_, d_3042_v56_})
                d_3047_v61_: D3
                d_3047_v61_ = D3_DC9(_dafny.Set({d_3042_v56_, d_3042_v56_}))
                d_3048_v62_: _dafny.Array
                nw486_ = _dafny.Array(None, 12)
                nw486_[int(0)] = d_3046_v60_
                nw486_[int(1)] = d_3046_v60_
                nw486_[int(2)] = (d_3046_v60_) - (_dafny.Set({d_3042_v56_, d_3042_v56_}))
                nw486_[int(3)] = (d_3046_v60_) - (d_3046_v60_)
                nw486_[int(4)] = d_3046_v60_
                nw486_[int(5)] = (d_3046_v60_) | (d_3046_v60_)
                nw486_[int(6)] = d_3046_v60_
                nw486_[int(7)] = _dafny.Set({d_3042_v56_, d_3042_v56_, d_3042_v56_, d_3042_v56_})
                nw486_[int(8)] = (d_3047_v61_).cf12
                nw486_[int(9)] = (d_3046_v60_) | (d_3046_v60_)
                nw486_[int(10)] = (d_3046_v60_) | (d_3046_v60_)
                nw486_[int(11)] = d_3046_v60_
                d_3048_v62_ = nw486_
                index555_ = default__.safeIndex(890, (d_3048_v62_).length(0))
                (d_3048_v62_)[index555_] = d_3046_v60_
                index556_ = default__.safeIndex(890, (d_3048_v62_).length(0))
                (d_3048_v62_)[index556_] = d_3046_v60_
                d_3049_v63_: _dafny.Set
                d_3049_v63_ = _dafny.Set({p0})
                index557_ = default__.safeIndex(272, (d_3042_v56_).length(0))
                (d_3042_v56_)[index557_] = (d_3049_v63_).isdisjoint(_dafny.Set({p0, p0}))
                d_3050_v64_: _dafny.Set
                d_3050_v64_ = _dafny.Set({False, d_3041_cf3_, p1, not(d_3044_v58_), (self).f24})
                index558_ = default__.safeIndex(272, (d_3042_v56_).length(0))
                (d_3042_v56_)[index558_] = ((d_3050_v64_).intersection(_dafny.Set({d_3041_cf3_}))).issubset(d_3050_v64_)
                d_3051_v65_: _dafny.Seq
                d_3051_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teauuxeu"))
                d_3045_v59_ = len(d_3051_v65_)
            elif True:
                d_3052_v66_: int
                d_3052_v66_ = 135
                d_3052_v66_ = p0
                d_3053_v67_: _dafny.Seq
                d_3053_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fo"))
                d_3054_v68_: _dafny.Seq
                d_3054_v68_ = _dafny.SeqWithoutIsStrInference([d_3053_v67_])
                d_3055_v69_: _dafny.Map
                d_3055_v69_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_3053_v67_])})
                d_3056_v70_: _dafny.Array
                nw487_ = _dafny.Array(None, 11)
                nw487_[int(0)] = d_3054_v68_
                nw487_[int(1)] = d_3054_v68_
                nw487_[int(2)] = d_3054_v68_
                nw487_[int(3)] = d_3054_v68_
                nw487_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_3053_v67_])) + (((d_3055_v69_)[690] if (690) in (d_3055_v69_) else d_3054_v68_))
                nw487_[int(5)] = d_3054_v68_
                nw487_[int(6)] = default__.fm6(globalState)
                nw487_[int(7)] = d_3054_v68_
                nw487_[int(8)] = d_3054_v68_
                nw487_[int(9)] = d_3054_v68_
                nw487_[int(10)] = (d_3054_v68_) + (_dafny.SeqWithoutIsStrInference([d_3053_v67_]))
                d_3056_v70_ = nw487_
                index559_ = default__.safeIndex(371, (d_3056_v70_).length(0))
                (d_3056_v70_)[index559_] = d_3054_v68_
                d_3057_v71_: _dafny.Array
                nw488_ = _dafny.Array(False, 27)
                d_3057_v71_ = nw488_
                d_3058_v72_: _dafny.Seq
                d_3058_v72_ = _dafny.SeqWithoutIsStrInference([-52])
                index560_ = default__.safeIndex(246, (d_3057_v71_).length(0))
                (d_3057_v71_)[index560_] = (d_3058_v72_) <= ((d_3058_v72_).set(default__.safeIndex(len((d_3054_v68_)[default__.safeIndex((_dafny.MultiSet(d_2959_v0_)).cardinality, len(d_3054_v68_))]), len(d_3058_v72_)), p0))
                index561_ = default__.safeIndex(371, (d_3056_v70_).length(0))
                index562_ = default__.safeIndex(246, (d_3057_v71_).length(0))
                rhs550_ = d_3054_v68_
                rhs551_ = p1
                lhs411_ = d_3056_v70_
                lhs412_ = default__.safeIndex(371, (d_3056_v70_).length(0))
                lhs413_ = d_3057_v71_
                lhs414_ = default__.safeIndex(246, (d_3057_v71_).length(0))
                lhs411_[lhs412_] = rhs550_
                lhs413_[lhs414_] = rhs551_
                d_3059_v73_: _dafny.Array
                d_3060_v74_: _dafny.Array
                d_3061_v75_: bool
                out55_: _dafny.Array
                out56_: _dafny.Array
                out57_: bool
                out55_, out56_, out57_ = default__.m0(globalState)
                d_3059_v73_ = out55_
                d_3060_v74_ = out56_
                d_3061_v75_ = out57_
                d_3062_v76_: _dafny.MultiSet
                d_3062_v76_ = _dafny.MultiSet([d_3052_v66_])
                (globalState).f12 = (d_2959_v0_)[default__.safeIndex(((d_3062_v76_).intersection(d_3062_v76_)).cardinality, len(d_2959_v0_))]
                d_3062_v76_ = (d_3062_v76_) - (d_3062_v76_)
            d_3063_v77_: C0
            nw489_ = C0()
            nw489_.ctor__()
            d_3063_v77_ = nw489_
            d_3064_v78_: int
            d_3064_v78_ = 513
            rhs552_ = p0
            rhs553_ = ((d_3064_v78_ if p1 else len(_dafny.SeqWithoutIsStrInference([d_3064_v78_ for d_3065_i5_ in range(default__.abs(-272))])))) - (d_3064_v78_)
            d_3064_v78_ = rhs552_
            d_3064_v78_ = rhs553_
            d_3066_v79_: _dafny.Array
            nw490_ = _dafny.Array(int(0), 11)
            d_3066_v79_ = nw490_
            index563_ = default__.safeIndex(258, (d_3066_v79_).length(0))
            (d_3066_v79_)[index563_] = len(d_2959_v0_)
            index564_ = default__.safeIndex(258, (d_3066_v79_).length(0))
            (d_3066_v79_)[index564_] = d_3064_v78_
        elif source42_.is_DC3:
            d_3067___mcc_h10_ = source42_.cf4
            d_3068___mcc_h11_ = source42_.cf5
            d_3069_cf5_ = d_3068___mcc_h11_
            d_3070_cf4_ = d_3067___mcc_h10_
            rhs554_ = (self).f24
            rhs555_ = (self).f24
            r0 = rhs554_
            r2 = rhs555_
            d_2968_v9_ = d_2968_v9_
            d_3071_v80_: _dafny.MultiSet
            d_3071_v80_ = _dafny.MultiSet([p0])
            d_3072_v81_: _dafny.Array
            nw491_ = _dafny.Array(None, 3)
            nw491_[int(0)] = d_2967_v8_
            nw491_[int(1)] = d_2967_v8_
            nw491_[int(2)] = default__.fm33(13, d_3069_cf5_, (d_3071_v80_).cardinality, globalState)
            d_3072_v81_ = nw491_
            index565_ = default__.safeIndex(466, (d_3072_v81_).length(0))
            (d_3072_v81_)[index565_] = d_2967_v8_
            index566_ = default__.safeIndex(466, (d_3072_v81_).length(0))
            (d_3072_v81_)[index566_] = default__.fm39((self).f24, globalState)
            d_3073_v82_: D1
            d_3073_v82_ = D1_DC5()
            d_3074_v83_: D8
            d_3074_v83_ = D8_DC24(d_3073_v82_)
            source43_ = d_3074_v83_
            if source43_.is_DC24:
                d_3075___mcc_h13_ = source43_.cf36
                d_3076_cf36_ = d_3075___mcc_h13_
                d_3077_v84_: _dafny.Array
                nw492_ = _dafny.Array(int(0), 26)
                d_3077_v84_ = nw492_
                index567_ = default__.safeIndex(666, (d_3077_v84_).length(0))
                (d_3077_v84_)[index567_] = default__.safeModuloInt(d_3070_cf4_, d_3070_cf4_)
                index568_ = default__.safeIndex(666, (d_3077_v84_).length(0))
                (d_3077_v84_)[index568_] = (p0) * ((0) - (p0))
                d_3078_v85_: _dafny.Set
                d_3078_v85_ = _dafny.Set({(self).f24})
                d_3079_v86_: _dafny.Seq
                d_3079_v86_ = _dafny.SeqWithoutIsStrInference([d_3078_v85_])
                d_3080_v87_: _dafny.Seq
                d_3080_v87_ = _dafny.SeqWithoutIsStrInference([(d_3077_v84_)[default__.safeIndex(666, (d_3077_v84_).length(0))]])
                d_3081_v88_: D9
                d_3081_v88_ = D9_DC27(len((d_3079_v86_)[default__.safeIndex(d_3070_cf4_, len(d_3079_v86_))]), ((d_3080_v87_)[default__.safeIndex(-873, len(d_3080_v87_))]) - (d_3070_cf4_))
                d_3081_v88_ = D9_DC27(p0, p0)
                d_3082_v89_: C5
                nw493_ = C5()
                nw493_.ctor__((p0) < (p0))
                d_3082_v89_ = nw493_
                d_3070_cf4_ = default__.safeDivisionInt(p0, (0) - (d_3070_cf4_))
            elif source43_.is_DC23:
                d_3083___mcc_h14_ = source43_.cf35
                d_3084_cf35_ = d_3083___mcc_h14_
                d_2968_v9_ = (_dafny.MultiSet([True, p1]) if (self).f24 else d_2968_v9_)
                d_3085_v90_: _dafny.Array
                d_3086_v91_: _dafny.Array
                d_3087_v92_: bool
                out58_: _dafny.Array
                out59_: _dafny.Array
                out60_: bool
                out58_, out59_, out60_ = default__.m0(globalState)
                d_3085_v90_ = out58_
                d_3086_v91_ = out59_
                d_3087_v92_ = out60_
                d_3088_v93_: D7
                d_3088_v93_ = D7_DC21(d_3086_v91_)
                d_3089_v94_: C9
                nw494_ = C9()
                nw494_.ctor__((d_3088_v93_).cf32, (-268) - (401))
                d_3089_v94_ = nw494_
                d_3090_v95_: _dafny.Seq
                d_3090_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grvwvats"))
                d_3091_v96_: _dafny.Map
                d_3091_v96_ = _dafny.Map({False: len(d_3090_v95_)})
                d_3092_v97_: _dafny.Set
                d_3092_v97_ = _dafny.Set({len(d_3091_v96_), p0})
                d_3093_v98_: _dafny.Map
                d_3093_v98_ = _dafny.Map({(d_3092_v97_).ispropersubset(d_3092_v97_): (d_2959_v0_)[default__.safeIndex(len(default__.fm2((0) - (p0), p0, globalState)), len(d_2959_v0_))]})
                d_3094_v99_: _dafny.Map
                d_3094_v99_ = _dafny.Map({(d_3072_v81_)[default__.safeIndex(466, (d_3072_v81_).length(0))]: (d_3089_v94_).f31})
                d_3093_v98_ = (d_3093_v98_).set((((d_3094_v99_)[(d_3072_v81_)[default__.safeIndex(466, (d_3072_v81_).length(0))]] if ((d_3072_v81_)[default__.safeIndex(466, (d_3072_v81_).length(0))]) in (d_3094_v99_) else len(d_3092_v97_))) != (default__.fm31(d_3090_v95_, globalState)), True)
            elif True:
                d_3095___mcc_h15_ = source43_.cf37
                d_3096_cf37_ = d_3095___mcc_h15_
                d_3097_v100_: _dafny.Map
                d_3097_v100_ = _dafny.Map({d_3070_cf4_: d_3070_cf4_})
                d_3098_v101_: _dafny.Array
                nw495_ = _dafny.Array(False, 7)
                d_3098_v101_ = nw495_
                d_3099_v102_: _dafny.Map
                d_3099_v102_ = _dafny.Map({d_3098_v101_: not(p1)})
                d_3100_v103_: _dafny.Seq
                d_3100_v103_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(d_3069_cf5_): len(_dafny.SeqWithoutIsStrInference([d_2967_v8_ for d_3101_i6_ in range(default__.abs(140))]))})), p0])
                d_3097_v100_ = (d_3097_v100_).set(d_3070_cf4_, (len(d_3099_v102_)) * (len(d_3100_v103_)))
                d_3102_v104_: _dafny.Set
                d_3102_v104_ = _dafny.Set({d_3013_v48_, d_3013_v48_})
                rhs556_ = d_3102_v104_
                rhs557_ = d_3070_cf4_
                d_3102_v104_ = rhs556_
                d_3070_cf4_ = rhs557_
                d_3103_v105_: _dafny.Map
                d_3103_v105_ = _dafny.Map({True: d_3070_cf4_})
                d_3070_cf4_ = (p0) - ((0) - (len(d_3103_v105_)))
                d_3103_v105_ = (d_3103_v105_).set(not((p0) != (d_3070_cf4_)), len(d_2959_v0_))
        elif True:
            d_3104___mcc_h12_ = source42_.cf0
            d_3105_cf0_ = d_3104___mcc_h12_
            d_3105_cf0_ = d_3105_cf0_
            d_2959_v0_ = (d_2959_v0_) + (d_2959_v0_)
            d_3106_v106_: int
            d_3106_v106_ = 461
            d_3106_v106_ = p0
            d_3107_v107_: _dafny.Set
            d_3107_v107_ = _dafny.Set({d_3105_cf0_})
            d_3108_v108_: _dafny.Map
            d_3108_v108_ = _dafny.Map({True: d_3105_cf0_})
            d_3109_v109_: _dafny.Map
            d_3109_v109_ = _dafny.Map({(self).f24: d_3108_v108_})
            rhs558_ = ((d_3107_v107_) | (d_3107_v107_)) | (d_3107_v107_)
            rhs559_ = p1
            rhs560_ = d_3109_v109_
            rhs561_ = p0
            d_3107_v107_ = rhs558_
            r2 = rhs559_
            d_3109_v109_ = rhs560_
            d_3106_v106_ = rhs561_
        d_3110_v110_: C11
        nw496_ = C11()
        nw496_.ctor__()
        d_3110_v110_ = nw496_
        d_3111_v111_: _dafny.Map
        d_3111_v111_ = _dafny.Map({d_2967_v8_: d_3110_v110_})
        d_3112_v112_: _dafny.Map
        d_3112_v112_ = _dafny.Map({p1: p0})
        r0 = (default__.safeDivisionInt(len(d_3111_v111_), len(d_3112_v112_))) > (default__.safeModuloInt(p0, p0))
        d_3113_v113_: _dafny.Seq
        d_3113_v113_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtjgpuljk"))
        d_3114_v114_: _dafny.Map
        d_3114_v114_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, p0]): d_3113_v113_})
        r1 = d_3114_v114_
        r2 = p1
        d_3115_v115_: D9
        d_3115_v115_ = D9_DC27(p0, p0)
        pat_let_tv82_ = d_2959_v0_
        pat_let_tv83_ = d_2959_v0_
        def lambda188_(source44_):
            if source44_.is_DC27:
                d_3116___mcc_h16_ = source44_.cf39
                d_3117___mcc_h17_ = source44_.cf40
                d_3118_cf40_ = d_3117___mcc_h17_
                d_3119_cf39_ = d_3116___mcc_h16_
                return pat_let_tv82_
            elif True:
                d_3120___mcc_h18_ = source44_.cf38
                d_3121_cf38_ = d_3120___mcc_h18_
                return pat_let_tv83_

        r3 = lambda188_(d_3115_v115_)
        return r0, r1, r2, r3

    @property
    def f24(self):
        return self._f24
