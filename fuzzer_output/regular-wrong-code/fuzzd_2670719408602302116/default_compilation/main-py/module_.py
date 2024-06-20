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
        return _dafny.MultiSet([not (not(not(not(not(True))))) or ((D2_DC6((0) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_0_i0_ in range(default__.abs(797))]))}))), True, not(True))).cf8), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqkb"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocpjiygo"))), (-380) != ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('n')]))))])

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i1_ in range(default__.abs(-386))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xk")))) for d_2_i0_ in range(default__.abs(284))])

    @staticmethod
    def fm5(globalState):
        return D0_DC3((D0_DC1(len(_dafny.Set({False})), 189) if True else D0_DC3(D0_DC2())))

    @staticmethod
    def fm6(p0, globalState):
        return 84

    @staticmethod
    def fm7(p0, globalState):
        return _dafny.Map({(_dafny.CodePoint('y')) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_3_i0_ in range(default__.abs(3))])): _dafny.CodePoint('l')})

    @staticmethod
    def fm11(p0, p1, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([False])) <= (_dafny.SeqWithoutIsStrInference([not(False), True, True, True])): _dafny.CodePoint('k')})

    @staticmethod
    def fm14(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))

    @staticmethod
    def fm17(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jflao"))))

    @staticmethod
    def fm18(p0, p1, globalState):
        return _dafny.CodePoint('e')

    @staticmethod
    def fm20(p0, globalState):
        return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jos"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_4_i0_ in range(default__.abs(500))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilgan")))))

    @staticmethod
    def fm21(p0, globalState):
        return _dafny.Map({False: False})

    @staticmethod
    def fm22(p0, globalState):
        return _dafny.MultiSet([(376) + (len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akm")): not(True)})))])

    @staticmethod
    def fm23(p0, globalState):
        source0_ = _dafny.SeqWithoutIsStrInference([len((D5_DC13((D5_DC13(_dafny.Map({True: False}))).cf19)).cf19), 382])
        d_5___mcc_h0_ = source0_
        d_6_cf24_ = d_5___mcc_h0_
        return (_dafny.SeqWithoutIsStrInference([True])) + ((D9_DC21(False, _dafny.SeqWithoutIsStrInference([True]), True)).cf27)

    @staticmethod
    def fm24(globalState):
        return D3_DC8(_dafny.Map({216: 372}))

    @staticmethod
    def fm25(globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_7_i0_ in range(default__.abs(-715))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unujxwuqx"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_8_i1_ in range(default__.abs(37))])))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return (default__.safeDivisionInt(-978, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_9_i0_ in range(default__.abs(293))]))))) * (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))), 312))

    @staticmethod
    def fm27(p0, globalState):
        source1_ = D19_DC39(False, 106)
        if source1_.is_DC39:
            d_10___mcc_h0_ = source1_.cf58
            d_11___mcc_h1_ = source1_.cf59
            d_12_cf59_ = d_11___mcc_h1_
            d_13_cf58_ = d_10___mcc_h0_
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([506, (_dafny.MultiSet([d_12_cf59_])).cardinality, d_12_cf59_, (_dafny.MultiSet([d_12_cf59_, len(_dafny.Map({d_13_cf58_: d_12_cf59_}))])).cardinality, d_12_cf59_])).cardinality, d_12_cf59_, (0) - (d_12_cf59_)]))
        elif source1_.is_DC38:
            d_14___mcc_h2_ = source1_.cf57
            d_15_cf57_ = d_14___mcc_h2_
            return (_dafny.MultiSet([(_dafny.MultiSet([True, True, True, not(False), True])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kh"))), -878, len(_dafny.Map({True: -273}))])) - (_dafny.MultiSet([(_dafny.MultiSet([939])).cardinality, len(_dafny.SeqWithoutIsStrInference([-120]))]))
        elif True:
            d_16___mcc_h3_ = source1_.cf60
            d_17_cf60_ = d_16___mcc_h3_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([687]))) - (_dafny.MultiSet([(D0_DC0(len(_dafny.Set({(0) - (-72)})))).cf0, 919]))

    @staticmethod
    def fm28(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(2, 844):
                d_18_v0_: int = compr_0_
                if ((2) <= (d_18_v0_)) and ((d_18_v0_) < (844)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_18_v0_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})))]))
            return _dafny.Set(coll0_)
        return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})) | (_dafny.Set({564, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(iife0_()
        ), 277])), len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qthfycdk"))) for d_19_i0_ in range(default__.abs(949))]): not(not(False))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "heoi")))])), 387, (0) - ((_dafny.MultiSet([not(False)])).cardinality)}))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([False, False, False])

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return D4_DC11(not((_dafny.MultiSet([-407])).ispropersubset(_dafny.MultiSet([708, 789, -937]))), (210) != (-759), 842, ((D10_DC24(len(_dafny.Map({True: -675})), True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('y')]))).cf33) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_20_i0_ in range(default__.abs(715))])))

    @staticmethod
    def fm31(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + ((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([True, False])))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        source2_ = D10_DC23(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sisutg"))]))
        if source2_.is_DC24:
            d_21___mcc_h0_ = source2_.cf31
            d_22___mcc_h1_ = source2_.cf32
            d_23___mcc_h2_ = source2_.cf33
            d_24_cf33_ = d_23___mcc_h2_
            d_25_cf32_ = d_22___mcc_h1_
            d_26_cf31_ = d_21___mcc_h0_
            return _dafny.Map({d_25_cf32_: _dafny.Map({d_25_cf32_: d_26_cf31_})})
        elif True:
            d_27___mcc_h3_ = source2_.cf30
            d_28_cf30_ = d_27___mcc_h3_
            return (_dafny.Map({(D2_DC7(_dafny.CodePoint('n'), True, -255)).cf10: _dafny.Map({True: 275})})) | (_dafny.Map({True: _dafny.Map({False: -616})}))

    @staticmethod
    def fm33(p0, p1, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True), not(False)]): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): False}))

    @staticmethod
    def fm34(p0, p1, globalState):
        return D0_DC0(997)

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False, True})) - (_dafny.Set({not(True), not(True), True, False, (D19_DC39(False, 831)).cf58}))) | ((_dafny.Set({False})).intersection(_dafny.Set({not(True), not(False)})))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return D10_DC23(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kueav"))]))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.MultiSet
            for compr_1_ in (_dafny.MultiSet([_dafny.MultiSet([False, False]), _dafny.MultiSet([False])])).Elements:
                d_29_v0_: _dafny.MultiSet = compr_1_
                if (d_29_v0_) in (_dafny.MultiSet([_dafny.MultiSet([False, False]), _dafny.MultiSet([False])])):
                    coll1_[d_29_v0_] = 737
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.MultiSet
            for compr_2_ in (_dafny.Map({_dafny.MultiSet([True]): len(_dafny.SeqWithoutIsStrInference([False]))})).keys.Elements:
                d_30_v1_: _dafny.MultiSet = compr_2_
                if (d_30_v1_) in (_dafny.Map({_dafny.MultiSet([True]): len(_dafny.SeqWithoutIsStrInference([False]))})):
                    coll2_[d_30_v1_] = len(_dafny.Map({True: True}))
            return _dafny.Map(coll2_)
        return ((iife1_()
        ) | (_dafny.Map({_dafny.MultiSet([True, True]): 790}))) | (iife2_()
        )

    @staticmethod
    def fm38(p0, globalState):
        return D0_DC3(D0_DC3(D0_DC3(D0_DC1(-725, 483))))

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        source3_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): -707}))])
        d_31___mcc_h0_ = source3_
        d_32_cf24_ = d_31___mcc_h0_
        return D12_DC26(_dafny.Map({False: 48}))

    @staticmethod
    def fm40(p0, globalState):
        source4_ = D7_DC17(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gndcte"))): not(True)}))
        if source4_.is_DC18:
            if True:
                return D2_DC7(_dafny.CodePoint('j'), not(True), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhiixh"))))
            elif True:
                return D2_DC7(_dafny.CodePoint('p'), True, -733)
        elif True:
            d_33___mcc_h0_ = source4_.cf23
            d_34_cf23_ = d_33___mcc_h0_
            return D2_DC7(_dafny.CodePoint('v'), False, len(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([True, True])).cardinality)])))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvh")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfclqlf"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxt")))})

    @staticmethod
    def fm42(p0, p1, p2, p3, globalState):
        source5_ = (D12_DC26(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([not(True)]))})) if True else D12_DC26(_dafny.Map({True: -332})))
        if source5_.is_DC27:
            d_35___mcc_h0_ = source5_.cf36
            d_36___mcc_h1_ = source5_.cf37
            d_37___mcc_h2_ = source5_.cf38
            d_38___mcc_h3_ = source5_.cf39
            d_39_cf39_ = d_38___mcc_h3_
            d_40_cf38_ = d_37___mcc_h2_
            d_41_cf37_ = d_36___mcc_h1_
            d_42_cf36_ = d_35___mcc_h0_
            def iife3_():
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: _dafny.Seq
                    for compr_5_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fup")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmp"))])).Elements:
                        d_43_v1_: _dafny.Seq = compr_5_
                        if (d_43_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fup")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmp"))])):
                            coll5_[d_43_v1_] = True
                    return _dafny.Map(coll5_)
                coll3_ = _dafny.Map()
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: _dafny.Seq
                    for compr_4_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fup")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmp"))])).Elements:
                        d_43_v1_: _dafny.Seq = compr_4_
                        if (d_43_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fup")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmp"))])):
                            coll4_[d_43_v1_] = True
                    return _dafny.Map(coll4_)
                compr_3_: _dafny.Seq
                for compr_3_ in (iife4_()
                ).keys.Elements:
                    d_44_v0_: _dafny.Seq = compr_3_
                    if (d_44_v0_) in (iife5_()
                    ):
                        coll3_[d_44_v0_] = True
                return _dafny.Map(coll3_)
            return (iife3_()
            ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbdnvowfp")): False}))
        elif source5_.is_DC26:
            d_45___mcc_h4_ = source5_.cf35
            d_46_cf35_ = d_45___mcc_h4_
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: _dafny.Seq
                for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xegsurtcy"))])).Elements:
                    d_47_v2_: _dafny.Seq = compr_6_
                    if (d_47_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xegsurtcy"))])):
                        coll6_[d_47_v2_] = True
                return _dafny.Map(coll6_)
            return iife6_()
            
        elif True:
            d_48___mcc_h5_ = source5_.cf40
            d_49_cf40_ = d_48___mcc_h5_
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: _dafny.Seq
                for compr_7_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rasmegb")): True})).keys.Elements:
                    d_50_v3_: _dafny.Seq = compr_7_
                    if (d_50_v3_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rasmegb")): True})):
                        coll7_[d_50_v3_] = True
                return _dafny.Map(coll7_)
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")): False})) | (iife7_()
            )

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        if (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))})).ispropersubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))})):
            return (_dafny.Map({315: _dafny.CodePoint('o')})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([6])): _dafny.CodePoint('s')}))
        elif True:
            def iife8_():
                def iife10_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in (_dafny.Map({-535: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})) for d_51_i0_ in range(default__.abs(425))]))})).keys.Elements:
                        d_52_v1_: int = compr_10_
                        if (d_52_v1_) in (_dafny.Map({-535: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})) for d_51_i0_ in range(default__.abs(425))]))})):
                            coll10_[(d_52_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkdcjuouw"))))] = (0) - (-371)
                    return _dafny.Map(coll10_)
                coll8_ = _dafny.Map()
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in (_dafny.Map({-535: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})) for d_51_i0_ in range(default__.abs(425))]))})).keys.Elements:
                        d_52_v1_: int = compr_9_
                        if (d_52_v1_) in (_dafny.Map({-535: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})) for d_51_i0_ in range(default__.abs(425))]))})):
                            coll9_[(d_52_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkdcjuouw"))))] = (0) - (-371)
                    return _dafny.Map(coll9_)
                compr_8_: int
                for compr_8_ in (iife9_()
                ).keys.Elements:
                    d_53_v0_: int = compr_8_
                    if (d_53_v0_) in (iife10_()
                    ):
                        coll8_[(d_53_v0_) * (len(_dafny.SeqWithoutIsStrInference([not(True)])))] = _dafny.CodePoint('e')
                return _dafny.Map(coll8_)
            return (iife8_()
            ) | (_dafny.Map({827: _dafny.CodePoint('b')}))

    @staticmethod
    def fm44(p0, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([not(True)])])

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: D12
            for compr_11_ in (_dafny.Set({D12_DC26(_dafny.Map({True: len(_dafny.Map({False: True}))}))})).Elements:
                d_54_v0_: D12 = compr_11_
                if (d_54_v0_) in (_dafny.Set({D12_DC26(_dafny.Map({True: len(_dafny.Map({False: True}))}))})):
                    coll11_[d_54_v0_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_55_i0_ in range(default__.abs(-170))]))
            return _dafny.Map(coll11_)
        def iife12_():
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in (_dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('t')})): _dafny.CodePoint('t')})).keys.Elements:
                    d_56_v2_: int = compr_14_
                    if (d_56_v2_) in (_dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('t')})): _dafny.CodePoint('t')})):
                        coll14_[(d_56_v2_) + (842)] = _dafny.CodePoint('s')
                return _dafny.Map(coll14_)
            coll12_ = _dafny.Map()
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in (_dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('t')})): _dafny.CodePoint('t')})).keys.Elements:
                    d_56_v2_: int = compr_13_
                    if (d_56_v2_) in (_dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('t')})): _dafny.CodePoint('t')})):
                        coll13_[(d_56_v2_) + (842)] = _dafny.CodePoint('s')
                return _dafny.Map(coll13_)
            compr_12_: D12
            for compr_12_ in (_dafny.Map({D12_DC26(_dafny.Map({True: (0) - (len(_dafny.Map({(0) - (len(_dafny.Map({True: 333}))): len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([D4_DC11(True, False, 251, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uid")))])).cardinality]))})))})): iife13_()
            })).keys.Elements:
                d_57_v1_: D12 = compr_12_
                if (d_57_v1_) in (_dafny.Map({D12_DC26(_dafny.Map({True: (0) - (len(_dafny.Map({(0) - (len(_dafny.Map({True: 333}))): len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([D4_DC11(True, False, 251, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uid")))])).cardinality]))})))})): iife14_()
                })):
                    coll12_[d_57_v1_] = 230
            return _dafny.Map(coll12_)
        def iife15_():
            def iife33_():
                def iife42_():
                    def iife46_():
                        def iife48_():
                            coll48_ = _dafny.Map()
                            compr_48_: int
                            for compr_48_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_48_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll48_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll48_)
                        coll46_ = _dafny.Map()
                        def iife47_():
                            coll47_ = _dafny.Map()
                            compr_47_: int
                            for compr_47_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_47_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll47_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll47_)
                        compr_46_: D12
                        for compr_46_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife47_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_46_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife48_()
)})): _dafny.CodePoint('d')})):
                                coll46_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll46_)
                    coll42_ = _dafny.Map()
                    def iife43_():
                        def iife45_():
                            coll45_ = _dafny.Map()
                            compr_45_: int
                            for compr_45_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_45_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll45_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll45_)
                        coll43_ = _dafny.Map()
                        def iife44_():
                            coll44_ = _dafny.Map()
                            compr_44_: int
                            for compr_44_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_44_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll44_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll44_)
                        compr_43_: D12
                        for compr_43_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife44_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_43_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife45_()
)})): _dafny.CodePoint('d')})):
                                coll43_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll43_)
                    compr_42_: D12
                    for compr_42_ in (iife43_()
                    ).keys.Elements:
                        d_60_v4_: D12 = compr_42_
                        if (d_60_v4_) in (iife46_()
                        ):
                            def iife49_():
                                coll49_ = _dafny.Map()
                                compr_49_: str
                                for compr_49_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))).Elements:
                                    d_61_v7_: str = compr_49_
                                    if (d_61_v7_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))):
                                        coll49_[d_61_v7_] = True
                                return _dafny.Map(coll49_)
                            coll42_[d_60_v4_] = len(iife49_()
                            )
                    return _dafny.Map(coll42_)
                coll33_ = _dafny.Map()
                def iife34_():
                    def iife38_():
                        def iife40_():
                            coll40_ = _dafny.Map()
                            compr_40_: int
                            for compr_40_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_40_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll40_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll40_)
                        coll38_ = _dafny.Map()
                        def iife39_():
                            coll39_ = _dafny.Map()
                            compr_39_: int
                            for compr_39_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_39_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll39_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll39_)
                        compr_38_: D12
                        for compr_38_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife39_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_38_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife40_()
)})): _dafny.CodePoint('d')})):
                                coll38_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll38_)
                    coll34_ = _dafny.Map()
                    def iife35_():
                        def iife37_():
                            coll37_ = _dafny.Map()
                            compr_37_: int
                            for compr_37_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_37_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll37_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll37_)
                        coll35_ = _dafny.Map()
                        def iife36_():
                            coll36_ = _dafny.Map()
                            compr_36_: int
                            for compr_36_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_36_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll36_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll36_)
                        compr_35_: D12
                        for compr_35_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife36_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_35_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife37_()
)})): _dafny.CodePoint('d')})):
                                coll35_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll35_)
                    compr_34_: D12
                    for compr_34_ in (iife35_()
                    ).keys.Elements:
                        d_60_v4_: D12 = compr_34_
                        if (d_60_v4_) in (iife38_()
                        ):
                            def iife41_():
                                coll41_ = _dafny.Map()
                                compr_41_: str
                                for compr_41_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))).Elements:
                                    d_61_v7_: str = compr_41_
                                    if (d_61_v7_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))):
                                        coll41_[d_61_v7_] = True
                                return _dafny.Map(coll41_)
                            coll34_[d_60_v4_] = len(iife41_()
                            )
                    return _dafny.Map(coll34_)
                compr_33_: D12
                for compr_33_ in (iife34_()
                ).keys.Elements:
                    d_62_v3_: D12 = compr_33_
                    if (d_62_v3_) in (iife42_()
                    ):
                        coll33_[d_62_v3_] = 754
                return _dafny.Map(coll33_)
            coll15_ = _dafny.Set()
            def iife16_():
                def iife25_():
                    def iife29_():
                        def iife31_():
                            coll31_ = _dafny.Map()
                            compr_31_: int
                            for compr_31_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_31_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll31_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll31_)
                        coll29_ = _dafny.Map()
                        def iife30_():
                            coll30_ = _dafny.Map()
                            compr_30_: int
                            for compr_30_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_30_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll30_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll30_)
                        compr_29_: D12
                        for compr_29_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife30_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_29_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife31_()
)})): _dafny.CodePoint('d')})):
                                coll29_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll29_)
                    coll25_ = _dafny.Map()
                    def iife26_():
                        def iife28_():
                            coll28_ = _dafny.Map()
                            compr_28_: int
                            for compr_28_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_28_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll28_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll28_)
                        coll26_ = _dafny.Map()
                        def iife27_():
                            coll27_ = _dafny.Map()
                            compr_27_: int
                            for compr_27_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_27_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll27_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll27_)
                        compr_26_: D12
                        for compr_26_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife27_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_26_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife28_()
)})): _dafny.CodePoint('d')})):
                                coll26_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll26_)
                    compr_25_: D12
                    for compr_25_ in (iife26_()
                    ).keys.Elements:
                        d_60_v4_: D12 = compr_25_
                        if (d_60_v4_) in (iife29_()
                        ):
                            def iife32_():
                                coll32_ = _dafny.Map()
                                compr_32_: str
                                for compr_32_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))).Elements:
                                    d_61_v7_: str = compr_32_
                                    if (d_61_v7_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))):
                                        coll32_[d_61_v7_] = True
                                return _dafny.Map(coll32_)
                            coll25_[d_60_v4_] = len(iife32_()
                            )
                    return _dafny.Map(coll25_)
                coll16_ = _dafny.Map()
                def iife17_():
                    def iife21_():
                        def iife23_():
                            coll23_ = _dafny.Map()
                            compr_23_: int
                            for compr_23_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_23_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll23_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll23_)
                        coll21_ = _dafny.Map()
                        def iife22_():
                            coll22_ = _dafny.Map()
                            compr_22_: int
                            for compr_22_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_22_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll22_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll22_)
                        compr_21_: D12
                        for compr_21_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife22_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_21_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife23_()
)})): _dafny.CodePoint('d')})):
                                coll21_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll21_)
                    coll17_ = _dafny.Map()
                    def iife18_():
                        def iife20_():
                            coll20_ = _dafny.Map()
                            compr_20_: int
                            for compr_20_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_20_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll20_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll20_)
                        coll18_ = _dafny.Map()
                        def iife19_():
                            coll19_ = _dafny.Map()
                            compr_19_: int
                            for compr_19_ in _dafny.IntegerRange(-428, 530):
                                d_58_v6_: int = compr_19_
                                if ((-428) <= (d_58_v6_)) and ((d_58_v6_) < (530)):
                                    coll19_[default__.safeDivisionInt(d_58_v6_, -206)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daydimu")))
                            return _dafny.Map(coll19_)
                        compr_18_: D12
                        for compr_18_ in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife19_()
)})): _dafny.CodePoint('d')})).keys.Elements:
                            d_59_v5_: D12 = compr_18_
                            if (d_59_v5_) in (_dafny.Map({D12_DC26(_dafny.Map({False: len(iife20_()
)})): _dafny.CodePoint('d')})):
                                coll18_[d_59_v5_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                        return _dafny.Map(coll18_)
                    compr_17_: D12
                    for compr_17_ in (iife18_()
                    ).keys.Elements:
                        d_60_v4_: D12 = compr_17_
                        if (d_60_v4_) in (iife21_()
                        ):
                            def iife24_():
                                coll24_ = _dafny.Map()
                                compr_24_: str
                                for compr_24_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))).Elements:
                                    d_61_v7_: str = compr_24_
                                    if (d_61_v7_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))):
                                        coll24_[d_61_v7_] = True
                                return _dafny.Map(coll24_)
                            coll17_[d_60_v4_] = len(iife24_()
                            )
                    return _dafny.Map(coll17_)
                compr_16_: D12
                for compr_16_ in (iife17_()
                ).keys.Elements:
                    d_62_v3_: D12 = compr_16_
                    if (d_62_v3_) in (iife25_()
                    ):
                        coll16_[d_62_v3_] = 754
                return _dafny.Map(coll16_)
            compr_15_: _dafny.Map
            for compr_15_ in (_dafny.Map({iife16_()
            : _dafny.CodePoint('m')})).keys.Elements:
                d_63_v8_: _dafny.Map = compr_15_
                if (d_63_v8_) in (_dafny.Map({iife33_()
                : _dafny.CodePoint('m')})):
                    coll15_ = coll15_.union(_dafny.Set([d_63_v8_]))
            return _dafny.Set(coll15_)
        return ((_dafny.Set({_dafny.Map({D12_DC26(_dafny.Map({False: 754})): 665}), _dafny.Map({D12_DC26(_dafny.Map({False: 202})): len(_dafny.Map({not(not(True)): 198}))})}) if not(not(True)) else _dafny.Set({_dafny.Map({D12_DC26(_dafny.Map({False: -239})): 616}), iife11_()
        }))) - ((_dafny.Set({iife12_()
        })).intersection(iife15_()
        ))

    @staticmethod
    def fm46(globalState):
        return ((_dafny.Map({False: (_dafny.MultiSet([not(True)])).cardinality})) | (_dafny.Map({True: 424}))) | ((_dafny.Map({True: (0) - (len(_dafny.Set({not(not(not(True)))})))})) | (_dafny.Map({False: 470})))

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('x')]))}): (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_64_i0_ in range(default__.abs(879))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjv")))})

    @staticmethod
    def fm48(p0, p1, globalState):
        return _dafny.Map({(True if False else not(False)): D7_DC18()})

    @staticmethod
    def fm49(p0, globalState):
        return ((_dafny.MultiSet([-701])).cardinality) > (len(_dafny.SeqWithoutIsStrInference([926 for d_65_i0_ in range(default__.abs(959))])))

    @staticmethod
    def fm50(globalState):
        if True:
            return D3_DC9()
        elif True:
            return D3_DC9()

    @staticmethod
    def fm51(p0, globalState):
        return False

    @staticmethod
    def fm52(globalState):
        source6_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))
        d_66___mcc_h0_ = source6_
        d_67_cf34_ = d_66___mcc_h0_
        def iife50_():
            coll50_ = _dafny.Set()
            compr_50_: int
            for compr_50_ in _dafny.IntegerRange(944, 690):
                d_68_v0_: int = compr_50_
                if ((944) <= (d_68_v0_)) and ((d_68_v0_) < (690)):
                    coll50_ = coll50_.union(_dafny.Set([default__.safeDivisionInt(d_68_v0_, len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgpvcbka")))})))]))
            return _dafny.Set(coll50_)
        return _dafny.Map({369: len(_dafny.SeqWithoutIsStrInference([len(iife50_()
        )]))})

    @staticmethod
    def fm53(p0, p1, globalState):
        def iife51_():
            coll51_ = _dafny.Map()
            compr_51_: int
            for compr_51_ in _dafny.IntegerRange(-717, 130):
                d_69_v0_: int = compr_51_
                if ((-717) <= (d_69_v0_)) and ((d_69_v0_) < (130)):
                    coll51_[(d_69_v0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_70_i0_ in range(default__.abs(728))])))] = True
            return _dafny.Map(coll51_)
        def iife52_():
            coll52_ = _dafny.Set()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(634, 678):
                d_71_v1_: int = compr_52_
                if ((634) <= (d_71_v1_)) and ((d_71_v1_) < (678)):
                    coll52_ = coll52_.union(_dafny.Set([default__.safeModuloInt(d_71_v1_, len(_dafny.Map({False: D0_DC0(len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oghwvum"))))])))})))]))
            return _dafny.Set(coll52_)
        return (_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kc"))): False}), iife51_()
        })).intersection(_dafny.Set({_dafny.Map({len(iife52_()
        ): True})}))

    @staticmethod
    def fm54(p0, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('d')]) if True else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('i')]))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_72_i0_ in range(default__.abs(-406))])))

    @staticmethod
    def m15(globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        r3: C0 = None
        d_73_v0_: _dafny.Seq
        d_73_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpud"))
        d_74_v1_: int
        d_74_v1_ = 36
        hi0_ = d_74_v1_
        for d_75_i0_ in range(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bheo"))) + (d_73_v0_)), hi0_):
            d_76_v2_: bool
            d_76_v2_ = True
            d_77_v3_: _dafny.Array
            nw0_ = _dafny.Array(None, 1)
            nw0_[int(0)] = d_76_v2_
            d_77_v3_ = nw0_
            index0_ = default__.safeIndex(925, (d_77_v3_).length(0))
            (d_77_v3_)[index0_] = d_76_v2_
            index1_ = default__.safeIndex(925, (d_77_v3_).length(0))
            (d_77_v3_)[index1_] = d_76_v2_
            (globalState).f8 = d_75_i0_
            d_78_v4_: _dafny.Seq
            d_78_v4_ = _dafny.SeqWithoutIsStrInference([d_76_v2_, d_76_v2_])
            d_79_v5_: _dafny.MultiSet
            d_79_v5_ = _dafny.MultiSet([True, d_76_v2_])
            d_80_v6_: _dafny.Seq
            d_80_v6_ = _dafny.SeqWithoutIsStrInference([d_74_v1_, len(_dafny.SeqWithoutIsStrInference([d_74_v1_]))])
            (globalState).f3 = ((len(d_78_v4_)) * ((_dafny.MultiSet([d_74_v1_, d_74_v1_, (d_79_v5_).cardinality, (d_79_v5_).cardinality])).cardinality)) - ((d_80_v6_)[default__.safeIndex(d_74_v1_, len(d_80_v6_))])
            d_76_v2_ = not(d_76_v2_)
        (globalState).f3 = d_74_v1_
        (globalState).f3 = len(d_73_v0_)
        d_81_v7_: bool
        d_81_v7_ = False
        d_82_v9_: _dafny.Map
        def iife53_():
            coll53_ = _dafny.Set()
            compr_53_: int
            for compr_53_ in _dafny.IntegerRange(-904, 973):
                d_83_v8_: int = compr_53_
                if ((-904) <= (d_83_v8_)) and ((d_83_v8_) < (973)):
                    coll53_ = coll53_.union(_dafny.Set([(d_83_v8_) * (d_74_v1_)]))
            return _dafny.Set(coll53_)
        d_82_v9_ = _dafny.Map({d_81_v7_: iife53_()
        })
        d_84_v10_: _dafny.Seq
        d_84_v10_ = _dafny.SeqWithoutIsStrInference([d_74_v1_, d_74_v1_, d_74_v1_, 657])
        (globalState).f3 = (0) - (((len(d_82_v9_)) * ((d_84_v10_)[default__.safeIndex(d_74_v1_, len(d_84_v10_))]) if True else d_74_v1_))
        d_85_v11_: _dafny.Array
        nw1_ = _dafny.Array(False, 21)
        d_85_v11_ = nw1_
        d_86_v12_: _dafny.Seq
        d_86_v12_ = _dafny.SeqWithoutIsStrInference([d_85_v11_])
        d_85_v11_ = (d_86_v12_)[default__.safeIndex((0) - (d_74_v1_), len(d_86_v12_))]
        d_87_v13_: _dafny.Array
        nw2_ = _dafny.Array(int(0), 11)
        d_87_v13_ = nw2_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_87_v13_).length(0)):
            d_88_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_88_i1_)) and ((d_88_i1_) < ((d_87_v13_).length(0)))):
                (d_87_v13_)[(d_88_i1_)] = default__.safeModuloInt(d_88_i1_, default__.safeModuloInt(d_74_v1_, d_74_v1_))
        r0 = d_74_v1_
        d_89_v14_: _dafny.Map
        d_89_v14_ = _dafny.Map({d_87_v13_: default__.fm54(d_74_v1_, globalState)})
        d_90_v15_: _dafny.Map
        d_90_v15_ = d_89_v14_
        r1 = (d_90_v15_)
        r2 = d_81_v7_
        d_91_v16_: C0
        nw3_ = C0()
        nw3_.ctor__()
        d_91_v16_ = nw3_
        r3 = (d_91_v16_ if d_81_v7_ else d_91_v16_)
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_92_v0_: int
        d_92_v0_ = 545
        d_93_v1_: _dafny.Seq
        d_93_v1_ = _dafny.SeqWithoutIsStrInference([d_92_v0_])
        d_94_v2_: bool
        d_94_v2_ = False
        d_95_v3_: _dafny.Map
        d_95_v3_ = _dafny.Map({d_94_v2_: d_92_v0_})
        d_96_v4_: D0
        d_96_v4_ = D0_DC0(len(d_95_v3_))
        d_97_v5_: _dafny.Seq
        d_97_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cj"))
        d_98_v6_: _dafny.Set
        d_98_v6_ = _dafny.Set({d_97_v5_})
        d_99_v7_: _dafny.MultiSet
        d_99_v7_ = _dafny.MultiSet([d_92_v0_])
        d_100_globalState_: GlobalState
        nw4_ = GlobalState()
        nw4_.ctor__(((d_93_v1_).set(default__.safeIndex((d_96_v4_).cf0, len(d_93_v1_)), d_92_v0_)).set(default__.safeIndex(d_92_v0_, len((d_93_v1_).set(default__.safeIndex((d_96_v4_).cf0, len(d_93_v1_)), d_92_v0_))), len(_dafny.Map({d_94_v2_: len(d_98_v6_)}))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) + (d_97_v5_), d_99_v7_, 742, False, False, True, 339, 344)
        d_100_globalState_ = nw4_
        d_101_v8_: _dafny.Array
        def lambda0_(d_102_v2_):
            def lambda1_(d_103_i0_):
                return (d_103_i0_) * (len(_dafny.Set({d_102_v2_, d_102_v2_})))

            return lambda1_

        init0_ = lambda0_(d_94_v2_)
        nw5_ = _dafny.Array(None, 16)
        for i0_0_ in range(nw5_.length(0)):
            nw5_[i0_0_] = init0_(i0_0_)
        d_101_v8_ = nw5_
        d_104_v9_: D0
        d_104_v9_ = D0_DC1(d_92_v0_, d_92_v0_)
        rhs0_ = d_101_v8_
        rhs1_ = (d_104_v9_).cf1
        lhs0_ = d_100_globalState_
        d_101_v8_ = rhs0_
        lhs0_.f3 = rhs1_
        d_105_i1_: int
        d_105_i1_ = 0
        with _dafny.label("0"):
            while d_94_v2_:
                with _dafny.c_label("0"):
                    if (d_105_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_105_i1_ = (d_105_i1_) + (1)
                    index2_ = default__.safeIndex(982, (d_101_v8_).length(0))
                    (d_101_v8_)[index2_] = d_92_v0_
                    index3_ = default__.safeIndex(982, (d_101_v8_).length(0))
                    (d_101_v8_)[index3_] = d_92_v0_
                    source7_ = d_96_v4_
                    if source7_.is_DC1:
                        d_106___mcc_h0_ = source7_.cf1
                        d_107___mcc_h1_ = source7_.cf2
                        d_108_cf2_ = d_107___mcc_h1_
                        d_109_cf1_ = d_106___mcc_h0_
                        d_110_v10_: _dafny.Array
                        nw6_ = _dafny.Array(int(0), 24)
                        d_110_v10_ = nw6_
                        def lambda2_(d_111_cf2_):
                            def lambda3_(d_112_i2_):
                                return (d_112_i2_) - (d_111_cf2_)

                            return lambda3_

                        init1_ = lambda2_(d_108_cf2_)
                        nw7_ = _dafny.Array(None, 14)
                        for i0_1_ in range(nw7_.length(0)):
                            nw7_[i0_1_] = init1_(i0_1_)
                        d_110_v10_ = nw7_
                        (d_100_globalState_).f3 = 716
                        (d_100_globalState_).f2 = d_99_v7_
                        index4_ = default__.safeIndex(982, (d_101_v8_).length(0))
                        rhs2_ = d_94_v2_
                        rhs3_ = (0) - (default__.safeModuloInt(d_108_cf2_, d_92_v0_))
                        rhs4_ = (d_92_v0_) != ((d_109_cf1_) * (d_109_cf1_))
                        lhs1_ = d_100_globalState_
                        lhs2_ = d_101_v8_
                        lhs3_ = default__.safeIndex(982, (d_101_v8_).length(0))
                        lhs4_ = d_100_globalState_
                        lhs1_.f4 = rhs2_
                        lhs2_[lhs3_] = rhs3_
                        lhs4_.f4 = rhs4_
                    elif source7_.is_DC2:
                        d_113_v11_: _dafny.Array
                        nw8_ = _dafny.Array(int(0), 3)
                        d_113_v11_ = nw8_
                        d_113_v11_ = d_113_v11_
                        d_95_v3_ = (d_95_v3_).set(not(not (d_94_v2_) or (d_94_v2_)), (d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))])
                        d_114_v12_: _dafny.Map
                        d_114_v12_ = _dafny.Map({(d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))]: d_92_v0_})
                        (d_100_globalState_).f0 = (d_93_v1_ if d_94_v2_ else _dafny.SeqWithoutIsStrInference([d_92_v0_, len(d_114_v12_)]))
                        d_95_v3_ = d_95_v3_
                    elif source7_.is_DC0:
                        d_115___mcc_h2_ = source7_.cf0
                        d_116_cf0_ = d_115___mcc_h2_
                        (d_100_globalState_).f8 = (default__.safeDivisionInt(d_116_cf0_, (d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))])) - ((D0_DC0((default__.fm0(d_116_cf0_, d_100_globalState_)).cardinality)).cf0)
                        (d_100_globalState_).f4 = d_94_v2_
                        d_117_v13_: _dafny.Array
                        nw9_ = _dafny.Array(None, 1)
                        nw9_[int(0)] = d_93_v1_
                        d_117_v13_ = nw9_
                        d_118_v14_: _dafny.Seq
                        d_118_v14_ = _dafny.SeqWithoutIsStrInference([d_94_v2_, d_94_v2_, d_94_v2_, False, False])
                        d_119_v15_: _dafny.Set
                        d_119_v15_ = _dafny.Set({d_94_v2_})
                        d_120_v16_: _dafny.Map
                        d_120_v16_ = _dafny.Map({len(d_119_v15_): d_116_cf0_})
                        index5_ = default__.safeIndex(529, (d_117_v13_).length(0))
                        (d_117_v13_)[index5_] = default__.fm1((0) - (len(_dafny.SeqWithoutIsStrInference([d_116_cf0_, (0) - (len(d_118_v14_)), len(d_97_v5_), 454, len(d_97_v5_)]))), _dafny.CodePoint('f'), not(d_94_v2_), d_120_v16_, d_100_globalState_)
                        index6_ = default__.safeIndex(529, (d_117_v13_).length(0))
                        (d_117_v13_)[index6_] = d_93_v1_
                        d_92_v0_ = (0) - (default__.safeModuloInt(681, (d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))]))
                    elif True:
                        d_121___mcc_h3_ = source7_.cf3
                        d_122_cf3_ = d_121___mcc_h3_
                        (d_100_globalState_).f8 = (0) - ((d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))])
                        d_123_v17_: C5
                        nw10_ = C5()
                        nw10_.ctor__()
                        d_123_v17_ = nw10_
                        d_124_v18_: _dafny.Set
                        d_124_v18_ = _dafny.Set({(d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))], d_92_v0_})
                        rhs5_ = ((d_97_v5_)[default__.safeIndex(d_92_v0_, len(d_97_v5_))]) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ethjt")))
                        rhs6_ = d_97_v5_
                        rhs7_ = (len(d_124_v18_)) >= (default__.safeDivisionInt((0) - ((d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))]), (d_101_v8_)[default__.safeIndex(982, (d_101_v8_).length(0))]))
                        rhs8_ = (d_97_v5_) + (d_97_v5_)
                        lhs5_ = d_100_globalState_
                        d_94_v2_ = rhs5_
                        d_97_v5_ = rhs6_
                        lhs5_.f4 = rhs7_
                        d_97_v5_ = rhs8_
                        (d_100_globalState_).f4 = d_94_v2_
                    d_125_v19_: C8
                    nw11_ = C8()
                    nw11_.ctor__()
                    d_125_v19_ = nw11_
                    d_126_v20_: _dafny.Array
                    def lambda4_(d_127_v2_):
                        def lambda5_(d_128_i3_):
                            return d_127_v2_

                        return lambda5_

                    init2_ = lambda4_(d_94_v2_)
                    nw12_ = _dafny.Array(None, 22)
                    for i0_2_ in range(nw12_.length(0)):
                        nw12_[i0_2_] = init2_(i0_2_)
                    d_126_v20_ = nw12_
                    index7_ = default__.safeIndex(998, (d_126_v20_).length(0))
                    (d_126_v20_)[index7_] = not(d_94_v2_)
                    d_129_v21_: _dafny.Seq
                    d_129_v21_ = _dafny.SeqWithoutIsStrInference([d_94_v2_, d_94_v2_, False, d_94_v2_])
                    d_130_v22_: _dafny.MultiSet
                    d_130_v22_ = _dafny.MultiSet([d_94_v2_])
                    d_131_v23_: _dafny.MultiSet
                    d_131_v23_ = _dafny.MultiSet([(_dafny.MultiSet(d_129_v21_)).set(d_94_v2_, default__.abs(d_92_v0_)), d_130_v22_, d_130_v22_])
                    index8_ = default__.safeIndex(998, (d_126_v20_).length(0))
                    (d_126_v20_)[index8_] = not (default__.fm49(default__.fm49(True, d_100_globalState_), d_100_globalState_)) or ((d_131_v23_).ispropersubset(d_131_v23_))
                    pass
            pass
        d_132_v24_: _dafny.Set
        d_132_v24_ = _dafny.Set({len(d_97_v5_), d_92_v0_})
        if not((d_132_v24_).issubset(d_132_v24_)):
            if d_94_v2_:
                d_133_v25_: _dafny.Array
                nw13_ = _dafny.Array(None, 2)
                nw13_[int(0)] = d_94_v2_
                nw13_[int(1)] = not(d_94_v2_)
                d_133_v25_ = nw13_
                d_134_v26_: _dafny.Map
                d_134_v26_ = _dafny.Map({d_92_v0_: d_133_v25_})
                d_134_v26_ = d_134_v26_
                d_135_v27_: C5
                nw14_ = C5()
                nw14_.ctor__()
                d_135_v27_ = nw14_
                d_136_v28_: _dafny.Array
                nw15_ = _dafny.Array(None, 5)
                nw15_[int(0)] = d_135_v27_
                nw15_[int(1)] = d_135_v27_
                nw15_[int(2)] = d_135_v27_
                nw15_[int(3)] = d_135_v27_
                nw15_[int(4)] = d_135_v27_
                d_136_v28_ = nw15_
                index9_ = default__.safeIndex(780, (d_136_v28_).length(0))
                (d_136_v28_)[index9_] = d_135_v27_
                index10_ = default__.safeIndex(780, (d_136_v28_).length(0))
                (d_136_v28_)[index10_] = d_135_v27_
                d_137_v29_: _dafny.Array
                def lambda6_(d_138_i4_):
                    return _dafny.CodePoint('l')

                init3_ = lambda6_
                nw16_ = _dafny.Array(None, 26)
                for i0_3_ in range(nw16_.length(0)):
                    nw16_[i0_3_] = init3_(i0_3_)
                d_137_v29_ = nw16_
                d_139_v30_: str
                d_139_v30_ = _dafny.CodePoint('d')
                index11_ = default__.safeIndex(346, (d_137_v29_).length(0))
                (d_137_v29_)[index11_] = d_139_v30_
                index12_ = default__.safeIndex(346, (d_137_v29_).length(0))
                (d_137_v29_)[index12_] = _dafny.CodePoint('h')
                d_140_v31_: C5
                nw17_ = C5()
                nw17_.ctor__()
                d_140_v31_ = nw17_
                d_141_v32_: C9
                nw18_ = C9()
                nw18_.ctor__(d_94_v2_)
                d_141_v32_ = nw18_
            elif True:
                d_142_v33_: D23
                d_142_v33_ = D23_DC46()
                d_143_v34_: _dafny.MultiSet
                d_143_v34_ = _dafny.MultiSet([d_142_v33_])
                d_144_v35_: _dafny.Map
                d_144_v35_ = _dafny.Map({True: d_143_v34_})
                d_145_v36_: _dafny.Map
                d_145_v36_ = _dafny.Map({len(d_144_v35_): (0) - (d_92_v0_)})
                d_145_v36_ = default__.fm52(d_100_globalState_)
                (d_100_globalState_).f3 = d_92_v0_
                d_146_v37_: _dafny.Array
                def lambda7_(d_147_v2_):
                    def lambda8_(d_148_i5_):
                        return d_147_v2_

                    return lambda8_

                init4_ = lambda7_(d_94_v2_)
                nw19_ = _dafny.Array(None, 6)
                for i0_4_ in range(nw19_.length(0)):
                    nw19_[i0_4_] = init4_(i0_4_)
                d_146_v37_ = nw19_
                index13_ = default__.safeIndex(438, (d_146_v37_).length(0))
                (d_146_v37_)[index13_] = not(d_94_v2_)
                index14_ = default__.safeIndex(438, (d_146_v37_).length(0))
                (d_146_v37_)[index14_] = (d_92_v0_) >= (d_92_v0_)
                d_149_v38_: _dafny.Map
                d_149_v38_ = _dafny.Map({d_101_v8_: d_146_v37_})
                d_149_v38_ = d_149_v38_
                index15_ = default__.safeIndex(438, (d_146_v37_).length(0))
                (d_146_v37_)[index15_] = (len(d_97_v5_)) < (d_92_v0_)
            d_150_v39_: int
            d_151_v40_: _dafny.Map
            d_152_v41_: bool
            d_153_v42_: C0
            out0_: int
            out1_: _dafny.Map
            out2_: bool
            out3_: C0
            out0_, out1_, out2_, out3_ = default__.m15(d_100_globalState_)
            d_150_v39_ = out0_
            d_151_v40_ = out1_
            d_152_v41_ = out2_
            d_153_v42_ = out3_
            (d_100_globalState_).f4 = d_152_v41_
            d_154_v43_: D10
            d_154_v43_ = D10_DC24(len(default__.fm23(d_97_v5_, d_100_globalState_)), d_152_v41_, d_97_v5_)
            d_155_v44_: _dafny.MultiSet
            d_155_v44_ = _dafny.MultiSet([d_152_v41_, (d_154_v43_).cf32, (d_97_v5_) < (d_97_v5_)])
            d_155_v44_ = ((d_155_v44_).set(not(d_94_v2_), default__.abs(len(default__.fm35(d_152_v41_, d_92_v0_, d_94_v2_, d_152_v41_, d_100_globalState_))))) - (d_155_v44_)
            d_156_v45_: int
            d_157_v46_: _dafny.Map
            d_158_v47_: bool
            d_159_v48_: C0
            out4_: int
            out5_: _dafny.Map
            out6_: bool
            out7_: C0
            out4_, out5_, out6_, out7_ = default__.m15(d_100_globalState_)
            d_156_v45_ = out4_
            d_157_v46_ = out5_
            d_158_v47_ = out6_
            d_159_v48_ = out7_
        elif True:
            if (default__.fm18(d_94_v2_, d_94_v2_, d_100_globalState_)) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))):
                (d_100_globalState_).f8 = (d_93_v1_)[default__.safeIndex(d_92_v0_, len(d_93_v1_))]
                (d_100_globalState_).f8 = (d_92_v0_) + (d_92_v0_)
                index16_ = default__.safeIndex(919, (d_101_v8_).length(0))
                (d_101_v8_)[index16_] = d_92_v0_
                index17_ = default__.safeIndex(919, (d_101_v8_).length(0))
                (d_101_v8_)[index17_] = default__.safeModuloInt((0) - (d_92_v0_), len(_dafny.SeqWithoutIsStrInference([len(d_95_v3_)])))
                (d_100_globalState_).f0 = d_93_v1_
                index18_ = default__.safeIndex(919, (d_101_v8_).length(0))
                (d_101_v8_)[index18_] = d_92_v0_
            elif True:
                d_160_v49_: C2
                nw20_ = C2()
                nw20_.ctor__()
                d_160_v49_ = nw20_
                d_161_v50_: C7
                nw21_ = C7()
                nw21_.ctor__()
                d_161_v50_ = nw21_
                d_162_v51_: _dafny.Map
                d_162_v51_ = _dafny.Map({d_161_v50_: (d_161_v50_).fm3(d_93_v1_, d_99_v7_, d_94_v2_, d_100_globalState_)})
                d_163_v52_: _dafny.Map
                d_163_v52_ = _dafny.Map({d_94_v2_: d_162_v51_})
                (d_100_globalState_).f4 = ((d_162_v51_) != (((d_163_v52_)[False] if (False) in (d_163_v52_) else d_162_v51_))) == (d_94_v2_)
                d_164_v53_: C9
                nw22_ = C9()
                nw22_.ctor__(d_94_v2_)
                d_164_v53_ = nw22_
                d_165_v54_: _dafny.Map
                d_165_v54_ = _dafny.Map({d_92_v0_: d_164_v53_})
                d_166_v55_: T0
                nw23_ = C5()
                nw23_.ctor__()
                d_166_v55_ = nw23_
                d_167_v56_: _dafny.Map
                d_167_v56_ = _dafny.Map({d_166_v55_: d_92_v0_})
                d_168_v57_: D26
                d_168_v57_ = D26_DC55(d_166_v55_)
                d_169_v58_: _dafny.Seq
                d_169_v58_ = _dafny.SeqWithoutIsStrInference([d_94_v2_, d_94_v2_])
                d_170_v59_: _dafny.Array
                nw24_ = _dafny.Array(None, 5)
                nw24_[int(0)] = d_165_v54_
                nw24_[int(1)] = _dafny.Map({((d_167_v56_)[(d_168_v57_).cf77] if ((d_168_v57_).cf77) in (d_167_v56_) else len(d_169_v58_)): d_164_v53_})
                nw24_[int(2)] = _dafny.Map({d_92_v0_: d_164_v53_})
                nw24_[int(3)] = d_165_v54_
                nw24_[int(4)] = d_165_v54_
                d_170_v59_ = nw24_
                d_170_v59_ = d_170_v59_
                d_97_v5_ = d_97_v5_
                (d_100_globalState_).f8 = d_92_v0_
            d_171_v60_: C8
            nw25_ = C8()
            nw25_.ctor__()
            d_171_v60_ = nw25_
            d_172_v61_: _dafny.Map
            d_172_v61_ = _dafny.Map({d_94_v2_: d_97_v5_})
            (d_171_v60_).m4(((0) - (d_92_v0_)) == (d_92_v0_), (_dafny.Map({d_94_v2_: d_97_v5_})) | (d_172_v61_), 709, d_100_globalState_)
            if d_94_v2_:
                d_101_v8_ = (d_101_v8_ if not(d_94_v2_) else d_101_v8_)
                (d_100_globalState_).f4 = not(d_94_v2_)
                (d_171_v60_).m4(d_94_v2_, d_172_v61_, d_92_v0_, d_100_globalState_)
                d_173_v62_: _dafny.Seq
                d_173_v62_ = _dafny.SeqWithoutIsStrInference([d_94_v2_])
                d_173_v62_ = (d_173_v62_) + (_dafny.SeqWithoutIsStrInference([not(d_94_v2_), False, d_94_v2_]))
                d_101_v8_ = d_101_v8_
            elif True:
                d_174_v63_: _dafny.Map
                d_174_v63_ = _dafny.Map({d_94_v2_: False})
                d_175_v64_: _dafny.Array
                nw26_ = _dafny.Array(None, 12)
                nw26_[int(0)] = d_95_v3_
                nw26_[int(1)] = (_dafny.Map({d_94_v2_: -686})).set(d_94_v2_, d_92_v0_)
                nw26_[int(2)] = (default__.fm46(d_100_globalState_)).set(((d_174_v63_)[not(d_94_v2_)] if (not(d_94_v2_)) in (d_174_v63_) else d_94_v2_), 379)
                nw26_[int(3)] = d_95_v3_
                nw26_[int(4)] = d_95_v3_
                nw26_[int(5)] = d_95_v3_
                nw26_[int(6)] = d_95_v3_
                nw26_[int(7)] = _dafny.Map({False: d_92_v0_})
                nw26_[int(8)] = _dafny.Map({d_94_v2_: d_92_v0_})
                nw26_[int(9)] = d_95_v3_
                nw26_[int(10)] = d_95_v3_
                nw26_[int(11)] = d_95_v3_
                d_175_v64_ = nw26_
                d_176_v65_: T0
                nw27_ = C4()
                nw27_.ctor__(d_175_v64_)
                d_176_v65_ = nw27_
                d_176_v65_ = d_176_v65_
                d_177_v66_: _dafny.Array
                nw28_ = _dafny.Array(None, 6)
                nw28_[int(0)] = (d_93_v1_ if True else d_93_v1_)
                nw28_[int(1)] = d_93_v1_
                nw28_[int(2)] = d_93_v1_
                nw28_[int(3)] = _dafny.SeqWithoutIsStrInference([d_92_v0_, d_92_v0_, d_92_v0_])
                nw28_[int(4)] = d_93_v1_
                nw28_[int(5)] = _dafny.SeqWithoutIsStrInference([d_92_v0_ for d_178_i6_ in range(default__.abs(813))])
                d_177_v66_ = nw28_
                index19_ = default__.safeIndex(459, (d_177_v66_).length(0))
                (d_177_v66_)[index19_] = d_93_v1_
                index20_ = default__.safeIndex(459, (d_177_v66_).length(0))
                (d_177_v66_)[index20_] = (d_93_v1_) + (d_93_v1_)
                d_179_v67_: _dafny.Array
                def lambda9_(d_180_v3_):
                    def lambda10_(d_181_i7_):
                        return (d_180_v3_) != (d_180_v3_)

                    return lambda10_

                init5_ = lambda9_(d_95_v3_)
                nw29_ = _dafny.Array(None, 21)
                for i0_5_ in range(nw29_.length(0)):
                    nw29_[i0_5_] = init5_(i0_5_)
                d_179_v67_ = nw29_
                index21_ = default__.safeIndex(229, (d_179_v67_).length(0))
                (d_179_v67_)[index21_] = (len(d_97_v5_)) <= (d_92_v0_)
                index22_ = default__.safeIndex(229, (d_179_v67_).length(0))
                (d_179_v67_)[index22_] = not(d_94_v2_)
                (d_100_globalState_).f4 = not(not ((d_92_v0_) <= (len(_dafny.SeqWithoutIsStrInference([((d_174_v63_)[d_94_v2_] if (d_94_v2_) in (d_174_v63_) else (d_179_v67_)[default__.safeIndex(229, (d_179_v67_).length(0))])])))) or ((d_179_v67_)[default__.safeIndex(229, (d_179_v67_).length(0))]))
                d_182_v68_: _dafny.MultiSet
                d_182_v68_ = _dafny.MultiSet([True])
                d_183_v69_: _dafny.Seq
                d_183_v69_ = _dafny.SeqWithoutIsStrInference([d_182_v68_])
                d_184_v70_: _dafny.Seq
                d_184_v70_ = _dafny.SeqWithoutIsStrInference([((d_177_v66_)[default__.safeIndex(459, (d_177_v66_).length(0))]).set(default__.safeIndex(d_92_v0_, len((d_177_v66_)[default__.safeIndex(459, (d_177_v66_).length(0))])), d_92_v0_)])
                rhs9_ = (((d_93_v1_) + (_dafny.SeqWithoutIsStrInference([((d_177_v66_)[default__.safeIndex(459, (d_177_v66_).length(0))])[default__.safeIndex(d_92_v0_, len((d_177_v66_)[default__.safeIndex(459, (d_177_v66_).length(0))]))]]))) + ((d_184_v70_)[default__.safeIndex(d_92_v0_, len(d_184_v70_))])).set(default__.safeIndex(d_92_v0_, len(((d_93_v1_) + (_dafny.SeqWithoutIsStrInference([((d_177_v66_)[default__.safeIndex(459, (d_177_v66_).length(0))])[default__.safeIndex(d_92_v0_, len((d_177_v66_)[default__.safeIndex(459, (d_177_v66_).length(0))]))]]))) + ((d_184_v70_)[default__.safeIndex(d_92_v0_, len(d_184_v70_))]))), d_92_v0_)
                rhs10_ = d_183_v69_
                rhs11_ = (d_92_v0_) + (d_92_v0_)
                lhs6_ = d_100_globalState_
                d_93_v1_ = rhs9_
                d_183_v69_ = rhs10_
                lhs6_.f3 = rhs11_
            if (d_94_v2_ if d_94_v2_ else d_94_v2_):
                d_185_v71_: str
                d_185_v71_ = _dafny.CodePoint('r')
                d_186_v72_: D10
                d_186_v72_ = D10_DC24(((d_99_v7_)[(0) - (d_92_v0_)] if ((0) - (d_92_v0_)) in (d_99_v7_) else default__.fm6(68, d_100_globalState_)), d_94_v2_, ((d_97_v5_) + (_dafny.SeqWithoutIsStrInference([d_185_v71_]))).set(default__.safeIndex(len(default__.fm7(d_185_v71_, d_100_globalState_)), len((d_97_v5_) + (_dafny.SeqWithoutIsStrInference([d_185_v71_])))), d_185_v71_))
                pat_let_tv0_ = d_94_v2_
                pat_let_tv1_ = d_92_v0_
                pat_let_tv2_ = d_92_v0_
                pat_let_tv3_ = d_94_v2_
                pat_let_tv4_ = d_92_v0_
                def iife56_(_pat_let2_0):
                    def iife57_(d_187_dt__update__tmp_h0_):
                        def iife58_(_pat_let3_0):
                            def iife59_(d_188_dt__update_hcf32_h0_):
                                def iife60_(_pat_let4_0):
                                    def iife61_(d_189_dt__update_hcf31_h0_):
                                        return D10_DC24(d_189_dt__update_hcf31_h0_, d_188_dt__update_hcf32_h0_, (d_187_dt__update__tmp_h0_).cf33)
                                    return iife61_(_pat_let4_0)
                                return iife60_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv1_, pat_let_tv2_])))
                            return iife59_(_pat_let3_0)
                        return iife58_(pat_let_tv0_)
                    return iife57_(_pat_let2_0)
                def iife55_(_pat_let1_0):
                    def iife62_(d_190_dt__update__tmp_h1_):
                        def iife63_(_pat_let5_0):
                            def iife64_(d_191_dt__update_hcf32_h1_):
                                return D10_DC24((d_190_dt__update__tmp_h1_).cf31, d_191_dt__update_hcf32_h1_, (d_190_dt__update__tmp_h1_).cf33)
                            return iife64_(_pat_let5_0)
                        return iife63_(pat_let_tv3_)
                    return iife62_(_pat_let1_0)
                def iife54_(_pat_let0_0):
                    def iife65_(d_192_dt__update__tmp_h2_):
                        def iife66_(_pat_let6_0):
                            def iife67_(d_193_dt__update_hcf31_h1_):
                                return D10_DC24(d_193_dt__update_hcf31_h1_, (d_192_dt__update__tmp_h2_).cf32, (d_192_dt__update__tmp_h2_).cf33)
                            return iife67_(_pat_let6_0)
                        return iife66_(pat_let_tv4_)
                    return iife65_(_pat_let0_0)
                d_186_v72_ = iife54_(iife55_(iife56_(d_186_v72_)))
                index23_ = default__.safeIndex(657, (d_101_v8_).length(0))
                (d_101_v8_)[index23_] = d_92_v0_
                index24_ = default__.safeIndex(657, (d_101_v8_).length(0))
                (d_101_v8_)[index24_] = default__.safeDivisionInt(d_92_v0_, d_92_v0_)
                d_194_v73_: _dafny.Array
                nw30_ = _dafny.Array(D2.default()(), 5)
                d_194_v73_ = nw30_
                d_195_v74_: _dafny.Array
                nw31_ = _dafny.Array(False, 25)
                d_195_v74_ = nw31_
                d_196_v75_: _dafny.Seq
                d_196_v75_ = _dafny.SeqWithoutIsStrInference([d_195_v74_])
                d_197_v76_: D2
                d_197_v76_ = D2_DC5((d_196_v75_)[default__.safeIndex((d_101_v8_)[default__.safeIndex(657, (d_101_v8_).length(0))], len(d_196_v75_))])
                index25_ = default__.safeIndex(747, (d_194_v73_).length(0))
                (d_194_v73_)[index25_] = d_197_v76_
                d_198_v77_: C1
                nw32_ = C1()
                nw32_.ctor__()
                d_198_v77_ = nw32_
                d_199_v78_: _dafny.Array
                nw33_ = _dafny.Array(None, 4)
                nw33_[int(0)] = (d_198_v77_ if d_94_v2_ else d_198_v77_)
                nw33_[int(1)] = d_198_v77_
                nw33_[int(2)] = d_198_v77_
                nw33_[int(3)] = d_198_v77_
                d_199_v78_ = nw33_
                index26_ = default__.safeIndex(751, (d_199_v78_).length(0))
                (d_199_v78_)[index26_] = d_198_v77_
                d_200_v79_: _dafny.MultiSet
                d_200_v79_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_185_v71_ for d_201_i8_ in range(default__.abs(665))]), d_97_v5_])
                d_202_v80_: C7
                nw34_ = C7()
                nw34_.ctor__()
                d_202_v80_ = nw34_
                d_203_v81_: _dafny.Map
                d_203_v81_ = _dafny.Map({d_92_v0_: d_202_v80_})
                index27_ = default__.safeIndex(747, (d_194_v73_).length(0))
                index28_ = default__.safeIndex(751, (d_199_v78_).length(0))
                rhs12_ = d_197_v76_
                rhs13_ = d_198_v77_
                rhs14_ = d_94_v2_
                rhs15_ = d_94_v2_
                rhs16_ = (((d_200_v79_)[d_97_v5_] if (d_97_v5_) in (d_200_v79_) else d_92_v0_)) - (default__.safeModuloInt(len(d_203_v81_), d_92_v0_))
                lhs7_ = d_194_v73_
                lhs8_ = default__.safeIndex(747, (d_194_v73_).length(0))
                lhs9_ = d_199_v78_
                lhs10_ = default__.safeIndex(751, (d_199_v78_).length(0))
                lhs11_ = d_100_globalState_
                lhs12_ = d_100_globalState_
                lhs7_[lhs8_] = rhs12_
                lhs9_[lhs10_] = rhs13_
                d_94_v2_ = rhs14_
                lhs11_.f4 = rhs15_
                lhs12_.f8 = rhs16_
                d_204_v82_: _dafny.Map
                d_204_v82_ = _dafny.Map({d_94_v2_: d_171_v60_})
                d_205_v83_: D0
                d_205_v83_ = D0_DC2()
                d_206_v84_: _dafny.Map
                d_206_v84_ = _dafny.Map({(d_204_v82_) | (d_204_v82_): d_205_v83_})
                index29_ = default__.safeIndex(657, (d_101_v8_).length(0))
                rhs17_ = d_94_v2_
                rhs18_ = default__.fm49(d_94_v2_, d_100_globalState_)
                rhs19_ = default__.safeDivisionInt((d_101_v8_)[default__.safeIndex(657, (d_101_v8_).length(0))], default__.safeDivisionInt(104, d_92_v0_))
                rhs20_ = not(d_94_v2_)
                rhs21_ = len(d_206_v84_)
                lhs13_ = d_100_globalState_
                lhs14_ = d_101_v8_
                lhs15_ = default__.safeIndex(657, (d_101_v8_).length(0))
                lhs16_ = d_100_globalState_
                lhs17_ = d_100_globalState_
                lhs13_.f4 = rhs17_
                d_94_v2_ = rhs18_
                lhs14_[lhs15_] = rhs19_
                lhs16_.f4 = rhs20_
                lhs17_.f3 = rhs21_
                d_207_v85_: _dafny.MultiSet
                d_207_v85_ = _dafny.MultiSet([default__.fm49(d_94_v2_, d_100_globalState_)])
                d_208_v86_: _dafny.Seq
                d_208_v86_ = _dafny.SeqWithoutIsStrInference([d_97_v5_])
                d_209_v87_: _dafny.Map
                d_209_v87_ = _dafny.Map({((d_207_v85_)[d_94_v2_] if (d_94_v2_) in (d_207_v85_) else d_92_v0_): (-212) * (len((d_208_v86_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjqbikwfu"))), len(d_208_v86_))]))})
                rhs22_ = d_94_v2_
                rhs23_ = ((d_209_v87_ if d_94_v2_ else d_209_v87_)).set(d_92_v0_, d_92_v0_)
                lhs18_ = d_100_globalState_
                lhs18_.f4 = rhs22_
                d_209_v87_ = rhs23_
            elif True:
                d_210_v88_: _dafny.Array
                d_210_v88_ = d_101_v8_
                d_211_v89_: _dafny.Seq
                d_211_v89_ = _dafny.SeqWithoutIsStrInference([d_101_v8_])
                d_212_v90_: _dafny.MultiSet
                d_212_v90_ = _dafny.MultiSet([d_94_v2_])
                d_213_v91_: _dafny.Array
                nw35_ = _dafny.Array(None, 8)
                nw35_[int(0)] = (d_101_v8_ if d_94_v2_ else d_101_v8_)
                nw35_[int(1)] = (d_210_v88_)
                nw35_[int(2)] = d_101_v8_
                nw35_[int(3)] = d_101_v8_
                nw35_[int(4)] = d_101_v8_
                nw35_[int(5)] = d_101_v8_
                nw35_[int(6)] = d_101_v8_
                nw35_[int(7)] = (d_211_v89_)[default__.safeIndex((d_212_v90_).cardinality, len(d_211_v89_))]
                d_213_v91_ = nw35_
                index30_ = default__.safeIndex(630, (d_213_v91_).length(0))
                (d_213_v91_)[index30_] = d_101_v8_
                index31_ = default__.safeIndex(630, (d_213_v91_).length(0))
                nw36_ = _dafny.Array(int(0), 25)
                (d_213_v91_)[index31_] = nw36_
                d_214_v93_: _dafny.Map
                def iife68_():
                    coll54_ = _dafny.Set()
                    compr_54_: int
                    for compr_54_ in (d_132_v24_).Elements:
                        d_215_v92_: int = compr_54_
                        if (d_215_v92_) in (d_132_v24_):
                            coll54_ = coll54_.union(_dafny.Set([(d_215_v92_) + (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([-834])})))]))
                    return _dafny.Set(coll54_)
                d_214_v93_ = _dafny.Map({(_dafny.Set({d_92_v0_, d_92_v0_, (0) - (len(iife68_()
                )), 605})) | (d_132_v24_): d_94_v2_})
                d_214_v93_ = (d_214_v93_) | (d_214_v93_)
                d_216_v94_: _dafny.Seq
                d_216_v94_ = _dafny.SeqWithoutIsStrInference([d_94_v2_, d_94_v2_])
                (d_100_globalState_).f8 = (d_92_v0_) - (((d_212_v90_)[(d_216_v94_)[default__.safeIndex(d_92_v0_, len(d_216_v94_))]] if ((d_216_v94_)[default__.safeIndex(d_92_v0_, len(d_216_v94_))]) in (d_212_v90_) else d_92_v0_))
                d_217_v95_: C9
                nw37_ = C9()
                nw37_.ctor__((False) and (d_94_v2_))
                d_217_v95_ = nw37_
                d_218_v96_: C8
                nw38_ = C8()
                nw38_.ctor__()
                d_218_v96_ = nw38_
                d_218_v96_ = d_218_v96_
        d_219_v97_: C0
        nw39_ = C0()
        nw39_.ctor__()
        d_219_v97_ = nw39_
        d_220_v98_: _dafny.Map
        d_220_v98_ = _dafny.Map({d_92_v0_: d_92_v0_})
        d_221_v99_: D3
        d_221_v99_ = D3_DC8(d_220_v98_)
        d_222_v100_: _dafny.Seq
        d_222_v100_ = _dafny.SeqWithoutIsStrInference([d_94_v2_])
        d_223_v101_: _dafny.Set
        d_223_v101_ = _dafny.Set({d_94_v2_})
        d_224_i9_: int
        d_224_i9_ = 0
        with _dafny.label("1"):
            while default__.fm49((d_223_v101_).issubset(default__.fm35(d_94_v2_, default__.fm26(d_221_v99_, d_92_v0_, _dafny.MultiSet([d_222_v100_, d_222_v100_, d_222_v100_]), d_100_globalState_), d_94_v2_, True, d_100_globalState_)), d_100_globalState_):
                with _dafny.c_label("1"):
                    if (d_224_i9_) >= (100):
                        raise _dafny.Break("1")
                    d_224_i9_ = (d_224_i9_) + (1)
                    d_225_v102_: _dafny.Array
                    nw40_ = _dafny.Array(_dafny.Map({}), 23)
                    d_225_v102_ = nw40_
                    d_226_v103_: C4
                    nw41_ = C4()
                    nw41_.ctor__(d_225_v102_)
                    d_226_v103_ = nw41_
                    d_227_v104_: _dafny.Seq
                    d_227_v104_ = _dafny.SeqWithoutIsStrInference([d_226_v103_])
                    d_226_v103_ = (d_227_v104_)[default__.safeIndex(208, len(d_227_v104_))]
                    (d_100_globalState_).f8 = (default__.safeDivisionInt(-265, d_92_v0_) if not(d_94_v2_) else (d_92_v0_) + (d_92_v0_))
                    d_228_v105_: T0
                    nw42_ = C7()
                    nw42_.ctor__()
                    d_228_v105_ = nw42_
                    d_229_v106_: str
                    d_229_v106_ = _dafny.CodePoint('w')
                    d_230_v107_: _dafny.Set
                    d_230_v107_ = _dafny.Set({_dafny.Map({d_228_v105_: d_229_v106_})})
                    d_231_v108_: _dafny.Map
                    d_231_v108_ = _dafny.Map({d_228_v105_: _dafny.CodePoint('q')})
                    d_230_v107_ = _dafny.Set({d_231_v108_, d_231_v108_, _dafny.Map({d_228_v105_: _dafny.CodePoint('y')}), _dafny.Map({d_228_v105_: d_229_v106_})})
                    (d_226_v103_).m9(default__.fm6(d_92_v0_, d_100_globalState_), d_92_v0_, d_100_globalState_)
                    pass
            pass
        d_232_v109_: D3
        d_232_v109_ = D3_DC9()
        d_232_v109_ = d_232_v109_
        d_233_v110_: D21
        d_233_v110_ = D21_DC43(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: d_94_v2_})) for d_234_i11_ in range(default__.abs(678))]), False)
        d_235_i10_: int
        d_235_i10_ = 0
        with _dafny.label("2"):
            pat_let_tv5_ = d_92_v0_
            pat_let_tv6_ = d_92_v0_
            pat_let_tv7_ = d_94_v2_
            pat_let_tv8_ = d_94_v2_
            pat_let_tv9_ = d_100_globalState_
            def lambda13_(source9_):
                if source9_.is_DC43:
                    d_285___mcc_h11_ = source9_.cf63
                    d_286___mcc_h12_ = source9_.cf64
                    d_287_cf64_ = d_286___mcc_h12_
                    d_288_cf63_ = d_285___mcc_h11_
                    return (pat_let_tv5_) >= (pat_let_tv6_)
                elif True:
                    d_289___mcc_h13_ = source9_.cf62
                    d_290_cf62_ = d_289___mcc_h13_
                    return pat_let_tv7_

            def iife70_(_pat_let7_0):
                def iife71_(d_291_dt__update__tmp_h3_):
                    def iife72_(_pat_let8_0):
                        def iife73_(d_292_dt__update_hcf64_h0_):
                            return D21_DC43((d_291_dt__update__tmp_h3_).cf63, d_292_dt__update_hcf64_h0_)
                        return iife73_(_pat_let8_0)
                    return iife72_(default__.fm49(pat_let_tv8_, pat_let_tv9_))
                return iife71_(_pat_let7_0)
            while lambda13_(iife70_(d_233_v110_)):
                with _dafny.c_label("2"):
                    if (d_235_i10_) >= (100):
                        raise _dafny.Break("2")
                    d_235_i10_ = (d_235_i10_) + (1)
                    if d_94_v2_:
                        d_236_v111_: _dafny.Array
                        nw43_ = _dafny.Array(_dafny.Map({}), 27)
                        d_236_v111_ = nw43_
                        d_237_v112_: C4
                        nw44_ = C4()
                        nw44_.ctor__(d_236_v111_)
                        d_237_v112_ = nw44_
                        (d_100_globalState_).f8 = (d_92_v0_ if d_94_v2_ else (0) - (d_92_v0_))
                        d_238_v113_: _dafny.Map
                        d_238_v113_ = _dafny.Map({d_92_v0_: (d_237_v112_).fm3(d_93_v1_, _dafny.MultiSet([d_92_v0_]), not(d_94_v2_), d_100_globalState_)})
                        d_239_v114_: _dafny.Set
                        d_239_v114_ = _dafny.Set({d_238_v113_, d_238_v113_, d_238_v113_, (d_238_v113_) | (d_238_v113_), d_238_v113_})
                        d_240_v115_: T0
                        nw45_ = C2()
                        nw45_.ctor__()
                        d_240_v115_ = nw45_
                        d_241_v116_: _dafny.Seq
                        d_241_v116_ = _dafny.SeqWithoutIsStrInference([d_240_v115_, d_240_v115_])
                        d_242_v117_: str
                        d_242_v117_ = _dafny.CodePoint('y')
                        d_243_v118_: C1
                        nw46_ = C1()
                        nw46_.ctor__()
                        d_243_v118_ = nw46_
                        d_244_v119_: _dafny.Map
                        d_244_v119_ = _dafny.Map({d_243_v118_: 110})
                        rhs24_ = (d_97_v5_) <= (default__.fm25(d_100_globalState_))
                        rhs25_ = (_dafny.Set({d_238_v113_})).intersection(default__.fm53(d_92_v0_, d_92_v0_, d_100_globalState_))
                        rhs26_ = (d_241_v116_)[default__.safeIndex(d_92_v0_, len(d_241_v116_))]
                        rhs27_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukahennov"))).set(default__.safeIndex(d_92_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukahennov")))), d_242_v117_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvcx"))) + (d_97_v5_))).set(default__.safeIndex(((d_244_v119_)[d_243_v118_] if (d_243_v118_) in (d_244_v119_) else (0) - (d_92_v0_)), len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukahennov"))).set(default__.safeIndex(d_92_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukahennov")))), d_242_v117_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvcx"))) + (d_97_v5_)))), d_242_v117_)
                        d_94_v2_ = rhs24_
                        d_239_v114_ = rhs25_
                        d_240_v115_ = rhs26_
                        d_97_v5_ = rhs27_
                        d_245_v120_: _dafny.Map
                        d_245_v120_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gofjuqlqk")): d_94_v2_})
                        (d_100_globalState_).f3 = len(d_245_v120_)
                        d_92_v0_ = -256
                    elif True:
                        d_246_v121_: _dafny.Array
                        nw47_ = _dafny.Array(False, 10)
                        d_246_v121_ = nw47_
                        index32_ = default__.safeIndex(429, (d_246_v121_).length(0))
                        (d_246_v121_)[index32_] = d_94_v2_
                        index33_ = default__.safeIndex(429, (d_246_v121_).length(0))
                        (d_246_v121_)[index33_] = d_94_v2_
                        d_247_v122_: int
                        d_248_v123_: _dafny.Map
                        d_249_v124_: bool
                        d_250_v125_: C0
                        out8_: int
                        out9_: _dafny.Map
                        out10_: bool
                        out11_: C0
                        out8_, out9_, out10_, out11_ = default__.m15(d_100_globalState_)
                        d_247_v122_ = out8_
                        d_248_v123_ = out9_
                        d_249_v124_ = out10_
                        d_250_v125_ = out11_
                        d_251_v126_: _dafny.Map
                        d_251_v126_ = _dafny.Map({d_249_v124_: (d_246_v121_)[default__.safeIndex(429, (d_246_v121_).length(0))]})
                        (d_100_globalState_).f4 = (d_251_v126_) != (d_251_v126_)
                        d_252_v127_: int
                        d_253_v128_: _dafny.Map
                        d_254_v129_: bool
                        d_255_v130_: C0
                        out12_: int
                        out13_: _dafny.Map
                        out14_: bool
                        out15_: C0
                        out12_, out13_, out14_, out15_ = default__.m15(d_100_globalState_)
                        d_252_v127_ = out12_
                        d_253_v128_ = out13_
                        d_254_v129_ = out14_
                        d_255_v130_ = out15_
                        index34_ = default__.safeIndex(49, (d_101_v8_).length(0))
                        (d_101_v8_)[index34_] = d_247_v122_
                        index35_ = default__.safeIndex(49, (d_101_v8_).length(0))
                        (d_101_v8_)[index35_] = d_252_v127_
                    rhs28_ = -265
                    rhs29_ = (0) - (d_92_v0_)
                    lhs19_ = d_100_globalState_
                    lhs20_ = d_100_globalState_
                    lhs19_.f8 = rhs28_
                    lhs20_.f8 = rhs29_
                    d_256_v131_: T0
                    nw48_ = C7()
                    nw48_.ctor__()
                    d_256_v131_ = nw48_
                    d_257_v132_: _dafny.Seq
                    d_257_v132_ = _dafny.SeqWithoutIsStrInference([d_256_v131_])
                    d_258_v133_: _dafny.Map
                    d_258_v133_ = _dafny.Map({d_94_v2_: d_94_v2_})
                    d_259_v134_: _dafny.MultiSet
                    d_259_v134_ = _dafny.MultiSet([d_94_v2_])
                    d_260_v135_: _dafny.Array
                    nw49_ = _dafny.Array(None, 29)
                    nw49_[int(0)] = (d_132_v24_).isdisjoint(d_132_v24_)
                    nw49_[int(1)] = (d_222_v100_) == (_dafny.SeqWithoutIsStrInference([d_94_v2_, d_94_v2_, d_94_v2_, d_94_v2_, d_94_v2_]))
                    nw49_[int(2)] = d_94_v2_
                    nw49_[int(3)] = ((d_258_v133_)[d_94_v2_] if (d_94_v2_) in (d_258_v133_) else d_94_v2_)
                    nw49_[int(4)] = d_94_v2_
                    nw49_[int(5)] = d_94_v2_
                    nw49_[int(6)] = (d_222_v100_) == (d_222_v100_)
                    nw49_[int(7)] = False
                    nw49_[int(8)] = d_94_v2_
                    nw49_[int(9)] = True
                    nw49_[int(10)] = (d_94_v2_ if d_94_v2_ else d_94_v2_)
                    nw49_[int(11)] = d_94_v2_
                    nw49_[int(12)] = d_94_v2_
                    nw49_[int(13)] = d_94_v2_
                    nw49_[int(14)] = (d_256_v131_).fm3(_dafny.SeqWithoutIsStrInference([len(d_97_v5_)]), d_99_v7_, d_94_v2_, d_100_globalState_)
                    nw49_[int(15)] = d_94_v2_
                    nw49_[int(16)] = d_94_v2_
                    nw49_[int(17)] = d_94_v2_
                    nw49_[int(18)] = d_94_v2_
                    nw49_[int(19)] = d_94_v2_
                    nw49_[int(20)] = d_94_v2_
                    nw49_[int(21)] = d_94_v2_
                    nw49_[int(22)] = not (d_94_v2_) or (d_94_v2_)
                    nw49_[int(23)] = d_94_v2_
                    nw49_[int(24)] = d_94_v2_
                    nw49_[int(25)] = False
                    nw49_[int(26)] = (_dafny.MultiSet([d_94_v2_])).isdisjoint(d_259_v134_)
                    nw49_[int(27)] = d_94_v2_
                    nw49_[int(28)] = d_94_v2_
                    d_260_v135_ = nw49_
                    index36_ = default__.safeIndex(299, (d_260_v135_).length(0))
                    (d_260_v135_)[index36_] = not(not (d_94_v2_) or (d_94_v2_))
                    d_261_v137_: _dafny.Array
                    def lambda11_(d_262_v2_):
                        def lambda12_(d_263_i12_):
                            def iife69_():
                                coll55_ = _dafny.Map()
                                compr_55_: str
                                for compr_55_ in (_dafny.MultiSet([_dafny.CodePoint('w')])).Elements:
                                    d_264_v136_: str = compr_55_
                                    if (d_264_v136_) in (_dafny.MultiSet([_dafny.CodePoint('w')])):
                                        coll55_[d_264_v136_] = d_262_v2_
                                return _dafny.Map(coll55_)
                            return iife69_()
                            

                        return lambda12_

                    init6_ = lambda11_(d_94_v2_)
                    nw50_ = _dafny.Array(None, 4)
                    for i0_6_ in range(nw50_.length(0)):
                        nw50_[i0_6_] = init6_(i0_6_)
                    d_261_v137_ = nw50_
                    d_265_v138_: str
                    d_265_v138_ = _dafny.CodePoint('q')
                    d_266_v139_: _dafny.Map
                    d_266_v139_ = _dafny.Map({d_265_v138_: d_94_v2_})
                    index37_ = default__.safeIndex(625, (d_261_v137_).length(0))
                    (d_261_v137_)[index37_] = (d_266_v139_).set(d_265_v138_, d_94_v2_)
                    d_267_v140_: D27
                    d_267_v140_ = D27_DC58(d_266_v139_)
                    index38_ = default__.safeIndex(299, (d_260_v135_).length(0))
                    index39_ = default__.safeIndex(625, (d_261_v137_).length(0))
                    rhs30_ = d_92_v0_
                    rhs31_ = _dafny.SeqWithoutIsStrInference([(d_256_v131_ if d_94_v2_ else d_256_v131_), d_256_v131_, d_256_v131_, d_256_v131_, d_256_v131_])
                    rhs32_ = False
                    rhs33_ = (d_256_v131_).fm3(d_93_v1_, d_99_v7_, d_94_v2_, d_100_globalState_)
                    rhs34_ = (d_267_v140_).cf82
                    lhs21_ = d_100_globalState_
                    lhs22_ = d_260_v135_
                    lhs23_ = default__.safeIndex(299, (d_260_v135_).length(0))
                    lhs24_ = d_261_v137_
                    lhs25_ = default__.safeIndex(625, (d_261_v137_).length(0))
                    lhs21_.f8 = rhs30_
                    d_257_v132_ = rhs31_
                    d_94_v2_ = rhs32_
                    lhs22_[lhs23_] = rhs33_
                    lhs24_[lhs25_] = rhs34_
                    d_268_v141_: D2
                    d_268_v141_ = D2_DC5(d_260_v135_)
                    source8_ = d_268_v141_
                    if source8_.is_DC6:
                        d_269___mcc_h4_ = source8_.cf6
                        d_270___mcc_h5_ = source8_.cf7
                        d_271___mcc_h6_ = source8_.cf8
                        d_272_cf8_ = d_271___mcc_h6_
                        d_273_cf7_ = d_270___mcc_h5_
                        d_274_cf6_ = d_269___mcc_h4_
                        d_95_v3_ = (d_95_v3_).set(d_273_cf7_, d_274_cf6_)
                        (d_100_globalState_).f8 = default__.safeDivisionInt(d_92_v0_, d_92_v0_)
                        d_275_v142_: _dafny.Array
                        nw51_ = _dafny.Array(_dafny.Array(None, 0), 25)
                        d_275_v142_ = nw51_
                        d_275_v142_ = d_275_v142_
                        index40_ = default__.safeIndex(851, (d_101_v8_).length(0))
                        (d_101_v8_)[index40_] = d_92_v0_
                        index41_ = default__.safeIndex(851, (d_101_v8_).length(0))
                        (d_101_v8_)[index41_] = d_274_cf6_
                    elif source8_.is_DC7:
                        d_276___mcc_h7_ = source8_.cf9
                        d_277___mcc_h8_ = source8_.cf10
                        d_278___mcc_h9_ = source8_.cf11
                        d_279_cf11_ = d_278___mcc_h9_
                        d_280_cf10_ = d_277___mcc_h8_
                        d_281_cf9_ = d_276___mcc_h7_
                        d_280_cf10_ = (d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))]
                        (d_100_globalState_).f4 = (d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))]
                        d_258_v133_ = (d_258_v133_).set(not ((d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))]) or (d_280_cf10_), (d_92_v0_) >= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmy")))))
                        (d_100_globalState_).f8 = (0) - (((d_95_v3_)[d_280_cf10_] if (d_280_cf10_) in (d_95_v3_) else d_92_v0_))
                    elif True:
                        d_282___mcc_h10_ = source8_.cf5
                        d_283_cf5_ = d_282___mcc_h10_
                        index42_ = default__.safeIndex(299, (d_260_v135_).length(0))
                        (d_260_v135_)[index42_] = ((d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))] if (d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))] else (d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))])
                        d_101_v8_ = d_101_v8_
                        d_284_v143_: D27
                        d_284_v143_ = D27_DC59(d_94_v2_, d_260_v135_, (d_260_v135_)[default__.safeIndex(299, (d_260_v135_).length(0))])
                        (d_100_globalState_).f4 = (d_284_v143_).cf83
                        (d_100_globalState_).f3 = default__.safeDivisionInt(d_92_v0_, d_92_v0_)
                    pass
            pass
        d_293_v144_: _dafny.Array
        def lambda14_(d_294_v2_):
            def lambda15_(d_295_i13_):
                return d_294_v2_

            return lambda15_

        init7_ = lambda14_(d_94_v2_)
        nw52_ = _dafny.Array(None, 12)
        for i0_7_ in range(nw52_.length(0)):
            nw52_[i0_7_] = init7_(i0_7_)
        d_293_v144_ = nw52_
        index43_ = default__.safeIndex(300, (d_293_v144_).length(0))
        (d_293_v144_)[index43_] = d_94_v2_
        index44_ = default__.safeIndex(300, (d_293_v144_).length(0))
        (d_293_v144_)[index44_] = d_94_v2_
        hi1_ = (d_92_v0_ if d_94_v2_ else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_296_i15_ in range(default__.abs(250))])))
        for d_297_i14_ in range(d_92_v0_, hi1_):
            d_298_v145_: _dafny.Seq
            d_298_v145_ = _dafny.SeqWithoutIsStrInference([d_293_v144_])
            (d_100_globalState_).f3 = len(_dafny.SeqWithoutIsStrInference([(d_298_v145_)[default__.safeIndex(d_92_v0_, len(d_298_v145_))], d_293_v144_]))
            d_299_v146_: str
            d_299_v146_ = _dafny.CodePoint('y')
            d_299_v146_ = _dafny.CodePoint('l')
            d_99_v7_ = d_99_v7_
            if (not(False)) not in (default__.fm23((d_97_v5_).set(default__.safeIndex(d_92_v0_, len(d_97_v5_)), d_299_v146_), d_100_globalState_)):
                d_300_v147_: _dafny.Array
                nw53_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_300_v147_ = nw53_
                index45_ = default__.safeIndex(431, (d_300_v147_).length(0))
                (d_300_v147_)[index45_] = d_101_v8_
                index46_ = default__.safeIndex(431, (d_300_v147_).length(0))
                (d_300_v147_)[index46_] = d_101_v8_
                d_301_v148_: C5
                nw54_ = C5()
                nw54_.ctor__()
                d_301_v148_ = nw54_
                d_302_v149_: _dafny.Array
                nw55_ = _dafny.Array(None, 1)
                nw55_[int(0)] = d_301_v148_
                d_302_v149_ = nw55_
                index47_ = default__.safeIndex(995, (d_302_v149_).length(0))
                (d_302_v149_)[index47_] = (d_301_v148_ if (d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))] else d_301_v148_)
                index48_ = default__.safeIndex(995, (d_302_v149_).length(0))
                (d_302_v149_)[index48_] = d_301_v148_
                (d_100_globalState_).f4 = (d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))]
                d_303_v150_: _dafny.MultiSet
                d_303_v150_ = _dafny.MultiSet([d_94_v2_])
                index49_ = default__.safeIndex(300, (d_293_v144_).length(0))
                rhs35_ = (d_92_v0_ if not((d_297_i14_) < (d_297_i14_)) else d_297_i14_)
                rhs36_ = ((_dafny.MultiSet(d_222_v100_)) - (d_303_v150_)).ispropersubset((d_303_v150_).set((d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))], default__.abs(d_297_i14_)))
                rhs37_ = (d_92_v0_) - (d_92_v0_)
                rhs38_ = ((d_222_v100_) + (d_222_v100_)) + (d_222_v100_)
                rhs39_ = len((d_222_v100_).set(default__.safeIndex((0) - (d_92_v0_), len(d_222_v100_)), d_94_v2_))
                lhs26_ = d_100_globalState_
                lhs27_ = d_293_v144_
                lhs28_ = default__.safeIndex(300, (d_293_v144_).length(0))
                lhs29_ = d_100_globalState_
                lhs30_ = d_100_globalState_
                lhs26_.f3 = rhs35_
                lhs27_[lhs28_] = rhs36_
                lhs29_.f3 = rhs37_
                d_222_v100_ = rhs38_
                lhs30_.f3 = rhs39_
                index50_ = default__.safeIndex(300, (d_293_v144_).length(0))
                (d_293_v144_)[index50_] = (d_303_v150_).isdisjoint(d_303_v150_)
            elif True:
                d_304_v151_: C8
                nw56_ = C8()
                nw56_.ctor__()
                d_304_v151_ = nw56_
                d_305_v152_: _dafny.Seq
                d_305_v152_ = _dafny.SeqWithoutIsStrInference([d_304_v151_, d_304_v151_])
                d_304_v151_ = (d_305_v152_)[default__.safeIndex(d_297_i14_, len(d_305_v152_))]
                (d_100_globalState_).f3 = d_92_v0_
                index51_ = default__.safeIndex(300, (d_293_v144_).length(0))
                (d_293_v144_)[index51_] = default__.fm49(d_94_v2_, d_100_globalState_)
                d_306_v153_: C3
                nw57_ = C3()
                nw57_.ctor__()
                d_306_v153_ = nw57_
                d_306_v153_ = d_306_v153_
                (d_100_globalState_).f8 = d_92_v0_
        d_307_v154_: C8
        nw58_ = C8()
        nw58_.ctor__()
        d_307_v154_ = nw58_
        d_308_v155_: _dafny.Array
        nw59_ = _dafny.Array(None, 3)
        nw59_[int(0)] = d_307_v154_
        nw59_[int(1)] = d_307_v154_
        nw59_[int(2)] = d_307_v154_
        d_308_v155_ = nw59_
        d_309_v156_: _dafny.Array
        nw60_ = _dafny.Array(None, 17)
        nw60_[int(0)] = d_308_v155_
        nw60_[int(1)] = d_308_v155_
        nw60_[int(2)] = d_308_v155_
        nw60_[int(3)] = d_308_v155_
        nw60_[int(4)] = d_308_v155_
        nw60_[int(5)] = d_308_v155_
        nw60_[int(6)] = d_308_v155_
        nw60_[int(7)] = d_308_v155_
        nw60_[int(8)] = d_308_v155_
        nw60_[int(9)] = d_308_v155_
        nw60_[int(10)] = d_308_v155_
        nw60_[int(11)] = d_308_v155_
        nw60_[int(12)] = d_308_v155_
        nw60_[int(13)] = d_308_v155_
        nw60_[int(14)] = d_308_v155_
        nw60_[int(15)] = d_308_v155_
        nw60_[int(16)] = d_308_v155_
        d_309_v156_ = nw60_
        index52_ = default__.safeIndex(845, (d_309_v156_).length(0))
        (d_309_v156_)[index52_] = d_308_v155_
        d_310_v157_: _dafny.Map
        d_310_v157_ = _dafny.Map({(d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))]: d_222_v100_})
        d_311_v158_: _dafny.MultiSet
        d_311_v158_ = _dafny.MultiSet([((d_310_v157_)[d_94_v2_] if (d_94_v2_) in (d_310_v157_) else d_222_v100_)])
        index53_ = default__.safeIndex(845, (d_309_v156_).length(0))
        rhs40_ = False
        rhs41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmhjp"))
        rhs42_ = default__.fm26(d_221_v99_, len(d_222_v100_), d_311_v158_, d_100_globalState_)
        rhs43_ = d_308_v155_
        lhs31_ = d_100_globalState_
        lhs32_ = d_100_globalState_
        lhs33_ = d_309_v156_
        lhs34_ = default__.safeIndex(845, (d_309_v156_).length(0))
        lhs31_.f4 = rhs40_
        d_97_v5_ = rhs41_
        lhs32_.f3 = rhs42_
        lhs33_[lhs34_] = rhs43_
        (d_100_globalState_).f4 = not(d_94_v2_)
        (d_100_globalState_).f3 = (d_92_v0_) * (d_92_v0_)
        (d_100_globalState_).f4 = True
        d_312_v159_: D2
        d_312_v159_ = D2_DC6(d_92_v0_, (d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))], (d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))])
        hi2_ = (d_92_v0_) - ((d_312_v159_).cf6)
        for d_313_i16_ in range(len(d_222_v100_), hi2_):
            if d_94_v2_:
                d_314_v160_: str
                d_314_v160_ = _dafny.CodePoint('x')
                d_315_v161_: _dafny.Map
                d_315_v161_ = _dafny.Map({d_314_v160_: d_97_v5_})
                d_316_v162_: _dafny.Map
                d_316_v162_ = _dafny.Map({((d_315_v161_)[d_314_v160_] if (d_314_v160_) in (d_315_v161_) else _dafny.SeqWithoutIsStrInference([(D16_DC35(d_92_v0_, (d_97_v5_)[default__.safeIndex(d_92_v0_, len(d_97_v5_))], len(_dafny.Set({d_94_v2_, d_94_v2_})), d_313_i16_, d_92_v0_)).cf51 for d_317_i17_ in range(default__.abs(-780))])): d_94_v2_})
                d_318_v163_: _dafny.Seq
                d_318_v163_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_92_v0_})])
                d_316_v162_ = default__.fm42(d_313_i16_, 198, len((d_318_v163_)[default__.safeIndex(d_313_i16_, len(d_318_v163_))]), True, d_100_globalState_)
                d_319_v164_: _dafny.Set
                d_319_v164_ = _dafny.Set({d_220_v98_, (d_220_v98_) | (d_220_v98_)})
                d_320_v165_: _dafny.Map
                d_320_v165_ = _dafny.Map({(d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))]: _dafny.Set({d_220_v98_})})
                d_319_v164_ = ((d_320_v165_)[not (d_94_v2_) or (True)] if (not (d_94_v2_) or (True)) in (d_320_v165_) else d_319_v164_)
                (d_100_globalState_).f3 = 359
                index54_ = default__.safeIndex(300, (d_293_v144_).length(0))
                (d_293_v144_)[index54_] = (d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))]
                d_321_v166_: _dafny.MultiSet
                d_321_v166_ = _dafny.MultiSet([d_97_v5_, d_97_v5_])
                (d_100_globalState_).f3 = default__.safeModuloInt((d_321_v166_).cardinality, d_313_i16_)
            elif True:
                (d_100_globalState_).f4 = (d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))]
                d_95_v3_ = (d_95_v3_).set(d_94_v2_, len(d_98_v6_))
                (d_100_globalState_).f4 = d_94_v2_
                d_322_v167_: D13
                d_322_v167_ = D13_DC30((d_293_v144_)[default__.safeIndex(300, (d_293_v144_).length(0))], d_94_v2_, d_132_v24_, d_219_v97_)
                (d_100_globalState_).f4 = not (not(d_94_v2_)) or ((d_322_v167_).cf43)
                d_323_v168_: _dafny.Map
                d_323_v168_ = _dafny.Map({d_94_v2_: d_97_v5_})
                (d_307_v154_).m4(d_94_v2_, d_323_v168_, 239, d_100_globalState_)
            (d_100_globalState_).f4 = True
            (d_100_globalState_).f4 = True
            d_101_v8_ = d_101_v8_
        d_324_v169_: C0
        nw61_ = C0()
        nw61_.ctor__()
        d_324_v169_ = nw61_
        d_325_v170_: _dafny.MultiSet
        d_325_v170_ = _dafny.MultiSet([d_97_v5_])
        (d_100_globalState_).f4 = (d_325_v170_).ispropersubset((d_325_v170_) | (d_325_v170_))
        _dafny.print(_dafny.string_of(d_92_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_93_v1_) == (_dafny.SeqWithoutIsStrInference([545, 545, 545, 545]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_94_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v3_) == (_dafny.Map({False: 545}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v4_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_97_v5_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v6_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cj"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v7_) == (_dafny.MultiSet([545]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_globalState_.f0) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_100_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_globalState_.f2) == (_dafny.MultiSet([545]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_100_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_100_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_100_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v8_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v9_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v9_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_105_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v24_) == (_dafny.Set({2, 545}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_220_v98_) == (_dafny.Map({545: 545}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_221_v99_).cf12) == (_dafny.Map({545: 545}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v100_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_223_v101_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_224_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_233_v110_).cf63) == (_dafny.SeqWithoutIsStrInference([1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v110_).cf64))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_235_i10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v144_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v157_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v158_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v159_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v159_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v159_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_325_v170_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmhjp"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0), False, False)
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

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(False, False, int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC11(D4, NamedTuple('DC11', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {self.cf17.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC14(D5, NamedTuple('DC14', [('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22
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
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC19(D8, NamedTuple('DC19', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC21(False, _dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC21(D9, NamedTuple('DC21', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC24(int(0), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC24(D10, NamedTuple('DC24', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {self.cf33.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)

class D11_DC25(D11, NamedTuple('DC25', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC27(int(0), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)

class D12_DC27(D12, NamedTuple('DC27', [('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC26(D12, NamedTuple('DC26', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC30(False, False, _dafny.Set({}), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)

class D13_DC30(D13, NamedTuple('DC30', [('cf42', Any), ('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC31(D13, NamedTuple('DC31', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC32(D14, NamedTuple('DC32', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)

class D15_DC33(D15, NamedTuple('DC33', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC35(int(0), _dafny.CodePoint('D'), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D16_DC34)

class D16_DC35(D16, NamedTuple('DC35', [('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC34(D16, NamedTuple('DC34', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC34({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC34) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
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
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D18_DC37)

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
        return lambda: D19_DC39(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D19_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D19_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D19_DC40)

class D19_DC39(D19, NamedTuple('DC39', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC39({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC39) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC38(D19, NamedTuple('DC38', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC38({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC38) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC40(D19, NamedTuple('DC40', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC40({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC40) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D20_DC41)

class D20_DC41(D20, NamedTuple('DC41', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC41({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC41) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC43(_dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D21_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D21_DC42)

class D21_DC43(D21, NamedTuple('DC43', [('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC43({_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC43) and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC42(D21, NamedTuple('DC42', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC42({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC42) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D22_DC44)

class D22_DC44(D22, NamedTuple('DC44', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC44({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC44) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC46()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D23_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D23_DC47)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D23_DC45)

class D23_DC46(D23, NamedTuple('DC46', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC46'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC46)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC47(D23, NamedTuple('DC47', [('cf67', Any), ('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC47({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC47) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC45(D23, NamedTuple('DC45', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC45({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC45) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC49()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D24_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D24_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D24_DC48)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D24_DC51)

class D24_DC49(D24, NamedTuple('DC49', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC49'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC49)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC50(D24, NamedTuple('DC50', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC50'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC50)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC48(D24, NamedTuple('DC48', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC48({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC48) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC51(D24, NamedTuple('DC51', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC51({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC51) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC53(_dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D25_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D25_DC52)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D25_DC54)

class D25_DC53(D25, NamedTuple('DC53', [('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC53({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC53) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC52(D25, NamedTuple('DC52', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC52({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC52) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC54(D25, NamedTuple('DC54', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC54({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC54) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC56(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D26_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D26_DC57)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D26_DC55)

class D26_DC56(D26, NamedTuple('DC56', [('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC56({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC56) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC57(D26, NamedTuple('DC57', [('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC57({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC57) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC55(D26, NamedTuple('DC55', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC55({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC55) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC59(False, _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D27_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D27_DC58)

class D27_DC59(D27, NamedTuple('DC59', [('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC59({_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC59) and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC58(D27, NamedTuple('DC58', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC58({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC58) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D28_DC60)

class D28_DC60(D28, NamedTuple('DC60', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC60({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC60) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Seq = _dafny.Seq({})
        self.f2: _dafny.MultiSet = _dafny.MultiSet({})
        self.f3: int = int(0)
        self.f4: bool = False
        self.f8: int = int(0)
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f5: bool = False
        self._f6: bool = False
        self._f7: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8

    @property
    def f1(self):
        return self._f1
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm19(self, globalState):
        return ((-686) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_326_i0_ in range(default__.abs(568))])))) - ((_dafny.MultiSet([529, (_dafny.MultiSet([True])).cardinality])).cardinality)


class C1(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm2(self, globalState):
        def iife74_():
            coll56_ = _dafny.Map()
            compr_56_: int
            for compr_56_ in _dafny.IntegerRange(-491, -220):
                d_327_v0_: int = compr_56_
                if ((-491) <= (d_327_v0_)) and ((d_327_v0_) < (-220)):
                    coll56_[(d_327_v0_) * (697)] = len((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([not(True), True])))
            return _dafny.Map(coll56_)
        return iife74_()
        

    def fm3(self, p0, p1, p2, globalState):
        return (not (True) or (not(True))) and ((_dafny.SeqWithoutIsStrInference([True, not(False)])) == (_dafny.SeqWithoutIsStrInference([True])))

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        if False:
            d_328_v0_: int
            d_328_v0_ = 847
            d_329_v1_: _dafny.Map
            d_329_v1_ = _dafny.Map({d_328_v0_: d_328_v0_})
            d_330_v2_: D3
            d_330_v2_ = D3_DC8(d_329_v1_)
            d_331_v3_: bool
            d_331_v3_ = True
            pat_let_tv10_ = d_329_v1_
            d_332_v4_: _dafny.Array
            nw62_ = _dafny.Array(None, 13)
            nw62_[int(0)] = d_330_v2_
            nw62_[int(1)] = (d_330_v2_ if not(d_331_v3_) else d_330_v2_)
            nw62_[int(2)] = d_330_v2_
            nw62_[int(3)] = D3_DC8(d_329_v1_)
            nw62_[int(4)] = d_330_v2_
            nw62_[int(5)] = d_330_v2_
            nw62_[int(6)] = d_330_v2_
            nw62_[int(7)] = d_330_v2_
            nw62_[int(8)] = default__.fm24(globalState)
            nw62_[int(9)] = d_330_v2_
            def iife75_(_pat_let9_0):
                def iife76_(d_333_dt__update__tmp_h0_):
                    def iife77_(_pat_let10_0):
                        def iife78_(d_334_dt__update_hcf12_h0_):
                            return D3_DC8(d_334_dt__update_hcf12_h0_)
                        return iife78_(_pat_let10_0)
                    return iife77_(pat_let_tv10_)
                return iife76_(_pat_let9_0)
            nw62_[int(10)] = (iife75_(d_330_v2_) if False else d_330_v2_)
            nw62_[int(11)] = d_330_v2_
            nw62_[int(12)] = d_330_v2_
            d_332_v4_ = nw62_
            index55_ = default__.safeIndex(504, (d_332_v4_).length(0))
            (d_332_v4_)[index55_] = d_330_v2_
            index56_ = default__.safeIndex(504, (d_332_v4_).length(0))
            (d_332_v4_)[index56_] = d_330_v2_
            d_335_v5_: C0
            nw63_ = C0()
            nw63_.ctor__()
            d_335_v5_ = nw63_
            d_336_v6_: _dafny.Set
            d_336_v6_ = _dafny.Set({d_328_v0_, d_328_v0_, default__.safeDivisionInt(d_328_v0_, d_328_v0_)})
            d_337_v7_: _dafny.Map
            d_337_v7_ = _dafny.Map({d_331_v3_: d_336_v6_})
            d_336_v6_ = (((d_337_v7_)[d_331_v3_] if (d_331_v3_) in (d_337_v7_) else d_336_v6_) if d_331_v3_ else (d_336_v6_) - (d_336_v6_))
            d_338_v8_: D3
            d_338_v8_ = D3_DC9()
            source10_ = d_338_v8_
            if source10_.is_DC9:
                d_339_v9_: _dafny.Set
                d_339_v9_ = _dafny.Set({d_331_v3_})
                d_340_v10_: _dafny.Array
                def lambda16_(d_341_v0_, d_342_v3_):
                    def lambda17_(d_343_i0_):
                        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_341_v0_, d_341_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_342_v3_]))), d_341_v0_]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_341_v0_])))

                    return lambda17_

                init8_ = lambda16_(d_328_v0_, d_331_v3_)
                nw64_ = _dafny.Array(None, 15)
                for i0_8_ in range(nw64_.length(0)):
                    nw64_[i0_8_] = init8_(i0_8_)
                d_340_v10_ = nw64_
                d_344_v11_: _dafny.Seq
                d_344_v11_ = _dafny.SeqWithoutIsStrInference([d_328_v0_, d_328_v0_, d_328_v0_, d_328_v0_, d_328_v0_])
                index57_ = default__.safeIndex(577, (d_340_v10_).length(0))
                (d_340_v10_)[index57_] = _dafny.MultiSet(d_344_v11_)
                d_345_v12_: _dafny.Seq
                d_345_v12_ = _dafny.SeqWithoutIsStrInference([d_331_v3_])
                d_346_v13_: _dafny.MultiSet
                d_346_v13_ = _dafny.MultiSet([d_328_v0_, d_328_v0_, len(d_339_v9_), d_328_v0_])
                d_347_v14_: D4
                d_347_v14_ = D4_DC10(d_345_v12_)
                d_348_v15_: _dafny.MultiSet
                d_348_v15_ = _dafny.MultiSet([d_347_v14_])
                d_349_v16_: _dafny.Map
                d_349_v16_ = _dafny.Map({d_348_v15_: d_328_v0_})
                d_350_v17_: _dafny.Map
                d_350_v17_ = _dafny.Map({d_328_v0_: d_349_v16_})
                index58_ = default__.safeIndex(577, (d_340_v10_).length(0))
                rhs44_ = len(d_345_v12_)
                rhs45_ = d_328_v0_
                rhs46_ = d_339_v9_
                rhs47_ = ((d_346_v13_) - (d_346_v13_)).intersection((d_346_v13_).intersection(d_346_v13_))
                rhs48_ = default__.safeModuloInt(d_328_v0_, len(((d_350_v17_)[(0) - (d_328_v0_)] if ((0) - (d_328_v0_)) in (d_350_v17_) else d_349_v16_)))
                lhs35_ = globalState
                lhs36_ = globalState
                lhs37_ = d_340_v10_
                lhs38_ = default__.safeIndex(577, (d_340_v10_).length(0))
                lhs39_ = globalState
                lhs35_.f8 = rhs44_
                lhs36_.f8 = rhs45_
                d_339_v9_ = rhs46_
                lhs37_[lhs38_] = rhs47_
                lhs39_.f8 = rhs48_
                d_351_v18_: _dafny.Seq
                d_351_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uou"))
                d_352_v19_: str
                d_352_v19_ = _dafny.CodePoint('r')
                rhs49_ = d_351_v18_
                rhs50_ = ((_dafny.CodePoint('j') if d_331_v3_ else d_352_v19_)) not in (d_351_v18_)
                lhs40_ = globalState
                d_351_v18_ = rhs49_
                lhs40_.f4 = rhs50_
                (globalState).f3 = (d_335_v5_).fm19(globalState)
                d_353_v20_: D2
                d_353_v20_ = D2_DC7(d_352_v19_, d_331_v3_, d_328_v0_)
                d_354_v21_: D4
                d_354_v21_ = D4_DC11(d_331_v3_, d_331_v3_, (d_344_v11_)[default__.safeIndex(d_328_v0_, len(d_344_v11_))], (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))).set(default__.safeIndex(d_328_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))), d_352_v19_))
                d_355_v22_: _dafny.Seq
                d_355_v22_ = d_344_v11_
                d_356_v23_: _dafny.Map
                d_356_v23_ = _dafny.Map({d_328_v0_: d_336_v6_})
                d_357_v24_: _dafny.Map
                d_357_v24_ = _dafny.Map({(0) - (d_328_v0_): d_346_v13_})
                d_358_v25_: _dafny.Seq
                d_358_v25_ = _dafny.SeqWithoutIsStrInference([d_357_v24_])
                d_359_v26_: _dafny.Array
                nw65_ = _dafny.Array(None, 28)
                nw65_[int(0)] = default__.safeDivisionInt(d_328_v0_, d_328_v0_)
                nw65_[int(1)] = d_328_v0_
                nw65_[int(2)] = d_328_v0_
                nw65_[int(3)] = len(_dafny.Map({d_328_v0_: d_328_v0_}))
                nw65_[int(4)] = (d_328_v0_) - (len(_dafny.SeqWithoutIsStrInference([d_352_v19_ for d_360_i1_ in range(default__.abs(294))])))
                nw65_[int(5)] = d_328_v0_
                nw65_[int(6)] = ((d_353_v20_).cf11 if d_331_v3_ else (d_354_v21_).cf16)
                nw65_[int(7)] = d_328_v0_
                nw65_[int(8)] = d_328_v0_
                nw65_[int(9)] = d_328_v0_
                nw65_[int(10)] = d_328_v0_
                nw65_[int(11)] = d_328_v0_
                nw65_[int(12)] = -245
                nw65_[int(13)] = ((d_329_v1_)[(0) - (((d_329_v1_)[d_328_v0_] if (d_328_v0_) in (d_329_v1_) else d_328_v0_))] if ((0) - (((d_329_v1_)[d_328_v0_] if (d_328_v0_) in (d_329_v1_) else d_328_v0_))) in (d_329_v1_) else d_328_v0_)
                nw65_[int(14)] = ((0) - ((((d_340_v10_)[default__.safeIndex(577, (d_340_v10_).length(0))])[d_328_v0_] if (d_328_v0_) in ((d_340_v10_)[default__.safeIndex(577, (d_340_v10_).length(0))]) else d_328_v0_))) - (d_328_v0_)
                nw65_[int(15)] = (d_328_v0_) + (d_328_v0_)
                nw65_[int(16)] = d_328_v0_
                nw65_[int(17)] = len((d_355_v22_))
                nw65_[int(18)] = d_328_v0_
                nw65_[int(19)] = d_328_v0_
                nw65_[int(20)] = d_328_v0_
                nw65_[int(21)] = default__.safeDivisionInt(len(d_356_v23_), d_328_v0_)
                nw65_[int(22)] = (d_328_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nalyv"))))
                nw65_[int(23)] = ((d_335_v5_).fm19(globalState) if True else d_328_v0_)
                nw65_[int(24)] = d_328_v0_
                nw65_[int(25)] = 715
                nw65_[int(26)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_352_v19_ for d_361_i2_ in range(default__.abs(-239))])), len((d_358_v25_)[default__.safeIndex(512, len(d_358_v25_))]))
                nw65_[int(27)] = (0) - (d_328_v0_)
                d_359_v26_ = nw65_
                index59_ = default__.safeIndex(376, (d_359_v26_).length(0))
                (d_359_v26_)[index59_] = default__.safeDivisionInt(d_328_v0_, d_328_v0_)
                index60_ = default__.safeIndex(376, (d_359_v26_).length(0))
                (d_359_v26_)[index60_] = (0) - ((865) - (len(d_351_v18_)))
            elif True:
                d_362___mcc_h0_ = source10_.cf12
                d_363_cf12_ = d_362___mcc_h0_
                d_364_v27_: _dafny.Array
                nw66_ = _dafny.Array(int(0), 15)
                d_364_v27_ = nw66_
                d_365_v28_: _dafny.Seq
                d_365_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqitorcr"))
                index61_ = default__.safeIndex(257, (d_364_v27_).length(0))
                (d_364_v27_)[index61_] = (len(d_365_v28_)) + (d_328_v0_)
                index62_ = default__.safeIndex(257, (d_364_v27_).length(0))
                (d_364_v27_)[index62_] = d_328_v0_
                d_366_v29_: _dafny.Set
                d_366_v29_ = _dafny.Set({d_331_v3_, d_331_v3_, d_331_v3_})
                d_331_v3_ = (d_366_v29_).issubset(_dafny.Set({not(d_331_v3_)}))
                (globalState).f8 = (d_364_v27_)[default__.safeIndex(257, (d_364_v27_).length(0))]
                d_367_v30_: _dafny.Array
                def lambda18_(d_368_v0_):
                    def lambda19_(d_369_i3_):
                        return (_dafny.Map({_dafny.CodePoint('a'): d_368_v0_})) | (_dafny.Map({_dafny.CodePoint('n'): -308}))

                    return lambda19_

                init9_ = lambda18_(d_328_v0_)
                nw67_ = _dafny.Array(None, 17)
                for i0_9_ in range(nw67_.length(0)):
                    nw67_[i0_9_] = init9_(i0_9_)
                d_367_v30_ = nw67_
                d_367_v30_ = d_367_v30_
            d_370_v31_: _dafny.Seq
            d_370_v31_ = _dafny.SeqWithoutIsStrInference([False, d_331_v3_])
            if (d_370_v31_)[default__.safeIndex(d_328_v0_, len(d_370_v31_))]:
                d_371_v32_: _dafny.Seq
                d_371_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prrqlk"))
                d_372_v33_: _dafny.Array
                nw68_ = _dafny.Array(False, 16)
                d_372_v33_ = nw68_
                d_373_v34_: _dafny.Map
                d_373_v34_ = _dafny.Map({d_331_v3_: d_372_v33_})
                d_374_v35_: str
                d_374_v35_ = _dafny.CodePoint('v')
                (globalState).f8 = len(((default__.fm1(d_328_v0_, _dafny.CodePoint('e'), d_331_v3_, d_329_v1_, globalState)).set(default__.safeIndex(-7, len(default__.fm1(d_328_v0_, _dafny.CodePoint('e'), d_331_v3_, d_329_v1_, globalState))), len(d_371_v32_))) + (default__.fm1(len(d_373_v34_), d_374_v35_, False, d_329_v1_, globalState)))
                d_375_v36_: _dafny.Map
                d_375_v36_ = _dafny.Map({d_331_v3_: d_331_v3_})
                d_376_v37_: _dafny.Map
                d_376_v37_ = _dafny.Map({d_375_v36_: d_331_v3_})
                d_377_v38_: _dafny.Array
                nw69_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                d_377_v38_ = nw69_
                index63_ = default__.safeIndex(347, (d_377_v38_).length(0))
                (d_377_v38_)[index63_] = d_374_v35_
                index64_ = default__.safeIndex(347, (d_377_v38_).length(0))
                rhs51_ = default__.fm25(globalState)
                rhs52_ = ((d_328_v0_) * (((d_329_v1_)[d_328_v0_] if (d_328_v0_) in (d_329_v1_) else d_328_v0_)) if d_331_v3_ else d_328_v0_)
                rhs53_ = (d_376_v37_) | (d_376_v37_)
                rhs54_ = d_374_v35_
                rhs55_ = default__.safeDivisionInt(d_328_v0_, (d_328_v0_) + ((0) - (d_328_v0_)))
                lhs41_ = globalState
                lhs42_ = d_377_v38_
                lhs43_ = default__.safeIndex(347, (d_377_v38_).length(0))
                lhs44_ = globalState
                d_371_v32_ = rhs51_
                lhs41_.f3 = rhs52_
                d_376_v37_ = rhs53_
                lhs42_[lhs43_] = rhs54_
                lhs44_.f8 = rhs55_
                d_378_v39_: _dafny.MultiSet
                d_378_v39_ = _dafny.MultiSet([d_328_v0_])
                d_379_v40_: _dafny.MultiSet
                d_379_v40_ = _dafny.MultiSet([not(d_331_v3_)])
                d_380_v41_: _dafny.Seq
                d_380_v41_ = _dafny.SeqWithoutIsStrInference([(0) - (d_328_v0_), (d_378_v39_).cardinality, d_328_v0_, d_328_v0_, ((d_379_v40_)[d_331_v3_] if (d_331_v3_) in (d_379_v40_) else (0) - (d_328_v0_))])
                (globalState).f8 = (d_380_v41_)[default__.safeIndex((d_328_v0_) * (d_328_v0_), len(d_380_v41_))]
                d_331_v3_ = d_331_v3_
                (globalState).f8 = (0) - (len(_dafny.Map({d_374_v35_: (d_371_v32_) == (d_371_v32_)})))
            elif True:
                d_381_v42_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_381_v42_ = nw70_
                d_382_v43_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_382_v43_ = nw71_
                index65_ = default__.safeIndex(183, (d_381_v42_).length(0))
                (d_381_v42_)[index65_] = d_382_v43_
                d_383_v44_: D5
                d_383_v44_ = D5_DC14(d_328_v0_, not(d_331_v3_))
                d_384_v45_: _dafny.Seq
                d_384_v45_ = _dafny.SeqWithoutIsStrInference([d_383_v44_])
                d_385_v46_: _dafny.Set
                d_385_v46_ = _dafny.Set({d_331_v3_, d_331_v3_, d_331_v3_, d_331_v3_})
                d_386_v47_: D9
                d_386_v47_ = D9_DC20(d_385_v46_)
                d_387_v48_: _dafny.Seq
                d_387_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fw"))
                d_388_v49_: _dafny.Seq
                d_388_v49_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_387_v48_))])
                pat_let_tv11_ = d_385_v46_
                pat_let_tv12_ = d_388_v49_
                d_389_v50_: _dafny.Array
                nw72_ = _dafny.Array(None, 19)
                nw72_[int(0)] = d_331_v3_
                nw72_[int(1)] = d_331_v3_
                nw72_[int(2)] = d_331_v3_
                nw72_[int(3)] = d_331_v3_
                nw72_[int(4)] = (d_331_v3_) or (d_331_v3_)
                nw72_[int(5)] = d_331_v3_
                nw72_[int(6)] = d_331_v3_
                nw72_[int(7)] = d_331_v3_
                nw72_[int(8)] = not(d_331_v3_)
                nw72_[int(9)] = not(d_331_v3_)
                nw72_[int(10)] = d_331_v3_
                nw72_[int(11)] = d_331_v3_
                nw72_[int(12)] = d_331_v3_
                nw72_[int(13)] = d_331_v3_
                nw72_[int(14)] = d_331_v3_
                def iife79_(_pat_let11_0):
                    def iife80_(d_390_dt__update__tmp_h1_):
                        def iife81_(_pat_let12_0):
                            def iife82_(d_391_dt__update_hcf25_h0_):
                                return D9_DC20(d_391_dt__update_hcf25_h0_)
                            return iife82_(_pat_let12_0)
                        return iife81_(pat_let_tv11_)
                    return iife80_(_pat_let11_0)
                nw72_[int(15)] = not(((iife79_(d_386_v47_)).cf25).ispropersubset(d_385_v46_))
                nw72_[int(16)] = d_331_v3_
                def iife83_(_pat_let13_0):
                    def iife84_(d_392_dt__update__tmp_h2_):
                        def iife85_(_pat_let14_0):
                            def iife86_(d_393_dt__update_hcf20_h0_):
                                return D5_DC14(d_393_dt__update_hcf20_h0_, (d_392_dt__update__tmp_h2_).cf21)
                            return iife86_(_pat_let14_0)
                        return iife85_(len(pat_let_tv12_))
                    return iife84_(_pat_let13_0)
                nw72_[int(17)] = (iife83_(d_383_v44_)).cf21
                nw72_[int(18)] = d_331_v3_
                d_389_v50_ = nw72_
                index66_ = default__.safeIndex(436, (d_389_v50_).length(0))
                (d_389_v50_)[index66_] = not(d_331_v3_)
                index67_ = default__.safeIndex(183, (d_381_v42_).length(0))
                index68_ = default__.safeIndex(436, (d_389_v50_).length(0))
                rhs56_ = d_331_v3_
                rhs57_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqv"))) <= (d_387_v48_)
                rhs58_ = d_382_v43_
                rhs59_ = (d_384_v45_) + (d_384_v45_)
                rhs60_ = False
                lhs45_ = d_381_v42_
                lhs46_ = default__.safeIndex(183, (d_381_v42_).length(0))
                lhs47_ = d_389_v50_
                lhs48_ = default__.safeIndex(436, (d_389_v50_).length(0))
                d_331_v3_ = rhs56_
                r0 = rhs57_
                lhs45_[lhs46_] = rhs58_
                d_384_v45_ = rhs59_
                lhs47_[lhs48_] = rhs60_
                d_394_v51_: C0
                nw73_ = C0()
                nw73_.ctor__()
                d_394_v51_ = nw73_
                d_395_v52_: _dafny.Array
                nw74_ = _dafny.Array(int(0), 2)
                d_395_v52_ = nw74_
                index69_ = default__.safeIndex(931, (d_395_v52_).length(0))
                (d_395_v52_)[index69_] = d_328_v0_
                index70_ = default__.safeIndex(931, (d_395_v52_).length(0))
                index71_ = default__.safeIndex(436, (d_389_v50_).length(0))
                rhs61_ = (901) * ((d_328_v0_) + (d_328_v0_))
                rhs62_ = d_328_v0_
                rhs63_ = (d_335_v5_).fm19(globalState)
                rhs64_ = (d_389_v50_)[default__.safeIndex(436, (d_389_v50_).length(0))]
                lhs49_ = globalState
                lhs50_ = globalState
                lhs51_ = d_395_v52_
                lhs52_ = default__.safeIndex(931, (d_395_v52_).length(0))
                lhs53_ = d_389_v50_
                lhs54_ = default__.safeIndex(436, (d_389_v50_).length(0))
                lhs49_.f8 = rhs61_
                lhs50_.f3 = rhs62_
                lhs51_[lhs52_] = rhs63_
                lhs53_[lhs54_] = rhs64_
                d_396_v53_: _dafny.MultiSet
                d_396_v53_ = _dafny.MultiSet([d_331_v3_])
                d_396_v53_ = _dafny.MultiSet([d_331_v3_, (d_389_v50_)[default__.safeIndex(436, (d_389_v50_).length(0))], not(d_331_v3_), ((d_389_v50_)[default__.safeIndex(436, (d_389_v50_).length(0))] if (d_370_v31_)[default__.safeIndex((d_395_v52_)[default__.safeIndex(931, (d_395_v52_).length(0))], len(d_370_v31_))] else d_331_v3_)])
                rhs65_ = d_331_v3_
                rhs66_ = False
                rhs67_ = d_395_v52_
                lhs55_ = globalState
                lhs55_.f4 = rhs65_
                r0 = rhs66_
                d_395_v52_ = rhs67_
        elif True:
            d_397_v54_: bool
            d_397_v54_ = False
            if d_397_v54_:
                d_398_v55_: int
                d_398_v55_ = -909
                d_399_v56_: _dafny.Map
                d_399_v56_ = _dafny.Map({d_398_v55_: _dafny.SeqWithoutIsStrInference([not(True)])})
                d_400_v57_: D4
                d_400_v57_ = D4_DC10(_dafny.SeqWithoutIsStrInference([False, d_397_v54_]))
                d_401_v58_: _dafny.Seq
                d_401_v58_ = _dafny.SeqWithoutIsStrInference([D4_DC10(((d_399_v56_)[d_398_v55_] if (d_398_v55_) in (d_399_v56_) else _dafny.SeqWithoutIsStrInference([d_397_v54_]))), d_400_v57_, d_400_v57_, d_400_v57_, d_400_v57_])
                d_401_v58_ = d_401_v58_
                d_402_v59_: _dafny.Seq
                d_402_v59_ = _dafny.SeqWithoutIsStrInference([d_398_v55_])
                d_403_v60_: _dafny.Map
                d_403_v60_ = _dafny.Map({d_397_v54_: d_397_v54_})
                d_404_v61_: _dafny.MultiSet
                d_404_v61_ = _dafny.MultiSet([-813, d_398_v55_, len(d_403_v60_)])
                d_405_v62_: _dafny.Map
                d_405_v62_ = _dafny.Map({(self).fm3(d_402_v59_, d_404_v61_, d_397_v54_, globalState): d_398_v55_})
                d_406_v63_: str
                d_406_v63_ = _dafny.CodePoint('y')
                d_407_v64_: _dafny.Map
                d_407_v64_ = _dafny.Map({(d_398_v55_) - (len(d_405_v62_)): (d_406_v63_ if d_397_v54_ else d_406_v63_)})
                d_407_v64_ = (d_407_v64_).set((d_398_v55_) * (d_398_v55_), d_406_v63_)
                d_408_v65_: _dafny.Array
                def lambda20_(d_409_v63_):
                    def lambda21_(d_410_i4_):
                        return d_409_v63_

                    return lambda21_

                init10_ = lambda20_(d_406_v63_)
                nw75_ = _dafny.Array(None, 26)
                for i0_10_ in range(nw75_.length(0)):
                    nw75_[i0_10_] = init10_(i0_10_)
                d_408_v65_ = nw75_
                index72_ = default__.safeIndex(51, (d_408_v65_).length(0))
                (d_408_v65_)[index72_] = d_406_v63_
                d_411_v66_: _dafny.Array
                nw76_ = _dafny.Array(False, 12)
                d_411_v66_ = nw76_
                index73_ = default__.safeIndex(51, (d_408_v65_).length(0))
                rhs68_ = (d_398_v55_) + (d_398_v55_)
                rhs69_ = d_406_v63_
                rhs70_ = d_411_v66_
                lhs56_ = globalState
                lhs57_ = d_408_v65_
                lhs58_ = default__.safeIndex(51, (d_408_v65_).length(0))
                lhs56_.f3 = rhs68_
                lhs57_[lhs58_] = rhs69_
                d_411_v66_ = rhs70_
                d_412_v67_: C0
                nw77_ = C0()
                nw77_.ctor__()
                d_412_v67_ = nw77_
                (globalState).f8 = (d_402_v59_)[default__.safeIndex((0) - (d_398_v55_), len(d_402_v59_))]
            elif True:
                d_413_v68_: _dafny.Set
                d_413_v68_ = _dafny.Set({True})
                d_414_v69_: _dafny.Seq
                d_414_v69_ = _dafny.SeqWithoutIsStrInference([d_397_v54_, (_dafny.Set({True, d_397_v54_})).isdisjoint(d_413_v68_)])
                d_415_v70_: int
                d_415_v70_ = 70
                (globalState).f4 = (d_414_v69_)[default__.safeIndex(d_415_v70_, len(d_414_v69_))]
                r0 = d_397_v54_
                d_416_v71_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                d_416_v71_ = nw78_
                d_417_v72_: str
                d_417_v72_ = _dafny.CodePoint('m')
                index74_ = default__.safeIndex(358, (d_416_v71_).length(0))
                (d_416_v71_)[index74_] = d_417_v72_
                index75_ = default__.safeIndex(358, (d_416_v71_).length(0))
                (d_416_v71_)[index75_] = d_417_v72_
                index76_ = default__.safeIndex(358, (d_416_v71_).length(0))
                (d_416_v71_)[index76_] = (d_416_v71_)[default__.safeIndex(358, (d_416_v71_).length(0))]
                d_418_v73_: _dafny.Seq
                d_418_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcrxtqw"))
                d_418_v73_ = d_418_v73_
            d_419_v74_: _dafny.Seq
            d_419_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilxgf"))
            d_420_v75_: int
            d_420_v75_ = 857
            d_421_v76_: _dafny.Map
            d_421_v76_ = _dafny.Map({len(d_419_v74_): d_420_v75_})
            d_422_v78_: str
            d_422_v78_ = _dafny.CodePoint('u')
            def iife87_():
                coll57_ = _dafny.Map()
                compr_57_: int
                for compr_57_ in _dafny.IntegerRange(513, 938):
                    d_423_v77_: int = compr_57_
                    if ((513) <= (d_423_v77_)) and ((d_423_v77_) < (938)):
                        coll57_[default__.safeModuloInt(d_423_v77_, ((D10_DC23(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_419_v74_])))).cf30).cardinality)] = d_397_v54_
                return _dafny.Map(coll57_)
            d_421_v76_ = (d_421_v76_).set(default__.safeDivisionInt(len(iife87_()
            ), d_420_v75_), len((_dafny.Map({d_422_v78_: d_422_v78_}) if d_397_v54_ else _dafny.Map({d_422_v78_: d_422_v78_}))))
            d_424_v79_: _dafny.Seq
            d_424_v79_ = _dafny.SeqWithoutIsStrInference([d_397_v54_, d_397_v54_])
            d_425_v80_: _dafny.Set
            d_425_v80_ = _dafny.Set({d_397_v54_})
            d_426_v81_: _dafny.Seq
            d_426_v81_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({d_397_v54_}) if d_397_v54_ else d_425_v80_)])
            rhs71_ = (0) - (default__.safeDivisionInt(default__.safeDivisionInt(d_420_v75_, 669), (len(d_424_v79_)) * (d_420_v75_)))
            rhs72_ = len(d_426_v81_)
            lhs59_ = globalState
            lhs60_ = globalState
            lhs59_.f8 = rhs71_
            lhs60_.f3 = rhs72_
            d_427_v82_: D3
            d_427_v82_ = D3_DC8(d_421_v76_)
            d_428_v83_: _dafny.MultiSet
            d_428_v83_ = _dafny.MultiSet([d_424_v79_])
            rhs73_ = d_420_v75_
            rhs74_ = default__.fm26(d_427_v82_, d_420_v75_, d_428_v83_, globalState)
            lhs61_ = globalState
            lhs61_.f3 = rhs73_
            d_420_v75_ = rhs74_
            d_429_v84_: _dafny.Seq
            d_429_v84_ = _dafny.SeqWithoutIsStrInference([d_420_v75_])
            d_430_v85_: D5
            d_430_v85_ = D5_DC15()
            source11_ = (d_430_v85_ if (733) not in (d_429_v84_) else d_430_v85_)
            if source11_.is_DC14:
                d_431___mcc_h1_ = source11_.cf20
                d_432___mcc_h2_ = source11_.cf21
                d_433_cf21_ = d_432___mcc_h2_
                d_434_cf20_ = d_431___mcc_h1_
                (globalState).f4 = d_433_cf21_
                d_435_v86_: _dafny.Array
                nw79_ = _dafny.Array(None, 23)
                nw79_[int(0)] = d_434_cf20_
                nw79_[int(1)] = 340
                nw79_[int(2)] = d_434_cf20_
                nw79_[int(3)] = d_434_cf20_
                nw79_[int(4)] = d_420_v75_
                nw79_[int(5)] = d_420_v75_
                nw79_[int(6)] = d_434_cf20_
                nw79_[int(7)] = d_420_v75_
                nw79_[int(8)] = default__.fm26(d_427_v82_, d_434_cf20_, d_428_v83_, globalState)
                nw79_[int(9)] = (0) - (d_434_cf20_)
                nw79_[int(10)] = d_434_cf20_
                nw79_[int(11)] = 78
                nw79_[int(12)] = d_420_v75_
                nw79_[int(13)] = d_420_v75_
                nw79_[int(14)] = d_434_cf20_
                nw79_[int(15)] = d_420_v75_
                nw79_[int(16)] = len(d_419_v74_)
                nw79_[int(17)] = d_434_cf20_
                nw79_[int(18)] = d_420_v75_
                nw79_[int(19)] = d_434_cf20_
                nw79_[int(20)] = d_420_v75_
                nw79_[int(21)] = d_434_cf20_
                nw79_[int(22)] = -573
                d_435_v86_ = nw79_
                d_436_v87_: _dafny.Map
                d_436_v87_ = _dafny.Map({d_420_v75_: d_435_v86_})
                d_437_v88_: _dafny.Map
                d_437_v88_ = _dafny.Map({_dafny.MultiSet([((d_436_v87_)[d_434_cf20_] if (d_434_cf20_) in (d_436_v87_) else d_435_v86_), d_435_v86_]): not(d_397_v54_)})
                d_438_v89_: _dafny.MultiSet
                d_438_v89_ = _dafny.MultiSet([d_435_v86_])
                d_433_cf21_ = ((d_437_v88_)[(d_438_v89_) | (d_438_v89_)] if ((d_438_v89_) | (d_438_v89_)) in (d_437_v88_) else d_397_v54_)
                r0 = d_397_v54_
                (globalState).f8 = d_434_cf20_
            elif source11_.is_DC15:
                (globalState).f3 = d_420_v75_
                (globalState).f3 = d_420_v75_
                d_439_v90_: D2
                d_439_v90_ = D2_DC7(d_422_v78_, d_397_v54_, d_420_v75_)
                d_439_v90_ = D2_DC7(d_422_v78_, d_397_v54_, d_420_v75_)
                (globalState).f3 = (0) - ((0) - ((len(d_419_v74_)) + (d_420_v75_)))
            elif True:
                d_440___mcc_h3_ = source11_.cf19
                d_441_cf19_ = d_440___mcc_h3_
                d_442_v91_: _dafny.Map
                d_442_v91_ = _dafny.Map({d_425_v80_: default__.fm26(d_427_v82_, len(d_419_v74_), d_428_v83_, globalState)})
                r0 = (d_442_v91_) == (d_442_v91_)
                d_443_v92_: _dafny.MultiSet
                d_443_v92_ = _dafny.MultiSet([d_397_v54_])
                r0 = ((d_420_v75_) + (len(d_419_v74_))) > ((d_443_v92_).cardinality)
                r0 = d_397_v54_
                d_444_v93_: _dafny.MultiSet
                d_444_v93_ = _dafny.MultiSet([627])
                (self).m13(d_420_v75_, (d_444_v93_).intersection(d_444_v93_), not((_dafny.Set({False})).isdisjoint(d_425_v80_)), globalState)
        d_445_v94_: bool
        d_445_v94_ = True
        (globalState).f4 = d_445_v94_
        d_446_v95_: _dafny.Array
        def lambda22_(d_447_v94_):
            def lambda23_(d_448_i5_):
                return d_447_v94_

            return lambda23_

        init11_ = lambda22_(d_445_v94_)
        nw80_ = _dafny.Array(None, 17)
        for i0_11_ in range(nw80_.length(0)):
            nw80_[i0_11_] = init11_(i0_11_)
        d_446_v95_ = nw80_
        index77_ = default__.safeIndex(782, (d_446_v95_).length(0))
        (d_446_v95_)[index77_] = (d_445_v94_ if d_445_v94_ else d_445_v94_)
        d_449_v96_: _dafny.Array
        nw81_ = _dafny.Array(_dafny.CodePoint('D'), 17)
        d_449_v96_ = nw81_
        d_450_v97_: _dafny.Array
        nw82_ = _dafny.Array(int(0), 18)
        d_450_v97_ = nw82_
        d_451_v98_: _dafny.Map
        d_451_v98_ = _dafny.Map({d_449_v96_: d_450_v97_})
        index78_ = default__.safeIndex(782, (d_446_v95_).length(0))
        (d_446_v95_)[index78_] = (d_451_v98_) == (d_451_v98_)
        d_452_i6_: int
        d_452_i6_ = 0
        with _dafny.label("3"):
            while (d_446_v95_)[default__.safeIndex(782, (d_446_v95_).length(0))]:
                with _dafny.c_label("3"):
                    if (d_452_i6_) >= (100):
                        raise _dafny.Break("3")
                    d_452_i6_ = (d_452_i6_) + (1)
                    d_453_v99_: int
                    d_453_v99_ = -172
                    d_454_v100_: _dafny.Seq
                    d_454_v100_ = _dafny.SeqWithoutIsStrInference([d_453_v99_, d_453_v99_, d_453_v99_])
                    d_455_v101_: _dafny.MultiSet
                    d_455_v101_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_453_v99_]), d_454_v100_])
                    d_456_v102_: _dafny.Map
                    d_456_v102_ = _dafny.Map({(0) - (d_453_v99_): d_455_v101_})
                    d_457_v103_: D3
                    d_457_v103_ = D3_DC8(_dafny.Map({d_453_v99_: d_453_v99_}))
                    d_458_v104_: _dafny.MultiSet
                    d_458_v104_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])
                    d_456_v102_ = (d_456_v102_).set(default__.fm26(d_457_v103_, d_453_v99_, d_458_v104_, globalState), d_455_v101_)
                    d_453_v99_ = d_453_v99_
                    d_459_v105_: C0
                    nw83_ = C0()
                    nw83_.ctor__()
                    d_459_v105_ = nw83_
                    (globalState).f8 = ((d_453_v99_) + (d_453_v99_)) + (d_453_v99_)
                    pass
            pass
        d_460_v106_: int
        d_460_v106_ = 160
        d_461_v107_: _dafny.Map
        d_461_v107_ = _dafny.Map({d_460_v106_: d_460_v106_})
        (globalState).f8 = (default__.safeModuloInt(d_460_v106_, d_460_v106_)) * ((602) * (len(d_461_v107_)))
        index79_ = default__.safeIndex(991, (d_450_v97_).length(0))
        (d_450_v97_)[index79_] = d_460_v106_
        d_462_v108_: _dafny.Seq
        d_462_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
        d_463_v109_: D2
        d_463_v109_ = D2_DC6(d_460_v106_, True, (d_446_v95_)[default__.safeIndex(782, (d_446_v95_).length(0))])
        index80_ = default__.safeIndex(991, (d_450_v97_).length(0))
        rhs75_ = (0) - (default__.safeModuloInt(d_460_v106_, (d_460_v106_) * ((0) - (d_460_v106_))))
        rhs76_ = (d_446_v95_)[default__.safeIndex(782, (d_446_v95_).length(0))]
        rhs77_ = (d_463_v109_).cf8
        rhs78_ = d_462_v108_
        rhs79_ = False
        lhs62_ = d_450_v97_
        lhs63_ = default__.safeIndex(991, (d_450_v97_).length(0))
        lhs64_ = globalState
        lhs62_[lhs63_] = rhs75_
        r1 = rhs76_
        lhs64_.f4 = rhs77_
        d_462_v108_ = rhs78_
        r1 = rhs79_
        r0 = d_445_v94_
        pat_let_tv13_ = d_450_v97_
        pat_let_tv14_ = d_450_v97_
        pat_let_tv15_ = d_445_v94_
        pat_let_tv16_ = d_446_v95_
        pat_let_tv17_ = d_446_v95_
        pat_let_tv18_ = d_445_v94_
        def lambda24_(source12_):
            if source12_.is_DC14:
                d_464___mcc_h4_ = source12_.cf20
                d_465___mcc_h5_ = source12_.cf21
                d_466_cf21_ = d_465___mcc_h5_
                d_467_cf20_ = d_464___mcc_h4_
                return (D2_DC6((pat_let_tv14_)[default__.safeIndex(991, (pat_let_tv13_).length(0))], True, d_466_cf21_)).cf8
            elif source12_.is_DC15:
                return (_dafny.Set({pat_let_tv15_})).issubset(_dafny.Set({(pat_let_tv17_)[default__.safeIndex(782, (pat_let_tv16_).length(0))]}))
            elif True:
                d_468___mcc_h6_ = source12_.cf19
                d_469_cf19_ = d_468___mcc_h6_
                return pat_let_tv18_

        r1 = lambda24_(D5_DC15())
        return r0, r1

    def m13(self, p0, p1, p2, globalState):
        d_470_v0_: _dafny.Seq
        d_470_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eahalakud"))
        if ((len(d_470_v0_)) - (150)) in (p1):
            d_471_v1_: _dafny.Array
            nw84_ = _dafny.Array(int(0), 4)
            d_471_v1_ = nw84_
            d_472_v2_: _dafny.Seq
            d_472_v2_ = _dafny.SeqWithoutIsStrInference([d_471_v1_])
            d_472_v2_ = d_472_v2_
            index81_ = default__.safeIndex(958, (d_471_v1_).length(0))
            (d_471_v1_)[index81_] = p0
            index82_ = default__.safeIndex(958, (d_471_v1_).length(0))
            (d_471_v1_)[index82_] = p0
            index83_ = default__.safeIndex(958, (d_471_v1_).length(0))
            (d_471_v1_)[index83_] = (d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))]
            def lambda25_(d_473_v1_):
                def lambda26_(d_474_i0_):
                    return (d_474_i0_) + ((d_473_v1_)[default__.safeIndex(958, (d_473_v1_).length(0))])

                return lambda26_

            init12_ = lambda25_(d_471_v1_)
            nw85_ = _dafny.Array(None, 24)
            for i0_12_ in range(nw85_.length(0)):
                nw85_[i0_12_] = init12_(i0_12_)
            d_471_v1_ = nw85_
            d_475_v3_: _dafny.Seq
            d_475_v3_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (-256)])
            d_476_v4_: D9
            d_476_v4_ = D9_DC21(False, _dafny.SeqWithoutIsStrInference([p2, p2, (self).fm3(d_475_v3_, p1, p2, globalState), p2]), p2)
            d_477_v5_: _dafny.Seq
            d_477_v5_ = _dafny.SeqWithoutIsStrInference([not(p2), not((self).fm3(_dafny.SeqWithoutIsStrInference([(d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))] for d_478_i1_ in range(default__.abs(586))]), p1, not(False), globalState)), p2])
            pat_let_tv19_ = p2
            def iife88_(_pat_let15_0):
                def iife89_(d_479_dt__update__tmp_h0_):
                    def iife90_(_pat_let16_0):
                        def iife91_(d_480_dt__update_hcf26_h0_):
                            return D9_DC21(d_480_dt__update_hcf26_h0_, (d_479_dt__update__tmp_h0_).cf27, (d_479_dt__update__tmp_h0_).cf28)
                        return iife91_(_pat_let16_0)
                    return iife90_(pat_let_tv19_)
                return iife89_(_pat_let15_0)
            if ((iife88_(d_476_v4_)).cf27) == ((d_477_v5_) + (d_477_v5_)):
                (globalState).f3 = (p0 if p2 else 655)
                d_471_v1_ = d_471_v1_
                (globalState).f4 = p2
                d_481_v6_: D2
                d_481_v6_ = D2_DC6(p0, (p0) == (p0), p2)
                d_481_v6_ = d_481_v6_
                (globalState).f8 = (d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))]
            elif True:
                d_482_v7_: C0
                nw86_ = C0()
                nw86_.ctor__()
                d_482_v7_ = nw86_
                d_483_v9_: _dafny.Array
                def lambda27_(d_484_p2_):
                    def lambda28_(d_485_i2_):
                        def iife92_():
                            coll58_ = _dafny.Map()
                            compr_58_: bool
                            for compr_58_ in (_dafny.Set({d_484_p2_, False})).Elements:
                                d_486_v8_: bool = compr_58_
                                if (d_486_v8_) in (_dafny.Set({d_484_p2_, False})):
                                    coll58_[d_486_v8_] = d_484_p2_
                            return _dafny.Map(coll58_)
                        return (_dafny.Map({d_484_p2_: d_484_p2_})) | (iife92_()
                        )

                    return lambda28_

                init13_ = lambda27_(p2)
                nw87_ = _dafny.Array(None, 20)
                for i0_13_ in range(nw87_.length(0)):
                    nw87_[i0_13_] = init13_(i0_13_)
                d_483_v9_ = nw87_
                d_487_v10_: bool
                d_487_v10_ = p2
                d_488_v11_: D2
                d_488_v11_ = D2_DC6(len(_dafny.Map({p0: not(p2)})), p2, p2)
                d_489_v12_: _dafny.Map
                d_489_v12_ = _dafny.Map({d_487_v10_: not((d_488_v11_).cf8)})
                index84_ = default__.safeIndex(378, (d_483_v9_).length(0))
                (d_483_v9_)[index84_] = d_489_v12_
                d_490_v14_: _dafny.Set
                d_490_v14_ = _dafny.Set({d_487_v10_, d_487_v10_})
                index85_ = default__.safeIndex(378, (d_483_v9_).length(0))
                def iife93_():
                    def iife95_():
                        coll61_ = _dafny.Set()
                        compr_61_: bool
                        for compr_61_ in (d_490_v14_).Elements:
                            d_493_v15_: bool = compr_61_
                            if (d_493_v15_) in (d_490_v14_):
                                coll61_ = coll61_.union(_dafny.Set([d_493_v15_]))
                        return _dafny.Set(coll61_)
                    coll59_ = _dafny.Map()
                    def iife94_():
                        coll60_ = _dafny.Set()
                        compr_60_: bool
                        for compr_60_ in (d_490_v14_).Elements:
                            d_491_v15_: bool = compr_60_
                            if (d_491_v15_) in (d_490_v14_):
                                coll60_ = coll60_.union(_dafny.Set([d_491_v15_]))
                        return _dafny.Set(coll60_)
                    compr_59_: bool
                    for compr_59_ in (iife94_()
                    ).Elements:
                        d_492_v13_: bool = compr_59_
                        if (d_492_v13_) in (iife95_()
                        ):
                            coll59_[d_492_v13_] = p2
                    return _dafny.Map(coll59_)
                (d_483_v9_)[index85_] = iife93_()
                
                (globalState).f8 = (d_482_v7_).fm19(globalState)
                d_494_v16_: str
                d_494_v16_ = _dafny.CodePoint('l')
                d_495_v17_: _dafny.Map
                d_495_v17_ = _dafny.Map({d_482_v7_: d_494_v16_})
                d_496_v18_: _dafny.Map
                d_496_v18_ = _dafny.Map({(d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))]: len(_dafny.SeqWithoutIsStrInference([d_494_v16_ for d_497_i5_ in range(default__.abs(-817))]))})
                d_498_v19_: _dafny.Map
                d_498_v19_ = _dafny.Map({p2: p2})
                d_499_v20_: _dafny.Map
                d_499_v20_ = _dafny.Map({d_498_v19_: d_494_v16_})
                d_500_v21_: _dafny.Array
                nw88_ = _dafny.Array(None, 15)
                nw88_[int(0)] = d_475_v3_
                nw88_[int(1)] = d_475_v3_
                nw88_[int(2)] = d_475_v3_
                nw88_[int(3)] = d_475_v3_
                nw88_[int(4)] = (d_475_v3_ if not((self).fm3(d_475_v3_, p1, p2, globalState)) else d_475_v3_)
                nw88_[int(5)] = d_475_v3_
                nw88_[int(6)] = _dafny.SeqWithoutIsStrInference([len(d_495_v17_), (d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))]])
                nw88_[int(7)] = _dafny.SeqWithoutIsStrInference([(d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))] for d_501_i3_ in range(default__.abs(435))])
                nw88_[int(8)] = (_dafny.SeqWithoutIsStrInference([782 for d_502_i4_ in range(default__.abs(817))])) + (d_475_v3_)
                nw88_[int(9)] = d_475_v3_
                nw88_[int(10)] = (default__.fm1((d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))], _dafny.CodePoint('j'), p2, d_496_v18_, globalState)) + (_dafny.SeqWithoutIsStrInference([(d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))], (d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))]]))
                nw88_[int(11)] = d_475_v3_
                nw88_[int(12)] = _dafny.SeqWithoutIsStrInference([(0) - ((d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))]), len(d_499_v20_)])
                nw88_[int(13)] = (d_475_v3_) + (d_475_v3_)
                nw88_[int(14)] = d_475_v3_
                d_500_v21_ = nw88_
                index86_ = default__.safeIndex(361, (d_500_v21_).length(0))
                (d_500_v21_)[index86_] = _dafny.SeqWithoutIsStrInference([(d_471_v1_)[default__.safeIndex(958, (d_471_v1_).length(0))] for d_503_i6_ in range(default__.abs(868))])
                index87_ = default__.safeIndex(361, (d_500_v21_).length(0))
                (d_500_v21_)[index87_] = d_475_v3_
                d_504_v22_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_504_v22_ = nw89_
                index88_ = default__.safeIndex(639, (d_504_v22_).length(0))
                (d_504_v22_)[index88_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fusvcouy"))
                index89_ = default__.safeIndex(639, (d_504_v22_).length(0))
                (d_504_v22_)[index89_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
        elif True:
            d_505_v23_: _dafny.Array
            def lambda29_(d_506_p0_):
                def lambda30_(d_507_i7_):
                    return (d_507_i7_) - (d_506_p0_)

                return lambda30_

            init14_ = lambda29_(p0)
            nw90_ = _dafny.Array(None, 2)
            for i0_14_ in range(nw90_.length(0)):
                nw90_[i0_14_] = init14_(i0_14_)
            d_505_v23_ = nw90_
            d_508_v24_: D0
            d_508_v24_ = D0_DC3(D0_DC1((0) - (p0), p0))
            d_509_v25_: _dafny.Seq
            d_509_v25_ = _dafny.SeqWithoutIsStrInference([d_508_v24_])
            index90_ = default__.safeIndex(119, (d_505_v23_).length(0))
            (d_505_v23_)[index90_] = len(d_509_v25_)
            d_510_v26_: _dafny.Seq
            d_510_v26_ = _dafny.SeqWithoutIsStrInference([p2, (p2) or (True), True])
            index91_ = default__.safeIndex(119, (d_505_v23_).length(0))
            rhs80_ = ((p0 if p2 else p0)) + (p0)
            rhs81_ = len(d_510_v26_)
            lhs65_ = globalState
            lhs66_ = d_505_v23_
            lhs67_ = default__.safeIndex(119, (d_505_v23_).length(0))
            lhs65_.f3 = rhs80_
            lhs66_[lhs67_] = rhs81_
            (globalState).f4 = p2
            d_511_v27_: _dafny.Set
            d_511_v27_ = _dafny.Set({p2})
            d_512_v28_: D9
            d_512_v28_ = D9_DC20(d_511_v27_)
            d_513_v29_: _dafny.Seq
            d_513_v29_ = _dafny.SeqWithoutIsStrInference([d_512_v28_])
            rhs82_ = (0) - ((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))])
            rhs83_ = d_513_v29_
            lhs68_ = globalState
            lhs68_.f8 = rhs82_
            d_513_v29_ = rhs83_
            d_514_v30_: _dafny.Seq
            d_514_v30_ = _dafny.SeqWithoutIsStrInference([(d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], p0, (d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], p0, p0])
            d_515_v31_: _dafny.Map
            d_515_v31_ = _dafny.Map({len(_dafny.Set({p0})): p0})
            d_516_v32_: D3
            d_516_v32_ = D3_DC8(d_515_v31_)
            d_517_v33_: _dafny.MultiSet
            d_517_v33_ = _dafny.MultiSet([d_510_v26_])
            d_518_v34_: _dafny.Array
            def lambda31_(d_519_i8_):
                return _dafny.CodePoint('t')

            init15_ = lambda31_
            nw91_ = _dafny.Array(None, 9)
            for i0_15_ in range(nw91_.length(0)):
                nw91_[i0_15_] = init15_(i0_15_)
            d_518_v34_ = nw91_
            d_520_v35_: _dafny.Map
            d_520_v35_ = _dafny.Map({d_518_v34_: -607})
            if (self).fm3(d_514_v30_, (p1) | (_dafny.MultiSet([p0, default__.fm26(d_516_v32_, default__.fm26(d_516_v32_, (d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_510_v26_])), globalState), d_517_v33_, globalState), len(d_520_v35_), p0, p0])), p2, globalState):
                (globalState).f4 = p2
                d_521_v36_: _dafny.Map
                d_521_v36_ = _dafny.Map({p2: (d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))]})
                (globalState).f3 = ((d_521_v36_)[p2] if (p2) in (d_521_v36_) else p0)
                d_470_v0_ = d_470_v0_
                d_470_v0_ = d_470_v0_
                d_522_v37_: _dafny.Set
                d_522_v37_ = _dafny.Set({(p1).cardinality})
                def iife96_():
                    coll62_ = _dafny.Set()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(331, 658):
                        d_523_v38_: int = compr_62_
                        if ((331) <= (d_523_v38_)) and ((d_523_v38_) < (658)):
                            coll62_ = coll62_.union(_dafny.Set([(d_523_v38_) - (len(_dafny.Map({len(d_470_v0_): p0})))]))
                    return _dafny.Set(coll62_)
                d_522_v37_ = (_dafny.Set({(d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], len(_dafny.Map({p0: len(iife96_()
                )})), (d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))]})) | (d_522_v37_)
            elif True:
                d_508_v24_ = d_508_v24_
                d_524_v39_: _dafny.Map
                d_524_v39_ = _dafny.Map({469: _dafny.SeqWithoutIsStrInference([p0])})
                d_525_v40_: _dafny.Array
                nw92_ = _dafny.Array(None, 9)
                nw92_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))]])
                nw92_[int(1)] = d_514_v30_
                nw92_[int(2)] = d_514_v30_
                nw92_[int(3)] = (d_514_v30_).set(default__.safeIndex((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], len(d_514_v30_)), 645)
                nw92_[int(4)] = d_514_v30_
                nw92_[int(5)] = (((d_524_v39_)[182] if (182) in (d_524_v39_) else d_514_v30_)) + ((d_514_v30_).set(default__.safeIndex(p0, len(d_514_v30_)), len(d_470_v0_)))
                nw92_[int(6)] = d_514_v30_
                nw92_[int(7)] = (d_514_v30_) + (d_514_v30_)
                nw92_[int(8)] = d_514_v30_
                d_525_v40_ = nw92_
                index92_ = default__.safeIndex(744, (d_525_v40_).length(0))
                (d_525_v40_)[index92_] = d_514_v30_
                d_526_v41_: str
                d_526_v41_ = _dafny.CodePoint('n')
                d_527_v42_: _dafny.Array
                nw93_ = _dafny.Array(None, 4)
                nw93_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eypc"))
                nw93_[int(1)] = (((d_470_v0_).set(default__.safeIndex(p0, len(d_470_v0_)), d_526_v41_)).set(default__.safeIndex((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], len((d_470_v0_).set(default__.safeIndex(p0, len(d_470_v0_)), d_526_v41_))), d_526_v41_)) + (d_470_v0_)
                nw93_[int(2)] = (d_470_v0_) + (d_470_v0_)
                nw93_[int(3)] = (d_470_v0_) + (d_470_v0_)
                d_527_v42_ = nw93_
                d_528_v43_: D2
                d_528_v43_ = D2_DC7(d_526_v41_, p2, (d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))])
                index93_ = default__.safeIndex(744, (d_525_v40_).length(0))
                index94_ = default__.safeIndex(119, (d_505_v23_).length(0))
                rhs84_ = d_514_v30_
                rhs85_ = default__.safeDivisionInt((0) - (p0), ((0) - ((0) - ((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))]))) + (p0))
                rhs86_ = d_527_v42_
                rhs87_ = ((d_528_v43_).cf11) not in (d_515_v31_)
                lhs69_ = d_525_v40_
                lhs70_ = default__.safeIndex(744, (d_525_v40_).length(0))
                lhs71_ = d_505_v23_
                lhs72_ = default__.safeIndex(119, (d_505_v23_).length(0))
                lhs73_ = globalState
                lhs69_[lhs70_] = rhs84_
                lhs71_[lhs72_] = rhs85_
                d_527_v42_ = rhs86_
                lhs73_.f4 = rhs87_
                d_529_v44_: bool
                out16_: bool
                out16_ = (self).m14(p2, (_dafny.Map({p2: p2})) == (_dafny.Map({True: p2})), p2, len(d_514_v30_), globalState)
                d_529_v44_ = out16_
                d_530_v45_: _dafny.Seq
                d_530_v45_ = _dafny.SeqWithoutIsStrInference([d_510_v26_])
                d_531_v46_: D5
                d_531_v46_ = D5_DC14(default__.fm26(d_516_v32_, (d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_510_v26_])), globalState), d_529_v44_)
                d_532_v47_: _dafny.Map
                d_532_v47_ = _dafny.Map({d_531_v46_: p0})
                d_533_v48_: _dafny.Map
                d_533_v48_ = _dafny.Map({d_530_v45_: d_532_v47_})
                d_533_v48_ = (d_533_v48_).set(d_530_v45_, d_532_v47_)
                (globalState).f4 = (628) != (default__.fm26(d_516_v32_, ((d_525_v40_)[default__.safeIndex(744, (d_525_v40_).length(0))])[default__.safeIndex((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))], len((d_525_v40_)[default__.safeIndex(744, (d_525_v40_).length(0))]))], d_517_v33_, globalState))
            d_534_v49_: D0
            d_534_v49_ = D0_DC0((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))])
            index95_ = default__.safeIndex(119, (d_505_v23_).length(0))
            rhs88_ = p0
            rhs89_ = len((d_470_v0_) + (d_470_v0_))
            rhs90_ = p2
            rhs91_ = (D0_DC0((d_505_v23_)[default__.safeIndex(119, (d_505_v23_).length(0))]) if p2 else d_534_v49_)
            rhs92_ = (d_514_v30_) + ((d_514_v30_) + (d_514_v30_))
            lhs74_ = d_505_v23_
            lhs75_ = default__.safeIndex(119, (d_505_v23_).length(0))
            lhs76_ = globalState
            lhs77_ = globalState
            lhs78_ = globalState
            lhs74_[lhs75_] = rhs88_
            lhs76_.f8 = rhs89_
            lhs77_.f4 = rhs90_
            d_534_v49_ = rhs91_
            lhs78_.f0 = rhs92_
        d_535_v50_: _dafny.Map
        d_535_v50_ = _dafny.Map({p0: p2})
        d_535_v50_ = (d_535_v50_).set(p0, (p2 if p2 else p2))
        (globalState).f4 = p2
        d_536_v51_: _dafny.Seq
        d_536_v51_ = _dafny.SeqWithoutIsStrInference([72, p0, p0, p0])
        d_537_v52_: _dafny.Map
        d_537_v52_ = _dafny.Map({(d_536_v51_).set(default__.safeIndex(p0, len(d_536_v51_)), p0): p2})
        d_538_v53_: _dafny.Map
        d_538_v53_ = _dafny.Map({d_537_v52_: p2})
        if (((d_538_v53_)[d_537_v52_] if (d_537_v52_) in (d_538_v53_) else p2)) and (p2):
            d_539_v54_: C0
            nw94_ = C0()
            nw94_.ctor__()
            d_539_v54_ = nw94_
            (globalState).f8 = p0
            if (self).fm3(d_536_v51_, p1, p2, globalState):
                d_540_v55_: C0
                nw95_ = C0()
                nw95_.ctor__()
                d_540_v55_ = nw95_
                rhs93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "meyetabe"))
                rhs94_ = (0) - (p0)
                lhs79_ = globalState
                d_470_v0_ = rhs93_
                lhs79_.f3 = rhs94_
                d_541_v56_: C0
                nw96_ = C0()
                nw96_.ctor__()
                d_541_v56_ = nw96_
                (globalState).f4 = (self).fm3(d_536_v51_, p1, (p2 if p2 else p2), globalState)
                d_542_v57_: _dafny.Array
                def lambda32_(d_543_p2_):
                    def lambda33_(d_544_i9_):
                        return d_543_p2_

                    return lambda33_

                init16_ = lambda32_(p2)
                nw97_ = _dafny.Array(None, 3)
                for i0_16_ in range(nw97_.length(0)):
                    nw97_[i0_16_] = init16_(i0_16_)
                d_542_v57_ = nw97_
                d_545_v58_: D5
                d_545_v58_ = D5_DC14(p0, p2)
                index96_ = default__.safeIndex(263, (d_542_v57_).length(0))
                (d_542_v57_)[index96_] = ((d_545_v58_).cf21) == (p2)
                index97_ = default__.safeIndex(986, (d_542_v57_).length(0))
                (d_542_v57_)[index97_] = p2
                d_546_v59_: _dafny.Map
                d_546_v59_ = _dafny.Map({default__.fm24(globalState): _dafny.MultiSet([p2, p2, p2, p2])})
                d_547_v60_: _dafny.MultiSet
                d_547_v60_ = _dafny.MultiSet([p2])
                d_548_v61_: _dafny.Seq
                d_548_v61_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p2, p2]), d_547_v60_])
                d_549_v62_: _dafny.MultiSet
                d_549_v62_ = d_547_v60_
                d_550_v63_: _dafny.Map
                d_550_v63_ = _dafny.Map({p0: d_547_v60_})
                d_551_v64_: _dafny.Array
                nw98_ = _dafny.Array(None, 27)
                nw98_[int(0)] = ((d_547_v60_).set(p2, default__.abs(p0))).intersection(default__.fm0(p0, globalState))
                nw98_[int(1)] = (d_548_v61_)[default__.safeIndex(p0, len(d_548_v61_))]
                nw98_[int(2)] = (d_547_v60_) - (_dafny.MultiSet([True, p2, p2, p2]))
                nw98_[int(3)] = (d_547_v60_ if p2 else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2, not(p2)])))
                nw98_[int(4)] = d_547_v60_
                nw98_[int(5)] = (d_549_v62_)
                nw98_[int(6)] = _dafny.MultiSet([p2])
                nw98_[int(7)] = (_dafny.MultiSet([p2, p2])) - (d_547_v60_)
                nw98_[int(8)] = (d_547_v60_).intersection(d_547_v60_)
                nw98_[int(9)] = d_547_v60_
                nw98_[int(10)] = d_547_v60_
                nw98_[int(11)] = d_547_v60_
                nw98_[int(12)] = d_547_v60_
                nw98_[int(13)] = ((d_548_v61_)[default__.safeIndex(len(d_536_v51_), len(d_548_v61_))] if p2 else d_547_v60_)
                nw98_[int(14)] = d_547_v60_
                nw98_[int(15)] = d_547_v60_
                nw98_[int(16)] = d_547_v60_
                nw98_[int(17)] = (d_547_v60_).intersection(d_547_v60_)
                nw98_[int(18)] = d_547_v60_
                nw98_[int(19)] = d_547_v60_
                nw98_[int(20)] = (d_547_v60_).set(p2, default__.abs(p0))
                nw98_[int(21)] = d_547_v60_
                nw98_[int(22)] = ((d_550_v63_)[p0] if (p0) in (d_550_v63_) else d_547_v60_)
                nw98_[int(23)] = d_547_v60_
                nw98_[int(24)] = d_547_v60_
                nw98_[int(25)] = (d_547_v60_).intersection(d_547_v60_)
                nw98_[int(26)] = (d_549_v62_)
                d_551_v64_ = nw98_
                index98_ = default__.safeIndex(392, (d_551_v64_).length(0))
                (d_551_v64_)[index98_] = default__.fm0(len(d_470_v0_), globalState)
                index99_ = default__.safeIndex(263, (d_542_v57_).length(0))
                index100_ = default__.safeIndex(986, (d_542_v57_).length(0))
                index101_ = default__.safeIndex(392, (d_551_v64_).length(0))
                rhs95_ = default__.fm27(True, globalState)
                rhs96_ = p2
                rhs97_ = not (p2) or (p2)
                rhs98_ = d_546_v59_
                rhs99_ = d_547_v60_
                lhs80_ = globalState
                lhs81_ = d_542_v57_
                lhs82_ = default__.safeIndex(263, (d_542_v57_).length(0))
                lhs83_ = d_542_v57_
                lhs84_ = default__.safeIndex(986, (d_542_v57_).length(0))
                lhs85_ = d_551_v64_
                lhs86_ = default__.safeIndex(392, (d_551_v64_).length(0))
                lhs80_.f2 = rhs95_
                lhs81_[lhs82_] = rhs96_
                lhs83_[lhs84_] = rhs97_
                d_546_v59_ = rhs98_
                lhs85_[lhs86_] = rhs99_
            elif True:
                d_552_v65_: _dafny.Array
                def lambda34_(d_553_p0_):
                    def lambda35_(d_554_i10_):
                        return (d_553_p0_) != (d_553_p0_)

                    return lambda35_

                init17_ = lambda34_(p0)
                nw99_ = _dafny.Array(None, 3)
                for i0_17_ in range(nw99_.length(0)):
                    nw99_[i0_17_] = init17_(i0_17_)
                d_552_v65_ = nw99_
                d_555_v66_: str
                d_555_v66_ = _dafny.CodePoint('h')
                index102_ = default__.safeIndex(578, (d_552_v65_).length(0))
                (d_552_v65_)[index102_] = (d_555_v66_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijcbji")))
                index103_ = default__.safeIndex(578, (d_552_v65_).length(0))
                (d_552_v65_)[index103_] = p2
                d_556_v67_: _dafny.MultiSet
                d_556_v67_ = _dafny.MultiSet([(d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))], (d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))], p2, (d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))]])
                d_557_v68_: D2
                d_557_v68_ = D2_DC6(p0, (d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))], p2)
                d_558_v69_: _dafny.Seq
                d_558_v69_ = _dafny.SeqWithoutIsStrInference([p2])
                d_559_v70_: _dafny.Array
                nw100_ = _dafny.Array(None, 20)
                nw100_[int(0)] = d_556_v67_
                nw100_[int(1)] = d_556_v67_
                nw100_[int(2)] = (_dafny.MultiSet([(d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))]])) - (d_556_v67_)
                nw100_[int(3)] = d_556_v67_
                nw100_[int(4)] = (_dafny.MultiSet([(d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))]])) | (_dafny.MultiSet([(d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))]]))
                nw100_[int(5)] = d_556_v67_
                nw100_[int(6)] = _dafny.MultiSet([p2, (d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))], False, not(p2)])
                nw100_[int(7)] = d_556_v67_
                nw100_[int(8)] = d_556_v67_
                nw100_[int(9)] = (d_556_v67_ if (d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))] else d_556_v67_)
                nw100_[int(10)] = _dafny.MultiSet([(d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))]])
                nw100_[int(11)] = (d_556_v67_) - (_dafny.MultiSet([(d_557_v68_).cf7, p2, (d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))]]))
                nw100_[int(12)] = _dafny.MultiSet([p2])
                nw100_[int(13)] = d_556_v67_
                nw100_[int(14)] = _dafny.MultiSet([(D9_DC21(not(not((d_552_v65_)[default__.safeIndex(578, (d_552_v65_).length(0))])), d_558_v69_, not(p2))).cf28])
                nw100_[int(15)] = d_556_v67_
                nw100_[int(16)] = d_556_v67_
                nw100_[int(17)] = _dafny.MultiSet(d_558_v69_)
                nw100_[int(18)] = d_556_v67_
                nw100_[int(19)] = d_556_v67_
                d_559_v70_ = nw100_
                index104_ = default__.safeIndex(293, (d_559_v70_).length(0))
                (d_559_v70_)[index104_] = d_556_v67_
                index105_ = default__.safeIndex(293, (d_559_v70_).length(0))
                (d_559_v70_)[index105_] = (d_556_v67_) | (_dafny.MultiSet([p2]))
                d_560_v71_: _dafny.Array
                nw101_ = _dafny.Array(int(0), 3)
                d_560_v71_ = nw101_
                index106_ = default__.safeIndex(169, (d_560_v71_).length(0))
                (d_560_v71_)[index106_] = p0
                d_561_v72_: _dafny.Map
                d_561_v72_ = _dafny.Map({-798: p0})
                index107_ = default__.safeIndex(169, (d_560_v71_).length(0))
                (d_560_v71_)[index107_] = (len((d_561_v72_) | (d_561_v72_))) + (p0)
                (globalState).f8 = (p0) * (975)
                d_470_v0_ = d_470_v0_
            d_539_v54_ = d_539_v54_
            d_562_v73_: _dafny.Map
            d_562_v73_ = _dafny.Map({p2: len(d_470_v0_)})
            def iife97_():
                coll63_ = _dafny.Map()
                compr_63_: int
                for compr_63_ in _dafny.IntegerRange(958, -919):
                    d_563_v74_: int = compr_63_
                    if ((958) <= (d_563_v74_)) and ((d_563_v74_) < (-919)):
                        coll63_[(d_563_v74_) + (p0)] = p2
                return _dafny.Map(coll63_)
            (globalState).f3 = (len((D12_DC26(d_562_v73_)).cf35)) + (len((d_535_v50_) | (iife97_()
            )))
        elif True:
            d_564_v75_: _dafny.Set
            d_564_v75_ = _dafny.Set({p0, (0) - (p0)})
            if (d_564_v75_).ispropersubset((d_564_v75_).intersection(default__.fm28(p0, globalState))):
                (globalState).f3 = (0) - (p0)
                d_565_v76_: _dafny.Set
                d_565_v76_ = _dafny.Set({p2, p2, p2})
                d_566_v77_: _dafny.MultiSet
                d_566_v77_ = _dafny.MultiSet([p2, True, p2, p2])
                d_567_v78_: _dafny.Seq
                d_567_v78_ = _dafny.SeqWithoutIsStrInference([d_536_v51_])
                d_568_v79_: _dafny.Map
                d_568_v79_ = _dafny.Map({p2: not(p2)})
                d_569_v80_: _dafny.Array
                nw102_ = _dafny.Array(None, 16)
                nw102_[int(0)] = (p0) > ((0) - (p0))
                nw102_[int(1)] = p2
                nw102_[int(2)] = (d_565_v76_).issubset(d_565_v76_)
                nw102_[int(3)] = p2
                nw102_[int(4)] = p2
                nw102_[int(5)] = (d_566_v77_).isdisjoint(d_566_v77_)
                nw102_[int(6)] = p2
                nw102_[int(7)] = True
                nw102_[int(8)] = (self).fm3(d_536_v51_, _dafny.MultiSet((d_567_v78_)[default__.safeIndex(p0, len(d_567_v78_))]), p2, globalState)
                nw102_[int(9)] = p2
                nw102_[int(10)] = p2
                nw102_[int(11)] = (p1).isdisjoint(p1)
                nw102_[int(12)] = ((d_568_v79_)[p2] if (p2) in (d_568_v79_) else p2)
                nw102_[int(13)] = p2
                nw102_[int(14)] = False
                nw102_[int(15)] = not(not (p2) or (p2))
                d_569_v80_ = nw102_
                index108_ = default__.safeIndex(463, (d_569_v80_).length(0))
                (d_569_v80_)[index108_] = p2
                index109_ = default__.safeIndex(463, (d_569_v80_).length(0))
                (d_569_v80_)[index109_] = not(False)
                d_570_v81_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.Map({}), 23)
                d_570_v81_ = nw103_
                index110_ = default__.safeIndex(820, (d_570_v81_).length(0))
                (d_570_v81_)[index110_] = d_568_v79_
                index111_ = default__.safeIndex(820, (d_570_v81_).length(0))
                (d_570_v81_)[index111_] = d_568_v79_
                d_571_v82_: _dafny.Array
                nw104_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_571_v82_ = nw104_
                d_572_v83_: str
                d_572_v83_ = _dafny.CodePoint('x')
                index112_ = default__.safeIndex(110, (d_571_v82_).length(0))
                (d_571_v82_)[index112_] = d_572_v83_
                index113_ = default__.safeIndex(110, (d_571_v82_).length(0))
                (d_571_v82_)[index113_] = d_572_v83_
                (globalState).f4 = ((d_566_v77_) | ((_dafny.MultiSet([p2])))) == (d_566_v77_)
            elif True:
                (globalState).f4 = (len(d_564_v75_)) not in (d_535_v50_)
                d_573_v84_: _dafny.MultiSet
                d_573_v84_ = _dafny.MultiSet([not (p2) or (p2), p2])
                d_573_v84_ = d_573_v84_
                d_574_v85_: str
                d_574_v85_ = _dafny.CodePoint('v')
                d_574_v85_ = d_574_v85_
                (globalState).f3 = p0
                d_575_v86_: _dafny.Array
                def lambda36_(d_576_p0_):
                    def lambda37_(d_577_i11_):
                        return default__.safeDivisionInt(d_577_i11_, d_576_p0_)

                    return lambda37_

                init18_ = lambda36_(p0)
                nw105_ = _dafny.Array(None, 16)
                for i0_18_ in range(nw105_.length(0)):
                    nw105_[i0_18_] = init18_(i0_18_)
                d_575_v86_ = nw105_
                index114_ = default__.safeIndex(852, (d_575_v86_).length(0))
                (d_575_v86_)[index114_] = (-359) + (p0)
                d_578_v87_: _dafny.Seq
                d_578_v87_ = _dafny.SeqWithoutIsStrInference([not(not(True)), p2, p2, p2, p2])
                d_579_v88_: _dafny.MultiSet
                d_579_v88_ = _dafny.MultiSet([d_578_v87_, d_578_v87_])
                d_580_v89_: _dafny.Map
                d_580_v89_ = _dafny.Map({p2: p2})
                d_581_v90_: _dafny.Map
                d_581_v90_ = _dafny.Map({len(d_580_v89_): d_578_v87_})
                index115_ = default__.safeIndex(852, (d_575_v86_).length(0))
                (d_575_v86_)[index115_] = ((d_579_v88_)[((d_581_v90_)[p0] if (p0) in (d_581_v90_) else d_578_v87_)] if (((d_581_v90_)[p0] if (p0) in (d_581_v90_) else d_578_v87_)) in (d_579_v88_) else p0)
            (globalState).f4 = p2
            (globalState).f0 = (d_536_v51_) + (d_536_v51_)
            (globalState).f4 = ((p2 if p2 else p2)) or (p2)
            d_582_v91_: _dafny.Array
            nw106_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_582_v91_ = nw106_
            index116_ = default__.safeIndex(24, (d_582_v91_).length(0))
            (d_582_v91_)[index116_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_583_i12_ in range(default__.abs(832))])
            d_584_v92_: _dafny.Array
            def lambda38_(d_585_p0_):
                def lambda39_(d_586_i13_):
                    return (d_585_p0_) == (d_585_p0_)

                return lambda39_

            init19_ = lambda38_(p0)
            nw107_ = _dafny.Array(None, 27)
            for i0_19_ in range(nw107_.length(0)):
                nw107_[i0_19_] = init19_(i0_19_)
            d_584_v92_ = nw107_
            index117_ = default__.safeIndex(212, (d_584_v92_).length(0))
            (d_584_v92_)[index117_] = True
            index118_ = default__.safeIndex(24, (d_582_v91_).length(0))
            index119_ = default__.safeIndex(212, (d_584_v92_).length(0))
            rhs100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "miiuxq"))
            rhs101_ = d_584_v92_
            rhs102_ = False
            lhs87_ = d_582_v91_
            lhs88_ = default__.safeIndex(24, (d_582_v91_).length(0))
            lhs89_ = d_584_v92_
            lhs90_ = default__.safeIndex(212, (d_584_v92_).length(0))
            lhs87_[lhs88_] = rhs100_
            d_584_v92_ = rhs101_
            lhs89_[lhs90_] = rhs102_
        d_587_i14_: int
        d_587_i14_ = 0
        with _dafny.label("4"):
            while (d_470_v0_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_597_i15_ in range(default__.abs(66))])):
                with _dafny.c_label("4"):
                    if (d_587_i14_) >= (100):
                        raise _dafny.Break("4")
                    d_587_i14_ = (d_587_i14_) + (1)
                    d_588_v93_: _dafny.Seq
                    d_588_v93_ = _dafny.SeqWithoutIsStrInference([not(True)])
                    d_589_v94_: _dafny.MultiSet
                    d_589_v94_ = _dafny.MultiSet([(d_588_v93_)[default__.safeIndex(813, len(d_588_v93_))], False])
                    d_590_v95_: _dafny.MultiSet
                    d_590_v95_ = d_589_v94_
                    d_591_v96_: _dafny.Array
                    nw108_ = _dafny.Array(None, 13)
                    nw108_[int(0)] = d_590_v95_
                    nw108_[int(1)] = d_589_v94_
                    nw108_[int(2)] = d_590_v95_
                    nw108_[int(3)] = default__.fm29(p0, p0, p2, p0, globalState)
                    nw108_[int(4)] = (d_590_v95_ if not(p2) else d_590_v95_)
                    nw108_[int(5)] = d_590_v95_
                    nw108_[int(6)] = d_590_v95_
                    nw108_[int(7)] = d_590_v95_
                    nw108_[int(8)] = d_590_v95_
                    nw108_[int(9)] = d_589_v94_
                    nw108_[int(10)] = d_590_v95_
                    nw108_[int(11)] = d_590_v95_
                    nw108_[int(12)] = _dafny.MultiSet([p2])
                    d_591_v96_ = nw108_
                    index120_ = default__.safeIndex(541, (d_591_v96_).length(0))
                    (d_591_v96_)[index120_] = d_590_v95_
                    index121_ = default__.safeIndex(541, (d_591_v96_).length(0))
                    (d_591_v96_)[index121_] = d_590_v95_
                    d_592_v97_: _dafny.Seq
                    d_592_v97_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_593_v98_: _dafny.Array
                    nw109_ = _dafny.Array(None, 13)
                    nw109_[int(0)] = p2
                    nw109_[int(1)] = p2
                    nw109_[int(2)] = p2
                    nw109_[int(3)] = p2
                    nw109_[int(4)] = (d_588_v93_)[default__.safeIndex(len(d_535_v50_), len(d_588_v93_))]
                    nw109_[int(5)] = p2
                    nw109_[int(6)] = (p1).issubset((d_592_v97_)[default__.safeIndex(p0, len(d_592_v97_))])
                    nw109_[int(7)] = p2
                    nw109_[int(8)] = p2
                    nw109_[int(9)] = p2
                    nw109_[int(10)] = p2
                    nw109_[int(11)] = p2
                    nw109_[int(12)] = True
                    d_593_v98_ = nw109_
                    rhs103_ = p2
                    rhs104_ = d_536_v51_
                    rhs105_ = d_593_v98_
                    lhs91_ = globalState
                    lhs92_ = globalState
                    lhs91_.f4 = rhs103_
                    lhs92_.f0 = rhs104_
                    d_593_v98_ = rhs105_
                    d_594_v99_: _dafny.Map
                    d_594_v99_ = _dafny.Map({p2: d_536_v51_})
                    d_595_v100_: _dafny.Map
                    d_595_v100_ = _dafny.Map({p0: (((d_594_v99_)[p2] if (p2) in (d_594_v99_) else d_536_v51_)).set(default__.safeIndex(p0, len(((d_594_v99_)[p2] if (p2) in (d_594_v99_) else d_536_v51_))), 288)})
                    (globalState).f0 = ((d_595_v100_)[p0] if (p0) in (d_595_v100_) else _dafny.SeqWithoutIsStrInference([p0 for d_596_i16_ in range(default__.abs(622))]))
                    (globalState).f4 = (p2 if p2 else p2)
                    pass
            pass
        d_598_v101_: _dafny.Array
        nw110_ = _dafny.Array(None, 5)
        nw110_[int(0)] = ((0) - (p0) if p2 else p0)
        nw110_[int(1)] = p0
        nw110_[int(2)] = (len((_dafny.Map({p2: p2})).set(p2, p2))) + (720)
        nw110_[int(3)] = p0
        nw110_[int(4)] = p0
        d_598_v101_ = nw110_
        d_598_v101_ = d_598_v101_

    def m14(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        (globalState).f8 = (p3) - (p3)
        hi3_ = p3
        for d_599_i0_ in range(p3, hi3_):
            d_600_v0_: C0
            nw111_ = C0()
            nw111_.ctor__()
            d_600_v0_ = nw111_
            (globalState).f4 = p0
            (globalState).f4 = p2
            (globalState).f3 = (0) - (d_599_i0_)
        if p1:
            d_601_v1_: _dafny.Array
            def lambda40_(d_602_p3_):
                def lambda41_(d_603_i1_):
                    return default__.safeDivisionInt(d_603_i1_, d_602_p3_)

                return lambda41_

            init20_ = lambda40_(p3)
            nw112_ = _dafny.Array(None, 1)
            for i0_20_ in range(nw112_.length(0)):
                nw112_[i0_20_] = init20_(i0_20_)
            d_601_v1_ = nw112_
            index122_ = default__.safeIndex(394, (d_601_v1_).length(0))
            (d_601_v1_)[index122_] = 199
            index123_ = default__.safeIndex(656, (d_601_v1_).length(0))
            (d_601_v1_)[index123_] = p3
            d_604_v2_: _dafny.Set
            d_604_v2_ = _dafny.Set({p3})
            d_605_v3_: _dafny.Seq
            d_605_v3_ = _dafny.SeqWithoutIsStrInference([d_604_v2_, d_604_v2_])
            d_606_v4_: _dafny.Seq
            d_606_v4_ = _dafny.SeqWithoutIsStrInference([p0, not(p1)])
            d_607_v5_: _dafny.Seq
            d_607_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auprui"))
            d_608_v6_: _dafny.Map
            d_608_v6_ = _dafny.Map({len(d_607_v5_): p3})
            d_609_v7_: D3
            d_609_v7_ = D3_DC8(d_608_v6_)
            d_610_v8_: _dafny.Map
            d_610_v8_ = _dafny.Map({d_606_v4_: d_609_v7_})
            d_611_v9_: D4
            d_611_v9_ = D4_DC10(d_606_v4_)
            d_612_v10_: D4
            d_612_v10_ = D4_DC12(d_611_v9_)
            d_613_v11_: _dafny.Map
            d_613_v11_ = _dafny.Map({d_612_v10_: d_604_v2_})
            d_614_v12_: D9
            d_614_v12_ = D9_DC21(p1, d_606_v4_, not(p0))
            d_615_v13_: _dafny.Seq
            d_615_v13_ = _dafny.SeqWithoutIsStrInference([p3, p3])
            d_616_v14_: _dafny.MultiSet
            d_616_v14_ = _dafny.MultiSet([p3, p3, p3, p3, p3])
            d_617_v15_: _dafny.Seq
            d_617_v15_ = _dafny.SeqWithoutIsStrInference([default__.fm31(p1, p1, p1, globalState)])
            pat_let_tv20_ = d_607_v5_
            pat_let_tv21_ = d_614_v12_
            pat_let_tv22_ = p1
            pat_let_tv23_ = d_607_v5_
            pat_let_tv24_ = globalState
            pat_let_tv25_ = d_607_v5_
            pat_let_tv26_ = d_614_v12_
            pat_let_tv27_ = p1
            pat_let_tv28_ = d_607_v5_
            pat_let_tv29_ = globalState
            d_618_v16_: _dafny.Array
            nw113_ = _dafny.Array(None, 29)
            def iife98_(_pat_let17_0):
                def iife99_(d_619_dt__update__tmp_h1_):
                    def iife100_(_pat_let18_0):
                        def iife101_(d_620_dt__update_hcf18_h1_):
                            return D4_DC12(d_620_dt__update_hcf18_h1_)
                        return iife101_(_pat_let18_0)
                    return iife100_(default__.fm30(len(pat_let_tv20_), pat_let_tv21_, pat_let_tv22_, len(pat_let_tv23_), pat_let_tv24_))
                return iife99_(_pat_let17_0)
            def iife102_(_pat_let19_0):
                def iife103_(d_621_dt__update__tmp_h0_):
                    def iife104_(_pat_let20_0):
                        def iife105_(d_622_dt__update_hcf18_h0_):
                            return D4_DC12(d_622_dt__update_hcf18_h0_)
                        return iife105_(_pat_let20_0)
                    return iife104_(default__.fm30(len(pat_let_tv25_), pat_let_tv26_, pat_let_tv27_, len(pat_let_tv28_), pat_let_tv29_))
                return iife103_(_pat_let19_0)
            nw113_[int(0)] = ((d_605_v3_)[default__.safeIndex(len(d_610_v8_), len(d_605_v3_))]).issubset(((d_613_v11_)[iife98_(D4_DC12(d_611_v9_))] if (iife102_(D4_DC12(d_611_v9_))) in (d_613_v11_) else d_604_v2_))
            nw113_[int(1)] = p0
            nw113_[int(2)] = p0
            nw113_[int(3)] = p0
            nw113_[int(4)] = p0
            nw113_[int(5)] = p2
            nw113_[int(6)] = not(p1)
            nw113_[int(7)] = p0
            nw113_[int(8)] = p1
            nw113_[int(9)] = p2
            nw113_[int(10)] = not(p1)
            nw113_[int(11)] = not(False)
            nw113_[int(12)] = (self).fm3(d_615_v13_, d_616_v14_, (self).fm3(d_615_v13_, d_616_v14_, p2, globalState), globalState)
            nw113_[int(13)] = p1
            nw113_[int(14)] = p1
            nw113_[int(15)] = p2
            nw113_[int(16)] = (p1) == (p1)
            nw113_[int(17)] = not(p1)
            nw113_[int(18)] = p1
            nw113_[int(19)] = (p3) == (len(d_617_v15_))
            nw113_[int(20)] = p1
            nw113_[int(21)] = p2
            nw113_[int(22)] = p1
            nw113_[int(23)] = (p1) and (p1)
            nw113_[int(24)] = p0
            nw113_[int(25)] = not(p2)
            nw113_[int(26)] = p0
            nw113_[int(27)] = (p3) in (d_615_v13_)
            nw113_[int(28)] = p2
            d_618_v16_ = nw113_
            index124_ = default__.safeIndex(355, (d_618_v16_).length(0))
            (d_618_v16_)[index124_] = (d_606_v4_)[default__.safeIndex(p3, len(d_606_v4_))]
            index125_ = default__.safeIndex(394, (d_601_v1_).length(0))
            index126_ = default__.safeIndex(656, (d_601_v1_).length(0))
            index127_ = default__.safeIndex(355, (d_618_v16_).length(0))
            rhs106_ = p3
            rhs107_ = default__.safeDivisionInt(len(_dafny.Set({p3, p3})), p3)
            rhs108_ = p3
            rhs109_ = not(p2)
            rhs110_ = p1
            lhs93_ = globalState
            lhs94_ = d_601_v1_
            lhs95_ = default__.safeIndex(394, (d_601_v1_).length(0))
            lhs96_ = d_601_v1_
            lhs97_ = default__.safeIndex(656, (d_601_v1_).length(0))
            lhs98_ = d_618_v16_
            lhs99_ = default__.safeIndex(355, (d_618_v16_).length(0))
            lhs100_ = globalState
            lhs93_.f8 = rhs106_
            lhs94_[lhs95_] = rhs107_
            lhs96_[lhs97_] = rhs108_
            lhs98_[lhs99_] = rhs109_
            lhs100_.f4 = rhs110_
            d_623_v17_: _dafny.Seq
            d_623_v17_ = _dafny.SeqWithoutIsStrInference([d_607_v5_, d_607_v5_, d_607_v5_, d_607_v5_, d_607_v5_])
            d_624_v18_: str
            d_624_v18_ = _dafny.CodePoint('y')
            d_625_v19_: _dafny.Map
            d_625_v19_ = _dafny.Map({p1: _dafny.Map({not(p0): (d_601_v1_)[default__.safeIndex(394, (d_601_v1_).length(0))]})})
            index128_ = default__.safeIndex(355, (d_618_v16_).length(0))
            (d_618_v16_)[index128_] = (default__.fm32(p3, d_623_v17_, d_624_v18_, p0, globalState)) != (d_625_v19_)
            d_626_v20_: _dafny.Set
            d_626_v20_ = _dafny.Set({p2})
            if ((d_626_v20_).ispropersubset(d_626_v20_)) or ((p2) in (d_626_v20_)):
                d_607_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mosxmc"))
                index129_ = default__.safeIndex(394, (d_601_v1_).length(0))
                (d_601_v1_)[index129_] = (0) - (default__.safeDivisionInt((0) - (default__.safeDivisionInt(p3, p3)), 837))
                d_627_v21_: _dafny.Map
                d_627_v21_ = _dafny.Map({d_618_v16_: d_624_v18_})
                (globalState).f8 = len(d_627_v21_)
                d_628_v22_: _dafny.Map
                d_628_v22_ = _dafny.Map({len(d_623_v17_): p1})
                d_629_v25_: _dafny.Set
                def iife106_():
                    coll64_ = _dafny.Set()
                    compr_64_: int
                    for compr_64_ in (d_628_v22_).keys.Elements:
                        d_630_v23_: int = compr_64_
                        if (d_630_v23_) in (d_628_v22_):
                            coll64_ = coll64_.union(_dafny.Set([default__.safeModuloInt(d_630_v23_, len(_dafny.Set({False, not(False), True, False, not(True)})))]))
                    return _dafny.Set(coll64_)
                def iife107_():
                    coll65_ = _dafny.Set()
                    compr_65_: int
                    for compr_65_ in _dafny.IntegerRange(280, 451):
                        d_631_v24_: int = compr_65_
                        if ((280) <= (d_631_v24_)) and ((d_631_v24_) < (451)):
                            coll65_ = coll65_.union(_dafny.Set([(d_631_v24_) * (p3)]))
                    return _dafny.Set(coll65_)
                d_629_v25_ = _dafny.Set({d_604_v2_, iife106_()
                , iife107_()
                })
                d_629_v25_ = d_629_v25_
                d_632_v26_: _dafny.Array
                nw114_ = _dafny.Array(_dafny.Map({}), 15)
                d_632_v26_ = nw114_
                index130_ = default__.safeIndex(333, (d_632_v26_).length(0))
                (d_632_v26_)[index130_] = _dafny.Map({(self).fm3(d_615_v13_, d_616_v14_, not((d_618_v16_)[default__.safeIndex(355, (d_618_v16_).length(0))]), globalState): False})
                index131_ = default__.safeIndex(333, (d_632_v26_).length(0))
                (d_632_v26_)[index131_] = _dafny.Map({(340) not in (d_608_v6_): p2})
            elif True:
                index132_ = default__.safeIndex(394, (d_601_v1_).length(0))
                (d_601_v1_)[index132_] = p3
                d_633_v27_: _dafny.MultiSet
                d_633_v27_ = _dafny.MultiSet([d_606_v4_])
                d_634_v28_: _dafny.Seq
                d_634_v28_ = _dafny.SeqWithoutIsStrInference([d_633_v27_, d_633_v27_, d_633_v27_])
                d_635_v29_: _dafny.Map
                d_635_v29_ = _dafny.Map({default__.fm26(d_609_v7_, 390, (d_634_v28_)[default__.safeIndex((d_601_v1_)[default__.safeIndex(394, (d_601_v1_).length(0))], len(d_634_v28_))], globalState): p1})
                (globalState).f3 = ((0) - (len(d_635_v29_))) - (126)
                d_636_v30_: _dafny.Array
                nw115_ = _dafny.Array(_dafny.Map({}), 12)
                d_636_v30_ = nw115_
                d_637_v31_: _dafny.Map
                d_637_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([False, (d_618_v16_)[default__.safeIndex(355, (d_618_v16_).length(0))]]): p1})
                index133_ = default__.safeIndex(154, (d_636_v30_).length(0))
                (d_636_v30_)[index133_] = d_637_v31_
                index134_ = default__.safeIndex(154, (d_636_v30_).length(0))
                def iife108_():
                    coll66_ = _dafny.Map()
                    compr_66_: _dafny.Seq
                    for compr_66_ in (d_617_v15_).Elements:
                        d_638_v32_: _dafny.Seq = compr_66_
                        if (d_638_v32_) in (d_617_v15_):
                            coll66_[d_638_v32_] = p1
                    return _dafny.Map(coll66_)
                rhs111_ = 393
                rhs112_ = (d_604_v2_) - (_dafny.Set({len(d_607_v5_)}))
                rhs113_ = (default__.fm33(p1, p0, globalState)) | ((iife108_()
                ) | (d_637_v31_))
                rhs114_ = d_624_v18_
                lhs101_ = globalState
                lhs102_ = d_636_v30_
                lhs103_ = default__.safeIndex(154, (d_636_v30_).length(0))
                lhs101_.f3 = rhs111_
                d_604_v2_ = rhs112_
                lhs102_[lhs103_] = rhs113_
                d_624_v18_ = rhs114_
                d_639_v33_: _dafny.Map
                d_639_v33_ = _dafny.Map({d_618_v16_: (0) - ((d_601_v1_)[default__.safeIndex(394, (d_601_v1_).length(0))])})
                d_639_v33_ = (d_639_v33_).set(d_618_v16_, default__.safeModuloInt((d_601_v1_)[default__.safeIndex(394, (d_601_v1_).length(0))], p3))
                d_640_v34_: C0
                nw116_ = C0()
                nw116_.ctor__()
                d_640_v34_ = nw116_
            d_641_v35_: _dafny.Map
            d_641_v35_ = _dafny.Map({len(d_626_v20_): d_624_v18_})
            index135_ = default__.safeIndex(394, (d_601_v1_).length(0))
            (d_601_v1_)[index135_] = default__.safeDivisionInt(len(d_641_v35_), (d_601_v1_)[default__.safeIndex(394, (d_601_v1_).length(0))])
            d_642_v37_: _dafny.MultiSet
            d_642_v37_ = _dafny.MultiSet([p3, len(d_607_v5_)])
            d_643_v38_: _dafny.MultiSet
            d_643_v38_ = _dafny.MultiSet([d_606_v4_])
            pat_let_tv30_ = p2
            pat_let_tv31_ = globalState
            pat_let_tv32_ = p2
            pat_let_tv33_ = globalState
            d_644_v39_: _dafny.Map
            def iife109_():
                def iife114_(_pat_let23_0):
                    def iife115_(d_648_dt__update__tmp_h2_):
                        def iife116_(_pat_let24_0):
                            def iife117_(d_649_dt__update_hcf22_h0_):
                                return d_649_dt__update_hcf22_h0_
                            return iife117_(_pat_let24_0)
                        return iife116_(default__.fm27(pat_let_tv32_, pat_let_tv33_))
                    return iife115_(_pat_let23_0)
                coll67_ = _dafny.Map()
                def iife110_(_pat_let21_0):
                    def iife111_(d_645_dt__update__tmp_h2_):
                        def iife112_(_pat_let22_0):
                            def iife113_(d_646_dt__update_hcf22_h0_):
                                return d_646_dt__update_hcf22_h0_
                            return iife113_(_pat_let22_0)
                        return iife112_(default__.fm27(pat_let_tv30_, pat_let_tv31_))
                    return iife111_(_pat_let21_0)
                compr_67_: _dafny.MultiSet
                for compr_67_ in (_dafny.SeqWithoutIsStrInference([d_642_v37_, d_616_v14_, iife110_(d_642_v37_)])).Elements:
                    d_647_v36_: _dafny.MultiSet = compr_67_
                    if (d_647_v36_) in (_dafny.SeqWithoutIsStrInference([d_642_v37_, d_616_v14_, iife114_(d_642_v37_)])):
                        coll67_[d_647_v36_] = True
                return _dafny.Map(coll67_)
            d_644_v39_ = _dafny.Map({p0: default__.fm26(d_609_v7_, len(iife109_()
            ), d_643_v38_, globalState)})
            d_644_v39_ = (d_644_v39_).set(p2, default__.safeModuloInt(p3, (d_601_v1_)[default__.safeIndex(394, (d_601_v1_).length(0))]))
        elif True:
            d_650_v40_: _dafny.Seq
            d_650_v40_ = _dafny.SeqWithoutIsStrInference([p3])
            d_651_v41_: _dafny.MultiSet
            d_651_v41_ = _dafny.MultiSet([p3, 492, p3, p3, p3])
            r0 = (self).fm3(d_650_v40_, d_651_v41_, p1, globalState)
            d_652_v42_: _dafny.Set
            d_652_v42_ = _dafny.Set({p0, p2, p1, (p1) and (p1)})
            d_653_v43_: _dafny.Array
            nw117_ = _dafny.Array(_dafny.Seq({}), 16)
            d_653_v43_ = nw117_
            d_654_v44_: _dafny.Map
            d_654_v44_ = _dafny.Map({p0: p1})
            rhs115_ = (d_652_v42_).intersection(d_652_v42_)
            rhs116_ = d_653_v43_
            rhs117_ = ((d_651_v41_)[default__.safeModuloInt(len(d_654_v44_), (0) - ((0) - (p3)))] if (default__.safeModuloInt(len(d_654_v44_), (0) - ((0) - (p3)))) in (d_651_v41_) else len(_dafny.Map({p3: p3})))
            lhs104_ = globalState
            d_652_v42_ = rhs115_
            d_653_v43_ = rhs116_
            lhs104_.f3 = rhs117_
            (globalState).f8 = p3
            d_655_v45_: _dafny.Array
            nw118_ = _dafny.Array(_dafny.CodePoint('D'), 28)
            d_655_v45_ = nw118_
            index136_ = default__.safeIndex(777, (d_655_v45_).length(0))
            (d_655_v45_)[index136_] = _dafny.CodePoint('h')
            d_656_v46_: str
            d_656_v46_ = _dafny.CodePoint('s')
            index137_ = default__.safeIndex(777, (d_655_v45_).length(0))
            (d_655_v45_)[index137_] = d_656_v46_
            (globalState).f8 = p3
        r0 = p0
        d_657_v47_: _dafny.Seq
        d_657_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjj"))
        d_658_v48_: _dafny.Array
        nw119_ = _dafny.Array(None, 13)
        nw119_[int(0)] = False
        nw119_[int(1)] = p1
        nw119_[int(2)] = not(p0)
        nw119_[int(3)] = p0
        nw119_[int(4)] = (_dafny.CodePoint('x')) not in (d_657_v47_)
        nw119_[int(5)] = p1
        nw119_[int(6)] = p2
        nw119_[int(7)] = (p3) < ((default__.fm34(p0, p1, globalState)).cf0)
        nw119_[int(8)] = p2
        nw119_[int(9)] = p1
        nw119_[int(10)] = p0
        nw119_[int(11)] = p0
        nw119_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvo"))) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_659_i2_ in range(default__.abs(702))]))
        d_658_v48_ = nw119_
        index138_ = default__.safeIndex(184, (d_658_v48_).length(0))
        (d_658_v48_)[index138_] = p1
        index139_ = default__.safeIndex(184, (d_658_v48_).length(0))
        (d_658_v48_)[index139_] = p2
        d_660_v49_: _dafny.Seq
        d_660_v49_ = _dafny.SeqWithoutIsStrInference([(d_657_v47_) < (d_657_v47_), (d_658_v48_)[default__.safeIndex(184, (d_658_v48_).length(0))], (d_658_v48_)[default__.safeIndex(184, (d_658_v48_).length(0))], p2])
        d_661_v50_: str
        d_661_v50_ = _dafny.CodePoint('f')
        d_662_v51_: _dafny.MultiSet
        d_662_v51_ = _dafny.MultiSet([(0) - (p3)])
        d_663_v52_: _dafny.Seq
        d_663_v52_ = _dafny.SeqWithoutIsStrInference([(d_662_v51_).cardinality])
        d_664_v53_: D2
        d_664_v53_ = D2_DC7(d_661_v50_, p1, len(d_663_v52_))
        d_665_v54_: D12
        d_665_v54_ = D12_DC27(p3, p3, p3, p3)
        pat_let_tv34_ = d_665_v54_
        index140_ = default__.safeIndex(184, (d_658_v48_).length(0))
        def iife119_(_pat_let26_0):
            def iife120_(d_666_dt__update__tmp_h3_):
                def iife121_(_pat_let27_0):
                    def iife122_(d_667_dt__update_hcf11_h0_):
                        return D2_DC7((d_666_dt__update__tmp_h3_).cf9, (d_666_dt__update__tmp_h3_).cf10, d_667_dt__update_hcf11_h0_)
                    return iife122_(_pat_let27_0)
                return iife121_((pat_let_tv34_).cf37)
            return iife120_(_pat_let26_0)
        def iife118_(_pat_let25_0):
            def iife123_(d_668_dt__update__tmp_h4_):
                def iife124_(_pat_let28_0):
                    def iife125_(d_669_dt__update_hcf10_h0_):
                        return D2_DC7((d_668_dt__update__tmp_h4_).cf9, d_669_dt__update_hcf10_h0_, (d_668_dt__update__tmp_h4_).cf11)
                    return iife125_(_pat_let28_0)
                return iife124_(True)
            return iife123_(_pat_let25_0)
        (d_658_v48_)[index140_] = (d_660_v49_)[default__.safeIndex((iife118_(iife119_(d_664_v53_))).cf11, len(d_660_v49_))]
        r0 = p1
        return r0


class C2(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self):
        pass
        pass

    def fm2(self, globalState):
        def iife126_():
            coll68_ = _dafny.Map()
            compr_68_: int
            for compr_68_ in (_dafny.Map({920: False})).keys.Elements:
                d_670_v0_: int = compr_68_
                if (d_670_v0_) in (_dafny.Map({920: False})):
                    coll68_[(d_670_v0_) - (851)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jf")))
            return _dafny.Map(coll68_)
        return ((iife126_()
        ) | (_dafny.Map({441: 596}))) | (_dafny.Map({len(_dafny.Map({-204: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ee")))})): 986}))

    def fm3(self, p0, p1, p2, globalState):
        source13_ = D4_DC11(True, True, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqonprou")))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svtwbn")))
        if source13_.is_DC11:
            d_671___mcc_h0_ = source13_.cf14
            d_672___mcc_h1_ = source13_.cf15
            d_673___mcc_h2_ = source13_.cf16
            d_674___mcc_h3_ = source13_.cf17
            d_675_cf17_ = d_674___mcc_h3_
            d_676_cf16_ = d_673___mcc_h2_
            d_677_cf15_ = d_672___mcc_h1_
            d_678_cf14_ = d_671___mcc_h0_
            return False
        elif source13_.is_DC10:
            d_679___mcc_h4_ = source13_.cf13
            d_680_cf13_ = d_679___mcc_h4_
            return (378) != (708)
        elif True:
            d_681___mcc_h5_ = source13_.cf18
            d_682_cf18_ = d_681___mcc_h5_
            return False

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_683_v0_: int
        d_683_v0_ = 558
        (globalState).f3 = d_683_v0_
        d_684_v2_: _dafny.MultiSet
        d_684_v2_ = _dafny.MultiSet([d_683_v0_])
        def iife127_():
            coll69_ = _dafny.Map()
            compr_69_: int
            for compr_69_ in _dafny.IntegerRange(764, 518):
                d_685_v1_: int = compr_69_
                if ((764) <= (d_685_v1_)) and ((d_685_v1_) < (518)):
                    coll69_[(d_685_v1_) - (d_683_v0_)] = _dafny.CodePoint('u')
            return _dafny.Map(coll69_)
        if (_dafny.MultiSet([d_683_v0_, d_683_v0_, d_683_v0_])).ispropersubset((_dafny.MultiSet([len(iife127_()
        )])) | (d_684_v2_)):
            d_686_v3_: _dafny.Seq
            d_686_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxhq"))
            d_687_v4_: _dafny.Seq
            d_687_v4_ = _dafny.SeqWithoutIsStrInference([d_686_v3_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_688_i0_ in range(default__.abs(-395))])])
            d_689_v5_: _dafny.Map
            d_689_v5_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilkioe")): (d_687_v4_)[default__.safeIndex(746, len(d_687_v4_))]})
            d_690_v6_: bool
            d_690_v6_ = True
            d_686_v3_ = ((d_689_v5_)[d_686_v3_] if (d_686_v3_) in (d_689_v5_) else default__.fm17(default__.fm18(d_690_v6_, d_690_v6_, globalState), globalState))
            d_691_v7_: _dafny.Seq
            d_691_v7_ = _dafny.SeqWithoutIsStrInference([d_683_v0_, d_683_v0_])
            d_692_v8_: _dafny.Seq
            d_692_v8_ = _dafny.SeqWithoutIsStrInference([(self).fm3(d_691_v7_, d_684_v2_, d_690_v6_, globalState)])
            d_693_v9_: D2
            d_693_v9_ = D2_DC6((d_683_v0_ if (d_692_v8_)[default__.safeIndex(len(d_686_v3_), len(d_692_v8_))] else 446), d_690_v6_, d_690_v6_)
            d_693_v9_ = d_693_v9_
            d_694_v10_: _dafny.Set
            d_694_v10_ = _dafny.Set({d_690_v6_})
            (globalState).f8 = default__.safeDivisionInt((d_683_v0_) + (d_683_v0_), len(d_694_v10_))
            if d_690_v6_:
                d_695_v11_: C0
                nw120_ = C0()
                nw120_.ctor__()
                d_695_v11_ = nw120_
                d_696_v12_: C0
                nw121_ = C0()
                nw121_.ctor__()
                d_696_v12_ = nw121_
                d_697_v13_: _dafny.Array
                nw122_ = _dafny.Array(False, 14)
                d_697_v13_ = nw122_
                index141_ = default__.safeIndex(665, (d_697_v13_).length(0))
                (d_697_v13_)[index141_] = d_690_v6_
                d_698_v14_: str
                d_698_v14_ = _dafny.CodePoint('j')
                d_699_v15_: _dafny.Array
                nw123_ = _dafny.Array(_dafny.Seq({}), 22)
                d_699_v15_ = nw123_
                d_700_v16_: _dafny.Seq
                d_700_v16_ = _dafny.SeqWithoutIsStrInference([d_697_v13_, d_697_v13_])
                index142_ = default__.safeIndex(378, (d_699_v15_).length(0))
                (d_699_v15_)[index142_] = d_700_v16_
                index143_ = default__.safeIndex(665, (d_697_v13_).length(0))
                index144_ = default__.safeIndex(378, (d_699_v15_).length(0))
                rhs118_ = d_683_v0_
                rhs119_ = 635
                rhs120_ = d_690_v6_
                rhs121_ = d_698_v14_
                rhs122_ = d_700_v16_
                lhs105_ = globalState
                lhs106_ = globalState
                lhs107_ = d_697_v13_
                lhs108_ = default__.safeIndex(665, (d_697_v13_).length(0))
                lhs109_ = d_699_v15_
                lhs110_ = default__.safeIndex(378, (d_699_v15_).length(0))
                lhs105_.f8 = rhs118_
                lhs106_.f3 = rhs119_
                lhs107_[lhs108_] = rhs120_
                d_698_v14_ = rhs121_
                lhs109_[lhs110_] = rhs122_
                index145_ = default__.safeIndex(665, (d_697_v13_).length(0))
                (d_697_v13_)[index145_] = (d_697_v13_)[default__.safeIndex(665, (d_697_v13_).length(0))]
                d_701_v17_: C0
                nw124_ = C0()
                nw124_.ctor__()
                d_701_v17_ = nw124_
            elif True:
                d_702_v18_: str
                d_702_v18_ = _dafny.CodePoint('a')
                d_686_v3_ = ((d_686_v3_).set(default__.safeIndex(d_683_v0_, len(d_686_v3_)), d_702_v18_)).set(default__.safeIndex(default__.fm20(not(False), globalState), len((d_686_v3_).set(default__.safeIndex(d_683_v0_, len(d_686_v3_)), d_702_v18_))), d_702_v18_)
                d_703_v19_: _dafny.Map
                d_703_v19_ = _dafny.Map({True: d_690_v6_})
                d_704_v20_: _dafny.Map
                d_704_v20_ = _dafny.Map({d_690_v6_: _dafny.Map({d_690_v6_: d_690_v6_})})
                d_705_v21_: D5
                d_705_v21_ = D5_DC13(_dafny.Map({d_690_v6_: True}))
                d_706_v22_: _dafny.Seq
                d_706_v22_ = _dafny.SeqWithoutIsStrInference([d_703_v19_])
                d_707_v23_: _dafny.Array
                nw125_ = _dafny.Array(None, 14)
                nw125_[int(0)] = d_703_v19_
                nw125_[int(1)] = d_703_v19_
                nw125_[int(2)] = ((d_704_v20_)[d_690_v6_] if (d_690_v6_) in (d_704_v20_) else _dafny.Map({d_690_v6_: d_690_v6_}))
                nw125_[int(3)] = d_703_v19_
                nw125_[int(4)] = _dafny.Map({d_690_v6_: d_690_v6_})
                nw125_[int(5)] = d_703_v19_
                nw125_[int(6)] = d_703_v19_
                nw125_[int(7)] = (d_705_v21_).cf19
                nw125_[int(8)] = d_703_v19_
                nw125_[int(9)] = default__.fm21(d_690_v6_, globalState)
                nw125_[int(10)] = d_703_v19_
                nw125_[int(11)] = (d_706_v22_)[default__.safeIndex(848, len(d_706_v22_))]
                nw125_[int(12)] = d_703_v19_
                nw125_[int(13)] = d_703_v19_
                d_707_v23_ = nw125_
                d_708_v24_: _dafny.Map
                d_708_v24_ = _dafny.Map({d_707_v23_: d_690_v6_})
                d_708_v24_ = (d_708_v24_).set(d_707_v23_, (d_683_v0_) == (len(d_686_v3_)))
                d_683_v0_ = d_683_v0_
                (globalState).f3 = d_683_v0_
                d_694_v10_ = d_694_v10_
            r0 = not(not((len(d_686_v3_)) <= (d_683_v0_)))
        elif True:
            d_709_v25_: bool
            d_709_v25_ = False
            (globalState).f3 = (d_683_v0_ if d_709_v25_ else d_683_v0_)
            d_710_v26_: _dafny.Array
            nw126_ = _dafny.Array(False, 20)
            d_710_v26_ = nw126_
            index146_ = default__.safeIndex(319, (d_710_v26_).length(0))
            (d_710_v26_)[index146_] = d_709_v25_
            index147_ = default__.safeIndex(319, (d_710_v26_).length(0))
            (d_710_v26_)[index147_] = d_709_v25_
            d_711_v27_: _dafny.Seq
            d_711_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijeb"))
            d_712_v28_: _dafny.Seq
            d_712_v28_ = _dafny.SeqWithoutIsStrInference([-273])
            d_713_v29_: _dafny.Seq
            d_713_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_711_v27_), d_683_v0_]), (d_712_v28_ if (d_710_v26_)[default__.safeIndex(319, (d_710_v26_).length(0))] else d_712_v28_)])
            (globalState).f0 = (d_713_v29_)[default__.safeIndex(d_683_v0_, len(d_713_v29_))]
            d_683_v0_ = d_683_v0_
            d_714_v30_: _dafny.Array
            nw127_ = _dafny.Array(_dafny.Seq({}), 27)
            d_714_v30_ = nw127_
            d_715_v31_: _dafny.Map
            d_715_v31_ = _dafny.Map({d_683_v0_: d_683_v0_})
            d_716_v32_: _dafny.Array
            nw128_ = _dafny.Array(None, 11)
            nw128_[int(0)] = d_683_v0_
            nw128_[int(1)] = d_683_v0_
            nw128_[int(2)] = len((_dafny.SeqWithoutIsStrInference([(d_710_v26_)[default__.safeIndex(319, (d_710_v26_).length(0))], d_709_v25_])) + (_dafny.SeqWithoutIsStrInference([(d_710_v26_)[default__.safeIndex(319, (d_710_v26_).length(0))]])))
            nw128_[int(3)] = default__.safeDivisionInt(d_683_v0_, default__.fm20((d_710_v26_)[default__.safeIndex(319, (d_710_v26_).length(0))], globalState))
            nw128_[int(4)] = default__.safeModuloInt(d_683_v0_, d_683_v0_)
            nw128_[int(5)] = d_683_v0_
            nw128_[int(6)] = (len(d_715_v31_)) + (523)
            nw128_[int(7)] = d_683_v0_
            nw128_[int(8)] = (0) - (default__.safeModuloInt(d_683_v0_, d_683_v0_))
            nw128_[int(9)] = default__.safeModuloInt(d_683_v0_, d_683_v0_)
            nw128_[int(10)] = d_683_v0_
            d_716_v32_ = nw128_
            index148_ = default__.safeIndex(813, (d_716_v32_).length(0))
            (d_716_v32_)[index148_] = default__.fm20((d_710_v26_)[default__.safeIndex(319, (d_710_v26_).length(0))], globalState)
            index149_ = default__.safeIndex(813, (d_716_v32_).length(0))
            rhs123_ = d_714_v30_
            rhs124_ = default__.safeDivisionInt(d_683_v0_, d_683_v0_)
            rhs125_ = d_683_v0_
            rhs126_ = d_709_v25_
            rhs127_ = (d_710_v26_ if (d_683_v0_) < (d_683_v0_) else d_710_v26_)
            lhs111_ = globalState
            lhs112_ = d_716_v32_
            lhs113_ = default__.safeIndex(813, (d_716_v32_).length(0))
            d_714_v30_ = rhs123_
            lhs111_.f3 = rhs124_
            lhs112_[lhs113_] = rhs125_
            d_709_v25_ = rhs126_
            d_710_v26_ = rhs127_
        d_717_v33_: bool
        d_717_v33_ = False
        d_718_v34_: _dafny.MultiSet
        d_718_v34_ = d_684_v2_
        d_719_v35_: _dafny.Map
        d_719_v35_ = _dafny.Map({d_683_v0_: _dafny.MultiSet([d_683_v0_, d_683_v0_])})
        d_720_v36_: _dafny.Map
        d_720_v36_ = _dafny.Map({False: d_717_v33_})
        d_721_v37_: _dafny.Seq
        d_721_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xebnstm"))
        d_722_v38_: D4
        d_722_v38_ = D4_DC11(d_717_v33_, not(True), d_683_v0_, d_721_v37_)
        d_723_v39_: _dafny.Seq
        d_723_v39_ = _dafny.SeqWithoutIsStrInference([not(d_717_v33_)])
        d_724_v40_: _dafny.Array
        nw129_ = _dafny.Array(None, 25)
        nw129_[int(0)] = not(d_717_v33_)
        nw129_[int(1)] = d_717_v33_
        nw129_[int(2)] = d_717_v33_
        nw129_[int(3)] = d_717_v33_
        nw129_[int(4)] = d_717_v33_
        nw129_[int(5)] = not(True)
        nw129_[int(6)] = ((d_718_v34_)).isdisjoint(((d_719_v35_)[679] if (679) in (d_719_v35_) else (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_725_i2_ in range(default__.abs(402))])) for d_726_i1_ in range(default__.abs(-318))]))).set(len(d_720_v36_), default__.abs(d_683_v0_))))
        nw129_[int(7)] = False
        nw129_[int(8)] = d_717_v33_
        nw129_[int(9)] = (d_722_v38_).cf14
        nw129_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_717_v33_, d_717_v33_])) <= (d_723_v39_)
        nw129_[int(11)] = d_717_v33_
        nw129_[int(12)] = d_717_v33_
        nw129_[int(13)] = d_717_v33_
        nw129_[int(14)] = d_717_v33_
        nw129_[int(15)] = d_717_v33_
        nw129_[int(16)] = d_717_v33_
        nw129_[int(17)] = not(d_717_v33_)
        nw129_[int(18)] = d_717_v33_
        nw129_[int(19)] = d_717_v33_
        nw129_[int(20)] = d_717_v33_
        nw129_[int(21)] = (d_683_v0_) == (d_683_v0_)
        nw129_[int(22)] = True
        nw129_[int(23)] = d_717_v33_
        nw129_[int(24)] = d_717_v33_
        d_724_v40_ = nw129_
        index150_ = default__.safeIndex(696, (d_724_v40_).length(0))
        (d_724_v40_)[index150_] = False
        d_727_v41_: _dafny.Seq
        d_727_v41_ = _dafny.SeqWithoutIsStrInference([d_683_v0_])
        index151_ = default__.safeIndex(696, (d_724_v40_).length(0))
        (d_724_v40_)[index151_] = ((len(d_727_v41_)) - (d_683_v0_)) >= (d_683_v0_)
        d_728_v42_: _dafny.Array
        def lambda42_(d_729_v0_):
            def lambda43_(d_730_i3_):
                return default__.safeDivisionInt(d_730_i3_, (0) - (d_729_v0_))

            return lambda43_

        init21_ = lambda42_(d_683_v0_)
        nw130_ = _dafny.Array(None, 19)
        for i0_21_ in range(nw130_.length(0)):
            nw130_[i0_21_] = init21_(i0_21_)
        d_728_v42_ = nw130_
        index152_ = default__.safeIndex(923, (d_728_v42_).length(0))
        (d_728_v42_)[index152_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))
        index153_ = default__.safeIndex(923, (d_728_v42_).length(0))
        (d_728_v42_)[index153_] = d_683_v0_
        d_731_i4_: int
        d_731_i4_ = 0
        with _dafny.label("5"):
            while (d_724_v40_)[default__.safeIndex(696, (d_724_v40_).length(0))]:
                with _dafny.c_label("5"):
                    if (d_731_i4_) >= (100):
                        raise _dafny.Break("5")
                    d_731_i4_ = (d_731_i4_) + (1)
                    r0 = d_717_v33_
                    (globalState).f8 = (d_683_v0_) - ((d_728_v42_)[default__.safeIndex(923, (d_728_v42_).length(0))])
                    d_732_v43_: _dafny.Map
                    d_732_v43_ = _dafny.Map({d_728_v42_: d_727_v41_})
                    d_733_v44_: _dafny.MultiSet
                    d_733_v44_ = _dafny.MultiSet([(d_724_v40_)[default__.safeIndex(696, (d_724_v40_).length(0))], d_717_v33_, d_717_v33_])
                    d_727_v41_ = ((d_732_v43_)[d_728_v42_] if (d_728_v42_) in (d_732_v43_) else ((d_727_v41_) + (d_727_v41_)).set(default__.safeIndex(((d_733_v44_)[(d_724_v40_)[default__.safeIndex(696, (d_724_v40_).length(0))]] if ((d_724_v40_)[default__.safeIndex(696, (d_724_v40_).length(0))]) in (d_733_v44_) else 690), len((d_727_v41_) + (d_727_v41_))), (d_728_v42_)[default__.safeIndex(923, (d_728_v42_).length(0))]))
                    index154_ = default__.safeIndex(923, (d_728_v42_).length(0))
                    (d_728_v42_)[index154_] = ((d_728_v42_)[default__.safeIndex(923, (d_728_v42_).length(0))]) * ((d_728_v42_)[default__.safeIndex(923, (d_728_v42_).length(0))])
                    pass
            pass
        d_734_v45_: str
        d_734_v45_ = _dafny.CodePoint('h')
        d_735_v46_: _dafny.Map
        d_735_v46_ = _dafny.Map({d_683_v0_: d_728_v42_})
        index155_ = default__.safeIndex(923, (d_728_v42_).length(0))
        rhs128_ = ((d_735_v46_)[d_683_v0_] if (d_683_v0_) in (d_735_v46_) else d_728_v42_)
        rhs129_ = _dafny.CodePoint('b')
        rhs130_ = d_734_v45_
        rhs131_ = ((d_728_v42_)[default__.safeIndex(923, (d_728_v42_).length(0))]) * (d_683_v0_)
        lhs114_ = d_728_v42_
        lhs115_ = default__.safeIndex(923, (d_728_v42_).length(0))
        d_728_v42_ = rhs128_
        d_734_v45_ = rhs129_
        d_734_v45_ = rhs130_
        lhs114_[lhs115_] = rhs131_
        r0 = d_717_v33_
        pat_let_tv35_ = d_734_v45_
        pat_let_tv36_ = d_717_v33_
        pat_let_tv37_ = d_724_v40_
        pat_let_tv38_ = d_724_v40_
        def lambda44_(source14_):
            if source14_.is_DC11:
                d_736___mcc_h0_ = source14_.cf14
                d_737___mcc_h1_ = source14_.cf15
                d_738___mcc_h2_ = source14_.cf16
                d_739___mcc_h3_ = source14_.cf17
                d_740_cf17_ = d_739___mcc_h3_
                d_741_cf16_ = d_738___mcc_h2_
                d_742_cf15_ = d_737___mcc_h1_
                d_743_cf14_ = d_736___mcc_h0_
                return (d_740_cf17_) < ((d_740_cf17_).set(default__.safeIndex(d_741_cf16_, len(d_740_cf17_)), pat_let_tv35_))
            elif source14_.is_DC10:
                d_744___mcc_h4_ = source14_.cf13
                d_745_cf13_ = d_744___mcc_h4_
                return pat_let_tv36_
            elif True:
                d_746___mcc_h5_ = source14_.cf18
                d_747_cf18_ = d_746___mcc_h5_
                return (pat_let_tv38_)[default__.safeIndex(696, (pat_let_tv37_).length(0))]

        r1 = lambda44_(d_722_v38_)
        return r0, r1

    def m11(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        d_748_v0_: bool
        d_748_v0_ = True
        (globalState).f4 = not(d_748_v0_)
        d_749_v1_: int
        d_749_v1_ = 159
        rhs132_ = d_749_v1_
        rhs133_ = d_749_v1_
        lhs116_ = globalState
        lhs116_.f8 = rhs132_
        r2 = rhs133_
        if d_748_v0_:
            d_750_v2_: _dafny.Set
            d_750_v2_ = _dafny.Set({d_748_v0_})
            d_750_v2_ = (d_750_v2_) - ((_dafny.Set({d_748_v0_}) if not(d_748_v0_) else d_750_v2_))
            d_751_v3_: _dafny.Seq
            d_751_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cooouputi"))
            d_752_v4_: _dafny.MultiSet
            d_752_v4_ = _dafny.MultiSet([d_749_v1_])
            d_753_v5_: _dafny.Map
            d_753_v5_ = _dafny.Map({d_748_v0_: (d_752_v4_).cardinality})
            d_754_v6_: _dafny.MultiSet
            d_754_v6_ = _dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ligquqw"))) + (d_751_v3_)), (d_749_v1_) * (((d_752_v4_)[d_749_v1_] if (d_749_v1_) in (d_752_v4_) else len(d_753_v5_))), default__.safeDivisionInt(d_749_v1_, d_749_v1_)])
            (globalState).f2 = d_752_v4_
            d_755_v7_: str
            d_755_v7_ = _dafny.CodePoint('n')
            d_756_v8_: _dafny.Array
            nw131_ = _dafny.Array(None, 9)
            nw131_[int(0)] = d_755_v7_
            nw131_[int(1)] = (d_755_v7_ if d_748_v0_ else d_755_v7_)
            nw131_[int(2)] = d_755_v7_
            nw131_[int(3)] = d_755_v7_
            nw131_[int(4)] = d_755_v7_
            nw131_[int(5)] = d_755_v7_
            nw131_[int(6)] = default__.fm18(d_748_v0_, d_748_v0_, globalState)
            nw131_[int(7)] = d_755_v7_
            nw131_[int(8)] = (_dafny.CodePoint('t') if d_748_v0_ else d_755_v7_)
            d_756_v8_ = nw131_
            index156_ = default__.safeIndex(437, (d_756_v8_).length(0))
            (d_756_v8_)[index156_] = _dafny.CodePoint('o')
            index157_ = default__.safeIndex(437, (d_756_v8_).length(0))
            (d_756_v8_)[index157_] = d_755_v7_
            (globalState).f4 = (d_749_v1_) == (d_749_v1_)
            d_757_v9_: _dafny.Array
            nw132_ = _dafny.Array(int(0), 26)
            d_757_v9_ = nw132_
            index158_ = default__.safeIndex(815, (d_757_v9_).length(0))
            (d_757_v9_)[index158_] = d_749_v1_
            index159_ = default__.safeIndex(815, (d_757_v9_).length(0))
            (d_757_v9_)[index159_] = default__.safeModuloInt(d_749_v1_, 217)
        elif True:
            d_758_v10_: _dafny.Map
            d_758_v10_ = _dafny.Map({346: d_749_v1_})
            d_759_v11_: D3
            d_759_v11_ = D3_DC8(d_758_v10_)
            d_760_v12_: _dafny.Seq
            d_760_v12_ = _dafny.SeqWithoutIsStrInference([(d_759_v11_ if d_748_v0_ else d_759_v11_), d_759_v11_])
            d_760_v12_ = (d_760_v12_) + (d_760_v12_)
            d_761_v13_: _dafny.Array
            def lambda45_(d_762_v1_):
                def lambda46_(d_763_i0_):
                    return _dafny.SeqWithoutIsStrInference([d_762_v1_, d_762_v1_])

                return lambda46_

            init22_ = lambda45_(d_749_v1_)
            nw133_ = _dafny.Array(None, 16)
            for i0_22_ in range(nw133_.length(0)):
                nw133_[i0_22_] = init22_(i0_22_)
            d_761_v13_ = nw133_
            d_764_v14_: _dafny.Seq
            d_764_v14_ = _dafny.SeqWithoutIsStrInference([d_749_v1_])
            index160_ = default__.safeIndex(828, (d_761_v13_).length(0))
            (d_761_v13_)[index160_] = (d_764_v14_) + (d_764_v14_)
            d_765_v15_: _dafny.MultiSet
            d_765_v15_ = _dafny.MultiSet([d_748_v0_])
            index161_ = default__.safeIndex(828, (d_761_v13_).length(0))
            rhs134_ = _dafny.SeqWithoutIsStrInference([d_749_v1_, d_749_v1_, d_749_v1_, ((d_765_v15_)[d_748_v0_] if (d_748_v0_) in (d_765_v15_) else d_749_v1_)])
            rhs135_ = (d_748_v0_) and (d_748_v0_)
            lhs117_ = d_761_v13_
            lhs118_ = default__.safeIndex(828, (d_761_v13_).length(0))
            lhs117_[lhs118_] = rhs134_
            d_748_v0_ = rhs135_
            d_766_v16_: _dafny.Map
            d_766_v16_ = _dafny.Map({d_748_v0_: d_749_v1_})
            d_766_v16_ = (d_766_v16_).set(d_748_v0_, len((p0) | (p0)))
            d_767_v17_: _dafny.Map
            d_767_v17_ = _dafny.Map({d_749_v1_: (self).fm3((d_761_v13_)[default__.safeIndex(828, (d_761_v13_).length(0))], (_dafny.MultiSet([d_749_v1_])).set(d_749_v1_, default__.abs(len((d_761_v13_)[default__.safeIndex(828, (d_761_v13_).length(0))]))), d_748_v0_, globalState)})
            d_768_v18_: D7
            d_768_v18_ = D7_DC17(_dafny.Map({d_749_v1_: d_748_v0_}))
            d_767_v17_ = (d_768_v18_).cf23
            d_769_v19_: _dafny.Seq
            d_769_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rd"))
            d_770_v20_: D2
            d_770_v20_ = D2_DC6((d_749_v1_) + (len(d_769_v19_)), not(d_748_v0_), not((d_749_v1_) < (d_749_v1_)))
            source15_ = d_770_v20_
            if source15_.is_DC6:
                d_771___mcc_h0_ = source15_.cf6
                d_772___mcc_h1_ = source15_.cf7
                d_773___mcc_h2_ = source15_.cf8
                d_774_cf8_ = d_773___mcc_h2_
                d_775_cf7_ = d_772___mcc_h1_
                d_776_cf6_ = d_771___mcc_h0_
                d_777_v21_: _dafny.Seq
                d_777_v21_ = _dafny.SeqWithoutIsStrInference([d_748_v0_, d_748_v0_, True, d_748_v0_, d_774_cf8_])
                d_778_v22_: _dafny.MultiSet
                d_778_v22_ = _dafny.MultiSet([(d_777_v21_) + (d_777_v21_)])
                d_778_v22_ = (d_778_v22_).set(d_777_v21_, default__.abs(d_749_v1_))
                d_779_v23_: _dafny.Array
                nw134_ = _dafny.Array(D5.default()(), 12)
                d_779_v23_ = nw134_
                index162_ = default__.safeIndex(885, (d_779_v23_).length(0))
                (d_779_v23_)[index162_] = D5_DC15()
                d_780_v24_: D5
                d_780_v24_ = D5_DC15()
                index163_ = default__.safeIndex(885, (d_779_v23_).length(0))
                (d_779_v23_)[index163_] = d_780_v24_
                d_781_v25_: int
                d_782_v26_: int
                d_783_v27_: int
                d_784_v28_: bool
                out17_: int
                out18_: int
                out19_: int
                out20_: bool
                out17_, out18_, out19_, out20_ = (self).m12(globalState)
                d_781_v25_ = out17_
                d_782_v26_ = out18_
                d_783_v27_ = out19_
                d_784_v28_ = out20_
                (globalState).f8 = d_782_v26_
            elif source15_.is_DC7:
                d_785___mcc_h3_ = source15_.cf9
                d_786___mcc_h4_ = source15_.cf10
                d_787___mcc_h5_ = source15_.cf11
                d_788_cf11_ = d_787___mcc_h5_
                d_789_cf10_ = d_786___mcc_h4_
                d_790_cf9_ = d_785___mcc_h3_
                d_791_v29_: _dafny.MultiSet
                d_791_v29_ = _dafny.MultiSet([730])
                d_792_v30_: _dafny.Map
                d_792_v30_ = _dafny.Map({(d_791_v29_).cardinality: d_790_cf9_})
                d_793_v31_: _dafny.MultiSet
                d_793_v31_ = _dafny.MultiSet([770, len(d_792_v30_)])
                d_794_v32_: _dafny.Map
                d_794_v32_ = _dafny.Map({not(d_748_v0_): d_748_v0_})
                d_795_v33_: _dafny.Seq
                d_795_v33_ = _dafny.SeqWithoutIsStrInference([d_793_v31_, default__.fm22(d_749_v1_, globalState)])
                d_796_v34_: _dafny.Seq
                d_796_v34_ = _dafny.SeqWithoutIsStrInference([((d_791_v29_).set(510, default__.abs(len(d_794_v32_)))).ispropersubset((d_795_v33_)[default__.safeIndex(len(d_769_v19_), len(d_795_v33_))])])
                d_797_v36_: _dafny.Map
                def iife128_():
                    coll70_ = _dafny.Map()
                    compr_70_: int
                    for compr_70_ in (d_758_v10_).keys.Elements:
                        d_798_v35_: int = compr_70_
                        if (d_798_v35_) in (d_758_v10_):
                            coll70_[(d_798_v35_) - (d_749_v1_)] = False
                    return _dafny.Map(coll70_)
                d_797_v36_ = _dafny.Map({d_749_v1_: (iife128_()
                ) | ((_dafny.Map({d_788_cf11_: d_789_cf10_})).set(d_749_v1_, False))})
                rhs136_ = default__.fm23(d_769_v19_, globalState)
                rhs137_ = d_748_v0_
                rhs138_ = (_dafny.Map({d_788_cf11_: d_767_v17_})) | (d_797_v36_)
                lhs119_ = globalState
                d_796_v34_ = rhs136_
                lhs119_.f4 = rhs137_
                d_797_v36_ = rhs138_
                d_790_cf9_ = _dafny.CodePoint('a')
                d_769_v19_ = default__.fm17(d_790_cf9_, globalState)
                pat_let_tv39_ = d_748_v0_
                d_799_v37_: _dafny.Map
                def iife129_(_pat_let29_0):
                    def iife130_(d_801_dt__update__tmp_h0_):
                        def iife131_(_pat_let30_0):
                            def iife132_(d_802_dt__update_hcf7_h0_):
                                def iife133_(_pat_let31_0):
                                    def iife134_(d_803_dt__update_hcf6_h0_):
                                        return D2_DC6(d_803_dt__update_hcf6_h0_, d_802_dt__update_hcf7_h0_, (d_801_dt__update__tmp_h0_).cf8)
                                    return iife134_(_pat_let31_0)
                                return iife133_(455)
                            return iife132_(_pat_let30_0)
                        return iife131_(pat_let_tv39_)
                    return iife130_(_pat_let29_0)
                d_799_v37_ = _dafny.Map({d_788_cf11_: ((_dafny.SeqWithoutIsStrInference([d_749_v1_ for d_800_i1_ in range(default__.abs(-851))])).set(default__.safeIndex((iife129_(d_770_v20_)).cf6, len(_dafny.SeqWithoutIsStrInference([d_749_v1_ for d_804_i1_ in range(default__.abs(-851))]))), d_749_v1_)) + (d_764_v14_)})
                d_799_v37_ = (d_799_v37_).set((0) - ((d_749_v1_) + (d_749_v1_)), (d_761_v13_)[default__.safeIndex(828, (d_761_v13_).length(0))])
            elif True:
                d_805___mcc_h6_ = source15_.cf5
                d_806_cf5_ = d_805___mcc_h6_
                d_807_v38_: _dafny.MultiSet
                d_807_v38_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_808_i2_ in range(default__.abs(-644))]))]))])
                d_809_v39_: str
                d_809_v39_ = _dafny.CodePoint('h')
                d_810_v40_: _dafny.Set
                d_810_v40_ = _dafny.Set({_dafny.CodePoint('y'), d_809_v39_})
                d_811_v41_: _dafny.Map
                d_811_v41_ = _dafny.Map({d_807_v38_: (len(d_810_v40_)) + (d_749_v1_)})
                d_811_v41_ = (d_811_v41_).set((default__.fm22(d_749_v1_, globalState)).set(d_749_v1_, default__.abs(d_749_v1_)), 850)
                (globalState).f8 = (d_749_v1_) + (d_749_v1_)
                index164_ = default__.safeIndex(338, (d_806_cf5_).length(0))
                (d_806_cf5_)[index164_] = d_748_v0_
                index165_ = default__.safeIndex(338, (d_806_cf5_).length(0))
                (d_806_cf5_)[index165_] = True
                rhs139_ = not((d_806_cf5_)[default__.safeIndex(338, (d_806_cf5_).length(0))])
                rhs140_ = d_748_v0_
                rhs141_ = ((d_761_v13_)[default__.safeIndex(828, (d_761_v13_).length(0))]) + (d_764_v14_)
                rhs142_ = d_748_v0_
                lhs120_ = globalState
                lhs121_ = globalState
                r1 = rhs139_
                lhs120_.f4 = rhs140_
                lhs121_.f0 = rhs141_
                d_748_v0_ = rhs142_
        (globalState).f4 = True
        d_812_v42_: _dafny.Array
        nw135_ = _dafny.Array(int(0), 11)
        d_812_v42_ = nw135_
        d_813_v43_: _dafny.Set
        d_813_v43_ = _dafny.Set({d_812_v42_, d_812_v42_, d_812_v42_, d_812_v42_})
        d_814_i3_: int
        d_814_i3_ = 0
        with _dafny.label("6"):
            while (d_812_v42_) not in ((d_813_v43_) - (d_813_v43_)):
                with _dafny.c_label("6"):
                    if (d_814_i3_) >= (100):
                        raise _dafny.Break("6")
                    d_814_i3_ = (d_814_i3_) + (1)
                    (globalState).f4 = d_748_v0_
                    d_815_v44_: _dafny.Seq
                    d_815_v44_ = _dafny.SeqWithoutIsStrInference([d_748_v0_])
                    d_816_v45_: _dafny.Seq
                    d_816_v45_ = _dafny.SeqWithoutIsStrInference([d_748_v0_, (d_815_v44_) <= (d_815_v44_), d_748_v0_])
                    d_816_v45_ = (d_815_v44_) + ((d_815_v44_).set(default__.safeIndex(d_749_v1_, len(d_815_v44_)), d_748_v0_))
                    d_817_v46_: _dafny.Array
                    nw136_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                    d_817_v46_ = nw136_
                    d_818_v47_: _dafny.Seq
                    d_818_v47_ = _dafny.SeqWithoutIsStrInference([d_749_v1_, len(d_816_v45_)])
                    d_819_v48_: _dafny.MultiSet
                    d_819_v48_ = _dafny.MultiSet([(0) - (d_749_v1_)])
                    d_820_v49_: _dafny.Array
                    nw137_ = _dafny.Array(None, 25)
                    nw137_[int(0)] = not(d_748_v0_)
                    nw137_[int(1)] = d_748_v0_
                    nw137_[int(2)] = d_748_v0_
                    nw137_[int(3)] = (d_748_v0_)
                    nw137_[int(4)] = False
                    nw137_[int(5)] = d_748_v0_
                    nw137_[int(6)] = not(d_748_v0_)
                    nw137_[int(7)] = d_748_v0_
                    nw137_[int(8)] = not(d_748_v0_)
                    nw137_[int(9)] = d_748_v0_
                    nw137_[int(10)] = True
                    nw137_[int(11)] = d_748_v0_
                    nw137_[int(12)] = d_748_v0_
                    nw137_[int(13)] = False
                    nw137_[int(14)] = not(d_748_v0_)
                    nw137_[int(15)] = d_748_v0_
                    nw137_[int(16)] = d_748_v0_
                    nw137_[int(17)] = not(True)
                    nw137_[int(18)] = d_748_v0_
                    nw137_[int(19)] = d_748_v0_
                    nw137_[int(20)] = d_748_v0_
                    nw137_[int(21)] = d_748_v0_
                    nw137_[int(22)] = d_748_v0_
                    nw137_[int(23)] = d_748_v0_
                    nw137_[int(24)] = (self).fm3(d_818_v47_, d_819_v48_, d_748_v0_, globalState)
                    d_820_v49_ = nw137_
                    d_821_v50_: D2
                    d_821_v50_ = D2_DC5(d_820_v49_)
                    d_822_v51_: _dafny.MultiSet
                    d_822_v51_ = _dafny.MultiSet([d_748_v0_, d_748_v0_])
                    d_823_v52_: _dafny.Map
                    d_823_v52_ = _dafny.Map({d_749_v1_: d_822_v51_})
                    d_824_v53_: _dafny.Set
                    d_824_v53_ = _dafny.Set({d_749_v1_})
                    rhs143_ = d_817_v46_
                    rhs144_ = (self).fm3(_dafny.SeqWithoutIsStrInference([d_749_v1_ for d_825_i4_ in range(default__.abs(895))]), d_819_v48_, (self).fm3((d_818_v47_).set(default__.safeIndex(d_749_v1_, len(d_818_v47_)), d_749_v1_), d_819_v48_, False, globalState), globalState)
                    rhs145_ = d_821_v50_
                    rhs146_ = default__.safeDivisionInt((101) - (782), d_749_v1_)
                    rhs147_ = (self).fm3((d_818_v47_).set(default__.safeIndex(len(d_823_v52_), len(d_818_v47_)), d_749_v1_), _dafny.MultiSet([len(d_824_v53_), d_749_v1_]), False, globalState)
                    lhs122_ = globalState
                    d_817_v46_ = rhs143_
                    d_748_v0_ = rhs144_
                    d_821_v50_ = rhs145_
                    lhs122_.f8 = rhs146_
                    r1 = rhs147_
                    if d_748_v0_:
                        (globalState).f8 = default__.safeDivisionInt(len(_dafny.Map({not(d_748_v0_): d_748_v0_})), 945)
                        d_826_v54_: T0
                        nw138_ = C1()
                        nw138_.ctor__()
                        d_826_v54_ = nw138_
                        d_827_v55_: _dafny.Seq
                        d_827_v55_ = _dafny.SeqWithoutIsStrInference([d_826_v54_])
                        d_828_v56_: _dafny.Seq
                        d_828_v56_ = _dafny.SeqWithoutIsStrInference([d_827_v55_])
                        d_829_v57_: _dafny.Seq
                        d_829_v57_ = _dafny.SeqWithoutIsStrInference([(d_828_v56_)[default__.safeIndex(d_749_v1_, len(d_828_v56_))]])
                        (globalState).f8 = (d_749_v1_) - ((0) - ((0) - (len((d_828_v56_) + (d_829_v57_)))))
                        d_830_v58_: str
                        d_830_v58_ = _dafny.CodePoint('t')
                        d_831_v59_: _dafny.MultiSet
                        d_831_v59_ = _dafny.MultiSet([d_830_v58_])
                        d_831_v59_ = ((d_831_v59_).intersection(d_831_v59_)) | (d_831_v59_)
                        d_832_v60_: C1
                        nw139_ = C1()
                        nw139_.ctor__()
                        d_832_v60_ = nw139_
                        index166_ = default__.safeIndex(526, (d_812_v42_).length(0))
                        (d_812_v42_)[index166_] = len(d_816_v45_)
                        index167_ = default__.safeIndex(526, (d_812_v42_).length(0))
                        rhs148_ = d_749_v1_
                        rhs149_ = d_820_v49_
                        rhs150_ = not(d_748_v0_)
                        lhs123_ = d_812_v42_
                        lhs124_ = default__.safeIndex(526, (d_812_v42_).length(0))
                        lhs123_[lhs124_] = rhs148_
                        d_820_v49_ = rhs149_
                        d_748_v0_ = rhs150_
                    elif True:
                        (globalState).f4 = (_dafny.Set({not(not((self).fm3(_dafny.SeqWithoutIsStrInference([d_749_v1_]), _dafny.MultiSet(d_818_v47_), d_748_v0_, globalState))), d_748_v0_})).ispropersubset(_dafny.Set({d_748_v0_}))
                        d_833_v61_: C1
                        nw140_ = C1()
                        nw140_.ctor__()
                        d_833_v61_ = nw140_
                        d_834_v62_: _dafny.Map
                        d_834_v62_ = _dafny.Map({d_749_v1_: d_749_v1_})
                        d_834_v62_ = _dafny.Map({d_749_v1_: default__.safeModuloInt(len(d_824_v53_), d_749_v1_)})
                        (globalState).f8 = ((d_819_v48_)[d_749_v1_] if (d_749_v1_) in (d_819_v48_) else (d_749_v1_ if d_748_v0_ else len(d_834_v62_)))
                        d_834_v62_ = (d_834_v62_) | (d_834_v62_)
                    pass
            pass
        d_835_v63_: _dafny.Array
        def lambda47_(d_836_v0_):
            def lambda48_(d_837_i5_):
                return d_836_v0_

            return lambda48_

        init23_ = lambda47_(d_748_v0_)
        nw141_ = _dafny.Array(None, 14)
        for i0_23_ in range(nw141_.length(0)):
            nw141_[i0_23_] = init23_(i0_23_)
        d_835_v63_ = nw141_
        d_838_v64_: _dafny.Seq
        d_838_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdjl"))
        index168_ = default__.safeIndex(775, (d_835_v63_).length(0))
        (d_835_v63_)[index168_] = (d_838_v64_) <= (d_838_v64_)
        index169_ = default__.safeIndex(775, (d_835_v63_).length(0))
        rhs151_ = d_748_v0_
        rhs152_ = d_748_v0_
        rhs153_ = d_748_v0_
        rhs154_ = default__.safeModuloInt((d_749_v1_) + (d_749_v1_), d_749_v1_)
        rhs155_ = d_748_v0_
        lhs125_ = d_835_v63_
        lhs126_ = default__.safeIndex(775, (d_835_v63_).length(0))
        d_748_v0_ = rhs151_
        d_748_v0_ = rhs152_
        r1 = rhs153_
        d_749_v1_ = rhs154_
        lhs125_[lhs126_] = rhs155_
        d_839_v65_: _dafny.Map
        d_839_v65_ = _dafny.Map({d_748_v0_: (d_835_v63_)[default__.safeIndex(775, (d_835_v63_).length(0))]})
        r0 = (704 if ((d_839_v65_)[d_748_v0_] if (d_748_v0_) in (d_839_v65_) else d_748_v0_) else d_749_v1_)
        r1 = d_748_v0_
        r2 = d_749_v1_
        return r0, r1, r2

    def m12(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_840_v0_: int
        d_840_v0_ = 912
        d_841_v1_: bool
        d_841_v1_ = True
        (globalState).f4 = (d_840_v0_) != (default__.fm20(d_841_v1_, globalState))
        d_842_v2_: _dafny.Array
        def lambda49_(d_843_v1_):
            def lambda50_(d_844_i0_):
                return d_843_v1_

            return lambda50_

        init24_ = lambda49_(d_841_v1_)
        nw142_ = _dafny.Array(None, 6)
        for i0_24_ in range(nw142_.length(0)):
            nw142_[i0_24_] = init24_(i0_24_)
        d_842_v2_ = nw142_
        index170_ = default__.safeIndex(787, (d_842_v2_).length(0))
        (d_842_v2_)[index170_] = d_841_v1_
        index171_ = default__.safeIndex(787, (d_842_v2_).length(0))
        (d_842_v2_)[index171_] = d_841_v1_
        d_845_v3_: _dafny.Seq
        d_845_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pncslplhn"))
        d_846_v4_: _dafny.Map
        d_846_v4_ = _dafny.Map({d_845_v3_: d_840_v0_})
        d_846_v4_ = (d_846_v4_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "momthrphd")), 92)
        d_847_v5_: _dafny.Seq
        d_847_v5_ = _dafny.SeqWithoutIsStrInference([-139])
        (globalState).f0 = (d_847_v5_) + (d_847_v5_)
        d_848_v6_: _dafny.Set
        d_848_v6_ = _dafny.Set({d_845_v3_})
        d_849_v7_: D13
        d_849_v7_ = D13_DC29(d_848_v6_)
        if (d_848_v6_) == (((d_849_v7_).cf41) - (d_848_v6_)):
            (globalState).f4 = (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]
            nw143_ = _dafny.Array(False, 5)
            d_842_v2_ = nw143_
            d_850_v8_: _dafny.Map
            d_850_v8_ = _dafny.Map({d_840_v0_: d_840_v0_})
            d_851_v9_: D3
            d_851_v9_ = D3_DC8(d_850_v8_)
            d_852_v10_: _dafny.Seq
            d_852_v10_ = _dafny.SeqWithoutIsStrInference([False])
            d_853_v11_: _dafny.MultiSet
            d_853_v11_ = _dafny.MultiSet([d_852_v10_, d_852_v10_, d_852_v10_, d_852_v10_])
            d_854_v12_: _dafny.MultiSet
            d_854_v12_ = _dafny.MultiSet([default__.fm26(d_851_v9_, d_840_v0_, d_853_v11_, globalState)])
            d_855_v13_: _dafny.MultiSet
            d_855_v13_ = d_854_v12_
            (globalState).f4 = not ((d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]) or (((d_855_v13_)).ispropersubset(d_854_v12_))
            d_856_v14_: _dafny.MultiSet
            d_856_v14_ = _dafny.MultiSet([d_845_v3_])
            d_857_v15_: _dafny.MultiSet
            d_857_v15_ = _dafny.MultiSet([(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], d_841_v1_, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], d_841_v1_])
            d_858_v16_: _dafny.Seq
            d_858_v16_ = _dafny.SeqWithoutIsStrInference([d_854_v12_, _dafny.MultiSet([(d_856_v14_).cardinality, ((d_857_v15_)[d_841_v1_] if (d_841_v1_) in (d_857_v15_) else d_840_v0_)])])
            (globalState).f8 = len(d_858_v16_)
            d_859_v17_: _dafny.Array
            nw144_ = _dafny.Array(int(0), 1)
            d_859_v17_ = nw144_
            index172_ = default__.safeIndex(346, (d_859_v17_).length(0))
            (d_859_v17_)[index172_] = (d_840_v0_ if d_841_v1_ else d_840_v0_)
            d_860_v18_: _dafny.Set
            d_860_v18_ = _dafny.Set({d_841_v1_})
            d_861_v19_: D9
            d_861_v19_ = D9_DC20(d_860_v18_)
            d_862_v20_: _dafny.Seq
            d_862_v20_ = _dafny.SeqWithoutIsStrInference([d_860_v18_, d_860_v18_, d_860_v18_])
            pat_let_tv40_ = d_860_v18_
            d_863_v21_: _dafny.Array
            nw145_ = _dafny.Array(None, 26)
            nw145_[int(0)] = d_861_v19_
            nw145_[int(1)] = D9_DC20(d_860_v18_)
            nw145_[int(2)] = D9_DC20(_dafny.Set({d_841_v1_, d_841_v1_, d_841_v1_, d_841_v1_, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]}))
            nw145_[int(3)] = d_861_v19_
            nw145_[int(4)] = d_861_v19_
            nw145_[int(5)] = d_861_v19_
            nw145_[int(6)] = d_861_v19_
            nw145_[int(7)] = D9_DC20(d_860_v18_)
            nw145_[int(8)] = d_861_v19_
            nw145_[int(9)] = d_861_v19_
            nw145_[int(10)] = D9_DC20(d_860_v18_)
            nw145_[int(11)] = d_861_v19_
            nw145_[int(12)] = d_861_v19_
            nw145_[int(13)] = d_861_v19_
            nw145_[int(14)] = d_861_v19_
            nw145_[int(15)] = D9_DC20(default__.fm35(d_841_v1_, d_840_v0_, d_841_v1_, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], globalState))
            nw145_[int(16)] = d_861_v19_
            def iife135_(_pat_let32_0):
                def iife136_(d_864_dt__update__tmp_h0_):
                    def iife137_(_pat_let33_0):
                        def iife138_(d_865_dt__update_hcf25_h0_):
                            return D9_DC20(d_865_dt__update_hcf25_h0_)
                        return iife138_(_pat_let33_0)
                    return iife137_(pat_let_tv40_)
                return iife136_(_pat_let32_0)
            nw145_[int(17)] = iife135_(d_861_v19_)
            nw145_[int(18)] = d_861_v19_
            nw145_[int(19)] = D9_DC20((d_862_v20_)[default__.safeIndex(d_840_v0_, len(d_862_v20_))])
            nw145_[int(20)] = d_861_v19_
            nw145_[int(21)] = d_861_v19_
            nw145_[int(22)] = d_861_v19_
            nw145_[int(23)] = d_861_v19_
            nw145_[int(24)] = d_861_v19_
            nw145_[int(25)] = d_861_v19_
            d_863_v21_ = nw145_
            index173_ = default__.safeIndex(366, (d_863_v21_).length(0))
            (d_863_v21_)[index173_] = d_861_v19_
            index174_ = default__.safeIndex(787, (d_842_v2_).length(0))
            index175_ = default__.safeIndex(346, (d_859_v17_).length(0))
            index176_ = default__.safeIndex(366, (d_863_v21_).length(0))
            rhs156_ = (default__.fm26(D3_DC8(d_850_v8_), d_840_v0_, _dafny.MultiSet([d_852_v10_, _dafny.SeqWithoutIsStrInference([d_841_v1_]), d_852_v10_]), globalState)) * ((d_840_v0_) - (859))
            rhs157_ = d_841_v1_
            rhs158_ = d_840_v0_
            rhs159_ = D9_DC20(d_860_v18_)
            lhs127_ = globalState
            lhs128_ = d_842_v2_
            lhs129_ = default__.safeIndex(787, (d_842_v2_).length(0))
            lhs130_ = d_859_v17_
            lhs131_ = default__.safeIndex(346, (d_859_v17_).length(0))
            lhs132_ = d_863_v21_
            lhs133_ = default__.safeIndex(366, (d_863_v21_).length(0))
            lhs127_.f8 = rhs156_
            lhs128_[lhs129_] = rhs157_
            lhs130_[lhs131_] = rhs158_
            lhs132_[lhs133_] = rhs159_
        elif True:
            d_866_v22_: _dafny.MultiSet
            d_866_v22_ = _dafny.MultiSet([d_841_v1_])
            d_867_v23_: C1
            nw146_ = C1()
            nw146_.ctor__()
            d_867_v23_ = nw146_
            d_868_v24_: _dafny.Map
            d_868_v24_ = _dafny.Map({d_867_v23_: (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]})
            d_869_v25_: _dafny.Map
            d_869_v25_ = _dafny.Map({(d_868_v24_).set(d_867_v23_, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]): _dafny.MultiSet([not((d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]), (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]])})
            rhs160_ = len(d_845_v3_)
            rhs161_ = 16
            rhs162_ = ((d_866_v22_) - (((d_869_v25_)[d_868_v24_] if (d_868_v24_) in (d_869_v25_) else d_866_v22_))).issubset(d_866_v22_)
            rhs163_ = (d_845_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_870_i1_ in range(default__.abs(307))]))
            r0 = rhs160_
            r1 = rhs161_
            r3 = rhs162_
            d_845_v3_ = rhs163_
            (globalState).f8 = d_840_v0_
            d_871_v26_: str
            d_871_v26_ = _dafny.CodePoint('n')
            d_872_v27_: _dafny.Map
            d_872_v27_ = _dafny.Map({(default__.fm17(d_871_v26_, globalState)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isrckjn"))): (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]})
            d_872_v27_ = (d_872_v27_).set(d_841_v1_, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))])
            d_873_v28_: _dafny.Array
            nw147_ = _dafny.Array(_dafny.Set({}), 21)
            d_873_v28_ = nw147_
            d_874_v29_: _dafny.Map
            d_874_v29_ = _dafny.Map({(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]: d_840_v0_})
            d_875_v30_: _dafny.Set
            d_875_v30_ = _dafny.Set({d_874_v29_})
            index177_ = default__.safeIndex(779, (d_873_v28_).length(0))
            (d_873_v28_)[index177_] = d_875_v30_
            index178_ = default__.safeIndex(779, (d_873_v28_).length(0))
            (d_873_v28_)[index178_] = ((d_875_v30_) | (d_875_v30_)) - (d_875_v30_)
            d_876_v31_: _dafny.Map
            d_876_v31_ = _dafny.Map({d_840_v0_: d_841_v1_})
            d_877_v32_: D7
            d_877_v32_ = D7_DC17((d_876_v31_).set(-103, d_841_v1_))
            source16_ = d_877_v32_
            if source16_.is_DC18:
                (globalState).f8 = default__.safeDivisionInt(d_840_v0_, d_840_v0_)
                d_878_v33_: C1
                nw148_ = C1()
                nw148_.ctor__()
                d_878_v33_ = nw148_
                d_879_v35_: _dafny.Seq
                d_879_v35_ = _dafny.SeqWithoutIsStrInference([(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]])
                d_880_v36_: _dafny.MultiSet
                d_880_v36_ = _dafny.MultiSet([d_879_v35_])
                d_881_v37_: _dafny.Set
                d_881_v37_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_882_i2_ in range(default__.abs(107))]))})
                d_883_v38_: _dafny.Set
                d_883_v38_ = _dafny.Set({len(d_881_v37_), d_840_v0_, d_840_v0_, len(d_848_v6_), d_840_v0_})
                d_884_v39_: _dafny.MultiSet
                d_884_v39_ = _dafny.MultiSet([d_840_v0_])
                d_885_v40_: _dafny.Array
                nw149_ = _dafny.Array(None, 23)
                nw149_[int(0)] = d_840_v0_
                nw149_[int(1)] = (0) - ((d_840_v0_) + (d_840_v0_))
                nw149_[int(2)] = d_840_v0_
                nw149_[int(3)] = d_840_v0_
                nw149_[int(4)] = d_840_v0_
                nw149_[int(5)] = default__.safeModuloInt(d_840_v0_, d_840_v0_)
                def iife139_():
                    coll71_ = _dafny.Map()
                    compr_71_: _dafny.Seq
                    for compr_71_ in (d_880_v36_).Elements:
                        d_886_v34_: _dafny.Seq = compr_71_
                        if (d_886_v34_) in (d_880_v36_):
                            coll71_[d_886_v34_] = d_840_v0_
                    return _dafny.Map(coll71_)
                nw149_[int(6)] = len(iife139_()
                )
                nw149_[int(7)] = len(d_881_v37_)
                nw149_[int(8)] = default__.safeDivisionInt(d_840_v0_, 371)
                nw149_[int(9)] = d_840_v0_
                nw149_[int(10)] = d_840_v0_
                nw149_[int(11)] = d_840_v0_
                nw149_[int(12)] = len(d_874_v29_)
                nw149_[int(13)] = d_840_v0_
                nw149_[int(14)] = len(_dafny.Map({d_840_v0_: d_841_v1_}))
                nw149_[int(15)] = -418
                nw149_[int(16)] = (d_840_v0_) + (len(_dafny.SeqWithoutIsStrInference([(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], True, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], d_841_v1_, d_841_v1_])))
                nw149_[int(17)] = d_840_v0_
                nw149_[int(18)] = d_840_v0_
                nw149_[int(19)] = (d_847_v5_)[default__.safeIndex(d_840_v0_, len(d_847_v5_))]
                nw149_[int(20)] = ((d_884_v39_)[d_840_v0_] if (d_840_v0_) in (d_884_v39_) else d_840_v0_)
                nw149_[int(21)] = d_840_v0_
                nw149_[int(22)] = ((d_874_v29_)[(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]] if ((d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]) in (d_874_v29_) else len(d_879_v35_))
                d_885_v40_ = nw149_
                index179_ = default__.safeIndex(473, (d_885_v40_).length(0))
                (d_885_v40_)[index179_] = d_840_v0_
                d_887_v41_: D5
                d_887_v41_ = D5_DC14(len(default__.fm17(d_871_v26_, globalState)), d_841_v1_)
                index180_ = default__.safeIndex(473, (d_885_v40_).length(0))
                (d_885_v40_)[index180_] = (0) - (default__.safeModuloInt((0) - ((d_840_v0_) - (len((d_845_v3_).set(default__.safeIndex((d_887_v41_).cf20, len(d_845_v3_)), d_871_v26_)))), d_840_v0_))
                d_888_v42_: C1
                nw150_ = C1()
                nw150_.ctor__()
                d_888_v42_ = nw150_
            elif True:
                d_889___mcc_h0_ = source16_.cf23
                d_890_cf23_ = d_889___mcc_h0_
                d_871_v26_ = d_871_v26_
                r1 = -346
                d_876_v31_ = d_876_v31_
                d_891_v43_: _dafny.Array
                nw151_ = _dafny.Array(_dafny.Set({}), 11)
                d_891_v43_ = nw151_
                d_892_v44_: D10
                d_892_v44_ = D10_DC24(-90, d_841_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')]))
                d_893_v45_: _dafny.Set
                d_893_v45_ = _dafny.Set({238, (_dafny.MultiSet([d_840_v0_])).cardinality})
                d_894_v46_: _dafny.Set
                d_894_v46_ = _dafny.Set({d_840_v0_, len(d_893_v45_)})
                d_895_v47_: C0
                nw152_ = C0()
                nw152_.ctor__()
                d_895_v47_ = nw152_
                d_896_v48_: D13
                d_896_v48_ = D13_DC30(d_841_v1_, (d_892_v44_).cf32, d_894_v46_, d_895_v47_)
                pat_let_tv41_ = d_841_v1_
                index181_ = default__.safeIndex(529, (d_891_v43_).length(0))
                def iife140_(_pat_let34_0):
                    def iife141_(d_897_dt__update__tmp_h1_):
                        def iife142_(_pat_let35_0):
                            def iife143_(d_898_dt__update_hcf42_h0_):
                                return D13_DC30(d_898_dt__update_hcf42_h0_, (d_897_dt__update__tmp_h1_).cf43, (d_897_dt__update__tmp_h1_).cf44, (d_897_dt__update__tmp_h1_).cf45)
                            return iife143_(_pat_let35_0)
                        return iife142_(pat_let_tv41_)
                    return iife141_(_pat_let34_0)
                (d_891_v43_)[index181_] = (iife140_(d_896_v48_)).cf44
                index182_ = default__.safeIndex(529, (d_891_v43_).length(0))
                (d_891_v43_)[index182_] = d_894_v46_
        d_899_v49_: _dafny.Seq
        d_899_v49_ = _dafny.SeqWithoutIsStrInference([(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], True, (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]])
        d_900_v50_: D4
        d_900_v50_ = D4_DC10(d_899_v49_)
        pat_let_tv42_ = d_899_v49_
        def iife144_(_pat_let36_0):
            def iife145_(d_901_dt__update__tmp_h2_):
                def iife146_(_pat_let37_0):
                    def iife147_(d_902_dt__update_hcf13_h0_):
                        return D4_DC10(d_902_dt__update_hcf13_h0_)
                    return iife147_(_pat_let37_0)
                return iife146_(pat_let_tv42_)
            return iife145_(_pat_let36_0)
        d_900_v50_ = iife144_(d_900_v50_)
        r0 = d_840_v0_
        r1 = len((d_899_v49_ if (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))] else (_dafny.SeqWithoutIsStrInference([(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))], (d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([(d_842_v2_)[default__.safeIndex(787, (d_842_v2_).length(0))]]))))
        d_903_v51_: D0
        d_903_v51_ = D0_DC0(d_840_v0_)
        pat_let_tv43_ = d_842_v2_
        pat_let_tv44_ = d_842_v2_
        pat_let_tv45_ = d_842_v2_
        pat_let_tv46_ = d_842_v2_
        pat_let_tv47_ = d_840_v0_
        pat_let_tv48_ = d_840_v0_
        pat_let_tv49_ = d_840_v0_
        def lambda51_(source17_):
            if source17_.is_DC1:
                d_904___mcc_h1_ = source17_.cf1
                d_905___mcc_h2_ = source17_.cf2
                d_906_cf2_ = d_905___mcc_h2_
                d_907_cf1_ = d_904___mcc_h1_
                return len((_dafny.Map({(pat_let_tv44_)[default__.safeIndex(787, (pat_let_tv43_).length(0))]: d_906_cf2_})) | (_dafny.Map({(pat_let_tv46_)[default__.safeIndex(787, (pat_let_tv45_).length(0))]: len(_dafny.Map({d_906_cf2_: (_dafny.MultiSet([d_906_cf2_, d_906_cf2_])).cardinality}))})))
            elif source17_.is_DC2:
                return pat_let_tv47_
            elif source17_.is_DC0:
                d_908___mcc_h3_ = source17_.cf0
                d_909_cf0_ = d_908___mcc_h3_
                return (d_909_cf0_) * ((0) - (-441))
            elif True:
                d_910___mcc_h4_ = source17_.cf3
                d_911_cf3_ = d_910___mcc_h4_
                return pat_let_tv48_

        def iife148_(_pat_let38_0):
            def iife149_(d_912_dt__update__tmp_h3_):
                def iife150_(_pat_let39_0):
                    def iife151_(d_913_dt__update_hcf0_h0_):
                        return D0_DC0(d_913_dt__update_hcf0_h0_)
                    return iife151_(_pat_let39_0)
                return iife150_(pat_let_tv49_)
            return iife149_(_pat_let38_0)
        r2 = lambda51_(iife148_(d_903_v51_))
        d_914_v52_: _dafny.Array
        nw153_ = _dafny.Array(None, 18)
        d_914_v52_ = nw153_
        d_915_v53_: _dafny.Seq
        d_915_v53_ = _dafny.SeqWithoutIsStrInference([d_914_v52_])
        r3 = ((d_915_v53_)[default__.safeIndex(d_840_v0_, len(d_915_v53_))]) in (d_915_v53_)
        return r0, r1, r2, r3


class C3(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm2(self, globalState):
        def iife152_():
            def iife154_():
                coll74_ = _dafny.Map()
                compr_74_: _dafny.Seq
                for compr_74_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC2())]), _dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC3(D0_DC2())), D0_DC3(D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))), D0_DC3(D0_DC0(603))])})).Elements:
                    d_916_v1_: _dafny.Seq = compr_74_
                    if (d_916_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC2())]), _dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC3(D0_DC2())), D0_DC3(D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))), D0_DC3(D0_DC0(603))])})):
                        coll74_[d_916_v1_] = False
                return _dafny.Map(coll74_)
            coll72_ = _dafny.Map()
            def iife153_():
                coll73_ = _dafny.Map()
                compr_73_: _dafny.Seq
                for compr_73_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC2())]), _dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC3(D0_DC2())), D0_DC3(D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))), D0_DC3(D0_DC0(603))])})).Elements:
                    d_916_v1_: _dafny.Seq = compr_73_
                    if (d_916_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC2())]), _dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC3(D0_DC2())), D0_DC3(D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))))), D0_DC3(D0_DC0(603))])})):
                        coll73_[d_916_v1_] = False
                return _dafny.Map(coll73_)
            compr_72_: int
            for compr_72_ in (_dafny.Map({len(iife153_()
            ): True})).keys.Elements:
                d_917_v0_: int = compr_72_
                if (d_917_v0_) in (_dafny.Map({len(iife154_()
                ): True})):
                    coll72_[default__.safeDivisionInt(d_917_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idrigq"))))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_918_i0_ in range(default__.abs(-517))]))
            return _dafny.Map(coll72_)
        return ((_dafny.Map({-878: 740})) | (_dafny.Map({663: len(_dafny.SeqWithoutIsStrInference([not(True), False]))}))) | (iife152_()
        )

    def fm3(self, p0, p1, p2, globalState):
        return not(not(False))

    def fm15(self, p0, globalState):
        return (_dafny.MultiSet([(_dafny.Map({401: 811})) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifkhv")))): len(_dafny.Map({(_dafny.MultiSet([False, False])).cardinality: False}))}))])).cardinality

    def fm16(self, p0, globalState):
        if True:
            return (True) == (not(True))
        elif True:
            return (len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('k'), _dafny.CodePoint('j'), _dafny.CodePoint('u'), _dafny.CodePoint('x')])}))) != (len(_dafny.SeqWithoutIsStrInference([False])))

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_919_v0_: _dafny.Array
        def lambda52_(d_920_i0_):
            return not(False)

        init25_ = lambda52_
        nw154_ = _dafny.Array(None, 29)
        for i0_25_ in range(nw154_.length(0)):
            nw154_[i0_25_] = init25_(i0_25_)
        d_919_v0_ = nw154_
        d_921_v1_: bool
        d_921_v1_ = True
        index183_ = default__.safeIndex(239, (d_919_v0_).length(0))
        (d_919_v0_)[index183_] = d_921_v1_
        d_922_v2_: int
        d_922_v2_ = 370
        index184_ = default__.safeIndex(239, (d_919_v0_).length(0))
        (d_919_v0_)[index184_] = (d_922_v2_) > ((0) - (d_922_v2_))
        hi4_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_923_i2_ in range(default__.abs(-926))]))
        for d_924_i1_ in range(d_922_v2_, hi4_):
            d_925_v3_: _dafny.Seq
            d_925_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fm"))
            d_926_v4_: _dafny.Seq
            d_926_v4_ = _dafny.SeqWithoutIsStrInference([d_925_v3_])
            if (((d_926_v4_)[default__.safeIndex(d_922_v2_, len(d_926_v4_))]) + (d_925_v3_)) <= (d_925_v3_):
                d_927_v5_: D0
                d_927_v5_ = D0_DC1(d_924_i1_, d_924_i1_)
                (globalState).f8 = (0) - ((d_927_v5_).cf1)
                d_928_v6_: str
                d_928_v6_ = _dafny.CodePoint('w')
                d_928_v6_ = d_928_v6_
                d_925_v3_ = ((d_925_v3_) + (d_925_v3_)).set(default__.safeIndex(d_922_v2_, len((d_925_v3_) + (d_925_v3_))), d_928_v6_)
                index185_ = default__.safeIndex(239, (d_919_v0_).length(0))
                (d_919_v0_)[index185_] = not(d_921_v1_)
                d_929_v7_: C0
                nw155_ = C0()
                nw155_.ctor__()
                d_929_v7_ = nw155_
            elif True:
                d_930_v8_: _dafny.Map
                d_930_v8_ = _dafny.Map({d_924_i1_: d_922_v2_})
                d_931_v9_: D10
                d_931_v9_ = D10_DC24(((d_930_v8_)[d_924_i1_] if (d_924_i1_) in (d_930_v8_) else d_922_v2_), (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], d_925_v3_)
                r0 = (d_931_v9_).cf32
                (globalState).f8 = d_922_v2_
                (globalState).f8 = d_922_v2_
                d_932_v10_: _dafny.Set
                d_932_v10_ = _dafny.Set({d_924_i1_, 903, (0) - (d_922_v2_)})
                d_933_v11_: C0
                nw156_ = C0()
                nw156_.ctor__()
                d_933_v11_ = nw156_
                d_934_v12_: D13
                d_934_v12_ = D13_DC30(True, d_921_v1_, d_932_v10_, d_933_v11_)
                r1 = (d_934_v12_).cf42
                d_935_v13_: _dafny.Map
                d_935_v13_ = _dafny.Map({(d_922_v2_) > (d_922_v2_): (d_933_v11_).fm19(globalState)})
                d_935_v13_ = ((d_935_v13_) | (d_935_v13_)) | (_dafny.Map({d_921_v1_: d_924_i1_}))
            d_936_v14_: _dafny.Map
            d_936_v14_ = _dafny.Map({d_922_v2_: d_921_v1_})
            d_937_v15_: _dafny.Seq
            d_937_v15_ = _dafny.SeqWithoutIsStrInference([d_936_v14_])
            d_938_v16_: _dafny.Map
            d_938_v16_ = _dafny.Map({(d_937_v15_) < (d_937_v15_): d_924_i1_})
            d_938_v16_ = (d_938_v16_).set(False, (d_922_v2_) - (d_922_v2_))
            index186_ = default__.safeIndex(239, (d_919_v0_).length(0))
            rhs164_ = not((d_924_i1_) > (d_924_i1_))
            rhs165_ = d_924_i1_
            rhs166_ = d_921_v1_
            rhs167_ = d_921_v1_
            rhs168_ = (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]
            lhs134_ = d_919_v0_
            lhs135_ = default__.safeIndex(239, (d_919_v0_).length(0))
            lhs136_ = globalState
            lhs137_ = globalState
            lhs138_ = globalState
            lhs134_[lhs135_] = rhs164_
            lhs136_.f8 = rhs165_
            lhs137_.f4 = rhs166_
            r0 = rhs167_
            lhs138_.f4 = rhs168_
            if (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]:
                d_939_v17_: _dafny.Set
                d_939_v17_ = _dafny.Set({d_925_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfqsa")), (d_925_v3_) + (d_925_v3_), d_925_v3_, d_925_v3_})
                d_939_v17_ = (_dafny.Set({d_925_v3_})) | (d_939_v17_)
                rhs169_ = d_924_i1_
                rhs170_ = (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]
                rhs171_ = d_922_v2_
                lhs139_ = globalState
                d_922_v2_ = rhs169_
                d_921_v1_ = rhs170_
                lhs139_.f8 = rhs171_
                d_940_v18_: _dafny.Map
                d_940_v18_ = _dafny.Map({d_924_i1_: d_924_i1_})
                d_941_v20_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.Seq({}), 16)
                d_941_v20_ = nw157_
                d_942_v21_: _dafny.Map
                d_942_v21_ = _dafny.Map({d_941_v20_: False})
                d_943_v22_: _dafny.Map
                def iife155_():
                    coll75_ = _dafny.Set()
                    compr_75_: int
                    for compr_75_ in (d_940_v18_).keys.Elements:
                        d_944_v19_: int = compr_75_
                        if (d_944_v19_) in (d_940_v18_):
                            coll75_ = coll75_.union(_dafny.Set([(d_944_v19_) * (664)]))
                    return _dafny.Set(coll75_)
                d_943_v22_ = _dafny.Map({iife155_()
                : ((d_942_v21_)[d_941_v20_] if (d_941_v20_) in (d_942_v21_) else d_921_v1_)})
                d_943_v22_ = d_943_v22_
                d_945_v23_: D3
                out21_: D3
                out21_ = (self).m10(d_924_i1_, (d_936_v14_) | (d_936_v14_), 329, globalState)
                d_945_v23_ = out21_
                d_946_v24_: _dafny.Seq
                d_946_v24_ = _dafny.SeqWithoutIsStrInference([(0) - (d_922_v2_), (0) - ((0) - (d_922_v2_)), len(d_925_v3_)])
                d_947_v25_: _dafny.Array
                nw158_ = _dafny.Array(None, 6)
                nw158_[int(0)] = d_922_v2_
                nw158_[int(1)] = (d_946_v24_)[default__.safeIndex(d_924_i1_, len(d_946_v24_))]
                nw158_[int(2)] = d_924_i1_
                nw158_[int(3)] = d_924_i1_
                nw158_[int(4)] = d_922_v2_
                nw158_[int(5)] = d_922_v2_
                d_947_v25_ = nw158_
                d_948_v26_: D3
                d_948_v26_ = D3_DC8(d_940_v18_)
                d_949_v27_: _dafny.Seq
                d_949_v27_ = _dafny.SeqWithoutIsStrInference([(d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]])
                d_950_v28_: _dafny.MultiSet
                d_950_v28_ = _dafny.MultiSet([d_949_v27_])
                index187_ = default__.safeIndex(239, (d_919_v0_).length(0))
                rhs172_ = default__.safeDivisionInt((d_922_v2_) * (default__.fm26(d_948_v26_, d_922_v2_, d_950_v28_, globalState)), 853)
                rhs173_ = (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]
                rhs174_ = d_947_v25_
                rhs175_ = ((d_925_v3_) + (d_925_v3_)) + (d_925_v3_)
                rhs176_ = False
                lhs140_ = globalState
                lhs141_ = d_919_v0_
                lhs142_ = default__.safeIndex(239, (d_919_v0_).length(0))
                lhs140_.f3 = rhs172_
                lhs141_[lhs142_] = rhs173_
                d_947_v25_ = rhs174_
                d_925_v3_ = rhs175_
                r0 = rhs176_
            elif True:
                d_921_v1_ = d_921_v1_
                d_922_v2_ = default__.fm20((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], globalState)
                d_951_v29_: _dafny.Array
                nw159_ = _dafny.Array(int(0), 11)
                d_951_v29_ = nw159_
                index188_ = default__.safeIndex(389, (d_951_v29_).length(0))
                (d_951_v29_)[index188_] = d_924_i1_
                index189_ = default__.safeIndex(389, (d_951_v29_).length(0))
                (d_951_v29_)[index189_] = len(d_925_v3_)
                d_951_v29_ = d_951_v29_
                d_952_v30_: str
                d_952_v30_ = _dafny.CodePoint('r')
                d_952_v30_ = d_952_v30_
        if (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]:
            (globalState).f8 = ((d_922_v2_) - (d_922_v2_)) - ((0) - ((0) - (default__.safeModuloInt(625, 995))))
            d_953_v31_: _dafny.MultiSet
            d_953_v31_ = _dafny.MultiSet([d_922_v2_, d_922_v2_])
            d_954_v32_: _dafny.Set
            d_954_v32_ = _dafny.Set({d_922_v2_})
            d_955_v33_: C0
            nw160_ = C0()
            nw160_.ctor__()
            d_955_v33_ = nw160_
            d_956_v34_: D13
            d_956_v34_ = D13_DC30(d_921_v1_, True, d_954_v32_, d_955_v33_)
            index190_ = default__.safeIndex(239, (d_919_v0_).length(0))
            rhs177_ = (0) - ((0) - ((d_922_v2_ if (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))] else 889)))
            rhs178_ = (d_953_v31_).issubset((d_953_v31_).set(d_922_v2_, default__.abs(d_922_v2_)))
            rhs179_ = (d_956_v34_).cf42
            rhs180_ = (d_922_v2_) == (d_922_v2_)
            lhs143_ = globalState
            lhs144_ = d_919_v0_
            lhs145_ = default__.safeIndex(239, (d_919_v0_).length(0))
            d_922_v2_ = rhs177_
            lhs143_.f4 = rhs178_
            lhs144_[lhs145_] = rhs179_
            d_921_v1_ = rhs180_
            d_957_v35_: _dafny.Array
            nw161_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_957_v35_ = nw161_
            d_958_v36_: _dafny.Array
            nw162_ = _dafny.Array(int(0), 28)
            d_958_v36_ = nw162_
            index191_ = default__.safeIndex(243, (d_957_v35_).length(0))
            (d_957_v35_)[index191_] = d_958_v36_
            index192_ = default__.safeIndex(243, (d_957_v35_).length(0))
            (d_957_v35_)[index192_] = d_958_v36_
            d_959_v37_: _dafny.Set
            d_959_v37_ = _dafny.Set({d_921_v1_})
            d_959_v37_ = (d_959_v37_) - (_dafny.Set({(d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]}))
            if (d_921_v1_) or (d_921_v1_):
                (globalState).f4 = not((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))])
                r1 = (self).fm16(d_921_v1_, globalState)
                d_960_v38_: C1
                nw163_ = C1()
                nw163_.ctor__()
                d_960_v38_ = nw163_
                d_961_v39_: D0
                d_961_v39_ = D0_DC2()
                d_962_v40_: D9
                d_962_v40_ = D9_DC21((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], _dafny.SeqWithoutIsStrInference([(d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], d_921_v1_, (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]]), d_921_v1_)
                rhs181_ = d_961_v39_
                rhs182_ = (0) - (d_922_v2_)
                rhs183_ = ((d_922_v2_) <= (d_922_v2_) if ((d_962_v40_).cf27) < (_dafny.SeqWithoutIsStrInference([d_921_v1_, (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], False])) else ((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]) and (False))
                rhs184_ = d_921_v1_
                lhs146_ = globalState
                d_961_v39_ = rhs181_
                lhs146_.f3 = rhs182_
                r0 = rhs183_
                r1 = rhs184_
                d_963_v41_: _dafny.Array
                def lambda53_(d_964_i3_):
                    return _dafny.CodePoint('b')

                init26_ = lambda53_
                nw164_ = _dafny.Array(None, 23)
                for i0_26_ in range(nw164_.length(0)):
                    nw164_[i0_26_] = init26_(i0_26_)
                d_963_v41_ = nw164_
                d_965_v42_: _dafny.Seq
                d_965_v42_ = _dafny.SeqWithoutIsStrInference([(d_963_v41_ if (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))] else d_963_v41_), d_963_v41_, d_963_v41_])
                rhs185_ = ((d_965_v42_) + (d_965_v42_)) + (d_965_v42_)
                rhs186_ = default__.safeModuloInt(d_922_v2_, 776)
                lhs147_ = globalState
                d_965_v42_ = rhs185_
                lhs147_.f8 = rhs186_
            elif True:
                d_966_v43_: _dafny.Seq
                d_966_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ae"))
                d_967_v44_: _dafny.Map
                d_967_v44_ = _dafny.Map({d_966_v43_: d_922_v2_})
                d_967_v44_ = d_967_v44_
                d_968_v45_: str
                d_968_v45_ = _dafny.CodePoint('g')
                d_969_v46_: D3
                d_969_v46_ = D3_DC9()
                d_970_v47_: _dafny.Map
                d_970_v47_ = _dafny.Map({d_968_v45_: d_969_v46_})
                d_970_v47_ = (d_970_v47_).set(d_968_v45_, d_969_v46_)
                d_971_v48_: _dafny.Map
                d_971_v48_ = _dafny.Map({not(d_921_v1_): d_922_v2_})
                d_972_v49_: _dafny.Map
                d_972_v49_ = _dafny.Map({d_922_v2_: len(d_971_v48_)})
                d_973_v50_: _dafny.Seq
                d_973_v50_ = _dafny.SeqWithoutIsStrInference([len((d_971_v48_).set((self).fm16((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], globalState), len(d_966_v43_))), len(d_972_v49_)])
                (globalState).f4 = (self).fm3(d_973_v50_, (d_953_v31_) - (d_953_v31_), d_921_v1_, globalState)
                index193_ = default__.safeIndex(522, (d_958_v36_).length(0))
                (d_958_v36_)[index193_] = len(d_972_v49_)
                index194_ = default__.safeIndex(522, (d_958_v36_).length(0))
                (d_958_v36_)[index194_] = default__.safeModuloInt(d_922_v2_, (d_973_v50_)[default__.safeIndex(d_922_v2_, len(d_973_v50_))])
                d_974_v51_: _dafny.Seq
                d_974_v51_ = _dafny.SeqWithoutIsStrInference([(d_957_v35_)[default__.safeIndex(243, (d_957_v35_).length(0))], ((d_957_v35_)[default__.safeIndex(243, (d_957_v35_).length(0))] if d_921_v1_ else (d_957_v35_)[default__.safeIndex(243, (d_957_v35_).length(0))]), (d_957_v35_)[default__.safeIndex(243, (d_957_v35_).length(0))], (d_957_v35_)[default__.safeIndex(243, (d_957_v35_).length(0))]])
                d_974_v51_ = ((d_974_v51_) + (d_974_v51_) if True else d_974_v51_)
        elif True:
            d_975_v52_: _dafny.Array
            nw165_ = _dafny.Array(_dafny.Map({}), 5)
            d_975_v52_ = nw165_
            d_976_v53_: _dafny.Map
            d_976_v53_ = _dafny.Map({d_922_v2_: d_975_v52_})
            d_976_v53_ = (d_976_v53_).set(d_922_v2_, d_975_v52_)
            d_977_v54_: _dafny.Seq
            d_977_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeeyibtfx"))
            d_978_v55_: _dafny.MultiSet
            d_978_v55_ = _dafny.MultiSet([d_977_v54_])
            d_979_v56_: D10
            d_979_v56_ = D10_DC23(d_978_v55_)
            pat_let_tv50_ = d_978_v55_
            def iife156_(_pat_let40_0):
                def iife157_(d_980_dt__update__tmp_h0_):
                    def iife158_(_pat_let41_0):
                        def iife159_(d_981_dt__update_hcf30_h0_):
                            return D10_DC23(d_981_dt__update_hcf30_h0_)
                        return iife159_(_pat_let41_0)
                    return iife158_(pat_let_tv50_)
                return iife157_(_pat_let40_0)
            d_979_v56_ = iife156_((default__.fm36(len(d_977_v54_), _dafny.Set({(d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]}), d_921_v1_, globalState) if d_921_v1_ else D10_DC23(d_978_v55_)))
            (globalState).f2 = default__.fm27(d_921_v1_, globalState)
            d_982_v57_: _dafny.Seq
            d_982_v57_ = _dafny.SeqWithoutIsStrInference([-589])
            d_983_v58_: _dafny.MultiSet
            d_983_v58_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wc"))), len(d_982_v57_)])
            d_984_v59_: _dafny.Seq
            d_984_v59_ = _dafny.SeqWithoutIsStrInference([(d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]])
            d_985_v60_: D10
            d_985_v60_ = D10_DC24(d_922_v2_, d_921_v1_, d_977_v54_)
            d_986_v61_: D9
            d_986_v61_ = D9_DC21(d_921_v1_, d_984_v59_, (d_985_v60_).cf32)
            index195_ = default__.safeIndex(239, (d_919_v0_).length(0))
            (d_919_v0_)[index195_] = (((self).fm3(d_982_v57_, d_983_v58_, d_921_v1_, globalState) if (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))] else (d_986_v61_).cf26)) or (False)
            d_987_v62_: str
            d_987_v62_ = _dafny.CodePoint('o')
            (globalState).f3 = (d_922_v2_) + (len((d_977_v54_).set(default__.safeIndex(d_922_v2_, len(d_977_v54_)), d_987_v62_)))
        (globalState).f3 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))
        hi5_ = d_922_v2_
        for d_988_i4_ in range(d_922_v2_, hi5_):
            (globalState).f3 = default__.safeModuloInt(d_988_i4_, (_dafny.MultiSet([d_921_v1_])).cardinality)
            r0 = (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]
            d_989_v63_: C2
            nw166_ = C2()
            nw166_.ctor__()
            d_989_v63_ = nw166_
            (globalState).f3 = d_988_i4_
        d_990_v64_: D2
        d_990_v64_ = D2_DC7(_dafny.CodePoint('l'), (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], default__.fm20((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], globalState))
        source18_ = d_990_v64_
        if source18_.is_DC6:
            d_991___mcc_h0_ = source18_.cf6
            d_992___mcc_h1_ = source18_.cf7
            d_993___mcc_h2_ = source18_.cf8
            d_994_cf8_ = d_993___mcc_h2_
            d_995_cf7_ = d_992___mcc_h1_
            d_996_cf6_ = d_991___mcc_h0_
            d_997_v65_: _dafny.Map
            d_997_v65_ = _dafny.Map({d_921_v1_: d_921_v1_})
            d_998_v66_: _dafny.MultiSet
            d_998_v66_ = _dafny.MultiSet([not(d_921_v1_), d_995_cf7_, d_995_cf7_, ((d_997_v65_)[d_995_cf7_] if (d_995_cf7_) in (d_997_v65_) else not(False))])
            d_999_v67_: _dafny.Map
            d_999_v67_ = _dafny.Map({(d_998_v66_) | (d_998_v66_): d_996_cf6_})
            d_1000_v68_: D0
            d_1000_v68_ = D0_DC1(d_996_cf6_, d_996_cf6_)
            d_1001_v69_: D0
            d_1001_v69_ = D0_DC3(d_1000_v68_)
            d_1002_v70_: _dafny.Seq
            d_1002_v70_ = _dafny.SeqWithoutIsStrInference([d_1001_v69_, d_1001_v69_, D0_DC3(d_1000_v68_), default__.fm38(d_921_v1_, globalState), D0_DC3(d_1000_v68_)])
            d_1003_v71_: str
            d_1003_v71_ = _dafny.CodePoint('o')
            d_999_v67_ = default__.fm37(d_1002_v70_, default__.fm17(d_1003_v71_, globalState), d_922_v2_, (d_922_v2_) + (d_996_cf6_), globalState)
            d_1004_v73_: _dafny.Map
            d_1004_v73_ = _dafny.Map({d_922_v2_: d_996_cf6_})
            d_1005_v74_: _dafny.Set
            def iife160_():
                coll76_ = _dafny.Map()
                compr_76_: int
                for compr_76_ in _dafny.IntegerRange(143, -189):
                    d_1006_v72_: int = compr_76_
                    if ((143) <= (d_1006_v72_)) and ((d_1006_v72_) < (-189)):
                        coll76_[(d_1006_v72_) * (len(_dafny.SeqWithoutIsStrInference([(d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))], not(d_995_cf7_), not(d_921_v1_), d_994_cf8_])))] = d_922_v2_
                return _dafny.Map(coll76_)
            d_1005_v74_ = _dafny.Set({(iife160_()
            ) | (d_1004_v73_)})
            d_1007_v75_: _dafny.Seq
            d_1007_v75_ = _dafny.SeqWithoutIsStrInference([d_1005_v74_])
            d_1005_v74_ = (d_1007_v75_)[default__.safeIndex(d_996_cf6_, len(d_1007_v75_))]
            d_1008_v76_: _dafny.Array
            nw167_ = _dafny.Array(int(0), 25)
            d_1008_v76_ = nw167_
            d_1008_v76_ = d_1008_v76_
            d_1009_v77_: _dafny.Seq
            d_1009_v77_ = _dafny.SeqWithoutIsStrInference([(d_994_cf8_ if (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))] else (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))])])
            d_1010_v78_: _dafny.Array
            nw168_ = _dafny.Array(_dafny.CodePoint('D'), 4)
            d_1010_v78_ = nw168_
            d_1011_v79_: _dafny.Set
            d_1011_v79_ = _dafny.Set({d_1010_v78_})
            d_1012_v80_: _dafny.Set
            d_1012_v80_ = _dafny.Set({d_995_cf7_, (d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))]})
            d_1009_v77_ = default__.fm31((d_1011_v79_).ispropersubset(_dafny.Set({d_1010_v78_, d_1010_v78_, d_1010_v78_})), (_dafny.Set({(d_1009_v77_)[default__.safeIndex((0) - (d_996_cf6_), len(d_1009_v77_))], d_994_cf8_})).issubset(d_1012_v80_), not(((d_919_v0_)[default__.safeIndex(239, (d_919_v0_).length(0))] if d_995_cf7_ else d_921_v1_)), globalState)
        elif source18_.is_DC7:
            d_1013___mcc_h3_ = source18_.cf9
            d_1014___mcc_h4_ = source18_.cf10
            d_1015___mcc_h5_ = source18_.cf11
            d_1016_cf11_ = d_1015___mcc_h5_
            d_1017_cf10_ = d_1014___mcc_h4_
            d_1018_cf9_ = d_1013___mcc_h3_
            d_1019_v81_: _dafny.Map
            d_1019_v81_ = _dafny.Map({d_1017_cf10_: d_1018_cf9_})
            d_1019_v81_ = (d_1019_v81_).set((self).fm16(d_1017_cf10_, globalState), d_1018_cf9_)
            (globalState).f4 = False
            d_1020_v82_: C1
            nw169_ = C1()
            nw169_.ctor__()
            d_1020_v82_ = nw169_
            d_1021_v83_: C0
            nw170_ = C0()
            nw170_.ctor__()
            d_1021_v83_ = nw170_
        elif True:
            d_1022___mcc_h6_ = source18_.cf5
            d_1023_cf5_ = d_1022___mcc_h6_
            d_1024_v84_: _dafny.Map
            d_1024_v84_ = _dafny.Map({d_922_v2_: d_922_v2_})
            d_1025_v85_: D3
            d_1025_v85_ = D3_DC8(d_1024_v84_)
            d_1025_v85_ = D3_DC8(d_1024_v84_)
            d_1026_v86_: str
            d_1026_v86_ = _dafny.CodePoint('t')
            d_1026_v86_ = d_1026_v86_
            d_1027_v87_: _dafny.Seq
            d_1027_v87_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exys"))])
            d_1028_v88_: _dafny.Map
            d_1028_v88_ = _dafny.Map({934: (d_1027_v87_)[default__.safeIndex(d_922_v2_, len(d_1027_v87_))]})
            d_1029_v89_: _dafny.Set
            d_1029_v89_ = _dafny.Set({d_1026_v86_})
            d_1030_v90_: _dafny.Seq
            d_1030_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjjgia"))
            d_1028_v88_ = (d_1028_v88_).set((d_922_v2_) + (len(d_1029_v89_)), d_1030_v90_)
            (globalState).f8 = d_922_v2_
        d_1031_v91_: _dafny.Set
        d_1031_v91_ = _dafny.Set({d_921_v1_})
        r0 = (d_1031_v91_).isdisjoint(d_1031_v91_)
        r1 = d_921_v1_
        return r0, r1

    def m10(self, p0, p1, p2, globalState):
        r0: D3 = D3.default()()
        d_1032_v0_: bool
        d_1032_v0_ = False
        d_1033_v1_: str
        d_1033_v1_ = _dafny.CodePoint('w')
        d_1034_v2_: _dafny.Seq
        d_1034_v2_ = _dafny.SeqWithoutIsStrInference([d_1033_v1_, d_1033_v1_])
        d_1035_v3_: D10
        d_1035_v3_ = D10_DC24(p2, d_1032_v0_, d_1034_v2_)
        if (d_1035_v3_).cf32:
            if (d_1032_v0_) or (False):
                d_1036_v4_: _dafny.Array
                nw171_ = _dafny.Array(False, 5)
                d_1036_v4_ = nw171_
                index196_ = default__.safeIndex(423, (d_1036_v4_).length(0))
                (d_1036_v4_)[index196_] = d_1032_v0_
                index197_ = default__.safeIndex(423, (d_1036_v4_).length(0))
                (d_1036_v4_)[index197_] = d_1032_v0_
                d_1037_v5_: _dafny.Array
                nw172_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1037_v5_ = nw172_
                d_1038_v6_: _dafny.Array
                nw173_ = _dafny.Array(int(0), 3)
                d_1038_v6_ = nw173_
                index198_ = default__.safeIndex(458, (d_1038_v6_).length(0))
                (d_1038_v6_)[index198_] = p0
                d_1039_v7_: _dafny.Seq
                d_1039_v7_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1040_v8_: _dafny.MultiSet
                d_1040_v8_ = _dafny.MultiSet([-279])
                index199_ = default__.safeIndex(423, (d_1036_v4_).length(0))
                index200_ = default__.safeIndex(458, (d_1038_v6_).length(0))
                rhs187_ = (self).fm3((d_1039_v7_) + (d_1039_v7_), d_1040_v8_, d_1032_v0_, globalState)
                rhs188_ = d_1037_v5_
                rhs189_ = (self).fm15(len((d_1034_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))), globalState)
                rhs190_ = p0
                lhs148_ = d_1036_v4_
                lhs149_ = default__.safeIndex(423, (d_1036_v4_).length(0))
                lhs150_ = d_1038_v6_
                lhs151_ = default__.safeIndex(458, (d_1038_v6_).length(0))
                lhs152_ = globalState
                lhs148_[lhs149_] = rhs187_
                d_1037_v5_ = rhs188_
                lhs150_[lhs151_] = rhs189_
                lhs152_.f3 = rhs190_
                (globalState).f0 = d_1039_v7_
                index201_ = default__.safeIndex(423, (d_1036_v4_).length(0))
                (d_1036_v4_)[index201_] = (d_1036_v4_)[default__.safeIndex(423, (d_1036_v4_).length(0))]
                d_1041_v9_: _dafny.Map
                d_1041_v9_ = _dafny.Map({p2: p2})
                d_1042_v10_: _dafny.Map
                d_1042_v10_ = _dafny.Map({d_1041_v9_: (d_1036_v4_)[default__.safeIndex(423, (d_1036_v4_).length(0))]})
                (globalState).f0 = (d_1039_v7_) + (((d_1039_v7_) + (d_1039_v7_)).set(default__.safeIndex((0) - (len(d_1042_v10_)), len((d_1039_v7_) + (d_1039_v7_))), p0))
            elif True:
                (globalState).f8 = default__.safeDivisionInt((p2) - (p0), p0)
                (globalState).f8 = p2
                d_1043_v11_: _dafny.Seq
                d_1043_v11_ = _dafny.SeqWithoutIsStrInference([d_1032_v0_, d_1032_v0_, not(d_1032_v0_)])
                d_1044_v12_: D4
                d_1044_v12_ = D4_DC10(d_1043_v11_)
                d_1045_v13_: _dafny.Seq
                d_1045_v13_ = _dafny.SeqWithoutIsStrInference([D4_DC12(d_1044_v12_)])
                d_1046_v15_: D9
                d_1046_v15_ = D9_DC22(d_1043_v11_)
                d_1047_v16_: _dafny.Map
                d_1047_v16_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1046_v15_])): len(d_1034_v2_)})
                d_1048_v17_: _dafny.Map
                d_1048_v17_ = _dafny.Map({d_1047_v16_: (self).fm16(d_1032_v0_, globalState)})
                d_1049_v18_: _dafny.Map
                def iife161_():
                    coll77_ = _dafny.Map()
                    compr_77_: _dafny.Map
                    for compr_77_ in (d_1048_v17_).keys.Elements:
                        d_1050_v14_: _dafny.Map = compr_77_
                        if (d_1050_v14_) in (d_1048_v17_):
                            coll77_[d_1050_v14_] = p2
                    return _dafny.Map(coll77_)
                d_1049_v18_ = _dafny.Map({d_1045_v13_: (iife161_()
                ).set(d_1047_v16_, p0)})
                d_1051_v19_: _dafny.Map
                d_1051_v19_ = _dafny.Map({d_1047_v16_: p0})
                d_1049_v18_ = (d_1049_v18_).set(d_1045_v13_, d_1051_v19_)
                d_1052_v20_: _dafny.MultiSet
                d_1052_v20_ = _dafny.MultiSet([default__.fm20(d_1032_v0_, globalState)])
                (globalState).f2 = (d_1052_v20_) - (d_1052_v20_)
                d_1053_v21_: _dafny.Map
                d_1053_v21_ = _dafny.Map({d_1032_v0_: len(d_1034_v2_)})
                d_1054_v22_: D12
                d_1054_v22_ = D12_DC26(d_1053_v21_)
                pat_let_tv51_ = d_1032_v0_
                pat_let_tv52_ = p2
                pat_let_tv53_ = d_1053_v21_
                d_1055_v23_: _dafny.Array
                nw174_ = _dafny.Array(None, 18)
                nw174_[int(0)] = d_1054_v22_
                nw174_[int(1)] = D12_DC26(_dafny.Map({d_1032_v0_: p2}))
                nw174_[int(2)] = d_1054_v22_
                nw174_[int(3)] = d_1054_v22_
                nw174_[int(4)] = d_1054_v22_
                nw174_[int(5)] = d_1054_v22_
                nw174_[int(6)] = d_1054_v22_
                def iife162_(_pat_let42_0):
                    def iife163_(d_1056_dt__update__tmp_h0_):
                        def iife164_(_pat_let43_0):
                            def iife165_(d_1057_dt__update_hcf35_h0_):
                                return D12_DC26(d_1057_dt__update_hcf35_h0_)
                            return iife165_(_pat_let43_0)
                        return iife164_(_dafny.Map({pat_let_tv51_: pat_let_tv52_}))
                    return iife163_(_pat_let42_0)
                nw174_[int(7)] = (d_1054_v22_ if d_1032_v0_ else iife162_(d_1054_v22_))
                def iife166_(_pat_let44_0):
                    def iife167_(d_1058_dt__update__tmp_h1_):
                        def iife168_(_pat_let45_0):
                            def iife169_(d_1059_dt__update_hcf35_h1_):
                                return D12_DC26(d_1059_dt__update_hcf35_h1_)
                            return iife169_(_pat_let45_0)
                        return iife168_(pat_let_tv53_)
                    return iife167_(_pat_let44_0)
                nw174_[int(8)] = iife166_(D12_DC26(d_1053_v21_))
                nw174_[int(9)] = D12_DC26(_dafny.Map({d_1032_v0_: p2}))
                nw174_[int(10)] = default__.fm39(p2, d_1032_v0_, d_1032_v0_, globalState)
                nw174_[int(11)] = d_1054_v22_
                nw174_[int(12)] = d_1054_v22_
                nw174_[int(13)] = d_1054_v22_
                nw174_[int(14)] = d_1054_v22_
                nw174_[int(15)] = default__.fm39(616, d_1032_v0_, d_1032_v0_, globalState)
                nw174_[int(16)] = d_1054_v22_
                nw174_[int(17)] = d_1054_v22_
                d_1055_v23_ = nw174_
                index202_ = default__.safeIndex(54, (d_1055_v23_).length(0))
                (d_1055_v23_)[index202_] = d_1054_v22_
                index203_ = default__.safeIndex(54, (d_1055_v23_).length(0))
                rhs191_ = (p0) + (len(d_1034_v2_))
                rhs192_ = D12_DC26(_dafny.Map({d_1032_v0_: p2}))
                rhs193_ = len(d_1034_v2_)
                lhs153_ = globalState
                lhs154_ = d_1055_v23_
                lhs155_ = default__.safeIndex(54, (d_1055_v23_).length(0))
                lhs156_ = globalState
                lhs153_.f3 = rhs191_
                lhs154_[lhs155_] = rhs192_
                lhs156_.f8 = rhs193_
            d_1032_v0_ = d_1032_v0_
            d_1060_v24_: _dafny.MultiSet
            d_1060_v24_ = _dafny.MultiSet([False, d_1032_v0_, d_1032_v0_, d_1032_v0_])
            d_1061_v25_: _dafny.Map
            d_1061_v25_ = _dafny.Map({default__.fm28(p0, globalState): (d_1060_v24_).isdisjoint(d_1060_v24_)})
            d_1062_v26_: _dafny.Map
            d_1062_v26_ = _dafny.Map({d_1032_v0_: 373})
            d_1063_v27_: _dafny.Array
            nw175_ = _dafny.Array(int(0), 26)
            d_1063_v27_ = nw175_
            d_1064_v28_: _dafny.Set
            d_1064_v28_ = _dafny.Set({d_1063_v27_, d_1063_v27_})
            d_1065_v29_: _dafny.Set
            d_1065_v29_ = _dafny.Set({869, 951, p2, len(d_1062_v26_), len(d_1064_v28_)})
            if ((d_1061_v25_)[d_1065_v29_] if (d_1065_v29_) in (d_1061_v25_) else d_1032_v0_):
                d_1066_v30_: C2
                nw176_ = C2()
                nw176_.ctor__()
                d_1066_v30_ = nw176_
                d_1067_v31_: _dafny.Seq
                d_1067_v31_ = _dafny.SeqWithoutIsStrInference([659, p0])
                d_1068_v32_: _dafny.MultiSet
                d_1068_v32_ = _dafny.MultiSet([p0])
                d_1032_v0_ = (self).fm3((d_1067_v31_) + (d_1067_v31_), d_1068_v32_, d_1032_v0_, globalState)
                d_1069_v33_: _dafny.Set
                d_1069_v33_ = _dafny.Set({d_1032_v0_, d_1032_v0_})
                d_1070_v34_: D9
                d_1070_v34_ = D9_DC20(d_1069_v33_)
                rhs194_ = (d_1032_v0_) or (d_1032_v0_)
                rhs195_ = ((d_1070_v34_).cf25).ispropersubset(d_1069_v33_)
                lhs157_ = globalState
                d_1032_v0_ = rhs194_
                lhs157_.f4 = rhs195_
                d_1071_v35_: _dafny.MultiSet
                d_1071_v35_ = d_1068_v32_
                d_1072_v36_: _dafny.Array
                nw177_ = _dafny.Array(False, 22)
                d_1072_v36_ = nw177_
                index204_ = default__.safeIndex(637, (d_1072_v36_).length(0))
                (d_1072_v36_)[index204_] = (d_1065_v29_).issubset(d_1065_v29_)
                index205_ = default__.safeIndex(637, (d_1072_v36_).length(0))
                rhs196_ = d_1071_v35_
                rhs197_ = False
                rhs198_ = d_1032_v0_
                rhs199_ = d_1032_v0_
                lhs158_ = d_1072_v36_
                lhs159_ = default__.safeIndex(637, (d_1072_v36_).length(0))
                d_1071_v35_ = rhs196_
                d_1032_v0_ = rhs197_
                d_1032_v0_ = rhs198_
                lhs158_[lhs159_] = rhs199_
                d_1073_v37_: _dafny.Seq
                d_1073_v37_ = _dafny.SeqWithoutIsStrInference([(d_1072_v36_)[default__.safeIndex(637, (d_1072_v36_).length(0))]])
                d_1074_v38_: _dafny.Map
                d_1074_v38_ = _dafny.Map({p0: p2})
                d_1073_v37_ = (_dafny.SeqWithoutIsStrInference([d_1032_v0_])) + (((d_1073_v37_).set(default__.safeIndex(len(default__.fm1(p0, d_1033_v1_, (d_1072_v36_)[default__.safeIndex(637, (d_1072_v36_).length(0))], d_1074_v38_, globalState)), len(d_1073_v37_)), (d_1072_v36_)[default__.safeIndex(637, (d_1072_v36_).length(0))])) + (default__.fm31(True, False, d_1032_v0_, globalState)))
            elif True:
                d_1032_v0_ = (d_1060_v24_).isdisjoint((d_1060_v24_) - (d_1060_v24_))
                d_1075_v39_: _dafny.Array
                nw178_ = _dafny.Array(False, 18)
                d_1075_v39_ = nw178_
                d_1076_v40_: _dafny.Map
                d_1076_v40_ = _dafny.Map({True: d_1075_v39_})
                d_1077_v41_: _dafny.Seq
                d_1077_v41_ = _dafny.SeqWithoutIsStrInference([d_1075_v39_])
                d_1076_v40_ = (d_1076_v40_).set(d_1032_v0_, (d_1077_v41_)[default__.safeIndex(p2, len(d_1077_v41_))])
                (globalState).f4 = d_1032_v0_
                d_1032_v0_ = (66) != (p2)
                d_1078_v42_: _dafny.Set
                d_1078_v42_ = _dafny.Set({d_1032_v0_})
                d_1079_v43_: _dafny.Array
                nw179_ = _dafny.Array(None, 3)
                nw179_[int(0)] = d_1078_v42_
                nw179_[int(1)] = d_1078_v42_
                nw179_[int(2)] = _dafny.Set({d_1032_v0_})
                d_1079_v43_ = nw179_
                d_1080_v44_: D4
                d_1080_v44_ = D4_DC11(True, d_1032_v0_, -888, d_1034_v2_)
                d_1081_v45_: _dafny.Map
                d_1081_v45_ = _dafny.Map({d_1079_v43_: len(_dafny.SeqWithoutIsStrInference([False, (d_1080_v44_).cf14, d_1032_v0_, d_1032_v0_]))})
                d_1081_v45_ = (d_1081_v45_).set(d_1079_v43_, p2)
            (globalState).f3 = p2
            d_1082_v46_: D5
            d_1082_v46_ = D5_DC15()
            source19_ = d_1082_v46_
            if source19_.is_DC14:
                d_1083___mcc_h0_ = source19_.cf20
                d_1084___mcc_h1_ = source19_.cf21
                d_1085_cf21_ = d_1084___mcc_h1_
                d_1086_cf20_ = d_1083___mcc_h0_
                d_1087_v47_: _dafny.Map
                d_1087_v47_ = _dafny.Map({d_1085_cf21_: d_1032_v0_})
                d_1088_v48_: _dafny.Seq
                d_1088_v48_ = _dafny.SeqWithoutIsStrInference([p2, (0) - (p2)])
                d_1089_v50_: D12
                def iife170_():
                    coll78_ = _dafny.Set()
                    compr_78_: int
                    for compr_78_ in (d_1088_v48_).Elements:
                        d_1090_v49_: int = compr_78_
                        if (d_1090_v49_) in (d_1088_v48_):
                            coll78_ = coll78_.union(_dafny.Set([(d_1090_v49_) + ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality)]))
                    return _dafny.Set(coll78_)
                d_1089_v50_ = D12_DC27(len(d_1034_v2_), len(d_1087_v47_), p0, len(iife170_()
))
                d_1091_v51_: D12
                d_1091_v51_ = D12_DC28(d_1089_v50_)
                d_1092_v52_: _dafny.Set
                d_1092_v52_ = _dafny.Set({d_1085_cf21_})
                rhs200_ = (default__.safeModuloInt(p0, p0)) + (d_1086_cf20_)
                rhs201_ = default__.safeModuloInt(len(d_1088_v48_), default__.safeDivisionInt(len(d_1092_v52_), (d_1060_v24_).cardinality))
                rhs202_ = default__.fm0(p2, globalState)
                rhs203_ = (_dafny.SeqWithoutIsStrInference([d_1033_v1_ for d_1093_i0_ in range(default__.abs(745))]) if ((0) - (((d_1062_v26_)[d_1032_v0_] if (d_1032_v0_) in (d_1062_v26_) else 461))) > (15) else d_1034_v2_)
                rhs204_ = d_1091_v51_
                lhs160_ = globalState
                lhs161_ = globalState
                lhs160_.f3 = rhs200_
                lhs161_.f8 = rhs201_
                d_1060_v24_ = rhs202_
                d_1034_v2_ = rhs203_
                d_1091_v51_ = rhs204_
                d_1087_v47_ = (d_1087_v47_).set(not (d_1085_cf21_) or (d_1085_cf21_), d_1032_v0_)
                (globalState).f8 = default__.safeModuloInt(d_1086_cf20_, default__.safeModuloInt(d_1086_cf20_, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))))
                d_1094_v53_: _dafny.Array
                nw180_ = _dafny.Array(False, 7)
                d_1094_v53_ = nw180_
                d_1095_v54_: _dafny.Map
                d_1095_v54_ = _dafny.Map({d_1032_v0_: d_1094_v53_})
                d_1095_v54_ = (d_1095_v54_).set(d_1032_v0_, d_1094_v53_)
            elif source19_.is_DC15:
                index206_ = default__.safeIndex(113, (d_1063_v27_).length(0))
                (d_1063_v27_)[index206_] = p2
                d_1096_v55_: _dafny.Array
                nw181_ = _dafny.Array(False, 4)
                d_1096_v55_ = nw181_
                d_1097_v56_: _dafny.Seq
                d_1097_v56_ = _dafny.SeqWithoutIsStrInference([d_1096_v55_])
                d_1098_v57_: _dafny.Seq
                d_1098_v57_ = _dafny.SeqWithoutIsStrInference([d_1096_v55_, d_1096_v55_])
                d_1099_v58_: D2
                d_1099_v58_ = D2_DC7(d_1033_v1_, (d_1032_v0_) or (d_1032_v0_), len((d_1097_v56_) + ((d_1098_v57_))))
                index207_ = default__.safeIndex(113, (d_1063_v27_).length(0))
                rhs205_ = p0
                rhs206_ = (d_1034_v2_) + ((d_1034_v2_) + (d_1034_v2_))
                rhs207_ = (default__.fm40(not(d_1032_v0_), globalState) if d_1032_v0_ else d_1099_v58_)
                lhs162_ = d_1063_v27_
                lhs163_ = default__.safeIndex(113, (d_1063_v27_).length(0))
                lhs162_[lhs163_] = rhs205_
                d_1034_v2_ = rhs206_
                d_1099_v58_ = rhs207_
                d_1100_v59_: _dafny.Seq
                d_1100_v59_ = _dafny.SeqWithoutIsStrInference([(d_1063_v27_)[default__.safeIndex(113, (d_1063_v27_).length(0))]])
                d_1101_v60_: _dafny.Array
                nw182_ = _dafny.Array(None, 3)
                nw182_[int(0)] = d_1100_v59_
                nw182_[int(1)] = _dafny.SeqWithoutIsStrInference([p0 for d_1102_i1_ in range(default__.abs(882))])
                nw182_[int(2)] = d_1100_v59_
                d_1101_v60_ = nw182_
                index208_ = default__.safeIndex(195, (d_1101_v60_).length(0))
                (d_1101_v60_)[index208_] = d_1100_v59_
                index209_ = default__.safeIndex(195, (d_1101_v60_).length(0))
                (d_1101_v60_)[index209_] = d_1100_v59_
                d_1103_v61_: _dafny.Map
                d_1103_v61_ = _dafny.Map({d_1032_v0_: d_1100_v59_})
                index210_ = default__.safeIndex(195, (d_1101_v60_).length(0))
                rhs208_ = (p0) >= ((d_1063_v27_)[default__.safeIndex(113, (d_1063_v27_).length(0))])
                rhs209_ = d_1033_v1_
                rhs210_ = p0
                rhs211_ = ((d_1103_v61_)[d_1032_v0_] if (d_1032_v0_) in (d_1103_v61_) else (d_1101_v60_)[default__.safeIndex(195, (d_1101_v60_).length(0))])
                lhs164_ = globalState
                lhs165_ = globalState
                lhs166_ = d_1101_v60_
                lhs167_ = default__.safeIndex(195, (d_1101_v60_).length(0))
                lhs164_.f4 = rhs208_
                d_1033_v1_ = rhs209_
                lhs165_.f3 = rhs210_
                lhs166_[lhs167_] = rhs211_
                d_1104_v62_: C1
                nw183_ = C1()
                nw183_.ctor__()
                d_1104_v62_ = nw183_
            elif True:
                d_1105___mcc_h2_ = source19_.cf19
                d_1106_cf19_ = d_1105___mcc_h2_
                d_1107_v63_: _dafny.Array
                nw184_ = _dafny.Array(D12.default()(), 13)
                d_1107_v63_ = nw184_
                d_1108_v64_: D12
                d_1108_v64_ = D12_DC26(_dafny.Map({False: p0}))
                index211_ = default__.safeIndex(145, (d_1107_v63_).length(0))
                (d_1107_v63_)[index211_] = d_1108_v64_
                index212_ = default__.safeIndex(145, (d_1107_v63_).length(0))
                (d_1107_v63_)[index212_] = d_1108_v64_
                d_1032_v0_ = d_1032_v0_
                d_1109_v65_: _dafny.Array
                nw185_ = _dafny.Array(False, 20)
                d_1109_v65_ = nw185_
                index213_ = default__.safeIndex(561, (d_1109_v65_).length(0))
                (d_1109_v65_)[index213_] = d_1032_v0_
                index214_ = default__.safeIndex(561, (d_1109_v65_).length(0))
                (d_1109_v65_)[index214_] = (d_1032_v0_) and (d_1032_v0_)
                rhs212_ = p0
                rhs213_ = (d_1109_v65_)[default__.safeIndex(561, (d_1109_v65_).length(0))]
                rhs214_ = (len(p1)) <= (p0)
                rhs215_ = (0) - (p0)
                lhs168_ = globalState
                lhs169_ = globalState
                lhs168_.f3 = rhs212_
                d_1032_v0_ = rhs213_
                d_1032_v0_ = rhs214_
                lhs169_.f8 = rhs215_
        elif True:
            d_1110_v66_: _dafny.Map
            d_1110_v66_ = _dafny.Map({not(d_1032_v0_): p2})
            d_1110_v66_ = (d_1110_v66_).set(d_1032_v0_, p0)
            d_1111_v67_: _dafny.Set
            d_1111_v67_ = _dafny.Set({d_1032_v0_, not(((p1)[p0] if (p0) in (p1) else d_1032_v0_)), False, d_1032_v0_})
            d_1112_v68_: _dafny.Map
            d_1112_v68_ = _dafny.Map({_dafny.Set({d_1032_v0_, d_1032_v0_, d_1032_v0_, d_1032_v0_}): (d_1111_v67_).ispropersubset(d_1111_v67_)})
            d_1113_v69_: D9
            d_1113_v69_ = D9_DC20(d_1111_v67_)
            d_1112_v68_ = (d_1112_v68_).set((d_1113_v69_).cf25, d_1032_v0_)
            d_1114_v70_: _dafny.Map
            d_1114_v70_ = _dafny.Map({len(d_1034_v2_): (self).fm15(p0, globalState)})
            d_1115_v71_: _dafny.Seq
            d_1115_v71_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1116_v72_: _dafny.Array
            nw186_ = _dafny.Array(D2.default()(), 14)
            d_1116_v72_ = nw186_
            d_1117_v73_: _dafny.Set
            d_1117_v73_ = _dafny.Set({d_1116_v72_})
            d_1118_v74_: _dafny.Array
            nw187_ = _dafny.Array(None, 25)
            nw187_[int(0)] = len(_dafny.Set({d_1032_v0_}))
            nw187_[int(1)] = p0
            nw187_[int(2)] = p0
            nw187_[int(3)] = ((_dafny.MultiSet([len(_dafny.Map({True: p2}))])).cardinality) + (p2)
            nw187_[int(4)] = (-628) + (p2)
            nw187_[int(5)] = len(_dafny.SeqWithoutIsStrInference([(0) - (p2), p2]))
            nw187_[int(6)] = p0
            nw187_[int(7)] = p0
            nw187_[int(8)] = p0
            nw187_[int(9)] = (0) - (p2)
            nw187_[int(10)] = len(_dafny.Set({d_1033_v1_, d_1033_v1_, _dafny.CodePoint('r')}))
            nw187_[int(11)] = ((d_1114_v70_)[p0] if (p0) in (d_1114_v70_) else p2)
            nw187_[int(12)] = -132
            nw187_[int(13)] = p0
            nw187_[int(14)] = p2
            nw187_[int(15)] = (p0) * (p2)
            nw187_[int(16)] = p0
            nw187_[int(17)] = (p2 if d_1032_v0_ else 319)
            nw187_[int(18)] = 607
            nw187_[int(19)] = (0) - ((p0) + (p0))
            nw187_[int(20)] = p2
            nw187_[int(21)] = p0
            nw187_[int(22)] = len(d_1115_v71_)
            nw187_[int(23)] = len(d_1117_v73_)
            nw187_[int(24)] = p2
            d_1118_v74_ = nw187_
            d_1118_v74_ = d_1118_v74_
            d_1119_v75_: _dafny.Array
            nw188_ = _dafny.Array(False, 17)
            d_1119_v75_ = nw188_
            index215_ = default__.safeIndex(779, (d_1119_v75_).length(0))
            (d_1119_v75_)[index215_] = d_1032_v0_
            d_1120_v76_: _dafny.MultiSet
            d_1120_v76_ = _dafny.MultiSet([d_1116_v72_])
            d_1121_v77_: _dafny.Set
            d_1121_v77_ = _dafny.Set({p0})
            index216_ = default__.safeIndex(779, (d_1119_v75_).length(0))
            (d_1119_v75_)[index216_] = (((d_1120_v76_).intersection((d_1120_v76_))).cardinality) not in (d_1121_v77_)
            if d_1032_v0_:
                d_1122_v78_: C0
                nw189_ = C0()
                nw189_.ctor__()
                d_1122_v78_ = nw189_
                d_1114_v70_ = (d_1114_v70_).set((len(d_1111_v67_)) + (p2), p0)
                d_1123_v79_: _dafny.Seq
                d_1123_v79_ = _dafny.SeqWithoutIsStrInference([False])
                d_1032_v0_ = (d_1123_v79_)[default__.safeIndex(p2, len(d_1123_v79_))]
                (globalState).f4 = (d_1123_v79_)[default__.safeIndex(p0, len(d_1123_v79_))]
                d_1124_v80_: _dafny.Map
                d_1124_v80_ = _dafny.Map({(d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]: (d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]})
                rhs216_ = p2
                rhs217_ = (d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]
                rhs218_ = (D13_DC30(((d_1124_v80_)[d_1032_v0_] if (d_1032_v0_) in (d_1124_v80_) else (d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]), d_1032_v0_, d_1121_v77_, d_1122_v78_)).cf44
                lhs170_ = globalState
                lhs171_ = globalState
                lhs170_.f3 = rhs216_
                lhs171_.f4 = rhs217_
                d_1121_v77_ = rhs218_
            elif True:
                d_1114_v70_ = (d_1114_v70_) | (d_1114_v70_)
                d_1125_v81_: _dafny.MultiSet
                d_1125_v81_ = default__.fm0((0) - (p2), globalState)
                d_1126_v82_: _dafny.MultiSet
                d_1126_v82_ = _dafny.MultiSet([(d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]])
                d_1125_v81_ = (d_1126_v82_) - (d_1126_v82_)
                d_1127_v83_: C2
                nw190_ = C2()
                nw190_.ctor__()
                d_1127_v83_ = nw190_
                d_1128_v84_: _dafny.Seq
                d_1128_v84_ = _dafny.SeqWithoutIsStrInference([(self).fm16(False, globalState)])
                d_1129_v85_: _dafny.Seq
                d_1129_v85_ = _dafny.SeqWithoutIsStrInference([d_1032_v0_, (True if (d_1128_v84_)[default__.safeIndex(default__.fm20((d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))], globalState), len(d_1128_v84_))] else (d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]), (d_1119_v75_)[default__.safeIndex(779, (d_1119_v75_).length(0))]])
                d_1032_v0_ = (d_1128_v84_)[default__.safeIndex(p2, len(d_1128_v84_))]
                d_1034_v2_ = d_1034_v2_
        d_1130_v86_: _dafny.Seq
        d_1130_v86_ = _dafny.SeqWithoutIsStrInference([d_1032_v0_])
        pat_let_tv54_ = d_1034_v2_
        pat_let_tv55_ = d_1130_v86_
        def iife171_(_pat_let46_0):
            def iife172_(d_1131_dt__update__tmp_h2_):
                def iife173_(_pat_let47_0):
                    def iife174_(d_1132_dt__update_hcf17_h0_):
                        def iife175_(_pat_let48_0):
                            def iife176_(d_1133_dt__update_hcf16_h0_):
                                return D4_DC11((d_1131_dt__update__tmp_h2_).cf14, (d_1131_dt__update__tmp_h2_).cf15, d_1133_dt__update_hcf16_h0_, d_1132_dt__update_hcf17_h0_)
                            return iife176_(_pat_let48_0)
                        return iife175_(len(pat_let_tv55_))
                    return iife174_(_pat_let47_0)
                return iife173_(pat_let_tv54_)
            return iife172_(_pat_let46_0)
        source20_ = iife171_(D4_DC11(d_1032_v0_, d_1032_v0_, p0, d_1034_v2_))
        if source20_.is_DC11:
            d_1134___mcc_h3_ = source20_.cf14
            d_1135___mcc_h4_ = source20_.cf15
            d_1136___mcc_h5_ = source20_.cf16
            d_1137___mcc_h6_ = source20_.cf17
            d_1138_cf17_ = d_1137___mcc_h6_
            d_1139_cf16_ = d_1136___mcc_h5_
            d_1140_cf15_ = d_1135___mcc_h4_
            d_1141_cf14_ = d_1134___mcc_h3_
            d_1142_v87_: _dafny.Array
            def lambda54_(d_1143_cf17_):
                def lambda55_(d_1144_i2_):
                    return d_1143_cf17_

                return lambda55_

            init27_ = lambda54_(d_1138_cf17_)
            nw191_ = _dafny.Array(None, 5)
            for i0_27_ in range(nw191_.length(0)):
                nw191_[i0_27_] = init27_(i0_27_)
            d_1142_v87_ = nw191_
            index217_ = default__.safeIndex(374, (d_1142_v87_).length(0))
            (d_1142_v87_)[index217_] = d_1034_v2_
            index218_ = default__.safeIndex(374, (d_1142_v87_).length(0))
            (d_1142_v87_)[index218_] = default__.fm25(globalState)
            d_1145_v88_: D9
            d_1145_v88_ = D9_DC22(d_1130_v86_)
            d_1146_v89_: _dafny.Seq
            d_1146_v89_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
            d_1147_v90_: _dafny.Map
            d_1147_v90_ = _dafny.Map({d_1139_cf16_: p0})
            d_1148_v91_: _dafny.Seq
            d_1148_v91_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1140_cf15_])])
            d_1149_v92_: _dafny.MultiSet
            d_1149_v92_ = _dafny.MultiSet([d_1130_v86_, (d_1148_v91_)[default__.safeIndex(default__.fm20(d_1141_cf14_, globalState), len(d_1148_v91_))], d_1130_v86_])
            pat_let_tv56_ = d_1147_v90_
            def iife177_(_pat_let49_0):
                def iife178_(d_1150_dt__update__tmp_h3_):
                    def iife179_(_pat_let50_0):
                        def iife180_(d_1151_dt__update_hcf12_h0_):
                            return D3_DC8(d_1151_dt__update_hcf12_h0_)
                        return iife180_(_pat_let50_0)
                    return iife179_(pat_let_tv56_)
                return iife178_(_pat_let49_0)
            rhs219_ = d_1146_v89_
            rhs220_ = default__.fm26(iife177_(D3_DC8(d_1147_v90_)), (d_1139_cf16_) + (p0), d_1149_v92_, globalState)
            rhs221_ = d_1145_v88_
            lhs172_ = globalState
            lhs173_ = globalState
            lhs172_.f0 = rhs219_
            lhs173_.f8 = rhs220_
            d_1145_v88_ = rhs221_
            d_1152_v93_: _dafny.Array
            nw192_ = _dafny.Array(False, 9)
            d_1152_v93_ = nw192_
            d_1153_v94_: D4
            d_1153_v94_ = D4_DC11(d_1141_cf14_, d_1140_cf15_, len(d_1138_cf17_), d_1034_v2_)
            index219_ = default__.safeIndex(634, (d_1152_v93_).length(0))
            (d_1152_v93_)[index219_] = (d_1153_v94_).cf14
            index220_ = default__.safeIndex(634, (d_1152_v93_).length(0))
            (d_1152_v93_)[index220_] = (not(d_1141_cf14_) if d_1032_v0_ else not((d_1130_v86_) <= (d_1130_v86_)))
            d_1154_v95_: _dafny.MultiSet
            d_1154_v95_ = _dafny.MultiSet([d_1032_v0_])
            d_1155_v96_: C2
            nw193_ = C2()
            nw193_.ctor__()
            d_1155_v96_ = nw193_
            d_1156_v97_: D16
            d_1156_v97_ = D16_DC34(d_1155_v96_)
            d_1157_v98_: _dafny.Map
            d_1157_v98_ = _dafny.Map({(d_1156_v97_).cf49: True})
            d_1158_v99_: T0
            nw194_ = C1()
            nw194_.ctor__()
            d_1158_v99_ = nw194_
            d_1159_v100_: _dafny.MultiSet
            d_1159_v100_ = _dafny.MultiSet([p0])
            d_1160_v101_: _dafny.Map
            d_1160_v101_ = _dafny.Map({d_1158_v99_: d_1159_v100_})
            d_1161_v102_: _dafny.Array
            nw195_ = _dafny.Array(None, 27)
            nw195_[int(0)] = (0) - (p0)
            nw195_[int(1)] = p2
            nw195_[int(2)] = ((d_1154_v95_)[False] if (False) in (d_1154_v95_) else len(d_1157_v98_))
            nw195_[int(3)] = d_1139_cf16_
            nw195_[int(4)] = len(_dafny.SeqWithoutIsStrInference([p0, p0]))
            nw195_[int(5)] = 141
            nw195_[int(6)] = d_1139_cf16_
            nw195_[int(7)] = (p2) * ((_dafny.MultiSet([d_1141_cf14_])).cardinality)
            nw195_[int(8)] = (p2) + (p2)
            nw195_[int(9)] = p0
            nw195_[int(10)] = p0
            nw195_[int(11)] = (d_1139_cf16_) + (d_1139_cf16_)
            nw195_[int(12)] = -851
            nw195_[int(13)] = (((d_1160_v101_)[d_1158_v99_] if (d_1158_v99_) in (d_1160_v101_) else d_1159_v100_)).cardinality
            nw195_[int(14)] = (0) - ((0) - (p0))
            nw195_[int(15)] = d_1139_cf16_
            nw195_[int(16)] = (-421) - (p2)
            nw195_[int(17)] = p0
            nw195_[int(18)] = (len(d_1147_v90_)) + (-24)
            nw195_[int(19)] = len(d_1146_v89_)
            nw195_[int(20)] = (0) - ((d_1139_cf16_) * (p0))
            nw195_[int(21)] = p0
            nw195_[int(22)] = p0
            nw195_[int(23)] = d_1139_cf16_
            nw195_[int(24)] = p2
            nw195_[int(25)] = p2
            nw195_[int(26)] = d_1139_cf16_
            d_1161_v102_ = nw195_
            index221_ = default__.safeIndex(367, (d_1161_v102_).length(0))
            (d_1161_v102_)[index221_] = (d_1146_v89_)[default__.safeIndex(p0, len(d_1146_v89_))]
            index222_ = default__.safeIndex(367, (d_1161_v102_).length(0))
            (d_1161_v102_)[index222_] = (p2) + (p0)
        elif source20_.is_DC10:
            d_1162___mcc_h7_ = source20_.cf13
            d_1163_cf13_ = d_1162___mcc_h7_
            d_1164_v103_: C0
            nw196_ = C0()
            nw196_.ctor__()
            d_1164_v103_ = nw196_
            d_1165_v104_: D0
            d_1165_v104_ = D0_DC2()
            source21_ = d_1165_v104_
            if source21_.is_DC1:
                d_1166___mcc_h9_ = source21_.cf1
                d_1167___mcc_h10_ = source21_.cf2
                d_1168_cf2_ = d_1167___mcc_h10_
                d_1169_cf1_ = d_1166___mcc_h9_
                d_1170_v105_: _dafny.Array
                nw197_ = _dafny.Array(None, 13)
                d_1170_v105_ = nw197_
                d_1171_v106_: _dafny.Set
                d_1171_v106_ = _dafny.Set({274, d_1168_cf2_, p2, 845, d_1168_cf2_})
                d_1172_v107_: D13
                d_1172_v107_ = D13_DC30(False, d_1032_v0_, d_1171_v106_, d_1164_v103_)
                index223_ = default__.safeIndex(562, (d_1170_v105_).length(0))
                (d_1170_v105_)[index223_] = d_1172_v107_
                d_1173_v108_: _dafny.Seq
                d_1173_v108_ = _dafny.SeqWithoutIsStrInference([d_1164_v103_, d_1164_v103_])
                index224_ = default__.safeIndex(562, (d_1170_v105_).length(0))
                (d_1170_v105_)[index224_] = D13_DC30(d_1032_v0_, d_1032_v0_, _dafny.Set({-260, p2}), (d_1173_v108_)[default__.safeIndex(d_1168_cf2_, len(d_1173_v108_))])
                d_1174_v109_: _dafny.Array
                nw198_ = _dafny.Array(False, 17)
                d_1174_v109_ = nw198_
                index225_ = default__.safeIndex(964, (d_1174_v109_).length(0))
                (d_1174_v109_)[index225_] = d_1032_v0_
                index226_ = default__.safeIndex(964, (d_1174_v109_).length(0))
                (d_1174_v109_)[index226_] = (d_1033_v1_) not in (d_1034_v2_)
                d_1175_v110_: _dafny.Map
                d_1175_v110_ = _dafny.Map({p2: 812})
                d_1176_v111_: D3
                d_1176_v111_ = D3_DC8(d_1175_v110_)
                d_1177_v112_: _dafny.MultiSet
                d_1177_v112_ = _dafny.MultiSet([d_1130_v86_])
                pat_let_tv57_ = d_1175_v110_
                def iife181_(_pat_let51_0):
                    def iife182_(d_1178_dt__update__tmp_h4_):
                        def iife183_(_pat_let52_0):
                            def iife184_(d_1179_dt__update_hcf12_h1_):
                                return D3_DC8(d_1179_dt__update_hcf12_h1_)
                            return iife184_(_pat_let52_0)
                        return iife183_(pat_let_tv57_)
                    return iife182_(_pat_let51_0)
                (globalState).f8 = default__.fm26(iife181_(d_1176_v111_), len(d_1034_v2_), d_1177_v112_, globalState)
                d_1180_v113_: bool
                d_1180_v113_ = not((d_1174_v109_)[default__.safeIndex(964, (d_1174_v109_).length(0))])
                d_1181_v114_: _dafny.Map
                d_1181_v114_ = _dafny.Map({d_1180_v113_: (d_1174_v109_)[default__.safeIndex(964, (d_1174_v109_).length(0))]})
                d_1182_v115_: _dafny.Seq
                d_1182_v115_ = _dafny.SeqWithoutIsStrInference([len(d_1034_v2_), p2, d_1169_cf1_])
                d_1181_v114_ = (d_1181_v114_).set(d_1180_v113_, (d_1182_v115_) <= (_dafny.SeqWithoutIsStrInference([len(d_1163_cf13_), p2])))
            elif source21_.is_DC2:
                d_1183_v116_: _dafny.Array
                nw199_ = _dafny.Array(int(0), 8)
                d_1183_v116_ = nw199_
                d_1184_v117_: _dafny.Array
                nw200_ = _dafny.Array(None, 5)
                nw200_[int(0)] = d_1183_v116_
                nw200_[int(1)] = d_1183_v116_
                nw200_[int(2)] = d_1183_v116_
                nw200_[int(3)] = d_1183_v116_
                nw200_[int(4)] = d_1183_v116_
                d_1184_v117_ = nw200_
                index227_ = default__.safeIndex(143, (d_1184_v117_).length(0))
                (d_1184_v117_)[index227_] = d_1183_v116_
                d_1185_v118_: _dafny.Seq
                d_1185_v118_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1186_v119_: _dafny.Map
                d_1186_v119_ = _dafny.Map({p0: (p0) * ((d_1185_v118_)[default__.safeIndex(len(d_1034_v2_), len(d_1185_v118_))])})
                d_1187_v120_: _dafny.Array
                def lambda56_(d_1188_v1_):
                    def lambda57_(d_1189_i3_):
                        return (_dafny.SeqWithoutIsStrInference([d_1188_v1_ for d_1190_i4_ in range(default__.abs(855))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('b'), d_1188_v1_, d_1188_v1_, _dafny.CodePoint('f')]))

                    return lambda57_

                init28_ = lambda56_(d_1033_v1_)
                nw201_ = _dafny.Array(None, 24)
                for i0_28_ in range(nw201_.length(0)):
                    nw201_[i0_28_] = init28_(i0_28_)
                d_1187_v120_ = nw201_
                index228_ = default__.safeIndex(206, (d_1187_v120_).length(0))
                (d_1187_v120_)[index228_] = d_1034_v2_
                d_1191_v121_: _dafny.Array
                d_1191_v121_ = d_1183_v116_
                index229_ = default__.safeIndex(143, (d_1184_v117_).length(0))
                index230_ = default__.safeIndex(206, (d_1187_v120_).length(0))
                rhs222_ = (0) - (len((default__.fm1(p0, d_1033_v1_, d_1032_v0_, d_1186_v119_, globalState)).set(default__.safeIndex(p0, len(default__.fm1(p0, d_1033_v1_, d_1032_v0_, d_1186_v119_, globalState))), 973)))
                rhs223_ = (len(d_1034_v2_)) - (len(d_1130_v86_))
                rhs224_ = (d_1191_v121_)
                rhs225_ = d_1186_v119_
                rhs226_ = d_1034_v2_
                lhs174_ = globalState
                lhs175_ = globalState
                lhs176_ = d_1184_v117_
                lhs177_ = default__.safeIndex(143, (d_1184_v117_).length(0))
                lhs178_ = d_1187_v120_
                lhs179_ = default__.safeIndex(206, (d_1187_v120_).length(0))
                lhs174_.f3 = rhs222_
                lhs175_.f8 = rhs223_
                lhs176_[lhs177_] = rhs224_
                d_1186_v119_ = rhs225_
                lhs178_[lhs179_] = rhs226_
                d_1192_v122_: _dafny.Array
                nw202_ = _dafny.Array(None, 13)
                nw202_[int(0)] = d_1032_v0_
                nw202_[int(1)] = d_1032_v0_
                nw202_[int(2)] = d_1032_v0_
                nw202_[int(3)] = d_1032_v0_
                nw202_[int(4)] = d_1032_v0_
                nw202_[int(5)] = d_1032_v0_
                nw202_[int(6)] = d_1032_v0_
                nw202_[int(7)] = d_1032_v0_
                nw202_[int(8)] = d_1032_v0_
                nw202_[int(9)] = d_1032_v0_
                nw202_[int(10)] = (self).fm16(d_1032_v0_, globalState)
                nw202_[int(11)] = d_1032_v0_
                nw202_[int(12)] = d_1032_v0_
                d_1192_v122_ = nw202_
                d_1193_v123_: _dafny.Map
                d_1193_v123_ = _dafny.Map({p0: d_1192_v122_})
                d_1193_v123_ = (d_1193_v123_) | (d_1193_v123_)
                d_1034_v2_ = (d_1187_v120_)[default__.safeIndex(206, (d_1187_v120_).length(0))]
                (globalState).f4 = d_1032_v0_
            elif source21_.is_DC0:
                d_1194___mcc_h11_ = source21_.cf0
                d_1195_cf0_ = d_1194___mcc_h11_
                (globalState).f4 = d_1032_v0_
                d_1196_v124_: _dafny.Array
                def lambda58_(d_1197_v2_):
                    def lambda59_(d_1198_i5_):
                        return d_1197_v2_

                    return lambda59_

                init29_ = lambda58_(d_1034_v2_)
                nw203_ = _dafny.Array(None, 16)
                for i0_29_ in range(nw203_.length(0)):
                    nw203_[i0_29_] = init29_(i0_29_)
                d_1196_v124_ = nw203_
                index231_ = default__.safeIndex(380, (d_1196_v124_).length(0))
                (d_1196_v124_)[index231_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwhy"))
                d_1199_v125_: _dafny.Map
                d_1199_v125_ = _dafny.Map({d_1032_v0_: d_1033_v1_})
                d_1200_v126_: _dafny.Map
                d_1200_v126_ = _dafny.Map({d_1034_v2_: (d_1032_v0_) not in (d_1199_v125_)})
                d_1201_v127_: _dafny.Array
                nw204_ = _dafny.Array(int(0), 8)
                d_1201_v127_ = nw204_
                index232_ = default__.safeIndex(530, (d_1201_v127_).length(0))
                (d_1201_v127_)[index232_] = p0
                d_1202_v128_: _dafny.Array
                d_1202_v128_ = d_1201_v127_
                d_1203_v129_: _dafny.Seq
                d_1203_v129_ = _dafny.SeqWithoutIsStrInference([d_1202_v128_])
                index233_ = default__.safeIndex(380, (d_1196_v124_).length(0))
                index234_ = default__.safeIndex(530, (d_1201_v127_).length(0))
                rhs227_ = d_1034_v2_
                rhs228_ = d_1200_v126_
                rhs229_ = len(((d_1130_v86_) + ((d_1163_cf13_) + (default__.fm31(d_1032_v0_, d_1032_v0_, d_1032_v0_, globalState)))).set(default__.safeIndex(len((d_1203_v129_ if d_1032_v0_ else d_1203_v129_)), len((d_1130_v86_) + ((d_1163_cf13_) + (default__.fm31(d_1032_v0_, d_1032_v0_, d_1032_v0_, globalState))))), d_1032_v0_))
                rhs230_ = d_1032_v0_
                lhs180_ = d_1196_v124_
                lhs181_ = default__.safeIndex(380, (d_1196_v124_).length(0))
                lhs182_ = d_1201_v127_
                lhs183_ = default__.safeIndex(530, (d_1201_v127_).length(0))
                lhs184_ = globalState
                lhs180_[lhs181_] = rhs227_
                d_1200_v126_ = rhs228_
                lhs182_[lhs183_] = rhs229_
                lhs184_.f4 = rhs230_
                (globalState).f4 = d_1032_v0_
                d_1204_v130_: _dafny.Seq
                d_1204_v130_ = _dafny.SeqWithoutIsStrInference([d_1195_cf0_, (d_1201_v127_)[default__.safeIndex(530, (d_1201_v127_).length(0))]])
                d_1205_v131_: _dafny.MultiSet
                d_1205_v131_ = _dafny.MultiSet([d_1195_cf0_])
                d_1206_v132_: _dafny.MultiSet
                d_1206_v132_ = _dafny.MultiSet([d_1205_v131_])
                d_1207_v133_: _dafny.Map
                d_1207_v133_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0 for d_1208_i6_ in range(default__.abs(-174))]): (0) - (d_1195_cf0_)})
                d_1209_v134_: _dafny.Map
                d_1209_v134_ = _dafny.Map({False: d_1032_v0_})
                d_1210_v135_: _dafny.Map
                d_1210_v135_ = _dafny.Map({d_1032_v0_: ((d_1209_v134_)[d_1032_v0_] if (d_1032_v0_) in (d_1209_v134_) else (self).fm3(d_1204_v130_, d_1205_v131_, d_1032_v0_, globalState))})
                d_1211_v136_: _dafny.Map
                d_1211_v136_ = _dafny.Map({p0: len(d_1210_v135_)})
                d_1212_v137_: _dafny.Set
                d_1212_v137_ = _dafny.Set({(_dafny.Map({d_1195_cf0_: p2})).set(len(d_1207_v133_), p0), d_1211_v136_})
                d_1213_v139_: _dafny.Array
                nw205_ = _dafny.Array(None, 20)
                nw205_[int(0)] = d_1032_v0_
                nw205_[int(1)] = not(d_1032_v0_)
                nw205_[int(2)] = (d_1204_v130_) < (_dafny.SeqWithoutIsStrInference([(d_1201_v127_)[default__.safeIndex(530, (d_1201_v127_).length(0))]]))
                nw205_[int(3)] = (d_1206_v132_).isdisjoint(d_1206_v132_)
                nw205_[int(4)] = d_1032_v0_
                nw205_[int(5)] = d_1032_v0_
                nw205_[int(6)] = d_1032_v0_
                def iife185_():
                    coll79_ = _dafny.Map()
                    compr_79_: int
                    for compr_79_ in _dafny.IntegerRange(-765, 433):
                        d_1214_v138_: int = compr_79_
                        if ((-765) <= (d_1214_v138_)) and ((d_1214_v138_) < (433)):
                            coll79_[(d_1214_v138_) - ((d_1201_v127_)[default__.safeIndex(530, (d_1201_v127_).length(0))])] = p0
                    return _dafny.Map(coll79_)
                nw205_[int(7)] = (_dafny.Set({iife185_()
                , d_1211_v136_})).issubset(d_1212_v137_)
                nw205_[int(8)] = True
                nw205_[int(9)] = not(not(d_1032_v0_))
                nw205_[int(10)] = d_1032_v0_
                nw205_[int(11)] = d_1032_v0_
                nw205_[int(12)] = (len(_dafny.SeqWithoutIsStrInference([d_1195_cf0_, p0]))) <= ((d_1201_v127_)[default__.safeIndex(530, (d_1201_v127_).length(0))])
                nw205_[int(13)] = d_1032_v0_
                nw205_[int(14)] = True
                nw205_[int(15)] = not(not(d_1032_v0_))
                nw205_[int(16)] = d_1032_v0_
                nw205_[int(17)] = d_1032_v0_
                nw205_[int(18)] = (_dafny.Set({d_1195_cf0_})).isdisjoint(_dafny.Set({p2, p2, len((d_1196_v124_)[default__.safeIndex(380, (d_1196_v124_).length(0))])}))
                nw205_[int(19)] = d_1032_v0_
                d_1213_v139_ = nw205_
                index235_ = default__.safeIndex(698, (d_1213_v139_).length(0))
                (d_1213_v139_)[index235_] = d_1032_v0_
                d_1215_v140_: D12
                d_1215_v140_ = D12_DC27(p0, d_1195_cf0_, len(_dafny.Map({d_1032_v0_: _dafny.Set({p0, d_1195_cf0_, p2})})), p0)
                d_1216_v141_: _dafny.Seq
                d_1216_v141_ = _dafny.SeqWithoutIsStrInference([d_1215_v140_])
                index236_ = default__.safeIndex(698, (d_1213_v139_).length(0))
                (d_1213_v139_)[index236_] = (d_1216_v141_) != ((d_1216_v141_) + (d_1216_v141_))
            elif True:
                d_1217___mcc_h12_ = source21_.cf3
                d_1218_cf3_ = d_1217___mcc_h12_
                d_1219_v142_: _dafny.Array
                nw206_ = _dafny.Array(False, 10)
                d_1219_v142_ = nw206_
                index237_ = default__.safeIndex(866, (d_1219_v142_).length(0))
                (d_1219_v142_)[index237_] = (D5_DC14(p0, d_1032_v0_)).cf21
                index238_ = default__.safeIndex(866, (d_1219_v142_).length(0))
                (d_1219_v142_)[index238_] = d_1032_v0_
                d_1032_v0_ = (d_1219_v142_)[default__.safeIndex(866, (d_1219_v142_).length(0))]
                d_1220_v143_: _dafny.Set
                d_1220_v143_ = _dafny.Set({(d_1219_v142_)[default__.safeIndex(866, (d_1219_v142_).length(0))], (d_1219_v142_)[default__.safeIndex(866, (d_1219_v142_).length(0))]})
                d_1221_v144_: _dafny.Array
                nw207_ = _dafny.Array(None, 10)
                nw207_[int(0)] = p0
                nw207_[int(1)] = p2
                nw207_[int(2)] = len(_dafny.SeqWithoutIsStrInference([d_1033_v1_ for d_1222_i7_ in range(default__.abs(772))]))
                nw207_[int(3)] = p2
                nw207_[int(4)] = len(_dafny.Set({p0}))
                nw207_[int(5)] = p2
                nw207_[int(6)] = p2
                nw207_[int(7)] = len(d_1034_v2_)
                nw207_[int(8)] = len(d_1163_cf13_)
                nw207_[int(9)] = len(d_1220_v143_)
                d_1221_v144_ = nw207_
                d_1223_v145_: _dafny.Map
                d_1223_v145_ = _dafny.Map({d_1221_v144_: d_1219_v142_})
                d_1224_v146_: _dafny.Seq
                d_1224_v146_ = _dafny.SeqWithoutIsStrInference([d_1223_v145_, d_1223_v145_])
                d_1225_v147_: _dafny.Array
                nw208_ = _dafny.Array(None, 26)
                nw208_[int(0)] = (_dafny.Map({d_1221_v144_: d_1219_v142_}) if d_1032_v0_ else d_1223_v145_)
                nw208_[int(1)] = d_1223_v145_
                nw208_[int(2)] = (d_1223_v145_) | ((d_1224_v146_)[default__.safeIndex(p0, len(d_1224_v146_))])
                nw208_[int(3)] = d_1223_v145_
                nw208_[int(4)] = (d_1223_v145_ if (d_1219_v142_)[default__.safeIndex(866, (d_1219_v142_).length(0))] else d_1223_v145_)
                nw208_[int(5)] = d_1223_v145_
                nw208_[int(6)] = d_1223_v145_
                nw208_[int(7)] = (_dafny.Map({d_1221_v144_: d_1219_v142_})) | (d_1223_v145_)
                nw208_[int(8)] = _dafny.Map({d_1221_v144_: d_1219_v142_})
                nw208_[int(9)] = d_1223_v145_
                nw208_[int(10)] = ((d_1223_v145_).set(d_1221_v144_, d_1219_v142_)) | (d_1223_v145_)
                nw208_[int(11)] = ((d_1223_v145_).set(d_1221_v144_, d_1219_v142_)) | (_dafny.Map({d_1221_v144_: d_1219_v142_}))
                nw208_[int(12)] = (d_1223_v145_) | (d_1223_v145_)
                nw208_[int(13)] = ((d_1223_v145_).set(d_1221_v144_, d_1219_v142_)) | (d_1223_v145_)
                nw208_[int(14)] = d_1223_v145_
                nw208_[int(15)] = d_1223_v145_
                nw208_[int(16)] = d_1223_v145_
                nw208_[int(17)] = _dafny.Map({d_1221_v144_: d_1219_v142_})
                nw208_[int(18)] = d_1223_v145_
                nw208_[int(19)] = d_1223_v145_
                nw208_[int(20)] = d_1223_v145_
                nw208_[int(21)] = (d_1223_v145_) | (d_1223_v145_)
                nw208_[int(22)] = (d_1223_v145_) | (_dafny.Map({d_1221_v144_: d_1219_v142_}))
                nw208_[int(23)] = d_1223_v145_
                nw208_[int(24)] = d_1223_v145_
                nw208_[int(25)] = d_1223_v145_
                d_1225_v147_ = nw208_
                index239_ = default__.safeIndex(257, (d_1225_v147_).length(0))
                (d_1225_v147_)[index239_] = _dafny.Map({d_1221_v144_: d_1219_v142_})
                d_1226_v148_: _dafny.MultiSet
                d_1226_v148_ = _dafny.MultiSet([(d_1219_v142_)[default__.safeIndex(866, (d_1219_v142_).length(0))]])
                d_1227_v149_: _dafny.MultiSet
                d_1227_v149_ = _dafny.MultiSet([p2, p2])
                d_1228_v150_: _dafny.MultiSet
                d_1228_v150_ = d_1227_v149_
                d_1229_v151_: _dafny.Set
                d_1229_v151_ = _dafny.Set({d_1228_v150_})
                index240_ = default__.safeIndex(866, (d_1219_v142_).length(0))
                index241_ = default__.safeIndex(257, (d_1225_v147_).length(0))
                rhs231_ = ((0) - (default__.safeDivisionInt((d_1226_v148_).cardinality, len(d_1229_v151_)))) > (632)
                rhs232_ = (d_1223_v145_) | ((_dafny.Map({d_1221_v144_: d_1219_v142_})) | (_dafny.Map({d_1221_v144_: d_1219_v142_})))
                lhs185_ = d_1219_v142_
                lhs186_ = default__.safeIndex(866, (d_1219_v142_).length(0))
                lhs187_ = d_1225_v147_
                lhs188_ = default__.safeIndex(257, (d_1225_v147_).length(0))
                lhs185_[lhs186_] = rhs231_
                lhs187_[lhs188_] = rhs232_
                index242_ = default__.safeIndex(420, (d_1221_v144_).length(0))
                (d_1221_v144_)[index242_] = default__.safeModuloInt(p0, p0)
                index243_ = default__.safeIndex(420, (d_1221_v144_).length(0))
                (d_1221_v144_)[index243_] = len(((d_1034_v2_) + (d_1034_v2_)) + (d_1034_v2_))
            d_1230_v152_: _dafny.MultiSet
            d_1230_v152_ = _dafny.MultiSet([(default__.fm27(d_1032_v0_, globalState)).cardinality, default__.safeModuloInt(p2, p2), ((self).fm15(len(p1), globalState)) * (p2), (p2) - (603), (p2) * (-108)])
            (globalState).f2 = d_1230_v152_
            d_1231_v153_: _dafny.MultiSet
            d_1231_v153_ = _dafny.MultiSet([(d_1230_v152_) - (d_1230_v152_)])
            d_1232_v154_: _dafny.Array
            nw209_ = _dafny.Array(int(0), 25)
            d_1232_v154_ = nw209_
            index244_ = default__.safeIndex(182, (d_1232_v154_).length(0))
            (d_1232_v154_)[index244_] = p2
            index245_ = default__.safeIndex(182, (d_1232_v154_).length(0))
            rhs233_ = (0) - ((p2 if d_1032_v0_ else (self).fm15(p2, globalState)))
            rhs234_ = (d_1034_v2_) + (d_1034_v2_)
            rhs235_ = (d_1231_v153_).set(d_1230_v152_, default__.abs(p2))
            rhs236_ = p2
            lhs189_ = globalState
            lhs190_ = d_1232_v154_
            lhs191_ = default__.safeIndex(182, (d_1232_v154_).length(0))
            lhs189_.f3 = rhs233_
            d_1034_v2_ = rhs234_
            d_1231_v153_ = rhs235_
            lhs190_[lhs191_] = rhs236_
        elif True:
            d_1233___mcc_h8_ = source20_.cf18
            d_1234_cf18_ = d_1233___mcc_h8_
            (globalState).f3 = p0
            d_1235_v155_: _dafny.Map
            d_1235_v155_ = _dafny.Map({default__.safeModuloInt(p0, (0) - (p0)): (d_1130_v86_) + (_dafny.SeqWithoutIsStrInference([d_1032_v0_, d_1032_v0_]))})
            d_1235_v155_ = (d_1235_v155_).set(p2, d_1130_v86_)
            (globalState).f4 = d_1032_v0_
            d_1236_v156_: C2
            nw210_ = C2()
            nw210_.ctor__()
            d_1236_v156_ = nw210_
        d_1237_v157_: _dafny.Array
        nw211_ = _dafny.Array(False, 9)
        d_1237_v157_ = nw211_
        d_1237_v157_ = d_1237_v157_
        d_1238_i8_: int
        d_1238_i8_ = 0
        with _dafny.label("7"):
            while ((p1)[p0] if (p0) in (p1) else (d_1130_v86_)[default__.safeIndex(p2, len(d_1130_v86_))]):
                with _dafny.c_label("7"):
                    if (d_1238_i8_) >= (100):
                        raise _dafny.Break("7")
                    d_1238_i8_ = (d_1238_i8_) + (1)
                    d_1239_v158_: _dafny.Map
                    d_1239_v158_ = _dafny.Map({d_1034_v2_: len(d_1034_v2_)})
                    d_1239_v158_ = (d_1239_v158_).set(d_1034_v2_, (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bly"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bly")))), d_1033_v1_)) if d_1032_v0_ else (0) - (p2)))
                    d_1240_v159_: _dafny.Set
                    d_1240_v159_ = _dafny.Set({p2, -6, p0})
                    d_1241_v160_: _dafny.Map
                    d_1241_v160_ = _dafny.Map({d_1033_v1_: d_1240_v159_})
                    d_1241_v160_ = (d_1241_v160_).set(d_1033_v1_, (d_1240_v159_).intersection(_dafny.Set({p2})))
                    d_1242_v161_: C0
                    nw212_ = C0()
                    nw212_.ctor__()
                    d_1242_v161_ = nw212_
                    (globalState).f4 = (len(p1)) <= (p2)
                    pass
            pass
        d_1243_v163_: D13
        def iife186_():
            coll80_ = _dafny.Set()
            compr_80_: _dafny.Seq
            for compr_80_ in (default__.fm41(d_1130_v86_, d_1033_v1_, p0, globalState)).Elements:
                d_1244_v162_: _dafny.Seq = compr_80_
                if (d_1244_v162_) in (default__.fm41(d_1130_v86_, d_1033_v1_, p0, globalState)):
                    coll80_ = coll80_.union(_dafny.Set([d_1244_v162_]))
            return _dafny.Set(coll80_)
        d_1243_v163_ = D13_DC29(iife186_()
)
        pat_let_tv58_ = d_1033_v1_
        pat_let_tv59_ = p2
        pat_let_tv60_ = d_1032_v0_
        pat_let_tv61_ = p0
        pat_let_tv62_ = d_1032_v0_
        def lambda60_(source23_):
            if source23_.is_DC30:
                d_1245___mcc_h16_ = source23_.cf42
                d_1246___mcc_h17_ = source23_.cf43
                d_1247___mcc_h18_ = source23_.cf44
                d_1248___mcc_h19_ = source23_.cf45
                d_1249_cf45_ = d_1248___mcc_h19_
                d_1250_cf44_ = d_1247___mcc_h18_
                d_1251_cf43_ = d_1246___mcc_h17_
                d_1252_cf42_ = d_1245___mcc_h16_
                return D5_DC14(len(_dafny.SeqWithoutIsStrInference([pat_let_tv58_ for d_1253_i9_ in range(default__.abs(145))])), not(d_1252_cf42_))
            elif source23_.is_DC29:
                d_1254___mcc_h20_ = source23_.cf41
                d_1255_cf41_ = d_1254___mcc_h20_
                return D5_DC14(pat_let_tv59_, pat_let_tv60_)
            elif True:
                d_1256___mcc_h21_ = source23_.cf46
                d_1257_cf46_ = d_1256___mcc_h21_
                return D5_DC14(pat_let_tv61_, pat_let_tv62_)

        source22_ = lambda60_(d_1243_v163_)
        if source22_.is_DC14:
            d_1258___mcc_h13_ = source22_.cf20
            d_1259___mcc_h14_ = source22_.cf21
            d_1260_cf21_ = d_1259___mcc_h14_
            d_1261_cf20_ = d_1258___mcc_h13_
            d_1032_v0_ = d_1260_cf21_
            d_1262_v164_: _dafny.Map
            d_1262_v164_ = _dafny.Map({909: not (False) or (d_1260_cf21_)})
            d_1262_v164_ = (d_1262_v164_).set(d_1261_cf20_, False)
            d_1263_v165_: _dafny.MultiSet
            d_1263_v165_ = _dafny.MultiSet([-720, p2, p0, d_1261_cf20_])
            (globalState).f2 = d_1263_v165_
            d_1264_v166_: _dafny.Set
            d_1264_v166_ = _dafny.Set({(self).fm16(d_1032_v0_, globalState)})
            d_1265_v167_: _dafny.Seq
            d_1265_v167_ = _dafny.SeqWithoutIsStrInference([p2, p2, len(d_1264_v166_)])
            d_1266_v168_: _dafny.MultiSet
            d_1266_v168_ = _dafny.MultiSet([(d_1032_v0_ if (self).fm3(d_1265_v167_, d_1263_v165_, d_1032_v0_, globalState) else d_1260_cf21_), d_1260_cf21_, d_1260_cf21_, (d_1130_v86_)[default__.safeIndex(p0, len(d_1130_v86_))], d_1032_v0_])
            d_1267_v169_: D2
            d_1267_v169_ = D2_DC7(d_1033_v1_, False, p0)
            rhs237_ = (d_1266_v168_) - (d_1266_v168_)
            rhs238_ = (d_1267_v169_).cf11
            rhs239_ = (self).fm15(default__.safeModuloInt(p2, 824), globalState)
            rhs240_ = (d_1264_v166_ if d_1032_v0_ else d_1264_v166_)
            lhs192_ = globalState
            d_1266_v168_ = rhs237_
            lhs192_.f3 = rhs238_
            d_1261_cf20_ = rhs239_
            d_1264_v166_ = rhs240_
        elif source22_.is_DC15:
            d_1268_v170_: C0
            nw213_ = C0()
            nw213_.ctor__()
            d_1268_v170_ = nw213_
            d_1269_v171_: _dafny.Seq
            d_1269_v171_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1270_v172_: _dafny.Map
            d_1270_v172_ = _dafny.Map({p0: d_1130_v86_})
            (globalState).f4 = (default__.safeDivisionInt(len(d_1269_v171_), p2)) != ((0) - ((0) - ((p0) * (len(((d_1270_v172_)[(d_1268_v170_).fm19(globalState)] if ((d_1268_v170_).fm19(globalState)) in (d_1270_v172_) else d_1130_v86_))))))
            d_1271_v173_: D3
            d_1271_v173_ = D3_DC8((self).fm2(globalState))
            d_1272_v174_: _dafny.MultiSet
            d_1272_v174_ = _dafny.MultiSet([d_1130_v86_, d_1130_v86_])
            (globalState).f3 = (0) - (default__.fm26(d_1271_v173_, p0, d_1272_v174_, globalState))
            d_1273_v175_: C2
            nw214_ = C2()
            nw214_.ctor__()
            d_1273_v175_ = nw214_
        elif True:
            d_1274___mcc_h15_ = source22_.cf19
            d_1275_cf19_ = d_1274___mcc_h15_
            (globalState).f4 = (d_1034_v2_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))
            d_1276_v176_: _dafny.Array
            nw215_ = _dafny.Array(D9.default()(), 3)
            d_1276_v176_ = nw215_
            d_1277_v177_: _dafny.Map
            d_1277_v177_ = _dafny.Map({d_1276_v176_: p0})
            d_1277_v177_ = ((d_1277_v177_) | (d_1277_v177_)) | (d_1277_v177_)
            d_1278_v178_: _dafny.Array
            def lambda61_(d_1279_p2_):
                def lambda62_(d_1280_i10_):
                    return (d_1280_i10_) + (d_1279_p2_)

                return lambda62_

            init30_ = lambda61_(p2)
            nw216_ = _dafny.Array(None, 17)
            for i0_30_ in range(nw216_.length(0)):
                nw216_[i0_30_] = init30_(i0_30_)
            d_1278_v178_ = nw216_
            index246_ = default__.safeIndex(165, (d_1278_v178_).length(0))
            (d_1278_v178_)[index246_] = 899
            index247_ = default__.safeIndex(165, (d_1278_v178_).length(0))
            rhs241_ = d_1032_v0_
            rhs242_ = d_1278_v178_
            rhs243_ = (0) - ((0) - ((p0 if True else p0)))
            lhs193_ = globalState
            lhs194_ = d_1278_v178_
            lhs195_ = default__.safeIndex(165, (d_1278_v178_).length(0))
            lhs193_.f4 = rhs241_
            d_1278_v178_ = rhs242_
            lhs194_[lhs195_] = rhs243_
            d_1281_v179_: _dafny.MultiSet
            d_1281_v179_ = _dafny.MultiSet([d_1032_v0_, d_1032_v0_])
            d_1282_v180_: _dafny.Map
            d_1282_v180_ = _dafny.Map({d_1033_v1_: d_1032_v0_})
            d_1283_v181_: _dafny.Seq
            d_1283_v181_ = _dafny.SeqWithoutIsStrInference([d_1282_v180_])
            d_1284_v182_: _dafny.MultiSet
            d_1284_v182_ = _dafny.MultiSet([d_1032_v0_, d_1032_v0_, d_1032_v0_, ((_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1033_v1_: d_1032_v0_})])).set(default__.safeIndex(((d_1281_v179_)[d_1032_v0_] if (d_1032_v0_) in (d_1281_v179_) else (d_1278_v178_)[default__.safeIndex(165, (d_1278_v178_).length(0))]), len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1033_v1_: d_1032_v0_})]))), d_1282_v180_)) < (d_1283_v181_)])
            d_1281_v179_ = default__.fm0(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldambf"))), (d_1278_v178_)[default__.safeIndex(165, (d_1278_v178_).length(0))]), globalState)
        d_1285_i11_: int
        d_1285_i11_ = 0
        with _dafny.label("8"):
            while not((387) in (default__.fm27(d_1032_v0_, globalState))):
                with _dafny.c_label("8"):
                    if (d_1285_i11_) >= (100):
                        raise _dafny.Break("8")
                    d_1285_i11_ = (d_1285_i11_) + (1)
                    (globalState).f3 = p2
                    d_1286_v183_: _dafny.MultiSet
                    d_1286_v183_ = _dafny.MultiSet([p1])
                    (globalState).f8 = ((d_1286_v183_).cardinality) + (p0)
                    index248_ = default__.safeIndex(483, (d_1237_v157_).length(0))
                    (d_1237_v157_)[index248_] = True
                    d_1287_v184_: D9
                    d_1287_v184_ = D9_DC21(d_1032_v0_, d_1130_v86_, d_1032_v0_)
                    index249_ = default__.safeIndex(483, (d_1237_v157_).length(0))
                    (d_1237_v157_)[index249_] = (d_1287_v184_).cf28
                    (globalState).f8 = p2
                    pass
            pass
        d_1288_v185_: D3
        d_1288_v185_ = D3_DC9()
        r0 = d_1288_v185_
        return r0


class C4(T0):
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f12):
        (self)._f12 = f12

    def fm2(self, globalState):
        if (_dafny.Set({False})) != (_dafny.Set({False, True, True})):
            def iife187_():
                coll81_ = _dafny.Map()
                compr_81_: int
                for compr_81_ in _dafny.IntegerRange(356, 2):
                    d_1290_v0_: int = compr_81_
                    if ((356) <= (d_1290_v0_)) and ((d_1290_v0_) < (2)):
                        coll81_[(d_1290_v0_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1291_i1_ in range(default__.abs(633))])))] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), 742])).cardinality})), len(_dafny.Set({759}))]))
                return _dafny.Map(coll81_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([206 for d_1289_i0_ in range(default__.abs(794))])): len(_dafny.Map({True: _dafny.CodePoint('l')}))})) | (iife187_()
            )
        elif True:
            return (_dafny.Map({-686: (0) - (len(_dafny.Map({959: len(_dafny.Map({len(_dafny.Map({len(_dafny.Map({False: True})): not(True)})): False}))})))})) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-167])), 717}))}))

    def fm3(self, p0, p1, p2, globalState):
        return ((False if True else True)) or ((len(_dafny.Map({264: True}))) < (-407))

    def fm12(self, p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([825])) + (_dafny.SeqWithoutIsStrInference([(0) - (-928), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpglxhpdq")))), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bw")))), len(_dafny.Map({not(not(True)): True}))]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trm"))))

    def fm13(self, p0, p1, p2, globalState):
        source24_ = D0_DC3(D0_DC0((_dafny.MultiSet([not(False), True])).cardinality))
        if source24_.is_DC1:
            d_1292___mcc_h0_ = source24_.cf1
            d_1293___mcc_h1_ = source24_.cf2
            d_1294_cf2_ = d_1293___mcc_h1_
            d_1295_cf1_ = d_1292___mcc_h0_
            return (d_1295_cf1_) >= (len(_dafny.Set({not(True), True, False, not(False)})))
        elif source24_.is_DC2:
            return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kammnnur")))) < (len(_dafny.Map({_dafny.Set({True, not(not(True)), not(False)}): False})))
        elif source24_.is_DC0:
            d_1296___mcc_h2_ = source24_.cf0
            d_1297_cf0_ = d_1296___mcc_h2_
            return not((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])).isdisjoint(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(D2_DC6(d_1297_cf0_, True, True)).cf8, False]), _dafny.SeqWithoutIsStrInference([False, True, True, True])])))
        elif True:
            d_1298___mcc_h3_ = source24_.cf3
            d_1299_cf3_ = d_1298___mcc_h3_
            return False

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_1300_v0_: _dafny.Array
        def lambda63_(d_1301_i0_):
            return (_dafny.MultiSet([True])) == (_dafny.MultiSet([True, True, False]))

        init31_ = lambda63_
        nw217_ = _dafny.Array(None, 23)
        for i0_31_ in range(nw217_.length(0)):
            nw217_[i0_31_] = init31_(i0_31_)
        d_1300_v0_ = nw217_
        d_1302_v1_: bool
        d_1302_v1_ = False
        index250_ = default__.safeIndex(660, (d_1300_v0_).length(0))
        (d_1300_v0_)[index250_] = (d_1302_v1_) == (not(not(not(d_1302_v1_))))
        d_1303_v2_: _dafny.Seq
        d_1303_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upyxghc"))
        d_1304_v3_: int
        d_1304_v3_ = -263
        d_1305_v4_: _dafny.Seq
        d_1305_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm14(d_1304_v3_, globalState), d_1303_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjftgkl"))])
        d_1306_v5_: _dafny.MultiSet
        d_1306_v5_ = _dafny.MultiSet([d_1303_v2_])
        index251_ = default__.safeIndex(660, (d_1300_v0_).length(0))
        (d_1300_v0_)[index251_] = ((_dafny.MultiSet([d_1303_v2_, d_1303_v2_])) | (_dafny.MultiSet(d_1305_v4_))).ispropersubset(d_1306_v5_)
        d_1303_v2_ = d_1303_v2_
        hi6_ = d_1304_v3_
        for d_1307_i1_ in range((d_1304_v3_) - (-637), hi6_):
            (globalState).f4 = d_1302_v1_
            d_1308_v6_: C1
            nw218_ = C1()
            nw218_.ctor__()
            d_1308_v6_ = nw218_
            (globalState).f4 = ((d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]) or (d_1302_v1_)
            d_1300_v0_ = d_1300_v0_
        d_1309_v7_: C3
        nw219_ = C3()
        nw219_.ctor__()
        d_1309_v7_ = nw219_
        if (d_1304_v3_) < ((_dafny.MultiSet([d_1304_v3_, 663, (0) - (d_1304_v3_)])).cardinality):
            d_1310_v8_: C1
            nw220_ = C1()
            nw220_.ctor__()
            d_1310_v8_ = nw220_
            (globalState).f4 = d_1302_v1_
            d_1311_v9_: _dafny.Map
            d_1311_v9_ = _dafny.Map({default__.safeDivisionInt(d_1304_v3_, default__.fm20(d_1302_v1_, globalState)): d_1302_v1_})
            index252_ = default__.safeIndex(660, (d_1300_v0_).length(0))
            (d_1300_v0_)[index252_] = ((d_1311_v9_)[d_1304_v3_] if (d_1304_v3_) in (d_1311_v9_) else True)
            d_1312_v10_: _dafny.MultiSet
            d_1312_v10_ = _dafny.MultiSet([(d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]])
            d_1313_v11_: _dafny.Map
            d_1313_v11_ = _dafny.Map({d_1300_v0_: default__.safeDivisionInt((d_1312_v10_).cardinality, d_1304_v3_)})
            d_1314_v12_: D2
            d_1314_v12_ = D2_DC5(d_1300_v0_)
            d_1315_v13_: _dafny.Map
            d_1315_v13_ = _dafny.Map({not((d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]): d_1304_v3_})
            d_1316_v14_: _dafny.Array
            def lambda64_(d_1317_v3_):
                def lambda65_(d_1318_i2_):
                    return default__.safeDivisionInt(d_1318_i2_, d_1317_v3_)

                return lambda65_

            init32_ = lambda64_(d_1304_v3_)
            nw221_ = _dafny.Array(None, 19)
            for i0_32_ in range(nw221_.length(0)):
                nw221_[i0_32_] = init32_(i0_32_)
            d_1316_v14_ = nw221_
            d_1319_v15_: _dafny.Map
            d_1319_v15_ = _dafny.Map({((d_1315_v13_)[(d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]] if ((d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]) in (d_1315_v13_) else d_1304_v3_): d_1316_v14_})
            d_1313_v11_ = (d_1313_v11_).set((d_1314_v12_).cf5, len(d_1319_v15_))
            (globalState).f3 = default__.safeModuloInt(len(d_1303_v2_), d_1304_v3_)
        elif True:
            d_1320_v16_: _dafny.Array
            nw222_ = _dafny.Array(int(0), 6)
            d_1320_v16_ = nw222_
            d_1321_v17_: _dafny.Map
            d_1321_v17_ = _dafny.Map({d_1304_v3_: (d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]})
            index253_ = default__.safeIndex(625, (d_1320_v16_).length(0))
            (d_1320_v16_)[index253_] = len(d_1321_v17_)
            index254_ = default__.safeIndex(625, (d_1320_v16_).length(0))
            (d_1320_v16_)[index254_] = (975) + ((d_1304_v3_) * (d_1304_v3_))
            (globalState).f4 = ((d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]) == (False)
            (globalState).f3 = d_1304_v3_
            d_1322_v18_: _dafny.Set
            d_1322_v18_ = _dafny.Set({207})
            d_1323_v19_: C0
            nw223_ = C0()
            nw223_.ctor__()
            d_1323_v19_ = nw223_
            d_1324_v20_: D13
            d_1324_v20_ = D13_DC30((d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))], (d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))], d_1322_v18_, d_1323_v19_)
            d_1325_v21_: D13
            d_1325_v21_ = D13_DC31(d_1324_v20_)
            d_1326_v23_: D0
            d_1326_v23_ = D0_DC1(d_1304_v3_, len(d_1303_v2_))
            def iife188_():
                coll82_ = _dafny.Set()
                compr_82_: int
                for compr_82_ in _dafny.IntegerRange(756, 690):
                    d_1327_v22_: int = compr_82_
                    if ((756) <= (d_1327_v22_)) and ((d_1327_v22_) < (690)):
                        coll82_ = coll82_.union(_dafny.Set([(d_1327_v22_) * (d_1304_v3_)]))
                return _dafny.Set(coll82_)
            d_1325_v21_ = ((d_1325_v21_ if d_1302_v1_ else d_1325_v21_) if (iife188_()
            ).isdisjoint(_dafny.Set({(d_1326_v23_).cf1})) else d_1325_v21_)
            if True:
                d_1328_v24_: _dafny.Seq
                d_1328_v24_ = _dafny.SeqWithoutIsStrInference([(d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]])
                rhs244_ = (d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]
                rhs245_ = d_1304_v3_
                rhs246_ = ((d_1328_v24_) + (d_1328_v24_)) + (d_1328_v24_)
                lhs196_ = globalState
                lhs197_ = globalState
                lhs196_.f4 = rhs244_
                lhs197_.f3 = rhs245_
                d_1328_v24_ = rhs246_
                d_1302_v1_ = (len(d_1328_v24_)) < (len((d_1303_v2_) + (default__.fm25(globalState))))
                d_1302_v1_ = d_1302_v1_
                d_1329_v25_: _dafny.Seq
                d_1329_v25_ = _dafny.SeqWithoutIsStrInference([d_1304_v3_, (d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))]])
                d_1330_v26_: _dafny.Map
                d_1330_v26_ = _dafny.Map({d_1300_v0_: d_1329_v25_})
                d_1330_v26_ = (d_1330_v26_).set(d_1300_v0_, (_dafny.SeqWithoutIsStrInference([(d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))] for d_1331_i3_ in range(default__.abs(629))])) + (d_1329_v25_))
                d_1332_v27_: _dafny.Array
                def lambda66_(d_1333_i4_):
                    return _dafny.CodePoint('g')

                init33_ = lambda66_
                nw224_ = _dafny.Array(None, 19)
                for i0_33_ in range(nw224_.length(0)):
                    nw224_[i0_33_] = init33_(i0_33_)
                d_1332_v27_ = nw224_
                d_1334_v28_: str
                d_1334_v28_ = _dafny.CodePoint('w')
                index255_ = default__.safeIndex(872, (d_1332_v27_).length(0))
                (d_1332_v27_)[index255_] = d_1334_v28_
                index256_ = default__.safeIndex(872, (d_1332_v27_).length(0))
                (d_1332_v27_)[index256_] = d_1334_v28_
            elif True:
                index257_ = default__.safeIndex(625, (d_1320_v16_).length(0))
                (d_1320_v16_)[index257_] = (default__.safeDivisionInt((d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))], (d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))])) - ((d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))])
                d_1335_v29_: _dafny.Seq
                d_1335_v29_ = _dafny.SeqWithoutIsStrInference([not(d_1302_v1_)])
                rhs247_ = d_1300_v0_
                rhs248_ = ((d_1335_v29_) + ((d_1335_v29_).set(default__.safeIndex(d_1304_v3_, len(d_1335_v29_)), d_1302_v1_))) < ((d_1335_v29_) + (d_1335_v29_))
                rhs249_ = d_1303_v2_
                lhs198_ = globalState
                d_1300_v0_ = rhs247_
                lhs198_.f4 = rhs248_
                d_1303_v2_ = rhs249_
                d_1336_v30_: bool
                d_1336_v30_ = d_1302_v1_
                index258_ = default__.safeIndex(625, (d_1320_v16_).length(0))
                (d_1320_v16_)[index258_] = (self).fm12((len(d_1303_v2_)) <= ((0) - ((d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))])), d_1304_v3_, d_1303_v2_, d_1336_v30_, globalState)
                d_1337_v31_: _dafny.Seq
                d_1337_v31_ = _dafny.SeqWithoutIsStrInference([(d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))], d_1304_v3_])
                d_1338_v32_: _dafny.Map
                d_1338_v32_ = _dafny.Map({d_1320_v16_: len(d_1337_v31_)})
                d_1339_v33_: _dafny.Map
                d_1339_v33_ = _dafny.Map({d_1304_v3_: len(d_1338_v32_)})
                (globalState).f8 = default__.safeDivisionInt((d_1320_v16_)[default__.safeIndex(625, (d_1320_v16_).length(0))], len((_dafny.Map({d_1304_v3_: d_1304_v3_})) | (d_1339_v33_)))
                d_1340_v34_: C2
                nw225_ = C2()
                nw225_.ctor__()
                d_1340_v34_ = nw225_
        d_1341_v35_: bool
        d_1341_v35_ = d_1302_v1_
        index259_ = default__.safeIndex(660, (d_1300_v0_).length(0))
        rhs250_ = d_1302_v1_
        rhs251_ = (d_1303_v2_ if (d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtakwmgb")))
        rhs252_ = (d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]
        rhs253_ = not((d_1341_v35_))
        rhs254_ = ((d_1300_v0_)[default__.safeIndex(660, (d_1300_v0_).length(0))]) or (d_1302_v1_)
        lhs199_ = globalState
        lhs200_ = d_1300_v0_
        lhs201_ = default__.safeIndex(660, (d_1300_v0_).length(0))
        lhs202_ = globalState
        lhs199_.f4 = rhs250_
        d_1303_v2_ = rhs251_
        lhs200_[lhs201_] = rhs252_
        r1 = rhs253_
        lhs202_.f4 = rhs254_
        r0 = d_1302_v1_
        r1 = (d_1304_v3_) < (-913)
        return r0, r1

    def m9(self, p0, p1, globalState):
        d_1342_v0_: _dafny.Array
        def lambda67_(d_1343_i0_):
            return (False) not in (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjxvcpd"))}))

        init34_ = lambda67_
        nw226_ = _dafny.Array(None, 2)
        for i0_34_ in range(nw226_.length(0)):
            nw226_[i0_34_] = init34_(i0_34_)
        d_1342_v0_ = nw226_
        index260_ = default__.safeIndex(364, (d_1342_v0_).length(0))
        (d_1342_v0_)[index260_] = not(False)
        d_1344_v1_: bool
        d_1344_v1_ = True
        index261_ = default__.safeIndex(364, (d_1342_v0_).length(0))
        (d_1342_v0_)[index261_] = d_1344_v1_
        (globalState).f8 = len(default__.fm31(d_1344_v1_, ((d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))] if (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))] else not(d_1344_v1_)), d_1344_v1_, globalState))
        hi7_ = p1
        for d_1345_i1_ in range(p0, hi7_):
            d_1346_v2_: _dafny.Seq
            d_1346_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koegpi"))
            d_1347_v3_: str
            d_1347_v3_ = _dafny.CodePoint('e')
            d_1348_v4_: _dafny.Seq
            d_1348_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smk"))])
            d_1349_v5_: C0
            nw227_ = C0()
            nw227_.ctor__()
            d_1349_v5_ = nw227_
            d_1350_v6_: _dafny.Seq
            d_1350_v6_ = _dafny.SeqWithoutIsStrInference([d_1349_v5_])
            d_1351_v7_: _dafny.Set
            d_1351_v7_ = _dafny.Set({d_1346_v2_, (d_1348_v4_)[default__.safeIndex((_dafny.MultiSet(d_1350_v6_)).cardinality, len(d_1348_v4_))]})
            d_1352_v8_: _dafny.Map
            d_1352_v8_ = _dafny.Map({(((d_1346_v2_).set(default__.safeIndex(p0, len(d_1346_v2_)), d_1347_v3_)).set(default__.safeIndex(len(_dafny.Set({d_1344_v1_})), len((d_1346_v2_).set(default__.safeIndex(p0, len(d_1346_v2_)), d_1347_v3_))), d_1347_v3_)) != (d_1346_v2_): d_1351_v7_})
            d_1353_v9_: _dafny.MultiSet
            d_1353_v9_ = _dafny.MultiSet([len(d_1346_v2_), d_1345_i1_, p0])
            d_1354_v10_: _dafny.Set
            d_1354_v10_ = _dafny.Set({d_1353_v9_})
            d_1352_v8_ = (d_1352_v8_).set((d_1354_v10_).issubset(d_1354_v10_), (d_1351_v7_) - (d_1351_v7_))
            d_1355_v11_: _dafny.Map
            d_1355_v11_ = _dafny.Map({(d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))]: (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))]})
            d_1356_v12_: _dafny.Seq
            d_1356_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({not(d_1344_v1_): (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))]}), d_1355_v11_])
            index262_ = default__.safeIndex(364, (d_1342_v0_).length(0))
            index263_ = default__.safeIndex(364, (d_1342_v0_).length(0))
            index264_ = default__.safeIndex(364, (d_1342_v0_).length(0))
            rhs255_ = not((d_1355_v11_) == ((d_1356_v12_)[default__.safeIndex(d_1345_i1_, len(d_1356_v12_))]))
            rhs256_ = (d_1346_v2_).set(default__.safeIndex(d_1345_i1_, len(d_1346_v2_)), d_1347_v3_)
            rhs257_ = d_1344_v1_
            rhs258_ = (d_1353_v9_).isdisjoint((d_1353_v9_).intersection(d_1353_v9_))
            rhs259_ = (p1) > (len(d_1346_v2_))
            lhs203_ = globalState
            lhs204_ = d_1342_v0_
            lhs205_ = default__.safeIndex(364, (d_1342_v0_).length(0))
            lhs206_ = d_1342_v0_
            lhs207_ = default__.safeIndex(364, (d_1342_v0_).length(0))
            lhs208_ = d_1342_v0_
            lhs209_ = default__.safeIndex(364, (d_1342_v0_).length(0))
            lhs203_.f4 = rhs255_
            d_1346_v2_ = rhs256_
            lhs204_[lhs205_] = rhs257_
            lhs206_[lhs207_] = rhs258_
            lhs208_[lhs209_] = rhs259_
            (globalState).f4 = (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))]
            d_1347_v3_ = d_1347_v3_
        d_1357_v13_: D12
        d_1357_v13_ = D12_DC28(D12_DC27(p0, p1, 1, (0) - (p1)))
        d_1358_v14_: D12
        d_1358_v14_ = D12_DC28(d_1357_v13_)
        source25_ = D12_DC28(d_1357_v13_)
        if source25_.is_DC27:
            d_1359___mcc_h0_ = source25_.cf36
            d_1360___mcc_h1_ = source25_.cf37
            d_1361___mcc_h2_ = source25_.cf38
            d_1362___mcc_h3_ = source25_.cf39
            d_1363_cf39_ = d_1362___mcc_h3_
            d_1364_cf38_ = d_1361___mcc_h2_
            d_1365_cf37_ = d_1360___mcc_h1_
            d_1366_cf36_ = d_1359___mcc_h0_
            d_1363_cf39_ = (0) - (d_1366_cf36_)
            (globalState).f8 = default__.fm20((d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))], globalState)
            d_1367_v15_: _dafny.Seq
            d_1367_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "io"))
            d_1364_cf38_ = default__.safeDivisionInt((0) - (d_1363_cf39_), len((d_1367_v15_) + (d_1367_v15_)))
            d_1368_v16_: C3
            nw228_ = C3()
            nw228_.ctor__()
            d_1368_v16_ = nw228_
        elif source25_.is_DC26:
            d_1369___mcc_h4_ = source25_.cf35
            d_1370_cf35_ = d_1369___mcc_h4_
            d_1371_v17_: _dafny.Array
            def lambda68_(d_1372_v0_):
                def lambda69_(d_1373_i2_):
                    return default__.safeDivisionInt(d_1373_i2_, (_dafny.MultiSet([(d_1372_v0_)[default__.safeIndex(364, (d_1372_v0_).length(0))], (d_1372_v0_)[default__.safeIndex(364, (d_1372_v0_).length(0))]])).cardinality)

                return lambda69_

            init35_ = lambda68_(d_1342_v0_)
            nw229_ = _dafny.Array(None, 26)
            for i0_35_ in range(nw229_.length(0)):
                nw229_[i0_35_] = init35_(i0_35_)
            d_1371_v17_ = nw229_
            d_1371_v17_ = d_1371_v17_
            d_1374_v18_: _dafny.Array
            def lambda70_(d_1375_p1_):
                def lambda71_(d_1376_i3_):
                    return D0_DC0(d_1375_p1_)

                return lambda71_

            init36_ = lambda70_(p1)
            nw230_ = _dafny.Array(None, 20)
            for i0_36_ in range(nw230_.length(0)):
                nw230_[i0_36_] = init36_(i0_36_)
            d_1374_v18_ = nw230_
            d_1377_v19_: _dafny.Seq
            d_1377_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ag"))
            d_1378_v20_: _dafny.Seq
            d_1378_v20_ = _dafny.SeqWithoutIsStrInference([p0, len(d_1377_v19_)])
            d_1379_v21_: _dafny.Seq
            d_1379_v21_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, (d_1378_v20_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1380_i4_ in range(default__.abs(555))])), len(d_1378_v20_))]])
            d_1381_v22_: D0
            d_1381_v22_ = D0_DC0(len(d_1379_v21_))
            index265_ = default__.safeIndex(867, (d_1374_v18_).length(0))
            (d_1374_v18_)[index265_] = d_1381_v22_
            index266_ = default__.safeIndex(867, (d_1374_v18_).length(0))
            (d_1374_v18_)[index266_] = d_1381_v22_
            d_1382_v23_: _dafny.MultiSet
            d_1382_v23_ = _dafny.MultiSet([(d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))], d_1344_v1_])
            d_1383_v24_: _dafny.Map
            d_1383_v24_ = _dafny.Map({p0: d_1382_v23_})
            (globalState).f8 = (0) - (((d_1382_v23_)[(d_1382_v23_).isdisjoint(((d_1383_v24_)[p1] if (p1) in (d_1383_v24_) else d_1382_v23_))] if ((d_1382_v23_).isdisjoint(((d_1383_v24_)[p1] if (p1) in (d_1383_v24_) else d_1382_v23_))) in (d_1382_v23_) else p0))
            d_1371_v17_ = d_1371_v17_
        elif True:
            d_1384___mcc_h5_ = source25_.cf40
            d_1385_cf40_ = d_1384___mcc_h5_
            (globalState).f3 = p0
            (globalState).f8 = -323
            (globalState).f8 = p0
            d_1386_v25_: _dafny.Seq
            d_1386_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdikjo"))
            d_1387_v26_: str
            d_1387_v26_ = _dafny.CodePoint('w')
            d_1388_v27_: _dafny.Seq
            d_1388_v27_ = _dafny.SeqWithoutIsStrInference([d_1386_v25_])
            d_1389_v28_: _dafny.Seq
            d_1389_v28_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1390_v29_: D10
            d_1390_v29_ = D10_DC24(p1, not(d_1344_v1_), d_1386_v25_)
            d_1391_v30_: _dafny.Array
            nw231_ = _dafny.Array(None, 20)
            nw231_[int(0)] = d_1386_v25_
            nw231_[int(1)] = d_1386_v25_
            nw231_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            nw231_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ceextbwbk"))
            nw231_[int(4)] = (d_1386_v25_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1392_i5_ in range(default__.abs(451))]))
            nw231_[int(5)] = d_1386_v25_
            nw231_[int(6)] = d_1386_v25_
            nw231_[int(7)] = d_1386_v25_
            nw231_[int(8)] = d_1386_v25_
            nw231_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lncxgmpd"))) + (d_1386_v25_)
            nw231_[int(10)] = (d_1386_v25_) + (d_1386_v25_)
            nw231_[int(11)] = d_1386_v25_
            nw231_[int(12)] = d_1386_v25_
            nw231_[int(13)] = default__.fm17(d_1387_v26_, globalState)
            nw231_[int(14)] = (d_1388_v27_)[default__.safeIndex((d_1389_v28_)[default__.safeIndex(p0, len(d_1389_v28_))], len(d_1388_v27_))]
            nw231_[int(15)] = d_1386_v25_
            nw231_[int(16)] = _dafny.SeqWithoutIsStrInference([d_1387_v26_ for d_1393_i6_ in range(default__.abs(-408))])
            nw231_[int(17)] = (d_1390_v29_).cf33
            nw231_[int(18)] = d_1386_v25_
            nw231_[int(19)] = d_1386_v25_
            d_1391_v30_ = nw231_
            index267_ = default__.safeIndex(666, (d_1391_v30_).length(0))
            (d_1391_v30_)[index267_] = (d_1386_v25_) + (d_1386_v25_)
            index268_ = default__.safeIndex(666, (d_1391_v30_).length(0))
            (d_1391_v30_)[index268_] = d_1386_v25_
        d_1394_v31_: _dafny.Array
        def lambda72_(d_1395_p1_, d_1396_v1_):
            def lambda73_(d_1397_i7_):
                return _dafny.Map({d_1395_p1_: d_1396_v1_})

            return lambda73_

        init37_ = lambda72_(p1, d_1344_v1_)
        nw232_ = _dafny.Array(None, 5)
        for i0_37_ in range(nw232_.length(0)):
            nw232_[i0_37_] = init37_(i0_37_)
        d_1394_v31_ = nw232_
        d_1398_v32_: _dafny.MultiSet
        d_1398_v32_ = _dafny.MultiSet([p1, p0])
        index269_ = default__.safeIndex(222, (d_1394_v31_).length(0))
        def iife189_():
            coll83_ = _dafny.Map()
            compr_83_: int
            for compr_83_ in _dafny.IntegerRange(58, 318):
                d_1399_v33_: int = compr_83_
                if ((58) <= (d_1399_v33_)) and ((d_1399_v33_) < (318)):
                    coll83_[default__.safeDivisionInt(d_1399_v33_, p1)] = d_1344_v1_
            return _dafny.Map(coll83_)
        (d_1394_v31_)[index269_] = (_dafny.Map({p0: (self).fm3(_dafny.SeqWithoutIsStrInference([p0]), d_1398_v32_, False, globalState)})) | (iife189_()
        )
        d_1400_v35_: _dafny.Seq
        d_1400_v35_ = _dafny.SeqWithoutIsStrInference([d_1344_v1_])
        d_1401_v36_: _dafny.Map
        d_1401_v36_ = _dafny.Map({-494: (_dafny.MultiSet(d_1400_v35_)).cardinality})
        d_1402_v37_: C3
        nw233_ = C3()
        nw233_.ctor__()
        d_1402_v37_ = nw233_
        d_1403_v38_: D0
        d_1403_v38_ = D0_DC1(p0, len(d_1400_v35_))
        d_1404_v39_: _dafny.Map
        d_1404_v39_ = _dafny.Map({d_1402_v37_: (d_1403_v38_).cf1})
        d_1405_v40_: C3
        d_1405_v40_ = d_1402_v37_
        d_1406_v41_: str
        d_1406_v41_ = _dafny.CodePoint('k')
        index270_ = default__.safeIndex(222, (d_1394_v31_).length(0))
        def iife190_():
            coll84_ = _dafny.Map()
            compr_84_: int
            for compr_84_ in (d_1401_v36_).keys.Elements:
                d_1407_v34_: int = compr_84_
                if (d_1407_v34_) in (d_1401_v36_):
                    coll84_[default__.safeModuloInt(d_1407_v34_, p0)] = d_1344_v1_
            return _dafny.Map(coll84_)
        (d_1394_v31_)[index270_] = (iife190_()
        ).set(((d_1404_v39_)[(d_1405_v40_)] if ((d_1405_v40_)) in (d_1404_v39_) else (0) - ((0) - (len(default__.fm1(p0, d_1406_v41_, (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))], d_1401_v36_, globalState))))), d_1344_v1_)
        d_1408_v42_: _dafny.Seq
        d_1408_v42_ = _dafny.SeqWithoutIsStrInference([default__.fm18((d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))], (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))], globalState), _dafny.CodePoint('a'), d_1406_v41_, d_1406_v41_])
        d_1409_v43_: D10
        d_1409_v43_ = D10_DC24(p0, d_1344_v1_, d_1408_v42_)
        d_1410_v44_: _dafny.Map
        d_1410_v44_ = _dafny.Map({(d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))]: (d_1409_v43_).cf33})
        d_1411_v45_: D2
        d_1411_v45_ = D2_DC7(d_1406_v41_, (d_1342_v0_)[default__.safeIndex(364, (d_1342_v0_).length(0))], p1)
        d_1410_v44_ = (d_1410_v44_).set((d_1411_v45_).cf10, d_1408_v42_)

    @property
    def f12(self):
        return self._f12

class C5(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def fm2(self, globalState):
        def iife191_():
            coll85_ = _dafny.Set()
            compr_85_: _dafny.Set
            for compr_85_ in (_dafny.Map({_dafny.Set({True}): True})).keys.Elements:
                d_1412_v0_: _dafny.Set = compr_85_
                if (d_1412_v0_) in (_dafny.Map({_dafny.Set({True}): True})):
                    coll85_ = coll85_.union(_dafny.Set([d_1412_v0_]))
            return _dafny.Set(coll85_)
        source26_ = D0_DC1(len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet((D4_DC10(_dafny.SeqWithoutIsStrInference([not(False), True]))).cf13)).cardinality, len(iife191_()
), len(_dafny.Set({True, True}))])), 277)
        if source26_.is_DC1:
            d_1413___mcc_h0_ = source26_.cf1
            d_1414___mcc_h1_ = source26_.cf2
            d_1415_cf2_ = d_1414___mcc_h1_
            d_1416_cf1_ = d_1413___mcc_h0_
            if not(not(True)):
                return _dafny.Map({d_1415_cf2_: d_1416_cf1_})
            elif True:
                return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, False])): d_1416_cf1_})
        elif source26_.is_DC2:
            def iife192_():
                coll86_ = _dafny.Map()
                compr_86_: int
                for compr_86_ in _dafny.IntegerRange(532, 298):
                    d_1417_v1_: int = compr_86_
                    if ((532) <= (d_1417_v1_)) and ((d_1417_v1_) < (298)):
                        coll86_[(d_1417_v1_) + (-161)] = -906
                return _dafny.Map(coll86_)
            return (iife192_()
            ) | (_dafny.Map({477: len(_dafny.SeqWithoutIsStrInference([-366 for d_1418_i0_ in range(default__.abs(-948))]))}))
        elif source26_.is_DC0:
            d_1419___mcc_h2_ = source26_.cf0
            d_1420_cf0_ = d_1419___mcc_h2_
            if True:
                return _dafny.Map({d_1420_cf0_: d_1420_cf0_})
            elif True:
                def iife193_():
                    coll87_ = _dafny.Map()
                    compr_87_: int
                    for compr_87_ in (_dafny.Set({d_1420_cf0_, 771, d_1420_cf0_})).Elements:
                        d_1421_v2_: int = compr_87_
                        if (d_1421_v2_) in (_dafny.Set({d_1420_cf0_, 771, d_1420_cf0_})):
                            coll87_[(d_1421_v2_) - (len(_dafny.Map({False: _dafny.CodePoint('k')})))] = d_1420_cf0_
                    return _dafny.Map(coll87_)
                return iife193_()
                
        elif True:
            d_1422___mcc_h3_ = source26_.cf3
            d_1423_cf3_ = d_1422___mcc_h3_
            return (_dafny.Map({len(_dafny.Map({False: True})): -365})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))): 561}))

    def fm3(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvoets"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))])) + (_dafny.SeqWithoutIsStrInference([-419 for d_1424_i0_ in range(default__.abs(-690))]))) == (_dafny.SeqWithoutIsStrInference([994, 515, len(_dafny.SeqWithoutIsStrInference([494 for d_1425_i1_ in range(default__.abs(453))])), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1426_i2_ in range(default__.abs(897))])))]))

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_1427_v0_: C1
        nw234_ = C1()
        nw234_.ctor__()
        d_1427_v0_ = nw234_
        d_1428_v1_: _dafny.Array
        nw235_ = _dafny.Array(D4.default()(), 22)
        d_1428_v1_ = nw235_
        d_1429_v2_: _dafny.Map
        d_1429_v2_ = _dafny.Map({d_1428_v1_: False})
        d_1430_v3_: _dafny.Seq
        d_1430_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyu"))
        d_1431_v4_: _dafny.Set
        d_1431_v4_ = _dafny.Set({d_1430_v3_})
        d_1432_v5_: int
        d_1432_v5_ = 112
        d_1433_v6_: str
        d_1433_v6_ = _dafny.CodePoint('o')
        d_1434_v7_: bool
        d_1434_v7_ = True
        def iife194_():
            coll88_ = _dafny.Set()
            compr_88_: _dafny.Seq
            for compr_88_ in (default__.fm42(d_1432_v5_, d_1432_v5_, len(_dafny.Map({d_1433_v6_: d_1432_v5_})), d_1434_v7_, globalState)).keys.Elements:
                d_1435_v8_: _dafny.Seq = compr_88_
                if (d_1435_v8_) in (default__.fm42(d_1432_v5_, d_1432_v5_, len(_dafny.Map({d_1433_v6_: d_1432_v5_})), d_1434_v7_, globalState)):
                    coll88_ = coll88_.union(_dafny.Set([d_1435_v8_]))
            return _dafny.Set(coll88_)
        d_1429_v2_ = (d_1429_v2_).set(d_1428_v1_, (d_1431_v4_) != (iife194_()
        ))
        d_1436_v9_: _dafny.MultiSet
        d_1436_v9_ = _dafny.MultiSet([d_1434_v7_])
        (globalState).f3 = ((-68) + (d_1432_v5_)) * (((d_1436_v9_) - (d_1436_v9_)).cardinality)
        d_1434_v7_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxdhgpiq"))) < (d_1430_v3_)
        d_1437_v10_: _dafny.Map
        d_1437_v10_ = _dafny.Map({d_1434_v7_: d_1434_v7_})
        d_1438_v11_: _dafny.Map
        d_1438_v11_ = _dafny.Map({87: _dafny.Map({len(d_1437_v10_): d_1433_v6_})})
        d_1439_v12_: _dafny.Map
        d_1439_v12_ = _dafny.Map({d_1432_v5_: _dafny.CodePoint('q')})
        d_1440_v13_: _dafny.Seq
        d_1440_v13_ = _dafny.SeqWithoutIsStrInference([d_1439_v12_, d_1439_v12_, d_1439_v12_, d_1439_v12_])
        d_1441_v15_: _dafny.Map
        d_1441_v15_ = _dafny.Map({-257: 869})
        d_1442_v16_: _dafny.Seq
        d_1442_v16_ = _dafny.SeqWithoutIsStrInference([d_1434_v7_, not(d_1434_v7_)])
        d_1443_v17_: _dafny.Map
        d_1443_v17_ = _dafny.Map({d_1434_v7_: (0) - ((0) - (default__.fm26(D3_DC8(d_1441_v15_), len(d_1430_v3_), _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([d_1434_v7_, d_1434_v7_])).set(default__.safeIndex(d_1432_v5_, len(_dafny.SeqWithoutIsStrInference([d_1434_v7_, d_1434_v7_]))), False), d_1442_v16_]), globalState)))})
        d_1444_v18_: _dafny.Seq
        d_1444_v18_ = _dafny.SeqWithoutIsStrInference([len(d_1430_v3_), d_1432_v5_, (0) - (len(d_1443_v17_)), d_1432_v5_, 340])
        d_1445_v19_: D19
        d_1445_v19_ = D19_DC38(d_1439_v12_)
        d_1446_v21_: _dafny.Map
        d_1446_v21_ = _dafny.Map({len(d_1442_v16_): d_1434_v7_})
        d_1447_v22_: _dafny.Array
        nw236_ = _dafny.Array(None, 20)
        nw236_[int(0)] = ((d_1438_v11_)[8] if (8) in (d_1438_v11_) else _dafny.Map({d_1432_v5_: d_1433_v6_}))
        nw236_[int(1)] = d_1439_v12_
        nw236_[int(2)] = d_1439_v12_
        nw236_[int(3)] = (d_1439_v12_) | (d_1439_v12_)
        nw236_[int(4)] = ((d_1440_v13_)[default__.safeIndex(d_1432_v5_, len(d_1440_v13_))]) | (d_1439_v12_)
        nw236_[int(5)] = d_1439_v12_
        nw236_[int(6)] = (d_1439_v12_) | (d_1439_v12_)
        nw236_[int(7)] = d_1439_v12_
        nw236_[int(8)] = d_1439_v12_
        nw236_[int(9)] = d_1439_v12_
        nw236_[int(10)] = default__.fm43(not(d_1434_v7_), d_1432_v5_, d_1433_v6_, globalState)
        nw236_[int(11)] = d_1439_v12_
        def iife195_():
            coll89_ = _dafny.Map()
            compr_89_: int
            for compr_89_ in (d_1444_v18_).Elements:
                d_1448_v14_: int = compr_89_
                if (d_1448_v14_) in (d_1444_v18_):
                    coll89_[(d_1448_v14_) * (d_1432_v5_)] = d_1433_v6_
            return _dafny.Map(coll89_)
        nw236_[int(12)] = iife195_()
        
        nw236_[int(13)] = d_1439_v12_
        nw236_[int(14)] = d_1439_v12_
        nw236_[int(15)] = d_1439_v12_
        nw236_[int(16)] = (_dafny.Map({d_1432_v5_: d_1433_v6_}) if d_1434_v7_ else (d_1445_v19_).cf57)
        nw236_[int(17)] = (d_1439_v12_) | (d_1439_v12_)
        nw236_[int(18)] = _dafny.Map({d_1432_v5_: d_1433_v6_})
        def iife196_():
            coll90_ = _dafny.Map()
            compr_90_: int
            for compr_90_ in (d_1446_v21_).keys.Elements:
                d_1449_v20_: int = compr_90_
                if (d_1449_v20_) in (d_1446_v21_):
                    coll90_[(d_1449_v20_) + (len(d_1439_v12_))] = _dafny.CodePoint('a')
            return _dafny.Map(coll90_)
        nw236_[int(19)] = ((iife196_()
        ).set(d_1432_v5_, d_1433_v6_)) | (d_1439_v12_)
        d_1447_v22_ = nw236_
        d_1447_v22_ = d_1447_v22_
        d_1450_v23_: D0
        d_1450_v23_ = D0_DC1(d_1432_v5_, d_1432_v5_)
        d_1451_v24_: _dafny.Set
        d_1451_v24_ = _dafny.Set({d_1450_v23_, d_1450_v23_})
        d_1452_v25_: _dafny.MultiSet
        d_1452_v25_ = _dafny.MultiSet([d_1432_v5_])
        d_1453_v26_: D19
        d_1453_v26_ = D19_DC39(d_1434_v7_, d_1432_v5_)
        (d_1427_v0_).m13(len(d_1451_v24_), ((d_1452_v25_).set(len(d_1430_v3_), default__.abs(d_1432_v5_)) if d_1434_v7_ else default__.fm27((d_1453_v26_).cf58, globalState)), d_1434_v7_, globalState)
        r0 = not(True)
        r1 = not((d_1434_v7_ if (d_1432_v5_) <= (d_1432_v5_) else (d_1427_v0_).fm3(d_1444_v18_, d_1452_v25_, d_1434_v7_, globalState)))
        return r0, r1


class C6(T0):
    def  __init__(self):
        self._f10: int = int(0)
        self._f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f10, f11):
        (self)._f10 = f10
        (self)._f11 = f11

    def fm2(self, globalState):
        return _dafny.Map({default__.safeDivisionInt((D0_DC1((self).f10, (self).f10)).cf2, 547): (self).f10})

    def fm3(self, p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True)])) <= (_dafny.SeqWithoutIsStrInference([True, False]))

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_1454_v0_: _dafny.Array
        nw237_ = _dafny.Array(int(0), 9)
        d_1454_v0_ = nw237_
        index271_ = default__.safeIndex(359, (d_1454_v0_).length(0))
        (d_1454_v0_)[index271_] = (self).f10
        d_1455_v1_: bool
        d_1455_v1_ = False
        d_1456_v2_: _dafny.Map
        d_1456_v2_ = _dafny.Map({d_1455_v1_: (self).f10})
        index272_ = default__.safeIndex(359, (d_1454_v0_).length(0))
        (d_1454_v0_)[index272_] = (len((self).f11) if d_1455_v1_ else ((d_1456_v2_)[d_1455_v1_] if (d_1455_v1_) in (d_1456_v2_) else 757))
        hi8_ = 149
        for d_1457_i0_ in range((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))], hi8_):
            d_1454_v0_ = (d_1454_v0_ if d_1455_v1_ else d_1454_v0_)
            d_1458_v3_: T0
            nw238_ = C3()
            nw238_.ctor__()
            d_1458_v3_ = nw238_
            d_1459_v4_: _dafny.Set
            d_1459_v4_ = _dafny.Set({d_1458_v3_})
            d_1459_v4_ = (d_1459_v4_).intersection(d_1459_v4_)
            d_1460_v5_: str
            d_1460_v5_ = _dafny.CodePoint('d')
            d_1461_v6_: _dafny.Map
            d_1461_v6_ = _dafny.Map({(self).f10: d_1457_i0_})
            d_1462_v7_: _dafny.MultiSet
            d_1462_v7_ = _dafny.MultiSet([(self).f10])
            d_1463_v8_: _dafny.Seq
            d_1463_v8_ = _dafny.SeqWithoutIsStrInference([d_1455_v1_, d_1455_v1_, d_1455_v1_, d_1455_v1_, d_1455_v1_])
            d_1464_v9_: _dafny.Seq
            d_1464_v9_ = _dafny.SeqWithoutIsStrInference([len((self).f11), (0) - (d_1457_i0_)])
            d_1465_v10_: _dafny.Map
            d_1465_v10_ = _dafny.Map({not(d_1455_v1_): d_1455_v1_})
            d_1466_v11_: _dafny.Array
            nw239_ = _dafny.Array(None, 23)
            nw239_[int(0)] = d_1455_v1_
            nw239_[int(1)] = not (d_1455_v1_) or (d_1455_v1_)
            nw239_[int(2)] = True
            nw239_[int(3)] = d_1455_v1_
            nw239_[int(4)] = (d_1455_v1_ if d_1455_v1_ else (d_1458_v3_).fm3(default__.fm1(len(_dafny.Map({242: (d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))]})), d_1460_v5_, d_1455_v1_, d_1461_v6_, globalState), d_1462_v7_, d_1455_v1_, globalState))
            nw239_[int(5)] = not(d_1455_v1_)
            nw239_[int(6)] = d_1455_v1_
            nw239_[int(7)] = (d_1463_v8_)[default__.safeIndex((self).f10, len(d_1463_v8_))]
            nw239_[int(8)] = d_1455_v1_
            nw239_[int(9)] = ((self).f10) <= ((self).f10)
            nw239_[int(10)] = d_1455_v1_
            nw239_[int(11)] = d_1455_v1_
            nw239_[int(12)] = (d_1458_v3_).fm3(d_1464_v9_, d_1462_v7_, d_1455_v1_, globalState)
            nw239_[int(13)] = d_1455_v1_
            nw239_[int(14)] = (d_1455_v1_) == (True)
            nw239_[int(15)] = d_1455_v1_
            nw239_[int(16)] = d_1455_v1_
            nw239_[int(17)] = (d_1465_v10_) != (d_1465_v10_)
            nw239_[int(18)] = (d_1455_v1_ if d_1455_v1_ else d_1455_v1_)
            nw239_[int(19)] = d_1455_v1_
            nw239_[int(20)] = d_1455_v1_
            nw239_[int(21)] = d_1455_v1_
            nw239_[int(22)] = d_1455_v1_
            d_1466_v11_ = nw239_
            index273_ = default__.safeIndex(304, (d_1466_v11_).length(0))
            (d_1466_v11_)[index273_] = d_1455_v1_
            index274_ = default__.safeIndex(304, (d_1466_v11_).length(0))
            (d_1466_v11_)[index274_] = not((((self).f10) not in (d_1462_v7_)) and (d_1455_v1_))
            if ((515) - ((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))])) == (-38):
                d_1467_v12_: _dafny.Map
                d_1467_v12_ = _dafny.Map({d_1455_v1_: d_1460_v5_})
                d_1467_v12_ = (d_1467_v12_).set(d_1455_v1_, d_1460_v5_)
                (globalState).f4 = (d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))]
                d_1468_v13_: _dafny.MultiSet
                d_1468_v13_ = _dafny.MultiSet([d_1455_v1_, False])
                d_1469_v14_: _dafny.Set
                d_1469_v14_ = _dafny.Set({(d_1468_v13_).cardinality, default__.safeDivisionInt((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))], (self).f10)})
                d_1470_v15_: _dafny.Seq
                d_1470_v15_ = _dafny.SeqWithoutIsStrInference([d_1454_v0_])
                index275_ = default__.safeIndex(304, (d_1466_v11_).length(0))
                rhs260_ = d_1469_v14_
                rhs261_ = not((d_1470_v15_) == (d_1470_v15_))
                lhs210_ = d_1466_v11_
                lhs211_ = default__.safeIndex(304, (d_1466_v11_).length(0))
                d_1469_v14_ = rhs260_
                lhs210_[lhs211_] = rhs261_
                index276_ = default__.safeIndex(304, (d_1466_v11_).length(0))
                (d_1466_v11_)[index276_] = ((self).f10) < (d_1457_i0_)
                d_1471_v16_: _dafny.MultiSet
                d_1471_v16_ = _dafny.MultiSet([d_1463_v8_])
                r1 = ((d_1471_v16_).set(d_1463_v8_, default__.abs((0) - ((self).f10)))).ispropersubset((d_1471_v16_).intersection(_dafny.MultiSet([d_1463_v8_, d_1463_v8_])))
            elif True:
                d_1472_v17_: _dafny.Map
                d_1472_v17_ = _dafny.Map({d_1460_v5_: d_1465_v10_})
                d_1473_v18_: _dafny.Map
                d_1473_v18_ = _dafny.Map({((d_1472_v17_)[d_1460_v5_] if (d_1460_v5_) in (d_1472_v17_) else d_1465_v10_): d_1454_v0_})
                (globalState).f4 = (d_1465_v10_) in (d_1473_v18_)
                r0 = not(d_1455_v1_)
                d_1474_v19_: _dafny.Map
                d_1474_v19_ = _dafny.Map({(d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))]: _dafny.SeqWithoutIsStrInference([len(d_1464_v9_), (0) - ((self).f10)])})
                rhs262_ = (d_1455_v1_) and (d_1455_v1_)
                rhs263_ = (d_1474_v19_) | (d_1474_v19_)
                lhs212_ = globalState
                lhs212_.f4 = rhs262_
                d_1474_v19_ = rhs263_
                d_1475_v20_: _dafny.Set
                d_1475_v20_ = _dafny.Set({(d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))], (0) - (len((self).f11)), (d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))], (d_1457_i0_) - ((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))])})
                d_1476_v21_: _dafny.MultiSet
                d_1476_v21_ = _dafny.MultiSet([d_1455_v1_])
                d_1477_v22_: _dafny.Array
                def lambda74_(d_1478_v0_, d_1479_v1_, d_1480_v5_):
                    def lambda75_(d_1481_i1_):
                        return ((D10_DC24((d_1478_v0_)[default__.safeIndex(359, (d_1478_v0_).length(0))], d_1479_v1_, (self).f11)).cf33) + (_dafny.SeqWithoutIsStrInference([d_1480_v5_ for d_1482_i2_ in range(default__.abs(834))]))

                    return lambda75_

                init38_ = lambda74_(d_1454_v0_, d_1455_v1_, d_1460_v5_)
                nw240_ = _dafny.Array(None, 21)
                for i0_38_ in range(nw240_.length(0)):
                    nw240_[i0_38_] = init38_(i0_38_)
                d_1477_v22_ = nw240_
                index277_ = default__.safeIndex(193, (d_1477_v22_).length(0))
                (d_1477_v22_)[index277_] = (self).f11
                d_1483_v23_: D19
                d_1483_v23_ = D19_DC39(True, (self).f10)
                d_1484_v24_: D19
                d_1484_v24_ = D19_DC40(d_1483_v23_)
                d_1485_v25_: _dafny.Set
                d_1485_v25_ = _dafny.Set({d_1484_v24_, d_1484_v24_, D19_DC40(d_1483_v23_), d_1484_v24_})
                d_1486_v26_: _dafny.Map
                d_1486_v26_ = _dafny.Map({d_1485_v25_: (_dafny.MultiSet([(d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))], d_1455_v1_, (d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))], False])) | (_dafny.MultiSet(default__.fm31(d_1455_v1_, (d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))], d_1455_v1_, globalState)))})
                d_1487_v27_: _dafny.Map
                d_1487_v27_ = _dafny.Map({(d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))]: d_1463_v8_})
                index278_ = default__.safeIndex(193, (d_1477_v22_).length(0))
                rhs264_ = d_1475_v20_
                rhs265_ = ((d_1486_v26_)[d_1485_v25_] if (d_1485_v25_) in (d_1486_v26_) else (d_1476_v21_) | (d_1476_v21_))
                rhs266_ = len((((d_1487_v27_)[(d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))]] if ((d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))]) in (d_1487_v27_) else d_1463_v8_) if (d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))] else d_1463_v8_))
                rhs267_ = default__.safeModuloInt((d_1462_v7_).cardinality, d_1457_i0_)
                rhs268_ = (_dafny.SeqWithoutIsStrInference([d_1460_v5_ for d_1488_i3_ in range(default__.abs(-313))])) + ((self).f11)
                lhs213_ = globalState
                lhs214_ = globalState
                lhs215_ = d_1477_v22_
                lhs216_ = default__.safeIndex(193, (d_1477_v22_).length(0))
                d_1475_v20_ = rhs264_
                d_1476_v21_ = rhs265_
                lhs213_.f3 = rhs266_
                lhs214_.f8 = rhs267_
                lhs215_[lhs216_] = rhs268_
                d_1489_v28_: _dafny.Set
                d_1489_v28_ = _dafny.Set({(d_1466_v11_)[default__.safeIndex(304, (d_1466_v11_).length(0))], True})
                d_1489_v28_ = ((d_1489_v28_).intersection(d_1489_v28_)).intersection(_dafny.Set({True}))
        d_1490_v29_: _dafny.Array
        nw241_ = _dafny.Array(None, 17)
        nw241_[int(0)] = False
        nw241_[int(1)] = d_1455_v1_
        nw241_[int(2)] = d_1455_v1_
        nw241_[int(3)] = d_1455_v1_
        nw241_[int(4)] = d_1455_v1_
        nw241_[int(5)] = d_1455_v1_
        nw241_[int(6)] = True
        nw241_[int(7)] = d_1455_v1_
        nw241_[int(8)] = d_1455_v1_
        nw241_[int(9)] = d_1455_v1_
        nw241_[int(10)] = d_1455_v1_
        nw241_[int(11)] = d_1455_v1_
        nw241_[int(12)] = d_1455_v1_
        nw241_[int(13)] = d_1455_v1_
        nw241_[int(14)] = d_1455_v1_
        nw241_[int(15)] = d_1455_v1_
        nw241_[int(16)] = d_1455_v1_
        d_1490_v29_ = nw241_
        d_1491_v30_: _dafny.Set
        d_1491_v30_ = _dafny.Set({d_1490_v29_})
        d_1492_v31_: D2
        d_1492_v31_ = D2_DC5(d_1490_v29_)
        d_1493_v32_: _dafny.Set
        d_1493_v32_ = _dafny.Set({d_1455_v1_, d_1455_v1_})
        d_1494_v34_: _dafny.Set
        d_1494_v34_ = _dafny.Set({(self).f10, (self).f10})
        d_1495_v35_: _dafny.Seq
        d_1495_v35_ = _dafny.SeqWithoutIsStrInference([d_1455_v1_])
        d_1496_v36_: _dafny.Seq
        d_1496_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1495_v35_)])])
        d_1497_v37_: str
        d_1497_v37_ = _dafny.CodePoint('y')
        d_1498_v38_: _dafny.Map
        d_1498_v38_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvptjdmb"))).set(default__.safeIndex(len((d_1496_v36_)[default__.safeIndex((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))], len(d_1496_v36_))]), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvptjdmb")))), d_1497_v37_): False})
        d_1499_v39_: _dafny.Array
        nw242_ = _dafny.Array(None, 9)
        nw242_[int(0)] = ((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))]) <= ((0) - ((self).f10))
        nw242_[int(1)] = (d_1491_v30_).isdisjoint(_dafny.Set({d_1490_v29_, (d_1492_v31_).cf5}))
        nw242_[int(2)] = (d_1493_v32_).isdisjoint(_dafny.Set({d_1455_v1_}))
        nw242_[int(3)] = d_1455_v1_
        nw242_[int(4)] = d_1455_v1_
        nw242_[int(5)] = d_1455_v1_
        nw242_[int(6)] = d_1455_v1_
        def iife197_():
            coll91_ = _dafny.Set()
            compr_91_: int
            for compr_91_ in _dafny.IntegerRange(964, -933):
                d_1500_v33_: int = compr_91_
                if ((964) <= (d_1500_v33_)) and ((d_1500_v33_) < (-933)):
                    coll91_ = coll91_.union(_dafny.Set([(d_1500_v33_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([d_1455_v1_]))))]))
            return _dafny.Set(coll91_)
        nw242_[int(7)] = not((d_1494_v34_).ispropersubset(iife197_()
        ))
        nw242_[int(8)] = (not(False) if d_1455_v1_ else ((d_1498_v38_)[(self).f11] if ((self).f11) in (d_1498_v38_) else d_1455_v1_))
        d_1499_v39_ = nw242_
        index279_ = default__.safeIndex(443, (d_1490_v29_).length(0))
        (d_1490_v29_)[index279_] = True
        index280_ = default__.safeIndex(443, (d_1490_v29_).length(0))
        (d_1490_v29_)[index280_] = False
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1454_v0_).length(0)):
            d_1501_i4_: int = guard_loop_1_
            if (True) and (((0) <= (d_1501_i4_)) and ((d_1501_i4_) < ((d_1454_v0_).length(0)))):
                (d_1454_v0_)[(d_1501_i4_)] = (d_1501_i4_) * ((self).f10)
        hi9_ = (d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))]
        for d_1502_i5_ in range((self).f10, hi9_):
            r0 = d_1455_v1_
            d_1503_v40_: _dafny.Seq
            d_1503_v40_ = _dafny.SeqWithoutIsStrInference([d_1454_v0_])
            d_1454_v0_ = ((d_1503_v40_)[default__.safeIndex((self).f10, len(d_1503_v40_))] if d_1455_v1_ else d_1454_v0_)
            d_1504_v41_: C0
            nw243_ = C0()
            nw243_.ctor__()
            d_1504_v41_ = nw243_
            d_1505_v42_: D7
            d_1505_v42_ = D7_DC18()
            source27_ = d_1505_v42_
            if source27_.is_DC18:
                (globalState).f3 = default__.fm20((d_1490_v29_)[default__.safeIndex(443, (d_1490_v29_).length(0))], globalState)
                d_1506_v43_: _dafny.Seq
                d_1506_v43_ = _dafny.SeqWithoutIsStrInference([(self).f10, d_1502_i5_])
                d_1507_v44_: _dafny.Map
                d_1507_v44_ = _dafny.Map({d_1455_v1_: (d_1490_v29_)[default__.safeIndex(443, (d_1490_v29_).length(0))]})
                index281_ = default__.safeIndex(443, (d_1490_v29_).length(0))
                (d_1490_v29_)[index281_] = (default__.safeModuloInt((0) - (d_1502_i5_), len((d_1506_v43_).set(default__.safeIndex(len(d_1495_v35_), len(d_1506_v43_)), len(d_1507_v44_))))) <= ((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))])
                d_1508_v45_: D9
                d_1508_v45_ = D9_DC20(d_1493_v32_)
                d_1508_v45_ = d_1508_v45_
                d_1509_v46_: _dafny.Seq
                d_1509_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mn"))
                d_1509_v46_ = d_1509_v46_
            elif True:
                d_1510___mcc_h0_ = source27_.cf23
                d_1511_cf23_ = d_1510___mcc_h0_
                d_1512_v47_: _dafny.MultiSet
                d_1512_v47_ = _dafny.MultiSet([(d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))], 534])
                d_1513_v48_: _dafny.MultiSet
                d_1513_v48_ = _dafny.MultiSet([(self).fm3(_dafny.SeqWithoutIsStrInference([634 for d_1514_i6_ in range(default__.abs(267))]), d_1512_v47_, (d_1490_v29_)[default__.safeIndex(443, (d_1490_v29_).length(0))], globalState), d_1455_v1_, ((self).f10) <= (790), (d_1490_v29_)[default__.safeIndex(443, (d_1490_v29_).length(0))], d_1455_v1_])
                index282_ = default__.safeIndex(359, (d_1454_v0_).length(0))
                (d_1454_v0_)[index282_] = (d_1513_v48_).cardinality
                index283_ = default__.safeIndex(359, (d_1454_v0_).length(0))
                (d_1454_v0_)[index283_] = (0) - ((817 if (d_1490_v29_)[default__.safeIndex(443, (d_1490_v29_).length(0))] else default__.safeDivisionInt((self).f10, d_1502_i5_)))
                (globalState).f8 = d_1502_i5_
                (globalState).f3 = (0) - ((d_1454_v0_)[default__.safeIndex(359, (d_1454_v0_).length(0))])
        d_1497_v37_ = d_1497_v37_
        r0 = ((self).f11) != ((self).f11)
        r1 = d_1455_v1_
        return r0, r1

    def m7(self, p0, p1, p2, globalState):
        d_1515_i0_: int
        d_1515_i0_ = 0
        with _dafny.label("9"):
            while p0:
                with _dafny.c_label("9"):
                    if (d_1515_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1515_i0_ = (d_1515_i0_) + (1)
                    d_1516_v0_: _dafny.Seq
                    d_1516_v0_ = _dafny.SeqWithoutIsStrInference([p0, (p0 if p0 else p0), p0, not(p0)])
                    d_1517_v1_: _dafny.Array
                    def lambda76_(d_1518_i1_):
                        return ((self).f11).set(default__.safeIndex((self).f10, len((self).f11)), _dafny.CodePoint('l'))

                    init39_ = lambda76_
                    nw244_ = _dafny.Array(None, 5)
                    for i0_39_ in range(nw244_.length(0)):
                        nw244_[i0_39_] = init39_(i0_39_)
                    d_1517_v1_ = nw244_
                    index284_ = default__.safeIndex(464, (d_1517_v1_).length(0))
                    (d_1517_v1_)[index284_] = (default__.fm14((self).f10, globalState)) + ((self).f11)
                    d_1519_v2_: C2
                    nw245_ = C2()
                    nw245_.ctor__()
                    d_1519_v2_ = nw245_
                    d_1520_v3_: D16
                    d_1520_v3_ = D16_DC34(d_1519_v2_)
                    d_1521_v4_: str
                    d_1521_v4_ = _dafny.CodePoint('x')
                    d_1522_v5_: _dafny.MultiSet
                    d_1522_v5_ = _dafny.MultiSet([(self).f10])
                    d_1523_v6_: _dafny.Map
                    d_1523_v6_ = _dafny.Map({(default__.fm27(p0, globalState)).ispropersubset(d_1522_v5_): d_1521_v4_})
                    index285_ = default__.safeIndex(464, (d_1517_v1_).length(0))
                    rhs269_ = (d_1516_v0_) + (d_1516_v0_)
                    rhs270_ = ((self).f11) + ((self).f11)
                    rhs271_ = d_1520_v3_
                    rhs272_ = ((d_1523_v6_)[(p0 if p0 else p0)] if ((p0 if p0 else p0)) in (d_1523_v6_) else d_1521_v4_)
                    rhs273_ = p0
                    lhs217_ = d_1517_v1_
                    lhs218_ = default__.safeIndex(464, (d_1517_v1_).length(0))
                    lhs219_ = globalState
                    d_1516_v0_ = rhs269_
                    lhs217_[lhs218_] = rhs270_
                    d_1520_v3_ = rhs271_
                    d_1521_v4_ = rhs272_
                    lhs219_.f4 = rhs273_
                    (globalState).f4 = p0
                    d_1524_v7_: _dafny.Seq
                    d_1524_v7_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                    if (d_1519_v2_).fm3(d_1524_v7_, d_1522_v5_, p0, globalState):
                        d_1525_v8_: int
                        d_1526_v9_: int
                        d_1527_v10_: int
                        d_1528_v11_: bool
                        out22_: int
                        out23_: int
                        out24_: int
                        out25_: bool
                        out22_, out23_, out24_, out25_ = (d_1519_v2_).m12(globalState)
                        d_1525_v8_ = out22_
                        d_1526_v9_ = out23_
                        d_1527_v10_ = out24_
                        d_1528_v11_ = out25_
                        (globalState).f4 = d_1528_v11_
                        d_1529_v12_: _dafny.Set
                        d_1529_v12_ = _dafny.Set({(self).f10})
                        d_1529_v12_ = (d_1529_v12_).intersection(d_1529_v12_)
                        (globalState).f8 = ((0) - (d_1526_v9_)) * (((d_1524_v7_)[default__.safeIndex(d_1527_v10_, len(d_1524_v7_))]) + ((0) - ((d_1522_v5_).cardinality)))
                        index286_ = default__.safeIndex(464, (d_1517_v1_).length(0))
                        (d_1517_v1_)[index286_] = (self).f11
                    elif True:
                        d_1521_v4_ = d_1521_v4_
                        (globalState).f4 = p0
                        d_1530_v13_: _dafny.MultiSet
                        d_1530_v13_ = _dafny.MultiSet([not(p0), p0])
                        d_1531_v14_: D9
                        d_1531_v14_ = D9_DC22((d_1516_v0_).set(default__.safeIndex((d_1530_v13_).cardinality, len(d_1516_v0_)), p0))
                        d_1532_v15_: _dafny.Set
                        d_1532_v15_ = _dafny.Set({p0})
                        d_1533_v16_: _dafny.Map
                        d_1533_v16_ = _dafny.Map({d_1531_v14_: ((self).f10) * (len(d_1532_v15_))})
                        d_1534_v18_: _dafny.Map
                        def iife198_():
                            coll92_ = _dafny.Map()
                            compr_92_: int
                            for compr_92_ in _dafny.IntegerRange(-921, 719):
                                d_1535_v17_: int = compr_92_
                                if ((-921) <= (d_1535_v17_)) and ((d_1535_v17_) < (719)):
                                    coll92_[default__.safeModuloInt(d_1535_v17_, len(d_1532_v15_))] = d_1521_v4_
                            return _dafny.Map(coll92_)
                        d_1534_v18_ = _dafny.Map({114: len(iife198_()
                        )})
                        d_1536_v19_: _dafny.MultiSet
                        d_1536_v19_ = _dafny.MultiSet([d_1516_v0_])
                        d_1533_v16_ = (d_1533_v16_).set(d_1531_v14_, default__.fm26(D3_DC8(d_1534_v18_), (self).f10, d_1536_v19_, globalState))
                        d_1537_v20_: _dafny.Set
                        d_1537_v20_ = _dafny.Set({(self).f10, (self).f10, (self).f10})
                        (globalState).f8 = ((self).f10) - (len(d_1537_v20_))
                        (globalState).f4 = ((self).f10) <= (((d_1534_v18_)[(self).f10] if ((self).f10) in (d_1534_v18_) else (self).f10))
                    d_1538_v21_: C3
                    nw246_ = C3()
                    nw246_.ctor__()
                    d_1538_v21_ = nw246_
                    pass
            pass
        d_1539_v22_: _dafny.Array
        def lambda77_(d_1540_i3_):
            return (((self).f11).set(default__.safeIndex((self).f10, len((self).f11)), _dafny.CodePoint('d'))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjr")))

        init40_ = lambda77_
        nw247_ = _dafny.Array(None, 14)
        for i0_40_ in range(nw247_.length(0)):
            nw247_[i0_40_] = init40_(i0_40_)
        d_1539_v22_ = nw247_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1539_v22_).length(0)):
            d_1541_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_1541_i2_)) and ((d_1541_i2_) < ((d_1539_v22_).length(0)))):
                (d_1539_v22_)[(d_1541_i2_)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1542_i4_ in range(default__.abs(-239))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxtmbhm"))) + ((self).f11))
        d_1543_v23_: _dafny.Array
        def lambda78_(d_1544_p0_):
            def lambda79_(d_1545_i5_):
                return (_dafny.Set({len(_dafny.Map({(_dafny.MultiSet([d_1544_p0_, d_1544_p0_])).cardinality: _dafny.SeqWithoutIsStrInference([(self).f10 for d_1546_i6_ in range(default__.abs(542))])}))})).issubset(_dafny.Set({(self).f10}))

            return lambda79_

        init41_ = lambda78_(p0)
        nw248_ = _dafny.Array(None, 12)
        for i0_41_ in range(nw248_.length(0)):
            nw248_[i0_41_] = init41_(i0_41_)
        d_1543_v23_ = nw248_
        index287_ = default__.safeIndex(981, (d_1543_v23_).length(0))
        (d_1543_v23_)[index287_] = p0
        d_1547_v24_: _dafny.Seq
        d_1547_v24_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
        d_1548_v25_: _dafny.Seq
        d_1548_v25_ = _dafny.SeqWithoutIsStrInference([len(d_1547_v24_)])
        index288_ = default__.safeIndex(981, (d_1543_v23_).length(0))
        rhs274_ = (self).f10
        rhs275_ = (p0 if p0 else (self).fm3(d_1548_v25_, _dafny.MultiSet([(self).f10]), p0, globalState))
        rhs276_ = (p0) and (p0)
        rhs277_ = default__.safeModuloInt((self).f10, (self).f10)
        lhs220_ = globalState
        lhs221_ = d_1543_v23_
        lhs222_ = default__.safeIndex(981, (d_1543_v23_).length(0))
        lhs223_ = globalState
        lhs224_ = globalState
        lhs220_.f8 = rhs274_
        lhs221_[lhs222_] = rhs275_
        lhs223_.f4 = rhs276_
        lhs224_.f8 = rhs277_
        d_1549_v26_: _dafny.Array
        nw249_ = _dafny.Array(int(0), 21)
        d_1549_v26_ = nw249_
        d_1550_v27_: _dafny.Array
        nw250_ = _dafny.Array(None, 18)
        nw250_[int(0)] = d_1549_v26_
        nw250_[int(1)] = d_1549_v26_
        nw250_[int(2)] = d_1549_v26_
        nw250_[int(3)] = d_1549_v26_
        nw250_[int(4)] = d_1549_v26_
        nw250_[int(5)] = d_1549_v26_
        nw250_[int(6)] = d_1549_v26_
        nw250_[int(7)] = d_1549_v26_
        nw250_[int(8)] = d_1549_v26_
        nw250_[int(9)] = d_1549_v26_
        nw250_[int(10)] = d_1549_v26_
        nw250_[int(11)] = d_1549_v26_
        nw250_[int(12)] = d_1549_v26_
        nw250_[int(13)] = d_1549_v26_
        nw250_[int(14)] = d_1549_v26_
        nw250_[int(15)] = d_1549_v26_
        nw250_[int(16)] = d_1549_v26_
        nw250_[int(17)] = d_1549_v26_
        d_1550_v27_ = nw250_
        d_1551_v28_: _dafny.Map
        d_1551_v28_ = _dafny.Map({d_1550_v27_: 179})
        hi10_ = (0) - (((self).f10) + ((self).f10))
        for d_1552_i7_ in range(((d_1551_v28_)[d_1550_v27_] if (d_1550_v27_) in (d_1551_v28_) else default__.fm20((d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))], globalState)), hi10_):
            d_1553_v29_: _dafny.MultiSet
            d_1553_v29_ = _dafny.MultiSet([p0])
            d_1554_v30_: _dafny.Map
            d_1554_v30_ = _dafny.Map({(d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]: (d_1553_v29_).issubset(d_1553_v29_)})
            d_1555_v31_: _dafny.MultiSet
            d_1555_v31_ = _dafny.MultiSet([(self).f10])
            index289_ = default__.safeIndex(981, (d_1543_v23_).length(0))
            rhs278_ = ((d_1552_i7_) + ((self).f10)) - (default__.safeDivisionInt(114, (d_1555_v31_).cardinality))
            rhs279_ = (self).f10
            rhs280_ = ((d_1554_v30_) | (d_1554_v30_)) | (d_1554_v30_)
            rhs281_ = (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]
            rhs282_ = (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]
            lhs225_ = globalState
            lhs226_ = globalState
            lhs227_ = globalState
            lhs228_ = d_1543_v23_
            lhs229_ = default__.safeIndex(981, (d_1543_v23_).length(0))
            lhs225_.f3 = rhs278_
            lhs226_.f8 = rhs279_
            d_1554_v30_ = rhs280_
            lhs227_.f4 = rhs281_
            lhs228_[lhs229_] = rhs282_
            d_1556_v32_: C5
            nw251_ = C5()
            nw251_.ctor__()
            d_1556_v32_ = nw251_
            d_1557_v33_: _dafny.Map
            d_1557_v33_ = _dafny.Map({d_1554_v30_: d_1552_i7_})
            d_1558_v34_: _dafny.Map
            d_1558_v34_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])): (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]})
            d_1559_v35_: _dafny.Map
            d_1559_v35_ = _dafny.Map({d_1555_v31_: not((_dafny.Map({(self).f10: p0})) == (d_1558_v34_))})
            d_1560_v36_: C0
            nw252_ = C0()
            nw252_.ctor__()
            d_1560_v36_ = nw252_
            d_1561_v37_: _dafny.Set
            d_1561_v37_ = _dafny.Set({d_1560_v36_, d_1560_v36_})
            rhs283_ = (d_1557_v33_) | ((d_1557_v33_) | (_dafny.Map({d_1554_v30_: len(_dafny.Map({(self).f10: len(d_1561_v37_)}))})))
            rhs284_ = _dafny.Map({(_dafny.MultiSet([d_1552_i7_])).intersection(d_1555_v31_): (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]})
            rhs285_ = d_1549_v26_
            rhs286_ = (d_1547_v24_) + (d_1548_v25_)
            lhs230_ = globalState
            d_1557_v33_ = rhs283_
            d_1559_v35_ = rhs284_
            d_1549_v26_ = rhs285_
            lhs230_.f0 = rhs286_
            d_1562_v38_: _dafny.Seq
            d_1562_v38_ = _dafny.SeqWithoutIsStrInference([d_1549_v26_])
            d_1563_v39_: _dafny.Map
            d_1563_v39_ = _dafny.Map({(d_1562_v38_) + (d_1562_v38_): (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]})
            d_1564_v40_: str
            d_1564_v40_ = _dafny.CodePoint('l')
            (globalState).f4 = ((d_1563_v39_)[d_1562_v38_] if (d_1562_v38_) in (d_1563_v39_) else ((self).f10) > (len(((self).f11).set(default__.safeIndex(d_1552_i7_, len((self).f11)), d_1564_v40_))))
        if True:
            d_1565_v41_: _dafny.Seq
            d_1565_v41_ = _dafny.SeqWithoutIsStrInference([(d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]])
            d_1566_v42_: D4
            d_1566_v42_ = D4_DC10((d_1565_v41_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f10), (self).f10])), len(d_1565_v41_)), p0))
            d_1567_v43_: _dafny.Map
            d_1567_v43_ = _dafny.Map({((self).f10) * ((self).f10): d_1566_v42_})
            d_1567_v43_ = (d_1567_v43_).set((self).f10, d_1566_v42_)
            d_1568_v44_: _dafny.Seq
            d_1568_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjqmhkj"))
            d_1568_v44_ = d_1568_v44_
            d_1569_v45_: _dafny.MultiSet
            d_1569_v45_ = _dafny.MultiSet([(self).f10, (0) - (len(d_1548_v25_)), (self).f10])
            index290_ = default__.safeIndex(981, (d_1543_v23_).length(0))
            (d_1543_v23_)[index290_] = ((self).f10) == ((d_1569_v45_).cardinality)
            (globalState).f4 = p0
            (globalState).f3 = (self).f10
        elif True:
            d_1570_v46_: str
            d_1570_v46_ = _dafny.CodePoint('i')
            d_1570_v46_ = (_dafny.CodePoint('q') if (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))] else d_1570_v46_)
            index291_ = default__.safeIndex(682, (d_1539_v22_).length(0))
            (d_1539_v22_)[index291_] = (self).f11
            index292_ = default__.safeIndex(682, (d_1539_v22_).length(0))
            (d_1539_v22_)[index292_] = (self).f11
            d_1571_v47_: _dafny.Set
            d_1571_v47_ = _dafny.Set({(self).f10})
            d_1572_v48_: D12
            d_1572_v48_ = D12_DC27((self).f10, len(d_1571_v47_), (self).f10, (self).f10)
            d_1573_v49_: _dafny.Map
            d_1573_v49_ = _dafny.Map({(_dafny.MultiSet([(self).f10])).cardinality: (d_1572_v48_).cf39})
            d_1574_v50_: D3
            d_1574_v50_ = D3_DC8(d_1573_v49_)
            (globalState).f8 = default__.fm26(d_1574_v50_, (self).f10, default__.fm44((self).f10, globalState), globalState)
            (globalState).f4 = ((d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))] if p0 else (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))])
            d_1575_v51_: _dafny.Seq
            d_1575_v51_ = _dafny.SeqWithoutIsStrInference([(d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))], p0])
            d_1576_v52_: _dafny.Seq
            d_1576_v52_ = _dafny.SeqWithoutIsStrInference([d_1575_v51_, d_1575_v51_, d_1575_v51_])
            d_1577_v53_: _dafny.Set
            d_1577_v53_ = _dafny.Set({(d_1576_v52_)[default__.safeIndex(len(_dafny.Set({False, p0, p0, (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))], (d_1543_v23_)[default__.safeIndex(981, (d_1543_v23_).length(0))]})), len(d_1576_v52_))]})
            (globalState).f3 = (len(d_1577_v53_)) * (len((self).f11))
        d_1578_v54_: str
        d_1578_v54_ = _dafny.CodePoint('k')
        d_1579_v55_: D2
        d_1579_v55_ = D2_DC7(d_1578_v54_, True, (self).f10)
        d_1580_v56_: _dafny.Set
        d_1580_v56_ = _dafny.Set({(self).f10, (d_1579_v55_).cf11})
        (globalState).f8 = len(d_1580_v56_)

    def m8(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_1581_v0_: bool
        d_1581_v0_ = True
        d_1582_v1_: T0
        nw253_ = C1()
        nw253_.ctor__()
        d_1582_v1_ = nw253_
        d_1583_v2_: _dafny.Seq
        d_1583_v2_ = _dafny.SeqWithoutIsStrInference([d_1582_v1_])
        d_1584_v3_: _dafny.Map
        d_1584_v3_ = _dafny.Map({d_1581_v0_: d_1583_v2_})
        (globalState).f8 = default__.safeModuloInt(len((self).f11), len(((d_1584_v3_)[d_1581_v0_] if (d_1581_v0_) in (d_1584_v3_) else (d_1583_v2_))))
        d_1585_v4_: _dafny.Seq
        d_1585_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1586_v5_: _dafny.Array
        nw254_ = _dafny.Array(_dafny.CodePoint('D'), 14)
        d_1586_v5_ = nw254_
        d_1587_v6_: _dafny.Seq
        d_1587_v6_ = _dafny.SeqWithoutIsStrInference([d_1586_v5_, d_1586_v5_])
        d_1588_v7_: D3
        d_1588_v7_ = D3_DC8(_dafny.Map({len(d_1587_v6_): len((self).f11)}))
        d_1589_v8_: _dafny.MultiSet
        d_1589_v8_ = _dafny.MultiSet([default__.fm31(d_1581_v0_, d_1581_v0_, d_1581_v0_, globalState)])
        (globalState).f3 = (0) - ((d_1585_v4_)[default__.safeIndex(default__.fm26(d_1588_v7_, (self).f10, d_1589_v8_, globalState), len(d_1585_v4_))])
        d_1590_v9_: C3
        nw255_ = C3()
        nw255_.ctor__()
        d_1590_v9_ = nw255_
        d_1591_i0_: int
        d_1591_i0_ = 0
        with _dafny.label("10"):
            while ((self).f10) == ((0) - ((389) * ((self).f10))):
                with _dafny.c_label("10"):
                    if (d_1591_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1591_i0_ = (d_1591_i0_) + (1)
                    d_1592_v10_: _dafny.Array
                    def lambda80_(d_1593_p0_):
                        def lambda81_(d_1594_i1_):
                            return _dafny.Map({not(False): d_1593_p0_})

                        return lambda81_

                    init42_ = lambda80_(p0)
                    nw256_ = _dafny.Array(None, 29)
                    for i0_42_ in range(nw256_.length(0)):
                        nw256_[i0_42_] = init42_(i0_42_)
                    d_1592_v10_ = nw256_
                    d_1595_v11_: C4
                    nw257_ = C4()
                    nw257_.ctor__(d_1592_v10_)
                    d_1595_v11_ = nw257_
                    d_1596_v12_: _dafny.Map
                    d_1596_v12_ = _dafny.Map({d_1581_v0_: (self).f10})
                    (globalState).f3 = default__.safeDivisionInt(p0, ((d_1596_v12_)[d_1581_v0_] if (d_1581_v0_) in (d_1596_v12_) else (self).f10))
                    d_1597_v13_: str
                    d_1597_v13_ = _dafny.CodePoint('f')
                    d_1598_v14_: _dafny.Map
                    d_1598_v14_ = _dafny.Map({(d_1595_v11_).fm13(d_1581_v0_, (self).f11, d_1581_v0_, globalState): d_1597_v13_})
                    d_1598_v14_ = (d_1598_v14_).set(d_1581_v0_, d_1597_v13_)
                    d_1599_v15_: _dafny.MultiSet
                    d_1599_v15_ = _dafny.MultiSet([(self).f10])
                    d_1600_v16_: _dafny.Seq
                    d_1600_v16_ = _dafny.SeqWithoutIsStrInference([not((d_1595_v11_).fm3(d_1585_v4_, d_1599_v15_, d_1581_v0_, globalState))])
                    d_1601_v17_: bool
                    d_1601_v17_ = d_1581_v0_
                    d_1602_v18_: D0
                    d_1602_v18_ = D0_DC1((d_1590_v9_).fm15(len(d_1600_v16_), globalState), (0) - (((d_1595_v11_).fm12(d_1581_v0_, p0, (self).f11, d_1601_v17_, globalState)) - (p0)))
                    d_1603_v19_: _dafny.Map
                    d_1603_v19_ = _dafny.Map({(self).f10: p0})
                    d_1604_v20_: _dafny.Set
                    d_1604_v20_ = _dafny.Set({p0})
                    d_1605_v21_: _dafny.MultiSet
                    d_1605_v21_ = _dafny.MultiSet([_dafny.Set({((d_1599_v15_)[p0] if (p0) in (d_1599_v15_) else p0), (self).f10, len(d_1603_v19_)}), d_1604_v20_, d_1604_v20_])
                    d_1606_v22_: _dafny.Map
                    d_1606_v22_ = _dafny.Map({D2_DC7(d_1597_v13_, d_1581_v0_, (self).f10): (d_1605_v21_).cardinality})
                    d_1607_v23_: _dafny.Set
                    d_1607_v23_ = _dafny.Set({((self).f10) - ((self).f10), len(d_1606_v22_)})
                    pat_let_tv63_ = d_1599_v15_
                    pat_let_tv64_ = d_1599_v15_
                    pat_let_tv65_ = p0
                    def iife199_(_pat_let53_0):
                        def iife200_(d_1608_dt__update__tmp_h0_):
                            def iife201_(_pat_let54_0):
                                def iife202_(d_1609_dt__update_hcf1_h0_):
                                    return D0_DC1(d_1609_dt__update_hcf1_h0_, (d_1608_dt__update__tmp_h0_).cf2)
                                return iife202_(_pat_let54_0)
                            return iife201_((0) - (((pat_let_tv63_)[(self).f10] if ((self).f10) in (pat_let_tv64_) else pat_let_tv65_)))
                        return iife200_(_pat_let53_0)
                    rhs287_ = d_1581_v0_
                    rhs288_ = not((d_1597_v13_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcumpsjw"))))
                    rhs289_ = iife199_(d_1602_v18_)
                    rhs290_ = d_1607_v23_
                    lhs231_ = globalState
                    d_1581_v0_ = rhs287_
                    lhs231_.f4 = rhs288_
                    d_1602_v18_ = rhs289_
                    d_1604_v20_ = rhs290_
                    pass
            pass
        d_1610_v24_: _dafny.Set
        d_1610_v24_ = _dafny.Set({d_1581_v0_, d_1581_v0_})
        d_1611_v25_: _dafny.Seq
        d_1611_v25_ = _dafny.SeqWithoutIsStrInference([not((d_1610_v24_).ispropersubset(d_1610_v24_))])
        d_1612_v26_: _dafny.Map
        d_1612_v26_ = _dafny.Map({d_1581_v0_: d_1581_v0_})
        d_1613_i2_: int
        d_1613_i2_ = 0
        with _dafny.label("11"):
            while (d_1611_v25_)[default__.safeIndex(len(d_1612_v26_), len(d_1611_v25_))]:
                with _dafny.c_label("11"):
                    if (d_1613_i2_) >= (100):
                        raise _dafny.Break("11")
                    d_1613_i2_ = (d_1613_i2_) + (1)
                    d_1614_v27_: _dafny.Array
                    def lambda82_(d_1615_v0_):
                        def lambda83_(d_1616_i3_):
                            return d_1615_v0_

                        return lambda83_

                    init43_ = lambda82_(d_1581_v0_)
                    nw258_ = _dafny.Array(None, 18)
                    for i0_43_ in range(nw258_.length(0)):
                        nw258_[i0_43_] = init43_(i0_43_)
                    d_1614_v27_ = nw258_
                    index293_ = default__.safeIndex(450, (d_1614_v27_).length(0))
                    (d_1614_v27_)[index293_] = d_1581_v0_
                    d_1617_v28_: _dafny.MultiSet
                    d_1617_v28_ = _dafny.MultiSet([(self).f10])
                    index294_ = default__.safeIndex(450, (d_1614_v27_).length(0))
                    (d_1614_v27_)[index294_] = ((d_1617_v28_).intersection(d_1617_v28_)).isdisjoint(d_1617_v28_)
                    (globalState).f8 = default__.safeModuloInt(463, len((self).f11))
                    d_1618_v29_: _dafny.Seq
                    d_1618_v29_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
                    source28_ = d_1618_v29_
                    d_1619___mcc_h0_ = source28_
                    d_1620_cf24_ = d_1619___mcc_h0_
                    d_1621_v30_: _dafny.Array
                    nw259_ = _dafny.Array(D2.default()(), 9)
                    d_1621_v30_ = nw259_
                    d_1622_v31_: _dafny.MultiSet
                    d_1622_v31_ = _dafny.MultiSet([d_1621_v30_])
                    d_1623_v32_: _dafny.MultiSet
                    d_1623_v32_ = d_1622_v31_
                    d_1623_v32_ = d_1623_v32_
                    r0 = d_1617_v28_
                    d_1624_v33_: _dafny.Seq
                    d_1624_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoxvu"))
                    d_1625_v34_: str
                    d_1625_v34_ = _dafny.CodePoint('e')
                    d_1626_v35_: D10
                    d_1626_v35_ = D10_DC24(p0, (d_1614_v27_)[default__.safeIndex(450, (d_1614_v27_).length(0))], _dafny.SeqWithoutIsStrInference([d_1625_v34_]))
                    d_1624_v33_ = (d_1626_v35_).cf33
                    (globalState).f8 = (0) - (((self).f10) * (default__.safeModuloInt(len((self).f11), len(d_1624_v33_))))
                    d_1627_v36_: _dafny.Array
                    nw260_ = _dafny.Array(_dafny.Array(None, 0), 18)
                    d_1627_v36_ = nw260_
                    d_1628_v37_: _dafny.Array
                    nw261_ = _dafny.Array(int(0), 25)
                    d_1628_v37_ = nw261_
                    index295_ = default__.safeIndex(506, (d_1627_v36_).length(0))
                    (d_1627_v36_)[index295_] = d_1628_v37_
                    index296_ = default__.safeIndex(506, (d_1627_v36_).length(0))
                    (d_1627_v36_)[index296_] = d_1628_v37_
                    pass
            pass
        d_1629_v38_: _dafny.Seq
        d_1629_v38_ = d_1583_v2_
        source29_ = d_1629_v38_
        d_1630___mcc_h1_ = source29_
        d_1631_cf61_ = d_1630___mcc_h1_
        if d_1581_v0_:
            d_1632_v39_: str
            d_1632_v39_ = _dafny.CodePoint('n')
            d_1632_v39_ = d_1632_v39_
            (globalState).f3 = p0
            d_1585_v4_ = d_1585_v4_
            (globalState).f8 = ((0) - ((self).f10)) + ((self).f10)
            d_1633_v40_: _dafny.Seq
            d_1633_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nujca"))
            rhs291_ = default__.safeModuloInt(-550, len(((self).f11).set(default__.safeIndex((0) - (len(d_1633_v40_)), len((self).f11)), d_1632_v39_)))
            rhs292_ = ((self).f11).set(default__.safeIndex(default__.fm26(d_1588_v7_, p0, default__.fm44(p0, globalState), globalState), len((self).f11)), d_1632_v39_)
            lhs232_ = globalState
            lhs232_.f3 = rhs291_
            d_1633_v40_ = rhs292_
        elif True:
            d_1634_v41_: str
            d_1634_v41_ = _dafny.CodePoint('a')
            d_1635_v42_: _dafny.Map
            d_1635_v42_ = _dafny.Map({d_1634_v41_: d_1581_v0_})
            (globalState).f0 = (_dafny.SeqWithoutIsStrInference([len(d_1610_v24_) for d_1636_i4_ in range(default__.abs(958))])) + (_dafny.SeqWithoutIsStrInference([(self).f10, len(d_1635_v42_), len(d_1611_v25_)]))
            d_1637_v43_: D5
            d_1637_v43_ = D5_DC13((d_1612_v26_) | (d_1612_v26_))
            d_1637_v43_ = d_1637_v43_
            d_1638_v44_: _dafny.Map
            d_1638_v44_ = _dafny.Map({len((self).f11): d_1581_v0_})
            d_1639_v45_: _dafny.Set
            d_1639_v45_ = _dafny.Set({p0})
            d_1640_v46_: _dafny.MultiSet
            d_1640_v46_ = _dafny.MultiSet([_dafny.CodePoint('j')])
            d_1638_v44_ = (d_1638_v44_).set(len((d_1639_v45_).intersection(d_1639_v45_)), (d_1640_v46_) == (_dafny.MultiSet([d_1634_v41_])))
            d_1641_v47_: _dafny.Map
            d_1641_v47_ = _dafny.Map({p0: (self).f11})
            d_1642_v48_: _dafny.Map
            d_1642_v48_ = _dafny.Map({((d_1641_v47_)[(self).f10] if ((self).f10) in (d_1641_v47_) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))).set(default__.safeIndex(len(_dafny.Set({not(True), True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), d_1634_v41_)): (0) - (((self).f10) + ((self).f10))})
            d_1642_v48_ = d_1642_v48_
            d_1643_v49_: _dafny.MultiSet
            d_1643_v49_ = _dafny.MultiSet([(self).f10, (self).f10, p0])
            d_1644_v50_: _dafny.Array
            nw262_ = _dafny.Array(None, 2)
            nw262_[int(0)] = d_1581_v0_
            nw262_[int(1)] = (d_1582_v1_).fm3((d_1585_v4_).set(default__.safeIndex(p0, len(d_1585_v4_)), (self).f10), d_1643_v49_, d_1581_v0_, globalState)
            d_1644_v50_ = nw262_
            index297_ = default__.safeIndex(879, (d_1644_v50_).length(0))
            (d_1644_v50_)[index297_] = (d_1590_v9_).fm16(d_1581_v0_, globalState)
            index298_ = default__.safeIndex(879, (d_1644_v50_).length(0))
            (d_1644_v50_)[index298_] = (((self).f11) < ((self).f11)) not in (d_1611_v25_)
        d_1581_v0_ = not(d_1581_v0_)
        d_1645_v51_: _dafny.Array
        def lambda84_(d_1646_v0_):
            def lambda85_(d_1647_i5_):
                return d_1646_v0_

            return lambda85_

        init44_ = lambda84_(d_1581_v0_)
        nw263_ = _dafny.Array(None, 17)
        for i0_44_ in range(nw263_.length(0)):
            nw263_[i0_44_] = init44_(i0_44_)
        d_1645_v51_ = nw263_
        index299_ = default__.safeIndex(708, (d_1645_v51_).length(0))
        (d_1645_v51_)[index299_] = d_1581_v0_
        index300_ = default__.safeIndex(708, (d_1645_v51_).length(0))
        (d_1645_v51_)[index300_] = d_1581_v0_
        d_1648_v52_: C1
        nw264_ = C1()
        nw264_.ctor__()
        d_1648_v52_ = nw264_
        d_1649_v53_: _dafny.Array
        nw265_ = _dafny.Array(int(0), 20)
        d_1649_v53_ = nw265_
        d_1650_v54_: _dafny.Seq
        d_1650_v54_ = _dafny.SeqWithoutIsStrInference([d_1649_v53_, d_1649_v53_])
        d_1651_v55_: _dafny.MultiSet
        d_1651_v55_ = _dafny.MultiSet([default__.safeModuloInt((d_1590_v9_).fm15(p0, globalState), len(d_1650_v54_))])
        r0 = d_1651_v55_
        return r0

    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11

class C7(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm2(self, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1652_i0_ in range(default__.abs(786))])): 529})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctfhl"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))}))) | ((D3_DC8(_dafny.Map({-955: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awvtjnrl"))))}))).cf12)

    def fm3(self, p0, p1, p2, globalState):
        def iife203_():
            coll93_ = _dafny.Set()
            compr_93_: int
            for compr_93_ in _dafny.IntegerRange(385, 693):
                d_1653_v0_: int = compr_93_
                if ((385) <= (d_1653_v0_)) and ((d_1653_v0_) < (693)):
                    coll93_ = coll93_.union(_dafny.Set([(d_1653_v0_) * (-762)]))
            return _dafny.Set(coll93_)
        def iife204_():
            coll94_ = _dafny.Map()
            compr_94_: int
            for compr_94_ in (_dafny.Map({len(_dafny.Map({646: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsmppmgyx")))})): 95})).keys.Elements:
                d_1656_v1_: int = compr_94_
                if (d_1656_v1_) in (_dafny.Map({len(_dafny.Map({646: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsmppmgyx")))})): 95})):
                    coll94_[(d_1656_v1_) - (-781)] = 797
            return _dafny.Map(coll94_)
        return ((_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no"))): False})), (D0_DC0(len(iife203_()
))).cf0, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1654_i1_ in range(default__.abs(173))]))) for d_1655_i0_ in range(default__.abs(501))]))])).cardinality, 300, 524])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: 795}))])))).isdisjoint(_dafny.MultiSet([len(_dafny.Map({470: True})), (0) - (len(_dafny.Map({not(False): False}))), -52, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({813: D3_DC8(iife204_()
)}))]))]))

    def fm10(self, globalState):
        return len(((_dafny.SeqWithoutIsStrInference([-113 for d_1657_i0_ in range(default__.abs(15))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(not(True))]))]))) + (_dafny.SeqWithoutIsStrInference([993])))

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_1658_v0_: int
        d_1658_v0_ = 905
        d_1659_v1_: D2
        d_1659_v1_ = D2_DC6(d_1658_v0_, False, True)
        d_1660_v2_: bool
        d_1660_v2_ = False
        d_1661_v3_: _dafny.MultiSet
        d_1661_v3_ = _dafny.MultiSet([d_1659_v1_, d_1659_v1_, d_1659_v1_, d_1659_v1_, D2_DC6(d_1658_v0_, d_1660_v2_, d_1660_v2_)])
        d_1661_v3_ = d_1661_v3_
        d_1660_v2_ = d_1660_v2_
        (globalState).f8 = d_1658_v0_
        d_1662_v4_: bool
        d_1662_v4_ = d_1660_v2_
        def lambda86_(source30_):
            d_1663___mcc_h6_ = source30_
            d_1664_cf4_ = d_1663___mcc_h6_
            return d_1664_cf4_

        if lambda86_(d_1662_v4_):
            if (True) and (d_1660_v2_):
                (globalState).f3 = (0) - (d_1658_v0_)
                d_1665_v5_: _dafny.Map
                d_1665_v5_ = _dafny.Map({not(d_1660_v2_): d_1658_v0_})
                d_1666_v6_: _dafny.Seq
                d_1666_v6_ = _dafny.SeqWithoutIsStrInference([d_1658_v0_, d_1658_v0_, d_1658_v0_])
                d_1667_v7_: _dafny.MultiSet
                d_1667_v7_ = _dafny.MultiSet([d_1658_v0_])
                d_1668_v8_: _dafny.Map
                d_1668_v8_ = _dafny.Map({d_1665_v5_: (self).fm3(d_1666_v6_, d_1667_v7_, d_1660_v2_, globalState)})
                d_1669_v9_: _dafny.Map
                d_1669_v9_ = _dafny.Map({d_1668_v8_: d_1658_v0_})
                d_1658_v0_ = (len(default__.fm11(d_1660_v2_, d_1660_v2_, globalState)) if (not(not(d_1660_v2_)) if d_1660_v2_ else not(d_1660_v2_)) else (0) - (((d_1669_v9_)[d_1668_v8_] if (d_1668_v8_) in (d_1669_v9_) else -971)))
                d_1670_v10_: _dafny.Array
                nw266_ = _dafny.Array(None, 5)
                nw266_[int(0)] = d_1665_v5_
                nw266_[int(1)] = d_1665_v5_
                nw266_[int(2)] = _dafny.Map({d_1660_v2_: 518})
                nw266_[int(3)] = d_1665_v5_
                nw266_[int(4)] = _dafny.Map({d_1660_v2_: d_1658_v0_})
                d_1670_v10_ = nw266_
                d_1671_v11_: C4
                nw267_ = C4()
                nw267_.ctor__(d_1670_v10_)
                d_1671_v11_ = nw267_
                d_1672_v12_: T0
                nw268_ = C4()
                nw268_.ctor__(d_1670_v10_)
                d_1672_v12_ = nw268_
                d_1673_v13_: _dafny.Seq
                d_1673_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nph"))
                d_1674_v14_: _dafny.Map
                d_1674_v14_ = _dafny.Map({d_1660_v2_: (d_1671_v11_).fm3(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1673_v13_, d_1673_v13_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmigiwsi")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1675_i1_ in range(default__.abs(211))]), d_1673_v13_])) for d_1676_i0_ in range(default__.abs(518))]), d_1667_v7_, d_1660_v2_, globalState)})
                rhs293_ = d_1672_v12_
                rhs294_ = not(d_1660_v2_)
                rhs295_ = (0) - (((d_1658_v0_) + (len(d_1673_v13_))) - (len((d_1674_v14_) | (default__.fm21(d_1660_v2_, globalState)))))
                rhs296_ = d_1658_v0_
                lhs233_ = globalState
                lhs234_ = globalState
                lhs235_ = globalState
                d_1672_v12_ = rhs293_
                lhs233_.f4 = rhs294_
                lhs234_.f8 = rhs295_
                lhs235_.f8 = rhs296_
                d_1677_v16_: _dafny.Array
                def lambda87_(d_1678_v2_, d_1679_v0_):
                    def lambda88_(d_1680_i2_):
                        def iife205_():
                            coll95_ = _dafny.Set()
                            compr_95_: int
                            for compr_95_ in (_dafny.Map({d_1679_v0_: (0) - (d_1679_v0_)})).keys.Elements:
                                d_1681_v15_: int = compr_95_
                                if (d_1681_v15_) in (_dafny.Map({d_1679_v0_: (0) - (d_1679_v0_)})):
                                    coll95_ = coll95_.union(_dafny.Set([default__.safeModuloInt(d_1681_v15_, 612)]))
                            return _dafny.Set(coll95_)
                        return _dafny.Set({iife205_()
                        , _dafny.Set({d_1679_v0_}), _dafny.Set({d_1679_v0_, len(_dafny.SeqWithoutIsStrInference([d_1678_v2_, d_1678_v2_]))})})

                    return lambda88_

                init45_ = lambda87_(d_1660_v2_, d_1658_v0_)
                nw269_ = _dafny.Array(None, 7)
                for i0_45_ in range(nw269_.length(0)):
                    nw269_[i0_45_] = init45_(i0_45_)
                d_1677_v16_ = nw269_
                d_1682_v17_: _dafny.Set
                d_1682_v17_ = _dafny.Set({len(_dafny.Set({True})), (0) - (d_1658_v0_), len(d_1673_v13_)})
                d_1683_v18_: _dafny.Set
                d_1683_v18_ = _dafny.Set({d_1682_v17_})
                index301_ = default__.safeIndex(883, (d_1677_v16_).length(0))
                (d_1677_v16_)[index301_] = d_1683_v18_
                d_1684_v19_: _dafny.Array
                nw270_ = _dafny.Array(False, 7)
                d_1684_v19_ = nw270_
                index302_ = default__.safeIndex(858, (d_1684_v19_).length(0))
                (d_1684_v19_)[index302_] = (d_1672_v12_).fm3(d_1666_v6_, _dafny.MultiSet([d_1658_v0_, d_1658_v0_]), d_1660_v2_, globalState)
                index303_ = default__.safeIndex(883, (d_1677_v16_).length(0))
                index304_ = default__.safeIndex(858, (d_1684_v19_).length(0))
                rhs297_ = ((d_1674_v14_)[d_1660_v2_] if (d_1660_v2_) in (d_1674_v14_) else not(d_1660_v2_))
                rhs298_ = d_1672_v12_
                rhs299_ = d_1683_v18_
                rhs300_ = d_1660_v2_
                lhs236_ = d_1677_v16_
                lhs237_ = default__.safeIndex(883, (d_1677_v16_).length(0))
                lhs238_ = d_1684_v19_
                lhs239_ = default__.safeIndex(858, (d_1684_v19_).length(0))
                r0 = rhs297_
                d_1672_v12_ = rhs298_
                lhs236_[lhs237_] = rhs299_
                lhs238_[lhs239_] = rhs300_
            elif True:
                d_1685_v20_: _dafny.Set
                d_1685_v20_ = _dafny.Set({d_1658_v0_})
                d_1686_v21_: _dafny.MultiSet
                d_1686_v21_ = _dafny.MultiSet([(_dafny.MultiSet([d_1660_v2_, d_1660_v2_])).cardinality])
                d_1687_v22_: _dafny.Seq
                d_1687_v22_ = _dafny.SeqWithoutIsStrInference([(d_1685_v20_).issubset(d_1685_v20_), not(not((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1658_v0_, d_1658_v0_]))).issubset(d_1686_v21_))), not(d_1660_v2_)])
                d_1687_v22_ = d_1687_v22_
                d_1688_v23_: C1
                nw271_ = C1()
                nw271_.ctor__()
                d_1688_v23_ = nw271_
                d_1689_v24_: _dafny.Map
                d_1689_v24_ = _dafny.Map({d_1660_v2_: d_1660_v2_})
                d_1690_v25_: _dafny.Map
                d_1690_v25_ = _dafny.Map({d_1688_v23_: ((d_1689_v24_)[d_1660_v2_] if (d_1660_v2_) in (d_1689_v24_) else d_1660_v2_)})
                d_1691_v26_: _dafny.Seq
                d_1691_v26_ = _dafny.SeqWithoutIsStrInference([d_1688_v23_, d_1688_v23_])
                d_1690_v25_ = (d_1690_v25_).set((d_1691_v26_)[default__.safeIndex(len(d_1685_v20_), len(d_1691_v26_))], not(d_1660_v2_))
                d_1692_v27_: _dafny.Map
                d_1692_v27_ = _dafny.Map({d_1660_v2_: d_1658_v0_})
                d_1693_v28_: D12
                d_1693_v28_ = D12_DC26(d_1692_v27_)
                d_1694_v29_: _dafny.Seq
                d_1694_v29_ = _dafny.SeqWithoutIsStrInference([d_1689_v24_, d_1689_v24_, d_1689_v24_])
                d_1695_v30_: _dafny.Map
                d_1695_v30_ = _dafny.Map({d_1693_v28_: default__.fm20(d_1660_v2_, globalState)})
                d_1696_v31_: _dafny.Set
                d_1696_v31_ = _dafny.Set({(_dafny.Map({d_1693_v28_: len((d_1694_v29_)[default__.safeIndex(d_1658_v0_, len(d_1694_v29_))])})).set(d_1693_v28_, d_1658_v0_), (d_1695_v30_ if d_1660_v2_ else d_1695_v30_)})
                d_1697_v32_: _dafny.Seq
                d_1697_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhgx"))
                d_1698_v33_: _dafny.Seq
                d_1698_v33_ = _dafny.SeqWithoutIsStrInference([d_1658_v0_, 841])
                d_1699_v34_: _dafny.Seq
                d_1699_v34_ = d_1698_v33_
                d_1700_v35_: _dafny.Map
                d_1700_v35_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([940 for d_1701_i3_ in range(default__.abs(599))]): d_1658_v0_})
                d_1702_v36_: _dafny.Map
                d_1702_v36_ = _dafny.Map({d_1658_v0_: d_1660_v2_})
                d_1696_v31_ = default__.fm45(d_1697_v32_, ((0) - ((0) - (len((d_1699_v34_))))) - (len(d_1700_v35_)), len(d_1702_v36_), globalState)
                d_1703_v37_: _dafny.Array
                nw272_ = _dafny.Array(None, 4)
                nw272_[int(0)] = False
                nw272_[int(1)] = d_1660_v2_
                nw272_[int(2)] = d_1660_v2_
                nw272_[int(3)] = d_1660_v2_
                d_1703_v37_ = nw272_
                d_1704_v38_: D2
                d_1704_v38_ = D2_DC5(d_1703_v37_)
                d_1705_v39_: _dafny.Map
                d_1705_v39_ = _dafny.Map({(d_1704_v38_).cf5: d_1660_v2_})
                d_1706_v40_: _dafny.Array
                nw273_ = _dafny.Array(None, 10)
                nw273_[int(0)] = d_1660_v2_
                nw273_[int(1)] = not(d_1660_v2_)
                nw273_[int(2)] = d_1660_v2_
                nw273_[int(3)] = d_1660_v2_
                nw273_[int(4)] = d_1660_v2_
                nw273_[int(5)] = ((d_1705_v39_)[d_1703_v37_] if (d_1703_v37_) in (d_1705_v39_) else d_1660_v2_)
                nw273_[int(6)] = d_1660_v2_
                nw273_[int(7)] = d_1660_v2_
                nw273_[int(8)] = d_1660_v2_
                nw273_[int(9)] = False
                d_1706_v40_ = nw273_
                d_1707_v41_: _dafny.Map
                d_1707_v41_ = _dafny.Map({d_1706_v40_: d_1658_v0_})
                d_1707_v41_ = (d_1707_v41_).set(d_1703_v37_, -116)
                d_1708_v42_: C5
                nw274_ = C5()
                nw274_.ctor__()
                d_1708_v42_ = nw274_
            d_1709_v43_: str
            d_1710_v44_: _dafny.Set
            d_1711_v45_: bool
            out26_: str
            out27_: _dafny.Set
            out28_: bool
            out26_, out27_, out28_ = (self).m5(d_1660_v2_, d_1658_v0_, len(default__.fm46(globalState)), d_1658_v0_, globalState)
            d_1709_v43_ = out26_
            d_1710_v44_ = out27_
            d_1711_v45_ = out28_
            d_1711_v45_ = d_1660_v2_
            d_1712_v46_: _dafny.Seq
            d_1712_v46_ = _dafny.SeqWithoutIsStrInference([d_1711_v45_, d_1660_v2_, d_1711_v45_])
            d_1713_v47_: D4
            d_1713_v47_ = D4_DC10(d_1712_v46_)
            d_1714_v48_: _dafny.Seq
            d_1714_v48_ = _dafny.SeqWithoutIsStrInference([d_1713_v47_])
            r1 = (len(d_1714_v48_)) <= (d_1658_v0_)
            d_1715_v49_: D4
            d_1715_v49_ = D4_DC10(d_1712_v46_)
            d_1716_v50_: D4
            d_1716_v50_ = D4_DC12(d_1715_v49_)
            d_1717_v51_: D4
            d_1717_v51_ = D4_DC12(d_1715_v49_)
            source31_ = (D4_DC12(D4_DC12(d_1715_v49_)) if d_1660_v2_ else d_1717_v51_)
            if source31_.is_DC11:
                d_1718___mcc_h0_ = source31_.cf14
                d_1719___mcc_h1_ = source31_.cf15
                d_1720___mcc_h2_ = source31_.cf16
                d_1721___mcc_h3_ = source31_.cf17
                d_1722_cf17_ = d_1721___mcc_h3_
                d_1723_cf16_ = d_1720___mcc_h2_
                d_1724_cf15_ = d_1719___mcc_h1_
                d_1725_cf14_ = d_1718___mcc_h0_
                d_1726_v52_: C5
                nw275_ = C5()
                nw275_.ctor__()
                d_1726_v52_ = nw275_
                (globalState).f3 = (d_1658_v0_) + (d_1658_v0_)
                (globalState).f4 = d_1660_v2_
                d_1727_v53_: _dafny.Map
                d_1727_v53_ = _dafny.Map({d_1660_v2_: (d_1723_cf16_) + (d_1723_cf16_)})
                d_1728_v54_: _dafny.MultiSet
                d_1728_v54_ = _dafny.MultiSet([d_1658_v0_])
                d_1727_v53_ = (d_1727_v53_).set((d_1726_v52_).fm3(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1711_v45_: d_1722_cf17_})) for d_1729_i4_ in range(default__.abs(259))]), d_1728_v54_, d_1660_v2_, globalState), d_1658_v0_)
            elif source31_.is_DC10:
                d_1730___mcc_h4_ = source31_.cf13
                d_1731_cf13_ = d_1730___mcc_h4_
                d_1732_v55_: _dafny.Array
                nw276_ = _dafny.Array(_dafny.Seq({}), 29)
                d_1732_v55_ = nw276_
                d_1733_v56_: _dafny.Array
                nw277_ = _dafny.Array(int(0), 15)
                d_1733_v56_ = nw277_
                d_1734_v57_: _dafny.Seq
                d_1734_v57_ = _dafny.SeqWithoutIsStrInference([d_1733_v56_, d_1733_v56_, d_1733_v56_])
                index305_ = default__.safeIndex(700, (d_1732_v55_).length(0))
                (d_1732_v55_)[index305_] = d_1734_v57_
                index306_ = default__.safeIndex(700, (d_1732_v55_).length(0))
                (d_1732_v55_)[index306_] = (d_1734_v57_).set(default__.safeIndex(-642, len(d_1734_v57_)), d_1733_v56_)
                d_1735_v58_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1735_v58_ = nw278_
                index307_ = default__.safeIndex(275, (d_1735_v58_).length(0))
                (d_1735_v58_)[index307_] = d_1733_v56_
                index308_ = default__.safeIndex(275, (d_1735_v58_).length(0))
                (d_1735_v58_)[index308_] = d_1733_v56_
                d_1736_v59_: _dafny.Seq
                d_1736_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkabsjr"))
                d_1737_v60_: _dafny.Map
                d_1737_v60_ = _dafny.Map({(d_1658_v0_) < (len(d_1736_v59_)): d_1711_v45_})
                d_1737_v60_ = (d_1737_v60_).set(d_1660_v2_, d_1660_v2_)
                d_1738_v61_: _dafny.Array
                nw279_ = _dafny.Array(None, 24)
                nw279_[int(0)] = d_1660_v2_
                nw279_[int(1)] = d_1711_v45_
                nw279_[int(2)] = d_1660_v2_
                nw279_[int(3)] = d_1660_v2_
                nw279_[int(4)] = d_1711_v45_
                nw279_[int(5)] = False
                nw279_[int(6)] = True
                nw279_[int(7)] = d_1711_v45_
                nw279_[int(8)] = d_1711_v45_
                nw279_[int(9)] = d_1711_v45_
                nw279_[int(10)] = False
                nw279_[int(11)] = True
                nw279_[int(12)] = d_1711_v45_
                nw279_[int(13)] = not(d_1711_v45_)
                nw279_[int(14)] = d_1711_v45_
                nw279_[int(15)] = d_1660_v2_
                nw279_[int(16)] = d_1660_v2_
                nw279_[int(17)] = d_1660_v2_
                nw279_[int(18)] = not(True)
                nw279_[int(19)] = not(d_1711_v45_)
                nw279_[int(20)] = d_1711_v45_
                nw279_[int(21)] = d_1711_v45_
                nw279_[int(22)] = d_1660_v2_
                nw279_[int(23)] = False
                d_1738_v61_ = nw279_
                (self).m6((self).fm10(globalState), d_1738_v61_, (d_1658_v0_) * (454), globalState)
            elif True:
                d_1739___mcc_h5_ = source31_.cf18
                d_1740_cf18_ = d_1739___mcc_h5_
                d_1741_v62_: _dafny.Seq
                d_1741_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwbkil"))
                d_1742_v63_: D4
                d_1742_v63_ = D4_DC11(d_1711_v45_, d_1711_v45_, d_1658_v0_, d_1741_v62_)
                d_1743_v64_: _dafny.MultiSet
                d_1743_v64_ = _dafny.MultiSet([d_1658_v0_, d_1658_v0_])
                d_1744_v65_: _dafny.Set
                d_1744_v65_ = _dafny.Set({d_1660_v2_, False})
                d_1745_v66_: _dafny.Map
                d_1745_v66_ = _dafny.Map({d_1743_v64_: (d_1744_v65_) - (d_1744_v65_)})
                d_1746_v67_: D9
                d_1746_v67_ = D9_DC21(d_1711_v45_, d_1712_v46_, d_1711_v45_)
                rhs301_ = default__.fm30((self).fm10(globalState), d_1746_v67_, d_1660_v2_, d_1658_v0_, globalState)
                rhs302_ = d_1745_v66_
                rhs303_ = default__.fm25(globalState)
                rhs304_ = ((d_1658_v0_) - ((0) - (d_1658_v0_))) + (len(((d_1742_v63_).cf17).set(default__.safeIndex(d_1658_v0_, len((d_1742_v63_).cf17)), d_1709_v43_)))
                rhs305_ = d_1658_v0_
                lhs240_ = globalState
                lhs241_ = globalState
                d_1742_v63_ = rhs301_
                d_1745_v66_ = rhs302_
                d_1741_v62_ = rhs303_
                lhs240_.f8 = rhs304_
                lhs241_.f8 = rhs305_
                d_1747_v68_: C2
                nw280_ = C2()
                nw280_.ctor__()
                d_1747_v68_ = nw280_
                rhs306_ = (d_1712_v46_)[default__.safeIndex(d_1658_v0_, len(d_1712_v46_))]
                rhs307_ = d_1660_v2_
                rhs308_ = default__.safeDivisionInt(len(d_1741_v62_), d_1658_v0_)
                rhs309_ = d_1744_v65_
                rhs310_ = (d_1711_v45_ if False else d_1711_v45_)
                lhs242_ = globalState
                lhs243_ = globalState
                lhs244_ = globalState
                lhs242_.f4 = rhs306_
                d_1711_v45_ = rhs307_
                lhs243_.f8 = rhs308_
                d_1744_v65_ = rhs309_
                lhs244_.f4 = rhs310_
                (globalState).f4 = (d_1741_v62_) in (_dafny.Set({d_1741_v62_, d_1741_v62_}))
        elif True:
            d_1748_v69_: _dafny.Map
            d_1748_v69_ = _dafny.Map({d_1660_v2_: d_1660_v2_})
            d_1749_v70_: _dafny.Seq
            d_1749_v70_ = _dafny.SeqWithoutIsStrInference([d_1748_v69_, d_1748_v69_])
            d_1750_v71_: _dafny.Array
            nw281_ = _dafny.Array(None, 3)
            nw281_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1748_v69_, d_1748_v69_, d_1748_v69_])
            nw281_[int(1)] = d_1749_v70_
            nw281_[int(2)] = d_1749_v70_
            d_1750_v71_ = nw281_
            index309_ = default__.safeIndex(932, (d_1750_v71_).length(0))
            (d_1750_v71_)[index309_] = d_1749_v70_
            index310_ = default__.safeIndex(932, (d_1750_v71_).length(0))
            (d_1750_v71_)[index310_] = ((d_1749_v70_) + (d_1749_v70_)) + (d_1749_v70_)
            (globalState).f8 = d_1658_v0_
            d_1751_v72_: C1
            nw282_ = C1()
            nw282_.ctor__()
            d_1751_v72_ = nw282_
            d_1752_v73_: _dafny.Seq
            d_1752_v73_ = _dafny.SeqWithoutIsStrInference([d_1660_v2_, d_1660_v2_, False])
            d_1753_v74_: _dafny.Seq
            d_1753_v74_ = _dafny.SeqWithoutIsStrInference([d_1658_v0_])
            d_1754_v75_: _dafny.Seq
            d_1754_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            d_1755_v76_: _dafny.MultiSet
            d_1755_v76_ = _dafny.MultiSet([len(d_1754_v75_)])
            d_1756_v77_: _dafny.Set
            d_1756_v77_ = _dafny.Set({not((d_1752_v73_)[default__.safeIndex(d_1658_v0_, len(d_1752_v73_))])})
            d_1757_v78_: str
            d_1757_v78_ = _dafny.CodePoint('v')
            d_1758_v79_: _dafny.MultiSet
            d_1758_v79_ = _dafny.MultiSet([d_1757_v78_])
            d_1759_v80_: _dafny.Array
            nw283_ = _dafny.Array(None, 23)
            nw283_[int(0)] = (d_1752_v73_)[default__.safeIndex(d_1658_v0_, len(d_1752_v73_))]
            nw283_[int(1)] = d_1660_v2_
            nw283_[int(2)] = not(d_1660_v2_)
            nw283_[int(3)] = d_1660_v2_
            nw283_[int(4)] = d_1660_v2_
            nw283_[int(5)] = d_1660_v2_
            nw283_[int(6)] = (d_1658_v0_) >= ((self).fm10(globalState))
            nw283_[int(7)] = d_1660_v2_
            nw283_[int(8)] = (d_1751_v72_).fm3(d_1753_v74_, d_1755_v76_, d_1660_v2_, globalState)
            nw283_[int(9)] = d_1660_v2_
            nw283_[int(10)] = d_1660_v2_
            nw283_[int(11)] = (d_1756_v77_).issubset(d_1756_v77_)
            nw283_[int(12)] = True
            nw283_[int(13)] = d_1660_v2_
            nw283_[int(14)] = (d_1758_v79_) != (d_1758_v79_)
            nw283_[int(15)] = not((d_1662_v4_))
            nw283_[int(16)] = d_1660_v2_
            nw283_[int(17)] = d_1660_v2_
            nw283_[int(18)] = not(d_1660_v2_)
            nw283_[int(19)] = d_1660_v2_
            nw283_[int(20)] = d_1660_v2_
            nw283_[int(21)] = d_1660_v2_
            nw283_[int(22)] = d_1660_v2_
            d_1759_v80_ = nw283_
            index311_ = default__.safeIndex(957, (d_1759_v80_).length(0))
            (d_1759_v80_)[index311_] = (d_1660_v2_) == ((d_1751_v72_).fm3(_dafny.SeqWithoutIsStrInference([d_1658_v0_ for d_1760_i5_ in range(default__.abs(-537))]), d_1755_v76_, d_1660_v2_, globalState))
            index312_ = default__.safeIndex(957, (d_1759_v80_).length(0))
            (d_1759_v80_)[index312_] = ((d_1754_v75_) + ((d_1754_v75_).set(default__.safeIndex(d_1658_v0_, len(d_1754_v75_)), _dafny.CodePoint('i')))) == ((default__.fm17(d_1757_v78_, globalState)) + (d_1754_v75_))
            d_1660_v2_ = (_dafny.MultiSet((d_1753_v74_) + (d_1753_v74_))).ispropersubset(_dafny.MultiSet((d_1753_v74_).set(default__.safeIndex(-860, len(d_1753_v74_)), d_1658_v0_)))
        if (d_1660_v2_ if d_1660_v2_ else True):
            d_1761_v81_: _dafny.Seq
            d_1761_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfusokk"))
            d_1762_v82_: str
            d_1762_v82_ = _dafny.CodePoint('y')
            d_1763_v83_: _dafny.Map
            d_1763_v83_ = _dafny.Map({d_1762_v82_: d_1660_v2_})
            d_1764_v84_: _dafny.Seq
            d_1764_v84_ = _dafny.SeqWithoutIsStrInference([len(d_1763_v83_)])
            d_1765_v85_: _dafny.Map
            d_1765_v85_ = _dafny.Map({d_1660_v2_: d_1660_v2_})
            d_1766_v86_: _dafny.MultiSet
            d_1766_v86_ = _dafny.MultiSet([len(d_1765_v85_)])
            d_1767_v87_: _dafny.Array
            nw284_ = _dafny.Array(None, 18)
            nw284_[int(0)] = d_1660_v2_
            nw284_[int(1)] = True
            nw284_[int(2)] = d_1660_v2_
            nw284_[int(3)] = d_1660_v2_
            nw284_[int(4)] = d_1660_v2_
            nw284_[int(5)] = d_1660_v2_
            nw284_[int(6)] = d_1660_v2_
            nw284_[int(7)] = d_1660_v2_
            nw284_[int(8)] = d_1660_v2_
            nw284_[int(9)] = d_1660_v2_
            nw284_[int(10)] = False
            nw284_[int(11)] = d_1660_v2_
            nw284_[int(12)] = d_1660_v2_
            nw284_[int(13)] = False
            nw284_[int(14)] = (self).fm3(d_1764_v84_, d_1766_v86_, d_1660_v2_, globalState)
            nw284_[int(15)] = d_1660_v2_
            nw284_[int(16)] = d_1660_v2_
            nw284_[int(17)] = d_1660_v2_
            d_1767_v87_ = nw284_
            d_1768_v88_: _dafny.Map
            d_1768_v88_ = _dafny.Map({d_1658_v0_: d_1767_v87_})
            d_1769_v89_: _dafny.Map
            d_1769_v89_ = _dafny.Map({(d_1761_v81_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1770_i6_ in range(default__.abs(565))])): d_1768_v88_})
            d_1769_v89_ = (d_1769_v89_).set(d_1761_v81_, d_1768_v88_)
            if d_1660_v2_:
                d_1761_v81_ = (d_1761_v81_) + (d_1761_v81_)
                (globalState).f3 = d_1658_v0_
                d_1761_v81_ = d_1761_v81_
                d_1771_v90_: _dafny.Array
                def lambda89_(d_1772_v82_):
                    def lambda90_(d_1773_i7_):
                        return D10_DC23(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1772_v82_ for d_1774_i8_ in range(default__.abs(93))]), _dafny.SeqWithoutIsStrInference([d_1772_v82_ for d_1775_i9_ in range(default__.abs(-809))])]))

                    return lambda90_

                init46_ = lambda89_(d_1762_v82_)
                nw285_ = _dafny.Array(None, 23)
                for i0_46_ in range(nw285_.length(0)):
                    nw285_[i0_46_] = init46_(i0_46_)
                d_1771_v90_ = nw285_
                d_1776_v91_: _dafny.MultiSet
                d_1776_v91_ = _dafny.MultiSet([d_1761_v81_])
                d_1777_v92_: D10
                d_1777_v92_ = D10_DC23(d_1776_v91_)
                index313_ = default__.safeIndex(5, (d_1771_v90_).length(0))
                (d_1771_v90_)[index313_] = d_1777_v92_
                index314_ = default__.safeIndex(5, (d_1771_v90_).length(0))
                (d_1771_v90_)[index314_] = d_1777_v92_
                (globalState).f4 = (d_1658_v0_) == (default__.safeModuloInt(d_1658_v0_, len(default__.fm25(globalState))))
            elif True:
                (globalState).f3 = d_1658_v0_
                (globalState).f8 = d_1658_v0_
                d_1778_v93_: _dafny.Set
                d_1778_v93_ = _dafny.Set({d_1658_v0_, d_1658_v0_, (len(d_1761_v81_)) + (d_1658_v0_), d_1658_v0_})
                (globalState).f8 = (0) - (len(d_1778_v93_))
                d_1779_v94_: _dafny.Seq
                d_1779_v94_ = _dafny.SeqWithoutIsStrInference([d_1660_v2_])
                index315_ = default__.safeIndex(903, (d_1767_v87_).length(0))
                (d_1767_v87_)[index315_] = not(((d_1779_v94_)[default__.safeIndex(d_1658_v0_, len(d_1779_v94_))]) in (d_1779_v94_))
                d_1780_v95_: _dafny.Set
                d_1780_v95_ = _dafny.Set({d_1660_v2_, d_1660_v2_})
                index316_ = default__.safeIndex(903, (d_1767_v87_).length(0))
                rhs311_ = (d_1660_v2_ if d_1660_v2_ else d_1660_v2_)
                rhs312_ = d_1780_v95_
                lhs245_ = d_1767_v87_
                lhs246_ = default__.safeIndex(903, (d_1767_v87_).length(0))
                lhs245_[lhs246_] = rhs311_
                d_1780_v95_ = rhs312_
                d_1765_v85_ = d_1765_v85_
            rhs313_ = True
            rhs314_ = (0) - (default__.safeModuloInt(d_1658_v0_, (d_1658_v0_) - ((0) - (d_1658_v0_))))
            lhs247_ = globalState
            d_1660_v2_ = rhs313_
            lhs247_.f8 = rhs314_
            d_1781_v96_: _dafny.Array
            nw286_ = _dafny.Array(int(0), 19)
            d_1781_v96_ = nw286_
            index317_ = default__.safeIndex(825, (d_1781_v96_).length(0))
            (d_1781_v96_)[index317_] = d_1658_v0_
            index318_ = default__.safeIndex(825, (d_1781_v96_).length(0))
            (d_1781_v96_)[index318_] = d_1658_v0_
            index319_ = default__.safeIndex(825, (d_1781_v96_).length(0))
            rhs315_ = (d_1781_v96_)[default__.safeIndex(825, (d_1781_v96_).length(0))]
            rhs316_ = d_1781_v96_
            lhs248_ = d_1781_v96_
            lhs249_ = default__.safeIndex(825, (d_1781_v96_).length(0))
            lhs248_[lhs249_] = rhs315_
            d_1781_v96_ = rhs316_
        elif True:
            d_1782_v97_: _dafny.Array
            nw287_ = _dafny.Array(_dafny.Map({}), 7)
            d_1782_v97_ = nw287_
            d_1783_v98_: _dafny.Set
            d_1783_v98_ = _dafny.Set({d_1658_v0_})
            d_1784_v99_: _dafny.Seq
            d_1784_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxcadue"))
            d_1785_v100_: _dafny.Map
            d_1785_v100_ = _dafny.Map({d_1783_v98_: d_1784_v99_})
            index320_ = default__.safeIndex(594, (d_1782_v97_).length(0))
            (d_1782_v97_)[index320_] = d_1785_v100_
            d_1786_v101_: _dafny.Array
            def lambda91_(d_1787_i10_):
                return D7_DC18()

            init47_ = lambda91_
            nw288_ = _dafny.Array(None, 17)
            for i0_47_ in range(nw288_.length(0)):
                nw288_[i0_47_] = init47_(i0_47_)
            d_1786_v101_ = nw288_
            d_1788_v102_: D7
            d_1788_v102_ = D7_DC18()
            index321_ = default__.safeIndex(331, (d_1786_v101_).length(0))
            (d_1786_v101_)[index321_] = (d_1788_v102_ if d_1660_v2_ else d_1788_v102_)
            index322_ = default__.safeIndex(594, (d_1782_v97_).length(0))
            index323_ = default__.safeIndex(331, (d_1786_v101_).length(0))
            rhs317_ = ((d_1785_v100_) | (default__.fm47(len(d_1784_v99_), d_1660_v2_, d_1660_v2_, globalState))) | (_dafny.Map({d_1783_v98_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1789_i11_ in range(default__.abs(53))])}))
            rhs318_ = not (d_1660_v2_) or (d_1660_v2_)
            rhs319_ = d_1788_v102_
            lhs250_ = d_1782_v97_
            lhs251_ = default__.safeIndex(594, (d_1782_v97_).length(0))
            lhs252_ = d_1786_v101_
            lhs253_ = default__.safeIndex(331, (d_1786_v101_).length(0))
            lhs250_[lhs251_] = rhs317_
            r0 = rhs318_
            lhs252_[lhs253_] = rhs319_
            d_1790_v103_: C1
            nw289_ = C1()
            nw289_.ctor__()
            d_1790_v103_ = nw289_
            r1 = True
            d_1791_v104_: _dafny.Map
            d_1791_v104_ = _dafny.Map({not (d_1660_v2_) or (d_1660_v2_): d_1658_v0_})
            d_1791_v104_ = d_1791_v104_
            d_1792_v105_: _dafny.Seq
            d_1792_v105_ = _dafny.SeqWithoutIsStrInference([d_1791_v104_])
            d_1793_v106_: _dafny.Seq
            d_1793_v106_ = _dafny.SeqWithoutIsStrInference([d_1658_v0_, d_1658_v0_, d_1658_v0_, d_1658_v0_, d_1658_v0_])
            d_1794_v107_: _dafny.Seq
            d_1794_v107_ = _dafny.SeqWithoutIsStrInference([(d_1792_v105_)[default__.safeIndex(d_1658_v0_, len(d_1792_v105_))], d_1791_v104_, d_1791_v104_, (_dafny.Map({d_1660_v2_: len(d_1793_v106_)})) | (d_1791_v104_)])
            d_1795_v108_: _dafny.Map
            d_1795_v108_ = _dafny.Map({d_1658_v0_: d_1660_v2_})
            d_1794_v107_ = (d_1792_v105_).set(default__.safeIndex(len((d_1795_v108_).set(507, d_1660_v2_)), len(d_1792_v105_)), default__.fm46(globalState))
        d_1796_v109_: _dafny.Array
        def lambda92_(d_1797_i13_):
            return _dafny.CodePoint('h')

        init48_ = lambda92_
        nw290_ = _dafny.Array(None, 5)
        for i0_48_ in range(nw290_.length(0)):
            nw290_[i0_48_] = init48_(i0_48_)
        d_1796_v109_ = nw290_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1796_v109_).length(0)):
            d_1798_i12_: int = guard_loop_3_
            if (True) and (((0) <= (d_1798_i12_)) and ((d_1798_i12_) < ((d_1796_v109_).length(0)))):
                (d_1796_v109_)[(d_1798_i12_)] = _dafny.CodePoint('h')
        r0 = not(d_1660_v2_)
        d_1799_v110_: _dafny.Map
        d_1799_v110_ = _dafny.Map({d_1659_v1_: d_1658_v0_})
        d_1800_v111_: _dafny.Set
        d_1800_v111_ = _dafny.Set({d_1799_v110_})
        r1 = (d_1800_v111_).ispropersubset(d_1800_v111_)
        return r0, r1

    def m5(self, p0, p1, p2, p3, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Set = _dafny.Set({})
        r2: bool = False
        (globalState).f8 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1801_i0_ in range(default__.abs(217))]))
        d_1802_v0_: _dafny.Set
        d_1802_v0_ = _dafny.Set({p2, p2})
        d_1803_v1_: _dafny.Array
        def lambda93_(d_1804_i2_):
            return True

        init49_ = lambda93_
        nw291_ = _dafny.Array(None, 5)
        for i0_49_ in range(nw291_.length(0)):
            nw291_[i0_49_] = init49_(i0_49_)
        d_1803_v1_ = nw291_
        d_1805_v2_: _dafny.Map
        d_1805_v2_ = _dafny.Map({d_1803_v1_: (0) - (p2)})
        d_1806_i1_: int
        d_1806_i1_ = 0
        with _dafny.label("12"):
            while (_dafny.Set({len(d_1805_v2_)})).issubset((d_1802_v0_) - (d_1802_v0_)):
                with _dafny.c_label("12"):
                    if (d_1806_i1_) >= (100):
                        raise _dafny.Break("12")
                    d_1806_i1_ = (d_1806_i1_) + (1)
                    d_1807_v3_: _dafny.Map
                    d_1807_v3_ = _dafny.Map({749: p3})
                    d_1808_v4_: D3
                    d_1808_v4_ = D3_DC8(d_1807_v3_)
                    d_1809_v5_: _dafny.Seq
                    d_1809_v5_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                    d_1810_v6_: _dafny.MultiSet
                    d_1810_v6_ = _dafny.MultiSet([d_1809_v5_])
                    (globalState).f8 = (default__.fm26(d_1808_v4_, p1, d_1810_v6_, globalState)) - (p1)
                    source32_ = p0
                    d_1811___mcc_h0_ = source32_
                    d_1812_cf4_ = d_1811___mcc_h0_
                    d_1813_v7_: _dafny.Seq
                    d_1813_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwtkmlf"))
                    d_1814_v8_: _dafny.MultiSet
                    d_1814_v8_ = _dafny.MultiSet([d_1813_v7_])
                    (globalState).f0 = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(p1, (d_1814_v8_).cardinality), ((d_1807_v3_)[p1] if (p1) in (d_1807_v3_) else p2), default__.fm20(p0, globalState), p3])
                    (globalState).f4 = p0
                    nw292_ = _dafny.Array(False, 14)
                    d_1803_v1_ = nw292_
                    d_1813_v7_ = default__.fm14(p3, globalState)
                    index324_ = default__.safeIndex(29, (d_1803_v1_).length(0))
                    (d_1803_v1_)[index324_] = p0
                    index325_ = default__.safeIndex(29, (d_1803_v1_).length(0))
                    (d_1803_v1_)[index325_] = p0
                    d_1815_v9_: str
                    d_1815_v9_ = _dafny.CodePoint('e')
                    d_1816_v10_: _dafny.Seq
                    d_1816_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrrugumfh"))
                    d_1817_v11_: _dafny.Map
                    d_1817_v11_ = _dafny.Map({d_1815_v9_: d_1816_v10_})
                    d_1818_v12_: _dafny.Seq
                    d_1818_v12_ = _dafny.SeqWithoutIsStrInference([p2, p2, len(((d_1817_v11_)[d_1815_v9_] if (d_1815_v9_) in (d_1817_v11_) else d_1816_v10_)), p2, len(d_1816_v10_)])
                    (globalState).f0 = (d_1818_v12_) + (d_1818_v12_)
                    pass
            pass
        d_1819_v13_: _dafny.Seq
        d_1819_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uba"))
        d_1820_v14_: _dafny.Set
        d_1820_v14_ = _dafny.Set({d_1819_v13_, d_1819_v13_})
        d_1821_v15_: D13
        d_1821_v15_ = D13_DC29(d_1820_v14_)
        d_1822_v16_: _dafny.Map
        d_1822_v16_ = _dafny.Map({p0: d_1821_v15_})
        d_1823_v17_: D13
        d_1823_v17_ = D13_DC31(((d_1822_v16_)[p0] if (p0) in (d_1822_v16_) else d_1821_v15_))
        source33_ = d_1823_v17_
        if source33_.is_DC30:
            d_1824___mcc_h1_ = source33_.cf42
            d_1825___mcc_h2_ = source33_.cf43
            d_1826___mcc_h3_ = source33_.cf44
            d_1827___mcc_h4_ = source33_.cf45
            d_1828_cf45_ = d_1827___mcc_h4_
            d_1829_cf44_ = d_1826___mcc_h3_
            d_1830_cf43_ = d_1825___mcc_h2_
            d_1831_cf42_ = d_1824___mcc_h1_
            (globalState).f3 = ((p3) * (p3)) + (p1)
            d_1819_v13_ = d_1819_v13_
            if p0:
                (globalState).f4 = not (p0) or (p0)
                d_1832_v18_: _dafny.MultiSet
                d_1832_v18_ = _dafny.MultiSet([638])
                d_1833_v19_: _dafny.Map
                d_1833_v19_ = _dafny.Map({(p1) * (p3): (d_1832_v18_).issubset(d_1832_v18_)})
                d_1834_v20_: _dafny.Seq
                d_1834_v20_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1835_v21_: _dafny.Map
                d_1835_v21_ = _dafny.Map({d_1831_cf42_: _dafny.MultiSet([len(d_1834_v20_)])})
                d_1833_v19_ = (d_1833_v19_).set(default__.safeDivisionInt(p3, p3), (((d_1835_v21_)[d_1830_cf43_] if (d_1830_cf43_) in (d_1835_v21_) else _dafny.MultiSet(d_1834_v20_))).ispropersubset(d_1832_v18_))
                d_1836_v22_: _dafny.Map
                d_1836_v22_ = _dafny.Map({_dafny.CodePoint('y'): d_1831_cf42_})
                d_1836_v22_ = (d_1836_v22_).set(_dafny.CodePoint('n'), d_1831_cf42_)
                d_1837_v23_: T0
                nw293_ = C2()
                nw293_.ctor__()
                d_1837_v23_ = nw293_
                d_1838_v24_: _dafny.Map
                d_1838_v24_ = _dafny.Map({not((self).fm3(d_1834_v20_, _dafny.MultiSet(d_1834_v20_), d_1831_cf42_, globalState)): _dafny.CodePoint('s')})
                d_1839_v25_: _dafny.Map
                d_1839_v25_ = _dafny.Map({d_1831_cf42_: d_1819_v13_})
                d_1840_v26_: D19
                d_1840_v26_ = D19_DC39(d_1831_cf42_, p1)
                d_1841_v27_: _dafny.Array
                nw294_ = _dafny.Array(int(0), 3)
                d_1841_v27_ = nw294_
                d_1842_v28_: _dafny.Map
                d_1842_v28_ = _dafny.Map({p1: d_1841_v27_})
                d_1843_v29_: _dafny.Map
                d_1843_v29_ = _dafny.Map({False: p3})
                d_1844_v30_: _dafny.Seq
                d_1844_v30_ = _dafny.SeqWithoutIsStrInference([((d_1842_v28_)[((d_1843_v29_)[d_1830_cf43_] if (d_1830_cf43_) in (d_1843_v29_) else p3)] if (((d_1843_v29_)[d_1830_cf43_] if (d_1830_cf43_) in (d_1843_v29_) else p3)) in (d_1842_v28_) else d_1841_v27_)])
                d_1845_v31_: _dafny.Map
                d_1845_v31_ = _dafny.Map({(d_1840_v26_).cf58: (d_1844_v30_)[default__.safeIndex(len(_dafny.Map({p3: d_1831_cf42_})), len(d_1844_v30_))]})
                d_1846_v32_: _dafny.Map
                d_1846_v32_ = _dafny.Map({d_1845_v31_: (d_1837_v23_ if p0 else d_1837_v23_)})
                d_1847_v33_: str
                d_1847_v33_ = _dafny.CodePoint('p')
                d_1848_v34_: _dafny.Map
                d_1848_v34_ = _dafny.Map({p3: d_1847_v33_})
                rhs320_ = (d_1829_cf44_ if (p2) != (len(((d_1839_v25_)[not(p0)] if (not(p0)) in (d_1839_v25_) else d_1819_v13_))) else d_1829_cf44_)
                rhs321_ = ((d_1846_v32_)[d_1845_v31_] if (d_1845_v31_) in (d_1846_v32_) else (d_1837_v23_ if d_1830_cf43_ else d_1837_v23_))
                rhs322_ = (default__.fm43(d_1830_cf43_, p3, d_1847_v33_, globalState)) == (d_1848_v34_)
                rhs323_ = (default__.fm11(d_1831_cf42_, d_1830_cf43_, globalState)) | (d_1838_v24_)
                rhs324_ = p0
                lhs254_ = globalState
                r1 = rhs320_
                d_1837_v23_ = rhs321_
                lhs254_.f4 = rhs322_
                d_1838_v24_ = rhs323_
                r2 = rhs324_
                index326_ = default__.safeIndex(783, (d_1841_v27_).length(0))
                (d_1841_v27_)[index326_] = 462
                index327_ = default__.safeIndex(783, (d_1841_v27_).length(0))
                (d_1841_v27_)[index327_] = ((680) - (p3) if False else p2)
            elif True:
                d_1849_v35_: _dafny.Array
                nw295_ = _dafny.Array(int(0), 15)
                d_1849_v35_ = nw295_
                d_1850_v36_: _dafny.Array
                d_1850_v36_ = d_1849_v35_
                d_1850_v36_ = d_1850_v36_
                d_1851_v37_: C2
                nw296_ = C2()
                nw296_.ctor__()
                d_1851_v37_ = nw296_
                d_1852_v38_: _dafny.Map
                d_1852_v38_ = _dafny.Map({d_1851_v37_: d_1819_v13_})
                d_1853_v39_: _dafny.MultiSet
                d_1853_v39_ = _dafny.MultiSet([False])
                d_1854_v40_: _dafny.Map
                d_1854_v40_ = _dafny.Map({(p1) > (p1): (d_1853_v39_).ispropersubset(d_1853_v39_)})
                index328_ = default__.safeIndex(205, (d_1849_v35_).length(0))
                (d_1849_v35_)[index328_] = (p2) * (-434)
                index329_ = default__.safeIndex(205, (d_1849_v35_).length(0))
                rhs325_ = d_1852_v38_
                rhs326_ = p0
                rhs327_ = ((d_1854_v40_) | (d_1854_v40_)).set(not(d_1830_cf43_), (p1) > (p2))
                rhs328_ = p2
                lhs255_ = globalState
                lhs256_ = d_1849_v35_
                lhs257_ = default__.safeIndex(205, (d_1849_v35_).length(0))
                d_1852_v38_ = rhs325_
                lhs255_.f4 = rhs326_
                d_1854_v40_ = rhs327_
                lhs256_[lhs257_] = rhs328_
                d_1831_cf42_ = d_1830_cf43_
                d_1855_v41_: _dafny.Seq
                d_1855_v41_ = _dafny.SeqWithoutIsStrInference([203, p2, p3])
                d_1856_v42_: _dafny.MultiSet
                d_1856_v42_ = _dafny.MultiSet([p2])
                index330_ = default__.safeIndex(205, (d_1849_v35_).length(0))
                (d_1849_v35_)[index330_] = ((d_1855_v41_)[default__.safeIndex(((d_1856_v42_)[(d_1849_v35_)[default__.safeIndex(205, (d_1849_v35_).length(0))]] if ((d_1849_v35_)[default__.safeIndex(205, (d_1849_v35_).length(0))]) in (d_1856_v42_) else (d_1849_v35_)[default__.safeIndex(205, (d_1849_v35_).length(0))]), len(d_1855_v41_))]) * ((-356) + (p2))
                d_1857_v43_: _dafny.MultiSet
                d_1857_v43_ = d_1856_v42_
                d_1858_v44_: _dafny.MultiSet
                d_1858_v44_ = _dafny.MultiSet([d_1857_v43_, d_1857_v43_, d_1857_v43_])
                d_1859_v45_: _dafny.Map
                d_1859_v45_ = _dafny.Map({d_1853_v39_: (d_1858_v44_).cardinality})
                d_1860_v46_: D12
                d_1860_v46_ = D12_DC27(p1, p2, 826, len(d_1859_v45_))
                d_1861_v47_: _dafny.Map
                d_1861_v47_ = _dafny.Map({not(p0): d_1860_v46_})
                d_1862_v48_: str
                d_1862_v48_ = _dafny.CodePoint('r')
                d_1863_v49_: _dafny.Map
                d_1863_v49_ = _dafny.Map({len(default__.fm31(d_1831_cf42_, p0, d_1830_cf43_, globalState)): p3})
                d_1864_v50_: D3
                d_1864_v50_ = D3_DC8(d_1863_v49_)
                d_1861_v47_ = (d_1861_v47_).set((p2) in (default__.fm1((d_1849_v35_)[default__.safeIndex(205, (d_1849_v35_).length(0))], d_1862_v48_, d_1830_cf43_, (d_1864_v50_).cf12, globalState)), d_1860_v46_)
            (globalState).f8 = p2
        elif source33_.is_DC29:
            d_1865___mcc_h5_ = source33_.cf41
            d_1866_cf41_ = d_1865___mcc_h5_
            (globalState).f8 = (self).fm10(globalState)
            d_1867_v51_: _dafny.Array
            nw297_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1867_v51_ = nw297_
            d_1868_v52_: _dafny.Array
            nw298_ = _dafny.Array(_dafny.Map({}), 9)
            d_1868_v52_ = nw298_
            index331_ = default__.safeIndex(990, (d_1867_v51_).length(0))
            (d_1867_v51_)[index331_] = d_1868_v52_
            index332_ = default__.safeIndex(990, (d_1867_v51_).length(0))
            (d_1867_v51_)[index332_] = d_1868_v52_
            d_1869_v53_: _dafny.Array
            def lambda94_(d_1870_p0_):
                def lambda95_(d_1871_i3_):
                    return (d_1871_i3_) + (len(_dafny.Map({True: d_1870_p0_})))

                return lambda95_

            init50_ = lambda94_(p0)
            nw299_ = _dafny.Array(None, 23)
            for i0_50_ in range(nw299_.length(0)):
                nw299_[i0_50_] = init50_(i0_50_)
            d_1869_v53_ = nw299_
            d_1872_v54_: D21
            d_1872_v54_ = D21_DC42(_dafny.Map({d_1869_v53_: 731}))
            (self).m6(default__.safeDivisionInt(724, len(d_1819_v13_)), d_1803_v1_, len((d_1872_v54_).cf62), globalState)
            (globalState).f8 = p1
        elif True:
            d_1873___mcc_h6_ = source33_.cf46
            d_1874_cf46_ = d_1873___mcc_h6_
            d_1819_v13_ = d_1819_v13_
            d_1875_v55_: C1
            nw300_ = C1()
            nw300_.ctor__()
            d_1875_v55_ = nw300_
            d_1876_v56_: D4
            d_1876_v56_ = D4_DC11(p0, True, -284, d_1819_v13_)
            d_1877_v57_: D4
            d_1877_v57_ = D4_DC12(d_1876_v56_)
            d_1878_v58_: D4
            d_1878_v58_ = D4_DC12(d_1876_v56_)
            d_1879_v59_: D4
            d_1879_v59_ = D4_DC12(d_1876_v56_)
            source34_ = (d_1879_v59_ if (False if p0 else p0) else d_1879_v59_)
            if source34_.is_DC11:
                d_1880___mcc_h7_ = source34_.cf14
                d_1881___mcc_h8_ = source34_.cf15
                d_1882___mcc_h9_ = source34_.cf16
                d_1883___mcc_h10_ = source34_.cf17
                d_1884_cf17_ = d_1883___mcc_h10_
                d_1885_cf16_ = d_1882___mcc_h9_
                d_1886_cf15_ = d_1881___mcc_h8_
                d_1887_cf14_ = d_1880___mcc_h7_
                index333_ = default__.safeIndex(268, (d_1803_v1_).length(0))
                (d_1803_v1_)[index333_] = False
                index334_ = default__.safeIndex(268, (d_1803_v1_).length(0))
                (d_1803_v1_)[index334_] = True
                d_1888_v60_: _dafny.Array
                def lambda96_(d_1889_cf14_, d_1890_p3_, d_1891_p2_):
                    def lambda97_(d_1892_i4_):
                        return _dafny.Map({d_1889_cf14_: len(_dafny.Map({d_1890_p3_: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1891_p2_]))])).cardinality}))})

                    return lambda97_

                init51_ = lambda96_(d_1887_cf14_, p3, p2)
                nw301_ = _dafny.Array(None, 5)
                for i0_51_ in range(nw301_.length(0)):
                    nw301_[i0_51_] = init51_(i0_51_)
                d_1888_v60_ = nw301_
                d_1893_v61_: C4
                nw302_ = C4()
                nw302_.ctor__(d_1888_v60_)
                d_1893_v61_ = nw302_
                d_1888_v60_ = (d_1888_v60_ if (d_1803_v1_)[default__.safeIndex(268, (d_1803_v1_).length(0))] else (d_1893_v61_).f12)
                d_1894_v62_: C3
                nw303_ = C3()
                nw303_.ctor__()
                d_1894_v62_ = nw303_
            elif source34_.is_DC10:
                d_1895___mcc_h11_ = source34_.cf13
                d_1896_cf13_ = d_1895___mcc_h11_
                d_1897_v63_: str
                d_1897_v63_ = _dafny.CodePoint('k')
                (globalState).f4 = (len((d_1819_v13_).set(default__.safeIndex(len(d_1819_v13_), len(d_1819_v13_)), d_1897_v63_))) != (p1)
                d_1898_v64_: C3
                nw304_ = C3()
                nw304_.ctor__()
                d_1898_v64_ = nw304_
                d_1899_v65_: _dafny.MultiSet
                d_1899_v65_ = _dafny.MultiSet([d_1898_v64_])
                d_1900_v66_: _dafny.MultiSet
                d_1900_v66_ = _dafny.MultiSet([p0])
                d_1901_v67_: _dafny.MultiSet
                d_1901_v67_ = d_1900_v66_
                d_1902_v68_: _dafny.Map
                d_1902_v68_ = _dafny.Map({p2: d_1901_v67_})
                rhs329_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1903_i5_ in range(default__.abs(78))])
                rhs330_ = d_1899_v65_
                rhs331_ = ((0) - (p3)) * (len(d_1902_v68_))
                lhs258_ = globalState
                d_1819_v13_ = rhs329_
                d_1899_v65_ = rhs330_
                lhs258_.f8 = rhs331_
                d_1904_v69_: _dafny.Map
                d_1904_v69_ = _dafny.Map({p3: p2})
                d_1905_v70_: _dafny.MultiSet
                d_1905_v70_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0, p0])])
                d_1906_v71_: _dafny.MultiSet
                d_1906_v71_ = _dafny.MultiSet([default__.fm26(D3_DC8(d_1904_v69_), p2, d_1905_v70_, globalState), p1, p1])
                (d_1875_v55_).m13(768, d_1906_v71_, p0, globalState)
                rhs332_ = not(p0)
                rhs333_ = default__.safeDivisionInt(p1, p3)
                rhs334_ = d_1897_v63_
                rhs335_ = default__.safeDivisionInt(p2, p3)
                lhs259_ = globalState
                lhs260_ = globalState
                lhs261_ = globalState
                lhs259_.f4 = rhs332_
                lhs260_.f8 = rhs333_
                r0 = rhs334_
                lhs261_.f3 = rhs335_
            elif True:
                d_1907___mcc_h12_ = source34_.cf18
                d_1908_cf18_ = d_1907___mcc_h12_
                d_1909_v72_: T0
                nw305_ = C1()
                nw305_.ctor__()
                d_1909_v72_ = nw305_
                d_1910_v73_: D2
                d_1910_v73_ = D2_DC6(len(d_1802_v0_), p0, p0)
                d_1911_v74_: _dafny.Map
                d_1911_v74_ = _dafny.Map({(d_1910_v73_).cf7: d_1909_v72_})
                d_1909_v72_ = (d_1909_v72_ if p0 else ((d_1911_v74_)[p0] if (p0) in (d_1911_v74_) else d_1909_v72_))
                d_1912_v75_: C3
                nw306_ = C3()
                nw306_.ctor__()
                d_1912_v75_ = nw306_
                (globalState).f4 = (p0 if p0 else p0)
                d_1913_v76_: _dafny.Map
                d_1913_v76_ = _dafny.Map({d_1875_v55_: p1})
                d_1914_v77_: _dafny.Map
                d_1914_v77_ = _dafny.Map({p0: True})
                d_1915_v78_: _dafny.MultiSet
                d_1915_v78_ = _dafny.MultiSet([p3, (0) - (p2), p2, 934, len(d_1914_v77_)])
                d_1916_v79_: _dafny.Array
                nw307_ = _dafny.Array(None, 8)
                nw307_[int(0)] = (len(d_1913_v76_)) + (p2)
                nw307_[int(1)] = p1
                nw307_[int(2)] = p2
                nw307_[int(3)] = (p1) - (p3)
                nw307_[int(4)] = 206
                nw307_[int(5)] = (833) - (len(d_1819_v13_))
                nw307_[int(6)] = (d_1915_v78_).cardinality
                nw307_[int(7)] = p3
                d_1916_v79_ = nw307_
                index335_ = default__.safeIndex(557, (d_1916_v79_).length(0))
                (d_1916_v79_)[index335_] = p1
                d_1917_v80_: _dafny.Seq
                d_1917_v80_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1918_v81_: _dafny.Map
                d_1918_v81_ = _dafny.Map({p2: len(d_1917_v80_)})
                d_1919_v82_: _dafny.Seq
                d_1919_v82_ = _dafny.SeqWithoutIsStrInference([(d_1912_v75_).fm15(p2, globalState)])
                d_1920_v83_: _dafny.Seq
                d_1920_v83_ = _dafny.SeqWithoutIsStrInference([(d_1919_v82_)[default__.safeIndex(782, len(d_1919_v82_))], p3, -919, (0) - (p1), p1])
                index336_ = default__.safeIndex(557, (d_1916_v79_).length(0))
                (d_1916_v79_)[index336_] = (len((d_1918_v81_).set(len(d_1919_v82_), -642))) - ((p1) + (160))
            d_1921_v84_: _dafny.Map
            d_1921_v84_ = _dafny.Map({p0: not(p0)})
            d_1921_v84_ = (d_1921_v84_).set(True, p0)
        d_1922_v85_: _dafny.Seq
        d_1922_v85_ = _dafny.SeqWithoutIsStrInference([p1, -829, p2, len(d_1819_v13_)])
        (globalState).f3 = (d_1922_v85_)[default__.safeIndex(162, len(d_1922_v85_))]
        (globalState).f3 = len((default__.fm28(p1, globalState)) | (d_1802_v0_))
        d_1923_v86_: _dafny.Map
        d_1923_v86_ = _dafny.Map({p3: p1})
        d_1924_v87_: D3
        d_1924_v87_ = D3_DC8(d_1923_v86_)
        d_1925_v89_: _dafny.Seq
        d_1925_v89_ = _dafny.SeqWithoutIsStrInference([True])
        d_1926_v90_: _dafny.MultiSet
        d_1926_v90_ = _dafny.MultiSet([p1])
        d_1927_v91_: _dafny.MultiSet
        d_1927_v91_ = _dafny.MultiSet([(d_1925_v89_).set(default__.safeIndex(p3, len(d_1925_v89_)), True), default__.fm31(p0, p0, (self).fm3(d_1922_v85_, d_1926_v90_, p0, globalState), globalState)])
        d_1928_v92_: str
        d_1928_v92_ = _dafny.CodePoint('v')
        pat_let_tv66_ = d_1923_v86_
        d_1929_v93_: C6
        nw308_ = C6()
        def iife206_(_pat_let55_0):
            def iife207_(d_1930_dt__update__tmp_h0_):
                def iife208_(_pat_let56_0):
                    def iife209_(d_1931_dt__update_hcf12_h0_):
                        return D3_DC8(d_1931_dt__update_hcf12_h0_)
                    return iife209_(_pat_let56_0)
                return iife208_(pat_let_tv66_)
            return iife207_(_pat_let55_0)
        def iife210_():
            coll96_ = _dafny.Map()
            compr_96_: int
            for compr_96_ in _dafny.IntegerRange(51, -278):
                d_1932_v88_: int = compr_96_
                if ((51) <= (d_1932_v88_)) and ((d_1932_v88_) < (-278)):
                    coll96_[default__.safeModuloInt(d_1932_v88_, p2)] = True
            return _dafny.Map(coll96_)
        nw308_.ctor__(p1, (d_1819_v13_).set(default__.safeIndex(default__.fm26(iife206_(d_1924_v87_), len(iife210_()
        ), d_1927_v91_, globalState), len(d_1819_v13_)), d_1928_v92_))
        d_1929_v93_ = nw308_
        r0 = d_1928_v92_
        r1 = d_1802_v0_
        r2 = (p3) < (p1)
        return r0, r1, r2

    def m6(self, p0, p1, p2, globalState):
        d_1933_v0_: bool
        d_1933_v0_ = True
        index337_ = default__.safeIndex(108, (p1).length(0))
        (p1)[index337_] = d_1933_v0_
        index338_ = default__.safeIndex(108, (p1).length(0))
        (p1)[index338_] = False
        d_1934_v1_: str
        d_1934_v1_ = _dafny.CodePoint('m')
        d_1935_v2_: _dafny.MultiSet
        d_1935_v2_ = _dafny.MultiSet([d_1934_v1_, d_1934_v1_, d_1934_v1_, d_1934_v1_])
        d_1936_v3_: _dafny.Seq
        d_1936_v3_ = _dafny.SeqWithoutIsStrInference([d_1934_v1_])
        d_1937_v4_: _dafny.MultiSet
        d_1937_v4_ = d_1935_v2_
        d_1938_v5_: _dafny.Map
        d_1938_v5_ = _dafny.Map({len(d_1936_v3_): p2})
        d_1939_v6_: D3
        d_1939_v6_ = D3_DC8(d_1938_v5_)
        d_1940_v7_: _dafny.Seq
        d_1940_v7_ = _dafny.SeqWithoutIsStrInference([d_1933_v0_, (p1)[default__.safeIndex(108, (p1).length(0))], d_1933_v0_])
        d_1941_v8_: _dafny.MultiSet
        d_1941_v8_ = _dafny.MultiSet([d_1940_v7_, d_1940_v7_])
        d_1942_v9_: D2
        d_1942_v9_ = D2_DC7(d_1934_v1_, False, p2)
        pat_let_tv67_ = p1
        pat_let_tv68_ = p1
        pat_let_tv69_ = d_1934_v1_
        index339_ = default__.safeIndex(108, (p1).length(0))
        def iife211_(_pat_let57_0):
            def iife212_(d_1943_dt__update__tmp_h0_):
                def iife213_(_pat_let58_0):
                    def iife214_(d_1944_dt__update_hcf10_h0_):
                        def iife215_(_pat_let59_0):
                            def iife216_(d_1945_dt__update_hcf9_h0_):
                                return D2_DC7(d_1945_dt__update_hcf9_h0_, d_1944_dt__update_hcf10_h0_, (d_1943_dt__update__tmp_h0_).cf11)
                            return iife216_(_pat_let59_0)
                        return iife215_(pat_let_tv69_)
                    return iife214_(_pat_let58_0)
                return iife213_((pat_let_tv68_)[default__.safeIndex(108, (pat_let_tv67_).length(0))])
            return iife212_(_pat_let57_0)
        rhs336_ = (d_1935_v2_) - ((_dafny.MultiSet(d_1936_v3_)) - ((d_1937_v4_)))
        rhs337_ = ((iife211_(d_1942_v9_)).cf10 if (default__.fm26(d_1939_v6_, p0, d_1941_v8_, globalState)) == (default__.fm20((p1)[default__.safeIndex(108, (p1).length(0))], globalState)) else (p1)[default__.safeIndex(108, (p1).length(0))])
        rhs338_ = d_1933_v0_
        lhs262_ = p1
        lhs263_ = default__.safeIndex(108, (p1).length(0))
        d_1935_v2_ = rhs336_
        d_1933_v0_ = rhs337_
        lhs262_[lhs263_] = rhs338_
        index340_ = default__.safeIndex(108, (p1).length(0))
        (p1)[index340_] = not((p1)[default__.safeIndex(108, (p1).length(0))])
        d_1946_v10_: C5
        nw309_ = C5()
        nw309_.ctor__()
        d_1946_v10_ = nw309_
        d_1947_v11_: _dafny.Map
        d_1947_v11_ = _dafny.Map({(p1)[default__.safeIndex(108, (p1).length(0))]: d_1946_v10_})
        hi11_ = p2
        for d_1948_i0_ in range(len(d_1947_v11_), hi11_):
            d_1949_v12_: _dafny.Array
            def lambda98_(d_1950_i1_):
                return (d_1950_i1_) - (310)

            init52_ = lambda98_
            nw310_ = _dafny.Array(None, 18)
            for i0_52_ in range(nw310_.length(0)):
                nw310_[i0_52_] = init52_(i0_52_)
            d_1949_v12_ = nw310_
            index341_ = default__.safeIndex(424, (d_1949_v12_).length(0))
            (d_1949_v12_)[index341_] = 632
            d_1951_v13_: D10
            d_1951_v13_ = D10_DC24(p0, d_1933_v0_, d_1936_v3_)
            index342_ = default__.safeIndex(424, (d_1949_v12_).length(0))
            (d_1949_v12_)[index342_] = default__.safeModuloInt(default__.fm26(d_1939_v6_, (d_1951_v13_).cf31, d_1941_v8_, globalState), p2)
            d_1952_v15_: _dafny.Array
            def lambda99_(d_1953_v1_, d_1954_p1_):
                def lambda100_(d_1955_i2_):
                    def iife217_():
                        coll97_ = _dafny.Map()
                        compr_97_: int
                        for compr_97_ in _dafny.IntegerRange(-720, 104):
                            d_1956_v14_: int = compr_97_
                            if ((-720) <= (d_1956_v14_)) and ((d_1956_v14_) < (104)):
                                coll97_[default__.safeDivisionInt(d_1956_v14_, len(_dafny.SeqWithoutIsStrInference([d_1953_v1_ for d_1957_i3_ in range(default__.abs(858))])))] = (d_1954_p1_)[default__.safeIndex(108, (d_1954_p1_).length(0))]
                        return _dafny.Map(coll97_)
                    return (iife217_()
                    ) | (_dafny.Map({27: (d_1954_p1_)[default__.safeIndex(108, (d_1954_p1_).length(0))]}))

                return lambda100_

            init53_ = lambda99_(d_1934_v1_, p1)
            nw311_ = _dafny.Array(None, 23)
            for i0_53_ in range(nw311_.length(0)):
                nw311_[i0_53_] = init53_(i0_53_)
            d_1952_v15_ = nw311_
            d_1958_v16_: _dafny.Map
            d_1958_v16_ = _dafny.Map({(d_1949_v12_)[default__.safeIndex(424, (d_1949_v12_).length(0))]: d_1933_v0_})
            d_1959_v17_: _dafny.Map
            d_1959_v17_ = _dafny.Map({-945: ((d_1958_v16_)[len(d_1940_v7_)] if (len(d_1940_v7_)) in (d_1958_v16_) else d_1933_v0_)})
            index343_ = default__.safeIndex(802, (d_1952_v15_).length(0))
            (d_1952_v15_)[index343_] = d_1959_v17_
            index344_ = default__.safeIndex(802, (d_1952_v15_).length(0))
            index345_ = default__.safeIndex(108, (p1).length(0))
            rhs339_ = d_1958_v16_
            rhs340_ = d_1933_v0_
            rhs341_ = (_dafny.SeqWithoutIsStrInference([d_1934_v1_ for d_1960_i4_ in range(default__.abs(139))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpl")) if d_1933_v0_ else d_1936_v3_))
            lhs264_ = d_1952_v15_
            lhs265_ = default__.safeIndex(802, (d_1952_v15_).length(0))
            lhs266_ = p1
            lhs267_ = default__.safeIndex(108, (p1).length(0))
            lhs264_[lhs265_] = rhs339_
            lhs266_[lhs267_] = rhs340_
            d_1936_v3_ = rhs341_
            d_1933_v0_ = d_1933_v0_
            d_1961_v18_: C3
            nw312_ = C3()
            nw312_.ctor__()
            d_1961_v18_ = nw312_
        d_1962_v19_: _dafny.Set
        d_1962_v19_ = _dafny.Set({d_1936_v3_})
        if (d_1962_v19_).isdisjoint(d_1962_v19_):
            d_1963_v20_: _dafny.Seq
            d_1963_v20_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1964_v21_: _dafny.MultiSet
            d_1964_v21_ = _dafny.MultiSet([p2])
            d_1965_v22_: _dafny.Set
            d_1965_v22_ = _dafny.Set({p2})
            d_1966_v23_: C0
            nw313_ = C0()
            nw313_.ctor__()
            d_1966_v23_ = nw313_
            d_1967_v24_: D13
            d_1967_v24_ = D13_DC30((p1)[default__.safeIndex(108, (p1).length(0))], (d_1946_v10_).fm3(d_1963_v20_, d_1964_v21_, (p1)[default__.safeIndex(108, (p1).length(0))], globalState), (d_1965_v22_) - (d_1965_v22_), d_1966_v23_)
            d_1967_v24_ = D13_DC30((d_1965_v22_).issubset(d_1965_v22_), (p1)[default__.safeIndex(108, (p1).length(0))], d_1965_v22_, d_1966_v23_)
            d_1968_v25_: _dafny.MultiSet
            d_1968_v25_ = _dafny.MultiSet([d_1936_v3_, d_1936_v3_])
            d_1969_v26_: D10
            d_1969_v26_ = D10_DC23(d_1968_v25_)
            source35_ = d_1969_v26_
            if source35_.is_DC24:
                d_1970___mcc_h0_ = source35_.cf31
                d_1971___mcc_h1_ = source35_.cf32
                d_1972___mcc_h2_ = source35_.cf33
                d_1973_cf33_ = d_1972___mcc_h2_
                d_1974_cf32_ = d_1971___mcc_h1_
                d_1975_cf31_ = d_1970___mcc_h0_
                d_1976_v27_: _dafny.Array
                def lambda101_(d_1977_p2_):
                    def lambda102_(d_1978_i5_):
                        return (d_1978_i5_) * (d_1977_p2_)

                    return lambda102_

                init54_ = lambda101_(p2)
                nw314_ = _dafny.Array(None, 2)
                for i0_54_ in range(nw314_.length(0)):
                    nw314_[i0_54_] = init54_(i0_54_)
                d_1976_v27_ = nw314_
                index346_ = default__.safeIndex(117, (d_1976_v27_).length(0))
                (d_1976_v27_)[index346_] = -106
                index347_ = default__.safeIndex(117, (d_1976_v27_).length(0))
                (d_1976_v27_)[index347_] = d_1975_cf31_
                index348_ = default__.safeIndex(108, (p1).length(0))
                (p1)[index348_] = (p1)[default__.safeIndex(108, (p1).length(0))]
                (globalState).f8 = p2
                d_1979_v28_: _dafny.Array
                def lambda103_(d_1980_v1_):
                    def lambda104_(d_1981_i6_):
                        return d_1980_v1_

                    return lambda104_

                init55_ = lambda103_(d_1934_v1_)
                nw315_ = _dafny.Array(None, 19)
                for i0_55_ in range(nw315_.length(0)):
                    nw315_[i0_55_] = init55_(i0_55_)
                d_1979_v28_ = nw315_
                d_1982_v29_: _dafny.Map
                d_1982_v29_ = _dafny.Map({d_1979_v28_: 459})
                d_1982_v29_ = d_1982_v29_
            elif True:
                d_1983___mcc_h3_ = source35_.cf30
                d_1984_cf30_ = d_1983___mcc_h3_
                d_1985_v30_: C0
                nw316_ = C0()
                nw316_.ctor__()
                d_1985_v30_ = nw316_
                d_1933_v0_ = (p1)[default__.safeIndex(108, (p1).length(0))]
                (globalState).f8 = p0
                d_1986_v31_: D2
                d_1986_v31_ = D2_DC5(p1)
                d_1986_v31_ = d_1986_v31_
            d_1987_v32_: C1
            nw317_ = C1()
            nw317_.ctor__()
            d_1987_v32_ = nw317_
            d_1988_v33_: D7
            d_1988_v33_ = D7_DC18()
            d_1989_v34_: _dafny.Map
            d_1989_v34_ = _dafny.Map({(d_1967_v24_).cf43: d_1988_v33_})
            d_1990_v35_: _dafny.Map
            d_1990_v35_ = _dafny.Map({len(d_1965_v22_): (p1)[default__.safeIndex(108, (p1).length(0))]})
            if not((d_1989_v34_) == (default__.fm48(d_1963_v20_, d_1990_v35_, globalState))):
                (globalState).f8 = p2
                d_1991_v36_: _dafny.Map
                d_1991_v36_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, p0, p2, p2, len(d_1940_v7_)]): (d_1964_v21_).intersection(d_1964_v21_)})
                d_1992_v37_: _dafny.Seq
                d_1992_v37_ = _dafny.SeqWithoutIsStrInference([d_1964_v21_])
                d_1991_v36_ = (d_1991_v36_).set(d_1963_v20_, (d_1992_v37_)[default__.safeIndex(p0, len(d_1992_v37_))])
                index349_ = default__.safeIndex(108, (p1).length(0))
                (p1)[index349_] = ((d_1987_v32_).fm3(d_1963_v20_, d_1964_v21_, (p1)[default__.safeIndex(108, (p1).length(0))], globalState))
                d_1993_v38_: _dafny.Array
                nw318_ = _dafny.Array(False, 27)
                d_1993_v38_ = nw318_
                d_1993_v38_ = d_1993_v38_
                d_1994_v39_: _dafny.Array
                nw319_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1994_v39_ = nw319_
                index350_ = default__.safeIndex(872, (d_1994_v39_).length(0))
                (d_1994_v39_)[index350_] = d_1940_v7_
                index351_ = default__.safeIndex(872, (d_1994_v39_).length(0))
                (d_1994_v39_)[index351_] = d_1940_v7_
            elif True:
                (globalState).f4 = True
                d_1995_v40_: D19
                d_1995_v40_ = D19_DC39(False, len(d_1936_v3_))
                d_1996_v41_: _dafny.Map
                d_1996_v41_ = _dafny.Map({d_1934_v1_: p2})
                (globalState).f8 = ((0) - ((d_1995_v40_).cf59)) * (((d_1996_v41_)[d_1934_v1_] if (d_1934_v1_) in (d_1996_v41_) else p2))
                index352_ = default__.safeIndex(108, (p1).length(0))
                (p1)[index352_] = True
                (globalState).f3 = (39) * (default__.fm20((p1)[default__.safeIndex(108, (p1).length(0))], globalState))
                index353_ = default__.safeIndex(108, (p1).length(0))
                (p1)[index353_] = (p2) not in (d_1990_v35_)
            d_1997_v42_: _dafny.Array
            def lambda105_(d_1998_v0_, d_1999_p2_):
                def lambda106_(d_2000_i7_):
                    return _dafny.Map({d_1998_v0_: d_1999_p2_})

                return lambda106_

            init56_ = lambda105_(d_1933_v0_, p2)
            nw320_ = _dafny.Array(None, 14)
            for i0_56_ in range(nw320_.length(0)):
                nw320_[i0_56_] = init56_(i0_56_)
            d_1997_v42_ = nw320_
            d_2001_v43_: _dafny.Map
            d_2001_v43_ = _dafny.Map({False: len(d_1936_v3_)})
            index354_ = default__.safeIndex(240, (d_1997_v42_).length(0))
            (d_1997_v42_)[index354_] = (d_2001_v43_) | (d_2001_v43_)
            index355_ = default__.safeIndex(240, (d_1997_v42_).length(0))
            (d_1997_v42_)[index355_] = d_2001_v43_
        elif True:
            d_2002_v44_: _dafny.Set
            d_2002_v44_ = _dafny.Set({d_1933_v0_, not(True), (p1)[default__.safeIndex(108, (p1).length(0))]})
            (globalState).f3 = len((d_2002_v44_) | (d_2002_v44_))
            d_2003_v45_: _dafny.Array
            nw321_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_2003_v45_ = nw321_
            d_2004_v46_: _dafny.MultiSet
            d_2004_v46_ = _dafny.MultiSet([244])
            d_2005_v47_: _dafny.Array
            nw322_ = _dafny.Array(None, 21)
            nw322_[int(0)] = d_2003_v45_
            nw322_[int(1)] = d_2003_v45_
            nw322_[int(2)] = d_2003_v45_
            nw322_[int(3)] = (d_2003_v45_ if (self).fm3(_dafny.SeqWithoutIsStrInference([p2, p0, p0]), d_2004_v46_, (p1)[default__.safeIndex(108, (p1).length(0))], globalState) else d_2003_v45_)
            nw322_[int(4)] = d_2003_v45_
            nw322_[int(5)] = (d_2003_v45_ if not(d_1933_v0_) else d_2003_v45_)
            nw322_[int(6)] = d_2003_v45_
            nw322_[int(7)] = d_2003_v45_
            nw322_[int(8)] = d_2003_v45_
            nw322_[int(9)] = d_2003_v45_
            nw322_[int(10)] = d_2003_v45_
            nw322_[int(11)] = d_2003_v45_
            nw322_[int(12)] = d_2003_v45_
            nw322_[int(13)] = d_2003_v45_
            nw322_[int(14)] = d_2003_v45_
            nw322_[int(15)] = d_2003_v45_
            nw322_[int(16)] = d_2003_v45_
            nw322_[int(17)] = d_2003_v45_
            nw322_[int(18)] = d_2003_v45_
            nw322_[int(19)] = d_2003_v45_
            nw322_[int(20)] = d_2003_v45_
            d_2005_v47_ = nw322_
            index356_ = default__.safeIndex(320, (d_2005_v47_).length(0))
            (d_2005_v47_)[index356_] = d_2003_v45_
            d_2006_v48_: _dafny.Array
            nw323_ = _dafny.Array(None, 11)
            nw323_[int(0)] = d_1934_v1_
            nw323_[int(1)] = d_1934_v1_
            nw323_[int(2)] = d_1934_v1_
            nw323_[int(3)] = d_1934_v1_
            nw323_[int(4)] = d_1934_v1_
            nw323_[int(5)] = d_1934_v1_
            nw323_[int(6)] = d_1934_v1_
            nw323_[int(7)] = d_1934_v1_
            nw323_[int(8)] = d_1934_v1_
            nw323_[int(9)] = d_1934_v1_
            nw323_[int(10)] = d_1934_v1_
            d_2006_v48_ = nw323_
            d_2007_v49_: _dafny.Set
            d_2007_v49_ = _dafny.Set({p1})
            d_2008_v50_: _dafny.Seq
            d_2008_v50_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1, p1}), d_2007_v49_])
            d_2009_v51_: _dafny.Seq
            d_2009_v51_ = _dafny.SeqWithoutIsStrInference([len((d_2008_v50_)[default__.safeIndex((_dafny.MultiSet(d_1940_v7_)).cardinality, len(d_2008_v50_))])])
            index357_ = default__.safeIndex(320, (d_2005_v47_).length(0))
            rhs342_ = d_2003_v45_
            rhs343_ = ((_dafny.SeqWithoutIsStrInference([d_1934_v1_ for d_2010_i8_ in range(default__.abs(772))]) if (d_1946_v10_).fm3(d_2009_v51_, _dafny.MultiSet([436]), (p1)[default__.safeIndex(108, (p1).length(0))], globalState) else (d_1936_v3_) + (d_1936_v3_))).set(default__.safeIndex(len(d_2009_v51_), len((_dafny.SeqWithoutIsStrInference([d_1934_v1_ for d_2011_i8_ in range(default__.abs(772))]) if (d_1946_v10_).fm3(d_2009_v51_, _dafny.MultiSet([436]), (p1)[default__.safeIndex(108, (p1).length(0))], globalState) else (d_1936_v3_) + (d_1936_v3_)))), d_1934_v1_)
            rhs344_ = (D23_DC45(d_2006_v48_)).cf66
            lhs268_ = d_2005_v47_
            lhs269_ = default__.safeIndex(320, (d_2005_v47_).length(0))
            lhs268_[lhs269_] = rhs342_
            d_1936_v3_ = rhs343_
            d_2006_v48_ = rhs344_
            d_2012_v52_: _dafny.Array
            nw324_ = _dafny.Array(int(0), 7)
            d_2012_v52_ = nw324_
            d_2013_v53_: _dafny.Array
            d_2013_v53_ = d_2012_v52_
            source36_ = d_2013_v53_
            d_2014___mcc_h4_ = source36_
            d_2015_cf55_ = d_2014___mcc_h4_
            (globalState).f4 = (p1)[default__.safeIndex(108, (p1).length(0))]
            d_2016_v54_: _dafny.Seq
            d_2016_v54_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2017_v55_: _dafny.Array
            nw325_ = _dafny.Array(None, 16)
            nw325_[int(0)] = p1
            nw325_[int(1)] = p1
            nw325_[int(2)] = p1
            nw325_[int(3)] = p1
            nw325_[int(4)] = p1
            nw325_[int(5)] = p1
            nw325_[int(6)] = p1
            nw325_[int(7)] = p1
            nw325_[int(8)] = (d_2016_v54_)[default__.safeIndex(190, len(d_2016_v54_))]
            nw325_[int(9)] = p1
            nw325_[int(10)] = p1
            nw325_[int(11)] = p1
            nw325_[int(12)] = p1
            nw325_[int(13)] = p1
            nw325_[int(14)] = p1
            nw325_[int(15)] = p1
            d_2017_v55_ = nw325_
            d_2018_v56_: _dafny.Array
            nw326_ = _dafny.Array(None, 13)
            d_2018_v56_ = nw326_
            d_2019_v57_: _dafny.Set
            d_2019_v57_ = _dafny.Set({d_2018_v56_, d_2018_v56_})
            rhs345_ = d_1936_v3_
            rhs346_ = d_2017_v55_
            rhs347_ = (len(d_2019_v57_)) - (((d_1938_v5_)[p2] if (p2) in (d_1938_v5_) else p2))
            lhs270_ = globalState
            d_1936_v3_ = rhs345_
            d_2017_v55_ = rhs346_
            lhs270_.f8 = rhs347_
            (globalState).f8 = p0
            d_2020_v58_: _dafny.Map
            d_2020_v58_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sl")): _dafny.SeqWithoutIsStrInference([d_1934_v1_ for d_2021_i9_ in range(default__.abs(711))])})
            d_1933_v0_ = ((d_1940_v7_)[default__.safeIndex(len(d_2020_v58_), len(d_1940_v7_))]) or (d_1933_v0_)
            (globalState).f4 = (d_2004_v46_).isdisjoint(d_2004_v46_)
            d_2022_v59_: _dafny.Map
            d_2022_v59_ = _dafny.Map({p0: d_2002_v44_})
            d_2023_v60_: _dafny.Seq
            d_2023_v60_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2024_v61_: _dafny.Array
            nw327_ = _dafny.Array(None, 22)
            nw327_[int(0)] = p1
            nw327_[int(1)] = p1
            nw327_[int(2)] = p1
            nw327_[int(3)] = p1
            nw327_[int(4)] = p1
            nw327_[int(5)] = p1
            nw327_[int(6)] = (d_2023_v60_)[default__.safeIndex(67, len(d_2023_v60_))]
            nw327_[int(7)] = p1
            nw327_[int(8)] = (d_2023_v60_)[default__.safeIndex(len(d_2009_v51_), len(d_2023_v60_))]
            nw327_[int(9)] = p1
            nw327_[int(10)] = p1
            nw327_[int(11)] = p1
            nw327_[int(12)] = p1
            nw327_[int(13)] = p1
            nw327_[int(14)] = p1
            nw327_[int(15)] = p1
            nw327_[int(16)] = p1
            nw327_[int(17)] = p1
            nw327_[int(18)] = p1
            nw327_[int(19)] = p1
            nw327_[int(20)] = p1
            nw327_[int(21)] = p1
            d_2024_v61_ = nw327_
            index358_ = default__.safeIndex(770, (d_2024_v61_).length(0))
            (d_2024_v61_)[index358_] = p1
            d_2025_v63_: D12
            def iife218_():
                coll98_ = _dafny.Map()
                compr_98_: int
                for compr_98_ in _dafny.IntegerRange(-216, 152):
                    d_2026_v62_: int = compr_98_
                    if ((-216) <= (d_2026_v62_)) and ((d_2026_v62_) < (152)):
                        coll98_[(d_2026_v62_) - (p2)] = (p1)[default__.safeIndex(108, (p1).length(0))]
                return _dafny.Map(coll98_)
            d_2025_v63_ = D12_DC27(len(iife218_()
), p2, len(d_1936_v3_), (0) - (p0))
            index359_ = default__.safeIndex(770, (d_2024_v61_).length(0))
            rhs348_ = (d_2025_v63_).cf39
            rhs349_ = ((d_2022_v59_).set(p0, d_2002_v44_)) | (d_2022_v59_)
            rhs350_ = False
            rhs351_ = d_1934_v1_
            rhs352_ = p1
            lhs271_ = globalState
            lhs272_ = globalState
            lhs273_ = d_2024_v61_
            lhs274_ = default__.safeIndex(770, (d_2024_v61_).length(0))
            lhs271_.f3 = rhs348_
            d_2022_v59_ = rhs349_
            lhs272_.f4 = rhs350_
            d_1934_v1_ = rhs351_
            lhs273_[lhs274_] = rhs352_
        d_2027_v65_: _dafny.Map
        def iife219_():
            coll99_ = _dafny.Map()
            compr_99_: str
            for compr_99_ in (d_1936_v3_).Elements:
                d_2028_v64_: str = compr_99_
                if (d_2028_v64_) in (d_1936_v3_):
                    coll99_[d_2028_v64_] = (p1)[default__.safeIndex(108, (p1).length(0))]
            return _dafny.Map(coll99_)
        d_2027_v65_ = _dafny.Map({d_1934_v1_: iife219_()
        })
        d_2029_v66_: _dafny.Seq
        d_2029_v66_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, len(d_2027_v65_)])
        if (d_2029_v66_) in (_dafny.SeqWithoutIsStrInference([d_2029_v66_ for d_2030_i10_ in range(default__.abs(724))])):
            (globalState).f8 = (0) - (p2)
            d_2031_v67_: _dafny.Map
            d_2031_v67_ = _dafny.Map({(p1)[default__.safeIndex(108, (p1).length(0))]: p0})
            d_2032_v68_: _dafny.Map
            d_2032_v68_ = _dafny.Map({d_2029_v66_: d_1946_v10_})
            (globalState).f8 = (default__.safeDivisionInt(((d_2031_v67_)[d_1933_v0_] if (d_1933_v0_) in (d_2031_v67_) else default__.fm26(d_1939_v6_, p2, d_1941_v8_, globalState)), len(d_1940_v7_))) + (default__.safeDivisionInt((self).fm10(globalState), len(d_2032_v68_)))
            if (p1)[default__.safeIndex(108, (p1).length(0))]:
                d_2033_v69_: D12
                d_2033_v69_ = D12_DC27(852, p0, 886, p0)
                pat_let_tv70_ = p0
                d_2034_v70_: _dafny.Array
                nw328_ = _dafny.Array(None, 18)
                nw328_[int(0)] = d_2033_v69_
                nw328_[int(1)] = d_2033_v69_
                nw328_[int(2)] = d_2033_v69_
                nw328_[int(3)] = d_2033_v69_
                nw328_[int(4)] = d_2033_v69_
                nw328_[int(5)] = D12_DC27(len(_dafny.Map({201: p2})), len(d_1936_v3_), default__.fm20(True, globalState), p0)
                nw328_[int(6)] = D12_DC27(len(d_1938_v5_), len(d_1936_v3_), p2, p2)
                nw328_[int(7)] = d_2033_v69_
                nw328_[int(8)] = d_2033_v69_
                nw328_[int(9)] = d_2033_v69_
                nw328_[int(10)] = d_2033_v69_
                nw328_[int(11)] = d_2033_v69_
                nw328_[int(12)] = d_2033_v69_
                nw328_[int(13)] = D12_DC27(p2, len(d_2031_v67_), p2, p0)
                nw328_[int(14)] = d_2033_v69_
                nw328_[int(15)] = d_2033_v69_
                nw328_[int(16)] = d_2033_v69_
                def iife220_(_pat_let60_0):
                    def iife221_(d_2035_dt__update__tmp_h2_):
                        def iife222_(_pat_let61_0):
                            def iife223_(d_2036_dt__update_hcf36_h0_):
                                def iife224_(_pat_let62_0):
                                    def iife225_(d_2037_dt__update_hcf39_h0_):
                                        return D12_DC27(d_2036_dt__update_hcf36_h0_, (d_2035_dt__update__tmp_h2_).cf37, (d_2035_dt__update__tmp_h2_).cf38, d_2037_dt__update_hcf39_h0_)
                                    return iife225_(_pat_let62_0)
                                return iife224_(pat_let_tv70_)
                            return iife223_(_pat_let61_0)
                        return iife222_(183)
                    return iife221_(_pat_let60_0)
                nw328_[int(17)] = iife220_(d_2033_v69_)
                d_2034_v70_ = nw328_
                d_2038_v71_: _dafny.Array
                nw329_ = _dafny.Array(None, 14)
                nw329_[int(0)] = (d_2034_v70_ if d_1933_v0_ else d_2034_v70_)
                nw329_[int(1)] = d_2034_v70_
                nw329_[int(2)] = d_2034_v70_
                nw329_[int(3)] = d_2034_v70_
                nw329_[int(4)] = d_2034_v70_
                nw329_[int(5)] = d_2034_v70_
                nw329_[int(6)] = d_2034_v70_
                nw329_[int(7)] = d_2034_v70_
                nw329_[int(8)] = d_2034_v70_
                nw329_[int(9)] = d_2034_v70_
                nw329_[int(10)] = d_2034_v70_
                nw329_[int(11)] = d_2034_v70_
                nw329_[int(12)] = (d_2034_v70_ if (p1)[default__.safeIndex(108, (p1).length(0))] else d_2034_v70_)
                nw329_[int(13)] = d_2034_v70_
                d_2038_v71_ = nw329_
                index360_ = default__.safeIndex(102, (d_2038_v71_).length(0))
                (d_2038_v71_)[index360_] = d_2034_v70_
                index361_ = default__.safeIndex(102, (d_2038_v71_).length(0))
                (d_2038_v71_)[index361_] = d_2034_v70_
                d_2031_v67_ = (d_2031_v67_).set(d_1933_v0_, p0)
                d_1933_v0_ = (p1)[default__.safeIndex(108, (p1).length(0))]
                def iife226_():
                    coll100_ = _dafny.Set()
                    compr_100_: str
                    for compr_100_ in (d_1936_v3_).Elements:
                        d_2039_v72_: str = compr_100_
                        if (d_2039_v72_) in (d_1936_v3_):
                            coll100_ = coll100_.union(_dafny.Set([d_2039_v72_]))
                    return _dafny.Set(coll100_)
                (globalState).f8 = len((_dafny.Set({d_1934_v1_})) | (iife226_()
                ))
                d_2040_v73_: _dafny.Map
                d_2040_v73_ = _dafny.Map({(p2) * (len(d_2029_v66_)): default__.fm46(globalState)})
                d_2041_v74_: _dafny.MultiSet
                d_2041_v74_ = _dafny.MultiSet([d_1933_v0_])
                d_2040_v73_ = (d_2040_v73_).set(default__.safeDivisionInt(600, (d_2041_v74_).cardinality), (d_2031_v67_).set(d_1933_v0_, p0))
            elif True:
                d_2042_v75_: _dafny.MultiSet
                d_2042_v75_ = _dafny.MultiSet([p2])
                d_2043_v76_: D4
                d_2043_v76_ = D4_DC10(_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(108, (p1).length(0))], (self).fm3(d_2029_v66_, d_2042_v75_, False, globalState), (p1)[default__.safeIndex(108, (p1).length(0))]]))
                d_2044_v77_: _dafny.Map
                d_2044_v77_ = _dafny.Map({(d_1936_v3_) + (d_1936_v3_): p0})
                rhs353_ = d_1936_v3_
                rhs354_ = D4_DC10((d_1940_v7_).set(default__.safeIndex(p0, len(d_1940_v7_)), not(d_1933_v0_)))
                rhs355_ = ((d_2044_v77_)[d_1936_v3_] if (d_1936_v3_) in (d_2044_v77_) else p2)
                rhs356_ = len(((d_1936_v3_) + (_dafny.SeqWithoutIsStrInference([d_1934_v1_ for d_2045_i11_ in range(default__.abs(230))])) if d_1933_v0_ else _dafny.SeqWithoutIsStrInference([d_1934_v1_ for d_2046_i12_ in range(default__.abs(908))])))
                lhs275_ = globalState
                lhs276_ = globalState
                d_1936_v3_ = rhs353_
                d_2043_v76_ = rhs354_
                lhs275_.f8 = rhs355_
                lhs276_.f8 = rhs356_
                d_1942_v9_ = default__.fm40(d_1933_v0_, globalState)
                d_2047_v78_: _dafny.Map
                d_2047_v78_ = _dafny.Map({p1: (self).fm10(globalState)})
                d_2047_v78_ = d_2047_v78_
                d_2048_v79_: _dafny.Array
                nw330_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_2048_v79_ = nw330_
                index362_ = default__.safeIndex(208, (d_2048_v79_).length(0))
                (d_2048_v79_)[index362_] = d_1934_v1_
                d_2049_v80_: _dafny.Array
                nw331_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2049_v80_ = nw331_
                d_2050_v81_: _dafny.Set
                d_2050_v81_ = _dafny.Set({p0, len(d_1936_v3_), p0})
                index363_ = default__.safeIndex(208, (d_2048_v79_).length(0))
                rhs357_ = _dafny.CodePoint('t')
                rhs358_ = (p1)[default__.safeIndex(108, (p1).length(0))]
                rhs359_ = p0
                rhs360_ = default__.fm18((len(d_1936_v3_)) <= (len(d_2050_v81_)), d_1933_v0_, globalState)
                rhs361_ = d_2049_v80_
                lhs277_ = globalState
                lhs278_ = d_2048_v79_
                lhs279_ = default__.safeIndex(208, (d_2048_v79_).length(0))
                d_1934_v1_ = rhs357_
                d_1933_v0_ = rhs358_
                lhs277_.f8 = rhs359_
                lhs278_[lhs279_] = rhs360_
                d_2049_v80_ = rhs361_
                d_2051_v82_: C1
                nw332_ = C1()
                nw332_.ctor__()
                d_2051_v82_ = nw332_
                d_2051_v82_ = d_2051_v82_
            (globalState).f4 = not(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgljestqn"))) + (d_1936_v3_)) != ((d_1936_v3_) + (d_1936_v3_)))
            d_2052_v83_: _dafny.Set
            d_2052_v83_ = _dafny.Set({not((p1)[default__.safeIndex(108, (p1).length(0))]), d_1933_v0_})
            d_2053_v84_: _dafny.Map
            d_2053_v84_ = _dafny.Map({p1: d_1933_v0_})
            d_2054_v85_: _dafny.Map
            d_2054_v85_ = _dafny.Map({p0: (p1)[default__.safeIndex(108, (p1).length(0))]})
            d_2055_v86_: _dafny.Set
            d_2055_v86_ = _dafny.Set({p2})
            d_2056_v87_: C0
            nw333_ = C0()
            nw333_.ctor__()
            d_2056_v87_ = nw333_
            d_2057_v88_: D13
            d_2057_v88_ = D13_DC30(d_1933_v0_, d_1933_v0_, d_2055_v86_, d_2056_v87_)
            d_2058_v89_: D13
            d_2058_v89_ = D13_DC31(d_2057_v88_)
            d_2059_v90_: _dafny.Set
            d_2059_v90_ = _dafny.Set({d_2058_v89_})
            d_2060_v92_: _dafny.Array
            nw334_ = _dafny.Array(None, 20)
            nw334_[int(0)] = (d_1933_v0_) and (d_1933_v0_)
            nw334_[int(1)] = (-957) <= (p2)
            nw334_[int(2)] = True
            nw334_[int(3)] = (len(d_2052_v83_)) <= (default__.fm20(d_1933_v0_, globalState))
            nw334_[int(4)] = ((d_2053_v84_)[p1] if (p1) in (d_2053_v84_) else (p1)[default__.safeIndex(108, (p1).length(0))])
            nw334_[int(5)] = (p1)[default__.safeIndex(108, (p1).length(0))]
            nw334_[int(6)] = d_1933_v0_
            nw334_[int(7)] = (695) <= (p0)
            nw334_[int(8)] = True
            nw334_[int(9)] = ((d_2054_v85_)[-954] if (-954) in (d_2054_v85_) else d_1933_v0_)
            nw334_[int(10)] = False
            nw334_[int(11)] = (p2) <= (p0)
            nw334_[int(12)] = (d_2059_v90_).issubset(d_2059_v90_)
            nw334_[int(13)] = (p1)[default__.safeIndex(108, (p1).length(0))]
            nw334_[int(14)] = True
            nw334_[int(15)] = d_1933_v0_
            nw334_[int(16)] = True
            def iife227_():
                coll101_ = _dafny.Set()
                compr_101_: int
                for compr_101_ in _dafny.IntegerRange(321, 692):
                    d_2061_v91_: int = compr_101_
                    if ((321) <= (d_2061_v91_)) and ((d_2061_v91_) < (692)):
                        coll101_ = coll101_.union(_dafny.Set([(d_2061_v91_) + (p2)]))
                return _dafny.Set(coll101_)
            nw334_[int(17)] = (len(iife227_()
            )) <= ((d_2029_v66_)[default__.safeIndex(p2, len(d_2029_v66_))])
            nw334_[int(18)] = (p1)[default__.safeIndex(108, (p1).length(0))]
            nw334_[int(19)] = (p1)[default__.safeIndex(108, (p1).length(0))]
            d_2060_v92_ = nw334_
            d_2062_v93_: _dafny.Map
            d_2062_v93_ = _dafny.Map({d_1936_v3_: (p1)[default__.safeIndex(108, (p1).length(0))]})
            d_2063_v94_: _dafny.Map
            d_2063_v94_ = _dafny.Map({d_2062_v93_: d_1933_v0_})
            index364_ = default__.safeIndex(108, (p1).length(0))
            rhs362_ = ((d_2063_v94_)[d_2062_v93_] if (d_2062_v93_) in (d_2063_v94_) else (p1)[default__.safeIndex(108, (p1).length(0))])
            rhs363_ = d_2060_v92_
            lhs280_ = p1
            lhs281_ = default__.safeIndex(108, (p1).length(0))
            lhs280_[lhs281_] = rhs362_
            d_2060_v92_ = rhs363_
        elif True:
            d_1933_v0_ = d_1933_v0_
            d_2064_v95_: C6
            nw335_ = C6()
            nw335_.ctor__(966, d_1936_v3_)
            d_2064_v95_ = nw335_
            (globalState).f8 = (0) - ((0) - ((d_2064_v95_).f10))
            d_2065_v96_: _dafny.Set
            d_2065_v96_ = _dafny.Set({p0})
            (globalState).f3 = (p0) - (default__.safeDivisionInt(len(d_2065_v96_), p2))
            d_2066_v97_: _dafny.Array
            nw336_ = _dafny.Array(int(0), 3)
            d_2066_v97_ = nw336_
            d_2067_v98_: _dafny.MultiSet
            d_2067_v98_ = _dafny.MultiSet([d_2066_v97_])
            rhs364_ = _dafny.MultiSet([d_2066_v97_, d_2066_v97_, d_2066_v97_])
            rhs365_ = default__.fm20(((p1)[default__.safeIndex(108, (p1).length(0))] if True else (p1)[default__.safeIndex(108, (p1).length(0))]), globalState)
            rhs366_ = -268
            rhs367_ = p2
            lhs282_ = globalState
            lhs283_ = globalState
            lhs284_ = globalState
            d_2067_v98_ = rhs364_
            lhs282_.f3 = rhs365_
            lhs283_.f3 = rhs366_
            lhs284_.f8 = rhs367_


class C8:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm8(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsbmpw"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2068_i0_ in range(default__.abs(-960))]))

    def fm9(self, p0, p1, globalState):
        return (len(_dafny.Map({(_dafny.MultiSet([579, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tspj")))])).cardinality: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coabyxfb")) for d_2069_i0_ in range(default__.abs(532))]))}))) + (len((_dafny.Map({False: True})) | (_dafny.Map({True: True}))))

    def m4(self, p0, p1, p2, globalState):
        hi12_ = (p2) - (p2)
        for d_2070_i0_ in range((p2) + (p2), hi12_):
            d_2071_v0_: D0
            d_2071_v0_ = D0_DC2()
            source37_ = d_2071_v0_
            if source37_.is_DC1:
                d_2072___mcc_h0_ = source37_.cf1
                d_2073___mcc_h1_ = source37_.cf2
                d_2074_cf2_ = d_2073___mcc_h1_
                d_2075_cf1_ = d_2072___mcc_h0_
                d_2076_v1_: _dafny.Array
                nw337_ = _dafny.Array(D0.default()(), 4)
                d_2076_v1_ = nw337_
                d_2077_v2_: D0
                d_2077_v2_ = D0_DC0(179)
                index365_ = default__.safeIndex(571, (d_2076_v1_).length(0))
                (d_2076_v1_)[index365_] = d_2077_v2_
                d_2078_v3_: str
                d_2078_v3_ = _dafny.CodePoint('n')
                index366_ = default__.safeIndex(571, (d_2076_v1_).length(0))
                rhs368_ = d_2077_v2_
                rhs369_ = d_2078_v3_
                lhs285_ = d_2076_v1_
                lhs286_ = default__.safeIndex(571, (d_2076_v1_).length(0))
                lhs285_[lhs286_] = rhs368_
                d_2078_v3_ = rhs369_
                d_2079_v4_: _dafny.Array
                nw338_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_2079_v4_ = nw338_
                d_2080_v5_: _dafny.Seq
                d_2080_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvute"))
                index367_ = default__.safeIndex(49, (d_2079_v4_).length(0))
                (d_2079_v4_)[index367_] = d_2080_v5_
                index368_ = default__.safeIndex(49, (d_2079_v4_).length(0))
                (d_2079_v4_)[index368_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjcj"))
                d_2081_v6_: _dafny.Array
                nw339_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_2081_v6_ = nw339_
                index369_ = default__.safeIndex(492, (d_2081_v6_).length(0))
                (d_2081_v6_)[index369_] = d_2078_v3_
                index370_ = default__.safeIndex(492, (d_2081_v6_).length(0))
                (d_2081_v6_)[index370_] = d_2078_v3_
                d_2082_v7_: _dafny.Map
                d_2082_v7_ = _dafny.Map({d_2074_cf2_: p0})
                d_2083_v8_: _dafny.Map
                d_2083_v8_ = _dafny.Map({d_2082_v7_: p0})
                (globalState).f4 = ((len(_dafny.SeqWithoutIsStrInference([d_2078_v3_ for d_2084_i1_ in range(default__.abs(49))]))) * (len(d_2083_v8_))) >= (p2)
            elif source37_.is_DC2:
                (globalState).f4 = (p0) or (p0)
                d_2085_v9_: _dafny.Map
                d_2085_v9_ = _dafny.Map({p2: p0})
                d_2086_v10_: _dafny.MultiSet
                d_2086_v10_ = _dafny.MultiSet([d_2085_v9_])
                d_2087_v11_: _dafny.Set
                d_2087_v11_ = _dafny.Set({d_2070_i0_, d_2070_i0_})
                d_2088_v12_: _dafny.MultiSet
                d_2088_v12_ = _dafny.MultiSet([p0])
                d_2089_v13_: _dafny.Seq
                d_2089_v13_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2090_v14_: _dafny.Array
                nw340_ = _dafny.Array(None, 18)
                nw340_[int(0)] = True
                nw340_[int(1)] = (False) and (p0)
                nw340_[int(2)] = p0
                nw340_[int(3)] = p0
                nw340_[int(4)] = True
                nw340_[int(5)] = p0
                nw340_[int(6)] = (d_2085_v9_) in (d_2086_v10_)
                nw340_[int(7)] = p0
                nw340_[int(8)] = p0
                nw340_[int(9)] = (d_2087_v11_).ispropersubset(d_2087_v11_)
                nw340_[int(10)] = (len(d_2087_v11_)) == (d_2070_i0_)
                nw340_[int(11)] = (((d_2088_v12_)[False] if (False) in (d_2088_v12_) else len(d_2089_v13_))) <= (d_2070_i0_)
                nw340_[int(12)] = p0
                nw340_[int(13)] = not (p0) or (not(p0))
                nw340_[int(14)] = (p0) == (p0)
                nw340_[int(15)] = p0
                nw340_[int(16)] = p0
                nw340_[int(17)] = True
                d_2090_v14_ = nw340_
                index371_ = default__.safeIndex(679, (d_2090_v14_).length(0))
                (d_2090_v14_)[index371_] = p0
                d_2091_v15_: bool
                d_2091_v15_ = p0
                d_2092_v16_: _dafny.Seq
                d_2092_v16_ = _dafny.SeqWithoutIsStrInference([p0, (d_2091_v15_)])
                d_2093_v17_: D2
                d_2093_v17_ = D2_DC5(d_2090_v14_)
                d_2094_v18_: _dafny.MultiSet
                d_2094_v18_ = _dafny.MultiSet([len(d_2089_v13_), len((self).fm8(p2, p0, d_2070_i0_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iutevq")) for d_2095_i2_ in range(default__.abs(472))]), globalState)), len((_dafny.Map({d_2090_v14_: (d_2093_v17_).cf5})).set(d_2090_v14_, d_2090_v14_))])
                index372_ = default__.safeIndex(679, (d_2090_v14_).length(0))
                (d_2090_v14_)[index372_] = (d_2092_v16_)[default__.safeIndex((0) - ((d_2094_v18_).cardinality), len(d_2092_v16_))]
                index373_ = default__.safeIndex(679, (d_2090_v14_).length(0))
                (d_2090_v14_)[index373_] = p0
                (globalState).f3 = default__.safeDivisionInt((((d_2094_v18_)[p2] if (p2) in (d_2094_v18_) else len(d_2089_v13_))) * (d_2070_i0_), d_2070_i0_)
            elif source37_.is_DC0:
                d_2096___mcc_h2_ = source37_.cf0
                d_2097_cf0_ = d_2096___mcc_h2_
                d_2098_v19_: _dafny.Seq
                d_2098_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlnuduto"))
                d_2099_v20_: D2
                d_2099_v20_ = D2_DC6(len(d_2098_v19_), p0, p0)
                d_2100_v21_: D0
                d_2100_v21_ = D0_DC1((d_2099_v20_).cf6, -464)
                d_2101_v22_: _dafny.Map
                d_2101_v22_ = _dafny.Map({p0: d_2100_v21_})
                d_2101_v22_ = d_2101_v22_
                (globalState).f3 = d_2070_i0_
                d_2102_v23_: _dafny.Array
                nw341_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_2102_v23_ = nw341_
                def lambda107_(d_2103_i3_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irpd"))

                init57_ = lambda107_
                nw342_ = _dafny.Array(None, 29)
                for i0_57_ in range(nw342_.length(0)):
                    nw342_[i0_57_] = init57_(i0_57_)
                d_2102_v23_ = nw342_
                d_2104_v24_: _dafny.Set
                d_2104_v24_ = _dafny.Set({p0, p0})
                d_2105_v25_: _dafny.Map
                d_2105_v25_ = _dafny.Map({d_2104_v24_: (self).fm9(p0, d_2104_v24_, globalState)})
                d_2105_v25_ = (d_2105_v25_).set(d_2104_v24_, d_2097_cf0_)
            elif True:
                d_2106___mcc_h3_ = source37_.cf3
                d_2107_cf3_ = d_2106___mcc_h3_
                d_2108_v26_: _dafny.Array
                nw343_ = _dafny.Array(_dafny.Map({}), 10)
                d_2108_v26_ = nw343_
                d_2109_v27_: C4
                nw344_ = C4()
                nw344_.ctor__(d_2108_v26_)
                d_2109_v27_ = nw344_
                d_2110_v28_: _dafny.MultiSet
                d_2110_v28_ = _dafny.MultiSet([d_2070_i0_, 492])
                d_2111_v29_: _dafny.Seq
                d_2111_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjohkdo"))
                d_2112_v30_: _dafny.Seq
                d_2112_v30_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2113_v31_: _dafny.Seq
                d_2113_v31_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2114_v32_: _dafny.Map
                d_2114_v32_ = _dafny.Map({(d_2110_v28_).cardinality: d_2111_v29_})
                d_2115_v33_: _dafny.Array
                nw345_ = _dafny.Array(None, 17)
                nw345_[int(0)] = (d_2070_i0_) in (d_2110_v28_)
                nw345_[int(1)] = p0
                nw345_[int(2)] = (d_2109_v27_).fm13(p0, d_2111_v29_, p0, globalState)
                nw345_[int(3)] = not(not(not(p0)))
                nw345_[int(4)] = not ((d_2112_v30_)[default__.safeIndex(len(d_2113_v31_), len(d_2112_v30_))]) or (p0)
                nw345_[int(5)] = (p2) == (d_2070_i0_)
                nw345_[int(6)] = (d_2109_v27_).fm13(p0, d_2111_v29_, p0, globalState)
                nw345_[int(7)] = True
                nw345_[int(8)] = p0
                nw345_[int(9)] = p0
                nw345_[int(10)] = (p0 if p0 else p0)
                nw345_[int(11)] = p0
                nw345_[int(12)] = p0
                nw345_[int(13)] = p0
                nw345_[int(14)] = p0
                nw345_[int(15)] = p0
                nw345_[int(16)] = (d_2109_v27_).fm13(p0, ((d_2114_v32_)[823] if (823) in (d_2114_v32_) else d_2111_v29_), p0, globalState)
                d_2115_v33_ = nw345_
                d_2116_v34_: _dafny.Set
                d_2116_v34_ = _dafny.Set({p0, True})
                rhs370_ = -61
                rhs371_ = (_dafny.Set({p0, False, not(p0), p0, p0})).issubset((d_2116_v34_).intersection(_dafny.Set({(d_2109_v27_).fm3(d_2113_v31_, _dafny.MultiSet([p2]), p0, globalState)})))
                rhs372_ = d_2115_v33_
                rhs373_ = d_2070_i0_
                lhs287_ = globalState
                lhs288_ = globalState
                lhs289_ = globalState
                lhs287_.f8 = rhs370_
                lhs288_.f4 = rhs371_
                d_2115_v33_ = rhs372_
                lhs289_.f8 = rhs373_
                d_2117_v35_: D4
                d_2117_v35_ = D4_DC10(d_2112_v30_)
                d_2117_v35_ = d_2117_v35_
                d_2118_v36_: _dafny.Map
                d_2118_v36_ = _dafny.Map({default__.safeModuloInt(p2, d_2070_i0_): p0})
                d_2118_v36_ = (d_2118_v36_).set(d_2070_i0_, (p2) != (d_2070_i0_))
            d_2119_v37_: str
            d_2119_v37_ = _dafny.CodePoint('s')
            d_2119_v37_ = d_2119_v37_
            d_2119_v37_ = d_2119_v37_
            d_2120_v38_: _dafny.Array
            nw346_ = _dafny.Array(_dafny.Seq({}), 26)
            d_2120_v38_ = nw346_
            index374_ = default__.safeIndex(175, (d_2120_v38_).length(0))
            (d_2120_v38_)[index374_] = default__.fm31(False, not(p0), p0, globalState)
            d_2121_v39_: _dafny.Seq
            d_2121_v39_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2122_v40_: _dafny.Map
            d_2122_v40_ = _dafny.Map({p0: p0})
            index375_ = default__.safeIndex(175, (d_2120_v38_).length(0))
            (d_2120_v38_)[index375_] = ((d_2121_v39_) + (((d_2121_v39_).set(default__.safeIndex(-670, len(d_2121_v39_)), ((d_2122_v40_)[p0] if (p0) in (d_2122_v40_) else p0))).set(default__.safeIndex(330, len((d_2121_v39_).set(default__.safeIndex(-670, len(d_2121_v39_)), ((d_2122_v40_)[p0] if (p0) in (d_2122_v40_) else p0)))), p0))) + ((d_2121_v39_) + (d_2121_v39_))
        hi13_ = (p2) * (p2)
        for d_2123_i4_ in range(p2, hi13_):
            d_2124_v41_: _dafny.Seq
            d_2124_v41_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2124_v41_ = d_2124_v41_
            (globalState).f8 = p2
            d_2125_v42_: str
            d_2125_v42_ = _dafny.CodePoint('a')
            d_2126_v43_: _dafny.Seq
            d_2126_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vstrg"))
            (globalState).f4 = (d_2125_v42_) not in (d_2126_v43_)
            d_2127_v44_: C6
            nw347_ = C6()
            nw347_.ctor__(d_2123_i4_, (d_2126_v43_) + (d_2126_v43_))
            d_2127_v44_ = nw347_
        (globalState).f4 = not(True)
        d_2128_v45_: _dafny.Seq
        d_2128_v45_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])
        d_2129_v46_: _dafny.Seq
        d_2129_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "am"))
        d_2130_v47_: str
        d_2130_v47_ = _dafny.CodePoint('b')
        d_2128_v45_ = ((_dafny.SeqWithoutIsStrInference([p0, p0])) + (d_2128_v45_)).set(default__.safeIndex(len(((((p1)[p0] if (p0) in (p1) else d_2129_v46_)) + (d_2129_v46_)).set(default__.safeIndex(p2, len((((p1)[p0] if (p0) in (p1) else d_2129_v46_)) + (d_2129_v46_))), d_2130_v47_)), len((_dafny.SeqWithoutIsStrInference([p0, p0])) + (d_2128_v45_))), False)
        if default__.fm49(p0, globalState):
            d_2131_v48_: _dafny.Array
            def lambda108_(d_2132_i5_):
                return False

            init58_ = lambda108_
            nw348_ = _dafny.Array(None, 29)
            for i0_58_ in range(nw348_.length(0)):
                nw348_[i0_58_] = init58_(i0_58_)
            d_2131_v48_ = nw348_
            index376_ = default__.safeIndex(334, (d_2131_v48_).length(0))
            (d_2131_v48_)[index376_] = p0
            index377_ = default__.safeIndex(283, (d_2131_v48_).length(0))
            (d_2131_v48_)[index377_] = p0
            d_2133_v49_: _dafny.Set
            d_2133_v49_ = _dafny.Set({d_2131_v48_})
            index378_ = default__.safeIndex(334, (d_2131_v48_).length(0))
            index379_ = default__.safeIndex(283, (d_2131_v48_).length(0))
            rhs374_ = len((_dafny.Set({d_2131_v48_})) | (d_2133_v49_))
            rhs375_ = (p2 if p0 else (p2) - (p2))
            rhs376_ = p0
            rhs377_ = p0
            lhs290_ = globalState
            lhs291_ = globalState
            lhs292_ = d_2131_v48_
            lhs293_ = default__.safeIndex(334, (d_2131_v48_).length(0))
            lhs294_ = d_2131_v48_
            lhs295_ = default__.safeIndex(283, (d_2131_v48_).length(0))
            lhs290_.f8 = rhs374_
            lhs291_.f3 = rhs375_
            lhs292_[lhs293_] = rhs376_
            lhs294_[lhs295_] = rhs377_
            d_2134_v50_: _dafny.Array
            def lambda109_(d_2135_p2_):
                def lambda110_(d_2136_i6_):
                    return default__.safeModuloInt(d_2136_i6_, d_2135_p2_)

                return lambda110_

            init59_ = lambda109_(p2)
            nw349_ = _dafny.Array(None, 5)
            for i0_59_ in range(nw349_.length(0)):
                nw349_[i0_59_] = init59_(i0_59_)
            d_2134_v50_ = nw349_
            d_2137_v51_: _dafny.Array
            d_2137_v51_ = d_2134_v50_
            d_2138_v52_: _dafny.Array
            nw350_ = _dafny.Array(None, 9)
            nw350_[int(0)] = d_2137_v51_
            nw350_[int(1)] = d_2137_v51_
            nw350_[int(2)] = (d_2134_v50_ if p0 else d_2137_v51_)
            nw350_[int(3)] = d_2137_v51_
            nw350_[int(4)] = d_2134_v50_
            nw350_[int(5)] = d_2134_v50_
            nw350_[int(6)] = d_2137_v51_
            nw350_[int(7)] = d_2137_v51_
            nw350_[int(8)] = d_2137_v51_
            d_2138_v52_ = nw350_
            index380_ = default__.safeIndex(306, (d_2138_v52_).length(0))
            (d_2138_v52_)[index380_] = d_2137_v51_
            index381_ = default__.safeIndex(306, (d_2138_v52_).length(0))
            (d_2138_v52_)[index381_] = d_2134_v50_
            d_2139_v53_: _dafny.Map
            d_2139_v53_ = _dafny.Map({p2: p2})
            d_2139_v53_ = (d_2139_v53_).set(len(d_2128_v45_), p2)
            d_2140_v54_: _dafny.Set
            d_2140_v54_ = _dafny.Set({-799, p2})
            d_2141_v55_: D3
            d_2141_v55_ = D3_DC9()
            d_2142_v56_: _dafny.Set
            d_2142_v56_ = _dafny.Set({default__.fm50(globalState), d_2141_v55_, d_2141_v55_})
            d_2143_v57_: _dafny.Seq
            d_2143_v57_ = _dafny.SeqWithoutIsStrInference([d_2142_v56_, d_2142_v56_])
            (globalState).f8 = (len((d_2140_v54_) | (d_2140_v54_))) - ((len((d_2143_v57_)[default__.safeIndex(p2, len(d_2143_v57_))]) if p0 else len(d_2128_v45_)))
            d_2144_v58_: _dafny.Array
            nw351_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_2144_v58_ = nw351_
            index382_ = default__.safeIndex(384, (d_2144_v58_).length(0))
            (d_2144_v58_)[index382_] = (d_2129_v46_) + (d_2129_v46_)
            index383_ = default__.safeIndex(384, (d_2144_v58_).length(0))
            (d_2144_v58_)[index383_] = (d_2129_v46_).set(default__.safeIndex(p2, len(d_2129_v46_)), d_2130_v47_)
        elif True:
            d_2145_v59_: _dafny.Map
            d_2145_v59_ = _dafny.Map({d_2129_v46_: p2})
            d_2146_v60_: _dafny.Map
            d_2146_v60_ = _dafny.Map({p0: p2})
            d_2145_v59_ = (d_2145_v59_).set(d_2129_v46_, (0) - (((d_2146_v60_)[p0] if (p0) in (d_2146_v60_) else (0) - (len((_dafny.SeqWithoutIsStrInference([p2, p2, len(_dafny.SeqWithoutIsStrInference([d_2130_v47_ for d_2147_i7_ in range(default__.abs(800))]))])).set(default__.safeIndex(len(d_2129_v46_), len(_dafny.SeqWithoutIsStrInference([p2, p2, len(_dafny.SeqWithoutIsStrInference([d_2130_v47_ for d_2148_i7_ in range(default__.abs(800))]))]))), p2))))))
            d_2149_v61_: T0
            nw352_ = C2()
            nw352_.ctor__()
            d_2149_v61_ = nw352_
            d_2150_v62_: _dafny.Seq
            d_2150_v62_ = _dafny.SeqWithoutIsStrInference([d_2149_v61_])
            d_2149_v61_ = (d_2150_v62_)[default__.safeIndex(p2, len(d_2150_v62_))]
            d_2129_v46_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvyfq"))) + (d_2129_v46_)) + (d_2129_v46_)
            d_2151_v63_: C3
            nw353_ = C3()
            nw353_.ctor__()
            d_2151_v63_ = nw353_
            d_2152_v64_: _dafny.Array
            nw354_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
            d_2152_v64_ = nw354_
            index384_ = default__.safeIndex(474, (d_2152_v64_).length(0))
            (d_2152_v64_)[index384_] = d_2129_v46_
            index385_ = default__.safeIndex(474, (d_2152_v64_).length(0))
            (d_2152_v64_)[index385_] = d_2129_v46_
        d_2153_v65_: _dafny.Array
        nw355_ = _dafny.Array(_dafny.Set({}), 26)
        d_2153_v65_ = nw355_
        d_2154_v66_: _dafny.Array
        nw356_ = _dafny.Array(None, 13)
        nw356_[int(0)] = d_2130_v47_
        nw356_[int(1)] = d_2130_v47_
        nw356_[int(2)] = d_2130_v47_
        nw356_[int(3)] = d_2130_v47_
        nw356_[int(4)] = d_2130_v47_
        nw356_[int(5)] = d_2130_v47_
        nw356_[int(6)] = d_2130_v47_
        nw356_[int(7)] = d_2130_v47_
        nw356_[int(8)] = d_2130_v47_
        nw356_[int(9)] = d_2130_v47_
        nw356_[int(10)] = d_2130_v47_
        nw356_[int(11)] = d_2130_v47_
        nw356_[int(12)] = d_2130_v47_
        d_2154_v66_ = nw356_
        d_2155_v67_: _dafny.Set
        d_2155_v67_ = _dafny.Set({d_2154_v66_, d_2154_v66_})
        index386_ = default__.safeIndex(546, (d_2153_v65_).length(0))
        (d_2153_v65_)[index386_] = d_2155_v67_
        d_2156_v68_: D19
        d_2156_v68_ = D19_DC39(p0, 762)
        d_2157_v69_: _dafny.Array
        nw357_ = _dafny.Array(None, 6)
        nw357_[int(0)] = not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxey"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brbh"))))
        nw357_[int(1)] = p0
        nw357_[int(2)] = (d_2156_v68_).cf58
        nw357_[int(3)] = p0
        nw357_[int(4)] = (default__.fm14(p2, globalState)) != (d_2129_v46_)
        nw357_[int(5)] = False
        d_2157_v69_ = nw357_
        index387_ = default__.safeIndex(712, (d_2157_v69_).length(0))
        (d_2157_v69_)[index387_] = p0
        d_2158_v70_: _dafny.Map
        d_2158_v70_ = _dafny.Map({p2: d_2155_v67_})
        index388_ = default__.safeIndex(546, (d_2153_v65_).length(0))
        index389_ = default__.safeIndex(712, (d_2157_v69_).length(0))
        rhs378_ = (d_2155_v67_) - (((d_2158_v70_)[p2] if (p2) in (d_2158_v70_) else d_2155_v67_))
        rhs379_ = p0
        rhs380_ = (default__.fm51(p2, globalState))
        lhs296_ = d_2153_v65_
        lhs297_ = default__.safeIndex(546, (d_2153_v65_).length(0))
        lhs298_ = globalState
        lhs299_ = d_2157_v69_
        lhs300_ = default__.safeIndex(712, (d_2157_v69_).length(0))
        lhs296_[lhs297_] = rhs378_
        lhs298_.f4 = rhs379_
        lhs299_[lhs300_] = rhs380_


class C9:
    def  __init__(self):
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self, f9):
        (self)._f9 = f9

    def m2(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        source38_ = default__.fm5(globalState)
        if source38_.is_DC1:
            d_2159___mcc_h0_ = source38_.cf1
            d_2160___mcc_h1_ = source38_.cf2
            d_2161_cf2_ = d_2160___mcc_h1_
            d_2162_cf1_ = d_2159___mcc_h0_
            d_2163_v0_: _dafny.Array
            nw358_ = _dafny.Array(False, 10)
            d_2163_v0_ = nw358_
            d_2163_v0_ = d_2163_v0_
            d_2164_v1_: str
            d_2164_v1_ = _dafny.CodePoint('h')
            d_2165_v2_: _dafny.Map
            d_2165_v2_ = _dafny.Map({(self).f9: d_2164_v1_})
            d_2166_v3_: _dafny.Map
            d_2166_v3_ = _dafny.Map({(d_2162_cf1_) + (d_2162_cf1_): d_2165_v2_})
            d_2167_v4_: _dafny.Map
            d_2167_v4_ = _dafny.Map({(self).f9: d_2162_cf1_})
            d_2168_v5_: _dafny.Seq
            d_2168_v5_ = _dafny.SeqWithoutIsStrInference([(self).f9])
            d_2166_v3_ = _dafny.Map({default__.fm6(((d_2167_v4_)[(self).f9] if ((self).f9) in (d_2167_v4_) else len(d_2168_v5_)), globalState): default__.fm7(d_2164_v1_, globalState)})
            r1 = d_2161_cf2_
            if (self).f9:
                d_2169_v6_: C7
                nw359_ = C7()
                nw359_.ctor__()
                d_2169_v6_ = nw359_
                d_2170_v7_: _dafny.Array
                nw360_ = _dafny.Array(None, 23)
                d_2170_v7_ = nw360_
                d_2171_v8_: C0
                nw361_ = C0()
                nw361_.ctor__()
                d_2171_v8_ = nw361_
                index390_ = default__.safeIndex(100, (d_2170_v7_).length(0))
                (d_2170_v7_)[index390_] = d_2171_v8_
                index391_ = default__.safeIndex(100, (d_2170_v7_).length(0))
                (d_2170_v7_)[index391_] = d_2171_v8_
                d_2172_v9_: _dafny.Seq
                d_2172_v9_ = _dafny.SeqWithoutIsStrInference([d_2162_cf1_])
                (globalState).f4 = (d_2169_v6_).fm3(d_2172_v9_, _dafny.MultiSet(d_2172_v9_), (self).f9, globalState)
                d_2173_v10_: C8
                nw362_ = C8()
                nw362_.ctor__()
                d_2173_v10_ = nw362_
                d_2174_v11_: D2
                d_2174_v11_ = D2_DC5(d_2163_v0_)
                d_2175_v12_: _dafny.Array
                nw363_ = _dafny.Array(None, 11)
                nw363_[int(0)] = d_2163_v0_
                nw363_[int(1)] = d_2163_v0_
                nw363_[int(2)] = d_2163_v0_
                nw363_[int(3)] = d_2163_v0_
                nw363_[int(4)] = d_2163_v0_
                nw363_[int(5)] = d_2163_v0_
                nw363_[int(6)] = d_2163_v0_
                nw363_[int(7)] = d_2163_v0_
                nw363_[int(8)] = d_2163_v0_
                nw363_[int(9)] = (d_2174_v11_).cf5
                nw363_[int(10)] = d_2163_v0_
                d_2175_v12_ = nw363_
                index392_ = default__.safeIndex(812, (d_2175_v12_).length(0))
                (d_2175_v12_)[index392_] = d_2163_v0_
                index393_ = default__.safeIndex(812, (d_2175_v12_).length(0))
                (d_2175_v12_)[index393_] = d_2163_v0_
            elif True:
                (globalState).f8 = default__.safeModuloInt(142, (0) - (d_2161_cf2_))
                d_2176_v13_: _dafny.Array
                def lambda111_(d_2177_cf1_):
                    def lambda112_(d_2178_i0_):
                        return (d_2178_i0_) * (d_2177_cf1_)

                    return lambda112_

                init60_ = lambda111_(d_2162_cf1_)
                nw364_ = _dafny.Array(None, 23)
                for i0_60_ in range(nw364_.length(0)):
                    nw364_[i0_60_] = init60_(i0_60_)
                d_2176_v13_ = nw364_
                d_2179_v14_: _dafny.Map
                d_2179_v14_ = _dafny.Map({(self).f9: False})
                d_2180_v15_: _dafny.Map
                d_2180_v15_ = _dafny.Map({d_2176_v13_: d_2179_v14_})
                d_2181_v16_: _dafny.Map
                d_2181_v16_ = _dafny.Map({(d_2162_cf1_) - ((0) - ((0) - (len(d_2180_v15_)))): (self).f9})
                d_2182_v18_: _dafny.Set
                d_2182_v18_ = _dafny.Set({-25})
                def iife228_():
                    coll102_ = _dafny.Map()
                    compr_102_: int
                    for compr_102_ in (d_2182_v18_).Elements:
                        d_2183_v17_: int = compr_102_
                        if (d_2183_v17_) in (d_2182_v18_):
                            coll102_[(d_2183_v17_) - (d_2162_cf1_)] = (self).f9
                    return _dafny.Map(coll102_)
                d_2181_v16_ = (d_2181_v16_) | ((iife228_()
                ).set(d_2162_cf1_, (self).f9))
                d_2184_v19_: _dafny.Seq
                d_2184_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fml"))
                d_2185_v20_: C6
                nw365_ = C6()
                nw365_.ctor__(d_2161_cf2_, (d_2184_v19_) + (d_2184_v19_))
                d_2185_v20_ = nw365_
                d_2186_v21_: _dafny.Set
                d_2186_v21_ = _dafny.Set({(self).f9, (self).f9, (self).f9, (d_2168_v5_)[default__.safeIndex((d_2185_v20_).f10, len(d_2168_v5_))], (self).f9})
                (globalState).f3 = default__.safeDivisionInt(len(d_2186_v21_), (-741) * (default__.fm6(d_2161_cf2_, globalState)))
                d_2187_v22_: D12
                d_2187_v22_ = D12_DC26(_dafny.Map({((d_2179_v14_)[(self).f9] if ((self).f9) in (d_2179_v14_) else (self).f9): d_2161_cf2_}))
                d_2188_v23_: _dafny.Map
                d_2188_v23_ = _dafny.Map({d_2187_v22_: True})
                d_2189_v24_: _dafny.Map
                d_2189_v24_ = _dafny.Map({d_2161_cf2_: d_2188_v23_})
                rhs381_ = len((d_2189_v24_) | (_dafny.Map({len((d_2185_v20_).f11): d_2188_v23_})))
                rhs382_ = (self).f9
                lhs301_ = globalState
                lhs301_.f3 = rhs381_
                r0 = rhs382_
        elif source38_.is_DC2:
            d_2190_v25_: _dafny.Seq
            d_2190_v25_ = _dafny.SeqWithoutIsStrInference([(self).f9, not((self).f9)])
            d_2191_v26_: _dafny.Map
            d_2191_v26_ = _dafny.Map({True: len(d_2190_v25_)})
            d_2192_v27_: int
            d_2192_v27_ = -901
            d_2193_v29_: _dafny.Seq
            d_2193_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivrfg"))
            d_2194_v30_: _dafny.Seq
            d_2194_v30_ = _dafny.SeqWithoutIsStrInference([len(d_2193_v29_)])
            d_2195_v31_: _dafny.Array
            nw366_ = _dafny.Array(None, 18)
            nw366_[int(0)] = d_2191_v26_
            nw366_[int(1)] = d_2191_v26_
            nw366_[int(2)] = d_2191_v26_
            nw366_[int(3)] = d_2191_v26_
            nw366_[int(4)] = default__.fm46(globalState)
            nw366_[int(5)] = d_2191_v26_
            nw366_[int(6)] = _dafny.Map({False: d_2192_v27_})
            nw366_[int(7)] = d_2191_v26_
            nw366_[int(8)] = d_2191_v26_
            nw366_[int(9)] = d_2191_v26_
            nw366_[int(10)] = _dafny.Map({(self).f9: d_2192_v27_})
            def iife229_():
                coll103_ = _dafny.Map()
                compr_103_: int
                for compr_103_ in _dafny.IntegerRange(993, 539):
                    d_2196_v28_: int = compr_103_
                    if ((993) <= (d_2196_v28_)) and ((d_2196_v28_) < (539)):
                        coll103_[(d_2196_v28_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkmh"))))] = (self).f9
                return _dafny.Map(coll103_)
            nw366_[int(11)] = _dafny.Map({(self).f9: len((_dafny.Map({d_2192_v27_: (self).f9})).set(len(iife229_()
            ), (self).f9))})
            nw366_[int(12)] = d_2191_v26_
            nw366_[int(13)] = d_2191_v26_
            nw366_[int(14)] = d_2191_v26_
            nw366_[int(15)] = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([d_2192_v27_, (_dafny.MultiSet(d_2194_v30_)).cardinality, d_2192_v27_, d_2192_v27_]))})
            nw366_[int(16)] = d_2191_v26_
            nw366_[int(17)] = d_2191_v26_
            d_2195_v31_ = nw366_
            d_2197_v32_: C4
            nw367_ = C4()
            nw367_.ctor__(d_2195_v31_)
            d_2197_v32_ = nw367_
            d_2198_v33_: _dafny.MultiSet
            d_2198_v33_ = _dafny.MultiSet([(self).f9])
            d_2199_v34_: _dafny.Map
            d_2199_v34_ = _dafny.Map({(self).f9: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fal"))})
            d_2200_v35_: D4
            d_2200_v35_ = D4_DC11((417) < ((((d_2198_v33_).set(True, default__.abs(836))).set((self).f9, default__.abs(d_2192_v27_))).cardinality), ((0) - ((0) - (len(((d_2199_v34_)[(self).f9] if ((self).f9) in (d_2199_v34_) else d_2193_v29_))))) >= (d_2192_v27_), d_2192_v27_, d_2193_v29_)
            d_2201_v36_: _dafny.Map
            d_2201_v36_ = _dafny.Map({d_2190_v25_: (self).f9})
            d_2202_v37_: _dafny.Map
            d_2202_v37_ = _dafny.Map({(self).f9: (self).f9})
            d_2203_v38_: D12
            d_2203_v38_ = D12_DC27(16, d_2192_v27_, d_2192_v27_, len(d_2202_v37_))
            d_2204_v39_: _dafny.Array
            nw368_ = _dafny.Array(None, 20)
            nw368_[int(0)] = (self).f9
            nw368_[int(1)] = (self).f9
            nw368_[int(2)] = (self).f9
            nw368_[int(3)] = (self).f9
            nw368_[int(4)] = (self).f9
            nw368_[int(5)] = (self).f9
            nw368_[int(6)] = (self).f9
            nw368_[int(7)] = (self).f9
            nw368_[int(8)] = (self).f9
            nw368_[int(9)] = (self).f9
            nw368_[int(10)] = (self).f9
            nw368_[int(11)] = (self).f9
            nw368_[int(12)] = (self).f9
            nw368_[int(13)] = False
            nw368_[int(14)] = False
            nw368_[int(15)] = (self).f9
            nw368_[int(16)] = (self).f9
            nw368_[int(17)] = (self).f9
            nw368_[int(18)] = (self).f9
            nw368_[int(19)] = (self).f9
            d_2204_v39_ = nw368_
            d_2205_v40_: _dafny.MultiSet
            d_2205_v40_ = _dafny.MultiSet([d_2204_v39_, d_2204_v39_, d_2204_v39_])
            d_2206_v41_: _dafny.Map
            d_2206_v41_ = _dafny.Map({d_2192_v27_: d_2192_v27_})
            d_2207_v42_: D3
            d_2207_v42_ = D3_DC8(d_2206_v41_)
            d_2208_v43_: _dafny.MultiSet
            d_2208_v43_ = _dafny.MultiSet([d_2190_v25_])
            d_2209_v44_: C6
            nw369_ = C6()
            nw369_.ctor__(d_2192_v27_, ((d_2199_v34_)[(self).f9] if ((self).f9) in (d_2199_v34_) else _dafny.SeqWithoutIsStrInference([(d_2193_v29_)[default__.safeIndex(d_2192_v27_, len(d_2193_v29_))] for d_2210_i1_ in range(default__.abs(706))])))
            d_2209_v44_ = nw369_
            d_2211_v45_: _dafny.Map
            d_2211_v45_ = _dafny.Map({(self).f9: d_2209_v44_})
            d_2212_v46_: _dafny.Array
            nw370_ = _dafny.Array(None, 23)
            nw370_[int(0)] = (0) - (len(d_2201_v36_))
            nw370_[int(1)] = ((d_2203_v38_).cf39) - (d_2192_v27_)
            nw370_[int(2)] = d_2192_v27_
            nw370_[int(3)] = ((d_2205_v40_).set(d_2204_v39_, default__.abs(d_2192_v27_))).cardinality
            nw370_[int(4)] = (len(d_2190_v25_) if (self).f9 else default__.fm6(d_2192_v27_, globalState))
            nw370_[int(5)] = default__.safeDivisionInt(len(d_2206_v41_), d_2192_v27_)
            nw370_[int(6)] = d_2192_v27_
            nw370_[int(7)] = -926
            nw370_[int(8)] = d_2192_v27_
            nw370_[int(9)] = ((d_2198_v33_)[(self).f9] if ((self).f9) in (d_2198_v33_) else 835)
            nw370_[int(10)] = default__.safeModuloInt(d_2192_v27_, d_2192_v27_)
            nw370_[int(11)] = default__.safeDivisionInt(default__.fm26(d_2207_v42_, d_2192_v27_, d_2208_v43_, globalState), d_2192_v27_)
            nw370_[int(12)] = (0) - ((0) - (len(((d_2211_v45_).set((self).f9, d_2209_v44_)).set((D9_DC21((self).f9, d_2190_v25_, True)).cf26, d_2209_v44_))))
            nw370_[int(13)] = d_2192_v27_
            nw370_[int(14)] = (d_2209_v44_).f10
            nw370_[int(15)] = (0) - (d_2192_v27_)
            nw370_[int(16)] = (0) - (d_2192_v27_)
            nw370_[int(17)] = (d_2192_v27_) - (52)
            nw370_[int(18)] = d_2192_v27_
            nw370_[int(19)] = d_2192_v27_
            nw370_[int(20)] = (0) - (d_2192_v27_)
            nw370_[int(21)] = len(_dafny.SeqWithoutIsStrInference([d_2209_v44_]))
            nw370_[int(22)] = d_2192_v27_
            d_2212_v46_ = nw370_
            d_2213_v47_: _dafny.Seq
            d_2213_v47_ = _dafny.SeqWithoutIsStrInference([d_2212_v46_, d_2212_v46_, d_2212_v46_, d_2212_v46_])
            d_2214_v48_: _dafny.Seq
            d_2214_v48_ = _dafny.SeqWithoutIsStrInference([d_2212_v46_, d_2212_v46_, d_2212_v46_, (d_2213_v47_)[default__.safeIndex((d_2209_v44_).f10, len(d_2213_v47_))]])
            rhs383_ = D4_DC11((self).f9, True, -501, ((d_2209_v44_).f11) + ((d_2209_v44_).f11))
            rhs384_ = (d_2213_v47_)[default__.safeIndex(len(d_2190_v25_), len(d_2213_v47_))]
            rhs385_ = (self).f9
            d_2200_v35_ = rhs383_
            d_2212_v46_ = rhs384_
            r0 = rhs385_
            d_2215_v49_: C4
            nw371_ = C4()
            nw371_.ctor__((d_2197_v32_).f12)
            d_2215_v49_ = nw371_
            r1 = default__.fm20((self).f9, globalState)
        elif source38_.is_DC0:
            d_2216___mcc_h2_ = source38_.cf0
            d_2217_cf0_ = d_2216___mcc_h2_
            d_2218_v50_: _dafny.Map
            d_2218_v50_ = _dafny.Map({(self).f9: (d_2217_cf0_) * (944)})
            d_2219_v51_: _dafny.Seq
            d_2219_v51_ = _dafny.SeqWithoutIsStrInference([d_2217_cf0_])
            d_2218_v50_ = (d_2218_v50_).set((self).f9, (d_2217_cf0_) + (len(d_2219_v51_)))
            (globalState).f8 = d_2217_cf0_
            (globalState).f4 = (d_2217_cf0_) <= ((0) - (d_2217_cf0_))
            d_2220_v52_: _dafny.Array
            def lambda113_(d_2221_i2_):
                return (self).f9

            init61_ = lambda113_
            nw372_ = _dafny.Array(None, 10)
            for i0_61_ in range(nw372_.length(0)):
                nw372_[i0_61_] = init61_(i0_61_)
            d_2220_v52_ = nw372_
            d_2222_v53_: _dafny.Seq
            d_2222_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwsx"))
            d_2223_v54_: _dafny.Map
            d_2223_v54_ = _dafny.Map({len(d_2222_v53_): d_2220_v52_})
            d_2224_v55_: _dafny.MultiSet
            d_2224_v55_ = _dafny.MultiSet([d_2217_cf0_, len(_dafny.SeqWithoutIsStrInference([-984, -528])), d_2217_cf0_])
            d_2225_v56_: _dafny.Array
            nw373_ = _dafny.Array(None, 14)
            nw373_[int(0)] = d_2220_v52_
            nw373_[int(1)] = d_2220_v52_
            nw373_[int(2)] = ((d_2223_v54_)[(d_2224_v55_).cardinality] if ((d_2224_v55_).cardinality) in (d_2223_v54_) else d_2220_v52_)
            nw373_[int(3)] = d_2220_v52_
            nw373_[int(4)] = d_2220_v52_
            nw373_[int(5)] = d_2220_v52_
            nw373_[int(6)] = d_2220_v52_
            nw373_[int(7)] = d_2220_v52_
            nw373_[int(8)] = d_2220_v52_
            nw373_[int(9)] = d_2220_v52_
            nw373_[int(10)] = (d_2220_v52_ if True else d_2220_v52_)
            nw373_[int(11)] = (D2_DC5(d_2220_v52_)).cf5
            nw373_[int(12)] = d_2220_v52_
            nw373_[int(13)] = d_2220_v52_
            d_2225_v56_ = nw373_
            index394_ = default__.safeIndex(707, (d_2225_v56_).length(0))
            (d_2225_v56_)[index394_] = d_2220_v52_
            index395_ = default__.safeIndex(707, (d_2225_v56_).length(0))
            (d_2225_v56_)[index395_] = d_2220_v52_
        elif True:
            d_2226___mcc_h3_ = source38_.cf3
            d_2227_cf3_ = d_2226___mcc_h3_
            if (self).f9:
                d_2228_v57_: _dafny.Array
                nw374_ = _dafny.Array(_dafny.Seq({}), 1)
                d_2228_v57_ = nw374_
                d_2228_v57_ = d_2228_v57_
                d_2229_v58_: _dafny.Seq
                d_2229_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljehk"))
                d_2230_v59_: _dafny.Seq
                d_2230_v59_ = _dafny.SeqWithoutIsStrInference([d_2229_v58_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbuyasjcd"))])
                d_2231_v60_: _dafny.Seq
                d_2231_v60_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2229_v58_, d_2229_v58_, d_2229_v58_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2232_i3_ in range(default__.abs(677))])])) != (d_2230_v59_)])
                d_2231_v60_ = d_2231_v60_
                d_2233_v61_: C5
                nw375_ = C5()
                nw375_.ctor__()
                d_2233_v61_ = nw375_
                d_2229_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrlflpdvu"))
                d_2234_v62_: str
                d_2234_v62_ = _dafny.CodePoint('m')
                d_2235_v63_: _dafny.Map
                d_2235_v63_ = _dafny.Map({(self).f9: d_2234_v62_})
                d_2236_v64_: _dafny.Map
                d_2236_v64_ = _dafny.Map({d_2235_v63_: d_2229_v58_})
                d_2236_v64_ = (d_2236_v64_).set((d_2235_v63_) | (d_2235_v63_), d_2229_v58_)
            elif True:
                d_2237_v65_: _dafny.Seq
                d_2237_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kypi"))
                d_2238_v66_: _dafny.Map
                d_2238_v66_ = _dafny.Map({len(d_2237_v65_): False})
                d_2239_v67_: int
                d_2239_v67_ = 831
                d_2238_v66_ = (d_2238_v66_).set(default__.fm6(d_2239_v67_, globalState), False)
                (globalState).f3 = (0) - (d_2239_v67_)
                (globalState).f3 = d_2239_v67_
                d_2240_v68_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.MultiSet({}), 24)
                d_2240_v68_ = nw376_
                d_2241_v69_: _dafny.Seq
                d_2241_v69_ = _dafny.SeqWithoutIsStrInference([(self).f9])
                index396_ = default__.safeIndex(716, (d_2240_v68_).length(0))
                (d_2240_v68_)[index396_] = _dafny.MultiSet(d_2241_v69_)
                index397_ = default__.safeIndex(716, (d_2240_v68_).length(0))
                (d_2240_v68_)[index397_] = (_dafny.MultiSet([(self).f9, (self).f9, not(not((self).f9))])).set((self).f9, default__.abs(d_2239_v67_))
                d_2242_v70_: _dafny.Set
                d_2242_v70_ = _dafny.Set({not((self).f9)})
                r2 = _dafny.Set({d_2239_v67_, default__.safeModuloInt(d_2239_v67_, d_2239_v67_), len((_dafny.Set({True, (self).f9, (self).f9, (self).f9, (self).f9})).intersection(d_2242_v70_)), (0) - ((d_2239_v67_) * (293)), d_2239_v67_})
            d_2243_v71_: _dafny.Seq
            d_2243_v71_ = _dafny.SeqWithoutIsStrInference([390])
            d_2244_v72_: int
            d_2244_v72_ = 603
            d_2245_v73_: _dafny.Set
            d_2245_v73_ = _dafny.Set({(d_2243_v71_) + (d_2243_v71_), (_dafny.SeqWithoutIsStrInference([-895 for d_2246_i4_ in range(default__.abs(395))])).set(default__.safeIndex(d_2244_v72_, len(_dafny.SeqWithoutIsStrInference([-895 for d_2247_i4_ in range(default__.abs(395))]))), 206), d_2243_v71_})
            d_2245_v73_ = d_2245_v73_
            d_2248_v74_: str
            d_2248_v74_ = _dafny.CodePoint('f')
            d_2249_v75_: _dafny.Map
            d_2249_v75_ = _dafny.Map({d_2248_v74_: d_2248_v74_})
            d_2250_v76_: _dafny.Array
            nw377_ = _dafny.Array(int(0), 18)
            d_2250_v76_ = nw377_
            index398_ = default__.safeIndex(282, (d_2250_v76_).length(0))
            (d_2250_v76_)[index398_] = (0) - (len(_dafny.Map({d_2250_v76_: (self).f9})))
            d_2251_v77_: _dafny.Array
            def lambda114_(d_2252_i5_):
                return (self).f9

            init62_ = lambda114_
            nw378_ = _dafny.Array(None, 3)
            for i0_62_ in range(nw378_.length(0)):
                nw378_[i0_62_] = init62_(i0_62_)
            d_2251_v77_ = nw378_
            d_2253_v78_: _dafny.Map
            d_2253_v78_ = _dafny.Map({d_2251_v77_: d_2249_v75_})
            d_2254_v79_: _dafny.Map
            d_2254_v79_ = _dafny.Map({d_2244_v72_: d_2244_v72_})
            d_2255_v80_: _dafny.Map
            d_2255_v80_ = _dafny.Map({len(d_2254_v79_): True})
            d_2256_v81_: _dafny.MultiSet
            d_2256_v81_ = _dafny.MultiSet([(self).f9])
            index399_ = default__.safeIndex(282, (d_2250_v76_).length(0))
            rhs386_ = ((d_2253_v78_)[d_2251_v77_] if (d_2251_v77_) in (d_2253_v78_) else _dafny.Map({d_2248_v74_: default__.fm18(((d_2255_v80_)[d_2244_v72_] if (d_2244_v72_) in (d_2255_v80_) else (self).f9), (self).f9, globalState)}))
            rhs387_ = ((d_2244_v72_) * ((d_2256_v81_).cardinality) if (self).f9 else (d_2243_v71_)[default__.safeIndex(371, len(d_2243_v71_))])
            lhs302_ = d_2250_v76_
            lhs303_ = default__.safeIndex(282, (d_2250_v76_).length(0))
            d_2249_v75_ = rhs386_
            lhs302_[lhs303_] = rhs387_
            d_2257_v82_: _dafny.Seq
            d_2257_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktuiokx"))
            d_2257_v82_ = default__.fm25(globalState)
        d_2258_v83_: _dafny.Array
        nw379_ = _dafny.Array(int(0), 27)
        d_2258_v83_ = nw379_
        index400_ = default__.safeIndex(232, (d_2258_v83_).length(0))
        (d_2258_v83_)[index400_] = default__.fm20((self).f9, globalState)
        d_2259_v84_: int
        d_2259_v84_ = 382
        d_2260_v85_: _dafny.Seq
        d_2260_v85_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_2261_v86_: _dafny.Seq
        d_2261_v86_ = _dafny.SeqWithoutIsStrInference([d_2259_v84_, d_2259_v84_, d_2259_v84_, len(d_2260_v85_)])
        d_2262_v87_: _dafny.Map
        d_2262_v87_ = _dafny.Map({(self).f9: (self).f9})
        index401_ = default__.safeIndex(232, (d_2258_v83_).length(0))
        (d_2258_v83_)[index401_] = default__.fm6(default__.safeModuloInt(d_2259_v84_, (_dafny.MultiSet((d_2261_v86_).set(default__.safeIndex(default__.fm6(len(d_2262_v87_), globalState), len(d_2261_v86_)), d_2259_v84_))).cardinality), globalState)
        d_2263_v88_: _dafny.Seq
        d_2263_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojpl"))
        d_2264_v89_: D4
        d_2264_v89_ = D4_DC11((self).f9, (self).f9, len(_dafny.SeqWithoutIsStrInference([d_2259_v84_, (d_2258_v83_)[default__.safeIndex(232, (d_2258_v83_).length(0))]])), d_2263_v88_)
        hi14_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2265_i7_ in range(default__.abs(722))]))
        for d_2266_i6_ in range((d_2264_v89_).cf16, hi14_):
            d_2267_v90_: D4
            d_2267_v90_ = D4_DC10(_dafny.SeqWithoutIsStrInference([(self).f9]))
            d_2268_v91_: _dafny.Map
            d_2268_v91_ = _dafny.Map({default__.fm49((self).f9, globalState): (_dafny.SeqWithoutIsStrInference([(self).f9])) + ((d_2267_v90_).cf13)})
            r1 = len(d_2268_v91_)
            r0 = (self).f9
            (globalState).f8 = d_2266_i6_
            d_2263_v88_ = d_2263_v88_
        index402_ = default__.safeIndex(232, (d_2258_v83_).length(0))
        rhs388_ = ((d_2258_v83_)[default__.safeIndex(232, (d_2258_v83_).length(0))]) + ((d_2258_v83_)[default__.safeIndex(232, (d_2258_v83_).length(0))])
        rhs389_ = len(((d_2260_v85_) + ((d_2260_v85_).set(default__.safeIndex(421, len(d_2260_v85_)), (self).f9))) + (_dafny.SeqWithoutIsStrInference([(self).f9])))
        lhs304_ = globalState
        lhs305_ = d_2258_v83_
        lhs306_ = default__.safeIndex(232, (d_2258_v83_).length(0))
        lhs304_.f8 = rhs388_
        lhs305_[lhs306_] = rhs389_
        d_2269_v92_: T0
        nw380_ = C7()
        nw380_.ctor__()
        d_2269_v92_ = nw380_
        d_2270_v93_: _dafny.Seq
        d_2270_v93_ = _dafny.SeqWithoutIsStrInference([d_2269_v92_])
        d_2271_v94_: _dafny.Seq
        d_2271_v94_ = d_2270_v93_
        d_2271_v94_ = d_2270_v93_
        (globalState).f3 = (d_2258_v83_)[default__.safeIndex(232, (d_2258_v83_).length(0))]
        d_2272_v95_: D2
        d_2272_v95_ = D2_DC6((d_2258_v83_)[default__.safeIndex(232, (d_2258_v83_).length(0))], (self).f9, (self).f9)
        pat_let_tv71_ = d_2258_v83_
        pat_let_tv72_ = d_2258_v83_
        def lambda115_(source39_):
            if source39_.is_DC6:
                d_2273___mcc_h4_ = source39_.cf6
                d_2274___mcc_h5_ = source39_.cf7
                d_2275___mcc_h6_ = source39_.cf8
                d_2276_cf8_ = d_2275___mcc_h6_
                d_2277_cf7_ = d_2274___mcc_h5_
                d_2278_cf6_ = d_2273___mcc_h4_
                if (self).f9:
                    return d_2276_cf8_
                elif True:
                    return not(d_2277_cf7_)
            elif source39_.is_DC7:
                d_2279___mcc_h7_ = source39_.cf9
                d_2280___mcc_h8_ = source39_.cf10
                d_2281___mcc_h9_ = source39_.cf11
                d_2282_cf11_ = d_2281___mcc_h9_
                d_2283_cf10_ = d_2280___mcc_h8_
                d_2284_cf9_ = d_2279___mcc_h7_
                return (self).f9
            elif True:
                d_2285___mcc_h10_ = source39_.cf5
                d_2286_cf5_ = d_2285___mcc_h10_
                return (_dafny.MultiSet([529])).isdisjoint(_dafny.MultiSet([(pat_let_tv72_)[default__.safeIndex(232, (pat_let_tv71_).length(0))]]))

        r0 = lambda115_(d_2272_v95_)
        r1 = d_2259_v84_
        d_2287_v96_: _dafny.Set
        d_2287_v96_ = _dafny.Set({d_2259_v84_})
        r2 = d_2287_v96_
        return r0, r1, r2

    def m3(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_2288_v0_: _dafny.Seq
        d_2288_v0_ = _dafny.SeqWithoutIsStrInference([(self).f9])
        d_2289_v1_: _dafny.Seq
        d_2289_v1_ = _dafny.SeqWithoutIsStrInference([d_2288_v0_, _dafny.SeqWithoutIsStrInference([False]), d_2288_v0_, d_2288_v0_, d_2288_v0_])
        d_2290_v2_: _dafny.Map
        d_2290_v2_ = _dafny.Map({d_2288_v0_: (0) - (len((d_2289_v1_)[default__.safeIndex(p1, len(d_2289_v1_))]))})
        d_2291_v3_: _dafny.Seq
        d_2291_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "julgoho"))
        d_2292_v4_: str
        d_2292_v4_ = _dafny.CodePoint('o')
        d_2293_v5_: _dafny.MultiSet
        d_2293_v5_ = _dafny.MultiSet([(((d_2290_v2_)[d_2288_v0_] if (d_2288_v0_) in (d_2290_v2_) else p0)) + (p0), -539, (len(d_2291_v3_) if (self).f9 else (0) - (len(default__.fm17(d_2292_v4_, globalState)))), (0) - (p0), p0])
        (globalState).f3 = ((d_2293_v5_)[p0] if (p0) in (d_2293_v5_) else p0)
        d_2294_v6_: _dafny.Set
        d_2294_v6_ = _dafny.Set({p0})
        if not ((self).f9) or ((d_2294_v6_).isdisjoint(d_2294_v6_)):
            d_2295_v7_: _dafny.Seq
            d_2295_v7_ = _dafny.SeqWithoutIsStrInference([989, p0])
            d_2296_v8_: _dafny.Map
            d_2296_v8_ = _dafny.Map({p0: 841})
            d_2297_v9_: D3
            d_2297_v9_ = D3_DC8(d_2296_v8_)
            d_2298_v10_: _dafny.MultiSet
            d_2298_v10_ = _dafny.MultiSet([d_2288_v0_])
            (globalState).f0 = ((d_2295_v7_) + (d_2295_v7_)).set(default__.safeIndex((p1) * (p1), len((d_2295_v7_) + (d_2295_v7_))), default__.fm26(d_2297_v9_, 687, d_2298_v10_, globalState))
            d_2299_v11_: _dafny.Array
            nw381_ = _dafny.Array(int(0), 17)
            d_2299_v11_ = nw381_
            d_2300_v12_: _dafny.Set
            d_2300_v12_ = _dafny.Set({d_2299_v11_, d_2299_v11_})
            d_2301_v13_: _dafny.Array
            nw382_ = _dafny.Array(_dafny.Map({}), 19)
            d_2301_v13_ = nw382_
            d_2302_v14_: T0
            nw383_ = C4()
            nw383_.ctor__(d_2301_v13_)
            d_2302_v14_ = nw383_
            d_2303_v15_: _dafny.Seq
            d_2303_v15_ = _dafny.SeqWithoutIsStrInference([d_2299_v11_])
            rhs390_ = (self).f9
            rhs391_ = (self).f9
            rhs392_ = _dafny.Set({d_2299_v11_, (d_2299_v11_ if (self).f9 else d_2299_v11_), (d_2303_v15_)[default__.safeIndex(p1, len(d_2303_v15_))]})
            rhs393_ = (0) - (-777)
            rhs394_ = d_2302_v14_
            lhs307_ = globalState
            lhs308_ = globalState
            r3 = rhs390_
            lhs307_.f4 = rhs391_
            d_2300_v12_ = rhs392_
            lhs308_.f3 = rhs393_
            d_2302_v14_ = rhs394_
            if ((p0) + (p0)) < (p0):
                (globalState).f3 = default__.safeModuloInt(p1, (default__.fm6(default__.fm20((self).f9, globalState), globalState)) - (p0))
                d_2304_v16_: _dafny.Set
                d_2304_v16_ = _dafny.Set({d_2292_v4_, default__.fm18((self).f9, (self).f9, globalState), default__.fm18((self).f9, False, globalState), d_2292_v4_})
                d_2304_v16_ = (d_2304_v16_).intersection((_dafny.Set({d_2292_v4_})) - (_dafny.Set({(D16_DC35(p0, d_2292_v4_, len((_dafny.SeqWithoutIsStrInference([d_2292_v4_ for d_2305_i0_ in range(default__.abs(-613))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_2292_v4_ for d_2306_i0_ in range(default__.abs(-613))]))), d_2292_v4_)), p0, p1)).cf51})))
                (globalState).f4 = ((self).f9 if (self).f9 else (self).f9)
                d_2307_v17_: _dafny.Array
                def lambda116_(d_2308_i1_):
                    return (self).f9

                init63_ = lambda116_
                nw384_ = _dafny.Array(None, 22)
                for i0_63_ in range(nw384_.length(0)):
                    nw384_[i0_63_] = init63_(i0_63_)
                d_2307_v17_ = nw384_
                d_2309_v18_: _dafny.Seq
                d_2309_v18_ = _dafny.SeqWithoutIsStrInference([d_2307_v17_, d_2307_v17_, d_2307_v17_])
                d_2309_v18_ = (d_2309_v18_ if (self).f9 else d_2309_v18_)
                d_2288_v0_ = (((d_2288_v0_).set(default__.safeIndex(p0, len(d_2288_v0_)), (self).f9)).set(default__.safeIndex(p0, len((d_2288_v0_).set(default__.safeIndex(p0, len(d_2288_v0_)), (self).f9))), (self).f9)).set(default__.safeIndex(default__.safeModuloInt(default__.fm20((self).f9, globalState), p0), len(((d_2288_v0_).set(default__.safeIndex(p0, len(d_2288_v0_)), (self).f9)).set(default__.safeIndex(p0, len((d_2288_v0_).set(default__.safeIndex(p0, len(d_2288_v0_)), (self).f9))), (self).f9))), (self).f9)
            elif True:
                d_2310_v19_: _dafny.Array
                nw385_ = _dafny.Array(_dafny.MultiSet({}), 11)
                d_2310_v19_ = nw385_
                d_2311_v20_: D21
                d_2311_v20_ = D21_DC43(d_2295_v7_, (self).f9)
                d_2312_v21_: _dafny.Array
                nw386_ = _dafny.Array(None, 21)
                nw386_[int(0)] = (self).f9
                nw386_[int(1)] = (d_2302_v14_).fm3((d_2311_v20_).cf63, d_2293_v5_, (self).f9, globalState)
                nw386_[int(2)] = False
                nw386_[int(3)] = (self).f9
                nw386_[int(4)] = (self).f9
                nw386_[int(5)] = (self).f9
                nw386_[int(6)] = (self).f9
                nw386_[int(7)] = (self).f9
                nw386_[int(8)] = (self).f9
                nw386_[int(9)] = (self).f9
                nw386_[int(10)] = False
                nw386_[int(11)] = (self).f9
                nw386_[int(12)] = (self).f9
                nw386_[int(13)] = (self).f9
                nw386_[int(14)] = (self).f9
                nw386_[int(15)] = (self).f9
                nw386_[int(16)] = (self).f9
                nw386_[int(17)] = (self).f9
                nw386_[int(18)] = (self).f9
                nw386_[int(19)] = (self).f9
                nw386_[int(20)] = (self).f9
                d_2312_v21_ = nw386_
                d_2313_v22_: _dafny.MultiSet
                d_2313_v22_ = _dafny.MultiSet([d_2312_v21_])
                index403_ = default__.safeIndex(722, (d_2310_v19_).length(0))
                (d_2310_v19_)[index403_] = d_2313_v22_
                index404_ = default__.safeIndex(722, (d_2310_v19_).length(0))
                (d_2310_v19_)[index404_] = _dafny.MultiSet([d_2312_v21_, d_2312_v21_, d_2312_v21_])
                d_2314_v23_: _dafny.Map
                d_2314_v23_ = _dafny.Map({d_2288_v0_: d_2295_v7_})
                d_2315_v24_: _dafny.Seq
                d_2315_v24_ = _dafny.SeqWithoutIsStrInference([d_2314_v23_])
                d_2314_v23_ = (d_2315_v24_)[default__.safeIndex(-356, len(d_2315_v24_))]
                index405_ = default__.safeIndex(664, (d_2312_v21_).length(0))
                (d_2312_v21_)[index405_] = default__.fm49(default__.fm49((self).f9, globalState), globalState)
                index406_ = default__.safeIndex(664, (d_2312_v21_).length(0))
                (d_2312_v21_)[index406_] = ((self).f9 if False else (self).f9)
                (globalState).f3 = p1
                (globalState).f4 = (self).f9
            d_2316_v25_: _dafny.Set
            d_2316_v25_ = _dafny.Set({(self).f9, (self).f9})
            d_2317_v26_: _dafny.Map
            d_2317_v26_ = _dafny.Map({p0: (d_2316_v25_).isdisjoint(d_2316_v25_)})
            d_2317_v26_ = (d_2317_v26_).set(len(_dafny.Map({d_2295_v7_: d_2299_v11_})), (self).f9)
            index407_ = default__.safeIndex(8, (d_2299_v11_).length(0))
            (d_2299_v11_)[index407_] = p0
            d_2318_v27_: _dafny.Seq
            d_2318_v27_ = _dafny.SeqWithoutIsStrInference([d_2291_v3_])
            d_2319_v28_: D10
            d_2319_v28_ = D10_DC24(p1, not((self).f9), (d_2318_v27_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2291_v3_ for d_2320_i2_ in range(default__.abs(699))])), len(d_2318_v27_))])
            index408_ = default__.safeIndex(8, (d_2299_v11_).length(0))
            (d_2299_v11_)[index408_] = (d_2319_v28_).cf31
        elif True:
            d_2291_v3_ = d_2291_v3_
            d_2321_v29_: _dafny.Array
            nw387_ = _dafny.Array(False, 5)
            d_2321_v29_ = nw387_
            index409_ = default__.safeIndex(645, (d_2321_v29_).length(0))
            (d_2321_v29_)[index409_] = (self).f9
            index410_ = default__.safeIndex(645, (d_2321_v29_).length(0))
            (d_2321_v29_)[index410_] = (d_2294_v6_) != (_dafny.Set({p0, p1, p1}))
            if (d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]:
                d_2322_v30_: _dafny.Map
                d_2322_v30_ = _dafny.Map({p1: not((d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))])})
                (globalState).f8 = (p1) * ((0) - (len(d_2322_v30_)))
                (globalState).f4 = (d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]
                d_2323_v31_: _dafny.Map
                d_2323_v31_ = _dafny.Map({p0: (d_2291_v3_) + (d_2291_v3_)})
                d_2323_v31_ = (d_2323_v31_).set(((d_2293_v5_ if (self).f9 else d_2293_v5_)).cardinality, d_2291_v3_)
                d_2324_v32_: _dafny.Array
                nw388_ = _dafny.Array(int(0), 25)
                d_2324_v32_ = nw388_
                index411_ = default__.safeIndex(30, (d_2324_v32_).length(0))
                (d_2324_v32_)[index411_] = (p0) * (p1)
                index412_ = default__.safeIndex(30, (d_2324_v32_).length(0))
                (d_2324_v32_)[index412_] = p0
                d_2325_v33_: C6
                nw389_ = C6()
                nw389_.ctor__(p1, d_2291_v3_)
                d_2325_v33_ = nw389_
                d_2326_v34_: _dafny.Map
                d_2326_v34_ = _dafny.Map({d_2325_v33_: not((d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))])})
                d_2326_v34_ = (d_2326_v34_).set(d_2325_v33_, (self).f9)
            elif True:
                d_2327_v35_: _dafny.Map
                d_2327_v35_ = _dafny.Map({p0: p1})
                d_2328_v36_: _dafny.Seq
                d_2328_v36_ = _dafny.SeqWithoutIsStrInference([(0) - (p1), p0])
                d_2329_v37_: D21
                d_2329_v37_ = D21_DC43(d_2328_v36_, (d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))])
                d_2330_v38_: _dafny.Seq
                d_2330_v38_ = d_2328_v36_
                index413_ = default__.safeIndex(645, (d_2321_v29_).length(0))
                rhs395_ = (_dafny.Map({p1: p1})) | (_dafny.Map({p1: p1}))
                rhs396_ = (self).f9
                rhs397_ = d_2321_v29_
                rhs398_ = D21_DC43((d_2328_v36_).set(default__.safeIndex(p1, len(d_2328_v36_)), p0), (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_2331_i3_ in range(default__.abs(822))]), d_2330_v38_, d_2330_v38_])) <= (_dafny.SeqWithoutIsStrInference([d_2330_v38_ for d_2332_i4_ in range(default__.abs(895))])))
                rhs399_ = (d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]
                lhs309_ = globalState
                lhs310_ = d_2321_v29_
                lhs311_ = default__.safeIndex(645, (d_2321_v29_).length(0))
                d_2327_v35_ = rhs395_
                lhs309_.f4 = rhs396_
                d_2321_v29_ = rhs397_
                d_2329_v37_ = rhs398_
                lhs310_[lhs311_] = rhs399_
                d_2333_v39_: _dafny.Array
                nw390_ = _dafny.Array(_dafny.Map({}), 18)
                d_2333_v39_ = nw390_
                d_2334_v40_: _dafny.Map
                d_2334_v40_ = _dafny.Map({(d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]: p1})
                index414_ = default__.safeIndex(277, (d_2333_v39_).length(0))
                (d_2333_v39_)[index414_] = d_2334_v40_
                d_2335_v41_: _dafny.Array
                nw391_ = _dafny.Array(_dafny.MultiSet({}), 26)
                d_2335_v41_ = nw391_
                d_2336_v42_: _dafny.MultiSet
                d_2336_v42_ = _dafny.MultiSet([(d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))], (self).f9])
                d_2337_v43_: _dafny.Map
                d_2337_v43_ = _dafny.Map({(d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]: (self).f9})
                index415_ = default__.safeIndex(277, (d_2333_v39_).length(0))
                rhs400_ = ((d_2336_v42_).intersection(d_2336_v42_)).ispropersubset(d_2336_v42_)
                rhs401_ = ((d_2328_v36_) + ((d_2328_v36_).set(default__.safeIndex(p1, len(d_2328_v36_)), p1))).set(default__.safeIndex(len(d_2337_v43_), len((d_2328_v36_) + ((d_2328_v36_).set(default__.safeIndex(p1, len(d_2328_v36_)), p1)))), p1)
                rhs402_ = (d_2334_v40_) | ((d_2334_v40_) | (d_2334_v40_))
                rhs403_ = d_2335_v41_
                lhs312_ = globalState
                lhs313_ = globalState
                lhs314_ = d_2333_v39_
                lhs315_ = default__.safeIndex(277, (d_2333_v39_).length(0))
                lhs312_.f4 = rhs400_
                lhs313_.f0 = rhs401_
                lhs314_[lhs315_] = rhs402_
                d_2335_v41_ = rhs403_
                d_2338_v45_: _dafny.Array
                def lambda117_(d_2339_v6_, d_2340_p0_):
                    def lambda118_(d_2341_i5_):
                        def iife230_():
                            coll104_ = _dafny.Map()
                            compr_104_: int
                            for compr_104_ in (d_2339_v6_).Elements:
                                d_2342_v44_: int = compr_104_
                                if (d_2342_v44_) in (d_2339_v6_):
                                    coll104_[(d_2342_v44_) - (d_2340_p0_)] = d_2340_p0_
                            return _dafny.Map(coll104_)
                        return iife230_()
                        

                    return lambda118_

                init64_ = lambda117_(d_2294_v6_, p0)
                nw392_ = _dafny.Array(None, 23)
                for i0_64_ in range(nw392_.length(0)):
                    nw392_[i0_64_] = init64_(i0_64_)
                d_2338_v45_ = nw392_
                index416_ = default__.safeIndex(856, (d_2338_v45_).length(0))
                (d_2338_v45_)[index416_] = d_2327_v35_
                d_2343_v46_: D9
                d_2343_v46_ = D9_DC21((d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))], _dafny.SeqWithoutIsStrInference([(d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]]), (self).f9)
                d_2344_v47_: _dafny.Set
                d_2344_v47_ = _dafny.Set({default__.fm30(p1, d_2343_v46_, (self).f9, p0, globalState)})
                d_2345_v50_: D3
                def iife231_():
                    def iife233_():
                        coll107_ = _dafny.Set()
                        compr_107_: int
                        for compr_107_ in _dafny.IntegerRange(924, 161):
                            d_2348_v49_: int = compr_107_
                            if ((924) <= (d_2348_v49_)) and ((d_2348_v49_) < (161)):
                                coll107_ = coll107_.union(_dafny.Set([default__.safeModuloInt(d_2348_v49_, len(d_2328_v36_))]))
                        return _dafny.Set(coll107_)
                    coll105_ = _dafny.Map()
                    def iife232_():
                        coll106_ = _dafny.Set()
                        compr_106_: int
                        for compr_106_ in _dafny.IntegerRange(924, 161):
                            d_2346_v49_: int = compr_106_
                            if ((924) <= (d_2346_v49_)) and ((d_2346_v49_) < (161)):
                                coll106_ = coll106_.union(_dafny.Set([default__.safeModuloInt(d_2346_v49_, len(d_2328_v36_))]))
                        return _dafny.Set(coll106_)
                    compr_105_: int
                    for compr_105_ in (iife232_()
                    ).Elements:
                        d_2347_v48_: int = compr_105_
                        if (d_2347_v48_) in (iife233_()
                        ):
                            coll105_[default__.safeModuloInt(d_2347_v48_, p1)] = p1
                    return _dafny.Map(coll105_)
                d_2345_v50_ = D3_DC8(iife231_()
)
                d_2349_v51_: _dafny.MultiSet
                d_2349_v51_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not((self).f9)])])
                d_2350_v52_: D24
                d_2350_v52_ = D24_DC48(d_2344_v47_)
                index417_ = default__.safeIndex(856, (d_2338_v45_).length(0))
                index418_ = default__.safeIndex(645, (d_2321_v29_).length(0))
                rhs404_ = d_2327_v35_
                rhs405_ = default__.fm26(d_2345_v50_, default__.safeDivisionInt(p1, p1), (d_2349_v51_) - (d_2349_v51_), globalState)
                rhs406_ = default__.fm49((self).f9, globalState)
                rhs407_ = default__.safeDivisionInt(p0, (0) - (p1))
                rhs408_ = ((d_2344_v47_) - ((d_2350_v52_).cf71)) | ((d_2344_v47_) | (d_2344_v47_))
                lhs316_ = d_2338_v45_
                lhs317_ = default__.safeIndex(856, (d_2338_v45_).length(0))
                lhs318_ = globalState
                lhs319_ = d_2321_v29_
                lhs320_ = default__.safeIndex(645, (d_2321_v29_).length(0))
                lhs321_ = globalState
                lhs316_[lhs317_] = rhs404_
                lhs318_.f3 = rhs405_
                lhs319_[lhs320_] = rhs406_
                lhs321_.f3 = rhs407_
                d_2344_v47_ = rhs408_
                (globalState).f4 = (d_2288_v0_)[default__.safeIndex(p1, len(d_2288_v0_))]
                rhs409_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhkfosxv"))), ((0) - (len(_dafny.SeqWithoutIsStrInference([(self).f9])))) + (p1))
                rhs410_ = -629
                lhs322_ = globalState
                lhs322_.f3 = rhs409_
                r2 = rhs410_
            d_2351_v53_: _dafny.Map
            d_2351_v53_ = _dafny.Map({(d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]: d_2294_v6_})
            d_2352_v54_: _dafny.Map
            d_2352_v54_ = _dafny.Map({p1: p0})
            rhs411_ = (d_2321_v29_)[default__.safeIndex(645, (d_2321_v29_).length(0))]
            rhs412_ = default__.safeModuloInt(len((_dafny.Map({p0: len(d_2351_v53_)})) | (d_2352_v54_)), default__.fm20(default__.fm49(True, globalState), globalState))
            lhs323_ = globalState
            r3 = rhs411_
            lhs323_.f3 = rhs412_
            d_2353_v55_: _dafny.Map
            d_2353_v55_ = _dafny.Map({(D0_DC1(630, p1)).cf2: d_2293_v5_})
            d_2354_v56_: _dafny.Map
            d_2354_v56_ = _dafny.Map({d_2353_v55_: (0) - (p1)})
            d_2354_v56_ = (d_2354_v56_).set((d_2353_v55_) | (d_2353_v55_), 652)
        d_2355_v57_: _dafny.Array
        nw393_ = _dafny.Array(False, 23)
        d_2355_v57_ = nw393_
        index419_ = default__.safeIndex(570, (d_2355_v57_).length(0))
        (d_2355_v57_)[index419_] = ((self).f9 if (self).f9 else (self).f9)
        index420_ = default__.safeIndex(570, (d_2355_v57_).length(0))
        (d_2355_v57_)[index420_] = False
        d_2356_v58_: _dafny.Array
        nw394_ = _dafny.Array(None, 4)
        nw394_[int(0)] = d_2292_v4_
        nw394_[int(1)] = _dafny.CodePoint('p')
        nw394_[int(2)] = d_2292_v4_
        nw394_[int(3)] = d_2292_v4_
        d_2356_v58_ = nw394_
        index421_ = default__.safeIndex(802, (d_2356_v58_).length(0))
        (d_2356_v58_)[index421_] = d_2292_v4_
        index422_ = default__.safeIndex(802, (d_2356_v58_).length(0))
        (d_2356_v58_)[index422_] = (d_2292_v4_ if (d_2355_v57_)[default__.safeIndex(570, (d_2355_v57_).length(0))] else (d_2292_v4_ if not((d_2355_v57_)[default__.safeIndex(570, (d_2355_v57_).length(0))]) else _dafny.CodePoint('r')))
        d_2357_v59_: _dafny.Seq
        d_2357_v59_ = _dafny.SeqWithoutIsStrInference([d_2355_v57_])
        d_2358_v60_: _dafny.Array
        nw395_ = _dafny.Array(None, 22)
        nw395_[int(0)] = d_2355_v57_
        nw395_[int(1)] = d_2355_v57_
        nw395_[int(2)] = d_2355_v57_
        nw395_[int(3)] = d_2355_v57_
        nw395_[int(4)] = d_2355_v57_
        nw395_[int(5)] = d_2355_v57_
        nw395_[int(6)] = d_2355_v57_
        nw395_[int(7)] = d_2355_v57_
        nw395_[int(8)] = d_2355_v57_
        nw395_[int(9)] = (d_2355_v57_ if (d_2355_v57_)[default__.safeIndex(570, (d_2355_v57_).length(0))] else d_2355_v57_)
        nw395_[int(10)] = d_2355_v57_
        nw395_[int(11)] = d_2355_v57_
        nw395_[int(12)] = d_2355_v57_
        nw395_[int(13)] = d_2355_v57_
        nw395_[int(14)] = d_2355_v57_
        nw395_[int(15)] = d_2355_v57_
        nw395_[int(16)] = d_2355_v57_
        nw395_[int(17)] = d_2355_v57_
        nw395_[int(18)] = (d_2357_v59_)[default__.safeIndex(len(d_2291_v3_), len(d_2357_v59_))]
        nw395_[int(19)] = d_2355_v57_
        nw395_[int(20)] = d_2355_v57_
        nw395_[int(21)] = d_2355_v57_
        d_2358_v60_ = nw395_
        d_2358_v60_ = d_2358_v60_
        d_2359_v61_: _dafny.Array
        nw396_ = _dafny.Array(D19.default()(), 10)
        d_2359_v61_ = nw396_
        d_2360_v62_: _dafny.Set
        d_2360_v62_ = _dafny.Set({d_2359_v61_})
        d_2361_v63_: _dafny.Map
        d_2361_v63_ = _dafny.Map({(self).f9: (d_2356_v58_)[default__.safeIndex(802, (d_2356_v58_).length(0))]})
        d_2362_v64_: _dafny.Array
        nw397_ = _dafny.Array(None, 20)
        nw397_[int(0)] = p0
        nw397_[int(1)] = (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f9, (self).f9])))
        nw397_[int(2)] = p0
        nw397_[int(3)] = p0
        nw397_[int(4)] = 889
        nw397_[int(5)] = p1
        nw397_[int(6)] = p0
        nw397_[int(7)] = p1
        nw397_[int(8)] = default__.fm20((d_2355_v57_)[default__.safeIndex(570, (d_2355_v57_).length(0))], globalState)
        nw397_[int(9)] = p0
        nw397_[int(10)] = p1
        nw397_[int(11)] = (0) - (p0)
        nw397_[int(12)] = p1
        nw397_[int(13)] = (0) - (p0)
        nw397_[int(14)] = p1
        nw397_[int(15)] = (0) - (p1)
        nw397_[int(16)] = p1
        nw397_[int(17)] = len(d_2361_v63_)
        nw397_[int(18)] = 240
        nw397_[int(19)] = p0
        d_2362_v64_ = nw397_
        d_2363_v65_: _dafny.Map
        d_2363_v65_ = _dafny.Map({d_2360_v62_: d_2362_v64_})
        r0 = ((d_2363_v65_)[(d_2360_v62_) - (d_2360_v62_)] if ((d_2360_v62_) - (d_2360_v62_)) in (d_2363_v65_) else d_2362_v64_)
        r0 = d_2362_v64_
        r1 = p0
        r2 = p1
        r3 = False
        return r0, r1, r2, r3

    @property
    def f9(self):
        return self._f9

class C10(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self):
        pass
        pass

    def fm2(self, globalState):
        def iife234_():
            coll108_ = _dafny.Map()
            compr_108_: int
            for compr_108_ in (_dafny.MultiSet([-85])).Elements:
                d_2364_v0_: int = compr_108_
                if (d_2364_v0_) in (_dafny.MultiSet([-85])):
                    coll108_[(d_2364_v0_) - (898)] = _dafny.CodePoint('t')
            return _dafny.Map(coll108_)
        return ((_dafny.Map({645: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([663])), len(_dafny.Set({(_dafny.MultiSet([True, True])).cardinality})), len(_dafny.Map({81: False})), (_dafny.MultiSet([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsdxerq"))])).cardinality, (0) - (len(_dafny.Map({-368: _dafny.Set({(_dafny.MultiSet([False, False])).cardinality})}))), len(iife234_()
        )])).cardinality]))})) | (_dafny.Map({len(_dafny.Set({_dafny.CodePoint('u'), _dafny.CodePoint('r')})): len(_dafny.SeqWithoutIsStrInference([23 for d_2365_i0_ in range(default__.abs(175))]))}))) | (_dafny.Map({-29: len(_dafny.Set({491}))}))

    def fm3(self, p0, p1, p2, globalState):
        return False

    def fm4(self, p0, globalState):
        def iife235_():
            coll109_ = _dafny.Map()
            compr_109_: D0
            for compr_109_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.Set({True, True})))])).Elements:
                d_2366_v0_: D0 = compr_109_
                if (d_2366_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(len(_dafny.Set({True, True})))])):
                    coll109_[d_2366_v0_] = True
            return _dafny.Map(coll109_)
        def iife236_():
            coll110_ = _dafny.Map()
            compr_110_: D0
            for compr_110_ in (_dafny.Map({D0_DC0(-672): _dafny.CodePoint('r')})).keys.Elements:
                d_2367_v1_: D0 = compr_110_
                if (d_2367_v1_) in (_dafny.Map({D0_DC0(-672): _dafny.CodePoint('r')})):
                    coll110_[d_2367_v1_] = False
            return _dafny.Map(coll110_)
        return (iife235_()
        ) | (iife236_()
        )

    def m0(self, globalState):
        r0: bool = False
        r1: bool = False
        d_2368_v0_: C0
        nw398_ = C0()
        nw398_.ctor__()
        d_2368_v0_ = nw398_
        d_2369_v1_: bool
        d_2369_v1_ = True
        d_2370_v2_: _dafny.MultiSet
        d_2370_v2_ = _dafny.MultiSet([d_2369_v1_])
        d_2370_v2_ = d_2370_v2_
        d_2371_v3_: int
        d_2371_v3_ = 428
        hi15_ = (d_2371_v3_) * ((0) - ((0) - (d_2371_v3_)))
        for d_2372_i0_ in range((d_2371_v3_) + (d_2371_v3_), hi15_):
            d_2371_v3_ = -567
            d_2373_v4_: _dafny.Seq
            d_2373_v4_ = _dafny.SeqWithoutIsStrInference([d_2369_v1_, d_2369_v1_])
            d_2374_v6_: _dafny.Map
            d_2374_v6_ = _dafny.Map({(0) - ((0) - (d_2372_i0_)): d_2369_v1_})
            d_2375_v7_: _dafny.Set
            def iife237_():
                coll111_ = _dafny.Map()
                compr_111_: int
                for compr_111_ in (d_2374_v6_).keys.Elements:
                    d_2376_v5_: int = compr_111_
                    if (d_2376_v5_) in (d_2374_v6_):
                        coll111_[(d_2376_v5_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2377_i1_ in range(default__.abs(319))])))] = d_2369_v1_
                return _dafny.Map(coll111_)
            d_2375_v7_ = _dafny.Set({len(iife237_()
            )})
            d_2378_v8_: D13
            d_2378_v8_ = D13_DC30(d_2369_v1_, (d_2373_v4_)[default__.safeIndex(392, len(d_2373_v4_))], d_2375_v7_, d_2368_v0_)
            d_2379_v9_: _dafny.Array
            nw399_ = _dafny.Array(None, 3)
            nw399_[int(0)] = d_2378_v8_
            nw399_[int(1)] = d_2378_v8_
            nw399_[int(2)] = d_2378_v8_
            d_2379_v9_ = nw399_
            index423_ = default__.safeIndex(563, (d_2379_v9_).length(0))
            (d_2379_v9_)[index423_] = d_2378_v8_
            d_2380_v10_: _dafny.Array
            nw400_ = _dafny.Array(False, 13)
            d_2380_v10_ = nw400_
            index424_ = default__.safeIndex(563, (d_2379_v9_).length(0))
            rhs413_ = d_2378_v8_
            rhs414_ = d_2371_v3_
            rhs415_ = (d_2380_v10_ if (d_2373_v4_) == (d_2373_v4_) else d_2380_v10_)
            rhs416_ = d_2371_v3_
            lhs324_ = d_2379_v9_
            lhs325_ = default__.safeIndex(563, (d_2379_v9_).length(0))
            lhs326_ = globalState
            lhs327_ = globalState
            lhs324_[lhs325_] = rhs413_
            lhs326_.f8 = rhs414_
            d_2380_v10_ = rhs415_
            lhs327_.f8 = rhs416_
            d_2381_v11_: _dafny.Array
            nw401_ = _dafny.Array(None, 17)
            d_2381_v11_ = nw401_
            d_2382_v12_: C1
            nw402_ = C1()
            nw402_.ctor__()
            d_2382_v12_ = nw402_
            index425_ = default__.safeIndex(458, (d_2381_v11_).length(0))
            (d_2381_v11_)[index425_] = d_2382_v12_
            index426_ = default__.safeIndex(458, (d_2381_v11_).length(0))
            rhs417_ = d_2372_i0_
            rhs418_ = d_2382_v12_
            lhs328_ = globalState
            lhs329_ = d_2381_v11_
            lhs330_ = default__.safeIndex(458, (d_2381_v11_).length(0))
            lhs328_.f3 = rhs417_
            lhs329_[lhs330_] = rhs418_
            d_2383_v13_: _dafny.Seq
            d_2383_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qajy"))
            d_2384_v14_: _dafny.Seq
            d_2384_v14_ = _dafny.SeqWithoutIsStrInference([63])
            d_2385_v15_: _dafny.Map
            d_2385_v15_ = _dafny.Map({default__.safeDivisionInt(len(d_2383_v13_), d_2372_i0_): (d_2384_v14_)[default__.safeIndex(d_2371_v3_, len(d_2384_v14_))]})
            d_2385_v15_ = (d_2385_v15_).set(d_2371_v3_, d_2372_i0_)
        if ((d_2370_v2_ if d_2369_v1_ else d_2370_v2_)).ispropersubset(d_2370_v2_):
            (globalState).f8 = d_2371_v3_
            if d_2369_v1_:
                d_2386_v16_: _dafny.Array
                def lambda119_(d_2387_v1_):
                    def lambda120_(d_2388_i2_):
                        return (_dafny.Map({d_2387_v1_: _dafny.CodePoint('h')})) | (_dafny.Map({d_2387_v1_: _dafny.CodePoint('b')}))

                    return lambda120_

                init65_ = lambda119_(d_2369_v1_)
                nw403_ = _dafny.Array(None, 21)
                for i0_65_ in range(nw403_.length(0)):
                    nw403_[i0_65_] = init65_(i0_65_)
                d_2386_v16_ = nw403_
                d_2386_v16_ = d_2386_v16_
                d_2389_v17_: C1
                nw404_ = C1()
                nw404_.ctor__()
                d_2389_v17_ = nw404_
                d_2390_v18_: _dafny.Map
                d_2390_v18_ = _dafny.Map({d_2389_v17_: 761})
                (globalState).f3 = (d_2371_v3_) * (((d_2390_v18_)[d_2389_v17_] if (d_2389_v17_) in (d_2390_v18_) else d_2371_v3_))
                rhs419_ = d_2369_v1_
                rhs420_ = d_2369_v1_
                lhs331_ = globalState
                r0 = rhs419_
                lhs331_.f4 = rhs420_
                d_2391_v19_: _dafny.Array
                nw405_ = _dafny.Array(_dafny.Map({}), 14)
                d_2391_v19_ = nw405_
                d_2392_v20_: T0
                nw406_ = C4()
                nw406_.ctor__(d_2391_v19_)
                d_2392_v20_ = nw406_
                d_2393_v21_: _dafny.Set
                d_2393_v21_ = _dafny.Set({d_2392_v20_})
                d_2394_v22_: _dafny.Map
                d_2394_v22_ = _dafny.Map({898: d_2393_v21_})
                d_2395_v23_: _dafny.Seq
                d_2395_v23_ = _dafny.SeqWithoutIsStrInference([d_2393_v21_])
                d_2396_v24_: _dafny.Array
                nw407_ = _dafny.Array(None, 22)
                nw407_[int(0)] = (d_2393_v21_) - (_dafny.Set({d_2392_v20_, d_2392_v20_, d_2392_v20_}))
                nw407_[int(1)] = d_2393_v21_
                nw407_[int(2)] = (d_2393_v21_) - (d_2393_v21_)
                nw407_[int(3)] = (d_2393_v21_).intersection(d_2393_v21_)
                nw407_[int(4)] = d_2393_v21_
                nw407_[int(5)] = (d_2393_v21_) | (d_2393_v21_)
                nw407_[int(6)] = (d_2393_v21_) - (d_2393_v21_)
                nw407_[int(7)] = d_2393_v21_
                nw407_[int(8)] = ((d_2394_v22_)[d_2371_v3_] if (d_2371_v3_) in (d_2394_v22_) else d_2393_v21_)
                nw407_[int(9)] = (d_2393_v21_) - (_dafny.Set({d_2392_v20_, d_2392_v20_, d_2392_v20_}))
                nw407_[int(10)] = d_2393_v21_
                nw407_[int(11)] = (d_2395_v23_)[default__.safeIndex(d_2371_v3_, len(d_2395_v23_))]
                nw407_[int(12)] = d_2393_v21_
                nw407_[int(13)] = d_2393_v21_
                nw407_[int(14)] = d_2393_v21_
                nw407_[int(15)] = d_2393_v21_
                nw407_[int(16)] = _dafny.Set({d_2392_v20_})
                nw407_[int(17)] = d_2393_v21_
                nw407_[int(18)] = d_2393_v21_
                nw407_[int(19)] = d_2393_v21_
                nw407_[int(20)] = _dafny.Set({d_2392_v20_})
                nw407_[int(21)] = (d_2393_v21_) - (_dafny.Set({d_2392_v20_}))
                d_2396_v24_ = nw407_
                index427_ = default__.safeIndex(499, (d_2396_v24_).length(0))
                (d_2396_v24_)[index427_] = d_2393_v21_
                index428_ = default__.safeIndex(499, (d_2396_v24_).length(0))
                (d_2396_v24_)[index428_] = _dafny.Set({d_2392_v20_})
                d_2397_v25_: str
                d_2397_v25_ = _dafny.CodePoint('h')
                d_2397_v25_ = d_2397_v25_
            elif True:
                d_2398_v26_: _dafny.Array
                nw408_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2398_v26_ = nw408_
                d_2399_v27_: _dafny.Seq
                d_2399_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdej"))
                d_2400_v28_: str
                d_2400_v28_ = _dafny.CodePoint('w')
                d_2401_v29_: D10
                d_2401_v29_ = D10_DC24(d_2371_v3_, d_2369_v1_, d_2399_v27_)
                d_2402_v30_: _dafny.Array
                nw409_ = _dafny.Array(None, 28)
                nw409_[int(0)] = d_2399_v27_
                nw409_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvvaguomr"))
                nw409_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vv"))
                nw409_[int(3)] = d_2399_v27_
                nw409_[int(4)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2403_i3_ in range(default__.abs(162))])).set(default__.safeIndex(d_2371_v3_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2404_i3_ in range(default__.abs(162))]))), d_2400_v28_)
                nw409_[int(5)] = d_2399_v27_
                nw409_[int(6)] = d_2399_v27_
                nw409_[int(7)] = d_2399_v27_
                nw409_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                nw409_[int(9)] = default__.fm25(globalState)
                nw409_[int(10)] = d_2399_v27_
                nw409_[int(11)] = _dafny.SeqWithoutIsStrInference([d_2400_v28_ for d_2405_i4_ in range(default__.abs(319))])
                nw409_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwsellrim"))
                nw409_[int(13)] = (d_2399_v27_).set(default__.safeIndex(d_2371_v3_, len(d_2399_v27_)), d_2400_v28_)
                nw409_[int(14)] = d_2399_v27_
                nw409_[int(15)] = (_dafny.SeqWithoutIsStrInference([d_2400_v28_ for d_2406_i5_ in range(default__.abs(-182))])).set(default__.safeIndex(352, len(_dafny.SeqWithoutIsStrInference([d_2400_v28_ for d_2407_i5_ in range(default__.abs(-182))]))), d_2400_v28_)
                nw409_[int(16)] = d_2399_v27_
                nw409_[int(17)] = d_2399_v27_
                nw409_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2400_v28_ for d_2408_i6_ in range(default__.abs(288))])
                nw409_[int(19)] = (d_2401_v29_).cf33
                nw409_[int(20)] = d_2399_v27_
                nw409_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                nw409_[int(22)] = d_2399_v27_
                nw409_[int(23)] = d_2399_v27_
                nw409_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuxk"))
                nw409_[int(25)] = d_2399_v27_
                nw409_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmxp"))
                nw409_[int(27)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsdb"))
                d_2402_v30_ = nw409_
                index429_ = default__.safeIndex(429, (d_2398_v26_).length(0))
                (d_2398_v26_)[index429_] = d_2402_v30_
                index430_ = default__.safeIndex(429, (d_2398_v26_).length(0))
                nw410_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                (d_2398_v26_)[index430_] = nw410_
                d_2409_v31_: _dafny.Map
                d_2409_v31_ = _dafny.Map({d_2371_v3_: True})
                d_2410_v32_: _dafny.Seq
                d_2410_v32_ = _dafny.SeqWithoutIsStrInference([d_2399_v27_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgmejkpo")), d_2399_v27_, d_2399_v27_, d_2399_v27_])
                d_2411_v33_: _dafny.Set
                d_2411_v33_ = _dafny.Set({d_2369_v1_, d_2369_v1_})
                d_2412_v34_: _dafny.Array
                nw411_ = _dafny.Array(None, 26)
                nw411_[int(0)] = (d_2370_v2_).cardinality
                nw411_[int(1)] = d_2371_v3_
                nw411_[int(2)] = (0) - (len(d_2409_v31_))
                nw411_[int(3)] = d_2371_v3_
                nw411_[int(4)] = (0) - (len(d_2410_v32_))
                nw411_[int(5)] = d_2371_v3_
                nw411_[int(6)] = d_2371_v3_
                nw411_[int(7)] = d_2371_v3_
                nw411_[int(8)] = d_2371_v3_
                nw411_[int(9)] = (0) - (d_2371_v3_)
                nw411_[int(10)] = d_2371_v3_
                nw411_[int(11)] = len(d_2411_v33_)
                nw411_[int(12)] = d_2371_v3_
                nw411_[int(13)] = len(_dafny.SeqWithoutIsStrInference([d_2371_v3_ for d_2413_i7_ in range(default__.abs(603))]))
                nw411_[int(14)] = d_2371_v3_
                nw411_[int(15)] = d_2371_v3_
                nw411_[int(16)] = d_2371_v3_
                nw411_[int(17)] = d_2371_v3_
                nw411_[int(18)] = d_2371_v3_
                nw411_[int(19)] = d_2371_v3_
                nw411_[int(20)] = d_2371_v3_
                nw411_[int(21)] = (d_2368_v0_).fm19(globalState)
                nw411_[int(22)] = d_2371_v3_
                nw411_[int(23)] = d_2371_v3_
                nw411_[int(24)] = d_2371_v3_
                nw411_[int(25)] = len(d_2399_v27_)
                d_2412_v34_ = nw411_
                d_2414_v35_: _dafny.Array
                nw412_ = _dafny.Array(None, 10)
                nw412_[int(0)] = d_2412_v34_
                nw412_[int(1)] = d_2412_v34_
                nw412_[int(2)] = d_2412_v34_
                nw412_[int(3)] = d_2412_v34_
                nw412_[int(4)] = d_2412_v34_
                nw412_[int(5)] = d_2412_v34_
                nw412_[int(6)] = d_2412_v34_
                nw412_[int(7)] = d_2412_v34_
                nw412_[int(8)] = d_2412_v34_
                nw412_[int(9)] = d_2412_v34_
                d_2414_v35_ = nw412_
                index431_ = default__.safeIndex(394, (d_2414_v35_).length(0))
                (d_2414_v35_)[index431_] = d_2412_v34_
                index432_ = default__.safeIndex(394, (d_2414_v35_).length(0))
                rhs421_ = d_2412_v34_
                rhs422_ = d_2369_v1_
                rhs423_ = (d_2371_v3_) < (d_2371_v3_)
                lhs332_ = d_2414_v35_
                lhs333_ = default__.safeIndex(394, (d_2414_v35_).length(0))
                lhs332_[lhs333_] = rhs421_
                r0 = rhs422_
                d_2369_v1_ = rhs423_
                d_2415_v36_: _dafny.Array
                nw413_ = _dafny.Array(D23.default()(), 12)
                d_2415_v36_ = nw413_
                d_2416_v37_: _dafny.Array
                nw414_ = _dafny.Array(None, 4)
                nw414_[int(0)] = d_2400_v28_
                nw414_[int(1)] = d_2400_v28_
                nw414_[int(2)] = d_2400_v28_
                nw414_[int(3)] = d_2400_v28_
                d_2416_v37_ = nw414_
                d_2417_v38_: D23
                d_2417_v38_ = D23_DC45(d_2416_v37_)
                index433_ = default__.safeIndex(282, (d_2415_v36_).length(0))
                (d_2415_v36_)[index433_] = d_2417_v38_
                pat_let_tv73_ = d_2416_v37_
                index434_ = default__.safeIndex(282, (d_2415_v36_).length(0))
                def iife238_(_pat_let63_0):
                    def iife239_(d_2418_dt__update__tmp_h0_):
                        def iife240_(_pat_let64_0):
                            def iife241_(d_2419_dt__update_hcf66_h0_):
                                return D23_DC45(d_2419_dt__update_hcf66_h0_)
                            return iife241_(_pat_let64_0)
                        return iife240_(pat_let_tv73_)
                    return iife239_(_pat_let63_0)
                (d_2415_v36_)[index434_] = iife238_(d_2417_v38_)
                d_2420_v39_: _dafny.Map
                d_2420_v39_ = _dafny.Map({d_2371_v3_: d_2371_v3_})
                d_2421_v40_: D2
                d_2421_v40_ = D2_DC7(d_2400_v28_, d_2369_v1_, 786)
                d_2422_v41_: _dafny.Seq
                d_2422_v41_ = _dafny.SeqWithoutIsStrInference([d_2411_v33_])
                d_2420_v39_ = (d_2420_v39_).set((d_2421_v40_).cf11, len((d_2411_v33_).intersection((d_2422_v41_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guqrkvlrc"))), len(d_2422_v41_))])))
                (globalState).f3 = (0) - (d_2371_v3_)
            d_2423_v42_: _dafny.Array
            def lambda121_(d_2424_v1_):
                def lambda122_(d_2425_i8_):
                    return d_2424_v1_

                return lambda122_

            init66_ = lambda121_(d_2369_v1_)
            nw415_ = _dafny.Array(None, 12)
            for i0_66_ in range(nw415_.length(0)):
                nw415_[i0_66_] = init66_(i0_66_)
            d_2423_v42_ = nw415_
            d_2423_v42_ = d_2423_v42_
            d_2426_v43_: _dafny.Array
            nw416_ = _dafny.Array(int(0), 3)
            d_2426_v43_ = nw416_
            index435_ = default__.safeIndex(857, (d_2426_v43_).length(0))
            (d_2426_v43_)[index435_] = (0) - (-180)
            d_2427_v44_: _dafny.Map
            d_2427_v44_ = _dafny.Map({d_2371_v3_: d_2371_v3_})
            d_2428_v45_: _dafny.Map
            d_2428_v45_ = _dafny.Map({d_2427_v44_: d_2371_v3_})
            index436_ = default__.safeIndex(857, (d_2426_v43_).length(0))
            (d_2426_v43_)[index436_] = ((d_2428_v45_)[d_2427_v44_] if (d_2427_v44_) in (d_2428_v45_) else d_2371_v3_)
            d_2429_v46_: str
            d_2429_v46_ = _dafny.CodePoint('d')
            d_2430_v47_: _dafny.Seq
            d_2430_v47_ = _dafny.SeqWithoutIsStrInference([d_2429_v46_, d_2429_v46_, d_2429_v46_])
            d_2431_v48_: _dafny.Seq
            d_2431_v48_ = _dafny.SeqWithoutIsStrInference([d_2429_v46_, d_2429_v46_, d_2429_v46_, (d_2430_v47_)[default__.safeIndex(d_2371_v3_, len(d_2430_v47_))]])
            d_2430_v47_ = d_2430_v47_
        elif True:
            if False:
                d_2432_v49_: _dafny.Array
                def lambda123_(d_2433_v3_):
                    def lambda124_(d_2434_i9_):
                        return default__.safeModuloInt(d_2434_i9_, d_2433_v3_)

                    return lambda124_

                init67_ = lambda123_(d_2371_v3_)
                nw417_ = _dafny.Array(None, 21)
                for i0_67_ in range(nw417_.length(0)):
                    nw417_[i0_67_] = init67_(i0_67_)
                d_2432_v49_ = nw417_
                index437_ = default__.safeIndex(192, (d_2432_v49_).length(0))
                (d_2432_v49_)[index437_] = d_2371_v3_
                index438_ = default__.safeIndex(192, (d_2432_v49_).length(0))
                (d_2432_v49_)[index438_] = d_2371_v3_
                d_2435_v50_: _dafny.Array
                def lambda125_(d_2436_v1_):
                    def lambda126_(d_2437_i10_):
                        return d_2436_v1_

                    return lambda126_

                init68_ = lambda125_(d_2369_v1_)
                nw418_ = _dafny.Array(None, 10)
                for i0_68_ in range(nw418_.length(0)):
                    nw418_[i0_68_] = init68_(i0_68_)
                d_2435_v50_ = nw418_
                index439_ = default__.safeIndex(256, (d_2435_v50_).length(0))
                (d_2435_v50_)[index439_] = d_2369_v1_
                index440_ = default__.safeIndex(136, (d_2435_v50_).length(0))
                (d_2435_v50_)[index440_] = d_2369_v1_
                d_2438_v51_: _dafny.Map
                d_2438_v51_ = _dafny.Map({(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]: d_2369_v1_})
                d_2439_v52_: _dafny.Seq
                d_2439_v52_ = _dafny.SeqWithoutIsStrInference([((d_2438_v51_)[d_2371_v3_] if (d_2371_v3_) in (d_2438_v51_) else d_2369_v1_)])
                d_2440_v53_: _dafny.Map
                d_2440_v53_ = _dafny.Map({d_2369_v1_: (d_2368_v0_).fm19(globalState)})
                d_2441_v54_: _dafny.MultiSet
                d_2441_v54_ = _dafny.MultiSet([(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))], d_2371_v3_, (d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))], len(d_2440_v53_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "youkw")))])
                d_2442_v55_: _dafny.MultiSet
                d_2442_v55_ = _dafny.MultiSet([(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))], (d_2441_v54_).cardinality, d_2371_v3_])
                d_2443_v56_: _dafny.MultiSet
                d_2443_v56_ = _dafny.MultiSet([d_2442_v55_, d_2441_v54_])
                d_2444_v57_: _dafny.Map
                d_2444_v57_ = _dafny.Map({(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]: d_2443_v56_})
                d_2445_v58_: _dafny.Seq
                d_2445_v58_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2446_i11_ in range(default__.abs(455))])])
                d_2447_v59_: _dafny.Seq
                d_2447_v59_ = _dafny.SeqWithoutIsStrInference([(((d_2444_v57_)[(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]] if ((d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]) in (d_2444_v57_) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2441_v54_, d_2441_v54_])))).cardinality, (d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))], (len(d_2445_v58_)) + ((d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))])])
                index441_ = default__.safeIndex(256, (d_2435_v50_).length(0))
                index442_ = default__.safeIndex(136, (d_2435_v50_).length(0))
                rhs424_ = (0) - ((0) - (-617))
                rhs425_ = d_2369_v1_
                rhs426_ = (d_2439_v52_)[default__.safeIndex(d_2371_v3_, len(d_2439_v52_))]
                rhs427_ = d_2369_v1_
                rhs428_ = (d_2447_v59_)[default__.safeIndex((d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))], len(d_2447_v59_))]
                lhs334_ = d_2435_v50_
                lhs335_ = default__.safeIndex(256, (d_2435_v50_).length(0))
                lhs336_ = d_2435_v50_
                lhs337_ = default__.safeIndex(136, (d_2435_v50_).length(0))
                lhs338_ = globalState
                d_2371_v3_ = rhs424_
                lhs334_[lhs335_] = rhs425_
                lhs336_[lhs337_] = rhs426_
                d_2369_v1_ = rhs427_
                lhs338_.f3 = rhs428_
                d_2448_v60_: str
                d_2448_v60_ = _dafny.CodePoint('n')
                d_2449_v61_: _dafny.Map
                d_2449_v61_ = _dafny.Map({d_2369_v1_: d_2448_v60_})
                d_2450_v62_: _dafny.Set
                d_2450_v62_ = _dafny.Set({d_2371_v3_, len(d_2449_v61_)})
                d_2451_v63_: _dafny.Set
                d_2451_v63_ = _dafny.Set({len(d_2450_v62_), (d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]})
                d_2452_v64_: _dafny.Set
                d_2452_v64_ = _dafny.Set({d_2450_v62_, (d_2451_v63_) | (_dafny.Set({len(d_2439_v52_)})), d_2450_v62_})
                d_2453_v65_: _dafny.Array
                def lambda127_(d_2454_v51_):
                    def lambda128_(d_2455_i12_):
                        return d_2454_v51_

                    return lambda128_

                init69_ = lambda127_(d_2438_v51_)
                nw419_ = _dafny.Array(None, 20)
                for i0_69_ in range(nw419_.length(0)):
                    nw419_[i0_69_] = init69_(i0_69_)
                d_2453_v65_ = nw419_
                index443_ = default__.safeIndex(200, (d_2453_v65_).length(0))
                def iife242_():
                    coll112_ = _dafny.Map()
                    compr_112_: int
                    for compr_112_ in _dafny.IntegerRange(137, 239):
                        d_2456_v66_: int = compr_112_
                        if ((137) <= (d_2456_v66_)) and ((d_2456_v66_) < (239)):
                            coll112_[(d_2456_v66_) + (520)] = d_2369_v1_
                    return _dafny.Map(coll112_)
                (d_2453_v65_)[index443_] = iife242_()
                
                index444_ = default__.safeIndex(200, (d_2453_v65_).length(0))
                rhs429_ = d_2452_v64_
                rhs430_ = (d_2438_v51_) | (d_2438_v51_)
                rhs431_ = (d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]
                lhs339_ = d_2453_v65_
                lhs340_ = default__.safeIndex(200, (d_2453_v65_).length(0))
                lhs341_ = globalState
                d_2452_v64_ = rhs429_
                lhs339_[lhs340_] = rhs430_
                lhs341_.f8 = rhs431_
                d_2457_v67_: _dafny.Seq
                d_2457_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "catkgef"))
                d_2458_v68_: _dafny.Map
                d_2458_v68_ = _dafny.Map({(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]: d_2457_v67_})
                index445_ = default__.safeIndex(192, (d_2432_v49_).length(0))
                (d_2432_v49_)[index445_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ds"))) + (((d_2458_v68_)[(d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]] if ((d_2432_v49_)[default__.safeIndex(192, (d_2432_v49_).length(0))]) in (d_2458_v68_) else d_2457_v67_)))
                d_2459_v69_: D2
                d_2459_v69_ = D2_DC5(d_2435_v50_)
                d_2460_v70_: _dafny.Array
                nw420_ = _dafny.Array(None, 13)
                nw420_[int(0)] = d_2435_v50_
                nw420_[int(1)] = (d_2435_v50_ if (d_2435_v50_)[default__.safeIndex(256, (d_2435_v50_).length(0))] else d_2435_v50_)
                nw420_[int(2)] = d_2435_v50_
                nw420_[int(3)] = d_2435_v50_
                nw420_[int(4)] = d_2435_v50_
                nw420_[int(5)] = d_2435_v50_
                nw420_[int(6)] = d_2435_v50_
                nw420_[int(7)] = d_2435_v50_
                nw420_[int(8)] = d_2435_v50_
                nw420_[int(9)] = (d_2459_v69_).cf5
                nw420_[int(10)] = d_2435_v50_
                nw420_[int(11)] = d_2435_v50_
                nw420_[int(12)] = d_2435_v50_
                d_2460_v70_ = nw420_
                index446_ = default__.safeIndex(183, (d_2460_v70_).length(0))
                (d_2460_v70_)[index446_] = d_2435_v50_
                index447_ = default__.safeIndex(183, (d_2460_v70_).length(0))
                (d_2460_v70_)[index447_] = d_2435_v50_
            elif True:
                d_2461_v71_: str
                d_2461_v71_ = _dafny.CodePoint('t')
                d_2462_v72_: _dafny.Seq
                d_2462_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gi"))
                d_2463_v73_: _dafny.Array
                nw421_ = _dafny.Array(None, 12)
                nw421_[int(0)] = d_2461_v71_
                nw421_[int(1)] = d_2461_v71_
                nw421_[int(2)] = d_2461_v71_
                nw421_[int(3)] = d_2461_v71_
                nw421_[int(4)] = d_2461_v71_
                nw421_[int(5)] = d_2461_v71_
                nw421_[int(6)] = d_2461_v71_
                nw421_[int(7)] = _dafny.CodePoint('p')
                nw421_[int(8)] = d_2461_v71_
                nw421_[int(9)] = (d_2462_v72_)[default__.safeIndex(650, len(d_2462_v72_))]
                nw421_[int(10)] = d_2461_v71_
                nw421_[int(11)] = (d_2461_v71_ if d_2369_v1_ else d_2461_v71_)
                d_2463_v73_ = nw421_
                index448_ = default__.safeIndex(195, (d_2463_v73_).length(0))
                (d_2463_v73_)[index448_] = _dafny.CodePoint('x')
                d_2464_v74_: _dafny.Map
                d_2464_v74_ = _dafny.Map({d_2371_v3_: (d_2462_v72_)[default__.safeIndex(d_2371_v3_, len(d_2462_v72_))]})
                index449_ = default__.safeIndex(195, (d_2463_v73_).length(0))
                (d_2463_v73_)[index449_] = ((d_2464_v74_)[d_2371_v3_] if (d_2371_v3_) in (d_2464_v74_) else d_2461_v71_)
                d_2465_v75_: _dafny.Array
                def lambda129_(d_2466_v3_, d_2467_v2_):
                    def lambda130_(d_2468_i13_):
                        return default__.safeDivisionInt(d_2468_i13_, (_dafny.MultiSet([(0) - (d_2466_v3_), 966, (d_2467_v2_).cardinality, d_2466_v3_])).cardinality)

                    return lambda130_

                init70_ = lambda129_(d_2371_v3_, d_2370_v2_)
                nw422_ = _dafny.Array(None, 26)
                for i0_70_ in range(nw422_.length(0)):
                    nw422_[i0_70_] = init70_(i0_70_)
                d_2465_v75_ = nw422_
                index450_ = default__.safeIndex(106, (d_2465_v75_).length(0))
                (d_2465_v75_)[index450_] = d_2371_v3_
                index451_ = default__.safeIndex(106, (d_2465_v75_).length(0))
                (d_2465_v75_)[index451_] = default__.safeModuloInt(default__.fm6(d_2371_v3_, globalState), (0) - (len(d_2462_v72_)))
                d_2469_v76_: C3
                nw423_ = C3()
                nw423_.ctor__()
                d_2469_v76_ = nw423_
                d_2470_v77_: C3
                d_2470_v77_ = d_2469_v76_
                d_2471_v78_: _dafny.Seq
                d_2471_v78_ = _dafny.SeqWithoutIsStrInference([d_2470_v77_, d_2469_v76_, d_2470_v77_, d_2470_v77_])
                d_2471_v78_ = d_2471_v78_
                (globalState).f3 = (len(_dafny.SeqWithoutIsStrInference([(d_2463_v73_)[default__.safeIndex(195, (d_2463_v73_).length(0))] for d_2472_i14_ in range(default__.abs(436))]))) - (d_2371_v3_)
                (globalState).f4 = d_2369_v1_
            (globalState).f3 = (d_2371_v3_) * ((d_2371_v3_) * (len(_dafny.Map({d_2371_v3_: d_2369_v1_}))))
            d_2473_v79_: _dafny.Map
            d_2473_v79_ = _dafny.Map({default__.fm49(d_2369_v1_, globalState): d_2369_v1_})
            (globalState).f4 = (d_2369_v1_) not in (d_2473_v79_)
            d_2474_v80_: _dafny.Array
            nw424_ = _dafny.Array(_dafny.Map({}), 22)
            d_2474_v80_ = nw424_
            d_2475_v81_: _dafny.Seq
            d_2475_v81_ = _dafny.SeqWithoutIsStrInference([d_2369_v1_])
            d_2476_v82_: _dafny.Seq
            d_2476_v82_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_2475_v81_)), d_2371_v3_])
            d_2477_v83_: _dafny.Map
            d_2477_v83_ = _dafny.Map({d_2369_v1_: d_2476_v82_})
            index452_ = default__.safeIndex(207, (d_2474_v80_).length(0))
            (d_2474_v80_)[index452_] = (d_2477_v83_).set(d_2369_v1_, d_2476_v82_)
            index453_ = default__.safeIndex(207, (d_2474_v80_).length(0))
            (d_2474_v80_)[index453_] = _dafny.Map({(d_2369_v1_) and (d_2369_v1_): _dafny.SeqWithoutIsStrInference([d_2371_v3_ for d_2478_i15_ in range(default__.abs(-944))])})
            d_2479_v84_: str
            d_2479_v84_ = _dafny.CodePoint('j')
            d_2480_v85_: _dafny.Seq
            d_2480_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pudx"))
            r1 = (d_2479_v84_) in (d_2480_v85_)
        d_2481_v86_: str
        d_2481_v86_ = _dafny.CodePoint('b')
        d_2482_v87_: _dafny.Map
        d_2482_v87_ = _dafny.Map({-713: d_2481_v86_})
        d_2483_v88_: D19
        d_2483_v88_ = D19_DC38(d_2482_v87_)
        d_2484_v89_: D19
        d_2484_v89_ = D19_DC40(d_2483_v88_)
        d_2485_v90_: D19
        d_2485_v90_ = D19_DC40(d_2483_v88_)
        source40_ = d_2485_v90_
        if source40_.is_DC39:
            d_2486___mcc_h0_ = source40_.cf58
            d_2487___mcc_h1_ = source40_.cf59
            d_2488_cf59_ = d_2487___mcc_h1_
            d_2489_cf58_ = d_2486___mcc_h0_
            d_2490_v91_: _dafny.Seq
            d_2490_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imuchiwj"))
            d_2491_v92_: _dafny.Array
            nw425_ = _dafny.Array(int(0), 18)
            d_2491_v92_ = nw425_
            d_2492_v93_: _dafny.Map
            d_2492_v93_ = _dafny.Map({d_2371_v3_: d_2491_v92_})
            d_2493_v94_: _dafny.Map
            d_2493_v94_ = _dafny.Map({d_2371_v3_: d_2371_v3_})
            d_2371_v3_ = (len((d_2490_v91_).set(default__.safeIndex(len(d_2492_v93_), len(d_2490_v91_)), d_2481_v86_))) * ((((d_2493_v94_)[default__.fm20(d_2489_cf58_, globalState)] if (default__.fm20(d_2489_cf58_, globalState)) in (d_2493_v94_) else d_2371_v3_)) - (890))
            r0 = d_2489_cf58_
            (globalState).f4 = d_2369_v1_
            d_2491_v92_ = d_2491_v92_
        elif source40_.is_DC38:
            d_2494___mcc_h2_ = source40_.cf57
            d_2495_cf57_ = d_2494___mcc_h2_
            d_2496_v95_: _dafny.MultiSet
            d_2496_v95_ = _dafny.MultiSet([(0) - (d_2371_v3_)])
            d_2497_v96_: _dafny.MultiSet
            d_2497_v96_ = d_2496_v95_
            d_2497_v96_ = d_2497_v96_
            d_2498_v97_: _dafny.Array
            nw426_ = _dafny.Array(int(0), 19)
            d_2498_v97_ = nw426_
            d_2499_v98_: _dafny.Map
            d_2499_v98_ = _dafny.Map({d_2498_v97_: d_2371_v3_})
            d_2500_v99_: D21
            d_2500_v99_ = D21_DC42(d_2499_v98_)
            d_2501_v100_: _dafny.Set
            d_2501_v100_ = _dafny.Set({d_2369_v1_})
            d_2502_v101_: _dafny.Map
            d_2502_v101_ = _dafny.Map({d_2369_v1_: False})
            d_2503_v102_: D0
            d_2503_v102_ = D0_DC0(len(d_2502_v101_))
            rhs432_ = ((d_2496_v95_ if d_2369_v1_ else d_2496_v95_)) - ((d_2496_v95_) - (_dafny.MultiSet([((d_2370_v2_)[d_2369_v1_] if (d_2369_v1_) in (d_2370_v2_) else d_2371_v3_), d_2371_v3_, len(d_2501_v100_), d_2371_v3_, 309])))
            rhs433_ = d_2500_v99_
            rhs434_ = (d_2503_v102_).cf0
            lhs342_ = globalState
            lhs343_ = globalState
            lhs342_.f2 = rhs432_
            d_2500_v99_ = rhs433_
            lhs343_.f3 = rhs434_
            d_2504_v103_: _dafny.Seq
            d_2504_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jo"))
            d_2505_v104_: _dafny.Map
            d_2505_v104_ = _dafny.Map({(len(d_2504_v103_)) != (d_2371_v3_): d_2371_v3_})
            (globalState).f3 = len(d_2505_v104_)
            (globalState).f8 = (0) - (default__.safeDivisionInt(default__.safeDivisionInt(d_2371_v3_, len(d_2502_v101_)), ((d_2370_v2_).set(d_2369_v1_, default__.abs(d_2371_v3_))).cardinality))
        elif True:
            d_2506___mcc_h3_ = source40_.cf60
            d_2507_cf60_ = d_2506___mcc_h3_
            d_2508_v105_: _dafny.Seq
            d_2508_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgp"))
            rhs435_ = d_2369_v1_
            rhs436_ = d_2369_v1_
            rhs437_ = (d_2508_v105_) < (((d_2508_v105_).set(default__.safeIndex(d_2371_v3_, len(d_2508_v105_)), d_2481_v86_)) + (_dafny.SeqWithoutIsStrInference([d_2481_v86_ for d_2509_i16_ in range(default__.abs(230))])))
            rhs438_ = d_2369_v1_
            lhs344_ = globalState
            lhs345_ = globalState
            d_2369_v1_ = rhs435_
            lhs344_.f4 = rhs436_
            r0 = rhs437_
            lhs345_.f4 = rhs438_
            d_2510_v106_: _dafny.Array
            nw427_ = _dafny.Array(int(0), 8)
            d_2510_v106_ = nw427_
            d_2510_v106_ = d_2510_v106_
            d_2511_v107_: _dafny.Set
            d_2511_v107_ = _dafny.Set({d_2481_v86_})
            d_2512_v108_: C3
            nw428_ = C3()
            nw428_.ctor__()
            d_2512_v108_ = nw428_
            d_2513_v109_: _dafny.Map
            d_2513_v109_ = _dafny.Map({(d_2511_v107_).intersection(d_2511_v107_): d_2512_v108_})
            d_2513_v109_ = (d_2513_v109_).set(d_2511_v107_, d_2512_v108_)
            (globalState).f4 = d_2369_v1_
        d_2514_v110_: _dafny.Map
        d_2514_v110_ = _dafny.Map({d_2369_v1_: d_2481_v86_})
        (globalState).f4 = (not(d_2369_v1_)) not in ((d_2514_v110_) | (d_2514_v110_))
        r0 = d_2369_v1_
        r1 = d_2369_v1_
        return r0, r1

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        d_2515_v0_: _dafny.Map
        d_2515_v0_ = _dafny.Map({p0: p1})
        d_2516_v1_: _dafny.MultiSet
        d_2516_v1_ = _dafny.MultiSet([((d_2515_v0_)[p0] if (p0) in (d_2515_v0_) else p1), p1, p1])
        source41_ = default__.fm40((d_2516_v1_).issubset(_dafny.MultiSet([p1])), globalState)
        if source41_.is_DC6:
            d_2517___mcc_h0_ = source41_.cf6
            d_2518___mcc_h1_ = source41_.cf7
            d_2519___mcc_h2_ = source41_.cf8
            d_2520_cf8_ = d_2519___mcc_h2_
            d_2521_cf7_ = d_2518___mcc_h1_
            d_2522_cf6_ = d_2517___mcc_h0_
            d_2523_v2_: C3
            nw429_ = C3()
            nw429_.ctor__()
            d_2523_v2_ = nw429_
            d_2524_v3_: _dafny.MultiSet
            d_2524_v3_ = _dafny.MultiSet([d_2523_v2_])
            d_2525_v4_: _dafny.Set
            d_2525_v4_ = _dafny.Set({(d_2522_cf6_) - (p0), (d_2522_cf6_) + ((d_2524_v3_).cardinality), p0, default__.safeDivisionInt(293, d_2522_cf6_), d_2522_cf6_})
            d_2526_v5_: _dafny.Array
            nw430_ = _dafny.Array(False, 23)
            d_2526_v5_ = nw430_
            index454_ = default__.safeIndex(739, (d_2526_v5_).length(0))
            (d_2526_v5_)[index454_] = not(p1)
            d_2527_v6_: _dafny.Array
            nw431_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_2527_v6_ = nw431_
            d_2528_v7_: _dafny.Seq
            d_2528_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbp"))
            index455_ = default__.safeIndex(235, (d_2527_v6_).length(0))
            (d_2527_v6_)[index455_] = d_2528_v7_
            index456_ = default__.safeIndex(739, (d_2526_v5_).length(0))
            index457_ = default__.safeIndex(235, (d_2527_v6_).length(0))
            rhs439_ = (d_2525_v4_) - (d_2525_v4_)
            rhs440_ = ((p0) - (p0)) != ((0) - ((p0) - (d_2522_cf6_)))
            rhs441_ = (d_2528_v7_) + (d_2528_v7_)
            rhs442_ = d_2520_cf8_
            lhs346_ = d_2526_v5_
            lhs347_ = default__.safeIndex(739, (d_2526_v5_).length(0))
            lhs348_ = d_2527_v6_
            lhs349_ = default__.safeIndex(235, (d_2527_v6_).length(0))
            lhs350_ = globalState
            d_2525_v4_ = rhs439_
            lhs346_[lhs347_] = rhs440_
            lhs348_[lhs349_] = rhs441_
            lhs350_.f4 = rhs442_
            d_2529_v8_: str
            d_2529_v8_ = _dafny.CodePoint('y')
            d_2530_v9_: _dafny.Map
            d_2530_v9_ = _dafny.Map({p1: (d_2529_v8_) in ((d_2527_v6_)[default__.safeIndex(235, (d_2527_v6_).length(0))])})
            d_2530_v9_ = default__.fm21(not((d_2526_v5_)[default__.safeIndex(739, (d_2526_v5_).length(0))]), globalState)
            index458_ = default__.safeIndex(739, (d_2526_v5_).length(0))
            (d_2526_v5_)[index458_] = not((d_2526_v5_)[default__.safeIndex(739, (d_2526_v5_).length(0))])
            d_2531_v11_: _dafny.Seq
            def iife243_():
                coll113_ = _dafny.Map()
                compr_113_: int
                for compr_113_ in _dafny.IntegerRange(-461, 451):
                    d_2532_v10_: int = compr_113_
                    if ((-461) <= (d_2532_v10_)) and ((d_2532_v10_) < (451)):
                        coll113_[(d_2532_v10_) - (p0)] = ((d_2527_v6_)[default__.safeIndex(235, (d_2527_v6_).length(0))]).set(default__.safeIndex(p0, len((d_2527_v6_)[default__.safeIndex(235, (d_2527_v6_).length(0))])), d_2529_v8_)
                return _dafny.Map(coll113_)
            d_2531_v11_ = _dafny.SeqWithoutIsStrInference([len(iife243_()
            )])
            (globalState).f0 = (d_2531_v11_).set(default__.safeIndex(720, len(d_2531_v11_)), p0)
        elif source41_.is_DC7:
            d_2533___mcc_h3_ = source41_.cf9
            d_2534___mcc_h4_ = source41_.cf10
            d_2535___mcc_h5_ = source41_.cf11
            d_2536_cf11_ = d_2535___mcc_h5_
            d_2537_cf10_ = d_2534___mcc_h4_
            d_2538_cf9_ = d_2533___mcc_h3_
            d_2539_v12_: C5
            nw432_ = C5()
            nw432_.ctor__()
            d_2539_v12_ = nw432_
            d_2540_v13_: C5
            nw433_ = C5()
            nw433_.ctor__()
            d_2540_v13_ = nw433_
            (globalState).f8 = (0) - (-857)
            if (d_2516_v1_) != (d_2516_v1_):
                d_2541_v14_: _dafny.Array
                def lambda131_(d_2542_cf11_):
                    def lambda132_(d_2543_i0_):
                        return default__.safeDivisionInt(d_2543_i0_, d_2542_cf11_)

                    return lambda132_

                init71_ = lambda131_(d_2536_cf11_)
                nw434_ = _dafny.Array(None, 21)
                for i0_71_ in range(nw434_.length(0)):
                    nw434_[i0_71_] = init71_(i0_71_)
                d_2541_v14_ = nw434_
                d_2541_v14_ = d_2541_v14_
                d_2544_v15_: _dafny.Array
                nw435_ = _dafny.Array(None, 11)
                d_2544_v15_ = nw435_
                d_2545_v16_: C1
                nw436_ = C1()
                nw436_.ctor__()
                d_2545_v16_ = nw436_
                index459_ = default__.safeIndex(53, (d_2544_v15_).length(0))
                (d_2544_v15_)[index459_] = d_2545_v16_
                d_2546_v17_: _dafny.Map
                d_2546_v17_ = _dafny.Map({p1: d_2541_v14_})
                d_2547_v18_: T0
                nw437_ = C3()
                nw437_.ctor__()
                d_2547_v18_ = nw437_
                d_2548_v19_: _dafny.Seq
                d_2548_v19_ = _dafny.SeqWithoutIsStrInference([d_2546_v17_, (d_2546_v17_ if d_2537_cf10_ else d_2546_v17_), d_2546_v17_])
                d_2549_v20_: _dafny.Seq
                d_2549_v20_ = _dafny.SeqWithoutIsStrInference([d_2541_v14_, ((d_2546_v17_)[p1] if (p1) in (d_2546_v17_) else d_2541_v14_)])
                index460_ = default__.safeIndex(53, (d_2544_v15_).length(0))
                rhs443_ = d_2545_v16_
                rhs444_ = d_2536_cf11_
                rhs445_ = (d_2548_v19_)[default__.safeIndex(default__.safeModuloInt(p0, d_2536_cf11_), len(d_2548_v19_))]
                rhs446_ = d_2547_v18_
                rhs447_ = (d_2541_v14_) not in (d_2549_v20_)
                lhs351_ = d_2544_v15_
                lhs352_ = default__.safeIndex(53, (d_2544_v15_).length(0))
                lhs353_ = globalState
                lhs351_[lhs352_] = rhs443_
                lhs353_.f3 = rhs444_
                d_2546_v17_ = rhs445_
                d_2547_v18_ = rhs446_
                d_2537_cf10_ = rhs447_
                d_2550_v21_: _dafny.Map
                d_2550_v21_ = _dafny.Map({(d_2537_cf10_) or (p1): (0) - (p0)})
                d_2550_v21_ = (d_2550_v21_).set(d_2537_cf10_, p0)
                index461_ = default__.safeIndex(734, (d_2541_v14_).length(0))
                (d_2541_v14_)[index461_] = (0) - (p0)
                index462_ = default__.safeIndex(734, (d_2541_v14_).length(0))
                (d_2541_v14_)[index462_] = default__.fm6(p0, globalState)
                r0 = 321
            elif True:
                d_2551_v22_: _dafny.Seq
                d_2551_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
                d_2551_v22_ = d_2551_v22_
                (globalState).f3 = p0
                (globalState).f8 = (default__.fm20(d_2537_cf10_, globalState)) + ((212) - (d_2536_cf11_))
                (globalState).f4 = d_2537_cf10_
                d_2552_v23_: D2
                d_2552_v23_ = D2_DC6(d_2536_cf11_, False, p1)
                d_2552_v23_ = d_2552_v23_
        elif True:
            d_2553___mcc_h6_ = source41_.cf5
            d_2554_cf5_ = d_2553___mcc_h6_
            d_2555_v24_: _dafny.Array
            nw438_ = _dafny.Array(_dafny.Map({}), 15)
            d_2555_v24_ = nw438_
            d_2556_v25_: _dafny.Map
            d_2556_v25_ = _dafny.Map({p1: d_2555_v24_})
            d_2557_v26_: C4
            nw439_ = C4()
            nw439_.ctor__(((d_2556_v25_)[p1] if (p1) in (d_2556_v25_) else d_2555_v24_))
            d_2557_v26_ = nw439_
            d_2558_v27_: _dafny.Array
            def lambda133_(d_2559_p0_):
                def lambda134_(d_2560_i1_):
                    return default__.safeModuloInt(d_2560_i1_, d_2559_p0_)

                return lambda134_

            init72_ = lambda133_(p0)
            nw440_ = _dafny.Array(None, 26)
            for i0_72_ in range(nw440_.length(0)):
                nw440_[i0_72_] = init72_(i0_72_)
            d_2558_v27_ = nw440_
            index463_ = default__.safeIndex(578, (d_2558_v27_).length(0))
            (d_2558_v27_)[index463_] = (p0) - (p0)
            d_2561_v28_: _dafny.Array
            nw441_ = _dafny.Array(None, 11)
            nw441_[int(0)] = d_2558_v27_
            nw441_[int(1)] = d_2558_v27_
            nw441_[int(2)] = d_2558_v27_
            nw441_[int(3)] = d_2558_v27_
            nw441_[int(4)] = d_2558_v27_
            nw441_[int(5)] = d_2558_v27_
            nw441_[int(6)] = d_2558_v27_
            nw441_[int(7)] = d_2558_v27_
            nw441_[int(8)] = d_2558_v27_
            nw441_[int(9)] = d_2558_v27_
            nw441_[int(10)] = d_2558_v27_
            d_2561_v28_ = nw441_
            index464_ = default__.safeIndex(789, (d_2561_v28_).length(0))
            (d_2561_v28_)[index464_] = d_2558_v27_
            index465_ = default__.safeIndex(578, (d_2558_v27_).length(0))
            index466_ = default__.safeIndex(789, (d_2561_v28_).length(0))
            rhs448_ = p0
            rhs449_ = default__.safeDivisionInt((0) - (p0), -463)
            rhs450_ = (p0) - (p0)
            rhs451_ = d_2558_v27_
            lhs354_ = d_2558_v27_
            lhs355_ = default__.safeIndex(578, (d_2558_v27_).length(0))
            lhs356_ = globalState
            lhs357_ = globalState
            lhs358_ = d_2561_v28_
            lhs359_ = default__.safeIndex(789, (d_2561_v28_).length(0))
            lhs354_[lhs355_] = rhs448_
            lhs356_.f3 = rhs449_
            lhs357_.f8 = rhs450_
            lhs358_[lhs359_] = rhs451_
            (globalState).f4 = p1
            d_2562_v29_: bool
            d_2563_v30_: bool
            out29_: bool
            out30_: bool
            out29_, out30_ = (d_2557_v26_).m0(globalState)
            d_2562_v29_ = out29_
            d_2563_v30_ = out30_
        hi16_ = (p0) + (default__.fm6(p0, globalState))
        for d_2564_i2_ in range(p0, hi16_):
            d_2565_v31_: _dafny.Array
            nw442_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_2565_v31_ = nw442_
            nw443_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_2565_v31_ = nw443_
            d_2566_v32_: _dafny.Array
            def lambda135_(d_2567_i3_):
                return True

            init73_ = lambda135_
            nw444_ = _dafny.Array(None, 14)
            for i0_73_ in range(nw444_.length(0)):
                nw444_[i0_73_] = init73_(i0_73_)
            d_2566_v32_ = nw444_
            d_2566_v32_ = (d_2566_v32_ if not(p1) else d_2566_v32_)
            if p1:
                (globalState).f4 = p1
                (globalState).f3 = p0
                d_2568_v33_: _dafny.Array
                def lambda136_(d_2569_i2_):
                    def lambda137_(d_2570_i4_):
                        return default__.safeDivisionInt(d_2570_i4_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2569_i2_ for d_2571_i5_ in range(default__.abs(745))]))).cardinality)

                    return lambda137_

                init74_ = lambda136_(d_2564_i2_)
                nw445_ = _dafny.Array(None, 26)
                for i0_74_ in range(nw445_.length(0)):
                    nw445_[i0_74_] = init74_(i0_74_)
                d_2568_v33_ = nw445_
                d_2572_v34_: _dafny.Seq
                d_2572_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqbcks"))
                d_2573_v35_: _dafny.Seq
                d_2573_v35_ = _dafny.SeqWithoutIsStrInference([d_2572_v34_])
                d_2574_v36_: _dafny.Array
                d_2574_v36_ = d_2568_v33_
                rhs452_ = (_dafny.SeqWithoutIsStrInference([d_2572_v34_])) <= (d_2573_v35_)
                rhs453_ = (d_2574_v36_)
                lhs360_ = globalState
                lhs360_.f4 = rhs452_
                d_2568_v33_ = rhs453_
                d_2575_v37_: _dafny.Array
                nw446_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_2575_v37_ = nw446_
                d_2575_v37_ = d_2575_v37_
                (globalState).f4 = (_dafny.SeqWithoutIsStrInference([(d_2572_v34_)[default__.safeIndex(p0, len(d_2572_v34_))] for d_2576_i6_ in range(default__.abs(586))])) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2577_i7_ in range(default__.abs(738))]))
            elif True:
                d_2578_v38_: str
                d_2578_v38_ = _dafny.CodePoint('l')
                d_2579_v39_: _dafny.Seq
                d_2579_v39_ = _dafny.SeqWithoutIsStrInference([d_2578_v38_, d_2578_v38_])
                d_2580_v40_: D10
                d_2580_v40_ = D10_DC24(p0, p1, d_2579_v39_)
                d_2581_v41_: _dafny.Seq
                d_2581_v41_ = _dafny.SeqWithoutIsStrInference([d_2579_v39_, ((d_2580_v40_).cf33).set(default__.safeIndex(d_2564_i2_, len((d_2580_v40_).cf33)), d_2578_v38_)])
                d_2582_v42_: _dafny.Map
                d_2582_v42_ = _dafny.Map({default__.fm20(p1, globalState): (_dafny.SeqWithoutIsStrInference([d_2578_v38_, d_2578_v38_, d_2578_v38_])) + ((d_2581_v41_)[default__.safeIndex(d_2564_i2_, len(d_2581_v41_))])})
                d_2582_v42_ = (d_2582_v42_).set((0) - (default__.safeModuloInt(p0, p0)), d_2579_v39_)
                (globalState).f4 = p1
                d_2578_v38_ = (d_2578_v38_ if p1 else d_2578_v38_)
                d_2583_v43_: D25
                d_2583_v43_ = D25_DC52(d_2565_v31_)
                d_2584_v44_: _dafny.Array
                nw447_ = _dafny.Array(None, 27)
                nw447_[int(0)] = (d_2583_v43_).cf73
                nw447_[int(1)] = d_2565_v31_
                nw447_[int(2)] = d_2565_v31_
                nw447_[int(3)] = d_2565_v31_
                nw447_[int(4)] = d_2565_v31_
                nw447_[int(5)] = d_2565_v31_
                nw447_[int(6)] = d_2565_v31_
                nw447_[int(7)] = d_2565_v31_
                nw447_[int(8)] = (d_2565_v31_ if p1 else d_2565_v31_)
                nw447_[int(9)] = d_2565_v31_
                nw447_[int(10)] = d_2565_v31_
                nw447_[int(11)] = d_2565_v31_
                nw447_[int(12)] = d_2565_v31_
                nw447_[int(13)] = d_2565_v31_
                nw447_[int(14)] = (d_2565_v31_ if p1 else d_2565_v31_)
                nw447_[int(15)] = d_2565_v31_
                nw447_[int(16)] = d_2565_v31_
                nw447_[int(17)] = d_2565_v31_
                nw447_[int(18)] = d_2565_v31_
                nw447_[int(19)] = d_2565_v31_
                nw447_[int(20)] = d_2565_v31_
                nw447_[int(21)] = d_2565_v31_
                nw447_[int(22)] = d_2565_v31_
                nw447_[int(23)] = d_2565_v31_
                nw447_[int(24)] = (d_2583_v43_).cf73
                nw447_[int(25)] = d_2565_v31_
                nw447_[int(26)] = d_2565_v31_
                d_2584_v44_ = nw447_
                index467_ = default__.safeIndex(26, (d_2584_v44_).length(0))
                (d_2584_v44_)[index467_] = d_2565_v31_
                index468_ = default__.safeIndex(26, (d_2584_v44_).length(0))
                (d_2584_v44_)[index468_] = d_2565_v31_
                (globalState).f8 = (p0) + (d_2564_i2_)
            d_2585_v45_: _dafny.Set
            d_2585_v45_ = _dafny.Set({p1})
            index469_ = default__.safeIndex(292, (d_2566_v32_).length(0))
            (d_2566_v32_)[index469_] = (d_2585_v45_).isdisjoint(_dafny.Set({p1, p1, False}))
            index470_ = default__.safeIndex(292, (d_2566_v32_).length(0))
            (d_2566_v32_)[index470_] = p1
        r0 = default__.safeDivisionInt(357, (663) * (p0))
        d_2586_v46_: str
        d_2586_v46_ = _dafny.CodePoint('h')
        d_2587_v47_: _dafny.Map
        d_2587_v47_ = _dafny.Map({(p0) + (p0): d_2586_v46_})
        d_2587_v47_ = (d_2587_v47_).set(p0, (_dafny.CodePoint('a') if p1 else d_2586_v46_))
        d_2588_v48_: _dafny.Array
        nw448_ = _dafny.Array(None, 26)
        d_2588_v48_ = nw448_
        d_2589_v49_: _dafny.Seq
        d_2589_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msd"))
        d_2590_v50_: _dafny.Seq
        d_2590_v50_ = _dafny.SeqWithoutIsStrInference([d_2589_v49_])
        d_2591_v51_: D4
        d_2591_v51_ = D4_DC11(p1, p1, p0, d_2589_v49_)
        d_2592_v52_: D4
        d_2592_v52_ = D4_DC12(d_2591_v51_)
        d_2593_v53_: _dafny.Seq
        d_2593_v53_ = _dafny.SeqWithoutIsStrInference([d_2592_v52_, D4_DC12(d_2591_v51_)])
        d_2594_v54_: _dafny.Map
        d_2594_v54_ = _dafny.Map({d_2588_v48_: (d_2590_v50_)[default__.safeIndex(len(d_2593_v53_), len(d_2590_v50_))]})
        d_2594_v54_ = (d_2594_v54_) | ((d_2594_v54_) | (d_2594_v54_))
        (globalState).f4 = default__.fm49(p1, globalState)
        d_2595_v55_: _dafny.Set
        d_2595_v55_ = _dafny.Set({167, p0})
        d_2596_v56_: _dafny.Map
        d_2596_v56_ = _dafny.Map({default__.fm49(p1, globalState): len(d_2595_v55_)})
        d_2597_v57_: _dafny.MultiSet
        d_2597_v57_ = _dafny.MultiSet([d_2596_v56_, (_dafny.Map({p1: p0})).set(p1, p0)])
        d_2598_v58_: _dafny.Seq
        d_2598_v58_ = _dafny.SeqWithoutIsStrInference([(d_2597_v57_).cardinality])
        r0 = (p0) + ((d_2598_v58_)[default__.safeIndex(p0, len(d_2598_v58_))])
        return r0

